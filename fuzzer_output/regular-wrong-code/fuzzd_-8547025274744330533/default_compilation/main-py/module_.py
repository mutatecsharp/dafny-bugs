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
        return ((_dafny.SeqWithoutIsStrInference([not(True), False, False, True])) + (_dafny.SeqWithoutIsStrInference([True, True, False]))) == ((_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([not(not(True))])))

    @staticmethod
    def fm1(p0, p1, p2, globalState):
        return (0) - (len(_dafny.Map({((0) - (len(_dafny.Map({True: True})))) != (589): _dafny.CodePoint('w')})))

    @staticmethod
    def fm10(p0, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(472, 729):
                d_0_v0_: int = compr_0_
                if ((472) <= (d_0_v0_)) and ((d_0_v0_) < (729)):
                    coll0_[(d_0_v0_) - (782)] = True
            return _dafny.Map(coll0_)
        return _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([-648, len(iife0_()
), 566])) for d_1_i0_ in range(default__.abs(701))])

    @staticmethod
    def fm11(p0, p1, p2, globalState):
        return D0_DC0((-796) + ((D18_DC44(False, -435, True)).cf61))

    @staticmethod
    def fm12(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([-332, -993, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([-913]))]))).cardinality])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({not(True), not(True)})), 937, len(_dafny.SeqWithoutIsStrInference([914 for d_2_i0_ in range(default__.abs(-329))])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))), 173]))

    @staticmethod
    def fm13(globalState):
        return _dafny.MultiSet([(len(_dafny.SeqWithoutIsStrInference([True]))) * (826), 58, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vb"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eyilrqfpv"))))])

    @staticmethod
    def fm14(p0, globalState):
        return _dafny.CodePoint('t')

    @staticmethod
    def fm15(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([D0_DC3(D0_DC3(D0_DC1(False))) for d_3_i0_ in range(default__.abs(538))])

    @staticmethod
    def fm16(p0, p1, globalState):
        return D0_DC3(D0_DC1(not(True)))

    @staticmethod
    def fm17(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([not(False), True])) + (_dafny.SeqWithoutIsStrInference([True]))

    @staticmethod
    def fm18(p0, globalState):
        return ((D13_DC27(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yvygfwnix")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_4_i0_ in range(default__.abs(113))])]))).cf37) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ypgkbwfy"))]))

    @staticmethod
    def fm19(p0, globalState):
        return (_dafny.Set({True})) | (_dafny.Set({False}))

    @staticmethod
    def fm20(globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.Map({True: len(_dafny.Set({-101, -828, 344, 303, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bqb"))))}))}), (_dafny.Map({True: 756})) | (_dafny.Map({not(False): -929})), _dafny.Map({False: -370}), _dafny.Map({not(False): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "shlrjox")))})])

    @staticmethod
    def fm21(p0, globalState):
        def iife1_():
            coll1_ = _dafny.Set()
            compr_1_: int
            for compr_1_ in _dafny.IntegerRange(482, 239):
                d_6_v0_: int = compr_1_
                if ((482) <= (d_6_v0_)) and ((d_6_v0_) < (239)):
                    coll1_ = coll1_.union(_dafny.Set([(d_6_v0_) * (874)]))
            return _dafny.Set(coll1_)
        return ((_dafny.Set({508, (_dafny.MultiSet([950])).cardinality})) - (_dafny.Set({len(_dafny.Map({_dafny.CodePoint('k'): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_5_i0_ in range(default__.abs(338))]))})), len(_dafny.Map({653: True}))}))) - ((iife1_()
        ) | (_dafny.Set({613})))

    @staticmethod
    def fm22(p0, globalState):
        return ((_dafny.MultiSet([D11_DC24(_dafny.Set({True}))])) | (_dafny.MultiSet([D11_DC24(_dafny.Set({True}))]))) | (_dafny.MultiSet([D11_DC24(_dafny.Set({False, False}))]))

    @staticmethod
    def fm23(p0, p1, p2, p3, globalState):
        def iife2_():
            coll2_ = _dafny.Set()
            compr_2_: D3
            for compr_2_ in (_dafny.SeqWithoutIsStrInference([D3_DC8(_dafny.Map({_dafny.SeqWithoutIsStrInference([603, len(_dafny.Set({794}))]): (_dafny.MultiSet([False, False])).cardinality}))])).Elements:
                d_7_v0_: D3 = compr_2_
                if (d_7_v0_) in (_dafny.SeqWithoutIsStrInference([D3_DC8(_dafny.Map({_dafny.SeqWithoutIsStrInference([603, len(_dafny.Set({794}))]): (_dafny.MultiSet([False, False])).cardinality}))])):
                    coll2_ = coll2_.union(_dafny.Set([d_7_v0_]))
            return _dafny.Set(coll2_)
        return (iife2_()
        ) | (_dafny.Set({D3_DC8(_dafny.Map({_dafny.SeqWithoutIsStrInference([724, -944]): (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vgs")))]))).cardinality})), D3_DC8(_dafny.Map({_dafny.SeqWithoutIsStrInference([496 for d_8_i0_ in range(default__.abs(214))]): 521}))}))

    @staticmethod
    def fm24(p0, globalState):
        if (True if False else False):
            return _dafny.CodePoint('k')
        elif True:
            return _dafny.CodePoint('x')

    @staticmethod
    def fm25(p0, p1, p2, p3, globalState):
        def iife3_():
            coll3_ = _dafny.Set()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(962, 585):
                d_9_v0_: int = compr_3_
                if ((962) <= (d_9_v0_)) and ((d_9_v0_) < (585)):
                    coll3_ = coll3_.union(_dafny.Set([(d_9_v0_) * ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_10_i0_ in range(default__.abs(-777))]))))]))
            return _dafny.Set(coll3_)
        if (iife3_()
        ).isdisjoint(_dafny.Set({(0) - (len(_dafny.Set({_dafny.SeqWithoutIsStrInference([not(False)]), _dafny.SeqWithoutIsStrInference([True])})))})):
            return (_dafny.Set({527})) | (_dafny.Set({-462}))
        elif True:
            def iife4_():
                coll4_ = _dafny.Set()
                compr_4_: int
                for compr_4_ in (_dafny.Map({243: True})).keys.Elements:
                    d_11_v1_: int = compr_4_
                    if (d_11_v1_) in (_dafny.Map({243: True})):
                        coll4_ = coll4_.union(_dafny.Set([(d_11_v1_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lrhi"))))]))
                return _dafny.Set(coll4_)
            return (_dafny.Set({438})) - (iife4_()
            )

    @staticmethod
    def fm26(globalState):
        return (_dafny.Set({False, False, not((D21_DC53(True, -400, False)).cf76)})).intersection(_dafny.Set({True}))

    @staticmethod
    def fm27(p0, p1, p2, globalState):
        def iife5_():
            coll5_ = _dafny.Map()
            compr_5_: int
            for compr_5_ in _dafny.IntegerRange(-163, 769):
                d_12_v0_: int = compr_5_
                if ((-163) <= (d_12_v0_)) and ((d_12_v0_) < (769)):
                    coll5_[(d_12_v0_) - (len(_dafny.SeqWithoutIsStrInference([653 for d_13_i0_ in range(default__.abs(398))])))] = _dafny.CodePoint('d')
            return _dafny.Map(coll5_)
        return ((_dafny.Map({_dafny.CodePoint('y'): iife5_()
        })) | (_dafny.Map({_dafny.CodePoint('l'): _dafny.Map({138: _dafny.CodePoint('x')})}))) | (_dafny.Map({_dafny.CodePoint('y'): _dafny.Map({583: _dafny.CodePoint('b')})}))

    @staticmethod
    def fm28(globalState):
        return _dafny.Set({((D6_DC13(_dafny.SeqWithoutIsStrInference([False]))).cf21) == (_dafny.SeqWithoutIsStrInference([not(False), True])), (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, False])]))).ispropersubset(_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([False, True, (D25_DC64(False, 793, _dafny.CodePoint('n'), False, len(_dafny.SeqWithoutIsStrInference([704])))).cf101]), _dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([True, not(False)]), _dafny.SeqWithoutIsStrInference([False, True])])), False})

    @staticmethod
    def fm29(p0, p1, globalState):
        def iife6_():
            coll6_ = _dafny.Map()
            compr_6_: int
            for compr_6_ in (_dafny.Set({len(_dafny.Set({False, True}))})).Elements:
                d_14_v0_: int = compr_6_
                if (d_14_v0_) in (_dafny.Set({len(_dafny.Set({False, True}))})):
                    coll6_[(d_14_v0_) - (len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([False, False]))})))] = False
            return _dafny.Map(coll6_)
        source0_ = D7_DC18(D7_DC15(iife6_()
))
        if source0_.is_DC16:
            return _dafny.CodePoint('y')
        elif source0_.is_DC17:
            d_15___mcc_h0_ = source0_.cf25
            d_16___mcc_h1_ = source0_.cf26
            d_17___mcc_h2_ = source0_.cf27
            d_18_cf27_ = d_17___mcc_h2_
            d_19_cf26_ = d_16___mcc_h1_
            d_20_cf25_ = d_15___mcc_h0_
            return _dafny.CodePoint('d')
        elif source0_.is_DC15:
            d_21___mcc_h3_ = source0_.cf24
            d_22_cf24_ = d_21___mcc_h3_
            return _dafny.CodePoint('k')
        elif True:
            d_23___mcc_h4_ = source0_.cf28
            d_24_cf28_ = d_23___mcc_h4_
            if True:
                return _dafny.CodePoint('o')
            elif True:
                return _dafny.CodePoint('w')

    @staticmethod
    def fm30(p0, p1, globalState):
        return ((_dafny.Set({D6_DC14(True, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iicnsdcau"))))}) if False else (D37_DC90(_dafny.Set({D6_DC14(True, -181), D6_DC14(False, 624)}))).cf152)) - (_dafny.Set({D6_DC14(False, len(_dafny.Map({92: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))})))}))

    @staticmethod
    def fm31(globalState):
        if (True if False else not(False)):
            return (_dafny.Map({False: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_25_i0_ in range(default__.abs(456))])})) | (_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mbjnqbggj"))}))
        elif True:
            return _dafny.Map({True: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_26_i1_ in range(default__.abs(-373))])})

    @staticmethod
    def fm32(p0, p1, p2, globalState):
        return (D6_DC13(_dafny.SeqWithoutIsStrInference([False, not(True)]))).cf21

    @staticmethod
    def fm33(p0, p1, globalState):
        return (D2_DC6(33, _dafny.SeqWithoutIsStrInference([378, ((D16_DC35(_dafny.MultiSet([len(_dafny.Set({False})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fccqm"))), 89, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cl")))]))).cf50).cardinality]), False, _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_27_i0_ in range(default__.abs(600))])), 490, 563, 270, 47]), False)).cf8

    @staticmethod
    def fm34(globalState):
        return (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([not(False), True])) + (_dafny.SeqWithoutIsStrInference([False, False, False, not(False)])))) | ((_dafny.MultiSet([not(True), False]) if True else _dafny.MultiSet([False, True])))

    @staticmethod
    def fm35(p0, p1, p2, globalState):
        return _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ddpj"))), (0) - (-679), 428})

    @staticmethod
    def fm37(p0, p1, globalState):
        def iife7_():
            coll7_ = _dafny.Map()
            compr_7_: str
            for compr_7_ in (_dafny.Set({_dafny.CodePoint('d')})).Elements:
                d_28_v0_: str = compr_7_
                if (d_28_v0_) in (_dafny.Set({_dafny.CodePoint('d')})):
                    coll7_[d_28_v0_] = (0) - (-521)
            return _dafny.Map(coll7_)
        source1_ = _dafny.MultiSet([_dafny.Map({_dafny.CodePoint('u'): 43}), iife7_()
        ])
        d_29___mcc_h0_ = source1_
        d_30_cf29_ = d_29___mcc_h0_
        if True:
            return _dafny.CodePoint('r')
        elif True:
            return _dafny.CodePoint('b')

    @staticmethod
    def fm38(p0, p1, globalState):
        return _dafny.Map({(not(True) if not(not(False)) else True): True})

    @staticmethod
    def fm39(p0, p1, p2, globalState):
        source2_ = D9_DC21(True, len(_dafny.Map({True: False})))
        if source2_.is_DC21:
            d_31___mcc_h0_ = source2_.cf31
            d_32___mcc_h1_ = source2_.cf32
            d_33_cf32_ = d_32___mcc_h1_
            d_34_cf31_ = d_31___mcc_h0_
            return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_33_cf32_]))) | (_dafny.MultiSet([d_33_cf32_, d_33_cf32_]))
        elif source2_.is_DC20:
            d_35___mcc_h2_ = source2_.cf30
            d_36_cf30_ = d_35___mcc_h2_
            if not(False):
                return _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([len((D23_DC59(_dafny.Map({True: False}))).cf90)])), 85, len(_dafny.Set({False})), 510])
            elif True:
                return _dafny.MultiSet([979, len(_dafny.SeqWithoutIsStrInference([449, len(_dafny.SeqWithoutIsStrInference([True])), len(_dafny.SeqWithoutIsStrInference([not(True)])), 823]))])
        elif True:
            d_37___mcc_h3_ = source2_.cf33
            d_38_cf33_ = d_37___mcc_h3_
            return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([782]))) - (_dafny.MultiSet([-333]))

    @staticmethod
    def fm40(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_39_i0_ in range(default__.abs(170))]))])) + ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_40_i1_ in range(default__.abs(500))])), (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "triaosgbv")))), (0) - ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qeqa")))])).cardinality)])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: 939})) for d_41_i2_ in range(default__.abs(235))])))

    @staticmethod
    def fm41(p0, globalState):
        return _dafny.Map({631: ((0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tydwljf"))), 817, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "erd")))]))) if False else -294)})

    @staticmethod
    def fm42(p0, globalState):
        def iife8_():
            coll8_ = _dafny.Set()
            compr_8_: int
            for compr_8_ in _dafny.IntegerRange(722, 174):
                d_42_v0_: int = compr_8_
                if ((722) <= (d_42_v0_)) and ((d_42_v0_) < (174)):
                    coll8_ = coll8_.union(_dafny.Set([(d_42_v0_) + (len(_dafny.Map({True: True})))]))
            return _dafny.Set(coll8_)
        return ((_dafny.Set({267})) - (iife8_()
        )) - (_dafny.Set({(0) - (-436), len(_dafny.Set({True, not(not(False)), False, not(False), True}))}))

    @staticmethod
    def fm43(p0, p1, globalState):
        def iife9_():
            coll9_ = _dafny.Set()
            compr_9_: int
            for compr_9_ in _dafny.IntegerRange(7, 204):
                d_43_v0_: int = compr_9_
                if ((7) <= (d_43_v0_)) and ((d_43_v0_) < (204)):
                    coll9_ = coll9_.union(_dafny.Set([(d_43_v0_) + (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: False}))])))]))
            return _dafny.Set(coll9_)
        def iife10_():
            coll10_ = _dafny.Set()
            compr_10_: int
            for compr_10_ in _dafny.IntegerRange(-153, 169):
                d_44_v1_: int = compr_10_
                if ((-153) <= (d_44_v1_)) and ((d_44_v1_) < (169)):
                    coll10_ = coll10_.union(_dafny.Set([default__.safeDivisionInt(d_44_v1_, -380)]))
            return _dafny.Set(coll10_)
        def iife11_():
            coll11_ = _dafny.Set()
            compr_11_: int
            for compr_11_ in _dafny.IntegerRange(127, -896):
                d_45_v2_: int = compr_11_
                if ((127) <= (d_45_v2_)) and ((d_45_v2_) < (-896)):
                    coll11_ = coll11_.union(_dafny.Set([(d_45_v2_) * (877)]))
            return _dafny.Set(coll11_)
        def iife12_():
            coll12_ = _dafny.Set()
            compr_12_: int
            for compr_12_ in _dafny.IntegerRange(592, 636):
                d_46_v3_: int = compr_12_
                if ((592) <= (d_46_v3_)) and ((d_46_v3_) < (636)):
                    coll12_ = coll12_.union(_dafny.Set([(d_46_v3_) + (len(_dafny.SeqWithoutIsStrInference([605])))]))
            return _dafny.Set(coll12_)
        return (_dafny.MultiSet([_dafny.Set({238}), iife9_()
        ])).intersection(_dafny.MultiSet([iife10_()
        , iife11_()
        , _dafny.Set({481, -552}), _dafny.Set({166}), iife12_()
        ]))

    @staticmethod
    def fm44(p0, p1, p2, p3, globalState):
        if (False) and (True):
            if True:
                return _dafny.CodePoint('c')
            elif True:
                return _dafny.CodePoint('t')
        elif True:
            return _dafny.CodePoint('o')
        elif True:
            return _dafny.CodePoint('b')

    @staticmethod
    def fm45(globalState):
        return D25_DC64((True if not(False) else not(not(True))), -250, _dafny.CodePoint('i'), False, 360)

    @staticmethod
    def fm46(p0, p1, globalState):
        source3_ = D2_DC5(_dafny.SeqWithoutIsStrInference([D0_DC3(D0_DC2(91, 261)), D0_DC3(D0_DC1(False)), D0_DC3(D0_DC1(True))]))
        if source3_.is_DC6:
            d_47___mcc_h0_ = source3_.cf7
            d_48___mcc_h1_ = source3_.cf8
            d_49___mcc_h2_ = source3_.cf9
            d_50___mcc_h3_ = source3_.cf10
            d_51___mcc_h4_ = source3_.cf11
            d_52_cf11_ = d_51___mcc_h4_
            d_53_cf10_ = d_50___mcc_h3_
            d_54_cf9_ = d_49___mcc_h2_
            d_55_cf8_ = d_48___mcc_h1_
            d_56_cf7_ = d_47___mcc_h0_
            return D0_DC2(d_56_cf7_, d_56_cf7_)
        elif source3_.is_DC5:
            d_57___mcc_h5_ = source3_.cf6
            d_58_cf6_ = d_57___mcc_h5_
            return D0_DC3(D0_DC3(D0_DC0(176)))
        elif True:
            d_59___mcc_h6_ = source3_.cf12
            d_60_cf12_ = d_59___mcc_h6_
            return D0_DC0(len(_dafny.Map({not(not(not(True))): False})))

    @staticmethod
    def fm47(p0, p1, p2, globalState):
        return D13_DC27(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))]))

    @staticmethod
    def fm48(p0, p1, p2, globalState):
        def iife13_():
            def iife15_():
                coll15_ = _dafny.Map()
                compr_15_: int
                for compr_15_ in _dafny.IntegerRange(144, 936):
                    d_61_v0_: int = compr_15_
                    if ((144) <= (d_61_v0_)) and ((d_61_v0_) < (936)):
                        coll15_[default__.safeModuloInt(d_61_v0_, len(_dafny.SeqWithoutIsStrInference([True])))] = False
                return _dafny.Map(coll15_)
            coll13_ = _dafny.Set()
            def iife14_():
                coll14_ = _dafny.Map()
                compr_14_: int
                for compr_14_ in _dafny.IntegerRange(144, 936):
                    d_61_v0_: int = compr_14_
                    if ((144) <= (d_61_v0_)) and ((d_61_v0_) < (936)):
                        coll14_[default__.safeModuloInt(d_61_v0_, len(_dafny.SeqWithoutIsStrInference([True])))] = False
                return _dafny.Map(coll14_)
            compr_13_: int
            for compr_13_ in (iife14_()
            ).keys.Elements:
                d_62_v1_: int = compr_13_
                if (d_62_v1_) in (iife15_()
                ):
                    coll13_ = coll13_.union(_dafny.Set([(d_62_v1_) + (347)]))
            return _dafny.Set(coll13_)
        return _dafny.Map({False: (0) - (len(_dafny.Map({-831: len(_dafny.SeqWithoutIsStrInference([729, len(iife13_()
        ), len(_dafny.Map({True: _dafny.Map({-318: True})})), -645]))})))})

    @staticmethod
    def fm49(p0, p1, p2, globalState):
        def iife16_():
            coll16_ = _dafny.Map()
            compr_16_: int
            for compr_16_ in _dafny.IntegerRange(-490, 374):
                d_64_v0_: int = compr_16_
                if ((-490) <= (d_64_v0_)) and ((d_64_v0_) < (374)):
                    coll16_[(d_64_v0_) * ((_dafny.MultiSet([False, False])).cardinality)] = False
            return _dafny.Map(coll16_)
        def iife17_():
            coll17_ = _dafny.Map()
            compr_17_: int
            for compr_17_ in _dafny.IntegerRange(-276, 89):
                d_65_v1_: int = compr_17_
                if ((-276) <= (d_65_v1_)) and ((d_65_v1_) < (89)):
                    coll17_[default__.safeModuloInt(d_65_v1_, (D14_DC31(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_66_i1_ in range(default__.abs(541))])))).cf43)] = False
            return _dafny.Map(coll17_)
        return ((_dafny.MultiSet([D21_DC55((0) - (len(_dafny.SeqWithoutIsStrInference([811])))), D21_DC55((_dafny.MultiSet([True])).cardinality)])).intersection(_dafny.MultiSet([D21_DC55((0) - (len(_dafny.Set({-91}))))]))) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([D21_DC55(len(_dafny.Set({(D30_DC74(False, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mqw"))): (D3_DC9(True, len(_dafny.Set({False})), len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True]) for d_63_i0_ in range(default__.abs(-393))])))).cf14})), not(True))).cf121, False}))), D21_DC55(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ubq")))), D21_DC55((0) - ((0) - ((0) - ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ft")))))))), D21_DC55(425), D21_DC55(len(_dafny.Map({len(_dafny.Set({iife16_()
, _dafny.Map({len(iife17_()
): False}), _dafny.Map({len(_dafny.Map({True: _dafny.CodePoint('q')})): False})})): 224})))])))

    @staticmethod
    def fm50(p0, p1, p2, globalState):
        return _dafny.Map({(0) - ((_dafny.MultiSet([528])).cardinality): 594})

    @staticmethod
    def fm51(p0, globalState):
        if (not(False) if True else not(True)):
            return D9_DC21(True, -998)
        elif True:
            return D9_DC21(True, 702)

    @staticmethod
    def fm52(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([True, not((_dafny.MultiSet([not(False)])).issubset(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False])))), (483) > (len(_dafny.SeqWithoutIsStrInference([not(not(not(not(False))))]))), False])

    @staticmethod
    def fm53(p0, p1, p2, p3, globalState):
        return D6_DC13(_dafny.SeqWithoutIsStrInference([False, True]))

    @staticmethod
    def fm54(p0, p1, p2, globalState):
        return _dafny.CodePoint('c')

    @staticmethod
    def fm55(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yoxdkwfwc"))))])) + (_dafny.SeqWithoutIsStrInference([-979, (_dafny.MultiSet([_dafny.Map({len(_dafny.SeqWithoutIsStrInference([406 for d_67_i0_ in range(default__.abs(951))])): _dafny.CodePoint('d')})])).cardinality]))) + ((_dafny.SeqWithoutIsStrInference([726])) + (_dafny.SeqWithoutIsStrInference([-893])))

    @staticmethod
    def fm56(p0, p1, p2, p3, globalState):
        return (_dafny.MultiSet([547, len(_dafny.SeqWithoutIsStrInference([D0_DC0(125), D0_DC0((0) - (len(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fejy")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdjamjft"))}))))])), (D18_DC43(-235)).cf59, 299])).intersection(_dafny.MultiSet([len(_dafny.Set({_dafny.CodePoint('a'), _dafny.CodePoint('t')})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uvmq"))), -153]))

    @staticmethod
    def fm57(p0, p1, p2, p3, globalState):
        def iife18_():
            coll18_ = _dafny.Map()
            compr_18_: int
            for compr_18_ in _dafny.IntegerRange(-249, 600):
                d_68_v0_: int = compr_18_
                if ((-249) <= (d_68_v0_)) and ((d_68_v0_) < (600)):
                    coll18_[(d_68_v0_) * (len(_dafny.SeqWithoutIsStrInference([True, False])))] = not(True)
            return _dafny.Map(coll18_)
        return ((_dafny.MultiSet([_dafny.Map({len(_dafny.SeqWithoutIsStrInference([532])): not(True)})])) - (_dafny.MultiSet([_dafny.Map({len(_dafny.Map({False: False})): False})]))) - (_dafny.MultiSet([_dafny.Map({311: False}), _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ghj"))): True}), _dafny.Map({608: False}), _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mwgasebhm"))): not(not(True))}), iife18_()
        ]))

    @staticmethod
    def fm58(p0, p1, p2, p3, globalState):
        if True:
            return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([not(False)]), _dafny.SeqWithoutIsStrInference([False])]))
        elif True:
            return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, not(True)])])

    @staticmethod
    def fm59(p0, p1, globalState):
        return D18_DC44((_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({True: True})))])) < (_dafny.SeqWithoutIsStrInference([27])), (111) - (len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_69_i0_ in range(default__.abs(744))])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xbknbs")))}))), (True) and (False))

    @staticmethod
    def fm60(p0, p1, globalState):
        if (_dafny.MultiSet([not(not(True)), not(True)])).ispropersubset(_dafny.MultiSet([False, True])):
            return (_dafny.Map({False: False})) | (_dafny.Map({not(not(not(not(False)))): not(True)}))
        elif True:
            return _dafny.Map({(D23_DC60(True, _dafny.CodePoint('r'))).cf91: False})

    @staticmethod
    def fm61(globalState):
        return _dafny.Map({False: (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kbjosn"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tnxxf")))})

    @staticmethod
    def fm62(p0, p1, p2, globalState):
        def iife19_():
            coll19_ = _dafny.Map()
            compr_19_: int
            for compr_19_ in (_dafny.Map({595: False})).keys.Elements:
                d_70_v0_: int = compr_19_
                if (d_70_v0_) in (_dafny.Map({595: False})):
                    coll19_[default__.safeModuloInt(d_70_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))))] = len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([not(True), False])): False}))
            return _dafny.Map(coll19_)
        return D21_DC53((_dafny.MultiSet([(_dafny.MultiSet([not(True), True, False])).cardinality])).isdisjoint(_dafny.MultiSet([(0) - (len(iife19_()
)), 237, 843, 283])), len((_dafny.Map({not(True): 836})) | (_dafny.Map({False: 539}))), False)

    @staticmethod
    def fm63(p0, p1, globalState):
        return D18_DC42(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lfutyswc")), False)

    @staticmethod
    def fm64(p0, p1, globalState):
        def iife20_():
            coll20_ = _dafny.Set()
            compr_20_: str
            for compr_20_ in (_dafny.Set({_dafny.CodePoint('j')})).Elements:
                d_71_v0_: str = compr_20_
                if (d_71_v0_) in (_dafny.Set({_dafny.CodePoint('j')})):
                    coll20_ = coll20_.union(_dafny.Set([d_71_v0_]))
            return _dafny.Set(coll20_)
        def iife21_():
            coll21_ = _dafny.Map()
            compr_21_: str
            for compr_21_ in (_dafny.Set({_dafny.CodePoint('b'), _dafny.CodePoint('y'), _dafny.CodePoint('n'), _dafny.CodePoint('e')})).Elements:
                d_72_v1_: str = compr_21_
                if (d_72_v1_) in (_dafny.Set({_dafny.CodePoint('b'), _dafny.CodePoint('y'), _dafny.CodePoint('n'), _dafny.CodePoint('e')})):
                    coll21_[d_72_v1_] = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ibeieeiw"))), -394])
            return _dafny.Map(coll21_)
        def iife22_():
            coll22_ = _dafny.Map()
            compr_22_: int
            for compr_22_ in _dafny.IntegerRange(506, 931):
                d_73_v2_: int = compr_22_
                if ((506) <= (d_73_v2_)) and ((d_73_v2_) < (931)):
                    coll22_[default__.safeModuloInt(d_73_v2_, 49)] = False
            return _dafny.Map(coll22_)
        return ((_dafny.Map({_dafny.CodePoint('b'): _dafny.SeqWithoutIsStrInference([-420, len(iife20_()
        )])})) | (iife21_()
        )) | (_dafny.Map({_dafny.CodePoint('p'): _dafny.SeqWithoutIsStrInference([363, len(iife22_()
        ), 729])}))

    @staticmethod
    def fm65(p0, p1, p2, globalState):
        source4_ = D21_DC53(not(True), 7, True)
        if source4_.is_DC53:
            d_74___mcc_h0_ = source4_.cf76
            d_75___mcc_h1_ = source4_.cf77
            d_76___mcc_h2_ = source4_.cf78
            d_77_cf78_ = d_76___mcc_h2_
            d_78_cf77_ = d_75___mcc_h1_
            d_79_cf76_ = d_74___mcc_h0_
            return D13_DC28(d_78_cf77_)
        elif source4_.is_DC54:
            d_80___mcc_h3_ = source4_.cf79
            d_81___mcc_h4_ = source4_.cf80
            d_82___mcc_h5_ = source4_.cf81
            d_83_cf81_ = d_82___mcc_h5_
            d_84_cf80_ = d_81___mcc_h4_
            d_85_cf79_ = d_80___mcc_h3_
            return D13_DC28(d_84_cf80_)
        elif source4_.is_DC55:
            d_86___mcc_h6_ = source4_.cf82
            d_87_cf82_ = d_86___mcc_h6_
            if not(True):
                return D13_DC28((0) - (d_87_cf82_))
            elif True:
                return D13_DC28(d_87_cf82_)
        elif True:
            d_88___mcc_h7_ = source4_.cf75
            d_89_cf75_ = d_88___mcc_h7_
            return D13_DC28(633)

    @staticmethod
    def fm66(p0, p1, globalState):
        def iife23_():
            coll23_ = _dafny.Map()
            compr_23_: str
            for compr_23_ in (_dafny.Map({_dafny.CodePoint('h'): True})).keys.Elements:
                d_90_v0_: str = compr_23_
                if (d_90_v0_) in (_dafny.Map({_dafny.CodePoint('h'): True})):
                    coll23_[d_90_v0_] = len(_dafny.SeqWithoutIsStrInference([False, not(False)]))
            return _dafny.Map(coll23_)
        return (iife23_()
        ) | ((_dafny.Map({_dafny.CodePoint('t'): len(_dafny.SeqWithoutIsStrInference([False]))}) if False else _dafny.Map({_dafny.CodePoint('x'): -413})))

    @staticmethod
    def fm67(p0, globalState):
        def iife24_():
            coll24_ = _dafny.Map()
            compr_24_: int
            for compr_24_ in _dafny.IntegerRange(564, 177):
                d_92_v0_: int = compr_24_
                if ((564) <= (d_92_v0_)) and ((d_92_v0_) < (177)):
                    coll24_[default__.safeModuloInt(d_92_v0_, 282)] = -299
            return _dafny.Map(coll24_)
        return _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([354, len(_dafny.Set({549, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_91_i1_ in range(default__.abs(626))]))})), 323, len(_dafny.Set({True}))])) + (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({893: len(_dafny.Map({_dafny.CodePoint('m'): len(iife24_()
)}))}))) for d_93_i2_ in range(default__.abs(590))])) for d_94_i0_ in range(default__.abs(-510))])

    @staticmethod
    def fm68(p0, p1, p2, p3, globalState):
        source5_ = D34_DC83(False)
        if source5_.is_DC83:
            d_95___mcc_h0_ = source5_.cf136
            d_96_cf136_ = d_95___mcc_h0_
            return D22_DC58(d_96_cf136_, 43, 995)
        elif source5_.is_DC84:
            d_97___mcc_h1_ = source5_.cf137
            d_98___mcc_h2_ = source5_.cf138
            d_99___mcc_h3_ = source5_.cf139
            d_100___mcc_h4_ = source5_.cf140
            d_101_cf140_ = d_100___mcc_h4_
            d_102_cf139_ = d_99___mcc_h3_
            d_103_cf138_ = d_98___mcc_h2_
            d_104_cf137_ = d_97___mcc_h1_
            return D22_DC58(True, (_dafny.MultiSet([True])).cardinality, d_104_cf137_)
        elif source5_.is_DC85:
            d_105___mcc_h5_ = source5_.cf141
            d_106___mcc_h6_ = source5_.cf142
            d_107___mcc_h7_ = source5_.cf143
            d_108_cf143_ = d_107___mcc_h7_
            d_109_cf142_ = d_106___mcc_h6_
            d_110_cf141_ = d_105___mcc_h5_
            return D22_DC58(d_109_cf142_, d_108_cf143_, d_108_cf143_)
        elif True:
            d_111___mcc_h8_ = source5_.cf135
            d_112_cf135_ = d_111___mcc_h8_
            return D22_DC58((d_112_cf135_).f34, len(_dafny.Set({-120})), len(_dafny.SeqWithoutIsStrInference([(d_112_cf135_).f34, (d_112_cf135_).f34])))

    @staticmethod
    def fm69(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.Map({716: len(_dafny.Map({False: 828}))}) for d_113_i0_ in range(default__.abs(647))])

    @staticmethod
    def fm70(globalState):
        def iife25_():
            coll25_ = _dafny.Map()
            compr_25_: int
            for compr_25_ in _dafny.IntegerRange(113, 619):
                d_114_v0_: int = compr_25_
                if ((113) <= (d_114_v0_)) and ((d_114_v0_) < (619)):
                    coll25_[(d_114_v0_) - (402)] = len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_115_i0_ in range(default__.abs(518))]))]))])): True}))
            return _dafny.Map(coll25_)
        def iife26_():
            coll26_ = _dafny.Set()
            compr_26_: int
            for compr_26_ in _dafny.IntegerRange(355, 863):
                d_121_v1_: int = compr_26_
                if ((355) <= (d_121_v1_)) and ((d_121_v1_) < (863)):
                    coll26_ = coll26_.union(_dafny.Set([(d_121_v1_) * (len(_dafny.Map({True: not(False)})))]))
            return _dafny.Set(coll26_)
        return D15_DC34(-417, (_dafny.Set({iife25_()
})).ispropersubset(_dafny.Set({_dafny.Map({len(_dafny.Set({len(_dafny.Map({True: 772})), len(_dafny.SeqWithoutIsStrInference([(D7_DC17(len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([-230 for d_116_i1_ in range(default__.abs(789))]))])), len(_dafny.Map({False: (_dafny.MultiSet([_dafny.CodePoint('j'), _dafny.CodePoint('u')])).cardinality})), False)).cf27])), len(_dafny.Map({-138: _dafny.CodePoint('t')}))})): 662})})), (0) - (len(_dafny.SeqWithoutIsStrInference([False, True, True, not(True)]))), default__.safeDivisionInt((0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: 610})), len(_dafny.SeqWithoutIsStrInference([D2_DC6(355, _dafny.SeqWithoutIsStrInference([703]), True, _dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_117_i3_ in range(default__.abs(483))]))})) for d_118_i2_ in range(default__.abs(121))]), False), D2_DC6(44, _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tmoernb"))) for d_119_i4_ in range(default__.abs(456))]), False, _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([False])).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wod")))]), not(False)), D2_DC6(405, _dafny.SeqWithoutIsStrInference([-354]), True, _dafny.SeqWithoutIsStrInference([494 for d_120_i5_ in range(default__.abs(-696))]), not(False)), D2_DC6(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fwgtmvno"))), _dafny.SeqWithoutIsStrInference([-364, -47]), False, _dafny.SeqWithoutIsStrInference([-744, (_dafny.MultiSet([2])).cardinality]), False)])), 31]))), len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))): False}))), (_dafny.Set({690, 14})).ispropersubset(iife26_()
))

    @staticmethod
    def fm71(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([D6_DC13(_dafny.SeqWithoutIsStrInference([False])) for d_122_i0_ in range(default__.abs(480))])) + (_dafny.SeqWithoutIsStrInference([D6_DC13(_dafny.SeqWithoutIsStrInference([False]))]))) + ((_dafny.SeqWithoutIsStrInference([D6_DC13(_dafny.SeqWithoutIsStrInference([True, True])), D6_DC13(_dafny.SeqWithoutIsStrInference([False])), D6_DC13(_dafny.SeqWithoutIsStrInference([False, True, True, not(not(False)), False])), D6_DC13(_dafny.SeqWithoutIsStrInference([False, True]))])) + (_dafny.SeqWithoutIsStrInference([D6_DC13(_dafny.SeqWithoutIsStrInference([True])), D6_DC13(_dafny.SeqWithoutIsStrInference([True])), D6_DC13(_dafny.SeqWithoutIsStrInference([False, False])), D6_DC13(_dafny.SeqWithoutIsStrInference([not(True), False]))])))

    @staticmethod
    def fm72(p0, p1, p2, p3, globalState):
        def iife27_():
            def iife29_():
                coll29_ = _dafny.Map()
                compr_29_: int
                for compr_29_ in _dafny.IntegerRange(339, 327):
                    d_123_v0_: int = compr_29_
                    if ((339) <= (d_123_v0_)) and ((d_123_v0_) < (327)):
                        coll29_[default__.safeDivisionInt(d_123_v0_, len(_dafny.Map({not(False): 464})))] = 405
                return _dafny.Map(coll29_)
            coll27_ = _dafny.Set()
            def iife28_():
                coll28_ = _dafny.Map()
                compr_28_: int
                for compr_28_ in _dafny.IntegerRange(339, 327):
                    d_123_v0_: int = compr_28_
                    if ((339) <= (d_123_v0_)) and ((d_123_v0_) < (327)):
                        coll28_[default__.safeDivisionInt(d_123_v0_, len(_dafny.Map({not(False): 464})))] = 405
                return _dafny.Map(coll28_)
            compr_27_: int
            for compr_27_ in (iife28_()
            ).keys.Elements:
                d_124_v1_: int = compr_27_
                if (d_124_v1_) in (iife29_()
                ):
                    coll27_ = coll27_.union(_dafny.Set([(d_124_v1_) + (60)]))
            return _dafny.Set(coll27_)
        return D2_DC6(default__.safeDivisionInt(len(_dafny.Set({False})), len(_dafny.SeqWithoutIsStrInference([len(iife27_()
) for d_125_i0_ in range(default__.abs(-305))]))), (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({_dafny.CodePoint('f')})), len(_dafny.SeqWithoutIsStrInference([858, 9]))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))).cardinality])) for d_126_i1_ in range(default__.abs(-811))])), (_dafny.SeqWithoutIsStrInference([False])) == (_dafny.SeqWithoutIsStrInference([False, False])), _dafny.SeqWithoutIsStrInference([290, (_dafny.MultiSet([True])).cardinality, 293]), True)

    @staticmethod
    def fm73(p0, p1, p2, globalState):
        source6_ = D7_DC16()
        if source6_.is_DC16:
            return (_dafny.MultiSet([_dafny.CodePoint('n')])) - (_dafny.MultiSet([_dafny.CodePoint('h'), _dafny.CodePoint('l')]))
        elif source6_.is_DC17:
            d_127___mcc_h0_ = source6_.cf25
            d_128___mcc_h1_ = source6_.cf26
            d_129___mcc_h2_ = source6_.cf27
            d_130_cf27_ = d_129___mcc_h2_
            d_131_cf26_ = d_128___mcc_h1_
            d_132_cf25_ = d_127___mcc_h0_
            return _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o'), _dafny.CodePoint('t')])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v')])))
        elif source6_.is_DC15:
            d_133___mcc_h3_ = source6_.cf24
            d_134_cf24_ = d_133___mcc_h3_
            return _dafny.MultiSet([_dafny.CodePoint('t'), _dafny.CodePoint('j')])
        elif True:
            d_135___mcc_h4_ = source6_.cf28
            d_136_cf28_ = d_135___mcc_h4_
            return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q')]))).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x'), _dafny.CodePoint('n')])))

    @staticmethod
    def fm74(p0, p1, p2, p3, globalState):
        return D23_DC60(not(False), (_dafny.CodePoint('p') if False else _dafny.CodePoint('h')))

    @staticmethod
    def fm75(p0, p1, p2, p3, globalState):
        if not(False):
            return D22_DC56(_dafny.Map({-225: _dafny.CodePoint('g')}))
        elif True:
            def iife30_():
                coll30_ = _dafny.Map()
                compr_30_: int
                for compr_30_ in (_dafny.Set({(0) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: False})), (0) - (len(_dafny.SeqWithoutIsStrInference([False])))]))).cardinality)})).Elements:
                    d_137_v0_: int = compr_30_
                    if (d_137_v0_) in (_dafny.Set({(0) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: False})), (0) - (len(_dafny.SeqWithoutIsStrInference([False])))]))).cardinality)})):
                        coll30_[(d_137_v0_) + (-579)] = _dafny.CodePoint('t')
                return _dafny.Map(coll30_)
            return D22_DC56(iife30_()
)

    @staticmethod
    def m0(p0, p1, globalState):
        r0: bool = False
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_138_v0_: bool
        d_138_v0_ = True
        d_139_v1_: C1
        nw0_ = C1()
        nw0_.ctor__(d_138_v0_)
        d_139_v1_ = nw0_
        d_140_v2_: int
        d_140_v2_ = 990
        d_141_v3_: _dafny.Seq
        d_141_v3_ = _dafny.SeqWithoutIsStrInference([d_140_v2_])
        hi0_ = (d_141_v3_)[default__.safeIndex(d_140_v2_, len(d_141_v3_))]
        for d_142_i0_ in range(d_140_v2_, hi0_):
            d_143_v4_: _dafny.Array
            def lambda0_(d_144_i0_):
                def lambda1_(d_145_i1_):
                    return (d_145_i1_) + (d_144_i0_)

                return lambda1_

            init0_ = lambda0_(d_142_i0_)
            nw1_ = _dafny.Array(None, 22)
            for i0_0_ in range(nw1_.length(0)):
                nw1_[i0_0_] = init0_(i0_0_)
            d_143_v4_ = nw1_
            d_146_v5_: _dafny.Array
            d_146_v5_ = d_143_v4_
            d_146_v5_ = d_143_v4_
            d_147_v6_: _dafny.MultiSet
            d_147_v6_ = _dafny.MultiSet([d_138_v0_, d_138_v0_, d_138_v0_, d_138_v0_])
            d_148_v7_: _dafny.Set
            d_148_v7_ = _dafny.Set({default__.safeModuloInt(d_142_i0_, -907), d_142_i0_, (-446) - ((d_147_v6_).cardinality)})
            d_148_v7_ = ((d_148_v7_) - (d_148_v7_)).intersection(d_148_v7_)
            d_149_v8_: _dafny.Array
            nw2_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 20)
            d_149_v8_ = nw2_
            d_150_v9_: _dafny.Seq
            d_150_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "faxgwscwa"))
            index0_ = default__.safeIndex(680, (d_149_v8_).length(0))
            (d_149_v8_)[index0_] = d_150_v9_
            index1_ = default__.safeIndex(680, (d_149_v8_).length(0))
            (d_149_v8_)[index1_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_151_i2_ in range(default__.abs(975))])
            d_152_v10_: str
            d_152_v10_ = _dafny.CodePoint('m')
            d_153_v11_: _dafny.Map
            d_153_v11_ = _dafny.Map({d_138_v0_: d_152_v10_})
            d_154_v12_: D9
            d_154_v12_ = D9_DC20(d_153_v11_)
            d_155_v13_: D9
            d_155_v13_ = D9_DC22(d_154_v12_)
            d_155_v13_ = d_155_v13_
        d_156_v14_: _dafny.Seq
        d_156_v14_ = _dafny.SeqWithoutIsStrInference([True])
        d_157_v15_: _dafny.MultiSet
        d_157_v15_ = _dafny.MultiSet([d_156_v14_])
        d_158_v16_: str
        d_158_v16_ = _dafny.CodePoint('w')
        d_159_v17_: _dafny.Map
        d_159_v17_ = _dafny.Map({(_dafny.MultiSet([True])).cardinality: (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cntdgr"))).set(default__.safeIndex(d_140_v2_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cntdgr")))), d_158_v16_)})
        d_160_v18_: _dafny.Seq
        d_160_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ah"))
        (globalState).f2 = ((d_157_v15_)[default__.fm17(d_140_v2_, ((d_159_v17_)[len(d_160_v18_)] if (len(d_160_v18_)) in (d_159_v17_) else _dafny.SeqWithoutIsStrInference([d_158_v16_ for d_161_i3_ in range(default__.abs(391))])), len(_dafny.SeqWithoutIsStrInference([d_138_v0_, d_138_v0_])), globalState)] if (default__.fm17(d_140_v2_, ((d_159_v17_)[len(d_160_v18_)] if (len(d_160_v18_)) in (d_159_v17_) else _dafny.SeqWithoutIsStrInference([d_158_v16_ for d_162_i3_ in range(default__.abs(391))])), len(_dafny.SeqWithoutIsStrInference([d_138_v0_, d_138_v0_])), globalState)) in (d_157_v15_) else d_140_v2_)
        hi1_ = default__.fm1(False, d_156_v14_, True, globalState)
        for d_163_i4_ in range(default__.safeDivisionInt(223, d_140_v2_), hi1_):
            d_164_v19_: _dafny.Array
            nw3_ = _dafny.Array(False, 29)
            d_164_v19_ = nw3_
            index2_ = default__.safeIndex(252, (d_164_v19_).length(0))
            (d_164_v19_)[index2_] = not(d_138_v0_)
            index3_ = default__.safeIndex(252, (d_164_v19_).length(0))
            (d_164_v19_)[index3_] = (d_139_v1_).fm4(D0_DC1(d_138_v0_), globalState)
            d_158_v16_ = d_158_v16_
            d_165_v20_: D11
            d_165_v20_ = D11_DC24(_dafny.Set({d_138_v0_}))
            d_166_v21_: _dafny.Set
            d_166_v21_ = _dafny.Set({(d_164_v19_)[default__.safeIndex(252, (d_164_v19_).length(0))]})
            pat_let_tv0_ = d_166_v21_
            def iife31_(_pat_let0_0):
                def iife32_(d_167_dt__update__tmp_h2_):
                    def iife33_(_pat_let1_0):
                        def iife34_(d_168_dt__update_hcf35_h0_):
                            return D11_DC24(d_168_dt__update_hcf35_h0_)
                        return iife34_(_pat_let1_0)
                    return iife33_(pat_let_tv0_)
                return iife32_(_pat_let0_0)
            source7_ = iife31_(d_165_v20_)
            if source7_.is_DC25:
                (globalState).f5 = (d_156_v14_)[default__.safeIndex(d_140_v2_, len(d_156_v14_))]
                (globalState).f2 = (0) - (d_163_i4_)
                d_169_v22_: _dafny.Set
                d_169_v22_ = _dafny.Set({d_163_i4_, d_140_v2_})
                d_170_v23_: _dafny.Set
                d_170_v23_ = _dafny.Set({(d_169_v22_) - (d_169_v22_)})
                d_170_v23_ = d_170_v23_
                (globalState).f2 = d_140_v2_
            elif True:
                d_171___mcc_h0_ = source7_.cf35
                d_172_cf35_ = d_171___mcc_h0_
                index4_ = default__.safeIndex(252, (d_164_v19_).length(0))
                (d_164_v19_)[index4_] = d_138_v0_
                (globalState).f5 = d_138_v0_
                d_173_v24_: _dafny.MultiSet
                d_173_v24_ = _dafny.MultiSet([(d_164_v19_)[default__.safeIndex(252, (d_164_v19_).length(0))], d_138_v0_])
                d_174_v25_: C4
                nw4_ = C4()
                nw4_.ctor__((d_173_v24_).set((d_164_v19_)[default__.safeIndex(252, (d_164_v19_).length(0))], default__.abs(len(_dafny.Map({d_138_v0_: 20})))), d_138_v0_)
                d_174_v25_ = nw4_
                d_175_v26_: _dafny.Map
                d_175_v26_ = _dafny.Map({d_138_v0_: d_158_v16_})
                d_176_v27_: T1
                nw5_ = C6()
                nw5_.ctor__(d_163_i4_, d_140_v2_, False)
                d_176_v27_ = nw5_
                d_177_v28_: _dafny.Seq
                d_177_v28_ = _dafny.SeqWithoutIsStrInference([d_176_v27_])
                d_178_v29_: _dafny.Seq
                d_178_v29_ = _dafny.SeqWithoutIsStrInference([d_176_v27_, (d_177_v28_)[default__.safeIndex((d_176_v27_).f29, len(d_177_v28_))]])
                d_179_v30_: _dafny.Map
                d_179_v30_ = _dafny.Map({(d_178_v29_)[default__.safeIndex((d_176_v27_).f28, len(d_178_v29_))]: d_175_v26_})
                d_175_v26_ = (((d_179_v30_)[d_176_v27_] if (d_176_v27_) in (d_179_v30_) else d_175_v26_)) | ((d_175_v26_) | (_dafny.Map({True: d_158_v16_})))
            if (d_139_v1_).fm4(p1, globalState):
                d_180_v31_: T1
                nw6_ = C5()
                nw6_.ctor__((d_164_v19_)[default__.safeIndex(252, (d_164_v19_).length(0))], d_140_v2_, d_163_i4_, d_138_v0_)
                d_180_v31_ = nw6_
                d_181_v32_: _dafny.Array
                nw7_ = _dafny.Array(int(0), 25)
                d_181_v32_ = nw7_
                d_182_v33_: _dafny.Map
                d_182_v33_ = _dafny.Map({(d_180_v31_).f29: (d_180_v31_).f27})
                nw8_ = C9()
                nw8_.ctor__(d_181_v32_, (d_140_v2_) * ((d_180_v31_).f28), len((d_182_v33_) | (d_182_v33_)), (986) - ((d_180_v31_).f28), not(d_138_v0_))
                d_180_v31_ = nw8_
                index5_ = default__.safeIndex(252, (d_164_v19_).length(0))
                (d_164_v19_)[index5_] = (d_180_v31_).f27
                d_183_v34_: _dafny.Array
                nw9_ = _dafny.Array(None, 22)
                nw9_[int(0)] = d_158_v16_
                nw9_[int(1)] = d_158_v16_
                nw9_[int(2)] = d_158_v16_
                nw9_[int(3)] = d_158_v16_
                nw9_[int(4)] = d_158_v16_
                nw9_[int(5)] = d_158_v16_
                nw9_[int(6)] = (d_160_v18_)[default__.safeIndex((d_180_v31_).f29, len(d_160_v18_))]
                nw9_[int(7)] = default__.fm14(len(d_160_v18_), globalState)
                nw9_[int(8)] = d_158_v16_
                nw9_[int(9)] = d_158_v16_
                nw9_[int(10)] = d_158_v16_
                nw9_[int(11)] = d_158_v16_
                nw9_[int(12)] = d_158_v16_
                nw9_[int(13)] = d_158_v16_
                nw9_[int(14)] = d_158_v16_
                nw9_[int(15)] = d_158_v16_
                nw9_[int(16)] = d_158_v16_
                nw9_[int(17)] = _dafny.CodePoint('k')
                nw9_[int(18)] = d_158_v16_
                nw9_[int(19)] = d_158_v16_
                nw9_[int(20)] = d_158_v16_
                nw9_[int(21)] = d_158_v16_
                d_183_v34_ = nw9_
                index6_ = default__.safeIndex(562, (d_181_v32_).length(0))
                (d_181_v32_)[index6_] = (d_180_v31_).f29
                d_184_v35_: _dafny.Set
                d_184_v35_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bxnlycd"))), 691})
                index7_ = default__.safeIndex(562, (d_181_v32_).length(0))
                rhs0_ = d_183_v34_
                rhs1_ = ((d_180_v31_).f29 if (d_160_v18_) <= ((d_160_v18_).set(default__.safeIndex((d_180_v31_).f29, len(d_160_v18_)), d_158_v16_)) else d_163_i4_)
                rhs2_ = d_180_v31_
                rhs3_ = d_158_v16_
                rhs4_ = default__.fm0(d_160_v18_, (d_180_v31_).f27, ((d_182_v33_)[d_163_i4_] if (d_163_i4_) in (d_182_v33_) else True), d_184_v35_, globalState)
                lhs0_ = d_181_v32_
                lhs1_ = default__.safeIndex(562, (d_181_v32_).length(0))
                lhs2_ = globalState
                d_183_v34_ = rhs0_
                lhs0_[lhs1_] = rhs1_
                d_180_v31_ = rhs2_
                d_158_v16_ = rhs3_
                lhs2_.f5 = rhs4_
                d_185_v36_: D21
                d_185_v36_ = D21_DC53(default__.fm0(_dafny.SeqWithoutIsStrInference([d_158_v16_ for d_186_i5_ in range(default__.abs(294))]), (d_180_v31_).f27, (d_180_v31_).f27, d_184_v35_, globalState), (d_180_v31_).f28, (d_164_v19_)[default__.safeIndex(252, (d_164_v19_).length(0))])
                d_187_v37_: _dafny.MultiSet
                d_187_v37_ = _dafny.MultiSet([(d_180_v31_).f27])
                d_156_v14_ = (default__.fm52((d_185_v36_).cf77, (d_180_v31_).f27, d_187_v37_, globalState)).set(default__.safeIndex(829, len(default__.fm52((d_185_v36_).cf77, (d_180_v31_).f27, d_187_v37_, globalState))), (d_180_v31_).f27)
                d_188_v38_: C9
                nw10_ = C9()
                nw10_.ctor__(d_181_v32_, (0) - ((d_180_v31_).f28), ((0) - ((d_180_v31_).f28)) - (d_140_v2_), (d_181_v32_)[default__.safeIndex(562, (d_181_v32_).length(0))], d_138_v0_)
                d_188_v38_ = nw10_
                d_188_v38_ = d_188_v38_
            elif True:
                d_189_v39_: C6
                nw11_ = C6()
                nw11_.ctor__(d_163_i4_, len((d_141_v3_) + (d_141_v3_)), not((d_164_v19_)[default__.safeIndex(252, (d_164_v19_).length(0))]))
                d_189_v39_ = nw11_
                d_190_v40_: _dafny.MultiSet
                d_190_v40_ = _dafny.MultiSet([(d_164_v19_)[default__.safeIndex(252, (d_164_v19_).length(0))]])
                d_191_v41_: T0
                nw12_ = C4()
                nw12_.ctor__(d_190_v40_, d_138_v0_)
                d_191_v41_ = nw12_
                d_191_v41_ = d_191_v41_
                d_192_v42_: C0
                nw13_ = C0()
                nw13_.ctor__()
                d_192_v42_ = nw13_
                (globalState).f25 = not((d_191_v41_).f27)
                (globalState).f2 = (0) - (default__.fm1((d_164_v19_)[default__.safeIndex(252, (d_164_v19_).length(0))], (d_156_v14_).set(default__.safeIndex(d_140_v2_, len(d_156_v14_)), (d_164_v19_)[default__.safeIndex(252, (d_164_v19_).length(0))]), True, globalState))
        d_193_v43_: _dafny.Array
        def lambda2_(d_194_v0_):
            def lambda3_(d_195_i7_):
                return (_dafny.Set({d_194_v0_, True})) | (_dafny.Set({d_194_v0_, d_194_v0_}))

            return lambda3_

        init1_ = lambda2_(d_138_v0_)
        nw14_ = _dafny.Array(None, 21)
        for i0_1_ in range(nw14_.length(0)):
            nw14_[i0_1_] = init1_(i0_1_)
        d_193_v43_ = nw14_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_193_v43_).length(0)):
            d_196_i6_: int = guard_loop_0_
            if (True) and (((0) <= (d_196_i6_)) and ((d_196_i6_) < ((d_193_v43_).length(0)))):
                (d_193_v43_)[(d_196_i6_)] = ((_dafny.Set({not(False), d_138_v0_, (d_156_v14_)[default__.safeIndex(786, len(d_156_v14_))]})) | (_dafny.Set({d_138_v0_, not(d_138_v0_), True})) if (True) == (d_138_v0_) else (_dafny.Set({d_138_v0_})) | (_dafny.Set({d_138_v0_})))
        d_197_v44_: _dafny.MultiSet
        d_197_v44_ = _dafny.MultiSet([d_138_v0_, d_138_v0_])
        d_198_v45_: _dafny.Set
        d_198_v45_ = _dafny.Set({(d_197_v44_).cardinality, 777})
        d_199_v46_: D25
        d_199_v46_ = D25_DC63(_dafny.Map({d_160_v18_: d_198_v45_}))
        d_200_v47_: _dafny.Map
        d_200_v47_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")): d_198_v45_})
        d_201_v49_: _dafny.Seq
        d_201_v49_ = _dafny.SeqWithoutIsStrInference([d_160_v18_, d_160_v18_])
        pat_let_tv1_ = d_200_v47_
        d_202_v50_: _dafny.Array
        nw15_ = _dafny.Array(None, 16)
        nw15_[int(0)] = d_199_v46_
        nw15_[int(1)] = D25_DC63(d_200_v47_)
        nw15_[int(2)] = d_199_v46_
        nw15_[int(3)] = d_199_v46_
        nw15_[int(4)] = d_199_v46_
        nw15_[int(5)] = d_199_v46_
        nw15_[int(6)] = d_199_v46_
        def iife35_():
            coll31_ = _dafny.Map()
            compr_31_: _dafny.Seq
            for compr_31_ in (d_201_v49_).Elements:
                d_203_v48_: _dafny.Seq = compr_31_
                if (d_203_v48_) in (d_201_v49_):
                    coll31_[d_203_v48_] = d_198_v45_
            return _dafny.Map(coll31_)
        nw15_[int(7)] = D25_DC63(iife35_()
)
        nw15_[int(8)] = d_199_v46_
        def iife36_(_pat_let2_0):
            def iife37_(d_204_dt__update__tmp_h3_):
                def iife38_(_pat_let3_0):
                    def iife39_(d_205_dt__update_hcf97_h0_):
                        return D25_DC63(d_205_dt__update_hcf97_h0_)
                    return iife39_(_pat_let3_0)
                return iife38_(pat_let_tv1_)
            return iife37_(_pat_let2_0)
        nw15_[int(9)] = iife36_(d_199_v46_)
        nw15_[int(10)] = d_199_v46_
        nw15_[int(11)] = d_199_v46_
        nw15_[int(12)] = d_199_v46_
        nw15_[int(13)] = d_199_v46_
        nw15_[int(14)] = d_199_v46_
        nw15_[int(15)] = d_199_v46_
        d_202_v50_ = nw15_
        d_206_v51_: _dafny.Array
        d_206_v51_ = (d_202_v50_ if d_138_v0_ else d_202_v50_)
        source8_ = d_206_v51_
        d_207___mcc_h1_ = source8_
        d_208_cf128_ = d_207___mcc_h1_
        d_141_v3_ = (d_141_v3_ if d_138_v0_ else d_141_v3_)
        (globalState).f6 = d_140_v2_
        d_209_v52_: C7
        nw16_ = C7()
        nw16_.ctor__(d_138_v0_, default__.safeModuloInt(d_140_v2_, d_140_v2_), d_140_v2_, (0) - ((0) - ((d_141_v3_)[default__.safeIndex(368, len(d_141_v3_))])), False)
        d_209_v52_ = nw16_
        d_197_v44_ = d_197_v44_
        r0 = d_138_v0_
        r1 = d_160_v18_
        return r0, r1

    @staticmethod
    def Main(noArgsParameter__):
        d_210_v1_: _dafny.Array
        def lambda4_(d_211_i0_):
            def iife40_():
                coll32_ = _dafny.Map()
                compr_32_: int
                for compr_32_ in _dafny.IntegerRange(-439, 38):
                    d_212_v0_: int = compr_32_
                    if ((-439) <= (d_212_v0_)) and ((d_212_v0_) < (38)):
                        coll32_[default__.safeModuloInt(d_212_v0_, len(_dafny.Map({not(True): (_dafny.MultiSet([470])).cardinality})))] = 481
                return _dafny.Map(coll32_)
            return (d_211_i0_) + (len(_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dmvo"))): len(_dafny.Map({(0) - (len(iife40_()
            )): True}))})])))

        init2_ = lambda4_
        nw17_ = _dafny.Array(None, 1)
        for i0_2_ in range(nw17_.length(0)):
            nw17_[i0_2_] = init2_(i0_2_)
        d_210_v1_ = nw17_
        d_213_v2_: int
        d_213_v2_ = 170
        d_214_v3_: _dafny.Set
        d_214_v3_ = _dafny.Set({d_213_v2_})
        d_215_v4_: _dafny.Map
        d_215_v4_ = _dafny.Map({d_214_v3_: d_214_v3_})
        d_216_v5_: bool
        d_216_v5_ = True
        d_217_v6_: _dafny.Map
        d_217_v6_ = _dafny.Map({not(False): d_216_v5_})
        d_218_v7_: _dafny.Seq
        d_218_v7_ = _dafny.SeqWithoutIsStrInference([d_216_v5_])
        d_219_v8_: _dafny.Map
        d_219_v8_ = _dafny.Map({(d_218_v7_)[default__.safeIndex(d_213_v2_, len(d_218_v7_))]: (0) - (d_213_v2_)})
        d_220_v9_: _dafny.Array
        nw18_ = _dafny.Array(_dafny.Map({}), 26)
        d_220_v9_ = nw18_
        d_221_v10_: _dafny.Seq
        d_221_v10_ = _dafny.SeqWithoutIsStrInference([d_213_v2_])
        d_222_v11_: _dafny.Seq
        d_222_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))
        d_223_v12_: _dafny.Seq
        d_223_v12_ = _dafny.SeqWithoutIsStrInference([d_222_v11_])
        d_224_v15_: _dafny.Map
        d_224_v15_ = _dafny.Map({(d_218_v7_)[default__.safeIndex(d_213_v2_, len(d_218_v7_))]: _dafny.SeqWithoutIsStrInference([d_218_v7_, d_218_v7_])})
        d_225_v16_: _dafny.Seq
        d_225_v16_ = _dafny.SeqWithoutIsStrInference([d_218_v7_])
        d_226_globalState_: GlobalState
        nw19_ = GlobalState()
        def iife41_():
            coll33_ = _dafny.Set()
            compr_33_: int
            for compr_33_ in (_dafny.Set({len(d_221_v10_), d_213_v2_, len(d_223_v12_), d_213_v2_, len(d_214_v3_)})).Elements:
                d_227_v13_: int = compr_33_
                if (d_227_v13_) in (_dafny.Set({len(d_221_v10_), d_213_v2_, len(d_223_v12_), d_213_v2_, len(d_214_v3_)})):
                    def iife42_():
                        coll34_ = _dafny.Map()
                        compr_34_: int
                        for compr_34_ in (_dafny.Set({159, (D0_DC2(456, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])) for d_228_i1_ in range(default__.abs(614))])): not(False)})))).cf3})).Elements:
                            d_229_v14_: int = compr_34_
                            if (d_229_v14_) in (_dafny.Set({159, (D0_DC2(456, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])) for d_228_i1_ in range(default__.abs(614))])): not(False)})))).cf3})):
                                coll34_[(d_229_v14_) + (-245)] = -836
                        return _dafny.Map(coll34_)
                    coll33_ = coll33_.union(_dafny.Set([default__.safeModuloInt(d_227_v13_, len(iife42_()
))]))
            return _dafny.Set(coll33_)
        nw19_.ctor__(831, d_210_v1_, 940, d_215_v4_, (d_214_v3_).intersection(d_214_v3_), True, 178, False, (d_217_v6_ if d_216_v5_ else d_217_v6_), -838, False, True, False, d_219_v8_, True, False, d_220_v9_, True, False, (d_214_v3_) - (iife41_()
        ), 78, False, d_210_v1_, ((d_224_v15_)[d_216_v5_] if (d_216_v5_) in (d_224_v15_) else d_225_v16_), d_219_v8_, True, _dafny.CodePoint('j'))
        d_226_globalState_ = nw19_
        d_230_v17_: _dafny.Map
        d_230_v17_ = _dafny.Map({d_213_v2_: d_216_v5_})
        d_231_v18_: _dafny.MultiSet
        d_231_v18_ = _dafny.MultiSet([len(d_214_v3_), d_213_v2_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_232_i2_ in range(default__.abs(554))]))])
        d_233_v19_: _dafny.Array
        nw20_ = _dafny.Array(None, 7)
        nw20_[int(0)] = (d_213_v2_) <= (len(d_222_v11_))
        nw20_[int(1)] = ((d_230_v17_)[d_213_v2_] if (d_213_v2_) in (d_230_v17_) else d_216_v5_)
        nw20_[int(2)] = d_216_v5_
        nw20_[int(3)] = d_216_v5_
        nw20_[int(4)] = (d_216_v5_ if not(d_216_v5_) else d_216_v5_)
        nw20_[int(5)] = (d_231_v18_).isdisjoint(d_231_v18_)
        nw20_[int(6)] = d_216_v5_
        d_233_v19_ = nw20_
        index8_ = default__.safeIndex(61, (d_233_v19_).length(0))
        (d_233_v19_)[index8_] = d_216_v5_
        index9_ = default__.safeIndex(61, (d_233_v19_).length(0))
        (d_233_v19_)[index9_] = d_216_v5_
        d_234_v20_: _dafny.Array
        nw21_ = _dafny.Array(_dafny.Map({}), 14)
        d_234_v20_ = nw21_
        index10_ = default__.safeIndex(505, (d_234_v20_).length(0))
        (d_234_v20_)[index10_] = (_dafny.Map({(d_233_v19_)[default__.safeIndex(61, (d_233_v19_).length(0))]: False})) | (d_217_v6_)
        index11_ = default__.safeIndex(505, (d_234_v20_).length(0))
        (d_234_v20_)[index11_] = d_217_v6_
        d_235_i3_: int
        d_235_i3_ = 0
        with _dafny.label("0"):
            while not(not((d_233_v19_)[default__.safeIndex(61, (d_233_v19_).length(0))])):
                with _dafny.c_label("0"):
                    if (d_235_i3_) >= (100):
                        raise _dafny.Break("0")
                    d_235_i3_ = (d_235_i3_) + (1)
                    index12_ = default__.safeIndex(353, (d_210_v1_).length(0))
                    (d_210_v1_)[index12_] = d_213_v2_
                    index13_ = default__.safeIndex(353, (d_210_v1_).length(0))
                    rhs5_ = ((d_219_v8_)[(d_233_v19_)[default__.safeIndex(61, (d_233_v19_).length(0))]] if ((d_233_v19_)[default__.safeIndex(61, (d_233_v19_).length(0))]) in (d_219_v8_) else (0) - (d_213_v2_))
                    rhs6_ = d_213_v2_
                    rhs7_ = d_213_v2_
                    lhs3_ = d_226_globalState_
                    lhs4_ = d_226_globalState_
                    lhs5_ = d_210_v1_
                    lhs6_ = default__.safeIndex(353, (d_210_v1_).length(0))
                    lhs3_.f6 = rhs5_
                    lhs4_.f2 = rhs6_
                    lhs5_[lhs6_] = rhs7_
                    if default__.fm0(d_222_v11_, (d_233_v19_)[default__.safeIndex(61, (d_233_v19_).length(0))], (default__.fm1(d_216_v5_, d_218_v7_, d_216_v5_, d_226_globalState_)) <= (d_213_v2_), d_214_v3_, d_226_globalState_):
                        d_236_v21_: str
                        d_236_v21_ = _dafny.CodePoint('p')
                        d_236_v21_ = d_236_v21_
                        (d_226_globalState_).f6 = default__.safeDivisionInt(((d_210_v1_)[default__.safeIndex(353, (d_210_v1_).length(0))]) * (d_213_v2_), default__.safeDivisionInt(len(d_222_v11_), len(d_222_v11_)))
                        d_237_v22_: D0
                        d_237_v22_ = D0_DC3(D0_DC1((d_233_v19_)[default__.safeIndex(61, (d_233_v19_).length(0))]))
                        d_238_v23_: D0
                        d_238_v23_ = D0_DC1(d_216_v5_)
                        d_239_v24_: bool
                        d_240_v25_: _dafny.Seq
                        out0_: bool
                        out1_: _dafny.Seq
                        out0_, out1_ = default__.m0(D0_DC3(d_237_v22_), d_238_v23_, d_226_globalState_)
                        d_239_v24_ = out0_
                        d_240_v25_ = out1_
                        d_221_v10_ = d_221_v10_
                        d_241_v26_: C0
                        nw22_ = C0()
                        nw22_.ctor__()
                        d_241_v26_ = nw22_
                    elif True:
                        (d_226_globalState_).f25 = (d_233_v19_)[default__.safeIndex(61, (d_233_v19_).length(0))]
                        d_216_v5_ = not((d_233_v19_)[default__.safeIndex(61, (d_233_v19_).length(0))])
                        d_222_v11_ = d_222_v11_
                        d_242_v27_: T0
                        nw23_ = C4()
                        nw23_.ctor__(_dafny.MultiSet(d_218_v7_), d_216_v5_)
                        d_242_v27_ = nw23_
                        d_243_v28_: str
                        d_243_v28_ = _dafny.CodePoint('w')
                        d_244_v29_: D25
                        d_244_v29_ = D25_DC64((d_233_v19_)[default__.safeIndex(61, (d_233_v19_).length(0))], (d_210_v1_)[default__.safeIndex(353, (d_210_v1_).length(0))], d_243_v28_, (d_242_v27_).f27, d_213_v2_)
                        pat_let_tv2_ = d_243_v28_
                        pat_let_tv3_ = d_222_v11_
                        index14_ = default__.safeIndex(61, (d_233_v19_).length(0))
                        def iife43_(_pat_let4_0):
                            def iife44_(d_245_dt__update__tmp_h0_):
                                def iife45_(_pat_let5_0):
                                    def iife46_(d_246_dt__update_hcf100_h0_):
                                        def iife47_(_pat_let6_0):
                                            def iife48_(d_247_dt__update_hcf99_h0_):
                                                return D25_DC64((d_245_dt__update__tmp_h0_).cf98, d_247_dt__update_hcf99_h0_, d_246_dt__update_hcf100_h0_, (d_245_dt__update__tmp_h0_).cf101, (d_245_dt__update__tmp_h0_).cf102)
                                            return iife48_(_pat_let6_0)
                                        return iife47_(len(pat_let_tv3_))
                                    return iife46_(_pat_let5_0)
                                return iife45_(pat_let_tv2_)
                            return iife44_(_pat_let4_0)
                        rhs8_ = (iife43_(d_244_v29_)).cf102
                        rhs9_ = (((d_218_v7_) + (d_218_v7_)).set(default__.safeIndex(d_213_v2_, len((d_218_v7_) + (d_218_v7_))), not(True))) != (d_218_v7_)
                        rhs10_ = d_216_v5_
                        rhs11_ = d_242_v27_
                        lhs7_ = d_226_globalState_
                        lhs8_ = d_226_globalState_
                        lhs9_ = d_233_v19_
                        lhs10_ = default__.safeIndex(61, (d_233_v19_).length(0))
                        lhs7_.f2 = rhs8_
                        lhs8_.f14 = rhs9_
                        lhs9_[lhs10_] = rhs10_
                        d_242_v27_ = rhs11_
                        rhs12_ = (d_216_v5_ if True else (d_233_v19_)[default__.safeIndex(61, (d_233_v19_).length(0))])
                        rhs13_ = not((d_233_v19_)[default__.safeIndex(61, (d_233_v19_).length(0))])
                        rhs14_ = (d_242_v27_).fm3(d_226_globalState_)
                        lhs11_ = d_226_globalState_
                        lhs12_ = d_226_globalState_
                        lhs11_.f14 = rhs12_
                        lhs12_.f25 = rhs13_
                        d_222_v11_ = rhs14_
                    d_248_v30_: T0
                    nw24_ = C3()
                    nw24_.ctor__((d_233_v19_)[default__.safeIndex(61, (d_233_v19_).length(0))], (d_233_v19_)[default__.safeIndex(61, (d_233_v19_).length(0))])
                    d_248_v30_ = nw24_
                    d_249_v31_: C3
                    nw25_ = C3()
                    nw25_.ctor__((d_233_v19_)[default__.safeIndex(61, (d_233_v19_).length(0))], (d_248_v30_).f27)
                    d_249_v31_ = nw25_
                    d_250_v32_: _dafny.Map
                    d_250_v32_ = _dafny.Map({d_248_v30_: d_249_v31_})
                    d_251_v33_: D36
                    d_251_v33_ = D36_DC87(d_250_v32_)
                    d_252_v34_: _dafny.Seq
                    d_252_v34_ = _dafny.SeqWithoutIsStrInference([d_249_v31_, d_249_v31_])
                    d_253_v35_: _dafny.Array
                    nw26_ = _dafny.Array(None, 16)
                    nw26_[int(0)] = (d_250_v32_) | (d_250_v32_)
                    nw26_[int(1)] = d_250_v32_
                    nw26_[int(2)] = d_250_v32_
                    nw26_[int(3)] = _dafny.Map({d_248_v30_: d_249_v31_})
                    nw26_[int(4)] = _dafny.Map({d_248_v30_: d_249_v31_})
                    nw26_[int(5)] = (_dafny.Map({d_248_v30_: d_249_v31_})) | (_dafny.Map({d_248_v30_: d_249_v31_}))
                    nw26_[int(6)] = d_250_v32_
                    nw26_[int(7)] = d_250_v32_
                    nw26_[int(8)] = d_250_v32_
                    nw26_[int(9)] = (d_251_v33_).cf145
                    nw26_[int(10)] = (d_250_v32_).set(d_248_v30_, d_249_v31_)
                    nw26_[int(11)] = (d_250_v32_ if (d_248_v30_).f27 else d_250_v32_)
                    nw26_[int(12)] = d_250_v32_
                    nw26_[int(13)] = d_250_v32_
                    nw26_[int(14)] = d_250_v32_
                    nw26_[int(15)] = (_dafny.Map({d_248_v30_: d_249_v31_})).set(d_248_v30_, (d_252_v34_)[default__.safeIndex(default__.fm1(True, d_218_v7_, d_249_v31_.f38, d_226_globalState_), len(d_252_v34_))])
                    d_253_v35_ = nw26_
                    nw27_ = _dafny.Array(_dafny.Map({}), 20)
                    d_253_v35_ = nw27_
                    if (d_248_v30_).f27:
                        (d_226_globalState_).f25 = (d_248_v30_).f27
                        d_254_v36_: _dafny.Array
                        d_255_v37_: _dafny.Seq
                        d_256_v38_: _dafny.Array
                        d_257_v39_: bool
                        out2_: _dafny.Array
                        out3_: _dafny.Seq
                        out4_: _dafny.Array
                        out5_: bool
                        out2_, out3_, out4_, out5_ = (d_249_v31_).m2(d_226_globalState_)
                        d_254_v36_ = out2_
                        d_255_v37_ = out3_
                        d_256_v38_ = out4_
                        d_257_v39_ = out5_
                        d_222_v11_ = d_222_v11_
                        d_258_v40_: _dafny.Array
                        nw28_ = _dafny.Array(None, 20)
                        nw28_[int(0)] = d_249_v31_
                        nw28_[int(1)] = d_249_v31_
                        nw28_[int(2)] = d_249_v31_
                        nw28_[int(3)] = d_249_v31_
                        nw28_[int(4)] = d_249_v31_
                        nw28_[int(5)] = d_249_v31_
                        nw28_[int(6)] = d_249_v31_
                        nw28_[int(7)] = d_249_v31_
                        nw28_[int(8)] = d_249_v31_
                        nw28_[int(9)] = (d_249_v31_ if default__.fm0(d_255_v37_, not(d_249_v31_.f38), d_249_v31_.f38, d_214_v3_, d_226_globalState_) else d_249_v31_)
                        nw28_[int(10)] = d_249_v31_
                        nw28_[int(11)] = d_249_v31_
                        nw28_[int(12)] = d_249_v31_
                        nw28_[int(13)] = d_249_v31_
                        nw28_[int(14)] = d_249_v31_
                        nw28_[int(15)] = d_249_v31_
                        nw28_[int(16)] = d_249_v31_
                        nw28_[int(17)] = d_249_v31_
                        nw28_[int(18)] = d_249_v31_
                        nw28_[int(19)] = d_249_v31_
                        d_258_v40_ = nw28_
                        index15_ = default__.safeIndex(854, (d_258_v40_).length(0))
                        (d_258_v40_)[index15_] = d_249_v31_
                        index16_ = default__.safeIndex(854, (d_258_v40_).length(0))
                        (d_258_v40_)[index16_] = d_249_v31_
                        d_259_v41_: C8
                        nw29_ = C8()
                        nw29_.ctor__((d_248_v30_).f27, (d_210_v1_)[default__.safeIndex(353, (d_210_v1_).length(0))], d_249_v31_.f38)
                        d_259_v41_ = nw29_
                        d_259_v41_ = d_259_v41_
                    elif True:
                        d_260_v42_: C6
                        nw30_ = C6()
                        nw30_.ctor__(d_213_v2_, ((d_210_v1_)[default__.safeIndex(353, (d_210_v1_).length(0))]) * (d_213_v2_), (d_248_v30_).f27)
                        d_260_v42_ = nw30_
                        d_261_v43_: T2
                        nw31_ = C2()
                        nw31_.ctor__((d_248_v30_).f27, not((d_248_v30_).f27))
                        d_261_v43_ = nw31_
                        d_262_v44_: _dafny.Seq
                        d_262_v44_ = _dafny.SeqWithoutIsStrInference([d_261_v43_, d_261_v43_, d_261_v43_])
                        index17_ = default__.safeIndex(353, (d_210_v1_).length(0))
                        (d_210_v1_)[index17_] = len(d_262_v44_)
                        (d_226_globalState_).f25 = d_249_v31_.f38
                        index18_ = default__.safeIndex(61, (d_233_v19_).length(0))
                        (d_233_v19_)[index18_] = (default__.fm1(not(False), d_218_v7_, (d_248_v30_).f27, d_226_globalState_)) == (d_213_v2_)
                        index19_ = default__.safeIndex(61, (d_233_v19_).length(0))
                        rhs15_ = default__.fm1((d_248_v30_).f27, (d_218_v7_) + (default__.fm32(d_249_v31_.f38, d_221_v10_, (d_248_v30_).f27, d_226_globalState_)), (default__.fm74(len(d_221_v10_), (d_248_v30_).f27, 576, d_249_v31_.f38, d_226_globalState_)).cf91, d_226_globalState_)
                        rhs16_ = not(d_216_v5_)
                        rhs17_ = (d_213_v2_ if d_216_v5_ else len(default__.fm52((d_210_v1_)[default__.safeIndex(353, (d_210_v1_).length(0))], (d_248_v30_).f27, default__.fm34(d_226_globalState_), d_226_globalState_)))
                        rhs18_ = not((867) < (((d_210_v1_)[default__.safeIndex(353, (d_210_v1_).length(0))]) - (d_213_v2_)))
                        rhs19_ = (d_248_v30_).f27
                        lhs13_ = d_226_globalState_
                        lhs14_ = d_249_v31_
                        lhs15_ = d_226_globalState_
                        lhs16_ = d_233_v19_
                        lhs17_ = default__.safeIndex(61, (d_233_v19_).length(0))
                        lhs13_.f2 = rhs15_
                        lhs14_.f38 = rhs16_
                        lhs15_.f6 = rhs17_
                        d_216_v5_ = rhs18_
                        lhs16_[lhs17_] = rhs19_
                    pass
            pass
        d_263_v45_: T2
        nw32_ = C1()
        nw32_.ctor__(d_216_v5_)
        d_263_v45_ = nw32_
        d_264_v46_: _dafny.MultiSet
        d_264_v46_ = _dafny.MultiSet([d_263_v45_])
        d_265_v47_: _dafny.Seq
        d_265_v47_ = _dafny.SeqWithoutIsStrInference([(d_264_v46_).set(d_263_v45_, default__.abs(d_213_v2_)), d_264_v46_, d_264_v46_, d_264_v46_, _dafny.MultiSet([d_263_v45_])])
        d_266_i4_: int
        d_266_i4_ = 0
        with _dafny.label("1"):
            while ((d_216_v5_) and ((d_263_v45_).f27) if (d_264_v46_).issubset((d_265_v47_)[default__.safeIndex(d_213_v2_, len(d_265_v47_))]) else (d_233_v19_)[default__.safeIndex(61, (d_233_v19_).length(0))]):
                with _dafny.c_label("1"):
                    if (d_266_i4_) >= (100):
                        raise _dafny.Break("1")
                    d_266_i4_ = (d_266_i4_) + (1)
                    d_267_v48_: C3
                    nw33_ = C3()
                    nw33_.ctor__(default__.fm0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wsikcuh")), (d_233_v19_)[default__.safeIndex(61, (d_233_v19_).length(0))], d_216_v5_, d_214_v3_, d_226_globalState_), (d_214_v3_) == (d_214_v3_))
                    d_267_v48_ = nw33_
                    d_268_v49_: T1
                    nw34_ = C6()
                    nw34_.ctor__(default__.safeDivisionInt(865, d_213_v2_), (d_221_v10_)[default__.safeIndex(d_213_v2_, len(d_221_v10_))], (d_216_v5_) and (d_267_v48_.f38))
                    d_268_v49_ = nw34_
                    rhs20_ = d_267_v48_
                    rhs21_ = (0) - (default__.fm1((d_263_v45_).f27, (d_218_v7_) + (d_218_v7_), (d_222_v11_) == (d_222_v11_), d_226_globalState_))
                    rhs22_ = d_268_v49_
                    rhs23_ = d_216_v5_
                    rhs24_ = (d_233_v19_)[default__.safeIndex(61, (d_233_v19_).length(0))]
                    lhs18_ = d_226_globalState_
                    lhs19_ = d_226_globalState_
                    d_267_v48_ = rhs20_
                    d_213_v2_ = rhs21_
                    d_268_v49_ = rhs22_
                    lhs18_.f14 = rhs23_
                    lhs19_.f21 = rhs24_
                    d_269_v50_: str
                    d_269_v50_ = _dafny.CodePoint('g')
                    rhs25_ = not((d_263_v45_).f27)
                    rhs26_ = d_269_v50_
                    rhs27_ = not((d_263_v45_).f27)
                    lhs20_ = d_226_globalState_
                    lhs21_ = d_226_globalState_
                    lhs20_.f5 = rhs25_
                    d_269_v50_ = rhs26_
                    lhs21_.f14 = rhs27_
                    d_233_v19_ = d_233_v19_
                    d_270_v51_: _dafny.Set
                    d_270_v51_ = _dafny.Set({d_222_v11_})
                    d_271_v53_: _dafny.Array
                    def lambda5_(d_272_v50_):
                        def lambda6_(d_273_i5_):
                            return d_272_v50_

                        return lambda6_

                    init3_ = lambda5_(d_269_v50_)
                    nw35_ = _dafny.Array(None, 10)
                    for i0_3_ in range(nw35_.length(0)):
                        nw35_[i0_3_] = init3_(i0_3_)
                    d_271_v53_ = nw35_
                    def iife49_():
                        coll35_ = _dafny.Set()
                        compr_35_: _dafny.Seq
                        for compr_35_ in (_dafny.MultiSet([d_222_v11_])).Elements:
                            d_274_v52_: _dafny.Seq = compr_35_
                            if (d_274_v52_) in (_dafny.MultiSet([d_222_v11_])):
                                coll35_ = coll35_.union(_dafny.Set([d_274_v52_]))
                        return _dafny.Set(coll35_)
                    (d_268_v49_).m3((d_270_v51_).isdisjoint(iife49_()
                    ), d_271_v53_, (d_268_v49_).f28, d_226_globalState_)
                    pass
            pass
        if False:
            d_210_v1_ = d_210_v1_
            (d_226_globalState_).f2 = (d_213_v2_) + (d_213_v2_)
            d_275_v54_: _dafny.Map
            d_275_v54_ = _dafny.Map({d_210_v1_: d_233_v19_})
            d_233_v19_ = ((((d_275_v54_)[d_210_v1_] if (d_210_v1_) in (d_275_v54_) else d_233_v19_) if (d_263_v45_).f27 else d_233_v19_) if False else d_233_v19_)
            index20_ = default__.safeIndex(61, (d_233_v19_).length(0))
            (d_233_v19_)[index20_] = ((d_263_v45_).f27) or (not(d_216_v5_))
            d_276_v55_: C8
            nw36_ = C8()
            nw36_.ctor__(((d_263_v45_).f27) or ((d_263_v45_).f27), d_213_v2_, (d_263_v45_).f27)
            d_276_v55_ = nw36_
        elif True:
            d_277_v56_: _dafny.Seq
            d_277_v56_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([218, d_213_v2_]), d_221_v10_])
            d_278_v57_: D23
            d_278_v57_ = D23_DC61(d_210_v1_, d_216_v5_, d_213_v2_)
            rhs28_ = (d_213_v2_) <= (default__.safeDivisionInt(687, (d_278_v57_).cf95))
            rhs29_ = (d_213_v2_) < (d_213_v2_)
            rhs30_ = _dafny.SeqWithoutIsStrInference([((d_221_v10_) + (d_221_v10_)).set(default__.safeIndex(len(d_222_v11_), len((d_221_v10_) + (d_221_v10_))), 133)])
            rhs31_ = d_222_v11_
            rhs32_ = (d_263_v45_).f27
            lhs22_ = d_226_globalState_
            lhs23_ = d_226_globalState_
            lhs22_.f5 = rhs28_
            lhs23_.f21 = rhs29_
            d_277_v56_ = rhs30_
            d_222_v11_ = rhs31_
            d_216_v5_ = rhs32_
            (d_226_globalState_).f2 = len(d_221_v10_)
            (d_226_globalState_).f2 = d_213_v2_
            index21_ = default__.safeIndex(956, (d_210_v1_).length(0))
            (d_210_v1_)[index21_] = d_213_v2_
            index22_ = default__.safeIndex(956, (d_210_v1_).length(0))
            (d_210_v1_)[index22_] = d_213_v2_
            (d_226_globalState_).f21 = d_216_v5_
        d_279_v58_: _dafny.Map
        d_279_v58_ = _dafny.Map({80: d_213_v2_})
        hi2_ = d_213_v2_
        for d_280_i6_ in range(len(d_279_v58_), hi2_):
            d_222_v11_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ab"))) + (d_222_v11_)) + (d_222_v11_)
            d_281_v59_: _dafny.Seq
            d_281_v59_ = _dafny.SeqWithoutIsStrInference([d_210_v1_])
            d_282_v60_: C7
            nw37_ = C7()
            nw37_.ctor__((d_263_v45_).f27, d_213_v2_, (d_213_v2_) + (d_280_i6_), len(d_281_v59_), (d_263_v45_).f27)
            d_282_v60_ = nw37_
            index23_ = default__.safeIndex(61, (d_233_v19_).length(0))
            (d_233_v19_)[index23_] = not((d_233_v19_)[default__.safeIndex(61, (d_233_v19_).length(0))])
            index24_ = default__.safeIndex(505, (d_234_v20_).length(0))
            (d_234_v20_)[index24_] = ((d_234_v20_)[default__.safeIndex(505, (d_234_v20_).length(0))]).set(d_282_v60_.f35, not (d_282_v60_.f35) or ((d_233_v19_)[default__.safeIndex(61, (d_233_v19_).length(0))]))
        d_283_v61_: str
        d_283_v61_ = _dafny.CodePoint('s')
        hi3_ = len(((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_284_i8_ in range(default__.abs(923))])).set(default__.safeIndex(d_213_v2_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_285_i8_ in range(default__.abs(923))]))), d_283_v61_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ht"))))
        for d_286_i7_ in range(d_213_v2_, hi3_):
            d_287_v62_: D0
            d_287_v62_ = D0_DC1((d_233_v19_)[default__.safeIndex(61, (d_233_v19_).length(0))])
            d_288_v63_: bool
            d_289_v64_: bool
            out6_: bool
            out7_: bool
            out6_, out7_ = (d_263_v45_).m4(D0_DC3(d_287_v62_), ((d_263_v45_).f27 if (d_233_v19_)[default__.safeIndex(61, (d_233_v19_).length(0))] else not((d_263_v45_).f27)), d_226_globalState_)
            d_288_v63_ = out6_
            d_289_v64_ = out7_
            index25_ = default__.safeIndex(61, (d_233_v19_).length(0))
            (d_233_v19_)[index25_] = d_216_v5_
            d_290_v65_: D18
            d_290_v65_ = D18_DC42(d_222_v11_, d_216_v5_)
            (d_226_globalState_).f25 = (d_290_v65_).cf58
            d_288_v63_ = d_216_v5_
        d_291_v66_: int
        out8_: int
        out8_ = (d_263_v45_).m1(d_226_globalState_)
        d_291_v66_ = out8_
        d_292_v67_: int
        out9_: int
        out9_ = (d_263_v45_).m1(d_226_globalState_)
        d_292_v67_ = out9_
        d_293_v68_: D0
        d_293_v68_ = D0_DC3(D0_DC0(d_213_v2_))
        d_294_v69_: bool
        d_295_v70_: bool
        out10_: bool
        out11_: bool
        out10_, out11_ = (d_263_v45_).m4(d_293_v68_, (d_231_v18_).ispropersubset(d_231_v18_), d_226_globalState_)
        d_294_v69_ = out10_
        d_295_v70_ = out11_
        d_296_v71_: _dafny.MultiSet
        d_296_v71_ = _dafny.MultiSet([d_295_v70_, d_295_v70_])
        d_297_v72_: C4
        nw38_ = C4()
        nw38_.ctor__(d_296_v71_, False)
        d_297_v72_ = nw38_
        d_298_v73_: _dafny.Seq
        d_298_v73_ = _dafny.SeqWithoutIsStrInference([d_297_v72_])
        d_297_v72_ = (d_298_v73_)[default__.safeIndex(d_292_v67_, len(d_298_v73_))]
        d_299_v74_: _dafny.Map
        d_299_v74_ = _dafny.Map({len(d_222_v11_): d_283_v61_})
        d_300_v75_: _dafny.Seq
        d_300_v75_ = _dafny.SeqWithoutIsStrInference([d_299_v74_])
        d_301_v77_: _dafny.Set
        def iife50_():
            coll36_ = _dafny.Set()
            compr_36_: _dafny.Map
            for compr_36_ in (d_300_v75_).Elements:
                d_302_v76_: _dafny.Map = compr_36_
                if (d_302_v76_) in (d_300_v75_):
                    coll36_ = coll36_.union(_dafny.Set([d_302_v76_]))
            return _dafny.Set(coll36_)
        d_301_v77_ = iife50_()
        
        d_303_v78_: _dafny.Seq
        d_303_v78_ = _dafny.SeqWithoutIsStrInference([d_301_v77_, d_301_v77_])
        d_304_v79_: _dafny.Map
        d_304_v79_ = _dafny.Map({(d_291_v66_) - (d_291_v66_): (d_222_v11_) + ((D30_DC75(d_213_v2_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_305_i9_ in range(default__.abs(309))]), d_303_v78_, d_233_v19_)).cf125)})
        index26_ = default__.safeIndex(61, (d_233_v19_).length(0))
        rhs33_ = d_304_v79_
        rhs34_ = d_295_v70_
        lhs24_ = d_233_v19_
        lhs25_ = default__.safeIndex(61, (d_233_v19_).length(0))
        d_304_v79_ = rhs33_
        lhs24_[lhs25_] = rhs34_
        d_306_v80_: _dafny.MultiSet
        d_306_v80_ = _dafny.MultiSet([default__.fm25(d_292_v67_, d_291_v66_, d_216_v5_, d_216_v5_, d_226_globalState_)])
        d_307_v81_: _dafny.MultiSet
        d_307_v81_ = _dafny.MultiSet([d_214_v3_])
        source9_ = d_307_v81_
        d_308___mcc_h0_ = source9_
        d_309_cf103_ = d_308___mcc_h0_
        if (not(not(((d_263_v45_).fm4(D0_DC1(d_295_v70_), d_226_globalState_)) == (True)))) and (d_216_v5_):
            d_310_v82_: C0
            nw39_ = C0()
            nw39_.ctor__()
            d_310_v82_ = nw39_
            d_311_v83_: int
            out12_: int
            out12_ = (d_297_v72_).m1(d_226_globalState_)
            d_311_v83_ = out12_
            d_312_v84_: _dafny.Array
            d_313_v85_: _dafny.Seq
            d_314_v86_: _dafny.Array
            d_315_v87_: bool
            out13_: _dafny.Array
            out14_: _dafny.Seq
            out15_: _dafny.Array
            out16_: bool
            out13_, out14_, out15_, out16_ = (d_263_v45_).m2(d_226_globalState_)
            d_312_v84_ = out13_
            d_313_v85_ = out14_
            d_314_v86_ = out15_
            d_315_v87_ = out16_
            d_316_v88_: _dafny.Map
            d_316_v88_ = _dafny.Map({default__.fm29(d_311_v83_, (d_263_v45_).f27, d_226_globalState_): d_210_v1_})
            d_317_v89_: D23
            d_317_v89_ = D23_DC61(((d_316_v88_)[d_283_v61_] if (d_283_v61_) in (d_316_v88_) else d_210_v1_), (d_263_v45_).f27, d_213_v2_)
            d_210_v1_ = (d_317_v89_).cf93
            d_292_v67_ = d_292_v67_
        elif True:
            (d_226_globalState_).f2 = (0) - ((0) - (d_292_v67_))
            d_318_v90_: C5
            nw40_ = C5()
            nw40_.ctor__((d_263_v45_).f27, len(d_222_v11_), d_213_v2_, (d_263_v45_).f27)
            d_318_v90_ = nw40_
            d_319_v91_: C5
            nw41_ = C5()
            nw41_.ctor__((d_263_v45_).f27, len(default__.fm28(d_226_globalState_)), d_213_v2_, (d_263_v45_).f27)
            d_319_v91_ = nw41_
            index27_ = default__.safeIndex(61, (d_233_v19_).length(0))
            (d_233_v19_)[index27_] = d_216_v5_
            d_320_v92_: _dafny.Array
            nw42_ = _dafny.Array(D29.default()(), 27)
            d_320_v92_ = nw42_
            d_321_v93_: D29
            d_321_v93_ = D29_DC72(D29_DC71(d_213_v2_, d_214_v3_, d_294_v69_, (d_263_v45_).f27, d_210_v1_))
            index28_ = default__.safeIndex(457, (d_320_v92_).length(0))
            (d_320_v92_)[index28_] = d_321_v93_
            d_322_v94_: D9
            d_322_v94_ = D9_DC21((d_233_v19_)[default__.safeIndex(61, (d_233_v19_).length(0))], (default__.fm1(False, (d_218_v7_).set(default__.safeIndex(d_291_v66_, len(d_218_v7_)), not((d_233_v19_)[default__.safeIndex(61, (d_233_v19_).length(0))])), (d_319_v91_).f39, d_226_globalState_)) * (len(d_222_v11_)))
            d_323_v95_: _dafny.Array
            def lambda7_(d_324_v11_):
                def lambda8_(d_325_i10_):
                    return d_324_v11_

                return lambda8_

            init4_ = lambda7_(d_222_v11_)
            nw43_ = _dafny.Array(None, 10)
            for i0_4_ in range(nw43_.length(0)):
                nw43_[i0_4_] = init4_(i0_4_)
            d_323_v95_ = nw43_
            index29_ = default__.safeIndex(56, (d_323_v95_).length(0))
            (d_323_v95_)[index29_] = d_222_v11_
            d_326_v96_: D19
            d_326_v96_ = D19_DC46(d_292_v67_, d_210_v1_, (d_318_v90_).f39)
            d_327_v97_: C6
            nw44_ = C6()
            nw44_.ctor__(d_292_v67_, (d_326_v96_).cf64, (447) > (d_292_v67_))
            d_327_v97_ = nw44_
            index30_ = default__.safeIndex(457, (d_320_v92_).length(0))
            index31_ = default__.safeIndex(56, (d_323_v95_).length(0))
            rhs35_ = d_321_v93_
            rhs36_ = d_322_v94_
            rhs37_ = (_dafny.SeqWithoutIsStrInference([d_283_v61_ for d_328_i11_ in range(default__.abs(523))])) + (d_222_v11_)
            rhs38_ = d_297_v72_
            rhs39_ = d_327_v97_
            lhs26_ = d_320_v92_
            lhs27_ = default__.safeIndex(457, (d_320_v92_).length(0))
            lhs28_ = d_323_v95_
            lhs29_ = default__.safeIndex(56, (d_323_v95_).length(0))
            lhs26_[lhs27_] = rhs35_
            d_322_v94_ = rhs36_
            lhs28_[lhs29_] = rhs37_
            d_297_v72_ = rhs38_
            d_327_v97_ = rhs39_
        d_210_v1_ = d_210_v1_
        d_329_v98_: int
        out17_: int
        out17_ = (d_263_v45_).m1(d_226_globalState_)
        d_329_v98_ = out17_
        d_292_v67_ = d_292_v67_
        (d_226_globalState_).f6 = 155
        hi4_ = d_292_v67_
        for d_330_i12_ in range((d_291_v66_ if d_295_v70_ else len(d_230_v17_)), hi4_):
            d_295_v70_ = (315) == (d_213_v2_)
            d_219_v8_ = (d_219_v8_).set(d_295_v70_, len((default__.fm17(d_291_v66_, d_222_v11_, d_213_v2_, d_226_globalState_)) + (d_218_v7_)))
            (d_226_globalState_).f5 = ((d_221_v10_)[default__.safeIndex(d_213_v2_, len(d_221_v10_))]) >= (d_213_v2_)
            d_331_v99_: _dafny.Set
            d_331_v99_ = _dafny.Set({not((d_218_v7_)[default__.safeIndex(d_213_v2_, len(d_218_v7_))])})
            d_332_v100_: _dafny.Array
            nw45_ = _dafny.Array(None, 6)
            nw45_[int(0)] = d_331_v99_
            nw45_[int(1)] = d_331_v99_
            nw45_[int(2)] = d_331_v99_
            nw45_[int(3)] = _dafny.Set({(d_263_v45_).f27, (d_263_v45_).f27, (d_263_v45_).f27})
            nw45_[int(4)] = d_331_v99_
            nw45_[int(5)] = (d_331_v99_) - (d_331_v99_)
            d_332_v100_ = nw45_
            d_333_v101_: _dafny.Seq
            d_333_v101_ = _dafny.SeqWithoutIsStrInference([d_331_v99_, d_331_v99_, d_331_v99_])
            index32_ = default__.safeIndex(2, (d_332_v100_).length(0))
            (d_332_v100_)[index32_] = (d_333_v101_)[default__.safeIndex(d_330_i12_, len(d_333_v101_))]
            index33_ = default__.safeIndex(2, (d_332_v100_).length(0))
            rhs40_ = ((d_331_v99_) | (d_331_v99_)) | (d_331_v99_)
            rhs41_ = ((d_263_v45_).f27) and (False)
            rhs42_ = d_291_v66_
            rhs43_ = (d_219_v8_) | (d_219_v8_)
            rhs44_ = d_221_v10_
            lhs30_ = d_332_v100_
            lhs31_ = default__.safeIndex(2, (d_332_v100_).length(0))
            lhs30_[lhs31_] = rhs40_
            d_295_v70_ = rhs41_
            d_291_v66_ = rhs42_
            d_219_v8_ = rhs43_
            d_221_v10_ = rhs44_
        if (d_292_v67_) != ((d_292_v67_) * (d_213_v2_)):
            d_334_v102_: D22
            d_334_v102_ = D22_DC57(d_294_v69_, (d_222_v11_)[default__.safeIndex(d_292_v67_, len(d_222_v11_))], (d_263_v45_).f27)
            d_335_v103_: D22
            d_335_v103_ = D22_DC56(d_299_v74_)
            d_336_v104_: _dafny.Seq
            d_336_v104_ = _dafny.SeqWithoutIsStrInference([default__.fm75((d_263_v45_).f27, (d_233_v19_)[default__.safeIndex(61, (d_233_v19_).length(0))], True, d_292_v67_, d_226_globalState_), d_335_v103_])
            d_337_v105_: _dafny.Seq
            d_337_v105_ = _dafny.SeqWithoutIsStrInference([(d_336_v104_).set(default__.safeIndex(d_213_v2_, len(d_336_v104_)), d_335_v103_)])
            d_338_v106_: _dafny.Seq
            d_338_v106_ = _dafny.SeqWithoutIsStrInference([(d_337_v105_)[default__.safeIndex(d_213_v2_, len(d_337_v105_))]])
            d_339_v107_: _dafny.Map
            d_339_v107_ = _dafny.Map({d_334_v102_: d_338_v106_})
            pat_let_tv4_ = d_283_v61_
            def iife51_(_pat_let7_0):
                def iife52_(d_340_dt__update__tmp_h2_):
                    def iife53_(_pat_let8_0):
                        def iife54_(d_341_dt__update_hcf85_h0_):
                            return D22_DC57((d_340_dt__update__tmp_h2_).cf84, d_341_dt__update_hcf85_h0_, (d_340_dt__update__tmp_h2_).cf86)
                        return iife54_(_pat_let8_0)
                    return iife53_(pat_let_tv4_)
                return iife52_(_pat_let7_0)
            d_339_v107_ = (d_339_v107_).set((iife51_(d_334_v102_) if d_295_v70_ else d_334_v102_), _dafny.SeqWithoutIsStrInference([d_336_v104_, d_336_v104_, d_336_v104_]))
            d_342_v108_: _dafny.Array
            nw46_ = _dafny.Array(_dafny.Set({}), 6)
            d_342_v108_ = nw46_
            d_343_v109_: _dafny.Set
            d_343_v109_ = _dafny.Set({(d_263_v45_).f27})
            index34_ = default__.safeIndex(125, (d_342_v108_).length(0))
            (d_342_v108_)[index34_] = d_343_v109_
            index35_ = default__.safeIndex(125, (d_342_v108_).length(0))
            index36_ = default__.safeIndex(61, (d_233_v19_).length(0))
            rhs45_ = ((d_343_v109_).intersection(d_343_v109_)) - (d_343_v109_)
            rhs46_ = d_216_v5_
            rhs47_ = (d_233_v19_ if not(False) else d_233_v19_)
            lhs32_ = d_342_v108_
            lhs33_ = default__.safeIndex(125, (d_342_v108_).length(0))
            lhs34_ = d_233_v19_
            lhs35_ = default__.safeIndex(61, (d_233_v19_).length(0))
            lhs32_[lhs33_] = rhs45_
            lhs34_[lhs35_] = rhs46_
            d_233_v19_ = rhs47_
            d_344_v110_: _dafny.Seq
            d_344_v110_ = _dafny.SeqWithoutIsStrInference([d_231_v18_])
            d_345_v111_: D21
            d_345_v111_ = D21_DC55(d_291_v66_)
            d_231_v18_ = ((d_344_v110_)[default__.safeIndex(d_213_v2_, len(d_344_v110_))]).set((d_345_v111_).cf82, default__.abs(d_292_v67_))
            d_346_v112_: D0
            d_346_v112_ = D0_DC0(d_292_v67_)
            d_347_v113_: D0
            d_347_v113_ = D0_DC1(d_295_v70_)
            d_348_v114_: bool
            d_349_v115_: _dafny.Seq
            out18_: bool
            out19_: _dafny.Seq
            out18_, out19_ = default__.m0(D0_DC3(d_346_v112_), d_347_v113_, d_226_globalState_)
            d_348_v114_ = out18_
            d_349_v115_ = out19_
            (d_226_globalState_).f5 = (d_291_v66_) > (d_292_v67_)
        elif True:
            d_350_v116_: _dafny.Array
            d_351_v117_: _dafny.Seq
            d_352_v118_: _dafny.Array
            d_353_v119_: bool
            out20_: _dafny.Array
            out21_: _dafny.Seq
            out22_: _dafny.Array
            out23_: bool
            out20_, out21_, out22_, out23_ = (d_297_v72_).m2(d_226_globalState_)
            d_350_v116_ = out20_
            d_351_v117_ = out21_
            d_352_v118_ = out22_
            d_353_v119_ = out23_
            d_354_v120_: _dafny.Map
            d_354_v120_ = _dafny.Map({d_222_v11_: d_213_v2_})
            d_354_v120_ = (d_354_v120_).set((d_263_v45_).fm3(d_226_globalState_), d_291_v66_)
            d_355_v121_: _dafny.Array
            nw47_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 28)
            d_355_v121_ = nw47_
            index37_ = default__.safeIndex(591, (d_355_v121_).length(0))
            (d_355_v121_)[index37_] = d_222_v11_
            d_356_v122_: _dafny.Map
            d_356_v122_ = _dafny.Map({(d_353_v119_) or (d_216_v5_): d_355_v121_})
            d_357_v123_: C2
            nw48_ = C2()
            nw48_.ctor__(not(d_295_v70_), d_295_v70_)
            d_357_v123_ = nw48_
            d_358_v124_: D34
            d_358_v124_ = D34_DC82(d_357_v123_)
            d_359_v125_: _dafny.MultiSet
            d_359_v125_ = _dafny.MultiSet([d_357_v123_, d_357_v123_, d_357_v123_, (d_358_v124_).cf135, d_357_v123_])
            index38_ = default__.safeIndex(591, (d_355_v121_).length(0))
            rhs48_ = ((d_356_v122_)[False] if (False) in (d_356_v122_) else d_355_v121_)
            rhs49_ = d_222_v11_
            rhs50_ = (((d_359_v125_) | (d_359_v125_)).cardinality) - (default__.safeModuloInt(len(d_221_v10_), (0) - (d_291_v66_)))
            lhs36_ = d_355_v121_
            lhs37_ = default__.safeIndex(591, (d_355_v121_).length(0))
            d_355_v121_ = rhs48_
            lhs36_[lhs37_] = rhs49_
            d_292_v67_ = rhs50_
            index39_ = default__.safeIndex(61, (d_233_v19_).length(0))
            (d_233_v19_)[index39_] = d_353_v119_
            d_360_v126_: C7
            nw49_ = C7()
            nw49_.ctor__(d_353_v119_, d_292_v67_, (d_357_v123_).fm9(d_213_v2_, d_226_globalState_), 966, (d_263_v45_).f27)
            d_360_v126_ = nw49_
        _dafny.print(_dafny.string_of((d_210_v1_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_213_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_214_v3_) == (_dafny.Set({170}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_215_v4_) == (_dafny.Map({_dafny.Set({170}): _dafny.Set({170})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_216_v5_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_217_v6_) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_218_v7_) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_219_v8_) == (_dafny.Map({True: -170, False: 4}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_221_v10_) == (_dafny.SeqWithoutIsStrInference([170]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_222_v11_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_223_v12_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_224_v15_) == (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([True])])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_225_v16_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_226_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_226_globalState_).f1)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_226_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_226_globalState_.f3) == (_dafny.Map({_dafny.Set({170}): _dafny.Set({170})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_226_globalState_).f4) == (_dafny.Set({170}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_226_globalState_.f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_226_globalState_.f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_226_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_226_globalState_).f8) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_226_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_226_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_226_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_226_globalState_.f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_226_globalState_.f13) == (_dafny.Map({True: -170}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_226_globalState_.f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_226_globalState_).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_226_globalState_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_226_globalState_).f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_226_globalState_).f19) == (_dafny.Set({170}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_226_globalState_).f20))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_226_globalState_.f21))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_226_globalState_).f22)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_226_globalState_).f23) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_226_globalState_).f24) == (_dafny.Map({True: -170}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_226_globalState_.f25))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_226_globalState_).f26))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_230_v17_) == (_dafny.Map({170: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_231_v18_) == (_dafny.MultiSet([1, 170, 554]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_233_v19_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_233_v19_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_233_v19_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_233_v19_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_233_v19_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_233_v19_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_233_v19_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_234_v20_)[1]) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_235_i3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_263_v45_).f27))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_264_v46_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_265_v47_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_266_i4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_279_v58_) == (_dafny.Map({80: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_283_v61_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_291_v66_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_292_v67_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_293_v68_).cf4).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_294_v69_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_295_v70_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_296_v71_) == (_dafny.MultiSet([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_297_v72_).f37) == (_dafny.MultiSet([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_298_v73_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_299_v74_) == (_dafny.Map({7: _dafny.CodePoint('s')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_300_v75_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({7: _dafny.CodePoint('s')})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_301_v77_)) == (_dafny.Set({_dafny.Map({7: _dafny.CodePoint('s')})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_303_v78_) == (_dafny.SeqWithoutIsStrInference([_dafny.Set({_dafny.Map({7: _dafny.CodePoint('s')})}), _dafny.Set({_dafny.Map({7: _dafny.CodePoint('s')})})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_304_v79_) == (_dafny.Map({0: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "whubvvjxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_306_v80_)) == (_dafny.MultiSet([_dafny.Set({527, -462})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_307_v81_) == (_dafny.MultiSet([_dafny.Set({170})]))))
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
    def is_DC2(self) -> bool:
        return isinstance(self, D0_DC2)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D0_DC3)

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC3(D0, NamedTuple('DC3', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC3({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC3) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)

class D1_DC4(D1, NamedTuple('DC4', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC6(int(0), _dafny.Seq({}), False, _dafny.Seq({}), False)
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

class D2_DC6(D2, NamedTuple('DC6', [('cf7', Any), ('cf8', Any), ('cf9', Any), ('cf10', Any), ('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC9(False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)

class D3_DC9(D3, NamedTuple('DC9', [('cf14', Any), ('cf15', Any), ('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf14 == __o.cf14 and self.cf15 == __o.cf15 and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC8(D3, NamedTuple('DC8', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)

class D4_DC10(D4, NamedTuple('DC10', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC12(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D5_DC12)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D5_DC11)

class D5_DC12(D5, NamedTuple('DC12', [('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC11(D5, NamedTuple('DC11', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC11({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC11) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC14(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D6_DC14)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D6_DC13)

class D6_DC14(D6, NamedTuple('DC14', [('cf22', Any), ('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC14({_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC14) and self.cf22 == __o.cf22 and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC13(D6, NamedTuple('DC13', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC13({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC13) and self.cf21 == __o.cf21
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
    def is_DC17(self) -> bool:
        return isinstance(self, D7_DC17)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D7_DC15)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D7_DC18)

class D7_DC16(D7, NamedTuple('DC16', [])):
    def __dafnystr__(self) -> str:
        return f'D7.DC16'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC16)
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC17(D7, NamedTuple('DC17', [('cf25', Any), ('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC17({_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC17) and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC15(D7, NamedTuple('DC15', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC15({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC15) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC18(D7, NamedTuple('DC18', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D8_DC19)

class D8_DC19(D8, NamedTuple('DC19', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC19({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC19) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC21(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D9_DC21)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D9_DC20)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D9_DC22)

class D9_DC21(D9, NamedTuple('DC21', [('cf31', Any), ('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC21({_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC21) and self.cf31 == __o.cf31 and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC20(D9, NamedTuple('DC20', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC20({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC20) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC22(D9, NamedTuple('DC22', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC22({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC22) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D10_DC23)

class D10_DC23(D10, NamedTuple('DC23', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC23({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC23) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC25()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D11_DC25)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D11_DC24)

class D11_DC25(D11, NamedTuple('DC25', [])):
    def __dafnystr__(self) -> str:
        return f'D11.DC25'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC25)
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC24(D11, NamedTuple('DC24', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC24({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC24) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.CodePoint('D')
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D12_DC26)

class D12_DC26(D12, NamedTuple('DC26', [('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC26({_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC26) and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC28(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D13_DC28)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D13_DC29)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D13_DC27)

class D13_DC28(D13, NamedTuple('DC28', [('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC28({_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC28) and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC29(D13, NamedTuple('DC29', [('cf39', Any), ('cf40', Any), ('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC29({_dafny.string_of(self.cf39)}, {_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC29) and self.cf39 == __o.cf39 and self.cf40 == __o.cf40 and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC27(D13, NamedTuple('DC27', [('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC27({_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC27) and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC31(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D14_DC31)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D14_DC30)

class D14_DC31(D14, NamedTuple('DC31', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC31({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC31) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC30(D14, NamedTuple('DC30', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC30({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC30) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC33()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D15_DC33)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D15_DC34)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D15_DC32)

class D15_DC33(D15, NamedTuple('DC33', [])):
    def __dafnystr__(self) -> str:
        return f'D15.DC33'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC33)
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC34(D15, NamedTuple('DC34', [('cf45', Any), ('cf46', Any), ('cf47', Any), ('cf48', Any), ('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC34({_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)}, {_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC34) and self.cf45 == __o.cf45 and self.cf46 == __o.cf46 and self.cf47 == __o.cf47 and self.cf48 == __o.cf48 and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC32(D15, NamedTuple('DC32', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC32({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC32) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC36(_dafny.Array(None, 0))
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

class D16_DC36(D16, NamedTuple('DC36', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC36({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC36) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC35(D16, NamedTuple('DC35', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC35({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC35) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC37(D16, NamedTuple('DC37', [('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC37({_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC37) and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC39(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D17_DC39)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D17_DC38)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D17_DC40)

class D17_DC39(D17, NamedTuple('DC39', [('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC39({self.cf54.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC39) and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC38(D17, NamedTuple('DC38', [('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC38({_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC38) and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC40(D17, NamedTuple('DC40', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC40({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC40) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC42(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D18_DC42)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D18_DC43)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D18_DC44)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D18_DC41)

class D18_DC42(D18, NamedTuple('DC42', [('cf57', Any), ('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC42({self.cf57.VerbatimString(True)}, {_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC42) and self.cf57 == __o.cf57 and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC43(D18, NamedTuple('DC43', [('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC43({_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC43) and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC44(D18, NamedTuple('DC44', [('cf60', Any), ('cf61', Any), ('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC44({_dafny.string_of(self.cf60)}, {_dafny.string_of(self.cf61)}, {_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC44) and self.cf60 == __o.cf60 and self.cf61 == __o.cf61 and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC41(D18, NamedTuple('DC41', [('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC41({_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC41) and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC46(int(0), _dafny.Array(None, 0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D19_DC46)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D19_DC45)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D19_DC47)

class D19_DC46(D19, NamedTuple('DC46', [('cf64', Any), ('cf65', Any), ('cf66', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC46({_dafny.string_of(self.cf64)}, {_dafny.string_of(self.cf65)}, {_dafny.string_of(self.cf66)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC46) and self.cf64 == __o.cf64 and self.cf65 == __o.cf65 and self.cf66 == __o.cf66
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC45(D19, NamedTuple('DC45', [('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC45({_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC45) and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC47(D19, NamedTuple('DC47', [('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC47({_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC47) and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC49(_dafny.CodePoint('D'), int(0), int(0), None, _dafny.Set({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D20_DC49)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D20_DC50)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D20_DC48)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D20_DC51)

class D20_DC49(D20, NamedTuple('DC49', [('cf69', Any), ('cf70', Any), ('cf71', Any), ('cf72', Any), ('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC49({_dafny.string_of(self.cf69)}, {_dafny.string_of(self.cf70)}, {_dafny.string_of(self.cf71)}, {_dafny.string_of(self.cf72)}, {_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC49) and self.cf69 == __o.cf69 and self.cf70 == __o.cf70 and self.cf71 == __o.cf71 and self.cf72 == __o.cf72 and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC50(D20, NamedTuple('DC50', [])):
    def __dafnystr__(self) -> str:
        return f'D20.DC50'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC50)
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC48(D20, NamedTuple('DC48', [('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC48({_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC48) and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC51(D20, NamedTuple('DC51', [('cf74', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC51({_dafny.string_of(self.cf74)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC51) and self.cf74 == __o.cf74
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC53(False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D21_DC53)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D21_DC54)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D21_DC55)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D21_DC52)

class D21_DC53(D21, NamedTuple('DC53', [('cf76', Any), ('cf77', Any), ('cf78', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC53({_dafny.string_of(self.cf76)}, {_dafny.string_of(self.cf77)}, {_dafny.string_of(self.cf78)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC53) and self.cf76 == __o.cf76 and self.cf77 == __o.cf77 and self.cf78 == __o.cf78
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC54(D21, NamedTuple('DC54', [('cf79', Any), ('cf80', Any), ('cf81', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC54({_dafny.string_of(self.cf79)}, {_dafny.string_of(self.cf80)}, {_dafny.string_of(self.cf81)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC54) and self.cf79 == __o.cf79 and self.cf80 == __o.cf80 and self.cf81 == __o.cf81
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC55(D21, NamedTuple('DC55', [('cf82', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC55({_dafny.string_of(self.cf82)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC55) and self.cf82 == __o.cf82
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC52(D21, NamedTuple('DC52', [('cf75', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC52({_dafny.string_of(self.cf75)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC52) and self.cf75 == __o.cf75
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC57(False, _dafny.CodePoint('D'), False)
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

class D22_DC57(D22, NamedTuple('DC57', [('cf84', Any), ('cf85', Any), ('cf86', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC57({_dafny.string_of(self.cf84)}, {_dafny.string_of(self.cf85)}, {_dafny.string_of(self.cf86)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC57) and self.cf84 == __o.cf84 and self.cf85 == __o.cf85 and self.cf86 == __o.cf86
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC58(D22, NamedTuple('DC58', [('cf87', Any), ('cf88', Any), ('cf89', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC58({_dafny.string_of(self.cf87)}, {_dafny.string_of(self.cf88)}, {_dafny.string_of(self.cf89)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC58) and self.cf87 == __o.cf87 and self.cf88 == __o.cf88 and self.cf89 == __o.cf89
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC56(D22, NamedTuple('DC56', [('cf83', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC56({_dafny.string_of(self.cf83)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC56) and self.cf83 == __o.cf83
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: D23_DC60(False, _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D23_DC60)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D23_DC61)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D23_DC59)

class D23_DC60(D23, NamedTuple('DC60', [('cf91', Any), ('cf92', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC60({_dafny.string_of(self.cf91)}, {_dafny.string_of(self.cf92)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC60) and self.cf91 == __o.cf91 and self.cf92 == __o.cf92
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC61(D23, NamedTuple('DC61', [('cf93', Any), ('cf94', Any), ('cf95', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC61({_dafny.string_of(self.cf93)}, {_dafny.string_of(self.cf94)}, {_dafny.string_of(self.cf95)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC61) and self.cf93 == __o.cf93 and self.cf94 == __o.cf94 and self.cf95 == __o.cf95
    def __hash__(self) -> int:
        return super().__hash__()

class D23_DC59(D23, NamedTuple('DC59', [('cf90', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC59({_dafny.string_of(self.cf90)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC59) and self.cf90 == __o.cf90
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D24_DC62)

class D24_DC62(D24, NamedTuple('DC62', [('cf96', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC62({_dafny.string_of(self.cf96)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC62) and self.cf96 == __o.cf96
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: D25_DC64(False, int(0), _dafny.CodePoint('D'), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC64(self) -> bool:
        return isinstance(self, D25_DC64)
    @property
    def is_DC63(self) -> bool:
        return isinstance(self, D25_DC63)

class D25_DC64(D25, NamedTuple('DC64', [('cf98', Any), ('cf99', Any), ('cf100', Any), ('cf101', Any), ('cf102', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC64({_dafny.string_of(self.cf98)}, {_dafny.string_of(self.cf99)}, {_dafny.string_of(self.cf100)}, {_dafny.string_of(self.cf101)}, {_dafny.string_of(self.cf102)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC64) and self.cf98 == __o.cf98 and self.cf99 == __o.cf99 and self.cf100 == __o.cf100 and self.cf101 == __o.cf101 and self.cf102 == __o.cf102
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC63(D25, NamedTuple('DC63', [('cf97', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC63({_dafny.string_of(self.cf97)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC63) and self.cf97 == __o.cf97
    def __hash__(self) -> int:
        return super().__hash__()


class D26:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC65(self) -> bool:
        return isinstance(self, D26_DC65)

class D26_DC65(D26, NamedTuple('DC65', [('cf103', Any)])):
    def __dafnystr__(self) -> str:
        return f'D26.DC65({_dafny.string_of(self.cf103)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D26_DC65) and self.cf103 == __o.cf103
    def __hash__(self) -> int:
        return super().__hash__()


class D27:
    @classmethod
    def default(cls, ):
        return lambda: D27_DC67(_dafny.Seq({}), _dafny.Set({}), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC67(self) -> bool:
        return isinstance(self, D27_DC67)
    @property
    def is_DC66(self) -> bool:
        return isinstance(self, D27_DC66)

class D27_DC67(D27, NamedTuple('DC67', [('cf105', Any), ('cf106', Any), ('cf107', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC67({_dafny.string_of(self.cf105)}, {_dafny.string_of(self.cf106)}, {_dafny.string_of(self.cf107)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC67) and self.cf105 == __o.cf105 and self.cf106 == __o.cf106 and self.cf107 == __o.cf107
    def __hash__(self) -> int:
        return super().__hash__()

class D27_DC66(D27, NamedTuple('DC66', [('cf104', Any)])):
    def __dafnystr__(self) -> str:
        return f'D27.DC66({_dafny.string_of(self.cf104)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D27_DC66) and self.cf104 == __o.cf104
    def __hash__(self) -> int:
        return super().__hash__()


class D28:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC68(self) -> bool:
        return isinstance(self, D28_DC68)

class D28_DC68(D28, NamedTuple('DC68', [('cf108', Any)])):
    def __dafnystr__(self) -> str:
        return f'D28.DC68({_dafny.string_of(self.cf108)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D28_DC68) and self.cf108 == __o.cf108
    def __hash__(self) -> int:
        return super().__hash__()


class D29:
    @classmethod
    def default(cls, ):
        return lambda: D29_DC70(D11.default()(), int(0), _dafny.Array(None, 0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC70(self) -> bool:
        return isinstance(self, D29_DC70)
    @property
    def is_DC71(self) -> bool:
        return isinstance(self, D29_DC71)
    @property
    def is_DC69(self) -> bool:
        return isinstance(self, D29_DC69)
    @property
    def is_DC72(self) -> bool:
        return isinstance(self, D29_DC72)

class D29_DC70(D29, NamedTuple('DC70', [('cf110', Any), ('cf111', Any), ('cf112', Any), ('cf113', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC70({_dafny.string_of(self.cf110)}, {_dafny.string_of(self.cf111)}, {_dafny.string_of(self.cf112)}, {_dafny.string_of(self.cf113)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC70) and self.cf110 == __o.cf110 and self.cf111 == __o.cf111 and self.cf112 == __o.cf112 and self.cf113 == __o.cf113
    def __hash__(self) -> int:
        return super().__hash__()

class D29_DC71(D29, NamedTuple('DC71', [('cf114', Any), ('cf115', Any), ('cf116', Any), ('cf117', Any), ('cf118', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC71({_dafny.string_of(self.cf114)}, {_dafny.string_of(self.cf115)}, {_dafny.string_of(self.cf116)}, {_dafny.string_of(self.cf117)}, {_dafny.string_of(self.cf118)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC71) and self.cf114 == __o.cf114 and self.cf115 == __o.cf115 and self.cf116 == __o.cf116 and self.cf117 == __o.cf117 and self.cf118 == __o.cf118
    def __hash__(self) -> int:
        return super().__hash__()

class D29_DC69(D29, NamedTuple('DC69', [('cf109', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC69({_dafny.string_of(self.cf109)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC69) and self.cf109 == __o.cf109
    def __hash__(self) -> int:
        return super().__hash__()

class D29_DC72(D29, NamedTuple('DC72', [('cf119', Any)])):
    def __dafnystr__(self) -> str:
        return f'D29.DC72({_dafny.string_of(self.cf119)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D29_DC72) and self.cf119 == __o.cf119
    def __hash__(self) -> int:
        return super().__hash__()


class D30:
    @classmethod
    def default(cls, ):
        return lambda: D30_DC74(False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC74(self) -> bool:
        return isinstance(self, D30_DC74)
    @property
    def is_DC75(self) -> bool:
        return isinstance(self, D30_DC75)
    @property
    def is_DC76(self) -> bool:
        return isinstance(self, D30_DC76)
    @property
    def is_DC73(self) -> bool:
        return isinstance(self, D30_DC73)

class D30_DC74(D30, NamedTuple('DC74', [('cf121', Any), ('cf122', Any), ('cf123', Any)])):
    def __dafnystr__(self) -> str:
        return f'D30.DC74({_dafny.string_of(self.cf121)}, {_dafny.string_of(self.cf122)}, {_dafny.string_of(self.cf123)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D30_DC74) and self.cf121 == __o.cf121 and self.cf122 == __o.cf122 and self.cf123 == __o.cf123
    def __hash__(self) -> int:
        return super().__hash__()

class D30_DC75(D30, NamedTuple('DC75', [('cf124', Any), ('cf125', Any), ('cf126', Any), ('cf127', Any)])):
    def __dafnystr__(self) -> str:
        return f'D30.DC75({_dafny.string_of(self.cf124)}, {self.cf125.VerbatimString(True)}, {_dafny.string_of(self.cf126)}, {_dafny.string_of(self.cf127)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D30_DC75) and self.cf124 == __o.cf124 and self.cf125 == __o.cf125 and self.cf126 == __o.cf126 and self.cf127 == __o.cf127
    def __hash__(self) -> int:
        return super().__hash__()

class D30_DC76(D30, NamedTuple('DC76', [])):
    def __dafnystr__(self) -> str:
        return f'D30.DC76'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D30_DC76)
    def __hash__(self) -> int:
        return super().__hash__()

class D30_DC73(D30, NamedTuple('DC73', [('cf120', Any)])):
    def __dafnystr__(self) -> str:
        return f'D30.DC73({_dafny.string_of(self.cf120)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D30_DC73) and self.cf120 == __o.cf120
    def __hash__(self) -> int:
        return super().__hash__()


class D31:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC77(self) -> bool:
        return isinstance(self, D31_DC77)

class D31_DC77(D31, NamedTuple('DC77', [('cf128', Any)])):
    def __dafnystr__(self) -> str:
        return f'D31.DC77({_dafny.string_of(self.cf128)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D31_DC77) and self.cf128 == __o.cf128
    def __hash__(self) -> int:
        return super().__hash__()


class D32:
    @classmethod
    def default(cls, ):
        return lambda: D32_DC79(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC79(self) -> bool:
        return isinstance(self, D32_DC79)
    @property
    def is_DC80(self) -> bool:
        return isinstance(self, D32_DC80)
    @property
    def is_DC78(self) -> bool:
        return isinstance(self, D32_DC78)

class D32_DC79(D32, NamedTuple('DC79', [('cf130', Any), ('cf131', Any)])):
    def __dafnystr__(self) -> str:
        return f'D32.DC79({_dafny.string_of(self.cf130)}, {_dafny.string_of(self.cf131)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D32_DC79) and self.cf130 == __o.cf130 and self.cf131 == __o.cf131
    def __hash__(self) -> int:
        return super().__hash__()

class D32_DC80(D32, NamedTuple('DC80', [('cf132', Any), ('cf133', Any)])):
    def __dafnystr__(self) -> str:
        return f'D32.DC80({_dafny.string_of(self.cf132)}, {_dafny.string_of(self.cf133)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D32_DC80) and self.cf132 == __o.cf132 and self.cf133 == __o.cf133
    def __hash__(self) -> int:
        return super().__hash__()

class D32_DC78(D32, NamedTuple('DC78', [('cf129', Any)])):
    def __dafnystr__(self) -> str:
        return f'D32.DC78({_dafny.string_of(self.cf129)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D32_DC78) and self.cf129 == __o.cf129
    def __hash__(self) -> int:
        return super().__hash__()


class D33:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC81(self) -> bool:
        return isinstance(self, D33_DC81)

class D33_DC81(D33, NamedTuple('DC81', [('cf134', Any)])):
    def __dafnystr__(self) -> str:
        return f'D33.DC81({_dafny.string_of(self.cf134)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D33_DC81) and self.cf134 == __o.cf134
    def __hash__(self) -> int:
        return super().__hash__()


class D34:
    @classmethod
    def default(cls, ):
        return lambda: D34_DC83(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC83(self) -> bool:
        return isinstance(self, D34_DC83)
    @property
    def is_DC84(self) -> bool:
        return isinstance(self, D34_DC84)
    @property
    def is_DC85(self) -> bool:
        return isinstance(self, D34_DC85)
    @property
    def is_DC82(self) -> bool:
        return isinstance(self, D34_DC82)

class D34_DC83(D34, NamedTuple('DC83', [('cf136', Any)])):
    def __dafnystr__(self) -> str:
        return f'D34.DC83({_dafny.string_of(self.cf136)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D34_DC83) and self.cf136 == __o.cf136
    def __hash__(self) -> int:
        return super().__hash__()

class D34_DC84(D34, NamedTuple('DC84', [('cf137', Any), ('cf138', Any), ('cf139', Any), ('cf140', Any)])):
    def __dafnystr__(self) -> str:
        return f'D34.DC84({_dafny.string_of(self.cf137)}, {_dafny.string_of(self.cf138)}, {_dafny.string_of(self.cf139)}, {_dafny.string_of(self.cf140)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D34_DC84) and self.cf137 == __o.cf137 and self.cf138 == __o.cf138 and self.cf139 == __o.cf139 and self.cf140 == __o.cf140
    def __hash__(self) -> int:
        return super().__hash__()

class D34_DC85(D34, NamedTuple('DC85', [('cf141', Any), ('cf142', Any), ('cf143', Any)])):
    def __dafnystr__(self) -> str:
        return f'D34.DC85({_dafny.string_of(self.cf141)}, {_dafny.string_of(self.cf142)}, {_dafny.string_of(self.cf143)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D34_DC85) and self.cf141 == __o.cf141 and self.cf142 == __o.cf142 and self.cf143 == __o.cf143
    def __hash__(self) -> int:
        return super().__hash__()

class D34_DC82(D34, NamedTuple('DC82', [('cf135', Any)])):
    def __dafnystr__(self) -> str:
        return f'D34.DC82({_dafny.string_of(self.cf135)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D34_DC82) and self.cf135 == __o.cf135
    def __hash__(self) -> int:
        return super().__hash__()


class D35:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC86(self) -> bool:
        return isinstance(self, D35_DC86)

class D35_DC86(D35, NamedTuple('DC86', [('cf144', Any)])):
    def __dafnystr__(self) -> str:
        return f'D35.DC86({_dafny.string_of(self.cf144)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D35_DC86) and self.cf144 == __o.cf144
    def __hash__(self) -> int:
        return super().__hash__()


class D36:
    @classmethod
    def default(cls, ):
        return lambda: D36_DC88(False, False, None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC88(self) -> bool:
        return isinstance(self, D36_DC88)
    @property
    def is_DC89(self) -> bool:
        return isinstance(self, D36_DC89)
    @property
    def is_DC87(self) -> bool:
        return isinstance(self, D36_DC87)

class D36_DC88(D36, NamedTuple('DC88', [('cf146', Any), ('cf147', Any), ('cf148', Any)])):
    def __dafnystr__(self) -> str:
        return f'D36.DC88({_dafny.string_of(self.cf146)}, {_dafny.string_of(self.cf147)}, {_dafny.string_of(self.cf148)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D36_DC88) and self.cf146 == __o.cf146 and self.cf147 == __o.cf147 and self.cf148 == __o.cf148
    def __hash__(self) -> int:
        return super().__hash__()

class D36_DC89(D36, NamedTuple('DC89', [('cf149', Any), ('cf150', Any), ('cf151', Any)])):
    def __dafnystr__(self) -> str:
        return f'D36.DC89({self.cf149.VerbatimString(True)}, {_dafny.string_of(self.cf150)}, {_dafny.string_of(self.cf151)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D36_DC89) and self.cf149 == __o.cf149 and self.cf150 == __o.cf150 and self.cf151 == __o.cf151
    def __hash__(self) -> int:
        return super().__hash__()

class D36_DC87(D36, NamedTuple('DC87', [('cf145', Any)])):
    def __dafnystr__(self) -> str:
        return f'D36.DC87({_dafny.string_of(self.cf145)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D36_DC87) and self.cf145 == __o.cf145
    def __hash__(self) -> int:
        return super().__hash__()


class D37:
    @classmethod
    def default(cls, ):
        return lambda: D37_DC91(False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC91(self) -> bool:
        return isinstance(self, D37_DC91)
    @property
    def is_DC92(self) -> bool:
        return isinstance(self, D37_DC92)
    @property
    def is_DC90(self) -> bool:
        return isinstance(self, D37_DC90)

class D37_DC91(D37, NamedTuple('DC91', [('cf153', Any), ('cf154', Any), ('cf155', Any)])):
    def __dafnystr__(self) -> str:
        return f'D37.DC91({_dafny.string_of(self.cf153)}, {_dafny.string_of(self.cf154)}, {_dafny.string_of(self.cf155)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D37_DC91) and self.cf153 == __o.cf153 and self.cf154 == __o.cf154 and self.cf155 == __o.cf155
    def __hash__(self) -> int:
        return super().__hash__()

class D37_DC92(D37, NamedTuple('DC92', [('cf156', Any)])):
    def __dafnystr__(self) -> str:
        return f'D37.DC92({_dafny.string_of(self.cf156)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D37_DC92) and self.cf156 == __o.cf156
    def __hash__(self) -> int:
        return super().__hash__()

class D37_DC90(D37, NamedTuple('DC90', [('cf152', Any)])):
    def __dafnystr__(self) -> str:
        return f'D37.DC90({_dafny.string_of(self.cf152)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D37_DC90) and self.cf152 == __o.cf152
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m1(self, globalState):
        pass

    def m2(self, globalState):
        pass


class T1:
    pass
    def m3(self, p0, p1, p2, globalState):
        pass


class T2:
    pass
    def m4(self, p0, p1, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f2: int = int(0)
        self.f3: _dafny.Map = _dafny.Map({})
        self.f5: bool = False
        self.f6: int = int(0)
        self.f12: bool = False
        self.f13: _dafny.Map = _dafny.Map({})
        self.f14: bool = False
        self.f21: bool = False
        self.f25: bool = False
        self._f0: int = int(0)
        self._f1: _dafny.Array = _dafny.Array(None, 0)
        self._f4: _dafny.Set = _dafny.Set({})
        self._f7: bool = False
        self._f8: _dafny.Map = _dafny.Map({})
        self._f9: int = int(0)
        self._f10: bool = False
        self._f11: bool = False
        self._f15: bool = False
        self._f16: _dafny.Array = _dafny.Array(None, 0)
        self._f17: bool = False
        self._f18: bool = False
        self._f19: _dafny.Set = _dafny.Set({})
        self._f20: int = int(0)
        self._f22: _dafny.Array = _dafny.Array(None, 0)
        self._f23: _dafny.Seq = _dafny.Seq({})
        self._f24: _dafny.Map = _dafny.Map({})
        self._f26: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25, f26):
        (self)._f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
        (self).f3 = f3
        (self)._f4 = f4
        (self).f5 = f5
        (self).f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self).f12 = f12
        (self).f13 = f13
        (self).f14 = f14
        (self)._f15 = f15
        (self)._f16 = f16
        (self)._f17 = f17
        (self)._f18 = f18
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
    def f4(self):
        return self._f4
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
    def f15(self):
        return self._f15
    @property
    def f16(self):
        return self._f16
    @property
    def f17(self):
        return self._f17
    @property
    def f18(self):
        return self._f18
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

class C0:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass


class C1(T2, T0):
    def  __init__(self):
        self._f27: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f27(self):
        return self._f27
    def ctor__(self, f27):
        (self)._f27 = f27

    def fm4(self, p0, globalState):
        return (self).f27

    def fm5(self, globalState):
        return D0_DC2((0) - (-728), default__.safeDivisionInt(-41, 424))

    def fm2(self, p0, p1, globalState):
        source10_ = D0_DC1((self).f27)
        if source10_.is_DC1:
            d_361___mcc_h0_ = source10_.cf1
            d_362_cf1_ = d_361___mcc_h0_
            return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_363_i0_ in range(default__.abs(965))])
        elif source10_.is_DC2:
            d_364___mcc_h1_ = source10_.cf2
            d_365___mcc_h2_ = source10_.cf3
            d_366_cf3_ = d_365___mcc_h2_
            d_367_cf2_ = d_364___mcc_h1_
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ui"))
        elif source10_.is_DC0:
            d_368___mcc_h3_ = source10_.cf0
            d_369_cf0_ = d_368___mcc_h3_
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jmsxq"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bebf")))
        elif True:
            d_370___mcc_h4_ = source10_.cf4
            d_371_cf4_ = d_370___mcc_h4_
            return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_372_i1_ in range(default__.abs(769))])

    def fm3(self, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_373_i0_ in range(default__.abs(33))]) if (self).f27 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l")))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_374_i1_ in range(default__.abs(853))])))

    def m4(self, p0, p1, globalState):
        r0: bool = False
        r1: bool = False
        source11_ = p0
        if source11_.is_DC1:
            d_375___mcc_h0_ = source11_.cf1
            d_376_cf1_ = d_375___mcc_h0_
            d_377_v0_: int
            d_377_v0_ = 902
            d_378_v1_: _dafny.Seq
            d_378_v1_ = _dafny.SeqWithoutIsStrInference([d_377_v0_, d_377_v0_])
            d_379_v2_: str
            d_379_v2_ = _dafny.CodePoint('v')
            d_380_v3_: _dafny.Map
            d_380_v3_ = _dafny.Map({d_378_v1_: d_379_v2_})
            d_380_v3_ = (d_380_v3_).set((d_378_v1_) + (d_378_v1_), d_379_v2_)
            d_381_v4_: _dafny.Map
            d_381_v4_ = _dafny.Map({d_376_cf1_: d_377_v0_})
            d_382_v5_: _dafny.Seq
            d_382_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dcrvqx"))
            if (d_377_v0_) < (((d_381_v4_)[p1] if (p1) in (d_381_v4_) else len(d_382_v5_))):
                d_383_v6_: _dafny.MultiSet
                d_383_v6_ = _dafny.MultiSet([d_379_v2_])
                d_384_v7_: _dafny.Seq
                d_384_v7_ = _dafny.SeqWithoutIsStrInference([d_383_v6_])
                d_385_v8_: _dafny.Set
                d_385_v8_ = _dafny.Set({462, d_377_v0_, d_377_v0_, (0) - (d_377_v0_)})
                d_386_v9_: _dafny.Seq
                d_386_v9_ = _dafny.SeqWithoutIsStrInference([p1])
                d_387_v10_: _dafny.Set
                d_387_v10_ = _dafny.Set({(self).f27, p1, False})
                d_388_v12_: _dafny.Array
                nw50_ = _dafny.Array(None, 28)
                nw50_[int(0)] = d_377_v0_
                nw50_[int(1)] = d_377_v0_
                nw50_[int(2)] = (D0_DC2(755, d_377_v0_)).cf2
                nw50_[int(3)] = len(d_382_v5_)
                nw50_[int(4)] = d_377_v0_
                nw50_[int(5)] = d_377_v0_
                nw50_[int(6)] = len(_dafny.Set({d_385_v8_, d_385_v8_}))
                nw50_[int(7)] = len(d_378_v1_)
                nw50_[int(8)] = len((_dafny.Map({len(d_386_v9_): d_387_v10_})).set(d_377_v0_, d_387_v10_))
                def iife55_():
                    coll37_ = _dafny.Map()
                    compr_37_: int
                    for compr_37_ in _dafny.IntegerRange(149, 49):
                        d_389_v11_: int = compr_37_
                        if ((149) <= (d_389_v11_)) and ((d_389_v11_) < (49)):
                            coll37_[(d_389_v11_) * (len(d_387_v10_))] = len(_dafny.Map({d_377_v0_: d_378_v1_}))
                    return _dafny.Map(coll37_)
                nw50_[int(9)] = default__.safeModuloInt(d_377_v0_, len(iife55_()
                ))
                nw50_[int(10)] = d_377_v0_
                nw50_[int(11)] = d_377_v0_
                nw50_[int(12)] = d_377_v0_
                nw50_[int(13)] = (0) - (d_377_v0_)
                nw50_[int(14)] = d_377_v0_
                nw50_[int(15)] = (d_377_v0_) * (len(d_382_v5_))
                nw50_[int(16)] = d_377_v0_
                nw50_[int(17)] = d_377_v0_
                nw50_[int(18)] = d_377_v0_
                nw50_[int(19)] = len(d_382_v5_)
                nw50_[int(20)] = 839
                nw50_[int(21)] = default__.safeDivisionInt((0) - (d_377_v0_), d_377_v0_)
                nw50_[int(22)] = default__.safeDivisionInt(len(d_387_v10_), d_377_v0_)
                nw50_[int(23)] = d_377_v0_
                nw50_[int(24)] = len(d_382_v5_)
                nw50_[int(25)] = ((d_378_v1_)[default__.safeIndex(d_377_v0_, len(d_378_v1_))]) - (d_377_v0_)
                nw50_[int(26)] = (0) - (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oplbec"))), d_377_v0_))
                nw50_[int(27)] = d_377_v0_
                d_388_v12_ = nw50_
                index40_ = default__.safeIndex(839, (d_388_v12_).length(0))
                (d_388_v12_)[index40_] = ((d_381_v4_)[False] if (False) in (d_381_v4_) else len(d_382_v5_))
                index41_ = default__.safeIndex(839, (d_388_v12_).length(0))
                rhs51_ = p1
                rhs52_ = (((d_384_v7_) + (d_384_v7_)) + (d_384_v7_)).set(default__.safeIndex(d_377_v0_, len(((d_384_v7_) + (d_384_v7_)) + (d_384_v7_))), (_dafny.MultiSet([d_379_v2_, d_379_v2_])).set(d_379_v2_, default__.abs(d_377_v0_)))
                rhs53_ = 589
                lhs38_ = globalState
                lhs39_ = d_388_v12_
                lhs40_ = default__.safeIndex(839, (d_388_v12_).length(0))
                lhs38_.f5 = rhs51_
                d_384_v7_ = rhs52_
                lhs39_[lhs40_] = rhs53_
                d_390_v13_: D0
                d_390_v13_ = D0_DC1(p1)
                def iife56_():
                    coll38_ = _dafny.Set()
                    compr_38_: int
                    for compr_38_ in (d_378_v1_).Elements:
                        d_391_v14_: int = compr_38_
                        if (d_391_v14_) in (d_378_v1_):
                            coll38_ = coll38_.union(_dafny.Set([(d_391_v14_) * (-665)]))
                    return _dafny.Set(coll38_)
                (globalState).f12 = default__.fm0(((self).fm3(globalState)) + (d_382_v5_), False, (self).fm4(d_390_v13_, globalState), (iife56_()
                ) - (d_385_v8_), globalState)
                (globalState).f12 = ((self).f27) or (p1)
                d_392_v15_: _dafny.Map
                d_392_v15_ = _dafny.Map({(self).f27: d_388_v12_})
                (globalState).f6 = (0) - (((d_377_v0_) + (len(d_392_v15_))) + ((d_388_v12_)[default__.safeIndex(839, (d_388_v12_).length(0))]))
                d_393_v16_: _dafny.Map
                d_393_v16_ = _dafny.Map({p1: p1})
                d_394_v17_: D0
                d_394_v17_ = D0_DC2(len(d_393_v16_), len(_dafny.SeqWithoutIsStrInference([(d_388_v12_)[default__.safeIndex(839, (d_388_v12_).length(0))] for d_395_i0_ in range(default__.abs(730))])))
                (globalState).f6 = (d_394_v17_).cf2
            elif True:
                d_396_v18_: _dafny.Map
                d_396_v18_ = _dafny.Map({d_377_v0_: True})
                d_397_v19_: _dafny.Set
                d_397_v19_ = _dafny.Set({d_396_v18_})
                (globalState).f2 = len(d_397_v19_)
                d_398_v20_: _dafny.Seq
                d_398_v20_ = _dafny.SeqWithoutIsStrInference([p1])
                d_399_v21_: D0
                d_399_v21_ = D0_DC0(len(d_398_v20_))
                pat_let_tv5_ = d_377_v0_
                def iife57_(_pat_let9_0):
                    def iife58_(d_400_dt__update__tmp_h0_):
                        def iife59_(_pat_let10_0):
                            def iife60_(d_401_dt__update_hcf0_h0_):
                                return D0_DC0(d_401_dt__update_hcf0_h0_)
                            return iife60_(_pat_let10_0)
                        return iife59_(pat_let_tv5_)
                    return iife58_(_pat_let9_0)
                d_399_v21_ = iife57_(d_399_v21_)
                d_402_v22_: _dafny.Array
                def lambda9_(d_403_v0_):
                    def lambda10_(d_404_i1_):
                        return _dafny.SeqWithoutIsStrInference([d_403_v0_ for d_405_i2_ in range(default__.abs(284))])

                    return lambda10_

                init5_ = lambda9_(d_377_v0_)
                nw51_ = _dafny.Array(None, 5)
                for i0_5_ in range(nw51_.length(0)):
                    nw51_[i0_5_] = init5_(i0_5_)
                d_402_v22_ = nw51_
                index42_ = default__.safeIndex(3, (d_402_v22_).length(0))
                (d_402_v22_)[index42_] = _dafny.SeqWithoutIsStrInference([d_377_v0_ for d_406_i3_ in range(default__.abs(218))])
                d_407_v23_: _dafny.Seq
                d_407_v23_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_377_v0_]), (d_378_v1_ if d_376_cf1_ else d_378_v1_), d_378_v1_])
                index43_ = default__.safeIndex(3, (d_402_v22_).length(0))
                rhs54_ = (((d_382_v5_) + (d_382_v5_)) + (d_382_v5_)).set(default__.safeIndex(d_377_v0_, len(((d_382_v5_) + (d_382_v5_)) + (d_382_v5_))), d_379_v2_)
                rhs55_ = (_dafny.SeqWithoutIsStrInference([d_377_v0_ for d_408_i4_ in range(default__.abs(591))])) + ((d_378_v1_) + (d_378_v1_))
                rhs56_ = d_407_v23_
                lhs41_ = d_402_v22_
                lhs42_ = default__.safeIndex(3, (d_402_v22_).length(0))
                d_382_v5_ = rhs54_
                lhs41_[lhs42_] = rhs55_
                d_407_v23_ = rhs56_
                d_409_v24_: _dafny.Array
                nw52_ = _dafny.Array(_dafny.Map({}), 11)
                d_409_v24_ = nw52_
                d_410_v26_: _dafny.Map
                def iife61_():
                    coll39_ = _dafny.Set()
                    compr_39_: int
                    for compr_39_ in _dafny.IntegerRange(259, 427):
                        d_411_v25_: int = compr_39_
                        if ((259) <= (d_411_v25_)) and ((d_411_v25_) < (427)):
                            coll39_ = coll39_.union(_dafny.Set([(d_411_v25_) + (d_377_v0_)]))
                    return _dafny.Set(coll39_)
                d_410_v26_ = _dafny.Map({d_376_cf1_: not(not(default__.fm0(d_382_v5_, (self).f27, p1, iife61_()
                , globalState)))})
                index44_ = default__.safeIndex(251, (d_409_v24_).length(0))
                (d_409_v24_)[index44_] = (d_410_v26_).set((self).f27, (self).f27)
                d_412_v27_: _dafny.Set
                d_412_v27_ = _dafny.Set({d_377_v0_, d_377_v0_})
                index45_ = default__.safeIndex(251, (d_409_v24_).length(0))
                (d_409_v24_)[index45_] = _dafny.Map({default__.fm0(d_382_v5_, False, (self).f27, d_412_v27_, globalState): p1})
                (globalState).f2 = ((d_377_v0_) - (len(d_382_v5_))) - (default__.safeModuloInt(d_377_v0_, len(default__.fm12((0) - (d_377_v0_), d_376_cf1_, True, globalState))))
            d_413_v28_: _dafny.Set
            d_413_v28_ = _dafny.Set({d_377_v0_})
            d_414_v30_: _dafny.Seq
            def iife62_():
                coll40_ = _dafny.Set()
                compr_40_: int
                for compr_40_ in (d_378_v1_).Elements:
                    d_415_v29_: int = compr_40_
                    if (d_415_v29_) in (d_378_v1_):
                        coll40_ = coll40_.union(_dafny.Set([default__.safeModuloInt(d_415_v29_, len(_dafny.SeqWithoutIsStrInference([not(False), False, False])))]))
                return _dafny.Set(coll40_)
            d_414_v30_ = _dafny.SeqWithoutIsStrInference([d_413_v28_, (d_413_v28_) - (d_413_v28_), (_dafny.Set({len(d_382_v5_), d_377_v0_})) - (d_413_v28_), d_413_v28_, iife62_()
            ])
            d_414_v30_ = ((_dafny.SeqWithoutIsStrInference([d_413_v28_ for d_416_i5_ in range(default__.abs(43))])) + (_dafny.SeqWithoutIsStrInference([d_413_v28_ for d_417_i6_ in range(default__.abs(551))])) if False else (_dafny.SeqWithoutIsStrInference([d_413_v28_])).set(default__.safeIndex(d_377_v0_, len(_dafny.SeqWithoutIsStrInference([d_413_v28_]))), (d_414_v30_)[default__.safeIndex(d_377_v0_, len(d_414_v30_))]))
            d_418_v31_: _dafny.MultiSet
            d_418_v31_ = _dafny.MultiSet([(self).f27, False, (self).f27])
            d_419_v32_: _dafny.Map
            d_419_v32_ = _dafny.Map({d_418_v31_: (665) * (len(d_382_v5_))})
            d_420_v33_: D0
            d_420_v33_ = D0_DC2(d_377_v0_, d_377_v0_)
            d_421_v35_: _dafny.Seq
            d_421_v35_ = _dafny.SeqWithoutIsStrInference([(self).f27, p1])
            def iife63_():
                coll41_ = _dafny.Map()
                compr_41_: _dafny.MultiSet
                for compr_41_ in (_dafny.Map({d_418_v31_: d_377_v0_})).keys.Elements:
                    d_422_v34_: _dafny.MultiSet = compr_41_
                    if (d_422_v34_) in (_dafny.Map({d_418_v31_: d_377_v0_})):
                        coll41_[d_422_v34_] = d_377_v0_
                return _dafny.Map(coll41_)
            rhs57_ = (d_420_v33_).cf2
            rhs58_ = (d_377_v0_) - (406)
            rhs59_ = ((_dafny.Map({d_418_v31_: d_377_v0_})) | (iife63_()
            )).set(d_418_v31_, default__.fm1(p1, d_421_v35_, (self).f27, globalState))
            rhs60_ = (d_378_v1_)[default__.safeIndex(d_377_v0_, len(d_378_v1_))]
            lhs43_ = globalState
            lhs44_ = globalState
            lhs45_ = globalState
            lhs43_.f2 = rhs57_
            lhs44_.f6 = rhs58_
            d_419_v32_ = rhs59_
            lhs45_.f6 = rhs60_
        elif source11_.is_DC2:
            d_423___mcc_h1_ = source11_.cf2
            d_424___mcc_h2_ = source11_.cf3
            d_425_cf3_ = d_424___mcc_h2_
            d_426_cf2_ = d_423___mcc_h1_
            d_427_v36_: _dafny.Array
            nw53_ = _dafny.Array(False, 14)
            d_427_v36_ = nw53_
            d_428_v37_: _dafny.Array
            nw54_ = _dafny.Array(_dafny.Array(None, 0), 14)
            d_428_v37_ = nw54_
            d_429_v38_: _dafny.Array
            nw55_ = _dafny.Array(int(0), 6)
            d_429_v38_ = nw55_
            d_430_v39_: _dafny.Array
            d_430_v39_ = d_429_v38_
            d_431_v40_: _dafny.Seq
            d_431_v40_ = _dafny.SeqWithoutIsStrInference([d_429_v38_, (d_430_v39_), d_429_v38_])
            index46_ = default__.safeIndex(85, (d_428_v37_).length(0))
            (d_428_v37_)[index46_] = (d_431_v40_)[default__.safeIndex(d_426_cf2_, len(d_431_v40_))]
            d_432_v41_: str
            d_432_v41_ = _dafny.CodePoint('l')
            d_433_v42_: _dafny.Map
            d_433_v42_ = _dafny.Map({d_425_cf3_: (self).fm3(globalState)})
            d_434_v43_: _dafny.Seq
            d_434_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "illg"))
            d_435_v44_: _dafny.MultiSet
            d_435_v44_ = _dafny.MultiSet([(d_432_v41_) in (((d_433_v42_)[(0) - (d_425_cf3_)] if ((0) - (d_425_cf3_)) in (d_433_v42_) else d_434_v43_)), (self).f27, (self).f27, p1])
            index47_ = default__.safeIndex(873, (d_429_v38_).length(0))
            (d_429_v38_)[index47_] = d_426_cf2_
            d_436_v45_: _dafny.Seq
            d_436_v45_ = _dafny.SeqWithoutIsStrInference([p1, (self).f27, not(True), not(p1)])
            index48_ = default__.safeIndex(85, (d_428_v37_).length(0))
            index49_ = default__.safeIndex(873, (d_429_v38_).length(0))
            rhs61_ = d_427_v36_
            rhs62_ = d_429_v38_
            rhs63_ = ((d_435_v44_).intersection(d_435_v44_)) | (_dafny.MultiSet(d_436_v45_))
            rhs64_ = d_425_cf3_
            lhs46_ = d_428_v37_
            lhs47_ = default__.safeIndex(85, (d_428_v37_).length(0))
            lhs48_ = d_429_v38_
            lhs49_ = default__.safeIndex(873, (d_429_v38_).length(0))
            d_427_v36_ = rhs61_
            lhs46_[lhs47_] = rhs62_
            d_435_v44_ = rhs63_
            lhs48_[lhs49_] = rhs64_
            d_437_v46_: _dafny.Seq
            d_437_v46_ = _dafny.SeqWithoutIsStrInference([(d_429_v38_)[default__.safeIndex(873, (d_429_v38_).length(0))], d_426_cf2_])
            d_438_v47_: _dafny.MultiSet
            d_438_v47_ = _dafny.MultiSet([(d_429_v38_)[default__.safeIndex(873, (d_429_v38_).length(0))], len(d_437_v46_), len(d_437_v46_)])
            d_439_v48_: _dafny.Seq
            d_439_v48_ = _dafny.SeqWithoutIsStrInference([d_438_v47_, default__.fm13(globalState)])
            d_440_v49_: _dafny.MultiSet
            d_440_v49_ = _dafny.MultiSet([_dafny.MultiSet(default__.fm12((d_429_v38_)[default__.safeIndex(873, (d_429_v38_).length(0))], True, p1, globalState)), d_438_v47_])
            d_441_v50_: _dafny.Map
            d_441_v50_ = _dafny.Map({(d_440_v49_).ispropersubset(_dafny.MultiSet((d_439_v48_).set(default__.safeIndex(default__.fm1(True, d_436_v45_, (self).f27, globalState), len(d_439_v48_)), default__.fm13(globalState)))): (False) and (p1)})
            d_441_v50_ = (d_441_v50_).set((p1) or (False), p1)
            (globalState).f12 = p1
            if p1:
                d_442_v51_: D0
                d_442_v51_ = D0_DC2(d_425_cf3_, ((d_429_v38_)[default__.safeIndex(873, (d_429_v38_).length(0))]) - ((d_429_v38_)[default__.safeIndex(873, (d_429_v38_).length(0))]))
                d_442_v51_ = D0_DC2((d_426_cf2_) * (d_426_cf2_), (d_425_cf3_) - (d_426_cf2_))
                (globalState).f6 = (d_429_v38_)[default__.safeIndex(873, (d_429_v38_).length(0))]
                d_432_v41_ = default__.fm14((d_429_v38_)[default__.safeIndex(873, (d_429_v38_).length(0))], globalState)
                index50_ = default__.safeIndex(873, (d_429_v38_).length(0))
                (d_429_v38_)[index50_] = d_426_cf2_
                (globalState).f14 = p1
            elif True:
                d_434_v43_ = ((d_434_v43_ if False else d_434_v43_)) + ((d_434_v43_) + (d_434_v43_))
                d_443_v52_: C0
                nw56_ = C0()
                nw56_.ctor__()
                d_443_v52_ = nw56_
                d_444_v53_: _dafny.Map
                d_444_v53_ = _dafny.Map({(d_429_v38_)[default__.safeIndex(873, (d_429_v38_).length(0))]: default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({824: (self).f27})) for d_445_i7_ in range(default__.abs(920))])), d_426_cf2_)})
                d_444_v53_ = (d_444_v53_).set(d_426_cf2_, len(_dafny.SeqWithoutIsStrInference([d_425_cf3_ for d_446_i8_ in range(default__.abs(507))])))
                d_447_v54_: _dafny.Seq
                d_447_v54_ = _dafny.SeqWithoutIsStrInference([d_430_v39_, d_430_v39_])
                d_448_v55_: _dafny.Seq
                d_448_v55_ = _dafny.SeqWithoutIsStrInference([d_434_v43_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bllcyut"))])
                d_449_v56_: _dafny.Map
                d_449_v56_ = _dafny.Map({(d_447_v54_) + (d_447_v54_): len((d_448_v55_)[default__.safeIndex((d_429_v38_)[default__.safeIndex(873, (d_429_v38_).length(0))], len(d_448_v55_))])})
                d_449_v56_ = (d_449_v56_).set(_dafny.SeqWithoutIsStrInference([d_430_v39_]), (d_429_v38_)[default__.safeIndex(873, (d_429_v38_).length(0))])
                d_450_v57_: int
                d_451_v58_: bool
                d_452_v59_: int
                d_453_v60_: str
                out24_: int
                out25_: bool
                out26_: int
                out27_: str
                out24_, out25_, out26_, out27_ = (self).m9(d_426_cf2_, globalState)
                d_450_v57_ = out24_
                d_451_v58_ = out25_
                d_452_v59_ = out26_
                d_453_v60_ = out27_
        elif source11_.is_DC0:
            d_454___mcc_h3_ = source11_.cf0
            d_455_cf0_ = d_454___mcc_h3_
            d_456_v61_: _dafny.Array
            nw57_ = _dafny.Array(_dafny.Array(None, 0), 2)
            d_456_v61_ = nw57_
            d_457_v62_: _dafny.Array
            nw58_ = _dafny.Array(int(0), 4)
            d_457_v62_ = nw58_
            index51_ = default__.safeIndex(564, (d_456_v61_).length(0))
            (d_456_v61_)[index51_] = d_457_v62_
            index52_ = default__.safeIndex(564, (d_456_v61_).length(0))
            (d_456_v61_)[index52_] = d_457_v62_
            (globalState).f5 = p1
            d_458_v63_: _dafny.Array
            d_458_v63_ = d_457_v62_
            index53_ = default__.safeIndex(564, (d_456_v61_).length(0))
            (d_456_v61_)[index53_] = (d_458_v63_)
            d_459_v64_: C0
            nw59_ = C0()
            nw59_.ctor__()
            d_459_v64_ = nw59_
        elif True:
            d_460___mcc_h4_ = source11_.cf4
            d_461_cf4_ = d_460___mcc_h4_
            d_462_v65_: int
            d_462_v65_ = 332
            (globalState).f6 = (d_462_v65_) * (416)
            d_463_v66_: _dafny.Array
            nw60_ = _dafny.Array(False, 18)
            d_463_v66_ = nw60_
            d_463_v66_ = d_463_v66_
            if not((p1 if (self).f27 else (p1) == (True))):
                d_464_v67_: _dafny.Array
                nw61_ = _dafny.Array(int(0), 17)
                d_464_v67_ = nw61_
                nw62_ = _dafny.Array(int(0), 7)
                d_464_v67_ = nw62_
                (globalState).f21 = (self).f27
                d_465_v68_: _dafny.Map
                d_465_v68_ = _dafny.Map({d_462_v65_: (self).f27})
                d_465_v68_ = (d_465_v68_).set((d_462_v65_) - (d_462_v65_), True)
                index54_ = default__.safeIndex(370, (d_464_v67_).length(0))
                (d_464_v67_)[index54_] = d_462_v65_
                d_466_v69_: _dafny.Map
                d_466_v69_ = _dafny.Map({d_462_v65_: d_462_v65_})
                d_467_v70_: _dafny.Map
                d_467_v70_ = _dafny.Map({d_462_v65_: d_466_v69_})
                d_468_v71_: _dafny.Seq
                d_468_v71_ = _dafny.SeqWithoutIsStrInference([d_467_v70_, d_467_v70_])
                index55_ = default__.safeIndex(370, (d_464_v67_).length(0))
                (d_464_v67_)[index55_] = default__.safeDivisionInt(d_462_v65_, len((d_468_v71_)[default__.safeIndex(-39, len(d_468_v71_))]))
                d_469_v72_: D0
                d_469_v72_ = D0_DC1((self).f27)
                d_470_v73_: bool
                d_471_v74_: _dafny.Array
                d_472_v75_: _dafny.Map
                d_473_v76_: _dafny.Seq
                out28_: bool
                out29_: _dafny.Array
                out30_: _dafny.Map
                out31_: _dafny.Seq
                out28_, out29_, out30_, out31_ = (self).m8(d_469_v72_, globalState)
                d_470_v73_ = out28_
                d_471_v74_ = out29_
                d_472_v75_ = out30_
                d_473_v76_ = out31_
            elif True:
                d_474_v77_: _dafny.Seq
                d_474_v77_ = _dafny.SeqWithoutIsStrInference([p0])
                d_475_v78_: D2
                d_475_v78_ = D2_DC5(d_474_v77_)
                d_476_v79_: _dafny.Array
                nw63_ = _dafny.Array(None, 14)
                nw63_[int(0)] = d_474_v77_
                nw63_[int(1)] = _dafny.SeqWithoutIsStrInference([p0])
                nw63_[int(2)] = default__.fm15(p1, p1, (self).f27, d_462_v65_, globalState)
                nw63_[int(3)] = d_474_v77_
                nw63_[int(4)] = d_474_v77_
                nw63_[int(5)] = d_474_v77_
                nw63_[int(6)] = d_474_v77_
                nw63_[int(7)] = d_474_v77_
                nw63_[int(8)] = (_dafny.SeqWithoutIsStrInference([p0, p0])).set(default__.safeIndex(d_462_v65_, len(_dafny.SeqWithoutIsStrInference([p0, p0]))), default__.fm16(p1, d_462_v65_, globalState))
                nw63_[int(9)] = d_474_v77_
                nw63_[int(10)] = d_474_v77_
                nw63_[int(11)] = (d_475_v78_).cf6
                nw63_[int(12)] = _dafny.SeqWithoutIsStrInference([p0])
                nw63_[int(13)] = d_474_v77_
                d_476_v79_ = nw63_
                d_477_v80_: _dafny.Seq
                d_477_v80_ = _dafny.SeqWithoutIsStrInference([d_476_v79_, d_476_v79_, d_476_v79_])
                d_478_v81_: _dafny.Array
                nw64_ = _dafny.Array(None, 22)
                nw64_[int(0)] = d_476_v79_
                nw64_[int(1)] = d_476_v79_
                nw64_[int(2)] = (d_477_v80_)[default__.safeIndex(d_462_v65_, len(d_477_v80_))]
                nw64_[int(3)] = d_476_v79_
                nw64_[int(4)] = (d_476_v79_ if (self).f27 else d_476_v79_)
                nw64_[int(5)] = d_476_v79_
                nw64_[int(6)] = d_476_v79_
                nw64_[int(7)] = d_476_v79_
                nw64_[int(8)] = d_476_v79_
                nw64_[int(9)] = d_476_v79_
                nw64_[int(10)] = d_476_v79_
                nw64_[int(11)] = d_476_v79_
                nw64_[int(12)] = d_476_v79_
                nw64_[int(13)] = d_476_v79_
                nw64_[int(14)] = d_476_v79_
                nw64_[int(15)] = d_476_v79_
                nw64_[int(16)] = d_476_v79_
                nw64_[int(17)] = d_476_v79_
                nw64_[int(18)] = d_476_v79_
                nw64_[int(19)] = (d_476_v79_ if (self).f27 else d_476_v79_)
                nw64_[int(20)] = d_476_v79_
                nw64_[int(21)] = d_476_v79_
                d_478_v81_ = nw64_
                index56_ = default__.safeIndex(905, (d_478_v81_).length(0))
                (d_478_v81_)[index56_] = d_476_v79_
                d_479_v82_: _dafny.Seq
                d_479_v82_ = _dafny.SeqWithoutIsStrInference([d_462_v65_])
                d_480_v83_: _dafny.Map
                d_480_v83_ = _dafny.Map({(d_479_v82_) + (d_479_v82_): (0) - ((d_462_v65_) - (d_462_v65_))})
                d_481_v84_: _dafny.Array
                nw65_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 10)
                d_481_v84_ = nw65_
                d_482_v85_: _dafny.Seq
                d_482_v85_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "txya"))
                index57_ = default__.safeIndex(580, (d_481_v84_).length(0))
                (d_481_v84_)[index57_] = d_482_v85_
                d_483_v87_: D3
                def iife64_():
                    coll42_ = _dafny.Map()
                    compr_42_: _dafny.Seq
                    for compr_42_ in (_dafny.Map({d_479_v82_: d_462_v65_})).keys.Elements:
                        d_484_v86_: _dafny.Seq = compr_42_
                        if (d_484_v86_) in (_dafny.Map({d_479_v82_: d_462_v65_})):
                            coll42_[d_484_v86_] = d_462_v65_
                    return _dafny.Map(coll42_)
                d_483_v87_ = D3_DC8(iife64_()
)
                d_485_v89_: _dafny.Set
                d_485_v89_ = _dafny.Set({d_479_v82_})
                pat_let_tv6_ = d_485_v89_
                pat_let_tv7_ = d_485_v89_
                pat_let_tv8_ = d_462_v65_
                index58_ = default__.safeIndex(905, (d_478_v81_).length(0))
                index59_ = default__.safeIndex(580, (d_481_v84_).length(0))
                def iife65_(_pat_let11_0):
                    def iife66_(d_486_dt__update__tmp_h1_):
                        def iife68_():
                            coll43_ = _dafny.Map()
                            compr_43_: _dafny.Seq
                            for compr_43_ in (pat_let_tv6_).Elements:
                                d_487_v88_: _dafny.Seq = compr_43_
                                if (d_487_v88_) in (pat_let_tv7_):
                                    coll43_[d_487_v88_] = pat_let_tv8_
                            return _dafny.Map(coll43_)
                        def iife67_(_pat_let12_0):
                            def iife69_(d_488_dt__update_hcf13_h0_):
                                return D3_DC8(d_488_dt__update_hcf13_h0_)
                            return iife69_(_pat_let12_0)
                        return iife67_(iife68_()
                        )
                    return iife66_(_pat_let11_0)
                rhs65_ = d_476_v79_
                rhs66_ = (iife65_(d_483_v87_)).cf13
                rhs67_ = p1
                rhs68_ = (self).fm3(globalState)
                lhs50_ = d_478_v81_
                lhs51_ = default__.safeIndex(905, (d_478_v81_).length(0))
                lhs52_ = globalState
                lhs53_ = d_481_v84_
                lhs54_ = default__.safeIndex(580, (d_481_v84_).length(0))
                lhs50_[lhs51_] = rhs65_
                d_480_v83_ = rhs66_
                lhs52_.f21 = rhs67_
                lhs53_[lhs54_] = rhs68_
                d_489_v90_: _dafny.Array
                nw66_ = _dafny.Array(int(0), 15)
                d_489_v90_ = nw66_
                index60_ = default__.safeIndex(258, (d_489_v90_).length(0))
                (d_489_v90_)[index60_] = d_462_v65_
                index61_ = default__.safeIndex(258, (d_489_v90_).length(0))
                (d_489_v90_)[index61_] = ((d_462_v65_) * (480) if not(p1) else (d_462_v65_) - (d_462_v65_))
                d_490_v91_: D0
                d_490_v91_ = D0_DC1(p1)
                index62_ = default__.safeIndex(431, (d_463_v66_).length(0))
                (d_463_v66_)[index62_] = (self).fm4(d_490_v91_, globalState)
                index63_ = default__.safeIndex(431, (d_463_v66_).length(0))
                (d_463_v66_)[index63_] = (len(d_479_v82_)) <= (default__.safeDivisionInt(611, -496))
                d_489_v90_ = d_489_v90_
                index64_ = default__.safeIndex(580, (d_481_v84_).length(0))
                def iife70_():
                    coll44_ = _dafny.Map()
                    compr_44_: str
                    for compr_44_ in ((d_481_v84_)[default__.safeIndex(580, (d_481_v84_).length(0))]).Elements:
                        d_491_v92_: str = compr_44_
                        if (d_491_v92_) in ((d_481_v84_)[default__.safeIndex(580, (d_481_v84_).length(0))]):
                            coll44_[d_491_v92_] = _dafny.Map({d_462_v65_: _dafny.CodePoint('m')})
                    return _dafny.Map(coll44_)
                (d_481_v84_)[index64_] = ((self).fm2(d_462_v65_, iife70_()
                , globalState)) + (((d_481_v84_)[default__.safeIndex(580, (d_481_v84_).length(0))]) + ((self).fm3(globalState)))
            if not(p1):
                (globalState).f21 = not((self).f27)
                d_492_v93_: _dafny.Seq
                d_492_v93_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eb"))
                d_493_v94_: _dafny.Seq
                d_493_v94_ = _dafny.SeqWithoutIsStrInference([p1, default__.fm0(d_492_v93_, (self).f27, (self).f27, _dafny.Set({d_462_v65_}), globalState)])
                (globalState).f2 = default__.fm1(False, (d_493_v94_) + (d_493_v94_), (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_494_i9_ in range(default__.abs(677))])) < (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_495_i10_ in range(default__.abs(200))])), globalState)
                d_496_v95_: C0
                nw67_ = C0()
                nw67_.ctor__()
                d_496_v95_ = nw67_
                d_497_v96_: D0
                d_497_v96_ = D0_DC1(p1)
                d_498_v97_: bool
                d_499_v98_: _dafny.Array
                d_500_v99_: _dafny.Map
                d_501_v100_: _dafny.Seq
                out32_: bool
                out33_: _dafny.Array
                out34_: _dafny.Map
                out35_: _dafny.Seq
                out32_, out33_, out34_, out35_ = (self).m8(d_497_v96_, globalState)
                d_498_v97_ = out32_
                d_499_v98_ = out33_
                d_500_v99_ = out34_
                d_501_v100_ = out35_
                d_492_v93_ = ((d_492_v93_) + (d_492_v93_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vfppletdw")))
            elif True:
                (globalState).f2 = default__.safeDivisionInt(default__.safeModuloInt(d_462_v65_, d_462_v65_), d_462_v65_)
                (globalState).f6 = 106
                (globalState).f12 = p1
                d_502_v101_: D0
                d_502_v101_ = D0_DC1((self).f27)
                d_503_v102_: bool
                d_504_v103_: _dafny.Seq
                out36_: bool
                out37_: _dafny.Seq
                out36_, out37_ = default__.m0(p0, d_502_v101_, globalState)
                d_503_v102_ = out36_
                d_504_v103_ = out37_
                (globalState).f14 = (self).f27
        d_505_v104_: int
        d_505_v104_ = 241
        if (d_505_v104_) <= ((d_505_v104_ if (self).f27 else d_505_v104_)):
            d_506_v105_: _dafny.Map
            d_506_v105_ = _dafny.Map({p1: d_505_v104_})
            (globalState).f6 = len((_dafny.Map({p1: d_505_v104_})) | ((d_506_v105_) | (d_506_v105_)))
            (globalState).f21 = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_507_i11_ in range(default__.abs(471))])) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ia")))
            (globalState).f6 = d_505_v104_
            d_508_v106_: D0
            d_508_v106_ = D0_DC1(p1)
            pat_let_tv9_ = d_508_v106_
            d_509_v107_: bool
            d_510_v108_: _dafny.Seq
            out38_: bool
            out39_: _dafny.Seq
            def iife71_(_pat_let13_0):
                def iife72_(d_511_dt__update__tmp_h2_):
                    def iife73_(_pat_let14_0):
                        def iife74_(d_512_dt__update_hcf4_h0_):
                            return D0_DC3(d_512_dt__update_hcf4_h0_)
                        return iife74_(_pat_let14_0)
                    return iife73_(pat_let_tv9_)
                return iife72_(_pat_let13_0)
            out38_, out39_ = default__.m0((p0 if p1 else iife71_(D0_DC3(d_508_v106_))), D0_DC1((self).f27), globalState)
            d_509_v107_ = out38_
            d_510_v108_ = out39_
            d_513_v109_: D0
            d_513_v109_ = D0_DC1(p1)
            d_514_v110_: bool
            d_515_v111_: _dafny.Array
            d_516_v112_: _dafny.Map
            d_517_v113_: _dafny.Seq
            out40_: bool
            out41_: _dafny.Array
            out42_: _dafny.Map
            out43_: _dafny.Seq
            out40_, out41_, out42_, out43_ = (self).m8(d_513_v109_, globalState)
            d_514_v110_ = out40_
            d_515_v111_ = out41_
            d_516_v112_ = out42_
            d_517_v113_ = out43_
        elif True:
            (globalState).f5 = ((self).f27 if (self).f27 else not((self).f27))
            d_518_v114_: _dafny.Seq
            d_518_v114_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fnu"))
            d_519_v115_: _dafny.Seq
            d_519_v115_ = _dafny.SeqWithoutIsStrInference([p1, p1])
            d_520_v116_: str
            d_520_v116_ = _dafny.CodePoint('b')
            d_518_v114_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_521_i12_ in range(default__.abs(329))])) + (((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kxslex"))).set(default__.safeIndex((0) - (default__.fm1((self).f27, d_519_v115_, False, globalState)), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kxslex")))), d_520_v116_)) + (d_518_v114_))
            d_522_v117_: C0
            nw68_ = C0()
            nw68_.ctor__()
            d_522_v117_ = nw68_
            d_523_v118_: _dafny.Map
            d_523_v118_ = _dafny.Map({d_505_v104_: d_522_v117_})
            d_524_v119_: _dafny.Map
            d_524_v119_ = d_523_v118_
            (globalState).f5 = (229) < (len((d_524_v119_)))
            d_525_v120_: D0
            d_525_v120_ = D0_DC1(p1)
            d_526_v121_: bool
            d_527_v122_: _dafny.Seq
            out44_: bool
            out45_: _dafny.Seq
            out44_, out45_ = default__.m0((p0 if p1 else p0), (d_525_v120_ if not((self).f27) else D0_DC1((self).f27)), globalState)
            d_526_v121_ = out44_
            d_527_v122_ = out45_
            d_528_v123_: _dafny.Map
            d_528_v123_ = _dafny.Map({d_505_v104_: d_520_v116_})
            d_529_v124_: _dafny.Map
            d_529_v124_ = _dafny.Map({p1: (self).f27})
            d_530_v125_: _dafny.Array
            nw69_ = _dafny.Array(None, 19)
            nw69_[int(0)] = d_520_v116_
            nw69_[int(1)] = d_520_v116_
            nw69_[int(2)] = default__.fm14(d_505_v104_, globalState)
            nw69_[int(3)] = d_520_v116_
            nw69_[int(4)] = d_520_v116_
            nw69_[int(5)] = d_520_v116_
            nw69_[int(6)] = d_520_v116_
            nw69_[int(7)] = d_520_v116_
            nw69_[int(8)] = _dafny.CodePoint('j')
            nw69_[int(9)] = d_520_v116_
            nw69_[int(10)] = d_520_v116_
            nw69_[int(11)] = ((d_528_v123_)[d_505_v104_] if (d_505_v104_) in (d_528_v123_) else d_520_v116_)
            nw69_[int(12)] = d_520_v116_
            nw69_[int(13)] = d_520_v116_
            nw69_[int(14)] = _dafny.CodePoint('c')
            nw69_[int(15)] = _dafny.CodePoint('c')
            nw69_[int(16)] = (d_520_v116_ if ((d_529_v124_)[not(p1)] if (not(p1)) in (d_529_v124_) else (self).f27) else default__.fm14(d_505_v104_, globalState))
            nw69_[int(17)] = _dafny.CodePoint('w')
            nw69_[int(18)] = d_520_v116_
            d_530_v125_ = nw69_
            index65_ = default__.safeIndex(581, (d_530_v125_).length(0))
            (d_530_v125_)[index65_] = d_520_v116_
            index66_ = default__.safeIndex(581, (d_530_v125_).length(0))
            rhs69_ = p1
            rhs70_ = d_520_v116_
            lhs55_ = d_530_v125_
            lhs56_ = default__.safeIndex(581, (d_530_v125_).length(0))
            r0 = rhs69_
            lhs55_[lhs56_] = rhs70_
        d_531_v126_: _dafny.Array
        nw70_ = _dafny.Array(_dafny.Array(None, 0), 2)
        d_531_v126_ = nw70_
        d_532_v127_: _dafny.Array
        nw71_ = _dafny.Array(False, 1)
        d_532_v127_ = nw71_
        index67_ = default__.safeIndex(660, (d_531_v126_).length(0))
        (d_531_v126_)[index67_] = d_532_v127_
        index68_ = default__.safeIndex(660, (d_531_v126_).length(0))
        nw72_ = _dafny.Array(False, 23)
        (d_531_v126_)[index68_] = nw72_
        d_533_v128_: _dafny.Seq
        d_533_v128_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ruxhr"))
        d_533_v128_ = (d_533_v128_ if ((self).f27) or (p1) else d_533_v128_)
        d_534_v129_: D0
        d_534_v129_ = D0_DC2(d_505_v104_, d_505_v104_)
        pat_let_tv10_ = d_505_v104_
        d_535_v130_: _dafny.Array
        nw73_ = _dafny.Array(None, 21)
        nw73_[int(0)] = d_534_v129_
        nw73_[int(1)] = d_534_v129_
        nw73_[int(2)] = d_534_v129_
        nw73_[int(3)] = (d_534_v129_ if (self).f27 else d_534_v129_)
        nw73_[int(4)] = d_534_v129_
        nw73_[int(5)] = d_534_v129_
        nw73_[int(6)] = d_534_v129_
        nw73_[int(7)] = D0_DC2(d_505_v104_, d_505_v104_)
        nw73_[int(8)] = d_534_v129_
        nw73_[int(9)] = d_534_v129_
        nw73_[int(10)] = d_534_v129_
        nw73_[int(11)] = d_534_v129_
        def iife75_(_pat_let15_0):
            def iife76_(d_536_dt__update__tmp_h3_):
                def iife77_(_pat_let16_0):
                    def iife78_(d_537_dt__update_hcf3_h0_):
                        return D0_DC2((d_536_dt__update__tmp_h3_).cf2, d_537_dt__update_hcf3_h0_)
                    return iife78_(_pat_let16_0)
                return iife77_(pat_let_tv10_)
            return iife76_(_pat_let15_0)
        nw73_[int(12)] = iife75_(d_534_v129_)
        nw73_[int(13)] = d_534_v129_
        nw73_[int(14)] = d_534_v129_
        nw73_[int(15)] = (self).fm5(globalState)
        nw73_[int(16)] = d_534_v129_
        nw73_[int(17)] = d_534_v129_
        nw73_[int(18)] = D0_DC2(d_505_v104_, d_505_v104_)
        nw73_[int(19)] = d_534_v129_
        nw73_[int(20)] = d_534_v129_
        d_535_v130_ = nw73_
        index69_ = default__.safeIndex(653, (d_535_v130_).length(0))
        (d_535_v130_)[index69_] = (self).fm5(globalState)
        index70_ = default__.safeIndex(653, (d_535_v130_).length(0))
        (d_535_v130_)[index70_] = d_534_v129_
        d_538_v131_: D0
        d_538_v131_ = D0_DC1((self).f27)
        d_539_i13_: int
        d_539_i13_ = 0
        with _dafny.label("2"):
            def lambda11_(source12_):
                if source12_.is_DC1:
                    d_549___mcc_h5_ = source12_.cf1
                    d_550_cf1_ = d_549___mcc_h5_
                    return False
                elif source12_.is_DC2:
                    d_551___mcc_h6_ = source12_.cf2
                    d_552___mcc_h7_ = source12_.cf3
                    d_553_cf3_ = d_552___mcc_h7_
                    d_554_cf2_ = d_551___mcc_h6_
                    return (self).f27
                elif source12_.is_DC0:
                    d_555___mcc_h8_ = source12_.cf0
                    d_556_cf0_ = d_555___mcc_h8_
                    return (self).f27
                elif True:
                    d_557___mcc_h9_ = source12_.cf4
                    d_558_cf4_ = d_557___mcc_h9_
                    return (self).f27

            while lambda11_(d_538_v131_):
                with _dafny.c_label("2"):
                    if (d_539_i13_) >= (100):
                        raise _dafny.Break("2")
                    d_539_i13_ = (d_539_i13_) + (1)
                    d_540_v132_: str
                    d_540_v132_ = _dafny.CodePoint('o')
                    d_540_v132_ = d_540_v132_
                    d_541_v133_: _dafny.Seq
                    d_541_v133_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wjentgq"))])
                    d_542_v134_: _dafny.Array
                    nw74_ = _dafny.Array(None, 16)
                    nw74_[int(0)] = (d_533_v128_) + ((d_541_v133_)[default__.safeIndex(d_505_v104_, len(d_541_v133_))])
                    nw74_[int(1)] = d_533_v128_
                    nw74_[int(2)] = (d_533_v128_) + (d_533_v128_)
                    nw74_[int(3)] = (d_533_v128_) + (d_533_v128_)
                    nw74_[int(4)] = d_533_v128_
                    nw74_[int(5)] = ((self).fm3(globalState)) + (_dafny.SeqWithoutIsStrInference([d_540_v132_ for d_543_i14_ in range(default__.abs(146))]))
                    nw74_[int(6)] = (d_533_v128_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vti")))
                    nw74_[int(7)] = (_dafny.SeqWithoutIsStrInference([d_540_v132_ for d_544_i15_ in range(default__.abs(131))])) + (d_533_v128_)
                    nw74_[int(8)] = (d_533_v128_) + (d_533_v128_)
                    nw74_[int(9)] = d_533_v128_
                    nw74_[int(10)] = (d_533_v128_) + (d_533_v128_)
                    nw74_[int(11)] = (self).fm3(globalState)
                    nw74_[int(12)] = d_533_v128_
                    nw74_[int(13)] = (d_533_v128_) + (d_533_v128_)
                    nw74_[int(14)] = d_533_v128_
                    nw74_[int(15)] = d_533_v128_
                    d_542_v134_ = nw74_
                    index71_ = default__.safeIndex(142, (d_542_v134_).length(0))
                    (d_542_v134_)[index71_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mw"))) + (d_533_v128_)
                    index72_ = default__.safeIndex(142, (d_542_v134_).length(0))
                    (d_542_v134_)[index72_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jp"))
                    (globalState).f21 = not((d_505_v104_) != ((d_505_v104_) + (-440)))
                    d_545_v135_: _dafny.Array
                    nw75_ = _dafny.Array(_dafny.Array(None, 0), 4)
                    d_545_v135_ = nw75_
                    d_546_v136_: _dafny.Map
                    d_546_v136_ = _dafny.Map({d_545_v135_: d_505_v104_})
                    d_547_v137_: _dafny.Map
                    d_547_v137_ = _dafny.Map({(self).f27: d_545_v135_})
                    d_548_v138_: D5
                    d_548_v138_ = D5_DC11(((d_547_v137_)[(self).f27] if ((self).f27) in (d_547_v137_) else d_545_v135_))
                    d_546_v136_ = (d_546_v136_).set((d_548_v138_).cf18, d_505_v104_)
                    pass
            pass
        r0 = p1
        r1 = p1
        return r0, r1

    def m1(self, globalState):
        r0: int = int(0)
        d_559_v0_: _dafny.Seq
        d_559_v0_ = _dafny.SeqWithoutIsStrInference([(self).f27])
        d_560_v1_: int
        d_560_v1_ = 884
        d_561_v2_: _dafny.Seq
        d_561_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gos"))
        d_562_v3_: str
        d_562_v3_ = _dafny.CodePoint('u')
        d_563_v4_: D0
        d_563_v4_ = D0_DC1((self).f27)
        d_564_v5_: _dafny.Array
        nw76_ = _dafny.Array(None, 7)
        nw76_[int(0)] = d_559_v0_
        nw76_[int(1)] = d_559_v0_
        nw76_[int(2)] = _dafny.SeqWithoutIsStrInference([(self).f27, (self).f27, (self).f27])
        nw76_[int(3)] = default__.fm17(d_560_v1_, (d_561_v2_).set(default__.safeIndex(d_560_v1_, len(d_561_v2_)), d_562_v3_), d_560_v1_, globalState)
        nw76_[int(4)] = (_dafny.SeqWithoutIsStrInference([(self).fm4(d_563_v4_, globalState), (self).f27])) + (d_559_v0_)
        nw76_[int(5)] = (d_559_v0_) + (default__.fm17(len(_dafny.SeqWithoutIsStrInference([(d_561_v2_)[default__.safeIndex(len(_dafny.Map({d_560_v1_: (self).f27})), len(d_561_v2_))] for d_565_i0_ in range(default__.abs(861))])), d_561_v2_, d_560_v1_, globalState))
        nw76_[int(6)] = default__.fm17(d_560_v1_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qnedubdqd"))).set(default__.safeIndex(d_560_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qnedubdqd")))), d_562_v3_), d_560_v1_, globalState)
        d_564_v5_ = nw76_
        index73_ = default__.safeIndex(92, (d_564_v5_).length(0))
        (d_564_v5_)[index73_] = d_559_v0_
        d_566_v6_: _dafny.Seq
        d_566_v6_ = _dafny.SeqWithoutIsStrInference([d_560_v1_])
        d_567_v7_: _dafny.Seq
        d_567_v7_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_560_v1_, d_560_v1_, len(_dafny.Set({not((self).f27), default__.fm0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ngyf")), (self).f27, (self).f27, _dafny.Set({len(d_566_v6_), d_560_v1_}), globalState)}))})])
        index74_ = default__.safeIndex(92, (d_564_v5_).length(0))
        (d_564_v5_)[index74_] = (d_559_v0_).set(default__.safeIndex((0) - ((len(d_567_v7_)) - (d_560_v1_)), len(d_559_v0_)), (self).f27)
        rhs71_ = d_562_v3_
        rhs72_ = ((_dafny.SeqWithoutIsStrInference([d_562_v3_ for d_568_i1_ in range(default__.abs(-158))]) if (self).f27 else _dafny.SeqWithoutIsStrInference([d_562_v3_ for d_569_i2_ in range(default__.abs(-1))]))) < (_dafny.SeqWithoutIsStrInference([d_562_v3_ for d_570_i3_ in range(default__.abs(-456))]))
        lhs57_ = globalState
        d_562_v3_ = rhs71_
        lhs57_.f21 = rhs72_
        d_559_v0_ = d_559_v0_
        d_571_v8_: _dafny.Array
        nw77_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 15)
        d_571_v8_ = nw77_
        index75_ = default__.safeIndex(359, (d_571_v8_).length(0))
        (d_571_v8_)[index75_] = d_561_v2_
        index76_ = default__.safeIndex(359, (d_571_v8_).length(0))
        (d_571_v8_)[index76_] = (d_561_v2_) + (d_561_v2_)
        if (self).f27:
            d_572_v9_: D3
            d_572_v9_ = D3_DC9((self).f27, d_560_v1_, d_560_v1_)
            d_573_v10_: _dafny.Seq
            d_573_v10_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_562_v3_])])
            d_574_v11_: D2
            d_574_v11_ = D2_DC6(d_560_v1_, d_566_v6_, (self).fm4(d_563_v4_, globalState), d_566_v6_, not((self).f27))
            d_575_v12_: _dafny.Map
            d_575_v12_ = _dafny.Map({d_560_v1_: (d_574_v11_).cf7})
            d_576_v13_: _dafny.Seq
            d_576_v13_ = _dafny.SeqWithoutIsStrInference([d_575_v12_])
            d_577_v14_: _dafny.MultiSet
            d_577_v14_ = _dafny.MultiSet([d_560_v1_])
            d_578_v15_: _dafny.Set
            d_578_v15_ = _dafny.Set({(self).f27, (self).f27, True})
            d_579_v16_: D5
            d_579_v16_ = D5_DC12(((d_564_v5_)[default__.safeIndex(92, (d_564_v5_).length(0))])[default__.safeIndex(d_560_v1_, len((d_564_v5_)[default__.safeIndex(92, (d_564_v5_).length(0))]))], -444)
            d_580_v17_: _dafny.MultiSet
            d_580_v17_ = _dafny.MultiSet([True, (d_579_v16_).cf19, True])
            d_581_v18_: _dafny.Array
            nw78_ = _dafny.Array(None, 23)
            nw78_[int(0)] = (self).f27
            nw78_[int(1)] = (self).f27
            nw78_[int(2)] = (d_572_v9_).cf14
            nw78_[int(3)] = (self).f27
            nw78_[int(4)] = (self).f27
            nw78_[int(5)] = (self).f27
            nw78_[int(6)] = not((self).f27)
            nw78_[int(7)] = (d_562_v3_) in ((d_573_v10_)[default__.safeIndex(d_560_v1_, len(d_573_v10_))])
            nw78_[int(8)] = (self).f27
            nw78_[int(9)] = (self).f27
            nw78_[int(10)] = ((d_571_v8_)[default__.safeIndex(359, (d_571_v8_).length(0))]) == ((d_571_v8_)[default__.safeIndex(359, (d_571_v8_).length(0))])
            nw78_[int(11)] = not((self).f27)
            nw78_[int(12)] = (345) in ((d_576_v13_)[default__.safeIndex((0) - (((d_577_v14_)[d_560_v1_] if (d_560_v1_) in (d_577_v14_) else len(d_578_v15_))), len(d_576_v13_))])
            nw78_[int(13)] = (self).f27
            nw78_[int(14)] = (self).f27
            nw78_[int(15)] = (self).f27
            nw78_[int(16)] = ((self).f27) in (d_580_v17_)
            nw78_[int(17)] = (self).f27
            nw78_[int(18)] = (self).f27
            nw78_[int(19)] = (self).f27
            nw78_[int(20)] = (self).f27
            nw78_[int(21)] = (self).f27
            nw78_[int(22)] = not (False) or ((self).f27)
            d_581_v18_ = nw78_
            d_581_v18_ = (d_581_v18_ if (self).f27 else d_581_v18_)
            (globalState).f2 = len(d_559_v0_)
            d_582_v19_: _dafny.Array
            def lambda12_(d_583_v1_):
                def lambda13_(d_584_i4_):
                    return (d_584_i4_) - (d_583_v1_)

                return lambda13_

            init6_ = lambda12_(d_560_v1_)
            nw79_ = _dafny.Array(None, 9)
            for i0_6_ in range(nw79_.length(0)):
                nw79_[i0_6_] = init6_(i0_6_)
            d_582_v19_ = nw79_
            index77_ = default__.safeIndex(855, (d_581_v18_).length(0))
            (d_581_v18_)[index77_] = (self).f27
            d_585_v20_: _dafny.Set
            d_585_v20_ = _dafny.Set({d_560_v1_, len(d_559_v0_)})
            index78_ = default__.safeIndex(359, (d_571_v8_).length(0))
            index79_ = default__.safeIndex(855, (d_581_v18_).length(0))
            rhs73_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
            rhs74_ = d_582_v19_
            rhs75_ = default__.safeDivisionInt(d_560_v1_, -804)
            rhs76_ = not (((d_563_v4_).cf1 if not((self).f27) else (self).f27)) or ((self).f27)
            rhs77_ = (default__.fm0((d_571_v8_)[default__.safeIndex(359, (d_571_v8_).length(0))], not((self).f27), (self).f27, d_585_v20_, globalState)) and ((self).f27)
            lhs58_ = d_571_v8_
            lhs59_ = default__.safeIndex(359, (d_571_v8_).length(0))
            lhs60_ = globalState
            lhs61_ = d_581_v18_
            lhs62_ = default__.safeIndex(855, (d_581_v18_).length(0))
            lhs63_ = globalState
            lhs58_[lhs59_] = rhs73_
            d_582_v19_ = rhs74_
            lhs60_.f2 = rhs75_
            lhs61_[lhs62_] = rhs76_
            lhs63_.f14 = rhs77_
            d_586_v21_: _dafny.Map
            d_586_v21_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_562_v3_ for d_587_i5_ in range(default__.abs(-85))])): not((d_581_v18_)[default__.safeIndex(855, (d_581_v18_).length(0))])})
            d_588_v22_: _dafny.Array
            nw80_ = _dafny.Array(None, 1)
            nw80_[int(0)] = d_586_v21_
            d_588_v22_ = nw80_
            index80_ = default__.safeIndex(6, (d_588_v22_).length(0))
            (d_588_v22_)[index80_] = d_586_v21_
            index81_ = default__.safeIndex(6, (d_588_v22_).length(0))
            (d_588_v22_)[index81_] = d_586_v21_
            d_574_v11_ = d_574_v11_
        elif True:
            d_589_v23_: _dafny.Array
            nw81_ = _dafny.Array(D3.default()(), 19)
            d_589_v23_ = nw81_
            d_590_v24_: D3
            d_590_v24_ = D3_DC9((self).f27, d_560_v1_, d_560_v1_)
            index82_ = default__.safeIndex(151, (d_589_v23_).length(0))
            (d_589_v23_)[index82_] = d_590_v24_
            index83_ = default__.safeIndex(151, (d_589_v23_).length(0))
            (d_589_v23_)[index83_] = D3_DC9(((0) - (d_560_v1_)) < (len(_dafny.SeqWithoutIsStrInference([len((d_564_v5_)[default__.safeIndex(92, (d_564_v5_).length(0))]) for d_591_i6_ in range(default__.abs(-307))]))), 452, d_560_v1_)
            (globalState).f21 = (self).f27
            d_592_v25_: _dafny.Map
            d_592_v25_ = _dafny.Map({d_560_v1_: d_560_v1_})
            d_593_v26_: _dafny.Map
            d_593_v26_ = _dafny.Map({-470: ((d_592_v25_)[d_560_v1_] if (d_560_v1_) in (d_592_v25_) else d_560_v1_)})
            d_594_v27_: _dafny.Map
            d_594_v27_ = _dafny.Map({((d_571_v8_)[default__.safeIndex(359, (d_571_v8_).length(0))])[default__.safeIndex(len(d_593_v26_), len((d_571_v8_)[default__.safeIndex(359, (d_571_v8_).length(0))]))]: d_560_v1_})
            d_595_v28_: D6
            d_595_v28_ = D6_DC13(d_559_v0_)
            d_596_v29_: C0
            nw82_ = C0()
            nw82_.ctor__()
            d_596_v29_ = nw82_
            d_597_v30_: _dafny.Map
            d_597_v30_ = _dafny.Map({d_560_v1_: d_596_v29_})
            d_598_v31_: _dafny.Seq
            d_598_v31_ = _dafny.SeqWithoutIsStrInference([d_596_v29_, ((d_597_v30_)[d_560_v1_] if (d_560_v1_) in (d_597_v30_) else d_596_v29_)])
            d_599_v32_: _dafny.MultiSet
            d_599_v32_ = _dafny.MultiSet([(self).f27])
            d_600_v33_: _dafny.Array
            nw83_ = _dafny.Array(None, 24)
            nw83_[int(0)] = d_560_v1_
            nw83_[int(1)] = len(d_561_v2_)
            nw83_[int(2)] = d_560_v1_
            nw83_[int(3)] = d_560_v1_
            nw83_[int(4)] = (_dafny.MultiSet(d_598_v31_)).cardinality
            nw83_[int(5)] = (0) - (d_560_v1_)
            nw83_[int(6)] = d_560_v1_
            nw83_[int(7)] = d_560_v1_
            nw83_[int(8)] = d_560_v1_
            nw83_[int(9)] = 544
            nw83_[int(10)] = d_560_v1_
            nw83_[int(11)] = len(d_561_v2_)
            nw83_[int(12)] = d_560_v1_
            nw83_[int(13)] = 427
            nw83_[int(14)] = d_560_v1_
            nw83_[int(15)] = d_560_v1_
            nw83_[int(16)] = d_560_v1_
            nw83_[int(17)] = default__.fm1((self).f27, (d_564_v5_)[default__.safeIndex(92, (d_564_v5_).length(0))], (self).f27, globalState)
            nw83_[int(18)] = d_560_v1_
            nw83_[int(19)] = -241
            nw83_[int(20)] = d_560_v1_
            nw83_[int(21)] = (d_599_v32_).cardinality
            nw83_[int(22)] = d_560_v1_
            nw83_[int(23)] = len(d_593_v26_)
            d_600_v33_ = nw83_
            d_601_v34_: _dafny.Map
            d_601_v34_ = _dafny.Map({d_599_v32_: d_559_v0_})
            d_602_v35_: _dafny.Map
            d_602_v35_ = _dafny.Map({d_600_v33_: ((d_601_v34_)[d_599_v32_] if (d_599_v32_) in (d_601_v34_) else (d_564_v5_)[default__.safeIndex(92, (d_564_v5_).length(0))])})
            d_603_v36_: D0
            d_603_v36_ = D0_DC0(len(d_602_v35_))
            d_604_v37_: _dafny.Map
            d_604_v37_ = _dafny.Map({d_560_v1_: d_562_v3_})
            d_605_v38_: _dafny.Map
            d_605_v38_ = _dafny.Map({(self).f27: d_560_v1_})
            d_606_v39_: _dafny.Seq
            d_606_v39_ = _dafny.SeqWithoutIsStrInference([d_605_v38_])
            d_607_v40_: _dafny.Map
            d_607_v40_ = _dafny.Map({len(d_593_v26_): (self).f27})
            d_608_v41_: _dafny.Array
            nw84_ = _dafny.Array(None, 23)
            nw84_[int(0)] = ((0) - (len(d_594_v27_)) if (self).f27 else d_560_v1_)
            nw84_[int(1)] = default__.safeDivisionInt(len((d_595_v28_).cf21), d_560_v1_)
            nw84_[int(2)] = d_560_v1_
            nw84_[int(3)] = default__.fm1(True, d_559_v0_, (self).f27, globalState)
            nw84_[int(4)] = d_560_v1_
            nw84_[int(5)] = d_560_v1_
            nw84_[int(6)] = len(d_593_v26_)
            nw84_[int(7)] = default__.fm1((self).f27, (d_564_v5_)[default__.safeIndex(92, (d_564_v5_).length(0))], (self).f27, globalState)
            nw84_[int(8)] = d_560_v1_
            nw84_[int(9)] = (0) - ((d_560_v1_) - (d_560_v1_))
            nw84_[int(10)] = d_560_v1_
            nw84_[int(11)] = (d_603_v36_).cf0
            nw84_[int(12)] = 525
            nw84_[int(13)] = d_560_v1_
            nw84_[int(14)] = ((d_593_v26_)[len(d_604_v37_)] if (len(d_604_v37_)) in (d_593_v26_) else d_560_v1_)
            nw84_[int(15)] = d_560_v1_
            nw84_[int(16)] = (d_560_v1_) - (d_560_v1_)
            nw84_[int(17)] = d_560_v1_
            nw84_[int(18)] = d_560_v1_
            nw84_[int(19)] = len((d_606_v39_) + (d_606_v39_))
            nw84_[int(20)] = (default__.fm13(globalState)).cardinality
            nw84_[int(21)] = len((D7_DC15(d_607_v40_)).cf24)
            nw84_[int(22)] = d_560_v1_
            d_608_v41_ = nw84_
            index84_ = default__.safeIndex(801, (d_600_v33_).length(0))
            (d_600_v33_)[index84_] = d_560_v1_
            index85_ = default__.safeIndex(801, (d_600_v33_).length(0))
            (d_600_v33_)[index85_] = d_560_v1_
            if (self).f27:
                d_609_v42_: _dafny.Map
                d_609_v42_ = _dafny.Map({(self).f27: (self).f27})
                d_610_v43_: _dafny.MultiSet
                d_610_v43_ = _dafny.MultiSet([((d_607_v40_).set(d_560_v1_, (self).f27)).set(638, (self).f27)])
                d_609_v42_ = (d_609_v42_).set(True, (d_610_v43_) != (d_610_v43_))
                (globalState).f2 = len(d_561_v2_)
                (globalState).f6 = (d_560_v1_) - ((d_600_v33_)[default__.safeIndex(801, (d_600_v33_).length(0))])
                d_611_v44_: _dafny.Array
                def lambda14_(d_612_v42_):
                    def lambda15_(d_613_i7_):
                        return D6_DC14((self).f27, len(d_612_v42_))

                    return lambda15_

                init7_ = lambda14_(d_609_v42_)
                nw85_ = _dafny.Array(None, 4)
                for i0_7_ in range(nw85_.length(0)):
                    nw85_[i0_7_] = init7_(i0_7_)
                d_611_v44_ = nw85_
                d_614_v45_: D6
                d_614_v45_ = D6_DC14(not((self).f27), len(d_561_v2_))
                index86_ = default__.safeIndex(993, (d_611_v44_).length(0))
                (d_611_v44_)[index86_] = d_614_v45_
                index87_ = default__.safeIndex(993, (d_611_v44_).length(0))
                (d_611_v44_)[index87_] = d_614_v45_
                (globalState).f21 = (self).f27
            elif True:
                (globalState).f6 = (d_566_v6_)[default__.safeIndex(824, len(d_566_v6_))]
                d_615_v46_: C0
                nw86_ = C0()
                nw86_.ctor__()
                d_615_v46_ = nw86_
                d_608_v41_ = d_608_v41_
                (globalState).f6 = (default__.safeDivisionInt((d_600_v33_)[default__.safeIndex(801, (d_600_v33_).length(0))], 356)) * ((0) - ((d_600_v33_)[default__.safeIndex(801, (d_600_v33_).length(0))]))
                d_616_v47_: C0
                nw87_ = C0()
                nw87_.ctor__()
                d_616_v47_ = nw87_
            index88_ = default__.safeIndex(801, (d_600_v33_).length(0))
            (d_600_v33_)[index88_] = (d_566_v6_)[default__.safeIndex(-171, len(d_566_v6_))]
        r0 = d_560_v1_
        r0 = d_560_v1_
        return r0

    def m2(self, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: bool = False
        d_617_v0_: _dafny.Seq
        d_617_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lachopxht"))
        d_618_i0_: int
        d_618_i0_ = 0
        with _dafny.label("3"):
            while not ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_635_i1_ in range(default__.abs(-316))])) != (d_617_v0_)) or (True):
                with _dafny.c_label("3"):
                    if (d_618_i0_) >= (100):
                        raise _dafny.Break("3")
                    d_618_i0_ = (d_618_i0_) + (1)
                    d_619_v1_: _dafny.Array
                    nw88_ = _dafny.Array(False, 20)
                    d_619_v1_ = nw88_
                    index89_ = default__.safeIndex(581, (d_619_v1_).length(0))
                    (d_619_v1_)[index89_] = (self).f27
                    d_620_v2_: int
                    d_620_v2_ = -741
                    d_621_v3_: _dafny.Seq
                    d_621_v3_ = _dafny.SeqWithoutIsStrInference([d_620_v2_, d_620_v2_])
                    index90_ = default__.safeIndex(581, (d_619_v1_).length(0))
                    (d_619_v1_)[index90_] = (d_621_v3_) < ((d_621_v3_) + (_dafny.SeqWithoutIsStrInference([d_620_v2_])))
                    d_622_v4_: str
                    d_622_v4_ = _dafny.CodePoint('d')
                    d_623_v5_: _dafny.Map
                    d_623_v5_ = _dafny.Map({d_622_v4_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gvfdsnym")))})
                    d_624_v6_: _dafny.Seq
                    d_624_v6_ = _dafny.SeqWithoutIsStrInference([d_623_v5_, _dafny.Map({d_622_v4_: -985})])
                    d_625_v7_: _dafny.Set
                    d_625_v7_ = _dafny.Set({(d_624_v6_)[default__.safeIndex(d_620_v2_, len(d_624_v6_))], d_623_v5_})
                    d_626_v8_: _dafny.Map
                    d_626_v8_ = _dafny.Map({d_620_v2_: (d_619_v1_)[default__.safeIndex(581, (d_619_v1_).length(0))]})
                    d_627_v9_: _dafny.MultiSet
                    d_627_v9_ = _dafny.MultiSet([d_623_v5_])
                    d_628_v10_: _dafny.MultiSet
                    d_628_v10_ = d_627_v9_
                    def iife79_():
                        coll45_ = _dafny.Set()
                        compr_45_: _dafny.Map
                        for compr_45_ in ((d_628_v10_)).Elements:
                            d_629_v11_: _dafny.Map = compr_45_
                            if (d_629_v11_) in ((d_628_v10_)):
                                coll45_ = coll45_.union(_dafny.Set([d_629_v11_]))
                        return _dafny.Set(coll45_)
                    d_625_v7_ = (_dafny.Set({d_623_v5_, d_623_v5_}) if ((d_626_v8_)[d_620_v2_] if (d_620_v2_) in (d_626_v8_) else (d_619_v1_)[default__.safeIndex(581, (d_619_v1_).length(0))]) else iife79_()
                    )
                    d_630_v12_: _dafny.Array
                    nw89_ = _dafny.Array(_dafny.Seq({}), 25)
                    d_630_v12_ = nw89_
                    d_630_v12_ = d_630_v12_
                    d_631_v13_: C0
                    nw90_ = C0()
                    nw90_.ctor__()
                    d_631_v13_ = nw90_
                    d_632_v14_: _dafny.Seq
                    d_632_v14_ = _dafny.SeqWithoutIsStrInference([d_631_v13_])
                    d_633_v15_: _dafny.Map
                    d_633_v15_ = _dafny.Map({d_622_v4_: (self).f27})
                    d_634_v16_: _dafny.Map
                    d_634_v16_ = _dafny.Map({len(d_633_v15_): d_622_v4_})
                    d_620_v2_ = ((len(d_632_v14_)) * (len(d_634_v16_))) + (d_620_v2_)
                    pass
            pass
        d_636_v17_: _dafny.Array
        nw91_ = _dafny.Array(int(0), 1)
        d_636_v17_ = nw91_
        d_637_v18_: int
        d_637_v18_ = 591
        rhs78_ = d_636_v17_
        rhs79_ = d_637_v18_
        rhs80_ = ((d_637_v18_) * (d_637_v18_)) + (d_637_v18_)
        rhs81_ = (0) - ((d_637_v18_) + (436))
        rhs82_ = d_637_v18_
        lhs64_ = globalState
        lhs65_ = globalState
        lhs66_ = globalState
        lhs67_ = globalState
        d_636_v17_ = rhs78_
        lhs64_.f2 = rhs79_
        lhs65_.f2 = rhs80_
        lhs66_.f6 = rhs81_
        lhs67_.f2 = rhs82_
        d_638_v19_: C0
        nw92_ = C0()
        nw92_.ctor__()
        d_638_v19_ = nw92_
        d_639_v20_: _dafny.Array
        nw93_ = _dafny.Array(_dafny.Map({}), 16)
        d_639_v20_ = nw93_
        d_640_v21_: _dafny.Array
        nw94_ = _dafny.Array(False, 23)
        d_640_v21_ = nw94_
        d_641_v22_: _dafny.Map
        d_641_v22_ = _dafny.Map({d_640_v21_: d_617_v0_})
        index91_ = default__.safeIndex(522, (d_639_v20_).length(0))
        (d_639_v20_)[index91_] = d_641_v22_
        index92_ = default__.safeIndex(522, (d_639_v20_).length(0))
        (d_639_v20_)[index92_] = d_641_v22_
        d_642_i2_: int
        d_642_i2_ = 0
        with _dafny.label("4"):
            while (self).f27:
                with _dafny.c_label("4"):
                    if (d_642_i2_) >= (100):
                        raise _dafny.Break("4")
                    d_642_i2_ = (d_642_i2_) + (1)
                    d_643_v23_: D0
                    d_643_v23_ = D0_DC1(True)
                    d_644_v24_: bool
                    d_645_v25_: _dafny.Array
                    d_646_v26_: _dafny.Map
                    d_647_v27_: _dafny.Seq
                    out46_: bool
                    out47_: _dafny.Array
                    out48_: _dafny.Map
                    out49_: _dafny.Seq
                    out46_, out47_, out48_, out49_ = (self).m8(d_643_v23_, globalState)
                    d_644_v24_ = out46_
                    d_645_v25_ = out47_
                    d_646_v26_ = out48_
                    d_647_v27_ = out49_
                    d_648_v28_: str
                    d_648_v28_ = _dafny.CodePoint('n')
                    d_649_v29_: _dafny.Map
                    d_649_v29_ = _dafny.Map({(self).f27: d_648_v28_})
                    d_650_v30_: D9
                    d_650_v30_ = D9_DC20(d_649_v29_)
                    pat_let_tv11_ = d_649_v29_
                    d_651_v31_: _dafny.MultiSet
                    def iife80_(_pat_let17_0):
                        def iife81_(d_652_dt__update__tmp_h0_):
                            def iife82_(_pat_let18_0):
                                def iife83_(d_653_dt__update_hcf30_h0_):
                                    return D9_DC20(d_653_dt__update_hcf30_h0_)
                                return iife83_(_pat_let18_0)
                            return iife82_(pat_let_tv11_)
                        return iife81_(_pat_let17_0)
                    d_651_v31_ = _dafny.MultiSet([d_644_v24_, (d_637_v18_) <= (len((iife80_(d_650_v30_)).cf30))])
                    d_651_v31_ = (d_651_v31_) - (d_651_v31_)
                    d_654_v32_: _dafny.Array
                    nw95_ = _dafny.Array(D6.default()(), 14)
                    d_654_v32_ = nw95_
                    d_655_v33_: _dafny.Array
                    nw96_ = _dafny.Array(None, 3)
                    nw96_[int(0)] = d_654_v32_
                    nw96_[int(1)] = (d_654_v32_ if True else d_654_v32_)
                    nw96_[int(2)] = d_654_v32_
                    d_655_v33_ = nw96_
                    index93_ = default__.safeIndex(416, (d_655_v33_).length(0))
                    (d_655_v33_)[index93_] = d_654_v32_
                    d_656_v34_: _dafny.Array
                    d_656_v34_ = d_654_v32_
                    index94_ = default__.safeIndex(416, (d_655_v33_).length(0))
                    (d_655_v33_)[index94_] = (d_656_v34_)
                    d_657_v35_: _dafny.MultiSet
                    d_657_v35_ = _dafny.MultiSet([226, d_637_v18_, d_637_v18_, d_637_v18_])
                    (globalState).f25 = (((d_657_v35_).set(d_637_v18_, default__.abs(d_637_v18_))) | (d_657_v35_)) != ((d_657_v35_) - (d_657_v35_))
                    pass
            pass
        d_658_v36_: _dafny.Seq
        d_658_v36_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({not((self).f27): d_637_v18_}))])
        d_659_v37_: _dafny.Map
        d_659_v37_ = _dafny.Map({d_658_v36_: d_637_v18_})
        d_660_v38_: D3
        d_660_v38_ = D3_DC8(d_659_v37_)
        pat_let_tv12_ = d_637_v18_
        def lambda16_(source14_):
            if source14_.is_DC9:
                d_661___mcc_h3_ = source14_.cf14
                d_662___mcc_h4_ = source14_.cf15
                d_663___mcc_h5_ = source14_.cf16
                d_664_cf16_ = d_663___mcc_h5_
                d_665_cf15_ = d_662___mcc_h4_
                d_666_cf14_ = d_661___mcc_h3_
                if (self).f27:
                    return D6_DC14((self).f27, d_665_cf15_)
                elif True:
                    return D6_DC14((self).f27, len(_dafny.SeqWithoutIsStrInference([d_666_cf14_, d_666_cf14_])))
            elif True:
                d_667___mcc_h6_ = source14_.cf13
                d_668_cf13_ = d_667___mcc_h6_
                return D6_DC14((self).f27, pat_let_tv12_)

        source13_ = lambda16_(d_660_v38_)
        if source13_.is_DC14:
            d_669___mcc_h0_ = source13_.cf22
            d_670___mcc_h1_ = source13_.cf23
            d_671_cf23_ = d_670___mcc_h1_
            d_672_cf22_ = d_669___mcc_h0_
            d_673_v39_: _dafny.Map
            d_673_v39_ = _dafny.Map({len(d_617_v0_): (self).f27})
            d_674_v40_: str
            d_674_v40_ = _dafny.CodePoint('i')
            d_675_v41_: _dafny.Set
            d_675_v41_ = _dafny.Set({d_674_v40_})
            d_676_v42_: D2
            d_676_v42_ = D2_DC6(d_637_v18_, d_658_v36_, (self).f27, _dafny.SeqWithoutIsStrInference([d_671_cf23_]), (self).f27)
            d_677_v43_: _dafny.MultiSet
            d_677_v43_ = _dafny.MultiSet([(_dafny.MultiSet([d_673_v39_])).cardinality, len(_dafny.SeqWithoutIsStrInference([d_617_v0_])), len(d_675_v41_), (d_676_v42_).cf7, (d_658_v36_)[default__.safeIndex(d_637_v18_, len(d_658_v36_))]])
            d_677_v43_ = (_dafny.MultiSet(d_658_v36_)).set(len(d_658_v36_), default__.abs(d_637_v18_))
            d_678_v44_: _dafny.Array
            nw97_ = _dafny.Array(_dafny.MultiSet({}), 14)
            d_678_v44_ = nw97_
            d_679_v45_: _dafny.MultiSet
            d_679_v45_ = _dafny.MultiSet([d_672_cf22_, False, d_672_cf22_])
            d_680_v46_: _dafny.Seq
            d_680_v46_ = _dafny.SeqWithoutIsStrInference([(self).f27, d_672_cf22_])
            d_681_v47_: D0
            d_681_v47_ = D0_DC1((self).f27)
            nw98_ = _dafny.Array(None, 14)
            nw98_[int(0)] = d_679_v45_
            nw98_[int(1)] = d_679_v45_
            nw98_[int(2)] = d_679_v45_
            nw98_[int(3)] = d_679_v45_
            nw98_[int(4)] = d_679_v45_
            nw98_[int(5)] = _dafny.MultiSet(d_680_v46_)
            nw98_[int(6)] = _dafny.MultiSet([d_672_cf22_, True, d_672_cf22_, d_672_cf22_, (self).f27])
            nw98_[int(7)] = d_679_v45_
            nw98_[int(8)] = (d_679_v45_).set((self).f27, default__.abs(d_671_cf23_))
            nw98_[int(9)] = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).fm4(d_681_v47_, globalState)]))
            nw98_[int(10)] = _dafny.MultiSet([(self).f27])
            nw98_[int(11)] = (_dafny.MultiSet([(self).f27, (self).f27, d_672_cf22_])).intersection(d_679_v45_)
            nw98_[int(12)] = d_679_v45_
            nw98_[int(13)] = (d_679_v45_).intersection(d_679_v45_)
            d_678_v44_ = nw98_
            d_682_v48_: D0
            d_682_v48_ = D0_DC2(556, (198 if d_672_cf22_ else -572))
            d_682_v48_ = D0_DC2(d_637_v18_, (d_637_v18_) + (default__.fm1(d_672_cf22_, d_680_v46_, (self).f27, globalState)))
            d_683_v49_: _dafny.Set
            d_683_v49_ = _dafny.Set({False, ((self).f27 if (self).f27 else not(d_672_cf22_)), d_672_cf22_})
            d_683_v49_ = (D11_DC24(d_683_v49_)).cf35
        elif True:
            d_684___mcc_h2_ = source13_.cf21
            d_685_cf21_ = d_684___mcc_h2_
            d_686_v50_: _dafny.Set
            d_686_v50_ = _dafny.Set({len(d_685_cf21_)})
            def iife84_():
                coll46_ = _dafny.Set()
                compr_46_: int
                for compr_46_ in _dafny.IntegerRange(728, 785):
                    d_687_v51_: int = compr_46_
                    if ((728) <= (d_687_v51_)) and ((d_687_v51_) < (785)):
                        def iife85_():
                            coll47_ = _dafny.Map()
                            compr_47_: int
                            for compr_47_ in _dafny.IntegerRange(924, -195):
                                d_688_v52_: int = compr_47_
                                if ((924) <= (d_688_v52_)) and ((d_688_v52_) < (-195)):
                                    coll47_[(d_688_v52_) * (-187)] = (self).f27
                            return _dafny.Map(coll47_)
                        coll46_ = coll46_.union(_dafny.Set([default__.safeDivisionInt(d_687_v51_, len(iife85_()
))]))
                return _dafny.Set(coll46_)
            if default__.fm0(d_617_v0_, (self).f27, (self).f27, (d_686_v50_) | (iife84_()
            ), globalState):
                d_689_v53_: D11
                d_689_v53_ = D11_DC25()
                d_690_v54_: _dafny.Set
                d_690_v54_ = _dafny.Set({True})
                d_691_v55_: D11
                d_691_v55_ = D11_DC24(d_690_v54_)
                d_692_v56_: _dafny.Seq
                d_692_v56_ = _dafny.SeqWithoutIsStrInference([d_691_v55_])
                d_693_v57_: _dafny.Map
                d_693_v57_ = _dafny.Map({d_692_v56_: (self).f27})
                d_694_v58_: _dafny.Map
                d_694_v58_ = _dafny.Map({d_689_v53_: (d_693_v57_ if (self).f27 else _dafny.Map({d_692_v56_: (self).f27}))})
                d_694_v58_ = (d_694_v58_).set(D11_DC25(), (d_693_v57_ if (self).f27 else d_693_v57_))
                (globalState).f6 = (0) - ((0) - (d_637_v18_))
                d_695_v59_: int
                d_696_v60_: bool
                d_697_v61_: int
                d_698_v62_: str
                out50_: int
                out51_: bool
                out52_: int
                out53_: str
                out50_, out51_, out52_, out53_ = (self).m9((0) - (d_637_v18_), globalState)
                d_695_v59_ = out50_
                d_696_v60_ = out51_
                d_697_v61_ = out52_
                d_698_v62_ = out53_
                d_695_v59_ = d_637_v18_
                (globalState).f6 = d_695_v59_
            elif True:
                d_699_v63_: D0
                d_699_v63_ = D0_DC1((d_685_cf21_)[default__.safeIndex(d_637_v18_, len(d_685_cf21_))])
                d_700_v64_: bool
                d_701_v65_: _dafny.Array
                d_702_v66_: _dafny.Map
                d_703_v67_: _dafny.Seq
                out54_: bool
                out55_: _dafny.Array
                out56_: _dafny.Map
                out57_: _dafny.Seq
                out54_, out55_, out56_, out57_ = (self).m8(d_699_v63_, globalState)
                d_700_v64_ = out54_
                d_701_v65_ = out55_
                d_702_v66_ = out56_
                d_703_v67_ = out57_
                d_704_v68_: D2
                d_704_v68_ = D2_DC6(d_637_v18_, _dafny.SeqWithoutIsStrInference([d_637_v18_]), d_700_v64_, d_703_v67_, (self).f27)
                d_705_v69_: D2
                d_705_v69_ = D2_DC6(d_637_v18_, (d_704_v68_).cf8, d_700_v64_, d_703_v67_, False)
                d_706_v70_: D2
                d_706_v70_ = D2_DC7(d_705_v69_)
                d_707_v71_: D2
                d_707_v71_ = D2_DC7(d_706_v70_)
                d_707_v71_ = d_707_v71_
                d_708_v72_: C0
                nw99_ = C0()
                nw99_.ctor__()
                d_708_v72_ = nw99_
                d_709_v73_: _dafny.Array
                nw100_ = _dafny.Array(_dafny.Array(None, 0), 12)
                d_709_v73_ = nw100_
                r2 = d_709_v73_
                (globalState).f2 = default__.safeModuloInt(d_637_v18_, d_637_v18_)
            if not(((self).f27) == ((D9_DC21((self).f27, d_637_v18_)).cf31)):
                d_710_v74_: _dafny.Map
                d_710_v74_ = _dafny.Map({620: (self).f27})
                d_710_v74_ = (d_710_v74_).set(d_637_v18_, ((self).f27 if not((self).f27) else (self).f27))
                d_711_v75_: _dafny.Set
                d_711_v75_ = _dafny.Set({d_636_v17_})
                d_712_v76_: _dafny.Array
                nw101_ = _dafny.Array(None, 19)
                nw101_[int(0)] = d_711_v75_
                nw101_[int(1)] = d_711_v75_
                nw101_[int(2)] = (d_711_v75_) - (d_711_v75_)
                nw101_[int(3)] = d_711_v75_
                nw101_[int(4)] = _dafny.Set({d_636_v17_, d_636_v17_})
                nw101_[int(5)] = d_711_v75_
                nw101_[int(6)] = d_711_v75_
                nw101_[int(7)] = (d_711_v75_ if (self).f27 else d_711_v75_)
                nw101_[int(8)] = (_dafny.Set({d_636_v17_})) - (d_711_v75_)
                nw101_[int(9)] = d_711_v75_
                nw101_[int(10)] = _dafny.Set({d_636_v17_, d_636_v17_})
                nw101_[int(11)] = d_711_v75_
                nw101_[int(12)] = d_711_v75_
                nw101_[int(13)] = (d_711_v75_) | (d_711_v75_)
                nw101_[int(14)] = d_711_v75_
                nw101_[int(15)] = d_711_v75_
                nw101_[int(16)] = (d_711_v75_ if True else d_711_v75_)
                nw101_[int(17)] = d_711_v75_
                nw101_[int(18)] = _dafny.Set({d_636_v17_, d_636_v17_})
                d_712_v76_ = nw101_
                index95_ = default__.safeIndex(587, (d_712_v76_).length(0))
                (d_712_v76_)[index95_] = d_711_v75_
                index96_ = default__.safeIndex(587, (d_712_v76_).length(0))
                (d_712_v76_)[index96_] = d_711_v75_
                d_658_v36_ = d_658_v36_
                (globalState).f6 = (0) - (-250)
                d_713_v77_: D0
                d_713_v77_ = D0_DC0(d_637_v18_)
                d_714_v78_: D0
                d_714_v78_ = D0_DC1((self).f27)
                d_715_v79_: D0
                d_715_v79_ = D0_DC1((self).fm4(d_714_v78_, globalState))
                d_716_v80_: bool
                d_717_v81_: _dafny.Seq
                out58_: bool
                out59_: _dafny.Seq
                out58_, out59_ = default__.m0(D0_DC3(d_713_v77_), d_715_v79_, globalState)
                d_716_v80_ = out58_
                d_717_v81_ = out59_
            elif True:
                d_718_v82_: str
                d_718_v82_ = _dafny.CodePoint('a')
                d_718_v82_ = default__.fm14(524, globalState)
                d_719_v83_: C0
                nw102_ = C0()
                nw102_.ctor__()
                d_719_v83_ = nw102_
                (globalState).f6 = d_637_v18_
                d_720_v84_: _dafny.Array
                nw103_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 28)
                d_720_v84_ = nw103_
                index97_ = default__.safeIndex(271, (d_720_v84_).length(0))
                (d_720_v84_)[index97_] = d_617_v0_
                index98_ = default__.safeIndex(271, (d_720_v84_).length(0))
                (d_720_v84_)[index98_] = (d_617_v0_).set(default__.safeIndex(len((d_685_cf21_ if (self).f27 else d_685_cf21_)), len(d_617_v0_)), d_718_v82_)
                r1 = d_617_v0_
            d_721_v85_: C0
            nw104_ = C0()
            nw104_.ctor__()
            d_721_v85_ = nw104_
            d_722_v86_: _dafny.MultiSet
            d_722_v86_ = _dafny.MultiSet([d_617_v0_])
            d_722_v86_ = d_722_v86_
        d_723_v87_: _dafny.Map
        d_723_v87_ = _dafny.Map({(self).f27: d_640_v21_})
        d_724_v88_: _dafny.Seq
        d_724_v88_ = _dafny.SeqWithoutIsStrInference([d_640_v21_])
        d_725_v89_: _dafny.Set
        d_725_v89_ = _dafny.Set({(self).f27, not(False)})
        d_726_v90_: _dafny.Array
        nw105_ = _dafny.Array(None, 27)
        nw105_[int(0)] = d_640_v21_
        nw105_[int(1)] = ((d_723_v87_)[(self).f27] if ((self).f27) in (d_723_v87_) else d_640_v21_)
        nw105_[int(2)] = d_640_v21_
        nw105_[int(3)] = d_640_v21_
        nw105_[int(4)] = d_640_v21_
        nw105_[int(5)] = d_640_v21_
        nw105_[int(6)] = d_640_v21_
        nw105_[int(7)] = d_640_v21_
        nw105_[int(8)] = d_640_v21_
        nw105_[int(9)] = d_640_v21_
        nw105_[int(10)] = d_640_v21_
        nw105_[int(11)] = d_640_v21_
        nw105_[int(12)] = d_640_v21_
        nw105_[int(13)] = (d_640_v21_ if (self).f27 else d_640_v21_)
        nw105_[int(14)] = d_640_v21_
        nw105_[int(15)] = d_640_v21_
        nw105_[int(16)] = d_640_v21_
        nw105_[int(17)] = d_640_v21_
        nw105_[int(18)] = d_640_v21_
        nw105_[int(19)] = d_640_v21_
        nw105_[int(20)] = d_640_v21_
        nw105_[int(21)] = d_640_v21_
        nw105_[int(22)] = d_640_v21_
        nw105_[int(23)] = (d_724_v88_)[default__.safeIndex((0) - (len(d_725_v89_)), len(d_724_v88_))]
        nw105_[int(24)] = d_640_v21_
        nw105_[int(25)] = (d_640_v21_ if (self).f27 else d_640_v21_)
        nw105_[int(26)] = d_640_v21_
        d_726_v90_ = nw105_
        r0 = d_726_v90_
        d_727_v91_: _dafny.Map
        d_727_v91_ = _dafny.Map({d_636_v17_: d_617_v0_})
        r1 = ((d_727_v91_)[d_636_v17_] if (d_636_v17_) in (d_727_v91_) else (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "urxfd")) if not((self).f27) else d_617_v0_))
        d_728_v92_: _dafny.Array
        nw106_ = _dafny.Array(_dafny.Array(None, 0), 24)
        d_728_v92_ = nw106_
        r2 = d_728_v92_
        r3 = False
        return r0, r1, r2, r3

    def m8(self, p0, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: _dafny.Map = _dafny.Map({})
        r3: _dafny.Seq = _dafny.Seq({})
        d_729_v0_: int
        d_729_v0_ = 55
        d_730_v1_: _dafny.Seq
        d_730_v1_ = _dafny.SeqWithoutIsStrInference([(self).f27])
        (globalState).f6 = (0) - (default__.fm1((d_729_v0_) <= (len((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False, (self).f27]))])).set(default__.safeIndex(d_729_v0_, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False, (self).f27]))]))), d_729_v0_))), (d_730_v1_).set(default__.safeIndex(d_729_v0_, len(d_730_v1_)), (self).f27), (self).f27, globalState))
        d_731_v2_: _dafny.Seq
        d_731_v2_ = _dafny.SeqWithoutIsStrInference([d_729_v0_])
        d_732_v3_: _dafny.Map
        d_732_v3_ = _dafny.Map({(d_729_v0_) <= ((0) - ((_dafny.MultiSet(d_731_v2_)).cardinality)): (self).f27})
        d_732_v3_ = (d_732_v3_).set((self).f27, (self).f27)
        d_733_v4_: _dafny.Array
        def lambda17_(d_734_v0_):
            def lambda18_(d_735_i1_):
                return (d_735_i1_) + (d_734_v0_)

            return lambda18_

        init8_ = lambda17_(d_729_v0_)
        nw107_ = _dafny.Array(None, 14)
        for i0_8_ in range(nw107_.length(0)):
            nw107_[i0_8_] = init8_(i0_8_)
        d_733_v4_ = nw107_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_733_v4_).length(0)):
            d_736_i0_: int = guard_loop_1_
            if (True) and (((0) <= (d_736_i0_)) and ((d_736_i0_) < ((d_733_v4_).length(0)))):
                (d_733_v4_)[(d_736_i0_)] = (d_736_i0_) * (d_729_v0_)
        d_737_v5_: D7
        d_737_v5_ = D7_DC17(d_729_v0_, d_729_v0_, (self).f27)
        pat_let_tv13_ = d_729_v0_
        pat_let_tv14_ = d_729_v0_
        pat_let_tv15_ = d_729_v0_
        pat_let_tv16_ = d_729_v0_
        pat_let_tv17_ = d_729_v0_
        pat_let_tv18_ = d_729_v0_
        def lambda19_(source15_):
            if source15_.is_DC16:
                return pat_let_tv13_
            elif source15_.is_DC17:
                d_738___mcc_h0_ = source15_.cf25
                d_739___mcc_h1_ = source15_.cf26
                d_740___mcc_h2_ = source15_.cf27
                d_741_cf27_ = d_740___mcc_h2_
                d_742_cf26_ = d_739___mcc_h1_
                d_743_cf25_ = d_738___mcc_h0_
                def iife86_():
                    coll48_ = _dafny.Map()
                    compr_48_: int
                    for compr_48_ in _dafny.IntegerRange(-814, 908):
                        d_744_v6_: int = compr_48_
                        if ((-814) <= (d_744_v6_)) and ((d_744_v6_) < (908)):
                            coll48_[default__.safeModuloInt(d_744_v6_, (_dafny.MultiSet([pat_let_tv15_])).cardinality)] = (self).f27
                    return _dafny.Map(coll48_)
                return len((_dafny.Map({d_741_cf27_: _dafny.Map({pat_let_tv14_: d_741_cf27_})})) | (_dafny.Map({d_741_cf27_: iife86_()
                })))
            elif source15_.is_DC15:
                d_745___mcc_h3_ = source15_.cf24
                d_746_cf24_ = d_745___mcc_h3_
                return (pat_let_tv16_) - (pat_let_tv17_)
            elif True:
                d_747___mcc_h4_ = source15_.cf28
                d_748_cf28_ = d_747___mcc_h4_
                return pat_let_tv18_

        def iife87_(_pat_let19_0):
            def iife88_(d_749_dt__update__tmp_h0_):
                def iife89_(_pat_let20_0):
                    def iife90_(d_750_dt__update_hcf25_h0_):
                        def iife91_(_pat_let21_0):
                            def iife92_(d_751_dt__update_hcf27_h0_):
                                return D7_DC17(d_750_dt__update_hcf25_h0_, (d_749_dt__update__tmp_h0_).cf26, d_751_dt__update_hcf27_h0_)
                            return iife92_(_pat_let21_0)
                        return iife91_((self).f27)
                    return iife90_(_pat_let20_0)
                return iife89_(179)
            return iife88_(_pat_let19_0)
        (globalState).f2 = lambda19_(iife87_(d_737_v5_))
        source16_ = d_733_v4_
        d_752___mcc_h5_ = source16_
        d_753_cf5_ = d_752___mcc_h5_
        d_754_v7_: _dafny.Array
        def lambda20_(d_755_v1_, d_756_v2_, d_757_v0_):
            def lambda21_(d_758_i2_):
                return D6_DC14((D2_DC6(len(_dafny.SeqWithoutIsStrInference([(self).f27])), _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(d_755_v1_)).cardinality for d_759_i3_ in range(default__.abs(723))]), (self).f27, (d_756_v2_).set(default__.safeIndex(d_757_v0_, len(d_756_v2_)), (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([d_757_v0_]))])).cardinality), (self).f27)).cf9, (d_756_v2_)[default__.safeIndex(len(_dafny.Set({not((self).f27)})), len(d_756_v2_))])

            return lambda21_

        init9_ = lambda20_(d_730_v1_, d_731_v2_, d_729_v0_)
        nw108_ = _dafny.Array(None, 25)
        for i0_9_ in range(nw108_.length(0)):
            nw108_[i0_9_] = init9_(i0_9_)
        d_754_v7_ = nw108_
        d_760_v8_: D6
        d_760_v8_ = D6_DC14((self).f27, 109)
        index99_ = default__.safeIndex(189, (d_754_v7_).length(0))
        (d_754_v7_)[index99_] = d_760_v8_
        index100_ = default__.safeIndex(189, (d_754_v7_).length(0))
        (d_754_v7_)[index100_] = d_760_v8_
        d_761_v9_: _dafny.Array
        def lambda22_(d_762_i4_):
            return True

        init10_ = lambda22_
        nw109_ = _dafny.Array(None, 17)
        for i0_10_ in range(nw109_.length(0)):
            nw109_[i0_10_] = init10_(i0_10_)
        d_761_v9_ = nw109_
        d_763_v10_: _dafny.MultiSet
        d_763_v10_ = _dafny.MultiSet([d_761_v9_, d_761_v9_])
        d_764_v11_: _dafny.Map
        d_764_v11_ = _dafny.Map({(d_763_v10_) == (d_763_v10_): _dafny.SeqWithoutIsStrInference([(self).f27])})
        d_764_v11_ = (d_764_v11_).set((self).f27, (_dafny.SeqWithoutIsStrInference([(self).f27]) if (self).f27 else d_730_v1_))
        index101_ = default__.safeIndex(490, (d_761_v9_).length(0))
        (d_761_v9_)[index101_] = not (False) or ((self).fm4(p0, globalState))
        index102_ = default__.safeIndex(490, (d_761_v9_).length(0))
        (d_761_v9_)[index102_] = not(not((self).f27))
        d_765_v12_: D0
        d_765_v12_ = D0_DC1((d_761_v9_)[default__.safeIndex(490, (d_761_v9_).length(0))])
        pat_let_tv19_ = d_765_v12_
        d_766_v13_: bool
        d_767_v14_: _dafny.Seq
        out60_: bool
        out61_: _dafny.Seq
        def iife93_(_pat_let22_0):
            def iife94_(d_768_dt__update__tmp_h1_):
                def iife95_(_pat_let23_0):
                    def iife96_(d_769_dt__update_hcf4_h0_):
                        return D0_DC3(d_769_dt__update_hcf4_h0_)
                    return iife96_(_pat_let23_0)
                return iife95_(pat_let_tv19_)
            return iife94_(_pat_let22_0)
        out60_, out61_ = default__.m0(iife93_(D0_DC3((self).fm5(globalState))), p0, globalState)
        d_766_v13_ = out60_
        d_767_v14_ = out61_
        (globalState).f25 = (self).f27
        r0 = not ((self).f27) or ((self).f27)
        d_770_v15_: _dafny.Seq
        d_770_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "okjnfjglr"))
        d_771_v16_: str
        d_771_v16_ = _dafny.CodePoint('v')
        nw110_ = _dafny.Array(None, 22)
        nw110_[int(0)] = d_770_v15_
        nw110_[int(1)] = d_770_v15_
        nw110_[int(2)] = d_770_v15_
        nw110_[int(3)] = d_770_v15_
        nw110_[int(4)] = d_770_v15_
        nw110_[int(5)] = (d_770_v15_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lik")))
        nw110_[int(6)] = d_770_v15_
        nw110_[int(7)] = d_770_v15_
        nw110_[int(8)] = d_770_v15_
        nw110_[int(9)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "orpsaowp"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vcnq")))
        nw110_[int(10)] = (d_770_v15_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wlkyidwjr")))
        nw110_[int(11)] = d_770_v15_
        nw110_[int(12)] = d_770_v15_
        nw110_[int(13)] = d_770_v15_
        nw110_[int(14)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fpxiwlyx"))
        nw110_[int(15)] = ((d_770_v15_) + (d_770_v15_)).set(default__.safeIndex(d_729_v0_, len((d_770_v15_) + (d_770_v15_))), d_771_v16_)
        nw110_[int(16)] = (d_770_v15_) + (d_770_v15_)
        nw110_[int(17)] = (d_770_v15_) + (d_770_v15_)
        nw110_[int(18)] = d_770_v15_
        nw110_[int(19)] = (d_770_v15_) + (d_770_v15_)
        nw110_[int(20)] = _dafny.SeqWithoutIsStrInference([(d_771_v16_) for d_772_i5_ in range(default__.abs(800))])
        nw110_[int(21)] = (d_770_v15_) + (d_770_v15_)
        r1 = nw110_
        d_773_v17_: _dafny.Map
        d_773_v17_ = _dafny.Map({(len(d_770_v15_)) * (d_729_v0_): d_729_v0_})
        r2 = d_773_v17_
        r3 = (_dafny.SeqWithoutIsStrInference([d_729_v0_, d_729_v0_])) + (d_731_v2_)
        return r0, r1, r2, r3

    def m9(self, p0, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: int = int(0)
        r3: str = _dafny.CodePoint('D')
        d_774_v0_: _dafny.Seq
        d_774_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yknidu"))
        d_775_v1_: _dafny.Seq
        d_775_v1_ = _dafny.SeqWithoutIsStrInference([d_774_v0_])
        d_776_v2_: str
        d_776_v2_ = _dafny.CodePoint('v')
        d_777_v3_: _dafny.MultiSet
        d_777_v3_ = _dafny.MultiSet([d_776_v2_, d_776_v2_, _dafny.CodePoint('n'), d_776_v2_])
        d_778_v4_: _dafny.Map
        d_778_v4_ = _dafny.Map({p0: (default__.fm18((self).f27, globalState)).set(default__.safeIndex(p0, len(default__.fm18((self).f27, globalState))), d_774_v0_)})
        d_779_v5_: _dafny.Seq
        d_779_v5_ = _dafny.SeqWithoutIsStrInference([d_775_v1_])
        d_780_v6_: D13
        d_780_v6_ = D13_DC27(d_775_v1_)
        d_781_v7_: _dafny.Array
        nw111_ = _dafny.Array(None, 28)
        nw111_[int(0)] = (d_775_v1_).set(default__.safeIndex(((d_777_v3_)[d_776_v2_] if (d_776_v2_) in (d_777_v3_) else p0), len(d_775_v1_)), d_774_v0_)
        nw111_[int(1)] = d_775_v1_
        nw111_[int(2)] = ((d_778_v4_)[p0] if (p0) in (d_778_v4_) else d_775_v1_)
        nw111_[int(3)] = (d_775_v1_) + (d_775_v1_)
        nw111_[int(4)] = _dafny.SeqWithoutIsStrInference([d_774_v0_ for d_782_i0_ in range(default__.abs(474))])
        nw111_[int(5)] = d_775_v1_
        nw111_[int(6)] = (d_775_v1_) + (d_775_v1_)
        nw111_[int(7)] = (d_775_v1_) + (d_775_v1_)
        nw111_[int(8)] = (d_775_v1_) + (d_775_v1_)
        nw111_[int(9)] = d_775_v1_
        nw111_[int(10)] = (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cixnateea")) for d_783_i1_ in range(default__.abs(989))]) if (self).f27 else (d_779_v5_)[default__.safeIndex(p0, len(d_779_v5_))])
        nw111_[int(11)] = d_775_v1_
        nw111_[int(12)] = d_775_v1_
        nw111_[int(13)] = d_775_v1_
        nw111_[int(14)] = (d_775_v1_) + (_dafny.SeqWithoutIsStrInference([d_774_v0_ for d_784_i2_ in range(default__.abs(512))]))
        nw111_[int(15)] = d_775_v1_
        nw111_[int(16)] = _dafny.SeqWithoutIsStrInference([(d_774_v0_).set(default__.safeIndex((0) - ((0) - (p0)), len(d_774_v0_)), d_776_v2_)])
        nw111_[int(17)] = d_775_v1_
        nw111_[int(18)] = (d_775_v1_) + (_dafny.SeqWithoutIsStrInference([d_774_v0_]))
        nw111_[int(19)] = (d_775_v1_) + (d_775_v1_)
        nw111_[int(20)] = d_775_v1_
        nw111_[int(21)] = d_775_v1_
        nw111_[int(22)] = _dafny.SeqWithoutIsStrInference([d_774_v0_, d_774_v0_])
        nw111_[int(23)] = (D13_DC27((d_775_v1_).set(default__.safeIndex(p0, len(d_775_v1_)), d_774_v0_))).cf37
        nw111_[int(24)] = d_775_v1_
        nw111_[int(25)] = d_775_v1_
        nw111_[int(26)] = (d_780_v6_).cf37
        nw111_[int(27)] = d_775_v1_
        d_781_v7_ = nw111_
        index103_ = default__.safeIndex(63, (d_781_v7_).length(0))
        (d_781_v7_)[index103_] = d_775_v1_
        index104_ = default__.safeIndex(63, (d_781_v7_).length(0))
        (d_781_v7_)[index104_] = (d_775_v1_) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "esfknh")) for d_785_i3_ in range(default__.abs(566))]))
        (globalState).f12 = (p0) == (default__.safeModuloInt(546, 503))
        d_786_v8_: _dafny.MultiSet
        d_786_v8_ = _dafny.MultiSet([(self).f27])
        d_787_v9_: _dafny.Seq
        d_787_v9_ = _dafny.SeqWithoutIsStrInference([(d_786_v8_).cardinality, p0])
        d_788_v10_: _dafny.Map
        d_788_v10_ = _dafny.Map({d_787_v9_: (self).f27})
        hi5_ = default__.safeModuloInt(len(d_788_v10_), p0)
        for d_789_i4_ in range((p0) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ifx")))), hi5_):
            d_790_v11_: C0
            nw112_ = C0()
            nw112_.ctor__()
            d_790_v11_ = nw112_
            d_791_v12_: _dafny.Map
            d_791_v12_ = _dafny.Map({((self).f27) or ((self).f27): (self).f27})
            d_791_v12_ = (d_791_v12_).set(not(not((self).f27)), True)
            d_792_v13_: C0
            nw113_ = C0()
            nw113_.ctor__()
            d_792_v13_ = nw113_
            d_793_v14_: D0
            d_793_v14_ = D0_DC1((self).f27)
            d_794_v15_: _dafny.Seq
            d_794_v15_ = _dafny.SeqWithoutIsStrInference([(self).f27, (self).f27, ((self).fm4(d_793_v14_, globalState)) == ((self).f27)])
            d_795_v16_: _dafny.Array
            nw114_ = _dafny.Array(int(0), 6)
            d_795_v16_ = nw114_
            index105_ = default__.safeIndex(423, (d_795_v16_).length(0))
            (d_795_v16_)[index105_] = p0
            index106_ = default__.safeIndex(423, (d_795_v16_).length(0))
            rhs83_ = (d_794_v15_) + (d_794_v15_)
            rhs84_ = d_795_v16_
            rhs85_ = (self).f27
            rhs86_ = (p0) - (d_789_i4_)
            rhs87_ = (d_794_v15_)[default__.safeIndex(p0, len(d_794_v15_))]
            lhs68_ = globalState
            lhs69_ = d_795_v16_
            lhs70_ = default__.safeIndex(423, (d_795_v16_).length(0))
            lhs71_ = globalState
            d_794_v15_ = rhs83_
            d_795_v16_ = rhs84_
            lhs68_.f14 = rhs85_
            lhs69_[lhs70_] = rhs86_
            lhs71_.f14 = rhs87_
        d_796_v17_: _dafny.Seq
        d_796_v17_ = _dafny.SeqWithoutIsStrInference([(self).f27, True, (self).f27, (self).f27, (self).f27])
        d_797_v18_: _dafny.Array
        nw115_ = _dafny.Array(int(0), 29)
        d_797_v18_ = nw115_
        d_798_v19_: _dafny.Array
        nw116_ = _dafny.Array(None, 2)
        nw116_[int(0)] = d_797_v18_
        nw116_[int(1)] = d_797_v18_
        d_798_v19_ = nw116_
        index107_ = default__.safeIndex(843, (d_798_v19_).length(0))
        (d_798_v19_)[index107_] = d_797_v18_
        d_799_v20_: _dafny.Set
        d_799_v20_ = _dafny.Set({(self).f27, (self).f27})
        d_800_v21_: D3
        d_800_v21_ = D3_DC9((self).f27, p0, len(d_799_v20_))
        d_801_v22_: D9
        d_801_v22_ = D9_DC21((self).f27, p0)
        d_802_v23_: _dafny.Set
        d_802_v23_ = _dafny.Set({p0, p0})
        d_803_v24_: _dafny.Array
        nw117_ = _dafny.Array(None, 29)
        nw117_[int(0)] = (self).f27
        nw117_[int(1)] = (self).f27
        nw117_[int(2)] = (self).f27
        nw117_[int(3)] = True
        nw117_[int(4)] = (self).f27
        nw117_[int(5)] = (d_800_v21_).cf14
        nw117_[int(6)] = (self).f27
        nw117_[int(7)] = (self).f27
        nw117_[int(8)] = (self).f27
        nw117_[int(9)] = (self).f27
        nw117_[int(10)] = (d_801_v22_).cf31
        nw117_[int(11)] = (self).f27
        nw117_[int(12)] = (self).f27
        nw117_[int(13)] = not((self).f27)
        nw117_[int(14)] = False
        nw117_[int(15)] = (self).f27
        nw117_[int(16)] = (self).f27
        nw117_[int(17)] = (self).f27
        nw117_[int(18)] = (self).f27
        nw117_[int(19)] = (self).f27
        nw117_[int(20)] = (self).f27
        nw117_[int(21)] = (self).f27
        nw117_[int(22)] = (self).f27
        nw117_[int(23)] = (self).f27
        nw117_[int(24)] = (self).f27
        nw117_[int(25)] = (self).f27
        nw117_[int(26)] = (self).f27
        nw117_[int(27)] = True
        nw117_[int(28)] = default__.fm0(d_774_v0_, (self).f27, (self).f27, d_802_v23_, globalState)
        d_803_v24_ = nw117_
        d_804_v25_: _dafny.Seq
        d_804_v25_ = _dafny.SeqWithoutIsStrInference([d_803_v24_, d_803_v24_, d_803_v24_, d_803_v24_, d_803_v24_])
        d_805_v26_: _dafny.Map
        d_805_v26_ = _dafny.Map({(self).f27: (_dafny.SeqWithoutIsStrInference([p0 for d_806_i5_ in range(default__.abs(190))])) + (d_787_v9_)})
        index108_ = default__.safeIndex(843, (d_798_v19_).length(0))
        rhs88_ = (d_796_v17_ if True else d_796_v17_)
        rhs89_ = p0
        rhs90_ = d_797_v18_
        rhs91_ = _dafny.SeqWithoutIsStrInference([(d_803_v24_ if (self).f27 else d_803_v24_), d_803_v24_])
        rhs92_ = (d_805_v26_ if (self).f27 else d_805_v26_)
        lhs72_ = d_798_v19_
        lhs73_ = default__.safeIndex(843, (d_798_v19_).length(0))
        d_796_v17_ = rhs88_
        r2 = rhs89_
        lhs72_[lhs73_] = rhs90_
        d_804_v25_ = rhs91_
        d_805_v26_ = rhs92_
        d_807_v27_: _dafny.MultiSet
        d_807_v27_ = _dafny.MultiSet([p0, 282])
        d_808_v28_: D0
        d_808_v28_ = D0_DC2(-830, p0)
        hi6_ = (d_808_v28_).cf2
        for d_809_i6_ in range((0) - (((d_807_v27_) | (d_807_v27_)).cardinality), hi6_):
            d_810_v29_: _dafny.Map
            d_810_v29_ = _dafny.Map({d_809_i6_: not((self).f27)})
            if ((d_810_v29_)[p0] if (p0) in (d_810_v29_) else not((self).f27)):
                index109_ = default__.safeIndex(0, (d_803_v24_).length(0))
                (d_803_v24_)[index109_] = (self).f27
                index110_ = default__.safeIndex(0, (d_803_v24_).length(0))
                (d_803_v24_)[index110_] = True
                d_811_v30_: str
                d_811_v30_ = d_776_v2_
                d_812_v31_: _dafny.Array
                nw118_ = _dafny.Array(None, 29)
                nw118_[int(0)] = d_776_v2_
                nw118_[int(1)] = d_776_v2_
                nw118_[int(2)] = d_776_v2_
                nw118_[int(3)] = d_776_v2_
                nw118_[int(4)] = _dafny.CodePoint('c')
                nw118_[int(5)] = d_776_v2_
                nw118_[int(6)] = d_776_v2_
                nw118_[int(7)] = d_776_v2_
                nw118_[int(8)] = d_776_v2_
                nw118_[int(9)] = d_776_v2_
                nw118_[int(10)] = d_776_v2_
                nw118_[int(11)] = d_776_v2_
                nw118_[int(12)] = d_776_v2_
                nw118_[int(13)] = d_776_v2_
                nw118_[int(14)] = _dafny.CodePoint('j')
                nw118_[int(15)] = (d_776_v2_ if (d_803_v24_)[default__.safeIndex(0, (d_803_v24_).length(0))] else d_776_v2_)
                nw118_[int(16)] = _dafny.CodePoint('h')
                nw118_[int(17)] = d_776_v2_
                nw118_[int(18)] = default__.fm14(p0, globalState)
                nw118_[int(19)] = d_776_v2_
                nw118_[int(20)] = d_776_v2_
                nw118_[int(21)] = d_776_v2_
                nw118_[int(22)] = _dafny.CodePoint('e')
                nw118_[int(23)] = d_776_v2_
                nw118_[int(24)] = d_776_v2_
                nw118_[int(25)] = d_776_v2_
                nw118_[int(26)] = d_776_v2_
                nw118_[int(27)] = (d_811_v30_)
                nw118_[int(28)] = default__.fm14(782, globalState)
                d_812_v31_ = nw118_
                index111_ = default__.safeIndex(908, (d_812_v31_).length(0))
                (d_812_v31_)[index111_] = default__.fm14(d_809_i6_, globalState)
                index112_ = default__.safeIndex(908, (d_812_v31_).length(0))
                (d_812_v31_)[index112_] = d_776_v2_
                index113_ = default__.safeIndex(0, (d_803_v24_).length(0))
                (d_803_v24_)[index113_] = ((d_803_v24_)[default__.safeIndex(0, (d_803_v24_).length(0))] if (self).f27 else (d_803_v24_)[default__.safeIndex(0, (d_803_v24_).length(0))])
                (globalState).f12 = (d_803_v24_)[default__.safeIndex(0, (d_803_v24_).length(0))]
                d_813_v32_: _dafny.Array
                def lambda23_(d_814_i6_):
                    def lambda24_(d_815_i7_):
                        return _dafny.Map({(self).f27: d_814_i6_})

                    return lambda24_

                init11_ = lambda23_(d_809_i6_)
                nw119_ = _dafny.Array(None, 8)
                for i0_11_ in range(nw119_.length(0)):
                    nw119_[i0_11_] = init11_(i0_11_)
                d_813_v32_ = nw119_
                d_816_v34_: D2
                d_816_v34_ = D2_DC6(776, _dafny.SeqWithoutIsStrInference([p0]), (self).f27, d_787_v9_, (d_803_v24_)[default__.safeIndex(0, (d_803_v24_).length(0))])
                d_817_v35_: _dafny.Map
                d_817_v35_ = _dafny.Map({(d_816_v34_).cf7: p0})
                d_818_v36_: _dafny.Map
                def iife97_():
                    coll49_ = _dafny.Map()
                    compr_49_: int
                    for compr_49_ in (d_817_v35_).keys.Elements:
                        d_819_v33_: int = compr_49_
                        if (d_819_v33_) in (d_817_v35_):
                            coll49_[default__.safeModuloInt(d_819_v33_, len(d_774_v0_))] = (d_812_v31_)[default__.safeIndex(908, (d_812_v31_).length(0))]
                    return _dafny.Map(coll49_)
                d_818_v36_ = _dafny.Map({d_776_v2_: iife97_()
                })
                d_820_v37_: D14
                d_820_v37_ = D14_DC30(d_818_v36_)
                index114_ = default__.safeIndex(0, (d_803_v24_).length(0))
                rhs93_ = (d_803_v24_)[default__.safeIndex(0, (d_803_v24_).length(0))]
                rhs94_ = (d_803_v24_)[default__.safeIndex(0, (d_803_v24_).length(0))]
                rhs95_ = (self).fm2(default__.fm1((d_803_v24_)[default__.safeIndex(0, (d_803_v24_).length(0))], d_796_v17_, default__.fm0((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))).set(default__.safeIndex(len(d_799_v20_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o")))), (d_812_v31_)[default__.safeIndex(908, (d_812_v31_).length(0))]), (d_803_v24_)[default__.safeIndex(0, (d_803_v24_).length(0))], (self).f27, _dafny.Set({p0, d_809_i6_}), globalState), globalState), (d_820_v37_).cf42, globalState)
                rhs96_ = d_813_v32_
                rhs97_ = ((default__.fm12(p0, (d_803_v24_)[default__.safeIndex(0, (d_803_v24_).length(0))], (d_803_v24_)[default__.safeIndex(0, (d_803_v24_).length(0))], globalState) if (self).f27 else d_787_v9_)) + (_dafny.SeqWithoutIsStrInference([len(d_774_v0_), 901, len(_dafny.Map({_dafny.CodePoint('x'): p0}))]))
                lhs74_ = globalState
                lhs75_ = d_803_v24_
                lhs76_ = default__.safeIndex(0, (d_803_v24_).length(0))
                lhs74_.f5 = rhs93_
                lhs75_[lhs76_] = rhs94_
                d_774_v0_ = rhs95_
                d_813_v32_ = rhs96_
                d_787_v9_ = rhs97_
            elif True:
                (globalState).f14 = (d_776_v2_) not in (d_774_v0_)
                d_821_v38_: _dafny.Array
                def lambda25_(d_822_p0_, d_823_i6_):
                    def lambda26_(d_824_i8_):
                        return D13_DC29(d_822_p0_, len(_dafny.SeqWithoutIsStrInference([d_823_i6_ for d_825_i9_ in range(default__.abs(375))])), (self).f27)

                    return lambda26_

                init12_ = lambda25_(p0, d_809_i6_)
                nw120_ = _dafny.Array(None, 27)
                for i0_12_ in range(nw120_.length(0)):
                    nw120_[i0_12_] = init12_(i0_12_)
                d_821_v38_ = nw120_
                d_826_v39_: _dafny.Map
                d_826_v39_ = _dafny.Map({(self).f27: d_821_v38_})
                d_827_v40_: _dafny.Seq
                d_827_v40_ = _dafny.SeqWithoutIsStrInference([d_826_v39_])
                d_828_v41_: _dafny.Seq
                d_828_v41_ = _dafny.SeqWithoutIsStrInference([d_826_v39_, ((d_827_v40_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([p0])), len(d_827_v40_))]).set(True, d_821_v38_), d_826_v39_])
                rhs98_ = (_dafny.Set({p0})).intersection((d_802_v23_) | (d_802_v23_))
                rhs99_ = (((d_827_v40_)[default__.safeIndex(len(d_774_v0_), len(d_827_v40_))]) | (d_826_v39_)) != (d_826_v39_)
                lhs77_ = globalState
                d_802_v23_ = rhs98_
                lhs77_.f25 = rhs99_
                d_829_v42_: _dafny.Map
                d_829_v42_ = _dafny.Map({d_803_v24_: d_809_i6_})
                d_830_v43_: _dafny.Seq
                d_830_v43_ = _dafny.SeqWithoutIsStrInference([d_829_v42_, d_829_v42_])
                d_831_v44_: D15
                d_831_v44_ = D15_DC32(d_829_v42_)
                pat_let_tv20_ = d_829_v42_
                d_832_v45_: _dafny.Map
                def iife98_(_pat_let24_0):
                    def iife99_(d_833_dt__update__tmp_h0_):
                        def iife100_(_pat_let25_0):
                            def iife101_(d_834_dt__update_hcf44_h0_):
                                return D15_DC32(d_834_dt__update_hcf44_h0_)
                            return iife101_(_pat_let25_0)
                        return iife100_(pat_let_tv20_)
                    return iife99_(_pat_let24_0)
                d_832_v45_ = _dafny.Map({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f27, (self).f27, True])): (iife98_(d_831_v44_)).cf44})
                d_835_v46_: _dafny.Array
                nw121_ = _dafny.Array(None, 14)
                nw121_[int(0)] = d_829_v42_
                nw121_[int(1)] = ((d_830_v43_)[default__.safeIndex(p0, len(d_830_v43_))]) | (d_829_v42_)
                nw121_[int(2)] = (d_830_v43_)[default__.safeIndex(p0, len(d_830_v43_))]
                nw121_[int(3)] = d_829_v42_
                nw121_[int(4)] = d_829_v42_
                nw121_[int(5)] = d_829_v42_
                nw121_[int(6)] = ((d_832_v45_)[d_786_v8_] if (d_786_v8_) in (d_832_v45_) else d_829_v42_)
                nw121_[int(7)] = (d_829_v42_).set(d_803_v24_, d_809_i6_)
                nw121_[int(8)] = d_829_v42_
                nw121_[int(9)] = d_829_v42_
                nw121_[int(10)] = (d_829_v42_) | (d_829_v42_)
                nw121_[int(11)] = d_829_v42_
                nw121_[int(12)] = _dafny.Map({d_803_v24_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hhka")))})
                nw121_[int(13)] = d_829_v42_
                d_835_v46_ = nw121_
                index115_ = default__.safeIndex(658, (d_835_v46_).length(0))
                (d_835_v46_)[index115_] = d_829_v42_
                index116_ = default__.safeIndex(658, (d_835_v46_).length(0))
                (d_835_v46_)[index116_] = (d_829_v42_) | (d_829_v42_)
                d_836_v47_: _dafny.Map
                d_836_v47_ = _dafny.Map({p0: d_776_v2_})
                d_836_v47_ = (d_836_v47_).set(d_809_i6_, d_776_v2_)
                d_837_v48_: _dafny.Map
                d_837_v48_ = _dafny.Map({(self).f27: d_809_i6_})
                arr0_ = (d_798_v19_)[default__.safeIndex(843, (d_798_v19_).length(0))]
                index117_ = default__.safeIndex(511, ((d_798_v19_)[default__.safeIndex(843, (d_798_v19_).length(0))]).length(0))
                arr0_[index117_] = (-973) - (len(d_837_v48_))
                d_838_v49_: _dafny.Seq
                d_838_v49_ = _dafny.SeqWithoutIsStrInference([d_786_v8_, (d_786_v8_).set((self).f27, default__.abs(360))])
                d_839_v50_: _dafny.MultiSet
                d_839_v50_ = _dafny.MultiSet([(d_838_v49_)[default__.safeIndex((d_786_v8_).cardinality, len(d_838_v49_))], (d_786_v8_).intersection(d_786_v8_), _dafny.MultiSet([True, not(True)])])
                arr1_ = (d_798_v19_)[default__.safeIndex(843, (d_798_v19_).length(0))]
                index118_ = default__.safeIndex(511, ((d_798_v19_)[default__.safeIndex(843, (d_798_v19_).length(0))]).length(0))
                rhs100_ = d_774_v0_
                rhs101_ = (self).f27
                rhs102_ = ((0) - (p0)) - (len(d_810_v29_))
                rhs103_ = (0) - ((p0) - (394))
                rhs104_ = d_839_v50_
                lhs78_ = globalState
                lhs79_ = (d_798_v19_)[default__.safeIndex(843, (d_798_v19_).length(0))]
                lhs80_ = default__.safeIndex(511, ((d_798_v19_)[default__.safeIndex(843, (d_798_v19_).length(0))]).length(0))
                d_774_v0_ = rhs100_
                lhs78_.f21 = rhs101_
                lhs79_[lhs80_] = rhs102_
                r2 = rhs103_
                d_839_v50_ = rhs104_
            rhs105_ = p0
            rhs106_ = default__.safeDivisionInt((0) - (len(d_774_v0_)), d_809_i6_)
            lhs81_ = globalState
            lhs82_ = globalState
            lhs81_.f6 = rhs105_
            lhs82_.f6 = rhs106_
            index119_ = default__.safeIndex(211, (d_803_v24_).length(0))
            (d_803_v24_)[index119_] = (self).f27
            index120_ = default__.safeIndex(211, (d_803_v24_).length(0))
            def iife102_():
                coll50_ = _dafny.Set()
                compr_50_: int
                for compr_50_ in _dafny.IntegerRange(48, 85):
                    d_840_v51_: int = compr_50_
                    if ((48) <= (d_840_v51_)) and ((d_840_v51_) < (85)):
                        coll50_ = coll50_.union(_dafny.Set([(d_840_v51_) * (p0)]))
                return _dafny.Set(coll50_)
            (d_803_v24_)[index120_] = (len(iife102_()
            )) not in (_dafny.SeqWithoutIsStrInference([-577, 347]))
            d_797_v18_ = (d_798_v19_)[default__.safeIndex(843, (d_798_v19_).length(0))]
        hi7_ = default__.safeModuloInt(p0, p0)
        for d_841_i10_ in range(len(d_774_v0_), hi7_):
            (globalState).f5 = (self).f27
            index121_ = default__.safeIndex(98, (d_797_v18_).length(0))
            (d_797_v18_)[index121_] = d_841_i10_
            index122_ = default__.safeIndex(98, (d_797_v18_).length(0))
            (d_797_v18_)[index122_] = d_841_i10_
            d_842_v52_: _dafny.Map
            d_842_v52_ = _dafny.Map({(self).f27: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_843_i11_ in range(default__.abs(9))])})
            d_844_v53_: _dafny.Map
            d_844_v53_ = _dafny.Map({d_776_v2_: len(d_842_v52_)})
            d_844_v53_ = (d_844_v53_).set(_dafny.CodePoint('j'), ((d_786_v8_).cardinality) * (d_841_i10_))
            d_845_v54_: _dafny.Map
            d_845_v54_ = _dafny.Map({d_774_v0_: (self).f27})
            (globalState).f21 = ((d_845_v54_)[(d_774_v0_) + (d_774_v0_)] if ((d_774_v0_) + (d_774_v0_)) in (d_845_v54_) else (self).f27)
        r0 = default__.fm1((self).f27, d_796_v17_, (self).f27, globalState)
        r1 = (self).f27
        r2 = default__.safeModuloInt((p0) * (p0), len(d_787_v9_))
        r3 = default__.fm14((p0) - (p0), globalState)
        return r0, r1, r2, r3


class C2(T2, T0):
    def  __init__(self):
        self._f27: bool = False
        self._f34: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f27(self):
        return self._f27
    def ctor__(self, f34, f27):
        (self)._f34 = f34
        (self)._f27 = f27

    def fm4(self, p0, globalState):
        return (self).f34

    def fm5(self, globalState):
        return D0_DC2(649, len((_dafny.SeqWithoutIsStrInference([(self).f27, True, not((self).f34), not((self).f34)])) + (_dafny.SeqWithoutIsStrInference([not((self).f27), (self).f27]))))

    def fm2(self, p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cvhj"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nt")))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dvjfsk"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hbwjq"))))

    def fm3(self, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hwouleacx"))

    def fm9(self, p0, globalState):
        def iife103_():
            coll51_ = _dafny.Map()
            compr_51_: int
            for compr_51_ in _dafny.IntegerRange(381, 646):
                d_846_v0_: int = compr_51_
                if ((381) <= (d_846_v0_)) and ((d_846_v0_) < (646)):
                    coll51_[(d_846_v0_) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_847_i0_ in range(default__.abs(522))])))] = not(False)
            return _dafny.Map(coll51_)
        return (len(iife103_()
        )) * (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uhak"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gw")))))

    def m4(self, p0, p1, globalState):
        r0: bool = False
        r1: bool = False
        d_848_v0_: int
        d_848_v0_ = -398
        d_849_v1_: D0
        d_849_v1_ = D0_DC0(d_848_v0_)
        pat_let_tv21_ = d_848_v0_
        pat_let_tv22_ = d_848_v0_
        pat_let_tv23_ = p1
        def lambda27_(source17_):
            if source17_.is_DC1:
                d_850___mcc_h5_ = source17_.cf1
                d_851_cf1_ = d_850___mcc_h5_
                return (pat_let_tv21_) == (pat_let_tv22_)
            elif source17_.is_DC2:
                d_852___mcc_h6_ = source17_.cf2
                d_853___mcc_h7_ = source17_.cf3
                d_854_cf3_ = d_853___mcc_h7_
                d_855_cf2_ = d_852___mcc_h6_
                return ((self).f34) == (pat_let_tv23_)
            elif source17_.is_DC0:
                d_856___mcc_h8_ = source17_.cf0
                d_857_cf0_ = d_856___mcc_h8_
                return (self).f34
            elif True:
                d_858___mcc_h9_ = source17_.cf4
                d_859_cf4_ = d_858___mcc_h9_
                return False

        if lambda27_(d_849_v1_):
            (globalState).f25 = (self).f27
            d_860_v2_: _dafny.Seq
            d_860_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rsf"))
            d_861_v3_: str
            d_861_v3_ = _dafny.CodePoint('b')
            d_862_v4_: _dafny.Set
            d_862_v4_ = _dafny.Set({default__.fm1((self).f34, _dafny.SeqWithoutIsStrInference([p1]), (self).f27, globalState)})
            d_863_v5_: _dafny.Set
            d_863_v5_ = _dafny.Set({True, (self).f27, p1})
            d_864_v6_: _dafny.Map
            d_864_v6_ = _dafny.Map({d_848_v0_: not((self).f34)})
            d_865_v7_: _dafny.Array
            nw122_ = _dafny.Array(None, 23)
            nw122_[int(0)] = not((self).f34)
            nw122_[int(1)] = (self).f34
            nw122_[int(2)] = (self).f27
            nw122_[int(3)] = (self).f34
            nw122_[int(4)] = True
            nw122_[int(5)] = ((self).f34 if (self).f27 else p1)
            nw122_[int(6)] = (self).f27
            nw122_[int(7)] = (self).f27
            nw122_[int(8)] = default__.fm0((d_860_v2_).set(default__.safeIndex(len(d_860_v2_), len(d_860_v2_)), d_861_v3_), False, not(False), d_862_v4_, globalState)
            nw122_[int(9)] = p1
            nw122_[int(10)] = (d_861_v3_) not in (_dafny.SeqWithoutIsStrInference([d_861_v3_ for d_866_i0_ in range(default__.abs(707))]))
            nw122_[int(11)] = False
            nw122_[int(12)] = (self).f34
            nw122_[int(13)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bnrusmb"))) != (d_860_v2_)
            nw122_[int(14)] = (self).f27
            nw122_[int(15)] = (self).f27
            nw122_[int(16)] = (d_863_v5_).issubset(d_863_v5_)
            nw122_[int(17)] = (self).f34
            nw122_[int(18)] = (self).f27
            nw122_[int(19)] = ((d_864_v6_)[d_848_v0_] if (d_848_v0_) in (d_864_v6_) else p1)
            nw122_[int(20)] = (self).f27
            nw122_[int(21)] = (self).f27
            nw122_[int(22)] = not ((self).f34) or ((self).f34)
            d_865_v7_ = nw122_
            index123_ = default__.safeIndex(560, (d_865_v7_).length(0))
            (d_865_v7_)[index123_] = (self).f34
            d_867_v8_: _dafny.Seq
            d_867_v8_ = _dafny.SeqWithoutIsStrInference([d_865_v7_])
            index124_ = default__.safeIndex(560, (d_865_v7_).length(0))
            (d_865_v7_)[index124_] = ((0) - (len((d_867_v8_) + (_dafny.SeqWithoutIsStrInference([d_865_v7_, d_865_v7_]))))) >= (d_848_v0_)
            d_868_v9_: D0
            d_868_v9_ = D0_DC1((d_865_v7_)[default__.safeIndex(560, (d_865_v7_).length(0))])
            pat_let_tv24_ = d_865_v7_
            pat_let_tv25_ = d_865_v7_
            def iife104_(_pat_let26_0):
                def iife105_(d_869_dt__update__tmp_h0_):
                    def iife106_(_pat_let27_0):
                        def iife107_(d_870_dt__update_hcf1_h0_):
                            return D0_DC1(d_870_dt__update_hcf1_h0_)
                        return iife107_(_pat_let27_0)
                    return iife106_((pat_let_tv25_)[default__.safeIndex(560, (pat_let_tv24_).length(0))])
                return iife105_(_pat_let26_0)
            source18_ = iife104_(d_868_v9_)
            if source18_.is_DC1:
                d_871___mcc_h0_ = source18_.cf1
                d_872_cf1_ = d_871___mcc_h0_
                rhs107_ = d_861_v3_
                rhs108_ = d_848_v0_
                rhs109_ = 868
                lhs83_ = globalState
                lhs84_ = globalState
                d_861_v3_ = rhs107_
                lhs83_.f6 = rhs108_
                lhs84_.f6 = rhs109_
                d_873_v10_: _dafny.Array
                nw123_ = _dafny.Array(_dafny.Map({}), 17)
                d_873_v10_ = nw123_
                d_874_v11_: _dafny.Map
                d_874_v11_ = _dafny.Map({True: (self).f34})
                index125_ = default__.safeIndex(736, (d_873_v10_).length(0))
                (d_873_v10_)[index125_] = d_874_v11_
                d_875_v12_: _dafny.Map
                d_875_v12_ = _dafny.Map({(self).f34: d_862_v4_})
                d_876_v14_: _dafny.Seq
                d_876_v14_ = _dafny.SeqWithoutIsStrInference([d_848_v0_, (0) - (d_848_v0_)])
                d_877_v15_: _dafny.MultiSet
                d_877_v15_ = _dafny.MultiSet([default__.fm10((self).f34, globalState), d_876_v14_, d_876_v14_])
                d_878_v16_: _dafny.Array
                def lambda28_(d_879_v0_):
                    def lambda29_(d_880_i1_):
                        return default__.safeDivisionInt(d_880_i1_, d_879_v0_)

                    return lambda29_

                init13_ = lambda28_(d_848_v0_)
                nw124_ = _dafny.Array(None, 2)
                for i0_13_ in range(nw124_.length(0)):
                    nw124_[i0_13_] = init13_(i0_13_)
                d_878_v16_ = nw124_
                d_881_v17_: _dafny.Map
                d_881_v17_ = _dafny.Map({d_861_v3_: d_878_v16_})
                d_882_v18_: _dafny.Map
                d_882_v18_ = _dafny.Map({default__.fm11((d_865_v7_)[default__.safeIndex(560, (d_865_v7_).length(0))], not((d_865_v7_)[default__.safeIndex(560, (d_865_v7_).length(0))]), len(d_881_v17_), globalState): (d_864_v6_) == (d_864_v6_)})
                index126_ = default__.safeIndex(736, (d_873_v10_).length(0))
                def iife108_():
                    coll52_ = _dafny.Set()
                    compr_52_: int
                    for compr_52_ in _dafny.IntegerRange(998, 193):
                        d_883_v13_: int = compr_52_
                        if ((998) <= (d_883_v13_)) and ((d_883_v13_) < (193)):
                            coll52_ = coll52_.union(_dafny.Set([default__.safeModuloInt(d_883_v13_, d_848_v0_)]))
                    return _dafny.Set(coll52_)
                rhs110_ = (((d_875_v12_)[d_872_cf1_] if (d_872_cf1_) in (d_875_v12_) else d_862_v4_)) - ((d_862_v4_) - (iife108_()
                ))
                rhs111_ = (_dafny.MultiSet([d_876_v14_])).issubset(((d_877_v15_).set(_dafny.SeqWithoutIsStrInference([210]), default__.abs(d_848_v0_)) if (self).f27 else d_877_v15_))
                rhs112_ = ((d_882_v18_)[d_849_v1_] if (d_849_v1_) in (d_882_v18_) else (d_865_v7_)[default__.safeIndex(560, (d_865_v7_).length(0))])
                rhs113_ = ((d_874_v11_) | (d_874_v11_)) | (d_874_v11_)
                lhs85_ = globalState
                lhs86_ = d_873_v10_
                lhs87_ = default__.safeIndex(736, (d_873_v10_).length(0))
                d_862_v4_ = rhs110_
                r0 = rhs111_
                lhs85_.f12 = rhs112_
                lhs86_[lhs87_] = rhs113_
                d_884_v19_: _dafny.Array
                nw125_ = _dafny.Array(None, 23)
                nw125_[int(0)] = d_876_v14_
                nw125_[int(1)] = d_876_v14_
                nw125_[int(2)] = default__.fm10(p1, globalState)
                nw125_[int(3)] = (d_876_v14_) + (d_876_v14_)
                nw125_[int(4)] = (d_876_v14_ if d_872_cf1_ else d_876_v14_)
                nw125_[int(5)] = d_876_v14_
                nw125_[int(6)] = (d_876_v14_) + (d_876_v14_)
                nw125_[int(7)] = (d_876_v14_) + (d_876_v14_)
                nw125_[int(8)] = d_876_v14_
                nw125_[int(9)] = (_dafny.SeqWithoutIsStrInference([d_848_v0_, d_848_v0_])) + (_dafny.SeqWithoutIsStrInference([d_848_v0_, 687]))
                nw125_[int(10)] = d_876_v14_
                nw125_[int(11)] = d_876_v14_
                nw125_[int(12)] = (d_876_v14_) + (d_876_v14_)
                nw125_[int(13)] = _dafny.SeqWithoutIsStrInference([d_848_v0_, d_848_v0_])
                nw125_[int(14)] = (_dafny.SeqWithoutIsStrInference([d_848_v0_ for d_885_i2_ in range(default__.abs(267))])) + (_dafny.SeqWithoutIsStrInference([-618 for d_886_i3_ in range(default__.abs(-702))]))
                nw125_[int(15)] = (d_876_v14_) + (_dafny.SeqWithoutIsStrInference([d_848_v0_, (0) - (len(d_860_v2_))]))
                nw125_[int(16)] = d_876_v14_
                nw125_[int(17)] = (_dafny.SeqWithoutIsStrInference([d_848_v0_])) + (d_876_v14_)
                nw125_[int(18)] = d_876_v14_
                nw125_[int(19)] = (d_876_v14_) + (d_876_v14_)
                nw125_[int(20)] = d_876_v14_
                nw125_[int(21)] = d_876_v14_
                nw125_[int(22)] = d_876_v14_
                d_884_v19_ = nw125_
                rhs114_ = (d_884_v19_ if (d_848_v0_) >= (d_848_v0_) else d_884_v19_)
                rhs115_ = d_872_cf1_
                lhs88_ = globalState
                d_884_v19_ = rhs114_
                lhs88_.f12 = rhs115_
                d_887_v20_: _dafny.Array
                def lambda30_(d_888_i4_):
                    return _dafny.SeqWithoutIsStrInference([(self).f27, True])

                init14_ = lambda30_
                nw126_ = _dafny.Array(None, 7)
                for i0_14_ in range(nw126_.length(0)):
                    nw126_[i0_14_] = init14_(i0_14_)
                d_887_v20_ = nw126_
                d_887_v20_ = d_887_v20_
            elif source18_.is_DC2:
                d_889___mcc_h1_ = source18_.cf2
                d_890___mcc_h2_ = source18_.cf3
                d_891_cf3_ = d_890___mcc_h2_
                d_892_cf2_ = d_889___mcc_h1_
                d_893_v21_: _dafny.Seq
                d_893_v21_ = _dafny.SeqWithoutIsStrInference([True, (d_865_v7_)[default__.safeIndex(560, (d_865_v7_).length(0))]])
                d_894_v22_: _dafny.Array
                def lambda31_(d_895_v21_):
                    def lambda32_(d_896_i5_):
                        return d_895_v21_

                    return lambda32_

                init15_ = lambda31_(d_893_v21_)
                nw127_ = _dafny.Array(None, 20)
                for i0_15_ in range(nw127_.length(0)):
                    nw127_[i0_15_] = init15_(i0_15_)
                d_894_v22_ = nw127_
                d_897_v23_: _dafny.Map
                d_897_v23_ = _dafny.Map({d_893_v21_: d_894_v22_})
                d_897_v23_ = (d_897_v23_).set((d_893_v21_ if (self).f27 else d_893_v21_), (d_894_v22_ if (self).f27 else d_894_v22_))
                d_898_v24_: _dafny.MultiSet
                d_898_v24_ = _dafny.MultiSet([d_891_cf3_])
                pat_let_tv26_ = d_848_v0_
                d_899_v25_: _dafny.Map
                def iife109_(_pat_let28_0):
                    def iife110_(d_900_dt__update__tmp_h1_):
                        def iife111_(_pat_let29_0):
                            def iife112_(d_901_dt__update_hcf2_h0_):
                                return D0_DC2(d_901_dt__update_hcf2_h0_, (d_900_dt__update__tmp_h1_).cf3)
                            return iife112_(_pat_let29_0)
                        return iife111_(pat_let_tv26_)
                    return iife110_(_pat_let28_0)
                d_899_v25_ = _dafny.Map({iife109_(D0_DC2(d_848_v0_, (0) - (d_891_cf3_))): (self).f27})
                d_902_v26_: _dafny.Seq
                d_902_v26_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))])
                rhs116_ = d_892_cf2_
                rhs117_ = ((0) - (len(d_860_v2_))) not in (d_898_v24_)
                rhs118_ = ((d_899_v25_)[D0_DC2(len(d_902_v26_), d_891_cf3_)] if (D0_DC2(len(d_902_v26_), d_891_cf3_)) in (d_899_v25_) else (d_865_v7_)[default__.safeIndex(560, (d_865_v7_).length(0))])
                lhs89_ = globalState
                lhs90_ = globalState
                d_891_cf3_ = rhs116_
                lhs89_.f21 = rhs117_
                lhs90_.f25 = rhs118_
                d_860_v2_ = (d_860_v2_).set(default__.safeIndex(d_892_cf2_, len(d_860_v2_)), (d_860_v2_)[default__.safeIndex(len(d_860_v2_), len(d_860_v2_))])
                d_903_v27_: D0
                d_903_v27_ = D0_DC2(d_892_cf2_, d_891_cf3_)
                d_904_v28_: D0
                d_904_v28_ = D0_DC0((d_903_v27_).cf3)
                d_905_v29_: D0
                d_905_v29_ = D0_DC3(d_904_v28_)
                d_906_v30_: D0
                d_906_v30_ = D0_DC3(d_905_v29_)
                d_906_v30_ = D0_DC3(D0_DC2(len(d_860_v2_), (self).fm9(d_892_cf2_, globalState)))
            elif source18_.is_DC0:
                d_907___mcc_h3_ = source18_.cf0
                d_908_cf0_ = d_907___mcc_h3_
                (globalState).f25 = (d_848_v0_) == (d_848_v0_)
                d_909_v31_: _dafny.Map
                d_909_v31_ = _dafny.Map({True: d_908_cf0_})
                d_909_v31_ = (d_909_v31_).set(not (False) or (default__.fm0(d_860_v2_, p1, (d_865_v7_)[default__.safeIndex(560, (d_865_v7_).length(0))], d_862_v4_, globalState)), default__.safeDivisionInt(d_908_cf0_, d_908_cf0_))
                d_910_v33_: C1
                nw128_ = C1()
                def iife113_():
                    coll53_ = _dafny.Set()
                    compr_53_: int
                    for compr_53_ in (d_864_v6_).keys.Elements:
                        d_912_v32_: int = compr_53_
                        if (d_912_v32_) in (d_864_v6_):
                            coll53_ = coll53_.union(_dafny.Set([(d_912_v32_) + (844)]))
                    return _dafny.Set(coll53_)
                nw128_.ctor__(default__.fm0(_dafny.SeqWithoutIsStrInference([d_861_v3_ for d_911_i6_ in range(default__.abs(616))]), p1, (self).f34, iife113_()
                , globalState))
                d_910_v33_ = nw128_
                r0 = not((d_865_v7_)[default__.safeIndex(560, (d_865_v7_).length(0))])
            elif True:
                d_913___mcc_h4_ = source18_.cf4
                d_914_cf4_ = d_913___mcc_h4_
                d_915_v34_: _dafny.MultiSet
                d_915_v34_ = _dafny.MultiSet([d_865_v7_, d_865_v7_])
                d_916_v35_: _dafny.MultiSet
                d_916_v35_ = _dafny.MultiSet([len(_dafny.Set({(self).f34, (self).f27, (self).f27, (self).f27})), d_848_v0_, d_848_v0_, d_848_v0_])
                index127_ = default__.safeIndex(560, (d_865_v7_).length(0))
                rhs119_ = (d_915_v34_).ispropersubset(d_915_v34_)
                rhs120_ = (0) - (d_848_v0_)
                rhs121_ = ((0) - ((d_916_v35_).cardinality)) == ((d_848_v0_) + (((d_916_v35_)[d_848_v0_] if (d_848_v0_) in (d_916_v35_) else d_848_v0_)))
                rhs122_ = default__.fm0(d_860_v2_, (self).f27, True, d_862_v4_, globalState)
                lhs91_ = globalState
                lhs92_ = globalState
                lhs93_ = d_865_v7_
                lhs94_ = default__.safeIndex(560, (d_865_v7_).length(0))
                lhs95_ = globalState
                lhs91_.f14 = rhs119_
                lhs92_.f2 = rhs120_
                lhs93_[lhs94_] = rhs121_
                lhs95_.f25 = rhs122_
                d_848_v0_ = d_848_v0_
                d_917_v36_: D11
                d_917_v36_ = D11_DC25()
                d_918_v37_: _dafny.Map
                d_918_v37_ = _dafny.Map({d_917_v36_: (d_862_v4_) | (d_862_v4_)})
                d_918_v37_ = (d_918_v37_).set(d_917_v36_, d_862_v4_)
                (globalState).f14 = False
            d_919_v38_: _dafny.Seq
            d_919_v38_ = _dafny.SeqWithoutIsStrInference([True, not((d_865_v7_)[default__.safeIndex(560, (d_865_v7_).length(0))])])
            d_920_v39_: _dafny.Seq
            d_920_v39_ = _dafny.SeqWithoutIsStrInference([d_919_v38_, d_919_v38_])
            d_921_v40_: _dafny.Array
            def lambda33_(d_922_v0_):
                def lambda34_(d_923_i7_):
                    return (d_923_i7_) * (d_922_v0_)

                return lambda34_

            init16_ = lambda33_(d_848_v0_)
            nw129_ = _dafny.Array(None, 28)
            for i0_16_ in range(nw129_.length(0)):
                nw129_[i0_16_] = init16_(i0_16_)
            d_921_v40_ = nw129_
            d_924_v41_: _dafny.Array
            nw130_ = _dafny.Array(None, 10)
            nw130_[int(0)] = d_848_v0_
            nw130_[int(1)] = d_848_v0_
            nw130_[int(2)] = d_848_v0_
            nw130_[int(3)] = d_848_v0_
            nw130_[int(4)] = d_848_v0_
            nw130_[int(5)] = d_848_v0_
            nw130_[int(6)] = default__.fm1((self).f34, d_919_v38_, p1, globalState)
            nw130_[int(7)] = (721) - (len(d_920_v39_))
            nw130_[int(8)] = (_dafny.MultiSet([d_921_v40_])).cardinality
            nw130_[int(9)] = d_848_v0_
            d_924_v41_ = nw130_
            index128_ = default__.safeIndex(930, (d_924_v41_).length(0))
            (d_924_v41_)[index128_] = d_848_v0_
            index129_ = default__.safeIndex(930, (d_924_v41_).length(0))
            (d_924_v41_)[index129_] = ((-985) - (d_848_v0_)) - (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "laqd")) if (self).f27 else d_860_v2_)))
            r0 = (self).f34
        elif True:
            (globalState).f14 = (self).f27
            d_925_v42_: _dafny.Array
            nw131_ = _dafny.Array(False, 20)
            d_925_v42_ = nw131_
            d_925_v42_ = d_925_v42_
            d_926_v43_: C0
            nw132_ = C0()
            nw132_.ctor__()
            d_926_v43_ = nw132_
            d_927_v44_: str
            d_927_v44_ = _dafny.CodePoint('q')
            d_928_v45_: _dafny.MultiSet
            d_928_v45_ = _dafny.MultiSet([d_927_v44_, d_927_v44_])
            d_929_v46_: _dafny.Map
            d_929_v46_ = _dafny.Map({(0) - (d_848_v0_): (self).f27})
            d_930_v47_: _dafny.Seq
            d_930_v47_ = _dafny.SeqWithoutIsStrInference([p1, (self).f27, ((d_929_v46_)[(0) - (d_848_v0_)] if ((0) - (d_848_v0_)) in (d_929_v46_) else (self).f27)])
            rhs123_ = p1
            rhs124_ = (d_930_v47_)[default__.safeIndex(d_848_v0_, len(d_930_v47_))]
            rhs125_ = d_928_v45_
            lhs96_ = globalState
            lhs97_ = globalState
            lhs96_.f21 = rhs123_
            lhs97_.f21 = rhs124_
            d_928_v45_ = rhs125_
            d_931_v48_: C0
            nw133_ = C0()
            nw133_.ctor__()
            d_931_v48_ = nw133_
        (globalState).f5 = not(p1)
        d_932_v49_: _dafny.Seq
        d_932_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kklmy"))
        d_933_v50_: _dafny.Seq
        d_933_v50_ = _dafny.SeqWithoutIsStrInference([len(default__.fm19(d_932_v49_, globalState)), d_848_v0_, d_848_v0_])
        d_934_v51_: D2
        d_934_v51_ = D2_DC6(d_848_v0_, d_933_v50_, (self).f34, d_933_v50_, p1)
        d_935_v52_: _dafny.Map
        d_935_v52_ = _dafny.Map({d_932_v49_: (d_934_v51_).cf11})
        d_935_v52_ = (d_935_v52_).set((d_932_v49_ if (self).f34 else d_932_v49_), not(p1))
        (globalState).f6 = d_848_v0_
        hi8_ = d_848_v0_
        for d_936_i8_ in range(d_848_v0_, hi8_):
            d_937_v53_: _dafny.Seq
            d_937_v53_ = _dafny.SeqWithoutIsStrInference([(self).f34, p1])
            d_938_v54_: _dafny.Seq
            d_938_v54_ = _dafny.SeqWithoutIsStrInference([d_932_v49_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))])
            d_939_v55_: _dafny.MultiSet
            d_939_v55_ = _dafny.MultiSet([d_936_i8_])
            d_940_v56_: _dafny.Array
            nw134_ = _dafny.Array(None, 3)
            nw134_[int(0)] = d_848_v0_
            nw134_[int(1)] = default__.fm1(p1, d_937_v53_, True, globalState)
            nw134_[int(2)] = len((d_938_v54_)[default__.safeIndex((0) - (((d_939_v55_)[-946] if (-946) in (d_939_v55_) else len(d_932_v49_))), len(d_938_v54_))])
            d_940_v56_ = nw134_
            d_941_v57_: _dafny.Map
            d_941_v57_ = _dafny.Map({d_940_v56_: ((d_933_v50_)[default__.safeIndex(579, len(d_933_v50_))]) - (d_848_v0_)})
            d_941_v57_ = (d_941_v57_).set(d_940_v56_, d_936_i8_)
            if (d_936_i8_) < (d_936_i8_):
                (globalState).f12 = (d_848_v0_) >= ((d_933_v50_)[default__.safeIndex(d_848_v0_, len(d_933_v50_))])
                d_942_v58_: _dafny.Array
                def lambda35_(d_943_v53_):
                    def lambda36_(d_944_i9_):
                        return d_943_v53_

                    return lambda36_

                init17_ = lambda35_(d_937_v53_)
                nw135_ = _dafny.Array(None, 3)
                for i0_17_ in range(nw135_.length(0)):
                    nw135_[i0_17_] = init17_(i0_17_)
                d_942_v58_ = nw135_
                index130_ = default__.safeIndex(611, (d_942_v58_).length(0))
                (d_942_v58_)[index130_] = d_937_v53_
                index131_ = default__.safeIndex(611, (d_942_v58_).length(0))
                (d_942_v58_)[index131_] = d_937_v53_
                d_932_v49_ = d_932_v49_
                d_945_v59_: _dafny.Array
                nw136_ = _dafny.Array(False, 11)
                d_945_v59_ = nw136_
                index132_ = default__.safeIndex(514, (d_945_v59_).length(0))
                (d_945_v59_)[index132_] = not(p1)
                index133_ = default__.safeIndex(514, (d_945_v59_).length(0))
                (d_945_v59_)[index133_] = (self).f34
                d_940_v56_ = d_940_v56_
            elif True:
                d_946_v60_: C1
                nw137_ = C1()
                nw137_.ctor__((self).f34)
                d_946_v60_ = nw137_
                rhs126_ = d_848_v0_
                rhs127_ = d_940_v56_
                lhs98_ = globalState
                lhs98_.f2 = rhs126_
                d_940_v56_ = rhs127_
                (globalState).f6 = d_936_i8_
                d_947_v61_: _dafny.Map
                d_947_v61_ = _dafny.Map({False: d_848_v0_})
                (globalState).f13 = (d_947_v61_) | (d_947_v61_)
                (globalState).f12 = (self).f27
            d_948_v62_: D11
            d_948_v62_ = D11_DC25()
            d_948_v62_ = d_948_v62_
            d_949_v63_: C0
            nw138_ = C0()
            nw138_.ctor__()
            d_949_v63_ = nw138_
        d_950_v64_: D0
        d_950_v64_ = D0_DC1(((self).f27 if (self).f27 else p1))
        d_951_v65_: _dafny.Map
        d_951_v65_ = _dafny.Map({d_848_v0_: (self).f27})
        d_952_v66_: D7
        d_952_v66_ = D7_DC15(d_951_v65_)
        pat_let_tv27_ = d_950_v64_
        pat_let_tv28_ = d_950_v64_
        pat_let_tv29_ = d_950_v64_
        pat_let_tv30_ = d_950_v64_
        def lambda37_(source19_):
            if source19_.is_DC16:
                return pat_let_tv27_
            elif source19_.is_DC17:
                d_953___mcc_h10_ = source19_.cf25
                d_954___mcc_h11_ = source19_.cf26
                d_955___mcc_h12_ = source19_.cf27
                d_956_cf27_ = d_955___mcc_h12_
                d_957_cf26_ = d_954___mcc_h11_
                d_958_cf25_ = d_953___mcc_h10_
                return pat_let_tv28_
            elif source19_.is_DC15:
                d_959___mcc_h13_ = source19_.cf24
                d_960_cf24_ = d_959___mcc_h13_
                return pat_let_tv29_
            elif True:
                d_961___mcc_h14_ = source19_.cf28
                d_962_cf28_ = d_961___mcc_h14_
                return pat_let_tv30_

        d_950_v64_ = lambda37_(d_952_v66_)
        r0 = (self).fm4(d_950_v64_, globalState)
        r1 = (self).f27
        return r0, r1

    def m1(self, globalState):
        r0: int = int(0)
        d_963_v0_: _dafny.Array
        def lambda38_(d_964_i0_):
            return not (False) or ((self).f34)

        init18_ = lambda38_
        nw139_ = _dafny.Array(None, 20)
        for i0_18_ in range(nw139_.length(0)):
            nw139_[i0_18_] = init18_(i0_18_)
        d_963_v0_ = nw139_
        d_965_v1_: _dafny.Map
        d_965_v1_ = _dafny.Map({(self).f34: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "of"))})
        d_966_v2_: _dafny.Seq
        d_966_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hqkivnnwe"))
        d_967_v3_: _dafny.Map
        d_967_v3_ = _dafny.Map({((d_965_v1_)[(self).f34] if ((self).f34) in (d_965_v1_) else d_966_v2_): (d_963_v0_ if True else d_963_v0_)})
        d_963_v0_ = ((d_967_v3_)[(d_966_v2_) + (d_966_v2_)] if ((d_966_v2_) + (d_966_v2_)) in (d_967_v3_) else (d_963_v0_ if (self).f27 else d_963_v0_))
        d_968_v4_: int
        d_968_v4_ = 299
        d_969_v5_: D13
        d_969_v5_ = D13_DC28(d_968_v4_)
        def lambda39_(source21_):
            if source21_.is_DC28:
                d_970___mcc_h6_ = source21_.cf38
                d_971_cf38_ = d_970___mcc_h6_
                return D15_DC33()
            elif source21_.is_DC29:
                d_972___mcc_h7_ = source21_.cf39
                d_973___mcc_h8_ = source21_.cf40
                d_974___mcc_h9_ = source21_.cf41
                d_975_cf41_ = d_974___mcc_h9_
                d_976_cf40_ = d_973___mcc_h8_
                d_977_cf39_ = d_972___mcc_h7_
                if (self).f34:
                    return D15_DC33()
                elif True:
                    return D15_DC33()
            elif True:
                d_978___mcc_h10_ = source21_.cf37
                d_979_cf37_ = d_978___mcc_h10_
                return D15_DC33()

        source20_ = lambda39_(d_969_v5_)
        if source20_.is_DC33:
            (globalState).f2 = default__.safeDivisionInt(d_968_v4_, d_968_v4_)
            d_980_v6_: C0
            nw140_ = C0()
            nw140_.ctor__()
            d_980_v6_ = nw140_
            d_980_v6_ = d_980_v6_
            d_981_v7_: _dafny.Array
            nw141_ = _dafny.Array(_dafny.Seq({}), 25)
            d_981_v7_ = nw141_
            d_982_v8_: _dafny.Seq
            d_982_v8_ = _dafny.SeqWithoutIsStrInference([d_980_v6_])
            index134_ = default__.safeIndex(534, (d_981_v7_).length(0))
            (d_981_v7_)[index134_] = d_982_v8_
            index135_ = default__.safeIndex(534, (d_981_v7_).length(0))
            (d_981_v7_)[index135_] = d_982_v8_
            (globalState).f2 = d_968_v4_
        elif source20_.is_DC34:
            d_983___mcc_h0_ = source20_.cf45
            d_984___mcc_h1_ = source20_.cf46
            d_985___mcc_h2_ = source20_.cf47
            d_986___mcc_h3_ = source20_.cf48
            d_987___mcc_h4_ = source20_.cf49
            d_988_cf49_ = d_987___mcc_h4_
            d_989_cf48_ = d_986___mcc_h3_
            d_990_cf47_ = d_985___mcc_h2_
            d_991_cf46_ = d_984___mcc_h1_
            d_992_cf45_ = d_983___mcc_h0_
            if d_991_cf46_:
                d_993_v9_: _dafny.Seq
                d_993_v9_ = _dafny.SeqWithoutIsStrInference([d_992_cf45_])
                d_994_v10_: _dafny.Map
                d_994_v10_ = _dafny.Map({d_991_cf46_: d_993_v9_})
                d_993_v9_ = ((_dafny.SeqWithoutIsStrInference([len(d_994_v10_)])) + (d_993_v9_)).set(default__.safeIndex(d_989_cf48_, len((_dafny.SeqWithoutIsStrInference([len(d_994_v10_)])) + (d_993_v9_))), (d_992_cf45_) + (d_992_cf45_))
                d_995_v11_: _dafny.MultiSet
                d_995_v11_ = _dafny.MultiSet([(self).f27, d_988_cf49_, d_988_cf49_])
                d_995_v11_ = d_995_v11_
                d_996_v12_: D6
                d_996_v12_ = D6_DC14(False, d_968_v4_)
                d_989_cf48_ = default__.safeDivisionInt(default__.safeModuloInt((d_996_v12_).cf23, d_989_cf48_), d_990_cf47_)
                d_997_v13_: _dafny.Seq
                d_997_v13_ = _dafny.SeqWithoutIsStrInference([(d_966_v2_) + (d_966_v2_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfabhwpmd")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_998_i1_ in range(default__.abs(887))])])
                d_997_v13_ = (_dafny.SeqWithoutIsStrInference([d_966_v2_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_999_i2_ in range(default__.abs(871))])])) + (d_997_v13_)
                (globalState).f14 = (d_988_cf49_) == (d_991_cf46_)
            elif True:
                d_1000_v14_: _dafny.Array
                nw142_ = _dafny.Array(_dafny.Seq({}), 26)
                d_1000_v14_ = nw142_
                index136_ = default__.safeIndex(52, (d_1000_v14_).length(0))
                (d_1000_v14_)[index136_] = default__.fm20(globalState)
                d_1001_v15_: _dafny.Map
                d_1001_v15_ = _dafny.Map({d_991_cf46_: len(d_966_v2_)})
                d_1002_v16_: _dafny.Seq
                d_1002_v16_ = _dafny.SeqWithoutIsStrInference([d_1001_v15_])
                d_1003_v17_: _dafny.Map
                d_1003_v17_ = _dafny.Map({d_991_cf46_: d_1001_v15_})
                d_1004_v18_: _dafny.Map
                d_1004_v18_ = _dafny.Map({(self).f27: ((d_1002_v16_) + (d_1002_v16_)).set(default__.safeIndex(418, len((d_1002_v16_) + (d_1002_v16_))), ((d_1003_v17_)[d_991_cf46_] if (d_991_cf46_) in (d_1003_v17_) else d_1001_v15_))})
                index137_ = default__.safeIndex(52, (d_1000_v14_).length(0))
                (d_1000_v14_)[index137_] = ((d_1004_v18_)[(self).f27] if ((self).f27) in (d_1004_v18_) else d_1002_v16_)
                d_1005_v19_: _dafny.Array
                def lambda40_(d_1006_v2_, d_1007_cf45_):
                    def lambda41_(d_1008_i3_):
                        return ((d_1006_v2_) + (d_1006_v2_)).set(default__.safeIndex(d_1007_cf45_, len((d_1006_v2_) + (d_1006_v2_))), _dafny.CodePoint('e'))

                    return lambda41_

                init19_ = lambda40_(d_966_v2_, d_992_cf45_)
                nw143_ = _dafny.Array(None, 2)
                for i0_19_ in range(nw143_.length(0)):
                    nw143_[i0_19_] = init19_(i0_19_)
                d_1005_v19_ = nw143_
                index138_ = default__.safeIndex(468, (d_1005_v19_).length(0))
                (d_1005_v19_)[index138_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uo"))) + (d_966_v2_)
                index139_ = default__.safeIndex(468, (d_1005_v19_).length(0))
                (d_1005_v19_)[index139_] = d_966_v2_
                d_1009_v20_: _dafny.Seq
                d_1009_v20_ = _dafny.SeqWithoutIsStrInference([(self).f27, (self).f27])
                d_1010_v21_: _dafny.Set
                d_1010_v21_ = _dafny.Set({d_988_cf49_})
                d_1011_v22_: _dafny.Set
                d_1011_v22_ = _dafny.Set({d_992_cf45_, d_990_cf47_, d_968_v4_, 423})
                d_1012_v23_: _dafny.Seq
                d_1012_v23_ = _dafny.SeqWithoutIsStrInference([default__.fm1(d_991_cf46_, d_1009_v20_, (self).f34, globalState), len(d_1010_v21_), d_990_cf47_, d_989_cf48_, len(d_1011_v22_)])
                rhs128_ = (d_1012_v23_)[default__.safeIndex(default__.safeDivisionInt(225, d_989_cf48_), len(d_1012_v23_))]
                rhs129_ = (d_968_v4_) * (d_989_cf48_)
                rhs130_ = default__.safeModuloInt(d_968_v4_, (0) - ((D15_DC34(-444, (D2_DC6(-644, d_1012_v23_, True, d_1012_v23_, d_991_cf46_)).cf9, d_992_cf45_, (0) - (d_990_cf47_), d_991_cf46_)).cf45))
                rhs131_ = default__.fm0((d_1005_v19_)[default__.safeIndex(468, (d_1005_v19_).length(0))], (self).f27, d_988_cf49_, d_1011_v22_, globalState)
                rhs132_ = d_989_cf48_
                lhs99_ = globalState
                lhs100_ = globalState
                d_989_cf48_ = rhs128_
                lhs99_.f6 = rhs129_
                d_990_cf47_ = rhs130_
                lhs100_.f21 = rhs131_
                d_989_cf48_ = rhs132_
                d_1013_v24_: str
                d_1013_v24_ = _dafny.CodePoint('i')
                d_1014_v25_: _dafny.Map
                d_1014_v25_ = _dafny.Map({143: d_1013_v24_})
                d_1015_v26_: _dafny.Seq
                d_1015_v26_ = _dafny.SeqWithoutIsStrInference([d_966_v2_, (self).fm2((0) - (d_989_cf48_), _dafny.Map({d_1013_v24_: d_1014_v25_}), globalState)])
                d_1001_v15_ = (d_1001_v15_).set(default__.fm0((d_1015_v26_)[default__.safeIndex(d_990_cf47_, len(d_1015_v26_))], (self).f27, (self).f27, d_1011_v22_, globalState), len((d_1012_v23_) + (d_1012_v23_)))
                d_1016_v27_: _dafny.Array
                def lambda42_(d_1017_v4_):
                    def lambda43_(d_1018_i4_):
                        return (d_1018_i4_) + (d_1017_v4_)

                    return lambda43_

                init20_ = lambda42_(d_968_v4_)
                nw144_ = _dafny.Array(None, 28)
                for i0_20_ in range(nw144_.length(0)):
                    nw144_[i0_20_] = init20_(i0_20_)
                d_1016_v27_ = nw144_
                d_1019_v28_: _dafny.MultiSet
                d_1019_v28_ = _dafny.MultiSet([d_1016_v27_, (d_1016_v27_ if (self).f27 else d_1016_v27_)])
                d_1019_v28_ = d_1019_v28_
            index140_ = default__.safeIndex(165, (d_963_v0_).length(0))
            (d_963_v0_)[index140_] = (self).f34
            index141_ = default__.safeIndex(165, (d_963_v0_).length(0))
            (d_963_v0_)[index141_] = not(d_991_cf46_)
            if (self).f27:
                d_1020_v29_: str
                d_1020_v29_ = _dafny.CodePoint('a')
                d_1021_v30_: _dafny.Array
                nw145_ = _dafny.Array(_dafny.MultiSet({}), 26)
                d_1021_v30_ = nw145_
                d_1022_v31_: _dafny.Array
                nw146_ = _dafny.Array(None, 6)
                nw146_[int(0)] = (d_963_v0_)[default__.safeIndex(165, (d_963_v0_).length(0))]
                nw146_[int(1)] = (self).f34
                nw146_[int(2)] = d_988_cf49_
                nw146_[int(3)] = True
                nw146_[int(4)] = d_991_cf46_
                nw146_[int(5)] = d_988_cf49_
                d_1022_v31_ = nw146_
                d_1023_v32_: _dafny.MultiSet
                d_1023_v32_ = _dafny.MultiSet([d_1022_v31_])
                index142_ = default__.safeIndex(670, (d_1021_v30_).length(0))
                (d_1021_v30_)[index142_] = d_1023_v32_
                index143_ = default__.safeIndex(165, (d_963_v0_).length(0))
                index144_ = default__.safeIndex(670, (d_1021_v30_).length(0))
                rhs133_ = (d_966_v2_) < (d_966_v2_)
                rhs134_ = d_1020_v29_
                rhs135_ = not((d_963_v0_)[default__.safeIndex(165, (d_963_v0_).length(0))])
                rhs136_ = (d_966_v2_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dif"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jvejx"))))
                rhs137_ = ((_dafny.MultiSet([d_1022_v31_, d_1022_v31_])).set(d_1022_v31_, default__.abs(d_989_cf48_))).set(d_1022_v31_, default__.abs(d_992_cf45_))
                lhs101_ = globalState
                lhs102_ = d_963_v0_
                lhs103_ = default__.safeIndex(165, (d_963_v0_).length(0))
                lhs104_ = d_1021_v30_
                lhs105_ = default__.safeIndex(670, (d_1021_v30_).length(0))
                lhs101_.f25 = rhs133_
                d_1020_v29_ = rhs134_
                lhs102_[lhs103_] = rhs135_
                d_966_v2_ = rhs136_
                lhs104_[lhs105_] = rhs137_
                (globalState).f6 = (self).fm9((d_990_cf47_) * (d_990_cf47_), globalState)
                d_1024_v33_: _dafny.Set
                d_1024_v33_ = _dafny.Set({d_990_cf47_, d_990_cf47_})
                d_1024_v33_ = (d_1024_v33_ if (d_966_v2_) == (d_966_v2_) else d_1024_v33_)
                (globalState).f6 = (949) - (d_992_cf45_)
                d_1025_v34_: D0
                d_1025_v34_ = D0_DC1(d_988_cf49_)
                (globalState).f5 = (self).fm4(d_1025_v34_, globalState)
            elif True:
                d_1026_v35_: _dafny.Array
                nw147_ = _dafny.Array(_dafny.Seq({}), 12)
                d_1026_v35_ = nw147_
                d_1027_v36_: _dafny.Seq
                d_1027_v36_ = _dafny.SeqWithoutIsStrInference([(self).f34, d_988_cf49_, (self).f27, d_991_cf46_])
                index145_ = default__.safeIndex(694, (d_1026_v35_).length(0))
                (d_1026_v35_)[index145_] = d_1027_v36_
                d_1028_v37_: D0
                d_1028_v37_ = D0_DC1((self).f34)
                d_1029_v38_: _dafny.Map
                d_1029_v38_ = _dafny.Map({(self).fm4(d_1028_v37_, globalState): d_989_cf48_})
                index146_ = default__.safeIndex(694, (d_1026_v35_).length(0))
                rhs138_ = not (d_988_cf49_) or (d_991_cf46_)
                rhs139_ = d_1029_v38_
                rhs140_ = (d_1027_v36_) + (d_1027_v36_)
                lhs106_ = globalState
                lhs107_ = globalState
                lhs108_ = d_1026_v35_
                lhs109_ = default__.safeIndex(694, (d_1026_v35_).length(0))
                lhs106_.f14 = rhs138_
                lhs107_.f13 = rhs139_
                lhs108_[lhs109_] = rhs140_
                (globalState).f2 = ((d_990_cf47_) + (d_989_cf48_)) * ((_dafny.MultiSet(d_1027_v36_)).cardinality)
                d_1030_v39_: _dafny.Array
                def lambda44_(d_1031_v2_):
                    def lambda45_(d_1032_i5_):
                        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bcmrhddxu"))) + (d_1031_v2_)

                    return lambda45_

                init21_ = lambda44_(d_966_v2_)
                nw148_ = _dafny.Array(None, 19)
                for i0_21_ in range(nw148_.length(0)):
                    nw148_[i0_21_] = init21_(i0_21_)
                d_1030_v39_ = nw148_
                d_1033_v40_: str
                d_1033_v40_ = _dafny.CodePoint('b')
                d_1034_v41_: _dafny.Map
                d_1034_v41_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_990_cf47_ for d_1035_i6_ in range(default__.abs(-679))])): d_1033_v40_})
                d_1036_v42_: _dafny.Map
                d_1036_v42_ = _dafny.Map({_dafny.CodePoint('y'): d_1034_v41_})
                index147_ = default__.safeIndex(368, (d_1030_v39_).length(0))
                (d_1030_v39_)[index147_] = (d_966_v2_) + ((self).fm2(d_990_cf47_, d_1036_v42_, globalState))
                index148_ = default__.safeIndex(368, (d_1030_v39_).length(0))
                (d_1030_v39_)[index148_] = d_966_v2_
                d_1037_v43_: C1
                nw149_ = C1()
                nw149_.ctor__(not((self).fm4(d_1028_v37_, globalState)))
                d_1037_v43_ = nw149_
                d_1038_v45_: _dafny.Set
                d_1038_v45_ = _dafny.Set({d_1033_v40_})
                def iife114_():
                    coll54_ = _dafny.Set()
                    compr_54_: str
                    for compr_54_ in ((d_1030_v39_)[default__.safeIndex(368, (d_1030_v39_).length(0))]).Elements:
                        d_1039_v44_: str = compr_54_
                        if (d_1039_v44_) in ((d_1030_v39_)[default__.safeIndex(368, (d_1030_v39_).length(0))]):
                            coll54_ = coll54_.union(_dafny.Set([d_1039_v44_]))
                    return _dafny.Set(coll54_)
                (globalState).f6 = default__.fm1((self).f27, ((d_1026_v35_)[default__.safeIndex(694, (d_1026_v35_).length(0))]) + (default__.fm17(36, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pxwepj")), d_990_cf47_, globalState)), (iife114_()
                ).isdisjoint(d_1038_v45_), globalState)
            d_1040_v46_: _dafny.Array
            nw150_ = _dafny.Array(int(0), 9)
            d_1040_v46_ = nw150_
            def lambda46_(d_1041_v4_):
                def lambda47_(d_1042_i7_):
                    return default__.safeModuloInt(d_1042_i7_, d_1041_v4_)

                return lambda47_

            init22_ = lambda46_(d_968_v4_)
            nw151_ = _dafny.Array(None, 18)
            for i0_22_ in range(nw151_.length(0)):
                nw151_[i0_22_] = init22_(i0_22_)
            d_1040_v46_ = nw151_
        elif True:
            d_1043___mcc_h5_ = source20_.cf44
            d_1044_cf44_ = d_1043___mcc_h5_
            (globalState).f2 = d_968_v4_
            d_1045_v47_: _dafny.MultiSet
            d_1045_v47_ = _dafny.MultiSet([default__.safeDivisionInt(len(d_966_v2_), d_968_v4_)])
            d_1046_v48_: _dafny.Seq
            d_1046_v48_ = _dafny.SeqWithoutIsStrInference([(d_1045_v47_).cardinality, d_968_v4_])
            rhs141_ = _dafny.MultiSet(d_1046_v48_)
            rhs142_ = d_968_v4_
            lhs110_ = globalState
            d_1045_v47_ = rhs141_
            lhs110_.f6 = rhs142_
            d_1047_v49_: _dafny.Map
            d_1047_v49_ = _dafny.Map({_dafny.Set({False}): d_968_v4_})
            d_1048_v50_: _dafny.Set
            d_1048_v50_ = _dafny.Set({(self).f27, (self).f27, (self).f34})
            d_1049_v51_: _dafny.Array
            def lambda48_(d_1050_v4_):
                def lambda49_(d_1051_i8_):
                    return (d_1051_i8_) * (d_1050_v4_)

                return lambda49_

            init23_ = lambda48_(d_968_v4_)
            nw152_ = _dafny.Array(None, 8)
            for i0_23_ in range(nw152_.length(0)):
                nw152_[i0_23_] = init23_(i0_23_)
            d_1049_v51_ = nw152_
            d_1052_v52_: _dafny.Map
            d_1052_v52_ = _dafny.Map({((d_1047_v49_)[d_1048_v50_] if (d_1048_v50_) in (d_1047_v49_) else (0) - ((self).fm9(d_968_v4_, globalState))): d_1049_v51_})
            d_1052_v52_ = (d_1052_v52_).set(d_968_v4_, d_1049_v51_)
            d_1053_v53_: _dafny.Set
            d_1053_v53_ = _dafny.Set({d_968_v4_, len(d_966_v2_)})
            d_1054_v54_: _dafny.Map
            d_1054_v54_ = _dafny.Map({(d_1053_v53_) | (d_1053_v53_): d_968_v4_})
            d_1055_v55_: _dafny.Seq
            d_1055_v55_ = _dafny.SeqWithoutIsStrInference([(self).f27])
            d_1054_v54_ = (d_1054_v54_).set(d_1053_v53_, default__.fm1((self).f34, d_1055_v55_, (self).f27, globalState))
        d_1056_v56_: _dafny.Seq
        d_1056_v56_ = _dafny.SeqWithoutIsStrInference([d_968_v4_])
        d_1057_v57_: D2
        d_1057_v57_ = D2_DC6(d_968_v4_, _dafny.SeqWithoutIsStrInference([d_968_v4_]), (self).f34, d_1056_v56_, True)
        pat_let_tv31_ = d_968_v4_
        pat_let_tv32_ = d_968_v4_
        def lambda50_(source22_):
            if source22_.is_DC6:
                d_1058___mcc_h11_ = source22_.cf7
                d_1059___mcc_h12_ = source22_.cf8
                d_1060___mcc_h13_ = source22_.cf9
                d_1061___mcc_h14_ = source22_.cf10
                d_1062___mcc_h15_ = source22_.cf11
                d_1063_cf11_ = d_1062___mcc_h15_
                d_1064_cf10_ = d_1061___mcc_h14_
                d_1065_cf9_ = d_1060___mcc_h13_
                d_1066_cf8_ = d_1059___mcc_h12_
                d_1067_cf7_ = d_1058___mcc_h11_
                return default__.safeDivisionInt((0) - (pat_let_tv31_), 979)
            elif source22_.is_DC5:
                d_1068___mcc_h16_ = source22_.cf6
                d_1069_cf6_ = d_1068___mcc_h16_
                return pat_let_tv32_
            elif True:
                d_1070___mcc_h17_ = source22_.cf12
                d_1071_cf12_ = d_1070___mcc_h17_
                return len(_dafny.SeqWithoutIsStrInference([(self).f34]))

        (globalState).f6 = lambda50_(D2_DC7(d_1057_v57_))
        if (d_966_v2_) < (d_966_v2_):
            if not(not((self).f27)):
                d_1072_v58_: _dafny.Seq
                d_1072_v58_ = _dafny.SeqWithoutIsStrInference([(self).f27, (self).f34, (self).f27, (self).f34, (self).f34])
                d_1073_v59_: _dafny.Map
                d_1073_v59_ = _dafny.Map({default__.safeDivisionInt(d_968_v4_, d_968_v4_): (_dafny.Map({d_1072_v58_: not((self).f27)})).set(d_1072_v58_, (self).f27)})
                d_1073_v59_ = (d_1073_v59_).set(d_968_v4_, _dafny.Map({(d_1072_v58_).set(default__.safeIndex(len(d_966_v2_), len(d_1072_v58_)), (self).f27): (self).f27}))
                d_1074_v60_: _dafny.Map
                d_1074_v60_ = _dafny.Map({d_963_v0_: d_1072_v58_})
                d_1075_v61_: _dafny.MultiSet
                d_1075_v61_ = _dafny.MultiSet([(self).f34])
                d_1076_v62_: _dafny.MultiSet
                d_1076_v62_ = _dafny.MultiSet([(d_1056_v56_)[default__.safeIndex(d_968_v4_, len(d_1056_v56_))], (len(d_1074_v60_)) * ((d_1075_v61_).cardinality), len(d_1056_v56_)])
                d_1077_v63_: _dafny.Map
                d_1077_v63_ = _dafny.Map({(self).f34: d_963_v0_})
                d_1076_v62_ = _dafny.MultiSet([d_968_v4_, len((d_1077_v63_ if (self).f27 else _dafny.Map({(self).f34: d_963_v0_}))), len(_dafny.SeqWithoutIsStrInference([d_1072_v58_, d_1072_v58_]))])
                (globalState).f21 = (d_968_v4_) < (default__.safeDivisionInt(d_968_v4_, d_968_v4_))
                (globalState).f21 = (d_1076_v62_) == ((d_1076_v62_) | (d_1076_v62_))
                (globalState).f2 = (d_968_v4_) * (d_968_v4_)
            elif True:
                d_1078_v64_: D15
                d_1078_v64_ = D15_DC33()
                d_1078_v64_ = d_1078_v64_
                d_1079_v65_: _dafny.Map
                d_1079_v65_ = _dafny.Map({_dafny.Map({(self).f27: (self).f34}): d_968_v4_})
                d_1080_v66_: _dafny.Map
                d_1080_v66_ = _dafny.Map({(self).f34: (self).f34})
                d_1079_v65_ = (d_1079_v65_).set((d_1080_v66_).set(not((self).f27), not((self).f27)), d_968_v4_)
                d_1081_v67_: _dafny.Map
                d_1081_v67_ = _dafny.Map({default__.safeDivisionInt(d_968_v4_, len(d_1056_v56_)): (918) * (d_968_v4_)})
                d_1081_v67_ = d_1081_v67_
                d_1082_v68_: _dafny.Seq
                d_1082_v68_ = _dafny.SeqWithoutIsStrInference([d_1056_v56_])
                d_1083_v69_: _dafny.Map
                d_1083_v69_ = _dafny.Map({(d_1082_v68_)[default__.safeIndex((0) - (d_968_v4_), len(d_1082_v68_))]: d_968_v4_})
                d_1084_v70_: D3
                d_1084_v70_ = D3_DC8(d_1083_v69_)
                d_1085_v71_: str
                d_1085_v71_ = _dafny.CodePoint('w')
                d_1086_v72_: _dafny.Array
                def lambda51_(d_1087_v4_):
                    def lambda52_(d_1088_i9_):
                        return (d_1088_i9_) + (d_1087_v4_)

                    return lambda52_

                init24_ = lambda51_(d_968_v4_)
                nw153_ = _dafny.Array(None, 12)
                for i0_24_ in range(nw153_.length(0)):
                    nw153_[i0_24_] = init24_(i0_24_)
                d_1086_v72_ = nw153_
                index149_ = default__.safeIndex(536, (d_1086_v72_).length(0))
                (d_1086_v72_)[index149_] = d_968_v4_
                d_1089_v73_: _dafny.Map
                d_1089_v73_ = _dafny.Map({(self).f34: (D2_DC6(d_968_v4_, d_1056_v56_, False, d_1056_v56_, (self).f27)).cf10})
                d_1090_v74_: _dafny.Map
                d_1090_v74_ = _dafny.Map({d_1086_v72_: (d_968_v4_) * (d_968_v4_)})
                index150_ = default__.safeIndex(536, (d_1086_v72_).length(0))
                rhs143_ = (self).f27
                rhs144_ = D3_DC8((d_1083_v69_).set(((d_1089_v73_)[(self).f34] if ((self).f34) in (d_1089_v73_) else d_1056_v56_), d_968_v4_))
                rhs145_ = (d_1085_v71_ if (self).f34 else d_1085_v71_)
                rhs146_ = (0) - (((d_1090_v74_)[d_1086_v72_] if (d_1086_v72_) in (d_1090_v74_) else d_968_v4_))
                rhs147_ = d_968_v4_
                lhs111_ = globalState
                lhs112_ = d_1086_v72_
                lhs113_ = default__.safeIndex(536, (d_1086_v72_).length(0))
                lhs111_.f12 = rhs143_
                d_1084_v70_ = rhs144_
                d_1085_v71_ = rhs145_
                lhs112_[lhs113_] = rhs146_
                d_968_v4_ = rhs147_
                d_1091_v75_: _dafny.Map
                d_1091_v75_ = _dafny.Map({_dafny.Set({d_968_v4_}): (self).f34})
                d_1092_v76_: _dafny.Map
                d_1092_v76_ = _dafny.Map({len((d_1091_v75_) | (d_1091_v75_)): (self).f34})
                d_1093_v77_: _dafny.Seq
                d_1093_v77_ = _dafny.SeqWithoutIsStrInference([(self).f27])
                d_1094_v78_: _dafny.MultiSet
                d_1094_v78_ = _dafny.MultiSet([d_968_v4_])
                d_1095_v79_: _dafny.Set
                d_1095_v79_ = _dafny.Set({len(_dafny.Set({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(d_1086_v72_)[default__.safeIndex(536, (d_1086_v72_).length(0))] for d_1096_i10_ in range(default__.abs(623))]))).cardinality}))})
                d_1092_v76_ = (d_1092_v76_).set(default__.safeModuloInt(default__.fm1((self).f27, d_1093_v77_, (self).f27, globalState), (d_1094_v78_).cardinality), default__.fm0(d_966_v2_, (self).f27, (self).f27, d_1095_v79_, globalState))
            d_1097_v81_: str
            d_1097_v81_ = _dafny.CodePoint('n')
            d_1098_v82_: _dafny.Map
            d_1098_v82_ = _dafny.Map({d_968_v4_: _dafny.CodePoint('b')})
            d_1099_v83_: _dafny.Map
            d_1099_v83_ = _dafny.Map({_dafny.CodePoint('k'): d_1098_v82_})
            def iife115_():
                coll55_ = _dafny.Map()
                compr_55_: str
                for compr_55_ in (_dafny.Set({d_1097_v81_})).Elements:
                    d_1100_v80_: str = compr_55_
                    if (d_1100_v80_) in (_dafny.Set({d_1097_v81_})):
                        coll55_[d_1100_v80_] = _dafny.Map({d_968_v4_: _dafny.CodePoint('r')})
                return _dafny.Map(coll55_)
            d_966_v2_ = (self).fm2(d_968_v4_, (iife115_()
             if True else d_1099_v83_), globalState)
            d_1101_v84_: _dafny.Seq
            d_1101_v84_ = _dafny.SeqWithoutIsStrInference([(self).f27, (self).f34])
            (globalState).f21 = (len((d_1101_v84_) + (d_1101_v84_))) == (len(d_966_v2_))
            d_1102_v85_: _dafny.Map
            d_1102_v85_ = _dafny.Map({(d_968_v4_) * (928): default__.safeDivisionInt(d_968_v4_, d_968_v4_)})
            d_968_v4_ = ((d_1102_v85_)[d_968_v4_] if (d_968_v4_) in (d_1102_v85_) else d_968_v4_)
            d_1103_v86_: D3
            d_1103_v86_ = D3_DC8(_dafny.Map({_dafny.SeqWithoutIsStrInference([130 for d_1104_i11_ in range(default__.abs(616))]): d_968_v4_}))
            d_1105_v88_: _dafny.Map
            d_1105_v88_ = _dafny.Map({(self).f27: (self).f34})
            d_1106_v89_: _dafny.Map
            d_1106_v89_ = _dafny.Map({(d_1056_v56_).set(default__.safeIndex(802, len(d_1056_v56_)), d_968_v4_): len(d_1105_v88_)})
            pat_let_tv33_ = d_1106_v89_
            pat_let_tv34_ = d_1106_v89_
            pat_let_tv35_ = d_968_v4_
            d_1107_v91_: _dafny.Array
            nw154_ = _dafny.Array(None, 7)
            def iife116_(_pat_let30_0):
                def iife117_(d_1108_dt__update__tmp_h0_):
                    def iife119_():
                        coll56_ = _dafny.Map()
                        compr_56_: _dafny.Seq
                        for compr_56_ in (pat_let_tv33_).keys.Elements:
                            d_1109_v87_: _dafny.Seq = compr_56_
                            if (d_1109_v87_) in (pat_let_tv34_):
                                coll56_[d_1109_v87_] = pat_let_tv35_
                        return _dafny.Map(coll56_)
                    def iife118_(_pat_let31_0):
                        def iife120_(d_1110_dt__update_hcf13_h0_):
                            return D3_DC8(d_1110_dt__update_hcf13_h0_)
                        return iife120_(_pat_let31_0)
                    return iife118_(iife119_()
                    )
                return iife117_(_pat_let30_0)
            nw154_[int(0)] = iife116_(d_1103_v86_)
            nw154_[int(1)] = d_1103_v86_
            nw154_[int(2)] = d_1103_v86_
            def iife121_():
                coll57_ = _dafny.Map()
                compr_57_: _dafny.Seq
                for compr_57_ in (d_1106_v89_).keys.Elements:
                    d_1111_v90_: _dafny.Seq = compr_57_
                    if (d_1111_v90_) in (d_1106_v89_):
                        coll57_[d_1111_v90_] = d_968_v4_
                return _dafny.Map(coll57_)
            nw154_[int(3)] = D3_DC8(iife121_()
)
            nw154_[int(4)] = d_1103_v86_
            nw154_[int(5)] = d_1103_v86_
            nw154_[int(6)] = d_1103_v86_
            d_1107_v91_ = nw154_
            index151_ = default__.safeIndex(169, (d_1107_v91_).length(0))
            (d_1107_v91_)[index151_] = d_1103_v86_
            index152_ = default__.safeIndex(169, (d_1107_v91_).length(0))
            (d_1107_v91_)[index152_] = d_1103_v86_
        elif True:
            (globalState).f21 = (self).f34
            (globalState).f6 = d_968_v4_
            d_1112_v92_: D9
            d_1112_v92_ = D9_DC21((self).f34, d_968_v4_)
            (globalState).f21 = (d_1112_v92_).cf31
            d_1113_v93_: _dafny.Set
            d_1113_v93_ = _dafny.Set({default__.safeModuloInt(d_968_v4_, d_968_v4_), d_968_v4_, d_968_v4_, d_968_v4_})
            d_1114_v94_: _dafny.Seq
            d_1114_v94_ = _dafny.SeqWithoutIsStrInference([False, (self).f27])
            d_1113_v93_ = default__.fm21(default__.safeModuloInt(default__.fm1((self).f27, d_1114_v94_, (self).f34, globalState), len(d_966_v2_)), globalState)
            d_1115_v95_: _dafny.Array
            nw155_ = _dafny.Array(int(0), 4)
            d_1115_v95_ = nw155_
            index153_ = default__.safeIndex(617, (d_963_v0_).length(0))
            (d_963_v0_)[index153_] = (665) > (d_968_v4_)
            d_1116_v96_: _dafny.MultiSet
            d_1116_v96_ = _dafny.MultiSet([(self).f34])
            d_1117_v97_: _dafny.MultiSet
            d_1117_v97_ = _dafny.MultiSet([d_966_v2_, d_966_v2_])
            d_1118_v98_: _dafny.Seq
            d_1118_v98_ = _dafny.SeqWithoutIsStrInference([d_1114_v94_])
            index154_ = default__.safeIndex(617, (d_963_v0_).length(0))
            def iife122_():
                coll58_ = _dafny.Set()
                compr_58_: int
                for compr_58_ in (d_1056_v56_).Elements:
                    d_1119_v99_: int = compr_58_
                    if (d_1119_v99_) in (d_1056_v56_):
                        coll58_ = coll58_.union(_dafny.Set([(d_1119_v99_) * (len(_dafny.SeqWithoutIsStrInference([True, True, not(True)])))]))
                return _dafny.Set(coll58_)
            rhs148_ = d_1115_v95_
            rhs149_ = (self).f27
            rhs150_ = default__.fm0(d_966_v2_, (d_1116_v96_) != (d_1116_v96_), (self).f27, (_dafny.Set({d_968_v4_, len(d_966_v2_), (0) - (((d_1117_v97_)[d_966_v2_] if (d_966_v2_) in (d_1117_v97_) else d_968_v4_)), d_968_v4_})) - (_dafny.Set({len((d_1118_v98_)[default__.safeIndex(d_968_v4_, len(d_1118_v98_))]), d_968_v4_})), globalState)
            rhs151_ = default__.fm0(d_966_v2_, (self).f34, not(False), iife122_()
            , globalState)
            lhs114_ = globalState
            lhs115_ = d_963_v0_
            lhs116_ = default__.safeIndex(617, (d_963_v0_).length(0))
            lhs117_ = globalState
            d_1115_v95_ = rhs148_
            lhs114_.f12 = rhs149_
            lhs115_[lhs116_] = rhs150_
            lhs117_.f14 = rhs151_
        d_1120_i12_: int
        d_1120_i12_ = 0
        with _dafny.label("5"):
            def iife123_(_pat_let32_0):
                def iife124_(d_1140_dt__update__tmp_h1_):
                    def iife125_(_pat_let33_0):
                        def iife126_(d_1141_dt__update_hcf19_h0_):
                            return D5_DC12(d_1141_dt__update_hcf19_h0_, (d_1140_dt__update__tmp_h1_).cf20)
                        return iife126_(_pat_let33_0)
                    return iife125_(True)
                return iife124_(_pat_let32_0)
            while (iife123_(D5_DC12((self).f27, 429))).cf19:
                with _dafny.c_label("5"):
                    if (d_1120_i12_) >= (100):
                        raise _dafny.Break("5")
                    d_1120_i12_ = (d_1120_i12_) + (1)
                    d_1121_v100_: _dafny.Set
                    d_1121_v100_ = _dafny.Set({(self).f27})
                    d_1122_v101_: _dafny.MultiSet
                    d_1122_v101_ = _dafny.MultiSet([D11_DC24(d_1121_v100_)])
                    d_1123_v102_: _dafny.Map
                    d_1123_v102_ = _dafny.Map({d_968_v4_: d_1122_v101_})
                    d_1124_v103_: _dafny.Array
                    nw156_ = _dafny.Array(int(0), 29)
                    d_1124_v103_ = nw156_
                    d_1125_v104_: _dafny.Map
                    d_1125_v104_ = _dafny.Map({d_1123_v102_: d_1124_v103_})
                    d_1126_v105_: _dafny.Map
                    d_1126_v105_ = _dafny.Map({d_968_v4_: _dafny.Map({d_968_v4_: default__.fm22(False, globalState)})})
                    d_1125_v104_ = (d_1125_v104_).set(((d_1126_v105_)[d_968_v4_] if (d_968_v4_) in (d_1126_v105_) else d_1123_v102_), (d_1124_v103_ if True else d_1124_v103_))
                    if default__.fm0((d_966_v2_) + (d_966_v2_), (self).f34, False, _dafny.Set({len(d_1121_v100_)}), globalState):
                        (globalState).f5 = (self).f27
                        index155_ = default__.safeIndex(140, (d_1124_v103_).length(0))
                        (d_1124_v103_)[index155_] = default__.safeDivisionInt(d_968_v4_, d_968_v4_)
                        d_1127_v106_: _dafny.Set
                        d_1127_v106_ = _dafny.Set({d_968_v4_, d_968_v4_})
                        index156_ = default__.safeIndex(140, (d_1124_v103_).length(0))
                        rhs152_ = d_968_v4_
                        rhs153_ = not ((self).f27) or (default__.fm0(d_966_v2_, (self).f34, (self).f34, d_1127_v106_, globalState))
                        rhs154_ = (d_968_v4_) * (d_968_v4_)
                        lhs118_ = globalState
                        lhs119_ = globalState
                        lhs120_ = d_1124_v103_
                        lhs121_ = default__.safeIndex(140, (d_1124_v103_).length(0))
                        lhs118_.f2 = rhs152_
                        lhs119_.f12 = rhs153_
                        lhs120_[lhs121_] = rhs154_
                        d_1128_v107_: C0
                        nw157_ = C0()
                        nw157_.ctor__()
                        d_1128_v107_ = nw157_
                        d_1129_v108_: _dafny.Seq
                        d_1129_v108_ = _dafny.SeqWithoutIsStrInference([(_dafny.Set({d_968_v4_, (0) - (len(d_966_v2_))})).ispropersubset(d_1127_v106_)])
                        d_1129_v108_ = _dafny.SeqWithoutIsStrInference([((self).f27 if (self).f34 else (self).f27), ((d_1124_v103_)[default__.safeIndex(140, (d_1124_v103_).length(0))]) <= (d_968_v4_)])
                        (globalState).f6 = (d_1124_v103_)[default__.safeIndex(140, (d_1124_v103_).length(0))]
                    elif True:
                        (globalState).f12 = (self).f27
                        d_1130_v109_: _dafny.Seq
                        d_1130_v109_ = _dafny.SeqWithoutIsStrInference([(self).f27, ((self).f27 if (self).f27 else (self).f34)])
                        d_1130_v109_ = d_1130_v109_
                        d_1131_v110_: str
                        d_1131_v110_ = _dafny.CodePoint('s')
                        d_1132_v111_: _dafny.Set
                        d_1132_v111_ = _dafny.Set({d_1131_v110_})
                        d_1133_v112_: _dafny.Map
                        d_1133_v112_ = _dafny.Map({(self).f27: d_968_v4_})
                        d_1134_v113_: _dafny.MultiSet
                        d_1134_v113_ = _dafny.MultiSet([len((d_1132_v111_) | (d_1132_v111_)), ((d_1133_v112_)[(self).f27] if ((self).f27) in (d_1133_v112_) else len(d_1130_v109_)), d_968_v4_])
                        d_1134_v113_ = d_1134_v113_
                        d_1135_v114_: _dafny.Map
                        d_1135_v114_ = _dafny.Map({d_1124_v103_: (self).f34})
                        d_1135_v114_ = (d_1135_v114_).set(d_1124_v103_, (self).f34)
                        index157_ = default__.safeIndex(135, (d_1124_v103_).length(0))
                        (d_1124_v103_)[index157_] = d_968_v4_
                        index158_ = default__.safeIndex(135, (d_1124_v103_).length(0))
                        (d_1124_v103_)[index158_] = d_968_v4_
                    d_1136_v115_: C1
                    nw158_ = C1()
                    nw158_.ctor__((self).f34)
                    d_1136_v115_ = nw158_
                    d_1137_v116_: _dafny.MultiSet
                    d_1137_v116_ = _dafny.MultiSet([(self).f27])
                    index159_ = default__.safeIndex(383, (d_1124_v103_).length(0))
                    (d_1124_v103_)[index159_] = (d_1056_v56_)[default__.safeIndex(((d_1137_v116_)[(self).f27] if ((self).f27) in (d_1137_v116_) else d_968_v4_), len(d_1056_v56_))]
                    d_1138_v117_: D0
                    d_1138_v117_ = D0_DC1((self).f27)
                    d_1139_v118_: _dafny.Seq
                    d_1139_v118_ = _dafny.SeqWithoutIsStrInference([(self).f34])
                    index160_ = default__.safeIndex(383, (d_1124_v103_).length(0))
                    (d_1124_v103_)[index160_] = (d_968_v4_) * ((0) - (default__.fm1((self).fm4(d_1138_v117_, globalState), d_1139_v118_, (self).f27, globalState)))
                    pass
            pass
        if (self).f27:
            d_1142_v119_: _dafny.Seq
            d_1142_v119_ = _dafny.SeqWithoutIsStrInference([d_966_v2_, d_966_v2_])
            d_1143_v120_: D13
            d_1143_v120_ = D13_DC27(d_1142_v119_)
            d_1144_v121_: _dafny.Seq
            d_1144_v121_ = _dafny.SeqWithoutIsStrInference([D13_DC27(d_1142_v119_), d_1143_v120_])
            d_1144_v121_ = (d_1144_v121_ if not((self).f34) else d_1144_v121_)
            d_1145_v122_: str
            d_1145_v122_ = _dafny.CodePoint('s')
            d_1146_v123_: _dafny.Map
            d_1146_v123_ = _dafny.Map({d_968_v4_: d_1145_v122_})
            d_1147_v124_: _dafny.MultiSet
            d_1147_v124_ = _dafny.MultiSet([False, (self).f27])
            d_1148_v125_: _dafny.Set
            d_1148_v125_ = _dafny.Set({_dafny.CodePoint('y'), d_1145_v122_})
            d_1149_v126_: _dafny.Map
            d_1149_v126_ = _dafny.Map({(d_1147_v124_).cardinality: d_1148_v125_})
            d_1150_v127_: _dafny.MultiSet
            d_1150_v127_ = _dafny.MultiSet([(d_968_v4_) * (d_968_v4_), (d_968_v4_ if not(not((self).f34)) else d_968_v4_), len(d_966_v2_), len(d_1146_v123_), len(((d_1149_v126_)[(_dafny.MultiSet([d_968_v4_, d_968_v4_])).cardinality] if ((_dafny.MultiSet([d_968_v4_, d_968_v4_])).cardinality) in (d_1149_v126_) else d_1148_v125_))])
            d_1151_v128_: D16
            d_1151_v128_ = D16_DC35(default__.fm13(globalState))
            d_1150_v127_ = (d_1151_v128_).cf50
            d_1152_v131_: _dafny.Map
            d_1152_v131_ = _dafny.Map({d_968_v4_: d_968_v4_})
            d_1153_v132_: _dafny.Map
            d_1153_v132_ = _dafny.Map({((d_1152_v131_)[d_968_v4_] if (d_968_v4_) in (d_1152_v131_) else d_968_v4_): (self).f27})
            d_1154_v133_: _dafny.Map
            d_1154_v133_ = _dafny.Map({(self).f27: d_968_v4_})
            def iife127_():
                coll59_ = _dafny.Map()
                compr_59_: int
                for compr_59_ in _dafny.IntegerRange(151, 401):
                    d_1155_v129_: int = compr_59_
                    if ((151) <= (d_1155_v129_)) and ((d_1155_v129_) < (401)):
                        coll59_[default__.safeDivisionInt(d_1155_v129_, d_968_v4_)] = (self).f34
                return _dafny.Map(coll59_)
            def iife128_():
                coll60_ = _dafny.Map()
                compr_60_: int
                for compr_60_ in (d_1150_v127_).Elements:
                    d_1156_v130_: int = compr_60_
                    if (d_1156_v130_) in (d_1150_v127_):
                        coll60_[(d_1156_v130_) - (d_968_v4_)] = (self).f27
                return _dafny.Map(coll60_)
            rhs155_ = ((iife127_()
            ) | (iife128_()
            )) == (d_1153_v132_)
            rhs156_ = d_1145_v122_
            rhs157_ = ((d_1150_v127_)[d_968_v4_] if (d_968_v4_) in (d_1150_v127_) else ((d_1154_v133_)[(self).f34] if ((self).f34) in (d_1154_v133_) else d_968_v4_))
            rhs158_ = d_968_v4_
            lhs122_ = globalState
            lhs123_ = globalState
            lhs122_.f14 = rhs155_
            d_1145_v122_ = rhs156_
            d_968_v4_ = rhs157_
            lhs123_.f2 = rhs158_
            (globalState).f21 = (self).f34
            index161_ = default__.safeIndex(550, (d_963_v0_).length(0))
            (d_963_v0_)[index161_] = not(False)
            d_1157_v134_: _dafny.Set
            d_1157_v134_ = _dafny.Set({(self).f34, (self).f34, (self).f34})
            d_1158_v135_: C1
            nw159_ = C1()
            nw159_.ctor__((self).f34)
            d_1158_v135_ = nw159_
            index162_ = default__.safeIndex(550, (d_963_v0_).length(0))
            rhs159_ = (self).f27
            rhs160_ = d_1157_v134_
            rhs161_ = d_1158_v135_
            rhs162_ = d_1157_v134_
            rhs163_ = default__.safeDivisionInt(d_968_v4_, -721)
            lhs124_ = d_963_v0_
            lhs125_ = default__.safeIndex(550, (d_963_v0_).length(0))
            lhs126_ = globalState
            lhs124_[lhs125_] = rhs159_
            d_1157_v134_ = rhs160_
            d_1158_v135_ = rhs161_
            d_1157_v134_ = rhs162_
            lhs126_.f6 = rhs163_
        elif True:
            r0 = -76
            d_1159_v136_: _dafny.MultiSet
            d_1159_v136_ = _dafny.MultiSet([d_966_v2_])
            d_1159_v136_ = d_1159_v136_
            d_1160_v137_: str
            d_1160_v137_ = _dafny.CodePoint('t')
            d_1161_v138_: str
            d_1161_v138_ = d_1160_v137_
            source23_ = d_1161_v138_
            d_1162___mcc_h18_ = source23_
            d_1163_cf36_ = d_1162___mcc_h18_
            d_1164_v139_: _dafny.Map
            d_1164_v139_ = _dafny.Map({d_1163_cf36_: (d_968_v4_) > (d_968_v4_)})
            d_1165_v140_: _dafny.Seq
            d_1165_v140_ = _dafny.SeqWithoutIsStrInference([(self).f34])
            d_1164_v139_ = (d_1164_v139_).set(_dafny.CodePoint('y'), (d_968_v4_) != (default__.fm1((self).f34, d_1165_v140_, (self).f34, globalState)))
            d_1166_v141_: _dafny.Array
            nw160_ = _dafny.Array(int(0), 8)
            d_1166_v141_ = nw160_
            index163_ = default__.safeIndex(276, (d_1166_v141_).length(0))
            (d_1166_v141_)[index163_] = len((d_966_v2_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lmbxf"))))
            index164_ = default__.safeIndex(276, (d_1166_v141_).length(0))
            (d_1166_v141_)[index164_] = (d_968_v4_ if (self).f27 else d_968_v4_)
            (globalState).f6 = d_968_v4_
            r0 = (0) - ((80 if (d_968_v4_) != ((d_1166_v141_)[default__.safeIndex(276, (d_1166_v141_).length(0))]) else default__.safeModuloInt(-846, d_968_v4_)))
            d_1167_v142_: _dafny.MultiSet
            d_1167_v142_ = _dafny.MultiSet([d_968_v4_, (0) - (d_968_v4_)])
            d_1168_v143_: _dafny.MultiSet
            d_1168_v143_ = _dafny.MultiSet([True])
            d_1169_v144_: _dafny.Map
            d_1169_v144_ = _dafny.Map({(self).f27: d_968_v4_})
            d_1170_v145_: _dafny.Array
            nw161_ = _dafny.Array(None, 19)
            nw161_[int(0)] = d_968_v4_
            nw161_[int(1)] = len(d_966_v2_)
            nw161_[int(2)] = d_968_v4_
            nw161_[int(3)] = d_968_v4_
            nw161_[int(4)] = len((d_1056_v56_) + (_dafny.SeqWithoutIsStrInference([d_968_v4_, -624])))
            nw161_[int(5)] = ((d_1167_v142_)[len(d_966_v2_)] if (len(d_966_v2_)) in (d_1167_v142_) else d_968_v4_)
            nw161_[int(6)] = d_968_v4_
            nw161_[int(7)] = d_968_v4_
            nw161_[int(8)] = (d_1168_v143_).cardinality
            nw161_[int(9)] = d_968_v4_
            nw161_[int(10)] = d_968_v4_
            nw161_[int(11)] = ((d_1169_v144_)[(self).f34] if ((self).f34) in (d_1169_v144_) else 356)
            nw161_[int(12)] = d_968_v4_
            nw161_[int(13)] = (self).fm9(-668, globalState)
            nw161_[int(14)] = d_968_v4_
            nw161_[int(15)] = d_968_v4_
            nw161_[int(16)] = len(_dafny.Set({len(d_1056_v56_)}))
            nw161_[int(17)] = d_968_v4_
            nw161_[int(18)] = d_968_v4_
            d_1170_v145_ = nw161_
            index165_ = default__.safeIndex(205, (d_1170_v145_).length(0))
            (d_1170_v145_)[index165_] = (d_968_v4_) + (d_968_v4_)
            index166_ = default__.safeIndex(205, (d_1170_v145_).length(0))
            (d_1170_v145_)[index166_] = len(d_966_v2_)
            d_1171_v146_: C1
            nw162_ = C1()
            nw162_.ctor__(False)
            d_1171_v146_ = nw162_
        r0 = len(d_966_v2_)
        return r0

    def m2(self, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: bool = False
        d_1172_v0_: _dafny.Array
        nw163_ = _dafny.Array(_dafny.Set({}), 3)
        d_1172_v0_ = nw163_
        d_1173_v1_: D7
        d_1173_v1_ = D7_DC16()
        d_1174_v2_: _dafny.Seq
        d_1174_v2_ = _dafny.SeqWithoutIsStrInference([d_1173_v1_, d_1173_v1_, d_1173_v1_])
        d_1175_v3_: str
        d_1175_v3_ = _dafny.CodePoint('v')
        d_1176_v4_: int
        d_1176_v4_ = 419
        d_1177_v5_: _dafny.Map
        d_1177_v5_ = _dafny.Map({d_1175_v3_: (0) - (d_1176_v4_)})
        d_1178_v6_: _dafny.MultiSet
        d_1178_v6_ = _dafny.MultiSet([d_1177_v5_, d_1177_v5_])
        d_1179_v7_: _dafny.MultiSet
        d_1179_v7_ = d_1178_v6_
        d_1180_v8_: _dafny.Seq
        d_1180_v8_ = _dafny.SeqWithoutIsStrInference([(d_1174_v2_) + (d_1174_v2_), d_1174_v2_, d_1174_v2_])
        pat_let_tv36_ = d_1176_v4_
        def lambda53_(source24_):
            d_1181___mcc_h0_ = source24_
            d_1182_cf29_ = d_1181___mcc_h0_
            return (306) <= (pat_let_tv36_)

        rhs164_ = lambda53_(d_1179_v7_)
        rhs165_ = d_1172_v0_
        rhs166_ = 839
        rhs167_ = (d_1180_v8_)[default__.safeIndex(d_1176_v4_, len(d_1180_v8_))]
        lhs127_ = globalState
        r3 = rhs164_
        d_1172_v0_ = rhs165_
        lhs127_.f6 = rhs166_
        d_1174_v2_ = rhs167_
        if (self).f34:
            d_1183_v9_: _dafny.Array
            nw164_ = _dafny.Array(None, 4)
            nw164_[int(0)] = (self).f34
            nw164_[int(1)] = (self).f34
            nw164_[int(2)] = True
            nw164_[int(3)] = (self).f27
            d_1183_v9_ = nw164_
            index167_ = default__.safeIndex(369, (d_1183_v9_).length(0))
            (d_1183_v9_)[index167_] = not ((self).f27) or ((self).f34)
            index168_ = default__.safeIndex(369, (d_1183_v9_).length(0))
            (d_1183_v9_)[index168_] = (self).f34
            d_1184_v10_: _dafny.Seq
            d_1184_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "usbcov"))
            d_1185_v11_: _dafny.Set
            d_1185_v11_ = _dafny.Set({d_1176_v4_})
            d_1186_v12_: _dafny.Seq
            d_1186_v12_ = _dafny.SeqWithoutIsStrInference([(d_1183_v9_)[default__.safeIndex(369, (d_1183_v9_).length(0))], default__.fm0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gkykhuj")), (self).f34, (self).f34, d_1185_v11_, globalState), (d_1183_v9_)[default__.safeIndex(369, (d_1183_v9_).length(0))]])
            d_1187_v13_: _dafny.Seq
            d_1187_v13_ = _dafny.SeqWithoutIsStrInference([d_1176_v4_])
            d_1188_v14_: _dafny.Map
            d_1188_v14_ = _dafny.Map({d_1187_v13_: d_1176_v4_})
            d_1189_v15_: _dafny.Array
            nw165_ = _dafny.Array(None, 13)
            nw165_[int(0)] = default__.fm1((d_1183_v9_)[default__.safeIndex(369, (d_1183_v9_).length(0))], default__.fm17(d_1176_v4_, d_1184_v10_, d_1176_v4_, globalState), (self).f27, globalState)
            nw165_[int(1)] = (self).fm9(d_1176_v4_, globalState)
            nw165_[int(2)] = default__.safeModuloInt(d_1176_v4_, d_1176_v4_)
            nw165_[int(3)] = (len(d_1186_v12_)) + ((self).fm9(d_1176_v4_, globalState))
            nw165_[int(4)] = 395
            nw165_[int(5)] = len(d_1187_v13_)
            nw165_[int(6)] = ((0) - (d_1176_v4_)) * (len(d_1188_v14_))
            nw165_[int(7)] = 551
            nw165_[int(8)] = d_1176_v4_
            nw165_[int(9)] = d_1176_v4_
            nw165_[int(10)] = d_1176_v4_
            nw165_[int(11)] = d_1176_v4_
            nw165_[int(12)] = default__.fm1((d_1186_v12_)[default__.safeIndex(d_1176_v4_, len(d_1186_v12_))], d_1186_v12_, not((d_1183_v9_)[default__.safeIndex(369, (d_1183_v9_).length(0))]), globalState)
            d_1189_v15_ = nw165_
            d_1189_v15_ = d_1189_v15_
            d_1190_v16_: _dafny.MultiSet
            d_1190_v16_ = _dafny.MultiSet([(d_1183_v9_)[default__.safeIndex(369, (d_1183_v9_).length(0))], ((d_1184_v10_)[default__.safeIndex(869, len(d_1184_v10_))]) not in (d_1184_v10_)])
            d_1190_v16_ = d_1190_v16_
            d_1191_v17_: C0
            nw166_ = C0()
            nw166_.ctor__()
            d_1191_v17_ = nw166_
            d_1189_v15_ = d_1189_v15_
        elif True:
            (globalState).f2 = -206
            d_1192_v18_: _dafny.Set
            d_1192_v18_ = _dafny.Set({(self).f34})
            (globalState).f2 = (d_1176_v4_) - ((len(_dafny.SeqWithoutIsStrInference([d_1175_v3_ for d_1193_i0_ in range(default__.abs(832))])) if (self).f34 else len(d_1192_v18_)))
            d_1194_v19_: _dafny.Array
            def lambda54_(d_1195_i1_):
                return (D0_DC1((self).f27)).cf1

            init25_ = lambda54_
            nw167_ = _dafny.Array(None, 17)
            for i0_25_ in range(nw167_.length(0)):
                nw167_[i0_25_] = init25_(i0_25_)
            d_1194_v19_ = nw167_
            index169_ = default__.safeIndex(730, (d_1194_v19_).length(0))
            (d_1194_v19_)[index169_] = True
            d_1196_v20_: _dafny.Seq
            d_1196_v20_ = _dafny.SeqWithoutIsStrInference([d_1176_v4_])
            d_1197_v21_: _dafny.Map
            d_1197_v21_ = _dafny.Map({(0) - ((0) - (d_1176_v4_)): d_1196_v20_})
            d_1198_v22_: _dafny.Set
            d_1198_v22_ = _dafny.Set({d_1196_v20_})
            index170_ = default__.safeIndex(730, (d_1194_v19_).length(0))
            rhs168_ = (((d_1197_v21_)[(0) - ((0) - ((self).fm9(d_1176_v4_, globalState)))] if ((0) - ((0) - ((self).fm9(d_1176_v4_, globalState)))) in (d_1197_v21_) else d_1196_v20_)) not in (d_1198_v22_)
            rhs169_ = (self).f34
            lhs128_ = globalState
            lhs129_ = d_1194_v19_
            lhs130_ = default__.safeIndex(730, (d_1194_v19_).length(0))
            lhs128_.f25 = rhs168_
            lhs129_[lhs130_] = rhs169_
            d_1199_v23_: _dafny.Seq
            d_1199_v23_ = _dafny.SeqWithoutIsStrInference([(self).f34, (self).f34])
            d_1200_v24_: _dafny.Seq
            d_1200_v24_ = _dafny.SeqWithoutIsStrInference([(d_1199_v23_)[default__.safeIndex(d_1176_v4_, len(d_1199_v23_))]])
            d_1201_v25_: _dafny.Seq
            d_1201_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tt"))
            d_1202_v26_: _dafny.Map
            d_1202_v26_ = _dafny.Map({len(default__.fm17(d_1176_v4_, d_1201_v25_, d_1176_v4_, globalState)): d_1176_v4_})
            d_1200_v24_ = (d_1199_v23_ if (d_1176_v4_) not in (d_1202_v26_) else default__.fm17(d_1176_v4_, d_1201_v25_, d_1176_v4_, globalState))
            d_1203_v27_: _dafny.Array
            nw168_ = _dafny.Array(int(0), 25)
            d_1203_v27_ = nw168_
            index171_ = default__.safeIndex(186, (d_1203_v27_).length(0))
            (d_1203_v27_)[index171_] = d_1176_v4_
            d_1204_v28_: _dafny.Map
            d_1204_v28_ = _dafny.Map({(self).f27: (self).f34})
            d_1205_v29_: _dafny.Map
            d_1205_v29_ = _dafny.Map({d_1176_v4_: d_1175_v3_})
            d_1206_v30_: _dafny.Map
            d_1206_v30_ = _dafny.Map({d_1175_v3_: d_1205_v29_})
            d_1207_v31_: D0
            d_1207_v31_ = D0_DC0(d_1176_v4_)
            index172_ = default__.safeIndex(186, (d_1203_v27_).length(0))
            rhs170_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mdentr")) if (((self).fm2(len(d_1204_v28_), d_1206_v30_, globalState)).set(default__.safeIndex((self).fm9((d_1207_v31_).cf0, globalState), len((self).fm2(len(d_1204_v28_), d_1206_v30_, globalState))), d_1175_v3_)) <= (d_1201_v25_) else d_1201_v25_)
            rhs171_ = len((default__.fm19(_dafny.SeqWithoutIsStrInference([d_1175_v3_ for d_1208_i2_ in range(default__.abs(94))]), globalState)) - (d_1192_v18_))
            rhs172_ = (d_1194_v19_)[default__.safeIndex(730, (d_1194_v19_).length(0))]
            rhs173_ = _dafny.SeqWithoutIsStrInference([d_1175_v3_ for d_1209_i3_ in range(default__.abs(771))])
            lhs131_ = d_1203_v27_
            lhs132_ = default__.safeIndex(186, (d_1203_v27_).length(0))
            lhs133_ = globalState
            r1 = rhs170_
            lhs131_[lhs132_] = rhs171_
            lhs133_.f12 = rhs172_
            r1 = rhs173_
        d_1210_v32_: _dafny.Array
        nw169_ = _dafny.Array(int(0), 7)
        d_1210_v32_ = nw169_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_1210_v32_).length(0)):
            d_1211_i4_: int = guard_loop_2_
            if (True) and (((0) <= (d_1211_i4_)) and ((d_1211_i4_) < ((d_1210_v32_).length(0)))):
                (d_1210_v32_)[(d_1211_i4_)] = (d_1211_i4_) - (len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lhdjnit"))).set(default__.safeIndex(d_1176_v4_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lhdjnit")))), _dafny.CodePoint('i'))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sgmljgpws")))))
        d_1212_i5_: int
        d_1212_i5_ = 0
        with _dafny.label("6"):
            while not (not(((self).f34) not in (_dafny.Map({not((self).f34): d_1176_v4_})))) or ((self).f27):
                with _dafny.c_label("6"):
                    if (d_1212_i5_) >= (100):
                        raise _dafny.Break("6")
                    d_1212_i5_ = (d_1212_i5_) + (1)
                    d_1213_v33_: _dafny.Map
                    d_1213_v33_ = _dafny.Map({(self).f34: d_1176_v4_})
                    d_1214_v34_: _dafny.Seq
                    d_1214_v34_ = _dafny.SeqWithoutIsStrInference([(self).f34, (self).f27, (d_1213_v33_) == (d_1213_v33_)])
                    d_1214_v34_ = (d_1214_v34_) + (d_1214_v34_)
                    d_1215_v35_: D0
                    d_1215_v35_ = D0_DC1((self).f34)
                    d_1216_v36_: C1
                    nw170_ = C1()
                    nw170_.ctor__((self).f34)
                    d_1216_v36_ = nw170_
                    d_1217_v37_: _dafny.Map
                    d_1217_v37_ = _dafny.Map({d_1176_v4_: d_1216_v36_})
                    d_1218_v38_: _dafny.Set
                    d_1218_v38_ = _dafny.Set({d_1176_v4_, (0) - (default__.safeDivisionInt(default__.fm1((self).fm4(d_1215_v35_, globalState), (d_1214_v34_).set(default__.safeIndex((0) - (len(d_1217_v37_)), len(d_1214_v34_)), (self).f34), True, globalState), d_1176_v4_))})
                    d_1219_v39_: _dafny.Seq
                    d_1219_v39_ = _dafny.SeqWithoutIsStrInference([d_1176_v4_])
                    d_1220_v40_: _dafny.Map
                    d_1220_v40_ = _dafny.Map({len(d_1219_v39_): d_1176_v4_})
                    d_1221_v41_: _dafny.Map
                    d_1221_v41_ = _dafny.Map({d_1176_v4_: d_1210_v32_})
                    d_1222_v42_: _dafny.Seq
                    d_1222_v42_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "js"))
                    d_1218_v38_ = (d_1218_v38_) - (_dafny.Set({len(d_1218_v38_), ((d_1220_v40_)[len(d_1221_v41_)] if (len(d_1221_v41_)) in (d_1220_v40_) else (_dafny.MultiSet([579, d_1176_v4_, len(_dafny.Map({default__.fm0(d_1222_v42_, (self).f27, (self).f27, d_1218_v38_, globalState): (self).f27})), d_1176_v4_, d_1176_v4_])).cardinality), d_1176_v4_, d_1176_v4_}))
                    d_1223_v43_: _dafny.Array
                    nw171_ = _dafny.Array(_dafny.Seq({}), 26)
                    d_1223_v43_ = nw171_
                    index173_ = default__.safeIndex(138, (d_1223_v43_).length(0))
                    (d_1223_v43_)[index173_] = d_1219_v39_
                    index174_ = default__.safeIndex(138, (d_1223_v43_).length(0))
                    (d_1223_v43_)[index174_] = (d_1219_v39_) + (_dafny.SeqWithoutIsStrInference([d_1176_v4_ for d_1224_i6_ in range(default__.abs(954))]))
                    d_1225_v44_: D0
                    d_1225_v44_ = D0_DC0((0) - (d_1176_v4_))
                    d_1226_v45_: D0
                    d_1226_v45_ = D0_DC3(d_1225_v44_)
                    d_1227_v46_: _dafny.Map
                    d_1227_v46_ = _dafny.Map({(self).f27: d_1226_v45_})
                    d_1227_v46_ = (d_1227_v46_).set((self).f27, default__.fm16((self).f27, d_1176_v4_, globalState))
                    pass
            pass
        d_1228_i7_: int
        d_1228_i7_ = 0
        with _dafny.label("7"):
            while (self).f27:
                with _dafny.c_label("7"):
                    if (d_1228_i7_) >= (100):
                        raise _dafny.Break("7")
                    d_1228_i7_ = (d_1228_i7_) + (1)
                    (globalState).f6 = d_1176_v4_
                    d_1176_v4_ = d_1176_v4_
                    d_1229_v47_: _dafny.Seq
                    d_1229_v47_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dgomlidx"))
                    rhs174_ = (self).f34
                    rhs175_ = (self).fm3(globalState)
                    rhs176_ = (self).f27
                    rhs177_ = len(d_1229_v47_)
                    lhs134_ = globalState
                    lhs135_ = globalState
                    lhs136_ = globalState
                    lhs134_.f21 = rhs174_
                    r1 = rhs175_
                    lhs135_.f25 = rhs176_
                    lhs136_.f2 = rhs177_
                    d_1230_v48_: C1
                    nw172_ = C1()
                    nw172_.ctor__((not((self).f34) if not((self).f34) else (self).f34))
                    d_1230_v48_ = nw172_
                    pass
            pass
        d_1231_v49_: _dafny.Seq
        d_1231_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))
        d_1232_v50_: _dafny.Seq
        d_1232_v50_ = _dafny.SeqWithoutIsStrInference([(self).f34, (self).f34, (self).f34, (self).f34, (self).f34])
        d_1233_v51_: _dafny.Seq
        d_1233_v51_ = _dafny.SeqWithoutIsStrInference([d_1176_v4_])
        d_1234_v52_: _dafny.MultiSet
        d_1234_v52_ = _dafny.MultiSet([len(d_1232_v50_), len(d_1233_v51_), d_1176_v4_, len(d_1232_v50_)])
        d_1235_v54_: _dafny.MultiSet
        def iife129_():
            coll61_ = _dafny.Set()
            compr_61_: int
            for compr_61_ in (d_1234_v52_).Elements:
                d_1236_v53_: int = compr_61_
                if (d_1236_v53_) in (d_1234_v52_):
                    coll61_ = coll61_.union(_dafny.Set([(d_1236_v53_) + (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jule")))])))]))
            return _dafny.Set(coll61_)
        d_1235_v54_ = _dafny.MultiSet([default__.fm0(d_1231_v49_, (self).f27, (self).f34, iife129_()
        , globalState)])
        index175_ = default__.safeIndex(708, (d_1210_v32_).length(0))
        (d_1210_v32_)[index175_] = ((d_1235_v54_)[(self).f27] if ((self).f27) in (d_1235_v54_) else -332)
        index176_ = default__.safeIndex(708, (d_1210_v32_).length(0))
        (d_1210_v32_)[index176_] = default__.safeModuloInt(d_1176_v4_, d_1176_v4_)
        d_1237_v55_: _dafny.Array
        def lambda55_(d_1238_i8_):
            return True

        init26_ = lambda55_
        nw173_ = _dafny.Array(None, 16)
        for i0_26_ in range(nw173_.length(0)):
            nw173_[i0_26_] = init26_(i0_26_)
        d_1237_v55_ = nw173_
        d_1239_v56_: D17
        d_1239_v56_ = D17_DC38(d_1237_v55_)
        d_1240_v57_: _dafny.Array
        nw174_ = _dafny.Array(None, 20)
        nw174_[int(0)] = d_1237_v55_
        nw174_[int(1)] = d_1237_v55_
        nw174_[int(2)] = (d_1239_v56_).cf53
        nw174_[int(3)] = (d_1239_v56_).cf53
        nw174_[int(4)] = d_1237_v55_
        nw174_[int(5)] = d_1237_v55_
        nw174_[int(6)] = d_1237_v55_
        nw174_[int(7)] = d_1237_v55_
        nw174_[int(8)] = (d_1239_v56_).cf53
        nw174_[int(9)] = d_1237_v55_
        nw174_[int(10)] = d_1237_v55_
        nw174_[int(11)] = d_1237_v55_
        nw174_[int(12)] = d_1237_v55_
        nw174_[int(13)] = d_1237_v55_
        nw174_[int(14)] = d_1237_v55_
        nw174_[int(15)] = d_1237_v55_
        nw174_[int(16)] = d_1237_v55_
        nw174_[int(17)] = d_1237_v55_
        nw174_[int(18)] = d_1237_v55_
        nw174_[int(19)] = d_1237_v55_
        d_1240_v57_ = nw174_
        r0 = d_1240_v57_
        r1 = d_1231_v49_
        d_1241_v58_: _dafny.Map
        d_1241_v58_ = _dafny.Map({True: (self).f27})
        d_1242_v59_: _dafny.Array
        nw175_ = _dafny.Array(_dafny.Array(None, 0), 15)
        d_1242_v59_ = nw175_
        d_1243_v60_: _dafny.Map
        d_1243_v60_ = _dafny.Map({((d_1241_v58_)[(self).f34] if ((self).f34) in (d_1241_v58_) else (self).f34): d_1242_v59_})
        r2 = ((d_1243_v60_)[((d_1210_v32_)[default__.safeIndex(708, (d_1210_v32_).length(0))]) >= (default__.fm1((self).f27, d_1232_v50_, (self).f27, globalState))] if (((d_1210_v32_)[default__.safeIndex(708, (d_1210_v32_).length(0))]) >= (default__.fm1((self).f27, d_1232_v50_, (self).f27, globalState))) in (d_1243_v60_) else d_1242_v59_)
        d_1244_v61_: _dafny.Map
        d_1244_v61_ = _dafny.Map({d_1176_v4_: not ((self).f27) or ((self).f27)})
        d_1245_v62_: _dafny.Map
        d_1245_v62_ = _dafny.Map({default__.fm1((self).f27, d_1232_v50_, (self).f27, globalState): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "voj"))})
        r3 = ((d_1244_v61_)[len(((d_1245_v62_)[d_1176_v4_] if (d_1176_v4_) in (d_1245_v62_) else d_1231_v49_))] if (len(((d_1245_v62_)[d_1176_v4_] if (d_1176_v4_) in (d_1245_v62_) else d_1231_v49_))) in (d_1244_v61_) else (self).f27)
        return r0, r1, r2, r3

    def m7(self, p0, p1, p2, p3, globalState):
        r0: _dafny.MultiSet = _dafny.MultiSet({})
        r1: bool = False
        r2: int = int(0)
        d_1246_v0_: D0
        d_1246_v0_ = D0_DC1(False)
        d_1247_v1_: _dafny.MultiSet
        d_1247_v1_ = _dafny.MultiSet([(self).f34, (self).fm4(d_1246_v0_, globalState), (self).f27, True, (p0) > (592)])
        r0 = d_1247_v1_
        d_1248_v2_: _dafny.Map
        d_1248_v2_ = _dafny.Map({D15_DC33(): (len((self).fm3(globalState))) > (p0)})
        d_1249_v3_: D15
        d_1249_v3_ = D15_DC33()
        d_1248_v2_ = (d_1248_v2_).set(d_1249_v3_, p3)
        d_1250_v4_: _dafny.Seq
        d_1250_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dtgjoqmwc"))
        if default__.fm0(d_1250_v4_, (self).f27, (self).f34, _dafny.Set({p0, p0, (self).fm9(p0, globalState), p0}), globalState):
            d_1251_v5_: _dafny.Array
            nw176_ = _dafny.Array(None, 8)
            nw176_[int(0)] = (self).f34
            nw176_[int(1)] = (self).f34
            nw176_[int(2)] = p3
            nw176_[int(3)] = (True) == (p3)
            nw176_[int(4)] = not(p3)
            nw176_[int(5)] = (d_1250_v4_) <= (d_1250_v4_)
            nw176_[int(6)] = p3
            nw176_[int(7)] = (p3) or ((self).f27)
            d_1251_v5_ = nw176_
            d_1251_v5_ = d_1251_v5_
            if p3:
                d_1252_v6_: _dafny.Seq
                d_1252_v6_ = _dafny.SeqWithoutIsStrInference([default__.fm21(p0, globalState), default__.fm21(p0, globalState)])
                d_1253_v7_: _dafny.Seq
                d_1253_v7_ = _dafny.SeqWithoutIsStrInference([p3, (D3_DC9((self).f27, p0, p0)).cf14])
                d_1254_v8_: _dafny.Set
                d_1254_v8_ = _dafny.Set({p0})
                d_1252_v6_ = ((d_1252_v6_).set(default__.safeIndex(len(d_1253_v7_), len(d_1252_v6_)), d_1254_v8_) if ((self).f34 if not(p3) else (self).f34) else _dafny.SeqWithoutIsStrInference([(D18_DC41(d_1254_v8_)).cf56 for d_1255_i0_ in range(default__.abs(59))]))
                (globalState).f2 = (0) - (default__.safeModuloInt(p0, 532))
                d_1250_v4_ = (self).fm3(globalState)
                d_1256_v9_: _dafny.Array
                nw177_ = _dafny.Array(_dafny.Array(None, 0), 22)
                d_1256_v9_ = nw177_
                d_1257_v10_: _dafny.Map
                d_1257_v10_ = _dafny.Map({p0: p0})
                d_1258_v11_: _dafny.Array
                nw178_ = _dafny.Array(None, 21)
                nw178_[int(0)] = p0
                nw178_[int(1)] = p0
                nw178_[int(2)] = p0
                nw178_[int(3)] = 343
                nw178_[int(4)] = 999
                nw178_[int(5)] = p0
                nw178_[int(6)] = p0
                nw178_[int(7)] = (_dafny.MultiSet(d_1250_v4_)).cardinality
                nw178_[int(8)] = p0
                nw178_[int(9)] = p0
                nw178_[int(10)] = (0) - (len(d_1257_v10_))
                nw178_[int(11)] = 373
                nw178_[int(12)] = -278
                nw178_[int(13)] = p0
                nw178_[int(14)] = (_dafny.MultiSet([(self).f34, (self).f34])).cardinality
                nw178_[int(15)] = p0
                nw178_[int(16)] = p0
                nw178_[int(17)] = 285
                nw178_[int(18)] = p0
                nw178_[int(19)] = p0
                nw178_[int(20)] = p0
                d_1258_v11_ = nw178_
                d_1259_v12_: _dafny.Array
                d_1259_v12_ = d_1258_v11_
                index177_ = default__.safeIndex(630, (d_1256_v9_).length(0))
                (d_1256_v9_)[index177_] = d_1259_v12_
                index178_ = default__.safeIndex(630, (d_1256_v9_).length(0))
                (d_1256_v9_)[index178_] = (d_1258_v11_ if (p3 if default__.fm0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ssmo")), p3, p3, d_1254_v8_, globalState) else (self).f27) else d_1259_v12_)
                (globalState).f2 = p0
            elif True:
                d_1260_v13_: str
                d_1260_v13_ = _dafny.CodePoint('c')
                d_1260_v13_ = default__.fm14(p0, globalState)
                d_1250_v4_ = (((d_1250_v4_).set(default__.safeIndex(p0, len(d_1250_v4_)), d_1260_v13_)) + (d_1250_v4_)) + ((d_1250_v4_).set(default__.safeIndex(-792, len(d_1250_v4_)), _dafny.CodePoint('j')))
                d_1261_v14_: _dafny.Seq
                d_1261_v14_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1261_v14_ = (d_1261_v14_) + ((d_1261_v14_) + (_dafny.SeqWithoutIsStrInference([p0, 270])))
                (globalState).f21 = (self).f34
                (globalState).f2 = p0
            d_1262_v15_: _dafny.Array
            nw179_ = _dafny.Array(int(0), 2)
            d_1262_v15_ = nw179_
            index179_ = default__.safeIndex(398, (d_1262_v15_).length(0))
            (d_1262_v15_)[index179_] = (p0) - (p0)
            index180_ = default__.safeIndex(398, (d_1262_v15_).length(0))
            (d_1262_v15_)[index180_] = ((d_1247_v1_)[(self).f34] if ((self).f34) in (d_1247_v1_) else len(d_1250_v4_))
            d_1263_v16_: _dafny.Set
            d_1263_v16_ = _dafny.Set({604})
            if default__.fm0(d_1250_v4_, ((self).f27) and ((self).f27), (self).f34, d_1263_v16_, globalState):
                d_1264_v17_: D17
                d_1264_v17_ = D17_DC39(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_1265_i1_ in range(default__.abs(830))]))
                d_1250_v4_ = (d_1264_v17_).cf54
                d_1266_v18_: str
                d_1266_v18_ = _dafny.CodePoint('w')
                d_1266_v18_ = d_1266_v18_
                d_1267_v19_: _dafny.MultiSet
                d_1267_v19_ = _dafny.MultiSet([d_1266_v18_])
                d_1268_v20_: _dafny.Map
                d_1268_v20_ = _dafny.Map({(self).f34: d_1267_v19_})
                index181_ = default__.safeIndex(398, (d_1262_v15_).length(0))
                (d_1262_v15_)[index181_] = ((D19_DC45(((d_1268_v20_)[False] if (False) in (d_1268_v20_) else _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(d_1250_v4_)[default__.safeIndex((d_1262_v15_)[default__.safeIndex(398, (d_1262_v15_).length(0))], len(d_1250_v4_))], d_1266_v18_, d_1266_v18_]))))).cf63).cardinality
                d_1269_v21_: _dafny.Array
                nw180_ = _dafny.Array(_dafny.Map({}), 14)
                d_1269_v21_ = nw180_
                index182_ = default__.safeIndex(930, (d_1269_v21_).length(0))
                (d_1269_v21_)[index182_] = _dafny.Map({p3: p3})
                d_1270_v22_: _dafny.Map
                d_1270_v22_ = _dafny.Map({p3: (self).f34})
                index183_ = default__.safeIndex(930, (d_1269_v21_).length(0))
                (d_1269_v21_)[index183_] = ((d_1270_v22_) | (_dafny.Map({(self).f34: p3}))) | ((_dafny.Map({p3: not(False)})) | (_dafny.Map({default__.fm0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fdvsrg")), p3, not((self).f34), d_1263_v16_, globalState): (self).f27})))
                d_1271_v23_: _dafny.Map
                d_1271_v23_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdvnyqfew"))): p0})
                d_1272_v24_: _dafny.Map
                d_1272_v24_ = _dafny.Map({d_1271_v23_: (d_1250_v4_) + (d_1250_v4_)})
                d_1273_v25_: D18
                d_1273_v25_ = D18_DC42(d_1250_v4_, (self).f34)
                d_1272_v24_ = (d_1272_v24_).set(d_1271_v23_, (d_1273_v25_).cf57)
            elif True:
                d_1274_v26_: C1
                nw181_ = C1()
                nw181_.ctor__((self).f27)
                d_1274_v26_ = nw181_
                (globalState).f6 = (d_1262_v15_)[default__.safeIndex(398, (d_1262_v15_).length(0))]
                index184_ = default__.safeIndex(398, (d_1262_v15_).length(0))
                (d_1262_v15_)[index184_] = (d_1262_v15_)[default__.safeIndex(398, (d_1262_v15_).length(0))]
                (globalState).f6 = (d_1262_v15_)[default__.safeIndex(398, (d_1262_v15_).length(0))]
                d_1275_v27_: C0
                nw182_ = C0()
                nw182_.ctor__()
                d_1275_v27_ = nw182_
            d_1276_v28_: _dafny.Seq
            d_1276_v28_ = _dafny.SeqWithoutIsStrInference([d_1251_v5_, d_1251_v5_, d_1251_v5_, d_1251_v5_, d_1251_v5_])
            d_1277_v29_: _dafny.Map
            d_1277_v29_ = _dafny.Map({(d_1262_v15_)[default__.safeIndex(398, (d_1262_v15_).length(0))]: d_1251_v5_})
            d_1278_v30_: _dafny.Map
            d_1278_v30_ = _dafny.Map({(self).f34: (d_1276_v28_)[default__.safeIndex((d_1262_v15_)[default__.safeIndex(398, (d_1262_v15_).length(0))], len(d_1276_v28_))]})
            d_1279_v31_: _dafny.Array
            nw183_ = _dafny.Array(None, 25)
            nw183_[int(0)] = d_1251_v5_
            nw183_[int(1)] = d_1251_v5_
            nw183_[int(2)] = d_1251_v5_
            nw183_[int(3)] = d_1251_v5_
            nw183_[int(4)] = d_1251_v5_
            nw183_[int(5)] = d_1251_v5_
            nw183_[int(6)] = d_1251_v5_
            nw183_[int(7)] = d_1251_v5_
            nw183_[int(8)] = d_1251_v5_
            nw183_[int(9)] = d_1251_v5_
            nw183_[int(10)] = d_1251_v5_
            nw183_[int(11)] = d_1251_v5_
            nw183_[int(12)] = (d_1276_v28_)[default__.safeIndex(534, len(d_1276_v28_))]
            nw183_[int(13)] = d_1251_v5_
            nw183_[int(14)] = d_1251_v5_
            nw183_[int(15)] = d_1251_v5_
            nw183_[int(16)] = d_1251_v5_
            nw183_[int(17)] = d_1251_v5_
            nw183_[int(18)] = d_1251_v5_
            nw183_[int(19)] = ((d_1277_v29_)[(d_1262_v15_)[default__.safeIndex(398, (d_1262_v15_).length(0))]] if ((d_1262_v15_)[default__.safeIndex(398, (d_1262_v15_).length(0))]) in (d_1277_v29_) else ((d_1278_v30_)[(self).f34] if ((self).f34) in (d_1278_v30_) else d_1251_v5_))
            nw183_[int(20)] = d_1251_v5_
            nw183_[int(21)] = d_1251_v5_
            nw183_[int(22)] = d_1251_v5_
            nw183_[int(23)] = d_1251_v5_
            nw183_[int(24)] = d_1251_v5_
            d_1279_v31_ = nw183_
            index185_ = default__.safeIndex(554, (d_1279_v31_).length(0))
            (d_1279_v31_)[index185_] = d_1251_v5_
            index186_ = default__.safeIndex(554, (d_1279_v31_).length(0))
            (d_1279_v31_)[index186_] = d_1251_v5_
        elif True:
            d_1280_v32_: _dafny.Array
            nw184_ = _dafny.Array(False, 27)
            d_1280_v32_ = nw184_
            d_1281_v33_: _dafny.Seq
            d_1281_v33_ = _dafny.SeqWithoutIsStrInference([p3, p3, (self).f34])
            nw185_ = _dafny.Array(None, 6)
            nw185_[int(0)] = (self).f27
            nw185_[int(1)] = (self).f34
            nw185_[int(2)] = (d_1281_v33_)[default__.safeIndex(len(_dafny.Map({False: p3})), len(d_1281_v33_))]
            nw185_[int(3)] = (self).f27
            nw185_[int(4)] = True
            nw185_[int(5)] = (self).fm4(d_1246_v0_, globalState)
            d_1280_v32_ = nw185_
            d_1282_v34_: _dafny.Seq
            d_1282_v34_ = _dafny.SeqWithoutIsStrInference([d_1281_v33_])
            (globalState).f2 = (p0) * (len(d_1282_v34_))
            d_1280_v32_ = d_1280_v32_
            d_1283_v35_: str
            d_1283_v35_ = _dafny.CodePoint('n')
            d_1284_v36_: _dafny.Array
            nw186_ = _dafny.Array(None, 11)
            nw186_[int(0)] = (d_1250_v4_) + (d_1250_v4_)
            nw186_[int(1)] = d_1250_v4_
            nw186_[int(2)] = (d_1250_v4_) + (d_1250_v4_)
            nw186_[int(3)] = d_1250_v4_
            nw186_[int(4)] = d_1250_v4_
            nw186_[int(5)] = d_1250_v4_
            nw186_[int(6)] = (d_1250_v4_) + (d_1250_v4_)
            nw186_[int(7)] = d_1250_v4_
            nw186_[int(8)] = ((self).fm3(globalState)).set(default__.safeIndex(p0, len((self).fm3(globalState))), d_1283_v35_)
            nw186_[int(9)] = (d_1250_v4_ if (self).f34 else d_1250_v4_)
            nw186_[int(10)] = (d_1250_v4_) + (d_1250_v4_)
            d_1284_v36_ = nw186_
            def lambda56_(d_1285_v4_, d_1286_v35_):
                def lambda57_(d_1287_i2_):
                    return ((d_1285_v4_).set(default__.safeIndex(846, len(d_1285_v4_)), d_1286_v35_)) + (_dafny.SeqWithoutIsStrInference([d_1286_v35_ for d_1288_i3_ in range(default__.abs(934))]))

                return lambda57_

            init27_ = lambda56_(d_1250_v4_, d_1283_v35_)
            nw187_ = _dafny.Array(None, 27)
            for i0_27_ in range(nw187_.length(0)):
                nw187_[i0_27_] = init27_(i0_27_)
            d_1284_v36_ = nw187_
            d_1289_v37_: C0
            nw188_ = C0()
            nw188_.ctor__()
            d_1289_v37_ = nw188_
        d_1290_v38_: _dafny.Array
        def lambda58_(d_1291_p0_):
            def lambda59_(d_1292_i4_):
                return (d_1292_i4_) - (d_1291_p0_)

            return lambda59_

        init28_ = lambda58_(p0)
        nw189_ = _dafny.Array(None, 23)
        for i0_28_ in range(nw189_.length(0)):
            nw189_[i0_28_] = init28_(i0_28_)
        d_1290_v38_ = nw189_
        d_1293_v39_: _dafny.Set
        d_1293_v39_ = _dafny.Set({d_1290_v38_})
        d_1294_v40_: D6
        d_1294_v40_ = D6_DC14((self).f27, p0)
        d_1293_v39_ = (d_1293_v39_ if (d_1294_v40_).cf22 else (d_1293_v39_) - (_dafny.Set({d_1290_v38_})))
        d_1295_v41_: _dafny.Set
        d_1295_v41_ = _dafny.Set({not ((self).f27) or (True), (D9_DC21((self).f27, p0)).cf31})
        d_1296_v42_: T2
        nw190_ = C1()
        nw190_.ctor__(not(p3))
        d_1296_v42_ = nw190_
        d_1297_v43_: _dafny.Array
        nw191_ = _dafny.Array(None, 1)
        nw191_[int(0)] = d_1296_v42_
        d_1297_v43_ = nw191_
        index187_ = default__.safeIndex(473, (d_1297_v43_).length(0))
        (d_1297_v43_)[index187_] = d_1296_v42_
        d_1298_v44_: C1
        nw192_ = C1()
        nw192_.ctor__(p3)
        d_1298_v44_ = nw192_
        d_1299_v45_: _dafny.Seq
        d_1299_v45_ = _dafny.SeqWithoutIsStrInference([d_1298_v44_])
        d_1300_v46_: _dafny.Map
        d_1300_v46_ = _dafny.Map({(d_1299_v45_) != (_dafny.SeqWithoutIsStrInference([d_1298_v44_, d_1298_v44_])): (d_1296_v42_).f27})
        d_1301_v47_: _dafny.Array
        nw193_ = _dafny.Array(_dafny.CodePoint('D'), 19)
        d_1301_v47_ = nw193_
        d_1302_v48_: str
        d_1302_v48_ = _dafny.CodePoint('q')
        index188_ = default__.safeIndex(548, (d_1301_v47_).length(0))
        (d_1301_v47_)[index188_] = d_1302_v48_
        index189_ = default__.safeIndex(473, (d_1297_v43_).length(0))
        index190_ = default__.safeIndex(548, (d_1301_v47_).length(0))
        rhs178_ = (_dafny.Set({(self).f34, (self).f27, (self).f27})) - ((d_1295_v41_).intersection(_dafny.Set({p3, (self).f27, p3, (d_1296_v42_).f27})))
        rhs179_ = (len(d_1250_v4_) if True else p0)
        rhs180_ = d_1296_v42_
        rhs181_ = d_1300_v46_
        rhs182_ = d_1302_v48_
        lhs137_ = globalState
        lhs138_ = d_1297_v43_
        lhs139_ = default__.safeIndex(473, (d_1297_v43_).length(0))
        lhs140_ = d_1301_v47_
        lhs141_ = default__.safeIndex(548, (d_1301_v47_).length(0))
        d_1295_v41_ = rhs178_
        lhs137_.f6 = rhs179_
        lhs138_[lhs139_] = rhs180_
        d_1300_v46_ = rhs181_
        lhs140_[lhs141_] = rhs182_
        d_1303_v49_: _dafny.MultiSet
        d_1303_v49_ = _dafny.MultiSet([p0, p0, 215])
        d_1303_v49_ = d_1303_v49_
        r0 = d_1247_v1_
        r1 = False
        r2 = len(d_1250_v4_)
        return r0, r1, r2

    @property
    def f34(self):
        return self._f34

class C3(T0):
    def  __init__(self):
        self._f27: bool = False
        self.f38: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f27(self):
        return self._f27
    def ctor__(self, f38, f27):
        (self).f38 = f38
        (self)._f27 = f27

    def fm2(self, p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_1304_i0_ in range(default__.abs(741))])

    def fm3(self, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aekmyum"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kfyqxtw")))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "taj")) if self.f38 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dvadpbsyb"))))

    def m1(self, globalState):
        r0: int = int(0)
        d_1305_v0_: _dafny.Set
        d_1305_v0_ = _dafny.Set({self.f38})
        (globalState).f25 = ((_dafny.Set({self.f38, True})).intersection(d_1305_v0_)).isdisjoint((d_1305_v0_).intersection(d_1305_v0_))
        d_1306_v1_: _dafny.Array
        nw194_ = _dafny.Array(False, 18)
        d_1306_v1_ = nw194_
        d_1307_v2_: int
        d_1307_v2_ = 742
        d_1308_v3_: D13
        d_1308_v3_ = D13_DC29(d_1307_v2_, d_1307_v2_, self.f38)
        d_1309_v4_: _dafny.Map
        d_1309_v4_ = _dafny.Map({d_1307_v2_: (d_1308_v3_).cf41})
        index191_ = default__.safeIndex(646, (d_1306_v1_).length(0))
        (d_1306_v1_)[index191_] = (not(self.f38) if self.f38 else ((d_1309_v4_)[198] if (198) in (d_1309_v4_) else (self).f27))
        index192_ = default__.safeIndex(646, (d_1306_v1_).length(0))
        (d_1306_v1_)[index192_] = False
        d_1310_v5_: _dafny.Array
        def lambda60_(d_1311_v1_):
            def lambda61_(d_1312_i0_):
                return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_1313_i1_ in range(default__.abs(189))]) if (d_1311_v1_)[default__.safeIndex(646, (d_1311_v1_).length(0))] else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "msfdrda")))

            return lambda61_

        init29_ = lambda60_(d_1306_v1_)
        nw195_ = _dafny.Array(None, 7)
        for i0_29_ in range(nw195_.length(0)):
            nw195_[i0_29_] = init29_(i0_29_)
        d_1310_v5_ = nw195_
        index193_ = default__.safeIndex(115, (d_1310_v5_).length(0))
        (d_1310_v5_)[index193_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oenhlrqyn"))
        d_1314_v6_: _dafny.Seq
        d_1314_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "orexglwba"))
        index194_ = default__.safeIndex(115, (d_1310_v5_).length(0))
        (d_1310_v5_)[index194_] = d_1314_v6_
        d_1315_v7_: _dafny.Map
        d_1315_v7_ = _dafny.Map({(self).f27: 860})
        d_1316_v8_: _dafny.Map
        d_1316_v8_ = _dafny.Map({d_1315_v7_: (d_1307_v2_ if self.f38 else d_1307_v2_)})
        d_1316_v8_ = (d_1316_v8_).set(d_1315_v7_, default__.safeDivisionInt(d_1307_v2_, d_1307_v2_))
        d_1317_v9_: D0
        d_1317_v9_ = D0_DC1((d_1306_v1_)[default__.safeIndex(646, (d_1306_v1_).length(0))])
        d_1318_v10_: D0
        d_1318_v10_ = D0_DC3(d_1317_v9_)
        d_1318_v10_ = d_1318_v10_
        d_1319_v11_: bool
        d_1320_v12_: _dafny.Seq
        out62_: bool
        out63_: _dafny.Seq
        out62_, out63_ = default__.m0(D0_DC3(D0_DC1((d_1306_v1_)[default__.safeIndex(646, (d_1306_v1_).length(0))])), D0_DC1(self.f38), globalState)
        d_1319_v11_ = out62_
        d_1320_v12_ = out63_
        r0 = d_1307_v2_
        return r0

    def m2(self, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: bool = False
        d_1321_v0_: int
        d_1321_v0_ = -919
        (globalState).f6 = default__.safeModuloInt(default__.safeModuloInt(d_1321_v0_, d_1321_v0_), d_1321_v0_)
        d_1322_v1_: _dafny.Array
        def lambda62_(d_1323_i0_):
            return not(not(self.f38))

        init30_ = lambda62_
        nw196_ = _dafny.Array(None, 22)
        for i0_30_ in range(nw196_.length(0)):
            nw196_[i0_30_] = init30_(i0_30_)
        d_1322_v1_ = nw196_
        d_1324_v2_: str
        d_1324_v2_ = _dafny.CodePoint('q')
        d_1325_v3_: _dafny.Map
        d_1325_v3_ = _dafny.Map({d_1321_v0_: d_1324_v2_})
        d_1326_v4_: _dafny.Set
        d_1326_v4_ = _dafny.Set({(d_1324_v2_ if self.f38 else ((d_1325_v3_)[d_1321_v0_] if (d_1321_v0_) in (d_1325_v3_) else _dafny.CodePoint('x'))), _dafny.CodePoint('i'), d_1324_v2_})
        rhs183_ = (0) - ((0) - (len(d_1326_v4_)))
        rhs184_ = d_1322_v1_
        d_1321_v0_ = rhs183_
        d_1322_v1_ = rhs184_
        d_1327_v5_: _dafny.Seq
        d_1327_v5_ = _dafny.SeqWithoutIsStrInference([self.f38, (self).f27])
        d_1328_v6_: _dafny.Map
        d_1328_v6_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(0) - (d_1321_v0_), d_1321_v0_, 796]): default__.fm1(self.f38, d_1327_v5_, self.f38, globalState)})
        d_1329_v7_: D3
        d_1329_v7_ = D3_DC8(d_1328_v6_)
        d_1330_v8_: _dafny.MultiSet
        d_1330_v8_ = _dafny.MultiSet([d_1321_v0_])
        d_1331_v9_: _dafny.Map
        d_1331_v9_ = _dafny.Map({d_1329_v7_: d_1330_v8_})
        source25_ = D20_DC48(d_1331_v9_)
        if source25_.is_DC49:
            d_1332___mcc_h0_ = source25_.cf69
            d_1333___mcc_h1_ = source25_.cf70
            d_1334___mcc_h2_ = source25_.cf71
            d_1335___mcc_h3_ = source25_.cf72
            d_1336___mcc_h4_ = source25_.cf73
            d_1337_cf73_ = d_1336___mcc_h4_
            d_1338_cf72_ = d_1335___mcc_h3_
            d_1339_cf71_ = d_1334___mcc_h2_
            d_1340_cf70_ = d_1333___mcc_h1_
            d_1341_cf69_ = d_1332___mcc_h0_
            d_1342_v10_: _dafny.Seq
            d_1342_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wqwcoqu"))
            d_1343_v11_: _dafny.Set
            d_1343_v11_ = _dafny.Set({(d_1338_cf72_).f27, default__.fm0(d_1342_v10_, True, self.f38, _dafny.Set({d_1339_cf71_, d_1340_cf70_}), globalState)})
            if (d_1343_v11_).isdisjoint(_dafny.Set({self.f38})):
                d_1344_v12_: _dafny.Seq
                d_1344_v12_ = _dafny.SeqWithoutIsStrInference([d_1342_v10_])
                d_1345_v13_: _dafny.Map
                d_1345_v13_ = _dafny.Map({(self).f27: (self).f27})
                d_1346_v14_: _dafny.Seq
                d_1346_v14_ = _dafny.SeqWithoutIsStrInference([len(d_1345_v13_), d_1339_cf71_])
                rhs185_ = (default__.fm1((self).f27, _dafny.SeqWithoutIsStrInference([(self).f27]), self.f38, globalState)) + ((d_1340_cf70_) + ((0) - (len(d_1344_v12_))))
                rhs186_ = (0) - (d_1339_cf71_)
                rhs187_ = ((d_1346_v14_) + (_dafny.SeqWithoutIsStrInference([d_1340_cf70_, d_1321_v0_, d_1340_cf70_]))) == (d_1346_v14_)
                rhs188_ = self.f38
                lhs142_ = globalState
                lhs143_ = globalState
                lhs144_ = globalState
                lhs142_.f6 = rhs185_
                d_1339_cf71_ = rhs186_
                lhs143_.f12 = rhs187_
                lhs144_.f5 = rhs188_
                d_1346_v14_ = d_1346_v14_
                (globalState).f2 = (d_1346_v14_)[default__.safeIndex((d_1321_v0_) * (d_1339_cf71_), len(d_1346_v14_))]
                d_1347_v15_: _dafny.Array
                nw197_ = _dafny.Array(_dafny.Array(None, 0), 10)
                d_1347_v15_ = nw197_
                d_1348_v16_: _dafny.Array
                def lambda63_(d_1349_cf72_):
                    def lambda64_(d_1350_i1_):
                        return _dafny.MultiSet([(d_1349_cf72_).f27, (d_1349_cf72_).f27, (d_1349_cf72_).f27, (self).f27, (d_1349_cf72_).f27])

                    return lambda64_

                init31_ = lambda63_(d_1338_cf72_)
                nw198_ = _dafny.Array(None, 22)
                for i0_31_ in range(nw198_.length(0)):
                    nw198_[i0_31_] = init31_(i0_31_)
                d_1348_v16_ = nw198_
                index195_ = default__.safeIndex(77, (d_1347_v15_).length(0))
                (d_1347_v15_)[index195_] = d_1348_v16_
                index196_ = default__.safeIndex(77, (d_1347_v15_).length(0))
                (d_1347_v15_)[index196_] = d_1348_v16_
                (globalState).f5 = (d_1338_cf72_).f27
            elif True:
                d_1351_v17_: _dafny.Array
                nw199_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 17)
                d_1351_v17_ = nw199_
                index197_ = default__.safeIndex(565, (d_1351_v17_).length(0))
                (d_1351_v17_)[index197_] = d_1342_v10_
                index198_ = default__.safeIndex(565, (d_1351_v17_).length(0))
                (d_1351_v17_)[index198_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xrr")) if (d_1338_cf72_).f27 else (d_1342_v10_) + (d_1342_v10_))
                d_1324_v2_ = _dafny.CodePoint('h')
                d_1352_v18_: _dafny.Array
                def lambda65_(d_1353_cf71_):
                    def lambda66_(d_1354_i2_):
                        return default__.safeDivisionInt(d_1354_i2_, d_1353_cf71_)

                    return lambda66_

                init32_ = lambda65_(d_1339_cf71_)
                nw200_ = _dafny.Array(None, 11)
                for i0_32_ in range(nw200_.length(0)):
                    nw200_[i0_32_] = init32_(i0_32_)
                d_1352_v18_ = nw200_
                d_1355_v19_: _dafny.Set
                d_1355_v19_ = _dafny.Set({d_1352_v18_, d_1352_v18_, d_1352_v18_, d_1352_v18_})
                d_1356_v20_: _dafny.Map
                d_1356_v20_ = _dafny.Map({(_dafny.Set({d_1352_v18_, d_1352_v18_, d_1352_v18_})) | (d_1355_v19_): d_1352_v18_})
                d_1356_v20_ = (d_1356_v20_).set(_dafny.Set({d_1352_v18_, d_1352_v18_}), d_1352_v18_)
                (globalState).f5 = default__.fm0((d_1351_v17_)[default__.safeIndex(565, (d_1351_v17_).length(0))], False, (not((self).f27)) and ((d_1338_cf72_).f27), d_1337_cf73_, globalState)
                (globalState).f5 = (d_1338_cf72_).f27
            d_1357_v21_: _dafny.Map
            d_1357_v21_ = _dafny.Map({d_1330_v8_: self.f38})
            d_1358_v22_: D22
            d_1358_v22_ = D22_DC56(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wjlqmie"))): d_1324_v2_}))
            d_1359_v23_: _dafny.Map
            d_1359_v23_ = _dafny.Map({d_1358_v22_: d_1339_cf71_})
            d_1360_v24_: _dafny.Map
            d_1360_v24_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m")): d_1337_cf73_})
            rhs189_ = default__.safeModuloInt(len(d_1357_v21_), d_1339_cf71_)
            rhs190_ = (len(d_1359_v23_)) > (d_1340_cf70_)
            rhs191_ = (len((d_1360_v24_) | ((D25_DC63(d_1360_v24_)).cf97))) + ((0) - (d_1321_v0_))
            lhs145_ = globalState
            d_1321_v0_ = rhs189_
            lhs145_.f21 = rhs190_
            d_1339_cf71_ = rhs191_
            d_1324_v2_ = (d_1341_cf69_ if (self).f27 else default__.fm44(not(True), self.f38, d_1342_v10_, d_1339_cf71_, globalState))
            (globalState).f2 = d_1339_cf71_
        elif source25_.is_DC50:
            (globalState).f6 = d_1321_v0_
            d_1361_v25_: _dafny.Seq
            d_1361_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hkt"))
            d_1362_v26_: _dafny.MultiSet
            d_1362_v26_ = _dafny.MultiSet([True])
            d_1363_v27_: _dafny.Set
            d_1363_v27_ = _dafny.Set({d_1321_v0_, (d_1362_v26_).cardinality, d_1321_v0_, d_1321_v0_})
            d_1364_v28_: _dafny.Seq
            d_1364_v28_ = _dafny.SeqWithoutIsStrInference([d_1363_v27_])
            d_1365_v29_: _dafny.Map
            d_1365_v29_ = _dafny.Map({(self).f27: default__.fm0(d_1361_v25_, (self).f27, self.f38, (d_1364_v28_)[default__.safeIndex(d_1321_v0_, len(d_1364_v28_))], globalState)})
            d_1365_v29_ = (d_1365_v29_).set(not(True), self.f38)
            d_1366_v30_: _dafny.MultiSet
            d_1366_v30_ = _dafny.MultiSet([d_1322_v1_, d_1322_v1_])
            d_1366_v30_ = d_1366_v30_
            d_1367_v31_: _dafny.Seq
            d_1367_v31_ = _dafny.SeqWithoutIsStrInference([d_1321_v0_])
            d_1368_v32_: _dafny.Map
            d_1368_v32_ = _dafny.Map({d_1327_v5_: d_1361_v25_})
            d_1367_v31_ = (d_1367_v31_ if (d_1368_v32_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([self.f38, (self).f27]): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vhhv"))})) else d_1367_v31_)
        elif source25_.is_DC48:
            d_1369___mcc_h5_ = source25_.cf68
            d_1370_cf68_ = d_1369___mcc_h5_
            d_1371_v33_: _dafny.Map
            d_1371_v33_ = _dafny.Map({d_1321_v0_: d_1322_v1_})
            d_1371_v33_ = (d_1371_v33_).set(default__.safeDivisionInt(d_1321_v0_, d_1321_v0_), d_1322_v1_)
            d_1372_v34_: _dafny.Seq
            d_1372_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))
            r1 = d_1372_v34_
            d_1373_v35_: _dafny.Array
            nw201_ = _dafny.Array(int(0), 23)
            d_1373_v35_ = nw201_
            index199_ = default__.safeIndex(98, (d_1373_v35_).length(0))
            (d_1373_v35_)[index199_] = (d_1321_v0_) * (d_1321_v0_)
            d_1374_v36_: D0
            d_1374_v36_ = D0_DC0(len(d_1372_v34_))
            index200_ = default__.safeIndex(98, (d_1373_v35_).length(0))
            (d_1373_v35_)[index200_] = (d_1374_v36_).cf0
            d_1375_v37_: _dafny.Map
            d_1375_v37_ = _dafny.Map({(d_1373_v35_)[default__.safeIndex(98, (d_1373_v35_).length(0))]: (d_1373_v35_)[default__.safeIndex(98, (d_1373_v35_).length(0))]})
            d_1376_v38_: _dafny.Set
            d_1376_v38_ = _dafny.Set({False, (self).f27})
            d_1377_v39_: _dafny.Seq
            d_1377_v39_ = _dafny.SeqWithoutIsStrInference([(0) - (len(d_1375_v37_)), len(d_1376_v38_)])
            d_1378_v40_: _dafny.Map
            d_1378_v40_ = _dafny.Map({d_1373_v35_: len(d_1377_v39_)})
            d_1378_v40_ = (d_1378_v40_).set(d_1373_v35_, d_1321_v0_)
        elif True:
            d_1379___mcc_h6_ = source25_.cf74
            d_1380_cf74_ = d_1379___mcc_h6_
            d_1381_v41_: _dafny.Array
            nw202_ = _dafny.Array(_dafny.Set({}), 8)
            d_1381_v41_ = nw202_
            d_1381_v41_ = d_1381_v41_
            d_1382_v42_: C2
            nw203_ = C2()
            nw203_.ctor__(False, (self).f27)
            d_1382_v42_ = nw203_
            if (self).f27:
                rhs192_ = d_1321_v0_
                rhs193_ = (0) - ((d_1382_v42_).fm9(d_1321_v0_, globalState))
                rhs194_ = not(not ((self).f27) or ((self).f27))
                rhs195_ = (d_1382_v42_).f34
                lhs146_ = globalState
                lhs147_ = globalState
                lhs148_ = globalState
                lhs149_ = globalState
                lhs146_.f2 = rhs192_
                lhs147_.f2 = rhs193_
                lhs148_.f5 = rhs194_
                lhs149_.f5 = rhs195_
                d_1383_v43_: _dafny.Array
                def lambda67_(d_1384_i3_):
                    return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vet"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vyfmnvnd")))

                init33_ = lambda67_
                nw204_ = _dafny.Array(None, 28)
                for i0_33_ in range(nw204_.length(0)):
                    nw204_[i0_33_] = init33_(i0_33_)
                d_1383_v43_ = nw204_
                d_1385_v44_: _dafny.Seq
                d_1385_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eumqhes"))
                d_1386_v45_: _dafny.Seq
                d_1386_v45_ = _dafny.SeqWithoutIsStrInference([d_1385_v44_])
                index201_ = default__.safeIndex(798, (d_1383_v43_).length(0))
                (d_1383_v43_)[index201_] = (d_1386_v45_)[default__.safeIndex(d_1321_v0_, len(d_1386_v45_))]
                index202_ = default__.safeIndex(798, (d_1383_v43_).length(0))
                (d_1383_v43_)[index202_] = d_1385_v44_
                d_1387_v46_: C2
                nw205_ = C2()
                nw205_.ctor__((self).f27, (652) == ((0) - ((d_1330_v8_).cardinality)))
                d_1387_v46_ = nw205_
                (globalState).f2 = len((d_1383_v43_)[default__.safeIndex(798, (d_1383_v43_).length(0))])
                d_1388_v47_: _dafny.Array
                nw206_ = _dafny.Array(int(0), 23)
                d_1388_v47_ = nw206_
                d_1389_v48_: _dafny.Map
                d_1389_v48_ = _dafny.Map({d_1321_v0_: (d_1382_v42_).f34})
                d_1390_v49_: _dafny.Seq
                d_1390_v49_ = _dafny.SeqWithoutIsStrInference([d_1321_v0_, len((d_1383_v43_)[default__.safeIndex(798, (d_1383_v43_).length(0))]), d_1321_v0_])
                d_1391_v50_: D0
                d_1391_v50_ = D0_DC1((d_1382_v42_).f34)
                d_1392_v51_: D18
                d_1392_v51_ = D18_DC43(d_1321_v0_)
                d_1393_v52_: D19
                d_1393_v52_ = D19_DC46((d_1392_v51_).cf59, d_1388_v47_, (d_1387_v46_).f34)
                d_1394_v53_: _dafny.Map
                d_1394_v53_ = _dafny.Map({d_1387_v46_: d_1321_v0_})
                d_1395_v54_: _dafny.Map
                d_1395_v54_ = _dafny.Map({d_1321_v0_: d_1321_v0_})
                d_1396_v55_: _dafny.MultiSet
                d_1396_v55_ = _dafny.MultiSet([(d_1382_v42_).f34])
                nw207_ = _dafny.Array(None, 26)
                nw207_[int(0)] = d_1321_v0_
                nw207_[int(1)] = d_1321_v0_
                nw207_[int(2)] = d_1321_v0_
                nw207_[int(3)] = d_1321_v0_
                nw207_[int(4)] = (len(d_1389_v48_) if not(not((d_1382_v42_).f34)) else d_1321_v0_)
                nw207_[int(5)] = d_1321_v0_
                nw207_[int(6)] = d_1321_v0_
                nw207_[int(7)] = (d_1390_v49_)[default__.safeIndex(d_1321_v0_, len(d_1390_v49_))]
                nw207_[int(8)] = d_1321_v0_
                nw207_[int(9)] = len(((d_1383_v43_)[default__.safeIndex(798, (d_1383_v43_).length(0))]) + ((d_1386_v45_)[default__.safeIndex(983, len(d_1386_v45_))]))
                nw207_[int(10)] = d_1321_v0_
                nw207_[int(11)] = len((_dafny.SeqWithoutIsStrInference([d_1321_v0_ for d_1397_i4_ in range(default__.abs(528))]) if (d_1387_v46_).fm4(d_1391_v50_, globalState) else _dafny.SeqWithoutIsStrInference([d_1321_v0_, len(d_1327_v5_)])))
                nw207_[int(12)] = d_1321_v0_
                nw207_[int(13)] = d_1321_v0_
                nw207_[int(14)] = (d_1382_v42_).fm9(d_1321_v0_, globalState)
                nw207_[int(15)] = d_1321_v0_
                nw207_[int(16)] = d_1321_v0_
                nw207_[int(17)] = (0) - (d_1321_v0_)
                nw207_[int(18)] = d_1321_v0_
                nw207_[int(19)] = (176) * (d_1321_v0_)
                nw207_[int(20)] = (d_1393_v52_).cf64
                nw207_[int(21)] = d_1321_v0_
                nw207_[int(22)] = ((d_1394_v53_)[d_1387_v46_] if (d_1387_v46_) in (d_1394_v53_) else (0) - (((d_1395_v54_)[(d_1396_v55_).cardinality] if ((d_1396_v55_).cardinality) in (d_1395_v54_) else 972)))
                nw207_[int(23)] = d_1321_v0_
                nw207_[int(24)] = len(d_1395_v54_)
                nw207_[int(25)] = d_1321_v0_
                d_1388_v47_ = nw207_
            elif True:
                d_1398_v56_: _dafny.Map
                d_1398_v56_ = _dafny.Map({d_1322_v1_: 27})
                d_1398_v56_ = ((d_1398_v56_) | (_dafny.Map({d_1322_v1_: d_1321_v0_}))) | (d_1398_v56_)
                index203_ = default__.safeIndex(209, (d_1322_v1_).length(0))
                (d_1322_v1_)[index203_] = True
                d_1399_v57_: _dafny.Map
                d_1399_v57_ = _dafny.Map({d_1321_v0_: d_1322_v1_})
                d_1400_v58_: _dafny.Seq
                d_1400_v58_ = _dafny.SeqWithoutIsStrInference([d_1322_v1_, d_1322_v1_, d_1322_v1_, d_1322_v1_, d_1322_v1_])
                index204_ = default__.safeIndex(209, (d_1322_v1_).length(0))
                rhs196_ = ((d_1399_v57_)[(d_1321_v0_) + (283)] if ((d_1321_v0_) + (283)) in (d_1399_v57_) else (d_1400_v58_)[default__.safeIndex((0) - (d_1321_v0_), len(d_1400_v58_))])
                rhs197_ = d_1324_v2_
                rhs198_ = (d_1382_v42_).f34
                rhs199_ = self.f38
                lhs150_ = globalState
                lhs151_ = d_1322_v1_
                lhs152_ = default__.safeIndex(209, (d_1322_v1_).length(0))
                d_1322_v1_ = rhs196_
                d_1324_v2_ = rhs197_
                lhs150_.f12 = rhs198_
                lhs151_[lhs152_] = rhs199_
                d_1401_v59_: _dafny.Set
                d_1401_v59_ = _dafny.Set({not((d_1382_v42_).f34)})
                (globalState).f25 = (d_1321_v0_) >= (len((d_1401_v59_).intersection(d_1401_v59_)))
                d_1402_v60_: D11
                d_1402_v60_ = D11_DC24(_dafny.Set({(d_1382_v42_).f34}))
                (globalState).f5 = ((d_1401_v59_).issubset((d_1402_v60_).cf35)) and ((d_1382_v42_).f34)
                d_1403_v61_: _dafny.Seq
                d_1403_v61_ = _dafny.SeqWithoutIsStrInference([d_1321_v0_])
                index205_ = default__.safeIndex(209, (d_1322_v1_).length(0))
                (d_1322_v1_)[index205_] = (d_1403_v61_) < ((d_1403_v61_) + (_dafny.SeqWithoutIsStrInference([d_1321_v0_])))
            index206_ = default__.safeIndex(886, (d_1322_v1_).length(0))
            (d_1322_v1_)[index206_] = (self.f38) or ((self).f27)
            index207_ = default__.safeIndex(886, (d_1322_v1_).length(0))
            (d_1322_v1_)[index207_] = self.f38
        d_1404_v62_: D25
        d_1404_v62_ = D25_DC64((self).f27, d_1321_v0_, d_1324_v2_, (self).f27, 744)
        (globalState).f21 = (d_1404_v62_).cf101
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_1322_v1_).length(0)):
            d_1405_i5_: int = guard_loop_3_
            if (True) and (((0) <= (d_1405_i5_)) and ((d_1405_i5_) < ((d_1322_v1_).length(0)))):
                (d_1322_v1_)[(d_1405_i5_)] = (d_1324_v2_) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jrle")))
        d_1406_v63_: _dafny.Seq
        d_1406_v63_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "edngyv"))
        d_1407_v64_: _dafny.Map
        d_1407_v64_ = _dafny.Map({(d_1406_v63_) == (d_1406_v63_): self.f38})
        (globalState).f25 = ((d_1407_v64_)[(self).f27] if ((self).f27) in (d_1407_v64_) else self.f38)
        d_1408_v65_: _dafny.Array
        nw208_ = _dafny.Array(None, 4)
        nw208_[int(0)] = d_1322_v1_
        nw208_[int(1)] = d_1322_v1_
        nw208_[int(2)] = d_1322_v1_
        nw208_[int(3)] = d_1322_v1_
        d_1408_v65_ = nw208_
        r0 = d_1408_v65_
        r1 = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yl"))
        d_1409_v67_: _dafny.Seq
        def iife130_():
            coll62_ = _dafny.Map()
            compr_62_: int
            for compr_62_ in _dafny.IntegerRange(531, 851):
                d_1410_v66_: int = compr_62_
                if ((531) <= (d_1410_v66_)) and ((d_1410_v66_) < (851)):
                    coll62_[default__.safeDivisionInt(d_1410_v66_, d_1321_v0_)] = True
            return _dafny.Map(coll62_)
        d_1409_v67_ = _dafny.SeqWithoutIsStrInference([iife130_()
        ])
        d_1411_v68_: _dafny.Set
        d_1411_v68_ = _dafny.Set({(0) - (d_1321_v0_)})
        d_1412_v69_: _dafny.Map
        d_1412_v69_ = _dafny.Map({d_1321_v0_: d_1411_v68_})
        d_1413_v70_: _dafny.Seq
        d_1413_v70_ = _dafny.SeqWithoutIsStrInference([d_1321_v0_])
        d_1414_v71_: D18
        d_1414_v71_ = D18_DC42(d_1406_v63_, (self).f27)
        d_1415_v72_: _dafny.Array
        nw209_ = _dafny.Array(None, 24)
        nw209_[int(0)] = len((d_1409_v67_)[default__.safeIndex(d_1321_v0_, len(d_1409_v67_))])
        nw209_[int(1)] = d_1321_v0_
        nw209_[int(2)] = d_1321_v0_
        nw209_[int(3)] = d_1321_v0_
        nw209_[int(4)] = d_1321_v0_
        nw209_[int(5)] = d_1321_v0_
        nw209_[int(6)] = d_1321_v0_
        nw209_[int(7)] = d_1321_v0_
        nw209_[int(8)] = d_1321_v0_
        nw209_[int(9)] = d_1321_v0_
        nw209_[int(10)] = d_1321_v0_
        nw209_[int(11)] = len(d_1412_v69_)
        nw209_[int(12)] = d_1321_v0_
        nw209_[int(13)] = d_1321_v0_
        nw209_[int(14)] = 198
        nw209_[int(15)] = d_1321_v0_
        nw209_[int(16)] = len(d_1413_v70_)
        nw209_[int(17)] = len((d_1414_v71_).cf57)
        nw209_[int(18)] = len(d_1413_v70_)
        nw209_[int(19)] = d_1321_v0_
        nw209_[int(20)] = d_1321_v0_
        nw209_[int(21)] = d_1321_v0_
        nw209_[int(22)] = d_1321_v0_
        nw209_[int(23)] = (0) - (d_1321_v0_)
        d_1415_v72_ = nw209_
        d_1416_v73_: _dafny.Array
        d_1416_v73_ = d_1415_v72_
        d_1417_v74_: _dafny.Seq
        d_1417_v74_ = _dafny.SeqWithoutIsStrInference([d_1415_v72_, d_1415_v72_])
        d_1418_v75_: _dafny.Array
        nw210_ = _dafny.Array(None, 28)
        nw210_[int(0)] = d_1415_v72_
        nw210_[int(1)] = d_1415_v72_
        nw210_[int(2)] = d_1415_v72_
        nw210_[int(3)] = d_1415_v72_
        nw210_[int(4)] = d_1415_v72_
        nw210_[int(5)] = (d_1416_v73_)
        nw210_[int(6)] = d_1415_v72_
        nw210_[int(7)] = d_1415_v72_
        nw210_[int(8)] = d_1415_v72_
        nw210_[int(9)] = d_1415_v72_
        nw210_[int(10)] = d_1415_v72_
        nw210_[int(11)] = d_1415_v72_
        nw210_[int(12)] = d_1415_v72_
        nw210_[int(13)] = d_1415_v72_
        nw210_[int(14)] = d_1415_v72_
        nw210_[int(15)] = (d_1417_v74_)[default__.safeIndex(d_1321_v0_, len(d_1417_v74_))]
        nw210_[int(16)] = d_1415_v72_
        nw210_[int(17)] = d_1415_v72_
        nw210_[int(18)] = d_1415_v72_
        nw210_[int(19)] = d_1415_v72_
        nw210_[int(20)] = d_1415_v72_
        nw210_[int(21)] = d_1415_v72_
        nw210_[int(22)] = d_1415_v72_
        nw210_[int(23)] = (d_1415_v72_ if False else d_1415_v72_)
        nw210_[int(24)] = d_1415_v72_
        nw210_[int(25)] = (d_1415_v72_ if self.f38 else d_1415_v72_)
        nw210_[int(26)] = d_1415_v72_
        nw210_[int(27)] = d_1415_v72_
        d_1418_v75_ = nw210_
        r2 = d_1418_v75_
        r3 = default__.fm0((d_1406_v63_ if (self).f27 else d_1406_v63_), self.f38, (self).f27, d_1411_v68_, globalState)
        return r0, r1, r2, r3


class C4(T0):
    def  __init__(self):
        self._f27: bool = False
        self._f37: _dafny.MultiSet = _dafny.MultiSet({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f27(self):
        return self._f27
    def ctor__(self, f37, f27):
        (self)._f37 = f37
        (self)._f27 = f27

    def fm2(self, p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_1419_i0_ in range(default__.abs(803))])

    def fm3(self, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "whubvvj"))

    def fm36(self, p0, p1, globalState):
        return len(((_dafny.Map({(self).f27: (self).f37})) | (_dafny.Map({(self).f27: _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f27, (self).f27]))})) if (self).f27 else (_dafny.Map({(self).f27: (self).f37})) | (_dafny.Map({(self).f27: (self).f37}))))

    def m1(self, globalState):
        r0: int = int(0)
        d_1420_v0_: int
        d_1420_v0_ = 590
        d_1421_v1_: str
        d_1421_v1_ = _dafny.CodePoint('c')
        hi9_ = (_dafny.MultiSet([_dafny.CodePoint('k'), d_1421_v1_])).cardinality
        for d_1422_i0_ in range(d_1420_v0_, hi9_):
            d_1423_v2_: _dafny.Seq
            d_1423_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "a"))
            d_1424_v3_: _dafny.Seq
            d_1424_v3_ = _dafny.SeqWithoutIsStrInference([d_1423_v2_, d_1423_v2_])
            d_1425_v4_: _dafny.Seq
            d_1425_v4_ = _dafny.SeqWithoutIsStrInference([((d_1424_v3_)[default__.safeIndex(d_1422_i0_, len(d_1424_v3_))]) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jgjxccpbs"))), d_1423_v2_, d_1423_v2_])
            d_1425_v4_ = (_dafny.SeqWithoutIsStrInference([d_1423_v2_]) if (self).f27 else _dafny.SeqWithoutIsStrInference([d_1423_v2_]))
            d_1426_v5_: _dafny.Array
            def lambda68_(d_1427_i0_):
                def lambda69_(d_1428_i1_):
                    return (d_1428_i1_) - (d_1427_i0_)

                return lambda69_

            init34_ = lambda68_(d_1422_i0_)
            nw211_ = _dafny.Array(None, 5)
            for i0_34_ in range(nw211_.length(0)):
                nw211_[i0_34_] = init34_(i0_34_)
            d_1426_v5_ = nw211_
            d_1429_v6_: _dafny.Seq
            d_1429_v6_ = _dafny.SeqWithoutIsStrInference([(self).f27])
            index208_ = default__.safeIndex(91, (d_1426_v5_).length(0))
            (d_1426_v5_)[index208_] = default__.fm1((self).f27, d_1429_v6_, (self).f27, globalState)
            index209_ = default__.safeIndex(91, (d_1426_v5_).length(0))
            (d_1426_v5_)[index209_] = d_1422_i0_
            d_1430_v7_: _dafny.Array
            nw212_ = _dafny.Array(_dafny.CodePoint('D'), 9)
            d_1430_v7_ = nw212_
            index210_ = default__.safeIndex(152, (d_1430_v7_).length(0))
            (d_1430_v7_)[index210_] = d_1421_v1_
            index211_ = default__.safeIndex(91, (d_1426_v5_).length(0))
            index212_ = default__.safeIndex(152, (d_1430_v7_).length(0))
            rhs200_ = d_1422_i0_
            rhs201_ = not(((45) + (d_1422_i0_)) > ((d_1422_i0_) + (d_1420_v0_)))
            rhs202_ = d_1421_v1_
            lhs153_ = d_1426_v5_
            lhs154_ = default__.safeIndex(91, (d_1426_v5_).length(0))
            lhs155_ = globalState
            lhs156_ = d_1430_v7_
            lhs157_ = default__.safeIndex(152, (d_1430_v7_).length(0))
            lhs153_[lhs154_] = rhs200_
            lhs155_.f21 = rhs201_
            lhs156_[lhs157_] = rhs202_
            index213_ = default__.safeIndex(91, (d_1426_v5_).length(0))
            (d_1426_v5_)[index213_] = ((d_1426_v5_)[default__.safeIndex(91, (d_1426_v5_).length(0))]) - ((0) - ((0) - (d_1420_v0_)))
        d_1431_v8_: _dafny.Map
        d_1431_v8_ = _dafny.Map({(self).f27: (self).f27})
        d_1432_v9_: _dafny.MultiSet
        d_1432_v9_ = _dafny.MultiSet([d_1420_v0_, 5])
        d_1433_i2_: int
        d_1433_i2_ = 0
        with _dafny.label("8"):
            while ((len(d_1431_v8_)) not in (d_1432_v9_)) not in ((self).f37):
                with _dafny.c_label("8"):
                    if (d_1433_i2_) >= (100):
                        raise _dafny.Break("8")
                    d_1433_i2_ = (d_1433_i2_) + (1)
                    d_1434_v10_: C1
                    nw213_ = C1()
                    nw213_.ctor__((self).f27)
                    d_1434_v10_ = nw213_
                    if (self).f27:
                        d_1435_v11_: _dafny.Array
                        nw214_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 7)
                        d_1435_v11_ = nw214_
                        d_1436_v12_: _dafny.Seq
                        d_1436_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gg"))
                        index214_ = default__.safeIndex(580, (d_1435_v11_).length(0))
                        (d_1435_v11_)[index214_] = d_1436_v12_
                        index215_ = default__.safeIndex(580, (d_1435_v11_).length(0))
                        (d_1435_v11_)[index215_] = d_1436_v12_
                        d_1437_v13_: _dafny.Map
                        d_1437_v13_ = _dafny.Map({(d_1435_v11_)[default__.safeIndex(580, (d_1435_v11_).length(0))]: d_1420_v0_})
                        d_1437_v13_ = (d_1437_v13_).set((d_1435_v11_)[default__.safeIndex(580, (d_1435_v11_).length(0))], len(_dafny.Map({897: True})))
                        d_1438_v14_: C1
                        nw215_ = C1()
                        nw215_.ctor__((self).f27)
                        d_1438_v14_ = nw215_
                        d_1439_v15_: _dafny.Seq
                        d_1439_v15_ = _dafny.SeqWithoutIsStrInference([(self).f27])
                        d_1440_v16_: D0
                        d_1440_v16_ = D0_DC1((self).f27)
                        d_1441_v17_: C2
                        nw216_ = C2()
                        nw216_.ctor__((d_1439_v15_)[default__.safeIndex(len(_dafny.Map({(d_1438_v14_).fm4(d_1440_v16_, globalState): (self).f27})), len(d_1439_v15_))], (self).f27)
                        d_1441_v17_ = nw216_
                        rhs203_ = d_1441_v17_
                        rhs204_ = (d_1441_v17_).f34
                        rhs205_ = d_1441_v17_
                        rhs206_ = not ((self).f27) or ((self).f27)
                        lhs158_ = globalState
                        lhs159_ = globalState
                        d_1441_v17_ = rhs203_
                        lhs158_.f12 = rhs204_
                        d_1441_v17_ = rhs205_
                        lhs159_.f12 = rhs206_
                        d_1442_v18_: _dafny.Array
                        nw217_ = _dafny.Array(_dafny.CodePoint('D'), 21)
                        d_1442_v18_ = nw217_
                        index216_ = default__.safeIndex(946, (d_1442_v18_).length(0))
                        (d_1442_v18_)[index216_] = d_1421_v1_
                        d_1443_v19_: _dafny.Set
                        d_1443_v19_ = _dafny.Set({d_1431_v8_, (default__.fm38(d_1420_v0_, d_1420_v0_, globalState)).set((d_1441_v17_).f34, (d_1441_v17_).f34), d_1431_v8_, d_1431_v8_})
                        d_1444_v20_: _dafny.Map
                        d_1444_v20_ = _dafny.Map({d_1421_v1_: (d_1441_v17_).f34})
                        d_1445_v21_: _dafny.Seq
                        d_1445_v21_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_1421_v1_: not((self).f27)}), d_1444_v20_, _dafny.Map({_dafny.CodePoint('i'): (self).f27})])
                        index217_ = default__.safeIndex(946, (d_1442_v18_).length(0))
                        rhs207_ = default__.fm37(d_1443_v19_, d_1445_v21_, globalState)
                        rhs208_ = (self).f27
                        rhs209_ = (d_1420_v0_) * (-203)
                        rhs210_ = ((default__.fm1((d_1441_v17_).f34, d_1439_v15_, (self).f27, globalState)) * (156)) == (default__.safeModuloInt(-584, d_1420_v0_))
                        rhs211_ = (d_1441_v17_).f34
                        lhs160_ = d_1442_v18_
                        lhs161_ = default__.safeIndex(946, (d_1442_v18_).length(0))
                        lhs162_ = globalState
                        lhs163_ = globalState
                        lhs164_ = globalState
                        lhs165_ = globalState
                        lhs160_[lhs161_] = rhs207_
                        lhs162_.f14 = rhs208_
                        lhs163_.f2 = rhs209_
                        lhs164_.f5 = rhs210_
                        lhs165_.f12 = rhs211_
                    elif True:
                        d_1446_v22_: _dafny.Seq
                        d_1446_v22_ = _dafny.SeqWithoutIsStrInference([d_1420_v0_, (0) - (d_1420_v0_), d_1420_v0_])
                        rhs212_ = -730
                        rhs213_ = d_1420_v0_
                        rhs214_ = (self).f27
                        rhs215_ = (self).f27
                        rhs216_ = default__.safeDivisionInt(((d_1446_v22_)[default__.safeIndex(d_1420_v0_, len(d_1446_v22_))]) - (d_1420_v0_), d_1420_v0_)
                        lhs166_ = globalState
                        lhs167_ = globalState
                        lhs168_ = globalState
                        lhs169_ = globalState
                        r0 = rhs212_
                        lhs166_.f6 = rhs213_
                        lhs167_.f25 = rhs214_
                        lhs168_.f25 = rhs215_
                        lhs169_.f6 = rhs216_
                        d_1447_v23_: _dafny.Seq
                        d_1447_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))
                        d_1448_v24_: _dafny.Seq
                        d_1448_v24_ = _dafny.SeqWithoutIsStrInference([d_1447_v23_, d_1447_v23_, d_1447_v23_])
                        d_1449_v25_: D13
                        d_1449_v25_ = D13_DC27(d_1448_v24_)
                        d_1449_v25_ = d_1449_v25_
                        d_1450_v26_: _dafny.Array
                        def lambda70_(d_1451_v0_):
                            def lambda71_(d_1452_i3_):
                                return default__.safeDivisionInt(d_1452_i3_, d_1451_v0_)

                            return lambda71_

                        init35_ = lambda70_(d_1420_v0_)
                        nw218_ = _dafny.Array(None, 20)
                        for i0_35_ in range(nw218_.length(0)):
                            nw218_[i0_35_] = init35_(i0_35_)
                        d_1450_v26_ = nw218_
                        index218_ = default__.safeIndex(46, (d_1450_v26_).length(0))
                        (d_1450_v26_)[index218_] = d_1420_v0_
                        index219_ = default__.safeIndex(46, (d_1450_v26_).length(0))
                        (d_1450_v26_)[index219_] = (d_1420_v0_) * (d_1420_v0_)
                        d_1453_v27_: _dafny.Array
                        nw219_ = _dafny.Array(False, 20)
                        d_1453_v27_ = nw219_
                        index220_ = default__.safeIndex(324, (d_1453_v27_).length(0))
                        (d_1453_v27_)[index220_] = (False) == (False)
                        d_1454_v28_: _dafny.Seq
                        d_1454_v28_ = _dafny.SeqWithoutIsStrInference([d_1450_v26_, d_1450_v26_, d_1450_v26_])
                        index221_ = default__.safeIndex(324, (d_1453_v27_).length(0))
                        rhs217_ = d_1447_v23_
                        rhs218_ = len(d_1454_v28_)
                        rhs219_ = default__.safeModuloInt(d_1420_v0_, d_1420_v0_)
                        rhs220_ = (d_1450_v26_)[default__.safeIndex(46, (d_1450_v26_).length(0))]
                        rhs221_ = (self).f27
                        lhs170_ = globalState
                        lhs171_ = globalState
                        lhs172_ = d_1453_v27_
                        lhs173_ = default__.safeIndex(324, (d_1453_v27_).length(0))
                        d_1447_v23_ = rhs217_
                        lhs170_.f2 = rhs218_
                        d_1420_v0_ = rhs219_
                        lhs171_.f6 = rhs220_
                        lhs172_[lhs173_] = rhs221_
                        d_1455_v29_: _dafny.Set
                        d_1455_v29_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([d_1421_v1_]))})
                        (globalState).f25 = ((self).f27 if (d_1453_v27_)[default__.safeIndex(324, (d_1453_v27_).length(0))] else default__.fm0(d_1447_v23_, (d_1453_v27_)[default__.safeIndex(324, (d_1453_v27_).length(0))], (d_1453_v27_)[default__.safeIndex(324, (d_1453_v27_).length(0))], d_1455_v29_, globalState))
                    d_1456_v30_: _dafny.Seq
                    d_1456_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hvywufea"))
                    (globalState).f14 = ((default__.fm39((self).f27, d_1456_v30_, d_1420_v0_, globalState)).intersection(d_1432_v9_)) != (d_1432_v9_)
                    d_1457_v31_: C1
                    nw220_ = C1()
                    nw220_.ctor__((self).f27)
                    d_1457_v31_ = nw220_
                    pass
            pass
        (globalState).f2 = 875
        d_1458_v32_: _dafny.Set
        d_1458_v32_ = _dafny.Set({d_1421_v1_})
        d_1459_v33_: D7
        d_1459_v33_ = D7_DC17(912, d_1420_v0_, (self).f27)
        d_1460_v34_: _dafny.Set
        d_1460_v34_ = _dafny.Set({d_1459_v33_, d_1459_v33_, d_1459_v33_})
        d_1461_v35_: _dafny.Seq
        d_1461_v35_ = _dafny.SeqWithoutIsStrInference([d_1420_v0_])
        d_1462_v36_: _dafny.Seq
        d_1462_v36_ = _dafny.SeqWithoutIsStrInference([(d_1458_v32_).ispropersubset(_dafny.Set({d_1421_v1_})), (len(d_1460_v34_)) not in (d_1461_v35_)])
        if (d_1462_v36_)[default__.safeIndex(d_1420_v0_, len(d_1462_v36_))]:
            d_1463_v37_: _dafny.Array
            def lambda72_(d_1464_v33_):
                def lambda73_(d_1465_i4_):
                    return d_1464_v33_

                return lambda73_

            init36_ = lambda72_(d_1459_v33_)
            nw221_ = _dafny.Array(None, 16)
            for i0_36_ in range(nw221_.length(0)):
                nw221_[i0_36_] = init36_(i0_36_)
            d_1463_v37_ = nw221_
            index222_ = default__.safeIndex(148, (d_1463_v37_).length(0))
            (d_1463_v37_)[index222_] = d_1459_v33_
            d_1466_v38_: _dafny.Array
            nw222_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 9)
            d_1466_v38_ = nw222_
            d_1467_v39_: _dafny.Seq
            d_1467_v39_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))
            index223_ = default__.safeIndex(440, (d_1466_v38_).length(0))
            (d_1466_v38_)[index223_] = d_1467_v39_
            index224_ = default__.safeIndex(148, (d_1463_v37_).length(0))
            index225_ = default__.safeIndex(440, (d_1466_v38_).length(0))
            rhs222_ = D7_DC17(-907, d_1420_v0_, (self).f27)
            rhs223_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "joyils"))) + (d_1467_v39_)
            lhs174_ = d_1463_v37_
            lhs175_ = default__.safeIndex(148, (d_1463_v37_).length(0))
            lhs176_ = d_1466_v38_
            lhs177_ = default__.safeIndex(440, (d_1466_v38_).length(0))
            lhs174_[lhs175_] = rhs222_
            lhs176_[lhs177_] = rhs223_
            d_1468_v40_: _dafny.MultiSet
            d_1468_v40_ = _dafny.MultiSet([d_1461_v35_, default__.fm40(422, False, globalState)])
            d_1469_v41_: _dafny.Set
            d_1469_v41_ = _dafny.Set({(self).f27, (self).f27})
            d_1470_v42_: _dafny.Array
            nw223_ = _dafny.Array(None, 12)
            nw223_[int(0)] = ((self).f27 if (self).f27 else True)
            nw223_[int(1)] = not((d_1468_v40_) == (d_1468_v40_))
            nw223_[int(2)] = (d_1467_v39_) < ((d_1467_v39_).set(default__.safeIndex(d_1420_v0_, len(d_1467_v39_)), d_1421_v1_))
            nw223_[int(3)] = False
            nw223_[int(4)] = (self).f27
            nw223_[int(5)] = (self).f27
            nw223_[int(6)] = False
            nw223_[int(7)] = (self).f27
            nw223_[int(8)] = (self).f27
            nw223_[int(9)] = (self).f27
            nw223_[int(10)] = not((d_1469_v41_).isdisjoint(d_1469_v41_))
            nw223_[int(11)] = (d_1462_v36_)[default__.safeIndex(((d_1432_v9_)[d_1420_v0_] if (d_1420_v0_) in (d_1432_v9_) else default__.fm1(False, d_1462_v36_, (self).f27, globalState)), len(d_1462_v36_))]
            d_1470_v42_ = nw223_
            d_1471_v43_: _dafny.Set
            d_1471_v43_ = _dafny.Set({d_1420_v0_})
            d_1472_v44_: _dafny.Map
            d_1472_v44_ = _dafny.Map({(self).f27: d_1469_v41_})
            d_1473_v45_: _dafny.Array
            nw224_ = _dafny.Array(None, 17)
            nw224_[int(0)] = (d_1420_v0_) + (d_1420_v0_)
            nw224_[int(1)] = (0) - (len((d_1471_v43_) - (d_1471_v43_)))
            nw224_[int(2)] = d_1420_v0_
            nw224_[int(3)] = d_1420_v0_
            nw224_[int(4)] = (d_1420_v0_ if (self).f27 else len(d_1467_v39_))
            nw224_[int(5)] = d_1420_v0_
            nw224_[int(6)] = d_1420_v0_
            nw224_[int(7)] = default__.safeDivisionInt((d_1461_v35_)[default__.safeIndex(d_1420_v0_, len(d_1461_v35_))], d_1420_v0_)
            nw224_[int(8)] = d_1420_v0_
            nw224_[int(9)] = d_1420_v0_
            nw224_[int(10)] = d_1420_v0_
            nw224_[int(11)] = (0) - (-448)
            nw224_[int(12)] = d_1420_v0_
            nw224_[int(13)] = len(d_1472_v44_)
            nw224_[int(14)] = d_1420_v0_
            nw224_[int(15)] = d_1420_v0_
            nw224_[int(16)] = (0) - (default__.safeDivisionInt(len((d_1466_v38_)[default__.safeIndex(440, (d_1466_v38_).length(0))]), d_1420_v0_))
            d_1473_v45_ = nw224_
            index226_ = default__.safeIndex(664, (d_1473_v45_).length(0))
            (d_1473_v45_)[index226_] = d_1420_v0_
            d_1474_v46_: C2
            nw225_ = C2()
            nw225_.ctor__((self).f27, not((self).f27))
            d_1474_v46_ = nw225_
            d_1475_v47_: _dafny.Array
            nw226_ = _dafny.Array(None, 11)
            nw226_[int(0)] = d_1474_v46_
            nw226_[int(1)] = d_1474_v46_
            nw226_[int(2)] = d_1474_v46_
            nw226_[int(3)] = d_1474_v46_
            nw226_[int(4)] = d_1474_v46_
            nw226_[int(5)] = d_1474_v46_
            nw226_[int(6)] = d_1474_v46_
            nw226_[int(7)] = d_1474_v46_
            nw226_[int(8)] = (d_1474_v46_ if (self).f27 else d_1474_v46_)
            nw226_[int(9)] = d_1474_v46_
            nw226_[int(10)] = d_1474_v46_
            d_1475_v47_ = nw226_
            index227_ = default__.safeIndex(313, (d_1475_v47_).length(0))
            (d_1475_v47_)[index227_] = d_1474_v46_
            d_1476_v48_: D0
            d_1476_v48_ = D0_DC1(False)
            d_1477_v49_: _dafny.Seq
            d_1477_v49_ = _dafny.SeqWithoutIsStrInference([d_1476_v48_])
            d_1478_v50_: D23
            d_1478_v50_ = D23_DC61(d_1473_v45_, (self).f27, d_1420_v0_)
            d_1479_v51_: _dafny.MultiSet
            d_1479_v51_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_1476_v48_ for d_1480_i5_ in range(default__.abs(12))])])
            index228_ = default__.safeIndex(664, (d_1473_v45_).length(0))
            index229_ = default__.safeIndex(313, (d_1475_v47_).length(0))
            rhs224_ = d_1470_v42_
            rhs225_ = ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_1477_v49_]))).intersection((d_1479_v51_ if (d_1478_v50_).cf94 else d_1479_v51_))).cardinality
            rhs226_ = d_1474_v46_
            lhs178_ = d_1473_v45_
            lhs179_ = default__.safeIndex(664, (d_1473_v45_).length(0))
            lhs180_ = d_1475_v47_
            lhs181_ = default__.safeIndex(313, (d_1475_v47_).length(0))
            d_1470_v42_ = rhs224_
            lhs178_[lhs179_] = rhs225_
            lhs180_[lhs181_] = rhs226_
            index230_ = default__.safeIndex(316, (d_1470_v42_).length(0))
            (d_1470_v42_)[index230_] = False
            d_1481_v52_: _dafny.Array
            nw227_ = _dafny.Array(_dafny.CodePoint('D'), 12)
            d_1481_v52_ = nw227_
            d_1482_v53_: _dafny.Map
            d_1482_v53_ = _dafny.Map({d_1481_v52_: d_1420_v0_})
            d_1483_v54_: _dafny.MultiSet
            d_1483_v54_ = _dafny.MultiSet([d_1421_v1_])
            d_1484_v55_: D22
            d_1484_v55_ = D22_DC58((d_1474_v46_).f34, len(d_1482_v53_), (d_1483_v54_).cardinality)
            d_1485_v56_: _dafny.MultiSet
            d_1485_v56_ = _dafny.MultiSet([((d_1431_v8_)[not((d_1484_v55_).cf87)] if (not((d_1484_v55_).cf87)) in (d_1431_v8_) else (d_1474_v46_).f34)])
            d_1486_v57_: D0
            d_1486_v57_ = D0_DC2((0) - ((d_1461_v35_)[default__.safeIndex(d_1420_v0_, len(d_1461_v35_))]), d_1420_v0_)
            index231_ = default__.safeIndex(316, (d_1470_v42_).length(0))
            rhs227_ = (d_1486_v57_).cf2
            rhs228_ = True
            rhs229_ = d_1485_v56_
            lhs182_ = globalState
            lhs183_ = d_1470_v42_
            lhs184_ = default__.safeIndex(316, (d_1470_v42_).length(0))
            lhs182_.f2 = rhs227_
            lhs183_[lhs184_] = rhs228_
            d_1485_v56_ = rhs229_
            d_1431_v8_ = (d_1431_v8_).set((d_1470_v42_)[default__.safeIndex(316, (d_1470_v42_).length(0))], (self).f27)
            d_1487_v58_: _dafny.Map
            d_1487_v58_ = _dafny.Map({d_1420_v0_: default__.fm0(d_1467_v39_, (self).f27, (d_1474_v46_).f34, d_1471_v43_, globalState)})
            d_1488_v59_: _dafny.Array
            nw228_ = _dafny.Array(None, 11)
            nw228_[int(0)] = d_1461_v35_
            nw228_[int(1)] = (d_1461_v35_).set(default__.safeIndex(d_1420_v0_, len(d_1461_v35_)), d_1420_v0_)
            nw228_[int(2)] = _dafny.SeqWithoutIsStrInference([(d_1473_v45_)[default__.safeIndex(664, (d_1473_v45_).length(0))]])
            nw228_[int(3)] = (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_1420_v0_ for d_1489_i6_ in range(default__.abs(-764))]))])).set(default__.safeIndex(len(d_1487_v58_), len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_1420_v0_ for d_1490_i6_ in range(default__.abs(-764))]))]))), (d_1473_v45_)[default__.safeIndex(664, (d_1473_v45_).length(0))])
            nw228_[int(4)] = d_1461_v35_
            nw228_[int(5)] = d_1461_v35_
            nw228_[int(6)] = (d_1461_v35_ if (d_1474_v46_).f34 else d_1461_v35_)
            nw228_[int(7)] = d_1461_v35_
            nw228_[int(8)] = _dafny.SeqWithoutIsStrInference([(d_1473_v45_)[default__.safeIndex(664, (d_1473_v45_).length(0))] for d_1491_i7_ in range(default__.abs(359))])
            nw228_[int(9)] = d_1461_v35_
            nw228_[int(10)] = d_1461_v35_
            d_1488_v59_ = nw228_
            index232_ = default__.safeIndex(185, (d_1488_v59_).length(0))
            (d_1488_v59_)[index232_] = d_1461_v35_
            index233_ = default__.safeIndex(185, (d_1488_v59_).length(0))
            (d_1488_v59_)[index233_] = (d_1461_v35_) + (d_1461_v35_)
        elif True:
            if not(not((self).f27)):
                d_1492_v60_: _dafny.Array
                def lambda74_(d_1493_v0_):
                    def lambda75_(d_1494_i8_):
                        return default__.safeModuloInt(d_1494_i8_, d_1493_v0_)

                    return lambda75_

                init37_ = lambda74_(d_1420_v0_)
                nw229_ = _dafny.Array(None, 27)
                for i0_37_ in range(nw229_.length(0)):
                    nw229_[i0_37_] = init37_(i0_37_)
                d_1492_v60_ = nw229_
                index234_ = default__.safeIndex(748, (d_1492_v60_).length(0))
                (d_1492_v60_)[index234_] = 327
                index235_ = default__.safeIndex(748, (d_1492_v60_).length(0))
                (d_1492_v60_)[index235_] = (0) - (d_1420_v0_)
                index236_ = default__.safeIndex(748, (d_1492_v60_).length(0))
                def iife131_():
                    coll63_ = _dafny.Set()
                    compr_63_: int
                    for compr_63_ in _dafny.IntegerRange(468, -354):
                        d_1495_v61_: int = compr_63_
                        if ((468) <= (d_1495_v61_)) and ((d_1495_v61_) < (-354)):
                            coll63_ = coll63_.union(_dafny.Set([(d_1495_v61_) + (len(_dafny.Set({len(_dafny.Map({len(_dafny.Set({False, (self).f27})): 82}))})))]))
                    return _dafny.Set(coll63_)
                (d_1492_v60_)[index236_] = default__.safeDivisionInt(default__.safeDivisionInt(len(iife131_()
                ), (d_1492_v60_)[default__.safeIndex(748, (d_1492_v60_).length(0))]), len(d_1461_v35_))
                rhs230_ = d_1420_v0_
                rhs231_ = ((0) - ((d_1420_v0_) * (d_1420_v0_))) <= (d_1420_v0_)
                lhs185_ = globalState
                lhs186_ = globalState
                lhs185_.f2 = rhs230_
                lhs186_.f25 = rhs231_
                (globalState).f14 = (d_1420_v0_) < (default__.safeModuloInt(d_1420_v0_, (d_1492_v60_)[default__.safeIndex(748, (d_1492_v60_).length(0))]))
                d_1496_v62_: _dafny.Array
                nw230_ = _dafny.Array(False, 18)
                d_1496_v62_ = nw230_
                index237_ = default__.safeIndex(698, (d_1496_v62_).length(0))
                (d_1496_v62_)[index237_] = ((self).f27 if (self).f27 else (self).f27)
                d_1497_v63_: _dafny.Seq
                d_1497_v63_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bmbg"))
                index238_ = default__.safeIndex(698, (d_1496_v62_).length(0))
                rhs232_ = (d_1492_v60_)[default__.safeIndex(748, (d_1492_v60_).length(0))]
                rhs233_ = ((d_1462_v36_) + ((d_1462_v36_).set(default__.safeIndex(d_1420_v0_, len(d_1462_v36_)), not(((d_1431_v8_)[(self).f27] if ((self).f27) in (d_1431_v8_) else (self).f27))))) + (d_1462_v36_)
                rhs234_ = (not(not(True)) if True else (self).f27)
                rhs235_ = d_1497_v63_
                lhs187_ = d_1496_v62_
                lhs188_ = default__.safeIndex(698, (d_1496_v62_).length(0))
                r0 = rhs232_
                d_1462_v36_ = rhs233_
                lhs187_[lhs188_] = rhs234_
                d_1497_v63_ = rhs235_
            elif True:
                d_1498_v64_: _dafny.Set
                d_1498_v64_ = _dafny.Set({345})
                (globalState).f25 = ((self).fm36(d_1498_v64_, d_1498_v64_, globalState)) <= (d_1420_v0_)
                d_1499_v65_: _dafny.Map
                d_1499_v65_ = _dafny.Map({(d_1420_v0_) + (d_1420_v0_): not((self).f27)})
                d_1499_v65_ = (d_1499_v65_).set((d_1420_v0_) + (len(_dafny.Map({(self).f27: (self).f27}))), (self).f27)
                d_1500_v66_: _dafny.Map
                d_1500_v66_ = _dafny.Map({d_1420_v0_: d_1420_v0_})
                (globalState).f14 = not((default__.fm41(True, globalState)) != ((d_1500_v66_) | (d_1500_v66_)))
                d_1500_v66_ = (d_1500_v66_).set(d_1420_v0_, d_1420_v0_)
                d_1501_v67_: D21
                d_1501_v67_ = D21_DC55(d_1420_v0_)
                (globalState).f25 = (((self).fm36(_dafny.Set({(d_1501_v67_).cf82}), d_1498_v64_, globalState)) * (d_1420_v0_)) == (len(d_1462_v36_))
            (globalState).f2 = default__.safeDivisionInt((0) - ((len(d_1461_v35_)) - (len(d_1431_v8_))), (0) - ((0) - (((d_1432_v9_)[d_1420_v0_] if (d_1420_v0_) in (d_1432_v9_) else d_1420_v0_))))
            if (self).f27:
                d_1502_v68_: _dafny.Array
                nw231_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 29)
                d_1502_v68_ = nw231_
                index239_ = default__.safeIndex(494, (d_1502_v68_).length(0))
                (d_1502_v68_)[index239_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jhiytukdf"))
                d_1503_v69_: _dafny.Seq
                d_1503_v69_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ersma"))
                index240_ = default__.safeIndex(494, (d_1502_v68_).length(0))
                (d_1502_v68_)[index240_] = ((d_1503_v69_) + (_dafny.SeqWithoutIsStrInference([d_1421_v1_ for d_1504_i9_ in range(default__.abs(51))]))) + ((d_1503_v69_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kcjyqsyxd"))))
                (globalState).f2 = (0) - (d_1420_v0_)
                (globalState).f12 = (self).f27
                d_1505_v70_: _dafny.Set
                d_1505_v70_ = _dafny.Set({(len(_dafny.Map({d_1420_v0_: (self).f27}))) * (d_1420_v0_)})
                d_1505_v70_ = d_1505_v70_
                d_1506_v71_: C1
                nw232_ = C1()
                nw232_.ctor__((d_1420_v0_) >= (d_1420_v0_))
                d_1506_v71_ = nw232_
            elif True:
                d_1507_v72_: _dafny.Set
                d_1507_v72_ = _dafny.Set({(d_1420_v0_) < (d_1420_v0_)})
                d_1508_v73_: _dafny.MultiSet
                d_1508_v73_ = _dafny.MultiSet([(self).f27])
                d_1509_v74_: _dafny.Array
                nw233_ = _dafny.Array(False, 18)
                d_1509_v74_ = nw233_
                d_1510_v75_: _dafny.Seq
                d_1510_v75_ = _dafny.SeqWithoutIsStrInference([d_1509_v74_])
                rhs236_ = d_1420_v0_
                rhs237_ = d_1507_v72_
                rhs238_ = ((d_1510_v75_).set(default__.safeIndex(d_1420_v0_, len(d_1510_v75_)), d_1509_v74_)) <= ((d_1510_v75_) + (d_1510_v75_))
                rhs239_ = ((self).f37 if (d_1462_v36_)[default__.safeIndex(d_1420_v0_, len(d_1462_v36_))] else d_1508_v73_)
                lhs189_ = globalState
                lhs190_ = globalState
                lhs189_.f2 = rhs236_
                d_1507_v72_ = rhs237_
                lhs190_.f25 = rhs238_
                d_1508_v73_ = rhs239_
                d_1511_v76_: C2
                nw234_ = C2()
                nw234_.ctor__(not((d_1507_v72_).ispropersubset(d_1507_v72_)), ((self).f27 if (self).f27 else True))
                d_1511_v76_ = nw234_
                (globalState).f14 = not((d_1511_v76_).f34)
                d_1512_v77_: C0
                nw235_ = C0()
                nw235_.ctor__()
                d_1512_v77_ = nw235_
                (globalState).f5 = (self).f27
            d_1513_v78_: C1
            nw236_ = C1()
            nw236_.ctor__((self).f27)
            d_1513_v78_ = nw236_
            d_1514_v79_: _dafny.Map
            d_1514_v79_ = _dafny.Map({d_1420_v0_: d_1421_v1_})
            d_1515_v80_: _dafny.Set
            d_1515_v80_ = _dafny.Set({(d_1514_v79_).set(d_1420_v0_, d_1421_v1_), d_1514_v79_})
            d_1516_v81_: _dafny.Set
            d_1516_v81_ = d_1515_v80_
            d_1517_v82_: D3
            d_1517_v82_ = D3_DC9((self).f27, d_1420_v0_, 649)
            d_1518_v83_: _dafny.Map
            d_1518_v83_ = _dafny.Map({len((d_1461_v35_).set(default__.safeIndex((d_1517_v82_).cf16, len(d_1461_v35_)), d_1420_v0_)): d_1420_v0_})
            d_1519_v84_: _dafny.Seq
            d_1519_v84_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "abol"))
            d_1520_v85_: _dafny.Set
            d_1520_v85_ = _dafny.Set({d_1519_v84_})
            d_1521_v86_: D0
            d_1521_v86_ = D0_DC0(len(_dafny.Set({d_1519_v84_})))
            d_1522_v87_: _dafny.Array
            nw237_ = _dafny.Array(None, 25)
            nw237_[int(0)] = (d_1420_v0_) - (d_1420_v0_)
            nw237_[int(1)] = d_1420_v0_
            nw237_[int(2)] = 558
            nw237_[int(3)] = (d_1420_v0_) + (d_1420_v0_)
            nw237_[int(4)] = d_1420_v0_
            nw237_[int(5)] = d_1420_v0_
            nw237_[int(6)] = (0) - (d_1420_v0_)
            nw237_[int(7)] = d_1420_v0_
            nw237_[int(8)] = default__.safeModuloInt(d_1420_v0_, d_1420_v0_)
            nw237_[int(9)] = len(d_1462_v36_)
            nw237_[int(10)] = d_1420_v0_
            nw237_[int(11)] = 165
            nw237_[int(12)] = d_1420_v0_
            nw237_[int(13)] = len((d_1516_v81_))
            nw237_[int(14)] = (d_1420_v0_) - (d_1420_v0_)
            nw237_[int(15)] = (d_1420_v0_) - (368)
            nw237_[int(16)] = (d_1420_v0_) - (d_1420_v0_)
            nw237_[int(17)] = ((d_1518_v83_)[d_1420_v0_] if (d_1420_v0_) in (d_1518_v83_) else d_1420_v0_)
            nw237_[int(18)] = d_1420_v0_
            nw237_[int(19)] = len(d_1520_v85_)
            nw237_[int(20)] = d_1420_v0_
            nw237_[int(21)] = d_1420_v0_
            nw237_[int(22)] = d_1420_v0_
            nw237_[int(23)] = (d_1521_v86_).cf0
            nw237_[int(24)] = d_1420_v0_
            d_1522_v87_ = nw237_
            index241_ = default__.safeIndex(643, (d_1522_v87_).length(0))
            (d_1522_v87_)[index241_] = d_1420_v0_
            d_1523_v88_: _dafny.Map
            d_1523_v88_ = _dafny.Map({(self).f27: d_1420_v0_})
            index242_ = default__.safeIndex(643, (d_1522_v87_).length(0))
            (d_1522_v87_)[index242_] = len(d_1523_v88_)
        d_1524_v89_: _dafny.Set
        d_1524_v89_ = _dafny.Set({(0) - (d_1420_v0_)})
        d_1525_v90_: _dafny.MultiSet
        d_1525_v90_ = _dafny.MultiSet([d_1524_v89_, d_1524_v89_])
        d_1526_v91_: _dafny.Seq
        d_1526_v91_ = _dafny.SeqWithoutIsStrInference([d_1524_v89_])
        d_1527_v94_: T0
        nw238_ = C3()
        nw238_.ctor__((self).f27, not((self).f27))
        d_1527_v94_ = nw238_
        d_1528_v95_: _dafny.MultiSet
        d_1528_v95_ = d_1525_v90_
        d_1529_v96_: _dafny.Array
        nw239_ = _dafny.Array(None, 27)
        nw239_[int(0)] = d_1525_v90_
        nw239_[int(1)] = (_dafny.MultiSet([d_1524_v89_])).intersection(d_1525_v90_)
        nw239_[int(2)] = _dafny.MultiSet(d_1526_v91_)
        nw239_[int(3)] = d_1525_v90_
        def iife132_():
            coll64_ = _dafny.Set()
            compr_64_: int
            for compr_64_ in (d_1461_v35_).Elements:
                d_1530_v92_: int = compr_64_
                if (d_1530_v92_) in (d_1461_v35_):
                    coll64_ = coll64_.union(_dafny.Set([default__.safeModuloInt(d_1530_v92_, len((D6_DC13(_dafny.SeqWithoutIsStrInference([False, False, False, True, True]))).cf21))]))
            return _dafny.Set(coll64_)
        nw239_[int(4)] = _dafny.MultiSet([default__.fm42((self).f27, globalState), d_1524_v89_, d_1524_v89_, iife132_()
        ])
        nw239_[int(5)] = d_1525_v90_
        nw239_[int(6)] = d_1525_v90_
        nw239_[int(7)] = d_1525_v90_
        nw239_[int(8)] = d_1525_v90_
        nw239_[int(9)] = d_1525_v90_
        nw239_[int(10)] = d_1525_v90_
        nw239_[int(11)] = _dafny.MultiSet(d_1526_v91_)
        nw239_[int(12)] = d_1525_v90_
        nw239_[int(13)] = d_1525_v90_
        nw239_[int(14)] = d_1525_v90_
        nw239_[int(15)] = default__.fm43(d_1420_v0_, (0) - (d_1420_v0_), globalState)
        def iife133_():
            coll65_ = _dafny.Set()
            compr_65_: int
            for compr_65_ in (_dafny.SeqWithoutIsStrInference([d_1420_v0_])).Elements:
                d_1531_v93_: int = compr_65_
                if (d_1531_v93_) in (_dafny.SeqWithoutIsStrInference([d_1420_v0_])):
                    coll65_ = coll65_.union(_dafny.Set([default__.safeDivisionInt(d_1531_v93_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jjhiaoqi"))))]))
            return _dafny.Set(coll65_)
        nw239_[int(16)] = _dafny.MultiSet([iife133_()
        ])
        nw239_[int(17)] = d_1525_v90_
        nw239_[int(18)] = (d_1525_v90_) | (d_1525_v90_)
        nw239_[int(19)] = (d_1525_v90_).intersection(d_1525_v90_)
        nw239_[int(20)] = d_1525_v90_
        nw239_[int(21)] = d_1525_v90_
        nw239_[int(22)] = (_dafny.MultiSet([d_1524_v89_])).set((D20_DC49(d_1421_v1_, d_1420_v0_, d_1420_v0_, d_1527_v94_, d_1524_v89_)).cf73, default__.abs(d_1420_v0_))
        nw239_[int(23)] = (d_1525_v90_).set(d_1524_v89_, default__.abs(d_1420_v0_))
        nw239_[int(24)] = (d_1525_v90_) | (d_1525_v90_)
        nw239_[int(25)] = (d_1525_v90_) | (d_1525_v90_)
        nw239_[int(26)] = (d_1528_v95_)
        d_1529_v96_ = nw239_
        index243_ = default__.safeIndex(446, (d_1529_v96_).length(0))
        (d_1529_v96_)[index243_] = (d_1525_v90_ if (d_1527_v94_).f27 else d_1525_v90_)
        d_1532_v97_: D9
        d_1532_v97_ = D9_DC21((d_1527_v94_).f27, d_1420_v0_)
        d_1533_v99_: _dafny.Set
        d_1533_v99_ = _dafny.Set({d_1532_v97_, d_1532_v97_})
        index244_ = default__.safeIndex(446, (d_1529_v96_).length(0))
        def iife134_():
            coll66_ = _dafny.Set()
            compr_66_: D9
            for compr_66_ in (_dafny.Map({d_1532_v97_: (d_1527_v94_).f27})).keys.Elements:
                d_1534_v98_: D9 = compr_66_
                if (d_1534_v98_) in (_dafny.Map({d_1532_v97_: (d_1527_v94_).f27})):
                    coll66_ = coll66_.union(_dafny.Set([d_1534_v98_]))
            return _dafny.Set(coll66_)
        rhs240_ = (d_1533_v99_).ispropersubset(iife134_()
        )
        rhs241_ = _dafny.MultiSet(d_1526_v91_)
        rhs242_ = d_1421_v1_
        rhs243_ = 929
        lhs191_ = globalState
        lhs192_ = d_1529_v96_
        lhs193_ = default__.safeIndex(446, (d_1529_v96_).length(0))
        lhs194_ = globalState
        lhs191_.f25 = rhs240_
        lhs192_[lhs193_] = rhs241_
        d_1421_v1_ = rhs242_
        lhs194_.f2 = rhs243_
        d_1535_v100_: _dafny.Seq
        d_1535_v100_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vinsagpn"))
        d_1536_v101_: _dafny.Set
        d_1536_v101_ = _dafny.Set({(d_1527_v94_).f27, True, not((d_1527_v94_).f27), (d_1527_v94_).f27})
        d_1537_v102_: _dafny.Map
        d_1537_v102_ = _dafny.Map({d_1420_v0_: d_1420_v0_})
        d_1538_v103_: _dafny.Map
        d_1538_v103_ = _dafny.Map({d_1421_v1_: d_1537_v102_})
        d_1539_v104_: _dafny.Map
        d_1539_v104_ = _dafny.Map({(self).f27: d_1538_v103_})
        d_1540_v105_: _dafny.Map
        d_1540_v105_ = _dafny.Map({(d_1527_v94_).f27: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hbbvql"))})
        d_1541_v106_: _dafny.Array
        nw240_ = _dafny.Array(None, 20)
        nw240_[int(0)] = d_1420_v0_
        nw240_[int(1)] = -680
        nw240_[int(2)] = len(d_1535_v100_)
        nw240_[int(3)] = d_1420_v0_
        nw240_[int(4)] = (0) - ((self).fm36(_dafny.Set({769}), d_1524_v89_, globalState))
        nw240_[int(5)] = len(d_1536_v101_)
        nw240_[int(6)] = len(((d_1539_v104_)[(self).f27] if ((self).f27) in (d_1539_v104_) else ((d_1539_v104_)[(self).f27] if ((self).f27) in (d_1539_v104_) else d_1538_v103_)))
        nw240_[int(7)] = default__.safeModuloInt(d_1420_v0_, d_1420_v0_)
        nw240_[int(8)] = d_1420_v0_
        nw240_[int(9)] = (0) - (d_1420_v0_)
        nw240_[int(10)] = d_1420_v0_
        nw240_[int(11)] = len(d_1535_v100_)
        nw240_[int(12)] = d_1420_v0_
        nw240_[int(13)] = d_1420_v0_
        nw240_[int(14)] = default__.safeDivisionInt(d_1420_v0_, d_1420_v0_)
        nw240_[int(15)] = d_1420_v0_
        nw240_[int(16)] = (0) - (d_1420_v0_)
        nw240_[int(17)] = d_1420_v0_
        nw240_[int(18)] = default__.fm1((self).f27, d_1462_v36_, False, globalState)
        nw240_[int(19)] = len(((d_1540_v105_)[(self).f27] if ((self).f27) in (d_1540_v105_) else d_1535_v100_))
        d_1541_v106_ = nw240_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_1541_v106_).length(0)):
            d_1542_i10_: int = guard_loop_4_
            if (True) and (((0) <= (d_1542_i10_)) and ((d_1542_i10_) < ((d_1541_v106_).length(0)))):
                def iife135_():
                    coll67_ = _dafny.Map()
                    compr_67_: _dafny.Seq
                    for compr_67_ in (_dafny.SeqWithoutIsStrInference([d_1535_v100_])).Elements:
                        d_1543_v107_: _dafny.Seq = compr_67_
                        if (d_1543_v107_) in (_dafny.SeqWithoutIsStrInference([d_1535_v100_])):
                            coll67_[d_1543_v107_] = _dafny.Set({d_1420_v0_})
                    return _dafny.Map(coll67_)
                (d_1541_v106_)[(d_1542_i10_)] = (d_1542_i10_) + ((d_1420_v0_) - (len(_dafny.Map({_dafny.SeqWithoutIsStrInference([(d_1527_v94_).f27, (d_1527_v94_).f27, (self).f27, True, (self).f27]): D25_DC63(iife135_()
)}))))
        r0 = d_1420_v0_
        return r0

    def m2(self, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: bool = False
        d_1544_v0_: int
        d_1544_v0_ = 674
        d_1545_v1_: _dafny.Seq
        d_1545_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hltvliyk"))
        d_1546_v2_: _dafny.Map
        d_1546_v2_ = _dafny.Map({d_1544_v0_: d_1545_v1_})
        d_1547_v3_: _dafny.Seq
        d_1547_v3_ = _dafny.SeqWithoutIsStrInference([656, d_1544_v0_])
        d_1548_v4_: _dafny.Map
        d_1548_v4_ = _dafny.Map({d_1544_v0_: False})
        d_1549_v5_: _dafny.Map
        d_1549_v5_ = _dafny.Map({d_1548_v4_: (self).f27})
        d_1550_v6_: _dafny.Seq
        d_1550_v6_ = _dafny.SeqWithoutIsStrInference([d_1549_v5_])
        d_1551_v7_: _dafny.Set
        d_1551_v7_ = _dafny.Set({d_1544_v0_})
        d_1552_v8_: str
        d_1552_v8_ = _dafny.CodePoint('u')
        d_1553_v9_: _dafny.Map
        d_1553_v9_ = _dafny.Map({_dafny.Map({not((self).f27): d_1544_v0_}): d_1552_v8_})
        d_1554_v11_: _dafny.Map
        def iife136_():
            coll68_ = _dafny.Set()
            compr_68_: int
            for compr_68_ in _dafny.IntegerRange(150, -379):
                d_1555_v10_: int = compr_68_
                if ((150) <= (d_1555_v10_)) and ((d_1555_v10_) < (-379)):
                    coll68_ = coll68_.union(_dafny.Set([default__.safeDivisionInt(d_1555_v10_, d_1544_v0_)]))
            return _dafny.Set(coll68_)
        d_1554_v11_ = _dafny.Map({(self).f27: len(iife136_()
        )})
        d_1556_v12_: _dafny.MultiSet
        d_1556_v12_ = _dafny.MultiSet([d_1544_v0_])
        d_1557_v13_: _dafny.Map
        d_1557_v13_ = _dafny.Map({(d_1556_v12_).cardinality: d_1552_v8_})
        d_1558_v14_: _dafny.Map
        d_1558_v14_ = _dafny.Map({((d_1553_v9_)[d_1554_v11_] if (d_1554_v11_) in (d_1553_v9_) else d_1552_v8_): d_1557_v13_})
        d_1559_v15_: _dafny.Array
        nw241_ = _dafny.Array(None, 12)
        nw241_[int(0)] = (len(((d_1546_v2_)[d_1544_v0_] if (d_1544_v0_) in (d_1546_v2_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ddn"))))) + (d_1544_v0_)
        nw241_[int(1)] = (len(d_1545_v1_)) * (len(d_1547_v3_))
        nw241_[int(2)] = d_1544_v0_
        nw241_[int(3)] = d_1544_v0_
        nw241_[int(4)] = (d_1544_v0_) + (d_1544_v0_)
        nw241_[int(5)] = d_1544_v0_
        nw241_[int(6)] = (len((d_1550_v6_)[default__.safeIndex(len(d_1551_v7_), len(d_1550_v6_))])) * (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_1560_i0_ in range(default__.abs(908))])))
        nw241_[int(7)] = ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_1561_i1_ in range(default__.abs(489))])))) * (d_1544_v0_)
        nw241_[int(8)] = d_1544_v0_
        nw241_[int(9)] = (0) - (len((d_1545_v1_) + ((self).fm2(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oilids"))), (d_1558_v14_).set(d_1552_v8_, d_1557_v13_), globalState))))
        nw241_[int(10)] = d_1544_v0_
        nw241_[int(11)] = (d_1544_v0_ if (self).f27 else (0) - (len(_dafny.SeqWithoutIsStrInference([(self).f27, (self).f27]))))
        d_1559_v15_ = nw241_
        index245_ = default__.safeIndex(903, (d_1559_v15_).length(0))
        (d_1559_v15_)[index245_] = d_1544_v0_
        d_1562_v16_: _dafny.Seq
        d_1562_v16_ = _dafny.SeqWithoutIsStrInference([(self).f27])
        index246_ = default__.safeIndex(903, (d_1559_v15_).length(0))
        (d_1559_v15_)[index246_] = default__.fm1((self).f27, d_1562_v16_, (d_1544_v0_) > (d_1544_v0_), globalState)
        d_1563_v17_: _dafny.Array
        nw242_ = _dafny.Array(_dafny.Map({}), 8)
        d_1563_v17_ = nw242_
        def iife137_():
            coll69_ = _dafny.Map()
            compr_69_: int
            for compr_69_ in _dafny.IntegerRange(932, 817):
                d_1565_v18_: int = compr_69_
                if ((932) <= (d_1565_v18_)) and ((d_1565_v18_) < (817)):
                    coll69_[(d_1565_v18_) + ((d_1559_v15_)[default__.safeIndex(903, (d_1559_v15_).length(0))])] = (len(_dafny.Map({(d_1562_v16_)[default__.safeIndex(d_1544_v0_, len(d_1562_v16_))]: d_1562_v16_}))) - (d_1544_v0_)
            return _dafny.Map(coll69_)
        _ingredientsd_0 = []
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_1563_v17_).length(0)):
            d_1564_i2_: int = guard_loop_5_
            if (True) and (((0) <= (d_1564_i2_)) and ((d_1564_i2_) < ((d_1563_v17_).length(0)))):
                _ingredientsd_0.append((d_1563_v17_, int(d_1564_i2_), iife137_()
                ))
        for _tupd_0 in _ingredientsd_0:
            _tupd_0[0][_tupd_0[1]] = _tupd_0[2]
        d_1566_i3_: int
        d_1566_i3_ = 0
        with _dafny.label("9"):
            while (self).f27:
                with _dafny.c_label("9"):
                    if (d_1566_i3_) >= (100):
                        raise _dafny.Break("9")
                    d_1566_i3_ = (d_1566_i3_) + (1)
                    d_1544_v0_ = (d_1559_v15_)[default__.safeIndex(903, (d_1559_v15_).length(0))]
                    def iife138_():
                        coll70_ = _dafny.Set()
                        compr_70_: int
                        for compr_70_ in _dafny.IntegerRange(801, 760):
                            d_1567_v19_: int = compr_70_
                            if ((801) <= (d_1567_v19_)) and ((d_1567_v19_) < (760)):
                                coll70_ = coll70_.union(_dafny.Set([(d_1567_v19_) - (len(_dafny.Map({(self).f27: d_1544_v0_})))]))
                        return _dafny.Set(coll70_)
                    if default__.fm0(d_1545_v1_, not((self).f27), not((self).f27), (d_1551_v7_) | (iife138_()
                    ), globalState):
                        d_1568_v20_: C2
                        nw243_ = C2()
                        nw243_.ctor__(not((not((self).f27) if not((self).f27) else (self).f27)), (self).f27)
                        d_1568_v20_ = nw243_
                        d_1569_v21_: _dafny.Array
                        nw244_ = _dafny.Array(None, 3)
                        nw244_[int(0)] = (d_1568_v20_).f34
                        nw244_[int(1)] = (d_1568_v20_).f34
                        nw244_[int(2)] = not ((d_1568_v20_).f34) or ((self).f27)
                        d_1569_v21_ = nw244_
                        d_1570_v22_: _dafny.Map
                        d_1570_v22_ = _dafny.Map({d_1544_v0_: 903})
                        index247_ = default__.safeIndex(230, (d_1569_v21_).length(0))
                        (d_1569_v21_)[index247_] = (len(d_1570_v22_)) >= ((d_1559_v15_)[default__.safeIndex(903, (d_1559_v15_).length(0))])
                        index248_ = default__.safeIndex(230, (d_1569_v21_).length(0))
                        (d_1569_v21_)[index248_] = (d_1570_v22_) == (d_1570_v22_)
                        d_1571_v23_: C3
                        nw245_ = C3()
                        nw245_.ctor__((d_1568_v20_).f34, (d_1569_v21_)[default__.safeIndex(230, (d_1569_v21_).length(0))])
                        d_1571_v23_ = nw245_
                        d_1572_v24_: C0
                        nw246_ = C0()
                        nw246_.ctor__()
                        d_1572_v24_ = nw246_
                        d_1552_v8_ = (default__.fm45(globalState)).cf100
                    elif True:
                        d_1573_v25_: _dafny.Array
                        nw247_ = _dafny.Array(_dafny.Array(None, 0), 11)
                        d_1573_v25_ = nw247_
                        r0 = d_1573_v25_
                        d_1559_v15_ = d_1559_v15_
                        d_1574_v26_: _dafny.MultiSet
                        d_1574_v26_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(d_1559_v15_)[default__.safeIndex(903, (d_1559_v15_).length(0))] for d_1575_i4_ in range(default__.abs(925))]), d_1547_v3_])
                        d_1576_v27_: _dafny.Map
                        d_1576_v27_ = _dafny.Map({((d_1574_v26_)[d_1547_v3_] if (d_1547_v3_) in (d_1574_v26_) else d_1544_v0_): (d_1556_v12_).cardinality})
                        def iife139_():
                            coll71_ = _dafny.Map()
                            compr_71_: int
                            for compr_71_ in _dafny.IntegerRange(908, -623):
                                d_1577_v28_: int = compr_71_
                                if ((908) <= (d_1577_v28_)) and ((d_1577_v28_) < (-623)):
                                    coll71_[(d_1577_v28_) + (d_1544_v0_)] = d_1544_v0_
                            return _dafny.Map(coll71_)
                        d_1576_v27_ = (_dafny.Map({(d_1559_v15_)[default__.safeIndex(903, (d_1559_v15_).length(0))]: d_1544_v0_})) | (iife139_()
                        )
                        d_1578_v29_: _dafny.Array
                        def lambda76_(d_1579_i5_):
                            return (self).f27

                        init38_ = lambda76_
                        nw248_ = _dafny.Array(None, 14)
                        for i0_38_ in range(nw248_.length(0)):
                            nw248_[i0_38_] = init38_(i0_38_)
                        d_1578_v29_ = nw248_
                        index249_ = default__.safeIndex(154, (d_1578_v29_).length(0))
                        (d_1578_v29_)[index249_] = (d_1544_v0_) >= (d_1544_v0_)
                        index250_ = default__.safeIndex(154, (d_1578_v29_).length(0))
                        (d_1578_v29_)[index250_] = (self).f27
                        index251_ = default__.safeIndex(154, (d_1578_v29_).length(0))
                        (d_1578_v29_)[index251_] = (d_1578_v29_)[default__.safeIndex(154, (d_1578_v29_).length(0))]
                    index252_ = default__.safeIndex(903, (d_1559_v15_).length(0))
                    (d_1559_v15_)[index252_] = (0) - (default__.safeModuloInt(d_1544_v0_, (d_1559_v15_)[default__.safeIndex(903, (d_1559_v15_).length(0))]))
                    d_1580_v30_: C1
                    nw249_ = C1()
                    nw249_.ctor__((self).f27)
                    d_1580_v30_ = nw249_
                    pass
            pass
        guard_loop_6_: int
        for guard_loop_6_ in _dafny.IntegerRange(0, (d_1559_v15_).length(0)):
            d_1581_i6_: int = guard_loop_6_
            if (True) and (((0) <= (d_1581_i6_)) and ((d_1581_i6_) < ((d_1559_v15_).length(0)))):
                (d_1559_v15_)[(d_1581_i6_)] = default__.safeDivisionInt(d_1581_i6_, (D21_DC53((self).f27, 530, (self).f27)).cf77)
        hi10_ = d_1544_v0_
        for d_1582_i7_ in range((d_1559_v15_)[default__.safeIndex(903, (d_1559_v15_).length(0))], hi10_):
            if (self).f27:
                d_1583_v31_: _dafny.Map
                d_1583_v31_ = _dafny.Map({(self).f27: (self).f27})
                d_1584_v32_: C1
                nw250_ = C1()
                nw250_.ctor__((d_1583_v31_) == (d_1583_v31_))
                d_1584_v32_ = nw250_
                d_1585_v33_: _dafny.Map
                d_1585_v33_ = _dafny.Map({d_1544_v0_: _dafny.MultiSet([(self).f27, (self).f27, (self).f27])})
                (globalState).f25 = ((self).f27) or ((((d_1585_v33_)[((d_1556_v12_)[len(d_1545_v1_)] if (len(d_1545_v1_)) in (d_1556_v12_) else len(_dafny.Map({d_1559_v15_: (d_1559_v15_)[default__.safeIndex(903, (d_1559_v15_).length(0))]})))] if (((d_1556_v12_)[len(d_1545_v1_)] if (len(d_1545_v1_)) in (d_1556_v12_) else len(_dafny.Map({d_1559_v15_: (d_1559_v15_)[default__.safeIndex(903, (d_1559_v15_).length(0))]})))) in (d_1585_v33_) else _dafny.MultiSet([(self).f27, (self).f27, not((self).f27), False]))) != ((self).f37))
                d_1586_v34_: _dafny.Array
                def lambda77_(d_1587_v8_):
                    def lambda78_(d_1588_i8_):
                        return d_1587_v8_

                    return lambda78_

                init39_ = lambda77_(d_1552_v8_)
                nw251_ = _dafny.Array(None, 24)
                for i0_39_ in range(nw251_.length(0)):
                    nw251_[i0_39_] = init39_(i0_39_)
                d_1586_v34_ = nw251_
                index253_ = default__.safeIndex(477, (d_1586_v34_).length(0))
                (d_1586_v34_)[index253_] = d_1552_v8_
                index254_ = default__.safeIndex(477, (d_1586_v34_).length(0))
                (d_1586_v34_)[index254_] = default__.fm44((self).f27, False, (_dafny.SeqWithoutIsStrInference([d_1552_v8_ for d_1589_i9_ in range(default__.abs(707))]) if (self).f27 else d_1545_v1_), len(d_1547_v3_), globalState)
                d_1590_v35_: _dafny.Array
                nw252_ = _dafny.Array(False, 21)
                d_1590_v35_ = nw252_
                d_1591_v36_: _dafny.Seq
                d_1591_v36_ = _dafny.SeqWithoutIsStrInference([d_1590_v35_, d_1590_v35_, d_1590_v35_, d_1590_v35_, d_1590_v35_])
                d_1591_v36_ = (d_1591_v36_) + (_dafny.SeqWithoutIsStrInference([d_1590_v35_]))
                (globalState).f12 = ((self).f27) or (((self).f27) and ((self).f27))
            elif True:
                (globalState).f6 = default__.fm1(False, (d_1562_v16_) + (d_1562_v16_), (self).f27, globalState)
                d_1592_v37_: D5
                d_1592_v37_ = D5_DC12((self).f27, (d_1559_v15_)[default__.safeIndex(903, (d_1559_v15_).length(0))])
                d_1544_v0_ = (d_1592_v37_).cf20
                index255_ = default__.safeIndex(903, (d_1559_v15_).length(0))
                (d_1559_v15_)[index255_] = (0) - (default__.safeModuloInt(default__.fm1((self).f27, d_1562_v16_, (self).f27, globalState), d_1582_i7_))
                d_1593_v38_: D0
                d_1593_v38_ = D0_DC3(default__.fm46(_dafny.MultiSet([(self).f27]), (self).f27, globalState))
                d_1594_v39_: D0
                d_1594_v39_ = D0_DC3(d_1593_v38_)
                d_1595_v40_: D0
                d_1595_v40_ = D0_DC3(d_1593_v38_)
                d_1596_v41_: D0
                d_1596_v41_ = D0_DC1((self).f27)
                d_1597_v42_: bool
                d_1598_v43_: _dafny.Seq
                out64_: bool
                out65_: _dafny.Seq
                out64_, out65_ = default__.m0(d_1595_v40_, d_1596_v41_, globalState)
                d_1597_v42_ = out64_
                d_1598_v43_ = out65_
                d_1599_v44_: _dafny.Array
                nw253_ = _dafny.Array(False, 23)
                d_1599_v44_ = nw253_
                d_1599_v44_ = d_1599_v44_
            d_1600_v45_: _dafny.Array
            def lambda79_(d_1601_v8_):
                def lambda80_(d_1602_i10_):
                    return d_1601_v8_

                return lambda80_

            init40_ = lambda79_(d_1552_v8_)
            nw254_ = _dafny.Array(None, 5)
            for i0_40_ in range(nw254_.length(0)):
                nw254_[i0_40_] = init40_(i0_40_)
            d_1600_v45_ = nw254_
            d_1603_v46_: _dafny.Seq
            d_1603_v46_ = _dafny.SeqWithoutIsStrInference([d_1600_v45_, d_1600_v45_, d_1600_v45_])
            d_1604_v47_: _dafny.Array
            nw255_ = _dafny.Array(None, 16)
            nw255_[int(0)] = d_1600_v45_
            nw255_[int(1)] = d_1600_v45_
            nw255_[int(2)] = d_1600_v45_
            nw255_[int(3)] = d_1600_v45_
            nw255_[int(4)] = d_1600_v45_
            nw255_[int(5)] = d_1600_v45_
            nw255_[int(6)] = d_1600_v45_
            nw255_[int(7)] = d_1600_v45_
            nw255_[int(8)] = d_1600_v45_
            nw255_[int(9)] = d_1600_v45_
            nw255_[int(10)] = d_1600_v45_
            nw255_[int(11)] = d_1600_v45_
            nw255_[int(12)] = d_1600_v45_
            nw255_[int(13)] = d_1600_v45_
            nw255_[int(14)] = d_1600_v45_
            nw255_[int(15)] = (d_1603_v46_)[default__.safeIndex((d_1559_v15_)[default__.safeIndex(903, (d_1559_v15_).length(0))], len(d_1603_v46_))]
            d_1604_v47_ = nw255_
            d_1604_v47_ = d_1604_v47_
            d_1605_v48_: _dafny.Array
            nw256_ = _dafny.Array(_dafny.Map({}), 16)
            d_1605_v48_ = nw256_
            index256_ = default__.safeIndex(83, (d_1605_v48_).length(0))
            (d_1605_v48_)[index256_] = d_1554_v11_
            index257_ = default__.safeIndex(83, (d_1605_v48_).length(0))
            rhs244_ = d_1552_v8_
            rhs245_ = d_1554_v11_
            rhs246_ = ((d_1559_v15_)[default__.safeIndex(903, (d_1559_v15_).length(0))]) - (d_1544_v0_)
            lhs195_ = d_1605_v48_
            lhs196_ = default__.safeIndex(83, (d_1605_v48_).length(0))
            d_1552_v8_ = rhs244_
            lhs195_[lhs196_] = rhs245_
            d_1544_v0_ = rhs246_
            d_1606_v49_: C1
            nw257_ = C1()
            nw257_.ctor__((self).f27)
            d_1606_v49_ = nw257_
        d_1607_v50_: _dafny.Array
        nw258_ = _dafny.Array(None, 1)
        nw258_[int(0)] = (self).f27
        d_1607_v50_ = nw258_
        d_1608_v51_: D17
        d_1608_v51_ = D17_DC38(d_1607_v50_)
        pat_let_tv37_ = d_1607_v50_
        def iife140_(_pat_let34_0):
            def iife141_(d_1609_dt__update__tmp_h0_):
                def iife142_(_pat_let35_0):
                    def iife143_(d_1610_dt__update_hcf53_h0_):
                        return D17_DC38(d_1610_dt__update_hcf53_h0_)
                    return iife143_(_pat_let35_0)
                return iife142_(pat_let_tv37_)
            return iife141_(_pat_let34_0)
        d_1607_v50_ = (iife140_(d_1608_v51_)).cf53
        d_1611_v52_: _dafny.Map
        d_1611_v52_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "catvgx")): d_1607_v50_})
        d_1612_v53_: _dafny.Array
        nw259_ = _dafny.Array(None, 21)
        nw259_[int(0)] = d_1607_v50_
        nw259_[int(1)] = d_1607_v50_
        nw259_[int(2)] = d_1607_v50_
        nw259_[int(3)] = d_1607_v50_
        nw259_[int(4)] = d_1607_v50_
        nw259_[int(5)] = d_1607_v50_
        nw259_[int(6)] = d_1607_v50_
        nw259_[int(7)] = d_1607_v50_
        nw259_[int(8)] = d_1607_v50_
        nw259_[int(9)] = d_1607_v50_
        nw259_[int(10)] = d_1607_v50_
        nw259_[int(11)] = d_1607_v50_
        nw259_[int(12)] = ((d_1611_v52_)[(d_1545_v1_).set(default__.safeIndex((d_1559_v15_)[default__.safeIndex(903, (d_1559_v15_).length(0))], len(d_1545_v1_)), d_1552_v8_)] if ((d_1545_v1_).set(default__.safeIndex((d_1559_v15_)[default__.safeIndex(903, (d_1559_v15_).length(0))], len(d_1545_v1_)), d_1552_v8_)) in (d_1611_v52_) else d_1607_v50_)
        nw259_[int(13)] = d_1607_v50_
        nw259_[int(14)] = d_1607_v50_
        nw259_[int(15)] = d_1607_v50_
        nw259_[int(16)] = d_1607_v50_
        nw259_[int(17)] = (d_1607_v50_ if (self).f27 else d_1607_v50_)
        nw259_[int(18)] = (d_1608_v51_).cf53
        nw259_[int(19)] = d_1607_v50_
        nw259_[int(20)] = d_1607_v50_
        d_1612_v53_ = nw259_
        r0 = d_1612_v53_
        r1 = d_1545_v1_
        d_1613_v54_: _dafny.Array
        nw260_ = _dafny.Array(_dafny.Array(None, 0), 18)
        d_1613_v54_ = nw260_
        r2 = (D5_DC11(d_1613_v54_)).cf18
        r3 = (self).f27
        return r0, r1, r2, r3

    def m12(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: bool = False
        r3: bool = False
        d_1614_v0_: _dafny.Set
        d_1614_v0_ = _dafny.Set({((self).f27) and (p2), False})
        d_1615_v1_: _dafny.Array
        nw261_ = _dafny.Array(int(0), 12)
        d_1615_v1_ = nw261_
        d_1616_v2_: _dafny.Seq
        d_1616_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vikv"))
        d_1617_v3_: _dafny.Map
        d_1617_v3_ = _dafny.Map({p2: p2})
        d_1618_v4_: _dafny.Array
        nw262_ = _dafny.Array(None, 8)
        nw262_[int(0)] = p0
        nw262_[int(1)] = (self).f27
        nw262_[int(2)] = default__.fm0(d_1616_v2_, p0, p0, _dafny.Set({(0) - (len(d_1617_v3_))}), globalState)
        nw262_[int(3)] = p0
        nw262_[int(4)] = p0
        nw262_[int(5)] = p0
        nw262_[int(6)] = p0
        nw262_[int(7)] = (self).f27
        d_1618_v4_ = nw262_
        d_1619_v5_: int
        d_1619_v5_ = -826
        d_1620_v6_: _dafny.Seq
        d_1620_v6_ = _dafny.SeqWithoutIsStrInference([d_1618_v4_])
        rhs247_ = d_1614_v0_
        rhs248_ = (0) - ((d_1619_v5_) * (len(_dafny.Map({d_1618_v4_: d_1619_v5_}))))
        rhs249_ = d_1615_v1_
        rhs250_ = d_1619_v5_
        rhs251_ = (d_1620_v6_)[default__.safeIndex(d_1619_v5_, len(d_1620_v6_))]
        lhs197_ = globalState
        lhs198_ = globalState
        d_1614_v0_ = rhs247_
        lhs197_.f2 = rhs248_
        d_1615_v1_ = rhs249_
        lhs198_.f2 = rhs250_
        d_1618_v4_ = rhs251_
        d_1621_v7_: _dafny.Array
        def lambda81_(d_1622_i1_):
            return _dafny.CodePoint('o')

        init41_ = lambda81_
        nw263_ = _dafny.Array(None, 29)
        for i0_41_ in range(nw263_.length(0)):
            nw263_[i0_41_] = init41_(i0_41_)
        d_1621_v7_ = nw263_
        guard_loop_7_: int
        for guard_loop_7_ in _dafny.IntegerRange(0, (d_1621_v7_).length(0)):
            d_1623_i0_: int = guard_loop_7_
            if (True) and (((0) <= (d_1623_i0_)) and ((d_1623_i0_) < ((d_1621_v7_).length(0)))):
                (d_1621_v7_)[(d_1623_i0_)] = _dafny.CodePoint('c')
        d_1624_v8_: D0
        d_1624_v8_ = D0_DC3(default__.fm46((self).f37, (self).f27, globalState))
        d_1625_v9_: D0
        d_1625_v9_ = D0_DC1(p2)
        d_1626_v10_: bool
        d_1627_v11_: _dafny.Seq
        out66_: bool
        out67_: _dafny.Seq
        out66_, out67_ = default__.m0(d_1624_v8_, d_1625_v9_, globalState)
        d_1626_v10_ = out66_
        d_1627_v11_ = out67_
        d_1628_v12_: _dafny.MultiSet
        d_1628_v12_ = _dafny.MultiSet([d_1618_v4_])
        d_1629_v13_: _dafny.Seq
        d_1629_v13_ = _dafny.SeqWithoutIsStrInference([d_1619_v5_])
        d_1630_v14_: _dafny.Array
        nw264_ = _dafny.Array(None, 13)
        nw264_[int(0)] = d_1628_v12_
        nw264_[int(1)] = d_1628_v12_
        nw264_[int(2)] = d_1628_v12_
        nw264_[int(3)] = d_1628_v12_
        nw264_[int(4)] = (d_1628_v12_) | (d_1628_v12_)
        nw264_[int(5)] = d_1628_v12_
        nw264_[int(6)] = (_dafny.MultiSet([d_1618_v4_, d_1618_v4_])).intersection(d_1628_v12_)
        nw264_[int(7)] = (d_1628_v12_) - (_dafny.MultiSet([d_1618_v4_]))
        nw264_[int(8)] = _dafny.MultiSet([d_1618_v4_, d_1618_v4_])
        nw264_[int(9)] = (d_1628_v12_) - (d_1628_v12_)
        nw264_[int(10)] = d_1628_v12_
        nw264_[int(11)] = _dafny.MultiSet([d_1618_v4_])
        nw264_[int(12)] = ((d_1628_v12_).set(d_1618_v4_, default__.abs(d_1619_v5_))).set(d_1618_v4_, default__.abs((d_1629_v13_)[default__.safeIndex(d_1619_v5_, len(d_1629_v13_))]))
        d_1630_v14_ = nw264_
        index258_ = default__.safeIndex(592, (d_1630_v14_).length(0))
        (d_1630_v14_)[index258_] = (d_1628_v12_).intersection(_dafny.MultiSet([d_1618_v4_]))
        d_1631_v15_: _dafny.Seq
        d_1631_v15_ = _dafny.SeqWithoutIsStrInference([d_1628_v12_])
        d_1632_v16_: _dafny.Map
        d_1632_v16_ = _dafny.Map({(d_1615_v1_ if (self).f27 else d_1615_v1_): (d_1631_v15_)[default__.safeIndex(len(d_1627_v11_), len(d_1631_v15_))]})
        index259_ = default__.safeIndex(592, (d_1630_v14_).length(0))
        (d_1630_v14_)[index259_] = ((d_1632_v16_)[d_1615_v1_] if (d_1615_v1_) in (d_1632_v16_) else (d_1628_v12_) | (d_1628_v12_))
        if p0:
            (globalState).f25 = not(p2)
            d_1633_v17_: _dafny.Map
            d_1633_v17_ = _dafny.Map({d_1619_v5_: d_1619_v5_})
            (globalState).f21 = default__.fm0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ssvibt")), not(not(d_1626_v10_)), p0, _dafny.Set({(0) - (((d_1633_v17_)[d_1619_v5_] if (d_1619_v5_) in (d_1633_v17_) else d_1619_v5_)), d_1619_v5_}), globalState)
            d_1634_v18_: _dafny.Map
            d_1634_v18_ = _dafny.Map({d_1629_v13_: d_1619_v5_})
            d_1634_v18_ = d_1634_v18_
            if ((0) - (d_1619_v5_)) not in (d_1629_v13_):
                d_1635_v19_: _dafny.Array
                nw265_ = _dafny.Array(None, 2)
                nw265_[int(0)] = d_1618_v4_
                nw265_[int(1)] = d_1618_v4_
                d_1635_v19_ = nw265_
                index260_ = default__.safeIndex(396, (d_1635_v19_).length(0))
                (d_1635_v19_)[index260_] = d_1618_v4_
                index261_ = default__.safeIndex(740, (d_1618_v4_).length(0))
                (d_1618_v4_)[index261_] = d_1626_v10_
                d_1636_v20_: D9
                d_1636_v20_ = D9_DC21(p0, d_1619_v5_)
                d_1637_v21_: D17
                d_1637_v21_ = D17_DC39(d_1627_v11_)
                index262_ = default__.safeIndex(396, (d_1635_v19_).length(0))
                index263_ = default__.safeIndex(740, (d_1618_v4_).length(0))
                rhs252_ = d_1627_v11_
                rhs253_ = d_1618_v4_
                rhs254_ = ((d_1636_v20_).cf32) == (d_1619_v5_)
                rhs255_ = len((d_1637_v21_).cf54)
                lhs199_ = d_1635_v19_
                lhs200_ = default__.safeIndex(396, (d_1635_v19_).length(0))
                lhs201_ = d_1618_v4_
                lhs202_ = default__.safeIndex(740, (d_1618_v4_).length(0))
                lhs203_ = globalState
                d_1627_v11_ = rhs252_
                lhs199_[lhs200_] = rhs253_
                lhs201_[lhs202_] = rhs254_
                lhs203_.f6 = rhs255_
                (globalState).f21 = False
                (globalState).f2 = d_1619_v5_
                index264_ = default__.safeIndex(396, (d_1635_v19_).length(0))
                (d_1635_v19_)[index264_] = (d_1620_v6_)[default__.safeIndex((d_1629_v13_)[default__.safeIndex(d_1619_v5_, len(d_1629_v13_))], len(d_1620_v6_))]
                d_1638_v23_: str
                d_1638_v23_ = _dafny.CodePoint('l')
                d_1639_v24_: _dafny.Map
                def iife144_():
                    coll72_ = _dafny.Map()
                    compr_72_: int
                    for compr_72_ in _dafny.IntegerRange(-974, 996):
                        d_1640_v22_: int = compr_72_
                        if ((-974) <= (d_1640_v22_)) and ((d_1640_v22_) < (996)):
                            coll72_[default__.safeModuloInt(d_1640_v22_, d_1619_v5_)] = _dafny.CodePoint('e')
                    return _dafny.Map(coll72_)
                d_1639_v24_ = _dafny.Map({_dafny.CodePoint('l'): (iife144_()
                ).set(d_1619_v5_, d_1638_v23_)})
                def iife145_():
                    coll73_ = _dafny.Set()
                    compr_73_: int
                    for compr_73_ in _dafny.IntegerRange(66, 783):
                        d_1641_v25_: int = compr_73_
                        if ((66) <= (d_1641_v25_)) and ((d_1641_v25_) < (783)):
                            coll73_ = coll73_.union(_dafny.Set([(d_1641_v25_) + (d_1619_v5_)]))
                    return _dafny.Set(coll73_)
                def iife146_():
                    coll74_ = _dafny.Set()
                    compr_74_: int
                    for compr_74_ in _dafny.IntegerRange(681, 56):
                        d_1642_v26_: int = compr_74_
                        if ((681) <= (d_1642_v26_)) and ((d_1642_v26_) < (56)):
                            coll74_ = coll74_.union(_dafny.Set([default__.safeModuloInt(d_1642_v26_, d_1619_v5_)]))
                    return _dafny.Set(coll74_)
                d_1626_v10_ = default__.fm0((self).fm2(-14, d_1639_v24_, globalState), d_1626_v10_, p2, (iife145_()
                ) | (iife146_()
                ), globalState)
            elif True:
                index265_ = default__.safeIndex(950, (d_1618_v4_).length(0))
                (d_1618_v4_)[index265_] = (d_1619_v5_) <= (d_1619_v5_)
                d_1643_v27_: _dafny.Set
                d_1643_v27_ = _dafny.Set({d_1619_v5_, len(p3)})
                d_1644_v28_: _dafny.MultiSet
                d_1644_v28_ = _dafny.MultiSet([d_1619_v5_])
                index266_ = default__.safeIndex(950, (d_1618_v4_).length(0))
                def iife147_():
                    coll75_ = _dafny.Set()
                    compr_75_: int
                    for compr_75_ in (d_1644_v28_).Elements:
                        d_1645_v29_: int = compr_75_
                        if (d_1645_v29_) in (d_1644_v28_):
                            coll75_ = coll75_.union(_dafny.Set([(d_1645_v29_) + ((0) - (-950))]))
                    return _dafny.Set(coll75_)
                rhs256_ = ((iife147_()
                ).intersection(d_1643_v27_)).ispropersubset(d_1643_v27_)
                rhs257_ = (p2) and (p2)
                lhs204_ = d_1618_v4_
                lhs205_ = default__.safeIndex(950, (d_1618_v4_).length(0))
                lhs206_ = globalState
                lhs204_[lhs205_] = rhs256_
                lhs206_.f21 = rhs257_
                d_1646_v31_: _dafny.Seq
                d_1646_v31_ = _dafny.SeqWithoutIsStrInference([D13_DC27(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sht"))]))])
                d_1647_v32_: _dafny.Map
                d_1647_v32_ = _dafny.Map({d_1619_v5_: d_1626_v10_})
                d_1648_v33_: _dafny.Map
                def iife148_():
                    coll76_ = _dafny.Map()
                    compr_76_: D13
                    for compr_76_ in (d_1646_v31_).Elements:
                        d_1649_v30_: D13 = compr_76_
                        if (d_1649_v30_) in (d_1646_v31_):
                            coll76_[d_1649_v30_] = False
                    return _dafny.Map(coll76_)
                d_1648_v33_ = _dafny.Map({iife148_()
                : len(_dafny.Map({d_1616_v2_: d_1647_v32_}))})
                d_1650_v34_: _dafny.Map
                d_1650_v34_ = _dafny.Map({default__.fm47(d_1626_v10_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p")), (d_1618_v4_)[default__.safeIndex(950, (d_1618_v4_).length(0))], globalState): (self).f27})
                d_1651_v35_: str
                d_1651_v35_ = _dafny.CodePoint('o')
                d_1652_v36_: _dafny.Seq
                d_1652_v36_ = _dafny.SeqWithoutIsStrInference([d_1616_v2_, (d_1627_v11_).set(default__.safeIndex(d_1619_v5_, len(d_1627_v11_)), d_1651_v35_)])
                d_1653_v37_: D13
                d_1653_v37_ = D13_DC27(d_1652_v36_)
                d_1648_v33_ = (d_1648_v33_).set((d_1650_v34_ if p0 else _dafny.Map({d_1653_v37_: p2})), len(d_1647_v32_))
                d_1617_v3_ = (d_1617_v3_).set((d_1618_v4_)[default__.safeIndex(950, (d_1618_v4_).length(0))], (self).f27)
                index267_ = default__.safeIndex(828, (d_1615_v1_).length(0))
                (d_1615_v1_)[index267_] = len(d_1616_v2_)
                d_1654_v38_: _dafny.Map
                d_1654_v38_ = _dafny.Map({_dafny.CodePoint('q'): (d_1619_v5_) >= (d_1619_v5_)})
                d_1655_v39_: _dafny.Map
                d_1655_v39_ = _dafny.Map({(self).f27: d_1619_v5_})
                d_1656_v40_: D18
                d_1656_v40_ = D18_DC42(d_1616_v2_, p2)
                d_1657_v41_: D18
                d_1657_v41_ = D18_DC43(len((d_1656_v40_).cf57))
                d_1658_v42_: _dafny.Map
                d_1658_v42_ = _dafny.Map({d_1657_v41_: default__.fm48(len(d_1617_v3_), not(p0), len(p3), globalState)})
                d_1659_v43_: D2
                d_1659_v43_ = D2_DC6(d_1619_v5_, d_1629_v13_, (d_1618_v4_)[default__.safeIndex(950, (d_1618_v4_).length(0))], _dafny.SeqWithoutIsStrInference([574]), d_1626_v10_)
                pat_let_tv38_ = d_1618_v4_
                pat_let_tv39_ = d_1618_v4_
                index268_ = default__.safeIndex(828, (d_1615_v1_).length(0))
                def iife149_(_pat_let36_0):
                    def iife150_(d_1660_dt__update__tmp_h0_):
                        def iife151_(_pat_let37_0):
                            def iife152_(d_1661_dt__update_hcf11_h0_):
                                return D2_DC6((d_1660_dt__update__tmp_h0_).cf7, (d_1660_dt__update__tmp_h0_).cf8, (d_1660_dt__update__tmp_h0_).cf9, (d_1660_dt__update__tmp_h0_).cf10, d_1661_dt__update_hcf11_h0_)
                            return iife152_(_pat_let37_0)
                        return iife151_((pat_let_tv39_)[default__.safeIndex(950, (pat_let_tv38_).length(0))])
                    return iife150_(_pat_let36_0)
                rhs258_ = (d_1655_v39_) | ((d_1655_v39_) | (((d_1658_v42_)[d_1657_v41_] if (d_1657_v41_) in (d_1658_v42_) else d_1655_v39_)))
                rhs259_ = d_1619_v5_
                rhs260_ = (d_1619_v5_) <= (d_1619_v5_)
                rhs261_ = d_1654_v38_
                rhs262_ = ((iife149_(d_1659_v43_) if (self).f27 else d_1659_v43_)).cf9
                lhs207_ = globalState
                lhs208_ = d_1615_v1_
                lhs209_ = default__.safeIndex(828, (d_1615_v1_).length(0))
                lhs210_ = globalState
                lhs211_ = globalState
                lhs207_.f13 = rhs258_
                lhs208_[lhs209_] = rhs259_
                lhs210_.f5 = rhs260_
                d_1654_v38_ = rhs261_
                lhs211_.f5 = rhs262_
                d_1619_v5_ = d_1619_v5_
            d_1662_v44_: D23
            d_1662_v44_ = D23_DC61(d_1615_v1_, p2, d_1619_v5_)
            d_1663_v45_: _dafny.Seq
            d_1663_v45_ = _dafny.SeqWithoutIsStrInference([d_1662_v44_, d_1662_v44_])
            d_1664_v46_: _dafny.MultiSet
            d_1664_v46_ = _dafny.MultiSet([d_1663_v45_, _dafny.SeqWithoutIsStrInference([d_1662_v44_])])
            d_1665_v47_: _dafny.Array
            nw266_ = _dafny.Array(None, 4)
            nw266_[int(0)] = (d_1664_v46_) - (d_1664_v46_)
            nw266_[int(1)] = (d_1664_v46_).intersection(d_1664_v46_)
            nw266_[int(2)] = d_1664_v46_
            nw266_[int(3)] = d_1664_v46_
            d_1665_v47_ = nw266_
            index269_ = default__.safeIndex(620, (d_1665_v47_).length(0))
            (d_1665_v47_)[index269_] = d_1664_v46_
            index270_ = default__.safeIndex(620, (d_1665_v47_).length(0))
            def iife153_(_pat_let38_0):
                def iife154_(d_1666_dt__update__tmp_h1_):
                    def iife155_(_pat_let39_0):
                        def iife156_(d_1667_dt__update_hcf95_h0_):
                            return D23_DC61((d_1666_dt__update__tmp_h1_).cf93, (d_1666_dt__update__tmp_h1_).cf94, d_1667_dt__update_hcf95_h0_)
                        return iife156_(_pat_let39_0)
                    return iife155_(798)
                return iife154_(_pat_let38_0)
            (d_1665_v47_)[index270_] = (_dafny.MultiSet([d_1663_v45_, d_1663_v45_, _dafny.SeqWithoutIsStrInference([iife153_(d_1662_v44_)]), d_1663_v45_, d_1663_v45_])).intersection(d_1664_v46_)
        elif True:
            d_1668_v48_: str
            d_1668_v48_ = _dafny.CodePoint('u')
            d_1668_v48_ = d_1668_v48_
            d_1669_v49_: D19
            d_1669_v49_ = D19_DC45(_dafny.MultiSet([d_1668_v48_]))
            d_1670_v50_: D19
            d_1670_v50_ = D19_DC47(d_1669_v49_)
            d_1671_v51_: D19
            d_1671_v51_ = D19_DC47(d_1669_v49_)
            d_1672_v52_: D19
            d_1672_v52_ = D19_DC47(d_1671_v51_)
            d_1673_v53_: D19
            d_1673_v53_ = D19_DC47(d_1670_v50_)
            d_1674_v54_: D19
            d_1674_v54_ = D19_DC47(d_1669_v49_)
            d_1675_v55_: D19
            d_1675_v55_ = D19_DC47(d_1671_v51_)
            d_1676_v56_: D19
            d_1676_v56_ = D19_DC47((d_1675_v55_).cf67)
            d_1676_v56_ = d_1675_v55_
            (globalState).f2 = d_1619_v5_
            pat_let_tv40_ = d_1626_v10_
            d_1677_v57_: bool
            d_1678_v58_: _dafny.Seq
            out68_: bool
            out69_: _dafny.Seq
            def iife157_(_pat_let40_0):
                def iife158_(d_1679_dt__update__tmp_h2_):
                    def iife159_(_pat_let41_0):
                        def iife160_(d_1680_dt__update_hcf1_h0_):
                            return D0_DC1(d_1680_dt__update_hcf1_h0_)
                        return iife160_(_pat_let41_0)
                    return iife159_(pat_let_tv40_)
                return iife158_(_pat_let40_0)
            out68_, out69_ = default__.m0(D0_DC3(D0_DC1(d_1626_v10_)), iife157_(d_1625_v9_), globalState)
            d_1677_v57_ = out68_
            d_1678_v58_ = out69_
            d_1681_v59_: _dafny.Set
            d_1681_v59_ = _dafny.Set({432, d_1619_v5_})
            def iife161_():
                coll77_ = _dafny.Set()
                compr_77_: int
                for compr_77_ in (d_1681_v59_).Elements:
                    d_1682_v60_: int = compr_77_
                    if (d_1682_v60_) in (d_1681_v59_):
                        coll77_ = coll77_.union(_dafny.Set([default__.safeModuloInt(d_1682_v60_, 31)]))
                return _dafny.Set(coll77_)
            (globalState).f14 = (iife161_()
            ).ispropersubset(d_1681_v59_)
        d_1683_v61_: _dafny.MultiSet
        d_1683_v61_ = _dafny.MultiSet([(d_1616_v2_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aapblyjy")))])
        d_1683_v61_ = (d_1683_v61_).set(d_1627_v11_, default__.abs(440))
        d_1684_v62_: D14
        d_1684_v62_ = D14_DC31(d_1619_v5_)
        d_1685_v63_: _dafny.Map
        d_1685_v63_ = _dafny.Map({p2: d_1684_v62_})
        r0 = (d_1629_v13_) + ((default__.fm40(len(d_1685_v63_), p2, globalState)) + (_dafny.SeqWithoutIsStrInference([d_1619_v5_])))
        d_1686_v64_: _dafny.Array
        nw267_ = _dafny.Array(_dafny.Map({}), 20)
        d_1686_v64_ = nw267_
        d_1687_v65_: D27
        d_1687_v65_ = D27_DC66(d_1686_v64_)
        r1 = (d_1687_v65_).cf104
        d_1688_v66_: _dafny.MultiSet
        d_1688_v66_ = _dafny.MultiSet([d_1619_v5_])
        r2 = (d_1688_v66_).ispropersubset((d_1688_v66_).set(d_1619_v5_, default__.abs(d_1619_v5_)))
        r3 = d_1626_v10_
        return r0, r1, r2, r3

    @property
    def f37(self):
        return self._f37

class C5(T1, T0):
    def  __init__(self):
        self._f27: bool = False
        self._f28: int = int(0)
        self._f29: int = int(0)
        self._f39: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f27(self):
        return self._f27
    @property
    def f28(self):
        return self._f28
    @property
    def f29(self):
        return self._f29
    def ctor__(self, f39, f28, f29, f27):
        (self)._f39 = f39
        (self)._f28 = f28
        (self)._f29 = f29
        (self)._f27 = f27

    def fm2(self, p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yrqenf"))

    def fm3(self, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kkcp"))

    def m3(self, p0, p1, p2, globalState):
        d_1689_v0_: _dafny.Array
        def lambda82_(d_1690_i1_):
            return (d_1690_i1_) * ((self).f29)

        init42_ = lambda82_
        nw268_ = _dafny.Array(None, 21)
        for i0_42_ in range(nw268_.length(0)):
            nw268_[i0_42_] = init42_(i0_42_)
        d_1689_v0_ = nw268_
        guard_loop_8_: int
        for guard_loop_8_ in _dafny.IntegerRange(0, (d_1689_v0_).length(0)):
            d_1691_i0_: int = guard_loop_8_
            if (True) and (((0) <= (d_1691_i0_)) and ((d_1691_i0_) < ((d_1689_v0_).length(0)))):
                (d_1689_v0_)[(d_1691_i0_)] = (d_1691_i0_) - ((0) - (((self).f29) - ((self).f29)))
        d_1692_v1_: _dafny.Seq
        d_1692_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "amkfn"))
        d_1693_v2_: str
        d_1693_v2_ = _dafny.CodePoint('g')
        d_1694_v3_: _dafny.Map
        d_1694_v3_ = _dafny.Map({d_1693_v2_: p0})
        d_1695_v4_: _dafny.Map
        d_1695_v4_ = _dafny.Map({(self).f29: (self).f27})
        d_1696_v5_: _dafny.Set
        d_1696_v5_ = _dafny.Set({(self).f28, p2})
        if default__.fm0(d_1692_v1_, not(((d_1694_v3_)[d_1693_v2_] if (d_1693_v2_) in (d_1694_v3_) else (self).f39)), True, (d_1696_v5_ if ((d_1695_v4_)[702] if (702) in (d_1695_v4_) else (self).f39) else d_1696_v5_), globalState):
            d_1697_v6_: _dafny.Map
            d_1697_v6_ = _dafny.Map({(self).f27: len(_dafny.Map({len(d_1695_v4_): p2}))})
            d_1698_v7_: _dafny.Map
            d_1698_v7_ = _dafny.Map({(self).f29: d_1697_v6_})
            (globalState).f13 = (((d_1698_v7_)[(self).f28] if ((self).f28) in (d_1698_v7_) else d_1697_v6_)) | (d_1697_v6_)
            d_1699_v8_: D21
            d_1699_v8_ = D21_DC55(p2)
            d_1700_v9_: _dafny.MultiSet
            d_1700_v9_ = _dafny.MultiSet([D21_DC55((self).f28), d_1699_v8_])
            def iife162_(_pat_let42_0):
                def iife163_(d_1701_dt__update__tmp_h0_):
                    def iife164_(_pat_let43_0):
                        def iife165_(d_1702_dt__update_hcf82_h0_):
                            return D21_DC55(d_1702_dt__update_hcf82_h0_)
                        return iife165_(_pat_let43_0)
                    return iife164_((self).f28)
                return iife163_(_pat_let42_0)
            (globalState).f2 = (0) - ((0) - ((((d_1700_v9_) | ((d_1700_v9_).set(iife162_(d_1699_v8_), default__.abs((self).f28)))) | (default__.fm49(not((self).f39), (self).f39, (self).f39, globalState))).cardinality))
            d_1703_v10_: _dafny.Seq
            d_1703_v10_ = _dafny.SeqWithoutIsStrInference([(self).f39, p0])
            d_1704_v11_: _dafny.Seq
            d_1704_v11_ = _dafny.SeqWithoutIsStrInference([(d_1703_v10_)[default__.safeIndex((self).f29, len(d_1703_v10_))], (self).f27, p0])
            d_1705_v12_: _dafny.Map
            d_1705_v12_ = _dafny.Map({(self).f29: d_1703_v10_})
            (globalState).f2 = default__.fm1((p2) == ((self).f28), ((d_1705_v12_)[len(d_1692_v1_)] if (len(d_1692_v1_)) in (d_1705_v12_) else d_1704_v11_), (self).f27, globalState)
            (globalState).f14 = default__.fm0(d_1692_v1_, p0, (d_1704_v11_)[default__.safeIndex((self).f29, len(d_1704_v11_))], d_1696_v5_, globalState)
            d_1706_v13_: _dafny.Map
            d_1706_v13_ = _dafny.Map({(self).f27: d_1696_v5_})
            d_1696_v5_ = (d_1696_v5_ if (self).f39 else (((d_1706_v13_)[p0] if (p0) in (d_1706_v13_) else d_1696_v5_)) | (d_1696_v5_))
        elif True:
            (globalState).f6 = (self).f28
            d_1707_v14_: D0
            d_1707_v14_ = D0_DC2(len(d_1692_v1_), (self).f28)
            d_1708_v15_: D0
            d_1708_v15_ = D0_DC3(d_1707_v14_)
            d_1709_v16_: D0
            d_1709_v16_ = D0_DC1(False)
            d_1710_v17_: bool
            d_1711_v18_: _dafny.Seq
            out70_: bool
            out71_: _dafny.Seq
            out70_, out71_ = default__.m0(d_1708_v15_, d_1709_v16_, globalState)
            d_1710_v17_ = out70_
            d_1711_v18_ = out71_
            (globalState).f6 = default__.safeModuloInt(len(d_1711_v18_), (self).f29)
            d_1712_v19_: _dafny.Seq
            d_1712_v19_ = _dafny.SeqWithoutIsStrInference([d_1710_v17_, p0])
            d_1713_v20_: D13
            d_1713_v20_ = D13_DC29(p2, (self).f28, (self).f39)
            d_1714_v21_: _dafny.Map
            d_1714_v21_ = _dafny.Map({d_1710_v17_: default__.fm1((self).f39, d_1712_v19_, (d_1713_v20_).cf41, globalState)})
            d_1715_v22_: _dafny.Map
            d_1715_v22_ = _dafny.Map({d_1714_v21_: d_1693_v2_})
            d_1716_v23_: _dafny.MultiSet
            d_1716_v23_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eumjd")), d_1692_v1_])
            d_1715_v22_ = _dafny.Map({(d_1714_v21_) | (_dafny.Map({(self).f27: ((d_1716_v23_)[d_1711_v18_] if (d_1711_v18_) in (d_1716_v23_) else (0) - (p2))})): d_1693_v2_})
            if (self).f27:
                d_1717_v24_: _dafny.MultiSet
                d_1717_v24_ = _dafny.MultiSet([d_1710_v17_, p0, (self).f39, p0, d_1710_v17_])
                d_1718_v25_: _dafny.Set
                d_1718_v25_ = _dafny.Set({p0})
                index271_ = default__.safeIndex(886, (d_1689_v0_).length(0))
                (d_1689_v0_)[index271_] = ((d_1717_v24_)[(self).f27] if ((self).f27) in (d_1717_v24_) else len(d_1718_v25_))
                index272_ = default__.safeIndex(886, (d_1689_v0_).length(0))
                (d_1689_v0_)[index272_] = ((self).f29) * ((self).f29)
                d_1719_v26_: C0
                nw269_ = C0()
                nw269_.ctor__()
                d_1719_v26_ = nw269_
                d_1720_v27_: C1
                nw270_ = C1()
                nw270_.ctor__((self).f39)
                d_1720_v27_ = nw270_
                d_1714_v21_ = d_1714_v21_
                d_1721_v28_: D18
                d_1721_v28_ = D18_DC42(d_1692_v1_, p0)
                d_1711_v18_ = ((d_1721_v28_).cf57) + (d_1711_v18_)
            elif True:
                d_1722_v29_: _dafny.Seq
                d_1722_v29_ = _dafny.SeqWithoutIsStrInference([d_1689_v0_])
                d_1695_v4_ = (d_1695_v4_).set((self).f29, (_dafny.SeqWithoutIsStrInference([d_1689_v0_])) < (d_1722_v29_))
                d_1723_v30_: _dafny.Map
                d_1723_v30_ = _dafny.Map({(p2) - ((self).f29): (self).f28})
                d_1723_v30_ = default__.fm50(p2, (self).f27, (d_1711_v18_) + (d_1711_v18_), globalState)
                d_1724_v31_: _dafny.Array
                def lambda83_(d_1725_v17_):
                    def lambda84_(d_1726_i2_):
                        return d_1725_v17_

                    return lambda84_

                init43_ = lambda83_(d_1710_v17_)
                nw271_ = _dafny.Array(None, 9)
                for i0_43_ in range(nw271_.length(0)):
                    nw271_[i0_43_] = init43_(i0_43_)
                d_1724_v31_ = nw271_
                index273_ = default__.safeIndex(251, (d_1724_v31_).length(0))
                (d_1724_v31_)[index273_] = default__.fm0(d_1711_v18_, d_1710_v17_, p0, _dafny.Set({len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vpnirac"))).set(default__.safeIndex((self).f28, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vpnirac")))), d_1693_v2_)), (self).f28}), globalState)
                index274_ = default__.safeIndex(251, (d_1724_v31_).length(0))
                (d_1724_v31_)[index274_] = ((94) - ((self).f28)) == (((default__.fm51((self).f29, globalState)).cf32) + ((0) - ((self).f29)))
                d_1727_v32_: _dafny.Set
                d_1727_v32_ = _dafny.Set({d_1710_v17_})
                d_1728_v33_: _dafny.Seq
                d_1728_v33_ = _dafny.SeqWithoutIsStrInference([d_1727_v32_, d_1727_v32_])
                (globalState).f25 = not ((d_1727_v32_) not in (d_1728_v33_)) or ((self).f39)
                d_1729_v34_: C0
                nw272_ = C0()
                nw272_.ctor__()
                d_1729_v34_ = nw272_
        (globalState).f12 = ((p2) * (p2)) >= ((0) - (p2))
        d_1730_v35_: D16
        d_1730_v35_ = D16_DC36(d_1689_v0_)
        source26_ = d_1730_v35_
        if source26_.is_DC36:
            d_1731___mcc_h0_ = source26_.cf51
            d_1732_cf51_ = d_1731___mcc_h0_
            d_1733_v36_: _dafny.Seq
            d_1733_v36_ = _dafny.SeqWithoutIsStrInference([(self).f28])
            d_1734_v37_: _dafny.Set
            d_1734_v37_ = _dafny.Set({d_1733_v36_})
            d_1735_v38_: D27
            d_1735_v38_ = D27_DC67(_dafny.SeqWithoutIsStrInference([d_1732_cf51_, d_1689_v0_, d_1732_cf51_]), d_1734_v37_, (self).f39)
            d_1736_v39_: _dafny.Array
            def lambda85_(d_1737_v36_):
                def lambda86_(d_1738_i3_):
                    return d_1737_v36_

                return lambda86_

            init44_ = lambda85_(d_1733_v36_)
            nw273_ = _dafny.Array(None, 29)
            for i0_44_ in range(nw273_.length(0)):
                nw273_[i0_44_] = init44_(i0_44_)
            d_1736_v39_ = nw273_
            d_1739_v40_: _dafny.Map
            d_1739_v40_ = _dafny.Map({(d_1735_v38_).cf107: d_1736_v39_})
            d_1740_v41_: _dafny.Seq
            d_1740_v41_ = _dafny.SeqWithoutIsStrInference([p0, (self).f27])
            d_1741_v42_: _dafny.Map
            d_1741_v42_ = _dafny.Map({((d_1739_v40_)[(self).f39] if ((self).f39) in (d_1739_v40_) else d_1736_v39_): len(d_1740_v41_)})
            d_1741_v42_ = (d_1741_v42_).set((d_1736_v39_ if p0 else d_1736_v39_), default__.safeModuloInt((self).f28, len(d_1692_v1_)))
            (globalState).f2 = ((self).f28) + (p2)
            (globalState).f2 = (self).f29
            (globalState).f25 = not((self).f27)
        elif source26_.is_DC35:
            d_1742___mcc_h1_ = source26_.cf50
            d_1743_cf50_ = d_1742___mcc_h1_
            d_1744_v43_: T0
            nw274_ = C3()
            nw274_.ctor__((self).f39, p0)
            d_1744_v43_ = nw274_
            d_1745_v44_: D20
            d_1745_v44_ = D20_DC49(d_1693_v2_, (self).f29, p2, d_1744_v43_, d_1696_v5_)
            source27_ = d_1745_v44_
            if source27_.is_DC49:
                d_1746___mcc_h3_ = source27_.cf69
                d_1747___mcc_h4_ = source27_.cf70
                d_1748___mcc_h5_ = source27_.cf71
                d_1749___mcc_h6_ = source27_.cf72
                d_1750___mcc_h7_ = source27_.cf73
                d_1751_cf73_ = d_1750___mcc_h7_
                d_1752_cf72_ = d_1749___mcc_h6_
                d_1753_cf71_ = d_1748___mcc_h5_
                d_1754_cf70_ = d_1747___mcc_h4_
                d_1755_cf69_ = d_1746___mcc_h3_
                index275_ = default__.safeIndex(758, (d_1689_v0_).length(0))
                (d_1689_v0_)[index275_] = p2
                index276_ = default__.safeIndex(758, (d_1689_v0_).length(0))
                (d_1689_v0_)[index276_] = (self).f28
                d_1756_v45_: C2
                nw275_ = C2()
                nw275_.ctor__((d_1752_cf72_).f27, (d_1752_cf72_).f27)
                d_1756_v45_ = nw275_
                d_1757_v46_: _dafny.Array
                nw276_ = _dafny.Array(_dafny.Seq({}), 7)
                d_1757_v46_ = nw276_
                d_1758_v47_: _dafny.Seq
                d_1758_v47_ = _dafny.SeqWithoutIsStrInference([(d_1744_v43_).f27, (d_1756_v45_).f34])
                d_1759_v48_: _dafny.Seq
                d_1759_v48_ = _dafny.SeqWithoutIsStrInference([(d_1758_v47_)[default__.safeIndex(p2, len(d_1758_v47_))], (d_1756_v45_).f34, (d_1744_v43_).f27])
                index277_ = default__.safeIndex(956, (d_1757_v46_).length(0))
                (d_1757_v46_)[index277_] = d_1759_v48_
                d_1760_v49_: _dafny.MultiSet
                d_1760_v49_ = _dafny.MultiSet([(d_1756_v45_).f34])
                d_1761_v50_: _dafny.Seq
                d_1761_v50_ = _dafny.SeqWithoutIsStrInference([d_1754_cf70_])
                d_1762_v51_: _dafny.Map
                d_1762_v51_ = _dafny.Map({d_1761_v50_: (self).f27})
                index278_ = default__.safeIndex(956, (d_1757_v46_).length(0))
                (d_1757_v46_)[index278_] = ((default__.fm52(-339, (d_1756_v45_).f34, d_1760_v49_, globalState)) + (_dafny.SeqWithoutIsStrInference([((d_1762_v51_)[d_1761_v50_] if (d_1761_v50_) in (d_1762_v51_) else (self).f39), True, not((self).f39), (self).f39]))) + ((d_1759_v48_) + (_dafny.SeqWithoutIsStrInference([(self).f39, (self).f39])))
                d_1763_v52_: _dafny.Map
                d_1763_v52_ = _dafny.Map({p2: d_1760_v49_})
                d_1764_v53_: _dafny.Array
                nw277_ = _dafny.Array(None, 10)
                nw277_[int(0)] = p2
                nw277_[int(1)] = d_1753_cf71_
                nw277_[int(2)] = (self).f28
                nw277_[int(3)] = (0) - ((self).f28)
                nw277_[int(4)] = len((d_1757_v46_)[default__.safeIndex(956, (d_1757_v46_).length(0))])
                nw277_[int(5)] = 795
                nw277_[int(6)] = p2
                nw277_[int(7)] = len(d_1759_v48_)
                nw277_[int(8)] = d_1753_cf71_
                nw277_[int(9)] = d_1754_cf70_
                d_1764_v53_ = nw277_
                d_1765_v54_: _dafny.Map
                d_1765_v54_ = _dafny.Map({((d_1763_v52_)[d_1753_cf71_] if (d_1753_cf71_) in (d_1763_v52_) else d_1760_v49_): d_1764_v53_})
                d_1765_v54_ = (d_1765_v54_).set(d_1760_v49_, d_1764_v53_)
            elif source27_.is_DC50:
                d_1766_v55_: _dafny.Array
                nw278_ = _dafny.Array(D7.default()(), 5)
                d_1766_v55_ = nw278_
                d_1767_v56_: _dafny.Set
                d_1767_v56_ = _dafny.Set({(d_1766_v55_ if p0 else d_1766_v55_)})
                d_1767_v56_ = d_1767_v56_
                d_1768_v57_: _dafny.Array
                def lambda87_(d_1769_v4_):
                    def lambda88_(d_1770_i4_):
                        return d_1769_v4_

                    return lambda88_

                init45_ = lambda87_(d_1695_v4_)
                nw279_ = _dafny.Array(None, 6)
                for i0_45_ in range(nw279_.length(0)):
                    nw279_[i0_45_] = init45_(i0_45_)
                d_1768_v57_ = nw279_
                d_1771_v58_: _dafny.Seq
                d_1771_v58_ = _dafny.SeqWithoutIsStrInference([(d_1744_v43_).f27])
                d_1772_v59_: _dafny.Seq
                d_1772_v59_ = _dafny.SeqWithoutIsStrInference([(self).f29, 499])
                rhs263_ = (d_1745_v44_).cf69
                rhs264_ = ((len(d_1771_v58_) if (d_1744_v43_).f27 else (self).f29)) * (default__.safeDivisionInt((d_1772_v59_)[default__.safeIndex((self).f29, len(d_1772_v59_))], (0) - (len(d_1771_v58_))))
                rhs265_ = (self).f28
                rhs266_ = d_1768_v57_
                lhs212_ = globalState
                lhs213_ = globalState
                d_1693_v2_ = rhs263_
                lhs212_.f2 = rhs264_
                lhs213_.f2 = rhs265_
                d_1768_v57_ = rhs266_
                d_1773_v60_: _dafny.Map
                d_1773_v60_ = _dafny.Map({not((self).f27): (d_1744_v43_).f27})
                d_1774_v61_: _dafny.Map
                d_1774_v61_ = _dafny.Map({p0: ((self).f28) != (len(d_1773_v60_))})
                d_1775_v62_: _dafny.Set
                d_1775_v62_ = _dafny.Set({(self).f39})
                d_1773_v60_ = (d_1773_v60_).set((d_1775_v62_).isdisjoint(d_1775_v62_), default__.fm0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "clfwfundw")), (d_1744_v43_).f27, (d_1744_v43_).f27, _dafny.Set({p2}), globalState))
                (globalState).f21 = ((d_1744_v43_).f27 if not((d_1744_v43_).f27) else (True if p0 else not(p0)))
            elif source27_.is_DC48:
                d_1776___mcc_h8_ = source27_.cf68
                d_1777_cf68_ = d_1776___mcc_h8_
                index279_ = default__.safeIndex(459, (d_1689_v0_).length(0))
                (d_1689_v0_)[index279_] = (self).f29
                d_1778_v63_: _dafny.Array
                def lambda89_(d_1779_p2_):
                    def lambda90_(d_1780_i5_):
                        return (d_1779_p2_) >= (796)

                    return lambda90_

                init46_ = lambda89_(p2)
                nw280_ = _dafny.Array(None, 1)
                for i0_46_ in range(nw280_.length(0)):
                    nw280_[i0_46_] = init46_(i0_46_)
                d_1778_v63_ = nw280_
                index280_ = default__.safeIndex(227, (d_1778_v63_).length(0))
                (d_1778_v63_)[index280_] = (d_1744_v43_).f27
                d_1781_v64_: _dafny.Seq
                d_1781_v64_ = _dafny.SeqWithoutIsStrInference([(self).f28, len(_dafny.Map({(self).f27: (self).f28}))])
                d_1782_v65_: _dafny.Map
                d_1782_v65_ = _dafny.Map({d_1693_v2_: d_1693_v2_})
                index281_ = default__.safeIndex(459, (d_1689_v0_).length(0))
                index282_ = default__.safeIndex(227, (d_1778_v63_).length(0))
                rhs267_ = (default__.safeDivisionInt(len(d_1692_v1_), (d_1781_v64_)[default__.safeIndex(p2, len(d_1781_v64_))])) + ((self).f29)
                rhs268_ = (((d_1782_v65_)[d_1693_v2_] if (d_1693_v2_) in (d_1782_v65_) else d_1693_v2_) if not(not(not(p0))) else d_1693_v2_)
                rhs269_ = d_1693_v2_
                rhs270_ = ((self).f29) == (784)
                rhs271_ = len(_dafny.SeqWithoutIsStrInference([d_1693_v2_ for d_1783_i6_ in range(default__.abs(243))]))
                lhs214_ = d_1689_v0_
                lhs215_ = default__.safeIndex(459, (d_1689_v0_).length(0))
                lhs216_ = d_1778_v63_
                lhs217_ = default__.safeIndex(227, (d_1778_v63_).length(0))
                lhs218_ = globalState
                lhs214_[lhs215_] = rhs267_
                d_1693_v2_ = rhs268_
                d_1693_v2_ = rhs269_
                lhs216_[lhs217_] = rhs270_
                lhs218_.f2 = rhs271_
                d_1781_v64_ = d_1781_v64_
                (globalState).f6 = p2
                d_1784_v66_: _dafny.MultiSet
                d_1784_v66_ = _dafny.MultiSet([(d_1744_v43_).f27])
                d_1785_v67_: C4
                nw281_ = C4()
                nw281_.ctor__(d_1784_v66_, (d_1744_v43_).f27)
                d_1785_v67_ = nw281_
            elif True:
                d_1786___mcc_h9_ = source27_.cf74
                d_1787_cf74_ = d_1786___mcc_h9_
                d_1788_v68_: _dafny.MultiSet
                d_1788_v68_ = _dafny.MultiSet([p0, (self).f39])
                d_1789_v69_: C4
                nw282_ = C4()
                nw282_.ctor__(d_1788_v68_, ((self).f39) and ((d_1744_v43_).f27))
                d_1789_v69_ = nw282_
                d_1790_v71_: _dafny.Array
                def lambda91_(d_1791_v43_):
                    def lambda92_(d_1792_i7_):
                        return not(not((d_1791_v43_).f27))

                    return lambda92_

                init47_ = lambda91_(d_1744_v43_)
                nw283_ = _dafny.Array(None, 29)
                for i0_47_ in range(nw283_.length(0)):
                    nw283_[i0_47_] = init47_(i0_47_)
                d_1790_v71_ = nw283_
                d_1793_v72_: _dafny.Map
                def iife166_():
                    coll78_ = _dafny.Set()
                    compr_78_: int
                    for compr_78_ in (d_1695_v4_).keys.Elements:
                        d_1794_v70_: int = compr_78_
                        if (d_1794_v70_) in (d_1695_v4_):
                            coll78_ = coll78_.union(_dafny.Set([(d_1794_v70_) + (len(_dafny.Map({359: _dafny.MultiSet([920])})))]))
                    return _dafny.Set(coll78_)
                d_1793_v72_ = _dafny.Map({(_dafny.Set({(self).f29, -293})) | (iife166_()
                ): d_1790_v71_})
                d_1795_v73_: _dafny.Map
                d_1795_v73_ = d_1793_v72_
                d_1793_v72_ = (d_1795_v73_)
                (globalState).f6 = default__.safeModuloInt((self).f28, (self).f29)
                d_1796_v74_: _dafny.Seq
                d_1796_v74_ = _dafny.SeqWithoutIsStrInference([(self).f29])
                d_1797_v75_: _dafny.MultiSet
                d_1797_v75_ = _dafny.MultiSet([d_1796_v74_])
                d_1797_v75_ = (_dafny.MultiSet([d_1796_v74_])) | (d_1797_v75_)
            d_1798_v76_: _dafny.Array
            nw284_ = _dafny.Array(_dafny.Map({}), 10)
            d_1798_v76_ = nw284_
            nw285_ = _dafny.Array(_dafny.Map({}), 14)
            d_1798_v76_ = nw285_
            d_1799_v77_: _dafny.Seq
            d_1799_v77_ = _dafny.SeqWithoutIsStrInference([(d_1744_v43_).f27, (d_1744_v43_).f27])
            d_1800_v78_: _dafny.MultiSet
            d_1800_v78_ = _dafny.MultiSet([d_1689_v0_])
            d_1801_v79_: _dafny.Seq
            d_1801_v79_ = _dafny.SeqWithoutIsStrInference([(0) - (len(d_1692_v1_)), (d_1800_v78_).cardinality])
            d_1802_v80_: _dafny.Seq
            d_1802_v80_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).f28), len(d_1801_v79_), (_dafny.MultiSet([(self).f29, (0) - ((self).f29)])).cardinality])
            d_1803_v81_: _dafny.Map
            d_1803_v81_ = _dafny.Map({(self).f39: ((0) - ((0) - (default__.fm1((d_1744_v43_).f27, d_1799_v77_, p0, globalState)))) * (len(d_1802_v80_))})
            (globalState).f13 = d_1803_v81_
            d_1804_v82_: D0
            d_1804_v82_ = D0_DC1((self).f39)
            d_1805_v83_: D0
            d_1805_v83_ = D0_DC3(d_1804_v82_)
            d_1806_v84_: _dafny.Map
            d_1806_v84_ = _dafny.Map({d_1805_v83_: (d_1744_v43_).f27})
            d_1807_v85_: _dafny.Set
            d_1807_v85_ = _dafny.Set({((d_1806_v84_)[d_1805_v83_] if (d_1805_v83_) in (d_1806_v84_) else (self).f39)})
            if (d_1807_v85_) != ((d_1807_v85_).intersection(d_1807_v85_)):
                d_1808_v86_: D0
                d_1808_v86_ = D0_DC1((self).f39)
                d_1809_v87_: _dafny.MultiSet
                d_1809_v87_ = _dafny.MultiSet([d_1808_v86_])
                (globalState).f6 = (((d_1809_v87_)[d_1808_v86_] if (d_1808_v86_) in (d_1809_v87_) else p2)) * ((self).f29)
                d_1810_v88_: _dafny.Array
                nw286_ = _dafny.Array(_dafny.Map({}), 18)
                d_1810_v88_ = nw286_
                d_1811_v89_: _dafny.MultiSet
                d_1811_v89_ = _dafny.MultiSet([(self).f27, (self).f27])
                d_1812_v90_: _dafny.Map
                d_1812_v90_ = _dafny.Map({d_1692_v1_: d_1811_v89_})
                d_1813_v91_: _dafny.Seq
                d_1813_v91_ = _dafny.SeqWithoutIsStrInference([d_1812_v90_, d_1812_v90_, d_1812_v90_])
                index283_ = default__.safeIndex(988, (d_1810_v88_).length(0))
                (d_1810_v88_)[index283_] = (d_1813_v91_)[default__.safeIndex(p2, len(d_1813_v91_))]
                index284_ = default__.safeIndex(988, (d_1810_v88_).length(0))
                (d_1810_v88_)[index284_] = (d_1812_v90_).set((d_1692_v1_).set(default__.safeIndex(640, len(d_1692_v1_)), d_1693_v2_), _dafny.MultiSet([(d_1744_v43_).f27, (d_1744_v43_).f27, p0, True, (d_1744_v43_).f27]))
                d_1814_v92_: C1
                nw287_ = C1()
                nw287_.ctor__((d_1744_v43_).f27)
                d_1814_v92_ = nw287_
                d_1815_v93_: C0
                nw288_ = C0()
                nw288_.ctor__()
                d_1815_v93_ = nw288_
                d_1816_v94_: _dafny.Array
                def lambda93_(d_1817_v43_):
                    def lambda94_(d_1818_i8_):
                        return (d_1817_v43_).f27

                    return lambda94_

                init48_ = lambda93_(d_1744_v43_)
                nw289_ = _dafny.Array(None, 25)
                for i0_48_ in range(nw289_.length(0)):
                    nw289_[i0_48_] = init48_(i0_48_)
                d_1816_v94_ = nw289_
                nw290_ = _dafny.Array(False, 17)
                d_1816_v94_ = nw290_
            elif True:
                index285_ = default__.safeIndex(479, (p1).length(0))
                (p1)[index285_] = d_1693_v2_
                index286_ = default__.safeIndex(479, (p1).length(0))
                (p1)[index286_] = d_1693_v2_
                d_1819_v95_: _dafny.Map
                d_1819_v95_ = _dafny.Map({d_1802_v80_: p2})
                d_1820_v96_: D3
                d_1820_v96_ = D3_DC8(d_1819_v95_)
                d_1821_v97_: _dafny.Map
                d_1821_v97_ = _dafny.Map({d_1820_v96_: d_1743_cf50_})
                d_1822_v98_: D20
                d_1822_v98_ = D20_DC48(d_1821_v97_)
                d_1823_v99_: _dafny.Map
                d_1823_v99_ = _dafny.Map({(p1)[default__.safeIndex(479, (p1).length(0))]: len(d_1692_v1_)})
                d_1824_v100_: D29
                d_1824_v100_ = D29_DC69(d_1823_v99_)
                d_1825_v101_: D18
                d_1825_v101_ = D18_DC44(p0, default__.fm1((d_1744_v43_).f27, d_1799_v77_, (d_1744_v43_).f27, globalState), True)
                d_1826_v102_: D15
                d_1826_v102_ = D15_DC34(187, (self).f27, len((d_1824_v100_).cf109), (self).f28, (d_1825_v101_).cf62)
                rhs272_ = (d_1826_v102_).cf49
                rhs273_ = d_1822_v98_
                lhs219_ = globalState
                lhs219_.f14 = rhs272_
                d_1822_v98_ = rhs273_
                (globalState).f6 = (self).f28
                (globalState).f2 = (self).f29
                d_1827_v103_: _dafny.Map
                d_1827_v103_ = _dafny.Map({_dafny.Set({(self).f27}): (self).f29})
                d_1828_v104_: D22
                d_1828_v104_ = D22_DC57((d_1744_v43_).f27, (p1)[default__.safeIndex(479, (p1).length(0))], p0)
                d_1695_v4_ = (d_1695_v4_).set(len(d_1827_v103_), ((0) - ((self).f28)) >= ((_dafny.MultiSet([d_1828_v104_, D22_DC57(not(not(p0)), d_1693_v2_, False)])).cardinality))
        elif True:
            d_1829___mcc_h2_ = source26_.cf52
            d_1830_cf52_ = d_1829___mcc_h2_
            (globalState).f2 = default__.safeModuloInt((len(d_1692_v1_)) + (417), (self).f28)
            d_1693_v2_ = d_1693_v2_
            rhs274_ = p0
            rhs275_ = d_1695_v4_
            lhs220_ = globalState
            lhs220_.f12 = rhs274_
            d_1695_v4_ = rhs275_
            d_1831_v105_: _dafny.Array
            nw291_ = _dafny.Array(_dafny.Array(None, 0), 18)
            d_1831_v105_ = nw291_
            d_1832_v106_: _dafny.Map
            d_1832_v106_ = _dafny.Map({True: (self).f27})
            d_1833_v107_: _dafny.Seq
            d_1833_v107_ = _dafny.SeqWithoutIsStrInference([d_1832_v106_, d_1832_v106_])
            d_1834_v108_: _dafny.Array
            nw292_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 29)
            d_1834_v108_ = nw292_
            d_1835_v109_: _dafny.Map
            d_1835_v109_ = _dafny.Map({d_1833_v107_: d_1834_v108_})
            index287_ = default__.safeIndex(383, (d_1831_v105_).length(0))
            (d_1831_v105_)[index287_] = ((d_1835_v109_)[_dafny.SeqWithoutIsStrInference([_dafny.Map({p0: (self).f39})])] if (_dafny.SeqWithoutIsStrInference([_dafny.Map({p0: (self).f39})])) in (d_1835_v109_) else d_1834_v108_)
            index288_ = default__.safeIndex(383, (d_1831_v105_).length(0))
            (d_1831_v105_)[index288_] = d_1834_v108_
        d_1836_v110_: C1
        nw293_ = C1()
        nw293_.ctor__(((self).f27) not in ((default__.fm53(d_1693_v2_, p2, p0, (self).f29, globalState)).cf21))
        d_1836_v110_ = nw293_
        d_1837_v111_: _dafny.Array
        nw294_ = _dafny.Array(False, 7)
        d_1837_v111_ = nw294_
        d_1838_v112_: _dafny.Seq
        d_1838_v112_ = _dafny.SeqWithoutIsStrInference([not(p0), False])
        d_1839_v113_: _dafny.Seq
        d_1839_v113_ = _dafny.SeqWithoutIsStrInference([p2])
        index289_ = default__.safeIndex(77, (d_1837_v111_).length(0))
        (d_1837_v111_)[index289_] = (d_1838_v112_)[default__.safeIndex((d_1839_v113_)[default__.safeIndex(p2, len(d_1839_v113_))], len(d_1838_v112_))]
        index290_ = default__.safeIndex(77, (d_1837_v111_).length(0))
        (d_1837_v111_)[index290_] = (self).f39

    def m1(self, globalState):
        r0: int = int(0)
        d_1840_v0_: _dafny.Map
        d_1840_v0_ = _dafny.Map({(self).f29: (self).f27})
        d_1841_v1_: D7
        d_1841_v1_ = D7_DC17((self).f29, 245, True)
        d_1842_v2_: _dafny.Seq
        d_1842_v2_ = _dafny.SeqWithoutIsStrInference([(d_1841_v1_).cf27, (self).f39])
        d_1840_v0_ = (d_1840_v0_).set(((self).f28) * ((0) - ((self).f28)), (d_1842_v2_)[default__.safeIndex((self).f29, len(d_1842_v2_))])
        d_1843_v3_: D6
        d_1843_v3_ = D6_DC13(d_1842_v2_)
        def lambda95_(source29_):
            if source29_.is_DC14:
                d_1844___mcc_h5_ = source29_.cf22
                d_1845___mcc_h6_ = source29_.cf23
                d_1846_cf23_ = d_1845___mcc_h6_
                d_1847_cf22_ = d_1844___mcc_h5_
                return D19_DC45(_dafny.MultiSet((D18_DC42(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_1848_i0_ in range(default__.abs(961))]), (self).f39)).cf57))
            elif True:
                d_1849___mcc_h7_ = source29_.cf21
                d_1850_cf21_ = d_1849___mcc_h7_
                return D19_DC45(_dafny.MultiSet([_dafny.CodePoint('l')]))

        source28_ = lambda95_(d_1843_v3_)
        if source28_.is_DC46:
            d_1851___mcc_h0_ = source28_.cf64
            d_1852___mcc_h1_ = source28_.cf65
            d_1853___mcc_h2_ = source28_.cf66
            d_1854_cf66_ = d_1853___mcc_h2_
            d_1855_cf65_ = d_1852___mcc_h1_
            d_1856_cf64_ = d_1851___mcc_h0_
            d_1857_v4_: _dafny.Array
            nw295_ = _dafny.Array(False, 20)
            d_1857_v4_ = nw295_
            index291_ = default__.safeIndex(608, (d_1857_v4_).length(0))
            (d_1857_v4_)[index291_] = d_1854_cf66_
            index292_ = default__.safeIndex(608, (d_1857_v4_).length(0))
            (d_1857_v4_)[index292_] = True
            (globalState).f6 = (0) - (d_1856_cf64_)
            (globalState).f6 = (self).f29
            d_1858_v5_: str
            d_1858_v5_ = _dafny.CodePoint('u')
            d_1859_v6_: _dafny.Seq
            d_1859_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "obgewrx"))
            d_1860_v7_: _dafny.Seq
            d_1860_v7_ = _dafny.SeqWithoutIsStrInference([len(d_1859_v6_), (self).f28])
            d_1861_v8_: _dafny.MultiSet
            d_1861_v8_ = _dafny.MultiSet([d_1860_v7_])
            d_1862_v9_: _dafny.Set
            d_1862_v9_ = _dafny.Set({(self).f29})
            d_1863_v10_: _dafny.Map
            d_1863_v10_ = _dafny.Map({(d_1861_v8_).cardinality: len(d_1862_v9_)})
            d_1864_v11_: D5
            d_1864_v11_ = D5_DC12(d_1854_cf66_, len(d_1862_v9_))
            d_1865_v12_: _dafny.Map
            d_1865_v12_ = _dafny.Map({(d_1864_v11_).cf19: d_1860_v7_})
            d_1866_v13_: _dafny.Set
            d_1866_v13_ = _dafny.Set({d_1854_cf66_})
            d_1867_v14_: _dafny.Map
            d_1867_v14_ = _dafny.Map({d_1858_v5_: ((d_1865_v12_)[(self).f27] if ((self).f27) in (d_1865_v12_) else _dafny.SeqWithoutIsStrInference([default__.fm1((self).f39, _dafny.SeqWithoutIsStrInference([(d_1857_v4_)[default__.safeIndex(608, (d_1857_v4_).length(0))], (D22_DC58(d_1854_cf66_, (self).f28, (d_1860_v7_)[default__.safeIndex(len(d_1866_v13_), len(d_1860_v7_))])).cf87]), (d_1857_v4_)[default__.safeIndex(608, (d_1857_v4_).length(0))], globalState), d_1856_cf64_]))})
            d_1868_v15_: _dafny.Map
            d_1868_v15_ = _dafny.Map({d_1863_v10_: d_1867_v14_})
            def iife167_():
                coll79_ = _dafny.Map()
                compr_79_: str
                for compr_79_ in (_dafny.MultiSet([_dafny.CodePoint('s'), _dafny.CodePoint('g'), d_1858_v5_])).Elements:
                    d_1869_v16_: str = compr_79_
                    if (d_1869_v16_) in (_dafny.MultiSet([_dafny.CodePoint('s'), _dafny.CodePoint('g'), d_1858_v5_])):
                        coll79_[d_1869_v16_] = d_1860_v7_
                return _dafny.Map(coll79_)
            d_1858_v5_ = (default__.fm54(((d_1840_v0_)[d_1856_cf64_] if (d_1856_cf64_) in (d_1840_v0_) else False), ((d_1868_v15_)[d_1863_v10_] if (d_1863_v10_) in (d_1868_v15_) else iife167_()
            ), (self).f29, globalState) if (d_1857_v4_)[default__.safeIndex(608, (d_1857_v4_).length(0))] else _dafny.CodePoint('j'))
        elif source28_.is_DC45:
            d_1870___mcc_h3_ = source28_.cf63
            d_1871_cf63_ = d_1870___mcc_h3_
            (globalState).f6 = (self).f28
            d_1872_v17_: C0
            nw296_ = C0()
            nw296_.ctor__()
            d_1872_v17_ = nw296_
            (globalState).f5 = (self).f39
            d_1873_v18_: _dafny.Array
            nw297_ = _dafny.Array(int(0), 20)
            d_1873_v18_ = nw297_
            index293_ = default__.safeIndex(290, (d_1873_v18_).length(0))
            (d_1873_v18_)[index293_] = (self).f29
            index294_ = default__.safeIndex(290, (d_1873_v18_).length(0))
            (d_1873_v18_)[index294_] = (self).f29
        elif True:
            d_1874___mcc_h4_ = source28_.cf67
            d_1875_cf67_ = d_1874___mcc_h4_
            d_1876_v19_: str
            d_1876_v19_ = _dafny.CodePoint('g')
            d_1877_v20_: _dafny.MultiSet
            d_1877_v20_ = _dafny.MultiSet([(0) - (len(default__.fm55((self).f29, d_1876_v19_, globalState)))])
            d_1877_v20_ = d_1877_v20_
            d_1878_v21_: _dafny.Seq
            d_1878_v21_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "icrqilo"))
            d_1879_v22_: D5
            d_1879_v22_ = D5_DC12((self).f27, (self).f28)
            d_1880_v23_: _dafny.Map
            d_1880_v23_ = _dafny.Map({d_1878_v21_: d_1879_v22_})
            d_1880_v23_ = (d_1880_v23_).set(d_1878_v21_, d_1879_v22_)
            d_1881_v24_: _dafny.Seq
            d_1881_v24_ = _dafny.SeqWithoutIsStrInference([(self).f28, (self).f29])
            d_1882_v25_: _dafny.Seq
            d_1882_v25_ = _dafny.SeqWithoutIsStrInference([len(d_1881_v24_)])
            d_1883_v26_: D21
            d_1883_v26_ = D21_DC54((self).f27, (self).f29, (d_1881_v24_)[default__.safeIndex((self).f28, len(d_1881_v24_))])
            (globalState).f6 = (0) - ((d_1881_v24_)[default__.safeIndex(len(_dafny.Map({(d_1883_v26_).cf80: (self).f27})), len(d_1881_v24_))])
            d_1884_v27_: _dafny.Seq
            d_1884_v27_ = _dafny.SeqWithoutIsStrInference([d_1878_v21_, d_1878_v21_])
            d_1885_v28_: _dafny.Seq
            d_1885_v28_ = _dafny.SeqWithoutIsStrInference([(d_1884_v27_) + (d_1884_v27_)])
            d_1884_v27_ = (d_1885_v28_)[default__.safeIndex((_dafny.MultiSet([d_1878_v21_, d_1878_v21_])).cardinality, len(d_1885_v28_))]
        d_1886_v29_: D6
        d_1886_v29_ = D6_DC14((self).f39, len(((_dafny.SeqWithoutIsStrInference([(self).f27, (self).f39])).set(default__.safeIndex(23, len(_dafny.SeqWithoutIsStrInference([(self).f27, (self).f39]))), (self).f39)).set(default__.safeIndex((self).f29, len((_dafny.SeqWithoutIsStrInference([(self).f27, (self).f39])).set(default__.safeIndex(23, len(_dafny.SeqWithoutIsStrInference([(self).f27, (self).f39]))), (self).f39))), (self).f27)))
        d_1887_v30_: str
        d_1887_v30_ = _dafny.CodePoint('q')
        d_1888_v31_: D23
        d_1888_v31_ = D23_DC60((self).f27, d_1887_v30_)
        (globalState).f2 = (default__.fm56((self).f27, default__.fm57(d_1886_v29_, (d_1888_v31_).cf91, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "befgiodrr")), (self).f28, globalState), ((self).f29 if not((self).f27) else (self).f28), (d_1842_v2_) + (d_1842_v2_), globalState)).cardinality
        d_1889_i1_: int
        d_1889_i1_ = 0
        with _dafny.label("10"):
            while ((self).f27) or ((self).f27):
                with _dafny.c_label("10"):
                    if (d_1889_i1_) >= (100):
                        raise _dafny.Break("10")
                    d_1889_i1_ = (d_1889_i1_) + (1)
                    d_1890_v32_: _dafny.Map
                    d_1890_v32_ = _dafny.Map({False: (self).f28})
                    d_1890_v32_ = (d_1890_v32_).set((self).f27, 505)
                    (globalState).f6 = default__.fm1(((self).f39 if (self).f39 else (self).f39), d_1842_v2_, (self).f27, globalState)
                    d_1891_v33_: _dafny.Map
                    d_1891_v33_ = _dafny.Map({(self).f27: d_1887_v30_})
                    d_1892_v34_: D9
                    d_1892_v34_ = D9_DC20(d_1891_v33_)
                    d_1893_v35_: _dafny.Seq
                    d_1893_v35_ = _dafny.SeqWithoutIsStrInference([d_1892_v34_, d_1892_v34_])
                    d_1894_v36_: D9
                    d_1894_v36_ = D9_DC22((d_1893_v35_)[default__.safeIndex(708, len(d_1893_v35_))])
                    d_1895_v37_: D9
                    d_1895_v37_ = D9_DC22(d_1894_v36_)
                    d_1895_v37_ = d_1895_v37_
                    d_1896_v38_: _dafny.MultiSet
                    d_1896_v38_ = _dafny.MultiSet([(self).f39, not((self).f39)])
                    d_1897_v39_: _dafny.Map
                    d_1897_v39_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([462]): (d_1896_v38_).cardinality})
                    d_1898_v40_: D3
                    d_1898_v40_ = D3_DC8(d_1897_v39_)
                    pat_let_tv41_ = d_1897_v39_
                    def iife168_(_pat_let44_0):
                        def iife169_(d_1899_dt__update__tmp_h0_):
                            def iife170_(_pat_let45_0):
                                def iife171_(d_1900_dt__update_hcf13_h0_):
                                    return D3_DC8(d_1900_dt__update_hcf13_h0_)
                                return iife171_(_pat_let45_0)
                            return iife170_(pat_let_tv41_)
                        return iife169_(_pat_let44_0)
                    source30_ = iife168_(d_1898_v40_)
                    if source30_.is_DC9:
                        d_1901___mcc_h8_ = source30_.cf14
                        d_1902___mcc_h9_ = source30_.cf15
                        d_1903___mcc_h10_ = source30_.cf16
                        d_1904_cf16_ = d_1903___mcc_h10_
                        d_1905_cf15_ = d_1902___mcc_h9_
                        d_1906_cf14_ = d_1901___mcc_h8_
                        (globalState).f5 = (self).f27
                        d_1907_v41_: _dafny.Set
                        d_1907_v41_ = _dafny.Set({(self).f27, (self).f39, d_1906_cf14_, d_1906_cf14_})
                        d_1907_v41_ = (d_1907_v41_) - (d_1907_v41_)
                        d_1908_v42_: _dafny.MultiSet
                        d_1908_v42_ = _dafny.MultiSet([d_1905_cf15_])
                        d_1909_v43_: _dafny.Array
                        nw298_ = _dafny.Array(None, 2)
                        nw298_[int(0)] = (self).f39
                        nw298_[int(1)] = (d_1908_v42_) == (d_1908_v42_)
                        d_1909_v43_ = nw298_
                        d_1909_v43_ = d_1909_v43_
                        d_1910_v44_: _dafny.Seq
                        d_1910_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yqklt"))
                        (globalState).f21 = ((d_1840_v0_)[(self).f29] if ((self).f29) in (d_1840_v0_) else (d_1910_v44_) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "drpv"))))
                    elif True:
                        d_1911___mcc_h11_ = source30_.cf13
                        d_1912_cf13_ = d_1911___mcc_h11_
                        (globalState).f2 = (self).f29
                        d_1913_v45_: C1
                        nw299_ = C1()
                        nw299_.ctor__((self).f27)
                        d_1913_v45_ = nw299_
                        d_1914_v46_: _dafny.Seq
                        d_1914_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dvmmqwece"))
                        d_1915_v47_: C3
                        nw300_ = C3()
                        nw300_.ctor__(True, (d_1914_v46_) == (d_1914_v46_))
                        d_1915_v47_ = nw300_
                        d_1916_v48_: _dafny.Seq
                        d_1916_v48_ = _dafny.SeqWithoutIsStrInference([d_1914_v46_, (d_1914_v46_).set(default__.safeIndex((self).f28, len(d_1914_v46_)), d_1887_v30_), _dafny.SeqWithoutIsStrInference([d_1887_v30_, d_1887_v30_])])
                        d_1914_v46_ = (d_1916_v48_)[default__.safeIndex(((self).f28) * ((self).f28), len(d_1916_v48_))]
                    pass
            pass
        d_1917_v49_: _dafny.MultiSet
        d_1917_v49_ = _dafny.MultiSet([(self).f29, (self).f29, (self).f29, (self).f29, (self).f29])
        d_1918_i2_: int
        d_1918_i2_ = 0
        with _dafny.label("11"):
            while not(((self).f29) > (((d_1917_v49_)[(self).f29] if ((self).f29) in (d_1917_v49_) else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pj")))))):
                with _dafny.c_label("11"):
                    if (d_1918_i2_) >= (100):
                        raise _dafny.Break("11")
                    d_1918_i2_ = (d_1918_i2_) + (1)
                    d_1919_v50_: _dafny.Map
                    d_1919_v50_ = _dafny.Map({(self).f27: (self).f28})
                    d_1919_v50_ = (d_1919_v50_).set((self).f39, -267)
                    d_1920_v51_: _dafny.Array
                    nw301_ = _dafny.Array(int(0), 26)
                    d_1920_v51_ = nw301_
                    d_1921_v52_: _dafny.Array
                    nw302_ = _dafny.Array(None, 2)
                    nw302_[int(0)] = d_1920_v51_
                    nw302_[int(1)] = d_1920_v51_
                    d_1921_v52_ = nw302_
                    d_1922_v53_: _dafny.Map
                    d_1922_v53_ = _dafny.Map({(self).f39: False})
                    d_1923_v54_: D23
                    d_1923_v54_ = D23_DC59(d_1922_v53_)
                    d_1924_v55_: _dafny.Map
                    d_1924_v55_ = _dafny.Map({d_1921_v52_: d_1923_v54_})
                    pat_let_tv42_ = d_1922_v53_
                    def iife172_(_pat_let46_0):
                        def iife173_(d_1925_dt__update__tmp_h1_):
                            def iife174_(_pat_let47_0):
                                def iife175_(d_1926_dt__update_hcf90_h0_):
                                    return D23_DC59(d_1926_dt__update_hcf90_h0_)
                                return iife175_(_pat_let47_0)
                            return iife174_(pat_let_tv42_)
                        return iife173_(_pat_let46_0)
                    d_1924_v55_ = (d_1924_v55_).set(d_1921_v52_, iife172_(d_1923_v54_))
                    d_1927_v56_: C0
                    nw303_ = C0()
                    nw303_.ctor__()
                    d_1927_v56_ = nw303_
                    if (self).f27:
                        (globalState).f2 = (self).f29
                        (globalState).f12 = not((self).f27)
                        (globalState).f2 = (((0) - ((self).f29)) * (len(d_1840_v0_)) if (self).f39 else (self).f29)
                        (globalState).f6 = (self).f29
                        d_1928_v57_: D16
                        d_1928_v57_ = D16_DC35(d_1917_v49_)
                        d_1928_v57_ = (d_1928_v57_ if not((self).f39) else (d_1928_v57_ if (self).f39 else d_1928_v57_))
                    elif True:
                        d_1929_v58_: _dafny.MultiSet
                        d_1929_v58_ = _dafny.MultiSet([not((d_1842_v2_)[default__.safeIndex(476, len(d_1842_v2_))]), (self).f39])
                        (globalState).f6 = ((d_1929_v58_) - ((d_1929_v58_).set((self).f27, default__.abs((self).f28)))).cardinality
                        index295_ = default__.safeIndex(801, (d_1920_v51_).length(0))
                        (d_1920_v51_)[index295_] = (0) - (((0) - ((self).f28)) - ((self).f28))
                        d_1930_v59_: _dafny.Seq
                        d_1930_v59_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hmtbj"))
                        d_1931_v60_: _dafny.Array
                        nw304_ = _dafny.Array(None, 19)
                        nw304_[int(0)] = d_1930_v59_
                        nw304_[int(1)] = ((d_1930_v59_).set(default__.safeIndex((self).f29, len(d_1930_v59_)), d_1887_v30_)) + (_dafny.SeqWithoutIsStrInference([d_1887_v30_ for d_1932_i3_ in range(default__.abs(721))]))
                        nw304_[int(2)] = d_1930_v59_
                        nw304_[int(3)] = d_1930_v59_
                        nw304_[int(4)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uql"))) + (d_1930_v59_)
                        nw304_[int(5)] = ((D18_DC42(d_1930_v59_, (self).f39)).cf57) + (_dafny.SeqWithoutIsStrInference([d_1887_v30_ for d_1933_i4_ in range(default__.abs(794))]))
                        nw304_[int(6)] = d_1930_v59_
                        nw304_[int(7)] = d_1930_v59_
                        nw304_[int(8)] = _dafny.SeqWithoutIsStrInference([d_1887_v30_ for d_1934_i5_ in range(default__.abs(425))])
                        nw304_[int(9)] = d_1930_v59_
                        nw304_[int(10)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dqupe")) if (self).f39 else d_1930_v59_)
                        nw304_[int(11)] = d_1930_v59_
                        nw304_[int(12)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))).set(default__.safeIndex((self).f29, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q")))), d_1887_v30_)
                        nw304_[int(13)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdulrfy"))
                        nw304_[int(14)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "laukebs"))
                        nw304_[int(15)] = d_1930_v59_
                        nw304_[int(16)] = d_1930_v59_
                        nw304_[int(17)] = _dafny.SeqWithoutIsStrInference([d_1887_v30_ for d_1935_i6_ in range(default__.abs(467))])
                        nw304_[int(18)] = (d_1930_v59_) + (d_1930_v59_)
                        d_1931_v60_ = nw304_
                        index296_ = default__.safeIndex(601, (d_1931_v60_).length(0))
                        (d_1931_v60_)[index296_] = d_1930_v59_
                        d_1936_v61_: _dafny.Seq
                        d_1936_v61_ = _dafny.SeqWithoutIsStrInference([(self).f29])
                        d_1937_v62_: _dafny.Map
                        d_1937_v62_ = _dafny.Map({(d_1930_v59_)[default__.safeIndex(len(d_1936_v61_), len(d_1930_v59_))]: (self).f28})
                        index297_ = default__.safeIndex(801, (d_1920_v51_).length(0))
                        index298_ = default__.safeIndex(601, (d_1931_v60_).length(0))
                        rhs276_ = (0) - (default__.safeModuloInt((0) - ((self).f28), default__.safeModuloInt((self).f29, (self).f28)))
                        rhs277_ = ((self).f39) or ((d_1929_v58_).issubset(d_1929_v58_))
                        rhs278_ = ((d_1930_v59_) + ((d_1930_v59_) + (d_1930_v59_))).set(default__.safeIndex((len(d_1937_v62_)) - ((self).f29), len((d_1930_v59_) + ((d_1930_v59_) + (d_1930_v59_)))), _dafny.CodePoint('h'))
                        rhs279_ = ((d_1929_v58_).intersection(_dafny.MultiSet([((d_1922_v53_)[(self).f39] if ((self).f39) in (d_1922_v53_) else (self).f27)]))).ispropersubset(d_1929_v58_)
                        lhs221_ = d_1920_v51_
                        lhs222_ = default__.safeIndex(801, (d_1920_v51_).length(0))
                        lhs223_ = globalState
                        lhs224_ = d_1931_v60_
                        lhs225_ = default__.safeIndex(601, (d_1931_v60_).length(0))
                        lhs226_ = globalState
                        lhs221_[lhs222_] = rhs276_
                        lhs223_.f12 = rhs277_
                        lhs224_[lhs225_] = rhs278_
                        lhs226_.f14 = rhs279_
                        d_1938_v63_: C0
                        nw305_ = C0()
                        nw305_.ctor__()
                        d_1938_v63_ = nw305_
                        d_1939_v64_: _dafny.Set
                        d_1939_v64_ = _dafny.Set({(self).f27})
                        d_1940_v65_: D11
                        d_1940_v65_ = D11_DC24(d_1939_v64_)
                        d_1941_v66_: _dafny.Map
                        d_1941_v66_ = _dafny.Map({(self).f39: _dafny.Map({(self).f28: (self).f27})})
                        d_1942_v67_: _dafny.Array
                        nw306_ = _dafny.Array(None, 8)
                        nw306_[int(0)] = (self).f29
                        nw306_[int(1)] = 620
                        nw306_[int(2)] = (d_1920_v51_)[default__.safeIndex(801, (d_1920_v51_).length(0))]
                        nw306_[int(3)] = (self).f28
                        nw306_[int(4)] = len(default__.fm58((self).f39, (self).f39, (self).f39, (self).f27, globalState))
                        nw306_[int(5)] = (d_1920_v51_)[default__.safeIndex(801, (d_1920_v51_).length(0))]
                        nw306_[int(6)] = (d_1920_v51_)[default__.safeIndex(801, (d_1920_v51_).length(0))]
                        nw306_[int(7)] = len(d_1941_v66_)
                        d_1942_v67_ = nw306_
                        d_1943_v68_: D29
                        d_1943_v68_ = D29_DC70(d_1940_v65_, (self).f28, d_1942_v67_, not(not((self).f39)))
                        d_1944_v69_: D29
                        d_1944_v69_ = D29_DC72(d_1943_v68_)
                        d_1945_v70_: D9
                        d_1945_v70_ = D9_DC21((self).f27, (self).f28)
                        rhs280_ = d_1842_v2_
                        rhs281_ = (self).f27
                        rhs282_ = (d_1944_v69_ if (self).f39 else d_1944_v69_)
                        rhs283_ = d_1945_v70_
                        rhs284_ = d_1887_v30_
                        lhs227_ = globalState
                        d_1842_v2_ = rhs280_
                        lhs227_.f21 = rhs281_
                        d_1944_v69_ = rhs282_
                        d_1945_v70_ = rhs283_
                        d_1887_v30_ = rhs284_
                        d_1946_v71_: _dafny.Set
                        d_1946_v71_ = _dafny.Set({(self).f29})
                        (globalState).f21 = (default__.fm1(default__.fm0((d_1931_v60_)[default__.safeIndex(601, (d_1931_v60_).length(0))], (self).f27, (self).f27, _dafny.Set({len(d_1936_v61_), (self).f29, len(d_1946_v71_)}), globalState), d_1842_v2_, (self).f39, globalState)) <= ((d_1920_v51_)[default__.safeIndex(801, (d_1920_v51_).length(0))])
                    pass
            pass
        d_1947_v72_: _dafny.Seq
        d_1947_v72_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iv"))
        d_1948_v73_: _dafny.MultiSet
        d_1948_v73_ = _dafny.MultiSet([(self).f39])
        r0 = ((len(d_1947_v72_)) + ((self).f28)) - (default__.fm1((self).f39, default__.fm52(len(d_1947_v72_), (self).f39, d_1948_v73_, globalState), (self).f27, globalState))
        r0 = 545
        return r0

    def m2(self, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: bool = False
        d_1949_v0_: C2
        nw307_ = C2()
        nw307_.ctor__((self).f39, (self).f39)
        d_1949_v0_ = nw307_
        d_1950_v1_: D21
        d_1950_v1_ = D21_DC53((not((self).f27)) == ((d_1949_v0_).f34), (self).f28, False)
        source31_ = d_1950_v1_
        if source31_.is_DC53:
            d_1951___mcc_h0_ = source31_.cf76
            d_1952___mcc_h1_ = source31_.cf77
            d_1953___mcc_h2_ = source31_.cf78
            d_1954_cf78_ = d_1953___mcc_h2_
            d_1955_cf77_ = d_1952___mcc_h1_
            d_1956_cf76_ = d_1951___mcc_h0_
            d_1957_v2_: _dafny.Array
            def lambda96_(d_1958_i0_):
                return default__.safeDivisionInt(d_1958_i0_, (self).f29)

            init49_ = lambda96_
            nw308_ = _dafny.Array(None, 16)
            for i0_49_ in range(nw308_.length(0)):
                nw308_[i0_49_] = init49_(i0_49_)
            d_1957_v2_ = nw308_
            index299_ = default__.safeIndex(896, (d_1957_v2_).length(0))
            (d_1957_v2_)[index299_] = d_1955_cf77_
            d_1959_v3_: _dafny.Seq
            d_1959_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "msuoeiyfs"))
            d_1960_v4_: _dafny.Seq
            d_1960_v4_ = _dafny.SeqWithoutIsStrInference([(d_1949_v0_).f34])
            d_1961_v5_: _dafny.Map
            d_1961_v5_ = _dafny.Map({d_1955_cf77_: default__.fm1(not((d_1949_v0_).f34), d_1960_v4_, (d_1949_v0_).f34, globalState)})
            d_1962_v6_: str
            d_1962_v6_ = _dafny.CodePoint('b')
            d_1963_v7_: _dafny.Set
            d_1963_v7_ = _dafny.Set({d_1962_v6_, d_1962_v6_, d_1962_v6_, d_1962_v6_, (d_1959_v3_)[default__.safeIndex((self).f28, len(d_1959_v3_))]})
            d_1964_v8_: _dafny.Map
            d_1964_v8_ = _dafny.Map({d_1956_cf76_: d_1954_cf78_})
            index300_ = default__.safeIndex(896, (d_1957_v2_).length(0))
            rhs285_ = ((d_1959_v3_).set(default__.safeIndex(((d_1961_v5_)[(self).f29] if ((self).f29) in (d_1961_v5_) else len(d_1963_v7_)), len(d_1959_v3_)), d_1962_v6_)) + ((d_1959_v3_) + (_dafny.SeqWithoutIsStrInference([d_1962_v6_ for d_1965_i1_ in range(default__.abs(450))])))
            rhs286_ = len(d_1964_v8_)
            lhs228_ = d_1957_v2_
            lhs229_ = default__.safeIndex(896, (d_1957_v2_).length(0))
            r1 = rhs285_
            lhs228_[lhs229_] = rhs286_
            d_1966_v9_: C0
            nw309_ = C0()
            nw309_.ctor__()
            d_1966_v9_ = nw309_
            d_1967_v10_: D3
            d_1967_v10_ = D3_DC9((len(d_1959_v3_)) <= ((self).f29), d_1955_cf77_, -547)
            d_1967_v10_ = D3_DC9((self).f27, (d_1957_v2_)[default__.safeIndex(896, (d_1957_v2_).length(0))], 81)
            d_1968_v11_: D0
            d_1968_v11_ = D0_DC3(D0_DC3((d_1949_v0_).fm5(globalState)))
            d_1969_v12_: _dafny.Seq
            d_1969_v12_ = _dafny.SeqWithoutIsStrInference([d_1968_v11_])
            d_1970_v13_: D2
            d_1970_v13_ = D2_DC5(d_1969_v12_)
            d_1971_v14_: D2
            d_1971_v14_ = D2_DC7(d_1970_v13_)
            d_1972_v15_: D2
            d_1972_v15_ = D2_DC7(d_1971_v14_)
            d_1973_v16_: D2
            d_1973_v16_ = D2_DC7((d_1972_v15_).cf12)
            d_1973_v16_ = d_1972_v15_
        elif source31_.is_DC54:
            d_1974___mcc_h3_ = source31_.cf79
            d_1975___mcc_h4_ = source31_.cf80
            d_1976___mcc_h5_ = source31_.cf81
            d_1977_cf81_ = d_1976___mcc_h5_
            d_1978_cf80_ = d_1975___mcc_h4_
            d_1979_cf79_ = d_1974___mcc_h3_
            d_1980_v17_: str
            d_1980_v17_ = _dafny.CodePoint('b')
            d_1981_v18_: D23
            d_1981_v18_ = D23_DC60((self).f27, d_1980_v17_)
            d_1982_v19_: _dafny.Map
            d_1982_v19_ = _dafny.Map({d_1981_v18_: (default__.fm59(d_1979_cf79_, d_1979_cf79_, globalState)).cf60})
            d_1983_v20_: _dafny.Seq
            d_1983_v20_ = _dafny.SeqWithoutIsStrInference([False])
            d_1984_v21_: D0
            d_1984_v21_ = D0_DC1((self).f39)
            d_1985_v22_: _dafny.MultiSet
            d_1985_v22_ = _dafny.MultiSet([(d_1949_v0_).fm4(d_1984_v21_, globalState), (self).f27])
            d_1986_v23_: C2
            nw310_ = C2()
            nw310_.ctor__(((self).f27 if ((d_1982_v19_)[d_1981_v18_] if (d_1981_v18_) in (d_1982_v19_) else (d_1949_v0_).f34) else (d_1983_v20_)[default__.safeIndex((d_1985_v22_).cardinality, len(d_1983_v20_))]), ((self).f28) != (d_1978_cf80_))
            d_1986_v23_ = nw310_
            d_1987_v24_: _dafny.MultiSet
            d_1987_v24_ = _dafny.MultiSet([d_1980_v17_])
            (globalState).f6 = (((d_1987_v24_) | (d_1987_v24_)).cardinality) - (d_1978_cf80_)
            d_1988_v25_: _dafny.Map
            d_1988_v25_ = _dafny.Map({(d_1986_v23_).f34: (self).f27})
            d_1989_v26_: _dafny.Set
            d_1989_v26_ = _dafny.Set({(self).f27})
            d_1990_v27_: _dafny.Map
            d_1990_v27_ = _dafny.Map({True: len(d_1989_v26_)})
            d_1991_v28_: _dafny.Map
            d_1991_v28_ = _dafny.Map({d_1980_v17_: ((d_1990_v27_)[(d_1949_v0_).f34] if ((d_1949_v0_).f34) in (d_1990_v27_) else len(d_1983_v20_))})
            d_1992_v29_: D29
            d_1992_v29_ = D29_DC69(d_1991_v28_)
            d_1993_v30_: _dafny.Map
            d_1993_v30_ = _dafny.Map({False: _dafny.Map({(d_1986_v23_).f34: d_1978_cf80_})})
            pat_let_tv43_ = d_1991_v28_
            d_1994_v31_: _dafny.MultiSet
            def iife176_(_pat_let48_0):
                def iife177_(d_1995_dt__update__tmp_h0_):
                    def iife178_(_pat_let49_0):
                        def iife179_(d_1996_dt__update_hcf109_h0_):
                            return D29_DC69(d_1996_dt__update_hcf109_h0_)
                        return iife179_(_pat_let49_0)
                    return iife178_(pat_let_tv43_)
                return iife177_(_pat_let48_0)
            d_1994_v31_ = _dafny.MultiSet([d_1988_v25_, default__.fm60(iife176_(d_1992_v29_), ((d_1993_v30_)[(d_1949_v0_).f34] if ((d_1949_v0_).f34) in (d_1993_v30_) else d_1990_v27_), globalState)])
            (globalState).f2 = ((d_1994_v31_)[d_1988_v25_] if (d_1988_v25_) in (d_1994_v31_) else (d_1986_v23_).fm9(len(d_1989_v26_), globalState))
            d_1997_v32_: _dafny.Array
            nw311_ = _dafny.Array(_dafny.Array(None, 0), 27)
            d_1997_v32_ = nw311_
            rhs287_ = default__.safeModuloInt((self).f29, (self).f28)
            rhs288_ = (self).f29
            rhs289_ = d_1997_v32_
            rhs290_ = (len(_dafny.SeqWithoutIsStrInference([d_1985_v22_, d_1985_v22_]))) * (d_1977_cf81_)
            lhs230_ = globalState
            lhs231_ = globalState
            lhs232_ = globalState
            lhs230_.f2 = rhs287_
            lhs231_.f6 = rhs288_
            r2 = rhs289_
            lhs232_.f6 = rhs290_
        elif source31_.is_DC55:
            d_1998___mcc_h6_ = source31_.cf82
            d_1999_cf82_ = d_1998___mcc_h6_
            d_2000_v33_: _dafny.Array
            nw312_ = _dafny.Array(False, 26)
            d_2000_v33_ = nw312_
            index301_ = default__.safeIndex(782, (d_2000_v33_).length(0))
            (d_2000_v33_)[index301_] = (d_1949_v0_).f34
            d_2001_v34_: _dafny.Map
            d_2001_v34_ = _dafny.Map({(self).f39: (self).f27})
            index302_ = default__.safeIndex(782, (d_2000_v33_).length(0))
            (d_2000_v33_)[index302_] = ((d_2001_v34_)[(d_1949_v0_).f34] if ((d_1949_v0_).f34) in (d_2001_v34_) else (self).f27)
            d_2002_v35_: _dafny.Array
            def lambda97_(d_2003_i2_):
                return default__.safeModuloInt(d_2003_i2_, (self).f29)

            init50_ = lambda97_
            nw313_ = _dafny.Array(None, 9)
            for i0_50_ in range(nw313_.length(0)):
                nw313_[i0_50_] = init50_(i0_50_)
            d_2002_v35_ = nw313_
            d_2002_v35_ = d_2002_v35_
            d_2004_v36_: str
            d_2004_v36_ = _dafny.CodePoint('m')
            d_2005_v37_: D9
            d_2005_v37_ = D9_DC20(_dafny.Map({not((d_1949_v0_).f34): d_2004_v36_}))
            d_2006_v38_: D9
            d_2006_v38_ = D9_DC22(d_2005_v37_)
            d_2007_v39_: D9
            d_2007_v39_ = D9_DC22(d_2005_v37_)
            d_2008_v40_: D9
            d_2008_v40_ = D9_DC22(d_2005_v37_)
            pat_let_tv44_ = d_2000_v33_
            pat_let_tv45_ = d_2000_v33_
            def iife180_(_pat_let50_0):
                def iife181_(d_2009_dt__update__tmp_h1_):
                    def iife182_(_pat_let51_0):
                        def iife183_(d_2010_dt__update_hcf33_h0_):
                            return D9_DC22(d_2010_dt__update_hcf33_h0_)
                        return iife183_(_pat_let51_0)
                    return iife182_(D9_DC21((pat_let_tv45_)[default__.safeIndex(782, (pat_let_tv44_).length(0))], (self).f28))
                return iife181_(_pat_let50_0)
            source32_ = iife180_(d_2008_v40_)
            if source32_.is_DC21:
                d_2011___mcc_h8_ = source32_.cf31
                d_2012___mcc_h9_ = source32_.cf32
                d_2013_cf32_ = d_2012___mcc_h9_
                d_2014_cf31_ = d_2011___mcc_h8_
                d_2015_v41_: _dafny.MultiSet
                d_2015_v41_ = _dafny.MultiSet([d_2000_v33_, d_2000_v33_])
                d_2015_v41_ = ((d_2015_v41_).intersection(d_2015_v41_)) - ((_dafny.MultiSet([d_2000_v33_, d_2000_v33_, d_2000_v33_, d_2000_v33_, d_2000_v33_])) - (d_2015_v41_))
                d_2016_v42_: _dafny.Seq
                d_2016_v42_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fuhnfbnru"))
                d_2017_v43_: _dafny.MultiSet
                d_2017_v43_ = _dafny.MultiSet([d_2016_v42_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xaubjnjj")), ((self).fm3(globalState)).set(default__.safeIndex((self).f28, len((self).fm3(globalState))), d_2004_v36_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))])
                d_2017_v43_ = (d_2017_v43_) | (d_2017_v43_)
                d_1999_cf82_ = ((0) - ((self).f29)) * (d_1999_cf82_)
                d_2018_v44_: _dafny.MultiSet
                d_2018_v44_ = _dafny.MultiSet([(d_1949_v0_).f34])
                d_2019_v45_: _dafny.Seq
                d_2019_v45_ = _dafny.SeqWithoutIsStrInference([False, not(True)])
                d_2020_v46_: D9
                d_2020_v46_ = D9_DC21((d_1949_v0_).f34, default__.safeDivisionInt(((d_2018_v44_)[(d_1949_v0_).f34] if ((d_1949_v0_).f34) in (d_2018_v44_) else (self).f28), len(d_2019_v45_)))
                d_2020_v46_ = d_2020_v46_
            elif source32_.is_DC20:
                d_2021___mcc_h10_ = source32_.cf30
                d_2022_cf30_ = d_2021___mcc_h10_
                d_2023_v47_: _dafny.Map
                d_2023_v47_ = _dafny.Map({(self).f29: (d_2000_v33_)[default__.safeIndex(782, (d_2000_v33_).length(0))]})
                d_2023_v47_ = d_2023_v47_
                d_2004_v36_ = d_2004_v36_
                d_2024_v48_: _dafny.MultiSet
                d_2024_v48_ = _dafny.MultiSet([d_1999_cf82_])
                d_2025_v49_: _dafny.Map
                d_2025_v49_ = _dafny.Map({(d_2024_v48_).ispropersubset(d_2024_v48_): d_1999_cf82_})
                (globalState).f13 = d_2025_v49_
                d_2026_v50_: _dafny.Seq
                d_2026_v50_ = _dafny.SeqWithoutIsStrInference([(d_1949_v0_).f34, (d_1949_v0_).f34])
                d_2027_v51_: _dafny.Set
                d_2027_v51_ = _dafny.Set({d_2026_v50_})
                d_2028_v52_: C2
                nw314_ = C2()
                nw314_.ctor__(not((d_2027_v51_).ispropersubset(d_2027_v51_)), (self).f39)
                d_2028_v52_ = nw314_
            elif True:
                d_2029___mcc_h11_ = source32_.cf33
                d_2030_cf33_ = d_2029___mcc_h11_
                d_2031_v53_: _dafny.MultiSet
                d_2031_v53_ = _dafny.MultiSet([default__.fm1((self).f39, _dafny.SeqWithoutIsStrInference([True, (d_1949_v0_).f34, (d_1949_v0_).f34, (d_1949_v0_).f34, (d_1949_v0_).f34]), (self).f39, globalState), 414])
                d_2032_v54_: _dafny.Set
                d_2032_v54_ = _dafny.Set({d_2031_v53_, d_2031_v53_, d_2031_v53_})
                d_2033_v55_: _dafny.Seq
                d_2033_v55_ = _dafny.SeqWithoutIsStrInference([(d_2000_v33_)[default__.safeIndex(782, (d_2000_v33_).length(0))]])
                rhs291_ = d_2032_v54_
                rhs292_ = (d_1949_v0_).f34
                rhs293_ = d_1999_cf82_
                rhs294_ = d_2002_v35_
                rhs295_ = d_2033_v55_
                lhs233_ = globalState
                lhs234_ = globalState
                d_2032_v54_ = rhs291_
                lhs233_.f14 = rhs292_
                lhs234_.f2 = rhs293_
                d_2002_v35_ = rhs294_
                d_2033_v55_ = rhs295_
                (globalState).f25 = not((self).f27)
                d_2001_v34_ = (d_2001_v34_).set((self).f27, (d_1949_v0_).f34)
                d_2034_v56_: _dafny.Seq
                d_2034_v56_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pb"))
                d_2035_v57_: _dafny.Map
                d_2035_v57_ = _dafny.Map({(self).f39: (d_2034_v56_) + (_dafny.SeqWithoutIsStrInference([d_2004_v36_ for d_2036_i3_ in range(default__.abs(-751))]))})
                d_2037_v58_: _dafny.Seq
                d_2037_v58_ = _dafny.SeqWithoutIsStrInference([default__.fm61(globalState)])
                d_2038_v59_: D0
                d_2038_v59_ = D0_DC1((d_1949_v0_).f34)
                d_2039_v60_: D6
                def iife184_(_pat_let52_0):
                    def iife185_(d_2040_dt__update__tmp_h2_):
                        def iife186_(_pat_let53_0):
                            def iife187_(d_2041_dt__update_hcf1_h0_):
                                return D0_DC1(d_2041_dt__update_hcf1_h0_)
                            return iife187_(_pat_let53_0)
                        return iife186_((self).f39)
                    return iife185_(_pat_let52_0)
                d_2039_v60_ = D6_DC13(_dafny.SeqWithoutIsStrInference([(d_1949_v0_).fm4(iife184_(d_2038_v59_), globalState)]))
                d_2042_v61_: _dafny.MultiSet
                d_2042_v61_ = _dafny.MultiSet([D6_DC13(d_2033_v55_), d_2039_v60_])
                d_2035_v57_ = (d_2037_v58_)[default__.safeIndex(((d_2042_v61_)[d_2039_v60_] if (d_2039_v60_) in (d_2042_v61_) else (self).f28), len(d_2037_v58_))]
            d_2043_v62_: _dafny.Map
            d_2043_v62_ = _dafny.Map({(d_1949_v0_).f34: d_1999_cf82_})
            d_2044_v63_: _dafny.Map
            d_2044_v63_ = _dafny.Map({d_2043_v62_: (self).f28})
            index303_ = default__.safeIndex(628, (d_2002_v35_).length(0))
            (d_2002_v35_)[index303_] = ((d_2044_v63_)[_dafny.Map({(d_1949_v0_).f34: (self).f29})] if (_dafny.Map({(d_1949_v0_).f34: (self).f29})) in (d_2044_v63_) else 845)
            d_2045_v64_: D19
            d_2045_v64_ = D19_DC46(d_1999_cf82_, d_2002_v35_, (d_2000_v33_)[default__.safeIndex(782, (d_2000_v33_).length(0))])
            index304_ = default__.safeIndex(628, (d_2002_v35_).length(0))
            (d_2002_v35_)[index304_] = (d_2045_v64_).cf64
        elif True:
            d_2046___mcc_h7_ = source31_.cf75
            d_2047_cf75_ = d_2046___mcc_h7_
            d_2048_v65_: _dafny.Array
            nw315_ = _dafny.Array(_dafny.Array(None, 0), 17)
            d_2048_v65_ = nw315_
            d_2049_v66_: _dafny.Array
            def lambda98_(d_2050_i4_):
                return (self).f27

            init51_ = lambda98_
            nw316_ = _dafny.Array(None, 24)
            for i0_51_ in range(nw316_.length(0)):
                nw316_[i0_51_] = init51_(i0_51_)
            d_2049_v66_ = nw316_
            index305_ = default__.safeIndex(422, (d_2048_v65_).length(0))
            (d_2048_v65_)[index305_] = d_2049_v66_
            index306_ = default__.safeIndex(422, (d_2048_v65_).length(0))
            (d_2048_v65_)[index306_] = d_2049_v66_
            d_2051_v67_: C1
            nw317_ = C1()
            nw317_.ctor__((self).f39)
            d_2051_v67_ = nw317_
            d_2052_v68_: C0
            nw318_ = C0()
            nw318_.ctor__()
            d_2052_v68_ = nw318_
            (globalState).f2 = (self).f29
        d_2053_v69_: _dafny.Array
        nw319_ = _dafny.Array(False, 20)
        d_2053_v69_ = nw319_
        d_2054_v70_: _dafny.Seq
        d_2054_v70_ = _dafny.SeqWithoutIsStrInference([d_2053_v69_, d_2053_v69_, d_2053_v69_, d_2053_v69_])
        d_2055_v71_: str
        d_2055_v71_ = _dafny.CodePoint('y')
        d_2056_v72_: D23
        d_2056_v72_ = D23_DC60((self).f27, d_2055_v71_)
        d_2057_v73_: _dafny.Map
        d_2057_v73_ = _dafny.Map({(self).f27: (self).f28})
        d_2058_v74_: _dafny.Seq
        d_2058_v74_ = _dafny.SeqWithoutIsStrInference([len(d_2057_v73_), (_dafny.MultiSet([(self).f29, (self).f28, (self).f29])).cardinality])
        d_2059_v75_: _dafny.Map
        d_2059_v75_ = _dafny.Map({d_2055_v71_: d_2058_v74_})
        d_2060_v76_: _dafny.Array
        nw320_ = _dafny.Array(None, 28)
        nw320_[int(0)] = d_2055_v71_
        nw320_[int(1)] = _dafny.CodePoint('d')
        nw320_[int(2)] = d_2055_v71_
        nw320_[int(3)] = d_2055_v71_
        nw320_[int(4)] = d_2055_v71_
        nw320_[int(5)] = _dafny.CodePoint('q')
        nw320_[int(6)] = d_2055_v71_
        nw320_[int(7)] = d_2055_v71_
        nw320_[int(8)] = d_2055_v71_
        nw320_[int(9)] = d_2055_v71_
        nw320_[int(10)] = d_2055_v71_
        nw320_[int(11)] = d_2055_v71_
        nw320_[int(12)] = d_2055_v71_
        nw320_[int(13)] = d_2055_v71_
        nw320_[int(14)] = d_2055_v71_
        nw320_[int(15)] = _dafny.CodePoint('y')
        nw320_[int(16)] = d_2055_v71_
        nw320_[int(17)] = d_2055_v71_
        nw320_[int(18)] = d_2055_v71_
        nw320_[int(19)] = (d_2056_v72_).cf92
        nw320_[int(20)] = d_2055_v71_
        nw320_[int(21)] = _dafny.CodePoint('g')
        nw320_[int(22)] = d_2055_v71_
        nw320_[int(23)] = _dafny.CodePoint('s')
        nw320_[int(24)] = d_2055_v71_
        nw320_[int(25)] = d_2055_v71_
        nw320_[int(26)] = default__.fm54((d_1949_v0_).f34, d_2059_v75_, (self).f28, globalState)
        nw320_[int(27)] = d_2055_v71_
        d_2060_v76_ = nw320_
        index307_ = default__.safeIndex(71, (d_2060_v76_).length(0))
        (d_2060_v76_)[index307_] = d_2055_v71_
        d_2061_v77_: D0
        d_2061_v77_ = D0_DC1((d_1949_v0_).f34)
        d_2062_v78_: _dafny.Set
        d_2062_v78_ = _dafny.Set({(self).f39})
        index308_ = default__.safeIndex(71, (d_2060_v76_).length(0))
        rhs296_ = d_2054_v70_
        rhs297_ = (_dafny.Set({(d_1949_v0_).fm4(d_2061_v77_, globalState), (d_1949_v0_).f34})).ispropersubset(d_2062_v78_)
        rhs298_ = (self).f28
        rhs299_ = _dafny.CodePoint('t')
        rhs300_ = (self).f29
        lhs235_ = globalState
        lhs236_ = globalState
        lhs237_ = d_2060_v76_
        lhs238_ = default__.safeIndex(71, (d_2060_v76_).length(0))
        lhs239_ = globalState
        d_2054_v70_ = rhs296_
        lhs235_.f25 = rhs297_
        lhs236_.f6 = rhs298_
        lhs237_[lhs238_] = rhs299_
        lhs239_.f6 = rhs300_
        hi11_ = (self).f29
        for d_2063_i5_ in range((self).f29, hi11_):
            (globalState).f6 = (d_2063_i5_ if (self).f27 else 327)
            d_2064_v79_: _dafny.Seq
            d_2064_v79_ = _dafny.SeqWithoutIsStrInference([(self).f27])
            (globalState).f2 = default__.fm1((self).f39, d_2064_v79_, False, globalState)
            d_2065_v80_: _dafny.Seq
            d_2065_v80_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sysaugy"))
            (globalState).f6 = len((d_2065_v80_) + (d_2065_v80_))
            if True:
                (globalState).f25 = (self).f27
                d_2066_v81_: _dafny.Array
                nw321_ = _dafny.Array(D13.default()(), 17)
                d_2066_v81_ = nw321_
                d_2067_v82_: _dafny.Map
                d_2067_v82_ = _dafny.Map({876: d_2066_v81_})
                d_2068_v83_: _dafny.Map
                d_2068_v83_ = _dafny.Map({((d_2067_v82_)[(self).f29] if ((self).f29) in (d_2067_v82_) else d_2066_v81_): (self).f29})
                d_2068_v83_ = (d_2068_v83_).set(d_2066_v81_, d_2063_i5_)
                d_2069_v84_: _dafny.Map
                d_2069_v84_ = _dafny.Map({143: (d_1949_v0_).f34})
                d_2070_v85_: _dafny.Map
                d_2070_v85_ = _dafny.Map({430: len(d_2069_v84_)})
                d_2071_v86_: _dafny.Map
                d_2071_v86_ = _dafny.Map({d_2070_v85_: 217})
                d_2072_v87_: _dafny.Array
                nw322_ = _dafny.Array(int(0), 28)
                d_2072_v87_ = nw322_
                index309_ = default__.safeIndex(883, (d_2072_v87_).length(0))
                (d_2072_v87_)[index309_] = (_dafny.MultiSet([(self).f39, (d_1949_v0_).fm4(D0_DC1((self).f27), globalState), (d_1949_v0_).f34])).cardinality
                d_2073_v88_: _dafny.Array
                nw323_ = _dafny.Array(D22.default()(), 21)
                d_2073_v88_ = nw323_
                d_2074_v89_: D22
                d_2074_v89_ = D22_DC58((d_1949_v0_).f34, d_2063_i5_, (self).f28)
                index310_ = default__.safeIndex(899, (d_2073_v88_).length(0))
                (d_2073_v88_)[index310_] = d_2074_v89_
                d_2075_v90_: D17
                d_2075_v90_ = D17_DC39(d_2065_v80_)
                index311_ = default__.safeIndex(883, (d_2072_v87_).length(0))
                index312_ = default__.safeIndex(899, (d_2073_v88_).length(0))
                rhs301_ = (d_2071_v86_) | (_dafny.Map({d_2070_v85_: len((d_2075_v90_).cf54)}))
                rhs302_ = (d_2063_i5_ if (self).f39 else (self).f28)
                rhs303_ = d_2063_i5_
                rhs304_ = 20
                rhs305_ = (D22_DC58((d_1949_v0_).f34, (self).f29, (self).f28) if (len(d_2062_v78_)) <= ((self).f29) else d_2074_v89_)
                lhs240_ = d_2072_v87_
                lhs241_ = default__.safeIndex(883, (d_2072_v87_).length(0))
                lhs242_ = globalState
                lhs243_ = globalState
                lhs244_ = d_2073_v88_
                lhs245_ = default__.safeIndex(899, (d_2073_v88_).length(0))
                d_2071_v86_ = rhs301_
                lhs240_[lhs241_] = rhs302_
                lhs242_.f2 = rhs303_
                lhs243_.f6 = rhs304_
                lhs244_[lhs245_] = rhs305_
                (globalState).f25 = (d_1949_v0_).fm4(d_2061_v77_, globalState)
                (globalState).f14 = ((d_2069_v84_)[d_2063_i5_] if (d_2063_i5_) in (d_2069_v84_) else (d_1949_v0_).f34)
            elif True:
                d_2076_v91_: _dafny.Array
                def lambda99_(d_2077_v74_):
                    def lambda100_(d_2078_i6_):
                        return (d_2077_v74_) + (_dafny.SeqWithoutIsStrInference([-669, (self).f29, len(d_2077_v74_), (self).f28]))

                    return lambda100_

                init52_ = lambda99_(d_2058_v74_)
                nw324_ = _dafny.Array(None, 14)
                for i0_52_ in range(nw324_.length(0)):
                    nw324_[i0_52_] = init52_(i0_52_)
                d_2076_v91_ = nw324_
                index313_ = default__.safeIndex(996, (d_2076_v91_).length(0))
                (d_2076_v91_)[index313_] = (d_2058_v74_).set(default__.safeIndex((0) - (len(d_2065_v80_)), len(d_2058_v74_)), default__.fm1((self).f39, d_2064_v79_, (self).f27, globalState))
                d_2079_v92_: _dafny.Map
                d_2079_v92_ = _dafny.Map({(self).f29: (self).f27})
                d_2080_v93_: D7
                d_2080_v93_ = D7_DC15(d_2079_v92_)
                d_2081_v94_: _dafny.MultiSet
                d_2081_v94_ = _dafny.MultiSet([(d_1949_v0_).f34])
                index314_ = default__.safeIndex(996, (d_2076_v91_).length(0))
                (d_2076_v91_)[index314_] = (((d_2058_v74_) + (_dafny.SeqWithoutIsStrInference([len((d_2080_v93_).cf24), d_2063_i5_, (d_2081_v94_).cardinality, (self).f28]))).set(default__.safeIndex((0) - (len(d_2065_v80_)), len((d_2058_v74_) + (_dafny.SeqWithoutIsStrInference([len((d_2080_v93_).cf24), d_2063_i5_, (d_2081_v94_).cardinality, (self).f28])))), -761)).set(default__.safeIndex((self).f28, len(((d_2058_v74_) + (_dafny.SeqWithoutIsStrInference([len((d_2080_v93_).cf24), d_2063_i5_, (d_2081_v94_).cardinality, (self).f28]))).set(default__.safeIndex((0) - (len(d_2065_v80_)), len((d_2058_v74_) + (_dafny.SeqWithoutIsStrInference([len((d_2080_v93_).cf24), d_2063_i5_, (d_2081_v94_).cardinality, (self).f28])))), -761))), (self).f28)
                (globalState).f5 = (d_1949_v0_).fm4(D0_DC1((d_1949_v0_).f34), globalState)
                d_2082_v95_: _dafny.Array
                nw325_ = _dafny.Array(None, 17)
                nw325_[int(0)] = d_1950_v1_
                nw325_[int(1)] = default__.fm62(633, (d_1949_v0_).f34, (d_2076_v91_)[default__.safeIndex(996, (d_2076_v91_).length(0))], globalState)
                nw325_[int(2)] = d_1950_v1_
                nw325_[int(3)] = d_1950_v1_
                nw325_[int(4)] = d_1950_v1_
                nw325_[int(5)] = d_1950_v1_
                nw325_[int(6)] = d_1950_v1_
                nw325_[int(7)] = D21_DC53((d_1949_v0_).f34, d_2063_i5_, (d_1949_v0_).f34)
                nw325_[int(8)] = d_1950_v1_
                nw325_[int(9)] = d_1950_v1_
                nw325_[int(10)] = d_1950_v1_
                nw325_[int(11)] = d_1950_v1_
                nw325_[int(12)] = d_1950_v1_
                nw325_[int(13)] = d_1950_v1_
                nw325_[int(14)] = d_1950_v1_
                nw325_[int(15)] = d_1950_v1_
                nw325_[int(16)] = d_1950_v1_
                d_2082_v95_ = nw325_
                index315_ = default__.safeIndex(360, (d_2082_v95_).length(0))
                (d_2082_v95_)[index315_] = d_1950_v1_
                index316_ = default__.safeIndex(360, (d_2082_v95_).length(0))
                (d_2082_v95_)[index316_] = (D21_DC53((d_1949_v0_).f34, len(d_2064_v79_), (self).f27) if (d_1949_v0_).f34 else d_1950_v1_)
                d_2083_v96_: _dafny.Array
                nw326_ = _dafny.Array(int(0), 13)
                d_2083_v96_ = nw326_
                index317_ = default__.safeIndex(575, (d_2083_v96_).length(0))
                (d_2083_v96_)[index317_] = (self).f29
                index318_ = default__.safeIndex(575, (d_2083_v96_).length(0))
                rhs306_ = _dafny.SeqWithoutIsStrInference([(d_1949_v0_).f34])
                rhs307_ = d_2063_i5_
                lhs246_ = d_2083_v96_
                lhs247_ = default__.safeIndex(575, (d_2083_v96_).length(0))
                d_2064_v79_ = rhs306_
                lhs246_[lhs247_] = rhs307_
                (globalState).f6 = d_2063_i5_
        d_2084_v97_: _dafny.MultiSet
        d_2084_v97_ = _dafny.MultiSet([(d_1949_v0_).fm9(len(d_2062_v78_), globalState), (self).f28])
        (globalState).f6 = (default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([True, (self).f27])), (self).f29)) - (((_dafny.MultiSet([(self).f28])).intersection(d_2084_v97_)).cardinality)
        d_2085_i7_: int
        d_2085_i7_ = 0
        with _dafny.label("12"):
            while (self).f27:
                with _dafny.c_label("12"):
                    if (d_2085_i7_) >= (100):
                        raise _dafny.Break("12")
                    d_2085_i7_ = (d_2085_i7_) + (1)
                    d_2086_v98_: _dafny.Set
                    d_2086_v98_ = _dafny.Set({383})
                    (globalState).f2 = len(d_2086_v98_)
                    d_2087_v99_: _dafny.MultiSet
                    d_2087_v99_ = _dafny.MultiSet([(self).f39, (d_1949_v0_).f34])
                    d_2088_v100_: C4
                    nw327_ = C4()
                    nw327_.ctor__((d_2087_v99_) | (d_2087_v99_), (d_1949_v0_).f34)
                    d_2088_v100_ = nw327_
                    d_2089_v101_: _dafny.Map
                    d_2089_v101_ = _dafny.Map({(d_1949_v0_).fm4(d_2061_v77_, globalState): (self).f27})
                    d_2090_v102_: _dafny.Array
                    def lambda101_(d_2091_i8_):
                        return (d_2091_i8_) - ((self).f29)

                    init53_ = lambda101_
                    nw328_ = _dafny.Array(None, 3)
                    for i0_53_ in range(nw328_.length(0)):
                        nw328_[i0_53_] = init53_(i0_53_)
                    d_2090_v102_ = nw328_
                    d_2092_v103_: _dafny.Map
                    d_2092_v103_ = _dafny.Map({d_2089_v101_: d_2090_v102_})
                    d_2092_v103_ = d_2092_v103_
                    d_2093_v104_: _dafny.Seq
                    d_2093_v104_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ks"))
                    d_2094_v105_: C1
                    nw329_ = C1()
                    nw329_.ctor__(((d_1949_v0_).f34 if (self).f39 else default__.fm0(d_2093_v104_, (self).f27, not((d_1949_v0_).f34), d_2086_v98_, globalState)))
                    d_2094_v105_ = nw329_
                    pass
            pass
        d_2095_v106_: _dafny.Seq
        d_2095_v106_ = _dafny.SeqWithoutIsStrInference([(self).f39])
        d_2096_v107_: _dafny.Map
        d_2096_v107_ = _dafny.Map({(self).f39: d_2095_v106_})
        d_2097_v108_: _dafny.Map
        d_2097_v108_ = _dafny.Map({(self).f28: d_2053_v69_})
        d_2098_v109_: _dafny.Array
        nw330_ = _dafny.Array(None, 7)
        nw330_[int(0)] = d_2053_v69_
        nw330_[int(1)] = (d_2054_v70_)[default__.safeIndex(len(d_2096_v107_), len(d_2054_v70_))]
        nw330_[int(2)] = d_2053_v69_
        nw330_[int(3)] = d_2053_v69_
        nw330_[int(4)] = ((d_2097_v108_)[(self).f28] if ((self).f28) in (d_2097_v108_) else d_2053_v69_)
        nw330_[int(5)] = d_2053_v69_
        nw330_[int(6)] = d_2053_v69_
        d_2098_v109_ = nw330_
        r0 = d_2098_v109_
        d_2099_v110_: _dafny.Seq
        d_2099_v110_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rjvisemgf"))
        r1 = (((d_2099_v110_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jdkcgk")))).set(default__.safeIndex((self).f29, len((d_2099_v110_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jdkcgk"))))), d_2055_v71_)) + ((d_2099_v110_ if (self).f39 else d_2099_v110_))
        nw331_ = _dafny.Array(_dafny.Array(None, 0), 21)
        r2 = nw331_
        r3 = (self).f27
        return r0, r1, r2, r3

    @property
    def f39(self):
        return self._f39

class C6(T1, T0):
    def  __init__(self):
        self._f27: bool = False
        self._f28: int = int(0)
        self._f29: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    @property
    def f27(self):
        return self._f27
    @property
    def f28(self):
        return self._f28
    @property
    def f29(self):
        return self._f29
    def ctor__(self, f28, f29, f27):
        (self)._f28 = f28
        (self)._f29 = f29
        (self)._f27 = f27

    def fm2(self, p0, p1, globalState):
        source33_ = D14_DC30(_dafny.Map({_dafny.CodePoint('u'): (D22_DC56(_dafny.Map({(self).f28: _dafny.CodePoint('j')}))).cf83}))
        if source33_.is_DC31:
            d_2100___mcc_h0_ = source33_.cf43
            d_2101_cf43_ = d_2100___mcc_h0_
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "igjvbx"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rhqhor")))
        elif True:
            d_2102___mcc_h1_ = source33_.cf42
            d_2103_cf42_ = d_2102___mcc_h1_
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "amsxhs"))

    def fm3(self, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qxawd")) if True else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lkrldf")))) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_2104_i0_ in range(default__.abs(982))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ygnxpcv"))))

    def m3(self, p0, p1, p2, globalState):
        d_2105_v0_: _dafny.Seq
        d_2105_v0_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True])])
        d_2106_v1_: _dafny.Map
        d_2106_v1_ = _dafny.Map({(self).f27: True})
        d_2107_v2_: _dafny.Seq
        d_2107_v2_ = _dafny.SeqWithoutIsStrInference([(self).f27, ((d_2106_v1_)[(self).f27] if ((self).f27) in (d_2106_v1_) else (self).f27), (self).f27, (self).f27])
        d_2108_v3_: _dafny.Seq
        d_2108_v3_ = _dafny.SeqWithoutIsStrInference([(self).f28, (self).f29, (self).f28])
        def iife188_():
            coll80_ = _dafny.Set()
            compr_80_: int
            for compr_80_ in (d_2108_v3_).Elements:
                d_2109_v4_: int = compr_80_
                if (d_2109_v4_) in (d_2108_v3_):
                    coll80_ = coll80_.union(_dafny.Set([(d_2109_v4_) * ((0) - (len(_dafny.SeqWithoutIsStrInference([not(not(not(not(False))))]))))]))
            return _dafny.Set(coll80_)
        if (_dafny.Set({p2, (self).f28, p2, default__.fm1(p0, (d_2105_v0_)[default__.safeIndex((self).f28, len(d_2105_v0_))], (self).f27, globalState), len(d_2107_v2_)})).isdisjoint(iife188_()
        ):
            (globalState).f14 = p0
            d_2110_v5_: _dafny.Seq
            d_2110_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tddmwx"))
            d_2111_v6_: _dafny.Seq
            d_2111_v6_ = _dafny.SeqWithoutIsStrInference([d_2110_v5_])
            d_2112_v7_: D13
            d_2112_v7_ = D13_DC27(d_2111_v6_)
            source34_ = (d_2112_v7_ if not(False) else d_2112_v7_)
            if source34_.is_DC28:
                d_2113___mcc_h0_ = source34_.cf38
                d_2114_cf38_ = d_2113___mcc_h0_
                d_2115_v8_: str
                d_2115_v8_ = _dafny.CodePoint('e')
                d_2116_v9_: _dafny.MultiSet
                d_2116_v9_ = _dafny.MultiSet([(self).f29])
                d_2117_v10_: _dafny.Map
                d_2117_v10_ = _dafny.Map({p2: d_2116_v9_})
                d_2118_v11_: _dafny.Map
                d_2118_v11_ = _dafny.Map({(0) - (len(d_2117_v10_)): d_2115_v8_})
                d_2119_v12_: _dafny.Map
                d_2119_v12_ = _dafny.Map({d_2115_v8_: d_2118_v11_})
                d_2120_v13_: D18
                d_2120_v13_ = D18_DC42(d_2110_v5_, (self).f27)
                d_2121_v14_: _dafny.Array
                nw332_ = _dafny.Array(None, 27)
                nw332_[int(0)] = (self).fm2(-514, default__.fm27(p0, len(default__.fm28(globalState)), (self).f27, globalState), globalState)
                nw332_[int(1)] = (self).fm2(default__.fm1(p0, d_2107_v2_, (self).f27, globalState), d_2119_v12_, globalState)
                nw332_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yi"))
                nw332_[int(3)] = d_2110_v5_
                nw332_[int(4)] = (self).fm3(globalState)
                nw332_[int(5)] = (d_2110_v5_) + (d_2110_v5_)
                nw332_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wxk"))
                nw332_[int(7)] = d_2110_v5_
                nw332_[int(8)] = d_2110_v5_
                nw332_[int(9)] = (d_2110_v5_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ptqgj")))
                nw332_[int(10)] = d_2110_v5_
                nw332_[int(11)] = (d_2110_v5_) + (d_2110_v5_)
                nw332_[int(12)] = _dafny.SeqWithoutIsStrInference([d_2115_v8_ for d_2122_i0_ in range(default__.abs(925))])
                nw332_[int(13)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wa"))
                nw332_[int(14)] = d_2110_v5_
                nw332_[int(15)] = _dafny.SeqWithoutIsStrInference([d_2115_v8_ for d_2123_i1_ in range(default__.abs(72))])
                nw332_[int(16)] = (self).fm3(globalState)
                nw332_[int(17)] = _dafny.SeqWithoutIsStrInference([(d_2110_v5_)[default__.safeIndex((self).f28, len(d_2110_v5_))] for d_2124_i2_ in range(default__.abs(-300))])
                nw332_[int(18)] = (d_2120_v13_).cf57
                nw332_[int(19)] = (_dafny.SeqWithoutIsStrInference([(d_2110_v5_)[default__.safeIndex((self).f28, len(d_2110_v5_))] for d_2125_i3_ in range(default__.abs(-231))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kpvmow")))
                nw332_[int(20)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lq"))
                nw332_[int(21)] = d_2110_v5_
                nw332_[int(22)] = d_2110_v5_
                nw332_[int(23)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))
                nw332_[int(24)] = d_2110_v5_
                nw332_[int(25)] = d_2110_v5_
                nw332_[int(26)] = (_dafny.SeqWithoutIsStrInference([d_2115_v8_ for d_2126_i4_ in range(default__.abs(-796))])) + (d_2110_v5_)
                d_2121_v14_ = nw332_
                index319_ = default__.safeIndex(381, (d_2121_v14_).length(0))
                (d_2121_v14_)[index319_] = d_2110_v5_
                d_2127_v15_: _dafny.Array
                nw333_ = _dafny.Array(int(0), 11)
                d_2127_v15_ = nw333_
                index320_ = default__.safeIndex(823, (d_2127_v15_).length(0))
                (d_2127_v15_)[index320_] = d_2114_cf38_
                d_2128_v16_: _dafny.Map
                d_2128_v16_ = _dafny.Map({d_2110_v5_: d_2115_v8_})
                index321_ = default__.safeIndex(381, (d_2121_v14_).length(0))
                index322_ = default__.safeIndex(823, (d_2127_v15_).length(0))
                rhs308_ = d_2110_v5_
                rhs309_ = 80
                rhs310_ = p0
                rhs311_ = d_2128_v16_
                lhs248_ = d_2121_v14_
                lhs249_ = default__.safeIndex(381, (d_2121_v14_).length(0))
                lhs250_ = d_2127_v15_
                lhs251_ = default__.safeIndex(823, (d_2127_v15_).length(0))
                lhs252_ = globalState
                lhs248_[lhs249_] = rhs308_
                lhs250_[lhs251_] = rhs309_
                lhs252_.f14 = rhs310_
                d_2128_v16_ = rhs311_
                d_2129_v17_: _dafny.Map
                d_2129_v17_ = _dafny.Map({d_2115_v8_: (self).f29})
                d_2130_v18_: _dafny.Set
                d_2130_v18_ = _dafny.Set({p0})
                d_2131_v19_: D13
                d_2131_v19_ = D13_DC29(len(d_2130_v18_), d_2114_cf38_, p0)
                d_2132_v20_: _dafny.Set
                def iife189_(_pat_let54_0):
                    def iife190_(d_2133_dt__update__tmp_h0_):
                        def iife191_(_pat_let55_0):
                            def iife192_(d_2134_dt__update_hcf39_h0_):
                                return D13_DC29(d_2134_dt__update_hcf39_h0_, (d_2133_dt__update__tmp_h0_).cf40, (d_2133_dt__update__tmp_h0_).cf41)
                            return iife192_(_pat_let55_0)
                        return iife191_((self).f29)
                    return iife190_(_pat_let54_0)
                d_2132_v20_ = _dafny.Set({(self).f27, (self).f27, p0, not((iife189_(d_2131_v19_)).cf41), (self).f27})
                d_2129_v17_ = (d_2129_v17_).set(default__.fm29((d_2127_v15_)[default__.safeIndex(823, (d_2127_v15_).length(0))], p0, globalState), (0) - (len(d_2132_v20_)))
                (globalState).f21 = p0
                (globalState).f6 = (p2) * (len((d_2108_v3_) + (_dafny.SeqWithoutIsStrInference([(self).f29 for d_2135_i5_ in range(default__.abs(-336))]))))
            elif source34_.is_DC29:
                d_2136___mcc_h1_ = source34_.cf39
                d_2137___mcc_h2_ = source34_.cf40
                d_2138___mcc_h3_ = source34_.cf41
                d_2139_cf41_ = d_2138___mcc_h3_
                d_2140_cf40_ = d_2137___mcc_h2_
                d_2141_cf39_ = d_2136___mcc_h1_
                (globalState).f6 = default__.fm1((p0 if p0 else False), d_2107_v2_, False, globalState)
                d_2142_v21_: str
                d_2142_v21_ = _dafny.CodePoint('x')
                d_2143_v22_: _dafny.Map
                d_2143_v22_ = _dafny.Map({d_2142_v21_: d_2139_cf41_})
                d_2144_v25_: _dafny.MultiSet
                d_2144_v25_ = _dafny.MultiSet([(self).f29, d_2141_cf39_])
                d_2145_v26_: _dafny.Seq
                d_2145_v26_ = _dafny.SeqWithoutIsStrInference([d_2106_v1_])
                d_2146_v27_: _dafny.Map
                def iife193_():
                    coll81_ = _dafny.Set()
                    compr_81_: str
                    for compr_81_ in (d_2143_v22_).keys.Elements:
                        d_2147_v23_: str = compr_81_
                        if (d_2147_v23_) in (d_2143_v22_):
                            coll81_ = coll81_.union(_dafny.Set([d_2147_v23_]))
                    return _dafny.Set(coll81_)
                def iife194_():
                    coll82_ = _dafny.Set()
                    compr_82_: str
                    for compr_82_ in (d_2110_v5_).Elements:
                        d_2148_v24_: str = compr_82_
                        if (d_2148_v24_) in (d_2110_v5_):
                            coll82_ = coll82_.union(_dafny.Set([d_2148_v24_]))
                    return _dafny.Set(coll82_)
                d_2146_v27_ = _dafny.Map({(iife193_()
                ) - (iife194_()
                ): (d_2144_v25_).set(len((d_2145_v26_)[default__.safeIndex(p2, len(d_2145_v26_))]), default__.abs(d_2141_cf39_))})
                d_2146_v27_ = d_2146_v27_
                (globalState).f6 = default__.safeDivisionInt(d_2141_cf39_, d_2141_cf39_)
                d_2110_v5_ = d_2110_v5_
            elif True:
                d_2149___mcc_h4_ = source34_.cf37
                d_2150_cf37_ = d_2149___mcc_h4_
                d_2151_v28_: _dafny.Set
                d_2151_v28_ = _dafny.Set({p2, (self).f29, (self).f29})
                d_2152_v29_: _dafny.Map
                d_2152_v29_ = _dafny.Map({d_2110_v5_: default__.fm0(d_2110_v5_, (self).f27, p0, d_2151_v28_, globalState)})
                d_2153_v30_: str
                d_2153_v30_ = _dafny.CodePoint('c')
                d_2154_v31_: _dafny.MultiSet
                d_2154_v31_ = _dafny.MultiSet([d_2153_v30_])
                (globalState).f14 = ((d_2152_v29_)[d_2110_v5_] if (d_2110_v5_) in (d_2152_v29_) else (d_2153_v30_) not in (d_2154_v31_))
                (globalState).f12 = (p2) == (672)
                (globalState).f21 = (self).f27
                d_2110_v5_ = (d_2110_v5_).set(default__.safeIndex((self).f29, len(d_2110_v5_)), d_2153_v30_)
            d_2155_v32_: _dafny.Array
            nw334_ = _dafny.Array(_dafny.MultiSet({}), 17)
            d_2155_v32_ = nw334_
            d_2156_v33_: _dafny.Map
            d_2156_v33_ = _dafny.Map({(0) - (p2): p0})
            d_2157_v34_: _dafny.Map
            d_2157_v34_ = _dafny.Map({len(d_2156_v33_): p0})
            d_2158_v35_: _dafny.MultiSet
            d_2158_v35_ = _dafny.MultiSet([len(d_2157_v34_)])
            index323_ = default__.safeIndex(745, (d_2155_v32_).length(0))
            (d_2155_v32_)[index323_] = (d_2158_v35_) | (d_2158_v35_)
            index324_ = default__.safeIndex(745, (d_2155_v32_).length(0))
            (d_2155_v32_)[index324_] = d_2158_v35_
            d_2159_v36_: _dafny.Map
            d_2159_v36_ = _dafny.Map({d_2106_v1_: (self).f27})
            d_2160_v37_: D23
            d_2160_v37_ = D23_DC59(d_2106_v1_)
            d_2161_v38_: _dafny.MultiSet
            d_2161_v38_ = _dafny.MultiSet([(self).f27, ((d_2159_v36_)[(d_2160_v37_).cf90] if ((d_2160_v37_).cf90) in (d_2159_v36_) else p0)])
            d_2162_v39_: _dafny.Set
            d_2162_v39_ = _dafny.Set({(d_2161_v38_).cardinality, len(_dafny.SeqWithoutIsStrInference([(self).f27, (self).f27, ((d_2156_v33_)[(self).f29] if ((self).f29) in (d_2156_v33_) else (self).f27)])), p2, -696})
            d_2163_v40_: _dafny.Map
            d_2163_v40_ = _dafny.Map({(d_2162_v39_) != (d_2162_v39_): (self).f28})
            def iife195_():
                coll83_ = _dafny.Set()
                compr_83_: int
                for compr_83_ in _dafny.IntegerRange(-43, 262):
                    d_2164_v41_: int = compr_83_
                    if ((-43) <= (d_2164_v41_)) and ((d_2164_v41_) < (262)):
                        coll83_ = coll83_.union(_dafny.Set([default__.safeDivisionInt(d_2164_v41_, p2)]))
                return _dafny.Set(coll83_)
            def iife196_():
                coll84_ = _dafny.Set()
                compr_84_: int
                for compr_84_ in _dafny.IntegerRange(-43, 262):
                    d_2165_v41_: int = compr_84_
                    if ((-43) <= (d_2165_v41_)) and ((d_2165_v41_) < (262)):
                        coll84_ = coll84_.union(_dafny.Set([default__.safeDivisionInt(d_2165_v41_, p2)]))
                return _dafny.Set(coll84_)
            (globalState).f6 = ((d_2163_v40_)[not(default__.fm0(d_2110_v5_, default__.fm0(d_2110_v5_, (self).f27, True, d_2162_v39_, globalState), (self).f27, iife195_()
            , globalState))] if (not(default__.fm0(d_2110_v5_, default__.fm0(d_2110_v5_, (self).f27, True, d_2162_v39_, globalState), (self).f27, iife196_()
            , globalState))) in (d_2163_v40_) else ((self).f29) * (-566))
            d_2166_v42_: D6
            d_2166_v42_ = D6_DC14(p0, 50)
            d_2167_v43_: _dafny.Set
            d_2167_v43_ = _dafny.Set({d_2166_v42_})
            d_2168_v44_: _dafny.Map
            d_2168_v44_ = _dafny.Map({((D0_DC2(p2, (self).f28)).cf2 if p0 else 722): (((d_2158_v35_)[(self).f29] if ((self).f29) in (d_2158_v35_) else (self).f29)) + ((self).f29)})
            rhs312_ = p2
            rhs313_ = d_2108_v3_
            rhs314_ = default__.fm30((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "citltq"))) < (d_2110_v5_), (self).f27, globalState)
            rhs315_ = (d_2168_v44_) | (_dafny.Map({(self).f28: (self).f28}))
            rhs316_ = (d_2108_v3_)[default__.safeIndex((0) - ((p2) * ((self).f29)), len(d_2108_v3_))]
            lhs253_ = globalState
            lhs254_ = globalState
            lhs253_.f2 = rhs312_
            d_2108_v3_ = rhs313_
            d_2167_v43_ = rhs314_
            d_2168_v44_ = rhs315_
            lhs254_.f2 = rhs316_
        elif True:
            d_2169_v45_: D15
            d_2169_v45_ = D15_DC34(p2, ((d_2106_v1_)[not((self).f27)] if (not((self).f27)) in (d_2106_v1_) else p0), p2, (self).f28, p0)
            d_2170_v46_: _dafny.Map
            d_2170_v46_ = _dafny.Map({d_2108_v3_: (d_2169_v45_).cf49})
            d_2171_v47_: _dafny.Seq
            d_2171_v47_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tdav"))
            d_2170_v46_ = (_dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f28 for d_2172_i6_ in range(default__.abs(936))]): (self).f27})).set(d_2108_v3_, (d_2171_v47_) != (d_2171_v47_))
            (globalState).f2 = len((d_2108_v3_) + (d_2108_v3_))
            d_2173_v48_: C2
            nw335_ = C2()
            nw335_.ctor__(False, False)
            d_2173_v48_ = nw335_
            d_2174_v49_: _dafny.Seq
            d_2174_v49_ = _dafny.SeqWithoutIsStrInference([d_2173_v48_])
            d_2175_v50_: _dafny.Map
            d_2175_v50_ = _dafny.Map({(self).f27: (d_2174_v49_)[default__.safeIndex(p2, len(d_2174_v49_))]})
            (globalState).f14 = not((len(d_2175_v50_)) < (p2))
            (globalState).f5 = (self).f27
            d_2176_v51_: str
            d_2176_v51_ = _dafny.CodePoint('l')
            d_2177_v52_: _dafny.Map
            d_2177_v52_ = _dafny.Map({d_2176_v51_: (self).f28})
            d_2178_v53_: _dafny.MultiSet
            d_2178_v53_ = _dafny.MultiSet([d_2177_v52_])
            d_2179_v54_: _dafny.Map
            d_2179_v54_ = _dafny.Map({(self).f29: d_2178_v53_})
            (globalState).f2 = (945) - (len(d_2179_v54_))
        d_2180_v55_: _dafny.Map
        d_2180_v55_ = _dafny.Map({913: p0})
        d_2181_v56_: _dafny.Map
        d_2181_v56_ = _dafny.Map({len((d_2180_v55_).set((self).f29, p0)): (self).f29})
        d_2182_v57_: _dafny.MultiSet
        d_2182_v57_ = _dafny.MultiSet([False])
        d_2183_v58_: D5
        d_2183_v58_ = D5_DC12((self).f27, (self).f28)
        d_2184_v59_: _dafny.Seq
        d_2184_v59_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "njjxwnxew"))
        d_2185_v60_: _dafny.Array
        nw336_ = _dafny.Array(None, 15)
        nw336_[int(0)] = p2
        nw336_[int(1)] = p2
        nw336_[int(2)] = (self).f28
        nw336_[int(3)] = (self).f29
        nw336_[int(4)] = 449
        nw336_[int(5)] = p2
        nw336_[int(6)] = (((d_2181_v56_)[302] if (302) in (d_2181_v56_) else p2)) * (((d_2182_v57_).set((d_2183_v58_).cf19, default__.abs(p2))).cardinality)
        nw336_[int(7)] = p2
        nw336_[int(8)] = (self).f28
        nw336_[int(9)] = p2
        nw336_[int(10)] = (self).f29
        nw336_[int(11)] = (self).f29
        nw336_[int(12)] = len(d_2184_v59_)
        nw336_[int(13)] = p2
        nw336_[int(14)] = 345
        d_2185_v60_ = nw336_
        index325_ = default__.safeIndex(63, (d_2185_v60_).length(0))
        (d_2185_v60_)[index325_] = -280
        index326_ = default__.safeIndex(63, (d_2185_v60_).length(0))
        (d_2185_v60_)[index326_] = (self).f28
        d_2186_v61_: _dafny.Array
        nw337_ = _dafny.Array(_dafny.Map({}), 16)
        d_2186_v61_ = nw337_
        d_2187_v62_: D0
        d_2187_v62_ = D0_DC1((self).f27)
        pat_let_tv46_ = p0
        d_2188_v63_: _dafny.Array
        nw338_ = _dafny.Array(None, 14)
        def iife197_(_pat_let56_0):
            def iife198_(d_2189_dt__update__tmp_h1_):
                def iife199_(_pat_let57_0):
                    def iife200_(d_2190_dt__update_hcf1_h0_):
                        return D0_DC1(d_2190_dt__update_hcf1_h0_)
                    return iife200_(_pat_let57_0)
                return iife199_(pat_let_tv46_)
            return iife198_(_pat_let56_0)
        nw338_[int(0)] = iife197_(D0_DC1(p0))
        nw338_[int(1)] = d_2187_v62_
        nw338_[int(2)] = d_2187_v62_
        nw338_[int(3)] = d_2187_v62_
        nw338_[int(4)] = d_2187_v62_
        nw338_[int(5)] = d_2187_v62_
        nw338_[int(6)] = d_2187_v62_
        nw338_[int(7)] = d_2187_v62_
        nw338_[int(8)] = d_2187_v62_
        nw338_[int(9)] = d_2187_v62_
        nw338_[int(10)] = d_2187_v62_
        nw338_[int(11)] = d_2187_v62_
        nw338_[int(12)] = d_2187_v62_
        nw338_[int(13)] = d_2187_v62_
        d_2188_v63_ = nw338_
        d_2191_v64_: _dafny.Map
        d_2191_v64_ = _dafny.Map({-33: d_2188_v63_})
        index327_ = default__.safeIndex(598, (d_2186_v61_).length(0))
        (d_2186_v61_)[index327_] = d_2191_v64_
        index328_ = default__.safeIndex(598, (d_2186_v61_).length(0))
        index329_ = default__.safeIndex(63, (d_2185_v60_).length(0))
        rhs317_ = d_2191_v64_
        rhs318_ = (d_2185_v60_)[default__.safeIndex(63, (d_2185_v60_).length(0))]
        rhs319_ = len(_dafny.SeqWithoutIsStrInference([D23_DC59(d_2106_v1_) for d_2192_i7_ in range(default__.abs(309))]))
        rhs320_ = default__.safeDivisionInt((self).f29, len((d_2184_v59_) + (d_2184_v59_)))
        rhs321_ = ((self).f27 if ((d_2106_v1_)[(self).f27] if ((self).f27) in (d_2106_v1_) else p0) else (self).f27)
        lhs255_ = d_2186_v61_
        lhs256_ = default__.safeIndex(598, (d_2186_v61_).length(0))
        lhs257_ = globalState
        lhs258_ = d_2185_v60_
        lhs259_ = default__.safeIndex(63, (d_2185_v60_).length(0))
        lhs260_ = globalState
        lhs261_ = globalState
        lhs255_[lhs256_] = rhs317_
        lhs257_.f2 = rhs318_
        lhs258_[lhs259_] = rhs319_
        lhs260_.f6 = rhs320_
        lhs261_.f21 = rhs321_
        d_2193_v65_: str
        d_2193_v65_ = _dafny.CodePoint('o')
        d_2194_v66_: _dafny.Map
        d_2194_v66_ = _dafny.Map({d_2193_v65_: (self).f27})
        rhs322_ = d_2182_v57_
        rhs323_ = p0
        rhs324_ = (p0) not in (_dafny.Map({p0: (self).f27}))
        rhs325_ = (len(d_2194_v66_)) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.Set({(d_2185_v60_)[default__.safeIndex(63, (d_2185_v60_).length(0))]})])) for d_2195_i8_ in range(default__.abs(540))]))
        lhs262_ = globalState
        lhs263_ = globalState
        lhs264_ = globalState
        d_2182_v57_ = rhs322_
        lhs262_.f25 = rhs323_
        lhs263_.f12 = rhs324_
        lhs264_.f5 = rhs325_
        index330_ = default__.safeIndex(63, (d_2185_v60_).length(0))
        (d_2185_v60_)[index330_] = (self).f29
        d_2196_v67_: _dafny.Map
        d_2196_v67_ = _dafny.Map({(self).f28: d_2193_v65_})
        d_2197_v68_: _dafny.Map
        d_2197_v68_ = _dafny.Map({d_2193_v65_: d_2196_v67_})
        d_2198_v69_: D14
        d_2198_v69_ = D14_DC30((d_2197_v68_) | (d_2197_v68_))
        source35_ = d_2198_v69_
        if source35_.is_DC31:
            d_2199___mcc_h5_ = source35_.cf43
            d_2200_cf43_ = d_2199___mcc_h5_
            d_2201_v70_: _dafny.Map
            d_2201_v70_ = _dafny.Map({p0: (0) - (d_2200_cf43_)})
            d_2201_v70_ = (d_2201_v70_).set(p0, (self).f29)
            (globalState).f12 = (self).f27
            d_2202_v71_: _dafny.Map
            d_2202_v71_ = _dafny.Map({(self).f27: d_2184_v59_})
            d_2202_v71_ = ((d_2202_v71_) | (_dafny.Map({p0: d_2184_v59_})) if (p0 if p0 else (self).f27) else default__.fm31(globalState))
            (globalState).f6 = (default__.fm1((self).f27, d_2107_v2_, p0, globalState)) * (d_2200_cf43_)
        elif True:
            d_2203___mcc_h6_ = source35_.cf42
            d_2204_cf42_ = d_2203___mcc_h6_
            d_2205_v72_: C2
            nw339_ = C2()
            nw339_.ctor__(p0, not(p0))
            d_2205_v72_ = nw339_
            d_2206_v73_: _dafny.MultiSet
            d_2206_v73_ = _dafny.MultiSet([d_2205_v72_, d_2205_v72_, d_2205_v72_, d_2205_v72_, d_2205_v72_])
            d_2207_v74_: _dafny.Map
            d_2207_v74_ = _dafny.Map({d_2206_v73_: (d_2184_v59_) + (d_2184_v59_)})
            rhs326_ = d_2185_v60_
            rhs327_ = ((d_2207_v74_)[d_2206_v73_] if (d_2206_v73_) in (d_2207_v74_) else d_2184_v59_)
            d_2185_v60_ = rhs326_
            d_2184_v59_ = rhs327_
            (globalState).f21 = (d_2107_v2_)[default__.safeIndex(p2, len(d_2107_v2_))]
            d_2208_v75_: C0
            nw340_ = C0()
            nw340_.ctor__()
            d_2208_v75_ = nw340_
            (globalState).f2 = -614

    def m1(self, globalState):
        r0: int = int(0)
        r0 = (self).f28
        d_2209_v0_: _dafny.Map
        d_2209_v0_ = _dafny.Map({(self).f29: (self).f27})
        d_2210_v1_: _dafny.Seq
        d_2210_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))
        d_2211_i0_: int
        d_2211_i0_ = 0
        with _dafny.label("13"):
            while ((d_2209_v0_)[len(d_2210_v1_)] if (len(d_2210_v1_)) in (d_2209_v0_) else (not((self).f27) if not((self).f27) else False)):
                with _dafny.c_label("13"):
                    if (d_2211_i0_) >= (100):
                        raise _dafny.Break("13")
                    d_2211_i0_ = (d_2211_i0_) + (1)
                    d_2212_v2_: _dafny.Array
                    def lambda102_(d_2213_i1_):
                        return (self).f27

                    init54_ = lambda102_
                    nw341_ = _dafny.Array(None, 26)
                    for i0_54_ in range(nw341_.length(0)):
                        nw341_[i0_54_] = init54_(i0_54_)
                    d_2212_v2_ = nw341_
                    d_2214_v3_: _dafny.Set
                    d_2214_v3_ = _dafny.Set({d_2212_v2_, d_2212_v2_})
                    d_2214_v3_ = (d_2214_v3_) | (_dafny.Set({d_2212_v2_}))
                    d_2215_v4_: C0
                    nw342_ = C0()
                    nw342_.ctor__()
                    d_2215_v4_ = nw342_
                    d_2215_v4_ = d_2215_v4_
                    d_2216_v5_: D15
                    d_2216_v5_ = D15_DC34((self).f29, not ((self).f27) or ((self).f27), ((self).f29) + ((self).f28), (self).f29, (self).f27)
                    source36_ = d_2216_v5_
                    if source36_.is_DC33:
                        d_2217_v6_: _dafny.Array
                        def lambda103_(d_2218_i2_):
                            return default__.safeDivisionInt(d_2218_i2_, -1)

                        init55_ = lambda103_
                        nw343_ = _dafny.Array(None, 15)
                        for i0_55_ in range(nw343_.length(0)):
                            nw343_[i0_55_] = init55_(i0_55_)
                        d_2217_v6_ = nw343_
                        d_2219_v7_: _dafny.Seq
                        d_2219_v7_ = _dafny.SeqWithoutIsStrInference([(self).f27])
                        index331_ = default__.safeIndex(979, (d_2217_v6_).length(0))
                        (d_2217_v6_)[index331_] = default__.fm1((self).f27, d_2219_v7_, not((d_2219_v7_)[default__.safeIndex((self).f28, len(d_2219_v7_))]), globalState)
                        d_2220_v8_: str
                        d_2220_v8_ = _dafny.CodePoint('m')
                        d_2221_v9_: _dafny.Array
                        nw344_ = _dafny.Array(None, 24)
                        d_2221_v9_ = nw344_
                        d_2222_v10_: C2
                        nw345_ = C2()
                        nw345_.ctor__((self).f27, (d_2219_v7_)[default__.safeIndex(len(d_2210_v1_), len(d_2219_v7_))])
                        d_2222_v10_ = nw345_
                        index332_ = default__.safeIndex(453, (d_2221_v9_).length(0))
                        (d_2221_v9_)[index332_] = d_2222_v10_
                        index333_ = default__.safeIndex(979, (d_2217_v6_).length(0))
                        index334_ = default__.safeIndex(453, (d_2221_v9_).length(0))
                        rhs328_ = (self).f28
                        rhs329_ = d_2220_v8_
                        rhs330_ = d_2222_v10_
                        lhs265_ = d_2217_v6_
                        lhs266_ = default__.safeIndex(979, (d_2217_v6_).length(0))
                        lhs267_ = d_2221_v9_
                        lhs268_ = default__.safeIndex(453, (d_2221_v9_).length(0))
                        lhs265_[lhs266_] = rhs328_
                        d_2220_v8_ = rhs329_
                        lhs267_[lhs268_] = rhs330_
                        d_2223_v11_: D20
                        d_2223_v11_ = D20_DC50()
                        d_2224_v12_: D20
                        d_2224_v12_ = D20_DC51(d_2223_v11_)
                        rhs331_ = (self).f27
                        rhs332_ = ((self).f28) + ((d_2217_v6_)[default__.safeIndex(979, (d_2217_v6_).length(0))])
                        rhs333_ = d_2224_v12_
                        rhs334_ = (d_2220_v8_ if (d_2222_v10_).f34 else d_2220_v8_)
                        lhs269_ = globalState
                        lhs270_ = globalState
                        lhs269_.f25 = rhs331_
                        lhs270_.f2 = rhs332_
                        d_2224_v12_ = rhs333_
                        d_2220_v8_ = rhs334_
                        d_2225_v13_: _dafny.Seq
                        d_2225_v13_ = _dafny.SeqWithoutIsStrInference([d_2222_v10_])
                        d_2226_v14_: _dafny.Map
                        d_2226_v14_ = _dafny.Map({d_2225_v13_: (d_2210_v1_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wrq")))})
                        d_2226_v14_ = d_2226_v14_
                        d_2227_v15_: _dafny.Set
                        d_2227_v15_ = _dafny.Set({(self).f27})
                        rhs335_ = default__.fm32((default__.fm33(False, (d_2222_v10_).f34, globalState)) == (_dafny.SeqWithoutIsStrInference([(0) - ((self).f29), (d_2217_v6_)[default__.safeIndex(979, (d_2217_v6_).length(0))], (self).f28, (d_2217_v6_)[default__.safeIndex(979, (d_2217_v6_).length(0))], (self).f28])), (default__.fm33((d_2222_v10_).f34, (d_2222_v10_).f34, globalState) if (d_2222_v10_).f34 else _dafny.SeqWithoutIsStrInference([(self).f29, (self).f29])), ((d_2222_v10_).fm9((self).f29, globalState)) > ((self).f28), globalState)
                        rhs336_ = (d_2222_v10_).f34
                        rhs337_ = (_dafny.Set({(self).f27})) | (d_2227_v15_)
                        rhs338_ = d_2217_v6_
                        lhs271_ = globalState
                        d_2219_v7_ = rhs335_
                        lhs271_.f14 = rhs336_
                        d_2227_v15_ = rhs337_
                        d_2217_v6_ = rhs338_
                    elif source36_.is_DC34:
                        d_2228___mcc_h0_ = source36_.cf45
                        d_2229___mcc_h1_ = source36_.cf46
                        d_2230___mcc_h2_ = source36_.cf47
                        d_2231___mcc_h3_ = source36_.cf48
                        d_2232___mcc_h4_ = source36_.cf49
                        d_2233_cf49_ = d_2232___mcc_h4_
                        d_2234_cf48_ = d_2231___mcc_h3_
                        d_2235_cf47_ = d_2230___mcc_h2_
                        d_2236_cf46_ = d_2229___mcc_h1_
                        d_2237_cf45_ = d_2228___mcc_h0_
                        d_2238_v16_: C0
                        nw346_ = C0()
                        nw346_.ctor__()
                        d_2238_v16_ = nw346_
                        d_2239_v17_: _dafny.Array
                        nw347_ = _dafny.Array(_dafny.MultiSet({}), 17)
                        d_2239_v17_ = nw347_
                        index335_ = default__.safeIndex(982, (d_2239_v17_).length(0))
                        (d_2239_v17_)[index335_] = _dafny.MultiSet([d_2234_cf48_, (self).f28])
                        index336_ = default__.safeIndex(982, (d_2239_v17_).length(0))
                        (d_2239_v17_)[index336_] = _dafny.MultiSet([d_2237_cf45_, default__.safeDivisionInt(461, d_2237_cf45_)])
                        d_2240_v18_: _dafny.Array
                        nw348_ = _dafny.Array(_dafny.Set({}), 26)
                        d_2240_v18_ = nw348_
                        d_2241_v19_: _dafny.Set
                        d_2241_v19_ = _dafny.Set({(self).f29, d_2234_cf48_})
                        index337_ = default__.safeIndex(370, (d_2240_v18_).length(0))
                        (d_2240_v18_)[index337_] = (d_2241_v19_) | (d_2241_v19_)
                        d_2242_v20_: _dafny.Seq
                        d_2242_v20_ = _dafny.SeqWithoutIsStrInference([d_2237_cf45_])
                        d_2243_v21_: _dafny.Seq
                        d_2243_v21_ = _dafny.SeqWithoutIsStrInference([d_2242_v20_, d_2242_v20_])
                        d_2244_v22_: C1
                        nw349_ = C1()
                        nw349_.ctor__((self).f27)
                        d_2244_v22_ = nw349_
                        d_2245_v23_: _dafny.Set
                        d_2245_v23_ = _dafny.Set({d_2244_v22_, d_2244_v22_, d_2244_v22_})
                        index338_ = default__.safeIndex(370, (d_2240_v18_).length(0))
                        rhs339_ = ((d_2209_v0_)[len(d_2243_v21_)] if (len(d_2243_v21_)) in (d_2209_v0_) else d_2236_cf46_)
                        rhs340_ = len((d_2245_v23_) - ((d_2245_v23_) | (_dafny.Set({d_2244_v22_}))))
                        rhs341_ = d_2241_v19_
                        lhs272_ = globalState
                        lhs273_ = globalState
                        lhs274_ = d_2240_v18_
                        lhs275_ = default__.safeIndex(370, (d_2240_v18_).length(0))
                        lhs272_.f12 = rhs339_
                        lhs273_.f6 = rhs340_
                        lhs274_[lhs275_] = rhs341_
                        d_2236_cf46_ = d_2236_cf46_
                    elif True:
                        d_2246___mcc_h5_ = source36_.cf44
                        d_2247_cf44_ = d_2246___mcc_h5_
                        (globalState).f14 = not ((self).f27) or ((self).f27)
                        d_2248_v26_: _dafny.Seq
                        d_2248_v26_ = _dafny.SeqWithoutIsStrInference([(self).f29, (self).f29])
                        d_2249_v27_: _dafny.Map
                        d_2249_v27_ = _dafny.Map({(self).f28: d_2248_v26_})
                        d_2250_v28_: _dafny.Map
                        def iife201_():
                            def iife203_():
                                coll87_ = _dafny.Set()
                                compr_87_: int
                                for compr_87_ in _dafny.IntegerRange(337, 318):
                                    d_2253_v24_: int = compr_87_
                                    if ((337) <= (d_2253_v24_)) and ((d_2253_v24_) < (318)):
                                        coll87_ = coll87_.union(_dafny.Set([default__.safeModuloInt(d_2253_v24_, (self).f29)]))
                                return _dafny.Set(coll87_)
                            coll85_ = _dafny.Set()
                            def iife202_():
                                coll86_ = _dafny.Set()
                                compr_86_: int
                                for compr_86_ in _dafny.IntegerRange(337, 318):
                                    d_2251_v24_: int = compr_86_
                                    if ((337) <= (d_2251_v24_)) and ((d_2251_v24_) < (318)):
                                        coll86_ = coll86_.union(_dafny.Set([default__.safeModuloInt(d_2251_v24_, (self).f29)]))
                                return _dafny.Set(coll86_)
                            compr_85_: int
                            for compr_85_ in (iife202_()
                            ).Elements:
                                d_2252_v25_: int = compr_85_
                                if (d_2252_v25_) in (iife203_()
                                ):
                                    coll85_ = coll85_.union(_dafny.Set([(d_2252_v25_) * (715)]))
                            return _dafny.Set(coll85_)
                        d_2250_v28_ = _dafny.Map({iife201_()
                        : ((d_2249_v27_)[(self).f28] if ((self).f28) in (d_2249_v27_) else d_2248_v26_)})
                        d_2254_v30_: _dafny.Set
                        d_2254_v30_ = _dafny.Set({(d_2248_v26_)[default__.safeIndex((self).f28, len(d_2248_v26_))]})
                        d_2255_v31_: _dafny.MultiSet
                        d_2255_v31_ = _dafny.MultiSet([d_2254_v30_])
                        def iife204_():
                            coll88_ = _dafny.Map()
                            compr_88_: _dafny.Set
                            for compr_88_ in (d_2255_v31_).Elements:
                                d_2256_v29_: _dafny.Set = compr_88_
                                if (d_2256_v29_) in (d_2255_v31_):
                                    coll88_[d_2256_v29_] = d_2248_v26_
                            return _dafny.Map(coll88_)
                        d_2250_v28_ = ((iife204_()
                        ) | (d_2250_v28_)) | (_dafny.Map({d_2254_v30_: d_2248_v26_}))
                        d_2257_v32_: _dafny.MultiSet
                        d_2257_v32_ = _dafny.MultiSet([False])
                        (globalState).f5 = (_dafny.MultiSet([(self).f27, (self).f27])).ispropersubset((d_2257_v32_) - (d_2257_v32_))
                        d_2258_v33_: str
                        d_2258_v33_ = _dafny.CodePoint('s')
                        d_2259_v34_: _dafny.Map
                        d_2259_v34_ = _dafny.Map({(self).f28: d_2258_v33_})
                        d_2260_v35_: _dafny.Map
                        d_2260_v35_ = _dafny.Map({d_2258_v33_: (d_2259_v34_).set((self).f28, d_2258_v33_)})
                        d_2210_v1_ = (self).fm2((self).f28, d_2260_v35_, globalState)
                    (globalState).f2 = (0) - ((self).f29)
                    pass
            pass
        if ((self).f29) != (602):
            d_2261_v36_: C0
            nw350_ = C0()
            nw350_.ctor__()
            d_2261_v36_ = nw350_
            d_2262_v37_: str
            d_2262_v37_ = _dafny.CodePoint('p')
            d_2263_v38_: _dafny.Map
            d_2263_v38_ = _dafny.Map({d_2262_v37_: False})
            d_2264_v39_: _dafny.Map
            d_2264_v39_ = _dafny.Map({d_2261_v36_: ((d_2263_v38_)[d_2262_v37_] if (d_2262_v37_) in (d_2263_v38_) else (self).f27)})
            d_2265_v40_: _dafny.Seq
            d_2265_v40_ = _dafny.SeqWithoutIsStrInference([len((d_2264_v39_).set(d_2261_v36_, (self).f27)), 767])
            rhs342_ = (self).f27
            rhs343_ = ((d_2265_v40_).set(default__.safeIndex((self).f28, len(d_2265_v40_)), (self).f29)) + (d_2265_v40_)
            rhs344_ = (self).f27
            lhs276_ = globalState
            lhs277_ = globalState
            lhs276_.f12 = rhs342_
            d_2265_v40_ = rhs343_
            lhs277_.f5 = rhs344_
            rhs345_ = (self).f28
            rhs346_ = d_2210_v1_
            rhs347_ = d_2262_v37_
            lhs278_ = globalState
            lhs278_.f6 = rhs345_
            d_2210_v1_ = rhs346_
            d_2262_v37_ = rhs347_
            d_2266_v41_: _dafny.MultiSet
            d_2266_v41_ = _dafny.MultiSet([(self).f28])
            d_2267_v42_: _dafny.MultiSet
            d_2267_v42_ = _dafny.MultiSet([_dafny.CodePoint('a')])
            d_2268_v43_: _dafny.Seq
            d_2268_v43_ = _dafny.SeqWithoutIsStrInference([True])
            d_2269_v44_: _dafny.Set
            d_2269_v44_ = _dafny.Set({-16})
            d_2270_v45_: _dafny.Array
            nw351_ = _dafny.Array(None, 20)
            nw351_[int(0)] = (0) - (((d_2266_v41_)[(self).f28] if ((self).f28) in (d_2266_v41_) else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kg")))))
            nw351_[int(1)] = (self).f28
            nw351_[int(2)] = ((self).f28 if (self).f27 else 187)
            nw351_[int(3)] = (self).f29
            nw351_[int(4)] = (self).f28
            nw351_[int(5)] = (self).f28
            nw351_[int(6)] = default__.safeModuloInt(((d_2267_v42_)[d_2262_v37_] if (d_2262_v37_) in (d_2267_v42_) else (self).f29), (self).f29)
            nw351_[int(7)] = (self).f29
            nw351_[int(8)] = (82) - ((self).f28)
            nw351_[int(9)] = (self).f28
            nw351_[int(10)] = 615
            nw351_[int(11)] = (self).f28
            nw351_[int(12)] = len(d_2268_v43_)
            nw351_[int(13)] = (self).f28
            nw351_[int(14)] = len((d_2210_v1_ if default__.fm0(d_2210_v1_, (self).f27, (self).f27, d_2269_v44_, globalState) else d_2210_v1_))
            nw351_[int(15)] = (self).f29
            nw351_[int(16)] = (self).f29
            nw351_[int(17)] = len(d_2269_v44_)
            nw351_[int(18)] = (self).f29
            nw351_[int(19)] = (self).f28
            d_2270_v45_ = nw351_
            index339_ = default__.safeIndex(247, (d_2270_v45_).length(0))
            (d_2270_v45_)[index339_] = 299
            index340_ = default__.safeIndex(247, (d_2270_v45_).length(0))
            (d_2270_v45_)[index340_] = (self).f28
            d_2271_v46_: D18
            d_2271_v46_ = D18_DC44((self).f27, (self).f28, (self).f27)
            d_2272_v47_: _dafny.Array
            nw352_ = _dafny.Array(None, 24)
            nw352_[int(0)] = (self).f27
            nw352_[int(1)] = (self).f27
            nw352_[int(2)] = (self).f27
            nw352_[int(3)] = (self).f27
            nw352_[int(4)] = (self).f27
            nw352_[int(5)] = (self).f27
            nw352_[int(6)] = (self).f27
            nw352_[int(7)] = (self).f27
            nw352_[int(8)] = default__.fm0(d_2210_v1_, (self).f27, (d_2268_v43_)[default__.safeIndex((self).f29, len(d_2268_v43_))], d_2269_v44_, globalState)
            nw352_[int(9)] = (self).f27
            nw352_[int(10)] = (self).f27
            nw352_[int(11)] = (self).f27
            nw352_[int(12)] = (self).f27
            nw352_[int(13)] = True
            nw352_[int(14)] = (self).f27
            nw352_[int(15)] = (d_2271_v46_).cf60
            nw352_[int(16)] = (d_2268_v43_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_2210_v1_])), len(d_2268_v43_))]
            nw352_[int(17)] = (self).f27
            nw352_[int(18)] = (self).f27
            nw352_[int(19)] = (self).f27
            nw352_[int(20)] = True
            nw352_[int(21)] = (self).f27
            nw352_[int(22)] = (self).f27
            nw352_[int(23)] = (self).f27
            d_2272_v47_ = nw352_
            d_2273_v48_: _dafny.Seq
            d_2273_v48_ = _dafny.SeqWithoutIsStrInference([d_2272_v47_])
            d_2274_v49_: _dafny.Map
            d_2274_v49_ = _dafny.Map({(_dafny.Set({(self).f29})) | (d_2269_v44_): (d_2273_v48_) + (d_2273_v48_)})
            def iife205_():
                coll89_ = _dafny.Set()
                compr_89_: int
                for compr_89_ in _dafny.IntegerRange(640, 618):
                    d_2275_v50_: int = compr_89_
                    if ((640) <= (d_2275_v50_)) and ((d_2275_v50_) < (618)):
                        coll89_ = coll89_.union(_dafny.Set([default__.safeModuloInt(d_2275_v50_, len(d_2269_v44_))]))
                return _dafny.Set(coll89_)
            def iife206_():
                coll90_ = _dafny.Set()
                compr_90_: int
                for compr_90_ in _dafny.IntegerRange(640, 618):
                    d_2276_v50_: int = compr_90_
                    if ((640) <= (d_2276_v50_)) and ((d_2276_v50_) < (618)):
                        coll90_ = coll90_.union(_dafny.Set([default__.safeModuloInt(d_2276_v50_, len(d_2269_v44_))]))
                return _dafny.Set(coll90_)
            (globalState).f6 = len(((d_2274_v49_)[iife205_()
            ] if (iife206_()
            ) in (d_2274_v49_) else (d_2273_v48_) + (d_2273_v48_)))
            d_2277_v51_: C1
            nw353_ = C1()
            nw353_.ctor__(not((self).f27))
            d_2277_v51_ = nw353_
            d_2278_v52_: _dafny.Seq
            d_2278_v52_ = _dafny.SeqWithoutIsStrInference([d_2277_v51_, d_2277_v51_])
            d_2277_v51_ = (d_2278_v52_)[default__.safeIndex((self).f29, len(d_2278_v52_))]
        elif True:
            d_2279_v53_: _dafny.Seq
            d_2279_v53_ = _dafny.SeqWithoutIsStrInference([(self).f29, (self).f28])
            d_2280_v54_: _dafny.Map
            d_2280_v54_ = _dafny.Map({(self).f27: (len(d_2279_v53_) if True else (self).f28)})
            d_2280_v54_ = (d_2280_v54_).set((d_2279_v53_) < (d_2279_v53_), (self).f29)
            if not(not ((self).f27) or ((self).f27)):
                (globalState).f2 = (self).f29
                d_2281_v55_: _dafny.Array
                def lambda104_(d_2282_i3_):
                    return default__.safeModuloInt(d_2282_i3_, (self).f29)

                init56_ = lambda104_
                nw354_ = _dafny.Array(None, 20)
                for i0_56_ in range(nw354_.length(0)):
                    nw354_[i0_56_] = init56_(i0_56_)
                d_2281_v55_ = nw354_
                d_2281_v55_ = d_2281_v55_
                (globalState).f6 = (self).f29
                d_2283_v56_: _dafny.Array
                def lambda105_(d_2284_i4_):
                    return _dafny.CodePoint('r')

                init57_ = lambda105_
                nw355_ = _dafny.Array(None, 17)
                for i0_57_ in range(nw355_.length(0)):
                    nw355_[i0_57_] = init57_(i0_57_)
                d_2283_v56_ = nw355_
                d_2285_v57_: _dafny.Array
                nw356_ = _dafny.Array(False, 6)
                d_2285_v57_ = nw356_
                index341_ = default__.safeIndex(524, (d_2285_v57_).length(0))
                (d_2285_v57_)[index341_] = not(not((self).f27))
                index342_ = default__.safeIndex(524, (d_2285_v57_).length(0))
                rhs348_ = d_2283_v56_
                rhs349_ = ((0) - ((self).f28)) > (default__.safeModuloInt((self).f29, (self).f28))
                rhs350_ = default__.safeDivisionInt((self).f28, default__.safeModuloInt((self).f29, (self).f28))
                lhs279_ = d_2285_v57_
                lhs280_ = default__.safeIndex(524, (d_2285_v57_).length(0))
                lhs281_ = globalState
                d_2283_v56_ = rhs348_
                lhs279_[lhs280_] = rhs349_
                lhs281_.f2 = rhs350_
                index343_ = default__.safeIndex(524, (d_2285_v57_).length(0))
                (d_2285_v57_)[index343_] = (self).f27
            elif True:
                d_2286_v59_: _dafny.Array
                def lambda106_(d_2287_i5_):
                    def iife207_():
                        coll91_ = _dafny.Set()
                        compr_91_: int
                        for compr_91_ in _dafny.IntegerRange(921, 857):
                            d_2288_v58_: int = compr_91_
                            if ((921) <= (d_2288_v58_)) and ((d_2288_v58_) < (857)):
                                coll91_ = coll91_.union(_dafny.Set([(d_2288_v58_) + ((D6_DC14((self).f27, (self).f29)).cf23)]))
                        return _dafny.Set(coll91_)
                    return iife207_()
                    

                init58_ = lambda106_
                nw357_ = _dafny.Array(None, 20)
                for i0_58_ in range(nw357_.length(0)):
                    nw357_[i0_58_] = init58_(i0_58_)
                d_2286_v59_ = nw357_
                d_2289_v60_: _dafny.Map
                d_2289_v60_ = _dafny.Map({(self).f27: d_2286_v59_})
                d_2289_v60_ = (d_2289_v60_).set(not(((self).f27) and ((self).f27)), d_2286_v59_)
                d_2290_v61_: _dafny.Array
                nw358_ = _dafny.Array(_dafny.Array(None, 0), 2)
                d_2290_v61_ = nw358_
                d_2291_v62_: D13
                d_2291_v62_ = D13_DC28((self).f29)
                d_2292_v63_: _dafny.Array
                nw359_ = _dafny.Array(None, 20)
                nw359_[int(0)] = (self).f28
                nw359_[int(1)] = (self).f28
                nw359_[int(2)] = (self).f29
                nw359_[int(3)] = (0) - ((self).f28)
                nw359_[int(4)] = (self).f28
                nw359_[int(5)] = (self).f28
                nw359_[int(6)] = len(d_2209_v0_)
                nw359_[int(7)] = (0) - ((self).f28)
                nw359_[int(8)] = (self).f29
                nw359_[int(9)] = (self).f28
                nw359_[int(10)] = (self).f28
                nw359_[int(11)] = (self).f28
                nw359_[int(12)] = (self).f29
                nw359_[int(13)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mrxmhjgqm")))
                nw359_[int(14)] = (self).f28
                nw359_[int(15)] = (d_2291_v62_).cf38
                nw359_[int(16)] = (self).f28
                nw359_[int(17)] = (self).f28
                nw359_[int(18)] = (self).f28
                nw359_[int(19)] = (self).f28
                d_2292_v63_ = nw359_
                index344_ = default__.safeIndex(171, (d_2290_v61_).length(0))
                (d_2290_v61_)[index344_] = d_2292_v63_
                index345_ = default__.safeIndex(171, (d_2290_v61_).length(0))
                (d_2290_v61_)[index345_] = d_2292_v63_
                (globalState).f25 = (self).f27
                d_2293_v64_: _dafny.Array
                nw360_ = _dafny.Array(None, 1)
                nw360_[int(0)] = (self).f27
                d_2293_v64_ = nw360_
                d_2294_v65_: _dafny.Map
                d_2294_v65_ = _dafny.Map({d_2293_v64_: (self).f29})
                d_2294_v65_ = d_2294_v65_
                d_2295_v66_: _dafny.MultiSet
                d_2295_v66_ = _dafny.MultiSet([(self).f27])
                d_2296_v67_: C0
                nw361_ = C0()
                nw361_.ctor__()
                d_2296_v67_ = nw361_
                d_2297_v68_: _dafny.Map
                d_2297_v68_ = _dafny.Map({(self).f27: d_2296_v67_})
                d_2298_v69_: _dafny.Map
                d_2298_v69_ = _dafny.Map({(self).f28: (d_2295_v66_).set((self).f27, default__.abs(len(d_2297_v68_)))})
                d_2298_v69_ = (d_2298_v69_).set((self).f28, d_2295_v66_)
            d_2299_v70_: _dafny.Seq
            d_2299_v70_ = _dafny.SeqWithoutIsStrInference([not((self).f27)])
            d_2300_v71_: _dafny.Set
            d_2300_v71_ = _dafny.Set({(self).f27, (self).f27, (self).f27})
            d_2301_v72_: _dafny.Set
            d_2301_v72_ = _dafny.Set({default__.fm1((self).f27, d_2299_v70_, (self).f27, globalState)})
            d_2302_v73_: str
            d_2302_v73_ = _dafny.CodePoint('g')
            d_2303_v74_: _dafny.Array
            nw362_ = _dafny.Array(None, 15)
            nw362_[int(0)] = 870
            nw362_[int(1)] = (self).f28
            nw362_[int(2)] = default__.fm1((self).f27, d_2299_v70_, (self).f27, globalState)
            nw362_[int(3)] = (self).f29
            nw362_[int(4)] = len(d_2300_v71_)
            nw362_[int(5)] = len(d_2301_v72_)
            nw362_[int(6)] = default__.fm1((self).f27, d_2299_v70_, (self).f27, globalState)
            nw362_[int(7)] = (self).f28
            nw362_[int(8)] = (len(d_2210_v1_)) - ((self).f29)
            nw362_[int(9)] = (self).f28
            nw362_[int(10)] = (self).f28
            nw362_[int(11)] = (self).f29
            nw362_[int(12)] = len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lwj"))) + (d_2210_v1_)).set(default__.safeIndex((self).f28, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lwj"))) + (d_2210_v1_))), d_2302_v73_))
            nw362_[int(13)] = (self).f29
            nw362_[int(14)] = (self).f28
            d_2303_v74_ = nw362_
            d_2304_v75_: _dafny.Seq
            d_2304_v75_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False]))])
            index346_ = default__.safeIndex(217, (d_2303_v74_).length(0))
            (d_2303_v74_)[index346_] = len((d_2304_v75_) + (d_2304_v75_))
            index347_ = default__.safeIndex(217, (d_2303_v74_).length(0))
            (d_2303_v74_)[index347_] = ((0) - (default__.fm1(not((self).f27), d_2299_v70_, (self).f27, globalState))) - (936)
            (globalState).f6 = default__.fm1((d_2302_v73_) not in (d_2210_v1_), d_2299_v70_, (d_2210_v1_) <= ((d_2210_v1_).set(default__.safeIndex((self).f28, len(d_2210_v1_)), d_2302_v73_)), globalState)
            (globalState).f12 = (self).f27
        hi12_ = (self).f29
        for d_2305_i6_ in range((self).f29, hi12_):
            d_2306_v76_: _dafny.Seq
            d_2306_v76_ = _dafny.SeqWithoutIsStrInference([(self).f27, True, (self).f27])
            d_2307_v77_: _dafny.Map
            d_2307_v77_ = _dafny.Map({d_2306_v76_: d_2210_v1_})
            d_2308_v78_: _dafny.Map
            d_2308_v78_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f29, 448, (self).f29, (self).f29, (self).f29]): d_2306_v76_})
            d_2307_v77_ = (d_2307_v77_).set(((d_2308_v78_)[_dafny.SeqWithoutIsStrInference([d_2305_i6_, d_2305_i6_])] if (_dafny.SeqWithoutIsStrInference([d_2305_i6_, d_2305_i6_])) in (d_2308_v78_) else _dafny.SeqWithoutIsStrInference([(self).f27])), d_2210_v1_)
            d_2309_v79_: _dafny.Set
            d_2309_v79_ = _dafny.Set({d_2305_i6_})
            d_2310_v80_: D18
            d_2310_v80_ = D18_DC41(d_2309_v79_)
            d_2311_v81_: _dafny.MultiSet
            d_2311_v81_ = _dafny.MultiSet([(self).f27, (default__.fm0(d_2210_v1_, True, (self).f27, (d_2310_v80_).cf56, globalState)) or ((self).f27), ((self).f27) == ((self).f27)])
            d_2312_v82_: _dafny.Array
            nw363_ = _dafny.Array(int(0), 3)
            d_2312_v82_ = nw363_
            index348_ = default__.safeIndex(68, (d_2312_v82_).length(0))
            (d_2312_v82_)[index348_] = default__.safeDivisionInt((self).f29, (0) - ((self).f28))
            d_2313_v83_: D21
            d_2313_v83_ = D21_DC54((self).f27, d_2305_i6_, d_2305_i6_)
            d_2314_v84_: D6
            d_2314_v84_ = D6_DC13(d_2306_v76_)
            index349_ = default__.safeIndex(68, (d_2312_v82_).length(0))
            rhs351_ = _dafny.MultiSet([(d_2313_v83_).cf79, (self).f27])
            rhs352_ = (self).f29
            rhs353_ = (0) - (len(((d_2306_v76_ if (self).f27 else d_2306_v76_)) + ((d_2314_v84_).cf21)))
            rhs354_ = (self).f27
            lhs282_ = d_2312_v82_
            lhs283_ = default__.safeIndex(68, (d_2312_v82_).length(0))
            lhs284_ = globalState
            lhs285_ = globalState
            d_2311_v81_ = rhs351_
            lhs282_[lhs283_] = rhs352_
            lhs284_.f6 = rhs353_
            lhs285_.f5 = rhs354_
            d_2315_v85_: _dafny.MultiSet
            d_2315_v85_ = _dafny.MultiSet([d_2210_v1_])
            d_2316_v86_: _dafny.Map
            d_2316_v86_ = _dafny.Map({(self).f27: not((self).f27)})
            d_2317_v87_: _dafny.Map
            d_2317_v87_ = _dafny.Map({_dafny.Map({721: (self).f27}): (self).f27})
            d_2318_v89_: _dafny.Array
            nw364_ = _dafny.Array(None, 8)
            nw364_[int(0)] = (d_2315_v85_).issubset(d_2315_v85_)
            nw364_[int(1)] = (self).f27
            def iife208_():
                coll92_ = _dafny.Map()
                compr_92_: int
                for compr_92_ in _dafny.IntegerRange(252, 568):
                    d_2319_v88_: int = compr_92_
                    if ((252) <= (d_2319_v88_)) and ((d_2319_v88_) < (568)):
                        coll92_[(d_2319_v88_) + ((0) - ((0) - ((self).f28)))] = (self).f27
                return _dafny.Map(coll92_)
            def iife209_():
                coll93_ = _dafny.Map()
                compr_93_: int
                for compr_93_ in _dafny.IntegerRange(252, 568):
                    d_2320_v88_: int = compr_93_
                    if ((252) <= (d_2320_v88_)) and ((d_2320_v88_) < (568)):
                        coll93_[(d_2320_v88_) + ((0) - ((0) - ((self).f28)))] = (self).f27
                return _dafny.Map(coll93_)
            nw364_[int(2)] = ((d_2316_v86_)[(self).f27] if ((self).f27) in (d_2316_v86_) else ((d_2317_v87_)[iife208_()
            ] if (iife209_()
            ) in (d_2317_v87_) else (self).f27))
            nw364_[int(3)] = (self).f27
            nw364_[int(4)] = (self).f27
            nw364_[int(5)] = (self).f27
            nw364_[int(6)] = (self).f27
            nw364_[int(7)] = (self).f27
            d_2318_v89_ = nw364_
            d_2321_v90_: _dafny.Map
            d_2321_v90_ = _dafny.Map({(self).f29: d_2318_v89_})
            d_2318_v89_ = ((d_2321_v90_)[d_2305_i6_] if (d_2305_i6_) in (d_2321_v90_) else d_2318_v89_)
            if (self).f27:
                index350_ = default__.safeIndex(68, (d_2312_v82_).length(0))
                (d_2312_v82_)[index350_] = d_2305_i6_
                d_2322_v91_: _dafny.Seq
                d_2322_v91_ = _dafny.SeqWithoutIsStrInference([len(d_2210_v1_), (0) - ((-374) - (d_2305_i6_)), len(d_2210_v1_), (d_2312_v82_)[default__.safeIndex(68, (d_2312_v82_).length(0))], default__.fm1(False, _dafny.SeqWithoutIsStrInference([True, (self).f27, (self).f27, (self).f27]), (self).f27, globalState)])
                d_2322_v91_ = d_2322_v91_
                d_2323_v92_: _dafny.Seq
                d_2323_v92_ = _dafny.SeqWithoutIsStrInference([d_2210_v1_, d_2210_v1_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_2324_i7_ in range(default__.abs(-903))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jya"))])
                d_2325_v93_: _dafny.MultiSet
                d_2325_v93_ = _dafny.MultiSet([d_2323_v92_])
                (globalState).f6 = ((d_2325_v93_)[d_2323_v92_] if (d_2323_v92_) in (d_2325_v93_) else (self).f29)
                (globalState).f12 = (self).f27
                (globalState).f6 = (d_2312_v82_)[default__.safeIndex(68, (d_2312_v82_).length(0))]
            elif True:
                d_2326_v94_: _dafny.Map
                d_2326_v94_ = _dafny.Map({d_2318_v89_: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_2327_i8_ in range(default__.abs(-545))])})
                d_2326_v94_ = (d_2326_v94_).set(d_2318_v89_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bojfwllht")))
                (globalState).f12 = ((self).f28) <= (d_2305_i6_)
                d_2209_v0_ = d_2209_v0_
                index351_ = default__.safeIndex(182, (d_2318_v89_).length(0))
                (d_2318_v89_)[index351_] = (self).f27
                index352_ = default__.safeIndex(182, (d_2318_v89_).length(0))
                rhs355_ = (0) - (d_2305_i6_)
                rhs356_ = (default__.fm34(globalState)) - (d_2311_v81_)
                rhs357_ = ((self).f28) > (58)
                rhs358_ = d_2210_v1_
                rhs359_ = (d_2312_v82_)[default__.safeIndex(68, (d_2312_v82_).length(0))]
                lhs286_ = d_2318_v89_
                lhs287_ = default__.safeIndex(182, (d_2318_v89_).length(0))
                lhs288_ = globalState
                r0 = rhs355_
                d_2311_v81_ = rhs356_
                lhs286_[lhs287_] = rhs357_
                d_2210_v1_ = rhs358_
                lhs288_.f2 = rhs359_
                (globalState).f5 = ((self).f29) > (d_2305_i6_)
        (globalState).f6 = (self).f28
        d_2328_v95_: str
        d_2328_v95_ = _dafny.CodePoint('h')
        d_2329_v96_: _dafny.Map
        d_2329_v96_ = _dafny.Map({(self).f28: d_2328_v95_})
        hi13_ = default__.safeModuloInt(len(d_2329_v96_), (self).f29)
        for d_2330_i9_ in range((self).f28, hi13_):
            if ((self).f29) >= (((self).f28) + (default__.fm1((self).f27, _dafny.SeqWithoutIsStrInference([(self).f27, (self).f27]), (self).f27, globalState))):
                (globalState).f6 = (self).f29
                d_2331_v97_: C2
                nw365_ = C2()
                nw365_.ctor__((self).f27, (self).f27)
                d_2331_v97_ = nw365_
                d_2332_v98_: _dafny.Seq
                d_2332_v98_ = _dafny.SeqWithoutIsStrInference([(self).f29])
                d_2333_v99_: _dafny.Array
                nw366_ = _dafny.Array(None, 5)
                nw366_[int(0)] = d_2330_i9_
                nw366_[int(1)] = (self).f29
                nw366_[int(2)] = (d_2332_v98_)[default__.safeIndex((self).f29, len(d_2332_v98_))]
                nw366_[int(3)] = -383
                nw366_[int(4)] = d_2330_i9_
                d_2333_v99_ = nw366_
                index353_ = default__.safeIndex(677, (d_2333_v99_).length(0))
                (d_2333_v99_)[index353_] = default__.safeDivisionInt(d_2330_i9_, (self).f29)
                index354_ = default__.safeIndex(677, (d_2333_v99_).length(0))
                (d_2333_v99_)[index354_] = d_2330_i9_
                d_2334_v100_: D0
                d_2334_v100_ = D0_DC2((self).f28, (self).f29)
                d_2335_v101_: _dafny.Seq
                d_2335_v101_ = _dafny.SeqWithoutIsStrInference([d_2334_v100_])
                index355_ = default__.safeIndex(677, (d_2333_v99_).length(0))
                index356_ = default__.safeIndex(677, (d_2333_v99_).length(0))
                rhs360_ = ((d_2333_v99_)[default__.safeIndex(677, (d_2333_v99_).length(0))]) + ((self).f29)
                rhs361_ = (self).f29
                rhs362_ = default__.fm0(d_2210_v1_, (d_2331_v97_).f34, (d_2331_v97_).f34, _dafny.Set({d_2330_i9_, len(d_2335_v101_), 930}), globalState)
                rhs363_ = d_2330_i9_
                rhs364_ = (d_2333_v99_)[default__.safeIndex(677, (d_2333_v99_).length(0))]
                lhs289_ = globalState
                lhs290_ = globalState
                lhs291_ = globalState
                lhs292_ = d_2333_v99_
                lhs293_ = default__.safeIndex(677, (d_2333_v99_).length(0))
                lhs294_ = d_2333_v99_
                lhs295_ = default__.safeIndex(677, (d_2333_v99_).length(0))
                lhs289_.f6 = rhs360_
                lhs290_.f6 = rhs361_
                lhs291_.f25 = rhs362_
                lhs292_[lhs293_] = rhs363_
                lhs294_[lhs295_] = rhs364_
                (globalState).f2 = (d_2333_v99_)[default__.safeIndex(677, (d_2333_v99_).length(0))]
            elif True:
                (globalState).f6 = (self).f28
                rhs365_ = (D17_DC39(d_2210_v1_)).cf54
                rhs366_ = (d_2330_i9_) != (d_2330_i9_)
                lhs296_ = globalState
                d_2210_v1_ = rhs365_
                lhs296_.f25 = rhs366_
                d_2336_v102_: _dafny.Array
                def lambda107_(d_2337_i10_):
                    return (d_2337_i10_) * ((D14_DC31((self).f28)).cf43)

                init59_ = lambda107_
                nw367_ = _dafny.Array(None, 2)
                for i0_59_ in range(nw367_.length(0)):
                    nw367_[i0_59_] = init59_(i0_59_)
                d_2336_v102_ = nw367_
                d_2338_v103_: _dafny.Seq
                d_2338_v103_ = _dafny.SeqWithoutIsStrInference([len(d_2210_v1_)])
                d_2339_v104_: _dafny.MultiSet
                d_2339_v104_ = _dafny.MultiSet([128])
                index357_ = default__.safeIndex(491, (d_2336_v102_).length(0))
                (d_2336_v102_)[index357_] = (d_2338_v103_)[default__.safeIndex(((d_2339_v104_)[-909] if (-909) in (d_2339_v104_) else (self).f29), len(d_2338_v103_))]
                d_2340_v105_: _dafny.Array
                def lambda108_(d_2341_i11_):
                    return D7_DC17((self).f29, (self).f29, (self).f27)

                init60_ = lambda108_
                nw368_ = _dafny.Array(None, 12)
                for i0_60_ in range(nw368_.length(0)):
                    nw368_[i0_60_] = init60_(i0_60_)
                d_2340_v105_ = nw368_
                d_2342_v106_: _dafny.Seq
                d_2342_v106_ = _dafny.SeqWithoutIsStrInference([(self).f27])
                d_2343_v107_: D7
                d_2343_v107_ = D7_DC17(default__.fm1(True, d_2342_v106_, (self).f27, globalState), d_2330_i9_, (self).f27)
                index358_ = default__.safeIndex(783, (d_2340_v105_).length(0))
                (d_2340_v105_)[index358_] = d_2343_v107_
                d_2344_v108_: _dafny.Map
                d_2344_v108_ = _dafny.Map({d_2342_v106_: (self).f29})
                d_2345_v109_: _dafny.Map
                d_2345_v109_ = _dafny.Map({(self).f27: (self).f29})
                d_2346_v110_: _dafny.Map
                d_2346_v110_ = _dafny.Map({(d_2339_v104_).intersection(d_2339_v104_): len(d_2345_v109_)})
                index359_ = default__.safeIndex(491, (d_2336_v102_).length(0))
                index360_ = default__.safeIndex(783, (d_2340_v105_).length(0))
                rhs367_ = d_2330_i9_
                rhs368_ = ((self).f29) - (len(d_2344_v108_))
                rhs369_ = d_2330_i9_
                rhs370_ = d_2343_v107_
                rhs371_ = ((d_2346_v110_)[d_2339_v104_] if (d_2339_v104_) in (d_2346_v110_) else (self).f28)
                lhs297_ = globalState
                lhs298_ = d_2336_v102_
                lhs299_ = default__.safeIndex(491, (d_2336_v102_).length(0))
                lhs300_ = globalState
                lhs301_ = d_2340_v105_
                lhs302_ = default__.safeIndex(783, (d_2340_v105_).length(0))
                lhs303_ = globalState
                lhs297_.f6 = rhs367_
                lhs298_[lhs299_] = rhs368_
                lhs300_.f2 = rhs369_
                lhs301_[lhs302_] = rhs370_
                lhs303_.f2 = rhs371_
                d_2347_v113_: _dafny.Array
                def lambda109_(d_2348_v1_, d_2349_i9_, d_2350_v106_):
                    def lambda110_(d_2351_i12_):
                        def iife210_():
                            def iife212_():
                                coll96_ = _dafny.Map()
                                compr_96_: _dafny.Seq
                                for compr_96_ in (_dafny.Map({d_2348_v1_: d_2349_i9_})).keys.Elements:
                                    d_2352_v111_: _dafny.Seq = compr_96_
                                    if (d_2352_v111_) in (_dafny.Map({d_2348_v1_: d_2349_i9_})):
                                        coll96_[d_2352_v111_] = d_2350_v106_
                                return _dafny.Map(coll96_)
                            coll94_ = _dafny.Set()
                            def iife211_():
                                coll95_ = _dafny.Map()
                                compr_95_: _dafny.Seq
                                for compr_95_ in (_dafny.Map({d_2348_v1_: d_2349_i9_})).keys.Elements:
                                    d_2352_v111_: _dafny.Seq = compr_95_
                                    if (d_2352_v111_) in (_dafny.Map({d_2348_v1_: d_2349_i9_})):
                                        coll95_[d_2352_v111_] = d_2350_v106_
                                return _dafny.Map(coll95_)
                            compr_94_: _dafny.Seq
                            for compr_94_ in (iife211_()
                            ).keys.Elements:
                                d_2353_v112_: _dafny.Seq = compr_94_
                                if (d_2353_v112_) in (iife212_()
                                ):
                                    coll94_ = coll94_.union(_dafny.Set([d_2353_v112_]))
                            return _dafny.Set(coll94_)
                        return iife210_()
                        

                    return lambda110_

                init61_ = lambda109_(d_2210_v1_, d_2330_i9_, d_2342_v106_)
                nw369_ = _dafny.Array(None, 19)
                for i0_61_ in range(nw369_.length(0)):
                    nw369_[i0_61_] = init61_(i0_61_)
                d_2347_v113_ = nw369_
                d_2354_v114_: _dafny.Set
                d_2354_v114_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "opk")), (self).fm3(globalState)})
                index361_ = default__.safeIndex(115, (d_2347_v113_).length(0))
                (d_2347_v113_)[index361_] = d_2354_v114_
                index362_ = default__.safeIndex(115, (d_2347_v113_).length(0))
                rhs372_ = (d_2342_v106_)[default__.safeIndex((d_2336_v102_)[default__.safeIndex(491, (d_2336_v102_).length(0))], len(d_2342_v106_))]
                rhs373_ = d_2354_v114_
                rhs374_ = (self).f27
                rhs375_ = (self).f27
                lhs304_ = globalState
                lhs305_ = d_2347_v113_
                lhs306_ = default__.safeIndex(115, (d_2347_v113_).length(0))
                lhs307_ = globalState
                lhs308_ = globalState
                lhs304_.f12 = rhs372_
                lhs305_[lhs306_] = rhs373_
                lhs307_.f21 = rhs374_
                lhs308_.f14 = rhs375_
                d_2210_v1_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xejhqra"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ibx")))).set(default__.safeIndex((self).f29, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xejhqra"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ibx"))))), _dafny.CodePoint('s'))
            if (self).f27:
                d_2355_v115_: _dafny.Array
                def lambda111_(d_2356_i13_):
                    return _dafny.SeqWithoutIsStrInference([(self).f28])

                init62_ = lambda111_
                nw370_ = _dafny.Array(None, 18)
                for i0_62_ in range(nw370_.length(0)):
                    nw370_[i0_62_] = init62_(i0_62_)
                d_2355_v115_ = nw370_
                d_2357_v116_: _dafny.Seq
                d_2357_v116_ = _dafny.SeqWithoutIsStrInference([(self).f28])
                index363_ = default__.safeIndex(854, (d_2355_v115_).length(0))
                (d_2355_v115_)[index363_] = d_2357_v116_
                index364_ = default__.safeIndex(854, (d_2355_v115_).length(0))
                (d_2355_v115_)[index364_] = ((d_2357_v116_) + (d_2357_v116_)) + (d_2357_v116_)
                (globalState).f2 = (self).f28
                d_2358_v117_: int
                d_2359_v118_: _dafny.Map
                out72_: int
                out73_: _dafny.Map
                out72_, out73_ = (self).m11(globalState)
                d_2358_v117_ = out72_
                d_2359_v118_ = out73_
                (globalState).f6 = (0) - ((d_2358_v117_) * (((self).f28) - (d_2358_v117_)))
                d_2360_v119_: D0
                d_2360_v119_ = D0_DC0((self).f29)
                d_2361_v120_: D0
                d_2361_v120_ = D0_DC3(d_2360_v119_)
                d_2362_v121_: bool
                d_2363_v122_: _dafny.Seq
                out74_: bool
                out75_: _dafny.Seq
                out74_, out75_ = default__.m0(d_2361_v120_, D0_DC1((self).f27), globalState)
                d_2362_v121_ = out74_
                d_2363_v122_ = out75_
            elif True:
                d_2364_v123_: _dafny.Array
                nw371_ = _dafny.Array(_dafny.Map({}), 4)
                d_2364_v123_ = nw371_
                d_2365_v125_: _dafny.Seq
                d_2365_v125_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qwoufuxcn")), _dafny.SeqWithoutIsStrInference([d_2328_v95_ for d_2366_i14_ in range(default__.abs(286))])])
                index365_ = default__.safeIndex(945, (d_2364_v123_).length(0))
                def iife213_():
                    coll97_ = _dafny.Map()
                    compr_97_: D13
                    for compr_97_ in (_dafny.Map({D13_DC27(d_2365_v125_): (self).f28})).keys.Elements:
                        d_2367_v124_: D13 = compr_97_
                        if (d_2367_v124_) in (_dafny.Map({D13_DC27(d_2365_v125_): (self).f28})):
                            coll97_[d_2367_v124_] = (self).f28
                    return _dafny.Map(coll97_)
                (d_2364_v123_)[index365_] = iife213_()
                
                d_2368_v126_: D13
                d_2368_v126_ = D13_DC27(d_2365_v125_)
                d_2369_v127_: _dafny.Map
                d_2369_v127_ = _dafny.Map({False: (self).f28})
                d_2370_v128_: _dafny.Map
                d_2370_v128_ = _dafny.Map({d_2368_v126_: len(d_2369_v127_)})
                index366_ = default__.safeIndex(945, (d_2364_v123_).length(0))
                (d_2364_v123_)[index366_] = (d_2370_v128_) | (d_2370_v128_)
                d_2371_v129_: _dafny.Array
                nw372_ = _dafny.Array(_dafny.Map({}), 18)
                d_2371_v129_ = nw372_
                d_2372_v130_: _dafny.Set
                d_2372_v130_ = _dafny.Set({(self).f27})
                d_2373_v131_: _dafny.Map
                d_2373_v131_ = _dafny.Map({d_2210_v1_: d_2372_v130_})
                index367_ = default__.safeIndex(348, (d_2371_v129_).length(0))
                (d_2371_v129_)[index367_] = d_2373_v131_
                d_2374_v132_: _dafny.Array
                nw373_ = _dafny.Array(int(0), 13)
                d_2374_v132_ = nw373_
                d_2375_v134_: _dafny.Map
                d_2375_v134_ = _dafny.Map({d_2328_v95_: d_2329_v96_})
                d_2376_v135_: _dafny.Map
                d_2376_v135_ = _dafny.Map({(self).fm2((0) - (d_2330_i9_), d_2375_v134_, globalState): (self).f28})
                d_2377_v136_: _dafny.Map
                def iife214_():
                    coll98_ = _dafny.Map()
                    compr_98_: _dafny.Seq
                    for compr_98_ in (d_2376_v135_).keys.Elements:
                        d_2378_v133_: _dafny.Seq = compr_98_
                        if (d_2378_v133_) in (d_2376_v135_):
                            coll98_[d_2378_v133_] = d_2372_v130_
                    return _dafny.Map(coll98_)
                d_2377_v136_ = _dafny.Map({d_2374_v132_: (iife214_()
                ).set(d_2210_v1_, _dafny.Set({(self).f27}))})
                d_2379_v137_: _dafny.Seq
                d_2379_v137_ = _dafny.SeqWithoutIsStrInference([d_2374_v132_])
                index368_ = default__.safeIndex(348, (d_2371_v129_).length(0))
                (d_2371_v129_)[index368_] = ((d_2377_v136_)[(d_2379_v137_)[default__.safeIndex((0) - ((self).f28), len(d_2379_v137_))]] if ((d_2379_v137_)[default__.safeIndex((0) - ((self).f28), len(d_2379_v137_))]) in (d_2377_v136_) else (d_2373_v131_).set(d_2210_v1_, d_2372_v130_))
                index369_ = default__.safeIndex(266, (d_2374_v132_).length(0))
                (d_2374_v132_)[index369_] = (self).f28
                index370_ = default__.safeIndex(266, (d_2374_v132_).length(0))
                (d_2374_v132_)[index370_] = default__.safeModuloInt((self).f28, d_2330_i9_)
                d_2380_v138_: _dafny.Seq
                d_2380_v138_ = _dafny.SeqWithoutIsStrInference([(self).f27, (self).f27])
                d_2381_v139_: _dafny.Array
                nw374_ = _dafny.Array(None, 11)
                nw374_[int(0)] = (self).f27
                nw374_[int(1)] = (self).f27
                nw374_[int(2)] = (self).f27
                nw374_[int(3)] = (self).f27
                nw374_[int(4)] = (self).f27
                nw374_[int(5)] = (self).f27
                nw374_[int(6)] = not ((self).f27) or (not((self).f27))
                nw374_[int(7)] = (d_2210_v1_) == (d_2210_v1_)
                nw374_[int(8)] = (d_2380_v138_)[default__.safeIndex((0) - ((self).f28), len(d_2380_v138_))]
                nw374_[int(9)] = (self).f27
                nw374_[int(10)] = (self).f27
                d_2381_v139_ = nw374_
                index371_ = default__.safeIndex(280, (d_2381_v139_).length(0))
                (d_2381_v139_)[index371_] = True
                d_2382_v140_: _dafny.Map
                d_2382_v140_ = _dafny.Map({(self).f28: (self).f28})
                d_2383_v141_: _dafny.Seq
                d_2383_v141_ = _dafny.SeqWithoutIsStrInference([(self).f29])
                d_2384_v142_: _dafny.Set
                d_2384_v142_ = _dafny.Set({len(d_2382_v140_), len(d_2383_v141_), (self).f28})
                index372_ = default__.safeIndex(280, (d_2381_v139_).length(0))
                (d_2381_v139_)[index372_] = default__.fm0((d_2210_v1_) + ((d_2365_v125_)[default__.safeIndex(len((self).fm2((self).f28, d_2375_v134_, globalState)), len(d_2365_v125_))]), True, default__.fm0(d_2210_v1_, (self).f27, (d_2380_v138_)[default__.safeIndex((self).f28, len(d_2380_v138_))], d_2384_v142_, globalState), default__.fm35(D13_DC28((d_2374_v132_)[default__.safeIndex(266, (d_2374_v132_).length(0))]), (self).f27, (self).f27, globalState), globalState)
                d_2328_v95_ = (d_2328_v95_ if (d_2381_v139_)[default__.safeIndex(280, (d_2381_v139_).length(0))] else d_2328_v95_)
            d_2385_v143_: D22
            d_2385_v143_ = D22_DC57((self).f27, d_2328_v95_, (self).f27)
            if (d_2385_v143_).cf86:
                d_2386_v144_: _dafny.Array
                nw375_ = _dafny.Array(False, 11)
                d_2386_v144_ = nw375_
                index373_ = default__.safeIndex(239, (d_2386_v144_).length(0))
                (d_2386_v144_)[index373_] = True
                d_2387_v145_: _dafny.Seq
                d_2387_v145_ = _dafny.SeqWithoutIsStrInference([d_2330_i9_])
                index374_ = default__.safeIndex(239, (d_2386_v144_).length(0))
                (d_2386_v144_)[index374_] = (not((len(_dafny.SeqWithoutIsStrInference([d_2328_v95_]))) in (d_2387_v145_))) == (not ((self).f27) or ((self).f27))
                d_2388_v146_: C1
                nw376_ = C1()
                nw376_.ctor__(False)
                d_2388_v146_ = nw376_
                d_2389_v147_: _dafny.MultiSet
                d_2389_v147_ = _dafny.MultiSet([(self).f27])
                d_2390_v148_: T0
                nw377_ = C4()
                nw377_.ctor__(d_2389_v147_, (self).f27)
                d_2390_v148_ = nw377_
                d_2391_v149_: _dafny.Set
                d_2391_v149_ = _dafny.Set({d_2330_i9_})
                d_2392_v150_: D20
                d_2392_v150_ = D20_DC49(_dafny.CodePoint('y'), (self).f29, 27, d_2390_v148_, (d_2391_v149_).intersection(d_2391_v149_))
                d_2393_v151_: _dafny.Map
                d_2393_v151_ = _dafny.Map({(d_2386_v144_)[default__.safeIndex(239, (d_2386_v144_).length(0))]: 224})
                index375_ = default__.safeIndex(239, (d_2386_v144_).length(0))
                rhs376_ = d_2392_v150_
                rhs377_ = default__.safeDivisionInt((self).f28, len(d_2393_v151_))
                rhs378_ = (self).f27
                rhs379_ = False
                lhs309_ = globalState
                lhs310_ = globalState
                lhs311_ = d_2386_v144_
                lhs312_ = default__.safeIndex(239, (d_2386_v144_).length(0))
                d_2392_v150_ = rhs376_
                lhs309_.f2 = rhs377_
                lhs310_.f21 = rhs378_
                lhs311_[lhs312_] = rhs379_
                d_2394_v152_: T1
                nw378_ = C5()
                nw378_.ctor__((self).f27, (self).f28, ((d_2393_v151_)[(d_2390_v148_).f27] if ((d_2390_v148_).f27) in (d_2393_v151_) else (self).f29), (d_2390_v148_).f27)
                d_2394_v152_ = nw378_
                d_2395_v153_: _dafny.Set
                d_2395_v153_ = _dafny.Set({d_2394_v152_})
                d_2396_v154_: C2
                nw379_ = C2()
                nw379_.ctor__((d_2395_v153_).isdisjoint(d_2395_v153_), (d_2328_v95_) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yauctxqe"))))
                d_2396_v154_ = nw379_
                index376_ = default__.safeIndex(239, (d_2386_v144_).length(0))
                rhs380_ = (self).f29
                rhs381_ = True
                lhs313_ = globalState
                lhs314_ = d_2386_v144_
                lhs315_ = default__.safeIndex(239, (d_2386_v144_).length(0))
                lhs313_.f6 = rhs380_
                lhs314_[lhs315_] = rhs381_
            elif True:
                d_2397_v155_: _dafny.Seq
                d_2397_v155_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f27])), d_2330_i9_])
                d_2398_v156_: _dafny.Map
                d_2398_v156_ = _dafny.Map({d_2328_v95_: (self).f27})
                d_2399_v157_: _dafny.Map
                d_2399_v157_ = _dafny.Map({d_2210_v1_: (d_2397_v155_) + (_dafny.SeqWithoutIsStrInference([len(d_2398_v156_), (self).f29]))})
                d_2400_v158_: _dafny.Map
                d_2400_v158_ = _dafny.Map({d_2328_v95_: (self).f28})
                d_2401_v159_: _dafny.Map
                d_2401_v159_ = _dafny.Map({D29_DC69(d_2400_v158_): (self).f28})
                d_2402_v160_: _dafny.Map
                d_2402_v160_ = _dafny.Map({d_2401_v159_: d_2397_v155_})
                def iife215_():
                    coll99_ = _dafny.Set()
                    compr_99_: int
                    for compr_99_ in _dafny.IntegerRange(86, 433):
                        d_2404_v161_: int = compr_99_
                        if ((86) <= (d_2404_v161_)) and ((d_2404_v161_) < (433)):
                            coll99_ = coll99_.union(_dafny.Set([(d_2404_v161_) + (len(_dafny.SeqWithoutIsStrInference([(self).f27, (self).f27])))]))
                    return _dafny.Set(coll99_)
                d_2399_v157_ = (d_2399_v157_).set((_dafny.SeqWithoutIsStrInference([d_2328_v95_ for d_2403_i15_ in range(default__.abs(186))])) + (d_2210_v1_), ((d_2397_v155_).set(default__.safeIndex((self).f28, len(d_2397_v155_)), (self).f29)) + (((d_2402_v160_)[d_2401_v159_] if (d_2401_v159_) in (d_2402_v160_) else _dafny.SeqWithoutIsStrInference([len(iife215_()
                ), (self).f29]))))
                d_2405_v162_: _dafny.Seq
                d_2405_v162_ = _dafny.SeqWithoutIsStrInference([(self).f27])
                (globalState).f2 = (0) - (len((d_2405_v162_) + (d_2405_v162_)))
                (globalState).f6 = (self).f28
                d_2406_v163_: _dafny.MultiSet
                d_2406_v163_ = _dafny.MultiSet([(self).f27])
                d_2407_v164_: C4
                nw380_ = C4()
                nw380_.ctor__(d_2406_v163_, (self).f27)
                d_2407_v164_ = nw380_
                d_2408_v165_: _dafny.Map
                d_2408_v165_ = _dafny.Map({True: (self).f28})
                d_2409_v166_: D18
                d_2409_v166_ = D18_DC42(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lxirqeov")), (self).f27)
                (globalState).f12 = ((0) - ((d_2330_i9_) + (len(d_2408_v165_)))) > ((len((d_2409_v166_).cf57)) * ((_dafny.MultiSet([d_2330_i9_])).cardinality))
            (globalState).f12 = (self).f27
        d_2410_v167_: _dafny.Seq
        d_2410_v167_ = _dafny.SeqWithoutIsStrInference([929])
        d_2411_v168_: _dafny.Map
        d_2411_v168_ = _dafny.Map({(self).f27: (self).f27})
        r0 = len((d_2410_v167_) + (((D2_DC6((self).f28, d_2410_v167_, (self).f27, _dafny.SeqWithoutIsStrInference([len(d_2411_v168_)]), (self).f27)).cf8) + (default__.fm33((self).f27, (self).f27, globalState))))
        return r0

    def m2(self, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: bool = False
        d_2412_v0_: _dafny.Array
        def lambda112_(d_2413_i1_):
            return not((self).f27)

        init63_ = lambda112_
        nw381_ = _dafny.Array(None, 6)
        for i0_63_ in range(nw381_.length(0)):
            nw381_[i0_63_] = init63_(i0_63_)
        d_2412_v0_ = nw381_
        d_2414_v1_: _dafny.MultiSet
        d_2414_v1_ = _dafny.MultiSet([d_2412_v0_, d_2412_v0_])
        hi14_ = default__.safeModuloInt((self).f28, (0) - ((self).f29))
        for d_2415_i0_ in range(((d_2414_v1_)[d_2412_v0_] if (d_2412_v0_) in (d_2414_v1_) else (self).f28), hi14_):
            d_2416_v2_: _dafny.Seq
            d_2416_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fxp"))
            (globalState).f6 = ((self).f29) * (len(d_2416_v2_))
            d_2417_v3_: C0
            nw382_ = C0()
            nw382_.ctor__()
            d_2417_v3_ = nw382_
            d_2418_v4_: _dafny.Array
            def lambda113_(d_2419_i2_):
                return _dafny.Map({(self).f27: 825})

            init64_ = lambda113_
            nw383_ = _dafny.Array(None, 22)
            for i0_64_ in range(nw383_.length(0)):
                nw383_[i0_64_] = init64_(i0_64_)
            d_2418_v4_ = nw383_
            d_2420_v5_: _dafny.Map
            d_2420_v5_ = _dafny.Map({(self).f27: (self).f29})
            index377_ = default__.safeIndex(511, (d_2418_v4_).length(0))
            (d_2418_v4_)[index377_] = d_2420_v5_
            index378_ = default__.safeIndex(511, (d_2418_v4_).length(0))
            (d_2418_v4_)[index378_] = (d_2420_v5_ if (self).f27 else (d_2420_v5_) | (d_2420_v5_))
            d_2421_v6_: D0
            d_2421_v6_ = D0_DC0(231)
            d_2422_v7_: D0
            d_2422_v7_ = D0_DC3(d_2421_v6_)
            d_2422_v7_ = D0_DC3((d_2422_v7_).cf4)
        (globalState).f14 = (self).f27
        d_2423_v8_: _dafny.Seq
        d_2423_v8_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).f29)])
        (globalState).f2 = default__.safeDivisionInt((self).f29, len(d_2423_v8_))
        d_2424_v9_: _dafny.Array
        nw384_ = _dafny.Array(_dafny.Seq({}), 22)
        d_2424_v9_ = nw384_
        guard_loop_9_: int
        for guard_loop_9_ in _dafny.IntegerRange(0, (d_2424_v9_).length(0)):
            d_2425_i3_: int = guard_loop_9_
            if (True) and (((0) <= (d_2425_i3_)) and ((d_2425_i3_) < ((d_2424_v9_).length(0)))):
                (d_2424_v9_)[(d_2425_i3_)] = (_dafny.SeqWithoutIsStrInference([not(not((self).f27))])) + (_dafny.SeqWithoutIsStrInference([(self).f27, (self).f27, (self).f27, (self).f27]))
        d_2426_v10_: _dafny.Seq
        d_2426_v10_ = _dafny.SeqWithoutIsStrInference([D0_DC0((self).f29)])
        d_2427_v11_: D21
        d_2427_v11_ = D21_DC53(not((self).f27), (self).f29, (self).f27)
        pat_let_tv47_ = d_2426_v10_
        pat_let_tv48_ = d_2426_v10_
        pat_let_tv49_ = d_2426_v10_
        pat_let_tv50_ = d_2426_v10_
        pat_let_tv51_ = d_2426_v10_
        pat_let_tv52_ = d_2426_v10_
        pat_let_tv53_ = d_2426_v10_
        def lambda114_(source37_):
            if source37_.is_DC53:
                d_2428___mcc_h0_ = source37_.cf76
                d_2429___mcc_h1_ = source37_.cf77
                d_2430___mcc_h2_ = source37_.cf78
                d_2431_cf78_ = d_2430___mcc_h2_
                d_2432_cf77_ = d_2429___mcc_h1_
                d_2433_cf76_ = d_2428___mcc_h0_
                return pat_let_tv47_
            elif source37_.is_DC54:
                d_2434___mcc_h3_ = source37_.cf79
                d_2435___mcc_h4_ = source37_.cf80
                d_2436___mcc_h5_ = source37_.cf81
                d_2437_cf81_ = d_2436___mcc_h5_
                d_2438_cf80_ = d_2435___mcc_h4_
                d_2439_cf79_ = d_2434___mcc_h3_
                return (pat_let_tv48_) + (pat_let_tv49_)
            elif source37_.is_DC55:
                d_2440___mcc_h6_ = source37_.cf82
                d_2441_cf82_ = d_2440___mcc_h6_
                return pat_let_tv50_
            elif True:
                d_2442___mcc_h7_ = source37_.cf75
                d_2443_cf75_ = d_2442___mcc_h7_
                if (self).f27:
                    return (pat_let_tv51_).set(default__.safeIndex((self).f28, len(pat_let_tv52_)), D0_DC0((self).f29))
                elif True:
                    return pat_let_tv53_

        d_2426_v10_ = lambda114_(d_2427_v11_)
        d_2412_v0_ = (D17_DC38(d_2412_v0_)).cf53
        d_2444_v12_: _dafny.Array
        nw385_ = _dafny.Array(_dafny.Array(None, 0), 19)
        d_2444_v12_ = nw385_
        r0 = d_2444_v12_
        d_2445_v13_: _dafny.Seq
        d_2445_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ui"))
        r1 = d_2445_v13_
        d_2446_v14_: _dafny.Array
        nw386_ = _dafny.Array(_dafny.Array(None, 0), 18)
        d_2446_v14_ = nw386_
        r2 = d_2446_v14_
        d_2447_v16_: _dafny.Seq
        d_2447_v16_ = _dafny.SeqWithoutIsStrInference([(self).f27, (self).f27, (self).f27, (self).f27, (self).f27])
        d_2448_v17_: _dafny.Map
        d_2448_v17_ = _dafny.Map({d_2447_v16_: len(d_2445_v13_)})
        d_2449_v18_: _dafny.MultiSet
        def iife216_():
            coll100_ = _dafny.Map()
            compr_100_: _dafny.Seq
            for compr_100_ in (d_2448_v17_).keys.Elements:
                d_2450_v15_: _dafny.Seq = compr_100_
                if (d_2450_v15_) in (d_2448_v17_):
                    coll100_[d_2450_v15_] = (self).f29
            return _dafny.Map(coll100_)
        d_2449_v18_ = _dafny.MultiSet([len(iife216_()
        )])
        d_2451_v19_: _dafny.Set
        d_2451_v19_ = _dafny.Set({(d_2449_v18_).cardinality, (self).f29})
        d_2452_v20_: _dafny.Array
        nw387_ = _dafny.Array(None, 10)
        nw387_[int(0)] = (self).f28
        nw387_[int(1)] = (0) - ((self).f28)
        nw387_[int(2)] = 200
        nw387_[int(3)] = (self).f28
        nw387_[int(4)] = 300
        nw387_[int(5)] = (self).f28
        nw387_[int(6)] = (self).f29
        nw387_[int(7)] = 752
        nw387_[int(8)] = (self).f28
        nw387_[int(9)] = (self).f28
        d_2452_v20_ = nw387_
        d_2453_v21_: _dafny.Seq
        d_2453_v21_ = _dafny.SeqWithoutIsStrInference([(D29_DC71((d_2423_v8_)[default__.safeIndex((self).f29, len(d_2423_v8_))], d_2451_v19_, (self).f27, (self).f27, d_2452_v20_)).cf117])
        r3 = ((self).f29) >= (len(d_2453_v21_))
        return r0, r1, r2, r3

    def m11(self, globalState):
        r0: int = int(0)
        r1: _dafny.Map = _dafny.Map({})
        (globalState).f6 = (self).f28
        d_2454_v0_: C0
        nw388_ = C0()
        nw388_.ctor__()
        d_2454_v0_ = nw388_
        d_2455_v1_: _dafny.Array
        def lambda115_(d_2456_i0_):
            return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h')])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_2457_i1_ in range(default__.abs(438))]))

        init65_ = lambda115_
        nw389_ = _dafny.Array(None, 8)
        for i0_65_ in range(nw389_.length(0)):
            nw389_[i0_65_] = init65_(i0_65_)
        d_2455_v1_ = nw389_
        d_2458_v2_: str
        d_2458_v2_ = _dafny.CodePoint('y')
        d_2459_v3_: _dafny.Seq
        d_2459_v3_ = _dafny.SeqWithoutIsStrInference([d_2458_v2_])
        d_2460_v4_: _dafny.Map
        d_2460_v4_ = _dafny.Map({(self).f29: d_2458_v2_})
        d_2461_v5_: _dafny.Map
        d_2461_v5_ = _dafny.Map({d_2458_v2_: d_2460_v4_})
        index379_ = default__.safeIndex(747, (d_2455_v1_).length(0))
        (d_2455_v1_)[index379_] = (d_2459_v3_) + ((self).fm2((self).f28, d_2461_v5_, globalState))
        d_2462_v6_: _dafny.Set
        d_2462_v6_ = _dafny.Set({(self).f28, (self).f29})
        d_2463_v7_: _dafny.Array
        nw390_ = _dafny.Array(False, 20)
        d_2463_v7_ = nw390_
        d_2464_v8_: _dafny.Map
        d_2464_v8_ = _dafny.Map({d_2462_v6_: d_2463_v7_})
        d_2465_v9_: _dafny.Map
        d_2465_v9_ = (d_2464_v8_) | (d_2464_v8_)
        d_2466_v10_: _dafny.Array
        def lambda116_(d_2467_i2_):
            return _dafny.SeqWithoutIsStrInference([(self).f28, (self).f28, (self).f28])

        init66_ = lambda116_
        nw391_ = _dafny.Array(None, 6)
        for i0_66_ in range(nw391_.length(0)):
            nw391_[i0_66_] = init66_(i0_66_)
        d_2466_v10_ = nw391_
        d_2468_v11_: _dafny.MultiSet
        d_2468_v11_ = _dafny.MultiSet([True, (self).f27])
        index380_ = default__.safeIndex(747, (d_2455_v1_).length(0))
        rhs382_ = (default__.fm63(410, (self).f28, globalState)).cf57
        rhs383_ = default__.fm0(d_2459_v3_, (self).f27, (self).f27, d_2462_v6_, globalState)
        rhs384_ = not((self).f27)
        rhs385_ = d_2464_v8_
        rhs386_ = (d_2466_v10_ if (d_2468_v11_).issubset(d_2468_v11_) else d_2466_v10_)
        lhs316_ = d_2455_v1_
        lhs317_ = default__.safeIndex(747, (d_2455_v1_).length(0))
        lhs318_ = globalState
        lhs319_ = globalState
        lhs316_[lhs317_] = rhs382_
        lhs318_.f25 = rhs383_
        lhs319_.f25 = rhs384_
        d_2465_v9_ = rhs385_
        d_2466_v10_ = rhs386_
        d_2469_v12_: _dafny.MultiSet
        d_2469_v12_ = _dafny.MultiSet([(self).f29, (self).f29])
        d_2470_v13_: _dafny.Map
        d_2470_v13_ = _dafny.Map({d_2469_v12_: 495})
        d_2471_v14_: _dafny.Map
        d_2471_v14_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([241]): len(d_2470_v13_)})
        d_2472_v15_: D3
        d_2472_v15_ = D3_DC8(d_2471_v14_)
        d_2473_v16_: _dafny.Seq
        d_2473_v16_ = _dafny.SeqWithoutIsStrInference([(self).f28])
        d_2474_v17_: _dafny.Map
        d_2474_v17_ = _dafny.Map({d_2472_v15_: _dafny.MultiSet(d_2473_v16_)})
        d_2475_v18_: _dafny.Seq
        d_2475_v18_ = _dafny.SeqWithoutIsStrInference([d_2474_v17_])
        d_2476_v19_: D20
        d_2476_v19_ = D20_DC48((d_2475_v18_)[default__.safeIndex((self).f29, len(d_2475_v18_))])
        d_2477_v20_: _dafny.Seq
        d_2477_v20_ = _dafny.SeqWithoutIsStrInference([d_2476_v19_, d_2476_v19_])
        d_2478_i3_: int
        d_2478_i3_ = 0
        with _dafny.label("14"):
            while not((len((d_2477_v20_) + (d_2477_v20_))) == (len(((d_2455_v1_)[default__.safeIndex(747, (d_2455_v1_).length(0))]) + ((d_2455_v1_)[default__.safeIndex(747, (d_2455_v1_).length(0))])))):
                with _dafny.c_label("14"):
                    if (d_2478_i3_) >= (100):
                        raise _dafny.Break("14")
                    d_2478_i3_ = (d_2478_i3_) + (1)
                    (globalState).f25 = (self).f27
                    d_2479_v21_: _dafny.Array
                    nw392_ = _dafny.Array(int(0), 12)
                    d_2479_v21_ = nw392_
                    d_2479_v21_ = d_2479_v21_
                    d_2480_v22_: D18
                    d_2480_v22_ = D18_DC42(d_2459_v3_, True)
                    d_2481_v23_: _dafny.Map
                    d_2481_v23_ = _dafny.Map({False: ((d_2480_v22_).cf58) and ((self).f27)})
                    if ((d_2481_v23_)[(self).f27] if ((self).f27) in (d_2481_v23_) else ((self).f29) == ((self).f29)):
                        (globalState).f2 = (self).f29
                        d_2482_v24_: _dafny.Seq
                        d_2482_v24_ = _dafny.SeqWithoutIsStrInference([(self).f27])
                        d_2483_v25_: _dafny.Seq
                        d_2483_v25_ = _dafny.SeqWithoutIsStrInference([(d_2482_v24_ if default__.fm0(_dafny.SeqWithoutIsStrInference([d_2458_v2_ for d_2484_i4_ in range(default__.abs(945))]), False, (self).f27, d_2462_v6_, globalState) else d_2482_v24_), (d_2482_v24_) + (d_2482_v24_), d_2482_v24_])
                        d_2483_v25_ = (d_2483_v25_) + (d_2483_v25_)
                        (globalState).f6 = (self).f29
                        (globalState).f21 = ((self).f27) and (not((self).f27))
                        d_2485_v26_: C2
                        nw393_ = C2()
                        nw393_.ctor__(True, True)
                        d_2485_v26_ = nw393_
                    elif True:
                        d_2486_v27_: _dafny.Array
                        nw394_ = _dafny.Array(_dafny.Array(None, 0), 14)
                        d_2486_v27_ = nw394_
                        index381_ = default__.safeIndex(607, (d_2486_v27_).length(0))
                        (d_2486_v27_)[index381_] = d_2479_v21_
                        index382_ = default__.safeIndex(607, (d_2486_v27_).length(0))
                        rhs387_ = d_2479_v21_
                        rhs388_ = _dafny.SeqWithoutIsStrInference([d_2458_v2_ for d_2487_i5_ in range(default__.abs(477))])
                        lhs320_ = d_2486_v27_
                        lhs321_ = default__.safeIndex(607, (d_2486_v27_).length(0))
                        lhs320_[lhs321_] = rhs387_
                        d_2459_v3_ = rhs388_
                        d_2488_v28_: _dafny.Seq
                        d_2488_v28_ = _dafny.SeqWithoutIsStrInference([(self).f27, (self).f27, (self).f27])
                        d_2489_v29_: _dafny.Map
                        d_2489_v29_ = _dafny.Map({(self).f28: (d_2488_v28_).set(default__.safeIndex((self).f28, len(d_2488_v28_)), (self).f27)})
                        d_2490_v30_: _dafny.Map
                        d_2490_v30_ = _dafny.Map({(self).f27: (d_2488_v28_).set(default__.safeIndex((self).f28, len(d_2488_v28_)), (self).f27)})
                        d_2491_v31_: D22
                        d_2491_v31_ = D22_DC57((self).f27, d_2458_v2_, (self).f27)
                        d_2492_v32_: D18
                        d_2492_v32_ = D18_DC43((self).f29)
                        d_2493_v33_: _dafny.Map
                        d_2493_v33_ = _dafny.Map({d_2492_v32_: d_2488_v28_})
                        d_2494_v34_: _dafny.Array
                        nw395_ = _dafny.Array(None, 19)
                        nw395_[int(0)] = d_2488_v28_
                        nw395_[int(1)] = (d_2488_v28_) + (d_2488_v28_)
                        nw395_[int(2)] = d_2488_v28_
                        nw395_[int(3)] = default__.fm32((self).f27, _dafny.SeqWithoutIsStrInference([(self).f29, len(d_2481_v23_)]), (self).f27, globalState)
                        nw395_[int(4)] = d_2488_v28_
                        nw395_[int(5)] = d_2488_v28_
                        nw395_[int(6)] = d_2488_v28_
                        nw395_[int(7)] = ((d_2489_v29_)[515] if (515) in (d_2489_v29_) else d_2488_v28_)
                        nw395_[int(8)] = d_2488_v28_
                        nw395_[int(9)] = (((d_2490_v30_)[(self).f27] if ((self).f27) in (d_2490_v30_) else _dafny.SeqWithoutIsStrInference([(self).f27, (self).f27])) if (self).f27 else _dafny.SeqWithoutIsStrInference([(d_2491_v31_).cf84, (self).f27]))
                        nw395_[int(10)] = _dafny.SeqWithoutIsStrInference([(self).f27, (self).f27])
                        nw395_[int(11)] = ((d_2493_v33_)[d_2492_v32_] if (d_2492_v32_) in (d_2493_v33_) else d_2488_v28_)
                        nw395_[int(12)] = d_2488_v28_
                        nw395_[int(13)] = d_2488_v28_
                        nw395_[int(14)] = ((d_2488_v28_) + (d_2488_v28_)).set(default__.safeIndex((self).f28, len((d_2488_v28_) + (d_2488_v28_))), (self).f27)
                        nw395_[int(15)] = d_2488_v28_
                        nw395_[int(16)] = (d_2488_v28_) + (d_2488_v28_)
                        nw395_[int(17)] = d_2488_v28_
                        nw395_[int(18)] = d_2488_v28_
                        d_2494_v34_ = nw395_
                        d_2494_v34_ = d_2494_v34_
                        (globalState).f14 = False
                        d_2495_v35_: _dafny.Map
                        d_2495_v35_ = _dafny.Map({((self).f28) not in (_dafny.Set({(self).f28, len(d_2459_v3_), (self).f29})): d_2458_v2_})
                        d_2495_v35_ = (d_2495_v35_).set((self).f27, d_2458_v2_)
                        (globalState).f25 = (self).f27
                    d_2496_v36_: _dafny.Map
                    d_2496_v36_ = _dafny.Map({d_2458_v2_: not(((self).f27 if (self).f27 else (self).f27))})
                    d_2496_v36_ = (d_2496_v36_).set(d_2458_v2_, (self).f27)
                    pass
            pass
        if not((self).f27):
            if (self).f27:
                d_2497_v37_: _dafny.Map
                d_2497_v37_ = _dafny.Map({d_2473_v16_: (self).f27})
                d_2497_v37_ = (d_2497_v37_).set((d_2473_v16_) + (d_2473_v16_), (self).f27)
                d_2497_v37_ = (d_2497_v37_).set((d_2473_v16_) + ((d_2473_v16_).set(default__.safeIndex((self).f29, len(d_2473_v16_)), (self).f29)), (self).f27)
                d_2498_v38_: _dafny.Seq
                d_2498_v38_ = _dafny.SeqWithoutIsStrInference([(self).f27])
                d_2499_v39_: _dafny.Map
                d_2499_v39_ = _dafny.Map({d_2498_v38_: (self).f27})
                d_2500_v40_: _dafny.Array
                nw396_ = _dafny.Array(None, 17)
                nw396_[int(0)] = len(default__.fm52((0) - (default__.fm1((self).f27, d_2498_v38_, (self).f27, globalState)), (self).f27, d_2468_v11_, globalState))
                nw396_[int(1)] = (self).f29
                nw396_[int(2)] = (self).f29
                nw396_[int(3)] = ((0) - ((self).f29)) + ((self).f28)
                nw396_[int(4)] = (self).f29
                nw396_[int(5)] = (self).f28
                nw396_[int(6)] = len(d_2499_v39_)
                nw396_[int(7)] = 217
                nw396_[int(8)] = (self).f28
                nw396_[int(9)] = (self).f29
                nw396_[int(10)] = (self).f29
                nw396_[int(11)] = (self).f28
                nw396_[int(12)] = (self).f29
                nw396_[int(13)] = ((self).f28) + ((self).f29)
                nw396_[int(14)] = (self).f29
                nw396_[int(15)] = 870
                nw396_[int(16)] = 413
                d_2500_v40_ = nw396_
                index383_ = default__.safeIndex(242, (d_2500_v40_).length(0))
                (d_2500_v40_)[index383_] = (self).f29
                index384_ = default__.safeIndex(242, (d_2500_v40_).length(0))
                (d_2500_v40_)[index384_] = (0) - (((self).f29) * ((self).f28))
                d_2501_v41_: _dafny.Array
                nw397_ = _dafny.Array(None, 19)
                nw397_[int(0)] = d_2465_v9_
                nw397_[int(1)] = d_2464_v8_
                nw397_[int(2)] = d_2465_v9_
                nw397_[int(3)] = d_2465_v9_
                nw397_[int(4)] = _dafny.Map({_dafny.Set({(d_2500_v40_)[default__.safeIndex(242, (d_2500_v40_).length(0))]}): d_2463_v7_})
                nw397_[int(5)] = _dafny.Map({d_2462_v6_: d_2463_v7_})
                nw397_[int(6)] = d_2465_v9_
                nw397_[int(7)] = d_2464_v8_
                nw397_[int(8)] = d_2465_v9_
                nw397_[int(9)] = d_2465_v9_
                nw397_[int(10)] = d_2465_v9_
                nw397_[int(11)] = d_2465_v9_
                nw397_[int(12)] = d_2465_v9_
                nw397_[int(13)] = d_2465_v9_
                nw397_[int(14)] = d_2465_v9_
                nw397_[int(15)] = d_2465_v9_
                nw397_[int(16)] = d_2465_v9_
                nw397_[int(17)] = d_2465_v9_
                nw397_[int(18)] = d_2465_v9_
                d_2501_v41_ = nw397_
                index385_ = default__.safeIndex(86, (d_2501_v41_).length(0))
                (d_2501_v41_)[index385_] = d_2464_v8_
                index386_ = default__.safeIndex(86, (d_2501_v41_).length(0))
                (d_2501_v41_)[index386_] = d_2465_v9_
                (globalState).f2 = (self).f28
            elif True:
                d_2502_v42_: D0
                d_2502_v42_ = D0_DC1((self).f27)
                d_2503_v43_: D0
                d_2503_v43_ = D0_DC3(d_2502_v42_)
                d_2504_v44_: D0
                d_2504_v44_ = D0_DC3(d_2502_v42_)
                d_2505_v45_: D0
                d_2505_v45_ = D0_DC3(d_2503_v43_)
                d_2506_v46_: _dafny.Seq
                d_2506_v46_ = _dafny.SeqWithoutIsStrInference([(self).f27])
                d_2507_v47_: D0
                d_2507_v47_ = D0_DC1((d_2506_v46_)[default__.safeIndex((self).f28, len(d_2506_v46_))])
                d_2508_v48_: bool
                d_2509_v49_: _dafny.Seq
                out76_: bool
                out77_: _dafny.Seq
                out76_, out77_ = default__.m0(d_2505_v45_, d_2507_v47_, globalState)
                d_2508_v48_ = out76_
                d_2509_v49_ = out77_
                d_2469_v12_ = (_dafny.MultiSet([(self).f28])) | (d_2469_v12_)
                index387_ = default__.safeIndex(747, (d_2455_v1_).length(0))
                (d_2455_v1_)[index387_] = (d_2459_v3_) + (((d_2455_v1_)[default__.safeIndex(747, (d_2455_v1_).length(0))]) + ((d_2455_v1_)[default__.safeIndex(747, (d_2455_v1_).length(0))]))
                d_2508_v48_ = (self).f27
                (globalState).f6 = (self).f29
            (globalState).f6 = default__.safeDivisionInt((self).f29, (self).f28)
            (globalState).f25 = ((self).f27 if (self).f27 else (self).f27)
            (globalState).f5 = (self).f27
            (globalState).f6 = (self).f28
        elif True:
            d_2510_v50_: _dafny.Map
            d_2510_v50_ = _dafny.Map({(self).f27: (self).f29})
            d_2511_v51_: _dafny.Map
            d_2511_v51_ = _dafny.Map({d_2454_v0_: d_2510_v50_})
            d_2512_v52_: _dafny.Map
            def iife217_(_pat_let58_0):
                def iife218_(d_2513_dt__update__tmp_h3_):
                    def iife219_(_pat_let59_0):
                        def iife220_(d_2514_dt__update_hcf31_h0_):
                            return D9_DC21(d_2514_dt__update_hcf31_h0_, (d_2513_dt__update__tmp_h3_).cf32)
                        return iife220_(_pat_let59_0)
                    return iife219_((self).f27)
                return iife218_(_pat_let58_0)
            d_2512_v52_ = _dafny.Map({iife217_(D9_DC21((self).f27, (self).f28)): ((d_2511_v51_)[d_2454_v0_] if (d_2454_v0_) in (d_2511_v51_) else d_2510_v50_)})
            d_2515_v53_: D9
            d_2515_v53_ = D9_DC21((self).f27, (self).f29)
            d_2512_v52_ = (d_2512_v52_).set(d_2515_v53_, d_2510_v50_)
            (globalState).f6 = default__.safeModuloInt((self).f28, len(d_2459_v3_))
            d_2516_v54_: D15
            d_2516_v54_ = D15_DC34(((d_2468_v11_)[(self).f27] if ((self).f27) in (d_2468_v11_) else (self).f28), (self).f27, (self).f28, (self).f29, False)
            d_2517_v55_: C5
            nw398_ = C5()
            nw398_.ctor__((self).f27, ((self).f29) - (317), ((self).f29) + ((d_2516_v54_).cf45), not((self).f27))
            d_2517_v55_ = nw398_
            d_2518_v56_: _dafny.Array
            nw399_ = _dafny.Array(None, 19)
            nw399_[int(0)] = d_2458_v2_
            nw399_[int(1)] = d_2458_v2_
            nw399_[int(2)] = d_2458_v2_
            nw399_[int(3)] = d_2458_v2_
            nw399_[int(4)] = d_2458_v2_
            nw399_[int(5)] = d_2458_v2_
            nw399_[int(6)] = d_2458_v2_
            nw399_[int(7)] = d_2458_v2_
            nw399_[int(8)] = ((d_2455_v1_)[default__.safeIndex(747, (d_2455_v1_).length(0))])[default__.safeIndex(len(_dafny.Set({(d_2517_v55_).f39})), len((d_2455_v1_)[default__.safeIndex(747, (d_2455_v1_).length(0))]))]
            nw399_[int(9)] = d_2458_v2_
            nw399_[int(10)] = d_2458_v2_
            nw399_[int(11)] = _dafny.CodePoint('e')
            nw399_[int(12)] = _dafny.CodePoint('x')
            nw399_[int(13)] = d_2458_v2_
            nw399_[int(14)] = d_2458_v2_
            nw399_[int(15)] = _dafny.CodePoint('c')
            nw399_[int(16)] = default__.fm54((d_2517_v55_).f39, default__.fm64((d_2517_v55_).f39, (self).f28, globalState), (0) - ((self).f28), globalState)
            nw399_[int(17)] = d_2458_v2_
            nw399_[int(18)] = d_2458_v2_
            d_2518_v56_ = nw399_
            d_2519_v57_: _dafny.Seq
            d_2519_v57_ = _dafny.SeqWithoutIsStrInference([(self).f27, (d_2517_v55_).f39])
            index388_ = default__.safeIndex(904, (d_2518_v56_).length(0))
            (d_2518_v56_)[index388_] = default__.fm29(default__.fm1(True, d_2519_v57_, (self).f27, globalState), (self).f27, globalState)
            index389_ = default__.safeIndex(904, (d_2518_v56_).length(0))
            rhs389_ = (d_2458_v2_ if (d_2517_v55_).f39 else d_2458_v2_)
            rhs390_ = d_2458_v2_
            rhs391_ = (self).f27
            lhs322_ = d_2518_v56_
            lhs323_ = default__.safeIndex(904, (d_2518_v56_).length(0))
            lhs324_ = globalState
            d_2458_v2_ = rhs389_
            lhs322_[lhs323_] = rhs390_
            lhs324_.f14 = rhs391_
            d_2520_v58_: D22
            d_2520_v58_ = D22_DC58((self).f27, (default__.fm1((d_2517_v55_).f39, d_2519_v57_, (self).f27, globalState) if (self).f27 else (self).f28), (0) - ((self).f28))
            d_2521_v59_: _dafny.Set
            d_2521_v59_ = _dafny.Set({(d_2455_v1_)[default__.safeIndex(747, (d_2455_v1_).length(0))], (d_2455_v1_)[default__.safeIndex(747, (d_2455_v1_).length(0))]})
            rhs392_ = d_2520_v58_
            rhs393_ = True
            rhs394_ = (d_2521_v59_) != ((d_2521_v59_) | (d_2521_v59_))
            lhs325_ = globalState
            lhs326_ = globalState
            d_2520_v58_ = rhs392_
            lhs325_.f5 = rhs393_
            lhs326_.f14 = rhs394_
        (globalState).f2 = len((d_2455_v1_)[default__.safeIndex(747, (d_2455_v1_).length(0))])
        d_2522_v61_: D15
        def iife221_():
            coll101_ = _dafny.Map()
            compr_101_: int
            for compr_101_ in _dafny.IntegerRange(340, 624):
                d_2523_v60_: int = compr_101_
                if ((340) <= (d_2523_v60_)) and ((d_2523_v60_) < (624)):
                    coll101_[(d_2523_v60_) + (len((d_2455_v1_)[default__.safeIndex(747, (d_2455_v1_).length(0))]))] = (self).f28
            return _dafny.Map(coll101_)
        d_2522_v61_ = D15_DC34((self).f28, (self).f27, len(iife221_()
), (self).f28, (self).f27)
        r0 = (0) - (((d_2522_v61_).cf45) * ((self).f28))
        d_2524_v62_: _dafny.Map
        d_2524_v62_ = _dafny.Map({(self).f27: (self).f27})
        d_2525_v63_: _dafny.Map
        d_2525_v63_ = _dafny.Map({D21_DC55((0) - (len(d_2524_v62_))): (0) - (((self).f29) + (len(default__.fm28(globalState))))})
        r1 = d_2525_v63_
        return r0, r1


class C7(T1, T0):
    def  __init__(self):
        self._f27: bool = False
        self._f28: int = int(0)
        self._f29: int = int(0)
        self.f35: bool = False
        self._f36: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    @property
    def f27(self):
        return self._f27
    @property
    def f28(self):
        return self._f28
    @property
    def f29(self):
        return self._f29
    def ctor__(self, f35, f36, f28, f29, f27):
        (self).f35 = f35
        (self)._f36 = f36
        (self)._f28 = f28
        (self)._f29 = f29
        (self)._f27 = f27

    def fm2(self, p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "djaohq"))) + ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_2526_i0_ in range(default__.abs(642))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))))

    def fm3(self, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_2527_i0_ in range(default__.abs(448))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_2528_i1_ in range(default__.abs(546))]))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "leypqst"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dqwrngp"))))

    def m3(self, p0, p1, p2, globalState):
        hi15_ = (0) - ((self).f28)
        for d_2529_i0_ in range(p2, hi15_):
            d_2530_v0_: str
            d_2530_v0_ = _dafny.CodePoint('j')
            index390_ = default__.safeIndex(900, (p1).length(0))
            (p1)[index390_] = d_2530_v0_
            index391_ = default__.safeIndex(900, (p1).length(0))
            (p1)[index391_] = d_2530_v0_
            d_2531_v1_: _dafny.Set
            d_2531_v1_ = _dafny.Set({p0, self.f35})
            d_2532_v2_: _dafny.Map
            d_2532_v2_ = _dafny.Map({p2: _dafny.Set({self.f35})})
            (globalState).f2 = (((self).f36) + (d_2529_i0_)) * (len((d_2531_v1_) - (((d_2532_v2_)[51] if (51) in (d_2532_v2_) else d_2531_v1_))))
            d_2533_v3_: C1
            nw400_ = C1()
            nw400_.ctor__(self.f35)
            d_2533_v3_ = nw400_
            d_2534_v4_: _dafny.Seq
            d_2534_v4_ = _dafny.SeqWithoutIsStrInference([(self).f27, (d_2529_i0_) != ((self).f29), p0, self.f35])
            d_2534_v4_ = d_2534_v4_
        (globalState).f25 = p0
        hi16_ = ((self).f36) * ((self).f29)
        for d_2535_i1_ in range((self).f36, hi16_):
            d_2536_v5_: _dafny.Seq
            d_2536_v5_ = _dafny.SeqWithoutIsStrInference([self.f35])
            d_2537_v6_: C1
            nw401_ = C1()
            nw401_.ctor__(((self).f28) <= (len(d_2536_v5_)))
            d_2537_v6_ = nw401_
            d_2538_v7_: D0
            d_2538_v7_ = D0_DC1(not(p0))
            (globalState).f2 = (default__.fm1((d_2537_v6_).fm4(d_2538_v7_, globalState), _dafny.SeqWithoutIsStrInference([self.f35, not(p0)]), self.f35, globalState)) + (d_2535_i1_)
            d_2539_v8_: _dafny.Seq
            d_2539_v8_ = _dafny.SeqWithoutIsStrInference([(self).f28, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aa")))])
            d_2540_v9_: _dafny.Seq
            d_2540_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsxgyd"))
            d_2541_v10_: _dafny.Map
            d_2541_v10_ = _dafny.Map({len(d_2540_v9_): d_2539_v8_})
            d_2542_v11_: _dafny.Set
            d_2542_v11_ = _dafny.Set({False, True})
            d_2543_v12_: _dafny.MultiSet
            d_2543_v12_ = _dafny.MultiSet([p2])
            d_2544_v13_: D17
            d_2544_v13_ = D17_DC39(d_2540_v9_)
            d_2545_v14_: str
            d_2545_v14_ = _dafny.CodePoint('c')
            d_2546_v15_: _dafny.Map
            d_2546_v15_ = _dafny.Map({(self).f27: d_2545_v14_})
            d_2547_v16_: _dafny.Map
            d_2547_v16_ = _dafny.Map({len(_dafny.Map({(self).f28: (d_2546_v15_).set((self).f27, default__.fm24(self.f35, globalState))})): p0})
            d_2548_v17_: _dafny.Array
            nw402_ = _dafny.Array(None, 18)
            nw402_[int(0)] = (d_2539_v8_) + (d_2539_v8_)
            nw402_[int(1)] = (d_2539_v8_) + (d_2539_v8_)
            nw402_[int(2)] = d_2539_v8_
            nw402_[int(3)] = ((d_2541_v10_)[(self).f36] if ((self).f36) in (d_2541_v10_) else d_2539_v8_)
            nw402_[int(4)] = d_2539_v8_
            nw402_[int(5)] = d_2539_v8_
            nw402_[int(6)] = d_2539_v8_
            nw402_[int(7)] = (d_2539_v8_) + (_dafny.SeqWithoutIsStrInference([len(d_2536_v5_), len(d_2536_v5_), len(d_2542_v11_), len(d_2540_v9_), -587]))
            nw402_[int(8)] = d_2539_v8_
            nw402_[int(9)] = d_2539_v8_
            nw402_[int(10)] = _dafny.SeqWithoutIsStrInference([p2 for d_2549_i2_ in range(default__.abs(719))])
            nw402_[int(11)] = d_2539_v8_
            nw402_[int(12)] = (d_2539_v8_).set(default__.safeIndex(default__.fm1((self).f27, d_2536_v5_, self.f35, globalState), len(d_2539_v8_)), (self).f28)
            nw402_[int(13)] = _dafny.SeqWithoutIsStrInference([(self).f29, (0) - (p2), p2, (self).f29, p2])
            nw402_[int(14)] = _dafny.SeqWithoutIsStrInference([(0) - (d_2535_i1_), (self).f28])
            nw402_[int(15)] = _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({(d_2543_v12_).cardinality: (d_2544_v13_).cf54}))), p2, (0) - (d_2535_i1_), (self).f29])
            nw402_[int(16)] = (d_2539_v8_) + ((d_2539_v8_).set(default__.safeIndex((self).f29, len(d_2539_v8_)), len(_dafny.SeqWithoutIsStrInference([self.f35]))))
            nw402_[int(17)] = (d_2539_v8_).set(default__.safeIndex(len(d_2547_v16_), len(d_2539_v8_)), d_2535_i1_)
            d_2548_v17_ = nw402_
            index392_ = default__.safeIndex(644, (d_2548_v17_).length(0))
            (d_2548_v17_)[index392_] = d_2539_v8_
            d_2550_v18_: _dafny.MultiSet
            d_2550_v18_ = _dafny.MultiSet([d_2539_v8_, d_2539_v8_])
            index393_ = default__.safeIndex(644, (d_2548_v17_).length(0))
            (d_2548_v17_)[index393_] = _dafny.SeqWithoutIsStrInference([(self).f28, (self).f28, ((d_2550_v18_)[d_2539_v8_] if (d_2539_v8_) in (d_2550_v18_) else 506), (self).f29, 656])
            d_2551_v19_: _dafny.Set
            d_2551_v19_ = _dafny.Set({(self).f36, -373})
            d_2552_v20_: _dafny.Seq
            d_2552_v20_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iva"))])
            d_2553_v21_: D18
            d_2553_v21_ = D18_DC42((d_2552_v20_)[default__.safeIndex(len(d_2540_v9_), len(d_2552_v20_))], self.f35)
            pat_let_tv54_ = d_2540_v9_
            def iife222_(_pat_let60_0):
                def iife223_(d_2554_dt__update__tmp_h0_):
                    def iife224_(_pat_let61_0):
                        def iife225_(d_2555_dt__update_hcf57_h0_):
                            return D18_DC42(d_2555_dt__update_hcf57_h0_, (d_2554_dt__update__tmp_h0_).cf58)
                        return iife225_(_pat_let61_0)
                    return iife224_(pat_let_tv54_)
                return iife223_(_pat_let60_0)
            if ((d_2551_v19_ if not((self).f27) else default__.fm25(p2, (self).f29, (self).f27, (iife222_(d_2553_v21_)).cf58, globalState))).ispropersubset(d_2551_v19_):
                (globalState).f25 = True
                (globalState).f2 = p2
                d_2556_v22_: _dafny.Array
                nw403_ = _dafny.Array(False, 20)
                d_2556_v22_ = nw403_
                d_2557_v23_: _dafny.Map
                d_2557_v23_ = _dafny.Map({(self).f29: d_2556_v22_})
                d_2558_v24_: _dafny.Seq
                d_2558_v24_ = _dafny.SeqWithoutIsStrInference([((d_2557_v23_)[p2] if (p2) in (d_2557_v23_) else d_2556_v22_), d_2556_v22_, d_2556_v22_])
                d_2559_v25_: _dafny.Map
                d_2559_v25_ = _dafny.Map({p0: d_2556_v22_})
                d_2558_v24_ = _dafny.SeqWithoutIsStrInference([d_2556_v22_, d_2556_v22_, d_2556_v22_, d_2556_v22_, (((d_2559_v25_)[p0] if (p0) in (d_2559_v25_) else d_2556_v22_) if (d_2536_v5_)[default__.safeIndex(d_2535_i1_, len(d_2536_v5_))] else d_2556_v22_)])
                d_2560_v26_: _dafny.Map
                d_2560_v26_ = _dafny.Map({d_2535_i1_: p2})
                rhs395_ = default__.safeDivisionInt((self).f36, len(d_2560_v26_))
                rhs396_ = self.f35
                lhs327_ = globalState
                lhs328_ = globalState
                lhs327_.f6 = rhs395_
                lhs328_.f12 = rhs396_
                d_2542_v11_ = (d_2542_v11_) | (d_2542_v11_)
            elif True:
                index394_ = default__.safeIndex(560, (p1).length(0))
                (p1)[index394_] = _dafny.CodePoint('s')
                d_2561_v27_: D2
                d_2561_v27_ = D2_DC6(len(d_2540_v9_), d_2539_v8_, p0, _dafny.SeqWithoutIsStrInference([(self).f36]), (self).f27)
                pat_let_tv55_ = p0
                d_2562_v28_: _dafny.Array
                nw404_ = _dafny.Array(None, 9)
                nw404_[int(0)] = self.f35
                nw404_[int(1)] = (d_2536_v5_)[default__.safeIndex(len(d_2540_v9_), len(d_2536_v5_))]
                nw404_[int(2)] = p0
                nw404_[int(3)] = (_dafny.SeqWithoutIsStrInference([d_2545_v14_ for d_2563_i3_ in range(default__.abs(-651))])) <= (d_2540_v9_)
                def iife226_(_pat_let62_0):
                    def iife227_(d_2564_dt__update__tmp_h1_):
                        def iife228_(_pat_let63_0):
                            def iife229_(d_2565_dt__update_hcf1_h0_):
                                return D0_DC1(d_2565_dt__update_hcf1_h0_)
                            return iife229_(_pat_let63_0)
                        return iife228_((self).f27)
                    return iife227_(_pat_let62_0)
                nw404_[int(4)] = (d_2537_v6_).fm4(iife226_(d_2538_v7_), globalState)
                nw404_[int(5)] = True
                nw404_[int(6)] = self.f35
                def iife230_(_pat_let64_0):
                    def iife231_(d_2566_dt__update__tmp_h2_):
                        def iife232_(_pat_let65_0):
                            def iife233_(d_2567_dt__update_hcf11_h0_):
                                return D2_DC6((d_2566_dt__update__tmp_h2_).cf7, (d_2566_dt__update__tmp_h2_).cf8, (d_2566_dt__update__tmp_h2_).cf9, (d_2566_dt__update__tmp_h2_).cf10, d_2567_dt__update_hcf11_h0_)
                            return iife233_(_pat_let65_0)
                        return iife232_(pat_let_tv55_)
                    return iife231_(_pat_let64_0)
                nw404_[int(7)] = (iife230_(d_2561_v27_)).cf11
                nw404_[int(8)] = (self).f27
                d_2562_v28_ = nw404_
                index395_ = default__.safeIndex(461, (d_2562_v28_).length(0))
                (d_2562_v28_)[index395_] = default__.fm0(d_2540_v9_, (self).f27, p0, d_2551_v19_, globalState)
                d_2568_v29_: _dafny.Map
                d_2568_v29_ = _dafny.Map({(self).f36: len(d_2542_v11_)})
                index396_ = default__.safeIndex(560, (p1).length(0))
                index397_ = default__.safeIndex(461, (d_2562_v28_).length(0))
                rhs397_ = d_2540_v9_
                rhs398_ = d_2545_v14_
                rhs399_ = default__.fm0((d_2540_v9_ if (self).f27 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yxnjamfg"))), p0, not(p0), d_2551_v19_, globalState)
                rhs400_ = (((d_2568_v29_)[(self).f36] if ((self).f36) in (d_2568_v29_) else (self).f36)) - ((0) - ((0) - ((self).f28)))
                lhs329_ = p1
                lhs330_ = default__.safeIndex(560, (p1).length(0))
                lhs331_ = d_2562_v28_
                lhs332_ = default__.safeIndex(461, (d_2562_v28_).length(0))
                lhs333_ = globalState
                d_2540_v9_ = rhs397_
                lhs329_[lhs330_] = rhs398_
                lhs331_[lhs332_] = rhs399_
                lhs333_.f2 = rhs400_
                d_2562_v28_ = d_2562_v28_
                (globalState).f2 = (self).f28
                d_2542_v11_ = d_2542_v11_
                (globalState).f14 = (d_2536_v5_)[default__.safeIndex(default__.safeModuloInt((self).f36, p2), len(d_2536_v5_))]
        d_2569_v30_: _dafny.Array
        def lambda117_(d_2570_i5_):
            return not ((self).f27) or (self.f35)

        init67_ = lambda117_
        nw405_ = _dafny.Array(None, 1)
        for i0_67_ in range(nw405_.length(0)):
            nw405_[i0_67_] = init67_(i0_67_)
        d_2569_v30_ = nw405_
        guard_loop_10_: int
        for guard_loop_10_ in _dafny.IntegerRange(0, (d_2569_v30_).length(0)):
            d_2571_i4_: int = guard_loop_10_
            if (True) and (((0) <= (d_2571_i4_)) and ((d_2571_i4_) < ((d_2569_v30_).length(0)))):
                (d_2569_v30_)[(d_2571_i4_)] = self.f35
        rhs401_ = (self.f35 if not((self).f27) else (self).f27)
        rhs402_ = False
        lhs334_ = globalState
        lhs335_ = globalState
        lhs334_.f25 = rhs401_
        lhs335_.f21 = rhs402_
        d_2572_i6_: int
        d_2572_i6_ = 0
        with _dafny.label("15"):
            while not (((self).f29) > (p2)) or (p0):
                with _dafny.c_label("15"):
                    if (d_2572_i6_) >= (100):
                        raise _dafny.Break("15")
                    d_2572_i6_ = (d_2572_i6_) + (1)
                    d_2573_v31_: C0
                    nw406_ = C0()
                    nw406_.ctor__()
                    d_2573_v31_ = nw406_
                    d_2574_v32_: _dafny.Map
                    d_2574_v32_ = _dafny.Map({(self).f28: p0})
                    d_2575_v33_: D7
                    d_2575_v33_ = D7_DC15((d_2574_v32_) | (_dafny.Map({-104: True})))
                    d_2575_v33_ = d_2575_v33_
                    d_2576_v34_: _dafny.Array
                    nw407_ = _dafny.Array(int(0), 4)
                    d_2576_v34_ = nw407_
                    index398_ = default__.safeIndex(275, (d_2576_v34_).length(0))
                    (d_2576_v34_)[index398_] = 808
                    index399_ = default__.safeIndex(275, (d_2576_v34_).length(0))
                    (d_2576_v34_)[index399_] = (0) - ((0) - ((p2) - ((self).f28)))
                    d_2577_v35_: _dafny.Seq
                    d_2577_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xu"))
                    d_2578_v37_: _dafny.MultiSet
                    def iife234_():
                        coll102_ = _dafny.Set()
                        compr_102_: int
                        for compr_102_ in _dafny.IntegerRange(-334, 175):
                            d_2579_v36_: int = compr_102_
                            if ((-334) <= (d_2579_v36_)) and ((d_2579_v36_) < (175)):
                                coll102_ = coll102_.union(_dafny.Set([default__.safeDivisionInt(d_2579_v36_, (d_2576_v34_)[default__.safeIndex(275, (d_2576_v34_).length(0))])]))
                        return _dafny.Set(coll102_)
                    d_2578_v37_ = _dafny.MultiSet([p0, default__.fm0(d_2577_v35_, p0, self.f35, iife234_()
                    , globalState), True])
                    d_2580_v38_: _dafny.Set
                    d_2580_v38_ = _dafny.Set({d_2578_v37_, d_2578_v37_})
                    d_2581_v39_: str
                    d_2581_v39_ = _dafny.CodePoint('b')
                    d_2582_v40_: _dafny.Map
                    d_2582_v40_ = _dafny.Map({(d_2576_v34_)[default__.safeIndex(275, (d_2576_v34_).length(0))]: d_2581_v39_})
                    d_2583_v41_: _dafny.Array
                    nw408_ = _dafny.Array(None, 17)
                    nw408_[int(0)] = (d_2577_v35_) + (d_2577_v35_)
                    nw408_[int(1)] = d_2577_v35_
                    nw408_[int(2)] = d_2577_v35_
                    nw408_[int(3)] = d_2577_v35_
                    nw408_[int(4)] = (self).fm2((self).f29, _dafny.Map({d_2581_v39_: d_2582_v40_}), globalState)
                    nw408_[int(5)] = (_dafny.SeqWithoutIsStrInference([d_2581_v39_ for d_2584_i7_ in range(default__.abs(-918))])) + (d_2577_v35_)
                    nw408_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))
                    nw408_[int(7)] = d_2577_v35_
                    nw408_[int(8)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_2585_i8_ in range(default__.abs(831))])
                    nw408_[int(9)] = d_2577_v35_
                    nw408_[int(10)] = d_2577_v35_
                    nw408_[int(11)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yhoandip"))
                    nw408_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hkgur"))
                    nw408_[int(13)] = d_2577_v35_
                    nw408_[int(14)] = d_2577_v35_
                    nw408_[int(15)] = (d_2577_v35_) + (d_2577_v35_)
                    nw408_[int(16)] = (d_2577_v35_).set(default__.safeIndex((d_2576_v34_)[default__.safeIndex(275, (d_2576_v34_).length(0))], len(d_2577_v35_)), d_2581_v39_)
                    d_2583_v41_ = nw408_
                    rhs403_ = _dafny.Set({(d_2578_v37_).intersection((D21_DC52(d_2578_v37_)).cf75)})
                    rhs404_ = (507) == ((d_2576_v34_)[default__.safeIndex(275, (d_2576_v34_).length(0))])
                    rhs405_ = p2
                    rhs406_ = d_2583_v41_
                    rhs407_ = self.f35
                    lhs336_ = globalState
                    lhs337_ = globalState
                    lhs338_ = globalState
                    d_2580_v38_ = rhs403_
                    lhs336_.f21 = rhs404_
                    lhs337_.f2 = rhs405_
                    d_2583_v41_ = rhs406_
                    lhs338_.f21 = rhs407_
                    pass
            pass

    def m1(self, globalState):
        r0: int = int(0)
        (globalState).f2 = (0) - (((self).f28) - ((self).f28))
        if not((self).f27):
            d_2586_v0_: _dafny.Array
            def lambda118_(d_2587_i0_):
                return default__.safeModuloInt(d_2587_i0_, 924)

            init68_ = lambda118_
            nw409_ = _dafny.Array(None, 24)
            for i0_68_ in range(nw409_.length(0)):
                nw409_[i0_68_] = init68_(i0_68_)
            d_2586_v0_ = nw409_
            d_2588_v1_: _dafny.Seq
            d_2588_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mf"))
            d_2589_v2_: _dafny.Seq
            d_2589_v2_ = _dafny.SeqWithoutIsStrInference([(self).f29])
            d_2590_v3_: _dafny.MultiSet
            d_2590_v3_ = _dafny.MultiSet([self.f35])
            nw410_ = _dafny.Array(None, 7)
            nw410_[int(0)] = (self).f28
            nw410_[int(1)] = len(d_2588_v1_)
            nw410_[int(2)] = (self).f36
            nw410_[int(3)] = (self).f36
            nw410_[int(4)] = len(((d_2589_v2_).set(default__.safeIndex(502, len(d_2589_v2_)), (self).f28)).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "esx"))), len((d_2589_v2_).set(default__.safeIndex(502, len(d_2589_v2_)), (self).f28))), (d_2590_v3_).cardinality))
            nw410_[int(5)] = (0) - ((self).f36)
            nw410_[int(6)] = -94
            d_2586_v0_ = nw410_
            (globalState).f6 = ((((_dafny.MultiSet([self.f35])).set(self.f35, default__.abs((self).f28))) - (d_2590_v3_)).cardinality) * ((self).f28)
            d_2591_v4_: _dafny.Map
            d_2591_v4_ = _dafny.Map({(self).f36: self.f35})
            d_2592_v5_: _dafny.Seq
            d_2592_v5_ = _dafny.SeqWithoutIsStrInference([self.f35, (self).f27])
            d_2593_v6_: _dafny.MultiSet
            d_2593_v6_ = _dafny.MultiSet([(self).f29, len(d_2592_v5_)])
            d_2589_v2_ = (d_2589_v2_) + (((d_2589_v2_).set(default__.safeIndex((self).f28, len(d_2589_v2_)), len(d_2588_v1_))) + (((d_2589_v2_).set(default__.safeIndex(len(d_2591_v4_), len(d_2589_v2_)), (self).f29)).set(default__.safeIndex((self).f36, len((d_2589_v2_).set(default__.safeIndex(len(d_2591_v4_), len(d_2589_v2_)), (self).f29))), (d_2593_v6_).cardinality)))
            (globalState).f5 = self.f35
            (globalState).f21 = (self).f27
        elif True:
            (globalState).f2 = ((self).f36) + (168)
            d_2594_v7_: _dafny.Seq
            d_2594_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rnlqn"))
            d_2595_v8_: _dafny.Map
            d_2595_v8_ = _dafny.Map({(self).f36: (self).f27})
            d_2596_v9_: _dafny.Set
            d_2596_v9_ = _dafny.Set({(self).f28, len(d_2595_v8_)})
            d_2597_v10_: _dafny.Seq
            d_2597_v10_ = _dafny.SeqWithoutIsStrInference([default__.fm0(d_2594_v7_, not(self.f35), (self).f27, d_2596_v9_, globalState)])
            (globalState).f2 = len(((d_2597_v10_) + (d_2597_v10_)) + ((d_2597_v10_) + (d_2597_v10_)))
            d_2594_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fkpbm"))
            d_2598_v11_: _dafny.Map
            d_2598_v11_ = _dafny.Map({self.f35: (self).f27})
            d_2598_v11_ = d_2598_v11_
            d_2599_v12_: _dafny.Map
            d_2599_v12_ = _dafny.Map({(_dafny.MultiSet([(self).f29])).cardinality: len(d_2596_v9_)})
            d_2600_v13_: _dafny.MultiSet
            d_2600_v13_ = _dafny.MultiSet([default__.fm0(d_2594_v7_, (self).f27, (self).f27, d_2596_v9_, globalState)])
            d_2601_v14_: _dafny.Map
            d_2601_v14_ = _dafny.Map({self.f35: d_2600_v13_})
            d_2602_v15_: _dafny.MultiSet
            d_2602_v15_ = _dafny.MultiSet([len(d_2601_v14_)])
            d_2603_v16_: _dafny.Array
            nw411_ = _dafny.Array(None, 29)
            nw411_[int(0)] = (self).f29
            nw411_[int(1)] = (self).f36
            nw411_[int(2)] = (self).f36
            nw411_[int(3)] = (self).f28
            nw411_[int(4)] = (self).f29
            nw411_[int(5)] = len(d_2596_v9_)
            nw411_[int(6)] = len(d_2594_v7_)
            nw411_[int(7)] = -148
            nw411_[int(8)] = default__.fm1((self).f27, d_2597_v10_, (self).f27, globalState)
            nw411_[int(9)] = len(_dafny.SeqWithoutIsStrInference([(self).f36 for d_2604_i1_ in range(default__.abs(665))]))
            nw411_[int(10)] = len(d_2599_v12_)
            nw411_[int(11)] = (self).f29
            nw411_[int(12)] = (self).f28
            nw411_[int(13)] = (self).f29
            nw411_[int(14)] = (self).f36
            nw411_[int(15)] = ((d_2602_v15_)[(self).f29] if ((self).f29) in (d_2602_v15_) else (self).f28)
            nw411_[int(16)] = (self).f29
            nw411_[int(17)] = (self).f28
            nw411_[int(18)] = (self).f28
            nw411_[int(19)] = (self).f36
            nw411_[int(20)] = (0) - (default__.fm1(not(self.f35), d_2597_v10_, self.f35, globalState))
            nw411_[int(21)] = len(_dafny.SeqWithoutIsStrInference([(self).f27, self.f35]))
            nw411_[int(22)] = (self).f36
            nw411_[int(23)] = (self).f29
            nw411_[int(24)] = (self).f36
            nw411_[int(25)] = (self).f29
            nw411_[int(26)] = (self).f28
            nw411_[int(27)] = (self).f28
            nw411_[int(28)] = (self).f28
            d_2603_v16_ = nw411_
            source38_ = d_2603_v16_
            d_2605___mcc_h0_ = source38_
            d_2606_cf5_ = d_2605___mcc_h0_
            d_2607_v17_: _dafny.Array
            def lambda119_(d_2608_i2_):
                return not(False)

            init69_ = lambda119_
            nw412_ = _dafny.Array(None, 23)
            for i0_69_ in range(nw412_.length(0)):
                nw412_[i0_69_] = init69_(i0_69_)
            d_2607_v17_ = nw412_
            index400_ = default__.safeIndex(968, (d_2607_v17_).length(0))
            (d_2607_v17_)[index400_] = (self).f27
            d_2609_v18_: _dafny.Set
            d_2609_v18_ = _dafny.Set({(self).f27, self.f35})
            index401_ = default__.safeIndex(968, (d_2607_v17_).length(0))
            rhs408_ = (d_2597_v10_) + (d_2597_v10_)
            rhs409_ = ((d_2609_v18_) | (default__.fm26(globalState))).issubset(d_2609_v18_)
            lhs339_ = d_2607_v17_
            lhs340_ = default__.safeIndex(968, (d_2607_v17_).length(0))
            d_2597_v10_ = rhs408_
            lhs339_[lhs340_] = rhs409_
            d_2610_v19_: str
            d_2610_v19_ = _dafny.CodePoint('c')
            d_2611_v20_: _dafny.Set
            d_2611_v20_ = _dafny.Set({d_2610_v19_, default__.fm24(self.f35, globalState)})
            d_2612_v21_: _dafny.Seq
            d_2612_v21_ = _dafny.SeqWithoutIsStrInference([(self).f36])
            d_2613_v22_: _dafny.Set
            d_2613_v22_ = _dafny.Set({d_2612_v21_, d_2612_v21_})
            rhs410_ = (self).f27
            rhs411_ = d_2611_v20_
            rhs412_ = ((len(d_2594_v7_)) - (-176)) < ((len(d_2596_v9_)) - (len(d_2613_v22_)))
            lhs341_ = globalState
            lhs342_ = globalState
            lhs341_.f12 = rhs410_
            d_2611_v20_ = rhs411_
            lhs342_.f21 = rhs412_
            d_2614_v23_: T1
            nw413_ = C5()
            nw413_.ctor__((d_2607_v17_)[default__.safeIndex(968, (d_2607_v17_).length(0))], (self).f36, default__.fm1(not(not(self.f35)), _dafny.SeqWithoutIsStrInference([(d_2607_v17_)[default__.safeIndex(968, (d_2607_v17_).length(0))], self.f35]), True, globalState), (self).f27)
            d_2614_v23_ = nw413_
            d_2615_v24_: _dafny.Map
            d_2615_v24_ = _dafny.Map({d_2614_v23_: (d_2614_v23_).f28})
            d_2615_v24_ = (d_2615_v24_).set(d_2614_v23_, (d_2614_v23_).f28)
            (globalState).f25 = self.f35
        d_2616_v25_: _dafny.Seq
        d_2616_v25_ = _dafny.SeqWithoutIsStrInference([self.f35])
        d_2617_v26_: _dafny.Map
        d_2617_v26_ = _dafny.Map({(self).f29: (d_2616_v25_) + (d_2616_v25_)})
        d_2618_v27_: _dafny.MultiSet
        d_2618_v27_ = _dafny.MultiSet([756])
        d_2617_v26_ = (d_2617_v26_) | (((d_2617_v26_).set((d_2618_v27_).cardinality, d_2616_v25_)) | (d_2617_v26_))
        (globalState).f14 = self.f35
        d_2619_v28_: D6
        d_2619_v28_ = D6_DC14(self.f35, (self).f36)
        d_2620_v29_: _dafny.Array
        nw414_ = _dafny.Array(None, 15)
        nw414_[int(0)] = d_2619_v28_
        def iife235_(_pat_let66_0):
            def iife236_(d_2621_dt__update__tmp_h0_):
                def iife237_(_pat_let67_0):
                    def iife238_(d_2622_dt__update_hcf22_h0_):
                        return D6_DC14(d_2622_dt__update_hcf22_h0_, (d_2621_dt__update__tmp_h0_).cf23)
                    return iife238_(_pat_let67_0)
                return iife237_((self).f27)
            return iife236_(_pat_let66_0)
        nw414_[int(1)] = iife235_(d_2619_v28_)
        nw414_[int(2)] = d_2619_v28_
        nw414_[int(3)] = d_2619_v28_
        nw414_[int(4)] = d_2619_v28_
        nw414_[int(5)] = d_2619_v28_
        nw414_[int(6)] = d_2619_v28_
        nw414_[int(7)] = D6_DC14((self).f27, (self).f36)
        nw414_[int(8)] = d_2619_v28_
        nw414_[int(9)] = d_2619_v28_
        nw414_[int(10)] = d_2619_v28_
        nw414_[int(11)] = d_2619_v28_
        nw414_[int(12)] = d_2619_v28_
        nw414_[int(13)] = d_2619_v28_
        nw414_[int(14)] = d_2619_v28_
        d_2620_v29_ = nw414_
        d_2623_v30_: _dafny.Array
        d_2623_v30_ = d_2620_v29_
        d_2623_v30_ = d_2623_v30_
        d_2624_v31_: _dafny.MultiSet
        d_2624_v31_ = _dafny.MultiSet([self.f35])
        d_2625_v32_: _dafny.Seq
        d_2625_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "raeq"))
        d_2626_v33_: _dafny.Set
        d_2626_v33_ = _dafny.Set({(self).f29, (self).f36})
        d_2627_v34_: _dafny.Map
        d_2627_v34_ = _dafny.Map({default__.fm0(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_2628_i3_ in range(default__.abs(794))]), not((self).f27), self.f35, default__.fm35(default__.fm65((self).f28, (d_2624_v31_).cardinality, (self).f36, globalState), (self).f27, (self).f27, globalState), globalState): ((self).f27) or (default__.fm0(d_2625_v32_, self.f35, (self).f27, d_2626_v33_, globalState))})
        d_2629_v35_: D21
        d_2629_v35_ = D21_DC54((self).f27, len(d_2625_v32_), default__.fm1((self).f27, _dafny.SeqWithoutIsStrInference([(self).f27]), self.f35, globalState))
        d_2627_v34_ = (d_2627_v34_).set(not((d_2629_v35_).cf79), not(not ((self).f27) or (self.f35)))
        r0 = (self).f28
        return r0

    def m2(self, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: bool = False
        if (self).f27:
            d_2630_v0_: _dafny.Set
            d_2630_v0_ = _dafny.Set({True, (self).f27})
            d_2631_v1_: _dafny.Seq
            d_2631_v1_ = _dafny.SeqWithoutIsStrInference([len(d_2630_v0_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rs")))])
            d_2632_v2_: _dafny.Map
            d_2632_v2_ = _dafny.Map({self.f35: not(self.f35)})
            d_2633_v3_: _dafny.Map
            d_2633_v3_ = _dafny.Map({len(d_2632_v2_): self.f35})
            d_2634_v4_: _dafny.Set
            d_2634_v4_ = _dafny.Set({len(d_2631_v1_), 112, (self).f28, len(d_2633_v3_), (self).f36})
            d_2635_v5_: _dafny.Map
            d_2635_v5_ = _dafny.Map({(self).f29: (self).f29})
            d_2636_v7_: _dafny.MultiSet
            def iife239_():
                coll103_ = _dafny.Set()
                compr_103_: int
                for compr_103_ in (d_2635_v5_).keys.Elements:
                    d_2637_v6_: int = compr_103_
                    if (d_2637_v6_) in (d_2635_v5_):
                        coll103_ = coll103_.union(_dafny.Set([default__.safeDivisionInt(d_2637_v6_, len(_dafny.Map({False: not(False)})))]))
                return _dafny.Set(coll103_)
            d_2636_v7_ = (_dafny.MultiSet([d_2634_v4_])).intersection(_dafny.MultiSet([iife239_()
            ]))
            d_2636_v7_ = d_2636_v7_
            d_2638_v8_: _dafny.Array
            nw415_ = _dafny.Array(D19.default()(), 8)
            d_2638_v8_ = nw415_
            d_2639_v9_: str
            d_2639_v9_ = _dafny.CodePoint('g')
            d_2640_v10_: _dafny.MultiSet
            d_2640_v10_ = _dafny.MultiSet([d_2639_v9_, d_2639_v9_])
            d_2641_v11_: D19
            d_2641_v11_ = D19_DC45(d_2640_v10_)
            d_2642_v12_: D19
            d_2642_v12_ = D19_DC47(d_2641_v11_)
            index402_ = default__.safeIndex(393, (d_2638_v8_).length(0))
            (d_2638_v8_)[index402_] = d_2642_v12_
            d_2643_v13_: _dafny.MultiSet
            d_2643_v13_ = _dafny.MultiSet([(self).f27])
            d_2644_v14_: C4
            nw416_ = C4()
            nw416_.ctor__(d_2643_v13_, False)
            d_2644_v14_ = nw416_
            index403_ = default__.safeIndex(393, (d_2638_v8_).length(0))
            rhs413_ = (self).f28
            rhs414_ = d_2642_v12_
            rhs415_ = d_2644_v14_
            lhs343_ = globalState
            lhs344_ = d_2638_v8_
            lhs345_ = default__.safeIndex(393, (d_2638_v8_).length(0))
            lhs343_.f6 = rhs413_
            lhs344_[lhs345_] = rhs414_
            d_2644_v14_ = rhs415_
            (globalState).f12 = ((self).f28) == ((self).f36)
            d_2645_v15_: _dafny.Map
            d_2645_v15_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([862]): (d_2631_v1_)[default__.safeIndex((self).f29, len(d_2631_v1_))]})
            d_2645_v15_ = (d_2645_v15_).set((d_2631_v1_).set(default__.safeIndex((self).f28, len(d_2631_v1_)), (self).f36), (self).f28)
            def iife240_(_pat_let68_0):
                def iife241_(d_2646_dt__update__tmp_h0_):
                    def iife242_(_pat_let69_0):
                        def iife243_(d_2647_dt__update_hcf77_h0_):
                            return D21_DC53((d_2646_dt__update__tmp_h0_).cf76, d_2647_dt__update_hcf77_h0_, (d_2646_dt__update__tmp_h0_).cf78)
                        return iife243_(_pat_let69_0)
                    return iife242_(((self).f28) * ((self).f36))
                return iife241_(_pat_let68_0)
            source39_ = iife240_(D21_DC53((self).f27, (self).f28, ((d_2633_v3_)[len(d_2630_v0_)] if (len(d_2630_v0_)) in (d_2633_v3_) else self.f35)))
            if source39_.is_DC53:
                d_2648___mcc_h0_ = source39_.cf76
                d_2649___mcc_h1_ = source39_.cf77
                d_2650___mcc_h2_ = source39_.cf78
                d_2651_cf78_ = d_2650___mcc_h2_
                d_2652_cf77_ = d_2649___mcc_h1_
                d_2653_cf76_ = d_2648___mcc_h0_
                d_2654_v16_: _dafny.Array
                nw417_ = _dafny.Array(int(0), 1)
                d_2654_v16_ = nw417_
                d_2655_v17_: _dafny.Seq
                d_2655_v17_ = _dafny.SeqWithoutIsStrInference([d_2654_v16_])
                d_2655_v17_ = d_2655_v17_
                d_2652_cf77_ = (self).f36
                d_2656_v18_: _dafny.Array
                def lambda120_(d_2657_v1_):
                    def lambda121_(d_2658_i0_):
                        return d_2657_v1_

                    return lambda121_

                init70_ = lambda120_(d_2631_v1_)
                nw418_ = _dafny.Array(None, 28)
                for i0_70_ in range(nw418_.length(0)):
                    nw418_[i0_70_] = init70_(i0_70_)
                d_2656_v18_ = nw418_
                d_2656_v18_ = d_2656_v18_
                index404_ = default__.safeIndex(140, (d_2654_v16_).length(0))
                (d_2654_v16_)[index404_] = (self).f28
                index405_ = default__.safeIndex(140, (d_2654_v16_).length(0))
                (d_2654_v16_)[index405_] = d_2652_cf77_
            elif source39_.is_DC54:
                d_2659___mcc_h3_ = source39_.cf79
                d_2660___mcc_h4_ = source39_.cf80
                d_2661___mcc_h5_ = source39_.cf81
                d_2662_cf81_ = d_2661___mcc_h5_
                d_2663_cf80_ = d_2660___mcc_h4_
                d_2664_cf79_ = d_2659___mcc_h3_
                d_2665_v19_: _dafny.Seq
                d_2665_v19_ = _dafny.SeqWithoutIsStrInference([(self).f27, False])
                d_2666_v20_: _dafny.Array
                nw419_ = _dafny.Array(None, 14)
                nw419_[int(0)] = d_2665_v19_
                nw419_[int(1)] = d_2665_v19_
                nw419_[int(2)] = d_2665_v19_
                nw419_[int(3)] = d_2665_v19_
                nw419_[int(4)] = d_2665_v19_
                nw419_[int(5)] = d_2665_v19_
                nw419_[int(6)] = _dafny.SeqWithoutIsStrInference([True])
                nw419_[int(7)] = d_2665_v19_
                nw419_[int(8)] = d_2665_v19_
                nw419_[int(9)] = d_2665_v19_
                nw419_[int(10)] = d_2665_v19_
                nw419_[int(11)] = d_2665_v19_
                nw419_[int(12)] = d_2665_v19_
                nw419_[int(13)] = d_2665_v19_
                d_2666_v20_ = nw419_
                d_2667_v21_: _dafny.MultiSet
                d_2667_v21_ = _dafny.MultiSet([d_2666_v20_])
                d_2668_v23_: _dafny.Map
                def iife244_():
                    coll104_ = _dafny.Map()
                    compr_104_: int
                    for compr_104_ in _dafny.IntegerRange(7, 67):
                        d_2669_v22_: int = compr_104_
                        if ((7) <= (d_2669_v22_)) and ((d_2669_v22_) < (67)):
                            coll104_[(d_2669_v22_) * ((self).f28)] = True
                    return _dafny.Map(coll104_)
                d_2668_v23_ = _dafny.Map({(self).f27: _dafny.Map({781: len(iife244_()
                )})})
                d_2670_v24_: _dafny.Map
                d_2670_v24_ = _dafny.Map({default__.safeModuloInt(((d_2667_v21_)[d_2666_v20_] if (d_2666_v20_) in (d_2667_v21_) else (self).f36), d_2662_cf81_): d_2668_v23_})
                d_2670_v24_ = (d_2670_v24_).set(default__.safeModuloInt((self).f29, (self).f28), d_2668_v23_)
                d_2671_v25_: D18
                d_2671_v25_ = D18_DC42(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fkyp")), self.f35)
                d_2671_v25_ = default__.fm63((self).f36, (self).f29, globalState)
                (globalState).f5 = ((self).f29) in (_dafny.SeqWithoutIsStrInference([394, 333, d_2663_cf80_, default__.fm1(d_2664_cf79_, d_2665_v19_, d_2664_cf79_, globalState), default__.fm1(d_2664_cf79_, d_2665_v19_, d_2664_cf79_, globalState)]))
                d_2672_v26_: _dafny.Array
                nw420_ = _dafny.Array(None, 6)
                nw420_[int(0)] = d_2664_cf79_
                nw420_[int(1)] = (self).f27
                nw420_[int(2)] = (self).f27
                nw420_[int(3)] = (False if d_2664_cf79_ else d_2664_cf79_)
                nw420_[int(4)] = self.f35
                nw420_[int(5)] = self.f35
                d_2672_v26_ = nw420_
                index406_ = default__.safeIndex(702, (d_2672_v26_).length(0))
                (d_2672_v26_)[index406_] = (self).f27
                index407_ = default__.safeIndex(702, (d_2672_v26_).length(0))
                (d_2672_v26_)[index407_] = (self).f27
            elif source39_.is_DC55:
                d_2673___mcc_h6_ = source39_.cf82
                d_2674_cf82_ = d_2673___mcc_h6_
                d_2675_v27_: D0
                d_2675_v27_ = D0_DC2((self).f29, (self).f36)
                d_2676_v28_: _dafny.Seq
                d_2676_v28_ = _dafny.SeqWithoutIsStrInference([False, self.f35, (self).f27])
                d_2677_v29_: D6
                d_2677_v29_ = D6_DC13(d_2676_v28_)
                pat_let_tv56_ = d_2676_v28_
                d_2678_v30_: _dafny.Array
                nw421_ = _dafny.Array(None, 11)
                nw421_[int(0)] = d_2677_v29_
                nw421_[int(1)] = d_2677_v29_
                nw421_[int(2)] = d_2677_v29_
                nw421_[int(3)] = D6_DC13(d_2676_v28_)
                nw421_[int(4)] = d_2677_v29_
                nw421_[int(5)] = d_2677_v29_
                nw421_[int(6)] = D6_DC13(_dafny.SeqWithoutIsStrInference([self.f35, (self).f27, self.f35, self.f35]))
                nw421_[int(7)] = D6_DC13(default__.fm32((self).f27, d_2631_v1_, (self).f27, globalState))
                nw421_[int(8)] = D6_DC13(d_2676_v28_)
                def iife245_(_pat_let70_0):
                    def iife246_(d_2679_dt__update__tmp_h1_):
                        def iife247_(_pat_let71_0):
                            def iife248_(d_2680_dt__update_hcf21_h0_):
                                return D6_DC13(d_2680_dt__update_hcf21_h0_)
                            return iife248_(_pat_let71_0)
                        return iife247_(pat_let_tv56_)
                    return iife246_(_pat_let70_0)
                nw421_[int(9)] = iife245_(d_2677_v29_)
                nw421_[int(10)] = d_2677_v29_
                d_2678_v30_ = nw421_
                index408_ = default__.safeIndex(456, (d_2678_v30_).length(0))
                (d_2678_v30_)[index408_] = d_2677_v29_
                d_2681_v31_: _dafny.Array
                def lambda122_(d_2682_cf82_):
                    def lambda123_(d_2683_i1_):
                        return (d_2683_i1_) + (d_2682_cf82_)

                    return lambda123_

                init71_ = lambda122_(d_2674_cf82_)
                nw422_ = _dafny.Array(None, 4)
                for i0_71_ in range(nw422_.length(0)):
                    nw422_[i0_71_] = init71_(i0_71_)
                d_2681_v31_ = nw422_
                index409_ = default__.safeIndex(104, (d_2681_v31_).length(0))
                (d_2681_v31_)[index409_] = default__.fm1(True, d_2676_v28_, self.f35, globalState)
                d_2684_v32_: _dafny.Array
                nw423_ = _dafny.Array(False, 22)
                d_2684_v32_ = nw423_
                d_2685_v33_: _dafny.Map
                d_2685_v33_ = _dafny.Map({(self).f27: d_2684_v32_})
                index410_ = default__.safeIndex(456, (d_2678_v30_).length(0))
                index411_ = default__.safeIndex(104, (d_2681_v31_).length(0))
                rhs416_ = not((self).f27)
                rhs417_ = d_2675_v27_
                rhs418_ = d_2677_v29_
                rhs419_ = len(d_2685_v33_)
                lhs346_ = globalState
                lhs347_ = d_2678_v30_
                lhs348_ = default__.safeIndex(456, (d_2678_v30_).length(0))
                lhs349_ = d_2681_v31_
                lhs350_ = default__.safeIndex(104, (d_2681_v31_).length(0))
                lhs346_.f12 = rhs416_
                d_2675_v27_ = rhs417_
                lhs347_[lhs348_] = rhs418_
                lhs349_[lhs350_] = rhs419_
                d_2686_v34_: _dafny.Seq
                d_2686_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dnxmgea"))
                d_2687_v35_: _dafny.Map
                d_2687_v35_ = _dafny.Map({len(d_2686_v34_): d_2681_v31_})
                d_2688_v36_: _dafny.Set
                d_2688_v36_ = _dafny.Set({((d_2687_v35_)[(d_2681_v31_)[default__.safeIndex(104, (d_2681_v31_).length(0))]] if ((d_2681_v31_)[default__.safeIndex(104, (d_2681_v31_).length(0))]) in (d_2687_v35_) else d_2681_v31_), d_2681_v31_, d_2681_v31_})
                d_2689_v37_: _dafny.Map
                d_2689_v37_ = _dafny.Map({(0) - ((self).f36): d_2688_v36_})
                d_2689_v37_ = (d_2689_v37_).set((self).f28, d_2688_v36_)
                d_2690_v38_: C2
                nw424_ = C2()
                nw424_.ctor__((self).f27, (d_2643_v13_) != ((d_2644_v14_).f37))
                d_2690_v38_ = nw424_
                index412_ = default__.safeIndex(104, (d_2681_v31_).length(0))
                (d_2681_v31_)[index412_] = (default__.fm1((self).f27, d_2676_v28_, (self).f27, globalState)) * (d_2674_cf82_)
            elif True:
                d_2691___mcc_h7_ = source39_.cf75
                d_2692_cf75_ = d_2691___mcc_h7_
                d_2643_v13_ = (_dafny.MultiSet([self.f35, self.f35, (self).f27, self.f35, self.f35])).intersection(((D21_DC52(d_2692_cf75_)).cf75) - (_dafny.MultiSet([self.f35])))
                (globalState).f21 = not((self).f27)
                d_2693_v39_: _dafny.Array
                nw425_ = _dafny.Array(_dafny.MultiSet({}), 22)
                d_2693_v39_ = nw425_
                d_2694_v40_: D30
                d_2694_v40_ = D30_DC73(d_2693_v39_)
                d_2693_v39_ = (d_2694_v40_).cf120
                d_2695_v41_: _dafny.Array
                nw426_ = _dafny.Array(_dafny.MultiSet({}), 23)
                d_2695_v41_ = nw426_
                d_2696_v42_: T0
                nw427_ = C4()
                nw427_.ctor__((d_2644_v14_).f37, (self).f27)
                d_2696_v42_ = nw427_
                d_2697_v43_: _dafny.Array
                def lambda124_(d_2698_v42_):
                    def lambda125_(d_2699_i2_):
                        return not((d_2698_v42_).f27)

                    return lambda125_

                init72_ = lambda124_(d_2696_v42_)
                nw428_ = _dafny.Array(None, 16)
                for i0_72_ in range(nw428_.length(0)):
                    nw428_[i0_72_] = init72_(i0_72_)
                d_2697_v43_ = nw428_
                d_2700_v44_: _dafny.Map
                d_2700_v44_ = _dafny.Map({d_2640_v10_: d_2697_v43_})
                d_2701_v45_: _dafny.MultiSet
                d_2701_v45_ = _dafny.MultiSet([d_2697_v43_, ((d_2700_v44_)[d_2640_v10_] if (d_2640_v10_) in (d_2700_v44_) else d_2697_v43_)])
                d_2702_v46_: _dafny.Map
                d_2702_v46_ = _dafny.Map({d_2696_v42_: d_2701_v45_})
                index413_ = default__.safeIndex(272, (d_2695_v41_).length(0))
                (d_2695_v41_)[index413_] = ((d_2702_v46_)[d_2696_v42_] if (d_2696_v42_) in (d_2702_v46_) else d_2701_v45_)
                index414_ = default__.safeIndex(272, (d_2695_v41_).length(0))
                (d_2695_v41_)[index414_] = d_2701_v45_
        elif True:
            d_2703_v47_: _dafny.MultiSet
            d_2703_v47_ = _dafny.MultiSet([(self).f27, self.f35])
            d_2704_v48_: _dafny.Seq
            d_2704_v48_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ufaa"))
            d_2705_v49_: _dafny.Seq
            d_2705_v49_ = _dafny.SeqWithoutIsStrInference([self.f35, self.f35])
            d_2706_v50_: _dafny.Set
            d_2706_v50_ = _dafny.Set({(self).f27})
            d_2707_v51_: T0
            nw429_ = C4()
            nw429_.ctor__(d_2703_v47_, default__.fm0(d_2704_v48_, (d_2705_v49_)[default__.safeIndex(default__.fm1((self).f27, _dafny.SeqWithoutIsStrInference([True]), False, globalState), len(d_2705_v49_))], (self).f27, _dafny.Set({len(d_2706_v50_)}), globalState))
            d_2707_v51_ = nw429_
            d_2707_v51_ = d_2707_v51_
            (globalState).f12 = self.f35
            if self.f35:
                (globalState).f6 = (0) - (((self).f36) - (len(d_2704_v48_)))
                d_2708_v52_: D18
                d_2708_v52_ = D18_DC42(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_2709_i3_ in range(default__.abs(260))]), (self).f27)
                d_2710_v53_: _dafny.MultiSet
                d_2710_v53_ = _dafny.MultiSet([d_2708_v52_, d_2708_v52_, d_2708_v52_, d_2708_v52_, d_2708_v52_])
                pat_let_tv57_ = d_2707_v51_
                def iife249_(_pat_let72_0):
                    def iife250_(d_2711_dt__update__tmp_h2_):
                        def iife251_(_pat_let73_0):
                            def iife252_(d_2712_dt__update_hcf58_h0_):
                                return D18_DC42((d_2711_dt__update__tmp_h2_).cf57, d_2712_dt__update_hcf58_h0_)
                            return iife252_(_pat_let73_0)
                        return iife251_((pat_let_tv57_).f27)
                    return iife250_(_pat_let72_0)
                d_2710_v53_ = _dafny.MultiSet([iife249_(d_2708_v52_), D18_DC42(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_2713_i4_ in range(default__.abs(766))]), not((d_2707_v51_).f27)), d_2708_v52_])
                d_2714_v54_: _dafny.Array
                nw430_ = _dafny.Array(_dafny.CodePoint('D'), 21)
                d_2714_v54_ = nw430_
                d_2715_v55_: _dafny.Map
                d_2715_v55_ = _dafny.Map({d_2714_v54_: (self).f29})
                (globalState).f2 = ((self).f28) - (((d_2715_v55_)[d_2714_v54_] if (d_2714_v54_) in (d_2715_v55_) else (self).f36))
                d_2716_v56_: _dafny.Array
                def lambda126_(d_2717_i5_):
                    return (d_2717_i5_) * ((self).f36)

                init73_ = lambda126_
                nw431_ = _dafny.Array(None, 6)
                for i0_73_ in range(nw431_.length(0)):
                    nw431_[i0_73_] = init73_(i0_73_)
                d_2716_v56_ = nw431_
                d_2718_v57_: _dafny.MultiSet
                d_2718_v57_ = _dafny.MultiSet([(self).f36])
                d_2719_v58_: _dafny.Map
                d_2719_v58_ = _dafny.Map({(self).f36: (self).f29})
                d_2720_v59_: _dafny.Map
                d_2720_v59_ = _dafny.Map({(self).f27: (d_2707_v51_).f27})
                d_2721_v60_: D18
                d_2721_v60_ = D18_DC44(False, (self).f28, (d_2707_v51_).f27)
                d_2722_v61_: _dafny.Map
                d_2722_v61_ = _dafny.Map({((d_2719_v58_)[len(d_2706_v50_)] if (len(d_2706_v50_)) in (d_2719_v58_) else len(d_2720_v59_)): d_2721_v60_})
                d_2723_v62_: _dafny.Seq
                d_2723_v62_ = _dafny.SeqWithoutIsStrInference([len(d_2722_v61_)])
                d_2724_v63_: _dafny.Seq
                d_2724_v63_ = _dafny.SeqWithoutIsStrInference([d_2723_v62_])
                d_2725_v64_: _dafny.Map
                d_2725_v64_ = _dafny.Map({(d_2718_v57_).cardinality: len(d_2724_v63_)})
                d_2726_v65_: _dafny.Map
                d_2726_v65_ = _dafny.Map({len(d_2725_v64_): (self).f27})
                d_2727_v66_: _dafny.Map
                d_2727_v66_ = _dafny.Map({self.f35: len(d_2705_v49_)})
                d_2728_v67_: D23
                d_2728_v67_ = D23_DC61(d_2716_v56_, ((d_2726_v65_)[len(d_2727_v66_)] if (len(d_2727_v66_)) in (d_2726_v65_) else (self).f27), (self).f36)
                index415_ = default__.safeIndex(5, (d_2716_v56_).length(0))
                (d_2716_v56_)[index415_] = default__.safeDivisionInt((self).f36, (0) - ((d_2728_v67_).cf95))
                index416_ = default__.safeIndex(5, (d_2716_v56_).length(0))
                rhs420_ = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rkdei")))) * ((self).f29)
                rhs421_ = default__.safeModuloInt((self).f29, (self).f29)
                lhs351_ = globalState
                lhs352_ = d_2716_v56_
                lhs353_ = default__.safeIndex(5, (d_2716_v56_).length(0))
                lhs351_.f2 = rhs420_
                lhs352_[lhs353_] = rhs421_
                d_2729_v68_: str
                d_2729_v68_ = _dafny.CodePoint('m')
                d_2730_v69_: _dafny.MultiSet
                d_2730_v69_ = _dafny.MultiSet([(d_2729_v68_ if False else d_2729_v68_), _dafny.CodePoint('e'), d_2729_v68_])
                d_2731_v70_: C2
                nw432_ = C2()
                nw432_.ctor__((d_2707_v51_).f27, (self).f27)
                d_2731_v70_ = nw432_
                index417_ = default__.safeIndex(5, (d_2716_v56_).length(0))
                rhs422_ = (d_2723_v62_) + (default__.fm40(len(d_2705_v49_), (d_2707_v51_).f27, globalState))
                rhs423_ = d_2730_v69_
                rhs424_ = (self).f28
                rhs425_ = (d_2716_v56_)[default__.safeIndex(5, (d_2716_v56_).length(0))]
                rhs426_ = (d_2731_v70_ if not((d_2731_v70_).f34) else d_2731_v70_)
                lhs354_ = d_2716_v56_
                lhs355_ = default__.safeIndex(5, (d_2716_v56_).length(0))
                lhs356_ = globalState
                d_2723_v62_ = rhs422_
                d_2730_v69_ = rhs423_
                lhs354_[lhs355_] = rhs424_
                lhs356_.f2 = rhs425_
                d_2731_v70_ = rhs426_
            elif True:
                d_2732_v71_: _dafny.Array
                nw433_ = _dafny.Array(int(0), 7)
                d_2732_v71_ = nw433_
                d_2733_v72_: _dafny.Map
                d_2733_v72_ = _dafny.Map({d_2732_v71_: (self).f27})
                d_2734_v73_: D19
                d_2734_v73_ = D19_DC46((self).f28, d_2732_v71_, (d_2707_v51_).f27)
                d_2733_v72_ = (d_2733_v72_).set((d_2734_v73_).cf65, True)
                d_2735_v74_: _dafny.Seq
                d_2735_v74_ = _dafny.SeqWithoutIsStrInference([len((D18_DC42(d_2704_v48_, (d_2707_v51_).f27)).cf57)])
                (globalState).f6 = (len(d_2735_v74_)) - (((self).f36) * ((self).f29))
                d_2736_v75_: _dafny.Map
                d_2736_v75_ = _dafny.Map({(d_2707_v51_).f27: (0) - ((self).f36)})
                d_2737_v76_: str
                d_2737_v76_ = _dafny.CodePoint('e')
                d_2738_v77_: _dafny.Map
                d_2738_v77_ = _dafny.Map({self.f35: d_2737_v76_})
                d_2739_v78_: _dafny.Array
                nw434_ = _dafny.Array(False, 9)
                d_2739_v78_ = nw434_
                d_2740_v79_: _dafny.Map
                d_2740_v79_ = _dafny.Map({d_2739_v78_: (self).f36})
                (globalState).f6 = (len((d_2736_v75_).set(self.f35, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sfxb"))).set(default__.safeIndex(-597, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sfxb")))), d_2737_v76_))))) * (default__.safeDivisionInt(len((_dafny.Map({(d_2707_v51_).f27: d_2737_v76_})).set((self).f27, ((d_2738_v77_)[(self).f27] if ((self).f27) in (d_2738_v77_) else d_2737_v76_))), len(d_2740_v79_)))
                r1 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))) + (d_2704_v48_)
                d_2741_v80_: D23
                d_2741_v80_ = D23_DC59(_dafny.Map({self.f35: (d_2707_v51_).f27}))
                d_2741_v80_ = d_2741_v80_
            d_2742_v81_: _dafny.Map
            d_2742_v81_ = _dafny.Map({(self).f28: self.f35})
            d_2743_v82_: _dafny.Seq
            d_2743_v82_ = _dafny.SeqWithoutIsStrInference([len(d_2704_v48_), 327])
            d_2744_v83_: D2
            d_2744_v83_ = D2_DC6(len(d_2742_v81_), _dafny.SeqWithoutIsStrInference([len(d_2704_v48_) for d_2745_i6_ in range(default__.abs(879))]), (d_2707_v51_).f27, d_2743_v82_, (self).f27)
            d_2746_v84_: _dafny.Map
            d_2746_v84_ = _dafny.Map({d_2744_v83_: (0) - (((self).f28 if (d_2707_v51_).f27 else (self).f28))})
            pat_let_tv58_ = d_2743_v82_
            def iife253_(_pat_let74_0):
                def iife254_(d_2747_dt__update__tmp_h3_):
                    def iife255_(_pat_let75_0):
                        def iife256_(d_2748_dt__update_hcf9_h0_):
                            def iife257_(_pat_let76_0):
                                def iife258_(d_2749_dt__update_hcf10_h0_):
                                    return D2_DC6((d_2747_dt__update__tmp_h3_).cf7, (d_2747_dt__update__tmp_h3_).cf8, d_2748_dt__update_hcf9_h0_, d_2749_dt__update_hcf10_h0_, (d_2747_dt__update__tmp_h3_).cf11)
                                return iife258_(_pat_let76_0)
                            return iife257_(pat_let_tv58_)
                        return iife256_(_pat_let75_0)
                    return iife255_(True)
                return iife254_(_pat_let74_0)
            d_2746_v84_ = (d_2746_v84_).set(iife253_(d_2744_v83_), (self).f36)
            d_2750_v85_: int
            out78_: int
            out78_ = (d_2707_v51_).m1(globalState)
            d_2750_v85_ = out78_
        d_2751_i7_: int
        d_2751_i7_ = 0
        with _dafny.label("16"):
            while self.f35:
                with _dafny.c_label("16"):
                    if (d_2751_i7_) >= (100):
                        raise _dafny.Break("16")
                    d_2751_i7_ = (d_2751_i7_) + (1)
                    d_2752_v86_: _dafny.Set
                    d_2752_v86_ = _dafny.Set({(self).f29})
                    d_2753_v87_: _dafny.Array
                    nw435_ = _dafny.Array(None, 8)
                    nw435_[int(0)] = (self).f27
                    nw435_[int(1)] = self.f35
                    nw435_[int(2)] = (self).f27
                    nw435_[int(3)] = (d_2752_v86_).isdisjoint(d_2752_v86_)
                    nw435_[int(4)] = self.f35
                    nw435_[int(5)] = not(self.f35)
                    nw435_[int(6)] = not (self.f35) or ((self).f27)
                    nw435_[int(7)] = (self).f27
                    d_2753_v87_ = nw435_
                    index418_ = default__.safeIndex(920, (d_2753_v87_).length(0))
                    (d_2753_v87_)[index418_] = (self).f27
                    index419_ = default__.safeIndex(920, (d_2753_v87_).length(0))
                    (d_2753_v87_)[index419_] = self.f35
                    d_2754_v88_: _dafny.Seq
                    d_2754_v88_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bcg"))
                    d_2755_v89_: _dafny.Map
                    d_2755_v89_ = _dafny.Map({d_2754_v88_: (-503) - ((self).f36)})
                    (globalState).f6 = ((d_2755_v89_)[(d_2754_v88_ if not(self.f35) else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_2756_i8_ in range(default__.abs(-988))]))] if ((d_2754_v88_ if not(self.f35) else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_2757_i8_ in range(default__.abs(-988))]))) in (d_2755_v89_) else (self).f28)
                    d_2758_v90_: _dafny.MultiSet
                    d_2758_v90_ = _dafny.MultiSet([(d_2753_v87_)[default__.safeIndex(920, (d_2753_v87_).length(0))]])
                    (globalState).f12 = ((self).f27) not in (d_2758_v90_)
                    d_2759_v91_: _dafny.Array
                    nw436_ = _dafny.Array(_dafny.MultiSet({}), 1)
                    d_2759_v91_ = nw436_
                    nw437_ = _dafny.Array(_dafny.MultiSet({}), 22)
                    d_2759_v91_ = nw437_
                    pass
            pass
        d_2760_v92_: _dafny.Set
        d_2760_v92_ = _dafny.Set({self.f35})
        d_2760_v92_ = (d_2760_v92_) - (d_2760_v92_)
        d_2761_v93_: _dafny.Array
        nw438_ = _dafny.Array(int(0), 25)
        d_2761_v93_ = nw438_
        index420_ = default__.safeIndex(735, (d_2761_v93_).length(0))
        (d_2761_v93_)[index420_] = (self).f28
        d_2762_v94_: str
        d_2762_v94_ = _dafny.CodePoint('x')
        d_2763_v95_: _dafny.Seq
        d_2763_v95_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pwlbj"))
        d_2764_v96_: _dafny.Map
        d_2764_v96_ = _dafny.Map({(self).f29: (self).f27})
        d_2765_v97_: _dafny.MultiSet
        d_2765_v97_ = _dafny.MultiSet([self.f35, not(self.f35), ((d_2764_v96_)[(self).f36] if ((self).f36) in (d_2764_v96_) else self.f35)])
        d_2766_v98_: _dafny.Seq
        d_2766_v98_ = _dafny.SeqWithoutIsStrInference([self.f35, (self).f27, self.f35])
        d_2767_v99_: _dafny.Map
        d_2767_v99_ = _dafny.Map({d_2763_v95_: self.f35})
        d_2768_v100_: _dafny.Seq
        d_2768_v100_ = _dafny.SeqWithoutIsStrInference([(self).f28, (self).f29])
        d_2769_v101_: D2
        d_2769_v101_ = D2_DC6((self).f36, d_2768_v100_, (self).f27, d_2768_v100_, (self).f27)
        d_2770_v102_: D23
        d_2770_v102_ = D23_DC61(d_2761_v93_, (self).f27, (self).f29)
        d_2771_v103_: _dafny.Array
        nw439_ = _dafny.Array(None, 10)
        nw439_[int(0)] = (self).f27
        nw439_[int(1)] = (self).f27
        nw439_[int(2)] = (self).f27
        nw439_[int(3)] = self.f35
        nw439_[int(4)] = (d_2770_v102_).cf94
        nw439_[int(5)] = (self).f27
        nw439_[int(6)] = not(self.f35)
        nw439_[int(7)] = True
        nw439_[int(8)] = True
        nw439_[int(9)] = True
        d_2771_v103_ = nw439_
        d_2772_v104_: _dafny.Map
        d_2772_v104_ = _dafny.Map({(self).f36: d_2771_v103_})
        index421_ = default__.safeIndex(735, (d_2761_v93_).length(0))
        rhs427_ = default__.fm1((default__.fm52(len(d_2763_v95_), self.f35, d_2765_v97_, globalState)) != (d_2766_v98_), d_2766_v98_, (self).f27, globalState)
        rhs428_ = d_2762_v94_
        rhs429_ = ((d_2767_v99_)[d_2763_v95_] if (d_2763_v95_) in (d_2767_v99_) else (self).f27)
        rhs430_ = (d_2769_v101_).cf11
        rhs431_ = len((d_2772_v104_) | (d_2772_v104_))
        lhs357_ = d_2761_v93_
        lhs358_ = default__.safeIndex(735, (d_2761_v93_).length(0))
        lhs359_ = globalState
        lhs360_ = self
        lhs361_ = globalState
        lhs357_[lhs358_] = rhs427_
        d_2762_v94_ = rhs428_
        lhs359_.f25 = rhs429_
        lhs360_.f35 = rhs430_
        lhs361_.f2 = rhs431_
        d_2773_v105_: _dafny.MultiSet
        d_2773_v105_ = _dafny.MultiSet([default__.fm66(not(self.f35), (self).f27, globalState), default__.fm66(self.f35, (self).f27, globalState), default__.fm66((self).f27, self.f35, globalState)])
        d_2774_v106_: _dafny.MultiSet
        d_2774_v106_ = d_2773_v105_
        source40_ = d_2773_v105_
        d_2775___mcc_h8_ = source40_
        d_2776_cf29_ = d_2775___mcc_h8_
        d_2777_v107_: _dafny.Array
        nw440_ = _dafny.Array(D6.default()(), 23)
        d_2777_v107_ = nw440_
        d_2778_v108_: _dafny.Array
        d_2778_v108_ = d_2777_v107_
        source41_ = d_2778_v108_
        d_2779___mcc_h9_ = source41_
        d_2780_cf34_ = d_2779___mcc_h9_
        (globalState).f6 = default__.safeDivisionInt((d_2761_v93_)[default__.safeIndex(735, (d_2761_v93_).length(0))], default__.safeModuloInt((0) - ((self).f36), (self).f36))
        d_2766_v98_ = ((d_2766_v98_) + (d_2766_v98_)) + ((d_2766_v98_ if False else d_2766_v98_))
        d_2768_v100_ = d_2768_v100_
        d_2781_v109_: _dafny.Map
        d_2781_v109_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f28]): (self).f29})
        d_2782_v110_: D3
        d_2782_v110_ = D3_DC8(d_2781_v109_)
        d_2783_v111_: _dafny.Seq
        d_2783_v111_ = _dafny.SeqWithoutIsStrInference([d_2782_v110_, D3_DC8((_dafny.Map({d_2768_v100_: (self).f28})).set(d_2768_v100_, len(_dafny.SeqWithoutIsStrInference([(self).f28]))))])
        (globalState).f2 = len(d_2783_v111_)
        index422_ = default__.safeIndex(735, (d_2761_v93_).length(0))
        (d_2761_v93_)[index422_] = default__.safeDivisionInt(default__.safeModuloInt((self).f36, (self).f28), 843)
        (globalState).f5 = (self).f27
        d_2784_v113_: _dafny.Array
        nw441_ = _dafny.Array(None, 1)
        def iife259_():
            coll105_ = _dafny.Map()
            compr_105_: int
            for compr_105_ in _dafny.IntegerRange(-821, 405):
                d_2785_v112_: int = compr_105_
                if ((-821) <= (d_2785_v112_)) and ((d_2785_v112_) < (405)):
                    coll105_[default__.safeDivisionInt(d_2785_v112_, (self).f29)] = (d_2761_v93_)[default__.safeIndex(735, (d_2761_v93_).length(0))]
            return _dafny.Map(coll105_)
        nw441_[int(0)] = iife259_()
        
        d_2784_v113_ = nw441_
        index423_ = default__.safeIndex(441, (d_2784_v113_).length(0))
        (d_2784_v113_)[index423_] = _dafny.Map({(self).f29: (self).f36})
        d_2786_v114_: _dafny.Map
        d_2786_v114_ = _dafny.Map({(0) - ((d_2761_v93_)[default__.safeIndex(735, (d_2761_v93_).length(0))]): len(_dafny.SeqWithoutIsStrInference([d_2762_v94_ for d_2787_i9_ in range(default__.abs(529))]))})
        index424_ = default__.safeIndex(441, (d_2784_v113_).length(0))
        rhs432_ = self.f35
        rhs433_ = (d_2786_v114_) | ((d_2786_v114_).set((self).f29, (self).f28))
        rhs434_ = self.f35
        rhs435_ = (d_2786_v114_) | (d_2786_v114_)
        lhs362_ = globalState
        lhs363_ = d_2784_v113_
        lhs364_ = default__.safeIndex(441, (d_2784_v113_).length(0))
        lhs365_ = globalState
        lhs362_.f5 = rhs432_
        lhs363_[lhs364_] = rhs433_
        lhs365_.f14 = rhs434_
        d_2786_v114_ = rhs435_
        d_2788_v115_: D0
        d_2788_v115_ = D0_DC0((0) - (len(d_2766_v98_)))
        source42_ = d_2788_v115_
        if source42_.is_DC1:
            d_2789___mcc_h10_ = source42_.cf1
            d_2790_cf1_ = d_2789___mcc_h10_
            d_2791_v116_: _dafny.Map
            d_2791_v116_ = _dafny.Map({510: (self).f36})
            d_2792_v117_: _dafny.Array
            nw442_ = _dafny.Array(_dafny.CodePoint('D'), 4)
            d_2792_v117_ = nw442_
            index425_ = default__.safeIndex(330, (d_2792_v117_).length(0))
            (d_2792_v117_)[index425_] = d_2762_v94_
            d_2793_v118_: _dafny.Map
            d_2793_v118_ = _dafny.Map({not(self.f35): (d_2761_v93_)[default__.safeIndex(735, (d_2761_v93_).length(0))]})
            d_2794_v119_: _dafny.Set
            d_2794_v119_ = _dafny.Set({(d_2765_v97_).cardinality, len(d_2793_v118_)})
            d_2795_v120_: _dafny.Map
            d_2795_v120_ = _dafny.Map({True: d_2794_v119_})
            d_2796_v121_: _dafny.MultiSet
            d_2796_v121_ = _dafny.MultiSet([(self).f28, -319])
            index426_ = default__.safeIndex(735, (d_2761_v93_).length(0))
            index427_ = default__.safeIndex(330, (d_2792_v117_).length(0))
            rhs436_ = (d_2793_v118_) | ((d_2793_v118_).set(default__.fm0(d_2763_v95_, d_2790_cf1_, d_2790_cf1_, d_2794_v119_, globalState), (self).f29))
            rhs437_ = len(((d_2795_v120_)[(self).f27] if ((self).f27) in (d_2795_v120_) else _dafny.Set({len(d_2768_v100_)})))
            rhs438_ = d_2791_v116_
            rhs439_ = (((self).f29) * (len(d_2763_v95_))) * (((d_2796_v121_)[(d_2761_v93_)[default__.safeIndex(735, (d_2761_v93_).length(0))]] if ((d_2761_v93_)[default__.safeIndex(735, (d_2761_v93_).length(0))]) in (d_2796_v121_) else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bpbgmwg")))))
            rhs440_ = (d_2763_v95_)[default__.safeIndex((self).f36, len(d_2763_v95_))]
            lhs366_ = globalState
            lhs367_ = globalState
            lhs368_ = d_2761_v93_
            lhs369_ = default__.safeIndex(735, (d_2761_v93_).length(0))
            lhs370_ = d_2792_v117_
            lhs371_ = default__.safeIndex(330, (d_2792_v117_).length(0))
            lhs366_.f13 = rhs436_
            lhs367_.f2 = rhs437_
            d_2791_v116_ = rhs438_
            lhs368_[lhs369_] = rhs439_
            lhs370_[lhs371_] = rhs440_
            index428_ = default__.safeIndex(735, (d_2761_v93_).length(0))
            (d_2761_v93_)[index428_] = default__.safeDivisionInt(default__.safeModuloInt((self).f29, (d_2761_v93_)[default__.safeIndex(735, (d_2761_v93_).length(0))]), (self).f29)
            d_2797_v122_: _dafny.Array
            nw443_ = _dafny.Array(_dafny.Map({}), 20)
            d_2797_v122_ = nw443_
            d_2798_v123_: _dafny.Map
            d_2798_v123_ = _dafny.Map({d_2771_v103_: d_2765_v97_})
            index429_ = default__.safeIndex(689, (d_2797_v122_).length(0))
            (d_2797_v122_)[index429_] = (d_2798_v123_).set(d_2771_v103_, d_2765_v97_)
            index430_ = default__.safeIndex(689, (d_2797_v122_).length(0))
            (d_2797_v122_)[index430_] = (d_2798_v123_) | (d_2798_v123_)
            d_2761_v93_ = d_2761_v93_
        elif source42_.is_DC2:
            d_2799___mcc_h11_ = source42_.cf2
            d_2800___mcc_h12_ = source42_.cf3
            d_2801_cf3_ = d_2800___mcc_h12_
            d_2802_cf2_ = d_2799___mcc_h11_
            d_2803_v124_: _dafny.Seq
            d_2803_v124_ = _dafny.SeqWithoutIsStrInference([d_2761_v93_, d_2761_v93_, d_2761_v93_])
            d_2804_v125_: T0
            nw444_ = C3()
            nw444_.ctor__((self).f27, self.f35)
            d_2804_v125_ = nw444_
            d_2805_v126_: _dafny.Map
            d_2805_v126_ = _dafny.Map({(self).f27: d_2804_v125_})
            d_2806_v127_: _dafny.Set
            d_2806_v127_ = _dafny.Set({(self).f28, len(d_2805_v126_)})
            d_2807_v128_: _dafny.Seq
            d_2807_v128_ = _dafny.SeqWithoutIsStrInference([d_2766_v98_, (d_2766_v98_ if (d_2804_v125_).f27 else _dafny.SeqWithoutIsStrInference([(d_2804_v125_).f27, (d_2804_v125_).f27]))])
            rhs441_ = (d_2801_cf3_) * (len(d_2806_v127_))
            rhs442_ = len(d_2807_v128_)
            rhs443_ = (_dafny.SeqWithoutIsStrInference([d_2761_v93_])) + ((d_2803_v124_) + (d_2803_v124_))
            rhs444_ = d_2801_cf3_
            lhs372_ = globalState
            lhs373_ = globalState
            lhs374_ = globalState
            lhs372_.f2 = rhs441_
            lhs373_.f6 = rhs442_
            d_2803_v124_ = rhs443_
            lhs374_.f6 = rhs444_
            d_2808_v129_: _dafny.Map
            d_2808_v129_ = _dafny.Map({not(self.f35): (d_2804_v125_).f27})
            d_2809_v130_: D7
            d_2809_v130_ = D7_DC17((self).f28, d_2802_cf2_, (self).f27)
            d_2810_v131_: _dafny.Map
            d_2810_v131_ = _dafny.Map({d_2808_v129_: (d_2809_v130_).cf26})
            d_2811_v132_: C6
            nw445_ = C6()
            nw445_.ctor__(len(d_2810_v131_), (d_2761_v93_)[default__.safeIndex(735, (d_2761_v93_).length(0))], (self).f27)
            d_2811_v132_ = nw445_
            d_2812_v133_: _dafny.Map
            d_2812_v133_ = _dafny.Map({d_2811_v132_: (d_2804_v125_).f27})
            d_2813_v134_: _dafny.Seq
            d_2813_v134_ = _dafny.SeqWithoutIsStrInference([d_2763_v95_, d_2763_v95_, d_2763_v95_, _dafny.SeqWithoutIsStrInference([d_2762_v94_ for d_2814_i10_ in range(default__.abs(621))]), _dafny.SeqWithoutIsStrInference([d_2762_v94_ for d_2815_i11_ in range(default__.abs(183))])])
            d_2816_v135_: _dafny.Seq
            d_2816_v135_ = _dafny.SeqWithoutIsStrInference([d_2804_v125_])
            d_2817_v136_: _dafny.MultiSet
            d_2817_v136_ = _dafny.MultiSet([d_2771_v103_])
            rhs445_ = (_dafny.Map({d_2811_v132_: (d_2804_v125_).f27}) if default__.fm0(d_2763_v95_, (self).f27, (self).f27, d_2806_v127_, globalState) else d_2812_v133_)
            rhs446_ = default__.fm0((d_2813_v134_)[default__.safeIndex((self).f36, len(d_2813_v134_))], False, (len((d_2816_v135_).set(default__.safeIndex(len(d_2766_v98_), len(d_2816_v135_)), d_2804_v125_))) >= ((d_2761_v93_)[default__.safeIndex(735, (d_2761_v93_).length(0))]), d_2806_v127_, globalState)
            rhs447_ = (d_2817_v136_).ispropersubset(d_2817_v136_)
            rhs448_ = ((d_2761_v93_)[default__.safeIndex(735, (d_2761_v93_).length(0))] if self.f35 else 538)
            lhs375_ = globalState
            lhs376_ = globalState
            lhs377_ = globalState
            d_2812_v133_ = rhs445_
            lhs375_.f14 = rhs446_
            lhs376_.f12 = rhs447_
            lhs377_.f6 = rhs448_
            (globalState).f5 = not(True)
            (globalState).f6 = d_2802_cf2_
        elif source42_.is_DC0:
            d_2818___mcc_h13_ = source42_.cf0
            d_2819_cf0_ = d_2818___mcc_h13_
            if ((0) - ((self).f29)) != (len(d_2768_v100_)):
                d_2820_v137_: _dafny.Set
                d_2820_v137_ = _dafny.Set({(self).f28})
                d_2821_v138_: _dafny.Map
                d_2821_v138_ = _dafny.Map({default__.fm0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mvnc")), (self).f27, (self).f27, d_2820_v137_, globalState): (self).f27})
                (globalState).f2 = default__.fm1(self.f35, d_2766_v98_, default__.fm0(d_2763_v95_, (self).f27, ((d_2821_v138_)[(self).f27] if ((self).f27) in (d_2821_v138_) else (self).f27), d_2820_v137_, globalState), globalState)
                d_2819_cf0_ = (default__.fm45(globalState)).cf102
                d_2822_v139_: _dafny.Array
                nw446_ = _dafny.Array(None, 23)
                nw446_[int(0)] = d_2761_v93_
                nw446_[int(1)] = d_2761_v93_
                nw446_[int(2)] = d_2761_v93_
                nw446_[int(3)] = d_2761_v93_
                nw446_[int(4)] = d_2761_v93_
                nw446_[int(5)] = d_2761_v93_
                nw446_[int(6)] = d_2761_v93_
                nw446_[int(7)] = (d_2761_v93_ if (self).f27 else d_2761_v93_)
                nw446_[int(8)] = d_2761_v93_
                nw446_[int(9)] = d_2761_v93_
                nw446_[int(10)] = d_2761_v93_
                nw446_[int(11)] = d_2761_v93_
                nw446_[int(12)] = d_2761_v93_
                nw446_[int(13)] = d_2761_v93_
                nw446_[int(14)] = d_2761_v93_
                nw446_[int(15)] = d_2761_v93_
                nw446_[int(16)] = d_2761_v93_
                nw446_[int(17)] = d_2761_v93_
                nw446_[int(18)] = d_2761_v93_
                nw446_[int(19)] = d_2761_v93_
                nw446_[int(20)] = d_2761_v93_
                nw446_[int(21)] = d_2761_v93_
                nw446_[int(22)] = (d_2761_v93_ if self.f35 else d_2761_v93_)
                d_2822_v139_ = nw446_
                index431_ = default__.safeIndex(704, (d_2822_v139_).length(0))
                (d_2822_v139_)[index431_] = d_2761_v93_
                d_2823_v140_: _dafny.Seq
                d_2823_v140_ = _dafny.SeqWithoutIsStrInference([d_2761_v93_, d_2761_v93_])
                d_2824_v141_: _dafny.Array
                d_2824_v141_ = (d_2823_v140_)[default__.safeIndex((self).f29, len(d_2823_v140_))]
                index432_ = default__.safeIndex(704, (d_2822_v139_).length(0))
                rhs449_ = d_2761_v93_
                rhs450_ = default__.safeDivisionInt((self).f29, len((D6_DC13(d_2766_v98_)).cf21))
                rhs451_ = (d_2824_v141_)
                rhs452_ = (self).f27
                lhs378_ = globalState
                lhs379_ = d_2822_v139_
                lhs380_ = default__.safeIndex(704, (d_2822_v139_).length(0))
                lhs381_ = globalState
                d_2761_v93_ = rhs449_
                lhs378_.f2 = rhs450_
                lhs379_[lhs380_] = rhs451_
                lhs381_.f25 = rhs452_
                (globalState).f5 = (self).f27
                index433_ = default__.safeIndex(220, (d_2771_v103_).length(0))
                (d_2771_v103_)[index433_] = not((len(d_2766_v98_)) < (d_2819_cf0_))
                index434_ = default__.safeIndex(220, (d_2771_v103_).length(0))
                rhs453_ = (_dafny.SeqWithoutIsStrInference([d_2762_v94_ for d_2825_i12_ in range(default__.abs(992))])) + ((d_2763_v95_) + (d_2763_v95_))
                rhs454_ = (d_2766_v98_)[default__.safeIndex((0) - (-608), len(d_2766_v98_))]
                lhs382_ = d_2771_v103_
                lhs383_ = default__.safeIndex(220, (d_2771_v103_).length(0))
                d_2763_v95_ = rhs453_
                lhs382_[lhs383_] = rhs454_
            elif True:
                index435_ = default__.safeIndex(735, (d_2761_v93_).length(0))
                (d_2761_v93_)[index435_] = (0) - (d_2819_cf0_)
                d_2766_v98_ = d_2766_v98_
                d_2826_v142_: _dafny.Seq
                d_2826_v142_ = _dafny.SeqWithoutIsStrInference([d_2768_v100_])
                d_2827_v143_: _dafny.Array
                nw447_ = _dafny.Array(None, 7)
                nw447_[int(0)] = default__.fm67((0) - (-346), globalState)
                nw447_[int(1)] = d_2826_v142_
                nw447_[int(2)] = d_2826_v142_
                nw447_[int(3)] = _dafny.SeqWithoutIsStrInference([d_2768_v100_])
                nw447_[int(4)] = d_2826_v142_
                nw447_[int(5)] = d_2826_v142_
                nw447_[int(6)] = _dafny.SeqWithoutIsStrInference([d_2768_v100_])
                d_2827_v143_ = nw447_
                d_2828_v144_: _dafny.Seq
                d_2828_v144_ = _dafny.SeqWithoutIsStrInference([d_2826_v142_])
                index436_ = default__.safeIndex(890, (d_2827_v143_).length(0))
                (d_2827_v143_)[index436_] = (d_2828_v144_)[default__.safeIndex(d_2819_cf0_, len(d_2828_v144_))]
                index437_ = default__.safeIndex(890, (d_2827_v143_).length(0))
                (d_2827_v143_)[index437_] = (((_dafny.SeqWithoutIsStrInference([d_2768_v100_])) + (_dafny.SeqWithoutIsStrInference([d_2768_v100_, d_2768_v100_])) if (self).f27 else d_2826_v142_)).set(default__.safeIndex(len(d_2766_v98_), len(((_dafny.SeqWithoutIsStrInference([d_2768_v100_])) + (_dafny.SeqWithoutIsStrInference([d_2768_v100_, d_2768_v100_])) if (self).f27 else d_2826_v142_))), (d_2768_v100_) + (d_2768_v100_))
                d_2829_v145_: _dafny.Map
                d_2829_v145_ = _dafny.Map({True: True})
                d_2830_v146_: _dafny.MultiSet
                d_2830_v146_ = _dafny.MultiSet([d_2829_v145_, _dafny.Map({(self).f27: (self).f27})])
                d_2831_v147_: D29
                d_2831_v147_ = D29_DC69(_dafny.Map({d_2762_v94_: (self).f28}))
                d_2832_v148_: _dafny.Map
                d_2832_v148_ = _dafny.Map({(self).f27: (self).f36})
                rhs455_ = self.f35
                rhs456_ = d_2763_v95_
                rhs457_ = (d_2830_v146_) == (_dafny.MultiSet([default__.fm60(d_2831_v147_, d_2832_v148_, globalState)]))
                lhs384_ = globalState
                lhs384_.f5 = rhs455_
                d_2763_v95_ = rhs456_
                r3 = rhs457_
                (globalState).f6 = (self).f28
            d_2833_v149_: _dafny.Map
            d_2833_v149_ = _dafny.Map({(self).f36: d_2763_v95_})
            index438_ = default__.safeIndex(735, (d_2761_v93_).length(0))
            rhs458_ = ((self).f28) * ((self).f29)
            rhs459_ = (len((_dafny.SeqWithoutIsStrInference([d_2763_v95_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nawoumcss")), d_2763_v95_, ((d_2833_v149_)[(self).f36] if ((self).f36) in (d_2833_v149_) else d_2763_v95_)]) if self.f35 else _dafny.SeqWithoutIsStrInference([d_2763_v95_ for d_2834_i13_ in range(default__.abs(282))])))) >= (default__.safeModuloInt((d_2761_v93_)[default__.safeIndex(735, (d_2761_v93_).length(0))], (d_2761_v93_)[default__.safeIndex(735, (d_2761_v93_).length(0))]))
            rhs460_ = not(self.f35)
            lhs385_ = d_2761_v93_
            lhs386_ = default__.safeIndex(735, (d_2761_v93_).length(0))
            lhs387_ = globalState
            lhs388_ = globalState
            lhs385_[lhs386_] = rhs458_
            lhs387_.f14 = rhs459_
            lhs388_.f25 = rhs460_
            d_2835_v150_: _dafny.Map
            d_2835_v150_ = _dafny.Map({(self).f28: ((0) - (d_2819_cf0_) if (self).f27 else (self).f29)})
            def iife260_():
                coll106_ = _dafny.Map()
                compr_106_: int
                for compr_106_ in ((d_2768_v100_) + (d_2768_v100_)).Elements:
                    d_2836_v151_: int = compr_106_
                    if (d_2836_v151_) in ((d_2768_v100_) + (d_2768_v100_)):
                        coll106_[(d_2836_v151_) + ((0) - ((self).f28))] = (self).f29
                return _dafny.Map(coll106_)
            d_2835_v150_ = iife260_()
            
            d_2837_v152_: _dafny.Map
            d_2837_v152_ = _dafny.Map({d_2771_v103_: default__.safeModuloInt((self).f29, len(d_2763_v95_))})
            d_2837_v152_ = (d_2837_v152_).set(d_2771_v103_, d_2819_cf0_)
        elif True:
            d_2838___mcc_h14_ = source42_.cf4
            d_2839_cf4_ = d_2838___mcc_h14_
            index439_ = default__.safeIndex(353, (d_2771_v103_).length(0))
            (d_2771_v103_)[index439_] = self.f35
            d_2840_v153_: D5
            d_2840_v153_ = D5_DC12(not(self.f35), (self).f28)
            index440_ = default__.safeIndex(353, (d_2771_v103_).length(0))
            (d_2771_v103_)[index440_] = (d_2840_v153_).cf19
            d_2841_v154_: C2
            nw448_ = C2()
            nw448_.ctor__(False, (d_2771_v103_)[default__.safeIndex(353, (d_2771_v103_).length(0))])
            d_2841_v154_ = nw448_
            (globalState).f2 = (d_2761_v93_)[default__.safeIndex(735, (d_2761_v93_).length(0))]
            index441_ = default__.safeIndex(735, (d_2761_v93_).length(0))
            (d_2761_v93_)[index441_] = (self).f36
        d_2842_v155_: _dafny.Array
        nw449_ = _dafny.Array(_dafny.Array(None, 0), 14)
        d_2842_v155_ = nw449_
        r0 = d_2842_v155_
        r1 = d_2763_v95_
        d_2843_v156_: _dafny.Seq
        d_2843_v156_ = _dafny.SeqWithoutIsStrInference([d_2761_v93_, d_2761_v93_, d_2761_v93_])
        d_2844_v157_: _dafny.Map
        d_2844_v157_ = _dafny.Map({d_2762_v94_: d_2761_v93_})
        d_2845_v158_: _dafny.Array
        nw450_ = _dafny.Array(None, 12)
        nw450_[int(0)] = d_2761_v93_
        nw450_[int(1)] = (d_2843_v156_)[default__.safeIndex((self).f29, len(d_2843_v156_))]
        nw450_[int(2)] = d_2761_v93_
        nw450_[int(3)] = d_2761_v93_
        nw450_[int(4)] = d_2761_v93_
        nw450_[int(5)] = d_2761_v93_
        nw450_[int(6)] = d_2761_v93_
        nw450_[int(7)] = d_2761_v93_
        nw450_[int(8)] = d_2761_v93_
        nw450_[int(9)] = d_2761_v93_
        nw450_[int(10)] = d_2761_v93_
        nw450_[int(11)] = ((d_2844_v157_)[d_2762_v94_] if (d_2762_v94_) in (d_2844_v157_) else d_2761_v93_)
        d_2845_v158_ = nw450_
        r2 = d_2845_v158_
        r3 = ((495) + ((self).f36)) >= ((self).f29)
        return r0, r1, r2, r3

    def m10(self, p0, p1, globalState):
        d_2846_v0_: _dafny.Array
        nw451_ = _dafny.Array(D29.default()(), 17)
        d_2846_v0_ = nw451_
        d_2847_v1_: str
        d_2847_v1_ = _dafny.CodePoint('d')
        index442_ = default__.safeIndex(152, (d_2846_v0_).length(0))
        (d_2846_v0_)[index442_] = D29_DC69(_dafny.Map({d_2847_v1_: (self).f29}))
        d_2848_v2_: _dafny.Map
        d_2848_v2_ = _dafny.Map({d_2847_v1_: 769})
        d_2849_v3_: D29
        d_2849_v3_ = D29_DC69((d_2848_v2_).set(d_2847_v1_, p1))
        index443_ = default__.safeIndex(152, (d_2846_v0_).length(0))
        rhs461_ = d_2849_v3_
        rhs462_ = p0
        lhs389_ = d_2846_v0_
        lhs390_ = default__.safeIndex(152, (d_2846_v0_).length(0))
        lhs391_ = globalState
        lhs389_[lhs390_] = rhs461_
        lhs391_.f12 = rhs462_
        if (self).f27:
            d_2850_v4_: _dafny.Seq
            d_2850_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mdeccx"))
            d_2851_v5_: _dafny.Array
            nw452_ = _dafny.Array(int(0), 5)
            d_2851_v5_ = nw452_
            d_2852_v6_: _dafny.Map
            d_2852_v6_ = _dafny.Map({d_2851_v5_: (d_2850_v4_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nmevtw")))})
            d_2850_v4_ = ((d_2852_v6_)[d_2851_v5_] if (d_2851_v5_) in (d_2852_v6_) else (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "brbn"))) + (d_2850_v4_))
            d_2853_v7_: _dafny.Seq
            d_2853_v7_ = _dafny.SeqWithoutIsStrInference([((self).f36) + ((self).f28), (self).f29])
            d_2854_v8_: _dafny.Array
            def lambda127_(d_2855_i0_):
                return D20_DC50()

            init74_ = lambda127_
            nw453_ = _dafny.Array(None, 19)
            for i0_74_ in range(nw453_.length(0)):
                nw453_[i0_74_] = init74_(i0_74_)
            d_2854_v8_ = nw453_
            rhs463_ = d_2853_v7_
            rhs464_ = (p1) + (default__.fm1(not(p0), _dafny.SeqWithoutIsStrInference([(self).f27]), (self).f27, globalState))
            rhs465_ = ((d_2850_v4_ if p0 else d_2850_v4_)) + (d_2850_v4_)
            rhs466_ = d_2854_v8_
            lhs392_ = globalState
            d_2853_v7_ = rhs463_
            lhs392_.f6 = rhs464_
            d_2850_v4_ = rhs465_
            d_2854_v8_ = rhs466_
            d_2856_v9_: _dafny.Set
            d_2856_v9_ = _dafny.Set({self.f35, (self).f27, self.f35})
            d_2857_v10_: _dafny.Set
            d_2857_v10_ = _dafny.Set({(self).f28})
            d_2858_v11_: _dafny.Seq
            d_2858_v11_ = _dafny.SeqWithoutIsStrInference([d_2856_v9_])
            d_2859_v12_: _dafny.Seq
            d_2859_v12_ = _dafny.SeqWithoutIsStrInference([(self).f27, self.f35])
            d_2860_v13_: _dafny.Array
            nw454_ = _dafny.Array(None, 1)
            nw454_[int(0)] = self.f35
            d_2860_v13_ = nw454_
            d_2861_v14_: _dafny.Seq
            d_2861_v14_ = _dafny.SeqWithoutIsStrInference([d_2860_v13_])
            d_2862_v15_: _dafny.Map
            d_2862_v15_ = _dafny.Map({(self).f29: d_2861_v14_})
            d_2863_v16_: _dafny.Map
            d_2863_v16_ = _dafny.Map({((d_2862_v15_)[(self).f28] if ((self).f28) in (d_2862_v15_) else d_2861_v14_): d_2856_v9_})
            d_2864_v17_: _dafny.Array
            nw455_ = _dafny.Array(None, 27)
            nw455_[int(0)] = d_2856_v9_
            nw455_[int(1)] = d_2856_v9_
            nw455_[int(2)] = d_2856_v9_
            nw455_[int(3)] = default__.fm26(globalState)
            nw455_[int(4)] = d_2856_v9_
            nw455_[int(5)] = _dafny.Set({(self).f27, p0, (self).f27, (self).f27})
            nw455_[int(6)] = (_dafny.Set({default__.fm0(d_2850_v4_, (self).f27, p0, d_2857_v10_, globalState)})) | (d_2856_v9_)
            nw455_[int(7)] = d_2856_v9_
            nw455_[int(8)] = (d_2856_v9_) | (d_2856_v9_)
            nw455_[int(9)] = (d_2856_v9_).intersection(d_2856_v9_)
            nw455_[int(10)] = _dafny.Set({p0})
            nw455_[int(11)] = (_dafny.Set({self.f35})) - (_dafny.Set({self.f35, (self).f27, False, p0}))
            nw455_[int(12)] = (d_2856_v9_).intersection(d_2856_v9_)
            nw455_[int(13)] = (d_2856_v9_) | (d_2856_v9_)
            nw455_[int(14)] = _dafny.Set({self.f35, (self).f27, self.f35, p0, (self).f27})
            nw455_[int(15)] = d_2856_v9_
            nw455_[int(16)] = d_2856_v9_
            nw455_[int(17)] = d_2856_v9_
            nw455_[int(18)] = d_2856_v9_
            nw455_[int(19)] = (d_2858_v11_)[default__.safeIndex(default__.fm1(self.f35, d_2859_v12_, True, globalState), len(d_2858_v11_))]
            nw455_[int(20)] = d_2856_v9_
            nw455_[int(21)] = d_2856_v9_
            nw455_[int(22)] = d_2856_v9_
            nw455_[int(23)] = (_dafny.Set({p0})) | (d_2856_v9_)
            nw455_[int(24)] = ((d_2863_v16_)[_dafny.SeqWithoutIsStrInference([d_2860_v13_, d_2860_v13_, d_2860_v13_, d_2860_v13_, d_2860_v13_])] if (_dafny.SeqWithoutIsStrInference([d_2860_v13_, d_2860_v13_, d_2860_v13_, d_2860_v13_, d_2860_v13_])) in (d_2863_v16_) else d_2856_v9_)
            nw455_[int(25)] = d_2856_v9_
            nw455_[int(26)] = d_2856_v9_
            d_2864_v17_ = nw455_
            d_2864_v17_ = d_2864_v17_
            (globalState).f6 = (self).f36
            index444_ = default__.safeIndex(203, (d_2860_v13_).length(0))
            (d_2860_v13_)[index444_] = False
            index445_ = default__.safeIndex(203, (d_2860_v13_).length(0))
            (d_2860_v13_)[index445_] = (self).f27
        elif True:
            d_2865_v18_: _dafny.Seq
            d_2865_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rqcea"))
            d_2866_v19_: _dafny.Set
            d_2866_v19_ = _dafny.Set({(self).f29, 282})
            d_2867_v20_: _dafny.Map
            d_2867_v20_ = _dafny.Map({(self).f29: p0})
            d_2868_v21_: _dafny.Array
            nw456_ = _dafny.Array(None, 11)
            nw456_[int(0)] = self.f35
            nw456_[int(1)] = self.f35
            nw456_[int(2)] = self.f35
            nw456_[int(3)] = (self).f27
            nw456_[int(4)] = self.f35
            nw456_[int(5)] = default__.fm0(d_2865_v18_, p0, False, d_2866_v19_, globalState)
            nw456_[int(6)] = p0
            nw456_[int(7)] = p0
            nw456_[int(8)] = ((d_2867_v20_)[p1] if (p1) in (d_2867_v20_) else p0)
            nw456_[int(9)] = False
            nw456_[int(10)] = (self).f27
            d_2868_v21_ = nw456_
            d_2869_v22_: _dafny.Seq
            d_2869_v22_ = _dafny.SeqWithoutIsStrInference([d_2868_v21_])
            d_2870_v23_: _dafny.Map
            d_2870_v23_ = _dafny.Map({d_2847_v1_: p0})
            d_2871_v24_: _dafny.Array
            nw457_ = _dafny.Array(None, 26)
            nw457_[int(0)] = d_2868_v21_
            nw457_[int(1)] = d_2868_v21_
            nw457_[int(2)] = d_2868_v21_
            nw457_[int(3)] = d_2868_v21_
            nw457_[int(4)] = d_2868_v21_
            nw457_[int(5)] = (d_2869_v22_)[default__.safeIndex(len(d_2870_v23_), len(d_2869_v22_))]
            nw457_[int(6)] = d_2868_v21_
            nw457_[int(7)] = d_2868_v21_
            nw457_[int(8)] = d_2868_v21_
            nw457_[int(9)] = d_2868_v21_
            nw457_[int(10)] = d_2868_v21_
            nw457_[int(11)] = d_2868_v21_
            nw457_[int(12)] = d_2868_v21_
            nw457_[int(13)] = d_2868_v21_
            nw457_[int(14)] = d_2868_v21_
            nw457_[int(15)] = d_2868_v21_
            nw457_[int(16)] = d_2868_v21_
            nw457_[int(17)] = d_2868_v21_
            nw457_[int(18)] = d_2868_v21_
            nw457_[int(19)] = d_2868_v21_
            nw457_[int(20)] = d_2868_v21_
            nw457_[int(21)] = d_2868_v21_
            nw457_[int(22)] = d_2868_v21_
            nw457_[int(23)] = d_2868_v21_
            nw457_[int(24)] = d_2868_v21_
            nw457_[int(25)] = d_2868_v21_
            d_2871_v24_ = nw457_
            d_2872_v25_: _dafny.MultiSet
            d_2872_v25_ = _dafny.MultiSet([self.f35])
            d_2873_v26_: _dafny.Map
            d_2873_v26_ = _dafny.Map({d_2871_v24_: d_2872_v25_})
            d_2874_v27_: D0
            d_2874_v27_ = D0_DC2((((d_2873_v26_)[d_2871_v24_] if (d_2871_v24_) in (d_2873_v26_) else d_2872_v25_)).cardinality, p1)
            d_2875_v28_: D0
            d_2875_v28_ = D0_DC3(d_2874_v27_)
            d_2876_v29_: bool
            d_2877_v30_: _dafny.Seq
            out79_: bool
            out80_: _dafny.Seq
            out79_, out80_ = default__.m0(d_2875_v28_, default__.fm46(d_2872_v25_, (self).f27, globalState), globalState)
            d_2876_v29_ = out79_
            d_2877_v30_ = out80_
            d_2878_v31_: _dafny.Map
            d_2878_v31_ = _dafny.Map({(self).f36: _dafny.CodePoint('g')})
            d_2879_v32_: _dafny.Set
            d_2879_v32_ = _dafny.Set({d_2878_v31_})
            d_2880_v33_: _dafny.Set
            d_2880_v33_ = d_2879_v32_
            d_2881_v34_: _dafny.Seq
            d_2881_v34_ = _dafny.SeqWithoutIsStrInference([d_2880_v33_])
            d_2882_v35_: D30
            d_2882_v35_ = D30_DC75((self).f36, d_2865_v18_, d_2881_v34_, d_2868_v21_)
            d_2883_v36_: _dafny.Map
            d_2883_v36_ = _dafny.Map({d_2882_v35_: not(False)})
            d_2883_v36_ = (d_2883_v36_).set(d_2882_v35_, (self).f27)
            d_2884_v37_: D23
            d_2884_v37_ = D23_DC60(d_2876_v29_, d_2847_v1_)
            (globalState).f21 = (d_2884_v37_).cf91
            d_2885_v38_: _dafny.Seq
            d_2885_v38_ = _dafny.SeqWithoutIsStrInference([p0, d_2876_v29_, self.f35])
            d_2886_v39_: D20
            d_2886_v39_ = D20_DC50()
            d_2887_v40_: D20
            d_2887_v40_ = D20_DC51(d_2886_v39_)
            pat_let_tv59_ = d_2886_v39_
            d_2888_v41_: _dafny.Set
            def iife261_(_pat_let77_0):
                def iife262_(d_2889_dt__update__tmp_h0_):
                    def iife263_(_pat_let78_0):
                        def iife264_(d_2890_dt__update_hcf74_h0_):
                            return D20_DC51(d_2890_dt__update_hcf74_h0_)
                        return iife264_(_pat_let78_0)
                    return iife263_(pat_let_tv59_)
                return iife262_(_pat_let77_0)
            d_2888_v41_ = _dafny.Set({_dafny.Map({len(d_2885_v38_): iife261_(d_2887_v40_)})})
            d_2888_v41_ = (d_2888_v41_) | (d_2888_v41_)
            d_2891_v42_: _dafny.Map
            d_2891_v42_ = _dafny.Map({p1: d_2877_v30_})
            d_2892_v43_: _dafny.MultiSet
            d_2892_v43_ = _dafny.MultiSet([(self).f29, (self).f28])
            d_2893_v44_: _dafny.Seq
            d_2893_v44_ = _dafny.SeqWithoutIsStrInference([default__.fm1(False, d_2885_v38_, self.f35, globalState), len(d_2867_v20_), (self).f36, 617, len(d_2877_v30_)])
            d_2894_v45_: _dafny.Map
            d_2894_v45_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iju")): ((d_2892_v43_)[p1] if (p1) in (d_2892_v43_) else len(d_2893_v44_))})
            d_2895_v46_: D22
            d_2895_v46_ = D22_DC58((self).f27, len(d_2894_v45_), len((d_2865_v18_).set(default__.safeIndex(len(d_2893_v44_), len(d_2865_v18_)), d_2847_v1_)))
            if ((d_2865_v18_) + (d_2865_v18_)) <= (((d_2891_v42_)[(d_2895_v46_).cf89] if ((d_2895_v46_).cf89) in (d_2891_v42_) else (self).fm3(globalState))):
                d_2896_v47_: C2
                nw458_ = C2()
                nw458_.ctor__((True if (self).f27 else self.f35), self.f35)
                d_2896_v47_ = nw458_
                (globalState).f14 = (_dafny.MultiSet([(self).f28])).ispropersubset(_dafny.MultiSet([p1, p1]))
                d_2897_v48_: _dafny.Array
                def lambda128_(d_2898_i1_):
                    return default__.safeModuloInt(d_2898_i1_, (self).f29)

                init75_ = lambda128_
                nw459_ = _dafny.Array(None, 4)
                for i0_75_ in range(nw459_.length(0)):
                    nw459_[i0_75_] = init75_(i0_75_)
                d_2897_v48_ = nw459_
                index446_ = default__.safeIndex(831, (d_2897_v48_).length(0))
                (d_2897_v48_)[index446_] = ((self).f29) + ((d_2896_v47_).fm9(116, globalState))
                index447_ = default__.safeIndex(831, (d_2897_v48_).length(0))
                (d_2897_v48_)[index447_] = (0) - (p1)
                (globalState).f5 = self.f35
                d_2899_v49_: _dafny.Seq
                d_2899_v49_ = _dafny.SeqWithoutIsStrInference([d_2877_v30_])
                d_2900_v50_: D13
                d_2900_v50_ = D13_DC27(d_2899_v49_)
                index448_ = default__.safeIndex(831, (d_2897_v48_).length(0))
                (d_2897_v48_)[index448_] = len((d_2900_v50_).cf37)
            elif True:
                d_2901_v51_: _dafny.Set
                d_2901_v51_ = _dafny.Set({self.f35, p0})
                d_2902_v53_: _dafny.MultiSet
                def iife265_():
                    coll107_ = _dafny.Set()
                    compr_107_: int
                    for compr_107_ in _dafny.IntegerRange(429, 278):
                        d_2903_v52_: int = compr_107_
                        if ((429) <= (d_2903_v52_)) and ((d_2903_v52_) < (278)):
                            coll107_ = coll107_.union(_dafny.Set([(d_2903_v52_) * (((d_2872_v25_)[False] if (False) in (d_2872_v25_) else (self).f29))]))
                    return _dafny.Set(coll107_)
                d_2902_v53_ = _dafny.MultiSet([_dafny.Set({default__.fm0(d_2865_v18_, p0, d_2876_v29_, iife265_()
                , globalState)})])
                d_2904_v54_: _dafny.Array
                nw460_ = _dafny.Array(None, 16)
                nw460_[int(0)] = (self).f28
                nw460_[int(1)] = ((self).f29) - (len(d_2901_v51_))
                nw460_[int(2)] = (p1) - ((self).f29)
                nw460_[int(3)] = p1
                nw460_[int(4)] = (self).f29
                nw460_[int(5)] = (d_2893_v44_)[default__.safeIndex(len(d_2885_v38_), len(d_2893_v44_))]
                nw460_[int(6)] = p1
                nw460_[int(7)] = ((d_2902_v53_).cardinality) - (len(_dafny.SeqWithoutIsStrInference([d_2847_v1_ for d_2905_i2_ in range(default__.abs(179))])))
                nw460_[int(8)] = (0) - (default__.safeDivisionInt((self).f28, (self).f29))
                nw460_[int(9)] = (self).f29
                nw460_[int(10)] = (self).f29
                nw460_[int(11)] = len(_dafny.Set({default__.fm1(p0, d_2885_v38_, self.f35, globalState), p1}))
                nw460_[int(12)] = p1
                nw460_[int(13)] = (self).f29
                nw460_[int(14)] = (self).f28
                nw460_[int(15)] = (self).f28
                d_2904_v54_ = nw460_
                index449_ = default__.safeIndex(855, (d_2904_v54_).length(0))
                (d_2904_v54_)[index449_] = (self).f29
                index450_ = default__.safeIndex(855, (d_2904_v54_).length(0))
                (d_2904_v54_)[index450_] = (((self).f28) - (len(d_2877_v30_))) * (p1)
                d_2906_v55_: _dafny.Map
                d_2906_v55_ = _dafny.Map({self.f35: (self).f28})
                d_2907_v56_: _dafny.Map
                d_2907_v56_ = _dafny.Map({(d_2892_v43_).cardinality: len((_dafny.Map({p0: ((d_2892_v43_)[(self).f28] if ((self).f28) in (d_2892_v43_) else (self).f28)})) | (d_2906_v55_))})
                d_2907_v56_ = (d_2907_v56_).set(699, (self).f29)
                d_2908_v57_: _dafny.Array
                nw461_ = _dafny.Array(_dafny.Map({}), 26)
                d_2908_v57_ = nw461_
                d_2909_v58_: _dafny.Map
                d_2909_v58_ = _dafny.Map({p0: not(p0)})
                index451_ = default__.safeIndex(137, (d_2908_v57_).length(0))
                (d_2908_v57_)[index451_] = d_2909_v58_
                d_2910_v59_: _dafny.Array
                nw462_ = _dafny.Array(_dafny.Seq({}), 6)
                d_2910_v59_ = nw462_
                index452_ = default__.safeIndex(851, (d_2910_v59_).length(0))
                (d_2910_v59_)[index452_] = d_2885_v38_
                d_2911_v60_: D23
                d_2911_v60_ = D23_DC61(d_2904_v54_, True, (self).f29)
                index453_ = default__.safeIndex(137, (d_2908_v57_).length(0))
                index454_ = default__.safeIndex(851, (d_2910_v59_).length(0))
                rhs467_ = default__.safeDivisionInt((self).f36, 169)
                rhs468_ = _dafny.Map({self.f35: (d_2911_v60_).cf94})
                rhs469_ = d_2885_v38_
                lhs393_ = globalState
                lhs394_ = d_2908_v57_
                lhs395_ = default__.safeIndex(137, (d_2908_v57_).length(0))
                lhs396_ = d_2910_v59_
                lhs397_ = default__.safeIndex(851, (d_2910_v59_).length(0))
                lhs393_.f6 = rhs467_
                lhs394_[lhs395_] = rhs468_
                lhs396_[lhs397_] = rhs469_
                d_2912_v61_: C3
                nw463_ = C3()
                nw463_.ctor__(p0, self.f35)
                d_2912_v61_ = nw463_
                def iife266_():
                    coll108_ = _dafny.Set()
                    compr_108_: int
                    for compr_108_ in _dafny.IntegerRange(254, -31):
                        d_2913_v62_: int = compr_108_
                        if ((254) <= (d_2913_v62_)) and ((d_2913_v62_) < (-31)):
                            coll108_ = coll108_.union(_dafny.Set([default__.safeModuloInt(d_2913_v62_, (d_2904_v54_)[default__.safeIndex(855, (d_2904_v54_).length(0))])]))
                    return _dafny.Set(coll108_)
                rhs470_ = d_2865_v18_
                rhs471_ = len(d_2906_v55_)
                rhs472_ = (d_2912_v61_ if (len(iife266_()
                )) > ((self).f29) else d_2912_v61_)
                lhs398_ = globalState
                d_2865_v18_ = rhs470_
                lhs398_.f6 = rhs471_
                d_2912_v61_ = rhs472_
                d_2914_v63_: _dafny.Array
                nw464_ = _dafny.Array(_dafny.Seq({}), 17)
                d_2914_v63_ = nw464_
                d_2915_v64_: _dafny.Seq
                d_2915_v64_ = _dafny.SeqWithoutIsStrInference([d_2904_v54_, d_2904_v54_])
                index455_ = default__.safeIndex(540, (d_2914_v63_).length(0))
                (d_2914_v63_)[index455_] = d_2915_v64_
                index456_ = default__.safeIndex(540, (d_2914_v63_).length(0))
                rhs473_ = d_2904_v54_
                rhs474_ = d_2915_v64_
                rhs475_ = not(((d_2865_v18_) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cwxi"))) if d_2912_v61_.f38 else p0))
                lhs399_ = d_2914_v63_
                lhs400_ = default__.safeIndex(540, (d_2914_v63_).length(0))
                lhs401_ = globalState
                d_2904_v54_ = rhs473_
                lhs399_[lhs400_] = rhs474_
                lhs401_.f12 = rhs475_
        d_2847_v1_ = d_2847_v1_
        d_2916_v65_: _dafny.Array
        nw465_ = _dafny.Array(None, 6)
        d_2916_v65_ = nw465_
        d_2917_v66_: _dafny.Map
        d_2917_v66_ = _dafny.Map({d_2916_v65_: self.f35})
        d_2918_v67_: _dafny.Array
        nw466_ = _dafny.Array(None, 2)
        nw466_[int(0)] = d_2917_v66_
        nw466_[int(1)] = (d_2917_v66_ if self.f35 else d_2917_v66_)
        d_2918_v67_ = nw466_
        index457_ = default__.safeIndex(447, (d_2918_v67_).length(0))
        (d_2918_v67_)[index457_] = _dafny.Map({d_2916_v65_: p0})
        index458_ = default__.safeIndex(447, (d_2918_v67_).length(0))
        (d_2918_v67_)[index458_] = (_dafny.Map({d_2916_v65_: p0})) | (d_2917_v66_)
        d_2919_v68_: _dafny.Map
        d_2919_v68_ = _dafny.Map({(self).f29: p0})
        d_2920_v69_: _dafny.Seq
        d_2920_v69_ = _dafny.SeqWithoutIsStrInference([d_2919_v68_])
        d_2921_v70_: _dafny.Seq
        d_2921_v70_ = _dafny.SeqWithoutIsStrInference([(self.f35) or ((self).f27), (self).f27, (self).f27, (d_2920_v69_) != (d_2920_v69_), self.f35])
        (globalState).f12 = not((d_2921_v70_)[default__.safeIndex((self).f28, len(d_2921_v70_))])
        d_2922_v71_: _dafny.MultiSet
        d_2922_v71_ = _dafny.MultiSet([(self).f29, (self).f28])
        d_2923_v72_: _dafny.Seq
        d_2923_v72_ = _dafny.SeqWithoutIsStrInference([len(d_2919_v68_)])
        hi17_ = (self).f28
        for d_2924_i3_ in range(default__.safeDivisionInt((self).f29, ((d_2922_v71_)[p1] if (p1) in (d_2922_v71_) else len(d_2923_v72_))), hi17_):
            d_2925_v73_: _dafny.MultiSet
            d_2925_v73_ = _dafny.MultiSet([d_2847_v1_, d_2847_v1_])
            (globalState).f25 = ((d_2925_v73_) != (d_2925_v73_) if (self).f27 else p0)
            d_2926_v74_: _dafny.MultiSet
            d_2926_v74_ = _dafny.MultiSet([False])
            d_2927_v75_: _dafny.Seq
            d_2927_v75_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(self).f27]), d_2926_v74_])
            d_2928_v76_: _dafny.Set
            d_2928_v76_ = _dafny.Set({len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wrielg"))).set(default__.safeIndex(len(d_2923_v72_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wrielg")))), d_2847_v1_))})
            d_2929_v77_: C4
            nw467_ = C4()
            nw467_.ctor__(((d_2927_v75_)[default__.safeIndex((self).f36, len(d_2927_v75_))]) | (_dafny.MultiSet([p0, (self).f27])), not(default__.fm0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vcwacxs")), self.f35, (self).f27, d_2928_v76_, globalState)))
            d_2929_v77_ = nw467_
            (globalState).f12 = (self).f27
            d_2930_v78_: _dafny.Seq
            d_2930_v78_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oidamohs"))
            d_2931_v79_: _dafny.Map
            d_2931_v79_ = _dafny.Map({(self).f36: len(d_2930_v78_)})
            d_2931_v79_ = d_2931_v79_

    @property
    def f36(self):
        return self._f36

class C8(T2, T0):
    def  __init__(self):
        self._f27: bool = False
        self._f32: bool = False
        self._f33: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    @property
    def f27(self):
        return self._f27
    def ctor__(self, f32, f33, f27):
        (self)._f32 = f32
        (self)._f33 = f33
        (self)._f27 = f27

    def fm4(self, p0, globalState):
        source43_ = D0_DC1((self).f27)
        if source43_.is_DC1:
            d_2932___mcc_h0_ = source43_.cf1
            d_2933_cf1_ = d_2932___mcc_h0_
            return (self).f32
        elif source43_.is_DC2:
            d_2934___mcc_h1_ = source43_.cf2
            d_2935___mcc_h2_ = source43_.cf3
            d_2936_cf3_ = d_2935___mcc_h2_
            d_2937_cf2_ = d_2934___mcc_h1_
            return (self).f32
        elif source43_.is_DC0:
            d_2938___mcc_h3_ = source43_.cf0
            d_2939_cf0_ = d_2938___mcc_h3_
            return (self).f32
        elif True:
            d_2940___mcc_h4_ = source43_.cf4
            d_2941_cf4_ = d_2940___mcc_h4_
            return (self).f32

    def fm5(self, globalState):
        return D0_DC2((0) - (((self).f33) * (294)), (len(_dafny.Map({(self).f32: D0_DC3(D0_DC1(True))}))) * ((self).f33))

    def fm2(self, p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "trfdrqixs"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_2942_i0_ in range(default__.abs(140))]))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "imhkwh"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_2943_i1_ in range(default__.abs(387))])))

    def fm3(self, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ntif"))

    def fm8(self, globalState):
        return (_dafny.Map({(self).f33: (0) - ((self).f33)})) == (_dafny.Map({(self).f33: (self).f33}))

    def m4(self, p0, p1, globalState):
        r0: bool = False
        r1: bool = False
        d_2944_v0_: _dafny.Array
        nw468_ = _dafny.Array(int(0), 15)
        d_2944_v0_ = nw468_
        d_2945_v1_: _dafny.Seq
        d_2945_v1_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f27: (self).f27})])
        d_2946_v2_: _dafny.Map
        d_2946_v2_ = _dafny.Map({(self).f32: (self).f32})
        d_2947_v3_: _dafny.Seq
        d_2947_v3_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.Map({p1: ((d_2946_v2_)[(self).f32] if ((self).f32) in (d_2946_v2_) else (self).f27)}), d_2946_v2_])])
        index459_ = default__.safeIndex(877, (d_2944_v0_).length(0))
        (d_2944_v0_)[index459_] = len(((d_2945_v1_).set(default__.safeIndex((self).f33, len(d_2945_v1_)), d_2946_v2_)) + ((d_2947_v3_)[default__.safeIndex((self).f33, len(d_2947_v3_))]))
        d_2948_v4_: _dafny.Array
        nw469_ = _dafny.Array(False, 11)
        d_2948_v4_ = nw469_
        index460_ = default__.safeIndex(78, (d_2948_v4_).length(0))
        (d_2948_v4_)[index460_] = (self).f27
        d_2949_v5_: _dafny.Seq
        d_2949_v5_ = _dafny.SeqWithoutIsStrInference([(self).f32, p1])
        d_2950_v6_: _dafny.MultiSet
        d_2950_v6_ = _dafny.MultiSet([p1])
        d_2951_v7_: _dafny.Seq
        d_2951_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vvqyorm"))
        d_2952_v8_: _dafny.Map
        d_2952_v8_ = _dafny.Map({(self).f32: d_2951_v7_})
        d_2953_v9_: _dafny.MultiSet
        d_2953_v9_ = _dafny.MultiSet([-327, (self).f33, (self).f33])
        d_2954_v10_: _dafny.Seq
        d_2954_v10_ = _dafny.SeqWithoutIsStrInference([d_2953_v9_])
        d_2955_v11_: _dafny.Seq
        d_2955_v11_ = _dafny.SeqWithoutIsStrInference([(d_2950_v6_).cardinality, (self).f33, len(d_2952_v8_), (self).f33, len((d_2954_v10_) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(_dafny.Map({_dafny.Set({len(_dafny.Map({(d_2950_v6_).cardinality: (self).f32})), (0) - ((self).f33)}): (self).f27}))]) for d_2956_i0_ in range(default__.abs(655))])))])
        index461_ = default__.safeIndex(877, (d_2944_v0_).length(0))
        index462_ = default__.safeIndex(78, (d_2948_v4_).length(0))
        rhs476_ = (self).f33
        rhs477_ = default__.fm1((self).f32, (_dafny.SeqWithoutIsStrInference([p1, p1])) + (d_2949_v5_), (self).f32, globalState)
        rhs478_ = (d_2955_v11_)[default__.safeIndex((self).f33, len(d_2955_v11_))]
        rhs479_ = (d_2955_v11_) <= (_dafny.SeqWithoutIsStrInference([(self).f33 for d_2957_i1_ in range(default__.abs(823))]))
        lhs402_ = globalState
        lhs403_ = d_2944_v0_
        lhs404_ = default__.safeIndex(877, (d_2944_v0_).length(0))
        lhs405_ = globalState
        lhs406_ = d_2948_v4_
        lhs407_ = default__.safeIndex(78, (d_2948_v4_).length(0))
        lhs402_.f2 = rhs476_
        lhs403_[lhs404_] = rhs477_
        lhs405_.f6 = rhs478_
        lhs406_[lhs407_] = rhs479_
        d_2958_v12_: str
        d_2958_v12_ = _dafny.CodePoint('f')
        if (d_2958_v12_) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wctmpoww"))):
            (globalState).f2 = default__.safeDivisionInt((self).f33, (default__.fm1((self).f27, d_2949_v5_, (d_2948_v4_)[default__.safeIndex(78, (d_2948_v4_).length(0))], globalState)) - ((self).f33))
            d_2959_v13_: T2
            nw470_ = C1()
            nw470_.ctor__((d_2948_v4_)[default__.safeIndex(78, (d_2948_v4_).length(0))])
            d_2959_v13_ = nw470_
            d_2959_v13_ = d_2959_v13_
            d_2960_v14_: _dafny.Map
            d_2960_v14_ = _dafny.Map({(self).f33: ((0) - ((self).f33)) * ((d_2944_v0_)[default__.safeIndex(877, (d_2944_v0_).length(0))])})
            d_2960_v14_ = (d_2960_v14_).set((d_2944_v0_)[default__.safeIndex(877, (d_2944_v0_).length(0))], (self).f33)
            d_2961_v15_: _dafny.Map
            d_2961_v15_ = _dafny.Map({(d_2951_v7_) + (d_2951_v7_): (self).f27})
            d_2961_v15_ = (d_2961_v15_).set(d_2951_v7_, (d_2959_v13_).f27)
            d_2962_v16_: D3
            d_2962_v16_ = D3_DC9((self).f27, (d_2944_v0_)[default__.safeIndex(877, (d_2944_v0_).length(0))], (self).f33)
            r1 = (((self).f33) * (default__.fm1((d_2962_v16_).cf14, _dafny.SeqWithoutIsStrInference([p1, not(not(True))]), False, globalState))) > ((0) - ((self).f33))
        elif True:
            (globalState).f14 = (self).f32
            d_2963_v17_: _dafny.Map
            d_2963_v17_ = _dafny.Map({d_2944_v0_: d_2958_v12_})
            d_2958_v12_ = ((d_2963_v17_)[d_2944_v0_] if (d_2944_v0_) in (d_2963_v17_) else d_2958_v12_)
            index463_ = default__.safeIndex(877, (d_2944_v0_).length(0))
            (d_2944_v0_)[index463_] = ((d_2944_v0_)[default__.safeIndex(877, (d_2944_v0_).length(0))]) + ((self).f33)
            index464_ = default__.safeIndex(877, (d_2944_v0_).length(0))
            (d_2944_v0_)[index464_] = ((0) - ((self).f33)) + ((self).f33)
            d_2964_v18_: _dafny.Set
            d_2964_v18_ = _dafny.Set({(d_2944_v0_)[default__.safeIndex(877, (d_2944_v0_).length(0))], (d_2944_v0_)[default__.safeIndex(877, (d_2944_v0_).length(0))], default__.fm1(True, _dafny.SeqWithoutIsStrInference([False]), (d_2948_v4_)[default__.safeIndex(78, (d_2948_v4_).length(0))], globalState), (d_2944_v0_)[default__.safeIndex(877, (d_2944_v0_).length(0))], (self).f33})
            index465_ = default__.safeIndex(78, (d_2948_v4_).length(0))
            (d_2948_v4_)[index465_] = default__.fm0((d_2951_v7_) + (d_2951_v7_), ((d_2948_v4_)[default__.safeIndex(78, (d_2948_v4_).length(0))]) == (default__.fm0(d_2951_v7_, True, p1, d_2964_v18_, globalState)), (self).f27, d_2964_v18_, globalState)
        index466_ = default__.safeIndex(877, (d_2944_v0_).length(0))
        (d_2944_v0_)[index466_] = ((self).f33) + (-741)
        d_2965_i2_: int
        d_2965_i2_ = 0
        with _dafny.label("17"):
            while not((self).f32):
                with _dafny.c_label("17"):
                    if (d_2965_i2_) >= (100):
                        raise _dafny.Break("17")
                    d_2965_i2_ = (d_2965_i2_) + (1)
                    if ((self).f33) == ((self).f33):
                        d_2966_v19_: D0
                        d_2966_v19_ = D0_DC0((d_2944_v0_)[default__.safeIndex(877, (d_2944_v0_).length(0))])
                        d_2967_v20_: _dafny.Map
                        d_2967_v20_ = _dafny.Map({d_2966_v19_: False})
                        d_2967_v20_ = (d_2967_v20_).set(d_2966_v19_, (self).f27)
                        d_2950_v6_ = _dafny.MultiSet([not((self).f27), False, (False) or ((d_2948_v4_)[default__.safeIndex(78, (d_2948_v4_).length(0))])])
                        (globalState).f6 = default__.safeDivisionInt(((self).f33) + ((self).f33), (self).f33)
                        d_2968_v21_: _dafny.Map
                        d_2968_v21_ = _dafny.Map({(d_2951_v7_) + (d_2951_v7_): _dafny.SeqWithoutIsStrInference([p1])})
                        d_2968_v21_ = (d_2968_v21_).set(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_2969_i3_ in range(default__.abs(-474))]), d_2949_v5_)
                        d_2970_v22_: _dafny.Set
                        d_2970_v22_ = _dafny.Set({(d_2944_v0_)[default__.safeIndex(877, (d_2944_v0_).length(0))]})
                        d_2951_v7_ = (d_2951_v7_ if default__.fm0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ptcbkywan")), (self).f32, (self).f27, d_2970_v22_, globalState) else d_2951_v7_)
                    elif True:
                        d_2971_v23_: _dafny.Array
                        nw471_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 22)
                        d_2971_v23_ = nw471_
                        index467_ = default__.safeIndex(833, (d_2971_v23_).length(0))
                        (d_2971_v23_)[index467_] = d_2951_v7_
                        d_2972_v24_: _dafny.Map
                        d_2972_v24_ = _dafny.Map({(self).f32: d_2958_v12_})
                        index468_ = default__.safeIndex(833, (d_2971_v23_).length(0))
                        (d_2971_v23_)[index468_] = (d_2951_v7_).set(default__.safeIndex((self).f33, len(d_2951_v7_)), ((d_2972_v24_)[True] if (True) in (d_2972_v24_) else d_2958_v12_))
                        d_2973_v25_: str
                        d_2973_v25_ = d_2958_v12_
                        d_2974_v26_: _dafny.Map
                        d_2974_v26_ = _dafny.Map({d_2973_v25_: ((d_2944_v0_)[default__.safeIndex(877, (d_2944_v0_).length(0))]) * ((self).f33)})
                        d_2975_v27_: _dafny.Map
                        d_2975_v27_ = _dafny.Map({(self).f27: (d_2953_v9_).cardinality})
                        d_2976_v28_: _dafny.MultiSet
                        d_2976_v28_ = _dafny.MultiSet([_dafny.MultiSet([(self).f27, (self).f32, (self).f32, (d_2948_v4_)[default__.safeIndex(78, (d_2948_v4_).length(0))], (self).f27]), d_2950_v6_, _dafny.MultiSet(d_2949_v5_), d_2950_v6_])
                        index469_ = default__.safeIndex(833, (d_2971_v23_).length(0))
                        index470_ = default__.safeIndex(877, (d_2944_v0_).length(0))
                        index471_ = default__.safeIndex(833, (d_2971_v23_).length(0))
                        rhs480_ = (_dafny.SeqWithoutIsStrInference([d_2958_v12_ for d_2977_i4_ in range(default__.abs(787))])).set(default__.safeIndex(len(d_2951_v7_), len(_dafny.SeqWithoutIsStrInference([d_2958_v12_ for d_2978_i4_ in range(default__.abs(787))]))), (d_2951_v7_)[default__.safeIndex(len(d_2975_v27_), len(d_2951_v7_))])
                        rhs481_ = _dafny.Map({d_2958_v12_: ((d_2976_v28_)[_dafny.MultiSet([not(True), (self).f27, (self).f32, (d_2948_v4_)[default__.safeIndex(78, (d_2948_v4_).length(0))], (self).f27])] if (_dafny.MultiSet([not(True), (self).f27, (self).f32, (d_2948_v4_)[default__.safeIndex(78, (d_2948_v4_).length(0))], (self).f27])) in (d_2976_v28_) else len(d_2951_v7_))})
                        rhs482_ = ((self).f33) * ((self).f33)
                        rhs483_ = (781) + ((self).f33)
                        rhs484_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hl"))
                        lhs408_ = d_2971_v23_
                        lhs409_ = default__.safeIndex(833, (d_2971_v23_).length(0))
                        lhs410_ = globalState
                        lhs411_ = d_2944_v0_
                        lhs412_ = default__.safeIndex(877, (d_2944_v0_).length(0))
                        lhs413_ = d_2971_v23_
                        lhs414_ = default__.safeIndex(833, (d_2971_v23_).length(0))
                        lhs408_[lhs409_] = rhs480_
                        d_2974_v26_ = rhs481_
                        lhs410_.f2 = rhs482_
                        lhs411_[lhs412_] = rhs483_
                        lhs413_[lhs414_] = rhs484_
                        (globalState).f6 = 146
                        d_2979_v29_: C1
                        nw472_ = C1()
                        nw472_.ctor__((self).f27)
                        d_2979_v29_ = nw472_
                        d_2980_v30_: D6
                        d_2980_v30_ = D6_DC14(((self).f32) and (True), default__.safeModuloInt((d_2944_v0_)[default__.safeIndex(877, (d_2944_v0_).length(0))], 173))
                        d_2980_v30_ = d_2980_v30_
                    d_2981_v31_: _dafny.Set
                    d_2981_v31_ = _dafny.Set({(self).f32, (self).f32, (d_2948_v4_)[default__.safeIndex(78, (d_2948_v4_).length(0))], (self).f27, p1})
                    d_2982_v32_: _dafny.Map
                    d_2982_v32_ = _dafny.Map({(d_2981_v31_) - (_dafny.Set({p1})): (d_2950_v6_).isdisjoint(_dafny.MultiSet([False]))})
                    d_2982_v32_ = (d_2982_v32_).set(d_2981_v31_, (d_2948_v4_)[default__.safeIndex(78, (d_2948_v4_).length(0))])
                    d_2983_v33_: _dafny.Array
                    nw473_ = _dafny.Array(None, 4)
                    nw473_[int(0)] = d_2981_v31_
                    nw473_[int(1)] = _dafny.Set({(d_2948_v4_)[default__.safeIndex(78, (d_2948_v4_).length(0))], p1})
                    nw473_[int(2)] = d_2981_v31_
                    nw473_[int(3)] = d_2981_v31_
                    d_2983_v33_ = nw473_
                    d_2984_v34_: _dafny.Map
                    d_2984_v34_ = _dafny.Map({d_2983_v33_: (d_2944_v0_)[default__.safeIndex(877, (d_2944_v0_).length(0))]})
                    (globalState).f6 = default__.safeModuloInt(((d_2984_v34_)[d_2983_v33_] if (d_2983_v33_) in (d_2984_v34_) else default__.fm1((self).f27, d_2949_v5_, p1, globalState)), len((d_2951_v7_) + (d_2951_v7_)))
                    d_2985_v35_: _dafny.Map
                    d_2985_v35_ = _dafny.Map({default__.fm10((self).f27, globalState): len(d_2949_v5_)})
                    d_2986_v36_: D3
                    d_2986_v36_ = D3_DC8(d_2985_v35_)
                    d_2987_v37_: _dafny.Set
                    d_2987_v37_ = _dafny.Set({d_2986_v36_, d_2986_v36_})
                    d_2988_v39_: _dafny.MultiSet
                    d_2988_v39_ = _dafny.MultiSet([d_2955_v11_])
                    d_2989_v40_: _dafny.Map
                    def iife267_():
                        coll109_ = _dafny.Map()
                        compr_109_: _dafny.Seq
                        for compr_109_ in (d_2988_v39_).Elements:
                            d_2990_v38_: _dafny.Seq = compr_109_
                            if (d_2990_v38_) in (d_2988_v39_):
                                coll109_[d_2990_v38_] = (d_2944_v0_)[default__.safeIndex(877, (d_2944_v0_).length(0))]
                        return _dafny.Map(coll109_)
                    d_2989_v40_ = _dafny.Map({D3_DC8(iife267_()
): False})
                    d_2991_v42_: _dafny.Map
                    d_2991_v42_ = _dafny.Map({d_2986_v36_: d_2953_v9_})
                    d_2992_v43_: D20
                    d_2992_v43_ = D20_DC48(d_2991_v42_)
                    d_2993_v44_: _dafny.Seq
                    d_2993_v44_ = _dafny.SeqWithoutIsStrInference([d_2991_v42_, (d_2992_v43_).cf68])
                    d_2994_v47_: _dafny.Array
                    nw474_ = _dafny.Array(None, 13)
                    nw474_[int(0)] = d_2987_v37_
                    nw474_[int(1)] = d_2987_v37_
                    nw474_[int(2)] = d_2987_v37_
                    def iife268_():
                        coll110_ = _dafny.Set()
                        compr_110_: D3
                        for compr_110_ in (d_2989_v40_).keys.Elements:
                            d_2995_v41_: D3 = compr_110_
                            if (d_2995_v41_) in (d_2989_v40_):
                                coll110_ = coll110_.union(_dafny.Set([d_2995_v41_]))
                        return _dafny.Set(coll110_)
                    nw474_[int(3)] = iife268_()
                    
                    nw474_[int(4)] = (d_2987_v37_) | (d_2987_v37_)
                    nw474_[int(5)] = d_2987_v37_
                    def iife269_():
                        def iife271_():
                            coll113_ = _dafny.Set()
                            compr_113_: D3
                            for compr_113_ in ((d_2993_v44_)[default__.safeIndex(595, len(d_2993_v44_))]).keys.Elements:
                                d_2998_v45_: D3 = compr_113_
                                if (d_2998_v45_) in ((d_2993_v44_)[default__.safeIndex(595, len(d_2993_v44_))]):
                                    coll113_ = coll113_.union(_dafny.Set([d_2998_v45_]))
                            return _dafny.Set(coll113_)
                        coll111_ = _dafny.Set()
                        def iife270_():
                            coll112_ = _dafny.Set()
                            compr_112_: D3
                            for compr_112_ in ((d_2993_v44_)[default__.safeIndex(595, len(d_2993_v44_))]).keys.Elements:
                                d_2996_v45_: D3 = compr_112_
                                if (d_2996_v45_) in ((d_2993_v44_)[default__.safeIndex(595, len(d_2993_v44_))]):
                                    coll112_ = coll112_.union(_dafny.Set([d_2996_v45_]))
                            return _dafny.Set(coll112_)
                        compr_111_: D3
                        for compr_111_ in (iife270_()
                        ).Elements:
                            d_2997_v46_: D3 = compr_111_
                            if (d_2997_v46_) in (iife271_()
                            ):
                                coll111_ = coll111_.union(_dafny.Set([d_2997_v46_]))
                        return _dafny.Set(coll111_)
                    nw474_[int(6)] = iife269_()
                    
                    nw474_[int(7)] = d_2987_v37_
                    nw474_[int(8)] = d_2987_v37_
                    nw474_[int(9)] = (d_2987_v37_) | (_dafny.Set({d_2986_v36_, d_2986_v36_}))
                    nw474_[int(10)] = (d_2987_v37_) | (d_2987_v37_)
                    nw474_[int(11)] = d_2987_v37_
                    nw474_[int(12)] = (default__.fm23(False, 287, len(d_2946_v2_), (self).f27, globalState)).intersection(d_2987_v37_)
                    d_2994_v47_ = nw474_
                    d_2994_v47_ = d_2994_v47_
                    pass
            pass
        hi18_ = (d_2950_v6_).cardinality
        for d_2999_i5_ in range((d_2944_v0_)[default__.safeIndex(877, (d_2944_v0_).length(0))], hi18_):
            (globalState).f21 = not((d_2951_v7_) == (d_2951_v7_))
            d_3000_v48_: C0
            nw475_ = C0()
            nw475_.ctor__()
            d_3000_v48_ = nw475_
            index472_ = default__.safeIndex(877, (d_2944_v0_).length(0))
            rhs485_ = (d_2999_i5_) * ((self).f33)
            rhs486_ = (default__.fm17((d_2944_v0_)[default__.safeIndex(877, (d_2944_v0_).length(0))], d_2951_v7_, (self).f33, globalState)) < ((d_2949_v5_) + (_dafny.SeqWithoutIsStrInference([(self).f32, (d_2948_v4_)[default__.safeIndex(78, (d_2948_v4_).length(0))], (self).f32, (self).f32, p1])))
            lhs415_ = d_2944_v0_
            lhs416_ = default__.safeIndex(877, (d_2944_v0_).length(0))
            lhs417_ = globalState
            lhs415_[lhs416_] = rhs485_
            lhs417_.f12 = rhs486_
            d_3001_v49_: _dafny.Set
            d_3001_v49_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([p1, p1])), (self).f33, (self).f33, (0) - ((self).f33)})
            d_3002_v50_: D0
            d_3002_v50_ = D0_DC1(default__.fm0(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qeaaow")), (self).f32, p1, d_3001_v49_, globalState))
            source44_ = d_3002_v50_
            if source44_.is_DC1:
                d_3003___mcc_h0_ = source44_.cf1
                d_3004_cf1_ = d_3003___mcc_h0_
                index473_ = default__.safeIndex(78, (d_2948_v4_).length(0))
                (d_2948_v4_)[index473_] = d_3004_cf1_
                d_3005_v51_: _dafny.Map
                d_3005_v51_ = _dafny.Map({len(d_2951_v7_): d_2951_v7_})
                d_3006_v52_: _dafny.Set
                d_3006_v52_ = _dafny.Set({(self).fm4(D0_DC1((d_2948_v4_)[default__.safeIndex(78, (d_2948_v4_).length(0))]), globalState)})
                (globalState).f2 = ((376) + (len(((d_3005_v51_)[(d_2944_v0_)[default__.safeIndex(877, (d_2944_v0_).length(0))]] if ((d_2944_v0_)[default__.safeIndex(877, (d_2944_v0_).length(0))]) in (d_3005_v51_) else (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eeghbvio"))).set(default__.safeIndex(len(d_3006_v52_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eeghbvio")))), default__.fm14(117, globalState)))))) - (d_2999_i5_)
                d_3007_v53_: _dafny.Array
                nw476_ = _dafny.Array(_dafny.Array(None, 0), 17)
                d_3007_v53_ = nw476_
                index474_ = default__.safeIndex(503, (d_3007_v53_).length(0))
                (d_3007_v53_)[index474_] = d_2944_v0_
                index475_ = default__.safeIndex(503, (d_3007_v53_).length(0))
                (d_3007_v53_)[index475_] = d_2944_v0_
                d_3008_v54_: D16
                d_3008_v54_ = D16_DC35(_dafny.MultiSet([(self).f33, (self).f33]))
                d_2953_v9_ = (d_3008_v54_).cf50
            elif source44_.is_DC2:
                d_3009___mcc_h1_ = source44_.cf2
                d_3010___mcc_h2_ = source44_.cf3
                d_3011_cf3_ = d_3010___mcc_h2_
                d_3012_cf2_ = d_3009___mcc_h1_
                index476_ = default__.safeIndex(78, (d_2948_v4_).length(0))
                (d_2948_v4_)[index476_] = (default__.fm0(d_2951_v7_, (self).fm8(globalState), (d_2948_v4_)[default__.safeIndex(78, (d_2948_v4_).length(0))], d_3001_v49_, globalState) if (self).f27 else (False if (d_2948_v4_)[default__.safeIndex(78, (d_2948_v4_).length(0))] else (self).fm4(d_3002_v50_, globalState)))
                (globalState).f2 = (0) - ((default__.fm1((self).f27, d_2949_v5_, (d_2948_v4_)[default__.safeIndex(78, (d_2948_v4_).length(0))], globalState)) - (d_3011_cf3_))
                d_3013_v55_: _dafny.Array
                nw477_ = _dafny.Array(_dafny.Seq({}), 13)
                d_3013_v55_ = nw477_
                d_3014_v56_: _dafny.Map
                d_3014_v56_ = _dafny.Map({d_3013_v55_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))})
                d_3014_v56_ = (d_3014_v56_).set(d_3013_v55_, (d_2951_v7_) + (d_2951_v7_))
                r0 = (d_3011_cf3_) <= ((d_2944_v0_)[default__.safeIndex(877, (d_2944_v0_).length(0))])
            elif source44_.is_DC0:
                d_3015___mcc_h3_ = source44_.cf0
                d_3016_cf0_ = d_3015___mcc_h3_
                (globalState).f2 = default__.safeDivisionInt(default__.safeDivisionInt((self).f33, (d_2944_v0_)[default__.safeIndex(877, (d_2944_v0_).length(0))]), (self).f33)
                d_3017_v57_: _dafny.Map
                d_3017_v57_ = _dafny.Map({(d_2955_v11_)[default__.safeIndex((d_2944_v0_)[default__.safeIndex(877, (d_2944_v0_).length(0))], len(d_2955_v11_))]: (self).f32})
                (globalState).f12 = ((d_3017_v57_)[d_3016_cf0_] if (d_3016_cf0_) in (d_3017_v57_) else (d_2948_v4_)[default__.safeIndex(78, (d_2948_v4_).length(0))])
                index477_ = default__.safeIndex(877, (d_2944_v0_).length(0))
                (d_2944_v0_)[index477_] = (self).f33
                d_3018_v58_: _dafny.Map
                d_3018_v58_ = _dafny.Map({546: default__.fm14(d_3016_cf0_, globalState)})
                d_3019_v59_: _dafny.Map
                d_3019_v59_ = _dafny.Map({d_2958_v12_: d_3018_v58_})
                d_3020_v60_: _dafny.Seq
                d_3020_v60_ = _dafny.SeqWithoutIsStrInference([d_2951_v7_, d_2951_v7_, (self).fm2((self).f33, d_3019_v59_, globalState)])
                d_3020_v60_ = (d_3020_v60_) + ((d_3020_v60_) + (d_3020_v60_))
            elif True:
                d_3021___mcc_h4_ = source44_.cf4
                d_3022_cf4_ = d_3021___mcc_h4_
                d_3023_v61_: _dafny.Map
                d_3023_v61_ = _dafny.Map({d_2958_v12_: (0) - (d_2999_i5_)})
                d_3024_v62_: _dafny.MultiSet
                d_3024_v62_ = _dafny.MultiSet([d_3023_v61_, d_3023_v61_])
                d_3025_v63_: _dafny.MultiSet
                d_3025_v63_ = (_dafny.MultiSet([d_3023_v61_])) | (d_3024_v62_)
                d_3025_v63_ = d_3025_v63_
                d_3026_v64_: D0
                d_3026_v64_ = D0_DC2((0) - ((d_2944_v0_)[default__.safeIndex(877, (d_2944_v0_).length(0))]), (self).f33)
                pat_let_tv60_ = d_3026_v64_
                d_3027_v65_: bool
                d_3028_v66_: _dafny.Seq
                out81_: bool
                out82_: _dafny.Seq
                def iife272_(_pat_let79_0):
                    def iife273_(d_3029_dt__update__tmp_h0_):
                        def iife274_(_pat_let80_0):
                            def iife275_(d_3030_dt__update_hcf4_h0_):
                                return D0_DC3(d_3030_dt__update_hcf4_h0_)
                            return iife275_(_pat_let80_0)
                        return iife274_(pat_let_tv60_)
                    return iife273_(_pat_let79_0)
                out81_, out82_ = default__.m0(iife272_(p0), d_3002_v50_, globalState)
                d_3027_v65_ = out81_
                d_3028_v66_ = out82_
                d_3031_v67_: C2
                nw478_ = C2()
                nw478_.ctor__(d_3027_v65_, p1)
                d_3031_v67_ = nw478_
                d_3032_v68_: _dafny.Seq
                d_3032_v68_ = _dafny.SeqWithoutIsStrInference([d_3028_v66_])
                d_3033_v69_: _dafny.Map
                d_3033_v69_ = _dafny.Map({d_3032_v68_: _dafny.SeqWithoutIsStrInference([d_2958_v12_ for d_3034_i6_ in range(default__.abs(893))])})
                d_3033_v69_ = (d_3033_v69_).set(d_3032_v68_, d_2951_v7_)
        d_3035_v70_: _dafny.Map
        d_3035_v70_ = _dafny.Map({(d_2949_v5_) == (default__.fm17((d_2944_v0_)[default__.safeIndex(877, (d_2944_v0_).length(0))], d_2951_v7_, (self).f33, globalState)): 445})
        index478_ = default__.safeIndex(877, (d_2944_v0_).length(0))
        (d_2944_v0_)[index478_] = len(d_3035_v70_)
        r0 = p1
        d_3036_v71_: D0
        d_3036_v71_ = D0_DC1((self).f27)
        pat_let_tv61_ = p1
        pat_let_tv62_ = p1
        pat_let_tv63_ = d_2949_v5_
        pat_let_tv64_ = d_2949_v5_
        def lambda129_(source45_):
            if source45_.is_DC1:
                d_3037___mcc_h5_ = source45_.cf1
                d_3038_cf1_ = d_3037___mcc_h5_
                return pat_let_tv61_
            elif source45_.is_DC2:
                d_3039___mcc_h6_ = source45_.cf2
                d_3040___mcc_h7_ = source45_.cf3
                d_3041_cf3_ = d_3040___mcc_h7_
                d_3042_cf2_ = d_3039___mcc_h6_
                return pat_let_tv62_
            elif source45_.is_DC0:
                d_3043___mcc_h8_ = source45_.cf0
                d_3044_cf0_ = d_3043___mcc_h8_
                return (pat_let_tv63_)[default__.safeIndex(d_3044_cf0_, len(pat_let_tv64_))]
            elif True:
                d_3045___mcc_h9_ = source45_.cf4
                d_3046_cf4_ = d_3045___mcc_h9_
                return not(True)

        r1 = lambda129_(d_3036_v71_)
        return r0, r1

    def m1(self, globalState):
        r0: int = int(0)
        d_3047_v0_: _dafny.Seq
        d_3047_v0_ = _dafny.SeqWithoutIsStrInference([((self).f33) >= ((self).f33)])
        if (d_3047_v0_)[default__.safeIndex((self).f33, len(d_3047_v0_))]:
            d_3048_v1_: C1
            nw479_ = C1()
            nw479_.ctor__(not((self).f27))
            d_3048_v1_ = nw479_
            d_3049_v2_: D3
            d_3049_v2_ = D3_DC8(_dafny.Map({default__.fm12((self).f33, (self).f32, (self).f32, globalState): len(d_3047_v0_)}))
            d_3050_v3_: _dafny.Map
            d_3050_v3_ = _dafny.Map({d_3049_v2_: _dafny.MultiSet([default__.fm1((self).f27, d_3047_v0_, (self).f32, globalState)])})
            d_3051_v4_: D20
            d_3051_v4_ = D20_DC48(d_3050_v3_)
            d_3051_v4_ = D20_DC48(d_3050_v3_)
            d_3052_v5_: _dafny.Set
            d_3052_v5_ = _dafny.Set({621, (self).f33, (self).f33, (self).f33})
            d_3053_v6_: D0
            d_3053_v6_ = D0_DC0(len(d_3052_v5_))
            (globalState).f6 = (d_3053_v6_).cf0
            d_3054_v7_: _dafny.MultiSet
            d_3054_v7_ = _dafny.MultiSet([(self).f27, (self).f32, not(not(not((self).f32))), (self).f27])
            d_3055_v8_: _dafny.Seq
            d_3055_v8_ = _dafny.SeqWithoutIsStrInference([(self).f33])
            d_3056_v9_: T1
            nw480_ = C7()
            nw480_.ctor__((self).f27, (d_3054_v7_).cardinality, len(d_3055_v8_), (0) - (((self).f33 if (self).fm8(globalState) else (self).f33)), not((self).f32))
            d_3056_v9_ = nw480_
            d_3056_v9_ = d_3056_v9_
            d_3057_v11_: _dafny.Array
            def lambda130_(d_3058_i0_):
                def iife276_():
                    coll114_ = _dafny.Set()
                    compr_114_: str
                    for compr_114_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d')])).Elements:
                        d_3059_v10_: str = compr_114_
                        if (d_3059_v10_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d')])):
                            coll114_ = coll114_.union(_dafny.Set([d_3059_v10_]))
                    return _dafny.Set(coll114_)
                return (d_3058_i0_) * (len(iife276_()
                ))

            init76_ = lambda130_
            nw481_ = _dafny.Array(None, 21)
            for i0_76_ in range(nw481_.length(0)):
                nw481_[i0_76_] = init76_(i0_76_)
            d_3057_v11_ = nw481_
            index479_ = default__.safeIndex(336, (d_3057_v11_).length(0))
            (d_3057_v11_)[index479_] = default__.safeDivisionInt(-926, (self).f33)
            index480_ = default__.safeIndex(336, (d_3057_v11_).length(0))
            (d_3057_v11_)[index480_] = (0) - ((0) - ((0) - ((d_3056_v9_).f28)))
        elif True:
            d_3060_v12_: str
            d_3060_v12_ = _dafny.CodePoint('q')
            d_3061_v13_: D22
            d_3061_v13_ = D22_DC58((self).f32, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "devlyg"))), 839)
            d_3062_v14_: _dafny.Set
            d_3062_v14_ = _dafny.Set({d_3061_v13_, d_3061_v13_, d_3061_v13_, d_3061_v13_})
            d_3063_v15_: D21
            d_3063_v15_ = D21_DC55(552)
            rhs487_ = ((self).f33 if (self).f32 else len(_dafny.Map({(self).f32: d_3060_v12_})))
            rhs488_ = (_dafny.Set({default__.fm68((self).f32, d_3063_v15_, d_3060_v12_, (self).f27, globalState)})).ispropersubset(d_3062_v14_)
            lhs418_ = globalState
            lhs419_ = globalState
            lhs418_.f2 = rhs487_
            lhs419_.f21 = rhs488_
            if (self).f32:
                d_3064_v16_: _dafny.Array
                nw482_ = _dafny.Array(None, 3)
                nw482_[int(0)] = not ((self).f27) or (not((self).f27))
                nw482_[int(1)] = (self).f32
                nw482_[int(2)] = ((0) - (-223)) >= ((self).f33)
                d_3064_v16_ = nw482_
                d_3064_v16_ = (d_3064_v16_ if (21) >= ((self).f33) else d_3064_v16_)
                (globalState).f2 = (self).f33
                d_3065_v17_: _dafny.Array
                nw483_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 18)
                d_3065_v17_ = nw483_
                d_3066_v18_: _dafny.Seq
                d_3066_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kveawytlh"))
                index481_ = default__.safeIndex(543, (d_3065_v17_).length(0))
                (d_3065_v17_)[index481_] = ((d_3066_v18_).set(default__.safeIndex(971, len(d_3066_v18_)), d_3060_v12_)).set(default__.safeIndex((self).f33, len((d_3066_v18_).set(default__.safeIndex(971, len(d_3066_v18_)), d_3060_v12_))), d_3060_v12_)
                index482_ = default__.safeIndex(543, (d_3065_v17_).length(0))
                (d_3065_v17_)[index482_] = (self).fm3(globalState)
                d_3067_v19_: _dafny.Array
                nw484_ = _dafny.Array(int(0), 17)
                d_3067_v19_ = nw484_
                index483_ = default__.safeIndex(438, (d_3067_v19_).length(0))
                (d_3067_v19_)[index483_] = (self).f33
                index484_ = default__.safeIndex(438, (d_3067_v19_).length(0))
                (d_3067_v19_)[index484_] = (((self).f33) + ((self).f33)) - ((self).f33)
                d_3068_v20_: _dafny.MultiSet
                d_3068_v20_ = _dafny.MultiSet([(self).f33, (self).f33, (d_3067_v19_)[default__.safeIndex(438, (d_3067_v19_).length(0))]])
                (globalState).f25 = not((d_3068_v20_).isdisjoint(((d_3068_v20_).set((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f33, (d_3067_v19_)[default__.safeIndex(438, (d_3067_v19_).length(0))], len(d_3047_v0_), (self).f33]))).cardinality, default__.abs((self).f33))) - (d_3068_v20_)))
            elif True:
                d_3069_v21_: _dafny.Seq
                d_3069_v21_ = _dafny.SeqWithoutIsStrInference([(self).f33, (self).f33])
                (globalState).f6 = (0) - ((len((_dafny.Map({(d_3069_v21_)[default__.safeIndex((self).f33, len(d_3069_v21_))]: (self).f33})) | (_dafny.Map({821: (self).f33})))) * ((0) - ((d_3069_v21_)[default__.safeIndex((self).f33, len(d_3069_v21_))])))
                d_3070_v22_: C1
                nw485_ = C1()
                nw485_.ctor__((self).f27)
                d_3070_v22_ = nw485_
                d_3071_v23_: _dafny.Array
                def lambda131_(d_3072_v0_):
                    def lambda132_(d_3073_i1_):
                        return ((d_3072_v0_)[default__.safeIndex((self).f33, len(d_3072_v0_))]) or ((d_3072_v0_)[default__.safeIndex((self).f33, len(d_3072_v0_))])

                    return lambda132_

                init77_ = lambda131_(d_3047_v0_)
                nw486_ = _dafny.Array(None, 28)
                for i0_77_ in range(nw486_.length(0)):
                    nw486_[i0_77_] = init77_(i0_77_)
                d_3071_v23_ = nw486_
                index485_ = default__.safeIndex(596, (d_3071_v23_).length(0))
                (d_3071_v23_)[index485_] = (self).f27
                d_3074_v24_: _dafny.Seq
                d_3074_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jcybyo"))
                index486_ = default__.safeIndex(596, (d_3071_v23_).length(0))
                (d_3071_v23_)[index486_] = (d_3074_v24_) <= ((d_3074_v24_).set(default__.safeIndex((self).f33, len(d_3074_v24_)), d_3060_v12_))
                d_3075_v25_: C7
                nw487_ = C7()
                nw487_.ctor__((d_3071_v23_)[default__.safeIndex(596, (d_3071_v23_).length(0))], (d_3069_v21_)[default__.safeIndex((self).f33, len(d_3069_v21_))], (self).f33, 298, ((self).f27) == ((self).f32))
                d_3075_v25_ = nw487_
                (globalState).f5 = (self).f32
            d_3076_v26_: _dafny.Array
            nw488_ = _dafny.Array(int(0), 18)
            d_3076_v26_ = nw488_
            index487_ = default__.safeIndex(250, (d_3076_v26_).length(0))
            def iife277_():
                coll115_ = _dafny.Map()
                compr_115_: int
                for compr_115_ in _dafny.IntegerRange(54, 507):
                    d_3077_v27_: int = compr_115_
                    if ((54) <= (d_3077_v27_)) and ((d_3077_v27_) < (507)):
                        coll115_[(d_3077_v27_) * ((self).f33)] = (self).f33
                return _dafny.Map(coll115_)
            (d_3076_v26_)[index487_] = ((self).f33) - ((0) - (len(iife277_()
            )))
            d_3078_v28_: _dafny.Set
            d_3078_v28_ = _dafny.Set({((self).f32) and ((self).f27)})
            index488_ = default__.safeIndex(250, (d_3076_v26_).length(0))
            (d_3076_v26_)[index488_] = len(d_3078_v28_)
            (globalState).f2 = (self).f33
            index489_ = default__.safeIndex(250, (d_3076_v26_).length(0))
            (d_3076_v26_)[index489_] = (0) - (((self).f33) + ((self).f33))
        d_3079_v29_: _dafny.MultiSet
        d_3079_v29_ = _dafny.MultiSet([(self).f27, (self).f27, (self).f32, (self).f27])
        d_3080_v30_: _dafny.Seq
        d_3080_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ncpnmtl"))
        d_3081_v31_: _dafny.Map
        d_3081_v31_ = _dafny.Map({d_3080_v30_: (self).f33})
        d_3082_v32_: _dafny.Seq
        d_3082_v32_ = _dafny.SeqWithoutIsStrInference([((d_3081_v31_)[d_3080_v30_] if (d_3080_v30_) in (d_3081_v31_) else (self).f33)])
        d_3083_v33_: _dafny.Set
        d_3083_v33_ = _dafny.Set({(self).f33})
        d_3084_v34_: _dafny.Array
        nw489_ = _dafny.Array(None, 8)
        nw489_[int(0)] = (self).fm8(globalState)
        nw489_[int(1)] = (d_3047_v0_)[default__.safeIndex((d_3079_v29_).cardinality, len(d_3047_v0_))]
        nw489_[int(2)] = (self).f32
        nw489_[int(3)] = (self).f32
        nw489_[int(4)] = ((self).f32) and ((self).f27)
        nw489_[int(5)] = True
        nw489_[int(6)] = (_dafny.SeqWithoutIsStrInference([(self).f33, (self).f33])) != (d_3082_v32_)
        nw489_[int(7)] = not((d_3083_v33_).isdisjoint(d_3083_v33_))
        d_3084_v34_ = nw489_
        index490_ = default__.safeIndex(758, (d_3084_v34_).length(0))
        (d_3084_v34_)[index490_] = (self).f27
        index491_ = default__.safeIndex(758, (d_3084_v34_).length(0))
        (d_3084_v34_)[index491_] = False
        d_3085_v35_: _dafny.Map
        d_3085_v35_ = _dafny.Map({default__.fm1(False, d_3047_v0_, (d_3084_v34_)[default__.safeIndex(758, (d_3084_v34_).length(0))], globalState): (self).f33})
        hi19_ = (len(d_3083_v33_)) * (len(d_3080_v30_))
        for d_3086_i2_ in range(default__.safeDivisionInt(((d_3079_v29_)[(self).f32] if ((self).f32) in (d_3079_v29_) else len(d_3085_v35_)), (self).f33), hi19_):
            d_3087_v36_: _dafny.Array
            nw490_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 23)
            d_3087_v36_ = nw490_
            d_3088_v37_: _dafny.Array
            nw491_ = _dafny.Array(None, 1)
            nw491_[int(0)] = d_3086_i2_
            d_3088_v37_ = nw491_
            d_3089_v38_: _dafny.Map
            d_3089_v38_ = _dafny.Map({(self).f33: d_3088_v37_})
            rhs489_ = (((self).f33) + (len(d_3089_v38_))) + ((self).f33)
            rhs490_ = d_3080_v30_
            rhs491_ = ((self).f33) - ((0) - ((self).f33))
            rhs492_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rsgqkf"))
            rhs493_ = d_3087_v36_
            lhs420_ = globalState
            lhs421_ = globalState
            lhs420_.f2 = rhs489_
            d_3080_v30_ = rhs490_
            lhs421_.f6 = rhs491_
            d_3080_v30_ = rhs492_
            d_3087_v36_ = rhs493_
            if not ((self).f27) or ((self).f27):
                d_3090_v39_: T1
                nw492_ = C7()
                nw492_.ctor__((d_3084_v34_)[default__.safeIndex(758, (d_3084_v34_).length(0))], (self).f33, (self).f33, 638, (self).f32)
                d_3090_v39_ = nw492_
                d_3090_v39_ = d_3090_v39_
                d_3091_v40_: _dafny.Array
                nw493_ = _dafny.Array(_dafny.Seq({}), 25)
                d_3091_v40_ = nw493_
                index492_ = default__.safeIndex(780, (d_3091_v40_).length(0))
                (d_3091_v40_)[index492_] = _dafny.SeqWithoutIsStrInference([d_3047_v0_, d_3047_v0_, d_3047_v0_])
                d_3092_v41_: _dafny.Seq
                d_3092_v41_ = _dafny.SeqWithoutIsStrInference([d_3047_v0_])
                index493_ = default__.safeIndex(780, (d_3091_v40_).length(0))
                (d_3091_v40_)[index493_] = (_dafny.SeqWithoutIsStrInference([d_3047_v0_, d_3047_v0_])) + (d_3092_v41_)
                r0 = ((d_3090_v39_).f28) * (d_3086_i2_)
                (globalState).f2 = default__.safeDivisionInt((d_3086_i2_) * ((d_3090_v39_).f29), (self).f33)
                d_3093_v42_: C2
                nw494_ = C2()
                nw494_.ctor__(not((d_3047_v0_)[default__.safeIndex((d_3090_v39_).f29, len(d_3047_v0_))]), (d_3090_v39_).f27)
                d_3093_v42_ = nw494_
            elif True:
                r0 = d_3086_i2_
                d_3094_v43_: str
                d_3094_v43_ = _dafny.CodePoint('e')
                d_3094_v43_ = _dafny.CodePoint('j')
                d_3095_v44_: _dafny.MultiSet
                d_3095_v44_ = _dafny.MultiSet([d_3086_i2_])
                d_3096_v45_: C2
                nw495_ = C2()
                nw495_.ctor__((d_3095_v44_) == (d_3095_v44_), not((d_3084_v34_)[default__.safeIndex(758, (d_3084_v34_).length(0))]))
                d_3096_v45_ = nw495_
                d_3096_v45_ = d_3096_v45_
                d_3097_v46_: _dafny.Map
                d_3097_v46_ = _dafny.Map({d_3086_i2_: d_3084_v34_})
                d_3098_v47_: _dafny.Map
                d_3098_v47_ = _dafny.Map({d_3082_v32_: len(d_3097_v46_)})
                d_3098_v47_ = (d_3098_v47_).set(d_3082_v32_, (self).f33)
                d_3099_v48_: int
                d_3100_v49_: int
                out83_: int
                out84_: int
                out83_, out84_ = (self).m6(globalState)
                d_3099_v48_ = out83_
                d_3100_v49_ = out84_
            (globalState).f21 = (self).f32
            d_3101_v50_: _dafny.Map
            d_3101_v50_ = _dafny.Map({(self).f33: ((self).f33) > (-331)})
            d_3102_v51_: T2
            nw496_ = C2()
            nw496_.ctor__((self).f32, (self).f32)
            d_3102_v51_ = nw496_
            d_3103_v52_: _dafny.Map
            d_3103_v52_ = _dafny.Map({len(d_3082_v32_): d_3102_v51_})
            d_3104_v53_: _dafny.MultiSet
            d_3104_v53_ = _dafny.MultiSet([d_3103_v52_])
            d_3101_v50_ = (d_3101_v50_).set(default__.safeModuloInt((d_3104_v53_).cardinality, len(d_3083_v33_)), (d_3084_v34_)[default__.safeIndex(758, (d_3084_v34_).length(0))])
        d_3105_v54_: _dafny.Array
        nw497_ = _dafny.Array(_dafny.Seq({}), 18)
        d_3105_v54_ = nw497_
        d_3106_v55_: _dafny.Map
        d_3106_v55_ = _dafny.Map({d_3105_v54_: (self).f32})
        d_3107_v56_: _dafny.MultiSet
        d_3107_v56_ = _dafny.MultiSet([(self).f33])
        d_3106_v55_ = (d_3106_v55_).set(d_3105_v54_, (d_3107_v56_).ispropersubset(d_3107_v56_))
        if (self).f27:
            d_3108_v58_: _dafny.Map
            def iife278_():
                coll116_ = _dafny.Set()
                compr_116_: int
                for compr_116_ in (d_3082_v32_).Elements:
                    d_3109_v57_: int = compr_116_
                    if (d_3109_v57_) in (d_3082_v32_):
                        coll116_ = coll116_.union(_dafny.Set([(d_3109_v57_) + (len(_dafny.Map({not(True): -415})))]))
                return _dafny.Set(coll116_)
            d_3108_v58_ = _dafny.Map({d_3080_v30_: iife278_()
            })
            pat_let_tv65_ = d_3108_v58_
            def iife279_(_pat_let81_0):
                def iife280_(d_3110_dt__update__tmp_h0_):
                    def iife281_(_pat_let82_0):
                        def iife282_(d_3111_dt__update_hcf97_h0_):
                            return D25_DC63(d_3111_dt__update_hcf97_h0_)
                        return iife282_(_pat_let82_0)
                    return iife281_(pat_let_tv65_)
                return iife280_(_pat_let81_0)
            source46_ = iife279_(D25_DC63(d_3108_v58_))
            if source46_.is_DC64:
                d_3112___mcc_h0_ = source46_.cf98
                d_3113___mcc_h1_ = source46_.cf99
                d_3114___mcc_h2_ = source46_.cf100
                d_3115___mcc_h3_ = source46_.cf101
                d_3116___mcc_h4_ = source46_.cf102
                d_3117_cf102_ = d_3116___mcc_h4_
                d_3118_cf101_ = d_3115___mcc_h3_
                d_3119_cf100_ = d_3114___mcc_h2_
                d_3120_cf99_ = d_3113___mcc_h1_
                d_3121_cf98_ = d_3112___mcc_h0_
                d_3122_v59_: _dafny.Map
                d_3122_v59_ = _dafny.Map({d_3121_cf98_: len(d_3085_v35_)})
                d_3123_v60_: _dafny.Seq
                d_3123_v60_ = _dafny.SeqWithoutIsStrInference([d_3047_v0_])
                d_3122_v59_ = (d_3122_v59_).set((d_3084_v34_)[default__.safeIndex(758, (d_3084_v34_).length(0))], default__.fm1((d_3084_v34_)[default__.safeIndex(758, (d_3084_v34_).length(0))], (d_3123_v60_)[default__.safeIndex((self).f33, len(d_3123_v60_))], (self).f27, globalState))
                d_3117_cf102_ = d_3120_cf99_
                (globalState).f21 = (self).f27
                (globalState).f2 = d_3117_cf102_
            elif True:
                d_3124___mcc_h5_ = source46_.cf97
                d_3125_cf97_ = d_3124___mcc_h5_
                d_3085_v35_ = (d_3085_v35_).set((self).f33, (self).f33)
                d_3126_v61_: str
                d_3126_v61_ = _dafny.CodePoint('s')
                rhs494_ = (d_3047_v0_)[default__.safeIndex((self).f33, len(d_3047_v0_))]
                rhs495_ = d_3126_v61_
                lhs422_ = globalState
                lhs422_.f5 = rhs494_
                d_3126_v61_ = rhs495_
                d_3127_v62_: _dafny.Array
                def lambda133_(d_3128_cf97_):
                    def lambda134_(d_3129_i3_):
                        return D25_DC63(d_3128_cf97_)

                    return lambda134_

                init78_ = lambda133_(d_3125_cf97_)
                nw498_ = _dafny.Array(None, 8)
                for i0_78_ in range(nw498_.length(0)):
                    nw498_[i0_78_] = init78_(i0_78_)
                d_3127_v62_ = nw498_
                d_3127_v62_ = (d_3127_v62_)
                (globalState).f14 = False
            index494_ = default__.safeIndex(758, (d_3084_v34_).length(0))
            (d_3084_v34_)[index494_] = ((self).f33) <= (((d_3079_v29_).cardinality) * (-182))
            d_3130_v63_: str
            d_3130_v63_ = _dafny.CodePoint('t')
            d_3131_v64_: _dafny.Map
            d_3131_v64_ = _dafny.Map({d_3130_v63_: (0) - (len(d_3080_v30_))})
            d_3131_v64_ = (d_3131_v64_).set(d_3130_v63_, (self).f33)
            d_3132_v65_: D14
            d_3132_v65_ = D14_DC31((self).f33)
            index495_ = default__.safeIndex(758, (d_3084_v34_).length(0))
            (d_3084_v34_)[index495_] = ((d_3132_v65_).cf43) != ((self).f33)
            d_3133_v66_: _dafny.Map
            d_3133_v66_ = _dafny.Map({(self).f32: d_3084_v34_})
            d_3134_v67_: _dafny.Map
            d_3134_v67_ = _dafny.Map({((d_3133_v66_)[True] if (True) in (d_3133_v66_) else d_3084_v34_): False})
            d_3135_v68_: C7
            nw499_ = C7()
            nw499_.ctor__((d_3084_v34_)[default__.safeIndex(758, (d_3084_v34_).length(0))], (self).f33, (self).f33, 407, (self).f27)
            d_3135_v68_ = nw499_
            d_3136_v69_: _dafny.Seq
            d_3136_v69_ = _dafny.SeqWithoutIsStrInference([d_3135_v68_, d_3135_v68_])
            d_3137_v70_: D2
            d_3137_v70_ = D2_DC6(len(d_3136_v69_), _dafny.SeqWithoutIsStrInference([(self).f33]), (self).f32, d_3082_v32_, (self).f27)
            d_3138_v71_: _dafny.Map
            d_3138_v71_ = _dafny.Map({(d_3134_v67_) == (d_3134_v67_): not((d_3082_v32_) <= ((d_3137_v70_).cf10))})
            d_3138_v71_ = ((d_3138_v71_) | (d_3138_v71_)) | ((d_3138_v71_ if (d_3084_v34_)[default__.safeIndex(758, (d_3084_v34_).length(0))] else d_3138_v71_))
        elif True:
            d_3139_v72_: _dafny.Map
            d_3139_v72_ = _dafny.Map({(self).f33: not((d_3084_v34_)[default__.safeIndex(758, (d_3084_v34_).length(0))])})
            (globalState).f14 = ((d_3139_v72_)[(self).f33] if ((self).f33) in (d_3139_v72_) else not((d_3084_v34_)[default__.safeIndex(758, (d_3084_v34_).length(0))]))
            d_3140_v73_: D0
            d_3140_v73_ = D0_DC0((0) - ((self).f33))
            d_3140_v73_ = d_3140_v73_
            d_3141_v74_: _dafny.Set
            d_3141_v74_ = _dafny.Set({(self).f27, True})
            d_3142_v75_: _dafny.Map
            d_3142_v75_ = _dafny.Map({((self).f27) not in (d_3141_v74_): _dafny.SeqWithoutIsStrInference([(self).f27])})
            rhs496_ = (self).f33
            rhs497_ = (self).f33
            rhs498_ = (0) - (len(d_3142_v75_))
            lhs423_ = globalState
            lhs424_ = globalState
            r0 = rhs496_
            lhs423_.f2 = rhs497_
            lhs424_.f6 = rhs498_
            (globalState).f6 = (self).f33
            d_3143_v76_: str
            d_3143_v76_ = _dafny.CodePoint('x')
            d_3144_v77_: _dafny.Map
            d_3144_v77_ = _dafny.Map({d_3143_v76_: (self).f27})
            d_3145_v78_: C3
            nw500_ = C3()
            nw500_.ctor__((len(d_3082_v32_)) >= (len(d_3144_v77_)), not((self).f32))
            d_3145_v78_ = nw500_
        if False:
            (globalState).f6 = 896
            d_3146_v79_: _dafny.Array
            nw501_ = _dafny.Array(D5.default()(), 6)
            d_3146_v79_ = nw501_
            d_3147_v80_: _dafny.Array
            nw502_ = _dafny.Array(_dafny.Array(None, 0), 20)
            d_3147_v80_ = nw502_
            d_3148_v81_: D5
            d_3148_v81_ = D5_DC11(d_3147_v80_)
            index496_ = default__.safeIndex(6, (d_3146_v79_).length(0))
            (d_3146_v79_)[index496_] = d_3148_v81_
            index497_ = default__.safeIndex(6, (d_3146_v79_).length(0))
            (d_3146_v79_)[index497_] = d_3148_v81_
            d_3149_v82_: str
            d_3149_v82_ = _dafny.CodePoint('m')
            d_3150_v83_: _dafny.Map
            d_3150_v83_ = _dafny.Map({(d_3084_v34_)[default__.safeIndex(758, (d_3084_v34_).length(0))]: d_3149_v82_})
            d_3151_v84_: D9
            d_3151_v84_ = D9_DC20(d_3150_v83_)
            d_3152_v85_: D9
            d_3152_v85_ = D9_DC22(d_3151_v84_)
            d_3153_v86_: D30
            d_3153_v86_ = D30_DC74((self).f32, 694, (self).f27)
            pat_let_tv66_ = d_3084_v34_
            pat_let_tv67_ = d_3084_v34_
            d_3154_v87_: C4
            nw503_ = C4()
            def iife283_(_pat_let83_0):
                def iife284_(d_3155_dt__update__tmp_h1_):
                    def iife285_(_pat_let84_0):
                        def iife286_(d_3156_dt__update_hcf121_h0_):
                            return D30_DC74(d_3156_dt__update_hcf121_h0_, (d_3155_dt__update__tmp_h1_).cf122, (d_3155_dt__update__tmp_h1_).cf123)
                        return iife286_(_pat_let84_0)
                    return iife285_((pat_let_tv67_)[default__.safeIndex(758, (pat_let_tv66_).length(0))])
                return iife284_(_pat_let83_0)
            nw503_.ctor__(d_3079_v29_, ((iife283_(d_3153_v86_)).cf121) == (False))
            d_3154_v87_ = nw503_
            d_3157_v88_: _dafny.Seq
            d_3157_v88_ = _dafny.SeqWithoutIsStrInference([d_3079_v29_, d_3079_v29_])
            d_3158_v89_: _dafny.Array
            nw504_ = _dafny.Array(None, 26)
            nw504_[int(0)] = default__.fm34(globalState)
            nw504_[int(1)] = d_3079_v29_
            nw504_[int(2)] = d_3079_v29_
            nw504_[int(3)] = ((d_3154_v87_).f37).set(False, default__.abs((self).f33))
            nw504_[int(4)] = _dafny.MultiSet((d_3047_v0_) + (d_3047_v0_))
            nw504_[int(5)] = _dafny.MultiSet(d_3047_v0_)
            nw504_[int(6)] = (d_3079_v29_).set((self).f27, default__.abs(len(d_3080_v30_)))
            nw504_[int(7)] = _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([(self).f32, (d_3084_v34_)[default__.safeIndex(758, (d_3084_v34_).length(0))], (self).fm8(globalState), (self).f32])) + (d_3047_v0_))
            nw504_[int(8)] = _dafny.MultiSet([(d_3084_v34_)[default__.safeIndex(758, (d_3084_v34_).length(0))]])
            nw504_[int(9)] = (d_3079_v29_).set((d_3084_v34_)[default__.safeIndex(758, (d_3084_v34_).length(0))], default__.abs((self).f33))
            nw504_[int(10)] = (d_3154_v87_).f37
            nw504_[int(11)] = (d_3154_v87_).f37
            nw504_[int(12)] = d_3079_v29_
            nw504_[int(13)] = d_3079_v29_
            nw504_[int(14)] = ((d_3154_v87_).f37) | ((d_3154_v87_).f37)
            nw504_[int(15)] = (d_3157_v88_)[default__.safeIndex((0) - ((self).f33), len(d_3157_v88_))]
            nw504_[int(16)] = ((d_3154_v87_).f37).set((self).fm8(globalState), default__.abs((self).f33))
            nw504_[int(17)] = ((d_3154_v87_).f37) - (_dafny.MultiSet([(d_3084_v34_)[default__.safeIndex(758, (d_3084_v34_).length(0))]]))
            nw504_[int(18)] = (d_3154_v87_).f37
            nw504_[int(19)] = d_3079_v29_
            nw504_[int(20)] = d_3079_v29_
            nw504_[int(21)] = ((d_3154_v87_).f37) | (d_3079_v29_)
            nw504_[int(22)] = default__.fm34(globalState)
            nw504_[int(23)] = d_3079_v29_
            nw504_[int(24)] = (d_3154_v87_).f37
            nw504_[int(25)] = (((d_3079_v29_).set((self).f27, default__.abs((self).f33))).set((d_3084_v34_)[default__.safeIndex(758, (d_3084_v34_).length(0))], default__.abs((self).f33))) - ((d_3154_v87_).f37)
            d_3158_v89_ = nw504_
            index498_ = default__.safeIndex(824, (d_3158_v89_).length(0))
            (d_3158_v89_)[index498_] = (d_3154_v87_).f37
            d_3159_v90_: D15
            d_3159_v90_ = D15_DC34((self).f33, (self).f32, (self).f33, (self).f33, False)
            d_3160_v91_: D21
            d_3160_v91_ = D21_DC53((d_3084_v34_)[default__.safeIndex(758, (d_3084_v34_).length(0))], (d_3159_v90_).cf45, not((self).f27))
            index499_ = default__.safeIndex(758, (d_3084_v34_).length(0))
            index500_ = default__.safeIndex(824, (d_3158_v89_).length(0))
            rhs499_ = D9_DC22(d_3151_v84_)
            rhs500_ = d_3154_v87_
            rhs501_ = (self).f27
            rhs502_ = d_3079_v29_
            rhs503_ = d_3160_v91_
            lhs425_ = d_3084_v34_
            lhs426_ = default__.safeIndex(758, (d_3084_v34_).length(0))
            lhs427_ = d_3158_v89_
            lhs428_ = default__.safeIndex(824, (d_3158_v89_).length(0))
            d_3152_v85_ = rhs499_
            d_3154_v87_ = rhs500_
            lhs425_[lhs426_] = rhs501_
            lhs427_[lhs428_] = rhs502_
            d_3160_v91_ = rhs503_
            d_3161_v92_: _dafny.Map
            d_3161_v92_ = _dafny.Map({d_3080_v30_: False})
            d_3161_v92_ = (d_3161_v92_).set((d_3080_v30_).set(default__.safeIndex(639, len(d_3080_v30_)), (d_3080_v30_)[default__.safeIndex((self).f33, len(d_3080_v30_))]), not((_dafny.MultiSet(d_3047_v0_)).isdisjoint(_dafny.MultiSet([not((d_3084_v34_)[default__.safeIndex(758, (d_3084_v34_).length(0))])]))))
            d_3162_v93_: _dafny.Seq
            d_3162_v93_ = _dafny.SeqWithoutIsStrInference([d_3047_v0_, d_3047_v0_, d_3047_v0_])
            index501_ = default__.safeIndex(758, (d_3084_v34_).length(0))
            (d_3084_v34_)[index501_] = (_dafny.SeqWithoutIsStrInference([d_3047_v0_ for d_3163_i4_ in range(default__.abs(83))])) == (d_3162_v93_)
        elif True:
            (globalState).f6 = (self).f33
            (globalState).f14 = (((self).f33) < ((self).f33) if (self).f27 else (d_3084_v34_)[default__.safeIndex(758, (d_3084_v34_).length(0))])
            d_3164_v94_: _dafny.Array
            nw505_ = _dafny.Array(None, 5)
            nw505_[int(0)] = len(_dafny.SeqWithoutIsStrInference([(0) - ((self).f33), (self).f33, (self).f33, (self).f33, (self).f33]))
            nw505_[int(1)] = (self).f33
            nw505_[int(2)] = (self).f33
            nw505_[int(3)] = (self).f33
            nw505_[int(4)] = (self).f33
            d_3164_v94_ = nw505_
            d_3165_v95_: _dafny.Map
            d_3165_v95_ = _dafny.Map({(self).f32: d_3164_v94_})
            d_3166_v96_: D29
            d_3166_v96_ = D29_DC71(980, (d_3083_v33_) | (d_3083_v33_), True, not ((self).f27) or ((d_3084_v34_)[default__.safeIndex(758, (d_3084_v34_).length(0))]), ((d_3165_v95_)[not((self).f32)] if (not((self).f32)) in (d_3165_v95_) else d_3164_v94_))
            source47_ = d_3166_v96_
            if source47_.is_DC70:
                d_3167___mcc_h6_ = source47_.cf110
                d_3168___mcc_h7_ = source47_.cf111
                d_3169___mcc_h8_ = source47_.cf112
                d_3170___mcc_h9_ = source47_.cf113
                d_3171_cf113_ = d_3170___mcc_h9_
                d_3172_cf112_ = d_3169___mcc_h8_
                d_3173_cf111_ = d_3168___mcc_h7_
                d_3174_cf110_ = d_3167___mcc_h6_
                d_3175_v97_: _dafny.Map
                d_3175_v97_ = _dafny.Map({(self).f27: (self).f27})
                d_3176_v98_: _dafny.MultiSet
                d_3176_v98_ = _dafny.MultiSet([d_3175_v97_, d_3175_v97_])
                d_3085_v35_ = (d_3085_v35_).set(default__.safeDivisionInt(d_3173_cf111_, d_3173_cf111_), default__.safeModuloInt(((d_3176_v98_)[_dafny.Map({(d_3084_v34_)[default__.safeIndex(758, (d_3084_v34_).length(0))]: (self).f27})] if (_dafny.Map({(d_3084_v34_)[default__.safeIndex(758, (d_3084_v34_).length(0))]: (self).f27})) in (d_3176_v98_) else (self).f33), (0) - (d_3173_cf111_)))
                index502_ = default__.safeIndex(758, (d_3084_v34_).length(0))
                (d_3084_v34_)[index502_] = (self).f32
                (globalState).f2 = -16
                d_3177_v99_: C7
                nw506_ = C7()
                nw506_.ctor__(((d_3084_v34_)[default__.safeIndex(758, (d_3084_v34_).length(0))] if (self).f27 else (d_3047_v0_)[default__.safeIndex(len(d_3175_v97_), len(d_3047_v0_))]), default__.safeModuloInt(((d_3107_v56_)[(0) - ((self).f33)] if ((0) - ((self).f33)) in (d_3107_v56_) else (self).f33), d_3173_cf111_), 411, (self).f33, True)
                d_3177_v99_ = nw506_
            elif source47_.is_DC71:
                d_3178___mcc_h10_ = source47_.cf114
                d_3179___mcc_h11_ = source47_.cf115
                d_3180___mcc_h12_ = source47_.cf116
                d_3181___mcc_h13_ = source47_.cf117
                d_3182___mcc_h14_ = source47_.cf118
                d_3183_cf118_ = d_3182___mcc_h14_
                d_3184_cf117_ = d_3181___mcc_h13_
                d_3185_cf116_ = d_3180___mcc_h12_
                d_3186_cf115_ = d_3179___mcc_h11_
                d_3187_cf114_ = d_3178___mcc_h10_
                d_3188_v100_: _dafny.Array
                nw507_ = _dafny.Array(_dafny.Map({}), 16)
                d_3188_v100_ = nw507_
                d_3189_v101_: D27
                d_3189_v101_ = D27_DC66(d_3188_v100_)
                d_3190_v102_: D17
                d_3190_v102_ = D17_DC38(d_3084_v34_)
                rhs504_ = (d_3190_v102_).cf53
                rhs505_ = d_3189_v101_
                d_3084_v34_ = rhs504_
                d_3189_v101_ = rhs505_
                d_3191_v103_: _dafny.Map
                d_3191_v103_ = _dafny.Map({True: d_3079_v29_})
                d_3192_v104_: _dafny.Array
                nw508_ = _dafny.Array(None, 28)
                nw508_[int(0)] = d_3164_v94_
                nw508_[int(1)] = d_3183_cf118_
                nw508_[int(2)] = d_3183_cf118_
                nw508_[int(3)] = d_3183_cf118_
                nw508_[int(4)] = d_3164_v94_
                nw508_[int(5)] = d_3183_cf118_
                nw508_[int(6)] = d_3164_v94_
                nw508_[int(7)] = d_3164_v94_
                nw508_[int(8)] = d_3183_cf118_
                nw508_[int(9)] = d_3164_v94_
                nw508_[int(10)] = d_3183_cf118_
                nw508_[int(11)] = d_3183_cf118_
                nw508_[int(12)] = d_3183_cf118_
                nw508_[int(13)] = d_3183_cf118_
                nw508_[int(14)] = d_3164_v94_
                nw508_[int(15)] = d_3164_v94_
                nw508_[int(16)] = d_3183_cf118_
                nw508_[int(17)] = d_3183_cf118_
                nw508_[int(18)] = d_3183_cf118_
                nw508_[int(19)] = d_3164_v94_
                nw508_[int(20)] = d_3183_cf118_
                nw508_[int(21)] = d_3164_v94_
                nw508_[int(22)] = d_3164_v94_
                nw508_[int(23)] = d_3183_cf118_
                nw508_[int(24)] = d_3164_v94_
                nw508_[int(25)] = d_3183_cf118_
                nw508_[int(26)] = d_3164_v94_
                nw508_[int(27)] = d_3164_v94_
                d_3192_v104_ = nw508_
                d_3193_v105_: _dafny.Map
                d_3193_v105_ = _dafny.Map({((d_3191_v103_)[(self).f27] if ((self).f27) in (d_3191_v103_) else d_3079_v29_): d_3192_v104_})
                d_3193_v105_ = (d_3193_v105_).set((d_3079_v29_) - (d_3079_v29_), d_3192_v104_)
                d_3194_v107_: _dafny.Map
                def iife287_():
                    coll117_ = _dafny.Map()
                    compr_117_: int
                    for compr_117_ in (d_3085_v35_).keys.Elements:
                        d_3195_v106_: int = compr_117_
                        if (d_3195_v106_) in (d_3085_v35_):
                            coll117_[default__.safeDivisionInt(d_3195_v106_, 243)] = _dafny.CodePoint('n')
                    return _dafny.Map(coll117_)
                d_3194_v107_ = _dafny.Map({_dafny.CodePoint('r'): iife287_()
                })
                d_3196_v108_: _dafny.Map
                d_3196_v108_ = _dafny.Map({(self).f27: d_3080_v30_})
                d_3197_v109_: _dafny.Array
                nw509_ = _dafny.Array(None, 14)
                nw509_[int(0)] = (self).fm2(d_3187_cf114_, d_3194_v107_, globalState)
                nw509_[int(1)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_3198_i5_ in range(default__.abs(9))])
                nw509_[int(2)] = ((d_3196_v108_)[d_3185_cf116_] if (d_3185_cf116_) in (d_3196_v108_) else d_3080_v30_)
                nw509_[int(3)] = d_3080_v30_
                nw509_[int(4)] = d_3080_v30_
                nw509_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fxy"))
                nw509_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ib"))
                nw509_[int(7)] = d_3080_v30_
                nw509_[int(8)] = d_3080_v30_
                nw509_[int(9)] = d_3080_v30_
                nw509_[int(10)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kjrmwjjl"))
                nw509_[int(11)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_3199_i6_ in range(default__.abs(141))])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_3200_i7_ in range(default__.abs(684))]))
                nw509_[int(12)] = (d_3080_v30_ if d_3185_cf116_ else d_3080_v30_)
                nw509_[int(13)] = d_3080_v30_
                d_3197_v109_ = nw509_
                index503_ = default__.safeIndex(405, (d_3197_v109_).length(0))
                (d_3197_v109_)[index503_] = d_3080_v30_
                index504_ = default__.safeIndex(405, (d_3197_v109_).length(0))
                (d_3197_v109_)[index504_] = (self).fm3(globalState)
                d_3201_v110_: C0
                nw510_ = C0()
                nw510_.ctor__()
                d_3201_v110_ = nw510_
            elif source47_.is_DC69:
                d_3202___mcc_h15_ = source47_.cf109
                d_3203_cf109_ = d_3202___mcc_h15_
                d_3204_v111_: _dafny.Map
                d_3204_v111_ = _dafny.Map({95: d_3164_v94_})
                d_3205_v112_: _dafny.Set
                d_3205_v112_ = _dafny.Set({((d_3204_v111_)[(self).f33] if ((self).f33) in (d_3204_v111_) else d_3164_v94_), d_3164_v94_})
                (globalState).f5 = (d_3164_v94_) in (d_3205_v112_)
                d_3206_v113_: _dafny.Array
                nw511_ = _dafny.Array(_dafny.Set({}), 6)
                d_3206_v113_ = nw511_
                d_3207_v114_: str
                d_3207_v114_ = _dafny.CodePoint('k')
                d_3208_v115_: _dafny.Set
                d_3208_v115_ = _dafny.Set({d_3207_v114_, d_3207_v114_})
                index505_ = default__.safeIndex(891, (d_3206_v113_).length(0))
                (d_3206_v113_)[index505_] = d_3208_v115_
                index506_ = default__.safeIndex(891, (d_3206_v113_).length(0))
                (d_3206_v113_)[index506_] = d_3208_v115_
                (globalState).f6 = (0) - ((0) - ((0) - (default__.safeDivisionInt((self).f33, default__.safeDivisionInt(253, len(d_3080_v30_))))))
                rhs506_ = (self).f32
                rhs507_ = d_3207_v114_
                lhs429_ = globalState
                lhs429_.f21 = rhs506_
                d_3207_v114_ = rhs507_
            elif True:
                d_3209___mcc_h16_ = source47_.cf119
                d_3210_cf119_ = d_3209___mcc_h16_
                (globalState).f25 = (self).f27
                d_3211_v116_: D15
                d_3211_v116_ = D15_DC32(_dafny.Map({d_3084_v34_: (0) - ((self).f33)}))
                rhs508_ = d_3211_v116_
                rhs509_ = (self).f33
                lhs430_ = globalState
                d_3211_v116_ = rhs508_
                lhs430_.f6 = rhs509_
                (globalState).f14 = True
                d_3047_v0_ = _dafny.SeqWithoutIsStrInference([(d_3084_v34_)[default__.safeIndex(758, (d_3084_v34_).length(0))]])
            r0 = ((self).f33) * ((self).f33)
            d_3212_v117_: _dafny.Set
            d_3212_v117_ = _dafny.Set({(self).f27})
            d_3213_v118_: _dafny.Seq
            d_3213_v118_ = _dafny.SeqWithoutIsStrInference([d_3212_v117_])
            d_3214_v119_: C5
            nw512_ = C5()
            nw512_.ctor__((_dafny.SeqWithoutIsStrInference([d_3212_v117_, _dafny.Set({True})])) <= (d_3213_v118_), default__.safeModuloInt(len(_dafny.Map({(self).f27: (self).f33})), len(d_3080_v30_)), default__.safeModuloInt((self).f33, (self).f33), (self).f32)
            d_3214_v119_ = nw512_
        r0 = default__.safeDivisionInt(-4, (self).f33)
        return r0

    def m2(self, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: bool = False
        d_3215_v0_: _dafny.Seq
        d_3215_v0_ = _dafny.SeqWithoutIsStrInference([(self).f33])
        d_3216_v1_: _dafny.Seq
        d_3216_v1_ = _dafny.SeqWithoutIsStrInference([(self).f32])
        d_3217_v2_: _dafny.Map
        d_3217_v2_ = _dafny.Map({(self).f32: (self).f33})
        d_3218_v3_: _dafny.Set
        d_3218_v3_ = _dafny.Set({d_3217_v2_})
        d_3219_v4_: _dafny.Array
        nw513_ = _dafny.Array(int(0), 4)
        d_3219_v4_ = nw513_
        d_3220_v5_: D5
        d_3220_v5_ = D5_DC12((self).f27, (self).f33)
        d_3221_v6_: _dafny.Array
        nw514_ = _dafny.Array(None, 15)
        nw514_[int(0)] = (self).f33
        nw514_[int(1)] = (self).f33
        nw514_[int(2)] = (self).f33
        nw514_[int(3)] = default__.safeModuloInt((self).f33, (self).f33)
        nw514_[int(4)] = (self).f33
        nw514_[int(5)] = (d_3215_v0_)[default__.safeIndex((self).f33, len(d_3215_v0_))]
        nw514_[int(6)] = default__.fm1(False, (d_3216_v1_).set(default__.safeIndex((self).f33, len(d_3216_v1_)), (self).f27), (self).f32, globalState)
        nw514_[int(7)] = 379
        nw514_[int(8)] = (self).f33
        nw514_[int(9)] = (self).f33
        nw514_[int(10)] = (self).f33
        nw514_[int(11)] = default__.fm1((self).f27, d_3216_v1_, (self).f27, globalState)
        nw514_[int(12)] = (0) - (default__.safeDivisionInt(len(d_3218_v3_), (D23_DC61(d_3219_v4_, (self).f27, (self).f33)).cf95))
        nw514_[int(13)] = (self).f33
        nw514_[int(14)] = (d_3220_v5_).cf20
        d_3221_v6_ = nw514_
        index507_ = default__.safeIndex(595, (d_3221_v6_).length(0))
        (d_3221_v6_)[index507_] = (self).f33
        index508_ = default__.safeIndex(595, (d_3221_v6_).length(0))
        (d_3221_v6_)[index508_] = 993
        d_3222_v7_: _dafny.Array
        nw515_ = _dafny.Array(False, 6)
        d_3222_v7_ = nw515_
        index509_ = default__.safeIndex(511, (d_3222_v7_).length(0))
        (d_3222_v7_)[index509_] = (self).f27
        index510_ = default__.safeIndex(511, (d_3222_v7_).length(0))
        (d_3222_v7_)[index510_] = (self).f32
        d_3223_i0_: int
        d_3223_i0_ = 0
        with _dafny.label("18"):
            while not(not((self).f32)):
                with _dafny.c_label("18"):
                    if (d_3223_i0_) >= (100):
                        raise _dafny.Break("18")
                    d_3223_i0_ = (d_3223_i0_) + (1)
                    d_3224_v8_: _dafny.Map
                    d_3224_v8_ = _dafny.Map({len(d_3215_v0_): (self).f33})
                    if (d_3216_v1_)[default__.safeIndex(((d_3224_v8_)[(d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))]] if ((d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))]) in (d_3224_v8_) else (self).f33), len(d_3216_v1_))]:
                        d_3225_v9_: _dafny.Map
                        d_3225_v9_ = _dafny.Map({(self).f32: (self).f27})
                        d_3226_v10_: _dafny.Array
                        nw516_ = _dafny.Array(_dafny.Seq({}), 27)
                        d_3226_v10_ = nw516_
                        index511_ = default__.safeIndex(827, (d_3226_v10_).length(0))
                        (d_3226_v10_)[index511_] = _dafny.SeqWithoutIsStrInference([(d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))]])
                        d_3227_v11_: _dafny.Array
                        nw517_ = _dafny.Array(_dafny.Array(None, 0), 6)
                        d_3227_v11_ = nw517_
                        index512_ = default__.safeIndex(779, (d_3227_v11_).length(0))
                        (d_3227_v11_)[index512_] = d_3222_v7_
                        d_3228_v12_: _dafny.Seq
                        d_3228_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xdbku"))
                        d_3229_v13_: _dafny.Seq
                        d_3229_v13_ = _dafny.SeqWithoutIsStrInference([d_3228_v12_])
                        d_3230_v14_: str
                        d_3230_v14_ = _dafny.CodePoint('h')
                        d_3231_v15_: _dafny.Seq
                        d_3231_v15_ = _dafny.SeqWithoutIsStrInference([d_3222_v7_, d_3222_v7_])
                        d_3232_v16_: _dafny.Map
                        d_3232_v16_ = _dafny.Map({d_3230_v14_: d_3231_v15_})
                        d_3233_v17_: _dafny.Map
                        d_3233_v17_ = _dafny.Map({len(d_3228_v12_): d_3230_v14_})
                        index513_ = default__.safeIndex(595, (d_3221_v6_).length(0))
                        index514_ = default__.safeIndex(827, (d_3226_v10_).length(0))
                        index515_ = default__.safeIndex(779, (d_3227_v11_).length(0))
                        rhs510_ = d_3225_v9_
                        rhs511_ = (0) - (((d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))]) * (len((d_3229_v13_) + (default__.fm18((self).f27, globalState)))))
                        rhs512_ = (self).f32
                        rhs513_ = (default__.fm55(((d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))]) + (len(((d_3232_v16_)[((d_3233_v17_)[(d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))]] if ((d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))]) in (d_3233_v17_) else d_3230_v14_)] if (((d_3233_v17_)[(d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))]] if ((d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))]) in (d_3233_v17_) else d_3230_v14_)) in (d_3232_v16_) else d_3231_v15_))), d_3230_v14_, globalState)).set(default__.safeIndex((d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))], len(default__.fm55(((d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))]) + (len(((d_3232_v16_)[((d_3233_v17_)[(d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))]] if ((d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))]) in (d_3233_v17_) else d_3230_v14_)] if (((d_3233_v17_)[(d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))]] if ((d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))]) in (d_3233_v17_) else d_3230_v14_)) in (d_3232_v16_) else d_3231_v15_))), d_3230_v14_, globalState))), default__.safeDivisionInt((self).f33, (self).f33))
                        rhs514_ = d_3222_v7_
                        lhs431_ = d_3221_v6_
                        lhs432_ = default__.safeIndex(595, (d_3221_v6_).length(0))
                        lhs433_ = globalState
                        lhs434_ = d_3226_v10_
                        lhs435_ = default__.safeIndex(827, (d_3226_v10_).length(0))
                        lhs436_ = d_3227_v11_
                        lhs437_ = default__.safeIndex(779, (d_3227_v11_).length(0))
                        d_3225_v9_ = rhs510_
                        lhs431_[lhs432_] = rhs511_
                        lhs433_.f21 = rhs512_
                        lhs434_[lhs435_] = rhs513_
                        lhs436_[lhs437_] = rhs514_
                        d_3234_v18_: _dafny.Array
                        nw518_ = _dafny.Array(None, 4)
                        d_3234_v18_ = nw518_
                        d_3235_v19_: T2
                        nw519_ = C1()
                        nw519_.ctor__((d_3222_v7_)[default__.safeIndex(511, (d_3222_v7_).length(0))])
                        d_3235_v19_ = nw519_
                        index516_ = default__.safeIndex(771, (d_3234_v18_).length(0))
                        (d_3234_v18_)[index516_] = d_3235_v19_
                        d_3236_v20_: _dafny.Set
                        d_3236_v20_ = _dafny.Set({(d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))], (d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))]})
                        index517_ = default__.safeIndex(771, (d_3234_v18_).length(0))
                        rhs515_ = default__.safeModuloInt(len(d_3236_v20_), (self).f33)
                        rhs516_ = d_3235_v19_
                        rhs517_ = (self).f33
                        lhs438_ = globalState
                        lhs439_ = d_3234_v18_
                        lhs440_ = default__.safeIndex(771, (d_3234_v18_).length(0))
                        lhs441_ = globalState
                        lhs438_.f6 = rhs515_
                        lhs439_[lhs440_] = rhs516_
                        lhs441_.f2 = rhs517_
                        (globalState).f14 = (d_3235_v19_).f27
                        d_3237_v22_: _dafny.Seq
                        def iife288_():
                            coll118_ = _dafny.Map()
                            compr_118_: int
                            for compr_118_ in _dafny.IntegerRange(856, 295):
                                d_3238_v21_: int = compr_118_
                                if ((856) <= (d_3238_v21_)) and ((d_3238_v21_) < (295)):
                                    coll118_[(d_3238_v21_) - ((d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))])] = len(_dafny.Map({(d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))]: (self).f27}))
                            return _dafny.Map(coll118_)
                        d_3237_v22_ = _dafny.SeqWithoutIsStrInference([iife288_()
                        , _dafny.Map({355: (self).f33}), d_3224_v8_])
                        d_3239_v23_: _dafny.Set
                        d_3239_v23_ = _dafny.Set({d_3228_v12_})
                        d_3240_v24_: T1
                        nw520_ = C7()
                        nw520_.ctor__((d_3222_v7_)[default__.safeIndex(511, (d_3222_v7_).length(0))], (self).f33, (d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))], len(d_3239_v23_), (self).f32)
                        d_3240_v24_ = nw520_
                        d_3241_v25_: _dafny.Map
                        d_3241_v25_ = _dafny.Map({d_3240_v24_: _dafny.CodePoint('y')})
                        index518_ = default__.safeIndex(595, (d_3221_v6_).length(0))
                        rhs518_ = ((d_3217_v2_)[((self).f32) or (not(not((self).fm4(default__.fm46(default__.fm34(globalState), (d_3240_v24_).f27, globalState), globalState))))] if (((self).f32) or (not(not((self).fm4(default__.fm46(default__.fm34(globalState), (d_3240_v24_).f27, globalState), globalState))))) in (d_3217_v2_) else (self).f33)
                        rhs519_ = ((_dafny.SeqWithoutIsStrInference([d_3224_v8_ for d_3242_i1_ in range(default__.abs(809))])) + (d_3237_v22_)) + (default__.fm69((d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))], _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jpet")), (self).f27, globalState))
                        rhs520_ = d_3241_v25_
                        rhs521_ = d_3221_v6_
                        lhs442_ = d_3221_v6_
                        lhs443_ = default__.safeIndex(595, (d_3221_v6_).length(0))
                        lhs442_[lhs443_] = rhs518_
                        d_3237_v22_ = rhs519_
                        d_3241_v25_ = rhs520_
                        d_3221_v6_ = rhs521_
                        d_3243_v26_: _dafny.MultiSet
                        d_3243_v26_ = _dafny.MultiSet([(d_3240_v24_).f28, (d_3240_v24_).f29])
                        d_3244_v27_: _dafny.Map
                        d_3244_v27_ = _dafny.Map({d_3243_v26_: (d_3222_v7_)[default__.safeIndex(511, (d_3222_v7_).length(0))]})
                        d_3245_v28_: D13
                        d_3245_v28_ = D13_DC28(175)
                        d_3246_v29_: D15
                        d_3246_v29_ = D15_DC34(len(d_3244_v27_), (self).f27, -302, (d_3245_v28_).cf38, (d_3235_v19_).f27)
                        d_3246_v29_ = default__.fm70(globalState)
                    elif True:
                        d_3247_v30_: str
                        d_3247_v30_ = _dafny.CodePoint('d')
                        d_3248_v31_: _dafny.MultiSet
                        d_3248_v31_ = _dafny.MultiSet([d_3247_v30_])
                        d_3249_v32_: _dafny.Map
                        d_3249_v32_ = _dafny.Map({D15_DC34(((d_3248_v31_)[d_3247_v30_] if (d_3247_v30_) in (d_3248_v31_) else 159), (self).f27, (d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))], (d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))], (self).f27): (self).f33})
                        def iife289_(_pat_let85_0):
                            def iife290_(d_3250_dt__update__tmp_h0_):
                                def iife291_(_pat_let86_0):
                                    def iife292_(d_3251_dt__update_hcf49_h0_):
                                        return D15_DC34((d_3250_dt__update__tmp_h0_).cf45, (d_3250_dt__update__tmp_h0_).cf46, (d_3250_dt__update__tmp_h0_).cf47, (d_3250_dt__update__tmp_h0_).cf48, d_3251_dt__update_hcf49_h0_)
                                    return iife292_(_pat_let86_0)
                                return iife291_(not((self).f27))
                            return iife290_(_pat_let85_0)
                        d_3249_v32_ = (d_3249_v32_).set(iife289_(default__.fm70(globalState)), len(_dafny.Set({(d_3222_v7_)[default__.safeIndex(511, (d_3222_v7_).length(0))]})))
                        d_3252_v33_: _dafny.Map
                        d_3252_v33_ = _dafny.Map({(d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))]: (d_3222_v7_)[default__.safeIndex(511, (d_3222_v7_).length(0))]})
                        d_3253_v34_: _dafny.Map
                        d_3253_v34_ = _dafny.Map({(self).f32: d_3252_v33_})
                        d_3254_v35_: D0
                        d_3254_v35_ = D0_DC1((self).f27)
                        d_3255_v37_: _dafny.Set
                        def iife293_():
                            coll119_ = _dafny.Map()
                            compr_119_: int
                            for compr_119_ in _dafny.IntegerRange(-780, 153):
                                d_3256_v36_: int = compr_119_
                                if ((-780) <= (d_3256_v36_)) and ((d_3256_v36_) < (153)):
                                    coll119_[(d_3256_v36_) - ((d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))])] = (self).f27
                            return _dafny.Map(coll119_)
                        d_3255_v37_ = _dafny.Set({d_3252_v33_, ((d_3253_v34_)[(self).fm4(d_3254_v35_, globalState)] if ((self).fm4(d_3254_v35_, globalState)) in (d_3253_v34_) else d_3252_v33_), iife293_()
                        , d_3252_v33_})
                        (globalState).f25 = (d_3252_v33_) not in (d_3255_v37_)
                        d_3257_v38_: str
                        d_3257_v38_ = _dafny.CodePoint('y')
                        d_3258_v39_: _dafny.Map
                        d_3258_v39_ = _dafny.Map({(self).f33: d_3257_v38_})
                        d_3259_v40_: _dafny.Set
                        d_3259_v40_ = _dafny.Set({d_3258_v39_})
                        d_3260_v41_: _dafny.Set
                        d_3260_v41_ = d_3259_v40_
                        d_3261_v42_: _dafny.Map
                        d_3261_v42_ = _dafny.Map({d_3258_v39_: (self).f27})
                        d_3262_v43_: D32
                        d_3262_v43_ = D32_DC78(d_3261_v42_)
                        d_3263_v45_: _dafny.Seq
                        def iife294_():
                            coll120_ = _dafny.Set()
                            compr_120_: _dafny.Map
                            for compr_120_ in ((d_3262_v43_).cf129).keys.Elements:
                                d_3264_v44_: _dafny.Map = compr_120_
                                if (d_3264_v44_) in ((d_3262_v43_).cf129):
                                    coll120_ = coll120_.union(_dafny.Set([d_3264_v44_]))
                            return _dafny.Set(coll120_)
                        d_3263_v45_ = _dafny.SeqWithoutIsStrInference([d_3260_v41_, d_3260_v41_, d_3260_v41_, iife294_()
                        , d_3260_v41_])
                        d_3265_v46_: D30
                        d_3265_v46_ = D30_DC75((d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))], _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_3266_i2_ in range(default__.abs(313))]), d_3263_v45_, d_3222_v7_)
                        pat_let_tv68_ = d_3222_v7_
                        pat_let_tv69_ = d_3263_v45_
                        def iife295_(_pat_let87_0):
                            def iife296_(d_3267_dt__update__tmp_h1_):
                                def iife297_(_pat_let88_0):
                                    def iife298_(d_3268_dt__update_hcf127_h0_):
                                        def iife299_(_pat_let89_0):
                                            def iife300_(d_3269_dt__update_hcf126_h0_):
                                                return D30_DC75((d_3267_dt__update__tmp_h1_).cf124, (d_3267_dt__update__tmp_h1_).cf125, d_3269_dt__update_hcf126_h0_, d_3268_dt__update_hcf127_h0_)
                                            return iife300_(_pat_let89_0)
                                        return iife299_(pat_let_tv69_)
                                    return iife298_(_pat_let88_0)
                                return iife297_(pat_let_tv68_)
                            return iife296_(_pat_let87_0)
                        d_3222_v7_ = (iife295_(d_3265_v46_)).cf127
                        d_3270_v47_: _dafny.Seq
                        d_3270_v47_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t')])
                        d_3271_v48_: _dafny.MultiSet
                        d_3271_v48_ = _dafny.MultiSet([len(d_3270_v47_), (d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))], (self).f33, len(_dafny.SeqWithoutIsStrInference([d_3257_v38_ for d_3272_i3_ in range(default__.abs(966))]))])
                        d_3273_v49_: C7
                        nw521_ = C7()
                        nw521_.ctor__((self).f32, (self).f33, (d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))], (d_3271_v48_).cardinality, (self).f32)
                        d_3273_v49_ = nw521_
                        d_3257_v38_ = d_3257_v38_
                    d_3274_v50_: D6
                    d_3274_v50_ = D6_DC13(d_3216_v1_)
                    d_3275_v51_: _dafny.Seq
                    d_3275_v51_ = _dafny.SeqWithoutIsStrInference([d_3274_v50_, d_3274_v50_])
                    d_3276_v52_: _dafny.Array
                    nw522_ = _dafny.Array(_dafny.Map({}), 27)
                    d_3276_v52_ = nw522_
                    d_3277_v53_: _dafny.Map
                    d_3277_v53_ = _dafny.Map({(self).f32: d_3222_v7_})
                    index519_ = default__.safeIndex(571, (d_3276_v52_).length(0))
                    (d_3276_v52_)[index519_] = (_dafny.Map({(d_3222_v7_)[default__.safeIndex(511, (d_3222_v7_).length(0))]: d_3222_v7_})) | (d_3277_v53_)
                    d_3278_v54_: _dafny.Set
                    d_3278_v54_ = _dafny.Set({not((self).f27)})
                    index520_ = default__.safeIndex(595, (d_3221_v6_).length(0))
                    index521_ = default__.safeIndex(571, (d_3276_v52_).length(0))
                    rhs522_ = -204
                    rhs523_ = len(_dafny.Map({(d_3222_v7_)[default__.safeIndex(511, (d_3222_v7_).length(0))]: _dafny.Map({len(d_3278_v54_): (self).f33})}))
                    rhs524_ = default__.fm71((d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))], (d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))], globalState)
                    rhs525_ = _dafny.Map({(self).f32: d_3222_v7_})
                    lhs444_ = globalState
                    lhs445_ = d_3221_v6_
                    lhs446_ = default__.safeIndex(595, (d_3221_v6_).length(0))
                    lhs447_ = d_3276_v52_
                    lhs448_ = default__.safeIndex(571, (d_3276_v52_).length(0))
                    lhs444_.f6 = rhs522_
                    lhs445_[lhs446_] = rhs523_
                    d_3275_v51_ = rhs524_
                    lhs447_[lhs448_] = rhs525_
                    (globalState).f2 = default__.safeDivisionInt((881) * ((d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))]), ((d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))]) - ((self).f33))
                    index522_ = default__.safeIndex(571, (d_3276_v52_).length(0))
                    (d_3276_v52_)[index522_] = (d_3276_v52_)[default__.safeIndex(571, (d_3276_v52_).length(0))]
                    pass
            pass
        hi20_ = 745
        for d_3279_i4_ in range((default__.fm1((d_3222_v7_)[default__.safeIndex(511, (d_3222_v7_).length(0))], _dafny.SeqWithoutIsStrInference([(self).f27, True, (self).f32]), (self).f27, globalState) if (d_3222_v7_)[default__.safeIndex(511, (d_3222_v7_).length(0))] else (self).f33), hi20_):
            d_3280_v55_: _dafny.Seq
            d_3280_v55_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aar"))
            d_3281_v56_: _dafny.Map
            d_3281_v56_ = _dafny.Map({d_3280_v55_: (self).f27})
            d_3282_v57_: C3
            nw523_ = C3()
            nw523_.ctor__((d_3222_v7_)[default__.safeIndex(511, (d_3222_v7_).length(0))], ((d_3281_v56_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yjcghnj"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yjcghnj"))) in (d_3281_v56_) else (self).f27))
            d_3282_v57_ = nw523_
            if ((d_3222_v7_)[default__.safeIndex(511, (d_3222_v7_).length(0))] if ((self).f33) == ((self).f33) else (self).f32):
                (globalState).f6 = len(_dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f33]): -381}))
                d_3283_v58_: C7
                nw524_ = C7()
                nw524_.ctor__(d_3282_v57_.f38, (0) - (d_3279_i4_), (d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))], len(d_3280_v55_), (d_3222_v7_)[default__.safeIndex(511, (d_3222_v7_).length(0))])
                d_3283_v58_ = nw524_
                d_3284_v59_: _dafny.Map
                d_3284_v59_ = _dafny.Map({d_3283_v58_: (self).f27})
                d_3285_v60_: C6
                nw525_ = C6()
                nw525_.ctor__(d_3279_i4_, (self).f33, d_3282_v57_.f38)
                d_3285_v60_ = nw525_
                d_3286_v61_: _dafny.Map
                d_3286_v61_ = _dafny.Map({(self).f27: d_3285_v60_})
                d_3287_v62_: D2
                d_3287_v62_ = D2_DC6(len(d_3286_v61_), d_3215_v0_, (d_3222_v7_)[default__.safeIndex(511, (d_3222_v7_).length(0))], d_3215_v0_, (d_3222_v7_)[default__.safeIndex(511, (d_3222_v7_).length(0))])
                d_3288_v63_: _dafny.Set
                d_3288_v63_ = _dafny.Set({(d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))], len(((d_3284_v59_).set(d_3283_v58_, (self).f27)).set(d_3283_v58_, (self).f27)), len((d_3287_v62_).cf8), len(d_3216_v1_), d_3279_i4_})
                index523_ = default__.safeIndex(511, (d_3222_v7_).length(0))
                (d_3222_v7_)[index523_] = not(((d_3288_v63_) - (_dafny.Set({(d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))]}))).ispropersubset(d_3288_v63_))
                d_3289_v64_: C0
                nw526_ = C0()
                nw526_.ctor__()
                d_3289_v64_ = nw526_
                (d_3282_v57_).f38 = d_3283_v58_.f35
                (d_3282_v57_).f38 = (self).f27
            elif True:
                d_3290_v65_: D9
                d_3290_v65_ = D9_DC22(D9_DC21(d_3282_v57_.f38, d_3279_i4_))
                d_3291_v66_: D9
                d_3291_v66_ = D9_DC22(d_3290_v65_)
                d_3292_v67_: D9
                d_3292_v67_ = D9_DC22(d_3291_v66_)
                pat_let_tv70_ = d_3290_v65_
                d_3293_v68_: _dafny.Map
                def iife301_(_pat_let90_0):
                    def iife302_(d_3294_dt__update__tmp_h2_):
                        def iife303_(_pat_let91_0):
                            def iife304_(d_3295_dt__update_hcf33_h0_):
                                return D9_DC22(d_3295_dt__update_hcf33_h0_)
                            return iife304_(_pat_let91_0)
                        return iife303_(pat_let_tv70_)
                    return iife302_(_pat_let90_0)
                d_3293_v68_ = _dafny.Map({d_3282_v57_.f38: (iife301_(d_3292_v67_) if d_3282_v57_.f38 else d_3292_v67_)})
                d_3293_v68_ = (d_3293_v68_).set((d_3222_v7_)[default__.safeIndex(511, (d_3222_v7_).length(0))], d_3292_v67_)
                index524_ = default__.safeIndex(595, (d_3221_v6_).length(0))
                def iife305_(_pat_let92_0):
                    def iife306_(d_3296_dt__update__tmp_h3_):
                        def iife307_(_pat_let93_0):
                            def iife308_(d_3297_dt__update_hcf20_h0_):
                                return D5_DC12((d_3296_dt__update__tmp_h3_).cf19, d_3297_dt__update_hcf20_h0_)
                            return iife308_(_pat_let93_0)
                        return iife307_(251)
                    return iife306_(_pat_let92_0)
                (d_3221_v6_)[index524_] = default__.safeModuloInt(((self).f33) - (len(d_3280_v55_)), ((self).f33) + ((iife305_(d_3220_v5_)).cf20))
                d_3298_v70_: _dafny.Set
                d_3298_v70_ = _dafny.Set({False})
                d_3299_v71_: C5
                nw527_ = C5()
                def iife309_():
                    coll121_ = _dafny.Set()
                    compr_121_: int
                    for compr_121_ in (_dafny.SeqWithoutIsStrInference([(self).f33, len(d_3280_v55_)])).Elements:
                        d_3300_v69_: int = compr_121_
                        if (d_3300_v69_) in (_dafny.SeqWithoutIsStrInference([(self).f33, len(d_3280_v55_)])):
                            coll121_ = coll121_.union(_dafny.Set([default__.safeModuloInt(d_3300_v69_, 217)]))
                    return _dafny.Set(coll121_)
                nw527_.ctor__((d_3282_v57_.f38) or ((self).f32), (0) - (default__.safeModuloInt((d_3215_v0_)[default__.safeIndex(-907, len(d_3215_v0_))], len(iife309_()
                ))), len(d_3298_v70_), d_3282_v57_.f38)
                d_3299_v71_ = nw527_
                r3 = not((len((d_3280_v55_) + (d_3280_v55_))) < (((d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))]) * ((self).f33)))
                index525_ = default__.safeIndex(511, (d_3222_v7_).length(0))
                rhs526_ = (self).f32
                rhs527_ = d_3222_v7_
                rhs528_ = (d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))]
                lhs449_ = d_3222_v7_
                lhs450_ = default__.safeIndex(511, (d_3222_v7_).length(0))
                lhs451_ = globalState
                lhs449_[lhs450_] = rhs526_
                d_3222_v7_ = rhs527_
                lhs451_.f6 = rhs528_
            (globalState).f12 = not((self).fm8(globalState))
            d_3301_v72_: D19
            d_3301_v72_ = D19_DC46((self).f33, d_3219_v4_, not(not(((d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))]) < (-880))))
            pat_let_tv71_ = d_3221_v6_
            index526_ = default__.safeIndex(595, (d_3221_v6_).length(0))
            def iife310_(_pat_let94_0):
                def iife311_(d_3302_dt__update__tmp_h4_):
                    def iife312_(_pat_let95_0):
                        def iife313_(d_3303_dt__update_hcf65_h0_):
                            return D19_DC46((d_3302_dt__update__tmp_h4_).cf64, d_3303_dt__update_hcf65_h0_, (d_3302_dt__update__tmp_h4_).cf66)
                        return iife313_(_pat_let95_0)
                    return iife312_(pat_let_tv71_)
                return iife311_(_pat_let94_0)
            rhs529_ = d_3282_v57_.f38
            rhs530_ = iife310_(d_3301_v72_)
            rhs531_ = ((self).f33) - (default__.safeModuloInt((self).f33, len(d_3280_v55_)))
            lhs452_ = globalState
            lhs453_ = d_3221_v6_
            lhs454_ = default__.safeIndex(595, (d_3221_v6_).length(0))
            lhs452_.f14 = rhs529_
            d_3301_v72_ = rhs530_
            lhs453_[lhs454_] = rhs531_
        (globalState).f6 = len(d_3215_v0_)
        d_3304_i5_: int
        d_3304_i5_ = 0
        with _dafny.label("19"):
            while (d_3222_v7_)[default__.safeIndex(511, (d_3222_v7_).length(0))]:
                with _dafny.c_label("19"):
                    if (d_3304_i5_) >= (100):
                        raise _dafny.Break("19")
                    d_3304_i5_ = (d_3304_i5_) + (1)
                    d_3305_v73_: _dafny.Seq
                    d_3305_v73_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ruxhpuhjs"))
                    d_3306_v74_: _dafny.Seq
                    d_3306_v74_ = _dafny.SeqWithoutIsStrInference([d_3305_v73_])
                    d_3307_v75_: D13
                    d_3307_v75_ = D13_DC29((d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))], (_dafny.MultiSet(d_3216_v1_)).cardinality, (d_3222_v7_)[default__.safeIndex(511, (d_3222_v7_).length(0))])
                    d_3308_v76_: _dafny.Map
                    d_3308_v76_ = _dafny.Map({(d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))]: (d_3307_v75_).cf41})
                    d_3309_v77_: _dafny.MultiSet
                    d_3309_v77_ = _dafny.MultiSet([(d_3222_v7_)[default__.safeIndex(511, (d_3222_v7_).length(0))], ((d_3308_v76_)[(self).f33] if ((self).f33) in (d_3308_v76_) else (d_3222_v7_)[default__.safeIndex(511, (d_3222_v7_).length(0))])])
                    d_3310_v78_: _dafny.MultiSet
                    d_3310_v78_ = _dafny.MultiSet([(self).f33, (self).f33])
                    d_3311_v79_: C7
                    nw528_ = C7()
                    nw528_.ctor__(not((self).f32), (d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))], (len(d_3306_v74_) if (self).f27 else (d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))]), ((d_3309_v77_)[(self).f27] if ((self).f27) in (d_3309_v77_) else (d_3310_v78_).cardinality), (self).f32)
                    d_3311_v79_ = nw528_
                    index527_ = default__.safeIndex(511, (d_3222_v7_).length(0))
                    (d_3222_v7_)[index527_] = (self).f27
                    d_3312_v80_: _dafny.Seq
                    d_3312_v80_ = _dafny.SeqWithoutIsStrInference([d_3215_v0_])
                    index528_ = default__.safeIndex(595, (d_3221_v6_).length(0))
                    (d_3221_v6_)[index528_] = default__.safeDivisionInt(default__.safeModuloInt(len((d_3312_v80_)[default__.safeIndex((d_3311_v79_).f36, len(d_3312_v80_))]), len(d_3215_v0_)), (self).f33)
                    (globalState).f25 = (self).f27
                    pass
            pass
        d_3313_v81_: _dafny.Array
        nw529_ = _dafny.Array(_dafny.Array(None, 0), 25)
        d_3313_v81_ = nw529_
        r0 = d_3313_v81_
        d_3314_v82_: _dafny.Seq
        d_3314_v82_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sphgobk"))
        d_3315_v83_: str
        d_3315_v83_ = _dafny.CodePoint('d')
        d_3316_v85_: _dafny.Map
        def iife314_():
            coll122_ = _dafny.Map()
            compr_122_: int
            for compr_122_ in (_dafny.MultiSet([len(d_3216_v1_), (self).f33])).Elements:
                d_3317_v84_: int = compr_122_
                if (d_3317_v84_) in (_dafny.MultiSet([len(d_3216_v1_), (self).f33])):
                    coll122_[default__.safeDivisionInt(d_3317_v84_, (_dafny.MultiSet([(self).f27])).cardinality)] = d_3315_v83_
            return _dafny.Map(coll122_)
        d_3316_v85_ = _dafny.Map({d_3315_v83_: iife314_()
        })
        r1 = (self).fm2(len(d_3314_v82_), d_3316_v85_, globalState)
        d_3318_v86_: _dafny.Array
        nw530_ = _dafny.Array(_dafny.Array(None, 0), 21)
        d_3318_v86_ = nw530_
        d_3319_v87_: D5
        d_3319_v87_ = D5_DC11(d_3318_v86_)
        r2 = (d_3319_v87_).cf18
        d_3320_v88_: D18
        d_3320_v88_ = D18_DC44((self).f27, (d_3221_v6_)[default__.safeIndex(595, (d_3221_v6_).length(0))], (d_3222_v7_)[default__.safeIndex(511, (d_3222_v7_).length(0))])
        r3 = (d_3320_v88_).cf62
        return r0, r1, r2, r3

    def m5(self, p0, globalState):
        r0: _dafny.Set = _dafny.Set({})
        d_3321_v0_: _dafny.Seq
        d_3321_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cnyrmgs"))
        d_3322_v1_: _dafny.Map
        d_3322_v1_ = _dafny.Map({(self).f33: d_3321_v0_})
        hi21_ = (self).f33
        for d_3323_i0_ in range((len(((d_3322_v1_)[(self).f33] if ((self).f33) in (d_3322_v1_) else d_3321_v0_))) * ((self).f33), hi21_):
            d_3324_v2_: _dafny.Array
            def lambda135_(d_3325_i1_):
                return (d_3325_i1_) - ((self).f33)

            init79_ = lambda135_
            nw531_ = _dafny.Array(None, 8)
            for i0_79_ in range(nw531_.length(0)):
                nw531_[i0_79_] = init79_(i0_79_)
            d_3324_v2_ = nw531_
            d_3326_v3_: _dafny.Seq
            d_3326_v3_ = _dafny.SeqWithoutIsStrInference([d_3323_i0_])
            d_3327_v4_: _dafny.Seq
            d_3327_v4_ = _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "whc")))), len(d_3326_v3_)])
            d_3328_v5_: _dafny.Map
            d_3328_v5_ = _dafny.Map({d_3324_v2_: len(d_3327_v4_)})
            d_3328_v5_ = d_3328_v5_
            (globalState).f6 = (self).f33
            d_3329_v6_: _dafny.Seq
            d_3329_v6_ = _dafny.SeqWithoutIsStrInference([(self).f27, (self).f32])
            d_3329_v6_ = d_3329_v6_
            d_3330_v7_: _dafny.Array
            nw532_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 4)
            d_3330_v7_ = nw532_
            index529_ = default__.safeIndex(529, (d_3330_v7_).length(0))
            (d_3330_v7_)[index529_] = d_3321_v0_
            index530_ = default__.safeIndex(529, (d_3330_v7_).length(0))
            (d_3330_v7_)[index530_] = d_3321_v0_
        d_3331_v8_: D21
        d_3331_v8_ = D21_DC55(857)
        hi22_ = (d_3331_v8_).cf82
        for d_3332_i2_ in range((self).f33, hi22_):
            d_3333_v9_: _dafny.Seq
            d_3333_v9_ = _dafny.SeqWithoutIsStrInference([841, (self).f33, d_3332_i2_, d_3332_i2_])
            d_3334_v10_: str
            d_3334_v10_ = _dafny.CodePoint('o')
            d_3335_v11_: T0
            nw533_ = C4()
            nw533_.ctor__(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f27])), (self).f32)
            d_3335_v11_ = nw533_
            d_3336_v13_: D20
            def iife315_():
                coll123_ = _dafny.Set()
                compr_123_: int
                for compr_123_ in _dafny.IntegerRange(208, 897):
                    d_3337_v12_: int = compr_123_
                    if ((208) <= (d_3337_v12_)) and ((d_3337_v12_) < (897)):
                        coll123_ = coll123_.union(_dafny.Set([(d_3337_v12_) + (-642)]))
                return _dafny.Set(coll123_)
            d_3336_v13_ = D20_DC49(d_3334_v10_, d_3332_i2_, (self).f33, d_3335_v11_, iife315_()
)
            (globalState).f2 = (d_3333_v9_)[default__.safeIndex((d_3336_v13_).cf71, len(d_3333_v9_))]
            d_3338_v14_: _dafny.Map
            d_3338_v14_ = _dafny.Map({d_3321_v0_: (self).f32})
            d_3339_v15_: _dafny.Seq
            d_3339_v15_ = _dafny.SeqWithoutIsStrInference([(self).f32, (d_3335_v11_).f27])
            d_3340_v16_: _dafny.Set
            d_3340_v16_ = _dafny.Set({len(d_3339_v15_), d_3332_i2_, d_3332_i2_, (d_3333_v9_)[default__.safeIndex(default__.fm1((self).f27, _dafny.SeqWithoutIsStrInference([(d_3335_v11_).f27]), (self).f27, globalState), len(d_3333_v9_))]})
            d_3341_v17_: _dafny.Map
            d_3341_v17_ = _dafny.Map({d_3321_v0_: default__.fm0(d_3321_v0_, (self).f27, (self).f27, d_3340_v16_, globalState)})
            d_3338_v14_ = (d_3341_v17_)
            d_3342_v18_: C2
            nw534_ = C2()
            nw534_.ctor__((d_3335_v11_).f27, (self).f27)
            d_3342_v18_ = nw534_
            (globalState).f2 = 602
        d_3343_v19_: str
        d_3343_v19_ = _dafny.CodePoint('n')
        d_3343_v19_ = d_3343_v19_
        d_3344_v21_: D20
        def iife316_():
            coll124_ = _dafny.Map()
            compr_124_: D3
            for compr_124_ in (_dafny.SeqWithoutIsStrInference([D3_DC8(_dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f33, (self).f33]): (self).f33})) for d_3345_i3_ in range(default__.abs(-435))])).Elements:
                d_3346_v20_: D3 = compr_124_
                if (d_3346_v20_) in (_dafny.SeqWithoutIsStrInference([D3_DC8(_dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f33, (self).f33]): (self).f33})) for d_3345_i3_ in range(default__.abs(-435))])):
                    coll124_[d_3346_v20_] = _dafny.MultiSet([(self).f33, (self).f33, (self).f33])
            return _dafny.Map(coll124_)
        d_3344_v21_ = D20_DC48(iife316_()
)
        d_3347_v22_: _dafny.Seq
        d_3347_v22_ = _dafny.SeqWithoutIsStrInference([d_3344_v21_])
        d_3348_v23_: _dafny.Seq
        d_3348_v23_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_3344_v21_]), d_3347_v22_])
        d_3348_v23_ = (d_3348_v23_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_3343_v19_ for d_3349_i4_ in range(default__.abs(659))])), len(d_3348_v23_)), d_3347_v22_)
        d_3350_v24_: _dafny.MultiSet
        d_3350_v24_ = _dafny.MultiSet([d_3343_v19_])
        d_3351_v25_: _dafny.Seq
        d_3351_v25_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f33, (self).f33, (self).f33]))])
        (globalState).f6 = (default__.safeDivisionInt((self).f33, (d_3350_v24_).cardinality) if (self).f32 else (d_3351_v25_)[default__.safeIndex((self).f33, len(d_3351_v25_))])
        d_3352_v26_: D0
        d_3352_v26_ = D0_DC1((self).f32)
        d_3353_v27_: D2
        d_3353_v27_ = D2_DC6(len(_dafny.SeqWithoutIsStrInference([d_3343_v19_ for d_3354_i5_ in range(default__.abs(-34))])), d_3351_v25_, (self).f32, d_3351_v25_, (self).f32)
        d_3355_v28_: _dafny.Map
        d_3355_v28_ = _dafny.Map({(self).f32: (default__.fm72(328, (self).f33, (self).f33, (self).f32, globalState) if (self).fm4(d_3352_v26_, globalState) else d_3353_v27_)})
        d_3356_v29_: _dafny.MultiSet
        d_3356_v29_ = _dafny.MultiSet([(self).f32, (self).f32])
        d_3355_v28_ = (d_3355_v28_).set((self).f32, D2_DC6((d_3356_v29_).cardinality, _dafny.SeqWithoutIsStrInference([(self).f33, (self).f33]), True, d_3351_v25_, (self).f27))
        d_3357_v30_: _dafny.Map
        d_3357_v30_ = _dafny.Map({(d_3321_v0_).set(default__.safeIndex((self).f33, len(d_3321_v0_)), d_3343_v19_): d_3321_v0_})
        d_3358_v31_: _dafny.Set
        d_3358_v31_ = _dafny.Set({d_3321_v0_, d_3321_v0_, _dafny.SeqWithoutIsStrInference([d_3343_v19_, d_3343_v19_]), ((self).fm3(globalState)) + (d_3321_v0_), ((d_3357_v30_)[d_3321_v0_] if (d_3321_v0_) in (d_3357_v30_) else d_3321_v0_)})
        r0 = d_3358_v31_
        return r0

    def m6(self, globalState):
        r0: int = int(0)
        r1: int = int(0)
        d_3359_v0_: _dafny.Map
        d_3359_v0_ = _dafny.Map({(self).f33: (self).f33})
        d_3360_v1_: _dafny.Set
        d_3360_v1_ = _dafny.Set({(self).f33, ((d_3359_v0_)[(0) - ((self).f33)] if ((0) - ((self).f33)) in (d_3359_v0_) else (self).f33)})
        d_3361_v2_: _dafny.Array
        def lambda136_(d_3362_i0_):
            return (d_3362_i0_) - ((self).f33)

        init80_ = lambda136_
        nw535_ = _dafny.Array(None, 26)
        for i0_80_ in range(nw535_.length(0)):
            nw535_[i0_80_] = init80_(i0_80_)
        d_3361_v2_ = nw535_
        d_3363_v3_: D0
        d_3363_v3_ = D0_DC1((D29_DC71((self).f33, d_3360_v1_, (self).f27, (self).f27, d_3361_v2_)).cf117)
        (globalState).f21 = (d_3363_v3_).cf1
        d_3361_v2_ = d_3361_v2_
        index531_ = default__.safeIndex(206, (d_3361_v2_).length(0))
        (d_3361_v2_)[index531_] = 715
        index532_ = default__.safeIndex(206, (d_3361_v2_).length(0))
        (d_3361_v2_)[index532_] = (self).f33
        if (D21_DC54((self).f32, (d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))], (d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))])).cf79:
            d_3364_v4_: str
            d_3364_v4_ = _dafny.CodePoint('b')
            d_3364_v4_ = d_3364_v4_
            d_3365_v5_: _dafny.Seq
            d_3365_v5_ = _dafny.SeqWithoutIsStrInference([(self).f32])
            d_3366_v6_: _dafny.Seq
            d_3366_v6_ = _dafny.SeqWithoutIsStrInference([d_3361_v2_])
            d_3367_v7_: _dafny.Map
            d_3367_v7_ = _dafny.Map({d_3366_v6_: (self).f33})
            d_3368_v8_: _dafny.Seq
            d_3368_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ukrw"))
            (globalState).f21 = (True) == (((self).f32 if (self).f32 else (d_3365_v5_)[default__.safeIndex(((d_3367_v7_)[(_dafny.SeqWithoutIsStrInference([d_3361_v2_])).set(default__.safeIndex((self).f33, len(_dafny.SeqWithoutIsStrInference([d_3361_v2_]))), d_3361_v2_)] if ((_dafny.SeqWithoutIsStrInference([d_3361_v2_])).set(default__.safeIndex((self).f33, len(_dafny.SeqWithoutIsStrInference([d_3361_v2_]))), d_3361_v2_)) in (d_3367_v7_) else len(d_3368_v8_)), len(d_3365_v5_))]))
            (globalState).f2 = (d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))]
            d_3369_v9_: _dafny.Map
            d_3369_v9_ = _dafny.Map({(self).f33: (self).f32})
            d_3370_v10_: D7
            d_3370_v10_ = D7_DC15(d_3369_v9_)
            d_3371_v11_: D7
            d_3371_v11_ = D7_DC18(d_3370_v10_)
            source48_ = d_3371_v11_
            if source48_.is_DC16:
                (globalState).f6 = (_dafny.MultiSet(d_3365_v5_)).cardinality
                d_3372_v12_: _dafny.Array
                def lambda137_(d_3373_v4_):
                    def lambda138_(d_3374_i1_):
                        return d_3373_v4_

                    return lambda138_

                init81_ = lambda137_(d_3364_v4_)
                nw536_ = _dafny.Array(None, 18)
                for i0_81_ in range(nw536_.length(0)):
                    nw536_[i0_81_] = init81_(i0_81_)
                d_3372_v12_ = nw536_
                d_3375_v13_: _dafny.MultiSet
                d_3375_v13_ = _dafny.MultiSet([(d_3365_v5_)[default__.safeIndex((self).f33, len(d_3365_v5_))], True, (self).f27, True, (self).f32])
                rhs532_ = d_3372_v12_
                rhs533_ = (default__.fm34(globalState)).isdisjoint(d_3375_v13_)
                rhs534_ = (d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))]
                lhs455_ = globalState
                lhs456_ = globalState
                d_3372_v12_ = rhs532_
                lhs455_.f14 = rhs533_
                lhs456_.f6 = rhs534_
                r0 = default__.safeDivisionInt((self).f33, ((self).f33) - ((0) - ((d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))])))
                d_3376_v14_: _dafny.Array
                nw537_ = _dafny.Array(False, 27)
                d_3376_v14_ = nw537_
                index533_ = default__.safeIndex(35, (d_3376_v14_).length(0))
                (d_3376_v14_)[index533_] = (self).f27
                index534_ = default__.safeIndex(35, (d_3376_v14_).length(0))
                (d_3376_v14_)[index534_] = not(((d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))]) < ((self).f33))
            elif source48_.is_DC17:
                d_3377___mcc_h0_ = source48_.cf25
                d_3378___mcc_h1_ = source48_.cf26
                d_3379___mcc_h2_ = source48_.cf27
                d_3380_cf27_ = d_3379___mcc_h2_
                d_3381_cf26_ = d_3378___mcc_h1_
                d_3382_cf25_ = d_3377___mcc_h0_
                index535_ = default__.safeIndex(206, (d_3361_v2_).length(0))
                (d_3361_v2_)[index535_] = d_3381_cf26_
                d_3383_v15_: _dafny.Array
                nw538_ = _dafny.Array(False, 9)
                d_3383_v15_ = nw538_
                index536_ = default__.safeIndex(178, (d_3383_v15_).length(0))
                (d_3383_v15_)[index536_] = (d_3382_cf25_) > (len(d_3368_v8_))
                index537_ = default__.safeIndex(178, (d_3383_v15_).length(0))
                (d_3383_v15_)[index537_] = (d_3380_cf27_) or ((self).f27)
                (globalState).f2 = (d_3382_cf25_ if True else d_3381_cf26_)
                d_3382_cf25_ = d_3382_cf25_
            elif source48_.is_DC15:
                d_3384___mcc_h3_ = source48_.cf24
                d_3385_cf24_ = d_3384___mcc_h3_
                index538_ = default__.safeIndex(206, (d_3361_v2_).length(0))
                (d_3361_v2_)[index538_] = (0) - ((d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))])
                d_3386_v16_: _dafny.MultiSet
                d_3386_v16_ = _dafny.MultiSet([d_3360_v1_])
                d_3387_v17_: _dafny.Map
                d_3387_v17_ = _dafny.Map({(0) - (default__.safeDivisionInt((d_3386_v16_).cardinality, (self).f33)): d_3368_v8_})
                d_3387_v17_ = (d_3387_v17_).set(default__.safeDivisionInt((self).f33, (d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))]), d_3368_v8_)
                d_3388_v18_: _dafny.Array
                nw539_ = _dafny.Array(None, 14)
                nw539_[int(0)] = (self).f27
                nw539_[int(1)] = (self).f27
                nw539_[int(2)] = (self).f27
                nw539_[int(3)] = (len(d_3360_v1_)) < ((self).f33)
                nw539_[int(4)] = (self).f27
                nw539_[int(5)] = (self).f27
                nw539_[int(6)] = (self).f27
                nw539_[int(7)] = ((self).f32) and ((self).f27)
                nw539_[int(8)] = (self).f27
                nw539_[int(9)] = (self).f27
                nw539_[int(10)] = (self).f27
                nw539_[int(11)] = (False if (self).f27 else (self).f32)
                nw539_[int(12)] = (d_3365_v5_)[default__.safeIndex((d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))], len(d_3365_v5_))]
                nw539_[int(13)] = (self).f27
                d_3388_v18_ = nw539_
                index539_ = default__.safeIndex(678, (d_3388_v18_).length(0))
                (d_3388_v18_)[index539_] = (self).f27
                index540_ = default__.safeIndex(678, (d_3388_v18_).length(0))
                (d_3388_v18_)[index540_] = (self).f27
                d_3368_v8_ = d_3368_v8_
            elif True:
                d_3389___mcc_h4_ = source48_.cf28
                d_3390_cf28_ = d_3389___mcc_h4_
                d_3391_v19_: _dafny.Array
                nw540_ = _dafny.Array(False, 13)
                d_3391_v19_ = nw540_
                d_3391_v19_ = d_3391_v19_
                d_3392_v20_: _dafny.MultiSet
                d_3392_v20_ = _dafny.MultiSet([(self).f33, 381, (d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))]])
                d_3393_v21_: _dafny.Map
                d_3393_v21_ = _dafny.Map({(d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))]: d_3392_v20_})
                d_3392_v20_ = ((d_3392_v20_).intersection(((d_3393_v21_)[(self).f33] if ((self).f33) in (d_3393_v21_) else d_3392_v20_))).set(len(d_3368_v8_), default__.abs(((d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))]) + ((d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))])))
                d_3394_v22_: _dafny.Map
                d_3394_v22_ = _dafny.Map({_dafny.Map({(d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))]: (d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))]}): False})
                d_3394_v22_ = (d_3394_v22_).set(_dafny.Map({(self).f33: len(_dafny.Set({(d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))]}))}), (self).f32)
                (globalState).f14 = (self).f27
            r0 = (d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))]
        elif True:
            (globalState).f6 = (d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))]
            (globalState).f2 = (d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))]
            if (self).f32:
                d_3395_v23_: _dafny.Array
                nw541_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 1)
                d_3395_v23_ = nw541_
                d_3395_v23_ = d_3395_v23_
                d_3396_v24_: _dafny.Array
                nw542_ = _dafny.Array(_dafny.Array(None, 0), 24)
                d_3396_v24_ = nw542_
                index541_ = default__.safeIndex(283, (d_3396_v24_).length(0))
                (d_3396_v24_)[index541_] = d_3361_v2_
                index542_ = default__.safeIndex(283, (d_3396_v24_).length(0))
                (d_3396_v24_)[index542_] = d_3361_v2_
                d_3397_v25_: _dafny.Array
                nw543_ = _dafny.Array(False, 29)
                d_3397_v25_ = nw543_
                d_3398_v26_: _dafny.Map
                d_3398_v26_ = _dafny.Map({(self).f27: (self).f32})
                d_3399_v27_: _dafny.Seq
                d_3399_v27_ = _dafny.SeqWithoutIsStrInference([(self).f27])
                d_3400_v28_: _dafny.Seq
                d_3400_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yuuk"))
                index543_ = default__.safeIndex(606, (d_3397_v25_).length(0))
                (d_3397_v25_)[index543_] = ((d_3398_v26_)[(d_3399_v27_)[default__.safeIndex((self).f33, len(d_3399_v27_))]] if ((d_3399_v27_)[default__.safeIndex((self).f33, len(d_3399_v27_))]) in (d_3398_v26_) else default__.fm0(d_3400_v28_, (self).f32, (self).f27, d_3360_v1_, globalState))
                d_3401_v29_: str
                d_3401_v29_ = _dafny.CodePoint('k')
                d_3402_v30_: D22
                d_3402_v30_ = D22_DC57((self).f32, d_3401_v29_, (self).f32)
                pat_let_tv72_ = d_3401_v29_
                index544_ = default__.safeIndex(606, (d_3397_v25_).length(0))
                def iife317_(_pat_let96_0):
                    def iife318_(d_3403_dt__update__tmp_h0_):
                        def iife319_(_pat_let97_0):
                            def iife320_(d_3404_dt__update_hcf85_h0_):
                                return D22_DC57((d_3403_dt__update__tmp_h0_).cf84, d_3404_dt__update_hcf85_h0_, (d_3403_dt__update__tmp_h0_).cf86)
                            return iife320_(_pat_let97_0)
                        return iife319_(pat_let_tv72_)
                    return iife318_(_pat_let96_0)
                (d_3397_v25_)[index544_] = (iife317_(d_3402_v30_)).cf84
                (globalState).f2 = ((d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))]) - ((self).f33)
                d_3405_v31_: _dafny.Array
                def lambda139_(d_3406_v26_):
                    def lambda140_(d_3407_i2_):
                        return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(d_3406_v26_) for d_3408_i3_ in range(default__.abs(784))]))])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f33, (self).f33])]))

                    return lambda140_

                init82_ = lambda139_(d_3398_v26_)
                nw544_ = _dafny.Array(None, 18)
                for i0_82_ in range(nw544_.length(0)):
                    nw544_[i0_82_] = init82_(i0_82_)
                d_3405_v31_ = nw544_
                d_3409_v32_: _dafny.Seq
                d_3409_v32_ = _dafny.SeqWithoutIsStrInference([(d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))], (self).f33, 419])
                d_3410_v33_: _dafny.Seq
                d_3410_v33_ = _dafny.SeqWithoutIsStrInference([d_3409_v32_])
                index545_ = default__.safeIndex(760, (d_3405_v31_).length(0))
                (d_3405_v31_)[index545_] = d_3410_v33_
                d_3411_v34_: _dafny.Seq
                d_3411_v34_ = _dafny.SeqWithoutIsStrInference([d_3400_v28_])
                index546_ = default__.safeIndex(206, (d_3361_v2_).length(0))
                index547_ = default__.safeIndex(760, (d_3405_v31_).length(0))
                rhs535_ = (0) - (-195)
                rhs536_ = default__.fm67((d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))], globalState)
                rhs537_ = (default__.safeDivisionInt((d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))], len((d_3411_v34_)[default__.safeIndex((0) - ((d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))]), len(d_3411_v34_))]))) * ((self).f33)
                lhs457_ = d_3361_v2_
                lhs458_ = default__.safeIndex(206, (d_3361_v2_).length(0))
                lhs459_ = d_3405_v31_
                lhs460_ = default__.safeIndex(760, (d_3405_v31_).length(0))
                lhs461_ = globalState
                lhs457_[lhs458_] = rhs535_
                lhs459_[lhs460_] = rhs536_
                lhs461_.f6 = rhs537_
            elif True:
                d_3412_v35_: _dafny.Seq
                d_3412_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wmgvagwiv"))
                d_3413_v36_: str
                d_3413_v36_ = _dafny.CodePoint('k')
                d_3414_v37_: _dafny.Map
                d_3414_v37_ = _dafny.Map({d_3412_v35_: d_3413_v36_})
                d_3415_v38_: _dafny.MultiSet
                d_3415_v38_ = _dafny.MultiSet([len(d_3414_v37_)])
                d_3416_v39_: _dafny.Seq
                d_3416_v39_ = _dafny.SeqWithoutIsStrInference([(self).f32])
                index548_ = default__.safeIndex(206, (d_3361_v2_).length(0))
                (d_3361_v2_)[index548_] = (0) - (default__.fm1(((d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))]) not in (d_3415_v38_), d_3416_v39_, (self).f32, globalState))
                d_3417_v40_: _dafny.MultiSet
                d_3417_v40_ = _dafny.MultiSet([d_3413_v36_])
                d_3418_v41_: _dafny.Set
                d_3418_v41_ = _dafny.Set({d_3361_v2_})
                d_3419_v42_: _dafny.Seq
                d_3419_v42_ = _dafny.SeqWithoutIsStrInference([default__.fm73(len(d_3418_v41_), (_dafny.MultiSet([len(d_3412_v35_)])).cardinality, 4, globalState)])
                d_3420_v43_: _dafny.Seq
                d_3420_v43_ = _dafny.SeqWithoutIsStrInference([default__.fm1((self).f27, (d_3416_v39_).set(default__.safeIndex((self).f33, len(d_3416_v39_)), (self).f27), (self).f27, globalState)])
                d_3421_v44_: _dafny.Map
                d_3421_v44_ = _dafny.Map({(d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))]: d_3417_v40_})
                d_3422_v45_: T0
                nw545_ = C3()
                nw545_.ctor__((self).f27, False)
                d_3422_v45_ = nw545_
                d_3423_v46_: D18
                d_3423_v46_ = D18_DC41(_dafny.Set({(d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))], 597}))
                d_3424_v47_: D20
                d_3424_v47_ = D20_DC49(default__.fm29((self).f33, (self).f27, globalState), len(d_3360_v1_), (d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))], d_3422_v45_, (d_3423_v46_).cf56)
                d_3425_v48_: _dafny.Array
                nw546_ = _dafny.Array(None, 22)
                nw546_[int(0)] = d_3417_v40_
                nw546_[int(1)] = ((d_3419_v42_)[default__.safeIndex((self).f33, len(d_3419_v42_))]) - (d_3417_v40_)
                nw546_[int(2)] = _dafny.MultiSet([d_3413_v36_, d_3413_v36_, d_3413_v36_, d_3413_v36_, _dafny.CodePoint('y')])
                nw546_[int(3)] = d_3417_v40_
                nw546_[int(4)] = _dafny.MultiSet(d_3412_v35_)
                nw546_[int(5)] = d_3417_v40_
                nw546_[int(6)] = (d_3417_v40_).set(d_3413_v36_, default__.abs((d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))]))
                nw546_[int(7)] = d_3417_v40_
                nw546_[int(8)] = (d_3417_v40_) | (d_3417_v40_)
                nw546_[int(9)] = ((d_3417_v40_).set(d_3413_v36_, default__.abs(len(d_3420_v43_))) if (self).f32 else _dafny.MultiSet([(d_3412_v35_)[default__.safeIndex((d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))], len(d_3412_v35_))], d_3413_v36_]))
                nw546_[int(10)] = (d_3417_v40_).intersection(((d_3421_v44_)[(d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))]] if ((d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))]) in (d_3421_v44_) else default__.fm73(246, (self).f33, (self).f33, globalState)))
                nw546_[int(11)] = (d_3417_v40_) - (_dafny.MultiSet((self).fm3(globalState)))
                nw546_[int(12)] = d_3417_v40_
                nw546_[int(13)] = d_3417_v40_
                nw546_[int(14)] = d_3417_v40_
                nw546_[int(15)] = _dafny.MultiSet([d_3413_v36_, d_3413_v36_, d_3413_v36_])
                nw546_[int(16)] = _dafny.MultiSet([d_3413_v36_])
                nw546_[int(17)] = d_3417_v40_
                nw546_[int(18)] = (_dafny.MultiSet([d_3413_v36_])) - (d_3417_v40_)
                nw546_[int(19)] = d_3417_v40_
                nw546_[int(20)] = _dafny.MultiSet([(d_3424_v47_).cf69, (d_3412_v35_)[default__.safeIndex((self).f33, len(d_3412_v35_))]])
                nw546_[int(21)] = d_3417_v40_
                d_3425_v48_ = nw546_
                index549_ = default__.safeIndex(6, (d_3425_v48_).length(0))
                (d_3425_v48_)[index549_] = d_3417_v40_
                index550_ = default__.safeIndex(6, (d_3425_v48_).length(0))
                (d_3425_v48_)[index550_] = _dafny.MultiSet([(d_3413_v36_ if (d_3422_v45_).f27 else d_3413_v36_)])
                d_3426_v49_: _dafny.Seq
                d_3426_v49_ = _dafny.SeqWithoutIsStrInference([d_3361_v2_])
                d_3427_v50_: _dafny.Set
                d_3427_v50_ = _dafny.Set({d_3420_v43_, d_3420_v43_})
                d_3428_v51_: D27
                d_3428_v51_ = D27_DC67(d_3426_v49_, d_3427_v50_, (d_3422_v45_).f27)
                d_3429_v52_: C2
                nw547_ = C2()
                nw547_.ctor__((D27_DC67(_dafny.SeqWithoutIsStrInference([d_3361_v2_, d_3361_v2_, d_3361_v2_]), (d_3428_v51_).cf106, True)).cf107, not((self).f32))
                d_3429_v52_ = nw547_
                d_3430_v53_: _dafny.Seq
                d_3430_v53_ = _dafny.SeqWithoutIsStrInference([d_3429_v52_, d_3429_v52_])
                d_3431_v54_: _dafny.MultiSet
                d_3431_v54_ = _dafny.MultiSet([False])
                d_3432_v55_: _dafny.Array
                nw548_ = _dafny.Array(None, 21)
                nw548_[int(0)] = d_3429_v52_
                nw548_[int(1)] = d_3429_v52_
                nw548_[int(2)] = d_3429_v52_
                nw548_[int(3)] = d_3429_v52_
                nw548_[int(4)] = d_3429_v52_
                nw548_[int(5)] = d_3429_v52_
                nw548_[int(6)] = d_3429_v52_
                nw548_[int(7)] = d_3429_v52_
                nw548_[int(8)] = (d_3430_v53_)[default__.safeIndex((d_3431_v54_).cardinality, len(d_3430_v53_))]
                nw548_[int(9)] = d_3429_v52_
                nw548_[int(10)] = d_3429_v52_
                nw548_[int(11)] = d_3429_v52_
                nw548_[int(12)] = d_3429_v52_
                nw548_[int(13)] = d_3429_v52_
                nw548_[int(14)] = d_3429_v52_
                nw548_[int(15)] = d_3429_v52_
                nw548_[int(16)] = d_3429_v52_
                nw548_[int(17)] = d_3429_v52_
                nw548_[int(18)] = (D34_DC82((d_3430_v53_)[default__.safeIndex((0) - ((self).f33), len(d_3430_v53_))])).cf135
                nw548_[int(19)] = d_3429_v52_
                nw548_[int(20)] = d_3429_v52_
                d_3432_v55_ = nw548_
                index551_ = default__.safeIndex(111, (d_3432_v55_).length(0))
                (d_3432_v55_)[index551_] = d_3429_v52_
                index552_ = default__.safeIndex(111, (d_3432_v55_).length(0))
                (d_3432_v55_)[index552_] = (d_3429_v52_ if not ((self).f27) or ((self).f27) else d_3429_v52_)
                d_3433_v56_: _dafny.Array
                def lambda141_(d_3434_v45_):
                    def lambda142_(d_3435_i4_):
                        return (d_3434_v45_).f27

                    return lambda142_

                init83_ = lambda141_(d_3422_v45_)
                nw549_ = _dafny.Array(None, 2)
                for i0_83_ in range(nw549_.length(0)):
                    nw549_[i0_83_] = init83_(i0_83_)
                d_3433_v56_ = nw549_
                index553_ = default__.safeIndex(569, (d_3433_v56_).length(0))
                (d_3433_v56_)[index553_] = (d_3422_v45_).f27
                d_3436_v57_: _dafny.Map
                d_3436_v57_ = _dafny.Map({d_3361_v2_: len(d_3416_v39_)})
                index554_ = default__.safeIndex(569, (d_3433_v56_).length(0))
                (d_3433_v56_)[index554_] = (d_3361_v2_) in (d_3436_v57_)
                (globalState).f6 = (d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))]
            d_3437_v58_: _dafny.Map
            d_3437_v58_ = _dafny.Map({(d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))]: not((self).f32)})
            d_3438_v59_: _dafny.Seq
            d_3438_v59_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dsq"))
            d_3439_v60_: C1
            nw550_ = C1()
            nw550_.ctor__(False)
            d_3439_v60_ = nw550_
            d_3440_v61_: _dafny.Map
            d_3440_v61_ = _dafny.Map({len(d_3438_v59_): d_3439_v60_})
            d_3441_v62_: str
            d_3441_v62_ = _dafny.CodePoint('j')
            d_3437_v58_ = (d_3437_v58_).set(len(d_3440_v61_), (D25_DC64((self).f32, (d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))], d_3441_v62_, (self).f27, (d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))])).cf98)
            d_3442_v63_: _dafny.Seq
            d_3442_v63_ = _dafny.SeqWithoutIsStrInference([d_3360_v1_])
            d_3443_v64_: _dafny.Map
            d_3443_v64_ = _dafny.Map({(self).f27: d_3360_v1_})
            d_3444_v65_: D13
            d_3444_v65_ = D13_DC28((d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))])
            d_3445_v66_: _dafny.Set
            d_3445_v66_ = _dafny.Set({(d_3442_v63_)[default__.safeIndex((self).f33, len(d_3442_v63_))], ((d_3443_v64_)[(self).f32] if ((self).f32) in (d_3443_v64_) else d_3360_v1_), default__.fm35(d_3444_v65_, (self).f32, (self).f27, globalState)})
            d_3445_v66_ = d_3445_v66_
        d_3446_v67_: _dafny.Map
        d_3446_v67_ = _dafny.Map({False: (self).f32})
        hi23_ = ((0) - (len(_dafny.Map({d_3446_v67_: (self).f27}))) if (self).f27 else (self).f33)
        for d_3447_i5_ in range((d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))], hi23_):
            (globalState).f14 = (self).f27
            d_3448_v68_: _dafny.Seq
            d_3448_v68_ = _dafny.SeqWithoutIsStrInference([(self).f32])
            d_3449_v69_: _dafny.Seq
            d_3449_v69_ = _dafny.SeqWithoutIsStrInference([d_3447_i5_])
            d_3448_v68_ = ((default__.fm32((self).f27, d_3449_v69_, (self).f32, globalState)) + (d_3448_v68_)) + (d_3448_v68_)
            (globalState).f2 = 646
            r1 = (((self).f33) + ((d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))])) + ((self).f33)
        d_3450_v70_: _dafny.Seq
        d_3450_v70_ = _dafny.SeqWithoutIsStrInference([(self).f27])
        d_3451_v71_: _dafny.MultiSet
        d_3451_v71_ = _dafny.MultiSet([d_3450_v70_, d_3450_v70_])
        d_3452_v72_: D0
        d_3452_v72_ = D0_DC2(((d_3451_v71_)[d_3450_v70_] if (d_3450_v70_) in (d_3451_v71_) else (d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))]), (d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))])
        d_3452_v72_ = D0_DC2((d_3361_v2_)[default__.safeIndex(206, (d_3361_v2_).length(0))], (self).f33)
        r0 = 616
        r1 = (self).f33
        return r0, r1

    @property
    def f32(self):
        return self._f32
    @property
    def f33(self):
        return self._f33

class C9(T1, T2, T0):
    def  __init__(self):
        self._f27: bool = False
        self._f28: int = int(0)
        self._f29: int = int(0)
        self._f30: _dafny.Array = _dafny.Array(None, 0)
        self._f31: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C9"
    @property
    def f27(self):
        return self._f27
    @property
    def f28(self):
        return self._f28
    @property
    def f29(self):
        return self._f29
    def ctor__(self, f30, f31, f28, f29, f27):
        (self)._f30 = f30
        (self)._f31 = f31
        (self)._f28 = f28
        (self)._f29 = f29
        (self)._f27 = f27

    def fm2(self, p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mvuomxtd"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xaie")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tv")))

    def fm3(self, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hytyaypu"))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x")) if (self).f27 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mkueo"))))

    def fm4(self, p0, globalState):
        return (self).f27

    def fm5(self, globalState):
        source49_ = D0_DC1((self).f27)
        if source49_.is_DC1:
            d_3453___mcc_h0_ = source49_.cf1
            d_3454_cf1_ = d_3453___mcc_h0_
            return D0_DC2((self).f31, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([(self).f27])): (_dafny.MultiSet([(self).f28, 604])).cardinality})))
        elif source49_.is_DC2:
            d_3455___mcc_h1_ = source49_.cf2
            d_3456___mcc_h2_ = source49_.cf3
            d_3457_cf3_ = d_3456___mcc_h2_
            d_3458_cf2_ = d_3455___mcc_h1_
            return D0_DC2((self).f31, (self).f28)
        elif source49_.is_DC0:
            d_3459___mcc_h3_ = source49_.cf0
            d_3460_cf0_ = d_3459___mcc_h3_
            return D0_DC2(d_3460_cf0_, (self).f28)
        elif True:
            d_3461___mcc_h4_ = source49_.cf4
            d_3462_cf4_ = d_3461___mcc_h4_
            return D0_DC2(len(_dafny.Map({(self).f27: (self).f27})), len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_3463_i0_ in range(default__.abs(-101))])), 704, (self).f31, (self).f28})))

    def fm6(self, p0, p1, globalState):
        if (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_3464_i0_ in range(default__.abs(-282))])) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uhwasda"))):
            return _dafny.CodePoint('b')
        elif True:
            return _dafny.CodePoint('o')

    def fm7(self, globalState):
        return (self).f27

    def m3(self, p0, p1, p2, globalState):
        d_3465_v0_: _dafny.Seq
        d_3465_v0_ = _dafny.SeqWithoutIsStrInference([p0, p0])
        d_3466_v1_: _dafny.Seq
        d_3466_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lffl"))
        d_3467_v2_: str
        d_3467_v2_ = _dafny.CodePoint('i')
        d_3468_v3_: C3
        nw551_ = C3()
        nw551_.ctor__((d_3465_v0_)[default__.safeIndex(len((d_3466_v1_).set(default__.safeIndex(-795, len(d_3466_v1_)), d_3467_v2_)), len(d_3465_v0_))], p0)
        d_3468_v3_ = nw551_
        if p0:
            d_3469_v4_: C4
            nw552_ = C4()
            nw552_.ctor__(_dafny.MultiSet([True, (self).f27, False, d_3468_v3_.f38]), not((d_3468_v3_.f38) == ((self).f27)))
            d_3469_v4_ = nw552_
            d_3469_v4_ = d_3469_v4_
            d_3470_v5_: T1
            nw553_ = C6()
            nw553_.ctor__(((D21_DC52((d_3469_v4_).f37)).cf75).cardinality, 253, False)
            d_3470_v5_ = nw553_
            d_3471_v6_: _dafny.Map
            d_3471_v6_ = _dafny.Map({(self).f27: d_3470_v5_})
            d_3471_v6_ = (d_3471_v6_).set(d_3468_v3_.f38, d_3470_v5_)
            (globalState).f5 = (d_3466_v1_) != ((d_3466_v1_ if (self).f27 else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_3472_i0_ in range(default__.abs(471))])))
            d_3473_v7_: _dafny.MultiSet
            d_3473_v7_ = _dafny.MultiSet([d_3467_v2_, d_3467_v2_, d_3467_v2_, d_3467_v2_])
            d_3473_v7_ = d_3473_v7_
            (globalState).f14 = ((self).f27 if not ((self).f27) or ((self).f27) else d_3468_v3_.f38)
        elif True:
            d_3474_v8_: _dafny.Array
            nw554_ = _dafny.Array(_dafny.Map({}), 2)
            d_3474_v8_ = nw554_
            d_3475_v9_: _dafny.Array
            nw555_ = _dafny.Array(None, 1)
            nw555_[int(0)] = d_3474_v8_
            d_3475_v9_ = nw555_
            index555_ = default__.safeIndex(76, (d_3475_v9_).length(0))
            (d_3475_v9_)[index555_] = d_3474_v8_
            index556_ = default__.safeIndex(76, (d_3475_v9_).length(0))
            nw556_ = _dafny.Array(_dafny.Map({}), 23)
            (d_3475_v9_)[index556_] = nw556_
            d_3476_v10_: C0
            nw557_ = C0()
            nw557_.ctor__()
            d_3476_v10_ = nw557_
            d_3476_v10_ = d_3476_v10_
            index557_ = default__.safeIndex(556, ((self).f30).length(0))
            ((self).f30)[index557_] = (self).f28
            d_3477_v11_: _dafny.MultiSet
            d_3477_v11_ = _dafny.MultiSet([not(False), p0])
            index558_ = default__.safeIndex(556, ((self).f30).length(0))
            ((self).f30)[index558_] = ((d_3477_v11_).cardinality) * (-280)
            d_3478_v12_: _dafny.Set
            d_3478_v12_ = _dafny.Set({((self).f30)[default__.safeIndex(556, ((self).f30).length(0))]})
            d_3479_v13_: _dafny.Map
            d_3479_v13_ = _dafny.Map({default__.fm0(d_3466_v1_, p0, d_3468_v3_.f38, d_3478_v12_, globalState): d_3468_v3_.f38})
            d_3479_v13_ = (d_3479_v13_).set(not (True) or (p0), d_3468_v3_.f38)
            d_3480_v14_: C7
            nw558_ = C7()
            nw558_.ctor__(p0, (self).f29, (0) - (p2), (0) - ((self).f29), not((self).f27))
            d_3480_v14_ = nw558_
            d_3481_v15_: C7
            d_3481_v15_ = d_3480_v14_
            d_3482_v16_: _dafny.Set
            d_3482_v16_ = _dafny.Set({d_3480_v14_, (d_3481_v15_)})
            d_3483_v17_: _dafny.Map
            d_3483_v17_ = _dafny.Map({(self).f27: d_3482_v16_})
            d_3482_v16_ = ((d_3482_v16_).intersection(d_3482_v16_)) | ((((d_3483_v17_)[False] if (False) in (d_3483_v17_) else d_3482_v16_)) - (d_3482_v16_))
        d_3484_v18_: D17
        d_3484_v18_ = D17_DC39(d_3466_v1_)
        pat_let_tv73_ = d_3466_v1_
        def lambda143_(source50_):
            if source50_.is_DC39:
                d_3485___mcc_h0_ = source50_.cf54
                d_3486_cf54_ = d_3485___mcc_h0_
                return (0) - ((self).f29)
            elif source50_.is_DC38:
                d_3487___mcc_h1_ = source50_.cf53
                d_3488_cf53_ = d_3487___mcc_h1_
                return (_dafny.MultiSet([len(pat_let_tv73_)])).cardinality
            elif True:
                d_3489___mcc_h2_ = source50_.cf55
                d_3490_cf55_ = d_3489___mcc_h2_
                return (self).f29

        (globalState).f2 = lambda143_(d_3484_v18_)
        d_3491_v19_: _dafny.MultiSet
        d_3491_v19_ = _dafny.MultiSet([True, (self).f27])
        d_3492_v20_: _dafny.Map
        d_3492_v20_ = _dafny.Map({(d_3491_v19_).cardinality: (self).f28})
        d_3493_v21_: _dafny.MultiSet
        d_3493_v21_ = _dafny.MultiSet([(self).f31, (0) - ((self).f31), 623, len(d_3492_v20_), (0) - ((self).f31)])
        d_3494_v22_: C2
        nw559_ = C2()
        nw559_.ctor__((_dafny.MultiSet([p2])).ispropersubset(d_3493_v21_), p0)
        d_3494_v22_ = nw559_
        d_3495_v24_: _dafny.Array
        def lambda144_(d_3496_v0_, d_3497_p2_, d_3498_v22_, d_3499_v3_):
            def lambda145_(d_3500_i2_):
                def iife321_():
                    coll125_ = _dafny.Map()
                    compr_125_: int
                    for compr_125_ in _dafny.IntegerRange(879, 784):
                        d_3501_v23_: int = compr_125_
                        if ((879) <= (d_3501_v23_)) and ((d_3501_v23_) < (784)):
                            coll125_[default__.safeModuloInt(d_3501_v23_, d_3497_p2_)] = (D5_DC12((d_3498_v22_).f34, (self).f29)).cf19
                    return _dafny.Map(coll125_)
                return (iife321_()
                ) != (_dafny.Map({960: (d_3496_v0_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([D30_DC74(True, (self).f29, d_3499_v3_.f38)])), len(d_3496_v0_))]}))

            return lambda145_

        init84_ = lambda144_(d_3465_v0_, p2, d_3494_v22_, d_3468_v3_)
        nw560_ = _dafny.Array(None, 4)
        for i0_84_ in range(nw560_.length(0)):
            nw560_[i0_84_] = init84_(i0_84_)
        d_3495_v24_ = nw560_
        guard_loop_11_: int
        for guard_loop_11_ in _dafny.IntegerRange(0, (d_3495_v24_).length(0)):
            d_3502_i1_: int = guard_loop_11_
            if (True) and (((0) <= (d_3502_i1_)) and ((d_3502_i1_) < ((d_3495_v24_).length(0)))):
                (d_3495_v24_)[(d_3502_i1_)] = False
        d_3503_v25_: _dafny.Map
        d_3503_v25_ = _dafny.Map({default__.safeModuloInt((self).f31, (self).f28): d_3495_v24_})
        d_3504_v26_: _dafny.Map
        d_3504_v26_ = _dafny.Map({p0: _dafny.Map({(self).f27: d_3468_v3_.f38})})
        d_3505_v27_: D25
        d_3505_v27_ = D25_DC64((self).f27, (self).f28, _dafny.CodePoint('i'), (d_3494_v22_).f34, -811)
        d_3503_v25_ = (d_3503_v25_).set((len(d_3504_v26_)) + ((d_3505_v27_).cf99), d_3495_v24_)

    def m1(self, globalState):
        r0: int = int(0)
        d_3506_v0_: _dafny.Array
        nw561_ = _dafny.Array(False, 20)
        d_3506_v0_ = nw561_
        index559_ = default__.safeIndex(504, (d_3506_v0_).length(0))
        (d_3506_v0_)[index559_] = (self).f27
        d_3507_v1_: D0
        d_3507_v1_ = D0_DC1((self).f27)
        d_3508_v2_: _dafny.Seq
        d_3508_v2_ = _dafny.SeqWithoutIsStrInference([(self).f27, not((self).fm4(d_3507_v1_, globalState)), ((self).f31) != (234), (self).f27, (self).f27])
        index560_ = default__.safeIndex(504, (d_3506_v0_).length(0))
        (d_3506_v0_)[index560_] = (d_3508_v2_)[default__.safeIndex((self).f31, len(d_3508_v2_))]
        d_3509_v3_: C3
        nw562_ = C3()
        nw562_.ctor__((d_3506_v0_)[default__.safeIndex(504, (d_3506_v0_).length(0))], (d_3506_v0_)[default__.safeIndex(504, (d_3506_v0_).length(0))])
        d_3509_v3_ = nw562_
        index561_ = default__.safeIndex(955, ((self).f30).length(0))
        ((self).f30)[index561_] = 597
        d_3510_v4_: _dafny.Seq
        d_3510_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dc"))
        index562_ = default__.safeIndex(504, (d_3506_v0_).length(0))
        index563_ = default__.safeIndex(955, ((self).f30).length(0))
        rhs538_ = 670
        rhs539_ = (len(d_3510_v4_)) != ((self).f29)
        rhs540_ = _dafny.SeqWithoutIsStrInference([(self).f27])
        rhs541_ = (self).f29
        lhs462_ = globalState
        lhs463_ = d_3506_v0_
        lhs464_ = default__.safeIndex(504, (d_3506_v0_).length(0))
        lhs465_ = (self).f30
        lhs466_ = default__.safeIndex(955, ((self).f30).length(0))
        lhs462_.f6 = rhs538_
        lhs463_[lhs464_] = rhs539_
        d_3508_v2_ = rhs540_
        lhs465_[lhs466_] = rhs541_
        (d_3509_v3_).f38 = (self).f27
        d_3511_v5_: str
        d_3511_v5_ = _dafny.CodePoint('s')
        pat_let_tv74_ = d_3510_v4_
        pat_let_tv75_ = d_3509_v3_
        def lambda146_(source51_):
            if source51_.is_DC64:
                d_3512___mcc_h0_ = source51_.cf98
                d_3513___mcc_h1_ = source51_.cf99
                d_3514___mcc_h2_ = source51_.cf100
                d_3515___mcc_h3_ = source51_.cf101
                d_3516___mcc_h4_ = source51_.cf102
                d_3517_cf102_ = d_3516___mcc_h4_
                d_3518_cf101_ = d_3515___mcc_h3_
                d_3519_cf100_ = d_3514___mcc_h2_
                d_3520_cf99_ = d_3513___mcc_h1_
                d_3521_cf98_ = d_3512___mcc_h0_
                return (pat_let_tv74_) < (_dafny.SeqWithoutIsStrInference([d_3519_cf100_ for d_3522_i0_ in range(default__.abs(36))]))
            elif True:
                d_3523___mcc_h5_ = source51_.cf97
                d_3524_cf97_ = d_3523___mcc_h5_
                return pat_let_tv75_.f38

        if lambda146_(D25_DC64((d_3506_v0_)[default__.safeIndex(504, (d_3506_v0_).length(0))], (self).f31, d_3511_v5_, False, 423)):
            d_3525_v6_: _dafny.Set
            d_3525_v6_ = _dafny.Set({(self).f29, (self).f29})
            index564_ = default__.safeIndex(504, (d_3506_v0_).length(0))
            (d_3506_v0_)[index564_] = (not((d_3525_v6_).ispropersubset(d_3525_v6_))) and (d_3509_v3_.f38)
            d_3526_v7_: _dafny.MultiSet
            d_3526_v7_ = _dafny.MultiSet([not((d_3506_v0_)[default__.safeIndex(504, (d_3506_v0_).length(0))]), False, d_3509_v3_.f38, (self).fm4(d_3507_v1_, globalState)])
            d_3527_v8_: _dafny.Map
            d_3527_v8_ = _dafny.Map({(self).f31: (self).f27})
            d_3528_v9_: C4
            nw563_ = C4()
            nw563_.ctor__(d_3526_v7_, ((d_3527_v8_)[(self).f31] if ((self).f31) in (d_3527_v8_) else (d_3506_v0_)[default__.safeIndex(504, (d_3506_v0_).length(0))]))
            d_3528_v9_ = nw563_
            d_3529_v10_: _dafny.Map
            d_3529_v10_ = _dafny.Map({((self).f30)[default__.safeIndex(955, ((self).f30).length(0))]: d_3510_v4_})
            d_3530_v11_: T0
            nw564_ = C4()
            nw564_.ctor__(((d_3528_v9_).f37).set((d_3506_v0_)[default__.safeIndex(504, (d_3506_v0_).length(0))], default__.abs(len(d_3529_v10_))), d_3509_v3_.f38)
            d_3530_v11_ = nw564_
            d_3531_v12_: _dafny.Map
            d_3531_v12_ = _dafny.Map({d_3530_v11_: (d_3526_v7_).cardinality})
            d_3532_v13_: _dafny.Seq
            d_3532_v13_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({(self).f29: (self).f28}))])
            d_3533_v14_: _dafny.Array
            nw565_ = _dafny.Array(None, 7)
            nw565_[int(0)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ar")))
            nw565_[int(1)] = (self).f29
            nw565_[int(2)] = ((self).f30)[default__.safeIndex(955, ((self).f30).length(0))]
            nw565_[int(3)] = (488) * ((self).f29)
            nw565_[int(4)] = ((self).f30)[default__.safeIndex(955, ((self).f30).length(0))]
            nw565_[int(5)] = ((d_3531_v12_)[d_3530_v11_] if (d_3530_v11_) in (d_3531_v12_) else (d_3532_v13_)[default__.safeIndex((self).f28, len(d_3532_v13_))])
            nw565_[int(6)] = (d_3532_v13_)[default__.safeIndex((self).f29, len(d_3532_v13_))]
            d_3533_v14_ = nw565_
            d_3533_v14_ = d_3533_v14_
            d_3534_v15_: _dafny.Map
            d_3534_v15_ = _dafny.Map({len(d_3510_v4_): (self).f29})
            d_3535_v17_: D25
            d_3535_v17_ = D25_DC64((self).f27, (self).f28, d_3511_v5_, False, (self).f28)
            d_3536_v18_: _dafny.Set
            def iife322_():
                coll126_ = _dafny.Set()
                compr_126_: int
                for compr_126_ in (_dafny.SeqWithoutIsStrInference([(self).f31])).Elements:
                    d_3537_v16_: int = compr_126_
                    if (d_3537_v16_) in (_dafny.SeqWithoutIsStrInference([(self).f31])):
                        coll126_ = coll126_.union(_dafny.Set([(d_3537_v16_) + (len(_dafny.SeqWithoutIsStrInference([True, True, True, True, True])))]))
                return _dafny.Set(coll126_)
            d_3536_v18_ = _dafny.Set({D25_DC64((d_3506_v0_)[default__.safeIndex(504, (d_3506_v0_).length(0))], len(d_3534_v15_), d_3511_v5_, d_3509_v3_.f38, (self).f28), D25_DC64(d_3509_v3_.f38, (self).f29, d_3511_v5_, (self).f27, len(iife322_()
)), d_3535_v17_})
            d_3538_v19_: _dafny.Seq
            d_3538_v19_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({((self).f30)[default__.safeIndex(955, ((self).f30).length(0))]: d_3509_v3_.f38})])
            d_3539_v20_: _dafny.Map
            d_3539_v20_ = _dafny.Map({d_3536_v18_: (len(d_3538_v19_)) * ((self).f31)})
            d_3539_v20_ = (d_3539_v20_).set(d_3536_v18_, (0) - ((((self).f30)[default__.safeIndex(955, ((self).f30).length(0))]) - ((self).f29)))
            index565_ = default__.safeIndex(955, ((self).f30).length(0))
            ((self).f30)[index565_] = (0) - (len(_dafny.SeqWithoutIsStrInference([(d_3510_v4_)[default__.safeIndex((self).f29, len(d_3510_v4_))] for d_3540_i1_ in range(default__.abs(947))])))
        elif True:
            d_3541_v21_: C0
            nw566_ = C0()
            nw566_.ctor__()
            d_3541_v21_ = nw566_
            d_3542_v22_: _dafny.Map
            d_3542_v22_ = _dafny.Map({((self).f30)[default__.safeIndex(955, ((self).f30).length(0))]: d_3541_v21_})
            d_3543_v23_: _dafny.Map
            d_3543_v23_ = ((d_3542_v22_)) | (_dafny.Map({(self).f31: d_3541_v21_}))
            d_3543_v23_ = _dafny.Map({(self).f31: d_3541_v21_})
            d_3544_v24_: _dafny.Set
            d_3544_v24_ = _dafny.Set({(self).f29, (self).f29})
            (globalState).f14 = (3) >= ((214) * (len(d_3544_v24_)))
            d_3545_v25_: _dafny.Map
            d_3545_v25_ = _dafny.Map({(d_3509_v3_.f38 if d_3509_v3_.f38 else d_3509_v3_.f38): d_3509_v3_.f38})
            d_3545_v25_ = (d_3545_v25_).set(d_3509_v3_.f38, (((self).f30)[default__.safeIndex(955, ((self).f30).length(0))]) == ((self).f29))
            d_3506_v0_ = d_3506_v0_
            (globalState).f21 = d_3509_v3_.f38
        d_3546_v26_: _dafny.Map
        d_3546_v26_ = _dafny.Map({((self).f30)[default__.safeIndex(955, ((self).f30).length(0))]: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ddodpw"))})
        d_3546_v26_ = d_3546_v26_
        r0 = len(d_3510_v4_)
        return r0

    def m2(self, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: _dafny.Array = _dafny.Array(None, 0)
        r3: bool = False
        d_3547_v0_: _dafny.Seq
        d_3547_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hvq"))
        d_3548_v1_: _dafny.Map
        d_3548_v1_ = _dafny.Map({(self).f27: d_3547_v0_})
        d_3549_v2_: _dafny.Array
        nw567_ = _dafny.Array(None, 3)
        nw567_[int(0)] = (self).f29
        nw567_[int(1)] = len(d_3548_v1_)
        nw567_[int(2)] = (self).f29
        d_3549_v2_ = nw567_
        rhs542_ = d_3549_v2_
        rhs543_ = d_3549_v2_
        d_3549_v2_ = rhs542_
        d_3549_v2_ = rhs543_
        d_3550_v3_: C1
        nw568_ = C1()
        nw568_.ctor__((self).f27)
        d_3550_v3_ = nw568_
        d_3551_v4_: _dafny.Map
        d_3551_v4_ = _dafny.Map({(self).f28: d_3550_v3_})
        hi24_ = ((self).f29) + (len(d_3551_v4_))
        for d_3552_i0_ in range(((self).f29) + ((self).f28), hi24_):
            index566_ = default__.safeIndex(31, (d_3549_v2_).length(0))
            (d_3549_v2_)[index566_] = 366
            d_3553_v5_: _dafny.Map
            d_3553_v5_ = _dafny.Map({(self).f29: d_3547_v0_})
            d_3554_v6_: _dafny.Set
            d_3554_v6_ = _dafny.Set({(0) - ((self).f31), (self).f28})
            d_3555_v7_: _dafny.Seq
            d_3555_v7_ = _dafny.SeqWithoutIsStrInference([(self).f27])
            index567_ = default__.safeIndex(31, (d_3549_v2_).length(0))
            (d_3549_v2_)[index567_] = default__.fm1(default__.fm0(((d_3553_v5_)[(self).f28] if ((self).f28) in (d_3553_v5_) else d_3547_v0_), (self).fm7(globalState), (self).f27, d_3554_v6_, globalState), d_3555_v7_, False, globalState)
            index568_ = default__.safeIndex(31, (d_3549_v2_).length(0))
            (d_3549_v2_)[index568_] = (0) - ((0) - ((len((d_3547_v0_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "re")))) if (d_3555_v7_)[default__.safeIndex(d_3552_i0_, len(d_3555_v7_))] else (d_3549_v2_)[default__.safeIndex(31, (d_3549_v2_).length(0))])))
            d_3556_v8_: C7
            nw569_ = C7()
            nw569_.ctor__((self).f27, (d_3549_v2_)[default__.safeIndex(31, (d_3549_v2_).length(0))], (d_3549_v2_)[default__.safeIndex(31, (d_3549_v2_).length(0))], len(d_3547_v0_), False)
            d_3556_v8_ = nw569_
            d_3557_v9_: _dafny.Array
            nw570_ = _dafny.Array(None, 11)
            nw570_[int(0)] = d_3556_v8_
            nw570_[int(1)] = d_3556_v8_
            nw570_[int(2)] = d_3556_v8_
            nw570_[int(3)] = d_3556_v8_
            nw570_[int(4)] = d_3556_v8_
            nw570_[int(5)] = d_3556_v8_
            nw570_[int(6)] = d_3556_v8_
            nw570_[int(7)] = d_3556_v8_
            nw570_[int(8)] = (d_3556_v8_ if d_3556_v8_.f35 else d_3556_v8_)
            nw570_[int(9)] = d_3556_v8_
            nw570_[int(10)] = d_3556_v8_
            d_3557_v9_ = nw570_
            index569_ = default__.safeIndex(99, (d_3557_v9_).length(0))
            (d_3557_v9_)[index569_] = d_3556_v8_
            index570_ = default__.safeIndex(99, (d_3557_v9_).length(0))
            (d_3557_v9_)[index570_] = d_3556_v8_
            d_3558_v10_: _dafny.Array
            def lambda147_(d_3559_i1_):
                return D20_DC50()

            init85_ = lambda147_
            nw571_ = _dafny.Array(None, 20)
            for i0_85_ in range(nw571_.length(0)):
                nw571_[i0_85_] = init85_(i0_85_)
            d_3558_v10_ = nw571_
            d_3558_v10_ = d_3558_v10_
        (globalState).f2 = 846
        guard_loop_12_: int
        for guard_loop_12_ in _dafny.IntegerRange(0, (d_3549_v2_).length(0)):
            d_3560_i2_: int = guard_loop_12_
            if (True) and (((0) <= (d_3560_i2_)) and ((d_3560_i2_) < ((d_3549_v2_).length(0)))):
                (d_3549_v2_)[(d_3560_i2_)] = (d_3560_i2_) * (len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_3561_i3_ in range(default__.abs(-84))])) + (d_3547_v0_)))
        d_3548_v1_ = (d_3548_v1_).set((self).f27, d_3547_v0_)
        hi25_ = (self).f29
        for d_3562_i4_ in range((self).f31, hi25_):
            d_3563_v11_: _dafny.Array
            nw572_ = _dafny.Array(False, 14)
            d_3563_v11_ = nw572_
            index571_ = default__.safeIndex(328, (d_3563_v11_).length(0))
            (d_3563_v11_)[index571_] = (self).f27
            d_3564_v12_: D20
            d_3564_v12_ = D20_DC50()
            d_3565_v13_: D20
            d_3565_v13_ = D20_DC51(d_3564_v12_)
            d_3566_v14_: _dafny.Set
            d_3566_v14_ = _dafny.Set({d_3565_v13_})
            index572_ = default__.safeIndex(328, (d_3563_v11_).length(0))
            (d_3563_v11_)[index572_] = not((D20_DC51(d_3564_v12_)) in (d_3566_v14_))
            (globalState).f25 = not((d_3563_v11_)[default__.safeIndex(328, (d_3563_v11_).length(0))])
            d_3567_v15_: _dafny.Seq
            d_3567_v15_ = _dafny.SeqWithoutIsStrInference([not((self).f27)])
            d_3568_v16_: _dafny.Map
            d_3568_v16_ = _dafny.Map({(self).f29: (self).f31})
            index573_ = default__.safeIndex(328, (d_3563_v11_).length(0))
            (d_3563_v11_)[index573_] = ((default__.fm1((self).f27, d_3567_v15_, (d_3563_v11_)[default__.safeIndex(328, (d_3563_v11_).length(0))], globalState)) != (((d_3568_v16_)[(self).f28] if ((self).f28) in (d_3568_v16_) else 710))) and (True)
            d_3569_v17_: _dafny.Array
            nw573_ = _dafny.Array(_dafny.Array(None, 0), 2)
            d_3569_v17_ = nw573_
            r2 = d_3569_v17_
        nw574_ = _dafny.Array(_dafny.Array(None, 0), 26)
        r0 = nw574_
        r1 = (d_3547_v0_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_3570_i5_ in range(default__.abs(5))]))
        d_3571_v18_: _dafny.Array
        nw575_ = _dafny.Array(_dafny.Array(None, 0), 22)
        d_3571_v18_ = nw575_
        r2 = d_3571_v18_
        r3 = (self).f27
        return r0, r1, r2, r3

    def m4(self, p0, p1, globalState):
        r0: bool = False
        r1: bool = False
        hi26_ = (self).f29
        for d_3572_i0_ in range((self).f28, hi26_):
            d_3573_v0_: _dafny.Seq
            d_3573_v0_ = _dafny.SeqWithoutIsStrInference([(self).f31, (self).f28, (self).f28])
            (globalState).f12 = ((_dafny.MultiSet(d_3573_v0_) if True else _dafny.MultiSet([d_3572_i0_]))).ispropersubset(_dafny.MultiSet([(self).f29]))
            d_3574_v1_: _dafny.MultiSet
            d_3574_v1_ = _dafny.MultiSet([(self).f27])
            d_3575_v2_: T0
            nw576_ = C4()
            nw576_.ctor__(d_3574_v1_, p1)
            d_3575_v2_ = nw576_
            d_3576_v3_: _dafny.Map
            d_3576_v3_ = _dafny.Map({(self).f31: d_3575_v2_})
            d_3576_v3_ = d_3576_v3_
            d_3577_v4_: _dafny.Seq
            d_3577_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xdiai"))
            if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ky"))) < (d_3577_v4_):
                d_3578_v5_: _dafny.Map
                d_3578_v5_ = _dafny.Map({(self).f30: (self).f29})
                d_3579_v6_: str
                d_3579_v6_ = _dafny.CodePoint('s')
                d_3580_v7_: _dafny.Map
                d_3580_v7_ = _dafny.Map({(d_3575_v2_).f27: d_3579_v6_})
                d_3581_v8_: D9
                d_3581_v8_ = D9_DC20(d_3580_v7_)
                pat_let_tv76_ = d_3580_v7_
                d_3582_v9_: _dafny.Set
                def iife323_(_pat_let98_0):
                    def iife324_(d_3583_dt__update__tmp_h0_):
                        def iife325_(_pat_let99_0):
                            def iife326_(d_3584_dt__update_hcf30_h0_):
                                return D9_DC20(d_3584_dt__update_hcf30_h0_)
                            return iife326_(_pat_let99_0)
                        return iife325_(pat_let_tv76_)
                    return iife324_(_pat_let98_0)
                d_3582_v9_ = _dafny.Set({iife323_(d_3581_v8_)})
                d_3585_v10_: _dafny.Map
                d_3585_v10_ = _dafny.Map({(d_3578_v5_) | (d_3578_v5_): (d_3582_v9_).intersection(d_3582_v9_)})
                d_3585_v10_ = (d_3585_v10_).set(d_3578_v5_, (_dafny.Set({d_3581_v8_, d_3581_v8_})) | (d_3582_v9_))
                d_3586_v11_: _dafny.Array
                def lambda148_(d_3587_v0_):
                    def lambda149_(d_3588_i1_):
                        return (((d_3587_v0_).set(default__.safeIndex((self).f29, len(d_3587_v0_)), (self).f28)) + (d_3587_v0_)).set(default__.safeIndex((0) - ((self).f28), len(((d_3587_v0_).set(default__.safeIndex((self).f29, len(d_3587_v0_)), (self).f28)) + (d_3587_v0_))), len(_dafny.Set({424})))

                    return lambda149_

                init86_ = lambda148_(d_3573_v0_)
                nw577_ = _dafny.Array(None, 21)
                for i0_86_ in range(nw577_.length(0)):
                    nw577_[i0_86_] = init86_(i0_86_)
                d_3586_v11_ = nw577_
                index574_ = default__.safeIndex(548, (d_3586_v11_).length(0))
                (d_3586_v11_)[index574_] = _dafny.SeqWithoutIsStrInference([(self).f28 for d_3589_i2_ in range(default__.abs(932))])
                index575_ = default__.safeIndex(548, (d_3586_v11_).length(0))
                (d_3586_v11_)[index575_] = ((d_3573_v0_) + (d_3573_v0_) if (self).f27 else (d_3573_v0_) + (_dafny.SeqWithoutIsStrInference([(self).f28, default__.fm1((d_3575_v2_).f27, default__.fm52(d_3572_i0_, True, d_3574_v1_, globalState), (self).f27, globalState)])))
                d_3590_v12_: _dafny.Map
                d_3590_v12_ = _dafny.Map({(d_3575_v2_).f27: len(d_3577_v4_)})
                (globalState).f13 = (d_3590_v12_) | (d_3590_v12_)
                d_3591_v13_: D0
                d_3591_v13_ = D0_DC0(d_3572_i0_)
                d_3591_v13_ = default__.fm11((d_3575_v2_).f27, (d_3575_v2_).f27, (0) - (((self).f28) * ((self).f28)), globalState)
                d_3592_v14_: _dafny.Array
                nw578_ = _dafny.Array(D29.default()(), 15)
                d_3592_v14_ = nw578_
                d_3592_v14_ = d_3592_v14_
            elif True:
                (globalState).f25 = ((d_3575_v2_).f27) == (p1)
                d_3593_v15_: _dafny.Set
                d_3593_v15_ = _dafny.Set({(self).f28})
                d_3594_v16_: _dafny.Map
                d_3594_v16_ = _dafny.Map({p1: False})
                d_3595_v17_: str
                d_3595_v17_ = _dafny.CodePoint('o')
                d_3596_v18_: _dafny.Map
                d_3596_v18_ = _dafny.Map({d_3594_v16_: d_3595_v17_})
                d_3597_v19_: _dafny.Set
                d_3597_v19_ = _dafny.Set({(self).f28, (0) - ((0) - (len(d_3593_v15_))), (_dafny.MultiSet([_dafny.Set({d_3572_i0_}), _dafny.Set({len(d_3573_v0_), len(d_3596_v18_)}), d_3593_v15_])).cardinality, -892})
                def iife327_():
                    coll127_ = _dafny.Set()
                    compr_127_: int
                    for compr_127_ in ((d_3573_v0_) + (d_3573_v0_)).Elements:
                        d_3598_v20_: int = compr_127_
                        if (d_3598_v20_) in ((d_3573_v0_) + (d_3573_v0_)):
                            coll127_ = coll127_.union(_dafny.Set([(d_3598_v20_) + (350)]))
                    return _dafny.Set(coll127_)
                d_3593_v15_ = iife327_()
                
                d_3599_v21_: _dafny.Array
                nw579_ = _dafny.Array(False, 27)
                d_3599_v21_ = nw579_
                index576_ = default__.safeIndex(561, (d_3599_v21_).length(0))
                (d_3599_v21_)[index576_] = ((self).f31) < ((self).f31)
                index577_ = default__.safeIndex(561, (d_3599_v21_).length(0))
                rhs544_ = p1
                rhs545_ = p1
                lhs467_ = globalState
                lhs468_ = d_3599_v21_
                lhs469_ = default__.safeIndex(561, (d_3599_v21_).length(0))
                lhs467_.f14 = rhs544_
                lhs468_[lhs469_] = rhs545_
                d_3600_v22_: D13
                d_3600_v22_ = D13_DC29((self).f31, (self).f28, p1)
                (globalState).f6 = (d_3600_v22_).cf40
                d_3601_v23_: _dafny.Map
                d_3601_v23_ = _dafny.Map({d_3572_i0_: (d_3599_v21_)[default__.safeIndex(561, (d_3599_v21_).length(0))]})
                (globalState).f6 = (0) - ((966 if (False) == (((d_3601_v23_)[(self).f31] if ((self).f31) in (d_3601_v23_) else (d_3575_v2_).f27)) else (self).f31))
            index578_ = default__.safeIndex(659, ((self).f30).length(0))
            ((self).f30)[index578_] = ((self).f28) - ((self).f28)
            index579_ = default__.safeIndex(659, ((self).f30).length(0))
            ((self).f30)[index579_] = d_3572_i0_
        d_3602_v24_: _dafny.Set
        d_3602_v24_ = _dafny.Set({(self).f29, -596})
        d_3603_v25_: C7
        nw580_ = C7()
        nw580_.ctor__(p1, (self).f31, (self).f31, len(d_3602_v24_), (self).f27)
        d_3603_v25_ = nw580_
        d_3604_v26_: _dafny.Map
        d_3604_v26_ = _dafny.Map({d_3603_v25_: (self).f27})
        d_3605_v27_: _dafny.Seq
        d_3605_v27_ = _dafny.SeqWithoutIsStrInference([d_3604_v26_])
        d_3606_v28_: _dafny.Map
        d_3606_v28_ = _dafny.Map({(d_3605_v27_)[default__.safeIndex((d_3603_v25_).f36, len(d_3605_v27_))]: (self).f27})
        hi27_ = len(_dafny.SeqWithoutIsStrInference([d_3603_v25_.f35, d_3603_v25_.f35]))
        for d_3607_i3_ in range(len(d_3606_v28_), hi27_):
            d_3608_v29_: _dafny.Seq
            d_3608_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ptgbtq"))
            d_3609_v30_: _dafny.Seq
            d_3609_v30_ = _dafny.SeqWithoutIsStrInference([(self).f27])
            d_3610_v31_: str
            d_3610_v31_ = _dafny.CodePoint('m')
            d_3608_v29_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "arltihk"))) + ((d_3608_v29_).set(default__.safeIndex(len(d_3609_v30_), len(d_3608_v29_)), d_3610_v31_))
            d_3611_v32_: _dafny.MultiSet
            d_3611_v32_ = _dafny.MultiSet([(self).f30])
            d_3612_v33_: D6
            d_3612_v33_ = D6_DC13(d_3609_v30_)
            index580_ = default__.safeIndex(592, ((self).f30).length(0))
            ((self).f30)[index580_] = ((d_3611_v32_)[(self).f30] if ((self).f30) in (d_3611_v32_) else len((d_3612_v33_).cf21))
            index581_ = default__.safeIndex(592, ((self).f30).length(0))
            ((self).f30)[index581_] = (d_3603_v25_).f36
            d_3613_v34_: _dafny.MultiSet
            d_3613_v34_ = _dafny.MultiSet([852, ((self).f30)[default__.safeIndex(592, ((self).f30).length(0))], (d_3603_v25_).f36])
            d_3613_v34_ = d_3613_v34_
            d_3614_v35_: D11
            d_3614_v35_ = D11_DC25()
            d_3615_v36_: _dafny.Map
            d_3615_v36_ = _dafny.Map({d_3614_v35_: d_3610_v31_})
            d_3616_v37_: _dafny.Seq
            d_3616_v37_ = _dafny.SeqWithoutIsStrInference([((self).f30)[default__.safeIndex(592, ((self).f30).length(0))], (d_3603_v25_).f36])
            d_3617_v38_: _dafny.Array
            def lambda150_(d_3618_i4_):
                return (d_3618_i4_) * (((self).f30)[default__.safeIndex(592, ((self).f30).length(0))])

            init87_ = lambda150_
            nw581_ = _dafny.Array(None, 26)
            for i0_87_ in range(nw581_.length(0)):
                nw581_[i0_87_] = init87_(i0_87_)
            d_3617_v38_ = nw581_
            d_3619_v39_: _dafny.Seq
            d_3619_v39_ = _dafny.SeqWithoutIsStrInference([d_3617_v38_, d_3617_v38_, d_3617_v38_, d_3617_v38_, d_3617_v38_])
            index582_ = default__.safeIndex(592, ((self).f30).length(0))
            rhs546_ = ((self).f29) == (default__.safeDivisionInt((d_3603_v25_).f36, d_3607_i3_))
            rhs547_ = (self).f31
            rhs548_ = ((len((d_3615_v36_).set(d_3614_v35_, d_3610_v31_))) + ((d_3616_v37_)[default__.safeIndex(len(d_3608_v29_), len(d_3616_v37_))])) != (default__.safeModuloInt(d_3607_i3_, ((self).f30)[default__.safeIndex(592, ((self).f30).length(0))]))
            rhs549_ = _dafny.MultiSet([(0) - (d_3607_i3_)])
            rhs550_ = (d_3617_v38_) not in (d_3619_v39_)
            lhs470_ = globalState
            lhs471_ = (self).f30
            lhs472_ = default__.safeIndex(592, ((self).f30).length(0))
            lhs473_ = globalState
            lhs470_.f5 = rhs546_
            lhs471_[lhs472_] = rhs547_
            lhs473_.f21 = rhs548_
            d_3613_v34_ = rhs549_
            r0 = rhs550_
        d_3620_v40_: C0
        nw582_ = C0()
        nw582_.ctor__()
        d_3620_v40_ = nw582_
        d_3621_v41_: _dafny.Map
        d_3621_v41_ = _dafny.Map({704: d_3620_v40_})
        d_3622_v42_: _dafny.Map
        d_3622_v42_ = d_3621_v41_
        d_3623_v43_: _dafny.Map
        d_3623_v43_ = _dafny.Map({(self).fm7(globalState): d_3622_v42_})
        d_3623_v43_ = (d_3623_v43_) | (d_3623_v43_)
        d_3624_v44_: C0
        nw583_ = C0()
        nw583_.ctor__()
        d_3624_v44_ = nw583_
        d_3625_v45_: _dafny.Array
        def lambda151_(d_3626_p1_):
            def lambda152_(d_3627_i5_):
                return d_3626_p1_

            return lambda152_

        init88_ = lambda151_(p1)
        nw584_ = _dafny.Array(None, 21)
        for i0_88_ in range(nw584_.length(0)):
            nw584_[i0_88_] = init88_(i0_88_)
        d_3625_v45_ = nw584_
        index583_ = default__.safeIndex(560, (d_3625_v45_).length(0))
        (d_3625_v45_)[index583_] = False
        d_3628_v46_: _dafny.Seq
        d_3628_v46_ = _dafny.SeqWithoutIsStrInference([(self).f27, d_3603_v25_.f35])
        d_3629_v47_: _dafny.Seq
        d_3629_v47_ = _dafny.SeqWithoutIsStrInference([(d_3628_v46_)[default__.safeIndex((self).f29, len(d_3628_v46_))]])
        index584_ = default__.safeIndex(560, (d_3625_v45_).length(0))
        (d_3625_v45_)[index584_] = ((self).f29) == (default__.fm1(d_3603_v25_.f35, d_3629_v47_, p1, globalState))
        d_3630_v48_: _dafny.Seq
        d_3630_v48_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bjoyujqtl"))
        d_3631_v49_: str
        d_3631_v49_ = _dafny.CodePoint('v')
        d_3632_v50_: _dafny.Map
        d_3632_v50_ = _dafny.Map({d_3631_v49_: _dafny.Map({(d_3603_v25_).f36: default__.fm24((d_3625_v45_)[default__.safeIndex(560, (d_3625_v45_).length(0))], globalState)})})
        d_3630_v48_ = (self).fm2(default__.safeModuloInt(811, (self).f31), d_3632_v50_, globalState)
        r0 = not((d_3625_v45_)[default__.safeIndex(560, (d_3625_v45_).length(0))])
        d_3633_v51_: _dafny.Seq
        d_3633_v51_ = _dafny.SeqWithoutIsStrInference([6, (self).f28])
        r1 = (((0) - ((d_3633_v51_)[default__.safeIndex((d_3603_v25_).f36, len(d_3633_v51_))])) - ((self).f31)) != ((0) - (len(_dafny.SeqWithoutIsStrInference([d_3603_v25_.f35, default__.fm0(d_3630_v48_, (d_3625_v45_)[default__.safeIndex(560, (d_3625_v45_).length(0))], (self).f27, default__.fm21((self).f28, globalState), globalState)]))))
        return r0, r1

    @property
    def f30(self):
        return self._f30
    @property
    def f31(self):
        return self._f31
