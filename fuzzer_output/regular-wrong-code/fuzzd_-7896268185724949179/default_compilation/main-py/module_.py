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
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(-842, 432):
                d_0_v0_: int = compr_0_
                if ((-842) <= (d_0_v0_)) and ((d_0_v0_) < (432)):
                    coll0_[default__.safeDivisionInt(d_0_v0_, 308)] = 166
            return _dafny.Map(coll0_)
        return (len(iife0_()
        )) + (len((D13_DC26(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gtqau")), True, True)).cf33))

    @staticmethod
    def fm1(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "juk"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "liutef")))) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_1_i0_ in range(default__.abs(902))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "olrx"))))

    @staticmethod
    def fm2(p0, globalState):
        return not((True if False else (False if False else True)))

    @staticmethod
    def fm3(p0, p1, p2, p3, globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: int
            for compr_1_ in _dafny.IntegerRange(310, 336):
                d_2_v0_: int = compr_1_
                if ((310) <= (d_2_v0_)) and ((d_2_v0_) < (336)):
                    coll1_[default__.safeDivisionInt(d_2_v0_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_4_i0_ in range(default__.abs(-856))])))] = (D24_DC64(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_3_i1_ in range(default__.abs(468))])), True, False)).cf93
            return _dafny.Map(coll1_)
        return ((_dafny.Map({len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cacr"))), len(_dafny.Map({False: False})), 792})): not(True)})) | (_dafny.Map({581: False}))) | (iife1_()
        )

    @staticmethod
    def fm4(p0, p1, globalState):
        if (False if True else True):
            return D0_DC0(not(True))
        elif True:
            return D0_DC1(D0_DC0(False))

    @staticmethod
    def fm11(p0, globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(500, 743):
                d_5_v0_: int = compr_2_
                if ((500) <= (d_5_v0_)) and ((d_5_v0_) < (743)):
                    coll2_[default__.safeDivisionInt(d_5_v0_, 893)] = _dafny.CodePoint('v')
            return _dafny.Map(coll2_)
        source0_ = D20_DC49(iife2_()
)
        if source0_.is_DC50:
            d_6___mcc_h0_ = source0_.cf68
            d_7___mcc_h1_ = source0_.cf69
            d_8___mcc_h2_ = source0_.cf70
            d_9_cf70_ = d_8___mcc_h2_
            d_10_cf69_ = d_7___mcc_h1_
            d_11_cf68_ = d_6___mcc_h0_
            return _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_10_cf69_]), _dafny.MultiSet([False]), _dafny.MultiSet([d_9_cf70_])])
        elif source0_.is_DC51:
            return (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False])])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True, True])]))
        elif True:
            d_12___mcc_h3_ = source0_.cf67
            d_13_cf67_ = d_12___mcc_h3_
            return (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True]), _dafny.MultiSet([True])])) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True])]))

    @staticmethod
    def fm12(p0, p1, p2, globalState):
        return (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([872])).cardinality, (_dafny.MultiSet([_dafny.CodePoint('v'), _dafny.CodePoint('n')])).cardinality, -229, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True, not(True)]))])), 358])) + (_dafny.SeqWithoutIsStrInference([636])))) - (_dafny.MultiSet([864, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mci")))]))

    @staticmethod
    def fm13(p0, p1, p2, p3, globalState):
        return ((_dafny.MultiSet([_dafny.CodePoint('h'), (D3_DC6(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_14_i0_ in range(default__.abs(743))]), _dafny.CodePoint('c'), False)).cf7])).intersection(_dafny.MultiSet([_dafny.CodePoint('w'), _dafny.CodePoint('q')]))) | (_dafny.MultiSet([_dafny.CodePoint('a')]))

    @staticmethod
    def fm14(p0, p1, globalState):
        return ((_dafny.Set({False}) if True else _dafny.Set({True}))) - ((_dafny.Set({True})).intersection(_dafny.Set({False})))

    @staticmethod
    def fm15(p0, p1, globalState):
        return _dafny.CodePoint('b')

    @staticmethod
    def fm16(p0, p1, globalState):
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: str
            for compr_3_ in (_dafny.Map({_dafny.CodePoint('m'): (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tynvdcxoj")))]))).cardinality})).keys.Elements:
                d_15_v0_: str = compr_3_
                if (d_15_v0_) in (_dafny.Map({_dafny.CodePoint('m'): (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tynvdcxoj")))]))).cardinality})):
                    coll3_[d_15_v0_] = True
            return _dafny.Map(coll3_)
        return D4_DC7((0) - ((D19_DC47(D1_DC2(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False])): True})), (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: D17_DC39(838, _dafny.Map({len(_dafny.Map({False: 353})): len(iife3_()
)}), not(False), _dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): D13_DC25(_dafny.Map({-51: (_dafny.MultiSet([52])).cardinality}))}), 358)})) for d_16_i0_ in range(default__.abs(352))]))])).cardinality, False)).cf63))

    @staticmethod
    def fm18(p0, p1, p2, globalState):
        if (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tfjbncjt")))]))) >= (-845):
            return _dafny.SeqWithoutIsStrInference([False])
        elif True:
            return (_dafny.SeqWithoutIsStrInference([True, False])) + (_dafny.SeqWithoutIsStrInference([True]))

    @staticmethod
    def fm19(p0, p1, p2, p3, globalState):
        return ((_dafny.Set({False})) | (_dafny.Set({True}))).intersection(_dafny.Set({False, True}))

    @staticmethod
    def fm20(p0, p1, p2, p3, globalState):
        return (_dafny.Map({True: True})) | ((_dafny.Map({False: True})) | (_dafny.Map({False: True})))

    @staticmethod
    def fm21(globalState):
        return (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([-380, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hf")))]))).cardinality for d_17_i0_ in range(default__.abs(171))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_18_i1_ in range(default__.abs(41))]))]))

    @staticmethod
    def fm22(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-136]), _dafny.SeqWithoutIsStrInference([-24]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([906])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xtfitcgkw")))]), _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([-784 for d_19_i0_ in range(default__.abs(44))]))})))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vpemvormm"))) for d_20_i1_ in range(default__.abs(844))])])

    @staticmethod
    def fm23(p0, p1, p2, p3, globalState):
        if not((_dafny.SeqWithoutIsStrInference([True])) <= (_dafny.SeqWithoutIsStrInference([False, True]))):
            return D13_DC26(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vmuja")), not(not(False)), False)
        elif True:
            return D13_DC26((D13_DC26(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lhaqxlf")), True, False)).cf33, True, True)

    @staticmethod
    def fm24(p0, p1, p2, p3, globalState):
        def iife4_():
            coll4_ = _dafny.Map()
            compr_4_: str
            for compr_4_ in (_dafny.Set({_dafny.CodePoint('c')})).Elements:
                d_21_v0_: str = compr_4_
                if (d_21_v0_) in (_dafny.Set({_dafny.CodePoint('c')})):
                    coll4_[d_21_v0_] = False
            return _dafny.Map(coll4_)
        def iife5_():
            coll5_ = _dafny.Map()
            compr_5_: str
            for compr_5_ in (_dafny.Map({_dafny.CodePoint('t'): True})).keys.Elements:
                d_22_v1_: str = compr_5_
                if (d_22_v1_) in (_dafny.Map({_dafny.CodePoint('t'): True})):
                    coll5_[d_22_v1_] = False
            return _dafny.Map(coll5_)
        return ((_dafny.MultiSet([_dafny.Map({_dafny.CodePoint('k'): not(True)})])).intersection(_dafny.MultiSet([_dafny.Map({_dafny.CodePoint('l'): not(True)}), _dafny.Map({_dafny.CodePoint('j'): False}), _dafny.Map({_dafny.CodePoint('j'): False})]))).intersection((_dafny.MultiSet([_dafny.Map({_dafny.CodePoint('n'): False}), _dafny.Map({_dafny.CodePoint('p'): True}), iife4_()
        , _dafny.Map({_dafny.CodePoint('j'): False})])) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([iife5_()
 for d_23_i0_ in range(default__.abs(922))]))))

    @staticmethod
    def fm25(p0, p1, p2, p3, globalState):
        return D3_DC6(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qs")), _dafny.CodePoint('p'), (_dafny.MultiSet([False])).ispropersubset(_dafny.MultiSet([True])))

    @staticmethod
    def fm26(p0, p1, globalState):
        return ((_dafny.Set({8, 496, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([-312, (_dafny.MultiSet([582, len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fre")))})), 784])).cardinality]))).cardinality})) - (_dafny.Set({(0) - (len((_dafny.Set({True, False})))), 386}))) - (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([False, True])).cardinality])), 190, -913}))

    @staticmethod
    def fm27(p0, p1, globalState):
        if (_dafny.Set({False, False, True})).issubset(_dafny.Set({True})):
            return _dafny.CodePoint('v')
        elif True:
            return _dafny.CodePoint('p')

    @staticmethod
    def fm28(p0, p1, p2, p3, globalState):
        source1_ = D18_DC42(_dafny.Map({True: False}))
        if source1_.is_DC43:
            return _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_24_i0_ in range(default__.abs(642))])), -885, len(_dafny.Set({len(_dafny.Map({True: False})), (0) - (len(_dafny.Map({True: False}))), 538, 588}))])
        elif source1_.is_DC44:
            return _dafny.MultiSet([158, len(_dafny.SeqWithoutIsStrInference([False]))])
        elif True:
            d_25___mcc_h0_ = source1_.cf58
            d_26_cf58_ = d_25___mcc_h0_
            return _dafny.MultiSet([(0) - (-301), -858])

    @staticmethod
    def fm30(p0, p1, p2, p3, globalState):
        return ((_dafny.Set({915, 103})).intersection(_dafny.Set({833}))).intersection(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rsvm"))) for d_27_i0_ in range(default__.abs(368))]))}))

    @staticmethod
    def fm31(p0, p1, globalState):
        return ((_dafny.Set({True})).intersection(_dafny.Set({True}))).intersection(_dafny.Set({False, True}))

    @staticmethod
    def fm32(p0, p1, p2, globalState):
        if (True) == (True):
            return D14_DC28(_dafny.Set({396}))
        elif True:
            def iife6_():
                coll6_ = _dafny.Set()
                compr_6_: int
                for compr_6_ in _dafny.IntegerRange(88, 346):
                    d_28_v0_: int = compr_6_
                    if ((88) <= (d_28_v0_)) and ((d_28_v0_) < (346)):
                        coll6_ = coll6_.union(_dafny.Set([(d_28_v0_) + (len(_dafny.SeqWithoutIsStrInference([610])))]))
                return _dafny.Set(coll6_)
            return D14_DC28(iife6_()
)
        elif True:
            def iife7_():
                coll7_ = _dafny.Set()
                compr_7_: int
                for compr_7_ in _dafny.IntegerRange(896, 614):
                    d_29_v1_: int = compr_7_
                    if ((896) <= (d_29_v1_)) and ((d_29_v1_) < (614)):
                        coll7_ = coll7_.union(_dafny.Set([(d_29_v1_) - (-280)]))
                return _dafny.Set(coll7_)
            return D14_DC28(iife7_()
)

    @staticmethod
    def fm33(globalState):
        return (D25_DC67(_dafny.MultiSet([-892, 188]))).cf98

    @staticmethod
    def fm34(globalState):
        return _dafny.CodePoint('u')

    @staticmethod
    def fm35(p0, p1, globalState):
        source2_ = D21_DC54(262, 529, not(not(True)))
        if source2_.is_DC53:
            d_30___mcc_h0_ = source2_.cf72
            d_31___mcc_h1_ = source2_.cf73
            d_32___mcc_h2_ = source2_.cf74
            d_33___mcc_h3_ = source2_.cf75
            d_34_cf75_ = d_33___mcc_h3_
            d_35_cf74_ = d_32___mcc_h2_
            d_36_cf73_ = d_31___mcc_h1_
            d_37_cf72_ = d_30___mcc_h0_
            return _dafny.Map({d_36_cf73_: 449})
        elif source2_.is_DC54:
            d_38___mcc_h4_ = source2_.cf76
            d_39___mcc_h5_ = source2_.cf77
            d_40___mcc_h6_ = source2_.cf78
            d_41_cf78_ = d_40___mcc_h6_
            d_42_cf77_ = d_39___mcc_h5_
            d_43_cf76_ = d_38___mcc_h4_
            return (_dafny.Map({d_41_cf78_: d_42_cf77_})) | (_dafny.Map({d_41_cf78_: d_42_cf77_}))
        elif source2_.is_DC52:
            d_44___mcc_h7_ = source2_.cf71
            d_45_cf71_ = d_44___mcc_h7_
            if True:
                def iife8_():
                    coll8_ = _dafny.Map()
                    compr_8_: str
                    for compr_8_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b'), _dafny.CodePoint('b')])).Elements:
                        d_46_v0_: str = compr_8_
                        if (d_46_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b'), _dafny.CodePoint('b')])):
                            coll8_[d_46_v0_] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_47_i0_ in range(default__.abs(-442))]))
                    return _dafny.Map(coll8_)
                return _dafny.Map({False: len(iife8_()
                )})
            elif True:
                return _dafny.Map({False: 278})
        elif True:
            d_48___mcc_h8_ = source2_.cf79
            d_49_cf79_ = d_48___mcc_h8_
            return (_dafny.Map({not(True): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "abldcs")))})) | (_dafny.Map({not(False): 542}))

    @staticmethod
    def fm36(p0, p1, p2, globalState):
        return _dafny.Map({_dafny.SeqWithoutIsStrInference([False]): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_50_i0_ in range(default__.abs(-663))])})

    @staticmethod
    def fm37(globalState):
        return _dafny.SeqWithoutIsStrInference([(0) - (-504)])

    @staticmethod
    def fm38(globalState):
        def iife9_():
            def iife11_():
                coll11_ = _dafny.Set()
                compr_11_: _dafny.Set
                for compr_11_ in (_dafny.SeqWithoutIsStrInference([_dafny.Set({_dafny.CodePoint('a')}), _dafny.Set({_dafny.CodePoint('j')})])).Elements:
                    d_53_v0_: _dafny.Set = compr_11_
                    if (d_53_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.Set({_dafny.CodePoint('a')}), _dafny.Set({_dafny.CodePoint('j')})])):
                        coll11_ = coll11_.union(_dafny.Set([d_53_v0_]))
                return _dafny.Set(coll11_)
            coll9_ = _dafny.Set()
            def iife10_():
                coll10_ = _dafny.Set()
                compr_10_: _dafny.Set
                for compr_10_ in (_dafny.SeqWithoutIsStrInference([_dafny.Set({_dafny.CodePoint('a')}), _dafny.Set({_dafny.CodePoint('j')})])).Elements:
                    d_51_v0_: _dafny.Set = compr_10_
                    if (d_51_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.Set({_dafny.CodePoint('a')}), _dafny.Set({_dafny.CodePoint('j')})])):
                        coll10_ = coll10_.union(_dafny.Set([d_51_v0_]))
                return _dafny.Set(coll10_)
            compr_9_: _dafny.Set
            for compr_9_ in (iife10_()
            ).Elements:
                d_52_v1_: _dafny.Set = compr_9_
                if (d_52_v1_) in (iife11_()
                ):
                    coll9_ = coll9_.union(_dafny.Set([d_52_v1_]))
            return _dafny.Set(coll9_)
        return (_dafny.Set({_dafny.Set({_dafny.CodePoint('i'), _dafny.CodePoint('g')}), _dafny.Set({_dafny.CodePoint('v')}), _dafny.Set({_dafny.CodePoint('p'), _dafny.CodePoint('y'), _dafny.CodePoint('w'), _dafny.CodePoint('e')})})).intersection(iife9_()
        )

    @staticmethod
    def fm39(p0, p1, p2, globalState):
        def iife12_():
            coll12_ = _dafny.Map()
            compr_12_: _dafny.Seq
            for compr_12_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-784]), _dafny.SeqWithoutIsStrInference([508, (0) - (len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o'), _dafny.CodePoint('l')])), 486})))])])).Elements:
                d_56_v0_: _dafny.Seq = compr_12_
                if (d_56_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([-784]), _dafny.SeqWithoutIsStrInference([508, (0) - (len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o'), _dafny.CodePoint('l')])), 486})))])])):
                    coll12_[d_56_v0_] = _dafny.SeqWithoutIsStrInference([False])
            return _dafny.Map(coll12_)
        def iife13_():
            coll13_ = _dafny.Map()
            compr_13_: _dafny.Seq
            for compr_13_ in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([317]), _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([(_dafny.MultiSet([960])).cardinality])).cardinality for d_57_i2_ in range(default__.abs(267))])])).Elements:
                d_58_v1_: _dafny.Seq = compr_13_
                if (d_58_v1_) in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([317]), _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([(_dafny.MultiSet([960])).cardinality])).cardinality for d_57_i2_ in range(default__.abs(267))])])):
                    coll13_[d_58_v1_] = _dafny.SeqWithoutIsStrInference([True])
            return _dafny.Map(coll13_)
        return ((_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_54_i0_ in range(default__.abs(-361))])), len(_dafny.Map({False: _dafny.SeqWithoutIsStrInference([476 for d_55_i1_ in range(default__.abs(336))])})), len(_dafny.SeqWithoutIsStrInference([759]))]): _dafny.SeqWithoutIsStrInference([not(True)])})) | (iife12_()
        )) | ((_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: (0) - (len(_dafny.Map({True: False})))}))]): _dafny.SeqWithoutIsStrInference([True, False])})) | (iife13_()
        ))

    @staticmethod
    def fm40(globalState):
        return D16_DC36(True, (_dafny.MultiSet([False, False])) | (_dafny.MultiSet([True])), (-574) != (499))

    @staticmethod
    def fm41(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([not(True)])

    @staticmethod
    def fm42(globalState):
        return D4_DC10((0) - ((0) - ((D20_DC50(len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uh")))])), False, True)).cf68)), (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "owettajm")))) * (len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True])]))), (False if not(False) else not(not(False))), not(True), (_dafny.MultiSet([False])).cardinality)

    @staticmethod
    def fm43(p0, p1, p2, p3, globalState):
        def iife14_():
            coll14_ = _dafny.Set()
            compr_14_: str
            for compr_14_ in ((D26_DC71(_dafny.MultiSet([_dafny.CodePoint('p')]))).cf100).Elements:
                d_59_v0_: str = compr_14_
                if (d_59_v0_) in ((D26_DC71(_dafny.MultiSet([_dafny.CodePoint('p')]))).cf100):
                    coll14_ = coll14_.union(_dafny.Set([d_59_v0_]))
            return _dafny.Set(coll14_)
        return (_dafny.Map({not(True): D4_DC10(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "memjruqa"))), -160, True, True, (_dafny.MultiSet([404, len(_dafny.Map({885: iife14_()
})), 955, 236])).cardinality)})) | ((_dafny.Map({False: D4_DC10(285, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "exg")))), True, True, 113)})) | (_dafny.Map({False: D4_DC10(655, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))), True, True, 849)})))

    @staticmethod
    def fm44(globalState):
        if not(True):
            return D13_DC27(True)
        elif True:
            return D13_DC27(True)

    @staticmethod
    def fm45(globalState):
        if False:
            def iife15_():
                coll15_ = _dafny.Map()
                compr_15_: int
                for compr_15_ in _dafny.IntegerRange(190, -939):
                    d_60_v0_: int = compr_15_
                    if ((190) <= (d_60_v0_)) and ((d_60_v0_) < (-939)):
                        coll15_[(d_60_v0_) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({743, -433})), len(_dafny.SeqWithoutIsStrInference([226, 423])), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_61_i0_ in range(default__.abs(622))]))]))).cardinality)] = -545
                return _dafny.Map(coll15_)
            return D13_DC25(iife15_()
)
        elif True:
            return D13_DC25(_dafny.Map({(0) - ((0) - (len(_dafny.Map({True: True})))): 652}))

    @staticmethod
    def fm46(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference([D14_DC30(_dafny.SeqWithoutIsStrInference([691, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_62_i1_ in range(default__.abs(-832))])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xxqknse")))]), _dafny.SeqWithoutIsStrInference([False])) for d_63_i0_ in range(default__.abs(323))])) + (_dafny.SeqWithoutIsStrInference([D14_DC30(_dafny.SeqWithoutIsStrInference([94]), _dafny.SeqWithoutIsStrInference([not(False), True, True])) for d_64_i2_ in range(default__.abs(-332))]))

    @staticmethod
    def fm47(p0, globalState):
        return (_dafny.Set({True, False}) if True else _dafny.Set({False}))

    @staticmethod
    def fm48(p0, p1, globalState):
        def iife16_():
            coll16_ = _dafny.Map()
            compr_16_: str
            for compr_16_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a'), _dafny.CodePoint('l')])).Elements:
                d_65_v0_: str = compr_16_
                if (d_65_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a'), _dafny.CodePoint('l')])):
                    coll16_[d_65_v0_] = True
            return _dafny.Map(coll16_)
        return _dafny.Map({(-648) - (583): (iife16_()
        ) | (_dafny.Map({_dafny.CodePoint('j'): False}))})

    @staticmethod
    def fm49(p0, p1, p2, p3, globalState):
        return D19_DC46((_dafny.Map({D23_DC60(_dafny.Map({False: _dafny.Map({False: not(not(False))})})): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "unbxcg")))})) != (_dafny.Map({D23_DC60(_dafny.Map({not(False): _dafny.Map({not(True): not(False)})})): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_66_i0_ in range(default__.abs(366))]))})), (len(_dafny.Map({len(_dafny.Set({D14_DC28(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([True]))}))})): False}))) * (711))

    @staticmethod
    def fm50(p0, p1, globalState):
        return (_dafny.Map({not(False): _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "obwgo")))])})) | (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([len(_dafny.Set({False, not(True)})) for d_67_i0_ in range(default__.abs(40))])}))

    @staticmethod
    def fm51(p0, p1, globalState):
        def iife17_():
            coll17_ = _dafny.Map()
            compr_17_: int
            for compr_17_ in _dafny.IntegerRange(620, 233):
                d_68_v0_: int = compr_17_
                if ((620) <= (d_68_v0_)) and ((d_68_v0_) < (233)):
                    coll17_[(d_68_v0_) - ((0) - (-232))] = not(True)
            return _dafny.Map(coll17_)
        return (_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([428]))})), len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dgjnixnjt"))), 0})), len(_dafny.Map({len(_dafny.Map({532: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "umwrjbo")))})): True}))])): True})])) + (_dafny.SeqWithoutIsStrInference([iife17_()
        ]))

    @staticmethod
    def fm52(p0, p1, globalState):
        return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(not(False)), False, True]))).intersection((_dafny.MultiSet([not(False)])) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))))

    @staticmethod
    def fm53(p0, globalState):
        return ((_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yuclgsvm")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))}) if (D23_DC61(_dafny.SeqWithoutIsStrInference([True]), len(_dafny.Map({True: _dafny.CodePoint('b')})), False)).cf88 else _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ndjpjd")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jgu")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))}))).intersection(_dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_69_i0_ in range(default__.abs(986))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "clwvk"))}))

    @staticmethod
    def fm54(globalState):
        return (D13_DC25(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([(D3_DC5(_dafny.CodePoint('n'))).cf5 for d_70_i0_ in range(default__.abs(14))])): len(_dafny.SeqWithoutIsStrInference([False]))}))).cf32

    @staticmethod
    def fm55(p0, p1, p2, globalState):
        return (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvrqpnj")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "db")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_71_i0_ in range(default__.abs(768))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hs"))])) - ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sfwcvbt"))])).intersection(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dntsofjge")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ctwuqpd"))])))

    @staticmethod
    def fm56(p0, p1, p2, globalState):
        return ((_dafny.MultiSet([D6_DC13(613, 12)])) - (_dafny.MultiSet([D6_DC13(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "itggvjgu"))), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_72_i0_ in range(default__.abs(36))])))]))) | (_dafny.MultiSet([D6_DC13(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wmqdtf"))), -359), D6_DC13((0) - (len(_dafny.Set({True, False, True}))), (0) - (len(_dafny.Map({True: _dafny.SeqWithoutIsStrInference([False])}))))]))

    @staticmethod
    def fm57(globalState):
        if (_dafny.MultiSet([564])).issubset(_dafny.MultiSet([519, len(_dafny.Set({43})), len(_dafny.Set({934, 879}))])):
            return D16_DC35(_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.Set({-867})), (D22_DC57(-174, len(_dafny.Set({False})), False)).cf81]): _dafny.SeqWithoutIsStrInference([True])}))
        elif True:
            return D16_DC35(_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.Map({553: True}))]): _dafny.SeqWithoutIsStrInference([False, False])}))

    @staticmethod
    def fm58(p0, globalState):
        return D12_DC24(D12_DC22(_dafny.Map({False: len(_dafny.Map({True: (0) - (len(_dafny.SeqWithoutIsStrInference([True])))}))})))

    @staticmethod
    def fm59(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mjaaqd"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bojv"))) for d_73_i0_ in range(default__.abs(-859))])

    @staticmethod
    def fm60(globalState):
        def iife18_():
            def iife21_():
                coll21_ = _dafny.Set()
                compr_21_: int
                for compr_21_ in _dafny.IntegerRange(805, 429):
                    d_77_v1_: int = compr_21_
                    if ((805) <= (d_77_v1_)) and ((d_77_v1_) < (429)):
                        coll21_ = coll21_.union(_dafny.Set([(d_77_v1_) + (-564)]))
                return _dafny.Set(coll21_)
            def iife22_():
                coll22_ = _dafny.Map()
                compr_22_: int
                for compr_22_ in _dafny.IntegerRange(754, 47):
                    d_75_v2_: int = compr_22_
                    if ((754) <= (d_75_v2_)) and ((d_75_v2_) < (47)):
                        coll22_[(d_75_v2_) - (len(_dafny.SeqWithoutIsStrInference([False, False])))] = _dafny.CodePoint('d')
                return _dafny.Map(coll22_)
            coll18_ = _dafny.Map()
            def iife19_():
                coll19_ = _dafny.Set()
                compr_19_: int
                for compr_19_ in _dafny.IntegerRange(805, 429):
                    d_74_v1_: int = compr_19_
                    if ((805) <= (d_74_v1_)) and ((d_74_v1_) < (429)):
                        coll19_ = coll19_.union(_dafny.Set([(d_74_v1_) + (-564)]))
                return _dafny.Set(coll19_)
            def iife20_():
                coll20_ = _dafny.Map()
                compr_20_: int
                for compr_20_ in _dafny.IntegerRange(754, 47):
                    d_75_v2_: int = compr_20_
                    if ((754) <= (d_75_v2_)) and ((d_75_v2_) < (47)):
                        coll20_[(d_75_v2_) - (len(_dafny.SeqWithoutIsStrInference([False, False])))] = _dafny.CodePoint('d')
                return _dafny.Map(coll20_)
            compr_18_: _dafny.Map
            for compr_18_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wodnahax"))): _dafny.CodePoint('s')}), _dafny.Map({len(_dafny.Map({len(iife19_()
            ): len(_dafny.Map({True: len(_dafny.Map({-498: len(_dafny.SeqWithoutIsStrInference([True]))}))}))})): _dafny.CodePoint('q')}), _dafny.Map({len(_dafny.SeqWithoutIsStrInference([131, 105])): _dafny.CodePoint('f')}), iife20_()
            , _dafny.Map({-540: _dafny.CodePoint('o')})])).Elements:
                d_76_v0_: _dafny.Map = compr_18_
                if (d_76_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wodnahax"))): _dafny.CodePoint('s')}), _dafny.Map({len(_dafny.Map({len(iife21_()
                ): len(_dafny.Map({True: len(_dafny.Map({-498: len(_dafny.SeqWithoutIsStrInference([True]))}))}))})): _dafny.CodePoint('q')}), _dafny.Map({len(_dafny.SeqWithoutIsStrInference([131, 105])): _dafny.CodePoint('f')}), iife22_()
                , _dafny.Map({-540: _dafny.CodePoint('o')})])):
                    coll18_[d_76_v0_] = _dafny.Map({False: not(False)})
            return _dafny.Map(coll18_)
        return (iife18_()
        ) | (_dafny.Map({_dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, not(True)]))).cardinality: _dafny.CodePoint('a')}): _dafny.Map({False: True})}))

    @staticmethod
    def fm61(p0, p1, p2, globalState):
        return _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(_dafny.Map({False: True})) | (_dafny.Map({False: False})) for d_78_i0_ in range(default__.abs(455))]))

    @staticmethod
    def fm62(p0, globalState):
        if True:
            return D18_DC43()
        elif True:
            return D18_DC43()

    @staticmethod
    def m0(p0, p1, p2, p3, globalState):
        d_79_v0_: _dafny.Seq
        d_79_v0_ = _dafny.SeqWithoutIsStrInference([p0])
        d_80_v1_: _dafny.Seq
        d_80_v1_ = _dafny.SeqWithoutIsStrInference([d_79_v0_])
        d_81_i0_: int
        d_81_i0_ = 0
        with _dafny.label("0"):
            while (((d_80_v1_)[default__.safeIndex(p1, len(d_80_v1_))]) + (_dafny.SeqWithoutIsStrInference([p0, p0]))) < (d_79_v0_):
                with _dafny.c_label("0"):
                    if (d_81_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_81_i0_ = (d_81_i0_) + (1)
                    d_82_v2_: str
                    d_82_v2_ = _dafny.CodePoint('f')
                    d_83_v3_: D3
                    d_83_v3_ = D3_DC5(d_82_v2_)
                    d_84_v4_: _dafny.Map
                    d_84_v4_ = _dafny.Map({p1: d_83_v3_})
                    d_84_v4_ = (d_84_v4_).set(p3, d_83_v3_)
                    d_85_v5_: _dafny.Array
                    nw0_ = _dafny.Array(False, 13)
                    d_85_v5_ = nw0_
                    d_86_v6_: _dafny.Array
                    def lambda0_(d_87_i1_):
                        return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rphheln")) for d_88_i2_ in range(default__.abs(261))])

                    init0_ = lambda0_
                    nw1_ = _dafny.Array(None, 23)
                    for i0_0_ in range(nw1_.length(0)):
                        nw1_[i0_0_] = init0_(i0_0_)
                    d_86_v6_ = nw1_
                    d_89_v7_: _dafny.Seq
                    d_89_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hne"))
                    index0_ = default__.safeIndex(422, (d_86_v6_).length(0))
                    (d_86_v6_)[index0_] = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lgtdubw")), d_89_v7_, d_89_v7_])
                    index1_ = default__.safeIndex(760, (d_85_v5_).length(0))
                    (d_85_v5_)[index1_] = (d_79_v0_)[default__.safeIndex(p1, len(d_79_v0_))]
                    d_90_v8_: _dafny.Seq
                    d_90_v8_ = _dafny.SeqWithoutIsStrInference([d_89_v7_])
                    d_91_v9_: _dafny.MultiSet
                    d_91_v9_ = _dafny.MultiSet([not(p0), p0])
                    index2_ = default__.safeIndex(422, (d_86_v6_).length(0))
                    index3_ = default__.safeIndex(760, (d_85_v5_).length(0))
                    rhs0_ = d_85_v5_
                    rhs1_ = p1
                    rhs2_ = (p1) * ((p1) + (len(d_79_v0_)))
                    rhs3_ = (((_dafny.SeqWithoutIsStrInference([d_89_v7_])) + (d_90_v8_)) + ((_dafny.SeqWithoutIsStrInference([d_89_v7_ for d_92_i3_ in range(default__.abs(674))]) if p0 else (d_90_v8_).set(default__.safeIndex(p1, len(d_90_v8_)), d_89_v7_)))).set(default__.safeIndex(p1, len(((_dafny.SeqWithoutIsStrInference([d_89_v7_])) + (d_90_v8_)) + ((_dafny.SeqWithoutIsStrInference([d_89_v7_ for d_93_i3_ in range(default__.abs(674))]) if p0 else (d_90_v8_).set(default__.safeIndex(p1, len(d_90_v8_)), d_89_v7_))))), (d_89_v7_ if p0 else d_89_v7_))
                    rhs4_ = ((_dafny.MultiSet([p0])) - (d_91_v9_)).issubset(d_91_v9_)
                    lhs0_ = globalState
                    lhs1_ = globalState
                    lhs2_ = d_86_v6_
                    lhs3_ = default__.safeIndex(422, (d_86_v6_).length(0))
                    lhs4_ = d_85_v5_
                    lhs5_ = default__.safeIndex(760, (d_85_v5_).length(0))
                    d_85_v5_ = rhs0_
                    lhs0_.f7 = rhs1_
                    lhs1_.f7 = rhs2_
                    lhs2_[lhs3_] = rhs3_
                    lhs4_[lhs5_] = rhs4_
                    (globalState).f13 = ((d_85_v5_)[default__.safeIndex(760, (d_85_v5_).length(0))]) == ((d_79_v0_)[default__.safeIndex(p3, len(d_79_v0_))])
                    d_94_v10_: _dafny.Seq
                    d_94_v10_ = _dafny.SeqWithoutIsStrInference([p1])
                    (globalState).f16 = (d_94_v10_)[default__.safeIndex(947, len(d_94_v10_))]
                    pass
            pass
        (globalState).f16 = p1
        d_95_v11_: _dafny.Array
        nw2_ = _dafny.Array(_dafny.Map({}), 11)
        d_95_v11_ = nw2_
        d_96_v13_: _dafny.Seq
        d_96_v13_ = _dafny.SeqWithoutIsStrInference([default__.fm1(712, globalState)])
        index4_ = default__.safeIndex(312, (d_95_v11_).length(0))
        def iife23_():
            coll23_ = _dafny.Map()
            compr_23_: _dafny.Seq
            for compr_23_ in (d_96_v13_).Elements:
                d_97_v12_: _dafny.Seq = compr_23_
                if (d_97_v12_) in (d_96_v13_):
                    coll23_[d_97_v12_] = p1
            return _dafny.Map(coll23_)
        (d_95_v11_)[index4_] = iife23_()
        
        index5_ = default__.safeIndex(312, (d_95_v11_).length(0))
        (d_95_v11_)[index5_] = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_98_i4_ in range(default__.abs(209))]): p1})
        pat_let_tv0_ = p1
        pat_let_tv1_ = p0
        pat_let_tv2_ = p0
        def lambda1_(source4_):
            if source4_.is_DC0:
                d_99___mcc_h2_ = source4_.cf0
                d_100_cf0_ = d_99___mcc_h2_
                return D1_DC2(_dafny.Map({pat_let_tv0_: pat_let_tv1_}))
            elif True:
                d_101___mcc_h3_ = source4_.cf1
                d_102_cf1_ = d_101___mcc_h3_
                return D1_DC2(_dafny.Map({845: pat_let_tv2_}))

        source3_ = lambda1_(D0_DC0(p0))
        if source3_.is_DC3:
            d_103___mcc_h0_ = source3_.cf3
            d_104_cf3_ = d_103___mcc_h0_
            d_105_v14_: _dafny.Array
            nw3_ = _dafny.Array(False, 18)
            d_105_v14_ = nw3_
            index6_ = default__.safeIndex(612, (d_105_v14_).length(0))
            (d_105_v14_)[index6_] = d_104_cf3_
            index7_ = default__.safeIndex(612, (d_105_v14_).length(0))
            (d_105_v14_)[index7_] = d_104_cf3_
            if (d_105_v14_)[default__.safeIndex(612, (d_105_v14_).length(0))]:
                d_106_v15_: _dafny.Array
                nw4_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 16)
                d_106_v15_ = nw4_
                d_107_v16_: str
                d_107_v16_ = _dafny.CodePoint('c')
                d_108_v17_: _dafny.Seq
                d_108_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tmcphea"))
                index8_ = default__.safeIndex(389, (d_106_v15_).length(0))
                (d_106_v15_)[index8_] = ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_109_i5_ in range(default__.abs(320))])).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_110_i5_ in range(default__.abs(320))]))), d_107_v16_)) + (d_108_v17_)
                d_111_v18_: _dafny.Set
                d_111_v18_ = _dafny.Set({p0})
                d_112_v19_: _dafny.Map
                d_112_v19_ = _dafny.Map({len(d_108_v17_): len(d_111_v18_)})
                d_113_v20_: _dafny.Map
                d_113_v20_ = _dafny.Map({True: False})
                d_114_v21_: _dafny.Set
                d_114_v21_ = _dafny.Set({_dafny.Map({p3: len(_dafny.Map({default__.fm2(d_113_v20_, globalState): p3}))}), _dafny.Map({p1: p3})})
                index9_ = default__.safeIndex(389, (d_106_v15_).length(0))
                (d_106_v15_)[index9_] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "on"))).set(default__.safeIndex(p3, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "on")))), d_107_v16_) if (_dafny.Set({d_112_v19_})) == (d_114_v21_) else d_108_v17_)
                d_113_v20_ = (d_113_v20_).set(d_104_cf3_, d_104_cf3_)
                (globalState).f20 = ((default__.fm1(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yootrwdif"))), globalState)) + ((d_106_v15_)[default__.safeIndex(389, (d_106_v15_).length(0))]) if (default__.fm0(p1, p3, len(d_113_v20_), (d_105_v14_)[default__.safeIndex(612, (d_105_v14_).length(0))], globalState)) < (p1) else ((d_108_v17_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r")))).set(default__.safeIndex(p1, len((d_108_v17_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))))), d_107_v16_))
                (globalState).f16 = (0) - (p3)
                (globalState).f2 = d_104_cf3_
            elif True:
                d_104_cf3_ = (d_104_cf3_ if p0 else False)
                (globalState).f0 = p0
                (globalState).f2 = not(False)
                d_115_v22_: _dafny.Map
                d_115_v22_ = _dafny.Map({(418) * ((0) - (p3)): p3})
                d_116_v23_: _dafny.Array
                def lambda2_(d_117_p0_):
                    def lambda3_(d_118_i6_):
                        return D0_DC0(d_117_p0_)

                    return lambda3_

                init1_ = lambda2_(p0)
                nw5_ = _dafny.Array(None, 21)
                for i0_1_ in range(nw5_.length(0)):
                    nw5_[i0_1_] = init1_(i0_1_)
                d_116_v23_ = nw5_
                d_119_v24_: _dafny.Map
                d_119_v24_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([p0])): not(p0)})
                d_120_v25_: D0
                d_120_v25_ = D0_DC0(((d_119_v24_)[147] if (147) in (d_119_v24_) else p0))
                pat_let_tv3_ = d_104_cf3_
                index10_ = default__.safeIndex(286, (d_116_v23_).length(0))
                def iife24_(_pat_let0_0):
                    def iife25_(d_121_dt__update__tmp_h0_):
                        def iife26_(_pat_let1_0):
                            def iife27_(d_122_dt__update_hcf0_h0_):
                                return D0_DC0(d_122_dt__update_hcf0_h0_)
                            return iife27_(_pat_let1_0)
                        return iife26_(pat_let_tv3_)
                    return iife25_(_pat_let0_0)
                (d_116_v23_)[index10_] = iife24_(d_120_v25_)
                d_123_v27_: _dafny.Array
                def lambda4_(d_124_v24_):
                    def lambda5_(d_125_i7_):
                        def iife28_():
                            coll24_ = _dafny.Set()
                            compr_24_: str
                            for compr_24_ in (_dafny.Map({_dafny.CodePoint('p'): d_124_v24_})).keys.Elements:
                                d_126_v26_: str = compr_24_
                                if (d_126_v26_) in (_dafny.Map({_dafny.CodePoint('p'): d_124_v24_})):
                                    coll24_ = coll24_.union(_dafny.Set([d_126_v26_]))
                            return _dafny.Set(coll24_)
                        return default__.safeModuloInt(d_125_i7_, len(iife28_()
                        ))

                    return lambda5_

                init2_ = lambda4_(d_119_v24_)
                nw6_ = _dafny.Array(None, 3)
                for i0_2_ in range(nw6_.length(0)):
                    nw6_[i0_2_] = init2_(i0_2_)
                d_123_v27_ = nw6_
                index11_ = default__.safeIndex(224, (d_123_v27_).length(0))
                (d_123_v27_)[index11_] = (0) - (p1)
                d_127_v28_: _dafny.Seq
                d_127_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ylg"))
                index12_ = default__.safeIndex(286, (d_116_v23_).length(0))
                index13_ = default__.safeIndex(612, (d_105_v14_).length(0))
                index14_ = default__.safeIndex(224, (d_123_v27_).length(0))
                index15_ = default__.safeIndex(612, (d_105_v14_).length(0))
                rhs5_ = (d_115_v22_).set((p3) - (-542), p3)
                rhs6_ = d_120_v25_
                rhs7_ = p0
                rhs8_ = (p1) - (p3)
                rhs9_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_128_i8_ in range(default__.abs(-830))])) < (d_127_v28_)
                lhs6_ = d_116_v23_
                lhs7_ = default__.safeIndex(286, (d_116_v23_).length(0))
                lhs8_ = d_105_v14_
                lhs9_ = default__.safeIndex(612, (d_105_v14_).length(0))
                lhs10_ = d_123_v27_
                lhs11_ = default__.safeIndex(224, (d_123_v27_).length(0))
                lhs12_ = d_105_v14_
                lhs13_ = default__.safeIndex(612, (d_105_v14_).length(0))
                d_115_v22_ = rhs5_
                lhs6_[lhs7_] = rhs6_
                lhs8_[lhs9_] = rhs7_
                lhs10_[lhs11_] = rhs8_
                lhs12_[lhs13_] = rhs9_
                index16_ = default__.safeIndex(612, (d_105_v14_).length(0))
                index17_ = default__.safeIndex(224, (d_123_v27_).length(0))
                rhs10_ = d_104_cf3_
                rhs11_ = (d_123_v27_)[default__.safeIndex(224, (d_123_v27_).length(0))]
                rhs12_ = default__.safeModuloInt((0) - (p3), p1)
                rhs13_ = d_123_v27_
                lhs14_ = d_105_v14_
                lhs15_ = default__.safeIndex(612, (d_105_v14_).length(0))
                lhs16_ = d_123_v27_
                lhs17_ = default__.safeIndex(224, (d_123_v27_).length(0))
                lhs18_ = globalState
                lhs14_[lhs15_] = rhs10_
                lhs16_[lhs17_] = rhs11_
                lhs18_.f7 = rhs12_
                d_123_v27_ = rhs13_
            d_129_v29_: _dafny.Array
            nw7_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 23)
            d_129_v29_ = nw7_
            d_130_v30_: _dafny.Seq
            d_130_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))
            index18_ = default__.safeIndex(916, (d_129_v29_).length(0))
            (d_129_v29_)[index18_] = (default__.fm1(p1, globalState)) + (d_130_v30_)
            index19_ = default__.safeIndex(916, (d_129_v29_).length(0))
            (d_129_v29_)[index19_] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_131_i9_ in range(default__.abs(90))])) + (d_130_v30_)
            index20_ = default__.safeIndex(612, (d_105_v14_).length(0))
            (d_105_v14_)[index20_] = True
        elif True:
            d_132___mcc_h1_ = source3_.cf2
            d_133_cf2_ = d_132___mcc_h1_
            (globalState).f16 = p3
            d_134_v31_: str
            d_134_v31_ = _dafny.CodePoint('q')
            d_135_v32_: _dafny.Map
            d_135_v32_ = _dafny.Map({(p3) + (p1): d_134_v31_})
            d_134_v31_ = ((d_135_v32_)[p3] if (p3) in (d_135_v32_) else d_134_v31_)
            d_136_v33_: _dafny.Map
            d_136_v33_ = _dafny.Map({p0: p0})
            d_137_v34_: T0
            nw8_ = C3()
            nw8_.ctor__()
            d_137_v34_ = nw8_
            d_138_v35_: _dafny.Array
            nw9_ = _dafny.Array(None, 18)
            nw9_[int(0)] = d_137_v34_
            nw9_[int(1)] = d_137_v34_
            nw9_[int(2)] = d_137_v34_
            nw9_[int(3)] = d_137_v34_
            nw9_[int(4)] = d_137_v34_
            nw9_[int(5)] = d_137_v34_
            nw9_[int(6)] = d_137_v34_
            nw9_[int(7)] = d_137_v34_
            nw9_[int(8)] = d_137_v34_
            nw9_[int(9)] = d_137_v34_
            nw9_[int(10)] = d_137_v34_
            nw9_[int(11)] = d_137_v34_
            nw9_[int(12)] = d_137_v34_
            nw9_[int(13)] = d_137_v34_
            nw9_[int(14)] = d_137_v34_
            nw9_[int(15)] = d_137_v34_
            nw9_[int(16)] = d_137_v34_
            nw9_[int(17)] = d_137_v34_
            d_138_v35_ = nw9_
            d_139_v36_: _dafny.Map
            d_139_v36_ = _dafny.Map({d_138_v35_: p0})
            d_140_v37_: _dafny.Map
            d_140_v37_ = _dafny.Map({default__.fm2(d_136_v33_, globalState): D22_DC56(d_139_v36_)})
            d_141_v38_: _dafny.Seq
            d_141_v38_ = _dafny.SeqWithoutIsStrInference([d_140_v37_])
            d_142_v39_: _dafny.Seq
            d_142_v39_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "frletcac"))
            d_143_v40_: C4
            nw10_ = C4()
            nw10_.ctor__(not(((d_141_v38_)[default__.safeIndex(len(d_142_v39_), len(d_141_v38_))]) != (d_140_v37_)), p0)
            d_143_v40_ = nw10_
            d_133_cf2_ = (d_133_cf2_).set((p3 if d_143_v40_.f25 else default__.fm0(len(d_79_v0_), 936, p3, d_143_v40_.f25, globalState)), False)
        d_144_v41_: C2
        nw11_ = C2()
        nw11_.ctor__(p1)
        d_144_v41_ = nw11_
        d_145_v42_: D18
        d_145_v42_ = D18_DC43()
        d_146_v43_: str
        d_146_v43_ = _dafny.CodePoint('l')
        d_147_v44_: _dafny.Map
        d_147_v44_ = _dafny.Map({p0: d_146_v43_})
        d_148_v45_: _dafny.Map
        d_148_v45_ = _dafny.Map({p0: len(d_147_v44_)})
        d_145_v42_ = default__.fm62(len((d_148_v45_) | (d_148_v45_)), globalState)

    @staticmethod
    def Main(noArgsParameter__):
        d_149_v0_: _dafny.Array
        def lambda6_(d_150_i0_):
            return (d_150_i0_) - (len(_dafny.Map({len(_dafny.Map({423: (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(False), not(not(True))]))).cardinality})): len(_dafny.Map({-866: 259}))})))

        init3_ = lambda6_
        nw12_ = _dafny.Array(None, 18)
        for i0_3_ in range(nw12_.length(0)):
            nw12_[i0_3_] = init3_(i0_3_)
        d_149_v0_ = nw12_
        d_151_v1_: _dafny.Seq
        d_151_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ucjarmpr"))
        d_152_v2_: bool
        d_152_v2_ = False
        d_153_v3_: _dafny.Set
        d_153_v3_ = _dafny.Set({d_151_v1_, d_151_v1_})
        d_154_v5_: int
        d_154_v5_ = 710
        d_155_v6_: str
        d_155_v6_ = _dafny.CodePoint('t')
        d_156_v7_: _dafny.Array
        nw13_ = _dafny.Array(_dafny.Map({}), 20)
        d_156_v7_ = nw13_
        d_157_globalState_: GlobalState
        nw14_ = GlobalState()
        def iife29_():
            coll25_ = _dafny.Set()
            compr_25_: _dafny.Seq
            for compr_25_ in (d_153_v3_).Elements:
                d_158_v4_: _dafny.Seq = compr_25_
                if (d_158_v4_) in (d_153_v3_):
                    coll25_ = coll25_.union(_dafny.Set([d_158_v4_]))
            return _dafny.Set(coll25_)
        nw14_.ctor__(True, d_149_v0_, False, 746, True, False, False, 729, d_151_v1_, -908, (d_153_v3_ if d_152_v2_ else iife29_()
        ), -466, 794, True, ((d_151_v1_).set(default__.safeIndex(d_154_v5_, len(d_151_v1_)), d_155_v6_)) + (d_151_v1_), 52, 251, False, 953, d_151_v1_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cesqe"))) + (d_151_v1_), False, d_156_v7_)
        d_157_globalState_ = nw14_
        d_159_v8_: _dafny.Seq
        d_159_v8_ = _dafny.SeqWithoutIsStrInference([d_154_v5_])
        d_154_v5_ = ((0) - (-504)) - ((d_159_v8_)[default__.safeIndex((d_159_v8_)[default__.safeIndex(d_154_v5_, len(d_159_v8_))], len(d_159_v8_))])
        d_160_v9_: _dafny.Array
        nw15_ = _dafny.Array(False, 7)
        d_160_v9_ = nw15_
        index21_ = default__.safeIndex(773, (d_160_v9_).length(0))
        (d_160_v9_)[index21_] = d_152_v2_
        index22_ = default__.safeIndex(773, (d_160_v9_).length(0))
        (d_160_v9_)[index22_] = d_152_v2_
        index23_ = default__.safeIndex(773, (d_160_v9_).length(0))
        (d_160_v9_)[index23_] = (d_152_v2_) and ((d_160_v9_)[default__.safeIndex(773, (d_160_v9_).length(0))])
        if d_152_v2_:
            (d_157_globalState_).f16 = d_154_v5_
            d_161_v10_: _dafny.Map
            d_161_v10_ = _dafny.Map({d_151_v1_: len(_dafny.Map({default__.fm0(727, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "faepiih"))), d_154_v5_, d_152_v2_, d_157_globalState_): d_152_v2_}))})
            d_162_v11_: _dafny.Map
            d_162_v11_ = _dafny.Map({d_151_v1_: len(d_161_v10_)})
            d_163_v12_: _dafny.Seq
            d_163_v12_ = _dafny.SeqWithoutIsStrInference([d_152_v2_])
            d_164_v13_: _dafny.Map
            d_164_v13_ = _dafny.Map({d_163_v12_: default__.fm1(d_154_v5_, d_157_globalState_)})
            d_165_v14_: _dafny.Seq
            d_165_v14_ = _dafny.SeqWithoutIsStrInference([d_163_v12_, d_163_v12_, d_163_v12_])
            default__.m0((d_160_v9_)[default__.safeIndex(773, (d_160_v9_).length(0))], len(d_162_v11_), (d_164_v13_).set((d_165_v14_)[default__.safeIndex(d_154_v5_, len(d_165_v14_))], d_151_v1_), (d_154_v5_ if d_152_v2_ else d_154_v5_), d_157_globalState_)
            (d_157_globalState_).f7 = (d_159_v8_)[default__.safeIndex(d_154_v5_, len(d_159_v8_))]
            d_166_v15_: D0
            d_166_v15_ = D0_DC0(d_152_v2_)
            d_167_v16_: D0
            d_167_v16_ = D0_DC1(D0_DC1(d_166_v15_))
            d_168_v17_: D0
            d_168_v17_ = D0_DC1(d_167_v16_)
            source5_ = d_168_v17_
            if source5_.is_DC0:
                d_169___mcc_h0_ = source5_.cf0
                d_170_cf0_ = d_169___mcc_h0_
                d_171_v18_: _dafny.Map
                d_171_v18_ = _dafny.Map({d_152_v2_: d_152_v2_})
                d_172_v19_: _dafny.Map
                d_172_v19_ = _dafny.Map({d_170_cf0_: ((d_171_v18_)[d_152_v2_] if (d_152_v2_) in (d_171_v18_) else d_152_v2_)})
                d_173_v20_: _dafny.Map
                d_173_v20_ = _dafny.Map({d_154_v5_: ((d_172_v19_)[d_152_v2_] if (d_152_v2_) in (d_172_v19_) else d_170_cf0_)})
                d_174_v21_: _dafny.Set
                d_174_v21_ = _dafny.Set({d_154_v5_, len(d_173_v20_)})
                default__.m0((d_174_v21_).ispropersubset(d_174_v21_), d_154_v5_, _dafny.Map({d_163_v12_: d_151_v1_}), d_154_v5_, d_157_globalState_)
                d_175_v22_: _dafny.Array
                def lambda7_(d_176_v20_):
                    def lambda8_(d_177_i1_):
                        return (D1_DC2(d_176_v20_)).cf2

                    return lambda8_

                init4_ = lambda7_(d_173_v20_)
                nw16_ = _dafny.Array(None, 2)
                for i0_4_ in range(nw16_.length(0)):
                    nw16_[i0_4_] = init4_(i0_4_)
                d_175_v22_ = nw16_
                d_178_v23_: _dafny.Map
                d_178_v23_ = _dafny.Map({d_154_v5_: d_154_v5_})
                d_179_v24_: _dafny.Seq
                d_179_v24_ = _dafny.SeqWithoutIsStrInference([d_178_v23_])
                d_180_v25_: _dafny.Seq
                d_180_v25_ = _dafny.SeqWithoutIsStrInference([(d_179_v24_) + (d_179_v24_)])
                d_181_v26_: _dafny.MultiSet
                d_181_v26_ = _dafny.MultiSet([d_149_v0_])
                d_182_v27_: _dafny.Map
                d_182_v27_ = _dafny.Map({d_170_cf0_: d_175_v22_})
                rhs14_ = (_dafny.MultiSet((d_180_v25_)[default__.safeIndex(d_154_v5_, len(d_180_v25_))])).cardinality
                rhs15_ = default__.fm2(_dafny.Map({d_152_v2_: (d_160_v9_)[default__.safeIndex(773, (d_160_v9_).length(0))]}), d_157_globalState_)
                rhs16_ = (d_175_v22_ if (d_163_v12_)[default__.safeIndex(((d_181_v26_)[d_149_v0_] if (d_149_v0_) in (d_181_v26_) else d_154_v5_), len(d_163_v12_))] else ((d_182_v27_)[d_152_v2_] if (d_152_v2_) in (d_182_v27_) else d_175_v22_))
                rhs17_ = default__.safeDivisionInt((d_154_v5_) + (d_154_v5_), d_154_v5_)
                rhs18_ = (d_160_v9_)[default__.safeIndex(773, (d_160_v9_).length(0))]
                lhs19_ = d_157_globalState_
                lhs20_ = d_157_globalState_
                lhs21_ = d_157_globalState_
                lhs22_ = d_157_globalState_
                lhs19_.f7 = rhs14_
                lhs20_.f2 = rhs15_
                d_175_v22_ = rhs16_
                lhs21_.f16 = rhs17_
                lhs22_.f2 = rhs18_
                d_183_v28_: _dafny.MultiSet
                d_183_v28_ = _dafny.MultiSet([d_171_v18_, _dafny.Map({(d_160_v9_)[default__.safeIndex(773, (d_160_v9_).length(0))]: d_152_v2_})])
                d_184_v29_: _dafny.MultiSet
                d_184_v29_ = _dafny.MultiSet([d_154_v5_, d_154_v5_])
                (d_157_globalState_).f16 = default__.fm0(d_154_v5_, (d_183_v28_).cardinality, (0) - (default__.safeModuloInt(d_154_v5_, d_154_v5_)), (d_184_v29_).ispropersubset(d_184_v29_), d_157_globalState_)
                (d_157_globalState_).f13 = (d_151_v1_) <= ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "urqdgor"))) + (_dafny.SeqWithoutIsStrInference([d_155_v6_ for d_185_i2_ in range(default__.abs(-451))])))
            elif True:
                d_186___mcc_h1_ = source5_.cf1
                d_187_cf1_ = d_186___mcc_h1_
                (d_157_globalState_).f21 = d_152_v2_
                d_154_v5_ = 439
                d_188_v30_: _dafny.Map
                d_188_v30_ = _dafny.Map({(d_160_v9_)[default__.safeIndex(773, (d_160_v9_).length(0))]: (d_160_v9_)[default__.safeIndex(773, (d_160_v9_).length(0))]})
                default__.m0(default__.fm2(d_188_v30_, d_157_globalState_), d_154_v5_, d_164_v13_, -198, d_157_globalState_)
                d_189_v31_: D1
                d_189_v31_ = D1_DC3((d_152_v2_) or ((d_160_v9_)[default__.safeIndex(773, (d_160_v9_).length(0))]))
                d_189_v31_ = d_189_v31_
            d_190_v32_: D0
            d_190_v32_ = D0_DC0((d_160_v9_)[default__.safeIndex(773, (d_160_v9_).length(0))])
            d_190_v32_ = d_190_v32_
        elif True:
            (d_157_globalState_).f21 = True
            d_151_v1_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bsi"))) + (d_151_v1_)
            d_152_v2_ = False
            index24_ = default__.safeIndex(441, (d_149_v0_).length(0))
            (d_149_v0_)[index24_] = (0) - (-193)
            d_191_v33_: _dafny.Set
            d_191_v33_ = _dafny.Set({d_152_v2_, d_152_v2_})
            d_192_v34_: _dafny.Map
            d_192_v34_ = _dafny.Map({len(d_191_v33_): (d_160_v9_)[default__.safeIndex(773, (d_160_v9_).length(0))]})
            d_193_v36_: D1
            def iife30_():
                coll26_ = _dafny.Map()
                compr_26_: int
                for compr_26_ in _dafny.IntegerRange(929, 635):
                    d_194_v35_: int = compr_26_
                    if ((929) <= (d_194_v35_)) and ((d_194_v35_) < (635)):
                        coll26_[default__.safeDivisionInt(d_194_v35_, d_154_v5_)] = (d_160_v9_)[default__.safeIndex(773, (d_160_v9_).length(0))]
                return _dafny.Map(coll26_)
            d_193_v36_ = D1_DC2((d_192_v34_) | (iife30_()
))
            index25_ = default__.safeIndex(441, (d_149_v0_).length(0))
            rhs19_ = (d_160_v9_)[default__.safeIndex(773, (d_160_v9_).length(0))]
            rhs20_ = ((_dafny.MultiSet([(0) - (d_154_v5_)])).cardinality if (d_152_v2_ if d_152_v2_ else (d_160_v9_)[default__.safeIndex(773, (d_160_v9_).length(0))]) else d_154_v5_)
            rhs21_ = d_193_v36_
            lhs23_ = d_157_globalState_
            lhs24_ = d_149_v0_
            lhs25_ = default__.safeIndex(441, (d_149_v0_).length(0))
            lhs23_.f21 = rhs19_
            lhs24_[lhs25_] = rhs20_
            d_193_v36_ = rhs21_
            d_195_v37_: _dafny.Array
            def lambda9_(d_196_v1_):
                def lambda10_(d_197_i3_):
                    return d_196_v1_

                return lambda10_

            init5_ = lambda9_(d_151_v1_)
            nw17_ = _dafny.Array(None, 11)
            for i0_5_ in range(nw17_.length(0)):
                nw17_[i0_5_] = init5_(i0_5_)
            d_195_v37_ = nw17_
            d_198_v38_: _dafny.Map
            d_198_v38_ = _dafny.Map({d_195_v37_: (d_151_v1_) + (d_151_v1_)})
            d_198_v38_ = (d_198_v38_).set(d_195_v37_, d_151_v1_)
        d_199_v39_: D0
        d_199_v39_ = D0_DC0(d_152_v2_)
        d_200_v40_: _dafny.Seq
        d_200_v40_ = _dafny.SeqWithoutIsStrInference([(d_160_v9_)[default__.safeIndex(773, (d_160_v9_).length(0))], False, d_152_v2_, False])
        d_201_v41_: _dafny.Map
        d_201_v41_ = _dafny.Map({d_200_v40_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ncojfqse"))})
        default__.m0((d_199_v39_).cf0, d_154_v5_, d_201_v41_, d_154_v5_, d_157_globalState_)
        _ingredientsd_0 = []
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_149_v0_).length(0)):
            d_202_i4_: int = guard_loop_0_
            if (True) and (((0) <= (d_202_i4_)) and ((d_202_i4_) < ((d_149_v0_).length(0)))):
                _ingredientsd_0.append((d_149_v0_, int(d_202_i4_), (d_202_i4_) + ((d_154_v5_ if (d_160_v9_)[default__.safeIndex(773, (d_160_v9_).length(0))] else (0) - (d_154_v5_)))))
        for _tupd_0 in _ingredientsd_0:
            _tupd_0[0][_tupd_0[1]] = _tupd_0[2]
        default__.m0((d_152_v2_ if False else True), len((_dafny.SeqWithoutIsStrInference([d_154_v5_])) + (d_159_v8_)), (d_201_v41_) | (d_201_v41_), d_154_v5_, d_157_globalState_)
        d_203_v42_: _dafny.Map
        d_203_v42_ = _dafny.Map({d_152_v2_: d_154_v5_})
        d_204_v43_: _dafny.Map
        d_204_v43_ = _dafny.Map({(d_151_v1_) + (_dafny.SeqWithoutIsStrInference([d_155_v6_ for d_205_i5_ in range(default__.abs(-461))])): not(((d_160_v9_)[default__.safeIndex(773, (d_160_v9_).length(0))]) in (d_203_v42_))})
        d_204_v43_ = d_204_v43_
        (d_157_globalState_).f13 = d_152_v2_
        d_206_v44_: _dafny.Map
        d_206_v44_ = _dafny.Map({d_154_v5_: (d_160_v9_)[default__.safeIndex(773, (d_160_v9_).length(0))]})
        d_207_v45_: D1
        d_207_v45_ = D1_DC2(d_206_v44_)
        d_208_v46_: _dafny.Set
        d_208_v46_ = _dafny.Set({d_154_v5_})
        d_207_v45_ = D1_DC2(default__.fm3(d_154_v5_, False, d_151_v1_, len(d_208_v46_), d_157_globalState_))
        (d_157_globalState_).f16 = d_154_v5_
        d_209_v47_: _dafny.MultiSet
        d_209_v47_ = _dafny.MultiSet([d_154_v5_])
        (d_157_globalState_).f13 = (d_209_v47_).isdisjoint(d_209_v47_)
        source6_ = D0_DC1(default__.fm4(default__.fm0(d_154_v5_, 593, d_154_v5_, d_152_v2_, d_157_globalState_), d_154_v5_, d_157_globalState_))
        if source6_.is_DC0:
            d_210___mcc_h2_ = source6_.cf0
            d_211_cf0_ = d_210___mcc_h2_
            pat_let_tv4_ = d_206_v44_
            def iife31_(_pat_let2_0):
                def iife32_(d_212_dt__update__tmp_h0_):
                    def iife33_(_pat_let3_0):
                        def iife34_(d_213_dt__update_hcf2_h0_):
                            return D1_DC2(d_213_dt__update_hcf2_h0_)
                        return iife34_(_pat_let3_0)
                    return iife33_(pat_let_tv4_)
                return iife32_(_pat_let2_0)
            source7_ = iife31_(d_207_v45_)
            if source7_.is_DC3:
                d_214___mcc_h4_ = source7_.cf3
                d_215_cf3_ = d_214___mcc_h4_
                d_149_v0_ = d_149_v0_
                d_155_v6_ = (d_155_v6_ if d_152_v2_ else d_155_v6_)
                nw18_ = _dafny.Array(False, 4)
                d_160_v9_ = nw18_
                index26_ = default__.safeIndex(575, (d_149_v0_).length(0))
                (d_149_v0_)[index26_] = d_154_v5_
                d_216_v48_: _dafny.MultiSet
                d_216_v48_ = _dafny.MultiSet([(d_160_v9_)[default__.safeIndex(773, (d_160_v9_).length(0))]])
                index27_ = default__.safeIndex(575, (d_149_v0_).length(0))
                (d_149_v0_)[index27_] = default__.safeDivisionInt(d_154_v5_, default__.safeModuloInt(default__.fm0(d_154_v5_, (d_216_v48_).cardinality, d_154_v5_, ((d_206_v44_)[len(d_151_v1_)] if (len(d_151_v1_)) in (d_206_v44_) else (d_160_v9_)[default__.safeIndex(773, (d_160_v9_).length(0))]), d_157_globalState_), (0) - (d_154_v5_)))
            elif True:
                d_217___mcc_h5_ = source7_.cf2
                d_218_cf2_ = d_217___mcc_h5_
                d_200_v40_ = (d_200_v40_) + (_dafny.SeqWithoutIsStrInference([(d_160_v9_)[default__.safeIndex(773, (d_160_v9_).length(0))]]))
                d_159_v8_ = (_dafny.SeqWithoutIsStrInference([d_154_v5_, d_154_v5_])) + (d_159_v8_)
                def iife35_():
                    coll27_ = _dafny.Set()
                    compr_27_: int
                    for compr_27_ in _dafny.IntegerRange(375, 820):
                        d_219_v49_: int = compr_27_
                        if ((375) <= (d_219_v49_)) and ((d_219_v49_) < (820)):
                            coll27_ = coll27_.union(_dafny.Set([(d_219_v49_) - (len(d_200_v40_))]))
                    return _dafny.Set(coll27_)
                (d_157_globalState_).f13 = not((d_208_v46_) != (iife35_()
                ))
                d_220_v50_: _dafny.MultiSet
                d_220_v50_ = _dafny.MultiSet([d_149_v0_])
                d_221_v51_: _dafny.Map
                d_221_v51_ = _dafny.Map({d_211_cf0_: (d_220_v50_)})
                d_222_v52_: _dafny.MultiSet
                d_222_v52_ = _dafny.MultiSet([d_149_v0_])
                d_221_v51_ = (d_221_v51_).set(d_152_v2_, d_222_v52_)
            (d_157_globalState_).f13 = d_152_v2_
            d_149_v0_ = d_149_v0_
            d_155_v6_ = _dafny.CodePoint('e')
        elif True:
            d_223___mcc_h3_ = source6_.cf1
            d_224_cf1_ = d_223___mcc_h3_
            d_225_v53_: _dafny.Array
            nw19_ = _dafny.Array(_dafny.Array(None, 0), 19)
            d_225_v53_ = nw19_
            index28_ = default__.safeIndex(229, (d_225_v53_).length(0))
            (d_225_v53_)[index28_] = d_149_v0_
            index29_ = default__.safeIndex(229, (d_225_v53_).length(0))
            (d_225_v53_)[index29_] = d_149_v0_
            default__.m0((d_160_v9_)[default__.safeIndex(773, (d_160_v9_).length(0))], -852, d_201_v41_, default__.safeModuloInt(d_154_v5_, d_154_v5_), d_157_globalState_)
            (d_157_globalState_).f0 = (d_160_v9_)[default__.safeIndex(773, (d_160_v9_).length(0))]
            d_226_v54_: _dafny.Map
            d_226_v54_ = _dafny.Map({d_155_v6_: d_152_v2_})
            d_200_v40_ = _dafny.SeqWithoutIsStrInference([(d_160_v9_)[default__.safeIndex(773, (d_160_v9_).length(0))], ((d_226_v54_)[d_155_v6_] if (d_155_v6_) in (d_226_v54_) else False)])
        d_227_i6_: int
        d_227_i6_ = 0
        with _dafny.label("1"):
            while (d_160_v9_)[default__.safeIndex(773, (d_160_v9_).length(0))]:
                with _dafny.c_label("1"):
                    if (d_227_i6_) >= (100):
                        raise _dafny.Break("1")
                    d_227_i6_ = (d_227_i6_) + (1)
                    pat_let_tv5_ = d_206_v44_
                    def iife36_(_pat_let4_0):
                        def iife37_(d_228_dt__update__tmp_h1_):
                            def iife38_(_pat_let5_0):
                                def iife39_(d_229_dt__update_hcf2_h1_):
                                    return D1_DC2(d_229_dt__update_hcf2_h1_)
                                return iife39_(_pat_let5_0)
                            return iife38_(pat_let_tv5_)
                        return iife37_(_pat_let4_0)
                    d_207_v45_ = iife36_(d_207_v45_)
                    d_207_v45_ = d_207_v45_
                    d_230_v55_: _dafny.Set
                    d_230_v55_ = _dafny.Set({d_152_v2_, (d_160_v9_)[default__.safeIndex(773, (d_160_v9_).length(0))]})
                    d_154_v5_ = len(d_230_v55_)
                    (d_157_globalState_).f16 = ((0) - (d_154_v5_)) - ((148) + ((0) - (d_154_v5_)))
                    pass
            pass
        d_231_v56_: _dafny.Map
        d_231_v56_ = _dafny.Map({d_160_v9_: d_154_v5_})
        d_232_v57_: D3
        d_232_v57_ = D3_DC6((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yswqe"))).set(default__.safeIndex(d_154_v5_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yswqe")))), d_155_v6_), d_155_v6_, d_152_v2_)
        d_233_v58_: _dafny.Array
        nw20_ = _dafny.Array(None, 21)
        nw20_[int(0)] = d_231_v56_
        nw20_[int(1)] = d_231_v56_
        nw20_[int(2)] = (d_231_v56_).set(d_160_v9_, d_154_v5_)
        nw20_[int(3)] = (d_231_v56_) | (d_231_v56_)
        nw20_[int(4)] = d_231_v56_
        nw20_[int(5)] = d_231_v56_
        nw20_[int(6)] = (d_231_v56_ if (d_160_v9_)[default__.safeIndex(773, (d_160_v9_).length(0))] else d_231_v56_)
        nw20_[int(7)] = (d_231_v56_) | ((d_231_v56_).set(d_160_v9_, (_dafny.MultiSet([(d_232_v57_).cf7])).cardinality))
        nw20_[int(8)] = (d_231_v56_) | (d_231_v56_)
        nw20_[int(9)] = d_231_v56_
        nw20_[int(10)] = d_231_v56_
        nw20_[int(11)] = d_231_v56_
        nw20_[int(12)] = d_231_v56_
        nw20_[int(13)] = d_231_v56_
        nw20_[int(14)] = (d_231_v56_).set(d_160_v9_, d_154_v5_)
        nw20_[int(15)] = d_231_v56_
        nw20_[int(16)] = _dafny.Map({d_160_v9_: d_154_v5_})
        nw20_[int(17)] = (d_231_v56_) | (d_231_v56_)
        nw20_[int(18)] = d_231_v56_
        nw20_[int(19)] = _dafny.Map({d_160_v9_: d_154_v5_})
        nw20_[int(20)] = _dafny.Map({d_160_v9_: 129})
        d_233_v58_ = nw20_
        index30_ = default__.safeIndex(421, (d_233_v58_).length(0))
        (d_233_v58_)[index30_] = d_231_v56_
        index31_ = default__.safeIndex(421, (d_233_v58_).length(0))
        index32_ = default__.safeIndex(773, (d_160_v9_).length(0))
        rhs22_ = d_231_v56_
        rhs23_ = not(d_152_v2_)
        rhs24_ = d_154_v5_
        lhs26_ = d_233_v58_
        lhs27_ = default__.safeIndex(421, (d_233_v58_).length(0))
        lhs28_ = d_160_v9_
        lhs29_ = default__.safeIndex(773, (d_160_v9_).length(0))
        lhs26_[lhs27_] = rhs22_
        lhs28_[lhs29_] = rhs23_
        d_154_v5_ = rhs24_
        d_152_v2_ = ((d_160_v9_)[default__.safeIndex(773, (d_160_v9_).length(0))]) and (d_152_v2_)
        _dafny.print(_dafny.string_of((d_149_v0_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_149_v0_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_149_v0_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_149_v0_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_149_v0_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_149_v0_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_149_v0_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_149_v0_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_149_v0_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_149_v0_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_149_v0_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_149_v0_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_149_v0_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_149_v0_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_149_v0_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_149_v0_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_149_v0_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_149_v0_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_151_v1_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_152_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_153_v3_) == (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ucjarmpr"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_154_v5_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_155_v6_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_157_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_157_globalState_).f1)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_157_globalState_).f1)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_157_globalState_).f1)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_157_globalState_).f1)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_157_globalState_).f1)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_157_globalState_).f1)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_157_globalState_).f1)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_157_globalState_).f1)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_157_globalState_).f1)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_157_globalState_).f1)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_157_globalState_).f1)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_157_globalState_).f1)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_157_globalState_).f1)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_157_globalState_).f1)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_157_globalState_).f1)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_157_globalState_).f1)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_157_globalState_).f1)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_157_globalState_).f1)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_157_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_157_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_157_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_157_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_157_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_157_globalState_.f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_157_globalState_.f8).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_157_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_157_globalState_.f10) == (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ucjarmpr"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_157_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_157_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_157_globalState_.f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_157_globalState_).f14).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_157_globalState_).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_157_globalState_.f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_157_globalState_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_157_globalState_).f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_157_globalState_.f19).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_157_globalState_.f20).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_157_globalState_.f21))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_159_v8_) == (_dafny.SeqWithoutIsStrInference([710]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v9_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v39_).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_200_v40_) == (_dafny.SeqWithoutIsStrInference([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_201_v41_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([False, False, False, False]): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ncojfqse"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_v42_) == (_dafny.Map({False: -206}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_v43_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bsiucjarmprttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt")): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v44_) == (_dafny.Map({-206: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_207_v45_).cf2) == (_dafny.Map({3: False, 581: False, 0: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_208_v46_) == (_dafny.Set({-206}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_209_v47_) == (_dafny.MultiSet([-206]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_227_i6_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_231_v56_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_232_v57_).cf6).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_v57_).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_v57_).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_233_v58_)[0])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_233_v58_)[1])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_233_v58_)[2])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_233_v58_)[3])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_233_v58_)[4])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_233_v58_)[5])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_233_v58_)[6])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_233_v58_)[7])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_233_v58_)[8])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_233_v58_)[9])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_233_v58_)[10])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_233_v58_)[11])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_233_v58_)[12])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_233_v58_)[13])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_233_v58_)[14])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_233_v58_)[15])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_233_v58_)[16])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_233_v58_)[17])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_233_v58_)[18])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_233_v58_)[19])))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_233_v58_)[20])))
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

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC3(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)

class D1_DC3(D1, NamedTuple('DC3', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC2(D1, NamedTuple('DC2', [('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D2_DC4)

class D2_DC4(D2, NamedTuple('DC4', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC6(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.CodePoint('D'), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D3_DC6)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D3_DC5)

class D3_DC6(D3, NamedTuple('DC6', [('cf6', Any), ('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC6({self.cf6.VerbatimString(True)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC6) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC5(D3, NamedTuple('DC5', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC5({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC5) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC8(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D4_DC8)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D4_DC9)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D4_DC7)

class D4_DC8(D4, NamedTuple('DC8', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC8({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC8) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC9(D4, NamedTuple('DC9', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC10(D4, NamedTuple('DC10', [('cf12', Any), ('cf13', Any), ('cf14', Any), ('cf15', Any), ('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf12 == __o.cf12 and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15 and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC7(D4, NamedTuple('DC7', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC7({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC7) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D5_DC11)

class D5_DC11(D5, NamedTuple('DC11', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC11({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC11) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC13(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D6_DC13)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D6_DC12)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D6_DC14)

class D6_DC13(D6, NamedTuple('DC13', [('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC13({_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC13) and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC12(D6, NamedTuple('DC12', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC12({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC12) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC14(D6, NamedTuple('DC14', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC14({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC14) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC16(_dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D7_DC16)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D7_DC15)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D7_DC17)

class D7_DC16(D7, NamedTuple('DC16', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC16({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC16) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC15(D7, NamedTuple('DC15', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC15({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC15) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC17(D7, NamedTuple('DC17', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC17({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC17) and self.cf24 == __o.cf24
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

class D8_DC18(D8, NamedTuple('DC18', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC18({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC18) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D9_DC19)

class D9_DC19(D9, NamedTuple('DC19', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC19({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC19) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D10_DC20)

class D10_DC20(D10, NamedTuple('DC20', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC20({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC20) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D11_DC21)

class D11_DC21(D11, NamedTuple('DC21', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC21({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC21) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC23(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D12_DC23)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D12_DC22)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D12_DC24)

class D12_DC23(D12, NamedTuple('DC23', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC23({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC23) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC22(D12, NamedTuple('DC22', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC22({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC22) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC24(D12, NamedTuple('DC24', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC24({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC24) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC26(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D13_DC26)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D13_DC27)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D13_DC25)

class D13_DC26(D13, NamedTuple('DC26', [('cf33', Any), ('cf34', Any), ('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC26({self.cf33.VerbatimString(True)}, {_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC26) and self.cf33 == __o.cf33 and self.cf34 == __o.cf34 and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC27(D13, NamedTuple('DC27', [('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC27({_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC27) and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC25(D13, NamedTuple('DC25', [('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC25({_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC25) and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC29(int(0), _dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D14_DC29)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D14_DC30)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D14_DC28)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D14_DC31)

class D14_DC29(D14, NamedTuple('DC29', [('cf38', Any), ('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC29({_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC29) and self.cf38 == __o.cf38 and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC30(D14, NamedTuple('DC30', [('cf40', Any), ('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC30({_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC30) and self.cf40 == __o.cf40 and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC28(D14, NamedTuple('DC28', [('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC28({_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC28) and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC31(D14, NamedTuple('DC31', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC31({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC31) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC33(_dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D15_DC33)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D15_DC32)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D15_DC34)

class D15_DC33(D15, NamedTuple('DC33', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC33({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC33) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC32(D15, NamedTuple('DC32', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC32({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC32) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC34(D15, NamedTuple('DC34', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC34({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC34) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC36(False, _dafny.MultiSet({}), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D16_DC36)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D16_DC35)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D16_DC37)

class D16_DC36(D16, NamedTuple('DC36', [('cf47', Any), ('cf48', Any), ('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC36({_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)}, {_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC36) and self.cf47 == __o.cf47 and self.cf48 == __o.cf48 and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC35(D16, NamedTuple('DC35', [('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC35({_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC35) and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC37(D16, NamedTuple('DC37', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC37({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC37) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC39(int(0), _dafny.Map({}), False, _dafny.Map({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D17_DC39)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D17_DC40)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D17_DC41)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D17_DC38)

class D17_DC39(D17, NamedTuple('DC39', [('cf52', Any), ('cf53', Any), ('cf54', Any), ('cf55', Any), ('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC39({_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)}, {_dafny.string_of(self.cf55)}, {_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC39) and self.cf52 == __o.cf52 and self.cf53 == __o.cf53 and self.cf54 == __o.cf54 and self.cf55 == __o.cf55 and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC40(D17, NamedTuple('DC40', [])):
    def __dafnystr__(self) -> str:
        return f'D17.DC40'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC40)
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC41(D17, NamedTuple('DC41', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC41({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC41) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC38(D17, NamedTuple('DC38', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC38({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC38) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC43()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D18_DC43)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D18_DC44)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D18_DC42)

class D18_DC43(D18, NamedTuple('DC43', [])):
    def __dafnystr__(self) -> str:
        return f'D18.DC43'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC43)
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC44(D18, NamedTuple('DC44', [])):
    def __dafnystr__(self) -> str:
        return f'D18.DC44'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC44)
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC42(D18, NamedTuple('DC42', [('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC42({_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC42) and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC46(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D19_DC46)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D19_DC47)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D19_DC48)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D19_DC45)

class D19_DC46(D19, NamedTuple('DC46', [('cf60', Any), ('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC46({_dafny.string_of(self.cf60)}, {_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC46) and self.cf60 == __o.cf60 and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC47(D19, NamedTuple('DC47', [('cf62', Any), ('cf63', Any), ('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC47({_dafny.string_of(self.cf62)}, {_dafny.string_of(self.cf63)}, {_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC47) and self.cf62 == __o.cf62 and self.cf63 == __o.cf63 and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC48(D19, NamedTuple('DC48', [('cf65', Any), ('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC48({_dafny.string_of(self.cf65)}, {_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC48) and self.cf65 == __o.cf65 and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC45(D19, NamedTuple('DC45', [('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC45({_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC45) and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC50(int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D20_DC50)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D20_DC51)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D20_DC49)

class D20_DC50(D20, NamedTuple('DC50', [('cf68', Any), ('cf69', Any), ('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC50({_dafny.string_of(self.cf68)}, {_dafny.string_of(self.cf69)}, {_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC50) and self.cf68 == __o.cf68 and self.cf69 == __o.cf69 and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC51(D20, NamedTuple('DC51', [])):
    def __dafnystr__(self) -> str:
        return f'D20.DC51'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC51)
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC49(D20, NamedTuple('DC49', [('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC49({_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC49) and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC53(_dafny.Array(None, 0), False, _dafny.MultiSet({}), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D21_DC53)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D21_DC54)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D21_DC52)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D21_DC55)

class D21_DC53(D21, NamedTuple('DC53', [('cf72', Any), ('cf73', Any), ('cf74', Any), ('cf75', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC53({_dafny.string_of(self.cf72)}, {_dafny.string_of(self.cf73)}, {_dafny.string_of(self.cf74)}, {_dafny.string_of(self.cf75)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC53) and self.cf72 == __o.cf72 and self.cf73 == __o.cf73 and self.cf74 == __o.cf74 and self.cf75 == __o.cf75
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC54(D21, NamedTuple('DC54', [('cf76', Any), ('cf77', Any), ('cf78', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC54({_dafny.string_of(self.cf76)}, {_dafny.string_of(self.cf77)}, {_dafny.string_of(self.cf78)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC54) and self.cf76 == __o.cf76 and self.cf77 == __o.cf77 and self.cf78 == __o.cf78
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC52(D21, NamedTuple('DC52', [('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC52({_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC52) and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC55(D21, NamedTuple('DC55', [('cf79', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC55({_dafny.string_of(self.cf79)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC55) and self.cf79 == __o.cf79
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC57(int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D22_DC57)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D22_DC58)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D22_DC56)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D22_DC59)

class D22_DC57(D22, NamedTuple('DC57', [('cf81', Any), ('cf82', Any), ('cf83', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC57({_dafny.string_of(self.cf81)}, {_dafny.string_of(self.cf82)}, {_dafny.string_of(self.cf83)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC57) and self.cf81 == __o.cf81 and self.cf82 == __o.cf82 and self.cf83 == __o.cf83
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC58(D22, NamedTuple('DC58', [])):
    def __dafnystr__(self) -> str:
        return f'D22.DC58'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC58)
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC56(D22, NamedTuple('DC56', [('cf80', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC56({_dafny.string_of(self.cf80)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC56) and self.cf80 == __o.cf80
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC59(D22, NamedTuple('DC59', [('cf84', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC59({_dafny.string_of(self.cf84)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC59) and self.cf84 == __o.cf84
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: D23_DC61(_dafny.Seq({}), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D23_DC61)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D23_DC60)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D23_DC62)

class D23_DC61(D23, NamedTuple('DC61', [('cf86', Any), ('cf87', Any), ('cf88', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC61({_dafny.string_of(self.cf86)}, {_dafny.string_of(self.cf87)}, {_dafny.string_of(self.cf88)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC61) and self.cf86 == __o.cf86 and self.cf87 == __o.cf87 and self.cf88 == __o.cf88
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC60(D23, NamedTuple('DC60', [('cf85', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC60({_dafny.string_of(self.cf85)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC60) and self.cf85 == __o.cf85
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC62(D23, NamedTuple('DC62', [('cf89', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC62({_dafny.string_of(self.cf89)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC62) and self.cf89 == __o.cf89
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: D24_DC64(int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC64(self) -> bool:
        return isinstance(self, D24_DC64)
    @property
    def is_DC65(self) -> bool:
        return isinstance(self, D24_DC65)
    @property
    def is_DC63(self) -> bool:
        return isinstance(self, D24_DC63)
    @property
    def is_DC66(self) -> bool:
        return isinstance(self, D24_DC66)

class D24_DC64(D24, NamedTuple('DC64', [('cf91', Any), ('cf92', Any), ('cf93', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC64({_dafny.string_of(self.cf91)}, {_dafny.string_of(self.cf92)}, {_dafny.string_of(self.cf93)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC64) and self.cf91 == __o.cf91 and self.cf92 == __o.cf92 and self.cf93 == __o.cf93
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC65(D24, NamedTuple('DC65', [('cf94', Any), ('cf95', Any), ('cf96', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC65({_dafny.string_of(self.cf94)}, {_dafny.string_of(self.cf95)}, {_dafny.string_of(self.cf96)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC65) and self.cf94 == __o.cf94 and self.cf95 == __o.cf95 and self.cf96 == __o.cf96
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC63(D24, NamedTuple('DC63', [('cf90', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC63({_dafny.string_of(self.cf90)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC63) and self.cf90 == __o.cf90
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC66(D24, NamedTuple('DC66', [('cf97', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC66({_dafny.string_of(self.cf97)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC66) and self.cf97 == __o.cf97
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: D25_DC68()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC68(self) -> bool:
        return isinstance(self, D25_DC68)
    @property
    def is_DC69(self) -> bool:
        return isinstance(self, D25_DC69)
    @property
    def is_DC67(self) -> bool:
        return isinstance(self, D25_DC67)
    @property
    def is_DC70(self) -> bool:
        return isinstance(self, D25_DC70)

class D25_DC68(D25, NamedTuple('DC68', [])):
    def __dafnystr__(self) -> str:
        return f'D25.DC68'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC68)
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC69(D25, NamedTuple('DC69', [])):
    def __dafnystr__(self) -> str:
        return f'D25.DC69'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC69)
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC67(D25, NamedTuple('DC67', [('cf98', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC67({_dafny.string_of(self.cf98)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC67) and self.cf98 == __o.cf98
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC70(D25, NamedTuple('DC70', [('cf99', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC70({_dafny.string_of(self.cf99)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC70) and self.cf99 == __o.cf99
    def __hash__(self) -> int:
        return super().__hash__()


class D26:
    @classmethod
    def default(cls, ):
        return lambda: D26_DC72(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC72(self) -> bool:
        return isinstance(self, D26_DC72)
    @property
    def is_DC73(self) -> bool:
        return isinstance(self, D26_DC73)
    @property
    def is_DC71(self) -> bool:
        return isinstance(self, D26_DC71)
    @property
    def is_DC74(self) -> bool:
        return isinstance(self, D26_DC74)

class D26_DC72(D26, NamedTuple('DC72', [('cf101', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC72({_dafny.string_of(self.cf101)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC72) and self.cf101 == __o.cf101
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC73(D26, NamedTuple('DC73', [('cf102', Any), ('cf103', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC73({_dafny.string_of(self.cf102)}, {_dafny.string_of(self.cf103)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC73) and self.cf102 == __o.cf102 and self.cf103 == __o.cf103
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC71(D26, NamedTuple('DC71', [('cf100', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC71({_dafny.string_of(self.cf100)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC71) and self.cf100 == __o.cf100
    def __hash__(self) -> int:
        return super().__hash__()

class D26_DC74(D26, NamedTuple('DC74', [('cf104', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC74({_dafny.string_of(self.cf104)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC74) and self.cf104 == __o.cf104
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m1(self, p0, globalState):
        pass


class T1:
    pass
    def m2(self, p0, p1, p2, p3, globalState):
        pass

    def m3(self, p0, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f0: bool = False
        self.f2: bool = False
        self.f7: int = int(0)
        self.f8: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f10: _dafny.Set = _dafny.Set({})
        self.f13: bool = False
        self.f16: int = int(0)
        self.f19: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f20: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f21: bool = False
        self.f22: _dafny.Array = _dafny.Array(None, 0)
        self._f1: _dafny.Array = _dafny.Array(None, 0)
        self._f3: int = int(0)
        self._f4: bool = False
        self._f5: bool = False
        self._f6: bool = False
        self._f9: int = int(0)
        self._f11: int = int(0)
        self._f12: int = int(0)
        self._f14: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f15: int = int(0)
        self._f17: bool = False
        self._f18: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22):
        (self).f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self).f7 = f7
        (self).f8 = f8
        (self)._f9 = f9
        (self).f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self).f13 = f13
        (self)._f14 = f14
        (self)._f15 = f15
        (self).f16 = f16
        (self)._f17 = f17
        (self)._f18 = f18
        (self).f19 = f19
        (self).f20 = f20
        (self).f21 = f21
        (self).f22 = f22

    @property
    def f1(self):
        return self._f1
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
    def f9(self):
        return self._f9
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
    def f15(self):
        return self._f15
    @property
    def f17(self):
        return self._f17
    @property
    def f18(self):
        return self._f18

class C0:
    def  __init__(self):
        self.f27: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f27):
        (self).f27 = f27


class C1(T1, T0):
    def  __init__(self):
        self.f29: int = int(0)
        self._f30: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self, f29, f30):
        (self).f29 = f29
        (self)._f30 = f30

    def fm7(self, p0, p1, p2, globalState):
        return True

    def fm5(self, p0, p1, p2, p3, globalState):
        return _dafny.Map({(False) == (not(True)): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ctdvjyvbi"))})

    def fm6(self, p0, globalState):
        return (self.f29) == (self.f29)

    def fm29(self, globalState):
        return len(_dafny.Map({_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([False, False, True])).cardinality]): (_dafny.Map({True: False})) | (_dafny.Map({True: True}))}))

    def m2(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: _dafny.Array = _dafny.Array(None, 0)
        if p1:
            if p1:
                (globalState).f2 = p1
                d_234_v0_: C0
                nw21_ = C0()
                nw21_.ctor__(not(True))
                d_234_v0_ = nw21_
                d_235_v1_: C0
                nw22_ = C0()
                nw22_.ctor__(p1)
                d_235_v1_ = nw22_
                d_236_v2_: str
                d_236_v2_ = _dafny.CodePoint('l')
                d_237_v3_: _dafny.Seq
                d_237_v3_ = _dafny.SeqWithoutIsStrInference([d_236_v2_])
                d_238_v4_: _dafny.MultiSet
                d_238_v4_ = _dafny.MultiSet([d_237_v3_])
                d_239_v5_: _dafny.Seq
                d_239_v5_ = _dafny.SeqWithoutIsStrInference([default__.safeModuloInt(p0, (d_238_v4_).cardinality)])
                d_240_v6_: _dafny.Array
                nw23_ = _dafny.Array(_dafny.Array(None, 0), 14)
                d_240_v6_ = nw23_
                d_241_v7_: _dafny.Array
                nw24_ = _dafny.Array(False, 3)
                d_241_v7_ = nw24_
                index33_ = default__.safeIndex(119, (d_240_v6_).length(0))
                (d_240_v6_)[index33_] = d_241_v7_
                d_242_v8_: _dafny.Map
                d_242_v8_ = _dafny.Map({p1: d_235_v1_.f27})
                d_243_v9_: _dafny.Array
                nw25_ = _dafny.Array(None, 1)
                nw25_[int(0)] = (default__.fm0(p3, 438, len(d_237_v3_), p1, globalState)) - (len(d_242_v8_))
                d_243_v9_ = nw25_
                d_244_v10_: _dafny.Map
                d_244_v10_ = _dafny.Map({p0: d_234_v0_.f27})
                index34_ = default__.safeIndex(119, (d_240_v6_).length(0))
                rhs25_ = _dafny.SeqWithoutIsStrInference([p3, len(d_244_v10_), p3])
                rhs26_ = d_241_v7_
                rhs27_ = p3
                rhs28_ = d_243_v9_
                lhs30_ = d_240_v6_
                lhs31_ = default__.safeIndex(119, (d_240_v6_).length(0))
                lhs32_ = globalState
                d_239_v5_ = rhs25_
                lhs30_[lhs31_] = rhs26_
                lhs32_.f7 = rhs27_
                d_243_v9_ = rhs28_
                (globalState).f21 = d_234_v0_.f27
            elif True:
                d_245_v11_: _dafny.Array
                def lambda11_(d_246_i0_):
                    return (d_246_i0_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ks"))))

                init6_ = lambda11_
                nw26_ = _dafny.Array(None, 29)
                for i0_6_ in range(nw26_.length(0)):
                    nw26_[i0_6_] = init6_(i0_6_)
                d_245_v11_ = nw26_
                nw27_ = _dafny.Array(int(0), 27)
                d_245_v11_ = nw27_
                (globalState).f0 = p1
                d_247_v12_: _dafny.Map
                d_247_v12_ = _dafny.Map({p1: p0})
                (globalState).f7 = default__.safeDivisionInt(default__.safeModuloInt(p0, len(d_247_v12_)), len(_dafny.Map({not(p1): True})))
                d_248_v13_: _dafny.Map
                d_248_v13_ = _dafny.Map({p1: p1})
                d_249_v14_: _dafny.MultiSet
                d_249_v14_ = _dafny.MultiSet([self.f29, (default__.fm0(p0, p3, 758, p1, globalState)) * (len(d_248_v13_)), p0])
                d_249_v14_ = d_249_v14_
                d_250_v15_: D3
                d_250_v15_ = D3_DC5(_dafny.CodePoint('m'))
                d_251_v16_: _dafny.MultiSet
                d_251_v16_ = _dafny.MultiSet([d_250_v15_])
                d_252_v17_: _dafny.Set
                d_252_v17_ = _dafny.Set({p3})
                d_253_v18_: _dafny.Seq
                d_253_v18_ = _dafny.SeqWithoutIsStrInference([p1, not((default__.fm30(p1, self.f29, (0) - (self.f29), d_251_v16_, globalState)) == (d_252_v17_)), p1, p1])
                d_253_v18_ = d_253_v18_
            d_254_v19_: _dafny.Map
            d_254_v19_ = _dafny.Map({p1: p3})
            rhs29_ = len((d_254_v19_ if p1 else d_254_v19_))
            rhs30_ = (0) - (len(default__.fm31(p0, not (p1) or (p1), globalState)))
            lhs33_ = globalState
            lhs33_.f16 = rhs29_
            r1 = rhs30_
            if p1:
                (globalState).f2 = p1
                d_255_v20_: _dafny.Seq
                d_255_v20_ = _dafny.SeqWithoutIsStrInference([p1])
                d_256_v21_: _dafny.Map
                d_256_v21_ = _dafny.Map({p0: d_255_v20_})
                d_257_v22_: _dafny.Map
                d_257_v22_ = _dafny.Map({((d_256_v21_)[142] if (142) in (d_256_v21_) else d_255_v20_): default__.fm1(self.f29, globalState)})
                default__.m0(p1, p0, d_257_v22_, (self).fm29(globalState), globalState)
                d_258_v23_: _dafny.Map
                d_258_v23_ = _dafny.Map({(d_255_v20_) + (_dafny.SeqWithoutIsStrInference([p1, p1, p1, p1, p1])): False})
                d_259_v24_: _dafny.Set
                d_259_v24_ = _dafny.Set({default__.fm0(p3, p3, p3, p1, globalState), len((d_255_v20_).set(default__.safeIndex(p3, len(d_255_v20_)), p1)), self.f29, self.f29})
                d_260_v25_: D14
                d_260_v25_ = D14_DC28(d_259_v24_)
                (globalState).f2 = ((d_258_v23_)[d_255_v20_] if (d_255_v20_) in (d_258_v23_) else ((d_260_v25_).cf37).ispropersubset(d_259_v24_))
                d_261_v26_: _dafny.Seq
                d_261_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ifljbrg"))
                d_262_v27_: _dafny.Seq
                d_262_v27_ = _dafny.SeqWithoutIsStrInference([self.f29, len(d_261_v26_)])
                (globalState).f7 = default__.fm0(p0, (d_262_v27_)[default__.safeIndex(p0, len(d_262_v27_))], 392, p1, globalState)
                index35_ = default__.safeIndex(828, (p2).length(0))
                (p2)[index35_] = p0
                index36_ = default__.safeIndex(828, (p2).length(0))
                (p2)[index36_] = p3
            elif True:
                d_263_v28_: _dafny.Seq
                d_263_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gtrohbu"))
                d_264_v29_: _dafny.Map
                d_264_v29_ = _dafny.Map({p1: d_263_v28_})
                d_265_v30_: str
                d_265_v30_ = _dafny.CodePoint('w')
                (globalState).f8 = ((((d_264_v29_)[p1] if (p1) in (d_264_v29_) else d_263_v28_)).set(default__.safeIndex(self.f29, len(((d_264_v29_)[p1] if (p1) in (d_264_v29_) else d_263_v28_))), d_265_v30_)) + (d_263_v28_)
                d_266_v31_: _dafny.Array
                nw28_ = _dafny.Array(_dafny.Set({}), 11)
                d_266_v31_ = nw28_
                d_267_v32_: _dafny.Map
                d_267_v32_ = _dafny.Map({p3: _dafny.SeqWithoutIsStrInference([self.f29, len(d_263_v28_)])})
                d_268_v33_: _dafny.Seq
                d_268_v33_ = _dafny.SeqWithoutIsStrInference([p1])
                d_269_v34_: D12
                d_269_v34_ = D12_DC23(528)
                d_270_v35_: _dafny.Map
                d_270_v35_ = _dafny.Map({p1: p1})
                d_271_v36_: _dafny.Seq
                d_271_v36_ = _dafny.SeqWithoutIsStrInference([p0, (d_269_v34_).cf30, len(d_270_v35_)])
                d_272_v37_: _dafny.Set
                d_272_v37_ = _dafny.Set({p0, p0, len(((d_267_v32_)[len(d_268_v33_)] if (len(d_268_v33_)) in (d_267_v32_) else _dafny.SeqWithoutIsStrInference([(d_271_v36_)[default__.safeIndex(p0, len(d_271_v36_))], p0])))})
                index37_ = default__.safeIndex(53, (d_266_v31_).length(0))
                (d_266_v31_)[index37_] = d_272_v37_
                index38_ = default__.safeIndex(53, (d_266_v31_).length(0))
                (d_266_v31_)[index38_] = d_272_v37_
                d_273_v38_: _dafny.Array
                nw29_ = _dafny.Array(D14.default()(), 3)
                d_273_v38_ = nw29_
                d_274_v39_: D14
                d_274_v39_ = D14_DC28((d_266_v31_)[default__.safeIndex(53, (d_266_v31_).length(0))])
                index39_ = default__.safeIndex(136, (d_273_v38_).length(0))
                (d_273_v38_)[index39_] = d_274_v39_
                d_275_v40_: _dafny.Map
                d_275_v40_ = _dafny.Map({(_dafny.Set({d_270_v35_, d_270_v35_, d_270_v35_, _dafny.Map({(self).fm6(p0, globalState): True}), d_270_v35_})).isdisjoint(_dafny.Set({(d_270_v35_).set(p1, p1), d_270_v35_})): p2})
                d_276_v41_: _dafny.MultiSet
                d_276_v41_ = _dafny.MultiSet([p1, (d_268_v33_)[default__.safeIndex(522, len(d_268_v33_))]])
                index40_ = default__.safeIndex(136, (d_273_v38_).length(0))
                rhs31_ = p1
                rhs32_ = default__.fm32((not(p1)) not in (_dafny.SeqWithoutIsStrInference([p1, False])), (d_276_v41_).cardinality, p3, globalState)
                rhs33_ = (d_275_v40_) | (d_275_v40_)
                lhs34_ = globalState
                lhs35_ = d_273_v38_
                lhs36_ = default__.safeIndex(136, (d_273_v38_).length(0))
                lhs34_.f2 = rhs31_
                lhs35_[lhs36_] = rhs32_
                d_275_v40_ = rhs33_
                (globalState).f2 = p1
                d_277_v42_: _dafny.Set
                d_277_v42_ = _dafny.Set({p1, p1, (self).fm7(p1, p3, p1, globalState)})
                d_278_v43_: _dafny.MultiSet
                d_278_v43_ = _dafny.MultiSet([d_277_v42_])
                d_279_v44_: _dafny.Set
                d_279_v44_ = d_277_v42_
                index41_ = default__.safeIndex(618, (p2).length(0))
                (p2)[index41_] = ((d_278_v43_)[d_279_v44_] if (d_279_v44_) in (d_278_v43_) else p0)
                d_280_v45_: _dafny.MultiSet
                d_280_v45_ = _dafny.MultiSet([(p3) + (self.f29), self.f29, (self.f29) + (p0)])
                index42_ = default__.safeIndex(618, (p2).length(0))
                (p2)[index42_] = (0) - (((d_280_v45_)[p3] if (p3) in (d_280_v45_) else default__.safeDivisionInt(((d_254_v19_)[p1] if (p1) in (d_254_v19_) else p3), self.f29)))
            (globalState).f13 = (len(_dafny.SeqWithoutIsStrInference([p0, len(d_254_v19_), self.f29, p3, self.f29]))) <= (p0)
            d_281_v46_: str
            d_281_v46_ = _dafny.CodePoint('t')
            d_282_v47_: _dafny.Seq
            d_282_v47_ = _dafny.SeqWithoutIsStrInference([d_281_v46_, d_281_v46_])
            r1 = default__.safeModuloInt(p0, len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_283_i1_ in range(default__.abs(141))])) + (d_282_v47_)))
        elif True:
            d_284_v48_: _dafny.Seq
            d_284_v48_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
            d_285_v49_: str
            d_285_v49_ = _dafny.CodePoint('i')
            d_286_v50_: D3
            d_286_v50_ = D3_DC6(d_284_v48_, d_285_v49_, p1)
            (globalState).f7 = len(_dafny.SeqWithoutIsStrInference([d_286_v50_]))
            index43_ = default__.safeIndex(348, (p2).length(0))
            (p2)[index43_] = (0) - ((0) - (p3))
            d_287_v51_: D4
            d_287_v51_ = D4_DC9(p3)
            d_288_v52_: _dafny.Seq
            d_288_v52_ = _dafny.SeqWithoutIsStrInference([p1, p1, p1])
            index44_ = default__.safeIndex(348, (p2).length(0))
            (p2)[index44_] = (p0) * (((d_287_v51_).cf11) + (len(d_288_v52_)))
            d_289_v53_: _dafny.Map
            d_289_v53_ = _dafny.Map({p1: True})
            (globalState).f21 = default__.fm2((d_289_v53_) | (d_289_v53_), globalState)
            (globalState).f7 = (0) - ((self.f29) - ((0) - ((p0) * (549))))
            d_290_v54_: _dafny.Array
            nw30_ = _dafny.Array(False, 21)
            d_290_v54_ = nw30_
            index45_ = default__.safeIndex(473, (d_290_v54_).length(0))
            (d_290_v54_)[index45_] = not(p1)
            d_291_v55_: _dafny.Seq
            d_291_v55_ = _dafny.SeqWithoutIsStrInference([(0) - ((p2)[default__.safeIndex(348, (p2).length(0))])])
            d_292_v56_: _dafny.Map
            d_292_v56_ = _dafny.Map({True: p3})
            d_293_v57_: _dafny.Map
            d_293_v57_ = _dafny.Map({d_284_v48_: _dafny.MultiSet([(d_291_v55_)[default__.safeIndex(len(d_292_v56_), len(d_291_v55_))], (p2)[default__.safeIndex(348, (p2).length(0))]])})
            d_294_v58_: _dafny.MultiSet
            d_294_v58_ = _dafny.MultiSet([p0, len(d_284_v48_)])
            index46_ = default__.safeIndex(473, (d_290_v54_).length(0))
            rhs34_ = (((d_293_v57_)[d_284_v48_] if (d_284_v48_) in (d_293_v57_) else d_294_v58_)).cardinality
            rhs35_ = p1
            lhs37_ = self
            lhs38_ = d_290_v54_
            lhs39_ = default__.safeIndex(473, (d_290_v54_).length(0))
            lhs37_.f29 = rhs34_
            lhs38_[lhs39_] = rhs35_
        d_295_v59_: _dafny.Map
        d_295_v59_ = _dafny.Map({p2: p0})
        d_295_v59_ = (d_295_v59_).set(p2, p0)
        d_296_v60_: _dafny.Seq
        d_296_v60_ = _dafny.SeqWithoutIsStrInference([self.f29])
        d_297_v61_: C0
        nw31_ = C0()
        nw31_.ctor__(p1)
        d_297_v61_ = nw31_
        d_298_v62_: _dafny.Seq
        d_298_v62_ = _dafny.SeqWithoutIsStrInference([p1, p1, p1, p1])
        d_299_v63_: _dafny.Map
        d_299_v63_ = _dafny.Map({(d_298_v62_)[default__.safeIndex(len(d_298_v62_), len(d_298_v62_))]: d_297_v61_.f27})
        d_300_v64_: _dafny.Map
        d_300_v64_ = _dafny.Map({d_297_v61_: d_299_v63_})
        d_301_v65_: _dafny.Seq
        d_301_v65_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fgkj"))
        d_302_v66_: _dafny.Map
        d_302_v66_ = _dafny.Map({d_298_v62_: d_301_v65_})
        default__.m0(p1, default__.safeDivisionInt(len(d_296_v60_), default__.fm0(-680, len(((d_300_v64_)[d_297_v61_] if (d_297_v61_) in (d_300_v64_) else d_299_v63_)), p3, d_297_v61_.f27, globalState)), d_302_v66_, len(d_296_v60_), globalState)
        d_303_v68_: str
        d_303_v68_ = _dafny.CodePoint('w')
        d_304_v69_: _dafny.Map
        d_304_v69_ = _dafny.Map({p3: d_303_v68_})
        d_305_v70_: _dafny.Map
        d_305_v70_ = _dafny.Map({self.f29: len(d_304_v69_)})
        d_306_v72_: _dafny.Map
        def iife40_():
            coll28_ = _dafny.Set()
            compr_28_: _dafny.Map
            for compr_28_ in (_dafny.SeqWithoutIsStrInference([d_305_v70_, d_305_v70_])).Elements:
                d_307_v71_: _dafny.Map = compr_28_
                if (d_307_v71_) in (_dafny.SeqWithoutIsStrInference([d_305_v70_, d_305_v70_])):
                    coll28_ = coll28_.union(_dafny.Set([d_307_v71_]))
            return _dafny.Set(coll28_)
        d_306_v72_ = _dafny.Map({p0: len(iife40_()
        )})
        def iife41_():
            coll29_ = _dafny.Map()
            compr_29_: int
            for compr_29_ in (_dafny.Map({((d_306_v72_)[p3] if (p3) in (d_306_v72_) else p0): p0})).keys.Elements:
                d_308_v67_: int = compr_29_
                if (d_308_v67_) in (_dafny.Map({((d_306_v72_)[p3] if (p3) in (d_306_v72_) else p0): p0})):
                    coll29_[default__.safeModuloInt(d_308_v67_, (d_296_v60_)[default__.safeIndex(p0, len(d_296_v60_))])] = p0
            return _dafny.Map(coll29_)
        if (len((iife41_()
        ).set(self.f29, p0))) <= (p3):
            (self).f29 = p3
            d_309_v73_: _dafny.Array
            nw32_ = _dafny.Array(False, 1)
            d_309_v73_ = nw32_
            d_310_v74_: _dafny.MultiSet
            d_310_v74_ = _dafny.MultiSet([p3, p0])
            index47_ = default__.safeIndex(758, (d_309_v73_).length(0))
            (d_309_v73_)[index47_] = (d_310_v74_).ispropersubset(default__.fm33(globalState))
            index48_ = default__.safeIndex(758, (d_309_v73_).length(0))
            (d_309_v73_)[index48_] = ((self).fm29(globalState)) >= ((((d_305_v70_)[p0] if (p0) in (d_305_v70_) else p3)) + (self.f29))
            (globalState).f13 = p1
            d_311_v75_: _dafny.MultiSet
            d_311_v75_ = _dafny.MultiSet([d_297_v61_.f27, (d_309_v73_)[default__.safeIndex(758, (d_309_v73_).length(0))]])
            if ((d_311_v75_).set(p1, default__.abs((0) - (p0)))).issubset(d_311_v75_):
                r1 = self.f29
                d_312_v76_: _dafny.Array
                nw33_ = _dafny.Array(_dafny.CodePoint('D'), 4)
                d_312_v76_ = nw33_
                index49_ = default__.safeIndex(844, (d_312_v76_).length(0))
                (d_312_v76_)[index49_] = d_303_v68_
                index50_ = default__.safeIndex(844, (d_312_v76_).length(0))
                (d_312_v76_)[index50_] = default__.fm34(globalState)
                d_313_v77_: _dafny.MultiSet
                d_313_v77_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True, p1]), d_298_v62_])
                (globalState).f0 = not((_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([d_298_v62_, d_298_v62_, _dafny.SeqWithoutIsStrInference([p1, p1, False]), d_298_v62_])).set(default__.safeIndex(len(d_296_v60_), len(_dafny.SeqWithoutIsStrInference([d_298_v62_, d_298_v62_, _dafny.SeqWithoutIsStrInference([p1, p1, False]), d_298_v62_]))), d_298_v62_))).ispropersubset(d_313_v77_))
                default__.m0((d_303_v68_) not in (d_301_v65_), (p3) + (self.f29), d_302_v66_, (d_311_v75_).cardinality, globalState)
                (globalState).f21 = not((self.f29) != ((self).fm29(globalState)))
            elif True:
                nw34_ = C0()
                nw34_.ctor__(p1)
                d_297_v61_ = nw34_
                d_314_v78_: _dafny.Array
                nw35_ = _dafny.Array(_dafny.Seq({}), 29)
                d_314_v78_ = nw35_
                d_315_v79_: D15
                d_315_v79_ = D15_DC32(d_314_v78_)
                d_314_v78_ = (d_315_v79_).cf43
                d_316_v80_: D14
                d_316_v80_ = D14_DC28(_dafny.Set({self.f29}))
                d_317_v81_: D7
                d_317_v81_ = D7_DC15(d_296_v60_)
                d_318_v82_: D7
                d_318_v82_ = D7_DC17(d_317_v81_)
                d_319_v83_: D7
                d_319_v83_ = D7_DC17(d_318_v82_)
                d_320_v84_: D7
                d_320_v84_ = D7_DC17(d_318_v82_)
                d_321_v85_: D7
                d_321_v85_ = D7_DC17(d_318_v82_)
                d_322_v86_: D7
                d_322_v86_ = D7_DC17(d_319_v83_)
                d_323_v87_: _dafny.Map
                d_323_v87_ = _dafny.Map({d_316_v80_: d_322_v86_})
                d_324_v88_: _dafny.Set
                d_324_v88_ = _dafny.Set({(_dafny.MultiSet(d_296_v60_)).cardinality})
                d_323_v87_ = (d_323_v87_).set(D14_DC28(d_324_v88_), D7_DC17(d_320_v84_))
                d_325_v89_: _dafny.Array
                nw36_ = _dafny.Array(None, 21)
                nw36_[int(0)] = d_301_v65_
                nw36_[int(1)] = d_301_v65_
                nw36_[int(2)] = d_301_v65_
                nw36_[int(3)] = (d_301_v65_) + (d_301_v65_)
                nw36_[int(4)] = _dafny.SeqWithoutIsStrInference([d_303_v68_, d_303_v68_, d_303_v68_])
                nw36_[int(5)] = d_301_v65_
                nw36_[int(6)] = _dafny.SeqWithoutIsStrInference([d_303_v68_ for d_326_i2_ in range(default__.abs(951))])
                nw36_[int(7)] = d_301_v65_
                nw36_[int(8)] = _dafny.SeqWithoutIsStrInference([d_303_v68_, d_303_v68_, d_303_v68_, d_303_v68_, d_303_v68_])
                nw36_[int(9)] = d_301_v65_
                nw36_[int(10)] = _dafny.SeqWithoutIsStrInference([d_303_v68_ for d_327_i3_ in range(default__.abs(631))])
                nw36_[int(11)] = d_301_v65_
                nw36_[int(12)] = d_301_v65_
                nw36_[int(13)] = d_301_v65_
                nw36_[int(14)] = d_301_v65_
                nw36_[int(15)] = _dafny.SeqWithoutIsStrInference([d_303_v68_ for d_328_i4_ in range(default__.abs(-368))])
                nw36_[int(16)] = d_301_v65_
                nw36_[int(17)] = d_301_v65_
                nw36_[int(18)] = d_301_v65_
                nw36_[int(19)] = _dafny.SeqWithoutIsStrInference([d_303_v68_])
                nw36_[int(20)] = _dafny.SeqWithoutIsStrInference([d_303_v68_ for d_329_i5_ in range(default__.abs(-300))])
                d_325_v89_ = nw36_
                index51_ = default__.safeIndex(257, (d_325_v89_).length(0))
                (d_325_v89_)[index51_] = _dafny.SeqWithoutIsStrInference([d_303_v68_ for d_330_i6_ in range(default__.abs(140))])
                d_331_v90_: D3
                d_331_v90_ = D3_DC5(d_303_v68_)
                index52_ = default__.safeIndex(257, (d_325_v89_).length(0))
                (d_325_v89_)[index52_] = (d_301_v65_ if d_297_v61_.f27 else _dafny.SeqWithoutIsStrInference([(d_331_v90_).cf5]))
                d_332_v91_: C0
                nw37_ = C0()
                nw37_.ctor__(d_297_v61_.f27)
                d_332_v91_ = nw37_
            d_333_v92_: _dafny.Seq
            d_333_v92_ = _dafny.SeqWithoutIsStrInference([(d_309_v73_ if d_297_v61_.f27 else d_309_v73_), d_309_v73_])
            d_309_v73_ = (d_333_v92_)[default__.safeIndex(p0, len(d_333_v92_))]
        elif True:
            d_303_v68_ = d_303_v68_
            d_334_v93_: C0
            nw38_ = C0()
            nw38_.ctor__(d_297_v61_.f27)
            d_334_v93_ = nw38_
            d_335_v94_: _dafny.MultiSet
            d_335_v94_ = _dafny.MultiSet([p3, p3])
            (globalState).f2 = (default__.fm0(268, p3, (d_335_v94_).cardinality, d_297_v61_.f27, globalState)) < ((36) + (self.f29))
            d_336_v95_: _dafny.Map
            d_336_v95_ = _dafny.Map({d_297_v61_.f27: p0})
            d_337_v96_: _dafny.Map
            d_337_v96_ = _dafny.Map({(d_336_v95_ if d_334_v93_.f27 else _dafny.Map({not(p1): p3})): p3})
            d_338_v97_: _dafny.Map
            d_338_v97_ = _dafny.Map({p0: D14_DC30(d_296_v60_, d_298_v62_)})
            d_337_v96_ = (d_337_v96_).set((default__.fm35(p1, p3, globalState)) | (d_336_v95_), (d_296_v60_)[default__.safeIndex(len(d_338_v97_), len(d_296_v60_))])
            (globalState).f7 = p3
        d_339_v98_: _dafny.Map
        d_339_v98_ = _dafny.Map({p3: d_297_v61_.f27})
        d_340_v99_: _dafny.Map
        d_340_v99_ = _dafny.Map({self.f29: (self).fm7(not(((d_339_v98_)[self.f29] if (self.f29) in (d_339_v98_) else d_297_v61_.f27)), p0, d_297_v61_.f27, globalState)})
        (globalState).f0 = (len(d_340_v99_)) == (p3)
        d_341_v100_: _dafny.Map
        d_341_v100_ = _dafny.Map({d_297_v61_.f27: p3})
        d_342_v101_: _dafny.MultiSet
        d_342_v101_ = _dafny.MultiSet([p2])
        d_343_v102_: _dafny.MultiSet
        d_343_v102_ = _dafny.MultiSet([p0, p3, self.f29])
        default__.m0(False, p0, default__.fm36(len(d_341_v100_), p0, ((d_342_v101_)[p2] if (p2) in (d_342_v101_) else p0), globalState), (d_343_v102_).cardinality, globalState)
        r0 = p1
        d_344_v103_: _dafny.MultiSet
        d_344_v103_ = _dafny.MultiSet([d_297_v61_.f27, p1])
        r1 = ((d_344_v103_)[d_297_v61_.f27] if (d_297_v61_.f27) in (d_344_v103_) else 359)
        d_345_v104_: _dafny.Array
        nw39_ = _dafny.Array(_dafny.Map({}), 8)
        d_345_v104_ = nw39_
        r2 = d_345_v104_
        return r0, r1, r2

    def m3(self, p0, globalState):
        d_346_v0_: _dafny.Map
        d_346_v0_ = _dafny.Map({self.f29: self.f29})
        d_347_v1_: _dafny.Seq
        d_347_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kx"))
        d_348_v2_: _dafny.Map
        d_348_v2_ = _dafny.Map({d_346_v0_: d_347_v1_})
        d_349_v3_: _dafny.Array
        def lambda12_(d_350_i1_):
            return False

        init7_ = lambda12_
        nw40_ = _dafny.Array(None, 5)
        for i0_7_ in range(nw40_.length(0)):
            nw40_[i0_7_] = init7_(i0_7_)
        d_349_v3_ = nw40_
        (globalState).f0 = ((len(((d_348_v2_)[d_346_v0_] if (d_346_v0_) in (d_348_v2_) else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_351_i0_ in range(default__.abs(-463))])))) + (self.f29)) in (_dafny.Map({self.f29: d_349_v3_}))
        index53_ = default__.safeIndex(727, (p0).length(0))
        (p0)[index53_] = self.f29
        index54_ = default__.safeIndex(727, (p0).length(0))
        (p0)[index54_] = (len(d_347_v1_)) - (self.f29)
        d_352_v4_: _dafny.MultiSet
        d_352_v4_ = _dafny.MultiSet([d_349_v3_])
        d_353_i2_: int
        d_353_i2_ = 0
        with _dafny.label("2"):
            while ((d_352_v4_).cardinality) == (523):
                with _dafny.c_label("2"):
                    if (d_353_i2_) >= (100):
                        raise _dafny.Break("2")
                    d_353_i2_ = (d_353_i2_) + (1)
                    index55_ = default__.safeIndex(727, (p0).length(0))
                    (p0)[index55_] = ((p0)[default__.safeIndex(727, (p0).length(0))]) + (self.f29)
                    d_354_v5_: str
                    d_354_v5_ = _dafny.CodePoint('d')
                    d_354_v5_ = d_354_v5_
                    d_355_v6_: bool
                    d_355_v6_ = True
                    d_356_v7_: C0
                    nw41_ = C0()
                    nw41_.ctor__(not(d_355_v6_))
                    d_356_v7_ = nw41_
                    d_357_v8_: _dafny.MultiSet
                    d_357_v8_ = _dafny.MultiSet([d_356_v7_.f27])
                    d_358_v9_: _dafny.Map
                    d_358_v9_ = _dafny.Map({((d_357_v8_)[d_355_v6_] if (d_355_v6_) in (d_357_v8_) else (p0)[default__.safeIndex(727, (p0).length(0))]): d_356_v7_.f27})
                    d_359_v10_: _dafny.Seq
                    d_359_v10_ = _dafny.SeqWithoutIsStrInference([self.f29, ((p0)[default__.safeIndex(727, (p0).length(0))]) * ((p0)[default__.safeIndex(727, (p0).length(0))]), (len(d_358_v9_)) + (self.f29), self.f29, (p0)[default__.safeIndex(727, (p0).length(0))]])
                    d_359_v10_ = (default__.fm37(globalState)) + (default__.fm37(globalState))
                    pass
            pass
        d_360_v11_: _dafny.Array
        nw42_ = _dafny.Array(int(0), 8)
        d_360_v11_ = nw42_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_360_v11_).length(0)):
            d_361_i3_: int = guard_loop_1_
            if (True) and (((0) <= (d_361_i3_)) and ((d_361_i3_) < ((d_360_v11_).length(0)))):
                (d_360_v11_)[(d_361_i3_)] = (d_361_i3_) - ((self.f29) - (self.f29))
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_349_v3_).length(0)):
            d_362_i4_: int = guard_loop_2_
            if (True) and (((0) <= (d_362_i4_)) and ((d_362_i4_) < ((d_349_v3_).length(0)))):
                (d_349_v3_)[(d_362_i4_)] = True
        d_363_v12_: bool
        d_363_v12_ = False
        d_364_v13_: _dafny.Seq
        d_364_v13_ = _dafny.SeqWithoutIsStrInference([False, d_363_v12_, d_363_v12_, (self).fm6(-244, globalState), d_363_v12_])
        d_365_v14_: _dafny.Map
        d_365_v14_ = _dafny.Map({not(d_363_v12_): d_364_v13_})
        hi0_ = ((p0)[default__.safeIndex(727, (p0).length(0))]) * (self.f29)
        for d_366_i5_ in range(len(((d_365_v14_)[d_363_v12_] if (d_363_v12_) in (d_365_v14_) else _dafny.SeqWithoutIsStrInference([d_363_v12_]))), hi0_):
            d_367_v15_: str
            d_367_v15_ = _dafny.CodePoint('e')
            d_368_v16_: _dafny.Map
            d_368_v16_ = _dafny.Map({d_366_i5_: d_367_v15_})
            d_368_v16_ = (d_368_v16_).set((0) - ((len(d_364_v13_)) * (-703)), d_367_v15_)
            d_369_v17_: D12
            d_369_v17_ = D12_DC23(len(d_347_v1_))
            (globalState).f7 = (d_369_v17_).cf30
            d_370_v18_: _dafny.Seq
            d_370_v18_ = _dafny.SeqWithoutIsStrInference([(p0)[default__.safeIndex(727, (p0).length(0))], 651, (p0)[default__.safeIndex(727, (p0).length(0))], self.f29])
            d_371_v19_: _dafny.Map
            d_371_v19_ = _dafny.Map({d_370_v18_: d_360_v11_})
            d_372_v20_: _dafny.Map
            d_372_v20_ = _dafny.Map({((d_371_v19_)[d_370_v18_] if (d_370_v18_) in (d_371_v19_) else d_360_v11_): _dafny.SeqWithoutIsStrInference([d_370_v18_, d_370_v18_])})
            d_373_v21_: _dafny.Seq
            d_373_v21_ = _dafny.SeqWithoutIsStrInference([d_370_v18_])
            d_372_v20_ = (d_372_v20_).set(d_360_v11_, (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(p0)[default__.safeIndex(727, (p0).length(0))], 368, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jcxjkvp")))]) for d_374_i6_ in range(default__.abs(880))])) + (d_373_v21_))
            d_375_v22_: _dafny.MultiSet
            d_375_v22_ = _dafny.MultiSet([d_370_v18_, d_370_v18_, _dafny.SeqWithoutIsStrInference([d_366_i5_ for d_376_i7_ in range(default__.abs(-816))]), _dafny.SeqWithoutIsStrInference([d_366_i5_, (0) - (d_366_i5_)]), (d_370_v18_).set(default__.safeIndex(938, len(d_370_v18_)), 206)])
            d_377_v23_: _dafny.MultiSet
            d_377_v23_ = _dafny.MultiSet([d_347_v1_])
            rhs36_ = (len(d_346_v0_)) == (509)
            rhs37_ = (((d_375_v22_)[d_370_v18_] if (d_370_v18_) in (d_375_v22_) else ((d_377_v23_)[d_347_v1_] if (d_347_v1_) in (d_377_v23_) else (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ilycvcw"))))))) >= ((33) - (self.f29))
            lhs40_ = globalState
            lhs41_ = globalState
            lhs40_.f21 = rhs36_
            lhs41_.f13 = rhs37_

    def m1(self, p0, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: bool = False
        d_378_v0_: bool
        d_378_v0_ = False
        (globalState).f0 = ((p0) != (p0) if d_378_v0_ else d_378_v0_)
        if (self).fm7(d_378_v0_, (self.f29 if d_378_v0_ else p0), (p0) == ((_dafny.MultiSet([p0])).cardinality), globalState):
            d_379_v1_: _dafny.MultiSet
            d_379_v1_ = _dafny.MultiSet([self.f29])
            d_380_v2_: _dafny.Map
            d_380_v2_ = _dafny.Map({d_379_v1_: d_378_v0_})
            d_381_v3_: _dafny.Set
            d_381_v3_ = _dafny.Set({False})
            d_382_v4_: _dafny.Seq
            d_382_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "by"))
            d_383_v5_: _dafny.Seq
            d_383_v5_ = _dafny.SeqWithoutIsStrInference([((d_379_v1_)[len(d_381_v3_)] if (len(d_381_v3_)) in (d_379_v1_) else p0), p0, len(d_382_v4_)])
            source8_ = (d_380_v2_).set(_dafny.MultiSet((d_383_v5_).set(default__.safeIndex(143, len(d_383_v5_)), self.f29)), d_378_v0_)
            d_384___mcc_h0_ = source8_
            d_385_cf27_ = d_384___mcc_h0_
            (globalState).f13 = d_378_v0_
            (globalState).f7 = p0
            d_386_v6_: C0
            nw43_ = C0()
            nw43_.ctor__(d_378_v0_)
            d_386_v6_ = nw43_
            (globalState).f21 = ((d_382_v4_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sbarm")))) == (d_382_v4_)
            d_387_v7_: D13
            d_387_v7_ = D13_DC26(d_382_v4_, d_378_v0_, d_378_v0_)
            d_388_v8_: _dafny.Set
            d_388_v8_ = _dafny.Set({d_382_v4_, (d_387_v7_).cf33, d_382_v4_})
            (globalState).f10 = (d_388_v8_).intersection(_dafny.Set({d_382_v4_, d_382_v4_}))
            d_389_v9_: C0
            nw44_ = C0()
            nw44_.ctor__(d_378_v0_)
            d_389_v9_ = nw44_
            (globalState).f7 = self.f29
            d_390_v10_: str
            d_390_v10_ = _dafny.CodePoint('l')
            d_391_v11_: _dafny.Set
            d_391_v11_ = _dafny.Set({d_390_v10_})
            d_392_v12_: _dafny.Set
            d_392_v12_ = _dafny.Set({(d_391_v11_) | (_dafny.Set({d_390_v10_}))})
            d_392_v12_ = default__.fm38(globalState)
        elif True:
            if False:
                (globalState).f7 = (366 if d_378_v0_ else p0)
                d_393_v13_: _dafny.Array
                nw45_ = _dafny.Array(_dafny.Map({}), 11)
                d_393_v13_ = nw45_
                d_394_v14_: str
                d_394_v14_ = _dafny.CodePoint('b')
                d_395_v15_: _dafny.Map
                d_395_v15_ = _dafny.Map({d_378_v0_: d_394_v14_})
                index56_ = default__.safeIndex(938, (d_393_v13_).length(0))
                (d_393_v13_)[index56_] = d_395_v15_
                d_396_v16_: _dafny.Map
                d_396_v16_ = _dafny.Map({True: d_378_v0_})
                index57_ = default__.safeIndex(938, (d_393_v13_).length(0))
                (d_393_v13_)[index57_] = (_dafny.Map({default__.fm2(d_396_v16_, globalState): d_394_v14_})) | (d_395_v15_)
                d_397_v17_: D7
                d_397_v17_ = D7_DC15(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kueylnvh"))) for d_398_i0_ in range(default__.abs(279))]))
                d_399_v18_: D7
                d_399_v18_ = D7_DC17(d_397_v17_)
                d_400_v19_: _dafny.Seq
                d_400_v19_ = _dafny.SeqWithoutIsStrInference([d_399_v18_])
                (globalState).f21 = (d_400_v19_) < ((d_400_v19_) + (d_400_v19_))
                d_401_v20_: _dafny.Array
                nw46_ = _dafny.Array(int(0), 3)
                d_401_v20_ = nw46_
                r0 = d_401_v20_
                (globalState).f0 = d_378_v0_
            elif True:
                (globalState).f16 = 748
                (globalState).f7 = self.f29
                d_402_v21_: _dafny.Array
                nw47_ = _dafny.Array(int(0), 6)
                d_402_v21_ = nw47_
                index58_ = default__.safeIndex(810, (d_402_v21_).length(0))
                (d_402_v21_)[index58_] = (0) - (self.f29)
                d_403_v22_: _dafny.Map
                d_403_v22_ = _dafny.Map({d_378_v0_: p0})
                index59_ = default__.safeIndex(810, (d_402_v21_).length(0))
                (d_402_v21_)[index59_] = default__.safeDivisionInt(((d_403_v22_)[d_378_v0_] if (d_378_v0_) in (d_403_v22_) else self.f29), p0)
                d_404_v23_: _dafny.MultiSet
                d_404_v23_ = _dafny.MultiSet([d_378_v0_])
                d_405_v24_: _dafny.Seq
                d_405_v24_ = _dafny.SeqWithoutIsStrInference([d_404_v23_])
                d_406_v25_: _dafny.MultiSet
                d_406_v25_ = (d_405_v24_)[default__.safeIndex(len(d_403_v22_), len(d_405_v24_))]
                d_406_v25_ = d_406_v25_
                (globalState).f13 = (d_378_v0_) or (d_378_v0_)
            d_407_v26_: C0
            nw48_ = C0()
            nw48_.ctor__(d_378_v0_)
            d_407_v26_ = nw48_
            (globalState).f7 = p0
            d_408_v27_: _dafny.Map
            d_408_v27_ = _dafny.Map({d_407_v26_.f27: (self).fm29(globalState)})
            d_409_v28_: _dafny.MultiSet
            d_409_v28_ = _dafny.MultiSet([d_408_v27_])
            d_410_v29_: _dafny.Map
            d_410_v29_ = _dafny.Map({(_dafny.MultiSet([self.f29])).cardinality: d_407_v26_.f27})
            d_411_v30_: _dafny.Seq
            d_411_v30_ = _dafny.SeqWithoutIsStrInference([len(d_410_v29_), self.f29])
            (globalState).f2 = (d_409_v28_).issubset((d_409_v28_) - (_dafny.MultiSet([_dafny.Map({True: (_dafny.MultiSet(d_411_v30_)).cardinality}), d_408_v27_, d_408_v27_, d_408_v27_, d_408_v27_])))
            (globalState).f7 = p0
        d_412_v31_: C0
        nw49_ = C0()
        nw49_.ctor__(d_378_v0_)
        d_412_v31_ = nw49_
        d_413_i1_: int
        d_413_i1_ = 0
        with _dafny.label("3"):
            while (self.f29) == ((self).fm29(globalState)):
                with _dafny.c_label("3"):
                    if (d_413_i1_) >= (100):
                        raise _dafny.Break("3")
                    d_413_i1_ = (d_413_i1_) + (1)
                    (globalState).f13 = not(d_378_v0_)
                    d_414_v32_: _dafny.Array
                    def lambda13_(d_415_v0_, d_416_v31_):
                        def lambda14_(d_417_i2_):
                            return default__.safeModuloInt(d_417_i2_, (_dafny.MultiSet([d_415_v0_, not(d_416_v31_.f27)])).cardinality)

                        return lambda14_

                    init8_ = lambda13_(d_378_v0_, d_412_v31_)
                    nw50_ = _dafny.Array(None, 2)
                    for i0_8_ in range(nw50_.length(0)):
                        nw50_[i0_8_] = init8_(i0_8_)
                    d_414_v32_ = nw50_
                    d_418_v33_: _dafny.Set
                    d_418_v33_ = _dafny.Set({d_414_v32_, d_414_v32_})
                    d_419_v34_: _dafny.Seq
                    d_419_v34_ = _dafny.SeqWithoutIsStrInference([d_378_v0_, d_378_v0_, d_412_v31_.f27])
                    d_420_v35_: _dafny.Seq
                    d_420_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xjuxav"))
                    d_421_v36_: _dafny.Array
                    nw51_ = _dafny.Array(None, 11)
                    nw51_[int(0)] = len(default__.fm1(len(d_418_v33_), globalState))
                    nw51_[int(1)] = p0
                    nw51_[int(2)] = len(d_419_v34_)
                    nw51_[int(3)] = (len(d_420_v35_)) + (-830)
                    nw51_[int(4)] = p0
                    nw51_[int(5)] = (len(d_420_v35_)) - (len(d_420_v35_))
                    nw51_[int(6)] = p0
                    nw51_[int(7)] = self.f29
                    nw51_[int(8)] = self.f29
                    nw51_[int(9)] = default__.fm0(-110, p0, self.f29, d_378_v0_, globalState)
                    nw51_[int(10)] = -680
                    d_421_v36_ = nw51_
                    index60_ = default__.safeIndex(572, (d_414_v32_).length(0))
                    (d_414_v32_)[index60_] = self.f29
                    index61_ = default__.safeIndex(572, (d_414_v32_).length(0))
                    (d_414_v32_)[index61_] = self.f29
                    d_422_v37_: _dafny.Map
                    d_422_v37_ = _dafny.Map({(d_414_v32_)[default__.safeIndex(572, (d_414_v32_).length(0))]: d_412_v31_.f27})
                    d_423_v38_: D1
                    d_423_v38_ = D1_DC2((d_422_v37_) | (d_422_v37_))
                    d_423_v38_ = D1_DC2(d_422_v37_)
                    d_424_v39_: _dafny.Map
                    d_424_v39_ = _dafny.Map({d_412_v31_.f27: self.f29})
                    d_424_v39_ = ((d_424_v39_) | (d_424_v39_)) | (_dafny.Map({d_378_v0_: 811}))
                    pass
            pass
        d_425_v40_: C0
        nw52_ = C0()
        nw52_.ctor__(d_412_v31_.f27)
        d_425_v40_ = nw52_
        d_426_v41_: _dafny.Map
        d_426_v41_ = _dafny.Map({not(d_378_v0_): d_425_v40_.f27})
        rhs38_ = (self).fm29(globalState)
        rhs39_ = d_426_v41_
        lhs42_ = self
        lhs42_.f29 = rhs38_
        d_426_v41_ = rhs39_
        d_427_v42_: _dafny.Array
        nw53_ = _dafny.Array(int(0), 27)
        d_427_v42_ = nw53_
        r0 = d_427_v42_
        r1 = d_425_v40_.f27
        return r0, r1

    @property
    def f30(self):
        return self._f30

class C2(T1, T0):
    def  __init__(self):
        self._f28: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    def ctor__(self, f28):
        (self)._f28 = f28

    def fm7(self, p0, p1, p2, globalState):
        return (((self).f28) > ((0) - ((self).f28))) == ((False if not(not(True)) else False))

    def fm5(self, p0, p1, p2, p3, globalState):
        source9_ = D1_DC2(_dafny.Map({-844: False}))
        if source9_.is_DC3:
            d_428___mcc_h0_ = source9_.cf3
            d_429_cf3_ = d_428___mcc_h0_
            return _dafny.Map({d_429_cf3_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lcwihbpu"))})
        elif True:
            d_430___mcc_h1_ = source9_.cf2
            d_431_cf2_ = d_430___mcc_h1_
            return (_dafny.Map({not(False): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "duyckw"))})) | (_dafny.Map({not(True): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mtfgtk"))}))

    def fm6(self, p0, globalState):
        return False

    def fm17(self, p0, p1, p2, p3, globalState):
        if not (True) or (False):
            return (self).f28
        elif True:
            return (self).f28

    def m2(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: _dafny.Array = _dafny.Array(None, 0)
        d_432_v0_: _dafny.Seq
        d_432_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "klcc"))
        hi1_ = len(d_432_v0_)
        for d_433_i0_ in range((self).f28, hi1_):
            (globalState).f21 = p1
            d_434_v1_: C0
            nw54_ = C0()
            nw54_.ctor__(p1)
            d_434_v1_ = nw54_
            d_435_v2_: _dafny.MultiSet
            d_435_v2_ = _dafny.MultiSet([d_434_v1_.f27])
            d_436_v3_: _dafny.MultiSet
            d_436_v3_ = _dafny.MultiSet([default__.safeDivisionInt((d_435_v2_).cardinality, (self).f28), d_433_i0_])
            d_437_v4_: D4
            d_437_v4_ = D4_DC9((0) - (p0))
            rhs40_ = d_434_v1_.f27
            rhs41_ = ((d_436_v3_)[(self).f28] if ((self).f28) in (d_436_v3_) else (0) - (((d_436_v3_)[p0] if (p0) in (d_436_v3_) else p3)))
            rhs42_ = ((d_437_v4_).cf11) + (p3)
            rhs43_ = (0) - (default__.safeModuloInt(p0, d_433_i0_))
            lhs43_ = d_434_v1_
            lhs44_ = globalState
            lhs45_ = globalState
            lhs46_ = globalState
            lhs43_.f27 = rhs40_
            lhs44_.f7 = rhs41_
            lhs45_.f7 = rhs42_
            lhs46_.f7 = rhs43_
            d_438_v5_: _dafny.Seq
            d_438_v5_ = _dafny.SeqWithoutIsStrInference([p1])
            d_439_v6_: _dafny.Seq
            d_439_v6_ = d_438_v5_
            d_440_v7_: str
            d_440_v7_ = _dafny.CodePoint('q')
            d_441_v8_: _dafny.Array
            nw55_ = _dafny.Array(None, 19)
            nw55_[int(0)] = d_438_v5_
            nw55_[int(1)] = d_438_v5_
            nw55_[int(2)] = (d_438_v5_).set(default__.safeIndex(973, len(d_438_v5_)), d_434_v1_.f27)
            nw55_[int(3)] = d_438_v5_
            nw55_[int(4)] = (d_439_v6_)
            nw55_[int(5)] = _dafny.SeqWithoutIsStrInference([d_434_v1_.f27])
            nw55_[int(6)] = (d_438_v5_) + ((default__.fm18(d_434_v1_.f27, _dafny.CodePoint('p'), d_434_v1_.f27, globalState)).set(default__.safeIndex(d_433_i0_, len(default__.fm18(d_434_v1_.f27, _dafny.CodePoint('p'), d_434_v1_.f27, globalState))), d_434_v1_.f27))
            nw55_[int(7)] = d_438_v5_
            nw55_[int(8)] = d_438_v5_
            nw55_[int(9)] = (d_438_v5_) + (d_438_v5_)
            nw55_[int(10)] = d_438_v5_
            nw55_[int(11)] = _dafny.SeqWithoutIsStrInference([d_434_v1_.f27])
            nw55_[int(12)] = d_438_v5_
            nw55_[int(13)] = _dafny.SeqWithoutIsStrInference([p1, d_434_v1_.f27, p1, d_434_v1_.f27])
            nw55_[int(14)] = (default__.fm18(not(d_434_v1_.f27), d_440_v7_, d_434_v1_.f27, globalState)) + (d_438_v5_)
            nw55_[int(15)] = (d_438_v5_).set(default__.safeIndex(p0, len(d_438_v5_)), d_434_v1_.f27)
            nw55_[int(16)] = _dafny.SeqWithoutIsStrInference([d_434_v1_.f27])
            nw55_[int(17)] = d_438_v5_
            nw55_[int(18)] = d_438_v5_
            d_441_v8_ = nw55_
            index62_ = default__.safeIndex(977, (d_441_v8_).length(0))
            (d_441_v8_)[index62_] = d_438_v5_
            index63_ = default__.safeIndex(977, (d_441_v8_).length(0))
            (d_441_v8_)[index63_] = d_438_v5_
        d_442_v9_: _dafny.Array
        nw56_ = _dafny.Array(D3.default()(), 20)
        d_442_v9_ = nw56_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_442_v9_).length(0)):
            d_443_i1_: int = guard_loop_3_
            if (True) and (((0) <= (d_443_i1_)) and ((d_443_i1_) < ((d_442_v9_).length(0)))):
                (d_442_v9_)[(d_443_i1_)] = D3_DC6(d_432_v0_, _dafny.CodePoint('v'), p1)
        d_444_v10_: D1
        d_444_v10_ = D1_DC3(p1)
        d_445_v11_: _dafny.Map
        d_445_v11_ = _dafny.Map({p0: (d_444_v10_).cf3})
        d_446_v12_: D4
        d_446_v12_ = D4_DC7(len(d_445_v11_))
        d_447_v13_: _dafny.Set
        d_447_v13_ = _dafny.Set({False})
        d_448_v14_: _dafny.Map
        d_448_v14_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_449_i2_ in range(default__.abs(-94))]): d_447_v13_})
        d_450_v15_: _dafny.Seq
        d_450_v15_ = _dafny.SeqWithoutIsStrInference([p0, p0])
        d_451_v16_: _dafny.Seq
        d_451_v16_ = _dafny.SeqWithoutIsStrInference([p1, p1, p1])
        d_452_v17_: _dafny.Array
        nw57_ = _dafny.Array(None, 13)
        nw57_[int(0)] = ((d_448_v14_)[d_432_v0_] if (d_432_v0_) in (d_448_v14_) else d_447_v13_)
        nw57_[int(1)] = (_dafny.Set({p1})).intersection(d_447_v13_)
        nw57_[int(2)] = (d_447_v13_).intersection(d_447_v13_)
        nw57_[int(3)] = d_447_v13_
        nw57_[int(4)] = d_447_v13_
        nw57_[int(5)] = d_447_v13_
        nw57_[int(6)] = d_447_v13_
        nw57_[int(7)] = d_447_v13_
        nw57_[int(8)] = default__.fm19(len(d_450_v15_), (d_451_v16_)[default__.safeIndex(702, len(d_451_v16_))], p1, -900, globalState)
        nw57_[int(9)] = ((d_447_v13_)) | (_dafny.Set({p1}))
        nw57_[int(10)] = d_447_v13_
        nw57_[int(11)] = (d_447_v13_) | (d_447_v13_)
        nw57_[int(12)] = (d_447_v13_).intersection(_dafny.Set({p1, p1, p1, p1, p1}))
        d_452_v17_ = nw57_
        index64_ = default__.safeIndex(693, (d_452_v17_).length(0))
        (d_452_v17_)[index64_] = d_447_v13_
        d_453_v18_: C0
        nw58_ = C0()
        nw58_.ctor__((p3) < ((self).f28))
        d_453_v18_ = nw58_
        index65_ = default__.safeIndex(693, (d_452_v17_).length(0))
        rhs44_ = (((self).f28) * (p3)) * (default__.safeDivisionInt(p3, len(d_450_v15_)))
        rhs45_ = p1
        rhs46_ = d_446_v12_
        rhs47_ = d_447_v13_
        rhs48_ = d_453_v18_
        lhs47_ = globalState
        lhs48_ = globalState
        lhs49_ = d_452_v17_
        lhs50_ = default__.safeIndex(693, (d_452_v17_).length(0))
        lhs47_.f7 = rhs44_
        lhs48_.f21 = rhs45_
        d_446_v12_ = rhs46_
        lhs49_[lhs50_] = rhs47_
        d_453_v18_ = rhs48_
        (d_453_v18_).f27 = d_453_v18_.f27
        d_454_v19_: _dafny.Array
        def lambda15_(d_455_v18_):
            def lambda16_(d_456_i3_):
                return (d_455_v18_.f27 if d_455_v18_.f27 else d_455_v18_.f27)

            return lambda16_

        init9_ = lambda15_(d_453_v18_)
        nw59_ = _dafny.Array(None, 17)
        for i0_9_ in range(nw59_.length(0)):
            nw59_[i0_9_] = init9_(i0_9_)
        d_454_v19_ = nw59_
        index66_ = default__.safeIndex(337, (d_454_v19_).length(0))
        (d_454_v19_)[index66_] = p1
        d_457_v20_: str
        d_457_v20_ = _dafny.CodePoint('a')
        d_458_v21_: _dafny.Seq
        d_458_v21_ = default__.fm18(d_453_v18_.f27, d_457_v20_, False, globalState)
        index67_ = default__.safeIndex(337, (d_454_v19_).length(0))
        rhs49_ = (self).fm17(d_453_v18_.f27, _dafny.SeqWithoutIsStrInference([(self).f28 for d_459_i4_ in range(default__.abs(518))]), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_460_i5_ in range(default__.abs(191))])), d_458_v21_, globalState)
        rhs50_ = (d_453_v18_.f27) or (d_453_v18_.f27)
        lhs51_ = globalState
        lhs52_ = d_454_v19_
        lhs53_ = default__.safeIndex(337, (d_454_v19_).length(0))
        lhs51_.f7 = rhs49_
        lhs52_[lhs53_] = rhs50_
        d_461_v22_: _dafny.Array
        def lambda17_(d_462_v10_):
            def lambda18_(d_463_i7_):
                return d_462_v10_

            return lambda18_

        init10_ = lambda17_(d_444_v10_)
        nw60_ = _dafny.Array(None, 2)
        for i0_10_ in range(nw60_.length(0)):
            nw60_[i0_10_] = init10_(i0_10_)
        d_461_v22_ = nw60_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_461_v22_).length(0)):
            d_464_i6_: int = guard_loop_4_
            if (True) and (((0) <= (d_464_i6_)) and ((d_464_i6_) < ((d_461_v22_).length(0)))):
                (d_461_v22_)[(d_464_i6_)] = (d_444_v10_ if ((d_445_v11_)[len(_dafny.Set({p0, (self).f28}))] if (len(_dafny.Set({p0, (self).f28}))) in (d_445_v11_) else d_453_v18_.f27) else d_444_v10_)
        d_465_v23_: _dafny.Map
        d_465_v23_ = _dafny.Map({p1: p1})
        r0 = not(default__.fm2((d_465_v23_).set(d_453_v18_.f27, d_453_v18_.f27), globalState))
        pat_let_tv6_ = p3
        pat_let_tv7_ = d_454_v19_
        pat_let_tv8_ = d_454_v19_
        pat_let_tv9_ = p3
        pat_let_tv10_ = p3
        def lambda19_(source10_):
            if source10_.is_DC8:
                d_466___mcc_h0_ = source10_.cf10
                d_467_cf10_ = d_466___mcc_h0_
                return pat_let_tv6_
            elif source10_.is_DC9:
                d_468___mcc_h1_ = source10_.cf11
                d_469_cf11_ = d_468___mcc_h1_
                return len((D12_DC22(_dafny.Map({(pat_let_tv8_)[default__.safeIndex(337, (pat_let_tv7_).length(0))]: pat_let_tv9_}))).cf29)
            elif source10_.is_DC10:
                d_470___mcc_h2_ = source10_.cf12
                d_471___mcc_h3_ = source10_.cf13
                d_472___mcc_h4_ = source10_.cf14
                d_473___mcc_h5_ = source10_.cf15
                d_474___mcc_h6_ = source10_.cf16
                d_475_cf16_ = d_474___mcc_h6_
                d_476_cf15_ = d_473___mcc_h5_
                d_477_cf14_ = d_472___mcc_h4_
                d_478_cf13_ = d_471___mcc_h3_
                d_479_cf12_ = d_470___mcc_h2_
                return (pat_let_tv10_) + (d_478_cf13_)
            elif True:
                d_480___mcc_h7_ = source10_.cf9
                d_481_cf9_ = d_480___mcc_h7_
                return (self).f28

        r1 = lambda19_(D4_DC7(406))
        d_482_v24_: _dafny.Array
        nw61_ = _dafny.Array(_dafny.Map({}), 27)
        d_482_v24_ = nw61_
        r2 = d_482_v24_
        return r0, r1, r2

    def m3(self, p0, globalState):
        (globalState).f16 = 644
        d_483_v0_: D12
        d_483_v0_ = D12_DC23(7)
        def iife42_(_pat_let6_0):
            def iife43_(d_484_dt__update__tmp_h0_):
                def iife44_(_pat_let7_0):
                    def iife45_(d_485_dt__update_hcf30_h0_):
                        return D12_DC23(d_485_dt__update_hcf30_h0_)
                    return iife45_(_pat_let7_0)
                return iife44_(((self).f28) * ((self).f28))
            return iife43_(_pat_let6_0)
        source11_ = iife42_(d_483_v0_)
        if source11_.is_DC23:
            d_486___mcc_h0_ = source11_.cf30
            d_487_cf30_ = d_486___mcc_h0_
            d_488_v1_: bool
            d_488_v1_ = True
            d_489_v2_: str
            d_489_v2_ = _dafny.CodePoint('n')
            d_490_v3_: _dafny.Map
            d_490_v3_ = _dafny.Map({not(d_488_v1_): d_489_v2_})
            d_491_v4_: _dafny.Map
            d_491_v4_ = _dafny.Map({d_487_cf30_: len(d_490_v3_)})
            d_492_v5_: _dafny.Map
            d_492_v5_ = _dafny.Map({d_488_v1_: len(d_491_v4_)})
            if ((d_492_v5_).set(default__.fm2(default__.fm20(d_487_cf30_, d_488_v1_, d_488_v1_, d_488_v1_, globalState), globalState), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bkpcvyger"))))) != (d_492_v5_):
                d_493_v6_: _dafny.MultiSet
                d_493_v6_ = _dafny.MultiSet([d_488_v1_])
                d_488_v1_ = ((d_493_v6_) - (_dafny.MultiSet([d_488_v1_]))).ispropersubset(d_493_v6_)
                d_494_v7_: _dafny.Array
                nw62_ = _dafny.Array(_dafny.Array(None, 0), 24)
                d_494_v7_ = nw62_
                index68_ = default__.safeIndex(115, (d_494_v7_).length(0))
                (d_494_v7_)[index68_] = p0
                index69_ = default__.safeIndex(115, (d_494_v7_).length(0))
                (d_494_v7_)[index69_] = p0
                d_495_v8_: _dafny.Array
                nw63_ = _dafny.Array(_dafny.Set({}), 6)
                d_495_v8_ = nw63_
                d_496_v9_: _dafny.Map
                d_496_v9_ = _dafny.Map({d_488_v1_: d_488_v1_})
                d_497_v10_: _dafny.Set
                d_497_v10_ = _dafny.Set({d_496_v9_})
                index70_ = default__.safeIndex(135, (d_495_v8_).length(0))
                (d_495_v8_)[index70_] = d_497_v10_
                index71_ = default__.safeIndex(135, (d_495_v8_).length(0))
                (d_495_v8_)[index71_] = d_497_v10_
                d_498_v11_: _dafny.Set
                d_498_v11_ = _dafny.Set({d_488_v1_})
                d_499_v12_: _dafny.Map
                d_499_v12_ = _dafny.Map({len(d_498_v11_): d_488_v1_})
                (globalState).f2 = (d_487_cf30_) not in (d_499_v12_)
                d_500_v13_: _dafny.Seq
                d_500_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aaacch"))
                d_501_v14_: _dafny.Seq
                d_501_v14_ = _dafny.SeqWithoutIsStrInference([(d_500_v13_) + (_dafny.SeqWithoutIsStrInference([d_489_v2_ for d_502_i0_ in range(default__.abs(966))])), d_500_v13_, d_500_v13_])
                index72_ = default__.safeIndex(115, (d_494_v7_).length(0))
                rhs51_ = d_488_v1_
                rhs52_ = not(d_488_v1_)
                rhs53_ = (d_494_v7_)[default__.safeIndex(115, (d_494_v7_).length(0))]
                rhs54_ = (d_501_v14_)[default__.safeIndex((0) - ((self).f28), len(d_501_v14_))]
                rhs55_ = d_487_cf30_
                lhs54_ = globalState
                lhs55_ = globalState
                lhs56_ = d_494_v7_
                lhs57_ = default__.safeIndex(115, (d_494_v7_).length(0))
                lhs58_ = globalState
                lhs59_ = globalState
                lhs54_.f0 = rhs51_
                lhs55_.f0 = rhs52_
                lhs56_[lhs57_] = rhs53_
                lhs58_.f19 = rhs54_
                lhs59_.f16 = rhs55_
            elif True:
                d_489_v2_ = _dafny.CodePoint('a')
                (self).m8(not(not(not(d_488_v1_))), d_488_v1_, d_488_v1_, globalState)
                d_503_v15_: _dafny.Seq
                d_503_v15_ = _dafny.SeqWithoutIsStrInference([d_488_v1_])
                rhs56_ = (0) - ((0) - ((0) - (default__.safeModuloInt(d_487_cf30_, (self).f28))))
                rhs57_ = len(d_503_v15_)
                lhs60_ = globalState
                lhs61_ = globalState
                lhs60_.f7 = rhs56_
                lhs61_.f7 = rhs57_
                d_504_v16_: _dafny.Seq
                d_504_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "olkogre"))
                d_505_v17_: _dafny.Map
                d_505_v17_ = _dafny.Map({(d_504_v16_).set(default__.safeIndex(d_487_cf30_, len(d_504_v16_)), d_489_v2_): (not(d_488_v1_) if False else False)})
                rhs58_ = d_488_v1_
                rhs59_ = d_505_v17_
                rhs60_ = d_488_v1_
                lhs62_ = globalState
                lhs63_ = globalState
                lhs62_.f0 = rhs58_
                d_505_v17_ = rhs59_
                lhs63_.f2 = rhs60_
                (globalState).f16 = (self).f28
            index73_ = default__.safeIndex(341, (p0).length(0))
            (p0)[index73_] = d_487_cf30_
            index74_ = default__.safeIndex(341, (p0).length(0))
            rhs61_ = (self).f28
            rhs62_ = (-20) - (d_487_cf30_)
            rhs63_ = d_489_v2_
            lhs64_ = p0
            lhs65_ = default__.safeIndex(341, (p0).length(0))
            lhs66_ = globalState
            lhs64_[lhs65_] = rhs61_
            lhs66_.f7 = rhs62_
            d_489_v2_ = rhs63_
            d_506_v18_: _dafny.Array
            nw64_ = _dafny.Array(int(0), 28)
            d_506_v18_ = nw64_
            d_507_v19_: _dafny.MultiSet
            d_507_v19_ = _dafny.MultiSet([d_506_v18_])
            d_508_v20_: _dafny.MultiSet
            d_508_v20_ = d_507_v19_
            d_509_v21_: _dafny.MultiSet
            d_509_v21_ = (d_508_v20_)
            source12_ = d_508_v20_
            d_510___mcc_h3_ = source12_
            d_511_cf4_ = d_510___mcc_h3_
            d_512_v22_: C0
            nw65_ = C0()
            nw65_.ctor__(not(d_488_v1_))
            d_512_v22_ = nw65_
            (globalState).f16 = d_487_cf30_
            (globalState).f16 = (self).f28
            d_513_v23_: _dafny.Seq
            d_513_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "am"))
            d_514_v24_: _dafny.Set
            d_514_v24_ = _dafny.Set({default__.safeDivisionInt(-250, (self).f28), (p0)[default__.safeIndex(341, (p0).length(0))], (self).f28, len(((d_513_v23_) + (d_513_v23_)).set(default__.safeIndex(d_487_cf30_, len((d_513_v23_) + (d_513_v23_))), d_489_v2_))})
            d_514_v24_ = ((d_514_v24_) - (d_514_v24_)) - (d_514_v24_)
            d_515_v25_: _dafny.Set
            d_515_v25_ = _dafny.Set({d_487_cf30_})
            d_516_v26_: _dafny.Seq
            d_516_v26_ = _dafny.SeqWithoutIsStrInference([d_488_v1_, True])
            d_517_v27_: _dafny.Seq
            d_517_v27_ = _dafny.SeqWithoutIsStrInference([len(d_516_v26_), d_487_cf30_])
            d_518_v28_: _dafny.MultiSet
            d_518_v28_ = _dafny.MultiSet([d_487_cf30_])
            d_519_v29_: _dafny.MultiSet
            d_519_v29_ = _dafny.MultiSet([(p0)[default__.safeIndex(341, (p0).length(0))], (d_517_v27_)[default__.safeIndex((d_518_v28_).cardinality, len(d_517_v27_))]])
            d_515_v25_ = ((d_515_v25_).intersection(_dafny.Set({d_487_cf30_, (self).f28, d_487_cf30_, (d_518_v28_).cardinality}))) | (d_515_v25_)
        elif source11_.is_DC22:
            d_520___mcc_h1_ = source11_.cf29
            d_521_cf29_ = d_520___mcc_h1_
            d_522_v30_: _dafny.Array
            def lambda20_(d_523_i1_):
                return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "liqnc"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mjrvidhd")))

            init11_ = lambda20_
            nw66_ = _dafny.Array(None, 17)
            for i0_11_ in range(nw66_.length(0)):
                nw66_[i0_11_] = init11_(i0_11_)
            d_522_v30_ = nw66_
            d_522_v30_ = d_522_v30_
            d_524_v31_: _dafny.Map
            d_524_v31_ = _dafny.Map({(0) - ((self).f28): (self).f28})
            d_525_v32_: bool
            d_525_v32_ = True
            d_526_v34_: D13
            def iife46_():
                coll30_ = _dafny.Map()
                compr_30_: int
                for compr_30_ in (default__.fm21(globalState)).Elements:
                    d_527_v33_: int = compr_30_
                    if (d_527_v33_) in (default__.fm21(globalState)):
                        coll30_[default__.safeModuloInt(d_527_v33_, (self).f28)] = (self).f28
                return _dafny.Map(coll30_)
            d_526_v34_ = D13_DC25(iife46_()
)
            d_528_v35_: _dafny.Map
            d_528_v35_ = _dafny.Map({(False) and (d_525_v32_): (d_526_v34_).cf32})
            d_524_v31_ = ((d_528_v35_)[d_525_v32_] if (d_525_v32_) in (d_528_v35_) else d_524_v31_)
            d_529_v36_: _dafny.Array
            nw67_ = _dafny.Array(_dafny.MultiSet({}), 3)
            d_529_v36_ = nw67_
            d_530_v37_: _dafny.MultiSet
            d_530_v37_ = _dafny.MultiSet([p0, p0, p0])
            d_531_v38_: _dafny.MultiSet
            d_531_v38_ = d_530_v37_
            index75_ = default__.safeIndex(962, (d_529_v36_).length(0))
            (d_529_v36_)[index75_] = d_531_v38_
            index76_ = default__.safeIndex(962, (d_529_v36_).length(0))
            (d_529_v36_)[index76_] = d_531_v38_
            if True:
                d_532_v39_: _dafny.Seq
                d_532_v39_ = _dafny.SeqWithoutIsStrInference([-367])
                d_533_v40_: _dafny.Seq
                d_533_v40_ = _dafny.SeqWithoutIsStrInference([d_525_v32_, d_525_v32_])
                d_534_v41_: _dafny.Seq
                d_534_v41_ = d_533_v40_
                (globalState).f7 = ((self).f28 if d_525_v32_ else default__.fm0((self).f28, (self).fm17(d_525_v32_, d_532_v39_, (self).f28, d_534_v41_, globalState), (self).f28, d_525_v32_, globalState))
                d_535_v42_: str
                d_535_v42_ = _dafny.CodePoint('n')
                d_536_v43_: _dafny.Map
                d_536_v43_ = _dafny.Map({(0) - ((self).f28): d_535_v42_})
                d_537_v44_: _dafny.Seq
                d_537_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mg"))
                d_536_v43_ = (d_536_v43_).set(366, (d_537_v44_)[default__.safeIndex((self).f28, len(d_537_v44_))])
                d_538_v45_: C0
                nw68_ = C0()
                nw68_.ctor__(d_525_v32_)
                d_538_v45_ = nw68_
                d_533_v40_ = d_533_v40_
                d_539_v46_: _dafny.Map
                d_539_v46_ = _dafny.Map({p0: d_525_v32_})
                d_539_v46_ = (_dafny.Map({p0: d_525_v32_})) | (d_539_v46_)
            elif True:
                d_540_v47_: C0
                nw69_ = C0()
                nw69_.ctor__(d_525_v32_)
                d_540_v47_ = nw69_
                d_541_v48_: _dafny.MultiSet
                d_541_v48_ = _dafny.MultiSet([d_540_v47_])
                (globalState).f7 = (((self).f28) * ((self).f28)) - (default__.safeDivisionInt(((d_541_v48_)[d_540_v47_] if (d_540_v47_) in (d_541_v48_) else (self).f28), (self).f28))
                d_542_v49_: _dafny.Seq
                d_542_v49_ = _dafny.SeqWithoutIsStrInference([False])
                d_543_v50_: _dafny.Seq
                d_543_v50_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))
                d_544_v51_: str
                d_544_v51_ = _dafny.CodePoint('s')
                d_545_v52_: D3
                d_545_v52_ = D3_DC6(d_543_v50_, d_544_v51_, d_540_v47_.f27)
                d_546_v53_: _dafny.Map
                d_546_v53_ = _dafny.Map({d_542_v49_: (d_545_v52_).cf6})
                default__.m0(d_540_v47_.f27, (self).f28, (_dafny.Map({(d_542_v49_).set(default__.safeIndex((self).f28, len(d_542_v49_)), d_540_v47_.f27): d_543_v50_}) if d_540_v47_.f27 else d_546_v53_), (self).f28, globalState)
                (globalState).f16 = (self).f28
                d_547_v54_: _dafny.Array
                nw70_ = _dafny.Array(_dafny.Map({}), 19)
                d_547_v54_ = nw70_
                d_547_v54_ = d_547_v54_
                (globalState).f21 = False
        elif True:
            d_548___mcc_h2_ = source11_.cf31
            d_549_cf31_ = d_548___mcc_h2_
            d_550_v55_: _dafny.Map
            d_550_v55_ = _dafny.Map({False: (self).f28})
            d_551_v56_: bool
            d_551_v56_ = False
            d_550_v55_ = (d_550_v55_).set(d_551_v56_, ((self).f28 if d_551_v56_ else (self).f28))
            (self).m8(d_551_v56_, d_551_v56_, d_551_v56_, globalState)
            d_552_v57_: _dafny.Seq
            d_552_v57_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "npq"))
            (globalState).f13 = (d_552_v57_) < (d_552_v57_)
            d_553_v58_: str
            d_553_v58_ = _dafny.CodePoint('d')
            d_554_v59_: _dafny.Seq
            d_554_v59_ = _dafny.SeqWithoutIsStrInference([d_551_v56_])
            d_555_v60_: D13
            d_555_v60_ = D13_DC26(d_552_v57_, (D3_DC6((d_552_v57_).set(default__.safeIndex((self).f28, len(d_552_v57_)), _dafny.CodePoint('e')), d_553_v58_, d_551_v56_)).cf8, (d_554_v59_)[default__.safeIndex((self).f28, len(d_554_v59_))])
            if (d_555_v60_).cf35:
                d_556_v61_: _dafny.Seq
                d_556_v61_ = _dafny.SeqWithoutIsStrInference([(self).f28])
                d_557_v62_: _dafny.Seq
                d_557_v62_ = _dafny.SeqWithoutIsStrInference([d_556_v61_, _dafny.SeqWithoutIsStrInference([(self).f28 for d_558_i2_ in range(default__.abs(886))])])
                d_559_v63_: _dafny.Seq
                d_559_v63_ = _dafny.SeqWithoutIsStrInference([d_557_v62_, d_557_v62_])
                d_560_v64_: _dafny.Array
                nw71_ = _dafny.Array(None, 7)
                nw71_[int(0)] = (d_559_v63_)[default__.safeIndex((self).f28, len(d_559_v63_))]
                nw71_[int(1)] = d_557_v62_
                nw71_[int(2)] = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_553_v58_ for d_561_i4_ in range(default__.abs(-177))]))])).set(default__.safeIndex((self).f28, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_553_v58_ for d_562_i4_ in range(default__.abs(-177))]))]))), (self).f28) for d_563_i3_ in range(default__.abs(-179))])
                nw71_[int(3)] = d_557_v62_
                nw71_[int(4)] = d_557_v62_
                nw71_[int(5)] = d_557_v62_
                nw71_[int(6)] = (d_557_v62_ if d_551_v56_ else (default__.fm22(d_551_v56_, d_551_v56_, (d_554_v59_)[default__.safeIndex((self).f28, len(d_554_v59_))], globalState)).set(default__.safeIndex((self).f28, len(default__.fm22(d_551_v56_, d_551_v56_, (d_554_v59_)[default__.safeIndex((self).f28, len(d_554_v59_))], globalState))), d_556_v61_))
                d_560_v64_ = nw71_
                d_564_v65_: _dafny.Map
                d_564_v65_ = _dafny.Map({(0) - (len(d_554_v59_)): d_557_v62_})
                index77_ = default__.safeIndex(451, (d_560_v64_).length(0))
                (d_560_v64_)[index77_] = (d_557_v62_) + (((d_564_v65_)[(self).f28] if ((self).f28) in (d_564_v65_) else d_557_v62_))
                index78_ = default__.safeIndex(451, (d_560_v64_).length(0))
                (d_560_v64_)[index78_] = d_557_v62_
                (globalState).f13 = d_551_v56_
                index79_ = default__.safeIndex(623, (p0).length(0))
                (p0)[index79_] = default__.safeModuloInt(976, (self).f28)
                index80_ = default__.safeIndex(623, (p0).length(0))
                (p0)[index80_] = (0) - ((self).f28)
                d_565_v66_: _dafny.MultiSet
                d_565_v66_ = _dafny.MultiSet([d_553_v58_, (d_552_v57_)[default__.safeIndex(313, len(d_552_v57_))], (d_553_v58_ if d_551_v56_ else d_553_v58_), d_553_v58_])
                (globalState).f16 = ((d_565_v66_)[(d_552_v57_)[default__.safeIndex((self).f28, len(d_552_v57_))]] if ((d_552_v57_)[default__.safeIndex((self).f28, len(d_552_v57_))]) in (d_565_v66_) else default__.safeDivisionInt((p0)[default__.safeIndex(623, (p0).length(0))], (self).f28))
                (globalState).f21 = not(((self).f28) != ((self).f28))
            elif True:
                (globalState).f2 = d_551_v56_
                d_566_v67_: _dafny.Set
                d_566_v67_ = _dafny.Set({d_553_v58_, d_553_v58_})
                d_567_v68_: _dafny.Seq
                d_567_v68_ = _dafny.SeqWithoutIsStrInference([d_566_v67_])
                (globalState).f2 = (_dafny.CodePoint('w')) in ((d_567_v68_)[default__.safeIndex(924, len(d_567_v68_))])
                d_568_v69_: _dafny.Seq
                d_568_v69_ = _dafny.SeqWithoutIsStrInference([(self).f28])
                d_569_v70_: _dafny.Set
                d_569_v70_ = _dafny.Set({d_551_v56_})
                d_570_v71_: _dafny.Map
                d_570_v71_ = _dafny.Map({(self).f28: d_569_v70_})
                d_571_v72_: _dafny.Array
                nw72_ = _dafny.Array(None, 9)
                nw72_[int(0)] = d_569_v70_
                nw72_[int(1)] = _dafny.Set({d_551_v56_, d_551_v56_})
                nw72_[int(2)] = _dafny.Set({d_551_v56_})
                nw72_[int(3)] = ((d_570_v71_)[(self).f28] if ((self).f28) in (d_570_v71_) else d_569_v70_)
                nw72_[int(4)] = d_569_v70_
                nw72_[int(5)] = d_569_v70_
                nw72_[int(6)] = d_569_v70_
                nw72_[int(7)] = d_569_v70_
                nw72_[int(8)] = d_569_v70_
                d_571_v72_ = nw72_
                d_572_v73_: D7
                d_572_v73_ = D7_DC16(d_571_v72_)
                d_573_v74_: _dafny.Array
                nw73_ = _dafny.Array(None, 13)
                nw73_[int(0)] = d_572_v73_
                nw73_[int(1)] = D7_DC16(d_571_v72_)
                nw73_[int(2)] = d_572_v73_
                nw73_[int(3)] = d_572_v73_
                nw73_[int(4)] = d_572_v73_
                nw73_[int(5)] = d_572_v73_
                nw73_[int(6)] = d_572_v73_
                nw73_[int(7)] = D7_DC16(d_571_v72_)
                nw73_[int(8)] = d_572_v73_
                nw73_[int(9)] = d_572_v73_
                nw73_[int(10)] = d_572_v73_
                nw73_[int(11)] = D7_DC16(d_571_v72_)
                nw73_[int(12)] = d_572_v73_
                d_573_v74_ = nw73_
                d_574_v75_: _dafny.Seq
                d_574_v75_ = _dafny.SeqWithoutIsStrInference([d_573_v74_, d_573_v74_])
                d_575_v76_: _dafny.Map
                d_575_v76_ = _dafny.Map({d_551_v56_: d_551_v56_})
                d_576_v77_: _dafny.MultiSet
                d_576_v77_ = _dafny.MultiSet([len(d_575_v76_), (self).f28])
                d_568_v69_ = _dafny.SeqWithoutIsStrInference([(self).f28, len(d_574_v75_), ((d_576_v77_).intersection(_dafny.MultiSet([(self).f28, len(d_552_v57_)]))).cardinality, (d_568_v69_)[default__.safeIndex((self).f28, len(d_568_v69_))]])
                d_577_v78_: C0
                nw74_ = C0()
                nw74_.ctor__(d_551_v56_)
                d_577_v78_ = nw74_
                d_568_v69_ = ((_dafny.SeqWithoutIsStrInference([len(d_554_v59_), (self).f28])) + (_dafny.SeqWithoutIsStrInference([-630]))) + (d_568_v69_)
        if True:
            d_578_v79_: _dafny.Array
            nw75_ = _dafny.Array(False, 27)
            d_578_v79_ = nw75_
            d_579_v80_: bool
            d_579_v80_ = False
            index81_ = default__.safeIndex(864, (d_578_v79_).length(0))
            (d_578_v79_)[index81_] = d_579_v80_
            d_580_v81_: _dafny.Seq
            d_580_v81_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "erwiuf"))
            d_581_v82_: _dafny.Map
            d_581_v82_ = _dafny.Map({d_579_v80_: d_580_v81_})
            index82_ = default__.safeIndex(864, (d_578_v79_).length(0))
            (d_578_v79_)[index82_] = (default__.fm23(d_579_v80_, ((d_581_v82_)[True] if (True) in (d_581_v82_) else d_580_v81_), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_582_i5_ in range(default__.abs(50))]), d_579_v80_, globalState)).cf34
            d_583_v83_: _dafny.Seq
            d_583_v83_ = _dafny.SeqWithoutIsStrInference([len(d_580_v81_), (self).f28])
            d_584_v84_: _dafny.Set
            d_584_v84_ = _dafny.Set({(d_583_v83_)[default__.safeIndex((self).f28, len(d_583_v83_))], (self).f28, (0) - ((self).f28)})
            index83_ = default__.safeIndex(268, (p0).length(0))
            (p0)[index83_] = (0) - ((0) - (len(d_584_v84_)))
            index84_ = default__.safeIndex(268, (p0).length(0))
            (p0)[index84_] = ((self).f28) + ((0) - ((self).f28))
            d_585_v85_: _dafny.Map
            d_585_v85_ = _dafny.Map({(d_578_v79_)[default__.safeIndex(864, (d_578_v79_).length(0))]: d_578_v79_})
            d_586_v86_: _dafny.Map
            d_586_v86_ = _dafny.Map({d_585_v85_: -210})
            d_586_v86_ = (d_586_v86_).set((d_585_v85_) | (d_585_v85_), (p0)[default__.safeIndex(268, (p0).length(0))])
            d_587_v87_: _dafny.Map
            d_587_v87_ = _dafny.Map({d_579_v80_: (0) - ((self).f28)})
            d_588_v88_: _dafny.Map
            d_588_v88_ = _dafny.Map({(self).f28: len(d_587_v87_)})
            d_588_v88_ = (d_588_v88_).set((p0)[default__.safeIndex(268, (p0).length(0))], (p0)[default__.safeIndex(268, (p0).length(0))])
            d_589_v89_: C0
            nw76_ = C0()
            nw76_.ctor__((d_578_v79_)[default__.safeIndex(864, (d_578_v79_).length(0))])
            d_589_v89_ = nw76_
        elif True:
            d_590_v90_: bool
            d_590_v90_ = True
            d_591_v91_: _dafny.Map
            d_591_v91_ = _dafny.Map({len(_dafny.Map({d_590_v90_: (self).f28})): _dafny.CodePoint('i')})
            d_592_v92_: str
            d_592_v92_ = _dafny.CodePoint('a')
            d_593_v93_: _dafny.Map
            d_593_v93_ = _dafny.Map({d_590_v90_: d_590_v90_})
            d_594_v94_: _dafny.Map
            d_594_v94_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cdoyd"))): d_593_v93_})
            d_595_v95_: _dafny.Map
            d_595_v95_ = _dafny.Map({d_592_v92_: d_590_v90_})
            d_596_v96_: _dafny.MultiSet
            d_596_v96_ = _dafny.MultiSet([d_595_v95_])
            (globalState).f21 = not((default__.fm24(((d_591_v91_)[493] if (493) in (d_591_v91_) else d_592_v92_), len(d_594_v94_), 342, 724, globalState)).isdisjoint(d_596_v96_))
            index85_ = default__.safeIndex(273, (p0).length(0))
            (p0)[index85_] = (self).f28
            d_597_v98_: _dafny.Map
            def iife47_():
                coll31_ = _dafny.Map()
                compr_31_: int
                for compr_31_ in _dafny.IntegerRange(-121, 315):
                    d_598_v97_: int = compr_31_
                    if ((-121) <= (d_598_v97_)) and ((d_598_v97_) < (315)):
                        coll31_[(d_598_v97_) - ((self).f28)] = False
                return _dafny.Map(coll31_)
            d_597_v98_ = _dafny.Map({len(iife47_()
            ): d_590_v90_})
            index86_ = default__.safeIndex(273, (p0).length(0))
            (p0)[index86_] = default__.safeDivisionInt(len((d_597_v98_) | (d_597_v98_)), (self).f28)
            (globalState).f21 = d_590_v90_
            (globalState).f0 = ((p0)[default__.safeIndex(273, (p0).length(0))]) > ((self).f28)
            d_599_v99_: _dafny.Seq
            d_599_v99_ = _dafny.SeqWithoutIsStrInference([d_590_v90_, not(d_590_v90_)])
            d_600_v100_: _dafny.Map
            d_600_v100_ = _dafny.Map({d_590_v90_: d_599_v99_})
            d_600_v100_ = (d_600_v100_).set(False, d_599_v99_)
        d_601_v101_: _dafny.Seq
        d_601_v101_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jeovgrwpv"))
        d_602_v102_: _dafny.MultiSet
        d_602_v102_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_603_i6_ in range(default__.abs(181))]), d_601_v101_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_604_i7_ in range(default__.abs(528))])])
        d_602_v102_ = ((d_602_v102_) - (d_602_v102_)) | (d_602_v102_)
        d_605_v103_: bool
        d_605_v103_ = True
        d_606_v104_: _dafny.MultiSet
        d_606_v104_ = _dafny.MultiSet([d_605_v103_])
        d_607_v105_: _dafny.Map
        d_607_v105_ = _dafny.Map({213: not(False)})
        d_608_v106_: _dafny.Seq
        d_608_v106_ = _dafny.SeqWithoutIsStrInference([default__.fm2(_dafny.Map({d_605_v103_: d_605_v103_}), globalState)])
        d_609_v107_: _dafny.Seq
        d_609_v107_ = d_608_v106_
        hi2_ = ((self).f28) + ((self).fm17(d_605_v103_, _dafny.SeqWithoutIsStrInference([(self).f28]), len(d_608_v106_), d_609_v107_, globalState))
        for d_610_i8_ in range((((d_606_v104_)[True] if (True) in (d_606_v104_) else ((d_606_v104_)[d_605_v103_] if (d_605_v103_) in (d_606_v104_) else len(d_607_v105_)))) + (921), hi2_):
            d_611_v108_: str
            d_611_v108_ = _dafny.CodePoint('o')
            (self).m8((d_611_v108_) in (default__.fm1(d_610_i8_, globalState)), d_605_v103_, d_605_v103_, globalState)
            d_612_v109_: _dafny.Map
            d_612_v109_ = _dafny.Map({d_605_v103_: (self).f28})
            (globalState).f16 = default__.safeModuloInt(len((d_612_v109_).set(d_605_v103_, (self).f28)), default__.safeModuloInt(762, d_610_i8_))
            d_613_v110_: D6
            d_613_v110_ = D6_DC13((self).f28, d_610_i8_)
            d_614_v111_: D6
            d_614_v111_ = D6_DC14(d_613_v110_)
            source13_ = d_614_v111_
            if source13_.is_DC13:
                d_615___mcc_h4_ = source13_.cf19
                d_616___mcc_h5_ = source13_.cf20
                d_617_cf20_ = d_616___mcc_h5_
                d_618_cf19_ = d_615___mcc_h4_
                d_619_v112_: _dafny.Seq
                d_619_v112_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                d_620_v113_: D6
                d_620_v113_ = D6_DC12(d_619_v112_)
                d_621_v114_: _dafny.MultiSet
                d_621_v114_ = _dafny.MultiSet([d_617_cf20_])
                pat_let_tv11_ = d_619_v112_
                pat_let_tv12_ = d_619_v112_
                pat_let_tv13_ = p0
                def iife48_(_pat_let8_0):
                    def iife49_(d_622_dt__update__tmp_h1_):
                        def iife50_(_pat_let9_0):
                            def iife51_(d_623_dt__update_hcf18_h0_):
                                return D6_DC12(d_623_dt__update_hcf18_h0_)
                            return iife51_(_pat_let9_0)
                        return iife50_((pat_let_tv11_).set(default__.safeIndex(677, len(pat_let_tv12_)), pat_let_tv13_))
                    return iife49_(_pat_let8_0)
                rhs64_ = d_617_cf20_
                rhs65_ = d_617_cf20_
                rhs66_ = (d_605_v103_) == (d_605_v103_)
                rhs67_ = (_dafny.MultiSet([len(d_601_v101_)])).ispropersubset(d_621_v114_)
                rhs68_ = iife48_(D6_DC12(d_619_v112_))
                lhs67_ = globalState
                lhs68_ = globalState
                lhs69_ = globalState
                lhs70_ = globalState
                lhs67_.f7 = rhs64_
                lhs68_.f7 = rhs65_
                lhs69_.f21 = rhs66_
                lhs70_.f21 = rhs67_
                d_620_v113_ = rhs68_
                (globalState).f13 = not(not (d_605_v103_) or ((d_605_v103_) == (d_605_v103_)))
                (globalState).f2 = d_605_v103_
                d_611_v108_ = d_611_v108_
            elif source13_.is_DC12:
                d_624___mcc_h6_ = source13_.cf18
                d_625_cf18_ = d_624___mcc_h6_
                d_626_v115_: _dafny.Set
                d_626_v115_ = _dafny.Set({d_605_v103_})
                d_627_v116_: C0
                nw77_ = C0()
                nw77_.ctor__((_dafny.Set({d_605_v103_, not(d_605_v103_), True})).issubset(d_626_v115_))
                d_627_v116_ = nw77_
                d_608_v106_ = (_dafny.SeqWithoutIsStrInference([d_627_v116_.f27])) + (d_608_v106_)
                (self).m8((d_606_v104_).ispropersubset(d_606_v104_), d_605_v103_, d_627_v116_.f27, globalState)
                d_628_v117_: _dafny.Array
                def lambda21_(d_629_i8_, d_630_v105_, d_631_v103_):
                    def lambda22_(d_632_i9_):
                        return ((d_630_v105_)[d_629_i8_] if (d_629_i8_) in (d_630_v105_) else d_631_v103_)

                    return lambda22_

                init12_ = lambda21_(d_610_i8_, d_607_v105_, d_605_v103_)
                nw78_ = _dafny.Array(None, 26)
                for i0_12_ in range(nw78_.length(0)):
                    nw78_[i0_12_] = init12_(i0_12_)
                d_628_v117_ = nw78_
                d_633_v118_: _dafny.Map
                d_633_v118_ = _dafny.Map({(self).f28: d_608_v106_})
                d_634_v119_: _dafny.Seq
                d_634_v119_ = _dafny.SeqWithoutIsStrInference([((d_633_v118_)[(0) - ((self).f28)] if ((0) - ((self).f28)) in (d_633_v118_) else d_608_v106_)])
                index87_ = default__.safeIndex(799, (d_628_v117_).length(0))
                (d_628_v117_)[index87_] = (d_605_v103_) in ((d_634_v119_)[default__.safeIndex(d_610_i8_, len(d_634_v119_))])
                index88_ = default__.safeIndex(799, (d_628_v117_).length(0))
                (d_628_v117_)[index88_] = ((d_605_v103_) and (d_605_v103_) if d_605_v103_ else (d_610_i8_) <= ((self).f28))
            elif True:
                d_635___mcc_h7_ = source13_.cf21
                d_636_cf21_ = d_635___mcc_h7_
                (globalState).f2 = (d_605_v103_) == (d_605_v103_)
                d_637_v120_: D0
                d_637_v120_ = D0_DC0(d_605_v103_)
                d_638_v121_: _dafny.Seq
                d_638_v121_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([not((d_637_v120_).cf0)])).set(not(d_605_v103_), default__.abs(d_610_i8_))])
                d_606_v104_ = (d_638_v121_)[default__.safeIndex(d_610_i8_, len(d_638_v121_))]
                d_639_v122_: _dafny.Array
                nw79_ = _dafny.Array(int(0), 26)
                d_639_v122_ = nw79_
                d_639_v122_ = (d_639_v122_ if ((self).fm6(554, globalState)) and (d_605_v103_) else p0)
                d_640_v123_: _dafny.Array
                nw80_ = _dafny.Array(False, 17)
                d_640_v123_ = nw80_
                d_641_v124_: _dafny.Map
                d_641_v124_ = _dafny.Map({d_640_v123_: d_610_i8_})
                d_641_v124_ = (d_641_v124_).set(d_640_v123_, d_610_i8_)
            d_642_v125_: C0
            nw81_ = C0()
            nw81_.ctor__(d_605_v103_)
            d_642_v125_ = nw81_
        d_643_v126_: _dafny.Array
        nw82_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 20)
        d_643_v126_ = nw82_
        index89_ = default__.safeIndex(873, (d_643_v126_).length(0))
        (d_643_v126_)[index89_] = (d_601_v101_) + (default__.fm1((self).f28, globalState))
        index90_ = default__.safeIndex(873, (d_643_v126_).length(0))
        (d_643_v126_)[index90_] = d_601_v101_

    def m1(self, p0, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: bool = False
        d_644_v0_: _dafny.Set
        d_644_v0_ = _dafny.Set({939, p0, (self).f28})
        d_645_v1_: bool
        d_645_v1_ = True
        d_646_v2_: str
        d_646_v2_ = _dafny.CodePoint('a')
        d_647_v3_: _dafny.Seq
        d_647_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "prnrf"))
        d_648_v4_: _dafny.Seq
        d_648_v4_ = _dafny.SeqWithoutIsStrInference([not(d_645_v1_), default__.fm2(_dafny.Map({d_645_v1_: d_645_v1_}), globalState)])
        d_649_v5_: D1
        d_649_v5_ = D1_DC3(d_645_v1_)
        d_650_v6_: _dafny.Set
        d_650_v6_ = _dafny.Set({True, d_645_v1_, d_645_v1_})
        d_651_v7_: _dafny.Array
        nw83_ = _dafny.Array(None, 16)
        nw83_[int(0)] = not((len(d_644_v0_)) <= ((self).f28))
        nw83_[int(1)] = d_645_v1_
        nw83_[int(2)] = (d_646_v2_) not in (d_647_v3_)
        nw83_[int(3)] = ((d_648_v4_)[default__.safeIndex((self).f28, len(d_648_v4_))] if d_645_v1_ else (d_649_v5_).cf3)
        nw83_[int(4)] = (d_645_v1_) == (d_645_v1_)
        nw83_[int(5)] = d_645_v1_
        nw83_[int(6)] = d_645_v1_
        nw83_[int(7)] = (self).fm6(len(d_650_v6_), globalState)
        nw83_[int(8)] = d_645_v1_
        nw83_[int(9)] = not (True) or (d_645_v1_)
        nw83_[int(10)] = d_645_v1_
        nw83_[int(11)] = d_645_v1_
        nw83_[int(12)] = d_645_v1_
        nw83_[int(13)] = d_645_v1_
        nw83_[int(14)] = d_645_v1_
        nw83_[int(15)] = d_645_v1_
        d_651_v7_ = nw83_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_651_v7_).length(0)):
            d_652_i0_: int = guard_loop_5_
            if (True) and (((0) <= (d_652_i0_)) and ((d_652_i0_) < ((d_651_v7_).length(0)))):
                (d_651_v7_)[(d_652_i0_)] = d_645_v1_
        d_653_v8_: _dafny.Map
        d_653_v8_ = _dafny.Map({(self).f28: (self).f28})
        d_654_v9_: _dafny.Array
        def lambda23_(d_655_v3_):
            def lambda24_(d_656_i1_):
                return (d_655_v3_)[default__.safeIndex((self).f28, len(d_655_v3_))]

            return lambda24_

        init13_ = lambda23_(d_647_v3_)
        nw84_ = _dafny.Array(None, 15)
        for i0_13_ in range(nw84_.length(0)):
            nw84_[i0_13_] = init13_(i0_13_)
        d_654_v9_ = nw84_
        d_657_v10_: _dafny.Map
        d_657_v10_ = _dafny.Map({p0: d_654_v9_})
        d_653_v8_ = (d_653_v8_).set(p0, (len(d_657_v10_)) + ((self).f28))
        if d_645_v1_:
            d_658_v11_: _dafny.Array
            def lambda25_(d_659_v1_, d_660_p0_):
                def lambda26_(d_661_i2_):
                    return default__.safeDivisionInt(d_661_i2_, len(_dafny.Map({d_659_v1_: d_660_p0_})))

                return lambda26_

            init14_ = lambda25_(d_645_v1_, p0)
            nw85_ = _dafny.Array(None, 28)
            for i0_14_ in range(nw85_.length(0)):
                nw85_[i0_14_] = init14_(i0_14_)
            d_658_v11_ = nw85_
            index91_ = default__.safeIndex(261, (d_658_v11_).length(0))
            (d_658_v11_)[index91_] = p0
            d_662_v12_: _dafny.Seq
            d_662_v12_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_646_v2_ for d_663_i3_ in range(default__.abs(-535))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lwkghweu")), d_647_v3_])
            d_664_v13_: _dafny.Seq
            d_664_v13_ = _dafny.SeqWithoutIsStrInference([(d_662_v12_)[default__.safeIndex(509, len(d_662_v12_))], d_647_v3_, d_647_v3_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hpd"))])
            d_665_v14_: _dafny.Map
            d_665_v14_ = _dafny.Map({d_648_v4_: d_651_v7_})
            d_666_v15_: _dafny.Map
            d_666_v15_ = _dafny.Map({d_653_v8_: d_646_v2_})
            d_667_v17_: _dafny.Set
            d_667_v17_ = _dafny.Set({_dafny.Map({p0: (0) - ((_dafny.MultiSet([d_645_v1_, d_645_v1_, False, d_645_v1_, True])).cardinality)}), d_653_v8_})
            index92_ = default__.safeIndex(261, (d_658_v11_).length(0))
            def iife52_():
                coll32_ = _dafny.Set()
                compr_32_: _dafny.Map
                for compr_32_ in (d_666_v15_).keys.Elements:
                    d_668_v16_: _dafny.Map = compr_32_
                    if (d_668_v16_) in (d_666_v15_):
                        coll32_ = coll32_.union(_dafny.Set([d_668_v16_]))
                return _dafny.Set(coll32_)
            rhs69_ = default__.safeModuloInt(p0, (self).f28)
            rhs70_ = (iife52_()
            ).isdisjoint(d_667_v17_)
            rhs71_ = d_662_v12_
            rhs72_ = ((d_665_v14_) | (d_665_v14_)) | (d_665_v14_)
            lhs71_ = d_658_v11_
            lhs72_ = default__.safeIndex(261, (d_658_v11_).length(0))
            lhs73_ = globalState
            lhs71_[lhs72_] = rhs69_
            lhs73_.f21 = rhs70_
            d_662_v12_ = rhs71_
            d_665_v14_ = rhs72_
            d_645_v1_ = False
            (globalState).f13 = not(d_645_v1_)
            (globalState).f21 = d_645_v1_
            d_669_v18_: _dafny.Array
            nw86_ = _dafny.Array(None, 10)
            nw86_[int(0)] = d_647_v3_
            nw86_[int(1)] = default__.fm1(len(d_647_v3_), globalState)
            nw86_[int(2)] = d_647_v3_
            nw86_[int(3)] = (default__.fm25(d_645_v1_, (self).f28, d_645_v1_, 842, globalState)).cf6
            nw86_[int(4)] = d_647_v3_
            nw86_[int(5)] = d_647_v3_
            nw86_[int(6)] = d_647_v3_
            nw86_[int(7)] = (d_647_v3_).set(default__.safeIndex(p0, len(d_647_v3_)), d_646_v2_)
            nw86_[int(8)] = d_647_v3_
            nw86_[int(9)] = d_647_v3_
            d_669_v18_ = nw86_
            index93_ = default__.safeIndex(327, (d_669_v18_).length(0))
            (d_669_v18_)[index93_] = d_647_v3_
            index94_ = default__.safeIndex(327, (d_669_v18_).length(0))
            (d_669_v18_)[index94_] = (d_647_v3_) + (d_647_v3_)
        elif True:
            d_670_v19_: _dafny.Array
            nw87_ = _dafny.Array(int(0), 28)
            d_670_v19_ = nw87_
            d_671_v20_: _dafny.MultiSet
            d_671_v20_ = _dafny.MultiSet([d_646_v2_])
            d_672_v21_: _dafny.Map
            d_672_v21_ = _dafny.Map({d_645_v1_: (d_671_v20_).cardinality})
            d_673_v22_: _dafny.Map
            d_673_v22_ = _dafny.Map({d_672_v21_: d_645_v1_})
            d_674_v23_: _dafny.Map
            d_674_v23_ = _dafny.Map({d_645_v1_: d_670_v19_})
            d_675_v24_: _dafny.Array
            nw88_ = _dafny.Array(None, 12)
            nw88_[int(0)] = d_670_v19_
            nw88_[int(1)] = d_670_v19_
            nw88_[int(2)] = d_670_v19_
            nw88_[int(3)] = d_670_v19_
            nw88_[int(4)] = d_670_v19_
            nw88_[int(5)] = d_670_v19_
            nw88_[int(6)] = (d_670_v19_ if (self).fm7(False, p0, ((d_673_v22_)[d_672_v21_] if (d_672_v21_) in (d_673_v22_) else d_645_v1_), globalState) else d_670_v19_)
            nw88_[int(7)] = d_670_v19_
            nw88_[int(8)] = d_670_v19_
            nw88_[int(9)] = ((d_674_v23_)[d_645_v1_] if (d_645_v1_) in (d_674_v23_) else d_670_v19_)
            nw88_[int(10)] = d_670_v19_
            nw88_[int(11)] = d_670_v19_
            d_675_v24_ = nw88_
            index95_ = default__.safeIndex(141, (d_675_v24_).length(0))
            (d_675_v24_)[index95_] = d_670_v19_
            d_676_v25_: _dafny.Seq
            d_676_v25_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_648_v4_)])
            d_677_v26_: _dafny.Seq
            d_677_v26_ = _dafny.SeqWithoutIsStrInference([445])
            d_678_v27_: _dafny.Map
            d_678_v27_ = _dafny.Map({553: d_677_v26_})
            d_679_v28_: _dafny.MultiSet
            d_679_v28_ = _dafny.MultiSet([True])
            d_680_v29_: D4
            d_680_v29_ = D4_DC10(p0, (self).f28, (self).fm7(d_645_v1_, (self).f28, d_645_v1_, globalState), d_645_v1_, p0)
            index96_ = default__.safeIndex(141, (d_675_v24_).length(0))
            rhs73_ = ((d_670_v19_ if d_645_v1_ else d_670_v19_) if d_645_v1_ else d_670_v19_)
            rhs74_ = ((d_676_v25_).set(default__.safeIndex(len(d_678_v27_), len(d_676_v25_)), (_dafny.MultiSet(d_648_v4_)) - (d_679_v28_))).set(default__.safeIndex(p0, len((d_676_v25_).set(default__.safeIndex(len(d_678_v27_), len(d_676_v25_)), (_dafny.MultiSet(d_648_v4_)) - (d_679_v28_)))), (d_679_v28_).set(d_645_v1_, default__.abs((self).f28)))
            rhs75_ = (d_680_v29_).cf16
            rhs76_ = d_645_v1_
            lhs74_ = d_675_v24_
            lhs75_ = default__.safeIndex(141, (d_675_v24_).length(0))
            lhs76_ = globalState
            lhs77_ = globalState
            lhs74_[lhs75_] = rhs73_
            d_676_v25_ = rhs74_
            lhs76_.f7 = rhs75_
            lhs77_.f0 = rhs76_
            d_681_v30_: C0
            nw89_ = C0()
            nw89_.ctor__(True)
            d_681_v30_ = nw89_
            d_682_v31_: C0
            nw90_ = C0()
            nw90_.ctor__(((self).f28) >= (p0))
            d_682_v31_ = nw90_
            if d_645_v1_:
                d_683_v32_: _dafny.Map
                d_683_v32_ = _dafny.Map({(p0) == ((self).f28): (self).fm6(len(d_647_v3_), globalState)})
                d_683_v32_ = (d_683_v32_).set((d_645_v1_) or (d_682_v31_.f27), d_682_v31_.f27)
                r1 = d_682_v31_.f27
                d_684_v33_: _dafny.Array
                nw91_ = _dafny.Array(_dafny.Array(None, 0), 3)
                d_684_v33_ = nw91_
                d_685_v34_: _dafny.Map
                d_685_v34_ = _dafny.Map({p0: (self).fm6((self).f28, globalState)})
                d_686_v37_: _dafny.Array
                nw92_ = _dafny.Array(None, 7)
                nw92_[int(0)] = d_644_v0_
                nw92_[int(1)] = _dafny.Set({len(d_685_v34_)})
                nw92_[int(2)] = d_644_v0_
                def iife53_():
                    coll33_ = _dafny.Set()
                    compr_33_: int
                    for compr_33_ in _dafny.IntegerRange(630, 607):
                        d_687_v35_: int = compr_33_
                        if ((630) <= (d_687_v35_)) and ((d_687_v35_) < (607)):
                            coll33_ = coll33_.union(_dafny.Set([(d_687_v35_) - ((self).f28)]))
                    return _dafny.Set(coll33_)
                nw92_[int(3)] = iife53_()
                
                nw92_[int(4)] = d_644_v0_
                nw92_[int(5)] = d_644_v0_
                def iife54_():
                    coll34_ = _dafny.Set()
                    compr_34_: int
                    for compr_34_ in _dafny.IntegerRange(910, 974):
                        d_688_v36_: int = compr_34_
                        if ((910) <= (d_688_v36_)) and ((d_688_v36_) < (974)):
                            coll34_ = coll34_.union(_dafny.Set([(d_688_v36_) - ((self).f28)]))
                    return _dafny.Set(coll34_)
                nw92_[int(6)] = iife54_()
                
                d_686_v37_ = nw92_
                index97_ = default__.safeIndex(650, (d_684_v33_).length(0))
                (d_684_v33_)[index97_] = d_686_v37_
                d_689_v38_: _dafny.Seq
                d_689_v38_ = _dafny.SeqWithoutIsStrInference([d_685_v34_, d_685_v34_])
                index98_ = default__.safeIndex(650, (d_684_v33_).length(0))
                rhs77_ = d_651_v7_
                rhs78_ = len((d_689_v38_) + (d_689_v38_))
                rhs79_ = (d_686_v37_ if d_682_v31_.f27 else d_686_v37_)
                lhs78_ = globalState
                lhs79_ = d_684_v33_
                lhs80_ = default__.safeIndex(650, (d_684_v33_).length(0))
                d_651_v7_ = rhs77_
                lhs78_.f7 = rhs78_
                lhs79_[lhs80_] = rhs79_
                d_651_v7_ = d_651_v7_
                d_690_v39_: _dafny.Map
                d_690_v39_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_682_v31_.f27]): (self).f28})
                d_690_v39_ = (d_690_v39_).set((_dafny.SeqWithoutIsStrInference([d_682_v31_.f27])) + (d_648_v4_), (self).f28)
            elif True:
                (globalState).f2 = True
                d_677_v26_ = d_677_v26_
                (globalState).f7 = 519
                d_651_v7_ = d_651_v7_
                d_691_v40_: _dafny.Map
                d_691_v40_ = _dafny.Map({d_681_v30_.f27: d_646_v2_})
                d_692_v41_: _dafny.Map
                d_692_v41_ = _dafny.Map({d_647_v3_: len(d_691_v40_)})
                d_692_v41_ = (d_692_v41_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vvrdpbh")), p0)
            arr0_ = (d_675_v24_)[default__.safeIndex(141, (d_675_v24_).length(0))]
            index99_ = default__.safeIndex(365, ((d_675_v24_)[default__.safeIndex(141, (d_675_v24_).length(0))]).length(0))
            arr0_[index99_] = default__.safeDivisionInt((self).f28, p0)
            d_693_v42_: D6
            d_693_v42_ = D6_DC13((self).f28, (_dafny.MultiSet([d_682_v31_.f27])).cardinality)
            d_694_v43_: D6
            d_694_v43_ = D6_DC14(d_693_v42_)
            d_695_v44_: D6
            d_695_v44_ = D6_DC14(d_693_v42_)
            d_696_v45_: _dafny.Seq
            d_696_v45_ = _dafny.SeqWithoutIsStrInference([d_647_v3_])
            d_697_v46_: _dafny.Seq
            d_697_v46_ = _dafny.SeqWithoutIsStrInference([(d_696_v45_)[default__.safeIndex((self).f28, len(d_696_v45_))], d_647_v3_, d_647_v3_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ikcvsk"))])
            arr1_ = (d_675_v24_)[default__.safeIndex(141, (d_675_v24_).length(0))]
            index100_ = default__.safeIndex(365, ((d_675_v24_)[default__.safeIndex(141, (d_675_v24_).length(0))]).length(0))
            rhs80_ = (0) - ((self).f28)
            rhs81_ = d_695_v44_
            rhs82_ = default__.safeModuloInt(p0, p0)
            rhs83_ = (d_697_v46_)[default__.safeIndex((self).f28, len(d_697_v46_))]
            lhs81_ = (d_675_v24_)[default__.safeIndex(141, (d_675_v24_).length(0))]
            lhs82_ = default__.safeIndex(365, ((d_675_v24_)[default__.safeIndex(141, (d_675_v24_).length(0))]).length(0))
            lhs83_ = globalState
            lhs84_ = globalState
            lhs81_[lhs82_] = rhs80_
            d_695_v44_ = rhs81_
            lhs83_.f7 = rhs82_
            lhs84_.f20 = rhs83_
        d_653_v8_ = (d_653_v8_).set((self).f28, p0)
        hi3_ = (self).f28
        for d_698_i4_ in range(p0, hi3_):
            d_699_v47_: _dafny.Array
            nw93_ = _dafny.Array(_dafny.Map({}), 1)
            d_699_v47_ = nw93_
            d_700_v48_: _dafny.Map
            d_700_v48_ = _dafny.Map({(self).f28: d_645_v1_})
            index101_ = default__.safeIndex(20, (d_699_v47_).length(0))
            (d_699_v47_)[index101_] = (d_700_v48_) | (d_700_v48_)
            index102_ = default__.safeIndex(20, (d_699_v47_).length(0))
            (d_699_v47_)[index102_] = (d_700_v48_) | (_dafny.Map({d_698_i4_: d_645_v1_}))
            (globalState).f21 = (False) and (d_645_v1_)
            d_701_v49_: C0
            nw94_ = C0()
            nw94_.ctor__(d_645_v1_)
            d_701_v49_ = nw94_
            (globalState).f19 = (d_647_v3_) + (d_647_v3_)
        d_702_v50_: _dafny.Map
        d_702_v50_ = _dafny.Map({p0: True})
        if (d_644_v0_).ispropersubset(default__.fm26(d_646_v2_, d_702_v50_, globalState)):
            d_703_v51_: _dafny.Map
            d_703_v51_ = _dafny.Map({(d_646_v2_) not in (d_647_v3_): default__.fm27(not(d_645_v1_), d_645_v1_, globalState)})
            d_703_v51_ = (d_703_v51_).set(d_645_v1_, (_dafny.CodePoint('q') if d_645_v1_ else d_646_v2_))
            (globalState).f8 = d_647_v3_
            d_704_v52_: _dafny.MultiSet
            d_704_v52_ = _dafny.MultiSet([d_645_v1_, d_645_v1_])
            d_705_v53_: C0
            nw95_ = C0()
            nw95_.ctor__(((self).f28) == ((d_704_v52_).cardinality))
            d_705_v53_ = nw95_
            d_706_v54_: C0
            nw96_ = C0()
            nw96_.ctor__(not((self).fm6((self).f28, globalState)))
            d_706_v54_ = nw96_
            d_707_v55_: C0
            nw97_ = C0()
            nw97_.ctor__(d_705_v53_.f27)
            d_707_v55_ = nw97_
        elif True:
            if d_645_v1_:
                (globalState).f16 = (0) - (p0)
                d_708_v56_: _dafny.Array
                nw98_ = _dafny.Array(int(0), 28)
                d_708_v56_ = nw98_
                d_709_v57_: _dafny.Map
                d_709_v57_ = _dafny.Map({d_708_v56_: len((d_647_v3_ if d_645_v1_ else d_647_v3_))})
                d_709_v57_ = (d_709_v57_).set(d_708_v56_, (-586) * (p0))
                d_710_v58_: bool
                d_711_v59_: bool
                d_712_v60_: bool
                d_713_v61_: bool
                out0_: bool
                out1_: bool
                out2_: bool
                out3_: bool
                out0_, out1_, out2_, out3_ = (self).m9(globalState)
                d_710_v58_ = out0_
                d_711_v59_ = out1_
                d_712_v60_ = out2_
                d_713_v61_ = out3_
                def iife55_():
                    coll35_ = _dafny.Map()
                    compr_35_: int
                    for compr_35_ in (d_644_v0_).Elements:
                        d_714_v62_: int = compr_35_
                        if (d_714_v62_) in (d_644_v0_):
                            coll35_[(d_714_v62_) - ((self).f28)] = ((d_653_v8_)[(0) - ((self).f28)] if ((0) - ((self).f28)) in (d_653_v8_) else (self).f28)
                    return _dafny.Map(coll35_)
                d_653_v8_ = ((d_653_v8_) | (iife55_()
                )) | ((d_653_v8_ if False else _dafny.Map({493: len(d_647_v3_)})))
                (globalState).f13 = default__.fm2(_dafny.Map({d_710_v58_: d_711_v59_}), globalState)
            elif True:
                d_715_v63_: D4
                d_715_v63_ = D4_DC7((self).f28)
                d_716_v64_: _dafny.MultiSet
                d_716_v64_ = _dafny.MultiSet([not(d_645_v1_)])
                d_717_v65_: _dafny.Array
                nw99_ = _dafny.Array(None, 9)
                nw99_[int(0)] = p0
                nw99_[int(1)] = (0) - (len(_dafny.Map({(d_715_v63_).cf9: True})))
                nw99_[int(2)] = (0) - (len(_dafny.SeqWithoutIsStrInference([d_646_v2_ for d_718_i5_ in range(default__.abs(986))])))
                nw99_[int(3)] = (self).f28
                nw99_[int(4)] = p0
                nw99_[int(5)] = p0
                nw99_[int(6)] = (d_716_v64_).cardinality
                nw99_[int(7)] = p0
                nw99_[int(8)] = (self).f28
                d_717_v65_ = nw99_
                r0 = (d_717_v65_ if d_645_v1_ else d_717_v65_)
                d_719_v66_: _dafny.Map
                d_719_v66_ = _dafny.Map({d_645_v1_: d_645_v1_})
                d_719_v66_ = (d_719_v66_).set(not (d_645_v1_) or (d_645_v1_), d_645_v1_)
                (globalState).f7 = (self).f28
                d_720_v67_: C0
                nw100_ = C0()
                nw100_.ctor__(d_645_v1_)
                d_720_v67_ = nw100_
                d_721_v68_: _dafny.Seq
                d_721_v68_ = _dafny.SeqWithoutIsStrInference([977, (self).f28])
                d_722_v69_: _dafny.Seq
                d_722_v69_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "efk"))) for d_723_i8_ in range(default__.abs(989))])).set(default__.safeIndex((self).f28, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "efk"))) for d_724_i8_ in range(default__.abs(989))]))), (self).f28), _dafny.SeqWithoutIsStrInference([p0 for d_725_i9_ in range(default__.abs(396))])])
                d_726_v70_: _dafny.Map
                d_726_v70_ = _dafny.Map({d_645_v1_: (d_722_v69_)[default__.safeIndex(len(d_721_v68_), len(d_722_v69_))]})
                d_727_v71_: _dafny.Array
                nw101_ = _dafny.Array(None, 22)
                nw101_[int(0)] = d_721_v68_
                nw101_[int(1)] = d_721_v68_
                nw101_[int(2)] = (d_721_v68_ if d_645_v1_ else d_721_v68_)
                nw101_[int(3)] = (d_721_v68_) + (d_721_v68_)
                nw101_[int(4)] = d_721_v68_
                nw101_[int(5)] = ((d_721_v68_) + (d_721_v68_)).set(default__.safeIndex(p0, len((d_721_v68_) + (d_721_v68_))), ((d_653_v8_)[p0] if (p0) in (d_653_v8_) else len(d_647_v3_)))
                nw101_[int(6)] = _dafny.SeqWithoutIsStrInference([p0 for d_728_i6_ in range(default__.abs(792))])
                nw101_[int(7)] = d_721_v68_
                nw101_[int(8)] = d_721_v68_
                nw101_[int(9)] = d_721_v68_
                nw101_[int(10)] = d_721_v68_
                nw101_[int(11)] = (d_721_v68_).set(default__.safeIndex((self).f28, len(d_721_v68_)), (self).f28)
                nw101_[int(12)] = d_721_v68_
                nw101_[int(13)] = (d_721_v68_) + (d_721_v68_)
                nw101_[int(14)] = (_dafny.SeqWithoutIsStrInference([p0 for d_729_i7_ in range(default__.abs(471))])) + (d_721_v68_)
                nw101_[int(15)] = d_721_v68_
                nw101_[int(16)] = d_721_v68_
                nw101_[int(17)] = ((d_726_v70_)[default__.fm2(d_719_v66_, globalState)] if (default__.fm2(d_719_v66_, globalState)) in (d_726_v70_) else d_721_v68_)
                nw101_[int(18)] = default__.fm21(globalState)
                nw101_[int(19)] = (d_721_v68_ if d_645_v1_ else d_721_v68_)
                nw101_[int(20)] = (d_721_v68_) + (d_721_v68_)
                nw101_[int(21)] = d_721_v68_
                d_727_v71_ = nw101_
                index103_ = default__.safeIndex(405, (d_727_v71_).length(0))
                (d_727_v71_)[index103_] = (d_722_v69_)[default__.safeIndex((0) - ((self).f28), len(d_722_v69_))]
                index104_ = default__.safeIndex(405, (d_727_v71_).length(0))
                (d_727_v71_)[index104_] = d_721_v68_
            d_730_v72_: _dafny.Array
            def lambda27_(d_731_i10_):
                return default__.safeDivisionInt(d_731_i10_, (self).f28)

            init15_ = lambda27_
            nw102_ = _dafny.Array(None, 23)
            for i0_15_ in range(nw102_.length(0)):
                nw102_[i0_15_] = init15_(i0_15_)
            d_730_v72_ = nw102_
            d_732_v73_: _dafny.Map
            d_732_v73_ = _dafny.Map({(self).f28: d_730_v72_})
            r0 = ((d_732_v73_)[p0] if (p0) in (d_732_v73_) else d_730_v72_)
            d_733_v74_: bool
            d_734_v75_: bool
            d_735_v76_: bool
            d_736_v77_: bool
            out4_: bool
            out5_: bool
            out6_: bool
            out7_: bool
            out4_, out5_, out6_, out7_ = (self).m9(globalState)
            d_733_v74_ = out4_
            d_734_v75_ = out5_
            d_735_v76_ = out6_
            d_736_v77_ = out7_
            if (d_648_v4_)[default__.safeIndex(p0, len(d_648_v4_))]:
                d_646_v2_ = d_646_v2_
                d_737_v78_: C0
                nw103_ = C0()
                nw103_.ctor__(d_736_v77_)
                d_737_v78_ = nw103_
                d_648_v4_ = _dafny.SeqWithoutIsStrInference([(self).fm6(-80, globalState), d_645_v1_])
                d_738_v79_: _dafny.Array
                nw104_ = _dafny.Array(None, 2)
                nw104_[int(0)] = (d_644_v0_) - (d_644_v0_)
                nw104_[int(1)] = d_644_v0_
                d_738_v79_ = nw104_
                index105_ = default__.safeIndex(956, (d_738_v79_).length(0))
                (d_738_v79_)[index105_] = d_644_v0_
                d_739_v80_: _dafny.Array
                nw105_ = _dafny.Array(_dafny.Array(None, 0), 12)
                d_739_v80_ = nw105_
                index106_ = default__.safeIndex(124, (d_739_v80_).length(0))
                (d_739_v80_)[index106_] = d_730_v72_
                index107_ = default__.safeIndex(956, (d_738_v79_).length(0))
                index108_ = default__.safeIndex(124, (d_739_v80_).length(0))
                rhs84_ = d_644_v0_
                rhs85_ = d_730_v72_
                rhs86_ = d_735_v76_
                lhs85_ = d_738_v79_
                lhs86_ = default__.safeIndex(956, (d_738_v79_).length(0))
                lhs87_ = d_739_v80_
                lhs88_ = default__.safeIndex(124, (d_739_v80_).length(0))
                lhs89_ = globalState
                lhs85_[lhs86_] = rhs84_
                lhs87_[lhs88_] = rhs85_
                lhs89_.f2 = rhs86_
                d_740_v81_: _dafny.Map
                d_740_v81_ = _dafny.Map({d_645_v1_: p0})
                d_741_v82_: _dafny.MultiSet
                d_741_v82_ = _dafny.MultiSet([((d_740_v81_)[True] if (True) in (d_740_v81_) else p0)])
                d_742_v83_: _dafny.MultiSet
                d_742_v83_ = _dafny.MultiSet([(d_741_v82_).cardinality, (0) - (p0)])
                d_743_v84_: _dafny.Seq
                d_743_v84_ = _dafny.SeqWithoutIsStrInference([d_730_v72_])
                d_744_v85_: _dafny.Seq
                d_744_v85_ = _dafny.SeqWithoutIsStrInference([(d_743_v84_)[default__.safeIndex(p0, len(d_743_v84_))]])
                index109_ = default__.safeIndex(124, (d_739_v80_).length(0))
                rhs87_ = d_651_v7_
                rhs88_ = (default__.fm28((d_738_v79_)[default__.safeIndex(956, (d_738_v79_).length(0))], p0, d_733_v74_, True, globalState)).issubset((d_741_v82_).set(len(d_647_v3_), default__.abs(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xnvuuooh"))))))
                rhs89_ = (d_743_v84_)[default__.safeIndex((0) - ((self).f28), len(d_743_v84_))]
                lhs90_ = globalState
                lhs91_ = d_739_v80_
                lhs92_ = default__.safeIndex(124, (d_739_v80_).length(0))
                d_651_v7_ = rhs87_
                lhs90_.f2 = rhs88_
                lhs91_[lhs92_] = rhs89_
            elif True:
                (globalState).f13 = ((479) >= ((0) - ((self).f28)) if not(not(d_735_v76_)) else d_734_v75_)
                d_745_v86_: _dafny.Seq
                d_745_v86_ = _dafny.SeqWithoutIsStrInference([d_702_v50_, d_702_v50_])
                d_746_v87_: T1
                nw106_ = C1()
                nw106_.ctor__(p0, d_745_v86_)
                d_746_v87_ = nw106_
                d_747_v88_: _dafny.Set
                d_747_v88_ = _dafny.Set({d_746_v87_})
                d_748_v89_: _dafny.Seq
                d_748_v89_ = _dafny.SeqWithoutIsStrInference([p0])
                d_749_v90_: _dafny.Seq
                d_749_v90_ = ((d_648_v4_).set(default__.safeIndex(len(d_653_v8_), len(d_648_v4_)), d_733_v74_)).set(default__.safeIndex((self).f28, len((d_648_v4_).set(default__.safeIndex(len(d_653_v8_), len(d_648_v4_)), d_733_v74_))), d_734_v75_)
                d_750_v91_: _dafny.MultiSet
                d_750_v91_ = _dafny.MultiSet([(p0) < (len(d_747_v88_)), not(((self).fm17(not(d_733_v74_), d_748_v89_, p0, d_749_v90_, globalState)) >= ((_dafny.MultiSet([p0])).cardinality))])
                d_751_v92_: _dafny.Map
                d_751_v92_ = _dafny.Map({d_645_v1_: p0})
                (globalState).f16 = ((d_750_v91_)[d_645_v1_] if (d_645_v1_) in (d_750_v91_) else ((d_751_v92_)[d_733_v74_] if (d_733_v74_) in (d_751_v92_) else p0))
                d_651_v7_ = d_651_v7_
                index110_ = default__.safeIndex(526, (d_730_v72_).length(0))
                (d_730_v72_)[index110_] = default__.fm0(len(d_647_v3_), (self).f28, p0, d_735_v76_, globalState)
                index111_ = default__.safeIndex(526, (d_730_v72_).length(0))
                (d_730_v72_)[index111_] = default__.safeDivisionInt((len(d_647_v3_) if not(d_734_v75_) else len(default__.fm35(d_735_v76_, (self).f28, globalState))), len(d_647_v3_))
                (globalState).f19 = (((d_647_v3_).set(default__.safeIndex((self).f28, len(d_647_v3_)), d_646_v2_)) + (d_647_v3_)) + ((d_647_v3_) + (d_647_v3_))
            if d_736_v77_:
                d_644_v0_ = d_644_v0_
                d_752_v93_: _dafny.Seq
                d_752_v93_ = _dafny.SeqWithoutIsStrInference([d_647_v3_])
                (globalState).f20 = (_dafny.SeqWithoutIsStrInference([d_646_v2_ for d_753_i11_ in range(default__.abs(-857))])) + ((d_752_v93_)[default__.safeIndex((self).f28, len(d_752_v93_))])
                d_754_v94_: _dafny.Seq
                d_754_v94_ = _dafny.SeqWithoutIsStrInference([d_702_v50_])
                d_755_v95_: C1
                nw107_ = C1()
                nw107_.ctor__((self).f28, d_754_v94_)
                d_755_v95_ = nw107_
                d_651_v7_ = d_651_v7_
                (globalState).f0 = True
            elif True:
                d_756_v96_: _dafny.Map
                d_756_v96_ = _dafny.Map({d_736_v77_: d_735_v76_})
                d_757_v98_: _dafny.MultiSet
                d_757_v98_ = _dafny.MultiSet([d_736_v77_, default__.fm2(d_756_v96_, globalState), d_735_v76_])
                d_758_v99_: T1
                nw108_ = C1()
                def iife56_():
                    coll36_ = _dafny.Map()
                    compr_36_: int
                    for compr_36_ in (_dafny.MultiSet([(0) - (((d_757_v98_)[d_645_v1_] if (d_645_v1_) in (d_757_v98_) else p0))])).Elements:
                        d_759_v97_: int = compr_36_
                        if (d_759_v97_) in (_dafny.MultiSet([(0) - (((d_757_v98_)[d_645_v1_] if (d_645_v1_) in (d_757_v98_) else p0))])):
                            coll36_[default__.safeDivisionInt(d_759_v97_, (self).f28)] = d_645_v1_
                    return _dafny.Map(coll36_)
                nw108_.ctor__(len(d_756_v96_), _dafny.SeqWithoutIsStrInference([iife56_()
                ]))
                d_758_v99_ = nw108_
                d_760_v100_: _dafny.Map
                d_760_v100_ = _dafny.Map({not(d_735_v76_): d_758_v99_})
                d_760_v100_ = (d_760_v100_) | (d_760_v100_)
                rhs90_ = p0
                rhs91_ = (0) - (len(d_647_v3_))
                lhs93_ = globalState
                lhs94_ = globalState
                lhs93_.f16 = rhs90_
                lhs94_.f16 = rhs91_
                d_645_v1_ = d_736_v77_
                d_734_v75_ = d_645_v1_
                d_761_v101_: C0
                nw109_ = C0()
                nw109_.ctor__((d_647_v3_) < (_dafny.SeqWithoutIsStrInference([d_646_v2_ for d_762_i12_ in range(default__.abs(-673))])))
                d_761_v101_ = nw109_
        d_763_v102_: _dafny.Array
        nw110_ = _dafny.Array(int(0), 10)
        d_763_v102_ = nw110_
        r0 = d_763_v102_
        r1 = d_645_v1_
        return r0, r1

    def m8(self, p0, p1, p2, globalState):
        d_764_v0_: _dafny.Seq
        d_764_v0_ = _dafny.SeqWithoutIsStrInference([p2, p0, p1, p0])
        d_764_v0_ = (d_764_v0_) + (d_764_v0_)
        d_765_v1_: _dafny.Map
        d_765_v1_ = _dafny.Map({p0: (self).f28})
        d_766_v2_: _dafny.Seq
        d_766_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "asmgvid"))
        d_767_v3_: _dafny.Array
        nw111_ = _dafny.Array(None, 8)
        nw111_[int(0)] = (self).f28
        nw111_[int(1)] = (self).f28
        nw111_[int(2)] = ((d_765_v1_)[p0] if (p0) in (d_765_v1_) else len(d_766_v2_))
        nw111_[int(3)] = -74
        nw111_[int(4)] = (self).f28
        nw111_[int(5)] = (self).f28
        nw111_[int(6)] = (self).f28
        nw111_[int(7)] = (self).f28
        d_767_v3_ = nw111_
        d_768_v4_: _dafny.Seq
        d_768_v4_ = _dafny.SeqWithoutIsStrInference([d_767_v3_, d_767_v3_])
        d_769_v5_: D6
        d_769_v5_ = D6_DC12(d_768_v4_)
        pat_let_tv14_ = d_768_v4_
        d_770_v6_: _dafny.Map
        def iife57_(_pat_let10_0):
            def iife58_(d_771_dt__update__tmp_h0_):
                def iife59_(_pat_let11_0):
                    def iife60_(d_772_dt__update_hcf18_h0_):
                        return D6_DC12(d_772_dt__update_hcf18_h0_)
                    return iife60_(_pat_let11_0)
                return iife59_(pat_let_tv14_)
            return iife58_(_pat_let10_0)
        d_770_v6_ = _dafny.Map({iife57_(d_769_v5_): (self).f28})
        d_773_v7_: _dafny.Seq
        d_773_v7_ = _dafny.SeqWithoutIsStrInference([d_770_v6_])
        d_774_v8_: _dafny.Map
        d_774_v8_ = _dafny.Map({177: p1})
        d_775_v9_: _dafny.Seq
        d_775_v9_ = _dafny.SeqWithoutIsStrInference([-739, (self).f28, (self).f28])
        d_776_v10_: _dafny.Array
        nw112_ = _dafny.Array(None, 19)
        nw112_[int(0)] = d_770_v6_
        nw112_[int(1)] = d_770_v6_
        nw112_[int(2)] = ((d_773_v7_)[default__.safeIndex((self).f28, len(d_773_v7_))]) | (d_770_v6_)
        nw112_[int(3)] = d_770_v6_
        nw112_[int(4)] = d_770_v6_
        nw112_[int(5)] = (d_770_v6_).set(d_769_v5_, (self).f28)
        nw112_[int(6)] = (d_770_v6_) | (d_770_v6_)
        nw112_[int(7)] = ((d_773_v7_)[default__.safeIndex(len(d_774_v8_), len(d_773_v7_))]) | (d_770_v6_)
        nw112_[int(8)] = (d_770_v6_) | (d_770_v6_)
        nw112_[int(9)] = d_770_v6_
        nw112_[int(10)] = d_770_v6_
        nw112_[int(11)] = (d_770_v6_).set(d_769_v5_, (self).f28)
        nw112_[int(12)] = d_770_v6_
        nw112_[int(13)] = (d_770_v6_).set(d_769_v5_, len(d_775_v9_))
        nw112_[int(14)] = d_770_v6_
        nw112_[int(15)] = (d_770_v6_) | (_dafny.Map({d_769_v5_: len(d_764_v0_)}))
        nw112_[int(16)] = d_770_v6_
        nw112_[int(17)] = (d_770_v6_) | (d_770_v6_)
        nw112_[int(18)] = d_770_v6_
        d_776_v10_ = nw112_
        d_776_v10_ = d_776_v10_
        d_777_v11_: _dafny.Array
        nw113_ = _dafny.Array(None, 21)
        nw113_[int(0)] = p2
        nw113_[int(1)] = p1
        nw113_[int(2)] = (511) >= ((0) - ((self).f28))
        nw113_[int(3)] = p0
        nw113_[int(4)] = p0
        nw113_[int(5)] = not(True)
        nw113_[int(6)] = p0
        nw113_[int(7)] = p1
        nw113_[int(8)] = p0
        nw113_[int(9)] = p1
        nw113_[int(10)] = p1
        nw113_[int(11)] = p2
        nw113_[int(12)] = p1
        nw113_[int(13)] = p1
        nw113_[int(14)] = p2
        nw113_[int(15)] = p0
        nw113_[int(16)] = True
        nw113_[int(17)] = (not(p2) if p2 else p1)
        nw113_[int(18)] = not(not(p2))
        nw113_[int(19)] = p2
        nw113_[int(20)] = p2
        d_777_v11_ = nw113_
        guard_loop_6_: int
        for guard_loop_6_ in _dafny.IntegerRange(0, (d_777_v11_).length(0)):
            d_778_i0_: int = guard_loop_6_
            if (True) and (((0) <= (d_778_i0_)) and ((d_778_i0_) < ((d_777_v11_).length(0)))):
                (d_777_v11_)[(d_778_i0_)] = not(((self).f28) > ((self).f28))
        d_779_v12_: D12
        d_779_v12_ = D12_DC23((self).f28)
        d_780_v13_: D12
        d_780_v13_ = D12_DC24(d_779_v12_)
        d_781_v14_: D12
        d_781_v14_ = D12_DC24(d_780_v13_)
        d_782_i1_: int
        d_782_i1_ = 0
        with _dafny.label("4"):
            pat_let_tv15_ = p1
            pat_let_tv16_ = p2
            pat_let_tv17_ = p0
            def lambda30_(source14_):
                if source14_.is_DC23:
                    d_798___mcc_h0_ = source14_.cf30
                    d_799_cf30_ = d_798___mcc_h0_
                    return pat_let_tv15_
                elif source14_.is_DC22:
                    d_800___mcc_h1_ = source14_.cf29
                    d_801_cf29_ = d_800___mcc_h1_
                    return pat_let_tv16_
                elif True:
                    d_802___mcc_h2_ = source14_.cf31
                    d_803_cf31_ = d_802___mcc_h2_
                    return not(pat_let_tv17_)

            while lambda30_(d_781_v14_):
                with _dafny.c_label("4"):
                    if (d_782_i1_) >= (100):
                        raise _dafny.Break("4")
                    d_782_i1_ = (d_782_i1_) + (1)
                    (globalState).f21 = p0
                    if p1:
                        (globalState).f7 = ((self).f28) + (((self).f28) - ((self).f28))
                        index112_ = default__.safeIndex(37, (d_777_v11_).length(0))
                        (d_777_v11_)[index112_] = p0
                        d_783_v15_: D1
                        d_783_v15_ = D1_DC3(p2)
                        index113_ = default__.safeIndex(37, (d_777_v11_).length(0))
                        (d_777_v11_)[index113_] = (d_783_v15_).cf3
                        (globalState).f21 = False
                        index114_ = default__.safeIndex(317, (d_767_v3_).length(0))
                        (d_767_v3_)[index114_] = default__.safeDivisionInt((self).f28, (self).f28)
                        index115_ = default__.safeIndex(317, (d_767_v3_).length(0))
                        (d_767_v3_)[index115_] = default__.safeDivisionInt(((self).f28) * ((self).f28), ((0) - ((self).f28)) + (len(d_765_v1_)))
                        (globalState).f21 = False
                    elif True:
                        (globalState).f16 = (self).f28
                        d_784_v16_: C0
                        nw114_ = C0()
                        nw114_.ctor__(p0)
                        d_784_v16_ = nw114_
                        (globalState).f7 = default__.safeModuloInt((self).f28, (self).f28)
                        d_785_v17_: _dafny.Set
                        d_785_v17_ = _dafny.Set({(self).f28})
                        d_786_v18_: _dafny.Map
                        d_786_v18_ = _dafny.Map({d_785_v17_: _dafny.Set({p2, p0})})
                        d_787_v19_: _dafny.Set
                        d_787_v19_ = _dafny.Set({d_784_v16_.f27, False, p1, p2, p1})
                        d_786_v18_ = (d_786_v18_).set(d_785_v17_, d_787_v19_)
                        d_788_v20_: _dafny.Array
                        nw115_ = _dafny.Array(_dafny.Seq({}), 13)
                        d_788_v20_ = nw115_
                        index116_ = default__.safeIndex(965, (d_788_v20_).length(0))
                        (d_788_v20_)[index116_] = d_764_v0_
                        index117_ = default__.safeIndex(965, (d_788_v20_).length(0))
                        (d_788_v20_)[index117_] = (d_764_v0_) + (d_764_v0_)
                    d_789_v21_: _dafny.Array
                    def lambda28_(d_790_p0_, d_791_p1_):
                        def lambda29_(d_792_i2_):
                            return _dafny.Set({d_790_p0_, d_791_p1_, not(d_790_p0_), d_791_p1_, d_790_p0_})

                        return lambda29_

                    init16_ = lambda28_(p0, p1)
                    nw116_ = _dafny.Array(None, 11)
                    for i0_16_ in range(nw116_.length(0)):
                        nw116_[i0_16_] = init16_(i0_16_)
                    d_789_v21_ = nw116_
                    d_793_v22_: D7
                    d_793_v22_ = D7_DC16(d_789_v21_)
                    d_793_v22_ = D7_DC16(d_789_v21_)
                    d_794_v23_: str
                    d_794_v23_ = _dafny.CodePoint('b')
                    d_795_v24_: _dafny.Map
                    d_795_v24_ = _dafny.Map({p2: False})
                    d_796_v25_: _dafny.Set
                    d_796_v25_ = _dafny.Set({p2})
                    d_797_v26_: _dafny.Map
                    d_797_v26_ = _dafny.Map({p2: d_796_v25_})
                    rhs92_ = d_794_v23_
                    rhs93_ = (0) - ((0) - (default__.safeModuloInt((self).f28, default__.fm0(len(d_795_v24_), default__.fm0((self).f28, 350, len(((d_797_v26_)[p2] if (p2) in (d_797_v26_) else d_796_v25_)), p1, globalState), (self).f28, p0, globalState))))
                    rhs94_ = p2
                    rhs95_ = (self).f28
                    lhs95_ = globalState
                    lhs96_ = globalState
                    lhs97_ = globalState
                    d_794_v23_ = rhs92_
                    lhs95_.f7 = rhs93_
                    lhs96_.f2 = rhs94_
                    lhs97_.f16 = rhs95_
                    pass
            pass
        d_804_v27_: _dafny.Array
        def lambda31_(d_805_p2_, d_806_p0_):
            def lambda32_(d_807_i3_):
                return _dafny.MultiSet([not(d_805_p2_), d_805_p2_, d_806_p0_])

            return lambda32_

        init17_ = lambda31_(p2, p0)
        nw117_ = _dafny.Array(None, 28)
        for i0_17_ in range(nw117_.length(0)):
            nw117_[i0_17_] = init17_(i0_17_)
        d_804_v27_ = nw117_
        d_808_v28_: _dafny.MultiSet
        d_808_v28_ = _dafny.MultiSet([p2])
        index118_ = default__.safeIndex(404, (d_804_v27_).length(0))
        (d_804_v27_)[index118_] = d_808_v28_
        d_809_v29_: _dafny.Map
        d_809_v29_ = _dafny.Map({p2: p1})
        d_810_v30_: _dafny.MultiSet
        d_810_v30_ = _dafny.MultiSet([default__.fm34(globalState)])
        d_811_v31_: _dafny.Seq
        d_811_v31_ = _dafny.SeqWithoutIsStrInference([d_810_v30_])
        d_812_v32_: _dafny.Map
        d_812_v32_ = _dafny.Map({(self).f28: d_766_v2_})
        d_813_v33_: _dafny.Map
        d_813_v33_ = _dafny.Map({len(d_812_v32_): len(d_766_v2_)})
        d_814_v34_: _dafny.Set
        d_814_v34_ = _dafny.Set({(self).f28, (self).f28, (len(d_774_v8_)) + (len(d_811_v31_)), default__.fm0(len(d_813_v33_), (self).f28, (self).f28, p0, globalState)})
        index119_ = default__.safeIndex(404, (d_804_v27_).length(0))
        rhs96_ = d_808_v28_
        rhs97_ = (d_809_v29_) | (d_809_v29_)
        rhs98_ = (d_814_v34_).intersection(_dafny.Set({(self).f28, (self).f28}))
        lhs98_ = d_804_v27_
        lhs99_ = default__.safeIndex(404, (d_804_v27_).length(0))
        lhs98_[lhs99_] = rhs96_
        d_809_v29_ = rhs97_
        d_814_v34_ = rhs98_
        (globalState).f2 = (self).fm7(((self).f28) <= ((self).f28), (self).f28, p1, globalState)

    def m9(self, globalState):
        r0: bool = False
        r1: bool = False
        r2: bool = False
        r3: bool = False
        d_815_v0_: _dafny.Seq
        d_815_v0_ = _dafny.SeqWithoutIsStrInference([(self).f28])
        d_816_i0_: int
        d_816_i0_ = 0
        with _dafny.label("5"):
            while not(((d_815_v0_) + ((d_815_v0_).set(default__.safeIndex((self).f28, len(d_815_v0_)), (self).f28))) == (default__.fm37(globalState))):
                with _dafny.c_label("5"):
                    if (d_816_i0_) >= (100):
                        raise _dafny.Break("5")
                    d_816_i0_ = (d_816_i0_) + (1)
                    d_817_v1_: _dafny.Array
                    def lambda33_(d_818_i1_):
                        return _dafny.CodePoint('o')

                    init18_ = lambda33_
                    nw118_ = _dafny.Array(None, 24)
                    for i0_18_ in range(nw118_.length(0)):
                        nw118_[i0_18_] = init18_(i0_18_)
                    d_817_v1_ = nw118_
                    d_819_v2_: str
                    d_819_v2_ = _dafny.CodePoint('h')
                    index120_ = default__.safeIndex(121, (d_817_v1_).length(0))
                    (d_817_v1_)[index120_] = d_819_v2_
                    d_820_v3_: bool
                    d_820_v3_ = False
                    index121_ = default__.safeIndex(121, (d_817_v1_).length(0))
                    (d_817_v1_)[index121_] = default__.fm27(((self).f28) == ((self).f28), d_820_v3_, globalState)
                    d_821_v4_: _dafny.Map
                    d_821_v4_ = _dafny.Map({d_820_v3_: d_820_v3_})
                    d_822_v5_: _dafny.Seq
                    d_822_v5_ = _dafny.SeqWithoutIsStrInference([((d_821_v4_)[d_820_v3_] if (d_820_v3_) in (d_821_v4_) else d_820_v3_), d_820_v3_])
                    d_823_v6_: _dafny.Seq
                    d_823_v6_ = d_822_v5_
                    d_824_v7_: _dafny.Map
                    d_824_v7_ = _dafny.Map({d_820_v3_: d_822_v5_})
                    d_825_v8_: D14
                    d_825_v8_ = D14_DC29((self).fm17(d_820_v3_, _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_820_v3_, False])) for d_826_i2_ in range(default__.abs(884))]), (self).f28, d_823_v6_, globalState), d_824_v7_)
                    (globalState).f16 = (0) - ((d_825_v8_).cf38)
                    source15_ = D14_DC30(d_815_v0_, (d_822_v5_).set(default__.safeIndex((self).f28, len(d_822_v5_)), not(d_820_v3_)))
                    if source15_.is_DC29:
                        d_827___mcc_h0_ = source15_.cf38
                        d_828___mcc_h1_ = source15_.cf39
                        d_829_cf39_ = d_828___mcc_h1_
                        d_830_cf38_ = d_827___mcc_h0_
                        d_831_v9_: _dafny.MultiSet
                        d_831_v9_ = _dafny.MultiSet([not (d_820_v3_) or (d_820_v3_)])
                        d_832_v10_: _dafny.Map
                        d_832_v10_ = _dafny.Map({(self).f28: (self).f28})
                        d_833_v11_: D13
                        d_833_v11_ = D13_DC25(d_832_v10_)
                        d_834_v12_: _dafny.Set
                        d_834_v12_ = _dafny.Set({d_833_v11_})
                        d_835_v13_: _dafny.Map
                        d_835_v13_ = _dafny.Map({d_834_v12_: d_819_v2_})
                        rhs99_ = (self).f28
                        rhs100_ = 209
                        rhs101_ = ((d_831_v9_)[d_820_v3_] if (d_820_v3_) in (d_831_v9_) else len((d_835_v13_).set(d_834_v12_, d_819_v2_)))
                        rhs102_ = (0) - (default__.fm0((d_830_cf38_) + ((self).f28), -485, ((_dafny.MultiSet(d_815_v0_)).cardinality if not(d_820_v3_) else d_830_cf38_), d_820_v3_, globalState))
                        lhs100_ = globalState
                        lhs101_ = globalState
                        lhs102_ = globalState
                        lhs100_.f16 = rhs99_
                        d_830_cf38_ = rhs100_
                        lhs101_.f7 = rhs101_
                        lhs102_.f7 = rhs102_
                        d_836_v14_: _dafny.Array
                        def lambda34_(d_837_v9_):
                            def lambda35_(d_838_i3_):
                                return (d_838_i3_) * ((d_837_v9_).cardinality)

                            return lambda35_

                        init19_ = lambda34_(d_831_v9_)
                        nw119_ = _dafny.Array(None, 11)
                        for i0_19_ in range(nw119_.length(0)):
                            nw119_[i0_19_] = init19_(i0_19_)
                        d_836_v14_ = nw119_
                        d_839_v15_: _dafny.MultiSet
                        d_839_v15_ = _dafny.MultiSet([d_830_cf38_])
                        index122_ = default__.safeIndex(93, (d_836_v14_).length(0))
                        (d_836_v14_)[index122_] = (d_830_cf38_) - ((d_839_v15_).cardinality)
                        index123_ = default__.safeIndex(93, (d_836_v14_).length(0))
                        (d_836_v14_)[index123_] = default__.safeModuloInt(d_830_cf38_, d_830_cf38_)
                        d_815_v0_ = d_815_v0_
                        d_840_v16_: _dafny.Map
                        d_840_v16_ = _dafny.Map({(((d_815_v0_).set(default__.safeIndex((d_836_v14_)[default__.safeIndex(93, (d_836_v14_).length(0))], len(d_815_v0_)), d_830_cf38_)).set(default__.safeIndex((d_836_v14_)[default__.safeIndex(93, (d_836_v14_).length(0))], len((d_815_v0_).set(default__.safeIndex((d_836_v14_)[default__.safeIndex(93, (d_836_v14_).length(0))], len(d_815_v0_)), d_830_cf38_))), (d_836_v14_)[default__.safeIndex(93, (d_836_v14_).length(0))])).set(default__.safeIndex((0) - (-205), len(((d_815_v0_).set(default__.safeIndex((d_836_v14_)[default__.safeIndex(93, (d_836_v14_).length(0))], len(d_815_v0_)), d_830_cf38_)).set(default__.safeIndex((d_836_v14_)[default__.safeIndex(93, (d_836_v14_).length(0))], len((d_815_v0_).set(default__.safeIndex((d_836_v14_)[default__.safeIndex(93, (d_836_v14_).length(0))], len(d_815_v0_)), d_830_cf38_))), (d_836_v14_)[default__.safeIndex(93, (d_836_v14_).length(0))]))), (self).f28): d_822_v5_})
                        d_841_v17_: _dafny.Seq
                        d_841_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gsaqxlayd"))
                        d_842_v18_: _dafny.Seq
                        d_842_v18_ = _dafny.SeqWithoutIsStrInference([d_840_v16_, d_840_v16_, d_840_v16_, ((d_840_v16_).set(d_815_v0_, d_822_v5_)).set(d_815_v0_, d_822_v5_)])
                        d_843_v19_: _dafny.Set
                        d_843_v19_ = _dafny.Set({d_820_v3_})
                        d_844_v21_: D4
                        d_844_v21_ = D4_DC9(len(d_815_v0_))
                        d_845_v22_: _dafny.Map
                        d_845_v22_ = _dafny.Map({default__.fm21(globalState): d_844_v21_})
                        d_846_v23_: D16
                        d_846_v23_ = D16_DC35(d_840_v16_)
                        d_847_v25_: _dafny.Seq
                        d_847_v25_ = _dafny.SeqWithoutIsStrInference([d_815_v0_])
                        d_848_v28_: _dafny.MultiSet
                        d_848_v28_ = _dafny.MultiSet([d_815_v0_])
                        pat_let_tv18_ = d_840_v16_
                        d_849_v29_: _dafny.Array
                        nw120_ = _dafny.Array(None, 27)
                        nw120_[int(0)] = d_840_v16_
                        nw120_[int(1)] = d_840_v16_
                        nw120_[int(2)] = d_840_v16_
                        nw120_[int(3)] = d_840_v16_
                        nw120_[int(4)] = _dafny.Map({d_815_v0_: _dafny.SeqWithoutIsStrInference([d_820_v3_, False, d_820_v3_])})
                        nw120_[int(5)] = ((d_840_v16_).set(_dafny.SeqWithoutIsStrInference([len(d_841_v17_), 205]), d_822_v5_)) | (d_840_v16_)
                        nw120_[int(6)] = d_840_v16_
                        nw120_[int(7)] = (d_842_v18_)[default__.safeIndex(len(d_843_v19_), len(d_842_v18_))]
                        nw120_[int(8)] = d_840_v16_
                        nw120_[int(9)] = d_840_v16_
                        nw120_[int(10)] = d_840_v16_
                        nw120_[int(11)] = d_840_v16_
                        nw120_[int(12)] = d_840_v16_
                        nw120_[int(13)] = default__.fm39((d_831_v9_).cardinality, d_820_v3_, (self).f28, globalState)
                        nw120_[int(14)] = _dafny.Map({d_815_v0_: d_822_v5_})
                        nw120_[int(15)] = d_840_v16_
                        def iife61_():
                            coll37_ = _dafny.Map()
                            compr_37_: _dafny.Seq
                            for compr_37_ in (d_845_v22_).keys.Elements:
                                d_850_v20_: _dafny.Seq = compr_37_
                                if (d_850_v20_) in (d_845_v22_):
                                    coll37_[d_850_v20_] = _dafny.SeqWithoutIsStrInference([d_820_v3_])
                            return _dafny.Map(coll37_)
                        nw120_[int(16)] = (iife61_()
                        ) | (d_840_v16_)
                        nw120_[int(17)] = _dafny.Map({d_815_v0_: d_822_v5_})
                        def iife62_(_pat_let12_0):
                            def iife63_(d_851_dt__update__tmp_h0_):
                                def iife64_(_pat_let13_0):
                                    def iife65_(d_852_dt__update_hcf46_h0_):
                                        return D16_DC35(d_852_dt__update_hcf46_h0_)
                                    return iife65_(_pat_let13_0)
                                return iife64_(pat_let_tv18_)
                            return iife63_(_pat_let12_0)
                        nw120_[int(18)] = (iife62_(d_846_v23_)).cf46
                        nw120_[int(19)] = d_840_v16_
                        def iife66_():
                            coll38_ = _dafny.Map()
                            compr_38_: _dafny.Seq
                            for compr_38_ in (d_847_v25_).Elements:
                                d_853_v24_: _dafny.Seq = compr_38_
                                if (d_853_v24_) in (d_847_v25_):
                                    coll38_[d_853_v24_] = d_822_v5_
                            return _dafny.Map(coll38_)
                        nw120_[int(20)] = (iife66_()
                        ) | (d_840_v16_)
                        nw120_[int(21)] = d_840_v16_
                        nw120_[int(22)] = d_840_v16_
                        nw120_[int(23)] = d_840_v16_
                        nw120_[int(24)] = d_840_v16_
                        def iife67_():
                            coll39_ = _dafny.Map()
                            compr_39_: _dafny.Seq
                            for compr_39_ in (d_847_v25_).Elements:
                                d_854_v26_: _dafny.Seq = compr_39_
                                if (d_854_v26_) in (d_847_v25_):
                                    coll39_[d_854_v26_] = d_822_v5_
                            return _dafny.Map(coll39_)
                        nw120_[int(25)] = (iife67_()
                        ).set(d_815_v0_, d_822_v5_)
                        def iife68_():
                            coll40_ = _dafny.Map()
                            compr_40_: _dafny.Seq
                            for compr_40_ in (d_848_v28_).Elements:
                                d_855_v27_: _dafny.Seq = compr_40_
                                if (d_855_v27_) in (d_848_v28_):
                                    coll40_[d_855_v27_] = _dafny.SeqWithoutIsStrInference([d_820_v3_, False])
                            return _dafny.Map(coll40_)
                        nw120_[int(26)] = (iife68_()
                        ) | (d_840_v16_)
                        d_849_v29_ = nw120_
                        index124_ = default__.safeIndex(867, (d_849_v29_).length(0))
                        (d_849_v29_)[index124_] = d_840_v16_
                        d_856_v30_: C0
                        nw121_ = C0()
                        nw121_.ctor__(False)
                        d_856_v30_ = nw121_
                        index125_ = default__.safeIndex(867, (d_849_v29_).length(0))
                        index126_ = default__.safeIndex(93, (d_836_v14_).length(0))
                        rhs103_ = (d_846_v23_).cf46
                        rhs104_ = d_819_v2_
                        rhs105_ = (d_836_v14_)[default__.safeIndex(93, (d_836_v14_).length(0))]
                        rhs106_ = ((self).f28) != (default__.safeDivisionInt((self).f28, len(d_841_v17_)))
                        rhs107_ = d_856_v30_
                        lhs103_ = d_849_v29_
                        lhs104_ = default__.safeIndex(867, (d_849_v29_).length(0))
                        lhs105_ = d_836_v14_
                        lhs106_ = default__.safeIndex(93, (d_836_v14_).length(0))
                        lhs103_[lhs104_] = rhs103_
                        d_819_v2_ = rhs104_
                        lhs105_[lhs106_] = rhs105_
                        r1 = rhs106_
                        d_856_v30_ = rhs107_
                    elif source15_.is_DC30:
                        d_857___mcc_h2_ = source15_.cf40
                        d_858___mcc_h3_ = source15_.cf41
                        d_859_cf41_ = d_858___mcc_h3_
                        d_860_cf40_ = d_857___mcc_h2_
                        d_861_v31_: _dafny.Map
                        d_861_v31_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(d_859_cf41_)[default__.safeIndex((0) - ((self).f28), len(d_859_cf41_))], (self).fm6((self).f28, globalState)]): _dafny.SeqWithoutIsStrInference([(d_817_v1_)[default__.safeIndex(121, (d_817_v1_).length(0))] for d_862_i4_ in range(default__.abs(375))])})
                        d_863_v32_: _dafny.Seq
                        d_863_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bfsvcv"))
                        d_864_v33_: _dafny.Map
                        d_864_v33_ = _dafny.Map({(self).f28: (self).f28})
                        default__.m0(not (d_820_v3_) or (d_820_v3_), -369, (d_861_v31_) | ((d_861_v31_).set(d_859_cf41_, d_863_v32_)), (811) - (len(d_864_v33_)), globalState)
                        (globalState).f7 = (0) - ((self).f28)
                        (globalState).f16 = (self).f28
                        d_865_v34_: _dafny.Array
                        nw122_ = _dafny.Array(False, 2)
                        d_865_v34_ = nw122_
                        d_866_v35_: _dafny.Array
                        nw123_ = _dafny.Array(None, 18)
                        nw123_[int(0)] = d_865_v34_
                        nw123_[int(1)] = (d_865_v34_ if d_820_v3_ else d_865_v34_)
                        nw123_[int(2)] = d_865_v34_
                        nw123_[int(3)] = d_865_v34_
                        nw123_[int(4)] = d_865_v34_
                        nw123_[int(5)] = d_865_v34_
                        nw123_[int(6)] = d_865_v34_
                        nw123_[int(7)] = d_865_v34_
                        nw123_[int(8)] = d_865_v34_
                        nw123_[int(9)] = d_865_v34_
                        nw123_[int(10)] = d_865_v34_
                        nw123_[int(11)] = d_865_v34_
                        nw123_[int(12)] = (d_865_v34_ if d_820_v3_ else d_865_v34_)
                        nw123_[int(13)] = d_865_v34_
                        nw123_[int(14)] = d_865_v34_
                        nw123_[int(15)] = d_865_v34_
                        nw123_[int(16)] = d_865_v34_
                        nw123_[int(17)] = d_865_v34_
                        d_866_v35_ = nw123_
                        index127_ = default__.safeIndex(926, (d_866_v35_).length(0))
                        (d_866_v35_)[index127_] = d_865_v34_
                        index128_ = default__.safeIndex(926, (d_866_v35_).length(0))
                        (d_866_v35_)[index128_] = d_865_v34_
                    elif source15_.is_DC28:
                        d_867___mcc_h4_ = source15_.cf37
                        d_868_cf37_ = d_867___mcc_h4_
                        (globalState).f7 = ((self).f28) * (default__.safeModuloInt((self).f28, 300))
                        d_869_v36_: _dafny.Set
                        d_869_v36_ = _dafny.Set({d_820_v3_})
                        d_870_v37_: _dafny.Array
                        nw124_ = _dafny.Array(False, 2)
                        d_870_v37_ = nw124_
                        d_871_v38_: _dafny.Map
                        d_871_v38_ = _dafny.Map({d_870_v37_: d_869_v36_})
                        d_872_v39_: _dafny.Array
                        nw125_ = _dafny.Array(None, 16)
                        nw125_[int(0)] = (d_869_v36_) - (d_869_v36_)
                        nw125_[int(1)] = d_869_v36_
                        nw125_[int(2)] = d_869_v36_
                        nw125_[int(3)] = (d_869_v36_) - (d_869_v36_)
                        nw125_[int(4)] = d_869_v36_
                        nw125_[int(5)] = _dafny.Set({((d_821_v4_)[d_820_v3_] if (d_820_v3_) in (d_821_v4_) else d_820_v3_)})
                        nw125_[int(6)] = default__.fm19(365, d_820_v3_, d_820_v3_, (self).f28, globalState)
                        nw125_[int(7)] = (d_869_v36_) - (d_869_v36_)
                        nw125_[int(8)] = (d_869_v36_ if d_820_v3_ else d_869_v36_)
                        nw125_[int(9)] = d_869_v36_
                        nw125_[int(10)] = _dafny.Set({d_820_v3_, (self).fm7(d_820_v3_, (self).f28, d_820_v3_, globalState)})
                        nw125_[int(11)] = d_869_v36_
                        nw125_[int(12)] = (d_869_v36_) | (d_869_v36_)
                        nw125_[int(13)] = d_869_v36_
                        nw125_[int(14)] = (d_869_v36_) - (d_869_v36_)
                        nw125_[int(15)] = ((d_871_v38_)[d_870_v37_] if (d_870_v37_) in (d_871_v38_) else d_869_v36_)
                        d_872_v39_ = nw125_
                        index129_ = default__.safeIndex(400, (d_872_v39_).length(0))
                        (d_872_v39_)[index129_] = (d_869_v36_).intersection(d_869_v36_)
                        index130_ = default__.safeIndex(400, (d_872_v39_).length(0))
                        (d_872_v39_)[index130_] = d_869_v36_
                        d_873_v40_: _dafny.Seq
                        d_873_v40_ = _dafny.SeqWithoutIsStrInference([d_815_v0_])
                        d_874_v41_: _dafny.Array
                        nw126_ = _dafny.Array(None, 4)
                        nw126_[int(0)] = (d_815_v0_) + (_dafny.SeqWithoutIsStrInference([len(d_868_cf37_), (self).f28, (self).f28]))
                        nw126_[int(1)] = (d_815_v0_) + (d_815_v0_)
                        nw126_[int(2)] = d_815_v0_
                        nw126_[int(3)] = (d_873_v40_)[default__.safeIndex((self).f28, len(d_873_v40_))]
                        d_874_v41_ = nw126_
                        index131_ = default__.safeIndex(791, (d_874_v41_).length(0))
                        (d_874_v41_)[index131_] = (default__.fm21(globalState)) + (d_815_v0_)
                        index132_ = default__.safeIndex(791, (d_874_v41_).length(0))
                        (d_874_v41_)[index132_] = d_815_v0_
                        d_875_v42_: C0
                        nw127_ = C0()
                        nw127_.ctor__((78) == ((self).f28))
                        d_875_v42_ = nw127_
                        d_875_v42_ = d_875_v42_
                    elif True:
                        d_876___mcc_h5_ = source15_.cf42
                        d_877_cf42_ = d_876___mcc_h5_
                        d_878_v43_: D15
                        d_878_v43_ = D15_DC33(d_815_v0_)
                        d_879_v44_: D7
                        d_879_v44_ = D7_DC15((d_878_v43_).cf44)
                        d_880_v45_: D7
                        d_880_v45_ = D7_DC17(d_879_v44_)
                        d_880_v45_ = d_880_v45_
                        d_881_v46_: _dafny.Map
                        d_881_v46_ = _dafny.Map({d_820_v3_: (0) - ((self).f28)})
                        rhs108_ = (d_815_v0_)[default__.safeIndex((self).f28, len(d_815_v0_))]
                        rhs109_ = ((d_881_v46_)[True] if (True) in (d_881_v46_) else (self).f28)
                        lhs107_ = globalState
                        lhs108_ = globalState
                        lhs107_.f7 = rhs108_
                        lhs108_.f16 = rhs109_
                        d_882_v47_: _dafny.Map
                        d_882_v47_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f28]): d_822_v5_})
                        d_883_v48_: D16
                        d_883_v48_ = D16_DC35(d_882_v47_)
                        d_884_v49_: D16
                        d_884_v49_ = D16_DC37(d_883_v48_)
                        d_884_v49_ = D16_DC37(d_883_v48_)
                        d_885_v50_: D12
                        d_885_v50_ = D12_DC22(d_881_v46_)
                        d_886_v51_: _dafny.Set
                        d_886_v51_ = _dafny.Set({d_885_v50_})
                        d_887_v52_: _dafny.Map
                        d_887_v52_ = _dafny.Map({d_885_v50_: d_820_v3_})
                        def iife69_():
                            coll41_ = _dafny.Set()
                            compr_41_: D12
                            for compr_41_ in (d_887_v52_).keys.Elements:
                                d_888_v53_: D12 = compr_41_
                                if (d_888_v53_) in (d_887_v52_):
                                    coll41_ = coll41_.union(_dafny.Set([d_888_v53_]))
                            return _dafny.Set(coll41_)
                        d_886_v51_ = ((_dafny.Set({D12_DC22(d_881_v46_)})) | (d_886_v51_)) - ((iife69_()
                        ) - (d_886_v51_))
                    d_889_v55_: D16
                    def iife70_():
                        coll42_ = _dafny.Map()
                        compr_42_: _dafny.Seq
                        for compr_42_ in (_dafny.Map({d_815_v0_: d_820_v3_})).keys.Elements:
                            d_890_v54_: _dafny.Seq = compr_42_
                            if (d_890_v54_) in (_dafny.Map({d_815_v0_: d_820_v3_})):
                                coll42_[d_890_v54_] = d_822_v5_
                        return _dafny.Map(coll42_)
                    d_889_v55_ = D16_DC37(D16_DC35(iife70_()
))
                    d_891_v56_: _dafny.MultiSet
                    d_891_v56_ = _dafny.MultiSet([d_889_v55_, d_889_v55_, d_889_v55_])
                    d_892_v57_: _dafny.Map
                    d_892_v57_ = _dafny.Map({d_820_v3_: (d_891_v56_) | (d_891_v56_)})
                    d_893_v58_: _dafny.Array
                    def lambda36_(d_894_v3_):
                        def lambda37_(d_895_i5_):
                            return d_894_v3_

                        return lambda37_

                    init20_ = lambda36_(d_820_v3_)
                    nw128_ = _dafny.Array(None, 8)
                    for i0_20_ in range(nw128_.length(0)):
                        nw128_[i0_20_] = init20_(i0_20_)
                    d_893_v58_ = nw128_
                    index133_ = default__.safeIndex(985, (d_893_v58_).length(0))
                    (d_893_v58_)[index133_] = not(not(not(d_820_v3_)))
                    d_896_v59_: _dafny.Seq
                    d_896_v59_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ji"))
                    d_897_v60_: _dafny.Map
                    d_897_v60_ = _dafny.Map({d_820_v3_: (self).f28})
                    d_898_v61_: _dafny.Array
                    nw129_ = _dafny.Array(None, 10)
                    nw129_[int(0)] = d_896_v59_
                    nw129_[int(1)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_899_i6_ in range(default__.abs(947))])
                    nw129_[int(2)] = d_896_v59_
                    nw129_[int(3)] = (d_896_v59_) + (d_896_v59_)
                    nw129_[int(4)] = _dafny.SeqWithoutIsStrInference([d_819_v2_ for d_900_i7_ in range(default__.abs(949))])
                    nw129_[int(5)] = d_896_v59_
                    nw129_[int(6)] = (d_896_v59_) + (_dafny.SeqWithoutIsStrInference([d_819_v2_ for d_901_i8_ in range(default__.abs(31))]))
                    nw129_[int(7)] = ((d_896_v59_).set(default__.safeIndex(len(d_897_v60_), len(d_896_v59_)), (d_817_v1_)[default__.safeIndex(121, (d_817_v1_).length(0))])).set(default__.safeIndex(len(d_822_v5_), len((d_896_v59_).set(default__.safeIndex(len(d_897_v60_), len(d_896_v59_)), (d_817_v1_)[default__.safeIndex(121, (d_817_v1_).length(0))]))), _dafny.CodePoint('d'))
                    nw129_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xry"))
                    nw129_[int(9)] = _dafny.SeqWithoutIsStrInference([d_819_v2_ for d_902_i9_ in range(default__.abs(49))])
                    d_898_v61_ = nw129_
                    index134_ = default__.safeIndex(629, (d_898_v61_).length(0))
                    (d_898_v61_)[index134_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_903_i10_ in range(default__.abs(284))])
                    d_904_v62_: D13
                    d_904_v62_ = D13_DC26(d_896_v59_, d_820_v3_, d_820_v3_)
                    pat_let_tv19_ = d_820_v3_
                    index135_ = default__.safeIndex(121, (d_817_v1_).length(0))
                    index136_ = default__.safeIndex(985, (d_893_v58_).length(0))
                    index137_ = default__.safeIndex(629, (d_898_v61_).length(0))
                    def iife71_(_pat_let14_0):
                        def iife72_(d_905_dt__update__tmp_h1_):
                            def iife73_(_pat_let15_0):
                                def iife74_(d_906_dt__update_hcf34_h0_):
                                    return D13_DC26((d_905_dt__update__tmp_h1_).cf33, d_906_dt__update_hcf34_h0_, (d_905_dt__update__tmp_h1_).cf35)
                                return iife74_(_pat_let15_0)
                            return iife73_(pat_let_tv19_)
                        return iife72_(_pat_let14_0)
                    rhs110_ = _dafny.Map({(iife71_(d_904_v62_)).cf34: d_891_v56_})
                    rhs111_ = ((d_817_v1_)[default__.safeIndex(121, (d_817_v1_).length(0))] if (d_820_v3_) and (d_820_v3_) else (d_817_v1_)[default__.safeIndex(121, (d_817_v1_).length(0))])
                    rhs112_ = (default__.fm2(d_821_v4_, globalState)) and ((d_904_v62_).cf35)
                    rhs113_ = (d_896_v59_).set(default__.safeIndex((self).f28, len(d_896_v59_)), (d_817_v1_)[default__.safeIndex(121, (d_817_v1_).length(0))])
                    lhs109_ = d_817_v1_
                    lhs110_ = default__.safeIndex(121, (d_817_v1_).length(0))
                    lhs111_ = d_893_v58_
                    lhs112_ = default__.safeIndex(985, (d_893_v58_).length(0))
                    lhs113_ = d_898_v61_
                    lhs114_ = default__.safeIndex(629, (d_898_v61_).length(0))
                    d_892_v57_ = rhs110_
                    lhs109_[lhs110_] = rhs111_
                    lhs111_[lhs112_] = rhs112_
                    lhs113_[lhs114_] = rhs113_
                    pass
            pass
        d_907_v63_: bool
        d_907_v63_ = True
        d_908_v64_: _dafny.Seq
        d_908_v64_ = _dafny.SeqWithoutIsStrInference([d_907_v63_, d_907_v63_])
        d_909_v65_: _dafny.Map
        d_909_v65_ = _dafny.Map({d_815_v0_: d_908_v64_})
        d_910_v66_: D16
        d_910_v66_ = D16_DC35(d_909_v65_)
        d_911_v67_: _dafny.Array
        nw130_ = _dafny.Array(_dafny.Array(None, 0), 28)
        d_911_v67_ = nw130_
        d_912_v68_: _dafny.MultiSet
        d_912_v68_ = _dafny.MultiSet([d_907_v63_, False])
        d_913_v69_: D16
        d_913_v69_ = D16_DC36(True, d_912_v68_, d_907_v63_)
        pat_let_tv20_ = d_907_v63_
        pat_let_tv21_ = d_907_v63_
        d_914_v70_: _dafny.Array
        nw131_ = _dafny.Array(None, 15)
        nw131_[int(0)] = d_913_v69_
        nw131_[int(1)] = d_913_v69_
        nw131_[int(2)] = d_913_v69_
        nw131_[int(3)] = d_913_v69_
        nw131_[int(4)] = d_913_v69_
        nw131_[int(5)] = D16_DC36(d_907_v63_, d_912_v68_, d_907_v63_)
        nw131_[int(6)] = d_913_v69_
        nw131_[int(7)] = d_913_v69_
        nw131_[int(8)] = default__.fm40(globalState)
        nw131_[int(9)] = d_913_v69_
        def iife75_(_pat_let16_0):
            def iife76_(d_915_dt__update__tmp_h2_):
                def iife77_(_pat_let17_0):
                    def iife78_(d_916_dt__update_hcf49_h0_):
                        return D16_DC36((d_915_dt__update__tmp_h2_).cf47, (d_915_dt__update__tmp_h2_).cf48, d_916_dt__update_hcf49_h0_)
                    return iife78_(_pat_let17_0)
                return iife77_(pat_let_tv20_)
            return iife76_(_pat_let16_0)
        nw131_[int(10)] = iife75_(d_913_v69_)
        def iife79_(_pat_let18_0):
            def iife80_(d_917_dt__update__tmp_h3_):
                def iife81_(_pat_let19_0):
                    def iife82_(d_918_dt__update_hcf47_h0_):
                        return D16_DC36(d_918_dt__update_hcf47_h0_, (d_917_dt__update__tmp_h3_).cf48, (d_917_dt__update__tmp_h3_).cf49)
                    return iife82_(_pat_let19_0)
                return iife81_(pat_let_tv21_)
            return iife80_(_pat_let18_0)
        nw131_[int(11)] = iife79_(d_913_v69_)
        nw131_[int(12)] = d_913_v69_
        nw131_[int(13)] = default__.fm40(globalState)
        nw131_[int(14)] = d_913_v69_
        d_914_v70_ = nw131_
        index138_ = default__.safeIndex(28, (d_911_v67_).length(0))
        (d_911_v67_)[index138_] = d_914_v70_
        d_919_v71_: D17
        d_919_v71_ = D17_DC38(d_914_v70_)
        index139_ = default__.safeIndex(28, (d_911_v67_).length(0))
        rhs114_ = d_910_v66_
        rhs115_ = (self).f28
        rhs116_ = (d_919_v71_).cf51
        lhs115_ = globalState
        lhs116_ = d_911_v67_
        lhs117_ = default__.safeIndex(28, (d_911_v67_).length(0))
        d_910_v66_ = rhs114_
        lhs115_.f7 = rhs115_
        lhs116_[lhs117_] = rhs116_
        d_920_v72_: str
        d_920_v72_ = _dafny.CodePoint('p')
        d_920_v72_ = d_920_v72_
        (globalState).f16 = (self).f28
        (globalState).f7 = len(_dafny.Set({d_907_v63_, (d_907_v63_ if d_907_v63_ else d_907_v63_), d_907_v63_, d_907_v63_}))
        d_921_v73_: D15
        d_921_v73_ = D15_DC33((_dafny.SeqWithoutIsStrInference([(self).f28])).set(default__.safeIndex((self).f28, len(_dafny.SeqWithoutIsStrInference([(self).f28]))), (self).f28))
        d_922_v74_: D15
        d_922_v74_ = D15_DC34(d_921_v73_)
        d_923_v75_: _dafny.Map
        d_923_v75_ = _dafny.Map({d_922_v74_: default__.fm0(-164, (self).f28, (self).f28, True, globalState)})
        r1 = not (d_907_v63_) or ((len(d_923_v75_)) != ((self).f28))
        r0 = d_907_v63_
        r1 = not(d_907_v63_)
        d_924_v76_: D4
        d_924_v76_ = D4_DC9((self).f28)
        r2 = (self).fm6((d_924_v76_).cf11, globalState)
        d_925_v77_: _dafny.Map
        d_925_v77_ = _dafny.Map({(self).f28: d_907_v63_})
        d_926_v78_: _dafny.Seq
        d_926_v78_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f28: d_907_v63_}), d_925_v77_])
        d_927_v79_: T1
        nw132_ = C1()
        nw132_.ctor__((self).f28, d_926_v78_)
        d_927_v79_ = nw132_
        d_928_v80_: _dafny.Seq
        d_928_v80_ = _dafny.SeqWithoutIsStrInference([d_927_v79_])
        r3 = not(not((len((d_928_v80_) + (d_928_v80_))) < (default__.safeModuloInt((0) - (-779), (self).f28))))
        return r0, r1, r2, r3

    @property
    def f28(self):
        return self._f28

class C3(T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    def ctor__(self):
        pass
        pass

    def fm5(self, p0, p1, p2, p3, globalState):
        return _dafny.Map({(523) == (len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False])]))): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_929_i0_ in range(default__.abs(992))])})

    def fm6(self, p0, globalState):
        return False

    def m1(self, p0, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: bool = False
        d_930_v0_: _dafny.Map
        d_930_v0_ = _dafny.Map({p0: p0})
        d_931_v1_: _dafny.Map
        d_931_v1_ = _dafny.Map({len(d_930_v0_): p0})
        d_932_v2_: _dafny.MultiSet
        d_932_v2_ = _dafny.MultiSet([d_930_v0_, d_930_v0_])
        d_933_v4_: bool
        d_933_v4_ = False
        d_934_v5_: _dafny.MultiSet
        d_934_v5_ = _dafny.MultiSet([True, d_933_v4_])
        d_935_v6_: _dafny.Map
        def iife83_():
            coll43_ = _dafny.Set()
            compr_43_: _dafny.Map
            for compr_43_ in (d_932_v2_).Elements:
                d_936_v3_: _dafny.Map = compr_43_
                if (d_936_v3_) in (d_932_v2_):
                    coll43_ = coll43_.union(_dafny.Set([d_936_v3_]))
            return _dafny.Set(coll43_)
        d_935_v6_ = _dafny.Map({(0) - ((len(iife83_()
        )) * (p0)): (d_934_v5_) | (d_934_v5_)})
        d_937_v7_: _dafny.Seq
        d_937_v7_ = _dafny.SeqWithoutIsStrInference([d_933_v4_])
        d_938_v8_: _dafny.Seq
        d_938_v8_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_937_v7_)])
        d_935_v6_ = (d_935_v6_).set(p0, (d_938_v8_)[default__.safeIndex(p0, len(d_938_v8_))])
        if d_933_v4_:
            d_939_v9_: _dafny.Array
            def lambda38_(d_940_p0_):
                def lambda39_(d_941_i0_):
                    return default__.safeModuloInt(d_941_i0_, d_940_p0_)

                return lambda39_

            init21_ = lambda38_(p0)
            nw133_ = _dafny.Array(None, 25)
            for i0_21_ in range(nw133_.length(0)):
                nw133_[i0_21_] = init21_(i0_21_)
            d_939_v9_ = nw133_
            d_942_v10_: _dafny.Seq
            d_942_v10_ = _dafny.SeqWithoutIsStrInference([d_939_v9_, d_939_v9_])
            d_943_v11_: _dafny.Map
            d_943_v11_ = _dafny.Map({d_933_v4_: d_933_v4_})
            d_944_v12_: D4
            d_944_v12_ = D4_DC10(p0, len(d_943_v11_), d_933_v4_, d_933_v4_, p0)
            d_945_v13_: _dafny.Seq
            d_945_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eofajvq"))
            d_946_v14_: D6
            d_946_v14_ = D6_DC12(d_942_v10_)
            d_947_v15_: _dafny.Seq
            d_947_v15_ = _dafny.SeqWithoutIsStrInference([(d_946_v14_).cf18, d_942_v10_, (d_942_v10_) + (d_942_v10_)])
            rhs117_ = (d_944_v12_).cf14
            rhs118_ = not(((len(d_945_v13_)) * (len(d_943_v11_))) <= (p0))
            rhs119_ = (d_947_v15_)[default__.safeIndex(p0, len(d_947_v15_))]
            lhs118_ = globalState
            r1 = rhs117_
            lhs118_.f0 = rhs118_
            d_942_v10_ = rhs119_
            (globalState).f16 = p0
            d_948_v16_: _dafny.Array
            nw134_ = _dafny.Array(False, 14)
            d_948_v16_ = nw134_
            d_949_v17_: _dafny.Set
            d_949_v17_ = _dafny.Set({d_948_v16_, d_948_v16_, d_948_v16_})
            rhs120_ = p0
            rhs121_ = (p0 if d_933_v4_ else (0) - ((420) * (p0)))
            rhs122_ = d_937_v7_
            rhs123_ = d_949_v17_
            rhs124_ = d_939_v9_
            lhs119_ = globalState
            lhs120_ = globalState
            lhs119_.f16 = rhs120_
            lhs120_.f7 = rhs121_
            d_937_v7_ = rhs122_
            d_949_v17_ = rhs123_
            r0 = rhs124_
            d_950_v18_: C0
            nw135_ = C0()
            nw135_.ctor__(d_933_v4_)
            d_950_v18_ = nw135_
            pat_let_tv22_ = d_942_v10_
            pat_let_tv23_ = d_942_v10_
            pat_let_tv24_ = p0
            pat_let_tv25_ = d_942_v10_
            pat_let_tv26_ = d_942_v10_
            pat_let_tv27_ = d_939_v9_
            def iife84_(_pat_let20_0):
                def iife85_(d_951_dt__update__tmp_h0_):
                    def iife86_(_pat_let21_0):
                        def iife87_(d_952_dt__update_hcf18_h0_):
                            return D6_DC12(d_952_dt__update_hcf18_h0_)
                        return iife87_(_pat_let21_0)
                    return iife86_(((pat_let_tv22_) + (pat_let_tv23_)).set(default__.safeIndex(pat_let_tv24_, len((pat_let_tv25_) + (pat_let_tv26_))), pat_let_tv27_))
                return iife85_(_pat_let20_0)
            d_946_v14_ = iife84_(d_946_v14_)
        elif True:
            (globalState).f7 = 639
            d_953_v19_: _dafny.Map
            d_953_v19_ = _dafny.Map({d_933_v4_: (p0) >= (default__.fm0(p0, p0, p0, d_933_v4_, globalState))})
            d_953_v19_ = (d_953_v19_).set(d_933_v4_, d_933_v4_)
            (globalState).f7 = (0) - (p0)
            (globalState).f21 = (p0) > (p0)
            d_954_v20_: _dafny.Set
            d_954_v20_ = _dafny.Set({d_933_v4_, not (d_933_v4_) or (False), d_933_v4_, (d_933_v4_ if d_933_v4_ else d_933_v4_), d_933_v4_})
            d_954_v20_ = _dafny.Set({d_933_v4_})
        (globalState).f21 = d_933_v4_
        d_955_v21_: D0
        d_955_v21_ = D0_DC0(d_933_v4_)
        pat_let_tv28_ = d_933_v4_
        pat_let_tv29_ = d_933_v4_
        def lambda40_(source16_):
            if source16_.is_DC0:
                d_956___mcc_h0_ = source16_.cf0
                d_957_cf0_ = d_956___mcc_h0_
                if pat_let_tv28_:
                    return d_957_cf0_
                elif True:
                    return pat_let_tv29_
            elif True:
                d_958___mcc_h1_ = source16_.cf1
                d_959_cf1_ = d_958___mcc_h1_
                return True

        if lambda40_(d_955_v21_):
            d_960_v22_: str
            d_960_v22_ = _dafny.CodePoint('e')
            d_960_v22_ = d_960_v22_
            d_961_v23_: C0
            nw136_ = C0()
            nw136_.ctor__(d_933_v4_)
            d_961_v23_ = nw136_
            (globalState).f2 = (True) == (d_961_v23_.f27)
            d_962_v24_: C0
            nw137_ = C0()
            nw137_.ctor__((d_961_v23_.f27) == (d_961_v23_.f27))
            d_962_v24_ = nw137_
            d_963_v25_: _dafny.Array
            def lambda41_(d_964_v23_, d_965_p0_):
                def lambda42_(d_966_i1_):
                    return default__.safeModuloInt(d_966_i1_, len(_dafny.Map({d_964_v23_.f27: D4_DC9(d_965_p0_)})))

                return lambda42_

            init22_ = lambda41_(d_961_v23_, p0)
            nw138_ = _dafny.Array(None, 9)
            for i0_22_ in range(nw138_.length(0)):
                nw138_[i0_22_] = init22_(i0_22_)
            d_963_v25_ = nw138_
            d_967_v26_: _dafny.Seq
            d_967_v26_ = _dafny.SeqWithoutIsStrInference([p0, p0])
            index140_ = default__.safeIndex(128, (d_963_v25_).length(0))
            (d_963_v25_)[index140_] = (len(d_967_v26_)) - (p0)
            d_968_v27_: _dafny.Set
            d_968_v27_ = _dafny.Set({d_933_v4_})
            index141_ = default__.safeIndex(128, (d_963_v25_).length(0))
            rhs125_ = (d_968_v27_).issubset((d_968_v27_) - (d_968_v27_))
            rhs126_ = default__.fm0((p0) - (p0), (_dafny.MultiSet([True])).cardinality, p0, d_962_v24_.f27, globalState)
            rhs127_ = d_961_v23_.f27
            lhs121_ = globalState
            lhs122_ = d_963_v25_
            lhs123_ = default__.safeIndex(128, (d_963_v25_).length(0))
            lhs124_ = globalState
            lhs121_.f21 = rhs125_
            lhs122_[lhs123_] = rhs126_
            lhs124_.f2 = rhs127_
        elif True:
            d_969_v28_: _dafny.Seq
            d_969_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tsxxruaee"))
            (globalState).f19 = (d_969_v28_) + ((d_969_v28_) + (d_969_v28_))
            d_970_v29_: _dafny.MultiSet
            d_970_v29_ = _dafny.MultiSet([p0, p0, p0, p0, p0])
            d_971_v30_: _dafny.Seq
            d_971_v30_ = _dafny.SeqWithoutIsStrInference([p0, p0])
            d_972_v31_: D7
            d_972_v31_ = D7_DC15(d_971_v30_)
            d_973_v32_: D7
            d_973_v32_ = D7_DC15((d_972_v31_).cf22)
            d_974_v33_: _dafny.Array
            nw139_ = _dafny.Array(None, 22)
            nw139_[int(0)] = d_970_v29_
            nw139_[int(1)] = d_970_v29_
            nw139_[int(2)] = d_970_v29_
            nw139_[int(3)] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p0]))
            nw139_[int(4)] = d_970_v29_
            nw139_[int(5)] = _dafny.MultiSet([p0, 790])
            nw139_[int(6)] = d_970_v29_
            nw139_[int(7)] = d_970_v29_
            nw139_[int(8)] = _dafny.MultiSet([p0, (0) - ((0) - (p0)), p0, len(d_969_v28_)])
            nw139_[int(9)] = (_dafny.MultiSet((d_972_v31_).cf22)) | (d_970_v29_)
            nw139_[int(10)] = (d_970_v29_ if d_933_v4_ else d_970_v29_)
            nw139_[int(11)] = d_970_v29_
            nw139_[int(12)] = (d_970_v29_) | (d_970_v29_)
            nw139_[int(13)] = (d_970_v29_) - (d_970_v29_)
            nw139_[int(14)] = d_970_v29_
            nw139_[int(15)] = (d_970_v29_).intersection(d_970_v29_)
            nw139_[int(16)] = _dafny.MultiSet([p0])
            nw139_[int(17)] = d_970_v29_
            nw139_[int(18)] = d_970_v29_
            nw139_[int(19)] = ((d_970_v29_).set(p0, default__.abs(len(d_969_v28_)))) - (d_970_v29_)
            nw139_[int(20)] = d_970_v29_
            nw139_[int(21)] = (default__.fm12(d_969_v28_, 453, d_933_v4_, globalState)) | (_dafny.MultiSet(d_971_v30_))
            d_974_v33_ = nw139_
            index142_ = default__.safeIndex(891, (d_974_v33_).length(0))
            (d_974_v33_)[index142_] = (d_970_v29_).intersection(d_970_v29_)
            index143_ = default__.safeIndex(891, (d_974_v33_).length(0))
            (d_974_v33_)[index143_] = _dafny.MultiSet(d_971_v30_)
            (globalState).f2 = d_933_v4_
            d_975_v34_: _dafny.Array
            def lambda43_(d_976_p0_):
                def lambda44_(d_977_i2_):
                    return default__.safeModuloInt(d_977_i2_, d_976_p0_)

                return lambda44_

            init23_ = lambda43_(p0)
            nw140_ = _dafny.Array(None, 19)
            for i0_23_ in range(nw140_.length(0)):
                nw140_[i0_23_] = init23_(i0_23_)
            d_975_v34_ = nw140_
            index144_ = default__.safeIndex(410, (d_975_v34_).length(0))
            (d_975_v34_)[index144_] = p0
            index145_ = default__.safeIndex(410, (d_975_v34_).length(0))
            (d_975_v34_)[index145_] = 62
            (globalState).f13 = d_933_v4_
        d_978_v35_: str
        d_978_v35_ = _dafny.CodePoint('b')
        d_979_v36_: _dafny.MultiSet
        d_979_v36_ = _dafny.MultiSet([d_978_v35_, _dafny.CodePoint('c'), d_978_v35_, _dafny.CodePoint('v')])
        d_980_v37_: _dafny.Seq
        d_980_v37_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xysyc"))
        d_981_v38_: D4
        d_981_v38_ = D4_DC10(p0, 30, not(d_933_v4_), d_933_v4_, (0) - (p0))
        d_982_v39_: _dafny.Seq
        d_982_v39_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_980_v37_), d_979_v36_])
        d_983_v40_: _dafny.Array
        nw141_ = _dafny.Array(None, 12)
        nw141_[int(0)] = d_979_v36_
        nw141_[int(1)] = d_979_v36_
        nw141_[int(2)] = d_979_v36_
        nw141_[int(3)] = d_979_v36_
        nw141_[int(4)] = d_979_v36_
        nw141_[int(5)] = default__.fm13(len(d_980_v37_), not(d_933_v4_), d_981_v38_, d_933_v4_, globalState)
        nw141_[int(6)] = d_979_v36_
        nw141_[int(7)] = d_979_v36_
        nw141_[int(8)] = (d_979_v36_) | ((d_982_v39_)[default__.safeIndex(p0, len(d_982_v39_))])
        nw141_[int(9)] = (d_979_v36_).intersection(d_979_v36_)
        nw141_[int(10)] = (d_979_v36_ if d_933_v4_ else d_979_v36_)
        nw141_[int(11)] = d_979_v36_
        d_983_v40_ = nw141_
        guard_loop_7_: int
        for guard_loop_7_ in _dafny.IntegerRange(0, (d_983_v40_).length(0)):
            d_984_i3_: int = guard_loop_7_
            if (True) and (((0) <= (d_984_i3_)) and ((d_984_i3_) < ((d_983_v40_).length(0)))):
                (d_983_v40_)[(d_984_i3_)] = _dafny.MultiSet([d_978_v35_])
        d_985_v41_: _dafny.Array
        nw142_ = _dafny.Array(_dafny.Set({}), 2)
        d_985_v41_ = nw142_
        d_986_v42_: D7
        d_986_v42_ = D7_DC16(d_985_v41_)
        d_987_v43_: D7
        d_987_v43_ = D7_DC17(d_986_v42_)
        d_988_v44_: D7
        d_988_v44_ = D7_DC17(d_987_v43_)
        source17_ = d_988_v44_
        if source17_.is_DC16:
            d_989___mcc_h2_ = source17_.cf23
            d_990_cf23_ = d_989___mcc_h2_
            d_991_v45_: _dafny.Map
            d_991_v45_ = _dafny.Map({d_980_v37_: p0})
            d_991_v45_ = (d_991_v45_).set(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jyay"))).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jyay")))), d_978_v35_)) + (d_980_v37_), p0)
            d_992_v46_: D6
            d_992_v46_ = D6_DC13(p0, p0)
            (globalState).f21 = (d_937_v7_)[default__.safeIndex((d_992_v46_).cf20, len(d_937_v7_))]
            d_993_v47_: _dafny.Map
            d_993_v47_ = _dafny.Map({d_933_v4_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eaxlkubrh"))})
            d_994_v48_: _dafny.Map
            d_994_v48_ = _dafny.Map({((((d_993_v47_)[d_933_v4_] if (d_933_v4_) in (d_993_v47_) else d_980_v37_)).set(default__.safeIndex(825, len(((d_993_v47_)[d_933_v4_] if (d_933_v4_) in (d_993_v47_) else d_980_v37_))), d_978_v35_)) + (d_980_v37_): (d_937_v7_)[default__.safeIndex(default__.fm0(-147, p0, p0, d_933_v4_, globalState), len(d_937_v7_))]})
            d_994_v48_ = (d_994_v48_).set(d_980_v37_, (p0) >= (p0))
            if d_933_v4_:
                (globalState).f2 = d_933_v4_
                d_995_v49_: _dafny.Array
                nw143_ = _dafny.Array(_dafny.Array(None, 0), 10)
                d_995_v49_ = nw143_
                d_996_v50_: _dafny.Array
                def lambda45_(d_997_p0_):
                    def lambda46_(d_998_i4_):
                        return (d_998_i4_) - (d_997_p0_)

                    return lambda46_

                init24_ = lambda45_(p0)
                nw144_ = _dafny.Array(None, 12)
                for i0_24_ in range(nw144_.length(0)):
                    nw144_[i0_24_] = init24_(i0_24_)
                d_996_v50_ = nw144_
                index146_ = default__.safeIndex(964, (d_995_v49_).length(0))
                (d_995_v49_)[index146_] = d_996_v50_
                index147_ = default__.safeIndex(964, (d_995_v49_).length(0))
                (d_995_v49_)[index147_] = d_996_v50_
                (globalState).f21 = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nxdew"))) + (d_980_v37_)) < (d_980_v37_)
                d_999_v51_: _dafny.Set
                d_999_v51_ = _dafny.Set({d_933_v4_, d_933_v4_, d_933_v4_})
                d_1000_v52_: _dafny.Set
                d_1000_v52_ = _dafny.Set({(d_999_v51_) != (default__.fm14((0) - (p0), d_933_v4_, globalState)), d_933_v4_, d_933_v4_, (p0) >= (580)})
                d_1000_v52_ = (_dafny.Set({d_933_v4_, d_933_v4_, d_933_v4_})) | ((_dafny.Set({not(d_933_v4_), d_933_v4_, d_933_v4_})).intersection(_dafny.Set({d_933_v4_})))
                d_1001_v53_: _dafny.MultiSet
                d_1001_v53_ = _dafny.MultiSet([(_dafny.MultiSet([p0, -127])).cardinality])
                d_1002_v54_: _dafny.Seq
                d_1002_v54_ = _dafny.SeqWithoutIsStrInference([d_996_v50_])
                d_1003_v55_: _dafny.Set
                d_1003_v55_ = _dafny.Set({p0, 318})
                d_1004_v56_: _dafny.Map
                d_1004_v56_ = _dafny.Map({d_1001_v53_: (_dafny.Set({p0, len(d_1002_v54_), default__.fm0((0) - (p0), p0, p0, d_933_v4_, globalState), p0, p0})).isdisjoint(d_1003_v55_)})
                d_1005_v57_: _dafny.Seq
                d_1005_v57_ = d_937_v7_
                d_1006_v58_: C0
                nw145_ = C0()
                nw145_.ctor__(d_933_v4_)
                d_1006_v58_ = nw145_
                d_1007_v59_: _dafny.Map
                d_1007_v59_ = _dafny.Map({d_1006_v58_: d_1001_v53_})
                d_1004_v56_ = (d_1004_v56_).set(_dafny.MultiSet([len((d_1005_v57_)), p0, len(d_1007_v59_)]), (self).fm6(len(d_980_v37_), globalState))
            elif True:
                d_1008_v60_: _dafny.Array
                def lambda47_(d_1009_p0_):
                    def lambda48_(d_1010_i5_):
                        return default__.safeModuloInt(d_1010_i5_, (D4_DC9(d_1009_p0_)).cf11)

                    return lambda48_

                init25_ = lambda47_(p0)
                nw146_ = _dafny.Array(None, 14)
                for i0_25_ in range(nw146_.length(0)):
                    nw146_[i0_25_] = init25_(i0_25_)
                d_1008_v60_ = nw146_
                index148_ = default__.safeIndex(556, (d_1008_v60_).length(0))
                (d_1008_v60_)[index148_] = p0
                index149_ = default__.safeIndex(556, (d_1008_v60_).length(0))
                (d_1008_v60_)[index149_] = p0
                d_1011_v61_: _dafny.Seq
                d_1011_v61_ = _dafny.SeqWithoutIsStrInference([(d_981_v38_).cf16])
                d_978_v35_ = default__.fm15((d_1011_v61_) + (d_1011_v61_), p0, globalState)
                d_1012_v62_: _dafny.Map
                d_1012_v62_ = _dafny.Map({d_933_v4_: d_1008_v60_})
                d_1013_v63_: _dafny.Map
                d_1013_v63_ = _dafny.Map({d_978_v35_: p0})
                d_1014_v64_: _dafny.MultiSet
                d_1014_v64_ = _dafny.MultiSet([372, ((d_1013_v63_)[d_978_v35_] if (d_978_v35_) in (d_1013_v63_) else (d_1008_v60_)[default__.safeIndex(556, (d_1008_v60_).length(0))])])
                d_1015_v65_: _dafny.Map
                d_1015_v65_ = _dafny.Map({d_933_v4_: d_1014_v64_})
                rhs128_ = (default__.fm0(p0, len(d_1015_v65_), (d_1008_v60_)[default__.safeIndex(556, (d_1008_v60_).length(0))], d_933_v4_, globalState)) >= (default__.safeModuloInt(((d_934_v5_)).cardinality, len(d_980_v37_)))
                rhs129_ = d_934_v5_
                rhs130_ = (d_1012_v62_) | (d_1012_v62_)
                lhs125_ = globalState
                lhs125_.f21 = rhs128_
                d_934_v5_ = rhs129_
                d_1012_v62_ = rhs130_
                (globalState).f7 = p0
                (globalState).f0 = d_933_v4_
        elif source17_.is_DC15:
            d_1016___mcc_h3_ = source17_.cf22
            d_1017_cf22_ = d_1016___mcc_h3_
            d_1018_v66_: _dafny.MultiSet
            d_1018_v66_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_978_v35_ for d_1019_i6_ in range(default__.abs(945))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tbsqpray")), d_980_v37_])
            (globalState).f7 = (d_1018_v66_).cardinality
            (globalState).f0 = (d_937_v7_)[default__.safeIndex(p0, len(d_937_v7_))]
            (globalState).f7 = p0
            d_1020_v67_: _dafny.Array
            nw147_ = _dafny.Array(int(0), 16)
            d_1020_v67_ = nw147_
            index150_ = default__.safeIndex(517, (d_1020_v67_).length(0))
            (d_1020_v67_)[index150_] = p0
            d_1021_v68_: _dafny.Array
            nw148_ = _dafny.Array(_dafny.Map({}), 15)
            d_1021_v68_ = nw148_
            d_1022_v69_: _dafny.Map
            d_1022_v69_ = _dafny.Map({not(d_933_v4_): d_1020_v67_})
            index151_ = default__.safeIndex(636, (d_1021_v68_).length(0))
            (d_1021_v68_)[index151_] = d_1022_v69_
            d_1023_v70_: _dafny.Seq
            d_1023_v70_ = _dafny.SeqWithoutIsStrInference([d_1022_v69_])
            index152_ = default__.safeIndex(517, (d_1020_v67_).length(0))
            index153_ = default__.safeIndex(636, (d_1021_v68_).length(0))
            rhs131_ = (len(default__.fm14(p0, d_933_v4_, globalState))) + ((0) - (p0))
            rhs132_ = d_1017_cf22_
            rhs133_ = (d_1023_v70_)[default__.safeIndex(default__.fm0(p0, p0, p0, d_933_v4_, globalState), len(d_1023_v70_))]
            lhs126_ = d_1020_v67_
            lhs127_ = default__.safeIndex(517, (d_1020_v67_).length(0))
            lhs128_ = d_1021_v68_
            lhs129_ = default__.safeIndex(636, (d_1021_v68_).length(0))
            lhs126_[lhs127_] = rhs131_
            d_1017_cf22_ = rhs132_
            lhs128_[lhs129_] = rhs133_
        elif True:
            d_1024___mcc_h4_ = source17_.cf24
            d_1025_cf24_ = d_1024___mcc_h4_
            if d_933_v4_:
                (globalState).f21 = (d_980_v37_) <= (d_980_v37_)
                d_1026_v71_: _dafny.Map
                d_1026_v71_ = _dafny.Map({d_933_v4_: d_933_v4_})
                d_1027_v72_: _dafny.Set
                d_1027_v72_ = _dafny.Set({d_1026_v71_})
                d_1028_v73_: C0
                nw149_ = C0()
                nw149_.ctor__(d_933_v4_)
                d_1028_v73_ = nw149_
                d_1029_v74_: _dafny.Array
                nw150_ = _dafny.Array(None, 24)
                nw150_[int(0)] = d_1028_v73_
                nw150_[int(1)] = d_1028_v73_
                nw150_[int(2)] = d_1028_v73_
                nw150_[int(3)] = d_1028_v73_
                nw150_[int(4)] = d_1028_v73_
                nw150_[int(5)] = d_1028_v73_
                nw150_[int(6)] = d_1028_v73_
                nw150_[int(7)] = d_1028_v73_
                nw150_[int(8)] = d_1028_v73_
                nw150_[int(9)] = d_1028_v73_
                nw150_[int(10)] = d_1028_v73_
                nw150_[int(11)] = d_1028_v73_
                nw150_[int(12)] = d_1028_v73_
                nw150_[int(13)] = d_1028_v73_
                nw150_[int(14)] = d_1028_v73_
                nw150_[int(15)] = d_1028_v73_
                nw150_[int(16)] = d_1028_v73_
                nw150_[int(17)] = d_1028_v73_
                nw150_[int(18)] = d_1028_v73_
                nw150_[int(19)] = d_1028_v73_
                nw150_[int(20)] = d_1028_v73_
                nw150_[int(21)] = d_1028_v73_
                nw150_[int(22)] = d_1028_v73_
                nw150_[int(23)] = d_1028_v73_
                d_1029_v74_ = nw150_
                d_1030_v75_: _dafny.Seq
                d_1030_v75_ = _dafny.SeqWithoutIsStrInference([d_1029_v74_, d_1029_v74_])
                rhs134_ = (d_1027_v72_).intersection((d_1027_v72_).intersection(d_1027_v72_))
                rhs135_ = d_980_v37_
                rhs136_ = d_933_v4_
                rhs137_ = len(((_dafny.SeqWithoutIsStrInference([d_1029_v74_, d_1029_v74_])) + (d_1030_v75_)) + (d_1030_v75_))
                lhs130_ = globalState
                lhs131_ = globalState
                lhs132_ = globalState
                d_1027_v72_ = rhs134_
                lhs130_.f8 = rhs135_
                lhs131_.f2 = rhs136_
                lhs132_.f7 = rhs137_
                (globalState).f13 = d_1028_v73_.f27
                d_1031_v76_: _dafny.Array
                def lambda49_(d_1032_v73_):
                    def lambda50_(d_1033_i7_):
                        return d_1032_v73_.f27

                    return lambda50_

                init26_ = lambda49_(d_1028_v73_)
                nw151_ = _dafny.Array(None, 19)
                for i0_26_ in range(nw151_.length(0)):
                    nw151_[i0_26_] = init26_(i0_26_)
                d_1031_v76_ = nw151_
                index154_ = default__.safeIndex(594, (d_1031_v76_).length(0))
                (d_1031_v76_)[index154_] = default__.fm2(d_1026_v71_, globalState)
                d_1034_v77_: _dafny.Set
                d_1034_v77_ = _dafny.Set({p0})
                d_1035_v78_: _dafny.MultiSet
                d_1035_v78_ = _dafny.MultiSet([d_1034_v77_, (d_1034_v77_) | (_dafny.Set({p0})), (_dafny.Set({p0, p0})) | (d_1034_v77_), d_1034_v77_])
                index155_ = default__.safeIndex(594, (d_1031_v76_).length(0))
                rhs138_ = 899
                rhs139_ = (d_980_v37_) < (d_980_v37_)
                rhs140_ = d_1035_v78_
                lhs133_ = globalState
                lhs134_ = d_1031_v76_
                lhs135_ = default__.safeIndex(594, (d_1031_v76_).length(0))
                lhs133_.f16 = rhs138_
                lhs134_[lhs135_] = rhs139_
                d_1035_v78_ = rhs140_
                d_1036_v79_: C0
                nw152_ = C0()
                nw152_.ctor__(d_933_v4_)
                d_1036_v79_ = nw152_
            elif True:
                d_981_v38_ = d_981_v38_
                (globalState).f13 = d_933_v4_
                (globalState).f8 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))) + (d_980_v37_)
                d_978_v35_ = d_978_v35_
                d_1037_v80_: _dafny.Seq
                d_1037_v80_ = _dafny.SeqWithoutIsStrInference([d_937_v7_, d_937_v7_])
                (globalState).f7 = default__.fm0(len(d_1037_v80_), (p0) - (len(_dafny.SeqWithoutIsStrInference([p0 for d_1038_i8_ in range(default__.abs(94))]))), ((d_934_v5_)[False] if (False) in (d_934_v5_) else p0), not((d_933_v4_) or (d_933_v4_)), globalState)
            d_1039_v81_: C0
            nw153_ = C0()
            nw153_.ctor__((d_933_v4_) or (d_933_v4_))
            d_1039_v81_ = nw153_
            d_1040_v82_: _dafny.Array
            def lambda51_(d_1041_p0_):
                def lambda52_(d_1042_i9_):
                    return (d_1042_i9_) - (d_1041_p0_)

                return lambda52_

            init27_ = lambda51_(p0)
            nw154_ = _dafny.Array(None, 6)
            for i0_27_ in range(nw154_.length(0)):
                nw154_[i0_27_] = init27_(i0_27_)
            d_1040_v82_ = nw154_
            d_1043_v83_: _dafny.Set
            d_1043_v83_ = _dafny.Set({d_1040_v82_})
            d_1044_v84_: _dafny.Array
            nw155_ = _dafny.Array(False, 5)
            d_1044_v84_ = nw155_
            rhs141_ = d_1043_v83_
            rhs142_ = d_978_v35_
            rhs143_ = d_1044_v84_
            d_1043_v83_ = rhs141_
            d_978_v35_ = rhs142_
            d_1044_v84_ = rhs143_
            d_1045_v85_: _dafny.Array
            nw156_ = _dafny.Array(_dafny.MultiSet({}), 25)
            d_1045_v85_ = nw156_
            index156_ = default__.safeIndex(480, (d_1045_v85_).length(0))
            (d_1045_v85_)[index156_] = (d_934_v5_).intersection(d_934_v5_)
            index157_ = default__.safeIndex(480, (d_1045_v85_).length(0))
            (d_1045_v85_)[index157_] = d_934_v5_
        d_1046_v86_: _dafny.Array
        nw157_ = _dafny.Array(int(0), 4)
        d_1046_v86_ = nw157_
        r0 = (d_1046_v86_ if d_933_v4_ else d_1046_v86_)
        r1 = d_933_v4_
        return r0, r1

    def m7(self, p0, p1, p2, p3, globalState):
        r0: D1 = D1.default()()
        r1: int = int(0)
        d_1047_v0_: _dafny.Array
        nw158_ = _dafny.Array(int(0), 28)
        d_1047_v0_ = nw158_
        d_1048_v1_: _dafny.Array
        nw159_ = _dafny.Array(_dafny.Array(None, 0), 18)
        d_1048_v1_ = nw159_
        d_1049_v2_: _dafny.Map
        d_1049_v2_ = _dafny.Map({d_1047_v0_: d_1048_v1_})
        d_1049_v2_ = (d_1049_v2_).set(d_1047_v0_, d_1048_v1_)
        hi4_ = p3
        for d_1050_i0_ in range(896, hi4_):
            index158_ = default__.safeIndex(593, (p0).length(0))
            (p0)[index158_] = _dafny.CodePoint('k')
            d_1051_v3_: str
            d_1051_v3_ = _dafny.CodePoint('o')
            index159_ = default__.safeIndex(593, (p0).length(0))
            (p0)[index159_] = d_1051_v3_
            d_1052_v4_: bool
            d_1052_v4_ = False
            if d_1052_v4_:
                (globalState).f2 = d_1052_v4_
                (globalState).f16 = -201
                d_1053_v5_: _dafny.MultiSet
                d_1053_v5_ = _dafny.MultiSet([463, 279, p2])
                d_1053_v5_ = d_1053_v5_
                d_1054_v6_: C0
                nw160_ = C0()
                nw160_.ctor__((d_1052_v4_ if not(d_1052_v4_) else d_1052_v4_))
                d_1054_v6_ = nw160_
                index160_ = default__.safeIndex(593, (p0).length(0))
                (p0)[index160_] = _dafny.CodePoint('n')
            elif True:
                d_1055_v7_: _dafny.Array
                def lambda53_(d_1056_p3_, d_1057_p1_, d_1058_v4_, d_1059_p2_):
                    def lambda54_(d_1060_i1_):
                        return not (False) or ((D4_DC10(d_1056_p3_, d_1057_p1_, not(d_1058_v4_), d_1058_v4_, (0) - (d_1059_p2_))).cf15)

                    return lambda54_

                init28_ = lambda53_(p3, p1, d_1052_v4_, p2)
                nw161_ = _dafny.Array(None, 29)
                for i0_28_ in range(nw161_.length(0)):
                    nw161_[i0_28_] = init28_(i0_28_)
                d_1055_v7_ = nw161_
                index161_ = default__.safeIndex(390, (d_1055_v7_).length(0))
                (d_1055_v7_)[index161_] = (d_1050_i0_) < (d_1050_i0_)
                index162_ = default__.safeIndex(390, (d_1055_v7_).length(0))
                rhs144_ = d_1052_v4_
                rhs145_ = (D0_DC0(d_1052_v4_)).cf0
                rhs146_ = d_1052_v4_
                lhs136_ = globalState
                lhs137_ = d_1055_v7_
                lhs138_ = default__.safeIndex(390, (d_1055_v7_).length(0))
                lhs139_ = globalState
                lhs136_.f2 = rhs144_
                lhs137_[lhs138_] = rhs145_
                lhs139_.f0 = rhs146_
                d_1061_v8_: _dafny.MultiSet
                d_1061_v8_ = _dafny.MultiSet([len(_dafny.Map({(d_1055_v7_)[default__.safeIndex(390, (d_1055_v7_).length(0))]: p3})), 879, len(_dafny.Set({False, d_1052_v4_})), p2, p3])
                d_1062_v9_: _dafny.Seq
                d_1062_v9_ = _dafny.SeqWithoutIsStrInference([d_1061_v8_, d_1061_v8_])
                d_1063_v10_: _dafny.Map
                d_1063_v10_ = _dafny.Map({(d_1062_v9_)[default__.safeIndex(p1, len(d_1062_v9_))]: d_1052_v4_})
                d_1064_v11_: _dafny.Map
                d_1064_v11_ = _dafny.Map({d_1052_v4_: d_1052_v4_})
                d_1065_v12_: _dafny.Map
                d_1065_v12_ = _dafny.Map({((d_1064_v11_)[d_1052_v4_] if (d_1052_v4_) in (d_1064_v11_) else (d_1055_v7_)[default__.safeIndex(390, (d_1055_v7_).length(0))]): d_1052_v4_})
                d_1066_v13_: C0
                nw162_ = C0()
                nw162_.ctor__((((d_1063_v10_)[_dafny.MultiSet([p2, p3])] if (_dafny.MultiSet([p2, p3])) in (d_1063_v10_) else d_1052_v4_)) or (default__.fm2(d_1064_v11_, globalState)))
                d_1066_v13_ = nw162_
                d_1067_v14_: _dafny.Array
                nw163_ = _dafny.Array(D4.default()(), 9)
                d_1067_v14_ = nw163_
                d_1068_v15_: _dafny.MultiSet
                d_1068_v15_ = _dafny.MultiSet([d_1066_v13_])
                d_1069_v16_: D4
                d_1069_v16_ = D4_DC7((d_1068_v15_).cardinality)
                index163_ = default__.safeIndex(568, (d_1067_v14_).length(0))
                (d_1067_v14_)[index163_] = d_1069_v16_
                index164_ = default__.safeIndex(568, (d_1067_v14_).length(0))
                (d_1067_v14_)[index164_] = default__.fm16(p1, d_1052_v4_, globalState)
                (globalState).f7 = p2
                index165_ = default__.safeIndex(390, (d_1055_v7_).length(0))
                (d_1055_v7_)[index165_] = (d_1055_v7_)[default__.safeIndex(390, (d_1055_v7_).length(0))]
            (globalState).f0 = d_1052_v4_
            d_1070_v17_: _dafny.Set
            d_1070_v17_ = _dafny.Set({p1, p2})
            if (d_1070_v17_).isdisjoint((d_1070_v17_) | (d_1070_v17_)):
                d_1071_v18_: _dafny.Seq
                d_1071_v18_ = _dafny.SeqWithoutIsStrInference([d_1050_i0_])
                r1 = (p1) * ((len(d_1071_v18_)) + (p2))
                d_1072_v19_: _dafny.Seq
                d_1072_v19_ = _dafny.SeqWithoutIsStrInference([True])
                (globalState).f16 = len(d_1072_v19_)
                d_1073_v20_: _dafny.MultiSet
                d_1073_v20_ = _dafny.MultiSet([default__.fm15(_dafny.SeqWithoutIsStrInference([p2 for d_1074_i2_ in range(default__.abs(20))]), (d_1071_v18_)[default__.safeIndex(p3, len(d_1071_v18_))], globalState)])
                d_1073_v20_ = d_1073_v20_
                d_1075_v21_: _dafny.Array
                nw164_ = _dafny.Array(D6.default()(), 7)
                d_1075_v21_ = nw164_
                d_1076_v22_: _dafny.Set
                d_1076_v22_ = _dafny.Set({d_1052_v4_})
                d_1077_v23_: D6
                d_1077_v23_ = D6_DC13(len(d_1076_v22_), len(_dafny.Set({p1})))
                d_1078_v24_: D6
                d_1078_v24_ = D6_DC14(d_1077_v23_)
                index166_ = default__.safeIndex(717, (d_1075_v21_).length(0))
                (d_1075_v21_)[index166_] = d_1078_v24_
                index167_ = default__.safeIndex(717, (d_1075_v21_).length(0))
                (d_1075_v21_)[index167_] = d_1078_v24_
                (globalState).f16 = p3
            elif True:
                d_1079_v25_: _dafny.Seq
                d_1079_v25_ = _dafny.SeqWithoutIsStrInference([d_1052_v4_, d_1052_v4_, d_1052_v4_])
                (globalState).f21 = (d_1079_v25_)[default__.safeIndex(p1, len(d_1079_v25_))]
                d_1080_v26_: _dafny.MultiSet
                d_1080_v26_ = _dafny.MultiSet([d_1050_i0_])
                d_1081_v27_: _dafny.Map
                d_1081_v27_ = _dafny.Map({d_1080_v26_: d_1052_v4_})
                d_1082_v28_: _dafny.Map
                d_1082_v28_ = d_1081_v27_
                (globalState).f2 = (len((d_1081_v27_))) == (default__.safeModuloInt(p2, p1))
                index168_ = default__.safeIndex(593, (p0).length(0))
                (p0)[index168_] = (p0)[default__.safeIndex(593, (p0).length(0))]
                d_1051_v3_ = d_1051_v3_
                (globalState).f0 = False
        d_1083_v29_: bool
        d_1083_v29_ = False
        d_1084_i3_: int
        d_1084_i3_ = 0
        with _dafny.label("6"):
            while not(d_1083_v29_):
                with _dafny.c_label("6"):
                    if (d_1084_i3_) >= (100):
                        raise _dafny.Break("6")
                    d_1084_i3_ = (d_1084_i3_) + (1)
                    (globalState).f13 = (d_1083_v29_ if d_1083_v29_ else (d_1083_v29_ if d_1083_v29_ else True))
                    d_1085_v30_: _dafny.Array
                    nw165_ = _dafny.Array(D4.default()(), 14)
                    d_1085_v30_ = nw165_
                    d_1086_v31_: D4
                    d_1086_v31_ = D4_DC8(d_1083_v29_)
                    index169_ = default__.safeIndex(56, (d_1085_v30_).length(0))
                    (d_1085_v30_)[index169_] = d_1086_v31_
                    d_1087_v32_: str
                    d_1087_v32_ = _dafny.CodePoint('d')
                    index170_ = default__.safeIndex(56, (d_1085_v30_).length(0))
                    rhs147_ = d_1086_v31_
                    rhs148_ = d_1087_v32_
                    lhs140_ = d_1085_v30_
                    lhs141_ = default__.safeIndex(56, (d_1085_v30_).length(0))
                    lhs140_[lhs141_] = rhs147_
                    d_1087_v32_ = rhs148_
                    d_1088_v33_: _dafny.Seq
                    d_1088_v33_ = _dafny.SeqWithoutIsStrInference([d_1083_v29_])
                    d_1089_v34_: _dafny.Set
                    d_1089_v34_ = _dafny.Set({d_1088_v33_, d_1088_v33_})
                    if (d_1083_v29_) or ((d_1089_v34_).issubset(d_1089_v34_)):
                        d_1090_v35_: C0
                        nw166_ = C0()
                        nw166_.ctor__((-32) >= ((0) - (p2)))
                        d_1090_v35_ = nw166_
                        d_1091_v36_: _dafny.Array
                        nw167_ = _dafny.Array(False, 18)
                        d_1091_v36_ = nw167_
                        d_1091_v36_ = d_1091_v36_
                        d_1047_v0_ = d_1047_v0_
                        index171_ = default__.safeIndex(753, (d_1091_v36_).length(0))
                        (d_1091_v36_)[index171_] = (d_1090_v35_.f27) or (d_1090_v35_.f27)
                        index172_ = default__.safeIndex(753, (d_1091_v36_).length(0))
                        (d_1091_v36_)[index172_] = (d_1088_v33_) != (d_1088_v33_)
                        d_1087_v32_ = d_1087_v32_
                    elif True:
                        d_1092_v37_: _dafny.Map
                        d_1092_v37_ = _dafny.Map({d_1083_v29_: d_1083_v29_})
                        d_1093_v38_: _dafny.Map
                        d_1093_v38_ = _dafny.Map({(0) - ((len(d_1092_v37_)) - (p2)): not(d_1083_v29_)})
                        d_1094_v39_: _dafny.Map
                        d_1094_v39_ = _dafny.Map({d_1083_v29_: p1})
                        d_1095_v40_: _dafny.Seq
                        d_1095_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vn"))
                        d_1096_v41_: _dafny.Set
                        d_1096_v41_ = _dafny.Set({-813})
                        d_1097_v43_: _dafny.Seq
                        def iife88_():
                            coll44_ = _dafny.Map()
                            compr_44_: int
                            for compr_44_ in _dafny.IntegerRange(316, 965):
                                d_1098_v42_: int = compr_44_
                                if ((316) <= (d_1098_v42_)) and ((d_1098_v42_) < (965)):
                                    coll44_[default__.safeModuloInt(d_1098_v42_, p1)] = True
                            return _dafny.Map(coll44_)
                        d_1097_v43_ = _dafny.SeqWithoutIsStrInference([iife88_()
                        ])
                        d_1099_v44_: T1
                        nw168_ = C1()
                        nw168_.ctor__(p2, d_1097_v43_)
                        d_1099_v44_ = nw168_
                        d_1100_v45_: _dafny.Seq
                        d_1100_v45_ = _dafny.SeqWithoutIsStrInference([d_1099_v44_])
                        d_1101_v46_: _dafny.MultiSet
                        d_1101_v46_ = _dafny.MultiSet([d_1099_v44_, (d_1100_v45_)[default__.safeIndex(-289, len(d_1100_v45_))]])
                        rhs149_ = ((d_1093_v38_)[p3] if (p3) in (d_1093_v38_) else d_1083_v29_)
                        rhs150_ = (0) - (((d_1094_v39_)[(self).fm6(p3, globalState)] if ((self).fm6(p3, globalState)) in (d_1094_v39_) else default__.fm0(187, -322, p3, d_1083_v29_, globalState)))
                        rhs151_ = (d_1095_v40_).set(default__.safeIndex(p1, len(d_1095_v40_)), d_1087_v32_)
                        rhs152_ = (p2) > ((len(d_1096_v41_)) - ((0) - (p2)))
                        rhs153_ = default__.safeModuloInt(((d_1101_v46_)[d_1099_v44_] if (d_1099_v44_) in (d_1101_v46_) else p3), p1)
                        lhs142_ = globalState
                        lhs143_ = globalState
                        lhs144_ = globalState
                        lhs145_ = globalState
                        lhs142_.f13 = rhs149_
                        r1 = rhs150_
                        lhs143_.f20 = rhs151_
                        lhs144_.f0 = rhs152_
                        lhs145_.f7 = rhs153_
                        r1 = p2
                        d_1102_v47_: _dafny.Set
                        d_1102_v47_ = _dafny.Set({p0})
                        d_1103_v48_: _dafny.MultiSet
                        d_1103_v48_ = _dafny.MultiSet([len(d_1102_v47_)])
                        rhs154_ = (260 if d_1083_v29_ else default__.safeDivisionInt(((d_1103_v48_)[p2] if (p2) in (d_1103_v48_) else p2), p1))
                        rhs155_ = default__.fm0(((0) - (-766)) + (p1), default__.fm0(p1, p3, p1, False, globalState), len((d_1094_v39_) | (d_1094_v39_)), not (d_1083_v29_) or (d_1083_v29_), globalState)
                        lhs146_ = globalState
                        lhs147_ = globalState
                        lhs146_.f16 = rhs154_
                        lhs147_.f16 = rhs155_
                        d_1104_v49_: C1
                        nw169_ = C1()
                        nw169_.ctor__(p1, ((_dafny.SeqWithoutIsStrInference([d_1093_v38_, _dafny.Map({p2: d_1083_v29_}), d_1093_v38_, (d_1093_v38_).set(p3, d_1083_v29_)])).set(default__.safeIndex(len(d_1095_v40_), len(_dafny.SeqWithoutIsStrInference([d_1093_v38_, _dafny.Map({p2: d_1083_v29_}), d_1093_v38_, (d_1093_v38_).set(p3, d_1083_v29_)]))), d_1093_v38_)).set(default__.safeIndex(p3, len((_dafny.SeqWithoutIsStrInference([d_1093_v38_, _dafny.Map({p2: d_1083_v29_}), d_1093_v38_, (d_1093_v38_).set(p3, d_1083_v29_)])).set(default__.safeIndex(len(d_1095_v40_), len(_dafny.SeqWithoutIsStrInference([d_1093_v38_, _dafny.Map({p2: d_1083_v29_}), d_1093_v38_, (d_1093_v38_).set(p3, d_1083_v29_)]))), d_1093_v38_))), d_1093_v38_))
                        d_1104_v49_ = nw169_
                        d_1105_v50_: _dafny.Array
                        def lambda55_(d_1106_i4_):
                            return True

                        init29_ = lambda55_
                        nw170_ = _dafny.Array(None, 5)
                        for i0_29_ in range(nw170_.length(0)):
                            nw170_[i0_29_] = init29_(i0_29_)
                        d_1105_v50_ = nw170_
                        index173_ = default__.safeIndex(897, (d_1105_v50_).length(0))
                        (d_1105_v50_)[index173_] = d_1083_v29_
                        index174_ = default__.safeIndex(897, (d_1105_v50_).length(0))
                        (d_1105_v50_)[index174_] = d_1083_v29_
                    if d_1083_v29_:
                        d_1107_v51_: _dafny.Seq
                        d_1107_v51_ = _dafny.SeqWithoutIsStrInference([p1])
                        d_1108_v52_: D15
                        d_1108_v52_ = D15_DC33((d_1107_v51_) + (d_1107_v51_))
                        d_1108_v52_ = d_1108_v52_
                        (globalState).f2 = d_1083_v29_
                        (globalState).f7 = p1
                        d_1109_v53_: _dafny.Array
                        def lambda56_(d_1110_v32_):
                            def lambda57_(d_1111_i5_):
                                return d_1110_v32_

                            return lambda57_

                        init30_ = lambda56_(d_1087_v32_)
                        nw171_ = _dafny.Array(None, 2)
                        for i0_30_ in range(nw171_.length(0)):
                            nw171_[i0_30_] = init30_(i0_30_)
                        d_1109_v53_ = nw171_
                        d_1109_v53_ = p0
                        (globalState).f16 = p2
                    elif True:
                        d_1112_v54_: _dafny.Map
                        d_1112_v54_ = _dafny.Map({p3: d_1083_v29_})
                        d_1113_v55_: _dafny.Seq
                        d_1113_v55_ = _dafny.SeqWithoutIsStrInference([d_1112_v54_])
                        d_1114_v56_: C1
                        nw172_ = C1()
                        nw172_.ctor__(default__.safeDivisionInt(p3, p3), (d_1113_v55_) + (d_1113_v55_))
                        d_1114_v56_ = nw172_
                        d_1115_v57_: D12
                        d_1115_v57_ = D12_DC23(p1)
                        d_1116_v58_: _dafny.Seq
                        d_1116_v58_ = _dafny.SeqWithoutIsStrInference([(d_1115_v57_).cf30, d_1114_v56_.f29])
                        index175_ = default__.safeIndex(485, (d_1047_v0_).length(0))
                        (d_1047_v0_)[index175_] = len(d_1116_v58_)
                        d_1117_v59_: _dafny.Seq
                        d_1117_v59_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ehgphob"))
                        index176_ = default__.safeIndex(485, (d_1047_v0_).length(0))
                        (d_1047_v0_)[index176_] = default__.safeModuloInt(len(d_1117_v59_), -723)
                        d_1118_v60_: C0
                        nw173_ = C0()
                        nw173_.ctor__(d_1083_v29_)
                        d_1118_v60_ = nw173_
                        d_1119_v62_: _dafny.Array
                        def lambda58_(d_1120_v56_, d_1121_p1_, d_1122_p3_, d_1123_v59_):
                            def lambda59_(d_1124_i6_):
                                def iife89_():
                                    coll45_ = _dafny.Map()
                                    compr_45_: int
                                    for compr_45_ in _dafny.IntegerRange(-631, 508):
                                        d_1125_v61_: int = compr_45_
                                        if ((-631) <= (d_1125_v61_)) and ((d_1125_v61_) < (508)):
                                            coll45_[(d_1125_v61_) - (d_1121_p1_)] = _dafny.MultiSet([d_1122_p3_])
                                    return _dafny.Map(coll45_)
                                return (iife89_()
                                ) | (_dafny.Map({d_1120_v56_.f29: _dafny.MultiSet([887, len(d_1123_v59_)])}))

                            return lambda59_

                        init31_ = lambda58_(d_1114_v56_, p1, p3, d_1117_v59_)
                        nw174_ = _dafny.Array(None, 14)
                        for i0_31_ in range(nw174_.length(0)):
                            nw174_[i0_31_] = init31_(i0_31_)
                        d_1119_v62_ = nw174_
                        d_1126_v63_: _dafny.MultiSet
                        d_1126_v63_ = _dafny.MultiSet([len(d_1116_v58_)])
                        d_1127_v64_: _dafny.Map
                        d_1127_v64_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p1 for d_1128_i7_ in range(default__.abs(472))]): (_dafny.Map({p3: d_1126_v63_})).set(884, d_1126_v63_)})
                        d_1129_v65_: _dafny.Seq
                        d_1129_v65_ = d_1088_v33_
                        d_1130_v66_: _dafny.Map
                        d_1130_v66_ = _dafny.Map({p1: _dafny.MultiSet([len((d_1129_v65_))])})
                        index177_ = default__.safeIndex(95, (d_1119_v62_).length(0))
                        (d_1119_v62_)[index177_] = ((d_1127_v64_)[d_1116_v58_] if (d_1116_v58_) in (d_1127_v64_) else d_1130_v66_)
                        index178_ = default__.safeIndex(95, (d_1119_v62_).length(0))
                        def iife90_():
                            coll46_ = _dafny.Map()
                            compr_46_: int
                            for compr_46_ in _dafny.IntegerRange(3, 524):
                                d_1131_v67_: int = compr_46_
                                if ((3) <= (d_1131_v67_)) and ((d_1131_v67_) < (524)):
                                    coll46_[default__.safeDivisionInt(d_1131_v67_, (0) - (p2))] = d_1126_v63_
                            return _dafny.Map(coll46_)
                        (d_1119_v62_)[index178_] = iife90_()
                        
                        (globalState).f21 = d_1118_v60_.f27
                    pass
            pass
        d_1132_v68_: C2
        nw175_ = C2()
        nw175_.ctor__(p3)
        d_1132_v68_ = nw175_
        d_1133_v69_: _dafny.Map
        d_1133_v69_ = _dafny.Map({(d_1132_v68_).f28: d_1083_v29_})
        d_1134_v70_: _dafny.Seq
        d_1134_v70_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dlb"))
        d_1135_v71_: _dafny.Seq
        d_1135_v71_ = _dafny.SeqWithoutIsStrInference([d_1133_v69_])
        d_1133_v69_ = ((default__.fm3(p2, d_1083_v29_, d_1134_v70_, (d_1132_v68_).f28, globalState)) | ((d_1135_v71_)[default__.safeIndex((d_1132_v68_).f28, len(d_1135_v71_))])) | (d_1133_v69_)
        guard_loop_8_: int
        for guard_loop_8_ in _dafny.IntegerRange(0, (d_1047_v0_).length(0)):
            d_1136_i8_: int = guard_loop_8_
            if (True) and (((0) <= (d_1136_i8_)) and ((d_1136_i8_) < ((d_1047_v0_).length(0)))):
                (d_1047_v0_)[(d_1136_i8_)] = (d_1136_i8_) - (p2)
        r0 = D1_DC2(d_1133_v69_)
        r1 = default__.fm0(p3, (d_1132_v68_).f28, len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1083_v29_]) for d_1137_i9_ in range(default__.abs(372))])), d_1083_v29_, globalState)
        return r0, r1


class C4:
    def  __init__(self):
        self.f25: bool = False
        self._f26: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    def ctor__(self, f25, f26):
        (self).f25 = f25
        (self)._f26 = f26

    def fm9(self, p0, globalState):
        source18_ = D3_DC6(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tfdl")), _dafny.CodePoint('r'), self.f25)
        if source18_.is_DC6:
            d_1138___mcc_h0_ = source18_.cf6
            d_1139___mcc_h1_ = source18_.cf7
            d_1140___mcc_h2_ = source18_.cf8
            d_1141_cf8_ = d_1140___mcc_h2_
            d_1142_cf7_ = d_1139___mcc_h1_
            d_1143_cf6_ = d_1138___mcc_h0_
            return True
        elif True:
            d_1144___mcc_h3_ = source18_.cf5
            d_1145_cf5_ = d_1144___mcc_h3_
            return (self).f26

    def fm10(self, p0, p1, p2, globalState):
        def iife91_():
            coll47_ = _dafny.Map()
            compr_47_: int
            for compr_47_ in _dafny.IntegerRange(-511, -748):
                d_1146_v0_: int = compr_47_
                if ((-511) <= (d_1146_v0_)) and ((d_1146_v0_) < (-748)):
                    coll47_[(d_1146_v0_) * (len(_dafny.Set({-345})))] = _dafny.CodePoint('g')
            return _dafny.Map(coll47_)
        def iife92_():
            coll48_ = _dafny.Map()
            compr_48_: int
            for compr_48_ in _dafny.IntegerRange(-284, 557):
                d_1147_v1_: int = compr_48_
                if ((-284) <= (d_1147_v1_)) and ((d_1147_v1_) < (557)):
                    coll48_[(d_1147_v1_) + (882)] = _dafny.CodePoint('a')
            return _dafny.Map(coll48_)
        return ((iife91_()
        ) | (_dafny.Map({387: _dafny.CodePoint('s')}))) | ((iife92_()
        ) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([self.f25])): _dafny.CodePoint('p')})))

    def m5(self, p0, globalState):
        d_1148_v0_: _dafny.Array
        nw176_ = _dafny.Array(None, 5)
        nw176_[int(0)] = p0
        nw176_[int(1)] = self.f25
        nw176_[int(2)] = False
        nw176_[int(3)] = (self).f26
        nw176_[int(4)] = self.f25
        d_1148_v0_ = nw176_
        index179_ = default__.safeIndex(513, (d_1148_v0_).length(0))
        (d_1148_v0_)[index179_] = not(p0)
        d_1149_v1_: _dafny.Seq
        d_1149_v1_ = _dafny.SeqWithoutIsStrInference([d_1148_v0_, d_1148_v0_])
        d_1150_v2_: int
        d_1150_v2_ = 683
        d_1151_v3_: _dafny.Map
        d_1151_v3_ = _dafny.Map({((d_1149_v1_)[default__.safeIndex(d_1150_v2_, len(d_1149_v1_))]): not(True)})
        d_1152_v4_: _dafny.Map
        d_1152_v4_ = _dafny.Map({d_1150_v2_: d_1150_v2_})
        d_1153_v5_: _dafny.MultiSet
        d_1153_v5_ = _dafny.MultiSet([d_1150_v2_, -508, len(d_1152_v4_)])
        d_1154_v6_: _dafny.Seq
        d_1154_v6_ = _dafny.SeqWithoutIsStrInference([p0])
        d_1155_v7_: _dafny.Map
        d_1155_v7_ = _dafny.Map({p0: 768})
        d_1156_v8_: _dafny.MultiSet
        d_1156_v8_ = _dafny.MultiSet([self.f25, p0])
        d_1157_v9_: _dafny.Seq
        d_1157_v9_ = _dafny.SeqWithoutIsStrInference([d_1156_v8_, d_1156_v8_])
        index180_ = default__.safeIndex(513, (d_1148_v0_).length(0))
        rhs156_ = ((d_1151_v3_)[d_1148_v0_] if (d_1148_v0_) in (d_1151_v3_) else not(False))
        rhs157_ = default__.safeDivisionInt((((d_1153_v5_)[len(d_1154_v6_)] if (len(d_1154_v6_)) in (d_1153_v5_) else d_1150_v2_)) + (len(d_1155_v7_)), ((d_1153_v5_)[d_1150_v2_] if (d_1150_v2_) in (d_1153_v5_) else d_1150_v2_))
        rhs158_ = (d_1154_v6_)[default__.safeIndex(d_1150_v2_, len(d_1154_v6_))]
        rhs159_ = (default__.fm11(d_1150_v2_, globalState)) <= (d_1157_v9_)
        lhs148_ = self
        lhs149_ = globalState
        lhs150_ = self
        lhs151_ = d_1148_v0_
        lhs152_ = default__.safeIndex(513, (d_1148_v0_).length(0))
        lhs148_.f25 = rhs156_
        lhs149_.f7 = rhs157_
        lhs150_.f25 = rhs158_
        lhs151_[lhs152_] = rhs159_
        guard_loop_9_: int
        for guard_loop_9_ in _dafny.IntegerRange(0, (d_1148_v0_).length(0)):
            d_1158_i0_: int = guard_loop_9_
            if (True) and (((0) <= (d_1158_i0_)) and ((d_1158_i0_) < ((d_1148_v0_).length(0)))):
                (d_1148_v0_)[(d_1158_i0_)] = p0
        hi5_ = d_1150_v2_
        for d_1159_i1_ in range(d_1150_v2_, hi5_):
            d_1160_v10_: _dafny.Array
            def lambda60_(d_1161_i1_):
                def lambda61_(d_1162_i2_):
                    return (d_1162_i2_) + (d_1161_i1_)

                return lambda61_

            init32_ = lambda60_(d_1159_i1_)
            nw177_ = _dafny.Array(None, 15)
            for i0_32_ in range(nw177_.length(0)):
                nw177_[i0_32_] = init32_(i0_32_)
            d_1160_v10_ = nw177_
            d_1163_v11_: _dafny.Seq
            d_1163_v11_ = _dafny.SeqWithoutIsStrInference([d_1159_i1_])
            index181_ = default__.safeIndex(387, (d_1160_v10_).length(0))
            (d_1160_v10_)[index181_] = ((d_1163_v11_)[default__.safeIndex(d_1150_v2_, len(d_1163_v11_))]) + (len(d_1152_v4_))
            index182_ = default__.safeIndex(387, (d_1160_v10_).length(0))
            (d_1160_v10_)[index182_] = ((d_1156_v8_)[(d_1148_v0_)[default__.safeIndex(513, (d_1148_v0_).length(0))]] if ((d_1148_v0_)[default__.safeIndex(513, (d_1148_v0_).length(0))]) in (d_1156_v8_) else d_1159_i1_)
            d_1164_v12_: _dafny.Map
            d_1164_v12_ = _dafny.Map({d_1148_v0_: default__.safeModuloInt(d_1150_v2_, 475)})
            d_1165_v13_: str
            d_1165_v13_ = _dafny.CodePoint('v')
            d_1166_v14_: _dafny.Map
            d_1166_v14_ = _dafny.Map({(d_1159_i1_) * (d_1150_v2_): (d_1164_v12_).set(d_1148_v0_, d_1159_i1_)})
            index183_ = default__.safeIndex(513, (d_1148_v0_).length(0))
            rhs160_ = False
            rhs161_ = ((d_1166_v14_)[(d_1160_v10_)[default__.safeIndex(387, (d_1160_v10_).length(0))]] if ((d_1160_v10_)[default__.safeIndex(387, (d_1160_v10_).length(0))]) in (d_1166_v14_) else ((d_1166_v14_)[d_1150_v2_] if (d_1150_v2_) in (d_1166_v14_) else d_1164_v12_))
            rhs162_ = (((d_1153_v5_)[11] if (11) in (d_1153_v5_) else -590)) * (d_1150_v2_)
            rhs163_ = (d_1154_v6_)[default__.safeIndex((d_1160_v10_)[default__.safeIndex(387, (d_1160_v10_).length(0))], len(d_1154_v6_))]
            rhs164_ = d_1165_v13_
            lhs153_ = d_1148_v0_
            lhs154_ = default__.safeIndex(513, (d_1148_v0_).length(0))
            lhs155_ = globalState
            lhs156_ = globalState
            lhs153_[lhs154_] = rhs160_
            d_1164_v12_ = rhs161_
            lhs155_.f7 = rhs162_
            lhs156_.f21 = rhs163_
            d_1165_v13_ = rhs164_
            index184_ = default__.safeIndex(387, (d_1160_v10_).length(0))
            (d_1160_v10_)[index184_] = (0) - (default__.safeDivisionInt(28, d_1150_v2_))
            d_1156_v8_ = ((_dafny.MultiSet([(d_1148_v0_)[default__.safeIndex(513, (d_1148_v0_).length(0))], (self).fm9(p0, globalState)])).set(self.f25, default__.abs(d_1159_i1_))) - ((d_1156_v8_) | (d_1156_v8_))
        hi6_ = (0) - (d_1150_v2_)
        for d_1167_i3_ in range(default__.safeModuloInt(529, d_1150_v2_), hi6_):
            d_1168_v15_: _dafny.Map
            d_1168_v15_ = _dafny.Map({(d_1150_v2_) + (d_1167_i3_): (d_1148_v0_)[default__.safeIndex(513, (d_1148_v0_).length(0))]})
            d_1168_v15_ = (d_1168_v15_).set(d_1150_v2_, p0)
            d_1169_v16_: _dafny.MultiSet
            d_1169_v16_ = _dafny.MultiSet([d_1148_v0_])
            d_1170_v17_: C2
            nw178_ = C2()
            nw178_.ctor__((d_1169_v16_).cardinality)
            d_1170_v17_ = nw178_
            index185_ = default__.safeIndex(513, (d_1148_v0_).length(0))
            (d_1148_v0_)[index185_] = (d_1148_v0_)[default__.safeIndex(513, (d_1148_v0_).length(0))]
            d_1171_v18_: D13
            d_1171_v18_ = D13_DC27(not(self.f25))
            source19_ = (d_1171_v18_ if self.f25 else d_1171_v18_)
            if source19_.is_DC26:
                d_1172___mcc_h0_ = source19_.cf33
                d_1173___mcc_h1_ = source19_.cf34
                d_1174___mcc_h2_ = source19_.cf35
                d_1175_cf35_ = d_1174___mcc_h2_
                d_1176_cf34_ = d_1173___mcc_h1_
                d_1177_cf33_ = d_1172___mcc_h0_
                d_1178_v19_: str
                d_1178_v19_ = _dafny.CodePoint('f')
                (globalState).f2 = not((d_1154_v6_)[default__.safeIndex(default__.safeDivisionInt((d_1170_v17_).f28, len(_dafny.Map({(d_1170_v17_).f28: d_1178_v19_}))), len(d_1154_v6_))])
                d_1179_v20_: C2
                nw179_ = C2()
                nw179_.ctor__((758) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tmdic")))))
                d_1179_v20_ = nw179_
                d_1180_v21_: _dafny.Array
                nw180_ = _dafny.Array(int(0), 24)
                d_1180_v21_ = nw180_
                (d_1179_v20_).m3(d_1180_v21_, globalState)
                d_1181_v22_: _dafny.Seq
                d_1181_v22_ = _dafny.SeqWithoutIsStrInference([d_1150_v2_])
                d_1182_v23_: _dafny.Map
                d_1182_v23_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "knj")): (d_1181_v22_) + (d_1181_v22_)})
                d_1182_v23_ = _dafny.Map({d_1177_cf33_: d_1181_v22_})
            elif source19_.is_DC27:
                d_1183___mcc_h3_ = source19_.cf36
                d_1184_cf36_ = d_1183___mcc_h3_
                (globalState).f13 = d_1184_cf36_
                (globalState).f7 = ((d_1170_v17_).f28) * (d_1150_v2_)
                d_1185_v24_: _dafny.Seq
                d_1185_v24_ = _dafny.SeqWithoutIsStrInference([d_1150_v2_, (d_1170_v17_).f28])
                d_1186_v25_: _dafny.Seq
                d_1186_v25_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([d_1167_i3_ for d_1187_i4_ in range(default__.abs(794))])) + (_dafny.SeqWithoutIsStrInference([(d_1170_v17_).f28 for d_1188_i5_ in range(default__.abs(-680))])), (d_1185_v24_ if self.f25 else d_1185_v24_), d_1185_v24_])
                d_1189_v26_: _dafny.Map
                d_1189_v26_ = _dafny.Map({len(d_1168_v15_): d_1154_v6_})
                d_1190_v27_: str
                d_1190_v27_ = _dafny.CodePoint('p')
                d_1191_v28_: _dafny.Map
                d_1191_v28_ = _dafny.Map({(d_1148_v0_)[default__.safeIndex(513, (d_1148_v0_).length(0))]: self.f25})
                d_1192_v29_: _dafny.Array
                nw181_ = _dafny.Array(None, 19)
                nw181_[int(0)] = (d_1154_v6_).set(default__.safeIndex(len((d_1189_v26_).set(default__.fm0(d_1150_v2_, d_1150_v2_, (d_1170_v17_).f28, False, globalState), d_1154_v6_)), len(d_1154_v6_)), p0)
                nw181_[int(1)] = (default__.fm18(d_1184_cf36_, d_1190_v27_, self.f25, globalState)) + (default__.fm18((self).f26, d_1190_v27_, p0, globalState))
                nw181_[int(2)] = (d_1154_v6_) + (_dafny.SeqWithoutIsStrInference([True, default__.fm2(d_1191_v28_, globalState), True]))
                nw181_[int(3)] = d_1154_v6_
                nw181_[int(4)] = d_1154_v6_
                nw181_[int(5)] = d_1154_v6_
                nw181_[int(6)] = (d_1154_v6_) + (d_1154_v6_)
                nw181_[int(7)] = d_1154_v6_
                nw181_[int(8)] = (_dafny.SeqWithoutIsStrInference([(d_1148_v0_)[default__.safeIndex(513, (d_1148_v0_).length(0))], not((d_1148_v0_)[default__.safeIndex(513, (d_1148_v0_).length(0))])])) + (default__.fm18(d_1184_cf36_, d_1190_v27_, p0, globalState))
                nw181_[int(9)] = d_1154_v6_
                nw181_[int(10)] = (_dafny.SeqWithoutIsStrInference([(d_1148_v0_)[default__.safeIndex(513, (d_1148_v0_).length(0))]])) + (d_1154_v6_)
                nw181_[int(11)] = d_1154_v6_
                nw181_[int(12)] = (d_1154_v6_) + (_dafny.SeqWithoutIsStrInference([(d_1148_v0_)[default__.safeIndex(513, (d_1148_v0_).length(0))]]))
                nw181_[int(13)] = (d_1154_v6_) + (d_1154_v6_)
                nw181_[int(14)] = (d_1154_v6_) + (d_1154_v6_)
                nw181_[int(15)] = (d_1154_v6_) + (d_1154_v6_)
                nw181_[int(16)] = d_1154_v6_
                nw181_[int(17)] = _dafny.SeqWithoutIsStrInference([self.f25])
                nw181_[int(18)] = _dafny.SeqWithoutIsStrInference([p0])
                d_1192_v29_ = nw181_
                index186_ = default__.safeIndex(952, (d_1192_v29_).length(0))
                (d_1192_v29_)[index186_] = d_1154_v6_
                d_1193_v30_: _dafny.Seq
                d_1193_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ytqvm"))
                d_1194_v31_: _dafny.Set
                d_1194_v31_ = _dafny.Set({d_1193_v30_, d_1193_v30_})
                index187_ = default__.safeIndex(952, (d_1192_v29_).length(0))
                rhs165_ = False
                rhs166_ = self.f25
                rhs167_ = default__.fm22(default__.fm2(_dafny.Map({d_1184_cf36_: (d_1148_v0_)[default__.safeIndex(513, (d_1148_v0_).length(0))]}), globalState), (d_1194_v31_).isdisjoint(d_1194_v31_), p0, globalState)
                rhs168_ = (d_1154_v6_) + (d_1154_v6_)
                lhs157_ = self
                lhs158_ = globalState
                lhs159_ = d_1192_v29_
                lhs160_ = default__.safeIndex(952, (d_1192_v29_).length(0))
                lhs157_.f25 = rhs165_
                lhs158_.f13 = rhs166_
                d_1186_v25_ = rhs167_
                lhs159_[lhs160_] = rhs168_
                d_1195_v32_: _dafny.Array
                nw182_ = _dafny.Array(None, 15)
                nw182_[int(0)] = 770
                nw182_[int(1)] = d_1150_v2_
                nw182_[int(2)] = (d_1170_v17_).f28
                nw182_[int(3)] = len(d_1193_v30_)
                nw182_[int(4)] = (d_1170_v17_).f28
                nw182_[int(5)] = (d_1170_v17_).f28
                nw182_[int(6)] = (d_1170_v17_).f28
                nw182_[int(7)] = d_1167_i3_
                nw182_[int(8)] = d_1150_v2_
                nw182_[int(9)] = 140
                nw182_[int(10)] = d_1150_v2_
                nw182_[int(11)] = (d_1170_v17_).f28
                nw182_[int(12)] = d_1167_i3_
                nw182_[int(13)] = d_1167_i3_
                nw182_[int(14)] = (d_1170_v17_).f28
                d_1195_v32_ = nw182_
                d_1196_v33_: _dafny.Set
                d_1196_v33_ = _dafny.Set({d_1195_v32_, d_1195_v32_, d_1195_v32_})
                (globalState).f13 = ((d_1148_v0_)[default__.safeIndex(513, (d_1148_v0_).length(0))] if not(((d_1191_v28_)[True] if (True) in (d_1191_v28_) else self.f25)) else (d_1195_v32_) in (d_1196_v33_))
            elif True:
                d_1197___mcc_h4_ = source19_.cf32
                d_1198_cf32_ = d_1197___mcc_h4_
                d_1199_v34_: _dafny.Array
                nw183_ = _dafny.Array(_dafny.Array(None, 0), 29)
                d_1199_v34_ = nw183_
                d_1200_v35_: _dafny.Array
                nw184_ = _dafny.Array(int(0), 20)
                d_1200_v35_ = nw184_
                index188_ = default__.safeIndex(292, (d_1199_v34_).length(0))
                (d_1199_v34_)[index188_] = d_1200_v35_
                index189_ = default__.safeIndex(122, (d_1200_v35_).length(0))
                (d_1200_v35_)[index189_] = (0) - (((d_1153_v5_)[(d_1170_v17_).f28] if ((d_1170_v17_).f28) in (d_1153_v5_) else d_1167_i3_))
                d_1201_v36_: _dafny.Array
                def lambda62_(d_1202_i6_):
                    return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_1203_i7_ in range(default__.abs(887))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yxeguqbr")))

                init33_ = lambda62_
                nw185_ = _dafny.Array(None, 8)
                for i0_33_ in range(nw185_.length(0)):
                    nw185_[i0_33_] = init33_(i0_33_)
                d_1201_v36_ = nw185_
                d_1204_v37_: _dafny.Seq
                d_1204_v37_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))
                index190_ = default__.safeIndex(540, (d_1201_v36_).length(0))
                (d_1201_v36_)[index190_] = d_1204_v37_
                d_1205_v38_: _dafny.Map
                d_1205_v38_ = _dafny.Map({d_1200_v35_: _dafny.CodePoint('d')})
                d_1206_v39_: _dafny.Map
                d_1206_v39_ = _dafny.Map({self.f25: (d_1148_v0_)[default__.safeIndex(513, (d_1148_v0_).length(0))]})
                d_1207_v40_: _dafny.Set
                d_1207_v40_ = _dafny.Set({d_1200_v35_})
                index191_ = default__.safeIndex(292, (d_1199_v34_).length(0))
                index192_ = default__.safeIndex(122, (d_1200_v35_).length(0))
                index193_ = default__.safeIndex(540, (d_1201_v36_).length(0))
                rhs169_ = d_1200_v35_
                rhs170_ = (0) - (len((d_1205_v38_) | (d_1205_v38_)))
                rhs171_ = default__.fm2(d_1206_v39_, globalState)
                rhs172_ = len(((d_1207_v40_) | (d_1207_v40_)).intersection(d_1207_v40_))
                rhs173_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iaka"))) + (d_1204_v37_)
                lhs161_ = d_1199_v34_
                lhs162_ = default__.safeIndex(292, (d_1199_v34_).length(0))
                lhs163_ = globalState
                lhs164_ = globalState
                lhs165_ = d_1200_v35_
                lhs166_ = default__.safeIndex(122, (d_1200_v35_).length(0))
                lhs167_ = d_1201_v36_
                lhs168_ = default__.safeIndex(540, (d_1201_v36_).length(0))
                lhs161_[lhs162_] = rhs169_
                lhs163_.f16 = rhs170_
                lhs164_.f2 = rhs171_
                lhs165_[lhs166_] = rhs172_
                lhs167_[lhs168_] = rhs173_
                d_1208_v41_: _dafny.Set
                d_1208_v41_ = _dafny.Set({(d_1148_v0_)[default__.safeIndex(513, (d_1148_v0_).length(0))], p0})
                d_1209_v42_: _dafny.Map
                d_1209_v42_ = _dafny.Map({d_1208_v41_: not (self.f25) or ((d_1148_v0_)[default__.safeIndex(513, (d_1148_v0_).length(0))])})
                d_1209_v42_ = (d_1209_v42_).set((d_1208_v41_).intersection(d_1208_v41_), p0)
                (globalState).f7 = default__.safeDivisionInt((d_1170_v17_).f28, (d_1170_v17_).f28)
                d_1210_v43_: C2
                nw186_ = C2()
                nw186_.ctor__((d_1200_v35_)[default__.safeIndex(122, (d_1200_v35_).length(0))])
                d_1210_v43_ = nw186_
        d_1148_v0_ = (d_1148_v0_ if False else d_1148_v0_)
        d_1211_v44_: _dafny.Map
        d_1211_v44_ = _dafny.Map({p0: p0})
        rhs174_ = (d_1154_v6_)[default__.safeIndex((160) * (d_1150_v2_), len(d_1154_v6_))]
        rhs175_ = _dafny.MultiSet([p0, (d_1150_v2_) < (len(d_1211_v44_)), p0])
        lhs169_ = globalState
        lhs169_.f21 = rhs174_
        d_1156_v8_ = rhs175_

    def m6(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: bool = False
        r3: bool = False
        d_1212_v0_: _dafny.Seq
        d_1212_v0_ = _dafny.SeqWithoutIsStrInference([default__.fm0(p3, 950, p3, self.f25, globalState)])
        d_1213_v1_: _dafny.Array
        def lambda63_(d_1214_i0_):
            return (self.f25 if (self).f26 else (self).f26)

        init34_ = lambda63_
        nw187_ = _dafny.Array(None, 1)
        for i0_34_ in range(nw187_.length(0)):
            nw187_[i0_34_] = init34_(i0_34_)
        d_1213_v1_ = nw187_
        d_1215_v2_: _dafny.Map
        d_1215_v2_ = _dafny.Map({p3: self.f25})
        d_1216_v5_: D1
        def iife93_():
            coll49_ = _dafny.Map()
            compr_49_: int
            for compr_49_ in (d_1212_v0_).Elements:
                d_1217_v4_: int = compr_49_
                if (d_1217_v4_) in (d_1212_v0_):
                    coll49_[(d_1217_v4_) + (p3)] = self.f25
            return _dafny.Map(coll49_)
        d_1216_v5_ = D1_DC2(iife93_()
)
        d_1218_v7_: _dafny.Seq
        d_1218_v7_ = _dafny.SeqWithoutIsStrInference([(self).f26, (self).f26])
        d_1219_v8_: _dafny.Map
        d_1219_v8_ = _dafny.Map({p3: d_1215_v2_})
        d_1220_v10_: _dafny.Map
        d_1220_v10_ = _dafny.Map({(self).f26: True})
        d_1221_v13_: _dafny.Array
        nw188_ = _dafny.Array(None, 23)
        nw188_[int(0)] = d_1215_v2_
        nw188_[int(1)] = d_1215_v2_
        nw188_[int(2)] = d_1215_v2_
        def iife94_():
            coll50_ = _dafny.Map()
            compr_50_: int
            for compr_50_ in _dafny.IntegerRange(306, 956):
                d_1222_v3_: int = compr_50_
                if ((306) <= (d_1222_v3_)) and ((d_1222_v3_) < (956)):
                    coll50_[default__.safeModuloInt(d_1222_v3_, p3)] = self.f25
            return _dafny.Map(coll50_)
        nw188_[int(3)] = iife94_()
        
        nw188_[int(4)] = (d_1216_v5_).cf2
        def iife95_():
            coll51_ = _dafny.Map()
            compr_51_: int
            for compr_51_ in _dafny.IntegerRange(766, 466):
                d_1223_v6_: int = compr_51_
                if ((766) <= (d_1223_v6_)) and ((d_1223_v6_) < (466)):
                    coll51_[default__.safeModuloInt(d_1223_v6_, p3)] = self.f25
            return _dafny.Map(coll51_)
        nw188_[int(5)] = (d_1215_v2_) | (iife95_()
        )
        nw188_[int(6)] = (d_1215_v2_) | (_dafny.Map({p3: (d_1218_v7_)[default__.safeIndex(p3, len(d_1218_v7_))]}))
        nw188_[int(7)] = (d_1215_v2_) | (_dafny.Map({p3: (self).f26}))
        nw188_[int(8)] = d_1215_v2_
        nw188_[int(9)] = d_1215_v2_
        nw188_[int(10)] = d_1215_v2_
        def iife96_():
            coll52_ = _dafny.Map()
            compr_52_: int
            for compr_52_ in _dafny.IntegerRange(658, 762):
                d_1224_v9_: int = compr_52_
                if ((658) <= (d_1224_v9_)) and ((d_1224_v9_) < (762)):
                    coll52_[default__.safeDivisionInt(d_1224_v9_, p3)] = (self).f26
            return _dafny.Map(coll52_)
        nw188_[int(11)] = (((d_1219_v8_)[p3] if (p3) in (d_1219_v8_) else (_dafny.Map({p3: (self).f26})).set(p3, True))) | (iife96_()
        )
        nw188_[int(12)] = d_1215_v2_
        nw188_[int(13)] = d_1215_v2_
        nw188_[int(14)] = d_1215_v2_
        nw188_[int(15)] = d_1215_v2_
        nw188_[int(16)] = d_1215_v2_
        nw188_[int(17)] = d_1215_v2_
        nw188_[int(18)] = _dafny.Map({len(d_1220_v10_): True})
        nw188_[int(19)] = d_1215_v2_
        nw188_[int(20)] = d_1215_v2_
        def iife97_():
            coll53_ = _dafny.Map()
            compr_53_: int
            for compr_53_ in _dafny.IntegerRange(444, 576):
                d_1225_v11_: int = compr_53_
                if ((444) <= (d_1225_v11_)) and ((d_1225_v11_) < (576)):
                    coll53_[default__.safeModuloInt(d_1225_v11_, len(_dafny.Map({self.f25: (self).f26})))] = self.f25
            return _dafny.Map(coll53_)
        nw188_[int(21)] = iife97_()
        
        def iife98_():
            coll54_ = _dafny.Map()
            compr_54_: int
            for compr_54_ in _dafny.IntegerRange(788, 326):
                d_1226_v12_: int = compr_54_
                if ((788) <= (d_1226_v12_)) and ((d_1226_v12_) < (326)):
                    coll54_[default__.safeModuloInt(d_1226_v12_, p3)] = self.f25
            return _dafny.Map(coll54_)
        nw188_[int(22)] = iife98_()
        
        d_1221_v13_ = nw188_
        d_1227_v14_: _dafny.MultiSet
        d_1227_v14_ = _dafny.MultiSet([not (self.f25) or (not((self).f26)), (self).fm9(self.f25, globalState), (self).f26, not(default__.fm2(d_1220_v10_, globalState)), (((d_1215_v2_)[p3] if (p3) in (d_1215_v2_) else (self).f26)) or (self.f25)])
        d_1228_v15_: _dafny.Seq
        d_1228_v15_ = _dafny.SeqWithoutIsStrInference([d_1213_v1_])
        rhs176_ = d_1212_v0_
        rhs177_ = (d_1228_v15_)[default__.safeIndex((p3) + (p3), len(d_1228_v15_))]
        rhs178_ = d_1221_v13_
        rhs179_ = d_1227_v14_
        d_1212_v0_ = rhs176_
        d_1213_v1_ = rhs177_
        d_1221_v13_ = rhs178_
        d_1227_v14_ = rhs179_
        d_1229_v16_: D12
        d_1229_v16_ = D12_DC23((0) - (p3))
        d_1230_v17_: _dafny.Map
        d_1230_v17_ = _dafny.Map({self.f25: d_1229_v16_})
        d_1230_v17_ = d_1230_v17_
        r1 = p3
        d_1231_v18_: _dafny.Array
        nw189_ = _dafny.Array(int(0), 7)
        d_1231_v18_ = nw189_
        d_1231_v18_ = p0
        d_1232_i1_: int
        d_1232_i1_ = 0
        with _dafny.label("7"):
            while (self).fm9((self.f25) or (self.f25), globalState):
                with _dafny.c_label("7"):
                    if (d_1232_i1_) >= (100):
                        raise _dafny.Break("7")
                    d_1232_i1_ = (d_1232_i1_) + (1)
                    d_1233_v19_: C3
                    nw190_ = C3()
                    nw190_.ctor__()
                    d_1233_v19_ = nw190_
                    if (_dafny.CodePoint('u')) not in (p1):
                        d_1213_v1_ = d_1213_v1_
                        index194_ = default__.safeIndex(813, (d_1213_v1_).length(0))
                        (d_1213_v1_)[index194_] = self.f25
                        index195_ = default__.safeIndex(813, (d_1213_v1_).length(0))
                        rhs180_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iongp"))) < ((p1) + (p1))
                        rhs181_ = (self).f26
                        rhs182_ = len((p1) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "osab"))))
                        lhs170_ = self
                        lhs171_ = d_1213_v1_
                        lhs172_ = default__.safeIndex(813, (d_1213_v1_).length(0))
                        lhs173_ = globalState
                        lhs170_.f25 = rhs180_
                        lhs171_[lhs172_] = rhs181_
                        lhs173_.f16 = rhs182_
                        d_1234_v20_: str
                        d_1234_v20_ = _dafny.CodePoint('q')
                        d_1235_v21_: _dafny.Map
                        d_1235_v21_ = _dafny.Map({not((d_1213_v1_)[default__.safeIndex(813, (d_1213_v1_).length(0))]): default__.fm0(p3, 169, p3, False, globalState)})
                        d_1236_v22_: _dafny.Map
                        d_1236_v22_ = _dafny.Map({d_1234_v20_: (0) - ((0) - ((0) - (len((d_1235_v21_) | (d_1235_v21_)))))})
                        d_1236_v22_ = (d_1236_v22_).set(d_1234_v20_, (-25) - (((d_1235_v21_)[(d_1213_v1_)[default__.safeIndex(813, (d_1213_v1_).length(0))]] if ((d_1213_v1_)[default__.safeIndex(813, (d_1213_v1_).length(0))]) in (d_1235_v21_) else p3)))
                        d_1237_v23_: _dafny.Array
                        def lambda64_(d_1238_v21_):
                            def lambda65_(d_1239_i2_):
                                return d_1238_v21_

                            return lambda65_

                        init35_ = lambda64_(d_1235_v21_)
                        nw191_ = _dafny.Array(None, 23)
                        for i0_35_ in range(nw191_.length(0)):
                            nw191_[i0_35_] = init35_(i0_35_)
                        d_1237_v23_ = nw191_
                        d_1237_v23_ = d_1237_v23_
                        d_1240_v24_: _dafny.Array
                        nw192_ = _dafny.Array(_dafny.CodePoint('D'), 1)
                        d_1240_v24_ = nw192_
                        d_1240_v24_ = d_1240_v24_
                    elif True:
                        d_1241_v25_: _dafny.Map
                        d_1241_v25_ = _dafny.Map({p3: (p3) * ((0) - (p3))})
                        d_1242_v27_: _dafny.Map
                        def iife99_():
                            coll55_ = _dafny.Set()
                            compr_55_: int
                            for compr_55_ in _dafny.IntegerRange(-718, -647):
                                d_1243_v26_: int = compr_55_
                                if ((-718) <= (d_1243_v26_)) and ((d_1243_v26_) < (-647)):
                                    coll55_ = coll55_.union(_dafny.Set([default__.safeDivisionInt(d_1243_v26_, p3)]))
                            return _dafny.Set(coll55_)
                        d_1242_v27_ = _dafny.Map({iife99_()
                        : self.f25})
                        d_1241_v25_ = (d_1241_v25_).set(len(d_1242_v27_), p3)
                        d_1244_v28_: D14
                        d_1244_v28_ = D14_DC30((d_1212_v0_) + (d_1212_v0_), (d_1218_v7_) + (d_1218_v7_))
                        d_1244_v28_ = d_1244_v28_
                        d_1245_v29_: _dafny.Array
                        nw193_ = _dafny.Array(None, 7)
                        nw193_[int(0)] = p2
                        nw193_[int(1)] = p2
                        nw193_[int(2)] = p0
                        nw193_[int(3)] = p2
                        nw193_[int(4)] = p2
                        nw193_[int(5)] = p0
                        nw193_[int(6)] = p2
                        d_1245_v29_ = nw193_
                        index196_ = default__.safeIndex(638, (d_1245_v29_).length(0))
                        (d_1245_v29_)[index196_] = p2
                        index197_ = default__.safeIndex(638, (d_1245_v29_).length(0))
                        (d_1245_v29_)[index197_] = d_1231_v18_
                        d_1246_v30_: _dafny.Array
                        nw194_ = _dafny.Array(_dafny.CodePoint('D'), 11)
                        d_1246_v30_ = nw194_
                        d_1247_v31_: str
                        d_1247_v31_ = _dafny.CodePoint('r')
                        index198_ = default__.safeIndex(501, (d_1246_v30_).length(0))
                        (d_1246_v30_)[index198_] = d_1247_v31_
                        index199_ = default__.safeIndex(501, (d_1246_v30_).length(0))
                        (d_1246_v30_)[index199_] = d_1247_v31_
                        d_1248_v33_: _dafny.Seq
                        def iife100_():
                            coll56_ = _dafny.Map()
                            compr_56_: int
                            for compr_56_ in _dafny.IntegerRange(-929, 930):
                                d_1249_v32_: int = compr_56_
                                if ((-929) <= (d_1249_v32_)) and ((d_1249_v32_) < (930)):
                                    coll56_[(d_1249_v32_) * (len(_dafny.Set({p3, p3})))] = self.f25
                            return _dafny.Map(coll56_)
                        d_1248_v33_ = _dafny.SeqWithoutIsStrInference([iife100_()
                        ])
                        d_1250_v34_: T1
                        nw195_ = C1()
                        nw195_.ctor__(len(p1), d_1248_v33_)
                        d_1250_v34_ = nw195_
                        d_1251_v35_: _dafny.Map
                        d_1251_v35_ = _dafny.Map({d_1250_v34_: p3})
                        d_1251_v35_ = (d_1251_v35_).set(d_1250_v34_, p3)
                    (globalState).f13 = self.f25
                    d_1252_v36_: _dafny.Array
                    nw196_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 2)
                    d_1252_v36_ = nw196_
                    index200_ = default__.safeIndex(892, (d_1252_v36_).length(0))
                    (d_1252_v36_)[index200_] = p1
                    d_1253_v37_: str
                    d_1253_v37_ = _dafny.CodePoint('d')
                    index201_ = default__.safeIndex(892, (d_1252_v36_).length(0))
                    (d_1252_v36_)[index201_] = ((p1) + (_dafny.SeqWithoutIsStrInference([d_1253_v37_, d_1253_v37_]))) + (p1)
                    pass
            pass
        d_1254_i3_: int
        d_1254_i3_ = 0
        with _dafny.label("8"):
            while ((self).f26 if self.f25 else True):
                with _dafny.c_label("8"):
                    if (d_1254_i3_) >= (100):
                        raise _dafny.Break("8")
                    d_1254_i3_ = (d_1254_i3_) + (1)
                    (globalState).f21 = self.f25
                    d_1255_v39_: _dafny.Seq
                    def iife101_():
                        coll57_ = _dafny.Map()
                        compr_57_: int
                        for compr_57_ in _dafny.IntegerRange(888, 272):
                            d_1256_v38_: int = compr_57_
                            if ((888) <= (d_1256_v38_)) and ((d_1256_v38_) < (272)):
                                coll57_[default__.safeDivisionInt(d_1256_v38_, p3)] = (self).f26
                        return _dafny.Map(coll57_)
                    d_1255_v39_ = _dafny.SeqWithoutIsStrInference([iife101_()
                    ])
                    d_1257_v40_: C1
                    nw197_ = C1()
                    nw197_.ctor__(p3, (_dafny.SeqWithoutIsStrInference([d_1215_v2_ for d_1258_i4_ in range(default__.abs(483))])) + (d_1255_v39_))
                    d_1257_v40_ = nw197_
                    r1 = (-433) * (default__.safeDivisionInt(p3, p3))
                    d_1259_v41_: _dafny.Set
                    d_1259_v41_ = _dafny.Set({(self).f26, self.f25})
                    d_1260_v42_: D4
                    d_1260_v42_ = D4_DC7(len(d_1259_v41_))
                    (d_1257_v40_).f29 = (d_1260_v42_).cf9
                    pass
            pass
        r0 = not(self.f25)
        r1 = (p3) * ((len(d_1218_v7_)) + (p3))
        d_1261_v43_: D18
        d_1261_v43_ = D18_DC42(d_1220_v10_)
        r2 = default__.fm2((d_1261_v43_).cf58, globalState)
        d_1262_v44_: _dafny.Set
        d_1262_v44_ = _dafny.Set({p2, p2})
        r3 = (d_1262_v44_).ispropersubset(d_1262_v44_)
        return r0, r1, r2, r3

    @property
    def f26(self):
        return self._f26

class C5(T1, T0):
    def  __init__(self):
        self._f24: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    def ctor__(self, f24):
        (self)._f24 = f24

    def fm7(self, p0, p1, p2, globalState):
        return (self).f24

    def fm5(self, p0, p1, p2, p3, globalState):
        return ((_dafny.Map({(self).f24: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wx"))})) | (_dafny.Map({(self).f24: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mjsnrtion"))}))) | ((_dafny.Map({(self).f24: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_1263_i0_ in range(default__.abs(378))])})) | (_dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sjvmhb"))})))

    def fm6(self, p0, globalState):
        source20_ = D3_DC5(_dafny.CodePoint('x'))
        if source20_.is_DC6:
            d_1264___mcc_h0_ = source20_.cf6
            d_1265___mcc_h1_ = source20_.cf7
            d_1266___mcc_h2_ = source20_.cf8
            d_1267_cf8_ = d_1266___mcc_h2_
            d_1268_cf7_ = d_1265___mcc_h1_
            d_1269_cf6_ = d_1264___mcc_h0_
            return (self).f24
        elif True:
            d_1270___mcc_h3_ = source20_.cf5
            d_1271_cf5_ = d_1270___mcc_h3_
            return ((self).f24) or ((self).f24)

    def fm8(self, p0, p1, p2, globalState):
        return (0) - (((len(_dafny.Map({(0) - ((_dafny.MultiSet([_dafny.CodePoint('t')])).cardinality): (self).f24}))) - (-530)) * (86))

    def m2(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: _dafny.Array = _dafny.Array(None, 0)
        d_1272_v0_: _dafny.Array
        nw198_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 10)
        d_1272_v0_ = nw198_
        guard_loop_10_: int
        for guard_loop_10_ in _dafny.IntegerRange(0, (d_1272_v0_).length(0)):
            d_1273_i0_: int = guard_loop_10_
            if (True) and (((0) <= (d_1273_i0_)) and ((d_1273_i0_) < ((d_1272_v0_).length(0)))):
                (d_1272_v0_)[(d_1273_i0_)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ckiixp"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ogsr")))
        d_1274_v1_: D4
        d_1274_v1_ = D4_DC7(p3)
        (globalState).f13 = (self).fm7(p1, (p3) - (p3), (215) != ((d_1274_v1_).cf9), globalState)
        if (852) != (p0):
            d_1275_v2_: _dafny.Map
            d_1275_v2_ = _dafny.Map({(self).f24: p1})
            (globalState).f0 = default__.fm2(d_1275_v2_, globalState)
            d_1276_v3_: _dafny.Map
            d_1276_v3_ = _dafny.Map({p0: p1})
            d_1277_v4_: _dafny.Set
            d_1277_v4_ = _dafny.Set({(self).fm6(len(_dafny.Map({d_1276_v3_: default__.fm0(p3, p0, p3, p1, globalState)})), globalState)})
            r1 = (p3) - (len((_dafny.Set({(self).f24, p1})) - (d_1277_v4_)))
            (globalState).f7 = default__.safeModuloInt((412) * (p3), p0)
            d_1278_v5_: _dafny.Array
            nw199_ = _dafny.Array(int(0), 4)
            d_1278_v5_ = nw199_
            d_1279_v6_: _dafny.Seq
            d_1279_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bqijp"))
            d_1280_v7_: str
            d_1280_v7_ = _dafny.CodePoint('s')
            d_1281_v8_: _dafny.Seq
            d_1281_v8_ = _dafny.SeqWithoutIsStrInference([not(True)])
            rhs183_ = default__.safeDivisionInt(len(d_1276_v3_), p0)
            rhs184_ = not ((p1 if p1 else p1)) or ((self).f24)
            rhs185_ = (self).fm7(not (p1) or (p1), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dcbbvmlo"))) + ((d_1279_v6_).set(default__.safeIndex(p0, len(d_1279_v6_)), d_1280_v7_))), True, globalState)
            rhs186_ = (d_1281_v8_)[default__.safeIndex((d_1274_v1_).cf9, len(d_1281_v8_))]
            rhs187_ = d_1278_v5_
            lhs174_ = globalState
            lhs175_ = globalState
            lhs176_ = globalState
            lhs174_.f16 = rhs183_
            lhs175_.f0 = rhs184_
            r0 = rhs185_
            lhs176_.f0 = rhs186_
            d_1278_v5_ = rhs187_
            d_1282_v9_: C3
            nw200_ = C3()
            nw200_.ctor__()
            d_1282_v9_ = nw200_
        elif True:
            d_1283_v10_: _dafny.Array
            nw201_ = _dafny.Array(None, 8)
            nw201_[int(0)] = p2
            nw201_[int(1)] = p2
            nw201_[int(2)] = p2
            nw201_[int(3)] = p2
            nw201_[int(4)] = p2
            nw201_[int(5)] = p2
            nw201_[int(6)] = p2
            nw201_[int(7)] = p2
            d_1283_v10_ = nw201_
            d_1283_v10_ = d_1283_v10_
            index202_ = default__.safeIndex(572, (p2).length(0))
            (p2)[index202_] = p3
            index203_ = default__.safeIndex(572, (p2).length(0))
            rhs188_ = (p3) - (p0)
            rhs189_ = p3
            rhs190_ = p3
            lhs177_ = globalState
            lhs178_ = globalState
            lhs179_ = p2
            lhs180_ = default__.safeIndex(572, (p2).length(0))
            lhs177_.f16 = rhs188_
            lhs178_.f16 = rhs189_
            lhs179_[lhs180_] = rhs190_
            d_1284_v11_: str
            d_1284_v11_ = _dafny.CodePoint('f')
            d_1284_v11_ = _dafny.CodePoint('v')
            index204_ = default__.safeIndex(572, (p2).length(0))
            (p2)[index204_] = p3
            d_1285_v12_: _dafny.Array
            def lambda66_(d_1286_p1_, d_1287_p3_):
                def lambda67_(d_1288_i1_):
                    return D12_DC22(_dafny.Map({d_1286_p1_: d_1287_p3_}))

                return lambda67_

            init36_ = lambda66_(p1, p3)
            nw202_ = _dafny.Array(None, 11)
            for i0_36_ in range(nw202_.length(0)):
                nw202_[i0_36_] = init36_(i0_36_)
            d_1285_v12_ = nw202_
            d_1289_v13_: _dafny.Map
            d_1289_v13_ = _dafny.Map({(p2)[default__.safeIndex(572, (p2).length(0))]: d_1285_v12_})
            d_1289_v13_ = (d_1289_v13_).set(p0, d_1285_v12_)
        if not(p1):
            d_1290_v14_: _dafny.Seq
            d_1290_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gus"))
            (globalState).f19 = ((d_1290_v14_) + (d_1290_v14_)) + (d_1290_v14_)
            d_1291_v15_: _dafny.Array
            def lambda68_(d_1292_p0_, d_1293_p3_):
                def lambda69_(d_1294_i2_):
                    return _dafny.SeqWithoutIsStrInference([d_1292_p0_, d_1292_p0_, d_1292_p0_, -190, d_1293_p3_])

                return lambda69_

            init37_ = lambda68_(p0, p3)
            nw203_ = _dafny.Array(None, 21)
            for i0_37_ in range(nw203_.length(0)):
                nw203_[i0_37_] = init37_(i0_37_)
            d_1291_v15_ = nw203_
            d_1295_v16_: D15
            d_1295_v16_ = D15_DC32(d_1291_v15_)
            d_1296_v17_: D15
            d_1296_v17_ = D15_DC34(d_1295_v16_)
            d_1297_v18_: _dafny.Map
            d_1297_v18_ = _dafny.Map({p1: d_1296_v17_})
            d_1297_v18_ = (d_1297_v18_).set(p1, d_1296_v17_)
            d_1298_v19_: str
            d_1298_v19_ = _dafny.CodePoint('n')
            d_1299_v20_: _dafny.Map
            d_1299_v20_ = _dafny.Map({len((d_1290_v14_) + ((d_1290_v14_).set(default__.safeIndex(p3, len(d_1290_v14_)), d_1298_v19_))): p3})
            d_1300_v21_: _dafny.Seq
            d_1300_v21_ = _dafny.SeqWithoutIsStrInference([d_1299_v20_])
            d_1299_v20_ = ((d_1299_v20_) | ((d_1300_v21_)[default__.safeIndex(p0, len(d_1300_v21_))])) | (d_1299_v20_)
            r1 = p0
            (globalState).f13 = p1
        elif True:
            index205_ = default__.safeIndex(96, (p2).length(0))
            (p2)[index205_] = p3
            d_1301_v22_: _dafny.Set
            d_1301_v22_ = _dafny.Set({p2, p2, p2})
            index206_ = default__.safeIndex(96, (p2).length(0))
            (p2)[index206_] = len(d_1301_v22_)
            d_1302_v23_: _dafny.Seq
            d_1302_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "coillmrw"))
            d_1303_v24_: _dafny.MultiSet
            d_1303_v24_ = _dafny.MultiSet([d_1302_v23_])
            d_1304_v25_: _dafny.Map
            d_1304_v25_ = _dafny.Map({p0: _dafny.CodePoint('l')})
            d_1305_v26_: str
            d_1305_v26_ = _dafny.CodePoint('k')
            d_1306_v27_: _dafny.MultiSet
            d_1306_v27_ = _dafny.MultiSet([((d_1303_v24_).cardinality if not((self).f24) else len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_1307_i3_ in range(default__.abs(183))])).set(default__.safeIndex((self).fm8((self).f24, (d_1304_v25_).set(p3, d_1305_v26_), d_1302_v23_, globalState), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_1308_i3_ in range(default__.abs(183))]))), default__.fm27((self).f24, not(False), globalState))))])
            d_1309_v28_: _dafny.Map
            d_1309_v28_ = _dafny.Map({(p2)[default__.safeIndex(96, (p2).length(0))]: p0})
            d_1310_v29_: _dafny.Map
            d_1310_v29_ = _dafny.Map({p1: (self).f24})
            index207_ = default__.safeIndex(96, (p2).length(0))
            (p2)[index207_] = ((d_1306_v27_)[p3] if (p3) in (d_1306_v27_) else ((d_1309_v28_)[881] if (881) in (d_1309_v28_) else len(d_1310_v29_)))
            d_1311_v30_: _dafny.Array
            def lambda70_(d_1312_p0_, d_1313_p2_):
                def lambda71_(d_1314_i4_):
                    return ((0) - (d_1312_p0_)) <= ((d_1313_p2_)[default__.safeIndex(96, (d_1313_p2_).length(0))])

                return lambda71_

            init38_ = lambda70_(p0, p2)
            nw204_ = _dafny.Array(None, 28)
            for i0_38_ in range(nw204_.length(0)):
                nw204_[i0_38_] = init38_(i0_38_)
            d_1311_v30_ = nw204_
            index208_ = default__.safeIndex(792, (d_1311_v30_).length(0))
            (d_1311_v30_)[index208_] = p1
            d_1315_v31_: _dafny.Map
            d_1315_v31_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1305_v26_ for d_1316_i5_ in range(default__.abs(488))]): p0})
            d_1317_v32_: C4
            nw205_ = C4()
            nw205_.ctor__(False, p1)
            d_1317_v32_ = nw205_
            d_1318_v33_: _dafny.Seq
            d_1318_v33_ = _dafny.SeqWithoutIsStrInference([(d_1317_v32_).f26])
            d_1319_v34_: _dafny.Seq
            d_1319_v34_ = _dafny.SeqWithoutIsStrInference([p1, (d_1318_v33_)[default__.safeIndex(p0, len(d_1318_v33_))], p1])
            d_1320_v35_: _dafny.Seq
            d_1320_v35_ = _dafny.SeqWithoutIsStrInference([p1, d_1317_v32_.f25])
            d_1321_v36_: _dafny.Map
            d_1321_v36_ = _dafny.Map({d_1320_v35_: d_1305_v26_})
            index209_ = default__.safeIndex(792, (d_1311_v30_).length(0))
            rhs191_ = True
            rhs192_ = d_1311_v30_
            rhs193_ = ((len(d_1315_v31_)) - ((0) - (len(_dafny.Set({d_1317_v32_, d_1317_v32_, d_1317_v32_, d_1317_v32_})))) if p1 else default__.safeDivisionInt(len(d_1318_v33_), len(_dafny.SeqWithoutIsStrInference([p1]))))
            rhs194_ = ((d_1321_v36_)[default__.fm41(not(True), _dafny.Set({default__.fm2(d_1310_v29_, globalState), True, (d_1317_v32_).f26, d_1317_v32_.f25}), (p2)[default__.safeIndex(96, (p2).length(0))], p3, globalState)] if (default__.fm41(not(True), _dafny.Set({default__.fm2(d_1310_v29_, globalState), True, (d_1317_v32_).f26, d_1317_v32_.f25}), (p2)[default__.safeIndex(96, (p2).length(0))], p3, globalState)) in (d_1321_v36_) else d_1305_v26_)
            lhs181_ = d_1311_v30_
            lhs182_ = default__.safeIndex(792, (d_1311_v30_).length(0))
            lhs183_ = globalState
            lhs181_[lhs182_] = rhs191_
            d_1311_v30_ = rhs192_
            lhs183_.f7 = rhs193_
            d_1305_v26_ = rhs194_
            (globalState).f8 = d_1302_v23_
            d_1322_v37_: _dafny.Array
            def lambda72_(d_1323_v32_):
                def lambda73_(d_1324_i6_):
                    return _dafny.Map({(self).f24: d_1323_v32_.f25})

                return lambda73_

            init39_ = lambda72_(d_1317_v32_)
            nw206_ = _dafny.Array(None, 23)
            for i0_39_ in range(nw206_.length(0)):
                nw206_[i0_39_] = init39_(i0_39_)
            d_1322_v37_ = nw206_
            d_1322_v37_ = (d_1322_v37_ if (d_1317_v32_).f26 else d_1322_v37_)
        if not (not(p1)) or (not(p1)):
            d_1325_v38_: str
            d_1325_v38_ = _dafny.CodePoint('m')
            d_1325_v38_ = d_1325_v38_
            d_1326_v39_: _dafny.Seq
            d_1326_v39_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))
            index210_ = default__.safeIndex(453, (d_1272_v0_).length(0))
            (d_1272_v0_)[index210_] = d_1326_v39_
            index211_ = default__.safeIndex(453, (d_1272_v0_).length(0))
            (d_1272_v0_)[index211_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "woo"))
            d_1327_v40_: _dafny.Array
            nw207_ = _dafny.Array(_dafny.Array(None, 0), 29)
            d_1327_v40_ = nw207_
            d_1327_v40_ = d_1327_v40_
            d_1328_v41_: _dafny.Map
            d_1328_v41_ = _dafny.Map({d_1325_v38_: p0})
            (globalState).f19 = ((d_1326_v39_).set(default__.safeIndex(((d_1328_v41_)[d_1325_v38_] if (d_1325_v38_) in (d_1328_v41_) else (0) - (p3)), len(d_1326_v39_)), d_1325_v38_)).set(default__.safeIndex(p3, len((d_1326_v39_).set(default__.safeIndex(((d_1328_v41_)[d_1325_v38_] if (d_1325_v38_) in (d_1328_v41_) else (0) - (p3)), len(d_1326_v39_)), d_1325_v38_))), d_1325_v38_)
            (globalState).f13 = (self).f24
        elif True:
            d_1329_v42_: _dafny.Array
            nw208_ = _dafny.Array(_dafny.Array(None, 0), 1)
            d_1329_v42_ = nw208_
            index212_ = default__.safeIndex(201, (d_1329_v42_).length(0))
            (d_1329_v42_)[index212_] = p2
            index213_ = default__.safeIndex(609, (p2).length(0))
            (p2)[index213_] = p0
            d_1330_v43_: _dafny.Seq
            d_1330_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krdj"))
            index214_ = default__.safeIndex(201, (d_1329_v42_).length(0))
            index215_ = default__.safeIndex(609, (p2).length(0))
            rhs195_ = (0) - (p0)
            rhs196_ = p2
            rhs197_ = (len(d_1330_v43_) if p1 else p3)
            lhs184_ = d_1329_v42_
            lhs185_ = default__.safeIndex(201, (d_1329_v42_).length(0))
            lhs186_ = p2
            lhs187_ = default__.safeIndex(609, (p2).length(0))
            r1 = rhs195_
            lhs184_[lhs185_] = rhs196_
            lhs186_[lhs187_] = rhs197_
            d_1331_v44_: _dafny.Array
            def lambda74_(d_1332_p1_):
                def lambda75_(d_1333_i7_):
                    return d_1332_p1_

                return lambda75_

            init40_ = lambda74_(p1)
            nw209_ = _dafny.Array(None, 11)
            for i0_40_ in range(nw209_.length(0)):
                nw209_[i0_40_] = init40_(i0_40_)
            d_1331_v44_ = nw209_
            d_1334_v45_: _dafny.Array
            nw210_ = _dafny.Array(None, 11)
            nw210_[int(0)] = d_1331_v44_
            nw210_[int(1)] = d_1331_v44_
            nw210_[int(2)] = d_1331_v44_
            nw210_[int(3)] = d_1331_v44_
            nw210_[int(4)] = d_1331_v44_
            nw210_[int(5)] = d_1331_v44_
            nw210_[int(6)] = d_1331_v44_
            nw210_[int(7)] = d_1331_v44_
            nw210_[int(8)] = d_1331_v44_
            nw210_[int(9)] = d_1331_v44_
            nw210_[int(10)] = d_1331_v44_
            d_1334_v45_ = nw210_
            d_1335_v46_: _dafny.Array
            d_1335_v46_ = d_1331_v44_
            index216_ = default__.safeIndex(879, (d_1334_v45_).length(0))
            (d_1334_v45_)[index216_] = (d_1335_v46_)
            d_1336_v47_: _dafny.Map
            d_1336_v47_ = _dafny.Map({p1: (self).f24})
            d_1337_v48_: D18
            d_1337_v48_ = D18_DC42((d_1336_v47_).set(p1, (self).f24))
            d_1338_v49_: _dafny.Seq
            d_1338_v49_ = _dafny.SeqWithoutIsStrInference([(self).f24, p1, True, p1])
            d_1339_v50_: _dafny.Seq
            d_1339_v50_ = _dafny.SeqWithoutIsStrInference([(d_1336_v47_) | ((d_1337_v48_).cf58), d_1336_v47_, (_dafny.Map({p1: (self).f24})).set((default__.fm25(True, p0, p1, len(d_1338_v49_), globalState)).cf8, (self).f24)])
            index217_ = default__.safeIndex(879, (d_1334_v45_).length(0))
            rhs198_ = (d_1331_v44_ if p1 else d_1331_v44_)
            rhs199_ = p1
            rhs200_ = p1
            rhs201_ = not((d_1330_v43_) == (d_1330_v43_))
            rhs202_ = (d_1339_v50_)[default__.safeIndex(p3, len(d_1339_v50_))]
            lhs188_ = d_1334_v45_
            lhs189_ = default__.safeIndex(879, (d_1334_v45_).length(0))
            lhs190_ = globalState
            lhs191_ = globalState
            lhs192_ = globalState
            lhs188_[lhs189_] = rhs198_
            lhs190_.f13 = rhs199_
            lhs191_.f13 = rhs200_
            lhs192_.f2 = rhs201_
            d_1336_v47_ = rhs202_
            (globalState).f0 = (self).fm6(p3, globalState)
            d_1340_v51_: _dafny.Map
            d_1340_v51_ = _dafny.Map({False: (p2)[default__.safeIndex(609, (p2).length(0))]})
            d_1341_v52_: D4
            d_1341_v52_ = D4_DC10(len(d_1340_v51_), (p2)[default__.safeIndex(609, (p2).length(0))], (self).f24, p1, (p2)[default__.safeIndex(609, (p2).length(0))])
            d_1342_v53_: _dafny.Seq
            d_1342_v53_ = _dafny.SeqWithoutIsStrInference([d_1341_v52_, default__.fm42(globalState)])
            d_1343_v54_: _dafny.MultiSet
            d_1343_v54_ = _dafny.MultiSet([p1])
            d_1344_v55_: D4
            d_1344_v55_ = D4_DC10(len(d_1342_v53_), p3, p1, (D16_DC36(not(p1), d_1343_v54_, p1)).cf49, (p2)[default__.safeIndex(609, (p2).length(0))])
            d_1345_v56_: _dafny.Map
            d_1345_v56_ = _dafny.Map({(self).f24: d_1344_v55_})
            d_1346_v57_: _dafny.Map
            d_1346_v57_ = _dafny.Map({p0: d_1345_v56_})
            d_1346_v57_ = _dafny.Map({p0: default__.fm43((self).f24, p3, len(d_1330_v43_), p1, globalState)})
            d_1347_v58_: _dafny.Map
            d_1347_v58_ = _dafny.Map({(self).f24: (d_1329_v42_)[default__.safeIndex(201, (d_1329_v42_).length(0))]})
            d_1347_v58_ = (d_1347_v58_).set((self).f24, (d_1329_v42_)[default__.safeIndex(201, (d_1329_v42_).length(0))])
        (globalState).f16 = p3
        r0 = not(((False if (self).f24 else p1)) == ((self).f24))
        r1 = (0) - (p0)
        def lambda76_(d_1348_i8_):
            return _dafny.Map({(self).f24: D3_DC6(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ysrfvc")), _dafny.CodePoint('w'), False)})

        init41_ = lambda76_
        nw211_ = _dafny.Array(None, 6)
        for i0_41_ in range(nw211_.length(0)):
            nw211_[i0_41_] = init41_(i0_41_)
        r2 = nw211_
        return r0, r1, r2

    def m3(self, p0, globalState):
        d_1349_v0_: int
        d_1349_v0_ = 48
        d_1350_v1_: _dafny.Set
        d_1350_v1_ = _dafny.Set({d_1349_v0_})
        d_1351_v2_: C4
        nw212_ = C4()
        nw212_.ctor__((d_1350_v1_) != (d_1350_v1_), (self).f24)
        d_1351_v2_ = nw212_
        d_1352_v3_: _dafny.MultiSet
        d_1352_v3_ = _dafny.MultiSet([(self).f24])
        d_1353_v4_: D16
        d_1353_v4_ = D16_DC36((d_1351_v2_).f26, d_1352_v3_, (d_1351_v2_).f26)
        source21_ = d_1353_v4_
        if source21_.is_DC36:
            d_1354___mcc_h0_ = source21_.cf47
            d_1355___mcc_h1_ = source21_.cf48
            d_1356___mcc_h2_ = source21_.cf49
            d_1357_cf49_ = d_1356___mcc_h2_
            d_1358_cf48_ = d_1355___mcc_h1_
            d_1359_cf47_ = d_1354___mcc_h0_
            (globalState).f16 = d_1349_v0_
            d_1360_v5_: _dafny.Seq
            d_1360_v5_ = _dafny.SeqWithoutIsStrInference([d_1357_cf49_, d_1359_cf47_])
            d_1358_cf48_ = ((d_1352_v3_ if d_1359_cf47_ else d_1358_cf48_)) - (_dafny.MultiSet(d_1360_v5_))
            d_1361_v6_: C0
            nw213_ = C0()
            nw213_.ctor__((d_1360_v5_)[default__.safeIndex(d_1349_v0_, len(d_1360_v5_))])
            d_1361_v6_ = nw213_
            d_1362_v7_: C0
            nw214_ = C0()
            nw214_.ctor__((d_1349_v0_) <= (74))
            d_1362_v7_ = nw214_
        elif source21_.is_DC35:
            d_1363___mcc_h3_ = source21_.cf46
            d_1364_cf46_ = d_1363___mcc_h3_
            d_1365_v8_: _dafny.Array
            nw215_ = _dafny.Array(D0.default()(), 13)
            d_1365_v8_ = nw215_
            d_1366_v9_: D0
            d_1366_v9_ = D0_DC1(default__.fm4(140, d_1349_v0_, globalState))
            d_1367_v10_: _dafny.Seq
            d_1367_v10_ = _dafny.SeqWithoutIsStrInference([d_1366_v9_, d_1366_v9_, D0_DC0(d_1351_v2_.f25)])
            d_1368_v11_: D0
            d_1368_v11_ = D0_DC1((d_1367_v10_)[default__.safeIndex(d_1349_v0_, len(d_1367_v10_))])
            index218_ = default__.safeIndex(212, (d_1365_v8_).length(0))
            (d_1365_v8_)[index218_] = d_1368_v11_
            pat_let_tv30_ = d_1366_v9_
            index219_ = default__.safeIndex(212, (d_1365_v8_).length(0))
            def iife102_(_pat_let22_0):
                def iife103_(d_1369_dt__update__tmp_h0_):
                    def iife104_(_pat_let23_0):
                        def iife105_(d_1370_dt__update_hcf1_h0_):
                            return D0_DC1(d_1370_dt__update_hcf1_h0_)
                        return iife105_(_pat_let23_0)
                    return iife104_(pat_let_tv30_)
                return iife103_(_pat_let22_0)
            (d_1365_v8_)[index219_] = iife102_(d_1368_v11_)
            d_1371_v12_: _dafny.Array
            def lambda77_(d_1372_v0_):
                def lambda78_(d_1373_i0_):
                    return _dafny.SeqWithoutIsStrInference([d_1372_v0_])

                return lambda78_

            init42_ = lambda77_(d_1349_v0_)
            nw216_ = _dafny.Array(None, 5)
            for i0_42_ in range(nw216_.length(0)):
                nw216_[i0_42_] = init42_(i0_42_)
            d_1371_v12_ = nw216_
            d_1374_v13_: D15
            d_1374_v13_ = D15_DC32(d_1371_v12_)
            d_1375_v14_: _dafny.MultiSet
            d_1375_v14_ = _dafny.MultiSet([d_1374_v13_])
            if not(((d_1375_v14_).set(d_1374_v13_, default__.abs(d_1349_v0_))) != (d_1375_v14_)):
                d_1376_v16_: _dafny.Seq
                d_1376_v16_ = _dafny.SeqWithoutIsStrInference([(d_1351_v2_).f26, not((d_1351_v2_).f26)])
                d_1377_v17_: _dafny.Seq
                d_1377_v17_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1378_v18_: _dafny.Array
                nw217_ = _dafny.Array(None, 20)
                nw217_[int(0)] = d_1351_v2_.f25
                nw217_[int(1)] = ((d_1351_v2_).f26) == ((self).f24)
                nw217_[int(2)] = d_1351_v2_.f25
                nw217_[int(3)] = (d_1351_v2_).f26
                nw217_[int(4)] = (d_1351_v2_).f26
                def iife106_():
                    coll58_ = _dafny.Set()
                    compr_58_: int
                    for compr_58_ in _dafny.IntegerRange(100, -738):
                        d_1379_v15_: int = compr_58_
                        if ((100) <= (d_1379_v15_)) and ((d_1379_v15_) < (-738)):
                            coll58_ = coll58_.union(_dafny.Set([(d_1379_v15_) - (d_1349_v0_)]))
                    return _dafny.Set(coll58_)
                nw217_[int(5)] = (iife106_()
                ).issubset(d_1350_v1_)
                nw217_[int(6)] = d_1351_v2_.f25
                nw217_[int(7)] = d_1351_v2_.f25
                nw217_[int(8)] = not(((d_1351_v2_).f26 if (self).fm7((d_1351_v2_).f26, d_1349_v0_, d_1351_v2_.f25, globalState) else (d_1351_v2_).f26))
                nw217_[int(9)] = (d_1351_v2_).f26
                nw217_[int(10)] = (d_1351_v2_).f26
                nw217_[int(11)] = (d_1376_v16_) < ((d_1376_v16_).set(default__.safeIndex(d_1349_v0_, len(d_1376_v16_)), False))
                nw217_[int(12)] = (d_1351_v2_).f26
                nw217_[int(13)] = (d_1377_v17_) <= (d_1377_v17_)
                nw217_[int(14)] = (d_1351_v2_).f26
                nw217_[int(15)] = d_1351_v2_.f25
                nw217_[int(16)] = (d_1351_v2_).f26
                nw217_[int(17)] = d_1351_v2_.f25
                nw217_[int(18)] = (d_1351_v2_).f26
                nw217_[int(19)] = (d_1351_v2_).f26
                d_1378_v18_ = nw217_
                d_1378_v18_ = d_1378_v18_
                d_1380_v19_: _dafny.Seq
                d_1380_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lb"))
                (globalState).f20 = d_1380_v19_
                (globalState).f7 = d_1349_v0_
                d_1381_v20_: _dafny.MultiSet
                d_1381_v20_ = _dafny.MultiSet([d_1349_v0_])
                d_1382_v21_: _dafny.Seq
                d_1382_v21_ = _dafny.SeqWithoutIsStrInference([(d_1352_v3_).cardinality, d_1349_v0_])
                (globalState).f16 = (default__.fm0((d_1381_v20_).cardinality, d_1349_v0_, len(d_1382_v21_), (d_1351_v2_).f26, globalState)) * (d_1349_v0_)
                d_1383_v22_: C4
                nw218_ = C4()
                nw218_.ctor__((d_1351_v2_).f26, not((d_1351_v2_).f26))
                d_1383_v22_ = nw218_
            elif True:
                (d_1351_v2_).m5(d_1351_v2_.f25, globalState)
                d_1384_v23_: _dafny.Seq
                d_1384_v23_ = _dafny.SeqWithoutIsStrInference([d_1351_v2_.f25])
                d_1385_v24_: _dafny.Set
                d_1385_v24_ = _dafny.Set({d_1351_v2_.f25, (self).f24, d_1351_v2_.f25})
                d_1386_v25_: _dafny.Array
                nw219_ = _dafny.Array(None, 14)
                nw219_[int(0)] = d_1351_v2_.f25
                nw219_[int(1)] = (d_1384_v23_) < (d_1384_v23_)
                nw219_[int(2)] = False
                nw219_[int(3)] = ((d_1351_v2_).f26) not in (d_1384_v23_)
                nw219_[int(4)] = not((d_1351_v2_).f26)
                nw219_[int(5)] = (d_1351_v2_).f26
                nw219_[int(6)] = (d_1351_v2_).f26
                nw219_[int(7)] = (d_1351_v2_).f26
                nw219_[int(8)] = ((d_1351_v2_).f26 if d_1351_v2_.f25 else True)
                nw219_[int(9)] = d_1351_v2_.f25
                nw219_[int(10)] = (d_1385_v24_) == (d_1385_v24_)
                nw219_[int(11)] = (d_1349_v0_) != (d_1349_v0_)
                nw219_[int(12)] = False
                nw219_[int(13)] = (self).f24
                d_1386_v25_ = nw219_
                index220_ = default__.safeIndex(438, (d_1386_v25_).length(0))
                (d_1386_v25_)[index220_] = (self).f24
                d_1387_v26_: _dafny.Seq
                d_1387_v26_ = _dafny.SeqWithoutIsStrInference([d_1349_v0_])
                index221_ = default__.safeIndex(438, (d_1386_v25_).length(0))
                (d_1386_v25_)[index221_] = (((0) - (d_1349_v0_)) == (d_1349_v0_)) or ((d_1387_v26_) <= (d_1387_v26_))
                d_1388_v27_: _dafny.Map
                d_1388_v27_ = _dafny.Map({d_1349_v0_: (self).f24})
                d_1389_v28_: _dafny.Map
                d_1389_v28_ = _dafny.Map({(d_1351_v2_).f26: False})
                rhs203_ = _dafny.Map({len((_dafny.Set({d_1351_v2_.f25})).intersection(d_1385_v24_)): (d_1386_v25_)[default__.safeIndex(438, (d_1386_v25_).length(0))]})
                rhs204_ = ((d_1389_v28_).set(d_1351_v2_.f25, (d_1351_v2_).f26)) | ((d_1389_v28_) | (d_1389_v28_))
                rhs205_ = (d_1349_v0_) - (155)
                rhs206_ = d_1349_v0_
                rhs207_ = (d_1351_v2_).f26
                lhs193_ = globalState
                lhs194_ = globalState
                lhs195_ = globalState
                d_1388_v27_ = rhs203_
                d_1389_v28_ = rhs204_
                lhs193_.f16 = rhs205_
                lhs194_.f7 = rhs206_
                lhs195_.f2 = rhs207_
                d_1390_v29_: _dafny.Seq
                d_1390_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "udixpjur"))
                d_1391_v30_: str
                d_1391_v30_ = _dafny.CodePoint('m')
                (globalState).f8 = (d_1390_v29_).set(default__.safeIndex(len(d_1385_v24_), len(d_1390_v29_)), d_1391_v30_)
                d_1392_v31_: _dafny.Map
                d_1392_v31_ = _dafny.Map({default__.fm32((d_1386_v25_)[default__.safeIndex(438, (d_1386_v25_).length(0))], d_1349_v0_, len(d_1387_v26_), globalState): d_1351_v2_.f25})
                d_1393_v32_: D14
                d_1393_v32_ = D14_DC28(_dafny.Set({717, d_1349_v0_}))
                d_1392_v31_ = (d_1392_v31_).set((d_1393_v32_ if (d_1386_v25_)[default__.safeIndex(438, (d_1386_v25_).length(0))] else d_1393_v32_), (default__.fm44(globalState)).cf36)
            index222_ = default__.safeIndex(774, (p0).length(0))
            (p0)[index222_] = d_1349_v0_
            index223_ = default__.safeIndex(774, (p0).length(0))
            (p0)[index223_] = d_1349_v0_
            d_1394_v33_: _dafny.Map
            d_1394_v33_ = _dafny.Map({(d_1351_v2_).f26: 262})
            d_1395_v34_: _dafny.Seq
            d_1395_v34_ = _dafny.SeqWithoutIsStrInference([d_1394_v33_])
            d_1396_v35_: _dafny.Seq
            d_1396_v35_ = _dafny.SeqWithoutIsStrInference([True])
            d_1397_v36_: _dafny.Set
            d_1397_v36_ = _dafny.Set({d_1351_v2_.f25, (d_1351_v2_).f26})
            d_1398_v37_: _dafny.Array
            nw220_ = _dafny.Array(None, 9)
            nw220_[int(0)] = (d_1351_v2_).f26
            nw220_[int(1)] = d_1351_v2_.f25
            nw220_[int(2)] = (_dafny.SeqWithoutIsStrInference([d_1394_v33_])) < (d_1395_v34_)
            nw220_[int(3)] = d_1351_v2_.f25
            nw220_[int(4)] = d_1351_v2_.f25
            nw220_[int(5)] = (d_1351_v2_.f25) or (d_1351_v2_.f25)
            nw220_[int(6)] = (self).f24
            nw220_[int(7)] = not ((d_1351_v2_).f26) or ((d_1351_v2_).f26)
            nw220_[int(8)] = (len(d_1396_v35_)) != (len(d_1397_v36_))
            d_1398_v37_ = nw220_
            index224_ = default__.safeIndex(566, (d_1398_v37_).length(0))
            (d_1398_v37_)[index224_] = d_1351_v2_.f25
            index225_ = default__.safeIndex(566, (d_1398_v37_).length(0))
            (d_1398_v37_)[index225_] = (d_1351_v2_).f26
        elif True:
            d_1399___mcc_h4_ = source21_.cf50
            d_1400_cf50_ = d_1399___mcc_h4_
            d_1401_v38_: _dafny.Seq
            d_1401_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vuuysgyyb"))
            d_1402_v39_: _dafny.Seq
            d_1402_v39_ = _dafny.SeqWithoutIsStrInference([d_1351_v2_.f25])
            d_1403_v40_: _dafny.MultiSet
            d_1403_v40_ = _dafny.MultiSet([(len(d_1401_v38_)) + (((d_1352_v3_)[d_1351_v2_.f25] if (d_1351_v2_.f25) in (d_1352_v3_) else d_1349_v0_)), default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([d_1402_v39_])), len(d_1401_v38_))])
            d_1403_v40_ = _dafny.MultiSet([d_1349_v0_, (d_1349_v0_) * (306), default__.fm0(d_1349_v0_, (0) - (d_1349_v0_), d_1349_v0_, not(d_1351_v2_.f25), globalState), 241])
            (globalState).f13 = d_1351_v2_.f25
            (globalState).f13 = (d_1351_v2_).f26
            nw221_ = C4()
            nw221_.ctor__((d_1351_v2_).f26, d_1351_v2_.f25)
            d_1351_v2_ = nw221_
        d_1404_v41_: _dafny.Map
        d_1404_v41_ = _dafny.Map({(d_1351_v2_).f26: (self).f24})
        if default__.fm2(d_1404_v41_, globalState):
            d_1405_v42_: _dafny.Array
            def lambda79_(d_1406_i1_):
                return _dafny.CodePoint('a')

            init43_ = lambda79_
            nw222_ = _dafny.Array(None, 6)
            for i0_43_ in range(nw222_.length(0)):
                nw222_[i0_43_] = init43_(i0_43_)
            d_1405_v42_ = nw222_
            d_1405_v42_ = d_1405_v42_
            d_1407_v43_: _dafny.MultiSet
            d_1407_v43_ = _dafny.MultiSet([d_1349_v0_, d_1349_v0_])
            if ((d_1349_v0_ if (d_1351_v2_).f26 else ((d_1407_v43_)[d_1349_v0_] if (d_1349_v0_) in (d_1407_v43_) else d_1349_v0_))) <= (161):
                d_1408_v44_: _dafny.Seq
                d_1408_v44_ = _dafny.SeqWithoutIsStrInference([d_1352_v3_])
                (globalState).f16 = default__.safeDivisionInt(len(d_1408_v44_), d_1349_v0_)
                (d_1351_v2_).f25 = d_1351_v2_.f25
                d_1409_v45_: _dafny.Map
                d_1409_v45_ = _dafny.Map({default__.safeDivisionInt(d_1349_v0_, 380): d_1349_v0_})
                d_1410_v46_: _dafny.MultiSet
                d_1410_v46_ = _dafny.MultiSet([p0])
                d_1409_v45_ = (d_1409_v45_).set(d_1349_v0_, (d_1349_v0_) - ((d_1410_v46_).cardinality))
                d_1411_v47_: _dafny.Array
                def lambda80_(d_1412_i2_):
                    return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hkapgn"))

                init44_ = lambda80_
                nw223_ = _dafny.Array(None, 17)
                for i0_44_ in range(nw223_.length(0)):
                    nw223_[i0_44_] = init44_(i0_44_)
                d_1411_v47_ = nw223_
                d_1413_v48_: str
                d_1413_v48_ = _dafny.CodePoint('r')
                d_1414_v49_: _dafny.Map
                d_1414_v49_ = _dafny.Map({d_1411_v47_: d_1413_v48_})
                d_1414_v49_ = (d_1414_v49_).set(d_1411_v47_, d_1413_v48_)
                index226_ = default__.safeIndex(102, (p0).length(0))
                (p0)[index226_] = -806
                index227_ = default__.safeIndex(102, (p0).length(0))
                (p0)[index227_] = d_1349_v0_
            elif True:
                d_1415_v50_: _dafny.Map
                d_1415_v50_ = _dafny.Map({d_1349_v0_: (d_1351_v2_).f26})
                d_1416_v52_: _dafny.Seq
                def iife107_():
                    coll59_ = _dafny.Map()
                    compr_59_: int
                    for compr_59_ in (d_1350_v1_).Elements:
                        d_1417_v51_: int = compr_59_
                        if (d_1417_v51_) in (d_1350_v1_):
                            coll59_[(d_1417_v51_) + (len(_dafny.SeqWithoutIsStrInference([d_1351_v2_.f25])))] = d_1351_v2_.f25
                    return _dafny.Map(coll59_)
                d_1416_v52_ = _dafny.SeqWithoutIsStrInference([d_1415_v50_, d_1415_v50_, iife107_()
                ])
                d_1418_v53_: C1
                nw224_ = C1()
                nw224_.ctor__((d_1349_v0_) * (d_1349_v0_), (_dafny.SeqWithoutIsStrInference([d_1415_v50_, d_1415_v50_, d_1415_v50_]) if d_1351_v2_.f25 else d_1416_v52_))
                d_1418_v53_ = nw224_
                d_1419_v54_: _dafny.Seq
                d_1419_v54_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1419_v54_ = d_1419_v54_
                d_1352_v3_ = d_1352_v3_
                (globalState).f0 = (d_1351_v2_).f26
                d_1420_v55_: C2
                nw225_ = C2()
                nw225_.ctor__(d_1349_v0_)
                d_1420_v55_ = nw225_
            source22_ = default__.fm45(globalState)
            if source22_.is_DC26:
                d_1421___mcc_h5_ = source22_.cf33
                d_1422___mcc_h6_ = source22_.cf34
                d_1423___mcc_h7_ = source22_.cf35
                d_1424_cf35_ = d_1423___mcc_h7_
                d_1425_cf34_ = d_1422___mcc_h6_
                d_1426_cf33_ = d_1421___mcc_h5_
                (globalState).f16 = d_1349_v0_
                (globalState).f13 = (d_1351_v2_).f26
                d_1427_v56_: str
                d_1427_v56_ = _dafny.CodePoint('w')
                d_1427_v56_ = (d_1426_cf33_)[default__.safeIndex((0) - (d_1349_v0_), len(d_1426_cf33_))]
                d_1428_v57_: _dafny.Seq
                d_1428_v57_ = _dafny.SeqWithoutIsStrInference([(305) * (d_1349_v0_)])
                d_1428_v57_ = (d_1428_v57_) + ((d_1428_v57_) + (d_1428_v57_))
            elif source22_.is_DC27:
                d_1429___mcc_h8_ = source22_.cf36
                d_1430_cf36_ = d_1429___mcc_h8_
                (globalState).f0 = not(((d_1351_v2_).f26) in (d_1352_v3_))
                d_1431_v58_: _dafny.Array
                nw226_ = _dafny.Array(False, 20)
                d_1431_v58_ = nw226_
                d_1432_v59_: _dafny.MultiSet
                d_1432_v59_ = _dafny.MultiSet([d_1431_v58_, d_1431_v58_, d_1431_v58_])
                d_1433_v60_: _dafny.Seq
                d_1433_v60_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_1431_v58_])])
                d_1432_v59_ = (d_1433_v60_)[default__.safeIndex((d_1349_v0_ if d_1351_v2_.f25 else d_1349_v0_), len(d_1433_v60_))]
                d_1434_v61_: _dafny.Map
                d_1434_v61_ = _dafny.Map({d_1349_v0_: (d_1351_v2_).f26})
                d_1435_v62_: _dafny.Seq
                d_1435_v62_ = _dafny.SeqWithoutIsStrInference([default__.fm2(d_1404_v41_, globalState), ((d_1434_v61_)[d_1349_v0_] if (d_1349_v0_) in (d_1434_v61_) else (self).f24), (d_1349_v0_) < (d_1349_v0_)])
                rhs208_ = (d_1435_v62_)[default__.safeIndex(d_1349_v0_, len(d_1435_v62_))]
                rhs209_ = (0) - (d_1349_v0_)
                rhs210_ = d_1431_v58_
                rhs211_ = d_1351_v2_
                rhs212_ = d_1349_v0_
                lhs196_ = d_1351_v2_
                lhs197_ = globalState
                lhs198_ = globalState
                lhs196_.f25 = rhs208_
                lhs197_.f7 = rhs209_
                d_1431_v58_ = rhs210_
                d_1351_v2_ = rhs211_
                lhs198_.f7 = rhs212_
                d_1436_v63_: str
                d_1436_v63_ = _dafny.CodePoint('c')
                d_1436_v63_ = d_1436_v63_
            elif True:
                d_1437___mcc_h9_ = source22_.cf32
                d_1438_cf32_ = d_1437___mcc_h9_
                d_1439_v64_: _dafny.Array
                nw227_ = _dafny.Array(_dafny.Array(None, 0), 16)
                d_1439_v64_ = nw227_
                d_1439_v64_ = d_1439_v64_
                d_1440_v65_: _dafny.Map
                d_1440_v65_ = _dafny.Map({d_1351_v2_.f25: d_1438_cf32_})
                d_1440_v65_ = (d_1440_v65_).set((d_1351_v2_).f26, d_1438_cf32_)
                d_1441_v66_: _dafny.Array
                nw228_ = _dafny.Array(_dafny.MultiSet({}), 11)
                d_1441_v66_ = nw228_
                rhs213_ = d_1351_v2_.f25
                rhs214_ = d_1441_v66_
                rhs215_ = (d_1351_v2_).f26
                rhs216_ = d_1351_v2_.f25
                lhs199_ = globalState
                lhs200_ = d_1351_v2_
                lhs201_ = globalState
                lhs199_.f2 = rhs213_
                d_1441_v66_ = rhs214_
                lhs200_.f25 = rhs215_
                lhs201_.f0 = rhs216_
                d_1442_v67_: C2
                nw229_ = C2()
                nw229_.ctor__((0) - (d_1349_v0_))
                d_1442_v67_ = nw229_
            d_1443_v68_: str
            d_1443_v68_ = _dafny.CodePoint('q')
            index228_ = default__.safeIndex(175, (d_1405_v42_).length(0))
            (d_1405_v42_)[index228_] = d_1443_v68_
            index229_ = default__.safeIndex(175, (d_1405_v42_).length(0))
            (d_1405_v42_)[index229_] = d_1443_v68_
            d_1444_v69_: _dafny.Seq
            d_1444_v69_ = _dafny.SeqWithoutIsStrInference([(d_1351_v2_).f26])
            d_1445_v70_: D14
            d_1445_v70_ = D14_DC30(_dafny.SeqWithoutIsStrInference([832 for d_1446_i3_ in range(default__.abs(116))]), d_1444_v69_)
            d_1447_v71_: _dafny.Array
            nw230_ = _dafny.Array(None, 20)
            nw230_[int(0)] = d_1444_v69_
            nw230_[int(1)] = d_1444_v69_
            nw230_[int(2)] = default__.fm18(d_1351_v2_.f25, _dafny.CodePoint('p'), (d_1351_v2_).f26, globalState)
            nw230_[int(3)] = _dafny.SeqWithoutIsStrInference([d_1351_v2_.f25])
            nw230_[int(4)] = _dafny.SeqWithoutIsStrInference([True, not(False), (self).f24, (d_1351_v2_).f26, True])
            nw230_[int(5)] = d_1444_v69_
            nw230_[int(6)] = (_dafny.SeqWithoutIsStrInference([d_1351_v2_.f25])) + (d_1444_v69_)
            nw230_[int(7)] = _dafny.SeqWithoutIsStrInference([(self).f24])
            nw230_[int(8)] = (d_1444_v69_) + (d_1444_v69_)
            nw230_[int(9)] = d_1444_v69_
            nw230_[int(10)] = d_1444_v69_
            nw230_[int(11)] = d_1444_v69_
            nw230_[int(12)] = d_1444_v69_
            nw230_[int(13)] = d_1444_v69_
            nw230_[int(14)] = (d_1444_v69_) + (_dafny.SeqWithoutIsStrInference([d_1351_v2_.f25, d_1351_v2_.f25]))
            nw230_[int(15)] = (default__.fm18(default__.fm2(d_1404_v41_, globalState), d_1443_v68_, d_1351_v2_.f25, globalState) if True else (d_1445_v70_).cf41)
            nw230_[int(16)] = _dafny.SeqWithoutIsStrInference([d_1351_v2_.f25])
            nw230_[int(17)] = d_1444_v69_
            nw230_[int(18)] = d_1444_v69_
            nw230_[int(19)] = d_1444_v69_
            d_1447_v71_ = nw230_
            index230_ = default__.safeIndex(1, (d_1447_v71_).length(0))
            (d_1447_v71_)[index230_] = d_1444_v69_
            d_1448_v72_: _dafny.Set
            d_1448_v72_ = _dafny.Set({(d_1351_v2_).f26})
            d_1449_v73_: _dafny.Set
            d_1449_v73_ = (d_1448_v72_) - (_dafny.Set({(d_1351_v2_).f26, d_1351_v2_.f25}))
            d_1450_v74_: _dafny.Seq
            d_1450_v74_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sgpusvt"))
            index231_ = default__.safeIndex(1, (d_1447_v71_).length(0))
            rhs217_ = d_1450_v74_
            rhs218_ = -46
            rhs219_ = d_1444_v69_
            rhs220_ = d_1449_v73_
            rhs221_ = d_1351_v2_.f25
            lhs202_ = globalState
            lhs203_ = globalState
            lhs204_ = d_1447_v71_
            lhs205_ = default__.safeIndex(1, (d_1447_v71_).length(0))
            lhs206_ = globalState
            lhs202_.f20 = rhs217_
            lhs203_.f7 = rhs218_
            lhs204_[lhs205_] = rhs219_
            d_1449_v73_ = rhs220_
            lhs206_.f21 = rhs221_
        elif True:
            d_1451_v75_: _dafny.Seq
            d_1451_v75_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bbe"))
            d_1452_v76_: _dafny.Map
            d_1452_v76_ = _dafny.Map({(d_1451_v75_) != (d_1451_v75_): d_1349_v0_})
            rhs222_ = default__.safeModuloInt(d_1349_v0_, d_1349_v0_)
            rhs223_ = (d_1351_v2_).f26
            rhs224_ = (d_1452_v76_) | (d_1452_v76_)
            lhs207_ = globalState
            lhs208_ = d_1351_v2_
            lhs207_.f7 = rhs222_
            lhs208_.f25 = rhs223_
            d_1452_v76_ = rhs224_
            d_1453_v77_: _dafny.Array
            def lambda81_(d_1454_v2_):
                def lambda82_(d_1455_i4_):
                    return (D13_DC27(d_1454_v2_.f25) if True else D13_DC27(d_1454_v2_.f25))

                return lambda82_

            init45_ = lambda81_(d_1351_v2_)
            nw231_ = _dafny.Array(None, 3)
            for i0_45_ in range(nw231_.length(0)):
                nw231_[i0_45_] = init45_(i0_45_)
            d_1453_v77_ = nw231_
            d_1456_v78_: D13
            d_1456_v78_ = D13_DC27(d_1351_v2_.f25)
            index232_ = default__.safeIndex(749, (d_1453_v77_).length(0))
            (d_1453_v77_)[index232_] = d_1456_v78_
            index233_ = default__.safeIndex(749, (d_1453_v77_).length(0))
            (d_1453_v77_)[index233_] = d_1456_v78_
            d_1457_v79_: _dafny.Seq
            d_1457_v79_ = _dafny.SeqWithoutIsStrInference([d_1351_v2_.f25])
            d_1458_v80_: _dafny.Seq
            d_1458_v80_ = _dafny.SeqWithoutIsStrInference([len(d_1457_v79_), d_1349_v0_])
            d_1459_v81_: _dafny.MultiSet
            d_1459_v81_ = _dafny.MultiSet([d_1349_v0_])
            d_1460_v82_: _dafny.Map
            d_1460_v82_ = _dafny.Map({len(d_1458_v80_): d_1459_v81_})
            d_1461_v83_: _dafny.Map
            d_1461_v83_ = _dafny.Map({d_1349_v0_: ((d_1460_v82_)[d_1349_v0_] if (d_1349_v0_) in (d_1460_v82_) else d_1459_v81_)})
            d_1462_v84_: _dafny.Set
            d_1462_v84_ = _dafny.Set({d_1457_v79_})
            d_1463_v85_: _dafny.Array
            nw232_ = _dafny.Array(None, 15)
            nw232_[int(0)] = (d_1459_v81_).issubset(((d_1461_v83_)[d_1349_v0_] if (d_1349_v0_) in (d_1461_v83_) else d_1459_v81_))
            nw232_[int(1)] = (d_1457_v79_) <= (_dafny.SeqWithoutIsStrInference([(self).f24]))
            nw232_[int(2)] = (d_1351_v2_).f26
            nw232_[int(3)] = (d_1349_v0_) <= (d_1349_v0_)
            nw232_[int(4)] = (_dafny.SeqWithoutIsStrInference([d_1351_v2_.f25, d_1351_v2_.f25, d_1351_v2_.f25])) == (_dafny.SeqWithoutIsStrInference([d_1351_v2_.f25, (d_1351_v2_).f26, d_1351_v2_.f25]))
            nw232_[int(5)] = (not((d_1351_v2_).f26)) or (False)
            nw232_[int(6)] = (self).f24
            nw232_[int(7)] = (d_1351_v2_).f26
            nw232_[int(8)] = ((self).f24 if d_1351_v2_.f25 else (self).f24)
            nw232_[int(9)] = (d_1457_v79_)[default__.safeIndex(d_1349_v0_, len(d_1457_v79_))]
            nw232_[int(10)] = (d_1351_v2_).f26
            nw232_[int(11)] = (d_1351_v2_).f26
            nw232_[int(12)] = (d_1462_v84_) == (d_1462_v84_)
            nw232_[int(13)] = (d_1351_v2_).f26
            nw232_[int(14)] = (d_1351_v2_).f26
            d_1463_v85_ = nw232_
            index234_ = default__.safeIndex(568, (p0).length(0))
            (p0)[index234_] = default__.safeModuloInt((0) - (d_1349_v0_), len(_dafny.SeqWithoutIsStrInference([(d_1451_v75_)[default__.safeIndex(((d_1452_v76_)[d_1351_v2_.f25] if (d_1351_v2_.f25) in (d_1452_v76_) else d_1349_v0_), len(d_1451_v75_))] for d_1464_i5_ in range(default__.abs(-100))])))
            d_1465_v86_: _dafny.Seq
            d_1465_v86_ = _dafny.SeqWithoutIsStrInference([d_1451_v75_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_1466_i6_ in range(default__.abs(441))])])
            d_1467_v87_: _dafny.MultiSet
            d_1467_v87_ = _dafny.MultiSet([d_1463_v85_])
            index235_ = default__.safeIndex(568, (p0).length(0))
            rhs225_ = not(((d_1467_v87_) | ((D19_DC45(d_1467_v87_)).cf59)).issubset(d_1467_v87_))
            rhs226_ = d_1463_v85_
            rhs227_ = default__.safeModuloInt(len(d_1451_v75_), d_1349_v0_)
            rhs228_ = ((d_1465_v86_) + (d_1465_v86_)) + (_dafny.SeqWithoutIsStrInference([d_1451_v75_, d_1451_v75_, d_1451_v75_]))
            lhs209_ = globalState
            lhs210_ = p0
            lhs211_ = default__.safeIndex(568, (p0).length(0))
            lhs209_.f13 = rhs225_
            d_1463_v85_ = rhs226_
            lhs210_[lhs211_] = rhs227_
            d_1465_v86_ = rhs228_
            index236_ = default__.safeIndex(383, (d_1463_v85_).length(0))
            (d_1463_v85_)[index236_] = d_1351_v2_.f25
            index237_ = default__.safeIndex(383, (d_1463_v85_).length(0))
            (d_1463_v85_)[index237_] = (d_1349_v0_) < ((d_1349_v0_) * (d_1349_v0_))
            d_1468_v88_: _dafny.Map
            d_1468_v88_ = _dafny.Map({len(d_1451_v75_): _dafny.CodePoint('r')})
            d_1469_v89_: _dafny.Set
            d_1469_v89_ = _dafny.Set({d_1351_v2_.f25})
            if ((d_1349_v0_) * (((d_1459_v81_)[(self).fm8(True, d_1468_v88_, (d_1451_v75_).set(default__.safeIndex((p0)[default__.safeIndex(568, (p0).length(0))], len(d_1451_v75_)), default__.fm34(globalState)), globalState)] if ((self).fm8(True, d_1468_v88_, (d_1451_v75_).set(default__.safeIndex((p0)[default__.safeIndex(568, (p0).length(0))], len(d_1451_v75_)), default__.fm34(globalState)), globalState)) in (d_1459_v81_) else len(d_1469_v89_)))) == ((p0)[default__.safeIndex(568, (p0).length(0))]):
                d_1470_v90_: D4
                d_1470_v90_ = D4_DC10(d_1349_v0_, (0) - (d_1349_v0_), d_1351_v2_.f25, False, len(d_1452_v76_))
                (globalState).f2 = (d_1470_v90_).cf14
                d_1471_v91_: _dafny.Map
                d_1471_v91_ = _dafny.Map({(d_1351_v2_).f26: d_1451_v75_})
                d_1472_v92_: str
                d_1472_v92_ = _dafny.CodePoint('y')
                d_1473_v93_: _dafny.Array
                nw233_ = _dafny.Array(None, 29)
                nw233_[int(0)] = d_1451_v75_
                nw233_[int(1)] = ((d_1471_v91_)[(d_1351_v2_).f26] if ((d_1351_v2_).f26) in (d_1471_v91_) else d_1451_v75_)
                nw233_[int(2)] = default__.fm1((p0)[default__.safeIndex(568, (p0).length(0))], globalState)
                nw233_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hyp"))
                nw233_[int(4)] = d_1451_v75_
                nw233_[int(5)] = d_1451_v75_
                nw233_[int(6)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_1474_i7_ in range(default__.abs(-529))])
                nw233_[int(7)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdnqhsuh"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fuy")))
                nw233_[int(8)] = d_1451_v75_
                nw233_[int(9)] = _dafny.SeqWithoutIsStrInference([(d_1451_v75_)[default__.safeIndex((p0)[default__.safeIndex(568, (p0).length(0))], len(d_1451_v75_))] for d_1475_i8_ in range(default__.abs(-316))])
                nw233_[int(10)] = d_1451_v75_
                nw233_[int(11)] = d_1451_v75_
                nw233_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cepxkgleo"))
                nw233_[int(13)] = (d_1451_v75_).set(default__.safeIndex((p0)[default__.safeIndex(568, (p0).length(0))], len(d_1451_v75_)), _dafny.CodePoint('d'))
                nw233_[int(14)] = (d_1451_v75_) + ((d_1465_v86_)[default__.safeIndex(d_1349_v0_, len(d_1465_v86_))])
                nw233_[int(15)] = d_1451_v75_
                nw233_[int(16)] = d_1451_v75_
                nw233_[int(17)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_1476_i9_ in range(default__.abs(116))])
                nw233_[int(18)] = d_1451_v75_
                nw233_[int(19)] = (d_1451_v75_).set(default__.safeIndex(d_1349_v0_, len(d_1451_v75_)), d_1472_v92_)
                nw233_[int(20)] = (d_1451_v75_) + (_dafny.SeqWithoutIsStrInference([d_1472_v92_ for d_1477_i10_ in range(default__.abs(807))]))
                nw233_[int(21)] = d_1451_v75_
                nw233_[int(22)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_1478_i11_ in range(default__.abs(885))])
                nw233_[int(23)] = ((D13_DC26(d_1451_v75_, d_1351_v2_.f25, d_1351_v2_.f25)).cf33) + (d_1451_v75_)
                nw233_[int(24)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ctaxiqtj"))
                nw233_[int(25)] = d_1451_v75_
                nw233_[int(26)] = d_1451_v75_
                nw233_[int(27)] = d_1451_v75_
                nw233_[int(28)] = d_1451_v75_
                d_1473_v93_ = nw233_
                index238_ = default__.safeIndex(501, (d_1473_v93_).length(0))
                (d_1473_v93_)[index238_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rckcmrpbm"))) + (d_1451_v75_)
                d_1479_v94_: _dafny.Map
                d_1479_v94_ = _dafny.Map({d_1349_v0_: d_1451_v75_})
                index239_ = default__.safeIndex(501, (d_1473_v93_).length(0))
                (d_1473_v93_)[index239_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "srrgrldpd"))) + (((d_1479_v94_)[d_1349_v0_] if (d_1349_v0_) in (d_1479_v94_) else d_1451_v75_))
                (globalState).f16 = default__.safeModuloInt((p0)[default__.safeIndex(568, (p0).length(0))], (p0)[default__.safeIndex(568, (p0).length(0))])
                d_1480_v95_: _dafny.Array
                nw234_ = _dafny.Array(_dafny.Array(None, 0), 18)
                d_1480_v95_ = nw234_
                d_1481_v96_: _dafny.Map
                d_1481_v96_ = _dafny.Map({(self).f24: d_1480_v95_})
                d_1481_v96_ = (d_1481_v96_).set((d_1351_v2_).f26, d_1480_v95_)
                index240_ = default__.safeIndex(568, (p0).length(0))
                (p0)[index240_] = (default__.safeModuloInt(d_1349_v0_, (p0)[default__.safeIndex(568, (p0).length(0))])) + (d_1349_v0_)
            elif True:
                d_1482_v97_: D14
                d_1482_v97_ = D14_DC30(d_1458_v80_, d_1457_v79_)
                d_1483_v98_: _dafny.MultiSet
                d_1483_v98_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([D14_DC30(d_1458_v80_, d_1457_v79_), d_1482_v97_, d_1482_v97_])])
                d_1484_v99_: str
                d_1484_v99_ = _dafny.CodePoint('x')
                d_1483_v98_ = ((d_1483_v98_) - (d_1483_v98_) if False else ((d_1483_v98_).set(default__.fm46(d_1351_v2_.f25, d_1484_v99_, globalState), default__.abs(843))) - (d_1483_v98_))
                d_1452_v76_ = (d_1452_v76_).set((d_1457_v79_) < (d_1457_v79_), 931)
                d_1485_v100_: C2
                nw235_ = C2()
                nw235_.ctor__((614 if (d_1463_v85_)[default__.safeIndex(383, (d_1463_v85_).length(0))] else (p0)[default__.safeIndex(568, (p0).length(0))]))
                d_1485_v100_ = nw235_
                d_1485_v100_ = d_1485_v100_
                (globalState).f13 = ((d_1451_v75_).set(default__.safeIndex(len(d_1457_v79_), len(d_1451_v75_)), d_1484_v99_)) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vjgipyft")))
                (d_1351_v2_).f25 = (d_1350_v1_).issubset(d_1350_v1_)
        d_1486_v101_: _dafny.Seq
        d_1486_v101_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ckhhh"))
        d_1487_v102_: _dafny.Map
        d_1487_v102_ = _dafny.Map({d_1486_v101_: (self).f24})
        hi7_ = d_1349_v0_
        for d_1488_i12_ in range(len(d_1487_v102_), hi7_):
            d_1350_v1_ = (d_1350_v1_) - (_dafny.Set({(0) - (d_1349_v0_)}))
            (d_1351_v2_).f25 = (d_1351_v2_).f26
            if d_1351_v2_.f25:
                d_1489_v103_: _dafny.Map
                d_1489_v103_ = _dafny.Map({d_1349_v0_: (d_1351_v2_).f26})
                d_1490_v104_: D1
                d_1490_v104_ = D1_DC2(d_1489_v103_)
                d_1490_v104_ = d_1490_v104_
                d_1491_v105_: _dafny.Array
                nw236_ = _dafny.Array(_dafny.Array(None, 0), 24)
                d_1491_v105_ = nw236_
                d_1492_v106_: _dafny.Array
                nw237_ = _dafny.Array(None, 27)
                nw237_[int(0)] = d_1351_v2_.f25
                nw237_[int(1)] = (d_1351_v2_).f26
                nw237_[int(2)] = d_1351_v2_.f25
                nw237_[int(3)] = d_1351_v2_.f25
                nw237_[int(4)] = (d_1351_v2_).f26
                nw237_[int(5)] = (self).f24
                nw237_[int(6)] = (d_1351_v2_).f26
                nw237_[int(7)] = not(d_1351_v2_.f25)
                nw237_[int(8)] = not(d_1351_v2_.f25)
                nw237_[int(9)] = (d_1351_v2_).f26
                nw237_[int(10)] = (d_1351_v2_).f26
                nw237_[int(11)] = (self).fm7(d_1351_v2_.f25, d_1488_i12_, False, globalState)
                nw237_[int(12)] = d_1351_v2_.f25
                nw237_[int(13)] = (d_1351_v2_).f26
                nw237_[int(14)] = (d_1351_v2_).f26
                nw237_[int(15)] = (d_1351_v2_).f26
                nw237_[int(16)] = (self).f24
                nw237_[int(17)] = (self).f24
                nw237_[int(18)] = (d_1351_v2_).f26
                nw237_[int(19)] = (d_1351_v2_).f26
                nw237_[int(20)] = (d_1351_v2_).f26
                nw237_[int(21)] = not(d_1351_v2_.f25)
                nw237_[int(22)] = (self).fm6(len(d_1486_v101_), globalState)
                nw237_[int(23)] = d_1351_v2_.f25
                nw237_[int(24)] = not((d_1351_v2_).f26)
                nw237_[int(25)] = (d_1351_v2_).f26
                nw237_[int(26)] = (d_1351_v2_).f26
                d_1492_v106_ = nw237_
                index241_ = default__.safeIndex(805, (d_1491_v105_).length(0))
                (d_1491_v105_)[index241_] = d_1492_v106_
                d_1493_v107_: str
                d_1493_v107_ = _dafny.CodePoint('o')
                d_1494_v108_: D13
                d_1494_v108_ = D13_DC26(d_1486_v101_, (d_1351_v2_).fm9(d_1351_v2_.f25, globalState), d_1351_v2_.f25)
                d_1495_v109_: D19
                d_1495_v109_ = D19_DC46(d_1351_v2_.f25, d_1349_v0_)
                index242_ = default__.safeIndex(805, (d_1491_v105_).length(0))
                rhs229_ = (d_1493_v107_) not in (((d_1494_v108_).cf33) + ((d_1486_v101_).set(default__.safeIndex(d_1488_i12_, len(d_1486_v101_)), d_1493_v107_)))
                rhs230_ = (d_1351_v2_).f26
                rhs231_ = ((d_1488_i12_ if (d_1351_v2_).f26 else (d_1495_v109_).cf61) if d_1351_v2_.f25 else d_1349_v0_)
                rhs232_ = d_1492_v106_
                lhs212_ = globalState
                lhs213_ = globalState
                lhs214_ = globalState
                lhs215_ = d_1491_v105_
                lhs216_ = default__.safeIndex(805, (d_1491_v105_).length(0))
                lhs212_.f2 = rhs229_
                lhs213_.f0 = rhs230_
                lhs214_.f7 = rhs231_
                lhs215_[lhs216_] = rhs232_
                d_1496_v110_: _dafny.Map
                d_1496_v110_ = _dafny.Map({(d_1351_v2_).f26: d_1349_v0_})
                d_1497_v111_: _dafny.Seq
                d_1497_v111_ = _dafny.SeqWithoutIsStrInference([default__.fm1(d_1488_i12_, globalState), d_1486_v101_])
                d_1498_v112_: _dafny.Map
                d_1498_v112_ = _dafny.Map({(d_1349_v0_) * (len(d_1496_v110_)): (0) - ((0) - ((_dafny.MultiSet(((d_1497_v111_).set(default__.safeIndex(d_1349_v0_, len(d_1497_v111_)), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "je")))) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xydhcd"))])))).cardinality))})
                d_1499_v113_: D19
                d_1499_v113_ = D19_DC47(d_1490_v104_, -809, True)
                d_1500_v114_: _dafny.Map
                d_1500_v114_ = _dafny.Map({(d_1491_v105_)[default__.safeIndex(805, (d_1491_v105_).length(0))]: d_1351_v2_.f25})
                (globalState).f16 = ((d_1498_v112_)[(d_1499_v113_).cf63] if ((d_1499_v113_).cf63) in (d_1498_v112_) else (0) - (((0) - (len(d_1404_v41_)) if (d_1351_v2_).f26 else len(d_1500_v114_))))
                d_1501_v115_: _dafny.Array
                nw238_ = _dafny.Array(int(0), 27)
                d_1501_v115_ = nw238_
                def lambda83_(d_1502_i12_):
                    def lambda84_(d_1503_i13_):
                        return (d_1503_i13_) + (d_1502_i12_)

                    return lambda84_

                init46_ = lambda83_(d_1488_i12_)
                nw239_ = _dafny.Array(None, 27)
                for i0_46_ in range(nw239_.length(0)):
                    nw239_[i0_46_] = init46_(i0_46_)
                d_1501_v115_ = nw239_
                d_1504_v116_: C0
                nw240_ = C0()
                nw240_.ctor__(default__.fm2(d_1404_v41_, globalState))
                d_1504_v116_ = nw240_
                d_1505_v117_: _dafny.Map
                d_1505_v117_ = _dafny.Map({d_1351_v2_.f25: d_1504_v116_})
                d_1505_v117_ = (d_1505_v117_).set(d_1504_v116_.f27, d_1504_v116_)
            elif True:
                d_1506_v118_: _dafny.Array
                nw241_ = _dafny.Array(D6.default()(), 23)
                d_1506_v118_ = nw241_
                index243_ = default__.safeIndex(13, (d_1506_v118_).length(0))
                (d_1506_v118_)[index243_] = D6_DC12(_dafny.SeqWithoutIsStrInference([p0]))
                d_1507_v119_: str
                d_1507_v119_ = _dafny.CodePoint('w')
                d_1508_v120_: _dafny.Seq
                d_1508_v120_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                d_1509_v121_: D6
                d_1509_v121_ = D6_DC12(d_1508_v120_)
                d_1510_v122_: _dafny.Seq
                d_1510_v122_ = _dafny.SeqWithoutIsStrInference([d_1349_v0_])
                d_1511_v123_: _dafny.Map
                d_1511_v123_ = _dafny.Map({(d_1510_v122_)[default__.safeIndex(d_1488_i12_, len(d_1510_v122_))]: d_1351_v2_.f25})
                d_1512_v124_: _dafny.Set
                d_1512_v124_ = _dafny.Set({d_1507_v119_})
                d_1513_v125_: D4
                d_1513_v125_ = D4_DC10(d_1349_v0_, (_dafny.MultiSet([(0) - (len(d_1512_v124_)), d_1488_i12_])).cardinality, d_1351_v2_.f25, d_1351_v2_.f25, d_1488_i12_)
                d_1514_v126_: C2
                nw242_ = C2()
                nw242_.ctor__(len(d_1511_v123_))
                d_1514_v126_ = nw242_
                d_1515_v127_: _dafny.Seq
                d_1515_v127_ = _dafny.SeqWithoutIsStrInference([(d_1351_v2_).f26])
                d_1516_v128_: _dafny.Seq
                d_1516_v128_ = _dafny.SeqWithoutIsStrInference([d_1515_v127_])
                pat_let_tv31_ = d_1349_v0_
                index244_ = default__.safeIndex(13, (d_1506_v118_).length(0))
                def iife108_(_pat_let24_0):
                    def iife109_(d_1517_dt__update__tmp_h1_):
                        def iife110_(_pat_let25_0):
                            def iife111_(d_1518_dt__update_hcf13_h0_):
                                return D4_DC10((d_1517_dt__update__tmp_h1_).cf12, d_1518_dt__update_hcf13_h0_, (d_1517_dt__update__tmp_h1_).cf14, (d_1517_dt__update__tmp_h1_).cf15, (d_1517_dt__update__tmp_h1_).cf16)
                            return iife111_(_pat_let25_0)
                        return iife110_(pat_let_tv31_)
                    return iife109_(_pat_let24_0)
                rhs233_ = (d_1509_v121_ if (d_1507_v119_) not in (d_1486_v101_) else d_1509_v121_)
                rhs234_ = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ys")))) < ((0) - ((self).fm8(((d_1511_v123_)[d_1488_i12_] if (d_1488_i12_) in (d_1511_v123_) else (d_1351_v2_).f26), _dafny.Map({(iife108_(d_1513_v125_)).cf12: d_1507_v119_}), d_1486_v101_, globalState)))
                rhs235_ = d_1488_i12_
                rhs236_ = ((d_1404_v41_)[(_dafny.Set({d_1514_v126_, d_1514_v126_, d_1514_v126_})).ispropersubset(_dafny.Set({d_1514_v126_}))] if ((_dafny.Set({d_1514_v126_, d_1514_v126_, d_1514_v126_})).ispropersubset(_dafny.Set({d_1514_v126_}))) in (d_1404_v41_) else (_dafny.SeqWithoutIsStrInference([d_1515_v127_, d_1515_v127_])) <= (d_1516_v128_))
                rhs237_ = d_1351_v2_.f25
                lhs217_ = d_1506_v118_
                lhs218_ = default__.safeIndex(13, (d_1506_v118_).length(0))
                lhs219_ = globalState
                lhs220_ = globalState
                lhs221_ = globalState
                lhs222_ = globalState
                lhs217_[lhs218_] = rhs233_
                lhs219_.f2 = rhs234_
                lhs220_.f7 = rhs235_
                lhs221_.f2 = rhs236_
                lhs222_.f13 = rhs237_
                (globalState).f21 = d_1351_v2_.f25
                d_1349_v0_ = d_1488_i12_
                d_1519_v129_: _dafny.Set
                d_1519_v129_ = _dafny.Set({(d_1351_v2_).f26, (d_1351_v2_).f26})
                d_1520_v130_: _dafny.Seq
                d_1520_v130_ = _dafny.SeqWithoutIsStrInference([d_1519_v129_])
                d_1521_v131_: _dafny.MultiSet
                d_1521_v131_ = _dafny.MultiSet([((d_1520_v130_)[default__.safeIndex(d_1488_i12_, len(d_1520_v130_))]).intersection(d_1519_v129_), d_1519_v129_, _dafny.Set({not(default__.fm2(d_1404_v41_, globalState))})])
                d_1521_v131_ = _dafny.MultiSet([d_1519_v129_])
                d_1522_v132_: _dafny.Array
                nw243_ = _dafny.Array(False, 8)
                d_1522_v132_ = nw243_
                d_1523_v133_: _dafny.Map
                d_1523_v133_ = _dafny.Map({d_1488_i12_: _dafny.CodePoint('c')})
                d_1524_v134_: _dafny.Map
                d_1524_v134_ = _dafny.Map({d_1522_v132_: (self).fm6((self).fm8((d_1351_v2_).f26, d_1523_v133_, d_1486_v101_, globalState), globalState)})
                d_1524_v134_ = (d_1524_v134_).set(d_1522_v132_, not(d_1351_v2_.f25))
            (globalState).f2 = d_1351_v2_.f25
        d_1525_v135_: _dafny.Set
        d_1525_v135_ = _dafny.Set({(self).f24, d_1351_v2_.f25, (self).f24, True})
        d_1526_v136_: _dafny.Map
        d_1526_v136_ = _dafny.Map({(self).f24: len(d_1525_v135_)})
        hi8_ = len((d_1526_v136_) | (d_1526_v136_))
        for d_1527_i14_ in range((d_1349_v0_) + (d_1349_v0_), hi8_):
            d_1528_v137_: _dafny.Set
            d_1528_v137_ = (d_1525_v135_).intersection(default__.fm31(d_1349_v0_, (d_1351_v2_).f26, globalState))
            d_1528_v137_ = d_1525_v135_
            d_1529_v138_: D13
            d_1529_v138_ = D13_DC27(False)
            d_1529_v138_ = (d_1529_v138_ if (d_1351_v2_).f26 else d_1529_v138_)
            d_1530_v139_: T0
            nw244_ = C3()
            nw244_.ctor__()
            d_1530_v139_ = nw244_
            nw245_ = C3()
            nw245_.ctor__()
            d_1530_v139_ = nw245_
            d_1531_v140_: _dafny.Array
            nw246_ = _dafny.Array(False, 17)
            d_1531_v140_ = nw246_
            index245_ = default__.safeIndex(495, (d_1531_v140_).length(0))
            (d_1531_v140_)[index245_] = False
            index246_ = default__.safeIndex(495, (d_1531_v140_).length(0))
            rhs238_ = (d_1351_v2_).f26
            rhs239_ = d_1349_v0_
            rhs240_ = (0) - ((0) - (d_1527_i14_))
            rhs241_ = not(False)
            lhs223_ = globalState
            lhs224_ = globalState
            lhs225_ = globalState
            lhs226_ = d_1531_v140_
            lhs227_ = default__.safeIndex(495, (d_1531_v140_).length(0))
            lhs223_.f13 = rhs238_
            lhs224_.f7 = rhs239_
            lhs225_.f16 = rhs240_
            lhs226_[lhs227_] = rhs241_
        hi9_ = d_1349_v0_
        for d_1532_i15_ in range(169, hi9_):
            if (d_1351_v2_).f26:
                d_1533_v141_: _dafny.Set
                d_1533_v141_ = d_1525_v135_
                d_1534_v142_: _dafny.Array
                nw247_ = _dafny.Array(None, 2)
                nw247_[int(0)] = (_dafny.Set({(d_1351_v2_).f26, True})).ispropersubset(d_1525_v135_)
                nw247_[int(1)] = (d_1351_v2_).f26
                d_1534_v142_ = nw247_
                d_1535_v143_: _dafny.Seq
                d_1535_v143_ = _dafny.SeqWithoutIsStrInference([(d_1351_v2_).f26, d_1351_v2_.f25])
                index247_ = default__.safeIndex(416, (d_1534_v142_).length(0))
                (d_1534_v142_)[index247_] = (d_1535_v143_) < (d_1535_v143_)
                index248_ = default__.safeIndex(416, (d_1534_v142_).length(0))
                def iife112_():
                    coll60_ = _dafny.Set()
                    compr_60_: int
                    for compr_60_ in _dafny.IntegerRange(872, 204):
                        d_1536_v144_: int = compr_60_
                        if ((872) <= (d_1536_v144_)) and ((d_1536_v144_) < (204)):
                            coll60_ = coll60_.union(_dafny.Set([(d_1536_v144_) - (d_1532_i15_)]))
                    return _dafny.Set(coll60_)
                rhs242_ = (d_1533_v141_ if (d_1349_v0_) >= (d_1532_i15_) else d_1533_v141_)
                rhs243_ = (iife112_()
                ) | ((_dafny.Set({d_1532_i15_})) | (d_1350_v1_))
                rhs244_ = (self).f24
                rhs245_ = d_1351_v2_.f25
                lhs228_ = d_1351_v2_
                lhs229_ = d_1534_v142_
                lhs230_ = default__.safeIndex(416, (d_1534_v142_).length(0))
                d_1533_v141_ = rhs242_
                d_1350_v1_ = rhs243_
                lhs228_.f25 = rhs244_
                lhs229_[lhs230_] = rhs245_
                (globalState).f2 = ((self).f24 if (d_1351_v2_).f26 else (d_1534_v142_)[default__.safeIndex(416, (d_1534_v142_).length(0))])
                index249_ = default__.safeIndex(416, (d_1534_v142_).length(0))
                (d_1534_v142_)[index249_] = ((d_1349_v0_) != (len(d_1350_v1_)) if not(d_1351_v2_.f25) else (len(d_1525_v135_)) != (d_1532_i15_))
                d_1537_v145_: _dafny.Seq
                d_1537_v145_ = _dafny.SeqWithoutIsStrInference([len(d_1486_v101_), len(_dafny.Map({d_1351_v2_.f25: d_1349_v0_})), d_1532_i15_, d_1349_v0_, d_1532_i15_])
                d_1404_v41_ = (d_1404_v41_).set(d_1351_v2_.f25, ((0) - (d_1349_v0_)) in (d_1537_v145_))
                index250_ = default__.safeIndex(625, (p0).length(0))
                (p0)[index250_] = (0) - (d_1532_i15_)
                index251_ = default__.safeIndex(625, (p0).length(0))
                (p0)[index251_] = d_1349_v0_
            elif True:
                d_1538_v146_: _dafny.Seq
                d_1538_v146_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({867})])
                rhs246_ = ((d_1525_v135_) | (d_1525_v135_)).intersection(default__.fm14(d_1532_i15_, (d_1351_v2_).f26, globalState))
                rhs247_ = (d_1538_v146_)[default__.safeIndex(d_1349_v0_, len(d_1538_v146_))]
                d_1525_v135_ = rhs246_
                d_1350_v1_ = rhs247_
                (globalState).f7 = (((d_1526_v136_)[(d_1351_v2_).f26] if ((d_1351_v2_).f26) in (d_1526_v136_) else d_1532_i15_)) * ((d_1349_v0_ if d_1351_v2_.f25 else d_1532_i15_))
                (globalState).f2 = d_1351_v2_.f25
                index252_ = default__.safeIndex(498, (p0).length(0))
                (p0)[index252_] = d_1532_i15_
                index253_ = default__.safeIndex(498, (p0).length(0))
                (p0)[index253_] = default__.safeModuloInt(d_1349_v0_, d_1349_v0_)
                (d_1351_v2_).m5(default__.fm2(d_1404_v41_, globalState), globalState)
            (globalState).f7 = (d_1349_v0_) + (default__.safeDivisionInt(559, d_1349_v0_))
            d_1539_v147_: _dafny.Array
            nw248_ = _dafny.Array(None, 22)
            nw248_[int(0)] = d_1404_v41_
            nw248_[int(1)] = (_dafny.Map({d_1351_v2_.f25: True})) | (d_1404_v41_)
            nw248_[int(2)] = (default__.fm20(d_1349_v0_, (self).fm6(d_1532_i15_, globalState), (d_1351_v2_).f26, d_1351_v2_.f25, globalState)).set((d_1351_v2_).f26, (self).fm7(d_1351_v2_.f25, (0) - (d_1349_v0_), not(not((d_1351_v2_).f26)), globalState))
            nw248_[int(3)] = (d_1404_v41_).set(not((self).f24), (d_1351_v2_).f26)
            nw248_[int(4)] = d_1404_v41_
            nw248_[int(5)] = (d_1404_v41_) | (d_1404_v41_)
            nw248_[int(6)] = d_1404_v41_
            nw248_[int(7)] = d_1404_v41_
            nw248_[int(8)] = d_1404_v41_
            nw248_[int(9)] = d_1404_v41_
            nw248_[int(10)] = (d_1404_v41_) | (d_1404_v41_)
            nw248_[int(11)] = d_1404_v41_
            nw248_[int(12)] = d_1404_v41_
            nw248_[int(13)] = d_1404_v41_
            nw248_[int(14)] = _dafny.Map({(d_1351_v2_).f26: False})
            nw248_[int(15)] = (d_1404_v41_) | (d_1404_v41_)
            nw248_[int(16)] = (d_1404_v41_) | (d_1404_v41_)
            nw248_[int(17)] = (d_1404_v41_) | (d_1404_v41_)
            nw248_[int(18)] = (d_1404_v41_) | (_dafny.Map({not((d_1351_v2_).f26): d_1351_v2_.f25}))
            nw248_[int(19)] = d_1404_v41_
            nw248_[int(20)] = d_1404_v41_
            nw248_[int(21)] = (d_1404_v41_) | (d_1404_v41_)
            d_1539_v147_ = nw248_
            d_1540_v148_: _dafny.Array
            nw249_ = _dafny.Array(_dafny.Set({}), 13)
            d_1540_v148_ = nw249_
            index254_ = default__.safeIndex(176, (d_1540_v148_).length(0))
            (d_1540_v148_)[index254_] = (d_1525_v135_) | (d_1525_v135_)
            d_1541_v149_: _dafny.Seq
            d_1541_v149_ = _dafny.SeqWithoutIsStrInference([d_1532_i15_])
            d_1542_v150_: _dafny.Seq
            d_1542_v150_ = _dafny.SeqWithoutIsStrInference([(self).f24])
            index255_ = default__.safeIndex(176, (d_1540_v148_).length(0))
            rhs248_ = d_1539_v147_
            rhs249_ = d_1351_v2_.f25
            rhs250_ = (d_1542_v150_)[default__.safeIndex(d_1532_i15_, len(d_1542_v150_))]
            rhs251_ = d_1525_v135_
            rhs252_ = d_1541_v149_
            lhs231_ = globalState
            lhs232_ = globalState
            lhs233_ = d_1540_v148_
            lhs234_ = default__.safeIndex(176, (d_1540_v148_).length(0))
            d_1539_v147_ = rhs248_
            lhs231_.f0 = rhs249_
            lhs232_.f2 = rhs250_
            lhs233_[lhs234_] = rhs251_
            d_1541_v149_ = rhs252_
            index256_ = default__.safeIndex(57, (p0).length(0))
            (p0)[index256_] = d_1349_v0_
            index257_ = default__.safeIndex(57, (p0).length(0))
            rhs253_ = (0) - (default__.safeDivisionInt(d_1349_v0_, 471))
            rhs254_ = default__.fm2(d_1404_v41_, globalState)
            lhs235_ = p0
            lhs236_ = default__.safeIndex(57, (p0).length(0))
            lhs237_ = globalState
            lhs235_[lhs236_] = rhs253_
            lhs237_.f21 = rhs254_

    def m1(self, p0, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: bool = False
        d_1543_v0_: _dafny.Seq
        d_1543_v0_ = _dafny.SeqWithoutIsStrInference([p0])
        d_1544_v1_: _dafny.Seq
        d_1544_v1_ = _dafny.SeqWithoutIsStrInference([(d_1543_v0_)[default__.safeIndex(p0, len(d_1543_v0_))]])
        hi10_ = p0
        for d_1545_i0_ in range((len(d_1544_v1_)) - (p0), hi10_):
            (globalState).f21 = (self).f24
            d_1546_v2_: C3
            nw250_ = C3()
            nw250_.ctor__()
            d_1546_v2_ = nw250_
            d_1547_v3_: C2
            nw251_ = C2()
            nw251_.ctor__(p0)
            d_1547_v3_ = nw251_
            if True:
                (globalState).f0 = True
                d_1548_v4_: str
                d_1548_v4_ = _dafny.CodePoint('k')
                d_1549_v5_: _dafny.MultiSet
                d_1549_v5_ = _dafny.MultiSet([len(_dafny.Map({(self).f24: d_1548_v4_}))])
                d_1550_v6_: _dafny.Set
                d_1550_v6_ = _dafny.Set({default__.safeDivisionInt((d_1547_v3_).f28, default__.fm0((d_1547_v3_).f28, (d_1549_v5_).cardinality, len(d_1544_v1_), (self).f24, globalState)), len((_dafny.SeqWithoutIsStrInference([len(default__.fm21(globalState)), (d_1547_v3_).f28])) + (d_1543_v0_)), len((d_1543_v0_) + (d_1544_v1_))})
                d_1551_v7_: _dafny.Map
                d_1551_v7_ = _dafny.Map({(d_1547_v3_).f28: d_1548_v4_})
                d_1552_v8_: _dafny.Seq
                d_1552_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gffxeber"))
                d_1553_v9_: _dafny.Map
                d_1553_v9_ = _dafny.Map({(d_1547_v3_).f28: (self).f24})
                d_1554_v11_: _dafny.Seq
                d_1554_v11_ = _dafny.SeqWithoutIsStrInference([(self).f24, (self).f24])
                d_1555_v12_: _dafny.Seq
                d_1555_v12_ = d_1554_v11_
                def iife113_():
                    coll61_ = _dafny.Set()
                    compr_61_: int
                    for compr_61_ in (d_1553_v9_).keys.Elements:
                        d_1556_v10_: int = compr_61_
                        if (d_1556_v10_) in (d_1553_v9_):
                            coll61_ = coll61_.union(_dafny.Set([default__.safeModuloInt(d_1556_v10_, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cftvrnjd"))): len(_dafny.Map({96: len(_dafny.Map({_dafny.CodePoint('i'): 379}))}))})) for d_1557_i1_ in range(default__.abs(879))]))).cardinality)]))
                    return _dafny.Set(coll61_)
                rhs255_ = (self).f24
                rhs256_ = (p0) <= ((p0) - ((self).fm8(False, d_1551_v7_, d_1552_v8_, globalState)))
                rhs257_ = ((d_1550_v6_) | (d_1550_v6_)) | (iife113_()
                )
                rhs258_ = ((d_1547_v3_).fm17((self).f24, _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f24])) for d_1558_i2_ in range(default__.abs(187))]), d_1545_i0_, d_1555_v12_, globalState)) + (default__.safeModuloInt((d_1547_v3_).f28, len(_dafny.Map({p0: False}))))
                lhs238_ = globalState
                lhs239_ = globalState
                lhs238_.f21 = rhs255_
                r1 = rhs256_
                d_1550_v6_ = rhs257_
                lhs239_.f7 = rhs258_
                d_1559_v13_: _dafny.Array
                def lambda85_(d_1560_v8_):
                    def lambda86_(d_1561_i3_):
                        return d_1560_v8_

                    return lambda86_

                init47_ = lambda85_(d_1552_v8_)
                nw252_ = _dafny.Array(None, 29)
                for i0_47_ in range(nw252_.length(0)):
                    nw252_[i0_47_] = init47_(i0_47_)
                d_1559_v13_ = nw252_
                index258_ = default__.safeIndex(731, (d_1559_v13_).length(0))
                (d_1559_v13_)[index258_] = d_1552_v8_
                index259_ = default__.safeIndex(731, (d_1559_v13_).length(0))
                (d_1559_v13_)[index259_] = _dafny.SeqWithoutIsStrInference([d_1548_v4_ for d_1562_i4_ in range(default__.abs(777))])
                d_1563_v14_: _dafny.Array
                nw253_ = _dafny.Array(False, 2)
                d_1563_v14_ = nw253_
                index260_ = default__.safeIndex(762, (d_1563_v14_).length(0))
                (d_1563_v14_)[index260_] = (self).f24
                index261_ = default__.safeIndex(762, (d_1563_v14_).length(0))
                (d_1563_v14_)[index261_] = (len(d_1543_v0_)) <= (-113)
                d_1564_v16_: _dafny.MultiSet
                d_1564_v16_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([p0])])
                d_1565_v17_: D16
                def iife114_():
                    coll62_ = _dafny.Map()
                    compr_62_: _dafny.Seq
                    for compr_62_ in (d_1564_v16_).Elements:
                        d_1566_v15_: _dafny.Seq = compr_62_
                        if (d_1566_v15_) in (d_1564_v16_):
                            coll62_[d_1566_v15_] = d_1554_v11_
                    return _dafny.Map(coll62_)
                d_1565_v17_ = D16_DC35(iife114_()
)
                d_1567_v18_: D16
                d_1567_v18_ = D16_DC37(d_1565_v17_)
                index262_ = default__.safeIndex(731, (d_1559_v13_).length(0))
                rhs259_ = (d_1552_v8_) + ((default__.fm1(d_1545_i0_, globalState)) + (d_1552_v8_))
                rhs260_ = (((d_1559_v13_)[default__.safeIndex(731, (d_1559_v13_).length(0))]) + (d_1552_v8_)) + (d_1552_v8_)
                rhs261_ = d_1548_v4_
                rhs262_ = d_1567_v18_
                lhs240_ = globalState
                lhs241_ = d_1559_v13_
                lhs242_ = default__.safeIndex(731, (d_1559_v13_).length(0))
                lhs240_.f8 = rhs259_
                lhs241_[lhs242_] = rhs260_
                d_1548_v4_ = rhs261_
                d_1567_v18_ = rhs262_
            elif True:
                d_1568_v19_: _dafny.Array
                nw254_ = _dafny.Array(_dafny.Map({}), 9)
                d_1568_v19_ = nw254_
                d_1569_v20_: _dafny.Map
                d_1569_v20_ = _dafny.Map({p0: d_1568_v19_})
                (globalState).f22 = ((d_1569_v20_)[default__.fm0((d_1547_v3_).f28, d_1545_i0_, p0, (self).f24, globalState)] if (default__.fm0((d_1547_v3_).f28, d_1545_i0_, p0, (self).f24, globalState)) in (d_1569_v20_) else d_1568_v19_)
                d_1570_v21_: _dafny.Array
                nw255_ = _dafny.Array(_dafny.Seq({}), 16)
                d_1570_v21_ = nw255_
                d_1571_v22_: _dafny.Seq
                d_1571_v22_ = _dafny.SeqWithoutIsStrInference([(self).f24])
                index263_ = default__.safeIndex(133, (d_1570_v21_).length(0))
                (d_1570_v21_)[index263_] = d_1571_v22_
                d_1572_v23_: _dafny.Array
                nw256_ = _dafny.Array(int(0), 27)
                d_1572_v23_ = nw256_
                d_1573_v24_: _dafny.Seq
                d_1573_v24_ = _dafny.SeqWithoutIsStrInference([d_1572_v23_, d_1572_v23_, d_1572_v23_, d_1572_v23_, d_1572_v23_])
                d_1574_v25_: _dafny.Set
                d_1574_v25_ = _dafny.Set({d_1545_i0_})
                index264_ = default__.safeIndex(133, (d_1570_v21_).length(0))
                rhs263_ = (d_1573_v24_)[default__.safeIndex(d_1545_i0_, len(d_1573_v24_))]
                rhs264_ = (d_1574_v25_).isdisjoint((D14_DC28(d_1574_v25_)).cf37)
                rhs265_ = p0
                rhs266_ = d_1571_v22_
                lhs243_ = globalState
                lhs244_ = globalState
                lhs245_ = d_1570_v21_
                lhs246_ = default__.safeIndex(133, (d_1570_v21_).length(0))
                r0 = rhs263_
                lhs243_.f13 = rhs264_
                lhs244_.f7 = rhs265_
                lhs245_[lhs246_] = rhs266_
                d_1575_v26_: _dafny.Map
                d_1575_v26_ = _dafny.Map({((d_1547_v3_).f28) * (899): not(((self).f24) or ((self).f24))})
                def iife115_():
                    def iife117_():
                        coll65_ = _dafny.Set()
                        compr_65_: int
                        for compr_65_ in (_dafny.MultiSet([(d_1547_v3_).f28, d_1545_i0_])).Elements:
                            d_1578_v28_: int = compr_65_
                            if (d_1578_v28_) in (_dafny.MultiSet([(d_1547_v3_).f28, d_1545_i0_])):
                                coll65_ = coll65_.union(_dafny.Set([(d_1578_v28_) - (227)]))
                        return _dafny.Set(coll65_)
                    coll63_ = _dafny.Map()
                    def iife116_():
                        coll64_ = _dafny.Set()
                        compr_64_: int
                        for compr_64_ in (_dafny.MultiSet([(d_1547_v3_).f28, d_1545_i0_])).Elements:
                            d_1576_v28_: int = compr_64_
                            if (d_1576_v28_) in (_dafny.MultiSet([(d_1547_v3_).f28, d_1545_i0_])):
                                coll64_ = coll64_.union(_dafny.Set([(d_1576_v28_) - (227)]))
                        return _dafny.Set(coll64_)
                    compr_63_: int
                    for compr_63_ in (iife116_()
                    ).Elements:
                        d_1577_v27_: int = compr_63_
                        if (d_1577_v27_) in (iife117_()
                        ):
                            coll63_[(d_1577_v27_) * (len(_dafny.Map({(d_1547_v3_).f28: d_1545_i0_})))] = (self).f24
                    return _dafny.Map(coll63_)
                d_1575_v26_ = (iife115_()
                 if not ((self).f24) or ((self).f24) else (d_1575_v26_) | (_dafny.Map({(d_1547_v3_).f28: (self).f24})))
                d_1579_v29_: _dafny.Seq
                d_1579_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ccmntf"))
                d_1580_v30_: _dafny.Map
                d_1580_v30_ = _dafny.Map({(d_1570_v21_)[default__.safeIndex(133, (d_1570_v21_).length(0))]: d_1579_v29_})
                d_1581_v31_: _dafny.MultiSet
                d_1581_v31_ = _dafny.MultiSet([p0, (d_1547_v3_).f28])
                default__.m0(((d_1570_v21_)[default__.safeIndex(133, (d_1570_v21_).length(0))])[default__.safeIndex((d_1547_v3_).f28, len((d_1570_v21_)[default__.safeIndex(133, (d_1570_v21_).length(0))]))], d_1545_i0_, d_1580_v30_, ((d_1581_v31_) - (d_1581_v31_)).cardinality, globalState)
                (globalState).f16 = d_1545_i0_
        d_1582_v32_: str
        d_1582_v32_ = _dafny.CodePoint('c')
        d_1582_v32_ = d_1582_v32_
        (globalState).f0 = not(not ((self).f24) or ((self).f24))
        d_1583_v33_: _dafny.Map
        d_1583_v33_ = _dafny.Map({p0: (self).f24})
        rhs267_ = p0
        rhs268_ = p0
        rhs269_ = default__.safeDivisionInt(p0, len((d_1583_v33_) | (d_1583_v33_)))
        lhs247_ = globalState
        lhs248_ = globalState
        lhs249_ = globalState
        lhs247_.f16 = rhs267_
        lhs248_.f16 = rhs268_
        lhs249_.f16 = rhs269_
        d_1584_v34_: _dafny.Array
        def lambda87_(d_1585_p0_):
            def lambda88_(d_1586_i5_):
                return (d_1586_i5_) - (d_1585_p0_)

            return lambda88_

        init48_ = lambda87_(p0)
        nw257_ = _dafny.Array(None, 19)
        for i0_48_ in range(nw257_.length(0)):
            nw257_[i0_48_] = init48_(i0_48_)
        d_1584_v34_ = nw257_
        r0 = d_1584_v34_
        if (self).f24:
            index265_ = default__.safeIndex(748, (d_1584_v34_).length(0))
            (d_1584_v34_)[index265_] = p0
            d_1587_v35_: _dafny.Array
            nw258_ = _dafny.Array(False, 13)
            d_1587_v35_ = nw258_
            index266_ = default__.safeIndex(633, (d_1587_v35_).length(0))
            (d_1587_v35_)[index266_] = (self).f24
            index267_ = default__.safeIndex(748, (d_1584_v34_).length(0))
            index268_ = default__.safeIndex(633, (d_1587_v35_).length(0))
            rhs270_ = (0) - (-510)
            rhs271_ = not((self).f24)
            rhs272_ = (p0) - (p0)
            rhs273_ = (p0) * ((p0) - (p0))
            lhs250_ = d_1584_v34_
            lhs251_ = default__.safeIndex(748, (d_1584_v34_).length(0))
            lhs252_ = d_1587_v35_
            lhs253_ = default__.safeIndex(633, (d_1587_v35_).length(0))
            lhs254_ = globalState
            lhs255_ = globalState
            lhs250_[lhs251_] = rhs270_
            lhs252_[lhs253_] = rhs271_
            lhs254_.f16 = rhs272_
            lhs255_.f16 = rhs273_
            d_1588_v36_: _dafny.Map
            d_1588_v36_ = _dafny.Map({not((self).f24): -656})
            (globalState).f7 = ((d_1588_v36_)[((self).f24) == ((d_1587_v35_)[default__.safeIndex(633, (d_1587_v35_).length(0))])] if (((self).f24) == ((d_1587_v35_)[default__.safeIndex(633, (d_1587_v35_).length(0))])) in (d_1588_v36_) else p0)
            d_1589_v37_: _dafny.Set
            d_1589_v37_ = _dafny.Set({533})
            d_1590_v38_: _dafny.Seq
            d_1590_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ospctjgb"))
            d_1591_v39_: _dafny.Array
            nw259_ = _dafny.Array(None, 19)
            nw259_[int(0)] = (0) - (len(d_1589_v37_))
            nw259_[int(1)] = len(d_1590_v38_)
            nw259_[int(2)] = (d_1584_v34_)[default__.safeIndex(748, (d_1584_v34_).length(0))]
            nw259_[int(3)] = (d_1584_v34_)[default__.safeIndex(748, (d_1584_v34_).length(0))]
            nw259_[int(4)] = (0) - (p0)
            nw259_[int(5)] = (d_1584_v34_)[default__.safeIndex(748, (d_1584_v34_).length(0))]
            nw259_[int(6)] = (d_1584_v34_)[default__.safeIndex(748, (d_1584_v34_).length(0))]
            nw259_[int(7)] = (d_1584_v34_)[default__.safeIndex(748, (d_1584_v34_).length(0))]
            nw259_[int(8)] = (0) - ((d_1584_v34_)[default__.safeIndex(748, (d_1584_v34_).length(0))])
            nw259_[int(9)] = len(d_1590_v38_)
            nw259_[int(10)] = (d_1584_v34_)[default__.safeIndex(748, (d_1584_v34_).length(0))]
            nw259_[int(11)] = (self).fm8((self).f24, (_dafny.Map({692: d_1582_v32_})).set(-65, d_1582_v32_), d_1590_v38_, globalState)
            nw259_[int(12)] = p0
            nw259_[int(13)] = -377
            nw259_[int(14)] = 143
            nw259_[int(15)] = len(_dafny.Map({not((d_1587_v35_)[default__.safeIndex(633, (d_1587_v35_).length(0))]): (self).f24}))
            nw259_[int(16)] = (d_1584_v34_)[default__.safeIndex(748, (d_1584_v34_).length(0))]
            nw259_[int(17)] = (d_1584_v34_)[default__.safeIndex(748, (d_1584_v34_).length(0))]
            nw259_[int(18)] = p0
            d_1591_v39_ = nw259_
            d_1592_v40_: D6
            d_1592_v40_ = D6_DC12(_dafny.SeqWithoutIsStrInference([d_1591_v39_]))
            d_1593_v42_: _dafny.Map
            d_1593_v42_ = _dafny.Map({d_1582_v32_: (d_1587_v35_)[default__.safeIndex(633, (d_1587_v35_).length(0))]})
            d_1594_v43_: _dafny.Map
            d_1594_v43_ = _dafny.Map({(d_1584_v34_)[default__.safeIndex(748, (d_1584_v34_).length(0))]: d_1593_v42_})
            index269_ = default__.safeIndex(633, (d_1587_v35_).length(0))
            def iife118_():
                coll66_ = _dafny.Map()
                compr_66_: int
                for compr_66_ in _dafny.IntegerRange(-222, 403):
                    d_1596_v41_: int = compr_66_
                    if ((-222) <= (d_1596_v41_)) and ((d_1596_v41_) < (403)):
                        coll66_[default__.safeDivisionInt(d_1596_v41_, (d_1584_v34_)[default__.safeIndex(748, (d_1584_v34_).length(0))])] = _dafny.Map({d_1582_v32_: (d_1587_v35_)[default__.safeIndex(633, (d_1587_v35_).length(0))]})
                return _dafny.Map(coll66_)
            rhs274_ = (d_1587_v35_)[default__.safeIndex(633, (d_1587_v35_).length(0))]
            rhs275_ = _dafny.SeqWithoutIsStrInference([d_1582_v32_ for d_1595_i6_ in range(default__.abs(231))])
            rhs276_ = not(((iife118_()
            ) | (d_1594_v43_)) != (default__.fm48(d_1582_v32_, (d_1584_v34_)[default__.safeIndex(748, (d_1584_v34_).length(0))], globalState)))
            rhs277_ = d_1592_v40_
            rhs278_ = not ((self).f24) or ((d_1587_v35_)[default__.safeIndex(633, (d_1587_v35_).length(0))])
            lhs256_ = d_1587_v35_
            lhs257_ = default__.safeIndex(633, (d_1587_v35_).length(0))
            lhs258_ = globalState
            lhs259_ = globalState
            lhs260_ = globalState
            lhs256_[lhs257_] = rhs274_
            lhs258_.f20 = rhs275_
            lhs259_.f0 = rhs276_
            d_1592_v40_ = rhs277_
            lhs260_.f13 = rhs278_
            (globalState).f13 = not (True) or ((p0) <= ((d_1584_v34_)[default__.safeIndex(748, (d_1584_v34_).length(0))]))
            d_1597_v44_: C2
            nw260_ = C2()
            nw260_.ctor__((d_1584_v34_)[default__.safeIndex(748, (d_1584_v34_).length(0))])
            d_1597_v44_ = nw260_
            d_1597_v44_ = d_1597_v44_
        elif True:
            d_1598_v45_: _dafny.Map
            d_1598_v45_ = _dafny.Map({(self).fm7((self).f24, len(_dafny.Map({_dafny.CodePoint('o'): (self).f24})), (self).f24, globalState): (self).f24})
            d_1599_v46_: _dafny.Set
            d_1599_v46_ = _dafny.Set({default__.fm2(d_1598_v45_, globalState)})
            d_1600_v47_: _dafny.Map
            d_1600_v47_ = _dafny.Map({d_1599_v46_: default__.fm18((self).f24, d_1582_v32_, (self).fm6(p0, globalState), globalState)})
            d_1601_v48_: _dafny.Seq
            d_1601_v48_ = _dafny.SeqWithoutIsStrInference([True])
            d_1600_v47_ = (d_1600_v47_).set(default__.fm14(p0, (self).f24, globalState), d_1601_v48_)
            d_1602_v49_: _dafny.Array
            nw261_ = _dafny.Array(_dafny.Array(None, 0), 7)
            d_1602_v49_ = nw261_
            d_1602_v49_ = d_1602_v49_
            index270_ = default__.safeIndex(70, (d_1584_v34_).length(0))
            (d_1584_v34_)[index270_] = p0
            index271_ = default__.safeIndex(70, (d_1584_v34_).length(0))
            (d_1584_v34_)[index271_] = len(((_dafny.Set({(self).f24, True})) - (d_1599_v46_)) - (d_1599_v46_))
            d_1601_v48_ = d_1601_v48_
            nw262_ = _dafny.Array(int(0), 15)
            r0 = nw262_
        r0 = d_1584_v34_
        r1 = (self).f24
        return r0, r1

    @property
    def f24(self):
        return self._f24

class C6(T1, T0):
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    def ctor__(self):
        pass
        pass

    def fm7(self, p0, p1, p2, globalState):
        return False

    def fm5(self, p0, p1, p2, p3, globalState):
        return ((_dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xdbnkvr"))}) if False else _dafny.Map({False: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_1603_i0_ in range(default__.abs(585))])}))) | (_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bgtt"))}))

    def fm6(self, p0, globalState):
        return True

    def m2(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: _dafny.Array = _dafny.Array(None, 0)
        (globalState).f13 = p1
        (globalState).f16 = -429
        hi11_ = p3
        for d_1604_i0_ in range(p3, hi11_):
            d_1605_v0_: _dafny.MultiSet
            d_1605_v0_ = _dafny.MultiSet([307])
            d_1606_v1_: _dafny.MultiSet
            d_1606_v1_ = _dafny.MultiSet([d_1605_v0_])
            (globalState).f0 = ((_dafny.MultiSet([d_1605_v0_, d_1605_v0_, d_1605_v0_, (d_1605_v0_).set(d_1604_i0_, default__.abs(-346))])).ispropersubset((d_1606_v1_).set(d_1605_v0_, default__.abs(p0)))) == (p1)
            d_1607_v2_: _dafny.Seq
            d_1607_v2_ = _dafny.SeqWithoutIsStrInference([not(p1), p1])
            d_1608_v3_: C0
            nw263_ = C0()
            nw263_.ctor__((d_1607_v2_)[default__.safeIndex(p0, len(d_1607_v2_))])
            d_1608_v3_ = nw263_
            d_1609_v4_: _dafny.Seq
            d_1609_v4_ = _dafny.SeqWithoutIsStrInference([p3])
            if (len(d_1609_v4_)) >= ((d_1604_i0_ if p1 else d_1604_i0_)):
                (globalState).f7 = p0
                d_1610_v5_: _dafny.Array
                nw264_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 27)
                d_1610_v5_ = nw264_
                d_1611_v6_: _dafny.Seq
                d_1611_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "odwrbay"))
                index272_ = default__.safeIndex(423, (d_1610_v5_).length(0))
                (d_1610_v5_)[index272_] = d_1611_v6_
                index273_ = default__.safeIndex(423, (d_1610_v5_).length(0))
                (d_1610_v5_)[index273_] = default__.fm1(151, globalState)
                (globalState).f16 = d_1604_i0_
                d_1612_v7_: _dafny.Array
                nw265_ = _dafny.Array(int(0), 20)
                d_1612_v7_ = nw265_
                d_1612_v7_ = d_1612_v7_
                d_1613_v8_: _dafny.Seq
                d_1613_v8_ = _dafny.SeqWithoutIsStrInference([(d_1610_v5_)[default__.safeIndex(423, (d_1610_v5_).length(0))], (d_1610_v5_)[default__.safeIndex(423, (d_1610_v5_).length(0))], (d_1610_v5_)[default__.safeIndex(423, (d_1610_v5_).length(0))], (d_1610_v5_)[default__.safeIndex(423, (d_1610_v5_).length(0))], default__.fm1(d_1604_i0_, globalState)])
                d_1614_v9_: _dafny.MultiSet
                d_1614_v9_ = _dafny.MultiSet([d_1608_v3_.f27])
                index274_ = default__.safeIndex(423, (d_1610_v5_).length(0))
                (d_1610_v5_)[index274_] = ((d_1613_v8_)[default__.safeIndex((0) - ((d_1609_v4_)[default__.safeIndex((d_1614_v9_).cardinality, len(d_1609_v4_))]), len(d_1613_v8_))]) + (d_1611_v6_)
            elif True:
                index275_ = default__.safeIndex(398, (p2).length(0))
                (p2)[index275_] = (d_1609_v4_)[default__.safeIndex(len(d_1607_v2_), len(d_1609_v4_))]
                index276_ = default__.safeIndex(398, (p2).length(0))
                (p2)[index276_] = p3
                (d_1608_v3_).f27 = d_1608_v3_.f27
                d_1615_v10_: _dafny.Set
                d_1615_v10_ = _dafny.Set({d_1608_v3_.f27, (self).fm6(p3, globalState)})
                d_1616_v11_: _dafny.Array
                nw266_ = _dafny.Array(None, 15)
                nw266_[int(0)] = d_1615_v10_
                nw266_[int(1)] = d_1615_v10_
                nw266_[int(2)] = d_1615_v10_
                nw266_[int(3)] = d_1615_v10_
                nw266_[int(4)] = d_1615_v10_
                nw266_[int(5)] = d_1615_v10_
                nw266_[int(6)] = d_1615_v10_
                nw266_[int(7)] = _dafny.Set({d_1608_v3_.f27, d_1608_v3_.f27, d_1608_v3_.f27})
                nw266_[int(8)] = d_1615_v10_
                nw266_[int(9)] = d_1615_v10_
                nw266_[int(10)] = d_1615_v10_
                nw266_[int(11)] = _dafny.Set({d_1608_v3_.f27})
                nw266_[int(12)] = d_1615_v10_
                nw266_[int(13)] = d_1615_v10_
                nw266_[int(14)] = d_1615_v10_
                d_1616_v11_ = nw266_
                d_1617_v12_: D7
                d_1617_v12_ = D7_DC16((d_1616_v11_ if d_1608_v3_.f27 else d_1616_v11_))
                d_1617_v12_ = d_1617_v12_
                d_1618_v13_: C4
                nw267_ = C4()
                nw267_.ctor__(p1, d_1608_v3_.f27)
                d_1618_v13_ = nw267_
                d_1619_v14_: _dafny.Map
                d_1619_v14_ = _dafny.Map({d_1608_v3_.f27: d_1618_v13_})
                d_1620_v15_: _dafny.Map
                d_1620_v15_ = _dafny.Map({((d_1619_v14_)[d_1618_v13_.f25] if (d_1618_v13_.f25) in (d_1619_v14_) else d_1618_v13_): p0})
                d_1620_v15_ = (d_1620_v15_).set(d_1618_v13_, (0) - ((0) - (default__.fm0(511, (p2)[default__.safeIndex(398, (p2).length(0))], (p2)[default__.safeIndex(398, (p2).length(0))], p1, globalState))))
                d_1621_v16_: _dafny.Map
                d_1621_v16_ = _dafny.Map({(p2)[default__.safeIndex(398, (p2).length(0))]: not(d_1608_v3_.f27)})
                d_1622_v17_: _dafny.Array
                nw268_ = _dafny.Array(_dafny.Seq({}), 25)
                d_1622_v17_ = nw268_
                d_1623_v18_: _dafny.Map
                d_1623_v18_ = _dafny.Map({d_1621_v16_: d_1622_v17_})
                d_1623_v18_ = (d_1623_v18_).set(d_1621_v16_, d_1622_v17_)
            d_1624_v19_: _dafny.Map
            d_1624_v19_ = _dafny.Map({d_1604_i0_: d_1608_v3_.f27})
            d_1625_v20_: _dafny.Map
            d_1625_v20_ = _dafny.Map({d_1624_v19_: p2})
            d_1625_v20_ = ((d_1625_v20_) | ((d_1625_v20_).set(_dafny.Map({d_1604_i0_: p1}), p2))) | (d_1625_v20_)
        d_1626_v21_: _dafny.Map
        d_1626_v21_ = _dafny.Map({p3: p1})
        (globalState).f2 = (D19_DC47(D1_DC2(d_1626_v21_), p3, p1)).cf64
        (globalState).f2 = (p1) == (p1)
        d_1627_v22_: _dafny.Array
        nw269_ = _dafny.Array(None, 3)
        nw269_[int(0)] = (False if p1 else not(p1))
        nw269_[int(1)] = p1
        nw269_[int(2)] = p1
        d_1627_v22_ = nw269_
        index277_ = default__.safeIndex(13, (d_1627_v22_).length(0))
        (d_1627_v22_)[index277_] = not (not(p1)) or (True)
        index278_ = default__.safeIndex(860, (d_1627_v22_).length(0))
        (d_1627_v22_)[index278_] = p1
        d_1628_v23_: str
        d_1628_v23_ = _dafny.CodePoint('e')
        d_1629_v24_: _dafny.Seq
        d_1629_v24_ = _dafny.SeqWithoutIsStrInference([p0, 325])
        d_1630_v25_: _dafny.Map
        d_1630_v25_ = _dafny.Map({d_1628_v23_: p1})
        d_1631_v26_: _dafny.MultiSet
        d_1631_v26_ = _dafny.MultiSet([(default__.fm49(d_1628_v23_, d_1629_v24_, 813, d_1628_v23_, globalState)).cf60, ((d_1630_v25_)[d_1628_v23_] if (d_1628_v23_) in (d_1630_v25_) else p1), p1])
        index279_ = default__.safeIndex(304, (p2).length(0))
        (p2)[index279_] = (d_1631_v26_).cardinality
        index280_ = default__.safeIndex(13, (d_1627_v22_).length(0))
        index281_ = default__.safeIndex(860, (d_1627_v22_).length(0))
        index282_ = default__.safeIndex(304, (p2).length(0))
        rhs279_ = p1
        rhs280_ = (p0) < (p0)
        rhs281_ = d_1631_v26_
        rhs282_ = p3
        lhs261_ = d_1627_v22_
        lhs262_ = default__.safeIndex(13, (d_1627_v22_).length(0))
        lhs263_ = d_1627_v22_
        lhs264_ = default__.safeIndex(860, (d_1627_v22_).length(0))
        lhs265_ = p2
        lhs266_ = default__.safeIndex(304, (p2).length(0))
        lhs261_[lhs262_] = rhs279_
        lhs263_[lhs264_] = rhs280_
        d_1631_v26_ = rhs281_
        lhs265_[lhs266_] = rhs282_
        r0 = ((p1) == (p1)) or (p1)
        r1 = default__.safeModuloInt((p0 if (d_1627_v22_)[default__.safeIndex(13, (d_1627_v22_).length(0))] else (p2)[default__.safeIndex(304, (p2).length(0))]), 133)
        d_1632_v27_: _dafny.Array
        def lambda89_(d_1633_v22_, d_1634_v23_, d_1635_p1_):
            def lambda90_(d_1636_i1_):
                return (_dafny.Map({(d_1633_v22_)[default__.safeIndex(13, (d_1633_v22_).length(0))]: D3_DC6(_dafny.SeqWithoutIsStrInference([d_1634_v23_ for d_1637_i2_ in range(default__.abs(490))]), _dafny.CodePoint('b'), d_1635_p1_)})) | (_dafny.Map({d_1635_p1_: D3_DC6(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "whrnsnix")), d_1634_v23_, (d_1633_v22_)[default__.safeIndex(13, (d_1633_v22_).length(0))])}))

            return lambda90_

        init49_ = lambda89_(d_1627_v22_, d_1628_v23_, p1)
        nw270_ = _dafny.Array(None, 8)
        for i0_49_ in range(nw270_.length(0)):
            nw270_[i0_49_] = init49_(i0_49_)
        d_1632_v27_ = nw270_
        r2 = d_1632_v27_
        return r0, r1, r2

    def m3(self, p0, globalState):
        d_1638_v0_: int
        d_1638_v0_ = 685
        d_1639_v1_: C5
        nw271_ = C5()
        nw271_.ctor__((d_1638_v0_) <= (d_1638_v0_))
        d_1639_v1_ = nw271_
        d_1640_v2_: C3
        nw272_ = C3()
        nw272_.ctor__()
        d_1640_v2_ = nw272_
        (globalState).f0 = True
        d_1641_v3_: _dafny.Map
        d_1641_v3_ = _dafny.Map({(d_1639_v1_).f24: False})
        d_1642_v4_: _dafny.Map
        d_1642_v4_ = _dafny.Map({default__.fm2(d_1641_v3_, globalState): default__.fm0(d_1638_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "brysun"))), d_1638_v0_, (d_1639_v1_).f24, globalState)})
        index283_ = default__.safeIndex(277, (p0).length(0))
        (p0)[index283_] = len(d_1642_v4_)
        d_1643_v5_: D19
        d_1643_v5_ = D19_DC48(d_1638_v0_, (d_1639_v1_).f24)
        d_1644_v6_: _dafny.Seq
        d_1644_v6_ = _dafny.SeqWithoutIsStrInference([(d_1639_v1_).f24, (d_1639_v1_).f24])
        d_1645_v7_: _dafny.MultiSet
        d_1645_v7_ = _dafny.MultiSet([d_1641_v3_])
        d_1646_v8_: _dafny.Map
        d_1646_v8_ = _dafny.Map({(d_1645_v7_).cardinality: d_1638_v0_})
        d_1647_v9_: _dafny.Set
        d_1647_v9_ = _dafny.Set({d_1646_v8_})
        d_1648_v10_: _dafny.Map
        d_1648_v10_ = _dafny.Map({len(d_1647_v9_): (d_1639_v1_).f24})
        d_1649_v11_: _dafny.Array
        nw273_ = _dafny.Array(None, 29)
        nw273_[int(0)] = (not((d_1639_v1_).f24) if not((d_1639_v1_).f24) else (d_1639_v1_).f24)
        nw273_[int(1)] = (d_1643_v5_).cf66
        nw273_[int(2)] = (d_1639_v1_).f24
        nw273_[int(3)] = (d_1639_v1_).f24
        nw273_[int(4)] = (d_1639_v1_).f24
        nw273_[int(5)] = (d_1639_v1_).f24
        nw273_[int(6)] = not((d_1639_v1_).f24)
        nw273_[int(7)] = (d_1639_v1_).f24
        nw273_[int(8)] = ((d_1639_v1_).f24 if ((d_1641_v3_)[(d_1639_v1_).f24] if ((d_1639_v1_).f24) in (d_1641_v3_) else (d_1639_v1_).f24) else False)
        nw273_[int(9)] = (d_1639_v1_).f24
        nw273_[int(10)] = (d_1639_v1_).f24
        nw273_[int(11)] = (d_1639_v1_).f24
        nw273_[int(12)] = (d_1639_v1_).f24
        nw273_[int(13)] = (d_1639_v1_).f24
        nw273_[int(14)] = (d_1639_v1_).f24
        nw273_[int(15)] = (d_1639_v1_).f24
        nw273_[int(16)] = False
        nw273_[int(17)] = not((d_1639_v1_).f24)
        nw273_[int(18)] = not((d_1644_v6_) <= (d_1644_v6_))
        nw273_[int(19)] = (d_1639_v1_).f24
        nw273_[int(20)] = (d_1639_v1_).f24
        nw273_[int(21)] = (d_1639_v1_).f24
        nw273_[int(22)] = (d_1639_v1_).f24
        nw273_[int(23)] = ((d_1639_v1_).f24 if False else (d_1639_v1_).f24)
        nw273_[int(24)] = (not((d_1639_v1_).f24)) or ((d_1639_v1_).f24)
        nw273_[int(25)] = ((d_1648_v10_)[d_1638_v0_] if (d_1638_v0_) in (d_1648_v10_) else False)
        nw273_[int(26)] = not(False)
        nw273_[int(27)] = False
        nw273_[int(28)] = (len(_dafny.SeqWithoutIsStrInference([(d_1639_v1_).f24]))) <= (d_1638_v0_)
        d_1649_v11_ = nw273_
        d_1650_v12_: str
        d_1650_v12_ = _dafny.CodePoint('n')
        d_1651_v13_: D20
        d_1651_v13_ = D20_DC49(_dafny.Map({d_1638_v0_: d_1650_v12_}))
        d_1652_v14_: _dafny.Map
        d_1652_v14_ = _dafny.Map({not((d_1639_v1_).f24): (d_1651_v13_).cf67})
        d_1653_v15_: _dafny.Seq
        d_1653_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "snhrb"))
        index284_ = default__.safeIndex(277, (p0).length(0))
        rhs283_ = default__.safeDivisionInt((d_1638_v0_) * (229), (0) - ((d_1639_v1_).fm8((d_1644_v6_)[default__.safeIndex(d_1638_v0_, len(d_1644_v6_))], ((d_1652_v14_)[(d_1639_v1_).f24] if ((d_1639_v1_).f24) in (d_1652_v14_) else _dafny.Map({d_1638_v0_: d_1650_v12_})), d_1653_v15_, globalState)))
        rhs284_ = ((d_1641_v3_)[(d_1639_v1_).f24] if ((d_1639_v1_).f24) in (d_1641_v3_) else (self).fm6(d_1638_v0_, globalState))
        rhs285_ = (d_1638_v0_) - (d_1638_v0_)
        rhs286_ = d_1649_v11_
        lhs267_ = p0
        lhs268_ = default__.safeIndex(277, (p0).length(0))
        lhs269_ = globalState
        lhs270_ = globalState
        lhs267_[lhs268_] = rhs283_
        lhs269_.f2 = rhs284_
        lhs270_.f16 = rhs285_
        d_1649_v11_ = rhs286_
        d_1654_i0_: int
        d_1654_i0_ = 0
        with _dafny.label("9"):
            while (d_1639_v1_).f24:
                with _dafny.c_label("9"):
                    if (d_1654_i0_) >= (100):
                        raise _dafny.Break("9")
                    d_1654_i0_ = (d_1654_i0_) + (1)
                    (globalState).f16 = default__.safeModuloInt((d_1638_v0_ if (d_1639_v1_).f24 else d_1638_v0_), d_1638_v0_)
                    d_1648_v10_ = (d_1648_v10_).set(145, (d_1639_v1_).f24)
                    d_1655_v16_: _dafny.Array
                    def lambda91_(d_1656_v1_, d_1657_v6_):
                        def lambda92_(d_1658_i1_):
                            return D16_DC36((d_1656_v1_).f24, _dafny.MultiSet(d_1657_v6_), (d_1656_v1_).f24)

                        return lambda92_

                    init50_ = lambda91_(d_1639_v1_, d_1644_v6_)
                    nw274_ = _dafny.Array(None, 10)
                    for i0_50_ in range(nw274_.length(0)):
                        nw274_[i0_50_] = init50_(i0_50_)
                    d_1655_v16_ = nw274_
                    d_1659_v17_: D17
                    d_1659_v17_ = D17_DC38(d_1655_v16_)
                    d_1660_v18_: _dafny.Map
                    d_1660_v18_ = _dafny.Map({d_1659_v17_: d_1642_v4_})
                    d_1661_v19_: _dafny.Seq
                    d_1661_v19_ = _dafny.SeqWithoutIsStrInference([d_1660_v18_, d_1660_v18_])
                    d_1662_v20_: _dafny.Map
                    d_1662_v20_ = _dafny.Map({d_1638_v0_: d_1642_v4_})
                    d_1663_v21_: _dafny.Map
                    d_1663_v21_ = _dafny.Map({(d_1639_v1_).f24: default__.fm35((d_1639_v1_).f24, len(d_1653_v15_), globalState)})
                    d_1664_v22_: _dafny.MultiSet
                    d_1664_v22_ = _dafny.MultiSet([True, (d_1639_v1_).f24, (d_1639_v1_).f24, (d_1639_v1_).f24, (d_1639_v1_).f24])
                    d_1665_v23_: _dafny.Array
                    nw275_ = _dafny.Array(None, 29)
                    nw275_[int(0)] = (d_1660_v18_) | (d_1660_v18_)
                    nw275_[int(1)] = d_1660_v18_
                    nw275_[int(2)] = _dafny.Map({d_1659_v17_: (d_1642_v4_).set((d_1639_v1_).f24, (p0)[default__.safeIndex(277, (p0).length(0))])})
                    nw275_[int(3)] = d_1660_v18_
                    nw275_[int(4)] = (d_1660_v18_).set(d_1659_v17_, d_1642_v4_)
                    nw275_[int(5)] = (d_1660_v18_) | (d_1660_v18_)
                    nw275_[int(6)] = (d_1661_v19_)[default__.safeIndex(d_1638_v0_, len(d_1661_v19_))]
                    nw275_[int(7)] = _dafny.Map({D17_DC38(d_1655_v16_): d_1642_v4_})
                    nw275_[int(8)] = d_1660_v18_
                    nw275_[int(9)] = (d_1660_v18_) | (d_1660_v18_)
                    nw275_[int(10)] = d_1660_v18_
                    nw275_[int(11)] = _dafny.Map({d_1659_v17_: d_1642_v4_})
                    nw275_[int(12)] = d_1660_v18_
                    nw275_[int(13)] = (_dafny.Map({d_1659_v17_: d_1642_v4_}) if (d_1639_v1_).f24 else _dafny.Map({d_1659_v17_: ((d_1662_v20_)[(p0)[default__.safeIndex(277, (p0).length(0))]] if ((p0)[default__.safeIndex(277, (p0).length(0))]) in (d_1662_v20_) else d_1642_v4_)}))
                    nw275_[int(14)] = d_1660_v18_
                    nw275_[int(15)] = d_1660_v18_
                    nw275_[int(16)] = d_1660_v18_
                    nw275_[int(17)] = d_1660_v18_
                    nw275_[int(18)] = d_1660_v18_
                    nw275_[int(19)] = ((_dafny.Map({d_1659_v17_: d_1642_v4_})).set(d_1659_v17_, (d_1642_v4_).set((d_1639_v1_).f24, (p0)[default__.safeIndex(277, (p0).length(0))]))) | (d_1660_v18_)
                    nw275_[int(20)] = (d_1660_v18_).set(D17_DC38(d_1655_v16_), ((d_1663_v21_)[(d_1639_v1_).f24] if ((d_1639_v1_).f24) in (d_1663_v21_) else d_1642_v4_))
                    nw275_[int(21)] = d_1660_v18_
                    nw275_[int(22)] = d_1660_v18_
                    nw275_[int(23)] = _dafny.Map({d_1659_v17_: _dafny.Map({(d_1639_v1_).f24: (0) - ((d_1664_v22_).cardinality)})})
                    nw275_[int(24)] = (_dafny.Map({d_1659_v17_: d_1642_v4_})) | (d_1660_v18_)
                    nw275_[int(25)] = (d_1660_v18_).set(d_1659_v17_, d_1642_v4_)
                    nw275_[int(26)] = (d_1660_v18_) | (d_1660_v18_)
                    nw275_[int(27)] = (d_1660_v18_).set(D17_DC38((d_1659_v17_).cf51), d_1642_v4_)
                    nw275_[int(28)] = d_1660_v18_
                    d_1665_v23_ = nw275_
                    index285_ = default__.safeIndex(907, (d_1665_v23_).length(0))
                    (d_1665_v23_)[index285_] = d_1660_v18_
                    index286_ = default__.safeIndex(907, (d_1665_v23_).length(0))
                    (d_1665_v23_)[index286_] = (d_1661_v19_)[default__.safeIndex((p0)[default__.safeIndex(277, (p0).length(0))], len(d_1661_v19_))]
                    (globalState).f19 = d_1653_v15_
                    pass
            pass
        hi12_ = len(d_1653_v15_)
        for d_1666_i2_ in range(((p0)[default__.safeIndex(277, (p0).length(0))]) * (211), hi12_):
            d_1667_v24_: _dafny.Map
            d_1667_v24_ = _dafny.Map({(d_1639_v1_).f24: _dafny.SeqWithoutIsStrInference([len(d_1653_v15_) for d_1668_i3_ in range(default__.abs(318))])})
            d_1669_v25_: _dafny.Map
            d_1669_v25_ = _dafny.Map({(d_1639_v1_).f24: d_1667_v24_})
            d_1670_v26_: _dafny.Set
            d_1670_v26_ = _dafny.Set({(d_1639_v1_).f24})
            d_1671_v27_: _dafny.Seq
            d_1671_v27_ = _dafny.SeqWithoutIsStrInference([len(d_1642_v4_), len(d_1670_v26_), (p0)[default__.safeIndex(277, (p0).length(0))]])
            d_1672_v28_: _dafny.Array
            nw276_ = _dafny.Array(None, 14)
            nw276_[int(0)] = d_1667_v24_
            nw276_[int(1)] = d_1667_v24_
            nw276_[int(2)] = ((d_1669_v25_)[(d_1639_v1_).f24] if ((d_1639_v1_).f24) in (d_1669_v25_) else d_1667_v24_)
            nw276_[int(3)] = d_1667_v24_
            nw276_[int(4)] = _dafny.Map({(d_1639_v1_).f24: d_1671_v27_})
            nw276_[int(5)] = (d_1667_v24_) | (d_1667_v24_)
            nw276_[int(6)] = (d_1667_v24_) | (_dafny.Map({(d_1639_v1_).f24: d_1671_v27_}))
            nw276_[int(7)] = (d_1667_v24_) | (d_1667_v24_)
            nw276_[int(8)] = _dafny.Map({(d_1639_v1_).f24: d_1671_v27_})
            nw276_[int(9)] = (default__.fm50(d_1666_i2_, d_1666_i2_, globalState)) | (d_1667_v24_)
            nw276_[int(10)] = d_1667_v24_
            nw276_[int(11)] = (d_1667_v24_) | (_dafny.Map({(d_1639_v1_).f24: d_1671_v27_}))
            nw276_[int(12)] = d_1667_v24_
            nw276_[int(13)] = d_1667_v24_
            d_1672_v28_ = nw276_
            index287_ = default__.safeIndex(498, (d_1672_v28_).length(0))
            (d_1672_v28_)[index287_] = d_1667_v24_
            d_1673_v29_: _dafny.Map
            d_1673_v29_ = _dafny.Map({d_1653_v15_: d_1646_v8_})
            index288_ = default__.safeIndex(498, (d_1672_v28_).length(0))
            rhs287_ = d_1650_v12_
            rhs288_ = (d_1639_v1_).f24
            rhs289_ = _dafny.Map({not((d_1639_v1_).f24): d_1671_v27_})
            rhs290_ = not((d_1639_v1_).f24)
            rhs291_ = (0) - ((0) - (len((d_1673_v29_) | ((_dafny.Map({(d_1653_v15_).set(default__.safeIndex((p0)[default__.safeIndex(277, (p0).length(0))], len(d_1653_v15_)), _dafny.CodePoint('i')): d_1646_v8_})) | (d_1673_v29_)))))
            lhs271_ = globalState
            lhs272_ = d_1672_v28_
            lhs273_ = default__.safeIndex(498, (d_1672_v28_).length(0))
            lhs274_ = globalState
            lhs275_ = globalState
            d_1650_v12_ = rhs287_
            lhs271_.f2 = rhs288_
            lhs272_[lhs273_] = rhs289_
            lhs274_.f21 = rhs290_
            lhs275_.f7 = rhs291_
            d_1674_v30_: _dafny.Seq
            d_1674_v30_ = _dafny.SeqWithoutIsStrInference([d_1648_v10_])
            d_1675_v31_: C1
            nw277_ = C1()
            nw277_.ctor__((d_1666_i2_ if (d_1639_v1_).f24 else len(d_1653_v15_)), (default__.fm51(_dafny.CodePoint('i'), (p0)[default__.safeIndex(277, (p0).length(0))], globalState)) + (d_1674_v30_))
            d_1675_v31_ = nw277_
            d_1676_v32_: _dafny.Map
            d_1676_v32_ = _dafny.Map({default__.fm34(globalState): len(d_1641_v3_)})
            d_1676_v32_ = (d_1676_v32_).set(d_1650_v12_, d_1675_v31_.f29)
            index289_ = default__.safeIndex(760, (d_1649_v11_).length(0))
            (d_1649_v11_)[index289_] = (d_1639_v1_).f24
            d_1677_v33_: _dafny.Seq
            d_1677_v33_ = _dafny.SeqWithoutIsStrInference([p0])
            d_1678_v34_: D6
            d_1678_v34_ = D6_DC12((d_1677_v33_).set(default__.safeIndex((p0)[default__.safeIndex(277, (p0).length(0))], len(d_1677_v33_)), p0))
            d_1679_v35_: D6
            d_1679_v35_ = D6_DC14(d_1678_v34_)
            d_1680_v36_: D6
            d_1680_v36_ = D6_DC14(d_1678_v34_)
            d_1681_v37_: D6
            d_1681_v37_ = D6_DC14(d_1679_v35_)
            d_1682_v38_: _dafny.Map
            d_1682_v38_ = _dafny.Map({d_1675_v31_.f29: p0})
            index290_ = default__.safeIndex(760, (d_1649_v11_).length(0))
            rhs292_ = ((d_1675_v31_).fm29(globalState)) + ((len(d_1648_v10_)) + (d_1666_i2_))
            rhs293_ = (((0) - (d_1675_v31_.f29) if (d_1639_v1_).f24 else d_1638_v0_)) in (d_1682_v38_)
            rhs294_ = D6_DC14(d_1678_v34_)
            rhs295_ = d_1653_v15_
            rhs296_ = d_1675_v31_.f29
            lhs276_ = d_1675_v31_
            lhs277_ = d_1649_v11_
            lhs278_ = default__.safeIndex(760, (d_1649_v11_).length(0))
            lhs279_ = globalState
            lhs280_ = globalState
            lhs276_.f29 = rhs292_
            lhs277_[lhs278_] = rhs293_
            d_1681_v37_ = rhs294_
            lhs279_.f20 = rhs295_
            lhs280_.f7 = rhs296_

    def m1(self, p0, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: bool = False
        d_1683_v0_: bool
        d_1683_v0_ = False
        d_1684_v1_: _dafny.Map
        d_1684_v1_ = _dafny.Map({d_1683_v0_: d_1683_v0_})
        d_1684_v1_ = (d_1684_v1_).set(d_1683_v0_, default__.fm2((d_1684_v1_).set(d_1683_v0_, not(d_1683_v0_)), globalState))
        d_1685_i0_: int
        d_1685_i0_ = 0
        with _dafny.label("10"):
            while d_1683_v0_:
                with _dafny.c_label("10"):
                    if (d_1685_i0_) >= (100):
                        raise _dafny.Break("10")
                    d_1685_i0_ = (d_1685_i0_) + (1)
                    d_1686_v2_: str
                    d_1686_v2_ = _dafny.CodePoint('t')
                    rhs297_ = d_1683_v0_
                    rhs298_ = _dafny.CodePoint('k')
                    lhs281_ = globalState
                    lhs281_.f0 = rhs297_
                    d_1686_v2_ = rhs298_
                    if True:
                        d_1687_v3_: _dafny.MultiSet
                        d_1687_v3_ = _dafny.MultiSet([len(d_1684_v1_)])
                        (globalState).f2 = ((d_1687_v3_).isdisjoint(d_1687_v3_)) or (d_1683_v0_)
                        d_1688_v5_: _dafny.Set
                        def iife119_():
                            coll67_ = _dafny.Map()
                            compr_67_: int
                            for compr_67_ in (d_1687_v3_).Elements:
                                d_1689_v4_: int = compr_67_
                                if (d_1689_v4_) in (d_1687_v3_):
                                    coll67_[(d_1689_v4_) + (p0)] = _dafny.SeqWithoutIsStrInference([p0 for d_1690_i1_ in range(default__.abs(351))])
                            return _dafny.Map(coll67_)
                        d_1688_v5_ = _dafny.Set({len(iife119_()
                        ), p0})
                        d_1691_v6_: _dafny.MultiSet
                        d_1691_v6_ = _dafny.MultiSet([d_1683_v0_, (d_1688_v5_).issubset(d_1688_v5_)])
                        d_1692_v7_: _dafny.Map
                        d_1692_v7_ = _dafny.Map({p0: (d_1691_v6_).set(d_1683_v0_, default__.abs(p0))})
                        d_1693_v8_: _dafny.Array
                        nw278_ = _dafny.Array(int(0), 16)
                        d_1693_v8_ = nw278_
                        d_1694_v9_: _dafny.Seq
                        d_1694_v9_ = _dafny.SeqWithoutIsStrInference([d_1683_v0_])
                        d_1695_v10_: _dafny.Map
                        d_1695_v10_ = _dafny.Map({d_1693_v8_: d_1694_v9_})
                        d_1691_v6_ = (((d_1692_v7_)[p0] if (p0) in (d_1692_v7_) else d_1691_v6_)) | (_dafny.MultiSet(((d_1695_v10_)[d_1693_v8_] if (d_1693_v8_) in (d_1695_v10_) else d_1694_v9_)))
                        d_1696_v11_: C0
                        nw279_ = C0()
                        nw279_.ctor__(d_1683_v0_)
                        d_1696_v11_ = nw279_
                        d_1696_v11_ = d_1696_v11_
                        d_1697_v12_: _dafny.Array
                        nw280_ = _dafny.Array(None, 1)
                        d_1697_v12_ = nw280_
                        d_1698_v13_: _dafny.Seq
                        d_1698_v13_ = _dafny.SeqWithoutIsStrInference([p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eoog")))])
                        d_1699_v14_: _dafny.Seq
                        d_1699_v14_ = _dafny.SeqWithoutIsStrInference([d_1698_v13_])
                        d_1700_v15_: _dafny.Map
                        d_1700_v15_ = _dafny.Map({d_1694_v9_: (d_1699_v14_)[default__.safeIndex(p0, len(d_1699_v14_))]})
                        d_1701_v16_: D21
                        d_1701_v16_ = D21_DC52(d_1700_v15_)
                        d_1702_v17_: C2
                        nw281_ = C2()
                        nw281_.ctor__(len((d_1701_v16_).cf71))
                        d_1702_v17_ = nw281_
                        index291_ = default__.safeIndex(504, (d_1697_v12_).length(0))
                        (d_1697_v12_)[index291_] = d_1702_v17_
                        index292_ = default__.safeIndex(504, (d_1697_v12_).length(0))
                        (d_1697_v12_)[index292_] = d_1702_v17_
                        (globalState).f19 = _dafny.SeqWithoutIsStrInference([d_1686_v2_ for d_1703_i2_ in range(default__.abs(-484))])
                    elif True:
                        d_1704_v18_: _dafny.Array
                        nw282_ = _dafny.Array(int(0), 20)
                        d_1704_v18_ = nw282_
                        r0 = d_1704_v18_
                        d_1705_v19_: _dafny.Set
                        d_1705_v19_ = _dafny.Set({p0, p0})
                        d_1706_v20_: D14
                        d_1706_v20_ = D14_DC28(d_1705_v19_)
                        d_1707_v21_: _dafny.Array
                        nw283_ = _dafny.Array(None, 25)
                        nw283_[int(0)] = d_1683_v0_
                        nw283_[int(1)] = d_1683_v0_
                        nw283_[int(2)] = not(d_1683_v0_)
                        nw283_[int(3)] = d_1683_v0_
                        nw283_[int(4)] = not(d_1683_v0_)
                        nw283_[int(5)] = (d_1705_v19_).ispropersubset((d_1706_v20_).cf37)
                        nw283_[int(6)] = (p0) > (p0)
                        nw283_[int(7)] = d_1683_v0_
                        nw283_[int(8)] = d_1683_v0_
                        nw283_[int(9)] = (not(True)) == (False)
                        nw283_[int(10)] = True
                        nw283_[int(11)] = not(True)
                        nw283_[int(12)] = d_1683_v0_
                        nw283_[int(13)] = d_1683_v0_
                        nw283_[int(14)] = not (d_1683_v0_) or (d_1683_v0_)
                        nw283_[int(15)] = d_1683_v0_
                        nw283_[int(16)] = d_1683_v0_
                        nw283_[int(17)] = d_1683_v0_
                        nw283_[int(18)] = d_1683_v0_
                        nw283_[int(19)] = False
                        nw283_[int(20)] = d_1683_v0_
                        nw283_[int(21)] = d_1683_v0_
                        nw283_[int(22)] = d_1683_v0_
                        nw283_[int(23)] = d_1683_v0_
                        nw283_[int(24)] = d_1683_v0_
                        d_1707_v21_ = nw283_
                        index293_ = default__.safeIndex(680, (d_1707_v21_).length(0))
                        (d_1707_v21_)[index293_] = d_1683_v0_
                        index294_ = default__.safeIndex(680, (d_1707_v21_).length(0))
                        (d_1707_v21_)[index294_] = default__.fm2(d_1684_v1_, globalState)
                        d_1708_v22_: T0
                        nw284_ = C3()
                        nw284_.ctor__()
                        d_1708_v22_ = nw284_
                        d_1709_v23_: _dafny.Map
                        d_1709_v23_ = _dafny.Map({d_1683_v0_: d_1708_v22_})
                        d_1710_v24_: _dafny.Seq
                        d_1710_v24_ = _dafny.SeqWithoutIsStrInference([d_1709_v23_])
                        d_1711_v25_: _dafny.Seq
                        d_1711_v25_ = _dafny.SeqWithoutIsStrInference([(d_1710_v24_)[default__.safeIndex(p0, len(d_1710_v24_))]])
                        d_1712_v26_: _dafny.Seq
                        d_1712_v26_ = _dafny.SeqWithoutIsStrInference([False])
                        d_1713_v27_: _dafny.Map
                        d_1713_v27_ = _dafny.Map({d_1712_v26_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ulcd"))})
                        d_1714_v28_: _dafny.Seq
                        d_1714_v28_ = _dafny.SeqWithoutIsStrInference([default__.fm27((d_1707_v21_)[default__.safeIndex(680, (d_1707_v21_).length(0))], (d_1707_v21_)[default__.safeIndex(680, (d_1707_v21_).length(0))], globalState)])
                        d_1715_v29_: _dafny.Set
                        d_1715_v29_ = _dafny.Set({d_1714_v28_, d_1714_v28_, d_1714_v28_})
                        default__.m0((p0) == (p0), default__.safeDivisionInt(len(d_1711_v25_), p0), d_1713_v27_, len((d_1715_v29_) - (d_1715_v29_)), globalState)
                        d_1716_v33_: _dafny.Seq
                        d_1716_v33_ = _dafny.SeqWithoutIsStrInference([p0])
                        d_1717_v34_: _dafny.Map
                        d_1717_v34_ = _dafny.Map({p0: 225})
                        d_1718_v35_: _dafny.Seq
                        def iife120_():
                            coll68_ = _dafny.Map()
                            compr_68_: int
                            for compr_68_ in (d_1716_v33_).Elements:
                                d_1719_v32_: int = compr_68_
                                if (d_1719_v32_) in (d_1716_v33_):
                                    coll68_[default__.safeModuloInt(d_1719_v32_, p0)] = 251
                            return _dafny.Map(coll68_)
                        d_1718_v35_ = _dafny.SeqWithoutIsStrInference([iife120_()
                        , d_1717_v34_])
                        d_1720_v36_: _dafny.Map
                        def iife121_():
                            def iife123_():
                                coll71_ = _dafny.Map()
                                compr_71_: _dafny.Map
                                for compr_71_ in (d_1718_v35_).Elements:
                                    d_1721_v31_: _dafny.Map = compr_71_
                                    if (d_1721_v31_) in (d_1718_v35_):
                                        coll71_[d_1721_v31_] = p0
                                return _dafny.Map(coll71_)
                            coll69_ = _dafny.Map()
                            def iife122_():
                                coll70_ = _dafny.Map()
                                compr_70_: _dafny.Map
                                for compr_70_ in (d_1718_v35_).Elements:
                                    d_1721_v31_: _dafny.Map = compr_70_
                                    if (d_1721_v31_) in (d_1718_v35_):
                                        coll70_[d_1721_v31_] = p0
                                return _dafny.Map(coll70_)
                            compr_69_: _dafny.Map
                            for compr_69_ in (iife122_()
                            ).keys.Elements:
                                d_1722_v30_: _dafny.Map = compr_69_
                                if (d_1722_v30_) in (iife123_()
                                ):
                                    coll69_[d_1722_v30_] = d_1683_v0_
                            return _dafny.Map(coll69_)
                        d_1720_v36_ = _dafny.Map({iife121_()
                        : True})
                        d_1723_v38_: _dafny.Map
                        def iife124_():
                            coll72_ = _dafny.Map()
                            compr_72_: int
                            for compr_72_ in _dafny.IntegerRange(83, 288):
                                d_1724_v37_: int = compr_72_
                                if ((83) <= (d_1724_v37_)) and ((d_1724_v37_) < (288)):
                                    coll72_[(d_1724_v37_) + (p0)] = p0
                            return _dafny.Map(coll72_)
                        d_1723_v38_ = _dafny.Map({iife124_()
                        : (d_1707_v21_)[default__.safeIndex(680, (d_1707_v21_).length(0))]})
                        d_1720_v36_ = (d_1720_v36_).set(d_1723_v38_, d_1683_v0_)
                        d_1725_v39_: C2
                        nw285_ = C2()
                        nw285_.ctor__(p0)
                        d_1725_v39_ = nw285_
                    d_1726_v40_: _dafny.Map
                    d_1726_v40_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1683_v0_]): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))})
                    d_1727_v41_: _dafny.Seq
                    d_1727_v41_ = _dafny.SeqWithoutIsStrInference([d_1726_v40_])
                    default__.m0(not (d_1683_v0_) or (d_1683_v0_), p0, (d_1727_v41_)[default__.safeIndex(p0, len(d_1727_v41_))], (p0) - (p0), globalState)
                    d_1728_v42_: _dafny.Seq
                    d_1728_v42_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))
                    d_1729_v43_: _dafny.Map
                    d_1729_v43_ = _dafny.Map({not(d_1683_v0_): d_1728_v42_})
                    (globalState).f20 = (d_1728_v42_) + (((d_1729_v43_)[d_1683_v0_] if (d_1683_v0_) in (d_1729_v43_) else d_1728_v42_))
                    pass
            pass
        d_1730_v44_: _dafny.Seq
        d_1730_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aapa"))
        d_1731_v45_: _dafny.Map
        d_1731_v45_ = _dafny.Map({p0: d_1683_v0_})
        d_1732_v46_: _dafny.Map
        d_1732_v46_ = _dafny.Map({len(d_1731_v45_): p0})
        hi13_ = p0
        for d_1733_i3_ in range(default__.safeDivisionInt(len(d_1730_v44_), default__.fm0(p0, len(d_1732_v46_), p0, d_1683_v0_, globalState)), hi13_):
            d_1734_v47_: T0
            nw286_ = C3()
            nw286_.ctor__()
            d_1734_v47_ = nw286_
            d_1734_v47_ = d_1734_v47_
            d_1735_v48_: _dafny.Array
            def lambda93_(d_1736_v44_):
                def lambda94_(d_1737_i4_):
                    return (_dafny.CodePoint('n')) in (d_1736_v44_)

                return lambda94_

            init51_ = lambda93_(d_1730_v44_)
            nw287_ = _dafny.Array(None, 18)
            for i0_51_ in range(nw287_.length(0)):
                nw287_[i0_51_] = init51_(i0_51_)
            d_1735_v48_ = nw287_
            index295_ = default__.safeIndex(567, (d_1735_v48_).length(0))
            (d_1735_v48_)[index295_] = d_1683_v0_
            d_1738_v49_: _dafny.Array
            nw288_ = _dafny.Array(None, 19)
            d_1738_v49_ = nw288_
            d_1739_v50_: _dafny.Map
            d_1739_v50_ = _dafny.Map({d_1738_v49_: d_1683_v0_})
            index296_ = default__.safeIndex(567, (d_1735_v48_).length(0))
            rhs299_ = d_1733_i3_
            rhs300_ = (d_1739_v50_) != ((d_1739_v50_) | ((D22_DC56(_dafny.Map({d_1738_v49_: False}))).cf80))
            lhs282_ = globalState
            lhs283_ = d_1735_v48_
            lhs284_ = default__.safeIndex(567, (d_1735_v48_).length(0))
            lhs282_.f16 = rhs299_
            lhs283_[lhs284_] = rhs300_
            d_1740_v51_: _dafny.Map
            d_1740_v51_ = _dafny.Map({d_1683_v0_: d_1733_i3_})
            d_1741_v52_: _dafny.MultiSet
            d_1741_v52_ = _dafny.MultiSet([d_1733_i3_])
            d_1742_v53_: _dafny.Set
            d_1742_v53_ = _dafny.Set({p0, p0, p0, p0})
            d_1743_v54_: _dafny.Seq
            d_1743_v54_ = _dafny.SeqWithoutIsStrInference([(d_1735_v48_)[default__.safeIndex(567, (d_1735_v48_).length(0))], (d_1735_v48_)[default__.safeIndex(567, (d_1735_v48_).length(0))]])
            d_1744_v55_: D4
            d_1744_v55_ = D4_DC9(d_1733_i3_)
            d_1745_v56_: _dafny.Seq
            d_1745_v56_ = _dafny.SeqWithoutIsStrInference([p0, 515])
            d_1746_v57_: D4
            d_1746_v57_ = D4_DC10(len(d_1743_v54_), len(d_1684_v1_), d_1683_v0_, (d_1735_v48_)[default__.safeIndex(567, (d_1735_v48_).length(0))], d_1733_i3_)
            d_1747_v58_: str
            d_1747_v58_ = _dafny.CodePoint('k')
            d_1748_v59_: _dafny.MultiSet
            d_1748_v59_ = _dafny.MultiSet([d_1747_v58_, d_1747_v58_, d_1747_v58_, d_1747_v58_])
            d_1749_v60_: _dafny.Set
            d_1749_v60_ = _dafny.Set({d_1683_v0_, d_1683_v0_, d_1683_v0_})
            d_1750_v61_: _dafny.Array
            nw289_ = _dafny.Array(None, 24)
            nw289_[int(0)] = d_1733_i3_
            nw289_[int(1)] = d_1733_i3_
            nw289_[int(2)] = len(d_1732_v46_)
            nw289_[int(3)] = (len(d_1740_v51_)) + ((d_1741_v52_).cardinality)
            nw289_[int(4)] = p0
            nw289_[int(5)] = default__.safeModuloInt(p0, len(d_1742_v53_))
            nw289_[int(6)] = (p0) * (len(d_1743_v54_))
            nw289_[int(7)] = 848
            nw289_[int(8)] = p0
            nw289_[int(9)] = (d_1744_v55_).cf11
            nw289_[int(10)] = (len(d_1732_v46_)) - (d_1733_i3_)
            nw289_[int(11)] = (d_1745_v56_)[default__.safeIndex(d_1733_i3_, len(d_1745_v56_))]
            nw289_[int(12)] = (d_1746_v57_).cf12
            nw289_[int(13)] = p0
            nw289_[int(14)] = (p0) - ((d_1748_v59_).cardinality)
            nw289_[int(15)] = ((d_1741_v52_)[d_1733_i3_] if (d_1733_i3_) in (d_1741_v52_) else p0)
            nw289_[int(16)] = d_1733_i3_
            nw289_[int(17)] = d_1733_i3_
            nw289_[int(18)] = p0
            nw289_[int(19)] = ((d_1732_v46_)[len(d_1749_v60_)] if (len(d_1749_v60_)) in (d_1732_v46_) else p0)
            nw289_[int(20)] = d_1733_i3_
            nw289_[int(21)] = d_1733_i3_
            nw289_[int(22)] = default__.fm0(p0, 827, p0, not((d_1735_v48_)[default__.safeIndex(567, (d_1735_v48_).length(0))]), globalState)
            nw289_[int(23)] = -744
            d_1750_v61_ = nw289_
            index297_ = default__.safeIndex(412, (d_1750_v61_).length(0))
            (d_1750_v61_)[index297_] = d_1733_i3_
            index298_ = default__.safeIndex(412, (d_1750_v61_).length(0))
            (d_1750_v61_)[index298_] = -85
            d_1740_v51_ = (d_1740_v51_).set(not(True), (d_1750_v61_)[default__.safeIndex(412, (d_1750_v61_).length(0))])
        d_1751_v62_: C0
        nw290_ = C0()
        nw290_.ctor__(d_1683_v0_)
        d_1751_v62_ = nw290_
        (globalState).f7 = p0
        d_1752_i5_: int
        d_1752_i5_ = 0
        with _dafny.label("11"):
            while d_1751_v62_.f27:
                with _dafny.c_label("11"):
                    if (d_1752_i5_) >= (100):
                        raise _dafny.Break("11")
                    d_1752_i5_ = (d_1752_i5_) + (1)
                    (globalState).f7 = p0
                    d_1753_v63_: _dafny.MultiSet
                    d_1753_v63_ = _dafny.MultiSet([d_1751_v62_.f27])
                    d_1753_v63_ = default__.fm52(d_1751_v62_.f27, p0, globalState)
                    d_1754_v64_: _dafny.Array
                    def lambda95_(d_1755_v44_):
                        def lambda96_(d_1756_i6_):
                            return _dafny.Set({d_1755_v44_, d_1755_v44_})

                        return lambda96_

                    init52_ = lambda95_(d_1730_v44_)
                    nw291_ = _dafny.Array(None, 29)
                    for i0_52_ in range(nw291_.length(0)):
                        nw291_[i0_52_] = init52_(i0_52_)
                    d_1754_v64_ = nw291_
                    d_1757_v65_: _dafny.Set
                    d_1757_v65_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ycoxy"))})
                    index299_ = default__.safeIndex(257, (d_1754_v64_).length(0))
                    def iife125_():
                        coll73_ = _dafny.Set()
                        compr_73_: _dafny.Seq
                        for compr_73_ in (d_1757_v65_).Elements:
                            d_1758_v66_: _dafny.Seq = compr_73_
                            if (d_1758_v66_) in (d_1757_v65_):
                                coll73_ = coll73_.union(_dafny.Set([d_1758_v66_]))
                        return _dafny.Set(coll73_)
                    (d_1754_v64_)[index299_] = (d_1757_v65_).intersection(iife125_()
                    )
                    index300_ = default__.safeIndex(257, (d_1754_v64_).length(0))
                    (d_1754_v64_)[index300_] = d_1757_v65_
                    d_1759_v67_: str
                    d_1759_v67_ = _dafny.CodePoint('b')
                    d_1759_v67_ = d_1759_v67_
                    pass
            pass
        d_1760_v68_: _dafny.Array
        nw292_ = _dafny.Array(None, 1)
        nw292_[int(0)] = (0) - (p0)
        d_1760_v68_ = nw292_
        r0 = d_1760_v68_
        r1 = (d_1751_v62_.f27 if d_1683_v0_ else d_1751_v62_.f27)
        return r0, r1


class C7(T1, T0):
    def  __init__(self):
        self._f23: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    def ctor__(self, f23):
        (self)._f23 = f23

    def fm7(self, p0, p1, p2, globalState):
        return False

    def fm5(self, p0, p1, p2, p3, globalState):
        return _dafny.Map({(not(True) if False else True): (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uqyvs"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "egwtgbjvl")))})

    def fm6(self, p0, globalState):
        return not(not (not(True)) or ((_dafny.Set({-747})).isdisjoint(_dafny.Set({158, -959}))))

    def m2(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: int = int(0)
        r2: _dafny.Array = _dafny.Array(None, 0)
        d_1761_v0_: C6
        nw293_ = C6()
        nw293_.ctor__()
        d_1761_v0_ = nw293_
        hi14_ = p3
        for d_1762_i0_ in range(p3, hi14_):
            (globalState).f16 = (914) - ((d_1762_i0_) + (len(_dafny.SeqWithoutIsStrInference([p3]))))
            d_1763_v1_: _dafny.Array
            def lambda97_(d_1764_i1_):
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rhc"))

            init53_ = lambda97_
            nw294_ = _dafny.Array(None, 5)
            for i0_53_ in range(nw294_.length(0)):
                nw294_[i0_53_] = init53_(i0_53_)
            d_1763_v1_ = nw294_
            d_1763_v1_ = (d_1763_v1_ if (d_1762_i0_) < (d_1762_i0_) else d_1763_v1_)
            d_1765_v2_: T1
            nw295_ = C5()
            nw295_.ctor__(p1)
            d_1765_v2_ = nw295_
            d_1765_v2_ = d_1765_v2_
            d_1766_v4_: _dafny.Array
            def lambda98_(d_1767_p1_):
                def lambda99_(d_1768_i2_):
                    def iife126_():
                        coll74_ = _dafny.Map()
                        compr_74_: _dafny.Set
                        for compr_74_ in (_dafny.Set({_dafny.Set({d_1767_p1_}), _dafny.Set({d_1767_p1_})})).Elements:
                            d_1769_v3_: _dafny.Set = compr_74_
                            if (d_1769_v3_) in (_dafny.Set({_dafny.Set({d_1767_p1_}), _dafny.Set({d_1767_p1_})})):
                                coll74_[d_1769_v3_] = d_1767_p1_
                        return _dafny.Map(coll74_)
                    return _dafny.SeqWithoutIsStrInference([iife126_()
                    , _dafny.Map({_dafny.Set({d_1767_p1_}): d_1767_p1_})])

                return lambda99_

            init54_ = lambda98_(p1)
            nw296_ = _dafny.Array(None, 27)
            for i0_54_ in range(nw296_.length(0)):
                nw296_[i0_54_] = init54_(i0_54_)
            d_1766_v4_ = nw296_
            d_1770_v5_: _dafny.Set
            d_1770_v5_ = _dafny.Set({False, p1})
            d_1771_v6_: _dafny.Map
            d_1771_v6_ = _dafny.Map({d_1770_v5_: p1})
            d_1772_v7_: _dafny.Seq
            d_1772_v7_ = _dafny.SeqWithoutIsStrInference([d_1771_v6_])
            d_1773_v8_: _dafny.Set
            d_1773_v8_ = d_1770_v5_
            index301_ = default__.safeIndex(54, (d_1766_v4_).length(0))
            (d_1766_v4_)[index301_] = (d_1772_v7_) + (_dafny.SeqWithoutIsStrInference([(_dafny.Map({d_1773_v8_: p1})).set(d_1773_v8_, False)]))
            index302_ = default__.safeIndex(54, (d_1766_v4_).length(0))
            (d_1766_v4_)[index302_] = d_1772_v7_
        d_1774_v9_: C2
        nw297_ = C2()
        nw297_.ctor__((0) - (p0))
        d_1774_v9_ = nw297_
        d_1775_v10_: _dafny.Set
        d_1775_v10_ = _dafny.Set({d_1774_v9_, d_1774_v9_, d_1774_v9_})
        if (d_1775_v10_).issubset(_dafny.Set({d_1774_v9_, d_1774_v9_})):
            (globalState).f16 = p3
            d_1776_v11_: _dafny.Seq
            d_1776_v11_ = _dafny.SeqWithoutIsStrInference([(d_1774_v9_).f28])
            d_1777_v12_: C0
            nw298_ = C0()
            nw298_.ctor__(((d_1774_v9_).f28) != (default__.fm0(p0, len(d_1776_v11_), default__.fm0(p3, p3, p0, p1, globalState), p1, globalState)))
            d_1777_v12_ = nw298_
            d_1778_v13_: _dafny.Array
            def lambda100_(d_1779_v12_, d_1780_p1_):
                def lambda101_(d_1781_i3_):
                    return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_1779_v12_.f27]))).isdisjoint(_dafny.MultiSet([d_1780_p1_, not(d_1779_v12_.f27)]))

                return lambda101_

            init55_ = lambda100_(d_1777_v12_, p1)
            nw299_ = _dafny.Array(None, 20)
            for i0_55_ in range(nw299_.length(0)):
                nw299_[i0_55_] = init55_(i0_55_)
            d_1778_v13_ = nw299_
            d_1782_v14_: C6
            nw300_ = C6()
            nw300_.ctor__()
            d_1782_v14_ = nw300_
            d_1783_v15_: _dafny.Seq
            d_1783_v15_ = _dafny.SeqWithoutIsStrInference([d_1777_v12_.f27])
            d_1784_v16_: _dafny.MultiSet
            d_1784_v16_ = _dafny.MultiSet([(d_1783_v15_ if d_1777_v12_.f27 else d_1783_v15_)])
            d_1785_v17_: _dafny.Seq
            d_1785_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mdrtqm"))
            rhs301_ = (0) - ((0) - ((d_1774_v9_).f28))
            rhs302_ = len(d_1785_v17_)
            rhs303_ = d_1778_v13_
            rhs304_ = d_1782_v14_
            rhs305_ = d_1784_v16_
            lhs285_ = globalState
            r1 = rhs301_
            lhs285_.f7 = rhs302_
            d_1778_v13_ = rhs303_
            d_1782_v14_ = rhs304_
            d_1784_v16_ = rhs305_
            d_1786_v18_: _dafny.Map
            d_1786_v18_ = _dafny.Map({p0: p1})
            d_1787_v19_: C1
            nw301_ = C1()
            nw301_.ctor__(((d_1774_v9_).f28) + (default__.fm0(p3, (d_1774_v9_).f28, (d_1774_v9_).f28, True, globalState)), _dafny.SeqWithoutIsStrInference([d_1786_v18_, d_1786_v18_, d_1786_v18_]))
            d_1787_v19_ = nw301_
            (globalState).f7 = (0) - ((d_1774_v9_).fm17(False, d_1776_v11_, p3, d_1783_v15_, globalState))
        elif True:
            d_1788_v20_: D17
            d_1788_v20_ = D17_DC40()
            source23_ = d_1788_v20_
            if source23_.is_DC39:
                d_1789___mcc_h0_ = source23_.cf52
                d_1790___mcc_h1_ = source23_.cf53
                d_1791___mcc_h2_ = source23_.cf54
                d_1792___mcc_h3_ = source23_.cf55
                d_1793___mcc_h4_ = source23_.cf56
                d_1794_cf56_ = d_1793___mcc_h4_
                d_1795_cf55_ = d_1792___mcc_h3_
                d_1796_cf54_ = d_1791___mcc_h2_
                d_1797_cf53_ = d_1790___mcc_h1_
                d_1798_cf52_ = d_1789___mcc_h0_
                d_1798_cf52_ = -36
                d_1799_v21_: _dafny.Map
                d_1799_v21_ = _dafny.Map({p1: False})
                d_1800_v22_: _dafny.Seq
                d_1800_v22_ = _dafny.SeqWithoutIsStrInference([d_1799_v21_, d_1799_v21_])
                (globalState).f13 = default__.fm2((d_1800_v22_)[default__.safeIndex(d_1798_cf52_, len(d_1800_v22_))], globalState)
                d_1801_v23_: _dafny.MultiSet
                d_1801_v23_ = _dafny.MultiSet([not(p1), p1])
                d_1802_v24_: C2
                nw302_ = C2()
                nw302_.ctor__((d_1801_v23_).cardinality)
                d_1802_v24_ = nw302_
                d_1803_v25_: _dafny.Seq
                d_1803_v25_ = _dafny.SeqWithoutIsStrInference([d_1798_cf52_, 955, (d_1774_v9_).f28, (d_1801_v23_).cardinality])
                (globalState).f0 = (d_1803_v25_) == (d_1803_v25_)
            elif source23_.is_DC40:
                d_1804_v26_: _dafny.Seq
                d_1804_v26_ = _dafny.SeqWithoutIsStrInference([p1, p1])
                d_1805_v27_: _dafny.Seq
                d_1805_v27_ = d_1804_v26_
                index303_ = default__.safeIndex(688, (p2).length(0))
                (p2)[index303_] = (d_1774_v9_).fm17(p1, _dafny.SeqWithoutIsStrInference([(d_1774_v9_).f28]), (d_1774_v9_).f28, d_1805_v27_, globalState)
                d_1806_v28_: _dafny.Seq
                d_1806_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ug"))
                index304_ = default__.safeIndex(688, (p2).length(0))
                (p2)[index304_] = len(d_1806_v28_)
                (globalState).f16 = p3
                d_1807_v29_: C5
                nw303_ = C5()
                nw303_.ctor__((self).fm7(p1, (0) - ((p2)[default__.safeIndex(688, (p2).length(0))]), p1, globalState))
                d_1807_v29_ = nw303_
                d_1808_v30_: str
                d_1808_v30_ = _dafny.CodePoint('u')
                d_1809_v31_: _dafny.Array
                nw304_ = _dafny.Array(_dafny.Map({}), 21)
                d_1809_v31_ = nw304_
                d_1810_v32_: _dafny.Map
                d_1810_v32_ = _dafny.Map({p1: (d_1807_v29_).f24})
                d_1811_v33_: _dafny.Map
                d_1811_v33_ = _dafny.Map({(d_1807_v29_).f24: d_1810_v32_})
                index305_ = default__.safeIndex(7, (d_1809_v31_).length(0))
                (d_1809_v31_)[index305_] = (D23_DC60(d_1811_v33_)).cf85
                d_1812_v34_: _dafny.Set
                d_1812_v34_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uvqcohcw"))})
                d_1813_v35_: D0
                d_1813_v35_ = D0_DC0(p1)
                d_1814_v36_: _dafny.Set
                d_1814_v36_ = _dafny.Set({d_1808_v30_})
                d_1815_v37_: _dafny.MultiSet
                d_1815_v37_ = _dafny.MultiSet([648])
                d_1816_v38_: _dafny.Map
                d_1816_v38_ = _dafny.Map({(d_1774_v9_).f28: p3})
                d_1817_v39_: _dafny.Map
                d_1817_v39_ = _dafny.Map({d_1816_v38_: p1})
                d_1818_v40_: _dafny.Map
                d_1818_v40_ = _dafny.Map({(d_1807_v29_).f24: len(default__.fm54(globalState))})
                d_1819_v41_: _dafny.MultiSet
                d_1819_v41_ = _dafny.MultiSet([True])
                d_1820_v42_: _dafny.Seq
                d_1820_v42_ = _dafny.SeqWithoutIsStrInference([(d_1819_v41_).cardinality])
                d_1821_v43_: D15
                d_1821_v43_ = D15_DC33(d_1820_v42_)
                d_1822_v44_: _dafny.Array
                nw305_ = _dafny.Array(None, 29)
                nw305_[int(0)] = (d_1807_v29_).f24
                nw305_[int(1)] = (d_1812_v34_).isdisjoint(default__.fm53(p1, globalState))
                nw305_[int(2)] = not((d_1807_v29_).f24)
                nw305_[int(3)] = p1
                nw305_[int(4)] = (d_1807_v29_).f24
                nw305_[int(5)] = p1
                nw305_[int(6)] = not(p1)
                nw305_[int(7)] = False
                nw305_[int(8)] = (d_1807_v29_).f24
                nw305_[int(9)] = p1
                nw305_[int(10)] = not((d_1813_v35_).cf0)
                nw305_[int(11)] = p1
                nw305_[int(12)] = (d_1814_v36_).issubset(d_1814_v36_)
                nw305_[int(13)] = (d_1807_v29_).f24
                nw305_[int(14)] = (d_1807_v29_).fm6(len(d_1804_v26_), globalState)
                nw305_[int(15)] = not((d_1815_v37_).isdisjoint(d_1815_v37_))
                nw305_[int(16)] = (d_1807_v29_).f24
                nw305_[int(17)] = False
                nw305_[int(18)] = (d_1807_v29_).f24
                nw305_[int(19)] = p1
                nw305_[int(20)] = (d_1807_v29_).f24
                nw305_[int(21)] = (d_1807_v29_).f24
                nw305_[int(22)] = (d_1816_v38_) in (d_1817_v39_)
                nw305_[int(23)] = p1
                nw305_[int(24)] = not(p1)
                nw305_[int(25)] = (d_1807_v29_).f24
                nw305_[int(26)] = (len(d_1818_v40_)) != ((d_1774_v9_).f28)
                nw305_[int(27)] = (self).fm6((d_1774_v9_).fm17(not(p1), (d_1821_v43_).cf44, (d_1774_v9_).f28, d_1805_v27_, globalState), globalState)
                nw305_[int(28)] = True
                d_1822_v44_ = nw305_
                index306_ = default__.safeIndex(329, (d_1822_v44_).length(0))
                (d_1822_v44_)[index306_] = (d_1807_v29_).f24
                d_1823_v45_: _dafny.Seq
                d_1823_v45_ = _dafny.SeqWithoutIsStrInference([d_1816_v38_])
                index307_ = default__.safeIndex(7, (d_1809_v31_).length(0))
                index308_ = default__.safeIndex(329, (d_1822_v44_).length(0))
                rhs306_ = (d_1808_v30_ if (d_1807_v29_).f24 else d_1808_v30_)
                rhs307_ = (d_1811_v33_) | ((d_1811_v33_ if p1 else d_1811_v33_))
                rhs308_ = not(not((d_1807_v29_).f24))
                rhs309_ = default__.safeDivisionInt((d_1820_v42_)[default__.safeIndex((_dafny.MultiSet(d_1823_v45_)).cardinality, len(d_1820_v42_))], default__.fm0((d_1820_v42_)[default__.safeIndex(len(d_1806_v28_), len(d_1820_v42_))], 431, p0, default__.fm2(d_1810_v32_, globalState), globalState))
                rhs310_ = (0) - ((p2)[default__.safeIndex(688, (p2).length(0))])
                lhs286_ = d_1809_v31_
                lhs287_ = default__.safeIndex(7, (d_1809_v31_).length(0))
                lhs288_ = d_1822_v44_
                lhs289_ = default__.safeIndex(329, (d_1822_v44_).length(0))
                lhs290_ = globalState
                d_1808_v30_ = rhs306_
                lhs286_[lhs287_] = rhs307_
                lhs288_[lhs289_] = rhs308_
                r1 = rhs309_
                lhs290_.f7 = rhs310_
            elif source23_.is_DC41:
                d_1824___mcc_h5_ = source23_.cf57
                d_1825_cf57_ = d_1824___mcc_h5_
                (globalState).f7 = 577
                d_1826_v46_: _dafny.Map
                d_1826_v46_ = _dafny.Map({-300: p1})
                d_1827_v47_: _dafny.Seq
                d_1827_v47_ = _dafny.SeqWithoutIsStrInference([d_1826_v46_])
                d_1828_v48_: C1
                nw306_ = C1()
                nw306_.ctor__(p0, d_1827_v47_)
                d_1828_v48_ = nw306_
                (globalState).f21 = p1
                d_1829_v49_: D19
                d_1829_v49_ = D19_DC48(p0, p1)
                r1 = (d_1829_v49_).cf65
            elif True:
                d_1830___mcc_h6_ = source23_.cf51
                d_1831_cf51_ = d_1830___mcc_h6_
                d_1832_v50_: _dafny.Seq
                d_1832_v50_ = _dafny.SeqWithoutIsStrInference([p1, p1, p1])
                d_1833_v51_: _dafny.MultiSet
                d_1833_v51_ = _dafny.MultiSet([p1])
                (globalState).f2 = ((d_1832_v50_)[default__.safeIndex((d_1833_v51_).cardinality, len(d_1832_v50_))]) == (p1)
                d_1834_v52_: _dafny.MultiSet
                d_1834_v52_ = _dafny.MultiSet([default__.fm0(p3, p0, (d_1774_v9_).f28, p1, globalState), (d_1774_v9_).f28, (d_1774_v9_).f28, (d_1774_v9_).f28, (d_1774_v9_).f28])
                (globalState).f2 = (d_1834_v52_).isdisjoint(_dafny.MultiSet([(d_1774_v9_).f28]))
                d_1835_v53_: str
                d_1835_v53_ = _dafny.CodePoint('n')
                d_1835_v53_ = default__.fm34(globalState)
                d_1836_v54_: _dafny.Array
                nw307_ = _dafny.Array(False, 3)
                d_1836_v54_ = nw307_
                index309_ = default__.safeIndex(65, (d_1836_v54_).length(0))
                (d_1836_v54_)[index309_] = not (p1) or (p1)
                d_1837_v55_: _dafny.Set
                d_1837_v55_ = _dafny.Set({True})
                index310_ = default__.safeIndex(65, (d_1836_v54_).length(0))
                (d_1836_v54_)[index310_] = ((d_1837_v55_).issubset(d_1837_v55_) if not(p1) else not(p1))
            d_1838_v56_: str
            d_1838_v56_ = _dafny.CodePoint('g')
            d_1839_v57_: _dafny.Array
            nw308_ = _dafny.Array(None, 13)
            nw308_[int(0)] = d_1838_v56_
            nw308_[int(1)] = d_1838_v56_
            nw308_[int(2)] = d_1838_v56_
            nw308_[int(3)] = d_1838_v56_
            nw308_[int(4)] = d_1838_v56_
            nw308_[int(5)] = d_1838_v56_
            nw308_[int(6)] = _dafny.CodePoint('k')
            nw308_[int(7)] = _dafny.CodePoint('f')
            nw308_[int(8)] = d_1838_v56_
            nw308_[int(9)] = d_1838_v56_
            nw308_[int(10)] = d_1838_v56_
            nw308_[int(11)] = d_1838_v56_
            nw308_[int(12)] = d_1838_v56_
            d_1839_v57_ = nw308_
            d_1840_v58_: _dafny.Map
            d_1840_v58_ = _dafny.Map({p3: d_1839_v57_})
            d_1839_v57_ = ((d_1840_v58_)[(d_1774_v9_).f28] if ((d_1774_v9_).f28) in (d_1840_v58_) else d_1839_v57_)
            d_1841_v59_: _dafny.Seq
            d_1841_v59_ = _dafny.SeqWithoutIsStrInference([p1])
            d_1842_v60_: _dafny.Seq
            d_1842_v60_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sbxhia"))
            d_1843_v61_: _dafny.MultiSet
            d_1843_v61_ = _dafny.MultiSet([len(d_1842_v60_), p3])
            d_1844_v62_: _dafny.Map
            d_1844_v62_ = _dafny.Map({(len(d_1841_v59_)) * ((d_1843_v61_).cardinality): d_1838_v56_})
            rhs311_ = ((d_1844_v62_)[default__.safeModuloInt((d_1774_v9_).f28, len(_dafny.SeqWithoutIsStrInference([len(d_1842_v60_), (d_1774_v9_).f28])))] if (default__.safeModuloInt((d_1774_v9_).f28, len(_dafny.SeqWithoutIsStrInference([len(d_1842_v60_), (d_1774_v9_).f28])))) in (d_1844_v62_) else d_1838_v56_)
            rhs312_ = (d_1774_v9_).fm6((d_1774_v9_).f28, globalState)
            lhs291_ = globalState
            d_1838_v56_ = rhs311_
            lhs291_.f2 = rhs312_
            index311_ = default__.safeIndex(357, (p2).length(0))
            (p2)[index311_] = p3
            index312_ = default__.safeIndex(357, (p2).length(0))
            (p2)[index312_] = len(d_1842_v60_)
            r1 = 794
        d_1845_v63_: C2
        nw309_ = C2()
        nw309_.ctor__(((d_1774_v9_).f28) + (p0))
        d_1845_v63_ = nw309_
        d_1846_v64_: _dafny.Map
        d_1846_v64_ = _dafny.Map({False: (d_1774_v9_).f28})
        d_1847_v65_: _dafny.Set
        d_1847_v65_ = _dafny.Set({_dafny.Map({p1: p0}), d_1846_v64_})
        d_1848_v66_: _dafny.Map
        d_1848_v66_ = _dafny.Map({p1: (d_1847_v65_).issubset(d_1847_v65_)})
        d_1849_i4_: int
        d_1849_i4_ = 0
        with _dafny.label("12"):
            while not(((d_1848_v66_)[p1] if (p1) in (d_1848_v66_) else True)):
                with _dafny.c_label("12"):
                    if (d_1849_i4_) >= (100):
                        raise _dafny.Break("12")
                    d_1849_i4_ = (d_1849_i4_) + (1)
                    d_1850_v67_: C3
                    nw310_ = C3()
                    nw310_.ctor__()
                    d_1850_v67_ = nw310_
                    (globalState).f16 = len(default__.fm1((d_1774_v9_).f28, globalState))
                    index313_ = default__.safeIndex(985, (p2).length(0))
                    (p2)[index313_] = p0
                    index314_ = default__.safeIndex(985, (p2).length(0))
                    (p2)[index314_] = p0
                    d_1851_v68_: _dafny.Array
                    nw311_ = _dafny.Array(False, 21)
                    d_1851_v68_ = nw311_
                    index315_ = default__.safeIndex(629, (d_1851_v68_).length(0))
                    (d_1851_v68_)[index315_] = p1
                    index316_ = default__.safeIndex(629, (d_1851_v68_).length(0))
                    (d_1851_v68_)[index316_] = p1
                    pass
            pass
        d_1852_v69_: _dafny.Seq
        d_1852_v69_ = _dafny.SeqWithoutIsStrInference([p1, True, (p1) and (p1)])
        d_1853_v70_: _dafny.Seq
        d_1853_v70_ = _dafny.SeqWithoutIsStrInference([(d_1845_v63_).f28, 773])
        d_1854_v71_: _dafny.Seq
        d_1854_v71_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ygs"))
        d_1855_i5_: int
        d_1855_i5_ = 0
        with _dafny.label("13"):
            while (d_1852_v69_)[default__.safeIndex(default__.safeDivisionInt((0) - ((d_1774_v9_).f28), (d_1853_v70_)[default__.safeIndex(len(d_1854_v71_), len(d_1853_v70_))]), len(d_1852_v69_))]:
                with _dafny.c_label("13"):
                    if (d_1855_i5_) >= (100):
                        raise _dafny.Break("13")
                    d_1855_i5_ = (d_1855_i5_) + (1)
                    d_1856_v72_: _dafny.Seq
                    d_1856_v72_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qxmym")), d_1854_v71_, d_1854_v71_, default__.fm1(len(d_1854_v71_), globalState), d_1854_v71_])
                    d_1857_v73_: _dafny.MultiSet
                    d_1857_v73_ = _dafny.MultiSet([(d_1856_v72_)[default__.safeIndex((d_1774_v9_).f28, len(d_1856_v72_))], _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_1858_i6_ in range(default__.abs(991))]), d_1854_v71_])
                    d_1859_v74_: _dafny.Map
                    d_1859_v74_ = _dafny.Map({(d_1774_v9_).f28: default__.fm55((d_1845_v63_).f28, p1, p1, globalState)})
                    d_1860_v75_: _dafny.Map
                    d_1860_v75_ = _dafny.Map({(d_1852_v69_)[default__.safeIndex(p0, len(d_1852_v69_))]: ((d_1859_v74_)[p3] if (p3) in (d_1859_v74_) else d_1857_v73_)})
                    d_1861_v76_: _dafny.Seq
                    d_1861_v76_ = _dafny.SeqWithoutIsStrInference([(d_1857_v73_).set(d_1854_v71_, default__.abs((d_1845_v63_).f28)), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_1854_v71_]))])
                    d_1862_v77_: _dafny.Array
                    nw312_ = _dafny.Array(None, 17)
                    nw312_[int(0)] = _dafny.MultiSet([d_1854_v71_, d_1854_v71_])
                    nw312_[int(1)] = d_1857_v73_
                    nw312_[int(2)] = d_1857_v73_
                    nw312_[int(3)] = (d_1857_v73_ if p1 else _dafny.MultiSet(d_1856_v72_))
                    nw312_[int(4)] = (d_1857_v73_).intersection(d_1857_v73_)
                    nw312_[int(5)] = d_1857_v73_
                    nw312_[int(6)] = _dafny.MultiSet([d_1854_v71_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "biloj"))])
                    nw312_[int(7)] = d_1857_v73_
                    nw312_[int(8)] = d_1857_v73_
                    nw312_[int(9)] = d_1857_v73_
                    nw312_[int(10)] = d_1857_v73_
                    nw312_[int(11)] = ((d_1860_v75_)[p1] if (p1) in (d_1860_v75_) else d_1857_v73_)
                    nw312_[int(12)] = (d_1861_v76_)[default__.safeIndex(257, len(d_1861_v76_))]
                    nw312_[int(13)] = d_1857_v73_
                    nw312_[int(14)] = _dafny.MultiSet([d_1854_v71_])
                    nw312_[int(15)] = (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_1863_i7_ in range(default__.abs(-724))])])).set(d_1854_v71_, default__.abs((d_1845_v63_).f28))
                    nw312_[int(16)] = d_1857_v73_
                    d_1862_v77_ = nw312_
                    index317_ = default__.safeIndex(233, (d_1862_v77_).length(0))
                    (d_1862_v77_)[index317_] = _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([d_1854_v71_])) + (d_1856_v72_))
                    index318_ = default__.safeIndex(233, (d_1862_v77_).length(0))
                    (d_1862_v77_)[index318_] = d_1857_v73_
                    (globalState).f16 = ((0) - ((d_1845_v63_).f28)) + (((d_1845_v63_).f28) + ((d_1845_v63_).f28))
                    (globalState).f13 = not(p1)
                    d_1864_v78_: T0
                    nw313_ = C3()
                    nw313_.ctor__()
                    d_1864_v78_ = nw313_
                    index319_ = default__.safeIndex(344, (p2).length(0))
                    (p2)[index319_] = (d_1845_v63_).f28
                    index320_ = default__.safeIndex(344, (p2).length(0))
                    rhs313_ = d_1864_v78_
                    rhs314_ = (_dafny.SeqWithoutIsStrInference([(d_1774_v9_).f28, p0])) < (_dafny.SeqWithoutIsStrInference([p3]))
                    rhs315_ = p3
                    rhs316_ = p1
                    rhs317_ = (d_1774_v9_).f28
                    lhs292_ = globalState
                    lhs293_ = p2
                    lhs294_ = default__.safeIndex(344, (p2).length(0))
                    lhs295_ = globalState
                    lhs296_ = globalState
                    d_1864_v78_ = rhs313_
                    lhs292_.f2 = rhs314_
                    lhs293_[lhs294_] = rhs315_
                    lhs295_.f13 = rhs316_
                    lhs296_.f7 = rhs317_
                    pass
            pass
        d_1865_v79_: C0
        nw314_ = C0()
        nw314_.ctor__(p1)
        d_1865_v79_ = nw314_
        d_1866_v80_: _dafny.Seq
        d_1866_v80_ = _dafny.SeqWithoutIsStrInference([d_1865_v79_])
        d_1867_v81_: _dafny.Map
        d_1867_v81_ = _dafny.Map({(d_1866_v80_)[default__.safeIndex(len(d_1854_v71_), len(d_1866_v80_))]: (d_1774_v9_).f28})
        r0 = (d_1852_v69_)[default__.safeIndex(len(d_1867_v81_), len(d_1852_v69_))]
        r1 = ((0) - (len(_dafny.Map({p1: (d_1845_v63_).f28})))) + ((d_1845_v63_).f28)
        def lambda102_(d_1868_v79_):
            def lambda103_(d_1869_i8_):
                return _dafny.Map({d_1868_v79_.f27: D3_DC6(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "twmsayjd")), _dafny.CodePoint('c'), d_1868_v79_.f27)})

            return lambda103_

        init56_ = lambda102_(d_1865_v79_)
        nw315_ = _dafny.Array(None, 10)
        for i0_56_ in range(nw315_.length(0)):
            nw315_[i0_56_] = init56_(i0_56_)
        r2 = nw315_
        return r0, r1, r2

    def m3(self, p0, globalState):
        d_1870_v0_: bool
        d_1870_v0_ = True
        d_1871_i0_: int
        d_1871_i0_ = 0
        with _dafny.label("14"):
            while not(d_1870_v0_):
                with _dafny.c_label("14"):
                    if (d_1871_i0_) >= (100):
                        raise _dafny.Break("14")
                    d_1871_i0_ = (d_1871_i0_) + (1)
                    if d_1870_v0_:
                        d_1872_v1_: _dafny.Array
                        def lambda104_(d_1873_v0_):
                            def lambda105_(d_1874_i1_):
                                return d_1873_v0_

                            return lambda105_

                        init57_ = lambda104_(d_1870_v0_)
                        nw316_ = _dafny.Array(None, 6)
                        for i0_57_ in range(nw316_.length(0)):
                            nw316_[i0_57_] = init57_(i0_57_)
                        d_1872_v1_ = nw316_
                        d_1875_v2_: _dafny.Seq
                        d_1875_v2_ = _dafny.SeqWithoutIsStrInference([d_1872_v1_])
                        d_1876_v3_: int
                        d_1876_v3_ = 533
                        d_1872_v1_ = (d_1875_v2_)[default__.safeIndex(d_1876_v3_, len(d_1875_v2_))]
                        d_1877_v4_: _dafny.Array
                        def lambda106_(d_1878_i2_):
                            return D15_DC33(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bauwltxil")))]))

                        init58_ = lambda106_
                        nw317_ = _dafny.Array(None, 20)
                        for i0_58_ in range(nw317_.length(0)):
                            nw317_[i0_58_] = init58_(i0_58_)
                        d_1877_v4_ = nw317_
                        index321_ = default__.safeIndex(559, (d_1877_v4_).length(0))
                        (d_1877_v4_)[index321_] = D15_DC33(default__.fm21(globalState))
                        d_1879_v5_: _dafny.Seq
                        d_1879_v5_ = _dafny.SeqWithoutIsStrInference([d_1870_v0_])
                        d_1880_v6_: _dafny.Map
                        d_1880_v6_ = _dafny.Map({D14_DC29(d_1876_v3_, _dafny.Map({False: d_1879_v5_})): -262})
                        d_1881_v7_: _dafny.Map
                        d_1881_v7_ = _dafny.Map({d_1870_v0_: p0})
                        d_1882_v8_: _dafny.Seq
                        d_1882_v8_ = _dafny.SeqWithoutIsStrInference([d_1876_v3_, len(d_1880_v6_), len(d_1881_v7_)])
                        d_1883_v9_: D15
                        d_1883_v9_ = D15_DC33(d_1882_v8_)
                        pat_let_tv32_ = d_1882_v8_
                        index322_ = default__.safeIndex(559, (d_1877_v4_).length(0))
                        def iife127_(_pat_let26_0):
                            def iife128_(d_1884_dt__update__tmp_h0_):
                                def iife129_(_pat_let27_0):
                                    def iife130_(d_1885_dt__update_hcf44_h0_):
                                        return D15_DC33(d_1885_dt__update_hcf44_h0_)
                                    return iife130_(_pat_let27_0)
                                return iife129_(pat_let_tv32_)
                            return iife128_(_pat_let26_0)
                        rhs318_ = d_1870_v0_
                        rhs319_ = (d_1876_v3_) < (d_1876_v3_)
                        rhs320_ = iife127_(d_1883_v9_)
                        rhs321_ = (0) - (default__.safeDivisionInt(d_1876_v3_, default__.safeDivisionInt(d_1876_v3_, d_1876_v3_)))
                        lhs297_ = globalState
                        lhs298_ = globalState
                        lhs299_ = d_1877_v4_
                        lhs300_ = default__.safeIndex(559, (d_1877_v4_).length(0))
                        lhs301_ = globalState
                        lhs297_.f13 = rhs318_
                        lhs298_.f21 = rhs319_
                        lhs299_[lhs300_] = rhs320_
                        lhs301_.f16 = rhs321_
                        d_1886_v10_: _dafny.Set
                        d_1886_v10_ = _dafny.Set({d_1876_v3_})
                        d_1887_v11_: C5
                        nw318_ = C5()
                        nw318_.ctor__((d_1886_v10_).ispropersubset(_dafny.Set({d_1876_v3_})))
                        d_1887_v11_ = nw318_
                        d_1870_v0_ = (d_1887_v11_).f24
                        index323_ = default__.safeIndex(37, (d_1872_v1_).length(0))
                        (d_1872_v1_)[index323_] = (d_1879_v5_)[default__.safeIndex(d_1876_v3_, len(d_1879_v5_))]
                        d_1888_v12_: _dafny.Set
                        d_1888_v12_ = _dafny.Set({d_1882_v8_})
                        d_1889_v13_: _dafny.Seq
                        d_1889_v13_ = _dafny.SeqWithoutIsStrInference([d_1882_v8_, d_1882_v8_, d_1882_v8_])
                        d_1890_v15_: _dafny.MultiSet
                        d_1890_v15_ = _dafny.MultiSet([d_1870_v0_, d_1870_v0_, (d_1887_v11_).f24])
                        index324_ = default__.safeIndex(37, (d_1872_v1_).length(0))
                        def iife131_():
                            coll75_ = _dafny.Set()
                            compr_75_: _dafny.Seq
                            for compr_75_ in (d_1889_v13_).Elements:
                                d_1891_v14_: _dafny.Seq = compr_75_
                                if (d_1891_v14_) in (d_1889_v13_):
                                    coll75_ = coll75_.union(_dafny.Set([d_1891_v14_]))
                            return _dafny.Set(coll75_)
                        rhs322_ = (d_1887_v11_).f24
                        rhs323_ = ((d_1888_v12_) - (iife131_()
                        )).ispropersubset(_dafny.Set({(d_1882_v8_).set(default__.safeIndex(d_1876_v3_, len(d_1882_v8_)), d_1876_v3_)}))
                        rhs324_ = (d_1890_v15_).isdisjoint(d_1890_v15_)
                        lhs302_ = globalState
                        lhs303_ = d_1872_v1_
                        lhs304_ = default__.safeIndex(37, (d_1872_v1_).length(0))
                        lhs305_ = globalState
                        lhs302_.f2 = rhs322_
                        lhs303_[lhs304_] = rhs323_
                        lhs305_.f13 = rhs324_
                    elif True:
                        d_1892_v16_: C4
                        nw319_ = C4()
                        nw319_.ctor__(d_1870_v0_, d_1870_v0_)
                        d_1892_v16_ = nw319_
                        d_1893_v17_: _dafny.Array
                        nw320_ = _dafny.Array(_dafny.Set({}), 16)
                        d_1893_v17_ = nw320_
                        d_1894_v18_: _dafny.Set
                        d_1894_v18_ = _dafny.Set({738})
                        index325_ = default__.safeIndex(679, (d_1893_v17_).length(0))
                        (d_1893_v17_)[index325_] = d_1894_v18_
                        index326_ = default__.safeIndex(679, (d_1893_v17_).length(0))
                        (d_1893_v17_)[index326_] = d_1894_v18_
                        d_1895_v19_: _dafny.Array
                        nw321_ = _dafny.Array(_dafny.CodePoint('D'), 3)
                        d_1895_v19_ = nw321_
                        d_1896_v20_: str
                        d_1896_v20_ = _dafny.CodePoint('f')
                        index327_ = default__.safeIndex(577, (d_1895_v19_).length(0))
                        (d_1895_v19_)[index327_] = d_1896_v20_
                        index328_ = default__.safeIndex(577, (d_1895_v19_).length(0))
                        (d_1895_v19_)[index328_] = d_1896_v20_
                        d_1897_v21_: int
                        d_1897_v21_ = 486
                        d_1898_v22_: _dafny.Map
                        d_1898_v22_ = _dafny.Map({d_1897_v21_: (d_1895_v19_)[default__.safeIndex(577, (d_1895_v19_).length(0))]})
                        d_1899_v23_: _dafny.Seq
                        d_1899_v23_ = _dafny.SeqWithoutIsStrInference([d_1870_v0_])
                        d_1900_v24_: _dafny.Seq
                        d_1900_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nwdg"))
                        d_1901_v25_: _dafny.Seq
                        d_1901_v25_ = _dafny.SeqWithoutIsStrInference([d_1900_v24_, d_1900_v24_])
                        d_1902_v26_: _dafny.Map
                        d_1902_v26_ = _dafny.Map({(d_1901_v25_)[default__.safeIndex(d_1897_v21_, len(d_1901_v25_))]: (d_1892_v16_).f26})
                        d_1903_v27_: C0
                        nw322_ = C0()
                        nw322_.ctor__(d_1892_v16_.f25)
                        d_1903_v27_ = nw322_
                        d_1904_v28_: _dafny.Map
                        d_1904_v28_ = _dafny.Map({d_1903_v27_: d_1870_v0_})
                        d_1905_v29_: _dafny.Map
                        d_1905_v29_ = _dafny.Map({d_1903_v27_.f27: not(d_1892_v16_.f25)})
                        d_1906_v30_: _dafny.Array
                        nw323_ = _dafny.Array(None, 7)
                        nw323_[int(0)] = (len(d_1898_v22_)) <= (default__.fm0(d_1897_v21_, d_1897_v21_, default__.fm0(d_1897_v21_, len((d_1893_v17_)[default__.safeIndex(679, (d_1893_v17_).length(0))]), len(_dafny.SeqWithoutIsStrInference([d_1897_v21_ for d_1907_i3_ in range(default__.abs(19))])), (d_1899_v23_)[default__.safeIndex(d_1897_v21_, len(d_1899_v23_))], globalState), (d_1892_v16_).f26, globalState))
                        nw323_[int(1)] = ((d_1902_v26_)[d_1900_v24_] if (d_1900_v24_) in (d_1902_v26_) else d_1870_v0_)
                        nw323_[int(2)] = not((not(((d_1904_v28_)[d_1903_v27_] if (d_1903_v27_) in (d_1904_v28_) else d_1892_v16_.f25))) == (False))
                        nw323_[int(3)] = default__.fm2(d_1905_v29_, globalState)
                        nw323_[int(4)] = (True) and (not(False))
                        nw323_[int(5)] = (d_1870_v0_) == (not(d_1870_v0_))
                        nw323_[int(6)] = False
                        d_1906_v30_ = nw323_
                        index329_ = default__.safeIndex(494, (d_1906_v30_).length(0))
                        (d_1906_v30_)[index329_] = (d_1896_v20_) not in (d_1900_v24_)
                        index330_ = default__.safeIndex(494, (d_1906_v30_).length(0))
                        (d_1906_v30_)[index330_] = (d_1892_v16_).f26
                        d_1908_v31_: _dafny.Array
                        nw324_ = _dafny.Array(int(0), 11)
                        d_1908_v31_ = nw324_
                        d_1908_v31_ = d_1908_v31_
                    d_1909_v32_: int
                    d_1909_v32_ = 678
                    (globalState).f7 = (0) - (d_1909_v32_)
                    d_1910_v33_: str
                    d_1910_v33_ = _dafny.CodePoint('p')
                    d_1911_v34_: _dafny.Map
                    d_1911_v34_ = _dafny.Map({d_1870_v0_: d_1870_v0_})
                    d_1912_v35_: _dafny.Array
                    nw325_ = _dafny.Array(None, 2)
                    nw325_[int(0)] = default__.fm18(True, d_1910_v33_, False, globalState)
                    nw325_[int(1)] = (_dafny.SeqWithoutIsStrInference([d_1870_v0_])) + (default__.fm18(True, d_1910_v33_, default__.fm2(d_1911_v34_, globalState), globalState))
                    d_1912_v35_ = nw325_
                    d_1913_v36_: _dafny.Seq
                    d_1913_v36_ = _dafny.SeqWithoutIsStrInference([d_1870_v0_, not(d_1870_v0_)])
                    index331_ = default__.safeIndex(847, (d_1912_v35_).length(0))
                    (d_1912_v35_)[index331_] = d_1913_v36_
                    index332_ = default__.safeIndex(847, (d_1912_v35_).length(0))
                    (d_1912_v35_)[index332_] = _dafny.SeqWithoutIsStrInference([d_1870_v0_, d_1870_v0_])
                    if d_1870_v0_:
                        d_1914_v37_: _dafny.Seq
                        d_1914_v37_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dhdknns"))
                        d_1915_v38_: _dafny.Map
                        d_1915_v38_ = _dafny.Map({d_1870_v0_: (d_1912_v35_)[default__.safeIndex(847, (d_1912_v35_).length(0))]})
                        d_1916_v39_: _dafny.Map
                        d_1916_v39_ = _dafny.Map({d_1909_v32_: d_1909_v32_})
                        d_1917_v41_: D13
                        d_1917_v41_ = D13_DC25(d_1916_v39_)
                        d_1918_v42_: _dafny.Map
                        def iife132_():
                            coll76_ = _dafny.Map()
                            compr_76_: int
                            for compr_76_ in _dafny.IntegerRange(368, 628):
                                d_1919_v40_: int = compr_76_
                                if ((368) <= (d_1919_v40_)) and ((d_1919_v40_) < (628)):
                                    coll76_[(d_1919_v40_) - (d_1909_v32_)] = d_1870_v0_
                            return _dafny.Map(coll76_)
                        d_1918_v42_ = _dafny.Map({len((d_1913_v36_).set(default__.safeIndex(len(iife132_()
                        ), len(d_1913_v36_)), d_1870_v0_)): d_1917_v41_})
                        d_1920_v43_: D17
                        d_1920_v43_ = D17_DC39(4, d_1916_v39_, d_1870_v0_, d_1918_v42_, d_1909_v32_)
                        d_1921_v44_: _dafny.Map
                        d_1921_v44_ = _dafny.Map({d_1909_v32_: d_1870_v0_})
                        d_1922_v45_: _dafny.Seq
                        d_1922_v45_ = _dafny.SeqWithoutIsStrInference([d_1921_v44_])
                        d_1923_v46_: C1
                        nw326_ = C1()
                        nw326_.ctor__(default__.safeDivisionInt(default__.fm0(len(d_1914_v37_), len(d_1915_v38_), d_1909_v32_, (d_1920_v43_).cf54, globalState), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cfemxm")))), (_dafny.SeqWithoutIsStrInference([d_1921_v44_])) + (d_1922_v45_))
                        d_1923_v46_ = nw326_
                        d_1924_v47_: _dafny.Seq
                        d_1924_v47_ = _dafny.SeqWithoutIsStrInference([d_1909_v32_])
                        d_1925_v48_: _dafny.Array
                        def lambda107_(d_1926_v0_):
                            def lambda108_(d_1927_i4_):
                                return d_1926_v0_

                            return lambda108_

                        init59_ = lambda107_(d_1870_v0_)
                        nw327_ = _dafny.Array(None, 1)
                        for i0_59_ in range(nw327_.length(0)):
                            nw327_[i0_59_] = init59_(i0_59_)
                        d_1925_v48_ = nw327_
                        d_1928_v49_: _dafny.Array
                        d_1928_v49_ = d_1925_v48_
                        d_1929_v50_: _dafny.Array
                        nw328_ = _dafny.Array(None, 9)
                        nw328_[int(0)] = d_1925_v48_
                        nw328_[int(1)] = d_1925_v48_
                        nw328_[int(2)] = d_1925_v48_
                        nw328_[int(3)] = d_1925_v48_
                        nw328_[int(4)] = (d_1928_v49_)
                        nw328_[int(5)] = d_1925_v48_
                        nw328_[int(6)] = d_1925_v48_
                        nw328_[int(7)] = d_1925_v48_
                        nw328_[int(8)] = d_1925_v48_
                        d_1929_v50_ = nw328_
                        d_1930_v51_: _dafny.Map
                        d_1930_v51_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([d_1923_v46_.f29, 310])) < (d_1924_v47_): d_1929_v50_})
                        d_1930_v51_ = (d_1930_v51_).set((d_1923_v46_.f29) <= (-382), d_1929_v50_)
                        d_1931_v52_: D18
                        d_1931_v52_ = D18_DC43()
                        d_1932_v53_: _dafny.Map
                        d_1932_v53_ = _dafny.Map({d_1909_v32_: d_1931_v52_})
                        (globalState).f2 = ((467 if d_1870_v0_ else d_1909_v32_)) not in (d_1932_v53_)
                        d_1933_v54_: _dafny.Array
                        nw329_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 16)
                        d_1933_v54_ = nw329_
                        d_1933_v54_ = d_1933_v54_
                        d_1934_v55_: C3
                        nw330_ = C3()
                        nw330_.ctor__()
                        d_1934_v55_ = nw330_
                    elif True:
                        d_1935_v56_: _dafny.Seq
                        d_1935_v56_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ap"))
                        d_1936_v57_: D13
                        d_1936_v57_ = D13_DC26(d_1935_v56_, d_1870_v0_, d_1870_v0_)
                        d_1937_v58_: _dafny.Map
                        d_1937_v58_ = _dafny.Map({(d_1936_v57_).cf33: d_1870_v0_})
                        d_1937_v58_ = (d_1937_v58_).set(d_1935_v56_, not(True))
                        d_1938_v59_: _dafny.Map
                        d_1938_v59_ = _dafny.Map({d_1909_v32_: d_1870_v0_})
                        d_1939_v60_: _dafny.Seq
                        d_1939_v60_ = _dafny.SeqWithoutIsStrInference([d_1938_v59_])
                        d_1940_v61_: C1
                        nw331_ = C1()
                        nw331_.ctor__(len(d_1935_v56_), d_1939_v60_)
                        d_1940_v61_ = nw331_
                        d_1941_v62_: _dafny.Seq
                        d_1941_v62_ = _dafny.SeqWithoutIsStrInference([d_1940_v61_, d_1940_v61_, d_1940_v61_, d_1940_v61_, d_1940_v61_])
                        d_1942_v63_: _dafny.Map
                        d_1942_v63_ = _dafny.Map({d_1913_v36_: d_1935_v56_})
                        default__.m0((d_1941_v62_) == (d_1941_v62_), d_1909_v32_, (d_1942_v63_).set((d_1912_v35_)[default__.safeIndex(847, (d_1912_v35_).length(0))], d_1935_v56_), len(d_1911_v34_), globalState)
                        (globalState).f19 = d_1935_v56_
                        d_1943_v64_: D0
                        d_1943_v64_ = D0_DC0(default__.fm2(d_1911_v34_, globalState))
                        d_1944_v65_: D0
                        d_1944_v65_ = D0_DC1(d_1943_v64_)
                        d_1945_v66_: _dafny.MultiSet
                        d_1945_v66_ = _dafny.MultiSet([d_1944_v65_, d_1944_v65_, d_1944_v65_])
                        (globalState).f16 = default__.safeDivisionInt(default__.safeModuloInt((0) - (-782), d_1909_v32_), (0) - (((d_1945_v66_)[d_1944_v65_] if (d_1944_v65_) in (d_1945_v66_) else d_1940_v61_.f29)))
                        d_1946_v67_: _dafny.Map
                        d_1946_v67_ = _dafny.Map({d_1870_v0_: d_1940_v61_.f29})
                        d_1947_v68_: _dafny.Map
                        d_1947_v68_ = _dafny.Map({default__.fm56(d_1940_v61_.f29, len(d_1946_v67_), d_1870_v0_, globalState): (d_1940_v61_).fm29(globalState)})
                        d_1947_v68_ = d_1947_v68_
                    pass
            pass
        d_1948_v69_: int
        d_1948_v69_ = 816
        index333_ = default__.safeIndex(213, (p0).length(0))
        (p0)[index333_] = d_1948_v69_
        d_1949_v70_: _dafny.Array
        nw332_ = _dafny.Array(_dafny.Set({}), 22)
        d_1949_v70_ = nw332_
        d_1950_v71_: D7
        d_1950_v71_ = D7_DC16(d_1949_v70_)
        d_1951_v72_: _dafny.MultiSet
        d_1951_v72_ = _dafny.MultiSet([d_1950_v71_])
        index334_ = default__.safeIndex(213, (p0).length(0))
        (p0)[index334_] = ((d_1951_v72_).set(d_1950_v71_, default__.abs(d_1948_v69_))).cardinality
        d_1952_v73_: C5
        nw333_ = C5()
        nw333_.ctor__(d_1870_v0_)
        d_1952_v73_ = nw333_
        d_1953_v74_: _dafny.MultiSet
        d_1953_v74_ = _dafny.MultiSet([p0])
        source24_ = d_1953_v74_
        d_1954___mcc_h0_ = source24_
        d_1955_cf4_ = d_1954___mcc_h0_
        d_1956_v75_: str
        d_1956_v75_ = _dafny.CodePoint('x')
        d_1957_v76_: _dafny.Seq
        d_1957_v76_ = _dafny.SeqWithoutIsStrInference([d_1956_v75_])
        d_1948_v69_ = default__.safeModuloInt(default__.safeDivisionInt(d_1948_v69_, len(d_1957_v76_)), d_1948_v69_)
        d_1958_v77_: _dafny.Array
        nw334_ = _dafny.Array(D19.default()(), 22)
        d_1958_v77_ = nw334_
        d_1959_v78_: _dafny.Array
        def lambda109_(d_1960_v73_):
            def lambda110_(d_1961_i5_):
                return (d_1960_v73_).f24

            return lambda110_

        init60_ = lambda109_(d_1952_v73_)
        nw335_ = _dafny.Array(None, 6)
        for i0_60_ in range(nw335_.length(0)):
            nw335_[i0_60_] = init60_(i0_60_)
        d_1959_v78_ = nw335_
        d_1962_v79_: D19
        d_1962_v79_ = D19_DC45(_dafny.MultiSet([d_1959_v78_]))
        index335_ = default__.safeIndex(212, (d_1958_v77_).length(0))
        (d_1958_v77_)[index335_] = d_1962_v79_
        index336_ = default__.safeIndex(212, (d_1958_v77_).length(0))
        (d_1958_v77_)[index336_] = d_1962_v79_
        d_1963_v80_: _dafny.MultiSet
        d_1963_v80_ = _dafny.MultiSet([(d_1952_v73_).f24])
        rhs325_ = d_1963_v80_
        rhs326_ = True
        lhs306_ = globalState
        d_1963_v80_ = rhs325_
        lhs306_.f2 = rhs326_
        (globalState).f13 = (d_1952_v73_).f24
        pat_let_tv33_ = d_1948_v69_
        pat_let_tv34_ = d_1952_v73_
        pat_let_tv35_ = d_1948_v69_
        pat_let_tv36_ = d_1948_v69_
        pat_let_tv37_ = d_1948_v69_
        def lambda111_(source25_):
            if source25_.is_DC36:
                d_1964___mcc_h1_ = source25_.cf47
                d_1965___mcc_h2_ = source25_.cf48
                d_1966___mcc_h3_ = source25_.cf49
                d_1967_cf49_ = d_1966___mcc_h3_
                d_1968_cf48_ = d_1965___mcc_h2_
                d_1969_cf47_ = d_1964___mcc_h1_
                return (pat_let_tv33_) * (-220)
            elif source25_.is_DC35:
                d_1970___mcc_h4_ = source25_.cf46
                d_1971_cf46_ = d_1970___mcc_h4_
                if (pat_let_tv34_).f24:
                    return pat_let_tv35_
                elif True:
                    return len(_dafny.SeqWithoutIsStrInference([pat_let_tv36_]))
            elif True:
                d_1972___mcc_h5_ = source25_.cf50
                d_1973_cf50_ = d_1972___mcc_h5_
                return (len(_dafny.Set({(0) - (pat_let_tv37_)}))) - (405)

        (globalState).f7 = lambda111_(default__.fm57(globalState))
        d_1974_v81_: str
        d_1974_v81_ = _dafny.CodePoint('y')
        d_1974_v81_ = _dafny.CodePoint('c')

    def m1(self, p0, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: bool = False
        d_1975_v0_: bool
        d_1975_v0_ = True
        d_1976_v1_: _dafny.Seq
        d_1976_v1_ = _dafny.SeqWithoutIsStrInference([d_1975_v0_])
        d_1977_v2_: _dafny.MultiSet
        d_1977_v2_ = _dafny.MultiSet([p0, p0])
        d_1978_v3_: C4
        nw336_ = C4()
        nw336_.ctor__(((d_1976_v1_)[default__.safeIndex(p0, len(d_1976_v1_))]) and (d_1975_v0_), (d_1977_v2_) in (_dafny.Map({d_1977_v2_: p0})))
        d_1978_v3_ = nw336_
        if (d_1978_v3_).f26:
            d_1979_v4_: _dafny.Map
            d_1979_v4_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iuyqqd"))): p0})
            d_1980_v5_: _dafny.Map
            d_1980_v5_ = _dafny.Map({p0: default__.fm45(globalState)})
            d_1981_v6_: _dafny.MultiSet
            d_1981_v6_ = _dafny.MultiSet([True, d_1978_v3_.f25, d_1975_v0_])
            d_1982_v7_: D17
            d_1982_v7_ = D17_DC39(p0, (d_1979_v4_) | (default__.fm54(globalState)), not((d_1978_v3_).f26), d_1980_v5_, (d_1981_v6_).cardinality)
            d_1982_v7_ = d_1982_v7_
            d_1983_v8_: _dafny.Set
            d_1983_v8_ = _dafny.Set({-285})
            d_1984_v9_: _dafny.Map
            d_1984_v9_ = _dafny.Map({d_1983_v8_: not(d_1978_v3_.f25)})
            d_1985_v10_: _dafny.Array
            nw337_ = _dafny.Array(None, 11)
            nw337_[int(0)] = (d_1975_v0_) or ((d_1978_v3_).f26)
            nw337_[int(1)] = False
            nw337_[int(2)] = (d_1983_v8_) not in (_dafny.Map({d_1983_v8_: p0}))
            nw337_[int(3)] = (d_1978_v3_).f26
            nw337_[int(4)] = False
            nw337_[int(5)] = d_1978_v3_.f25
            nw337_[int(6)] = ((d_1984_v9_)[d_1983_v8_] if (d_1983_v8_) in (d_1984_v9_) else d_1975_v0_)
            nw337_[int(7)] = d_1978_v3_.f25
            nw337_[int(8)] = d_1978_v3_.f25
            nw337_[int(9)] = (_dafny.Set({p0})).ispropersubset(d_1983_v8_)
            nw337_[int(10)] = d_1978_v3_.f25
            d_1985_v10_ = nw337_
            d_1985_v10_ = d_1985_v10_
            d_1986_v11_: C0
            nw338_ = C0()
            nw338_.ctor__((d_1978_v3_).f26)
            d_1986_v11_ = nw338_
            d_1987_v12_: _dafny.Map
            d_1987_v12_ = _dafny.Map({d_1986_v11_: d_1978_v3_.f25})
            d_1988_v13_: _dafny.Map
            d_1988_v13_ = _dafny.Map({(d_1987_v12_) | (d_1987_v12_): d_1986_v11_.f27})
            d_1989_v14_: _dafny.Map
            d_1989_v14_ = _dafny.Map({p0: d_1987_v12_})
            d_1990_v15_: _dafny.Map
            d_1990_v15_ = _dafny.Map({p0: d_1978_v3_.f25})
            d_1991_v16_: _dafny.Seq
            d_1991_v16_ = _dafny.SeqWithoutIsStrInference([len(d_1976_v1_), len(d_1990_v15_)])
            d_1992_v17_: _dafny.Map
            d_1992_v17_ = _dafny.Map({d_1976_v1_: d_1991_v16_})
            d_1993_v18_: D21
            d_1993_v18_ = D21_DC52(d_1992_v17_)
            pat_let_tv38_ = d_1992_v17_
            d_1994_v19_: _dafny.Map
            def iife133_(_pat_let28_0):
                def iife134_(d_1995_dt__update__tmp_h0_):
                    def iife135_(_pat_let29_0):
                        def iife136_(d_1996_dt__update_hcf71_h0_):
                            return D21_DC52(d_1996_dt__update_hcf71_h0_)
                        return iife136_(_pat_let29_0)
                    return iife135_(pat_let_tv38_)
                return iife134_(_pat_let28_0)
            d_1994_v19_ = _dafny.Map({iife133_(d_1993_v18_): d_1986_v11_})
            d_1997_v20_: _dafny.Set
            d_1997_v20_ = _dafny.Set({((d_1994_v19_)[d_1993_v18_] if (d_1993_v18_) in (d_1994_v19_) else d_1986_v11_), d_1986_v11_})
            d_1998_v21_: _dafny.Map
            d_1998_v21_ = _dafny.Map({d_1986_v11_.f27: d_1997_v20_})
            d_1988_v13_ = (d_1988_v13_).set((d_1987_v12_) | (((d_1989_v14_)[p0] if (p0) in (d_1989_v14_) else d_1987_v12_)), (((d_1998_v21_)[d_1986_v11_.f27] if (d_1986_v11_.f27) in (d_1998_v21_) else d_1997_v20_)).issubset(d_1997_v20_))
            d_1999_v22_: _dafny.Array
            nw339_ = _dafny.Array(int(0), 13)
            d_1999_v22_ = nw339_
            d_2000_v23_: _dafny.Seq
            d_2000_v23_ = _dafny.SeqWithoutIsStrInference([d_1999_v22_])
            d_2001_v24_: D6
            d_2001_v24_ = D6_DC12(d_2000_v23_)
            d_2002_v25_: D6
            d_2002_v25_ = D6_DC14(d_2001_v24_)
            d_2003_v26_: D0
            d_2003_v26_ = D0_DC0(d_1978_v3_.f25)
            d_2004_v27_: _dafny.Map
            d_2004_v27_ = _dafny.Map({d_2003_v26_: p0})
            d_2005_v28_: str
            d_2005_v28_ = _dafny.CodePoint('n')
            d_2006_v29_: _dafny.Set
            d_2006_v29_ = _dafny.Set({d_2005_v28_})
            d_2007_v30_: _dafny.MultiSet
            d_2007_v30_ = _dafny.MultiSet([d_2006_v29_, _dafny.Set({_dafny.CodePoint('x')}), _dafny.Set({d_2005_v28_})])
            d_2008_v31_: _dafny.MultiSet
            d_2008_v31_ = _dafny.MultiSet([d_1985_v10_])
            rhs327_ = ((p0) + (691)) in (_dafny.SeqWithoutIsStrInference([970, p0, len(d_2004_v27_), p0, (0) - (default__.fm0(p0, (0) - (p0), p0, (self).fm7(False, p0, d_1978_v3_.f25, globalState), globalState))]))
            rhs328_ = (0) - (((d_2007_v30_)[d_2006_v29_] if (d_2006_v29_) in (d_2007_v30_) else (p0) - (p0)))
            rhs329_ = d_2002_v25_
            rhs330_ = p0
            rhs331_ = ((d_2008_v31_)[d_1985_v10_] if (d_1985_v10_) in (d_2008_v31_) else 384)
            lhs307_ = globalState
            lhs308_ = globalState
            lhs309_ = globalState
            lhs310_ = globalState
            lhs307_.f21 = rhs327_
            lhs308_.f16 = rhs328_
            d_2002_v25_ = rhs329_
            lhs309_.f7 = rhs330_
            lhs310_.f7 = rhs331_
            d_2009_v32_: _dafny.Seq
            d_2009_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "as"))
            (globalState).f0 = ((d_2009_v32_) <= (d_2009_v32_) if not(not (d_1978_v3_.f25) or (d_1978_v3_.f25)) else (d_1978_v3_).f26)
        elif True:
            d_2010_v33_: _dafny.Array
            def lambda112_(d_2011_v2_):
                def lambda113_(d_2012_i0_):
                    return d_2011_v2_

                return lambda113_

            init61_ = lambda112_(d_1977_v2_)
            nw340_ = _dafny.Array(None, 2)
            for i0_61_ in range(nw340_.length(0)):
                nw340_[i0_61_] = init61_(i0_61_)
            d_2010_v33_ = nw340_
            d_2010_v33_ = d_2010_v33_
            d_2013_v34_: _dafny.Array
            nw341_ = _dafny.Array(D12.default()(), 11)
            d_2013_v34_ = nw341_
            d_2014_v35_: D12
            d_2014_v35_ = D12_DC23(p0)
            d_2015_v36_: D12
            d_2015_v36_ = D12_DC24(d_2014_v35_)
            index337_ = default__.safeIndex(627, (d_2013_v34_).length(0))
            (d_2013_v34_)[index337_] = d_2015_v36_
            pat_let_tv39_ = d_2014_v35_
            index338_ = default__.safeIndex(627, (d_2013_v34_).length(0))
            def iife137_(_pat_let30_0):
                def iife138_(d_2016_dt__update__tmp_h1_):
                    def iife139_(_pat_let31_0):
                        def iife140_(d_2017_dt__update_hcf31_h0_):
                            return D12_DC24(d_2017_dt__update_hcf31_h0_)
                        return iife140_(_pat_let31_0)
                    return iife139_(pat_let_tv39_)
                return iife138_(_pat_let30_0)
            (d_2013_v34_)[index338_] = iife137_(D12_DC24(d_2014_v35_))
            d_2018_v37_: C6
            nw342_ = C6()
            nw342_.ctor__()
            d_2018_v37_ = nw342_
            d_2019_v38_: _dafny.Set
            d_2019_v38_ = _dafny.Set({d_1978_v3_.f25, (d_1978_v3_).f26})
            (d_1978_v3_).m5((d_2019_v38_).isdisjoint(d_2019_v38_), globalState)
            d_2020_v39_: _dafny.Map
            d_2020_v39_ = _dafny.Map({(d_1978_v3_).f26: (d_1978_v3_).f26})
            (globalState).f21 = default__.fm2(d_2020_v39_, globalState)
        d_2021_v40_: str
        d_2021_v40_ = _dafny.CodePoint('c')
        d_2022_v41_: _dafny.Map
        d_2022_v41_ = _dafny.Map({d_2021_v40_: not(d_1975_v0_)})
        d_2022_v41_ = (d_2022_v41_).set(d_2021_v40_, d_1975_v0_)
        if (p0) == (306):
            (globalState).f16 = 836
            d_2023_v42_: _dafny.Map
            d_2023_v42_ = _dafny.Map({d_1977_v2_: d_1978_v3_.f25})
            d_2024_v43_: _dafny.Map
            d_2024_v43_ = (_dafny.Map({d_1977_v2_: (d_1978_v3_).f26})) | (d_2023_v42_)
            d_2025_v44_: D12
            d_2025_v44_ = D12_DC23(p0)
            d_2026_v45_: D12
            d_2026_v45_ = D12_DC24(d_2025_v44_)
            rhs332_ = d_2024_v43_
            rhs333_ = default__.fm58(p0, globalState)
            d_2024_v43_ = rhs332_
            d_2026_v45_ = rhs333_
            d_2027_v46_: _dafny.Array
            nw343_ = _dafny.Array(False, 16)
            d_2027_v46_ = nw343_
            d_2027_v46_ = d_2027_v46_
            d_2028_v47_: T0
            nw344_ = C3()
            nw344_.ctor__()
            d_2028_v47_ = nw344_
            d_2029_v48_: _dafny.Seq
            d_2029_v48_ = _dafny.SeqWithoutIsStrInference([d_2028_v47_])
            d_2030_v49_: _dafny.Seq
            d_2030_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ggly"))
            d_2031_v50_: D24
            d_2031_v50_ = D24_DC63(d_2028_v47_)
            d_2032_v51_: _dafny.Array
            nw345_ = _dafny.Array(None, 18)
            nw345_[int(0)] = (d_2029_v48_)[default__.safeIndex(len(d_2030_v49_), len(d_2029_v48_))]
            nw345_[int(1)] = d_2028_v47_
            nw345_[int(2)] = d_2028_v47_
            nw345_[int(3)] = d_2028_v47_
            nw345_[int(4)] = d_2028_v47_
            nw345_[int(5)] = d_2028_v47_
            nw345_[int(6)] = d_2028_v47_
            nw345_[int(7)] = (d_2031_v50_).cf90
            nw345_[int(8)] = d_2028_v47_
            nw345_[int(9)] = (d_2028_v47_ if d_1978_v3_.f25 else d_2028_v47_)
            nw345_[int(10)] = d_2028_v47_
            nw345_[int(11)] = (d_2031_v50_).cf90
            nw345_[int(12)] = d_2028_v47_
            nw345_[int(13)] = d_2028_v47_
            nw345_[int(14)] = d_2028_v47_
            nw345_[int(15)] = d_2028_v47_
            nw345_[int(16)] = d_2028_v47_
            nw345_[int(17)] = (d_2028_v47_ if not(d_1978_v3_.f25) else d_2028_v47_)
            d_2032_v51_ = nw345_
            index339_ = default__.safeIndex(741, (d_2032_v51_).length(0))
            (d_2032_v51_)[index339_] = d_2028_v47_
            index340_ = default__.safeIndex(741, (d_2032_v51_).length(0))
            (d_2032_v51_)[index340_] = d_2028_v47_
            if not(not(False)):
                index341_ = default__.safeIndex(952, (d_2027_v46_).length(0))
                (d_2027_v46_)[index341_] = not (d_1978_v3_.f25) or ((d_1978_v3_).f26)
                index342_ = default__.safeIndex(952, (d_2027_v46_).length(0))
                (d_2027_v46_)[index342_] = not(not(d_1978_v3_.f25))
                (globalState).f0 = ((d_2030_v49_) <= (d_2030_v49_)) == ((d_1978_v3_).f26)
                d_2033_v52_: _dafny.Map
                d_2033_v52_ = _dafny.Map({d_1975_v0_: p0})
                d_2034_v53_: _dafny.Set
                d_2034_v53_ = _dafny.Set({len(d_2030_v49_)})
                d_2035_v54_: _dafny.Set
                d_2035_v54_ = _dafny.Set({_dafny.Set({p0, p0})})
                rhs334_ = (d_1976_v1_ if (d_2034_v53_) != (d_2034_v53_) else d_1976_v1_)
                rhs335_ = default__.safeModuloInt(p0, default__.fm0(p0, p0, p0, (d_1978_v3_).f26, globalState))
                rhs336_ = default__.fm35(d_1975_v0_, len(d_2035_v54_), globalState)
                rhs337_ = p0
                rhs338_ = d_2030_v49_
                lhs311_ = globalState
                lhs312_ = globalState
                lhs313_ = globalState
                d_1976_v1_ = rhs334_
                lhs311_.f16 = rhs335_
                d_2033_v52_ = rhs336_
                lhs312_.f16 = rhs337_
                lhs313_.f8 = rhs338_
                (globalState).f7 = (179) + (p0)
                d_2036_v55_: _dafny.Array
                def lambda114_(d_2037_v49_):
                    def lambda115_(d_2038_i1_):
                        return _dafny.SeqWithoutIsStrInference([d_2037_v49_, d_2037_v49_])

                    return lambda115_

                init62_ = lambda114_(d_2030_v49_)
                nw346_ = _dafny.Array(None, 27)
                for i0_62_ in range(nw346_.length(0)):
                    nw346_[i0_62_] = init62_(i0_62_)
                d_2036_v55_ = nw346_
                index343_ = default__.safeIndex(981, (d_2036_v55_).length(0))
                (d_2036_v55_)[index343_] = _dafny.SeqWithoutIsStrInference([d_2030_v49_])
                d_2039_v56_: _dafny.Seq
                d_2039_v56_ = _dafny.SeqWithoutIsStrInference([d_2030_v49_, d_2030_v49_, d_2030_v49_, d_2030_v49_])
                index344_ = default__.safeIndex(981, (d_2036_v55_).length(0))
                (d_2036_v55_)[index344_] = (d_2039_v56_) + (default__.fm59((d_1978_v3_).f26, globalState))
            elif True:
                d_2040_v57_: _dafny.MultiSet
                d_2040_v57_ = _dafny.MultiSet([d_1975_v0_])
                d_2041_v58_: _dafny.Array
                nw347_ = _dafny.Array(None, 1)
                nw347_[int(0)] = p0
                d_2041_v58_ = nw347_
                d_2042_v59_: D21
                d_2042_v59_ = D21_DC53(d_2027_v46_, False, ((d_2040_v57_).set(d_1975_v0_, default__.abs(len(d_2030_v49_)))) - (d_2040_v57_), d_2041_v58_)
                d_2042_v59_ = (d_2042_v59_ if not ((d_1978_v3_).f26) or (d_1978_v3_.f25) else d_2042_v59_)
                (globalState).f16 = p0
                d_2043_v60_: _dafny.Array
                nw348_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 11)
                d_2043_v60_ = nw348_
                d_2043_v60_ = d_2043_v60_
                (globalState).f8 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ydtter"))
                d_2044_v61_: C0
                nw349_ = C0()
                nw349_.ctor__(d_1975_v0_)
                d_2044_v61_ = nw349_
        elif True:
            d_2045_v62_: _dafny.Array
            def lambda116_(d_2046_p0_):
                def lambda117_(d_2047_i2_):
                    return (d_2047_i2_) + (len(_dafny.Map({d_2046_p0_: d_2046_p0_})))

                return lambda117_

            init63_ = lambda116_(p0)
            nw350_ = _dafny.Array(None, 7)
            for i0_63_ in range(nw350_.length(0)):
                nw350_[i0_63_] = init63_(i0_63_)
            d_2045_v62_ = nw350_
            d_2048_v63_: _dafny.Seq
            out8_: _dafny.Seq
            out8_ = (self).m4((_dafny.MultiSet([d_2045_v62_, d_2045_v62_])).cardinality, d_1975_v0_, 934, globalState)
            d_2048_v63_ = out8_
            d_2049_v64_: _dafny.Seq
            d_2049_v64_ = _dafny.SeqWithoutIsStrInference([d_2045_v62_])
            r0 = (d_2049_v64_)[default__.safeIndex(p0, len(d_2049_v64_))]
            d_2050_v65_: _dafny.Map
            d_2050_v65_ = _dafny.Map({p0: _dafny.CodePoint('v')})
            d_2051_v66_: _dafny.Map
            d_2051_v66_ = _dafny.Map({d_1978_v3_.f25: d_1978_v3_.f25})
            d_2052_v67_: _dafny.Map
            d_2052_v67_ = _dafny.Map({(d_2050_v65_) | (d_2050_v65_): d_2051_v66_})
            d_2052_v67_ = default__.fm60(globalState)
            d_2053_v68_: _dafny.Set
            d_2053_v68_ = _dafny.Set({d_1978_v3_.f25})
            d_2054_v69_: _dafny.Map
            d_2054_v69_ = _dafny.Map({p0: len(d_2053_v68_)})
            d_2055_v70_: _dafny.Seq
            d_2055_v70_ = _dafny.SeqWithoutIsStrInference([d_2054_v69_])
            rhs339_ = (d_1978_v3_).f26
            rhs340_ = (p0) - (p0)
            rhs341_ = (len(d_2055_v70_)) - (p0)
            lhs314_ = d_1978_v3_
            lhs315_ = globalState
            lhs316_ = globalState
            lhs314_.f25 = rhs339_
            lhs315_.f16 = rhs340_
            lhs316_.f16 = rhs341_
            index345_ = default__.safeIndex(612, (d_2045_v62_).length(0))
            (d_2045_v62_)[index345_] = ((d_1977_v2_).intersection(d_1977_v2_)).cardinality
            index346_ = default__.safeIndex(612, (d_2045_v62_).length(0))
            (d_2045_v62_)[index346_] = p0
        d_2056_v71_: _dafny.MultiSet
        d_2056_v71_ = _dafny.MultiSet([d_2021_v40_, d_2021_v40_, default__.fm15(_dafny.SeqWithoutIsStrInference([214 for d_2057_i3_ in range(default__.abs(-90))]), p0, globalState)])
        d_2058_v72_: _dafny.Seq
        d_2058_v72_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))
        d_2059_v73_: _dafny.MultiSet
        d_2059_v73_ = _dafny.MultiSet([(d_1978_v3_).f26])
        rhs342_ = ((_dafny.MultiSet([(d_2058_v72_)[default__.safeIndex((d_2059_v73_).cardinality, len(d_2058_v72_))]])).set(d_2021_v40_, default__.abs(p0))) | ((d_2056_v71_).set(d_2021_v40_, default__.abs(len(default__.fm1(502, globalState)))))
        rhs343_ = d_1978_v3_.f25
        rhs344_ = (_dafny.CodePoint('x') if not((706) < (212)) else d_2021_v40_)
        rhs345_ = p0
        lhs317_ = d_1978_v3_
        lhs318_ = globalState
        d_2056_v71_ = rhs342_
        lhs317_.f25 = rhs343_
        d_2021_v40_ = rhs344_
        lhs318_.f7 = rhs345_
        d_2060_v74_: _dafny.Array
        nw351_ = _dafny.Array(int(0), 15)
        d_2060_v74_ = nw351_
        guard_loop_11_: int
        for guard_loop_11_ in _dafny.IntegerRange(0, (d_2060_v74_).length(0)):
            d_2061_i4_: int = guard_loop_11_
            if (True) and (((0) <= (d_2061_i4_)) and ((d_2061_i4_) < ((d_2060_v74_).length(0)))):
                (d_2060_v74_)[(d_2061_i4_)] = (d_2061_i4_) - (p0)
        r0 = d_2060_v74_
        d_2062_v75_: _dafny.Set
        d_2062_v75_ = _dafny.Set({d_2060_v74_, d_2060_v74_})
        r1 = not((p0) != (len(d_2062_v75_)))
        return r0, r1

    def m4(self, p0, p1, p2, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_2063_v0_: _dafny.Map
        d_2063_v0_ = _dafny.Map({False: p1})
        d_2064_v1_: _dafny.Set
        d_2064_v1_ = _dafny.Set({p1})
        d_2065_v2_: str
        d_2065_v2_ = _dafny.CodePoint('d')
        d_2066_v3_: _dafny.Seq
        d_2066_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yjscrou"))
        d_2067_v4_: _dafny.Array
        nw352_ = _dafny.Array(None, 16)
        nw352_[int(0)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_2068_i0_ in range(default__.abs(-831))]))
        nw352_[int(1)] = p2
        nw352_[int(2)] = (0) - (p2)
        nw352_[int(3)] = default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([False, True])), default__.fm0(p2, p2, p2, p1, globalState))
        nw352_[int(4)] = (len(d_2063_v0_)) * (p2)
        nw352_[int(5)] = p0
        nw352_[int(6)] = p2
        nw352_[int(7)] = (0) - (p0)
        nw352_[int(8)] = len((_dafny.Set({False})) | (d_2064_v1_))
        nw352_[int(9)] = 359
        nw352_[int(10)] = p2
        nw352_[int(11)] = (len(_dafny.Map({d_2065_v2_: d_2065_v2_}))) - (303)
        nw352_[int(12)] = p0
        nw352_[int(13)] = p0
        nw352_[int(14)] = (0) - (len(d_2066_v3_))
        nw352_[int(15)] = p0
        d_2067_v4_ = nw352_
        index347_ = default__.safeIndex(676, (d_2067_v4_).length(0))
        (d_2067_v4_)[index347_] = p0
        index348_ = default__.safeIndex(676, (d_2067_v4_).length(0))
        (d_2067_v4_)[index348_] = default__.fm0(p2, p2, p0, (p1 if p1 else p1), globalState)
        d_2069_v5_: _dafny.Seq
        d_2069_v5_ = _dafny.SeqWithoutIsStrInference([d_2067_v4_, d_2067_v4_, d_2067_v4_, d_2067_v4_, d_2067_v4_])
        d_2070_v6_: D6
        d_2070_v6_ = D6_DC12(d_2069_v5_)
        pat_let_tv40_ = d_2067_v4_
        def iife141_(_pat_let32_0):
            def iife142_(d_2071_dt__update__tmp_h0_):
                def iife143_(_pat_let33_0):
                    def iife144_(d_2072_dt__update_hcf18_h0_):
                        return D6_DC12(d_2072_dt__update_hcf18_h0_)
                    return iife144_(_pat_let33_0)
                return iife143_(_dafny.SeqWithoutIsStrInference([pat_let_tv40_]))
            return iife142_(_pat_let32_0)
        d_2069_v5_ = ((d_2069_v5_) + (d_2069_v5_) if p1 else (iife141_(d_2070_v6_)).cf18)
        (globalState).f7 = ((d_2067_v4_)[default__.safeIndex(676, (d_2067_v4_).length(0))] if p1 else p0)
        d_2073_v7_: C6
        nw353_ = C6()
        nw353_.ctor__()
        d_2073_v7_ = nw353_
        d_2073_v7_ = d_2073_v7_
        guard_loop_12_: int
        for guard_loop_12_ in _dafny.IntegerRange(0, (d_2067_v4_).length(0)):
            d_2074_i1_: int = guard_loop_12_
            if (True) and (((0) <= (d_2074_i1_)) and ((d_2074_i1_) < ((d_2067_v4_).length(0)))):
                (d_2067_v4_)[(d_2074_i1_)] = (d_2074_i1_) + (len(d_2066_v3_))
        pat_let_tv41_ = d_2069_v5_
        def iife145_(_pat_let34_0):
            def iife146_(d_2075_dt__update__tmp_h1_):
                def iife147_(_pat_let35_0):
                    def iife148_(d_2076_dt__update_hcf18_h1_):
                        return D6_DC12(d_2076_dt__update_hcf18_h1_)
                    return iife148_(_pat_let35_0)
                return iife147_(pat_let_tv41_)
            return iife146_(_pat_let34_0)
        source26_ = iife145_(d_2070_v6_)
        if source26_.is_DC13:
            d_2077___mcc_h0_ = source26_.cf19
            d_2078___mcc_h1_ = source26_.cf20
            d_2079_cf20_ = d_2078___mcc_h1_
            d_2080_cf19_ = d_2077___mcc_h0_
            index349_ = default__.safeIndex(676, (d_2067_v4_).length(0))
            (d_2067_v4_)[index349_] = (0) - ((0) - (p0))
            (globalState).f2 = False
            d_2081_v8_: _dafny.Seq
            d_2081_v8_ = _dafny.SeqWithoutIsStrInference([p1, p1, True])
            d_2082_v9_: _dafny.Seq
            d_2082_v9_ = _dafny.SeqWithoutIsStrInference([d_2081_v8_])
            d_2083_v10_: _dafny.Seq
            d_2083_v10_ = _dafny.SeqWithoutIsStrInference([len((d_2082_v9_)[default__.safeIndex(d_2080_cf19_, len(d_2082_v9_))]), -542, (d_2067_v4_)[default__.safeIndex(676, (d_2067_v4_).length(0))]])
            d_2084_v11_: _dafny.Map
            d_2084_v11_ = _dafny.Map({(d_2083_v10_)[default__.safeIndex(len(d_2066_v3_), len(d_2083_v10_))]: p2})
            d_2085_v12_: _dafny.Array
            def lambda118_(d_2086_v2_):
                def lambda119_(d_2087_i2_):
                    return d_2086_v2_

                return lambda119_

            init64_ = lambda118_(d_2065_v2_)
            nw354_ = _dafny.Array(None, 14)
            for i0_64_ in range(nw354_.length(0)):
                nw354_[i0_64_] = init64_(i0_64_)
            d_2085_v12_ = nw354_
            d_2088_v13_: _dafny.Map
            d_2088_v13_ = _dafny.Map({p1: d_2085_v12_})
            d_2084_v11_ = (d_2084_v11_).set(len(d_2088_v13_), p2)
            d_2089_v14_: _dafny.MultiSet
            d_2089_v14_ = _dafny.MultiSet([p1, p1, p1])
            d_2090_v15_: _dafny.Map
            d_2090_v15_ = _dafny.Map({p1: (d_2089_v14_) | (d_2089_v14_)})
            d_2090_v15_ = (d_2090_v15_).set((True if p1 else p1), default__.fm52(p1, default__.fm0(len(d_2066_v3_), (0) - (d_2080_cf19_), p2, (d_2081_v8_)[default__.safeIndex(d_2080_cf19_, len(d_2081_v8_))], globalState), globalState))
        elif source26_.is_DC12:
            d_2091___mcc_h2_ = source26_.cf18
            d_2092_cf18_ = d_2091___mcc_h2_
            (globalState).f2 = (p1 if p1 else p1)
            d_2093_v16_: _dafny.Array
            nw355_ = _dafny.Array(None, 25)
            nw355_[int(0)] = p1
            nw355_[int(1)] = p1
            nw355_[int(2)] = p1
            nw355_[int(3)] = p1
            nw355_[int(4)] = (p2) <= ((d_2067_v4_)[default__.safeIndex(676, (d_2067_v4_).length(0))])
            nw355_[int(5)] = (p1) or (p1)
            nw355_[int(6)] = default__.fm2(d_2063_v0_, globalState)
            nw355_[int(7)] = p1
            nw355_[int(8)] = p1
            nw355_[int(9)] = p1
            nw355_[int(10)] = p1
            nw355_[int(11)] = p1
            nw355_[int(12)] = True
            nw355_[int(13)] = p1
            nw355_[int(14)] = p1
            nw355_[int(15)] = p1
            nw355_[int(16)] = p1
            nw355_[int(17)] = p1
            nw355_[int(18)] = p1
            nw355_[int(19)] = not(p1)
            nw355_[int(20)] = False
            nw355_[int(21)] = p1
            nw355_[int(22)] = p1
            nw355_[int(23)] = p1
            nw355_[int(24)] = p1
            d_2093_v16_ = nw355_
            d_2093_v16_ = d_2093_v16_
            d_2094_v17_: _dafny.Seq
            d_2094_v17_ = _dafny.SeqWithoutIsStrInference([d_2066_v3_, d_2066_v3_])
            d_2095_v18_: _dafny.Map
            d_2095_v18_ = _dafny.Map({102: p1})
            d_2096_v19_: _dafny.Map
            d_2096_v19_ = _dafny.Map({(d_2067_v4_)[default__.safeIndex(676, (d_2067_v4_).length(0))]: d_2095_v18_})
            d_2097_v20_: _dafny.Map
            d_2097_v20_ = _dafny.Map({len(_dafny.Map({p1: len(((d_2096_v19_)[170] if (170) in (d_2096_v19_) else d_2095_v18_))})): (d_2067_v4_)[default__.safeIndex(676, (d_2067_v4_).length(0))]})
            d_2098_v21_: D3
            d_2098_v21_ = D3_DC6(d_2066_v3_, d_2065_v2_, p1)
            d_2099_v22_: _dafny.Array
            nw356_ = _dafny.Array(None, 26)
            nw356_[int(0)] = ((d_2094_v17_)[default__.safeIndex(331, len(d_2094_v17_))]) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hesfue")))
            nw356_[int(1)] = (d_2066_v3_ if p1 else (d_2066_v3_).set(default__.safeIndex(p0, len(d_2066_v3_)), d_2065_v2_))
            nw356_[int(2)] = _dafny.SeqWithoutIsStrInference([d_2065_v2_ for d_2100_i3_ in range(default__.abs(-666))])
            nw356_[int(3)] = d_2066_v3_
            nw356_[int(4)] = (default__.fm1(len(d_2066_v3_), globalState)) + (d_2066_v3_)
            nw356_[int(5)] = (d_2066_v3_) + (d_2066_v3_)
            nw356_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "thegxd"))
            nw356_[int(7)] = (d_2066_v3_) + (d_2066_v3_)
            nw356_[int(8)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "npsshuqfe"))).set(default__.safeIndex(len(d_2097_v20_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "npsshuqfe")))), d_2065_v2_)
            nw356_[int(9)] = d_2066_v3_
            nw356_[int(10)] = d_2066_v3_
            nw356_[int(11)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cshpf"))
            nw356_[int(12)] = d_2066_v3_
            nw356_[int(13)] = (_dafny.SeqWithoutIsStrInference([d_2065_v2_ for d_2101_i4_ in range(default__.abs(-731))])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([d_2065_v2_ for d_2102_i4_ in range(default__.abs(-731))]))), d_2065_v2_)
            nw356_[int(14)] = d_2066_v3_
            nw356_[int(15)] = d_2066_v3_
            nw356_[int(16)] = (d_2066_v3_) + (d_2066_v3_)
            nw356_[int(17)] = (d_2094_v17_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kiiymoo"))), len(d_2094_v17_))]
            nw356_[int(18)] = d_2066_v3_
            nw356_[int(19)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "svnce"))
            nw356_[int(20)] = ((default__.fm1(p0, globalState)).set(default__.safeIndex(p0, len(default__.fm1(p0, globalState))), _dafny.CodePoint('d'))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xkvbkkxeo")))
            nw356_[int(21)] = d_2066_v3_
            nw356_[int(22)] = (d_2098_v21_).cf6
            nw356_[int(23)] = (d_2066_v3_).set(default__.safeIndex(p0, len(d_2066_v3_)), d_2065_v2_)
            nw356_[int(24)] = d_2066_v3_
            nw356_[int(25)] = d_2066_v3_
            d_2099_v22_ = nw356_
            index350_ = default__.safeIndex(809, (d_2099_v22_).length(0))
            (d_2099_v22_)[index350_] = d_2066_v3_
            d_2103_v23_: _dafny.Map
            d_2103_v23_ = _dafny.Map({(p0) * ((d_2067_v4_)[default__.safeIndex(676, (d_2067_v4_).length(0))]): (d_2066_v3_) + (((d_2066_v3_).set(default__.safeIndex(832, len(d_2066_v3_)), d_2065_v2_)).set(default__.safeIndex(p2, len((d_2066_v3_).set(default__.safeIndex(832, len(d_2066_v3_)), d_2065_v2_))), d_2065_v2_))})
            index351_ = default__.safeIndex(809, (d_2099_v22_).length(0))
            (d_2099_v22_)[index351_] = ((d_2103_v23_)[(d_2067_v4_)[default__.safeIndex(676, (d_2067_v4_).length(0))]] if ((d_2067_v4_)[default__.safeIndex(676, (d_2067_v4_).length(0))]) in (d_2103_v23_) else d_2066_v3_)
            (globalState).f2 = p1
        elif True:
            d_2104___mcc_h3_ = source26_.cf21
            d_2105_cf21_ = d_2104___mcc_h3_
            if not(p1):
                (globalState).f20 = (d_2066_v3_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "isojrib"))) + (d_2066_v3_))
                (globalState).f16 = (0) - ((0) - (default__.safeModuloInt(p2, -215)))
                d_2106_v24_: _dafny.Map
                d_2106_v24_ = _dafny.Map({p0: p1})
                d_2107_v25_: _dafny.Map
                d_2107_v25_ = _dafny.Map({default__.fm0(p0, p0, (0) - (p2), not(p1), globalState): d_2066_v3_})
                d_2106_v24_ = (d_2106_v24_).set(len(d_2107_v25_), p1)
                d_2108_v26_: _dafny.MultiSet
                d_2108_v26_ = _dafny.MultiSet([p1, p1])
                d_2109_v27_: _dafny.Seq
                d_2109_v27_ = _dafny.SeqWithoutIsStrInference([d_2108_v26_])
                d_2110_v28_: _dafny.Map
                d_2110_v28_ = _dafny.Map({((d_2109_v27_)[default__.safeIndex(p2, len(d_2109_v27_))]).ispropersubset(d_2108_v26_): _dafny.SeqWithoutIsStrInference([p1, False])})
                d_2111_v29_: _dafny.Seq
                d_2111_v29_ = _dafny.SeqWithoutIsStrInference([p1, p1])
                d_2110_v28_ = (d_2110_v28_).set(p1, d_2111_v29_)
                d_2112_v30_: _dafny.Map
                d_2112_v30_ = _dafny.Map({d_2067_v4_: d_2073_v7_})
                (globalState).f16 = default__.fm0(((d_2108_v26_)[p1] if (p1) in (d_2108_v26_) else (d_2067_v4_)[default__.safeIndex(676, (d_2067_v4_).length(0))]), default__.fm0(default__.fm0(p2, p0, (d_2067_v4_)[default__.safeIndex(676, (d_2067_v4_).length(0))], True, globalState), len(d_2106_v24_), len(d_2112_v30_), p1, globalState), default__.fm0((d_2067_v4_)[default__.safeIndex(676, (d_2067_v4_).length(0))], (d_2067_v4_)[default__.safeIndex(676, (d_2067_v4_).length(0))], (d_2067_v4_)[default__.safeIndex(676, (d_2067_v4_).length(0))], p1, globalState), p1, globalState)
            elif True:
                d_2113_v31_: _dafny.MultiSet
                d_2113_v31_ = _dafny.MultiSet([p1, p1])
                d_2114_v32_: _dafny.Map
                d_2114_v32_ = _dafny.Map({d_2067_v4_: (p1) not in (d_2113_v31_)})
                d_2114_v32_ = (d_2114_v32_).set(d_2067_v4_, p1)
                d_2115_v33_: _dafny.Map
                d_2115_v33_ = _dafny.Map({267: p2})
                d_2116_v34_: _dafny.Seq
                d_2116_v34_ = _dafny.SeqWithoutIsStrInference([((d_2115_v33_)[len((d_2115_v33_).set((d_2067_v4_)[default__.safeIndex(676, (d_2067_v4_).length(0))], p0))] if (len((d_2115_v33_).set((d_2067_v4_)[default__.safeIndex(676, (d_2067_v4_).length(0))], p0))) in (d_2115_v33_) else p2), (d_2067_v4_)[default__.safeIndex(676, (d_2067_v4_).length(0))], (d_2067_v4_)[default__.safeIndex(676, (d_2067_v4_).length(0))]])
                d_2117_v35_: _dafny.Map
                d_2117_v35_ = _dafny.Map({(d_2116_v34_)[default__.safeIndex(len(d_2064_v1_), len(d_2116_v34_))]: p1})
                d_2117_v35_ = (d_2117_v35_).set((d_2067_v4_)[default__.safeIndex(676, (d_2067_v4_).length(0))], p1)
                d_2118_v36_: C4
                nw357_ = C4()
                nw357_.ctor__(p1, p1)
                d_2118_v36_ = nw357_
                d_2119_v37_: _dafny.Set
                d_2119_v37_ = _dafny.Set({d_2118_v36_})
                d_2119_v37_ = d_2119_v37_
                (globalState).f16 = (d_2116_v34_)[default__.safeIndex((p0) - (p0), len(d_2116_v34_))]
                d_2120_v38_: _dafny.Array
                nw358_ = _dafny.Array(_dafny.Map({}), 25)
                d_2120_v38_ = nw358_
                d_2121_v40_: _dafny.Map
                d_2121_v40_ = _dafny.Map({d_2065_v2_: False})
                d_2122_v41_: D22
                d_2122_v41_ = D22_DC59(D22_DC57((d_2067_v4_)[default__.safeIndex(676, (d_2067_v4_).length(0))], 767, d_2118_v36_.f25))
                d_2123_v42_: _dafny.Map
                def iife149_():
                    coll77_ = _dafny.Map()
                    compr_77_: str
                    for compr_77_ in (d_2121_v40_).keys.Elements:
                        d_2124_v39_: str = compr_77_
                        if (d_2124_v39_) in (d_2121_v40_):
                            coll77_[d_2124_v39_] = (d_2118_v36_).f26
                    return _dafny.Map(coll77_)
                d_2123_v42_ = _dafny.Map({iife149_()
                : d_2122_v41_})
                index352_ = default__.safeIndex(187, (d_2120_v38_).length(0))
                (d_2120_v38_)[index352_] = (d_2123_v42_) | (d_2123_v42_)
                d_2125_v43_: _dafny.Seq
                d_2125_v43_ = _dafny.SeqWithoutIsStrInference([(d_2118_v36_).f26])
                d_2126_v44_: _dafny.Map
                d_2126_v44_ = _dafny.Map({(d_2118_v36_).f26: d_2125_v43_})
                d_2127_v45_: D16
                d_2127_v45_ = D16_DC36((d_2118_v36_).f26, (d_2113_v31_) | (_dafny.MultiSet(((d_2126_v44_)[(self).fm7(p1, len(d_2125_v43_), d_2118_v36_.f25, globalState)] if ((self).fm7(p1, len(d_2125_v43_), d_2118_v36_.f25, globalState)) in (d_2126_v44_) else d_2125_v43_))), d_2118_v36_.f25)
                d_2128_v46_: _dafny.Array
                nw359_ = _dafny.Array(False, 3)
                d_2128_v46_ = nw359_
                d_2129_v47_: _dafny.Seq
                d_2129_v47_ = _dafny.SeqWithoutIsStrInference([d_2128_v46_])
                d_2130_v48_: _dafny.Array
                nw360_ = _dafny.Array(None, 26)
                nw360_[int(0)] = d_2128_v46_
                nw360_[int(1)] = d_2128_v46_
                nw360_[int(2)] = d_2128_v46_
                nw360_[int(3)] = d_2128_v46_
                nw360_[int(4)] = d_2128_v46_
                nw360_[int(5)] = d_2128_v46_
                nw360_[int(6)] = d_2128_v46_
                nw360_[int(7)] = d_2128_v46_
                nw360_[int(8)] = d_2128_v46_
                nw360_[int(9)] = d_2128_v46_
                nw360_[int(10)] = d_2128_v46_
                nw360_[int(11)] = d_2128_v46_
                nw360_[int(12)] = d_2128_v46_
                nw360_[int(13)] = d_2128_v46_
                nw360_[int(14)] = d_2128_v46_
                nw360_[int(15)] = d_2128_v46_
                nw360_[int(16)] = d_2128_v46_
                nw360_[int(17)] = d_2128_v46_
                nw360_[int(18)] = (d_2129_v47_)[default__.safeIndex(p2, len(d_2129_v47_))]
                nw360_[int(19)] = d_2128_v46_
                nw360_[int(20)] = d_2128_v46_
                nw360_[int(21)] = d_2128_v46_
                nw360_[int(22)] = d_2128_v46_
                nw360_[int(23)] = d_2128_v46_
                nw360_[int(24)] = d_2128_v46_
                nw360_[int(25)] = d_2128_v46_
                d_2130_v48_ = nw360_
                index353_ = default__.safeIndex(558, (d_2130_v48_).length(0))
                (d_2130_v48_)[index353_] = d_2128_v46_
                pat_let_tv42_ = d_2125_v43_
                pat_let_tv43_ = d_2067_v4_
                pat_let_tv44_ = d_2067_v4_
                pat_let_tv45_ = d_2125_v43_
                index354_ = default__.safeIndex(187, (d_2120_v38_).length(0))
                index355_ = default__.safeIndex(558, (d_2130_v48_).length(0))
                def iife150_(_pat_let36_0):
                    def iife151_(d_2131_dt__update__tmp_h2_):
                        def iife152_(_pat_let37_0):
                            def iife153_(d_2132_dt__update_hcf49_h0_):
                                return D16_DC36((d_2131_dt__update__tmp_h2_).cf47, (d_2131_dt__update__tmp_h2_).cf48, d_2132_dt__update_hcf49_h0_)
                            return iife153_(_pat_let37_0)
                        return iife152_((pat_let_tv42_)[default__.safeIndex((pat_let_tv44_)[default__.safeIndex(676, (pat_let_tv43_).length(0))], len(pat_let_tv45_))])
                    return iife151_(_pat_let36_0)
                rhs346_ = (_dafny.Map({d_2121_v40_: d_2122_v41_})) | (d_2123_v42_)
                rhs347_ = d_2125_v43_
                rhs348_ = (d_2066_v3_)[default__.safeIndex((p0) - (p2), len(d_2066_v3_))]
                rhs349_ = iife150_(D16_DC36(not((d_2118_v36_).f26), d_2113_v31_, d_2118_v36_.f25))
                rhs350_ = d_2128_v46_
                lhs319_ = d_2120_v38_
                lhs320_ = default__.safeIndex(187, (d_2120_v38_).length(0))
                lhs321_ = d_2130_v48_
                lhs322_ = default__.safeIndex(558, (d_2130_v48_).length(0))
                lhs319_[lhs320_] = rhs346_
                d_2125_v43_ = rhs347_
                d_2065_v2_ = rhs348_
                d_2127_v45_ = rhs349_
                lhs321_[lhs322_] = rhs350_
            index356_ = default__.safeIndex(676, (d_2067_v4_).length(0))
            def iife154_():
                coll78_ = _dafny.Map()
                compr_78_: _dafny.Seq
                for compr_78_ in (_dafny.SeqWithoutIsStrInference([d_2066_v3_])).Elements:
                    d_2133_v49_: _dafny.Seq = compr_78_
                    if (d_2133_v49_) in (_dafny.SeqWithoutIsStrInference([d_2066_v3_])):
                        coll78_[d_2133_v49_] = d_2065_v2_
                return _dafny.Map(coll78_)
            (d_2067_v4_)[index356_] = default__.safeModuloInt((d_2067_v4_)[default__.safeIndex(676, (d_2067_v4_).length(0))], default__.safeModuloInt(len(iife154_()
            ), p2))
            d_2134_v50_: _dafny.MultiSet
            d_2134_v50_ = _dafny.MultiSet([d_2067_v4_, d_2067_v4_])
            index357_ = default__.safeIndex(676, (d_2067_v4_).length(0))
            (d_2067_v4_)[index357_] = (0) - (((d_2067_v4_)[default__.safeIndex(676, (d_2067_v4_).length(0))]) + (((d_2134_v50_) - (d_2134_v50_)).cardinality))
            if (d_2073_v7_).fm7(not (p1) or (not(not(p1))), p0, p1, globalState):
                d_2135_v51_: T0
                nw361_ = C3()
                nw361_.ctor__()
                d_2135_v51_ = nw361_
                d_2135_v51_ = d_2135_v51_
                index358_ = default__.safeIndex(676, (d_2067_v4_).length(0))
                (d_2067_v4_)[index358_] = p0
                d_2136_v52_: _dafny.Set
                d_2136_v52_ = _dafny.Set({p2, p2})
                d_2137_v53_: _dafny.Map
                d_2137_v53_ = _dafny.Map({p1: (0) - ((d_2067_v4_)[default__.safeIndex(676, (d_2067_v4_).length(0))])})
                d_2138_v54_: _dafny.Map
                d_2138_v54_ = _dafny.Map({p2: len(d_2066_v3_)})
                d_2139_v55_: _dafny.Seq
                d_2139_v55_ = _dafny.SeqWithoutIsStrInference([(d_2067_v4_)[default__.safeIndex(676, (d_2067_v4_).length(0))], len(d_2138_v54_)])
                d_2140_v56_: _dafny.Map
                d_2140_v56_ = _dafny.Map({((d_2137_v53_)[p1] if (p1) in (d_2137_v53_) else (d_2139_v55_)[default__.safeIndex(len(d_2066_v3_), len(d_2139_v55_))]): not(p1)})
                d_2136_v52_ = default__.fm26(d_2065_v2_, d_2140_v56_, globalState)
                index359_ = default__.safeIndex(676, (d_2067_v4_).length(0))
                (d_2067_v4_)[index359_] = ((default__.fm61(748, d_2065_v2_, d_2065_v2_, globalState)).set(((d_2063_v0_).set(True, p1)) | ((_dafny.Map({p1: p1})).set(True, p1)), default__.abs(p0))).cardinality
                d_2141_v57_: _dafny.MultiSet
                d_2141_v57_ = _dafny.MultiSet([p0])
                d_2142_v58_: _dafny.Map
                d_2142_v58_ = _dafny.Map({p1: d_2063_v0_})
                d_2143_v59_: D23
                d_2143_v59_ = D23_DC60(d_2142_v58_)
                d_2144_v60_: _dafny.Set
                d_2144_v60_ = _dafny.Set({d_2143_v59_, d_2143_v59_, d_2143_v59_})
                d_2145_v61_: C2
                nw362_ = C2()
                nw362_.ctor__(len((d_2144_v60_) | (d_2144_v60_)))
                d_2145_v61_ = nw362_
                rhs351_ = (d_2141_v57_) - ((d_2141_v57_) | (d_2141_v57_))
                rhs352_ = not(not(p1))
                rhs353_ = (d_2141_v57_) - ((D25_DC67(d_2141_v57_)).cf98)
                rhs354_ = d_2145_v61_
                lhs323_ = globalState
                d_2141_v57_ = rhs351_
                lhs323_.f21 = rhs352_
                d_2141_v57_ = rhs353_
                d_2145_v61_ = rhs354_
            elif True:
                d_2065_v2_ = d_2065_v2_
                d_2063_v0_ = (d_2063_v0_).set(p1, p1)
                d_2146_v62_: _dafny.Array
                nw363_ = _dafny.Array(_dafny.Array(None, 0), 21)
                d_2146_v62_ = nw363_
                index360_ = default__.safeIndex(510, (d_2146_v62_).length(0))
                (d_2146_v62_)[index360_] = d_2067_v4_
                index361_ = default__.safeIndex(510, (d_2146_v62_).length(0))
                (d_2146_v62_)[index361_] = d_2067_v4_
                (globalState).f13 = p1
                d_2147_v63_: _dafny.Array
                nw364_ = _dafny.Array(_dafny.Map({}), 29)
                d_2147_v63_ = nw364_
                d_2147_v63_ = d_2147_v63_
        r0 = ((d_2066_v3_ if p1 else (d_2066_v3_).set(default__.safeIndex((d_2067_v4_)[default__.safeIndex(676, (d_2067_v4_).length(0))], len(d_2066_v3_)), _dafny.CodePoint('n')))) + (d_2066_v3_)
        return r0

    @property
    def f23(self):
        return self._f23
