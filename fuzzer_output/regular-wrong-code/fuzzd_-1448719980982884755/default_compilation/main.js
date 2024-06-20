// Dafny program main.dfy compiled into JavaScript
// Copyright by the contributors to the Dafny Project
// SPDX-License-Identifier: MIT

const BigNumber = require('bignumber.js');
BigNumber.config({ MODULO_MODE: BigNumber.EUCLID })
let _dafny = (function() {
  let $module = {};
  $module.areEqual = function(a, b) {
    if (typeof a === 'string' && b instanceof _dafny.Seq) {
      // Seq.equals(string) works as expected,
      // and the catch-all else block handles that direction.
      // But the opposite direction doesn't work; handle it here.
      return b.equals(a);
    } else if (typeof a === 'number' && BigNumber.isBigNumber(b)) {
      // This conditional would be correct even without the `typeof a` part,
      // but in most cases it's probably faster to short-circuit on a `typeof`
      // than to call `isBigNumber`. (But it remains to properly test this.)
      return b.isEqualTo(a);
    } else if (typeof a !== 'object' || a === null || b === null) {
      return a === b;
    } else if (BigNumber.isBigNumber(a)) {
      return a.isEqualTo(b);
    } else if (a._tname !== undefined || (Array.isArray(a) && a.constructor.name == "Array")) {
      return a === b;  // pointer equality
    } else {
      return a.equals(b);  // value-type equality
    }
  }
  $module.toString = function(a) {
    if (a === null) {
      return "null";
    } else if (typeof a === "number") {
      return a.toFixed();
    } else if (BigNumber.isBigNumber(a)) {
      return a.toFixed();
    } else if (a._tname !== undefined) {
      return a._tname;
    } else {
      return a.toString();
    }
  }
  $module.escapeCharacter = function(cp) {
    let s = String.fromCodePoint(cp.value)
    switch (s) {
      case '\n': return "\\n";
      case '\r': return "\\r";
      case '\t': return "\\t";
      case '\0': return "\\0";
      case '\'': return "\\'";
      case '\"': return "\\\"";
      case '\\': return "\\\\";
      default: return s;
    };
  }
  $module.NewObject = function() {
    return { _tname: "object" };
  }
  $module.InstanceOfTrait = function(obj, trait) {
    return obj._parentTraits !== undefined && obj._parentTraits().includes(trait);
  }
  $module.Rtd_bool = class {
    static get Default() { return false; }
  }
  $module.Rtd_char = class {
    static get Default() { return 'D'; }  // See CharType.DefaultValue in Dafny source code
  }
  $module.Rtd_codepoint = class {
    static get Default() { return new _dafny.CodePoint('D'.codePointAt(0)); }
  }
  $module.Rtd_int = class {
    static get Default() { return BigNumber(0); }
  }
  $module.Rtd_number = class {
    static get Default() { return 0; }
  }
  $module.Rtd_ref = class {
    static get Default() { return null; }
  }
  $module.Rtd_array = class {
    static get Default() { return []; }
  }
  $module.ZERO = new BigNumber(0);
  $module.ONE = new BigNumber(1);
  $module.NUMBER_LIMIT = new BigNumber(0x20).multipliedBy(0x1000000000000);  // 2^53
  $module.Tuple = class Tuple extends Array {
    constructor(...elems) {
      super(...elems);
    }
    toString() {
      return "(" + arrayElementsToString(this) + ")";
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static Default(...values) {
      return Tuple.of(...values);
    }
    static Rtd(...rtdArgs) {
      return {
        Default: Tuple.from(rtdArgs, rtd => rtd.Default)
      };
    }
  }
  $module.Set = class Set extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Set.Empty;
    }
    toString() {
      return "{" + arrayElementsToString(this) + "}";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Set();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new Set();
      for (let k of elmts) {
        s.add(k);
      }
      return s;
    }
    contains(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i], k)) {
          return true;
        }
      }
      return false;
    }
    add(k) {  // mutates the Set; use only during construction
      if (!this.contains(k)) {
        this.push(k);
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        if (!other.contains(e)) {
          return false;
        }
      }
      return true;
    }
    get Elements() {
      return this;
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = Set.of(...this);
        for (let k of that) {
          s.add(k);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new Set();
        for (let k of this) {
          if (that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length == 0 || that.length == 0) {
        return this;
      } else {
        let s = new Set();
        for (let k of this) {
          if (!that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      for (let k of this) {
        if (that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsSubsetOf(that) {
      if (that.length < this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      if (that.length <= this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    get AllSubsets() {
      return this.AllSubsets_();
    }
    *AllSubsets_() {
      // Start by putting all set elements into a list, but don't include null
      let elmts = Array.of(...this);
      let n = elmts.length;
      let which = new Array(n);
      which.fill(false);
      let a = [];
      while (true) {
        yield Set.of(...a);
        // "add 1" to "which", as if doing a carry chain.  For every digit changed, change the membership of the corresponding element in "a".
        let i = 0;
        for (; i < n && which[i]; i++) {
          which[i] = false;
          // remove elmts[i] from a
          for (let j = 0; j < a.length; j++) {
            if (_dafny.areEqual(a[j], elmts[i])) {
              // move the last element of a into slot j
              a[j] = a[-1];
              a.pop();
              break;
            }
          }
        }
        if (i === n) {
          // we have cycled through all the subsets
          break;
        }
        which[i] = true;
        a.push(elmts[i]);
      }
    }
  }
  $module.MultiSet = class MultiSet extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return MultiSet.Empty;
    }
    toString() {
      let s = "multiset{";
      let sep = "";
      for (let e of this) {
        let [k, n] = e;
        let ks = _dafny.toString(k);
        while (!n.isZero()) {
          n = n.minus(1);
          s += sep + ks;
          sep = ", ";
        }
      }
      s += "}";
      return s;
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new MultiSet();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new MultiSet();
      for (let e of elmts) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    static FromArray(arr) {
      let s = new MultiSet();
      for (let e of arr) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    cardinality() {
      let c = _dafny.ZERO;
      for (let e of this) {
        let [k, n] = e;
        c = c.plus(n);
      }
      return c;
    }
    clone() {
      let s = new MultiSet();
      for (let e of this) {
        let [k, n] = e;
        s.push([k, n]);  // make sure to create a new array [k, n] here
      }
      return s;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return _dafny.ZERO;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return !this.get(k).isZero();
    }
    add(k, n) {
      let i = this.findIndex(k);
      if (i === this.length) {
        this.push([k, n]);
      } else {
        let m = this[i][1];
        this[i] = [k, m.plus(n)];
      }
    }
    update(k, n) {
      let i = this.findIndex(k);
      if (i < this.length && this[i][1].isEqualTo(n)) {
        return this;
      } else if (i === this.length && n.isZero()) {
        return this;
      } else if (i === this.length) {
        let m = this.slice();
        m.push([k, n]);
        return m;
      } else {
        let m = this.slice();
        m[i] = [k, n];
        return m;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let e of this) {
        let [k, n] = e;
        let m = other.get(k);
        if (!n.isEqualTo(m)) {
          return false;
        }
      }
      return this.cardinality().isEqualTo(other.cardinality());
    }
    get Elements() {
      return this.Elements_();
    }
    *Elements_() {
      for (let i = 0; i < this.length; i++) {
        let [k, n] = this[i];
        while (!n.isZero()) {
          yield k;
          n = n.minus(1);
        }
      }
    }
    get UniqueElements() {
      return this.UniqueElements_();
    }
    *UniqueElements_() {
      for (let e of this) {
        let [k, n] = e;
        if (!n.isZero()) {
          yield k;
        }
      }
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = this.clone();
        for (let e of that) {
          let [k, n] = e;
          s.add(k, n);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let m = that.get(k);
          if (!m.isZero()) {
            s.push([k, m.isLessThan(n) ? m : n]);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length === 0 || that.length === 0) {
        return this;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let d = n.minus(that.get(k));
          if (d.isGreaterThan(0)) {
            s.push([k, d]);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      let intersection = this.Intersect(that);
      return intersection.cardinality().isZero();
    }
    IsSubsetOf(that) {
      for (let e of this) {
        let [k, n] = e;
        let m = that.get(k);
        if (!n.isLessThanOrEqualTo(m)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      return this.IsSubsetOf(that) && this.cardinality().isLessThan(that.cardinality());
    }
  }
  $module.CodePoint = class CodePoint {
    constructor(value) {
      this.value = value
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      return this.value === other.value
    }
    isLessThan(other) {
      return this.value < other.value
    }
    isLessThanOrEqual(other) {
      return this.value <= other.value
    }
    toString() {
      return "'" + $module.escapeCharacter(this) + "'";
    }
    static isCodePoint(i) {
      return (
        (_dafny.ZERO.isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0xD800))) ||
        (new BigNumber(0xE000).isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0x11_0000))))
    }
  }
  $module.Seq = class Seq extends Array {
    constructor(...elems) {
      super(...elems);
    }
    static get Default() {
      return Seq.of();
    }
    static Create(n, init) {
      return Seq.from({length: n}, (_, i) => init(new BigNumber(i)));
    }
    static UnicodeFromString(s) {
      return new Seq(...([...s].map(c => new _dafny.CodePoint(c.codePointAt(0)))))
    }
    toString() {
      return "[" + arrayElementsToString(this) + "]";
    }
    toVerbatimString(asLiteral) {
      if (asLiteral) {
        return '"' + this.map(c => _dafny.escapeCharacter(c)).join("") + '"';
      } else {
        return this.map(c => String.fromCodePoint(c.value)).join("");
      }
    }
    static update(s, i, v) {
      if (typeof s === "string") {
        let p = s.slice(0, i);
        let q = s.slice(i.toNumber() + 1);
        return p.concat(v, q);
      } else {
        let t = s.slice();
        t[i] = v;
        return t;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static contains(s, k) {
      if (typeof s === "string") {
        return s.includes(k);
      } else {
        for (let x of s) {
          if (_dafny.areEqual(x, k)) {
            return true;
          }
        }
        return false;
      }
    }
    get Elements() {
      return this;
    }
    get UniqueElements() {
      return _dafny.Set.fromElements(...this);
    }
    static Concat(a, b) {
      if (typeof a === "string" || typeof b === "string") {
        // string concatenation, so make sure both operands are strings before concatenating
        if (typeof a !== "string") {
          // a must be a Seq
          a = a.join("");
        }
        if (typeof b !== "string") {
          // b must be a Seq
          b = b.join("");
        }
        return a + b;
      } else {
        // ordinary concatenation
        let r = Seq.of(...a);
        r.push(...b);
        return r;
      }
    }
    static JoinIfPossible(x) {
      try { return x.join(""); } catch(_error) { return x; }
    }
    static IsPrefixOf(a, b) {
      if (b.length < a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
    static IsProperPrefixOf(a, b) {
      if (b.length <= a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
  }
  $module.Map = class Map extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Map.of();
    }
    toString() {
      return "map[" + this.map(maplet => _dafny.toString(maplet[0]) + " := " + _dafny.toString(maplet[1])).join(", ") + "]";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Map();
      }
      return this._empty;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return undefined;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return this.findIndex(k) < this.length;
    }
    update(k, v) {
      let m = this.slice();
      m.updateUnsafe(k, v);
      return m;
    }
    // Similar to update, but make the modification in-place.
    // Meant to be used in the map constructor.
    updateUnsafe(k, v) {
      let m = this;
      let i = m.findIndex(k);
      m[i] = [k, v];
      return m;
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        let [k, v] = e;
        let w = other.get(k);
        if (w === undefined || !_dafny.areEqual(v, w)) {
          return false;
        }
      }
      return true;
    }
    get Keys() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(k);
      }
      return s;
    }
    get Values() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.add(v);
      }
      return s;
    }
    get Items() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(_dafny.Tuple.of(k, v));
      }
      return s;
    }
    Merge(that) {
      let m = that.slice();
      for (let e of this) {
        let [k, v] = e;
        let i = m.findIndex(k);
        if (i == m.length) {
          m[i] = [k, v];
        }
      }
      return m;
    }
    Subtract(keys) {
      if (this.length === 0 || keys.length === 0) {
        return this;
      }
      let m = new Map();
      for (let e of this) {
        let [k, v] = e;
        if (!keys.contains(k)) {
          m[m.length] = e;
        }
      }
      return m;
    }
  }
  $module.newArray = function(initValue, ...dims) {
    return { dims: dims, elmts: buildArray(initValue, ...dims) };
  }
  $module.BigOrdinal = class BigOrdinal {
    static get Default() {
      return _dafny.ZERO;
    }
    static IsLimit(ord) {
      return ord.isZero();
    }
    static IsSucc(ord) {
      return ord.isGreaterThan(0);
    }
    static Offset(ord) {
      return ord;
    }
    static IsNat(ord) {
      return true;  // at run time, every ORDINAL is a natural number
    }
  }
  $module.BigRational = class BigRational {
    static get ZERO() {
      if (this._zero === undefined) {
        this._zero = new BigRational(_dafny.ZERO);
      }
      return this._zero;
    }
    constructor (n, d) {
      // requires d === undefined || 1 <= d
      this.num = n;
      this.den = d === undefined ? _dafny.ONE : d;
      // invariant 1 <= den || (num == 0 && den == 0)
    }
    static get Default() {
      return _dafny.BigRational.ZERO;
    }
    // We need to deal with the special case `num == 0 && den == 0`, because
    // that's what C#'s default struct constructor will produce for BigRational. :(
    // To deal with it, we ignore `den` when `num` is 0.
    toString() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num.toFixed() + ".0";
      }
      let answer = this.dividesAPowerOf10(this.den);
      if (answer !== undefined) {
        let n = this.num.multipliedBy(answer[0]);
        let log10 = answer[1];
        let sign, digits;
        if (this.num.isLessThan(0)) {
          sign = "-"; digits = n.negated().toFixed();
        } else {
          sign = ""; digits = n.toFixed();
        }
        if (log10 < digits.length) {
          let digitCount = digits.length - log10;
          return sign + digits.slice(0, digitCount) + "." + digits.slice(digitCount);
        } else {
          return sign + "0." + "0".repeat(log10 - digits.length) + digits;
        }
      } else {
        return "(" + this.num.toFixed() + ".0 / " + this.den.toFixed() + ".0)";
      }
    }
    isPowerOf10(x) {
      if (x.isZero()) {
        return undefined;
      }
      let log10 = 0;
      while (true) {  // invariant: x != 0 && x * 10^log10 == old(x)
        if (x.isEqualTo(1)) {
          return log10;
        } else if (x.mod(10).isZero()) {
          log10++;
          x = x.dividedToIntegerBy(10);
        } else {
          return undefined;
        }
      }
    }
    dividesAPowerOf10(i) {
      let factor = _dafny.ONE;
      let log10 = 0;
      if (i.isLessThanOrEqualTo(_dafny.ZERO)) {
        return undefined;
      }

      // invariant: 1 <= i && i * 10^log10 == factor * old(i)
      while (i.mod(10).isZero()) {
        i = i.dividedToIntegerBy(10);
       log10++;
      }

      while (i.mod(5).isZero()) {
        i = i.dividedToIntegerBy(5);
        factor = factor.multipliedBy(2);
        log10++;
      }
      while (i.mod(2).isZero()) {
        i = i.dividedToIntegerBy(2);
        factor = factor.multipliedBy(5);
        log10++;
      }

      if (i.isEqualTo(_dafny.ONE)) {
        return [factor, log10];
      } else {
        return undefined;
      }
    }
    toBigNumber() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num;
      } else if (this.num.isGreaterThan(0)) {
        return this.num.dividedToIntegerBy(this.den);
      } else {
        return this.num.minus(this.den).plus(1).dividedToIntegerBy(this.den);
      }
    }
    isInteger() {
      return this.equals(new _dafny.BigRational(this.toBigNumber(), _dafny.ONE));
    }
    // Returns values such that aa/dd == a and bb/dd == b.
    normalize(b) {
      let a = this;
      let aa, bb, dd;
      if (a.num.isZero()) {
        aa = a.num;
        bb = b.num;
        dd = b.den;
      } else if (b.num.isZero()) {
        aa = a.num;
        dd = a.den;
        bb = b.num;
      } else {
        let gcd = BigNumberGcd(a.den, b.den);
        let xx = a.den.dividedToIntegerBy(gcd);
        let yy = b.den.dividedToIntegerBy(gcd);
        // We now have a == a.num / (xx * gcd) and b == b.num / (yy * gcd).
        aa = a.num.multipliedBy(yy);
        bb = b.num.multipliedBy(xx);
        dd = a.den.multipliedBy(yy);
      }
      return [aa, bb, dd];
    }
    compareTo(that) {
      // simple things first
      let asign = this.num.isZero() ? 0 : this.num.isLessThan(0) ? -1 : 1;
      let bsign = that.num.isZero() ? 0 : that.num.isLessThan(0) ? -1 : 1;
      if (asign < 0 && 0 <= bsign) {
        return -1;
      } else if (asign <= 0 && 0 < bsign) {
        return -1;
      } else if (bsign < 0 && 0 <= asign) {
        return 1;
      } else if (bsign <= 0 && 0 < asign) {
        return 1;
      }
      let [aa, bb, dd] = this.normalize(that);
      if (aa.isLessThan(bb)) {
        return -1;
      } else if (aa.isEqualTo(bb)){
        return 0;
      } else {
        return 1;
      }
    }
    equals(that) {
      return this.compareTo(that) === 0;
    }
    isLessThan(that) {
      return this.compareTo(that) < 0;
    }
    isAtMost(that) {
      return this.compareTo(that) <= 0;
    }
    plus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.plus(bb), dd);
    }
    minus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.minus(bb), dd);
    }
    negated() {
      return new BigRational(this.num.negated(), this.den);
    }
    multipliedBy(b) {
      return new BigRational(this.num.multipliedBy(b.num), this.den.multipliedBy(b.den));
    }
    dividedBy(b) {
      let a = this;
      // Compute the reciprocal of b
      let bReciprocal;
      if (b.num.isGreaterThan(0)) {
        bReciprocal = new BigRational(b.den, b.num);
      } else {
        // this is the case b.num < 0
        bReciprocal = new BigRational(b.den.negated(), b.num.negated());
      }
      return a.multipliedBy(bReciprocal);
    }
  }
  $module.EuclideanDivisionNumber = function(a, b) {
    if (0 <= a) {
      if (0 <= b) {
        // +a +b: a/b
        return Math.floor(a / b);
      } else {
        // +a -b: -(a/(-b))
        return -Math.floor(a / -b);
      }
    } else {
      if (0 <= b) {
        // -a +b: -((-a-1)/b) - 1
        return -Math.floor((-a-1) / b) - 1;
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return Math.floor((-a-1) / -b) + 1;
      }
    }
  }
  $module.EuclideanDivision = function(a, b) {
    if (a.isGreaterThanOrEqualTo(0)) {
      if (b.isGreaterThanOrEqualTo(0)) {
        // +a +b: a/b
        return a.dividedToIntegerBy(b);
      } else {
        // +a -b: -(a/(-b))
        return a.dividedToIntegerBy(b.negated()).negated();
      }
    } else {
      if (b.isGreaterThanOrEqualTo(0)) {
        // -a +b: -((-a-1)/b) - 1
        return a.negated().minus(1).dividedToIntegerBy(b).negated().minus(1);
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return a.negated().minus(1).dividedToIntegerBy(b.negated()).plus(1);
      }
    }
  }
  $module.EuclideanModuloNumber = function(a, b) {
    let bp = Math.abs(b);
    if (0 <= a) {
      // +a: a % bp
      return a % bp;
    } else {
      // c = ((-a) % bp)
      // -a: bp - c if c > 0
      // -a: 0 if c == 0
      let c = (-a) % bp;
      return c === 0 ? c : bp - c;
    }
  }
  $module.ShiftLeft = function(b, n) {
    return b.multipliedBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.ShiftRight = function(b, n) {
    return b.dividedToIntegerBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.RotateLeft = function(b, n, w) {  // truncate(b << n) | (b >> (w - n))
    let x = _dafny.ShiftLeft(b, n).mod(new BigNumber(2).exponentiatedBy(w));
    let y = _dafny.ShiftRight(b, w - n);
    return x.plus(y);
  }
  $module.RotateRight = function(b, n, w) {  // (b >> n) | truncate(b << (w - n))
    let x = _dafny.ShiftRight(b, n);
    let y = _dafny.ShiftLeft(b, w - n).mod(new BigNumber(2).exponentiatedBy(w));;
    return x.plus(y);
  }
  $module.BitwiseAnd = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 & b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    return r;
  }
  $module.BitwiseOr = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 | b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseXor = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 ^ b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseNot = function(a, bits) {
    let r = _dafny.ZERO;
    let h = _dafny.ONE;
    for (let i = 0; i < bits; i++) {
      let bit = a.mod(2);
      if (bit.isZero()) {
        r = r.plus(h);
      }
      a = a.dividedToIntegerBy(2);
      h = h.multipliedBy(2);
    }
    return r;
  }
  $module.Quantifier = function(vals, frall, pred) {
    for (let u of vals) {
      if (pred(u) !== frall) { return !frall; }
    }
    return frall;
  }
  $module.PlusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) + b.charCodeAt(0));
  }
  $module.UnicodePlusChar = function(a, b) {
    return new _dafny.CodePoint(a.value + b.value);
  }
  $module.MinusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) - b.charCodeAt(0));
  }
  $module.UnicodeMinusChar = function(a, b) {
    return new _dafny.CodePoint(a.value - b.value);
  }
  $module.AllBooleans = function*() {
    yield false;
    yield true;
  }
  $module.AllChars = function*() {
    for (let i = 0; i < 0x10000; i++) {
      yield String.fromCharCode(i);
    }
  }
  $module.AllUnicodeChars = function*() {
    for (let i = 0; i < 0xD800; i++) {
      yield new _dafny.CodePoint(i);
    }
    for (let i = 0xE0000; i < 0x110000; i++) {
      yield new _dafny.CodePoint(i);
    }
  }
  $module.AllIntegers = function*() {
    yield _dafny.ZERO;
    for (let j = _dafny.ONE;; j = j.plus(1)) {
      yield j;
      yield j.negated();
    }
  }
  $module.IntegerRange = function*(lo, hi) {
    if (lo === null) {
      while (true) {
        hi = hi.minus(1);
        yield hi;
      }
    } else if (hi === null) {
      while (true) {
        yield lo;
        lo = lo.plus(1);
      }
    } else {
      while (lo.isLessThan(hi)) {
        yield lo;
        lo = lo.plus(1);
      }
    }
  }
  $module.SingleValue = function*(v) {
    yield v;
  }
  $module.HaltException = class HaltException extends Error {
    constructor(message) {
      super(message)
    }
  }
  $module.HandleHaltExceptions = function(f) {
    try {
      f()
    } catch (e) {
      if (e instanceof _dafny.HaltException) {
        process.stdout.write("[Program halted] " + e.message + "\n")
        process.exitCode = 1
      } else {
        throw e
      }
    }
  }
  $module.FromMainArguments = function(args) {
    var a = [...args];
    a.splice(0, 2, args[0] + " " + args[1]);
    return a;
  }
  $module.UnicodeFromMainArguments = function(args) {
    return $module.FromMainArguments(args).map(_dafny.Seq.UnicodeFromString);
  }
  return $module;

  // What follows are routines private to the Dafny runtime
  function buildArray(initValue, ...dims) {
    if (dims.length === 0) {
      return initValue;
    } else {
      let a = Array(dims[0].toNumber());
      let b = Array.from(a, (x) => buildArray(initValue, ...dims.slice(1)));
      return b;
    }
  }
  function arrayElementsToString(a) {
    // like `a.join(", ")`, but calling _dafny.toString(x) on every element x instead of x.toString()
    let s = "";
    let sep = "";
    for (let x of a) {
      s += sep + _dafny.toString(x);
      sep = ", ";
    }
    return s;
  }
  function BigNumberGcd(a, b){  // gcd of two non-negative BigNumber's
    while (true) {
      if (a.isZero()) {
        return b;
      } else if (b.isZero()) {
        return a;
      }
      if (a.isLessThan(b)) {
        b = b.modulo(a);
      } else {
        a = a.modulo(b);
      }
    }
  }
})();
// Dafny program systemModulePopulator.dfy compiled into JavaScript
let _System = (function() {
  let $module = {};

  $module.nat = class nat {
    constructor () {
    }
    static get Default() {
      return _dafny.ZERO;
    }
    static _Is(__source) {
      let _0_x = (__source);
      return (_dafny.ZERO).isLessThanOrEqualTo(_0_x);
    }
  };

  return $module;
})(); // end of module _System
let _module = (function() {
  let $module = {};

  $module.__default = class __default {
    constructor () {
      this._tname = "_module._default";
    }
    _parentTraits() {
      return [];
    }
    static abs(x) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return (new BigNumber(-1)).multipliedBy(x);
      } else {
        return x;
      }
    };
    static safeIndex(x, length) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return _dafny.ZERO;
      } else if ((length).isLessThanOrEqualTo(x)) {
        return (x).mod(length);
      } else {
        return x;
      }
    };
    static safeDivisionInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return _dafny.EuclideanDivision(x1, x2);
      }
    };
    static safeModuloInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return (x1).mod(x2);
      }
    };
    static fm0(p0, p1, p2, globalState) {
      return (_module.D1.create_DC2(new BigNumber(121))).dtor_cf2;
    };
    static fm2(globalState) {
      return (!(!(!((_dafny.MultiSet.fromElements(new BigNumber(-776))).IsProperSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(448)), function (_0_i0) {
        return new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(function () {
          let _coll0 = new _dafny.Set();
          for (const _compr_0 of _dafny.IntegerRange(new BigNumber(686), new BigNumber(805))) {
            let _1_v0 = _compr_0;
            if (((new BigNumber(686)).isLessThanOrEqualTo(_1_v0)) && ((_1_v0).isLessThan(new BigNumber(805)))) {
              _coll0.add((_1_v0).minus(new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(323),true))).length)));
            }
          }
          return _coll0;
        }()))).cardinality());
      }))))))) && (!(!(((false) ? (false) : (false)))));
    };
    static fm3(p0, globalState) {
      return _dafny.Set.fromElements(_dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("ifnkwt"), _dafny.Seq.UnicodeFromString("hammgsnta")), ((_dafny.ZERO).minus(new BigNumber(-352))).isLessThanOrEqualTo((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(false)).length))), (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(false)).length))).IsSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(240)), function (_2_i0) {
        return new BigNumber(277);
      }))), false, true);
    };
    static fm4(p0, p1, globalState) {
      let _source0 = _module.D4.create_DC11(new BigNumber(823));
      if (_source0.is_DC9) {
        let _3___mcc_h0 = (_source0).cf10;
        let _4___mcc_h1 = (_source0).cf11;
        let _5___mcc_h2 = (_source0).cf12;
        let _6_cf12 = _5___mcc_h2;
        let _7_cf11 = _4___mcc_h1;
        let _8_cf10 = _3___mcc_h0;
        return _dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("aqlkl")).length));
      } else if (_source0.is_DC10) {
        let _9___mcc_h3 = (_source0).cf13;
        let _10___mcc_h4 = (_source0).cf14;
        let _11___mcc_h5 = (_source0).cf15;
        let _12___mcc_h6 = (_source0).cf16;
        let _13___mcc_h7 = (_source0).cf17;
        let _14_cf17 = _13___mcc_h7;
        let _15_cf16 = _12___mcc_h6;
        let _16_cf15 = _11___mcc_h5;
        let _17_cf14 = _10___mcc_h4;
        let _18_cf13 = _9___mcc_h3;
        return (_dafny.Set.fromElements(_16_cf15)).Intersect(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(611)), ((_19_cf14) => function (_20_i0) {
          return _19_cf14;
        })(_17_cf14))).length), _16_cf15));
      } else if (_source0.is_DC11) {
        let _21___mcc_h8 = (_source0).cf18;
        let _22_cf18 = _21___mcc_h8;
        return _dafny.Set.fromElements(_22_cf18, _22_cf18, _22_cf18, (_dafny.ZERO).minus(_22_cf18), (_dafny.ZERO).minus((_dafny.ZERO).minus(_22_cf18)));
      } else if (_source0.is_DC8) {
        let _23___mcc_h9 = (_source0).cf9;
        let _24_cf9 = _23___mcc_h9;
        return _dafny.Set.fromElements(new BigNumber(-944), new BigNumber(396), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_24_cf9,new BigNumber(228))).length));
      } else {
        let _25___mcc_h10 = (_source0).cf19;
        let _26_cf19 = _25___mcc_h10;
        return function () {
          let _coll1 = new _dafny.Set();
          for (const _compr_1 of _dafny.IntegerRange(new BigNumber(73), new BigNumber(-142))) {
            let _27_v0 = _compr_1;
            if (((new BigNumber(73)).isLessThanOrEqualTo(_27_v0)) && ((_27_v0).isLessThan(new BigNumber(-142)))) {
              _coll1.add(_module.__default.safeModuloInt(_27_v0, new BigNumber(972)));
            }
          }
          return _coll1;
        }();
      }
    };
    static fm5(globalState) {
      if (!((new BigNumber(350)).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true, true, true, true, true)).length)), new BigNumber(207))).length)))) {
        return (_dafny.Map.Empty.slice().updateUnsafe(true,function () {
          let _coll2 = new _dafny.Map();
          for (const _compr_2 of _dafny.IntegerRange(new BigNumber(75), new BigNumber(606))) {
            let _28_v0 = _compr_2;
            if (((new BigNumber(75)).isLessThanOrEqualTo(_28_v0)) && ((_28_v0).isLessThan(new BigNumber(606)))) {
              _coll2.push([_module.__default.safeDivisionInt(_28_v0, new BigNumber(436)),true]);
            }
          }
          return _coll2;
        }())).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("bkytbt")).length),false)));
      } else {
        return (_dafny.Map.Empty.slice().updateUnsafe(true,function () {
          let _coll3 = new _dafny.Map();
          for (const _compr_3 of (_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-230), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-856))).length))).cardinality()), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(514),false)).length))).Elements) {
            let _29_v1 = _compr_3;
            if ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-230), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-856))).length))).cardinality()), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(514),false)).length))).contains(_29_v1)) {
              _coll3.push([_module.__default.safeModuloInt(_29_v1, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(340))).length))).length))),false]);
            }
          }
          return _coll3;
        }())).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("qevan")).length),false)));
      }
    };
    static fm6(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(165)), function (_30_i0) {
        return ((!(false)) ? (new _dafny.CodePoint('y'.codePointAt(0))) : (new _dafny.CodePoint('g'.codePointAt(0))));
      });
    };
    static fm7(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(914), (_dafny.ZERO).minus(new BigNumber(-417))), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(400)), _dafny.Seq.of(new BigNumber(553))));
    };
    static fm8(p0, p1, globalState) {
      return _module.D3.create_DC6(new BigNumber(522));
    };
    static fm9(p0, p1, p2, globalState) {
      if (false) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('q'.codePointAt(0));
      }
    };
    static fm10(p0, p1, p2, globalState) {
      return _module.D2.create_DC3(_dafny.Seq.UnicodeFromString("gywvd"));
    };
    static fm11(p0, p1, p2, p3, globalState) {
      return _dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(817),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(583)), function (_31_i0) {
        return new BigNumber(147);
      })).length))).length))).length))).length)), _dafny.Seq.of(new BigNumber(-387), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('x'.codePointAt(0)))).length), new BigNumber(868), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true)).length), new BigNumber((_dafny.Seq.UnicodeFromString("oendfkfg")).length))).length))).cardinality()))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(559)), function (_32_i1) {
        return new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("htm")).length))).length);
      }), _dafny.Seq.of(new BigNumber(766))));
    };
    static fm12(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of(true)), _dafny.Seq.of(!(!(true))));
    };
    static fm13(p0, globalState) {
      return (((_module.D8.create_DC19(_dafny.Set.fromElements(_dafny.Seq.of(true, false, false, true), _dafny.Seq.of(true)))).dtor_cf26).Union(function () {
        let _coll4 = new _dafny.Set();
        for (const _compr_4 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false),true)).Keys.Elements) {
          let _33_v0 = _compr_4;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false),true)).contains(_33_v0)) {
            _coll4.add(_33_v0);
          }
        }
        return _coll4;
      }())).Intersect(_dafny.Set.fromElements(_dafny.Seq.of(false, true, true, true, !(true)), _dafny.Seq.of(true, false), _dafny.Seq.of(true, true, false)));
    };
    static fm14(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,false))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,!(false)));
    };
    static m0(globalState) {
      let _34_v0;
      _34_v0 = false;
      let _35_v1;
      let _nw0 = new _module.C0();
      _nw0.__ctor(_34_v0);
      _35_v1 = _nw0;
      let _36_v2;
      let _init0 = function (_37_i0) {
        return _dafny.MultiSet.fromElements(new BigNumber(497), new BigNumber((_dafny.Set.fromElements(new BigNumber(-320), new BigNumber(-68), new BigNumber(69))).length), new BigNumber(-332));
      };
      let _nw1 = Array((new BigNumber(5)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw1.length); _i0_0++) {
        _nw1[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _36_v2 = _nw1;
      let _38_v3;
      _38_v3 = _dafny.Map.Empty.slice().updateUnsafe(_35_v1,_36_v2);
      _38_v3 = (_38_v3).update(_35_v1, _36_v2);
      let _39_v4;
      _39_v4 = new BigNumber(570);
      let _40_v5;
      _40_v5 = new _dafny.CodePoint('h'.codePointAt(0));
      let _pat_let_tv0 = _35_v1;
      let _pat_let_tv1 = _35_v1;
      if (function (_source1) {
        if (_source1.is_DC4) {
          let _41___mcc_h0 = (_source1).cf4;
          let _42___mcc_h1 = (_source1).cf5;
          let _43_cf5 = _42___mcc_h1;
          let _44_cf4 = _41___mcc_h0;
          return _pat_let_tv0.f15;
        } else {
          let _45___mcc_h2 = (_source1).cf3;
          let _46_cf3 = _45___mcc_h2;
          return !(_pat_let_tv1.f15);
        }
      }(_module.__default.fm10(_39_v4, _40_v5, _34_v0, globalState))) {
        let _47_v6;
        let _init1 = function (_48_i1) {
          return _dafny.Seq.UnicodeFromString("i");
        };
        let _nw2 = Array((new BigNumber(20)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw2.length); _i0_1++) {
          _nw2[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _47_v6 = _nw2;
        _47_v6 = _47_v6;
        let _49_v7;
        _49_v7 = _dafny.Seq.UnicodeFromString("rxcmd");
        let _50_v8;
        let _init2 = function (_51_i3) {
          return true;
        };
        let _nw3 = Array((new BigNumber(27)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw3.length); _i0_2++) {
          _nw3[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _50_v8 = _nw3;
        let _52_v9;
        _52_v9 = _module.D4.create_DC9(_50_v8, _35_v1, _34_v0);
        let _53_v10;
        _53_v10 = _dafny.Seq.of(_35_v1.f15, _35_v1.f15, _35_v1.f15);
        let _54_v11;
        _54_v11 = _dafny.MultiSet.fromElements(_35_v1.f15);
        let _55_v12;
        _55_v12 = _dafny.Set.fromElements(_35_v1, _35_v1);
        let _56_v13;
        _56_v13 = _dafny.Set.fromElements(_35_v1.f15);
        let _pat_let_tv2 = _35_v1;
        let _pat_let_tv3 = _50_v8;
        let _57_v14;
        let _nw4 = Array((new BigNumber(28)).toNumber());
        _nw4[(_dafny.ZERO).toNumber()] = _dafny.areEqual(_49_v7, _dafny.Seq.Create(_module.__default.abs(new BigNumber(391)), function (_58_i2) {
          return new _dafny.CodePoint('m'.codePointAt(0));
        }));
        _nw4[(_dafny.ONE).toNumber()] = _35_v1.f15;
        _nw4[(new BigNumber(2)).toNumber()] = _35_v1.f15;
        _nw4[(new BigNumber(3)).toNumber()] = (function (_pat_let0_0) {
          return function (_59_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_60_dt__update_hcf11_h0) {
                return function (_pat_let2_0) {
                  return function (_61_dt__update_hcf10_h0) {
                    return _module.D4.create_DC9(_61_dt__update_hcf10_h0, _60_dt__update_hcf11_h0, (_59_dt__update__tmp_h0).dtor_cf12);
                  }(_pat_let2_0);
                }(_pat_let_tv3);
              }(_pat_let1_0);
            }(_pat_let_tv2);
          }(_pat_let0_0);
        }(_52_v9)).dtor_cf12;
        _nw4[(new BigNumber(4)).toNumber()] = ((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_34_v0, _35_v1.f15, _35_v1.f15, _35_v1.f15, _35_v1.f15)).length))).isEqualTo(new BigNumber((_49_v7).length));
        _nw4[(new BigNumber(5)).toNumber()] = !_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("difl"), _40_v5);
        _nw4[(new BigNumber(6)).toNumber()] = true;
        _nw4[(new BigNumber(7)).toNumber()] = _35_v1.f15;
        _nw4[(new BigNumber(8)).toNumber()] = _35_v1.f15;
        _nw4[(new BigNumber(9)).toNumber()] = (_35_v1.f15) && ((_53_v10)[_module.__default.safeIndex(_module.__default.fm0(_34_v0, _39_v4, _35_v1.f15, globalState), new BigNumber((_53_v10).length))]);
        _nw4[(new BigNumber(10)).toNumber()] = (_54_v11).IsSubsetOf((_54_v11).update(_34_v0, _module.__default.abs(_39_v4)));
        _nw4[(new BigNumber(11)).toNumber()] = _34_v0;
        _nw4[(new BigNumber(12)).toNumber()] = _35_v1.f15;
        _nw4[(new BigNumber(13)).toNumber()] = _35_v1.f15;
        _nw4[(new BigNumber(14)).toNumber()] = (!(_34_v0)) || (true);
        _nw4[(new BigNumber(15)).toNumber()] = _module.__default.fm2(globalState);
        _nw4[(new BigNumber(16)).toNumber()] = _35_v1.f15;
        _nw4[(new BigNumber(17)).toNumber()] = _34_v0;
        _nw4[(new BigNumber(18)).toNumber()] = (_55_v12).IsDisjointFrom(_55_v12);
        _nw4[(new BigNumber(19)).toNumber()] = _34_v0;
        _nw4[(new BigNumber(20)).toNumber()] = (_56_v13).equals(_56_v13);
        _nw4[(new BigNumber(21)).toNumber()] = !(_35_v1.f15) || (_35_v1.f15);
        _nw4[(new BigNumber(22)).toNumber()] = _35_v1.f15;
        _nw4[(new BigNumber(23)).toNumber()] = _35_v1.f15;
        _nw4[(new BigNumber(24)).toNumber()] = _35_v1.f15;
        _nw4[(new BigNumber(25)).toNumber()] = _34_v0;
        _nw4[(new BigNumber(26)).toNumber()] = _35_v1.f15;
        _nw4[(new BigNumber(27)).toNumber()] = _35_v1.f15;
        _57_v14 = _nw4;
        let _index0 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_50_v8).length));
        (_50_v8)[_index0] = _34_v0;
        let _index1 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_50_v8).length));
        (_50_v8)[_index1] = ((!((new BigNumber(332)).isEqualTo(_39_v4))) ? (true) : ((_35_v1.f15) || (_34_v0)));
        let _index2 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_50_v8).length));
        (_50_v8)[_index2] = _module.__default.fm2(globalState);
        _39_v4 = _39_v4;
        let _62_v15;
        let _nw5 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
        _62_v15 = _nw5;
        let _index3 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_62_v15).length));
        (_62_v15)[_index3] = _39_v4;
        let _63_v16;
        _63_v16 = _dafny.Seq.of(_35_v1);
        let _64_v17;
        _64_v17 = _dafny.MultiSet.fromElements(_63_v16, _63_v16, _63_v16);
        let _65_v18;
        _65_v18 = _dafny.Seq.of(_63_v16);
        let _index4 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_62_v15).length));
        let _rhs0 = _62_v15;
        let _rhs1 = _39_v4;
        let _rhs2 = ((_64_v17).Intersect(_dafny.MultiSet.FromArray(_65_v18))).IsSubsetOf((_64_v17).Intersect(_64_v17));
        let _rhs3 = _39_v4;
        let _lhs0 = _62_v15;
        let _lhs1 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_62_v15).length));
        let _lhs2 = globalState;
        let _lhs3 = globalState;
        _62_v15 = _rhs0;
        _lhs0[_lhs1] = _rhs1;
        _lhs2.f8 = _rhs2;
        _lhs3.f0 = _rhs3;
      } else {
        let _66_v19;
        _66_v19 = _dafny.Map.Empty.slice().updateUnsafe(_34_v0,_module.__default.fm2(globalState));
        let _67_v20;
        _67_v20 = _dafny.Seq.of(_66_v19);
        let _68_v22;
        _68_v22 = _module.D1.create_DC2(_39_v4);
        (globalState).f1 = (new BigNumber((function () {
          let _coll5 = new _dafny.Set();
          for (const _compr_5 of (_67_v20).Elements) {
            let _69_v21 = _compr_5;
            if (_dafny.Seq.contains(_67_v20, _69_v21)) {
              _coll5.add(_69_v21);
            }
          }
          return _coll5;
        }()).length)).isLessThanOrEqualTo((_dafny.ZERO).minus(new BigNumber((_module.__default.fm6(_39_v4, _35_v1.f15, _68_v22, false, globalState)).length)));
        let _70_v23;
        _70_v23 = _dafny.MultiSet.fromElements(false);
        let _71_v24;
        _71_v24 = _dafny.Seq.UnicodeFromString("kv");
        (globalState).f10 = ((((_70_v23).contains(_34_v0)) ? ((_70_v23).get(_34_v0)) : (new BigNumber((_71_v24).length)))).isEqualTo(new BigNumber(717));
        let _72_v25;
        _72_v25 = _dafny.Map.Empty.slice().updateUnsafe(_40_v5,_35_v1);
        let _73_v26;
        _73_v26 = _dafny.Map.Empty.slice().updateUnsafe(_72_v25,_34_v0);
        let _74_v27;
        _74_v27 = _module.D5.create_DC13(_73_v26);
        _73_v26 = (_74_v27).dtor_cf20;
        _35_v1 = _35_v1;
        _66_v19 = (_66_v19).update(true, _35_v1.f15);
      }
      let _75_i4;
      _75_i4 = _dafny.ZERO;
      L0: {
        while (!(_34_v0)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_75_i4)) {
              break L0;
            }
            _75_i4 = (_75_i4).plus(_dafny.ONE);
            let _76_v28;
            let _nw6 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
            _76_v28 = _nw6;
            let _77_v29;
            _77_v29 = _dafny.Seq.UnicodeFromString("brhkdmqxe");
            let _index5 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_76_v28).length));
            (_76_v28)[_index5] = _77_v29;
            let _index6 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_76_v28).length));
            (_76_v28)[_index6] = _77_v29;
            let _78_v30;
            _78_v30 = _dafny.Map.Empty.slice().updateUnsafe(_34_v0,_40_v5);
            let _79_v31;
            let _nw7 = Array((new BigNumber(22)).toNumber());
            _nw7[(_dafny.ZERO).toNumber()] = _40_v5;
            _nw7[(_dafny.ONE).toNumber()] = _40_v5;
            _nw7[(new BigNumber(2)).toNumber()] = _40_v5;
            _nw7[(new BigNumber(3)).toNumber()] = _40_v5;
            _nw7[(new BigNumber(4)).toNumber()] = _module.__default.fm9(new BigNumber(398), (_76_v28)[_module.__default.safeIndex(new BigNumber(983), new BigNumber((_76_v28).length))], _39_v4, globalState);
            _nw7[(new BigNumber(5)).toNumber()] = _40_v5;
            _nw7[(new BigNumber(6)).toNumber()] = _40_v5;
            _nw7[(new BigNumber(7)).toNumber()] = _40_v5;
            _nw7[(new BigNumber(8)).toNumber()] = _40_v5;
            _nw7[(new BigNumber(9)).toNumber()] = _40_v5;
            _nw7[(new BigNumber(10)).toNumber()] = _40_v5;
            _nw7[(new BigNumber(11)).toNumber()] = _40_v5;
            _nw7[(new BigNumber(12)).toNumber()] = ((_34_v0) ? (new _dafny.CodePoint('s'.codePointAt(0))) : (new _dafny.CodePoint('l'.codePointAt(0))));
            _nw7[(new BigNumber(13)).toNumber()] = _40_v5;
            _nw7[(new BigNumber(14)).toNumber()] = _40_v5;
            _nw7[(new BigNumber(15)).toNumber()] = ((_34_v0) ? (new _dafny.CodePoint('v'.codePointAt(0))) : (_40_v5));
            _nw7[(new BigNumber(16)).toNumber()] = new _dafny.CodePoint('r'.codePointAt(0));
            _nw7[(new BigNumber(17)).toNumber()] = new _dafny.CodePoint('o'.codePointAt(0));
            _nw7[(new BigNumber(18)).toNumber()] = _40_v5;
            _nw7[(new BigNumber(19)).toNumber()] = (_77_v29)[_module.__default.safeIndex(_39_v4, new BigNumber((_77_v29).length))];
            _nw7[(new BigNumber(20)).toNumber()] = _40_v5;
            _nw7[(new BigNumber(21)).toNumber()] = (((_78_v30).contains(_34_v0)) ? ((_78_v30).get(_34_v0)) : (_40_v5));
            _79_v31 = _nw7;
            let _index7 = _module.__default.safeIndex(new BigNumber(452), new BigNumber((_79_v31).length));
            (_79_v31)[_index7] = _40_v5;
            let _index8 = _module.__default.safeIndex(new BigNumber(452), new BigNumber((_79_v31).length));
            (_79_v31)[_index8] = _40_v5;
            let _80_v32;
            let _nw8 = new _module.C0();
            _nw8.__ctor(_34_v0);
            _80_v32 = _nw8;
            let _index9 = _module.__default.safeIndex(new BigNumber(452), new BigNumber((_79_v31).length));
            (_79_v31)[_index9] = _40_v5;
          }
        }
      }
      if (_34_v0) {
        let _81_v33;
        let _init3 = ((_82_v1) => function (_83_i5) {
          return _82_v1.f15;
        })(_35_v1);
        let _nw9 = Array((new BigNumber(16)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw9.length); _i0_3++) {
          _nw9[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _81_v33 = _nw9;
        let _84_v34;
        _84_v34 = _dafny.Map.Empty.slice().updateUnsafe(_39_v4,_dafny.Seq.update(_dafny.Seq.of(_35_v1.f15, _34_v0, _35_v1.f15), _module.__default.safeIndex((_dafny.ZERO).minus(_39_v4), new BigNumber((_dafny.Seq.of(_35_v1.f15, _34_v0, _35_v1.f15)).length)), _34_v0));
        let _85_v35;
        _85_v35 = _dafny.Seq.of(_34_v0, _35_v1.f15);
        let _86_v36;
        _86_v36 = _dafny.Set.fromElements(_85_v35);
        let _index10 = _module.__default.safeIndex(new BigNumber(800), new BigNumber((_81_v33).length));
        (_81_v33)[_index10] = !(_86_v36).contains((((_84_v34).contains(new BigNumber((_dafny.MultiSet.fromElements(!(_35_v1.f15), _34_v0)).cardinality()))) ? ((_84_v34).get(new BigNumber((_dafny.MultiSet.fromElements(!(_35_v1.f15), _34_v0)).cardinality()))) : (_85_v35)));
        let _index11 = _module.__default.safeIndex(new BigNumber(800), new BigNumber((_81_v33).length));
        (_81_v33)[_index11] = _34_v0;
        let _87_v37;
        _87_v37 = _dafny.Seq.of(_39_v4);
        let _88_v38;
        _88_v38 = _dafny.MultiSet.fromElements(_87_v37, _87_v37);
        let _89_v39;
        _89_v39 = _dafny.Seq.of(_87_v37);
        let _90_v40;
        _90_v40 = _dafny.Map.Empty.slice().updateUnsafe(_35_v1.f15,_39_v4);
        let _91_v41;
        _91_v41 = _dafny.Map.Empty.slice().updateUnsafe(_39_v4,_39_v4);
        let _92_v42;
        let _nw10 = Array((new BigNumber(13)).toNumber());
        _nw10[(_dafny.ZERO).toNumber()] = _88_v38;
        _nw10[(_dafny.ONE).toNumber()] = _88_v38;
        _nw10[(new BigNumber(2)).toNumber()] = _dafny.MultiSet.FromArray(_89_v39);
        _nw10[(new BigNumber(3)).toNumber()] = _88_v38;
        _nw10[(new BigNumber(4)).toNumber()] = _88_v38;
        _nw10[(new BigNumber(5)).toNumber()] = _88_v38;
        _nw10[(new BigNumber(6)).toNumber()] = _dafny.MultiSet.FromArray(_89_v39);
        _nw10[(new BigNumber(7)).toNumber()] = _module.__default.fm11(_39_v4, _39_v4, _34_v0, _34_v0, globalState);
        _nw10[(new BigNumber(8)).toNumber()] = _dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber((_90_v40).length), new BigNumber((_91_v41).length)), _87_v37, _dafny.Seq.of(_39_v4), _dafny.Seq.update(_dafny.Seq.of(_39_v4), _module.__default.safeIndex(_39_v4, new BigNumber((_dafny.Seq.of(_39_v4)).length)), _module.__default.fm0((_81_v33)[_module.__default.safeIndex(new BigNumber(800), new BigNumber((_81_v33).length))], _39_v4, _35_v1.f15, globalState)), _87_v37);
        _nw10[(new BigNumber(9)).toNumber()] = (_88_v38).Difference(_88_v38);
        _nw10[(new BigNumber(10)).toNumber()] = _88_v38;
        _nw10[(new BigNumber(11)).toNumber()] = _dafny.MultiSet.fromElements(_87_v37);
        _nw10[(new BigNumber(12)).toNumber()] = ((_35_v1.f15) ? (_88_v38) : (_88_v38));
        _92_v42 = _nw10;
        let _index12 = _module.__default.safeIndex(new BigNumber(616), new BigNumber((_92_v42).length));
        (_92_v42)[_index12] = _module.__default.fm11(_39_v4, _39_v4, !(_35_v1.f15), !(false), globalState);
        let _index13 = _module.__default.safeIndex(new BigNumber(616), new BigNumber((_92_v42).length));
        (_92_v42)[_index13] = (_88_v38).Difference((_88_v38).Union(_88_v38));
        (globalState).f4 = _39_v4;
        (globalState).f1 = (_81_v33)[_module.__default.safeIndex(new BigNumber(800), new BigNumber((_81_v33).length))];
        (globalState).f0 = _39_v4;
      } else {
        let _93_v43;
        _93_v43 = _dafny.Seq.of(_39_v4, _39_v4);
        let _94_v44;
        _94_v44 = _dafny.Map.Empty.slice().updateUnsafe(_39_v4,_dafny.Map.Empty.slice().updateUnsafe(_40_v5,_93_v43));
        let _95_v45;
        _95_v45 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_93_v43).length),_39_v4);
        let _96_v46;
        _96_v46 = _dafny.Seq.of(_35_v1.f15);
        _94_v44 = (_94_v44).update(_39_v4, _dafny.Map.Empty.slice().updateUnsafe(_40_v5,_dafny.Seq.of(new BigNumber((_95_v45).length), (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_96_v46).length))), _39_v4, _39_v4, _39_v4)));
        let _97_v47;
        let _nw11 = Array((new BigNumber(25)).toNumber()).fill(false);
        _97_v47 = _nw11;
        let _index14 = _module.__default.safeIndex(new BigNumber(157), new BigNumber((_97_v47).length));
        (_97_v47)[_index14] = false;
        let _index15 = _module.__default.safeIndex(new BigNumber(157), new BigNumber((_97_v47).length));
        (_97_v47)[_index15] = !(true) || (false);
        let _98_v48;
        let _nw12 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
        _98_v48 = _nw12;
        let _index16 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_98_v48).length));
        (_98_v48)[_index16] = _module.__default.fm0((_97_v47)[_module.__default.safeIndex(new BigNumber(157), new BigNumber((_97_v47).length))], new BigNumber(604), (_97_v47)[_module.__default.safeIndex(new BigNumber(157), new BigNumber((_97_v47).length))], globalState);
        let _index17 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_98_v48).length));
        (_98_v48)[_index17] = _module.__default.safeModuloInt((_dafny.ZERO).minus(_39_v4), _39_v4);
        let _99_v49;
        _99_v49 = _dafny.Map.Empty.slice().updateUnsafe((_97_v47)[_module.__default.safeIndex(new BigNumber(157), new BigNumber((_97_v47).length))],_35_v1.f15);
        (globalState).f0 = new BigNumber((_dafny.Seq.update(_96_v46, _module.__default.safeIndex((((((_99_v49).contains(_34_v0)) ? ((_99_v49).get(_34_v0)) : (_35_v1.f15))) ? (_39_v4) : ((_dafny.ZERO).minus(_39_v4))), new BigNumber((_96_v46).length)), _35_v1.f15)).length);
        if (_34_v0) {
          let _rhs4 = _40_v5;
          let _rhs5 = _95_v45;
          _40_v5 = _rhs4;
          _95_v45 = _rhs5;
          let _100_v50;
          _100_v50 = _dafny.Set.fromElements(_39_v4);
          let _101_v51;
          _101_v51 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm12(new BigNumber((_100_v50).length), globalState),_dafny.Set.fromElements((_97_v47)[_module.__default.safeIndex(new BigNumber(157), new BigNumber((_97_v47).length))]));
          let _102_v52;
          _102_v52 = _dafny.Set.fromElements(true);
          _101_v51 = (_101_v51).update(_96_v46, _102_v52);
          let _103_v53;
          let _nw13 = new _module.C0();
          _nw13.__ctor((new BigNumber(367)).isEqualTo(new BigNumber(643)));
          _103_v53 = _nw13;
          _96_v46 = _dafny.Seq.Concat(_96_v46, _96_v46);
          (globalState).f11 = _97_v47;
        } else {
          let _104_v54;
          _104_v54 = _dafny.Map.Empty.slice().updateUnsafe((_98_v48)[_module.__default.safeIndex(new BigNumber(259), new BigNumber((_98_v48).length))],_35_v1);
          let _105_v55;
          _105_v55 = _dafny.Set.fromElements(_35_v1.f15, !(_35_v1.f15));
          let _106_v56;
          _106_v56 = _dafny.Map.Empty.slice().updateUnsafe(_105_v55,(_98_v48)[_module.__default.safeIndex(new BigNumber(259), new BigNumber((_98_v48).length))]);
          let _107_v57;
          _107_v57 = _dafny.Set.fromElements(_35_v1, (((_104_v54).contains(new BigNumber((_106_v56).length))) ? ((_104_v54).get(new BigNumber((_106_v56).length))) : (_35_v1)), _35_v1);
          let _108_v58;
          _108_v58 = _dafny.MultiSet.fromElements((_97_v47)[_module.__default.safeIndex(new BigNumber(157), new BigNumber((_97_v47).length))], (_97_v47)[_module.__default.safeIndex(new BigNumber(157), new BigNumber((_97_v47).length))]);
          let _rhs6 = (_dafny.ZERO).minus((_98_v48)[_module.__default.safeIndex(new BigNumber(259), new BigNumber((_98_v48).length))]);
          let _rhs7 = (((_104_v54).contains(_module.__default.safeDivisionInt(_39_v4, new BigNumber((_108_v58).cardinality())))) ? ((_104_v54).get(_module.__default.safeDivisionInt(_39_v4, new BigNumber((_108_v58).cardinality())))) : (_35_v1));
          let _rhs8 = _39_v4;
          let _rhs9 = ((_107_v57).Union(_107_v57)).Union(_107_v57);
          let _lhs4 = globalState;
          let _lhs5 = globalState;
          _lhs4.f4 = _rhs6;
          _35_v1 = _rhs7;
          _lhs5.f6 = _rhs8;
          _107_v57 = _rhs9;
          let _109_v59;
          _109_v59 = _dafny.Map.Empty.slice().updateUnsafe(_97_v47,(_98_v48)[_module.__default.safeIndex(new BigNumber(259), new BigNumber((_98_v48).length))]);
          _109_v59 = (_109_v59).update(_97_v47, new BigNumber(282));
          let _110_v60;
          _110_v60 = _dafny.Map.Empty.slice().updateUnsafe((_39_v4).isLessThanOrEqualTo(_39_v4),_97_v47);
          _110_v60 = (_110_v60).update(_module.__default.fm2(globalState), _97_v47);
          (globalState).f6 = _module.__default.fm0(_35_v1.f15, (_93_v43)[_module.__default.safeIndex(new BigNumber(-243), new BigNumber((_93_v43).length))], _35_v1.f15, globalState);
          let _111_v61;
          _111_v61 = _dafny.Seq.UnicodeFromString("mkynuw");
          _95_v45 = (_95_v45).update((_98_v48)[_module.__default.safeIndex(new BigNumber(259), new BigNumber((_98_v48).length))], new BigNumber((_dafny.Seq.Concat(_111_v61, _111_v61)).length));
        }
      }
      if (_35_v1.f15) {
        let _112_v62;
        let _init4 = ((_113_v4) => function (_114_i6) {
          return (_114_i6).minus(_113_v4);
        })(_39_v4);
        let _nw14 = Array((new BigNumber(10)).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw14.length); _i0_4++) {
          _nw14[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _112_v62 = _nw14;
        let _115_v63;
        _115_v63 = _dafny.MultiSet.fromElements(_40_v5);
        let _116_v64;
        _116_v64 = _dafny.Set.fromElements(_39_v4);
        let _117_v65;
        _117_v65 = _dafny.Seq.UnicodeFromString("f");
        let _118_v66;
        _118_v66 = _module.D1.create_DC2(_39_v4);
        let _119_v67;
        _119_v67 = _dafny.Map.Empty.slice().updateUnsafe(_34_v0,_35_v1.f15);
        let _120_v68;
        let _nw15 = Array((new BigNumber(25)).toNumber());
        _nw15[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_117_v65, _117_v65);
        _nw15[(_dafny.ONE).toNumber()] = _module.__default.fm6(_39_v4, _35_v1.f15, _118_v66, true, globalState);
        _nw15[(new BigNumber(2)).toNumber()] = _dafny.Seq.UnicodeFromString("kov");
        _nw15[(new BigNumber(3)).toNumber()] = _117_v65;
        _nw15[(new BigNumber(4)).toNumber()] = _117_v65;
        _nw15[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("ndxuxbq");
        _nw15[(new BigNumber(6)).toNumber()] = _117_v65;
        _nw15[(new BigNumber(7)).toNumber()] = _117_v65;
        _nw15[(new BigNumber(8)).toNumber()] = _117_v65;
        _nw15[(new BigNumber(9)).toNumber()] = _117_v65;
        _nw15[(new BigNumber(10)).toNumber()] = _117_v65;
        _nw15[(new BigNumber(11)).toNumber()] = _module.__default.fm6(_39_v4, _35_v1.f15, _module.D1.create_DC2(new BigNumber((_117_v65).length)), _35_v1.f15, globalState);
        _nw15[(new BigNumber(12)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(372)), ((_121_v5) => function (_122_i7) {
          return _121_v5;
        })(_40_v5));
        _nw15[(new BigNumber(13)).toNumber()] = _117_v65;
        _nw15[(new BigNumber(14)).toNumber()] = _117_v65;
        _nw15[(new BigNumber(15)).toNumber()] = _117_v65;
        _nw15[(new BigNumber(16)).toNumber()] = ((!(_35_v1.f15)) ? (_117_v65) : (_117_v65));
        _nw15[(new BigNumber(17)).toNumber()] = _117_v65;
        _nw15[(new BigNumber(18)).toNumber()] = _117_v65;
        _nw15[(new BigNumber(19)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(123)), ((_123_v5) => function (_124_i8) {
          return _123_v5;
        })(_40_v5)), _module.__default.safeIndex(_39_v4, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(123)), ((_125_v5) => function (_126_i8) {
          return _125_v5;
        })(_40_v5))).length)), _40_v5);
        _nw15[(new BigNumber(20)).toNumber()] = _117_v65;
        _nw15[(new BigNumber(21)).toNumber()] = _117_v65;
        _nw15[(new BigNumber(22)).toNumber()] = _module.__default.fm6(_39_v4, _module.__default.fm2(globalState), _module.D1.create_DC2(new BigNumber((_119_v67).length)), _35_v1.f15, globalState);
        _nw15[(new BigNumber(23)).toNumber()] = _module.__default.fm6(_39_v4, true, _118_v66, _35_v1.f15, globalState);
        _nw15[(new BigNumber(24)).toNumber()] = _117_v65;
        _120_v68 = _nw15;
        let _index18 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_120_v68).length));
        (_120_v68)[_index18] = _dafny.Seq.UnicodeFromString("ju");
        let _index19 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_112_v62).length));
        (_112_v62)[_index19] = _39_v4;
        let _index20 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_120_v68).length));
        let _index21 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_112_v62).length));
        let _rhs10 = _112_v62;
        let _rhs11 = _115_v63;
        let _rhs12 = (_dafny.Set.fromElements(_39_v4)).Union(_116_v64);
        let _rhs13 = _117_v65;
        let _rhs14 = (_39_v4).plus(_39_v4);
        let _lhs6 = _120_v68;
        let _lhs7 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_120_v68).length));
        let _lhs8 = _112_v62;
        let _lhs9 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_112_v62).length));
        _112_v62 = _rhs10;
        _115_v63 = _rhs11;
        _116_v64 = _rhs12;
        _lhs6[_lhs7] = _rhs13;
        _lhs8[_lhs9] = _rhs14;
        let _127_v69;
        let _nw16 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.of());
        _127_v69 = _nw16;
        let _128_v70;
        _128_v70 = _dafny.Seq.of((_112_v62)[_module.__default.safeIndex(new BigNumber(132), new BigNumber((_112_v62).length))], _module.__default.fm0(_35_v1.f15, _39_v4, !(_34_v0), globalState), _39_v4, (_112_v62)[_module.__default.safeIndex(new BigNumber(132), new BigNumber((_112_v62).length))], (_112_v62)[_module.__default.safeIndex(new BigNumber(132), new BigNumber((_112_v62).length))]);
        let _index22 = _module.__default.safeIndex(new BigNumber(550), new BigNumber((_127_v69).length));
        (_127_v69)[_index22] = _128_v70;
        let _index23 = _module.__default.safeIndex(new BigNumber(550), new BigNumber((_127_v69).length));
        (_127_v69)[_index23] = _128_v70;
        let _129_v71;
        let _nw17 = Array((new BigNumber(6)).toNumber()).fill(false);
        _129_v71 = _nw17;
        let _index24 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_129_v71).length));
        (_129_v71)[_index24] = !(_35_v1.f15);
        let _index25 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_129_v71).length));
        (_129_v71)[_index25] = _34_v0;
        let _130_v72;
        _130_v72 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_120_v68)[_module.__default.safeIndex(new BigNumber(985), new BigNumber((_120_v68).length))]).length),_35_v1.f15);
        let _131_v73;
        _131_v73 = _dafny.Seq.of(_112_v62, _112_v62);
        let _132_v74;
        _132_v74 = _module.D6.create_DC15(_112_v62);
        let _133_v75;
        let _nw18 = Array((new BigNumber(16)).toNumber());
        _nw18[(_dafny.ZERO).toNumber()] = _112_v62;
        _nw18[(_dafny.ONE).toNumber()] = _112_v62;
        _nw18[(new BigNumber(2)).toNumber()] = _112_v62;
        _nw18[(new BigNumber(3)).toNumber()] = (((((_130_v72).contains(_39_v4)) ? ((_130_v72).get(_39_v4)) : (_35_v1.f15))) ? (_112_v62) : ((_131_v73)[_module.__default.safeIndex((_112_v62)[_module.__default.safeIndex(new BigNumber(132), new BigNumber((_112_v62).length))], new BigNumber((_131_v73).length))]));
        _nw18[(new BigNumber(4)).toNumber()] = _112_v62;
        _nw18[(new BigNumber(5)).toNumber()] = _112_v62;
        _nw18[(new BigNumber(6)).toNumber()] = _112_v62;
        _nw18[(new BigNumber(7)).toNumber()] = _112_v62;
        _nw18[(new BigNumber(8)).toNumber()] = _112_v62;
        _nw18[(new BigNumber(9)).toNumber()] = (((_129_v71)[_module.__default.safeIndex(new BigNumber(460), new BigNumber((_129_v71).length))]) ? (_112_v62) : ((_132_v74).dtor_cf21));
        _nw18[(new BigNumber(10)).toNumber()] = _112_v62;
        _nw18[(new BigNumber(11)).toNumber()] = _112_v62;
        _nw18[(new BigNumber(12)).toNumber()] = _112_v62;
        _nw18[(new BigNumber(13)).toNumber()] = _112_v62;
        _nw18[(new BigNumber(14)).toNumber()] = _112_v62;
        _nw18[(new BigNumber(15)).toNumber()] = _112_v62;
        _133_v75 = _nw18;
        let _index26 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_133_v75).length));
        (_133_v75)[_index26] = _112_v62;
        let _134_v77;
        let _init5 = ((_135_v1, _136_v62, _137_v5) => function (_138_i9) {
          return function () {
            let _coll6 = new _dafny.Set();
            for (const _compr_6 of (_dafny.MultiSet.fromElements(_dafny.Seq.of(_135_v1.f15, (_module.D4.create_DC10((_136_v62)[_module.__default.safeIndex(new BigNumber(132), new BigNumber((_136_v62).length))], _137_v5, (_136_v62)[_module.__default.safeIndex(new BigNumber(132), new BigNumber((_136_v62).length))], _135_v1.f15, _dafny.Seq.of(_module.D3.create_DC6((_136_v62)[_module.__default.safeIndex(new BigNumber(132), new BigNumber((_136_v62).length))])))).dtor_cf16))).Elements) {
              let _139_v76 = _compr_6;
              if ((_dafny.MultiSet.fromElements(_dafny.Seq.of(_135_v1.f15, (_module.D4.create_DC10((_136_v62)[_module.__default.safeIndex(new BigNumber(132), new BigNumber((_136_v62).length))], _137_v5, (_136_v62)[_module.__default.safeIndex(new BigNumber(132), new BigNumber((_136_v62).length))], _135_v1.f15, _dafny.Seq.of(_module.D3.create_DC6((_136_v62)[_module.__default.safeIndex(new BigNumber(132), new BigNumber((_136_v62).length))])))).dtor_cf16))).contains(_139_v76)) {
                _coll6.add(_139_v76);
              }
            }
            return _coll6;
          }();
        })(_35_v1, _112_v62, _40_v5);
        let _nw19 = Array((_dafny.ONE).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw19.length); _i0_5++) {
          _nw19[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _134_v77 = _nw19;
        let _140_v78;
        _140_v78 = _dafny.Seq.of(true, _35_v1.f15, (_129_v71)[_module.__default.safeIndex(new BigNumber(460), new BigNumber((_129_v71).length))]);
        let _141_v79;
        _141_v79 = _dafny.Set.fromElements(_140_v78, _140_v78, _140_v78);
        let _index27 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_134_v77).length));
        (_134_v77)[_index27] = _141_v79;
        let _index28 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_133_v75).length));
        let _index29 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_134_v77).length));
        let _index30 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_129_v71).length));
        let _rhs15 = (_132_v74).dtor_cf21;
        let _rhs16 = _35_v1.f15;
        let _rhs17 = ((_35_v1.f15) ? ((_112_v62)[_module.__default.safeIndex(new BigNumber(132), new BigNumber((_112_v62).length))]) : ((_112_v62)[_module.__default.safeIndex(new BigNumber(132), new BigNumber((_112_v62).length))]));
        let _rhs18 = (_module.__default.fm13((_112_v62)[_module.__default.safeIndex(new BigNumber(132), new BigNumber((_112_v62).length))], globalState)).Difference(_141_v79);
        let _rhs19 = (_129_v71)[_module.__default.safeIndex(new BigNumber(460), new BigNumber((_129_v71).length))];
        let _lhs10 = _133_v75;
        let _lhs11 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_133_v75).length));
        let _lhs12 = globalState;
        let _lhs13 = globalState;
        let _lhs14 = _134_v77;
        let _lhs15 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_134_v77).length));
        let _lhs16 = _129_v71;
        let _lhs17 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_129_v71).length));
        _lhs10[_lhs11] = _rhs15;
        _lhs12.f10 = _rhs16;
        _lhs13.f0 = _rhs17;
        _lhs14[_lhs15] = _rhs18;
        _lhs16[_lhs17] = _rhs19;
        (globalState).f10 = (_129_v71)[_module.__default.safeIndex(new BigNumber(460), new BigNumber((_129_v71).length))];
      } else {
        (_35_v1).f15 = _35_v1.f15;
        if (_35_v1.f15) {
          let _142_v80;
          let _nw20 = Array((_dafny.ONE).toNumber());
          _nw20[(_dafny.ZERO).toNumber()] = _39_v4;
          _142_v80 = _nw20;
          let _143_v81;
          _143_v81 = _dafny.Map.Empty.slice().updateUnsafe(_39_v4,_142_v80);
          let _144_v82;
          let _nw21 = Array((new BigNumber(25)).toNumber());
          _nw21[(_dafny.ZERO).toNumber()] = _35_v1.f15;
          _nw21[(_dafny.ONE).toNumber()] = _35_v1.f15;
          _nw21[(new BigNumber(2)).toNumber()] = _35_v1.f15;
          _nw21[(new BigNumber(3)).toNumber()] = false;
          _nw21[(new BigNumber(4)).toNumber()] = _35_v1.f15;
          _nw21[(new BigNumber(5)).toNumber()] = _35_v1.f15;
          _nw21[(new BigNumber(6)).toNumber()] = _35_v1.f15;
          _nw21[(new BigNumber(7)).toNumber()] = _35_v1.f15;
          _nw21[(new BigNumber(8)).toNumber()] = _35_v1.f15;
          _nw21[(new BigNumber(9)).toNumber()] = _34_v0;
          _nw21[(new BigNumber(10)).toNumber()] = _35_v1.f15;
          _nw21[(new BigNumber(11)).toNumber()] = _34_v0;
          _nw21[(new BigNumber(12)).toNumber()] = _35_v1.f15;
          _nw21[(new BigNumber(13)).toNumber()] = _35_v1.f15;
          _nw21[(new BigNumber(14)).toNumber()] = true;
          _nw21[(new BigNumber(15)).toNumber()] = _34_v0;
          _nw21[(new BigNumber(16)).toNumber()] = _35_v1.f15;
          _nw21[(new BigNumber(17)).toNumber()] = _34_v0;
          _nw21[(new BigNumber(18)).toNumber()] = _35_v1.f15;
          _nw21[(new BigNumber(19)).toNumber()] = _34_v0;
          _nw21[(new BigNumber(20)).toNumber()] = _35_v1.f15;
          _nw21[(new BigNumber(21)).toNumber()] = _35_v1.f15;
          _nw21[(new BigNumber(22)).toNumber()] = _34_v0;
          _nw21[(new BigNumber(23)).toNumber()] = _34_v0;
          _nw21[(new BigNumber(24)).toNumber()] = false;
          _144_v82 = _nw21;
          let _145_v83;
          _145_v83 = _144_v82;
          let _146_v84;
          let _nw22 = Array((new BigNumber(18)).toNumber());
          _nw22[(_dafny.ZERO).toNumber()] = _142_v80;
          _nw22[(_dafny.ONE).toNumber()] = _142_v80;
          _nw22[(new BigNumber(2)).toNumber()] = (((_143_v81).contains(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_145_v83,_35_v1.f15)).length))) ? ((_143_v81).get(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_145_v83,_35_v1.f15)).length))) : (_142_v80));
          _nw22[(new BigNumber(3)).toNumber()] = _142_v80;
          _nw22[(new BigNumber(4)).toNumber()] = _142_v80;
          _nw22[(new BigNumber(5)).toNumber()] = _142_v80;
          _nw22[(new BigNumber(6)).toNumber()] = _142_v80;
          _nw22[(new BigNumber(7)).toNumber()] = _142_v80;
          _nw22[(new BigNumber(8)).toNumber()] = _142_v80;
          _nw22[(new BigNumber(9)).toNumber()] = _142_v80;
          _nw22[(new BigNumber(10)).toNumber()] = _142_v80;
          _nw22[(new BigNumber(11)).toNumber()] = _142_v80;
          _nw22[(new BigNumber(12)).toNumber()] = _142_v80;
          _nw22[(new BigNumber(13)).toNumber()] = _142_v80;
          _nw22[(new BigNumber(14)).toNumber()] = _142_v80;
          _nw22[(new BigNumber(15)).toNumber()] = _142_v80;
          _nw22[(new BigNumber(16)).toNumber()] = _142_v80;
          _nw22[(new BigNumber(17)).toNumber()] = _142_v80;
          _146_v84 = _nw22;
          let _147_v85;
          _147_v85 = _dafny.Seq.of(_146_v84);
          let _148_v86;
          _148_v86 = _dafny.Seq.of(_35_v1.f15, _35_v1.f15);
          let _149_v87;
          _149_v87 = _dafny.Map.Empty.slice().updateUnsafe((_147_v85)[_module.__default.safeIndex(_39_v4, new BigNumber((_147_v85).length))],_dafny.Seq.Concat(_148_v86, _148_v86));
          _149_v87 = (_149_v87).update(_146_v84, _148_v86);
          _34_v0 = (_39_v4).isEqualTo(_39_v4);
          let _150_v88;
          _150_v88 = _dafny.Map.Empty.slice().updateUnsafe(_35_v1.f15,_35_v1.f15);
          _150_v88 = (_150_v88).update(_35_v1.f15, _35_v1.f15);
          let _151_v89;
          _151_v89 = _module.D3.create_DC5(_148_v86);
          (globalState).f0 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_148_v86,_module.D3.create_DC7(_151_v89))).length);
          let _index31 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_144_v82).length));
          (_144_v82)[_index31] = _35_v1.f15;
          let _152_v90;
          _152_v90 = _module.D1.create_DC2(_39_v4);
          let _153_v91;
          _153_v91 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(249),_39_v4);
          let _154_v92;
          let _nw23 = Array((new BigNumber(2)).toNumber());
          _nw23[(_dafny.ZERO).toNumber()] = _152_v90;
          _nw23[(_dafny.ONE).toNumber()] = _module.D1.create_DC2(new BigNumber((_153_v91).length));
          _154_v92 = _nw23;
          let _155_v93;
          _155_v93 = _dafny.MultiSet.fromElements(_154_v92);
          let _index32 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_144_v82).length));
          let _rhs20 = _39_v4;
          let _rhs21 = _35_v1.f15;
          let _rhs22 = (((true) ? (_dafny.MultiSet.fromElements(_154_v92)) : (_155_v93))).IsDisjointFrom(_155_v93);
          let _rhs23 = _142_v80;
          let _lhs18 = globalState;
          let _lhs19 = _144_v82;
          let _lhs20 = _module.__default.safeIndex(new BigNumber(630), new BigNumber((_144_v82).length));
          let _lhs21 = _35_v1;
          _lhs18.f4 = _rhs20;
          _lhs19[_lhs20] = _rhs21;
          _lhs21.f15 = _rhs22;
          _142_v80 = _rhs23;
        } else {
          let _156_v94;
          _156_v94 = _dafny.Map.Empty.slice().updateUnsafe(_35_v1.f15,_35_v1);
          let _157_v95;
          let _nw24 = Array((new BigNumber(14)).toNumber());
          _nw24[(_dafny.ZERO).toNumber()] = (((_156_v94).contains(!(true))) ? ((_156_v94).get(!(true))) : (_35_v1));
          _nw24[(_dafny.ONE).toNumber()] = _35_v1;
          _nw24[(new BigNumber(2)).toNumber()] = _35_v1;
          _nw24[(new BigNumber(3)).toNumber()] = _35_v1;
          _nw24[(new BigNumber(4)).toNumber()] = _35_v1;
          _nw24[(new BigNumber(5)).toNumber()] = _35_v1;
          _nw24[(new BigNumber(6)).toNumber()] = _35_v1;
          _nw24[(new BigNumber(7)).toNumber()] = _35_v1;
          _nw24[(new BigNumber(8)).toNumber()] = _35_v1;
          _nw24[(new BigNumber(9)).toNumber()] = (((_156_v94).contains(_35_v1.f15)) ? ((_156_v94).get(_35_v1.f15)) : (_35_v1));
          _nw24[(new BigNumber(10)).toNumber()] = _35_v1;
          _nw24[(new BigNumber(11)).toNumber()] = _35_v1;
          _nw24[(new BigNumber(12)).toNumber()] = _35_v1;
          _nw24[(new BigNumber(13)).toNumber()] = _35_v1;
          _157_v95 = _nw24;
          let _158_v96;
          _158_v96 = _dafny.Seq.of(_35_v1, _35_v1);
          let _index33 = _module.__default.safeIndex(new BigNumber(101), new BigNumber((_157_v95).length));
          (_157_v95)[_index33] = (_158_v96)[_module.__default.safeIndex(_39_v4, new BigNumber((_158_v96).length))];
          let _159_v97;
          _159_v97 = _dafny.Seq.UnicodeFromString("iaufjsuky");
          let _160_v98;
          _160_v98 = _dafny.Seq.of(_159_v97, _159_v97);
          let _161_v99;
          _161_v99 = _module.D2.create_DC3(_159_v97);
          let _162_v100;
          let _nw25 = Array((new BigNumber(5)).toNumber());
          _nw25[(_dafny.ZERO).toNumber()] = (_160_v98)[_module.__default.safeIndex(new BigNumber((_159_v97).length), new BigNumber((_160_v98).length))];
          _nw25[(_dafny.ONE).toNumber()] = _159_v97;
          _nw25[(new BigNumber(2)).toNumber()] = _159_v97;
          _nw25[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(214)), ((_163_v5) => function (_164_i10) {
            return _163_v5;
          })(_40_v5)), _159_v97), _module.__default.safeIndex(_39_v4, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(214)), ((_165_v5) => function (_166_i10) {
            return _165_v5;
          })(_40_v5)), _159_v97)).length)), _40_v5);
          _nw25[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat((_161_v99).dtor_cf3, _dafny.Seq.UnicodeFromString("rtb"));
          _162_v100 = _nw25;
          let _index34 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_162_v100).length));
          (_162_v100)[_index34] = _159_v97;
          let _index35 = _module.__default.safeIndex(new BigNumber(101), new BigNumber((_157_v95).length));
          let _index36 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_162_v100).length));
          let _rhs24 = _35_v1.f15;
          let _rhs25 = (_module.__default.fm0(_35_v1.f15, _39_v4, _35_v1.f15, globalState)).minus((_dafny.ZERO).minus(_39_v4));
          let _rhs26 = _35_v1;
          let _rhs27 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("mjfq"), _159_v97);
          let _lhs22 = _35_v1;
          let _lhs23 = globalState;
          let _lhs24 = _157_v95;
          let _lhs25 = _module.__default.safeIndex(new BigNumber(101), new BigNumber((_157_v95).length));
          let _lhs26 = _162_v100;
          let _lhs27 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_162_v100).length));
          _lhs22.f15 = _rhs24;
          _lhs23.f4 = _rhs25;
          _lhs24[_lhs25] = _rhs26;
          _lhs26[_lhs27] = _rhs27;
          let _167_v101;
          _167_v101 = _dafny.Seq.of(_35_v1.f15, _35_v1.f15, _module.__default.fm2(globalState));
          _167_v101 = _167_v101;
          let _rhs28 = ((_35_v1.f15) || (_34_v0)) || (_35_v1.f15);
          let _rhs29 = ((_35_v1.f15) ? (_module.__default.safeDivisionInt(new BigNumber(((_162_v100)[_module.__default.safeIndex(new BigNumber(222), new BigNumber((_162_v100).length))]).length), _39_v4)) : (_39_v4));
          let _lhs28 = globalState;
          _lhs28.f1 = _rhs28;
          _39_v4 = _rhs29;
          let _168_v102;
          _168_v102 = _dafny.Seq.of(_39_v4);
          let _169_v103;
          _169_v103 = _dafny.Set.fromElements(!(false) || (false), _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-681)), ((_170_v4) => function (_171_i11) {
            return _170_v4;
          })(_39_v4)), _168_v102));
          let _172_v104;
          _172_v104 = _dafny.Set.fromElements(_39_v4, new BigNumber(741), _module.__default.safeDivisionInt(_39_v4, new BigNumber(((_162_v100)[_module.__default.safeIndex(new BigNumber(222), new BigNumber((_162_v100).length))]).length)));
          let _rhs30 = (_dafny.ZERO).minus(new BigNumber((_169_v103).length));
          let _rhs31 = (_dafny.ZERO).minus(_39_v4);
          let _rhs32 = new BigNumber((_172_v104).length);
          let _lhs29 = globalState;
          let _lhs30 = globalState;
          let _lhs31 = globalState;
          _lhs29.f6 = _rhs30;
          _lhs30.f0 = _rhs31;
          _lhs31.f0 = _rhs32;
          let _173_v105;
          let _init6 = ((_174_v4) => function (_175_i12) {
            return _dafny.Map.Empty.slice().updateUnsafe(_174_v4,_174_v4);
          })(_39_v4);
          let _nw26 = Array((new BigNumber(15)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw26.length); _i0_6++) {
            _nw26[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _173_v105 = _nw26;
          let _176_v106;
          _176_v106 = _dafny.Map.Empty.slice().updateUnsafe(_35_v1.f15,_35_v1.f15);
          let _177_v107;
          _177_v107 = _module.D1.create_DC2(_39_v4);
          let _rhs33 = _173_v105;
          let _rhs34 = (_module.__default.fm0(_35_v1.f15, (_dafny.ZERO).minus((_dafny.ZERO).minus(_39_v4)), _35_v1.f15, globalState)).multipliedBy((_dafny.ZERO).minus(new BigNumber((_module.__default.fm6(new BigNumber((_176_v106).length), _35_v1.f15, _177_v107, _35_v1.f15, globalState)).length)));
          let _rhs35 = _39_v4;
          let _rhs36 = _39_v4;
          let _lhs32 = globalState;
          let _lhs33 = globalState;
          let _lhs34 = globalState;
          _173_v105 = _rhs33;
          _lhs32.f0 = _rhs34;
          _lhs33.f6 = _rhs35;
          _lhs34.f4 = _rhs36;
        }
        let _178_v108;
        let _nw27 = Array((new BigNumber(25)).toNumber()).fill(false);
        _178_v108 = _nw27;
        let _179_v109;
        _179_v109 = _178_v108;
        let _source2 = _179_v109;
        let _180___mcc_h3 = _source2;
        let _181_cf0 = _180___mcc_h3;
        let _182_v110;
        let _nw28 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _182_v110 = _nw28;
        let _183_v111;
        _183_v111 = _dafny.Seq.UnicodeFromString("mftvud");
        let _index37 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_182_v110).length));
        (_182_v110)[_index37] = ((_35_v1.f15) ? (_183_v111) : (_183_v111));
        let _index38 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_181_cf0).length));
        (_181_cf0)[_index38] = _35_v1.f15;
        let _index39 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_182_v110).length));
        let _index40 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_181_cf0).length));
        let _rhs37 = _183_v111;
        let _rhs38 = _39_v4;
        let _rhs39 = _module.__default.safeDivisionInt(_39_v4, _module.__default.safeDivisionInt(_39_v4, _39_v4));
        let _rhs40 = !(!(_35_v1.f15));
        let _lhs35 = _182_v110;
        let _lhs36 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_182_v110).length));
        let _lhs37 = globalState;
        let _lhs38 = _181_cf0;
        let _lhs39 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_181_cf0).length));
        _lhs35[_lhs36] = _rhs37;
        _39_v4 = _rhs38;
        _lhs37.f0 = _rhs39;
        _lhs38[_lhs39] = _rhs40;
        let _184_v112;
        _184_v112 = _dafny.Map.Empty.slice().updateUnsafe(((!(_34_v0)) ? (_182_v110) : (_182_v110)),(_dafny.ZERO).minus(_39_v4));
        _184_v112 = (_184_v112).update(_182_v110, _module.__default.fm0(false, _39_v4, _34_v0, globalState));
        let _185_v113;
        let _nw29 = new _module.C0();
        _nw29.__ctor(!(!(_35_v1.f15) || (_35_v1.f15)));
        _185_v113 = _nw29;
        let _186_v114;
        let _nw30 = Array((new BigNumber(4)).toNumber()).fill([]);
        _186_v114 = _nw30;
        let _187_v115;
        let _nw31 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
        _187_v115 = _nw31;
        let _index41 = _module.__default.safeIndex(new BigNumber(611), new BigNumber((_186_v114).length));
        (_186_v114)[_index41] = _187_v115;
        let _index42 = _module.__default.safeIndex(new BigNumber(611), new BigNumber((_186_v114).length));
        (_186_v114)[_index42] = _187_v115;
        if (_35_v1.f15) {
          _34_v0 = !(_35_v1.f15);
          (globalState).f4 = _39_v4;
          let _188_v116;
          _188_v116 = _dafny.Map.Empty.slice().updateUnsafe(_35_v1.f15,_39_v4);
          let _189_v117;
          _189_v117 = _dafny.Seq.UnicodeFromString("shhs");
          let _190_v118;
          _190_v118 = _module.D1.create_DC1(_39_v4);
          let _191_v119;
          _191_v119 = _dafny.Set.fromElements(_35_v1.f15, _35_v1.f15);
          let _192_v120;
          let _nw32 = Array((new BigNumber(23)).toNumber());
          _nw32[(_dafny.ZERO).toNumber()] = new BigNumber(115);
          _nw32[(_dafny.ONE).toNumber()] = new BigNumber((_188_v116).length);
          _nw32[(new BigNumber(2)).toNumber()] = new BigNumber((_189_v117).length);
          _nw32[(new BigNumber(3)).toNumber()] = _39_v4;
          _nw32[(new BigNumber(4)).toNumber()] = new BigNumber((_189_v117).length);
          _nw32[(new BigNumber(5)).toNumber()] = (_190_v118).dtor_cf1;
          _nw32[(new BigNumber(6)).toNumber()] = new BigNumber((_189_v117).length);
          _nw32[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(_39_v4);
          _nw32[(new BigNumber(8)).toNumber()] = _39_v4;
          _nw32[(new BigNumber(9)).toNumber()] = _39_v4;
          _nw32[(new BigNumber(10)).toNumber()] = _39_v4;
          _nw32[(new BigNumber(11)).toNumber()] = _39_v4;
          _nw32[(new BigNumber(12)).toNumber()] = new BigNumber((_module.__default.fm4(_35_v1.f15, _39_v4, globalState)).length);
          _nw32[(new BigNumber(13)).toNumber()] = new BigNumber((_191_v119).length);
          _nw32[(new BigNumber(14)).toNumber()] = _39_v4;
          _nw32[(new BigNumber(15)).toNumber()] = _39_v4;
          _nw32[(new BigNumber(16)).toNumber()] = _module.__default.fm0(_35_v1.f15, _39_v4, _35_v1.f15, globalState);
          _nw32[(new BigNumber(17)).toNumber()] = _39_v4;
          _nw32[(new BigNumber(18)).toNumber()] = _39_v4;
          _nw32[(new BigNumber(19)).toNumber()] = (_dafny.ZERO).minus(_39_v4);
          _nw32[(new BigNumber(20)).toNumber()] = _39_v4;
          _nw32[(new BigNumber(21)).toNumber()] = _39_v4;
          _nw32[(new BigNumber(22)).toNumber()] = _39_v4;
          _192_v120 = _nw32;
          let _193_v121;
          let _nw33 = Array((new BigNumber(17)).toNumber());
          _nw33[(_dafny.ZERO).toNumber()] = _192_v120;
          _nw33[(_dafny.ONE).toNumber()] = _192_v120;
          _nw33[(new BigNumber(2)).toNumber()] = _192_v120;
          _nw33[(new BigNumber(3)).toNumber()] = _192_v120;
          _nw33[(new BigNumber(4)).toNumber()] = _192_v120;
          _nw33[(new BigNumber(5)).toNumber()] = _192_v120;
          _nw33[(new BigNumber(6)).toNumber()] = _192_v120;
          _nw33[(new BigNumber(7)).toNumber()] = _192_v120;
          _nw33[(new BigNumber(8)).toNumber()] = _192_v120;
          _nw33[(new BigNumber(9)).toNumber()] = _192_v120;
          _nw33[(new BigNumber(10)).toNumber()] = _192_v120;
          _nw33[(new BigNumber(11)).toNumber()] = _192_v120;
          _nw33[(new BigNumber(12)).toNumber()] = _192_v120;
          _nw33[(new BigNumber(13)).toNumber()] = _192_v120;
          _nw33[(new BigNumber(14)).toNumber()] = _192_v120;
          _nw33[(new BigNumber(15)).toNumber()] = _192_v120;
          _nw33[(new BigNumber(16)).toNumber()] = _192_v120;
          _193_v121 = _nw33;
          let _194_v122;
          _194_v122 = _dafny.Seq.of(_193_v121, _193_v121, _193_v121);
          let _195_v123;
          _195_v123 = _module.D7.create_DC17(_193_v121);
          let _196_v124;
          let _nw34 = Array((new BigNumber(16)).toNumber());
          _nw34[(_dafny.ZERO).toNumber()] = _193_v121;
          _nw34[(_dafny.ONE).toNumber()] = _193_v121;
          _nw34[(new BigNumber(2)).toNumber()] = (_194_v122)[_module.__default.safeIndex(_39_v4, new BigNumber((_194_v122).length))];
          _nw34[(new BigNumber(3)).toNumber()] = _193_v121;
          _nw34[(new BigNumber(4)).toNumber()] = _193_v121;
          _nw34[(new BigNumber(5)).toNumber()] = _193_v121;
          _nw34[(new BigNumber(6)).toNumber()] = _193_v121;
          _nw34[(new BigNumber(7)).toNumber()] = _193_v121;
          _nw34[(new BigNumber(8)).toNumber()] = _193_v121;
          _nw34[(new BigNumber(9)).toNumber()] = _193_v121;
          _nw34[(new BigNumber(10)).toNumber()] = _193_v121;
          _nw34[(new BigNumber(11)).toNumber()] = _193_v121;
          _nw34[(new BigNumber(12)).toNumber()] = _193_v121;
          _nw34[(new BigNumber(13)).toNumber()] = (_195_v123).dtor_cf25;
          _nw34[(new BigNumber(14)).toNumber()] = _193_v121;
          _nw34[(new BigNumber(15)).toNumber()] = _193_v121;
          _196_v124 = _nw34;
          let _index43 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_196_v124).length));
          (_196_v124)[_index43] = _193_v121;
          let _197_v125;
          let _nw35 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _197_v125 = _nw35;
          let _index44 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_196_v124).length));
          let _rhs41 = _40_v5;
          let _rhs42 = _193_v121;
          let _rhs43 = _197_v125;
          let _rhs44 = _39_v4;
          let _rhs45 = (_module.D2.create_DC3(_189_v117)).dtor_cf3;
          let _lhs40 = _196_v124;
          let _lhs41 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_196_v124).length));
          let _lhs42 = globalState;
          _40_v5 = _rhs41;
          _lhs40[_lhs41] = _rhs42;
          _197_v125 = _rhs43;
          _lhs42.f6 = _rhs44;
          _189_v117 = _rhs45;
          let _198_v126;
          _198_v126 = _dafny.Seq.of(_34_v0);
          let _199_v127;
          _199_v127 = _module.D3.create_DC6(_39_v4);
          let _200_v128;
          _200_v128 = _dafny.Seq.of(_199_v127);
          let _201_v129;
          _201_v129 = _module.D4.create_DC10(_39_v4, _40_v5, _39_v4, _35_v1.f15, _200_v128);
          let _202_v130;
          _202_v130 = _module.D3.create_DC7(_module.D3.create_DC6(_39_v4));
          let _203_v131;
          _203_v131 = _module.D3.create_DC7(_202_v130);
          let _204_v132;
          _204_v132 = _module.D3.create_DC7(_202_v130);
          let _pat_let_tv4 = _202_v130;
          let _205_v133;
          _205_v133 = _dafny.Set.fromElements(_204_v132, function (_pat_let3_0) {
            return function (_206_dt__update__tmp_h1) {
              return function (_pat_let4_0) {
                return function (_207_dt__update_hcf8_h0) {
                  return _module.D3.create_DC7(_207_dt__update_hcf8_h0);
                }(_pat_let4_0);
              }(_pat_let_tv4);
            }(_pat_let3_0);
          }(_204_v132), _204_v132, _204_v132);
          let _rhs46 = (_201_v129).dtor_cf16;
          let _rhs47 = _module.__default.safeModuloInt(new BigNumber(45), (_39_v4).minus(new BigNumber((_189_v117).length)));
          let _rhs48 = _dafny.Seq.update(_189_v117, _module.__default.safeIndex(new BigNumber(((_205_v133).Union(_205_v133)).length), new BigNumber((_189_v117).length)), _40_v5);
          let _rhs49 = _35_v1;
          let _rhs50 = _198_v126;
          let _lhs43 = globalState;
          let _lhs44 = globalState;
          _lhs43.f8 = _rhs46;
          _lhs44.f0 = _rhs47;
          _189_v117 = _rhs48;
          _35_v1 = _rhs49;
          _198_v126 = _rhs50;
          let _208_v134;
          _208_v134 = _module.D4.create_DC8(_34_v0);
          (globalState).f10 = (_208_v134).dtor_cf9;
        } else {
          let _209_v135;
          let _nw36 = new _module.C0();
          _nw36.__ctor(false);
          _209_v135 = _nw36;
          let _210_v136;
          _210_v136 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(globalState),_35_v1.f15);
          let _211_v137;
          _211_v137 = _dafny.Set.fromElements(_40_v5);
          _210_v136 = (_210_v136).update(_209_v135.f15, (_211_v137).IsProperSubsetOf(_211_v137));
          let _212_v138;
          _212_v138 = _dafny.Map.Empty.slice().updateUnsafe(_39_v4,_209_v135);
          let _213_v139;
          _213_v139 = _dafny.Seq.UnicodeFromString("medr");
          let _214_v140;
          _214_v140 = _dafny.Map.Empty.slice().updateUnsafe(((!(_35_v1.f15)) ? (_39_v4) : (_39_v4)),(((_212_v138).contains(new BigNumber((_213_v139).length))) ? ((_212_v138).get(new BigNumber((_213_v139).length))) : (_35_v1)));
          _212_v138 = (_212_v138).update(_39_v4, _35_v1);
          let _215_v141;
          _215_v141 = _dafny.Seq.of(_209_v135.f15);
          let _index45 = _module.__default.safeIndex(new BigNumber(536), new BigNumber((_178_v108).length));
          (_178_v108)[_index45] = _dafny.Seq.IsPrefixOf(_215_v141, _215_v141);
          let _index46 = _module.__default.safeIndex(new BigNumber(536), new BigNumber((_178_v108).length));
          let _rhs51 = _213_v139;
          let _rhs52 = new BigNumber(959);
          let _rhs53 = _35_v1.f15;
          let _lhs45 = globalState;
          let _lhs46 = _178_v108;
          let _lhs47 = _module.__default.safeIndex(new BigNumber(536), new BigNumber((_178_v108).length));
          _213_v139 = _rhs51;
          _lhs45.f0 = _rhs52;
          _lhs46[_lhs47] = _rhs53;
          let _216_v142;
          _216_v142 = _dafny.Set.fromElements(_34_v0);
          let _217_v143;
          let _nw37 = Array((new BigNumber(8)).toNumber());
          _nw37[(_dafny.ZERO).toNumber()] = _39_v4;
          _nw37[(_dafny.ONE).toNumber()] = _module.__default.fm0(true, _39_v4, (_215_v141)[_module.__default.safeIndex(_module.__default.fm0(_209_v135.f15, _39_v4, (_215_v141)[_module.__default.safeIndex(new BigNumber((_216_v142).length), new BigNumber((_215_v141).length))], globalState), new BigNumber((_215_v141).length))], globalState);
          _nw37[(new BigNumber(2)).toNumber()] = (_39_v4).plus(new BigNumber((_213_v139).length));
          _nw37[(new BigNumber(3)).toNumber()] = _39_v4;
          _nw37[(new BigNumber(4)).toNumber()] = _module.__default.safeModuloInt(_39_v4, _39_v4);
          _nw37[(new BigNumber(5)).toNumber()] = _39_v4;
          _nw37[(new BigNumber(6)).toNumber()] = ((_35_v1.f15) ? (_39_v4) : (_39_v4));
          _nw37[(new BigNumber(7)).toNumber()] = _39_v4;
          _217_v143 = _nw37;
          let _index47 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_217_v143).length));
          (_217_v143)[_index47] = _module.__default.safeDivisionInt(_39_v4, _39_v4);
          let _218_v144;
          _218_v144 = _dafny.Map.Empty.slice().updateUnsafe(_39_v4,_39_v4);
          let _219_v146;
          _219_v146 = _module.D4.create_DC8(true);
          let _index48 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_217_v143).length));
          let _rhs54 = !(_218_v144).equals(function () {
            let _coll7 = new _dafny.Map();
            for (const _compr_7 of _dafny.IntegerRange(new BigNumber(-450), new BigNumber(868))) {
              let _220_v145 = _compr_7;
              if (((new BigNumber(-450)).isLessThanOrEqualTo(_220_v145)) && ((_220_v145).isLessThan(new BigNumber(868)))) {
                _coll7.push([_module.__default.safeModuloInt(_220_v145, _39_v4),new BigNumber((_dafny.MultiSet.fromElements(_39_v4, _39_v4)).cardinality())]);
              }
            }
            return _coll7;
          }());
          let _rhs55 = _module.__default.safeDivisionInt(_39_v4, _module.__default.safeModuloInt(_39_v4, _39_v4));
          let _rhs56 = (_35_v1.f15) || ((_219_v146).dtor_cf9);
          let _rhs57 = _module.__default.safeModuloInt(_39_v4, new BigNumber((_module.__default.fm14(globalState)).length));
          let _lhs48 = globalState;
          let _lhs49 = _217_v143;
          let _lhs50 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_217_v143).length));
          let _lhs51 = globalState;
          let _lhs52 = globalState;
          _lhs48.f1 = _rhs54;
          _lhs49[_lhs50] = _rhs55;
          _lhs51.f8 = _rhs56;
          _lhs52.f6 = _rhs57;
        }
        let _221_v147;
        let _nw38 = Array((new BigNumber(2)).toNumber());
        _221_v147 = _nw38;
        let _index49 = _module.__default.safeIndex(new BigNumber(141), new BigNumber((_221_v147).length));
        (_221_v147)[_index49] = _35_v1;
        let _index50 = _module.__default.safeIndex(new BigNumber(141), new BigNumber((_221_v147).length));
        (_221_v147)[_index50] = _35_v1;
      }
      let _222_v148;
      _222_v148 = _dafny.Seq.UnicodeFromString("usstt");
      _39_v4 = (_39_v4).plus(new BigNumber((_dafny.Seq.Concat(_222_v148, _222_v148)).length));
      return;
    }
    static Main(__noArgsParameter) {
      let _223_v0;
      let _init7 = function (_224_i0) {
        return !(!(!(true)));
      };
      let _nw39 = Array((new BigNumber(18)).toNumber());
      for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw39.length); _i0_7++) {
        _nw39[_i0_7] = _init7(new BigNumber(_i0_7));
      }
      _223_v0 = _nw39;
      let _225_v1;
      _225_v1 = _223_v0;
      let _226_v2;
      _226_v2 = true;
      let _227_v3;
      _227_v3 = _dafny.Map.Empty.slice().updateUnsafe(_226_v2,true);
      let _228_globalState;
      let _nw40 = new _module.GlobalState();
      _nw40.__ctor(new BigNumber(527), false, true, new BigNumber(871), new BigNumber(272), new BigNumber(844), new BigNumber(508), true, true, _223_v0, false, (_225_v1), (_227_v3).Merge(_227_v3), new BigNumber(357), false);
      _228_globalState = _nw40;
      let _229_v4;
      _229_v4 = new BigNumber(-669);
      (_228_globalState).f6 = _module.__default.safeDivisionInt(_229_v4, new BigNumber(-539));
      let _230_v5;
      _230_v5 = _dafny.Map.Empty.slice().updateUnsafe(_226_v2,_module.__default.fm0(true, _229_v4, _226_v2, _228_globalState));
      if ((_226_v2) === ((_230_v5).contains(_226_v2))) {
        _226_v2 = _226_v2;
        if (_226_v2) {
          let _231_v6;
          _231_v6 = _dafny.Seq.of(_226_v2);
          _226_v2 = _dafny.Seq.contains(_231_v6, _226_v2);
          let _232_v7;
          _232_v7 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_223_v0,_229_v4));
          let _233_v8;
          _233_v8 = _dafny.Map.Empty.slice().updateUnsafe((_232_v7)[_module.__default.safeIndex(_229_v4, new BigNumber((_232_v7).length))],false);
          let _234_v9;
          _234_v9 = new _dafny.CodePoint('h'.codePointAt(0));
          let _rhs58 = _226_v2;
          let _rhs59 = _233_v8;
          let _rhs60 = _234_v9;
          let _rhs61 = _229_v4;
          let _lhs53 = _228_globalState;
          let _lhs54 = _228_globalState;
          _lhs53.f1 = _rhs58;
          _233_v8 = _rhs59;
          _234_v9 = _rhs60;
          _lhs54.f6 = _rhs61;
          let _235_v10;
          let _nw41 = new _module.C0();
          _nw41.__ctor(_module.__default.fm2(_228_globalState));
          _235_v10 = _nw41;
          let _236_v11;
          _236_v11 = _dafny.Seq.UnicodeFromString("m");
          let _237_v12;
          _237_v12 = _dafny.Set.fromElements(new BigNumber(893), (new BigNumber((_236_v11).length)).plus(_229_v4), _229_v4, _229_v4, _229_v4);
          let _238_v13;
          _238_v13 = _dafny.Map.Empty.slice().updateUnsafe(_235_v10,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(588)), ((_239_v9) => function (_240_i1) {
            return _239_v9;
          })(_234_v9))).length));
          _237_v12 = _dafny.Set.fromElements(new BigNumber((_238_v13).length), _229_v4);
          (_228_globalState).f0 = _229_v4;
        } else {
          let _241_v14;
          _241_v14 = _module.D1.create_DC1(new BigNumber(-329));
          (_228_globalState).f6 = (_241_v14).dtor_cf1;
          let _242_v15;
          _242_v15 = _dafny.Seq.UnicodeFromString("pwxho");
          let _243_v16;
          _243_v16 = _dafny.Map.Empty.slice().updateUnsafe(_229_v4,_226_v2);
          let _244_v17;
          _244_v17 = _module.D2.create_DC3(_dafny.Seq.update(_242_v15, _module.__default.safeIndex(_229_v4, new BigNumber((_242_v15).length)), new _dafny.CodePoint('b'.codePointAt(0))));
          let _245_v18;
          _245_v18 = _dafny.Seq.of(_226_v2);
          let _rhs62 = (_244_v17).dtor_cf3;
          let _rhs63 = new BigNumber((_dafny.Seq.Concat(_245_v18, _245_v18)).length);
          let _rhs64 = (!(_226_v2)) === (!(_226_v2));
          let _rhs65 = _243_v16;
          let _lhs55 = _228_globalState;
          let _lhs56 = _228_globalState;
          _242_v15 = _rhs62;
          _lhs55.f6 = _rhs63;
          _lhs56.f8 = _rhs64;
          _243_v16 = _rhs65;
          let _246_v19;
          _246_v19 = _dafny.Seq.of(_229_v4);
          let _247_v20;
          _247_v20 = new _dafny.CodePoint('h'.codePointAt(0));
          let _248_v21;
          _248_v21 = _dafny.Map.Empty.slice().updateUnsafe(_247_v20,_229_v4);
          let _249_v22;
          let _nw42 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
          _249_v22 = _nw42;
          let _250_v23;
          _250_v23 = _dafny.Map.Empty.slice().updateUnsafe(_249_v22,_229_v4);
          _246_v19 = _dafny.Seq.Concat(_246_v19, _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(943)), function (_251_i2) {
            return new _dafny.CodePoint('n'.codePointAt(0));
          })).length), new BigNumber((_dafny.Seq.of(new BigNumber((_248_v21).length), new BigNumber((_227_v3).length))).length), new BigNumber((_250_v23).length)), _246_v19));
          let _252_v24;
          let _nw43 = new _module.C0();
          _nw43.__ctor(!(_226_v2) || (_226_v2));
          _252_v24 = _nw43;
          (_228_globalState).f8 = _252_v24.f15;
        }
        if (_226_v2) {
          let _253_v25;
          _253_v25 = _dafny.Seq.UnicodeFromString("xygbbejn");
          let _254_v26;
          _254_v26 = _dafny.Map.Empty.slice().updateUnsafe(_253_v25,_226_v2);
          let _255_v27;
          let _nw44 = new _module.C0();
          _nw44.__ctor((((_254_v26).contains(_253_v25)) ? ((_254_v26).get(_253_v25)) : (_226_v2)));
          _255_v27 = _nw44;
          let _256_v28;
          _256_v28 = _dafny.Map.Empty.slice().updateUnsafe(_255_v27,_255_v27.f15);
          _256_v28 = (_256_v28).update(_255_v27, _module.__default.fm2(_228_globalState));
          _module.__default.m0(_228_globalState);
          let _257_v29;
          _257_v29 = _dafny.Set.fromElements(!(_226_v2) || (!(_226_v2)));
          _257_v29 = (_module.__default.fm3(_229_v4, _228_globalState)).Difference((_257_v29).Union(_dafny.Set.fromElements(_255_v27.f15)));
          (_228_globalState).f8 = _255_v27.f15;
          let _258_v30;
          let _init8 = ((_259_v4) => function (_260_i3) {
            return (_260_i3).plus(_259_v4);
          })(_229_v4);
          let _nw45 = Array((new BigNumber(14)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw45.length); _i0_8++) {
            _nw45[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _258_v30 = _nw45;
          let _261_v31;
          _261_v31 = _dafny.MultiSet.fromElements(_258_v30);
          let _262_v32;
          _262_v32 = _dafny.Seq.of(_261_v31, _261_v31);
          let _263_v33;
          _263_v33 = _dafny.Seq.of(_229_v4, _229_v4, _229_v4, new BigNumber(((_262_v32)[_module.__default.safeIndex(_229_v4, new BigNumber((_262_v32).length))]).cardinality()), _229_v4);
          _263_v33 = _263_v33;
        } else {
          let _264_v34;
          let _nw46 = Array((new BigNumber(27)).toNumber()).fill([]);
          _264_v34 = _nw46;
          let _index51 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_264_v34).length));
          (_264_v34)[_index51] = _223_v0;
          let _265_v35;
          _265_v35 = _dafny.Set.fromElements(true);
          let _266_v36;
          _266_v36 = _dafny.Set.fromElements(_module.__default.fm0(_226_v2, _229_v4, _226_v2, _228_globalState), (_229_v4).multipliedBy(new BigNumber((_265_v35).length)), _229_v4, (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_230_v5).length), _229_v4)));
          let _267_v37;
          _267_v37 = _dafny.Seq.UnicodeFromString("xd");
          let _index52 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_264_v34).length));
          let _rhs66 = _223_v0;
          let _rhs67 = _229_v4;
          let _rhs68 = (_module.__default.fm4(_226_v2, _229_v4, _228_globalState)).Union((_266_v36).Intersect(_266_v36));
          let _rhs69 = _226_v2;
          let _rhs70 = ((_dafny.Seq.IsProperPrefixOf(_267_v37, _267_v37)) ? (_226_v2) : (_226_v2));
          let _lhs57 = _264_v34;
          let _lhs58 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_264_v34).length));
          let _lhs59 = _228_globalState;
          let _lhs60 = _228_globalState;
          let _lhs61 = _228_globalState;
          _lhs57[_lhs58] = _rhs66;
          _lhs59.f0 = _rhs67;
          _266_v36 = _rhs68;
          _lhs60.f10 = _rhs69;
          _lhs61.f8 = _rhs70;
          let _268_v38;
          let _nw47 = new _module.C0();
          _nw47.__ctor(_226_v2);
          _268_v38 = _nw47;
          let _269_v39;
          let _nw48 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
          _269_v39 = _nw48;
          let _270_v40;
          _270_v40 = _dafny.Map.Empty.slice().updateUnsafe(_229_v4,_229_v4);
          let _index53 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_269_v39).length));
          (_269_v39)[_index53] = new BigNumber((_270_v40).length);
          let _271_v41;
          _271_v41 = _module.D2.create_DC4(new BigNumber(613), _229_v4);
          let _272_v42;
          _272_v42 = _dafny.Seq.of(_226_v2, _226_v2, _268_v38.f15);
          let _273_v43;
          _273_v43 = _dafny.Map.Empty.slice().updateUnsafe(_271_v41,_272_v42);
          let _274_v44;
          _274_v44 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_268_v38.f15, _268_v38.f15, _268_v38.f15), _272_v42);
          let _index54 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_269_v39).length));
          let _rhs71 = _268_v38;
          let _rhs72 = _269_v39;
          let _rhs73 = (new BigNumber((_274_v44).cardinality())).minus(_229_v4);
          let _rhs74 = _273_v43;
          let _lhs62 = _269_v39;
          let _lhs63 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_269_v39).length));
          _268_v38 = _rhs71;
          _269_v39 = _rhs72;
          _lhs62[_lhs63] = _rhs73;
          _273_v43 = _rhs74;
          (_228_globalState).f1 = _module.__default.fm2(_228_globalState);
          _module.__default.m0(_228_globalState);
          let _275_v45;
          let _init9 = ((_276_v42, _277_v39, _278_v38) => function (_279_i4) {
            return _dafny.Seq.update(_dafny.Seq.Concat(_276_v42, _276_v42), _module.__default.safeIndex((_277_v39)[_module.__default.safeIndex(new BigNumber(726), new BigNumber((_277_v39).length))], new BigNumber((_dafny.Seq.Concat(_276_v42, _276_v42)).length)), _278_v38.f15);
          })(_272_v42, _269_v39, _268_v38);
          let _nw49 = Array((new BigNumber(10)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw49.length); _i0_9++) {
            _nw49[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _275_v45 = _nw49;
          let _index55 = _module.__default.safeIndex(new BigNumber(25), new BigNumber((_275_v45).length));
          (_275_v45)[_index55] = _272_v42;
          let _index56 = _module.__default.safeIndex(new BigNumber(25), new BigNumber((_275_v45).length));
          (_275_v45)[_index56] = _272_v42;
        }
        let _280_v46;
        _280_v46 = _module.D2.create_DC3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(707)), function (_281_i5) {
  return new _dafny.CodePoint('q'.codePointAt(0));
}));
        let _282_v47;
        _282_v47 = _dafny.Seq.UnicodeFromString("yqrgy");
        let _pat_let_tv5 = _282_v47;
        let _source3 = function (_pat_let5_0) {
          return function (_283_dt__update__tmp_h0) {
            return function (_pat_let6_0) {
              return function (_284_dt__update_hcf3_h0) {
                return _module.D2.create_DC3(_284_dt__update_hcf3_h0);
              }(_pat_let6_0);
            }(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("b"), _pat_let_tv5));
          }(_pat_let5_0);
        }(_280_v46);
        if (_source3.is_DC4) {
          let _285___mcc_h0 = (_source3).cf4;
          let _286___mcc_h1 = (_source3).cf5;
          let _287_cf5 = _286___mcc_h1;
          let _288_cf4 = _285___mcc_h0;
          let _289_v48;
          _289_v48 = new _dafny.CodePoint('d'.codePointAt(0));
          let _rhs75 = _module.__default.safeDivisionInt(new BigNumber(654), new BigNumber(-318));
          let _rhs76 = _dafny.Seq.contains(_282_v47, _289_v48);
          let _rhs77 = !((_229_v4).isLessThanOrEqualTo(_288_cf4));
          let _rhs78 = _226_v2;
          let _rhs79 = _288_cf4;
          let _lhs64 = _228_globalState;
          let _lhs65 = _228_globalState;
          let _lhs66 = _228_globalState;
          let _lhs67 = _228_globalState;
          let _lhs68 = _228_globalState;
          _lhs64.f4 = _rhs75;
          _lhs65.f1 = _rhs76;
          _lhs66.f10 = _rhs77;
          _lhs67.f1 = _rhs78;
          _lhs68.f6 = _rhs79;
          let _290_v49;
          _290_v49 = _dafny.Seq.of(_223_v0);
          _290_v49 = _dafny.Seq.of(_223_v0, (_225_v1), _223_v0);
          (_228_globalState).f6 = _288_cf4;
          let _291_v50;
          _291_v50 = _module.D1.create_DC2(_288_cf4);
          let _292_v51;
          _292_v51 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_282_v47).length),new BigNumber((_282_v47).length));
          let _293_v52;
          _293_v52 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_291_v50)).length),_292_v51);
          let _294_v53;
          _294_v53 = _dafny.Set.fromElements(_226_v2, _226_v2);
          let _295_v54;
          let _nw50 = Array((new BigNumber(14)).toNumber());
          _nw50[(_dafny.ZERO).toNumber()] = _287_cf5;
          _nw50[(_dafny.ONE).toNumber()] = new BigNumber(-32);
          _nw50[(new BigNumber(2)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(893), _287_cf5);
          _nw50[(new BigNumber(3)).toNumber()] = _288_cf4;
          _nw50[(new BigNumber(4)).toNumber()] = _module.__default.fm0(_226_v2, _287_cf5, _module.__default.fm2(_228_globalState), _228_globalState);
          _nw50[(new BigNumber(5)).toNumber()] = _287_cf5;
          _nw50[(new BigNumber(6)).toNumber()] = _module.__default.fm0(_226_v2, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(230)), ((_296_cf5) => function (_297_i6) {
            return _296_cf5;
          })(_287_cf5))).length), _226_v2, _228_globalState);
          _nw50[(new BigNumber(7)).toNumber()] = _288_cf4;
          _nw50[(new BigNumber(8)).toNumber()] = _229_v4;
          _nw50[(new BigNumber(9)).toNumber()] = _module.__default.safeDivisionInt(_229_v4, _229_v4);
          _nw50[(new BigNumber(10)).toNumber()] = new BigNumber((_293_v52).length);
          _nw50[(new BigNumber(11)).toNumber()] = _229_v4;
          _nw50[(new BigNumber(12)).toNumber()] = (_287_cf5).plus(new BigNumber((_294_v53).length));
          _nw50[(new BigNumber(13)).toNumber()] = _229_v4;
          _295_v54 = _nw50;
          let _index57 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_295_v54).length));
          (_295_v54)[_index57] = _229_v4;
          let _index58 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_295_v54).length));
          (_295_v54)[_index58] = (_288_cf4).multipliedBy(_288_cf4);
        } else {
          let _298___mcc_h2 = (_source3).cf3;
          let _299_cf3 = _298___mcc_h2;
          (_228_globalState).f4 = _229_v4;
          let _300_v55;
          _300_v55 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_226_v2, _229_v4, _226_v2, _228_globalState),_226_v2);
          let _301_v56;
          _301_v56 = _dafny.Map.Empty.slice().updateUnsafe(_226_v2,_300_v55);
          let _302_v57;
          let _nw51 = Array((new BigNumber(17)).toNumber());
          _nw51[(_dafny.ZERO).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(_228_globalState),_dafny.Map.Empty.slice().updateUnsafe(_229_v4,!(_226_v2)))).Merge(_301_v56);
          _nw51[(_dafny.ONE).toNumber()] = _301_v56;
          _nw51[(new BigNumber(2)).toNumber()] = _301_v56;
          _nw51[(new BigNumber(3)).toNumber()] = _301_v56;
          _nw51[(new BigNumber(4)).toNumber()] = _301_v56;
          _nw51[(new BigNumber(5)).toNumber()] = _301_v56;
          _nw51[(new BigNumber(6)).toNumber()] = (_301_v56).Merge(_301_v56);
          _nw51[(new BigNumber(7)).toNumber()] = _301_v56;
          _nw51[(new BigNumber(8)).toNumber()] = _301_v56;
          _nw51[(new BigNumber(9)).toNumber()] = _301_v56;
          _nw51[(new BigNumber(10)).toNumber()] = _301_v56;
          _nw51[(new BigNumber(11)).toNumber()] = _module.__default.fm5(_228_globalState);
          _nw51[(new BigNumber(12)).toNumber()] = (_301_v56).Merge(_dafny.Map.Empty.slice().updateUnsafe(_226_v2,_300_v55));
          _nw51[(new BigNumber(13)).toNumber()] = _301_v56;
          _nw51[(new BigNumber(14)).toNumber()] = (_module.__default.fm5(_228_globalState)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_226_v2,_300_v55));
          _nw51[(new BigNumber(15)).toNumber()] = _301_v56;
          _nw51[(new BigNumber(16)).toNumber()] = _301_v56;
          _302_v57 = _nw51;
          let _index59 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_302_v57).length));
          (_302_v57)[_index59] = _301_v56;
          let _index60 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_302_v57).length));
          (_302_v57)[_index60] = ((_301_v56).Merge(_301_v56)).Merge(_301_v56);
          (_228_globalState).f0 = new BigNumber(-334);
          let _303_v58;
          let _nw52 = new _module.C0();
          _nw52.__ctor(_226_v2);
          _303_v58 = _nw52;
        }
        let _304_v59;
        _304_v59 = _module.D1.create_DC2(_229_v4);
        (_228_globalState).f8 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_282_v47, _module.__default.fm6(_229_v4, _226_v2, _304_v59, _226_v2, _228_globalState)), _282_v47);
      } else {
        let _index61 = _module.__default.safeIndex(new BigNumber(865), new BigNumber((_223_v0).length));
        (_223_v0)[_index61] = _226_v2;
        let _305_v60;
        _305_v60 = _dafny.Seq.of(_226_v2);
        let _index62 = _module.__default.safeIndex(new BigNumber(865), new BigNumber((_223_v0).length));
        (_223_v0)[_index62] = _dafny.Seq.IsPrefixOf(((_226_v2) ? (_305_v60) : (_dafny.Seq.of(_module.__default.fm2(_228_globalState), _226_v2))), _305_v60);
        (_228_globalState).f6 = (_module.__default.safeModuloInt(new BigNumber(-760), (_dafny.ZERO).minus(_229_v4))).multipliedBy((((_223_v0)[_module.__default.safeIndex(new BigNumber(865), new BigNumber((_223_v0).length))]) ? (_229_v4) : (_229_v4)));
        let _306_v61;
        _306_v61 = _dafny.Seq.UnicodeFromString("dqluto");
        _306_v61 = _306_v61;
        let _307_v62;
        _307_v62 = _dafny.Seq.of(_229_v4);
        let _308_v63;
        _308_v63 = _dafny.Map.Empty.slice().updateUnsafe((_307_v62)[_module.__default.safeIndex(_229_v4, new BigNumber((_307_v62).length))],(_223_v0)[_module.__default.safeIndex(new BigNumber(865), new BigNumber((_223_v0).length))]);
        let _309_v64;
        _309_v64 = _module.D3.create_DC5(_305_v60);
        _308_v63 = (_308_v63).update((_dafny.ZERO).minus((_229_v4).plus((_dafny.ZERO).minus(_229_v4))), _dafny.Seq.IsPrefixOf(_305_v60, (_309_v64).dtor_cf6));
        (_228_globalState).f4 = (_229_v4).minus(_229_v4);
      }
      (_228_globalState).f1 = _226_v2;
      let _310_i7;
      _310_i7 = _dafny.ZERO;
      L1: {
        while (!(_229_v4).isEqualTo(new BigNumber(439))) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_310_i7)) {
              break L1;
            }
            _310_i7 = (_310_i7).plus(_dafny.ONE);
            let _311_v65;
            _311_v65 = _dafny.Seq.UnicodeFromString("l");
            (_228_globalState).f8 = !_dafny.areEqual(_311_v65, _dafny.Seq.UnicodeFromString("avcf"));
            let _312_v66;
            let _nw53 = new _module.C0();
            _nw53.__ctor((_226_v2) || (_226_v2));
            _312_v66 = _nw53;
            _312_v66 = _312_v66;
            _227_v3 = (_227_v3).update(_226_v2, _312_v66.f15);
            let _313_v67;
            let _init10 = ((_314_v4) => function (_315_i8) {
              return _dafny.Seq.Concat(_dafny.Seq.of(_314_v4), _dafny.Seq.Create(_module.__default.abs(new BigNumber(159)), function (_316_i9) {
                return new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)), new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)))).cardinality());
              }));
            })(_229_v4);
            let _nw54 = Array((new BigNumber(10)).toNumber());
            for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw54.length); _i0_10++) {
              _nw54[_i0_10] = _init10(new BigNumber(_i0_10));
            }
            _313_v67 = _nw54;
            let _317_v68;
            _317_v68 = _dafny.Seq.of(_229_v4, _229_v4, _229_v4, _229_v4, _229_v4);
            let _index63 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_313_v67).length));
            (_313_v67)[_index63] = _317_v68;
            let _318_v69;
            _318_v69 = _dafny.Seq.of(_312_v66.f15, _312_v66.f15, _module.__default.fm2(_228_globalState), _226_v2);
            let _index64 = _module.__default.safeIndex(new BigNumber(30), new BigNumber((_313_v67).length));
            (_313_v67)[_index64] = _module.__default.fm7((_229_v4).plus(_229_v4), true, _318_v69, new BigNumber((_dafny.Seq.UnicodeFromString("lhvsj")).length), _228_globalState);
          }
        }
      }
      let _319_v70;
      let _nw55 = new _module.C0();
      _nw55.__ctor(false);
      _319_v70 = _nw55;
      let _320_v71;
      _320_v71 = _module.D4.create_DC9(_223_v0, _319_v70, !(_319_v70.f15));
      let _321_v72;
      let _nw56 = new _module.C0();
      _nw56.__ctor((_320_v71).dtor_cf12);
      _321_v72 = _nw56;
      let _322_v73;
      let _init11 = ((_323_v4) => function (_324_i10) {
        return _module.__default.safeModuloInt(_324_i10, _323_v4);
      })(_229_v4);
      let _nw57 = Array((new BigNumber(21)).toNumber());
      for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw57.length); _i0_11++) {
        _nw57[_i0_11] = _init11(new BigNumber(_i0_11));
      }
      _322_v73 = _nw57;
      let _index65 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_322_v73).length));
      (_322_v73)[_index65] = _229_v4;
      let _index66 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_322_v73).length));
      (_322_v73)[_index66] = _module.__default.fm0((false) || (_319_v70.f15), _229_v4, _321_v72.f15, _228_globalState);
      let _325_v74;
      _325_v74 = new _dafny.CodePoint('f'.codePointAt(0));
      let _326_v75;
      _326_v75 = _dafny.Map.Empty.slice().updateUnsafe(_321_v72.f15,_325_v74);
      let _pat_let_tv6 = _322_v73;
      let _pat_let_tv7 = _322_v73;
      let _327_v76;
      _327_v76 = _module.D4.create_DC10(_229_v4, _325_v74, _229_v4, _226_v2, _dafny.Seq.of(function (_pat_let7_0) {
  return function (_328_dt__update__tmp_h1) {
    return function (_pat_let8_0) {
      return function (_329_dt__update_hcf7_h0) {
        return _module.D3.create_DC6(_329_dt__update_hcf7_h0);
      }(_pat_let8_0);
    }((_pat_let_tv7)[_module.__default.safeIndex(new BigNumber(626), new BigNumber((_pat_let_tv6).length))]);
  }(_pat_let7_0);
}(_module.D3.create_DC6(new BigNumber((_dafny.Seq.UnicodeFromString("kjuqikydf")).length))), _module.__default.fm8(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_319_v70.f15,(_322_v73)[_module.__default.safeIndex(new BigNumber(626), new BigNumber((_322_v73).length))])).length),_module.__default.fm2(_228_globalState))).length), _226_v2, _228_globalState)));
      let _330_v77;
      _330_v77 = _module.D3.create_DC5(_dafny.Seq.of(_321_v72.f15));
      let _pat_let_tv8 = _322_v73;
      let _pat_let_tv9 = _322_v73;
      let _pat_let_tv10 = _229_v4;
      let _pat_let_tv11 = _229_v4;
      let _index67 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_322_v73).length));
      let _index68 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_322_v73).length));
      let _index69 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_322_v73).length));
      let _rhs80 = (_229_v4).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),_dafny.MultiSet.FromArray(_dafny.Seq.of(_225_v1, _225_v1, _225_v1)))).length));
      let _rhs81 = (_229_v4).minus((_229_v4).plus(new BigNumber((_326_v75).length)));
      let _rhs82 = _321_v72.f15;
      let _rhs83 = (_327_v76).dtor_cf13;
      let _rhs84 = function (_source4) {
        if (_source4.is_DC6) {
          let _331___mcc_h3 = (_source4).cf7;
          let _332_cf7 = _331___mcc_h3;
          return ((_pat_let_tv9)[_module.__default.safeIndex(new BigNumber(626), new BigNumber((_pat_let_tv8).length))]).multipliedBy(_332_cf7);
        } else if (_source4.is_DC5) {
          let _333___mcc_h4 = (_source4).cf6;
          let _334_cf6 = _333___mcc_h4;
          return _pat_let_tv10;
        } else {
          let _335___mcc_h5 = (_source4).cf8;
          let _336_cf8 = _335___mcc_h5;
          return (_dafny.ZERO).minus(_pat_let_tv11);
        }
      }(_330_v77);
      let _lhs69 = _322_v73;
      let _lhs70 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_322_v73).length));
      let _lhs71 = _322_v73;
      let _lhs72 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_322_v73).length));
      let _lhs73 = _228_globalState;
      let _lhs74 = _322_v73;
      let _lhs75 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_322_v73).length));
      _lhs69[_lhs70] = _rhs80;
      _lhs71[_lhs72] = _rhs81;
      _lhs73.f1 = _rhs82;
      _229_v4 = _rhs83;
      _lhs74[_lhs75] = _rhs84;
      (_228_globalState).f10 = (_229_v4).isLessThanOrEqualTo(_229_v4);
      (_228_globalState).f6 = ((_322_v73)[_module.__default.safeIndex(new BigNumber(626), new BigNumber((_322_v73).length))]).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_322_v73)[_module.__default.safeIndex(new BigNumber(626), new BigNumber((_322_v73).length))],_229_v4)).length)));
      let _hi0 = _229_v4;
      for (let _337_i11 = new BigNumber(685); _337_i11.isLessThan(_hi0); _337_i11 = _337_i11.plus(_dafny.ONE)) {
        let _338_v78;
        _338_v78 = _dafny.Map.Empty.slice().updateUnsafe(_337_i11,new BigNumber(-16));
        let _339_v79;
        _339_v79 = _dafny.Map.Empty.slice().updateUnsafe(_325_v74,_319_v70.f15);
        let _340_v80;
        _340_v80 = _dafny.MultiSet.fromElements((_322_v73)[_module.__default.safeIndex(new BigNumber(626), new BigNumber((_322_v73).length))], _229_v4, _229_v4, new BigNumber((_339_v79).length));
        let _341_v81;
        _341_v81 = _dafny.Map.Empty.slice().updateUnsafe(_319_v70,new BigNumber(159));
        let _342_v82;
        let _nw58 = Array((new BigNumber(17)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _342_v82 = _nw58;
        let _343_v83;
        _343_v83 = _dafny.Seq.of(_321_v72.f15, _319_v70.f15);
        let _344_v84;
        _344_v84 = _dafny.Map.Empty.slice().updateUnsafe(_342_v82,new BigNumber((_343_v83).length));
        let _345_v85;
        _345_v85 = _dafny.Map.Empty.slice().updateUnsafe((((_338_v78).contains(_229_v4)) ? ((_338_v78).get(_229_v4)) : ((((_340_v80).contains(new BigNumber(8))) ? ((_340_v80).get(new BigNumber(8))) : (_337_i11)))),((_341_v81).update(_319_v70, (_dafny.ZERO).minus(new BigNumber((_344_v84).length)))).Merge(_341_v81));
        _345_v85 = (_345_v85).update((_dafny.ZERO).minus((_322_v73)[_module.__default.safeIndex(new BigNumber(626), new BigNumber((_322_v73).length))]), ((_319_v70.f15) ? (_341_v81) : (_341_v81)));
        _module.__default.m0(_228_globalState);
        _340_v80 = _340_v80;
        let _346_v86;
        _346_v86 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(824)), ((_347_v74) => function (_348_i13) {
          return _347_v74;
        })(_325_v74)));
        let _349_v87;
        _349_v87 = _dafny.Seq.UnicodeFromString("kqs");
        let _350_v88;
        _350_v88 = _dafny.Set.fromElements(_226_v2, _319_v70.f15, _319_v70.f15, _321_v72.f15, _319_v70.f15);
        let _351_v89;
        _351_v89 = _module.D1.create_DC2(new BigNumber((_350_v88).length));
        let _352_v90;
        let _nw59 = Array((new BigNumber(11)).toNumber());
        _nw59[(_dafny.ZERO).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(687)), ((_353_v74) => function (_354_i12) {
          return _353_v74;
        })(_325_v74));
        _nw59[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat((_346_v86)[_module.__default.safeIndex(_337_i11, new BigNumber((_346_v86).length))], _349_v87);
        _nw59[(new BigNumber(2)).toNumber()] = _349_v87;
        _nw59[(new BigNumber(3)).toNumber()] = _349_v87;
        _nw59[(new BigNumber(4)).toNumber()] = _349_v87;
        _nw59[(new BigNumber(5)).toNumber()] = _349_v87;
        _nw59[(new BigNumber(6)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(847)), ((_355_v74) => function (_356_i14) {
          return _355_v74;
        })(_325_v74));
        _nw59[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_349_v87, _dafny.Seq.update(_349_v87, _module.__default.safeIndex(_229_v4, new BigNumber((_349_v87).length)), _325_v74));
        _nw59[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_349_v87, _349_v87);
        _nw59[(new BigNumber(9)).toNumber()] = _349_v87;
        _nw59[(new BigNumber(10)).toNumber()] = _module.__default.fm6(_229_v4, (_320_v71).dtor_cf12, _351_v89, _226_v2, _228_globalState);
        _352_v90 = _nw59;
        let _index70 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_352_v90).length));
        (_352_v90)[_index70] = _349_v87;
        let _index71 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_352_v90).length));
        let _rhs85 = _319_v70.f15;
        let _rhs86 = _319_v70.f15;
        let _rhs87 = (_226_v2) || (!((_337_i11).isLessThanOrEqualTo(_337_i11)));
        let _rhs88 = _349_v87;
        let _rhs89 = (((_322_v73)[_module.__default.safeIndex(new BigNumber(626), new BigNumber((_322_v73).length))]).plus((_322_v73)[_module.__default.safeIndex(new BigNumber(626), new BigNumber((_322_v73).length))])).multipliedBy(_229_v4);
        let _lhs76 = _321_v72;
        let _lhs77 = _228_globalState;
        let _lhs78 = _228_globalState;
        let _lhs79 = _352_v90;
        let _lhs80 = _module.__default.safeIndex(new BigNumber(219), new BigNumber((_352_v90).length));
        let _lhs81 = _228_globalState;
        _lhs76.f15 = _rhs85;
        _lhs77.f10 = _rhs86;
        _lhs78.f10 = _rhs87;
        _lhs79[_lhs80] = _rhs88;
        _lhs81.f0 = _rhs89;
      }
      let _357_v91;
      _357_v91 = _dafny.Seq.UnicodeFromString("ymky");
      let _hi1 = _module.__default.fm0(_319_v70.f15, new BigNumber((_357_v91).length), _module.__default.fm2(_228_globalState), _228_globalState);
      for (let _358_i15 = _module.__default.safeModuloInt(_229_v4, (_322_v73)[_module.__default.safeIndex(new BigNumber(626), new BigNumber((_322_v73).length))]); _358_i15.isLessThan(_hi1); _358_i15 = _358_i15.plus(_dafny.ONE)) {
        let _359_v92;
        _359_v92 = _dafny.Map.Empty.slice().updateUnsafe(_226_v2,_321_v72);
        let _360_v93;
        _360_v93 = _dafny.Set.fromElements((_359_v92).update(_319_v70.f15, _321_v72), _359_v92);
        (_228_globalState).f6 = new BigNumber((_360_v93).length);
        _322_v73 = _322_v73;
        _module.__default.m0(_228_globalState);
        let _361_v94;
        _361_v94 = _dafny.Seq.of(_319_v70.f15);
        let _362_v95;
        _362_v95 = _dafny.Seq.of((_322_v73)[_module.__default.safeIndex(new BigNumber(626), new BigNumber((_322_v73).length))]);
        let _363_v96;
        _363_v96 = _module.D1.create_DC2(new BigNumber((_dafny.MultiSet.FromArray(_362_v95)).cardinality()));
        let _pat_let_tv12 = _321_v72;
        let _pat_let_tv13 = _229_v4;
        let _pat_let_tv14 = _321_v72;
        let _pat_let_tv15 = _228_globalState;
        _357_v91 = _dafny.Seq.Concat(_module.__default.fm6(new BigNumber(882), (_361_v94)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(new BigNumber((_357_v91).length), (_322_v73)[_module.__default.safeIndex(new BigNumber(626), new BigNumber((_322_v73).length))], (_322_v73)[_module.__default.safeIndex(new BigNumber(626), new BigNumber((_322_v73).length))], new BigNumber((_230_v5).length))).length), new BigNumber((_361_v94).length))], function (_pat_let9_0) {
          return function (_364_dt__update__tmp_h2) {
            return function (_pat_let10_0) {
              return function (_365_dt__update_hcf2_h0) {
                return _module.D1.create_DC2(_365_dt__update_hcf2_h0);
              }(_pat_let10_0);
            }(_module.__default.fm0(_pat_let_tv12.f15, _pat_let_tv13, _pat_let_tv14.f15, _pat_let_tv15));
          }(_pat_let9_0);
        }(_363_v96), _226_v2, _228_globalState), _357_v91);
      }
      let _index72 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_322_v73).length));
      let _rhs90 = (_322_v73)[_module.__default.safeIndex(new BigNumber(626), new BigNumber((_322_v73).length))];
      let _rhs91 = !(_226_v2) || (_319_v70.f15);
      let _lhs82 = _322_v73;
      let _lhs83 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_322_v73).length));
      _lhs82[_lhs83] = _rhs90;
      _226_v2 = _rhs91;
      let _366_v97;
      _366_v97 = _dafny.MultiSet.fromElements(_319_v70.f15, false);
      let _367_v98;
      let _nw60 = new _module.C0();
      _nw60.__ctor((_366_v97).IsSubsetOf(_366_v97));
      _367_v98 = _nw60;
      let _368_v99;
      _368_v99 = _dafny.Seq.of(_321_v72);
      let _369_v100;
      _369_v100 = _dafny.Seq.of(new BigNumber((_368_v99).length));
      _325_v74 = _module.__default.fm9(_module.__default.safeModuloInt(new BigNumber((_369_v100).length), (_322_v73)[_module.__default.safeIndex(new BigNumber(626), new BigNumber((_322_v73).length))]), _dafny.Seq.UnicodeFromString("xxgmqx"), _229_v4, _228_globalState);
      (_228_globalState).f10 = !(!(_367_v98.f15));
      _357_v91 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("sbo"), _357_v91), _module.__default.safeIndex(_229_v4, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("sbo"), _357_v91)).length)), new _dafny.CodePoint('e'.codePointAt(0))), _357_v91);
      process.stdout.write(_dafny.toString((_223_v0)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_223_v0)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_223_v0)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_223_v0)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_223_v0)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_223_v0)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_223_v0)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_223_v0)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_223_v0)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_223_v0)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_223_v0)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_223_v0)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_223_v0)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_223_v0)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_223_v0)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_223_v0)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_223_v0)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_223_v0)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_225_v1))[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_225_v1))[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_225_v1))[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_225_v1))[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_225_v1))[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_225_v1))[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_225_v1))[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_225_v1))[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_225_v1))[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_225_v1))[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_225_v1))[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_225_v1))[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_225_v1))[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_225_v1))[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_225_v1))[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_225_v1))[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_225_v1))[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_225_v1))[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_226_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_227_v3).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_228_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_228_globalState.f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_228_globalState.f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_228_globalState.f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_228_globalState.f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f9)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f9)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f9)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f9)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f9)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f9)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f9)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f9)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f9)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f9)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f9)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f9)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f9)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f9)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f9)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f9)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f9)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f9)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_228_globalState.f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f11)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f11)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f11)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f11)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f11)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f11)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f11)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f11)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f11)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f11)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f11)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f11)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f11)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f11)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f11)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f11)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f11)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState.f11)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_228_globalState).f12).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_228_globalState).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_229_v4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_230_v5).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(121)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_310_i7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_319_v70.f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_320_v71).dtor_cf10)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_320_v71).dtor_cf10)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_320_v71).dtor_cf10)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_320_v71).dtor_cf10)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_320_v71).dtor_cf10)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_320_v71).dtor_cf10)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_320_v71).dtor_cf10)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_320_v71).dtor_cf10)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_320_v71).dtor_cf10)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_320_v71).dtor_cf10)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_320_v71).dtor_cf10)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_320_v71).dtor_cf10)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_320_v71).dtor_cf10)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_320_v71).dtor_cf10)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_320_v71).dtor_cf10)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_320_v71).dtor_cf10)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_320_v71).dtor_cf10)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_320_v71).dtor_cf10)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_320_v71).dtor_cf11.f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_320_v71).dtor_cf12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_321_v72.f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_322_v73)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_322_v73)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_322_v73)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_322_v73)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_322_v73)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_322_v73)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_322_v73)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_322_v73)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_322_v73)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_322_v73)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_322_v73)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_322_v73)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_322_v73)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_322_v73)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_322_v73)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_322_v73)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_322_v73)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_322_v73)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_322_v73)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_322_v73)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_322_v73)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_325_v74));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_326_v75).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('f'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_327_v76).dtor_cf13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_327_v76).dtor_cf14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_327_v76).dtor_cf15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_327_v76).dtor_cf16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_327_v76).dtor_cf17, _dafny.Seq.of(_module.D3.create_DC6(new BigNumber(121)), _module.D3.create_DC6(new BigNumber(522))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_330_v77).dtor_cf6, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_357_v91).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_366_v97).equals(_dafny.MultiSet.fromElements(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_367_v98.f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_368_v99).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_369_v100, _dafny.Seq.of(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC0(cf0) {
      let $dt = new D0(0);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC0() { return this.$tag === 0; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf0 === other.cf0;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return [];
    }
    static Rtd() {
      return class {
        static get Default() {
          return D0.Default();
        }
      };
    }
  }

  $module.D1 = class D1 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC2(cf2) {
      let $dt = new D1(0);
      $dt.cf2 = cf2;
      return $dt;
    }
    static create_DC1(cf1) {
      let $dt = new D1(1);
      $dt.cf1 = cf1;
      return $dt;
    }
    get is_DC2() { return this.$tag === 0; }
    get is_DC1() { return this.$tag === 1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf1() { return this.cf1; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf2) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf2, other.cf2);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf1, other.cf1);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC2(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D1.Default();
        }
      };
    }
  }

  $module.D2 = class D2 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC4(cf4, cf5) {
      let $dt = new D2(0);
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC3(cf3) {
      let $dt = new D2(1);
      $dt.cf3 = cf3;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf3() { return this.cf3; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC4" + "(" + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC3" + "(" + this.cf3.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf4, other.cf4) && _dafny.areEqual(this.cf5, other.cf5);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf3, other.cf3);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC4(_dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC6(cf7) {
      let $dt = new D3(0);
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC5(cf6) {
      let $dt = new D3(1);
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC7(cf8) {
      let $dt = new D3(2);
      $dt.cf8 = cf8;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC7() { return this.$tag === 2; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf8() { return this.cf8; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC6" + "(" + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC5" + "(" + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC7" + "(" + _dafny.toString(this.cf8) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf7, other.cf7);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf6, other.cf6);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf8, other.cf8);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC6(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D3.Default();
        }
      };
    }
  }

  $module.D4 = class D4 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC9(cf10, cf11, cf12) {
      let $dt = new D4(0);
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC10(cf13, cf14, cf15, cf16, cf17) {
      let $dt = new D4(1);
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC11(cf18) {
      let $dt = new D4(2);
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC8(cf9) {
      let $dt = new D4(3);
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC12(cf19) {
      let $dt = new D4(4);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get is_DC11() { return this.$tag === 2; }
    get is_DC8() { return this.$tag === 3; }
    get is_DC12() { return this.$tag === 4; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC9" + "(" + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 3) {
        return "D4.DC8" + "(" + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 4) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf19) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf10 === other.cf10 && this.cf11 === other.cf11 && this.cf12 === other.cf12;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf13, other.cf13) && _dafny.areEqual(this.cf14, other.cf14) && _dafny.areEqual(this.cf15, other.cf15) && this.cf16 === other.cf16 && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf18, other.cf18);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf9 === other.cf9;
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf19, other.cf19);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC9([], null, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D4.Default();
        }
      };
    }
  }

  $module.D5 = class D5 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC14() {
      let $dt = new D5(0);
      return $dt;
    }
    static create_DC13(cf20) {
      let $dt = new D5(1);
      $dt.cf20 = cf20;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get dtor_cf20() { return this.cf20; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC14";
      } else if (this.$tag === 1) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf20) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf20, other.cf20);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC14();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC16(cf22, cf23, cf24) {
      let $dt = new D6(0);
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC15(cf21) {
      let $dt = new D6(1);
      $dt.cf21 = cf21;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf21() { return this.cf21; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC15" + "(" + _dafny.toString(this.cf21) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf22 === other.cf22 && this.cf23 === other.cf23 && _dafny.areEqual(this.cf24, other.cf24);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf21 === other.cf21;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC16(false, null, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D6.Default();
        }
      };
    }
  }

  $module.D7 = class D7 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC18() {
      let $dt = new D7(0);
      return $dt;
    }
    static create_DC17(cf25) {
      let $dt = new D7(1);
      $dt.cf25 = cf25;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get is_DC17() { return this.$tag === 1; }
    get dtor_cf25() { return this.cf25; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC18";
      } else if (this.$tag === 1) {
        return "D7.DC17" + "(" + _dafny.toString(this.cf25) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf25 === other.cf25;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC18();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D7.Default();
        }
      };
    }
  }

  $module.D8 = class D8 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC20(cf27, cf28, cf29) {
      let $dt = new D8(0);
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC21(cf30) {
      let $dt = new D8(1);
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC19(cf26) {
      let $dt = new D8(2);
      $dt.cf26 = cf26;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get is_DC21() { return this.$tag === 1; }
    get is_DC19() { return this.$tag === 2; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf26() { return this.cf26; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC20" + "(" + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC21" + "(" + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC19" + "(" + _dafny.toString(this.cf26) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf27, other.cf27) && _dafny.areEqual(this.cf28, other.cf28) && _dafny.areEqual(this.cf29, other.cf29);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf30 === other.cf30;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf26, other.cf26);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC20(_dafny.ZERO, _dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D8.Default();
        }
      };
    }
  }

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f0 = _dafny.ZERO;
      this.f1 = false;
      this.f4 = _dafny.ZERO;
      this.f6 = _dafny.ZERO;
      this.f8 = false;
      this.f9 = [];
      this.f10 = false;
      this.f11 = [];
      this._f2 = false;
      this._f3 = _dafny.ZERO;
      this._f5 = _dafny.ZERO;
      this._f7 = false;
      this._f12 = _dafny.Map.Empty;
      this._f13 = _dafny.ZERO;
      this._f14 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14) {
      let _this = this;
      (_this).f0 = f0;
      (_this).f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this).f4 = f4;
      (_this)._f5 = f5;
      (_this).f6 = f6;
      (_this)._f7 = f7;
      (_this).f8 = f8;
      (_this).f9 = f9;
      (_this).f10 = f10;
      (_this).f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      (_this)._f14 = f14;
      return;
    }
    get f2() {
      let _this = this;
      return _this._f2;
    };
    get f3() {
      let _this = this;
      return _this._f3;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this.f15 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f15) {
      let _this = this;
      (_this).f15 = f15;
      return;
    }
    fm1(p0, p1, globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f15,new BigNumber(-178))).length)).minus(new BigNumber((_dafny.Seq.UnicodeFromString("dbuyqaqah")).length)),_this.f15);
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
