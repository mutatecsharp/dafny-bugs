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
        return (not(True)) or ((not(True)) or (False))

    @staticmethod
    def fm5(p0, p1, globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(955, 500):
                d_1_v0_: int = compr_0_
                if ((955) <= (d_1_v0_)) and ((d_1_v0_) < (500)):
                    coll0_ = coll0_.union(_dafny.Set([(d_1_v0_) + (len(_dafny.Set({True, not(True)})))]))
            return _dafny.Set(coll0_)
        def iife1_():
            coll1_ = _dafny.Set()
            compr_1_: int
            for compr_1_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: True}))])).Elements:
                d_2_v1_: int = compr_1_
                if (d_2_v1_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: True}))])):
                    coll1_ = coll1_.union(_dafny.Set([(d_2_v1_) - (601)]))
            return _dafny.Set(coll1_)
        return len((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.Set({299, 553}), _dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_0_i0_ in range(default__.abs(162))]))}), iife0_()
        , iife1_()
        ])), 528])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_3_i1_ in range(default__.abs(429))])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "smswnlmsn")))])))

    @staticmethod
    def fm7(globalState):
        if False:
            return (_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([True]))
        elif True:
            return _dafny.SeqWithoutIsStrInference([True])

    @staticmethod
    def fm8(p0, p1, p2, globalState):
        return ((_dafny.MultiSet([not(True)])).intersection(_dafny.MultiSet([False]))).intersection((_dafny.MultiSet([False])) | (_dafny.MultiSet([not(True), False])))

    @staticmethod
    def fm9(p0, p1, globalState):
        def iife2_():
            coll2_ = _dafny.Set()
            compr_2_: _dafny.Map
            for compr_2_ in (_dafny.Set({_dafny.Map({713: True})})).Elements:
                d_4_v0_: _dafny.Map = compr_2_
                if (d_4_v0_) in (_dafny.Set({_dafny.Map({713: True})})):
                    coll2_ = coll2_.union(_dafny.Set([d_4_v0_]))
            return _dafny.Set(coll2_)
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(-871, -831):
                d_5_v1_: int = compr_3_
                if ((-871) <= (d_5_v1_)) and ((d_5_v1_) < (-831)):
                    coll3_[(d_5_v1_) + (-737)] = False
            return _dafny.Map(coll3_)
        return ((iife2_()
        ).intersection(_dafny.Set({_dafny.Map({(_dafny.MultiSet([False])).cardinality: False})}))) - (_dafny.Set({_dafny.Map({413: True}), _dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")))): False}), iife3_()
        }))

    @staticmethod
    def fm10(p0, p1, p2, p3, globalState):
        return _dafny.CodePoint('n')

    @staticmethod
    def fm11(p0, p1, globalState):
        def iife4_():
            coll4_ = _dafny.Set()
            compr_4_: int
            for compr_4_ in _dafny.IntegerRange(71, 108):
                d_6_v0_: int = compr_4_
                if ((71) <= (d_6_v0_)) and ((d_6_v0_) < (108)):
                    coll4_ = coll4_.union(_dafny.Set([(d_6_v0_) + (len(_dafny.Map({False: not(not(True))})))]))
            return _dafny.Set(coll4_)
        def iife5_():
            coll5_ = _dafny.Set()
            compr_5_: int
            for compr_5_ in _dafny.IntegerRange(388, 551):
                d_7_v1_: int = compr_5_
                if ((388) <= (d_7_v1_)) and ((d_7_v1_) < (551)):
                    def iife6_():
                        coll6_ = _dafny.Map()
                        compr_6_: str
                        for compr_6_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g'), _dafny.CodePoint('e')])).Elements:
                            d_8_v2_: str = compr_6_
                            if (d_8_v2_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g'), _dafny.CodePoint('e')])):
                                coll6_[d_8_v2_] = len(_dafny.Map({True: _dafny.CodePoint('h')}))
                        return _dafny.Map(coll6_)
                    coll5_ = coll5_.union(_dafny.Set([default__.safeDivisionInt(d_7_v1_, len(iife6_()
))]))
            return _dafny.Set(coll5_)
        return _dafny.Map({255: len((_dafny.Set({D12_DC26(_dafny.Set({387})), D12_DC26(_dafny.Set({len(_dafny.Map({569: True}))})), D12_DC26(_dafny.Set({427, 687, 659, 249})), D12_DC26(_dafny.Set({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.Set({True})])): False}))}))}) if not(not(True)) else _dafny.Set({D12_DC26(iife4_()
), D12_DC26(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cao"))), 353})), D12_DC26(iife5_()
), D12_DC26(_dafny.Set({129}))})))})

    @staticmethod
    def fm12(p0, p1, p2, globalState):
        return ((_dafny.Set({True})).intersection(_dafny.Set({True}))) | ((_dafny.Set({False, True, not(True)})) - (_dafny.Set({not(not(True))})))

    @staticmethod
    def fm13(p0, globalState):
        def iife7_():
            coll7_ = _dafny.Map()
            compr_7_: _dafny.Seq
            for compr_7_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tmyf")): not(False)})).keys.Elements:
                d_10_v0_: _dafny.Seq = compr_7_
                if (d_10_v0_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tmyf")): not(False)})):
                    coll7_[d_10_v0_] = (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ecjrmps"))))
            return _dafny.Map(coll7_)
        return ((_dafny.Map({False: D3_DC6((_dafny.MultiSet([not(True)])).cardinality, -906, False)})) | (_dafny.Map({False: D3_DC6(799, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wmk"))) for d_9_i0_ in range(default__.abs(-306))])), True)}))) | (_dafny.Map({False: D3_DC6(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qi"))), len(iife7_()
), True)}))

    @staticmethod
    def fm14(p0, globalState):
        source0_ = D6_DC12(True, 932)
        if source0_.is_DC12:
            d_11___mcc_h0_ = source0_.cf17
            d_12___mcc_h1_ = source0_.cf18
            d_13_cf18_ = d_12___mcc_h1_
            d_14_cf17_ = d_11___mcc_h0_
            return D10_DC18(_dafny.Set({d_14_cf17_}))
        elif True:
            d_15___mcc_h2_ = source0_.cf16
            d_16_cf16_ = d_15___mcc_h2_
            return D10_DC18(_dafny.Set({not(True)}))

    @staticmethod
    def fm15(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_17_i0_ in range(default__.abs(10))])) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_18_i1_ in range(default__.abs(-385))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_19_i2_ in range(default__.abs(-51))])))

    @staticmethod
    def fm16(p0, globalState):
        return D10_DC20((_dafny.Set({266})).isdisjoint(_dafny.Set({246})), (-107) + (449))

    @staticmethod
    def fm17(p0, p1, p2, globalState):
        source1_ = (D2_DC2(False) if True else D2_DC2(True))
        if source1_.is_DC3:
            d_20___mcc_h0_ = source1_.cf3
            d_21___mcc_h1_ = source1_.cf4
            d_22_cf4_ = d_21___mcc_h1_
            d_23_cf3_ = d_20___mcc_h0_
            return d_23_cf3_
        elif True:
            d_24___mcc_h2_ = source1_.cf2
            d_25_cf2_ = d_24___mcc_h2_
            return (0) - ((D11_DC24(d_25_cf2_, len(_dafny.Map({d_25_cf2_: 205})))).cf37)

    @staticmethod
    def fm19(p0, globalState):
        return (_dafny.Map({True: not(not(True))})) | ((_dafny.Map({True: True}) if True else _dafny.Map({(D11_DC23(355, not(True))).cf35: not(False)})))

    @staticmethod
    def fm20(p0, p1, globalState):
        source2_ = (D11_DC24(True, 961) if not(True) else D11_DC24(True, len(_dafny.Map({(_dafny.MultiSet([_dafny.CodePoint('a'), _dafny.CodePoint('l'), _dafny.CodePoint('a')])).cardinality: 634}))))
        if source2_.is_DC23:
            d_26___mcc_h0_ = source2_.cf34
            d_27___mcc_h1_ = source2_.cf35
            d_28_cf35_ = d_27___mcc_h1_
            d_29_cf34_ = d_26___mcc_h0_
            return D10_DC19(_dafny.CodePoint('r'))
        elif source2_.is_DC24:
            d_30___mcc_h2_ = source2_.cf36
            d_31___mcc_h3_ = source2_.cf37
            d_32_cf37_ = d_31___mcc_h3_
            d_33_cf36_ = d_30___mcc_h2_
            return D10_DC19(_dafny.CodePoint('x'))
        elif source2_.is_DC22:
            d_34___mcc_h4_ = source2_.cf33
            d_35_cf33_ = d_34___mcc_h4_
            return D10_DC19(_dafny.CodePoint('k'))
        elif True:
            d_36___mcc_h5_ = source2_.cf38
            d_37_cf38_ = d_36___mcc_h5_
            return D10_DC19(_dafny.CodePoint('x'))

    @staticmethod
    def fm21(p0, p1, globalState):
        source3_ = (D9_DC16(_dafny.Map({True: D3_DC6(len(_dafny.Map({46: 327})), 723, not(False))})) if True else D9_DC16(_dafny.Map({False: D3_DC6(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nmqjwa"))), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_38_i0_ in range(default__.abs(210))])), False)})))
        if source3_.is_DC17:
            d_39___mcc_h0_ = source3_.cf25
            d_40_cf25_ = d_39___mcc_h0_
            def iife8_():
                coll8_ = _dafny.Map()
                compr_8_: int
                for compr_8_ in (_dafny.SeqWithoutIsStrInference([d_40_cf25_])).Elements:
                    d_41_v0_: int = compr_8_
                    if (d_41_v0_) in (_dafny.SeqWithoutIsStrInference([d_40_cf25_])):
                        coll8_[(d_41_v0_) - (d_40_cf25_)] = False
                return _dafny.Map(coll8_)
            return (_dafny.SeqWithoutIsStrInference([D7_DC13(iife8_()
)])) + (_dafny.SeqWithoutIsStrInference([D7_DC13(_dafny.Map({d_40_cf25_: not(False)}))]))
        elif True:
            d_42___mcc_h1_ = source3_.cf24
            d_43_cf24_ = d_42___mcc_h1_
            return _dafny.SeqWithoutIsStrInference([D7_DC13(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_44_i1_ in range(default__.abs(-943))])): False})), D7_DC13(_dafny.Map({853: True}))])

    @staticmethod
    def fm22(p0, p1, p2, globalState):
        return D11_DC23((0) - (len(_dafny.Map({_dafny.SeqWithoutIsStrInference([True, not(False), True]): 132}))), (len(_dafny.Set({249}))) != (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")))))

    @staticmethod
    def fm23(globalState):
        return D3_DC6((8) + (len(_dafny.SeqWithoutIsStrInference([True]))), len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j'), _dafny.CodePoint('k'), _dafny.CodePoint('x')]) if False else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l')]))), False)

    @staticmethod
    def fm24(p0, p1, p2, globalState):
        if (_dafny.Map({False: (_dafny.MultiSet([290, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_45_i0_ in range(default__.abs(977))])), -498, (_dafny.MultiSet([752])).cardinality])).cardinality})) != (_dafny.Map({False: 618})):
            return D9_DC16(_dafny.Map({False: D3_DC6((_dafny.MultiSet([False])).cardinality, 630, True)}))
        elif True:
            return D9_DC16(_dafny.Map({False: D3_DC6(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cnempwcsy"))), len(_dafny.Set({True})), False)}))

    @staticmethod
    def fm25(p0, p1, p2, globalState):
        def iife9_():
            def iife11_():
                coll11_ = _dafny.Set()
                compr_11_: str
                for compr_11_ in (_dafny.Map({_dafny.CodePoint('i'): (_dafny.MultiSet([633])).cardinality})).keys.Elements:
                    d_48_v1_: str = compr_11_
                    if (d_48_v1_) in (_dafny.Map({_dafny.CodePoint('i'): (_dafny.MultiSet([633])).cardinality})):
                        coll11_ = coll11_.union(_dafny.Set([d_48_v1_]))
                return _dafny.Set(coll11_)
            coll9_ = _dafny.Map()
            def iife10_():
                coll10_ = _dafny.Set()
                compr_10_: str
                for compr_10_ in (_dafny.Map({_dafny.CodePoint('i'): (_dafny.MultiSet([633])).cardinality})).keys.Elements:
                    d_46_v1_: str = compr_10_
                    if (d_46_v1_) in (_dafny.Map({_dafny.CodePoint('i'): (_dafny.MultiSet([633])).cardinality})):
                        coll10_ = coll10_.union(_dafny.Set([d_46_v1_]))
                return _dafny.Set(coll10_)
            compr_9_: str
            for compr_9_ in (iife10_()
            ).Elements:
                d_47_v0_: str = compr_9_
                if (d_47_v0_) in (iife11_()
                ):
                    coll9_[d_47_v0_] = False
            return _dafny.Map(coll9_)
        return ((iife9_()
        ) | (_dafny.Map({_dafny.CodePoint('h'): True}))) | ((_dafny.Map({_dafny.CodePoint('n'): True})) | (_dafny.Map({_dafny.CodePoint('p'): False})))

    @staticmethod
    def fm26(globalState):
        return (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({639: False}))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({not(True)})) for d_49_i0_ in range(default__.abs(680))]))

    @staticmethod
    def fm27(p0, globalState):
        return (_dafny.Map({_dafny.MultiSet([True]): (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.Set({len(_dafny.Set({False}))})])))})) | (_dafny.Map({_dafny.MultiSet([False, False, False, False]): -332}))

    @staticmethod
    def fm28(p0, p1, p2, globalState):
        def iife12_():
            def iife14_():
                coll14_ = _dafny.Set()
                compr_14_: _dafny.Seq
                for compr_14_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ujowcoxtg")) for d_50_i0_ in range(default__.abs(152))])).Elements:
                    d_53_v1_: _dafny.Seq = compr_14_
                    if (d_53_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ujowcoxtg")) for d_50_i0_ in range(default__.abs(152))])):
                        coll14_ = coll14_.union(_dafny.Set([d_53_v1_]))
                return _dafny.Set(coll14_)
            coll12_ = _dafny.Map()
            def iife13_():
                coll13_ = _dafny.Set()
                compr_13_: _dafny.Seq
                for compr_13_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ujowcoxtg")) for d_50_i0_ in range(default__.abs(152))])).Elements:
                    d_51_v1_: _dafny.Seq = compr_13_
                    if (d_51_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ujowcoxtg")) for d_50_i0_ in range(default__.abs(152))])):
                        coll13_ = coll13_.union(_dafny.Set([d_51_v1_]))
                return _dafny.Set(coll13_)
            compr_12_: D9
            for compr_12_ in (_dafny.Set({D9_DC17(859), D9_DC17(413), D9_DC17(len(iife13_()
))})).Elements:
                d_52_v0_: D9 = compr_12_
                if (d_52_v0_) in (_dafny.Set({D9_DC17(859), D9_DC17(413), D9_DC17(len(iife14_()
))})):
                    coll12_[d_52_v0_] = not(True)
            return _dafny.Map(coll12_)
        return (_dafny.Map({D9_DC17(len(_dafny.Set({False, True}))): False})) | ((_dafny.Map({D9_DC17(288): not(True)})) | (iife12_()
        ))

    @staticmethod
    def fm29(p0, globalState):
        if False:
            return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_54_i0_ in range(default__.abs(857))])
        elif True:
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qt"))

    @staticmethod
    def fm30(p0, p1, p2, p3, globalState):
        def iife15_():
            def iife17_():
                coll17_ = _dafny.Set()
                compr_17_: _dafny.Map
                for compr_17_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({not(True): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ro"))}) for d_55_i0_ in range(default__.abs(590))])).Elements:
                    d_58_v1_: _dafny.Map = compr_17_
                    if (d_58_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({not(True): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ro"))}) for d_55_i0_ in range(default__.abs(590))])):
                        coll17_ = coll17_.union(_dafny.Set([d_58_v1_]))
                return _dafny.Set(coll17_)
            coll15_ = _dafny.Map()
            def iife16_():
                coll16_ = _dafny.Set()
                compr_16_: _dafny.Map
                for compr_16_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({not(True): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ro"))}) for d_55_i0_ in range(default__.abs(590))])).Elements:
                    d_56_v1_: _dafny.Map = compr_16_
                    if (d_56_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({not(True): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ro"))}) for d_55_i0_ in range(default__.abs(590))])):
                        coll16_ = coll16_.union(_dafny.Set([d_56_v1_]))
                return _dafny.Set(coll16_)
            compr_15_: _dafny.Map
            for compr_15_ in ((iife16_()
            ).intersection(_dafny.Set({_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qduxisjho"))})}))).Elements:
                d_57_v0_: _dafny.Map = compr_15_
                if (d_57_v0_) in ((iife17_()
                ).intersection(_dafny.Set({_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qduxisjho"))})}))):
                    coll15_[d_57_v0_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fyriywjbp"))) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mniamr")))
            return _dafny.Map(coll15_)
        return iife15_()
        

    @staticmethod
    def fm31(p0, p1, p2, p3, globalState):
        return (_dafny.Map({True: -432})) | ((_dafny.Map({True: 152})) | (_dafny.Map({False: (0) - ((_dafny.MultiSet([True, not(False)])).cardinality)})))

    @staticmethod
    def fm32(p0, p1, globalState):
        def iife18_():
            coll18_ = _dafny.Map()
            compr_18_: int
            for compr_18_ in _dafny.IntegerRange(672, 958):
                d_59_v0_: int = compr_18_
                if ((672) <= (d_59_v0_)) and ((d_59_v0_) < (958)):
                    coll18_[default__.safeModuloInt(d_59_v0_, len(_dafny.SeqWithoutIsStrInference([False])))] = _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([not(True)])) + (_dafny.SeqWithoutIsStrInference([True, False])))
            return _dafny.Map(coll18_)
        return iife18_()
        

    @staticmethod
    def fm33(globalState):
        def iife19_():
            coll19_ = _dafny.Map()
            compr_19_: int
            for compr_19_ in _dafny.IntegerRange(370, -582):
                d_60_v0_: int = compr_19_
                if ((370) <= (d_60_v0_)) and ((d_60_v0_) < (-582)):
                    coll19_[default__.safeDivisionInt(d_60_v0_, -876)] = False
            return _dafny.Map(coll19_)
        def iife20_():
            coll20_ = _dafny.Map()
            compr_20_: _dafny.Seq
            for compr_20_ in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xtjcm")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_61_i0_ in range(default__.abs(607))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_62_i1_ in range(default__.abs(662))])])).Elements:
                d_63_v1_: _dafny.Seq = compr_20_
                if (d_63_v1_) in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xtjcm")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_61_i0_ in range(default__.abs(607))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_62_i1_ in range(default__.abs(662))])])):
                    coll20_[d_63_v1_] = len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([473 for d_64_i2_ in range(default__.abs(656))])): True}))
            return _dafny.Map(coll20_)
        return ((_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gvtucu")): len(iife19_()
        )})) | ((D16_DC33(iife20_()
)).cf49)) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o")): 612}))

    @staticmethod
    def fm34(p0, p1, p2, globalState):
        def iife21_():
            coll21_ = _dafny.Set()
            compr_21_: int
            for compr_21_ in _dafny.IntegerRange(918, 553):
                d_65_v0_: int = compr_21_
                if ((918) <= (d_65_v0_)) and ((d_65_v0_) < (553)):
                    coll21_ = coll21_.union(_dafny.Set([(d_65_v0_) + (-644)]))
            return _dafny.Set(coll21_)
        def iife22_():
            def iife24_():
                coll24_ = _dafny.Set()
                compr_24_: _dafny.Seq
                for compr_24_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference([601])})).Elements:
                    d_68_v1_: _dafny.Seq = compr_24_
                    if (d_68_v1_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference([601])})):
                        coll24_ = coll24_.union(_dafny.Set([d_68_v1_]))
                return _dafny.Set(coll24_)
            coll22_ = _dafny.Set()
            def iife23_():
                coll23_ = _dafny.Set()
                compr_23_: _dafny.Seq
                for compr_23_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference([601])})).Elements:
                    d_66_v1_: _dafny.Seq = compr_23_
                    if (d_66_v1_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference([601])})):
                        coll23_ = coll23_.union(_dafny.Set([d_66_v1_]))
                return _dafny.Set(coll23_)
            compr_22_: _dafny.Seq
            for compr_22_ in (iife23_()
            ).Elements:
                d_67_v2_: _dafny.Seq = compr_22_
                if (d_67_v2_) in (iife24_()
                ):
                    coll22_ = coll22_.union(_dafny.Set([d_67_v2_]))
            return _dafny.Set(coll22_)
        return ((_dafny.Set({_dafny.SeqWithoutIsStrInference([-661]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qhld"))), (0) - (len(_dafny.Set({False, True})))])})).intersection(_dafny.Set({_dafny.SeqWithoutIsStrInference([len(iife21_()
        )])}))) | ((iife22_()
        ) - (_dafny.Set({_dafny.SeqWithoutIsStrInference([894]), _dafny.SeqWithoutIsStrInference([len((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False]))])))]), _dafny.SeqWithoutIsStrInference([-476, 198]), _dafny.SeqWithoutIsStrInference([439])})))

    @staticmethod
    def m6(p0, p1, p2, p3, globalState):
        r0: str = _dafny.CodePoint('D')
        hi0_ = (0) - (p0)
        for d_69_i0_ in range(p0, hi0_):
            if p1:
                d_70_v0_: _dafny.Array
                nw0_ = _dafny.Array(False, 16)
                d_70_v0_ = nw0_
                index0_ = default__.safeIndex(394, (d_70_v0_).length(0))
                (d_70_v0_)[index0_] = p2
                index1_ = default__.safeIndex(394, (d_70_v0_).length(0))
                (d_70_v0_)[index1_] = p3
                d_71_v1_: str
                d_71_v1_ = _dafny.CodePoint('f')
                d_72_v2_: C3
                nw1_ = C3()
                nw1_.ctor__(default__.fm0(d_69_i0_, (d_70_v0_)[default__.safeIndex(394, (d_70_v0_).length(0))], p3, p0, globalState), d_71_v1_, p0)
                d_72_v2_ = nw1_
                d_72_v2_ = d_72_v2_
                d_73_v3_: _dafny.MultiSet
                d_73_v3_ = _dafny.MultiSet([default__.fm0(p0, (d_70_v0_)[default__.safeIndex(394, (d_70_v0_).length(0))], d_72_v2_.f31, -620, globalState), p3, default__.fm0(p0, (d_70_v0_)[default__.safeIndex(394, (d_70_v0_).length(0))], d_72_v2_.f31, d_69_i0_, globalState), p1])
                d_74_v4_: _dafny.Map
                d_74_v4_ = _dafny.Map({d_73_v3_: d_69_i0_})
                d_75_v5_: _dafny.Seq
                d_75_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vah"))
                d_76_v6_: C4
                nw2_ = C4()
                nw2_.ctor__(d_74_v4_, (d_75_v5_) < (d_75_v5_), d_71_v1_, p0)
                d_76_v6_ = nw2_
                d_77_v7_: T1
                nw3_ = C3()
                nw3_.ctor__(p2, d_71_v1_, p0)
                d_77_v7_ = nw3_
                d_78_v8_: _dafny.Seq
                d_78_v8_ = _dafny.SeqWithoutIsStrInference([d_77_v7_, d_77_v7_])
                d_79_v9_: C5
                nw4_ = C5()
                nw4_.ctor__((d_76_v6_).f30, d_71_v1_, (d_69_i0_) * (len(d_78_v8_)))
                d_79_v9_ = nw4_
                d_79_v9_ = d_79_v9_
            elif True:
                r0 = default__.fm10((d_69_i0_) >= (d_69_i0_), p2, p1, p3, globalState)
                d_80_v10_: _dafny.Seq
                d_80_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rthyuxu"))
                d_81_v11_: _dafny.MultiSet
                d_81_v11_ = _dafny.MultiSet([default__.fm5(len(d_80_v10_), d_69_i0_, globalState)])
                d_82_v12_: _dafny.Array
                nw5_ = _dafny.Array(None, 8)
                nw5_[int(0)] = p2
                nw5_[int(1)] = (p3 if p2 else p3)
                nw5_[int(2)] = (d_81_v11_).issubset(d_81_v11_)
                nw5_[int(3)] = p2
                nw5_[int(4)] = True
                nw5_[int(5)] = p1
                nw5_[int(6)] = (p0) != ((0) - ((0) - (d_69_i0_)))
                nw5_[int(7)] = (p3) or (p3)
                d_82_v12_ = nw5_
                index2_ = default__.safeIndex(676, (d_82_v12_).length(0))
                (d_82_v12_)[index2_] = p2
                index3_ = default__.safeIndex(676, (d_82_v12_).length(0))
                (d_82_v12_)[index3_] = p3
                index4_ = default__.safeIndex(676, (d_82_v12_).length(0))
                (d_82_v12_)[index4_] = p2
                d_83_v13_: _dafny.Map
                d_83_v13_ = _dafny.Map({-334: p1})
                d_83_v13_ = (d_83_v13_).set(d_69_i0_, (d_82_v12_)[default__.safeIndex(676, (d_82_v12_).length(0))])
                d_84_v14_: _dafny.Array
                def lambda0_(d_85_v12_, d_86_p1_, d_87_p2_):
                    def lambda1_(d_88_i1_):
                        return _dafny.SeqWithoutIsStrInference([not((d_85_v12_)[default__.safeIndex(676, (d_85_v12_).length(0))]), d_86_p1_, d_87_p2_, not(d_86_p1_)])

                    return lambda1_

                init0_ = lambda0_(d_82_v12_, p1, p2)
                nw6_ = _dafny.Array(None, 1)
                for i0_0_ in range(nw6_.length(0)):
                    nw6_[i0_0_] = init0_(i0_0_)
                d_84_v14_ = nw6_
                d_89_v15_: _dafny.Seq
                d_89_v15_ = _dafny.SeqWithoutIsStrInference([p3])
                index5_ = default__.safeIndex(590, (d_84_v14_).length(0))
                (d_84_v14_)[index5_] = d_89_v15_
                d_90_v16_: _dafny.Set
                d_90_v16_ = _dafny.Set({d_83_v13_})
                d_91_v17_: D3
                d_91_v17_ = D3_DC4(d_90_v16_)
                d_92_v18_: C0
                nw7_ = C0()
                nw7_.ctor__()
                d_92_v18_ = nw7_
                d_93_v19_: str
                d_93_v19_ = _dafny.CodePoint('w')
                d_94_v21_: _dafny.Set
                d_94_v21_ = _dafny.Set({d_93_v19_})
                d_95_v22_: D10
                def iife25_():
                    coll25_ = _dafny.Set()
                    compr_25_: str
                    for compr_25_ in (_dafny.Map({d_93_v19_: 736})).keys.Elements:
                        d_96_v20_: str = compr_25_
                        if (d_96_v20_) in (_dafny.Map({d_93_v19_: 736})):
                            coll25_ = coll25_.union(_dafny.Set([d_96_v20_]))
                    return _dafny.Set(coll25_)
                d_95_v22_ = D10_DC21(787, (d_69_i0_) == (296), (iife25_()
) != (d_94_v21_))
                index6_ = default__.safeIndex(590, (d_84_v14_).length(0))
                index7_ = default__.safeIndex(676, (d_82_v12_).length(0))
                rhs0_ = d_89_v15_
                rhs1_ = ((p0) < (d_69_i0_)) or ((p0) >= (len(d_80_v10_)))
                rhs2_ = d_91_v17_
                rhs3_ = d_92_v18_
                rhs4_ = d_95_v22_
                lhs0_ = d_84_v14_
                lhs1_ = default__.safeIndex(590, (d_84_v14_).length(0))
                lhs2_ = d_82_v12_
                lhs3_ = default__.safeIndex(676, (d_82_v12_).length(0))
                lhs0_[lhs1_] = rhs0_
                lhs2_[lhs3_] = rhs1_
                d_91_v17_ = rhs2_
                d_92_v18_ = rhs3_
                d_95_v22_ = rhs4_
            d_97_v23_: _dafny.Seq
            d_97_v23_ = _dafny.SeqWithoutIsStrInference([not(p1), p3])
            d_98_v24_: _dafny.Array
            nw8_ = _dafny.Array(None, 17)
            nw8_[int(0)] = ((d_97_v23_).set(default__.safeIndex(p0, len(d_97_v23_)), p1)) + (d_97_v23_)
            nw8_[int(1)] = (d_97_v23_) + (d_97_v23_)
            nw8_[int(2)] = d_97_v23_
            nw8_[int(3)] = (d_97_v23_) + (d_97_v23_)
            nw8_[int(4)] = (d_97_v23_ if p2 else d_97_v23_)
            nw8_[int(5)] = d_97_v23_
            nw8_[int(6)] = d_97_v23_
            nw8_[int(7)] = d_97_v23_
            nw8_[int(8)] = d_97_v23_
            nw8_[int(9)] = ((d_97_v23_).set(default__.safeIndex(p0, len(d_97_v23_)), not(p1))) + (d_97_v23_)
            nw8_[int(10)] = ((_dafny.SeqWithoutIsStrInference([p3])) + (d_97_v23_)).set(default__.safeIndex(p0, len((_dafny.SeqWithoutIsStrInference([p3])) + (d_97_v23_))), p1)
            nw8_[int(11)] = (d_97_v23_).set(default__.safeIndex(d_69_i0_, len(d_97_v23_)), p1)
            nw8_[int(12)] = d_97_v23_
            nw8_[int(13)] = d_97_v23_
            nw8_[int(14)] = d_97_v23_
            nw8_[int(15)] = d_97_v23_
            nw8_[int(16)] = d_97_v23_
            d_98_v24_ = nw8_
            index8_ = default__.safeIndex(814, (d_98_v24_).length(0))
            (d_98_v24_)[index8_] = d_97_v23_
            index9_ = default__.safeIndex(814, (d_98_v24_).length(0))
            (d_98_v24_)[index9_] = default__.fm7(globalState)
            d_99_v25_: D11
            d_99_v25_ = D11_DC23(p0, p1)
            d_100_v26_: D11
            d_100_v26_ = D11_DC25(d_99_v25_)
            d_101_v27_: _dafny.Map
            d_101_v27_ = _dafny.Map({d_100_v26_: p1})
            rhs5_ = (0) - (default__.fm17((d_101_v27_) == (d_101_v27_), not (p1) or (p1), p3, globalState))
            rhs6_ = p3
            lhs4_ = globalState
            lhs5_ = globalState
            lhs4_.f2 = rhs5_
            lhs5_.f18 = rhs6_
            d_102_v28_: _dafny.Seq
            d_102_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rkvodkl"))
            if (d_102_v28_) != (d_102_v28_):
                d_103_v29_: C2
                nw9_ = C2()
                nw9_.ctor__((d_69_i0_) + (d_69_i0_), not (p2) or (p2))
                d_103_v29_ = nw9_
                (globalState).f18 = True
                d_104_v31_: str
                d_104_v31_ = _dafny.CodePoint('l')
                d_105_v32_: C1
                nw10_ = C1()
                nw10_.ctor__(d_104_v31_, p0)
                d_105_v32_ = nw10_
                d_106_v33_: _dafny.Map
                d_106_v33_ = _dafny.Map({330: d_105_v32_})
                d_107_v34_: _dafny.Map
                d_107_v34_ = _dafny.Map({p0: (d_103_v29_).f28})
                d_108_v35_: C6
                nw11_ = C6()
                def iife26_():
                    coll26_ = _dafny.Map()
                    compr_26_: int
                    for compr_26_ in _dafny.IntegerRange(606, 984):
                        d_109_v30_: int = compr_26_
                        if ((606) <= (d_109_v30_)) and ((d_109_v30_) < (984)):
                            coll26_[(d_109_v30_) * ((d_103_v29_).f27)] = not(False)
                    return _dafny.Map(coll26_)
                nw11_.ctor__(((iife26_()
                ).set(p0, (d_97_v23_)[default__.safeIndex(len(d_106_v33_), len(d_97_v23_))])) | (d_107_v34_))
                d_108_v35_ = nw11_
                (globalState).f3 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rohb")))
                (globalState).f16 = (0) - ((default__.safeDivisionInt((d_103_v29_).f27, d_69_i0_)) - ((p0 if p3 else (0) - (p0))))
            elif True:
                d_110_v36_: _dafny.MultiSet
                d_110_v36_ = _dafny.MultiSet([p3])
                d_111_v37_: _dafny.Map
                d_111_v37_ = _dafny.Map({p0: d_110_v36_})
                (globalState).f18 = (((d_111_v37_)[-178] if (-178) in (d_111_v37_) else _dafny.MultiSet([False]))).isdisjoint(d_110_v36_)
                d_112_v38_: _dafny.Seq
                d_112_v38_ = _dafny.SeqWithoutIsStrInference([((d_98_v24_)[default__.safeIndex(814, (d_98_v24_).length(0))]) + (d_97_v23_), _dafny.SeqWithoutIsStrInference([not(p2), p1, p3, p3, p2]), _dafny.SeqWithoutIsStrInference([default__.fm0(len(d_102_v28_), p1, p3, d_69_i0_, globalState)]), d_97_v23_, d_97_v23_])
                d_112_v38_ = d_112_v38_
                d_113_v39_: _dafny.Set
                d_113_v39_ = _dafny.Set({d_102_v28_})
                d_114_v40_: _dafny.Map
                d_114_v40_ = _dafny.Map({len(d_113_v39_): d_69_i0_})
                d_115_v41_: _dafny.Seq
                d_115_v41_ = _dafny.SeqWithoutIsStrInference([d_114_v40_])
                (globalState).f18 = default__.fm0(d_69_i0_, True, p1, (p0) * (len((d_115_v41_)[default__.safeIndex(p0, len(d_115_v41_))])), globalState)
                (globalState).f10 = (p0) + (d_69_i0_)
                (globalState).f0 = (d_69_i0_) - ((0) - (p0))
        d_116_v42_: _dafny.Seq
        d_116_v42_ = _dafny.SeqWithoutIsStrInference([p1, p3])
        d_117_v43_: C5
        nw12_ = C5()
        nw12_.ctor__(p3, default__.fm10(True, (d_116_v42_)[default__.safeIndex(p0, len(d_116_v42_))], p1, False, globalState), default__.safeDivisionInt(p0, 12))
        d_117_v43_ = nw12_
        if p2:
            (globalState).f10 = (0) - (p0)
            d_118_v44_: str
            d_118_v44_ = _dafny.CodePoint('r')
            d_119_v45_: C3
            nw13_ = C3()
            nw13_.ctor__(True, d_118_v44_, (p0) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qicyr")))))
            d_119_v45_ = nw13_
            d_120_v46_: _dafny.MultiSet
            d_120_v46_ = _dafny.MultiSet([p1])
            d_121_v47_: _dafny.Map
            d_121_v47_ = _dafny.Map({d_120_v46_: p0})
            d_122_v48_: T0
            nw14_ = C4()
            nw14_.ctor__((d_121_v47_) | (d_121_v47_), (d_117_v43_).f26, d_118_v44_, p0)
            d_122_v48_ = nw14_
            d_122_v48_ = d_122_v48_
            d_123_v49_: _dafny.Set
            d_123_v49_ = _dafny.Set({(d_117_v43_).f26})
            d_124_v50_: _dafny.Seq
            d_124_v50_ = _dafny.SeqWithoutIsStrInference([754, d_122_v48_.f25])
            d_125_v51_: C4
            nw15_ = C4()
            nw15_.ctor__(_dafny.Map({d_120_v46_: d_122_v48_.f25}), (p0) < (len(d_123_v49_)), default__.fm10(not(p3), not(p1), (d_119_v45_).fm18(globalState), (d_117_v43_).f26, globalState), len(d_124_v50_))
            d_125_v51_ = nw15_
            d_126_v52_: _dafny.Array
            def lambda2_(d_127_v51_):
                def lambda3_(d_128_i2_):
                    return (d_127_v51_).f30

                return lambda3_

            init1_ = lambda2_(d_125_v51_)
            nw16_ = _dafny.Array(None, 3)
            for i0_1_ in range(nw16_.length(0)):
                nw16_[i0_1_] = init1_(i0_1_)
            d_126_v52_ = nw16_
            index10_ = default__.safeIndex(36, (d_126_v52_).length(0))
            (d_126_v52_)[index10_] = (default__.fm17(p1, p3, not((d_117_v43_).f26), globalState)) >= (default__.fm17((d_117_v43_).f26, p3, d_119_v45_.f31, globalState))
            d_129_v53_: _dafny.Array
            def lambda4_(d_130_i3_):
                return (d_130_i3_) + (928)

            init2_ = lambda4_
            nw17_ = _dafny.Array(None, 19)
            for i0_2_ in range(nw17_.length(0)):
                nw17_[i0_2_] = init2_(i0_2_)
            d_129_v53_ = nw17_
            index11_ = default__.safeIndex(568, (d_129_v53_).length(0))
            (d_129_v53_)[index11_] = p0
            d_131_v54_: _dafny.Seq
            d_131_v54_ = _dafny.SeqWithoutIsStrInference([(d_124_v50_).set(default__.safeIndex(p0, len(d_124_v50_)), 86), d_124_v50_, _dafny.SeqWithoutIsStrInference([p0]), d_124_v50_, d_124_v50_])
            index12_ = default__.safeIndex(36, (d_126_v52_).length(0))
            index13_ = default__.safeIndex(568, (d_129_v53_).length(0))
            rhs7_ = d_122_v48_.f25
            rhs8_ = d_125_v51_
            rhs9_ = (p1) in (_dafny.Set({p1}))
            rhs10_ = -14
            rhs11_ = (0) - (default__.safeModuloInt(d_122_v48_.f25, len((d_131_v54_) + (d_131_v54_))))
            lhs6_ = globalState
            lhs7_ = d_126_v52_
            lhs8_ = default__.safeIndex(36, (d_126_v52_).length(0))
            lhs9_ = globalState
            lhs10_ = d_129_v53_
            lhs11_ = default__.safeIndex(568, (d_129_v53_).length(0))
            lhs6_.f10 = rhs7_
            d_125_v51_ = rhs8_
            lhs7_[lhs8_] = rhs9_
            lhs9_.f11 = rhs10_
            lhs10_[lhs11_] = rhs11_
            d_132_v55_: _dafny.Map
            d_132_v55_ = _dafny.Map({p2: (d_129_v53_)[default__.safeIndex(568, (d_129_v53_).length(0))]})
            d_133_v56_: _dafny.Array
            nw18_ = _dafny.Array(None, 2)
            nw18_[int(0)] = (d_132_v55_ if (d_125_v51_).f30 else d_132_v55_)
            nw18_[int(1)] = d_132_v55_
            d_133_v56_ = nw18_
            d_134_v57_: _dafny.Map
            d_134_v57_ = _dafny.Map({(0) - (len(d_116_v42_)): not((d_117_v43_).f26)})
            index14_ = default__.safeIndex(152, (d_133_v56_).length(0))
            (d_133_v56_)[index14_] = default__.fm31(len(d_134_v57_), (d_117_v43_).f26, default__.fm0(967, p3, p3, len(d_116_v42_), globalState), 262, globalState)
            index15_ = default__.safeIndex(152, (d_133_v56_).length(0))
            (d_133_v56_)[index15_] = d_132_v55_
        elif True:
            d_135_v58_: _dafny.Array
            def lambda5_(d_136_p0_):
                def lambda6_(d_137_i4_):
                    return default__.safeDivisionInt(d_137_i4_, d_136_p0_)

                return lambda6_

            init3_ = lambda5_(p0)
            nw19_ = _dafny.Array(None, 5)
            for i0_3_ in range(nw19_.length(0)):
                nw19_[i0_3_] = init3_(i0_3_)
            d_135_v58_ = nw19_
            d_138_v59_: _dafny.Seq
            d_138_v59_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qdj"))
            d_139_v60_: _dafny.Seq
            d_139_v60_ = _dafny.SeqWithoutIsStrInference([d_138_v59_])
            d_140_v61_: str
            d_140_v61_ = _dafny.CodePoint('q')
            d_141_v62_: _dafny.Map
            d_141_v62_ = _dafny.Map({p0: d_140_v61_})
            d_142_v63_: _dafny.Array
            nw20_ = _dafny.Array(None, 13)
            nw20_[int(0)] = p2
            nw20_[int(1)] = True
            nw20_[int(2)] = p2
            nw20_[int(3)] = (d_139_v60_) <= (d_139_v60_)
            nw20_[int(4)] = p3
            nw20_[int(5)] = True
            nw20_[int(6)] = p2
            nw20_[int(7)] = (d_117_v43_).f26
            nw20_[int(8)] = (d_140_v61_) not in (d_138_v59_)
            nw20_[int(9)] = p3
            nw20_[int(10)] = (p2) not in (d_116_v42_)
            nw20_[int(11)] = (d_141_v62_) != (d_141_v62_)
            nw20_[int(12)] = (p0) <= (-372)
            d_142_v63_ = nw20_
            rhs12_ = not ((d_117_v43_).f26) or (p3)
            rhs13_ = d_135_v58_
            rhs14_ = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_143_i5_ in range(default__.abs(34))]))
            rhs15_ = (p0) + (p0)
            rhs16_ = d_142_v63_
            lhs12_ = globalState
            lhs13_ = globalState
            lhs14_ = globalState
            lhs12_.f18 = rhs12_
            d_135_v58_ = rhs13_
            lhs13_.f16 = rhs14_
            lhs14_.f11 = rhs15_
            d_142_v63_ = rhs16_
            d_144_v64_: D13
            d_144_v64_ = D13_DC29((d_117_v43_).f26, (d_117_v43_).f26, d_142_v63_)
            d_145_v65_: _dafny.Map
            d_145_v65_ = _dafny.Map({D13_DC29(p2, default__.fm0(p0, p3, p2, p0, globalState), d_142_v63_): d_135_v58_})
            rhs17_ = p1
            rhs18_ = (d_144_v64_) in (d_145_v65_)
            rhs19_ = p0
            rhs20_ = (p0 if not(p1) else default__.fm5(p0, p0, globalState))
            lhs15_ = globalState
            lhs16_ = globalState
            lhs17_ = globalState
            lhs18_ = globalState
            lhs15_.f18 = rhs17_
            lhs16_.f18 = rhs18_
            lhs17_.f16 = rhs19_
            lhs18_.f22 = rhs20_
            d_146_v66_: C2
            nw21_ = C2()
            nw21_.ctor__(p0, (d_117_v43_).f26)
            d_146_v66_ = nw21_
            d_147_v67_: _dafny.Set
            d_147_v67_ = _dafny.Set({d_146_v66_, d_146_v66_, d_146_v66_, d_146_v66_, d_146_v66_})
            d_148_v68_: _dafny.Seq
            d_148_v68_ = _dafny.SeqWithoutIsStrInference([d_147_v67_, d_147_v67_, d_147_v67_, d_147_v67_, d_147_v67_])
            d_148_v68_ = (d_148_v68_) + (d_148_v68_)
            d_149_v69_: _dafny.Map
            d_149_v69_ = _dafny.Map({(d_146_v66_).f28: (d_117_v43_).f26})
            d_149_v69_ = (d_149_v69_).set(p1, p1)
            d_150_v70_: C3
            nw22_ = C3()
            nw22_.ctor__(p1, d_140_v61_, p0)
            d_150_v70_ = nw22_
            d_151_v71_: _dafny.Seq
            d_151_v71_ = _dafny.SeqWithoutIsStrInference([d_150_v70_, d_150_v70_, d_150_v70_, d_150_v70_])
            d_152_v72_: _dafny.Map
            d_152_v72_ = _dafny.Map({(d_117_v43_).f26: d_151_v71_})
            d_153_v73_: _dafny.Seq
            d_153_v73_ = _dafny.SeqWithoutIsStrInference([d_151_v71_, d_151_v71_, d_151_v71_, d_151_v71_])
            d_154_v74_: _dafny.Map
            d_154_v74_ = _dafny.Map({(d_146_v66_).f27: d_151_v71_})
            d_155_v75_: T0
            nw23_ = C1()
            nw23_.ctor__(_dafny.CodePoint('g'), p0)
            d_155_v75_ = nw23_
            d_156_v76_: _dafny.Map
            d_156_v76_ = _dafny.Map({d_155_v75_: d_151_v71_})
            d_157_v77_: _dafny.Array
            nw24_ = _dafny.Array(None, 21)
            nw24_[int(0)] = (d_151_v71_ if p2 else d_151_v71_)
            nw24_[int(1)] = d_151_v71_
            nw24_[int(2)] = d_151_v71_
            nw24_[int(3)] = d_151_v71_
            nw24_[int(4)] = d_151_v71_
            nw24_[int(5)] = _dafny.SeqWithoutIsStrInference([d_150_v70_])
            nw24_[int(6)] = d_151_v71_
            nw24_[int(7)] = d_151_v71_
            nw24_[int(8)] = (d_151_v71_) + ((d_151_v71_).set(default__.safeIndex((d_146_v66_).f27, len(d_151_v71_)), d_150_v70_))
            nw24_[int(9)] = d_151_v71_
            nw24_[int(10)] = d_151_v71_
            nw24_[int(11)] = _dafny.SeqWithoutIsStrInference([d_150_v70_, d_150_v70_])
            nw24_[int(12)] = (((d_152_v72_)[not(p2)] if (not(p2)) in (d_152_v72_) else d_151_v71_)) + (((d_153_v73_)[default__.safeIndex(p0, len(d_153_v73_))]).set(default__.safeIndex(p0, len((d_153_v73_)[default__.safeIndex(p0, len(d_153_v73_))])), d_150_v70_))
            nw24_[int(13)] = d_151_v71_
            nw24_[int(14)] = ((d_154_v74_)[(d_146_v66_).f27] if ((d_146_v66_).f27) in (d_154_v74_) else d_151_v71_)
            nw24_[int(15)] = d_151_v71_
            nw24_[int(16)] = _dafny.SeqWithoutIsStrInference([d_150_v70_])
            nw24_[int(17)] = d_151_v71_
            nw24_[int(18)] = d_151_v71_
            nw24_[int(19)] = ((d_156_v76_)[d_155_v75_] if (d_155_v75_) in (d_156_v76_) else d_151_v71_)
            nw24_[int(20)] = d_151_v71_
            d_157_v77_ = nw24_
            d_157_v77_ = d_157_v77_
        d_158_v78_: _dafny.Array
        def lambda7_(d_159_v43_):
            def lambda8_(d_160_i7_):
                return (d_159_v43_).f26

            return lambda8_

        init4_ = lambda7_(d_117_v43_)
        nw25_ = _dafny.Array(None, 22)
        for i0_4_ in range(nw25_.length(0)):
            nw25_[i0_4_] = init4_(i0_4_)
        d_158_v78_ = nw25_
        source4_ = D13_DC29(p2, (_dafny.SeqWithoutIsStrInference([p0 for d_161_i6_ in range(default__.abs(995))])) != (_dafny.SeqWithoutIsStrInference([7, p0])), d_158_v78_)
        if source4_.is_DC29:
            d_162___mcc_h0_ = source4_.cf43
            d_163___mcc_h1_ = source4_.cf44
            d_164___mcc_h2_ = source4_.cf45
            d_165_cf45_ = d_164___mcc_h2_
            d_166_cf44_ = d_163___mcc_h1_
            d_167_cf43_ = d_162___mcc_h0_
            d_168_v79_: _dafny.Map
            d_168_v79_ = _dafny.Map({True: p0})
            d_169_v80_: _dafny.Map
            d_169_v80_ = _dafny.Map({p0: d_168_v79_})
            d_170_v81_: _dafny.Seq
            d_170_v81_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mw"))
            (globalState).f16 = len((d_168_v79_) | (((d_169_v80_)[len(d_170_v81_)] if (len(d_170_v81_)) in (d_169_v80_) else d_168_v79_)))
            d_171_v82_: C0
            nw26_ = C0()
            nw26_.ctor__()
            d_171_v82_ = nw26_
            d_172_v83_: _dafny.Array
            nw27_ = _dafny.Array(_dafny.Set({}), 3)
            d_172_v83_ = nw27_
            d_173_v84_: D10
            d_173_v84_ = D10_DC20(d_166_cf44_, default__.fm5(p0, p0, globalState))
            d_174_v85_: _dafny.Set
            d_174_v85_ = _dafny.Set({d_173_v84_})
            index16_ = default__.safeIndex(375, (d_172_v83_).length(0))
            (d_172_v83_)[index16_] = d_174_v85_
            d_175_v87_: _dafny.Seq
            d_175_v87_ = _dafny.SeqWithoutIsStrInference([d_173_v84_, d_173_v84_])
            d_176_v88_: _dafny.Map
            def iife27_():
                coll27_ = _dafny.Set()
                compr_27_: int
                for compr_27_ in _dafny.IntegerRange(961, 914):
                    d_177_v86_: int = compr_27_
                    if ((961) <= (d_177_v86_)) and ((d_177_v86_) < (914)):
                        coll27_ = coll27_.union(_dafny.Set([(d_177_v86_) + (len(d_170_v81_))]))
                return _dafny.Set(coll27_)
            d_176_v88_ = _dafny.Map({(0) - (len(iife27_()
            )): d_175_v87_})
            d_178_v89_: _dafny.Seq
            d_178_v89_ = _dafny.SeqWithoutIsStrInference([p0])
            index17_ = default__.safeIndex(375, (d_172_v83_).length(0))
            def iife28_():
                coll28_ = _dafny.Set()
                compr_28_: D10
                for compr_28_ in (((d_176_v88_)[len(d_178_v89_)] if (len(d_178_v89_)) in (d_176_v88_) else _dafny.SeqWithoutIsStrInference([D10_DC20((D6_DC12(d_166_cf44_, len(_dafny.Map({(d_117_v43_).f26: p2})))).cf17, p0) for d_179_i8_ in range(default__.abs(782))]))).Elements:
                    d_180_v90_: D10 = compr_28_
                    if (d_180_v90_) in (((d_176_v88_)[len(d_178_v89_)] if (len(d_178_v89_)) in (d_176_v88_) else _dafny.SeqWithoutIsStrInference([D10_DC20((D6_DC12(d_166_cf44_, len(_dafny.Map({(d_117_v43_).f26: p2})))).cf17, p0) for d_179_i8_ in range(default__.abs(782))]))):
                        coll28_ = coll28_.union(_dafny.Set([d_180_v90_]))
                return _dafny.Set(coll28_)
            (d_172_v83_)[index17_] = (d_174_v85_) | (iife28_()
            )
            d_181_v91_: _dafny.Array
            nw28_ = _dafny.Array(int(0), 28)
            d_181_v91_ = nw28_
            d_182_v92_: _dafny.Map
            d_182_v92_ = _dafny.Map({p0: d_117_v43_})
            d_183_v93_: D2
            d_183_v93_ = D2_DC3(len(d_170_v81_), len(d_182_v92_))
            d_184_v94_: D7
            d_184_v94_ = D7_DC14(p2, p0, (d_183_v93_).cf3)
            index18_ = default__.safeIndex(765, (d_181_v91_).length(0))
            (d_181_v91_)[index18_] = default__.safeDivisionInt((d_184_v94_).cf22, len(default__.fm26(globalState)))
            d_185_v95_: C3
            nw29_ = C3()
            nw29_.ctor__(d_166_cf44_, (d_170_v81_)[default__.safeIndex(p0, len(d_170_v81_))], 547)
            d_185_v95_ = nw29_
            d_186_v96_: _dafny.Map
            d_186_v96_ = _dafny.Map({d_185_v95_: d_170_v81_})
            index19_ = default__.safeIndex(765, (d_181_v91_).length(0))
            (d_181_v91_)[index19_] = default__.fm17((d_185_v95_) in (d_186_v96_), p2, (d_117_v43_).f26, globalState)
        elif source4_.is_DC28:
            d_187___mcc_h3_ = source4_.cf42
            d_188_cf42_ = d_187___mcc_h3_
            (globalState).f10 = (p0) * ((667) - (450))
            d_189_v97_: _dafny.MultiSet
            d_189_v97_ = _dafny.MultiSet([p2])
            d_190_v98_: _dafny.Map
            d_190_v98_ = _dafny.Map({d_189_v97_: p0})
            d_191_v99_: str
            d_191_v99_ = _dafny.CodePoint('v')
            d_192_v100_: _dafny.Map
            d_192_v100_ = _dafny.Map({p0: p3})
            d_193_v101_: T1
            nw30_ = C4()
            nw30_.ctor__(d_190_v98_, p3, d_191_v99_, p0)
            d_193_v101_ = nw30_
            d_194_v102_: _dafny.Map
            d_194_v102_ = _dafny.Map({d_193_v101_: p0})
            d_195_v103_: _dafny.MultiSet
            d_195_v103_ = _dafny.MultiSet([len(d_194_v102_), d_193_v101_.f25])
            d_196_v104_: _dafny.Seq
            d_196_v104_ = _dafny.SeqWithoutIsStrInference([p0, len(d_192_v100_), (d_195_v103_).cardinality])
            d_197_v105_: T0
            nw31_ = C4()
            nw31_.ctor__(d_190_v98_, p3, d_191_v99_, len(d_196_v104_))
            d_197_v105_ = nw31_
            d_198_v106_: _dafny.MultiSet
            d_198_v106_ = _dafny.MultiSet([(d_197_v105_ if p2 else d_197_v105_), d_197_v105_])
            rhs21_ = (d_198_v106_) - (d_198_v106_)
            rhs22_ = ((p0) + ((0) - (d_197_v105_.f25))) != (p0)
            lhs19_ = globalState
            d_198_v106_ = rhs21_
            lhs19_.f18 = rhs22_
            if not((not(p3)) and (((d_117_v43_).f26 if p3 else p2))):
                d_199_v107_: C3
                nw32_ = C3()
                nw32_.ctor__((d_117_v43_).f26, (d_197_v105_).f24, d_197_v105_.f25)
                d_199_v107_ = nw32_
                d_200_v108_: C3
                d_200_v108_ = d_199_v107_
                d_201_v109_: _dafny.Map
                d_201_v109_ = _dafny.Map({(d_200_v108_): _dafny.SeqWithoutIsStrInference([563])})
                d_202_v110_: _dafny.MultiSet
                d_202_v110_ = _dafny.MultiSet([d_196_v104_, d_196_v104_, _dafny.SeqWithoutIsStrInference([d_197_v105_.f25])])
                d_203_v111_: _dafny.Array
                nw33_ = _dafny.Array(None, 20)
                nw33_[int(0)] = (d_196_v104_).set(default__.safeIndex(-929, len(d_196_v104_)), (d_196_v104_)[default__.safeIndex(default__.fm17((d_117_v43_).f26, p1, False, globalState), len(d_196_v104_))])
                nw33_[int(1)] = d_196_v104_
                nw33_[int(2)] = _dafny.SeqWithoutIsStrInference([d_193_v101_.f25 for d_204_i9_ in range(default__.abs(817))])
                nw33_[int(3)] = ((d_201_v109_)[d_199_v107_] if (d_199_v107_) in (d_201_v109_) else d_196_v104_)
                nw33_[int(4)] = d_196_v104_
                nw33_[int(5)] = d_196_v104_
                nw33_[int(6)] = d_196_v104_
                nw33_[int(7)] = d_196_v104_
                nw33_[int(8)] = d_196_v104_
                nw33_[int(9)] = _dafny.SeqWithoutIsStrInference([(d_189_v97_).cardinality for d_205_i10_ in range(default__.abs(-330))])
                nw33_[int(10)] = d_196_v104_
                nw33_[int(11)] = d_196_v104_
                nw33_[int(12)] = d_196_v104_
                nw33_[int(13)] = _dafny.SeqWithoutIsStrInference([d_197_v105_.f25, len(d_116_v42_), d_193_v101_.f25])
                nw33_[int(14)] = d_196_v104_
                nw33_[int(15)] = d_196_v104_
                nw33_[int(16)] = _dafny.SeqWithoutIsStrInference([(d_196_v104_)[default__.safeIndex(237, len(d_196_v104_))] for d_206_i11_ in range(default__.abs(-866))])
                nw33_[int(17)] = d_196_v104_
                nw33_[int(18)] = d_196_v104_
                nw33_[int(19)] = (d_196_v104_).set(default__.safeIndex(d_197_v105_.f25, len(d_196_v104_)), (d_202_v110_).cardinality)
                d_203_v111_ = nw33_
                d_207_v112_: _dafny.Map
                d_207_v112_ = _dafny.Map({(d_117_v43_).f26: d_203_v111_})
                d_207_v112_ = (d_207_v112_).set((d_117_v43_).f26, d_203_v111_)
                index20_ = default__.safeIndex(434, (d_203_v111_).length(0))
                (d_203_v111_)[index20_] = d_196_v104_
                index21_ = default__.safeIndex(434, (d_203_v111_).length(0))
                (d_203_v111_)[index21_] = (_dafny.SeqWithoutIsStrInference([d_197_v105_.f25, d_193_v101_.f25])) + ((d_196_v104_) + (d_196_v104_))
                d_208_v113_: _dafny.Map
                d_208_v113_ = _dafny.Map({p3: (d_116_v42_)[default__.safeIndex(p0, len(d_116_v42_))]})
                d_209_v114_: D11
                d_209_v114_ = D11_DC24((d_117_v43_).f26, len(default__.fm34(p2, len(d_208_v113_), (d_117_v43_).f26, globalState)))
                d_210_v115_: C1
                nw34_ = C1()
                nw34_.ctor__((d_193_v101_).f24, (d_193_v101_.f25) - ((d_209_v114_).cf37))
                d_210_v115_ = nw34_
                d_211_v116_: D3
                d_211_v116_ = D3_DC5(d_116_v42_, d_193_v101_.f25, d_193_v101_.f25)
                d_212_v117_: D10
                d_212_v117_ = D10_DC20((d_117_v43_).f26, (d_211_v116_).cf7)
                (globalState).f18 = (d_212_v117_).cf28
                d_191_v99_ = _dafny.CodePoint('a')
            elif True:
                d_213_v118_: _dafny.Map
                d_213_v118_ = _dafny.Map({(d_117_v43_).f26: p0})
                (globalState).f11 = ((d_213_v118_)[p3] if (p3) in (d_213_v118_) else d_193_v101_.f25)
                d_214_v119_: D3
                d_214_v119_ = D3_DC5(d_116_v42_, d_197_v105_.f25, d_197_v105_.f25)
                d_214_v119_ = d_214_v119_
                index22_ = default__.safeIndex(116, (d_158_v78_).length(0))
                (d_158_v78_)[index22_] = (d_189_v97_).isdisjoint(d_189_v97_)
                index23_ = default__.safeIndex(116, (d_158_v78_).length(0))
                (d_158_v78_)[index23_] = (d_117_v43_).f26
                d_215_v120_: _dafny.Array
                nw35_ = _dafny.Array(int(0), 19)
                d_215_v120_ = nw35_
                index24_ = default__.safeIndex(232, (d_215_v120_).length(0))
                (d_215_v120_)[index24_] = d_193_v101_.f25
                d_216_v121_: D10
                d_216_v121_ = D10_DC21(d_193_v101_.f25, not((d_158_v78_)[default__.safeIndex(116, (d_158_v78_).length(0))]), (d_117_v43_).f26)
                d_217_v122_: D11
                d_217_v122_ = D11_DC25(D11_DC24(p3, default__.fm17((D11_DC24(p3, d_197_v105_.f25)).cf36, (d_216_v121_).cf32, p2, globalState)))
                d_218_v123_: _dafny.MultiSet
                d_218_v123_ = _dafny.MultiSet([d_217_v122_, d_217_v122_, d_217_v122_])
                index25_ = default__.safeIndex(232, (d_215_v120_).length(0))
                rhs23_ = 795
                rhs24_ = ((d_213_v118_)[(d_117_v43_).f26] if ((d_117_v43_).f26) in (d_213_v118_) else (d_218_v123_).cardinality)
                rhs25_ = default__.safeModuloInt((-545) + (d_197_v105_.f25), d_193_v101_.f25)
                lhs20_ = d_215_v120_
                lhs21_ = default__.safeIndex(232, (d_215_v120_).length(0))
                lhs22_ = globalState
                lhs23_ = globalState
                lhs20_[lhs21_] = rhs23_
                lhs22_.f22 = rhs24_
                lhs23_.f10 = rhs25_
                d_219_v124_: _dafny.Array
                def lambda9_(d_220_v101_, d_221_v120_):
                    def lambda10_(d_222_i12_):
                        return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Set({(d_221_v120_)[default__.safeIndex(232, (d_221_v120_).length(0))], d_220_v101_.f25}), _dafny.Set({d_220_v101_.f25})]))).issubset(_dafny.MultiSet([_dafny.Set({d_220_v101_.f25})]))

                    return lambda10_

                init5_ = lambda9_(d_193_v101_, d_215_v120_)
                nw36_ = _dafny.Array(None, 29)
                for i0_5_ in range(nw36_.length(0)):
                    nw36_[i0_5_] = init5_(i0_5_)
                d_219_v124_ = nw36_
                d_219_v124_ = (d_219_v124_ if p1 else d_219_v124_)
            d_223_v125_: _dafny.Array
            def lambda11_(d_224_v104_):
                def lambda12_(d_225_i13_):
                    return d_224_v104_

                return lambda12_

            init6_ = lambda11_(d_196_v104_)
            nw37_ = _dafny.Array(None, 23)
            for i0_6_ in range(nw37_.length(0)):
                nw37_[i0_6_] = init6_(i0_6_)
            d_223_v125_ = nw37_
            index26_ = default__.safeIndex(828, (d_223_v125_).length(0))
            (d_223_v125_)[index26_] = d_196_v104_
            d_226_v126_: D10
            d_226_v126_ = D10_DC19((d_197_v105_).f24)
            d_227_v127_: _dafny.Map
            d_227_v127_ = _dafny.Map({d_158_v78_: (d_117_v43_).f26})
            index27_ = default__.safeIndex(828, (d_223_v125_).length(0))
            rhs26_ = (d_197_v105_.f25) * (d_197_v105_.f25)
            rhs27_ = (0) - (p0)
            rhs28_ = _dafny.SeqWithoutIsStrInference([(703) + (-504), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pvsknr")))])
            rhs29_ = (len(d_227_v127_)) + (d_197_v105_.f25)
            rhs30_ = d_226_v126_
            lhs24_ = globalState
            lhs25_ = d_193_v101_
            lhs26_ = d_223_v125_
            lhs27_ = default__.safeIndex(828, (d_223_v125_).length(0))
            lhs28_ = globalState
            lhs24_.f0 = rhs26_
            lhs25_.f25 = rhs27_
            lhs26_[lhs27_] = rhs28_
            lhs28_.f2 = rhs29_
            d_226_v126_ = rhs30_
        elif True:
            d_228___mcc_h4_ = source4_.cf46
            d_229_cf46_ = d_228___mcc_h4_
            d_230_v128_: _dafny.Set
            d_230_v128_ = _dafny.Set({p0})
            d_231_v129_: _dafny.Map
            d_231_v129_ = _dafny.Map({True: len(d_230_v128_)})
            d_232_v130_: D11
            d_232_v130_ = D11_DC24(p2, len(d_116_v42_))
            d_233_v131_: _dafny.Map
            d_233_v131_ = _dafny.Map({((d_231_v129_)[(d_232_v130_).cf36] if ((d_232_v130_).cf36) in (d_231_v129_) else p0): p0})
            (globalState).f10 = ((d_233_v131_)[p0] if (p0) in (d_233_v131_) else p0)
            index28_ = default__.safeIndex(236, (d_158_v78_).length(0))
            (d_158_v78_)[index28_] = p2
            index29_ = default__.safeIndex(236, (d_158_v78_).length(0))
            rhs31_ = (d_117_v43_).f26
            rhs32_ = (0) - ((0) - (p0))
            lhs29_ = d_158_v78_
            lhs30_ = default__.safeIndex(236, (d_158_v78_).length(0))
            lhs31_ = globalState
            lhs29_[lhs30_] = rhs31_
            lhs31_.f0 = rhs32_
            d_234_v132_: _dafny.Map
            d_234_v132_ = _dafny.Map({p0: p3})
            d_235_v133_: _dafny.Set
            d_235_v133_ = _dafny.Set({d_234_v132_})
            d_236_v134_: D3
            d_236_v134_ = D3_DC4(d_235_v133_)
            d_237_v135_: _dafny.Seq
            d_237_v135_ = d_116_v42_
            d_238_v136_: _dafny.Map
            d_238_v136_ = _dafny.Map({(0) - (len((d_237_v135_))): D3_DC4(d_235_v133_)})
            d_239_v137_: _dafny.Array
            def lambda13_(d_240_i14_):
                return _dafny.CodePoint('t')

            init7_ = lambda13_
            nw38_ = _dafny.Array(None, 16)
            for i0_7_ in range(nw38_.length(0)):
                nw38_[i0_7_] = init7_(i0_7_)
            d_239_v137_ = nw38_
            d_241_v138_: _dafny.Seq
            d_241_v138_ = _dafny.SeqWithoutIsStrInference([d_239_v137_, d_239_v137_])
            d_242_v139_: bool
            d_243_v140_: _dafny.Seq
            d_244_v141_: bool
            out0_: bool
            out1_: _dafny.Seq
            out2_: bool
            out0_, out1_, out2_ = (d_117_v43_).m1((_dafny.Map({p0: d_236_v134_})) | (d_238_v136_), ((d_116_v42_)[default__.safeIndex(p0, len(d_116_v42_))]) == (p3), d_241_v138_, globalState)
            d_242_v139_ = out0_
            d_243_v140_ = out1_
            d_244_v141_ = out2_
            d_245_v142_: C0
            nw39_ = C0()
            nw39_.ctor__()
            d_245_v142_ = nw39_
        d_246_v143_: _dafny.Map
        d_246_v143_ = _dafny.Map({p0: p3})
        d_247_v144_: _dafny.Map
        d_247_v144_ = _dafny.Map({d_246_v143_: d_158_v78_})
        def iife29_():
            coll29_ = _dafny.Map()
            compr_29_: int
            for compr_29_ in _dafny.IntegerRange(691, 699):
                d_248_v145_: int = compr_29_
                if ((691) <= (d_248_v145_)) and ((d_248_v145_) < (699)):
                    coll29_[(d_248_v145_) * (919)] = p3
            return _dafny.Map(coll29_)
        d_247_v144_ = (d_247_v144_).set(iife29_()
        , d_158_v78_)
        d_249_v146_: _dafny.Array
        def lambda14_(d_250_p0_):
            def lambda15_(d_251_i15_):
                return default__.safeModuloInt(d_251_i15_, d_250_p0_)

            return lambda15_

        init8_ = lambda14_(p0)
        nw40_ = _dafny.Array(None, 10)
        for i0_8_ in range(nw40_.length(0)):
            nw40_[i0_8_] = init8_(i0_8_)
        d_249_v146_ = nw40_
        index30_ = default__.safeIndex(237, (d_249_v146_).length(0))
        (d_249_v146_)[index30_] = p0
        index31_ = default__.safeIndex(237, (d_249_v146_).length(0))
        (d_249_v146_)[index31_] = (_dafny.MultiSet(((d_116_v42_) + (d_116_v42_)).set(default__.safeIndex(p0, len((d_116_v42_) + (d_116_v42_))), p3))).cardinality
        d_252_v147_: D10
        d_252_v147_ = D10_DC21(124, p2, p1)
        def lambda16_(source5_):
            if source5_.is_DC19:
                d_253___mcc_h5_ = source5_.cf27
                d_254_cf27_ = d_253___mcc_h5_
                return d_254_cf27_
            elif source5_.is_DC20:
                d_255___mcc_h6_ = source5_.cf28
                d_256___mcc_h7_ = source5_.cf29
                d_257_cf29_ = d_256___mcc_h7_
                d_258_cf28_ = d_255___mcc_h6_
                return _dafny.CodePoint('l')
            elif source5_.is_DC21:
                d_259___mcc_h8_ = source5_.cf30
                d_260___mcc_h9_ = source5_.cf31
                d_261___mcc_h10_ = source5_.cf32
                d_262_cf32_ = d_261___mcc_h10_
                d_263_cf31_ = d_260___mcc_h9_
                d_264_cf30_ = d_259___mcc_h8_
                return _dafny.CodePoint('n')
            elif True:
                d_265___mcc_h11_ = source5_.cf26
                d_266_cf26_ = d_265___mcc_h11_
                return _dafny.CodePoint('o')

        r0 = lambda16_(d_252_v147_)
        return r0

    @staticmethod
    def Main(noArgsParameter__):
        d_267_v0_: _dafny.Seq
        d_267_v0_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w')])
        d_268_v1_: bool
        d_268_v1_ = True
        d_269_v2_: int
        d_269_v2_ = 742
        d_270_v3_: _dafny.Map
        d_270_v3_ = _dafny.Map({d_268_v1_: d_269_v2_})
        d_271_v4_: _dafny.Array
        nw41_ = _dafny.Array(False, 27)
        d_271_v4_ = nw41_
        d_272_v5_: _dafny.Array
        nw42_ = _dafny.Array(int(0), 25)
        d_272_v5_ = nw42_
        d_273_v6_: _dafny.Map
        d_273_v6_ = _dafny.Map({d_271_v4_: d_272_v5_})
        d_274_v7_: _dafny.MultiSet
        d_274_v7_ = _dafny.MultiSet([d_268_v1_])
        d_275_v8_: _dafny.Seq
        d_275_v8_ = _dafny.SeqWithoutIsStrInference([(d_274_v7_).cardinality])
        d_276_v9_: _dafny.Set
        d_276_v9_ = _dafny.Set({d_275_v8_})
        d_277_globalState_: GlobalState
        nw43_ = GlobalState()
        nw43_.ctor__(-928, 141, 124, (d_267_v0_) + (d_267_v0_), 227, False, _dafny.CodePoint('d'), d_270_v3_, d_273_v6_, -50, -492, -627, False, (d_270_v3_) | (d_270_v3_), d_276_v9_, 968, 343, False, False, d_275_v8_, 425, True, -118)
        d_277_globalState_ = nw43_
        d_278_v10_: _dafny.MultiSet
        d_278_v10_ = _dafny.MultiSet([d_269_v2_, d_269_v2_])
        d_279_v11_: _dafny.Seq
        d_279_v11_ = _dafny.SeqWithoutIsStrInference([d_268_v1_, default__.fm0((d_278_v10_).cardinality, d_268_v1_, d_268_v1_, len(d_270_v3_), d_277_globalState_)])
        d_280_v12_: _dafny.Seq
        d_280_v12_ = d_279_v11_
        d_281_v13_: _dafny.Seq
        d_281_v13_ = _dafny.SeqWithoutIsStrInference([d_272_v5_, d_272_v5_, d_272_v5_, d_272_v5_])
        d_282_v14_: _dafny.Seq
        d_282_v14_ = d_267_v0_
        d_283_v15_: _dafny.Seq
        d_283_v15_ = _dafny.SeqWithoutIsStrInference([(d_282_v14_), d_267_v0_, d_267_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fln")), d_267_v0_])
        d_284_v16_: str
        d_284_v16_ = _dafny.CodePoint('j')
        rhs33_ = d_267_v0_
        rhs34_ = d_269_v2_
        rhs35_ = d_269_v2_
        rhs36_ = (0) - ((d_269_v2_) - (len(((d_280_v12_)) + (d_279_v11_))))
        rhs37_ = (d_281_v13_)[default__.safeIndex(len(((d_283_v15_)[default__.safeIndex(d_269_v2_, len(d_283_v15_))]).set(default__.safeIndex(len((d_283_v15_)[default__.safeIndex(d_269_v2_, len(d_283_v15_))]), len((d_283_v15_)[default__.safeIndex(d_269_v2_, len(d_283_v15_))])), d_284_v16_)), len(d_281_v13_))]
        lhs32_ = d_277_globalState_
        lhs33_ = d_277_globalState_
        lhs34_ = d_277_globalState_
        d_267_v0_ = rhs33_
        lhs32_.f11 = rhs34_
        lhs33_.f10 = rhs35_
        lhs34_.f16 = rhs36_
        d_272_v5_ = rhs37_
        (d_277_globalState_).f2 = (0) - (len(((d_275_v8_) + (d_275_v8_)) + (d_275_v8_)))
        d_285_i0_: int
        d_285_i0_ = 0
        with _dafny.label("0"):
            while True:
                with _dafny.c_label("0"):
                    if (d_285_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_285_i0_ = (d_285_i0_) + (1)
                    d_286_v17_: C4
                    nw44_ = C4()
                    nw44_.ctor__(_dafny.Map({d_274_v7_: len(d_279_v11_)}), not(d_268_v1_), d_284_v16_, d_269_v2_)
                    d_286_v17_ = nw44_
                    d_284_v16_ = (d_284_v16_ if not (not(d_268_v1_)) or (d_268_v1_) else d_284_v16_)
                    source6_ = d_282_v14_
                    d_287___mcc_h0_ = source6_
                    d_288_cf1_ = d_287___mcc_h0_
                    index32_ = default__.safeIndex(947, (d_271_v4_).length(0))
                    (d_271_v4_)[index32_] = False
                    d_289_v18_: _dafny.Map
                    d_289_v18_ = _dafny.Map({d_269_v2_: ((d_270_v3_)[(d_279_v11_)[default__.safeIndex(d_269_v2_, len(d_279_v11_))]] if ((d_279_v11_)[default__.safeIndex(d_269_v2_, len(d_279_v11_))]) in (d_270_v3_) else d_269_v2_)})
                    d_290_v19_: D3
                    d_290_v19_ = D3_DC6(((d_289_v18_)[d_269_v2_] if (d_269_v2_) in (d_289_v18_) else d_269_v2_), d_269_v2_, (d_286_v17_).f30)
                    index33_ = default__.safeIndex(947, (d_271_v4_).length(0))
                    (d_271_v4_)[index33_] = (d_290_v19_).cf11
                    (d_277_globalState_).f0 = (d_275_v8_)[default__.safeIndex(d_269_v2_, len(d_275_v8_))]
                    (d_277_globalState_).f15 = d_269_v2_
                    index34_ = default__.safeIndex(947, (d_271_v4_).length(0))
                    (d_271_v4_)[index34_] = False
                    d_291_v20_: D5
                    d_291_v20_ = D5_DC10(False, True)
                    d_292_v21_: _dafny.Map
                    d_292_v21_ = _dafny.Map({not(not((d_286_v17_).f30)): d_291_v20_})
                    (d_286_v17_).m5((d_292_v21_) | (d_292_v21_), True, d_269_v2_, d_277_globalState_)
                    pass
            pass
        d_293_v22_: _dafny.Set
        d_293_v22_ = _dafny.Set({d_269_v2_, len(d_279_v11_)})
        d_294_v23_: _dafny.Map
        d_294_v23_ = _dafny.Map({(d_275_v8_)[default__.safeIndex(d_269_v2_, len(d_275_v8_))]: default__.fm0(d_269_v2_, d_268_v1_, False, ((d_270_v3_)[d_268_v1_] if (d_268_v1_) in (d_270_v3_) else d_269_v2_), d_277_globalState_)})
        (d_277_globalState_).f2 = (len(d_293_v22_) if False else len(d_294_v23_))
        d_295_v24_: str
        out3_: str
        out3_ = default__.m6(d_269_v2_, d_268_v1_, d_268_v1_, d_268_v1_, d_277_globalState_)
        d_295_v24_ = out3_
        d_296_v25_: D11
        d_296_v25_ = D11_DC23(d_269_v2_, d_268_v1_)
        d_297_v26_: D11
        d_297_v26_ = D11_DC25(d_296_v25_)
        d_298_v27_: _dafny.Array
        nw45_ = _dafny.Array(None, 4)
        nw45_[int(0)] = d_297_v26_
        nw45_[int(1)] = d_297_v26_
        nw45_[int(2)] = d_297_v26_
        nw45_[int(3)] = D11_DC25(d_296_v25_)
        d_298_v27_ = nw45_
        d_299_v28_: _dafny.Set
        d_299_v28_ = _dafny.Set({d_298_v27_, d_298_v27_})
        d_300_v29_: _dafny.Map
        d_300_v29_ = _dafny.Map({d_268_v1_: d_299_v28_})
        d_300_v29_ = (d_300_v29_).set(d_268_v1_, d_299_v28_)
        if d_268_v1_:
            d_301_v30_: T1
            nw46_ = C3()
            nw46_.ctor__(True, d_284_v16_, (d_275_v8_)[default__.safeIndex(d_269_v2_, len(d_275_v8_))])
            d_301_v30_ = nw46_
            d_302_v31_: _dafny.MultiSet
            d_302_v31_ = _dafny.MultiSet([d_301_v30_])
            d_303_v32_: str
            out4_: str
            out4_ = default__.m6(d_269_v2_, (d_269_v2_) == (default__.fm5(d_269_v2_, (d_302_v31_).cardinality, d_277_globalState_)), (d_267_v0_) == (d_267_v0_), (d_269_v2_) < (d_269_v2_), d_277_globalState_)
            d_303_v32_ = out4_
            if ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xvkixbphw"))) + (_dafny.SeqWithoutIsStrInference([d_295_v24_ for d_304_i1_ in range(default__.abs(618))]))) < ((d_267_v0_) + (_dafny.SeqWithoutIsStrInference([d_284_v16_ for d_305_i2_ in range(default__.abs(912))]))):
                d_268_v1_ = ((len(_dafny.Set({d_301_v30_}))) * (d_301_v30_.f25)) in ((d_275_v8_) + (d_275_v8_))
                (d_277_globalState_).f18 = d_268_v1_
                d_306_v33_: str
                out5_: str
                out5_ = default__.m6((d_301_v30_.f25) * (d_269_v2_), d_268_v1_, (d_274_v7_).isdisjoint(_dafny.MultiSet([d_268_v1_, d_268_v1_])), (d_269_v2_) >= (d_301_v30_.f25), d_277_globalState_)
                d_306_v33_ = out5_
                d_307_v34_: _dafny.Map
                d_307_v34_ = _dafny.Map({d_274_v7_: d_301_v30_.f25})
                d_308_v36_: _dafny.Set
                d_308_v36_ = _dafny.Set({default__.fm8(d_301_v30_.f25, True, _dafny.SeqWithoutIsStrInference([d_268_v1_]), d_277_globalState_)})
                d_309_v37_: C4
                nw47_ = C4()
                def iife30_():
                    coll30_ = _dafny.Map()
                    compr_30_: _dafny.MultiSet
                    for compr_30_ in (d_308_v36_).Elements:
                        d_310_v35_: _dafny.MultiSet = compr_30_
                        if (d_310_v35_) in (d_308_v36_):
                            coll30_[d_310_v35_] = d_301_v30_.f25
                    return _dafny.Map(coll30_)
                nw47_.ctor__((d_307_v34_) | (iife30_()
                ), (d_268_v1_ if True else d_268_v1_), _dafny.CodePoint('b'), d_301_v30_.f25)
                d_309_v37_ = nw47_
                (d_277_globalState_).f3 = _dafny.SeqWithoutIsStrInference([(d_301_v30_).f24 for d_311_i3_ in range(default__.abs(941))])
            elif True:
                d_312_v38_: _dafny.Map
                d_312_v38_ = _dafny.Map({(d_278_v10_).isdisjoint(d_278_v10_): d_271_v4_})
                d_313_v39_: _dafny.Map
                d_313_v39_ = _dafny.Map({(d_301_v30_).f24: d_268_v1_})
                d_314_v40_: D13
                d_314_v40_ = D13_DC29(d_268_v1_, d_268_v1_, d_271_v4_)
                d_312_v38_ = ((_dafny.Map({d_268_v1_: d_271_v4_})).set(((d_313_v39_)[(d_301_v30_).f24] if ((d_301_v30_).f24) in (d_313_v39_) else d_268_v1_), (d_314_v40_).cf45)) | (d_312_v38_)
                (d_277_globalState_).f18 = d_268_v1_
                d_315_v41_: C6
                nw48_ = C6()
                nw48_.ctor__(d_294_v23_)
                d_315_v41_ = nw48_
                d_316_v42_: _dafny.Map
                d_316_v42_ = _dafny.Map({d_268_v1_: _dafny.CodePoint('u')})
                d_316_v42_ = (d_316_v42_).set(d_268_v1_, (d_301_v30_).f24)
                d_317_v43_: _dafny.Array
                nw49_ = _dafny.Array(None, 3)
                d_317_v43_ = nw49_
                d_317_v43_ = (d_317_v43_ if False else d_317_v43_)
            (d_277_globalState_).f18 = (d_295_v24_) in (_dafny.Map({d_284_v16_: d_268_v1_}))
            (d_277_globalState_).f16 = d_269_v2_
            if ((d_268_v1_ if d_268_v1_ else d_268_v1_)) == (False):
                d_318_v44_: _dafny.Set
                d_318_v44_ = _dafny.Set({(d_301_v30_.f25) != (d_269_v2_)})
                d_319_v45_: _dafny.Array
                nw50_ = _dafny.Array(_dafny.CodePoint('D'), 13)
                d_319_v45_ = nw50_
                index35_ = default__.safeIndex(843, (d_319_v45_).length(0))
                (d_319_v45_)[index35_] = d_303_v32_
                index36_ = default__.safeIndex(843, (d_319_v45_).length(0))
                rhs38_ = ((d_269_v2_) + (d_269_v2_) if (d_279_v11_)[default__.safeIndex(d_269_v2_, len(d_279_v11_))] else d_301_v30_.f25)
                rhs39_ = ((_dafny.Set({d_268_v1_})) - (d_318_v44_)) - (d_318_v44_)
                rhs40_ = d_268_v1_
                rhs41_ = _dafny.CodePoint('y')
                rhs42_ = (d_272_v5_ if d_268_v1_ else d_272_v5_)
                lhs35_ = d_277_globalState_
                lhs36_ = d_277_globalState_
                lhs37_ = d_319_v45_
                lhs38_ = default__.safeIndex(843, (d_319_v45_).length(0))
                lhs35_.f16 = rhs38_
                d_318_v44_ = rhs39_
                lhs36_.f18 = rhs40_
                lhs37_[lhs38_] = rhs41_
                d_272_v5_ = rhs42_
                index37_ = default__.safeIndex(553, (d_272_v5_).length(0))
                (d_272_v5_)[index37_] = len((d_275_v8_) + (d_275_v8_))
                index38_ = default__.safeIndex(553, (d_272_v5_).length(0))
                (d_272_v5_)[index38_] = d_301_v30_.f25
                d_320_v46_: _dafny.MultiSet
                d_320_v46_ = _dafny.MultiSet([d_267_v0_])
                (d_277_globalState_).f16 = (0) - ((d_320_v46_).cardinality)
                d_294_v23_ = (d_294_v23_).set(d_301_v30_.f25, d_268_v1_)
                (d_277_globalState_).f18 = d_268_v1_
            elif True:
                d_321_v47_: _dafny.Map
                d_321_v47_ = _dafny.Map({d_274_v7_: d_301_v30_.f25})
                d_322_v48_: D3
                d_322_v48_ = D3_DC6(d_269_v2_, d_301_v30_.f25, d_268_v1_)
                d_323_v49_: D5
                d_323_v49_ = D5_DC10(d_268_v1_, (d_322_v48_).cf11)
                d_324_v50_: T0
                nw51_ = C4()
                nw51_.ctor__((_dafny.Map({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([((d_294_v23_)[d_269_v2_] if (d_269_v2_) in (d_294_v23_) else d_268_v1_), d_268_v1_])): d_269_v2_})) | (d_321_v47_), d_268_v1_, default__.fm10(d_268_v1_, d_268_v1_, not((d_323_v49_).cf15), d_268_v1_, d_277_globalState_), d_269_v2_)
                d_324_v50_ = nw51_
                d_325_v51_: _dafny.Set
                d_325_v51_ = _dafny.Set({True})
                d_326_v52_: D6
                d_326_v52_ = D6_DC11(d_324_v50_)
                d_324_v50_ = (d_324_v50_ if default__.fm0(len(d_325_v51_), False, d_268_v1_, d_269_v2_, d_277_globalState_) else (d_326_v52_).cf16)
                (d_277_globalState_).f18 = d_268_v1_
                index39_ = default__.safeIndex(375, (d_272_v5_).length(0))
                (d_272_v5_)[index39_] = (d_269_v2_) * (d_324_v50_.f25)
                index40_ = default__.safeIndex(375, (d_272_v5_).length(0))
                (d_272_v5_)[index40_] = (0) - (len(d_279_v11_))
                (d_277_globalState_).f18 = d_268_v1_
                d_327_v53_: C0
                nw52_ = C0()
                nw52_.ctor__()
                d_327_v53_ = nw52_
        elif True:
            index41_ = default__.safeIndex(814, (d_272_v5_).length(0))
            (d_272_v5_)[index41_] = d_269_v2_
            index42_ = default__.safeIndex(814, (d_272_v5_).length(0))
            (d_272_v5_)[index42_] = default__.safeModuloInt(default__.safeDivisionInt(161, len(_dafny.SeqWithoutIsStrInference([d_269_v2_]))), d_269_v2_)
            (d_277_globalState_).f11 = len(d_267_v0_)
            d_328_v54_: C2
            nw53_ = C2()
            nw53_.ctor__(len(default__.fm7(d_277_globalState_)), d_268_v1_)
            d_328_v54_ = nw53_
            d_329_v55_: _dafny.Seq
            d_329_v55_ = _dafny.SeqWithoutIsStrInference([(d_328_v54_ if d_268_v1_ else d_328_v54_)])
            d_329_v55_ = (d_329_v55_) + ((d_329_v55_) + (d_329_v55_))
            d_279_v11_ = d_279_v11_
            d_328_v54_ = d_328_v54_
        if d_268_v1_:
            d_330_v56_: str
            out6_: str
            out6_ = default__.m6(138, not((-166) > (d_269_v2_)), (d_268_v1_) == (d_268_v1_), d_268_v1_, d_277_globalState_)
            d_330_v56_ = out6_
            d_331_v57_: C0
            nw54_ = C0()
            nw54_.ctor__()
            d_331_v57_ = nw54_
            d_332_v58_: _dafny.Map
            d_332_v58_ = _dafny.Map({d_331_v57_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kmijxxnsm"))})
            d_332_v58_ = (d_332_v58_).set(d_331_v57_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ogexx")))
            d_333_v59_: C5
            nw55_ = C5()
            nw55_.ctor__(d_268_v1_, d_295_v24_, d_269_v2_)
            d_333_v59_ = nw55_
            d_333_v59_ = d_333_v59_
            d_272_v5_ = d_272_v5_
            d_334_v60_: _dafny.Set
            d_334_v60_ = _dafny.Set({d_272_v5_})
            d_335_v61_: D13
            d_335_v61_ = D13_DC28(d_334_v60_)
            d_335_v61_ = d_335_v61_
        elif True:
            d_336_v62_: _dafny.Map
            d_336_v62_ = _dafny.Map({d_269_v2_: d_274_v7_})
            d_337_v63_: _dafny.Map
            d_337_v63_ = _dafny.Map({d_336_v62_: len(d_267_v0_)})
            d_337_v63_ = (d_337_v63_).set(default__.fm32(d_268_v1_, d_268_v1_, d_277_globalState_), (0) - (((d_274_v7_)[d_268_v1_] if (d_268_v1_) in (d_274_v7_) else d_269_v2_)))
            d_269_v2_ = (0) - (d_269_v2_)
            (d_277_globalState_).f19 = _dafny.SeqWithoutIsStrInference([(d_269_v2_) * (d_269_v2_) for d_338_i4_ in range(default__.abs(2))])
            d_268_v1_ = not(d_268_v1_)
            d_268_v1_ = (len(default__.fm15(len(d_270_v3_), d_269_v2_, d_269_v2_, d_277_globalState_))) < (d_269_v2_)
        d_339_v64_: D3
        d_339_v64_ = D3_DC6(d_269_v2_, (0) - (d_269_v2_), d_268_v1_)
        d_340_v65_: _dafny.Map
        d_340_v65_ = _dafny.Map({d_269_v2_: d_269_v2_})
        d_341_v66_: _dafny.Map
        d_341_v66_ = _dafny.Map({d_269_v2_: len(d_340_v65_)})
        d_342_v67_: str
        out7_: str
        out7_ = default__.m6(d_269_v2_, default__.fm0(d_269_v2_, default__.fm0((d_339_v64_).cf10, d_268_v1_, d_268_v1_, d_269_v2_, d_277_globalState_), d_268_v1_, len(d_267_v0_), d_277_globalState_), d_268_v1_, (((d_341_v66_)[default__.fm5(d_269_v2_, default__.fm17(d_268_v1_, d_268_v1_, d_268_v1_, d_277_globalState_), d_277_globalState_)] if (default__.fm5(d_269_v2_, default__.fm17(d_268_v1_, d_268_v1_, d_268_v1_, d_277_globalState_), d_277_globalState_)) in (d_341_v66_) else len(d_267_v0_))) <= (d_269_v2_), d_277_globalState_)
        d_342_v67_ = out7_
        d_343_i5_: int
        d_343_i5_ = 0
        with _dafny.label("1"):
            while d_268_v1_:
                with _dafny.c_label("1"):
                    if (d_343_i5_) >= (100):
                        raise _dafny.Break("1")
                    d_343_i5_ = (d_343_i5_) + (1)
                    if (False) or (d_268_v1_):
                        (d_277_globalState_).f18 = d_268_v1_
                        d_344_v68_: _dafny.Map
                        d_344_v68_ = _dafny.Map({d_274_v7_: d_269_v2_})
                        d_345_v69_: C4
                        nw56_ = C4()
                        nw56_.ctor__(d_344_v68_, not((d_268_v1_) == (d_268_v1_)), _dafny.CodePoint('f'), (0) - (default__.safeModuloInt(d_269_v2_, d_269_v2_)))
                        d_345_v69_ = nw56_
                        d_346_v70_: D12
                        d_346_v70_ = D12_DC27(d_269_v2_, d_268_v1_)
                        (d_277_globalState_).f15 = (d_269_v2_) - ((d_346_v70_).cf40)
                        d_347_v71_: D5
                        d_347_v71_ = D5_DC10((d_345_v69_).f30, d_268_v1_)
                        pat_let_tv0_ = d_268_v1_
                        d_348_v72_: _dafny.Map
                        def iife31_(_pat_let0_0):
                            def iife32_(d_349_dt__update__tmp_h0_):
                                def iife33_(_pat_let1_0):
                                    def iife34_(d_350_dt__update_hcf14_h0_):
                                        return D5_DC10(d_350_dt__update_hcf14_h0_, (d_349_dt__update__tmp_h0_).cf15)
                                    return iife34_(_pat_let1_0)
                                return iife33_(pat_let_tv0_)
                            return iife32_(_pat_let0_0)
                        d_348_v72_ = _dafny.Map({(d_345_v69_).f30: iife31_(d_347_v71_)})
                        (d_345_v69_).m5(d_348_v72_, d_268_v1_, default__.safeModuloInt(d_269_v2_, d_269_v2_), d_277_globalState_)
                        d_281_v13_ = d_281_v13_
                    elif True:
                        d_351_v73_: C1
                        nw57_ = C1()
                        nw57_.ctor__(d_284_v16_, (d_269_v2_) - (len(d_267_v0_)))
                        d_351_v73_ = nw57_
                        (d_277_globalState_).f10 = len((d_267_v0_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "untmsi"))))
                        d_352_v74_: _dafny.Map
                        d_352_v74_ = _dafny.Map({d_295_v24_: d_268_v1_})
                        d_353_v75_: _dafny.Array
                        nw58_ = _dafny.Array(_dafny.CodePoint('D'), 25)
                        d_353_v75_ = nw58_
                        d_354_v76_: _dafny.Map
                        d_354_v76_ = _dafny.Map({(len(d_352_v74_) if d_268_v1_ else d_269_v2_): d_353_v75_})
                        d_354_v76_ = (d_354_v76_).set((d_351_v73_).fm6(d_277_globalState_), d_353_v75_)
                        d_355_v77_: T1
                        nw59_ = C3()
                        nw59_.ctor__(True, d_295_v24_, len(d_275_v8_))
                        d_355_v77_ = nw59_
                        d_356_v78_: _dafny.Map
                        d_356_v78_ = _dafny.Map({d_355_v77_: d_268_v1_})
                        d_268_v1_ = ((d_355_v77_) in (d_356_v78_) if d_268_v1_ else (((d_294_v23_)[d_269_v2_] if (d_269_v2_) in (d_294_v23_) else d_268_v1_)) == (default__.fm0(763, True, d_268_v1_, d_355_v77_.f25, d_277_globalState_)))
                        d_357_v79_: _dafny.Map
                        d_357_v79_ = _dafny.Map({d_274_v7_: d_355_v77_.f25})
                        d_358_v81_: _dafny.Set
                        d_358_v81_ = _dafny.Set({_dafny.MultiSet([not(d_268_v1_), d_268_v1_])})
                        d_359_v82_: C4
                        nw60_ = C4()
                        def iife35_():
                            coll31_ = _dafny.Map()
                            compr_31_: _dafny.MultiSet
                            for compr_31_ in (d_358_v81_).Elements:
                                d_360_v80_: _dafny.MultiSet = compr_31_
                                if (d_360_v80_) in (d_358_v81_):
                                    coll31_[d_360_v80_] = d_355_v77_.f25
                            return _dafny.Map(coll31_)
                        nw60_.ctor__((d_357_v79_) | (iife35_()
                        ), d_268_v1_, d_295_v24_, d_269_v2_)
                        d_359_v82_ = nw60_
                    d_361_v83_: _dafny.Set
                    d_361_v83_ = _dafny.Set({d_272_v5_, d_272_v5_, d_272_v5_})
                    d_362_v84_: _dafny.Map
                    d_362_v84_ = _dafny.Map({((d_279_v11_)[default__.safeIndex(len(d_267_v0_), len(d_279_v11_))]) or (d_268_v1_): d_361_v83_})
                    d_362_v84_ = (d_362_v84_).set(d_268_v1_, (d_361_v83_ if d_268_v1_ else d_361_v83_))
                    index43_ = default__.safeIndex(816, (d_271_v4_).length(0))
                    (d_271_v4_)[index43_] = d_268_v1_
                    d_363_v85_: _dafny.Array
                    def lambda17_(d_364_v0_):
                        def lambda18_(d_365_i6_):
                            return d_364_v0_

                        return lambda18_

                    init9_ = lambda17_(d_267_v0_)
                    nw61_ = _dafny.Array(None, 6)
                    for i0_9_ in range(nw61_.length(0)):
                        nw61_[i0_9_] = init9_(i0_9_)
                    d_363_v85_ = nw61_
                    index44_ = default__.safeIndex(206, (d_363_v85_).length(0))
                    (d_363_v85_)[index44_] = d_267_v0_
                    d_366_v86_: _dafny.MultiSet
                    d_366_v86_ = _dafny.MultiSet([d_342_v67_])
                    index45_ = default__.safeIndex(816, (d_271_v4_).length(0))
                    index46_ = default__.safeIndex(206, (d_363_v85_).length(0))
                    rhs43_ = default__.fm5(d_269_v2_, ((d_366_v86_) - (_dafny.MultiSet([_dafny.CodePoint('f')]))).cardinality, d_277_globalState_)
                    rhs44_ = d_268_v1_
                    rhs45_ = (d_267_v0_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wdfjarp")))
                    lhs39_ = d_277_globalState_
                    lhs40_ = d_271_v4_
                    lhs41_ = default__.safeIndex(816, (d_271_v4_).length(0))
                    lhs42_ = d_363_v85_
                    lhs43_ = default__.safeIndex(206, (d_363_v85_).length(0))
                    lhs39_.f2 = rhs43_
                    lhs40_[lhs41_] = rhs44_
                    lhs42_[lhs43_] = rhs45_
                    index47_ = default__.safeIndex(950, (d_272_v5_).length(0))
                    (d_272_v5_)[index47_] = (d_275_v8_)[default__.safeIndex(d_269_v2_, len(d_275_v8_))]
                    index48_ = default__.safeIndex(950, (d_272_v5_).length(0))
                    (d_272_v5_)[index48_] = ((d_278_v10_)[d_269_v2_] if (d_269_v2_) in (d_278_v10_) else d_269_v2_)
                    pass
            pass
        d_269_v2_ = d_269_v2_
        (d_277_globalState_).f15 = (d_269_v2_ if d_268_v1_ else d_269_v2_)
        d_367_v87_: _dafny.Map
        d_367_v87_ = _dafny.Map({False: d_272_v5_})
        d_368_v88_: _dafny.Array
        nw62_ = _dafny.Array(None, 11)
        nw62_[int(0)] = d_272_v5_
        nw62_[int(1)] = d_272_v5_
        nw62_[int(2)] = d_272_v5_
        nw62_[int(3)] = d_272_v5_
        nw62_[int(4)] = ((d_367_v87_)[d_268_v1_] if (d_268_v1_) in (d_367_v87_) else d_272_v5_)
        nw62_[int(5)] = d_272_v5_
        nw62_[int(6)] = d_272_v5_
        nw62_[int(7)] = d_272_v5_
        nw62_[int(8)] = d_272_v5_
        nw62_[int(9)] = d_272_v5_
        nw62_[int(10)] = d_272_v5_
        d_368_v88_ = nw62_
        d_368_v88_ = d_368_v88_
        d_369_v89_: _dafny.Array
        def lambda19_(d_370_i8_):
            return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_371_i9_ in range(default__.abs(903))])

        init10_ = lambda19_
        nw63_ = _dafny.Array(None, 14)
        for i0_10_ in range(nw63_.length(0)):
            nw63_[i0_10_] = init10_(i0_10_)
        d_369_v89_ = nw63_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_369_v89_).length(0)):
            d_372_i7_: int = guard_loop_0_
            if (True) and (((0) <= (d_372_i7_)) and ((d_372_i7_) < ((d_369_v89_).length(0)))):
                (d_369_v89_)[(d_372_i7_)] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sn"))) + (d_267_v0_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "infvwkvh")))
        d_373_v90_: _dafny.Map
        d_373_v90_ = _dafny.Map({(d_279_v11_)[default__.safeIndex(default__.fm17(not(d_268_v1_), d_268_v1_, d_268_v1_, d_277_globalState_), len(d_279_v11_))]: d_268_v1_})
        if not(((d_373_v90_)[d_268_v1_] if (d_268_v1_) in (d_373_v90_) else (d_269_v2_) > ((0) - (d_269_v2_)))):
            d_374_v91_: C6
            nw64_ = C6()
            nw64_.ctor__(((d_294_v23_).set(d_269_v2_, False) if d_268_v1_ else d_294_v23_))
            d_374_v91_ = nw64_
            d_375_v92_: C3
            nw65_ = C3()
            nw65_.ctor__(d_268_v1_, d_342_v67_, 224)
            d_375_v92_ = nw65_
            d_376_v93_: C1
            nw66_ = C1()
            nw66_.ctor__(d_342_v67_, (d_275_v8_)[default__.safeIndex(d_269_v2_, len(d_275_v8_))])
            d_376_v93_ = nw66_
            d_376_v93_ = (d_376_v93_ if not(d_375_v92_.f31) else d_376_v93_)
            index49_ = default__.safeIndex(919, (d_272_v5_).length(0))
            (d_272_v5_)[index49_] = d_269_v2_
            index50_ = default__.safeIndex(919, (d_272_v5_).length(0))
            (d_272_v5_)[index50_] = ((d_340_v65_)[d_269_v2_] if (d_269_v2_) in (d_340_v65_) else default__.safeDivisionInt(d_269_v2_, len(d_279_v11_)))
            (d_277_globalState_).f15 = (((d_272_v5_)[default__.safeIndex(919, (d_272_v5_).length(0))]) - (d_269_v2_)) - (d_269_v2_)
        elif True:
            d_377_v94_: D7
            d_377_v94_ = D7_DC13(d_294_v23_)
            d_378_v95_: C1
            nw67_ = C1()
            nw67_.ctor__(d_295_v24_, (len((d_377_v94_).cf19)) - (d_269_v2_))
            d_378_v95_ = nw67_
            index51_ = default__.safeIndex(706, (d_272_v5_).length(0))
            (d_272_v5_)[index51_] = d_269_v2_
            index52_ = default__.safeIndex(617, (d_369_v89_).length(0))
            (d_369_v89_)[index52_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yj"))) + (d_267_v0_)
            d_379_v96_: D10
            d_379_v96_ = D10_DC21(((d_270_v3_)[d_268_v1_] if (d_268_v1_) in (d_270_v3_) else d_269_v2_), d_268_v1_, not(False))
            d_380_v97_: _dafny.MultiSet
            d_380_v97_ = _dafny.MultiSet([d_379_v96_, d_379_v96_, d_379_v96_, d_379_v96_, D10_DC21(d_269_v2_, not(default__.fm0((_dafny.MultiSet(d_279_v11_)).cardinality, d_268_v1_, d_268_v1_, len(d_267_v0_), d_277_globalState_)), d_268_v1_)])
            pat_let_tv1_ = d_268_v1_
            index53_ = default__.safeIndex(706, (d_272_v5_).length(0))
            index54_ = default__.safeIndex(617, (d_369_v89_).length(0))
            def iife36_(_pat_let2_0):
                def iife37_(d_381_dt__update__tmp_h1_):
                    def iife38_(_pat_let3_0):
                        def iife39_(d_382_dt__update_hcf32_h0_):
                            return D10_DC21((d_381_dt__update__tmp_h1_).cf30, (d_381_dt__update__tmp_h1_).cf31, d_382_dt__update_hcf32_h0_)
                        return iife39_(_pat_let3_0)
                    return iife38_(pat_let_tv1_)
                return iife37_(_pat_let2_0)
            rhs46_ = d_269_v2_
            rhs47_ = d_267_v0_
            rhs48_ = (d_293_v22_) | (d_293_v22_)
            rhs49_ = iife36_(d_379_v96_)
            rhs50_ = ((d_380_v97_ if d_268_v1_ else d_380_v97_)).intersection((_dafny.MultiSet([d_379_v96_])) | (d_380_v97_))
            lhs44_ = d_272_v5_
            lhs45_ = default__.safeIndex(706, (d_272_v5_).length(0))
            lhs46_ = d_369_v89_
            lhs47_ = default__.safeIndex(617, (d_369_v89_).length(0))
            lhs44_[lhs45_] = rhs46_
            lhs46_[lhs47_] = rhs47_
            d_293_v22_ = rhs48_
            d_379_v96_ = rhs49_
            d_380_v97_ = rhs50_
            d_383_v98_: _dafny.Map
            d_383_v98_ = _dafny.Map({d_274_v7_: 683})
            d_384_v99_: T1
            nw68_ = C4()
            nw68_.ctor__(d_383_v98_, d_268_v1_, d_284_v16_, 350)
            d_384_v99_ = nw68_
            d_385_v100_: _dafny.Map
            d_385_v100_ = _dafny.Map({482: d_384_v99_})
            d_385_v100_ = (d_385_v100_).set((len(d_270_v3_) if d_268_v1_ else d_269_v2_), d_384_v99_)
            d_386_v101_: _dafny.Map
            d_386_v101_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))).set(default__.safeIndex(d_269_v2_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")))), d_284_v16_): d_269_v2_})
            d_387_v102_: _dafny.Map
            d_387_v102_ = _dafny.Map({(d_378_v95_).fm6(d_277_globalState_): d_386_v101_})
            d_387_v102_ = (d_387_v102_).set(d_384_v99_.f25, default__.fm33(d_277_globalState_))
            d_267_v0_ = (d_369_v89_)[default__.safeIndex(617, (d_369_v89_).length(0))]
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_272_v5_).length(0)):
            d_388_i10_: int = guard_loop_1_
            if (True) and (((0) <= (d_388_i10_)) and ((d_388_i10_) < ((d_272_v5_).length(0)))):
                (d_272_v5_)[(d_388_i10_)] = default__.safeModuloInt(d_388_i10_, d_269_v2_)
        _dafny.print(_dafny.string_of((d_267_v0_) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_268_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_269_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v3_) == (_dafny.Map({True: 742}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_271_v4_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_271_v4_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v5_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v5_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v5_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v5_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v5_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v5_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v5_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v5_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v5_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v5_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v5_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v5_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v5_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v5_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v5_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v5_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v5_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v5_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v5_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v5_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v5_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v5_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v5_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v5_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v5_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_273_v6_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_274_v7_) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_275_v8_) == (_dafny.SeqWithoutIsStrInference([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_276_v9_) == (_dafny.Set({_dafny.SeqWithoutIsStrInference([1])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_277_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_277_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_277_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_277_globalState_.f3) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w'), _dafny.CodePoint('w')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_277_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_277_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_277_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_277_globalState_).f7) == (_dafny.Map({True: 742}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_277_globalState_).f8)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_277_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_277_globalState_.f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_277_globalState_.f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_277_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_277_globalState_.f13) == (_dafny.Map({True: 742}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_277_globalState_).f14) == (_dafny.Set({_dafny.SeqWithoutIsStrInference([1])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_277_globalState_.f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_277_globalState_.f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_277_globalState_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_277_globalState_.f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_277_globalState_.f19) == (_dafny.SeqWithoutIsStrInference([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_277_globalState_).f20))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_277_globalState_).f21))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_277_globalState_.f22))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_278_v10_) == (_dafny.MultiSet([742, 742]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_279_v11_) == (_dafny.SeqWithoutIsStrInference([True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_280_v12_)) == (_dafny.SeqWithoutIsStrInference([True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_281_v13_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_282_v14_)) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_283_v15_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w')]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w')]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w')]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fln")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w')])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_284_v16_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_285_i0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_293_v22_) == (_dafny.Set({742, 2}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_294_v23_) == (_dafny.Map({1: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_295_v24_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_296_v25_).cf34))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_296_v25_).cf35))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_297_v26_).cf38).cf34))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_297_v26_).cf38).cf35))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_298_v27_)[0]).cf38).cf34))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_298_v27_)[0]).cf38).cf35))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_298_v27_)[1]).cf38).cf34))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_298_v27_)[1]).cf38).cf35))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_298_v27_)[2]).cf38).cf34))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_298_v27_)[2]).cf38).cf35))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_298_v27_)[3]).cf38).cf34))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_298_v27_)[3]).cf38).cf35))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_299_v28_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_300_v29_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_339_v64_).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_339_v64_).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_339_v64_).cf11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_340_v65_) == (_dafny.Map({742: 742}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_341_v66_) == (_dafny.Map({742: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_342_v67_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_343_i5_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_367_v87_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[0])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[0])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[0])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[0])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[0])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[0])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[0])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[0])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[0])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[0])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[0])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[0])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[0])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[0])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[0])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[0])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[0])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[0])[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[0])[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[0])[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[0])[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[0])[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[0])[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[0])[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[0])[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[1])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[1])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[1])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[1])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[1])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[1])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[1])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[1])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[1])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[1])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[1])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[1])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[1])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[1])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[1])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[1])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[1])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[1])[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[1])[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[1])[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[1])[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[1])[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[1])[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[1])[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[1])[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[2])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[2])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[2])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[2])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[2])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[2])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[2])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[2])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[2])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[2])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[2])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[2])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[2])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[2])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[2])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[2])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[2])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[2])[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[2])[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[2])[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[2])[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[2])[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[2])[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[2])[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[2])[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[3])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[3])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[3])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[3])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[3])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[3])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[3])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[3])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[3])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[3])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[3])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[3])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[3])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[3])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[3])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[3])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[3])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[3])[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[3])[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[3])[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[3])[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[3])[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[3])[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[3])[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[3])[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[4])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[4])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[4])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[4])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[4])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[4])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[4])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[4])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[4])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[4])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[4])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[4])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[4])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[4])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[4])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[4])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[4])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[4])[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[4])[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[4])[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[4])[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[4])[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[4])[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[4])[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[4])[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[5])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[5])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[5])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[5])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[5])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[5])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[5])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[5])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[5])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[5])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[5])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[5])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[5])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[5])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[5])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[5])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[5])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[5])[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[5])[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[5])[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[5])[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[5])[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[5])[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[5])[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[5])[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[6])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[6])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[6])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[6])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[6])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[6])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[6])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[6])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[6])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[6])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[6])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[6])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[6])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[6])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[6])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[6])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[6])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[6])[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[6])[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[6])[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[6])[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[6])[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[6])[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[6])[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[6])[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[7])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[7])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[7])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[7])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[7])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[7])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[7])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[7])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[7])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[7])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[7])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[7])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[7])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[7])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[7])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[7])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[7])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[7])[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[7])[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[7])[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[7])[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[7])[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[7])[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[7])[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[7])[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[8])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[8])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[8])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[8])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[8])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[8])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[8])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[8])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[8])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[8])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[8])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[8])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[8])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[8])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[8])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[8])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[8])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[8])[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[8])[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[8])[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[8])[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[8])[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[8])[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[8])[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[8])[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[9])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[9])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[9])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[9])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[9])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[9])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[9])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[9])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[9])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[9])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[9])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[9])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[9])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[9])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[9])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[9])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[9])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[9])[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[9])[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[9])[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[9])[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[9])[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[9])[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[9])[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[9])[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[10])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[10])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[10])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[10])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[10])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[10])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[10])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[10])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[10])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[10])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[10])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[10])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[10])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[10])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[10])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[10])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[10])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[10])[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[10])[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[10])[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[10])[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[10])[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[10])[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[10])[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_368_v88_)[10])[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_369_v89_)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_369_v89_)[1]) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_369_v89_)[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_369_v89_)[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_369_v89_)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_369_v89_)[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_369_v89_)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_369_v89_)[7]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_369_v89_)[8]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_369_v89_)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_369_v89_)[10]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_369_v89_)[11]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_369_v89_)[12]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_369_v89_)[13]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_373_v90_) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
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
        return lambda: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D1_DC1)

class D1_DC1(D1, NamedTuple('DC1', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC1({self.cf1.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC1) and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC3(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D2_DC3)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D2_DC2)

class D2_DC3(D2, NamedTuple('DC3', [('cf3', Any), ('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC3({_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC3) and self.cf3 == __o.cf3 and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC2(D2, NamedTuple('DC2', [('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC2({_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC2) and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC5(_dafny.Seq({}), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D3_DC5)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D3_DC6)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D3_DC4)

class D3_DC5(D3, NamedTuple('DC5', [('cf6', Any), ('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC5({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC5) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC6(D3, NamedTuple('DC6', [('cf9', Any), ('cf10', Any), ('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC6({_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC6) and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC4(D3, NamedTuple('DC4', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC4({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC4) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC8()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D4_DC8)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D4_DC7)

class D4_DC8(D4, NamedTuple('DC8', [])):
    def __dafnystr__(self) -> str:
        return f'D4.DC8'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC8)
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
        return lambda: D5_DC10(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D5_DC10)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D5_DC9)

class D5_DC10(D5, NamedTuple('DC10', [('cf14', Any), ('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC10({_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC10) and self.cf14 == __o.cf14 and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC9(D5, NamedTuple('DC9', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC9({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC9) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC12(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D6_DC12)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D6_DC11)

class D6_DC12(D6, NamedTuple('DC12', [('cf17', Any), ('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC12({_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC12) and self.cf17 == __o.cf17 and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC11(D6, NamedTuple('DC11', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC11({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC11) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC14(False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D7_DC14)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D7_DC13)

class D7_DC14(D7, NamedTuple('DC14', [('cf20', Any), ('cf21', Any), ('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC14({_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC14) and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC13(D7, NamedTuple('DC13', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC13({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC13) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D8_DC15)

class D8_DC15(D8, NamedTuple('DC15', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC15({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC15) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC17(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D9_DC17)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D9_DC16)

class D9_DC17(D9, NamedTuple('DC17', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC17({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC17) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC16(D9, NamedTuple('DC16', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC16({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC16) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC19(_dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D10_DC19)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D10_DC20)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D10_DC21)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D10_DC18)

class D10_DC19(D10, NamedTuple('DC19', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC19({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC19) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC20(D10, NamedTuple('DC20', [('cf28', Any), ('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC20({_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC20) and self.cf28 == __o.cf28 and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC21(D10, NamedTuple('DC21', [('cf30', Any), ('cf31', Any), ('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC21({_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC21) and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC18(D10, NamedTuple('DC18', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC18({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC18) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC23(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D11_DC23)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D11_DC24)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D11_DC22)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D11_DC25)

class D11_DC23(D11, NamedTuple('DC23', [('cf34', Any), ('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC23({_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC23) and self.cf34 == __o.cf34 and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC24(D11, NamedTuple('DC24', [('cf36', Any), ('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC24({_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC24) and self.cf36 == __o.cf36 and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC22(D11, NamedTuple('DC22', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC22({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC22) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC25(D11, NamedTuple('DC25', [('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC25({_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC25) and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC27(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D12_DC27)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D12_DC26)

class D12_DC27(D12, NamedTuple('DC27', [('cf40', Any), ('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC27({_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC27) and self.cf40 == __o.cf40 and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC26(D12, NamedTuple('DC26', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC26({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC26) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC29(False, False, _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D13_DC29)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D13_DC28)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D13_DC30)

class D13_DC29(D13, NamedTuple('DC29', [('cf43', Any), ('cf44', Any), ('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC29({_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)}, {_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC29) and self.cf43 == __o.cf43 and self.cf44 == __o.cf44 and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC28(D13, NamedTuple('DC28', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC28({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC28) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC30(D13, NamedTuple('DC30', [('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC30({_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC30) and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D14_DC31)

class D14_DC31(D14, NamedTuple('DC31', [('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC31({_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC31) and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D15_DC32)

class D15_DC32(D15, NamedTuple('DC32', [('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC32({_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC32) and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC34(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D16_DC34)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D16_DC33)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D16_DC35)

class D16_DC34(D16, NamedTuple('DC34', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC34({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC34) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC33(D16, NamedTuple('DC33', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC33({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC33) and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC35(D16, NamedTuple('DC35', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC35({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC35) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    @property
    def f25(self):
        return self._f25
    @f25.setter
    def f25(self, value):
        self._f25 = value

class T1:
    pass

class GlobalState:
    def  __init__(self):
        self.f0: int = int(0)
        self.f2: int = int(0)
        self.f3: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f10: int = int(0)
        self.f11: int = int(0)
        self.f13: _dafny.Map = _dafny.Map({})
        self.f15: int = int(0)
        self.f16: int = int(0)
        self.f18: bool = False
        self.f19: _dafny.Seq = _dafny.Seq({})
        self.f22: int = int(0)
        self._f1: int = int(0)
        self._f4: int = int(0)
        self._f5: bool = False
        self._f6: str = _dafny.CodePoint('D')
        self._f7: _dafny.Map = _dafny.Map({})
        self._f8: _dafny.Map = _dafny.Map({})
        self._f9: int = int(0)
        self._f12: bool = False
        self._f14: _dafny.Set = _dafny.Set({})
        self._f17: bool = False
        self._f20: int = int(0)
        self._f21: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22):
        (self).f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
        (self).f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self).f10 = f10
        (self).f11 = f11
        (self)._f12 = f12
        (self).f13 = f13
        (self)._f14 = f14
        (self).f15 = f15
        (self).f16 = f16
        (self)._f17 = f17
        (self).f18 = f18
        (self).f19 = f19
        (self)._f20 = f20
        (self)._f21 = f21
        (self).f22 = f22

    @property
    def f1(self):
        return self._f1
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
    def f12(self):
        return self._f12
    @property
    def f14(self):
        return self._f14
    @property
    def f17(self):
        return self._f17
    @property
    def f20(self):
        return self._f20
    @property
    def f21(self):
        return self._f21

class C0:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass


class C1(T0):
    def  __init__(self):
        self._f25: int = int(0)
        self._f24: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f25(self):
        return self._f25
    @f25.setter
    def f25(self, value):
        self._f25 = value
    @property
    def f24(self):
        return self._f24
    def ctor__(self, f24, f25):
        (self)._f24 = f24
        (self).f25 = f25

    def fm3(self, p0, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ywk"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nyrhpcet")))

    def fm6(self, globalState):
        return default__.safeModuloInt(default__.safeModuloInt(self.f25, self.f25), len(_dafny.Map({_dafny.SeqWithoutIsStrInference([False, True]): len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([self.f25 for d_389_i0_ in range(default__.abs(-385))]))]))})))

    def m4(self, p0, p1, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: int = int(0)
        r3: int = int(0)
        d_390_v0_: _dafny.Map
        d_390_v0_ = _dafny.Map({p0: p0})
        d_391_v1_: _dafny.Array
        def lambda20_(d_392_p0_):
            def lambda21_(d_393_i0_):
                return d_392_p0_

            return lambda21_

        init11_ = lambda20_(p0)
        nw69_ = _dafny.Array(None, 9)
        for i0_11_ in range(nw69_.length(0)):
            nw69_[i0_11_] = init11_(i0_11_)
        d_391_v1_ = nw69_
        d_394_v2_: _dafny.Set
        d_394_v2_ = _dafny.Set({d_391_v1_})
        d_390_v0_ = (d_390_v0_).set((d_394_v2_).issubset(d_394_v2_), p0)
        (self).f25 = p1
        hi1_ = (self).fm6(globalState)
        for d_395_i1_ in range(p1, hi1_):
            d_396_v3_: _dafny.Array
            nw70_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 6)
            d_396_v3_ = nw70_
            d_396_v3_ = d_396_v3_
            d_397_v4_: D5
            d_397_v4_ = D5_DC10(p0, not(p0))
            source7_ = d_397_v4_
            if source7_.is_DC10:
                d_398___mcc_h0_ = source7_.cf14
                d_399___mcc_h1_ = source7_.cf15
                d_400_cf15_ = d_399___mcc_h1_
                d_401_cf14_ = d_398___mcc_h0_
                d_402_v5_: _dafny.Array
                nw71_ = _dafny.Array(int(0), 29)
                d_402_v5_ = nw71_
                index55_ = default__.safeIndex(929, (d_402_v5_).length(0))
                (d_402_v5_)[index55_] = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_400_cf15_]))).cardinality
                index56_ = default__.safeIndex(844, (d_391_v1_).length(0))
                (d_391_v1_)[index56_] = d_401_cf14_
                index57_ = default__.safeIndex(163, (d_402_v5_).length(0))
                (d_402_v5_)[index57_] = d_395_i1_
                d_403_v6_: _dafny.Map
                d_403_v6_ = _dafny.Map({self.f25: p0})
                d_404_v7_: _dafny.Seq
                d_404_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tpjon"))
                d_405_v8_: _dafny.Seq
                d_405_v8_ = _dafny.SeqWithoutIsStrInference([d_400_cf15_, p0])
                d_406_v10_: _dafny.Seq
                d_406_v10_ = _dafny.SeqWithoutIsStrInference([self.f25])
                d_407_v11_: _dafny.Set
                d_407_v11_ = _dafny.Set({p0})
                d_408_v12_: D7
                d_408_v12_ = D7_DC13(d_403_v6_)
                d_409_v14_: _dafny.Array
                nw72_ = _dafny.Array(None, 28)
                nw72_[int(0)] = d_403_v6_
                nw72_[int(1)] = d_403_v6_
                nw72_[int(2)] = d_403_v6_
                nw72_[int(3)] = _dafny.Map({len(d_404_v7_): (d_405_v8_)[default__.safeIndex(p1, len(d_405_v8_))]})
                nw72_[int(4)] = d_403_v6_
                nw72_[int(5)] = d_403_v6_
                nw72_[int(6)] = d_403_v6_
                nw72_[int(7)] = d_403_v6_
                nw72_[int(8)] = d_403_v6_
                nw72_[int(9)] = d_403_v6_
                nw72_[int(10)] = d_403_v6_
                nw72_[int(11)] = d_403_v6_
                nw72_[int(12)] = d_403_v6_
                nw72_[int(13)] = _dafny.Map({len(d_404_v7_): d_401_cf14_})
                nw72_[int(14)] = d_403_v6_
                nw72_[int(15)] = d_403_v6_
                nw72_[int(16)] = (d_403_v6_).set((self).fm6(globalState), d_400_cf15_)
                def iife40_():
                    coll32_ = _dafny.Map()
                    compr_32_: int
                    for compr_32_ in (d_406_v10_).Elements:
                        d_410_v9_: int = compr_32_
                        if (d_410_v9_) in (d_406_v10_):
                            coll32_[default__.safeModuloInt(d_410_v9_, 560)] = d_401_cf14_
                    return _dafny.Map(coll32_)
                nw72_[int(17)] = iife40_()
                
                nw72_[int(18)] = (d_403_v6_).set(len(d_407_v11_), d_400_cf15_)
                nw72_[int(19)] = d_403_v6_
                nw72_[int(20)] = d_403_v6_
                nw72_[int(21)] = d_403_v6_
                nw72_[int(22)] = d_403_v6_
                nw72_[int(23)] = d_403_v6_
                nw72_[int(24)] = _dafny.Map({p1: d_401_cf14_})
                nw72_[int(25)] = (d_408_v12_).cf19
                nw72_[int(26)] = d_403_v6_
                def iife41_():
                    coll33_ = _dafny.Map()
                    compr_33_: int
                    for compr_33_ in (d_406_v10_).Elements:
                        d_411_v13_: int = compr_33_
                        if (d_411_v13_) in (d_406_v10_):
                            coll33_[(d_411_v13_) - (d_395_i1_)] = d_401_cf14_
                    return _dafny.Map(coll33_)
                nw72_[int(27)] = iife41_()
                
                d_409_v14_ = nw72_
                d_412_v15_: D4
                d_412_v15_ = D4_DC7(d_409_v14_)
                d_413_v16_: _dafny.Map
                d_413_v16_ = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference([d_395_i1_])})
                index58_ = default__.safeIndex(929, (d_402_v5_).length(0))
                index59_ = default__.safeIndex(844, (d_391_v1_).length(0))
                index60_ = default__.safeIndex(163, (d_402_v5_).length(0))
                rhs51_ = (d_395_i1_) - ((d_395_i1_) - (d_395_i1_))
                rhs52_ = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "duj")))
                rhs53_ = d_401_cf14_
                rhs54_ = (0) - ((d_395_i1_) * (default__.safeModuloInt(len(d_413_v16_), self.f25)))
                rhs55_ = d_412_v15_
                lhs48_ = d_402_v5_
                lhs49_ = default__.safeIndex(929, (d_402_v5_).length(0))
                lhs50_ = globalState
                lhs51_ = d_391_v1_
                lhs52_ = default__.safeIndex(844, (d_391_v1_).length(0))
                lhs53_ = d_402_v5_
                lhs54_ = default__.safeIndex(163, (d_402_v5_).length(0))
                lhs48_[lhs49_] = rhs51_
                lhs50_.f22 = rhs52_
                lhs51_[lhs52_] = rhs53_
                lhs53_[lhs54_] = rhs54_
                d_412_v15_ = rhs55_
                d_414_v17_: _dafny.Array
                nw73_ = _dafny.Array(_dafny.Array(None, 0), 19)
                d_414_v17_ = nw73_
                d_414_v17_ = d_414_v17_
                index61_ = default__.safeIndex(929, (d_402_v5_).length(0))
                (d_402_v5_)[index61_] = (default__.safeDivisionInt((d_402_v5_)[default__.safeIndex(929, (d_402_v5_).length(0))], p1)) * (d_395_i1_)
                d_402_v5_ = (d_402_v5_ if ((d_391_v1_)[default__.safeIndex(844, (d_391_v1_).length(0))] if d_400_cf15_ else True) else d_402_v5_)
            elif True:
                d_415___mcc_h2_ = source7_.cf13
                d_416_cf13_ = d_415___mcc_h2_
                index62_ = default__.safeIndex(487, (d_416_cf13_).length(0))
                (d_416_cf13_)[index62_] = d_395_i1_
                index63_ = default__.safeIndex(487, (d_416_cf13_).length(0))
                (d_416_cf13_)[index63_] = (0) - ((d_395_i1_) - (d_395_i1_))
                (globalState).f2 = (d_416_cf13_)[default__.safeIndex(487, (d_416_cf13_).length(0))]
                d_417_v18_: _dafny.Map
                d_417_v18_ = _dafny.Map({p0: d_416_cf13_})
                d_418_v19_: _dafny.MultiSet
                d_418_v19_ = _dafny.MultiSet([p0])
                d_419_v20_: _dafny.Map
                d_419_v20_ = _dafny.Map({d_395_i1_: True})
                index64_ = default__.safeIndex(487, (d_416_cf13_).length(0))
                rhs56_ = ((d_418_v19_)[not (True) or (default__.fm0(d_395_i1_, p0, p0, self.f25, globalState))] if (not (True) or (default__.fm0(d_395_i1_, p0, p0, self.f25, globalState))) in (d_418_v19_) else len(d_419_v20_))
                rhs57_ = (d_417_v18_) | (d_417_v18_)
                rhs58_ = d_395_i1_
                rhs59_ = p0
                rhs60_ = d_416_cf13_
                lhs55_ = globalState
                lhs56_ = d_416_cf13_
                lhs57_ = default__.safeIndex(487, (d_416_cf13_).length(0))
                lhs58_ = globalState
                lhs55_.f15 = rhs56_
                d_417_v18_ = rhs57_
                lhs56_[lhs57_] = rhs58_
                lhs58_.f18 = rhs59_
                d_416_cf13_ = rhs60_
                d_420_v21_: C0
                nw74_ = C0()
                nw74_.ctor__()
                d_420_v21_ = nw74_
            d_421_v22_: _dafny.Array
            nw75_ = _dafny.Array(None, 21)
            nw75_[int(0)] = (d_391_v1_ if p0 else d_391_v1_)
            nw75_[int(1)] = d_391_v1_
            nw75_[int(2)] = (d_391_v1_ if p0 else d_391_v1_)
            nw75_[int(3)] = d_391_v1_
            nw75_[int(4)] = d_391_v1_
            nw75_[int(5)] = (d_391_v1_ if p0 else d_391_v1_)
            nw75_[int(6)] = d_391_v1_
            nw75_[int(7)] = d_391_v1_
            nw75_[int(8)] = d_391_v1_
            nw75_[int(9)] = d_391_v1_
            nw75_[int(10)] = d_391_v1_
            nw75_[int(11)] = d_391_v1_
            nw75_[int(12)] = d_391_v1_
            nw75_[int(13)] = d_391_v1_
            nw75_[int(14)] = d_391_v1_
            nw75_[int(15)] = d_391_v1_
            nw75_[int(16)] = d_391_v1_
            nw75_[int(17)] = d_391_v1_
            nw75_[int(18)] = d_391_v1_
            nw75_[int(19)] = d_391_v1_
            nw75_[int(20)] = d_391_v1_
            d_421_v22_ = nw75_
            d_422_v23_: _dafny.Seq
            d_422_v23_ = _dafny.SeqWithoutIsStrInference([D5_DC10(p0, False)])
            d_423_v24_: _dafny.Seq
            d_423_v24_ = _dafny.SeqWithoutIsStrInference([not(p0)])
            d_424_v25_: _dafny.Seq
            d_424_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nnhp"))
            d_425_v26_: _dafny.Seq
            d_425_v26_ = _dafny.SeqWithoutIsStrInference([len(d_424_v25_)])
            d_426_v27_: _dafny.Array
            nw76_ = _dafny.Array(None, 14)
            nw76_[int(0)] = (d_395_i1_) - (self.f25)
            nw76_[int(1)] = d_395_i1_
            nw76_[int(2)] = self.f25
            nw76_[int(3)] = (len(d_422_v23_) if p0 else self.f25)
            nw76_[int(4)] = d_395_i1_
            nw76_[int(5)] = self.f25
            nw76_[int(6)] = len(d_423_v24_)
            nw76_[int(7)] = d_395_i1_
            nw76_[int(8)] = p1
            nw76_[int(9)] = d_395_i1_
            nw76_[int(10)] = (d_395_i1_) - (d_395_i1_)
            nw76_[int(11)] = p1
            nw76_[int(12)] = self.f25
            nw76_[int(13)] = (d_425_v26_)[default__.safeIndex(p1, len(d_425_v26_))]
            d_426_v27_ = nw76_
            index65_ = default__.safeIndex(949, (d_426_v27_).length(0))
            (d_426_v27_)[index65_] = self.f25
            d_427_v28_: _dafny.Set
            d_427_v28_ = _dafny.Set({p0, False, p0, p0})
            index66_ = default__.safeIndex(949, (d_426_v27_).length(0))
            rhs61_ = ((d_427_v28_).intersection(d_427_v28_)).ispropersubset((_dafny.Set({p0})) - (d_427_v28_))
            rhs62_ = d_421_v22_
            rhs63_ = ((self.f25) - ((0) - (p1))) * (-958)
            rhs64_ = d_397_v4_
            lhs59_ = globalState
            lhs60_ = d_426_v27_
            lhs61_ = default__.safeIndex(949, (d_426_v27_).length(0))
            lhs59_.f18 = rhs61_
            d_421_v22_ = rhs62_
            lhs60_[lhs61_] = rhs63_
            d_397_v4_ = rhs64_
            d_428_v29_: D6
            d_428_v29_ = D6_DC12(p0, self.f25)
            (globalState).f18 = (d_428_v29_).cf17
        d_429_v30_: _dafny.Set
        d_429_v30_ = _dafny.Set({not(p0)})
        d_430_v31_: _dafny.Seq
        d_430_v31_ = _dafny.SeqWithoutIsStrInference([len(d_429_v30_), p1, p1])
        r3 = (-490 if p0 else (d_430_v31_)[default__.safeIndex((0) - (p1), len(d_430_v31_))])
        (globalState).f18 = p0
        d_431_v32_: _dafny.Map
        d_431_v32_ = _dafny.Map({295: self.f25})
        if (self.f25) >= (((d_431_v32_)[p1] if (p1) in (d_431_v32_) else self.f25)):
            if not((p0 if p0 else p0)):
                index67_ = default__.safeIndex(397, (d_391_v1_).length(0))
                (d_391_v1_)[index67_] = (p0 if p0 else p0)
                d_432_v33_: _dafny.Map
                d_432_v33_ = _dafny.Map({self.f25: p0})
                index68_ = default__.safeIndex(397, (d_391_v1_).length(0))
                (d_391_v1_)[index68_] = (((d_432_v33_).set(self.f25, False) if p0 else d_432_v33_)) == (_dafny.Map({176: True}))
                d_433_v34_: _dafny.Array
                def lambda22_(d_434_v1_):
                    def lambda23_(d_435_i2_):
                        return (d_434_v1_)[default__.safeIndex(397, (d_434_v1_).length(0))]

                    return lambda23_

                init12_ = lambda22_(d_391_v1_)
                nw77_ = _dafny.Array(None, 13)
                for i0_12_ in range(nw77_.length(0)):
                    nw77_[i0_12_] = init12_(i0_12_)
                d_433_v34_ = nw77_
                d_433_v34_ = d_433_v34_
                r0 = -777
                d_436_v35_: _dafny.Seq
                d_436_v35_ = _dafny.SeqWithoutIsStrInference([not((d_391_v1_)[default__.safeIndex(397, (d_391_v1_).length(0))])])
                d_436_v35_ = default__.fm7(globalState)
                (globalState).f18 = default__.fm0((p1) * (self.f25), (d_391_v1_)[default__.safeIndex(397, (d_391_v1_).length(0))], p0, (self).fm6(globalState), globalState)
            elif True:
                d_437_v36_: _dafny.MultiSet
                d_437_v36_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([(self).f24 for d_438_i3_ in range(default__.abs(309))]))])
                d_439_v37_: D7
                d_439_v37_ = D7_DC14(p0, 882, self.f25)
                (globalState).f18 = (len(d_429_v30_)) > (((d_437_v36_).cardinality) - ((d_439_v37_).cf21))
                (globalState).f18 = p0
                d_440_v38_: _dafny.Seq
                d_440_v38_ = _dafny.SeqWithoutIsStrInference([p0])
                d_441_v39_: _dafny.Seq
                d_441_v39_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ybuioxvd"))
                d_442_v40_: _dafny.MultiSet
                d_442_v40_ = _dafny.MultiSet([p0, p0, p0, p0])
                d_443_v41_: _dafny.Array
                nw78_ = _dafny.Array(None, 12)
                nw78_[int(0)] = (len(_dafny.SeqWithoutIsStrInference([self.f25]))) + (self.f25)
                nw78_[int(1)] = (self.f25) - ((0) - (p1))
                nw78_[int(2)] = (self).fm6(globalState)
                nw78_[int(3)] = p1
                nw78_[int(4)] = self.f25
                nw78_[int(5)] = ((default__.fm8(len(d_441_v39_), p0, d_440_v38_, globalState)) - (d_442_v40_)).cardinality
                nw78_[int(6)] = p1
                nw78_[int(7)] = (self.f25) + (p1)
                nw78_[int(8)] = (687) * (self.f25)
                nw78_[int(9)] = (self).fm6(globalState)
                nw78_[int(10)] = p1
                nw78_[int(11)] = (0) - (p1)
                d_443_v41_ = nw78_
                index69_ = default__.safeIndex(18, (d_443_v41_).length(0))
                (d_443_v41_)[index69_] = (self).fm6(globalState)
                d_444_v42_: D6
                d_444_v42_ = D6_DC12(p0, p1)
                d_445_v43_: D3
                d_445_v43_ = D3_DC4(default__.fm9(True, (d_440_v38_).set(default__.safeIndex(p1, len(d_440_v38_)), (d_444_v42_).cf17), globalState))
                index70_ = default__.safeIndex(18, (d_443_v41_).length(0))
                rhs65_ = (_dafny.SeqWithoutIsStrInference([p0])).set(default__.safeIndex(self.f25, len(_dafny.SeqWithoutIsStrInference([p0]))), default__.fm0(p1, False, p0, self.f25, globalState))
                rhs66_ = (self).fm3((self).f24, globalState)
                rhs67_ = (self).fm6(globalState)
                rhs68_ = d_445_v43_
                lhs62_ = globalState
                lhs63_ = d_443_v41_
                lhs64_ = default__.safeIndex(18, (d_443_v41_).length(0))
                d_440_v38_ = rhs65_
                lhs62_.f3 = rhs66_
                lhs63_[lhs64_] = rhs67_
                d_445_v43_ = rhs68_
                d_446_v44_: D5
                d_446_v44_ = D5_DC9(d_443_v41_)
                pat_let_tv2_ = d_443_v41_
                d_447_v45_: _dafny.Array
                nw79_ = _dafny.Array(None, 16)
                nw79_[int(0)] = D5_DC9(d_443_v41_)
                nw79_[int(1)] = d_446_v44_
                nw79_[int(2)] = d_446_v44_
                nw79_[int(3)] = D5_DC9(d_443_v41_)
                nw79_[int(4)] = D5_DC9(d_443_v41_)
                nw79_[int(5)] = d_446_v44_
                nw79_[int(6)] = d_446_v44_
                nw79_[int(7)] = d_446_v44_
                nw79_[int(8)] = D5_DC9(d_443_v41_)
                def iife42_(_pat_let4_0):
                    def iife43_(d_448_dt__update__tmp_h0_):
                        def iife44_(_pat_let5_0):
                            def iife45_(d_449_dt__update_hcf13_h0_):
                                return D5_DC9(d_449_dt__update_hcf13_h0_)
                            return iife45_(_pat_let5_0)
                        return iife44_(pat_let_tv2_)
                    return iife43_(_pat_let4_0)
                nw79_[int(9)] = iife42_(D5_DC9(d_443_v41_))
                nw79_[int(10)] = d_446_v44_
                nw79_[int(11)] = d_446_v44_
                nw79_[int(12)] = d_446_v44_
                nw79_[int(13)] = d_446_v44_
                nw79_[int(14)] = d_446_v44_
                nw79_[int(15)] = D5_DC9(d_443_v41_)
                d_447_v45_ = nw79_
                index71_ = default__.safeIndex(440, (d_447_v45_).length(0))
                (d_447_v45_)[index71_] = d_446_v44_
                pat_let_tv3_ = d_443_v41_
                index72_ = default__.safeIndex(440, (d_447_v45_).length(0))
                def iife46_(_pat_let6_0):
                    def iife47_(d_450_dt__update__tmp_h1_):
                        def iife48_(_pat_let7_0):
                            def iife49_(d_451_dt__update_hcf13_h1_):
                                return D5_DC9(d_451_dt__update_hcf13_h1_)
                            return iife49_(_pat_let7_0)
                        return iife48_(pat_let_tv3_)
                    return iife47_(_pat_let6_0)
                (d_447_v45_)[index72_] = iife46_(d_446_v44_)
                index73_ = default__.safeIndex(887, (d_391_v1_).length(0))
                (d_391_v1_)[index73_] = p0
                index74_ = default__.safeIndex(887, (d_391_v1_).length(0))
                index75_ = default__.safeIndex(18, (d_443_v41_).length(0))
                rhs69_ = p0
                rhs70_ = (0) - (default__.safeModuloInt(self.f25, (d_443_v41_)[default__.safeIndex(18, (d_443_v41_).length(0))]))
                rhs71_ = p1
                rhs72_ = p1
                rhs73_ = p0
                lhs65_ = d_391_v1_
                lhs66_ = default__.safeIndex(887, (d_391_v1_).length(0))
                lhs67_ = d_443_v41_
                lhs68_ = default__.safeIndex(18, (d_443_v41_).length(0))
                lhs69_ = self
                lhs70_ = globalState
                lhs65_[lhs66_] = rhs69_
                lhs67_[lhs68_] = rhs70_
                lhs69_.f25 = rhs71_
                r0 = rhs72_
                lhs70_.f18 = rhs73_
            d_452_v46_: _dafny.Array
            nw80_ = _dafny.Array(None, 10)
            nw80_[int(0)] = _dafny.CodePoint('u')
            nw80_[int(1)] = _dafny.CodePoint('e')
            nw80_[int(2)] = (self).f24
            nw80_[int(3)] = (self).f24
            nw80_[int(4)] = default__.fm10(p0, p0, p0, p0, globalState)
            nw80_[int(5)] = (self).f24
            nw80_[int(6)] = (self).f24
            nw80_[int(7)] = ((self).f24 if p0 else (self).f24)
            nw80_[int(8)] = (self).f24
            nw80_[int(9)] = (self).f24
            d_452_v46_ = nw80_
            d_452_v46_ = d_452_v46_
            (globalState).f18 = not(True)
            d_453_v47_: D7
            d_453_v47_ = D7_DC14(p0, p1, p1)
            (globalState).f18 = (d_453_v47_).cf20
            d_454_v48_: D2
            d_454_v48_ = D2_DC2(p0)
            d_454_v48_ = d_454_v48_
        elif True:
            d_455_v49_: C0
            nw81_ = C0()
            nw81_.ctor__()
            d_455_v49_ = nw81_
            (globalState).f18 = p0
            d_456_v50_: _dafny.Seq
            d_456_v50_ = _dafny.SeqWithoutIsStrInference([p0])
            d_457_v51_: _dafny.Map
            d_457_v51_ = _dafny.Map({(0) - (p1): d_456_v50_})
            d_458_v52_: _dafny.Seq
            d_458_v52_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fwhru"))
            rhs74_ = not((_dafny.MultiSet([p1, p1])).isdisjoint(_dafny.MultiSet(d_430_v31_)))
            rhs75_ = len(d_458_v52_)
            rhs76_ = d_457_v51_
            lhs71_ = globalState
            lhs72_ = globalState
            lhs71_.f18 = rhs74_
            lhs72_.f11 = rhs75_
            d_457_v51_ = rhs76_
            d_459_v53_: _dafny.MultiSet
            d_459_v53_ = _dafny.MultiSet([not(p0)])
            rhs77_ = (self.f25) + (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([self.f25, -356]))])))
            rhs78_ = p1
            rhs79_ = (default__.fm8(self.f25, p0, d_456_v50_, globalState)).issubset((d_459_v53_) | (d_459_v53_))
            lhs73_ = globalState
            lhs74_ = globalState
            lhs75_ = globalState
            lhs73_.f22 = rhs77_
            lhs74_.f16 = rhs78_
            lhs75_.f18 = rhs79_
            if p0:
                d_460_v54_: _dafny.Array
                nw82_ = _dafny.Array(None, 22)
                nw82_[int(0)] = d_431_v32_
                nw82_[int(1)] = (_dafny.Map({self.f25: p1})) | ((d_431_v32_).set((self).fm6(globalState), p1))
                nw82_[int(2)] = (d_431_v32_) | (d_431_v32_)
                nw82_[int(3)] = d_431_v32_
                nw82_[int(4)] = d_431_v32_
                nw82_[int(5)] = d_431_v32_
                nw82_[int(6)] = d_431_v32_
                nw82_[int(7)] = d_431_v32_
                nw82_[int(8)] = (d_431_v32_) | (d_431_v32_)
                nw82_[int(9)] = d_431_v32_
                nw82_[int(10)] = (_dafny.Map({p1: p1})) | (d_431_v32_)
                nw82_[int(11)] = _dafny.Map({p1: self.f25})
                nw82_[int(12)] = d_431_v32_
                nw82_[int(13)] = d_431_v32_
                nw82_[int(14)] = (d_431_v32_) | (_dafny.Map({(self).fm6(globalState): len(_dafny.Map({self.f25: d_391_v1_}))}))
                nw82_[int(15)] = (default__.fm11(_dafny.CodePoint('c'), p0, globalState)) | (d_431_v32_)
                nw82_[int(16)] = _dafny.Map({self.f25: self.f25})
                nw82_[int(17)] = d_431_v32_
                nw82_[int(18)] = (d_431_v32_ if True else (d_431_v32_).set(p1, len(_dafny.Set({p0, p0, p0}))))
                nw82_[int(19)] = d_431_v32_
                nw82_[int(20)] = d_431_v32_
                nw82_[int(21)] = _dafny.Map({p1: p1})
                d_460_v54_ = nw82_
                index76_ = default__.safeIndex(7, (d_460_v54_).length(0))
                (d_460_v54_)[index76_] = _dafny.Map({p1: p1})
                d_461_v55_: _dafny.Map
                d_461_v55_ = _dafny.Map({self.f25: (self).f24})
                d_462_v56_: _dafny.Map
                d_462_v56_ = _dafny.Map({p1: _dafny.Map({(self).f24: True})})
                index77_ = default__.safeIndex(7, (d_460_v54_).length(0))
                rhs80_ = default__.fm11(((d_461_v55_)[p1] if (p1) in (d_461_v55_) else _dafny.CodePoint('t')), p0, globalState)
                rhs81_ = (len(_dafny.Map({len(d_462_v56_): (self).fm6(globalState)})) if p0 else (248) - (281))
                rhs82_ = p0
                rhs83_ = d_455_v49_
                lhs76_ = d_460_v54_
                lhs77_ = default__.safeIndex(7, (d_460_v54_).length(0))
                lhs78_ = globalState
                lhs79_ = globalState
                lhs76_[lhs77_] = rhs80_
                lhs78_.f22 = rhs81_
                lhs79_.f18 = rhs82_
                d_455_v49_ = rhs83_
                (globalState).f18 = p0
                (globalState).f18 = p0
                d_463_v57_: _dafny.Map
                d_463_v57_ = _dafny.Map({(p1) - (p1): d_390_v0_})
                d_463_v57_ = d_463_v57_
                (globalState).f18 = p0
            elif True:
                index78_ = default__.safeIndex(118, (d_391_v1_).length(0))
                (d_391_v1_)[index78_] = (self.f25) <= (p1)
                index79_ = default__.safeIndex(118, (d_391_v1_).length(0))
                (d_391_v1_)[index79_] = (d_456_v50_)[default__.safeIndex(len((D3_DC5(d_456_v50_, -73, -879)).cf6), len(d_456_v50_))]
                (globalState).f0 = default__.safeDivisionInt((self.f25 if p0 else (self).fm6(globalState)), default__.safeDivisionInt(p1, p1))
                (globalState).f11 = self.f25
                d_464_v58_: _dafny.Seq
                d_464_v58_ = _dafny.SeqWithoutIsStrInference([d_429_v30_])
                d_465_v59_: _dafny.Array
                nw83_ = _dafny.Array(None, 20)
                nw83_[int(0)] = d_429_v30_
                nw83_[int(1)] = (d_429_v30_) - (d_429_v30_)
                nw83_[int(2)] = (d_429_v30_) - (_dafny.Set({(d_391_v1_)[default__.safeIndex(118, (d_391_v1_).length(0))], p0}))
                nw83_[int(3)] = _dafny.Set({not((d_391_v1_)[default__.safeIndex(118, (d_391_v1_).length(0))])})
                nw83_[int(4)] = _dafny.Set({(d_391_v1_)[default__.safeIndex(118, (d_391_v1_).length(0))]})
                nw83_[int(5)] = d_429_v30_
                nw83_[int(6)] = (d_464_v58_)[default__.safeIndex((_dafny.MultiSet([len(d_430_v31_), p1])).cardinality, len(d_464_v58_))]
                nw83_[int(7)] = d_429_v30_
                nw83_[int(8)] = d_429_v30_
                nw83_[int(9)] = d_429_v30_
                nw83_[int(10)] = _dafny.Set({(d_391_v1_)[default__.safeIndex(118, (d_391_v1_).length(0))], (d_391_v1_)[default__.safeIndex(118, (d_391_v1_).length(0))]})
                nw83_[int(11)] = (_dafny.Set({(d_391_v1_)[default__.safeIndex(118, (d_391_v1_).length(0))], p0})).intersection(d_429_v30_)
                nw83_[int(12)] = d_429_v30_
                nw83_[int(13)] = d_429_v30_
                nw83_[int(14)] = (d_429_v30_) | (d_429_v30_)
                nw83_[int(15)] = d_429_v30_
                nw83_[int(16)] = d_429_v30_
                nw83_[int(17)] = d_429_v30_
                nw83_[int(18)] = d_429_v30_
                nw83_[int(19)] = d_429_v30_
                d_465_v59_ = nw83_
                index80_ = default__.safeIndex(772, (d_465_v59_).length(0))
                (d_465_v59_)[index80_] = (d_429_v30_) - (d_429_v30_)
                index81_ = default__.safeIndex(772, (d_465_v59_).length(0))
                rhs84_ = (default__.fm12((d_391_v1_)[default__.safeIndex(118, (d_391_v1_).length(0))], p0, True, globalState)) | ((d_429_v30_).intersection(d_429_v30_))
                rhs85_ = (d_391_v1_)[default__.safeIndex(118, (d_391_v1_).length(0))]
                lhs80_ = d_465_v59_
                lhs81_ = default__.safeIndex(772, (d_465_v59_).length(0))
                lhs82_ = globalState
                lhs80_[lhs81_] = rhs84_
                lhs82_.f18 = rhs85_
                index82_ = default__.safeIndex(118, (d_391_v1_).length(0))
                (d_391_v1_)[index82_] = default__.fm0(default__.safeDivisionInt(self.f25, len(d_458_v52_)), (658) in (_dafny.Map({p1: self.f25})), not((d_429_v30_).issubset((d_465_v59_)[default__.safeIndex(772, (d_465_v59_).length(0))])), (0) - (p1), globalState)
        r0 = default__.safeModuloInt(self.f25, p1)
        r1 = (((self).fm6(globalState) if True else p1)) - ((p1 if default__.fm0(p1, p0, p0, self.f25, globalState) else self.f25))
        r2 = (0) - ((0) - (default__.safeDivisionInt(p1, self.f25)))
        r3 = self.f25
        return r0, r1, r2, r3


class C2:
    def  __init__(self):
        self._f27: int = int(0)
        self._f28: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    def ctor__(self, f27, f28):
        (self)._f27 = f27
        (self)._f28 = f28

    def fm4(self, p0, p1, p2, globalState):
        return (_dafny.MultiSet([(self).f28, (self).f28, (self).f28, (self).f28, (self).f28])).intersection((_dafny.MultiSet([(self).f28, (self).f28, True, (self).f28])) | (_dafny.MultiSet([(self).f28])))

    def m3(self, p0, p1, globalState):
        r0: int = int(0)
        r1: int = int(0)
        d_466_v0_: str
        d_466_v0_ = _dafny.CodePoint('m')
        d_467_v1_: _dafny.Seq
        d_467_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gxn"))
        d_468_v2_: _dafny.Seq
        d_468_v2_ = _dafny.SeqWithoutIsStrInference([d_466_v0_ for d_469_i0_ in range(default__.abs(269))])
        d_470_v3_: _dafny.Array
        nw84_ = _dafny.Array(None, 12)
        nw84_[int(0)] = d_467_v1_
        nw84_[int(1)] = d_468_v2_
        nw84_[int(2)] = d_468_v2_
        nw84_[int(3)] = d_467_v1_
        nw84_[int(4)] = d_468_v2_
        nw84_[int(5)] = d_468_v2_
        nw84_[int(6)] = d_468_v2_
        nw84_[int(7)] = d_468_v2_
        nw84_[int(8)] = d_467_v1_
        nw84_[int(9)] = d_468_v2_
        nw84_[int(10)] = d_468_v2_
        nw84_[int(11)] = d_468_v2_
        d_470_v3_ = nw84_
        d_471_v4_: _dafny.Map
        d_471_v4_ = _dafny.Map({((self).f27) + ((self).f27): d_470_v3_})
        d_472_v5_: _dafny.Set
        d_472_v5_ = _dafny.Set({(self).f27, (self).f27, (self).f27, -942})
        rhs86_ = (self).f28
        rhs87_ = _dafny.CodePoint('i')
        rhs88_ = d_467_v1_
        rhs89_ = ((d_471_v4_)[len(d_472_v5_)] if (len(d_472_v5_)) in (d_471_v4_) else d_470_v3_)
        lhs83_ = globalState
        lhs84_ = globalState
        lhs83_.f18 = rhs86_
        d_466_v0_ = rhs87_
        lhs84_.f3 = rhs88_
        d_470_v3_ = rhs89_
        d_473_v6_: _dafny.MultiSet
        d_473_v6_ = _dafny.MultiSet([p0, (self).f28])
        d_474_v7_: D3
        d_474_v7_ = D3_DC5(_dafny.SeqWithoutIsStrInference([(self).f28]), (d_473_v6_).cardinality, (self).f27)
        d_475_v8_: _dafny.Seq
        d_475_v8_ = (d_474_v7_).cf6
        d_476_v9_: _dafny.Seq
        d_476_v9_ = _dafny.SeqWithoutIsStrInference([(self).f28])
        source8_ = d_476_v9_
        d_477___mcc_h0_ = source8_
        d_478_cf0_ = d_477___mcc_h0_
        if (d_467_v1_) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))):
            d_479_v10_: _dafny.Array
            nw85_ = _dafny.Array(_dafny.Set({}), 5)
            d_479_v10_ = nw85_
            d_479_v10_ = d_479_v10_
            d_480_v11_: _dafny.Map
            d_480_v11_ = _dafny.Map({d_467_v1_: (d_474_v7_).cf7})
            d_480_v11_ = (d_480_v11_).set((d_467_v1_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_481_i1_ in range(default__.abs(59))])), (0) - (((self).f27) * ((self).f27)))
            (globalState).f3 = d_467_v1_
            d_482_v12_: _dafny.Array
            nw86_ = _dafny.Array(_dafny.Map({}), 18)
            d_482_v12_ = nw86_
            d_483_v13_: D4
            d_483_v13_ = D4_DC7(d_482_v12_)
            d_482_v12_ = (d_483_v13_).cf12
            d_484_v14_: _dafny.Set
            d_484_v14_ = _dafny.Set({(self).f28, p0, (not(p1)) and (p1), ((self).f28) == (p0), (self).f28})
            (globalState).f22 = len(d_484_v14_)
        elif True:
            (globalState).f0 = (self).f27
            d_485_v15_: C0
            nw87_ = C0()
            nw87_.ctor__()
            d_485_v15_ = nw87_
            d_473_v6_ = (_dafny.MultiSet((d_476_v9_) + ((D3_DC5(d_478_cf0_, (self).f27, (self).f27)).cf6))).intersection(d_473_v6_)
            d_486_v16_: _dafny.Array
            nw88_ = _dafny.Array(int(0), 1)
            d_486_v16_ = nw88_
            d_487_v17_: D5
            d_487_v17_ = D5_DC9(d_486_v16_)
            d_488_v18_: _dafny.Seq
            d_488_v18_ = _dafny.SeqWithoutIsStrInference([(d_487_v17_).cf13])
            (globalState).f11 = len(d_488_v18_)
            d_489_v19_: D5
            d_489_v19_ = D5_DC10(p1, (d_478_cf0_)[default__.safeIndex((self).f27, len(d_478_cf0_))])
            d_490_v20_: _dafny.Array
            nw89_ = _dafny.Array(None, 19)
            nw89_[int(0)] = p1
            nw89_[int(1)] = not(((self).f28) or (p0))
            nw89_[int(2)] = not(p0)
            nw89_[int(3)] = ((self).f28 if (d_478_cf0_)[default__.safeIndex((self).f27, len(d_478_cf0_))] else not((self).f28))
            nw89_[int(4)] = not((d_472_v5_).isdisjoint(_dafny.Set({(self).f27})))
            nw89_[int(5)] = (_dafny.MultiSet([not((self).f28), (d_478_cf0_)[default__.safeIndex((self).f27, len(d_478_cf0_))]])) == (d_473_v6_)
            nw89_[int(6)] = (self).f28
            nw89_[int(7)] = not (p1) or (p1)
            nw89_[int(8)] = (self).f28
            nw89_[int(9)] = (d_489_v19_).cf14
            nw89_[int(10)] = (self).f28
            nw89_[int(11)] = p1
            nw89_[int(12)] = (self).f28
            nw89_[int(13)] = default__.fm0((self).f27, p1, (self).f28, (self).f27, globalState)
            nw89_[int(14)] = p1
            nw89_[int(15)] = not(p0)
            nw89_[int(16)] = ((self).f27) > ((self).f27)
            nw89_[int(17)] = p1
            nw89_[int(18)] = (self).f28
            d_490_v20_ = nw89_
            index83_ = default__.safeIndex(602, (d_490_v20_).length(0))
            (d_490_v20_)[index83_] = ((self).f27) > ((self).f27)
            index84_ = default__.safeIndex(602, (d_490_v20_).length(0))
            (d_490_v20_)[index84_] = (p0 if (self).f28 else p0)
        (globalState).f19 = _dafny.SeqWithoutIsStrInference([(self).f27, (self).f27])
        d_491_v21_: _dafny.Array
        nw90_ = _dafny.Array(_dafny.CodePoint('D'), 20)
        d_491_v21_ = nw90_
        def lambda24_(d_492_v0_):
            def lambda25_(d_493_i2_):
                return d_492_v0_

            return lambda25_

        init13_ = lambda24_(d_466_v0_)
        nw91_ = _dafny.Array(None, 3)
        for i0_13_ in range(nw91_.length(0)):
            nw91_[i0_13_] = init13_(i0_13_)
        d_491_v21_ = nw91_
        d_494_v22_: D3
        d_494_v22_ = D3_DC6(617, (0) - ((self).f27), p0)
        d_495_v23_: _dafny.Map
        d_495_v23_ = _dafny.Map({(self).f28: p1})
        d_496_v24_: _dafny.Map
        d_496_v24_ = _dafny.Map({p0: 381})
        d_497_v25_: _dafny.MultiSet
        d_497_v25_ = _dafny.MultiSet([(d_494_v22_).cf10, (self).f27, len(d_495_v23_), len((d_472_v5_) | (d_472_v5_)), len(((d_496_v24_).set(p1, (self).f27)) | ((d_496_v24_).set(p0, (self).f27)))])
        d_498_v26_: _dafny.Seq
        d_498_v26_ = _dafny.SeqWithoutIsStrInference([d_491_v21_])
        d_499_v27_: _dafny.Seq
        d_499_v27_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({p0: (0) - ((self).f27)}), _dafny.Map({(self).f28: (self).f27}), d_496_v24_])
        rhs90_ = ((d_497_v25_) | (d_497_v25_)) | ((d_497_v25_).intersection(_dafny.MultiSet([((d_497_v25_)[(self).f27] if ((self).f27) in (d_497_v25_) else (self).f27)])))
        rhs91_ = _dafny.SeqWithoutIsStrInference([d_491_v21_])
        rhs92_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))).set(default__.safeIndex(len((d_499_v27_) + (_dafny.SeqWithoutIsStrInference([d_496_v24_, d_496_v24_]))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o")))), d_466_v0_)
        rhs93_ = not(p1)
        lhs85_ = globalState
        lhs86_ = globalState
        d_497_v25_ = rhs90_
        d_498_v26_ = rhs91_
        lhs85_.f3 = rhs92_
        lhs86_.f18 = rhs93_
        d_500_v28_: _dafny.Map
        d_500_v28_ = _dafny.Map({p1: p1})
        d_501_v29_: _dafny.Set
        d_501_v29_ = _dafny.Set({d_500_v28_})
        if not(not((d_501_v29_).ispropersubset(d_501_v29_))):
            d_502_v30_: _dafny.MultiSet
            d_502_v30_ = _dafny.MultiSet([(0) - ((self).f27), -840])
            d_503_v31_: _dafny.Seq
            d_503_v31_ = _dafny.SeqWithoutIsStrInference([d_472_v5_])
            d_504_v32_: _dafny.Array
            nw92_ = _dafny.Array(None, 24)
            nw92_[int(0)] = (d_502_v30_).cardinality
            nw92_[int(1)] = len(_dafny.Map({(self).f27: 872}))
            nw92_[int(2)] = (self).f27
            nw92_[int(3)] = len(d_467_v1_)
            nw92_[int(4)] = (self).f27
            nw92_[int(5)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l")))
            nw92_[int(6)] = (self).f27
            nw92_[int(7)] = (0) - ((self).f27)
            nw92_[int(8)] = (self).f27
            nw92_[int(9)] = (self).f27
            nw92_[int(10)] = (self).f27
            nw92_[int(11)] = (self).f27
            nw92_[int(12)] = (self).f27
            nw92_[int(13)] = (self).f27
            nw92_[int(14)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xrr")))
            nw92_[int(15)] = (self).f27
            nw92_[int(16)] = (self).f27
            nw92_[int(17)] = (self).f27
            nw92_[int(18)] = (self).f27
            nw92_[int(19)] = len(d_472_v5_)
            nw92_[int(20)] = len(d_503_v31_)
            nw92_[int(21)] = (self).f27
            nw92_[int(22)] = (self).f27
            nw92_[int(23)] = (self).f27
            d_504_v32_ = nw92_
            d_505_v33_: D5
            d_505_v33_ = D5_DC9(d_504_v32_)
            source9_ = d_505_v33_
            if source9_.is_DC10:
                d_506___mcc_h1_ = source9_.cf14
                d_507___mcc_h2_ = source9_.cf15
                d_508_cf15_ = d_507___mcc_h2_
                d_509_cf14_ = d_506___mcc_h1_
                d_510_v34_: _dafny.Set
                d_510_v34_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ofy")), d_467_v1_})
                d_511_v35_: D3
                d_511_v35_ = D3_DC6(((self).f27 if p0 else (self).f27), len((d_510_v34_) | (d_510_v34_)), True)
                d_512_v36_: _dafny.Array
                nw93_ = _dafny.Array(_dafny.Array(None, 0), 5)
                d_512_v36_ = nw93_
                d_513_v37_: _dafny.Array
                def lambda26_(d_514_v30_):
                    def lambda27_(d_515_i3_):
                        return d_514_v30_

                    return lambda27_

                init14_ = lambda26_(d_502_v30_)
                nw94_ = _dafny.Array(None, 27)
                for i0_14_ in range(nw94_.length(0)):
                    nw94_[i0_14_] = init14_(i0_14_)
                d_513_v37_ = nw94_
                index85_ = default__.safeIndex(46, (d_512_v36_).length(0))
                (d_512_v36_)[index85_] = d_513_v37_
                index86_ = default__.safeIndex(46, (d_512_v36_).length(0))
                rhs94_ = not(p1)
                rhs95_ = d_511_v35_
                rhs96_ = d_513_v37_
                lhs87_ = globalState
                lhs88_ = d_512_v36_
                lhs89_ = default__.safeIndex(46, (d_512_v36_).length(0))
                lhs87_.f18 = rhs94_
                d_511_v35_ = rhs95_
                lhs88_[lhs89_] = rhs96_
                d_516_v38_: _dafny.Map
                d_516_v38_ = _dafny.Map({(self).f27: 200})
                (globalState).f18 = ((d_502_v30_).intersection(_dafny.MultiSet([(self).f27, (self).f27, (self).f27, (self).f27, default__.fm5((self).f27, (self).f27, globalState)]))) == (_dafny.MultiSet([(self).f27, len(d_516_v38_)]))
                d_517_v39_: _dafny.Map
                d_517_v39_ = _dafny.Map({(self).f27: p1})
                d_518_v40_: _dafny.Set
                d_518_v40_ = _dafny.Set({d_517_v39_, _dafny.Map({(self).f27: p0}), d_517_v39_, _dafny.Map({((d_516_v38_)[(self).f27] if ((self).f27) in (d_516_v38_) else (self).f27): (self).f28})})
                d_519_v41_: D3
                d_519_v41_ = D3_DC4(d_518_v40_)
                d_520_v42_: _dafny.Map
                d_520_v42_ = _dafny.Map({(self).f27: d_519_v41_})
                d_520_v42_ = (d_520_v42_).set((self).f27, D3_DC4(d_518_v40_))
                d_521_v43_: _dafny.MultiSet
                d_521_v43_ = _dafny.MultiSet([d_504_v32_, d_504_v32_, d_504_v32_])
                rhs97_ = d_509_cf14_
                rhs98_ = not(p1)
                rhs99_ = p0
                rhs100_ = _dafny.MultiSet([(self).f27, (self).f27, ((d_521_v43_).intersection(d_521_v43_)).cardinality])
                lhs90_ = globalState
                lhs91_ = globalState
                lhs90_.f18 = rhs97_
                d_509_cf14_ = rhs98_
                lhs91_.f18 = rhs99_
                d_502_v30_ = rhs100_
            elif True:
                d_522___mcc_h3_ = source9_.cf13
                d_523_cf13_ = d_522___mcc_h3_
                d_468_v2_ = d_468_v2_
                d_524_v44_: _dafny.Seq
                d_524_v44_ = _dafny.SeqWithoutIsStrInference([len(d_476_v9_)])
                d_525_v45_: _dafny.Map
                d_525_v45_ = _dafny.Map({d_524_v44_: (len(d_467_v1_)) >= ((self).f27)})
                d_525_v45_ = (d_525_v45_).set((d_524_v44_) + (d_524_v44_), True)
                (globalState).f22 = 409
                (globalState).f0 = (self).f27
            (globalState).f15 = (len(d_467_v1_)) + (((self).f27) * ((self).f27))
            d_526_v46_: _dafny.Seq
            d_526_v46_ = _dafny.SeqWithoutIsStrInference([(self).f27, (self).f27])
            d_527_v47_: D2
            d_527_v47_ = D2_DC2((p0 if p1 else default__.fm0(len(d_526_v46_), p0, p0, (self).f27, globalState)))
            source10_ = d_527_v47_
            if source10_.is_DC3:
                d_528___mcc_h4_ = source10_.cf3
                d_529___mcc_h5_ = source10_.cf4
                d_530_cf4_ = d_529___mcc_h5_
                d_531_cf3_ = d_528___mcc_h4_
                d_532_v48_: _dafny.Map
                d_532_v48_ = _dafny.Map({p1: d_504_v32_})
                d_532_v48_ = (d_532_v48_).set(p1, d_504_v32_)
                d_533_v49_: C0
                nw95_ = C0()
                nw95_.ctor__()
                d_533_v49_ = nw95_
                index87_ = default__.safeIndex(89, (d_504_v32_).length(0))
                (d_504_v32_)[index87_] = (d_531_cf3_) + (len(_dafny.Set({False})))
                index88_ = default__.safeIndex(89, (d_504_v32_).length(0))
                (d_504_v32_)[index88_] = d_530_cf4_
                d_534_v50_: _dafny.Seq
                d_534_v50_ = _dafny.SeqWithoutIsStrInference([d_467_v1_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sjtnwy")), d_467_v1_])
                index89_ = default__.safeIndex(89, (d_504_v32_).length(0))
                rhs101_ = (self).f27
                rhs102_ = (d_504_v32_)[default__.safeIndex(89, (d_504_v32_).length(0))]
                rhs103_ = (self).f27
                rhs104_ = d_466_v0_
                rhs105_ = (d_534_v50_)[default__.safeIndex(default__.safeModuloInt(len(d_526_v46_), 682), len(d_534_v50_))]
                lhs92_ = globalState
                lhs93_ = d_504_v32_
                lhs94_ = default__.safeIndex(89, (d_504_v32_).length(0))
                lhs95_ = globalState
                lhs96_ = globalState
                lhs92_.f16 = rhs101_
                lhs93_[lhs94_] = rhs102_
                lhs95_.f22 = rhs103_
                d_466_v0_ = rhs104_
                lhs96_.f3 = rhs105_
            elif True:
                d_535___mcc_h6_ = source10_.cf2
                d_536_cf2_ = d_535___mcc_h6_
                (globalState).f16 = (self).f27
                d_537_v51_: _dafny.Map
                d_537_v51_ = _dafny.Map({d_467_v1_: p1})
                (globalState).f18 = ((d_537_v51_)[((d_467_v1_) + (d_467_v1_)).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mffkpiqe"))), len((d_467_v1_) + (d_467_v1_))), d_466_v0_)] if (((d_467_v1_) + (d_467_v1_)).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mffkpiqe"))), len((d_467_v1_) + (d_467_v1_))), d_466_v0_)) in (d_537_v51_) else (self).f28)
                d_538_v52_: _dafny.Array
                def lambda28_(d_539_i4_):
                    return ((self).f27) != (73)

                init15_ = lambda28_
                nw96_ = _dafny.Array(None, 28)
                for i0_15_ in range(nw96_.length(0)):
                    nw96_[i0_15_] = init15_(i0_15_)
                d_538_v52_ = nw96_
                d_538_v52_ = d_538_v52_
                d_540_v53_: T0
                nw97_ = C1()
                nw97_.ctor__(_dafny.CodePoint('a'), (self).f27)
                d_540_v53_ = nw97_
                d_541_v54_: D6
                d_541_v54_ = D6_DC11(d_540_v53_)
                d_542_v55_: _dafny.Seq
                d_542_v55_ = _dafny.SeqWithoutIsStrInference([d_540_v53_])
                d_543_v56_: _dafny.Seq
                d_543_v56_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_541_v54_).cf16, d_540_v53_]), d_542_v55_, d_542_v55_, d_542_v55_, d_542_v55_])
                d_544_v57_: _dafny.MultiSet
                d_544_v57_ = _dafny.MultiSet([d_542_v55_, d_542_v55_, d_542_v55_])
                d_545_v58_: _dafny.Array
                nw98_ = _dafny.Array(None, 10)
                nw98_[int(0)] = _dafny.MultiSet(d_543_v56_)
                nw98_[int(1)] = d_544_v57_
                nw98_[int(2)] = d_544_v57_
                nw98_[int(3)] = d_544_v57_
                nw98_[int(4)] = d_544_v57_
                nw98_[int(5)] = _dafny.MultiSet(d_543_v56_)
                nw98_[int(6)] = d_544_v57_
                nw98_[int(7)] = (d_544_v57_ if p1 else _dafny.MultiSet([(d_542_v55_).set(default__.safeIndex(len(d_501_v29_), len(d_542_v55_)), d_540_v53_), d_542_v55_, _dafny.SeqWithoutIsStrInference([d_540_v53_, d_540_v53_]), d_542_v55_, d_542_v55_]))
                nw98_[int(8)] = (_dafny.MultiSet(d_543_v56_)).set(d_542_v55_, default__.abs(d_540_v53_.f25))
                nw98_[int(9)] = d_544_v57_
                d_545_v58_ = nw98_
                d_546_v59_: _dafny.Seq
                d_546_v59_ = d_542_v55_
                index90_ = default__.safeIndex(844, (d_545_v58_).length(0))
                (d_545_v58_)[index90_] = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_540_v53_]), (d_546_v59_), d_542_v55_])
                index91_ = default__.safeIndex(844, (d_545_v58_).length(0))
                rhs106_ = d_544_v57_
                rhs107_ = d_473_v6_
                lhs97_ = d_545_v58_
                lhs98_ = default__.safeIndex(844, (d_545_v58_).length(0))
                lhs97_[lhs98_] = rhs106_
                d_473_v6_ = rhs107_
            d_547_v60_: T0
            nw99_ = C1()
            nw99_.ctor__(d_466_v0_, (self).f27)
            d_547_v60_ = nw99_
            d_548_v61_: _dafny.Map
            d_548_v61_ = _dafny.Map({d_547_v60_: p0})
            d_549_v62_: D6
            d_549_v62_ = D6_DC12(((d_548_v61_)[d_547_v60_] if (d_547_v60_) in (d_548_v61_) else not((self).f28)), default__.fm5(650, (0) - (default__.fm5(d_547_v60_.f25, d_547_v60_.f25, globalState)), globalState))
            (globalState).f18 = (d_549_v62_).cf17
            (globalState).f18 = ((0) - (((self).f27) * (d_547_v60_.f25))) <= (((self).f27) - ((self).f27))
        elif True:
            (globalState).f18 = not(((d_500_v28_) != (d_500_v28_) if (self).f28 else (p1) or (not((self).f28))))
            (globalState).f18 = not(p1)
            if not (False) or ((self).f28):
                d_550_v63_: C0
                nw100_ = C0()
                nw100_.ctor__()
                d_550_v63_ = nw100_
                d_551_v64_: D4
                d_551_v64_ = D4_DC8()
                d_552_v65_: _dafny.Map
                d_552_v65_ = _dafny.Map({(self).f27: default__.fm0(-11, p0, p1, (self).f27, globalState)})
                d_553_v66_: _dafny.MultiSet
                d_553_v66_ = _dafny.MultiSet([len(d_552_v65_)])
                d_554_v67_: _dafny.Map
                d_554_v67_ = _dafny.Map({(d_553_v66_) | (d_553_v66_): p1})
                d_555_v68_: _dafny.Seq
                d_555_v68_ = _dafny.SeqWithoutIsStrInference([(self).f27])
                rhs108_ = D4_DC8()
                rhs109_ = d_554_v67_
                rhs110_ = (((self).f27) - (337)) + (default__.safeDivisionInt(len(d_555_v68_), (self).f27))
                lhs99_ = globalState
                d_551_v64_ = rhs108_
                d_554_v67_ = rhs109_
                lhs99_.f15 = rhs110_
                (globalState).f16 = (self).f27
                d_556_v69_: _dafny.Array
                nw101_ = _dafny.Array(None, 5)
                nw101_[int(0)] = d_466_v0_
                nw101_[int(1)] = default__.fm10((self).f28, p0, p0, p1, globalState)
                nw101_[int(2)] = d_466_v0_
                nw101_[int(3)] = d_466_v0_
                nw101_[int(4)] = d_466_v0_
                d_556_v69_ = nw101_
                d_557_v70_: _dafny.Map
                d_557_v70_ = _dafny.Map({(self).f27: d_556_v69_})
                d_557_v70_ = (d_557_v70_).set((self).f27, d_556_v69_)
                (globalState).f18 = ((d_555_v68_)[default__.safeIndex((self).f27, len(d_555_v68_))]) >= ((self).f27)
            elif True:
                d_558_v71_: T0
                nw102_ = C1()
                nw102_.ctor__(d_466_v0_, (d_473_v6_).cardinality)
                d_558_v71_ = nw102_
                d_559_v72_: _dafny.Map
                d_559_v72_ = _dafny.Map({default__.fm5((self).f27, (self).f27, globalState): d_558_v71_})
                d_560_v73_: D7
                d_560_v73_ = D7_DC14(p0, len(d_559_v72_), d_558_v71_.f25)
                d_561_v74_: _dafny.Map
                d_561_v74_ = _dafny.Map({p1: D3_DC6((self).f27, (self).f27, (d_560_v73_).cf20)})
                d_562_v75_: D9
                d_562_v75_ = D9_DC16(default__.fm13(d_467_v1_, globalState))
                d_561_v74_ = (d_562_v75_).cf24
                (globalState).f2 = len(d_467_v1_)
                d_563_v76_: _dafny.Array
                nw103_ = _dafny.Array(_dafny.Seq({}), 25)
                d_563_v76_ = nw103_
                d_564_v77_: _dafny.Array
                def lambda29_(d_565_i5_):
                    return False

                init16_ = lambda29_
                nw104_ = _dafny.Array(None, 15)
                for i0_16_ in range(nw104_.length(0)):
                    nw104_[i0_16_] = init16_(i0_16_)
                d_564_v77_ = nw104_
                d_566_v78_: _dafny.Seq
                d_566_v78_ = _dafny.SeqWithoutIsStrInference([d_564_v77_])
                index92_ = default__.safeIndex(684, (d_563_v76_).length(0))
                (d_563_v76_)[index92_] = (d_566_v78_ if p1 else d_566_v78_)
                index93_ = default__.safeIndex(684, (d_563_v76_).length(0))
                rhs111_ = not(p1)
                rhs112_ = not(p0)
                rhs113_ = (self).f27
                rhs114_ = d_566_v78_
                lhs100_ = globalState
                lhs101_ = globalState
                lhs102_ = globalState
                lhs103_ = d_563_v76_
                lhs104_ = default__.safeIndex(684, (d_563_v76_).length(0))
                lhs100_.f18 = rhs111_
                lhs101_.f18 = rhs112_
                lhs102_.f22 = rhs113_
                lhs103_[lhs104_] = rhs114_
                d_567_v79_: _dafny.Map
                d_567_v79_ = _dafny.Map({(self).f27: (self).f28})
                d_568_v80_: _dafny.Map
                d_568_v80_ = _dafny.Map({(self).f28: (d_567_v79_).set(d_558_v71_.f25, p1)})
                index94_ = default__.safeIndex(241, (d_564_v77_).length(0))
                (d_564_v77_)[index94_] = (_dafny.Map({(self).f28: d_567_v79_})) != (d_568_v80_)
                d_569_v81_: D3
                d_569_v81_ = D3_DC6(len(d_567_v79_), (self).f27, p0)
                index95_ = default__.safeIndex(241, (d_564_v77_).length(0))
                rhs115_ = (-589) * ((615) - ((self).f27))
                rhs116_ = (d_569_v81_).cf9
                rhs117_ = p0
                rhs118_ = (d_558_v71_.f25) * ((self).f27)
                lhs105_ = d_558_v71_
                lhs106_ = globalState
                lhs107_ = d_564_v77_
                lhs108_ = default__.safeIndex(241, (d_564_v77_).length(0))
                lhs109_ = globalState
                lhs105_.f25 = rhs115_
                lhs106_.f0 = rhs116_
                lhs107_[lhs108_] = rhs117_
                lhs109_.f22 = rhs118_
                d_570_v82_: D7
                d_570_v82_ = D7_DC13(d_567_v79_)
                d_571_v83_: _dafny.MultiSet
                d_571_v83_ = _dafny.MultiSet([d_570_v82_, d_570_v82_])
                d_572_v84_: C1
                nw105_ = C1()
                nw105_.ctor__(d_466_v0_, ((d_571_v83_)[d_570_v82_] if (d_570_v82_) in (d_571_v83_) else len(d_467_v1_)))
                d_572_v84_ = nw105_
            d_573_v85_: _dafny.Set
            d_573_v85_ = _dafny.Set({p0, p0})
            rhs119_ = (self).f28
            rhs120_ = p1
            rhs121_ = (0) - ((self).f27)
            rhs122_ = default__.fm0((0) - (((self).f27) + ((self).f27)), p1, (d_573_v85_).isdisjoint((default__.fm14((self).f28, globalState)).cf26), (self).f27, globalState)
            lhs110_ = globalState
            lhs111_ = globalState
            lhs112_ = globalState
            lhs113_ = globalState
            lhs110_.f18 = rhs119_
            lhs111_.f18 = rhs120_
            lhs112_.f22 = rhs121_
            lhs113_.f18 = rhs122_
            d_574_v86_: _dafny.MultiSet
            d_574_v86_ = _dafny.MultiSet([d_466_v0_])
            if (default__.fm5(-135, (d_574_v86_).cardinality, globalState)) >= ((self).f27):
                (globalState).f18 = not(p0)
                (globalState).f18 = p1
                index96_ = default__.safeIndex(398, (d_470_v3_).length(0))
                (d_470_v3_)[index96_] = d_468_v2_
                index97_ = default__.safeIndex(398, (d_470_v3_).length(0))
                (d_470_v3_)[index97_] = d_468_v2_
                d_575_v87_: D5
                d_575_v87_ = D5_DC10(p1, p1)
                d_576_v88_: _dafny.Map
                d_576_v88_ = _dafny.Map({(self).f27: d_467_v1_})
                d_577_v89_: _dafny.Array
                nw106_ = _dafny.Array(None, 9)
                nw106_[int(0)] = ((d_467_v1_) + (d_467_v1_)).set(default__.safeIndex((self).f27, len((d_467_v1_) + (d_467_v1_))), d_466_v0_)
                nw106_[int(1)] = d_467_v1_
                nw106_[int(2)] = d_467_v1_
                nw106_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "htgefcx"))
                nw106_[int(4)] = ((d_470_v3_)[default__.safeIndex(398, (d_470_v3_).length(0))])
                nw106_[int(5)] = d_467_v1_
                nw106_[int(6)] = d_467_v1_
                nw106_[int(7)] = ((d_576_v88_)[(self).f27] if ((self).f27) in (d_576_v88_) else default__.fm15((self).f27, (self).f27, (0) - ((self).f27), globalState))
                nw106_[int(8)] = (d_467_v1_) + (d_467_v1_)
                d_577_v89_ = nw106_
                index98_ = default__.safeIndex(367, (d_577_v89_).length(0))
                (d_577_v89_)[index98_] = (d_467_v1_) + (d_467_v1_)
                index99_ = default__.safeIndex(367, (d_577_v89_).length(0))
                def iife50_(_pat_let8_0):
                    def iife51_(d_578_dt__update__tmp_h2_):
                        def iife52_(_pat_let9_0):
                            def iife53_(d_579_dt__update_hcf15_h0_):
                                return D5_DC10((d_578_dt__update__tmp_h2_).cf14, d_579_dt__update_hcf15_h0_)
                            return iife53_(_pat_let9_0)
                        return iife52_((self).f28)
                    return iife51_(_pat_let8_0)
                rhs123_ = iife50_(d_575_v87_)
                rhs124_ = (self).f27
                rhs125_ = p0
                rhs126_ = (d_467_v1_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_580_i6_ in range(default__.abs(-798))]))
                lhs114_ = globalState
                lhs115_ = globalState
                lhs116_ = d_577_v89_
                lhs117_ = default__.safeIndex(367, (d_577_v89_).length(0))
                d_575_v87_ = rhs123_
                lhs114_.f10 = rhs124_
                lhs115_.f18 = rhs125_
                lhs116_[lhs117_] = rhs126_
                d_581_v90_: _dafny.Map
                d_581_v90_ = _dafny.Map({(p0) or (p1): (d_476_v9_) + (d_476_v9_)})
                d_581_v90_ = (d_581_v90_).set(False, d_476_v9_)
            elif True:
                d_582_v91_: _dafny.Array
                nw107_ = _dafny.Array(None, 27)
                nw107_[int(0)] = _dafny.SeqWithoutIsStrInference([p1, p0, p1])
                nw107_[int(1)] = (d_476_v9_) + (_dafny.SeqWithoutIsStrInference([p1, p1]))
                nw107_[int(2)] = d_476_v9_
                nw107_[int(3)] = d_476_v9_
                nw107_[int(4)] = d_476_v9_
                nw107_[int(5)] = d_476_v9_
                nw107_[int(6)] = d_476_v9_
                nw107_[int(7)] = d_476_v9_
                nw107_[int(8)] = (d_476_v9_) + (default__.fm7(globalState))
                nw107_[int(9)] = _dafny.SeqWithoutIsStrInference([p0, (self).f28, p0, p0])
                nw107_[int(10)] = d_476_v9_
                nw107_[int(11)] = (d_476_v9_ if p1 else d_476_v9_)
                nw107_[int(12)] = d_476_v9_
                nw107_[int(13)] = d_476_v9_
                nw107_[int(14)] = d_476_v9_
                nw107_[int(15)] = d_476_v9_
                nw107_[int(16)] = d_476_v9_
                nw107_[int(17)] = (_dafny.SeqWithoutIsStrInference([p1])) + (d_476_v9_)
                nw107_[int(18)] = d_476_v9_
                nw107_[int(19)] = d_476_v9_
                nw107_[int(20)] = d_476_v9_
                nw107_[int(21)] = d_476_v9_
                nw107_[int(22)] = (d_476_v9_) + (d_476_v9_)
                nw107_[int(23)] = (d_476_v9_) + (d_476_v9_)
                nw107_[int(24)] = d_476_v9_
                nw107_[int(25)] = d_476_v9_
                nw107_[int(26)] = _dafny.SeqWithoutIsStrInference([p0])
                d_582_v91_ = nw107_
                index100_ = default__.safeIndex(27, (d_582_v91_).length(0))
                (d_582_v91_)[index100_] = d_476_v9_
                index101_ = default__.safeIndex(27, (d_582_v91_).length(0))
                (d_582_v91_)[index101_] = d_476_v9_
                (globalState).f18 = p1
                d_583_v92_: _dafny.Map
                d_583_v92_ = _dafny.Map({not(p0): (d_467_v1_).set(default__.safeIndex(-629, len(d_467_v1_)), d_466_v0_)})
                d_583_v92_ = (d_583_v92_).set(p1, (d_467_v1_) + (d_467_v1_))
                (globalState).f0 = (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "html"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i")))) if not((self).f28) else (len(_dafny.SeqWithoutIsStrInference([p0, (self).f28]))) * (86))
                (globalState).f18 = not((self).f28)
        d_584_v93_: _dafny.Seq
        d_584_v93_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fyp")), d_467_v1_])
        d_585_v94_: D5
        d_585_v94_ = D5_DC10(p1, False)
        if not(((((d_467_v1_) + (_dafny.SeqWithoutIsStrInference([d_466_v0_ for d_586_i7_ in range(default__.abs(280))]))).set(default__.safeIndex((self).f27, len((d_467_v1_) + (_dafny.SeqWithoutIsStrInference([d_466_v0_ for d_587_i7_ in range(default__.abs(280))])))), d_466_v0_)).set(default__.safeIndex(len((d_584_v93_)[default__.safeIndex((self).f27, len(d_584_v93_))]), len(((d_467_v1_) + (_dafny.SeqWithoutIsStrInference([d_466_v0_ for d_588_i7_ in range(default__.abs(280))]))).set(default__.safeIndex((self).f27, len((d_467_v1_) + (_dafny.SeqWithoutIsStrInference([d_466_v0_ for d_589_i7_ in range(default__.abs(280))])))), d_466_v0_))), d_466_v0_)) != ((((d_467_v1_) if (d_585_v94_).cf15 else d_467_v1_)).set(default__.safeIndex((self).f27, len(((d_467_v1_) if (d_585_v94_).cf15 else d_467_v1_))), d_466_v0_))):
            (globalState).f15 = (self).f27
            (globalState).f18 = False
            d_590_v95_: _dafny.Seq
            d_590_v95_ = _dafny.SeqWithoutIsStrInference([(self).f27])
            d_591_v96_: _dafny.Array
            nw108_ = _dafny.Array(None, 15)
            nw108_[int(0)] = p1
            nw108_[int(1)] = (self).f28
            nw108_[int(2)] = default__.fm0((self).f27, p1, p1, (self).f27, globalState)
            nw108_[int(3)] = p1
            nw108_[int(4)] = p0
            nw108_[int(5)] = not((self).f28)
            nw108_[int(6)] = p1
            nw108_[int(7)] = (self).f28
            nw108_[int(8)] = p1
            nw108_[int(9)] = p0
            nw108_[int(10)] = (self).f28
            nw108_[int(11)] = not(p1)
            nw108_[int(12)] = (self).f28
            nw108_[int(13)] = p1
            nw108_[int(14)] = (d_585_v94_).cf15
            d_591_v96_ = nw108_
            d_592_v97_: _dafny.Map
            d_592_v97_ = _dafny.Map({d_591_v96_: p1})
            d_593_v98_: _dafny.Map
            d_593_v98_ = _dafny.Map({(((d_473_v6_)[(self).f28] if ((self).f28) in (d_473_v6_) else (self).f27)) - (len(d_590_v95_)): (d_591_v96_) in (d_592_v97_)})
            d_594_v99_: _dafny.Array
            nw109_ = _dafny.Array(int(0), 20)
            d_594_v99_ = nw109_
            d_595_v100_: _dafny.Map
            d_595_v100_ = _dafny.Map({d_594_v99_: default__.fm0((self).f27, not(False), (self).f28, (self).f27, globalState)})
            d_593_v98_ = (d_593_v98_).set((self).f27, ((0) - ((self).f27)) < (len(d_595_v100_)))
            rhs127_ = _dafny.CodePoint('b')
            rhs128_ = d_467_v1_
            lhs118_ = globalState
            d_466_v0_ = rhs127_
            lhs118_.f3 = rhs128_
            if p1:
                (globalState).f11 = (0) - (default__.safeModuloInt(len(default__.fm15(len(d_467_v1_), (self).f27, default__.fm5(len(d_590_v95_), (self).f27, globalState), globalState)), ((d_473_v6_)[p0] if (p0) in (d_473_v6_) else (self).f27)))
                d_596_v101_: _dafny.Array
                nw110_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 17)
                d_596_v101_ = nw110_
                index102_ = default__.safeIndex(63, (d_596_v101_).length(0))
                (d_596_v101_)[index102_] = d_467_v1_
                index103_ = default__.safeIndex(63, (d_596_v101_).length(0))
                (d_596_v101_)[index103_] = d_467_v1_
                (globalState).f10 = (0) - ((self).f27)
                (globalState).f18 = False
                d_597_v102_: _dafny.MultiSet
                d_597_v102_ = _dafny.MultiSet([d_475_v8_, d_476_v9_, d_475_v8_])
                d_598_v103_: _dafny.MultiSet
                d_598_v103_ = _dafny.MultiSet([(self).f27, (d_597_v102_).cardinality])
                d_598_v103_ = d_598_v103_
            elif True:
                (globalState).f18 = (p1) == (p1)
                index104_ = default__.safeIndex(854, (d_594_v99_).length(0))
                (d_594_v99_)[index104_] = 213
                d_599_v104_: _dafny.MultiSet
                d_599_v104_ = _dafny.MultiSet([(self).f27])
                index105_ = default__.safeIndex(854, (d_594_v99_).length(0))
                (d_594_v99_)[index105_] = ((self).f27 if (self).f28 else ((d_599_v104_).cardinality if p1 else (self).f27))
                (globalState).f2 = default__.safeDivisionInt(((d_594_v99_)[default__.safeIndex(854, (d_594_v99_).length(0))]) - ((0) - ((d_594_v99_)[default__.safeIndex(854, (d_594_v99_).length(0))])), default__.safeModuloInt((d_594_v99_)[default__.safeIndex(854, (d_594_v99_).length(0))], 588))
                d_600_v105_: D3
                d_600_v105_ = D3_DC6(len(d_467_v1_), (d_594_v99_)[default__.safeIndex(854, (d_594_v99_).length(0))], p0)
                d_601_v106_: D9
                d_601_v106_ = D9_DC16(_dafny.Map({(self).f28: d_600_v105_}))
                d_601_v106_ = d_601_v106_
                d_602_v107_: D10
                d_602_v107_ = D10_DC20(p1, default__.fm5((self).f27, (d_594_v99_)[default__.safeIndex(854, (d_594_v99_).length(0))], globalState))
                d_602_v107_ = d_602_v107_
        elif True:
            d_500_v28_ = (d_500_v28_) | ((d_500_v28_) | (d_500_v28_))
            d_603_v108_: D2
            d_603_v108_ = D2_DC3((self).f27, (self).f27)
            d_604_v109_: _dafny.Map
            d_604_v109_ = _dafny.Map({d_603_v108_: (0) - ((self).f27)})
            d_604_v109_ = (d_604_v109_).set(d_603_v108_, 446)
            d_605_v110_: _dafny.Array
            nw111_ = _dafny.Array(False, 20)
            d_605_v110_ = nw111_
            index106_ = default__.safeIndex(331, (d_605_v110_).length(0))
            (d_605_v110_)[index106_] = p1
            index107_ = default__.safeIndex(331, (d_605_v110_).length(0))
            (d_605_v110_)[index107_] = p0
            d_606_v111_: _dafny.Seq
            d_606_v111_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([not((self).f28)]))])
            (globalState).f0 = (d_606_v111_)[default__.safeIndex(len(d_476_v9_), len(d_606_v111_))]
            d_607_v112_: _dafny.Map
            d_607_v112_ = _dafny.Map({False: (self).f27})
            (globalState).f15 = (len(d_607_v112_)) * ((self).f27)
        d_608_v113_: C1
        nw112_ = C1()
        nw112_.ctor__(d_466_v0_, ((self).f27) - ((self).f27))
        d_608_v113_ = nw112_
        d_609_i8_: int
        d_609_i8_ = 0
        with _dafny.label("2"):
            while ((d_473_v6_) - (((_dafny.MultiSet([p1])).set(p1, default__.abs((self).f27))).set(p0, default__.abs(103)))).isdisjoint(d_473_v6_):
                with _dafny.c_label("2"):
                    if (d_609_i8_) >= (100):
                        raise _dafny.Break("2")
                    d_609_i8_ = (d_609_i8_) + (1)
                    (globalState).f18 = p1
                    (globalState).f15 = len(d_476_v9_)
                    (globalState).f18 = ((d_473_v6_) - (d_473_v6_)).isdisjoint(d_473_v6_)
                    (globalState).f10 = (self).f27
                    pass
            pass
        def lambda30_(source11_):
            if source11_.is_DC3:
                d_610___mcc_h7_ = source11_.cf3
                d_611___mcc_h8_ = source11_.cf4
                d_612_cf4_ = d_611___mcc_h8_
                d_613_cf3_ = d_610___mcc_h7_
                return (self).f27
            elif True:
                d_614___mcc_h9_ = source11_.cf2
                d_615_cf2_ = d_614___mcc_h9_
                return (0) - ((self).f27)

        r0 = lambda30_(D2_DC2((self).f28))
        d_616_v114_: _dafny.MultiSet
        d_616_v114_ = _dafny.MultiSet([d_466_v0_])
        r1 = (0) - (((d_616_v114_)[d_466_v0_] if (d_466_v0_) in (d_616_v114_) else (self).f27))
        return r0, r1

    @property
    def f27(self):
        return self._f27
    @property
    def f28(self):
        return self._f28

class C3(T1, T0):
    def  __init__(self):
        self._f25: int = int(0)
        self._f24: str = _dafny.CodePoint('D')
        self.f31: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f25(self):
        return self._f25
    @f25.setter
    def f25(self, value):
        self._f25 = value
    @property
    def f24(self):
        return self._f24
    def ctor__(self, f31, f24, f25):
        (self).f31 = f31
        (self)._f24 = f24
        (self).f25 = f25

    def fm3(self, p0, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pcxk"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_617_i0_ in range(default__.abs(893))]))

    def fm18(self, globalState):
        return not(self.f31)


class C4(T0, T1):
    def  __init__(self):
        self._f25: int = int(0)
        self._f24: str = _dafny.CodePoint('D')
        self.f29: _dafny.Map = _dafny.Map({})
        self._f30: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f25(self):
        return self._f25
    @f25.setter
    def f25(self, value):
        self._f25 = value
    @property
    def f24(self):
        return self._f24
    def ctor__(self, f29, f30, f24, f25):
        (self).f29 = f29
        (self)._f30 = f30
        (self)._f24 = f24
        (self).f25 = f25

    def fm3(self, p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "if"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kaxkef")))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oqanv"))))

    def m5(self, p0, p1, p2, globalState):
        (globalState).f11 = p2
        d_618_i0_: int
        d_618_i0_ = 0
        with _dafny.label("3"):
            while p1:
                with _dafny.c_label("3"):
                    if (d_618_i0_) >= (100):
                        raise _dafny.Break("3")
                    d_618_i0_ = (d_618_i0_) + (1)
                    d_619_v0_: _dafny.Map
                    d_619_v0_ = _dafny.Map({D2_DC2(p1): p2})
                    d_620_v1_: _dafny.Map
                    d_620_v1_ = _dafny.Map({p2: (0) - ((default__.fm16(d_619_v0_, globalState)).cf29)})
                    d_621_v2_: _dafny.Seq
                    d_621_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aqdbq"))
                    d_622_v3_: _dafny.Map
                    d_622_v3_ = _dafny.Map({d_621_v2_: default__.fm0(p2, (self).f30, p1, len(_dafny.SeqWithoutIsStrInference([self.f25 for d_623_i1_ in range(default__.abs(644))])), globalState)})
                    d_620_v1_ = (d_620_v1_).set(len((d_622_v3_) | (d_622_v3_)), p2)
                    d_624_v4_: _dafny.Array
                    nw113_ = _dafny.Array(None, 24)
                    d_624_v4_ = nw113_
                    d_625_v5_: _dafny.Map
                    d_625_v5_ = _dafny.Map({(self).f30: d_624_v4_})
                    d_624_v4_ = ((d_625_v5_)[(self).f30] if ((self).f30) in (d_625_v5_) else d_624_v4_)
                    d_620_v1_ = (d_620_v1_).set(self.f25, default__.safeDivisionInt(self.f25, (0) - (p2)))
                    d_626_v6_: _dafny.Seq
                    d_626_v6_ = _dafny.SeqWithoutIsStrInference([(self).f30])
                    d_627_v7_: _dafny.Array
                    def lambda31_(d_628_v1_):
                        def lambda32_(d_629_i2_):
                            return d_628_v1_

                        return lambda32_

                    init17_ = lambda31_(d_620_v1_)
                    nw114_ = _dafny.Array(None, 16)
                    for i0_17_ in range(nw114_.length(0)):
                        nw114_[i0_17_] = init17_(i0_17_)
                    d_627_v7_ = nw114_
                    d_630_v8_: _dafny.MultiSet
                    d_630_v8_ = _dafny.MultiSet([d_627_v7_])
                    d_631_v9_: _dafny.Set
                    d_631_v9_ = _dafny.Set({not((self).f30), (self).f30})
                    d_632_v10_: _dafny.MultiSet
                    d_632_v10_ = _dafny.MultiSet([self.f25])
                    d_633_v11_: _dafny.Set
                    d_633_v11_ = _dafny.Set({489, len(d_621_v2_)})
                    d_634_v12_: _dafny.Map
                    d_634_v12_ = _dafny.Map({p1: p2})
                    d_635_v13_: _dafny.Seq
                    d_635_v13_ = _dafny.SeqWithoutIsStrInference([d_634_v12_])
                    d_636_v14_: _dafny.Map
                    d_636_v14_ = _dafny.Map({(d_635_v13_)[default__.safeIndex(self.f25, len(d_635_v13_))]: (self).f24})
                    d_637_v15_: _dafny.Seq
                    d_637_v15_ = _dafny.SeqWithoutIsStrInference([p2])
                    d_638_v16_: _dafny.Map
                    d_638_v16_ = _dafny.Map({self.f25: (self).f24})
                    d_639_v17_: _dafny.Array
                    nw115_ = _dafny.Array(None, 27)
                    nw115_[int(0)] = default__.safeDivisionInt(self.f25, self.f25)
                    nw115_[int(1)] = p2
                    nw115_[int(2)] = len(d_621_v2_)
                    nw115_[int(3)] = default__.safeModuloInt(len(d_626_v6_), self.f25)
                    nw115_[int(4)] = self.f25
                    nw115_[int(5)] = len(d_621_v2_)
                    nw115_[int(6)] = default__.safeModuloInt(((d_630_v8_)[d_627_v7_] if (d_627_v7_) in (d_630_v8_) else len(d_631_v9_)), 108)
                    nw115_[int(7)] = self.f25
                    nw115_[int(8)] = (p2) * (p2)
                    nw115_[int(9)] = ((d_632_v10_)[p2] if (p2) in (d_632_v10_) else p2)
                    nw115_[int(10)] = default__.safeModuloInt(len(d_633_v11_), p2)
                    nw115_[int(11)] = -441
                    nw115_[int(12)] = len(d_636_v14_)
                    nw115_[int(13)] = p2
                    nw115_[int(14)] = 141
                    nw115_[int(15)] = (0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([499 for d_640_i3_ in range(default__.abs(592))]))))
                    nw115_[int(16)] = self.f25
                    nw115_[int(17)] = p2
                    nw115_[int(18)] = self.f25
                    nw115_[int(19)] = (0) - (default__.safeModuloInt(len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_641_i4_ in range(default__.abs(707))])).set(default__.safeIndex(self.f25, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_642_i4_ in range(default__.abs(707))]))), (self).f24)), 257))
                    nw115_[int(20)] = p2
                    nw115_[int(21)] = (p2 if (self).f30 else default__.fm17((self).f30, True, p1, globalState))
                    nw115_[int(22)] = len(d_637_v15_)
                    nw115_[int(23)] = (0) - (p2)
                    nw115_[int(24)] = (len(d_638_v16_)) - (624)
                    nw115_[int(25)] = (286) + (default__.fm17(p1, (self).f30, p1, globalState))
                    nw115_[int(26)] = self.f25
                    d_639_v17_ = nw115_
                    index108_ = default__.safeIndex(280, (d_639_v17_).length(0))
                    (d_639_v17_)[index108_] = p2
                    index109_ = default__.safeIndex(280, (d_639_v17_).length(0))
                    (d_639_v17_)[index109_] = self.f25
                    pass
            pass
        d_643_v18_: _dafny.Map
        d_643_v18_ = _dafny.Map({(self).f30: p1})
        d_644_v19_: T1
        nw116_ = C3()
        nw116_.ctor__(p1, (self).f24, len(d_643_v18_))
        d_644_v19_ = nw116_
        d_645_v20_: _dafny.Seq
        d_645_v20_ = _dafny.SeqWithoutIsStrInference([d_644_v19_])
        (globalState).f11 = (0) - ((self.f25) * (len(_dafny.SeqWithoutIsStrInference([d_645_v20_, d_645_v20_, d_645_v20_, d_645_v20_]))))
        d_646_v21_: _dafny.Map
        d_646_v21_ = _dafny.Map({682: p1})
        d_647_v22_: _dafny.Array
        nw117_ = _dafny.Array(None, 3)
        nw117_[int(0)] = (((d_646_v21_)[self.f25] if (self.f25) in (d_646_v21_) else p1)) and (p1)
        nw117_[int(1)] = ((d_643_v18_)[(self).f30] if ((self).f30) in (d_643_v18_) else (self).f30)
        nw117_[int(2)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uxddo"))) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "teel")))
        d_647_v22_ = nw117_
        index110_ = default__.safeIndex(612, (d_647_v22_).length(0))
        (d_647_v22_)[index110_] = p1
        d_648_v23_: _dafny.Map
        d_648_v23_ = _dafny.Map({p2: p2})
        index111_ = default__.safeIndex(612, (d_647_v22_).length(0))
        (d_647_v22_)[index111_] = default__.fm0(len(d_648_v23_), p1, p1, d_644_v19_.f25, globalState)
        d_646_v21_ = (d_646_v21_).set(d_644_v19_.f25, (self.f25) >= ((0) - (d_644_v19_.f25)))
        d_649_v24_: C1
        nw118_ = C1()
        nw118_.ctor__((d_644_v19_).f24, default__.safeDivisionInt(d_644_v19_.f25, len(d_643_v18_)))
        d_649_v24_ = nw118_

    @property
    def f30(self):
        return self._f30

class C5(T1, T0):
    def  __init__(self):
        self._f25: int = int(0)
        self._f24: str = _dafny.CodePoint('D')
        self._f26: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f25(self):
        return self._f25
    @f25.setter
    def f25(self, value):
        self._f25 = value
    @property
    def f24(self):
        return self._f24
    def ctor__(self, f26, f24, f25):
        (self)._f26 = f26
        (self)._f24 = f24
        (self).f25 = f25

    def fm3(self, p0, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "enhmpxpn"))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bob")) if (self).f26 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ihqo"))))

    def m1(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: _dafny.Seq = _dafny.Seq({})
        r2: bool = False
        d_650_v0_: C2
        nw119_ = C2()
        nw119_.ctor__(self.f25, p1)
        d_650_v0_ = nw119_
        d_651_i0_: int
        d_651_i0_ = 0
        with _dafny.label("4"):
            pat_let_tv4_ = d_650_v0_
            pat_let_tv5_ = d_650_v0_
            pat_let_tv6_ = p1
            pat_let_tv7_ = d_650_v0_
            pat_let_tv8_ = d_650_v0_
            def lambda36_(source12_):
                if source12_.is_DC19:
                    d_680___mcc_h0_ = source12_.cf27
                    d_681_cf27_ = d_680___mcc_h0_
                    return not((_dafny.Set({(pat_let_tv4_).f28})) != (_dafny.Set({(pat_let_tv5_).f28, pat_let_tv6_})))
                elif source12_.is_DC20:
                    d_682___mcc_h1_ = source12_.cf28
                    d_683___mcc_h2_ = source12_.cf29
                    d_684_cf29_ = d_683___mcc_h2_
                    d_685_cf28_ = d_682___mcc_h1_
                    return not(True)
                elif source12_.is_DC21:
                    d_686___mcc_h3_ = source12_.cf30
                    d_687___mcc_h4_ = source12_.cf31
                    d_688___mcc_h5_ = source12_.cf32
                    d_689_cf32_ = d_688___mcc_h5_
                    d_690_cf31_ = d_687___mcc_h4_
                    d_691_cf30_ = d_686___mcc_h3_
                    return (pat_let_tv7_).f28
                elif True:
                    d_692___mcc_h6_ = source12_.cf26
                    d_693_cf26_ = d_692___mcc_h6_
                    return (pat_let_tv8_).f28

            while not(lambda36_(default__.fm14((self).f26, globalState))):
                with _dafny.c_label("4"):
                    if (d_651_i0_) >= (100):
                        raise _dafny.Break("4")
                    d_651_i0_ = (d_651_i0_) + (1)
                    d_652_v1_: _dafny.Array
                    def lambda33_(d_653_i1_):
                        return (self).f26

                    init18_ = lambda33_
                    nw120_ = _dafny.Array(None, 22)
                    for i0_18_ in range(nw120_.length(0)):
                        nw120_[i0_18_] = init18_(i0_18_)
                    d_652_v1_ = nw120_
                    d_654_v2_: _dafny.Seq
                    d_654_v2_ = _dafny.SeqWithoutIsStrInference([d_652_v1_])
                    d_655_v3_: T1
                    nw121_ = C3()
                    nw121_.ctor__((d_650_v0_).f28, (self).f24, ((d_650_v0_).f27) - ((0) - (self.f25)))
                    d_655_v3_ = nw121_
                    rhs129_ = d_655_v3_.f25
                    rhs130_ = not((d_650_v0_).f28)
                    rhs131_ = p1
                    rhs132_ = d_654_v2_
                    rhs133_ = d_655_v3_
                    lhs119_ = globalState
                    lhs120_ = globalState
                    lhs119_.f0 = rhs129_
                    lhs120_.f18 = rhs130_
                    r0 = rhs131_
                    d_654_v2_ = rhs132_
                    d_655_v3_ = rhs133_
                    r0 = False
                    d_656_v4_: _dafny.Seq
                    d_656_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vfafw"))
                    (globalState).f3 = ((d_656_v4_) + (_dafny.SeqWithoutIsStrInference([(d_655_v3_).f24 for d_657_i2_ in range(default__.abs(-183))]))) + ((d_656_v4_ if (self).f26 else d_656_v4_))
                    if (d_650_v0_).f28:
                        r2 = (d_650_v0_).f28
                        (globalState).f18 = not((((d_650_v0_).f27) - ((0) - (d_655_v3_.f25))) > ((d_650_v0_).f27))
                        (globalState).f0 = default__.safeModuloInt((d_650_v0_).f27, (d_650_v0_).f27)
                        (globalState).f22 = (d_655_v3_.f25) + ((0) - (default__.safeDivisionInt(783, (d_650_v0_).f27)))
                        d_658_v5_: _dafny.Map
                        d_658_v5_ = _dafny.Map({(d_650_v0_).f28: p1})
                        d_659_v6_: _dafny.Array
                        nw122_ = _dafny.Array(None, 4)
                        d_659_v6_ = nw122_
                        d_660_v7_: _dafny.Map
                        d_660_v7_ = _dafny.Map({d_658_v5_: d_659_v6_})
                        d_661_v8_: _dafny.Seq
                        d_661_v8_ = _dafny.SeqWithoutIsStrInference([d_659_v6_])
                        d_662_v9_: C0
                        nw123_ = C0()
                        nw123_.ctor__()
                        d_662_v9_ = nw123_
                        d_663_v10_: _dafny.Seq
                        d_663_v10_ = _dafny.SeqWithoutIsStrInference([d_662_v9_])
                        d_664_v11_: _dafny.Map
                        d_664_v11_ = _dafny.Map({(d_650_v0_).f27: len(d_663_v10_)})
                        d_665_v12_: D11
                        d_665_v12_ = D11_DC22(d_659_v6_)
                        d_666_v13_: _dafny.Array
                        nw124_ = _dafny.Array(None, 25)
                        nw124_[int(0)] = (d_660_v7_) | (d_660_v7_)
                        nw124_[int(1)] = (d_660_v7_) | (d_660_v7_)
                        nw124_[int(2)] = (_dafny.Map({d_658_v5_: d_659_v6_})).set(d_658_v5_, d_659_v6_)
                        nw124_[int(3)] = (d_660_v7_ if False else d_660_v7_)
                        nw124_[int(4)] = _dafny.Map({d_658_v5_: (d_661_v8_)[default__.safeIndex(len(d_664_v11_), len(d_661_v8_))]})
                        nw124_[int(5)] = d_660_v7_
                        nw124_[int(6)] = d_660_v7_
                        nw124_[int(7)] = (d_660_v7_) | (d_660_v7_)
                        nw124_[int(8)] = (d_660_v7_) | (_dafny.Map({d_658_v5_: d_659_v6_}))
                        nw124_[int(9)] = d_660_v7_
                        nw124_[int(10)] = d_660_v7_
                        nw124_[int(11)] = d_660_v7_
                        nw124_[int(12)] = d_660_v7_
                        nw124_[int(13)] = d_660_v7_
                        nw124_[int(14)] = d_660_v7_
                        nw124_[int(15)] = d_660_v7_
                        nw124_[int(16)] = (d_660_v7_) | (d_660_v7_)
                        nw124_[int(17)] = d_660_v7_
                        nw124_[int(18)] = d_660_v7_
                        nw124_[int(19)] = (d_660_v7_).set(d_658_v5_, (d_665_v12_).cf33)
                        nw124_[int(20)] = (d_660_v7_) | (d_660_v7_)
                        nw124_[int(21)] = (d_660_v7_) | ((d_660_v7_).set(d_658_v5_, d_659_v6_))
                        nw124_[int(22)] = d_660_v7_
                        nw124_[int(23)] = d_660_v7_
                        nw124_[int(24)] = d_660_v7_
                        d_666_v13_ = nw124_
                        index112_ = default__.safeIndex(347, (d_666_v13_).length(0))
                        (d_666_v13_)[index112_] = (_dafny.Map({default__.fm19((d_650_v0_).f28, globalState): d_659_v6_})) | (d_660_v7_)
                        index113_ = default__.safeIndex(347, (d_666_v13_).length(0))
                        (d_666_v13_)[index113_] = (d_660_v7_) | (d_660_v7_)
                    elif True:
                        d_667_v14_: D4
                        d_667_v14_ = D4_DC8()
                        d_668_v15_: _dafny.MultiSet
                        d_668_v15_ = _dafny.MultiSet([D4_DC8(), d_667_v14_])
                        d_669_v16_: _dafny.Seq
                        d_669_v16_ = _dafny.SeqWithoutIsStrInference([d_667_v14_, d_667_v14_])
                        rhs134_ = p1
                        rhs135_ = (_dafny.MultiSet(d_669_v16_)) | (_dafny.MultiSet([d_667_v14_, d_667_v14_, d_667_v14_, d_667_v14_, d_667_v14_]))
                        rhs136_ = (self).f26
                        rhs137_ = p1
                        lhs121_ = globalState
                        lhs122_ = globalState
                        lhs121_.f18 = rhs134_
                        d_668_v15_ = rhs135_
                        r2 = rhs136_
                        lhs122_.f18 = rhs137_
                        d_670_v17_: _dafny.Array
                        def lambda34_(d_671_v0_):
                            def lambda35_(d_672_i3_):
                                return default__.safeDivisionInt(d_672_i3_, (0) - ((d_671_v0_).f27))

                            return lambda35_

                        init19_ = lambda34_(d_650_v0_)
                        nw125_ = _dafny.Array(None, 27)
                        for i0_19_ in range(nw125_.length(0)):
                            nw125_[i0_19_] = init19_(i0_19_)
                        d_670_v17_ = nw125_
                        nw126_ = _dafny.Array(int(0), 12)
                        d_670_v17_ = nw126_
                        d_673_v18_: D10
                        d_673_v18_ = D10_DC19((d_655_v3_).f24)
                        d_674_v19_: D10
                        d_674_v19_ = D10_DC21(d_655_v3_.f25, (self).f26, (d_650_v0_).f28)
                        def iife54_(_pat_let10_0):
                            def iife55_(d_675_dt__update__tmp_h0_):
                                def iife56_(_pat_let11_0):
                                    def iife57_(d_676_dt__update_hcf27_h0_):
                                        return D10_DC19(d_676_dt__update_hcf27_h0_)
                                    return iife57_(_pat_let11_0)
                                return iife56_(_dafny.CodePoint('g'))
                            return iife55_(_pat_let10_0)
                        d_673_v18_ = (d_673_v18_ if not((d_650_v0_).f28) else (iife54_(d_673_v18_) if (d_674_v19_).cf32 else default__.fm20(d_655_v3_.f25, True, globalState)))
                        d_677_v20_: int
                        d_678_v21_: int
                        out8_: int
                        out9_: int
                        out8_, out9_ = (d_650_v0_).m3((self).f26, (self).f26, globalState)
                        d_677_v20_ = out8_
                        d_678_v21_ = out9_
                        d_679_v22_: C0
                        nw127_ = C0()
                        nw127_.ctor__()
                        d_679_v22_ = nw127_
                        d_679_v22_ = d_679_v22_
                    pass
            pass
        d_694_v23_: _dafny.Seq
        d_694_v23_ = _dafny.SeqWithoutIsStrInference([p1, (self).f26, p1, p1, p1])
        d_695_v24_: _dafny.Map
        d_695_v24_ = _dafny.Map({self.f25: p1})
        d_696_v25_: _dafny.Seq
        d_696_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pdxyd"))
        d_697_v26_: _dafny.Array
        nw128_ = _dafny.Array(None, 20)
        nw128_[int(0)] = not(p1)
        nw128_[int(1)] = not(((d_650_v0_).f28 if p1 else p1))
        nw128_[int(2)] = (self).f26
        nw128_[int(3)] = ((0) - (self.f25)) >= (self.f25)
        nw128_[int(4)] = not(not((self.f25) < (self.f25)))
        nw128_[int(5)] = (self).f26
        nw128_[int(6)] = (self).f26
        nw128_[int(7)] = (self).f26
        nw128_[int(8)] = (self).f26
        nw128_[int(9)] = p1
        nw128_[int(10)] = p1
        nw128_[int(11)] = ((d_650_v0_).f28) == (default__.fm0((d_650_v0_).f27, p1, (d_650_v0_).f28, self.f25, globalState))
        nw128_[int(12)] = ((self).f26) or (p1)
        nw128_[int(13)] = (d_694_v23_)[default__.safeIndex((d_650_v0_).f27, len(d_694_v23_))]
        nw128_[int(14)] = not((self).f26)
        nw128_[int(15)] = (d_650_v0_).f28
        nw128_[int(16)] = ((d_695_v24_)[self.f25] if (self.f25) in (d_695_v24_) else (self).f26)
        nw128_[int(17)] = p1
        nw128_[int(18)] = (self.f25) >= (len(d_696_v25_))
        nw128_[int(19)] = (d_650_v0_).f28
        d_697_v26_ = nw128_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_697_v26_).length(0)):
            d_698_i4_: int = guard_loop_2_
            if (True) and (((0) <= (d_698_i4_)) and ((d_698_i4_) < ((d_697_v26_).length(0)))):
                (d_697_v26_)[(d_698_i4_)] = ((_dafny.MultiSet([(d_650_v0_).f28]) if (d_650_v0_).f28 else _dafny.MultiSet([(d_650_v0_).f28, not(p1), (d_650_v0_).f28]))).ispropersubset((_dafny.MultiSet([(d_650_v0_).f28])).intersection(_dafny.MultiSet(d_694_v23_)))
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_697_v26_).length(0)):
            d_699_i5_: int = guard_loop_3_
            if (True) and (((0) <= (d_699_i5_)) and ((d_699_i5_) < ((d_697_v26_).length(0)))):
                (d_697_v26_)[(d_699_i5_)] = (self).f26
        d_700_v27_: _dafny.Map
        d_700_v27_ = _dafny.Map({(self.f25) * ((d_650_v0_).f27): d_650_v0_})
        d_650_v0_ = ((d_700_v27_)[(d_650_v0_).f27] if ((d_650_v0_).f27) in (d_700_v27_) else d_650_v0_)
        if p1:
            d_701_v28_: str
            d_701_v28_ = _dafny.CodePoint('d')
            d_701_v28_ = (self).f24
            d_702_v29_: D7
            d_702_v29_ = D7_DC14((self).f26, (0) - (self.f25), (d_650_v0_).f27)
            d_703_v30_: _dafny.Map
            d_703_v30_ = _dafny.Map({_dafny.MultiSet(d_694_v23_): (d_702_v29_).cf22})
            d_704_v31_: _dafny.Map
            d_704_v31_ = _dafny.Map({((d_695_v24_)[(d_650_v0_).f27] if ((d_650_v0_).f27) in (d_695_v24_) else True): not((d_650_v0_).f28)})
            d_705_v32_: C4
            nw129_ = C4()
            nw129_.ctor__(d_703_v30_, ((d_704_v31_)[(self).f26] if ((self).f26) in (d_704_v31_) else p1), _dafny.CodePoint('f'), default__.safeDivisionInt(656, self.f25))
            d_705_v32_ = nw129_
            d_705_v32_ = d_705_v32_
            if not(False):
                (globalState).f11 = self.f25
                d_706_v33_: D7
                d_706_v33_ = D7_DC13(d_695_v24_)
                d_707_v34_: _dafny.MultiSet
                d_707_v34_ = _dafny.MultiSet([(d_706_v33_ if (self).f26 else d_706_v33_)])
                d_708_v35_: _dafny.Array
                def lambda37_(d_709_i6_):
                    return (d_709_i6_) - (-813)

                init20_ = lambda37_
                nw130_ = _dafny.Array(None, 10)
                for i0_20_ in range(nw130_.length(0)):
                    nw130_[i0_20_] = init20_(i0_20_)
                d_708_v35_ = nw130_
                d_710_v36_: D6
                d_710_v36_ = D6_DC12(p1, (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pshrqd"))).set(default__.safeIndex((d_650_v0_).f27, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pshrqd")))), d_701_v28_))) + ((d_650_v0_).f27))
                rhs138_ = _dafny.MultiSet(default__.fm21(default__.safeModuloInt((d_650_v0_).f27, (d_650_v0_).f27), default__.safeModuloInt((d_650_v0_).f27, (d_650_v0_).f27), globalState))
                rhs139_ = d_708_v35_
                rhs140_ = d_710_v36_
                d_707_v34_ = rhs138_
                d_708_v35_ = rhs139_
                d_710_v36_ = rhs140_
                (globalState).f3 = ((d_696_v25_) + (d_696_v25_)) + (d_696_v25_)
                d_711_v37_: D10
                d_711_v37_ = D10_DC21(len((d_704_v31_).set(p1, (d_650_v0_).f28)), (d_650_v0_).f28, (d_650_v0_).f28)
                d_712_v38_: _dafny.Set
                d_712_v38_ = _dafny.Set({d_711_v37_})
                d_712_v38_ = d_712_v38_
                (globalState).f18 = (d_694_v23_) <= (default__.fm7(globalState))
            elif True:
                d_713_v39_: _dafny.Array
                nw131_ = _dafny.Array(int(0), 10)
                d_713_v39_ = nw131_
                d_714_v40_: _dafny.MultiSet
                d_714_v40_ = _dafny.MultiSet([(d_650_v0_).f28])
                index114_ = default__.safeIndex(824, (d_713_v39_).length(0))
                (d_713_v39_)[index114_] = default__.safeModuloInt((d_714_v40_).cardinality, 515)
                index115_ = default__.safeIndex(824, (d_713_v39_).length(0))
                (d_713_v39_)[index115_] = 460
                d_715_v41_: T0
                nw132_ = C1()
                nw132_.ctor__((self).f24, (len(d_694_v23_)) * ((d_650_v0_).f27))
                d_715_v41_ = nw132_
                d_716_v42_: _dafny.Map
                d_716_v42_ = _dafny.Map({(d_650_v0_).f27: (self).f24})
                d_717_v43_: _dafny.Set
                d_717_v43_ = _dafny.Set({len(d_696_v25_), (d_650_v0_).f27})
                nw133_ = C4()
                nw133_.ctor__(_dafny.Map({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(d_650_v0_).f28, (d_650_v0_).f28])): d_715_v41_.f25}), (d_650_v0_).f28, ((d_716_v42_)[len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kanvy")))] if (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kanvy")))) in (d_716_v42_) else d_701_v28_), (len(d_717_v43_)) + (len(_dafny.SeqWithoutIsStrInference([d_701_v28_ for d_718_i7_ in range(default__.abs(408))]))))
                d_715_v41_ = nw133_
                index116_ = default__.safeIndex(824, (d_713_v39_).length(0))
                (d_713_v39_)[index116_] = (d_650_v0_).f27
                d_719_v44_: _dafny.Array
                def lambda38_(d_720_v41_, d_721_v39_, d_722_v0_):
                    def lambda39_(d_723_i8_):
                        return D3_DC6(d_720_v41_.f25, (d_721_v39_)[default__.safeIndex(824, (d_721_v39_).length(0))], (d_722_v0_).f28)

                    return lambda39_

                init21_ = lambda38_(d_715_v41_, d_713_v39_, d_650_v0_)
                nw134_ = _dafny.Array(None, 1)
                for i0_21_ in range(nw134_.length(0)):
                    nw134_[i0_21_] = init21_(i0_21_)
                d_719_v44_ = nw134_
                index117_ = default__.safeIndex(811, (d_719_v44_).length(0))
                (d_719_v44_)[index117_] = D3_DC6(default__.fm17((d_650_v0_).f28, (d_705_v32_).f30, p1, globalState), self.f25, default__.fm0(self.f25, not((d_705_v32_).f30), False, self.f25, globalState))
                d_724_v45_: _dafny.Map
                d_724_v45_ = _dafny.Map({(d_650_v0_).f28: (d_713_v39_)[default__.safeIndex(824, (d_713_v39_).length(0))]})
                d_725_v46_: D3
                d_725_v46_ = D3_DC6((d_650_v0_).f27, d_715_v41_.f25, p1)
                index118_ = default__.safeIndex(811, (d_719_v44_).length(0))
                rhs141_ = d_724_v45_
                rhs142_ = d_725_v46_
                lhs123_ = globalState
                lhs124_ = d_719_v44_
                lhs125_ = default__.safeIndex(811, (d_719_v44_).length(0))
                lhs123_.f13 = rhs141_
                lhs124_[lhs125_] = rhs142_
                d_726_v47_: _dafny.Array
                nw135_ = _dafny.Array(None, 24)
                d_726_v47_ = nw135_
                d_727_v48_: D11
                d_727_v48_ = D11_DC22(d_726_v47_)
                d_728_v49_: _dafny.Seq
                d_728_v49_ = _dafny.SeqWithoutIsStrInference([d_727_v48_, d_727_v48_])
                d_729_v50_: _dafny.Seq
                d_729_v50_ = _dafny.SeqWithoutIsStrInference([(0) - ((d_713_v39_)[default__.safeIndex(824, (d_713_v39_).length(0))])])
                index119_ = default__.safeIndex(824, (d_713_v39_).length(0))
                rhs143_ = (d_728_v49_) + ((d_728_v49_).set(default__.safeIndex(len(d_729_v50_), len(d_728_v49_)), d_727_v48_))
                rhs144_ = ((d_650_v0_).f27) * (d_715_v41_.f25)
                rhs145_ = d_697_v26_
                lhs126_ = d_713_v39_
                lhs127_ = default__.safeIndex(824, (d_713_v39_).length(0))
                d_728_v49_ = rhs143_
                lhs126_[lhs127_] = rhs144_
                d_697_v26_ = rhs145_
            d_730_v51_: C1
            nw136_ = C1()
            nw136_.ctor__(_dafny.CodePoint('v'), len(d_696_v25_))
            d_730_v51_ = nw136_
            d_731_v53_: _dafny.MultiSet
            d_731_v53_ = _dafny.MultiSet([(d_650_v0_).f28, (d_650_v0_).f28, not((d_650_v0_).f28), not(True), ((d_704_v31_)[(d_650_v0_).f28] if ((d_650_v0_).f28) in (d_704_v31_) else not((self).f26))])
            def iife58_():
                coll34_ = _dafny.Set()
                compr_34_: int
                for compr_34_ in _dafny.IntegerRange(93, 576):
                    d_732_v52_: int = compr_34_
                    if ((93) <= (d_732_v52_)) and ((d_732_v52_) < (576)):
                        coll34_ = coll34_.union(_dafny.Set([default__.safeModuloInt(d_732_v52_, (d_650_v0_).f27)]))
                return _dafny.Set(coll34_)
            rhs146_ = (0) - ((d_650_v0_).f27)
            rhs147_ = not(p1)
            rhs148_ = d_697_v26_
            rhs149_ = ((d_650_v0_).fm4(iife58_()
            , len(d_696_v25_), (d_650_v0_).f28, globalState)).isdisjoint(d_731_v53_)
            lhs128_ = globalState
            lhs129_ = globalState
            lhs128_.f22 = rhs146_
            r0 = rhs147_
            d_697_v26_ = rhs148_
            lhs129_.f18 = rhs149_
        elif True:
            r1 = d_694_v23_
            (globalState).f2 = (((d_650_v0_).f27) - (default__.fm17(not(True), (self).f26, p1, globalState))) + ((d_650_v0_).f27)
            d_733_v54_: _dafny.Map
            d_733_v54_ = _dafny.Map({(d_650_v0_).f28: len(d_696_v25_)})
            d_734_v55_: _dafny.Map
            d_734_v55_ = _dafny.Map({not(default__.fm0(-691, (d_694_v23_)[default__.safeIndex(-920, len(d_694_v23_))], (self).f26, ((d_733_v54_)[True] if (True) in (d_733_v54_) else (0) - ((d_650_v0_).f27)), globalState)): (d_650_v0_).f27})
            (globalState).f13 = d_733_v54_
            if p1:
                d_735_v56_: _dafny.Map
                d_735_v56_ = _dafny.Map({(self).f24: (d_650_v0_).f27})
                d_736_v57_: _dafny.MultiSet
                d_736_v57_ = _dafny.MultiSet([d_694_v23_, d_694_v23_])
                d_737_v58_: _dafny.MultiSet
                d_737_v58_ = _dafny.MultiSet([(d_650_v0_).f28, (d_650_v0_).f28, p1, default__.fm0((d_650_v0_).f27, (d_650_v0_).f28, (d_650_v0_).f28, (d_650_v0_).f27, globalState)])
                d_738_v59_: _dafny.Map
                d_738_v59_ = _dafny.Map({(d_650_v0_).f27: (d_650_v0_).f27})
                d_739_v60_: _dafny.Set
                d_739_v60_ = _dafny.Set({(d_650_v0_).f27})
                d_740_v61_: _dafny.Seq
                d_740_v61_ = _dafny.SeqWithoutIsStrInference([((d_738_v59_)[len(d_739_v60_)] if (len(d_739_v60_)) in (d_738_v59_) else 125), len(_dafny.SeqWithoutIsStrInference([True, (self).f26, (d_650_v0_).f28, p1, (d_650_v0_).f28])), (d_650_v0_).f27, (d_650_v0_).f27, len(d_738_v59_)])
                d_741_v62_: _dafny.Array
                nw137_ = _dafny.Array(None, 26)
                nw137_[int(0)] = default__.fm5((d_650_v0_).f27, (d_650_v0_).f27, globalState)
                nw137_[int(1)] = (d_650_v0_).f27
                nw137_[int(2)] = self.f25
                nw137_[int(3)] = ((d_736_v57_)[_dafny.SeqWithoutIsStrInference([(self).f26])] if (_dafny.SeqWithoutIsStrInference([(self).f26])) in (d_736_v57_) else (d_650_v0_).f27)
                nw137_[int(4)] = (d_650_v0_).f27
                nw137_[int(5)] = (d_650_v0_).f27
                nw137_[int(6)] = (d_650_v0_).f27
                nw137_[int(7)] = (d_650_v0_).f27
                nw137_[int(8)] = self.f25
                nw137_[int(9)] = (d_650_v0_).f27
                nw137_[int(10)] = len(p2)
                nw137_[int(11)] = (d_650_v0_).f27
                nw137_[int(12)] = (d_650_v0_).f27
                nw137_[int(13)] = (d_650_v0_).f27
                nw137_[int(14)] = (d_650_v0_).f27
                nw137_[int(15)] = (d_737_v58_).cardinality
                nw137_[int(16)] = (d_650_v0_).f27
                nw137_[int(17)] = (d_650_v0_).f27
                nw137_[int(18)] = (d_650_v0_).f27
                nw137_[int(19)] = len(d_694_v23_)
                nw137_[int(20)] = self.f25
                nw137_[int(21)] = self.f25
                nw137_[int(22)] = (0) - (len(d_740_v61_))
                nw137_[int(23)] = self.f25
                nw137_[int(24)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uxb")))
                nw137_[int(25)] = (d_650_v0_).f27
                d_741_v62_ = nw137_
                d_742_v63_: _dafny.Array
                nw138_ = _dafny.Array(None, 13)
                nw138_[int(0)] = d_741_v62_
                nw138_[int(1)] = d_741_v62_
                nw138_[int(2)] = d_741_v62_
                nw138_[int(3)] = d_741_v62_
                nw138_[int(4)] = d_741_v62_
                nw138_[int(5)] = d_741_v62_
                nw138_[int(6)] = d_741_v62_
                nw138_[int(7)] = d_741_v62_
                nw138_[int(8)] = d_741_v62_
                nw138_[int(9)] = d_741_v62_
                nw138_[int(10)] = d_741_v62_
                nw138_[int(11)] = d_741_v62_
                nw138_[int(12)] = d_741_v62_
                d_742_v63_ = nw138_
                d_743_v64_: _dafny.Seq
                d_743_v64_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bedhejjw"))
                d_744_v65_: _dafny.Map
                d_744_v65_ = _dafny.Map({(d_743_v64_): (_dafny.Map({(self).f24: (d_650_v0_).f27})) | (d_735_v56_)})
                rhs150_ = d_650_v0_
                rhs151_ = ((d_744_v65_)[(d_696_v25_) + (d_696_v25_)] if ((d_696_v25_) + (d_696_v25_)) in (d_744_v65_) else d_735_v56_)
                rhs152_ = (d_650_v0_).f27
                rhs153_ = d_742_v63_
                lhs130_ = globalState
                d_650_v0_ = rhs150_
                d_735_v56_ = rhs151_
                lhs130_.f10 = rhs152_
                d_742_v63_ = rhs153_
                index120_ = default__.safeIndex(340, (d_741_v62_).length(0))
                (d_741_v62_)[index120_] = (d_650_v0_).f27
                index121_ = default__.safeIndex(340, (d_741_v62_).length(0))
                (d_741_v62_)[index121_] = (self.f25) * (self.f25)
                d_745_v66_: _dafny.Map
                d_746_v67_: bool
                out10_: _dafny.Map
                out11_: bool
                out10_, out11_ = (self).m2(globalState)
                d_745_v66_ = out10_
                d_746_v67_ = out11_
                d_747_v68_: C3
                nw139_ = C3()
                nw139_.ctor__(d_746_v67_, (self).f24, self.f25)
                d_747_v68_ = nw139_
                d_650_v0_ = d_650_v0_
            elif True:
                d_748_v69_: _dafny.Map
                d_748_v69_ = _dafny.Map({(self).f26: default__.fm10((self).f26, (d_650_v0_).f28, False, (d_650_v0_).f28, globalState)})
                d_749_v70_: _dafny.Seq
                d_749_v70_ = _dafny.SeqWithoutIsStrInference([d_748_v69_, d_748_v69_, d_748_v69_])
                d_749_v70_ = d_749_v70_
                d_750_v71_: C1
                nw140_ = C1()
                nw140_.ctor__((self).f24, (d_650_v0_).f27)
                d_750_v71_ = nw140_
                d_751_v72_: _dafny.Map
                d_751_v72_ = _dafny.Map({p1: d_750_v71_})
                d_750_v71_ = ((d_751_v72_)[(self).f26] if ((self).f26) in (d_751_v72_) else d_750_v71_)
                d_752_v73_: _dafny.Map
                d_752_v73_ = _dafny.Map({(self).f26: d_694_v23_})
                d_752_v73_ = (d_752_v73_).set(False, d_694_v23_)
                d_753_v74_: _dafny.Set
                d_753_v74_ = _dafny.Set({(self).f26})
                r2 = (d_753_v74_).issubset(d_753_v74_)
                index122_ = default__.safeIndex(163, (d_697_v26_).length(0))
                (d_697_v26_)[index122_] = (self).f26
                index123_ = default__.safeIndex(163, (d_697_v26_).length(0))
                (d_697_v26_)[index123_] = (d_650_v0_).f28
            d_754_v75_: _dafny.Array
            nw141_ = _dafny.Array(_dafny.Array(None, 0), 17)
            d_754_v75_ = nw141_
            index124_ = default__.safeIndex(244, (d_754_v75_).length(0))
            (d_754_v75_)[index124_] = d_697_v26_
            d_755_v76_: _dafny.Array
            nw142_ = _dafny.Array(int(0), 9)
            d_755_v76_ = nw142_
            index125_ = default__.safeIndex(26, (d_755_v76_).length(0))
            (d_755_v76_)[index125_] = -380
            d_756_v77_: _dafny.Seq
            d_756_v77_ = _dafny.SeqWithoutIsStrInference([d_697_v26_, d_697_v26_])
            d_757_v78_: D3
            d_757_v78_ = D3_DC5(d_694_v23_, self.f25, (d_650_v0_).f27)
            index126_ = default__.safeIndex(244, (d_754_v75_).length(0))
            index127_ = default__.safeIndex(26, (d_755_v76_).length(0))
            rhs154_ = (d_756_v77_)[default__.safeIndex((d_650_v0_).f27, len(d_756_v77_))]
            rhs155_ = self.f25
            rhs156_ = default__.safeDivisionInt((self.f25) * ((d_650_v0_).f27), (d_650_v0_).f27)
            rhs157_ = (d_694_v23_)[default__.safeIndex((d_757_v78_).cf8, len(d_694_v23_))]
            lhs131_ = d_754_v75_
            lhs132_ = default__.safeIndex(244, (d_754_v75_).length(0))
            lhs133_ = d_755_v76_
            lhs134_ = default__.safeIndex(26, (d_755_v76_).length(0))
            lhs135_ = globalState
            lhs131_[lhs132_] = rhs154_
            lhs133_[lhs134_] = rhs155_
            lhs135_.f15 = rhs156_
            r2 = rhs157_
        r0 = (self).f26
        r1 = (d_694_v23_).set(default__.safeIndex((self.f25) - (758), len(d_694_v23_)), (d_650_v0_).f28)
        r2 = (d_650_v0_).f28
        return r0, r1, r2

    def m2(self, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: bool = False
        d_758_v0_: D7
        d_758_v0_ = D7_DC14((self).f26, len((self).fm3((self).f24, globalState)), self.f25)
        if ((d_758_v0_).cf20 if (self).f26 else not(((0) - (self.f25)) > (961))):
            d_759_v2_: _dafny.Map
            d_759_v2_ = _dafny.Map({_dafny.CodePoint('n'): (self).f26})
            d_760_v3_: _dafny.Set
            def iife59_():
                coll35_ = _dafny.Map()
                compr_35_: str
                for compr_35_ in (d_759_v2_).keys.Elements:
                    d_761_v1_: str = compr_35_
                    if (d_761_v1_) in (d_759_v2_):
                        coll35_[d_761_v1_] = (0) - (self.f25)
                return _dafny.Map(coll35_)
            d_760_v3_ = _dafny.Set({self.f25, len(iife59_()
            )})
            d_762_v4_: D11
            d_762_v4_ = D11_DC23(len(d_760_v3_), (self).f26)
            d_763_v5_: C3
            nw143_ = C3()
            nw143_.ctor__(not((d_762_v4_).cf35), (self).f24, (self.f25 if (self).f26 else self.f25))
            d_763_v5_ = nw143_
            d_764_v6_: str
            d_764_v6_ = _dafny.CodePoint('f')
            rhs158_ = (default__.fm22(949, not(not((self).f26)), d_763_v5_.f31, globalState)).cf34
            rhs159_ = d_764_v6_
            rhs160_ = d_763_v5_.f31
            rhs161_ = (self).f26
            lhs136_ = globalState
            lhs137_ = globalState
            lhs138_ = globalState
            lhs136_.f0 = rhs158_
            d_764_v6_ = rhs159_
            lhs137_.f18 = rhs160_
            lhs138_.f18 = rhs161_
            d_765_v7_: _dafny.MultiSet
            d_765_v7_ = _dafny.MultiSet([self.f25, self.f25])
            d_766_v8_: _dafny.Seq
            d_766_v8_ = _dafny.SeqWithoutIsStrInference([self.f25])
            (self).f25 = default__.fm17((self).f26, (self).f26, (_dafny.MultiSet(d_766_v8_)).ispropersubset(d_765_v7_), globalState)
            d_767_v9_: _dafny.Map
            d_767_v9_ = _dafny.Map({d_766_v8_: d_763_v5_.f31})
            if (d_767_v9_) == (d_767_v9_):
                d_768_v10_: _dafny.Seq
                d_768_v10_ = _dafny.SeqWithoutIsStrInference([d_763_v5_.f31])
                d_768_v10_ = ((d_768_v10_) + (d_768_v10_)) + (d_768_v10_)
                (globalState).f18 = d_763_v5_.f31
                d_769_v11_: _dafny.Map
                d_769_v11_ = _dafny.Map({(self).f24: ((d_768_v10_).set(default__.safeIndex(self.f25, len(d_768_v10_)), d_763_v5_.f31)) + (d_768_v10_)})
                d_769_v11_ = (d_769_v11_).set((self).f24, d_768_v10_)
                d_770_v12_: _dafny.Array
                nw144_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 15)
                d_770_v12_ = nw144_
                d_770_v12_ = d_770_v12_
                (d_763_v5_).f31 = d_763_v5_.f31
            elif True:
                d_771_v13_: _dafny.Map
                d_771_v13_ = _dafny.Map({default__.fm17(d_763_v5_.f31, d_763_v5_.f31, d_763_v5_.f31, globalState): False})
                d_771_v13_ = d_771_v13_
                d_764_v6_ = (self).f24
                (globalState).f10 = (self.f25) - (self.f25)
                d_772_v14_: _dafny.Array
                def lambda40_(d_773_v5_):
                    def lambda41_(d_774_i0_):
                        return d_773_v5_.f31

                    return lambda41_

                init22_ = lambda40_(d_763_v5_)
                nw145_ = _dafny.Array(None, 20)
                for i0_22_ in range(nw145_.length(0)):
                    nw145_[i0_22_] = init22_(i0_22_)
                d_772_v14_ = nw145_
                d_775_v15_: _dafny.Seq
                d_775_v15_ = _dafny.SeqWithoutIsStrInference([d_765_v7_])
                d_776_v16_: _dafny.Map
                d_776_v16_ = _dafny.Map({(d_775_v15_) == (d_775_v15_): d_772_v14_})
                d_772_v14_ = ((d_776_v16_)[(self).f26] if ((self).f26) in (d_776_v16_) else d_772_v14_)
                d_777_v17_: _dafny.Map
                d_777_v17_ = _dafny.Map({_dafny.MultiSet([d_763_v5_.f31]): len(d_760_v3_)})
                d_778_v18_: _dafny.Map
                d_778_v18_ = _dafny.Map({default__.fm10((self).f26, d_763_v5_.f31, True, (self).f26, globalState): (self).f24})
                d_779_v19_: T0
                nw146_ = C4()
                nw146_.ctor__(d_777_v17_, d_763_v5_.f31, ((d_778_v18_)[(self).f24] if ((self).f24) in (d_778_v18_) else d_764_v6_), self.f25)
                d_779_v19_ = nw146_
                d_780_v20_: _dafny.Map
                d_780_v20_ = _dafny.Map({d_779_v19_: (self).f24})
                d_780_v20_ = (d_780_v20_).set(d_779_v19_, (d_779_v19_).f24)
            d_781_v21_: _dafny.Array
            nw147_ = _dafny.Array(None, 1)
            nw147_[int(0)] = d_763_v5_.f31
            d_781_v21_ = nw147_
            d_782_v22_: _dafny.Array
            nw148_ = _dafny.Array(D11.default()(), 21)
            d_782_v22_ = nw148_
            d_783_v23_: _dafny.Array
            nw149_ = _dafny.Array(None, 27)
            d_783_v23_ = nw149_
            d_784_v24_: D11
            d_784_v24_ = D11_DC22(d_783_v23_)
            index128_ = default__.safeIndex(483, (d_782_v22_).length(0))
            (d_782_v22_)[index128_] = d_784_v24_
            d_785_v25_: _dafny.Array
            nw150_ = _dafny.Array(_dafny.Set({}), 28)
            d_785_v25_ = nw150_
            d_786_v26_: _dafny.Array
            nw151_ = _dafny.Array(None, 4)
            d_786_v26_ = nw151_
            d_787_v27_: _dafny.Map
            d_787_v27_ = _dafny.Map({(self).f26: d_786_v26_})
            d_788_v28_: _dafny.Seq
            d_788_v28_ = _dafny.SeqWithoutIsStrInference([d_781_v21_])
            d_789_v29_: _dafny.Map
            d_789_v29_ = _dafny.Map({(d_787_v27_) | (d_787_v27_): (d_788_v28_)[default__.safeIndex(616, len(d_788_v28_))]})
            index129_ = default__.safeIndex(483, (d_782_v22_).length(0))
            rhs162_ = ((d_789_v29_)[_dafny.Map({False: d_786_v26_})] if (_dafny.Map({False: d_786_v26_})) in (d_789_v29_) else d_781_v21_)
            rhs163_ = d_784_v24_
            rhs164_ = default__.safeModuloInt((0) - (self.f25), self.f25)
            rhs165_ = d_785_v25_
            lhs139_ = d_782_v22_
            lhs140_ = default__.safeIndex(483, (d_782_v22_).length(0))
            lhs141_ = globalState
            d_781_v21_ = rhs162_
            lhs139_[lhs140_] = rhs163_
            lhs141_.f16 = rhs164_
            d_785_v25_ = rhs165_
        elif True:
            d_790_v30_: _dafny.Array
            nw152_ = _dafny.Array(_dafny.Seq({}), 26)
            d_790_v30_ = nw152_
            d_791_v31_: _dafny.Seq
            d_791_v31_ = _dafny.SeqWithoutIsStrInference([D10_DC21(self.f25, (self).f26, (self).f26)])
            index130_ = default__.safeIndex(189, (d_790_v30_).length(0))
            (d_790_v30_)[index130_] = (d_791_v31_) + (d_791_v31_)
            d_792_v32_: D10
            d_792_v32_ = D10_DC21(745, (self).f26, (self).f26)
            index131_ = default__.safeIndex(189, (d_790_v30_).length(0))
            (d_790_v30_)[index131_] = _dafny.SeqWithoutIsStrInference([d_792_v32_, d_792_v32_, d_792_v32_, d_792_v32_, d_792_v32_])
            d_793_v33_: _dafny.Seq
            d_793_v33_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qxbymslm"))
            d_794_v34_: D11
            d_794_v34_ = D11_DC24((self).f26, 802)
            d_795_v35_: D4
            d_795_v35_ = D4_DC8()
            d_796_v36_: _dafny.Map
            d_796_v36_ = _dafny.Map({(d_794_v34_).cf36: d_795_v35_})
            d_797_v37_: _dafny.Map
            d_797_v37_ = _dafny.Map({d_793_v33_: d_796_v36_})
            d_797_v37_ = (d_797_v37_).set(d_793_v33_, d_796_v36_)
            d_798_v38_: _dafny.Map
            d_798_v38_ = _dafny.Map({self.f25: (self).f26})
            d_799_v39_: _dafny.Map
            d_799_v39_ = _dafny.Map({False: (self).f26})
            d_800_v40_: _dafny.Seq
            d_800_v40_ = _dafny.SeqWithoutIsStrInference([len(d_798_v38_), self.f25, self.f25, len(d_799_v39_)])
            d_801_v41_: _dafny.Seq
            d_801_v41_ = _dafny.SeqWithoutIsStrInference([(0) - (self.f25), len(d_800_v40_), -217])
            d_802_v42_: _dafny.Set
            d_802_v42_ = _dafny.Set({(d_801_v41_)[default__.safeIndex(self.f25, len(d_801_v41_))]})
            d_803_v43_: _dafny.Map
            d_803_v43_ = _dafny.Map({(self).f26: len(d_802_v42_)})
            (globalState).f22 = (((d_803_v43_)[(self).f26] if ((self).f26) in (d_803_v43_) else self.f25) if default__.fm0(self.f25, (self).f26, (self).f26, self.f25, globalState) else self.f25)
            d_804_v44_: C1
            nw153_ = C1()
            nw153_.ctor__((self).f24, default__.safeDivisionInt(self.f25, self.f25))
            d_804_v44_ = nw153_
            d_805_v45_: _dafny.Array
            nw154_ = _dafny.Array(int(0), 8)
            d_805_v45_ = nw154_
            index132_ = default__.safeIndex(102, (d_805_v45_).length(0))
            (d_805_v45_)[index132_] = default__.fm5(len(d_793_v33_), self.f25, globalState)
            index133_ = default__.safeIndex(102, (d_805_v45_).length(0))
            (d_805_v45_)[index133_] = (0) - ((0) - ((self.f25) + (default__.safeDivisionInt(self.f25, self.f25))))
        d_806_v46_: D10
        d_806_v46_ = D10_DC21(self.f25, not((self).f26), (self).f26)
        d_807_v47_: _dafny.Array
        def lambda42_(d_808_i1_):
            return _dafny.Map({self.f25: not((self).f26)})

        init23_ = lambda42_
        nw155_ = _dafny.Array(None, 14)
        for i0_23_ in range(nw155_.length(0)):
            nw155_[i0_23_] = init23_(i0_23_)
        d_807_v47_ = nw155_
        source13_ = D4_DC7((d_807_v47_ if (d_806_v46_).cf32 else d_807_v47_))
        if source13_.is_DC8:
            d_809_v48_: _dafny.Array
            nw156_ = _dafny.Array(None, 1)
            nw156_[int(0)] = default__.safeDivisionInt(self.f25, self.f25)
            d_809_v48_ = nw156_
            d_810_v49_: _dafny.Seq
            d_810_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))
            d_811_v50_: _dafny.MultiSet
            d_811_v50_ = _dafny.MultiSet([self.f25, len(d_810_v49_)])
            index134_ = default__.safeIndex(676, (d_809_v48_).length(0))
            (d_809_v48_)[index134_] = ((d_811_v50_).cardinality) - (self.f25)
            index135_ = default__.safeIndex(676, (d_809_v48_).length(0))
            (d_809_v48_)[index135_] = self.f25
            d_812_v51_: D10
            d_812_v51_ = D10_DC19((self).f24)
            d_813_v52_: _dafny.Array
            nw157_ = _dafny.Array(None, 25)
            nw157_[int(0)] = (self).f24
            nw157_[int(1)] = (self).f24
            nw157_[int(2)] = (self).f24
            nw157_[int(3)] = (self).f24
            nw157_[int(4)] = (self).f24
            nw157_[int(5)] = _dafny.CodePoint('q')
            nw157_[int(6)] = (self).f24
            nw157_[int(7)] = (self).f24
            nw157_[int(8)] = (self).f24
            nw157_[int(9)] = default__.fm10((self).f26, (self).f26, (self).f26, (self).f26, globalState)
            nw157_[int(10)] = (self).f24
            nw157_[int(11)] = (self).f24
            nw157_[int(12)] = (self).f24
            nw157_[int(13)] = (self).f24
            nw157_[int(14)] = default__.fm10((self).f26, (self).f26, (self).f26, (self).f26, globalState)
            nw157_[int(15)] = (self).f24
            nw157_[int(16)] = (self).f24
            nw157_[int(17)] = _dafny.CodePoint('u')
            nw157_[int(18)] = (self).f24
            nw157_[int(19)] = (self).f24
            nw157_[int(20)] = (self).f24
            nw157_[int(21)] = (d_812_v51_).cf27
            nw157_[int(22)] = (self).f24
            nw157_[int(23)] = (self).f24
            nw157_[int(24)] = (self).f24
            d_813_v52_ = nw157_
            d_814_v53_: _dafny.Seq
            d_814_v53_ = _dafny.SeqWithoutIsStrInference([d_810_v49_, d_810_v49_, d_810_v49_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kucbeyarw")), (self).fm3((self).f24, globalState)])
            d_815_v54_: _dafny.MultiSet
            d_815_v54_ = _dafny.MultiSet([d_814_v53_, d_814_v53_, d_814_v53_, d_814_v53_, d_814_v53_])
            rhs166_ = d_813_v52_
            rhs167_ = (0) - (((d_815_v54_)[_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f24 for d_816_i2_ in range(default__.abs(920))])])] if (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f24 for d_817_i2_ in range(default__.abs(920))])])) in (d_815_v54_) else self.f25))
            rhs168_ = (self).f26
            lhs142_ = globalState
            d_813_v52_ = rhs166_
            lhs142_.f16 = rhs167_
            r1 = rhs168_
            index136_ = default__.safeIndex(676, (d_809_v48_).length(0))
            (d_809_v48_)[index136_] = default__.safeModuloInt(267, len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_818_i3_ in range(default__.abs(845))])) + (d_810_v49_)))
            d_819_v55_: D3
            d_819_v55_ = D3_DC6((d_809_v48_)[default__.safeIndex(676, (d_809_v48_).length(0))], (d_809_v48_)[default__.safeIndex(676, (d_809_v48_).length(0))], True)
            if (d_819_v55_).cf11:
                d_820_v56_: _dafny.Map
                d_820_v56_ = _dafny.Map({default__.fm0(len(d_810_v49_), (self).f26, (self).f26, self.f25, globalState): d_809_v48_})
                d_809_v48_ = ((d_820_v56_)[(self).f26] if ((self).f26) in (d_820_v56_) else d_809_v48_)
                d_821_v57_: _dafny.Map
                d_821_v57_ = _dafny.Map({(self).f26: not(default__.fm0(self.f25, (self).f26, (self).f26, self.f25, globalState))})
                d_821_v57_ = (d_821_v57_).set(True, (self).f26)
                d_822_v58_: _dafny.Map
                d_822_v58_ = _dafny.Map({(self).f26: d_819_v55_})
                d_823_v59_: D9
                d_823_v59_ = D9_DC16((_dafny.Map({(self).f26: d_819_v55_})) | ((d_822_v58_).set(not((self).f26), default__.fm23(globalState))))
                d_823_v59_ = default__.fm24((self).f26, (d_809_v48_)[default__.safeIndex(676, (d_809_v48_).length(0))], d_810_v49_, globalState)
                d_824_v60_: _dafny.Map
                d_824_v60_ = _dafny.Map({(self).f24: (self).f26})
                d_825_v61_: _dafny.Map
                d_825_v61_ = _dafny.Map({len(d_824_v60_): self.f25})
                d_826_v64_: _dafny.Set
                d_826_v64_ = _dafny.Set({(self).f24, (self).f24, (self).f24, (self).f24})
                d_827_v65_: _dafny.Seq
                d_827_v65_ = _dafny.SeqWithoutIsStrInference([d_826_v64_, d_826_v64_])
                d_828_v66_: _dafny.Map
                d_828_v66_ = _dafny.Map({D10_DC19(_dafny.CodePoint('c')): False})
                d_829_v67_: _dafny.Map
                d_829_v67_ = _dafny.Map({not((self).f26): d_828_v66_})
                index137_ = default__.safeIndex(676, (d_809_v48_).length(0))
                def iife60_():
                    coll36_ = _dafny.Map()
                    compr_36_: int
                    for compr_36_ in (_dafny.Map({self.f25: self.f25})).keys.Elements:
                        d_830_v62_: int = compr_36_
                        if (d_830_v62_) in (_dafny.Map({self.f25: self.f25})):
                            def iife61_():
                                coll37_ = _dafny.Map()
                                compr_37_: int
                                for compr_37_ in (d_811_v50_).Elements:
                                    d_831_v63_: int = compr_37_
                                    if (d_831_v63_) in (d_811_v50_):
                                        coll37_[(d_831_v63_) + (self.f25)] = (self).f26
                                return _dafny.Map(coll37_)
                            coll36_[(d_830_v62_) + (self.f25)] = len((iife61_()
                            ) | (_dafny.Map({len(_dafny.Set({(d_809_v48_)[default__.safeIndex(676, (d_809_v48_).length(0))]})): (self).f26})))
                    return _dafny.Map(coll36_)
                rhs169_ = iife60_()

                rhs170_ = (d_827_v65_) == ((d_827_v65_) + (d_827_v65_))
                rhs171_ = default__.safeModuloInt(len(d_829_v67_), (d_809_v48_)[default__.safeIndex(676, (d_809_v48_).length(0))])
                lhs143_ = globalState
                lhs144_ = d_809_v48_
                lhs145_ = default__.safeIndex(676, (d_809_v48_).length(0))
                d_825_v61_ = rhs169_
                lhs143_.f18 = rhs170_
                lhs144_[lhs145_] = rhs171_
                d_832_v68_: _dafny.Seq
                d_832_v68_ = _dafny.SeqWithoutIsStrInference([(0) - (self.f25), len(d_810_v49_)])
                d_833_v69_: _dafny.Set
                d_833_v69_ = _dafny.Set({(self).f26, (self).f26})
                d_834_v70_: _dafny.MultiSet
                d_834_v70_ = _dafny.MultiSet([(self).f26])
                d_835_v71_: _dafny.Map
                d_835_v71_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([self.f25, len(default__.fm25(len(d_833_v69_), (self).f26, (d_809_v48_)[default__.safeIndex(676, (d_809_v48_).length(0))], globalState))]): d_834_v70_})
                d_836_v72_: _dafny.Seq
                d_836_v72_ = _dafny.SeqWithoutIsStrInference([(self).f26])
                d_837_v73_: _dafny.Map
                d_837_v73_ = _dafny.Map({(d_809_v48_)[default__.safeIndex(676, (d_809_v48_).length(0))]: (self).f26})
                index138_ = default__.safeIndex(676, (d_809_v48_).length(0))
                rhs172_ = 33
                rhs173_ = ((d_825_v61_)[(0) - (((d_809_v48_)[default__.safeIndex(676, (d_809_v48_).length(0))]) * (len(d_832_v68_)))] if ((0) - (((d_809_v48_)[default__.safeIndex(676, (d_809_v48_).length(0))]) * (len(d_832_v68_)))) in (d_825_v61_) else ((((d_835_v71_)[default__.fm26(globalState)] if (default__.fm26(globalState)) in (d_835_v71_) else d_834_v70_)).intersection(_dafny.MultiSet(d_836_v72_))).cardinality)
                rhs174_ = ((d_837_v73_)[((d_834_v70_)[(self).f26] if ((self).f26) in (d_834_v70_) else (0) - ((d_809_v48_)[default__.safeIndex(676, (d_809_v48_).length(0))]))] if (((d_834_v70_)[(self).f26] if ((self).f26) in (d_834_v70_) else (0) - ((d_809_v48_)[default__.safeIndex(676, (d_809_v48_).length(0))]))) in (d_837_v73_) else not(((d_821_v57_)[(self).f26] if ((self).f26) in (d_821_v57_) else (self).f26)))
                lhs146_ = globalState
                lhs147_ = d_809_v48_
                lhs148_ = default__.safeIndex(676, (d_809_v48_).length(0))
                lhs149_ = globalState
                lhs146_.f2 = rhs172_
                lhs147_[lhs148_] = rhs173_
                lhs149_.f18 = rhs174_
            elif True:
                d_838_v74_: _dafny.Array
                def lambda43_(d_839_i4_):
                    return _dafny.SeqWithoutIsStrInference([self.f25, self.f25])

                init24_ = lambda43_
                nw158_ = _dafny.Array(None, 27)
                for i0_24_ in range(nw158_.length(0)):
                    nw158_[i0_24_] = init24_(i0_24_)
                d_838_v74_ = nw158_
                d_840_v75_: _dafny.Map
                d_840_v75_ = _dafny.Map({d_838_v74_: True})
                (globalState).f18 = ((d_840_v75_)[d_838_v74_] if (d_838_v74_) in (d_840_v75_) else (self).f26)
                rhs175_ = _dafny.SeqWithoutIsStrInference([d_810_v49_ for d_841_i5_ in range(default__.abs(913))])
                rhs176_ = (self).f26
                d_814_v53_ = rhs175_
                r1 = rhs176_
                (globalState).f2 = self.f25
                d_842_v76_: C0
                nw159_ = C0()
                nw159_.ctor__()
                d_842_v76_ = nw159_
                d_842_v76_ = d_842_v76_
                d_843_v77_: _dafny.Seq
                d_843_v77_ = _dafny.SeqWithoutIsStrInference([not((self).f26)])
                d_844_v78_: _dafny.Array
                nw160_ = _dafny.Array(None, 11)
                nw160_[int(0)] = d_843_v77_
                nw160_[int(1)] = default__.fm7(globalState)
                nw160_[int(2)] = d_843_v77_
                nw160_[int(3)] = d_843_v77_
                nw160_[int(4)] = ((d_843_v77_).set(default__.safeIndex((d_809_v48_)[default__.safeIndex(676, (d_809_v48_).length(0))], len(d_843_v77_)), (self).f26)) + (d_843_v77_)
                nw160_[int(5)] = _dafny.SeqWithoutIsStrInference([not((self).f26), (self).f26, (self).f26, (self).f26, (self).f26])
                nw160_[int(6)] = d_843_v77_
                nw160_[int(7)] = _dafny.SeqWithoutIsStrInference([False, (self).f26])
                nw160_[int(8)] = d_843_v77_
                nw160_[int(9)] = ((_dafny.SeqWithoutIsStrInference([(self).f26, (self).f26])).set(default__.safeIndex(-115, len(_dafny.SeqWithoutIsStrInference([(self).f26, (self).f26]))), (self).f26)) + (d_843_v77_)
                nw160_[int(10)] = (d_843_v77_) + (d_843_v77_)
                d_844_v78_ = nw160_
                index139_ = default__.safeIndex(568, (d_844_v78_).length(0))
                (d_844_v78_)[index139_] = d_843_v77_
                index140_ = default__.safeIndex(568, (d_844_v78_).length(0))
                (d_844_v78_)[index140_] = d_843_v77_
        elif True:
            d_845___mcc_h0_ = source13_.cf12
            d_846_cf12_ = d_845___mcc_h0_
            (globalState).f19 = _dafny.SeqWithoutIsStrInference([len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uuvoe"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "knnyuorh")))) for d_847_i6_ in range(default__.abs(323))])
            d_848_v79_: _dafny.Array
            def lambda44_(d_849_i7_):
                return default__.safeDivisionInt(d_849_i7_, len(_dafny.Set({(self).f26, (self).f26})))

            init25_ = lambda44_
            nw161_ = _dafny.Array(None, 2)
            for i0_25_ in range(nw161_.length(0)):
                nw161_[i0_25_] = init25_(i0_25_)
            d_848_v79_ = nw161_
            d_848_v79_ = d_848_v79_
            r1 = (self).f26
            if (self).f26:
                d_850_v80_: _dafny.Seq
                d_850_v80_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ksmojox"))
                (globalState).f10 = len(d_850_v80_)
                d_851_v81_: _dafny.Seq
                d_851_v81_ = _dafny.SeqWithoutIsStrInference([(self).f26, not((self).f26)])
                d_852_v82_: _dafny.Array
                nw162_ = _dafny.Array(None, 10)
                nw162_[int(0)] = (502) > (self.f25)
                nw162_[int(1)] = True
                nw162_[int(2)] = (self).f26
                nw162_[int(3)] = (self).f26
                nw162_[int(4)] = (self).f26
                nw162_[int(5)] = (self).f26
                nw162_[int(6)] = True
                nw162_[int(7)] = ((self).f26 if (self).f26 else (self).f26)
                nw162_[int(8)] = (self).f26
                nw162_[int(9)] = (d_851_v81_) == (d_851_v81_)
                d_852_v82_ = nw162_
                index141_ = default__.safeIndex(344, (d_852_v82_).length(0))
                (d_852_v82_)[index141_] = (self).f26
                index142_ = default__.safeIndex(344, (d_852_v82_).length(0))
                (d_852_v82_)[index142_] = (self).f26
                d_853_v83_: _dafny.Array
                nw163_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 6)
                d_853_v83_ = nw163_
                index143_ = default__.safeIndex(195, (d_853_v83_).length(0))
                (d_853_v83_)[index143_] = (d_850_v80_) + (d_850_v80_)
                index144_ = default__.safeIndex(195, (d_853_v83_).length(0))
                rhs177_ = d_852_v82_
                rhs178_ = d_850_v80_
                lhs150_ = d_853_v83_
                lhs151_ = default__.safeIndex(195, (d_853_v83_).length(0))
                d_852_v82_ = rhs177_
                lhs150_[lhs151_] = rhs178_
                d_854_v84_: C4
                nw164_ = C4()
                nw164_.ctor__(default__.fm27(self.f25, globalState), default__.fm0(self.f25, (self).f26, False, self.f25, globalState), (self).f24, self.f25)
                d_854_v84_ = nw164_
                d_855_v85_: C3
                nw165_ = C3()
                nw165_.ctor__((self).f26, default__.fm10((d_854_v84_).f30, (d_854_v84_).f30, (self).f26, (d_852_v82_)[default__.safeIndex(344, (d_852_v82_).length(0))], globalState), default__.safeModuloInt(len(d_850_v80_), self.f25))
                d_855_v85_ = nw165_
            elif True:
                d_856_v86_: _dafny.Seq
                d_856_v86_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xukawq"))
                d_857_v87_: _dafny.MultiSet
                d_857_v87_ = _dafny.MultiSet([d_856_v86_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ylh")), d_856_v86_, d_856_v86_, d_856_v86_])
                d_858_v88_: _dafny.Seq
                d_858_v88_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qxgte"))
                d_859_v89_: _dafny.Seq
                d_859_v89_ = _dafny.SeqWithoutIsStrInference([(self).f26, (self).f26])
                d_860_v90_: C3
                nw166_ = C3()
                nw166_.ctor__(not((d_857_v87_).isdisjoint(_dafny.MultiSet([(d_858_v88_)]))), (self).f24, len(d_859_v89_))
                d_860_v90_ = nw166_
                d_861_v91_: _dafny.Array
                def lambda45_(d_862_v86_):
                    def lambda46_(d_863_i8_):
                        return d_862_v86_

                    return lambda46_

                init26_ = lambda45_(d_856_v86_)
                nw167_ = _dafny.Array(None, 26)
                for i0_26_ in range(nw167_.length(0)):
                    nw167_[i0_26_] = init26_(i0_26_)
                d_861_v91_ = nw167_
                index145_ = default__.safeIndex(692, (d_861_v91_).length(0))
                (d_861_v91_)[index145_] = d_856_v86_
                index146_ = default__.safeIndex(692, (d_861_v91_).length(0))
                (d_861_v91_)[index146_] = (d_856_v86_) + ((d_856_v86_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ixxmbm"))))
                (globalState).f22 = len(d_856_v86_)
                (globalState).f3 = d_856_v86_
                (globalState).f18 = (self).f26
        d_864_v92_: _dafny.Array
        nw168_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 6)
        d_864_v92_ = nw168_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_864_v92_).length(0)):
            d_865_i9_: int = guard_loop_4_
            if (True) and (((0) <= (d_865_i9_)) and ((d_865_i9_) < ((d_864_v92_).length(0)))):
                (d_864_v92_)[(d_865_i9_)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gvibnejd"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "polnxwjt")))
        d_866_v94_: _dafny.Array
        def lambda47_(d_867_i10_):
            def iife62_():
                coll38_ = _dafny.Map()
                compr_38_: str
                for compr_38_ in (_dafny.Map({(self).f24: (self).f26})).keys.Elements:
                    d_868_v93_: str = compr_38_
                    if (d_868_v93_) in (_dafny.Map({(self).f24: (self).f26})):
                        coll38_[d_868_v93_] = (self).f26
                return _dafny.Map(coll38_)
            return (_dafny.Map({(self).f24: (self).f26})) | (iife62_()
            )

        init27_ = lambda47_
        nw169_ = _dafny.Array(None, 7)
        for i0_27_ in range(nw169_.length(0)):
            nw169_[i0_27_] = init27_(i0_27_)
        d_866_v94_ = nw169_
        d_866_v94_ = d_866_v94_
        d_869_v95_: _dafny.Array
        nw170_ = _dafny.Array(False, 21)
        d_869_v95_ = nw170_
        index147_ = default__.safeIndex(908, (d_869_v95_).length(0))
        (d_869_v95_)[index147_] = (self).f26
        d_870_v96_: _dafny.Set
        d_870_v96_ = _dafny.Set({(self).f26, (self).f26})
        d_871_v97_: _dafny.Map
        d_871_v97_ = _dafny.Map({d_870_v96_: self.f25})
        index148_ = default__.safeIndex(908, (d_869_v95_).length(0))
        (d_869_v95_)[index148_] = (default__.fm0(len(d_871_v97_), (self).f26, False, self.f25, globalState) if (self).f26 else (self).f26)
        d_872_v98_: _dafny.Map
        d_872_v98_ = _dafny.Map({(self.f25) * (self.f25): 198})
        d_872_v98_ = (d_872_v98_).set(self.f25, default__.safeModuloInt(self.f25, self.f25))
        d_873_v99_: _dafny.Map
        d_873_v99_ = _dafny.Map({d_869_v95_: _dafny.Map({False: True})})
        r0 = (d_873_v99_) | ((d_873_v99_) | (d_873_v99_))
        r1 = False
        return r0, r1

    @property
    def f26(self):
        return self._f26

class C6:
    def  __init__(self):
        self.f23: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    def ctor__(self, f23):
        (self).f23 = f23

    def fm1(self, p0, globalState):
        return ((_dafny.Map({False: -509})) | (_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "clwiitdbt")))}))) != ((_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_874_i0_ in range(default__.abs(-196))]))]))})) | (_dafny.Map({not(False): -610})))

    def fm2(self, globalState):
        def iife63_():
            coll39_ = _dafny.Map()
            compr_39_: _dafny.Map
            for compr_39_ in (_dafny.MultiSet([self.f23])).Elements:
                d_876_v0_: _dafny.Map = compr_39_
                if (d_876_v0_) in (_dafny.MultiSet([self.f23])):
                    coll39_[d_876_v0_] = (0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yf"))), (0) - (len(_dafny.Map({True: _dafny.CodePoint('t')})))])))
            return _dafny.Map(coll39_)
        return (len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_875_i0_ in range(default__.abs(339))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rideqr"))))) <= ((-714) + (len(iife63_()
        )))

    def m0(self, p0, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_877_v0_: bool
        d_877_v0_ = True
        if d_877_v0_:
            (globalState).f2 = default__.safeDivisionInt(default__.safeDivisionInt(p0, p0), 203)
            d_878_v1_: _dafny.Seq
            d_878_v1_ = _dafny.SeqWithoutIsStrInference([d_877_v0_])
            if (d_878_v1_)[default__.safeIndex(p0, len(d_878_v1_))]:
                (globalState).f19 = _dafny.SeqWithoutIsStrInference([(p0) * (p0) for d_879_i0_ in range(default__.abs(806))])
                (globalState).f10 = (p0) * ((p0) - (p0))
                d_880_v2_: _dafny.Map
                d_880_v2_ = _dafny.Map({((self.f23)[p0] if (p0) in (self.f23) else d_877_v0_): d_877_v0_})
                d_881_v3_: D2
                d_881_v3_ = D2_DC2(d_877_v0_)
                d_882_v4_: _dafny.Seq
                d_882_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yxdeimc"))
                d_883_v5_: str
                d_883_v5_ = _dafny.CodePoint('n')
                d_884_v6_: _dafny.Array
                nw171_ = _dafny.Array(None, 15)
                nw171_[int(0)] = True
                nw171_[int(1)] = d_877_v0_
                nw171_[int(2)] = ((d_880_v2_)[d_877_v0_] if (d_877_v0_) in (d_880_v2_) else d_877_v0_)
                nw171_[int(3)] = d_877_v0_
                nw171_[int(4)] = d_877_v0_
                nw171_[int(5)] = d_877_v0_
                nw171_[int(6)] = d_877_v0_
                nw171_[int(7)] = (d_881_v3_).cf2
                nw171_[int(8)] = not(not((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lkxsn"))) == (d_882_v4_)))
                nw171_[int(9)] = True
                nw171_[int(10)] = (d_877_v0_ if d_877_v0_ else True)
                nw171_[int(11)] = (d_883_v5_) not in (d_882_v4_)
                nw171_[int(12)] = d_877_v0_
                nw171_[int(13)] = d_877_v0_
                nw171_[int(14)] = d_877_v0_
                d_884_v6_ = nw171_
                d_885_v7_: _dafny.Seq
                d_885_v7_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                index149_ = default__.safeIndex(590, (d_884_v6_).length(0))
                (d_884_v6_)[index149_] = (d_885_v7_) != (d_885_v7_)
                index150_ = default__.safeIndex(590, (d_884_v6_).length(0))
                (d_884_v6_)[index150_] = (p0) <= (p0)
                d_886_v9_: _dafny.Set
                def iife64_():
                    coll40_ = _dafny.Map()
                    compr_40_: int
                    for compr_40_ in _dafny.IntegerRange(807, 609):
                        d_887_v8_: int = compr_40_
                        if ((807) <= (d_887_v8_)) and ((d_887_v8_) < (609)):
                            coll40_[default__.safeModuloInt(d_887_v8_, p0)] = (d_884_v6_)[default__.safeIndex(590, (d_884_v6_).length(0))]
                    return _dafny.Map(coll40_)
                d_886_v9_ = _dafny.Set({(_dafny.Map({p0: (d_884_v6_)[default__.safeIndex(590, (d_884_v6_).length(0))]}) if d_877_v0_ else self.f23), self.f23, self.f23, iife64_()
                , self.f23})
                d_888_v10_: _dafny.Array
                nw172_ = _dafny.Array(_dafny.Array(None, 0), 8)
                d_888_v10_ = nw172_
                d_889_v11_: _dafny.Array
                def lambda48_(d_890_v5_):
                    def lambda49_(d_891_i1_):
                        return d_890_v5_

                    return lambda49_

                init28_ = lambda48_(d_883_v5_)
                nw173_ = _dafny.Array(None, 18)
                for i0_28_ in range(nw173_.length(0)):
                    nw173_[i0_28_] = init28_(i0_28_)
                d_889_v11_ = nw173_
                index151_ = default__.safeIndex(953, (d_888_v10_).length(0))
                (d_888_v10_)[index151_] = d_889_v11_
                d_892_v12_: _dafny.Seq
                d_892_v12_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_893_i2_ in range(default__.abs(218))])
                d_894_v13_: _dafny.MultiSet
                d_894_v13_ = _dafny.MultiSet([d_892_v12_, d_892_v12_])
                d_895_v14_: _dafny.Seq
                d_895_v14_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).fm1(d_894_v13_, globalState), d_877_v0_, False])])
                d_896_v15_: _dafny.Map
                d_896_v15_ = _dafny.Map({p0: (d_895_v14_)[default__.safeIndex((0) - (p0), len(d_895_v14_))]})
                index152_ = default__.safeIndex(953, (d_888_v10_).length(0))
                index153_ = default__.safeIndex(590, (d_884_v6_).length(0))
                rhs179_ = (d_886_v9_).intersection((D3_DC4(d_886_v9_)).cf5)
                rhs180_ = d_889_v11_
                rhs181_ = (d_885_v7_) != (d_885_v7_)
                rhs182_ = d_877_v0_
                rhs183_ = (d_878_v1_) + (((d_896_v15_)[p0] if (p0) in (d_896_v15_) else _dafny.SeqWithoutIsStrInference([(d_884_v6_)[default__.safeIndex(590, (d_884_v6_).length(0))], (d_884_v6_)[default__.safeIndex(590, (d_884_v6_).length(0))]])))
                lhs152_ = d_888_v10_
                lhs153_ = default__.safeIndex(953, (d_888_v10_).length(0))
                lhs154_ = d_884_v6_
                lhs155_ = default__.safeIndex(590, (d_884_v6_).length(0))
                d_886_v9_ = rhs179_
                lhs152_[lhs153_] = rhs180_
                lhs154_[lhs155_] = rhs181_
                d_877_v0_ = rhs182_
                d_878_v1_ = rhs183_
                rhs184_ = (p0) + (p0)
                rhs185_ = d_885_v7_
                lhs156_ = globalState
                lhs156_.f2 = rhs184_
                d_885_v7_ = rhs185_
            elif True:
                d_897_v16_: D11
                d_897_v16_ = D11_DC24(d_877_v0_, p0)
                d_898_v17_: _dafny.Map
                d_898_v17_ = _dafny.Map({d_897_v16_: d_877_v0_})
                d_899_v18_: str
                d_899_v18_ = _dafny.CodePoint('p')
                d_900_v19_: _dafny.Seq
                d_900_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "csjg"))
                d_901_v20_: C4
                nw174_ = C4()
                nw174_.ctor__(_dafny.Map({_dafny.MultiSet([False]): 707}), ((d_898_v17_)[d_897_v16_] if (d_897_v16_) in (d_898_v17_) else d_877_v0_), d_899_v18_, default__.safeModuloInt(len(d_900_v19_), p0))
                d_901_v20_ = nw174_
                d_902_v21_: _dafny.Map
                d_902_v21_ = _dafny.Map({(0) - ((0) - (p0)): _dafny.CodePoint('k')})
                d_902_v21_ = (d_902_v21_).set((0) - (default__.safeDivisionInt(p0, p0)), _dafny.CodePoint('e'))
                d_903_v22_: _dafny.Set
                d_903_v22_ = _dafny.Set({p0})
                d_904_v23_: D12
                d_904_v23_ = D12_DC26(d_903_v22_)
                r2 = default__.fm15(p0, (0) - (p0), len((d_904_v23_).cf39), globalState)
                d_905_v24_: _dafny.Array
                nw175_ = _dafny.Array(_dafny.Array(None, 0), 2)
                d_905_v24_ = nw175_
                d_905_v24_ = d_905_v24_
                d_906_v25_: D7
                d_906_v25_ = D7_DC13(_dafny.Map({703: d_877_v0_}))
                d_907_v26_: _dafny.Map
                d_907_v26_ = _dafny.Map({d_906_v25_: d_878_v1_})
                d_908_v27_: _dafny.Map
                d_908_v27_ = _dafny.Map({(d_901_v20_).f30: (d_907_v26_) | (_dafny.Map({d_906_v25_: d_878_v1_}))})
                d_908_v27_ = (d_908_v27_).set((d_901_v20_).f30, (d_907_v26_).set(d_906_v25_, d_878_v1_))
            d_909_v28_: _dafny.Array
            def lambda50_(d_910_v1_):
                def lambda51_(d_911_i3_):
                    return d_910_v1_

                return lambda51_

            init29_ = lambda50_(d_878_v1_)
            nw176_ = _dafny.Array(None, 16)
            for i0_29_ in range(nw176_.length(0)):
                nw176_[i0_29_] = init29_(i0_29_)
            d_909_v28_ = nw176_
            d_912_v29_: _dafny.Seq
            d_912_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "telw"))
            d_913_v30_: _dafny.Seq
            d_913_v30_ = _dafny.SeqWithoutIsStrInference([d_878_v1_])
            nw177_ = _dafny.Array(None, 22)
            nw177_[int(0)] = d_878_v1_
            nw177_[int(1)] = d_878_v1_
            nw177_[int(2)] = _dafny.SeqWithoutIsStrInference([d_877_v0_, d_877_v0_, d_877_v0_])
            nw177_[int(3)] = d_878_v1_
            nw177_[int(4)] = d_878_v1_
            nw177_[int(5)] = (d_878_v1_) + ((d_878_v1_).set(default__.safeIndex(len(d_912_v29_), len(d_878_v1_)), d_877_v0_))
            nw177_[int(6)] = d_878_v1_
            nw177_[int(7)] = d_878_v1_
            nw177_[int(8)] = d_878_v1_
            nw177_[int(9)] = _dafny.SeqWithoutIsStrInference([True, d_877_v0_, d_877_v0_, not(((self.f23)[len(d_878_v1_)] if (len(d_878_v1_)) in (self.f23) else d_877_v0_))])
            nw177_[int(10)] = d_878_v1_
            nw177_[int(11)] = d_878_v1_
            nw177_[int(12)] = (d_878_v1_) + (d_878_v1_)
            nw177_[int(13)] = d_878_v1_
            nw177_[int(14)] = ((d_913_v30_)[default__.safeIndex(p0, len(d_913_v30_))]) + (d_878_v1_)
            nw177_[int(15)] = d_878_v1_
            nw177_[int(16)] = (d_878_v1_) + (d_878_v1_)
            nw177_[int(17)] = _dafny.SeqWithoutIsStrInference([d_877_v0_])
            nw177_[int(18)] = (d_878_v1_) + (d_878_v1_)
            nw177_[int(19)] = default__.fm7(globalState)
            nw177_[int(20)] = d_878_v1_
            nw177_[int(21)] = d_878_v1_
            d_909_v28_ = nw177_
            d_914_v31_: _dafny.Seq
            d_914_v31_ = _dafny.SeqWithoutIsStrInference([(0) - (p0), len(_dafny.SeqWithoutIsStrInference([d_877_v0_, False]))])
            rhs186_ = default__.safeDivisionInt((d_914_v31_)[default__.safeIndex(p0, len(d_914_v31_))], p0)
            rhs187_ = d_877_v0_
            rhs188_ = d_877_v0_
            lhs157_ = globalState
            lhs158_ = globalState
            lhs157_.f10 = rhs186_
            lhs158_.f18 = rhs187_
            d_877_v0_ = rhs188_
            r1 = p0
        elif True:
            d_915_v32_: _dafny.Seq
            d_915_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kysw"))
            d_916_v33_: D9
            d_916_v33_ = D9_DC17(len(d_915_v32_))
            d_917_v34_: _dafny.Seq
            d_917_v34_ = _dafny.SeqWithoutIsStrInference([p0])
            d_918_v35_: _dafny.Map
            d_918_v35_ = _dafny.Map({d_916_v33_: (p0) in (d_917_v34_)})
            d_918_v35_ = ((_dafny.Map({D9_DC17(p0): d_877_v0_})) | (d_918_v35_)) | (default__.fm28(p0, not(d_877_v0_), d_877_v0_, globalState))
            d_877_v0_ = False
            d_919_v36_: _dafny.MultiSet
            d_919_v36_ = _dafny.MultiSet([d_877_v0_, d_877_v0_, True, d_877_v0_])
            d_920_v37_: _dafny.Map
            d_920_v37_ = _dafny.Map({d_919_v36_: p0})
            d_921_v38_: str
            d_921_v38_ = _dafny.CodePoint('u')
            d_922_v39_: T0
            nw178_ = C4()
            nw178_.ctor__(d_920_v37_, d_877_v0_, d_921_v38_, p0)
            d_922_v39_ = nw178_
            d_923_v40_: _dafny.Seq
            d_923_v40_ = _dafny.SeqWithoutIsStrInference([d_922_v39_, d_922_v39_, d_922_v39_, d_922_v39_])
            d_924_v41_: _dafny.Seq
            d_924_v41_ = d_923_v40_
            d_925_v42_: _dafny.Set
            d_925_v42_ = _dafny.Set({d_924_v41_})
            (globalState).f18 = not(not((d_925_v42_).isdisjoint(d_925_v42_)))
            d_926_v44_: _dafny.Seq
            d_926_v44_ = _dafny.SeqWithoutIsStrInference([d_919_v36_])
            d_927_v45_: C4
            nw179_ = C4()
            def iife65_():
                coll41_ = _dafny.Map()
                compr_41_: _dafny.MultiSet
                for compr_41_ in (d_926_v44_).Elements:
                    d_928_v43_: _dafny.MultiSet = compr_41_
                    if (d_928_v43_) in (d_926_v44_):
                        coll41_[d_928_v43_] = 943
                return _dafny.Map(coll41_)
            nw179_.ctor__((iife65_()
            ) | (_dafny.Map({d_919_v36_: p0})), d_877_v0_, default__.fm10(not(d_877_v0_), d_877_v0_, d_877_v0_, d_877_v0_, globalState), p0)
            d_927_v45_ = nw179_
            d_929_v46_: C0
            nw180_ = C0()
            nw180_.ctor__()
            d_929_v46_ = nw180_
            d_930_v47_: _dafny.Map
            d_930_v47_ = _dafny.Map({(d_927_v45_).f30: d_929_v46_})
            rhs189_ = p0
            rhs190_ = (d_922_v39_).f24
            rhs191_ = d_927_v45_
            rhs192_ = ((d_930_v47_)[d_877_v0_] if (d_877_v0_) in (d_930_v47_) else d_929_v46_)
            lhs159_ = globalState
            lhs159_.f11 = rhs189_
            d_921_v38_ = rhs190_
            d_927_v45_ = rhs191_
            d_929_v46_ = rhs192_
            (globalState).f10 = len((d_915_v32_) + ((d_915_v32_) + (_dafny.SeqWithoutIsStrInference([d_921_v38_ for d_931_i4_ in range(default__.abs(564))]))))
        if d_877_v0_:
            d_932_v48_: _dafny.Seq
            d_932_v48_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sepaclhxe"))
            d_933_v49_: str
            d_933_v49_ = _dafny.CodePoint('g')
            d_934_v50_: _dafny.Map
            d_934_v50_ = _dafny.Map({p0: d_933_v49_})
            d_935_v51_: C5
            nw181_ = C5()
            nw181_.ctor__(default__.fm0(len(d_932_v48_), not(d_877_v0_), d_877_v0_, (0) - (p0), globalState), ((d_934_v50_)[p0] if (p0) in (d_934_v50_) else d_933_v49_), p0)
            d_935_v51_ = nw181_
            d_936_v52_: _dafny.Set
            d_936_v52_ = _dafny.Set({((0) - (p0)) > (p0)})
            d_936_v52_ = d_936_v52_
            d_937_v53_: _dafny.Array
            nw182_ = _dafny.Array(D9.default()(), 18)
            d_937_v53_ = nw182_
            d_938_v54_: _dafny.Map
            d_938_v54_ = _dafny.Map({(d_935_v51_).f26: D3_DC6(p0, p0, False)})
            d_939_v55_: D9
            d_939_v55_ = D9_DC16(d_938_v54_)
            index154_ = default__.safeIndex(705, (d_937_v53_).length(0))
            (d_937_v53_)[index154_] = d_939_v55_
            index155_ = default__.safeIndex(705, (d_937_v53_).length(0))
            (d_937_v53_)[index155_] = D9_DC16((D9_DC16(d_938_v54_)).cf24)
            (globalState).f18 = not(not(d_877_v0_))
            d_940_v56_: _dafny.Array
            nw183_ = _dafny.Array(None, 2)
            nw183_[int(0)] = d_933_v49_
            nw183_[int(1)] = d_933_v49_
            d_940_v56_ = nw183_
            index156_ = default__.safeIndex(842, (d_940_v56_).length(0))
            (d_940_v56_)[index156_] = d_933_v49_
            index157_ = default__.safeIndex(842, (d_940_v56_).length(0))
            (d_940_v56_)[index157_] = d_933_v49_
        elif True:
            d_941_v57_: _dafny.MultiSet
            d_941_v57_ = _dafny.MultiSet([p0, p0])
            d_942_v58_: D7
            d_942_v58_ = D7_DC14(d_877_v0_, p0, p0)
            source14_ = (d_942_v58_ if not((d_941_v57_) == (d_941_v57_)) else d_942_v58_)
            if source14_.is_DC14:
                d_943___mcc_h0_ = source14_.cf20
                d_944___mcc_h1_ = source14_.cf21
                d_945___mcc_h2_ = source14_.cf22
                d_946_cf22_ = d_945___mcc_h2_
                d_947_cf21_ = d_944___mcc_h1_
                d_948_cf20_ = d_943___mcc_h0_
                d_949_v59_: str
                d_949_v59_ = _dafny.CodePoint('w')
                d_949_v59_ = (d_949_v59_ if (d_948_cf20_) or (d_877_v0_) else d_949_v59_)
                d_949_v59_ = (d_949_v59_ if False else d_949_v59_)
                (globalState).f18 = not (False) or ((d_948_cf20_) and (d_877_v0_))
                d_950_v60_: _dafny.Array
                nw184_ = _dafny.Array(_dafny.Array(None, 0), 15)
                d_950_v60_ = nw184_
                d_951_v61_: _dafny.MultiSet
                d_951_v61_ = _dafny.MultiSet([d_948_cf20_])
                d_952_v62_: _dafny.Map
                d_952_v62_ = _dafny.Map({d_951_v61_: d_946_cf22_})
                d_953_v63_: C4
                nw185_ = C4()
                nw185_.ctor__(d_952_v62_, d_877_v0_, d_949_v59_, d_946_cf22_)
                d_953_v63_ = nw185_
                d_954_v64_: _dafny.Map
                d_954_v64_ = _dafny.Map({d_953_v63_: (d_953_v63_).f30})
                d_955_v65_: _dafny.Map
                d_955_v65_ = _dafny.Map({596: d_954_v64_})
                d_956_v66_: _dafny.Seq
                d_956_v66_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))
                d_957_v67_: _dafny.Set
                d_957_v67_ = _dafny.Set({440})
                d_958_v68_: _dafny.Seq
                d_958_v68_ = _dafny.SeqWithoutIsStrInference([(d_953_v63_).f30, True, (d_953_v63_).f30, d_877_v0_, d_948_cf20_])
                d_959_v69_: _dafny.Array
                nw186_ = _dafny.Array(None, 8)
                nw186_[int(0)] = len(_dafny.SeqWithoutIsStrInference([False, d_948_cf20_, d_948_cf20_]))
                nw186_[int(1)] = p0
                nw186_[int(2)] = (_dafny.MultiSet([d_948_cf20_])).cardinality
                nw186_[int(3)] = len(d_955_v65_)
                nw186_[int(4)] = len(d_956_v66_)
                nw186_[int(5)] = len(d_957_v67_)
                nw186_[int(6)] = default__.fm17(d_877_v0_, (d_953_v63_).f30, default__.fm0(len(d_958_v68_), d_948_cf20_, d_948_cf20_, p0, globalState), globalState)
                nw186_[int(7)] = len(d_958_v68_)
                d_959_v69_ = nw186_
                index158_ = default__.safeIndex(669, (d_950_v60_).length(0))
                (d_950_v60_)[index158_] = d_959_v69_
                d_960_v70_: D5
                d_960_v70_ = D5_DC9(d_959_v69_)
                pat_let_tv9_ = d_959_v69_
                index159_ = default__.safeIndex(669, (d_950_v60_).length(0))
                def iife66_(_pat_let12_0):
                    def iife67_(d_961_dt__update__tmp_h0_):
                        def iife68_(_pat_let13_0):
                            def iife69_(d_962_dt__update_hcf13_h0_):
                                return D5_DC9(d_962_dt__update_hcf13_h0_)
                            return iife69_(_pat_let13_0)
                        return iife68_(pat_let_tv9_)
                    return iife67_(_pat_let12_0)
                (d_950_v60_)[index159_] = (iife66_(d_960_v70_)).cf13
            elif True:
                d_963___mcc_h3_ = source14_.cf19
                d_964_cf19_ = d_963___mcc_h3_
                d_965_v71_: _dafny.Array
                nw187_ = _dafny.Array(None, 4)
                nw187_[int(0)] = d_877_v0_
                nw187_[int(1)] = d_877_v0_
                nw187_[int(2)] = d_877_v0_
                nw187_[int(3)] = d_877_v0_
                d_965_v71_ = nw187_
                d_966_v72_: _dafny.Map
                d_966_v72_ = _dafny.Map({d_877_v0_: p0})
                d_967_v73_: _dafny.Map
                d_967_v73_ = _dafny.Map({p0: p0})
                d_968_v74_: _dafny.Map
                d_968_v74_ = _dafny.Map({d_965_v71_: (_dafny.Map({len(d_966_v72_): len(d_967_v73_)})) | (d_967_v73_)})
                d_969_v75_: _dafny.Array
                nw188_ = _dafny.Array(_dafny.Array(None, 0), 19)
                d_969_v75_ = nw188_
                d_970_v76_: _dafny.Array
                def lambda52_(d_971_i5_):
                    return _dafny.CodePoint('q')

                init30_ = lambda52_
                nw189_ = _dafny.Array(None, 14)
                for i0_30_ in range(nw189_.length(0)):
                    nw189_[i0_30_] = init30_(i0_30_)
                d_970_v76_ = nw189_
                index160_ = default__.safeIndex(903, (d_969_v75_).length(0))
                (d_969_v75_)[index160_] = d_970_v76_
                d_972_v77_: _dafny.Seq
                d_972_v77_ = _dafny.SeqWithoutIsStrInference([d_968_v74_, d_968_v74_, _dafny.Map({d_965_v71_: (d_967_v73_).set(p0, (D11_DC24(default__.fm0(p0, d_877_v0_, d_877_v0_, p0, globalState), p0)).cf37)}), d_968_v74_])
                d_973_v78_: _dafny.Seq
                d_973_v78_ = _dafny.SeqWithoutIsStrInference([(d_968_v74_) | (d_968_v74_), (d_968_v74_) | (d_968_v74_), (d_972_v77_)[default__.safeIndex((0) - (p0), len(d_972_v77_))], d_968_v74_])
                d_974_v79_: _dafny.Array
                nw190_ = _dafny.Array(None, 12)
                nw190_[int(0)] = p0
                nw190_[int(1)] = p0
                nw190_[int(2)] = p0
                nw190_[int(3)] = p0
                nw190_[int(4)] = p0
                nw190_[int(5)] = 838
                nw190_[int(6)] = p0
                nw190_[int(7)] = p0
                nw190_[int(8)] = p0
                nw190_[int(9)] = p0
                nw190_[int(10)] = p0
                nw190_[int(11)] = p0
                d_974_v79_ = nw190_
                d_975_v80_: D13
                d_975_v80_ = D13_DC28(_dafny.Set({d_974_v79_}))
                d_976_v81_: D5
                d_976_v81_ = D5_DC9(d_974_v79_)
                d_977_v82_: _dafny.Set
                d_977_v82_ = _dafny.Set({d_974_v79_, (d_976_v81_).cf13, d_974_v79_, d_974_v79_})
                d_978_v83_: _dafny.MultiSet
                d_978_v83_ = _dafny.MultiSet([d_877_v0_])
                d_979_v84_: _dafny.Seq
                d_979_v84_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sc"))
                d_980_v85_: D12
                d_980_v85_ = D12_DC27(((d_941_v57_)[p0] if (p0) in (d_941_v57_) else ((d_978_v83_)[d_877_v0_] if (d_877_v0_) in (d_978_v83_) else len(d_979_v84_))), d_877_v0_)
                d_981_v86_: _dafny.Map
                d_981_v86_ = _dafny.Map({False: d_877_v0_})
                d_982_v87_: _dafny.Map
                d_982_v87_ = _dafny.Map({p0: d_981_v86_})
                index161_ = default__.safeIndex(903, (d_969_v75_).length(0))
                rhs193_ = (d_973_v78_)[default__.safeIndex(p0, len(d_973_v78_))]
                rhs194_ = d_970_v76_
                rhs195_ = p0
                rhs196_ = ((d_975_v80_).cf42).isdisjoint(d_977_v82_)
                rhs197_ = (D7_DC14((d_980_v85_).cf41, p0, len(((d_982_v87_)[p0] if (p0) in (d_982_v87_) else d_981_v86_)))).cf20
                lhs160_ = d_969_v75_
                lhs161_ = default__.safeIndex(903, (d_969_v75_).length(0))
                lhs162_ = globalState
                lhs163_ = globalState
                lhs164_ = globalState
                d_968_v74_ = rhs193_
                lhs160_[lhs161_] = rhs194_
                lhs162_.f16 = rhs195_
                lhs163_.f18 = rhs196_
                lhs164_.f18 = rhs197_
                d_983_v88_: str
                d_983_v88_ = _dafny.CodePoint('k')
                d_983_v88_ = default__.fm10(d_877_v0_, (False) or (d_877_v0_), d_877_v0_, d_877_v0_, globalState)
                d_984_v89_: _dafny.MultiSet
                d_984_v89_ = _dafny.MultiSet([default__.fm29(p0, globalState)])
                d_967_v73_ = default__.fm11(d_983_v88_, (self).fm1(d_984_v89_, globalState), globalState)
                index162_ = default__.safeIndex(901, (d_965_v71_).length(0))
                (d_965_v71_)[index162_] = not(d_877_v0_)
                index163_ = default__.safeIndex(901, (d_965_v71_).length(0))
                (d_965_v71_)[index163_] = not(d_877_v0_)
            d_985_v90_: _dafny.Seq
            d_985_v90_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sc"))
            d_986_v91_: _dafny.Seq
            d_986_v91_ = _dafny.SeqWithoutIsStrInference([len(d_985_v90_)])
            (globalState).f11 = default__.safeDivisionInt(len(d_986_v91_), p0)
            d_987_v92_: D11
            d_987_v92_ = D11_DC24(d_877_v0_, p0)
            d_988_v93_: _dafny.Seq
            d_988_v93_ = _dafny.SeqWithoutIsStrInference([(d_987_v92_).cf36])
            if (d_988_v93_) < (d_988_v93_):
                (globalState).f11 = p0
                d_989_v94_: C2
                nw191_ = C2()
                nw191_.ctor__(p0, d_877_v0_)
                d_989_v94_ = nw191_
                d_990_v95_: _dafny.Map
                d_990_v95_ = _dafny.Map({d_877_v0_: d_989_v94_})
                d_991_v96_: _dafny.Map
                d_991_v96_ = _dafny.Map({d_877_v0_: d_877_v0_})
                d_992_v97_: _dafny.Set
                d_992_v97_ = _dafny.Set({d_985_v90_})
                d_993_v98_: _dafny.MultiSet
                d_993_v98_ = _dafny.MultiSet([d_877_v0_])
                d_994_v99_: _dafny.Seq
                d_994_v99_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tkymklej"))
                d_995_v100_: _dafny.MultiSet
                d_995_v100_ = _dafny.MultiSet([d_994_v99_, d_985_v90_, d_985_v90_, d_994_v99_, d_994_v99_])
                d_996_v101_: _dafny.Array
                nw192_ = _dafny.Array(None, 20)
                nw192_[int(0)] = (len(d_990_v95_)) != ((d_989_v94_).f27)
                nw192_[int(1)] = (d_989_v94_).f28
                nw192_[int(2)] = ((d_991_v96_)[True] if (True) in (d_991_v96_) else (d_989_v94_).f28)
                nw192_[int(3)] = (d_989_v94_).f28
                nw192_[int(4)] = (d_877_v0_) or ((d_989_v94_).f28)
                nw192_[int(5)] = (d_989_v94_).f28
                nw192_[int(6)] = d_877_v0_
                nw192_[int(7)] = ((d_989_v94_).f28) not in (d_988_v93_)
                nw192_[int(8)] = (d_989_v94_).f28
                nw192_[int(9)] = (d_989_v94_).f28
                nw192_[int(10)] = d_877_v0_
                nw192_[int(11)] = (d_992_v97_).isdisjoint(d_992_v97_)
                nw192_[int(12)] = (d_989_v94_).f28
                nw192_[int(13)] = default__.fm0(p0, (d_989_v94_).f28, d_877_v0_, p0, globalState)
                nw192_[int(14)] = (d_993_v98_).isdisjoint(d_993_v98_)
                nw192_[int(15)] = d_877_v0_
                nw192_[int(16)] = not((self).fm1(d_995_v100_, globalState))
                nw192_[int(17)] = d_877_v0_
                nw192_[int(18)] = (p0) < (348)
                nw192_[int(19)] = not(((d_989_v94_).f28) and (False))
                d_996_v101_ = nw192_
                index164_ = default__.safeIndex(517, (d_996_v101_).length(0))
                (d_996_v101_)[index164_] = d_877_v0_
                index165_ = default__.safeIndex(517, (d_996_v101_).length(0))
                (d_996_v101_)[index165_] = not (d_877_v0_) or (False)
                d_997_v102_: D9
                d_997_v102_ = D9_DC17(338)
                d_998_v103_: _dafny.Array
                def lambda53_(d_999_v94_, d_1000_p0_, d_1001_v91_):
                    def lambda54_(d_1002_i6_):
                        return (_dafny.SeqWithoutIsStrInference([(d_999_v94_).f27, d_1000_p0_])) + (d_1001_v91_)

                    return lambda54_

                init31_ = lambda53_(d_989_v94_, p0, d_986_v91_)
                nw193_ = _dafny.Array(None, 23)
                for i0_31_ in range(nw193_.length(0)):
                    nw193_[i0_31_] = init31_(i0_31_)
                d_998_v103_ = nw193_
                d_1003_v104_: _dafny.Seq
                d_1003_v104_ = _dafny.SeqWithoutIsStrInference([(d_989_v94_).f27 for d_1004_i7_ in range(default__.abs(820))])
                index166_ = default__.safeIndex(24, (d_998_v103_).length(0))
                (d_998_v103_)[index166_] = (d_1003_v104_)
                d_1005_v105_: D13
                d_1005_v105_ = D13_DC29((d_996_v101_)[default__.safeIndex(517, (d_996_v101_).length(0))], d_877_v0_, d_996_v101_)
                d_1006_v106_: _dafny.Map
                d_1006_v106_ = _dafny.Map({(d_1005_v105_).cf45: (d_989_v94_).f28})
                d_1007_v107_: _dafny.Seq
                d_1007_v107_ = _dafny.SeqWithoutIsStrInference([d_993_v98_])
                index167_ = default__.safeIndex(24, (d_998_v103_).length(0))
                rhs198_ = (d_997_v102_ if ((d_1006_v106_)[d_996_v101_] if (d_996_v101_) in (d_1006_v106_) else (d_989_v94_).f28) else D9_DC17(len(_dafny.Map({(d_996_v101_)[default__.safeIndex(517, (d_996_v101_).length(0))]: (d_989_v94_).f28}))))
                rhs199_ = _dafny.SeqWithoutIsStrInference([p0])
                rhs200_ = ((len(d_1007_v107_) if (d_989_v94_).f28 else (d_989_v94_).f27)) - (p0)
                rhs201_ = d_996_v101_
                lhs165_ = d_998_v103_
                lhs166_ = default__.safeIndex(24, (d_998_v103_).length(0))
                lhs167_ = globalState
                d_997_v102_ = rhs198_
                lhs165_[lhs166_] = rhs199_
                lhs167_.f16 = rhs200_
                d_996_v101_ = rhs201_
                (globalState).f18 = (p0) > (p0)
                (globalState).f0 = (0) - ((d_989_v94_).f27)
            elif True:
                d_1008_v108_: _dafny.Array
                nw194_ = _dafny.Array(None, 21)
                nw194_[int(0)] = (d_986_v91_) + (d_986_v91_)
                nw194_[int(1)] = (d_986_v91_) + (d_986_v91_)
                nw194_[int(2)] = _dafny.SeqWithoutIsStrInference([(0) - (p0)])
                nw194_[int(3)] = (d_986_v91_) + (d_986_v91_)
                nw194_[int(4)] = (_dafny.SeqWithoutIsStrInference([p0])).set(default__.safeIndex(default__.fm5((d_941_v57_).cardinality, p0, globalState), len(_dafny.SeqWithoutIsStrInference([p0]))), len(d_986_v91_))
                nw194_[int(5)] = (d_986_v91_) + (d_986_v91_)
                nw194_[int(6)] = d_986_v91_
                nw194_[int(7)] = d_986_v91_
                nw194_[int(8)] = d_986_v91_
                nw194_[int(9)] = d_986_v91_
                nw194_[int(10)] = (d_986_v91_) + (_dafny.SeqWithoutIsStrInference([p0 for d_1009_i8_ in range(default__.abs(229))]))
                nw194_[int(11)] = d_986_v91_
                nw194_[int(12)] = d_986_v91_
                nw194_[int(13)] = ((d_986_v91_) + (d_986_v91_)).set(default__.safeIndex(p0, len((d_986_v91_) + (d_986_v91_))), p0)
                nw194_[int(14)] = d_986_v91_
                nw194_[int(15)] = _dafny.SeqWithoutIsStrInference([p0])
                nw194_[int(16)] = d_986_v91_
                nw194_[int(17)] = d_986_v91_
                nw194_[int(18)] = d_986_v91_
                nw194_[int(19)] = _dafny.SeqWithoutIsStrInference([(d_986_v91_)[default__.safeIndex(p0, len(d_986_v91_))] for d_1010_i9_ in range(default__.abs(370))])
                nw194_[int(20)] = d_986_v91_
                d_1008_v108_ = nw194_
                index168_ = default__.safeIndex(306, (d_1008_v108_).length(0))
                (d_1008_v108_)[index168_] = _dafny.SeqWithoutIsStrInference([p0 for d_1011_i10_ in range(default__.abs(483))])
                index169_ = default__.safeIndex(306, (d_1008_v108_).length(0))
                (d_1008_v108_)[index169_] = d_986_v91_
                d_877_v0_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rbg"))) <= ((d_985_v90_ if d_877_v0_ else d_985_v90_))
                d_1012_v109_: C5
                nw195_ = C5()
                nw195_.ctor__(d_877_v0_, (d_985_v90_)[default__.safeIndex(p0, len(d_985_v90_))], p0)
                d_1012_v109_ = nw195_
                d_1013_v110_: C0
                nw196_ = C0()
                nw196_.ctor__()
                d_1013_v110_ = nw196_
                d_1014_v111_: _dafny.Map
                d_1014_v111_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ocqonr")): (d_1012_v109_).f26})
                d_1014_v111_ = (d_1014_v111_) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_1015_i11_ in range(default__.abs(928))]): not(False)}))
            d_1016_v112_: _dafny.Seq
            d_1016_v112_ = d_985_v90_
            d_1017_v113_: _dafny.Map
            d_1017_v113_ = _dafny.Map({d_877_v0_: d_1016_v112_})
            d_1018_v114_: _dafny.Map
            d_1018_v114_ = _dafny.Map({d_1017_v113_: d_877_v0_})
            (globalState).f0 = len((d_1018_v114_) | ((d_1018_v114_) | (default__.fm30(746, p0, d_877_v0_, d_877_v0_, globalState))))
            (globalState).f2 = p0
        d_1019_v115_: str
        d_1019_v115_ = _dafny.CodePoint('p')
        d_1020_v116_: D10
        d_1020_v116_ = D10_DC19(d_1019_v115_)
        d_1021_v117_: _dafny.Seq
        d_1021_v117_ = _dafny.SeqWithoutIsStrInference([d_1020_v116_])
        d_1022_v118_: _dafny.Seq
        d_1022_v118_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "okgjhwo"))
        d_1023_v119_: _dafny.Map
        d_1023_v119_ = _dafny.Map({d_1019_v115_: d_1022_v118_})
        d_1024_v120_: _dafny.Map
        d_1024_v120_ = _dafny.Map({(d_1021_v117_) < (d_1021_v117_): (d_1023_v119_) | (d_1023_v119_)})
        d_1024_v120_ = (d_1024_v120_).set(d_877_v0_, d_1023_v119_)
        d_1025_v121_: _dafny.Set
        d_1025_v121_ = _dafny.Set({default__.fm17(not(False), True, False, globalState)})
        d_1026_v122_: _dafny.Map
        d_1026_v122_ = _dafny.Map({d_877_v0_: d_877_v0_})
        d_1027_v123_: _dafny.MultiSet
        d_1027_v123_ = _dafny.MultiSet([default__.fm5(len(d_1026_v122_), p0, globalState)])
        def iife70_():
            coll42_ = _dafny.Set()
            compr_42_: int
            for compr_42_ in (d_1027_v123_).Elements:
                d_1028_v124_: int = compr_42_
                if (d_1028_v124_) in (d_1027_v123_):
                    def iife71_():
                        coll43_ = _dafny.Set()
                        compr_43_: int
                        for compr_43_ in (_dafny.Set({len(_dafny.Set({(0) - (-941), len(_dafny.SeqWithoutIsStrInference([not(False)])), 774, 187}))})).Elements:
                            d_1029_v125_: int = compr_43_
                            if (d_1029_v125_) in (_dafny.Set({len(_dafny.Set({(0) - (-941), len(_dafny.SeqWithoutIsStrInference([not(False)])), 774, 187}))})):
                                coll43_ = coll43_.union(_dafny.Set([(d_1029_v125_) * (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xn"))), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_1030_i12_ in range(default__.abs(780))]))])))]))
                        return _dafny.Set(coll43_)
                    coll42_ = coll42_.union(_dafny.Set([default__.safeDivisionInt(d_1028_v124_, (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e")))) + (len(_dafny.Map({D3_DC5(_dafny.SeqWithoutIsStrInference([False]), 492, len(iife71_()
)): not(not(False))}))))]))
            return _dafny.Set(coll42_)
        d_1025_v121_ = iife70_()
        
        d_1031_v126_: _dafny.Seq
        d_1031_v126_ = _dafny.SeqWithoutIsStrInference([d_877_v0_, not (d_877_v0_) or (d_877_v0_), d_877_v0_, d_877_v0_])
        if (d_1031_v126_)[default__.safeIndex((p0) + (p0), len(d_1031_v126_))]:
            (globalState).f0 = 330
            d_1032_v127_: C0
            nw197_ = C0()
            nw197_.ctor__()
            d_1032_v127_ = nw197_
            r2 = (d_1022_v118_) + (d_1022_v118_)
            d_1027_v123_ = d_1027_v123_
            d_1033_v128_: _dafny.Map
            d_1033_v128_ = _dafny.Map({d_877_v0_: p0})
            (globalState).f15 = (((d_1033_v128_)[d_877_v0_] if (d_877_v0_) in (d_1033_v128_) else len((d_1022_v118_).set(default__.safeIndex(p0, len(d_1022_v118_)), d_1019_v115_)))) + (p0)
        elif True:
            d_1034_v129_: C3
            nw198_ = C3()
            nw198_.ctor__(True, _dafny.CodePoint('r'), p0)
            d_1034_v129_ = nw198_
            d_1035_v130_: _dafny.Map
            d_1035_v130_ = _dafny.Map({d_1034_v129_: (d_1034_v129_).fm18(globalState)})
            (globalState).f18 = (len((_dafny.Map({d_1034_v129_: True})) | (d_1035_v130_))) != (p0)
            d_1036_v131_: _dafny.Seq
            d_1036_v131_ = _dafny.SeqWithoutIsStrInference([p0, p0])
            d_1037_v132_: C1
            nw199_ = C1()
            nw199_.ctor__(d_1019_v115_, len(d_1036_v131_))
            d_1037_v132_ = nw199_
            d_1038_v133_: _dafny.Map
            d_1038_v133_ = _dafny.Map({True: d_1037_v132_})
            d_1039_v134_: _dafny.Set
            d_1039_v134_ = _dafny.Set({((d_1038_v133_)[d_1034_v129_.f31] if (d_1034_v129_.f31) in (d_1038_v133_) else d_1037_v132_), d_1037_v132_})
            d_1040_v135_: D7
            d_1040_v135_ = D7_DC13((_dafny.Map({len(d_1039_v134_): d_877_v0_})).set((0) - (p0), d_1034_v129_.f31))
            pat_let_tv10_ = d_1023_v119_
            pat_let_tv11_ = d_877_v0_
            def iife72_(_pat_let14_0):
                def iife73_(d_1041_dt__update__tmp_h1_):
                    def iife74_(_pat_let15_0):
                        def iife75_(d_1042_dt__update_hcf19_h0_):
                            return D7_DC13(d_1042_dt__update_hcf19_h0_)
                        return iife75_(_pat_let15_0)
                    return iife74_(_dafny.Map({len(pat_let_tv10_): pat_let_tv11_}))
                return iife73_(_pat_let14_0)
            rhs202_ = iife72_(d_1040_v135_)
            rhs203_ = d_1034_v129_.f31
            lhs168_ = d_1034_v129_
            d_1040_v135_ = rhs202_
            lhs168_.f31 = rhs203_
            d_877_v0_ = not (d_1034_v129_.f31) or ((d_1034_v129_.f31 if d_1034_v129_.f31 else d_1034_v129_.f31))
            (globalState).f18 = d_1034_v129_.f31
            d_1043_v136_: C1
            nw200_ = C1()
            nw200_.ctor__(d_1019_v115_, p0)
            d_1043_v136_ = nw200_
        d_1044_v137_: D9
        d_1044_v137_ = D9_DC16(_dafny.Map({d_877_v0_: D3_DC6(p0, p0, not(d_877_v0_))}))
        pat_let_tv12_ = d_1025_v121_
        pat_let_tv13_ = d_877_v0_
        def lambda55_(source15_):
            if source15_.is_DC17:
                d_1045___mcc_h5_ = source15_.cf25
                d_1046_cf25_ = d_1045___mcc_h5_
                return (len(pat_let_tv12_)) > (619)
            elif True:
                d_1047___mcc_h6_ = source15_.cf24
                d_1048_cf24_ = d_1047___mcc_h6_
                return pat_let_tv13_

        if lambda55_(d_1044_v137_):
            d_1049_v138_: _dafny.Seq
            d_1049_v138_ = d_1022_v118_
            source16_ = d_1049_v138_
            d_1050___mcc_h4_ = source16_
            d_1051_cf1_ = d_1050___mcc_h4_
            d_1052_v139_: _dafny.Array
            nw201_ = _dafny.Array(_dafny.CodePoint('D'), 14)
            d_1052_v139_ = nw201_
            d_1053_v140_: _dafny.Seq
            d_1053_v140_ = _dafny.SeqWithoutIsStrInference([default__.fm5(p0, p0, globalState)])
            d_1054_v141_: _dafny.Map
            d_1054_v141_ = _dafny.Map({d_1052_v139_: (len(d_1053_v140_) if True else p0)})
            r1 = len(d_1054_v141_)
            d_1027_v123_ = (d_1027_v123_) - (d_1027_v123_)
            (self).f23 = (self.f23).set(len(d_1051_cf1_), d_877_v0_)
            r2 = d_1022_v118_
            r1 = (p0) + (default__.safeModuloInt(p0, (0) - (-498)))
            (globalState).f18 = not(not(True))
            d_1055_v142_: _dafny.Array
            nw202_ = _dafny.Array(None, 4)
            nw202_[int(0)] = d_877_v0_
            nw202_[int(1)] = True
            nw202_[int(2)] = d_877_v0_
            nw202_[int(3)] = d_877_v0_
            d_1055_v142_ = nw202_
            d_1056_v143_: _dafny.Seq
            d_1056_v143_ = _dafny.SeqWithoutIsStrInference([d_1055_v142_])
            d_1057_v144_: _dafny.Map
            d_1057_v144_ = _dafny.Map({len(d_1056_v143_): d_1019_v115_})
            d_1057_v144_ = (d_1057_v144_).set(len(d_1022_v118_), default__.fm10(d_877_v0_, d_877_v0_, False, d_877_v0_, globalState))
            d_1058_v145_: _dafny.Array
            nw203_ = _dafny.Array(int(0), 9)
            d_1058_v145_ = nw203_
            d_1059_v146_: _dafny.Map
            d_1059_v146_ = _dafny.Map({d_877_v0_: d_1058_v145_})
            if ((d_1059_v146_) | (d_1059_v146_)) != ((d_1059_v146_) | ((d_1059_v146_).set(False, d_1058_v145_))):
                d_1060_v147_: _dafny.MultiSet
                d_1060_v147_ = _dafny.MultiSet([d_877_v0_, d_877_v0_])
                d_1060_v147_ = (d_1060_v147_) - (d_1060_v147_)
                (globalState).f18 = d_877_v0_
                d_1061_v148_: _dafny.Array
                def lambda56_(d_1062_v118_):
                    def lambda57_(d_1063_i13_):
                        return d_1062_v118_

                    return lambda57_

                init32_ = lambda56_(d_1022_v118_)
                nw204_ = _dafny.Array(None, 16)
                for i0_32_ in range(nw204_.length(0)):
                    nw204_[i0_32_] = init32_(i0_32_)
                d_1061_v148_ = nw204_
                index170_ = default__.safeIndex(110, (d_1061_v148_).length(0))
                (d_1061_v148_)[index170_] = d_1022_v118_
                index171_ = default__.safeIndex(110, (d_1061_v148_).length(0))
                (d_1061_v148_)[index171_] = d_1022_v118_
                d_1064_v149_: _dafny.Array
                nw205_ = _dafny.Array(_dafny.Set({}), 8)
                d_1064_v149_ = nw205_
                d_1065_v150_: _dafny.Seq
                d_1065_v150_ = _dafny.SeqWithoutIsStrInference([d_1025_v121_, _dafny.Set({(0) - (p0)})])
                d_1066_v151_: _dafny.Seq
                d_1066_v151_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                d_1067_v152_: _dafny.Seq
                d_1067_v152_ = _dafny.SeqWithoutIsStrInference([d_1066_v151_])
                d_1068_v153_: _dafny.Seq
                d_1068_v153_ = _dafny.SeqWithoutIsStrInference([(d_1067_v152_)[default__.safeIndex((0) - (len(d_1026_v122_)), len(d_1067_v152_))], _dafny.SeqWithoutIsStrInference([p0])])
                d_1069_v154_: _dafny.Seq
                d_1069_v154_ = _dafny.SeqWithoutIsStrInference([d_1025_v121_, d_1025_v121_, (d_1065_v150_)[default__.safeIndex(433, len(d_1065_v150_))], _dafny.Set({len(d_1067_v152_), p0}), _dafny.Set({p0, p0})])
                d_1070_v155_: _dafny.Map
                d_1070_v155_ = _dafny.Map({(_dafny.MultiSet([d_877_v0_])).cardinality: p0})
                index172_ = default__.safeIndex(251, (d_1064_v149_).length(0))
                (d_1064_v149_)[index172_] = (d_1065_v150_)[default__.safeIndex(len(d_1070_v155_), len(d_1065_v150_))]
                index173_ = default__.safeIndex(802, (d_1055_v142_).length(0))
                (d_1055_v142_)[index173_] = d_877_v0_
                d_1071_v156_: _dafny.Seq
                d_1071_v156_ = _dafny.SeqWithoutIsStrInference([self.f23, self.f23, self.f23])
                index174_ = default__.safeIndex(251, (d_1064_v149_).length(0))
                index175_ = default__.safeIndex(802, (d_1055_v142_).length(0))
                rhs204_ = (d_1025_v121_).intersection((d_1025_v121_) | (_dafny.Set({p0, p0})))
                rhs205_ = 374
                rhs206_ = ((self.f23) | (self.f23)) in (d_1071_v156_)
                lhs169_ = d_1064_v149_
                lhs170_ = default__.safeIndex(251, (d_1064_v149_).length(0))
                lhs171_ = globalState
                lhs172_ = d_1055_v142_
                lhs173_ = default__.safeIndex(802, (d_1055_v142_).length(0))
                lhs169_[lhs170_] = rhs204_
                lhs171_.f10 = rhs205_
                lhs172_[lhs173_] = rhs206_
                (globalState).f10 = ((d_1060_v147_ if (d_1055_v142_)[default__.safeIndex(802, (d_1055_v142_).length(0))] else d_1060_v147_)).cardinality
            elif True:
                index176_ = default__.safeIndex(262, (d_1055_v142_).length(0))
                (d_1055_v142_)[index176_] = d_877_v0_
                index177_ = default__.safeIndex(262, (d_1055_v142_).length(0))
                (d_1055_v142_)[index177_] = not(d_877_v0_)
                d_1072_v157_: _dafny.Map
                d_1072_v157_ = _dafny.Map({_dafny.CodePoint('g'): p0})
                d_1073_v158_: _dafny.Seq
                d_1073_v158_ = _dafny.SeqWithoutIsStrInference([d_1072_v157_])
                d_1074_v159_: _dafny.Seq
                d_1074_v159_ = _dafny.SeqWithoutIsStrInference([794, len((d_1072_v157_) | ((d_1073_v158_)[default__.safeIndex(p0, len(d_1073_v158_))])), p0, p0])
                (globalState).f15 = (d_1074_v159_)[default__.safeIndex(p0, len(d_1074_v159_))]
                d_1075_v160_: C0
                nw206_ = C0()
                nw206_.ctor__()
                d_1075_v160_ = nw206_
                index178_ = default__.safeIndex(262, (d_1055_v142_).length(0))
                rhs207_ = d_877_v0_
                rhs208_ = (d_1074_v159_) + (d_1074_v159_)
                rhs209_ = d_1075_v160_
                rhs210_ = (d_1055_v142_)[default__.safeIndex(262, (d_1055_v142_).length(0))]
                rhs211_ = (d_1027_v123_).isdisjoint((d_1027_v123_).set(p0, default__.abs((0) - (len(default__.fm31(len(d_1074_v159_), (d_1055_v142_)[default__.safeIndex(262, (d_1055_v142_).length(0))], d_877_v0_, len(d_1022_v118_), globalState))))))
                lhs174_ = globalState
                lhs175_ = d_1055_v142_
                lhs176_ = default__.safeIndex(262, (d_1055_v142_).length(0))
                lhs177_ = globalState
                d_877_v0_ = rhs207_
                lhs174_.f19 = rhs208_
                d_1075_v160_ = rhs209_
                lhs175_[lhs176_] = rhs210_
                lhs177_.f18 = rhs211_
                d_1076_v161_: _dafny.Array
                nw207_ = _dafny.Array(None, 6)
                nw207_[int(0)] = (d_1055_v142_)[default__.safeIndex(262, (d_1055_v142_).length(0))]
                nw207_[int(1)] = d_877_v0_
                nw207_[int(2)] = (d_1055_v142_)[default__.safeIndex(262, (d_1055_v142_).length(0))]
                nw207_[int(3)] = d_877_v0_
                nw207_[int(4)] = False
                nw207_[int(5)] = False
                d_1076_v161_ = nw207_
                d_1077_v162_: _dafny.Map
                d_1077_v162_ = _dafny.Map({d_1076_v161_: (d_1027_v123_).cardinality})
                d_1077_v162_ = (d_1077_v162_).set(d_1076_v161_, (p0) - ((0) - (p0)))
                d_1078_v163_: _dafny.Array
                nw208_ = _dafny.Array(_dafny.Map({}), 6)
                d_1078_v163_ = nw208_
                d_1079_v164_: _dafny.Map
                d_1079_v164_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oo")): (d_1055_v142_)[default__.safeIndex(262, (d_1055_v142_).length(0))]})
                index179_ = default__.safeIndex(958, (d_1078_v163_).length(0))
                (d_1078_v163_)[index179_] = d_1079_v164_
                index180_ = default__.safeIndex(958, (d_1078_v163_).length(0))
                (d_1078_v163_)[index180_] = ((d_1079_v164_).set(d_1022_v118_, (d_1055_v142_)[default__.safeIndex(262, (d_1055_v142_).length(0))])) | (d_1079_v164_)
        elif True:
            d_1080_v165_: _dafny.Array
            nw209_ = _dafny.Array(False, 29)
            d_1080_v165_ = nw209_
            index181_ = default__.safeIndex(172, (d_1080_v165_).length(0))
            (d_1080_v165_)[index181_] = d_877_v0_
            index182_ = default__.safeIndex(172, (d_1080_v165_).length(0))
            (d_1080_v165_)[index182_] = not(d_877_v0_)
            (globalState).f16 = default__.safeModuloInt(len(d_1026_v122_), len(_dafny.SeqWithoutIsStrInference([(d_1080_v165_)[default__.safeIndex(172, (d_1080_v165_).length(0))]])))
            (globalState).f18 = not(default__.fm0(p0, (p0) < (p0), False, p0, globalState))
            d_1022_v118_ = (d_1022_v118_ if (d_1080_v165_)[default__.safeIndex(172, (d_1080_v165_).length(0))] else (d_1022_v118_) + (d_1022_v118_))
            d_1081_v166_: C3
            nw210_ = C3()
            nw210_.ctor__(d_877_v0_, d_1019_v115_, p0)
            d_1081_v166_ = nw210_
            d_1081_v166_ = (d_1081_v166_ if (d_1080_v165_)[default__.safeIndex(172, (d_1080_v165_).length(0))] else d_1081_v166_)
        r0 = p0
        r1 = default__.fm17(d_877_v0_, d_877_v0_, (d_1031_v126_)[default__.safeIndex(len((d_1022_v118_).set(default__.safeIndex(611, len(d_1022_v118_)), default__.fm10(d_877_v0_, d_877_v0_, d_877_v0_, d_877_v0_, globalState))), len(d_1031_v126_))], globalState)
        r2 = d_1022_v118_
        return r0, r1, r2

