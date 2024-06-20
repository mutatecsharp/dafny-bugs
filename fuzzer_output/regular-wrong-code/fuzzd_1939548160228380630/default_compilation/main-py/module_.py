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
            for compr_0_ in _dafny.IntegerRange(985, 907):
                d_0_v0_: int = compr_0_
                if ((985) <= (d_0_v0_)) and ((d_0_v0_) < (907)):
                    coll0_[(d_0_v0_) - (-431)] = _dafny.CodePoint('k')
            return _dafny.Map(coll0_)
        return (len(iife0_()
        )) - ((0) - ((-875) + (32)))

    @staticmethod
    def fm3(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([(0) - (-909), len(_dafny.SeqWithoutIsStrInference([False, True])), -14, 948, (0) - ((0) - ((_dafny.MultiSet([False])).cardinality))])) <= (_dafny.SeqWithoutIsStrInference([502, 581, -87, 242]))) and ((len(_dafny.Map({True: len(_dafny.Map({not(False): False}))}))) == (-105))

    @staticmethod
    def fm5(globalState):
        return _dafny.CodePoint('t')

    @staticmethod
    def fm13(p0, globalState):
        source0_ = D7_DC16(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_1_i0_ in range(default__.abs(453))])))
        if source0_.is_DC16:
            d_2___mcc_h0_ = source0_.cf29
            d_3_cf29_ = d_2___mcc_h0_
            if True:
                return _dafny.CodePoint('w')
            elif True:
                return _dafny.CodePoint('r')
        elif source0_.is_DC15:
            d_4___mcc_h1_ = source0_.cf28
            d_5_cf28_ = d_4___mcc_h1_
            return _dafny.CodePoint('i')
        elif True:
            d_6___mcc_h2_ = source0_.cf30
            d_7_cf30_ = d_6___mcc_h2_
            return _dafny.CodePoint('q')

    @staticmethod
    def fm17(globalState):
        return True

    @staticmethod
    def fm18(p0, p1, p2, p3, globalState):
        if not(not((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "og"))])).ispropersubset(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sple"))])))):
            return _dafny.CodePoint('l')
        elif True:
            return _dafny.CodePoint('t')

    @staticmethod
    def fm19(globalState):
        return ((_dafny.MultiSet([(0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_8_i0_ in range(default__.abs(8))])))), (0) - (-377)])).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([218, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aqrya")))]))])))) - ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c")))])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({-649, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mbapixuot"))), 394})), 923, 695]))))

    @staticmethod
    def fm20(globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ofvstaor"))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rfjmdefgm"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))))

    @staticmethod
    def fm21(globalState):
        return D1_DC1((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sflqu"))))

    @staticmethod
    def fm22(p0, p1, globalState):
        source1_ = (D6_DC11(_dafny.MultiSet([344])) if False else D6_DC11(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([-70]))))
        if source1_.is_DC12:
            d_9___mcc_h0_ = source1_.cf18
            d_10___mcc_h1_ = source1_.cf19
            d_11___mcc_h2_ = source1_.cf20
            d_12___mcc_h3_ = source1_.cf21
            d_13_cf21_ = d_12___mcc_h3_
            d_14_cf20_ = d_11___mcc_h2_
            d_15_cf19_ = d_10___mcc_h1_
            d_16_cf18_ = d_9___mcc_h0_
            return _dafny.Set({d_15_cf19_})
        elif source1_.is_DC13:
            d_17___mcc_h4_ = source1_.cf22
            d_18___mcc_h5_ = source1_.cf23
            d_19___mcc_h6_ = source1_.cf24
            d_20___mcc_h7_ = source1_.cf25
            d_21___mcc_h8_ = source1_.cf26
            d_22_cf26_ = d_21___mcc_h8_
            d_23_cf25_ = d_20___mcc_h7_
            d_24_cf24_ = d_19___mcc_h6_
            d_25_cf23_ = d_18___mcc_h5_
            d_26_cf22_ = d_17___mcc_h4_
            return _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yyjhaqfxv"))})
        elif source1_.is_DC11:
            d_27___mcc_h9_ = source1_.cf17
            d_28_cf17_ = d_27___mcc_h9_
            return _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ycub")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dvtlsfkq"))})
        elif True:
            d_29___mcc_h10_ = source1_.cf27
            d_30_cf27_ = d_29___mcc_h10_
            def iife1_():
                coll1_ = _dafny.Set()
                compr_1_: _dafny.Seq
                for compr_1_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vecee")): True})).keys.Elements:
                    d_31_v0_: _dafny.Seq = compr_1_
                    if (d_31_v0_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vecee")): True})):
                        coll1_ = coll1_.union(_dafny.Set([d_31_v0_]))
                return _dafny.Set(coll1_)
            return iife1_()
            

    @staticmethod
    def fm23(p0, globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([((D8_DC19(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([True]))}), _dafny.Set({True, False}), _dafny.MultiSet([296]), D1_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lqbbxmxcd"))))).cf34).cardinality]))).Elements:
                d_32_v0_: int = compr_2_
                if (d_32_v0_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([((D8_DC19(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([True]))}), _dafny.Set({True, False}), _dafny.MultiSet([296]), D1_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lqbbxmxcd"))))).cf34).cardinality]))):
                    coll2_[default__.safeModuloInt(d_32_v0_, (_dafny.MultiSet([576, 33])).cardinality)] = _dafny.CodePoint('f')
            return _dafny.Map(coll2_)
        return ((iife2_()
        ) | (_dafny.Map({-86: _dafny.CodePoint('f')}))) | ((_dafny.Map({137: _dafny.CodePoint('n')}) if False else _dafny.Map({len(_dafny.Map({True: -424})): _dafny.CodePoint('w')})))

    @staticmethod
    def fm24(p0, globalState):
        return ((_dafny.Map({137: len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gfpmde"))): True}))})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_33_i0_ in range(default__.abs(508))])): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qtnakf")))}))) | (_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bug")))): 798}))

    @staticmethod
    def fm25(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "puihdeaw"))): not(not(False))}) for d_34_i0_ in range(default__.abs(-151))])

    @staticmethod
    def fm26(p0, p1, p2, globalState):
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: _dafny.Seq
            for compr_3_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_35_i0_ in range(default__.abs(757))])])).Elements:
                d_36_v0_: _dafny.Seq = compr_3_
                if (d_36_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_35_i0_ in range(default__.abs(757))])])):
                    coll3_[d_36_v0_] = not(True)
            return _dafny.Map(coll3_)
        def iife4_():
            coll4_ = _dafny.Map()
            compr_4_: _dafny.Seq
            for compr_4_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tio")): _dafny.Set({True})})).keys.Elements:
                d_37_v1_: _dafny.Seq = compr_4_
                if (d_37_v1_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tio")): _dafny.Set({True})})):
                    coll4_[d_37_v1_] = False
            return _dafny.Map(coll4_)
        return ((iife3_()
        ) | (iife4_()
        )) | ((_dafny.Map({(D9_DC22(True, len(_dafny.Map({(_dafny.MultiSet([False, True])).cardinality: True})), not(True), _dafny.Set({not(True)}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gc")))).cf41: True})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vgp")): not(False)})))

    @staticmethod
    def fm27(globalState):
        return (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_38_i0_ in range(default__.abs(538))])), -901])) + (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({len(_dafny.Set({len((D8_DC19(_dafny.Map({True: 104}), _dafny.Set({False}), _dafny.MultiSet([len(_dafny.Set({445})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dixgeiij")))]), D1_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lqypklsa"))))).cf32)})): (0) - (len(_dafny.SeqWithoutIsStrInference([True, True, True])))}))) for d_39_i1_ in range(default__.abs(-458))]))

    @staticmethod
    def fm28(p0, p1, p2, globalState):
        return ((_dafny.MultiSet([False])).intersection(_dafny.MultiSet([False]))).intersection((_dafny.MultiSet([False, False, True])) | (_dafny.MultiSet([(D6_DC12(True, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_40_i0_ in range(default__.abs(186))]), 574, True)).cf18, True, False, False, True])))

    @staticmethod
    def fm29(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([not(False)]))

    @staticmethod
    def fm30(p0, globalState):
        return False

    @staticmethod
    def fm31(globalState):
        def iife5_():
            coll5_ = _dafny.Map()
            compr_5_: _dafny.Seq
            for compr_5_ in ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_41_i0_ in range(default__.abs(-335))])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aavabqni"))]))).Elements:
                d_42_v0_: _dafny.Seq = compr_5_
                if (d_42_v0_) in ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_41_i0_ in range(default__.abs(-335))])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aavabqni"))]))):
                    coll5_[d_42_v0_] = _dafny.Map({False: _dafny.CodePoint('n')})
            return _dafny.Map(coll5_)
        return iife5_()
        

    @staticmethod
    def fm32(p0, globalState):
        source2_ = D16_DC41(True)
        if source2_.is_DC41:
            d_43___mcc_h0_ = source2_.cf84
            d_44_cf84_ = d_43___mcc_h0_
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sqsbuy"))
        elif source2_.is_DC42:
            d_45___mcc_h1_ = source2_.cf85
            d_46_cf85_ = d_45___mcc_h1_
            return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_47_i0_ in range(default__.abs(660))])
        elif source2_.is_DC40:
            d_48___mcc_h2_ = source2_.cf83
            d_49_cf83_ = d_48___mcc_h2_
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sidmalhfq"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nga")))
        elif True:
            d_50___mcc_h3_ = source2_.cf86
            d_51_cf86_ = d_50___mcc_h3_
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bnryd"))

    @staticmethod
    def fm33(p0, p1, globalState):
        return _dafny.CodePoint('o')

    @staticmethod
    def fm34(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pdggpy")))])) + ((_dafny.SeqWithoutIsStrInference([859])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([497 for d_52_i1_ in range(default__.abs(-560))])) for d_53_i0_ in range(default__.abs(-621))])))

    @staticmethod
    def fm35(p0, p1, p2, p3, globalState):
        return _dafny.Map({(True) == (False): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "coubrsv"))})

    @staticmethod
    def fm36(p0, p1, p2, globalState):
        def iife6_():
            coll6_ = _dafny.Map()
            compr_6_: int
            for compr_6_ in _dafny.IntegerRange(789, 442):
                d_54_v1_: int = compr_6_
                if ((789) <= (d_54_v1_)) and ((d_54_v1_) < (442)):
                    coll6_[(d_54_v1_) * (-848)] = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([True, True])).cardinality])), 12]))).cardinality
            return _dafny.Map(coll6_)
        def iife7_():
            coll7_ = _dafny.Set()
            compr_7_: int
            for compr_7_ in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hvfvypqh"))): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "weervqcj"))})).keys.Elements:
                d_55_v0_: int = compr_7_
                if (d_55_v0_) in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hvfvypqh"))): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "weervqcj"))})):
                    coll7_ = coll7_.union(_dafny.Set([default__.safeModuloInt(d_55_v0_, (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lhu")))])).cardinality)]))
            return _dafny.Set(coll7_)
        if (_dafny.Set({len(iife6_()
        )})).issubset(iife7_()
        ):
            return (_dafny.Set({True, not(False), False})) | (_dafny.Set({True}))
        elif True:
            return _dafny.Set({True})

    @staticmethod
    def fm37(p0, p1, p2, p3, globalState):
        return D1_DC2(len(_dafny.Map({D3_DC7(): _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ti"))), 332])})), (-570) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wmlctwlx")))), False)

    @staticmethod
    def fm38(p0, p1, p2, p3, globalState):
        def iife8_():
            coll8_ = _dafny.Set()
            compr_8_: int
            for compr_8_ in _dafny.IntegerRange(201, -244):
                d_56_v0_: int = compr_8_
                if ((201) <= (d_56_v0_)) and ((d_56_v0_) < (-244)):
                    coll8_ = coll8_.union(_dafny.Set([(d_56_v0_) * ((0) - (len(_dafny.SeqWithoutIsStrInference([-467, 716]))))]))
            return _dafny.Set(coll8_)
        def iife9_():
            coll9_ = _dafny.Set()
            compr_9_: int
            for compr_9_ in _dafny.IntegerRange(371, 582):
                d_58_v1_: int = compr_9_
                if ((371) <= (d_58_v1_)) and ((d_58_v1_) < (582)):
                    coll9_ = coll9_.union(_dafny.Set([(d_58_v1_) - (len(_dafny.Map({False: 704})))]))
            return _dafny.Set(coll9_)
        return _dafny.SeqWithoutIsStrInference([(iife8_()
        ).isdisjoint(_dafny.Set({289, len(_dafny.Map({len(_dafny.Set({len(_dafny.Map({-587: _dafny.CodePoint('q')})), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_57_i0_ in range(default__.abs(969))])), len(iife9_()
        )})): False}))}))])

    @staticmethod
    def fm39(p0, p1, globalState):
        return D6_DC13((-93) == (37), False, (0) - (len(_dafny.SeqWithoutIsStrInference([-741 for d_59_i0_ in range(default__.abs(933))]))), 232, not(not (True) or (not(False))))

    @staticmethod
    def fm40(globalState):
        return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, True]))) | (_dafny.MultiSet([True, not(False)]))

    @staticmethod
    def fm41(p0, p1, globalState):
        def iife10_():
            coll10_ = _dafny.Set()
            compr_10_: int
            for compr_10_ in _dafny.IntegerRange(272, 733):
                d_60_v0_: int = compr_10_
                if ((272) <= (d_60_v0_)) and ((d_60_v0_) < (733)):
                    coll10_ = coll10_.union(_dafny.Set([default__.safeDivisionInt(d_60_v0_, 509)]))
            return _dafny.Set(coll10_)
        return ((D21_DC54(_dafny.Map({not(True): _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fdkkibdc")))})}))).cf105) | ((_dafny.Map({False: _dafny.Set({len(_dafny.Map({True: 958})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dg")))})})) | (_dafny.Map({True: _dafny.Set({len(iife10_()
        ), len(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "win"))}))})})))

    @staticmethod
    def fm44(p0, globalState):
        source3_ = (D21_DC56(True) if True else D21_DC56(False))
        if source3_.is_DC55:
            d_61___mcc_h0_ = source3_.cf106
            d_62___mcc_h1_ = source3_.cf107
            d_63___mcc_h2_ = source3_.cf108
            d_64_cf108_ = d_63___mcc_h2_
            d_65_cf107_ = d_62___mcc_h1_
            d_66_cf106_ = d_61___mcc_h0_
            def iife11_():
                def iife13_():
                    coll13_ = _dafny.Set()
                    compr_13_: _dafny.Set
                    for compr_13_ in (_dafny.Map({_dafny.Set({d_66_cf106_}): d_66_cf106_})).keys.Elements:
                        d_69_v1_: _dafny.Set = compr_13_
                        if (d_69_v1_) in (_dafny.Map({_dafny.Set({d_66_cf106_}): d_66_cf106_})):
                            coll13_ = coll13_.union(_dafny.Set([d_69_v1_]))
                    return _dafny.Set(coll13_)
                coll11_ = _dafny.Map()
                def iife12_():
                    coll12_ = _dafny.Set()
                    compr_12_: _dafny.Set
                    for compr_12_ in (_dafny.Map({_dafny.Set({d_66_cf106_}): d_66_cf106_})).keys.Elements:
                        d_67_v1_: _dafny.Set = compr_12_
                        if (d_67_v1_) in (_dafny.Map({_dafny.Set({d_66_cf106_}): d_66_cf106_})):
                            coll12_ = coll12_.union(_dafny.Set([d_67_v1_]))
                    return _dafny.Set(coll12_)
                compr_11_: _dafny.Set
                for compr_11_ in (iife12_()
                ).Elements:
                    d_68_v0_: _dafny.Set = compr_11_
                    if (d_68_v0_) in (iife13_()
                    ):
                        coll11_[d_68_v0_] = (0) - (d_64_cf108_)
                return _dafny.Map(coll11_)
            def iife14_():
                coll14_ = _dafny.Map()
                compr_14_: _dafny.Set
                for compr_14_ in (_dafny.SeqWithoutIsStrInference([_dafny.Set({not(True)})])).Elements:
                    d_70_v2_: _dafny.Set = compr_14_
                    if (d_70_v2_) in (_dafny.SeqWithoutIsStrInference([_dafny.Set({not(True)})])):
                        coll14_[d_70_v2_] = d_64_cf108_
                return _dafny.Map(coll14_)
            return (iife11_()
            ) | ((D22_DC58(iife14_()
)).cf112)
        elif source3_.is_DC56:
            d_71___mcc_h3_ = source3_.cf109
            d_72_cf109_ = d_71___mcc_h3_
            return _dafny.Map({_dafny.Set({d_72_cf109_, d_72_cf109_}): (_dafny.MultiSet([True])).cardinality})
        elif source3_.is_DC57:
            d_73___mcc_h4_ = source3_.cf110
            d_74___mcc_h5_ = source3_.cf111
            d_75_cf111_ = d_74___mcc_h5_
            d_76_cf110_ = d_73___mcc_h4_
            return _dafny.Map({_dafny.Set({d_76_cf110_}): (_dafny.MultiSet([d_76_cf110_])).cardinality})
        elif True:
            d_77___mcc_h6_ = source3_.cf105
            d_78_cf105_ = d_77___mcc_h6_
            return _dafny.Map({_dafny.Set({True}): 345})

    @staticmethod
    def fm45(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_79_i0_ in range(default__.abs(483))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_80_i1_ in range(default__.abs(797))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wenrpms"))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "shmvcrb"))]))

    @staticmethod
    def fm46(p0, p1, p2, globalState):
        return D9_DC22((False if False else True), (0) - (len((_dafny.SeqWithoutIsStrInference([-266])) + (_dafny.SeqWithoutIsStrInference([997, 737])))), (len(_dafny.Map({not(True): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uxklvq")))}))) <= (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_81_i0_ in range(default__.abs(264))]))), _dafny.Set({False, False}), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vot")))

    @staticmethod
    def fm47(p0, globalState):
        def iife15_():
            coll15_ = _dafny.Set()
            compr_15_: int
            for compr_15_ in (_dafny.Set({(_dafny.MultiSet([(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_82_i0_ in range(default__.abs(-279))])))])).cardinality, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_83_i1_ in range(default__.abs(564))]))})).Elements:
                d_84_v0_: int = compr_15_
                if (d_84_v0_) in (_dafny.Set({(_dafny.MultiSet([(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_82_i0_ in range(default__.abs(-279))])))])).cardinality, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_83_i1_ in range(default__.abs(564))]))})):
                    coll15_ = coll15_.union(_dafny.Set([(d_84_v0_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "thfmph"))))]))
            return _dafny.Set(coll15_)
        return (_dafny.Set({167, (0) - (-372)})).intersection((_dafny.Set({-541})) - (iife15_()
        ))

    @staticmethod
    def fm48(globalState):
        def iife16_():
            coll16_ = _dafny.Set()
            compr_16_: _dafny.Seq
            for compr_16_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))]))).Elements:
                d_86_v0_: _dafny.Seq = compr_16_
                if (d_86_v0_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))]))):
                    coll16_ = coll16_.union(_dafny.Set([d_86_v0_]))
            return _dafny.Set(coll16_)
        return ((_dafny.Set({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_85_i0_ in range(default__.abs(727))])})).intersection(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "utp"))}))).intersection((iife16_()
        ) | (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ogjtqc")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hyaujykjq"))})))

    @staticmethod
    def fm49(p0, p1, globalState):
        return D3_DC8(((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "idqyyo")))])) | (_dafny.MultiSet([(_dafny.MultiSet([True])).cardinality]))).cardinality)

    @staticmethod
    def fm50(globalState):
        return _dafny.Map({((_dafny.MultiSet([956])).cardinality) + (415): (_dafny.CodePoint('s')) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v")))})

    @staticmethod
    def fm51(globalState):
        return D9_DC23((416 if True else -787), default__.safeDivisionInt(371, 820), ((_dafny.MultiSet([True, False, not(True)])) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, not(not(True)), True, False])))).cardinality, (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_87_i0_ in range(default__.abs(-494))]))) * (839), 109)

    @staticmethod
    def fm54(p0, globalState):
        return ((_dafny.Set({not(not(False))})).intersection(_dafny.Set({False, True}))) - (_dafny.Set({True}))

    @staticmethod
    def fm55(p0, p1, p2, globalState):
        return _dafny.MultiSet([True])

    @staticmethod
    def fm56(p0, p1, globalState):
        return ((_dafny.MultiSet([len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bdebqr")))})), (0) - (len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hixovsg")))})))])) - (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([-869 for d_88_i0_ in range(default__.abs(651))])), (_dafny.MultiSet([False, False, False])).cardinality]))) - (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([False])), 941]))

    @staticmethod
    def fm57(globalState):
        source4_ = D21_DC56(True)
        if source4_.is_DC55:
            d_89___mcc_h0_ = source4_.cf106
            d_90___mcc_h1_ = source4_.cf107
            d_91___mcc_h2_ = source4_.cf108
            d_92_cf108_ = d_91___mcc_h2_
            d_93_cf107_ = d_90___mcc_h1_
            d_94_cf106_ = d_89___mcc_h0_
            return D12_DC32(d_94_cf106_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))), False, d_94_cf106_, d_92_cf108_)
        elif source4_.is_DC56:
            d_95___mcc_h3_ = source4_.cf109
            d_96_cf109_ = d_95___mcc_h3_
            return D12_DC32(d_96_cf109_, 649, d_96_cf109_, d_96_cf109_, (_dafny.MultiSet([_dafny.Map({d_96_cf109_: d_96_cf109_}), _dafny.Map({d_96_cf109_: d_96_cf109_}), _dafny.Map({d_96_cf109_: d_96_cf109_}), _dafny.Map({d_96_cf109_: d_96_cf109_}), _dafny.Map({d_96_cf109_: d_96_cf109_})])).cardinality)
        elif source4_.is_DC57:
            d_97___mcc_h4_ = source4_.cf110
            d_98___mcc_h5_ = source4_.cf111
            d_99_cf111_ = d_98___mcc_h5_
            d_100_cf110_ = d_97___mcc_h4_
            return D12_DC32(not(d_100_cf110_), len((D22_DC60(-656, _dafny.SeqWithoutIsStrInference([False, not(d_100_cf110_)]), False)).cf116), d_100_cf110_, d_100_cf110_, 559)
        elif True:
            d_101___mcc_h6_ = source4_.cf105
            d_102_cf105_ = d_101___mcc_h6_
            return D12_DC32(False, -50, False, True, len(_dafny.Set({698, 259})))

    @staticmethod
    def fm58(globalState):
        if (len(_dafny.SeqWithoutIsStrInference([325]))) > ((_dafny.MultiSet([len(_dafny.Map({_dafny.SeqWithoutIsStrInference([D22_DC59(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_103_i1_ in range(default__.abs(506))])), False) for d_104_i0_ in range(default__.abs(-639))]): _dafny.SeqWithoutIsStrInference([668])}))])).cardinality):
            return _dafny.CodePoint('i')
        elif not(False):
            return _dafny.CodePoint('f')
        elif True:
            return _dafny.CodePoint('q')

    @staticmethod
    def fm59(p0, p1, p2, p3, globalState):
        return _dafny.Map({_dafny.CodePoint('l'): True})

    @staticmethod
    def fm60(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_105_i0_ in range(default__.abs(182))])

    @staticmethod
    def fm61(p0, p1, p2, globalState):
        source5_ = D6_DC11(_dafny.MultiSet([-259]))
        if source5_.is_DC12:
            d_106___mcc_h0_ = source5_.cf18
            d_107___mcc_h1_ = source5_.cf19
            d_108___mcc_h2_ = source5_.cf20
            d_109___mcc_h3_ = source5_.cf21
            d_110_cf21_ = d_109___mcc_h3_
            d_111_cf20_ = d_108___mcc_h2_
            d_112_cf19_ = d_107___mcc_h1_
            d_113_cf18_ = d_106___mcc_h0_
            return _dafny.SeqWithoutIsStrInference([not(d_110_cf21_), True, d_110_cf21_, False, False])
        elif source5_.is_DC13:
            d_114___mcc_h4_ = source5_.cf22
            d_115___mcc_h5_ = source5_.cf23
            d_116___mcc_h6_ = source5_.cf24
            d_117___mcc_h7_ = source5_.cf25
            d_118___mcc_h8_ = source5_.cf26
            d_119_cf26_ = d_118___mcc_h8_
            d_120_cf25_ = d_117___mcc_h7_
            d_121_cf24_ = d_116___mcc_h6_
            d_122_cf23_ = d_115___mcc_h5_
            d_123_cf22_ = d_114___mcc_h4_
            return (_dafny.SeqWithoutIsStrInference([d_123_cf22_])) + (_dafny.SeqWithoutIsStrInference([d_119_cf26_]))
        elif source5_.is_DC11:
            d_124___mcc_h9_ = source5_.cf17
            d_125_cf17_ = d_124___mcc_h9_
            return (_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([True]))
        elif True:
            d_126___mcc_h10_ = source5_.cf27
            d_127_cf27_ = d_126___mcc_h10_
            return _dafny.SeqWithoutIsStrInference([not(True)])

    @staticmethod
    def fm62(p0, p1, globalState):
        return D1_DC1((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cb"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rjfj"))))

    @staticmethod
    def fm63(p0, p1, globalState):
        if not(False):
            return (_dafny.SeqWithoutIsStrInference([489 for d_128_i0_ in range(default__.abs(791))])) + (_dafny.SeqWithoutIsStrInference([31, (0) - (len(_dafny.Map({449: True})))]))
        elif True:
            return (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gqn"))) for d_129_i1_ in range(default__.abs(995))])) + (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([False])).cardinality]))

    @staticmethod
    def fm64(globalState):
        return (_dafny.Set({787})).intersection((_dafny.Set({350, -91, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w")))})).intersection((D19_DC49(_dafny.Set({(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_130_i0_ in range(default__.abs(816))]))])).cardinality}))).cf93))

    @staticmethod
    def fm65(p0, p1, p2, p3, globalState):
        def iife17_():
            coll17_ = _dafny.Set()
            compr_17_: int
            for compr_17_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({649}))])).Elements:
                d_131_v0_: int = compr_17_
                if (d_131_v0_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({649}))])):
                    coll17_ = coll17_.union(_dafny.Set([(d_131_v0_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xfi"))))]))
            return _dafny.Set(coll17_)
        def iife18_():
            coll18_ = _dafny.Set()
            compr_18_: int
            for compr_18_ in _dafny.IntegerRange(210, 31):
                d_133_v1_: int = compr_18_
                if ((210) <= (d_133_v1_)) and ((d_133_v1_) < (31)):
                    coll18_ = coll18_.union(_dafny.Set([(d_133_v1_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jrgri"))))]))
            return _dafny.Set(coll18_)
        def iife19_():
            def iife21_():
                coll21_ = _dafny.Map()
                compr_21_: _dafny.Seq
                for compr_21_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference([293])})).Elements:
                    d_134_v2_: _dafny.Seq = compr_21_
                    if (d_134_v2_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference([293])})):
                        coll21_[d_134_v2_] = True
                return _dafny.Map(coll21_)
            coll19_ = _dafny.Set()
            def iife20_():
                coll20_ = _dafny.Map()
                compr_20_: _dafny.Seq
                for compr_20_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference([293])})).Elements:
                    d_134_v2_: _dafny.Seq = compr_20_
                    if (d_134_v2_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference([293])})):
                        coll20_[d_134_v2_] = True
                return _dafny.Map(coll20_)
            compr_19_: _dafny.Seq
            for compr_19_ in (iife20_()
            ).keys.Elements:
                d_135_v3_: _dafny.Seq = compr_19_
                if (d_135_v3_) in (iife21_()
                ):
                    coll19_ = coll19_.union(_dafny.Set([d_135_v3_]))
            return _dafny.Set(coll19_)
        def iife22_():
            coll22_ = _dafny.Map()
            compr_22_: _dafny.Seq
            for compr_22_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "glco")): 914})).keys.Elements:
                d_138_v4_: _dafny.Seq = compr_22_
                if (d_138_v4_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "glco")): 914})):
                    coll22_[d_138_v4_] = False
            return _dafny.Map(coll22_)
        return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len((D19_DC50(True, (_dafny.MultiSet([107])).cardinality, _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([(0) - (len(iife17_()
)), len(_dafny.Set({74}))])).cardinality for d_132_i0_ in range(default__.abs(885))]))).cf96), 719, len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([False, True]), _dafny.SeqWithoutIsStrInference([True])]))]), (D11_DC27(True, _dafny.CodePoint('o'), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ycrugglv"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nokerh"))), _dafny.SeqWithoutIsStrInference([453]))).cf55, (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ggmughh"))), len(_dafny.SeqWithoutIsStrInference([(D3_DC6((_dafny.MultiSet([True, True])).cardinality, False)).cf13, True])), len(iife18_()
        ), (0) - ((0) - ((_dafny.MultiSet([True])).cardinality)), len(_dafny.Set({True, True, False}))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({False, False}))])), _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([len(iife19_()
        ), -985, (0) - (-569), len(_dafny.SeqWithoutIsStrInference([106 for d_136_i1_ in range(default__.abs(375))])), len(_dafny.Set({True}))])).cardinality]), (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: len(_dafny.Map({True: len(_dafny.Map({True: _dafny.CodePoint('m')}))}))})) for d_137_i2_ in range(default__.abs(-157))]) if False else _dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: (_dafny.MultiSet([True])).cardinality})), len(iife22_()
        ), len(_dafny.SeqWithoutIsStrInference([751])), len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vqoyap")))]))]))])

    @staticmethod
    def fm66(p0, globalState):
        return ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nnwdhunb")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dpcfmoy"))])) | (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ehco")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eo")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_139_i0_ in range(default__.abs(16))]), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_140_i1_ in range(default__.abs(487))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bj"))]))) | (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bs")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rlwyon"))]))

    @staticmethod
    def fm67(p0, globalState):
        return D8_DC19(_dafny.Map({False: (0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_141_i0_ in range(default__.abs(693))])))}), (_dafny.Set({True, True})) - (_dafny.Set({False, False})), (_dafny.MultiSet([(0) - (-601), -799])) | (_dafny.MultiSet([-453])), D1_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))))

    @staticmethod
    def m0(p0, p1, globalState):
        r0: _dafny.MultiSet = _dafny.MultiSet({})
        d_142_v0_: _dafny.Array
        nw0_ = _dafny.Array(int(0), 5)
        d_142_v0_ = nw0_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_142_v0_).length(0)):
            d_143_i0_: int = guard_loop_0_
            if (True) and (((0) <= (d_143_i0_)) and ((d_143_i0_) < ((d_142_v0_).length(0)))):
                (d_142_v0_)[(d_143_i0_)] = (d_143_i0_) + (605)
        d_144_i1_: int
        d_144_i1_ = 0
        with _dafny.label("0"):
            while (p1) or (p1):
                with _dafny.c_label("0"):
                    if (d_144_i1_) >= (100):
                        raise _dafny.Break("0")
                    d_144_i1_ = (d_144_i1_) + (1)
                    (globalState).f7 = -8
                    d_142_v0_ = d_142_v0_
                    d_145_v1_: int
                    d_145_v1_ = 859
                    d_146_v2_: _dafny.Seq
                    d_146_v2_ = _dafny.SeqWithoutIsStrInference([d_145_v1_])
                    d_147_v3_: _dafny.Map
                    d_147_v3_ = _dafny.Map({p1: d_146_v2_})
                    d_148_v4_: _dafny.Array
                    nw1_ = _dafny.Array(None, 2)
                    nw1_[int(0)] = d_147_v3_
                    nw1_[int(1)] = d_147_v3_
                    d_148_v4_ = nw1_
                    index0_ = default__.safeIndex(937, (d_148_v4_).length(0))
                    (d_148_v4_)[index0_] = d_147_v3_
                    index1_ = default__.safeIndex(937, (d_148_v4_).length(0))
                    (d_148_v4_)[index1_] = d_147_v3_
                    d_149_v5_: _dafny.Map
                    d_149_v5_ = _dafny.Map({d_145_v1_: p1})
                    d_150_v6_: _dafny.Array
                    nw2_ = _dafny.Array(None, 8)
                    nw2_[int(0)] = p1
                    nw2_[int(1)] = ((d_149_v5_)[d_145_v1_] if (d_145_v1_) in (d_149_v5_) else p1)
                    nw2_[int(2)] = False
                    nw2_[int(3)] = p1
                    nw2_[int(4)] = p1
                    nw2_[int(5)] = p1
                    nw2_[int(6)] = p1
                    nw2_[int(7)] = p1
                    d_150_v6_ = nw2_
                    d_151_v7_: _dafny.Set
                    d_151_v7_ = _dafny.Set({d_145_v1_})
                    d_152_v8_: D1
                    d_152_v8_ = D1_DC2(d_145_v1_, d_145_v1_, p1)
                    d_153_v9_: T2
                    nw3_ = C5()
                    nw3_.ctor__(p1, d_150_v6_, d_151_v7_, d_152_v8_, p1)
                    d_153_v9_ = nw3_
                    d_154_v10_: _dafny.Seq
                    d_154_v10_ = _dafny.SeqWithoutIsStrInference([d_153_v9_])
                    d_155_v11_: _dafny.Map
                    d_155_v11_ = _dafny.Map({d_154_v10_: d_145_v1_})
                    d_156_v12_: _dafny.Map
                    d_156_v12_ = _dafny.Map({p1: (499) < (len(d_155_v11_))})
                    d_156_v12_ = (d_156_v12_).set(p1, ((d_156_v12_)[p1] if (p1) in (d_156_v12_) else p1))
                    pass
            pass
        d_157_v13_: _dafny.Array
        nw4_ = _dafny.Array(None, 12)
        nw4_[int(0)] = p1
        nw4_[int(1)] = p1
        nw4_[int(2)] = p1
        nw4_[int(3)] = True
        nw4_[int(4)] = p1
        nw4_[int(5)] = p1
        nw4_[int(6)] = p1
        nw4_[int(7)] = p1
        nw4_[int(8)] = p1
        nw4_[int(9)] = p1
        nw4_[int(10)] = p1
        nw4_[int(11)] = p1
        d_157_v13_ = nw4_
        d_158_v14_: int
        d_158_v14_ = -466
        d_159_v15_: _dafny.Set
        d_159_v15_ = _dafny.Set({d_158_v14_, d_158_v14_})
        d_160_v16_: _dafny.Seq
        d_160_v16_ = _dafny.SeqWithoutIsStrInference([p1, p1])
        d_161_v17_: D1
        d_161_v17_ = D1_DC2(len(d_160_v16_), d_158_v14_, p1)
        d_162_v18_: C5
        nw5_ = C5()
        nw5_.ctor__(p1, d_157_v13_, d_159_v15_, d_161_v17_, p1)
        d_162_v18_ = nw5_
        d_163_v19_: _dafny.Seq
        d_163_v19_ = _dafny.SeqWithoutIsStrInference([d_162_v18_])
        d_164_v20_: _dafny.Map
        d_164_v20_ = _dafny.Map({not(p1): (d_163_v19_)[default__.safeIndex(d_158_v14_, len(d_163_v19_))]})
        d_165_v21_: T2
        nw6_ = C5()
        nw6_.ctor__(p1, d_157_v13_, d_159_v15_, d_161_v17_, p1)
        d_165_v21_ = nw6_
        d_166_v22_: _dafny.Set
        d_166_v22_ = _dafny.Set({d_165_v21_})
        hi0_ = len((d_166_v22_ if d_162_v18_.f24 else d_166_v22_))
        for d_167_i2_ in range(len(d_164_v20_), hi0_):
            d_168_v23_: _dafny.Map
            d_168_v23_ = _dafny.Map({d_162_v18_.f24: False})
            d_169_v24_: _dafny.Seq
            d_169_v24_ = _dafny.SeqWithoutIsStrInference([d_168_v23_])
            d_170_v25_: _dafny.Map
            d_170_v25_ = _dafny.Map({(d_165_v21_).f13: _dafny.SeqWithoutIsStrInference([d_168_v23_, d_168_v23_])})
            d_171_v26_: _dafny.Set
            d_171_v26_ = _dafny.Set({d_160_v16_, _dafny.SeqWithoutIsStrInference([False, (d_165_v21_).f13]), d_160_v16_})
            d_172_v27_: _dafny.Set
            d_172_v27_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([(d_162_v18_).fm52(d_167_i2_, d_171_v26_, globalState), d_162_v18_.f24])})
            d_173_v28_: C0
            nw7_ = C0()
            nw7_.ctor__((d_169_v24_) != (((d_170_v25_)[d_162_v18_.f24] if (d_162_v18_.f24) in (d_170_v25_) else d_169_v24_)), _dafny.CodePoint('b'), d_165_v21_.f12, (d_162_v18_).fm52(d_158_v14_, d_172_v27_, globalState))
            d_173_v28_ = nw7_
            d_174_v29_: _dafny.Seq
            d_174_v29_ = _dafny.SeqWithoutIsStrInference([d_173_v28_])
            d_173_v28_ = (d_174_v29_)[default__.safeIndex((d_173_v28_).fm16((d_165_v21_).f13, d_167_i2_, globalState), len(d_174_v29_))]
            d_175_v30_: _dafny.Array
            nw8_ = _dafny.Array(_dafny.Array(None, 0), 2)
            d_175_v30_ = nw8_
            d_176_v31_: _dafny.Array
            nw9_ = _dafny.Array(None, 24)
            nw9_[int(0)] = d_175_v30_
            nw9_[int(1)] = d_175_v30_
            nw9_[int(2)] = d_175_v30_
            nw9_[int(3)] = d_175_v30_
            nw9_[int(4)] = d_175_v30_
            nw9_[int(5)] = d_175_v30_
            nw9_[int(6)] = d_175_v30_
            nw9_[int(7)] = d_175_v30_
            nw9_[int(8)] = d_175_v30_
            nw9_[int(9)] = d_175_v30_
            nw9_[int(10)] = d_175_v30_
            nw9_[int(11)] = d_175_v30_
            nw9_[int(12)] = d_175_v30_
            nw9_[int(13)] = d_175_v30_
            nw9_[int(14)] = d_175_v30_
            nw9_[int(15)] = d_175_v30_
            nw9_[int(16)] = d_175_v30_
            nw9_[int(17)] = d_175_v30_
            nw9_[int(18)] = d_175_v30_
            nw9_[int(19)] = d_175_v30_
            nw9_[int(20)] = d_175_v30_
            nw9_[int(21)] = d_175_v30_
            nw9_[int(22)] = d_175_v30_
            nw9_[int(23)] = d_175_v30_
            d_176_v31_ = nw9_
            index2_ = default__.safeIndex(722, (d_176_v31_).length(0))
            (d_176_v31_)[index2_] = d_175_v30_
            index3_ = default__.safeIndex(722, (d_176_v31_).length(0))
            (d_176_v31_)[index3_] = d_175_v30_
            d_177_v32_: _dafny.Seq
            d_177_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qrcl"))
            d_177_v32_ = default__.fm20(globalState)
            (globalState).f7 = d_158_v14_
        (d_162_v18_).f24 = (not(d_162_v18_.f24) if p1 else (p1) and (d_162_v18_.f24))
        d_178_v33_: _dafny.Seq
        d_178_v33_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ulnyf"))
        d_179_v34_: _dafny.Map
        d_179_v34_ = _dafny.Map({len(d_178_v33_): d_158_v14_})
        (d_162_v18_).f24 = default__.fm30((0) - (len((d_179_v34_ if d_162_v18_.f24 else d_179_v34_))), globalState)
        d_180_v35_: D1
        d_180_v35_ = D1_DC1(d_178_v33_)
        d_181_v36_: C3
        nw10_ = C3()
        nw10_.ctor__((d_165_v21_).f13, d_180_v35_, d_162_v18_.f24, d_157_v13_, d_159_v15_, d_161_v17_, (d_162_v18_.f24 if p1 else p1))
        d_181_v36_ = nw10_
        pat_let_tv0_ = d_181_v36_
        pat_let_tv1_ = d_162_v18_
        pat_let_tv2_ = d_165_v21_
        pat_let_tv3_ = d_181_v36_
        pat_let_tv4_ = d_162_v18_
        pat_let_tv5_ = d_181_v36_
        pat_let_tv6_ = d_160_v16_
        pat_let_tv7_ = d_158_v14_
        pat_let_tv8_ = d_160_v16_
        def lambda0_(source6_):
            if source6_.is_DC30:
                d_182___mcc_h0_ = source6_.cf58
                d_183___mcc_h1_ = source6_.cf59
                d_184_cf59_ = d_183___mcc_h1_
                d_185_cf58_ = d_182___mcc_h0_
                d_186_dt__update__tmp_h0_ = _dafny.MultiSet([(pat_let_tv0_).f18, pat_let_tv1_.f24])
                d_187_dt__update_hcf0_h0_ = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(pat_let_tv2_).f13]))
                return d_187_dt__update_hcf0_h0_
            elif source6_.is_DC31:
                d_188___mcc_h2_ = source6_.cf60
                d_189___mcc_h3_ = source6_.cf61
                d_190___mcc_h4_ = source6_.cf62
                d_191___mcc_h5_ = source6_.cf63
                d_192___mcc_h6_ = source6_.cf64
                d_193_cf64_ = d_192___mcc_h6_
                d_194_cf63_ = d_191___mcc_h5_
                d_195_cf62_ = d_190___mcc_h4_
                d_196_cf61_ = d_189___mcc_h3_
                d_197_cf60_ = d_188___mcc_h2_
                return _dafny.MultiSet([(pat_let_tv3_).f18, d_197_cf60_])
            elif source6_.is_DC32:
                d_198___mcc_h7_ = source6_.cf65
                d_199___mcc_h8_ = source6_.cf66
                d_200___mcc_h9_ = source6_.cf67
                d_201___mcc_h10_ = source6_.cf68
                d_202___mcc_h11_ = source6_.cf69
                d_203_cf69_ = d_202___mcc_h11_
                d_204_cf68_ = d_201___mcc_h10_
                d_205_cf67_ = d_200___mcc_h9_
                d_206_cf66_ = d_199___mcc_h8_
                d_207_cf65_ = d_198___mcc_h7_
                return _dafny.MultiSet([not(True), d_205_cf67_, pat_let_tv4_.f24])
            elif True:
                d_208___mcc_h12_ = source6_.cf57
                d_209_cf57_ = d_208___mcc_h12_
                return _dafny.MultiSet([(pat_let_tv5_).f18, (pat_let_tv6_)[default__.safeIndex(pat_let_tv7_, len(pat_let_tv8_))]])

        r0 = lambda0_(D12_DC32(p1, len(d_178_v33_), (d_181_v36_).f18, d_162_v18_.f24, d_158_v14_))
        return r0

    @staticmethod
    def Main(noArgsParameter__):
        d_210_v1_: _dafny.Array
        def lambda1_(d_211_i0_):
            def iife23_():
                coll23_ = _dafny.Map()
                compr_23_: int
                for compr_23_ in (_dafny.Map({633: ((_dafny.MultiSet([False]))).cardinality})).keys.Elements:
                    d_212_v0_: int = compr_23_
                    if (d_212_v0_) in (_dafny.Map({633: ((_dafny.MultiSet([False]))).cardinality})):
                        coll23_[default__.safeDivisionInt(d_212_v0_, 805)] = False
                return _dafny.Map(coll23_)
            return iife23_()
            

        init0_ = lambda1_
        nw11_ = _dafny.Array(None, 2)
        for i0_0_ in range(nw11_.length(0)):
            nw11_[i0_0_] = init0_(i0_0_)
        d_210_v1_ = nw11_
        d_213_v2_: bool
        d_213_v2_ = True
        d_214_v3_: _dafny.Seq
        d_214_v3_ = _dafny.SeqWithoutIsStrInference([d_213_v2_])
        d_215_v4_: _dafny.Seq
        d_215_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dgxao"))
        d_216_v5_: _dafny.Array
        nw12_ = _dafny.Array(int(0), 10)
        d_216_v5_ = nw12_
        d_217_v6_: _dafny.MultiSet
        d_217_v6_ = _dafny.MultiSet([d_216_v5_, d_216_v5_])
        d_218_v7_: _dafny.Map
        d_218_v7_ = _dafny.Map({d_213_v2_: d_215_v4_})
        d_219_v8_: str
        d_219_v8_ = _dafny.CodePoint('g')
        d_220_v9_: int
        d_220_v9_ = 336
        d_221_v10_: _dafny.Map
        d_221_v10_ = _dafny.Map({d_220_v9_: d_215_v4_})
        d_222_v11_: D1
        d_222_v11_ = D1_DC1(d_215_v4_)
        d_223_v12_: _dafny.Array
        nw13_ = _dafny.Array(None, 28)
        nw13_[int(0)] = ((d_218_v7_)[d_213_v2_] if (d_213_v2_) in (d_218_v7_) else d_215_v4_)
        nw13_[int(1)] = d_215_v4_
        nw13_[int(2)] = d_215_v4_
        nw13_[int(3)] = d_215_v4_
        nw13_[int(4)] = d_215_v4_
        nw13_[int(5)] = d_215_v4_
        nw13_[int(6)] = d_215_v4_
        nw13_[int(7)] = d_215_v4_
        nw13_[int(8)] = d_215_v4_
        nw13_[int(9)] = d_215_v4_
        nw13_[int(10)] = d_215_v4_
        nw13_[int(11)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yauffnrew"))).set(default__.safeIndex(752, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yauffnrew")))), d_219_v8_)
        nw13_[int(12)] = d_215_v4_
        nw13_[int(13)] = d_215_v4_
        nw13_[int(14)] = ((d_221_v10_)[d_220_v9_] if (d_220_v9_) in (d_221_v10_) else d_215_v4_)
        nw13_[int(15)] = _dafny.SeqWithoutIsStrInference([d_219_v8_ for d_224_i1_ in range(default__.abs(51))])
        nw13_[int(16)] = d_215_v4_
        nw13_[int(17)] = d_215_v4_
        nw13_[int(18)] = (d_222_v11_).cf1
        nw13_[int(19)] = d_215_v4_
        nw13_[int(20)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tqk"))
        nw13_[int(21)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dgjly"))
        nw13_[int(22)] = d_215_v4_
        nw13_[int(23)] = _dafny.SeqWithoutIsStrInference([d_219_v8_ for d_225_i2_ in range(default__.abs(658))])
        nw13_[int(24)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dffircy"))
        nw13_[int(25)] = d_215_v4_
        nw13_[int(26)] = d_215_v4_
        nw13_[int(27)] = d_215_v4_
        d_223_v12_ = nw13_
        d_226_globalState_: GlobalState
        nw14_ = GlobalState()
        nw14_.ctor__(d_210_v1_, True, d_214_v3_, d_215_v4_, True, False, 142, 530, (d_215_v4_) + (d_215_v4_), d_217_v6_, d_223_v12_)
        d_226_globalState_ = nw14_
        d_227_v13_: _dafny.MultiSet
        out0_: _dafny.MultiSet
        out0_ = default__.m0(_dafny.CodePoint('g'), d_213_v2_, d_226_globalState_)
        d_227_v13_ = out0_
        d_228_v14_: _dafny.Array
        nw15_ = _dafny.Array(None, 6)
        nw15_[int(0)] = not(d_213_v2_)
        nw15_[int(1)] = not(d_213_v2_)
        nw15_[int(2)] = d_213_v2_
        nw15_[int(3)] = not(d_213_v2_)
        nw15_[int(4)] = True
        nw15_[int(5)] = d_213_v2_
        d_228_v14_ = nw15_
        d_229_v15_: D2
        d_229_v15_ = D2_DC3(d_228_v14_)
        d_230_v16_: _dafny.Seq
        d_230_v16_ = _dafny.SeqWithoutIsStrInference([d_228_v14_])
        d_231_v17_: _dafny.Set
        d_231_v17_ = _dafny.Set({d_213_v2_})
        d_232_v18_: _dafny.MultiSet
        d_232_v18_ = _dafny.MultiSet([len(d_231_v17_)])
        d_233_v19_: _dafny.Map
        d_233_v19_ = _dafny.Map({d_232_v18_: d_213_v2_})
        d_234_v20_: D2
        d_234_v20_ = D2_DC4(d_216_v5_, d_213_v2_, d_220_v9_, False, d_228_v14_)
        d_235_v21_: _dafny.Array
        nw16_ = _dafny.Array(None, 29)
        nw16_[int(0)] = (d_228_v14_ if d_213_v2_ else d_228_v14_)
        nw16_[int(1)] = d_228_v14_
        nw16_[int(2)] = d_228_v14_
        nw16_[int(3)] = d_228_v14_
        nw16_[int(4)] = d_228_v14_
        nw16_[int(5)] = d_228_v14_
        nw16_[int(6)] = d_228_v14_
        nw16_[int(7)] = d_228_v14_
        nw16_[int(8)] = d_228_v14_
        nw16_[int(9)] = d_228_v14_
        nw16_[int(10)] = d_228_v14_
        nw16_[int(11)] = d_228_v14_
        nw16_[int(12)] = d_228_v14_
        nw16_[int(13)] = d_228_v14_
        nw16_[int(14)] = (d_228_v14_ if d_213_v2_ else d_228_v14_)
        nw16_[int(15)] = (d_229_v15_).cf5
        nw16_[int(16)] = d_228_v14_
        nw16_[int(17)] = d_228_v14_
        nw16_[int(18)] = d_228_v14_
        nw16_[int(19)] = d_228_v14_
        nw16_[int(20)] = d_228_v14_
        nw16_[int(21)] = d_228_v14_
        nw16_[int(22)] = d_228_v14_
        nw16_[int(23)] = d_228_v14_
        nw16_[int(24)] = d_228_v14_
        nw16_[int(25)] = (d_230_v16_)[default__.safeIndex(len(d_233_v19_), len(d_230_v16_))]
        nw16_[int(26)] = d_228_v14_
        nw16_[int(27)] = d_228_v14_
        nw16_[int(28)] = (d_234_v20_).cf10
        d_235_v21_ = nw16_
        index4_ = default__.safeIndex(193, (d_235_v21_).length(0))
        (d_235_v21_)[index4_] = d_228_v14_
        index5_ = default__.safeIndex(193, (d_235_v21_).length(0))
        rhs0_ = d_228_v14_
        rhs1_ = default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "egift"))), 294)
        lhs0_ = d_235_v21_
        lhs1_ = default__.safeIndex(193, (d_235_v21_).length(0))
        lhs2_ = d_226_globalState_
        lhs0_[lhs1_] = rhs0_
        lhs2_.f7 = rhs1_
        index6_ = default__.safeIndex(777, (d_216_v5_).length(0))
        (d_216_v5_)[index6_] = d_220_v9_
        index7_ = default__.safeIndex(777, (d_216_v5_).length(0))
        (d_216_v5_)[index7_] = d_220_v9_
        d_236_v22_: _dafny.Map
        d_236_v22_ = _dafny.Map({d_213_v2_: d_214_v3_})
        hi1_ = (d_220_v9_) * (default__.fm0(((d_236_v22_)[d_213_v2_] if (d_213_v2_) in (d_236_v22_) else d_214_v3_), d_213_v2_, d_213_v2_, (d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))], d_226_globalState_))
        for d_237_i3_ in range(len(d_231_v17_), hi1_):
            d_213_v2_ = d_213_v2_
            source7_ = D1_DC1((d_215_v4_) + (d_215_v4_))
            if source7_.is_DC2:
                d_238___mcc_h0_ = source7_.cf2
                d_239___mcc_h1_ = source7_.cf3
                d_240___mcc_h2_ = source7_.cf4
                d_241_cf4_ = d_240___mcc_h2_
                d_242_cf3_ = d_239___mcc_h1_
                d_243_cf2_ = d_238___mcc_h0_
                d_244_v23_: _dafny.Map
                d_244_v23_ = _dafny.Map({len(d_215_v4_): d_228_v14_})
                d_245_v24_: D1
                d_245_v24_ = D1_DC2(87, len(d_215_v4_), d_241_cf4_)
                d_246_v25_: C2
                nw17_ = C2()
                nw17_.ctor__(d_242_cf3_, not(not(d_241_cf4_)), d_222_v11_, d_213_v2_, (d_228_v14_ if d_241_cf4_ else ((d_244_v23_)[d_243_cf2_] if (d_243_cf2_) in (d_244_v23_) else (d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))])), _dafny.Set({len(_dafny.Set({d_241_cf4_, d_241_cf4_}))}), d_245_v24_, d_241_cf4_)
                d_246_v25_ = nw17_
                d_247_v26_: _dafny.Map
                d_247_v26_ = _dafny.Map({d_241_cf4_: d_220_v9_})
                d_247_v26_ = (d_247_v26_).set(not((d_246_v25_).f22), (d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))])
                d_242_cf3_ = (d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))]
                d_216_v5_ = d_216_v5_
            elif True:
                d_248___mcc_h3_ = source7_.cf1
                d_249_cf1_ = d_248___mcc_h3_
                d_250_v27_: C5
                nw18_ = C5()
                nw18_.ctor__(d_213_v2_, d_228_v14_, default__.fm64(d_226_globalState_), default__.fm37((0) - ((0) - ((d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))])), d_237_i3_, (0) - (d_220_v9_), (d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))], d_226_globalState_), ((d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))]) > (d_237_i3_))
                d_250_v27_ = nw18_
                d_251_v29_: D1
                d_251_v29_ = D1_DC2(d_237_i3_, len(d_215_v4_), d_213_v2_)
                d_252_v30_: C2
                nw19_ = C2()
                def iife24_():
                    coll24_ = _dafny.Set()
                    compr_24_: int
                    for compr_24_ in _dafny.IntegerRange(536, 571):
                        d_253_v28_: int = compr_24_
                        if ((536) <= (d_253_v28_)) and ((d_253_v28_) < (571)):
                            coll24_ = coll24_.union(_dafny.Set([(d_253_v28_) + (d_220_v9_)]))
                    return _dafny.Set(coll24_)
                nw19_.ctor__(d_237_i3_, d_250_v27_.f24, D1_DC1(d_249_cf1_), d_213_v2_, (d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))], iife24_()
                , d_251_v29_, d_250_v27_.f24)
                d_252_v30_ = nw19_
                d_254_v31_: _dafny.Map
                d_254_v31_ = _dafny.Map({d_250_v27_.f24: d_252_v30_})
                d_255_v32_: _dafny.Seq
                d_255_v32_ = _dafny.SeqWithoutIsStrInference([len(d_254_v31_), (d_252_v30_).f21])
                d_256_v33_: _dafny.Array
                nw20_ = _dafny.Array(None, 5)
                nw20_[int(0)] = (D11_DC27(True, d_219_v8_, d_220_v9_, (d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))], d_255_v32_)).cf52
                nw20_[int(1)] = d_219_v8_
                nw20_[int(2)] = (d_215_v4_)[default__.safeIndex(-56, len(d_215_v4_))]
                nw20_[int(3)] = d_219_v8_
                nw20_[int(4)] = d_219_v8_
                d_256_v33_ = nw20_
                d_256_v33_ = d_256_v33_
                d_257_v34_: _dafny.Array
                def lambda2_(d_258_v30_, d_259_v32_):
                    def lambda3_(d_260_i4_):
                        return _dafny.MultiSet([D19_DC50((d_258_v30_).f22, -629, d_259_v32_), D19_DC50((d_258_v30_).f22, (d_258_v30_).f21, d_259_v32_)])

                    return lambda3_

                init1_ = lambda2_(d_252_v30_, d_255_v32_)
                nw21_ = _dafny.Array(None, 1)
                for i0_1_ in range(nw21_.length(0)):
                    nw21_[i0_1_] = init1_(i0_1_)
                d_257_v34_ = nw21_
                index8_ = default__.safeIndex(193, (d_235_v21_).length(0))
                rhs2_ = d_216_v5_
                rhs3_ = (d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))]
                rhs4_ = (d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))]
                rhs5_ = d_257_v34_
                lhs3_ = d_235_v21_
                lhs4_ = default__.safeIndex(193, (d_235_v21_).length(0))
                lhs5_ = d_226_globalState_
                d_216_v5_ = rhs2_
                lhs3_[lhs4_] = rhs3_
                lhs5_.f7 = rhs4_
                d_257_v34_ = rhs5_
                d_261_v35_: int
                out1_: int
                out1_ = (d_252_v30_).m6((d_252_v30_).f22, d_226_globalState_)
                d_261_v35_ = out1_
            d_262_v36_: _dafny.Set
            d_262_v36_ = _dafny.Set({d_220_v9_})
            d_263_v37_: D19
            d_263_v37_ = D19_DC49(d_262_v36_)
            pat_let_tv9_ = d_262_v36_
            arr0_ = (d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))]
            index9_ = default__.safeIndex(315, ((d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))]).length(0))
            def iife25_(_pat_let0_0):
                def iife26_(d_264_dt__update__tmp_h0_):
                    def iife27_(_pat_let1_0):
                        def iife28_(d_265_dt__update_hcf93_h0_):
                            return D19_DC49(d_265_dt__update_hcf93_h0_)
                        return iife28_(_pat_let1_0)
                    return iife27_(pat_let_tv9_)
                return iife26_(_pat_let0_0)
            arr0_[index9_] = ((iife25_(d_263_v37_)).cf93).isdisjoint(d_262_v36_)
            index10_ = default__.safeIndex(777, (d_216_v5_).length(0))
            arr1_ = (d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))]
            index11_ = default__.safeIndex(315, ((d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))]).length(0))
            rhs6_ = d_220_v9_
            rhs7_ = (d_220_v9_ if d_213_v2_ else -487)
            rhs8_ = not(True)
            rhs9_ = False
            lhs6_ = d_226_globalState_
            lhs7_ = d_216_v5_
            lhs8_ = default__.safeIndex(777, (d_216_v5_).length(0))
            lhs9_ = (d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))]
            lhs10_ = default__.safeIndex(315, ((d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))]).length(0))
            lhs6_.f7 = rhs6_
            lhs7_[lhs8_] = rhs7_
            d_213_v2_ = rhs8_
            lhs9_[lhs10_] = rhs9_
            index12_ = default__.safeIndex(777, (d_216_v5_).length(0))
            (d_216_v5_)[index12_] = d_220_v9_
        d_266_i5_: int
        d_266_i5_ = 0
        with _dafny.label("1"):
            while False:
                with _dafny.c_label("1"):
                    if (d_266_i5_) >= (100):
                        raise _dafny.Break("1")
                    d_266_i5_ = (d_266_i5_) + (1)
                    arr2_ = (d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))]
                    index13_ = default__.safeIndex(232, ((d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))]).length(0))
                    arr2_[index13_] = d_213_v2_
                    arr3_ = (d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))]
                    index14_ = default__.safeIndex(232, ((d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))]).length(0))
                    arr3_[index14_] = d_213_v2_
                    if d_213_v2_:
                        d_267_v38_: _dafny.Map
                        d_267_v38_ = _dafny.Map({(d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))]: d_220_v9_})
                        d_268_v40_: D1
                        d_268_v40_ = D1_DC2(d_220_v9_, (_dafny.MultiSet([d_214_v3_])).cardinality, not(((d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))])[default__.safeIndex(232, ((d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))]).length(0))]))
                        d_269_v41_: C6
                        nw22_ = C6()
                        def iife29_():
                            coll25_ = _dafny.Set()
                            compr_25_: int
                            for compr_25_ in (d_267_v38_).keys.Elements:
                                d_270_v39_: int = compr_25_
                                if (d_270_v39_) in (d_267_v38_):
                                    coll25_ = coll25_.union(_dafny.Set([(d_270_v39_) * ((0) - (-879))]))
                            return _dafny.Set(coll25_)
                        nw22_.ctor__(d_222_v11_, default__.fm17(d_226_globalState_), d_228_v14_, iife29_()
                        , d_268_v40_, ((d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))])[default__.safeIndex(232, ((d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))]).length(0))])
                        d_269_v41_ = nw22_
                        d_220_v9_ = len(d_231_v17_)
                        d_271_v42_: _dafny.MultiSet
                        d_271_v42_ = _dafny.MultiSet([d_228_v14_, d_228_v14_])
                        arr4_ = (d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))]
                        index15_ = default__.safeIndex(232, ((d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))]).length(0))
                        arr4_[index15_] = ((d_271_v42_) - (d_271_v42_)).issubset(d_271_v42_)
                        d_272_v43_: _dafny.Array
                        nw23_ = _dafny.Array(_dafny.Seq({}), 24)
                        d_272_v43_ = nw23_
                        d_272_v43_ = d_272_v43_
                        d_273_v44_: _dafny.Map
                        d_273_v44_ = _dafny.Map({d_219_v8_: d_220_v9_})
                        d_274_v45_: _dafny.Seq
                        d_274_v45_ = _dafny.SeqWithoutIsStrInference([d_273_v44_])
                        d_274_v45_ = d_274_v45_
                    elif True:
                        d_215_v4_ = (d_215_v4_) + (d_215_v4_)
                        d_275_v46_: C4
                        nw24_ = C4()
                        nw24_.ctor__()
                        d_275_v46_ = nw24_
                        d_276_v47_: _dafny.Map
                        d_276_v47_ = _dafny.Map({d_220_v9_: not(((d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))])[default__.safeIndex(232, ((d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))]).length(0))])})
                        d_277_v48_: _dafny.Map
                        d_277_v48_ = _dafny.Map({((d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))])[default__.safeIndex(232, ((d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))]).length(0))]: d_276_v47_})
                        d_278_v49_: _dafny.Map
                        d_278_v49_ = _dafny.Map({d_275_v46_: len(((d_277_v48_)[d_213_v2_] if (d_213_v2_) in (d_277_v48_) else _dafny.Map({d_220_v9_: ((d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))])[default__.safeIndex(232, ((d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))]).length(0))]})))})
                        d_279_v50_: D1
                        d_279_v50_ = D1_DC2((0) - ((d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))]), (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_220_v9_]))).cardinality, d_213_v2_)
                        d_280_v51_: C0
                        nw25_ = C0()
                        nw25_.ctor__(((d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))])[default__.safeIndex(232, ((d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))]).length(0))], (d_215_v4_)[default__.safeIndex(((d_278_v49_)[d_275_v46_] if (d_275_v46_) in (d_278_v49_) else 86), len(d_215_v4_))], d_279_v50_, (((d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))])[default__.safeIndex(232, ((d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))]).length(0))]) == (True))
                        d_280_v51_ = nw25_
                        d_280_v51_ = d_280_v51_
                        d_281_v52_: _dafny.Map
                        d_281_v52_ = _dafny.Map({d_220_v9_: -643})
                        d_282_v53_: _dafny.Seq
                        d_282_v53_ = _dafny.SeqWithoutIsStrInference([(0) - ((d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))]), len(d_281_v52_)])
                        d_283_v54_: _dafny.Set
                        d_283_v54_ = _dafny.Set({(d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))]})
                        d_284_v55_: C3
                        nw26_ = C3()
                        nw26_.ctor__(d_280_v51_.f19, d_222_v11_, default__.fm30(len(d_282_v53_), d_226_globalState_), d_228_v14_, d_283_v54_, d_279_v50_, ((d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))])[default__.safeIndex(232, ((d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))]).length(0))])
                        d_284_v55_ = nw26_
                        d_285_v56_: _dafny.Seq
                        d_285_v56_ = _dafny.SeqWithoutIsStrInference([d_284_v55_, d_284_v55_, d_284_v55_])
                        d_284_v55_ = (d_285_v56_)[default__.safeIndex(830, len(d_285_v56_))]
                        d_286_v57_: _dafny.MultiSet
                        d_286_v57_ = _dafny.MultiSet([d_213_v2_])
                        (d_280_v51_).f19 = ((d_286_v57_) - (d_286_v57_)).isdisjoint((d_286_v57_).intersection(d_286_v57_))
                        (d_226_globalState_).f2 = d_214_v3_
                    arr5_ = (d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))]
                    index16_ = default__.safeIndex(232, ((d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))]).length(0))
                    arr5_[index16_] = not(not(default__.fm3((d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))], d_226_globalState_)))
                    d_213_v2_ = False
                    pass
            pass
        d_287_v58_: _dafny.Map
        d_287_v58_ = _dafny.Map({d_213_v2_: d_213_v2_})
        index17_ = default__.safeIndex(919, (d_228_v14_).length(0))
        (d_228_v14_)[index17_] = ((d_287_v58_)[d_213_v2_] if (d_213_v2_) in (d_287_v58_) else d_213_v2_)
        index18_ = default__.safeIndex(919, (d_228_v14_).length(0))
        (d_228_v14_)[index18_] = ((d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))]) == (len(_dafny.SeqWithoutIsStrInference([d_213_v2_])))
        d_287_v58_ = (d_287_v58_).set(d_213_v2_, ((d_228_v14_)[default__.safeIndex(919, (d_228_v14_).length(0))]) == (d_213_v2_))
        hi2_ = d_220_v9_
        for d_288_i6_ in range(d_220_v9_, hi2_):
            d_213_v2_ = (d_230_v16_) <= (d_230_v16_)
            d_289_v59_: _dafny.MultiSet
            out2_: _dafny.MultiSet
            out2_ = default__.m0(d_219_v8_, default__.fm17(d_226_globalState_), d_226_globalState_)
            d_289_v59_ = out2_
            d_290_v60_: _dafny.Array
            nw27_ = _dafny.Array(_dafny.Array(None, 0), 22)
            d_290_v60_ = nw27_
            index19_ = default__.safeIndex(703, (d_290_v60_).length(0))
            (d_290_v60_)[index19_] = d_216_v5_
            index20_ = default__.safeIndex(703, (d_290_v60_).length(0))
            (d_290_v60_)[index20_] = d_216_v5_
            index21_ = default__.safeIndex(703, (d_290_v60_).length(0))
            (d_290_v60_)[index21_] = d_216_v5_
        d_291_v61_: D19
        d_291_v61_ = D19_DC50((d_228_v14_)[default__.safeIndex(919, (d_228_v14_).length(0))], (d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))], _dafny.SeqWithoutIsStrInference([len(default__.fm24(997, d_226_globalState_)), d_220_v9_]))
        d_292_v62_: _dafny.MultiSet
        out3_: _dafny.MultiSet
        out3_ = default__.m0(d_219_v8_, (d_291_v61_).cf94, d_226_globalState_)
        d_292_v62_ = out3_
        index22_ = default__.safeIndex(919, (d_228_v14_).length(0))
        rhs10_ = default__.fm0((d_214_v3_ if default__.fm17(d_226_globalState_) else d_214_v3_), (d_228_v14_)[default__.safeIndex(919, (d_228_v14_).length(0))], not((d_228_v14_)[default__.safeIndex(919, (d_228_v14_).length(0))]), default__.fm0(d_214_v3_, (d_228_v14_)[default__.safeIndex(919, (d_228_v14_).length(0))], (d_228_v14_)[default__.safeIndex(919, (d_228_v14_).length(0))], -259, d_226_globalState_), d_226_globalState_)
        rhs11_ = default__.fm3((d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))], d_226_globalState_)
        rhs12_ = len(_dafny.Map({(d_228_v14_)[default__.safeIndex(919, (d_228_v14_).length(0))]: d_213_v2_}))
        lhs11_ = d_226_globalState_
        lhs12_ = d_228_v14_
        lhs13_ = default__.safeIndex(919, (d_228_v14_).length(0))
        lhs11_.f7 = rhs10_
        lhs12_[lhs13_] = rhs11_
        d_220_v9_ = rhs12_
        d_293_v63_: _dafny.Array
        def lambda4_(d_294_v2_, d_295_v14_):
            def lambda5_(d_296_i8_):
                return _dafny.Map({not(not(d_294_v2_)): (d_295_v14_)[default__.safeIndex(919, (d_295_v14_).length(0))]})

            return lambda5_

        init2_ = lambda4_(d_213_v2_, d_228_v14_)
        nw28_ = _dafny.Array(None, 12)
        for i0_2_ in range(nw28_.length(0)):
            nw28_[i0_2_] = init2_(i0_2_)
        d_293_v63_ = nw28_
        _ingredientsd_0 = []
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_293_v63_).length(0)):
            d_297_i7_: int = guard_loop_1_
            if (True) and (((0) <= (d_297_i7_)) and ((d_297_i7_) < ((d_293_v63_).length(0)))):
                _ingredientsd_0.append((d_293_v63_, int(d_297_i7_), _dafny.Map({(d_231_v17_) == (d_231_v17_): (False if (d_228_v14_)[default__.safeIndex(919, (d_228_v14_).length(0))] else d_213_v2_)})))
        for _tupd_0 in _ingredientsd_0:
            _tupd_0[0][_tupd_0[1]] = _tupd_0[2]
        (d_226_globalState_).f7 = ((0) - ((d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))])) * (((d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))]) * ((d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))]))
        def iife30_():
            coll26_ = _dafny.Set()
            compr_26_: int
            for compr_26_ in (_dafny.SeqWithoutIsStrInference([(0) - ((d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))]) for d_299_i10_ in range(default__.abs(280))])).Elements:
                d_300_v64_: int = compr_26_
                if (d_300_v64_) in (_dafny.SeqWithoutIsStrInference([(0) - ((d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))]) for d_299_i10_ in range(default__.abs(280))])):
                    coll26_ = coll26_.union(_dafny.Set([(d_300_v64_) + (len(_dafny.SeqWithoutIsStrInference([51])))]))
            return _dafny.Set(coll26_)
        _ingredientsd_1 = []
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_228_v14_).length(0)):
            d_298_i9_: int = guard_loop_2_
            if (True) and (((0) <= (d_298_i9_)) and ((d_298_i9_) < ((d_228_v14_).length(0)))):
                _ingredientsd_1.append((d_228_v14_, int(d_298_i9_), ((iife30_()
                ) | (_dafny.Set({-321}))).ispropersubset((_dafny.Set({len(d_287_v58_)})) | (_dafny.Set({d_220_v9_})))))
        for _tupd_1 in _ingredientsd_1:
            _tupd_1[0][_tupd_1[1]] = _tupd_1[2]
        d_301_v65_: _dafny.Map
        d_301_v65_ = _dafny.Map({d_213_v2_: d_220_v9_})
        hi3_ = len(_dafny.SeqWithoutIsStrInference([(d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))]]))
        for d_302_i11_ in range(default__.safeModuloInt((d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))], len(d_301_v65_)), hi3_):
            d_303_v67_: _dafny.Seq
            d_303_v67_ = _dafny.SeqWithoutIsStrInference([d_220_v9_, (d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))]])
            d_304_v68_: D13
            d_304_v68_ = D13_DC34(d_220_v9_, default__.fm65(d_213_v2_, d_303_v67_, d_219_v8_, (d_228_v14_)[default__.safeIndex(919, (d_228_v14_).length(0))], d_226_globalState_))
            d_305_v69_: D1
            d_305_v69_ = D1_DC2(d_220_v9_, (d_304_v68_).cf71, d_213_v2_)
            d_306_v70_: T3
            nw29_ = C6()
            def iife31_():
                coll27_ = _dafny.Set()
                compr_27_: int
                for compr_27_ in _dafny.IntegerRange(163, 958):
                    d_307_v66_: int = compr_27_
                    if ((163) <= (d_307_v66_)) and ((d_307_v66_) < (958)):
                        coll27_ = coll27_.union(_dafny.Set([default__.safeModuloInt(d_307_v66_, d_302_i11_)]))
                return _dafny.Set(coll27_)
            nw29_.ctor__(d_222_v11_, (d_228_v14_)[default__.safeIndex(919, (d_228_v14_).length(0))], (d_235_v21_)[default__.safeIndex(193, (d_235_v21_).length(0))], iife31_()
            , d_305_v69_, (d_228_v14_)[default__.safeIndex(919, (d_228_v14_).length(0))])
            d_306_v70_ = nw29_
            d_308_v71_: _dafny.Map
            d_308_v71_ = _dafny.Map({d_302_i11_: d_306_v70_})
            pat_let_tv10_ = d_215_v4_
            d_309_v72_: T1
            nw30_ = C3()
            def iife32_(_pat_let2_0):
                def iife33_(d_310_dt__update__tmp_h1_):
                    def iife34_(_pat_let3_0):
                        def iife35_(d_311_dt__update_hcf1_h0_):
                            return D1_DC1(d_311_dt__update_hcf1_h0_)
                        return iife35_(_pat_let3_0)
                    return iife34_(pat_let_tv10_)
                return iife33_(_pat_let2_0)
            nw30_.ctor__(not((d_308_v71_) == (d_308_v71_)), iife32_(D1_DC1(d_215_v4_)), (d_306_v70_).f13, (d_306_v70_).f14, (d_306_v70_).f15, d_305_v69_, (d_228_v14_)[default__.safeIndex(919, (d_228_v14_).length(0))])
            d_309_v72_ = nw30_
            d_309_v72_ = (d_309_v72_ if (d_306_v70_).f17 else d_309_v72_)
            d_232_v18_ = d_232_v18_
            d_213_v2_ = (d_220_v9_) > (d_220_v9_)
            d_312_v73_: _dafny.Seq
            d_312_v73_ = _dafny.SeqWithoutIsStrInference([d_301_v65_])
            d_313_v74_: C8
            nw31_ = C8()
            nw31_.ctor__((d_306_v70_).f17)
            d_313_v74_ = nw31_
            d_314_v75_: _dafny.Set
            d_314_v75_ = _dafny.Set({d_313_v74_})
            d_315_v76_: _dafny.Array
            nw32_ = _dafny.Array(None, 16)
            nw32_[int(0)] = (d_301_v65_) | ((d_301_v65_).set((d_309_v72_).f13, d_302_i11_))
            nw32_[int(1)] = _dafny.Map({True: d_302_i11_})
            nw32_[int(2)] = (d_306_v70_).fm10((d_306_v70_).f13, (d_228_v14_)[default__.safeIndex(919, (d_228_v14_).length(0))], _dafny.SeqWithoutIsStrInference([(d_306_v70_).f17]), d_226_globalState_)
            nw32_[int(3)] = d_301_v65_
            nw32_[int(4)] = d_301_v65_
            nw32_[int(5)] = d_301_v65_
            nw32_[int(6)] = d_301_v65_
            nw32_[int(7)] = (d_301_v65_) | (d_301_v65_)
            nw32_[int(8)] = (d_301_v65_) | ((d_312_v73_)[default__.safeIndex((d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))], len(d_312_v73_))])
            nw32_[int(9)] = (d_301_v65_) | (d_301_v65_)
            nw32_[int(10)] = (_dafny.Map({(d_306_v70_).f17: len(d_314_v75_)})) | ((default__.fm67((d_228_v14_)[default__.safeIndex(919, (d_228_v14_).length(0))], d_226_globalState_)).cf32)
            nw32_[int(11)] = d_301_v65_
            nw32_[int(12)] = d_301_v65_
            nw32_[int(13)] = d_301_v65_
            nw32_[int(14)] = d_301_v65_
            nw32_[int(15)] = d_301_v65_
            d_315_v76_ = nw32_
            d_315_v76_ = d_315_v76_
        if d_213_v2_:
            d_316_v77_: _dafny.MultiSet
            out4_: _dafny.MultiSet
            out4_ = default__.m0((d_215_v4_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oxurkuj"))), len(d_215_v4_))], (d_228_v14_)[default__.safeIndex(919, (d_228_v14_).length(0))], d_226_globalState_)
            d_316_v77_ = out4_
            (d_226_globalState_).f7 = (d_220_v9_) * (d_220_v9_)
            d_317_v78_: _dafny.MultiSet
            out5_: _dafny.MultiSet
            out5_ = default__.m0(d_219_v8_, d_213_v2_, d_226_globalState_)
            d_317_v78_ = out5_
            d_318_v79_: D16
            d_318_v79_ = D16_DC41(not(d_213_v2_))
            d_213_v2_ = (d_318_v79_).cf84
            index23_ = default__.safeIndex(919, (d_228_v14_).length(0))
            (d_228_v14_)[index23_] = (d_213_v2_ if ((0) - (d_220_v9_)) < (d_220_v9_) else not((d_215_v4_) != (_dafny.SeqWithoutIsStrInference([d_219_v8_ for d_319_i12_ in range(default__.abs(-115))]))))
        elif True:
            d_320_v80_: _dafny.Map
            d_320_v80_ = _dafny.Map({_dafny.Map({(d_228_v14_)[default__.safeIndex(919, (d_228_v14_).length(0))]: d_231_v17_}): ((d_228_v14_)[default__.safeIndex(919, (d_228_v14_).length(0))]) or ((d_228_v14_)[default__.safeIndex(919, (d_228_v14_).length(0))])})
            d_321_v81_: _dafny.Map
            d_321_v81_ = _dafny.Map({(d_228_v14_)[default__.safeIndex(919, (d_228_v14_).length(0))]: d_231_v17_})
            d_320_v80_ = (d_320_v80_).set(d_321_v81_, d_213_v2_)
            d_215_v4_ = (d_215_v4_) + (((d_215_v4_ if d_213_v2_ else d_215_v4_)).set(default__.safeIndex((d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))], len((d_215_v4_ if d_213_v2_ else d_215_v4_))), d_219_v8_))
            index24_ = default__.safeIndex(919, (d_228_v14_).length(0))
            rhs13_ = False
            rhs14_ = (d_228_v14_)[default__.safeIndex(919, (d_228_v14_).length(0))]
            lhs14_ = d_228_v14_
            lhs15_ = default__.safeIndex(919, (d_228_v14_).length(0))
            d_213_v2_ = rhs13_
            lhs14_[lhs15_] = rhs14_
            d_322_v82_: _dafny.Map
            d_322_v82_ = _dafny.Map({d_219_v8_: (d_232_v18_).ispropersubset(d_232_v18_)})
            d_213_v2_ = ((d_322_v82_)[_dafny.CodePoint('x')] if (_dafny.CodePoint('x')) in (d_322_v82_) else d_213_v2_)
            d_323_v83_: _dafny.Array
            nw33_ = _dafny.Array(_dafny.MultiSet({}), 7)
            d_323_v83_ = nw33_
            d_324_v84_: _dafny.MultiSet
            d_324_v84_ = _dafny.MultiSet([d_215_v4_, d_215_v4_])
            index25_ = default__.safeIndex(931, (d_323_v83_).length(0))
            (d_323_v83_)[index25_] = d_324_v84_
            index26_ = default__.safeIndex(931, (d_323_v83_).length(0))
            (d_323_v83_)[index26_] = d_324_v84_
        d_325_v85_: _dafny.Map
        d_325_v85_ = _dafny.Map({d_220_v9_: (d_228_v14_)[default__.safeIndex(919, (d_228_v14_).length(0))]})
        d_326_v86_: _dafny.Seq
        d_326_v86_ = _dafny.SeqWithoutIsStrInference([(0) - (d_220_v9_), d_220_v9_, d_220_v9_, len(d_325_v85_)])
        hi4_ = (d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))]
        for d_327_i13_ in range((0) - (len((d_326_v86_) + (_dafny.SeqWithoutIsStrInference([(d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))]])))), hi4_):
            d_328_v87_: _dafny.Array
            nw34_ = _dafny.Array(D1.default()(), 4)
            d_328_v87_ = nw34_
            d_329_v88_: C9
            nw35_ = C9()
            nw35_.ctor__()
            d_329_v88_ = nw35_
            d_330_v89_: _dafny.Map
            d_330_v89_ = _dafny.Map({(d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))]: d_328_v87_})
            d_331_v90_: D1
            d_331_v90_ = D1_DC2(d_327_i13_, (d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))], d_213_v2_)
            d_332_v91_: C0
            nw36_ = C0()
            nw36_.ctor__((d_228_v14_)[default__.safeIndex(919, (d_228_v14_).length(0))], default__.fm58(d_226_globalState_), d_331_v90_, d_213_v2_)
            d_332_v91_ = nw36_
            d_333_v92_: _dafny.Set
            d_333_v92_ = _dafny.Set({d_332_v91_, d_332_v91_})
            d_334_v93_: _dafny.Map
            d_334_v93_ = _dafny.Map({d_329_v88_: ((d_330_v89_)[len((D20_DC52(d_333_v92_)).cf99)] if (len((D20_DC52(d_333_v92_)).cf99)) in (d_330_v89_) else d_328_v87_)})
            index27_ = default__.safeIndex(777, (d_216_v5_).length(0))
            rhs15_ = (d_220_v9_) + (d_220_v9_)
            rhs16_ = not((d_228_v14_)[default__.safeIndex(919, (d_228_v14_).length(0))])
            rhs17_ = (0) - ((0) - (d_327_i13_))
            rhs18_ = ((d_334_v93_)[d_329_v88_] if (d_329_v88_) in (d_334_v93_) else d_328_v87_)
            lhs16_ = d_216_v5_
            lhs17_ = default__.safeIndex(777, (d_216_v5_).length(0))
            lhs16_[lhs17_] = rhs15_
            d_213_v2_ = rhs16_
            d_220_v9_ = rhs17_
            d_328_v87_ = rhs18_
            d_326_v86_ = d_326_v86_
            d_220_v9_ = (d_220_v9_ if default__.fm17(d_226_globalState_) else (d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))])
            d_335_v94_: C4
            nw37_ = C4()
            nw37_.ctor__()
            d_335_v94_ = nw37_
            index28_ = default__.safeIndex(777, (d_216_v5_).length(0))
            rhs19_ = (d_216_v5_)[default__.safeIndex(777, (d_216_v5_).length(0))]
            rhs20_ = d_327_i13_
            rhs21_ = d_216_v5_
            rhs22_ = (d_335_v94_ if d_332_v91_.f19 else d_335_v94_)
            lhs18_ = d_226_globalState_
            lhs19_ = d_216_v5_
            lhs20_ = default__.safeIndex(777, (d_216_v5_).length(0))
            lhs18_.f7 = rhs19_
            lhs19_[lhs20_] = rhs20_
            d_216_v5_ = rhs21_
            d_335_v94_ = rhs22_
        _dafny.print(_dafny.string_of(((d_210_v1_)[0]) == (_dafny.Map({0: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_210_v1_)[1]) == (_dafny.Map({0: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_213_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_214_v3_) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_215_v4_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_216_v5_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_217_v6_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_218_v7_) == (_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dgxao"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_219_v8_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_220_v9_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_v10_) == (_dafny.Map({336: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dgxao"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_222_v11_).cf1).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_223_v12_)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_223_v12_)[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_223_v12_)[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_223_v12_)[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_223_v12_)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_223_v12_)[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_223_v12_)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_223_v12_)[7]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_223_v12_)[8]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_223_v12_)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_223_v12_)[10]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_223_v12_)[11]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_223_v12_)[12]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_223_v12_)[13]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_223_v12_)[14]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_223_v12_)[15]) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_223_v12_)[16]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_223_v12_)[17]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_223_v12_)[18]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_223_v12_)[19]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_223_v12_)[20]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_223_v12_)[21]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_223_v12_)[22]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_223_v12_)[23]) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_223_v12_)[24]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_223_v12_)[25]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_223_v12_)[26]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_223_v12_)[27]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_226_globalState_).f0)[0]) == (_dafny.Map({0: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_226_globalState_).f0)[1]) == (_dafny.Map({0: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_226_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_226_globalState_.f2) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_226_globalState_).f3).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_226_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_226_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_226_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_226_globalState_.f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_226_globalState_).f8).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_226_globalState_).f9).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_226_globalState_).f10)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_226_globalState_).f10)[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_226_globalState_).f10)[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_226_globalState_).f10)[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_226_globalState_).f10)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_226_globalState_).f10)[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_226_globalState_).f10)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_226_globalState_).f10)[7]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_226_globalState_).f10)[8]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_226_globalState_).f10)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_226_globalState_).f10)[10]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_226_globalState_).f10)[11]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_226_globalState_).f10)[12]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_226_globalState_).f10)[13]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_226_globalState_).f10)[14]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_226_globalState_).f10)[15]) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_226_globalState_).f10)[16]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_226_globalState_).f10)[17]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_226_globalState_).f10)[18]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_226_globalState_).f10)[19]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_226_globalState_).f10)[20]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_226_globalState_).f10)[21]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_226_globalState_).f10)[22]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_226_globalState_).f10)[23]) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_226_globalState_).f10)[24]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_226_globalState_).f10)[25]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_226_globalState_).f10)[26]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((((d_226_globalState_).f10)[27]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_227_v13_)) == (_dafny.MultiSet([False, False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v14_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v14_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v14_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v14_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v14_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_228_v14_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_229_v15_).cf5)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_229_v15_).cf5)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_229_v15_).cf5)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_229_v15_).cf5)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_229_v15_).cf5)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_229_v15_).cf5)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_230_v16_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_231_v17_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_232_v18_) == (_dafny.MultiSet([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_233_v19_) == (_dafny.Map({_dafny.MultiSet([1]): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_234_v20_).cf6)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_234_v20_).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_234_v20_).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_234_v20_).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_234_v20_).cf10)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_234_v20_).cf10)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_234_v20_).cf10)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_234_v20_).cf10)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_234_v20_).cf10)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_234_v20_).cf10)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[0])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[0])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[0])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[0])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[0])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[0])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[1])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[1])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[1])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[1])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[1])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[1])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[2])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[2])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[2])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[2])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[2])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[2])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[3])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[3])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[3])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[3])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[3])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[3])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[4])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[4])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[4])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[4])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[4])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[4])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[5])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[5])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[5])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[5])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[5])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[5])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[6])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[6])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[6])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[6])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[6])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[6])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[7])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[7])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[7])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[7])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[7])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[7])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[8])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[8])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[8])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[8])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[8])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[8])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[9])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[9])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[9])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[9])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[9])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[9])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[10])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[10])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[10])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[10])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[10])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[10])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[11])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[11])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[11])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[11])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[11])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[11])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[12])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[12])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[12])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[12])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[12])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[12])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[13])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[13])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[13])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[13])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[13])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[13])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[14])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[14])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[14])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[14])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[14])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[14])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[15])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[15])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[15])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[15])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[15])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[15])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[16])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[16])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[16])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[16])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[16])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[16])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[17])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[17])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[17])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[17])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[17])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[17])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[18])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[18])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[18])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[18])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[18])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[18])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[19])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[19])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[19])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[19])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[19])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[19])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[20])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[20])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[20])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[20])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[20])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[20])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[21])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[21])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[21])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[21])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[21])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[21])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[22])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[22])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[22])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[22])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[22])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[22])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[23])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[23])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[23])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[23])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[23])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[23])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[24])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[24])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[24])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[24])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[24])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[24])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[25])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[25])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[25])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[25])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[25])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[25])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[26])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[26])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[26])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[26])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[26])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[26])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[27])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[27])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[27])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[27])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[27])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[27])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[28])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[28])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[28])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[28])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[28])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_235_v21_)[28])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_236_v22_) == (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([True])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_266_i5_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_287_v58_) == (_dafny.Map({True: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_291_v61_).cf94))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_291_v61_).cf95))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_291_v61_).cf96) == (_dafny.SeqWithoutIsStrInference([3, 336]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_292_v62_)) == (_dafny.MultiSet([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_293_v63_)[0]) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_293_v63_)[1]) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_293_v63_)[2]) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_293_v63_)[3]) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_293_v63_)[4]) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_293_v63_)[5]) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_293_v63_)[6]) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_293_v63_)[7]) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_293_v63_)[8]) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_293_v63_)[9]) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_293_v63_)[10]) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_293_v63_)[11]) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_301_v65_) == (_dafny.Map({True: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_325_v85_) == (_dafny.Map({1: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_326_v86_) == (_dafny.SeqWithoutIsStrInference([-1, 1, 1, 1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
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
        return lambda: D1_DC2(int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D1_DC1)

class D1_DC2(D1, NamedTuple('DC2', [('cf2', Any), ('cf3', Any), ('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

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
        return lambda: D2_DC4(_dafny.Array(None, 0), False, int(0), False, _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D2_DC4)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D2_DC3)

class D2_DC4(D2, NamedTuple('DC4', [('cf6', Any), ('cf7', Any), ('cf8', Any), ('cf9', Any), ('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC3(D2, NamedTuple('DC3', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC3({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC3) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC6(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D3_DC6)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D3_DC5)

class D3_DC6(D3, NamedTuple('DC6', [('cf12', Any), ('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC6({_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC6) and self.cf12 == __o.cf12 and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC7(D3, NamedTuple('DC7', [])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7)
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC8(D3, NamedTuple('DC8', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC5(D3, NamedTuple('DC5', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC5({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC5) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D4_DC9)

class D4_DC9(D4, NamedTuple('DC9', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D5_DC10)

class D5_DC10(D5, NamedTuple('DC10', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC10({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC10) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC12(False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0), False)
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
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D6_DC14)

class D6_DC12(D6, NamedTuple('DC12', [('cf18', Any), ('cf19', Any), ('cf20', Any), ('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC12({_dafny.string_of(self.cf18)}, {self.cf19.VerbatimString(True)}, {_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC12) and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20 and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC13(D6, NamedTuple('DC13', [('cf22', Any), ('cf23', Any), ('cf24', Any), ('cf25', Any), ('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC13({_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC13) and self.cf22 == __o.cf22 and self.cf23 == __o.cf23 and self.cf24 == __o.cf24 and self.cf25 == __o.cf25 and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC11(D6, NamedTuple('DC11', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC11({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC11) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC14(D6, NamedTuple('DC14', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC14({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC14) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC16(int(0))
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

class D7_DC16(D7, NamedTuple('DC16', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC16({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC16) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC15(D7, NamedTuple('DC15', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC15({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC15) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC17(D7, NamedTuple('DC17', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC17({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC17) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC19(_dafny.Map({}), _dafny.Set({}), _dafny.MultiSet({}), D1.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D8_DC19)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D8_DC18)

class D8_DC19(D8, NamedTuple('DC19', [('cf32', Any), ('cf33', Any), ('cf34', Any), ('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC19({_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC19) and self.cf32 == __o.cf32 and self.cf33 == __o.cf33 and self.cf34 == __o.cf34 and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC18(D8, NamedTuple('DC18', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC18({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC18) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC21()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D9_DC21)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D9_DC22)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D9_DC23)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D9_DC20)

class D9_DC21(D9, NamedTuple('DC21', [])):
    def __dafnystr__(self) -> str:
        return f'D9.DC21'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC21)
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC22(D9, NamedTuple('DC22', [('cf37', Any), ('cf38', Any), ('cf39', Any), ('cf40', Any), ('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC22({_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)}, {_dafny.string_of(self.cf40)}, {self.cf41.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC22) and self.cf37 == __o.cf37 and self.cf38 == __o.cf38 and self.cf39 == __o.cf39 and self.cf40 == __o.cf40 and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC23(D9, NamedTuple('DC23', [('cf42', Any), ('cf43', Any), ('cf44', Any), ('cf45', Any), ('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC23({_dafny.string_of(self.cf42)}, {_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)}, {_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC23) and self.cf42 == __o.cf42 and self.cf43 == __o.cf43 and self.cf44 == __o.cf44 and self.cf45 == __o.cf45 and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC20(D9, NamedTuple('DC20', [('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC20({_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC20) and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D10_DC24)

class D10_DC24(D10, NamedTuple('DC24', [('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC24({_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC24) and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC26(_dafny.Array(None, 0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
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

class D11_DC26(D11, NamedTuple('DC26', [('cf49', Any), ('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC26({_dafny.string_of(self.cf49)}, {self.cf50.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC26) and self.cf49 == __o.cf49 and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC27(D11, NamedTuple('DC27', [('cf51', Any), ('cf52', Any), ('cf53', Any), ('cf54', Any), ('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC27({_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)}, {_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC27) and self.cf51 == __o.cf51 and self.cf52 == __o.cf52 and self.cf53 == __o.cf53 and self.cf54 == __o.cf54 and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC25(D11, NamedTuple('DC25', [('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC25({_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC25) and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC28(D11, NamedTuple('DC28', [('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC28({_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC28) and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC30(_dafny.Array(None, 0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D12_DC30)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D12_DC31)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D12_DC32)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D12_DC29)

class D12_DC30(D12, NamedTuple('DC30', [('cf58', Any), ('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC30({_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC30) and self.cf58 == __o.cf58 and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC31(D12, NamedTuple('DC31', [('cf60', Any), ('cf61', Any), ('cf62', Any), ('cf63', Any), ('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC31({_dafny.string_of(self.cf60)}, {_dafny.string_of(self.cf61)}, {_dafny.string_of(self.cf62)}, {_dafny.string_of(self.cf63)}, {_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC31) and self.cf60 == __o.cf60 and self.cf61 == __o.cf61 and self.cf62 == __o.cf62 and self.cf63 == __o.cf63 and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC32(D12, NamedTuple('DC32', [('cf65', Any), ('cf66', Any), ('cf67', Any), ('cf68', Any), ('cf69', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC32({_dafny.string_of(self.cf65)}, {_dafny.string_of(self.cf66)}, {_dafny.string_of(self.cf67)}, {_dafny.string_of(self.cf68)}, {_dafny.string_of(self.cf69)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC32) and self.cf65 == __o.cf65 and self.cf66 == __o.cf66 and self.cf67 == __o.cf67 and self.cf68 == __o.cf68 and self.cf69 == __o.cf69
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC29(D12, NamedTuple('DC29', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC29({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC29) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC34(int(0), _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D13_DC34)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D13_DC33)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D13_DC35)

class D13_DC34(D13, NamedTuple('DC34', [('cf71', Any), ('cf72', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC34({_dafny.string_of(self.cf71)}, {_dafny.string_of(self.cf72)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC34) and self.cf71 == __o.cf71 and self.cf72 == __o.cf72
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC33(D13, NamedTuple('DC33', [('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC33({_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC33) and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC35(D13, NamedTuple('DC35', [('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC35({_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC35) and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC37(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D14_DC37)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D14_DC36)

class D14_DC37(D14, NamedTuple('DC37', [('cf75', Any), ('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC37({self.cf75.VerbatimString(True)}, {_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC37) and self.cf75 == __o.cf75 and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC36(D14, NamedTuple('DC36', [('cf74', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC36({_dafny.string_of(self.cf74)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC36) and self.cf74 == __o.cf74
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC39(int(0), _dafny.CodePoint('D'), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D15_DC39)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D15_DC38)

class D15_DC39(D15, NamedTuple('DC39', [('cf78', Any), ('cf79', Any), ('cf80', Any), ('cf81', Any), ('cf82', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC39({_dafny.string_of(self.cf78)}, {_dafny.string_of(self.cf79)}, {self.cf80.VerbatimString(True)}, {_dafny.string_of(self.cf81)}, {_dafny.string_of(self.cf82)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC39) and self.cf78 == __o.cf78 and self.cf79 == __o.cf79 and self.cf80 == __o.cf80 and self.cf81 == __o.cf81 and self.cf82 == __o.cf82
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC38(D15, NamedTuple('DC38', [('cf77', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC38({_dafny.string_of(self.cf77)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC38) and self.cf77 == __o.cf77
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC41(False)
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

class D16_DC41(D16, NamedTuple('DC41', [('cf84', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC41({_dafny.string_of(self.cf84)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC41) and self.cf84 == __o.cf84
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC42(D16, NamedTuple('DC42', [('cf85', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC42({_dafny.string_of(self.cf85)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC42) and self.cf85 == __o.cf85
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC40(D16, NamedTuple('DC40', [('cf83', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC40({_dafny.string_of(self.cf83)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC40) and self.cf83 == __o.cf83
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC43(D16, NamedTuple('DC43', [('cf86', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC43({_dafny.string_of(self.cf86)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC43) and self.cf86 == __o.cf86
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC45(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D17_DC45)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D17_DC44)

class D17_DC45(D17, NamedTuple('DC45', [('cf88', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC45({_dafny.string_of(self.cf88)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC45) and self.cf88 == __o.cf88
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC44(D17, NamedTuple('DC44', [('cf87', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC44({_dafny.string_of(self.cf87)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC44) and self.cf87 == __o.cf87
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC47(int(0), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D18_DC47)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D18_DC46)

class D18_DC47(D18, NamedTuple('DC47', [('cf90', Any), ('cf91', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC47({_dafny.string_of(self.cf90)}, {_dafny.string_of(self.cf91)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC47) and self.cf90 == __o.cf90 and self.cf91 == __o.cf91
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC46(D18, NamedTuple('DC46', [('cf89', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC46({_dafny.string_of(self.cf89)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC46) and self.cf89 == __o.cf89
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC49(_dafny.Set({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D19_DC49)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D19_DC50)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D19_DC51)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D19_DC48)

class D19_DC49(D19, NamedTuple('DC49', [('cf93', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC49({_dafny.string_of(self.cf93)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC49) and self.cf93 == __o.cf93
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC50(D19, NamedTuple('DC50', [('cf94', Any), ('cf95', Any), ('cf96', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC50({_dafny.string_of(self.cf94)}, {_dafny.string_of(self.cf95)}, {_dafny.string_of(self.cf96)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC50) and self.cf94 == __o.cf94 and self.cf95 == __o.cf95 and self.cf96 == __o.cf96
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC51(D19, NamedTuple('DC51', [('cf97', Any), ('cf98', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC51({_dafny.string_of(self.cf97)}, {_dafny.string_of(self.cf98)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC51) and self.cf97 == __o.cf97 and self.cf98 == __o.cf98
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC48(D19, NamedTuple('DC48', [('cf92', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC48({_dafny.string_of(self.cf92)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC48) and self.cf92 == __o.cf92
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC53(None, False, int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D20_DC53)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D20_DC52)

class D20_DC53(D20, NamedTuple('DC53', [('cf100', Any), ('cf101', Any), ('cf102', Any), ('cf103', Any), ('cf104', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC53({_dafny.string_of(self.cf100)}, {_dafny.string_of(self.cf101)}, {_dafny.string_of(self.cf102)}, {_dafny.string_of(self.cf103)}, {_dafny.string_of(self.cf104)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC53) and self.cf100 == __o.cf100 and self.cf101 == __o.cf101 and self.cf102 == __o.cf102 and self.cf103 == __o.cf103 and self.cf104 == __o.cf104
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC52(D20, NamedTuple('DC52', [('cf99', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC52({_dafny.string_of(self.cf99)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC52) and self.cf99 == __o.cf99
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC55(False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D21_DC55)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D21_DC56)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D21_DC57)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D21_DC54)

class D21_DC55(D21, NamedTuple('DC55', [('cf106', Any), ('cf107', Any), ('cf108', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC55({_dafny.string_of(self.cf106)}, {_dafny.string_of(self.cf107)}, {_dafny.string_of(self.cf108)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC55) and self.cf106 == __o.cf106 and self.cf107 == __o.cf107 and self.cf108 == __o.cf108
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC56(D21, NamedTuple('DC56', [('cf109', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC56({_dafny.string_of(self.cf109)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC56) and self.cf109 == __o.cf109
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC57(D21, NamedTuple('DC57', [('cf110', Any), ('cf111', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC57({_dafny.string_of(self.cf110)}, {_dafny.string_of(self.cf111)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC57) and self.cf110 == __o.cf110 and self.cf111 == __o.cf111
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC54(D21, NamedTuple('DC54', [('cf105', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC54({_dafny.string_of(self.cf105)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC54) and self.cf105 == __o.cf105
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC59(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D22_DC59)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D22_DC60)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D22_DC58)

class D22_DC59(D22, NamedTuple('DC59', [('cf113', Any), ('cf114', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC59({_dafny.string_of(self.cf113)}, {_dafny.string_of(self.cf114)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC59) and self.cf113 == __o.cf113 and self.cf114 == __o.cf114
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC60(D22, NamedTuple('DC60', [('cf115', Any), ('cf116', Any), ('cf117', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC60({_dafny.string_of(self.cf115)}, {_dafny.string_of(self.cf116)}, {_dafny.string_of(self.cf117)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC60) and self.cf115 == __o.cf115 and self.cf116 == __o.cf116 and self.cf117 == __o.cf117
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC58(D22, NamedTuple('DC58', [('cf112', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC58({_dafny.string_of(self.cf112)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC58) and self.cf112 == __o.cf112
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    @property
    def f12(self):
        return self._f12
    @f12.setter
    def f12(self, value):
        self._f12 = value

class T1:
    pass

class T2:
    pass
    def m5(self, p0, globalState):
        pass


class T3:
    pass
    @property
    def f16(self):
        return self._f16
    @f16.setter
    def f16(self, value):
        self._f16 = value
    def m6(self, p0, globalState):
        pass

    def m7(self, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f2: _dafny.Seq = _dafny.Seq({})
        self.f7: int = int(0)
        self._f0: _dafny.Array = _dafny.Array(None, 0)
        self._f1: bool = False
        self._f3: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f4: bool = False
        self._f5: bool = False
        self._f6: int = int(0)
        self._f8: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f9: _dafny.MultiSet = _dafny.MultiSet({})
        self._f10: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10):
        (self)._f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self).f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10

    @property
    def f0(self):
        return self._f0
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
    def f8(self):
        return self._f8
    @property
    def f9(self):
        return self._f9
    @property
    def f10(self):
        return self._f10

class C0(T0):
    def  __init__(self):
        self._f12: D1 = D1.default()()
        self._f13: bool = False
        self.f19: bool = False
        self._f20: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    @property
    def f12(self):
        return self._f12
    @f12.setter
    def f12(self, value):
        self._f12 = value
    @property
    def f13(self):
        return self._f13
    def ctor__(self, f19, f20, f12, f13):
        (self).f19 = f19
        (self)._f20 = f20
        (self).f12 = f12
        (self)._f13 = f13

    def fm16(self, p0, p1, globalState):
        def iife36_():
            coll28_ = _dafny.Set()
            compr_28_: _dafny.Seq
            for compr_28_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bxjtdswbn")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")), _dafny.SeqWithoutIsStrInference([(self).f20 for d_336_i0_ in range(default__.abs(560))])])).Elements:
                d_337_v0_: _dafny.Seq = compr_28_
                if (d_337_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bxjtdswbn")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")), _dafny.SeqWithoutIsStrInference([(self).f20 for d_336_i0_ in range(default__.abs(560))])])):
                    coll28_ = coll28_.union(_dafny.Set([d_337_v0_]))
            return _dafny.Set(coll28_)
        return (0) - (default__.safeDivisionInt(len((iife36_()
        ) | (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vlirfxd"))}))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ivnwymujb")))))

    @property
    def f20(self):
        return self._f20

class C1(T1, T0):
    def  __init__(self):
        self._f12: D1 = D1.default()()
        self._f13: bool = False
        self.f23: str = _dafny.CodePoint('D')
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
    def f13(self):
        return self._f13
    def ctor__(self, f23, f12, f13):
        (self).f23 = f23
        (self).f12 = f12
        (self)._f13 = f13

    def fm8(self, p0, p1, globalState):
        return default__.safeDivisionInt(default__.safeModuloInt(len(_dafny.Set({_dafny.CodePoint('f'), self.f23, self.f23, self.f23})), len(_dafny.Map({not((self).f13): (self).f13}))), (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cug")))))

    def fm9(self, p0, globalState):
        return (len(_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f13, (self).f13, (self).f13]))]): 454}))) + (-762)

    def fm42(self, globalState):
        def iife37_():
            coll29_ = _dafny.Map()
            compr_29_: int
            for compr_29_ in _dafny.IntegerRange(-83, -824):
                d_338_v0_: int = compr_29_
                if ((-83) <= (d_338_v0_)) and ((d_338_v0_) < (-824)):
                    coll29_[(d_338_v0_) * ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lx")))))] = (self).f13
            return _dafny.Map(coll29_)
        def iife38_():
            coll30_ = _dafny.Map()
            compr_30_: int
            for compr_30_ in (_dafny.Set({436, 724})).Elements:
                d_339_v1_: int = compr_30_
                if (d_339_v1_) in (_dafny.Set({436, 724})):
                    coll30_[(d_339_v1_) + (len(_dafny.Map({931: 668})))] = len(_dafny.Map({D7_DC15(_dafny.SeqWithoutIsStrInference([477])): len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([(self).f13]))}))}))
            return _dafny.Map(coll30_)
        return _dafny.MultiSet([97, 693, default__.safeModuloInt((0) - (len(iife37_()
        )), (0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f13])), (_dafny.MultiSet([(self).f13])).cardinality])))), default__.safeDivisionInt(len(_dafny.Map({False: iife38_()
        })), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iik")))), default__.safeModuloInt(-784, -992)])

    def fm43(self, globalState):
        source8_ = D3_DC6(209, (self).f13)
        if source8_.is_DC6:
            d_340___mcc_h0_ = source8_.cf12
            d_341___mcc_h1_ = source8_.cf13
            d_342_cf13_ = d_341___mcc_h1_
            d_343_cf12_ = d_340___mcc_h0_
            return d_342_cf13_
        elif source8_.is_DC7:
            return (self).f13
        elif source8_.is_DC8:
            d_344___mcc_h2_ = source8_.cf14
            d_345_cf14_ = d_344___mcc_h2_
            return (self).f13
        elif True:
            d_346___mcc_h3_ = source8_.cf11
            d_347_cf11_ = d_346___mcc_h3_
            return True


class C2(T3, T2, T1, T0):
    def  __init__(self):
        self._f12: D1 = D1.default()()
        self._f16: D1 = D1.default()()
        self._f14: _dafny.Array = _dafny.Array(None, 0)
        self._f15: _dafny.Set = _dafny.Set({})
        self._f13: bool = False
        self._f17: bool = False
        self._f21: int = int(0)
        self._f22: bool = False
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
    def f16(self):
        return self._f16
    @f16.setter
    def f16(self, value):
        self._f16 = value
    @property
    def f14(self):
        return self._f14
    @property
    def f15(self):
        return self._f15
    @property
    def f13(self):
        return self._f13
    @property
    def f17(self):
        return self._f17
    def ctor__(self, f21, f22, f16, f17, f14, f15, f12, f13):
        (self)._f21 = f21
        (self)._f22 = f22
        (self).f16 = f16
        (self)._f17 = f17
        (self)._f14 = f14
        (self)._f15 = f15
        (self).f12 = f12
        (self)._f13 = f13

    def fm10(self, p0, p1, p2, globalState):
        return (_dafny.Map({not((self).f22): (self).f21})) | ((_dafny.Map({(self).f22: len(_dafny.Set({(self).f17}))})) | (_dafny.Map({(self).f22: (self).f21})))

    def fm8(self, p0, p1, globalState):
        return (self).f21

    def fm9(self, p0, globalState):
        return (self).f21

    def m6(self, p0, globalState):
        r0: int = int(0)
        d_348_i0_: int
        d_348_i0_ = 0
        with _dafny.label("2"):
            while (self).f22:
                with _dafny.c_label("2"):
                    if (d_348_i0_) >= (100):
                        raise _dafny.Break("2")
                    d_348_i0_ = (d_348_i0_) + (1)
                    d_349_v0_: _dafny.Array
                    nw38_ = _dafny.Array(int(0), 18)
                    d_349_v0_ = nw38_
                    d_350_v1_: _dafny.Seq
                    d_350_v1_ = _dafny.SeqWithoutIsStrInference([(self).f17])
                    d_351_v2_: str
                    d_351_v2_ = _dafny.CodePoint('k')
                    d_352_v3_: _dafny.MultiSet
                    d_352_v3_ = _dafny.MultiSet([_dafny.CodePoint('r'), d_351_v2_, d_351_v2_])
                    source9_ = D2_DC4(d_349_v0_, (self).f22, default__.fm0(d_350_v1_, False, (self).f13, (d_352_v3_).cardinality, globalState), (self).f22, (self).f14)
                    if source9_.is_DC4:
                        d_353___mcc_h0_ = source9_.cf6
                        d_354___mcc_h1_ = source9_.cf7
                        d_355___mcc_h2_ = source9_.cf8
                        d_356___mcc_h3_ = source9_.cf9
                        d_357___mcc_h4_ = source9_.cf10
                        d_358_cf10_ = d_357___mcc_h4_
                        d_359_cf9_ = d_356___mcc_h3_
                        d_360_cf8_ = d_355___mcc_h2_
                        d_361_cf7_ = d_354___mcc_h1_
                        d_362_cf6_ = d_353___mcc_h0_
                        d_363_v4_: C0
                        nw39_ = C0()
                        nw39_.ctor__(default__.fm30((self).f21, globalState), _dafny.CodePoint('p'), self.f12, d_359_cf9_)
                        d_363_v4_ = nw39_
                        d_364_v5_: _dafny.MultiSet
                        d_364_v5_ = _dafny.MultiSet([(self).f21, len(d_350_v1_)])
                        d_365_v6_: _dafny.Map
                        d_365_v6_ = _dafny.Map({D6_DC11(d_364_v5_): d_351_v2_})
                        d_366_v7_: D6
                        d_366_v7_ = D6_DC11(d_364_v5_)
                        d_367_v8_: _dafny.Seq
                        d_367_v8_ = _dafny.SeqWithoutIsStrInference([(self).f21])
                        d_368_v9_: _dafny.Seq
                        d_368_v9_ = _dafny.SeqWithoutIsStrInference([d_367_v8_])
                        d_369_v10_: _dafny.Map
                        d_369_v10_ = _dafny.Map({(d_367_v8_) + ((d_368_v9_)[default__.safeIndex(d_360_cf8_, len(d_368_v9_))]): d_360_cf8_})
                        rhs23_ = d_359_cf9_
                        rhs24_ = d_360_cf8_
                        rhs25_ = (d_363_v4_).fm16(d_361_cf7_, len((d_365_v6_) | (_dafny.Map({d_366_v7_: (d_363_v4_).f20}))), globalState)
                        rhs26_ = d_363_v4_
                        rhs27_ = ((d_369_v10_)[(d_367_v8_) + (d_367_v8_)] if ((d_367_v8_) + (d_367_v8_)) in (d_369_v10_) else (self).f21)
                        lhs21_ = globalState
                        lhs22_ = globalState
                        d_361_cf7_ = rhs23_
                        lhs21_.f7 = rhs24_
                        lhs22_.f7 = rhs25_
                        d_363_v4_ = rhs26_
                        r0 = rhs27_
                        d_370_v11_: D2
                        d_370_v11_ = D2_DC3(d_358_cf10_)
                        d_371_v12_: _dafny.Map
                        d_371_v12_ = _dafny.Map({(True if (self).f22 else (self).f22): d_370_v11_})
                        d_371_v12_ = (d_371_v12_).set((d_350_v1_)[default__.safeIndex((self).f21, len(d_350_v1_))], d_370_v11_)
                        d_372_v13_: _dafny.Map
                        d_372_v13_ = _dafny.Map({(d_367_v8_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(0) - ((self).f21) for d_373_i1_ in range(default__.abs(677))])), len(d_367_v8_))]: d_351_v2_})
                        d_374_v14_: _dafny.Map
                        d_374_v14_ = _dafny.Map({(self).f22: d_351_v2_})
                        d_372_v13_ = (d_372_v13_).set((0) - ((0) - (((self).f21 if d_361_cf7_ else d_360_cf8_))), ((d_374_v14_)[(self).f17] if ((self).f17) in (d_374_v14_) else (d_363_v4_).f20))
                        d_375_v15_: _dafny.Seq
                        d_375_v15_ = _dafny.SeqWithoutIsStrInference([(d_363_v4_).f20])
                        d_375_v15_ = d_375_v15_
                    elif True:
                        d_376___mcc_h5_ = source9_.cf5
                        d_377_cf5_ = d_376___mcc_h5_
                        index29_ = default__.safeIndex(639, (d_349_v0_).length(0))
                        (d_349_v0_)[index29_] = (self).f21
                        index30_ = default__.safeIndex(639, (d_349_v0_).length(0))
                        (d_349_v0_)[index30_] = (self).f21
                        d_378_v16_: _dafny.Map
                        d_378_v16_ = _dafny.Map({True: p0})
                        rhs28_ = ((d_349_v0_)[default__.safeIndex(639, (d_349_v0_).length(0))]) + (len((d_378_v16_) | (d_378_v16_)))
                        rhs29_ = (d_349_v0_)[default__.safeIndex(639, (d_349_v0_).length(0))]
                        lhs23_ = globalState
                        lhs24_ = globalState
                        lhs23_.f7 = rhs28_
                        lhs24_.f7 = rhs29_
                        d_379_v17_: _dafny.Map
                        d_379_v17_ = _dafny.Map({(self).f21: d_351_v2_})
                        d_380_v18_: _dafny.Map
                        d_380_v18_ = _dafny.Map({((self).f21) * (len((self).f15)): d_379_v17_})
                        d_380_v18_ = d_380_v18_
                        d_381_v19_: _dafny.MultiSet
                        out6_: _dafny.MultiSet
                        out6_ = default__.m0(d_351_v2_, not(False), globalState)
                        d_381_v19_ = out6_
                    (globalState).f2 = ((d_350_v1_) + (d_350_v1_)) + (d_350_v1_)
                    d_351_v2_ = d_351_v2_
                    d_382_v20_: C0
                    nw40_ = C0()
                    nw40_.ctor__((self).f17, d_351_v2_, self.f12, (self).f17)
                    d_382_v20_ = nw40_
                    pass
            pass
        d_383_v21_: _dafny.Map
        d_383_v21_ = _dafny.Map({66: (self).f13})
        d_384_v22_: _dafny.Map
        d_384_v22_ = _dafny.Map({len(d_383_v21_): not(False)})
        d_385_v23_: D6
        d_385_v23_ = D6_DC12((self).f13, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_386_i2_ in range(default__.abs(685))]), (self).f21, p0)
        d_387_v24_: _dafny.Seq
        d_387_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mpspf"))
        pat_let_tv11_ = p0
        pat_let_tv12_ = d_387_v24_
        pat_let_tv13_ = p0
        pat_let_tv14_ = d_387_v24_
        d_388_v25_: _dafny.Array
        nw41_ = _dafny.Array(None, 2)
        nw41_[int(0)] = d_383_v21_
        def iife40_(_pat_let5_0):
            def iife41_(d_389_dt__update__tmp_h0_):
                def iife42_(_pat_let6_0):
                    def iife43_(d_390_dt__update_hcf20_h0_):
                        def iife44_(_pat_let7_0):
                            def iife45_(d_391_dt__update_hcf18_h0_):
                                def iife46_(_pat_let8_0):
                                    def iife47_(d_392_dt__update_hcf19_h0_):
                                        return D6_DC12(d_391_dt__update_hcf18_h0_, d_392_dt__update_hcf19_h0_, d_390_dt__update_hcf20_h0_, (d_389_dt__update__tmp_h0_).cf21)
                                    return iife47_(_pat_let8_0)
                                return iife46_(pat_let_tv14_)
                            return iife45_(_pat_let7_0)
                        return iife44_(False)
                    return iife43_(_pat_let6_0)
                return iife42_((D6_DC12(pat_let_tv11_, pat_let_tv12_, (self).f21, not(pat_let_tv13_))).cf20)
            return iife41_(_pat_let5_0)
        def iife39_(_pat_let4_0):
            def iife48_(d_393_dt__update__tmp_h1_):
                def iife49_(_pat_let9_0):
                    def iife50_(d_394_dt__update_hcf20_h1_):
                        return D6_DC12((d_393_dt__update__tmp_h1_).cf18, (d_393_dt__update__tmp_h1_).cf19, d_394_dt__update_hcf20_h1_, (d_393_dt__update__tmp_h1_).cf21)
                    return iife50_(_pat_let9_0)
                return iife49_((self).f21)
            return iife48_(_pat_let4_0)
        nw41_[int(1)] = _dafny.Map({(self).f21: (iife39_(iife40_(d_385_v23_))).cf21})
        d_388_v25_ = nw41_
        d_395_v26_: str
        d_395_v26_ = _dafny.CodePoint('u')
        d_396_v27_: T0
        nw42_ = C0()
        nw42_.ctor__((self).f17, d_395_v26_, self.f12, p0)
        d_396_v27_ = nw42_
        d_397_v28_: _dafny.Seq
        d_397_v28_ = _dafny.SeqWithoutIsStrInference([(self).f15, (self).f15, (self).f15])
        d_398_v29_: _dafny.Map
        d_398_v29_ = _dafny.Map({(self).f13: p0})
        d_399_v30_: _dafny.MultiSet
        d_399_v30_ = _dafny.MultiSet([(self).f21, len(d_398_v29_)])
        d_400_v31_: D6
        d_400_v31_ = D6_DC11(d_399_v30_)
        d_401_v32_: _dafny.Map
        d_401_v32_ = _dafny.Map({d_396_v27_: (0) - (len((d_397_v28_)[default__.safeIndex(((d_400_v31_).cf17).cardinality, len(d_397_v28_))]))})
        d_402_v33_: _dafny.Seq
        d_402_v33_ = _dafny.SeqWithoutIsStrInference([d_401_v32_])
        d_403_v34_: _dafny.Seq
        d_403_v34_ = _dafny.SeqWithoutIsStrInference([(self).f21, (self).f21])
        d_404_v35_: _dafny.Seq
        d_404_v35_ = _dafny.SeqWithoutIsStrInference([len(d_403_v34_), (self).f21])
        d_405_v36_: _dafny.Map
        d_405_v36_ = _dafny.Map({len(d_404_v35_): d_387_v24_})
        rhs30_ = d_388_v25_
        rhs31_ = ((_dafny.Map({d_396_v27_: (self).f21})) | ((d_402_v33_)[default__.safeIndex((self).f21, len(d_402_v33_))])) | ((d_401_v32_) | (d_401_v32_))
        rhs32_ = len(((d_405_v36_)[(self).f21] if ((self).f21) in (d_405_v36_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uxadk"))))
        lhs25_ = globalState
        d_388_v25_ = rhs30_
        d_401_v32_ = rhs31_
        lhs25_.f7 = rhs32_
        (globalState).f7 = (self).f21
        d_384_v22_ = (d_384_v22_).set(default__.safeDivisionInt((self).fm9(d_395_v26_, globalState), (self).f21), (self).f17)
        d_406_v37_: C0
        nw43_ = C0()
        nw43_.ctor__(not((self).f17), d_395_v26_, self.f12, (d_396_v27_).f13)
        d_406_v37_ = nw43_
        if (self).f22:
            d_407_v38_: D3
            d_407_v38_ = D3_DC8((self).f21)
            d_408_v39_: _dafny.Map
            d_408_v39_ = _dafny.Map({(d_396_v27_).f13: (self).f21})
            d_409_v40_: _dafny.Map
            d_409_v40_ = _dafny.Map({(d_396_v27_).f13: d_395_v26_})
            d_410_v41_: _dafny.MultiSet
            d_410_v41_ = _dafny.MultiSet([(self).f22, (d_396_v27_).f13])
            d_411_v43_: _dafny.Map
            d_411_v43_ = _dafny.Map({(d_406_v37_).f20: (self).f21})
            d_412_v44_: _dafny.Seq
            def iife51_():
                coll31_ = _dafny.Map()
                compr_31_: str
                for compr_31_ in (d_411_v43_).keys.Elements:
                    d_413_v42_: str = compr_31_
                    if (d_413_v42_) in (d_411_v43_):
                        coll31_[d_413_v42_] = (self).f21
                return _dafny.Map(coll31_)
            d_412_v44_ = _dafny.SeqWithoutIsStrInference([iife51_()
            ])
            d_414_v45_: _dafny.Set
            d_414_v45_ = _dafny.Set({False, (d_396_v27_).f13})
            d_415_v46_: _dafny.Map
            d_415_v46_ = _dafny.Map({d_398_v29_: d_414_v45_})
            d_416_v47_: _dafny.Seq
            d_416_v47_ = _dafny.SeqWithoutIsStrInference([d_406_v37_.f19])
            d_417_v48_: _dafny.Array
            nw44_ = _dafny.Array(None, 23)
            def iife52_(_pat_let10_0):
                def iife53_(d_418_dt__update__tmp_h2_):
                    def iife54_(_pat_let11_0):
                        def iife55_(d_419_dt__update_hcf14_h0_):
                            return D3_DC8(d_419_dt__update_hcf14_h0_)
                        return iife55_(_pat_let11_0)
                    return iife54_(381)
                return iife53_(_pat_let10_0)
            nw44_[int(0)] = (iife52_(d_407_v38_)).cf14
            nw44_[int(1)] = ((0) - ((self).f21)) + ((self).f21)
            nw44_[int(2)] = len(((d_403_v34_).set(default__.safeIndex((self).f21, len(d_403_v34_)), (self).f21)).set(default__.safeIndex((self).f21, len((d_403_v34_).set(default__.safeIndex((self).f21, len(d_403_v34_)), (self).f21))), (self).f21))
            nw44_[int(3)] = (self).f21
            nw44_[int(4)] = (0) - ((self).f21)
            nw44_[int(5)] = (self).f21
            nw44_[int(6)] = default__.safeModuloInt((self).f21, (self).f21)
            nw44_[int(7)] = ((d_408_v39_)[(d_396_v27_).f13] if ((d_396_v27_).f13) in (d_408_v39_) else (self).f21)
            nw44_[int(8)] = 670
            nw44_[int(9)] = len(d_387_v24_)
            nw44_[int(10)] = (self).f21
            nw44_[int(11)] = (self).f21
            nw44_[int(12)] = (self).f21
            nw44_[int(13)] = len((_dafny.Map({d_387_v24_: d_409_v40_})) | (default__.fm31(globalState)))
            nw44_[int(14)] = (self).f21
            nw44_[int(15)] = ((d_410_v41_).cardinality) * (len(d_412_v44_))
            nw44_[int(16)] = default__.safeModuloInt((self).f21, (self).f21)
            nw44_[int(17)] = default__.safeDivisionInt(len(d_387_v24_), (0) - ((self).f21))
            nw44_[int(18)] = default__.safeDivisionInt((self).f21, (self).f21)
            nw44_[int(19)] = (len(_dafny.SeqWithoutIsStrInference([((d_415_v46_)[d_398_v29_] if (d_398_v29_) in (d_415_v46_) else _dafny.Set({d_406_v37_.f19}))]))) * ((self).f21)
            nw44_[int(20)] = default__.safeModuloInt((D3_DC6((0) - ((self).f21), (d_396_v27_).f13)).cf12, 703)
            nw44_[int(21)] = len((self).f15)
            nw44_[int(22)] = len((d_416_v47_) + (d_416_v47_))
            d_417_v48_ = nw44_
            d_420_v49_: _dafny.Map
            d_420_v49_ = _dafny.Map({d_395_v26_: d_406_v37_.f19})
            index31_ = default__.safeIndex(442, (d_417_v48_).length(0))
            (d_417_v48_)[index31_] = default__.safeDivisionInt(len(d_420_v49_), (self).f21)
            index32_ = default__.safeIndex(442, (d_417_v48_).length(0))
            (d_417_v48_)[index32_] = (0) - ((self).f21)
            d_421_v50_: C0
            nw45_ = C0()
            nw45_.ctor__((d_396_v27_).f13, (d_406_v37_).f20, d_396_v27_.f12, (self).f17)
            d_421_v50_ = nw45_
            d_422_v51_: _dafny.Array
            nw46_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 21)
            d_422_v51_ = nw46_
            index33_ = default__.safeIndex(949, (d_422_v51_).length(0))
            (d_422_v51_)[index33_] = d_387_v24_
            index34_ = default__.safeIndex(949, (d_422_v51_).length(0))
            (d_422_v51_)[index34_] = (d_387_v24_) + ((((d_405_v36_)[(d_417_v48_)[default__.safeIndex(442, (d_417_v48_).length(0))]] if ((d_417_v48_)[default__.safeIndex(442, (d_417_v48_).length(0))]) in (d_405_v36_) else _dafny.SeqWithoutIsStrInference([(d_406_v37_).f20 for d_423_i3_ in range(default__.abs(611))])) if False else d_387_v24_))
            d_417_v48_ = d_417_v48_
            (globalState).f7 = (d_417_v48_)[default__.safeIndex(442, (d_417_v48_).length(0))]
        elif True:
            d_424_v52_: _dafny.MultiSet
            d_424_v52_ = _dafny.MultiSet([(d_397_v28_)[default__.safeIndex((self).f21, len(d_397_v28_))]])
            d_425_v53_: _dafny.Seq
            d_425_v53_ = _dafny.SeqWithoutIsStrInference([(d_396_v27_).f13, (self).f22])
            d_426_v54_: _dafny.Map
            d_426_v54_ = _dafny.Map({(d_396_v27_).f13: d_424_v52_})
            d_424_v52_ = (_dafny.MultiSet([_dafny.Set({len(d_425_v53_)})])).intersection(((d_426_v54_)[not(p0)] if (not(p0)) in (d_426_v54_) else d_424_v52_))
            d_427_v56_: _dafny.Map
            def iife56_():
                coll32_ = _dafny.Set()
                compr_32_: int
                for compr_32_ in (_dafny.MultiSet([(self).f21])).Elements:
                    d_428_v55_: int = compr_32_
                    if (d_428_v55_) in (_dafny.MultiSet([(self).f21])):
                        coll32_ = coll32_.union(_dafny.Set([default__.safeModuloInt(d_428_v55_, 781)]))
                return _dafny.Set(coll32_)
            d_427_v56_ = _dafny.Map({(iife56_()
            ) | ((self).f15): _dafny.Map({d_387_v24_: (d_406_v37_).f20})})
            d_429_v57_: _dafny.Map
            d_429_v57_ = _dafny.Map({d_387_v24_: d_395_v26_})
            d_427_v56_ = (d_427_v56_).set((self).f15, d_429_v57_)
            d_430_v58_: _dafny.Array
            nw47_ = _dafny.Array(_dafny.Array(None, 0), 19)
            d_430_v58_ = nw47_
            d_431_v59_: _dafny.Map
            d_431_v59_ = _dafny.Map({d_425_v53_: (self).f21})
            d_432_v60_: D7
            d_432_v60_ = D7_DC15(_dafny.SeqWithoutIsStrInference([((d_431_v59_)[d_425_v53_] if (d_425_v53_) in (d_431_v59_) else (self).f21)]))
            d_433_v61_: _dafny.Array
            nw48_ = _dafny.Array(None, 16)
            nw48_[int(0)] = d_403_v34_
            nw48_[int(1)] = d_403_v34_
            nw48_[int(2)] = d_404_v35_
            nw48_[int(3)] = d_404_v35_
            nw48_[int(4)] = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.Map({(d_406_v37_).f20: (d_406_v37_).f20}): True})) for d_434_i4_ in range(default__.abs(389))])
            nw48_[int(5)] = d_403_v34_
            nw48_[int(6)] = d_404_v35_
            nw48_[int(7)] = (d_432_v60_).cf28
            nw48_[int(8)] = d_403_v34_
            nw48_[int(9)] = d_404_v35_
            nw48_[int(10)] = d_404_v35_
            nw48_[int(11)] = d_404_v35_
            nw48_[int(12)] = _dafny.SeqWithoutIsStrInference([(self).f21 for d_435_i5_ in range(default__.abs(954))])
            nw48_[int(13)] = d_403_v34_
            nw48_[int(14)] = d_403_v34_
            nw48_[int(15)] = d_403_v34_
            d_433_v61_ = nw48_
            index35_ = default__.safeIndex(140, (d_430_v58_).length(0))
            (d_430_v58_)[index35_] = d_433_v61_
            index36_ = default__.safeIndex(140, (d_430_v58_).length(0))
            (d_430_v58_)[index36_] = d_433_v61_
            (d_406_v37_).f19 = (((d_399_v30_)[(self).f21] if ((self).f21) in (d_399_v30_) else (self).f21)) >= ((len(d_387_v24_) if (self).f17 else (self).f21))
            if not(((d_387_v24_ if (self.f12).cf4 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "agacmtia")))) == (d_387_v24_)):
                d_436_v62_: _dafny.Seq
                d_436_v62_ = _dafny.SeqWithoutIsStrInference([(self).f14])
                d_437_v63_: _dafny.Set
                d_437_v63_ = _dafny.Set({False, d_406_v37_.f19, True})
                d_438_v64_: _dafny.Map
                d_438_v64_ = _dafny.Map({(self).f21: (d_436_v62_)[default__.safeIndex(len(d_437_v63_), len(d_436_v62_))]})
                d_439_v65_: D6
                d_439_v65_ = D6_DC11(d_399_v30_)
                d_440_v66_: D6
                d_440_v66_ = D6_DC14(d_439_v65_)
                d_441_v67_: _dafny.Array
                nw49_ = _dafny.Array(None, 20)
                nw49_[int(0)] = (self).f14
                nw49_[int(1)] = (self).f14
                nw49_[int(2)] = (self).f14
                nw49_[int(3)] = (self).f14
                nw49_[int(4)] = (self).f14
                nw49_[int(5)] = (self).f14
                nw49_[int(6)] = (self).f14
                nw49_[int(7)] = (self).f14
                nw49_[int(8)] = (self).f14
                nw49_[int(9)] = ((d_438_v64_)[(self).f21] if ((self).f21) in (d_438_v64_) else (self).f14)
                nw49_[int(10)] = (self).f14
                nw49_[int(11)] = (self).f14
                nw49_[int(12)] = (self).f14
                nw49_[int(13)] = (self).f14
                nw49_[int(14)] = (self).f14
                nw49_[int(15)] = (self).f14
                nw49_[int(16)] = (self).f14
                nw49_[int(17)] = (self).f14
                nw49_[int(18)] = (self).f14
                nw49_[int(19)] = (d_436_v62_)[default__.safeIndex(len(default__.fm32(d_440_v66_, globalState)), len(d_436_v62_))]
                d_441_v67_ = nw49_
                d_442_v68_: _dafny.Array
                d_442_v68_ = d_441_v67_
                d_442_v68_ = d_442_v68_
                (d_406_v37_).f19 = False
                (globalState).f7 = ((self).f21) + ((self).f21)
                d_443_v69_: _dafny.Array
                def lambda6_(d_444_i6_):
                    return default__.safeDivisionInt(d_444_i6_, (self).f21)

                init3_ = lambda6_
                nw50_ = _dafny.Array(None, 17)
                for i0_3_ in range(nw50_.length(0)):
                    nw50_[i0_3_] = init3_(i0_3_)
                d_443_v69_ = nw50_
                d_443_v69_ = (d_443_v69_ if (self).f22 else d_443_v69_)
                d_445_v70_: _dafny.Seq
                d_445_v70_ = _dafny.SeqWithoutIsStrInference([d_385_v23_])
                d_446_v71_: _dafny.Seq
                d_446_v71_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_385_v23_ for d_447_i7_ in range(default__.abs(541))]), d_445_v70_])
                d_448_v72_: _dafny.Map
                d_448_v72_ = _dafny.Map({len(_dafny.Map({p0: d_406_v37_.f19})): (d_446_v71_)[default__.safeIndex((self).f21, len(d_446_v71_))]})
                d_449_v73_: C0
                nw51_ = C0()
                nw51_.ctor__(((d_396_v27_).f13 if (self).f17 else (d_396_v27_).f13), d_395_v26_, d_396_v27_.f12, (693) not in (d_448_v72_))
                d_449_v73_ = nw51_
            elif True:
                index37_ = default__.safeIndex(376, ((self).f14).length(0))
                ((self).f14)[index37_] = (d_396_v27_).f13
                d_450_v74_: _dafny.Map
                d_450_v74_ = _dafny.Map({(self).f21: (self).f21})
                d_451_v75_: _dafny.Set
                d_451_v75_ = _dafny.Set({d_406_v37_.f19, p0, (self).f17, False})
                d_452_v76_: _dafny.Seq
                d_452_v76_ = _dafny.SeqWithoutIsStrInference([d_387_v24_])
                index38_ = default__.safeIndex(376, ((self).f14).length(0))
                rhs33_ = (d_396_v27_).f13
                rhs34_ = (d_451_v75_) == ((d_451_v75_) | (d_451_v75_))
                rhs35_ = d_450_v74_
                rhs36_ = ((d_425_v53_)[default__.safeIndex((self).f21, len(d_425_v53_))] if (self).f22 else not(((self).f21) >= (len(d_452_v76_))))
                lhs26_ = d_406_v37_
                lhs27_ = (self).f14
                lhs28_ = default__.safeIndex(376, ((self).f14).length(0))
                lhs29_ = d_406_v37_
                lhs26_.f19 = rhs33_
                lhs27_[lhs28_] = rhs34_
                d_450_v74_ = rhs35_
                lhs29_.f19 = rhs36_
                d_387_v24_ = d_387_v24_
                d_453_v77_: _dafny.Array
                def lambda7_(d_454_i8_):
                    return (d_454_i8_) - ((self).f21)

                init4_ = lambda7_
                nw52_ = _dafny.Array(None, 21)
                for i0_4_ in range(nw52_.length(0)):
                    nw52_[i0_4_] = init4_(i0_4_)
                d_453_v77_ = nw52_
                d_453_v77_ = d_453_v77_
                d_455_v78_: _dafny.Array
                def lambda8_(d_456_v37_):
                    def lambda9_(d_457_i9_):
                        return (d_456_v37_).f20

                    return lambda9_

                init5_ = lambda8_(d_406_v37_)
                nw53_ = _dafny.Array(None, 5)
                for i0_5_ in range(nw53_.length(0)):
                    nw53_[i0_5_] = init5_(i0_5_)
                d_455_v78_ = nw53_
                index39_ = default__.safeIndex(178, (d_455_v78_).length(0))
                (d_455_v78_)[index39_] = (d_406_v37_).f20
                index40_ = default__.safeIndex(178, (d_455_v78_).length(0))
                rhs37_ = (d_406_v37_).f20
                rhs38_ = d_425_v53_
                rhs39_ = (d_406_v37_).f20
                lhs30_ = d_455_v78_
                lhs31_ = default__.safeIndex(178, (d_455_v78_).length(0))
                lhs32_ = globalState
                lhs30_[lhs31_] = rhs37_
                lhs32_.f2 = rhs38_
                d_395_v26_ = rhs39_
                (globalState).f7 = (0) - ((self).f21)
        r0 = 139
        return r0

    def m7(self, globalState):
        d_458_v0_: _dafny.Seq
        d_458_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rxto"))
        d_459_v1_: _dafny.Array
        def lambda10_(d_460_i0_):
            return (d_460_i0_) * (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_461_i1_ in range(default__.abs(-952))])))

        init6_ = lambda10_
        nw54_ = _dafny.Array(None, 9)
        for i0_6_ in range(nw54_.length(0)):
            nw54_[i0_6_] = init6_(i0_6_)
        d_459_v1_ = nw54_
        d_462_v2_: _dafny.Seq
        d_462_v2_ = _dafny.SeqWithoutIsStrInference([d_459_v1_, d_459_v1_, d_459_v1_])
        d_463_v3_: D6
        d_463_v3_ = D6_DC12((self).f17, d_458_v0_, len(d_462_v2_), (self).f22)
        source10_ = d_463_v3_
        if source10_.is_DC12:
            d_464___mcc_h0_ = source10_.cf18
            d_465___mcc_h1_ = source10_.cf19
            d_466___mcc_h2_ = source10_.cf20
            d_467___mcc_h3_ = source10_.cf21
            d_468_cf21_ = d_467___mcc_h3_
            d_469_cf20_ = d_466___mcc_h2_
            d_470_cf19_ = d_465___mcc_h1_
            d_471_cf18_ = d_464___mcc_h0_
            index41_ = default__.safeIndex(186, (d_459_v1_).length(0))
            (d_459_v1_)[index41_] = (self).f21
            index42_ = default__.safeIndex(186, (d_459_v1_).length(0))
            (d_459_v1_)[index42_] = d_469_cf20_
            d_472_v4_: _dafny.Map
            d_472_v4_ = _dafny.Map({d_469_cf20_: (self).f14})
            d_473_v5_: _dafny.Map
            d_473_v5_ = _dafny.Map({d_468_cf21_: len(d_472_v4_)})
            d_474_v6_: _dafny.Map
            d_474_v6_ = _dafny.Map({d_473_v5_: (_dafny.MultiSet([877])).cardinality})
            d_475_v7_: _dafny.Map
            d_475_v7_ = _dafny.Map({d_471_cf18_: False})
            d_476_v8_: _dafny.Seq
            d_476_v8_ = _dafny.SeqWithoutIsStrInference([(self).f21, (d_459_v1_)[default__.safeIndex(186, (d_459_v1_).length(0))], d_469_cf20_, (self).f21, d_469_cf20_])
            d_477_v9_: _dafny.MultiSet
            d_477_v9_ = _dafny.MultiSet([d_469_cf20_])
            d_478_v10_: _dafny.Array
            nw55_ = _dafny.Array(None, 18)
            nw55_[int(0)] = (d_459_v1_)[default__.safeIndex(186, (d_459_v1_).length(0))]
            nw55_[int(1)] = len(d_470_cf19_)
            nw55_[int(2)] = ((d_459_v1_)[default__.safeIndex(186, (d_459_v1_).length(0))] if True else ((d_474_v6_)[d_473_v5_] if (d_473_v5_) in (d_474_v6_) else (self).f21))
            nw55_[int(3)] = ((d_459_v1_)[default__.safeIndex(186, (d_459_v1_).length(0))]) + (len(d_475_v7_))
            nw55_[int(4)] = (d_459_v1_)[default__.safeIndex(186, (d_459_v1_).length(0))]
            nw55_[int(5)] = d_469_cf20_
            nw55_[int(6)] = (d_469_cf20_) - ((self).f21)
            nw55_[int(7)] = (d_459_v1_)[default__.safeIndex(186, (d_459_v1_).length(0))]
            nw55_[int(8)] = (self).f21
            nw55_[int(9)] = len(((d_476_v8_).set(default__.safeIndex((d_459_v1_)[default__.safeIndex(186, (d_459_v1_).length(0))], len(d_476_v8_)), d_469_cf20_)) + (d_476_v8_))
            nw55_[int(10)] = (d_469_cf20_) * (len((d_476_v8_).set(default__.safeIndex(d_469_cf20_, len(d_476_v8_)), d_469_cf20_)))
            nw55_[int(11)] = default__.safeModuloInt(d_469_cf20_, d_469_cf20_)
            nw55_[int(12)] = (547 if True else d_469_cf20_)
            nw55_[int(13)] = (d_477_v9_).cardinality
            nw55_[int(14)] = (d_476_v8_)[default__.safeIndex(d_469_cf20_, len(d_476_v8_))]
            nw55_[int(15)] = 865
            nw55_[int(16)] = d_469_cf20_
            nw55_[int(17)] = (self).f21
            d_478_v10_ = nw55_
            d_478_v10_ = (D2_DC4(d_478_v10_, True, (self).f21, (self).f22, (self).f14)).cf6
            (globalState).f7 = 527
            d_479_v11_: _dafny.Seq
            d_479_v11_ = _dafny.SeqWithoutIsStrInference([not((self).f13), d_471_cf18_])
            d_480_v12_: C0
            nw56_ = C0()
            nw56_.ctor__(((d_479_v11_)[default__.safeIndex((0) - ((self).f21), len(d_479_v11_))] if (self).f22 else False), (d_458_v0_)[default__.safeIndex((d_459_v1_)[default__.safeIndex(186, (d_459_v1_).length(0))], len(d_458_v0_))], self.f12, d_471_cf18_)
            d_480_v12_ = nw56_
        elif source10_.is_DC13:
            d_481___mcc_h4_ = source10_.cf22
            d_482___mcc_h5_ = source10_.cf23
            d_483___mcc_h6_ = source10_.cf24
            d_484___mcc_h7_ = source10_.cf25
            d_485___mcc_h8_ = source10_.cf26
            d_486_cf26_ = d_485___mcc_h8_
            d_487_cf25_ = d_484___mcc_h7_
            d_488_cf24_ = d_483___mcc_h6_
            d_489_cf23_ = d_482___mcc_h5_
            d_490_cf22_ = d_481___mcc_h4_
            d_489_cf23_ = (self).f22
            d_491_v13_: _dafny.Set
            d_491_v13_ = _dafny.Set({(self).f17})
            d_492_v14_: bool
            out7_: bool
            out7_ = (self).m13((self).f13, default__.fm33((self).f13, d_487_cf25_, globalState), d_491_v13_, (0) - (d_488_cf24_), globalState)
            d_492_v14_ = out7_
            index43_ = default__.safeIndex(295, (d_459_v1_).length(0))
            (d_459_v1_)[index43_] = d_488_cf24_
            d_493_v15_: _dafny.Seq
            d_493_v15_ = _dafny.SeqWithoutIsStrInference([(self).f17])
            index44_ = default__.safeIndex(295, (d_459_v1_).length(0))
            (d_459_v1_)[index44_] = (0) - (len(_dafny.Set({not(((self).f22 if (d_493_v15_)[default__.safeIndex((self).f21, len(d_493_v15_))] else not(d_489_cf23_))), (d_487_cf25_) == ((self).f21), (d_486_cf26_) or (d_492_v14_)})))
            d_492_v14_ = (d_489_cf23_) == ((self).f13)
        elif source10_.is_DC11:
            d_494___mcc_h9_ = source10_.cf17
            d_495_cf17_ = d_494___mcc_h9_
            d_496_v16_: _dafny.Map
            d_496_v16_ = _dafny.Map({True: (self).f21})
            (globalState).f7 = ((d_496_v16_)[(self).f22] if ((self).f22) in (d_496_v16_) else (self).f21)
            d_497_v17_: _dafny.Map
            d_497_v17_ = _dafny.Map({(self).f14: self.f12})
            d_498_v18_: str
            d_498_v18_ = _dafny.CodePoint('c')
            d_499_v19_: C0
            nw57_ = C0()
            nw57_.ctor__((d_497_v17_) != (d_497_v17_), d_498_v18_, D1_DC2((self).f21, (self).f21, True), (self).f17)
            d_499_v19_ = nw57_
            d_500_v20_: D6
            d_500_v20_ = D6_DC11(d_495_cf17_)
            source11_ = d_500_v20_
            if source11_.is_DC12:
                d_501___mcc_h11_ = source11_.cf18
                d_502___mcc_h12_ = source11_.cf19
                d_503___mcc_h13_ = source11_.cf20
                d_504___mcc_h14_ = source11_.cf21
                d_505_cf21_ = d_504___mcc_h14_
                d_506_cf20_ = d_503___mcc_h13_
                d_507_cf19_ = d_502___mcc_h12_
                d_508_cf18_ = d_501___mcc_h11_
                d_509_v21_: _dafny.Array
                def lambda11_(d_510_i2_):
                    return _dafny.CodePoint('t')

                init7_ = lambda11_
                nw58_ = _dafny.Array(None, 5)
                for i0_7_ in range(nw58_.length(0)):
                    nw58_[i0_7_] = init7_(i0_7_)
                d_509_v21_ = nw58_
                index45_ = default__.safeIndex(445, (d_509_v21_).length(0))
                (d_509_v21_)[index45_] = d_498_v18_
                index46_ = default__.safeIndex(445, (d_509_v21_).length(0))
                (d_509_v21_)[index46_] = _dafny.CodePoint('v')
                d_511_v22_: _dafny.MultiSet
                d_511_v22_ = _dafny.MultiSet([(self).f13, d_508_cf18_])
                d_512_v23_: _dafny.Seq
                d_512_v23_ = _dafny.SeqWithoutIsStrInference([d_508_cf18_, d_505_cf21_])
                d_513_v24_: _dafny.Seq
                d_513_v24_ = _dafny.SeqWithoutIsStrInference([d_506_cf20_, len(d_458_v0_)])
                d_514_v25_: D6
                d_514_v25_ = D6_DC14(D6_DC11(_dafny.MultiSet(d_513_v24_)))
                d_515_v26_: D6
                d_515_v26_ = D6_DC14(d_514_v25_)
                d_516_v27_: _dafny.Map
                d_516_v27_ = _dafny.Map({(((d_511_v22_)[(self).f17] if ((self).f17) in (d_511_v22_) else default__.fm0(d_512_v23_, (self).f13, (self).f22, d_506_cf20_, globalState))) + (len(d_513_v24_)): (len((default__.fm34(d_506_cf20_, globalState)).set(default__.safeIndex((self).f21, len(default__.fm34(d_506_cf20_, globalState))), (self).f21))) - (len(default__.fm32(d_515_v26_, globalState)))})
                d_516_v27_ = (d_516_v27_).set(d_506_cf20_, len(d_507_cf19_))
                d_517_v28_: _dafny.Array
                nw59_ = _dafny.Array(D7.default()(), 12)
                d_517_v28_ = nw59_
                d_518_v29_: D7
                d_518_v29_ = D7_DC16(d_506_cf20_)
                d_519_v30_: D7
                d_519_v30_ = D7_DC17(d_518_v29_)
                d_520_v31_: D7
                d_520_v31_ = D7_DC17(d_519_v30_)
                index47_ = default__.safeIndex(529, (d_517_v28_).length(0))
                (d_517_v28_)[index47_] = d_520_v31_
                d_521_v32_: _dafny.Map
                d_521_v32_ = _dafny.Map({(self).f13: d_507_cf19_})
                d_522_v33_: _dafny.Seq
                d_522_v33_ = _dafny.SeqWithoutIsStrInference([d_458_v0_, d_458_v0_, d_507_cf19_, (d_458_v0_) + (d_507_cf19_), d_458_v0_])
                index48_ = default__.safeIndex(529, (d_517_v28_).length(0))
                rhs40_ = d_506_cf20_
                rhs41_ = len((d_522_v33_)[default__.safeIndex((d_513_v24_)[default__.safeIndex(391, len(d_513_v24_))], len(d_522_v33_))])
                rhs42_ = d_520_v31_
                rhs43_ = default__.fm35((d_463_v3_).cf20, (len(d_512_v23_) if d_508_cf18_ else d_506_cf20_), (d_499_v19_).f20, not(default__.fm30((0) - ((self).f21), globalState)), globalState)
                lhs33_ = globalState
                lhs34_ = globalState
                lhs35_ = d_517_v28_
                lhs36_ = default__.safeIndex(529, (d_517_v28_).length(0))
                lhs33_.f7 = rhs40_
                lhs34_.f7 = rhs41_
                lhs35_[lhs36_] = rhs42_
                d_521_v32_ = rhs43_
                d_515_v26_ = d_515_v26_
            elif source11_.is_DC13:
                d_523___mcc_h15_ = source11_.cf22
                d_524___mcc_h16_ = source11_.cf23
                d_525___mcc_h17_ = source11_.cf24
                d_526___mcc_h18_ = source11_.cf25
                d_527___mcc_h19_ = source11_.cf26
                d_528_cf26_ = d_527___mcc_h19_
                d_529_cf25_ = d_526___mcc_h18_
                d_530_cf24_ = d_525___mcc_h17_
                d_531_cf23_ = d_524___mcc_h16_
                d_532_cf22_ = d_523___mcc_h15_
                d_533_v34_: _dafny.Array
                def lambda12_(d_534_cf23_, d_535_cf25_):
                    def lambda13_(d_536_i3_):
                        return (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_534_cf23_: True}))])) + (_dafny.SeqWithoutIsStrInference([(self).f21, (self).f21, d_535_cf25_]))

                    return lambda13_

                init8_ = lambda12_(d_531_cf23_, d_529_cf25_)
                nw60_ = _dafny.Array(None, 14)
                for i0_8_ in range(nw60_.length(0)):
                    nw60_[i0_8_] = init8_(i0_8_)
                d_533_v34_ = nw60_
                d_537_v35_: _dafny.Seq
                d_537_v35_ = _dafny.SeqWithoutIsStrInference([(self).f21])
                index49_ = default__.safeIndex(7, (d_533_v34_).length(0))
                (d_533_v34_)[index49_] = d_537_v35_
                index50_ = default__.safeIndex(7, (d_533_v34_).length(0))
                (d_533_v34_)[index50_] = d_537_v35_
                d_538_v36_: _dafny.Map
                d_538_v36_ = _dafny.Map({d_529_cf25_: not((self).f17)})
                d_539_v37_: _dafny.Seq
                d_539_v37_ = _dafny.SeqWithoutIsStrInference([(self).f13, ((d_538_v36_)[d_530_cf24_] if (d_530_cf24_) in (d_538_v36_) else d_499_v19_.f19), d_528_cf26_, (self).f22, d_528_cf26_])
                (d_499_v19_).f19 = (len((self).fm10(d_531_cf23_, d_528_cf26_, d_539_v37_, globalState))) < (d_530_cf24_)
                d_531_cf23_ = ((self).f21) == (d_529_cf25_)
                d_540_v38_: _dafny.Array
                nw61_ = _dafny.Array(None, 29)
                nw61_[int(0)] = (self).f14
                nw61_[int(1)] = (self).f14
                nw61_[int(2)] = (self).f14
                nw61_[int(3)] = (self).f14
                nw61_[int(4)] = (self).f14
                nw61_[int(5)] = (self).f14
                nw61_[int(6)] = (self).f14
                nw61_[int(7)] = (self).f14
                nw61_[int(8)] = (self).f14
                nw61_[int(9)] = (self).f14
                nw61_[int(10)] = (self).f14
                nw61_[int(11)] = (self).f14
                nw61_[int(12)] = (D2_DC4(d_459_v1_, (self).f22, len(d_458_v0_), d_499_v19_.f19, (self).f14)).cf10
                nw61_[int(13)] = (self).f14
                nw61_[int(14)] = (self).f14
                nw61_[int(15)] = (self).f14
                nw61_[int(16)] = (self).f14
                nw61_[int(17)] = (self).f14
                nw61_[int(18)] = (self).f14
                nw61_[int(19)] = (self).f14
                nw61_[int(20)] = (self).f14
                nw61_[int(21)] = (self).f14
                nw61_[int(22)] = (self).f14
                nw61_[int(23)] = (self).f14
                nw61_[int(24)] = (self).f14
                nw61_[int(25)] = (self).f14
                nw61_[int(26)] = (self).f14
                nw61_[int(27)] = (self).f14
                nw61_[int(28)] = (self).f14
                d_540_v38_ = nw61_
                d_541_v39_: _dafny.Array
                d_541_v39_ = d_540_v38_
                d_542_v40_: _dafny.Array
                nw62_ = _dafny.Array(None, 27)
                nw62_[int(0)] = d_541_v39_
                nw62_[int(1)] = d_541_v39_
                nw62_[int(2)] = d_540_v38_
                nw62_[int(3)] = d_541_v39_
                nw62_[int(4)] = d_541_v39_
                nw62_[int(5)] = d_541_v39_
                nw62_[int(6)] = d_540_v38_
                nw62_[int(7)] = d_541_v39_
                nw62_[int(8)] = d_541_v39_
                nw62_[int(9)] = d_541_v39_
                nw62_[int(10)] = d_541_v39_
                nw62_[int(11)] = d_541_v39_
                nw62_[int(12)] = d_541_v39_
                nw62_[int(13)] = d_541_v39_
                nw62_[int(14)] = d_541_v39_
                nw62_[int(15)] = d_541_v39_
                nw62_[int(16)] = d_541_v39_
                nw62_[int(17)] = d_541_v39_
                nw62_[int(18)] = d_541_v39_
                nw62_[int(19)] = d_540_v38_
                nw62_[int(20)] = d_541_v39_
                nw62_[int(21)] = d_541_v39_
                nw62_[int(22)] = d_541_v39_
                nw62_[int(23)] = d_541_v39_
                nw62_[int(24)] = d_541_v39_
                nw62_[int(25)] = d_541_v39_
                nw62_[int(26)] = d_541_v39_
                d_542_v40_ = nw62_
                index51_ = default__.safeIndex(257, (d_542_v40_).length(0))
                (d_542_v40_)[index51_] = d_541_v39_
                index52_ = default__.safeIndex(257, (d_542_v40_).length(0))
                (d_542_v40_)[index52_] = d_540_v38_
            elif source11_.is_DC11:
                d_543___mcc_h20_ = source11_.cf17
                d_544_cf17_ = d_543___mcc_h20_
                index53_ = default__.safeIndex(312, ((self).f14).length(0))
                ((self).f14)[index53_] = (self).f22
                index54_ = default__.safeIndex(312, ((self).f14).length(0))
                ((self).f14)[index54_] = d_499_v19_.f19
                (globalState).f7 = (self).f21
                d_545_v41_: _dafny.MultiSet
                d_545_v41_ = _dafny.MultiSet([(d_458_v0_) + (d_458_v0_), d_458_v0_, d_458_v0_, d_458_v0_, d_458_v0_])
                d_546_v42_: _dafny.Map
                d_546_v42_ = _dafny.Map({(658) <= ((self).f21): d_459_v1_})
                d_547_v43_: _dafny.Set
                d_547_v43_ = _dafny.Set({(self).f17})
                d_548_v44_: _dafny.Seq
                d_548_v44_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({False}), d_547_v43_])
                d_549_v45_: _dafny.MultiSet
                d_549_v45_ = _dafny.MultiSet([((d_548_v44_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(self).f21])), len(d_548_v44_))]).isdisjoint(d_547_v43_), d_499_v19_.f19, ((self).f15).issubset((self).f15), d_499_v19_.f19, default__.fm30((self).f21, globalState)])
                d_550_v46_: _dafny.Seq
                d_550_v46_ = _dafny.SeqWithoutIsStrInference([d_549_v45_, d_549_v45_, (d_549_v45_) - (d_549_v45_)])
                index55_ = default__.safeIndex(312, ((self).f14).length(0))
                def iife57_():
                    coll33_ = _dafny.Map()
                    compr_33_: int
                    for compr_33_ in _dafny.IntegerRange(391, 894):
                        d_551_v47_: int = compr_33_
                        if ((391) <= (d_551_v47_)) and ((d_551_v47_) < (894)):
                            coll33_[(d_551_v47_) * ((self).f21)] = d_549_v45_
                    return _dafny.Map(coll33_)
                rhs44_ = False
                rhs45_ = ((d_545_v41_ if d_499_v19_.f19 else (d_545_v41_).set(d_458_v0_, default__.abs((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f21]))).cardinality)))).set((self.f16).cf1, default__.abs((self).f21))
                rhs46_ = d_546_v42_
                rhs47_ = (d_550_v46_)[default__.safeIndex((self).f21, len(d_550_v46_))]
                rhs48_ = not ((iife57_()
                ) != (_dafny.Map({(self).f21: d_549_v45_}))) or ((self).f13)
                lhs37_ = (self).f14
                lhs38_ = default__.safeIndex(312, ((self).f14).length(0))
                lhs39_ = d_499_v19_
                lhs37_[lhs38_] = rhs44_
                d_545_v41_ = rhs45_
                d_546_v42_ = rhs46_
                d_549_v45_ = rhs47_
                lhs39_.f19 = rhs48_
                d_552_v48_: _dafny.Seq
                d_552_v48_ = _dafny.SeqWithoutIsStrInference([(self).f17, (self).f13])
                index56_ = default__.safeIndex(312, ((self).f14).length(0))
                ((self).f14)[index56_] = ((default__.fm36((d_552_v48_)[default__.safeIndex(117, len(d_552_v48_))], (d_499_v19_).f20, default__.fm30((self).f21, globalState), globalState)) == (d_547_v43_)) and (True)
            elif True:
                d_553___mcc_h21_ = source11_.cf27
                d_554_cf27_ = d_553___mcc_h21_
                d_555_v49_: _dafny.Array
                nw63_ = _dafny.Array(_dafny.MultiSet({}), 20)
                d_555_v49_ = nw63_
                rhs49_ = d_555_v49_
                rhs50_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mxwbndl"))
                rhs51_ = d_499_v19_.f19
                lhs40_ = d_499_v19_
                d_555_v49_ = rhs49_
                d_458_v0_ = rhs50_
                lhs40_.f19 = rhs51_
                d_556_v50_: _dafny.Seq
                d_556_v50_ = _dafny.SeqWithoutIsStrInference([(self).f21])
                d_557_v51_: _dafny.Map
                d_557_v51_ = _dafny.Map({d_458_v0_: len((d_556_v50_) + (d_556_v50_))})
                d_557_v51_ = (d_557_v51_).set(d_458_v0_, (self).f21)
                (d_499_v19_).f19 = d_499_v19_.f19
                d_558_v52_: T0
                nw64_ = C0()
                nw64_.ctor__((self).f22, (d_499_v19_).f20, default__.fm37((self).f21, (self).f21, (self).f21, (d_499_v19_).fm16((self).f22, (self).f21, globalState), globalState), d_499_v19_.f19)
                d_558_v52_ = nw64_
                d_559_v53_: _dafny.MultiSet
                d_559_v53_ = _dafny.MultiSet([d_558_v52_, d_558_v52_, d_558_v52_])
                d_559_v53_ = (d_559_v53_).intersection(d_559_v53_)
            d_560_v54_: _dafny.Array
            def lambda14_(d_561_v18_):
                def lambda15_(d_562_i4_):
                    return d_561_v18_

                return lambda15_

            init9_ = lambda14_(d_498_v18_)
            nw65_ = _dafny.Array(None, 14)
            for i0_9_ in range(nw65_.length(0)):
                nw65_[i0_9_] = init9_(i0_9_)
            d_560_v54_ = nw65_
            d_563_v55_: _dafny.MultiSet
            d_563_v55_ = _dafny.MultiSet([d_560_v54_])
            d_564_v56_: _dafny.Map
            d_564_v56_ = _dafny.Map({(self).f14: d_563_v55_})
            d_565_v57_: _dafny.MultiSet
            d_565_v57_ = _dafny.MultiSet([(self).f13, not((self).f13)])
            if ((d_560_v54_) in (((d_564_v56_)[(self).f14] if ((self).f14) in (d_564_v56_) else d_563_v55_))) not in (d_565_v57_):
                d_566_v58_: C0
                nw66_ = C0()
                nw66_.ctor__(not (d_499_v19_.f19) or ((self).f13), (d_499_v19_).f20, D1_DC2((self).f21, 609, default__.fm30((self).f21, globalState)), d_499_v19_.f19)
                d_566_v58_ = nw66_
                d_498_v18_ = (d_499_v19_).f20
                d_458_v0_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "avvykq"))) + ((d_458_v0_) + (d_458_v0_))
                d_567_v59_: _dafny.Seq
                d_567_v59_ = _dafny.SeqWithoutIsStrInference([(self).f21])
                d_568_v60_: _dafny.Map
                d_568_v60_ = _dafny.Map({d_458_v0_: d_567_v59_})
                d_569_v62_: _dafny.Map
                def iife58_():
                    coll34_ = _dafny.Set()
                    compr_34_: _dafny.Seq
                    for compr_34_ in (d_568_v60_).keys.Elements:
                        d_570_v61_: _dafny.Seq = compr_34_
                        if (d_570_v61_) in (d_568_v60_):
                            coll34_ = coll34_.union(_dafny.Set([d_570_v61_]))
                    return _dafny.Set(coll34_)
                d_569_v62_ = _dafny.Map({(self).f21: iife58_()
                })
                d_571_v63_: _dafny.Set
                d_571_v63_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mbugy")), d_458_v0_, d_458_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fw")), d_458_v0_})
                d_569_v62_ = (d_569_v62_).set((self).f21, d_571_v63_)
                rhs52_ = (self).f21
                rhs53_ = len((self).f15)
                lhs41_ = globalState
                lhs42_ = globalState
                lhs41_.f7 = rhs52_
                lhs42_.f7 = rhs53_
            elif True:
                d_572_v64_: _dafny.Map
                d_572_v64_ = _dafny.Map({self.f16: (self).f21})
                d_573_v65_: _dafny.Map
                d_573_v65_ = _dafny.Map({d_499_v19_.f19: d_572_v64_})
                d_573_v65_ = d_573_v65_
                index57_ = default__.safeIndex(242, (d_459_v1_).length(0))
                (d_459_v1_)[index57_] = (self).f21
                index58_ = default__.safeIndex(577, (d_459_v1_).length(0))
                (d_459_v1_)[index58_] = (self).f21
                index59_ = default__.safeIndex(242, (d_459_v1_).length(0))
                index60_ = default__.safeIndex(577, (d_459_v1_).length(0))
                rhs54_ = (self).f21
                rhs55_ = (((self).f21) - ((self).f21)) + ((0) - ((self).f21))
                lhs43_ = d_459_v1_
                lhs44_ = default__.safeIndex(242, (d_459_v1_).length(0))
                lhs45_ = d_459_v1_
                lhs46_ = default__.safeIndex(577, (d_459_v1_).length(0))
                lhs43_[lhs44_] = rhs54_
                lhs45_[lhs46_] = rhs55_
                d_574_v66_: _dafny.Seq
                d_574_v66_ = _dafny.SeqWithoutIsStrInference([(self).f17])
                d_575_v67_: _dafny.Set
                d_575_v67_ = _dafny.Set({(d_574_v66_)[default__.safeIndex((self).f21, len(d_574_v66_))], True})
                d_575_v67_ = (_dafny.Set({d_499_v19_.f19})).intersection(d_575_v67_)
                d_576_v68_: _dafny.Map
                d_576_v68_ = _dafny.Map({(self).f13: (self).f15})
                d_577_v69_: _dafny.Array
                nw67_ = _dafny.Array(None, 24)
                nw67_[int(0)] = len((d_458_v0_).set(default__.safeIndex((d_459_v1_)[default__.safeIndex(242, (d_459_v1_).length(0))], len(d_458_v0_)), (d_499_v19_).f20))
                nw67_[int(1)] = (self).f21
                nw67_[int(2)] = (self).f21
                nw67_[int(3)] = len(d_496_v16_)
                nw67_[int(4)] = (0) - ((self).f21)
                nw67_[int(5)] = (d_459_v1_)[default__.safeIndex(242, (d_459_v1_).length(0))]
                nw67_[int(6)] = 256
                nw67_[int(7)] = 565
                nw67_[int(8)] = (d_459_v1_)[default__.safeIndex(242, (d_459_v1_).length(0))]
                nw67_[int(9)] = 780
                nw67_[int(10)] = (d_459_v1_)[default__.safeIndex(242, (d_459_v1_).length(0))]
                nw67_[int(11)] = (self).f21
                nw67_[int(12)] = default__.fm0(d_574_v66_, d_499_v19_.f19, (self).f22, (self).fm8(default__.fm30(len(d_574_v66_), globalState), self.f16, globalState), globalState)
                nw67_[int(13)] = len(d_576_v68_)
                nw67_[int(14)] = ((d_496_v16_)[False] if (False) in (d_496_v16_) else (d_459_v1_)[default__.safeIndex(242, (d_459_v1_).length(0))])
                nw67_[int(15)] = (self).f21
                nw67_[int(16)] = (self).f21
                nw67_[int(17)] = (self).f21
                nw67_[int(18)] = (d_459_v1_)[default__.safeIndex(242, (d_459_v1_).length(0))]
                nw67_[int(19)] = (d_459_v1_)[default__.safeIndex(242, (d_459_v1_).length(0))]
                nw67_[int(20)] = (d_459_v1_)[default__.safeIndex(242, (d_459_v1_).length(0))]
                nw67_[int(21)] = (d_459_v1_)[default__.safeIndex(242, (d_459_v1_).length(0))]
                nw67_[int(22)] = (d_459_v1_)[default__.safeIndex(242, (d_459_v1_).length(0))]
                nw67_[int(23)] = (d_459_v1_)[default__.safeIndex(242, (d_459_v1_).length(0))]
                d_577_v69_ = nw67_
                d_578_v70_: D2
                d_578_v70_ = D2_DC4(d_577_v69_, ((self).f21) == (len(d_575_v67_)), (d_459_v1_)[default__.safeIndex(242, (d_459_v1_).length(0))], default__.fm30((self).f21, globalState), (self).f14)
                d_579_v71_: _dafny.Seq
                d_579_v71_ = _dafny.SeqWithoutIsStrInference([(d_459_v1_)[default__.safeIndex(242, (d_459_v1_).length(0))]])
                d_580_v72_: _dafny.Map
                d_580_v72_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "txsjiq")): (self).f22})
                d_581_v73_: _dafny.Seq
                d_581_v73_ = _dafny.SeqWithoutIsStrInference([d_580_v72_])
                d_582_v74_: _dafny.Seq
                d_582_v74_ = _dafny.SeqWithoutIsStrInference([d_580_v72_, (d_581_v73_)[default__.safeIndex((self).f21, len(d_581_v73_))], d_580_v72_])
                d_583_v75_: _dafny.Map
                d_583_v75_ = _dafny.Map({d_577_v69_: (d_495_cf17_).isdisjoint((_dafny.MultiSet(d_579_v71_)).set((d_459_v1_)[default__.safeIndex(242, (d_459_v1_).length(0))], default__.abs(len((d_582_v74_)[default__.safeIndex((d_459_v1_)[default__.safeIndex(242, (d_459_v1_).length(0))], len(d_582_v74_))]))))})
                rhs56_ = D2_DC4(d_577_v69_, (self).f22, (d_459_v1_)[default__.safeIndex(242, (d_459_v1_).length(0))], ((d_574_v66_)[default__.safeIndex((d_459_v1_)[default__.safeIndex(242, (d_459_v1_).length(0))], len(d_574_v66_))] if (self).f22 else d_499_v19_.f19), (self).f14)
                rhs57_ = (d_583_v75_ if (default__.fm38(False, d_499_v19_.f19, d_499_v19_.f19, (self).f22, globalState)) == (d_574_v66_) else (d_583_v75_) | (d_583_v75_))
                d_578_v70_ = rhs56_
                d_583_v75_ = rhs57_
                (globalState).f2 = default__.fm38((self).f22, (self).f17, False, (self).f22, globalState)
        elif True:
            d_584___mcc_h10_ = source10_.cf27
            d_585_cf27_ = d_584___mcc_h10_
            d_586_v76_: _dafny.Seq
            d_586_v76_ = _dafny.SeqWithoutIsStrInference([self.f12])
            d_586_v76_ = d_586_v76_
            d_587_v77_: _dafny.Array
            nw68_ = _dafny.Array(_dafny.Array(None, 0), 7)
            d_587_v77_ = nw68_
            index61_ = default__.safeIndex(98, (d_587_v77_).length(0))
            (d_587_v77_)[index61_] = d_459_v1_
            index62_ = default__.safeIndex(98, (d_587_v77_).length(0))
            (d_587_v77_)[index62_] = d_459_v1_
            index63_ = default__.safeIndex(801, ((self).f14).length(0))
            ((self).f14)[index63_] = (self).f22
            index64_ = default__.safeIndex(801, ((self).f14).length(0))
            ((self).f14)[index64_] = (self).f17
            if not(not((self).f22)):
                d_588_v78_: _dafny.Array
                nw69_ = _dafny.Array(None, 11)
                d_588_v78_ = nw69_
                d_589_v79_: str
                d_589_v79_ = _dafny.CodePoint('b')
                d_590_v80_: C0
                nw70_ = C0()
                nw70_.ctor__(False, d_589_v79_, self.f12, (self).f22)
                d_590_v80_ = nw70_
                index65_ = default__.safeIndex(657, (d_588_v78_).length(0))
                (d_588_v78_)[index65_] = d_590_v80_
                index66_ = default__.safeIndex(801, ((self).f14).length(0))
                index67_ = default__.safeIndex(657, (d_588_v78_).length(0))
                rhs58_ = ((self).f14)[default__.safeIndex(801, ((self).f14).length(0))]
                rhs59_ = d_590_v80_
                lhs47_ = (self).f14
                lhs48_ = default__.safeIndex(801, ((self).f14).length(0))
                lhs49_ = d_588_v78_
                lhs50_ = default__.safeIndex(657, (d_588_v78_).length(0))
                lhs47_[lhs48_] = rhs58_
                lhs49_[lhs50_] = rhs59_
                (d_590_v80_).f19 = ((d_590_v80_).f20) in (_dafny.SeqWithoutIsStrInference([(d_590_v80_).f20 for d_591_i5_ in range(default__.abs(329))]))
                d_592_v81_: _dafny.Set
                d_592_v81_ = _dafny.Set({((self).f14)[default__.safeIndex(801, ((self).f14).length(0))], (self).f13})
                (d_590_v80_).f19 = ((self).f21) > ((len(d_592_v81_) if ((self).f14)[default__.safeIndex(801, ((self).f14).length(0))] else len(d_458_v0_)))
                d_593_v82_: _dafny.Seq
                d_593_v82_ = _dafny.SeqWithoutIsStrInference([len(d_458_v0_)])
                d_594_v83_: _dafny.Map
                d_594_v83_ = _dafny.Map({d_593_v82_: d_590_v80_.f19})
                d_595_v84_: _dafny.Seq
                d_595_v84_ = _dafny.SeqWithoutIsStrInference([d_593_v82_])
                d_594_v83_ = (d_594_v83_).set((d_595_v84_)[default__.safeIndex((self).f21, len(d_595_v84_))], (self).f22)
                d_596_v85_: _dafny.Array
                nw71_ = _dafny.Array(_dafny.CodePoint('D'), 10)
                d_596_v85_ = nw71_
                index68_ = default__.safeIndex(790, (d_596_v85_).length(0))
                (d_596_v85_)[index68_] = (d_590_v80_).f20
                d_597_v86_: _dafny.Set
                d_597_v86_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_589_v79_ for d_598_i6_ in range(default__.abs(846))])})
                index69_ = default__.safeIndex(790, (d_596_v85_).length(0))
                index70_ = default__.safeIndex(801, ((self).f14).length(0))
                rhs60_ = (self).fm9(_dafny.CodePoint('y'), globalState)
                rhs61_ = d_589_v79_
                rhs62_ = False
                rhs63_ = (d_597_v86_).issubset(_dafny.Set({d_458_v0_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_599_i7_ in range(default__.abs(726))])}))
                rhs64_ = not((self).f22)
                lhs51_ = globalState
                lhs52_ = d_596_v85_
                lhs53_ = default__.safeIndex(790, (d_596_v85_).length(0))
                lhs54_ = (self).f14
                lhs55_ = default__.safeIndex(801, ((self).f14).length(0))
                lhs56_ = d_590_v80_
                lhs57_ = d_590_v80_
                lhs51_.f7 = rhs60_
                lhs52_[lhs53_] = rhs61_
                lhs54_[lhs55_] = rhs62_
                lhs56_.f19 = rhs63_
                lhs57_.f19 = rhs64_
            elif True:
                d_600_v87_: _dafny.Seq
                d_600_v87_ = _dafny.SeqWithoutIsStrInference([(self).f21])
                d_601_v88_: D7
                d_601_v88_ = D7_DC15(d_600_v87_)
                pat_let_tv15_ = globalState
                def iife59_(_pat_let12_0):
                    def iife60_(d_602_dt__update__tmp_h3_):
                        def iife61_(_pat_let13_0):
                            def iife62_(d_603_dt__update_hcf28_h0_):
                                return D7_DC15(d_603_dt__update_hcf28_h0_)
                            return iife62_(_pat_let13_0)
                        return iife61_(default__.fm34((self).f21, pat_let_tv15_))
                    return iife60_(_pat_let12_0)
                d_601_v88_ = (iife59_(d_601_v88_) if ((self).f14)[default__.safeIndex(801, ((self).f14).length(0))] else d_601_v88_)
                d_604_v89_: _dafny.Map
                d_604_v89_ = _dafny.Map({(self).f17: (self).f13})
                d_605_v90_: D6
                d_605_v90_ = D6_DC13((self).f17, (self).f22, len(d_458_v0_), (self).f21, ((d_604_v89_)[(self).f13] if ((self).f13) in (d_604_v89_) else (self).f17))
                index71_ = default__.safeIndex(801, ((self).f14).length(0))
                ((self).f14)[index71_] = (d_605_v90_).cf26
                d_606_v91_: _dafny.Array
                nw72_ = _dafny.Array(_dafny.Seq({}), 5)
                d_606_v91_ = nw72_
                d_607_v92_: _dafny.Seq
                d_607_v92_ = _dafny.SeqWithoutIsStrInference([True, (self).f17, True, ((self).f14)[default__.safeIndex(801, ((self).f14).length(0))], ((self).f14)[default__.safeIndex(801, ((self).f14).length(0))]])
                index72_ = default__.safeIndex(422, (d_606_v91_).length(0))
                (d_606_v91_)[index72_] = (d_607_v92_) + (d_607_v92_)
                index73_ = default__.safeIndex(422, (d_606_v91_).length(0))
                index74_ = default__.safeIndex(801, ((self).f14).length(0))
                index75_ = default__.safeIndex(801, ((self).f14).length(0))
                rhs65_ = d_607_v92_
                rhs66_ = ((self).f14)[default__.safeIndex(801, ((self).f14).length(0))]
                rhs67_ = (d_458_v0_ if default__.fm30((self).f21, globalState) else d_458_v0_)
                rhs68_ = not(not((self).f13))
                lhs58_ = d_606_v91_
                lhs59_ = default__.safeIndex(422, (d_606_v91_).length(0))
                lhs60_ = (self).f14
                lhs61_ = default__.safeIndex(801, ((self).f14).length(0))
                lhs62_ = (self).f14
                lhs63_ = default__.safeIndex(801, ((self).f14).length(0))
                lhs58_[lhs59_] = rhs65_
                lhs60_[lhs61_] = rhs66_
                d_458_v0_ = rhs67_
                lhs62_[lhs63_] = rhs68_
                d_608_v93_: _dafny.Map
                d_608_v93_ = _dafny.Map({((self).f14)[default__.safeIndex(801, ((self).f14).length(0))]: (self).f21})
                arr6_ = (d_587_v77_)[default__.safeIndex(98, (d_587_v77_).length(0))]
                index76_ = default__.safeIndex(181, ((d_587_v77_)[default__.safeIndex(98, (d_587_v77_).length(0))]).length(0))
                arr6_[index76_] = len(d_608_v93_)
                index77_ = default__.safeIndex(799, (d_459_v1_).length(0))
                (d_459_v1_)[index77_] = (d_600_v87_)[default__.safeIndex(226, len(d_600_v87_))]
                arr7_ = (d_587_v77_)[default__.safeIndex(98, (d_587_v77_).length(0))]
                index78_ = default__.safeIndex(109, ((d_587_v77_)[default__.safeIndex(98, (d_587_v77_).length(0))]).length(0))
                arr7_[index78_] = len(_dafny.Map({(self).f21: (self).f17}))
                d_609_v94_: str
                d_609_v94_ = _dafny.CodePoint('x')
                d_610_v95_: D6
                d_610_v95_ = D6_DC14(default__.fm39(False, ((self).f14)[default__.safeIndex(801, ((self).f14).length(0))], globalState))
                d_611_v96_: _dafny.Array
                nw73_ = _dafny.Array(_dafny.CodePoint('D'), 22)
                d_611_v96_ = nw73_
                d_612_v97_: _dafny.Seq
                d_612_v97_ = _dafny.SeqWithoutIsStrInference([d_611_v96_, d_611_v96_, d_611_v96_])
                arr8_ = (d_587_v77_)[default__.safeIndex(98, (d_587_v77_).length(0))]
                index79_ = default__.safeIndex(181, ((d_587_v77_)[default__.safeIndex(98, (d_587_v77_).length(0))]).length(0))
                index80_ = default__.safeIndex(801, ((self).f14).length(0))
                index81_ = default__.safeIndex(799, (d_459_v1_).length(0))
                arr9_ = (d_587_v77_)[default__.safeIndex(98, (d_587_v77_).length(0))]
                index82_ = default__.safeIndex(109, ((d_587_v77_)[default__.safeIndex(98, (d_587_v77_).length(0))]).length(0))
                rhs69_ = (self).f21
                rhs70_ = (d_458_v0_) in (_dafny.SeqWithoutIsStrInference([(d_458_v0_).set(default__.safeIndex((self).f21, len(d_458_v0_)), d_609_v94_), d_458_v0_, default__.fm32(d_610_v95_, globalState)]))
                rhs71_ = (len(d_600_v87_)) + (((self).f21) + ((self).f21))
                rhs72_ = len(d_612_v97_)
                lhs64_ = (d_587_v77_)[default__.safeIndex(98, (d_587_v77_).length(0))]
                lhs65_ = default__.safeIndex(181, ((d_587_v77_)[default__.safeIndex(98, (d_587_v77_).length(0))]).length(0))
                lhs66_ = (self).f14
                lhs67_ = default__.safeIndex(801, ((self).f14).length(0))
                lhs68_ = d_459_v1_
                lhs69_ = default__.safeIndex(799, (d_459_v1_).length(0))
                lhs70_ = (d_587_v77_)[default__.safeIndex(98, (d_587_v77_).length(0))]
                lhs71_ = default__.safeIndex(109, ((d_587_v77_)[default__.safeIndex(98, (d_587_v77_).length(0))]).length(0))
                lhs64_[lhs65_] = rhs69_
                lhs66_[lhs67_] = rhs70_
                lhs68_[lhs69_] = rhs71_
                lhs70_[lhs71_] = rhs72_
                index83_ = default__.safeIndex(799, (d_459_v1_).length(0))
                (d_459_v1_)[index83_] = default__.safeDivisionInt(-317, ((d_587_v77_)[default__.safeIndex(98, (d_587_v77_).length(0))])[default__.safeIndex(181, ((d_587_v77_)[default__.safeIndex(98, (d_587_v77_).length(0))]).length(0))])
        d_613_v98_: bool
        d_613_v98_ = False
        d_614_v99_: _dafny.Seq
        d_614_v99_ = _dafny.SeqWithoutIsStrInference([False, (self).f13, d_613_v98_, d_613_v98_, (self).f13])
        d_615_v100_: _dafny.Map
        d_615_v100_ = _dafny.Map({(self).f21: d_614_v99_})
        d_616_v101_: _dafny.Map
        d_616_v101_ = _dafny.Map({len(d_614_v99_): (self).f22})
        d_613_v98_ = ((((d_615_v100_)[len(d_616_v101_)] if (len(d_616_v101_)) in (d_615_v100_) else d_614_v99_)).set(default__.safeIndex((self).f21, len(((d_615_v100_)[len(d_616_v101_)] if (len(d_616_v101_)) in (d_615_v100_) else d_614_v99_))), (self).f22)) == ((d_614_v99_) + (d_614_v99_))
        d_617_v102_: _dafny.MultiSet
        d_617_v102_ = _dafny.MultiSet([(self).f17])
        d_618_v103_: _dafny.MultiSet
        d_618_v103_ = d_617_v102_
        d_619_v104_: _dafny.Array
        nw74_ = _dafny.Array(None, 16)
        nw74_[int(0)] = d_617_v102_
        nw74_[int(1)] = d_617_v102_
        nw74_[int(2)] = (d_617_v102_) - ((_dafny.MultiSet([(self).f22, True])).set(True, default__.abs(len(_dafny.SeqWithoutIsStrInference([(self).f21 for d_620_i9_ in range(default__.abs(124))])))))
        nw74_[int(3)] = d_617_v102_
        nw74_[int(4)] = (d_617_v102_).intersection(d_617_v102_)
        nw74_[int(5)] = (d_618_v103_)
        nw74_[int(6)] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f13]))
        nw74_[int(7)] = d_617_v102_
        nw74_[int(8)] = _dafny.MultiSet([(self).f22])
        nw74_[int(9)] = (default__.fm40(globalState)).intersection(d_617_v102_)
        nw74_[int(10)] = d_617_v102_
        nw74_[int(11)] = (d_617_v102_) - (d_617_v102_)
        nw74_[int(12)] = d_617_v102_
        nw74_[int(13)] = _dafny.MultiSet(d_614_v99_)
        nw74_[int(14)] = (_dafny.MultiSet([(self).f22])) | (d_617_v102_)
        nw74_[int(15)] = (d_617_v102_) - (_dafny.MultiSet([(self).f13]))
        d_619_v104_ = nw74_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_619_v104_).length(0)):
            d_621_i8_: int = guard_loop_3_
            if (True) and (((0) <= (d_621_i8_)) and ((d_621_i8_) < ((d_619_v104_).length(0)))):
                (d_619_v104_)[(d_621_i8_)] = (d_617_v102_) | (d_617_v102_)
        d_622_v105_: _dafny.Array
        def lambda16_(d_623_i10_):
            return (self.f16).cf1

        init10_ = lambda16_
        nw75_ = _dafny.Array(None, 29)
        for i0_10_ in range(nw75_.length(0)):
            nw75_[i0_10_] = init10_(i0_10_)
        d_622_v105_ = nw75_
        index84_ = default__.safeIndex(719, (d_622_v105_).length(0))
        (d_622_v105_)[index84_] = (d_458_v0_ if d_613_v98_ else d_458_v0_)
        d_624_v106_: D6
        d_624_v106_ = D6_DC14(D6_DC13((self).f17, (self).f13, (self).f21, (self).f21, d_613_v98_))
        d_625_v107_: D6
        d_625_v107_ = D6_DC14(d_624_v106_)
        d_626_v108_: D6
        d_626_v108_ = D6_DC14(d_624_v106_)
        pat_let_tv16_ = d_458_v0_
        pat_let_tv17_ = d_458_v0_
        index85_ = default__.safeIndex(719, (d_622_v105_).length(0))
        def lambda17_(source12_):
            if source12_.is_DC12:
                d_627___mcc_h22_ = source12_.cf18
                d_628___mcc_h23_ = source12_.cf19
                d_629___mcc_h24_ = source12_.cf20
                d_630___mcc_h25_ = source12_.cf21
                d_631_cf21_ = d_630___mcc_h25_
                d_632_cf20_ = d_629___mcc_h24_
                d_633_cf19_ = d_628___mcc_h23_
                d_634_cf18_ = d_627___mcc_h22_
                return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_635_i11_ in range(default__.abs(86))])
            elif source12_.is_DC13:
                d_636___mcc_h26_ = source12_.cf22
                d_637___mcc_h27_ = source12_.cf23
                d_638___mcc_h28_ = source12_.cf24
                d_639___mcc_h29_ = source12_.cf25
                d_640___mcc_h30_ = source12_.cf26
                d_641_cf26_ = d_640___mcc_h30_
                d_642_cf25_ = d_639___mcc_h29_
                d_643_cf24_ = d_638___mcc_h28_
                d_644_cf23_ = d_637___mcc_h27_
                d_645_cf22_ = d_636___mcc_h26_
                return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_646_i12_ in range(default__.abs(-521))])
            elif source12_.is_DC11:
                d_647___mcc_h31_ = source12_.cf17
                d_648_cf17_ = d_647___mcc_h31_
                return pat_let_tv16_
            elif True:
                d_649___mcc_h32_ = source12_.cf27
                d_650_cf27_ = d_649___mcc_h32_
                return pat_let_tv17_

        (d_622_v105_)[index85_] = lambda17_(d_626_v108_)
        d_651_v109_: _dafny.Array
        nw76_ = _dafny.Array(_dafny.Map({}), 24)
        d_651_v109_ = nw76_
        d_652_v110_: str
        d_652_v110_ = _dafny.CodePoint('u')
        d_653_v111_: _dafny.Map
        d_653_v111_ = _dafny.Map({(self).f22: (self).f15})
        index86_ = default__.safeIndex(24, (d_651_v109_).length(0))
        (d_651_v109_)[index86_] = (default__.fm41((self).f21, d_652_v110_, globalState)) | (d_653_v111_)
        d_654_v112_: _dafny.Seq
        d_654_v112_ = _dafny.SeqWithoutIsStrInference([d_458_v0_, d_458_v0_])
        d_655_v113_: _dafny.Set
        d_655_v113_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_652_v110_ for d_656_i13_ in range(default__.abs(924))]), (d_654_v112_)[default__.safeIndex((self).f21, len(d_654_v112_))]})
        d_657_v114_: _dafny.Map
        d_657_v114_ = _dafny.Map({(self).f13: True})
        d_658_v115_: D3
        d_658_v115_ = D3_DC5(default__.fm33((self).f22, len(d_657_v114_), globalState))
        pat_let_tv18_ = d_653_v111_
        pat_let_tv19_ = d_653_v111_
        pat_let_tv20_ = d_653_v111_
        pat_let_tv21_ = d_653_v111_
        pat_let_tv22_ = d_653_v111_
        pat_let_tv23_ = d_653_v111_
        index87_ = default__.safeIndex(24, (d_651_v109_).length(0))
        def lambda18_(source13_):
            d_659___mcc_h33_ = source13_
            d_660_cf16_ = d_659___mcc_h33_
            return (self).f22

        def lambda19_(source14_):
            if source14_.is_DC6:
                d_661___mcc_h34_ = source14_.cf12
                d_662___mcc_h35_ = source14_.cf13
                d_663_cf13_ = d_662___mcc_h35_
                d_664_cf12_ = d_661___mcc_h34_
                return (pat_let_tv18_) | (pat_let_tv19_)
            elif source14_.is_DC7:
                return pat_let_tv20_
            elif source14_.is_DC8:
                d_665___mcc_h36_ = source14_.cf14
                d_666_cf14_ = d_665___mcc_h36_
                return pat_let_tv21_
            elif True:
                d_667___mcc_h37_ = source14_.cf11
                d_668_cf11_ = d_667___mcc_h37_
                if (self).f17:
                    return pat_let_tv22_
                elif True:
                    return pat_let_tv23_

        rhs73_ = lambda18_(d_655_v113_)
        rhs74_ = lambda19_(d_658_v115_)
        rhs75_ = (self).f21
        rhs76_ = (self).f21
        lhs72_ = d_651_v109_
        lhs73_ = default__.safeIndex(24, (d_651_v109_).length(0))
        lhs74_ = globalState
        lhs75_ = globalState
        d_613_v98_ = rhs73_
        lhs72_[lhs73_] = rhs74_
        lhs74_.f7 = rhs75_
        lhs75_.f7 = rhs76_
        d_669_i14_: int
        d_669_i14_ = 0
        with _dafny.label("3"):
            while (self).f17:
                with _dafny.c_label("3"):
                    if (d_669_i14_) >= (100):
                        raise _dafny.Break("3")
                    d_669_i14_ = (d_669_i14_) + (1)
                    d_670_v116_: _dafny.Map
                    d_670_v116_ = _dafny.Map({((self).f22 if (self).f13 else not(not(((d_616_v101_)[(self).f21] if ((self).f21) in (d_616_v101_) else d_613_v98_)))): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gr"))})
                    d_670_v116_ = (d_670_v116_).set((d_614_v99_)[default__.safeIndex((self).f21, len(d_614_v99_))], ((d_622_v105_)[default__.safeIndex(719, (d_622_v105_).length(0))]) + ((d_622_v105_)[default__.safeIndex(719, (d_622_v105_).length(0))]))
                    index88_ = default__.safeIndex(607, (d_459_v1_).length(0))
                    (d_459_v1_)[index88_] = (self).f21
                    index89_ = default__.safeIndex(607, (d_459_v1_).length(0))
                    (d_459_v1_)[index89_] = (self).f21
                    d_613_v98_ = (((d_459_v1_)[default__.safeIndex(607, (d_459_v1_).length(0))] if d_613_v98_ else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "av"))))) != (540)
                    if (self).f17:
                        pat_let_tv24_ = d_652_v110_
                        pat_let_tv25_ = d_652_v110_
                        def iife64_(_pat_let15_0):
                            def iife65_(d_671_dt__update__tmp_h4_):
                                def iife66_(_pat_let16_0):
                                    def iife67_(d_672_dt__update_hcf11_h0_):
                                        return D3_DC5(d_672_dt__update_hcf11_h0_)
                                    return iife67_(_pat_let16_0)
                                return iife66_(pat_let_tv24_)
                            return iife65_(_pat_let15_0)
                        def iife63_(_pat_let14_0):
                            def iife68_(d_673_dt__update__tmp_h5_):
                                def iife69_(_pat_let17_0):
                                    def iife70_(d_674_dt__update_hcf11_h1_):
                                        return D3_DC5(d_674_dt__update_hcf11_h1_)
                                    return iife70_(_pat_let17_0)
                                return iife69_(pat_let_tv25_)
                            return iife68_(_pat_let14_0)
                        d_658_v115_ = iife63_(iife64_(d_658_v115_))
                        index90_ = default__.safeIndex(879, ((self).f14).length(0))
                        ((self).f14)[index90_] = False
                        index91_ = default__.safeIndex(879, ((self).f14).length(0))
                        ((self).f14)[index91_] = not((self).f13)
                        index92_ = default__.safeIndex(879, ((self).f14).length(0))
                        ((self).f14)[index92_] = (self).f17
                        d_675_v117_: _dafny.Seq
                        d_675_v117_ = _dafny.SeqWithoutIsStrInference([(self).f21])
                        d_676_v118_: _dafny.Map
                        d_676_v118_ = _dafny.Map({len(d_675_v117_): d_616_v101_})
                        d_676_v118_ = (d_676_v118_).set((self).f21, (d_616_v101_) | (d_616_v101_))
                        d_613_v98_ = not((self).f13)
                    elif True:
                        index93_ = default__.safeIndex(607, (d_459_v1_).length(0))
                        (d_459_v1_)[index93_] = (self).f21
                        d_616_v101_ = (d_616_v101_).set((d_459_v1_)[default__.safeIndex(607, (d_459_v1_).length(0))], (self).f22)
                        d_677_v119_: C0
                        nw77_ = C0()
                        nw77_.ctor__(not((self).f13), d_652_v110_, self.f12, (not(True) if not((self).f22) else (self).f22))
                        d_677_v119_ = nw77_
                        d_613_v98_ = (self).f17
                        d_678_v120_: _dafny.MultiSet
                        d_678_v120_ = _dafny.MultiSet([449, (d_459_v1_)[default__.safeIndex(607, (d_459_v1_).length(0))], len((self).f15), (self).f21])
                        d_679_v121_: _dafny.Array
                        nw78_ = _dafny.Array(None, 21)
                        nw78_[int(0)] = d_652_v110_
                        nw78_[int(1)] = (d_677_v119_).f20
                        nw78_[int(2)] = d_652_v110_
                        nw78_[int(3)] = d_652_v110_
                        nw78_[int(4)] = d_652_v110_
                        nw78_[int(5)] = (d_677_v119_).f20
                        nw78_[int(6)] = _dafny.CodePoint('u')
                        nw78_[int(7)] = (d_677_v119_).f20
                        nw78_[int(8)] = (d_677_v119_).f20
                        nw78_[int(9)] = d_652_v110_
                        nw78_[int(10)] = (d_677_v119_).f20
                        nw78_[int(11)] = d_652_v110_
                        nw78_[int(12)] = (d_677_v119_).f20
                        nw78_[int(13)] = _dafny.CodePoint('v')
                        nw78_[int(14)] = (d_677_v119_).f20
                        nw78_[int(15)] = d_652_v110_
                        nw78_[int(16)] = (d_677_v119_).f20
                        nw78_[int(17)] = (d_677_v119_).f20
                        nw78_[int(18)] = d_652_v110_
                        nw78_[int(19)] = _dafny.CodePoint('p')
                        nw78_[int(20)] = _dafny.CodePoint('b')
                        d_679_v121_ = nw78_
                        rhs77_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pwfti")) if d_613_v98_ else (d_458_v0_) + (d_458_v0_))
                        rhs78_ = not((d_678_v120_).ispropersubset((d_678_v120_).set(365, default__.abs((_dafny.MultiSet([d_679_v121_])).cardinality))))
                        lhs76_ = d_677_v119_
                        d_458_v0_ = rhs77_
                        lhs76_.f19 = rhs78_
                    pass
            pass

    def m5(self, p0, globalState):
        r0: bool = False
        d_680_v0_: _dafny.Array
        nw79_ = _dafny.Array(None, 3)
        nw79_[int(0)] = (self).f21
        nw79_[int(1)] = (self).f21
        nw79_[int(2)] = (self).f21
        d_680_v0_ = nw79_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_680_v0_).length(0)):
            d_681_i0_: int = guard_loop_4_
            if (True) and (((0) <= (d_681_i0_)) and ((d_681_i0_) < ((d_680_v0_).length(0)))):
                (d_680_v0_)[(d_681_i0_)] = (d_681_i0_) * ((p0) + (p0))
        (globalState).f7 = ((self).f21) + ((0) - (p0))
        d_682_v1_: str
        d_682_v1_ = _dafny.CodePoint('m')
        d_683_v2_: C0
        nw80_ = C0()
        nw80_.ctor__((self).f22, d_682_v1_, (self.f12 if (self).f22 else D1_DC2((self).f21, (self).fm8((self).f22, self.f16, globalState), (self).f13)), (self).f13)
        d_683_v2_ = nw80_
        d_684_v3_: _dafny.Array
        def lambda20_(d_685_v2_):
            def lambda21_(d_686_i1_):
                return _dafny.SeqWithoutIsStrInference([d_685_v2_.f19, (self).f17])

            return lambda21_

        init11_ = lambda20_(d_683_v2_)
        nw81_ = _dafny.Array(None, 18)
        for i0_11_ in range(nw81_.length(0)):
            nw81_[i0_11_] = init11_(i0_11_)
        d_684_v3_ = nw81_
        d_687_v4_: D8
        d_687_v4_ = D8_DC18(d_684_v3_)
        d_688_v5_: _dafny.Map
        d_688_v5_ = _dafny.Map({(d_687_v4_).cf31: _dafny.SeqWithoutIsStrInference([p0 for d_689_i2_ in range(default__.abs(762))])})
        d_690_v6_: _dafny.Seq
        d_690_v6_ = _dafny.SeqWithoutIsStrInference([p0])
        d_691_v7_: _dafny.Seq
        d_691_v7_ = _dafny.SeqWithoutIsStrInference([p0, (self).f21, (self).f21, len(d_690_v6_), p0])
        d_688_v5_ = (d_688_v5_).set(d_684_v3_, d_690_v6_)
        if (self).f17:
            d_692_v8_: _dafny.Seq
            d_692_v8_ = _dafny.SeqWithoutIsStrInference([d_683_v2_.f19, d_683_v2_.f19, (self).f13])
            index94_ = default__.safeIndex(132, (d_684_v3_).length(0))
            (d_684_v3_)[index94_] = d_692_v8_
            d_693_v9_: _dafny.Array
            def lambda22_(d_694_v2_, d_695_v1_):
                def lambda23_(d_696_i3_):
                    return _dafny.Map({(d_694_v2_).f20: _dafny.SeqWithoutIsStrInference([d_695_v1_])})

                return lambda23_

            init12_ = lambda22_(d_683_v2_, d_682_v1_)
            nw82_ = _dafny.Array(None, 9)
            for i0_12_ in range(nw82_.length(0)):
                nw82_[i0_12_] = init12_(i0_12_)
            d_693_v9_ = nw82_
            d_697_v10_: _dafny.Seq
            d_697_v10_ = _dafny.SeqWithoutIsStrInference([(d_683_v2_).f20, default__.fm33(d_683_v2_.f19, p0, globalState)])
            d_698_v11_: _dafny.Map
            d_698_v11_ = _dafny.Map({d_682_v1_: d_697_v10_})
            index95_ = default__.safeIndex(77, (d_693_v9_).length(0))
            (d_693_v9_)[index95_] = d_698_v11_
            index96_ = default__.safeIndex(132, (d_684_v3_).length(0))
            index97_ = default__.safeIndex(77, (d_693_v9_).length(0))
            rhs79_ = default__.fm0(((d_692_v8_).set(default__.safeIndex((self).f21, len(d_692_v8_)), d_683_v2_.f19)).set(default__.safeIndex((self).f21, len((d_692_v8_).set(default__.safeIndex((self).f21, len(d_692_v8_)), d_683_v2_.f19))), (self).f17), not(((self).f17 if not((self).f13) else d_683_v2_.f19)), (default__.fm30((d_690_v6_)[default__.safeIndex(p0, len(d_690_v6_))], globalState) if d_683_v2_.f19 else d_683_v2_.f19), default__.safeModuloInt(p0, p0), globalState)
            rhs80_ = (d_692_v8_) + (d_692_v8_)
            rhs81_ = d_698_v11_
            lhs77_ = globalState
            lhs78_ = d_684_v3_
            lhs79_ = default__.safeIndex(132, (d_684_v3_).length(0))
            lhs80_ = d_693_v9_
            lhs81_ = default__.safeIndex(77, (d_693_v9_).length(0))
            lhs77_.f7 = rhs79_
            lhs78_[lhs79_] = rhs80_
            lhs80_[lhs81_] = rhs81_
            (globalState).f7 = (self).fm9((d_683_v2_).f20, globalState)
            d_699_v12_: _dafny.MultiSet
            d_699_v12_ = _dafny.MultiSet([(self).f21, (self).f21])
            d_700_v13_: _dafny.MultiSet
            d_700_v13_ = _dafny.MultiSet([(_dafny.MultiSet([len((d_684_v3_)[default__.safeIndex(132, (d_684_v3_).length(0))]), (_dafny.MultiSet((d_684_v3_)[default__.safeIndex(132, (d_684_v3_).length(0))])).cardinality, (self).f21])).intersection((d_699_v12_).set(p0, default__.abs(len(d_697_v10_))))])
            (globalState).f7 = ((d_700_v13_)[d_699_v12_] if (d_699_v12_) in (d_700_v13_) else (self).f21)
            d_701_v14_: _dafny.Set
            d_701_v14_ = _dafny.Set({(d_697_v10_) + (d_697_v10_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wk")), d_697_v10_, (d_697_v10_) + (d_697_v10_)})
            d_701_v14_ = d_701_v14_
            (d_683_v2_).f19 = (self).f22
        elif True:
            d_702_v15_: _dafny.MultiSet
            d_702_v15_ = _dafny.MultiSet([(self).f22, d_683_v2_.f19])
            d_703_v16_: _dafny.Map
            d_703_v16_ = _dafny.Map({d_702_v15_: d_683_v2_.f19})
            d_704_v17_: _dafny.Map
            d_704_v17_ = _dafny.Map({(self).f21: _dafny.Map({d_702_v15_: d_683_v2_.f19})})
            d_703_v16_ = ((d_704_v17_)[(self).f21] if ((self).f21) in (d_704_v17_) else d_703_v16_)
            if (self).f13:
                d_705_v18_: _dafny.Seq
                d_705_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))
                d_706_v19_: D6
                d_706_v19_ = D6_DC13((self).f17, False, (self).f21, (self).f21, d_683_v2_.f19)
                d_707_v20_: _dafny.Set
                d_707_v20_ = _dafny.Set({d_683_v2_.f19})
                d_708_v21_: T1
                nw83_ = C1()
                nw83_.ctor__((d_683_v2_).f20, self.f12, (self).f17)
                d_708_v21_ = nw83_
                d_709_v22_: _dafny.MultiSet
                d_709_v22_ = _dafny.MultiSet([d_708_v21_])
                d_710_v23_: _dafny.Array
                nw84_ = _dafny.Array(None, 22)
                nw84_[int(0)] = d_683_v2_.f19
                nw84_[int(1)] = (self).f13
                nw84_[int(2)] = (d_705_v18_) <= (default__.fm32(D6_DC14(d_706_v19_), globalState))
                nw84_[int(3)] = d_683_v2_.f19
                nw84_[int(4)] = (_dafny.SeqWithoutIsStrInference([d_683_v2_.f19, (self).f22])) < (default__.fm38((self).f13, (self).f22, d_683_v2_.f19, (self).f17, globalState))
                nw84_[int(5)] = (self).f22
                nw84_[int(6)] = default__.fm30(-545, globalState)
                nw84_[int(7)] = default__.fm30(p0, globalState)
                nw84_[int(8)] = d_683_v2_.f19
                nw84_[int(9)] = (self).f13
                nw84_[int(10)] = (not(d_683_v2_.f19) if (self).f13 else True)
                nw84_[int(11)] = (p0) >= ((d_690_v6_)[default__.safeIndex(333, len(d_690_v6_))])
                nw84_[int(12)] = ((self).f21) <= (274)
                nw84_[int(13)] = (self).f17
                nw84_[int(14)] = (d_707_v20_).issubset(d_707_v20_)
                nw84_[int(15)] = (self).f13
                nw84_[int(16)] = (self).f17
                nw84_[int(17)] = (d_709_v22_).issubset(d_709_v22_)
                nw84_[int(18)] = ((self).f21) > (p0)
                nw84_[int(19)] = (len(d_691_v7_)) <= ((self).f21)
                nw84_[int(20)] = d_683_v2_.f19
                nw84_[int(21)] = (d_705_v18_) < (d_705_v18_)
                d_710_v23_ = nw84_
                d_710_v23_ = d_710_v23_
                d_711_v24_: _dafny.Array
                nw85_ = _dafny.Array(_dafny.Array(None, 0), 23)
                d_711_v24_ = nw85_
                d_712_v25_: _dafny.Array
                nw86_ = _dafny.Array(_dafny.Map({}), 3)
                d_712_v25_ = nw86_
                index98_ = default__.safeIndex(815, (d_711_v24_).length(0))
                (d_711_v24_)[index98_] = d_712_v25_
                index99_ = default__.safeIndex(815, (d_711_v24_).length(0))
                (d_711_v24_)[index99_] = d_712_v25_
                (d_683_v2_).f19 = (not((d_708_v21_).f13) if d_683_v2_.f19 else default__.fm30(p0, globalState))
                index100_ = default__.safeIndex(147, (d_710_v23_).length(0))
                (d_710_v23_)[index100_] = not((self).f13)
                index101_ = default__.safeIndex(147, (d_710_v23_).length(0))
                (d_710_v23_)[index101_] = (self).f17
                d_713_v26_: _dafny.Map
                d_713_v26_ = _dafny.Map({False: (d_708_v21_).f13})
                d_690_v6_ = ((d_690_v6_) + (_dafny.SeqWithoutIsStrInference([p0 for d_714_i4_ in range(default__.abs(-120))]))).set(default__.safeIndex(p0, len((d_690_v6_) + (_dafny.SeqWithoutIsStrInference([p0 for d_715_i4_ in range(default__.abs(-120))])))), len(d_713_v26_))
            elif True:
                d_716_v27_: _dafny.Seq
                d_716_v27_ = _dafny.SeqWithoutIsStrInference([True, (self).f22, (self).f22])
                (globalState).f7 = default__.fm0(d_716_v27_, False, (self).f22, (p0) - ((self).f21), globalState)
                d_717_v28_: _dafny.Array
                nw87_ = _dafny.Array(_dafny.Seq({}), 24)
                d_717_v28_ = nw87_
                d_718_v29_: _dafny.Set
                d_718_v29_ = _dafny.Set({d_683_v2_.f19})
                d_719_v30_: D8
                d_719_v30_ = D8_DC19(_dafny.Map({not(d_683_v2_.f19): (self).f21}), d_718_v29_, _dafny.MultiSet([680]), self.f16)
                d_720_v31_: _dafny.Seq
                d_720_v31_ = _dafny.SeqWithoutIsStrInference([d_718_v29_, (d_719_v30_).cf33, d_718_v29_, d_718_v29_, d_718_v29_])
                index102_ = default__.safeIndex(846, (d_717_v28_).length(0))
                (d_717_v28_)[index102_] = d_720_v31_
                index103_ = default__.safeIndex(846, (d_717_v28_).length(0))
                (d_717_v28_)[index103_] = d_720_v31_
                d_721_v32_: _dafny.MultiSet
                out8_: _dafny.MultiSet
                out8_ = default__.m0((d_683_v2_).f20, (self).f22, globalState)
                d_721_v32_ = out8_
                (globalState).f7 = (0) - (((p0) * (p0)) * (p0))
                (globalState).f7 = (0) - ((p0 if not ((self).f22) or (d_683_v2_.f19) else p0))
            d_722_v33_: D6
            d_722_v33_ = D6_DC11(_dafny.MultiSet((d_690_v6_) + (d_690_v6_)))
            d_722_v33_ = d_722_v33_
            r0 = d_683_v2_.f19
            (globalState).f7 = ((self).f21) - ((0) - (p0))
        (globalState).f7 = (-221) - (p0)
        r0 = (self).f22
        return r0

    def m13(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        d_723_v0_: _dafny.Array
        nw88_ = _dafny.Array(False, 4)
        d_723_v0_ = nw88_
        d_723_v0_ = d_723_v0_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_723_v0_).length(0)):
            d_724_i0_: int = guard_loop_5_
            if (True) and (((0) <= (d_724_i0_)) and ((d_724_i0_) < ((d_723_v0_).length(0)))):
                (d_723_v0_)[(d_724_i0_)] = (self).f17
        d_725_v1_: _dafny.Map
        d_725_v1_ = _dafny.Map({(self).f21: (self).f21})
        d_726_v2_: _dafny.Seq
        d_726_v2_ = _dafny.SeqWithoutIsStrInference([len(d_725_v1_), (self).f21, (self).f21])
        d_726_v2_ = d_726_v2_
        d_727_v3_: _dafny.Map
        d_727_v3_ = _dafny.Map({not((908) == (p3)): (self).f22})
        d_727_v3_ = (d_727_v3_).set(((self).f17) == (p0), (self).f17)
        if (self).f22:
            (globalState).f7 = (self).fm8((self).f17, self.f16, globalState)
            d_728_v4_: _dafny.Array
            nw89_ = _dafny.Array(_dafny.Seq({}), 23)
            d_728_v4_ = nw89_
            d_728_v4_ = d_728_v4_
            d_729_v5_: D3
            d_729_v5_ = D3_DC5(p1)
            d_730_v6_: C1
            nw90_ = C1()
            nw90_.ctor__((d_729_v5_).cf11, D1_DC2(-728, p3, (self).f17), True)
            d_730_v6_ = nw90_
            d_731_v7_: _dafny.MultiSet
            out9_: _dafny.MultiSet
            out9_ = default__.m0(d_730_v6_.f23, (self).f22, globalState)
            d_731_v7_ = out9_
            r0 = default__.fm30(p3, globalState)
        elif True:
            d_732_v8_: _dafny.Seq
            d_732_v8_ = _dafny.SeqWithoutIsStrInference([p0])
            (globalState).f7 = ((0) - ((self).f21)) - (default__.fm0(d_732_v8_, (self).f17, (self).f13, (self).f21, globalState))
            d_733_v9_: _dafny.Seq
            d_733_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rkdajmk"))
            (globalState).f7 = len((d_733_v9_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lvxjwypfe"))))
            d_734_v10_: _dafny.Map
            d_734_v10_ = _dafny.Map({(self).f21: d_723_v0_})
            d_734_v10_ = (d_734_v10_).set((self).f21, (self).f14)
            (globalState).f7 = p3
            d_735_v11_: _dafny.Array
            nw91_ = _dafny.Array(int(0), 21)
            d_735_v11_ = nw91_
            d_736_v12_: _dafny.Map
            d_736_v12_ = _dafny.Map({d_723_v0_: d_735_v11_})
            d_736_v12_ = (d_736_v12_) | (_dafny.Map({d_723_v0_: d_735_v11_}))
        r0 = not((self).f13)
        r0 = (self).f22
        return r0

    @property
    def f21(self):
        return self._f21
    @property
    def f22(self):
        return self._f22

class C3(T3, T0, T1, T2):
    def  __init__(self):
        self._f12: D1 = D1.default()()
        self._f16: D1 = D1.default()()
        self._f14: _dafny.Array = _dafny.Array(None, 0)
        self._f15: _dafny.Set = _dafny.Set({})
        self._f13: bool = False
        self._f17: bool = False
        self._f18: bool = False
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
    def f16(self):
        return self._f16
    @f16.setter
    def f16(self, value):
        self._f16 = value
    @property
    def f14(self):
        return self._f14
    @property
    def f15(self):
        return self._f15
    @property
    def f13(self):
        return self._f13
    @property
    def f17(self):
        return self._f17
    def ctor__(self, f18, f16, f17, f14, f15, f12, f13):
        (self)._f18 = f18
        (self).f16 = f16
        (self)._f17 = f17
        (self)._f14 = f14
        (self)._f15 = f15
        (self).f12 = f12
        (self)._f13 = f13

    def fm10(self, p0, p1, p2, globalState):
        return _dafny.Map({not((self).f18): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tfdhco")))})

    def fm8(self, p0, p1, globalState):
        return default__.safeModuloInt((144) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t")))), 150)

    def fm9(self, p0, globalState):
        return ((0) - ((0) - ((874) + ((_dafny.MultiSet([len(_dafny.Map({(self).f18: (self).f17}))])).cardinality)))) * ((len(_dafny.Map({(self).f17: (self).f17})) if False else 537))

    def fm15(self, globalState):
        return (0) - ((default__.safeModuloInt(981, 611)) - (default__.safeDivisionInt(-191, 434)))

    def m6(self, p0, globalState):
        r0: int = int(0)
        if (self.f12).cf4:
            d_737_v0_: _dafny.Array
            def lambda24_(d_738_p0_):
                def lambda25_(d_739_i0_):
                    return (d_739_i0_) - ((_dafny.MultiSet([(self).f13, d_738_p0_, (self).f18, (self).f17])).cardinality)

                return lambda25_

            init13_ = lambda24_(p0)
            nw92_ = _dafny.Array(None, 15)
            for i0_13_ in range(nw92_.length(0)):
                nw92_[i0_13_] = init13_(i0_13_)
            d_737_v0_ = nw92_
            d_740_v1_: int
            d_740_v1_ = 673
            index104_ = default__.safeIndex(717, (d_737_v0_).length(0))
            (d_737_v0_)[index104_] = (0) - (default__.safeDivisionInt(d_740_v1_, len(_dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f13: len(_dafny.SeqWithoutIsStrInference([p0]))}) for d_741_i1_ in range(default__.abs(31))]))))
            d_742_v2_: _dafny.Map
            d_742_v2_ = _dafny.Map({not (True) or ((self).f13): not((self).f17)})
            d_743_v3_: _dafny.Seq
            d_743_v3_ = _dafny.SeqWithoutIsStrInference([111])
            index105_ = default__.safeIndex(717, (d_737_v0_).length(0))
            rhs82_ = (default__.safeModuloInt(len(d_743_v3_), d_740_v1_)) + ((d_743_v3_)[default__.safeIndex((0) - (d_740_v1_), len(d_743_v3_))])
            rhs83_ = d_742_v2_
            rhs84_ = d_740_v1_
            lhs82_ = d_737_v0_
            lhs83_ = default__.safeIndex(717, (d_737_v0_).length(0))
            lhs82_[lhs83_] = rhs82_
            d_742_v2_ = rhs83_
            d_740_v1_ = rhs84_
            d_744_v4_: bool
            d_744_v4_ = True
            d_745_v5_: _dafny.Seq
            d_745_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "db"))
            d_744_v4_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rgxktld"))) == (d_745_v5_)
            d_746_v6_: _dafny.Map
            d_746_v6_ = _dafny.Map({(d_737_v0_)[default__.safeIndex(717, (d_737_v0_).length(0))]: ((d_737_v0_)[default__.safeIndex(717, (d_737_v0_).length(0))]) - (d_740_v1_)})
            d_747_v7_: _dafny.Seq
            d_747_v7_ = _dafny.SeqWithoutIsStrInference([(self).f17, (self).f13])
            d_746_v6_ = (d_746_v6_).set((d_737_v0_)[default__.safeIndex(717, (d_737_v0_).length(0))], default__.safeModuloInt(default__.fm0(d_747_v7_, not(True), (self).f17, d_740_v1_, globalState), (d_737_v0_)[default__.safeIndex(717, (d_737_v0_).length(0))]))
            d_748_v8_: str
            d_748_v8_ = _dafny.CodePoint('w')
            d_749_v9_: T0
            nw93_ = C0()
            nw93_.ctor__((self).f13, d_748_v8_, self.f12, (self).f18)
            d_749_v9_ = nw93_
            d_750_v10_: _dafny.Seq
            d_750_v10_ = _dafny.SeqWithoutIsStrInference([d_749_v9_, d_749_v9_])
            d_751_v11_: _dafny.Map
            d_751_v11_ = _dafny.Map({(self).f17: len(d_750_v10_)})
            index106_ = default__.safeIndex(717, (d_737_v0_).length(0))
            (d_737_v0_)[index106_] = (d_743_v3_)[default__.safeIndex(len(d_751_v11_), len(d_743_v3_))]
            r0 = d_740_v1_
        elif True:
            d_752_v12_: int
            d_752_v12_ = 352
            d_753_v13_: _dafny.MultiSet
            d_753_v13_ = _dafny.MultiSet([p0])
            d_754_v14_: _dafny.Seq
            d_754_v14_ = _dafny.SeqWithoutIsStrInference([d_753_v13_])
            d_755_v15_: _dafny.Map
            d_755_v15_ = _dafny.Map({d_752_v12_: len(d_754_v14_)})
            (globalState).f7 = default__.safeDivisionInt(len(d_755_v15_), 6)
            r0 = d_752_v12_
            if (self).f18:
                rhs85_ = len((self).f15)
                rhs86_ = (0) - (-213)
                lhs84_ = globalState
                lhs84_.f7 = rhs85_
                d_752_v12_ = rhs86_
                d_756_v16_: _dafny.Seq
                d_756_v16_ = _dafny.SeqWithoutIsStrInference([(self).f14])
                d_752_v12_ = len(d_756_v16_)
                d_757_v17_: bool
                d_757_v17_ = True
                d_757_v17_ = default__.fm17(globalState)
                d_758_v18_: _dafny.Seq
                d_758_v18_ = _dafny.SeqWithoutIsStrInference([d_752_v12_])
                d_759_v19_: _dafny.MultiSet
                d_759_v19_ = _dafny.MultiSet([len(d_758_v18_)])
                d_760_v20_: _dafny.Seq
                d_760_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))
                d_761_v21_: _dafny.Array
                nw94_ = _dafny.Array(int(0), 12)
                d_761_v21_ = nw94_
                d_762_v22_: D2
                d_762_v22_ = D2_DC4(d_761_v21_, (self).f18, 945, default__.fm17(globalState), (self).f14)
                d_763_v23_: _dafny.Array
                nw95_ = _dafny.Array(None, 17)
                nw95_[int(0)] = len(d_758_v18_)
                nw95_[int(1)] = (0) - (d_752_v12_)
                nw95_[int(2)] = d_752_v12_
                nw95_[int(3)] = d_752_v12_
                nw95_[int(4)] = default__.safeDivisionInt(d_752_v12_, d_752_v12_)
                nw95_[int(5)] = len(d_755_v15_)
                nw95_[int(6)] = d_752_v12_
                nw95_[int(7)] = d_752_v12_
                nw95_[int(8)] = d_752_v12_
                nw95_[int(9)] = default__.safeDivisionInt(d_752_v12_, d_752_v12_)
                nw95_[int(10)] = (d_759_v19_).cardinality
                nw95_[int(11)] = default__.safeDivisionInt(d_752_v12_, d_752_v12_)
                nw95_[int(12)] = len(d_760_v20_)
                nw95_[int(13)] = (d_762_v22_).cf8
                nw95_[int(14)] = len((d_758_v18_).set(default__.safeIndex(len(d_760_v20_), len(d_758_v18_)), 263))
                nw95_[int(15)] = d_752_v12_
                nw95_[int(16)] = d_752_v12_
                d_763_v23_ = nw95_
                index107_ = default__.safeIndex(198, (d_761_v21_).length(0))
                (d_761_v21_)[index107_] = 680
                index108_ = default__.safeIndex(198, (d_761_v21_).length(0))
                rhs87_ = d_752_v12_
                rhs88_ = d_752_v12_
                lhs85_ = d_761_v21_
                lhs86_ = default__.safeIndex(198, (d_761_v21_).length(0))
                lhs85_[lhs86_] = rhs87_
                d_752_v12_ = rhs88_
                d_764_v24_: str
                d_764_v24_ = _dafny.CodePoint('l')
                d_765_v25_: D3
                d_765_v25_ = D3_DC5(_dafny.CodePoint('s'))
                d_764_v24_ = (d_765_v25_).cf11
            elif True:
                d_766_v26_: str
                d_766_v26_ = _dafny.CodePoint('w')
                d_767_v27_: _dafny.Map
                d_767_v27_ = _dafny.Map({(self).f18: d_766_v26_})
                d_768_v28_: _dafny.Seq
                d_768_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fshc"))
                d_769_v29_: _dafny.Seq
                d_769_v29_ = _dafny.SeqWithoutIsStrInference([(self).f13])
                d_770_v30_: _dafny.Array
                nw96_ = _dafny.Array(None, 25)
                nw96_[int(0)] = d_766_v26_
                nw96_[int(1)] = d_766_v26_
                nw96_[int(2)] = d_766_v26_
                nw96_[int(3)] = (d_766_v26_ if not((self).f13) else ((d_767_v27_)[p0] if (p0) in (d_767_v27_) else d_766_v26_))
                nw96_[int(4)] = d_766_v26_
                nw96_[int(5)] = d_766_v26_
                nw96_[int(6)] = d_766_v26_
                nw96_[int(7)] = _dafny.CodePoint('a')
                nw96_[int(8)] = d_766_v26_
                nw96_[int(9)] = _dafny.CodePoint('k')
                nw96_[int(10)] = d_766_v26_
                nw96_[int(11)] = _dafny.CodePoint('i')
                nw96_[int(12)] = d_766_v26_
                nw96_[int(13)] = d_766_v26_
                nw96_[int(14)] = d_766_v26_
                nw96_[int(15)] = _dafny.CodePoint('b')
                nw96_[int(16)] = default__.fm18(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukupbqlom"))), d_768_v28_, (self).f17, (self).fm10((self).f13, (self).f13, d_769_v29_, globalState), globalState)
                nw96_[int(17)] = _dafny.CodePoint('w')
                nw96_[int(18)] = (d_766_v26_ if not((self).f18) else d_766_v26_)
                nw96_[int(19)] = (d_768_v28_)[default__.safeIndex((self).fm9(d_766_v26_, globalState), len(d_768_v28_))]
                nw96_[int(20)] = d_766_v26_
                nw96_[int(21)] = d_766_v26_
                nw96_[int(22)] = d_766_v26_
                nw96_[int(23)] = d_766_v26_
                nw96_[int(24)] = d_766_v26_
                d_770_v30_ = nw96_
                d_770_v30_ = d_770_v30_
                d_771_v31_: bool
                d_771_v31_ = False
                d_771_v31_ = ((d_752_v12_) + (d_752_v12_)) <= (d_752_v12_)
                d_771_v31_ = d_771_v31_
                d_772_v32_: _dafny.Array
                nw97_ = _dafny.Array(_dafny.Seq({}), 10)
                d_772_v32_ = nw97_
                d_773_v33_: _dafny.MultiSet
                d_773_v33_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lrwaorxc")))])
                d_774_v34_: _dafny.Seq
                d_774_v34_ = _dafny.SeqWithoutIsStrInference([d_773_v33_])
                index109_ = default__.safeIndex(133, (d_772_v32_).length(0))
                (d_772_v32_)[index109_] = (d_774_v34_ if (self).f17 else _dafny.SeqWithoutIsStrInference([d_773_v33_, d_773_v33_]))
                index110_ = default__.safeIndex(133, (d_772_v32_).length(0))
                (d_772_v32_)[index110_] = d_774_v34_
                d_775_v35_: _dafny.Array
                nw98_ = _dafny.Array(_dafny.Array(None, 0), 28)
                d_775_v35_ = nw98_
                d_775_v35_ = d_775_v35_
            d_776_v36_: str
            d_776_v36_ = _dafny.CodePoint('g')
            d_777_v37_: C0
            nw99_ = C0()
            nw99_.ctor__((d_752_v12_) > (d_752_v12_), d_776_v36_, D1_DC2(d_752_v12_, d_752_v12_, p0), p0)
            d_777_v37_ = nw99_
            d_778_v38_: _dafny.Seq
            d_778_v38_ = _dafny.SeqWithoutIsStrInference([(self).f14])
            d_779_v39_: _dafny.Seq
            d_779_v39_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))
            d_780_v40_: _dafny.Array
            nw100_ = _dafny.Array(None, 14)
            nw100_[int(0)] = (self).f14
            nw100_[int(1)] = (self).f14
            nw100_[int(2)] = (self).f14
            nw100_[int(3)] = (self).f14
            nw100_[int(4)] = (self).f14
            nw100_[int(5)] = (self).f14
            nw100_[int(6)] = (self).f14
            nw100_[int(7)] = (self).f14
            nw100_[int(8)] = (self).f14
            nw100_[int(9)] = (self).f14
            nw100_[int(10)] = (self).f14
            nw100_[int(11)] = (self).f14
            nw100_[int(12)] = (d_778_v38_)[default__.safeIndex(len(d_779_v39_), len(d_778_v38_))]
            nw100_[int(13)] = (self).f14
            d_780_v40_ = nw100_
            d_781_v41_: _dafny.Array
            d_781_v41_ = d_780_v40_
            d_782_v42_: _dafny.Seq
            d_782_v42_ = _dafny.SeqWithoutIsStrInference([d_780_v40_, d_780_v40_, (d_781_v41_), d_780_v40_])
            d_783_v43_: _dafny.Array
            nw101_ = _dafny.Array(None, 14)
            nw101_[int(0)] = d_780_v40_
            nw101_[int(1)] = d_780_v40_
            nw101_[int(2)] = d_780_v40_
            nw101_[int(3)] = d_780_v40_
            nw101_[int(4)] = d_780_v40_
            nw101_[int(5)] = (d_782_v42_)[default__.safeIndex(d_752_v12_, len(d_782_v42_))]
            nw101_[int(6)] = d_780_v40_
            nw101_[int(7)] = d_780_v40_
            nw101_[int(8)] = (d_782_v42_)[default__.safeIndex(d_752_v12_, len(d_782_v42_))]
            nw101_[int(9)] = d_780_v40_
            nw101_[int(10)] = d_780_v40_
            nw101_[int(11)] = d_780_v40_
            nw101_[int(12)] = d_780_v40_
            nw101_[int(13)] = (d_782_v42_)[default__.safeIndex(d_752_v12_, len(d_782_v42_))]
            d_783_v43_ = nw101_
            index111_ = default__.safeIndex(976, (d_783_v43_).length(0))
            (d_783_v43_)[index111_] = d_780_v40_
            index112_ = default__.safeIndex(976, (d_783_v43_).length(0))
            (d_783_v43_)[index112_] = d_780_v40_
        d_784_v44_: _dafny.Array
        def lambda26_(d_785_p0_):
            def lambda27_(d_786_i2_):
                return _dafny.Map({len(_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.Map({False: (self).f18}), _dafny.Map({(self).f17: (self).f13})]))): (0) - (len(_dafny.Map({d_785_p0_: len(_dafny.SeqWithoutIsStrInference([(self).f18]))})))})): not(d_785_p0_)})

            return lambda27_

        init14_ = lambda26_(p0)
        nw102_ = _dafny.Array(None, 4)
        for i0_14_ in range(nw102_.length(0)):
            nw102_[i0_14_] = init14_(i0_14_)
        d_784_v44_ = nw102_
        d_787_v45_: int
        d_787_v45_ = 10
        d_788_v46_: _dafny.Map
        d_788_v46_ = _dafny.Map({d_784_v44_: d_787_v45_})
        d_788_v46_ = (d_788_v46_).set(d_784_v44_, default__.safeModuloInt(771, 190))
        if default__.fm17(globalState):
            d_789_v47_: _dafny.Array
            nw103_ = _dafny.Array(_dafny.Array(None, 0), 22)
            d_789_v47_ = nw103_
            d_790_v48_: _dafny.Array
            d_790_v48_ = d_789_v47_
            d_791_v49_: _dafny.Array
            nw104_ = _dafny.Array(None, 4)
            nw104_[int(0)] = d_789_v47_
            nw104_[int(1)] = d_790_v48_
            nw104_[int(2)] = d_790_v48_
            nw104_[int(3)] = d_789_v47_
            d_791_v49_ = nw104_
            rhs89_ = d_791_v49_
            rhs90_ = d_787_v45_
            lhs87_ = globalState
            d_791_v49_ = rhs89_
            lhs87_.f7 = rhs90_
            d_792_v50_: _dafny.Map
            d_792_v50_ = _dafny.Map({(self).f17: d_787_v45_})
            d_793_v51_: _dafny.Seq
            d_793_v51_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pwgijra"))
            d_794_v52_: _dafny.MultiSet
            d_794_v52_ = _dafny.MultiSet([(self).f13])
            d_795_v53_: _dafny.Seq
            d_795_v53_ = _dafny.SeqWithoutIsStrInference([d_794_v52_, d_794_v52_, d_794_v52_])
            d_796_v54_: _dafny.Array
            nw105_ = _dafny.Array(int(0), 17)
            d_796_v54_ = nw105_
            d_797_v55_: D3
            d_797_v55_ = D3_DC7()
            d_798_v56_: _dafny.Map
            d_798_v56_ = _dafny.Map({d_796_v54_: d_797_v55_})
            d_799_v57_: _dafny.Seq
            d_799_v57_ = _dafny.SeqWithoutIsStrInference([d_787_v45_])
            d_800_v58_: str
            d_800_v58_ = _dafny.CodePoint('p')
            d_801_v59_: _dafny.Seq
            d_801_v59_ = _dafny.SeqWithoutIsStrInference([d_796_v54_])
            d_802_v60_: D3
            d_802_v60_ = D3_DC8(698)
            d_803_v61_: _dafny.Array
            nw106_ = _dafny.Array(None, 26)
            nw106_[int(0)] = d_787_v45_
            nw106_[int(1)] = 268
            nw106_[int(2)] = (d_787_v45_) * (d_787_v45_)
            nw106_[int(3)] = d_787_v45_
            nw106_[int(4)] = d_787_v45_
            nw106_[int(5)] = d_787_v45_
            nw106_[int(6)] = default__.safeModuloInt(d_787_v45_, d_787_v45_)
            nw106_[int(7)] = ((d_792_v50_)[(self).f17] if ((self).f17) in (d_792_v50_) else len(d_793_v51_))
            nw106_[int(8)] = ((d_795_v53_)[default__.safeIndex((default__.fm19(globalState)).cardinality, len(d_795_v53_))]).cardinality
            nw106_[int(9)] = len(d_798_v56_)
            nw106_[int(10)] = d_787_v45_
            nw106_[int(11)] = d_787_v45_
            nw106_[int(12)] = d_787_v45_
            nw106_[int(13)] = d_787_v45_
            nw106_[int(14)] = d_787_v45_
            nw106_[int(15)] = (d_799_v57_)[default__.safeIndex(len((default__.fm20(globalState)).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_804_i3_ in range(default__.abs(90))])), len(default__.fm20(globalState))), d_800_v58_)), len(d_799_v57_))]
            nw106_[int(16)] = (self).fm8((self).f18, default__.fm21(globalState), globalState)
            nw106_[int(17)] = d_787_v45_
            nw106_[int(18)] = d_787_v45_
            nw106_[int(19)] = d_787_v45_
            nw106_[int(20)] = (self).fm8(default__.fm17(globalState), self.f16, globalState)
            nw106_[int(21)] = d_787_v45_
            nw106_[int(22)] = ((_dafny.MultiSet([d_787_v45_])) | (_dafny.MultiSet([d_787_v45_]))).cardinality
            nw106_[int(23)] = default__.safeDivisionInt(d_787_v45_, len(d_801_v59_))
            nw106_[int(24)] = d_787_v45_
            nw106_[int(25)] = (d_802_v60_).cf14
            d_803_v61_ = nw106_
            index113_ = default__.safeIndex(374, (d_796_v54_).length(0))
            (d_796_v54_)[index113_] = (((d_792_v50_)[(self).f17] if ((self).f17) in (d_792_v50_) else len(d_799_v57_))) - (d_787_v45_)
            index114_ = default__.safeIndex(374, (d_796_v54_).length(0))
            (d_796_v54_)[index114_] = (d_787_v45_) - ((d_787_v45_ if p0 else len(default__.fm20(globalState))))
            (globalState).f7 = (0) - ((d_796_v54_)[default__.safeIndex(374, (d_796_v54_).length(0))])
            (globalState).f7 = d_787_v45_
            d_805_v63_: _dafny.Array
            def lambda28_(d_806_p0_, d_807_v58_, d_808_v51_):
                def lambda29_(d_809_i4_):
                    def iife71_():
                        coll35_ = _dafny.Set()
                        compr_35_: str
                        for compr_35_ in (_dafny.Map({(d_808_v51_)[default__.safeIndex(637, len(d_808_v51_))]: not((self).f17)})).keys.Elements:
                            d_810_v62_: str = compr_35_
                            if (d_810_v62_) in (_dafny.Map({(d_808_v51_)[default__.safeIndex(637, len(d_808_v51_))]: not((self).f17)})):
                                coll35_ = coll35_.union(_dafny.Set([d_810_v62_]))
                        return _dafny.Set(coll35_)
                    return (iife71_()
                     if d_806_p0_ else _dafny.Set({d_807_v58_, d_807_v58_}))

                return lambda29_

            init15_ = lambda28_(p0, d_800_v58_, d_793_v51_)
            nw107_ = _dafny.Array(None, 19)
            for i0_15_ in range(nw107_.length(0)):
                nw107_[i0_15_] = init15_(i0_15_)
            d_805_v63_ = nw107_
            d_811_v64_: _dafny.Map
            d_811_v64_ = _dafny.Map({494: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uh"))})
            d_812_v65_: _dafny.Set
            d_812_v65_ = _dafny.Set({((d_811_v64_)[d_787_v45_] if (d_787_v45_) in (d_811_v64_) else d_793_v51_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dgxndwfh")), d_793_v51_})
            d_813_v66_: _dafny.Set
            d_813_v66_ = d_812_v65_
            d_814_v67_: _dafny.Set
            d_814_v67_ = _dafny.Set({not(default__.fm17(globalState)), True, (self).f13, (self).f13})
            rhs91_ = d_805_v63_
            rhs92_ = (d_793_v51_) + ((d_793_v51_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wxpgj"))))
            rhs93_ = (d_813_v66_)
            rhs94_ = d_796_v54_
            rhs95_ = default__.safeModuloInt((d_787_v45_) - (len(d_814_v67_)), (d_796_v54_)[default__.safeIndex(374, (d_796_v54_).length(0))])
            d_805_v63_ = rhs91_
            d_793_v51_ = rhs92_
            d_812_v65_ = rhs93_
            d_803_v61_ = rhs94_
            d_787_v45_ = rhs95_
        elif True:
            d_815_v69_: _dafny.Seq
            d_815_v69_ = _dafny.SeqWithoutIsStrInference([d_787_v45_, d_787_v45_])
            d_816_v70_: _dafny.MultiSet
            d_816_v70_ = _dafny.MultiSet([len(d_815_v69_), (0) - (d_787_v45_)])
            d_817_v71_: _dafny.Seq
            d_817_v71_ = _dafny.SeqWithoutIsStrInference([d_816_v70_])
            d_818_v73_: D6
            d_818_v73_ = D6_DC11(d_816_v70_)
            d_819_v74_: _dafny.MultiSet
            d_819_v74_ = _dafny.MultiSet([(d_818_v73_).cf17])
            d_820_v75_: _dafny.Map
            def iife72_():
                coll36_ = _dafny.Map()
                compr_36_: _dafny.MultiSet
                for compr_36_ in (d_817_v71_).Elements:
                    d_821_v68_: _dafny.MultiSet = compr_36_
                    if (d_821_v68_) in (d_817_v71_):
                        coll36_[d_821_v68_] = True
                return _dafny.Map(coll36_)
            def iife73_():
                coll37_ = _dafny.Map()
                compr_37_: _dafny.MultiSet
                for compr_37_ in (d_819_v74_).Elements:
                    d_822_v72_: _dafny.MultiSet = compr_37_
                    if (d_822_v72_) in (d_819_v74_):
                        coll37_[d_822_v72_] = (self).f13
                return _dafny.Map(coll37_)
            d_820_v75_ = _dafny.Map({(iife72_()
            ) | (iife73_()
            ): False})
            d_823_v76_: _dafny.Map
            d_823_v76_ = _dafny.Map({d_816_v70_: (self).f13})
            d_820_v75_ = (d_820_v75_).set(d_823_v76_, p0)
            d_824_v77_: D2
            d_824_v77_ = D2_DC3((self).f14)
            source15_ = d_824_v77_
            if source15_.is_DC4:
                d_825___mcc_h0_ = source15_.cf6
                d_826___mcc_h1_ = source15_.cf7
                d_827___mcc_h2_ = source15_.cf8
                d_828___mcc_h3_ = source15_.cf9
                d_829___mcc_h4_ = source15_.cf10
                d_830_cf10_ = d_829___mcc_h4_
                d_831_cf9_ = d_828___mcc_h3_
                d_832_cf8_ = d_827___mcc_h2_
                d_833_cf7_ = d_826___mcc_h1_
                d_834_cf6_ = d_825___mcc_h0_
                d_835_v78_: _dafny.Seq
                d_835_v78_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fkeo"))
                d_836_v79_: _dafny.Set
                d_836_v79_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gn")), d_835_v78_})
                d_837_v80_: _dafny.Set
                d_837_v80_ = d_836_v79_
                d_838_v81_: _dafny.Map
                d_838_v81_ = _dafny.Map({_dafny.MultiSet([d_837_v80_, d_837_v80_, default__.fm22(d_831_cf9_, _dafny.SeqWithoutIsStrInference([-328 for d_839_i5_ in range(default__.abs(-793))]), globalState), d_837_v80_, d_836_v79_]): p0})
                d_840_v82_: _dafny.MultiSet
                d_840_v82_ = _dafny.MultiSet([d_837_v80_, d_837_v80_])
                d_841_v83_: _dafny.Seq
                d_841_v83_ = _dafny.SeqWithoutIsStrInference([d_831_cf9_])
                d_838_v81_ = (d_838_v81_).set((d_840_v82_).intersection(d_840_v82_), (d_841_v83_)[default__.safeIndex((_dafny.MultiSet([d_832_cf8_])).cardinality, len(d_841_v83_))])
                d_842_v84_: str
                d_842_v84_ = _dafny.CodePoint('b')
                d_843_v85_: _dafny.Set
                d_843_v85_ = _dafny.Set({d_834_cf6_, d_834_cf6_})
                d_844_v86_: C0
                nw108_ = C0()
                nw108_.ctor__((self).f18, d_842_v84_, self.f12, (d_834_cf6_) not in (d_843_v85_))
                d_844_v86_ = nw108_
                d_845_v87_: _dafny.Seq
                d_845_v87_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))), d_835_v78_, d_835_v78_, ((self.f16).cf1) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "te"))), _dafny.SeqWithoutIsStrInference([(D3_DC5((d_844_v86_).f20)).cf11 for d_846_i6_ in range(default__.abs(733))])])
                d_847_v88_: _dafny.Map
                d_847_v88_ = _dafny.Map({d_832_cf8_: d_845_v87_})
                d_845_v87_ = ((d_847_v88_)[d_787_v45_] if (d_787_v45_) in (d_847_v88_) else d_845_v87_)
                d_831_cf9_ = d_833_cf7_
            elif True:
                d_848___mcc_h5_ = source15_.cf5
                d_849_cf5_ = d_848___mcc_h5_
                rhs96_ = d_787_v45_
                rhs97_ = (0) - (d_787_v45_)
                lhs88_ = globalState
                lhs88_.f7 = rhs96_
                r0 = rhs97_
                d_850_v89_: _dafny.Seq
                d_850_v89_ = _dafny.SeqWithoutIsStrInference([(self).f13])
                (globalState).f2 = d_850_v89_
                d_851_v90_: bool
                d_851_v90_ = True
                d_852_v91_: str
                d_852_v91_ = _dafny.CodePoint('x')
                d_853_v92_: _dafny.Array
                def lambda30_(d_854_i7_):
                    return self.f12

                init16_ = lambda30_
                nw109_ = _dafny.Array(None, 29)
                for i0_16_ in range(nw109_.length(0)):
                    nw109_[i0_16_] = init16_(i0_16_)
                d_853_v92_ = nw109_
                index115_ = default__.safeIndex(193, (d_853_v92_).length(0))
                (d_853_v92_)[index115_] = self.f12
                d_855_v93_: _dafny.Array
                def lambda31_(d_856_v45_):
                    def lambda32_(d_857_i8_):
                        return (d_857_i8_) + (d_856_v45_)

                    return lambda32_

                init17_ = lambda31_(d_787_v45_)
                nw110_ = _dafny.Array(None, 28)
                for i0_17_ in range(nw110_.length(0)):
                    nw110_[i0_17_] = init17_(i0_17_)
                d_855_v93_ = nw110_
                index116_ = default__.safeIndex(644, (d_855_v93_).length(0))
                (d_855_v93_)[index116_] = ((0) - (d_787_v45_)) * (d_787_v45_)
                pat_let_tv26_ = d_787_v45_
                index117_ = default__.safeIndex(193, (d_853_v92_).length(0))
                index118_ = default__.safeIndex(644, (d_855_v93_).length(0))
                def iife74_(_pat_let18_0):
                    def iife75_(d_858_dt__update__tmp_h1_):
                        def iife76_(_pat_let19_0):
                            def iife77_(d_859_dt__update_hcf4_h0_):
                                def iife78_(_pat_let20_0):
                                    def iife79_(d_860_dt__update_hcf2_h0_):
                                        return D1_DC2(d_860_dt__update_hcf2_h0_, (d_858_dt__update__tmp_h1_).cf3, d_859_dt__update_hcf4_h0_)
                                    return iife79_(_pat_let20_0)
                                return iife78_(pat_let_tv26_)
                            return iife77_(_pat_let19_0)
                        return iife76_((self).f17)
                    return iife75_(_pat_let18_0)
                rhs98_ = d_851_v90_
                rhs99_ = d_787_v45_
                rhs100_ = d_852_v91_
                rhs101_ = iife74_(self.f12)
                rhs102_ = d_787_v45_
                lhs89_ = d_853_v92_
                lhs90_ = default__.safeIndex(193, (d_853_v92_).length(0))
                lhs91_ = d_855_v93_
                lhs92_ = default__.safeIndex(644, (d_855_v93_).length(0))
                d_851_v90_ = rhs98_
                d_787_v45_ = rhs99_
                d_852_v91_ = rhs100_
                lhs89_[lhs90_] = rhs101_
                lhs91_[lhs92_] = rhs102_
                (globalState).f7 = 785
            d_861_v94_: _dafny.Array
            nw111_ = _dafny.Array(int(0), 18)
            d_861_v94_ = nw111_
            index119_ = default__.safeIndex(543, (d_861_v94_).length(0))
            (d_861_v94_)[index119_] = (d_787_v45_) + (d_787_v45_)
            index120_ = default__.safeIndex(543, (d_861_v94_).length(0))
            (d_861_v94_)[index120_] = (520) * (d_787_v45_)
            d_862_v95_: str
            d_862_v95_ = _dafny.CodePoint('t')
            d_863_v96_: C0
            nw112_ = C0()
            nw112_.ctor__((self).f13, d_862_v95_, self.f12, default__.fm17(globalState))
            d_863_v96_ = nw112_
            d_863_v96_ = d_863_v96_
        def iife80_(_pat_let21_0):
            def iife81_(d_864_dt__update__tmp_h2_):
                def iife82_(_pat_let22_0):
                    def iife83_(d_865_dt__update_hcf4_h1_):
                        return D1_DC2((d_864_dt__update__tmp_h2_).cf2, (d_864_dt__update__tmp_h2_).cf3, d_865_dt__update_hcf4_h1_)
                    return iife83_(_pat_let22_0)
                return iife82_((self).f13)
            return iife81_(_pat_let21_0)
        (self).f12 = iife80_(D1_DC2(len(_dafny.Map({(self).f17: (self).f13})), d_787_v45_, (self).f13))
        if not((d_787_v45_) < (d_787_v45_)):
            d_866_v97_: _dafny.Map
            d_866_v97_ = _dafny.Map({p0: (self).f18})
            d_867_v98_: _dafny.Array
            nw113_ = _dafny.Array(int(0), 23)
            d_867_v98_ = nw113_
            d_868_v99_: _dafny.MultiSet
            d_868_v99_ = _dafny.MultiSet([d_867_v98_])
            d_869_v100_: _dafny.Map
            d_869_v100_ = _dafny.Map({(d_866_v97_ if (self).f17 else _dafny.Map({(self).f13: (self).f18})): default__.safeModuloInt((d_868_v99_).cardinality, 527)})
            d_870_v101_: _dafny.Map
            d_870_v101_ = _dafny.Map({default__.fm23((self).f13, globalState): d_869_v100_})
            d_871_v102_: str
            d_871_v102_ = _dafny.CodePoint('h')
            d_872_v103_: _dafny.Map
            d_872_v103_ = _dafny.Map({d_787_v45_: d_871_v102_})
            d_873_v105_: _dafny.Map
            d_873_v105_ = _dafny.Map({d_787_v45_: (self).f17})
            def iife84_():
                coll38_ = _dafny.Map()
                compr_38_: int
                for compr_38_ in (d_873_v105_).keys.Elements:
                    d_874_v104_: int = compr_38_
                    if (d_874_v104_) in (d_873_v105_):
                        coll38_[(d_874_v104_) - (len(_dafny.SeqWithoutIsStrInference([(self).f13, (self).f17])))] = d_871_v102_
                return _dafny.Map(coll38_)
            def iife85_():
                coll39_ = _dafny.Map()
                compr_39_: int
                for compr_39_ in (d_873_v105_).keys.Elements:
                    d_875_v104_: int = compr_39_
                    if (d_875_v104_) in (d_873_v105_):
                        coll39_[(d_875_v104_) - (len(_dafny.SeqWithoutIsStrInference([(self).f13, (self).f17])))] = d_871_v102_
                return _dafny.Map(coll39_)
            def iife86_():
                coll40_ = _dafny.Map()
                compr_40_: _dafny.Map
                for compr_40_ in (d_869_v100_).keys.Elements:
                    d_876_v106_: _dafny.Map = compr_40_
                    if (d_876_v106_) in (d_869_v100_):
                        coll40_[d_876_v106_] = len(_dafny.Map({d_871_v102_: (self).f13}))
                return _dafny.Map(coll40_)
            d_869_v100_ = ((d_870_v101_)[(d_872_v103_) | (iife84_()
            )] if ((d_872_v103_) | (iife85_()
            )) in (d_870_v101_) else iife86_()
            )
            d_877_v107_: C0
            nw114_ = C0()
            nw114_.ctor__((self).f18, d_871_v102_, self.f12, (self).f17)
            d_877_v107_ = nw114_
            d_878_v108_: _dafny.Map
            d_878_v108_ = _dafny.Map({d_787_v45_: d_787_v45_})
            d_879_v109_: _dafny.Seq
            d_879_v109_ = _dafny.SeqWithoutIsStrInference([(self).f13, d_877_v107_.f19])
            rhs103_ = default__.fm17(globalState)
            rhs104_ = d_787_v45_
            rhs105_ = default__.fm24(d_787_v45_, globalState)
            rhs106_ = (default__.fm0(d_879_v109_, (self).f17, (self).f18, d_787_v45_, globalState)) + (len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_880_i9_ in range(default__.abs(892))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e")))))
            lhs93_ = d_877_v107_
            lhs93_.f19 = rhs103_
            d_787_v45_ = rhs104_
            d_878_v108_ = rhs105_
            r0 = rhs106_
            (d_877_v107_).f19 = d_877_v107_.f19
            d_873_v105_ = (d_873_v105_).set(d_787_v45_, (self).f17)
        elif True:
            index121_ = default__.safeIndex(83, ((self).f14).length(0))
            ((self).f14)[index121_] = p0
            index122_ = default__.safeIndex(83, ((self).f14).length(0))
            ((self).f14)[index122_] = p0
            d_881_v110_: str
            d_881_v110_ = _dafny.CodePoint('w')
            d_882_v111_: C0
            nw115_ = C0()
            nw115_.ctor__((not(((self).f14)[default__.safeIndex(83, ((self).f14).length(0))])) or ((self).f18), d_881_v110_, self.f12, (self).f18)
            d_882_v111_ = nw115_
            d_883_v112_: _dafny.MultiSet
            d_883_v112_ = _dafny.MultiSet([True])
            d_884_v113_: _dafny.Seq
            d_884_v113_ = _dafny.SeqWithoutIsStrInference([d_883_v112_])
            d_884_v113_ = _dafny.SeqWithoutIsStrInference([d_883_v112_ for d_885_i10_ in range(default__.abs(368))])
            d_886_v114_: _dafny.MultiSet
            d_886_v114_ = _dafny.MultiSet([d_787_v45_, d_787_v45_])
            d_887_v115_: _dafny.Seq
            d_887_v115_ = _dafny.SeqWithoutIsStrInference([(self).f17, (self).f17])
            d_888_v116_: _dafny.Map
            d_888_v116_ = _dafny.Map({False: (self).f17})
            d_889_v117_: _dafny.Set
            d_889_v117_ = _dafny.Set({default__.fm17(globalState), ((d_888_v116_)[(self).f17] if ((self).f17) in (d_888_v116_) else True)})
            d_890_v118_: _dafny.Map
            d_890_v118_ = _dafny.Map({d_882_v111_.f19: d_787_v45_})
            d_891_v119_: _dafny.Array
            nw116_ = _dafny.Array(None, 26)
            nw116_[int(0)] = d_787_v45_
            nw116_[int(1)] = (0) - (d_787_v45_)
            nw116_[int(2)] = d_787_v45_
            nw116_[int(3)] = ((d_886_v114_)[d_787_v45_] if (d_787_v45_) in (d_886_v114_) else (_dafny.MultiSet([(self).f13])).cardinality)
            nw116_[int(4)] = len(d_887_v115_)
            nw116_[int(5)] = d_787_v45_
            nw116_[int(6)] = d_787_v45_
            nw116_[int(7)] = d_787_v45_
            nw116_[int(8)] = len(d_889_v117_)
            nw116_[int(9)] = d_787_v45_
            nw116_[int(10)] = d_787_v45_
            nw116_[int(11)] = d_787_v45_
            nw116_[int(12)] = -840
            nw116_[int(13)] = d_787_v45_
            nw116_[int(14)] = len(d_890_v118_)
            nw116_[int(15)] = d_787_v45_
            nw116_[int(16)] = d_787_v45_
            nw116_[int(17)] = d_787_v45_
            nw116_[int(18)] = d_787_v45_
            nw116_[int(19)] = 787
            nw116_[int(20)] = d_787_v45_
            nw116_[int(21)] = d_787_v45_
            nw116_[int(22)] = d_787_v45_
            nw116_[int(23)] = d_787_v45_
            nw116_[int(24)] = d_787_v45_
            nw116_[int(25)] = d_787_v45_
            d_891_v119_ = nw116_
            d_892_v120_: _dafny.Array
            def lambda33_(d_893_i11_):
                return (self).f13

            init18_ = lambda33_
            nw117_ = _dafny.Array(None, 20)
            for i0_18_ in range(nw117_.length(0)):
                nw117_[i0_18_] = init18_(i0_18_)
            d_892_v120_ = nw117_
            d_894_v121_: D2
            d_894_v121_ = D2_DC4(d_891_v119_, (self).f17, (0) - (d_787_v45_), d_882_v111_.f19, d_892_v120_)
            (d_882_v111_).f19 = ((d_894_v121_).cf7 if (p0) and (d_882_v111_.f19) else True)
            d_787_v45_ = (0) - ((0) - (default__.safeDivisionInt(d_787_v45_, d_787_v45_)))
        if p0:
            d_895_v122_: _dafny.Map
            d_895_v122_ = _dafny.Map({(self).f13: (self).f13})
            r0 = (0) - ((d_787_v45_ if (self).f18 else default__.safeDivisionInt(len(d_895_v122_), d_787_v45_)))
            d_896_v123_: _dafny.Map
            d_896_v123_ = _dafny.Map({d_787_v45_: (self).f17})
            d_897_v125_: _dafny.Seq
            def iife87_():
                coll41_ = _dafny.Map()
                compr_41_: int
                for compr_41_ in _dafny.IntegerRange(627, 879):
                    d_898_v124_: int = compr_41_
                    if ((627) <= (d_898_v124_)) and ((d_898_v124_) < (879)):
                        coll41_[default__.safeDivisionInt(d_898_v124_, d_787_v45_)] = not((self).f13)
                return _dafny.Map(coll41_)
            d_897_v125_ = _dafny.SeqWithoutIsStrInference([d_896_v123_, (d_896_v123_) | (d_896_v123_), d_896_v123_, (iife87_()
            ).set(d_787_v45_, (self).f18)])
            d_899_v126_: _dafny.Seq
            d_899_v126_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fonto"))
            d_897_v125_ = default__.fm25(d_899_v126_, p0, (0) - (d_787_v45_), globalState)
            d_900_v127_: str
            d_900_v127_ = _dafny.CodePoint('a')
            pat_let_tv27_ = d_787_v45_
            pat_let_tv28_ = d_787_v45_
            pat_let_tv29_ = d_787_v45_
            d_901_v128_: C0
            nw118_ = C0()
            def iife89_(_pat_let24_0):
                def iife90_(d_902_dt__update__tmp_h3_):
                    def iife91_(_pat_let25_0):
                        def iife92_(d_903_dt__update_hcf2_h1_):
                            return D1_DC2(d_903_dt__update_hcf2_h1_, (d_902_dt__update__tmp_h3_).cf3, (d_902_dt__update__tmp_h3_).cf4)
                        return iife92_(_pat_let25_0)
                    return iife91_(pat_let_tv27_)
                return iife90_(_pat_let24_0)
            def iife88_(_pat_let23_0):
                def iife93_(d_904_dt__update__tmp_h4_):
                    def iife94_(_pat_let26_0):
                        def iife95_(d_905_dt__update_hcf3_h0_):
                            def iife96_(_pat_let27_0):
                                def iife97_(d_906_dt__update_hcf2_h2_):
                                    return D1_DC2(d_906_dt__update_hcf2_h2_, d_905_dt__update_hcf3_h0_, (d_904_dt__update__tmp_h4_).cf4)
                                return iife97_(_pat_let27_0)
                            return iife96_(-76)
                        return iife95_(_pat_let26_0)
                    return iife94_((D1_DC2(pat_let_tv28_, pat_let_tv29_, (self).f13)).cf3)
                return iife93_(_pat_let23_0)
            nw118_.ctor__(p0, d_900_v127_, iife88_(iife89_(self.f12)), (self).f13)
            d_901_v128_ = nw118_
            d_907_v129_: _dafny.Seq
            d_907_v129_ = _dafny.SeqWithoutIsStrInference([not(p0)])
            d_908_v130_: _dafny.Array
            nw119_ = _dafny.Array(None, 18)
            nw119_[int(0)] = (d_787_v45_ if p0 else len(d_907_v129_))
            nw119_[int(1)] = 923
            nw119_[int(2)] = len(d_899_v126_)
            nw119_[int(3)] = 581
            nw119_[int(4)] = default__.safeModuloInt(d_787_v45_, d_787_v45_)
            nw119_[int(5)] = d_787_v45_
            nw119_[int(6)] = d_787_v45_
            nw119_[int(7)] = len(default__.fm26(d_787_v45_, d_787_v45_, p0, globalState))
            nw119_[int(8)] = (d_787_v45_ if p0 else d_787_v45_)
            nw119_[int(9)] = d_787_v45_
            nw119_[int(10)] = d_787_v45_
            nw119_[int(11)] = d_787_v45_
            nw119_[int(12)] = d_787_v45_
            nw119_[int(13)] = d_787_v45_
            nw119_[int(14)] = d_787_v45_
            nw119_[int(15)] = d_787_v45_
            nw119_[int(16)] = (d_787_v45_) * (d_787_v45_)
            nw119_[int(17)] = d_787_v45_
            d_908_v130_ = nw119_
            index123_ = default__.safeIndex(639, (d_908_v130_).length(0))
            (d_908_v130_)[index123_] = d_787_v45_
            index124_ = default__.safeIndex(639, (d_908_v130_).length(0))
            (d_908_v130_)[index124_] = d_787_v45_
            d_909_v131_: _dafny.Array
            nw120_ = _dafny.Array(_dafny.Set({}), 24)
            d_909_v131_ = nw120_
            d_910_v132_: _dafny.Map
            d_910_v132_ = _dafny.Map({d_787_v45_: (d_908_v130_)[default__.safeIndex(639, (d_908_v130_).length(0))]})
            d_911_v133_: _dafny.Seq
            d_911_v133_ = _dafny.SeqWithoutIsStrInference([len(d_910_v132_)])
            d_912_v134_: _dafny.Set
            d_912_v134_ = _dafny.Set({d_911_v133_})
            index125_ = default__.safeIndex(702, (d_909_v131_).length(0))
            (d_909_v131_)[index125_] = (d_912_v134_) | (d_912_v134_)
            index126_ = default__.safeIndex(702, (d_909_v131_).length(0))
            (d_909_v131_)[index126_] = d_912_v134_
        elif True:
            d_913_v135_: _dafny.Array
            nw121_ = _dafny.Array(_dafny.Array(None, 0), 28)
            d_913_v135_ = nw121_
            d_914_v136_: _dafny.Array
            nw122_ = _dafny.Array(None, 1)
            nw122_[int(0)] = self.f12
            d_914_v136_ = nw122_
            index127_ = default__.safeIndex(324, (d_913_v135_).length(0))
            (d_913_v135_)[index127_] = d_914_v136_
            index128_ = default__.safeIndex(324, (d_913_v135_).length(0))
            (d_913_v135_)[index128_] = d_914_v136_
            d_915_v137_: D6
            d_915_v137_ = D6_DC12(p0, default__.fm20(globalState), d_787_v45_, (self).f13)
            source16_ = d_915_v137_
            if source16_.is_DC12:
                d_916___mcc_h6_ = source16_.cf18
                d_917___mcc_h7_ = source16_.cf19
                d_918___mcc_h8_ = source16_.cf20
                d_919___mcc_h9_ = source16_.cf21
                d_920_cf21_ = d_919___mcc_h9_
                d_921_cf20_ = d_918___mcc_h8_
                d_922_cf19_ = d_917___mcc_h7_
                d_923_cf18_ = d_916___mcc_h6_
                d_924_v138_: _dafny.Array
                def lambda34_(d_925_cf19_):
                    def lambda35_(d_926_i12_):
                        return d_925_cf19_

                    return lambda35_

                init19_ = lambda34_(d_922_cf19_)
                nw123_ = _dafny.Array(None, 3)
                for i0_19_ in range(nw123_.length(0)):
                    nw123_[i0_19_] = init19_(i0_19_)
                d_924_v138_ = nw123_
                index129_ = default__.safeIndex(804, (d_924_v138_).length(0))
                (d_924_v138_)[index129_] = d_922_cf19_
                index130_ = default__.safeIndex(804, (d_924_v138_).length(0))
                (d_924_v138_)[index130_] = d_922_cf19_
                (self).m12(globalState)
                d_927_v139_: str
                d_927_v139_ = _dafny.CodePoint('b')
                d_928_v140_: C0
                nw124_ = C0()
                nw124_.ctor__((self).f18, d_927_v139_, self.f12, (d_787_v45_) < (d_787_v45_))
                d_928_v140_ = nw124_
                d_929_v141_: C0
                nw125_ = C0()
                nw125_.ctor__(d_923_cf18_, d_927_v139_, self.f12, d_923_cf18_)
                d_929_v141_ = nw125_
            elif source16_.is_DC13:
                d_930___mcc_h10_ = source16_.cf22
                d_931___mcc_h11_ = source16_.cf23
                d_932___mcc_h12_ = source16_.cf24
                d_933___mcc_h13_ = source16_.cf25
                d_934___mcc_h14_ = source16_.cf26
                d_935_cf26_ = d_934___mcc_h14_
                d_936_cf25_ = d_933___mcc_h13_
                d_937_cf24_ = d_932___mcc_h12_
                d_938_cf23_ = d_931___mcc_h11_
                d_939_cf22_ = d_930___mcc_h10_
                d_940_v142_: _dafny.Array
                nw126_ = _dafny.Array(False, 12)
                d_940_v142_ = nw126_
                d_940_v142_ = d_940_v142_
                d_941_v143_: _dafny.Seq
                d_941_v143_ = _dafny.SeqWithoutIsStrInference([d_937_cf24_])
                d_941_v143_ = (_dafny.SeqWithoutIsStrInference([d_787_v45_, d_936_cf25_])) + (d_941_v143_)
                d_942_v144_: str
                d_942_v144_ = _dafny.CodePoint('t')
                d_943_v145_: _dafny.Seq
                d_943_v145_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tpil"))
                d_935_cf26_ = (d_942_v144_) not in ((d_943_v145_) + (d_943_v145_))
                d_935_cf26_ = ((self).f18 if not((self).f17) else (d_787_v45_) <= (d_787_v45_))
            elif source16_.is_DC11:
                d_944___mcc_h15_ = source16_.cf17
                d_945_cf17_ = d_944___mcc_h15_
                d_946_v146_: _dafny.Array
                nw127_ = _dafny.Array(_dafny.Array(None, 0), 28)
                d_946_v146_ = nw127_
                d_947_v147_: _dafny.Array
                nw128_ = _dafny.Array(int(0), 22)
                d_947_v147_ = nw128_
                index131_ = default__.safeIndex(90, (d_946_v146_).length(0))
                (d_946_v146_)[index131_] = d_947_v147_
                index132_ = default__.safeIndex(90, (d_946_v146_).length(0))
                (d_946_v146_)[index132_] = d_947_v147_
                index133_ = default__.safeIndex(933, (d_947_v147_).length(0))
                (d_947_v147_)[index133_] = d_787_v45_
                d_948_v148_: _dafny.Array
                nw129_ = _dafny.Array(D6.default()(), 20)
                d_948_v148_ = nw129_
                d_949_v149_: _dafny.Seq
                d_949_v149_ = _dafny.SeqWithoutIsStrInference([d_948_v148_])
                d_950_v150_: _dafny.Seq
                d_950_v150_ = _dafny.SeqWithoutIsStrInference([d_949_v149_, d_949_v149_, d_949_v149_])
                index134_ = default__.safeIndex(933, (d_947_v147_).length(0))
                (d_947_v147_)[index134_] = default__.safeModuloInt(default__.safeModuloInt(d_787_v45_, len((d_950_v150_)[default__.safeIndex(d_787_v45_, len(d_950_v150_))])), d_787_v45_)
                d_951_v151_: _dafny.Map
                d_951_v151_ = _dafny.Map({p0: p0})
                d_952_v152_: _dafny.Map
                d_952_v152_ = _dafny.Map({True: (d_947_v147_)[default__.safeIndex(933, (d_947_v147_).length(0))]})
                index135_ = default__.safeIndex(711, ((self).f14).length(0))
                ((self).f14)[index135_] = ((0) - (len(d_951_v151_))) > (((d_952_v152_)[(self).f18] if ((self).f18) in (d_952_v152_) else d_787_v45_))
                index136_ = default__.safeIndex(711, ((self).f14).length(0))
                ((self).f14)[index136_] = p0
                d_953_v153_: _dafny.Seq
                d_953_v153_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))
                d_954_v154_: _dafny.Map
                d_954_v154_ = _dafny.Map({d_953_v153_: d_947_v147_})
                (globalState).f7 = len((d_954_v154_).set((d_953_v153_) + (d_953_v153_), (d_946_v146_)[default__.safeIndex(90, (d_946_v146_).length(0))]))
            elif True:
                d_955___mcc_h16_ = source16_.cf27
                d_956_cf27_ = d_955___mcc_h16_
                (globalState).f7 = d_787_v45_
                d_957_v155_: _dafny.Seq
                d_957_v155_ = _dafny.SeqWithoutIsStrInference([(self).f14, (self).f14])
                d_958_v156_: _dafny.Seq
                d_958_v156_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xwoa"))
                d_959_v157_: _dafny.Map
                d_959_v157_ = _dafny.Map({(d_957_v155_)[default__.safeIndex(len(d_958_v156_), len(d_957_v155_))]: _dafny.SeqWithoutIsStrInference([(0) - (d_787_v45_), d_787_v45_])})
                d_959_v157_ = (d_959_v157_).set((self).f14, _dafny.SeqWithoutIsStrInference([447 for d_960_i13_ in range(default__.abs(315))]))
                d_961_v158_: bool
                d_961_v158_ = True
                d_962_v159_: _dafny.Array
                def lambda36_(d_963_v45_):
                    def lambda37_(d_964_i14_):
                        return _dafny.SeqWithoutIsStrInference([d_963_v45_])

                    return lambda37_

                init20_ = lambda36_(d_787_v45_)
                nw130_ = _dafny.Array(None, 26)
                for i0_20_ in range(nw130_.length(0)):
                    nw130_[i0_20_] = init20_(i0_20_)
                d_962_v159_ = nw130_
                d_965_v160_: _dafny.Seq
                d_965_v160_ = _dafny.SeqWithoutIsStrInference([d_787_v45_])
                index137_ = default__.safeIndex(565, (d_962_v159_).length(0))
                (d_962_v159_)[index137_] = d_965_v160_
                d_966_v161_: _dafny.Seq
                d_966_v161_ = _dafny.SeqWithoutIsStrInference([default__.fm17(globalState)])
                d_967_v162_: _dafny.Map
                d_967_v162_ = _dafny.Map({d_787_v45_: p0})
                d_968_v163_: _dafny.MultiSet
                d_968_v163_ = _dafny.MultiSet([(d_966_v161_).set(default__.safeIndex(len(d_967_v162_), len(d_966_v161_)), True)])
                index138_ = default__.safeIndex(565, (d_962_v159_).length(0))
                rhs107_ = ((self).f17) and (not(p0))
                rhs108_ = (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_787_v45_: p0})), (d_968_v163_).cardinality, d_787_v45_, d_787_v45_])) + ((default__.fm27(globalState)).set(default__.safeIndex(d_787_v45_, len(default__.fm27(globalState))), d_787_v45_))
                rhs109_ = ((d_787_v45_) - (d_787_v45_)) - (default__.safeDivisionInt(d_787_v45_, d_787_v45_))
                lhs94_ = d_962_v159_
                lhs95_ = default__.safeIndex(565, (d_962_v159_).length(0))
                d_961_v158_ = rhs107_
                lhs94_[lhs95_] = rhs108_
                d_787_v45_ = rhs109_
                d_969_v164_: C0
                nw131_ = C0()
                nw131_.ctor__(p0, _dafny.CodePoint('n'), self.f12, d_961_v158_)
                d_969_v164_ = nw131_
                d_970_v165_: _dafny.Map
                d_970_v165_ = _dafny.Map({d_969_v164_: d_958_v156_})
                d_970_v165_ = (d_970_v165_).set(d_969_v164_, (d_958_v156_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))))
            d_971_v166_: _dafny.MultiSet
            d_971_v166_ = _dafny.MultiSet([(self).f13, (self).f18])
            index139_ = default__.safeIndex(413, ((self).f14).length(0))
            ((self).f14)[index139_] = (self).f17
            d_972_v167_: bool
            d_972_v167_ = False
            d_973_v168_: str
            d_973_v168_ = _dafny.CodePoint('b')
            d_974_v169_: _dafny.Map
            d_974_v169_ = _dafny.Map({d_973_v168_: (self).f18})
            index140_ = default__.safeIndex(413, ((self).f14).length(0))
            rhs110_ = default__.fm28(((d_974_v169_)[d_973_v168_] if (d_973_v168_) in (d_974_v169_) else (self).f18), (0) - (default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ybya"))), d_787_v45_)), not ((self).f17) or (not(p0)), globalState)
            rhs111_ = (self).f13
            rhs112_ = default__.fm17(globalState)
            rhs113_ = ((d_787_v45_) + (d_787_v45_)) < (d_787_v45_)
            rhs114_ = (self).f13
            lhs96_ = (self).f14
            lhs97_ = default__.safeIndex(413, ((self).f14).length(0))
            d_971_v166_ = rhs110_
            lhs96_[lhs97_] = rhs111_
            d_972_v167_ = rhs112_
            d_972_v167_ = rhs113_
            d_972_v167_ = rhs114_
            if ((self).f14)[default__.safeIndex(413, ((self).f14).length(0))]:
                d_975_v170_: _dafny.Array
                def lambda38_(d_976_v45_):
                    def lambda39_(d_977_i15_):
                        return (d_977_i15_) + (d_976_v45_)

                    return lambda39_

                init21_ = lambda38_(d_787_v45_)
                nw132_ = _dafny.Array(None, 25)
                for i0_21_ in range(nw132_.length(0)):
                    nw132_[i0_21_] = init21_(i0_21_)
                d_975_v170_ = nw132_
                index141_ = default__.safeIndex(362, (d_975_v170_).length(0))
                (d_975_v170_)[index141_] = ((self).fm15(globalState)) - ((self).fm9(d_973_v168_, globalState))
                index142_ = default__.safeIndex(362, (d_975_v170_).length(0))
                (d_975_v170_)[index142_] = d_787_v45_
                index143_ = default__.safeIndex(413, ((self).f14).length(0))
                ((self).f14)[index143_] = d_972_v167_
                d_978_v171_: _dafny.Seq
                d_978_v171_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsq"))
                d_978_v171_ = (d_978_v171_).set(default__.safeIndex(d_787_v45_, len(d_978_v171_)), d_973_v168_)
                d_979_v172_: D6
                d_979_v172_ = D6_DC13(p0, p0, d_787_v45_, d_787_v45_, default__.fm17(globalState))
                index144_ = default__.safeIndex(413, ((self).f14).length(0))
                rhs115_ = (d_972_v167_ if p0 else ((d_975_v170_)[default__.safeIndex(362, (d_975_v170_).length(0))]) < (d_787_v45_))
                rhs116_ = ((0) - ((0) - ((len(_dafny.Set({d_975_v170_, d_975_v170_}))) * (d_787_v45_)))) + ((d_979_v172_).cf25)
                lhs98_ = (self).f14
                lhs99_ = default__.safeIndex(413, ((self).f14).length(0))
                lhs100_ = globalState
                lhs98_[lhs99_] = rhs115_
                lhs100_.f7 = rhs116_
                d_980_v173_: _dafny.Array
                nw133_ = _dafny.Array(False, 23)
                d_980_v173_ = nw133_
                d_981_v174_: _dafny.Map
                d_981_v174_ = _dafny.Map({d_972_v167_: d_980_v173_})
                d_981_v174_ = ((_dafny.Map({(self).f13: d_980_v173_})) | (d_981_v174_)) | (d_981_v174_)
            elif True:
                d_982_v175_: _dafny.Seq
                d_982_v175_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cmyb"))
                pat_let_tv30_ = d_982_v175_
                d_983_v176_: _dafny.Map
                def iife98_(_pat_let28_0):
                    def iife99_(d_984_dt__update__tmp_h5_):
                        def iife100_(_pat_let29_0):
                            def iife101_(d_985_dt__update_hcf19_h0_):
                                def iife102_(_pat_let30_0):
                                    def iife103_(d_986_dt__update_hcf21_h0_):
                                        return D6_DC12((d_984_dt__update__tmp_h5_).cf18, d_985_dt__update_hcf19_h0_, (d_984_dt__update__tmp_h5_).cf20, d_986_dt__update_hcf21_h0_)
                                    return iife103_(_pat_let30_0)
                                return iife102_((self).f17)
                            return iife101_(_pat_let29_0)
                        return iife100_(pat_let_tv30_)
                    return iife99_(_pat_let28_0)
                d_983_v176_ = _dafny.Map({default__.fm29((iife98_(d_915_v137_)).cf20, globalState): d_787_v45_})
                d_987_v177_: _dafny.Seq
                d_987_v177_ = _dafny.SeqWithoutIsStrInference([d_972_v167_, not(p0), True, not(p0)])
                d_983_v176_ = (d_983_v176_).set(d_987_v177_, d_787_v45_)
                d_988_v178_: _dafny.Map
                d_988_v178_ = _dafny.Map({d_972_v167_: d_982_v175_})
                d_989_v179_: _dafny.Array
                nw134_ = _dafny.Array(None, 13)
                nw134_[int(0)] = (d_982_v175_) + (d_982_v175_)
                nw134_[int(1)] = d_982_v175_
                nw134_[int(2)] = (d_982_v175_) + (d_982_v175_)
                nw134_[int(3)] = d_982_v175_
                nw134_[int(4)] = d_982_v175_
                nw134_[int(5)] = d_982_v175_
                nw134_[int(6)] = d_982_v175_
                nw134_[int(7)] = d_982_v175_
                nw134_[int(8)] = ((d_988_v178_)[p0] if (p0) in (d_988_v178_) else _dafny.SeqWithoutIsStrInference([d_973_v168_ for d_990_i16_ in range(default__.abs(44))]))
                nw134_[int(9)] = (d_982_v175_) + (_dafny.SeqWithoutIsStrInference([d_973_v168_ for d_991_i17_ in range(default__.abs(104))]))
                nw134_[int(10)] = (d_982_v175_ if (self).f13 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "plvqc")))
                nw134_[int(11)] = d_982_v175_
                nw134_[int(12)] = (d_982_v175_) + (_dafny.SeqWithoutIsStrInference([d_973_v168_ for d_992_i18_ in range(default__.abs(-228))]))
                d_989_v179_ = nw134_
                index145_ = default__.safeIndex(367, (d_989_v179_).length(0))
                (d_989_v179_)[index145_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yhhsck"))
                index146_ = default__.safeIndex(367, (d_989_v179_).length(0))
                rhs117_ = d_973_v168_
                rhs118_ = d_982_v175_
                rhs119_ = d_982_v175_
                lhs101_ = d_989_v179_
                lhs102_ = default__.safeIndex(367, (d_989_v179_).length(0))
                d_973_v168_ = rhs117_
                d_982_v175_ = rhs118_
                lhs101_[lhs102_] = rhs119_
                (globalState).f7 = default__.safeModuloInt((len((d_989_v179_)[default__.safeIndex(367, (d_989_v179_).length(0))])) * (d_787_v45_), d_787_v45_)
                d_972_v167_ = d_972_v167_
                d_972_v167_ = False
            d_993_v180_: _dafny.MultiSet
            d_993_v180_ = _dafny.MultiSet([d_787_v45_, 882, d_787_v45_, d_787_v45_])
            d_994_v181_: _dafny.Seq
            d_994_v181_ = _dafny.SeqWithoutIsStrInference([d_787_v45_, ((d_993_v180_)[d_787_v45_] if (d_787_v45_) in (d_993_v180_) else d_787_v45_), d_787_v45_])
            d_994_v181_ = (default__.fm27(globalState)).set(default__.safeIndex((d_994_v181_)[default__.safeIndex(d_787_v45_, len(d_994_v181_))], len(default__.fm27(globalState))), d_787_v45_)
        d_995_v182_: _dafny.Seq
        d_995_v182_ = _dafny.SeqWithoutIsStrInference([(self).f17])
        r0 = default__.fm0(d_995_v182_, (self).f13, (self).f13, d_787_v45_, globalState)
        return r0

    def m7(self, globalState):
        if (self).f17:
            d_996_v0_: int
            d_996_v0_ = 554
            d_997_v1_: D6
            d_997_v1_ = D6_DC13(not((self).f17), not((272) >= (383)), d_996_v0_, d_996_v0_, (self).f17)
            d_997_v1_ = d_997_v1_
            d_998_v2_: _dafny.Seq
            d_998_v2_ = _dafny.SeqWithoutIsStrInference([len((self).fm10((self.f12).cf4, default__.fm17(globalState), (_dafny.SeqWithoutIsStrInference([(self).f13, (self).f13])).set(default__.safeIndex(d_996_v0_, len(_dafny.SeqWithoutIsStrInference([(self).f13, (self).f13]))), (self).f13), globalState))])
            d_999_v3_: _dafny.Seq
            d_999_v3_ = _dafny.SeqWithoutIsStrInference([(self).f18])
            d_1000_v4_: _dafny.Seq
            d_1000_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xdavfw"))
            d_1001_v5_: str
            d_1001_v5_ = _dafny.CodePoint('e')
            d_1002_v6_: _dafny.Map
            d_1002_v6_ = _dafny.Map({d_1000_v4_: d_1001_v5_})
            d_1003_v7_: _dafny.Array
            nw135_ = _dafny.Array(None, 17)
            nw135_[int(0)] = (d_996_v0_ if (self).f17 else d_996_v0_)
            nw135_[int(1)] = d_996_v0_
            nw135_[int(2)] = d_996_v0_
            nw135_[int(3)] = d_996_v0_
            nw135_[int(4)] = d_996_v0_
            nw135_[int(5)] = len(d_998_v2_)
            nw135_[int(6)] = d_996_v0_
            nw135_[int(7)] = (d_996_v0_) + (d_996_v0_)
            nw135_[int(8)] = len(d_999_v3_)
            nw135_[int(9)] = (d_998_v2_)[default__.safeIndex(d_996_v0_, len(d_998_v2_))]
            nw135_[int(10)] = d_996_v0_
            nw135_[int(11)] = d_996_v0_
            nw135_[int(12)] = -137
            nw135_[int(13)] = len(d_1002_v6_)
            nw135_[int(14)] = d_996_v0_
            nw135_[int(15)] = 440
            nw135_[int(16)] = (d_996_v0_) - (d_996_v0_)
            d_1003_v7_ = nw135_
            index147_ = default__.safeIndex(674, (d_1003_v7_).length(0))
            (d_1003_v7_)[index147_] = len((D1_DC1(d_1000_v4_)).cf1)
            index148_ = default__.safeIndex(594, (d_1003_v7_).length(0))
            (d_1003_v7_)[index148_] = d_996_v0_
            d_1004_v8_: _dafny.MultiSet
            d_1004_v8_ = _dafny.MultiSet([d_1001_v5_, d_1001_v5_, d_1001_v5_])
            index149_ = default__.safeIndex(674, (d_1003_v7_).length(0))
            index150_ = default__.safeIndex(594, (d_1003_v7_).length(0))
            rhs120_ = d_996_v0_
            rhs121_ = (len((d_1000_v4_) + (d_1000_v4_))) + (d_996_v0_)
            rhs122_ = (_dafny.MultiSet([d_1001_v5_])) - (_dafny.MultiSet((d_1000_v4_) + (_dafny.SeqWithoutIsStrInference([d_1001_v5_, d_1001_v5_, d_1001_v5_]))))
            lhs103_ = d_1003_v7_
            lhs104_ = default__.safeIndex(674, (d_1003_v7_).length(0))
            lhs105_ = d_1003_v7_
            lhs106_ = default__.safeIndex(594, (d_1003_v7_).length(0))
            lhs103_[lhs104_] = rhs120_
            lhs105_[lhs106_] = rhs121_
            d_1004_v8_ = rhs122_
            (globalState).f7 = (d_1003_v7_)[default__.safeIndex(674, (d_1003_v7_).length(0))]
            d_1005_v9_: bool
            d_1005_v9_ = True
            d_1005_v9_ = (self).f13
            d_1006_v10_: _dafny.Map
            d_1006_v10_ = _dafny.Map({(d_1003_v7_)[default__.safeIndex(674, (d_1003_v7_).length(0))]: len(d_1000_v4_)})
            d_1007_v11_: _dafny.Seq
            d_1007_v11_ = _dafny.SeqWithoutIsStrInference([d_1000_v4_, d_1000_v4_])
            d_1006_v10_ = (d_1006_v10_).set(-212, len(d_1007_v11_))
        elif True:
            d_1008_v12_: _dafny.Seq
            d_1008_v12_ = _dafny.SeqWithoutIsStrInference([(self).f13])
            d_1009_v13_: _dafny.Array
            nw136_ = _dafny.Array(None, 3)
            nw136_[int(0)] = ((self).f17) == ((self).f18)
            nw136_[int(1)] = (self).f13
            nw136_[int(2)] = (d_1008_v12_)[default__.safeIndex(562, len(d_1008_v12_))]
            d_1009_v13_ = nw136_
            d_1010_v14_: bool
            d_1010_v14_ = False
            d_1011_v15_: _dafny.Array
            def lambda40_(d_1012_i0_):
                return _dafny.CodePoint('w')

            init22_ = lambda40_
            nw137_ = _dafny.Array(None, 23)
            for i0_22_ in range(nw137_.length(0)):
                nw137_[i0_22_] = init22_(i0_22_)
            d_1011_v15_ = nw137_
            d_1013_v16_: str
            d_1013_v16_ = _dafny.CodePoint('c')
            d_1014_v17_: _dafny.Seq
            d_1014_v17_ = _dafny.SeqWithoutIsStrInference([d_1013_v16_, d_1013_v16_])
            d_1015_v18_: int
            d_1015_v18_ = 677
            index151_ = default__.safeIndex(921, (d_1011_v15_).length(0))
            (d_1011_v15_)[index151_] = (d_1014_v17_)[default__.safeIndex(d_1015_v18_, len(d_1014_v17_))]
            d_1016_v19_: _dafny.Map
            d_1016_v19_ = _dafny.Map({709: not((not(d_1010_v14_)) or ((self).f18))})
            index152_ = default__.safeIndex(921, (d_1011_v15_).length(0))
            rhs123_ = (_dafny.SeqWithoutIsStrInference([d_1010_v14_])) + ((_dafny.SeqWithoutIsStrInference([d_1010_v14_])) + (default__.fm29(d_1015_v18_, globalState)))
            rhs124_ = d_1009_v13_
            rhs125_ = ((d_1016_v19_)[d_1015_v18_] if (d_1015_v18_) in (d_1016_v19_) else (self).f17)
            rhs126_ = d_1013_v16_
            lhs107_ = globalState
            lhs108_ = d_1011_v15_
            lhs109_ = default__.safeIndex(921, (d_1011_v15_).length(0))
            lhs107_.f2 = rhs123_
            d_1009_v13_ = rhs124_
            d_1010_v14_ = rhs125_
            lhs108_[lhs109_] = rhs126_
            d_1017_v20_: _dafny.Set
            d_1017_v20_ = _dafny.Set({(self).f13})
            (globalState).f7 = (default__.safeModuloInt((self).fm9(d_1013_v16_, globalState), 757)) * (len(d_1017_v20_))
            d_1018_v21_: _dafny.MultiSet
            d_1018_v21_ = _dafny.MultiSet([d_1015_v18_])
            d_1019_v22_: _dafny.Seq
            d_1019_v22_ = _dafny.SeqWithoutIsStrInference([d_1015_v18_])
            d_1020_v23_: _dafny.Array
            nw138_ = _dafny.Array(None, 6)
            nw138_[int(0)] = len(_dafny.SeqWithoutIsStrInference([d_1013_v16_ for d_1021_i1_ in range(default__.abs(733))]))
            nw138_[int(1)] = len(d_1014_v17_)
            nw138_[int(2)] = d_1015_v18_
            nw138_[int(3)] = d_1015_v18_
            nw138_[int(4)] = (d_1018_v21_).cardinality
            nw138_[int(5)] = (d_1019_v22_)[default__.safeIndex(len(d_1014_v17_), len(d_1019_v22_))]
            d_1020_v23_ = nw138_
            d_1022_v24_: D2
            d_1022_v24_ = D2_DC4(d_1020_v23_, default__.fm17(globalState), (d_1015_v18_) + ((d_1018_v21_).cardinality), (d_1015_v18_) in (_dafny.Map({d_1015_v18_: d_1015_v18_})), d_1009_v13_)
            source17_ = d_1022_v24_
            if source17_.is_DC4:
                d_1023___mcc_h0_ = source17_.cf6
                d_1024___mcc_h1_ = source17_.cf7
                d_1025___mcc_h2_ = source17_.cf8
                d_1026___mcc_h3_ = source17_.cf9
                d_1027___mcc_h4_ = source17_.cf10
                d_1028_cf10_ = d_1027___mcc_h4_
                d_1029_cf9_ = d_1026___mcc_h3_
                d_1030_cf8_ = d_1025___mcc_h2_
                d_1031_cf7_ = d_1024___mcc_h1_
                d_1032_cf6_ = d_1023___mcc_h0_
                (globalState).f7 = d_1015_v18_
                d_1033_v25_: _dafny.Array
                nw139_ = _dafny.Array(_dafny.Array(None, 0), 24)
                d_1033_v25_ = nw139_
                index153_ = default__.safeIndex(612, (d_1033_v25_).length(0))
                (d_1033_v25_)[index153_] = d_1009_v13_
                index154_ = default__.safeIndex(612, (d_1033_v25_).length(0))
                (d_1033_v25_)[index154_] = d_1028_cf10_
                d_1034_v26_: _dafny.Seq
                d_1034_v26_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({_dafny.SeqWithoutIsStrInference([(d_1011_v15_)[default__.safeIndex(921, (d_1011_v15_).length(0))] for d_1035_i2_ in range(default__.abs(288))])})])
                d_1036_v27_: D3
                d_1036_v27_ = D3_DC7()
                d_1037_v28_: _dafny.Set
                d_1037_v28_ = _dafny.Set({d_1036_v27_})
                rhs127_ = d_1015_v18_
                rhs128_ = ((d_1015_v18_) + (len(d_1034_v26_))) + (d_1015_v18_)
                rhs129_ = len(d_1037_v28_)
                lhs110_ = globalState
                d_1030_cf8_ = rhs127_
                d_1030_cf8_ = rhs128_
                lhs110_.f7 = rhs129_
                (globalState).f7 = d_1015_v18_
            elif True:
                d_1038___mcc_h5_ = source17_.cf5
                d_1039_cf5_ = d_1038___mcc_h5_
                (self).m11(d_1009_v13_, d_1010_v14_, (len(d_1008_v12_)) >= (d_1015_v18_), globalState)
                d_1039_cf5_ = (self).f14
                (globalState).f7 = (d_1015_v18_) + (default__.safeModuloInt((0) - (d_1015_v18_), d_1015_v18_))
                d_1039_cf5_ = (self).f14
            d_1040_v29_: _dafny.Array
            nw140_ = _dafny.Array(None, 14)
            d_1040_v29_ = nw140_
            index155_ = default__.safeIndex(723, (d_1020_v23_).length(0))
            (d_1020_v23_)[index155_] = d_1015_v18_
            index156_ = default__.safeIndex(723, (d_1020_v23_).length(0))
            rhs130_ = d_1040_v29_
            rhs131_ = (self).f14
            rhs132_ = d_1015_v18_
            rhs133_ = (d_1019_v22_)[default__.safeIndex(default__.safeDivisionInt(d_1015_v18_, d_1015_v18_), len(d_1019_v22_))]
            lhs111_ = d_1020_v23_
            lhs112_ = default__.safeIndex(723, (d_1020_v23_).length(0))
            lhs113_ = globalState
            d_1040_v29_ = rhs130_
            d_1009_v13_ = rhs131_
            lhs111_[lhs112_] = rhs132_
            lhs113_.f7 = rhs133_
            d_1010_v14_ = (self).f18
        d_1041_v30_: bool
        d_1041_v30_ = False
        d_1041_v30_ = True
        d_1042_v31_: D6
        d_1042_v31_ = D6_DC11(default__.fm19(globalState))
        d_1042_v31_ = d_1042_v31_
        source18_ = self.f12
        if source18_.is_DC2:
            d_1043___mcc_h6_ = source18_.cf2
            d_1044___mcc_h7_ = source18_.cf3
            d_1045___mcc_h8_ = source18_.cf4
            d_1046_cf4_ = d_1045___mcc_h8_
            d_1047_cf3_ = d_1044___mcc_h7_
            d_1048_cf2_ = d_1043___mcc_h6_
            d_1049_v32_: str
            d_1049_v32_ = _dafny.CodePoint('f')
            d_1049_v32_ = d_1049_v32_
            d_1050_v33_: _dafny.Seq
            d_1050_v33_ = _dafny.SeqWithoutIsStrInference([d_1047_cf3_])
            d_1041_v30_ = ((d_1050_v33_) + (d_1050_v33_)) != (d_1050_v33_)
            d_1051_v34_: T3
            nw141_ = C2()
            nw141_.ctor__(-855, d_1046_cf4_, self.f16, d_1046_cf4_, (self).f14, (self).f15, self.f12, d_1041_v30_)
            d_1051_v34_ = nw141_
            d_1052_v35_: _dafny.Array
            def lambda41_(d_1053_cf2_):
                def lambda42_(d_1054_i3_):
                    return default__.safeDivisionInt(d_1054_i3_, d_1053_cf2_)

                return lambda42_

            init23_ = lambda41_(d_1048_cf2_)
            nw142_ = _dafny.Array(None, 21)
            for i0_23_ in range(nw142_.length(0)):
                nw142_[i0_23_] = init23_(i0_23_)
            d_1052_v35_ = nw142_
            d_1055_v36_: _dafny.Map
            d_1055_v36_ = _dafny.Map({(d_1051_v34_ if (self).f17 else d_1051_v34_): d_1052_v35_})
            (globalState).f7 = (0) - (len(d_1055_v36_))
            d_1056_v37_: _dafny.Seq
            d_1056_v37_ = _dafny.SeqWithoutIsStrInference([((self).f13 if d_1046_cf4_ else d_1046_cf4_)])
            d_1041_v30_ = not(not((d_1056_v37_)[default__.safeIndex(d_1048_cf2_, len(d_1056_v37_))]))
        elif True:
            d_1057___mcc_h9_ = source18_.cf1
            d_1058_cf1_ = d_1057___mcc_h9_
            d_1041_v30_ = (d_1041_v30_ if not(False) else default__.fm30(38, globalState))
            d_1059_v38_: str
            d_1059_v38_ = _dafny.CodePoint('v')
            d_1060_v39_: C0
            nw143_ = C0()
            nw143_.ctor__((self).f17, d_1059_v38_, self.f12, ((self).f13 if (self).f13 else d_1041_v30_))
            d_1060_v39_ = nw143_
            d_1061_v40_: int
            d_1061_v40_ = 872
            d_1062_v41_: _dafny.Set
            d_1062_v41_ = _dafny.Set({(self).f13, not(not((self).f13))})
            d_1063_v42_: C2
            nw144_ = C2()
            nw144_.ctor__((0) - (d_1061_v40_), (_dafny.Set({(self).f17, not(False)})).issubset(d_1062_v41_), D1_DC1(d_1058_cf1_), (d_1058_cf1_) <= (d_1058_cf1_), (self).f14, (self).f15, self.f12, d_1041_v30_)
            d_1063_v42_ = nw144_
            d_1064_v43_: _dafny.Map
            d_1064_v43_ = _dafny.Map({(d_1063_v42_).f21: d_1041_v30_})
            d_1065_v44_: _dafny.Map
            d_1065_v44_ = _dafny.Map({d_1064_v43_: (d_1061_v40_) + (d_1061_v40_)})
            d_1066_v45_: _dafny.Map
            d_1066_v45_ = _dafny.Map({(self).f18: (self).f13})
            d_1065_v44_ = (d_1065_v44_).set(d_1064_v43_, (len(d_1066_v45_)) - (d_1061_v40_))
        d_1067_v46_: int
        d_1067_v46_ = -61
        d_1068_v47_: _dafny.Seq
        d_1068_v47_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hu"))
        d_1069_v48_: _dafny.Map
        d_1069_v48_ = _dafny.Map({d_1067_v46_: d_1068_v47_})
        hi5_ = default__.safeModuloInt(len(d_1069_v48_), d_1067_v46_)
        for d_1070_i4_ in range((0) - (d_1067_v46_), hi5_):
            d_1071_v49_: _dafny.MultiSet
            d_1071_v49_ = _dafny.MultiSet([(self).f14])
            d_1072_v50_: _dafny.Map
            d_1072_v50_ = _dafny.Map({d_1071_v49_: (self).f14})
            d_1072_v50_ = (d_1072_v50_).set(d_1071_v49_, (self).f14)
            d_1073_v51_: _dafny.Array
            def lambda43_(d_1074_i5_):
                return _dafny.Map({(self).f17: 323})

            init24_ = lambda43_
            nw145_ = _dafny.Array(None, 3)
            for i0_24_ in range(nw145_.length(0)):
                nw145_[i0_24_] = init24_(i0_24_)
            d_1073_v51_ = nw145_
            d_1075_v53_: _dafny.Map
            def iife104_():
                coll42_ = _dafny.Map()
                compr_42_: int
                for compr_42_ in (_dafny.Map({len((self).f15): (self).f17})).keys.Elements:
                    d_1076_v52_: int = compr_42_
                    if (d_1076_v52_) in (_dafny.Map({len((self).f15): (self).f17})):
                        coll42_[(d_1076_v52_) * (-276)] = False
                return _dafny.Map(coll42_)
            d_1075_v53_ = _dafny.Map({d_1073_v51_: _dafny.Map({(iife104_()
            ).set(d_1067_v46_, (self).f13): d_1068_v47_})})
            d_1077_v54_: _dafny.MultiSet
            d_1077_v54_ = _dafny.MultiSet([-748])
            d_1078_v55_: _dafny.Map
            d_1078_v55_ = _dafny.Map({(0) - ((d_1077_v54_).cardinality): d_1041_v30_})
            d_1079_v57_: _dafny.Seq
            d_1079_v57_ = _dafny.SeqWithoutIsStrInference([d_1041_v30_])
            d_1080_v58_: _dafny.Map
            d_1080_v58_ = _dafny.Map({d_1078_v55_: d_1079_v57_})
            def iife105_():
                coll43_ = _dafny.Map()
                compr_43_: _dafny.Map
                for compr_43_ in (d_1080_v58_).keys.Elements:
                    d_1081_v56_: _dafny.Map = compr_43_
                    if (d_1081_v56_) in (d_1080_v58_):
                        coll43_[d_1081_v56_] = d_1068_v47_
                return _dafny.Map(coll43_)
            d_1075_v53_ = (d_1075_v53_).set(d_1073_v51_, (_dafny.Map({d_1078_v55_: d_1068_v47_})) | ((iife105_()
            ).set(d_1078_v55_, d_1068_v47_)))
            d_1082_v59_: _dafny.Set
            d_1082_v59_ = _dafny.Set({d_1041_v30_, (self).f13})
            d_1041_v30_ = ((0) - (d_1067_v46_)) <= (len(d_1082_v59_))
            d_1041_v30_ = (self).f17
        d_1083_v60_: D6
        d_1083_v60_ = D6_DC12(False, d_1068_v47_, 629, d_1041_v30_)
        index157_ = default__.safeIndex(901, ((self).f14).length(0))
        ((self).f14)[index157_] = (d_1083_v60_).cf21
        index158_ = default__.safeIndex(901, ((self).f14).length(0))
        ((self).f14)[index158_] = (_dafny.Set({d_1067_v46_})).isdisjoint((D9_DC20((self).f15)).cf36)

    def m5(self, p0, globalState):
        r0: bool = False
        source19_ = self.f16
        if source19_.is_DC2:
            d_1084___mcc_h0_ = source19_.cf2
            d_1085___mcc_h1_ = source19_.cf3
            d_1086___mcc_h2_ = source19_.cf4
            d_1087_cf4_ = d_1086___mcc_h2_
            d_1088_cf3_ = d_1085___mcc_h1_
            d_1089_cf2_ = d_1084___mcc_h0_
            d_1090_v0_: _dafny.Seq
            d_1090_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "havb"))
            d_1091_v1_: _dafny.MultiSet
            d_1091_v1_ = _dafny.MultiSet([d_1090_v0_, d_1090_v0_])
            d_1087_cf4_ = not((d_1091_v1_).ispropersubset(d_1091_v1_))
            d_1092_v2_: _dafny.Array
            nw146_ = _dafny.Array(int(0), 14)
            d_1092_v2_ = nw146_
            d_1092_v2_ = d_1092_v2_
            d_1093_v3_: _dafny.Set
            d_1093_v3_ = _dafny.Set({d_1092_v2_})
            d_1089_cf2_ = len(d_1093_v3_)
            d_1094_v4_: _dafny.MultiSet
            d_1094_v4_ = _dafny.MultiSet([(d_1089_cf2_) == ((self).fm8((self).f17, self.f16, globalState)), (self).f17])
            d_1095_v5_: _dafny.Seq
            d_1095_v5_ = _dafny.SeqWithoutIsStrInference([d_1087_cf4_, d_1087_cf4_])
            d_1096_v6_: _dafny.Seq
            d_1096_v6_ = _dafny.SeqWithoutIsStrInference([d_1095_v5_, d_1095_v5_])
            rhs134_ = (d_1094_v4_ if True else (_dafny.MultiSet((d_1096_v6_)[default__.safeIndex(d_1088_cf3_, len(d_1096_v6_))])) - (d_1094_v4_))
            rhs135_ = (self).f13
            rhs136_ = True
            d_1094_v4_ = rhs134_
            d_1087_cf4_ = rhs135_
            d_1087_cf4_ = rhs136_
        elif True:
            d_1097___mcc_h3_ = source19_.cf1
            d_1098_cf1_ = d_1097___mcc_h3_
            d_1099_v7_: _dafny.Array
            def lambda44_(d_1100_p0_):
                def lambda45_(d_1101_i0_):
                    return (d_1101_i0_) + (d_1100_p0_)

                return lambda45_

            init25_ = lambda44_(p0)
            nw147_ = _dafny.Array(None, 4)
            for i0_25_ in range(nw147_.length(0)):
                nw147_[i0_25_] = init25_(i0_25_)
            d_1099_v7_ = nw147_
            index159_ = default__.safeIndex(708, (d_1099_v7_).length(0))
            (d_1099_v7_)[index159_] = p0
            index160_ = default__.safeIndex(708, (d_1099_v7_).length(0))
            (d_1099_v7_)[index160_] = p0
            d_1102_v8_: _dafny.Set
            d_1102_v8_ = _dafny.Set({(self).f17, (d_1098_cf1_) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_1103_i1_ in range(default__.abs(681))])), (self).f13, (self).f18})
            d_1102_v8_ = d_1102_v8_
            d_1098_cf1_ = d_1098_cf1_
            d_1104_v9_: str
            d_1104_v9_ = _dafny.CodePoint('h')
            d_1104_v9_ = d_1104_v9_
        d_1105_v10_: _dafny.Array
        nw148_ = _dafny.Array(_dafny.CodePoint('D'), 23)
        d_1105_v10_ = nw148_
        d_1105_v10_ = (d_1105_v10_ if (p0) < (p0) else d_1105_v10_)
        d_1106_v11_: _dafny.Array
        def lambda46_(d_1107_p0_):
            def lambda47_(d_1108_i3_):
                return (d_1108_i3_) * (d_1107_p0_)

            return lambda47_

        init26_ = lambda46_(p0)
        nw149_ = _dafny.Array(None, 23)
        for i0_26_ in range(nw149_.length(0)):
            nw149_[i0_26_] = init26_(i0_26_)
        d_1106_v11_ = nw149_
        d_1109_v12_: D2
        d_1109_v12_ = D2_DC4(d_1106_v11_, (self).f17, (0) - (p0), True, (self).f14)
        d_1110_v13_: _dafny.MultiSet
        d_1110_v13_ = _dafny.MultiSet([(self).f17])
        d_1111_v14_: _dafny.MultiSet
        d_1111_v14_ = d_1110_v13_
        d_1112_v15_: _dafny.Map
        d_1112_v15_ = _dafny.Map({p0: not((self).f18)})
        d_1113_v16_: D3
        d_1113_v16_ = D3_DC6((0) - (p0), (self.f12).cf4)
        d_1114_v17_: _dafny.Array
        nw150_ = _dafny.Array(None, 14)
        nw150_[int(0)] = d_1109_v12_
        nw150_[int(1)] = d_1109_v12_
        nw150_[int(2)] = D2_DC4(d_1106_v11_, (self).f17, (0) - (len(_dafny.Map({(self).f17: d_1111_v14_}))), ((d_1112_v15_)[p0] if (p0) in (d_1112_v15_) else False), (self).f14)
        nw150_[int(3)] = D2_DC4(d_1106_v11_, (self).f18, p0, (self).f17, (self).f14)
        nw150_[int(4)] = D2_DC4(d_1106_v11_, (self).f18, p0, (d_1113_v16_).cf13, (self).f14)
        nw150_[int(5)] = d_1109_v12_
        nw150_[int(6)] = d_1109_v12_
        nw150_[int(7)] = d_1109_v12_
        nw150_[int(8)] = d_1109_v12_
        nw150_[int(9)] = d_1109_v12_
        nw150_[int(10)] = d_1109_v12_
        nw150_[int(11)] = d_1109_v12_
        nw150_[int(12)] = D2_DC4(d_1106_v11_, (self).f13, p0, (self).f13, (d_1109_v12_).cf10)
        nw150_[int(13)] = d_1109_v12_
        d_1114_v17_ = nw150_
        d_1115_v18_: _dafny.Set
        d_1115_v18_ = _dafny.Set({d_1114_v17_})
        d_1116_i2_: int
        d_1116_i2_ = 0
        with _dafny.label("4"):
            while (_dafny.Set({d_1114_v17_})).issubset(d_1115_v18_):
                with _dafny.c_label("4"):
                    if (d_1116_i2_) >= (100):
                        raise _dafny.Break("4")
                    d_1116_i2_ = (d_1116_i2_) + (1)
                    d_1112_v15_ = (d_1112_v15_).set(p0, (self).f17)
                    d_1117_v19_: _dafny.Map
                    d_1117_v19_ = _dafny.Map({p0: p0})
                    d_1118_v20_: _dafny.Seq
                    d_1118_v20_ = _dafny.SeqWithoutIsStrInference([(self).f17])
                    d_1117_v19_ = (d_1117_v19_).set(default__.fm0(d_1118_v20_, (self).f18, (self).f13, p0, globalState), 114)
                    r0 = (self).f18
                    d_1119_v21_: _dafny.MultiSet
                    d_1119_v21_ = _dafny.MultiSet([p0, p0])
                    r0 = ((self).f18 if (d_1119_v21_) != (d_1119_v21_) else ((self).f15).issubset((self).f15))
                    pass
            pass
        d_1120_v22_: _dafny.Seq
        d_1120_v22_ = _dafny.SeqWithoutIsStrInference([not(False)])
        d_1121_v23_: _dafny.Map
        d_1121_v23_ = _dafny.Map({-575: p0})
        d_1122_v24_: _dafny.Map
        d_1122_v24_ = _dafny.Map({D9_DC23(p0, p0, (self).fm15(globalState), default__.fm0(d_1120_v22_, (self).f17, False, len(d_1121_v23_), globalState), (0) - (p0)): p0})
        d_1123_v25_: _dafny.Seq
        d_1123_v25_ = _dafny.SeqWithoutIsStrInference([len(d_1122_v24_), p0])
        d_1124_v26_: _dafny.Map
        d_1124_v26_ = _dafny.Map({d_1123_v25_: p0})
        index161_ = default__.safeIndex(895, (d_1106_v11_).length(0))
        (d_1106_v11_)[index161_] = len(d_1124_v26_)
        index162_ = default__.safeIndex(895, (d_1106_v11_).length(0))
        (d_1106_v11_)[index162_] = p0
        d_1125_v27_: _dafny.Seq
        d_1125_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "itq"))
        d_1126_v28_: D6
        d_1126_v28_ = D6_DC12((self).f13, d_1125_v27_, p0, (self).f18)
        d_1127_v29_: C2
        nw151_ = C2()
        nw151_.ctor__(default__.safeModuloInt(p0, (d_1106_v11_)[default__.safeIndex(895, (d_1106_v11_).length(0))]), (self).f13, self.f16, ((self).f13 if (d_1126_v28_).cf21 else (self).f13), (self).f14, ((self).f15) | ((self).f15), self.f12, True)
        d_1127_v29_ = nw151_
        d_1125_v27_ = (d_1125_v27_) + (d_1125_v27_)
        r0 = (self).f17
        return r0

    def m11(self, p0, p1, p2, globalState):
        d_1128_v0_: int
        d_1128_v0_ = 685
        if default__.fm30(d_1128_v0_, globalState):
            d_1129_v1_: _dafny.Map
            d_1129_v1_ = _dafny.Map({default__.fm30(d_1128_v0_, globalState): p1})
            d_1130_v2_: _dafny.Seq
            d_1130_v2_ = _dafny.SeqWithoutIsStrInference([True, False, p2, (self).f18, True])
            d_1129_v1_ = (d_1129_v1_).set((d_1130_v2_)[default__.safeIndex(d_1128_v0_, len(d_1130_v2_))], (self).f18)
            d_1131_v3_: _dafny.Seq
            d_1131_v3_ = _dafny.SeqWithoutIsStrInference([-44])
            d_1132_v4_: _dafny.Map
            d_1132_v4_ = _dafny.Map({p1: d_1128_v0_})
            d_1133_v5_: _dafny.Set
            d_1133_v5_ = _dafny.Set({p2})
            d_1134_v6_: _dafny.Map
            d_1134_v6_ = _dafny.Map({d_1133_v5_: d_1128_v0_})
            d_1135_v7_: _dafny.MultiSet
            d_1135_v7_ = _dafny.MultiSet([False])
            d_1136_v8_: _dafny.Seq
            d_1136_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jlbyvmu"))
            d_1137_v9_: _dafny.Array
            def lambda48_(d_1138_v0_):
                def lambda49_(d_1139_i0_):
                    return default__.safeDivisionInt(d_1139_i0_, d_1138_v0_)

                return lambda49_

            init27_ = lambda48_(d_1128_v0_)
            nw152_ = _dafny.Array(None, 9)
            for i0_27_ in range(nw152_.length(0)):
                nw152_[i0_27_] = init27_(i0_27_)
            d_1137_v9_ = nw152_
            d_1140_v10_: _dafny.Map
            d_1140_v10_ = _dafny.Map({d_1128_v0_: d_1128_v0_})
            d_1141_v11_: _dafny.Seq
            d_1141_v11_ = _dafny.SeqWithoutIsStrInference([d_1132_v4_])
            d_1142_v12_: D9
            d_1142_v12_ = D9_DC23(d_1128_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))), d_1128_v0_, len(d_1141_v11_), d_1128_v0_)
            d_1143_v13_: _dafny.Array
            nw153_ = _dafny.Array(None, 24)
            nw153_[int(0)] = ((d_1131_v3_)[default__.safeIndex(((d_1132_v4_)[(self).f17] if ((self).f17) in (d_1132_v4_) else d_1128_v0_), len(d_1131_v3_))] if p2 else 272)
            nw153_[int(1)] = len((default__.fm44((self).f17, globalState)) | (d_1134_v6_))
            nw153_[int(2)] = default__.fm0(d_1130_v2_, p1, False, d_1128_v0_, globalState)
            nw153_[int(3)] = d_1128_v0_
            nw153_[int(4)] = d_1128_v0_
            nw153_[int(5)] = d_1128_v0_
            nw153_[int(6)] = d_1128_v0_
            nw153_[int(7)] = (d_1128_v0_) * (d_1128_v0_)
            nw153_[int(8)] = default__.safeDivisionInt(d_1128_v0_, d_1128_v0_)
            nw153_[int(9)] = d_1128_v0_
            nw153_[int(10)] = (0) - ((305) + (len((self).f15)))
            nw153_[int(11)] = 770
            nw153_[int(12)] = default__.safeDivisionInt((d_1135_v7_).cardinality, (0) - (d_1128_v0_))
            nw153_[int(13)] = default__.fm0(d_1130_v2_, (self).f13, False, default__.fm0(d_1130_v2_, (self).f13, default__.fm30(len(d_1136_v8_), globalState), d_1128_v0_, globalState), globalState)
            nw153_[int(14)] = d_1128_v0_
            nw153_[int(15)] = d_1128_v0_
            nw153_[int(16)] = len(_dafny.Set({(D2_DC4(d_1137_v9_, not((self).f17), d_1128_v0_, (self).f18, (self).f14)).cf8}))
            nw153_[int(17)] = d_1128_v0_
            nw153_[int(18)] = (d_1128_v0_) - (998)
            nw153_[int(19)] = (_dafny.MultiSet([d_1128_v0_, d_1128_v0_, len(d_1140_v10_), -793, len((self).f15)])).cardinality
            nw153_[int(20)] = (d_1128_v0_) * (d_1128_v0_)
            nw153_[int(21)] = ((d_1140_v10_)[(d_1142_v12_).cf45] if ((d_1142_v12_).cf45) in (d_1140_v10_) else len(d_1136_v8_))
            nw153_[int(22)] = d_1128_v0_
            nw153_[int(23)] = 630
            d_1143_v13_ = nw153_
            index163_ = default__.safeIndex(465, (d_1143_v13_).length(0))
            (d_1143_v13_)[index163_] = d_1128_v0_
            index164_ = default__.safeIndex(465, (d_1143_v13_).length(0))
            (d_1143_v13_)[index164_] = ((self).fm15(globalState)) * (d_1128_v0_)
            d_1144_v14_: C1
            nw154_ = C1()
            nw154_.ctor__(_dafny.CodePoint('j'), self.f12, (self).f17)
            d_1144_v14_ = nw154_
            d_1145_v15_: _dafny.Array
            nw155_ = _dafny.Array(None, 18)
            nw155_[int(0)] = d_1144_v14_
            nw155_[int(1)] = d_1144_v14_
            nw155_[int(2)] = d_1144_v14_
            nw155_[int(3)] = d_1144_v14_
            nw155_[int(4)] = d_1144_v14_
            nw155_[int(5)] = d_1144_v14_
            nw155_[int(6)] = d_1144_v14_
            nw155_[int(7)] = d_1144_v14_
            nw155_[int(8)] = d_1144_v14_
            nw155_[int(9)] = d_1144_v14_
            nw155_[int(10)] = d_1144_v14_
            nw155_[int(11)] = d_1144_v14_
            nw155_[int(12)] = d_1144_v14_
            nw155_[int(13)] = d_1144_v14_
            nw155_[int(14)] = d_1144_v14_
            nw155_[int(15)] = d_1144_v14_
            nw155_[int(16)] = d_1144_v14_
            nw155_[int(17)] = d_1144_v14_
            d_1145_v15_ = nw155_
            d_1146_v16_: _dafny.MultiSet
            d_1146_v16_ = _dafny.MultiSet([d_1145_v15_, d_1145_v15_])
            d_1147_v17_: bool
            d_1147_v17_ = False
            rhs137_ = (d_1143_v13_)[default__.safeIndex(465, (d_1143_v13_).length(0))]
            rhs138_ = d_1146_v16_
            rhs139_ = ((d_1143_v13_)[default__.safeIndex(465, (d_1143_v13_).length(0))]) > ((d_1143_v13_)[default__.safeIndex(465, (d_1143_v13_).length(0))])
            rhs140_ = ((d_1135_v7_).cardinality) == ((d_1143_v13_)[default__.safeIndex(465, (d_1143_v13_).length(0))])
            rhs141_ = p1
            lhs114_ = globalState
            lhs114_.f7 = rhs137_
            d_1146_v16_ = rhs138_
            d_1147_v17_ = rhs139_
            d_1147_v17_ = rhs140_
            d_1147_v17_ = rhs141_
            if (self).f13:
                index165_ = default__.safeIndex(465, (d_1143_v13_).length(0))
                (d_1143_v13_)[index165_] = -903
                (globalState).f7 = (d_1143_v13_)[default__.safeIndex(465, (d_1143_v13_).length(0))]
                d_1147_v17_ = p2
                d_1148_v18_: _dafny.Map
                d_1148_v18_ = _dafny.Map({len((self).f15): p2})
                rhs142_ = ((d_1144_v14_.f23 if p2 else d_1144_v14_.f23) if ((d_1148_v18_)[(d_1135_v7_).cardinality] if ((d_1135_v7_).cardinality) in (d_1148_v18_) else (self).f13) else d_1144_v14_.f23)
                rhs143_ = d_1144_v14_.f23
                lhs115_ = d_1144_v14_
                lhs116_ = d_1144_v14_
                lhs115_.f23 = rhs142_
                lhs116_.f23 = rhs143_
                d_1149_v19_: _dafny.Array
                nw156_ = _dafny.Array(False, 19)
                d_1149_v19_ = nw156_
                d_1149_v19_ = p0
            elif True:
                d_1150_v20_: _dafny.Map
                d_1150_v20_ = _dafny.Map({(d_1143_v13_)[default__.safeIndex(465, (d_1143_v13_).length(0))]: p0})
                d_1150_v20_ = (d_1150_v20_).set(len(d_1130_v2_), (self).f14)
                d_1151_v21_: _dafny.Set
                d_1151_v21_ = _dafny.Set({d_1136_v8_})
                d_1152_v22_: _dafny.Set
                d_1152_v22_ = d_1151_v21_
                pat_let_tv31_ = d_1143_v13_
                pat_let_tv32_ = d_1143_v13_
                def iife106_(_pat_let31_0):
                    def iife107_(d_1153_dt__update__tmp_h0_):
                        def iife108_(_pat_let32_0):
                            def iife109_(d_1154_dt__update_hcf2_h0_):
                                return D1_DC2(d_1154_dt__update_hcf2_h0_, (d_1153_dt__update__tmp_h0_).cf3, (d_1153_dt__update__tmp_h0_).cf4)
                            return iife109_(_pat_let32_0)
                        return iife108_((pat_let_tv32_)[default__.safeIndex(465, (pat_let_tv31_).length(0))])
                    return iife107_(_pat_let31_0)
                d_1152_v22_ = (d_1151_v21_ if (iife106_(self.f12)).cf4 else d_1152_v22_)
                d_1155_v23_: _dafny.Map
                d_1155_v23_ = _dafny.Map({d_1144_v14_: d_1128_v0_})
                d_1155_v23_ = (d_1155_v23_).set(d_1144_v14_, (d_1143_v13_)[default__.safeIndex(465, (d_1143_v13_).length(0))])
                d_1156_v24_: _dafny.Array
                nw157_ = _dafny.Array(None, 20)
                nw157_[int(0)] = d_1137_v9_
                nw157_[int(1)] = d_1137_v9_
                nw157_[int(2)] = d_1137_v9_
                nw157_[int(3)] = d_1137_v9_
                nw157_[int(4)] = d_1137_v9_
                nw157_[int(5)] = d_1137_v9_
                nw157_[int(6)] = d_1137_v9_
                nw157_[int(7)] = d_1143_v13_
                nw157_[int(8)] = d_1143_v13_
                nw157_[int(9)] = d_1143_v13_
                nw157_[int(10)] = (D2_DC4(d_1137_v9_, (self).f18, (self).fm8((self).f13, self.f16, globalState), not((self).f18), (self).f14)).cf6
                nw157_[int(11)] = d_1143_v13_
                nw157_[int(12)] = d_1137_v9_
                nw157_[int(13)] = d_1143_v13_
                nw157_[int(14)] = d_1137_v9_
                nw157_[int(15)] = d_1137_v9_
                nw157_[int(16)] = d_1143_v13_
                nw157_[int(17)] = d_1143_v13_
                nw157_[int(18)] = d_1137_v9_
                nw157_[int(19)] = d_1143_v13_
                d_1156_v24_ = nw157_
                index166_ = default__.safeIndex(149, (d_1156_v24_).length(0))
                (d_1156_v24_)[index166_] = d_1137_v9_
                d_1157_v25_: _dafny.Map
                d_1157_v25_ = _dafny.Map({d_1128_v0_: d_1143_v13_})
                index167_ = default__.safeIndex(149, (d_1156_v24_).length(0))
                (d_1156_v24_)[index167_] = ((d_1157_v25_)[d_1128_v0_] if (d_1128_v0_) in (d_1157_v25_) else (D2_DC4(d_1143_v13_, (self).f13, 612, (self).f13, (self).f14)).cf6)
                d_1158_v26_: _dafny.Map
                d_1158_v26_ = _dafny.Map({(self).f13: _dafny.Set({(d_1143_v13_)[default__.safeIndex(465, (d_1143_v13_).length(0))]})})
                d_1159_v27_: _dafny.Map
                d_1159_v27_ = _dafny.Map({(d_1144_v14_).fm43(globalState): ((d_1158_v26_)[(self).f17] if ((self).f17) in (d_1158_v26_) else (self).f15)})
                d_1158_v26_ = _dafny.Map({(self).f13: (self).f15})
            (self).m12(globalState)
        elif True:
            d_1160_v28_: _dafny.Seq
            d_1160_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lxk"))
            d_1161_v29_: _dafny.Seq
            d_1161_v29_ = _dafny.SeqWithoutIsStrInference([d_1128_v0_])
            d_1162_v30_: D3
            d_1162_v30_ = D3_DC5(_dafny.CodePoint('c'))
            d_1163_v31_: _dafny.MultiSet
            d_1163_v31_ = _dafny.MultiSet([(d_1162_v30_).cf11])
            d_1164_v32_: str
            d_1164_v32_ = _dafny.CodePoint('q')
            d_1165_v33_: _dafny.Map
            d_1165_v33_ = _dafny.Map({997: D7_DC17(D7_DC15(d_1161_v29_))})
            d_1166_v34_: _dafny.MultiSet
            d_1166_v34_ = _dafny.MultiSet([d_1128_v0_, d_1128_v0_])
            d_1167_v35_: D7
            d_1167_v35_ = D7_DC16(d_1128_v0_)
            d_1168_v36_: D7
            d_1168_v36_ = D7_DC17(d_1167_v35_)
            d_1169_v37_: _dafny.Array
            nw158_ = _dafny.Array(None, 20)
            nw158_[int(0)] = d_1128_v0_
            nw158_[int(1)] = d_1128_v0_
            nw158_[int(2)] = default__.safeDivisionInt(642, d_1128_v0_)
            nw158_[int(3)] = default__.safeModuloInt(d_1128_v0_, d_1128_v0_)
            nw158_[int(4)] = (0) - (d_1128_v0_)
            nw158_[int(5)] = d_1128_v0_
            nw158_[int(6)] = len(d_1160_v28_)
            nw158_[int(7)] = (len(d_1161_v29_)) - ((0) - (d_1128_v0_))
            nw158_[int(8)] = d_1128_v0_
            nw158_[int(9)] = d_1128_v0_
            nw158_[int(10)] = (((d_1163_v31_)[d_1164_v32_] if (d_1164_v32_) in (d_1163_v31_) else d_1128_v0_) if p1 else (self).fm9(d_1164_v32_, globalState))
            nw158_[int(11)] = d_1128_v0_
            nw158_[int(12)] = len((self).f15)
            nw158_[int(13)] = d_1128_v0_
            nw158_[int(14)] = (d_1128_v0_) + (d_1128_v0_)
            nw158_[int(15)] = d_1128_v0_
            nw158_[int(16)] = -151
            nw158_[int(17)] = len((d_1160_v28_).set(default__.safeIndex(len((d_1165_v33_).set((d_1166_v34_).cardinality, d_1168_v36_)), len(d_1160_v28_)), d_1164_v32_))
            nw158_[int(18)] = d_1128_v0_
            nw158_[int(19)] = (d_1128_v0_) + (len(d_1160_v28_))
            d_1169_v37_ = nw158_
            index168_ = default__.safeIndex(920, (d_1169_v37_).length(0))
            (d_1169_v37_)[index168_] = d_1128_v0_
            index169_ = default__.safeIndex(920, (d_1169_v37_).length(0))
            (d_1169_v37_)[index169_] = d_1128_v0_
            d_1128_v0_ = len(default__.fm20(globalState))
            d_1170_v38_: _dafny.Array
            nw159_ = _dafny.Array(_dafny.CodePoint('D'), 19)
            d_1170_v38_ = nw159_
            index170_ = default__.safeIndex(494, (d_1170_v38_).length(0))
            (d_1170_v38_)[index170_] = d_1164_v32_
            index171_ = default__.safeIndex(494, (d_1170_v38_).length(0))
            (d_1170_v38_)[index171_] = d_1164_v32_
            d_1171_v39_: _dafny.Array
            nw160_ = _dafny.Array(_dafny.Array(None, 0), 11)
            d_1171_v39_ = nw160_
            d_1172_v40_: _dafny.Array
            d_1172_v40_ = d_1171_v39_
            source20_ = d_1172_v40_
            d_1173___mcc_h0_ = source20_
            d_1174_cf15_ = d_1173___mcc_h0_
            d_1175_v41_: _dafny.Set
            d_1175_v41_ = _dafny.Set({(self).f18})
            d_1128_v0_ = ((d_1166_v34_) | ((_dafny.MultiSet([len(d_1175_v41_)])).intersection(_dafny.MultiSet([(d_1169_v37_)[default__.safeIndex(920, (d_1169_v37_).length(0))]])))).cardinality
            index172_ = default__.safeIndex(426, (p0).length(0))
            (p0)[index172_] = not(p1)
            index173_ = default__.safeIndex(426, (p0).length(0))
            (p0)[index173_] = p1
            d_1176_v42_: _dafny.Array
            def lambda50_(d_1177_i1_):
                return False

            init28_ = lambda50_
            nw161_ = _dafny.Array(None, 11)
            for i0_28_ in range(nw161_.length(0)):
                nw161_[i0_28_] = init28_(i0_28_)
            d_1176_v42_ = nw161_
            d_1178_v43_: _dafny.Array
            nw162_ = _dafny.Array(None, 1)
            nw162_[int(0)] = d_1169_v37_
            d_1178_v43_ = nw162_
            index174_ = default__.safeIndex(414, (d_1178_v43_).length(0))
            (d_1178_v43_)[index174_] = d_1169_v37_
            d_1179_v44_: _dafny.MultiSet
            d_1179_v44_ = _dafny.MultiSet([(self).f14])
            d_1180_v45_: _dafny.MultiSet
            d_1180_v45_ = _dafny.MultiSet([d_1176_v42_])
            d_1181_v46_: _dafny.MultiSet
            d_1181_v46_ = _dafny.MultiSet([d_1161_v29_])
            d_1182_v47_: _dafny.Array
            nw163_ = _dafny.Array(D3.default()(), 5)
            d_1182_v47_ = nw163_
            d_1183_v48_: _dafny.Map
            d_1183_v48_ = _dafny.Map({d_1182_v47_: p1})
            d_1184_v49_: _dafny.Map
            d_1184_v49_ = _dafny.Map({d_1183_v48_: d_1176_v42_})
            d_1185_v50_: _dafny.Seq
            d_1185_v50_ = _dafny.SeqWithoutIsStrInference([d_1169_v37_, (d_1169_v37_ if p1 else d_1169_v37_)])
            index175_ = default__.safeIndex(426, (p0).length(0))
            index176_ = default__.safeIndex(414, (d_1178_v43_).length(0))
            rhs144_ = default__.safeModuloInt((d_1169_v37_)[default__.safeIndex(920, (d_1169_v37_).length(0))], (0) - (((d_1180_v45_)).cardinality))
            rhs145_ = (d_1181_v46_).isdisjoint(d_1181_v46_)
            rhs146_ = ((d_1184_v49_)[d_1183_v48_] if (d_1183_v48_) in (d_1184_v49_) else (self).f14)
            rhs147_ = (d_1169_v37_)[default__.safeIndex(920, (d_1169_v37_).length(0))]
            rhs148_ = (d_1185_v50_)[default__.safeIndex(d_1128_v0_, len(d_1185_v50_))]
            lhs117_ = p0
            lhs118_ = default__.safeIndex(426, (p0).length(0))
            lhs119_ = globalState
            lhs120_ = d_1178_v43_
            lhs121_ = default__.safeIndex(414, (d_1178_v43_).length(0))
            d_1128_v0_ = rhs144_
            lhs117_[lhs118_] = rhs145_
            d_1176_v42_ = rhs146_
            lhs119_.f7 = rhs147_
            lhs120_[lhs121_] = rhs148_
            d_1186_v51_: _dafny.Seq
            d_1186_v51_ = _dafny.SeqWithoutIsStrInference([p1, (self).f13])
            d_1187_v52_: _dafny.Map
            d_1187_v52_ = _dafny.Map({(d_1176_v42_ if (self).f18 else d_1176_v42_): (0) - (len(d_1186_v51_))})
            d_1187_v52_ = (d_1187_v52_).set(d_1176_v42_, default__.safeModuloInt(d_1128_v0_, (d_1169_v37_)[default__.safeIndex(920, (d_1169_v37_).length(0))]))
            if p1:
                d_1188_v53_: bool
                d_1188_v53_ = False
                d_1188_v53_ = p2
                d_1160_v28_ = default__.fm20(globalState)
                d_1164_v32_ = d_1164_v32_
                index177_ = default__.safeIndex(920, (d_1169_v37_).length(0))
                (d_1169_v37_)[index177_] = (0) - ((default__.safeDivisionInt(41, d_1128_v0_)) * ((0) - (d_1128_v0_)))
                d_1188_v53_ = ((d_1166_v34_ if (self).f18 else _dafny.MultiSet([670]))).isdisjoint(default__.fm19(globalState))
            elif True:
                d_1189_v54_: _dafny.Map
                d_1189_v54_ = _dafny.Map({D6_DC11(d_1166_v34_): False})
                d_1190_v55_: C2
                nw164_ = C2()
                nw164_.ctor__(d_1128_v0_, (self).f18, self.f16, ((d_1189_v54_)[D6_DC11(d_1166_v34_)] if (D6_DC11(d_1166_v34_)) in (d_1189_v54_) else p1), p0, ((self).f15) - ((self).f15), self.f12, ((0) - (d_1128_v0_)) == ((d_1169_v37_)[default__.safeIndex(920, (d_1169_v37_).length(0))]))
                d_1190_v55_ = nw164_
                d_1191_v56_: _dafny.Map
                d_1191_v56_ = _dafny.Map({d_1169_v37_: (d_1169_v37_)[default__.safeIndex(920, (d_1169_v37_).length(0))]})
                d_1191_v56_ = (d_1191_v56_).set(d_1169_v37_, default__.safeDivisionInt((d_1169_v37_)[default__.safeIndex(920, (d_1169_v37_).length(0))], d_1128_v0_))
                d_1192_v57_: _dafny.Map
                d_1192_v57_ = _dafny.Map({(d_1169_v37_)[default__.safeIndex(920, (d_1169_v37_).length(0))]: (d_1190_v55_).f22})
                d_1192_v57_ = d_1192_v57_
                d_1193_v58_: _dafny.Map
                d_1193_v58_ = _dafny.Map({True: (d_1190_v55_).f21})
                d_1193_v58_ = (d_1193_v58_).set((self).f13, (d_1190_v55_).f21)
                (globalState).f7 = default__.safeDivisionInt(d_1128_v0_, (d_1169_v37_)[default__.safeIndex(920, (d_1169_v37_).length(0))])
        d_1194_v59_: _dafny.Map
        d_1194_v59_ = _dafny.Map({d_1128_v0_: _dafny.CodePoint('h')})
        d_1195_v60_: _dafny.Seq
        d_1195_v60_ = _dafny.SeqWithoutIsStrInference([d_1128_v0_, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f17, (self).f13, (self).f13, (self).f18]))).cardinality])
        d_1196_v61_: _dafny.Array
        nw165_ = _dafny.Array(None, 9)
        nw165_[int(0)] = d_1128_v0_
        nw165_[int(1)] = len((d_1194_v59_) | (d_1194_v59_))
        nw165_[int(2)] = (0) - ((d_1195_v60_)[default__.safeIndex(d_1128_v0_, len(d_1195_v60_))])
        nw165_[int(3)] = (0) - ((0) - ((0) - (d_1128_v0_)))
        nw165_[int(4)] = default__.fm0(_dafny.SeqWithoutIsStrInference([(self).f13]), True, (self).f17, 384, globalState)
        nw165_[int(5)] = d_1128_v0_
        nw165_[int(6)] = d_1128_v0_
        nw165_[int(7)] = d_1128_v0_
        nw165_[int(8)] = d_1128_v0_
        d_1196_v61_ = nw165_
        index178_ = default__.safeIndex(210, (d_1196_v61_).length(0))
        (d_1196_v61_)[index178_] = (d_1128_v0_) + (d_1128_v0_)
        d_1197_v62_: _dafny.MultiSet
        d_1197_v62_ = _dafny.MultiSet([(self).f14])
        d_1198_v63_: bool
        d_1198_v63_ = True
        d_1199_v64_: _dafny.Seq
        d_1199_v64_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "paa"))
        index179_ = default__.safeIndex(70, (d_1196_v61_).length(0))
        (d_1196_v61_)[index179_] = len(d_1199_v64_)
        index180_ = default__.safeIndex(210, (d_1196_v61_).length(0))
        index181_ = default__.safeIndex(70, (d_1196_v61_).length(0))
        rhs149_ = d_1128_v0_
        rhs150_ = d_1197_v62_
        rhs151_ = p1
        rhs152_ = False
        rhs153_ = d_1128_v0_
        lhs122_ = d_1196_v61_
        lhs123_ = default__.safeIndex(210, (d_1196_v61_).length(0))
        lhs124_ = d_1196_v61_
        lhs125_ = default__.safeIndex(70, (d_1196_v61_).length(0))
        lhs122_[lhs123_] = rhs149_
        d_1197_v62_ = rhs150_
        d_1198_v63_ = rhs151_
        d_1198_v63_ = rhs152_
        lhs124_[lhs125_] = rhs153_
        d_1200_v65_: _dafny.Seq
        d_1200_v65_ = _dafny.SeqWithoutIsStrInference([d_1198_v63_])
        d_1201_v66_: _dafny.Map
        d_1201_v66_ = _dafny.Map({(d_1196_v61_)[default__.safeIndex(210, (d_1196_v61_).length(0))]: (d_1200_v65_)[default__.safeIndex(d_1128_v0_, len(d_1200_v65_))]})
        if ((d_1201_v66_)[d_1128_v0_] if (d_1128_v0_) in (d_1201_v66_) else d_1198_v63_):
            (globalState).f7 = (d_1196_v61_)[default__.safeIndex(210, (d_1196_v61_).length(0))]
            index182_ = default__.safeIndex(210, (d_1196_v61_).length(0))
            (d_1196_v61_)[index182_] = len(d_1195_v60_)
            index183_ = default__.safeIndex(210, (d_1196_v61_).length(0))
            (d_1196_v61_)[index183_] = default__.safeDivisionInt((self).fm15(globalState), (d_1196_v61_)[default__.safeIndex(210, (d_1196_v61_).length(0))])
            d_1202_v67_: _dafny.Map
            d_1202_v67_ = _dafny.Map({default__.fm21(globalState): d_1198_v63_})
            d_1202_v67_ = (d_1202_v67_).set(self.f16, (810) < (d_1128_v0_))
            (globalState).f7 = default__.safeModuloInt((d_1196_v61_)[default__.safeIndex(210, (d_1196_v61_).length(0))], d_1128_v0_)
        elif True:
            d_1203_v68_: _dafny.Seq
            d_1203_v68_ = _dafny.SeqWithoutIsStrInference([d_1199_v64_])
            d_1204_v69_: _dafny.Array
            nw166_ = _dafny.Array(None, 3)
            nw166_[int(0)] = (d_1203_v68_) + (d_1203_v68_)
            nw166_[int(1)] = d_1203_v68_
            nw166_[int(2)] = _dafny.SeqWithoutIsStrInference([d_1199_v64_])
            d_1204_v69_ = nw166_
            index184_ = default__.safeIndex(207, (d_1204_v69_).length(0))
            (d_1204_v69_)[index184_] = d_1203_v68_
            d_1205_v70_: D6
            d_1205_v70_ = D6_DC11(_dafny.MultiSet([(d_1196_v61_)[default__.safeIndex(210, (d_1196_v61_).length(0))]]))
            d_1206_v71_: _dafny.Map
            d_1206_v71_ = _dafny.Map({d_1205_v70_: d_1199_v64_})
            d_1207_v72_: D6
            d_1207_v72_ = D6_DC12((self).f13, ((d_1206_v71_)[d_1205_v70_] if (d_1205_v70_) in (d_1206_v71_) else d_1199_v64_), (d_1196_v61_)[default__.safeIndex(210, (d_1196_v61_).length(0))], d_1198_v63_)
            index185_ = default__.safeIndex(207, (d_1204_v69_).length(0))
            (d_1204_v69_)[index185_] = default__.fm45(d_1207_v72_, d_1128_v0_, globalState)
            d_1208_v73_: _dafny.Map
            d_1208_v73_ = _dafny.Map({332: -981})
            index186_ = default__.safeIndex(222, ((self).f14).length(0))
            ((self).f14)[index186_] = not(default__.fm30(len(d_1208_v73_), globalState))
            d_1209_v74_: str
            d_1209_v74_ = _dafny.CodePoint('y')
            d_1210_v75_: D9
            d_1210_v75_ = D9_DC22((self).f17, (d_1196_v61_)[default__.safeIndex(210, (d_1196_v61_).length(0))], (self).f18, default__.fm36(p2, _dafny.CodePoint('k'), p1, globalState), (d_1199_v64_).set(default__.safeIndex(d_1128_v0_, len(d_1199_v64_)), d_1209_v74_))
            index187_ = default__.safeIndex(222, ((self).f14).length(0))
            ((self).f14)[index187_] = (d_1210_v75_).cf37
            d_1128_v0_ = (d_1128_v0_) + (398)
            d_1211_v76_: _dafny.Map
            d_1211_v76_ = _dafny.Map({d_1198_v63_: d_1128_v0_})
            d_1212_v77_: _dafny.MultiSet
            d_1212_v77_ = _dafny.MultiSet([(d_1196_v61_)[default__.safeIndex(210, (d_1196_v61_).length(0))], 32, len(d_1211_v76_), (d_1196_v61_)[default__.safeIndex(210, (d_1196_v61_).length(0))]])
            d_1213_v78_: _dafny.MultiSet
            d_1213_v78_ = _dafny.MultiSet([(d_1212_v77_).cardinality])
            d_1214_v79_: _dafny.Map
            d_1214_v79_ = _dafny.Map({121: d_1213_v78_})
            d_1215_v80_: _dafny.Seq
            d_1215_v80_ = _dafny.SeqWithoutIsStrInference([d_1211_v76_, d_1211_v76_, d_1211_v76_, d_1211_v76_])
            d_1216_v81_: _dafny.MultiSet
            d_1216_v81_ = _dafny.MultiSet([d_1211_v76_, d_1211_v76_, d_1211_v76_, (d_1215_v80_)[default__.safeIndex((d_1207_v72_).cf20, len(d_1215_v80_))]])
            index188_ = default__.safeIndex(210, (d_1196_v61_).length(0))
            rhs154_ = (0) - (((self).fm8(d_1198_v63_, self.f16, globalState)) - (len((d_1214_v79_) | ((d_1214_v79_).set((d_1196_v61_)[default__.safeIndex(210, (d_1196_v61_).length(0))], d_1213_v78_)))))
            rhs155_ = ((default__.fm0(d_1200_v65_, (self).f18, (self).f18, (d_1216_v81_).cardinality, globalState)) + (d_1128_v0_)) * ((0) - ((d_1196_v61_)[default__.safeIndex(210, (d_1196_v61_).length(0))]))
            lhs126_ = globalState
            lhs127_ = d_1196_v61_
            lhs128_ = default__.safeIndex(210, (d_1196_v61_).length(0))
            lhs126_.f7 = rhs154_
            lhs127_[lhs128_] = rhs155_
            d_1217_v82_: D6
            d_1217_v82_ = D6_DC13(((self).f14)[default__.safeIndex(222, ((self).f14).length(0))], (self).f18, (d_1196_v61_)[default__.safeIndex(210, (d_1196_v61_).length(0))], (d_1196_v61_)[default__.safeIndex(210, (d_1196_v61_).length(0))], p2)
            d_1218_v83_: D6
            d_1218_v83_ = D6_DC14(d_1217_v82_)
            d_1199_v64_ = ((default__.fm32(d_1218_v83_, globalState) if ((self).f14)[default__.safeIndex(222, ((self).f14).length(0))] else d_1199_v64_)) + ((d_1203_v68_)[default__.safeIndex(d_1128_v0_, len(d_1203_v68_))])
        d_1219_v84_: _dafny.Map
        d_1219_v84_ = _dafny.Map({True: (self).fm8(False, self.f16, globalState)})
        d_1220_v85_: _dafny.Seq
        d_1220_v85_ = _dafny.SeqWithoutIsStrInference([(d_1219_v84_) | (d_1219_v84_)])
        d_1221_v86_: _dafny.Map
        d_1221_v86_ = _dafny.Map({(d_1196_v61_)[default__.safeIndex(210, (d_1196_v61_).length(0))]: _dafny.SeqWithoutIsStrInference([d_1219_v84_])})
        d_1220_v85_ = (((d_1221_v86_)[d_1128_v0_] if (d_1128_v0_) in (d_1221_v86_) else _dafny.SeqWithoutIsStrInference([(d_1220_v85_)[default__.safeIndex((0) - ((d_1196_v61_)[default__.safeIndex(210, (d_1196_v61_).length(0))]), len(d_1220_v85_))]]))) + (d_1220_v85_)
        d_1222_i2_: int
        d_1222_i2_ = 0
        with _dafny.label("5"):
            while p1:
                with _dafny.c_label("5"):
                    if (d_1222_i2_) >= (100):
                        raise _dafny.Break("5")
                    d_1222_i2_ = (d_1222_i2_) + (1)
                    index189_ = default__.safeIndex(274, (p0).length(0))
                    (p0)[index189_] = False
                    index190_ = default__.safeIndex(917, (p0).length(0))
                    (p0)[index190_] = p2
                    d_1223_v87_: _dafny.MultiSet
                    d_1223_v87_ = _dafny.MultiSet([(self).f18, True])
                    index191_ = default__.safeIndex(274, (p0).length(0))
                    index192_ = default__.safeIndex(917, (p0).length(0))
                    rhs156_ = (self).f13
                    rhs157_ = True
                    rhs158_ = (self).f18
                    rhs159_ = ((d_1195_v60_) <= ((default__.fm34((d_1196_v61_)[default__.safeIndex(210, (d_1196_v61_).length(0))], globalState)).set(default__.safeIndex(d_1128_v0_, len(default__.fm34((d_1196_v61_)[default__.safeIndex(210, (d_1196_v61_).length(0))], globalState))), (d_1196_v61_)[default__.safeIndex(210, (d_1196_v61_).length(0))]))) and ((len(d_1199_v64_)) >= ((d_1223_v87_).cardinality))
                    rhs160_ = (d_1195_v60_)[default__.safeIndex((250 if (self).f13 else d_1128_v0_), len(d_1195_v60_))]
                    lhs129_ = p0
                    lhs130_ = default__.safeIndex(274, (p0).length(0))
                    lhs131_ = p0
                    lhs132_ = default__.safeIndex(917, (p0).length(0))
                    lhs129_[lhs130_] = rhs156_
                    lhs131_[lhs132_] = rhs157_
                    d_1198_v63_ = rhs158_
                    d_1198_v63_ = rhs159_
                    d_1128_v0_ = rhs160_
                    d_1224_v88_: _dafny.Array
                    def lambda51_(d_1225_v0_):
                        def lambda52_(d_1226_i3_):
                            return D6_DC11(_dafny.MultiSet([313, d_1225_v0_]))

                        return lambda52_

                    init29_ = lambda51_(d_1128_v0_)
                    nw167_ = _dafny.Array(None, 6)
                    for i0_29_ in range(nw167_.length(0)):
                        nw167_[i0_29_] = init29_(i0_29_)
                    d_1224_v88_ = nw167_
                    d_1224_v88_ = d_1224_v88_
                    (globalState).f2 = d_1200_v65_
                    (globalState).f7 = (-41) * (((d_1196_v61_)[default__.safeIndex(210, (d_1196_v61_).length(0))]) * (d_1128_v0_))
                    pass
            pass
        d_1227_v89_: _dafny.Map
        d_1227_v89_ = _dafny.Map({d_1128_v0_: d_1199_v64_})
        d_1227_v89_ = (d_1227_v89_).set((d_1196_v61_)[default__.safeIndex(210, (d_1196_v61_).length(0))], d_1199_v64_)

    def m12(self, globalState):
        d_1228_v0_: str
        d_1228_v0_ = _dafny.CodePoint('j')
        d_1229_v1_: C0
        nw168_ = C0()
        nw168_.ctor__(True, (_dafny.CodePoint('p') if (self).f13 else d_1228_v0_), self.f12, (self).f18)
        d_1229_v1_ = nw168_
        d_1230_v2_: int
        d_1230_v2_ = 385
        (globalState).f7 = d_1230_v2_
        (d_1229_v1_).f19 = (self).f18
        d_1231_v3_: _dafny.Seq
        d_1231_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))
        d_1232_v4_: _dafny.Map
        d_1232_v4_ = _dafny.Map({d_1231_v3_: (self).f18})
        d_1232_v4_ = d_1232_v4_
        d_1231_v3_ = _dafny.SeqWithoutIsStrInference([d_1228_v0_ for d_1233_i0_ in range(default__.abs(818))])
        hi6_ = d_1230_v2_
        for d_1234_i1_ in range(d_1230_v2_, hi6_):
            (globalState).f7 = default__.safeModuloInt(d_1234_i1_, -145)
            d_1235_v5_: _dafny.Array
            nw169_ = _dafny.Array(int(0), 21)
            d_1235_v5_ = nw169_
            d_1236_v6_: _dafny.Set
            d_1236_v6_ = _dafny.Set({d_1229_v1_.f19, d_1229_v1_.f19})
            d_1237_v7_: _dafny.Seq
            d_1237_v7_ = _dafny.SeqWithoutIsStrInference([len(d_1236_v6_), d_1230_v2_, d_1230_v2_, d_1230_v2_])
            d_1238_v8_: _dafny.Map
            d_1238_v8_ = _dafny.Map({d_1235_v5_: d_1237_v7_})
            d_1239_v9_: _dafny.Seq
            d_1239_v9_ = _dafny.SeqWithoutIsStrInference([d_1230_v2_, len(d_1238_v8_)])
            d_1240_v10_: D9
            d_1240_v10_ = D9_DC22(d_1229_v1_.f19, (d_1237_v7_)[default__.safeIndex(d_1230_v2_, len(d_1237_v7_))], (self).f13, d_1236_v6_, d_1231_v3_)
            d_1241_v11_: _dafny.Map
            d_1241_v11_ = _dafny.Map({(d_1230_v2_) - (len((_dafny.SeqWithoutIsStrInference([(self).f18])).set(default__.safeIndex(d_1234_i1_, len(_dafny.SeqWithoutIsStrInference([(self).f18]))), (self).f13))): (d_1240_v10_).cf38})
            d_1241_v11_ = d_1241_v11_
            (globalState).f7 = d_1234_i1_
            def iife110_():
                coll44_ = _dafny.Set()
                compr_44_: int
                for compr_44_ in _dafny.IntegerRange(892, 210):
                    d_1242_v12_: int = compr_44_
                    if ((892) <= (d_1242_v12_)) and ((d_1242_v12_) < (210)):
                        coll44_ = coll44_.union(_dafny.Set([(d_1242_v12_) * (d_1230_v2_)]))
                return _dafny.Set(coll44_)
            if ((iife110_()
            ) - ((self).f15)) == (((self).f15) - ((self).f15)):
                d_1243_v13_: D3
                d_1243_v13_ = D3_DC6(d_1234_i1_, d_1229_v1_.f19)
                d_1244_v14_: _dafny.Array
                nw170_ = _dafny.Array(None, 29)
                nw170_[int(0)] = (d_1229_v1_).f20
                nw170_[int(1)] = (d_1229_v1_).f20
                nw170_[int(2)] = (d_1229_v1_).f20
                nw170_[int(3)] = (d_1229_v1_).f20
                nw170_[int(4)] = (d_1229_v1_).f20
                nw170_[int(5)] = (d_1229_v1_).f20
                nw170_[int(6)] = d_1228_v0_
                nw170_[int(7)] = (d_1229_v1_).f20
                nw170_[int(8)] = (d_1229_v1_).f20
                nw170_[int(9)] = _dafny.CodePoint('y')
                nw170_[int(10)] = d_1228_v0_
                nw170_[int(11)] = (d_1229_v1_).f20
                nw170_[int(12)] = (d_1229_v1_).f20
                nw170_[int(13)] = (d_1229_v1_).f20
                nw170_[int(14)] = (d_1229_v1_).f20
                nw170_[int(15)] = (d_1229_v1_).f20
                nw170_[int(16)] = (d_1229_v1_).f20
                nw170_[int(17)] = default__.fm18(d_1230_v2_, d_1231_v3_, (self).f17, _dafny.Map({(self).f13: (0) - (d_1234_i1_)}), globalState)
                nw170_[int(18)] = default__.fm33((d_1243_v13_).cf13, d_1234_i1_, globalState)
                nw170_[int(19)] = (d_1229_v1_).f20
                nw170_[int(20)] = _dafny.CodePoint('n')
                nw170_[int(21)] = (d_1229_v1_).f20
                nw170_[int(22)] = (d_1229_v1_).f20
                nw170_[int(23)] = (d_1229_v1_).f20
                nw170_[int(24)] = d_1228_v0_
                nw170_[int(25)] = (d_1231_v3_)[default__.safeIndex(d_1234_i1_, len(d_1231_v3_))]
                nw170_[int(26)] = (d_1229_v1_).f20
                nw170_[int(27)] = (d_1229_v1_).f20
                nw170_[int(28)] = (d_1229_v1_).f20
                d_1244_v14_ = nw170_
                d_1245_v15_: D3
                d_1245_v15_ = D3_DC5((d_1229_v1_).f20)
                index193_ = default__.safeIndex(287, (d_1244_v14_).length(0))
                (d_1244_v14_)[index193_] = (d_1245_v15_).cf11
                d_1246_v16_: _dafny.Seq
                d_1246_v16_ = _dafny.SeqWithoutIsStrInference([d_1229_v1_.f19, not((self).f17), default__.fm30(d_1230_v2_, globalState)])
                d_1247_v17_: _dafny.MultiSet
                d_1247_v17_ = _dafny.MultiSet([(self).f17])
                d_1248_v18_: _dafny.Map
                d_1248_v18_ = _dafny.Map({d_1229_v1_.f19: (d_1247_v17_).cardinality})
                d_1249_v19_: _dafny.Seq
                d_1249_v19_ = _dafny.SeqWithoutIsStrInference([(self).f18, (self).f13, not(d_1229_v1_.f19), (self).f13, (d_1246_v16_)[default__.safeIndex(len(d_1248_v18_), len(d_1246_v16_))]])
                d_1250_v20_: _dafny.Seq
                d_1250_v20_ = _dafny.SeqWithoutIsStrInference([d_1246_v16_, d_1249_v19_])
                index194_ = default__.safeIndex(287, (d_1244_v14_).length(0))
                rhs161_ = (d_1246_v16_) + ((d_1250_v20_)[default__.safeIndex(d_1234_i1_, len(d_1250_v20_))])
                rhs162_ = (d_1240_v10_).cf39
                rhs163_ = d_1228_v0_
                rhs164_ = (d_1229_v1_).f20
                lhs133_ = globalState
                lhs134_ = d_1229_v1_
                lhs135_ = d_1244_v14_
                lhs136_ = default__.safeIndex(287, (d_1244_v14_).length(0))
                lhs133_.f2 = rhs161_
                lhs134_.f19 = rhs162_
                lhs135_[lhs136_] = rhs163_
                d_1228_v0_ = rhs164_
                (d_1229_v1_).f19 = (self).f17
                d_1251_v21_: _dafny.Map
                d_1251_v21_ = _dafny.Map({d_1234_i1_: d_1229_v1_.f19})
                (d_1229_v1_).f19 = ((d_1251_v21_)[len(_dafny.Map({len(d_1249_v19_): d_1234_i1_}))] if (len(_dafny.Map({len(d_1249_v19_): d_1234_i1_}))) in (d_1251_v21_) else (d_1237_v7_) == (_dafny.SeqWithoutIsStrInference([len(d_1231_v3_) for d_1252_i2_ in range(default__.abs(230))])))
                index195_ = default__.safeIndex(78, ((self).f14).length(0))
                ((self).f14)[index195_] = d_1229_v1_.f19
                index196_ = default__.safeIndex(78, ((self).f14).length(0))
                ((self).f14)[index196_] = not (not((d_1230_v2_) <= (509))) or ((d_1246_v16_)[default__.safeIndex(d_1230_v2_, len(d_1246_v16_))])
                (d_1229_v1_).f19 = ((self).f18) and (d_1229_v1_.f19)
            elif True:
                (d_1229_v1_).f19 = (self.f12).cf4
                (globalState).f7 = ((d_1234_i1_) + (len(d_1231_v3_))) * (d_1230_v2_)
                d_1253_v22_: _dafny.Seq
                d_1253_v22_ = _dafny.SeqWithoutIsStrInference([d_1231_v3_])
                rhs165_ = (d_1231_v3_) + (d_1231_v3_)
                rhs166_ = d_1234_i1_
                rhs167_ = (d_1253_v22_) + (_dafny.SeqWithoutIsStrInference([d_1231_v3_ for d_1254_i3_ in range(default__.abs(-342))]))
                lhs137_ = globalState
                d_1231_v3_ = rhs165_
                lhs137_.f7 = rhs166_
                d_1253_v22_ = rhs167_
                (d_1229_v1_).f19 = d_1229_v1_.f19
                d_1231_v3_ = d_1231_v3_

    @property
    def f18(self):
        return self._f18

class C4:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    def ctor__(self):
        pass
        pass

    def fm14(self, p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bq"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_1255_i0_ in range(default__.abs(963))]))

    def m10(self, p0, p1, p2, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: _dafny.Seq = _dafny.Seq({})
        d_1256_v0_: _dafny.Seq
        d_1256_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gwx"))
        d_1256_v0_ = d_1256_v0_
        d_1257_v1_: _dafny.Array
        nw171_ = _dafny.Array(int(0), 15)
        d_1257_v1_ = nw171_
        d_1258_v2_: _dafny.Seq
        d_1258_v2_ = _dafny.SeqWithoutIsStrInference([d_1257_v1_])
        hi7_ = len(d_1258_v2_)
        for d_1259_i0_ in range((450 if p0 else -597), hi7_):
            d_1260_v3_: _dafny.Seq
            d_1260_v3_ = _dafny.SeqWithoutIsStrInference([d_1259_i0_])
            if (_dafny.SeqWithoutIsStrInference([-808])) < (d_1260_v3_):
                d_1261_v4_: bool
                d_1261_v4_ = False
                d_1261_v4_ = d_1261_v4_
                d_1262_v5_: _dafny.Array
                nw172_ = _dafny.Array(False, 29)
                d_1262_v5_ = nw172_
                d_1262_v5_ = d_1262_v5_
                d_1263_v6_: D1
                d_1263_v6_ = D1_DC2(130, d_1259_i0_, p0)
                d_1264_v7_: C1
                nw173_ = C1()
                nw173_.ctor__(p2, d_1263_v6_, p0)
                d_1264_v7_ = nw173_
                index197_ = default__.safeIndex(645, (d_1262_v5_).length(0))
                (d_1262_v5_)[index197_] = not(not((p0) not in (_dafny.SeqWithoutIsStrInference([not(d_1261_v4_)]))))
                d_1265_v8_: _dafny.Set
                d_1265_v8_ = _dafny.Set({d_1261_v4_})
                index198_ = default__.safeIndex(645, (d_1262_v5_).length(0))
                rhs168_ = (len(d_1265_v8_)) >= (len(d_1256_v0_))
                rhs169_ = d_1256_v0_
                rhs170_ = (d_1260_v3_)[default__.safeIndex(d_1259_i0_, len(d_1260_v3_))]
                lhs138_ = d_1262_v5_
                lhs139_ = default__.safeIndex(645, (d_1262_v5_).length(0))
                lhs140_ = globalState
                lhs138_[lhs139_] = rhs168_
                d_1256_v0_ = rhs169_
                lhs140_.f7 = rhs170_
                index199_ = default__.safeIndex(645, (d_1262_v5_).length(0))
                (d_1262_v5_)[index199_] = not(p0)
            elif True:
                d_1266_v9_: _dafny.Seq
                d_1266_v9_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1267_v10_: D1
                d_1267_v10_ = D1_DC1(d_1256_v0_)
                d_1268_v11_: _dafny.Array
                nw174_ = _dafny.Array(False, 15)
                d_1268_v11_ = nw174_
                d_1269_v12_: _dafny.Set
                d_1269_v12_ = _dafny.Set({d_1259_i0_})
                pat_let_tv33_ = d_1256_v0_
                d_1270_v13_: C2
                nw175_ = C2()
                def iife111_(_pat_let33_0):
                    def iife112_(d_1271_dt__update__tmp_h0_):
                        def iife113_(_pat_let34_0):
                            def iife114_(d_1272_dt__update_hcf1_h0_):
                                return D1_DC1(d_1272_dt__update_hcf1_h0_)
                            return iife114_(_pat_let34_0)
                        return iife113_(pat_let_tv33_)
                    return iife112_(_pat_let33_0)
                nw175_.ctor__(d_1259_i0_, (d_1266_v9_)[default__.safeIndex(d_1259_i0_, len(d_1266_v9_))], iife111_(d_1267_v10_), p0, d_1268_v11_, d_1269_v12_, D1_DC2(d_1259_i0_, d_1259_i0_, p0), not (p0) or (p0))
                d_1270_v13_ = nw175_
                d_1273_v14_: _dafny.Set
                d_1273_v14_ = _dafny.Set({True})
                d_1274_v15_: D9
                d_1274_v15_ = D9_DC22(p0, (d_1270_v13_).f21, p0, d_1273_v14_, _dafny.SeqWithoutIsStrInference([p2 for d_1275_i1_ in range(default__.abs(172))]))
                d_1276_v16_: _dafny.Set
                d_1276_v16_ = _dafny.Set({d_1274_v15_, d_1274_v15_, D9_DC22(p0, (d_1270_v13_).f21, p0, d_1273_v14_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m")))})
                d_1277_v17_: D3
                d_1277_v17_ = D3_DC5(p2)
                d_1278_v18_: _dafny.Array
                nw176_ = _dafny.Array(None, 16)
                nw176_[int(0)] = p2
                nw176_[int(1)] = (d_1277_v17_).cf11
                nw176_[int(2)] = p2
                nw176_[int(3)] = p2
                nw176_[int(4)] = p2
                nw176_[int(5)] = p2
                nw176_[int(6)] = p2
                nw176_[int(7)] = p2
                nw176_[int(8)] = p2
                nw176_[int(9)] = _dafny.CodePoint('t')
                nw176_[int(10)] = p2
                nw176_[int(11)] = p2
                nw176_[int(12)] = _dafny.CodePoint('g')
                nw176_[int(13)] = p2
                nw176_[int(14)] = p2
                nw176_[int(15)] = p2
                d_1278_v18_ = nw176_
                d_1279_v19_: _dafny.Map
                d_1279_v19_ = _dafny.Map({d_1278_v18_: len(d_1260_v3_)})
                d_1280_v20_: _dafny.MultiSet
                d_1280_v20_ = _dafny.MultiSet([len(d_1276_v16_), ((d_1279_v19_)[d_1278_v18_] if (d_1278_v18_) in (d_1279_v19_) else d_1259_i0_)])
                d_1281_v21_: D7
                d_1281_v21_ = D7_DC16((d_1270_v13_).f21)
                d_1282_v22_: _dafny.Map
                d_1282_v22_ = _dafny.Map({p0: d_1280_v20_})
                d_1283_v23_: D6
                d_1283_v23_ = D6_DC11(_dafny.MultiSet([d_1259_i0_, len(d_1256_v0_), (d_1270_v13_).f21]))
                d_1284_v24_: _dafny.Array
                nw177_ = _dafny.Array(None, 22)
                nw177_[int(0)] = d_1280_v20_
                nw177_[int(1)] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_1259_i0_, len(d_1273_v14_)]))
                nw177_[int(2)] = d_1280_v20_
                nw177_[int(3)] = (d_1280_v20_).intersection(d_1280_v20_)
                nw177_[int(4)] = (d_1280_v20_) | (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([981, d_1259_i0_, len(_dafny.SeqWithoutIsStrInference([not((d_1270_v13_).f22)]))])), (d_1281_v21_).cf29]))
                nw177_[int(5)] = d_1280_v20_
                nw177_[int(6)] = d_1280_v20_
                nw177_[int(7)] = d_1280_v20_
                nw177_[int(8)] = d_1280_v20_
                nw177_[int(9)] = d_1280_v20_
                nw177_[int(10)] = d_1280_v20_
                nw177_[int(11)] = (_dafny.MultiSet([d_1259_i0_])) - (d_1280_v20_)
                nw177_[int(12)] = d_1280_v20_
                nw177_[int(13)] = ((d_1282_v22_)[(d_1270_v13_).f22] if ((d_1270_v13_).f22) in (d_1282_v22_) else d_1280_v20_)
                nw177_[int(14)] = d_1280_v20_
                nw177_[int(15)] = (d_1280_v20_) | ((d_1280_v20_).set((d_1280_v20_).cardinality, default__.abs(473)))
                nw177_[int(16)] = d_1280_v20_
                nw177_[int(17)] = (d_1280_v20_).intersection(d_1280_v20_)
                nw177_[int(18)] = (d_1280_v20_).intersection(d_1280_v20_)
                nw177_[int(19)] = d_1280_v20_
                nw177_[int(20)] = default__.fm19(globalState)
                nw177_[int(21)] = ((d_1283_v23_).cf17) - (_dafny.MultiSet(d_1260_v3_))
                d_1284_v24_ = nw177_
                index200_ = default__.safeIndex(506, (d_1284_v24_).length(0))
                (d_1284_v24_)[index200_] = d_1280_v20_
                index201_ = default__.safeIndex(506, (d_1284_v24_).length(0))
                (d_1284_v24_)[index201_] = d_1280_v20_
                d_1285_v25_: bool
                d_1285_v25_ = False
                d_1286_v26_: _dafny.Array
                nw178_ = _dafny.Array(_dafny.Array(None, 0), 28)
                d_1286_v26_ = nw178_
                d_1287_v27_: _dafny.Array
                nw179_ = _dafny.Array(_dafny.Set({}), 23)
                d_1287_v27_ = nw179_
                index202_ = default__.safeIndex(460, (d_1286_v26_).length(0))
                (d_1286_v26_)[index202_] = d_1287_v27_
                index203_ = default__.safeIndex(743, (d_1268_v11_).length(0))
                (d_1268_v11_)[index203_] = (d_1270_v13_).f22
                d_1288_v28_: _dafny.Map
                d_1288_v28_ = _dafny.Map({(d_1259_i0_) - (-417): d_1287_v27_})
                index204_ = default__.safeIndex(460, (d_1286_v26_).length(0))
                index205_ = default__.safeIndex(743, (d_1268_v11_).length(0))
                rhs171_ = (p2) in (d_1256_v0_)
                rhs172_ = ((d_1288_v28_)[183] if (183) in (d_1288_v28_) else d_1287_v27_)
                rhs173_ = False
                rhs174_ = d_1257_v1_
                lhs141_ = d_1286_v26_
                lhs142_ = default__.safeIndex(460, (d_1286_v26_).length(0))
                lhs143_ = d_1268_v11_
                lhs144_ = default__.safeIndex(743, (d_1268_v11_).length(0))
                d_1285_v25_ = rhs171_
                lhs141_[lhs142_] = rhs172_
                lhs143_[lhs144_] = rhs173_
                d_1257_v1_ = rhs174_
                d_1289_v29_: _dafny.Map
                d_1289_v29_ = _dafny.Map({(d_1270_v13_).f22: (d_1268_v11_)[default__.safeIndex(743, (d_1268_v11_).length(0))]})
                index206_ = default__.safeIndex(743, (d_1268_v11_).length(0))
                index207_ = default__.safeIndex(743, (d_1268_v11_).length(0))
                rhs175_ = (d_1268_v11_)[default__.safeIndex(743, (d_1268_v11_).length(0))]
                rhs176_ = d_1280_v20_
                rhs177_ = not(((d_1270_v13_).f21) <= ((d_1270_v13_).f21))
                rhs178_ = not(((d_1259_i0_) not in (d_1260_v3_) if (d_1270_v13_).f22 else ((d_1289_v29_)[d_1285_v25_] if (d_1285_v25_) in (d_1289_v29_) else (d_1268_v11_)[default__.safeIndex(743, (d_1268_v11_).length(0))])))
                lhs145_ = d_1268_v11_
                lhs146_ = default__.safeIndex(743, (d_1268_v11_).length(0))
                lhs147_ = d_1268_v11_
                lhs148_ = default__.safeIndex(743, (d_1268_v11_).length(0))
                lhs145_[lhs146_] = rhs175_
                d_1280_v20_ = rhs176_
                lhs147_[lhs148_] = rhs177_
                d_1285_v25_ = rhs178_
                d_1290_v30_: _dafny.MultiSet
                d_1290_v30_ = _dafny.MultiSet([(d_1270_v13_).f22])
                d_1291_v31_: D1
                d_1291_v31_ = D1_DC2(236, (d_1270_v13_).f21, (d_1270_v13_).f22)
                d_1292_v32_: _dafny.Set
                d_1292_v32_ = _dafny.Set({_dafny.CodePoint('c'), p2, p2})
                pat_let_tv34_ = d_1256_v0_
                d_1293_v33_: C3
                nw180_ = C3()
                def iife115_(_pat_let35_0):
                    def iife116_(d_1294_dt__update__tmp_h1_):
                        def iife117_(_pat_let36_0):
                            def iife118_(d_1295_dt__update_hcf1_h1_):
                                return D1_DC1(d_1295_dt__update_hcf1_h1_)
                            return iife118_(_pat_let36_0)
                        return iife117_(pat_let_tv34_)
                    return iife116_(_pat_let35_0)
                nw180_.ctor__((d_1290_v30_).isdisjoint(d_1290_v30_), iife115_(d_1267_v10_), (d_1270_v13_).f22, d_1268_v11_, d_1269_v12_, d_1291_v31_, (_dafny.Set({default__.fm33((d_1268_v11_)[default__.safeIndex(743, (d_1268_v11_).length(0))], (d_1270_v13_).f21, globalState)})).issubset(d_1292_v32_))
                d_1293_v33_ = nw180_
            (globalState).f7 = (d_1259_i0_ if p0 else d_1259_i0_)
            d_1296_v34_: D6
            d_1296_v34_ = D6_DC12(p0, d_1256_v0_, d_1259_i0_, p0)
            d_1297_v35_: _dafny.Set
            d_1297_v35_ = _dafny.Set({235, d_1259_i0_})
            d_1298_v36_: _dafny.Array
            nw181_ = _dafny.Array(None, 4)
            nw181_[int(0)] = p0
            nw181_[int(1)] = (d_1296_v34_).cf21
            nw181_[int(2)] = p0
            nw181_[int(3)] = (d_1297_v35_).ispropersubset(d_1297_v35_)
            d_1298_v36_ = nw181_
            index208_ = default__.safeIndex(230, (d_1298_v36_).length(0))
            (d_1298_v36_)[index208_] = not(p0)
            index209_ = default__.safeIndex(230, (d_1298_v36_).length(0))
            (d_1298_v36_)[index209_] = not((p0) or (p0))
            d_1299_v37_: D1
            d_1299_v37_ = D1_DC2(d_1259_i0_, d_1259_i0_, (d_1298_v36_)[default__.safeIndex(230, (d_1298_v36_).length(0))])
            d_1300_v38_: C1
            nw182_ = C1()
            nw182_.ctor__(p2, d_1299_v37_, p0)
            d_1300_v38_ = nw182_
        d_1301_v39_: _dafny.Set
        d_1301_v39_ = _dafny.Set({p0})
        d_1302_v40_: int
        d_1302_v40_ = 706
        rhs179_ = d_1301_v39_
        rhs180_ = d_1302_v40_
        rhs181_ = d_1302_v40_
        lhs149_ = globalState
        lhs150_ = globalState
        d_1301_v39_ = rhs179_
        lhs149_.f7 = rhs180_
        lhs150_.f7 = rhs181_
        d_1303_v41_: D1
        d_1303_v41_ = D1_DC1(d_1256_v0_)
        d_1304_v42_: _dafny.Array
        def lambda53_(d_1305_p0_):
            def lambda54_(d_1306_i2_):
                return d_1305_p0_

            return lambda54_

        init30_ = lambda53_(p0)
        nw183_ = _dafny.Array(None, 4)
        for i0_30_ in range(nw183_.length(0)):
            nw183_[i0_30_] = init30_(i0_30_)
        d_1304_v42_ = nw183_
        d_1307_v43_: _dafny.Set
        d_1307_v43_ = _dafny.Set({d_1302_v40_, d_1302_v40_})
        d_1308_v44_: D1
        d_1308_v44_ = D1_DC2(d_1302_v40_, d_1302_v40_, not(p0))
        d_1309_v45_: C3
        nw184_ = C3()
        nw184_.ctor__((p0) not in (d_1301_v39_), d_1303_v41_, (p0 if p0 else not(p0)), d_1304_v42_, d_1307_v43_, d_1308_v44_, True)
        d_1309_v45_ = nw184_
        d_1310_v46_: bool
        d_1310_v46_ = True
        d_1310_v46_ = p0
        hi8_ = 881
        for d_1311_i3_ in range(d_1302_v40_, hi8_):
            (globalState).f7 = d_1311_i3_
            d_1312_v47_: str
            d_1312_v47_ = _dafny.CodePoint('f')
            d_1312_v47_ = p2
            d_1313_v48_: _dafny.Array
            nw185_ = _dafny.Array(_dafny.MultiSet({}), 12)
            d_1313_v48_ = nw185_
            rhs182_ = d_1313_v48_
            rhs183_ = p0
            rhs184_ = (d_1256_v0_) < (d_1256_v0_)
            d_1313_v48_ = rhs182_
            d_1310_v46_ = rhs183_
            d_1310_v46_ = rhs184_
            if (d_1309_v45_).f18:
                d_1314_v49_: _dafny.Set
                d_1314_v49_ = _dafny.Set({p2, p2})
                d_1314_v49_ = d_1314_v49_
                (globalState).f7 = d_1302_v40_
                d_1310_v46_ = d_1310_v46_
                d_1315_v50_: _dafny.Seq
                d_1315_v50_ = _dafny.SeqWithoutIsStrInference([d_1310_v46_, False])
                d_1316_v51_: _dafny.Seq
                d_1316_v51_ = _dafny.SeqWithoutIsStrInference([(d_1302_v40_) != (len(d_1315_v50_)), (d_1309_v45_).f18, False])
                (globalState).f2 = d_1316_v51_
                d_1317_v52_: _dafny.MultiSet
                d_1317_v52_ = _dafny.MultiSet([d_1257_v1_])
                d_1318_v53_: C2
                nw186_ = C2()
                nw186_.ctor__(d_1311_i3_, not (not((d_1309_v45_).f18)) or ((d_1309_v45_).f18), d_1303_v41_, (d_1316_v51_)[default__.safeIndex(d_1311_i3_, len(d_1316_v51_))], d_1304_v42_, _dafny.Set({d_1302_v40_}), d_1308_v44_, (d_1317_v52_).ispropersubset(d_1317_v52_))
                d_1318_v53_ = nw186_
            elif True:
                (globalState).f7 = 75
                d_1310_v46_ = (d_1309_v45_).f18
                d_1319_v54_: C0
                nw187_ = C0()
                nw187_.ctor__((d_1309_v45_).f18, p2, D1_DC2(d_1302_v40_, d_1311_i3_, (d_1309_v45_).f18), not(default__.fm17(globalState)))
                d_1319_v54_ = nw187_
                d_1320_v55_: _dafny.Array
                nw188_ = _dafny.Array(_dafny.Map({}), 11)
                d_1320_v55_ = nw188_
                d_1320_v55_ = d_1320_v55_
                d_1321_v56_: _dafny.Map
                d_1321_v56_ = _dafny.Map({d_1319_v54_.f19: d_1302_v40_})
                d_1322_v57_: _dafny.Map
                d_1322_v57_ = _dafny.Map({d_1321_v56_: -568})
                d_1323_v58_: _dafny.Map
                d_1323_v58_ = _dafny.Map({d_1322_v57_: d_1310_v46_})
                d_1323_v58_ = (d_1323_v58_).set(d_1322_v57_, (d_1309_v45_).f18)
        r0 = p1
        r1 = d_1258_v2_
        return r0, r1


class C5(T2, T1, T0):
    def  __init__(self):
        self._f12: D1 = D1.default()()
        self._f14: _dafny.Array = _dafny.Array(None, 0)
        self._f15: _dafny.Set = _dafny.Set({})
        self._f13: bool = False
        self.f24: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f12(self):
        return self._f12
    @f12.setter
    def f12(self, value):
        self._f12 = value
    @property
    def f14(self):
        return self._f14
    @property
    def f15(self):
        return self._f15
    @property
    def f13(self):
        return self._f13
    def ctor__(self, f24, f14, f15, f12, f13):
        (self).f24 = f24
        (self)._f14 = f14
        (self)._f15 = f15
        (self).f12 = f12
        (self)._f13 = f13

    def fm10(self, p0, p1, p2, globalState):
        return (D8_DC19(_dafny.Map({(self).f13: len(_dafny.Map({-193: 525}))}), _dafny.Set({False, self.f24, (D12_DC32(self.f24, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wbv"))), False, self.f24, len(_dafny.Map({(self).f13: 732})))).cf67, self.f24, self.f24}), _dafny.MultiSet([144, (0) - ((_dafny.MultiSet([-496, 687])).cardinality), len(_dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f13: -869})])), len(_dafny.Map({342: self.f24})), (_dafny.MultiSet([(self).f13, self.f24])).cardinality]), D1_DC1(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_1324_i0_ in range(default__.abs(113))])))).cf32

    def fm8(self, p0, p1, globalState):
        return len(((self).f15) - ((self).f15))

    def fm9(self, p0, globalState):
        return len(_dafny.Map({(D12_DC29(_dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f13]): (self).f13})) if self.f24 else D12_DC29(_dafny.Map({_dafny.SeqWithoutIsStrInference([True]): self.f24}))): ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([D7_DC17(D7_DC17(D7_DC17(D7_DC16(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_1325_i0_ in range(default__.abs(460))])))))), D7_DC17(D7_DC15(_dafny.SeqWithoutIsStrInference([324])))])])) | (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([D7_DC17(D7_DC16(len(_dafny.Set({self.f24, (self).f13})))), D7_DC17(D7_DC15(_dafny.SeqWithoutIsStrInference([-703, 177]))), D7_DC17(D7_DC17(D7_DC17(D7_DC17(D7_DC15(_dafny.SeqWithoutIsStrInference([387 for d_1326_i1_ in range(default__.abs(-543))]))))))]), _dafny.SeqWithoutIsStrInference([D7_DC17(D7_DC16((0) - ((0) - (len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x')]))}))))))])]))).cardinality}))

    def fm52(self, p0, p1, globalState):
        return (731) <= (default__.safeDivisionInt(459, len(_dafny.SeqWithoutIsStrInference([-683]))))

    def fm53(self, p0, globalState):
        return (0) - ((347) + ((111) + (19)))

    def m5(self, p0, globalState):
        r0: bool = False
        index210_ = default__.safeIndex(855, ((self).f14).length(0))
        ((self).f14)[index210_] = False
        d_1327_v0_: _dafny.Seq
        d_1327_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "av"))
        d_1328_v1_: _dafny.Seq
        d_1328_v1_ = _dafny.SeqWithoutIsStrInference([False, (self).f13])
        d_1329_v2_: _dafny.Map
        d_1329_v2_ = _dafny.Map({self.f24: p0})
        d_1330_v3_: _dafny.Array
        nw189_ = _dafny.Array(None, 14)
        nw189_[int(0)] = (0) - (p0)
        nw189_[int(1)] = (self).fm8(self.f24, D1_DC1(d_1327_v0_), globalState)
        nw189_[int(2)] = default__.safeDivisionInt(p0, p0)
        nw189_[int(3)] = (len(d_1328_v1_)) * (len(d_1329_v2_))
        nw189_[int(4)] = (p0) * (p0)
        nw189_[int(5)] = p0
        nw189_[int(6)] = 370
        nw189_[int(7)] = p0
        nw189_[int(8)] = p0
        nw189_[int(9)] = p0
        nw189_[int(10)] = p0
        nw189_[int(11)] = p0
        nw189_[int(12)] = p0
        nw189_[int(13)] = p0
        d_1330_v3_ = nw189_
        index211_ = default__.safeIndex(773, (d_1330_v3_).length(0))
        (d_1330_v3_)[index211_] = (p0 if self.f24 else len(d_1329_v2_))
        d_1331_v4_: _dafny.Array
        nw190_ = _dafny.Array(D12.default()(), 12)
        d_1331_v4_ = nw190_
        index212_ = default__.safeIndex(969, (d_1331_v4_).length(0))
        (d_1331_v4_)[index212_] = D12_DC30((self).f14, (self).f13)
        d_1332_v5_: str
        d_1332_v5_ = _dafny.CodePoint('f')
        d_1333_v6_: _dafny.MultiSet
        d_1333_v6_ = _dafny.MultiSet([d_1332_v5_])
        d_1334_v7_: _dafny.MultiSet
        d_1334_v7_ = _dafny.MultiSet([False, False, self.f24, (D12_DC31((self).f13, d_1330_v3_, (self).f13, _dafny.CodePoint('l'), d_1333_v6_)).cf62])
        d_1335_v8_: _dafny.MultiSet
        d_1335_v8_ = _dafny.MultiSet([d_1330_v3_])
        d_1336_v9_: _dafny.Seq
        d_1336_v9_ = _dafny.SeqWithoutIsStrInference([d_1335_v8_, d_1335_v8_])
        index213_ = default__.safeIndex(855, ((self).f14).length(0))
        index214_ = default__.safeIndex(773, (d_1330_v3_).length(0))
        index215_ = default__.safeIndex(969, (d_1331_v4_).length(0))
        rhs185_ = p0
        rhs186_ = not(not(not (False) or ((self.f24 if (self).f13 else self.f24))))
        rhs187_ = default__.safeModuloInt(len(d_1328_v1_), ((d_1334_v7_ if self.f24 else d_1334_v7_)).cardinality)
        rhs188_ = D12_DC30(((self).f14 if (self).f13 else (self).f14), ((d_1336_v9_)[default__.safeIndex(p0, len(d_1336_v9_))]).ispropersubset(d_1335_v8_))
        lhs151_ = globalState
        lhs152_ = (self).f14
        lhs153_ = default__.safeIndex(855, ((self).f14).length(0))
        lhs154_ = d_1330_v3_
        lhs155_ = default__.safeIndex(773, (d_1330_v3_).length(0))
        lhs156_ = d_1331_v4_
        lhs157_ = default__.safeIndex(969, (d_1331_v4_).length(0))
        lhs151_.f7 = rhs185_
        lhs152_[lhs153_] = rhs186_
        lhs154_[lhs155_] = rhs187_
        lhs156_[lhs157_] = rhs188_
        d_1337_v10_: _dafny.Seq
        d_1337_v10_ = _dafny.SeqWithoutIsStrInference([p0])
        d_1338_v11_: D1
        d_1338_v11_ = D1_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rkrn")))
        d_1339_v12_: _dafny.Array
        nw191_ = _dafny.Array(False, 13)
        d_1339_v12_ = nw191_
        pat_let_tv35_ = d_1332_v5_
        d_1340_v13_: C2
        nw192_ = C2()
        def iife119_(_pat_let37_0):
            def iife120_(d_1341_dt__update__tmp_h0_):
                def iife121_(_pat_let38_0):
                    def iife122_(d_1343_dt__update_hcf1_h0_):
                        return D1_DC1(d_1343_dt__update_hcf1_h0_)
                    return iife122_(_pat_let38_0)
                return iife121_(_dafny.SeqWithoutIsStrInference([pat_let_tv35_ for d_1342_i0_ in range(default__.abs(629))]))
            return iife120_(_pat_let37_0)
        nw192_.ctor__((len(d_1337_v10_)) + ((self.f12).cf3), self.f24, iife119_(d_1338_v11_), not(((d_1330_v3_)[default__.safeIndex(773, (d_1330_v3_).length(0))]) < ((0) - (len(d_1327_v0_)))), d_1339_v12_, (self).f15, self.f12, (self).f13)
        d_1340_v13_ = nw192_
        d_1344_v14_: D9
        d_1344_v14_ = D9_DC22((self).f13, (d_1330_v3_)[default__.safeIndex(773, (d_1330_v3_).length(0))], (self).f13, default__.fm54(len(_dafny.SeqWithoutIsStrInference([len((d_1327_v0_).set(default__.safeIndex((d_1340_v13_).f21, len(d_1327_v0_)), (d_1327_v0_)[default__.safeIndex((d_1340_v13_).f21, len(d_1327_v0_))])) for d_1345_i1_ in range(default__.abs(623))])), globalState), d_1327_v0_)
        r0 = (d_1344_v14_).cf37
        index216_ = default__.safeIndex(773, (d_1330_v3_).length(0))
        (d_1330_v3_)[index216_] = p0
        index217_ = default__.safeIndex(773, (d_1330_v3_).length(0))
        (d_1330_v3_)[index217_] = (d_1330_v3_)[default__.safeIndex(773, (d_1330_v3_).length(0))]
        d_1346_v15_: _dafny.Seq
        d_1346_v15_ = _dafny.SeqWithoutIsStrInference([d_1339_v12_, d_1339_v12_])
        d_1347_v16_: _dafny.Set
        d_1347_v16_ = _dafny.Set({d_1339_v12_, (d_1346_v15_)[default__.safeIndex(753, len(d_1346_v15_))], d_1339_v12_, d_1339_v12_, d_1339_v12_})
        index218_ = default__.safeIndex(773, (d_1330_v3_).length(0))
        rhs189_ = d_1347_v16_
        rhs190_ = default__.safeModuloInt(len(d_1328_v1_), (d_1330_v3_)[default__.safeIndex(773, (d_1330_v3_).length(0))])
        rhs191_ = (d_1330_v3_)[default__.safeIndex(773, (d_1330_v3_).length(0))]
        lhs158_ = globalState
        lhs159_ = d_1330_v3_
        lhs160_ = default__.safeIndex(773, (d_1330_v3_).length(0))
        d_1347_v16_ = rhs189_
        lhs158_.f7 = rhs190_
        lhs159_[lhs160_] = rhs191_
        r0 = (d_1340_v13_).f22
        return r0

    def m14(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        d_1348_v0_: int
        out10_: int
        out10_ = (self).m15(globalState)
        d_1348_v0_ = out10_
        d_1349_i0_: int
        d_1349_i0_ = 0
        with _dafny.label("6"):
            while self.f24:
                with _dafny.c_label("6"):
                    if (d_1349_i0_) >= (100):
                        raise _dafny.Break("6")
                    d_1349_i0_ = (d_1349_i0_) + (1)
                    d_1350_v1_: str
                    d_1350_v1_ = _dafny.CodePoint('e')
                    d_1351_v2_: D3
                    d_1351_v2_ = D3_DC5(d_1350_v1_)
                    index219_ = default__.safeIndex(698, ((self).f14).length(0))
                    ((self).f14)[index219_] = self.f24
                    d_1352_v3_: _dafny.Map
                    d_1352_v3_ = _dafny.Map({d_1350_v1_: 78})
                    d_1353_v4_: _dafny.Set
                    d_1353_v4_ = _dafny.Set({d_1352_v3_})
                    d_1354_v5_: _dafny.Array
                    nw193_ = _dafny.Array(None, 1)
                    nw193_[int(0)] = (default__.fm55((self).fm9(d_1350_v1_, globalState), d_1353_v4_, p0, globalState)).cardinality
                    d_1354_v5_ = nw193_
                    index220_ = default__.safeIndex(641, (d_1354_v5_).length(0))
                    (d_1354_v5_)[index220_] = d_1348_v0_
                    index221_ = default__.safeIndex(698, ((self).f14).length(0))
                    index222_ = default__.safeIndex(641, (d_1354_v5_).length(0))
                    rhs192_ = d_1351_v2_
                    rhs193_ = (True if p0 else (self).f13)
                    rhs194_ = d_1348_v0_
                    lhs161_ = (self).f14
                    lhs162_ = default__.safeIndex(698, ((self).f14).length(0))
                    lhs163_ = d_1354_v5_
                    lhs164_ = default__.safeIndex(641, (d_1354_v5_).length(0))
                    d_1351_v2_ = rhs192_
                    lhs161_[lhs162_] = rhs193_
                    lhs163_[lhs164_] = rhs194_
                    index223_ = default__.safeIndex(698, ((self).f14).length(0))
                    ((self).f14)[index223_] = (self).f13
                    d_1355_v6_: D11
                    d_1355_v6_ = D11_DC26(d_1354_v5_, p3)
                    d_1354_v5_ = ((d_1355_v6_ if p0 else d_1355_v6_)).cf49
                    r0 = (d_1354_v5_)[default__.safeIndex(641, (d_1354_v5_).length(0))]
                    pass
            pass
        d_1356_v7_: _dafny.Set
        d_1356_v7_ = _dafny.Set({p2})
        d_1357_v8_: _dafny.Set
        d_1357_v8_ = (d_1356_v7_)
        source21_ = d_1357_v8_
        d_1358___mcc_h0_ = source21_
        d_1359_cf16_ = d_1358___mcc_h0_
        d_1360_v9_: _dafny.MultiSet
        d_1360_v9_ = _dafny.MultiSet([d_1348_v0_, d_1348_v0_, d_1348_v0_])
        d_1361_v10_: _dafny.Array
        def lambda55_(d_1362_v0_):
            def lambda56_(d_1363_i1_):
                return (d_1363_i1_) * (d_1362_v0_)

            return lambda56_

        init31_ = lambda55_(d_1348_v0_)
        nw194_ = _dafny.Array(None, 6)
        for i0_31_ in range(nw194_.length(0)):
            nw194_[i0_31_] = init31_(i0_31_)
        d_1361_v10_ = nw194_
        d_1364_v11_: D11
        d_1364_v11_ = D11_DC26(d_1361_v10_, p3)
        d_1365_v12_: _dafny.Seq
        d_1365_v12_ = _dafny.SeqWithoutIsStrInference([d_1364_v11_, D11_DC26(d_1361_v10_, p3)])
        d_1360_v9_ = (_dafny.MultiSet([d_1348_v0_, d_1348_v0_]) if self.f24 else default__.fm56(len(d_1365_v12_), (self).f15, globalState))
        index224_ = default__.safeIndex(761, ((self).f14).length(0))
        ((self).f14)[index224_] = (self).f13
        index225_ = default__.safeIndex(761, ((self).f14).length(0))
        ((self).f14)[index225_] = p0
        d_1366_v13_: _dafny.Set
        d_1366_v13_ = _dafny.Set({D1_DC1(p3)})
        if (d_1366_v13_).isdisjoint(d_1366_v13_):
            d_1367_v14_: D3
            d_1367_v14_ = D3_DC6(d_1348_v0_, self.f24)
            d_1368_v15_: D7
            d_1368_v15_ = D7_DC16((d_1367_v14_).cf12)
            (globalState).f7 = (d_1368_v15_).cf29
            d_1369_v16_: _dafny.Set
            d_1369_v16_ = _dafny.Set({False, self.f24, p0})
            d_1370_v17_: D9
            d_1370_v17_ = D9_DC22((self).f13, d_1348_v0_, False, d_1369_v16_, p2)
            d_1371_v18_: _dafny.Map
            d_1371_v18_ = _dafny.Map({(d_1348_v0_) == (d_1348_v0_): ((self).f13) and ((d_1370_v17_).cf37)})
            d_1371_v18_ = (d_1371_v18_) | (d_1371_v18_)
            index226_ = default__.safeIndex(761, ((self).f14).length(0))
            ((self).f14)[index226_] = (default__.fm57(globalState)).cf65
            (self).f24 = not(not(not (True) or ((((self).f14)[default__.safeIndex(761, ((self).f14).length(0))] if p0 else (self).f13))))
            (globalState).f7 = default__.safeModuloInt(d_1348_v0_, d_1348_v0_)
        elif True:
            index227_ = default__.safeIndex(761, ((self).f14).length(0))
            ((self).f14)[index227_] = False
            d_1372_v19_: _dafny.MultiSet
            out11_: _dafny.MultiSet
            out11_ = default__.m0(_dafny.CodePoint('j'), p0, globalState)
            d_1372_v19_ = out11_
            d_1373_v20_: _dafny.Map
            d_1373_v20_ = _dafny.Map({d_1348_v0_: d_1361_v10_})
            index228_ = default__.safeIndex(721, (d_1361_v10_).length(0))
            (d_1361_v10_)[index228_] = len(d_1373_v20_)
            d_1374_v21_: _dafny.Map
            d_1374_v21_ = _dafny.Map({self.f24: p0})
            d_1375_v22_: _dafny.Seq
            d_1375_v22_ = _dafny.SeqWithoutIsStrInference([len(d_1374_v21_), len(_dafny.SeqWithoutIsStrInference([d_1361_v10_, d_1361_v10_]))])
            index229_ = default__.safeIndex(721, (d_1361_v10_).length(0))
            (d_1361_v10_)[index229_] = (len(d_1375_v22_)) + (d_1348_v0_)
            index230_ = default__.safeIndex(761, ((self).f14).length(0))
            ((self).f14)[index230_] = not((self).f13)
            d_1376_v23_: _dafny.Array
            nw195_ = _dafny.Array(_dafny.CodePoint('D'), 25)
            d_1376_v23_ = nw195_
            d_1377_v24_: str
            d_1377_v24_ = _dafny.CodePoint('o')
            index231_ = default__.safeIndex(192, (d_1376_v23_).length(0))
            (d_1376_v23_)[index231_] = d_1377_v24_
            d_1378_v25_: D12
            d_1378_v25_ = D12_DC32(not (self.f24) or (p0), (d_1361_v10_)[default__.safeIndex(721, (d_1361_v10_).length(0))], ((d_1361_v10_)[default__.safeIndex(721, (d_1361_v10_).length(0))]) < ((0) - (len(d_1375_v22_))), (self).f13, (d_1361_v10_)[default__.safeIndex(721, (d_1361_v10_).length(0))])
            d_1379_v26_: _dafny.Map
            d_1379_v26_ = _dafny.Map({default__.fm58(globalState): (self).f13})
            d_1380_v27_: _dafny.Array
            nw196_ = _dafny.Array(None, 6)
            nw196_[int(0)] = d_1379_v26_
            nw196_[int(1)] = _dafny.Map({d_1377_v24_: False})
            nw196_[int(2)] = (d_1379_v26_) | (d_1379_v26_)
            nw196_[int(3)] = (default__.fm59(self.f24, p0, (self).f13, d_1377_v24_, globalState)) | (d_1379_v26_)
            nw196_[int(4)] = d_1379_v26_
            nw196_[int(5)] = default__.fm59(p0, True, p0, d_1377_v24_, globalState)
            d_1380_v27_ = nw196_
            d_1381_v28_: _dafny.Seq
            d_1381_v28_ = _dafny.SeqWithoutIsStrInference([(self).f13, (self).f13, (self).f13, False, (self).f13])
            index232_ = default__.safeIndex(192, (d_1376_v23_).length(0))
            rhs195_ = (d_1377_v24_ if (len(d_1381_v28_)) < (d_1348_v0_) else d_1377_v24_)
            rhs196_ = _dafny.SeqWithoutIsStrInference([(d_1361_v10_)[default__.safeIndex(721, (d_1361_v10_).length(0))], (d_1348_v0_) - ((d_1361_v10_)[default__.safeIndex(721, (d_1361_v10_).length(0))]), default__.safeDivisionInt(d_1348_v0_, d_1348_v0_)])
            rhs197_ = default__.fm57(globalState)
            rhs198_ = (d_1361_v10_)[default__.safeIndex(721, (d_1361_v10_).length(0))]
            rhs199_ = d_1380_v27_
            lhs165_ = d_1376_v23_
            lhs166_ = default__.safeIndex(192, (d_1376_v23_).length(0))
            lhs165_[lhs166_] = rhs195_
            d_1375_v22_ = rhs196_
            d_1378_v25_ = rhs197_
            r0 = rhs198_
            d_1380_v27_ = rhs199_
        d_1382_v29_: _dafny.Map
        d_1382_v29_ = _dafny.Map({d_1361_v10_: ((self).f14)[default__.safeIndex(761, ((self).f14).length(0))]})
        d_1383_v30_: D14
        d_1383_v30_ = D14_DC36((d_1382_v29_).set(d_1361_v10_, (self).f13))
        d_1382_v29_ = (_dafny.Map({d_1361_v10_: ((self).f14)[default__.safeIndex(761, ((self).f14).length(0))]})) | ((d_1383_v30_).cf74)
        (self).f24 = (not(p0)) or (p0)
        d_1384_v31_: _dafny.Array
        def lambda57_(d_1385_i2_):
            return _dafny.CodePoint('o')

        init32_ = lambda57_
        nw197_ = _dafny.Array(None, 1)
        for i0_32_ in range(nw197_.length(0)):
            nw197_[i0_32_] = init32_(i0_32_)
        d_1384_v31_ = nw197_
        d_1386_v32_: str
        d_1386_v32_ = _dafny.CodePoint('b')
        d_1387_v33_: _dafny.Seq
        d_1387_v33_ = _dafny.SeqWithoutIsStrInference([d_1348_v0_, d_1348_v0_, d_1348_v0_, d_1348_v0_, d_1348_v0_])
        d_1388_v34_: D11
        d_1388_v34_ = D11_DC27(p0, d_1386_v32_, d_1348_v0_, len(p1), d_1387_v33_)
        nw198_ = _dafny.Array(None, 28)
        nw198_[int(0)] = d_1386_v32_
        nw198_[int(1)] = d_1386_v32_
        nw198_[int(2)] = d_1386_v32_
        nw198_[int(3)] = d_1386_v32_
        nw198_[int(4)] = d_1386_v32_
        nw198_[int(5)] = d_1386_v32_
        nw198_[int(6)] = (p1)[default__.safeIndex(len(p2), len(p1))]
        nw198_[int(7)] = d_1386_v32_
        nw198_[int(8)] = d_1386_v32_
        nw198_[int(9)] = d_1386_v32_
        nw198_[int(10)] = d_1386_v32_
        nw198_[int(11)] = d_1386_v32_
        nw198_[int(12)] = d_1386_v32_
        nw198_[int(13)] = d_1386_v32_
        nw198_[int(14)] = (d_1388_v34_).cf52
        nw198_[int(15)] = d_1386_v32_
        nw198_[int(16)] = d_1386_v32_
        nw198_[int(17)] = d_1386_v32_
        nw198_[int(18)] = d_1386_v32_
        nw198_[int(19)] = d_1386_v32_
        nw198_[int(20)] = d_1386_v32_
        nw198_[int(21)] = d_1386_v32_
        nw198_[int(22)] = _dafny.CodePoint('i')
        nw198_[int(23)] = d_1386_v32_
        nw198_[int(24)] = d_1386_v32_
        nw198_[int(25)] = d_1386_v32_
        nw198_[int(26)] = d_1386_v32_
        nw198_[int(27)] = d_1386_v32_
        d_1384_v31_ = nw198_
        d_1389_v35_: _dafny.Seq
        d_1389_v35_ = _dafny.SeqWithoutIsStrInference([p0, p0])
        (globalState).f2 = d_1389_v35_
        r0 = d_1348_v0_
        return r0

    def m15(self, globalState):
        r0: int = int(0)
        d_1390_v0_: _dafny.MultiSet
        d_1390_v0_ = _dafny.MultiSet([(self).f13])
        d_1391_v1_: _dafny.Seq
        d_1391_v1_ = _dafny.SeqWithoutIsStrInference([d_1390_v0_])
        d_1392_v2_: int
        d_1392_v2_ = -765
        d_1393_v3_: _dafny.MultiSet
        d_1393_v3_ = (d_1391_v1_)[default__.safeIndex(d_1392_v2_, len(d_1391_v1_))]
        source22_ = d_1393_v3_
        d_1394___mcc_h0_ = source22_
        d_1395_cf0_ = d_1394___mcc_h0_
        d_1396_v4_: str
        d_1396_v4_ = _dafny.CodePoint('u')
        d_1397_v5_: D3
        d_1397_v5_ = D3_DC5(d_1396_v4_)
        source23_ = d_1397_v5_
        if source23_.is_DC6:
            d_1398___mcc_h1_ = source23_.cf12
            d_1399___mcc_h2_ = source23_.cf13
            d_1400_cf13_ = d_1399___mcc_h2_
            d_1401_cf12_ = d_1398___mcc_h1_
            d_1402_v7_: _dafny.Seq
            d_1402_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cdytr"))
            d_1403_v8_: _dafny.Map
            d_1403_v8_ = _dafny.Map({d_1402_v7_: d_1396_v4_})
            def iife123_():
                coll45_ = _dafny.Map()
                compr_45_: _dafny.Seq
                for compr_45_ in ((d_1403_v8_) | (_dafny.Map({d_1402_v7_: d_1396_v4_}))).keys.Elements:
                    d_1404_v6_: _dafny.Seq = compr_45_
                    if (d_1404_v6_) in ((d_1403_v8_) | (_dafny.Map({d_1402_v7_: d_1396_v4_}))):
                        coll45_[d_1404_v6_] = (len(_dafny.SeqWithoutIsStrInference([self.f24, (self).f13]))) * (d_1392_v2_)
                return _dafny.Map(coll45_)
            d_1392_v2_ = (0) - (len(iife123_()
            ))
            d_1405_v9_: _dafny.Array
            def lambda58_(d_1406_i0_):
                return (self).f13

            init33_ = lambda58_
            nw199_ = _dafny.Array(None, 7)
            for i0_33_ in range(nw199_.length(0)):
                nw199_[i0_33_] = init33_(i0_33_)
            d_1405_v9_ = nw199_
            d_1405_v9_ = (d_1405_v9_ if self.f24 else (self).f14)
            d_1407_v10_: _dafny.Seq
            d_1407_v10_ = _dafny.SeqWithoutIsStrInference([d_1400_cf13_])
            d_1408_v11_: _dafny.Set
            d_1408_v11_ = _dafny.Set({d_1407_v10_})
            index233_ = default__.safeIndex(319, ((self).f14).length(0))
            ((self).f14)[index233_] = (self).fm52((0) - ((self).fm53(d_1401_cf12_, globalState)), d_1408_v11_, globalState)
            d_1409_v12_: _dafny.Map
            d_1409_v12_ = _dafny.Map({(self).f13: (self).f13})
            d_1410_v13_: _dafny.MultiSet
            d_1410_v13_ = _dafny.MultiSet([d_1402_v7_])
            d_1411_v14_: _dafny.Map
            d_1411_v14_ = _dafny.Map({d_1401_cf12_: d_1400_cf13_})
            d_1412_v16_: _dafny.Set
            def iife124_():
                coll46_ = _dafny.Set()
                compr_46_: int
                for compr_46_ in _dafny.IntegerRange(996, -82):
                    d_1413_v15_: int = compr_46_
                    if ((996) <= (d_1413_v15_)) and ((d_1413_v15_) < (-82)):
                        coll46_ = coll46_.union(_dafny.Set([(d_1413_v15_) * (d_1392_v2_)]))
                return _dafny.Set(coll46_)
            d_1412_v16_ = _dafny.Set({iife124_()
            })
            index234_ = default__.safeIndex(319, ((self).f14).length(0))
            rhs200_ = not(((d_1409_v12_)[(self).fm52(d_1392_v2_, d_1408_v11_, globalState)] if ((self).fm52(d_1392_v2_, d_1408_v11_, globalState)) in (d_1409_v12_) else (self).f13))
            rhs201_ = (d_1396_v4_) in ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kv"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nolgwiuyb"))))
            rhs202_ = ((d_1410_v13_)[(d_1402_v7_) + (default__.fm60(self.f24, (self).f13, d_1411_v14_, len(d_1412_v16_), globalState))] if ((d_1402_v7_) + (default__.fm60(self.f24, (self).f13, d_1411_v14_, len(d_1412_v16_), globalState))) in (d_1410_v13_) else d_1392_v2_)
            lhs167_ = (self).f14
            lhs168_ = default__.safeIndex(319, ((self).f14).length(0))
            lhs169_ = globalState
            d_1400_cf13_ = rhs200_
            lhs167_[lhs168_] = rhs201_
            lhs169_.f7 = rhs202_
            d_1414_v17_: C4
            nw200_ = C4()
            nw200_.ctor__()
            d_1414_v17_ = nw200_
            d_1415_v18_: _dafny.Map
            d_1415_v18_ = _dafny.Map({d_1414_v17_: d_1392_v2_})
            d_1416_v19_: _dafny.Seq
            d_1416_v19_ = _dafny.SeqWithoutIsStrInference([len(d_1415_v18_)])
            d_1417_v20_: _dafny.Map
            d_1417_v20_ = _dafny.Map({self.f24: (d_1416_v19_)[default__.safeIndex(988, len(d_1416_v19_))]})
            d_1417_v20_ = d_1417_v20_
        elif source23_.is_DC7:
            d_1418_v21_: _dafny.Map
            d_1418_v21_ = _dafny.Map({937: (0) - (-550)})
            d_1419_v22_: _dafny.Set
            d_1419_v22_ = _dafny.Set({default__.fm61(self.f24, len(d_1418_v21_), _dafny.MultiSet([d_1396_v4_, d_1396_v4_]), globalState)})
            d_1420_v23_: _dafny.Map
            d_1420_v23_ = _dafny.Map({(self).fm52((0) - (d_1392_v2_), d_1419_v22_, globalState): (self).f14})
            d_1421_v24_: C2
            nw201_ = C2()
            nw201_.ctor__(d_1392_v2_, (self.f24 if True else self.f24), D1_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uy"))), (self).f13, ((d_1420_v23_)[(self).f13] if ((self).f13) in (d_1420_v23_) else (self).f14), (self).f15, self.f12, self.f24)
            d_1421_v24_ = nw201_
            d_1422_v25_: _dafny.Array
            nw202_ = _dafny.Array(_dafny.Seq({}), 5)
            d_1422_v25_ = nw202_
            d_1423_v26_: D8
            d_1423_v26_ = D8_DC18(d_1422_v25_)
            d_1423_v26_ = d_1423_v26_
            index235_ = default__.safeIndex(101, ((self).f14).length(0))
            ((self).f14)[index235_] = (self).f13
            index236_ = default__.safeIndex(101, ((self).f14).length(0))
            rhs203_ = not(self.f24)
            rhs204_ = self.f24
            rhs205_ = False
            lhs170_ = (self).f14
            lhs171_ = default__.safeIndex(101, ((self).f14).length(0))
            lhs172_ = self
            lhs173_ = self
            lhs170_[lhs171_] = rhs203_
            lhs172_.f24 = rhs204_
            lhs173_.f24 = rhs205_
            d_1424_v27_: _dafny.MultiSet
            d_1424_v27_ = _dafny.MultiSet([d_1392_v2_, d_1392_v2_])
            (globalState).f7 = (((d_1424_v27_)[(d_1421_v24_).f21] if ((d_1421_v24_).f21) in (d_1424_v27_) else d_1392_v2_)) * (d_1392_v2_)
        elif source23_.is_DC8:
            d_1425___mcc_h3_ = source23_.cf14
            d_1426_cf14_ = d_1425___mcc_h3_
            d_1427_v28_: _dafny.Seq
            d_1427_v28_ = _dafny.SeqWithoutIsStrInference([self.f24, self.f24])
            d_1428_v29_: D15
            d_1428_v29_ = D15_DC38(d_1427_v28_)
            d_1429_v30_: _dafny.Set
            d_1429_v30_ = _dafny.Set({(d_1428_v29_).cf77, _dafny.SeqWithoutIsStrInference([self.f24])})
            d_1430_v31_: _dafny.Seq
            d_1430_v31_ = _dafny.SeqWithoutIsStrInference([d_1429_v30_, _dafny.Set({d_1427_v28_, d_1427_v28_, d_1427_v28_})])
            d_1431_v32_: _dafny.Set
            d_1431_v32_ = _dafny.Set({self.f24})
            d_1432_v33_: _dafny.Seq
            d_1432_v33_ = _dafny.SeqWithoutIsStrInference([d_1431_v32_])
            (self).f24 = not((self).fm52(d_1392_v2_, (d_1430_v31_)[default__.safeIndex(len((d_1432_v33_)[default__.safeIndex(d_1392_v2_, len(d_1432_v33_))]), len(d_1430_v31_))], globalState))
            d_1433_v34_: _dafny.Seq
            d_1433_v34_ = _dafny.SeqWithoutIsStrInference([d_1392_v2_, len((_dafny.SeqWithoutIsStrInference([True, (self).f13])) + (d_1427_v28_)), d_1426_cf14_])
            (globalState).f7 = (d_1433_v34_)[default__.safeIndex((0) - (d_1426_cf14_), len(d_1433_v34_))]
            d_1392_v2_ = d_1426_cf14_
            rhs206_ = (((self).f15).ispropersubset((self).f15)) or ((self).f13)
            rhs207_ = default__.safeModuloInt(d_1392_v2_, d_1426_cf14_)
            rhs208_ = d_1427_v28_
            lhs174_ = self
            lhs175_ = globalState
            lhs176_ = globalState
            lhs174_.f24 = rhs206_
            lhs175_.f7 = rhs207_
            lhs176_.f2 = rhs208_
        elif True:
            d_1434___mcc_h4_ = source23_.cf11
            d_1435_cf11_ = d_1434___mcc_h4_
            d_1436_v35_: _dafny.Seq
            d_1436_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dffjrx"))
            d_1437_v36_: _dafny.Seq
            d_1437_v36_ = _dafny.SeqWithoutIsStrInference([d_1392_v2_])
            d_1438_v37_: _dafny.Map
            d_1438_v37_ = _dafny.Map({(_dafny.MultiSet(d_1437_v36_)).cardinality: d_1396_v4_})
            d_1439_v38_: _dafny.Map
            d_1439_v38_ = _dafny.Map({len(d_1438_v37_): (self).f13})
            d_1440_v39_: _dafny.Map
            d_1440_v39_ = _dafny.Map({self.f24: d_1392_v2_})
            d_1441_v40_: D6
            d_1441_v40_ = D6_DC13(self.f24, self.f24, d_1392_v2_, len(_dafny.SeqWithoutIsStrInference([(self).f13])), not((self).f13))
            d_1442_v41_: T0
            nw203_ = C3()
            nw203_.ctor__((self).f13, default__.fm62(self.f24, d_1436_v35_, globalState), (self).f13, (self).f14, (self).f15, self.f12, self.f24)
            d_1442_v41_ = nw203_
            d_1443_v42_: _dafny.Map
            d_1443_v42_ = _dafny.Map({d_1442_v41_: (self).f13})
            d_1444_v43_: _dafny.MultiSet
            d_1444_v43_ = _dafny.MultiSet([d_1392_v2_])
            d_1445_v44_: _dafny.Set
            d_1445_v44_ = _dafny.Set({(_dafny.SeqWithoutIsStrInference([(d_1442_v41_).f13, True, (self).f13, (self).f13])).set(default__.safeIndex((d_1444_v43_).cardinality, len(_dafny.SeqWithoutIsStrInference([(d_1442_v41_).f13, True, (self).f13, (self).f13]))), (self).f13)})
            d_1446_v45_: _dafny.Map
            d_1446_v45_ = _dafny.Map({(self).f13: (self).fm52(d_1392_v2_, d_1445_v44_, globalState)})
            d_1447_v46_: _dafny.Array
            nw204_ = _dafny.Array(None, 28)
            nw204_[int(0)] = (len(d_1436_v35_)) + (d_1392_v2_)
            nw204_[int(1)] = len((d_1439_v38_) | (d_1439_v38_))
            nw204_[int(2)] = d_1392_v2_
            nw204_[int(3)] = ((d_1440_v39_)[self.f24] if (self.f24) in (d_1440_v39_) else d_1392_v2_)
            nw204_[int(4)] = d_1392_v2_
            nw204_[int(5)] = d_1392_v2_
            nw204_[int(6)] = (428) + ((self).fm8((self).f13, D1_DC1(d_1436_v35_), globalState))
            nw204_[int(7)] = d_1392_v2_
            nw204_[int(8)] = d_1392_v2_
            nw204_[int(9)] = d_1392_v2_
            nw204_[int(10)] = ((0) - ((d_1441_v40_).cf25)) + (d_1392_v2_)
            nw204_[int(11)] = d_1392_v2_
            nw204_[int(12)] = (0) - ((d_1392_v2_) + (d_1392_v2_))
            nw204_[int(13)] = d_1392_v2_
            nw204_[int(14)] = d_1392_v2_
            nw204_[int(15)] = d_1392_v2_
            nw204_[int(16)] = d_1392_v2_
            nw204_[int(17)] = len(_dafny.Map({self.f24: d_1392_v2_}))
            nw204_[int(18)] = d_1392_v2_
            nw204_[int(19)] = d_1392_v2_
            nw204_[int(20)] = (d_1392_v2_ if (self).f13 else d_1392_v2_)
            nw204_[int(21)] = len(d_1443_v42_)
            nw204_[int(22)] = len((d_1446_v45_) | (d_1446_v45_))
            nw204_[int(23)] = d_1392_v2_
            nw204_[int(24)] = d_1392_v2_
            nw204_[int(25)] = (105) - (d_1392_v2_)
            nw204_[int(26)] = (0) - (d_1392_v2_)
            nw204_[int(27)] = d_1392_v2_
            d_1447_v46_ = nw204_
            index237_ = default__.safeIndex(642, (d_1447_v46_).length(0))
            (d_1447_v46_)[index237_] = d_1392_v2_
            index238_ = default__.safeIndex(642, (d_1447_v46_).length(0))
            (d_1447_v46_)[index238_] = ((d_1395_cf0_).cardinality) + (d_1392_v2_)
            d_1444_v43_ = d_1444_v43_
            index239_ = default__.safeIndex(642, (d_1447_v46_).length(0))
            (d_1447_v46_)[index239_] = (0) - ((d_1447_v46_)[default__.safeIndex(642, (d_1447_v46_).length(0))])
            d_1448_v47_: _dafny.Map
            d_1448_v47_ = _dafny.Map({(d_1447_v46_)[default__.safeIndex(642, (d_1447_v46_).length(0))]: (0) - (d_1392_v2_)})
            d_1449_v48_: _dafny.Map
            d_1449_v48_ = _dafny.Map({d_1435_cf11_: (d_1437_v36_)[default__.safeIndex(len(d_1448_v47_), len(d_1437_v36_))]})
            d_1449_v48_ = (d_1449_v48_).set(d_1396_v4_, d_1392_v2_)
        index240_ = default__.safeIndex(234, ((self).f14).length(0))
        ((self).f14)[index240_] = self.f24
        d_1450_v49_: _dafny.Map
        d_1450_v49_ = _dafny.Map({(self).f13: d_1392_v2_})
        d_1451_v50_: D6
        d_1451_v50_ = D6_DC13((self).f13, (self).f13, d_1392_v2_, len(d_1450_v49_), not((self).f13))
        d_1452_v51_: _dafny.Map
        d_1452_v51_ = _dafny.Map({False: self.f24})
        d_1453_v52_: _dafny.Seq
        d_1453_v52_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hmg"))
        pat_let_tv36_ = d_1453_v52_
        pat_let_tv37_ = d_1452_v51_
        pat_let_tv38_ = d_1452_v51_
        d_1454_v53_: _dafny.Map
        def iife125_(_pat_let39_0):
            def iife126_(d_1455_dt__update__tmp_h0_):
                def iife127_(_pat_let40_0):
                    def iife128_(d_1456_dt__update_hcf24_h0_):
                        def iife129_(_pat_let41_0):
                            def iife130_(d_1457_dt__update_hcf22_h0_):
                                return D6_DC13(d_1457_dt__update_hcf22_h0_, (d_1455_dt__update__tmp_h0_).cf23, d_1456_dt__update_hcf24_h0_, (d_1455_dt__update__tmp_h0_).cf25, (d_1455_dt__update__tmp_h0_).cf26)
                            return iife130_(_pat_let41_0)
                        return iife129_(((pat_let_tv37_)[self.f24] if (self.f24) in (pat_let_tv38_) else not(not((self).f13))))
                    return iife128_(_pat_let40_0)
                return iife127_(len(pat_let_tv36_))
            return iife126_(_pat_let39_0)
        d_1454_v53_ = _dafny.Map({d_1392_v2_: (iife125_(d_1451_v50_)).cf26})
        index241_ = default__.safeIndex(234, ((self).f14).length(0))
        ((self).f14)[index241_] = ((d_1454_v53_)[(d_1392_v2_) - (d_1392_v2_)] if ((d_1392_v2_) - (d_1392_v2_)) in (d_1454_v53_) else not(False))
        (self).f24 = (not((self).f13) if (self).f13 else self.f24)
        d_1458_v54_: _dafny.Seq
        d_1458_v54_ = _dafny.SeqWithoutIsStrInference([((self).f14)[default__.safeIndex(234, ((self).f14).length(0))]])
        d_1459_v55_: _dafny.MultiSet
        d_1459_v55_ = _dafny.MultiSet([len((d_1453_v52_).set(default__.safeIndex(d_1392_v2_, len(d_1453_v52_)), d_1396_v4_)), d_1392_v2_])
        if (d_1459_v55_).ispropersubset((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_1392_v2_ for d_1460_i1_ in range(default__.abs(424))]))).intersection(default__.fm56(len(d_1458_v54_), (self).f15, globalState))):
            d_1461_v56_: _dafny.Seq
            d_1461_v56_ = _dafny.SeqWithoutIsStrInference([d_1392_v2_, 703, d_1392_v2_, d_1392_v2_, d_1392_v2_])
            d_1461_v56_ = d_1461_v56_
            d_1462_v57_: _dafny.Array
            def lambda59_(d_1463_v2_):
                def lambda60_(d_1464_i2_):
                    return (d_1464_i2_) + (d_1463_v2_)

                return lambda60_

            init34_ = lambda59_(d_1392_v2_)
            nw205_ = _dafny.Array(None, 4)
            for i0_34_ in range(nw205_.length(0)):
                nw205_[i0_34_] = init34_(i0_34_)
            d_1462_v57_ = nw205_
            index242_ = default__.safeIndex(548, (d_1462_v57_).length(0))
            (d_1462_v57_)[index242_] = (d_1392_v2_) - (len(d_1450_v49_))
            index243_ = default__.safeIndex(548, (d_1462_v57_).length(0))
            (d_1462_v57_)[index243_] = (d_1392_v2_) * (d_1392_v2_)
            (globalState).f7 = len((self).f15)
            d_1453_v52_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_1465_i3_ in range(default__.abs(483))])
            index244_ = default__.safeIndex(548, (d_1462_v57_).length(0))
            (d_1462_v57_)[index244_] = (d_1462_v57_)[default__.safeIndex(548, (d_1462_v57_).length(0))]
        elif True:
            index245_ = default__.safeIndex(234, ((self).f14).length(0))
            ((self).f14)[index245_] = ((self).f14)[default__.safeIndex(234, ((self).f14).length(0))]
            d_1466_v58_: _dafny.Seq
            d_1466_v58_ = _dafny.SeqWithoutIsStrInference([d_1392_v2_])
            index246_ = default__.safeIndex(234, ((self).f14).length(0))
            ((self).f14)[index246_] = not((len(d_1466_v58_)) <= (d_1392_v2_))
            r0 = d_1392_v2_
            (globalState).f7 = d_1392_v2_
            d_1467_v59_: D7
            d_1467_v59_ = D7_DC16(d_1392_v2_)
            d_1467_v59_ = d_1467_v59_
        d_1468_v60_: _dafny.Seq
        d_1468_v60_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jnxuox"))
        d_1469_v61_: _dafny.Array
        nw206_ = _dafny.Array(int(0), 21)
        d_1469_v61_ = nw206_
        d_1470_v62_: _dafny.Map
        d_1470_v62_ = _dafny.Map({default__.safeDivisionInt(d_1392_v2_, len(d_1468_v60_)): d_1469_v61_})
        d_1471_v63_: _dafny.Map
        d_1471_v63_ = _dafny.Map({False: d_1392_v2_})
        d_1472_v64_: D1
        d_1472_v64_ = D1_DC1(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_1473_i4_ in range(default__.abs(734))]))
        d_1474_v65_: C2
        nw207_ = C2()
        nw207_.ctor__(len(d_1471_v63_), self.f24, d_1472_v64_, not(True), (self).f14, (self).f15, self.f12, self.f24)
        d_1474_v65_ = nw207_
        d_1475_v66_: _dafny.Map
        d_1475_v66_ = _dafny.Map({d_1474_v65_: _dafny.SeqWithoutIsStrInference([(self).f13])})
        d_1476_v67_: _dafny.Seq
        d_1476_v67_ = _dafny.SeqWithoutIsStrInference([self.f24])
        d_1477_v68_: _dafny.Set
        d_1477_v68_ = _dafny.Set({((d_1475_v66_)[d_1474_v65_] if (d_1474_v65_) in (d_1475_v66_) else d_1476_v67_)})
        d_1478_v69_: D6
        d_1478_v69_ = D6_DC14(D6_DC13((self).f13, (self).fm52(d_1392_v2_, d_1477_v68_, globalState), (_dafny.MultiSet(d_1476_v67_)).cardinality, d_1392_v2_, (d_1474_v65_).f22))
        d_1479_v70_: _dafny.Map
        d_1479_v70_ = _dafny.Map({d_1392_v2_: d_1478_v69_})
        d_1470_v62_ = (d_1470_v62_).set(len(d_1479_v70_), d_1469_v61_)
        hi9_ = d_1392_v2_
        for d_1480_i5_ in range(d_1392_v2_, hi9_):
            index247_ = default__.safeIndex(473, (d_1469_v61_).length(0))
            (d_1469_v61_)[index247_] = d_1480_i5_
            d_1481_v71_: str
            d_1481_v71_ = _dafny.CodePoint('t')
            index248_ = default__.safeIndex(473, (d_1469_v61_).length(0))
            rhs209_ = not ((d_1468_v60_) < ((d_1468_v60_).set(default__.safeIndex((d_1474_v65_).f21, len(d_1468_v60_)), d_1481_v71_))) or ((self).fm52(d_1392_v2_, d_1477_v68_, globalState))
            rhs210_ = d_1469_v61_
            rhs211_ = d_1468_v60_
            rhs212_ = (d_1474_v65_).f22
            rhs213_ = (d_1474_v65_).f21
            lhs177_ = self
            lhs178_ = self
            lhs179_ = d_1469_v61_
            lhs180_ = default__.safeIndex(473, (d_1469_v61_).length(0))
            lhs177_.f24 = rhs209_
            d_1469_v61_ = rhs210_
            d_1468_v60_ = rhs211_
            lhs178_.f24 = rhs212_
            lhs179_[lhs180_] = rhs213_
            d_1482_v72_: _dafny.Array
            nw208_ = _dafny.Array(int(0), 19)
            d_1482_v72_ = nw208_
            d_1483_v73_: D11
            d_1483_v73_ = D11_DC26(d_1482_v72_, d_1468_v60_)
            d_1484_v74_: _dafny.Array
            nw209_ = _dafny.Array(None, 8)
            nw209_[int(0)] = d_1482_v72_
            nw209_[int(1)] = d_1482_v72_
            nw209_[int(2)] = d_1482_v72_
            nw209_[int(3)] = d_1482_v72_
            nw209_[int(4)] = d_1482_v72_
            nw209_[int(5)] = (d_1483_v73_).cf49
            nw209_[int(6)] = d_1482_v72_
            nw209_[int(7)] = d_1482_v72_
            d_1484_v74_ = nw209_
            index249_ = default__.safeIndex(938, (d_1484_v74_).length(0))
            (d_1484_v74_)[index249_] = d_1482_v72_
            d_1485_v75_: _dafny.Array
            nw210_ = _dafny.Array(_dafny.Array(None, 0), 9)
            d_1485_v75_ = nw210_
            d_1486_v76_: _dafny.Map
            d_1486_v76_ = _dafny.Map({(d_1474_v65_).f21: 501})
            index250_ = default__.safeIndex(938, (d_1484_v74_).length(0))
            rhs214_ = self.f24
            rhs215_ = d_1482_v72_
            rhs216_ = (d_1485_v75_ if (len(_dafny.Map({(self).f14: True}))) > ((0) - (len(d_1486_v76_))) else d_1485_v75_)
            lhs181_ = self
            lhs182_ = d_1484_v74_
            lhs183_ = default__.safeIndex(938, (d_1484_v74_).length(0))
            lhs181_.f24 = rhs214_
            lhs182_[lhs183_] = rhs215_
            d_1485_v75_ = rhs216_
            d_1487_v77_: D2
            d_1487_v77_ = D2_DC4((d_1484_v74_)[default__.safeIndex(938, (d_1484_v74_).length(0))], (self).f13, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iqftrwd"))), True, (self).f14)
            d_1488_v78_: _dafny.Map
            d_1488_v78_ = _dafny.Map({(d_1481_v71_ if not(self.f24) else d_1481_v71_): (d_1487_v77_).cf7})
            d_1488_v78_ = (d_1488_v78_).set((d_1481_v71_ if (self).f13 else d_1481_v71_), (d_1474_v65_).f22)
            (globalState).f7 = d_1480_i5_
        d_1489_i6_: int
        d_1489_i6_ = 0
        with _dafny.label("7"):
            while not((d_1474_v65_).f22):
                with _dafny.c_label("7"):
                    if (d_1489_i6_) >= (100):
                        raise _dafny.Break("7")
                    d_1489_i6_ = (d_1489_i6_) + (1)
                    if True:
                        index251_ = default__.safeIndex(866, ((self).f14).length(0))
                        ((self).f14)[index251_] = (d_1474_v65_).f22
                        d_1490_v79_: _dafny.Set
                        d_1490_v79_ = _dafny.Set({d_1469_v61_, d_1469_v61_, d_1469_v61_})
                        index252_ = default__.safeIndex(866, ((self).f14).length(0))
                        ((self).f14)[index252_] = (_dafny.Set({d_1469_v61_})) == (d_1490_v79_)
                        d_1491_v80_: str
                        d_1491_v80_ = _dafny.CodePoint('x')
                        d_1492_v81_: T0
                        nw211_ = C0()
                        nw211_.ctor__(self.f24, d_1491_v80_, self.f12, (d_1474_v65_).f22)
                        d_1492_v81_ = nw211_
                        d_1493_v82_: D16
                        d_1493_v82_ = D16_DC40(d_1492_v81_)
                        d_1492_v81_ = (d_1493_v82_).cf83
                        (globalState).f7 = (d_1474_v65_).f21
                        d_1494_v83_: _dafny.Array
                        nw212_ = _dafny.Array(_dafny.Array(None, 0), 16)
                        d_1494_v83_ = nw212_
                        d_1495_v84_: _dafny.Array
                        nw213_ = _dafny.Array(False, 20)
                        d_1495_v84_ = nw213_
                        index253_ = default__.safeIndex(222, (d_1494_v83_).length(0))
                        (d_1494_v83_)[index253_] = d_1495_v84_
                        index254_ = default__.safeIndex(222, (d_1494_v83_).length(0))
                        (d_1494_v83_)[index254_] = d_1495_v84_
                        d_1479_v70_ = (d_1479_v70_).set((d_1474_v65_).f21, d_1478_v69_)
                    elif True:
                        d_1496_v85_: _dafny.MultiSet
                        d_1496_v85_ = _dafny.MultiSet([(self).f14, (self).f14])
                        (globalState).f7 = default__.safeModuloInt(822, ((d_1496_v85_)[(self).f14] if ((self).f14) in (d_1496_v85_) else (d_1474_v65_).f21))
                        d_1497_v86_: _dafny.Map
                        d_1497_v86_ = _dafny.Map({d_1468_v60_: (self).f13})
                        rhs217_ = (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_1498_i7_ in range(default__.abs(601))]))) - (len(d_1497_v86_))
                        rhs218_ = (((d_1474_v65_).f21) + (d_1392_v2_)) - ((0) - ((d_1474_v65_).f21))
                        rhs219_ = (self).f13
                        rhs220_ = (0) - (default__.safeDivisionInt((d_1474_v65_).f21, d_1392_v2_))
                        rhs221_ = False
                        lhs184_ = globalState
                        lhs185_ = self
                        lhs186_ = self
                        r0 = rhs217_
                        lhs184_.f7 = rhs218_
                        lhs185_.f24 = rhs219_
                        r0 = rhs220_
                        lhs186_.f24 = rhs221_
                        def iife131_():
                            coll47_ = _dafny.Map()
                            compr_47_: int
                            for compr_47_ in _dafny.IntegerRange(155, -903):
                                d_1499_v87_: int = compr_47_
                                if ((155) <= (d_1499_v87_)) and ((d_1499_v87_) < (-903)):
                                    coll47_[default__.safeDivisionInt(d_1499_v87_, d_1392_v2_)] = d_1478_v69_
                            return _dafny.Map(coll47_)
                        (globalState).f7 = len((d_1479_v70_ if not((d_1474_v65_).f22) else (d_1479_v70_ if self.f24 else iife131_()
                        )))
                        index255_ = default__.safeIndex(941, ((self).f14).length(0))
                        ((self).f14)[index255_] = self.f24
                        d_1500_v88_: _dafny.Seq
                        d_1500_v88_ = _dafny.SeqWithoutIsStrInference([d_1468_v60_, d_1468_v60_, d_1468_v60_])
                        index256_ = default__.safeIndex(941, ((self).f14).length(0))
                        rhs222_ = (d_1500_v88_)[default__.safeIndex(default__.safeModuloInt(d_1392_v2_, (d_1474_v65_).f21), len(d_1500_v88_))]
                        rhs223_ = ((d_1474_v65_).f21) >= ((0) - (default__.safeDivisionInt((d_1474_v65_).f21, (d_1474_v65_).f21)))
                        lhs187_ = (self).f14
                        lhs188_ = default__.safeIndex(941, ((self).f14).length(0))
                        d_1468_v60_ = rhs222_
                        lhs187_[lhs188_] = rhs223_
                        (self).f24 = ((d_1474_v65_).f21) == ((d_1474_v65_).f21)
                    (self).f24 = (self).fm52((d_1474_v65_).f21, d_1477_v68_, globalState)
                    d_1501_v89_: _dafny.Seq
                    d_1501_v89_ = _dafny.SeqWithoutIsStrInference([d_1392_v2_, d_1392_v2_, (d_1474_v65_).f21, (d_1474_v65_).f21, (d_1474_v65_).f21])
                    d_1502_v90_: _dafny.MultiSet
                    d_1502_v90_ = _dafny.MultiSet([d_1501_v89_])
                    if ((_dafny.MultiSet([d_1501_v89_]) if self.f24 else d_1502_v90_)).ispropersubset(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(d_1474_v65_).f21 for d_1503_i8_ in range(default__.abs(99))]), default__.fm63((d_1474_v65_).f22, (d_1474_v65_).f22, globalState)])):
                        d_1504_v91_: _dafny.Map
                        d_1504_v91_ = _dafny.Map({(d_1474_v65_).f22: (self).f13})
                        d_1504_v91_ = (d_1504_v91_).set((d_1474_v65_).f22, (d_1474_v65_).f22)
                        d_1505_v92_: str
                        d_1505_v92_ = _dafny.CodePoint('t')
                        d_1506_v93_: C1
                        nw214_ = C1()
                        nw214_.ctor__(d_1505_v92_, D1_DC2((d_1474_v65_).f21, (d_1474_v65_).f21, (d_1474_v65_).f22), (self).f13)
                        d_1506_v93_ = nw214_
                        d_1507_v94_: _dafny.Map
                        d_1507_v94_ = _dafny.Map({d_1506_v93_: (d_1474_v65_).f22})
                        d_1508_v95_: _dafny.Map
                        d_1508_v95_ = _dafny.Map({(d_1507_v94_) | (d_1507_v94_): (d_1474_v65_).f21})
                        rhs224_ = (d_1506_v93_).fm43(globalState)
                        rhs225_ = (d_1508_v95_) | (d_1508_v95_)
                        lhs189_ = self
                        lhs189_.f24 = rhs224_
                        d_1508_v95_ = rhs225_
                        (globalState).f7 = (d_1474_v65_).f21
                        (self).f24 = (self.f24) or ((self).f13)
                        d_1509_v96_: _dafny.Map
                        d_1509_v96_ = _dafny.Map({502: (d_1474_v65_).f21})
                        d_1510_v97_: _dafny.Map
                        d_1510_v97_ = _dafny.Map({d_1468_v60_: d_1509_v96_})
                        d_1510_v97_ = (d_1510_v97_).set(d_1468_v60_, d_1509_v96_)
                    elif True:
                        d_1511_v98_: _dafny.MultiSet
                        d_1511_v98_ = _dafny.MultiSet([d_1476_v67_, d_1476_v67_])
                        index257_ = default__.safeIndex(476, ((self).f14).length(0))
                        ((self).f14)[index257_] = (d_1511_v98_) == (_dafny.MultiSet([d_1476_v67_]))
                        index258_ = default__.safeIndex(476, ((self).f14).length(0))
                        ((self).f14)[index258_] = (self.f24) and ((self).fm52((_dafny.MultiSet([(d_1474_v65_).f22])).cardinality, d_1477_v68_, globalState))
                        r0 = d_1392_v2_
                        (globalState).f7 = (d_1474_v65_).f21
                        d_1512_v99_: _dafny.MultiSet
                        d_1512_v99_ = _dafny.MultiSet([default__.safeModuloInt((d_1474_v65_).f21, (d_1474_v65_).f21)])
                        d_1392_v2_ = (d_1512_v99_).cardinality
                        d_1468_v60_ = (d_1468_v60_ if (d_1390_v0_).ispropersubset(_dafny.MultiSet([(self).f13])) else d_1468_v60_)
                    d_1513_v100_: _dafny.Map
                    d_1513_v100_ = _dafny.Map({False: self.f24})
                    d_1513_v100_ = _dafny.Map({(self).f13: (d_1390_v0_).isdisjoint(_dafny.MultiSet([(self).f13]))})
                    pass
            pass
        (globalState).f7 = (len(default__.fm63(True, self.f24, globalState))) + (d_1392_v2_)
        d_1514_v101_: D11
        d_1514_v101_ = D11_DC27(False, _dafny.CodePoint('q'), 890, d_1392_v2_, _dafny.SeqWithoutIsStrInference([118]))
        d_1515_i9_: int
        d_1515_i9_ = 0
        with _dafny.label("8"):
            while ((d_1390_v0_) != (d_1390_v0_)) and ((d_1514_v101_).cf51):
                with _dafny.c_label("8"):
                    if (d_1515_i9_) >= (100):
                        raise _dafny.Break("8")
                    d_1515_i9_ = (d_1515_i9_) + (1)
                    d_1516_v102_: _dafny.Array
                    nw215_ = _dafny.Array(False, 14)
                    d_1516_v102_ = nw215_
                    def lambda61_(d_1517_v101_):
                        def lambda62_(d_1518_i10_):
                            return (d_1517_v101_).cf51

                        return lambda62_

                    init35_ = lambda61_(d_1514_v101_)
                    nw216_ = _dafny.Array(None, 1)
                    for i0_35_ in range(nw216_.length(0)):
                        nw216_[i0_35_] = init35_(i0_35_)
                    d_1516_v102_ = nw216_
                    d_1468_v60_ = (d_1468_v60_ if (d_1474_v65_).f22 else d_1468_v60_)
                    d_1519_v103_: str
                    d_1519_v103_ = _dafny.CodePoint('i')
                    d_1520_v104_: C0
                    nw217_ = C0()
                    nw217_.ctor__((d_1474_v65_).f22, d_1519_v103_, self.f12, (d_1476_v67_)[default__.safeIndex(-991, len(d_1476_v67_))])
                    d_1520_v104_ = nw217_
                    d_1521_v105_: _dafny.Map
                    d_1521_v105_ = _dafny.Map({(d_1474_v65_).f21: (_dafny.Set({(self).fm53(d_1392_v2_, globalState)})) | ((self).f15)})
                    d_1521_v105_ = (d_1521_v105_).set((d_1474_v65_).f21, (default__.fm64(globalState)) - (_dafny.Set({d_1392_v2_})))
                    pass
            pass
        d_1522_v106_: _dafny.Map
        d_1522_v106_ = _dafny.Map({134: (d_1474_v65_).f21})
        r0 = len((d_1522_v106_) | (d_1522_v106_))
        return r0


class C6(T3, T2, T1, T0):
    def  __init__(self):
        self._f12: D1 = D1.default()()
        self._f16: D1 = D1.default()()
        self._f14: _dafny.Array = _dafny.Array(None, 0)
        self._f15: _dafny.Set = _dafny.Set({})
        self._f13: bool = False
        self._f17: bool = False
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
    def f16(self):
        return self._f16
    @f16.setter
    def f16(self, value):
        self._f16 = value
    @property
    def f14(self):
        return self._f14
    @property
    def f15(self):
        return self._f15
    @property
    def f13(self):
        return self._f13
    @property
    def f17(self):
        return self._f17
    def ctor__(self, f16, f17, f14, f15, f12, f13):
        (self).f16 = f16
        (self)._f17 = f17
        (self)._f14 = f14
        (self)._f15 = f15
        (self).f12 = f12
        (self)._f13 = f13

    def fm10(self, p0, p1, p2, globalState):
        return _dafny.Map({(_dafny.MultiSet([(self).f13, (self).f17])).issubset(_dafny.MultiSet([(self).f17])): len(_dafny.SeqWithoutIsStrInference([705, (0) - (-805)]))})

    def fm8(self, p0, p1, globalState):
        if (self).f17:
            return -683
        elif True:
            return (0) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({(self).f13: _dafny.SeqWithoutIsStrInference([(self).f17, (self).f13])})))]))).cardinality)

    def fm9(self, p0, globalState):
        def iife132_():
            coll48_ = _dafny.Map()
            compr_48_: int
            for compr_48_ in _dafny.IntegerRange(-200, 231):
                d_1524_v0_: int = compr_48_
                if ((-200) <= (d_1524_v0_)) and ((d_1524_v0_) < (231)):
                    coll48_[default__.safeDivisionInt(d_1524_v0_, 374)] = 805
            return _dafny.Map(coll48_)
        return (0) - (default__.safeDivisionInt(((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_1523_i0_ in range(default__.abs(711))])))) * (len(iife132_()
        )), len((_dafny.Set({(self).f13, not((self).f17)})).intersection(_dafny.Set({(self).f13})))))

    def fm11(self, globalState):
        return (self).f13

    def fm12(self, p0, p1, globalState):
        return (self).f17

    def m6(self, p0, globalState):
        r0: int = int(0)
        d_1525_v0_: int
        d_1525_v0_ = 622
        hi10_ = (d_1525_v0_) - (d_1525_v0_)
        for d_1526_i0_ in range((d_1525_v0_) - (d_1525_v0_), hi10_):
            d_1527_v1_: _dafny.Array
            nw218_ = _dafny.Array(_dafny.Set({}), 1)
            d_1527_v1_ = nw218_
            d_1528_v2_: _dafny.Set
            d_1528_v2_ = _dafny.Set({p0, (self).f17})
            index259_ = default__.safeIndex(183, (d_1527_v1_).length(0))
            (d_1527_v1_)[index259_] = (d_1528_v2_) | (d_1528_v2_)
            index260_ = default__.safeIndex(183, (d_1527_v1_).length(0))
            (d_1527_v1_)[index260_] = d_1528_v2_
            d_1529_v3_: bool
            d_1529_v3_ = False
            d_1530_v4_: _dafny.Array
            nw219_ = _dafny.Array(int(0), 2)
            d_1530_v4_ = nw219_
            d_1531_v5_: _dafny.MultiSet
            d_1531_v5_ = _dafny.MultiSet([(self).f13])
            index261_ = default__.safeIndex(344, (d_1530_v4_).length(0))
            (d_1530_v4_)[index261_] = (d_1531_v5_).cardinality
            d_1532_v6_: _dafny.Seq
            d_1532_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pjvi"))
            d_1533_v7_: _dafny.Map
            d_1533_v7_ = _dafny.Map({len(d_1532_v6_): (self).f13})
            d_1534_v8_: _dafny.Seq
            d_1534_v8_ = _dafny.SeqWithoutIsStrInference([len((d_1527_v1_)[default__.safeIndex(183, (d_1527_v1_).length(0))])])
            d_1535_v9_: _dafny.MultiSet
            d_1535_v9_ = _dafny.MultiSet([d_1526_i0_])
            d_1536_v10_: _dafny.Seq
            d_1536_v10_ = _dafny.SeqWithoutIsStrInference([not(((d_1533_v7_)[d_1525_v0_] if (d_1525_v0_) in (d_1533_v7_) else p0)), (d_1535_v9_).ispropersubset(_dafny.MultiSet(d_1534_v8_))])
            d_1537_v11_: _dafny.Map
            d_1537_v11_ = _dafny.Map({default__.fm13(d_1532_v6_, globalState): (self).f13})
            index262_ = default__.safeIndex(344, (d_1530_v4_).length(0))
            rhs226_ = (d_1536_v10_)[default__.safeIndex(d_1526_i0_, len(d_1536_v10_))]
            rhs227_ = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sn")))
            rhs228_ = (0) - ((0) - (d_1525_v0_))
            rhs229_ = d_1525_v0_
            rhs230_ = not (p0) or ((self).fm12(d_1537_v11_, self.f12, globalState))
            lhs190_ = globalState
            lhs191_ = d_1530_v4_
            lhs192_ = default__.safeIndex(344, (d_1530_v4_).length(0))
            d_1529_v3_ = rhs226_
            r0 = rhs227_
            lhs190_.f7 = rhs228_
            lhs191_[lhs192_] = rhs229_
            d_1529_v3_ = rhs230_
            d_1538_v12_: str
            d_1538_v12_ = _dafny.CodePoint('v')
            d_1539_v13_: _dafny.Map
            d_1539_v13_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_1540_i2_ in range(default__.abs(486))])) for d_1541_i1_ in range(default__.abs(181))])): d_1538_v12_})
            d_1542_v14_: _dafny.Map
            d_1542_v14_ = _dafny.Map({p0: (d_1530_v4_)[default__.safeIndex(344, (d_1530_v4_).length(0))]})
            d_1543_v15_: _dafny.Set
            d_1543_v15_ = _dafny.Set({d_1530_v4_})
            d_1544_v16_: D2
            d_1544_v16_ = D2_DC4(d_1530_v4_, p0, d_1526_i0_, d_1529_v3_, (self).f14)
            d_1545_v17_: _dafny.Array
            nw220_ = _dafny.Array(None, 17)
            nw220_[int(0)] = d_1529_v3_
            nw220_[int(1)] = d_1529_v3_
            nw220_[int(2)] = (d_1532_v6_) <= (d_1532_v6_)
            nw220_[int(3)] = p0
            nw220_[int(4)] = (d_1529_v3_) or (d_1529_v3_)
            nw220_[int(5)] = d_1529_v3_
            nw220_[int(6)] = d_1529_v3_
            nw220_[int(7)] = p0
            nw220_[int(8)] = (False) and (d_1529_v3_)
            nw220_[int(9)] = (self).f17
            nw220_[int(10)] = (self).f13
            nw220_[int(11)] = (_dafny.MultiSet([len(d_1539_v13_), (0) - ((d_1534_v8_)[default__.safeIndex((d_1530_v4_)[default__.safeIndex(344, (d_1530_v4_).length(0))], len(d_1534_v8_))])])).isdisjoint(_dafny.MultiSet([(0) - (d_1526_i0_)]))
            nw220_[int(12)] = d_1529_v3_
            nw220_[int(13)] = (((d_1542_v14_)[(self).f17] if ((self).f17) in (d_1542_v14_) else default__.fm0((d_1536_v10_).set(default__.safeIndex((d_1530_v4_)[default__.safeIndex(344, (d_1530_v4_).length(0))], len(d_1536_v10_)), (self).f17), (d_1536_v10_)[default__.safeIndex((d_1530_v4_)[default__.safeIndex(344, (d_1530_v4_).length(0))], len(d_1536_v10_))], (self).f13, len(d_1543_v15_), globalState))) >= (d_1525_v0_)
            nw220_[int(14)] = (self).f13
            nw220_[int(15)] = not((d_1536_v10_)[default__.safeIndex(980, len(d_1536_v10_))])
            nw220_[int(16)] = (d_1544_v16_).cf9
            d_1545_v17_ = nw220_
            d_1545_v17_ = (self).f14
            index263_ = default__.safeIndex(344, (d_1530_v4_).length(0))
            (d_1530_v4_)[index263_] = (d_1530_v4_)[default__.safeIndex(344, (d_1530_v4_).length(0))]
        d_1546_v18_: _dafny.Map
        d_1546_v18_ = _dafny.Map({True: p0})
        d_1547_v19_: _dafny.Seq
        d_1547_v19_ = _dafny.SeqWithoutIsStrInference([((d_1546_v18_)[p0] if (p0) in (d_1546_v18_) else not(False)), (self).f13])
        if (d_1547_v19_) <= ((d_1547_v19_) + (_dafny.SeqWithoutIsStrInference([(self).f17]))):
            d_1548_v20_: str
            d_1548_v20_ = _dafny.CodePoint('b')
            d_1549_v21_: _dafny.Seq
            d_1549_v21_ = _dafny.SeqWithoutIsStrInference([d_1525_v0_, d_1525_v0_, d_1525_v0_, d_1525_v0_])
            d_1550_v22_: _dafny.Seq
            d_1550_v22_ = _dafny.SeqWithoutIsStrInference([(d_1549_v21_)[default__.safeIndex(d_1525_v0_, len(d_1549_v21_))], d_1525_v0_])
            d_1551_v23_: C1
            nw221_ = C1()
            nw221_.ctor__(d_1548_v20_, self.f12, (d_1550_v22_) != (d_1549_v21_))
            d_1551_v23_ = nw221_
            d_1525_v0_ = (d_1525_v0_) + (d_1525_v0_)
            index264_ = default__.safeIndex(608, ((self).f14).length(0))
            ((self).f14)[index264_] = (self).f17
            d_1552_v24_: _dafny.Seq
            d_1552_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mttwbblrp"))
            d_1553_v25_: _dafny.Map
            d_1553_v25_ = _dafny.Map({True: d_1525_v0_})
            index265_ = default__.safeIndex(608, ((self).f14).length(0))
            rhs231_ = default__.fm0((d_1547_v19_) + (d_1547_v19_), not((self).f17), (self).f13, d_1525_v0_, globalState)
            rhs232_ = (self).f13
            rhs233_ = default__.fm18((0) - ((0) - (d_1525_v0_)), d_1552_v24_, (self).f17, (d_1553_v25_ if (self).f13 else (self).fm10((self).f17, (self).f13, d_1547_v19_, globalState)), globalState)
            rhs234_ = d_1525_v0_
            lhs193_ = globalState
            lhs194_ = (self).f14
            lhs195_ = default__.safeIndex(608, ((self).f14).length(0))
            lhs196_ = d_1551_v23_
            lhs193_.f7 = rhs231_
            lhs194_[lhs195_] = rhs232_
            lhs196_.f23 = rhs233_
            r0 = rhs234_
            d_1554_v26_: _dafny.Array
            nw222_ = _dafny.Array(None, 4)
            nw222_[int(0)] = d_1525_v0_
            nw222_[int(1)] = d_1525_v0_
            nw222_[int(2)] = d_1525_v0_
            nw222_[int(3)] = (0) - (d_1525_v0_)
            d_1554_v26_ = nw222_
            d_1555_v27_: _dafny.Array
            nw223_ = _dafny.Array(False, 17)
            d_1555_v27_ = nw223_
            d_1556_v28_: D2
            d_1556_v28_ = D2_DC4(d_1554_v26_, (d_1547_v19_)[default__.safeIndex(d_1525_v0_, len(d_1547_v19_))], d_1525_v0_, (self).f13, d_1555_v27_)
            pat_let_tv39_ = d_1555_v27_
            pat_let_tv40_ = d_1525_v0_
            pat_let_tv41_ = d_1554_v26_
            def iife133_(_pat_let42_0):
                def iife134_(d_1557_dt__update__tmp_h0_):
                    def iife135_(_pat_let43_0):
                        def iife136_(d_1558_dt__update_hcf10_h0_):
                            def iife137_(_pat_let44_0):
                                def iife138_(d_1559_dt__update_hcf7_h0_):
                                    def iife139_(_pat_let45_0):
                                        def iife140_(d_1560_dt__update_hcf8_h0_):
                                            def iife141_(_pat_let46_0):
                                                def iife142_(d_1561_dt__update_hcf6_h0_):
                                                    return D2_DC4(d_1561_dt__update_hcf6_h0_, d_1559_dt__update_hcf7_h0_, d_1560_dt__update_hcf8_h0_, (d_1557_dt__update__tmp_h0_).cf9, d_1558_dt__update_hcf10_h0_)
                                                return iife142_(_pat_let46_0)
                                            return iife141_(pat_let_tv41_)
                                        return iife140_(_pat_let45_0)
                                    return iife139_(pat_let_tv40_)
                                return iife138_(_pat_let44_0)
                            return iife137_(((self).f14)[default__.safeIndex(608, ((self).f14).length(0))])
                        return iife136_(_pat_let43_0)
                    return iife135_(pat_let_tv39_)
                return iife134_(_pat_let42_0)
            d_1553_v25_ = (d_1553_v25_).set((iife133_(d_1556_v28_)).cf9, default__.safeDivisionInt(d_1525_v0_, d_1525_v0_))
            (globalState).f7 = (d_1525_v0_) + (d_1525_v0_)
        elif True:
            d_1562_v29_: C4
            nw224_ = C4()
            nw224_.ctor__()
            d_1562_v29_ = nw224_
            nw225_ = C4()
            nw225_.ctor__()
            d_1562_v29_ = nw225_
            d_1563_v30_: _dafny.Seq
            d_1563_v30_ = _dafny.SeqWithoutIsStrInference([(d_1525_v0_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tfumdn")))), d_1525_v0_])
            d_1563_v30_ = (d_1563_v30_) + (d_1563_v30_)
            (globalState).f7 = d_1525_v0_
            r0 = (d_1525_v0_) - ((self.f12).cf3)
            d_1564_v31_: D2
            d_1564_v31_ = D2_DC3((self).f14)
            d_1565_v32_: _dafny.Array
            def lambda63_(d_1566_i3_):
                return (d_1566_i3_) + (118)

            init36_ = lambda63_
            nw226_ = _dafny.Array(None, 10)
            for i0_36_ in range(nw226_.length(0)):
                nw226_[i0_36_] = init36_(i0_36_)
            d_1565_v32_ = nw226_
            d_1567_v33_: D2
            d_1567_v33_ = D2_DC4(d_1565_v32_, (self).f17, d_1525_v0_, False, (self).f14)
            d_1568_v34_: _dafny.Map
            d_1568_v34_ = _dafny.Map({not((self).f17): (self).f14})
            d_1569_v35_: _dafny.Array
            nw227_ = _dafny.Array(None, 17)
            nw227_[int(0)] = (self).f14
            nw227_[int(1)] = (d_1564_v31_).cf5
            nw227_[int(2)] = (self).f14
            nw227_[int(3)] = (d_1567_v33_).cf10
            nw227_[int(4)] = (self).f14
            nw227_[int(5)] = (self).f14
            nw227_[int(6)] = (d_1567_v33_).cf10
            nw227_[int(7)] = (self).f14
            nw227_[int(8)] = (self).f14
            nw227_[int(9)] = (self).f14
            nw227_[int(10)] = (self).f14
            nw227_[int(11)] = ((d_1568_v34_)[True] if (True) in (d_1568_v34_) else (self).f14)
            nw227_[int(12)] = (self).f14
            nw227_[int(13)] = (self).f14
            nw227_[int(14)] = (self).f14
            nw227_[int(15)] = (self).f14
            nw227_[int(16)] = (self).f14
            d_1569_v35_ = nw227_
            index266_ = default__.safeIndex(959, (d_1569_v35_).length(0))
            (d_1569_v35_)[index266_] = (self).f14
            index267_ = default__.safeIndex(959, (d_1569_v35_).length(0))
            (d_1569_v35_)[index267_] = (self).f14
        d_1570_v36_: _dafny.Array
        nw228_ = _dafny.Array(_dafny.Map({}), 26)
        d_1570_v36_ = nw228_
        nw229_ = _dafny.Array(_dafny.Map({}), 1)
        d_1570_v36_ = nw229_
        d_1571_v37_: _dafny.Array
        nw230_ = _dafny.Array(int(0), 11)
        d_1571_v37_ = nw230_
        guard_loop_6_: int
        for guard_loop_6_ in _dafny.IntegerRange(0, (d_1571_v37_).length(0)):
            d_1572_i4_: int = guard_loop_6_
            if (True) and (((0) <= (d_1572_i4_)) and ((d_1572_i4_) < ((d_1571_v37_).length(0)))):
                (d_1571_v37_)[(d_1572_i4_)] = default__.safeModuloInt(d_1572_i4_, default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([d_1525_v0_])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dscsy")))))
        hi11_ = d_1525_v0_
        for d_1573_i5_ in range(d_1525_v0_, hi11_):
            d_1574_v38_: bool
            d_1574_v38_ = False
            d_1574_v38_ = not(p0)
            d_1575_v39_: _dafny.MultiSet
            d_1575_v39_ = _dafny.MultiSet([d_1525_v0_])
            d_1576_v40_: _dafny.Map
            d_1576_v40_ = _dafny.Map({p0: d_1575_v39_})
            (globalState).f2 = (d_1547_v19_).set(default__.safeIndex((((d_1576_v40_)[not(p0)] if (not(p0)) in (d_1576_v40_) else d_1575_v39_)).cardinality, len(d_1547_v19_)), (d_1574_v38_ if True else (self).f13))
            d_1577_v41_: _dafny.Seq
            d_1577_v41_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fbqky"))
            d_1578_v42_: _dafny.Map
            d_1578_v42_ = _dafny.Map({-435: p0})
            rhs235_ = default__.safeModuloInt(d_1525_v0_, d_1573_i5_)
            rhs236_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_1579_i6_ in range(default__.abs(-811))])) != (d_1577_v41_)
            rhs237_ = (d_1546_v18_ if ((d_1578_v42_)[d_1573_i5_] if (d_1573_i5_) in (d_1578_v42_) else d_1574_v38_) else (d_1546_v18_) | (_dafny.Map({(self).f13: (self).f17})))
            d_1525_v0_ = rhs235_
            d_1574_v38_ = rhs236_
            d_1546_v18_ = rhs237_
            d_1580_v43_: D2
            d_1580_v43_ = D2_DC3((self).f14)
            def iife143_(_pat_let47_0):
                def iife144_(d_1581_dt__update__tmp_h1_):
                    def iife145_(_pat_let48_0):
                        def iife146_(d_1582_dt__update_hcf5_h0_):
                            return D2_DC3(d_1582_dt__update_hcf5_h0_)
                        return iife146_(_pat_let48_0)
                    return iife145_((self).f14)
                return iife144_(_pat_let47_0)
            d_1580_v43_ = iife143_(d_1580_v43_)
        d_1583_v44_: _dafny.Seq
        d_1583_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "axb"))
        if (len((d_1583_v44_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gsllko"))))) > ((d_1525_v0_) * (len(_dafny.SeqWithoutIsStrInference([d_1525_v0_ for d_1584_i7_ in range(default__.abs(435))])))):
            d_1585_v45_: str
            d_1585_v45_ = _dafny.CodePoint('h')
            d_1585_v45_ = d_1585_v45_
            d_1586_v46_: _dafny.MultiSet
            d_1586_v46_ = _dafny.MultiSet([d_1571_v37_])
            d_1587_v47_: _dafny.Map
            d_1587_v47_ = _dafny.Map({d_1585_v45_: (self).f13})
            d_1588_v48_: C0
            nw231_ = C0()
            nw231_.ctor__((d_1586_v46_).issubset(d_1586_v46_), (d_1585_v45_ if (self).f13 else d_1585_v45_), self.f12, (self).fm12(d_1587_v47_, self.f12, globalState))
            d_1588_v48_ = nw231_
            d_1589_v49_: _dafny.Array
            def lambda64_(d_1590_v0_):
                def lambda65_(d_1591_i8_):
                    return D7_DC16(d_1590_v0_)

                return lambda65_

            init37_ = lambda64_(d_1525_v0_)
            nw232_ = _dafny.Array(None, 17)
            for i0_37_ in range(nw232_.length(0)):
                nw232_[i0_37_] = init37_(i0_37_)
            d_1589_v49_ = nw232_
            d_1592_v50_: _dafny.Map
            d_1592_v50_ = _dafny.Map({default__.fm46(d_1525_v0_, ((d_1546_v18_)[(self).f17] if ((self).f17) in (d_1546_v18_) else (self.f12).cf4), (self).f17, globalState): d_1589_v49_})
            index268_ = default__.safeIndex(942, (d_1571_v37_).length(0))
            (d_1571_v37_)[index268_] = len(d_1592_v50_)
            index269_ = default__.safeIndex(942, (d_1571_v37_).length(0))
            (d_1571_v37_)[index269_] = len(default__.fm47(p0, globalState))
            d_1593_v51_: _dafny.Map
            d_1593_v51_ = _dafny.Map({(d_1571_v37_)[default__.safeIndex(942, (d_1571_v37_).length(0))]: d_1583_v44_})
            d_1594_v52_: _dafny.Array
            nw233_ = _dafny.Array(None, 23)
            nw233_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wnydrhnob"))
            nw233_[int(1)] = d_1583_v44_
            nw233_[int(2)] = _dafny.SeqWithoutIsStrInference([(d_1588_v48_).f20 for d_1595_i9_ in range(default__.abs(-269))])
            nw233_[int(3)] = d_1583_v44_
            nw233_[int(4)] = d_1583_v44_
            nw233_[int(5)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fh"))) + (d_1583_v44_)
            nw233_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "orofsd"))
            nw233_[int(7)] = ((d_1583_v44_) + (d_1583_v44_)).set(default__.safeIndex((d_1571_v37_)[default__.safeIndex(942, (d_1571_v37_).length(0))], len((d_1583_v44_) + (d_1583_v44_))), _dafny.CodePoint('f'))
            nw233_[int(8)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ftv"))) + (d_1583_v44_)
            nw233_[int(9)] = (d_1583_v44_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_1596_i10_ in range(default__.abs(891))]))
            nw233_[int(10)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lgqso"))
            nw233_[int(11)] = d_1583_v44_
            nw233_[int(12)] = (d_1583_v44_) + ((d_1583_v44_).set(default__.safeIndex(d_1525_v0_, len(d_1583_v44_)), _dafny.CodePoint('f')))
            nw233_[int(13)] = d_1583_v44_
            nw233_[int(14)] = ((d_1593_v51_)[(0) - ((d_1571_v37_)[default__.safeIndex(942, (d_1571_v37_).length(0))])] if ((0) - ((d_1571_v37_)[default__.safeIndex(942, (d_1571_v37_).length(0))])) in (d_1593_v51_) else d_1583_v44_)
            nw233_[int(15)] = d_1583_v44_
            nw233_[int(16)] = _dafny.SeqWithoutIsStrInference([(d_1588_v48_).f20 for d_1597_i11_ in range(default__.abs(-534))])
            nw233_[int(17)] = d_1583_v44_
            nw233_[int(18)] = d_1583_v44_
            nw233_[int(19)] = d_1583_v44_
            nw233_[int(20)] = d_1583_v44_
            nw233_[int(21)] = (d_1583_v44_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xbclijk")))
            nw233_[int(22)] = d_1583_v44_
            d_1594_v52_ = nw233_
            d_1594_v52_ = d_1594_v52_
            d_1598_v53_: _dafny.Map
            d_1598_v53_ = _dafny.Map({(self).f14: p0})
            d_1599_v54_: _dafny.Array
            nw234_ = _dafny.Array(int(0), 22)
            d_1599_v54_ = nw234_
            d_1600_v55_: D2
            d_1600_v55_ = D2_DC4(d_1599_v54_, d_1588_v48_.f19, d_1525_v0_, d_1588_v48_.f19, (self).f14)
            pat_let_tv42_ = d_1571_v37_
            pat_let_tv43_ = d_1571_v37_
            def iife147_(_pat_let49_0):
                def iife148_(d_1601_dt__update__tmp_h2_):
                    def iife149_(_pat_let50_0):
                        def iife150_(d_1602_dt__update_hcf8_h1_):
                            def iife151_(_pat_let51_0):
                                def iife152_(d_1603_dt__update_hcf9_h0_):
                                    return D2_DC4((d_1601_dt__update__tmp_h2_).cf6, (d_1601_dt__update__tmp_h2_).cf7, d_1602_dt__update_hcf8_h1_, d_1603_dt__update_hcf9_h0_, (d_1601_dt__update__tmp_h2_).cf10)
                                return iife152_(_pat_let51_0)
                            return iife151_((self).f17)
                        return iife150_(_pat_let50_0)
                    return iife149_((pat_let_tv43_)[default__.safeIndex(942, (pat_let_tv42_).length(0))])
                return iife148_(_pat_let49_0)
            d_1598_v53_ = (d_1598_v53_).set((self).f14, (iife147_(d_1600_v55_)).cf9)
        elif True:
            d_1604_v56_: bool
            d_1604_v56_ = False
            d_1604_v56_ = default__.fm17(globalState)
            r0 = d_1525_v0_
            d_1605_v57_: str
            d_1605_v57_ = _dafny.CodePoint('q')
            rhs238_ = 486
            rhs239_ = d_1605_v57_
            r0 = rhs238_
            d_1605_v57_ = rhs239_
            d_1606_v58_: _dafny.Seq
            d_1606_v58_ = _dafny.SeqWithoutIsStrInference([d_1525_v0_])
            def iife153_():
                coll49_ = _dafny.Set()
                compr_49_: int
                for compr_49_ in (_dafny.Map({(d_1606_v58_)[default__.safeIndex(d_1525_v0_, len(d_1606_v58_))]: p0})).keys.Elements:
                    d_1607_v59_: int = compr_49_
                    if (d_1607_v59_) in (_dafny.Map({(d_1606_v58_)[default__.safeIndex(d_1525_v0_, len(d_1606_v58_))]: p0})):
                        coll49_ = coll49_.union(_dafny.Set([(d_1607_v59_) + (-552)]))
                return _dafny.Set(coll49_)
            if (iife153_()
            ).isdisjoint((self).f15):
                d_1608_v60_: _dafny.Map
                d_1608_v60_ = _dafny.Map({309: d_1525_v0_})
                d_1609_v61_: _dafny.MultiSet
                d_1609_v61_ = _dafny.MultiSet([((d_1608_v60_)[(0) - (d_1525_v0_)] if ((0) - (d_1525_v0_)) in (d_1608_v60_) else d_1525_v0_), (0) - (len((default__.fm27(globalState)) + (_dafny.SeqWithoutIsStrInference([126, d_1525_v0_]))))])
                d_1609_v61_ = (d_1609_v61_) | (d_1609_v61_)
                index270_ = default__.safeIndex(981, ((self).f14).length(0))
                ((self).f14)[index270_] = (self).f13
                index271_ = default__.safeIndex(981, ((self).f14).length(0))
                ((self).f14)[index271_] = (self).f13
                index272_ = default__.safeIndex(981, ((self).f14).length(0))
                ((self).f14)[index272_] = not(not(((self).f14)[default__.safeIndex(981, ((self).f14).length(0))]))
                index273_ = default__.safeIndex(116, (d_1571_v37_).length(0))
                (d_1571_v37_)[index273_] = d_1525_v0_
                index274_ = default__.safeIndex(116, (d_1571_v37_).length(0))
                (d_1571_v37_)[index274_] = d_1525_v0_
                (globalState).f7 = len(((d_1583_v44_) + ((d_1583_v44_).set(default__.safeIndex(d_1525_v0_, len(d_1583_v44_)), d_1605_v57_))).set(default__.safeIndex(d_1525_v0_, len((d_1583_v44_) + ((d_1583_v44_).set(default__.safeIndex(d_1525_v0_, len(d_1583_v44_)), d_1605_v57_)))), d_1605_v57_))
            elif True:
                index275_ = default__.safeIndex(411, ((self).f14).length(0))
                ((self).f14)[index275_] = not (p0) or (p0)
                index276_ = default__.safeIndex(411, ((self).f14).length(0))
                ((self).f14)[index276_] = (self).f13
                (globalState).f7 = (self).fm8((self).f17, self.f16, globalState)
                d_1610_v62_: _dafny.Seq
                d_1610_v62_ = _dafny.SeqWithoutIsStrInference([d_1583_v44_])
                d_1611_v63_: _dafny.Map
                d_1611_v63_ = _dafny.Map({default__.fm38((self).f13, ((self).f14)[default__.safeIndex(411, ((self).f14).length(0))], False, ((self).f14)[default__.safeIndex(411, ((self).f14).length(0))], globalState): (d_1583_v44_) not in ((d_1610_v62_).set(default__.safeIndex(d_1525_v0_, len(d_1610_v62_)), d_1583_v44_))})
                d_1612_v64_: _dafny.Map
                d_1612_v64_ = _dafny.Map({len(d_1547_v19_): not(((self).f14)[default__.safeIndex(411, ((self).f14).length(0))])})
                d_1613_v65_: _dafny.Map
                d_1613_v65_ = _dafny.Map({len((d_1612_v64_).set(len(d_1547_v19_), (self).f13)): (self).f13})
                d_1611_v63_ = (d_1611_v63_).set((d_1547_v19_) + (d_1547_v19_), (d_1604_v56_ if (self).f13 else ((d_1613_v65_)[d_1525_v0_] if (d_1525_v0_) in (d_1613_v65_) else ((self).f14)[default__.safeIndex(411, ((self).f14).length(0))])))
                index277_ = default__.safeIndex(584, (d_1571_v37_).length(0))
                (d_1571_v37_)[index277_] = ((0) - (len(default__.fm20(globalState)))) + (d_1525_v0_)
                index278_ = default__.safeIndex(584, (d_1571_v37_).length(0))
                rhs240_ = default__.safeDivisionInt((self).fm9(_dafny.CodePoint('g'), globalState), d_1525_v0_)
                rhs241_ = d_1605_v57_
                rhs242_ = (d_1525_v0_) - (-984)
                lhs197_ = d_1571_v37_
                lhs198_ = default__.safeIndex(584, (d_1571_v37_).length(0))
                lhs197_[lhs198_] = rhs240_
                d_1605_v57_ = rhs241_
                d_1525_v0_ = rhs242_
                d_1614_v66_: _dafny.Set
                d_1614_v66_ = _dafny.Set({(0) - ((d_1571_v37_)[default__.safeIndex(584, (d_1571_v37_).length(0))])})
                d_1614_v66_ = (self).f15
            if d_1604_v56_:
                d_1615_v67_: D9
                d_1615_v67_ = D9_DC20((self).f15)
                d_1616_v68_: _dafny.MultiSet
                d_1616_v68_ = _dafny.MultiSet([d_1615_v67_, d_1615_v67_])
                index279_ = default__.safeIndex(863, (d_1571_v37_).length(0))
                (d_1571_v37_)[index279_] = d_1525_v0_
                index280_ = default__.safeIndex(863, (d_1571_v37_).length(0))
                rhs243_ = ((d_1616_v68_).intersection(d_1616_v68_)) - (d_1616_v68_)
                rhs244_ = d_1525_v0_
                rhs245_ = True
                rhs246_ = d_1547_v19_
                lhs199_ = d_1571_v37_
                lhs200_ = default__.safeIndex(863, (d_1571_v37_).length(0))
                lhs201_ = globalState
                d_1616_v68_ = rhs243_
                lhs199_[lhs200_] = rhs244_
                d_1604_v56_ = rhs245_
                lhs201_.f2 = rhs246_
                index281_ = default__.safeIndex(181, ((self).f14).length(0))
                ((self).f14)[index281_] = True
                d_1617_v69_: C3
                nw235_ = C3()
                nw235_.ctor__((self).f13, self.f16, d_1604_v56_, (self).f14, (self).f15, D1_DC2(d_1525_v0_, (d_1571_v37_)[default__.safeIndex(863, (d_1571_v37_).length(0))], (self).f17), d_1604_v56_)
                d_1617_v69_ = nw235_
                index282_ = default__.safeIndex(181, ((self).f14).length(0))
                ((self).f14)[index282_] = default__.fm30(((d_1571_v37_)[default__.safeIndex(863, (d_1571_v37_).length(0))]) - (len(_dafny.Map({(self).f17: d_1617_v69_}))), globalState)
                d_1618_v70_: _dafny.Map
                d_1618_v70_ = _dafny.Map({(d_1571_v37_)[default__.safeIndex(863, (d_1571_v37_).length(0))]: (d_1571_v37_)[default__.safeIndex(863, (d_1571_v37_).length(0))]})
                d_1619_v71_: _dafny.Map
                d_1619_v71_ = _dafny.Map({d_1618_v70_: default__.fm0(d_1547_v19_, not(True), p0, (d_1571_v37_)[default__.safeIndex(863, (d_1571_v37_).length(0))], globalState)})
                d_1620_v72_: _dafny.MultiSet
                d_1620_v72_ = _dafny.MultiSet([p0, d_1604_v56_, p0])
                index283_ = default__.safeIndex(181, ((self).f14).length(0))
                index284_ = default__.safeIndex(181, ((self).f14).length(0))
                rhs247_ = default__.fm38(not (False) or ((self).f13), ((self).f17 if ((self).f14)[default__.safeIndex(181, ((self).f14).length(0))] else d_1604_v56_), d_1604_v56_, d_1604_v56_, globalState)
                rhs248_ = (((self).f14)[default__.safeIndex(181, ((self).f14).length(0))]) and ((d_1617_v69_).f18)
                rhs249_ = (D11_DC25(d_1619_v71_)).cf48
                rhs250_ = (d_1546_v18_) | ((d_1546_v18_).set((self).f13, (self).f17))
                rhs251_ = ((d_1620_v72_).intersection(_dafny.MultiSet(d_1547_v19_))) == (d_1620_v72_)
                lhs202_ = globalState
                lhs203_ = (self).f14
                lhs204_ = default__.safeIndex(181, ((self).f14).length(0))
                lhs205_ = (self).f14
                lhs206_ = default__.safeIndex(181, ((self).f14).length(0))
                lhs202_.f2 = rhs247_
                lhs203_[lhs204_] = rhs248_
                d_1619_v71_ = rhs249_
                d_1546_v18_ = rhs250_
                lhs205_[lhs206_] = rhs251_
                d_1621_v73_: _dafny.Map
                d_1621_v73_ = _dafny.Map({len(((self).f15) | ((self).f15)): ((self).f13) and ((d_1617_v69_).f18)})
                d_1621_v73_ = (d_1621_v73_) | (d_1621_v73_)
                d_1621_v73_ = (d_1621_v73_).set(default__.safeModuloInt(d_1525_v0_, 671), True)
            elif True:
                d_1622_v74_: D11
                d_1622_v74_ = D11_DC26(d_1571_v37_, d_1583_v44_)
                d_1623_v75_: D11
                d_1623_v75_ = D11_DC28(d_1622_v74_)
                d_1624_v76_: _dafny.Array
                nw236_ = _dafny.Array(None, 9)
                nw236_[int(0)] = d_1623_v75_
                nw236_[int(1)] = d_1623_v75_
                nw236_[int(2)] = d_1623_v75_
                nw236_[int(3)] = d_1623_v75_
                nw236_[int(4)] = d_1623_v75_
                nw236_[int(5)] = D11_DC28(d_1622_v74_)
                nw236_[int(6)] = d_1623_v75_
                nw236_[int(7)] = D11_DC28(d_1622_v74_)
                nw236_[int(8)] = D11_DC28(d_1622_v74_)
                d_1624_v76_ = nw236_
                d_1625_v77_: _dafny.Seq
                d_1625_v77_ = _dafny.SeqWithoutIsStrInference([d_1624_v76_])
                d_1625_v77_ = d_1625_v77_
                d_1604_v56_ = ((self).f15).isdisjoint(default__.fm47((self).f13, globalState))
                d_1571_v37_ = d_1571_v37_
                d_1626_v78_: _dafny.Map
                d_1626_v78_ = _dafny.Map({p0: (d_1525_v0_) * (d_1525_v0_)})
                d_1626_v78_ = (d_1626_v78_).set((self).f17, d_1525_v0_)
                d_1627_v79_: _dafny.Map
                d_1627_v79_ = _dafny.Map({d_1525_v0_: d_1605_v57_})
                d_1628_v80_: D3
                d_1628_v80_ = D3_DC5(d_1605_v57_)
                d_1627_v79_ = (d_1627_v79_).set(d_1525_v0_, (d_1628_v80_).cf11)
        r0 = 411
        return r0

    def m7(self, globalState):
        d_1629_i0_: int
        d_1629_i0_ = 0
        with _dafny.label("9"):
            while (self).f13:
                with _dafny.c_label("9"):
                    if (d_1629_i0_) >= (100):
                        raise _dafny.Break("9")
                    d_1629_i0_ = (d_1629_i0_) + (1)
                    d_1630_v0_: str
                    d_1630_v0_ = _dafny.CodePoint('c')
                    d_1631_v1_: C0
                    nw237_ = C0()
                    nw237_.ctor__((d_1630_v0_) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "erlwcj"))), d_1630_v0_, self.f12, (self).f13)
                    d_1631_v1_ = nw237_
                    d_1632_v2_: int
                    d_1632_v2_ = 697
                    d_1633_v3_: C1
                    nw238_ = C1()
                    nw238_.ctor__(d_1630_v0_, D1_DC2(d_1632_v2_, d_1632_v2_, (self).f13), (self).f17)
                    d_1633_v3_ = nw238_
                    d_1634_v4_: _dafny.Map
                    d_1634_v4_ = _dafny.Map({not(d_1631_v1_.f19): d_1633_v3_})
                    d_1635_v5_: _dafny.Map
                    d_1635_v5_ = _dafny.Map({d_1630_v0_: d_1631_v1_.f19})
                    d_1636_v6_: _dafny.Seq
                    d_1636_v6_ = _dafny.SeqWithoutIsStrInference([(self).f17])
                    d_1634_v4_ = (d_1634_v4_).set((self).fm12(d_1635_v5_, D1_DC2(len(d_1636_v6_), d_1632_v2_, (self).f17), globalState), d_1633_v3_)
                    d_1637_v7_: D3
                    d_1637_v7_ = D3_DC5(default__.fm33(d_1631_v1_.f19, 4, globalState))
                    d_1637_v7_ = d_1637_v7_
                    d_1638_v8_: C3
                    nw239_ = C3()
                    nw239_.ctor__((self).f17, D1_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yitlqi"))), (self).fm12(d_1635_v5_, D1_DC2(d_1632_v2_, 219, (self).f17), globalState), (self).f14, (self).f15, D1_DC2(d_1632_v2_, 478, False), d_1631_v1_.f19)
                    d_1638_v8_ = nw239_
                    pass
            pass
        d_1639_v9_: int
        d_1639_v9_ = 993
        (globalState).f7 = d_1639_v9_
        d_1639_v9_ = d_1639_v9_
        d_1640_v10_: _dafny.Seq
        d_1640_v10_ = _dafny.SeqWithoutIsStrInference([True])
        d_1641_v11_: _dafny.Seq
        d_1641_v11_ = _dafny.SeqWithoutIsStrInference([(d_1640_v10_) + (d_1640_v10_)])
        (globalState).f2 = (d_1641_v11_)[default__.safeIndex(d_1639_v9_, len(d_1641_v11_))]
        d_1642_v12_: _dafny.Seq
        d_1642_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))
        d_1643_v13_: _dafny.Set
        d_1643_v13_ = _dafny.Set({d_1642_v12_, d_1642_v12_})
        d_1644_v14_: _dafny.Set
        d_1644_v14_ = (default__.fm48(globalState) if (self).f13 else d_1643_v13_)
        d_1644_v14_ = d_1644_v14_
        d_1645_v15_: _dafny.Array
        nw240_ = _dafny.Array(_dafny.Map({}), 7)
        d_1645_v15_ = nw240_
        guard_loop_7_: int
        for guard_loop_7_ in _dafny.IntegerRange(0, (d_1645_v15_).length(0)):
            d_1646_i1_: int = guard_loop_7_
            if (True) and (((0) <= (d_1646_i1_)) and ((d_1646_i1_) < ((d_1645_v15_).length(0)))):
                (d_1645_v15_)[(d_1646_i1_)] = ((_dafny.Map({(self).f13: d_1639_v9_})) | (_dafny.Map({True: -165}))) | ((_dafny.Map({(self).f13: d_1639_v9_})) | (_dafny.Map({(self).f17: d_1639_v9_})))

    def m5(self, p0, globalState):
        r0: bool = False
        d_1647_i0_: int
        d_1647_i0_ = 0
        with _dafny.label("10"):
            while (self).f17:
                with _dafny.c_label("10"):
                    if (d_1647_i0_) >= (100):
                        raise _dafny.Break("10")
                    d_1647_i0_ = (d_1647_i0_) + (1)
                    d_1648_v0_: _dafny.Seq
                    d_1648_v0_ = _dafny.SeqWithoutIsStrInference([(self).f17, (self).f13])
                    d_1649_v1_: str
                    d_1649_v1_ = _dafny.CodePoint('o')
                    d_1650_v2_: D3
                    d_1650_v2_ = D3_DC5(d_1649_v1_)
                    d_1651_v3_: _dafny.MultiSet
                    d_1651_v3_ = _dafny.MultiSet([(self).f17, (d_1648_v0_)[default__.safeIndex(p0, len(d_1648_v0_))], (self).fm12(_dafny.Map({(d_1650_v2_).cf11: (self).f17}), D1_DC2(p0, p0, (self).f13), globalState)])
                    d_1652_v4_: _dafny.Map
                    d_1652_v4_ = _dafny.Map({d_1649_v1_: p0})
                    d_1653_v5_: _dafny.Array
                    def lambda66_(d_1654_i1_):
                        return default__.safeModuloInt(d_1654_i1_, len(_dafny.Set({(self).f17, not((self).f13), (self).f17, (self).f13})))

                    init38_ = lambda66_
                    nw241_ = _dafny.Array(None, 20)
                    for i0_38_ in range(nw241_.length(0)):
                        nw241_[i0_38_] = init38_(i0_38_)
                    d_1653_v5_ = nw241_
                    d_1655_v6_: _dafny.Map
                    d_1655_v6_ = _dafny.Map({d_1652_v4_: d_1653_v5_})
                    d_1656_v7_: _dafny.Seq
                    d_1656_v7_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                    d_1657_v8_: _dafny.Set
                    d_1657_v8_ = _dafny.Set({(self).f17, (self).f17})
                    d_1658_v9_: _dafny.Map
                    d_1658_v9_ = _dafny.Map({d_1657_v8_: _dafny.Set({p0})})
                    d_1659_v10_: _dafny.Array
                    nw242_ = _dafny.Array(None, 13)
                    nw242_[int(0)] = (0) - ((d_1651_v3_).cardinality)
                    nw242_[int(1)] = len((d_1655_v6_ if False else d_1655_v6_))
                    nw242_[int(2)] = p0
                    nw242_[int(3)] = (p0) * (p0)
                    nw242_[int(4)] = len(_dafny.Set({(d_1656_v7_)[default__.safeIndex(p0, len(d_1656_v7_))], p0}))
                    nw242_[int(5)] = ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([811 for d_1660_i2_ in range(default__.abs(740))]))).cardinality) * (p0)
                    nw242_[int(6)] = len(d_1657_v8_)
                    nw242_[int(7)] = p0
                    nw242_[int(8)] = -766
                    nw242_[int(9)] = p0
                    nw242_[int(10)] = (d_1656_v7_)[default__.safeIndex(p0, len(d_1656_v7_))]
                    nw242_[int(11)] = p0
                    nw242_[int(12)] = len((_dafny.Map({d_1657_v8_: (self).f15})) | (d_1658_v9_))
                    d_1659_v10_ = nw242_
                    index285_ = default__.safeIndex(801, (d_1653_v5_).length(0))
                    (d_1653_v5_)[index285_] = p0
                    d_1661_v11_: _dafny.Seq
                    d_1661_v11_ = _dafny.SeqWithoutIsStrInference([d_1657_v8_, d_1657_v8_])
                    index286_ = default__.safeIndex(801, (d_1653_v5_).length(0))
                    (d_1653_v5_)[index286_] = (D9_DC22((self).f17, p0, False, (d_1661_v11_)[default__.safeIndex(p0, len(d_1661_v11_))], _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "apxjf")))).cf38
                    d_1651_v3_ = (_dafny.MultiSet(default__.fm38((self).f17, (self).f13, (self).f17, default__.fm30(596, globalState), globalState))).set(False, default__.abs(p0))
                    d_1662_v12_: _dafny.Array
                    nw243_ = _dafny.Array(_dafny.Map({}), 9)
                    d_1662_v12_ = nw243_
                    index287_ = default__.safeIndex(801, (d_1653_v5_).length(0))
                    rhs252_ = p0
                    rhs253_ = d_1662_v12_
                    rhs254_ = False
                    lhs207_ = d_1653_v5_
                    lhs208_ = default__.safeIndex(801, (d_1653_v5_).length(0))
                    lhs207_[lhs208_] = rhs252_
                    d_1662_v12_ = rhs253_
                    r0 = rhs254_
                    d_1663_v13_: _dafny.Map
                    d_1663_v13_ = _dafny.Map({p0: 814})
                    index288_ = default__.safeIndex(801, (d_1653_v5_).length(0))
                    (d_1653_v5_)[index288_] = default__.safeDivisionInt(((d_1663_v13_)[(d_1653_v5_)[default__.safeIndex(801, (d_1653_v5_).length(0))]] if ((d_1653_v5_)[default__.safeIndex(801, (d_1653_v5_).length(0))]) in (d_1663_v13_) else (d_1653_v5_)[default__.safeIndex(801, (d_1653_v5_).length(0))]), (d_1656_v7_)[default__.safeIndex(p0, len(d_1656_v7_))])
                    pass
            pass
        d_1664_v14_: _dafny.Array
        nw244_ = _dafny.Array(_dafny.Array(None, 0), 26)
        d_1664_v14_ = nw244_
        d_1665_v15_: _dafny.Array
        def lambda67_(d_1666_i3_):
            return (d_1666_i3_) + (683)

        init39_ = lambda67_
        nw245_ = _dafny.Array(None, 19)
        for i0_39_ in range(nw245_.length(0)):
            nw245_[i0_39_] = init39_(i0_39_)
        d_1665_v15_ = nw245_
        index289_ = default__.safeIndex(307, (d_1664_v14_).length(0))
        (d_1664_v14_)[index289_] = d_1665_v15_
        d_1667_v16_: _dafny.MultiSet
        d_1667_v16_ = _dafny.MultiSet([self.f12])
        d_1668_v18_: _dafny.Map
        d_1668_v18_ = _dafny.Map({self.f12: True})
        index290_ = default__.safeIndex(307, (d_1664_v14_).length(0))
        def iife154_():
            coll50_ = _dafny.Set()
            compr_50_: D1
            for compr_50_ in ((d_1667_v16_).set(self.f12, default__.abs(p0))).Elements:
                d_1669_v17_: D1 = compr_50_
                if (d_1669_v17_) in ((d_1667_v16_).set(self.f12, default__.abs(p0))):
                    coll50_ = coll50_.union(_dafny.Set([d_1669_v17_]))
            return _dafny.Set(coll50_)
        def iife155_():
            coll51_ = _dafny.Set()
            compr_51_: D1
            for compr_51_ in (d_1668_v18_).keys.Elements:
                d_1670_v19_: D1 = compr_51_
                if (d_1670_v19_) in (d_1668_v18_):
                    coll51_ = coll51_.union(_dafny.Set([d_1670_v19_]))
            return _dafny.Set(coll51_)
        (d_1664_v14_)[index290_] = (d_1665_v15_ if (iife154_()
        ) == (iife155_()
        ) else d_1665_v15_)
        d_1671_v20_: D7
        d_1671_v20_ = D7_DC16(p0)
        d_1672_v21_: D7
        d_1672_v21_ = D7_DC17(d_1671_v20_)
        pat_let_tv44_ = d_1671_v20_
        def iife156_(_pat_let52_0):
            def iife157_(d_1673_dt__update__tmp_h0_):
                def iife158_(_pat_let53_0):
                    def iife159_(d_1674_dt__update_hcf30_h0_):
                        return D7_DC17(d_1674_dt__update_hcf30_h0_)
                    return iife159_(_pat_let53_0)
                return iife158_(pat_let_tv44_)
            return iife157_(_pat_let52_0)
        source24_ = iife156_(d_1672_v21_)
        if source24_.is_DC16:
            d_1675___mcc_h0_ = source24_.cf29
            d_1676_cf29_ = d_1675___mcc_h0_
            d_1665_v15_ = d_1665_v15_
            d_1677_v22_: _dafny.Seq
            d_1677_v22_ = _dafny.SeqWithoutIsStrInference([(self).f13])
            d_1678_v23_: _dafny.Map
            d_1678_v23_ = _dafny.Map({p0: (self).f17})
            d_1679_v24_: _dafny.Map
            d_1679_v24_ = _dafny.Map({((d_1678_v23_)[d_1676_cf29_] if (d_1676_cf29_) in (d_1678_v23_) else (self).f17): p0})
            d_1680_v25_: _dafny.Seq
            d_1680_v25_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x')])
            d_1681_v26_: _dafny.Map
            d_1681_v26_ = _dafny.Map({(self).f17: d_1680_v25_})
            index291_ = default__.safeIndex(307, (d_1664_v14_).length(0))
            nw246_ = _dafny.Array(None, 7)
            nw246_[int(0)] = default__.safeModuloInt(d_1676_cf29_, (0) - (default__.fm0(d_1677_v22_, not((self).f13), (self).f17, ((d_1679_v24_)[(self).f17] if ((self).f17) in (d_1679_v24_) else d_1676_cf29_), globalState)))
            nw246_[int(1)] = d_1676_cf29_
            nw246_[int(2)] = len((d_1681_v26_) | (d_1681_v26_))
            nw246_[int(3)] = p0
            nw246_[int(4)] = p0
            nw246_[int(5)] = (0) - (len((self).f15))
            nw246_[int(6)] = d_1676_cf29_
            (d_1664_v14_)[index291_] = nw246_
            arr10_ = (d_1664_v14_)[default__.safeIndex(307, (d_1664_v14_).length(0))]
            index292_ = default__.safeIndex(281, ((d_1664_v14_)[default__.safeIndex(307, (d_1664_v14_).length(0))]).length(0))
            arr10_[index292_] = d_1676_cf29_
            arr11_ = (d_1664_v14_)[default__.safeIndex(307, (d_1664_v14_).length(0))]
            index293_ = default__.safeIndex(281, ((d_1664_v14_)[default__.safeIndex(307, (d_1664_v14_).length(0))]).length(0))
            arr11_[index293_] = 943
            d_1680_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))
        elif source24_.is_DC15:
            d_1682___mcc_h1_ = source24_.cf28
            d_1683_cf28_ = d_1682___mcc_h1_
            d_1684_v27_: _dafny.Map
            d_1684_v27_ = _dafny.Map({(self).f15: not((self).f17)})
            d_1684_v27_ = (d_1684_v27_).set((self).f15, (self).f13)
            d_1685_v28_: _dafny.Array
            def lambda68_(d_1686_i4_):
                return (self).f17

            init40_ = lambda68_
            nw247_ = _dafny.Array(None, 11)
            for i0_40_ in range(nw247_.length(0)):
                nw247_[i0_40_] = init40_(i0_40_)
            d_1685_v28_ = nw247_
            rhs255_ = (self).f17
            rhs256_ = d_1685_v28_
            r0 = rhs255_
            d_1685_v28_ = rhs256_
            d_1687_v29_: str
            d_1687_v29_ = _dafny.CodePoint('p')
            d_1688_v30_: _dafny.Seq
            d_1688_v30_ = _dafny.SeqWithoutIsStrInference([d_1687_v29_, _dafny.CodePoint('k'), d_1687_v29_])
            d_1689_v31_: C0
            nw248_ = C0()
            nw248_.ctor__((self).f13, (d_1688_v30_)[default__.safeIndex(p0, len(d_1688_v30_))], self.f12, True)
            d_1689_v31_ = nw248_
            d_1690_v32_: _dafny.Seq
            d_1690_v32_ = _dafny.SeqWithoutIsStrInference([d_1689_v31_, d_1689_v31_])
            d_1691_v33_: _dafny.Array
            nw249_ = _dafny.Array(None, 14)
            nw249_[int(0)] = (d_1690_v32_)[default__.safeIndex(2, len(d_1690_v32_))]
            nw249_[int(1)] = d_1689_v31_
            nw249_[int(2)] = d_1689_v31_
            nw249_[int(3)] = d_1689_v31_
            nw249_[int(4)] = (d_1689_v31_ if (self).f13 else d_1689_v31_)
            nw249_[int(5)] = d_1689_v31_
            nw249_[int(6)] = d_1689_v31_
            nw249_[int(7)] = d_1689_v31_
            nw249_[int(8)] = d_1689_v31_
            nw249_[int(9)] = d_1689_v31_
            nw249_[int(10)] = d_1689_v31_
            nw249_[int(11)] = d_1689_v31_
            nw249_[int(12)] = d_1689_v31_
            nw249_[int(13)] = d_1689_v31_
            d_1691_v33_ = nw249_
            index294_ = default__.safeIndex(528, (d_1691_v33_).length(0))
            (d_1691_v33_)[index294_] = d_1689_v31_
            index295_ = default__.safeIndex(182, (d_1665_v15_).length(0))
            (d_1665_v15_)[index295_] = p0
            d_1692_v34_: _dafny.Array
            nw250_ = _dafny.Array(_dafny.Seq({}), 19)
            d_1692_v34_ = nw250_
            d_1693_v35_: _dafny.Seq
            d_1693_v35_ = _dafny.SeqWithoutIsStrInference([d_1692_v34_])
            d_1694_v36_: D8
            d_1694_v36_ = D8_DC18((d_1693_v35_)[default__.safeIndex(p0, len(d_1693_v35_))])
            index296_ = default__.safeIndex(528, (d_1691_v33_).length(0))
            index297_ = default__.safeIndex(182, (d_1665_v15_).length(0))
            rhs257_ = 883
            rhs258_ = d_1689_v31_
            rhs259_ = (0) - (((p0) + (751) if (self).f13 else (p0) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ffp"))))))
            rhs260_ = D8_DC18(d_1692_v34_)
            rhs261_ = d_1689_v31_.f19
            lhs209_ = globalState
            lhs210_ = d_1691_v33_
            lhs211_ = default__.safeIndex(528, (d_1691_v33_).length(0))
            lhs212_ = d_1665_v15_
            lhs213_ = default__.safeIndex(182, (d_1665_v15_).length(0))
            lhs214_ = d_1689_v31_
            lhs209_.f7 = rhs257_
            lhs210_[lhs211_] = rhs258_
            lhs212_[lhs213_] = rhs259_
            d_1694_v36_ = rhs260_
            lhs214_.f19 = rhs261_
            if d_1689_v31_.f19:
                d_1695_v37_: _dafny.Seq
                d_1695_v37_ = _dafny.SeqWithoutIsStrInference([d_1689_v31_.f19, d_1689_v31_.f19])
                d_1696_v38_: D3
                d_1696_v38_ = D3_DC8((0) - (default__.fm0(d_1695_v37_, d_1689_v31_.f19, not(d_1689_v31_.f19), (d_1665_v15_)[default__.safeIndex(182, (d_1665_v15_).length(0))], globalState)))
                d_1697_v39_: _dafny.Set
                d_1697_v39_ = _dafny.Set({(d_1689_v31_).f20, (d_1689_v31_).f20})
                d_1698_v40_: _dafny.Map
                d_1698_v40_ = _dafny.Map({d_1687_v29_: default__.fm17(globalState)})
                pat_let_tv45_ = d_1665_v15_
                pat_let_tv46_ = d_1665_v15_
                pat_let_tv47_ = d_1665_v15_
                pat_let_tv48_ = d_1665_v15_
                index298_ = default__.safeIndex(182, (d_1665_v15_).length(0))
                def iife160_(_pat_let54_0):
                    def iife161_(d_1699_dt__update__tmp_h1_):
                        def iife162_(_pat_let55_0):
                            def iife163_(d_1700_dt__update_hcf3_h0_):
                                def iife164_(_pat_let56_0):
                                    def iife165_(d_1701_dt__update_hcf2_h0_):
                                        return D1_DC2(d_1701_dt__update_hcf2_h0_, d_1700_dt__update_hcf3_h0_, (d_1699_dt__update__tmp_h1_).cf4)
                                    return iife165_(_pat_let56_0)
                                return iife164_((pat_let_tv48_)[default__.safeIndex(182, (pat_let_tv47_).length(0))])
                            return iife163_(_pat_let55_0)
                        return iife162_((pat_let_tv46_)[default__.safeIndex(182, (pat_let_tv45_).length(0))])
                    return iife161_(_pat_let54_0)
                rhs262_ = default__.fm49(self.f16, d_1697_v39_, globalState)
                rhs263_ = (d_1689_v31_).fm16((self).fm12(d_1698_v40_, iife160_(self.f12), globalState), p0, globalState)
                rhs264_ = d_1689_v31_.f19
                lhs215_ = d_1665_v15_
                lhs216_ = default__.safeIndex(182, (d_1665_v15_).length(0))
                lhs217_ = d_1689_v31_
                d_1696_v38_ = rhs262_
                lhs215_[lhs216_] = rhs263_
                lhs217_.f19 = rhs264_
                index299_ = default__.safeIndex(182, (d_1665_v15_).length(0))
                (d_1665_v15_)[index299_] = (0) - ((0) - (((d_1683_cf28_)[default__.safeIndex((d_1665_v15_)[default__.safeIndex(182, (d_1665_v15_).length(0))], len(d_1683_cf28_))] if (default__.fm17(globalState) if d_1689_v31_.f19 else d_1689_v31_.f19) else default__.fm0(_dafny.SeqWithoutIsStrInference([d_1689_v31_.f19, False, (d_1695_v37_)[default__.safeIndex(len(d_1695_v37_), len(d_1695_v37_))]]), (self).f17, (self).f17, (d_1665_v15_)[default__.safeIndex(182, (d_1665_v15_).length(0))], globalState))))
                d_1702_v41_: _dafny.Map
                d_1702_v41_ = _dafny.Map({p0: d_1695_v37_})
                def iife166_():
                    coll52_ = _dafny.Set()
                    compr_52_: int
                    for compr_52_ in _dafny.IntegerRange(736, 816):
                        d_1703_v42_: int = compr_52_
                        if ((736) <= (d_1703_v42_)) and ((d_1703_v42_) < (816)):
                            coll52_ = coll52_.union(_dafny.Set([default__.safeModuloInt(d_1703_v42_, len(_dafny.SeqWithoutIsStrInference([d_1689_v31_.f19])))]))
                    return _dafny.Set(coll52_)
                d_1702_v41_ = (d_1702_v41_).set(len(iife166_()
                ), (d_1695_v37_) + (d_1695_v37_))
                d_1704_v43_: _dafny.MultiSet
                d_1704_v43_ = _dafny.MultiSet([(d_1665_v15_)[default__.safeIndex(182, (d_1665_v15_).length(0))]])
                d_1705_v44_: D6
                d_1705_v44_ = D6_DC11(d_1704_v43_)
                d_1706_v45_: D6
                d_1706_v45_ = D6_DC14(d_1705_v44_)
                d_1707_v46_: _dafny.Array
                nw251_ = _dafny.Array(_dafny.Map({}), 25)
                d_1707_v46_ = nw251_
                index300_ = default__.safeIndex(303, (d_1707_v46_).length(0))
                (d_1707_v46_)[index300_] = default__.fm50(globalState)
                d_1708_v47_: _dafny.MultiSet
                d_1708_v47_ = _dafny.MultiSet([(d_1689_v31_.f19 if not(d_1689_v31_.f19) else ((d_1698_v40_)[d_1687_v29_] if (d_1687_v29_) in (d_1698_v40_) else (self).f17)), False, False, not(((self).f17) and ((self).f17)), (d_1688_v30_) < ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ido"))).set(default__.safeIndex((d_1665_v15_)[default__.safeIndex(182, (d_1665_v15_).length(0))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ido")))), (d_1689_v31_).f20))])
                d_1709_v48_: _dafny.Map
                d_1709_v48_ = _dafny.Map({p0: (self).f17})
                index301_ = default__.safeIndex(303, (d_1707_v46_).length(0))
                rhs265_ = (d_1689_v31_).f20
                rhs266_ = D6_DC14(d_1705_v44_)
                rhs267_ = ((d_1709_v48_).set(376, d_1689_v31_.f19)) | (_dafny.Map({(d_1665_v15_)[default__.safeIndex(182, (d_1665_v15_).length(0))]: d_1689_v31_.f19}))
                rhs268_ = (d_1708_v47_ if ((self).f15).ispropersubset((self).f15) else d_1708_v47_)
                lhs218_ = d_1707_v46_
                lhs219_ = default__.safeIndex(303, (d_1707_v46_).length(0))
                d_1687_v29_ = rhs265_
                d_1706_v45_ = rhs266_
                lhs218_[lhs219_] = rhs267_
                d_1708_v47_ = rhs268_
                rhs269_ = (d_1665_v15_)[default__.safeIndex(182, (d_1665_v15_).length(0))]
                rhs270_ = _dafny.Map({d_1687_v29_: (self).f13})
                rhs271_ = (d_1706_v45_ if not(False) else d_1706_v45_)
                rhs272_ = (d_1695_v37_)[default__.safeIndex(((d_1665_v15_)[default__.safeIndex(182, (d_1665_v15_).length(0))] if (self).f13 else (d_1665_v15_)[default__.safeIndex(182, (d_1665_v15_).length(0))]), len(d_1695_v37_))]
                lhs220_ = globalState
                lhs220_.f7 = rhs269_
                d_1698_v40_ = rhs270_
                d_1706_v45_ = rhs271_
                r0 = rhs272_
            elif True:
                d_1710_v49_: C1
                nw252_ = C1()
                nw252_.ctor__((d_1689_v31_).f20, self.f12, not((self).f13))
                d_1710_v49_ = nw252_
                d_1711_v50_: _dafny.Map
                d_1711_v50_ = _dafny.Map({p0: (d_1665_v15_)[default__.safeIndex(182, (d_1665_v15_).length(0))]})
                d_1711_v50_ = d_1711_v50_
                (d_1689_v31_).f19 = d_1689_v31_.f19
                d_1712_v51_: _dafny.Map
                d_1712_v51_ = _dafny.Map({d_1689_v31_.f19: p0})
                d_1712_v51_ = (d_1712_v51_).set((self).f13, p0)
                (d_1689_v31_).f19 = not ((self).f17) or ((d_1688_v30_) == (d_1688_v30_))
        elif True:
            d_1713___mcc_h2_ = source24_.cf30
            d_1714_cf30_ = d_1713___mcc_h2_
            index302_ = default__.safeIndex(104, ((self).f14).length(0))
            ((self).f14)[index302_] = (self).f17
            index303_ = default__.safeIndex(104, ((self).f14).length(0))
            ((self).f14)[index303_] = (self).f17
            if ((self).f14)[default__.safeIndex(104, ((self).f14).length(0))]:
                d_1715_v52_: _dafny.Array
                nw253_ = _dafny.Array(_dafny.CodePoint('D'), 9)
                d_1715_v52_ = nw253_
                d_1716_v53_: str
                d_1716_v53_ = _dafny.CodePoint('l')
                index304_ = default__.safeIndex(404, (d_1715_v52_).length(0))
                (d_1715_v52_)[index304_] = d_1716_v53_
                d_1717_v54_: _dafny.Seq
                d_1717_v54_ = _dafny.SeqWithoutIsStrInference([(self).f17])
                index305_ = default__.safeIndex(404, (d_1715_v52_).length(0))
                rhs273_ = (p0) < (default__.fm0(d_1717_v54_, ((self).f14)[default__.safeIndex(104, ((self).f14).length(0))], ((self).f14)[default__.safeIndex(104, ((self).f14).length(0))], p0, globalState))
                rhs274_ = d_1716_v53_
                rhs275_ = default__.safeDivisionInt(p0, p0)
                lhs221_ = d_1715_v52_
                lhs222_ = default__.safeIndex(404, (d_1715_v52_).length(0))
                lhs223_ = globalState
                r0 = rhs273_
                lhs221_[lhs222_] = rhs274_
                lhs223_.f7 = rhs275_
                (globalState).f7 = (default__.safeModuloInt(832, p0)) * (p0)
                d_1718_v55_: _dafny.Map
                d_1718_v55_ = _dafny.Map({(0) - (p0): (self).f17})
                d_1719_v56_: _dafny.Seq
                d_1719_v56_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uwc"))
                d_1720_v57_: _dafny.Array
                nw254_ = _dafny.Array(False, 22)
                d_1720_v57_ = nw254_
                pat_let_tv49_ = d_1719_v56_
                d_1721_v58_: C2
                nw255_ = C2()
                def iife167_(_pat_let57_0):
                    def iife168_(d_1722_dt__update__tmp_h2_):
                        def iife169_(_pat_let58_0):
                            def iife170_(d_1723_dt__update_hcf1_h0_):
                                return D1_DC1(d_1723_dt__update_hcf1_h0_)
                            return iife170_(_pat_let58_0)
                        return iife169_(pat_let_tv49_)
                    return iife168_(_pat_let57_0)
                nw255_.ctor__(p0, ((d_1718_v55_)[-780] if (-780) in (d_1718_v55_) else (self).f13), iife167_(self.f16), (self).f17, d_1720_v57_, default__.fm47(not((self).f13), globalState), self.f12, ((self).f14)[default__.safeIndex(104, ((self).f14).length(0))])
                d_1721_v58_ = nw255_
                d_1724_v59_: C3
                nw256_ = C3()
                nw256_.ctor__((self).f17, (self.f16 if ((self).f14)[default__.safeIndex(104, ((self).f14).length(0))] else D1_DC1(d_1719_v56_)), (self).f13, d_1720_v57_, ((self).f15) - ((self).f15), D1_DC2(p0, (d_1721_v58_).f21, ((self).f14)[default__.safeIndex(104, ((self).f14).length(0))]), (self).f13)
                d_1724_v59_ = nw256_
                index306_ = default__.safeIndex(104, ((self).f14).length(0))
                ((self).f14)[index306_] = (p0) != (p0)
            elif True:
                index307_ = default__.safeIndex(104, ((self).f14).length(0))
                ((self).f14)[index307_] = ((self).f14)[default__.safeIndex(104, ((self).f14).length(0))]
                d_1725_v60_: _dafny.MultiSet
                d_1725_v60_ = _dafny.MultiSet([(self).f13, ((self).f14)[default__.safeIndex(104, ((self).f14).length(0))]])
                d_1725_v60_ = default__.fm40(globalState)
                d_1726_v61_: _dafny.Map
                d_1726_v61_ = _dafny.Map({5: ((self).f14)[default__.safeIndex(104, ((self).f14).length(0))]})
                d_1727_v62_: _dafny.Seq
                d_1727_v62_ = _dafny.SeqWithoutIsStrInference([(self).f13])
                d_1726_v61_ = (d_1726_v61_).set(p0, (d_1727_v62_)[default__.safeIndex(p0, len(d_1727_v62_))])
                d_1728_v63_: _dafny.MultiSet
                d_1728_v63_ = _dafny.MultiSet([(0) - (p0), (0) - (p0)])
                d_1729_v64_: D6
                d_1729_v64_ = D6_DC11(d_1728_v63_)
                d_1730_v65_: D6
                d_1730_v65_ = D6_DC14(D6_DC11(((d_1729_v64_).cf17).set(p0, default__.abs(default__.fm0(_dafny.SeqWithoutIsStrInference([not((self).f17)]), (self).f13, False, p0, globalState)))))
                d_1731_v66_: _dafny.MultiSet
                out12_: _dafny.MultiSet
                out12_ = default__.m0(default__.fm13(default__.fm32(d_1730_v65_, globalState), globalState), (d_1727_v62_)[default__.safeIndex(156, len(d_1727_v62_))], globalState)
                d_1731_v66_ = out12_
                d_1732_v67_: _dafny.Map
                d_1732_v67_ = _dafny.Map({d_1727_v62_: not((self).f13)})
                d_1733_v68_: _dafny.Seq
                d_1733_v68_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))
                d_1734_v69_: str
                d_1734_v69_ = _dafny.CodePoint('k')
                d_1735_v70_: _dafny.Map
                d_1735_v70_ = _dafny.Map({d_1734_v69_: ((self).f14)[default__.safeIndex(104, ((self).f14).length(0))]})
                d_1736_v71_: _dafny.Map
                d_1736_v71_ = _dafny.Map({len(d_1733_v68_): d_1734_v69_})
                d_1737_v72_: _dafny.MultiSet
                d_1737_v72_ = _dafny.MultiSet([(d_1733_v68_).set(default__.safeIndex(len((_dafny.Map({default__.fm0((d_1727_v62_).set(default__.safeIndex(default__.fm0(d_1727_v62_, (self).f17, ((self).f14)[default__.safeIndex(104, ((self).f14).length(0))], p0, globalState), len(d_1727_v62_)), (self).f17), (self).f17, (self).f13, p0, globalState): p0})).set(p0, p0)), len(d_1733_v68_)), default__.fm33(((d_1735_v70_)[((d_1736_v71_)[len(d_1733_v68_)] if (len(d_1733_v68_)) in (d_1736_v71_) else d_1734_v69_)] if (((d_1736_v71_)[len(d_1733_v68_)] if (len(d_1733_v68_)) in (d_1736_v71_) else d_1734_v69_)) in (d_1735_v70_) else default__.fm17(globalState)), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jmgntc"))), globalState)), d_1733_v68_, d_1733_v68_])
                d_1738_v73_: D12
                d_1738_v73_ = D12_DC29(d_1732_v67_)
                index308_ = default__.safeIndex(104, ((self).f14).length(0))
                rhs276_ = ((148 if False else p0)) == ((0) - ((0) - (((d_1737_v72_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qxcd"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qxcd"))) in (d_1737_v72_) else (0) - (p0)))))
                rhs277_ = p0
                rhs278_ = (d_1738_v73_).cf57
                lhs224_ = (self).f14
                lhs225_ = default__.safeIndex(104, ((self).f14).length(0))
                lhs226_ = globalState
                lhs224_[lhs225_] = rhs276_
                lhs226_.f7 = rhs277_
                d_1732_v67_ = rhs278_
            d_1739_v74_: _dafny.Map
            d_1739_v74_ = _dafny.Map({p0: ((self).f14)[default__.safeIndex(104, ((self).f14).length(0))]})
            d_1740_v75_: _dafny.Seq
            d_1740_v75_ = _dafny.SeqWithoutIsStrInference([True])
            index309_ = default__.safeIndex(104, ((self).f14).length(0))
            index310_ = default__.safeIndex(104, ((self).f14).length(0))
            rhs279_ = 243
            rhs280_ = True
            rhs281_ = ((d_1739_v74_)[p0] if (p0) in (d_1739_v74_) else (d_1740_v75_) <= (d_1740_v75_))
            lhs227_ = globalState
            lhs228_ = (self).f14
            lhs229_ = default__.safeIndex(104, ((self).f14).length(0))
            lhs230_ = (self).f14
            lhs231_ = default__.safeIndex(104, ((self).f14).length(0))
            lhs227_.f7 = rhs279_
            lhs228_[lhs229_] = rhs280_
            lhs230_[lhs231_] = rhs281_
            source25_ = default__.fm51(globalState)
            if source25_.is_DC21:
                index311_ = default__.safeIndex(384, (d_1665_v15_).length(0))
                (d_1665_v15_)[index311_] = p0
                index312_ = default__.safeIndex(384, (d_1665_v15_).length(0))
                (d_1665_v15_)[index312_] = p0
                d_1741_v76_: str
                d_1741_v76_ = _dafny.CodePoint('l')
                d_1742_v77_: _dafny.MultiSet
                out13_: _dafny.MultiSet
                out13_ = default__.m0(d_1741_v76_, False, globalState)
                d_1742_v77_ = out13_
                (globalState).f7 = (0) - (p0)
                index313_ = default__.safeIndex(104, ((self).f14).length(0))
                ((self).f14)[index313_] = not((self).f17)
            elif source25_.is_DC22:
                d_1743___mcc_h3_ = source25_.cf37
                d_1744___mcc_h4_ = source25_.cf38
                d_1745___mcc_h5_ = source25_.cf39
                d_1746___mcc_h6_ = source25_.cf40
                d_1747___mcc_h7_ = source25_.cf41
                d_1748_cf41_ = d_1747___mcc_h7_
                d_1749_cf40_ = d_1746___mcc_h6_
                d_1750_cf39_ = d_1745___mcc_h5_
                d_1751_cf38_ = d_1744___mcc_h4_
                d_1752_cf37_ = d_1743___mcc_h3_
                index314_ = default__.safeIndex(141, (d_1665_v15_).length(0))
                (d_1665_v15_)[index314_] = (0) - (d_1751_cf38_)
                index315_ = default__.safeIndex(141, (d_1665_v15_).length(0))
                (d_1665_v15_)[index315_] = 879
                (globalState).f2 = (d_1740_v75_) + (_dafny.SeqWithoutIsStrInference([d_1752_cf37_, (self).f17, default__.fm30((d_1665_v15_)[default__.safeIndex(141, (d_1665_v15_).length(0))], globalState)]))
                d_1753_v78_: _dafny.Array
                nw257_ = _dafny.Array(False, 9)
                d_1753_v78_ = nw257_
                rhs282_ = d_1753_v78_
                rhs283_ = p0
                lhs232_ = globalState
                d_1753_v78_ = rhs282_
                lhs232_.f7 = rhs283_
                r0 = True
            elif source25_.is_DC23:
                d_1754___mcc_h8_ = source25_.cf42
                d_1755___mcc_h9_ = source25_.cf43
                d_1756___mcc_h10_ = source25_.cf44
                d_1757___mcc_h11_ = source25_.cf45
                d_1758___mcc_h12_ = source25_.cf46
                d_1759_cf46_ = d_1758___mcc_h12_
                d_1760_cf45_ = d_1757___mcc_h11_
                d_1761_cf44_ = d_1756___mcc_h10_
                d_1762_cf43_ = d_1755___mcc_h9_
                d_1763_cf42_ = d_1754___mcc_h8_
                index316_ = default__.safeIndex(104, ((self).f14).length(0))
                ((self).f14)[index316_] = (True if (self).f13 else (self).f17)
                d_1764_v79_: D6
                d_1764_v79_ = D6_DC13(not((self).f17), ((self).f14)[default__.safeIndex(104, ((self).f14).length(0))], 556, d_1759_cf46_, (self).f17)
                d_1764_v79_ = d_1764_v79_
                index317_ = default__.safeIndex(104, ((self).f14).length(0))
                ((self).f14)[index317_] = ((self).f14)[default__.safeIndex(104, ((self).f14).length(0))]
                d_1765_v80_: _dafny.Array
                def lambda69_(d_1766_i5_):
                    return ((self).f14)[default__.safeIndex(104, ((self).f14).length(0))]

                init41_ = lambda69_
                nw258_ = _dafny.Array(None, 4)
                for i0_41_ in range(nw258_.length(0)):
                    nw258_[i0_41_] = init41_(i0_41_)
                d_1765_v80_ = nw258_
                d_1767_v81_: _dafny.MultiSet
                d_1767_v81_ = _dafny.MultiSet([d_1765_v80_])
                d_1768_v82_: _dafny.MultiSet
                d_1768_v82_ = d_1767_v81_
                d_1769_v83_: _dafny.Map
                d_1769_v83_ = _dafny.Map({d_1767_v81_: len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uothuadbe"))))})
                d_1770_v84_: str
                d_1770_v84_ = _dafny.CodePoint('j')
                d_1771_v85_: C1
                nw259_ = C1()
                nw259_.ctor__(d_1770_v84_, self.f12, (self).f17)
                d_1771_v85_ = nw259_
                d_1772_v86_: _dafny.Seq
                d_1772_v86_ = _dafny.SeqWithoutIsStrInference([d_1771_v85_])
                d_1769_v83_ = (d_1769_v83_).set(d_1768_v82_, len(d_1772_v86_))
            elif True:
                d_1773___mcc_h13_ = source25_.cf36
                d_1774_cf36_ = d_1773___mcc_h13_
                index318_ = default__.safeIndex(104, ((self).f14).length(0))
                ((self).f14)[index318_] = (((self).f15) | (_dafny.Set({p0}))).isdisjoint((self).f15)
                index319_ = default__.safeIndex(22, (d_1665_v15_).length(0))
                (d_1665_v15_)[index319_] = p0
                index320_ = default__.safeIndex(22, (d_1665_v15_).length(0))
                (d_1665_v15_)[index320_] = p0
                d_1775_v87_: _dafny.MultiSet
                d_1775_v87_ = _dafny.MultiSet([p0])
                d_1776_v88_: _dafny.MultiSet
                d_1776_v88_ = _dafny.MultiSet([(d_1665_v15_)[default__.safeIndex(22, (d_1665_v15_).length(0))], ((d_1775_v87_)[(0) - (p0)] if ((0) - (p0)) in (d_1775_v87_) else p0)])
                d_1777_v89_: _dafny.Set
                d_1777_v89_ = _dafny.Set({(d_1775_v87_).ispropersubset(d_1775_v87_)})
                d_1778_v90_: _dafny.Map
                d_1778_v90_ = _dafny.Map({p0: _dafny.Set({(self).f13})})
                d_1779_v91_: _dafny.Seq
                d_1779_v91_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kxckk"))
                d_1780_v92_: str
                d_1780_v92_ = _dafny.CodePoint('a')
                d_1781_v93_: _dafny.Seq
                d_1781_v93_ = _dafny.SeqWithoutIsStrInference([p0])
                rhs284_ = ((self).f14)[default__.safeIndex(104, ((self).f14).length(0))]
                rhs285_ = ((d_1778_v90_)[len((d_1779_v91_).set(default__.safeIndex(p0, len(d_1779_v91_)), d_1780_v92_))] if (len((d_1779_v91_).set(default__.safeIndex(p0, len(d_1779_v91_)), d_1780_v92_))) in (d_1778_v90_) else _dafny.Set({((self).f14)[default__.safeIndex(104, ((self).f14).length(0))], ((self).f14)[default__.safeIndex(104, ((self).f14).length(0))]}))
                rhs286_ = ((_dafny.MultiSet(d_1781_v93_)).intersection((d_1776_v88_).set(901, default__.abs((d_1665_v15_)[default__.safeIndex(22, (d_1665_v15_).length(0))])))) - (d_1775_v87_)
                r0 = rhs284_
                d_1777_v89_ = rhs285_
                d_1776_v88_ = rhs286_
                index321_ = default__.safeIndex(22, (d_1665_v15_).length(0))
                (d_1665_v15_)[index321_] = default__.fm0(_dafny.SeqWithoutIsStrInference([(self).f13]), ((self).f14)[default__.safeIndex(104, ((self).f14).length(0))], True, (d_1665_v15_)[default__.safeIndex(22, (d_1665_v15_).length(0))], globalState)
        r0 = ((self).f13) or ((self).f17)
        index322_ = default__.safeIndex(701, ((self).f14).length(0))
        ((self).f14)[index322_] = not((self).f17)
        index323_ = default__.safeIndex(701, ((self).f14).length(0))
        ((self).f14)[index323_] = (self).f17
        index324_ = default__.safeIndex(701, ((self).f14).length(0))
        ((self).f14)[index324_] = ((self).f14)[default__.safeIndex(701, ((self).f14).length(0))]
        r0 = ((self).f14)[default__.safeIndex(701, ((self).f14).length(0))]
        return r0

    def m8(self, p0, p1, globalState):
        r0: bool = False
        d_1782_v0_: _dafny.Seq
        d_1782_v0_ = _dafny.SeqWithoutIsStrInference([(self).f17])
        (globalState).f7 = default__.fm0(d_1782_v0_, (self).f17, (self).f17, p0, globalState)
        d_1783_i0_: int
        d_1783_i0_ = 0
        with _dafny.label("11"):
            while not((self).f17):
                with _dafny.c_label("11"):
                    if (d_1783_i0_) >= (100):
                        raise _dafny.Break("11")
                    d_1783_i0_ = (d_1783_i0_) + (1)
                    d_1784_v1_: _dafny.MultiSet
                    d_1784_v1_ = _dafny.MultiSet([(self).f17])
                    r0 = ((d_1784_v1_) | (d_1784_v1_)).issubset(d_1784_v1_)
                    (globalState).f7 = p0
                    d_1785_v2_: _dafny.Array
                    nw260_ = _dafny.Array(_dafny.Array(None, 0), 2)
                    d_1785_v2_ = nw260_
                    d_1786_v3_: _dafny.Array
                    nw261_ = _dafny.Array(_dafny.MultiSet({}), 24)
                    d_1786_v3_ = nw261_
                    index325_ = default__.safeIndex(736, (d_1785_v2_).length(0))
                    (d_1785_v2_)[index325_] = d_1786_v3_
                    d_1787_v4_: _dafny.Array
                    nw262_ = _dafny.Array(int(0), 19)
                    d_1787_v4_ = nw262_
                    d_1788_v5_: str
                    d_1788_v5_ = _dafny.CodePoint('h')
                    index326_ = default__.safeIndex(888, (d_1787_v4_).length(0))
                    (d_1787_v4_)[index326_] = (self).fm9(d_1788_v5_, globalState)
                    d_1789_v6_: _dafny.Map
                    d_1789_v6_ = _dafny.Map({p0: (self).f17})
                    d_1790_v7_: _dafny.Map
                    d_1790_v7_ = _dafny.Map({default__.fm17(globalState): (_dafny.MultiSet([d_1789_v6_, _dafny.Map({p0: (self).f17}), _dafny.Map({p0: (self).f13})])).issubset(_dafny.MultiSet([d_1789_v6_]))})
                    d_1791_v8_: D2
                    d_1791_v8_ = D2_DC4(d_1787_v4_, (self).f17, p0, (self).f13, (self).f14)
                    index327_ = default__.safeIndex(736, (d_1785_v2_).length(0))
                    index328_ = default__.safeIndex(888, (d_1787_v4_).length(0))
                    rhs287_ = (self).f17
                    rhs288_ = ((d_1790_v7_)[(self).f17] if ((self).f17) in (d_1790_v7_) else (self).f13)
                    rhs289_ = d_1786_v3_
                    rhs290_ = (d_1791_v8_).cf8
                    rhs291_ = (p0) * (2)
                    lhs233_ = d_1785_v2_
                    lhs234_ = default__.safeIndex(736, (d_1785_v2_).length(0))
                    lhs235_ = d_1787_v4_
                    lhs236_ = default__.safeIndex(888, (d_1787_v4_).length(0))
                    lhs237_ = globalState
                    r0 = rhs287_
                    r0 = rhs288_
                    lhs233_[lhs234_] = rhs289_
                    lhs235_[lhs236_] = rhs290_
                    lhs237_.f7 = rhs291_
                    index329_ = default__.safeIndex(888, (d_1787_v4_).length(0))
                    (d_1787_v4_)[index329_] = (d_1787_v4_)[default__.safeIndex(888, (d_1787_v4_).length(0))]
                    pass
            pass
        d_1792_v9_: str
        d_1792_v9_ = _dafny.CodePoint('e')
        d_1793_v10_: _dafny.Map
        d_1793_v10_ = _dafny.Map({d_1792_v9_: _dafny.CodePoint('i')})
        d_1793_v10_ = (d_1793_v10_).set(d_1792_v9_, d_1792_v9_)
        d_1794_v11_: _dafny.Array
        nw263_ = _dafny.Array(_dafny.Array(None, 0), 20)
        d_1794_v11_ = nw263_
        d_1795_v12_: _dafny.Array
        d_1795_v12_ = d_1794_v11_
        d_1796_v13_: _dafny.Map
        d_1796_v13_ = _dafny.Map({(d_1795_v12_): not(False)})
        d_1797_v14_: _dafny.Set
        d_1797_v14_ = _dafny.Set({(self).f17})
        d_1796_v13_ = (d_1796_v13_).set(d_1794_v11_, (d_1797_v14_) == (d_1797_v14_))
        (globalState).f7 = (self).fm9(d_1792_v9_, globalState)
        hi12_ = p0
        for d_1798_i1_ in range(p0, hi12_):
            d_1799_v15_: _dafny.Map
            d_1799_v15_ = _dafny.Map({(self).fm11(globalState): d_1798_i1_})
            d_1799_v15_ = (d_1799_v15_).set((self).f13, d_1798_i1_)
            (globalState).f7 = d_1798_i1_
            d_1800_v16_: D7
            d_1800_v16_ = D7_DC16(d_1798_i1_)
            d_1801_v17_: D9
            d_1801_v17_ = D9_DC23(default__.safeDivisionInt(p0, p0), p0, ((d_1800_v16_).cf29) - (len(_dafny.Set({p0}))), d_1798_i1_, default__.safeDivisionInt(default__.fm0(d_1782_v0_, (self).f17, not(False), p0, globalState), d_1798_i1_))
            d_1802_v18_: _dafny.Array
            nw264_ = _dafny.Array(_dafny.Array(None, 0), 23)
            d_1802_v18_ = nw264_
            d_1803_v19_: _dafny.Array
            nw265_ = _dafny.Array(None, 27)
            nw265_[int(0)] = d_1792_v9_
            nw265_[int(1)] = d_1792_v9_
            nw265_[int(2)] = d_1792_v9_
            nw265_[int(3)] = _dafny.CodePoint('f')
            nw265_[int(4)] = d_1792_v9_
            nw265_[int(5)] = _dafny.CodePoint('j')
            nw265_[int(6)] = d_1792_v9_
            nw265_[int(7)] = d_1792_v9_
            nw265_[int(8)] = d_1792_v9_
            nw265_[int(9)] = d_1792_v9_
            nw265_[int(10)] = d_1792_v9_
            nw265_[int(11)] = d_1792_v9_
            nw265_[int(12)] = d_1792_v9_
            nw265_[int(13)] = d_1792_v9_
            nw265_[int(14)] = d_1792_v9_
            nw265_[int(15)] = d_1792_v9_
            nw265_[int(16)] = d_1792_v9_
            nw265_[int(17)] = d_1792_v9_
            nw265_[int(18)] = d_1792_v9_
            nw265_[int(19)] = d_1792_v9_
            nw265_[int(20)] = _dafny.CodePoint('b')
            nw265_[int(21)] = d_1792_v9_
            nw265_[int(22)] = d_1792_v9_
            nw265_[int(23)] = d_1792_v9_
            nw265_[int(24)] = d_1792_v9_
            nw265_[int(25)] = d_1792_v9_
            nw265_[int(26)] = d_1792_v9_
            d_1803_v19_ = nw265_
            index330_ = default__.safeIndex(698, (d_1802_v18_).length(0))
            (d_1802_v18_)[index330_] = (D13_DC33(d_1803_v19_)).cf70
            d_1804_v20_: _dafny.MultiSet
            d_1804_v20_ = _dafny.MultiSet([True, (self).f17, (self).f13])
            d_1805_v21_: _dafny.Set
            d_1805_v21_ = _dafny.Set({d_1804_v20_, d_1804_v20_})
            d_1806_v22_: D6
            d_1806_v22_ = D6_DC11((_dafny.MultiSet([p0])).set(p0, default__.abs(d_1798_i1_)))
            d_1807_v23_: D6
            d_1807_v23_ = D6_DC14(d_1806_v22_)
            index331_ = default__.safeIndex(698, (d_1802_v18_).length(0))
            rhs292_ = d_1801_v17_
            rhs293_ = d_1803_v19_
            rhs294_ = d_1792_v9_
            rhs295_ = (d_1805_v21_) - (d_1805_v21_)
            rhs296_ = default__.fm30((len(default__.fm32(D6_DC14(d_1806_v22_), globalState)) if (self).f17 else d_1798_i1_), globalState)
            lhs238_ = d_1802_v18_
            lhs239_ = default__.safeIndex(698, (d_1802_v18_).length(0))
            d_1801_v17_ = rhs292_
            lhs238_[lhs239_] = rhs293_
            d_1792_v9_ = rhs294_
            d_1805_v21_ = rhs295_
            r0 = rhs296_
            d_1808_v24_: C4
            nw266_ = C4()
            nw266_.ctor__()
            d_1808_v24_ = nw266_
        r0 = ((self).f15).ispropersubset((self).f15)
        return r0

    def m9(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: bool = False
        d_1809_v0_: _dafny.Seq
        d_1809_v0_ = _dafny.SeqWithoutIsStrInference([(self).f17])
        d_1810_v1_: _dafny.Seq
        d_1810_v1_ = _dafny.SeqWithoutIsStrInference([p1, p1])
        d_1811_v2_: _dafny.Map
        d_1811_v2_ = _dafny.Map({p3: len(d_1810_v1_)})
        d_1812_v3_: _dafny.Map
        d_1812_v3_ = _dafny.Map({p1: (d_1811_v2_).set((self).f13, p1)})
        d_1813_i0_: int
        d_1813_i0_ = 0
        with _dafny.label("12"):
            while (len((d_1809_v0_) + (d_1809_v0_))) == (len(d_1812_v3_)):
                with _dafny.c_label("12"):
                    if (d_1813_i0_) >= (100):
                        raise _dafny.Break("12")
                    d_1813_i0_ = (d_1813_i0_) + (1)
                    d_1814_v4_: T2
                    nw267_ = C5()
                    nw267_.ctor__(not((self).f13), (self).f14, (self).f15, D1_DC2(p1, (_dafny.MultiSet([p3])).cardinality, (self).f17), p2)
                    d_1814_v4_ = nw267_
                    d_1814_v4_ = d_1814_v4_
                    d_1815_v5_: _dafny.Array
                    nw268_ = _dafny.Array(_dafny.CodePoint('D'), 20)
                    d_1815_v5_ = nw268_
                    d_1816_v6_: _dafny.Map
                    d_1816_v6_ = _dafny.Map({(d_1814_v4_).f13: d_1815_v5_})
                    (globalState).f7 = (p1) * ((0) - (default__.safeDivisionInt(len(d_1816_v6_), p1)))
                    d_1817_v7_: _dafny.Seq
                    d_1817_v7_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "thmbvav"))])
                    d_1818_v8_: _dafny.Seq
                    d_1818_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vhrtlm"))
                    d_1817_v7_ = (d_1817_v7_) + ((d_1817_v7_).set(default__.safeIndex(p1, len(d_1817_v7_)), d_1818_v8_))
                    r1 = (p1) == (p1)
                    pass
            pass
        def iife171_():
            coll53_ = _dafny.Map()
            compr_53_: int
            for compr_53_ in (_dafny.SeqWithoutIsStrInference([len(d_1809_v0_) for d_1819_i1_ in range(default__.abs(-360))])).Elements:
                d_1820_v9_: int = compr_53_
                if (d_1820_v9_) in (_dafny.SeqWithoutIsStrInference([len(d_1809_v0_) for d_1819_i1_ in range(default__.abs(-360))])):
                    coll53_[default__.safeDivisionInt(d_1820_v9_, p1)] = p0
            return _dafny.Map(coll53_)
        if (p1) not in (iife171_()
        ):
            r1 = False
            if not (p3) or ((self).f17):
                d_1821_v10_: _dafny.Array
                def lambda70_(d_1822_p1_):
                    def lambda71_(d_1823_i2_):
                        return default__.safeDivisionInt(d_1823_i2_, d_1822_p1_)

                    return lambda71_

                init42_ = lambda70_(p1)
                nw269_ = _dafny.Array(None, 24)
                for i0_42_ in range(nw269_.length(0)):
                    nw269_[i0_42_] = init42_(i0_42_)
                d_1821_v10_ = nw269_
                index332_ = default__.safeIndex(441, (d_1821_v10_).length(0))
                (d_1821_v10_)[index332_] = p1
                index333_ = default__.safeIndex(441, (d_1821_v10_).length(0))
                (d_1821_v10_)[index333_] = (p1) - (default__.safeDivisionInt(p1, p1))
                d_1824_v11_: str
                d_1824_v11_ = _dafny.CodePoint('d')
                d_1825_v12_: _dafny.MultiSet
                d_1825_v12_ = _dafny.MultiSet([_dafny.CodePoint('l'), d_1824_v11_, _dafny.CodePoint('q'), _dafny.CodePoint('u')])
                d_1825_v12_ = d_1825_v12_
                r1 = False
                d_1826_v13_: _dafny.Array
                def lambda72_(d_1827_p1_):
                    def lambda73_(d_1828_i3_):
                        return D7_DC15(_dafny.SeqWithoutIsStrInference([d_1827_p1_]))

                    return lambda73_

                init43_ = lambda72_(p1)
                nw270_ = _dafny.Array(None, 24)
                for i0_43_ in range(nw270_.length(0)):
                    nw270_[i0_43_] = init43_(i0_43_)
                d_1826_v13_ = nw270_
                d_1829_v14_: _dafny.Set
                d_1829_v14_ = _dafny.Set({d_1826_v13_, d_1826_v13_, d_1826_v13_})
                index334_ = default__.safeIndex(441, (d_1821_v10_).length(0))
                rhs297_ = (d_1821_v10_)[default__.safeIndex(441, (d_1821_v10_).length(0))]
                rhs298_ = (d_1829_v14_) != (d_1829_v14_)
                lhs240_ = d_1821_v10_
                lhs241_ = default__.safeIndex(441, (d_1821_v10_).length(0))
                lhs240_[lhs241_] = rhs297_
                r1 = rhs298_
                d_1830_v15_: _dafny.Array
                def lambda74_(d_1831_p1_):
                    def lambda75_(d_1832_i4_):
                        return (545) <= (d_1831_p1_)

                    return lambda75_

                init44_ = lambda74_(p1)
                nw271_ = _dafny.Array(None, 3)
                for i0_44_ in range(nw271_.length(0)):
                    nw271_[i0_44_] = init44_(i0_44_)
                d_1830_v15_ = nw271_
                d_1833_v16_: _dafny.Seq
                d_1833_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fxdirlqrk"))
                d_1834_v17_: _dafny.Array
                nw272_ = _dafny.Array(_dafny.Array(None, 0), 15)
                d_1834_v17_ = nw272_
                index335_ = default__.safeIndex(495, (d_1834_v17_).length(0))
                (d_1834_v17_)[index335_] = d_1830_v15_
                index336_ = default__.safeIndex(495, (d_1834_v17_).length(0))
                rhs299_ = (default__.safeDivisionInt(p1, p1)) == (p1)
                rhs300_ = d_1821_v10_
                rhs301_ = d_1830_v15_
                rhs302_ = d_1833_v16_
                rhs303_ = (self).f14
                lhs242_ = d_1834_v17_
                lhs243_ = default__.safeIndex(495, (d_1834_v17_).length(0))
                r1 = rhs299_
                d_1821_v10_ = rhs300_
                d_1830_v15_ = rhs301_
                d_1833_v16_ = rhs302_
                lhs242_[lhs243_] = rhs303_
            elif True:
                r0 = p1
                d_1835_v18_: _dafny.Map
                d_1835_v18_ = _dafny.Map({default__.safeDivisionInt(p1, p1): (p1) + (p1)})
                d_1835_v18_ = d_1835_v18_
                d_1836_v19_: _dafny.MultiSet
                d_1836_v19_ = _dafny.MultiSet([p2, (self).f13])
                d_1837_v20_: D16
                d_1837_v20_ = D16_DC42(d_1836_v19_)
                d_1837_v20_ = d_1837_v20_
                r1 = (self).f13
                d_1838_v21_: _dafny.Seq
                d_1838_v21_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "olfi"))
                d_1839_v22_: _dafny.Seq
                d_1839_v22_ = _dafny.SeqWithoutIsStrInference([d_1838_v21_, d_1838_v21_])
                d_1840_v23_: D6
                d_1840_v23_ = D6_DC12((self).f13, d_1838_v21_, p1, p3)
                d_1841_v24_: _dafny.Array
                nw273_ = _dafny.Array(None, 16)
                nw273_[int(0)] = d_1839_v22_
                nw273_[int(1)] = (_dafny.SeqWithoutIsStrInference([d_1838_v21_])) + (d_1839_v22_)
                nw273_[int(2)] = d_1839_v22_
                nw273_[int(3)] = d_1839_v22_
                nw273_[int(4)] = default__.fm45(d_1840_v23_, p1, globalState)
                nw273_[int(5)] = d_1839_v22_
                nw273_[int(6)] = d_1839_v22_
                nw273_[int(7)] = d_1839_v22_
                nw273_[int(8)] = d_1839_v22_
                nw273_[int(9)] = (d_1839_v22_) + (d_1839_v22_)
                nw273_[int(10)] = d_1839_v22_
                nw273_[int(11)] = d_1839_v22_
                nw273_[int(12)] = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "msjf")) for d_1842_i5_ in range(default__.abs(-12))])
                nw273_[int(13)] = ((d_1839_v22_).set(default__.safeIndex(len(d_1838_v21_), len(d_1839_v22_)), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vqwt")))) + (d_1839_v22_)
                nw273_[int(14)] = (_dafny.SeqWithoutIsStrInference([d_1838_v21_])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qpiiluosn")) for d_1843_i6_ in range(default__.abs(721))]))
                nw273_[int(15)] = default__.fm45(d_1840_v23_, len(_dafny.SeqWithoutIsStrInference([d_1835_v18_, default__.fm24(p1, globalState)])), globalState)
                d_1841_v24_ = nw273_
                index337_ = default__.safeIndex(149, (d_1841_v24_).length(0))
                (d_1841_v24_)[index337_] = d_1839_v22_
                index338_ = default__.safeIndex(149, (d_1841_v24_).length(0))
                (d_1841_v24_)[index338_] = d_1839_v22_
            d_1844_v25_: str
            d_1844_v25_ = _dafny.CodePoint('k')
            d_1845_v26_: C1
            nw274_ = C1()
            nw274_.ctor__(d_1844_v25_, self.f12, (p1) >= ((self).fm9(d_1844_v25_, globalState)))
            d_1845_v26_ = nw274_
            r1 = (self).f17
            d_1846_v27_: _dafny.Map
            d_1846_v27_ = _dafny.Map({p3: default__.fm17(globalState)})
            d_1847_v28_: _dafny.Map
            d_1847_v28_ = _dafny.Map({p1: (self).f17})
            d_1848_v29_: _dafny.Map
            d_1848_v29_ = _dafny.Map({d_1845_v26_.f23: p3})
            d_1849_v30_: D13
            d_1849_v30_ = D13_DC34(p1, default__.fm65(p3, _dafny.SeqWithoutIsStrInference([len(d_1848_v29_)]), d_1844_v25_, not(p2), globalState))
            pat_let_tv50_ = p1
            d_1850_v31_: _dafny.Array
            nw275_ = _dafny.Array(None, 29)
            nw275_[int(0)] = 105
            nw275_[int(1)] = p1
            nw275_[int(2)] = p1
            nw275_[int(3)] = p1
            nw275_[int(4)] = p1
            nw275_[int(5)] = p1
            nw275_[int(6)] = p1
            nw275_[int(7)] = p1
            nw275_[int(8)] = p1
            nw275_[int(9)] = (len(d_1846_v27_)) * (p1)
            nw275_[int(10)] = (0) - ((p1) + (p1))
            nw275_[int(11)] = p1
            nw275_[int(12)] = 916
            nw275_[int(13)] = p1
            nw275_[int(14)] = len(d_1847_v28_)
            nw275_[int(15)] = p1
            nw275_[int(16)] = 99
            nw275_[int(17)] = p1
            nw275_[int(18)] = p1
            nw275_[int(19)] = p1
            nw275_[int(20)] = p1
            nw275_[int(21)] = -859
            nw275_[int(22)] = p1
            nw275_[int(23)] = p1
            nw275_[int(24)] = -842
            nw275_[int(25)] = default__.fm0(d_1809_v0_, p3, not(p3), p1, globalState)
            nw275_[int(26)] = p1
            nw275_[int(27)] = (0) - ((0) - (p1))
            def iife172_(_pat_let59_0):
                def iife173_(d_1851_dt__update__tmp_h0_):
                    def iife174_(_pat_let60_0):
                        def iife175_(d_1852_dt__update_hcf71_h0_):
                            return D13_DC34(d_1852_dt__update_hcf71_h0_, (d_1851_dt__update__tmp_h0_).cf72)
                        return iife175_(_pat_let60_0)
                    return iife174_(pat_let_tv50_)
                return iife173_(_pat_let59_0)
            nw275_[int(28)] = (iife172_(d_1849_v30_)).cf71
            d_1850_v31_ = nw275_
            d_1850_v31_ = d_1850_v31_
        elif True:
            r1 = (self).f17
            r1 = (self).f17
            d_1853_v32_: str
            d_1853_v32_ = _dafny.CodePoint('r')
            d_1854_v33_: _dafny.Map
            d_1854_v33_ = _dafny.Map({(self).fm8((self).f17, self.f16, globalState): d_1853_v32_})
            d_1855_v34_: _dafny.MultiSet
            d_1855_v34_ = _dafny.MultiSet([p1, p1])
            d_1854_v33_ = (d_1854_v33_).set((d_1855_v34_).cardinality, d_1853_v32_)
            d_1856_v35_: _dafny.Seq
            d_1856_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vquk"))
            d_1857_v36_: _dafny.Seq
            d_1857_v36_ = _dafny.SeqWithoutIsStrInference([d_1856_v35_])
            d_1858_v37_: _dafny.Map
            d_1858_v37_ = _dafny.Map({p0: (self).f17})
            d_1859_v38_: _dafny.Array
            nw276_ = _dafny.Array(None, 8)
            nw276_[int(0)] = 466
            nw276_[int(1)] = len(d_1858_v37_)
            nw276_[int(2)] = p1
            nw276_[int(3)] = p1
            nw276_[int(4)] = p1
            nw276_[int(5)] = p1
            nw276_[int(6)] = (0) - (len(d_1856_v35_))
            nw276_[int(7)] = -758
            d_1859_v38_ = nw276_
            d_1860_v39_: _dafny.Map
            d_1860_v39_ = _dafny.Map({(d_1857_v36_)[default__.safeIndex(p1, len(d_1857_v36_))]: _dafny.Map({d_1859_v38_: p1})})
            d_1860_v39_ = (d_1860_v39_).set(_dafny.SeqWithoutIsStrInference([d_1853_v32_ for d_1861_i7_ in range(default__.abs(937))]), _dafny.Map({d_1859_v38_: len(d_1856_v35_)}))
            pat_let_tv51_ = p1
            d_1862_v40_: C2
            nw277_ = C2()
            def iife176_(_pat_let61_0):
                def iife177_(d_1863_dt__update__tmp_h1_):
                    def iife178_(_pat_let62_0):
                        def iife179_(d_1864_dt__update_hcf2_h0_):
                            return D1_DC2(d_1864_dt__update_hcf2_h0_, (d_1863_dt__update__tmp_h1_).cf3, (d_1863_dt__update__tmp_h1_).cf4)
                        return iife179_(_pat_let62_0)
                    return iife178_(pat_let_tv51_)
                return iife177_(_pat_let61_0)
            nw277_.ctor__(default__.safeDivisionInt(p1, 18), (p3) or (not(p0)), self.f16, p2, (self).f14, ((self).f15).intersection((self).f15), iife176_(D1_DC2(p1, p1, not((self).f13))), (self).f13)
            d_1862_v40_ = nw277_
        r1 = (self).f17
        d_1865_v41_: str
        d_1865_v41_ = _dafny.CodePoint('a')
        d_1866_v42_: _dafny.Set
        d_1866_v42_ = _dafny.Set({D3_DC5(d_1865_v41_)})
        d_1867_v43_: _dafny.MultiSet
        d_1867_v43_ = _dafny.MultiSet([p1, p1])
        pat_let_tv52_ = d_1866_v42_
        pat_let_tv53_ = d_1865_v41_
        pat_let_tv54_ = d_1865_v41_
        pat_let_tv55_ = d_1865_v41_
        pat_let_tv56_ = d_1865_v41_
        pat_let_tv57_ = d_1865_v41_
        pat_let_tv58_ = d_1866_v42_
        def lambda76_(source26_):
            if source26_.is_DC12:
                d_1868___mcc_h0_ = source26_.cf18
                d_1869___mcc_h1_ = source26_.cf19
                d_1870___mcc_h2_ = source26_.cf20
                d_1871___mcc_h3_ = source26_.cf21
                d_1872_cf21_ = d_1871___mcc_h3_
                d_1873_cf20_ = d_1870___mcc_h2_
                d_1874_cf19_ = d_1869___mcc_h1_
                d_1875_cf18_ = d_1868___mcc_h0_
                return pat_let_tv52_
            elif source26_.is_DC13:
                d_1876___mcc_h4_ = source26_.cf22
                d_1877___mcc_h5_ = source26_.cf23
                d_1878___mcc_h6_ = source26_.cf24
                d_1879___mcc_h7_ = source26_.cf25
                d_1880___mcc_h8_ = source26_.cf26
                d_1881_cf26_ = d_1880___mcc_h8_
                d_1882_cf25_ = d_1879___mcc_h7_
                d_1883_cf24_ = d_1878___mcc_h6_
                d_1884_cf23_ = d_1877___mcc_h5_
                d_1885_cf22_ = d_1876___mcc_h4_
                return _dafny.Set({D3_DC5(pat_let_tv53_), D3_DC5(pat_let_tv54_)})
            elif source26_.is_DC11:
                d_1886___mcc_h9_ = source26_.cf17
                d_1887_cf17_ = d_1886___mcc_h9_
                return _dafny.Set({D3_DC5(pat_let_tv55_), D3_DC5(pat_let_tv56_), D3_DC5(pat_let_tv57_)})
            elif True:
                d_1888___mcc_h10_ = source26_.cf27
                d_1889_cf27_ = d_1888___mcc_h10_
                return pat_let_tv58_

        rhs304_ = p1
        rhs305_ = 209
        rhs306_ = lambda76_(D6_DC11(d_1867_v43_))
        rhs307_ = p0
        r0 = rhs304_
        r0 = rhs305_
        d_1866_v42_ = rhs306_
        r1 = rhs307_
        r1 = p3
        if (p3) or (not((self).f17)):
            d_1890_v44_: _dafny.Seq
            d_1890_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wypej"))
            d_1891_v45_: _dafny.Array
            nw278_ = _dafny.Array(int(0), 12)
            d_1891_v45_ = nw278_
            index339_ = default__.safeIndex(216, (d_1891_v45_).length(0))
            (d_1891_v45_)[index339_] = (_dafny.MultiSet([(self).f13, p3])).cardinality
            index340_ = default__.safeIndex(216, (d_1891_v45_).length(0))
            rhs308_ = default__.safeDivisionInt((p1 if (self).f17 else 36), p1)
            rhs309_ = ((d_1890_v44_) + (d_1890_v44_)) + (d_1890_v44_)
            rhs310_ = d_1865_v41_
            rhs311_ = (p1) == (p1)
            rhs312_ = len(((_dafny.SeqWithoutIsStrInference([True])) + (d_1809_v0_)) + (d_1809_v0_))
            lhs244_ = d_1891_v45_
            lhs245_ = default__.safeIndex(216, (d_1891_v45_).length(0))
            r0 = rhs308_
            d_1890_v44_ = rhs309_
            d_1865_v41_ = rhs310_
            r1 = rhs311_
            lhs244_[lhs245_] = rhs312_
            d_1892_v46_: C1
            nw279_ = C1()
            nw279_.ctor__(d_1865_v41_, self.f12, p2)
            d_1892_v46_ = nw279_
            d_1893_v47_: _dafny.Seq
            d_1893_v47_ = _dafny.SeqWithoutIsStrInference([d_1892_v46_])
            d_1894_v48_: _dafny.Array
            nw280_ = _dafny.Array(None, 13)
            nw280_[int(0)] = d_1892_v46_
            nw280_[int(1)] = d_1892_v46_
            nw280_[int(2)] = d_1892_v46_
            nw280_[int(3)] = d_1892_v46_
            nw280_[int(4)] = d_1892_v46_
            nw280_[int(5)] = d_1892_v46_
            nw280_[int(6)] = d_1892_v46_
            nw280_[int(7)] = d_1892_v46_
            nw280_[int(8)] = d_1892_v46_
            nw280_[int(9)] = d_1892_v46_
            nw280_[int(10)] = (d_1893_v47_)[default__.safeIndex(p1, len(d_1893_v47_))]
            nw280_[int(11)] = d_1892_v46_
            nw280_[int(12)] = d_1892_v46_
            d_1894_v48_ = nw280_
            d_1895_v49_: D17
            d_1895_v49_ = D17_DC44(d_1894_v48_)
            d_1896_v50_: _dafny.Map
            d_1896_v50_ = _dafny.Map({default__.fm0(_dafny.SeqWithoutIsStrInference([p3]), (self).f13, p3, len(_dafny.SeqWithoutIsStrInference([d_1865_v41_ for d_1897_i8_ in range(default__.abs(527))])), globalState): d_1894_v48_})
            d_1898_v51_: _dafny.Array
            nw281_ = _dafny.Array(None, 14)
            nw281_[int(0)] = d_1894_v48_
            nw281_[int(1)] = d_1894_v48_
            nw281_[int(2)] = d_1894_v48_
            nw281_[int(3)] = d_1894_v48_
            nw281_[int(4)] = d_1894_v48_
            nw281_[int(5)] = (d_1895_v49_).cf87
            nw281_[int(6)] = d_1894_v48_
            nw281_[int(7)] = d_1894_v48_
            nw281_[int(8)] = ((d_1896_v50_)[(d_1891_v45_)[default__.safeIndex(216, (d_1891_v45_).length(0))]] if ((d_1891_v45_)[default__.safeIndex(216, (d_1891_v45_).length(0))]) in (d_1896_v50_) else d_1894_v48_)
            nw281_[int(9)] = d_1894_v48_
            nw281_[int(10)] = d_1894_v48_
            nw281_[int(11)] = d_1894_v48_
            nw281_[int(12)] = d_1894_v48_
            nw281_[int(13)] = d_1894_v48_
            d_1898_v51_ = nw281_
            index341_ = default__.safeIndex(337, (d_1898_v51_).length(0))
            (d_1898_v51_)[index341_] = d_1894_v48_
            d_1899_v52_: _dafny.Seq
            d_1899_v52_ = _dafny.SeqWithoutIsStrInference([(d_1894_v48_ if False else d_1894_v48_), d_1894_v48_, d_1894_v48_, d_1894_v48_, d_1894_v48_])
            index342_ = default__.safeIndex(337, (d_1898_v51_).length(0))
            (d_1898_v51_)[index342_] = (d_1899_v52_)[default__.safeIndex(p1, len(d_1899_v52_))]
            index343_ = default__.safeIndex(216, (d_1891_v45_).length(0))
            (d_1891_v45_)[index343_] = (0) - (p1)
            r1 = (self).f17
            (globalState).f7 = p1
        elif True:
            d_1900_v53_: _dafny.Seq
            d_1900_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uvv"))
            (globalState).f7 = len((d_1900_v53_ if p0 else d_1900_v53_))
            d_1901_v54_: _dafny.Map
            d_1901_v54_ = _dafny.Map({p1: (self).f17})
            d_1902_v55_: C2
            nw282_ = C2()
            nw282_.ctor__(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bndxiw")))), p0, self.f16, ((d_1901_v54_)[p1] if (p1) in (d_1901_v54_) else (self).f13), (self).f14, (self).f15, self.f12, (self).f17)
            d_1902_v55_ = nw282_
            d_1902_v55_ = d_1902_v55_
            d_1903_v56_: _dafny.Map
            d_1903_v56_ = _dafny.Map({d_1902_v55_: (self).f17})
            d_1904_v57_: D6
            d_1904_v57_ = D6_DC12(((d_1902_v55_).f21) != ((d_1902_v55_).f21), (d_1900_v53_) + ((d_1900_v53_).set(default__.safeIndex((d_1902_v55_).f21, len(d_1900_v53_)), d_1865_v41_)), len((d_1903_v56_) | (d_1903_v56_)), p3)
            d_1904_v57_ = default__.fm39(not (p3) or (False), p3, globalState)
            if p0:
                d_1905_v58_: _dafny.Array
                def lambda77_(d_1906_v41_):
                    def lambda78_(d_1907_i9_):
                        return d_1906_v41_

                    return lambda78_

                init45_ = lambda77_(d_1865_v41_)
                nw283_ = _dafny.Array(None, 17)
                for i0_45_ in range(nw283_.length(0)):
                    nw283_[i0_45_] = init45_(i0_45_)
                d_1905_v58_ = nw283_
                d_1908_v59_: D13
                d_1908_v59_ = D13_DC33(d_1905_v58_)
                d_1909_v60_: D13
                d_1909_v60_ = D13_DC35(d_1908_v59_)
                d_1910_v61_: D13
                d_1910_v61_ = D13_DC35((d_1909_v60_).cf73)
                d_1910_v61_ = d_1910_v61_
                index344_ = default__.safeIndex(354, ((self).f14).length(0))
                ((self).f14)[index344_] = (self).f13
                d_1911_v62_: _dafny.MultiSet
                d_1911_v62_ = _dafny.MultiSet([(d_1902_v55_).f22, not((self).f13), not((d_1902_v55_).f22)])
                d_1912_v63_: D6
                d_1912_v63_ = D6_DC13(p2, p2, (d_1902_v55_).f21, (d_1911_v62_).cardinality, not(p0))
                index345_ = default__.safeIndex(354, ((self).f14).length(0))
                rhs313_ = (d_1912_v63_).cf23
                rhs314_ = d_1902_v55_
                rhs315_ = (d_1902_v55_).f22
                lhs246_ = (self).f14
                lhs247_ = default__.safeIndex(354, ((self).f14).length(0))
                r1 = rhs313_
                d_1902_v55_ = rhs314_
                lhs246_[lhs247_] = rhs315_
                d_1913_v64_: _dafny.Array
                def lambda79_(d_1914_v55_):
                    def lambda80_(d_1915_i10_):
                        return (d_1915_i10_) + (len(_dafny.SeqWithoutIsStrInference([(d_1914_v55_).f22])))

                    return lambda80_

                init46_ = lambda79_(d_1902_v55_)
                nw284_ = _dafny.Array(None, 6)
                for i0_46_ in range(nw284_.length(0)):
                    nw284_[i0_46_] = init46_(i0_46_)
                d_1913_v64_ = nw284_
                d_1913_v64_ = d_1913_v64_
                d_1916_v65_: _dafny.Map
                d_1916_v65_ = _dafny.Map({(d_1902_v55_).f21: (d_1902_v55_).f21})
                d_1917_v66_: D11
                d_1917_v66_ = D11_DC27((self).f13, d_1865_v41_, ((d_1916_v65_)[(d_1902_v55_).f21] if ((d_1902_v55_).f21) in (d_1916_v65_) else (0) - ((d_1902_v55_).f21)), (d_1902_v55_).f21, d_1810_v1_)
                d_1917_v66_ = d_1917_v66_
                d_1918_v67_: _dafny.Array
                nw285_ = _dafny.Array(_dafny.Array(None, 0), 10)
                d_1918_v67_ = nw285_
                d_1918_v67_ = d_1918_v67_
            elif True:
                index346_ = default__.safeIndex(678, ((self).f14).length(0))
                ((self).f14)[index346_] = False
                d_1919_v68_: _dafny.Array
                nw286_ = _dafny.Array(False, 7)
                d_1919_v68_ = nw286_
                d_1920_v69_: _dafny.Array
                nw287_ = _dafny.Array(None, 15)
                nw287_[int(0)] = d_1867_v43_
                nw287_[int(1)] = d_1867_v43_
                nw287_[int(2)] = (_dafny.MultiSet(d_1810_v1_)).intersection(d_1867_v43_)
                nw287_[int(3)] = d_1867_v43_
                nw287_[int(4)] = (d_1867_v43_).intersection(default__.fm19(globalState))
                nw287_[int(5)] = (d_1867_v43_) | (d_1867_v43_)
                nw287_[int(6)] = d_1867_v43_
                nw287_[int(7)] = d_1867_v43_
                nw287_[int(8)] = d_1867_v43_
                nw287_[int(9)] = (d_1867_v43_) | (d_1867_v43_)
                nw287_[int(10)] = d_1867_v43_
                nw287_[int(11)] = (d_1867_v43_) | (_dafny.MultiSet(default__.fm27(globalState)))
                nw287_[int(12)] = d_1867_v43_
                nw287_[int(13)] = (d_1867_v43_).set((d_1902_v55_).f21, default__.abs((d_1902_v55_).f21))
                nw287_[int(14)] = (d_1867_v43_).set((0) - (p1), default__.abs((0) - ((d_1902_v55_).f21)))
                d_1920_v69_ = nw287_
                index347_ = default__.safeIndex(544, (d_1920_v69_).length(0))
                (d_1920_v69_)[index347_] = (default__.fm56(p1, (self).f15, globalState)).intersection(d_1867_v43_)
                d_1921_v70_: _dafny.Map
                d_1921_v70_ = _dafny.Map({(d_1902_v55_).f21: (0) - ((_dafny.MultiSet(d_1809_v0_)).cardinality)})
                index348_ = default__.safeIndex(678, ((self).f14).length(0))
                index349_ = default__.safeIndex(544, (d_1920_v69_).length(0))
                rhs316_ = default__.safeDivisionInt(((d_1902_v55_).f21) * ((d_1902_v55_).f21), 721)
                rhs317_ = default__.fm30((len(_dafny.Map({len((d_1921_v70_).set((d_1902_v55_).f21, 217)): p1}))) * ((d_1902_v55_).f21), globalState)
                rhs318_ = (self).f14
                rhs319_ = (d_1867_v43_) | (d_1867_v43_)
                rhs320_ = d_1865_v41_
                lhs248_ = globalState
                lhs249_ = (self).f14
                lhs250_ = default__.safeIndex(678, ((self).f14).length(0))
                lhs251_ = d_1920_v69_
                lhs252_ = default__.safeIndex(544, (d_1920_v69_).length(0))
                lhs248_.f7 = rhs316_
                lhs249_[lhs250_] = rhs317_
                d_1919_v68_ = rhs318_
                lhs251_[lhs252_] = rhs319_
                d_1865_v41_ = rhs320_
                r1 = (p1) > (((d_1902_v55_).f21) * ((0) - ((d_1902_v55_).fm9(default__.fm18((d_1902_v55_).f21, d_1900_v53_, p2, d_1811_v2_, globalState), globalState))))
                r0 = ((0) - ((d_1902_v55_).f21) if p2 else (d_1902_v55_).f21)
                (globalState).f7 = (d_1902_v55_).f21
                d_1922_v71_: _dafny.Array
                def lambda81_(d_1923_v55_):
                    def lambda82_(d_1924_i11_):
                        return default__.safeModuloInt(d_1924_i11_, (d_1923_v55_).f21)

                    return lambda82_

                init47_ = lambda81_(d_1902_v55_)
                nw288_ = _dafny.Array(None, 20)
                for i0_47_ in range(nw288_.length(0)):
                    nw288_[i0_47_] = init47_(i0_47_)
                d_1922_v71_ = nw288_
                index350_ = default__.safeIndex(602, (d_1922_v71_).length(0))
                (d_1922_v71_)[index350_] = p1
                d_1925_v72_: _dafny.Map
                d_1925_v72_ = _dafny.Map({not(((self).f14)[default__.safeIndex(678, ((self).f14).length(0))]): d_1811_v2_})
                index351_ = default__.safeIndex(602, (d_1922_v71_).length(0))
                rhs321_ = default__.safeModuloInt((0) - (default__.safeModuloInt((self).fm9(d_1865_v41_, globalState), (d_1902_v55_).f21)), (p1) * ((d_1902_v55_).f21))
                rhs322_ = d_1919_v68_
                rhs323_ = (d_1902_v55_).f21
                rhs324_ = ((d_1925_v72_ if (self).f13 else d_1925_v72_)).set((d_1902_v55_).f22, (d_1811_v2_).set((self).f17, len(d_1810_v1_)))
                rhs325_ = (d_1902_v55_).f22
                lhs253_ = globalState
                lhs254_ = d_1922_v71_
                lhs255_ = default__.safeIndex(602, (d_1922_v71_).length(0))
                lhs253_.f7 = rhs321_
                d_1919_v68_ = rhs322_
                lhs254_[lhs255_] = rhs323_
                d_1925_v72_ = rhs324_
                r1 = rhs325_
            (globalState).f7 = (default__.safeDivisionInt(((d_1811_v2_)[p2] if (p2) in (d_1811_v2_) else p1), (d_1902_v55_).f21)) - (p1)
        r0 = p1
        r1 = ((default__.fm56(414, (self).f15, globalState)).cardinality) != (756)
        return r0, r1


class C7:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    def ctor__(self):
        pass
        pass

    def fm6(self, p0, globalState):
        if not(not(True)):
            return (0) - (((0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d')]))])))) * (len(_dafny.SeqWithoutIsStrInference([True]))))
        elif True:
            return (len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pkq")))}))) * (959)

    def fm7(self, p0, p1, p2, p3, globalState):
        return ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_1926_i0_ in range(default__.abs(-407))])))) >= (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sgoyyjn")))))

    def m4(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: bool = False
        d_1927_v0_: _dafny.Array
        def lambda83_(d_1928_p3_):
            def lambda84_(d_1929_i0_):
                return default__.safeDivisionInt(d_1929_i0_, d_1928_p3_)

            return lambda84_

        init48_ = lambda83_(p3)
        nw289_ = _dafny.Array(None, 2)
        for i0_48_ in range(nw289_.length(0)):
            nw289_[i0_48_] = init48_(i0_48_)
        d_1927_v0_ = nw289_
        d_1930_v1_: _dafny.Map
        d_1930_v1_ = _dafny.Map({d_1927_v0_: p0})
        d_1930_v1_ = (d_1930_v1_).set(d_1927_v0_, p0)
        d_1931_i1_: int
        d_1931_i1_ = 0
        with _dafny.label("13"):
            while not(p0):
                with _dafny.c_label("13"):
                    if (d_1931_i1_) >= (100):
                        raise _dafny.Break("13")
                    d_1931_i1_ = (d_1931_i1_) + (1)
                    d_1932_v2_: str
                    d_1932_v2_ = _dafny.CodePoint('j')
                    d_1933_v3_: _dafny.Seq
                    d_1933_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yqxdb"))
                    d_1932_v2_ = (d_1933_v3_)[default__.safeIndex(p2, len(d_1933_v3_))]
                    d_1934_v4_: _dafny.Array
                    nw290_ = _dafny.Array(False, 3)
                    d_1934_v4_ = nw290_
                    d_1935_v5_: _dafny.Map
                    d_1935_v5_ = _dafny.Map({p0: True})
                    d_1936_v6_: _dafny.Set
                    d_1936_v6_ = _dafny.Set({p0, p0})
                    index352_ = default__.safeIndex(870, (d_1934_v4_).length(0))
                    (d_1934_v4_)[index352_] = (self).fm7(p2, d_1935_v5_, d_1933_v3_, d_1936_v6_, globalState)
                    d_1937_v7_: D2
                    d_1937_v7_ = D2_DC4(d_1927_v0_, p0, p3, p0, d_1934_v4_)
                    index353_ = default__.safeIndex(870, (d_1934_v4_).length(0))
                    (d_1934_v4_)[index353_] = (d_1937_v7_).cf9
                    d_1933_v3_ = d_1933_v3_
                    index354_ = default__.safeIndex(80, (d_1927_v0_).length(0))
                    (d_1927_v0_)[index354_] = p3
                    index355_ = default__.safeIndex(80, (d_1927_v0_).length(0))
                    (d_1927_v0_)[index355_] = default__.safeModuloInt(default__.safeModuloInt(901, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vtl")))), p3)
                    pass
            pass
        d_1938_v8_: _dafny.Array
        nw291_ = _dafny.Array(False, 14)
        d_1938_v8_ = nw291_
        d_1939_v9_: D2
        d_1939_v9_ = D2_DC4(d_1927_v0_, not(p0), p3, p0, d_1938_v8_)
        source27_ = d_1939_v9_
        if source27_.is_DC4:
            d_1940___mcc_h0_ = source27_.cf6
            d_1941___mcc_h1_ = source27_.cf7
            d_1942___mcc_h2_ = source27_.cf8
            d_1943___mcc_h3_ = source27_.cf9
            d_1944___mcc_h4_ = source27_.cf10
            d_1945_cf10_ = d_1944___mcc_h4_
            d_1946_cf9_ = d_1943___mcc_h3_
            d_1947_cf8_ = d_1942___mcc_h2_
            d_1948_cf7_ = d_1941___mcc_h1_
            d_1949_cf6_ = d_1940___mcc_h0_
            d_1947_cf8_ = p3
            d_1950_v10_: str
            d_1950_v10_ = _dafny.CodePoint('o')
            d_1951_v11_: _dafny.MultiSet
            out14_: _dafny.MultiSet
            out14_ = default__.m0(d_1950_v10_, p0, globalState)
            d_1951_v11_ = out14_
            d_1952_v12_: _dafny.Seq
            d_1952_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))
            d_1953_v13_: D1
            d_1953_v13_ = D1_DC1(d_1952_v12_)
            d_1954_v14_: _dafny.Set
            d_1954_v14_ = _dafny.Set({p2})
            d_1955_v15_: _dafny.Map
            d_1955_v15_ = _dafny.Map({d_1948_cf7_: p3})
            d_1956_v16_: D1
            d_1956_v16_ = D1_DC2(len(d_1955_v15_), d_1947_cf8_, False)
            pat_let_tv59_ = d_1952_v12_
            d_1957_v17_: C6
            nw292_ = C6()
            def iife180_(_pat_let63_0):
                def iife181_(d_1958_dt__update__tmp_h0_):
                    def iife182_(_pat_let64_0):
                        def iife183_(d_1959_dt__update_hcf1_h0_):
                            return D1_DC1(d_1959_dt__update_hcf1_h0_)
                        return iife183_(_pat_let64_0)
                    return iife182_(pat_let_tv59_)
                return iife181_(_pat_let63_0)
            nw292_.ctor__(iife180_(d_1953_v13_), d_1946_cf9_, d_1938_v8_, d_1954_v14_, d_1956_v16_, True)
            d_1957_v17_ = nw292_
            index356_ = default__.safeIndex(95, (d_1949_cf6_).length(0))
            (d_1949_cf6_)[index356_] = d_1947_cf8_
            d_1960_v18_: _dafny.Seq
            d_1960_v18_ = _dafny.SeqWithoutIsStrInference([p1])
            index357_ = default__.safeIndex(95, (d_1949_cf6_).length(0))
            (d_1949_cf6_)[index357_] = default__.fm0((d_1960_v18_)[default__.safeIndex(395, len(d_1960_v18_))], d_1946_cf9_, d_1948_cf7_, (0) - (p2), globalState)
        elif True:
            d_1961___mcc_h5_ = source27_.cf5
            d_1962_cf5_ = d_1961___mcc_h5_
            d_1963_v19_: _dafny.Seq
            d_1963_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pliy"))
            d_1964_v20_: _dafny.MultiSet
            d_1964_v20_ = _dafny.MultiSet([d_1963_v19_, d_1963_v19_])
            d_1965_v21_: _dafny.MultiSet
            d_1965_v21_ = _dafny.MultiSet([p0])
            d_1966_v22_: _dafny.Map
            d_1966_v22_ = _dafny.Map({d_1964_v20_: d_1965_v21_})
            d_1967_v23_: str
            d_1967_v23_ = _dafny.CodePoint('h')
            d_1968_v24_: _dafny.Map
            d_1968_v24_ = _dafny.Map({d_1967_v23_: True})
            d_1969_v25_: _dafny.Map
            d_1969_v25_ = _dafny.Map({len(d_1968_v24_): p0})
            d_1970_v26_: _dafny.Set
            d_1970_v26_ = _dafny.Set({len(d_1969_v25_), p3})
            d_1971_v27_: T2
            nw293_ = C5()
            nw293_.ctor__(p0, d_1962_cf5_, d_1970_v26_, D1_DC2(-499, len(_dafny.Set({d_1967_v23_})), p0), p0)
            d_1971_v27_ = nw293_
            d_1972_v28_: _dafny.Map
            d_1972_v28_ = _dafny.Map({p0: d_1971_v27_})
            d_1966_v22_ = (d_1966_v22_).set(default__.fm66(p3, globalState), (d_1965_v21_).set(p0, default__.abs(len(d_1972_v28_))))
            d_1973_v29_: _dafny.Array
            nw294_ = _dafny.Array(None, 13)
            d_1973_v29_ = nw294_
            d_1974_v30_: D17
            d_1974_v30_ = D17_DC44(d_1973_v29_)
            source28_ = d_1974_v30_
            if source28_.is_DC45:
                d_1975___mcc_h6_ = source28_.cf88
                d_1976_cf88_ = d_1975___mcc_h6_
                d_1977_v31_: _dafny.Map
                d_1977_v31_ = _dafny.Map({d_1976_cf88_: p3})
                d_1978_v32_: _dafny.Seq
                d_1978_v32_ = _dafny.SeqWithoutIsStrInference([p3])
                d_1979_v33_: _dafny.Set
                d_1979_v33_ = _dafny.Set({d_1978_v32_})
                d_1980_v34_: _dafny.Seq
                d_1980_v34_ = _dafny.SeqWithoutIsStrInference([d_1979_v33_, d_1979_v33_, d_1979_v33_])
                rhs326_ = (self).fm6(d_1963_v19_, globalState)
                rhs327_ = len((d_1963_v19_ if (default__.fm0(p1, (d_1971_v27_).f13, p0, ((d_1965_v21_)[not((d_1971_v27_).f13)] if (not((d_1971_v27_).f13)) in (d_1965_v21_) else p2), globalState)) == (len(d_1977_v31_)) else d_1963_v19_))
                rhs328_ = ((d_1979_v33_).intersection(d_1979_v33_)).issubset(((d_1980_v34_)[default__.safeIndex(p2, len(d_1980_v34_))] if d_1976_cf88_ else d_1979_v33_))
                lhs256_ = globalState
                lhs257_ = globalState
                lhs256_.f7 = rhs326_
                lhs257_.f7 = rhs327_
                r2 = rhs328_
                r2 = p0
                index358_ = default__.safeIndex(197, ((d_1971_v27_).f14).length(0))
                ((d_1971_v27_).f14)[index358_] = (d_1971_v27_).f13
                index359_ = default__.safeIndex(197, ((d_1971_v27_).f14).length(0))
                ((d_1971_v27_).f14)[index359_] = d_1976_cf88_
                d_1981_v35_: _dafny.Seq
                d_1981_v35_ = _dafny.SeqWithoutIsStrInference([d_1963_v19_])
                d_1982_v36_: _dafny.Map
                d_1982_v36_ = _dafny.Map({D7_DC16(p3): d_1981_v35_})
                d_1982_v36_ = (d_1982_v36_).set(D7_DC16(p3), (d_1981_v35_) + (_dafny.SeqWithoutIsStrInference([d_1963_v19_])))
            elif True:
                d_1983___mcc_h7_ = source28_.cf87
                d_1984_cf87_ = d_1983___mcc_h7_
                d_1985_v37_: _dafny.MultiSet
                d_1985_v37_ = _dafny.MultiSet([d_1962_cf5_])
                d_1986_v38_: _dafny.MultiSet
                d_1986_v38_ = d_1985_v37_
                d_1987_v39_: _dafny.Seq
                d_1987_v39_ = _dafny.SeqWithoutIsStrInference([d_1986_v38_, d_1986_v38_])
                d_1988_v40_: _dafny.Set
                d_1988_v40_ = _dafny.Set({d_1963_v19_})
                d_1989_v41_: _dafny.Set
                d_1989_v41_ = _dafny.Set({((0) - (len(d_1987_v39_))) < (488), (len(d_1988_v40_)) > (p2)})
                d_1989_v41_ = d_1989_v41_
                d_1990_v42_: _dafny.MultiSet
                d_1990_v42_ = _dafny.MultiSet([p2])
                d_1991_v43_: C0
                nw295_ = C0()
                nw295_.ctor__(True, d_1967_v23_, d_1971_v27_.f12, (d_1990_v42_).ispropersubset(d_1990_v42_))
                d_1991_v43_ = nw295_
                d_1992_v44_: _dafny.Seq
                d_1992_v44_ = _dafny.SeqWithoutIsStrInference([d_1927_v0_])
                d_1992_v44_ = (d_1992_v44_) + (d_1992_v44_)
                d_1993_v45_: _dafny.Map
                d_1993_v45_ = _dafny.Map({d_1991_v43_.f19: (d_1971_v27_).f13})
                d_1993_v45_ = (d_1993_v45_).set(p0, (d_1971_v27_).f13)
            r2 = (d_1971_v27_).f13
            d_1994_v46_: _dafny.Array
            nw296_ = _dafny.Array(_dafny.Array(None, 0), 15)
            d_1994_v46_ = nw296_
            index360_ = default__.safeIndex(107, (d_1994_v46_).length(0))
            (d_1994_v46_)[index360_] = d_1938_v8_
            index361_ = default__.safeIndex(107, (d_1994_v46_).length(0))
            rhs329_ = d_1967_v23_
            rhs330_ = d_1938_v8_
            lhs258_ = d_1994_v46_
            lhs259_ = default__.safeIndex(107, (d_1994_v46_).length(0))
            d_1967_v23_ = rhs329_
            lhs258_[lhs259_] = rhs330_
        d_1995_v47_: _dafny.Array
        nw297_ = _dafny.Array(_dafny.Map({}), 24)
        d_1995_v47_ = nw297_
        guard_loop_8_: int
        for guard_loop_8_ in _dafny.IntegerRange(0, (d_1995_v47_).length(0)):
            d_1996_i2_: int = guard_loop_8_
            if (True) and (((0) <= (d_1996_i2_)) and ((d_1996_i2_) < ((d_1995_v47_).length(0)))):
                (d_1995_v47_)[(d_1996_i2_)] = ((_dafny.Map({p0: True})) | (_dafny.Map({p0: p0}))) | ((_dafny.Map({False: p0})) | (_dafny.Map({p0: p0})))
        d_1997_i3_: int
        d_1997_i3_ = 0
        with _dafny.label("14"):
            while (p3) <= (p3):
                with _dafny.c_label("14"):
                    if (d_1997_i3_) >= (100):
                        raise _dafny.Break("14")
                    d_1997_i3_ = (d_1997_i3_) + (1)
                    d_1998_v48_: _dafny.MultiSet
                    d_1998_v48_ = _dafny.MultiSet([p3])
                    d_1998_v48_ = d_1998_v48_
                    d_1999_v49_: _dafny.Seq
                    d_1999_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lvqcjon"))
                    d_2000_v50_: str
                    d_2000_v50_ = _dafny.CodePoint('w')
                    r1 = (521) * (len(((d_1999_v49_).set(default__.safeIndex(p2, len(d_1999_v49_)), d_2000_v50_)) + (d_1999_v49_)))
                    (globalState).f7 = (0) - (p2)
                    d_2001_v51_: _dafny.Map
                    d_2001_v51_ = _dafny.Map({default__.safeDivisionInt(p3, p3): not(True)})
                    r2 = ((d_2001_v51_)[p2] if (p2) in (d_2001_v51_) else p0)
                    pass
            pass
        d_2002_v52_: _dafny.Seq
        d_2002_v52_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rxfde"))
        hi13_ = len(d_2002_v52_)
        for d_2003_i4_ in range((p3) * (325), hi13_):
            r0 = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_2004_i5_ in range(default__.abs(2))]))
            d_2005_v53_: D15
            d_2005_v53_ = D15_DC39(p3, _dafny.CodePoint('h'), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uny")), p3, 80)
            d_2006_v54_: _dafny.MultiSet
            out15_: _dafny.MultiSet
            out15_ = default__.m0((d_2005_v53_).cf79, not(p0), globalState)
            d_2006_v54_ = out15_
            (globalState).f2 = p1
            r0 = (0) - (default__.safeDivisionInt(p2, d_2003_i4_))
        d_2007_v55_: _dafny.Set
        d_2007_v55_ = _dafny.Set({p0})
        r0 = ((p3) - (len(d_2007_v55_))) * (p3)
        r1 = p2
        r2 = p0
        return r0, r1, r2


class C8:
    def  __init__(self):
        self.f11: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    def ctor__(self, f11):
        (self).f11 = f11

    def fm4(self, p0, p1, p2, p3, globalState):
        return self.f11

    def m2(self, p0, p1, p2, globalState):
        r0: bool = False
        d_2008_v0_: _dafny.Seq
        d_2008_v0_ = _dafny.SeqWithoutIsStrInference([p2, self.f11])
        hi14_ = p1
        for d_2009_i0_ in range((default__.fm0(d_2008_v0_, self.f11, self.f11, (0) - (p0), globalState)) * (-964), hi14_):
            (self).f11 = self.f11
            d_2010_v1_: _dafny.MultiSet
            d_2010_v1_ = _dafny.MultiSet([p1])
            d_2011_v2_: _dafny.Seq
            d_2011_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nmecnbdo"))
            d_2012_v3_: _dafny.Map
            d_2012_v3_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(self).fm4(d_2008_v0_, (d_2010_v1_).cardinality, d_2011_v2_, self.f11, globalState), p2]): p1})
            d_2012_v3_ = d_2012_v3_
            (globalState).f7 = p0
            if self.f11:
                (globalState).f7 = p1
                d_2013_v4_: _dafny.Array
                nw298_ = _dafny.Array(int(0), 10)
                d_2013_v4_ = nw298_
                d_2013_v4_ = d_2013_v4_
                d_2014_v5_: _dafny.Map
                d_2014_v5_ = _dafny.Map({(d_2011_v2_) != (d_2011_v2_): default__.safeModuloInt(len(d_2008_v0_), (0) - (d_2009_i0_))})
                d_2014_v5_ = (d_2014_v5_).set(p2, (p1) + (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "npixpjow"))).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "npixpjow")))), default__.fm5(globalState)))))
                r0 = p2
                d_2015_v6_: _dafny.Array
                nw299_ = _dafny.Array(_dafny.Seq({}), 8)
                d_2015_v6_ = nw299_
                d_2016_v7_: _dafny.Seq
                d_2016_v7_ = _dafny.SeqWithoutIsStrInference([d_2008_v0_])
                index362_ = default__.safeIndex(433, (d_2015_v6_).length(0))
                (d_2015_v6_)[index362_] = d_2016_v7_
                d_2017_v8_: _dafny.Seq
                d_2017_v8_ = _dafny.SeqWithoutIsStrInference([p0])
                index363_ = default__.safeIndex(433, (d_2015_v6_).length(0))
                (d_2015_v6_)[index363_] = (d_2016_v7_) + (((_dafny.SeqWithoutIsStrInference([d_2008_v0_])) + (d_2016_v7_)).set(default__.safeIndex((d_2017_v8_)[default__.safeIndex(279, len(d_2017_v8_))], len((_dafny.SeqWithoutIsStrInference([d_2008_v0_])) + (d_2016_v7_))), d_2008_v0_))
            elif True:
                d_2018_v9_: _dafny.Map
                d_2018_v9_ = _dafny.Map({p2: d_2009_i0_})
                d_2019_v10_: D6
                d_2019_v10_ = D6_DC12(self.f11, d_2011_v2_, default__.fm0(d_2008_v0_, self.f11, self.f11, p0, globalState), self.f11)
                d_2020_v11_: D1
                d_2020_v11_ = D1_DC1(d_2011_v2_)
                d_2021_v12_: _dafny.Array
                nw300_ = _dafny.Array(None, 17)
                nw300_[int(0)] = p2
                nw300_[int(1)] = p2
                nw300_[int(2)] = True
                nw300_[int(3)] = self.f11
                nw300_[int(4)] = self.f11
                nw300_[int(5)] = self.f11
                nw300_[int(6)] = self.f11
                nw300_[int(7)] = not(self.f11)
                nw300_[int(8)] = p2
                nw300_[int(9)] = p2
                nw300_[int(10)] = self.f11
                nw300_[int(11)] = p2
                nw300_[int(12)] = p2
                nw300_[int(13)] = self.f11
                nw300_[int(14)] = p2
                nw300_[int(15)] = not(True)
                nw300_[int(16)] = p2
                d_2021_v12_ = nw300_
                d_2022_v13_: _dafny.Set
                d_2022_v13_ = _dafny.Set({p0})
                pat_let_tv60_ = p2
                d_2023_v14_: C2
                nw301_ = C2()
                def iife184_(_pat_let65_0):
                    def iife185_(d_2024_dt__update__tmp_h0_):
                        def iife186_(_pat_let66_0):
                            def iife187_(d_2025_dt__update_hcf4_h0_):
                                def iife188_(_pat_let67_0):
                                    def iife189_(d_2026_dt__update_hcf3_h0_):
                                        return D1_DC2((d_2024_dt__update__tmp_h0_).cf2, d_2026_dt__update_hcf3_h0_, d_2025_dt__update_hcf4_h0_)
                                    return iife189_(_pat_let67_0)
                                return iife188_(d_2009_i0_)
                            return iife187_(_pat_let66_0)
                        return iife186_(pat_let_tv60_)
                    return iife185_(_pat_let65_0)
                nw301_.ctor__((0) - ((len(_dafny.Set({d_2011_v2_}))) * (len(d_2018_v9_))), (self.f11 if (d_2019_v10_).cf21 else self.f11), d_2020_v11_, default__.fm30(p1, globalState), d_2021_v12_, (d_2022_v13_) - (_dafny.Set({p1})), iife184_(default__.fm37(p1, p0, len(d_2011_v2_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "knl"))), globalState)), p2)
                d_2023_v14_ = nw301_
                (self).f11 = not(self.f11)
                d_2021_v12_ = d_2021_v12_
                (globalState).f7 = default__.safeDivisionInt(650, (p1) + ((d_2023_v14_).f21))
                d_2027_v15_: _dafny.Set
                d_2027_v15_ = _dafny.Set({p2})
                d_2028_v16_: C2
                nw302_ = C2()
                nw302_.ctor__((0) - (len((d_2011_v2_ if p2 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "le"))))), (self).fm4(d_2008_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ajkrilfy"))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dli")), not((d_2023_v14_).f22), globalState), D1_DC1(d_2011_v2_), not(default__.fm17(globalState)), d_2021_v12_, (d_2022_v13_).intersection(d_2022_v13_), D1_DC2(default__.fm0(_dafny.SeqWithoutIsStrInference([self.f11, (d_2023_v14_).f22]), p2, self.f11, len(d_2027_v15_), globalState), (d_2023_v14_).f21, (d_2023_v14_).f22), (d_2023_v14_).f22)
                d_2028_v16_ = nw302_
        d_2029_v17_: _dafny.MultiSet
        d_2029_v17_ = _dafny.MultiSet([_dafny.MultiSet([self.f11, p2, self.f11]), _dafny.MultiSet([self.f11, self.f11])])
        d_2030_v18_: _dafny.MultiSet
        d_2030_v18_ = _dafny.MultiSet([p2])
        d_2031_v19_: _dafny.Set
        d_2031_v19_ = _dafny.Set({p1, p0})
        d_2032_v20_: _dafny.Seq
        d_2032_v20_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([((d_2029_v17_)[d_2030_v18_] if (d_2030_v18_) in (d_2029_v17_) else p0), default__.fm0(d_2008_v0_, self.f11, p2, len(d_2031_v19_), globalState), p0, p1, p0])])
        d_2033_v21_: _dafny.Seq
        d_2033_v21_ = _dafny.SeqWithoutIsStrInference([p0, len(d_2008_v0_)])
        d_2034_v22_: str
        d_2034_v22_ = _dafny.CodePoint('a')
        d_2035_v23_: _dafny.Seq
        d_2035_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vrdiples"))
        d_2036_v24_: D1
        d_2036_v24_ = D1_DC2(len(d_2035_v23_), p1, self.f11)
        d_2037_v25_: C1
        nw303_ = C1()
        nw303_.ctor__(d_2034_v22_, d_2036_v24_, self.f11)
        d_2037_v25_ = nw303_
        d_2038_v26_: _dafny.MultiSet
        d_2038_v26_ = _dafny.MultiSet([d_2037_v25_, d_2037_v25_])
        d_2039_v27_: _dafny.Array
        nw304_ = _dafny.Array(None, 25)
        nw304_[int(0)] = 791
        nw304_[int(1)] = p1
        nw304_[int(2)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w")))
        nw304_[int(3)] = len(_dafny.SeqWithoutIsStrInference([p0]))
        nw304_[int(4)] = len(d_2032_v20_)
        nw304_[int(5)] = p1
        nw304_[int(6)] = p0
        nw304_[int(7)] = (0) - (p1)
        nw304_[int(8)] = p1
        nw304_[int(9)] = 341
        nw304_[int(10)] = p0
        nw304_[int(11)] = p1
        nw304_[int(12)] = p1
        nw304_[int(13)] = p1
        nw304_[int(14)] = p0
        nw304_[int(15)] = len(_dafny.SeqWithoutIsStrInference([(0) - (p1), p1]))
        nw304_[int(16)] = default__.fm0(d_2008_v0_, p2, self.f11, p0, globalState)
        nw304_[int(17)] = p1
        nw304_[int(18)] = p0
        nw304_[int(19)] = p1
        nw304_[int(20)] = p1
        nw304_[int(21)] = p1
        nw304_[int(22)] = (d_2033_v21_)[default__.safeIndex((d_2038_v26_).cardinality, len(d_2033_v21_))]
        nw304_[int(23)] = p1
        nw304_[int(24)] = p0
        d_2039_v27_ = nw304_
        d_2040_v28_: _dafny.Map
        d_2040_v28_ = _dafny.Map({p0: d_2039_v27_})
        hi15_ = len((d_2033_v21_) + (d_2033_v21_))
        for d_2041_i1_ in range(len((d_2040_v28_ if self.f11 else _dafny.Map({((d_2030_v18_)[p2] if (p2) in (d_2030_v18_) else p1): d_2039_v27_}))), hi15_):
            (globalState).f7 = d_2041_i1_
            (globalState).f7 = p1
            d_2042_v29_: _dafny.MultiSet
            d_2042_v29_ = _dafny.MultiSet([d_2034_v22_, d_2034_v22_, (d_2037_v25_.f23 if self.f11 else d_2037_v25_.f23), _dafny.CodePoint('f')])
            d_2043_v30_: _dafny.Seq
            d_2043_v30_ = _dafny.SeqWithoutIsStrInference([d_2042_v29_])
            d_2042_v29_ = (d_2043_v30_)[default__.safeIndex((0) - (default__.safeDivisionInt((0) - ((0) - (len(_dafny.Set({self.f11})))), 980)), len(d_2043_v30_))]
            d_2044_v31_: _dafny.MultiSet
            d_2044_v31_ = _dafny.MultiSet([-531])
            (self).f11 = (self.f11) == ((d_2044_v31_) == (d_2044_v31_))
        d_2045_v32_: _dafny.Set
        d_2045_v32_ = _dafny.Set({self.f11, self.f11})
        d_2046_v33_: _dafny.Map
        d_2046_v33_ = _dafny.Map({d_2045_v32_: (d_2033_v21_) + (d_2033_v21_)})
        d_2046_v33_ = (d_2046_v33_).set(d_2045_v32_, _dafny.SeqWithoutIsStrInference([len(_dafny.Set({p0})), p0, 128, (d_2033_v21_)[default__.safeIndex(404, len(d_2033_v21_))], p0]))
        (self).f11 = False
        d_2047_v34_: _dafny.Map
        d_2047_v34_ = _dafny.Map({not(self.f11): p2})
        d_2048_v35_: _dafny.Seq
        d_2048_v35_ = _dafny.SeqWithoutIsStrInference([d_2047_v34_, d_2047_v34_, d_2047_v34_])
        d_2049_v36_: _dafny.Map
        d_2049_v36_ = _dafny.Map({p1: d_2048_v35_})
        d_2048_v35_ = ((d_2049_v36_)[p1] if (p1) in (d_2049_v36_) else (_dafny.SeqWithoutIsStrInference([d_2047_v34_, d_2047_v34_])).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference([d_2047_v34_, d_2047_v34_]))), _dafny.Map({p2: self.f11})))
        d_2050_v37_: D1
        d_2050_v37_ = D1_DC1(_dafny.SeqWithoutIsStrInference([d_2034_v22_ for d_2051_i2_ in range(default__.abs(715))]))
        d_2052_v38_: _dafny.Array
        def lambda85_(d_2053_p2_):
            def lambda86_(d_2054_i3_):
                return d_2053_p2_

            return lambda86_

        init49_ = lambda85_(p2)
        nw305_ = _dafny.Array(None, 29)
        for i0_49_ in range(nw305_.length(0)):
            nw305_[i0_49_] = init49_(i0_49_)
        d_2052_v38_ = nw305_
        d_2055_v39_: C3
        nw306_ = C3()
        nw306_.ctor__(True, default__.fm62(not(False), d_2035_v23_, globalState), p2, d_2052_v38_, d_2031_v19_, d_2036_v24_, self.f11)
        d_2055_v39_ = nw306_
        d_2056_v40_: _dafny.MultiSet
        d_2056_v40_ = _dafny.MultiSet([d_2055_v39_])
        d_2057_v41_: _dafny.Seq
        d_2057_v41_ = _dafny.SeqWithoutIsStrInference([d_2055_v39_])
        d_2058_v42_: _dafny.Set
        d_2058_v42_ = _dafny.Set({_dafny.MultiSet(d_2057_v41_), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_2055_v39_]))})
        d_2059_v43_: C2
        nw307_ = C2()
        nw307_.ctor__(p1, p2, d_2050_v37_, (d_2056_v40_) in (d_2058_v42_), d_2052_v38_, _dafny.Set({p0}), d_2036_v24_, (d_2055_v39_).f18)
        d_2059_v43_ = nw307_
        r0 = (d_2059_v43_).f22
        return r0

    def m3(self, p0, globalState):
        r0: _dafny.Map = _dafny.Map({})
        r1: D2 = D2.default()()
        if True:
            (globalState).f7 = 56
            d_2060_v0_: _dafny.Seq
            d_2060_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yhesaykny"))
            d_2061_v1_: _dafny.Array
            def lambda87_(d_2062_i0_):
                return _dafny.Map({_dafny.Set({self.f11, self.f11, self.f11, self.f11}): self.f11})

            init50_ = lambda87_
            nw308_ = _dafny.Array(None, 23)
            for i0_50_ in range(nw308_.length(0)):
                nw308_[i0_50_] = init50_(i0_50_)
            d_2061_v1_ = nw308_
            d_2063_v2_: _dafny.Set
            d_2063_v2_ = _dafny.Set({self.f11, not(self.f11), not(self.f11)})
            d_2064_v3_: _dafny.Map
            d_2064_v3_ = _dafny.Map({d_2063_v2_: self.f11})
            index364_ = default__.safeIndex(668, (d_2061_v1_).length(0))
            (d_2061_v1_)[index364_] = d_2064_v3_
            d_2065_v5_: _dafny.Map
            d_2065_v5_ = _dafny.Map({d_2063_v2_: 153})
            index365_ = default__.safeIndex(668, (d_2061_v1_).length(0))
            def iife190_():
                coll54_ = _dafny.Map()
                compr_54_: _dafny.Set
                for compr_54_ in (d_2065_v5_).keys.Elements:
                    d_2066_v4_: _dafny.Set = compr_54_
                    if (d_2066_v4_) in (d_2065_v5_):
                        coll54_[d_2066_v4_] = not(self.f11)
                return _dafny.Map(coll54_)
            rhs331_ = self.f11
            rhs332_ = d_2060_v0_
            rhs333_ = ((_dafny.Map({d_2063_v2_: self.f11})) | (iife190_()
            )) | ((d_2064_v3_).set(d_2063_v2_, self.f11))
            rhs334_ = True
            rhs335_ = False
            lhs260_ = self
            lhs261_ = d_2061_v1_
            lhs262_ = default__.safeIndex(668, (d_2061_v1_).length(0))
            lhs263_ = self
            lhs264_ = self
            lhs260_.f11 = rhs331_
            d_2060_v0_ = rhs332_
            lhs261_[lhs262_] = rhs333_
            lhs263_.f11 = rhs334_
            lhs264_.f11 = rhs335_
            d_2067_v6_: _dafny.Seq
            d_2067_v6_ = _dafny.SeqWithoutIsStrInference([self.f11])
            d_2068_v7_: _dafny.Map
            d_2068_v7_ = _dafny.Map({(d_2067_v6_)[default__.safeIndex(p0, len(d_2067_v6_))]: self.f11})
            d_2069_v8_: _dafny.Map
            d_2069_v8_ = _dafny.Map({self.f11: p0})
            d_2070_v9_: _dafny.Map
            d_2070_v9_ = _dafny.Map({default__.fm18(len(d_2068_v7_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "narktrxe")), self.f11, d_2069_v8_, globalState): self.f11})
            d_2071_v10_: str
            d_2071_v10_ = _dafny.CodePoint('k')
            d_2070_v9_ = (_dafny.Map({d_2071_v10_: self.f11})) | (d_2070_v9_)
            d_2072_v11_: D6
            d_2072_v11_ = D6_DC12(self.f11, ((d_2060_v0_) + (d_2060_v0_)).set(default__.safeIndex(p0, len((d_2060_v0_) + (d_2060_v0_))), d_2071_v10_), p0, self.f11)
            source29_ = d_2072_v11_
            if source29_.is_DC12:
                d_2073___mcc_h0_ = source29_.cf18
                d_2074___mcc_h1_ = source29_.cf19
                d_2075___mcc_h2_ = source29_.cf20
                d_2076___mcc_h3_ = source29_.cf21
                d_2077_cf21_ = d_2076___mcc_h3_
                d_2078_cf20_ = d_2075___mcc_h2_
                d_2079_cf19_ = d_2074___mcc_h1_
                d_2080_cf18_ = d_2073___mcc_h0_
                d_2081_v12_: _dafny.Array
                def lambda88_(d_2082_p0_):
                    def lambda89_(d_2083_i1_):
                        return default__.safeModuloInt(d_2083_i1_, (0) - (d_2082_p0_))

                    return lambda89_

                init51_ = lambda88_(p0)
                nw309_ = _dafny.Array(None, 29)
                for i0_51_ in range(nw309_.length(0)):
                    nw309_[i0_51_] = init51_(i0_51_)
                d_2081_v12_ = nw309_
                d_2081_v12_ = d_2081_v12_
                index366_ = default__.safeIndex(613, (d_2081_v12_).length(0))
                (d_2081_v12_)[index366_] = d_2078_cf20_
                index367_ = default__.safeIndex(613, (d_2081_v12_).length(0))
                (d_2081_v12_)[index367_] = (0) - (default__.safeDivisionInt(len(d_2060_v0_), d_2078_cf20_))
                d_2084_v13_: _dafny.Array
                nw310_ = _dafny.Array(_dafny.Set({}), 12)
                d_2084_v13_ = nw310_
                nw311_ = _dafny.Array(_dafny.Set({}), 14)
                d_2084_v13_ = nw311_
                d_2085_v14_: _dafny.Array
                nw312_ = _dafny.Array(_dafny.Array(None, 0), 24)
                d_2085_v14_ = nw312_
                d_2085_v14_ = d_2085_v14_
            elif source29_.is_DC13:
                d_2086___mcc_h4_ = source29_.cf22
                d_2087___mcc_h5_ = source29_.cf23
                d_2088___mcc_h6_ = source29_.cf24
                d_2089___mcc_h7_ = source29_.cf25
                d_2090___mcc_h8_ = source29_.cf26
                d_2091_cf26_ = d_2090___mcc_h8_
                d_2092_cf25_ = d_2089___mcc_h7_
                d_2093_cf24_ = d_2088___mcc_h6_
                d_2094_cf23_ = d_2087___mcc_h5_
                d_2095_cf22_ = d_2086___mcc_h4_
                d_2096_v15_: _dafny.MultiSet
                d_2096_v15_ = _dafny.MultiSet([default__.safeDivisionInt(d_2092_cf25_, p0), ((0) - (d_2092_cf25_)) + (d_2092_cf25_), d_2093_cf24_, (p0) - (p0)])
                d_2096_v15_ = d_2096_v15_
                d_2097_v16_: _dafny.Map
                d_2097_v16_ = _dafny.Map({d_2067_v6_: d_2095_cf22_})
                d_2098_v17_: D12
                d_2098_v17_ = D12_DC29(d_2097_v16_)
                d_2099_v18_: _dafny.Map
                d_2099_v18_ = _dafny.Map({d_2098_v17_: (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "swnt"))).set(default__.safeIndex(d_2092_cf25_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "swnt")))), d_2071_v10_)})
                d_2099_v18_ = (d_2099_v18_).set(d_2098_v17_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dii")) if d_2091_cf26_ else d_2060_v0_))
                d_2100_v19_: C7
                nw313_ = C7()
                nw313_.ctor__()
                d_2100_v19_ = nw313_
                d_2101_v20_: _dafny.Map
                d_2101_v20_ = _dafny.Map({(d_2093_cf24_) + (d_2092_cf25_): default__.safeModuloInt(p0, p0)})
                d_2101_v20_ = (d_2101_v20_).set(p0, p0)
            elif source29_.is_DC11:
                d_2102___mcc_h9_ = source29_.cf17
                d_2103_cf17_ = d_2102___mcc_h9_
                d_2104_v22_: _dafny.Seq
                d_2104_v22_ = _dafny.SeqWithoutIsStrInference([p0])
                d_2105_v23_: _dafny.MultiSet
                def iife191_():
                    coll55_ = _dafny.Map()
                    compr_55_: int
                    for compr_55_ in (d_2104_v22_).Elements:
                        d_2106_v21_: int = compr_55_
                        if (d_2106_v21_) in (d_2104_v22_):
                            coll55_[(d_2106_v21_) + (p0)] = self.f11
                    return _dafny.Map(coll55_)
                d_2105_v23_ = _dafny.MultiSet([iife191_()
                ])
                d_2107_v24_: _dafny.Map
                d_2107_v24_ = _dafny.Map({p0: self.f11})
                d_2108_v25_: bool
                out16_: bool
                out16_ = (self).m2(((d_2105_v23_) - (_dafny.MultiSet([d_2107_v24_]))).cardinality, len(d_2060_v0_), not(self.f11), globalState)
                d_2108_v25_ = out16_
                d_2109_v26_: _dafny.Array
                nw314_ = _dafny.Array(_dafny.Array(None, 0), 12)
                d_2109_v26_ = nw314_
                d_2110_v27_: _dafny.Array
                nw315_ = _dafny.Array(None, 3)
                nw315_[int(0)] = p0
                nw315_[int(1)] = p0
                nw315_[int(2)] = p0
                d_2110_v27_ = nw315_
                index368_ = default__.safeIndex(399, (d_2109_v26_).length(0))
                (d_2109_v26_)[index368_] = d_2110_v27_
                index369_ = default__.safeIndex(399, (d_2109_v26_).length(0))
                rhs336_ = not(d_2108_v25_)
                rhs337_ = d_2110_v27_
                lhs265_ = self
                lhs266_ = d_2109_v26_
                lhs267_ = default__.safeIndex(399, (d_2109_v26_).length(0))
                lhs265_.f11 = rhs336_
                lhs266_[lhs267_] = rhs337_
                d_2104_v22_ = _dafny.SeqWithoutIsStrInference([533, 549])
                d_2111_v28_: _dafny.Array
                def lambda90_(d_2112_v0_):
                    def lambda91_(d_2113_i2_):
                        return _dafny.Set({len(d_2112_v0_)})

                    return lambda91_

                init52_ = lambda90_(d_2060_v0_)
                nw316_ = _dafny.Array(None, 5)
                for i0_52_ in range(nw316_.length(0)):
                    nw316_[i0_52_] = init52_(i0_52_)
                d_2111_v28_ = nw316_
                d_2111_v28_ = d_2111_v28_
            elif True:
                d_2114___mcc_h10_ = source29_.cf27
                d_2115_cf27_ = d_2114___mcc_h10_
                d_2116_v29_: D1
                d_2116_v29_ = D1_DC1((d_2060_v0_).set(default__.safeIndex(-526, len(d_2060_v0_)), _dafny.CodePoint('s')))
                d_2117_v30_: _dafny.Array
                def lambda92_(d_2118_i3_):
                    return not(self.f11)

                init53_ = lambda92_
                nw317_ = _dafny.Array(None, 9)
                for i0_53_ in range(nw317_.length(0)):
                    nw317_[i0_53_] = init53_(i0_53_)
                d_2117_v30_ = nw317_
                d_2119_v31_: _dafny.Set
                d_2119_v31_ = _dafny.Set({p0})
                d_2120_v32_: C6
                nw318_ = C6()
                nw318_.ctor__(d_2116_v29_, self.f11, d_2117_v30_, d_2119_v31_, D1_DC2(p0, p0, self.f11), self.f11)
                d_2120_v32_ = nw318_
                d_2121_v33_: _dafny.Map
                d_2121_v33_ = _dafny.Map({d_2120_v32_: d_2068_v7_})
                d_2122_v34_: _dafny.MultiSet
                d_2122_v34_ = _dafny.MultiSet([self.f11, (d_2120_v32_).fm11(globalState)])
                d_2123_v35_: _dafny.Map
                d_2123_v35_ = _dafny.Map({(d_2122_v34_).cardinality: p0})
                d_2124_v36_: _dafny.Seq
                d_2124_v36_ = _dafny.SeqWithoutIsStrInference([len(d_2123_v35_), p0, 352, p0])
                d_2125_v37_: _dafny.Array
                nw319_ = _dafny.Array(None, 13)
                nw319_[int(0)] = len(d_2060_v0_)
                nw319_[int(1)] = len(d_2121_v33_)
                nw319_[int(2)] = default__.safeModuloInt(p0, len(_dafny.Map({True: p0})))
                nw319_[int(3)] = p0
                nw319_[int(4)] = len(d_2124_v36_)
                nw319_[int(5)] = p0
                nw319_[int(6)] = (len(d_2060_v0_)) - (983)
                nw319_[int(7)] = len(d_2060_v0_)
                nw319_[int(8)] = p0
                nw319_[int(9)] = p0
                nw319_[int(10)] = p0
                nw319_[int(11)] = len((d_2068_v7_) | (_dafny.Map({self.f11: True})))
                nw319_[int(12)] = p0
                d_2125_v37_ = nw319_
                d_2126_v38_: D2
                d_2126_v38_ = D2_DC4(d_2125_v37_, self.f11, p0, self.f11, d_2117_v30_)
                index370_ = default__.safeIndex(300, (d_2125_v37_).length(0))
                (d_2125_v37_)[index370_] = (((d_2123_v35_)[p0] if (p0) in (d_2123_v35_) else p0) if self.f11 else (d_2126_v38_).cf8)
                index371_ = default__.safeIndex(300, (d_2125_v37_).length(0))
                (d_2125_v37_)[index371_] = ((p0) * (p0)) - (p0)
                (globalState).f7 = (d_2125_v37_)[default__.safeIndex(300, (d_2125_v37_).length(0))]
                (globalState).f7 = default__.safeDivisionInt(len(d_2060_v0_), (d_2125_v37_)[default__.safeIndex(300, (d_2125_v37_).length(0))])
                index372_ = default__.safeIndex(300, (d_2125_v37_).length(0))
                (d_2125_v37_)[index372_] = p0
            d_2127_v39_: _dafny.Array
            def lambda93_(d_2128_v0_):
                def lambda94_(d_2129_i4_):
                    return d_2128_v0_

                return lambda94_

            init54_ = lambda93_(d_2060_v0_)
            nw320_ = _dafny.Array(None, 1)
            for i0_54_ in range(nw320_.length(0)):
                nw320_[i0_54_] = init54_(i0_54_)
            d_2127_v39_ = nw320_
            d_2127_v39_ = d_2127_v39_
        elif True:
            d_2130_v40_: _dafny.Seq
            d_2130_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dpqp"))
            d_2131_v41_: str
            d_2131_v41_ = _dafny.CodePoint('y')
            d_2132_v42_: _dafny.Array
            nw321_ = _dafny.Array(False, 11)
            d_2132_v42_ = nw321_
            d_2133_v43_: D12
            d_2133_v43_ = D12_DC30(d_2132_v42_, self.f11)
            d_2134_v44_: _dafny.Map
            d_2134_v44_ = _dafny.Map({p0: False})
            d_2135_v45_: _dafny.Map
            d_2135_v45_ = _dafny.Map({self.f11: self.f11})
            d_2136_v46_: _dafny.Array
            nw322_ = _dafny.Array(None, 27)
            nw322_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sq"))
            nw322_[int(1)] = d_2130_v40_
            nw322_[int(2)] = d_2130_v40_
            nw322_[int(3)] = d_2130_v40_
            nw322_[int(4)] = (d_2130_v40_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ugei")))
            nw322_[int(5)] = d_2130_v40_
            nw322_[int(6)] = d_2130_v40_
            nw322_[int(7)] = (d_2130_v40_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vjrdhd")))
            nw322_[int(8)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hd"))) + (d_2130_v40_)
            nw322_[int(9)] = d_2130_v40_
            nw322_[int(10)] = d_2130_v40_
            nw322_[int(11)] = d_2130_v40_
            nw322_[int(12)] = (d_2130_v40_).set(default__.safeIndex(p0, len(d_2130_v40_)), d_2131_v41_)
            nw322_[int(13)] = d_2130_v40_
            nw322_[int(14)] = d_2130_v40_
            nw322_[int(15)] = d_2130_v40_
            nw322_[int(16)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "scbj"))).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "scbj")))), d_2131_v41_)
            nw322_[int(17)] = d_2130_v40_
            nw322_[int(18)] = d_2130_v40_
            nw322_[int(19)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bbwe")) if self.f11 else d_2130_v40_)
            nw322_[int(20)] = d_2130_v40_
            nw322_[int(21)] = ((d_2130_v40_).set(default__.safeIndex(p0, len(d_2130_v40_)), d_2131_v41_) if self.f11 else _dafny.SeqWithoutIsStrInference([d_2131_v41_ for d_2137_i5_ in range(default__.abs(34))]))
            nw322_[int(22)] = default__.fm60(self.f11, not(not((d_2133_v43_).cf59)), d_2134_v44_, p0, globalState)
            nw322_[int(23)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ihv"))).set(default__.safeIndex(len(d_2135_v45_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ihv")))), _dafny.CodePoint('e'))
            nw322_[int(24)] = d_2130_v40_
            nw322_[int(25)] = (d_2130_v40_) + (d_2130_v40_)
            nw322_[int(26)] = (d_2130_v40_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kf")))
            d_2136_v46_ = nw322_
            index373_ = default__.safeIndex(779, (d_2136_v46_).length(0))
            (d_2136_v46_)[index373_] = (d_2130_v40_) + (_dafny.SeqWithoutIsStrInference([d_2131_v41_ for d_2138_i6_ in range(default__.abs(927))]))
            d_2139_v47_: _dafny.Array
            nw323_ = _dafny.Array(_dafny.CodePoint('D'), 17)
            d_2139_v47_ = nw323_
            d_2140_v48_: _dafny.Map
            d_2140_v48_ = _dafny.Map({828: d_2139_v47_})
            d_2141_v49_: _dafny.MultiSet
            d_2141_v49_ = _dafny.MultiSet([d_2139_v47_, d_2139_v47_, ((d_2140_v48_)[p0] if (p0) in (d_2140_v48_) else d_2139_v47_)])
            d_2142_v50_: _dafny.MultiSet
            d_2142_v50_ = _dafny.MultiSet([773])
            index374_ = default__.safeIndex(779, (d_2136_v46_).length(0))
            rhs338_ = _dafny.SeqWithoutIsStrInference([d_2131_v41_ for d_2143_i7_ in range(default__.abs(994))])
            rhs339_ = d_2141_v49_
            rhs340_ = ((d_2134_v44_)[p0] if (p0) in (d_2134_v44_) else (d_2142_v50_).isdisjoint(d_2142_v50_))
            rhs341_ = self.f11
            lhs268_ = d_2136_v46_
            lhs269_ = default__.safeIndex(779, (d_2136_v46_).length(0))
            lhs270_ = self
            lhs271_ = self
            lhs268_[lhs269_] = rhs338_
            d_2141_v49_ = rhs339_
            lhs270_.f11 = rhs340_
            lhs271_.f11 = rhs341_
            index375_ = default__.safeIndex(228, (d_2132_v42_).length(0))
            (d_2132_v42_)[index375_] = (self.f11 if self.f11 else self.f11)
            d_2144_v51_: _dafny.Array
            def lambda95_(d_2145_p0_):
                def lambda96_(d_2146_i8_):
                    return default__.safeDivisionInt(d_2146_i8_, d_2145_p0_)

                return lambda96_

            init55_ = lambda95_(p0)
            nw324_ = _dafny.Array(None, 23)
            for i0_55_ in range(nw324_.length(0)):
                nw324_[i0_55_] = init55_(i0_55_)
            d_2144_v51_ = nw324_
            index376_ = default__.safeIndex(314, (d_2144_v51_).length(0))
            (d_2144_v51_)[index376_] = p0
            d_2147_v52_: _dafny.Seq
            d_2147_v52_ = _dafny.SeqWithoutIsStrInference([True, False, True])
            d_2148_v53_: _dafny.Map
            d_2148_v53_ = _dafny.Map({default__.fm30(596, globalState): d_2132_v42_})
            index377_ = default__.safeIndex(228, (d_2132_v42_).length(0))
            index378_ = default__.safeIndex(314, (d_2144_v51_).length(0))
            rhs342_ = not (True) or ((len(d_2147_v52_)) >= (p0))
            rhs343_ = p0
            rhs344_ = (p0) * (default__.safeModuloInt(default__.fm0(d_2147_v52_, self.f11, self.f11, p0, globalState), p0))
            rhs345_ = ((d_2148_v53_)[default__.fm30((0) - (p0), globalState)] if (default__.fm30((0) - (p0), globalState)) in (d_2148_v53_) else d_2132_v42_)
            rhs346_ = default__.fm0(d_2147_v52_, (self.f11) and (self.f11), False, 953, globalState)
            lhs272_ = d_2132_v42_
            lhs273_ = default__.safeIndex(228, (d_2132_v42_).length(0))
            lhs274_ = globalState
            lhs275_ = globalState
            lhs276_ = d_2144_v51_
            lhs277_ = default__.safeIndex(314, (d_2144_v51_).length(0))
            lhs272_[lhs273_] = rhs342_
            lhs274_.f7 = rhs343_
            lhs275_.f7 = rhs344_
            d_2132_v42_ = rhs345_
            lhs276_[lhs277_] = rhs346_
            (globalState).f7 = ((d_2144_v51_)[default__.safeIndex(314, (d_2144_v51_).length(0))] if (len(d_2130_v40_)) <= (p0) else default__.fm0(d_2147_v52_, self.f11, (d_2132_v42_)[default__.safeIndex(228, (d_2132_v42_).length(0))], p0, globalState))
            index379_ = default__.safeIndex(228, (d_2132_v42_).length(0))
            (d_2132_v42_)[index379_] = self.f11
            d_2134_v44_ = (d_2134_v44_).set(p0, (d_2132_v42_)[default__.safeIndex(228, (d_2132_v42_).length(0))])
        d_2149_v54_: _dafny.Array
        def lambda97_(d_2150_i10_):
            return (_dafny.MultiSet([self.f11])).ispropersubset(_dafny.MultiSet([self.f11]))

        init56_ = lambda97_
        nw325_ = _dafny.Array(None, 26)
        for i0_56_ in range(nw325_.length(0)):
            nw325_[i0_56_] = init56_(i0_56_)
        d_2149_v54_ = nw325_
        guard_loop_9_: int
        for guard_loop_9_ in _dafny.IntegerRange(0, (d_2149_v54_).length(0)):
            d_2151_i9_: int = guard_loop_9_
            if (True) and (((0) <= (d_2151_i9_)) and ((d_2151_i9_) < ((d_2149_v54_).length(0)))):
                (d_2149_v54_)[(d_2151_i9_)] = self.f11
        d_2152_v55_: _dafny.Array
        def lambda98_(d_2153_p0_):
            def lambda99_(d_2154_i12_):
                return _dafny.Set({d_2153_p0_, d_2153_p0_})

            return lambda99_

        init57_ = lambda98_(p0)
        nw326_ = _dafny.Array(None, 17)
        for i0_57_ in range(nw326_.length(0)):
            nw326_[i0_57_] = init57_(i0_57_)
        d_2152_v55_ = nw326_
        guard_loop_10_: int
        for guard_loop_10_ in _dafny.IntegerRange(0, (d_2152_v55_).length(0)):
            d_2155_i11_: int = guard_loop_10_
            if (True) and (((0) <= (d_2155_i11_)) and ((d_2155_i11_) < ((d_2152_v55_).length(0)))):
                (d_2152_v55_)[(d_2155_i11_)] = _dafny.Set({(0) - (default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([True])), 235)), p0})
        (self).f11 = (p0) != ((-217) * (p0))
        if True:
            d_2156_v56_: str
            d_2156_v56_ = _dafny.CodePoint('n')
            d_2157_v57_: _dafny.Map
            d_2157_v57_ = _dafny.Map({True: p0})
            d_2158_v58_: D1
            d_2158_v58_ = D1_DC2(((d_2157_v57_)[self.f11] if (self.f11) in (d_2157_v57_) else p0), (0) - (p0), self.f11)
            d_2159_v59_: C0
            nw327_ = C0()
            nw327_.ctor__(self.f11, d_2156_v56_, d_2158_v58_, not(self.f11))
            d_2159_v59_ = nw327_
            d_2159_v59_ = d_2159_v59_
            (globalState).f7 = p0
            if d_2159_v59_.f19:
                d_2160_v60_: _dafny.Seq
                d_2160_v60_ = _dafny.SeqWithoutIsStrInference([d_2159_v59_.f19, d_2159_v59_.f19, self.f11])
                (globalState).f7 = default__.fm0(d_2160_v60_, d_2159_v59_.f19, self.f11, (len(d_2160_v60_)) - (p0), globalState)
                d_2161_v61_: _dafny.Seq
                d_2161_v61_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ped"))
                rhs347_ = (len(_dafny.Map({p0: default__.fm30(len(_dafny.Set({d_2161_v61_, d_2161_v61_})), globalState)}))) + (default__.fm0(_dafny.SeqWithoutIsStrInference([d_2159_v59_.f19, d_2159_v59_.f19]), self.f11, self.f11, p0, globalState))
                rhs348_ = p0
                lhs278_ = globalState
                lhs279_ = globalState
                lhs278_.f7 = rhs347_
                lhs279_.f7 = rhs348_
                d_2162_v62_: _dafny.Seq
                d_2162_v62_ = _dafny.SeqWithoutIsStrInference([7])
                d_2163_v63_: D7
                d_2163_v63_ = D7_DC15(d_2162_v62_)
                d_2164_v64_: _dafny.Map
                d_2164_v64_ = _dafny.Map({d_2159_v59_.f19: (d_2159_v59_).f20})
                d_2165_v65_: C4
                nw328_ = C4()
                nw328_.ctor__()
                d_2165_v65_ = nw328_
                d_2166_v66_: _dafny.Seq
                d_2166_v66_ = _dafny.SeqWithoutIsStrInference([d_2165_v65_])
                d_2167_v67_: _dafny.Array
                nw329_ = _dafny.Array(None, 20)
                nw329_[int(0)] = (d_2162_v62_) + (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([self.f11, d_2159_v59_.f19])).cardinality, default__.fm0(d_2160_v60_, self.f11, d_2159_v59_.f19, p0, globalState)]))
                nw329_[int(1)] = d_2162_v62_
                nw329_[int(2)] = d_2162_v62_
                nw329_[int(3)] = d_2162_v62_
                nw329_[int(4)] = ((d_2163_v63_).cf28) + (d_2162_v62_)
                nw329_[int(5)] = _dafny.SeqWithoutIsStrInference([p0 for d_2168_i13_ in range(default__.abs(828))])
                nw329_[int(6)] = d_2162_v62_
                nw329_[int(7)] = d_2162_v62_
                nw329_[int(8)] = (d_2162_v62_) + (d_2162_v62_)
                nw329_[int(9)] = d_2162_v62_
                nw329_[int(10)] = d_2162_v62_
                nw329_[int(11)] = d_2162_v62_
                nw329_[int(12)] = (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({432, len(_dafny.Set({d_2159_v59_.f19, self.f11, d_2159_v59_.f19, self.f11, d_2159_v59_.f19})), p0})) for d_2169_i14_ in range(default__.abs(973))])) + (d_2162_v62_)
                nw329_[int(13)] = d_2162_v62_
                nw329_[int(14)] = _dafny.SeqWithoutIsStrInference([p0])
                nw329_[int(15)] = (d_2162_v62_) + ((d_2162_v62_).set(default__.safeIndex(len(d_2164_v64_), len(d_2162_v62_)), p0))
                nw329_[int(16)] = d_2162_v62_
                nw329_[int(17)] = d_2162_v62_
                nw329_[int(18)] = d_2162_v62_
                nw329_[int(19)] = _dafny.SeqWithoutIsStrInference([p0, len(d_2166_v66_)])
                d_2167_v67_ = nw329_
                rhs349_ = (0) - (p0)
                rhs350_ = d_2167_v67_
                lhs280_ = globalState
                lhs280_.f7 = rhs349_
                d_2167_v67_ = rhs350_
                index380_ = default__.safeIndex(752, (d_2149_v54_).length(0))
                (d_2149_v54_)[index380_] = d_2159_v59_.f19
                d_2170_v68_: _dafny.Map
                d_2170_v68_ = _dafny.Map({p0: d_2159_v59_.f19})
                index381_ = default__.safeIndex(752, (d_2149_v54_).length(0))
                (d_2149_v54_)[index381_] = (d_2162_v62_) == ((_dafny.SeqWithoutIsStrInference([p0, len(d_2170_v68_), p0])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([p0, len(d_2170_v68_), p0]))), 725))
                d_2171_v69_: _dafny.Set
                d_2171_v69_ = _dafny.Set({d_2159_v59_.f19, True, d_2159_v59_.f19})
                d_2172_v70_: _dafny.MultiSet
                d_2172_v70_ = _dafny.MultiSet([self.f11, self.f11])
                index382_ = default__.safeIndex(752, (d_2149_v54_).length(0))
                index383_ = default__.safeIndex(752, (d_2149_v54_).length(0))
                rhs351_ = ((d_2149_v54_)[default__.safeIndex(752, (d_2149_v54_).length(0))]) not in (d_2171_v69_)
                rhs352_ = ((default__.fm40(globalState)).set(self.f11, default__.abs(len(d_2171_v69_)))).issubset(d_2172_v70_)
                lhs281_ = d_2149_v54_
                lhs282_ = default__.safeIndex(752, (d_2149_v54_).length(0))
                lhs283_ = d_2149_v54_
                lhs284_ = default__.safeIndex(752, (d_2149_v54_).length(0))
                lhs281_[lhs282_] = rhs351_
                lhs283_[lhs284_] = rhs352_
            elif True:
                d_2173_v71_: _dafny.Seq
                d_2173_v71_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "arsssdc"))
                d_2156_v56_ = (d_2173_v71_)[default__.safeIndex(p0, len(d_2173_v71_))]
                d_2174_v72_: _dafny.Set
                d_2174_v72_ = _dafny.Set({d_2173_v71_})
                d_2175_v73_: _dafny.Map
                d_2175_v73_ = _dafny.Map({(d_2174_v72_) - (_dafny.Set({d_2173_v71_, d_2173_v71_})): d_2159_v59_.f19})
                d_2175_v73_ = (d_2175_v73_).set(d_2174_v72_, d_2159_v59_.f19)
                (d_2159_v59_).f19 = (False if d_2159_v59_.f19 else False)
                index384_ = default__.safeIndex(399, (d_2149_v54_).length(0))
                (d_2149_v54_)[index384_] = False
                index385_ = default__.safeIndex(399, (d_2149_v54_).length(0))
                (d_2149_v54_)[index385_] = self.f11
                d_2176_v74_: _dafny.Set
                d_2176_v74_ = _dafny.Set({p0})
                d_2177_v75_: _dafny.Map
                d_2177_v75_ = _dafny.Map({self.f11: (d_2176_v74_).issubset(d_2176_v74_)})
                d_2178_v76_: _dafny.Array
                def lambda100_(d_2179_p0_):
                    def lambda101_(d_2180_i15_):
                        return (d_2180_i15_) - (d_2179_p0_)

                    return lambda101_

                init58_ = lambda100_(p0)
                nw330_ = _dafny.Array(None, 8)
                for i0_58_ in range(nw330_.length(0)):
                    nw330_[i0_58_] = init58_(i0_58_)
                d_2178_v76_ = nw330_
                d_2181_v77_: _dafny.MultiSet
                d_2181_v77_ = _dafny.MultiSet([(d_2159_v59_).f20])
                d_2182_v78_: _dafny.Map
                d_2182_v78_ = _dafny.Map({self.f11: d_2181_v77_})
                d_2183_v79_: D12
                d_2183_v79_ = D12_DC31(self.f11, d_2178_v76_, d_2159_v59_.f19, (d_2159_v59_).f20, ((d_2182_v78_)[d_2159_v59_.f19] if (d_2159_v59_.f19) in (d_2182_v78_) else _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(d_2159_v59_).f20]))))
                d_2184_v80_: _dafny.Set
                d_2184_v80_ = _dafny.Set({True, (d_2149_v54_)[default__.safeIndex(399, (d_2149_v54_).length(0))], (d_2149_v54_)[default__.safeIndex(399, (d_2149_v54_).length(0))], d_2159_v59_.f19, True})
                index386_ = default__.safeIndex(399, (d_2149_v54_).length(0))
                (d_2149_v54_)[index386_] = ((d_2177_v75_)[not ((d_2149_v54_)[default__.safeIndex(399, (d_2149_v54_).length(0))]) or ((d_2183_v79_).cf62)] if (not ((d_2149_v54_)[default__.safeIndex(399, (d_2149_v54_).length(0))]) or ((d_2183_v79_).cf62)) in (d_2177_v75_) else not(default__.fm30(len(d_2184_v80_), globalState)))
            d_2185_v81_: _dafny.Seq
            d_2185_v81_ = _dafny.SeqWithoutIsStrInference([d_2159_v59_.f19, True])
            d_2186_v82_: _dafny.MultiSet
            d_2186_v82_ = _dafny.MultiSet([len(d_2185_v81_)])
            d_2187_v83_: _dafny.Map
            d_2187_v83_ = _dafny.Map({d_2186_v82_: d_2149_v54_})
            d_2188_v84_: _dafny.Set
            d_2188_v84_ = _dafny.Set({p0})
            d_2189_v85_: C5
            nw331_ = C5()
            nw331_.ctor__(d_2159_v59_.f19, ((d_2187_v83_)[d_2186_v82_] if (d_2186_v82_) in (d_2187_v83_) else d_2149_v54_), (d_2188_v84_) - (d_2188_v84_), d_2158_v58_, d_2159_v59_.f19)
            d_2189_v85_ = nw331_
            if self.f11:
                d_2190_v86_: _dafny.Seq
                d_2190_v86_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hhqingxd"))
                index387_ = default__.safeIndex(6, (d_2149_v54_).length(0))
                (d_2149_v54_)[index387_] = (self).fm4(d_2185_v81_, p0, d_2190_v86_, d_2189_v85_.f24, globalState)
                d_2191_v87_: _dafny.Seq
                d_2191_v87_ = _dafny.SeqWithoutIsStrInference([d_2190_v86_, d_2190_v86_, d_2190_v86_, d_2190_v86_, d_2190_v86_])
                d_2192_v88_: _dafny.Map
                d_2192_v88_ = _dafny.Map({(d_2191_v87_)[default__.safeIndex(p0, len(d_2191_v87_))]: d_2157_v57_})
                index388_ = default__.safeIndex(6, (d_2149_v54_).length(0))
                (d_2149_v54_)[index388_] = (d_2190_v86_) in (d_2192_v88_)
                d_2193_v89_: _dafny.Seq
                d_2193_v89_ = _dafny.SeqWithoutIsStrInference([p0, (0) - (p0)])
                d_2194_v90_: _dafny.Map
                d_2194_v90_ = _dafny.Map({p0: not(d_2189_v85_.f24)})
                d_2195_v91_: _dafny.Set
                d_2195_v91_ = _dafny.Set({False, d_2189_v85_.f24})
                d_2196_v92_: _dafny.Map
                d_2196_v92_ = _dafny.Map({len(d_2195_v91_): p0})
                d_2197_v93_: _dafny.Seq
                d_2197_v93_ = _dafny.SeqWithoutIsStrInference([d_2186_v82_, d_2186_v82_])
                d_2198_v94_: _dafny.Map
                d_2198_v94_ = _dafny.Map({d_2195_v91_: len(d_2185_v81_)})
                d_2199_v95_: _dafny.Array
                nw332_ = _dafny.Array(None, 16)
                nw332_[int(0)] = (d_2193_v89_)[default__.safeIndex((d_2189_v85_).fm9((d_2159_v59_).f20, globalState), len(d_2193_v89_))]
                nw332_[int(1)] = p0
                nw332_[int(2)] = len(d_2194_v90_)
                nw332_[int(3)] = p0
                nw332_[int(4)] = (0) - ((d_2189_v85_).fm9((d_2159_v59_).f20, globalState))
                nw332_[int(5)] = p0
                nw332_[int(6)] = (p0) + ((0) - (((d_2196_v92_)[p0] if (p0) in (d_2196_v92_) else p0)))
                nw332_[int(7)] = p0
                nw332_[int(8)] = (0) - (len((_dafny.SeqWithoutIsStrInference([p0, p0])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([p0, p0]))), (0) - (p0))))
                nw332_[int(9)] = (p0) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))))
                nw332_[int(10)] = p0
                nw332_[int(11)] = (0) - (p0)
                nw332_[int(12)] = len(d_2194_v90_)
                nw332_[int(13)] = 366
                nw332_[int(14)] = default__.safeDivisionInt(((d_2197_v93_)[default__.safeIndex(p0, len(d_2197_v93_))]).cardinality, len(default__.fm64(globalState)))
                nw332_[int(15)] = len((d_2198_v94_) | (d_2198_v94_))
                d_2199_v95_ = nw332_
                d_2199_v95_ = d_2199_v95_
                (globalState).f7 = (p0) * (p0)
                rhs353_ = not(self.f11)
                rhs354_ = p0
                lhs285_ = d_2189_v85_
                lhs286_ = globalState
                lhs285_.f24 = rhs353_
                lhs286_.f7 = rhs354_
                (self).f11 = (d_2159_v59_.f19) == ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sftaoljs"))) <= ((d_2191_v87_)[default__.safeIndex(len(d_2195_v91_), len(d_2191_v87_))]))
            elif True:
                d_2200_v96_: _dafny.Map
                d_2200_v96_ = _dafny.Map({True: self.f11})
                d_2201_v97_: _dafny.MultiSet
                d_2201_v97_ = _dafny.MultiSet([(d_2200_v96_) | (d_2200_v96_)])
                d_2202_v98_: _dafny.Seq
                d_2202_v98_ = _dafny.SeqWithoutIsStrInference([d_2200_v96_, d_2200_v96_, d_2200_v96_, _dafny.Map({d_2159_v59_.f19: d_2159_v59_.f19}), (d_2200_v96_).set(d_2189_v85_.f24, self.f11)])
                d_2201_v97_ = (_dafny.MultiSet(d_2202_v98_)) - (d_2201_v97_)
                d_2203_v99_: _dafny.MultiSet
                d_2203_v99_ = _dafny.MultiSet([d_2149_v54_, d_2149_v54_, d_2149_v54_, d_2149_v54_])
                d_2203_v99_ = d_2203_v99_
                d_2204_v100_: _dafny.Seq
                d_2204_v100_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bvpukkl"))
                d_2205_v102_: _dafny.Set
                d_2205_v102_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ipcpel")), d_2204_v100_})
                d_2206_v103_: _dafny.MultiSet
                d_2206_v103_ = _dafny.MultiSet([d_2204_v100_])
                d_2207_v105_: _dafny.Seq
                def iife192_():
                    coll56_ = _dafny.Set()
                    compr_56_: _dafny.Seq
                    for compr_56_ in (d_2206_v103_).Elements:
                        d_2209_v104_: _dafny.Seq = compr_56_
                        if (d_2209_v104_) in (d_2206_v103_):
                            coll56_ = coll56_.union(_dafny.Set([d_2209_v104_]))
                    return _dafny.Set(coll56_)
                d_2207_v105_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_2204_v100_, _dafny.SeqWithoutIsStrInference([(d_2159_v59_).f20 for d_2208_i16_ in range(default__.abs(-402))]), d_2204_v100_}), iife192_()
                ])
                d_2210_v106_: _dafny.Seq
                d_2210_v106_ = _dafny.SeqWithoutIsStrInference([(0) - (p0), p0])
                d_2211_v108_: _dafny.Map
                d_2211_v108_ = _dafny.Map({p0: p0})
                d_2212_v109_: C7
                nw333_ = C7()
                nw333_.ctor__()
                d_2212_v109_ = nw333_
                d_2213_v110_: _dafny.Map
                d_2213_v110_ = _dafny.Map({d_2204_v100_: d_2212_v109_})
                d_2214_v111_: _dafny.Map
                d_2214_v111_ = _dafny.Map({p0: ((d_2211_v108_)[p0] if (p0) in (d_2211_v108_) else len(d_2213_v110_))})
                d_2215_v112_: _dafny.Array
                nw334_ = _dafny.Array(None, 17)
                def iife193_():
                    coll57_ = _dafny.Set()
                    compr_57_: _dafny.Seq
                    for compr_57_ in (_dafny.SeqWithoutIsStrInference([d_2204_v100_])).Elements:
                        d_2216_v101_: _dafny.Seq = compr_57_
                        if (d_2216_v101_) in (_dafny.SeqWithoutIsStrInference([d_2204_v100_])):
                            coll57_ = coll57_.union(_dafny.Set([d_2216_v101_]))
                    return _dafny.Set(coll57_)
                nw334_[int(0)] = iife193_()
                
                nw334_[int(1)] = default__.fm48(globalState)
                nw334_[int(2)] = (d_2205_v102_) | (d_2205_v102_)
                nw334_[int(3)] = d_2205_v102_
                nw334_[int(4)] = d_2205_v102_
                nw334_[int(5)] = ((d_2207_v105_)[default__.safeIndex((d_2210_v106_)[default__.safeIndex(p0, len(d_2210_v106_))], len(d_2207_v105_))]) - (d_2205_v102_)
                def iife194_():
                    coll58_ = _dafny.Set()
                    compr_58_: _dafny.Seq
                    for compr_58_ in (d_2206_v103_).Elements:
                        d_2217_v107_: _dafny.Seq = compr_58_
                        if (d_2217_v107_) in (d_2206_v103_):
                            coll58_ = coll58_.union(_dafny.Set([d_2217_v107_]))
                    return _dafny.Set(coll58_)
                nw334_[int(6)] = iife194_()
                
                nw334_[int(7)] = d_2205_v102_
                nw334_[int(8)] = d_2205_v102_
                nw334_[int(9)] = d_2205_v102_
                nw334_[int(10)] = d_2205_v102_
                nw334_[int(11)] = d_2205_v102_
                nw334_[int(12)] = (d_2207_v105_)[default__.safeIndex(len(d_2211_v108_), len(d_2207_v105_))]
                nw334_[int(13)] = d_2205_v102_
                nw334_[int(14)] = d_2205_v102_
                nw334_[int(15)] = d_2205_v102_
                nw334_[int(16)] = (d_2205_v102_).intersection(default__.fm48(globalState))
                d_2215_v112_ = nw334_
                index389_ = default__.safeIndex(813, (d_2215_v112_).length(0))
                (d_2215_v112_)[index389_] = d_2205_v102_
                index390_ = default__.safeIndex(813, (d_2215_v112_).length(0))
                (d_2215_v112_)[index390_] = (d_2205_v102_ if (self).fm4(d_2185_v81_, p0, d_2204_v100_, d_2159_v59_.f19, globalState) else (_dafny.Set({d_2204_v100_, d_2204_v100_})) - (d_2205_v102_))
                d_2218_v113_: _dafny.Array
                nw335_ = _dafny.Array(int(0), 18)
                d_2218_v113_ = nw335_
                d_2219_v114_: _dafny.MultiSet
                d_2219_v114_ = _dafny.MultiSet([d_2218_v113_])
                d_2219_v114_ = (d_2219_v114_) - (d_2219_v114_)
                (self).f11 = not(False)
        elif True:
            d_2220_v115_: str
            d_2220_v115_ = _dafny.CodePoint('i')
            d_2221_v116_: _dafny.Map
            d_2221_v116_ = _dafny.Map({(p0) <= (680): d_2220_v115_})
            d_2221_v116_ = d_2221_v116_
            d_2222_v117_: _dafny.Array
            nw336_ = _dafny.Array(int(0), 4)
            d_2222_v117_ = nw336_
            d_2223_v118_: _dafny.Set
            d_2223_v118_ = _dafny.Set({self.f11})
            index391_ = default__.safeIndex(574, (d_2222_v117_).length(0))
            (d_2222_v117_)[index391_] = (len(d_2223_v118_)) - (p0)
            d_2224_v119_: _dafny.Map
            d_2224_v119_ = _dafny.Map({p0: d_2223_v118_})
            d_2225_v120_: _dafny.Seq
            d_2225_v120_ = _dafny.SeqWithoutIsStrInference([self.f11, self.f11])
            d_2226_v121_: _dafny.Map
            d_2226_v121_ = _dafny.Map({self.f11: d_2223_v118_})
            index392_ = default__.safeIndex(574, (d_2222_v117_).length(0))
            (d_2222_v117_)[index392_] = len(((d_2224_v119_)[(0) - (default__.fm0(d_2225_v120_, self.f11, self.f11, p0, globalState))] if ((0) - (default__.fm0(d_2225_v120_, self.f11, self.f11, p0, globalState))) in (d_2224_v119_) else (((d_2226_v121_)[False] if (False) in (d_2226_v121_) else d_2223_v118_) if self.f11 else d_2223_v118_)))
            d_2227_v122_: _dafny.Set
            d_2227_v122_ = _dafny.Set({(0) - (len(d_2223_v118_)), (d_2222_v117_)[default__.safeIndex(574, (d_2222_v117_).length(0))], p0})
            rhs355_ = d_2227_v122_
            rhs356_ = ((p0) + (p0)) - (p0)
            rhs357_ = self.f11
            lhs287_ = globalState
            lhs288_ = self
            d_2227_v122_ = rhs355_
            lhs287_.f7 = rhs356_
            lhs288_.f11 = rhs357_
            if self.f11:
                (self).f11 = ((not(self.f11) if self.f11 else self.f11) if self.f11 else (p0) != (p0))
                d_2228_v123_: _dafny.Seq
                d_2228_v123_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lf"))
                index393_ = default__.safeIndex(574, (d_2222_v117_).length(0))
                (d_2222_v117_)[index393_] = len(_dafny.SeqWithoutIsStrInference([(0) - ((default__.fm55(len(d_2228_v123_), _dafny.Set({_dafny.Map({d_2220_v115_: p0})}), self.f11, globalState)).cardinality), p0, (d_2222_v117_)[default__.safeIndex(574, (d_2222_v117_).length(0))]]))
                d_2229_v124_: _dafny.Array
                nw337_ = _dafny.Array(_dafny.Set({}), 15)
                d_2229_v124_ = nw337_
                d_2230_v125_: _dafny.Seq
                d_2230_v125_ = _dafny.SeqWithoutIsStrInference([d_2223_v118_])
                index394_ = default__.safeIndex(957, (d_2229_v124_).length(0))
                (d_2229_v124_)[index394_] = (d_2230_v125_)[default__.safeIndex(p0, len(d_2230_v125_))]
                d_2231_v126_: _dafny.Seq
                d_2231_v126_ = _dafny.SeqWithoutIsStrInference([d_2149_v54_])
                d_2232_v127_: _dafny.Seq
                d_2232_v127_ = _dafny.SeqWithoutIsStrInference([d_2228_v123_, d_2228_v123_])
                d_2233_v128_: _dafny.Map
                d_2233_v128_ = _dafny.Map({not(self.f11): _dafny.SeqWithoutIsStrInference([self.f11, self.f11])})
                index395_ = default__.safeIndex(957, (d_2229_v124_).length(0))
                index396_ = default__.safeIndex(574, (d_2222_v117_).length(0))
                rhs358_ = d_2223_v118_
                rhs359_ = (d_2231_v126_)[default__.safeIndex((0) - ((len((d_2232_v127_)[default__.safeIndex(p0, len(d_2232_v127_))])) - (p0)), len(d_2231_v126_))]
                rhs360_ = default__.safeModuloInt((d_2222_v117_)[default__.safeIndex(574, (d_2222_v117_).length(0))], (d_2222_v117_)[default__.safeIndex(574, (d_2222_v117_).length(0))])
                rhs361_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))
                rhs362_ = (len((((d_2233_v128_)[self.f11] if (self.f11) in (d_2233_v128_) else d_2225_v120_)) + (d_2225_v120_))) > ((p0) + ((d_2222_v117_)[default__.safeIndex(574, (d_2222_v117_).length(0))]))
                lhs289_ = d_2229_v124_
                lhs290_ = default__.safeIndex(957, (d_2229_v124_).length(0))
                lhs291_ = d_2222_v117_
                lhs292_ = default__.safeIndex(574, (d_2222_v117_).length(0))
                lhs293_ = self
                lhs289_[lhs290_] = rhs358_
                d_2149_v54_ = rhs359_
                lhs291_[lhs292_] = rhs360_
                d_2228_v123_ = rhs361_
                lhs293_.f11 = rhs362_
                d_2222_v117_ = d_2222_v117_
                d_2234_v129_: _dafny.Array
                def lambda102_(d_2235_v117_):
                    def lambda103_(d_2236_i17_):
                        return _dafny.MultiSet([(d_2235_v117_)[default__.safeIndex(574, (d_2235_v117_).length(0))]])

                    return lambda103_

                init59_ = lambda102_(d_2222_v117_)
                nw338_ = _dafny.Array(None, 14)
                for i0_59_ in range(nw338_.length(0)):
                    nw338_[i0_59_] = init59_(i0_59_)
                d_2234_v129_ = nw338_
                d_2237_v130_: _dafny.MultiSet
                d_2237_v130_ = _dafny.MultiSet([False, self.f11, self.f11])
                d_2238_v131_: _dafny.MultiSet
                d_2238_v131_ = _dafny.MultiSet([4, p0, (d_2222_v117_)[default__.safeIndex(574, (d_2222_v117_).length(0))], (d_2222_v117_)[default__.safeIndex(574, (d_2222_v117_).length(0))], ((d_2237_v130_)[self.f11] if (self.f11) in (d_2237_v130_) else (d_2222_v117_)[default__.safeIndex(574, (d_2222_v117_).length(0))])])
                index397_ = default__.safeIndex(912, (d_2234_v129_).length(0))
                (d_2234_v129_)[index397_] = d_2238_v131_
                index398_ = default__.safeIndex(912, (d_2234_v129_).length(0))
                (d_2234_v129_)[index398_] = d_2238_v131_
            elif True:
                d_2239_v132_: D7
                d_2239_v132_ = D7_DC16((d_2222_v117_)[default__.safeIndex(574, (d_2222_v117_).length(0))])
                d_2240_v133_: _dafny.Map
                d_2240_v133_ = _dafny.Map({D2_DC3(d_2149_v54_): self.f11})
                d_2241_v134_: D2
                d_2241_v134_ = D2_DC3(d_2149_v54_)
                d_2242_v135_: _dafny.Seq
                d_2242_v135_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_2241_v134_: True}), _dafny.Map({D2_DC3(d_2149_v54_): True})])
                d_2243_v136_: _dafny.MultiSet
                d_2243_v136_ = _dafny.MultiSet([d_2220_v115_])
                d_2244_v137_: _dafny.Array
                nw339_ = _dafny.Array(None, 24)
                nw339_[int(0)] = d_2240_v133_
                nw339_[int(1)] = d_2240_v133_
                nw339_[int(2)] = (d_2242_v135_)[default__.safeIndex((d_2243_v136_).cardinality, len(d_2242_v135_))]
                nw339_[int(3)] = _dafny.Map({d_2241_v134_: self.f11})
                nw339_[int(4)] = _dafny.Map({d_2241_v134_: False})
                nw339_[int(5)] = d_2240_v133_
                nw339_[int(6)] = d_2240_v133_
                nw339_[int(7)] = d_2240_v133_
                nw339_[int(8)] = (d_2240_v133_).set(D2_DC3(d_2149_v54_), self.f11)
                nw339_[int(9)] = d_2240_v133_
                nw339_[int(10)] = d_2240_v133_
                nw339_[int(11)] = d_2240_v133_
                nw339_[int(12)] = d_2240_v133_
                nw339_[int(13)] = _dafny.Map({d_2241_v134_: self.f11})
                nw339_[int(14)] = d_2240_v133_
                nw339_[int(15)] = d_2240_v133_
                nw339_[int(16)] = d_2240_v133_
                nw339_[int(17)] = d_2240_v133_
                nw339_[int(18)] = d_2240_v133_
                nw339_[int(19)] = d_2240_v133_
                nw339_[int(20)] = _dafny.Map({d_2241_v134_: self.f11})
                nw339_[int(21)] = d_2240_v133_
                nw339_[int(22)] = d_2240_v133_
                nw339_[int(23)] = d_2240_v133_
                d_2244_v137_ = nw339_
                d_2245_v138_: _dafny.Map
                d_2245_v138_ = _dafny.Map({default__.safeModuloInt((d_2222_v117_)[default__.safeIndex(574, (d_2222_v117_).length(0))], (d_2239_v132_).cf29): d_2244_v137_})
                d_2245_v138_ = (d_2245_v138_).set((d_2222_v117_)[default__.safeIndex(574, (d_2222_v117_).length(0))], d_2244_v137_)
                index399_ = default__.safeIndex(462, (d_2149_v54_).length(0))
                (d_2149_v54_)[index399_] = not(self.f11)
                d_2246_v139_: _dafny.MultiSet
                d_2246_v139_ = _dafny.MultiSet([True])
                d_2247_v140_: _dafny.Map
                d_2247_v140_ = _dafny.Map({self.f11: d_2246_v139_})
                d_2248_v141_: _dafny.MultiSet
                d_2248_v141_ = _dafny.MultiSet([d_2227_v122_])
                d_2249_v142_: D18
                d_2249_v142_ = D18_DC46(d_2248_v141_)
                d_2250_v143_: _dafny.Seq
                d_2250_v143_ = _dafny.SeqWithoutIsStrInference([((d_2249_v142_).cf89).cardinality])
                index400_ = default__.safeIndex(462, (d_2149_v54_).length(0))
                rhs363_ = (0) - ((d_2222_v117_)[default__.safeIndex(574, (d_2222_v117_).length(0))])
                rhs364_ = _dafny.Set({(((d_2247_v140_)[self.f11] if (self.f11) in (d_2247_v140_) else d_2246_v139_)).isdisjoint(_dafny.MultiSet([False])), (default__.fm63(False, (d_2225_v120_)[default__.safeIndex((d_2222_v117_)[default__.safeIndex(574, (d_2222_v117_).length(0))], len(d_2225_v120_))], globalState)) != (d_2250_v143_)})
                rhs365_ = (d_2222_v117_)[default__.safeIndex(574, (d_2222_v117_).length(0))]
                rhs366_ = self.f11
                rhs367_ = self.f11
                lhs294_ = globalState
                lhs295_ = globalState
                lhs296_ = d_2149_v54_
                lhs297_ = default__.safeIndex(462, (d_2149_v54_).length(0))
                lhs298_ = self
                lhs294_.f7 = rhs363_
                d_2223_v118_ = rhs364_
                lhs295_.f7 = rhs365_
                lhs296_[lhs297_] = rhs366_
                lhs298_.f11 = rhs367_
                d_2251_v144_: _dafny.Array
                def lambda104_(d_2252_v54_):
                    def lambda105_(d_2253_i18_):
                        return (d_2252_v54_)[default__.safeIndex(462, (d_2252_v54_).length(0))]

                    return lambda105_

                init60_ = lambda104_(d_2149_v54_)
                nw340_ = _dafny.Array(None, 10)
                for i0_60_ in range(nw340_.length(0)):
                    nw340_[i0_60_] = init60_(i0_60_)
                d_2251_v144_ = nw340_
                d_2254_v145_: D2
                d_2254_v145_ = D2_DC4(d_2222_v117_, (d_2149_v54_)[default__.safeIndex(462, (d_2149_v54_).length(0))], (d_2222_v117_)[default__.safeIndex(574, (d_2222_v117_).length(0))], False, d_2251_v144_)
                d_2255_v146_: _dafny.Array
                nw341_ = _dafny.Array(None, 8)
                nw341_[int(0)] = d_2222_v117_
                nw341_[int(1)] = d_2222_v117_
                nw341_[int(2)] = d_2222_v117_
                nw341_[int(3)] = d_2222_v117_
                nw341_[int(4)] = d_2222_v117_
                nw341_[int(5)] = (d_2254_v145_).cf6
                nw341_[int(6)] = d_2222_v117_
                nw341_[int(7)] = d_2222_v117_
                d_2255_v146_ = nw341_
                index401_ = default__.safeIndex(65, (d_2255_v146_).length(0))
                (d_2255_v146_)[index401_] = d_2222_v117_
                index402_ = default__.safeIndex(65, (d_2255_v146_).length(0))
                nw342_ = _dafny.Array(int(0), 13)
                (d_2255_v146_)[index402_] = nw342_
                index403_ = default__.safeIndex(462, (d_2149_v54_).length(0))
                (d_2149_v54_)[index403_] = (d_2225_v120_) <= (d_2225_v120_)
                d_2256_v147_: _dafny.Array
                nw343_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 8)
                d_2256_v147_ = nw343_
                d_2257_v148_: _dafny.Seq
                d_2257_v148_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "psmsf"))
                index404_ = default__.safeIndex(782, (d_2256_v147_).length(0))
                (d_2256_v147_)[index404_] = d_2257_v148_
                index405_ = default__.safeIndex(782, (d_2256_v147_).length(0))
                (d_2256_v147_)[index405_] = d_2257_v148_
            d_2258_v149_: _dafny.Array
            def lambda106_(d_2259_i19_):
                return (_dafny.Map({_dafny.MultiSet([self.f11]): self.f11})) | (_dafny.Map({_dafny.MultiSet([self.f11]): self.f11}))

            init61_ = lambda106_
            nw344_ = _dafny.Array(None, 22)
            for i0_61_ in range(nw344_.length(0)):
                nw344_[i0_61_] = init61_(i0_61_)
            d_2258_v149_ = nw344_
            d_2260_v150_: _dafny.MultiSet
            d_2260_v150_ = _dafny.MultiSet([self.f11])
            d_2261_v151_: _dafny.Map
            d_2261_v151_ = _dafny.Map({d_2260_v150_: self.f11})
            index406_ = default__.safeIndex(414, (d_2258_v149_).length(0))
            (d_2258_v149_)[index406_] = d_2261_v151_
            index407_ = default__.safeIndex(414, (d_2258_v149_).length(0))
            (d_2258_v149_)[index407_] = _dafny.Map({(_dafny.MultiSet([self.f11, not(self.f11)])) | (d_2260_v150_): self.f11})
        d_2262_v152_: str
        d_2262_v152_ = _dafny.CodePoint('k')
        d_2263_v153_: _dafny.Map
        d_2263_v153_ = _dafny.Map({self.f11: d_2262_v152_})
        d_2264_v154_: _dafny.Seq
        d_2264_v154_ = _dafny.SeqWithoutIsStrInference([p0, len(_dafny.SeqWithoutIsStrInference([d_2263_v153_]))])
        d_2265_v155_: D9
        d_2265_v155_ = D9_DC23((p0) - ((d_2264_v154_)[default__.safeIndex(691, len(d_2264_v154_))]), p0, p0, (p0) + (p0), 988)
        d_2265_v155_ = d_2265_v155_
        d_2266_v156_: _dafny.Set
        d_2266_v156_ = _dafny.Set({(0) - (p0)})
        d_2267_v157_: _dafny.MultiSet
        d_2267_v157_ = _dafny.MultiSet([p0])
        d_2268_v158_: _dafny.Seq
        d_2268_v158_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gvrrji"))
        d_2269_v159_: _dafny.Map
        d_2269_v159_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([len(d_2264_v154_)]): _dafny.Map({self.f11: self.f11})})
        d_2270_v160_: _dafny.Map
        d_2270_v160_ = _dafny.Map({self.f11: self.f11})
        d_2271_v161_: _dafny.Map
        d_2271_v161_ = _dafny.Map({len(d_2268_v158_): len(((d_2269_v159_)[d_2264_v154_] if (d_2264_v154_) in (d_2269_v159_) else d_2270_v160_))})
        d_2272_v162_: _dafny.Map
        d_2272_v162_ = _dafny.Map({_dafny.Map({p0: D1_DC2(p0, len(d_2266_v156_), self.f11)}): (d_2267_v157_).set(len(d_2268_v158_), default__.abs(((d_2271_v161_)[p0] if (p0) in (d_2271_v161_) else p0)))})
        r0 = (d_2272_v162_) | ((d_2272_v162_) | (d_2272_v162_))
        d_2273_v163_: D2
        d_2273_v163_ = D2_DC3(d_2149_v54_)
        pat_let_tv61_ = d_2149_v54_
        def iife195_(_pat_let68_0):
            def iife196_(d_2274_dt__update__tmp_h0_):
                def iife197_(_pat_let69_0):
                    def iife198_(d_2275_dt__update_hcf5_h0_):
                        return D2_DC3(d_2275_dt__update_hcf5_h0_)
                    return iife198_(_pat_let69_0)
                return iife197_(pat_let_tv61_)
            return iife196_(_pat_let68_0)
        r1 = (iife195_(d_2273_v163_) if not(self.f11) else d_2273_v163_)
        return r0, r1


class C9:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C9"
    def ctor__(self):
        pass
        pass

    def fm1(self, p0, globalState):
        source30_ = D1_DC2(len(_dafny.SeqWithoutIsStrInference([_dafny.Map({True: True}) for d_2276_i0_ in range(default__.abs(260))])), -126, False)
        if source30_.is_DC2:
            d_2277___mcc_h0_ = source30_.cf2
            d_2278___mcc_h1_ = source30_.cf3
            d_2279___mcc_h2_ = source30_.cf4
            d_2280_cf4_ = d_2279___mcc_h2_
            d_2281_cf3_ = d_2278___mcc_h1_
            d_2282_cf2_ = d_2277___mcc_h0_
            return (d_2282_cf2_) - (d_2282_cf2_)
        elif True:
            d_2283___mcc_h3_ = source30_.cf1
            d_2284_cf1_ = d_2283___mcc_h3_
            return 734

    def fm2(self, globalState):
        if (len(_dafny.Set({True}))) >= (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vahhmfdxj")))):
            return _dafny.MultiSet([_dafny.CodePoint('h')])
        elif True:
            return _dafny.MultiSet([_dafny.CodePoint('n'), _dafny.CodePoint('r')])

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: bool = False
        r3: int = int(0)
        d_2285_v0_: bool
        d_2285_v0_ = True
        d_2286_i0_: int
        d_2286_i0_ = 0
        with _dafny.label("15"):
            while d_2285_v0_:
                with _dafny.c_label("15"):
                    if (d_2286_i0_) >= (100):
                        raise _dafny.Break("15")
                    d_2286_i0_ = (d_2286_i0_) + (1)
                    index408_ = default__.safeIndex(846, (p1).length(0))
                    (p1)[index408_] = d_2285_v0_
                    index409_ = default__.safeIndex(846, (p1).length(0))
                    (p1)[index409_] = (51) < (p0)
                    r3 = p0
                    d_2287_v1_: _dafny.Array
                    def lambda107_(d_2288_p1_):
                        def lambda108_(d_2289_i1_):
                            return (D1_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hrlebc"))) if (d_2288_p1_)[default__.safeIndex(846, (d_2288_p1_).length(0))] else D1_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "houbdlof"))))

                        return lambda108_

                    init62_ = lambda107_(p1)
                    nw345_ = _dafny.Array(None, 17)
                    for i0_62_ in range(nw345_.length(0)):
                        nw345_[i0_62_] = init62_(i0_62_)
                    d_2287_v1_ = nw345_
                    d_2290_v2_: _dafny.Seq
                    d_2290_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rpfcdldk"))
                    d_2291_v3_: D1
                    d_2291_v3_ = D1_DC1(d_2290_v2_)
                    index410_ = default__.safeIndex(673, (d_2287_v1_).length(0))
                    (d_2287_v1_)[index410_] = d_2291_v3_
                    d_2292_v4_: _dafny.Map
                    d_2292_v4_ = _dafny.Map({d_2285_v0_: p0})
                    d_2293_v5_: _dafny.MultiSet
                    d_2293_v5_ = _dafny.MultiSet([(d_2292_v4_).set(default__.fm3(p0, globalState), p0)])
                    d_2294_v6_: _dafny.Array
                    def lambda109_(d_2295_p0_):
                        def lambda110_(d_2296_i2_):
                            return default__.safeModuloInt(d_2296_i2_, d_2295_p0_)

                        return lambda110_

                    init63_ = lambda109_(p0)
                    nw346_ = _dafny.Array(None, 13)
                    for i0_63_ in range(nw346_.length(0)):
                        nw346_[i0_63_] = init63_(i0_63_)
                    d_2294_v6_ = nw346_
                    d_2297_v7_: _dafny.Seq
                    d_2297_v7_ = _dafny.SeqWithoutIsStrInference([d_2294_v6_])
                    index411_ = default__.safeIndex(673, (d_2287_v1_).length(0))
                    rhs368_ = d_2291_v3_
                    rhs369_ = d_2293_v5_
                    rhs370_ = (d_2291_v3_).cf1
                    rhs371_ = ((p0) + (len(d_2290_v2_))) >= (p0)
                    rhs372_ = (d_2297_v7_)[default__.safeIndex((0) - (p0), len(d_2297_v7_))]
                    lhs299_ = d_2287_v1_
                    lhs300_ = default__.safeIndex(673, (d_2287_v1_).length(0))
                    lhs299_[lhs300_] = rhs368_
                    d_2293_v5_ = rhs369_
                    d_2290_v2_ = rhs370_
                    r2 = rhs371_
                    d_2294_v6_ = rhs372_
                    r0 = p0
                    pass
            pass
        d_2298_v8_: _dafny.Array
        nw347_ = _dafny.Array(int(0), 9)
        d_2298_v8_ = nw347_
        d_2299_v9_: D2
        d_2299_v9_ = D2_DC4(d_2298_v8_, d_2285_v0_, p0, d_2285_v0_, p1)
        pat_let_tv62_ = p0
        pat_let_tv63_ = d_2285_v0_
        pat_let_tv64_ = p1
        def iife199_(_pat_let70_0):
            def iife200_(d_2300_dt__update__tmp_h0_):
                def iife201_(_pat_let71_0):
                    def iife202_(d_2301_dt__update_hcf8_h0_):
                        def iife203_(_pat_let72_0):
                            def iife204_(d_2302_dt__update_hcf9_h0_):
                                def iife205_(_pat_let73_0):
                                    def iife206_(d_2303_dt__update_hcf10_h0_):
                                        return D2_DC4((d_2300_dt__update__tmp_h0_).cf6, (d_2300_dt__update__tmp_h0_).cf7, d_2301_dt__update_hcf8_h0_, d_2302_dt__update_hcf9_h0_, d_2303_dt__update_hcf10_h0_)
                                    return iife206_(_pat_let73_0)
                                return iife205_(pat_let_tv64_)
                            return iife204_(_pat_let72_0)
                        return iife203_(not(pat_let_tv63_))
                    return iife202_(_pat_let71_0)
                return iife201_(pat_let_tv62_)
            return iife200_(_pat_let70_0)
        d_2299_v9_ = iife199_(d_2299_v9_)
        d_2304_v10_: _dafny.Map
        d_2304_v10_ = _dafny.Map({d_2285_v0_: not (d_2285_v0_) or (d_2285_v0_)})
        d_2304_v10_ = (d_2304_v10_).set(d_2285_v0_, d_2285_v0_)
        r2 = d_2285_v0_
        d_2305_v11_: _dafny.Seq
        d_2305_v11_ = _dafny.SeqWithoutIsStrInference([p0])
        hi16_ = (d_2305_v11_)[default__.safeIndex(p0, len(d_2305_v11_))]
        for d_2306_i3_ in range(p0, hi16_):
            d_2307_v12_: _dafny.Array
            def lambda111_(d_2308_i4_):
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ouwtuvdwo"))

            init64_ = lambda111_
            nw348_ = _dafny.Array(None, 16)
            for i0_64_ in range(nw348_.length(0)):
                nw348_[i0_64_] = init64_(i0_64_)
            d_2307_v12_ = nw348_
            d_2309_v13_: _dafny.Seq
            d_2309_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mm"))
            d_2310_v14_: _dafny.Map
            d_2310_v14_ = _dafny.Map({d_2307_v12_: len((d_2309_v13_) + (d_2309_v13_))})
            d_2310_v14_ = (d_2310_v14_).set(d_2307_v12_, p0)
            pat_let_tv65_ = d_2309_v13_
            def iife207_(_pat_let74_0):
                def iife208_(d_2311_dt__update__tmp_h1_):
                    def iife209_(_pat_let75_0):
                        def iife210_(d_2312_dt__update_hcf1_h0_):
                            return D1_DC1(d_2312_dt__update_hcf1_h0_)
                        return iife210_(_pat_let75_0)
                    return iife209_(pat_let_tv65_)
                return iife208_(_pat_let74_0)
            source31_ = iife207_(D1_DC1(d_2309_v13_))
            if source31_.is_DC2:
                d_2313___mcc_h0_ = source31_.cf2
                d_2314___mcc_h1_ = source31_.cf3
                d_2315___mcc_h2_ = source31_.cf4
                d_2316_cf4_ = d_2315___mcc_h2_
                d_2317_cf3_ = d_2314___mcc_h1_
                d_2318_cf2_ = d_2313___mcc_h0_
                r3 = d_2318_cf2_
                d_2319_v15_: _dafny.MultiSet
                d_2319_v15_ = _dafny.MultiSet([931, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dhfqdpmbx")))])
                d_2320_v16_: _dafny.Map
                d_2320_v16_ = _dafny.Map({(d_2319_v15_).issubset(d_2319_v15_): p0})
                d_2320_v16_ = (d_2320_v16_).set(d_2285_v0_, d_2306_i3_)
                d_2321_v17_: D1
                d_2321_v17_ = D1_DC1(d_2309_v13_)
                d_2322_v18_: D1
                d_2322_v18_ = D1_DC2(d_2318_cf2_, d_2317_cf3_, d_2285_v0_)
                d_2323_v19_: C6
                nw349_ = C6()
                nw349_.ctor__(d_2321_v17_, d_2316_cf4_, p1, _dafny.Set({d_2317_cf3_}), d_2322_v18_, d_2316_cf4_)
                d_2323_v19_ = nw349_
                d_2324_v20_: _dafny.Seq
                d_2324_v20_ = _dafny.SeqWithoutIsStrInference([d_2323_v19_, d_2323_v19_, d_2323_v19_, d_2323_v19_])
                d_2325_v21_: _dafny.Seq
                d_2325_v21_ = _dafny.SeqWithoutIsStrInference([d_2285_v0_])
                d_2326_v22_: _dafny.Set
                d_2326_v22_ = _dafny.Set({p0, d_2318_cf2_, default__.fm0(d_2325_v21_, d_2285_v0_, True, d_2306_i3_, globalState)})
                d_2327_v23_: C2
                nw350_ = C2()
                nw350_.ctor__(len(d_2324_v20_), d_2316_cf4_, d_2321_v17_, d_2316_cf4_, p1, d_2326_v22_, d_2322_v18_, d_2316_cf4_)
                d_2327_v23_ = nw350_
                d_2328_v24_: _dafny.Set
                d_2328_v24_ = _dafny.Set({d_2327_v23_})
                d_2329_v25_: _dafny.Seq
                d_2329_v25_ = _dafny.SeqWithoutIsStrInference([d_2328_v24_])
                d_2330_v26_: str
                d_2330_v26_ = _dafny.CodePoint('l')
                d_2331_v27_: C0
                nw351_ = C0()
                nw351_.ctor__((_dafny.Set({d_2327_v23_, d_2327_v23_})).ispropersubset((d_2329_v25_)[default__.safeIndex((d_2327_v23_).f21, len(d_2329_v25_))]), d_2330_v26_, d_2322_v18_, (d_2327_v23_).f22)
                d_2331_v27_ = nw351_
                d_2332_v28_: _dafny.Seq
                d_2332_v28_ = _dafny.SeqWithoutIsStrInference([d_2305_v11_])
                d_2333_v29_: D11
                d_2333_v29_ = D11_DC26(d_2298_v8_, d_2309_v13_)
                d_2334_v30_: _dafny.Map
                d_2334_v30_ = _dafny.Map({not((d_2306_i3_) <= ((D13_DC34(d_2306_i3_, d_2332_v28_)).cf71)): d_2333_v29_})
                d_2334_v30_ = (d_2334_v30_).set(d_2285_v0_, d_2333_v29_)
            elif True:
                d_2335___mcc_h3_ = source31_.cf1
                d_2336_cf1_ = d_2335___mcc_h3_
                d_2337_v31_: _dafny.Set
                d_2337_v31_ = _dafny.Set({d_2306_i3_})
                d_2338_v32_: D1
                d_2338_v32_ = D1_DC2(d_2306_i3_, -87, d_2285_v0_)
                d_2339_v33_: T3
                nw352_ = C2()
                nw352_.ctor__(p0, d_2285_v0_, D1_DC1(d_2336_cf1_), d_2285_v0_, p1, d_2337_v31_, d_2338_v32_, d_2285_v0_)
                d_2339_v33_ = nw352_
                d_2340_v34_: _dafny.Map
                d_2340_v34_ = _dafny.Map({d_2306_i3_: d_2339_v33_})
                d_2340_v34_ = d_2340_v34_
                d_2341_v35_: C4
                nw353_ = C4()
                nw353_.ctor__()
                d_2341_v35_ = nw353_
                d_2342_v36_: _dafny.Array
                nw354_ = _dafny.Array(D3.default()(), 2)
                d_2342_v36_ = nw354_
                d_2343_v37_: D3
                d_2343_v37_ = D3_DC8(p0)
                index412_ = default__.safeIndex(637, (d_2342_v36_).length(0))
                (d_2342_v36_)[index412_] = d_2343_v37_
                index413_ = default__.safeIndex(637, (d_2342_v36_).length(0))
                (d_2342_v36_)[index413_] = d_2343_v37_
                d_2344_v38_: _dafny.MultiSet
                d_2344_v38_ = _dafny.MultiSet([len(d_2305_v11_), len(d_2337_v31_), -180, d_2306_i3_])
                d_2345_v40_: C2
                nw355_ = C2()
                def iife211_():
                    coll59_ = _dafny.Map()
                    compr_59_: int
                    for compr_59_ in _dafny.IntegerRange(186, 82):
                        d_2347_v39_: int = compr_59_
                        if ((186) <= (d_2347_v39_)) and ((d_2347_v39_) < (82)):
                            coll59_[default__.safeDivisionInt(d_2347_v39_, d_2306_i3_)] = (D15_DC39(p0, _dafny.CodePoint('q'), d_2309_v13_, p0, d_2306_i3_)).cf82
                    return _dafny.Map(coll59_)
                nw355_.ctor__(d_2306_i3_, False, d_2339_v33_.f16, not(d_2285_v0_), p1, d_2337_v31_, D1_DC2(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_2346_i5_ in range(default__.abs(873))])), (d_2344_v38_).cardinality, (d_2339_v33_).f13), (_dafny.SeqWithoutIsStrInference([len(iife211_()
                )])) == (d_2305_v11_))
                d_2345_v40_ = nw355_
                d_2348_v41_: D9
                d_2348_v41_ = D9_DC23(570, (d_2306_i3_) + (p0), (d_2345_v40_).f21, d_2306_i3_, (d_2345_v40_).f21)
                rhs373_ = d_2345_v40_
                rhs374_ = d_2348_v41_
                d_2345_v40_ = rhs373_
                d_2348_v41_ = rhs374_
            index414_ = default__.safeIndex(484, (d_2298_v8_).length(0))
            (d_2298_v8_)[index414_] = (0) - (d_2306_i3_)
            d_2349_v42_: _dafny.Map
            d_2349_v42_ = _dafny.Map({d_2285_v0_: d_2309_v13_})
            d_2350_v43_: _dafny.Map
            d_2350_v43_ = _dafny.Map({p0: ((d_2349_v42_)[not(d_2285_v0_)] if (not(d_2285_v0_)) in (d_2349_v42_) else d_2309_v13_)})
            index415_ = default__.safeIndex(484, (d_2298_v8_).length(0))
            (d_2298_v8_)[index415_] = (default__.safeModuloInt(d_2306_i3_, len(d_2350_v43_))) + (p0)
            d_2351_v44_: str
            d_2351_v44_ = _dafny.CodePoint('a')
            d_2352_v45_: D1
            d_2352_v45_ = D1_DC2(d_2306_i3_, p0, False)
            d_2353_v46_: C0
            nw356_ = C0()
            nw356_.ctor__(d_2285_v0_, d_2351_v44_, d_2352_v45_, d_2285_v0_)
            d_2353_v46_ = nw356_
            d_2354_v47_: _dafny.Map
            d_2354_v47_ = _dafny.Map({d_2353_v46_.f19: d_2353_v46_})
            d_2355_v48_: D19
            d_2355_v48_ = D19_DC48(d_2353_v46_)
            d_2356_v49_: _dafny.Array
            nw357_ = _dafny.Array(None, 25)
            nw357_[int(0)] = d_2353_v46_
            nw357_[int(1)] = d_2353_v46_
            nw357_[int(2)] = d_2353_v46_
            nw357_[int(3)] = d_2353_v46_
            nw357_[int(4)] = d_2353_v46_
            nw357_[int(5)] = d_2353_v46_
            nw357_[int(6)] = d_2353_v46_
            nw357_[int(7)] = d_2353_v46_
            nw357_[int(8)] = d_2353_v46_
            nw357_[int(9)] = d_2353_v46_
            nw357_[int(10)] = d_2353_v46_
            nw357_[int(11)] = d_2353_v46_
            nw357_[int(12)] = d_2353_v46_
            nw357_[int(13)] = d_2353_v46_
            nw357_[int(14)] = ((d_2354_v47_)[True] if (True) in (d_2354_v47_) else d_2353_v46_)
            nw357_[int(15)] = d_2353_v46_
            nw357_[int(16)] = d_2353_v46_
            nw357_[int(17)] = (d_2353_v46_ if d_2285_v0_ else d_2353_v46_)
            nw357_[int(18)] = d_2353_v46_
            nw357_[int(19)] = d_2353_v46_
            nw357_[int(20)] = d_2353_v46_
            nw357_[int(21)] = (d_2355_v48_).cf92
            nw357_[int(22)] = d_2353_v46_
            nw357_[int(23)] = d_2353_v46_
            nw357_[int(24)] = d_2353_v46_
            d_2356_v49_ = nw357_
            d_2356_v49_ = d_2356_v49_
        d_2357_v50_: C4
        nw358_ = C4()
        nw358_.ctor__()
        d_2357_v50_ = nw358_
        d_2358_v51_: _dafny.MultiSet
        d_2358_v51_ = _dafny.MultiSet([p0, p0, (0) - (p0), (0) - (p0), p0])
        r0 = (d_2305_v11_)[default__.safeIndex(((d_2358_v51_)[653] if (653) in (d_2358_v51_) else p0), len(d_2305_v11_))]
        r1 = d_2285_v0_
        r2 = ((d_2304_v10_)[True] if (True) in (d_2304_v10_) else (p0) < (940))
        d_2359_v52_: _dafny.Seq
        d_2359_v52_ = _dafny.SeqWithoutIsStrInference([d_2285_v0_])
        d_2360_v53_: _dafny.Seq
        d_2360_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sigql"))
        d_2361_v54_: _dafny.Map
        d_2361_v54_ = _dafny.Map({p0: (0) - (len(d_2360_v53_))})
        r3 = (len((d_2359_v52_) + (d_2359_v52_))) * ((p0 if d_2285_v0_ else (0) - (((d_2361_v54_)[p0] if (p0) in (d_2361_v54_) else (0) - (p0)))))
        return r0, r1, r2, r3

