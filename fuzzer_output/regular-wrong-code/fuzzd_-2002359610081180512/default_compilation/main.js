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
    static fm0(globalState) {
      return _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(false))).length), ((!(!(false))) ? (new BigNumber((_dafny.Seq.UnicodeFromString("rakhhgk")).length)) : (new BigNumber((_dafny.Seq.UnicodeFromString("ry")).length))));
    };
    static fm1(p0, p1, globalState) {
      return _dafny.Seq.contains(_dafny.Seq.UnicodeFromString("nkmmphev"), new _dafny.CodePoint('d'.codePointAt(0)));
    };
    static fm6(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(110)), function (_0_i0) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      }), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("olpkpd"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(775)), function (_1_i1) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      })));
    };
    static fm7(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(981)), function (_2_i0) {
        return _module.D1.create_DC2(new BigNumber(61), true, new BigNumber(-978), new BigNumber(52));
      }), _dafny.Seq.Concat(_dafny.Seq.of(_module.D1.create_DC2(new BigNumber((_dafny.Set.fromElements(false)).length), true, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(false)).length))).length), new BigNumber((_dafny.Seq.UnicodeFromString("ropqb")).length)), _module.D1.create_DC2(new BigNumber(422), false, new BigNumber(-728), new BigNumber(-268)), _module.D1.create_DC2(new BigNumber(157), false, new BigNumber(267), new BigNumber(-998))), _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(179), true, new BigNumber(207), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-541)), function (_3_i1) {
  return new BigNumber((_dafny.Seq.UnicodeFromString("xv")).length);
})).length)))));
    };
    static fm9(p0, p1, p2, globalState) {
      if (((!(true)) ? (false) : (false))) {
        return _dafny.Seq.of(new BigNumber(-925));
      } else {
        return _dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(420), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-368))).length), (_dafny.ZERO).minus(new BigNumber(-659)), new BigNumber(-693), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(548))).length))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(false, true, true)).cardinality())),_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),!(true))).length))),new BigNumber(916)))).length));
      }
    };
    static fm10(p0, globalState) {
      return (_dafny.Set.fromElements(true)).Union(_dafny.Set.fromElements(!(false), !(true)));
    };
    static fm11(p0, globalState) {
      let _source0 = ((true) ? (_module.D16.create_DC45(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true)).length),new BigNumber(142)))) : (_module.D16.create_DC45(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(87),false)).length),new BigNumber(220)))));
      if (_source0.is_DC46) {
        let _4___mcc_h0 = (_source0).cf79;
        let _5_cf79 = _4___mcc_h0;
        return new _dafny.CodePoint('d'.codePointAt(0));
      } else if (_source0.is_DC45) {
        let _6___mcc_h1 = (_source0).cf78;
        let _7_cf78 = _6___mcc_h1;
        return new _dafny.CodePoint('g'.codePointAt(0));
      } else {
        let _8___mcc_h2 = (_source0).cf80;
        let _9_cf80 = _8___mcc_h2;
        return new _dafny.CodePoint('a'.codePointAt(0));
      }
    };
    static fm12(p0, p1, p2, p3, globalState) {
      return _dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("gighne")).length), new BigNumber(946));
    };
    static fm13(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(756),_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(289)), function (_10_i0) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      })).length))).length))))).Merge(function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of (_dafny.Seq.of(new BigNumber(855), new BigNumber(-260), (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(310),new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("dnvebqbov")).length))).length))).length))), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(509),new BigNumber(851))).length), new BigNumber(965))).Elements) {
          let _11_v0 = _compr_0;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(855), new BigNumber(-260), (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(310),new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("dnvebqbov")).length))).length))).length))), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(509),new BigNumber(851))).length), new BigNumber(965)), _11_v0)) {
            _coll0.push([(_11_v0).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(!(!(false)))).length))).length)),_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("n")).length))]);
          }
        }
        return _coll0;
      }());
    };
    static fm14(p0, p1, p2, globalState) {
      return _dafny.Set.fromElements(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.of(false)).length), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()))), new BigNumber((_dafny.Seq.UnicodeFromString("ahpb")).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(202)), function (_12_i0) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      })).length)), (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(!(false), false))).cardinality())))).Intersect(_dafny.MultiSet.fromElements(new BigNumber(904)))).cardinality()))));
    };
    static fm15(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)))).length),true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(816)), function (_13_i0) {
        return new BigNumber(627);
      })).length),false));
    };
    static fm16(p0, p1, globalState) {
      return (_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(814), (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-505))), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(704),new BigNumber(455))).length)), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("gf"), _dafny.Seq.UnicodeFromString("s"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-308)), function (_14_i0) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      }))).cardinality())), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),false)).length)))).Intersect(_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(893))).length), new BigNumber(646))));
    };
    static fm17(p0, p1, p2, p3, globalState) {
      return function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of (_dafny.Seq.of(_module.D3.create_DC8())).Elements) {
          let _15_v0 = _compr_1;
          if (_dafny.Seq.contains(_dafny.Seq.of(_module.D3.create_DC8()), _15_v0)) {
            _coll1.push([_15_v0,((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(true)).length))).plus(new BigNumber(558))]);
          }
        }
        return _coll1;
      }();
    };
    static fm18(p0, p1, p2, globalState) {
      return _module.D1.create_DC1(new BigNumber(230));
    };
    static fm19(p0, globalState) {
      return (((false) ? (_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("gxerujrc"))) : (_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("nwiu"), _dafny.Seq.UnicodeFromString("fghlrjlu"))))).Difference(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("efts")));
    };
    static fm20(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(false));
    };
    static fm21(p0, p1, globalState) {
      return ((_dafny.MultiSet.fromElements(true)).Intersect(_dafny.MultiSet.fromElements(true))).Union((_dafny.MultiSet.fromElements(true)).Difference(_dafny.MultiSet.fromElements(false)));
    };
    static fm22(globalState) {
      return (function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber(558), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(417)), function (_16_i0) {
          return new _dafny.CodePoint('r'.codePointAt(0));
        })).length))).length)),new BigNumber((_dafny.Seq.of(false, false)).length))).Keys.Elements) {
          let _17_v0 = _compr_2;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber(558), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(417)), function (_16_i0) {
            return new _dafny.CodePoint('r'.codePointAt(0));
          })).length))).length)),new BigNumber((_dafny.Seq.of(false, false)).length))).contains(_17_v0)) {
            _coll2.push([_17_v0,!(true)]);
          }
        }
        return _coll2;
      }()).Merge(function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(57),true)).length), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false)).length))).length)),new BigNumber(656))).Keys.Elements) {
          let _18_v1 = _compr_3;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(57),true)).length), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false)).length))).length)),new BigNumber(656))).contains(_18_v1)) {
            _coll3.push([_18_v1,true]);
          }
        }
        return _coll3;
      }());
    };
    static fm23(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(false)).length))).Merge((_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber(911))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(559))));
    };
    static fm24(p0, p1, p2, p3, globalState) {
      return (_dafny.Set.fromElements(!(true))).Intersect(_dafny.Set.fromElements(true, false, false, false));
    };
    static fm25(p0, p1, p2, p3, globalState) {
      return _module.D1.create_DC2(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(957))).length), (!(false)) && (false), (new BigNumber(568)).minus(new BigNumber(800)), new BigNumber(154));
    };
    static fm26(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_module.D1.create_DC2(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('j'.codePointAt(0)))).length), false, new BigNumber((_dafny.Set.fromElements(true, true, !(false), false, true)).length), new BigNumber(15)), _module.D1.create_DC2(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("xc"),new BigNumber((_dafny.Set.fromElements(true)).length))).length), !(true), new BigNumber(679), new BigNumber(163))), _dafny.Seq.of(_module.D1.create_DC2(new BigNumber((_dafny.Set.fromElements(function () {
  let _coll4 = new _dafny.Set();
  for (const _compr_4 of (_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(837))).length))).Elements) {
    let _19_v0 = _compr_4;
    if ((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(837))).length))).contains(_19_v0)) {
      _coll4.add(_module.__default.safeModuloInt(_19_v0, new BigNumber(866)));
    }
  }
  return _coll4;
}(), _dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("xydlmxbg")).length)))).length), false, new BigNumber(683), new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("mxrqqpsh"))).cardinality())), _module.D1.create_DC2(new BigNumber((_dafny.Seq.of(new BigNumber(31), new BigNumber(106))).length), true, new BigNumber(860), new BigNumber(-392)), _module.D1.create_DC2(new BigNumber(116), true, new BigNumber((_dafny.Seq.UnicodeFromString("mn")).length), new BigNumber(-325)))), _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(627), !(false), new BigNumber(837), new BigNumber(429))));
    };
    static fm27(globalState) {
      return _module.D3.create_DC10();
    };
    static fm28(p0, p1, globalState) {
      return _dafny.Seq.UnicodeFromString("wfu");
    };
    static fm29(p0, p1, globalState) {
      return _dafny.MultiSet.fromElements(_module.__default.safeModuloInt(new BigNumber(897), new BigNumber(-821)));
    };
    static fm30(globalState) {
      if (true) {
        return (_dafny.Map.Empty.slice().updateUnsafe(false,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,false));
      } else {
        return _dafny.Map.Empty.slice().updateUnsafe(true,!(false));
      }
    };
    static fm31(globalState) {
      let _source1 = _module.D12.create_DC32(new BigNumber(-862), true, new BigNumber((_dafny.Seq.UnicodeFromString("knta")).length));
      if (_source1.is_DC32) {
        let _20___mcc_h0 = (_source1).cf53;
        let _21___mcc_h1 = (_source1).cf54;
        let _22___mcc_h2 = (_source1).cf55;
        let _23_cf55 = _22___mcc_h2;
        let _24_cf54 = _21___mcc_h1;
        let _25_cf53 = _20___mcc_h0;
        return new _dafny.CodePoint('h'.codePointAt(0));
      } else if (_source1.is_DC31) {
        let _26___mcc_h3 = (_source1).cf52;
        let _27_cf52 = _26___mcc_h3;
        return new _dafny.CodePoint('p'.codePointAt(0));
      } else {
        let _28___mcc_h4 = (_source1).cf56;
        let _29_cf56 = _28___mcc_h4;
        return new _dafny.CodePoint('c'.codePointAt(0));
      }
    };
    static fm34(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("opsm")), _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("tmbalwqgj"), _dafny.Seq.UnicodeFromString("uetnpshyt")), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("s"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-573)), function (_30_i0) {
        return new _dafny.CodePoint('w'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("s"))));
    };
    static fm35(globalState) {
      if (((true) ? (true) : (!(!(true))))) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('b'.codePointAt(0));
      }
    };
    static fm36(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(901),new BigNumber(-213))).length)),new BigNumber(-297))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-323)),new BigNumber((_dafny.Seq.of(_dafny.Seq.UnicodeFromString("nxungteu"))).length)));
    };
    static fm37(p0, p1, globalState) {
      return _dafny.Set.fromElements(new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)));
    };
    static fm38(p0, p1, globalState) {
      return _module.D1.create_DC2((_dafny.ZERO).minus(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-290))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(!(!(!(false)))),new BigNumber(519)))).length)), true, new BigNumber((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("pntnad"), _dafny.Seq.UnicodeFromString("vo"))).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("qvbi")).length), new BigNumber(448))).cardinality()));
    };
    static fm39(globalState) {
      return (_dafny.Set.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.of((_module.D12.create_DC32(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Map.Empty.slice().updateUnsafe(false,true))).length), false, new BigNumber((function () {
  let _coll5 = new _dafny.Set();
  for (const _compr_5 of _dafny.IntegerRange(new BigNumber(37), new BigNumber(25))) {
    let _31_v0 = _compr_5;
    if (((new BigNumber(37)).isLessThanOrEqualTo(_31_v0)) && ((_31_v0).isLessThan(new BigNumber(25)))) {
      _coll5.add(_module.__default.safeDivisionInt(_31_v0, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(207))).cardinality())));
    }
  }
  return _coll5;
}()).length))).dtor_cf54, false)))).Intersect(_dafny.Set.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.of(true))));
    };
    static fm40(globalState) {
      return _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(213)), function (_32_i0) {
        return new _dafny.CodePoint('l'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("rp")));
    };
    static fm41(p0, p1, p2, globalState) {
      return _dafny.Seq.of(false);
    };
    static fm42(p0, p1, p2, globalState) {
      return (_dafny.MultiSet.fromElements(false, true, false));
    };
    static fm43(p0, p1, p2, globalState) {
      return (((!(false)) ? (_dafny.MultiSet.fromElements(_dafny.Set.fromElements(true, false, !(false)))) : (_dafny.MultiSet.fromElements(_dafny.Set.fromElements(false, true))))).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Set.fromElements(true, true, false, true), _dafny.Set.fromElements(!(false)), _dafny.Set.fromElements(false, false, false, !(false)), _dafny.Set.fromElements(true), _dafny.Set.fromElements(!(!(false))))));
    };
    static fm44(globalState) {
      let _source2 = _module.D14.create_DC38(new _dafny.CodePoint('p'.codePointAt(0)), !(true), _dafny.Seq.of((_module.D7.create_DC20(true, new _dafny.CodePoint('k'.codePointAt(0)))).dtor_cf36));
      if (_source2.is_DC38) {
        let _33___mcc_h0 = (_source2).cf67;
        let _34___mcc_h1 = (_source2).cf68;
        let _35___mcc_h2 = (_source2).cf69;
        let _36_cf69 = _35___mcc_h2;
        let _37_cf68 = _34___mcc_h1;
        let _38_cf67 = _33___mcc_h0;
        return _module.D7.create_DC20(false, _38_cf67);
      } else if (_source2.is_DC39) {
        let _39___mcc_h3 = (_source2).cf70;
        let _40___mcc_h4 = (_source2).cf71;
        let _41_cf71 = _40___mcc_h4;
        let _42_cf70 = _39___mcc_h3;
        if (false) {
          return _module.D7.create_DC20(true, new _dafny.CodePoint('h'.codePointAt(0)));
        } else {
          return _module.D7.create_DC20(false, new _dafny.CodePoint('t'.codePointAt(0)));
        }
      } else {
        let _43___mcc_h5 = (_source2).cf66;
        let _44_cf66 = _43___mcc_h5;
        return _module.D7.create_DC20(false, new _dafny.CodePoint('g'.codePointAt(0)));
      }
    };
    static fm45(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_module.D1.create_DC3(true, _dafny.Seq.UnicodeFromString("dwhk"), new BigNumber(-798))), _dafny.Seq.of(_module.D1.create_DC3(!(false), _dafny.Seq.UnicodeFromString("cc"), (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("nsv")).length)))))), _dafny.Seq.of(_module.D1.create_DC3(!(true), (_module.D13.create_DC36(_dafny.Seq.Create(_module.__default.abs(_dafny.ZERO), function (_45_i0) {
  return new _dafny.CodePoint('s'.codePointAt(0));
}), new BigNumber(369), new BigNumber((_dafny.Seq.of(false)).length), new BigNumber((_dafny.MultiSet.fromElements(true, false)).cardinality()))).dtor_cf62, new BigNumber(246)), _module.D1.create_DC3(true, _dafny.Seq.UnicodeFromString("e"), new BigNumber(89))));
    };
    static fm46(p0, p1, p2, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
        let _coll6 = new _dafny.Map();
        for (const _compr_6 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(703)), function (_46_i0) {
          return new BigNumber(945);
        })).Elements) {
          let _47_v0 = _compr_6;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(703)), function (_46_i0) {
            return new BigNumber(945);
          }), _47_v0)) {
            _coll6.push([_module.__default.safeModuloInt(_47_v0, new BigNumber((_dafny.Seq.of(new BigNumber(655), new BigNumber(235))).length)),(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("bwex")).length),new BigNumber(493))).length))]);
          }
        }
        return _coll6;
      }()).length),_dafny.Seq.Create(_module.__default.abs(new BigNumber(673)), function (_48_i1) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      }))).Merge(function () {
        let _coll7 = new _dafny.Map();
        for (const _compr_7 of _dafny.IntegerRange(new BigNumber(706), new BigNumber(867))) {
          let _49_v1 = _compr_7;
          if (((new BigNumber(706)).isLessThanOrEqualTo(_49_v1)) && ((_49_v1).isLessThan(new BigNumber(867)))) {
            _coll7.push([(_49_v1).multipliedBy(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.of(new BigNumber(941)), _dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(!(true)))).cardinality()), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(420)), function (_50_i2) {
              return new _dafny.CodePoint('r'.codePointAt(0));
            })).length)), _dafny.Seq.of(new BigNumber(284))))).cardinality())),_dafny.Seq.UnicodeFromString("pkrufmho")]);
          }
        }
        return _coll7;
      }())).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(true, false, false)).cardinality()),_dafny.Seq.Create(_module.__default.abs(new BigNumber(492)), function (_51_i3) {
        return new _dafny.CodePoint('i'.codePointAt(0));
      })));
    };
    static fm47(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(936)),new BigNumber((_dafny.Seq.UnicodeFromString("ti")).length))).Merge(function () {
        let _coll8 = new _dafny.Map();
        for (const _compr_8 of (_dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(true, false)).length))).cardinality())))).Elements) {
          let _52_v0 = _compr_8;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(true, false)).length))).cardinality()))), _52_v0)) {
            _coll8.push([_52_v0,new BigNumber(180)]);
          }
        }
        return _coll8;
      }())).Merge((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(347), new BigNumber((_dafny.Set.fromElements(new BigNumber(-530))).length)),new BigNumber((_dafny.Set.fromElements(_dafny.MultiSet.fromElements(true))).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(100)), function (_53_i0) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      })).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((function () {
        let _coll9 = new _dafny.Set();
        for (const _compr_9 of (_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(531),true)).length))).Elements) {
          let _54_v1 = _compr_9;
          if ((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(531),true)).length))).contains(_54_v1)) {
            _coll9.add((_54_v1).multipliedBy(new BigNumber(801)));
          }
        }
        return _coll9;
      }()).length))).cardinality()),!(false))).length)),new BigNumber(660))));
    };
    static fm48(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("ks")).length),(_dafny.ZERO).minus((new BigNumber((_dafny.Set.fromElements(new BigNumber(-433), new BigNumber(647), new BigNumber(919))).length)).minus(new BigNumber(852))));
    };
    static m0(p0, globalState) {
      let _55_v1;
      _55_v1 = new BigNumber(-50);
      let _56_v2;
      _56_v2 = _dafny.Map.Empty.slice().updateUnsafe(_55_v1,_55_v1);
      let _57_v3;
      _57_v3 = _dafny.Map.Empty.slice().updateUnsafe(function () {
        let _coll10 = new _dafny.Map();
        for (const _compr_10 of (_dafny.Seq.of(_55_v1)).Elements) {
          let _58_v0 = _compr_10;
          if (_dafny.Seq.contains(_dafny.Seq.of(_55_v1), _58_v0)) {
            _coll10.push([(_58_v0).multipliedBy(_55_v1),p0]);
          }
        }
        return _coll10;
      }(),(_56_v2).Merge((_56_v2).update(_55_v1, _55_v1)));
      let _59_v4;
      _59_v4 = _dafny.MultiSet.fromElements(_55_v1);
      let _60_v5;
      _60_v5 = _dafny.Map.Empty.slice().updateUnsafe(_55_v1,_module.__default.fm1(_59_v4, _55_v1, globalState));
      let _61_v6;
      _61_v6 = _module.D18.create_DC52(_57_v3);
      _57_v3 = ((_dafny.Map.Empty.slice().updateUnsafe(_60_v5,_56_v2)).Merge((_61_v6).dtor_cf88)).Merge(((_57_v3).update(_60_v5, _56_v2)).Merge(_57_v3));
      _55_v1 = _55_v1;
      if (p0) {
        let _62_v7;
        _62_v7 = _dafny.Seq.UnicodeFromString("dlxanf");
        _55_v1 = _module.__default.safeDivisionInt((((_56_v2).contains((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(p0, p0)).cardinality())))) ? ((_56_v2).get((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(p0, p0)).cardinality())))) : (_55_v1)), new BigNumber((_62_v7).length));
        let _63_v8;
        _63_v8 = _dafny.Seq.of(_55_v1);
        _55_v1 = (new BigNumber((_dafny.MultiSet.FromArray(_63_v8)).cardinality())).multipliedBy(_55_v1);
        let _64_v9;
        _64_v9 = _dafny.Map.Empty.slice().updateUnsafe(p0,!(_55_v1).isEqualTo(_55_v1));
        (globalState).f2 = (((_64_v9).contains(p0)) ? ((_64_v9).get(p0)) : (_module.__default.fm1(_59_v4, _55_v1, globalState)));
        let _65_v10;
        _65_v10 = _dafny.Seq.of(p0);
        _65_v10 = _65_v10;
        let _66_v11;
        _66_v11 = new _dafny.CodePoint('f'.codePointAt(0));
        let _67_v12;
        let _nw0 = new _module.C1();
        _nw0.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(644)), ((_68_p0, _69_v4, _70_v1) => function (_71_i0) {
          return _module.D1.create_DC2((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_68_p0,_68_p0)).length)), _68_p0, new BigNumber((_dafny.Set.fromElements(new BigNumber((_69_v4).cardinality()))).length), _70_v1);
        })(p0, _59_v4, _55_v1)));
        _67_v12 = _nw0;
        let _72_v13;
        _72_v13 = _module.D17.create_DC50(_66_v11, p0, _62_v7, _55_v1, _67_v12);
        _62_v7 = _dafny.Seq.update((_72_v13).dtor_cf84, _module.__default.safeIndex(_55_v1, new BigNumber(((_72_v13).dtor_cf84).length)), _66_v11);
      } else {
        let _73_v14;
        _73_v14 = _module.D1.create_DC2(_55_v1, (_55_v1).isLessThanOrEqualTo(_55_v1), _55_v1, _55_v1);
        let _74_v15;
        _74_v15 = _dafny.MultiSet.fromElements(p0, p0);
        let _75_v16;
        let _nw1 = new _module.C4();
        _nw1.__ctor(p0);
        _75_v16 = _nw1;
        let _76_v17;
        _76_v17 = _dafny.Map.Empty.slice().updateUnsafe(_75_v16,p0);
        _73_v14 = _module.D1.create_DC2((_55_v1).multipliedBy(_55_v1), p0, _55_v1, (new BigNumber((_74_v15).cardinality())).plus(new BigNumber((_76_v17).length)));
        (globalState).f2 = _module.__default.fm1(_59_v4, _module.__default.safeModuloInt(_55_v1, _55_v1), globalState);
        let _77_v18;
        _77_v18 = _dafny.Seq.of(_73_v14);
        let _78_v19;
        let _nw2 = new _module.C1();
        _nw2.__ctor(_77_v18);
        _78_v19 = _nw2;
        let _79_v20;
        _79_v20 = _dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(p0, p0, p0, p0)).cardinality()), _55_v1, new BigNumber(419), new BigNumber((_dafny.Seq.update(_dafny.Seq.of(p0), _module.__default.safeIndex(_55_v1, new BigNumber((_dafny.Seq.of(p0)).length)), true)).length));
        let _80_v21;
        _80_v21 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsProperPrefixOf(_79_v20, _79_v20),_55_v1);
        let _81_v22;
        _81_v22 = _dafny.Seq.UnicodeFromString("kqwwfaqw");
        let _82_v23;
        let _nw3 = new _module.C5();
        _nw3.__ctor(_55_v1, p0);
        _82_v23 = _nw3;
        let _83_v24;
        _83_v24 = _dafny.Seq.of(_82_v23, _82_v23, _82_v23);
        let _rhs0 = p0;
        let _rhs1 = _80_v21;
        let _rhs2 = p0;
        let _rhs3 = _81_v22;
        let _rhs4 = new BigNumber((_83_v24).length);
        let _lhs0 = globalState;
        let _lhs1 = globalState;
        _lhs0.f2 = _rhs0;
        _80_v21 = _rhs1;
        _lhs1.f2 = _rhs2;
        _81_v22 = _rhs3;
        _55_v1 = _rhs4;
        let _84_v25;
        _84_v25 = _dafny.Set.fromElements(_79_v20);
        let _85_v26;
        _85_v26 = _dafny.Seq.of(p0, p0, false, p0, p0);
        let _86_v27;
        _86_v27 = _dafny.Seq.of(!(p0), (_85_v26)[_module.__default.safeIndex((_82_v23).f10, new BigNumber((_85_v26).length))], true);
        let _87_v28;
        _87_v28 = _module.D11.create_DC30();
        if (!((_82_v23).fm5((_84_v25).Difference(_84_v25), _dafny.Map.Empty.slice().updateUnsafe((_82_v23).fm4(p0, p0, p0, p0, globalState),_85_v26), new BigNumber((((p0) ? (_dafny.Seq.of(_87_v28)) : (_dafny.Seq.of(_87_v28)))).length), globalState))) {
          let _88_v29;
          _88_v29 = new _dafny.CodePoint('x'.codePointAt(0));
          let _89_v30;
          _89_v30 = _dafny.Map.Empty.slice().updateUnsafe(p0,_88_v29);
          let _90_v31;
          let _nw4 = Array((new BigNumber(23)).toNumber());
          _nw4[(_dafny.ZERO).toNumber()] = _88_v29;
          _nw4[(_dafny.ONE).toNumber()] = (((_89_v30).contains(p0)) ? ((_89_v30).get(p0)) : (_88_v29));
          _nw4[(new BigNumber(2)).toNumber()] = _88_v29;
          _nw4[(new BigNumber(3)).toNumber()] = _88_v29;
          _nw4[(new BigNumber(4)).toNumber()] = _88_v29;
          _nw4[(new BigNumber(5)).toNumber()] = ((!(!(!(p0)))) ? (_88_v29) : (_88_v29));
          _nw4[(new BigNumber(6)).toNumber()] = _88_v29;
          _nw4[(new BigNumber(7)).toNumber()] = _88_v29;
          _nw4[(new BigNumber(8)).toNumber()] = _88_v29;
          _nw4[(new BigNumber(9)).toNumber()] = _88_v29;
          _nw4[(new BigNumber(10)).toNumber()] = _88_v29;
          _nw4[(new BigNumber(11)).toNumber()] = _88_v29;
          _nw4[(new BigNumber(12)).toNumber()] = _88_v29;
          _nw4[(new BigNumber(13)).toNumber()] = ((p0) ? (_88_v29) : (_88_v29));
          _nw4[(new BigNumber(14)).toNumber()] = _88_v29;
          _nw4[(new BigNumber(15)).toNumber()] = new _dafny.CodePoint('s'.codePointAt(0));
          _nw4[(new BigNumber(16)).toNumber()] = new _dafny.CodePoint('d'.codePointAt(0));
          _nw4[(new BigNumber(17)).toNumber()] = _88_v29;
          _nw4[(new BigNumber(18)).toNumber()] = new _dafny.CodePoint('i'.codePointAt(0));
          _nw4[(new BigNumber(19)).toNumber()] = new _dafny.CodePoint('r'.codePointAt(0));
          _nw4[(new BigNumber(20)).toNumber()] = _88_v29;
          _nw4[(new BigNumber(21)).toNumber()] = _88_v29;
          _nw4[(new BigNumber(22)).toNumber()] = (_81_v22)[_module.__default.safeIndex((_82_v23).f10, new BigNumber((_81_v22).length))];
          _90_v31 = _nw4;
          let _index0 = _module.__default.safeIndex(new BigNumber(554), new BigNumber((_90_v31).length));
          (_90_v31)[_index0] = _88_v29;
          let _index1 = _module.__default.safeIndex(new BigNumber(554), new BigNumber((_90_v31).length));
          (_90_v31)[_index1] = _88_v29;
          _79_v20 = _79_v20;
          let _91_v32;
          let _nw5 = new _module.C2();
          _nw5.__ctor((_82_v23).f10, !(p0) || (p0));
          _91_v32 = _nw5;
          let _92_v33;
          let _nw6 = Array((new BigNumber(3)).toNumber()).fill(false);
          _92_v33 = _nw6;
          let _index2 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_92_v33).length));
          (_92_v33)[_index2] = ((_82_v23).f10).isLessThanOrEqualTo((_79_v20)[_module.__default.safeIndex((_82_v23).f10, new BigNumber((_79_v20).length))]);
          let _index3 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_92_v33).length));
          (_92_v33)[_index3] = false;
          (globalState).f2 = _dafny.areEqual(_dafny.Seq.UnicodeFromString("jb"), _module.__default.fm6(_86_v27, globalState));
        } else {
          let _93_v34;
          let _nw7 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
          _93_v34 = _nw7;
          let _94_v35;
          _94_v35 = _dafny.Set.fromElements(_93_v34, _93_v34);
          let _95_v36;
          _95_v36 = new _dafny.CodePoint('m'.codePointAt(0));
          let _96_v37;
          _96_v37 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements(_93_v34)).Union(_94_v35),_dafny.Seq.update(_81_v22, _module.__default.safeIndex(_55_v1, new BigNumber((_81_v22).length)), _95_v36));
          _96_v37 = (_96_v37).update((_94_v35).Intersect(_94_v35), _81_v22);
          let _97_v38;
          let _nw8 = new _module.C0();
          _nw8.__ctor(_84_v25, ((_82_v23).f10).minus((_82_v23).f10));
          _97_v38 = _nw8;
          let _98_v39;
          _98_v39 = _dafny.Seq.of((_93_v34));
          let _rhs5 = p0;
          let _rhs6 = (_98_v39)[_module.__default.safeIndex((_82_v23).f10, new BigNumber((_98_v39).length))];
          let _rhs7 = _79_v20;
          let _rhs8 = _97_v38;
          let _rhs9 = new BigNumber(-96);
          let _lhs2 = globalState;
          let _lhs3 = _97_v38;
          _lhs2.f2 = _rhs5;
          _93_v34 = _rhs6;
          _79_v20 = _rhs7;
          _97_v38 = _rhs8;
          _lhs3.f12 = _rhs9;
          (globalState).f2 = p0;
          let _rhs10 = ((_82_v23).f10).minus((new BigNumber(603)).plus(_97_v38.f12));
          let _rhs11 = _55_v1;
          let _lhs4 = _97_v38;
          _55_v1 = _rhs10;
          _lhs4.f12 = _rhs11;
          _55_v1 = (_dafny.ZERO).minus((_82_v23).f10);
        }
      }
      let _99_v40;
      _99_v40 = _dafny.Set.fromElements(new BigNumber(-54), _55_v1, _55_v1, _55_v1, _55_v1);
      let _100_i1;
      _100_i1 = _dafny.ZERO;
      L0: {
        while (_module.__default.fm1(_59_v4, _module.__default.safeModuloInt(new BigNumber((_99_v40).length), _55_v1), globalState)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_100_i1)) {
              break L0;
            }
            _100_i1 = (_100_i1).plus(_dafny.ONE);
            let _101_v41;
            let _init0 = ((_102_p0) => function (_103_i2) {
              return _102_p0;
            })(p0);
            let _nw9 = Array((new BigNumber(15)).toNumber());
            for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw9.length); _i0_0++) {
              _nw9[_i0_0] = _init0(new BigNumber(_i0_0));
            }
            _101_v41 = _nw9;
            let _index4 = _module.__default.safeIndex(new BigNumber(901), new BigNumber((_101_v41).length));
            (_101_v41)[_index4] = p0;
            let _index5 = _module.__default.safeIndex(new BigNumber(901), new BigNumber((_101_v41).length));
            (_101_v41)[_index5] = !(p0);
            (globalState).f2 = p0;
            _55_v1 = (_55_v1).plus(_55_v1);
            let _index6 = _module.__default.safeIndex(new BigNumber(901), new BigNumber((_101_v41).length));
            (_101_v41)[_index6] = p0;
          }
        }
      }
      if (!(p0) || (p0)) {
        let _104_v42;
        _104_v42 = _99_v40;
        let _105_v44;
        let _nw10 = Array((new BigNumber(4)).toNumber());
        _nw10[(_dafny.ZERO).toNumber()] = _99_v40;
        _nw10[(_dafny.ONE).toNumber()] = (_104_v42);
        _nw10[(new BigNumber(2)).toNumber()] = _99_v40;
        _nw10[(new BigNumber(3)).toNumber()] = function () {
          let _coll11 = new _dafny.Set();
          for (const _compr_11 of (_56_v2).Keys.Elements) {
            let _106_v43 = _compr_11;
            if ((_56_v2).contains(_106_v43)) {
              _coll11.add((_106_v43).multipliedBy(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber(92), new BigNumber((_dafny.MultiSet.fromElements(false, false)).cardinality()))).length), new BigNumber((_dafny.Seq.UnicodeFromString("oaq")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(false, false)).length), new BigNumber((_dafny.Set.fromElements(new BigNumber(125))).length))).length)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(290))).length))).length), new BigNumber(623), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length))).length)));
            }
          }
          return _coll11;
        }();
        _105_v44 = _nw10;
        let _107_v45;
        let _nw11 = new _module.C3();
        _nw11.__ctor(p0, _105_v44, !(p0));
        _107_v45 = _nw11;
        _55_v1 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.fm0(globalState)));
        _55_v1 = (_55_v1).multipliedBy(_55_v1);
        (globalState).f2 = ((p0) ? (!(p0) || (p0)) : (p0));
        (globalState).f2 = !((_107_v45).f15);
      } else {
        let _108_v46;
        _108_v46 = _dafny.Map.Empty.slice().updateUnsafe(_99_v40,_55_v1);
        _108_v46 = (_108_v46).Merge(_108_v46);
        let _109_v47;
        _109_v47 = _dafny.Seq.of(p0);
        let _110_v48;
        _110_v48 = _dafny.Map.Empty.slice().updateUnsafe(p0,_109_v47);
        let _111_v49;
        _111_v49 = _dafny.MultiSet.fromElements(p0);
        let _112_v50;
        let _nw12 = Array((new BigNumber(10)).toNumber());
        _nw12[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.FromArray((((_110_v48).contains(p0)) ? ((_110_v48).get(p0)) : (_dafny.Seq.of(p0))));
        _nw12[(_dafny.ONE).toNumber()] = (_111_v49).Difference(_dafny.MultiSet.fromElements(p0));
        _nw12[(new BigNumber(2)).toNumber()] = (_111_v49).update(p0, _module.__default.abs(_55_v1));
        _nw12[(new BigNumber(3)).toNumber()] = (_111_v49).Intersect(_111_v49);
        _nw12[(new BigNumber(4)).toNumber()] = _111_v49;
        _nw12[(new BigNumber(5)).toNumber()] = _111_v49;
        _nw12[(new BigNumber(6)).toNumber()] = (_111_v49).Union(_111_v49);
        _nw12[(new BigNumber(7)).toNumber()] = _dafny.MultiSet.FromArray(((p0) ? (_109_v47) : (_109_v47)));
        _nw12[(new BigNumber(8)).toNumber()] = _111_v49;
        _nw12[(new BigNumber(9)).toNumber()] = _111_v49;
        _112_v50 = _nw12;
        let _index7 = _module.__default.safeIndex(new BigNumber(402), new BigNumber((_112_v50).length));
        (_112_v50)[_index7] = _111_v49;
        let _index8 = _module.__default.safeIndex(new BigNumber(402), new BigNumber((_112_v50).length));
        (_112_v50)[_index8] = _111_v49;
        let _113_v51;
        _113_v51 = _dafny.Set.fromElements(p0, p0);
        let _114_v52;
        _114_v52 = _module.D1.create_DC2(_55_v1, p0, new BigNumber((_113_v51).length), _55_v1);
        let _115_v53;
        _115_v53 = _dafny.Seq.of(_114_v52);
        let _116_v54;
        let _nw13 = new _module.C1();
        _nw13.__ctor(_115_v53);
        _116_v54 = _nw13;
        let _117_v55;
        _117_v55 = _dafny.Set.fromElements(_116_v54);
        _55_v1 = new BigNumber((_117_v55).length);
        let _118_v56;
        _118_v56 = _module.D10.create_DC27(_dafny.Map.Empty.slice().updateUnsafe(p0,_55_v1));
        let _119_v57;
        _119_v57 = _dafny.Map.Empty.slice().updateUnsafe(true,_55_v1);
        let _pat_let_tv0 = _119_v57;
        let _pat_let_tv1 = _119_v57;
        let _source3 = function (_pat_let0_0) {
          return function (_120_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_121_dt__update_hcf50_h0) {
                return _module.D10.create_DC27(_121_dt__update_hcf50_h0);
              }(_pat_let1_0);
            }((_pat_let_tv0).Merge(_pat_let_tv1));
          }(_pat_let0_0);
        }(_118_v56);
        if (_source3.is_DC28) {
          let _122_v58;
          let _init1 = function (_123_i3) {
            return _dafny.Seq.UnicodeFromString("fvtkev");
          };
          let _nw14 = Array((new BigNumber(14)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw14.length); _i0_1++) {
            _nw14[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _122_v58 = _nw14;
          let _index9 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_122_v58).length));
          (_122_v58)[_index9] = (_module.D13.create_DC36(_dafny.Seq.Create(_module.__default.abs(new BigNumber(579)), function (_124_i4) {
  return new _dafny.CodePoint('l'.codePointAt(0));
}), new BigNumber(327), _55_v1, _55_v1)).dtor_cf62;
          let _125_v59;
          _125_v59 = new _dafny.CodePoint('m'.codePointAt(0));
          let _126_v60;
          _126_v60 = _dafny.Map.Empty.slice().updateUnsafe(_55_v1,_116_v54);
          let _127_v61;
          _127_v61 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_126_v60).length),_125_v59);
          let _index10 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_122_v58).length));
          let _rhs12 = ((false) ? (new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_55_v1,_125_v59)).Merge(_127_v61)).length)) : (_55_v1));
          let _rhs13 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(158)), function (_128_i5) {
            return new _dafny.CodePoint('s'.codePointAt(0));
          });
          let _rhs14 = _module.__default.fm1(_59_v4, (((((_60_v5).contains(_55_v1)) ? ((_60_v5).get(_55_v1)) : (p0))) ? (new BigNumber((_59_v4).cardinality())) : (_55_v1)), globalState);
          let _lhs5 = _122_v58;
          let _lhs6 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_122_v58).length));
          let _lhs7 = globalState;
          _55_v1 = _rhs12;
          _lhs5[_lhs6] = _rhs13;
          _lhs7.f2 = _rhs14;
          _55_v1 = _55_v1;
          _113_v51 = _113_v51;
          (globalState).f2 = ((p0) ? (!(p0)) : (p0));
        } else {
          let _129___mcc_h0 = (_source3).cf50;
          let _130_cf50 = _129___mcc_h0;
          let _131_v62;
          let _nw15 = new _module.C4();
          _nw15.__ctor(true);
          _131_v62 = _nw15;
          let _132_v63;
          let _133_v64;
          let _134_v65;
          let _135_v66;
          let _out0;
          let _out1;
          let _out2;
          let _out3;
          let _outcollector0 = (_131_v62).m7(_55_v1, _55_v1, false, globalState);
          _out0 = _outcollector0[0];
          _out1 = _outcollector0[1];
          _out2 = _outcollector0[2];
          _out3 = _outcollector0[3];
          _132_v63 = _out0;
          _133_v64 = _out1;
          _134_v65 = _out2;
          _135_v66 = _out3;
          (globalState).f2 = !(new BigNumber(-737)).isEqualTo(new BigNumber((_111_v49).cardinality()));
          let _136_v67;
          let _137_v68;
          let _138_v69;
          let _139_v70;
          let _out4;
          let _out5;
          let _out6;
          let _out7;
          let _outcollector1 = (_131_v62).m7(_132_v63, _55_v1, p0, globalState);
          _out4 = _outcollector1[0];
          _out5 = _outcollector1[1];
          _out6 = _outcollector1[2];
          _out7 = _outcollector1[3];
          _136_v67 = _out4;
          _137_v68 = _out5;
          _138_v69 = _out6;
          _139_v70 = _out7;
        }
        let _140_v71;
        _140_v71 = _module.D16.create_DC46(_55_v1);
        let _141_v72;
        _141_v72 = _dafny.Map.Empty.slice().updateUnsafe(_140_v71,p0);
        let _rhs15 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_55_v1));
        let _rhs16 = (_55_v1).multipliedBy(((p0) ? (_module.__default.fm0(globalState)) : (_55_v1)));
        let _rhs17 = _module.__default.fm1(_59_v4, _55_v1, globalState);
        let _rhs18 = (((_module.__default.fm0(globalState)).isLessThanOrEqualTo(_55_v1)) ? (_141_v72) : (_141_v72));
        let _lhs8 = globalState;
        _55_v1 = _rhs15;
        _55_v1 = _rhs16;
        _lhs8.f2 = _rhs17;
        _141_v72 = _rhs18;
      }
      let _142_i6;
      _142_i6 = _dafny.ZERO;
      L1: {
        while (p0) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_142_i6)) {
              break L1;
            }
            _142_i6 = (_142_i6).plus(_dafny.ONE);
            let _143_v73;
            _143_v73 = _dafny.Seq.of(false, p0);
            let _144_v74;
            let _nw16 = new _module.C5();
            _nw16.__ctor(_55_v1, p0);
            _144_v74 = _nw16;
            let _145_v75;
            _145_v75 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_143_v73).length),_144_v74);
            let _146_v76;
            _146_v76 = _dafny.Seq.UnicodeFromString("mivytavqf");
            let _147_v77;
            let _nw17 = new _module.C4();
            _nw17.__ctor(p0);
            _147_v77 = _nw17;
            let _148_v78;
            _148_v78 = _dafny.Seq.of(_147_v77, _147_v77);
            let _149_v79;
            _149_v79 = _dafny.Seq.of(_dafny.Seq.of(_147_v77), _148_v78, _dafny.Seq.of(_147_v77, _147_v77, _147_v77));
            let _150_v80;
            _150_v80 = _dafny.Map.Empty.slice().updateUnsafe(_146_v76,_55_v1);
            let _151_v81;
            _151_v81 = _dafny.Seq.of(_146_v76);
            let _152_v82;
            let _nw18 = Array((new BigNumber(29)).toNumber());
            _nw18[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Set.fromElements(_143_v73, _143_v73)).length);
            _nw18[(_dafny.ONE).toNumber()] = new BigNumber(616);
            _nw18[(new BigNumber(2)).toNumber()] = (_55_v1).multipliedBy(_55_v1);
            _nw18[(new BigNumber(3)).toNumber()] = new BigNumber((((_145_v75).update(new BigNumber((_59_v4).cardinality()), _144_v74)).Merge(_145_v75)).length);
            _nw18[(new BigNumber(4)).toNumber()] = _55_v1;
            _nw18[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_55_v1).multipliedBy(_55_v1)));
            _nw18[(new BigNumber(6)).toNumber()] = (_55_v1).plus(_55_v1);
            _nw18[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(569)), function (_153_i7) {
              return new _dafny.CodePoint('v'.codePointAt(0));
            }), _146_v76)).length);
            _nw18[(new BigNumber(8)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_module.__default.fm6(_143_v73, globalState)).length), new BigNumber((_149_v79).length));
            _nw18[(new BigNumber(9)).toNumber()] = _55_v1;
            _nw18[(new BigNumber(10)).toNumber()] = _55_v1;
            _nw18[(new BigNumber(11)).toNumber()] = _55_v1;
            _nw18[(new BigNumber(12)).toNumber()] = new BigNumber((_146_v76).length);
            _nw18[(new BigNumber(13)).toNumber()] = (_55_v1).minus(_55_v1);
            _nw18[(new BigNumber(14)).toNumber()] = new BigNumber(989);
            _nw18[(new BigNumber(15)).toNumber()] = new BigNumber(912);
            _nw18[(new BigNumber(16)).toNumber()] = (_55_v1).minus((((_150_v80).contains((_151_v81)[_module.__default.safeIndex(_55_v1, new BigNumber((_151_v81).length))])) ? ((_150_v80).get((_151_v81)[_module.__default.safeIndex(_55_v1, new BigNumber((_151_v81).length))])) : (new BigNumber(967))));
            _nw18[(new BigNumber(17)).toNumber()] = _55_v1;
            _nw18[(new BigNumber(18)).toNumber()] = new BigNumber(530);
            _nw18[(new BigNumber(19)).toNumber()] = _55_v1;
            _nw18[(new BigNumber(20)).toNumber()] = new BigNumber(322);
            _nw18[(new BigNumber(21)).toNumber()] = _55_v1;
            _nw18[(new BigNumber(22)).toNumber()] = (new BigNumber((_module.__default.fm20(p0, !((_144_v74).f9), new BigNumber(-541), _146_v76, globalState)).length)).multipliedBy(_55_v1);
            _nw18[(new BigNumber(23)).toNumber()] = _module.__default.fm0(globalState);
            _nw18[(new BigNumber(24)).toNumber()] = (_dafny.ZERO).minus(_55_v1);
            _nw18[(new BigNumber(25)).toNumber()] = (_55_v1).multipliedBy(_55_v1);
            _nw18[(new BigNumber(26)).toNumber()] = _55_v1;
            _nw18[(new BigNumber(27)).toNumber()] = _55_v1;
            _nw18[(new BigNumber(28)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("k")).length);
            _152_v82 = _nw18;
            let _154_v83;
            _154_v83 = new _dafny.CodePoint('a'.codePointAt(0));
            let _155_v84;
            _155_v84 = _dafny.Set.fromElements(_154_v83, _154_v83);
            let _156_v85;
            _156_v85 = _dafny.Map.Empty.slice().updateUnsafe(p0,_155_v84);
            let _rhs19 = _152_v82;
            let _rhs20 = (_156_v85).equals(_dafny.Map.Empty.slice().updateUnsafe(!((_144_v74).f9),_155_v84));
            let _lhs9 = globalState;
            _152_v82 = _rhs19;
            _lhs9.f2 = _rhs20;
            let _157_v86;
            _157_v86 = _dafny.Map.Empty.slice().updateUnsafe((_144_v74).f9,_144_v74);
            _55_v1 = (new BigNumber(((_157_v86).Merge(_157_v86)).length)).minus(_55_v1);
            let _158_v87;
            _158_v87 = _dafny.MultiSet.fromElements(_56_v2, _56_v2);
            _55_v1 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt(_55_v1, new BigNumber((_146_v76).length)), (((_158_v87).contains(_56_v2)) ? ((_158_v87).get(_56_v2)) : (new BigNumber(540))));
            let _159_v88;
            let _nw19 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Map.Empty);
            _159_v88 = _nw19;
            let _index11 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_159_v88).length));
            (_159_v88)[_index11] = (_60_v5).Merge(_module.__default.fm15((_144_v74).f9, globalState));
            let _index12 = _module.__default.safeIndex(new BigNumber(640), new BigNumber((_159_v88).length));
            (_159_v88)[_index12] = (function () {
              let _coll12 = new _dafny.Map();
              for (const _compr_12 of (_60_v5).Keys.Elements) {
                let _160_v89 = _compr_12;
                if ((_60_v5).contains(_160_v89)) {
                  _coll12.push([(_160_v89).plus(_55_v1),p0]);
                }
              }
              return _coll12;
            }()).Merge(_60_v5);
          }
        }
      }
      return;
    }
    static Main(__noArgsParameter) {
      let _161_v0;
      _161_v0 = true;
      let _162_v1;
      let _nw20 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Map.Empty);
      _162_v1 = _nw20;
      let _163_v2;
      _163_v2 = _dafny.Seq.UnicodeFromString("pa");
      let _164_globalState;
      let _nw21 = new _module.GlobalState();
      _nw21.__ctor(true, new BigNumber(162), true, false, ((_161_v0) ? (_162_v1) : (_162_v1)), true, new BigNumber(253), _163_v2, false);
      _164_globalState = _nw21;
      let _165_v3;
      let _init2 = ((_166_v0) => function (_167_i1) {
        return (_167_i1).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_166_v0,new BigNumber(882))).length)));
      })(_161_v0);
      let _nw22 = Array((new BigNumber(7)).toNumber());
      for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw22.length); _i0_2++) {
        _nw22[_i0_2] = _init2(new BigNumber(_i0_2));
      }
      _165_v3 = _nw22;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_165_v3).length))) {
        let _168_i0 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_168_i0)) && ((_168_i0).isLessThan(new BigNumber((_165_v3).length))))) {
          (_165_v3)[(_168_i0)] = _module.__default.safeDivisionInt(_168_i0, new BigNumber((((_161_v0) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(105)), ((_169_v0) => function (_170_i2) {
            return _dafny.MultiSet.fromElements(_169_v0, _169_v0);
          })(_161_v0))) : (_dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.of(_161_v0, _161_v0)), _dafny.MultiSet.fromElements(_161_v0, _161_v0, _161_v0))))).length));
        }
      }
      let _171_v4;
      _171_v4 = new BigNumber(490);
      let _hi0 = _module.__default.fm0(_164_globalState);
      for (let _172_i3 = _171_v4; _172_i3.isLessThan(_hi0); _172_i3 = _172_i3.plus(_dafny.ONE)) {
        _module.__default.m0(_161_v0, _164_globalState);
        let _index13 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((_165_v3).length));
        (_165_v3)[_index13] = (_dafny.ZERO).minus((_172_i3).multipliedBy(_172_i3));
        let _173_v5;
        let _init3 = function (_174_i4) {
          return false;
        };
        let _nw23 = Array((new BigNumber(9)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw23.length); _i0_3++) {
          _nw23[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _173_v5 = _nw23;
        let _175_v6;
        _175_v6 = new _dafny.CodePoint('g'.codePointAt(0));
        let _176_v7;
        _176_v7 = _dafny.Map.Empty.slice().updateUnsafe(_173_v5,_175_v6);
        let _index14 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((_165_v3).length));
        let _rhs21 = new BigNumber((_176_v7).length);
        let _rhs22 = _172_i3;
        let _lhs10 = _165_v3;
        let _lhs11 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((_165_v3).length));
        _171_v4 = _rhs21;
        _lhs10[_lhs11] = _rhs22;
        let _177_v8;
        _177_v8 = _dafny.Map.Empty.slice().updateUnsafe(_171_v4,(_165_v3)[_module.__default.safeIndex(new BigNumber(487), new BigNumber((_165_v3).length))]);
        let _178_v9;
        _178_v9 = _dafny.Seq.of(_173_v5, _173_v5);
        _171_v4 = _module.__default.safeDivisionInt(new BigNumber((_177_v8).length), new BigNumber((_dafny.Seq.Concat(_178_v9, _178_v9)).length));
        let _179_v10;
        _179_v10 = _dafny.MultiSet.fromElements((_165_v3)[_module.__default.safeIndex(new BigNumber(487), new BigNumber((_165_v3).length))], (_165_v3)[_module.__default.safeIndex(new BigNumber(487), new BigNumber((_165_v3).length))]);
        _161_v0 = _module.__default.fm1(((_161_v0) ? (_179_v10) : (_179_v10)), (_165_v3)[_module.__default.safeIndex(new BigNumber(487), new BigNumber((_165_v3).length))], _164_globalState);
      }
      let _180_i5;
      _180_i5 = _dafny.ZERO;
      L2: {
        while (_161_v0) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_180_i5)) {
              break L2;
            }
            _180_i5 = (_180_i5).plus(_dafny.ONE);
            let _181_v11;
            _181_v11 = new _dafny.CodePoint('j'.codePointAt(0));
            let _182_v12;
            _182_v12 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_181_v11,_171_v4),_161_v0);
            let _183_v13;
            _183_v13 = _dafny.Seq.of(_163_v2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(655)), ((_184_v11) => function (_185_i6) {
              return _184_v11;
            })(_181_v11)), _163_v2, _163_v2, _163_v2);
            let _186_v15;
            _186_v15 = _dafny.Map.Empty.slice().updateUnsafe(_181_v11,new BigNumber((function () {
              let _coll13 = new _dafny.Set();
              for (const _compr_13 of (_183_v13).Elements) {
                let _187_v14 = _compr_13;
                if (_dafny.Seq.contains(_183_v13, _187_v14)) {
                  _coll13.add(_187_v14);
                }
              }
              return _coll13;
            }()).length));
            _182_v12 = (_182_v12).update((_186_v15).Merge(_186_v15), _161_v0);
            let _188_v16;
            let _nw24 = new _module.C4();
            _nw24.__ctor(_161_v0);
            _188_v16 = _nw24;
            (_164_globalState).f2 = _dafny.Seq.IsPrefixOf(_dafny.Seq.update(_dafny.Seq.Concat(_module.__default.fm28(_171_v4, !(_161_v0), _164_globalState), _163_v2), _module.__default.safeIndex(_171_v4, new BigNumber((_dafny.Seq.Concat(_module.__default.fm28(_171_v4, !(_161_v0), _164_globalState), _163_v2)).length)), _181_v11), _dafny.Seq.of(_181_v11));
            if (!(_161_v0) || (_161_v0)) {
              let _189_v17;
              _189_v17 = _dafny.Seq.of(_161_v0);
              let _190_v18;
              _190_v18 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(_171_v4)).cardinality()), new BigNumber((_189_v17).length));
              let _191_v19;
              _191_v19 = _dafny.Set.fromElements(_dafny.Seq.of(new BigNumber((_190_v18).cardinality())), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-165)), ((_192_v4) => function (_193_i7) {
                return _192_v4;
              })(_171_v4)));
              let _194_v20;
              let _nw25 = new _module.C0();
              _nw25.__ctor(_191_v19, (_dafny.ZERO).minus(_171_v4));
              _194_v20 = _nw25;
              let _195_v21;
              _195_v21 = _module.D1.create_DC2(_194_v20.f12, (_194_v20.f12).isEqualTo(_171_v4), _module.__default.safeModuloInt(_171_v4, _194_v20.f12), _171_v4);
              _195_v21 = _195_v21;
              let _196_v22;
              _196_v22 = _dafny.Map.Empty.slice().updateUnsafe(_171_v4,_165_v3);
              let _rhs23 = (_196_v22).Merge((_196_v22).Merge(_196_v22));
              let _rhs24 = _162_v1;
              _196_v22 = _rhs23;
              _162_v1 = _rhs24;
              let _197_v23;
              let _nw26 = new _module.C2();
              _nw26.__ctor(_171_v4, _161_v0);
              _197_v23 = _nw26;
              _197_v23 = _197_v23;
              _161_v0 = (_171_v4).isLessThanOrEqualTo(new BigNumber((_163_v2).length));
            } else {
              let _198_v24;
              let _nw27 = Array((new BigNumber(28)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
              _198_v24 = _nw27;
              let _index15 = _module.__default.safeIndex(new BigNumber(589), new BigNumber((_198_v24).length));
              (_198_v24)[_index15] = ((_161_v0) ? (_181_v11) : (new _dafny.CodePoint('o'.codePointAt(0))));
              let _index16 = _module.__default.safeIndex(new BigNumber(589), new BigNumber((_198_v24).length));
              (_198_v24)[_index16] = _181_v11;
              let _rhs25 = _171_v4;
              let _rhs26 = _163_v2;
              _171_v4 = _rhs25;
              _163_v2 = _rhs26;
              let _199_v25;
              _199_v25 = _dafny.Seq.of(_161_v0);
              let _200_v26;
              let _nw28 = new _module.C5();
              _nw28.__ctor(new BigNumber((_199_v25).length), (_161_v0) || (!(_161_v0)));
              _200_v26 = _nw28;
              let _201_v27;
              _201_v27 = _module.D1.create_DC2((_200_v26).f10, true, _171_v4, (_dafny.ZERO).minus((_200_v26).f10));
              let _202_v28;
              _202_v28 = _dafny.Seq.of(_201_v27);
              let _203_v29;
              let _nw29 = new _module.C1();
              _nw29.__ctor(_202_v28);
              _203_v29 = _nw29;
              let _204_v30;
              _204_v30 = _dafny.Map.Empty.slice().updateUnsafe(_203_v29,_171_v4);
              _204_v30 = (_204_v30).update(_203_v29, new BigNumber(324));
              let _205_v31;
              _205_v31 = _dafny.Seq.of((_200_v26).f10);
              let _206_v32;
              _206_v32 = _dafny.Set.fromElements(_205_v31);
              let _207_v33;
              _207_v33 = _dafny.Map.Empty.slice().updateUnsafe((_200_v26).f10,_171_v4);
              let _208_v34;
              _208_v34 = _dafny.MultiSet.fromElements(_165_v3, _165_v3);
              let _209_v35;
              let _nw30 = new _module.C0();
              _nw30.__ctor(_206_v32, (((_207_v33).contains(new BigNumber(901))) ? ((_207_v33).get(new BigNumber(901))) : ((((_208_v34).contains(_165_v3)) ? ((_208_v34).get(_165_v3)) : ((_dafny.ZERO).minus(_171_v4))))));
              _209_v35 = _nw30;
            }
          }
        }
      }
      let _hi1 = _171_v4;
      for (let _210_i8 = new BigNumber((_163_v2).length); _210_i8.isLessThan(_hi1); _210_i8 = _210_i8.plus(_dafny.ONE)) {
        let _211_v36;
        _211_v36 = new _dafny.CodePoint('n'.codePointAt(0));
        let _212_v37;
        _212_v37 = _dafny.MultiSet.fromElements(_211_v36, _211_v36, _211_v36);
        let _213_v38;
        _213_v38 = _dafny.Seq.of(_module.D1.create_DC2(_210_i8, _161_v0, _171_v4, new BigNumber((_212_v37).cardinality())));
        let _214_v39;
        let _nw31 = new _module.C1();
        _nw31.__ctor(_dafny.Seq.Concat(_213_v38, _213_v38));
        _214_v39 = _nw31;
        let _215_v41;
        _215_v41 = _module.D17.create_DC48(_214_v39);
        let _rhs27 = (_module.__default.fm46(_161_v0, _161_v0, !(!(_161_v0)), _164_globalState)).equals(function () {
          let _coll14 = new _dafny.Map();
          for (const _compr_14 of _dafny.IntegerRange(new BigNumber(-582), new BigNumber(489))) {
            let _216_v40 = _compr_14;
            if (((new BigNumber(-582)).isLessThanOrEqualTo(_216_v40)) && ((_216_v40).isLessThan(new BigNumber(489)))) {
              _coll14.push([_module.__default.safeModuloInt(_216_v40, _171_v4),_163_v2]);
            }
          }
          return _coll14;
        }());
        let _rhs28 = _171_v4;
        let _rhs29 = (_215_v41).dtor_cf81;
        let _lhs12 = _164_globalState;
        _lhs12.f2 = _rhs27;
        _171_v4 = _rhs28;
        _214_v39 = _rhs29;
        _module.__default.m0(_161_v0, _164_globalState);
        (_164_globalState).f2 = !(_210_i8).isEqualTo(_171_v4);
        let _217_v42;
        _217_v42 = _dafny.Seq.of(_161_v0);
        let _rhs30 = _210_i8;
        let _rhs31 = _161_v0;
        let _rhs32 = !(_161_v0) || (((_161_v0) ? ((_217_v42)[_module.__default.safeIndex(_171_v4, new BigNumber((_217_v42).length))]) : (_161_v0)));
        let _lhs13 = _164_globalState;
        _171_v4 = _rhs30;
        _lhs13.f2 = _rhs31;
        _161_v0 = _rhs32;
      }
      _module.__default.m0(_161_v0, _164_globalState);
      let _218_i9;
      _218_i9 = _dafny.ZERO;
      L3: {
        while (_161_v0) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_218_i9)) {
              break L3;
            }
            _218_i9 = (_218_i9).plus(_dafny.ONE);
            let _219_v43;
            let _nw32 = new _module.C2();
            _nw32.__ctor(new BigNumber(98), !(_161_v0));
            _219_v43 = _nw32;
            let _220_v44;
            _220_v44 = _dafny.Seq.of(_219_v43, _219_v43, _219_v43, _219_v43, _219_v43);
            _220_v44 = _220_v44;
            let _221_v45;
            _221_v45 = _dafny.MultiSet.fromElements(_171_v4);
            let _222_v47;
            let _init4 = ((_223_v43) => function (_224_i10) {
              return function () {
                let _coll15 = new _dafny.Set();
                for (const _compr_15 of _dafny.IntegerRange(new BigNumber(809), new BigNumber(-805))) {
                  let _225_v46 = _compr_15;
                  if (((new BigNumber(809)).isLessThanOrEqualTo(_225_v46)) && ((_225_v46).isLessThan(new BigNumber(-805)))) {
                    _coll15.add((_225_v46).minus((_223_v43).f14));
                  }
                }
                return _coll15;
              }();
            })(_219_v43);
            let _nw33 = Array((new BigNumber(11)).toNumber());
            for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw33.length); _i0_4++) {
              _nw33[_i0_4] = _init4(new BigNumber(_i0_4));
            }
            _222_v47 = _nw33;
            let _226_v48;
            let _nw34 = new _module.C3();
            _nw34.__ctor(false, _222_v47, _161_v0);
            _226_v48 = _nw34;
            let _227_v49;
            _227_v49 = _dafny.Map.Empty.slice().updateUnsafe((_219_v43).f14,_226_v48);
            let _228_v50;
            _228_v50 = _dafny.Map.Empty.slice().updateUnsafe(_161_v0,_163_v2);
            let _229_v51;
            _229_v51 = _dafny.Seq.of(new BigNumber(((((_228_v50).contains((_226_v48).f15)) ? ((_228_v50).get((_226_v48).f15)) : (_163_v2))).length), _171_v4);
            let _230_v52;
            _230_v52 = _dafny.Seq.of(new BigNumber((_229_v51).length));
            let _231_v53;
            _231_v53 = _module.D9.create_DC24(_module.__default.fm1(_221_v45, new BigNumber((_227_v49).length), _164_globalState), new BigNumber((_229_v51).length), _163_v2);
            _module.__default.m0((_231_v53).dtor_cf41, _164_globalState);
            _163_v2 = _163_v2;
            let _232_v54;
            let _nw35 = new _module.C3();
            _nw35.__ctor((_226_v48).f15, _226_v48.f16, _161_v0);
            _232_v54 = _nw35;
          }
        }
      }
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_165_v3).length))) {
        let _233_i11 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_233_i11)) && ((_233_i11).isLessThan(new BigNumber((_165_v3).length))))) {
          (_165_v3)[(_233_i11)] = _module.__default.safeDivisionInt(_233_i11, _171_v4);
        }
      }
      let _index17 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_165_v3).length));
      (_165_v3)[_index17] = new BigNumber(77);
      let _234_v55;
      _234_v55 = new _dafny.CodePoint('m'.codePointAt(0));
      let _235_v56;
      _235_v56 = _module.D1.create_DC2(_171_v4, _161_v0, _171_v4, new BigNumber((_163_v2).length));
      let _236_v57;
      _236_v57 = _dafny.Seq.of(_235_v56);
      let _237_v58;
      let _nw36 = new _module.C1();
      _nw36.__ctor(_236_v57);
      _237_v58 = _nw36;
      let _238_v59;
      _238_v59 = _module.D17.create_DC50(_234_v55, _161_v0, _dafny.Seq.UnicodeFromString("awbtnvw"), _171_v4, _237_v58);
      let _index18 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_165_v3).length));
      let _rhs33 = _171_v4;
      let _rhs34 = _238_v59;
      let _lhs14 = _165_v3;
      let _lhs15 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_165_v3).length));
      _lhs14[_lhs15] = _rhs33;
      _238_v59 = _rhs34;
      let _239_v60;
      let _nw37 = new _module.C1();
      _nw37.__ctor(_236_v57);
      _239_v60 = _nw37;
      let _240_v61;
      _240_v61 = _module.D13.create_DC35(_161_v0, _234_v55, _239_v60, _161_v0);
      let _241_v62;
      let _nw38 = Array((new BigNumber(25)).toNumber());
      _nw38[(_dafny.ZERO).toNumber()] = _239_v60;
      _nw38[(_dafny.ONE).toNumber()] = _239_v60;
      _nw38[(new BigNumber(2)).toNumber()] = _239_v60;
      _nw38[(new BigNumber(3)).toNumber()] = _239_v60;
      _nw38[(new BigNumber(4)).toNumber()] = _239_v60;
      _nw38[(new BigNumber(5)).toNumber()] = _239_v60;
      _nw38[(new BigNumber(6)).toNumber()] = _239_v60;
      _nw38[(new BigNumber(7)).toNumber()] = _239_v60;
      _nw38[(new BigNumber(8)).toNumber()] = _239_v60;
      _nw38[(new BigNumber(9)).toNumber()] = _239_v60;
      _nw38[(new BigNumber(10)).toNumber()] = _239_v60;
      _nw38[(new BigNumber(11)).toNumber()] = _239_v60;
      _nw38[(new BigNumber(12)).toNumber()] = _239_v60;
      _nw38[(new BigNumber(13)).toNumber()] = _239_v60;
      _nw38[(new BigNumber(14)).toNumber()] = _239_v60;
      _nw38[(new BigNumber(15)).toNumber()] = ((false) ? (_239_v60) : (_239_v60));
      _nw38[(new BigNumber(16)).toNumber()] = _239_v60;
      _nw38[(new BigNumber(17)).toNumber()] = _239_v60;
      _nw38[(new BigNumber(18)).toNumber()] = (_240_v61).dtor_cf60;
      _nw38[(new BigNumber(19)).toNumber()] = _239_v60;
      _nw38[(new BigNumber(20)).toNumber()] = _239_v60;
      _nw38[(new BigNumber(21)).toNumber()] = _239_v60;
      _nw38[(new BigNumber(22)).toNumber()] = _239_v60;
      _nw38[(new BigNumber(23)).toNumber()] = _239_v60;
      _nw38[(new BigNumber(24)).toNumber()] = _239_v60;
      _241_v62 = _nw38;
      _241_v62 = _241_v62;
      let _242_v63;
      _242_v63 = _dafny.Seq.of(_165_v3);
      let _243_i12;
      _243_i12 = _dafny.ZERO;
      L4: {
        while (((_161_v0) ? (false) : ((new BigNumber((_242_v63).length)).isLessThanOrEqualTo(_module.__default.fm0(_164_globalState))))) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_243_i12)) {
              break L4;
            }
            _243_i12 = (_243_i12).plus(_dafny.ONE);
            let _244_v64;
            let _nw39 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Seq.of());
            _244_v64 = _nw39;
            _244_v64 = _244_v64;
            let _245_v65;
            _245_v65 = _module.D16.create_DC46(_171_v4);
            let _246_v66;
            _246_v66 = _module.D16.create_DC47(_245_v65);
            let _247_v67;
            _247_v67 = _module.D16.create_DC47(_246_v66);
            let _248_v68;
            _248_v68 = _dafny.Map.Empty.slice().updateUnsafe(_171_v4,_247_v67);
            _248_v68 = (_248_v68).update(_171_v4, _247_v67);
            if (!(_161_v0)) {
              let _249_v69;
              _249_v69 = _dafny.Map.Empty.slice().updateUnsafe(_161_v0,_161_v0);
              _249_v69 = (_249_v69).update(_161_v0, _161_v0);
              let _250_v70;
              _250_v70 = _dafny.Seq.of((_165_v3)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_165_v3).length))]);
              let _251_v71;
              _251_v71 = _dafny.Set.fromElements(_250_v70, _dafny.Seq.of(new BigNumber(948), (_165_v3)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_165_v3).length))]), _250_v70);
              let _252_v72;
              let _nw40 = new _module.C0();
              _nw40.__ctor(_251_v71, ((_165_v3)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_165_v3).length))]).plus(new BigNumber((_249_v69).length)));
              _252_v72 = _nw40;
              let _253_v73;
              _253_v73 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_module.__default.fm0(_164_globalState)),_161_v0);
              let _254_v74;
              _254_v74 = _module.D7.create_DC18(_253_v73);
              let _255_v75;
              _255_v75 = _dafny.Map.Empty.slice().updateUnsafe(_171_v4,_254_v74);
              _255_v75 = (_255_v75).update((_165_v3)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_165_v3).length))], _254_v74);
              let _pat_let_tv2 = _164_globalState;
              _254_v74 = function (_pat_let2_0) {
                return function (_256_dt__update__tmp_h0) {
                  return function (_pat_let3_0) {
                    return function (_257_dt__update_hcf32_h0) {
                      return _module.D7.create_DC18(_257_dt__update_hcf32_h0);
                    }(_pat_let3_0);
                  }(_module.__default.fm22(_pat_let_tv2));
                }(_pat_let2_0);
              }(_254_v74);
              (_252_v72).f12 = _171_v4;
            } else {
              _161_v0 = _161_v0;
              let _258_v76;
              let _init5 = ((_259_v3, _260_v4) => function (_261_i13) {
                return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_259_v3)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_259_v3).length))]),(_259_v3)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_259_v3).length))])).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(694)), ((_262_v4) => function (_263_i14) {
                  return _262_v4;
                })(_260_v4)),new BigNumber((_dafny.Set.fromElements((_259_v3)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_259_v3).length))])).length)));
              })(_165_v3, _171_v4);
              let _nw41 = Array((new BigNumber(28)).toNumber());
              for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw41.length); _i0_5++) {
                _nw41[_i0_5] = _init5(new BigNumber(_i0_5));
              }
              _258_v76 = _nw41;
              let _index19 = _module.__default.safeIndex(new BigNumber(880), new BigNumber((_258_v76).length));
              (_258_v76)[_index19] = _module.__default.fm47(_164_globalState);
              let _264_v77;
              _264_v77 = _dafny.Seq.of(_171_v4, (_165_v3)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_165_v3).length))]);
              let _265_v78;
              _265_v78 = _264_v77;
              let _266_v79;
              _266_v79 = _dafny.Map.Empty.slice().updateUnsafe(_265_v78,_171_v4);
              let _index20 = _module.__default.safeIndex(new BigNumber(880), new BigNumber((_258_v76).length));
              (_258_v76)[_index20] = ((_dafny.Map.Empty.slice().updateUnsafe(_265_v78,_171_v4)).Merge((_266_v79).update(_264_v77, new BigNumber((_264_v77).length)))).update(_265_v78, _171_v4);
              let _267_v80;
              _267_v80 = _dafny.MultiSet.fromElements((_165_v3)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_165_v3).length))]);
              let _268_v81;
              _268_v81 = _165_v3;
              let _rhs35 = _161_v0;
              let _rhs36 = _161_v0;
              let _rhs37 = ((_165_v3)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_165_v3).length))]).isLessThanOrEqualTo(new BigNumber((_267_v80).cardinality()));
              let _rhs38 = _161_v0;
              let _rhs39 = (_268_v81);
              let _lhs16 = _164_globalState;
              let _lhs17 = _164_globalState;
              let _lhs18 = _164_globalState;
              _lhs16.f2 = _rhs35;
              _lhs17.f2 = _rhs36;
              _161_v0 = _rhs37;
              _lhs18.f2 = _rhs38;
              _165_v3 = _rhs39;
              let _269_v82;
              _269_v82 = _module.D15.create_DC43();
              let _270_v83;
              _270_v83 = _module.D15.create_DC44(_269_v82);
              let _271_v84;
              let _nw42 = new _module.C2();
              _nw42.__ctor(new BigNumber((_163_v2).length), !(true));
              _271_v84 = _nw42;
              let _272_v85;
              _272_v85 = _dafny.Seq.of(_271_v84, _271_v84);
              let _273_v86;
              let _nw43 = new _module.C2();
              _nw43.__ctor(_module.__default.fm0(_164_globalState), true);
              _273_v86 = _nw43;
              let _274_v87;
              _274_v87 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat(_272_v85, _272_v85)).length),_273_v86);
              let _275_v88;
              let _init6 = ((_276_v2) => function (_277_i15) {
                return _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_276_v2).length)));
              })(_163_v2);
              let _nw44 = Array((new BigNumber(3)).toNumber());
              for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw44.length); _i0_6++) {
                _nw44[_i0_6] = _init6(new BigNumber(_i0_6));
              }
              _275_v88 = _nw44;
              let _278_v89;
              let _nw45 = new _module.C3();
              _nw45.__ctor(false, _275_v88, (_271_v84).f9);
              _278_v89 = _nw45;
              let _279_v90;
              _279_v90 = _module.D15.create_DC42(_161_v0, (_165_v3)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_165_v3).length))], _278_v89, _171_v4);
              let _280_v91;
              _280_v91 = _dafny.Map.Empty.slice().updateUnsafe(_279_v90,((_165_v3)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_165_v3).length))]).isLessThan((_dafny.ZERO).minus(_171_v4)));
              let _pat_let_tv3 = _269_v82;
              let _rhs40 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt(_171_v4, ((_273_v86).f14).multipliedBy((_273_v86).f14))));
              let _rhs41 = function (_pat_let4_0) {
                return function (_281_dt__update__tmp_h2) {
                  return function (_pat_let5_0) {
                    return function (_282_dt__update_hcf77_h0) {
                      return _module.D15.create_DC44(_282_dt__update_hcf77_h0);
                    }(_pat_let5_0);
                  }(_pat_let_tv3);
                }(_pat_let4_0);
              }(_270_v83);
              let _rhs42 = (((_280_v91).contains(_279_v90)) ? ((_280_v91).get(_279_v90)) : ((_278_v89).f9));
              let _rhs43 = (_274_v87).Merge(_274_v87);
              let _lhs19 = _164_globalState;
              _171_v4 = _rhs40;
              _270_v83 = _rhs41;
              _lhs19.f2 = _rhs42;
              _274_v87 = _rhs43;
              let _283_v92;
              let _nw46 = Array((new BigNumber(2)).toNumber());
              _nw46[(_dafny.ZERO).toNumber()] = (_278_v89).f9;
              _nw46[(_dafny.ONE).toNumber()] = (_278_v89).f9;
              _283_v92 = _nw46;
              let _284_v93;
              let _285_v94;
              let _out8;
              let _out9;
              let _outcollector2 = (_239_v60).m6(((_161_v0) ? (_237_v58) : (_237_v58)), _283_v92, (_271_v84).f9, _164_globalState);
              _out8 = _outcollector2[0];
              _out9 = _outcollector2[1];
              _284_v93 = _out8;
              _285_v94 = _out9;
            }
            _171_v4 = (_165_v3)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_165_v3).length))];
          }
        }
      }
      let _286_v95;
      _286_v95 = _dafny.Map.Empty.slice().updateUnsafe(_171_v4,(new BigNumber((_module.__default.fm23(_161_v0, _161_v0, _171_v4, _164_globalState)).length)).plus(_171_v4));
      let _287_v96;
      _287_v96 = _dafny.Map.Empty.slice().updateUnsafe(_171_v4,_161_v0);
      let _288_v97;
      _288_v97 = _module.D4.create_DC13(_171_v4);
      let _289_v98;
      let _nw47 = new _module.C2();
      _nw47.__ctor((_165_v3)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_165_v3).length))], _161_v0);
      _289_v98 = _nw47;
      let _290_v99;
      _290_v99 = _dafny.MultiSet.fromElements(_289_v98);
      let _291_v100;
      _291_v100 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_161_v0,(_289_v98).f14));
      let _292_v101;
      _292_v101 = _dafny.Seq.of((_291_v100)[_module.__default.safeIndex((_289_v98).f14, new BigNumber((_291_v100).length))]);
      let _293_v102;
      _293_v102 = _dafny.Map.Empty.slice().updateUnsafe(_161_v0,new BigNumber((_dafny.Set.fromElements(_161_v0, _161_v0)).length));
      let _294_v103;
      _294_v103 = _dafny.Map.Empty.slice().updateUnsafe(_161_v0,_161_v0);
      let _295_v104;
      _295_v104 = _module.D9.create_DC24(_161_v0, (_165_v3)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_165_v3).length))], _dafny.Seq.UnicodeFromString("fqe"));
      let _pat_let_tv4 = _165_v3;
      let _pat_let_tv5 = _165_v3;
      let _index21 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_165_v3).length));
      let _rhs44 = (_286_v95).Merge((_module.__default.fm48(_161_v0, _module.__default.fm11(false, _164_globalState), _161_v0, (((_286_v95).contains((_dafny.ZERO).minus((_165_v3)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_165_v3).length))]))) ? ((_286_v95).get((_dafny.ZERO).minus((_165_v3)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_165_v3).length))]))) : (new BigNumber((_287_v96).length))), _164_globalState)).update(new BigNumber(-614), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(function (_pat_let6_0) {
        return function (_296_dt__update__tmp_h3) {
          return function (_pat_let7_0) {
            return function (_297_dt__update_hcf25_h0) {
              return _module.D4.create_DC13(_297_dt__update_hcf25_h0);
            }(_pat_let7_0);
          }((_pat_let_tv5)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_pat_let_tv4).length))]);
        }(_pat_let6_0);
      }(_288_v97),_237_v58)).length)));
      let _rhs45 = (_165_v3)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_165_v3).length))];
      let _rhs46 = _234_v55;
      let _rhs47 = (((_290_v99).contains(_289_v98)) ? ((_290_v99).get(_289_v98)) : (new BigNumber((((false) ? ((_291_v100)[_module.__default.safeIndex(_171_v4, new BigNumber((_291_v100).length))]) : (_293_v102))).length)));
      let _rhs48 = !(((_161_v0) ? (false) : (!((_295_v104).dtor_cf41)))) || (((_165_v3)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_165_v3).length))]).isLessThan(new BigNumber((_294_v103).length)));
      let _lhs20 = _165_v3;
      let _lhs21 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_165_v3).length));
      let _lhs22 = _164_globalState;
      _286_v95 = _rhs44;
      _171_v4 = _rhs45;
      _234_v55 = _rhs46;
      _lhs20[_lhs21] = _rhs47;
      _lhs22.f2 = _rhs48;
      let _298_v105;
      let _nw48 = Array((new BigNumber(7)).toNumber()).fill(false);
      _298_v105 = _nw48;
      let _index22 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_165_v3).length));
      let _rhs49 = _298_v105;
      let _rhs50 = (_165_v3)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_165_v3).length))];
      let _rhs51 = _module.D4.create_DC13(_module.__default.safeModuloInt(new BigNumber((_163_v2).length), (_289_v98).f14));
      let _rhs52 = !(_161_v0);
      let _lhs23 = _165_v3;
      let _lhs24 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_165_v3).length));
      _298_v105 = _rhs49;
      _lhs23[_lhs24] = _rhs50;
      _288_v97 = _rhs51;
      _161_v0 = _rhs52;
      let _299_v106;
      let _nw49 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Map.Empty);
      _299_v106 = _nw49;
      _299_v106 = _299_v106;
      let _300_v107;
      _300_v107 = _module.D15.create_DC41();
      let _301_v108;
      _301_v108 = _module.D15.create_DC44(_300_v107);
      let _302_v109;
      _302_v109 = _dafny.Set.fromElements(_161_v0, _161_v0, _161_v0, _161_v0);
      let _303_v110;
      _303_v110 = _dafny.Map.Empty.slice().updateUnsafe(_301_v108,(_302_v109).Difference(_302_v109));
      _303_v110 = (_303_v110).Merge(_303_v110);
      let _304_v111;
      _304_v111 = _module.D17.create_DC48(((_161_v0) ? (_237_v58) : (_237_v58)));
      _304_v111 = _304_v111;
      let _index23 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_165_v3).length));
      (_165_v3)[_index23] = _171_v4;
      process.stdout.write(_dafny.toString(_161_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_163_v2).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_164_globalState.f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_164_globalState).f7).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_164_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_165_v3)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_165_v3)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_165_v3)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_165_v3)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_165_v3)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_165_v3)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_165_v3)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_171_v4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_180_i5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_218_i9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_234_v55));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_235_v56).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_235_v56).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_235_v56).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_235_v56).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_236_v57, _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(489), true, new BigNumber(489), new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_v59).dtor_cf82));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_v59).dtor_cf83));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_238_v59).dtor_cf84).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_238_v59).dtor_cf85));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_239_v60.f13, _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(489), true, new BigNumber(489), new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_240_v61).dtor_cf58));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_240_v61).dtor_cf59));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_240_v61).dtor_cf60.f13, _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(489), true, new BigNumber(489), new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_240_v61).dtor_cf61));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_241_v62)[_dafny.ZERO].f13, _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(489), true, new BigNumber(489), new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_241_v62)[_dafny.ONE].f13, _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(489), true, new BigNumber(489), new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_241_v62)[new BigNumber(2)].f13, _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(489), true, new BigNumber(489), new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_241_v62)[new BigNumber(3)].f13, _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(489), true, new BigNumber(489), new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_241_v62)[new BigNumber(4)].f13, _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(489), true, new BigNumber(489), new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_241_v62)[new BigNumber(5)].f13, _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(489), true, new BigNumber(489), new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_241_v62)[new BigNumber(6)].f13, _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(489), true, new BigNumber(489), new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_241_v62)[new BigNumber(7)].f13, _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(489), true, new BigNumber(489), new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_241_v62)[new BigNumber(8)].f13, _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(489), true, new BigNumber(489), new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_241_v62)[new BigNumber(9)].f13, _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(489), true, new BigNumber(489), new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_241_v62)[new BigNumber(10)].f13, _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(489), true, new BigNumber(489), new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_241_v62)[new BigNumber(11)].f13, _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(489), true, new BigNumber(489), new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_241_v62)[new BigNumber(12)].f13, _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(489), true, new BigNumber(489), new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_241_v62)[new BigNumber(13)].f13, _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(489), true, new BigNumber(489), new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_241_v62)[new BigNumber(14)].f13, _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(489), true, new BigNumber(489), new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_241_v62)[new BigNumber(15)].f13, _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(489), true, new BigNumber(489), new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_241_v62)[new BigNumber(16)].f13, _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(489), true, new BigNumber(489), new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_241_v62)[new BigNumber(17)].f13, _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(489), true, new BigNumber(489), new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_241_v62)[new BigNumber(18)].f13, _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(489), true, new BigNumber(489), new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_241_v62)[new BigNumber(19)].f13, _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(489), true, new BigNumber(489), new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_241_v62)[new BigNumber(20)].f13, _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(489), true, new BigNumber(489), new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_241_v62)[new BigNumber(21)].f13, _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(489), true, new BigNumber(489), new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_241_v62)[new BigNumber(22)].f13, _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(489), true, new BigNumber(489), new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_241_v62)[new BigNumber(23)].f13, _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(489), true, new BigNumber(489), new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_241_v62)[new BigNumber(24)].f13, _dafny.Seq.of(_module.D1.create_DC2(new BigNumber(489), true, new BigNumber(489), new BigNumber(2))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_242_v63).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_243_i12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_286_v95).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(489),new BigNumber(491)).updateUnsafe(new BigNumber(2),new BigNumber(849)).updateUnsafe(new BigNumber(-614),_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_287_v96).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(489),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_288_v97).dtor_cf25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_289_v98).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_290_v99).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_291_v100, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(489))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_292_v101, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(489))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_293_v102).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_294_v103).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_295_v104).dtor_cf41));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_295_v104).dtor_cf42));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_295_v104).dtor_cf43).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_302_v109).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_303_v110).length)));
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
        return other.$tag === 0 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Seq.of();
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
    static create_DC2(cf2, cf3, cf4, cf5) {
      let $dt = new D1(0);
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC3(cf6, cf7, cf8) {
      let $dt = new D1(1);
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC4(cf9, cf10, cf11, cf12, cf13) {
      let $dt = new D1(2);
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC1(cf1) {
      let $dt = new D1(3);
      $dt.cf1 = cf1;
      return $dt;
    }
    static create_DC5(cf14) {
      let $dt = new D1(4);
      $dt.cf14 = cf14;
      return $dt;
    }
    get is_DC2() { return this.$tag === 0; }
    get is_DC3() { return this.$tag === 1; }
    get is_DC4() { return this.$tag === 2; }
    get is_DC1() { return this.$tag === 3; }
    get is_DC5() { return this.$tag === 4; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf14() { return this.cf14; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf6) + ", " + this.cf7.toVerbatimString(true) + ", " + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 3) {
        return "D1.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else if (this.$tag === 4) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf14) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf2, other.cf2) && this.cf3 === other.cf3 && _dafny.areEqual(this.cf4, other.cf4) && _dafny.areEqual(this.cf5, other.cf5);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf6 === other.cf6 && _dafny.areEqual(this.cf7, other.cf7) && _dafny.areEqual(this.cf8, other.cf8);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf9, other.cf9) && this.cf10 === other.cf10 && this.cf11 === other.cf11 && this.cf12 === other.cf12 && this.cf13 === other.cf13;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf1, other.cf1);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf14, other.cf14);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC2(_dafny.ZERO, false, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC6(cf15) {
      let $dt = new D2(0);
      $dt.cf15 = cf15;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get dtor_cf15() { return this.cf15; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf15) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf15, other.cf15);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Set.Empty;
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
    static create_DC8() {
      let $dt = new D3(0);
      return $dt;
    }
    static create_DC9(cf17, cf18) {
      let $dt = new D3(1);
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC10() {
      let $dt = new D3(2);
      return $dt;
    }
    static create_DC7(cf16) {
      let $dt = new D3(3);
      $dt.cf16 = cf16;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get is_DC10() { return this.$tag === 2; }
    get is_DC7() { return this.$tag === 3; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf16() { return this.cf16; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC8";
      } else if (this.$tag === 1) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC10";
      } else if (this.$tag === 3) {
        return "D3.DC7" + "(" + _dafny.toString(this.cf16) + ")";
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
        return other.$tag === 1 && this.cf17 === other.cf17 && _dafny.areEqual(this.cf18, other.cf18);
      } else if (this.$tag === 2) {
        return other.$tag === 2;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf16, other.cf16);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC8();
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
    static create_DC12(cf20, cf21, cf22, cf23, cf24) {
      let $dt = new D4(0);
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC13(cf25) {
      let $dt = new D4(1);
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC14(cf26, cf27, cf28) {
      let $dt = new D4(2);
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC11(cf19) {
      let $dt = new D4(3);
      $dt.cf19 = cf19;
      return $dt;
    }
    static create_DC15(cf29) {
      let $dt = new D4(4);
      $dt.cf29 = cf29;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get is_DC14() { return this.$tag === 2; }
    get is_DC11() { return this.$tag === 3; }
    get is_DC15() { return this.$tag === 4; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf29() { return this.cf29; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC14" + "(" + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 3) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf19) + ")";
      } else if (this.$tag === 4) {
        return "D4.DC15" + "(" + _dafny.toString(this.cf29) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf20 === other.cf20 && this.cf21 === other.cf21 && _dafny.areEqual(this.cf22, other.cf22) && _dafny.areEqual(this.cf23, other.cf23) && this.cf24 === other.cf24;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf25, other.cf25);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf26 === other.cf26 && this.cf27 === other.cf27 && _dafny.areEqual(this.cf28, other.cf28);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf19 === other.cf19;
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf29, other.cf29);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC12(false, false, _dafny.ZERO, _dafny.ZERO, false);
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
    static create_DC16(cf30) {
      let $dt = new D5(0);
      $dt.cf30 = cf30;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get dtor_cf30() { return this.cf30; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC16" + "(" + _dafny.toString(this.cf30) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf30 === other.cf30;
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
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC17(cf31) {
      let $dt = new D6(0);
      $dt.cf31 = cf31;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get dtor_cf31() { return this.cf31; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf31) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf31, other.cf31);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.MultiSet.Empty;
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
    static create_DC19(cf33, cf34, cf35) {
      let $dt = new D7(0);
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC20(cf36, cf37) {
      let $dt = new D7(1);
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC18(cf32) {
      let $dt = new D7(2);
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC21(cf38) {
      let $dt = new D7(3);
      $dt.cf38 = cf38;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get is_DC20() { return this.$tag === 1; }
    get is_DC18() { return this.$tag === 2; }
    get is_DC21() { return this.$tag === 3; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf38() { return this.cf38; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC19" + "(" + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ", " + this.cf35.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC20" + "(" + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC18" + "(" + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 3) {
        return "D7.DC21" + "(" + _dafny.toString(this.cf38) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf33 === other.cf33 && this.cf34 === other.cf34 && _dafny.areEqual(this.cf35, other.cf35);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf36 === other.cf36 && _dafny.areEqual(this.cf37, other.cf37);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf32, other.cf32);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf38, other.cf38);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC19([], false, _dafny.Seq.UnicodeFromString(""));
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
    static create_DC22(cf39) {
      let $dt = new D8(0);
      $dt.cf39 = cf39;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get dtor_cf39() { return this.cf39; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC22" + "(" + _dafny.toString(this.cf39) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf39 === other.cf39;
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
          return D8.Default();
        }
      };
    }
  }

  $module.D9 = class D9 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC24(cf41, cf42, cf43) {
      let $dt = new D9(0);
      $dt.cf41 = cf41;
      $dt.cf42 = cf42;
      $dt.cf43 = cf43;
      return $dt;
    }
    static create_DC25(cf44, cf45, cf46, cf47, cf48) {
      let $dt = new D9(1);
      $dt.cf44 = cf44;
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      return $dt;
    }
    static create_DC23(cf40) {
      let $dt = new D9(2);
      $dt.cf40 = cf40;
      return $dt;
    }
    static create_DC26(cf49) {
      let $dt = new D9(3);
      $dt.cf49 = cf49;
      return $dt;
    }
    get is_DC24() { return this.$tag === 0; }
    get is_DC25() { return this.$tag === 1; }
    get is_DC23() { return this.$tag === 2; }
    get is_DC26() { return this.$tag === 3; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf49() { return this.cf49; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC24" + "(" + _dafny.toString(this.cf41) + ", " + _dafny.toString(this.cf42) + ", " + this.cf43.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC25" + "(" + _dafny.toString(this.cf44) + ", " + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC23" + "(" + _dafny.toString(this.cf40) + ")";
      } else if (this.$tag === 3) {
        return "D9.DC26" + "(" + _dafny.toString(this.cf49) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf41 === other.cf41 && _dafny.areEqual(this.cf42, other.cf42) && _dafny.areEqual(this.cf43, other.cf43);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf44, other.cf44) && _dafny.areEqual(this.cf45, other.cf45) && _dafny.areEqual(this.cf46, other.cf46) && _dafny.areEqual(this.cf47, other.cf47) && _dafny.areEqual(this.cf48, other.cf48);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf40, other.cf40);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf49, other.cf49);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC24(false, _dafny.ZERO, _dafny.Seq.UnicodeFromString(""));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D9.Default();
        }
      };
    }
  }

  $module.D10 = class D10 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC28() {
      let $dt = new D10(0);
      return $dt;
    }
    static create_DC27(cf50) {
      let $dt = new D10(1);
      $dt.cf50 = cf50;
      return $dt;
    }
    get is_DC28() { return this.$tag === 0; }
    get is_DC27() { return this.$tag === 1; }
    get dtor_cf50() { return this.cf50; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC28";
      } else if (this.$tag === 1) {
        return "D10.DC27" + "(" + _dafny.toString(this.cf50) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf50, other.cf50);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC28();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D10.Default();
        }
      };
    }
  }

  $module.D11 = class D11 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC30() {
      let $dt = new D11(0);
      return $dt;
    }
    static create_DC29(cf51) {
      let $dt = new D11(1);
      $dt.cf51 = cf51;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get is_DC29() { return this.$tag === 1; }
    get dtor_cf51() { return this.cf51; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC30";
      } else if (this.$tag === 1) {
        return "D11.DC29" + "(" + _dafny.toString(this.cf51) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf51, other.cf51);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC30();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D11.Default();
        }
      };
    }
  }

  $module.D12 = class D12 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC32(cf53, cf54, cf55) {
      let $dt = new D12(0);
      $dt.cf53 = cf53;
      $dt.cf54 = cf54;
      $dt.cf55 = cf55;
      return $dt;
    }
    static create_DC31(cf52) {
      let $dt = new D12(1);
      $dt.cf52 = cf52;
      return $dt;
    }
    static create_DC33(cf56) {
      let $dt = new D12(2);
      $dt.cf56 = cf56;
      return $dt;
    }
    get is_DC32() { return this.$tag === 0; }
    get is_DC31() { return this.$tag === 1; }
    get is_DC33() { return this.$tag === 2; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf56() { return this.cf56; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC32" + "(" + _dafny.toString(this.cf53) + ", " + _dafny.toString(this.cf54) + ", " + _dafny.toString(this.cf55) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC31" + "(" + _dafny.toString(this.cf52) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC33" + "(" + _dafny.toString(this.cf56) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf53, other.cf53) && this.cf54 === other.cf54 && _dafny.areEqual(this.cf55, other.cf55);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf52, other.cf52);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf56, other.cf56);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC32(_dafny.ZERO, false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D12.Default();
        }
      };
    }
  }

  $module.D13 = class D13 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC35(cf58, cf59, cf60, cf61) {
      let $dt = new D13(0);
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      $dt.cf60 = cf60;
      $dt.cf61 = cf61;
      return $dt;
    }
    static create_DC36(cf62, cf63, cf64, cf65) {
      let $dt = new D13(1);
      $dt.cf62 = cf62;
      $dt.cf63 = cf63;
      $dt.cf64 = cf64;
      $dt.cf65 = cf65;
      return $dt;
    }
    static create_DC34(cf57) {
      let $dt = new D13(2);
      $dt.cf57 = cf57;
      return $dt;
    }
    get is_DC35() { return this.$tag === 0; }
    get is_DC36() { return this.$tag === 1; }
    get is_DC34() { return this.$tag === 2; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf57() { return this.cf57; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC35" + "(" + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ", " + _dafny.toString(this.cf60) + ", " + _dafny.toString(this.cf61) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC36" + "(" + this.cf62.toVerbatimString(true) + ", " + _dafny.toString(this.cf63) + ", " + _dafny.toString(this.cf64) + ", " + _dafny.toString(this.cf65) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC34" + "(" + _dafny.toString(this.cf57) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf58 === other.cf58 && _dafny.areEqual(this.cf59, other.cf59) && this.cf60 === other.cf60 && this.cf61 === other.cf61;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf62, other.cf62) && _dafny.areEqual(this.cf63, other.cf63) && _dafny.areEqual(this.cf64, other.cf64) && _dafny.areEqual(this.cf65, other.cf65);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf57 === other.cf57;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC35(false, new _dafny.CodePoint('D'.codePointAt(0)), null, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D13.Default();
        }
      };
    }
  }

  $module.D14 = class D14 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC38(cf67, cf68, cf69) {
      let $dt = new D14(0);
      $dt.cf67 = cf67;
      $dt.cf68 = cf68;
      $dt.cf69 = cf69;
      return $dt;
    }
    static create_DC39(cf70, cf71) {
      let $dt = new D14(1);
      $dt.cf70 = cf70;
      $dt.cf71 = cf71;
      return $dt;
    }
    static create_DC37(cf66) {
      let $dt = new D14(2);
      $dt.cf66 = cf66;
      return $dt;
    }
    get is_DC38() { return this.$tag === 0; }
    get is_DC39() { return this.$tag === 1; }
    get is_DC37() { return this.$tag === 2; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf66() { return this.cf66; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC38" + "(" + _dafny.toString(this.cf67) + ", " + _dafny.toString(this.cf68) + ", " + _dafny.toString(this.cf69) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC39" + "(" + _dafny.toString(this.cf70) + ", " + _dafny.toString(this.cf71) + ")";
      } else if (this.$tag === 2) {
        return "D14.DC37" + "(" + _dafny.toString(this.cf66) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf67, other.cf67) && this.cf68 === other.cf68 && _dafny.areEqual(this.cf69, other.cf69);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf70, other.cf70) && _dafny.areEqual(this.cf71, other.cf71);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf66, other.cf66);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC38(new _dafny.CodePoint('D'.codePointAt(0)), false, _dafny.Seq.of());
    }
    static Rtd() {
      return class {
        static get Default() {
          return D14.Default();
        }
      };
    }
  }

  $module.D15 = class D15 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC41() {
      let $dt = new D15(0);
      return $dt;
    }
    static create_DC42(cf73, cf74, cf75, cf76) {
      let $dt = new D15(1);
      $dt.cf73 = cf73;
      $dt.cf74 = cf74;
      $dt.cf75 = cf75;
      $dt.cf76 = cf76;
      return $dt;
    }
    static create_DC43() {
      let $dt = new D15(2);
      return $dt;
    }
    static create_DC40(cf72) {
      let $dt = new D15(3);
      $dt.cf72 = cf72;
      return $dt;
    }
    static create_DC44(cf77) {
      let $dt = new D15(4);
      $dt.cf77 = cf77;
      return $dt;
    }
    get is_DC41() { return this.$tag === 0; }
    get is_DC42() { return this.$tag === 1; }
    get is_DC43() { return this.$tag === 2; }
    get is_DC40() { return this.$tag === 3; }
    get is_DC44() { return this.$tag === 4; }
    get dtor_cf73() { return this.cf73; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf77() { return this.cf77; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC41";
      } else if (this.$tag === 1) {
        return "D15.DC42" + "(" + _dafny.toString(this.cf73) + ", " + _dafny.toString(this.cf74) + ", " + _dafny.toString(this.cf75) + ", " + _dafny.toString(this.cf76) + ")";
      } else if (this.$tag === 2) {
        return "D15.DC43";
      } else if (this.$tag === 3) {
        return "D15.DC40" + "(" + _dafny.toString(this.cf72) + ")";
      } else if (this.$tag === 4) {
        return "D15.DC44" + "(" + _dafny.toString(this.cf77) + ")";
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
        return other.$tag === 1 && this.cf73 === other.cf73 && _dafny.areEqual(this.cf74, other.cf74) && this.cf75 === other.cf75 && _dafny.areEqual(this.cf76, other.cf76);
      } else if (this.$tag === 2) {
        return other.$tag === 2;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf72, other.cf72);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf77, other.cf77);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC41();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D15.Default();
        }
      };
    }
  }

  $module.D16 = class D16 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC46(cf79) {
      let $dt = new D16(0);
      $dt.cf79 = cf79;
      return $dt;
    }
    static create_DC45(cf78) {
      let $dt = new D16(1);
      $dt.cf78 = cf78;
      return $dt;
    }
    static create_DC47(cf80) {
      let $dt = new D16(2);
      $dt.cf80 = cf80;
      return $dt;
    }
    get is_DC46() { return this.$tag === 0; }
    get is_DC45() { return this.$tag === 1; }
    get is_DC47() { return this.$tag === 2; }
    get dtor_cf79() { return this.cf79; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf80() { return this.cf80; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC46" + "(" + _dafny.toString(this.cf79) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC45" + "(" + _dafny.toString(this.cf78) + ")";
      } else if (this.$tag === 2) {
        return "D16.DC47" + "(" + _dafny.toString(this.cf80) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf79, other.cf79);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf78, other.cf78);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf80, other.cf80);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC46(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D16.Default();
        }
      };
    }
  }

  $module.D17 = class D17 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC49() {
      let $dt = new D17(0);
      return $dt;
    }
    static create_DC50(cf82, cf83, cf84, cf85, cf86) {
      let $dt = new D17(1);
      $dt.cf82 = cf82;
      $dt.cf83 = cf83;
      $dt.cf84 = cf84;
      $dt.cf85 = cf85;
      $dt.cf86 = cf86;
      return $dt;
    }
    static create_DC48(cf81) {
      let $dt = new D17(2);
      $dt.cf81 = cf81;
      return $dt;
    }
    static create_DC51(cf87) {
      let $dt = new D17(3);
      $dt.cf87 = cf87;
      return $dt;
    }
    get is_DC49() { return this.$tag === 0; }
    get is_DC50() { return this.$tag === 1; }
    get is_DC48() { return this.$tag === 2; }
    get is_DC51() { return this.$tag === 3; }
    get dtor_cf82() { return this.cf82; }
    get dtor_cf83() { return this.cf83; }
    get dtor_cf84() { return this.cf84; }
    get dtor_cf85() { return this.cf85; }
    get dtor_cf86() { return this.cf86; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf87() { return this.cf87; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC49";
      } else if (this.$tag === 1) {
        return "D17.DC50" + "(" + _dafny.toString(this.cf82) + ", " + _dafny.toString(this.cf83) + ", " + this.cf84.toVerbatimString(true) + ", " + _dafny.toString(this.cf85) + ", " + _dafny.toString(this.cf86) + ")";
      } else if (this.$tag === 2) {
        return "D17.DC48" + "(" + _dafny.toString(this.cf81) + ")";
      } else if (this.$tag === 3) {
        return "D17.DC51" + "(" + _dafny.toString(this.cf87) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf82, other.cf82) && this.cf83 === other.cf83 && _dafny.areEqual(this.cf84, other.cf84) && _dafny.areEqual(this.cf85, other.cf85) && this.cf86 === other.cf86;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf81 === other.cf81;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf87, other.cf87);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D17.create_DC49();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D17.Default();
        }
      };
    }
  }

  $module.D18 = class D18 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC53(cf89, cf90, cf91, cf92) {
      let $dt = new D18(0);
      $dt.cf89 = cf89;
      $dt.cf90 = cf90;
      $dt.cf91 = cf91;
      $dt.cf92 = cf92;
      return $dt;
    }
    static create_DC52(cf88) {
      let $dt = new D18(1);
      $dt.cf88 = cf88;
      return $dt;
    }
    get is_DC53() { return this.$tag === 0; }
    get is_DC52() { return this.$tag === 1; }
    get dtor_cf89() { return this.cf89; }
    get dtor_cf90() { return this.cf90; }
    get dtor_cf91() { return this.cf91; }
    get dtor_cf92() { return this.cf92; }
    get dtor_cf88() { return this.cf88; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC53" + "(" + _dafny.toString(this.cf89) + ", " + _dafny.toString(this.cf90) + ", " + _dafny.toString(this.cf91) + ", " + _dafny.toString(this.cf92) + ")";
      } else if (this.$tag === 1) {
        return "D18.DC52" + "(" + _dafny.toString(this.cf88) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf89, other.cf89) && this.cf90 === other.cf90 && _dafny.areEqual(this.cf91, other.cf91) && this.cf92 === other.cf92;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf88, other.cf88);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC53(_dafny.ZERO, null, _dafny.Seq.of(), false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D18.Default();
        }
      };
    }
  }

  $module.D19 = class D19 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC54(cf93) {
      let $dt = new D19(0);
      $dt.cf93 = cf93;
      return $dt;
    }
    get is_DC54() { return this.$tag === 0; }
    get dtor_cf93() { return this.cf93; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC54" + "(" + _dafny.toString(this.cf93) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf93, other.cf93);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Map.Empty;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D19.Default();
        }
      };
    }
  }

  $module.D20 = class D20 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC55(cf94) {
      let $dt = new D20(0);
      $dt.cf94 = cf94;
      return $dt;
    }
    get is_DC55() { return this.$tag === 0; }
    get dtor_cf94() { return this.cf94; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC55" + "(" + _dafny.toString(this.cf94) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf94, other.cf94);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.MultiSet.Empty;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D20.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.T1 = class T1 {
  };

  $module.T2 = class T2 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f2 = false;
      this._f0 = false;
      this._f1 = _dafny.ZERO;
      this._f3 = false;
      this._f4 = [];
      this._f5 = false;
      this._f6 = _dafny.ZERO;
      this._f7 = _dafny.Seq.UnicodeFromString("");
      this._f8 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      return;
    }
    get f0() {
      let _this = this;
      return _this._f0;
    };
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f3() {
      let _this = this;
      return _this._f3;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this.f11 = _dafny.Set.Empty;
      this.f12 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f11, f12) {
      let _this = this;
      (_this).f11 = f11;
      (_this).f12 = f12;
      return;
    }
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      let _source4 = _module.D1.create_DC2(_this.f12, !(!(true)), _this.f12, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(67)), function (_305_i0) {
  return new _dafny.CodePoint('r'.codePointAt(0));
})).length));
      if (_source4.is_DC2) {
        let _306___mcc_h0 = (_source4).cf2;
        let _307___mcc_h1 = (_source4).cf3;
        let _308___mcc_h2 = (_source4).cf4;
        let _309___mcc_h3 = (_source4).cf5;
        let _310_cf5 = _309___mcc_h3;
        let _311_cf4 = _308___mcc_h2;
        let _312_cf3 = _307___mcc_h1;
        let _313_cf2 = _306___mcc_h0;
        return _312_cf3;
      } else if (_source4.is_DC3) {
        let _314___mcc_h4 = (_source4).cf6;
        let _315___mcc_h5 = (_source4).cf7;
        let _316___mcc_h6 = (_source4).cf8;
        let _317_cf8 = _316___mcc_h6;
        let _318_cf7 = _315___mcc_h5;
        let _319_cf6 = _314___mcc_h4;
        if (_319_cf6) {
          return _319_cf6;
        } else {
          return _319_cf6;
        }
      } else if (_source4.is_DC4) {
        let _320___mcc_h7 = (_source4).cf9;
        let _321___mcc_h8 = (_source4).cf10;
        let _322___mcc_h9 = (_source4).cf11;
        let _323___mcc_h10 = (_source4).cf12;
        let _324___mcc_h11 = (_source4).cf13;
        let _325_cf13 = _324___mcc_h11;
        let _326_cf12 = _323___mcc_h10;
        let _327_cf11 = _322___mcc_h9;
        let _328_cf10 = _321___mcc_h8;
        let _329_cf9 = _320___mcc_h7;
        return true;
      } else if (_source4.is_DC1) {
        let _330___mcc_h12 = (_source4).cf1;
        let _331_cf1 = _330___mcc_h12;
        return (_331_cf1).isLessThanOrEqualTo(_this.f12);
      } else {
        let _332___mcc_h13 = (_source4).cf14;
        let _333_cf14 = _332___mcc_h13;
        return (_module.D1.create_DC3((_module.D1.create_DC4(_this.f12, false, true, false, false)).dtor_cf11, _dafny.Seq.UnicodeFromString("wfyipkfqa"), _this.f12)).dtor_cf6;
      }
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this.f13 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f13) {
      let _this = this;
      (_this).f13 = f13;
      return;
    }
    fm2(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("eiomh")).length)), _dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(628))))).cardinality()),false)).length)), _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(649)), function (_334_i0) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      })).length))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(56)), function (_335_i1) {
        return new BigNumber((_dafny.Set.fromElements(false, !(!(!(true))), !(false), !(!((_module.D1.create_DC4(new BigNumber((_dafny.Seq.of(new BigNumber(976))).length), false, false, false, false)).dtor_cf12)))).length);
      }));
    };
    m6(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let _336_v0;
      _336_v0 = new BigNumber(226);
      let _337_v1;
      _337_v1 = _dafny.Seq.of(_336_v0);
      let _338_v2;
      _338_v2 = _dafny.Seq.of(_337_v1);
      let _339_v3;
      _339_v3 = _dafny.MultiSet.fromElements(_336_v0);
      let _340_v4;
      _340_v4 = _dafny.Seq.of(!(p2), p2, p2);
      let _341_v5;
      let _nw50 = Array((new BigNumber(23)).toNumber());
      _nw50[(_dafny.ZERO).toNumber()] = _336_v0;
      _nw50[(_dafny.ONE).toNumber()] = new BigNumber((_338_v2).length);
      _nw50[(new BigNumber(2)).toNumber()] = _336_v0;
      _nw50[(new BigNumber(3)).toNumber()] = _336_v0;
      _nw50[(new BigNumber(4)).toNumber()] = _336_v0;
      _nw50[(new BigNumber(5)).toNumber()] = _336_v0;
      _nw50[(new BigNumber(6)).toNumber()] = _336_v0;
      _nw50[(new BigNumber(7)).toNumber()] = _336_v0;
      _nw50[(new BigNumber(8)).toNumber()] = _336_v0;
      _nw50[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_339_v3).cardinality()));
      _nw50[(new BigNumber(10)).toNumber()] = _336_v0;
      _nw50[(new BigNumber(11)).toNumber()] = _336_v0;
      _nw50[(new BigNumber(12)).toNumber()] = new BigNumber(893);
      _nw50[(new BigNumber(13)).toNumber()] = _336_v0;
      _nw50[(new BigNumber(14)).toNumber()] = _336_v0;
      _nw50[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus(_336_v0);
      _nw50[(new BigNumber(16)).toNumber()] = new BigNumber((_340_v4).length);
      _nw50[(new BigNumber(17)).toNumber()] = _336_v0;
      _nw50[(new BigNumber(18)).toNumber()] = _336_v0;
      _nw50[(new BigNumber(19)).toNumber()] = _336_v0;
      _nw50[(new BigNumber(20)).toNumber()] = _336_v0;
      _nw50[(new BigNumber(21)).toNumber()] = _336_v0;
      _nw50[(new BigNumber(22)).toNumber()] = (_dafny.ZERO).minus(_module.__default.fm0(globalState));
      _341_v5 = _nw50;
      let _342_v6;
      _342_v6 = _dafny.Map.Empty.slice().updateUnsafe(_341_v5,_341_v5);
      _342_v6 = (_342_v6).update(_341_v5, _341_v5);
      let _343_v7;
      let _nw51 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Set.Empty);
      _343_v7 = _nw51;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_343_v7).length))) {
        let _344_i0 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_344_i0)) && ((_344_i0).isLessThan(new BigNumber((_343_v7).length))))) {
          (_343_v7)[(_344_i0)] = ((_dafny.Set.fromElements(p2)).Union(_dafny.Set.fromElements(p2, p2, p2, p2, p2))).Difference(_dafny.Set.fromElements(p2, p2));
        }
      }
      _336_v0 = ((p2) ? (_module.__default.fm0(globalState)) : (_module.__default.fm0(globalState)));
      if (((_336_v0).multipliedBy(_336_v0)).isLessThanOrEqualTo((_dafny.ZERO).minus(_336_v0))) {
        let _345_v8;
        _345_v8 = _dafny.MultiSet.fromElements(p1, p1, p1);
        _345_v8 = ((_module.D3.create_DC7(_dafny.MultiSet.fromElements(p1))).dtor_cf16).Difference(_345_v8);
        let _index24 = _module.__default.safeIndex(new BigNumber(713), new BigNumber((_341_v5).length));
        (_341_v5)[_index24] = ((_dafny.ZERO).minus(_336_v0)).plus(_336_v0);
        let _346_v9;
        _346_v9 = _dafny.Map.Empty.slice().updateUnsafe(_336_v0,p2);
        let _347_v10;
        _347_v10 = _dafny.Map.Empty.slice().updateUnsafe(_336_v0,new BigNumber((_346_v9).length));
        let _index25 = _module.__default.safeIndex(new BigNumber(713), new BigNumber((_341_v5).length));
        (_341_v5)[_index25] = (_336_v0).minus((((_347_v10).contains(new BigNumber(180))) ? ((_347_v10).get(new BigNumber(180))) : ((_337_v1)[_module.__default.safeIndex(_336_v0, new BigNumber((_337_v1).length))])));
        let _348_v11;
        _348_v11 = _dafny.Set.fromElements(_337_v1);
        let _349_v12;
        let _nw52 = new _module.C0();
        _nw52.__ctor((_348_v11).Union(_dafny.Set.fromElements(_337_v1, _337_v1)), _module.__default.fm0(globalState));
        _349_v12 = _nw52;
        let _350_v13;
        _350_v13 = new _dafny.CodePoint('j'.codePointAt(0));
        let _351_v14;
        _351_v14 = _module.D1.create_DC3(false, _dafny.Seq.UnicodeFromString("hgnqr"), (_341_v5)[_module.__default.safeIndex(new BigNumber(713), new BigNumber((_341_v5).length))]);
        let _352_v15;
        _352_v15 = _dafny.Map.Empty.slice().updateUnsafe(_350_v13,!((_dafny.ZERO).minus(_349_v12.f12)).isEqualTo((_module.D1.create_DC4(_336_v0, (_351_v14).dtor_cf6, p2, p2, p2)).dtor_cf9));
        let _353_v16;
        _353_v16 = _dafny.Set.fromElements(_349_v12, _349_v12, _349_v12);
        let _354_v17;
        _354_v17 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements(_349_v12)).IsDisjointFrom(_353_v16),((p2) ? (_349_v12) : (_349_v12)));
        let _index26 = _module.__default.safeIndex(new BigNumber(713), new BigNumber((_341_v5).length));
        let _rhs53 = (((_354_v17).contains(p2)) ? ((_354_v17).get(p2)) : (_349_v12));
        let _rhs54 = _349_v12.f12;
        let _rhs55 = p2;
        let _rhs56 = _dafny.Map.Empty.slice().updateUnsafe(_350_v13,!(p2));
        let _lhs25 = _341_v5;
        let _lhs26 = _module.__default.safeIndex(new BigNumber(713), new BigNumber((_341_v5).length));
        let _lhs27 = globalState;
        _349_v12 = _rhs53;
        _lhs25[_lhs26] = _rhs54;
        _lhs27.f2 = _rhs55;
        _352_v15 = _rhs56;
        r1 = p2;
      } else {
        let _index27 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((p1).length));
        (p1)[_index27] = p2;
        let _355_v18;
        _355_v18 = _dafny.MultiSet.fromElements(p2, !(p2));
        let _index28 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((p1).length));
        let _rhs57 = p2;
        let _rhs58 = p2;
        let _rhs59 = new BigNumber((((_355_v18).Union(_dafny.MultiSet.FromArray(_340_v4))).Intersect(_355_v18)).cardinality());
        let _rhs60 = (new BigNumber((_339_v3).cardinality())).multipliedBy((_336_v0).multipliedBy(_336_v0));
        let _rhs61 = _336_v0;
        let _lhs28 = globalState;
        let _lhs29 = p1;
        let _lhs30 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((p1).length));
        _lhs28.f2 = _rhs57;
        _lhs29[_lhs30] = _rhs58;
        _336_v0 = _rhs59;
        _336_v0 = _rhs60;
        _336_v0 = _rhs61;
        let _rhs62 = (p1)[_module.__default.safeIndex(new BigNumber(602), new BigNumber((p1).length))];
        let _rhs63 = _336_v0;
        let _rhs64 = _339_v3;
        let _lhs31 = globalState;
        _lhs31.f2 = _rhs62;
        _336_v0 = _rhs63;
        _339_v3 = _rhs64;
        let _356_v20;
        _356_v20 = _dafny.Set.fromElements(p2);
        let _357_v21;
        _357_v21 = _dafny.MultiSet.fromElements(_356_v20, _356_v20, _356_v20, _356_v20);
        _336_v0 = (new BigNumber((function () {
          let _coll16 = new _dafny.Map();
          for (const _compr_16 of (_357_v21).Elements) {
            let _358_v19 = _compr_16;
            if ((_357_v21).contains(_358_v19)) {
              _coll16.push([_358_v19,(p1)[_module.__default.safeIndex(new BigNumber(602), new BigNumber((p1).length))]]);
            }
          }
          return _coll16;
        }()).length)).multipliedBy(_module.__default.safeDivisionInt(new BigNumber(47), _336_v0));
        let _359_v22;
        _359_v22 = _dafny.Set.fromElements(_dafny.Seq.of(_336_v0));
        let _360_v23;
        let _nw53 = new _module.C0();
        _nw53.__ctor(_359_v22, _336_v0);
        _360_v23 = _nw53;
        let _361_v24;
        _361_v24 = _dafny.Map.Empty.slice().updateUnsafe(_360_v23,(p1)[_module.__default.safeIndex(new BigNumber(602), new BigNumber((p1).length))]);
        _361_v24 = (_361_v24).update(_360_v23, !(((_dafny.ZERO).minus(_336_v0)).isLessThanOrEqualTo(_360_v23.f12)));
        (_360_v23).f12 = _360_v23.f12;
      }
      let _362_v25;
      _362_v25 = _337_v1;
      let _source5 = _362_v25;
      let _363___mcc_h0 = _source5;
      let _364_cf0 = _363___mcc_h0;
      let _365_v26;
      _365_v26 = new _dafny.CodePoint('l'.codePointAt(0));
      let _366_v27;
      _366_v27 = _dafny.Map.Empty.slice().updateUnsafe(_336_v0,_365_v26);
      _336_v0 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((new BigNumber((_366_v27).length)).multipliedBy((_dafny.ZERO).minus(_336_v0)), _336_v0));
      let _rhs65 = (false) || (true);
      let _rhs66 = ((!(p2)) ? (_336_v0) : ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat((_this).fm2(_336_v0, p2, p2, new BigNumber((_340_v4).length), globalState), _364_cf0)).length))));
      let _rhs67 = _module.__default.safeModuloInt(_module.__default.fm0(globalState), (_dafny.ZERO).minus(_336_v0));
      let _rhs68 = (!(p2)) && (p2);
      let _lhs32 = globalState;
      let _lhs33 = globalState;
      _lhs32.f2 = _rhs65;
      _336_v0 = _rhs66;
      _336_v0 = _rhs67;
      _lhs33.f2 = _rhs68;
      r0 = p2;
      let _367_v28;
      _367_v28 = _module.D3.create_DC10();
      let _368_v29;
      _368_v29 = _dafny.Map.Empty.slice().updateUnsafe(p2,_367_v28);
      let _369_v30;
      _369_v30 = _dafny.Seq.of(_368_v29);
      let _370_v31;
      _370_v31 = _dafny.Map.Empty.slice().updateUnsafe(_336_v0,_336_v0);
      r1 = !((_369_v30)[_module.__default.safeIndex(new BigNumber((_370_v31).length), new BigNumber((_369_v30).length))]).contains(p2);
      if (true) {
        let _371_v32;
        _371_v32 = new _dafny.CodePoint('t'.codePointAt(0));
        let _372_v33;
        _372_v33 = _dafny.Map.Empty.slice().updateUnsafe(_371_v32,p2);
        let _index29 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((p1).length));
        (p1)[_index29] = (((_372_v33).contains(_371_v32)) ? ((_372_v33).get(_371_v32)) : (p2));
        let _373_v34;
        _373_v34 = _dafny.Set.fromElements(new BigNumber(((p0).fm2((_dafny.ZERO).minus(_336_v0), true, p2, _336_v0, globalState)).length));
        let _index30 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((p1).length));
        (p1)[_index30] = ((_dafny.Set.fromElements(_336_v0, (_dafny.ZERO).minus(_336_v0))).Difference(_373_v34)).IsSubsetOf(_373_v34);
        let _374_v35;
        _374_v35 = _dafny.Seq.UnicodeFromString("dcyxdyb");
        let _375_v36;
        _375_v36 = _dafny.Map.Empty.slice().updateUnsafe(_336_v0,true);
        let _376_v37;
        let _nw54 = Array((new BigNumber(2)).toNumber());
        _nw54[(_dafny.ZERO).toNumber()] = p2;
        _nw54[(_dafny.ONE).toNumber()] = (_dafny.Set.fromElements(_336_v0, new BigNumber((_337_v1).length))).IsDisjointFrom(_module.__default.fm14(_336_v0, _336_v0, new BigNumber((_375_v36).length), globalState));
        _376_v37 = _nw54;
        let _rhs69 = _dafny.Seq.Concat(_dafny.Seq.Concat(_374_v35, _dafny.Seq.Create(_module.__default.abs(new BigNumber(404)), ((_377_v32) => function (_378_i1) {
          return _377_v32;
        })(_371_v32))), _374_v35);
        let _rhs70 = _376_v37;
        _374_v35 = _rhs69;
        _376_v37 = _rhs70;
        if ((p1)[_module.__default.safeIndex(new BigNumber(7), new BigNumber((p1).length))]) {
          let _379_v38;
          _379_v38 = _dafny.Set.fromElements(_337_v1, _dafny.Seq.of(_336_v0), _337_v1, _337_v1);
          let _380_v39;
          let _nw55 = new _module.C0();
          _nw55.__ctor((_379_v38).Intersect(_379_v38), (_336_v0).plus(new BigNumber(73)));
          _380_v39 = _nw55;
          let _381_v40;
          _381_v40 = _dafny.Map.Empty.slice().updateUnsafe((p1)[_module.__default.safeIndex(new BigNumber(7), new BigNumber((p1).length))],p2);
          (_380_v39).f12 = ((p2) ? (_module.__default.safeModuloInt(new BigNumber(144), new BigNumber((_dafny.Seq.of(_380_v39.f12)).length))) : ((new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("aajbnx"), _module.__default.safeIndex(new BigNumber((_381_v40).length), new BigNumber((_dafny.Seq.UnicodeFromString("aajbnx")).length)), _371_v32)).length)).minus(_336_v0)));
          let _index31 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((p1).length));
          (p1)[_index31] = false;
          let _index32 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((p1).length));
          (p1)[_index32] = _module.__default.fm1(_339_v3, new BigNumber((_dafny.Seq.UnicodeFromString("fgnkdwmjc")).length), globalState);
          let _382_v41;
          _382_v41 = _dafny.Map.Empty.slice().updateUnsafe(_380_v39.f12,_374_v35);
          _382_v41 = (_382_v41).update(_380_v39.f12, _374_v35);
        } else {
          let _383_v42;
          let _nw56 = Array((new BigNumber(23)).toNumber()).fill([]);
          _383_v42 = _nw56;
          let _384_v43;
          let _nw57 = Array((new BigNumber(24)).toNumber()).fill([]);
          _384_v43 = _nw57;
          let _index33 = _module.__default.safeIndex(new BigNumber(417), new BigNumber((_383_v42).length));
          (_383_v42)[_index33] = _384_v43;
          let _index34 = _module.__default.safeIndex(new BigNumber(417), new BigNumber((_383_v42).length));
          (_383_v42)[_index34] = _384_v43;
          let _index35 = _module.__default.safeIndex(new BigNumber(343), new BigNumber((_341_v5).length));
          (_341_v5)[_index35] = new BigNumber((_dafny.Seq.of(p2, p2)).length);
          let _index36 = _module.__default.safeIndex(new BigNumber(343), new BigNumber((_341_v5).length));
          (_341_v5)[_index36] = ((_module.__default.fm1(_339_v3, new BigNumber((_373_v34).length), globalState)) ? (new BigNumber((_337_v1).length)) : (_336_v0));
          let _385_v44;
          let _init7 = ((_386_v35) => function (_387_i2) {
            return (_387_i2).minus((_module.D1.create_DC1(new BigNumber((_386_v35).length))).dtor_cf1);
          })(_374_v35);
          let _nw58 = Array((new BigNumber(10)).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw58.length); _i0_7++) {
            _nw58[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _385_v44 = _nw58;
          let _388_v45;
          _388_v45 = _dafny.Map.Empty.slice().updateUnsafe((p1)[_module.__default.safeIndex(new BigNumber(7), new BigNumber((p1).length))],_385_v44);
          _388_v45 = (_388_v45).update(_module.__default.fm1(_339_v3, (_341_v5)[_module.__default.safeIndex(new BigNumber(343), new BigNumber((_341_v5).length))], globalState), (((_388_v45).contains(!(p2))) ? ((_388_v45).get(!(p2))) : (_385_v44)));
          let _389_v47;
          _389_v47 = _dafny.Set.fromElements(_337_v1);
          let _390_v48;
          let _nw59 = new _module.C0();
          _nw59.__ctor(_389_v47, (_341_v5)[_module.__default.safeIndex(new BigNumber(343), new BigNumber((_341_v5).length))]);
          _390_v48 = _nw59;
          let _391_v49;
          _391_v49 = _dafny.Map.Empty.slice().updateUnsafe(_390_v48,p2);
          let _index37 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((p1).length));
          let _index38 = _module.__default.safeIndex(new BigNumber(343), new BigNumber((_341_v5).length));
          let _index39 = _module.__default.safeIndex(new BigNumber(343), new BigNumber((_341_v5).length));
          let _rhs71 = ((true) ? (_module.__default.fm15(p2, globalState)) : ((_375_v36).Merge(function () {
            let _coll17 = new _dafny.Map();
            for (const _compr_17 of _dafny.IntegerRange(new BigNumber(-724), new BigNumber(374))) {
              let _392_v46 = _compr_17;
              if (((new BigNumber(-724)).isLessThanOrEqualTo(_392_v46)) && ((_392_v46).isLessThan(new BigNumber(374)))) {
                _coll17.push([_module.__default.safeDivisionInt(_392_v46, _336_v0),(p1)[_module.__default.safeIndex(new BigNumber(7), new BigNumber((p1).length))]]);
              }
            }
            return _coll17;
          }())));
          let _rhs72 = (_340_v4)[_module.__default.safeIndex(new BigNumber(((_391_v49).Merge(_dafny.Map.Empty.slice().updateUnsafe(_390_v48,(p1)[_module.__default.safeIndex(new BigNumber(7), new BigNumber((p1).length))]))).length), new BigNumber((_340_v4).length))];
          let _rhs73 = (_390_v48.f12).multipliedBy((_390_v48.f12).plus((_341_v5)[_module.__default.safeIndex(new BigNumber(343), new BigNumber((_341_v5).length))]));
          let _rhs74 = (_341_v5)[_module.__default.safeIndex(new BigNumber(343), new BigNumber((_341_v5).length))];
          let _lhs34 = p1;
          let _lhs35 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((p1).length));
          let _lhs36 = _341_v5;
          let _lhs37 = _module.__default.safeIndex(new BigNumber(343), new BigNumber((_341_v5).length));
          let _lhs38 = _341_v5;
          let _lhs39 = _module.__default.safeIndex(new BigNumber(343), new BigNumber((_341_v5).length));
          _375_v36 = _rhs71;
          _lhs34[_lhs35] = _rhs72;
          _lhs36[_lhs37] = _rhs73;
          _lhs38[_lhs39] = _rhs74;
          let _393_v50;
          _393_v50 = _dafny.Map.Empty.slice().updateUnsafe((p1)[_module.__default.safeIndex(new BigNumber(7), new BigNumber((p1).length))],new BigNumber(521));
          r1 = ((!(_393_v50).equals(_393_v50)) ? (true) : ((p1)[_module.__default.safeIndex(new BigNumber(7), new BigNumber((p1).length))]));
        }
        _336_v0 = _module.__default.safeModuloInt(_336_v0, (_dafny.ZERO).minus(_336_v0));
        _336_v0 = _module.__default.safeModuloInt(_336_v0, _336_v0);
      } else {
        let _394_v51;
        let _init8 = function (_395_i3) {
          return _module.D3.create_DC8();
        };
        let _nw60 = Array((new BigNumber(9)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw60.length); _i0_8++) {
          _nw60[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _394_v51 = _nw60;
        _394_v51 = _394_v51;
        let _396_v52;
        _396_v52 = _dafny.Map.Empty.slice().updateUnsafe(p2,_336_v0);
        _396_v52 = _396_v52;
        let _397_v53;
        let _init9 = function (_398_i4) {
          return new _dafny.CodePoint('v'.codePointAt(0));
        };
        let _nw61 = Array((new BigNumber(22)).toNumber());
        for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw61.length); _i0_9++) {
          _nw61[_i0_9] = _init9(new BigNumber(_i0_9));
        }
        _397_v53 = _nw61;
        let _399_v54;
        _399_v54 = new _dafny.CodePoint('v'.codePointAt(0));
        let _index40 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_397_v53).length));
        (_397_v53)[_index40] = _399_v54;
        let _index41 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_397_v53).length));
        (_397_v53)[_index41] = _399_v54;
        let _400_v55;
        _400_v55 = _dafny.MultiSet.fromElements(p2, p2, p2, true, p2);
        let _401_v56;
        _401_v56 = _module.D1.create_DC4(new BigNumber((_400_v55).cardinality()), p2, p2, p2, p2);
        let _402_v57;
        _402_v57 = _module.D1.create_DC5(_401_v56);
        let _403_v58;
        _403_v58 = _dafny.Map.Empty.slice().updateUnsafe(_399_v54,p2);
        let _404_v59;
        let _nw62 = new _module.C0();
        _nw62.__ctor(_module.__default.fm16(_403_v58, p2, globalState), _336_v0);
        _404_v59 = _nw62;
        let _405_v60;
        _405_v60 = _dafny.Map.Empty.slice().updateUnsafe(((_module.__default.fm1(_339_v3, _336_v0, globalState)) ? (_402_v57) : (_402_v57)),_404_v59);
        _405_v60 = _405_v60;
        if (!((!(p2)) === (p2))) {
          let _406_v61;
          _406_v61 = _dafny.Map.Empty.slice().updateUnsafe(p1,_404_v59.f12);
          let _407_v62;
          _407_v62 = _dafny.Seq.of(p1, p1);
          let _408_v63;
          _408_v63 = _dafny.Seq.UnicodeFromString("xuorgwewj");
          let _409_v64;
          _409_v64 = _dafny.Map.Empty.slice().updateUnsafe(_404_v59.f12,p2);
          let _410_v65;
          _410_v65 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(p2, (_404_v59).fm8(p2, p2, _409_v64, globalState), p2, p2, p2)).cardinality()),p2);
          let _411_v66;
          _411_v66 = _dafny.Seq.of(_409_v64);
          let _412_v67;
          _412_v67 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
          let _413_v68;
          _413_v68 = _module.D1.create_DC4(new BigNumber((_337_v1).length), p2, (((_412_v67).contains(p1)) ? ((_412_v67).get(p1)) : (p2)), p2, (_404_v59).fm8(p2, p2, _410_v65, globalState));
          let _rhs75 = new BigNumber(((_406_v61).Merge(_dafny.Map.Empty.slice().updateUnsafe((_407_v62)[_module.__default.safeIndex(new BigNumber((_408_v63).length), new BigNumber((_407_v62).length))],_404_v59.f12))).length);
          let _rhs76 = ((p2) ? (p2) : (!(new BigNumber((_411_v66).length)).isEqualTo(new BigNumber((_dafny.Seq.of(_413_v68, _413_v68)).length))));
          let _rhs77 = _341_v5;
          let _lhs40 = globalState;
          _336_v0 = _rhs75;
          _lhs40.f2 = _rhs76;
          _341_v5 = _rhs77;
          let _414_v69;
          _414_v69 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_340_v4, _module.__default.safeIndex((((_339_v3).contains(_336_v0)) ? ((_339_v3).get(_336_v0)) : (_404_v59.f12)), new BigNumber((_340_v4).length)), p2),!(p2));
          _414_v69 = (_414_v69).update(_dafny.Seq.Concat(_dafny.Seq.of(p2), _340_v4), (p2) && (p2));
          let _415_v70;
          _415_v70 = _dafny.Set.fromElements(p1, p1, p1);
          (globalState).f2 = !(p2) || (((_dafny.ZERO).minus(new BigNumber((_415_v70).length))).isLessThanOrEqualTo(_404_v59.f12));
          (_404_v59).f12 = (_dafny.ZERO).minus(_336_v0);
          let _416_v71;
          _416_v71 = _dafny.Map.Empty.slice().updateUnsafe(_337_v1,_404_v59.f12);
          let _417_v72;
          _417_v72 = _dafny.Map.Empty.slice().updateUnsafe(_336_v0,_404_v59.f12);
          _416_v71 = (_dafny.Map.Empty.slice().updateUnsafe(_337_v1,_404_v59.f12)).Merge((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_404_v59.f12, (((_417_v72).contains(new BigNumber((_module.__default.fm15(p2, globalState)).length))) ? ((_417_v72).get(new BigNumber((_module.__default.fm15(p2, globalState)).length))) : (_336_v0))),_336_v0)).Merge(_416_v71));
        } else {
          let _418_v73;
          _418_v73 = _dafny.Set.fromElements(_404_v59.f12, _336_v0);
          let _419_v74;
          _419_v74 = _418_v73;
          let _420_v75;
          _420_v75 = _dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC8(),_336_v0);
          let _421_v76;
          let _init10 = ((_422_v73) => function (_423_i5) {
            return _422_v73;
          })(_418_v73);
          let _nw63 = Array((new BigNumber(12)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw63.length); _i0_10++) {
            _nw63[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _421_v76 = _nw63;
          let _index42 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_421_v76).length));
          (_421_v76)[_index42] = _418_v73;
          let _424_v77;
          _424_v77 = _dafny.Seq.UnicodeFromString("rrhtmhjuf");
          let _425_v78;
          _425_v78 = _module.D1.create_DC3(!(_module.__default.fm1(_339_v3, _404_v59.f12, globalState)), _424_v77, new BigNumber(404));
          let _426_v79;
          _426_v79 = _dafny.Map.Empty.slice().updateUnsafe(_424_v77,p2);
          let _427_v80;
          _427_v80 = _dafny.Map.Empty.slice().updateUnsafe(_404_v59.f12,_424_v77);
          let _index43 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_421_v76).length));
          let _rhs78 = _419_v74;
          let _rhs79 = _module.__default.fm17(_404_v59.f12, _426_v79, p2, (((_400_v55).contains(p2)) ? ((_400_v55).get(p2)) : (new BigNumber((_427_v80).length))), globalState);
          let _rhs80 = (_dafny.ZERO).minus((new BigNumber(((_338_v2)[_module.__default.safeIndex(_404_v59.f12, new BigNumber((_338_v2).length))]).length)).minus(((p2) ? (_404_v59.f12) : (new BigNumber((_dafny.Seq.UnicodeFromString("d")).length)))));
          let _rhs81 = (_418_v73).Intersect(_418_v73);
          let _rhs82 = _425_v78;
          let _lhs41 = _404_v59;
          let _lhs42 = _421_v76;
          let _lhs43 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_421_v76).length));
          _419_v74 = _rhs78;
          _420_v75 = _rhs79;
          _lhs41.f12 = _rhs80;
          _lhs42[_lhs43] = _rhs81;
          _425_v78 = _rhs82;
          _337_v1 = _dafny.Seq.of(_404_v59.f12);
          let _428_v81;
          _428_v81 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(p2),false)).length),p2);
          _module.__default.m0((((_428_v81).contains(new BigNumber(701))) ? ((_428_v81).get(new BigNumber(701))) : (p2)), globalState);
          let _rhs83 = (new BigNumber(217)).multipliedBy(_module.__default.fm0(globalState));
          let _rhs84 = p2;
          let _rhs85 = new BigNumber(-297);
          let _lhs44 = _404_v59;
          _lhs44.f12 = _rhs83;
          r0 = _rhs84;
          _336_v0 = _rhs85;
          let _429_v82;
          _429_v82 = _dafny.Map.Empty.slice().updateUnsafe(false,_404_v59);
          _429_v82 = _dafny.Map.Empty.slice().updateUnsafe(p2,_404_v59);
        }
      }
      let _430_v83;
      _430_v83 = new _dafny.CodePoint('x'.codePointAt(0));
      let _431_v84;
      _431_v84 = _dafny.MultiSet.fromElements(_430_v83, _430_v83);
      r0 = ((!((_336_v0).isLessThanOrEqualTo(new BigNumber(565)))) ? (p2) : (!(new BigNumber((function () {
        let _coll18 = new _dafny.Set();
        for (const _compr_18 of (_431_v84).Elements) {
          let _432_v85 = _compr_18;
          if ((_431_v84).contains(_432_v85)) {
            _coll18.add(_432_v85);
          }
        }
        return _coll18;
      }()).length)).isEqualTo(_336_v0)));
      r1 = p2;
      return [r0, r1];
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f9 = false;
      this._f14 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f9() {
      let _this = this;
      return _this._f9;
    };
    __ctor(f14, f9) {
      let _this = this;
      (_this)._f14 = f14;
      (_this)._f9 = f9;
      return;
    }
    fm2(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.of((_this).f14), _dafny.Seq.of((_this).f14, (_this).f14, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(460)), function (_433_i0) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      })).length)));
    };
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _434_v0;
      _434_v0 = new BigNumber(884);
      let _435_v1;
      _435_v1 = _dafny.Seq.UnicodeFromString("vqsey");
      let _436_v2;
      _436_v2 = _dafny.Seq.of(p0);
      let _437_v3;
      _437_v3 = _dafny.Set.fromElements(_436_v2);
      let _438_v4;
      let _nw64 = new _module.C0();
      _nw64.__ctor(_437_v3, _module.__default.fm0(globalState));
      _438_v4 = _nw64;
      let _439_v5;
      _439_v5 = _dafny.Seq.of(_438_v4, _438_v4);
      let _440_v6;
      _440_v6 = new _dafny.CodePoint('o'.codePointAt(0));
      let _441_v7;
      _441_v7 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_435_v1, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_439_v5, _module.__default.safeIndex(p0, new BigNumber((_439_v5).length)), _438_v4), _module.__default.safeIndex(_438_v4.f12, new BigNumber((_dafny.Seq.update(_439_v5, _module.__default.safeIndex(p0, new BigNumber((_439_v5).length)), _438_v4)).length)), _438_v4)).length)), new BigNumber((_435_v1).length)), _440_v6), _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f14,p1)).length), new BigNumber((_dafny.Seq.update(_435_v1, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_439_v5, _module.__default.safeIndex(p0, new BigNumber((_439_v5).length)), _438_v4), _module.__default.safeIndex(_438_v4.f12, new BigNumber((_dafny.Seq.update(_439_v5, _module.__default.safeIndex(p0, new BigNumber((_439_v5).length)), _438_v4)).length)), _438_v4)).length)), new BigNumber((_435_v1).length)), _440_v6)).length)), _440_v6)).length));
      _434_v0 = ((p2) ? ((((_441_v7).contains((_this).f14)) ? ((_441_v7).get((_this).f14)) : ((_this).f14))) : (new BigNumber((_435_v1).length)));
      (globalState).f2 = !(p2);
      if (p1) {
        let _442_v8;
        let _nw65 = Array((new BigNumber(9)).toNumber()).fill([]);
        _442_v8 = _nw65;
        let _443_v9;
        _443_v9 = _442_v8;
        let _444_v10;
        let _nw66 = Array((new BigNumber(25)).toNumber()).fill(false);
        _444_v10 = _nw66;
        let _index44 = _module.__default.safeIndex(new BigNumber(992), new BigNumber((_444_v10).length));
        (_444_v10)[_index44] = (_this).f9;
        let _445_v11;
        _445_v11 = _dafny.Map.Empty.slice().updateUnsafe(p1,_440_v6);
        let _index45 = _module.__default.safeIndex(new BigNumber(992), new BigNumber((_444_v10).length));
        let _rhs86 = _443_v9;
        let _rhs87 = !(_445_v11).equals(_445_v11);
        let _lhs45 = _444_v10;
        let _lhs46 = _module.__default.safeIndex(new BigNumber(992), new BigNumber((_444_v10).length));
        _443_v9 = _rhs86;
        _lhs45[_lhs46] = _rhs87;
        let _446_v12;
        _446_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,((_module.__default.fm1(_441_v7, new BigNumber(620), globalState)) ? ((_this).f9) : ((_444_v10)[_module.__default.safeIndex(new BigNumber(992), new BigNumber((_444_v10).length))])));
        _446_v12 = _446_v12;
        let _447_v13;
        _447_v13 = _dafny.Set.fromElements(_434_v0);
        (globalState).f2 = (_447_v13).equals(_447_v13);
        if (true) {
          (_438_v4).f12 = _438_v4.f12;
          let _448_v14;
          _448_v14 = _dafny.Set.fromElements(_435_v1);
          let _449_v15;
          let _nw67 = Array((new BigNumber(2)).toNumber());
          _nw67[(_dafny.ZERO).toNumber()] = new BigNumber((_448_v14).length);
          _nw67[(_dafny.ONE).toNumber()] = _module.__default.safeModuloInt(_438_v4.f12, new BigNumber(679));
          _449_v15 = _nw67;
          let _index46 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_449_v15).length));
          (_449_v15)[_index46] = new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(_438_v4.f12), _438_v4.f12, new BigNumber((_436_v2).length), _438_v4.f12, _438_v4.f12)).length);
          let _index47 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_449_v15).length));
          (_449_v15)[_index47] = _438_v4.f12;
          let _450_v16;
          _450_v16 = _module.D1.create_DC2((_this).f14, p2, _438_v4.f12, _438_v4.f12);
          _450_v16 = _module.D1.create_DC2(new BigNumber(718), (_444_v10)[_module.__default.safeIndex(new BigNumber(992), new BigNumber((_444_v10).length))], _438_v4.f12, (_438_v4.f12).multipliedBy((_dafny.ZERO).minus(_434_v0)));
          let _451_v17;
          let _nw68 = Array((new BigNumber(23)).toNumber());
          _451_v17 = _nw68;
          let _index48 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_451_v17).length));
          (_451_v17)[_index48] = _438_v4;
          let _452_v18;
          _452_v18 = _dafny.MultiSet.fromElements(p1, _dafny.areEqual(_435_v1, _435_v1));
          let _453_v19;
          _453_v19 = _dafny.Map.Empty.slice().updateUnsafe(_441_v7,(_444_v10)[_module.__default.safeIndex(new BigNumber(992), new BigNumber((_444_v10).length))]);
          let _454_v20;
          _454_v20 = _module.D7.create_DC18(_453_v19);
          let _455_v21;
          _455_v21 = _dafny.Seq.of((_444_v10)[_module.__default.safeIndex(new BigNumber(992), new BigNumber((_444_v10).length))], p1, p1, (_this).f9, false);
          let _index49 = _module.__default.safeIndex(new BigNumber(992), new BigNumber((_444_v10).length));
          let _index50 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_449_v15).length));
          let _index51 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_451_v17).length));
          let _rhs88 = (_444_v10)[_module.__default.safeIndex(new BigNumber(992), new BigNumber((_444_v10).length))];
          let _rhs89 = _438_v4.f12;
          let _rhs90 = _438_v4;
          let _rhs91 = ((_452_v18).Intersect(_452_v18)).Union((_dafny.MultiSet.FromArray(_455_v21)).Difference(_dafny.MultiSet.fromElements((_this).f9)));
          let _rhs92 = _454_v20;
          let _lhs47 = _444_v10;
          let _lhs48 = _module.__default.safeIndex(new BigNumber(992), new BigNumber((_444_v10).length));
          let _lhs49 = _449_v15;
          let _lhs50 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_449_v15).length));
          let _lhs51 = _451_v17;
          let _lhs52 = _module.__default.safeIndex(new BigNumber(246), new BigNumber((_451_v17).length));
          _lhs47[_lhs48] = _rhs88;
          _lhs49[_lhs50] = _rhs89;
          _lhs51[_lhs52] = _rhs90;
          _452_v18 = _rhs91;
          _454_v20 = _rhs92;
          _455_v21 = _455_v21;
        } else {
          let _456_v22;
          _456_v22 = _dafny.Map.Empty.slice().updateUnsafe(_435_v1,_444_v10);
          _456_v22 = _456_v22;
          let _457_v23;
          let _nw69 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Map.Empty);
          _457_v23 = _nw69;
          let _nw70 = Array((_dafny.ONE).toNumber()).fill(_dafny.Map.Empty);
          _457_v23 = _nw70;
          let _458_v24;
          _458_v24 = _dafny.Set.fromElements(!(true), (_dafny.MultiSet.fromElements((_dafny.ZERO).minus(_438_v4.f12))).IsDisjointFrom(_441_v7), p1, p1, p1);
          _458_v24 = _458_v24;
          (globalState).f2 = p1;
          let _459_v25;
          _459_v25 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
          let _460_v26;
          let _nw71 = Array((new BigNumber(21)).toNumber());
          _nw71[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.update(_436_v2, _module.__default.safeIndex(_438_v4.f12, new BigNumber((_436_v2).length)), new BigNumber((_436_v2).length))).length);
          _nw71[(_dafny.ONE).toNumber()] = _434_v0;
          _nw71[(new BigNumber(2)).toNumber()] = new BigNumber(-78);
          _nw71[(new BigNumber(3)).toNumber()] = _438_v4.f12;
          _nw71[(new BigNumber(4)).toNumber()] = _434_v0;
          _nw71[(new BigNumber(5)).toNumber()] = p0;
          _nw71[(new BigNumber(6)).toNumber()] = p0;
          _nw71[(new BigNumber(7)).toNumber()] = _438_v4.f12;
          _nw71[(new BigNumber(8)).toNumber()] = _module.__default.fm0(globalState);
          _nw71[(new BigNumber(9)).toNumber()] = _438_v4.f12;
          _nw71[(new BigNumber(10)).toNumber()] = _434_v0;
          _nw71[(new BigNumber(11)).toNumber()] = _438_v4.f12;
          _nw71[(new BigNumber(12)).toNumber()] = (_this).f14;
          _nw71[(new BigNumber(13)).toNumber()] = new BigNumber((_459_v25).length);
          _nw71[(new BigNumber(14)).toNumber()] = _438_v4.f12;
          _nw71[(new BigNumber(15)).toNumber()] = (_this).f14;
          _nw71[(new BigNumber(16)).toNumber()] = p0;
          _nw71[(new BigNumber(17)).toNumber()] = p0;
          _nw71[(new BigNumber(18)).toNumber()] = _434_v0;
          _nw71[(new BigNumber(19)).toNumber()] = _438_v4.f12;
          _nw71[(new BigNumber(20)).toNumber()] = new BigNumber(976);
          _460_v26 = _nw71;
          let _461_v27;
          _461_v27 = _dafny.MultiSet.fromElements(_460_v26, _460_v26);
          (_this).m8(!((_this).f9), _461_v27, globalState);
        }
        (_438_v4).f12 = _438_v4.f12;
      } else {
        let _462_v28;
        _462_v28 = _dafny.Seq.of((_this).f9);
        let _463_v29;
        let _nw72 = Array((new BigNumber(21)).toNumber());
        _nw72[(_dafny.ZERO).toNumber()] = (_this).f14;
        _nw72[(_dafny.ONE).toNumber()] = p0;
        _nw72[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_462_v28).length));
        _nw72[(new BigNumber(3)).toNumber()] = _438_v4.f12;
        _nw72[(new BigNumber(4)).toNumber()] = _438_v4.f12;
        _nw72[(new BigNumber(5)).toNumber()] = (_this).f14;
        _nw72[(new BigNumber(6)).toNumber()] = p0;
        _nw72[(new BigNumber(7)).toNumber()] = _438_v4.f12;
        _nw72[(new BigNumber(8)).toNumber()] = new BigNumber((_435_v1).length);
        _nw72[(new BigNumber(9)).toNumber()] = _438_v4.f12;
        _nw72[(new BigNumber(10)).toNumber()] = (_this).f14;
        _nw72[(new BigNumber(11)).toNumber()] = new BigNumber((_436_v2).length);
        _nw72[(new BigNumber(12)).toNumber()] = _434_v0;
        _nw72[(new BigNumber(13)).toNumber()] = _438_v4.f12;
        _nw72[(new BigNumber(14)).toNumber()] = _module.__default.fm0(globalState);
        _nw72[(new BigNumber(15)).toNumber()] = new BigNumber(958);
        _nw72[(new BigNumber(16)).toNumber()] = _438_v4.f12;
        _nw72[(new BigNumber(17)).toNumber()] = new BigNumber((_dafny.Seq.of(p0)).length);
        _nw72[(new BigNumber(18)).toNumber()] = _438_v4.f12;
        _nw72[(new BigNumber(19)).toNumber()] = new BigNumber((_435_v1).length);
        _nw72[(new BigNumber(20)).toNumber()] = new BigNumber(55);
        _463_v29 = _nw72;
        let _464_v30;
        _464_v30 = _dafny.Seq.of(_463_v29);
        (_438_v4).f12 = (((_this).f9) ? (new BigNumber((_module.__default.fm24(_438_v4.f12, !(true), (_this).f9, p1, globalState)).length)) : ((new BigNumber((_464_v30).length)).multipliedBy(_438_v4.f12)));
        (globalState).f2 = (_this).f9;
        let _465_v31;
        _465_v31 = _dafny.MultiSet.fromElements(p2, p2);
        let _466_v32;
        _466_v32 = _module.D1.create_DC3((_this).f9, _dafny.Seq.UnicodeFromString("doecrfd"), p0);
        let _rhs93 = _435_v1;
        let _rhs94 = _dafny.Seq.update(_dafny.Seq.Concat(_436_v2, (_this).fm2(_438_v4.f12, p1, (_this).f9, new BigNumber((_465_v31).cardinality()), globalState)), _module.__default.safeIndex((new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("e"), _module.__default.safeIndex(new BigNumber(533), new BigNumber((_dafny.Seq.UnicodeFromString("e")).length)), _440_v6), _module.__default.safeIndex(new BigNumber(-52), new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("e"), _module.__default.safeIndex(new BigNumber(533), new BigNumber((_dafny.Seq.UnicodeFromString("e")).length)), _440_v6)).length)), _440_v6)).length)).minus(_438_v4.f12), new BigNumber((_dafny.Seq.Concat(_436_v2, (_this).fm2(_438_v4.f12, p1, (_this).f9, new BigNumber((_465_v31).cardinality()), globalState))).length)), (_this).f14);
        let _rhs95 = !(p1);
        let _rhs96 = _dafny.Seq.Concat((_466_v32).dtor_cf7, _435_v1);
        let _lhs53 = globalState;
        _435_v1 = _rhs93;
        _436_v2 = _rhs94;
        _lhs53.f2 = _rhs95;
        _435_v1 = _rhs96;
        let _467_v33;
        _467_v33 = _module.D4.create_DC12((_this).f9, p1, (_this).f14, new BigNumber((_436_v2).length), (_this).f9);
        let _468_v34;
        _468_v34 = _module.D1.create_DC2(_module.__default.safeModuloInt((_this).f14, _438_v4.f12), (_462_v28)[_module.__default.safeIndex(_module.__default.fm0(globalState), new BigNumber((_462_v28).length))], ((p1) ? ((_dafny.ZERO).minus(_438_v4.f12)) : (_434_v0)), ((_this).f14).minus((_467_v33).dtor_cf23));
        let _469_v35;
        _469_v35 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
        let _470_v36;
        _470_v36 = _dafny.Set.fromElements(new BigNumber((_469_v35).length));
        let _471_v37;
        _471_v37 = _dafny.MultiSet.fromElements(_463_v29, _463_v29);
        let _pat_let_tv6 = _471_v37;
        let _pat_let_tv7 = _471_v37;
        let _pat_let_tv8 = p1;
        _468_v34 = function (_pat_let8_0) {
          return function (_472_dt__update__tmp_h0) {
            return function (_pat_let9_0) {
              return function (_473_dt__update_hcf3_h0) {
                return function (_pat_let10_0) {
                  return function (_474_dt__update_hcf4_h0) {
                    return _module.D1.create_DC2((_472_dt__update__tmp_h0).dtor_cf2, _473_dt__update_hcf3_h0, _474_dt__update_hcf4_h0, (_472_dt__update__tmp_h0).dtor_cf5);
                  }(_pat_let10_0);
                }(new BigNumber((_dafny.Seq.of((_this).f9, true, _pat_let_tv8)).length));
              }(_pat_let9_0);
            }((_pat_let_tv6).IsProperSubsetOf(_pat_let_tv7));
          }(_pat_let8_0);
        }(_module.__default.fm25(new BigNumber((_470_v36).length), p1, p1, _441_v7, globalState));
        let _475_v38;
        _475_v38 = _dafny.Map.Empty.slice().updateUnsafe((p2) === (!(p1)),(_this).f14);
        _475_v38 = (_475_v38).update(p1, _module.__default.fm0(globalState));
      }
      let _476_v39;
      let _init11 = ((_477_p2) => function (_478_i0) {
        return _477_p2;
      })(p2);
      let _nw73 = Array((new BigNumber(2)).toNumber());
      for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw73.length); _i0_11++) {
        _nw73[_i0_11] = _init11(new BigNumber(_i0_11));
      }
      _476_v39 = _nw73;
      let _index52 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_476_v39).length));
      (_476_v39)[_index52] = p1;
      let _index53 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_476_v39).length));
      (_476_v39)[_index53] = !((_this).f9);
      let _479_v40;
      let _nw74 = Array((_dafny.ONE).toNumber());
      _479_v40 = _nw74;
      let _480_v41;
      _480_v41 = _module.D4.create_DC11(_479_v40);
      let _481_v42;
      let _nw75 = Array((new BigNumber(7)).toNumber());
      _nw75[(_dafny.ZERO).toNumber()] = _479_v40;
      _nw75[(_dafny.ONE).toNumber()] = (_480_v41).dtor_cf19;
      _nw75[(new BigNumber(2)).toNumber()] = _479_v40;
      _nw75[(new BigNumber(3)).toNumber()] = (_480_v41).dtor_cf19;
      _nw75[(new BigNumber(4)).toNumber()] = _479_v40;
      _nw75[(new BigNumber(5)).toNumber()] = _479_v40;
      _nw75[(new BigNumber(6)).toNumber()] = _479_v40;
      _481_v42 = _nw75;
      let _source6 = _481_v42;
      let _482___mcc_h0 = _source6;
      let _483_cf30 = _482___mcc_h0;
      (_438_v4).f12 = _434_v0;
      let _484_v43;
      _484_v43 = _module.D4.create_DC12(!((_this).f9), p2, (_dafny.ZERO).minus(_438_v4.f12), (_this).f14, false);
      (_438_v4).f12 = ((_484_v43).dtor_cf23).plus(_module.__default.safeModuloInt(_438_v4.f12, new BigNumber((_435_v1).length)));
      let _485_v44;
      _485_v44 = _dafny.Seq.of(_438_v4.f11, _438_v4.f11);
      let _486_v46;
      let _nw76 = new _module.C0();
      _nw76.__ctor(function () {
        let _coll19 = new _dafny.Set();
        for (const _compr_19 of ((_485_v44)[_module.__default.safeIndex(new BigNumber((_435_v1).length), new BigNumber((_485_v44).length))]).Elements) {
          let _487_v45 = _compr_19;
          if (((_485_v44)[_module.__default.safeIndex(new BigNumber((_435_v1).length), new BigNumber((_485_v44).length))]).contains(_487_v45)) {
            _coll19.add(_487_v45);
          }
        }
        return _coll19;
      }(), _438_v4.f12);
      _486_v46 = _nw76;
      let _488_v47;
      _488_v47 = _dafny.Map.Empty.slice().updateUnsafe(_440_v6,p2);
      let _489_v48;
      _489_v48 = _dafny.Set.fromElements(_488_v47, _488_v47);
      let _index54 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_476_v39).length));
      (_476_v39)[_index54] = !(((_489_v48).Difference(_489_v48)).IsSubsetOf(_489_v48));
      let _490_v49;
      _490_v49 = _dafny.Set.fromElements((_this).f14, (_this).f14);
      let _491_v50;
      _491_v50 = _dafny.Seq.of((_this).f9);
      let _492_v51;
      _492_v51 = _dafny.MultiSet.fromElements((_476_v39)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_476_v39).length))]);
      let _493_v52;
      let _nw77 = Array((new BigNumber(17)).toNumber());
      _nw77[(_dafny.ZERO).toNumber()] = (_434_v0).plus(new BigNumber((_491_v50).length));
      _nw77[(_dafny.ONE).toNumber()] = p0;
      _nw77[(new BigNumber(2)).toNumber()] = new BigNumber((_435_v1).length);
      _nw77[(new BigNumber(3)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_492_v51).cardinality()), new BigNumber((_436_v2).length));
      _nw77[(new BigNumber(4)).toNumber()] = p0;
      _nw77[(new BigNumber(5)).toNumber()] = _module.__default.safeModuloInt((_this).f14, p0);
      _nw77[(new BigNumber(6)).toNumber()] = _434_v0;
      _nw77[(new BigNumber(7)).toNumber()] = _434_v0;
      _nw77[(new BigNumber(8)).toNumber()] = (_434_v0).plus((_this).f14);
      _nw77[(new BigNumber(9)).toNumber()] = _438_v4.f12;
      _nw77[(new BigNumber(10)).toNumber()] = (_this).f14;
      _nw77[(new BigNumber(11)).toNumber()] = _434_v0;
      _nw77[(new BigNumber(12)).toNumber()] = _module.__default.fm0(globalState);
      _nw77[(new BigNumber(13)).toNumber()] = _434_v0;
      _nw77[(new BigNumber(14)).toNumber()] = _module.__default.safeModuloInt(_438_v4.f12, _438_v4.f12);
      _nw77[(new BigNumber(15)).toNumber()] = _434_v0;
      _nw77[(new BigNumber(16)).toNumber()] = (((_476_v39)[_module.__default.safeIndex(new BigNumber(459), new BigNumber((_476_v39).length))]) ? ((_this).f14) : ((_dafny.ZERO).minus(_module.__default.fm0(globalState))));
      _493_v52 = _nw77;
      let _494_v53;
      let _nw78 = new _module.C1();
      _nw78.__ctor(_module.__default.fm26(globalState));
      _494_v53 = _nw78;
      let _495_v54;
      _495_v54 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(414)), ((_496_v6) => function (_497_i1) {
        return _496_v6;
      })(_440_v6)),_dafny.Map.Empty.slice().updateUnsafe(_494_v53,_438_v4.f12));
      let _498_v55;
      _498_v55 = _dafny.Map.Empty.slice().updateUnsafe(_494_v53,p0);
      let _index55 = _module.__default.safeIndex(new BigNumber(751), new BigNumber((_493_v52).length));
      (_493_v52)[_index55] = new BigNumber(((((_495_v54).contains(_dafny.Seq.UnicodeFromString("nukgadch"))) ? ((_495_v54).get(_dafny.Seq.UnicodeFromString("nukgadch"))) : (_498_v55))).length);
      let _499_v56;
      _499_v56 = _dafny.Seq.of(_dafny.Set.fromElements(_434_v0), _490_v49, _490_v49);
      let _index56 = _module.__default.safeIndex(new BigNumber(751), new BigNumber((_493_v52).length));
      let _rhs97 = (_490_v49).Difference((_499_v56)[_module.__default.safeIndex(_434_v0, new BigNumber((_499_v56).length))]);
      let _rhs98 = p0;
      let _lhs54 = _493_v52;
      let _lhs55 = _module.__default.safeIndex(new BigNumber(751), new BigNumber((_493_v52).length));
      _490_v49 = _rhs97;
      _lhs54[_lhs55] = _rhs98;
      return;
    }
    m8(p0, p1, globalState) {
      let _this = this;
      let _500_v0;
      _500_v0 = _dafny.MultiSet.fromElements((_this).f14, (_this).f14);
      let _501_v1;
      _501_v1 = _dafny.Seq.of(p0);
      let _502_v2;
      _502_v2 = _dafny.Set.fromElements((_this).f14);
      let _503_v3;
      _503_v3 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f9);
      let _504_v4;
      _504_v4 = _dafny.MultiSet.fromElements((((_503_v3).contains((_this).f9)) ? ((_503_v3).get((_this).f9)) : ((_this).f9)), (_this).f9);
      let _505_v5;
      let _nw79 = Array((new BigNumber(3)).toNumber());
      _nw79[(_dafny.ZERO).toNumber()] = (_500_v0).IsProperSubsetOf(_500_v0);
      _nw79[(_dafny.ONE).toNumber()] = ((_this).f14).isLessThanOrEqualTo((_this).f14);
      _nw79[(new BigNumber(2)).toNumber()] = ((_dafny.MultiSet.FromArray(_dafny.Seq.update(_501_v1, _module.__default.safeIndex((_this).f14, new BigNumber((_501_v1).length)), (_this).f9))).update((_this).f9, _module.__default.abs(new BigNumber((_502_v2).length)))).IsDisjointFrom(_504_v4);
      _505_v5 = _nw79;
      let _index57 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((_505_v5).length));
      (_505_v5)[_index57] = (_this).f9;
      let _506_v6;
      let _nw80 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
      _506_v6 = _nw80;
      let _index58 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_506_v6).length));
      (_506_v6)[_index58] = new BigNumber(-89);
      let _507_v7;
      let _nw81 = Array((new BigNumber(12)).toNumber()).fill(_dafny.MultiSet.Empty);
      _507_v7 = _nw81;
      let _508_v8;
      _508_v8 = _dafny.Seq.of(_507_v7);
      let _509_v9;
      let _nw82 = Array((new BigNumber(16)).toNumber());
      _nw82[(_dafny.ZERO).toNumber()] = _507_v7;
      _nw82[(_dafny.ONE).toNumber()] = _507_v7;
      _nw82[(new BigNumber(2)).toNumber()] = _507_v7;
      _nw82[(new BigNumber(3)).toNumber()] = _507_v7;
      _nw82[(new BigNumber(4)).toNumber()] = _507_v7;
      _nw82[(new BigNumber(5)).toNumber()] = _507_v7;
      _nw82[(new BigNumber(6)).toNumber()] = _507_v7;
      _nw82[(new BigNumber(7)).toNumber()] = _507_v7;
      _nw82[(new BigNumber(8)).toNumber()] = (_508_v8)[_module.__default.safeIndex((_this).f14, new BigNumber((_508_v8).length))];
      _nw82[(new BigNumber(9)).toNumber()] = _507_v7;
      _nw82[(new BigNumber(10)).toNumber()] = _507_v7;
      _nw82[(new BigNumber(11)).toNumber()] = _507_v7;
      _nw82[(new BigNumber(12)).toNumber()] = _507_v7;
      _nw82[(new BigNumber(13)).toNumber()] = _507_v7;
      _nw82[(new BigNumber(14)).toNumber()] = _507_v7;
      _nw82[(new BigNumber(15)).toNumber()] = _507_v7;
      _509_v9 = _nw82;
      let _index59 = _module.__default.safeIndex(new BigNumber(218), new BigNumber((_509_v9).length));
      (_509_v9)[_index59] = _507_v7;
      let _510_v10;
      _510_v10 = _dafny.Seq.of((_this).f14);
      let _index60 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((_505_v5).length));
      let _index61 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_506_v6).length));
      let _index62 = _module.__default.safeIndex(new BigNumber(218), new BigNumber((_509_v9).length));
      let _rhs99 = (((_this).f9) ? ((_this).f9) : (p0));
      let _rhs100 = _module.__default.fm1(_500_v0, (_dafny.ZERO).minus((_this).f14), globalState);
      let _rhs101 = _module.__default.safeDivisionInt((_this).f14, (_dafny.ZERO).minus((new BigNumber((_510_v10).length)).minus((_this).f14)));
      let _rhs102 = _507_v7;
      let _lhs56 = globalState;
      let _lhs57 = _505_v5;
      let _lhs58 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((_505_v5).length));
      let _lhs59 = _506_v6;
      let _lhs60 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_506_v6).length));
      let _lhs61 = _509_v9;
      let _lhs62 = _module.__default.safeIndex(new BigNumber(218), new BigNumber((_509_v9).length));
      _lhs56.f2 = _rhs99;
      _lhs57[_lhs58] = _rhs100;
      _lhs59[_lhs60] = _rhs101;
      _lhs61[_lhs62] = _rhs102;
      let _index63 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_506_v6).length));
      (_506_v6)[_index63] = (_506_v6)[_module.__default.safeIndex(new BigNumber(781), new BigNumber((_506_v6).length))];
      let _511_v11;
      _511_v11 = _module.D1.create_DC5(_module.D1.create_DC4((_this).f14, (_this).f9, true, (_module.D3.create_DC9((_this).f9, _dafny.Seq.of(p0, p0))).dtor_cf17, (_this).f9));
      let _512_v12;
      _512_v12 = _module.D1.create_DC5(_511_v11);
      let _pat_let_tv9 = _511_v11;
      let _pat_let_tv10 = _511_v11;
      let _513_v13;
      _513_v13 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(function (_pat_let11_0) {
        return function (_514_dt__update__tmp_h0) {
          return function (_pat_let12_0) {
            return function (_515_dt__update_hcf14_h0) {
              return _module.D1.create_DC5(_515_dt__update_hcf14_h0);
            }(_pat_let12_0);
          }(_pat_let_tv9);
        }(_pat_let11_0);
      }(_512_v12), _512_v12, function (_pat_let13_0) {
        return function (_516_dt__update__tmp_h1) {
          return function (_pat_let14_0) {
            return function (_517_dt__update_hcf14_h1) {
              return _module.D1.create_DC5(_517_dt__update_hcf14_h1);
            }(_pat_let14_0);
          }(_pat_let_tv10);
        }(_pat_let13_0);
      }(_512_v12)),_500_v0);
      let _arr0 = (_509_v9)[_module.__default.safeIndex(new BigNumber(218), new BigNumber((_509_v9).length))];
      let _index64 = _module.__default.safeIndex(new BigNumber(912), new BigNumber(((_509_v9)[_module.__default.safeIndex(new BigNumber(218), new BigNumber((_509_v9).length))]).length));
      _arr0[_index64] = (((_513_v13).contains(_dafny.MultiSet.fromElements(_512_v12, _512_v12, _512_v12))) ? ((_513_v13).get(_dafny.MultiSet.fromElements(_512_v12, _512_v12, _512_v12))) : (_dafny.MultiSet.fromElements((_this).f14, (_506_v6)[_module.__default.safeIndex(new BigNumber(781), new BigNumber((_506_v6).length))])));
      let _arr1 = (_509_v9)[_module.__default.safeIndex(new BigNumber(218), new BigNumber((_509_v9).length))];
      let _index65 = _module.__default.safeIndex(new BigNumber(912), new BigNumber(((_509_v9)[_module.__default.safeIndex(new BigNumber(218), new BigNumber((_509_v9).length))]).length));
      _arr1[_index65] = (_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_510_v10, _dafny.Seq.of((_this).f14)))).Difference(_500_v0);
      let _518_i0;
      _518_i0 = _dafny.ZERO;
      L5: {
        while ((_505_v5)[_module.__default.safeIndex(new BigNumber(557), new BigNumber((_505_v5).length))]) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_518_i0)) {
              break L5;
            }
            _518_i0 = (_518_i0).plus(_dafny.ONE);
            let _519_v14;
            _519_v14 = _dafny.Seq.UnicodeFromString("a");
            let _520_v15;
            _520_v15 = _module.D1.create_DC4(new BigNumber((_519_v14).length), !(p0), p0, (_505_v5)[_module.__default.safeIndex(new BigNumber(557), new BigNumber((_505_v5).length))], (_505_v5)[_module.__default.safeIndex(new BigNumber(557), new BigNumber((_505_v5).length))]);
            if (_module.__default.fm1((((_509_v9)[_module.__default.safeIndex(new BigNumber(218), new BigNumber((_509_v9).length))])[_module.__default.safeIndex(new BigNumber(912), new BigNumber(((_509_v9)[_module.__default.safeIndex(new BigNumber(218), new BigNumber((_509_v9).length))]).length))]).Union(_500_v0), _module.__default.safeDivisionInt((_520_v15).dtor_cf9, (_this).f14), globalState)) {
              let _521_v16;
              _521_v16 = _dafny.Map.Empty.slice().updateUnsafe(_500_v0,(_this).f9);
              let _index66 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_506_v6).length));
              (_506_v6)[_index66] = ((_506_v6)[_module.__default.safeIndex(new BigNumber(781), new BigNumber((_506_v6).length))]).plus((_dafny.ZERO).minus((new BigNumber((_519_v14).length)).minus(new BigNumber((_521_v16).length))));
              let _522_v17;
              _522_v17 = _module.D3.create_DC9((_this).f9, _501_v1);
              let _523_v18;
              _523_v18 = _dafny.MultiSet.fromElements(_505_v5);
              let _524_v19;
              _524_v19 = _module.D3.create_DC7(_523_v18);
              let _525_v20;
              _525_v20 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_519_v14, _519_v14),p0);
              let _526_v21;
              _526_v21 = _module.D9.create_DC23(_525_v20);
              let _pat_let_tv11 = _523_v18;
              let _index67 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_506_v6).length));
              let _rhs103 = ((_506_v6)[_module.__default.safeIndex(new BigNumber(781), new BigNumber((_506_v6).length))]).multipliedBy((_506_v6)[_module.__default.safeIndex(new BigNumber(781), new BigNumber((_506_v6).length))]);
              let _rhs104 = _522_v17;
              let _rhs105 = function (_pat_let15_0) {
                return function (_527_dt__update__tmp_h2) {
                  return function (_pat_let16_0) {
                    return function (_528_dt__update_hcf16_h0) {
                      return _module.D3.create_DC7(_528_dt__update_hcf16_h0);
                    }(_pat_let16_0);
                  }(_pat_let_tv11);
                }(_pat_let15_0);
              }(_524_v19);
              let _rhs106 = (_526_v21).dtor_cf40;
              let _lhs63 = _506_v6;
              let _lhs64 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_506_v6).length));
              _lhs63[_lhs64] = _rhs103;
              _522_v17 = _rhs104;
              _524_v19 = _rhs105;
              _525_v20 = _rhs106;
              let _529_v22;
              _529_v22 = new _dafny.CodePoint('e'.codePointAt(0));
              let _530_v23;
              _530_v23 = _dafny.Seq.of(_dafny.Seq.update(_519_v14, _module.__default.safeIndex((_this).f14, new BigNumber((_519_v14).length)), _529_v22));
              let _531_v24;
              _531_v24 = _dafny.Map.Empty.slice().updateUnsafe((_500_v0).update((_506_v6)[_module.__default.safeIndex(new BigNumber(781), new BigNumber((_506_v6).length))], _module.__default.abs(new BigNumber((_530_v23).length))),(_dafny.ZERO).minus((_506_v6)[_module.__default.safeIndex(new BigNumber(781), new BigNumber((_506_v6).length))]));
              let _532_v25;
              _532_v25 = _dafny.Map.Empty.slice().updateUnsafe(_531_v24,(_this).f9);
              _532_v25 = (_532_v25).update(_dafny.Map.Empty.slice().updateUnsafe(((_509_v9)[_module.__default.safeIndex(new BigNumber(218), new BigNumber((_509_v9).length))])[_module.__default.safeIndex(new BigNumber(912), new BigNumber(((_509_v9)[_module.__default.safeIndex(new BigNumber(218), new BigNumber((_509_v9).length))]).length))],(_506_v6)[_module.__default.safeIndex(new BigNumber(781), new BigNumber((_506_v6).length))]), (_this).f9);
              let _533_v26;
              let _nw83 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Set.Empty);
              _533_v26 = _nw83;
              let _534_v27;
              _534_v27 = _dafny.Map.Empty.slice().updateUnsafe((_506_v6)[_module.__default.safeIndex(new BigNumber(781), new BigNumber((_506_v6).length))],_529_v22);
              let _535_v28;
              _535_v28 = _dafny.Set.fromElements(_534_v27, _534_v27, _534_v27);
              let _index68 = _module.__default.safeIndex(new BigNumber(435), new BigNumber((_533_v26).length));
              (_533_v26)[_index68] = _535_v28;
              let _536_v29;
              let _init12 = ((_537_v14) => function (_538_i1) {
                return _537_v14;
              })(_519_v14);
              let _nw84 = Array((new BigNumber(15)).toNumber());
              for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw84.length); _i0_12++) {
                _nw84[_i0_12] = _init12(new BigNumber(_i0_12));
              }
              _536_v29 = _nw84;
              let _index69 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_536_v29).length));
              (_536_v29)[_index69] = _519_v14;
              let _index70 = _module.__default.safeIndex(new BigNumber(435), new BigNumber((_533_v26).length));
              let _index71 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_536_v29).length));
              let _index72 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((_505_v5).length));
              let _rhs107 = (_505_v5)[_module.__default.safeIndex(new BigNumber(557), new BigNumber((_505_v5).length))];
              let _rhs108 = (_535_v28).Intersect((_535_v28).Intersect(_535_v28));
              let _rhs109 = _519_v14;
              let _rhs110 = (((_this).f9) ? (new _dafny.CodePoint('r'.codePointAt(0))) : (_529_v22));
              let _rhs111 = !((((_505_v5)[_module.__default.safeIndex(new BigNumber(557), new BigNumber((_505_v5).length))]) ? (p0) : ((_505_v5)[_module.__default.safeIndex(new BigNumber(557), new BigNumber((_505_v5).length))])));
              let _lhs65 = globalState;
              let _lhs66 = _533_v26;
              let _lhs67 = _module.__default.safeIndex(new BigNumber(435), new BigNumber((_533_v26).length));
              let _lhs68 = _536_v29;
              let _lhs69 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_536_v29).length));
              let _lhs70 = _505_v5;
              let _lhs71 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((_505_v5).length));
              _lhs65.f2 = _rhs107;
              _lhs66[_lhs67] = _rhs108;
              _lhs68[_lhs69] = _rhs109;
              _529_v22 = _rhs110;
              _lhs70[_lhs71] = _rhs111;
              (globalState).f2 = (_505_v5)[_module.__default.safeIndex(new BigNumber(557), new BigNumber((_505_v5).length))];
            } else {
              _519_v14 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(337)), function (_539_i2) {
                return new _dafny.CodePoint('k'.codePointAt(0));
              }), _519_v14);
              let _540_v30;
              _540_v30 = new _dafny.CodePoint('w'.codePointAt(0));
              (globalState).f2 = !(_dafny.areEqual(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.update(_519_v14, _module.__default.safeIndex(new BigNumber((_519_v14).length), new BigNumber((_519_v14).length)), _540_v30), _dafny.Seq.Create(_module.__default.abs(new BigNumber(766)), ((_541_v30) => function (_542_i3) {
                return _541_v30;
              })(_540_v30))), _module.__default.safeIndex((_this).f14, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_519_v14, _module.__default.safeIndex(new BigNumber((_519_v14).length), new BigNumber((_519_v14).length)), _540_v30), _dafny.Seq.Create(_module.__default.abs(new BigNumber(766)), ((_543_v30) => function (_544_i3) {
                return _543_v30;
              })(_540_v30)))).length)), _540_v30), _dafny.Seq.Concat(_519_v14, _519_v14)));
              let _545_v31;
              _545_v31 = _506_v6;
              let _546_v32;
              _546_v32 = _dafny.Map.Empty.slice().updateUnsafe((_505_v5)[_module.__default.safeIndex(new BigNumber(557), new BigNumber((_505_v5).length))],(_this).f14);
              let _index73 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_506_v6).length));
              let _rhs112 = (((_546_v32).contains(p0)) ? ((_546_v32).get(p0)) : (new BigNumber(738)));
              let _rhs113 = _545_v31;
              let _lhs72 = _506_v6;
              let _lhs73 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_506_v6).length));
              _lhs72[_lhs73] = _rhs112;
              _545_v31 = _rhs113;
              let _547_v33;
              let _nw85 = new _module.C1();
              _nw85.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(324)), ((_548_p0, _549_v6) => function (_550_i4) {
                return _module.D1.create_DC2(new BigNumber((_dafny.Set.fromElements(_548_p0)).length), _548_p0, (_549_v6)[_module.__default.safeIndex(new BigNumber(781), new BigNumber((_549_v6).length))], (_this).f14);
              })(p0, _506_v6)));
              _547_v33 = _nw85;
              let _551_v34;
              _551_v34 = _dafny.MultiSet.fromElements(_547_v33);
              let _index74 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_506_v6).length));
              (_506_v6)[_index74] = (((_551_v34).IsSubsetOf(_551_v34)) ? ((_this).f14) : ((_this).f14));
              let _552_v35;
              _552_v35 = _dafny.Set.fromElements(_506_v6);
              let _553_v36;
              _553_v36 = _dafny.Map.Empty.slice().updateUnsafe((_552_v35).Intersect(_552_v35),(_this).f14);
              let _index75 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_506_v6).length));
              (_506_v6)[_index75] = (((_553_v36).contains(_552_v35)) ? ((_553_v36).get(_552_v35)) : ((_506_v6)[_module.__default.safeIndex(new BigNumber(781), new BigNumber((_506_v6).length))]));
            }
            _506_v6 = _506_v6;
            let _554_v37;
            _554_v37 = _module.D3.create_DC10();
            _554_v37 = _module.__default.fm27(globalState);
            (globalState).f2 = (_this).f9;
          }
        }
      }
      let _555_v38;
      _555_v38 = new _dafny.CodePoint('k'.codePointAt(0));
      let _556_v39;
      _556_v39 = _dafny.Set.fromElements(_555_v38, new _dafny.CodePoint('o'.codePointAt(0)), _555_v38, _555_v38, _555_v38);
      _556_v39 = _556_v39;
      let _557_v40;
      _557_v40 = _module.D4.create_DC14((_505_v5)[_module.__default.safeIndex(new BigNumber(557), new BigNumber((_505_v5).length))], false, (_506_v6)[_module.__default.safeIndex(new BigNumber(781), new BigNumber((_506_v6).length))]);
      let _558_v41;
      _558_v41 = _dafny.MultiSet.fromElements(_555_v38, _555_v38);
      let _559_v42;
      _559_v42 = _module.D1.create_DC3((_this).f9, _module.__default.fm28(new BigNumber((_558_v41).cardinality()), true, globalState), (_506_v6)[_module.__default.safeIndex(new BigNumber(781), new BigNumber((_506_v6).length))]);
      let _560_v43;
      _560_v43 = _dafny.MultiSet.fromElements(_559_v42);
      let _index76 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((_505_v5).length));
      let _rhs114 = _module.D4.create_DC14((((_this).f9) ? ((_505_v5)[_module.__default.safeIndex(new BigNumber(557), new BigNumber((_505_v5).length))]) : (!(p0))), (_505_v5)[_module.__default.safeIndex(new BigNumber(557), new BigNumber((_505_v5).length))], (_506_v6)[_module.__default.safeIndex(new BigNumber(781), new BigNumber((_506_v6).length))]);
      let _rhs115 = (_dafny.MultiSet.fromElements(_559_v42)).update(_559_v42, _module.__default.abs((_this).f14));
      let _rhs116 = p0;
      let _lhs74 = _505_v5;
      let _lhs75 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((_505_v5).length));
      _557_v40 = _rhs114;
      _560_v43 = _rhs115;
      _lhs74[_lhs75] = _rhs116;
      return;
    }
    get f14() {
      let _this = this;
      return _this._f14;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f9 = false;
      this.f16 = [];
      this._f15 = false;
    }
    _parentTraits() {
      return [_module.T2, _module.T1, _module.T0];
    }
    get f9() {
      let _this = this;
      return _this._f9;
    };
    __ctor(f15, f16, f9) {
      let _this = this;
      (_this)._f15 = f15;
      (_this).f16 = f16;
      (_this)._f9 = f9;
      return;
    }
    fm3(globalState) {
      let _this = this;
      return (_module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-746)), function (_561_i0) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      })).length)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.D10.create_DC28(),(_this).f9)).length))).isLessThanOrEqualTo(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_this).f15,new BigNumber((_dafny.Seq.UnicodeFromString("kthtybr")).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f15,new BigNumber((function () {
        let _coll20 = new _dafny.Set();
        for (const _compr_20 of (_dafny.MultiSet.fromElements(_dafny.Set.fromElements(true, (_this).f9), _dafny.Set.fromElements(false))).Elements) {
          let _562_v0 = _compr_20;
          if ((_dafny.MultiSet.fromElements(_dafny.Set.fromElements(true, (_this).f9), _dafny.Set.fromElements(false))).contains(_562_v0)) {
            _coll20.add(_562_v0);
          }
        }
        return _coll20;
      }()).length)))).length));
    };
    fm2(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(true)).length))), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(286)), function (_563_i0) {
        return new BigNumber(-817);
      }), _dafny.Seq.of(new BigNumber(611))));
    };
    fm32(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !(!((_this).f9));
    };
    fm33(globalState) {
      let _this = this;
      return _dafny.Seq.UnicodeFromString("chkqh");
    };
    m2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let r3 = _dafny.Set.Empty;
      let _564_v0;
      _564_v0 = new BigNumber(418);
      let _565_v1;
      _565_v1 = _dafny.MultiSet.fromElements(_564_v0);
      let _566_v2;
      _566_v2 = _dafny.Map.Empty.slice().updateUnsafe(_565_v1,(_this).f9);
      let _567_v3;
      _567_v3 = _module.D7.create_DC18(_566_v2);
      _567_v3 = _567_v3;
      let _568_v4;
      let _nw86 = Array((new BigNumber(11)).toNumber());
      _nw86[(_dafny.ZERO).toNumber()] = (_this).f15;
      _nw86[(_dafny.ONE).toNumber()] = (_this).f15;
      _nw86[(new BigNumber(2)).toNumber()] = (_this).f15;
      _nw86[(new BigNumber(3)).toNumber()] = (_this).f9;
      _nw86[(new BigNumber(4)).toNumber()] = (_this).f15;
      _nw86[(new BigNumber(5)).toNumber()] = (_this).f15;
      _nw86[(new BigNumber(6)).toNumber()] = (_this).f9;
      _nw86[(new BigNumber(7)).toNumber()] = (_this).f15;
      _nw86[(new BigNumber(8)).toNumber()] = (_this).f9;
      _nw86[(new BigNumber(9)).toNumber()] = (_this).f15;
      _nw86[(new BigNumber(10)).toNumber()] = (_this).f9;
      _568_v4 = _nw86;
      let _569_v5;
      _569_v5 = _dafny.MultiSet.fromElements(p1, p1);
      let _index77 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_568_v4).length));
      (_568_v4)[_index77] = !(!(!(_569_v5).contains(p1)));
      let _index78 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_568_v4).length));
      (_568_v4)[_index78] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("jrugtgog"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(882)), ((_570_p2) => function (_571_i0) {
        return _570_p2;
      })(p2)));
      let _index79 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_568_v4).length));
      (_568_v4)[_index79] = (_this).f15;
      let _572_v6;
      _572_v6 = _dafny.Set.fromElements(_564_v0, _564_v0);
      let _573_i1;
      _573_i1 = _dafny.ZERO;
      L6: {
        while ((_572_v6).IsProperSubsetOf(_572_v6)) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_573_i1)) {
              break L6;
            }
            _573_i1 = (_573_i1).plus(_dafny.ONE);
            let _574_v7;
            let _init13 = ((_575_v0) => function (_576_i2) {
              return (_576_i2).minus(_575_v0);
            })(_564_v0);
            let _nw87 = Array((new BigNumber(15)).toNumber());
            for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw87.length); _i0_13++) {
              _nw87[_i0_13] = _init13(new BigNumber(_i0_13));
            }
            _574_v7 = _nw87;
            let _index80 = _module.__default.safeIndex(new BigNumber(527), new BigNumber((_574_v7).length));
            (_574_v7)[_index80] = _564_v0;
            let _index81 = _module.__default.safeIndex(new BigNumber(527), new BigNumber((_574_v7).length));
            (_574_v7)[_index81] = _564_v0;
            let _577_v8;
            _577_v8 = _module.D9.create_DC25(_564_v0, (_574_v7)[_module.__default.safeIndex(new BigNumber(527), new BigNumber((_574_v7).length))], _565_v1, _dafny.Map.Empty.slice().updateUnsafe((_574_v7)[_module.__default.safeIndex(new BigNumber(527), new BigNumber((_574_v7).length))],p2), new BigNumber(-777));
            let _578_v9;
            _578_v9 = _dafny.Map.Empty.slice().updateUnsafe((_574_v7)[_module.__default.safeIndex(new BigNumber(527), new BigNumber((_574_v7).length))],_577_v8);
            _578_v9 = _578_v9;
            let _579_v10;
            _579_v10 = _module.D9.create_DC24(false, (_574_v7)[_module.__default.safeIndex(new BigNumber(527), new BigNumber((_574_v7).length))], _dafny.Seq.UnicodeFromString("prdj"));
            let _580_v11;
            _580_v11 = _dafny.Set.fromElements(p1, p1, (_579_v10).dtor_cf43);
            let _581_v12;
            _581_v12 = _dafny.Seq.of(new BigNumber((_580_v11).length), new BigNumber((_dafny.Seq.of(p2, p2)).length));
            let _582_v13;
            _582_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_564_v0);
            let _583_v14;
            _583_v14 = _dafny.Seq.of(_dafny.Seq.of(new BigNumber((_582_v13).length)), p3, p3, p3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(775)), ((_584_v0) => function (_585_i3) {
              return _584_v0;
            })(_564_v0)));
            _581_v12 = _dafny.Seq.Concat((_583_v14)[_module.__default.safeIndex(_564_v0, new BigNumber((_583_v14).length))], p3);
            let _586_v15;
            _586_v15 = _dafny.Seq.of(p1);
            let _587_v16;
            _587_v16 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f15) || ((_this).f15),_586_v15);
            let _588_v17;
            _588_v17 = _dafny.Seq.of((_this).f15);
            let _589_v18;
            _589_v18 = _dafny.Map.Empty.slice().updateUnsafe(_564_v0,new BigNumber(859));
            _587_v16 = (_587_v16).update((_588_v17)[_module.__default.safeIndex(new BigNumber((_589_v18).length), new BigNumber((_588_v17).length))], _dafny.Seq.update(_module.__default.fm34((_dafny.ZERO).minus(_564_v0), (_568_v4)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((_568_v4).length))], (_this).f15, globalState), _module.__default.safeIndex(_564_v0, new BigNumber((_module.__default.fm34((_dafny.ZERO).minus(_564_v0), (_568_v4)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((_568_v4).length))], (_this).f15, globalState)).length)), _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("sooucaqvp"), _module.__default.safeIndex(_564_v0, new BigNumber((_dafny.Seq.UnicodeFromString("sooucaqvp")).length)), p2), _module.__default.safeIndex(_564_v0, new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("sooucaqvp"), _module.__default.safeIndex(_564_v0, new BigNumber((_dafny.Seq.UnicodeFromString("sooucaqvp")).length)), p2)).length)), p2)));
          }
        }
      }
      if ((_568_v4)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((_568_v4).length))]) {
        let _590_v19;
        _590_v19 = _dafny.Seq.of(_module.__default.fm0(globalState));
        let _591_v20;
        _591_v20 = new _dafny.CodePoint('s'.codePointAt(0));
        let _592_v21;
        _592_v21 = _dafny.Seq.of(true);
        let _rhs117 = _dafny.Seq.Concat(_dafny.Seq.of(_564_v0), p3);
        let _rhs118 = p2;
        let _rhs119 = _592_v21;
        _590_v19 = _rhs117;
        _591_v20 = _rhs118;
        _592_v21 = _rhs119;
        if ((_this).f9) {
          let _593_v22;
          _593_v22 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_564_v0).minus(_564_v0)),(_this).f15);
          _593_v22 = (_593_v22).update(_564_v0, (_568_v4)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((_568_v4).length))]);
          _568_v4 = _568_v4;
          let _594_v23;
          let _nw88 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
          _594_v23 = _nw88;
          let _index82 = _module.__default.safeIndex(new BigNumber(233), new BigNumber((_594_v23).length));
          (_594_v23)[_index82] = _564_v0;
          let _index83 = _module.__default.safeIndex(new BigNumber(233), new BigNumber((_594_v23).length));
          (_594_v23)[_index83] = _564_v0;
          (globalState).f2 = (_this).f9;
          r2 = _564_v0;
        } else {
          let _595_v24;
          _595_v24 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-632)), ((_596_v0) => function (_597_i4) {
            return _596_v0;
          })(_564_v0))), new BigNumber(308), globalState),(_568_v4)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((_568_v4).length))]);
          r2 = new BigNumber((_595_v24).length);
          let _598_v25;
          let _nw89 = Array((new BigNumber(22)).toNumber());
          _nw89[(_dafny.ZERO).toNumber()] = p1;
          _nw89[(_dafny.ONE).toNumber()] = _dafny.Seq.UnicodeFromString("qvqafp");
          _nw89[(new BigNumber(2)).toNumber()] = _dafny.Seq.update(p1, _module.__default.safeIndex(_564_v0, new BigNumber((p1).length)), _591_v20);
          _nw89[(new BigNumber(3)).toNumber()] = _dafny.Seq.UnicodeFromString("n");
          _nw89[(new BigNumber(4)).toNumber()] = p1;
          _nw89[(new BigNumber(5)).toNumber()] = p1;
          _nw89[(new BigNumber(6)).toNumber()] = p1;
          _nw89[(new BigNumber(7)).toNumber()] = p1;
          _nw89[(new BigNumber(8)).toNumber()] = p1;
          _nw89[(new BigNumber(9)).toNumber()] = p1;
          _nw89[(new BigNumber(10)).toNumber()] = p1;
          _nw89[(new BigNumber(11)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(51)), ((_599_p2) => function (_600_i5) {
            return _599_p2;
          })(p2));
          _nw89[(new BigNumber(12)).toNumber()] = (_this).fm33(globalState);
          _nw89[(new BigNumber(13)).toNumber()] = p1;
          _nw89[(new BigNumber(14)).toNumber()] = (_this).fm33(globalState);
          _nw89[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("osyhtuwn"), p1);
          _nw89[(new BigNumber(16)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(p1, p1), _module.__default.safeIndex(_564_v0, new BigNumber((_dafny.Seq.Concat(p1, p1)).length)), p2);
          _nw89[(new BigNumber(17)).toNumber()] = _dafny.Seq.UnicodeFromString("iqiqgyo");
          _nw89[(new BigNumber(18)).toNumber()] = _dafny.Seq.UnicodeFromString("amy");
          _nw89[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(p1, (_this).fm33(globalState));
          _nw89[(new BigNumber(20)).toNumber()] = p1;
          _nw89[(new BigNumber(21)).toNumber()] = _dafny.Seq.Concat(p1, p1);
          _598_v25 = _nw89;
          _598_v25 = _598_v25;
          let _601_v26;
          _601_v26 = _dafny.Set.fromElements(false, true);
          (globalState).f2 = (_601_v26).IsSubsetOf((_601_v26).Intersect(_601_v26));
          let _602_v27;
          let _nw90 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
          _602_v27 = _nw90;
          _602_v27 = _602_v27;
          let _index84 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_568_v4).length));
          (_568_v4)[_index84] = false;
        }
        let _603_v28;
        _603_v28 = _dafny.Map.Empty.slice().updateUnsafe(_564_v0,new BigNumber(986));
        let _604_v29;
        _604_v29 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_564_v0);
        let _605_v30;
        let _nw91 = Array((new BigNumber(9)).toNumber());
        _nw91[(_dafny.ZERO).toNumber()] = (((_603_v28).contains(_564_v0)) ? ((_603_v28).get(_564_v0)) : (_564_v0));
        _nw91[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus((((_this).f15) ? (new BigNumber((_dafny.Seq.of(new BigNumber((_604_v29).length))).length)) : (_564_v0)))));
        _nw91[(new BigNumber(2)).toNumber()] = _564_v0;
        _nw91[(new BigNumber(3)).toNumber()] = _564_v0;
        _nw91[(new BigNumber(4)).toNumber()] = _564_v0;
        _nw91[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(899)), ((_606_p1) => function (_607_i6) {
          return (_module.D1.create_DC3((_this).f9, _606_p1, new BigNumber(973))).dtor_cf7;
        })(p1))).length);
        _nw91[(new BigNumber(6)).toNumber()] = _564_v0;
        _nw91[(new BigNumber(7)).toNumber()] = _564_v0;
        _nw91[(new BigNumber(8)).toNumber()] = _module.__default.fm0(globalState);
        _605_v30 = _nw91;
        let _nw92 = Array((_dafny.ONE).toNumber());
        _nw92[(_dafny.ZERO).toNumber()] = _module.__default.safeDivisionInt(_564_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_564_v0,(_this).f15)).length));
        _605_v30 = _nw92;
        r0 = !(!((_568_v4)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((_568_v4).length))]) || (!((_this).f9)));
        (globalState).f2 = !_dafny.Seq.contains(_dafny.Seq.Concat(p1, (_this).fm33(globalState)), (((_this).f9) ? ((p1)[_module.__default.safeIndex(_564_v0, new BigNumber((p1).length))]) : (p2)));
      } else {
        if (!(false)) {
          r1 = ((_564_v0).multipliedBy(_564_v0)).isLessThanOrEqualTo(_564_v0);
          let _608_v31;
          _608_v31 = _dafny.Map.Empty.slice().updateUnsafe(_568_v4,_564_v0);
          let _609_v32;
          let _nw93 = Array((new BigNumber(27)).toNumber()).fill([]);
          _609_v32 = _nw93;
          let _610_v33;
          _610_v33 = _609_v32;
          let _611_v34;
          _611_v34 = _dafny.Set.fromElements(_610_v33);
          let _612_v35;
          let _nw94 = Array((new BigNumber(9)).toNumber());
          _nw94[(_dafny.ZERO).toNumber()] = _564_v0;
          _nw94[(_dafny.ONE).toNumber()] = new BigNumber((_611_v34).length);
          _nw94[(new BigNumber(2)).toNumber()] = _564_v0;
          _nw94[(new BigNumber(3)).toNumber()] = _564_v0;
          _nw94[(new BigNumber(4)).toNumber()] = _564_v0;
          _nw94[(new BigNumber(5)).toNumber()] = _564_v0;
          _nw94[(new BigNumber(6)).toNumber()] = _564_v0;
          _nw94[(new BigNumber(7)).toNumber()] = new BigNumber(323);
          _nw94[(new BigNumber(8)).toNumber()] = _564_v0;
          _612_v35 = _nw94;
          let _613_v36;
          let _nw95 = new _module.C2();
          _nw95.__ctor(_564_v0, (_568_v4)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((_568_v4).length))]);
          _613_v36 = _nw95;
          let _614_v37;
          _614_v37 = _dafny.Seq.of(_613_v36, _613_v36, _613_v36);
          let _615_v38;
          _615_v38 = _dafny.Map.Empty.slice().updateUnsafe(_612_v35,new BigNumber((_614_v37).length));
          let _616_v39;
          _616_v39 = _module.D11.create_DC29(_615_v38);
          let _rhs120 = ((_dafny.Map.Empty.slice().updateUnsafe(_568_v4,_564_v0)).Merge(_608_v31)).Merge((((_568_v4)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((_568_v4).length))]) ? (_dafny.Map.Empty.slice().updateUnsafe(_568_v4,(_613_v36).f14)) : (_608_v31)));
          let _rhs121 = (_616_v39).dtor_cf51;
          let _rhs122 = ((((_613_v36).f14).isLessThanOrEqualTo((_613_v36).f14)) ? ((_613_v36).f14) : ((_613_v36).f14));
          _608_v31 = _rhs120;
          _615_v38 = _rhs121;
          _564_v0 = _rhs122;
          let _617_v40;
          _617_v40 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_564_v0),_564_v0);
          let _618_v41;
          _618_v41 = _dafny.Seq.of(_617_v40, _617_v40, _617_v40);
          let _619_v42;
          _619_v42 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_618_v41).length)),(_this).f15);
          (globalState).f2 = ((((_568_v4)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((_568_v4).length))]) === ((_this).f9)) ? ((((_619_v42).contains(new BigNumber(22))) ? ((_619_v42).get(new BigNumber(22))) : ((_this).f9))) : (!((_this).f15)));
          let _620_v43;
          let _nw96 = Array((new BigNumber(13)).toNumber());
          _nw96[(_dafny.ZERO).toNumber()] = p1;
          _nw96[(_dafny.ONE).toNumber()] = p1;
          _nw96[(new BigNumber(2)).toNumber()] = p1;
          _nw96[(new BigNumber(3)).toNumber()] = p1;
          _nw96[(new BigNumber(4)).toNumber()] = p1;
          _nw96[(new BigNumber(5)).toNumber()] = p1;
          _nw96[(new BigNumber(6)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.update(p1, _module.__default.safeIndex((_613_v36).f14, new BigNumber((p1).length)), p2), p1), _module.__default.safeIndex(_564_v0, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(p1, _module.__default.safeIndex((_613_v36).f14, new BigNumber((p1).length)), p2), p1)).length)), p2);
          _nw96[(new BigNumber(7)).toNumber()] = p1;
          _nw96[(new BigNumber(8)).toNumber()] = p1;
          _nw96[(new BigNumber(9)).toNumber()] = p1;
          _nw96[(new BigNumber(10)).toNumber()] = _dafny.Seq.UnicodeFromString("ofiikthw");
          _nw96[(new BigNumber(11)).toNumber()] = p1;
          _nw96[(new BigNumber(12)).toNumber()] = _dafny.Seq.update(p1, _module.__default.safeIndex(_564_v0, new BigNumber((p1).length)), p2);
          _620_v43 = _nw96;
          _620_v43 = _620_v43;
          _565_v1 = _565_v1;
        } else {
          let _index85 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_568_v4).length));
          (_568_v4)[_index85] = (_565_v1).IsSubsetOf(_565_v1);
          let _index86 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_568_v4).length));
          (_568_v4)[_index86] = (_this).f15;
          let _621_v44;
          _621_v44 = _dafny.Map.Empty.slice().updateUnsafe((_568_v4)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((_568_v4).length))],_568_v4);
          r2 = _module.__default.safeDivisionInt(new BigNumber(((_621_v44).Merge(_621_v44)).length), (_564_v0).multipliedBy((_dafny.ZERO).minus(_564_v0)));
          let _622_v45;
          _622_v45 = _dafny.Seq.of(_dafny.Seq.Concat(p1, p1), p1, p1, p1, p1);
          r2 = new BigNumber(((_622_v45)[_module.__default.safeIndex((((_565_v1).contains(new BigNumber((p3).length))) ? ((_565_v1).get(new BigNumber((p3).length))) : (new BigNumber(275))), new BigNumber((_622_v45).length))]).length);
          let _623_v46;
          _623_v46 = _module.D1.create_DC2(_564_v0, (_this).f9, new BigNumber((p1).length), _564_v0);
          let _624_v47;
          _624_v47 = _dafny.Seq.of(_623_v46, _623_v46, _623_v46, _623_v46);
          let _625_v48;
          _625_v48 = _module.D12.create_DC31(_624_v47);
          let _626_v49;
          let _nw97 = new _module.C1();
          _nw97.__ctor(_dafny.Seq.update((_625_v48).dtor_cf52, _module.__default.safeIndex(_564_v0, new BigNumber(((_625_v48).dtor_cf52).length)), _623_v46));
          _626_v49 = _nw97;
        }
        let _627_v50;
        _627_v50 = new _dafny.CodePoint('n'.codePointAt(0));
        let _rhs123 = _564_v0;
        let _rhs124 = (_this).f9;
        let _rhs125 = (_this).f9;
        let _rhs126 = (p1)[_module.__default.safeIndex(_564_v0, new BigNumber((p1).length))];
        _564_v0 = _rhs123;
        r0 = _rhs124;
        r1 = _rhs125;
        _627_v50 = _rhs126;
        r2 = (_564_v0).multipliedBy(_564_v0);
        _module.__default.m0((_568_v4)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((_568_v4).length))], globalState);
        r0 = (((((_this).f15) ? ((_this).f15) : ((_this).f9))) ? ((_this).f15) : ((_this).f15));
      }
      let _628_v51;
      _628_v51 = _dafny.Seq.of((_this).f9, (_568_v4)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((_568_v4).length))]);
      let _629_v52;
      _629_v52 = _dafny.Set.fromElements((_568_v4)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((_568_v4).length))], (_628_v51)[_module.__default.safeIndex(_564_v0, new BigNumber((_628_v51).length))], (_this).f15, (((_568_v4)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((_568_v4).length))]) ? (true) : ((_this).f15)), (_568_v4)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((_568_v4).length))]);
      _629_v52 = (_629_v52).Union(_629_v52);
      r0 = (_568_v4)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((_568_v4).length))];
      r1 = (_568_v4)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((_568_v4).length))];
      r2 = _module.__default.safeDivisionInt(_564_v0, new BigNumber(186));
      r3 = (_572_v6).Difference(_572_v6);
      return [r0, r1, r2, r3];
    }
    m3(p0, globalState) {
      let _this = this;
      let r0 = false;
      let _630_v0;
      let _nw98 = new _module.C2();
      _nw98.__ctor(new BigNumber(660), (_this).f9);
      _630_v0 = _nw98;
      (globalState).f2 = (_this).f15;
      let _631_v1;
      _631_v1 = _dafny.Set.fromElements(new BigNumber(-456), (_630_v0).f14);
      let _632_v2;
      _632_v2 = _dafny.Seq.UnicodeFromString("amua");
      if ((_this).fm32(_dafny.Seq.of(_631_v1), true, (_630_v0).f14, (new BigNumber((_632_v2).length)).isLessThanOrEqualTo((_630_v0).f14), globalState)) {
        let _633_v3;
        _633_v3 = new BigNumber(-903);
        _633_v3 = (_630_v0).f14;
        let _634_v4;
        _634_v4 = new _dafny.CodePoint('d'.codePointAt(0));
        let _635_v5;
        _635_v5 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(369),_634_v4);
        let _636_v6;
        _636_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(((_635_v5).contains(_633_v3)) ? ((_635_v5).get(_633_v3)) : (_module.__default.fm35(globalState))));
        let _637_v7;
        _637_v7 = _dafny.Seq.of(_636_v6);
        _637_v7 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(369)), ((_638_v6) => function (_639_i0) {
          return _638_v6;
        })(_636_v6));
        let _640_v8;
        let _nw99 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
        _640_v8 = _nw99;
        _640_v8 = _640_v8;
        let _641_v9;
        let _nw100 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _641_v9 = _nw100;
        let _index87 = _module.__default.safeIndex(new BigNumber(496), new BigNumber((_641_v9).length));
        (_641_v9)[_index87] = _632_v2;
        let _index88 = _module.__default.safeIndex(new BigNumber(496), new BigNumber((_641_v9).length));
        (_641_v9)[_index88] = _632_v2;
        let _642_v10;
        _642_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_630_v0).f14);
        let _643_v11;
        _643_v11 = _dafny.MultiSet.fromElements(new BigNumber(992));
        let _644_v12;
        _644_v12 = _dafny.Map.Empty.slice().updateUnsafe(_642_v10,(_module.D9.create_DC25(_633_v3, (_630_v0).f14, (_643_v11).update(new BigNumber((_643_v11).cardinality()), _module.__default.abs(_633_v3)), _635_v5, (_630_v0).f14)).dtor_cf45);
        _633_v3 = (new BigNumber(((_module.__default.fm36((_this).f15, (_this).f15, p0, globalState)).Merge(_644_v12)).length)).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus(_633_v3), _633_v3));
      } else {
        (globalState).f2 = ((_this).f15) || (((_630_v0).f14).isLessThan((_630_v0).f14));
        let _645_v13;
        let _nw101 = new _module.C2();
        _nw101.__ctor(((_dafny.ZERO).minus((_630_v0).f14)).multipliedBy((_630_v0).f14), !(((_dafny.ZERO).minus((_630_v0).f14)).isEqualTo((_630_v0).f14)));
        _645_v13 = _nw101;
        let _646_v14;
        _646_v14 = new BigNumber(678);
        _646_v14 = (_630_v0).f14;
        let _647_v15;
        _647_v15 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_this).f9);
        let _648_v16;
        _648_v16 = _dafny.Map.Empty.slice().updateUnsafe((_630_v0).f14,!((_this).f15) || ((((_647_v15).contains(!(true))) ? ((_647_v15).get(!(true))) : ((_this).f15))));
        let _649_v17;
        _649_v17 = _dafny.Map.Empty.slice().updateUnsafe((_630_v0).f14,new BigNumber((_648_v16).length));
        let _650_v18;
        _650_v18 = _dafny.Map.Empty.slice().updateUnsafe(p0,_649_v17);
        if ((((_648_v16).contains(_646_v14)) ? ((_648_v16).get(_646_v14)) : (!(_650_v18).contains((_this).f9)))) {
          (globalState).f2 = false;
          (globalState).f2 = ((_this).f9) === (!((_this).f9) || (p0));
          let _651_v19;
          _651_v19 = _dafny.MultiSet.fromElements((_this).f9);
          let _652_v20;
          let _nw102 = Array((new BigNumber(13)).toNumber());
          _nw102[(_dafny.ZERO).toNumber()] = _646_v14;
          _nw102[(_dafny.ONE).toNumber()] = new BigNumber(-248);
          _nw102[(new BigNumber(2)).toNumber()] = new BigNumber(-829);
          _nw102[(new BigNumber(3)).toNumber()] = new BigNumber((_dafny.Seq.of(new BigNumber((_651_v19).cardinality()))).length);
          _nw102[(new BigNumber(4)).toNumber()] = _646_v14;
          _nw102[(new BigNumber(5)).toNumber()] = (_645_v13).f14;
          _nw102[(new BigNumber(6)).toNumber()] = _646_v14;
          _nw102[(new BigNumber(7)).toNumber()] = (_645_v13).f14;
          _nw102[(new BigNumber(8)).toNumber()] = (_645_v13).f14;
          _nw102[(new BigNumber(9)).toNumber()] = (_630_v0).f14;
          _nw102[(new BigNumber(10)).toNumber()] = (_630_v0).f14;
          _nw102[(new BigNumber(11)).toNumber()] = _646_v14;
          _nw102[(new BigNumber(12)).toNumber()] = (_645_v13).f14;
          _652_v20 = _nw102;
          let _653_v21;
          _653_v21 = _dafny.MultiSet.fromElements(_652_v20, _652_v20, _652_v20);
          (_630_v0).m8((_this).f9, (_653_v21).Union(_653_v21), globalState);
          let _654_v22;
          let _nw103 = Array((new BigNumber(29)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _654_v22 = _nw103;
          let _655_v23;
          let _nw104 = Array((new BigNumber(14)).toNumber());
          _nw104[(_dafny.ZERO).toNumber()] = _654_v22;
          _nw104[(_dafny.ONE).toNumber()] = _654_v22;
          _nw104[(new BigNumber(2)).toNumber()] = (((_this).f9) ? (_654_v22) : (_654_v22));
          _nw104[(new BigNumber(3)).toNumber()] = _654_v22;
          _nw104[(new BigNumber(4)).toNumber()] = _654_v22;
          _nw104[(new BigNumber(5)).toNumber()] = _654_v22;
          _nw104[(new BigNumber(6)).toNumber()] = _654_v22;
          _nw104[(new BigNumber(7)).toNumber()] = _654_v22;
          _nw104[(new BigNumber(8)).toNumber()] = _654_v22;
          _nw104[(new BigNumber(9)).toNumber()] = _654_v22;
          _nw104[(new BigNumber(10)).toNumber()] = _654_v22;
          _nw104[(new BigNumber(11)).toNumber()] = _654_v22;
          _nw104[(new BigNumber(12)).toNumber()] = _654_v22;
          _nw104[(new BigNumber(13)).toNumber()] = _654_v22;
          _655_v23 = _nw104;
          let _index89 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_655_v23).length));
          (_655_v23)[_index89] = _654_v22;
          let _656_v24;
          _656_v24 = _module.D13.create_DC34(_654_v22);
          let _657_v25;
          _657_v25 = _dafny.Seq.of(_654_v22, _654_v22, (_656_v24).dtor_cf57, _654_v22, _654_v22);
          let _index90 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_655_v23).length));
          (_655_v23)[_index90] = (_657_v25)[_module.__default.safeIndex((_630_v0).f14, new BigNumber((_657_v25).length))];
          _646_v14 = ((_645_v13).f14).minus(_646_v14);
        } else {
          let _658_v26;
          _658_v26 = new _dafny.CodePoint('q'.codePointAt(0));
          _658_v26 = _658_v26;
          let _659_v27;
          _659_v27 = _dafny.Set.fromElements(_dafny.Seq.of((_645_v13).f14, (_630_v0).f14), _dafny.Seq.of(_646_v14));
          let _660_v28;
          let _nw105 = new _module.C0();
          _nw105.__ctor(_659_v27, (_645_v13).f14);
          _660_v28 = _nw105;
          let _661_v29;
          _661_v29 = _dafny.MultiSet.fromElements(_660_v28);
          let _662_v30;
          _662_v30 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,_module.__default.fm37(p0, (_this).f9, globalState));
          (globalState).f2 = !(_662_v30).contains((_661_v29).IsDisjointFrom(_dafny.MultiSet.fromElements(_660_v28)));
          let _663_v31;
          let _nw106 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Seq.of());
          _663_v31 = _nw106;
          let _664_v32;
          _664_v32 = _dafny.Seq.of(_660_v28, _660_v28, _660_v28);
          let _index91 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_663_v31).length));
          (_663_v31)[_index91] = _664_v32;
          let _index92 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_663_v31).length));
          (_663_v31)[_index92] = _664_v32;
          let _665_v33;
          let _nw107 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
          _665_v33 = _nw107;
          let _index93 = _module.__default.safeIndex(new BigNumber(861), new BigNumber((_665_v33).length));
          (_665_v33)[_index93] = (_630_v0).f14;
          let _666_v34;
          _666_v34 = _dafny.Seq.of((_this).f15, p0);
          let _667_v35;
          _667_v35 = _module.D1.create_DC3((_this).f15, _632_v2, (_645_v13).f14);
          let _668_v36;
          _668_v36 = _module.D4.create_DC14(!((_this).f9), (_this).f15, (_630_v0).f14);
          let _index94 = _module.__default.safeIndex(new BigNumber(861), new BigNumber((_665_v33).length));
          let _rhs127 = (new BigNumber((_666_v34).length)).minus((_645_v13).f14);
          let _rhs128 = _module.__default.fm35(globalState);
          let _rhs129 = _module.__default.safeDivisionInt((_630_v0).f14, new BigNumber(928));
          let _rhs130 = !((((_630_v0).f14).plus((_668_v36).dtor_cf28)).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("oyablh"), (_667_v35).dtor_cf7)).length)));
          let _rhs131 = _658_v26;
          let _lhs76 = _660_v28;
          let _lhs77 = _665_v33;
          let _lhs78 = _module.__default.safeIndex(new BigNumber(861), new BigNumber((_665_v33).length));
          let _lhs79 = globalState;
          _lhs76.f12 = _rhs127;
          _658_v26 = _rhs128;
          _lhs77[_lhs78] = _rhs129;
          _lhs79.f2 = _rhs130;
          _658_v26 = _rhs131;
          (globalState).f2 = !((_this).f9);
        }
        _646_v14 = new BigNumber(335);
      }
      r0 = !((p0) === ((_this).f15));
      let _rhs132 = p0;
      let _rhs133 = !((_this).f9);
      let _rhs134 = (_this).f9;
      let _rhs135 = _632_v2;
      let _lhs80 = globalState;
      let _lhs81 = globalState;
      let _lhs82 = globalState;
      _lhs80.f2 = _rhs132;
      _lhs81.f2 = _rhs133;
      _lhs82.f2 = _rhs134;
      _632_v2 = _rhs135;
      let _hi2 = _module.__default.safeDivisionInt(_module.__default.fm0(globalState), (_630_v0).f14);
      for (let _669_i1 = (_630_v0).f14; _669_i1.isLessThan(_hi2); _669_i1 = _669_i1.plus(_dafny.ONE)) {
        let _670_v37;
        _670_v37 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,!((_this).f9));
        let _671_v38;
        _671_v38 = _dafny.MultiSet.fromElements((_630_v0).f14);
        let _672_v39;
        _672_v39 = _dafny.Map.Empty.slice().updateUnsafe((_this).f15,(_630_v0).f14);
        _670_v37 = (_670_v37).update((_671_v38).IsDisjointFrom(_671_v38), (_672_v39).equals(_672_v39));
        let _673_v40;
        _673_v40 = new BigNumber(-700);
        _673_v40 = _673_v40;
        let _674_v41;
        _674_v41 = _dafny.MultiSet.fromElements(p0, (_this).f9, (_this).f9);
        let _675_v42;
        _675_v42 = _dafny.Seq.of((_dafny.ZERO).minus((new BigNumber(530)).multipliedBy(_669_i1)), _module.__default.safeDivisionInt(new BigNumber((_670_v37).length), (_630_v0).f14), new BigNumber((_674_v41).cardinality()));
        _673_v40 = new BigNumber((_675_v42).length);
        let _676_v43;
        _676_v43 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(new BigNumber(378), _669_i1),!((_this).f9));
        let _rhs136 = (_this).f9;
        let _rhs137 = (((true) ? (_676_v43) : (_676_v43))).update(_669_i1, (_669_i1).isLessThan(_673_v40));
        let _lhs83 = globalState;
        _lhs83.f2 = _rhs136;
        _676_v43 = _rhs137;
      }
      r0 = (_this).f15;
      return r0;
    }
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _677_v0;
      _677_v0 = new BigNumber(605);
      let _678_v1;
      _678_v1 = _dafny.Map.Empty.slice().updateUnsafe(p2,_677_v0);
      let _679_v2;
      _679_v2 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(p2,p0), _678_v1, _678_v1, _678_v1);
      let _680_v3;
      _680_v3 = _module.D14.create_DC37(_679_v2);
      let _681_v4;
      _681_v4 = _dafny.Set.fromElements(_677_v0);
      let _pat_let_tv12 = _679_v2;
      let _pat_let_tv13 = _677_v0;
      let _pat_let_tv14 = _679_v2;
      let _pat_let_tv15 = _678_v1;
      let _pat_let_tv16 = _681_v4;
      let _pat_let_tv17 = _678_v1;
      _677_v0 = new BigNumber((_dafny.Seq.Concat(_679_v2, (function (_pat_let17_0) {
        return function (_682_dt__update__tmp_h0) {
          return function (_pat_let18_0) {
            return function (_683_dt__update_hcf66_h0) {
              return _module.D14.create_DC37(_683_dt__update_hcf66_h0);
            }(_pat_let18_0);
          }(_dafny.Seq.of((_pat_let_tv12)[_module.__default.safeIndex(_pat_let_tv13, new BigNumber((_pat_let_tv14).length))], (_pat_let_tv15).update((_this).f9, new BigNumber((_pat_let_tv16).length)), _pat_let_tv17));
        }(_pat_let17_0);
      }(_680_v3)).dtor_cf66)).length);
      let _684_v5;
      _684_v5 = _dafny.Seq.UnicodeFromString("thqlb");
      let _685_v6;
      _685_v6 = _dafny.MultiSet.fromElements(_677_v0);
      let _686_v7;
      _686_v7 = _module.D1.create_DC2(new BigNumber(937), p1, new BigNumber((_685_v6).cardinality()), _677_v0);
      let _687_v8;
      let _nw108 = new _module.C1();
      _nw108.__ctor(_dafny.Seq.of(_module.__default.fm38((_this).f15, new BigNumber((_684_v5).length), globalState), _686_v7, _686_v7));
      _687_v8 = _nw108;
      let _688_v9;
      _688_v9 = _dafny.Seq.of(_687_v8);
      let _689_v10;
      _689_v10 = _dafny.Seq.of(p1);
      let _690_v11;
      _690_v11 = _dafny.Seq.of(p0);
      let _691_v12;
      let _nw109 = Array((new BigNumber(15)).toNumber());
      _nw109[(_dafny.ZERO).toNumber()] = p2;
      _nw109[(_dafny.ONE).toNumber()] = true;
      _nw109[(new BigNumber(2)).toNumber()] = (_this).f15;
      _nw109[(new BigNumber(3)).toNumber()] = !((_this).f15);
      _nw109[(new BigNumber(4)).toNumber()] = (_this).f9;
      _nw109[(new BigNumber(5)).toNumber()] = (_this).f15;
      _nw109[(new BigNumber(6)).toNumber()] = (_this).f15;
      _nw109[(new BigNumber(7)).toNumber()] = p2;
      _nw109[(new BigNumber(8)).toNumber()] = (_this).f9;
      _nw109[(new BigNumber(9)).toNumber()] = (_this).f15;
      _nw109[(new BigNumber(10)).toNumber()] = (_this).f9;
      _nw109[(new BigNumber(11)).toNumber()] = (_this).f9;
      _nw109[(new BigNumber(12)).toNumber()] = (_this).f15;
      _nw109[(new BigNumber(13)).toNumber()] = (_this).f9;
      _nw109[(new BigNumber(14)).toNumber()] = p1;
      _691_v12 = _nw109;
      let _692_v13;
      _692_v13 = _dafny.MultiSet.fromElements(_691_v12, _691_v12);
      let _693_v14;
      _693_v14 = _dafny.Map.Empty.slice().updateUnsafe(_677_v0,p1);
      let _694_v15;
      let _nw110 = Array((new BigNumber(25)).toNumber());
      _nw110[(_dafny.ZERO).toNumber()] = (((_this).f9) ? (_677_v0) : (_677_v0));
      _nw110[(_dafny.ONE).toNumber()] = new BigNumber((_688_v9).length);
      _nw110[(new BigNumber(2)).toNumber()] = new BigNumber((_689_v10).length);
      _nw110[(new BigNumber(3)).toNumber()] = new BigNumber((_690_v11).length);
      _nw110[(new BigNumber(4)).toNumber()] = _677_v0;
      _nw110[(new BigNumber(5)).toNumber()] = (p0).plus(_677_v0);
      _nw110[(new BigNumber(6)).toNumber()] = p0;
      _nw110[(new BigNumber(7)).toNumber()] = p0;
      _nw110[(new BigNumber(8)).toNumber()] = new BigNumber((_692_v13).cardinality());
      _nw110[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_684_v5, _684_v5)).length);
      _nw110[(new BigNumber(10)).toNumber()] = _677_v0;
      _nw110[(new BigNumber(11)).toNumber()] = _677_v0;
      _nw110[(new BigNumber(12)).toNumber()] = _677_v0;
      _nw110[(new BigNumber(13)).toNumber()] = p0;
      _nw110[(new BigNumber(14)).toNumber()] = _677_v0;
      _nw110[(new BigNumber(15)).toNumber()] = _677_v0;
      _nw110[(new BigNumber(16)).toNumber()] = new BigNumber((_dafny.Seq.of(new BigNumber(((_687_v8).fm2(_677_v0, true, p1, p0, globalState)).length))).length);
      _nw110[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus(_677_v0);
      _nw110[(new BigNumber(18)).toNumber()] = ((_dafny.ZERO).minus(new BigNumber((_693_v14).length))).multipliedBy(new BigNumber(-316));
      _nw110[(new BigNumber(19)).toNumber()] = p0;
      _nw110[(new BigNumber(20)).toNumber()] = new BigNumber(44);
      _nw110[(new BigNumber(21)).toNumber()] = new BigNumber(660);
      _nw110[(new BigNumber(22)).toNumber()] = new BigNumber((_module.__default.fm39(globalState)).length);
      _nw110[(new BigNumber(23)).toNumber()] = p0;
      _nw110[(new BigNumber(24)).toNumber()] = _677_v0;
      _694_v15 = _nw110;
      let _index95 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_694_v15).length));
      (_694_v15)[_index95] = p0;
      let _695_v16;
      _695_v16 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("afneutgkf"), _dafny.Seq.UnicodeFromString("eaxwwu"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-209)), function (_696_i0) {
        return new _dafny.CodePoint('p'.codePointAt(0));
      }));
      let _697_v17;
      let _init14 = ((_698_p2) => function (_699_i1) {
        return _dafny.Map.Empty.slice().updateUnsafe(_698_p2,true);
      })(p2);
      let _nw111 = Array((new BigNumber(4)).toNumber());
      for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw111.length); _i0_14++) {
        _nw111[_i0_14] = _init14(new BigNumber(_i0_14));
      }
      _697_v17 = _nw111;
      let _700_v18;
      _700_v18 = _module.D7.create_DC19(_697_v17, p1, _684_v5);
      let _index96 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_694_v15).length));
      let _rhs138 = _module.__default.safeDivisionInt((_690_v11)[_module.__default.safeIndex(_677_v0, new BigNumber((_690_v11).length))], p0);
      let _rhs139 = (_dafny.MultiSet.fromElements(_684_v5, (_700_v18).dtor_cf35, _dafny.Seq.Create(_module.__default.abs(new BigNumber(651)), function (_701_i2) {
        return new _dafny.CodePoint('e'.codePointAt(0));
      }), _684_v5)).Union(_module.__default.fm40(globalState));
      let _rhs140 = (_this).f15;
      let _rhs141 = (_677_v0).minus((_690_v11)[_module.__default.safeIndex(p0, new BigNumber((_690_v11).length))]);
      let _lhs84 = _694_v15;
      let _lhs85 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_694_v15).length));
      let _lhs86 = globalState;
      _lhs84[_lhs85] = _rhs138;
      _695_v16 = _rhs139;
      _lhs86.f2 = _rhs140;
      _677_v0 = _rhs141;
      let _702_v19;
      let _init15 = function (_703_i3) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      };
      let _nw112 = Array((new BigNumber(26)).toNumber());
      for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw112.length); _i0_15++) {
        _nw112[_i0_15] = _init15(new BigNumber(_i0_15));
      }
      _702_v19 = _nw112;
      let _704_v20;
      _704_v20 = new _dafny.CodePoint('m'.codePointAt(0));
      let _index97 = _module.__default.safeIndex(new BigNumber(966), new BigNumber((_702_v19).length));
      (_702_v19)[_index97] = _704_v20;
      let _index98 = _module.__default.safeIndex(new BigNumber(966), new BigNumber((_702_v19).length));
      (_702_v19)[_index98] = _704_v20;
      let _init16 = ((_705_v15) => function (_706_i4) {
        return _module.__default.safeModuloInt(_706_i4, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_705_v15)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((_705_v15).length))],(_705_v15)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((_705_v15).length))])).length));
      })(_694_v15);
      let _nw113 = Array((new BigNumber(6)).toNumber());
      for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw113.length); _i0_16++) {
        _nw113[_i0_16] = _init16(new BigNumber(_i0_16));
      }
      _694_v15 = _nw113;
      let _707_v21;
      _707_v21 = _module.D10.create_DC27(_dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber(((_this).fm2((_dafny.ZERO).minus(new BigNumber((_684_v5).length)), (_this).f9, (_this).f9, _677_v0, globalState)).length)));
      let _708_v22;
      _708_v22 = _module.D14.create_DC38((_702_v19)[_module.__default.safeIndex(new BigNumber(966), new BigNumber((_702_v19).length))], (_this).f15, _module.__default.fm41(_707_v21, _677_v0, _690_v11, globalState));
      let _pat_let_tv18 = p1;
      let _pat_let_tv19 = _689_v10;
      let _709_v23;
      _709_v23 = _dafny.Set.fromElements(_708_v22, _708_v22, _708_v22, _708_v22, function (_pat_let19_0) {
        return function (_710_dt__update__tmp_h1) {
          return function (_pat_let20_0) {
            return function (_711_dt__update_hcf68_h0) {
              return function (_pat_let21_0) {
                return function (_712_dt__update_hcf69_h0) {
                  return _module.D14.create_DC38((_710_dt__update__tmp_h1).dtor_cf67, _711_dt__update_hcf68_h0, _712_dt__update_hcf69_h0);
                }(_pat_let21_0);
              }(_pat_let_tv19);
            }(_pat_let20_0);
          }(_pat_let_tv18);
        }(_pat_let19_0);
      }(_708_v22));
      _709_v23 = (_709_v23).Intersect((_709_v23).Difference(_709_v23));
      let _713_v24;
      _713_v24 = _dafny.Seq.of(_685_v6, _dafny.MultiSet.FromArray(_690_v11), _685_v6);
      let _714_v25;
      _714_v25 = (_713_v24)[_module.__default.safeIndex((_694_v15)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((_694_v15).length))], new BigNumber((_713_v24).length))];
      if (function (_source7) {
        let _715___mcc_h0 = _source7;
        let _716_cf31 = _715___mcc_h0;
        return false;
      }(_714_v25)) {
        (globalState).f2 = (_677_v0).isLessThanOrEqualTo((_694_v15)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((_694_v15).length))]);
        if (_dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-978)), ((_717_v22) => function (_718_i5) {
          return (_717_v22).dtor_cf67;
        })(_708_v22)), _684_v5)) {
          let _719_v26;
          _719_v26 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),p1);
          let _720_v27;
          _720_v27 = _dafny.Set.fromElements(!(p1));
          let _721_v28;
          _721_v28 = _module.D15.create_DC40(_dafny.Set.fromElements((_this).f9));
          let _722_v29;
          _722_v29 = _dafny.Seq.of(_720_v27, _720_v27, (_721_v28).dtor_cf72, _720_v27);
          let _index99 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_694_v15).length));
          (_694_v15)[_index99] = _module.__default.safeDivisionInt(new BigNumber((_719_v26).length), new BigNumber(((_722_v29)[_module.__default.safeIndex(_677_v0, new BigNumber((_722_v29).length))]).length));
          let _723_v30;
          _723_v30 = _dafny.Map.Empty.slice().updateUnsafe(p1,_704_v20);
          (globalState).f2 = !((p2) || (_dafny.Seq.contains((_this).fm33(globalState), (((_723_v30).contains(p1)) ? ((_723_v30).get(p1)) : ((_702_v19)[_module.__default.safeIndex(new BigNumber(966), new BigNumber((_702_v19).length))])))));
          _677_v0 = (_694_v15)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((_694_v15).length))];
          let _724_v31;
          _724_v31 = _dafny.Map.Empty.slice().updateUnsafe(_684_v5,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(511)), ((_725_p0) => function (_726_i6) {
            return _725_p0;
          })(p0))).length));
          let _index100 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_694_v15).length));
          (_694_v15)[_index100] = (_dafny.ZERO).minus((((_724_v31).contains(_dafny.Seq.Concat(_684_v5, (_this).fm33(globalState)))) ? ((_724_v31).get(_dafny.Seq.Concat(_684_v5, (_this).fm33(globalState)))) : (new BigNumber((_module.__default.fm42(p2, p1, new BigNumber((_689_v10).length), globalState)).cardinality()))));
          let _nw114 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
          _694_v15 = _nw114;
        } else {
          let _727_v32;
          _727_v32 = _dafny.Map.Empty.slice().updateUnsafe(!(p1),p1);
          let _index101 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_694_v15).length));
          (_694_v15)[_index101] = new BigNumber((_727_v32).length);
          let _index102 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_694_v15).length));
          (_694_v15)[_index102] = p0;
          let _728_v33;
          _728_v33 = _dafny.Set.fromElements(((_694_v15)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((_694_v15).length))]).isLessThan(_677_v0));
          let _index103 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_694_v15).length));
          let _rhs142 = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(p1, p1)).length), _677_v0);
          let _rhs143 = _728_v33;
          let _rhs144 = _694_v15;
          let _rhs145 = (_this).f9;
          let _rhs146 = (p0).isLessThanOrEqualTo(_677_v0);
          let _lhs87 = _694_v15;
          let _lhs88 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_694_v15).length));
          let _lhs89 = globalState;
          let _lhs90 = globalState;
          _lhs87[_lhs88] = _rhs142;
          _728_v33 = _rhs143;
          _694_v15 = _rhs144;
          _lhs89.f2 = _rhs145;
          _lhs90.f2 = _rhs146;
          let _729_v34;
          let _nw115 = Array((new BigNumber(29)).toNumber()).fill(_module.D3.Default());
          _729_v34 = _nw115;
          let _730_v35;
          _730_v35 = _module.D3.create_DC8();
          let _index104 = _module.__default.safeIndex(new BigNumber(147), new BigNumber((_729_v34).length));
          (_729_v34)[_index104] = _730_v35;
          let _index105 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_694_v15).length));
          let _index106 = _module.__default.safeIndex(new BigNumber(147), new BigNumber((_729_v34).length));
          let _rhs147 = (p0).isLessThanOrEqualTo(_module.__default.safeModuloInt((_694_v15)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((_694_v15).length))], p0));
          let _rhs148 = _727_v32;
          let _rhs149 = ((p1) ? (p0) : ((_694_v15)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((_694_v15).length))]));
          let _rhs150 = ((p1) ? (_730_v35) : (_730_v35));
          let _lhs91 = globalState;
          let _lhs92 = _694_v15;
          let _lhs93 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_694_v15).length));
          let _lhs94 = _729_v34;
          let _lhs95 = _module.__default.safeIndex(new BigNumber(147), new BigNumber((_729_v34).length));
          _lhs91.f2 = _rhs147;
          _727_v32 = _rhs148;
          _lhs92[_lhs93] = _rhs149;
          _lhs94[_lhs95] = _rhs150;
          let _index107 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_694_v15).length));
          (_694_v15)[_index107] = (_dafny.ZERO).minus(_module.__default.fm0(globalState));
        }
        let _731_v36;
        _731_v36 = _module.D4.create_DC13(_677_v0);
        _677_v0 = (_dafny.ZERO).minus((_731_v36).dtor_cf25);
        let _732_v37;
        let _nw116 = Array((new BigNumber(8)).toNumber());
        _nw116[(_dafny.ZERO).toNumber()] = _684_v5;
        _nw116[(_dafny.ONE).toNumber()] = _dafny.Seq.UnicodeFromString("rmco");
        _nw116[(new BigNumber(2)).toNumber()] = (_this).fm33(globalState);
        _nw116[(new BigNumber(3)).toNumber()] = _684_v5;
        _nw116[(new BigNumber(4)).toNumber()] = _684_v5;
        _nw116[(new BigNumber(5)).toNumber()] = _684_v5;
        _nw116[(new BigNumber(6)).toNumber()] = _684_v5;
        _nw116[(new BigNumber(7)).toNumber()] = _684_v5;
        _732_v37 = _nw116;
        let _733_v38;
        _733_v38 = _dafny.Seq.of(_732_v37, _732_v37, _732_v37);
        let _734_v39;
        _734_v39 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_687_v8,new BigNumber(590)),(_733_v38)[_module.__default.safeIndex((_694_v15)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((_694_v15).length))], new BigNumber((_733_v38).length))]);
        let _735_v40;
        _735_v40 = _dafny.Map.Empty.slice().updateUnsafe(_687_v8,p0);
        _734_v39 = (_734_v39).update(_735_v40, _732_v37);
        let _index108 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_694_v15).length));
        (_694_v15)[_index108] = _677_v0;
      } else {
        let _index109 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_694_v15).length));
        (_694_v15)[_index109] = p0;
        let _736_v41;
        _736_v41 = _dafny.Map.Empty.slice().updateUnsafe((_694_v15)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((_694_v15).length))],(_694_v15)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((_694_v15).length))]);
        let _737_v42;
        _737_v42 = _dafny.Map.Empty.slice().updateUnsafe(((_module.D16.create_DC45(_736_v41)).dtor_cf78).Merge(_736_v41),((_this).f9) || ((_this).f15));
        (globalState).f2 = (((_737_v42).contains(_736_v41)) ? ((_737_v42).get(_736_v41)) : (p1));
        _677_v0 = (_694_v15)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((_694_v15).length))];
        _677_v0 = _module.__default.fm0(globalState);
        _690_v11 = _690_v11;
      }
      return;
    }
    get f15() {
      let _this = this;
      return _this._f15;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f9 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f9() {
      let _this = this;
      return _this._f9;
    };
    __ctor(f9) {
      let _this = this;
      (_this)._f9 = f9;
      return;
    }
    fm2(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(222))).length), new BigNumber((_dafny.MultiSet.fromElements(!((_this).f9))).cardinality())), _dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('v'.codePointAt(0)))).cardinality()))), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(413)), function (_738_i0) {
        return new BigNumber(725);
      }), _dafny.Seq.of(new BigNumber(381))));
    };
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _739_v0;
      _739_v0 = _dafny.MultiSet.fromElements(p0);
      let _740_i0;
      _740_i0 = _dafny.ZERO;
      L7: {
        while ((new BigNumber(((_739_v0).Intersect(_739_v0)).cardinality())).isLessThanOrEqualTo(p0)) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_740_i0)) {
              break L7;
            }
            _740_i0 = (_740_i0).plus(_dafny.ONE);
            let _741_v1;
            let _init17 = function (_742_i1) {
              return new _dafny.CodePoint('x'.codePointAt(0));
            };
            let _nw117 = Array((new BigNumber(14)).toNumber());
            for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw117.length); _i0_17++) {
              _nw117[_i0_17] = _init17(new BigNumber(_i0_17));
            }
            _741_v1 = _nw117;
            _741_v1 = _741_v1;
            let _743_v2;
            let _init18 = ((_744_p2) => function (_745_i2) {
              return _744_p2;
            })(p2);
            let _nw118 = Array((new BigNumber(26)).toNumber());
            for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw118.length); _i0_18++) {
              _nw118[_i0_18] = _init18(new BigNumber(_i0_18));
            }
            _743_v2 = _nw118;
            let _746_v3;
            _746_v3 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
            let _747_v4;
            _747_v4 = _dafny.Seq.of(new BigNumber((_746_v3).length));
            let _748_v5;
            _748_v5 = _dafny.Seq.of((p0).plus(p0), p0, (p0).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,_743_v2)).length)), (_747_v4)[_module.__default.safeIndex(p0, new BigNumber((_747_v4).length))], new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f9)).length));
            _747_v4 = _dafny.Seq.Concat(_747_v4, _748_v5);
            let _749_v6;
            _749_v6 = new BigNumber(442);
            _749_v6 = (new BigNumber((_747_v4).length)).multipliedBy(p0);
            let _750_v7;
            let _init19 = function (_751_i3) {
              return _dafny.Seq.UnicodeFromString("kwxxvgrr");
            };
            let _nw119 = Array((new BigNumber(16)).toNumber());
            for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw119.length); _i0_19++) {
              _nw119[_i0_19] = _init19(new BigNumber(_i0_19));
            }
            _750_v7 = _nw119;
            let _752_v8;
            _752_v8 = _module.D1.create_DC3(p2, _dafny.Seq.Create(_module.__default.abs(new BigNumber(388)), function (_753_i4) {
  return new _dafny.CodePoint('e'.codePointAt(0));
}), _749_v6);
            let _index110 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_750_v7).length));
            (_750_v7)[_index110] = (_752_v8).dtor_cf7;
            let _index111 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_750_v7).length));
            (_750_v7)[_index111] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(997)), function (_754_i5) {
              return new _dafny.CodePoint('d'.codePointAt(0));
            });
          }
        }
      }
      (globalState).f2 = !(!(true));
      let _755_v9;
      _755_v9 = _dafny.Seq.of(p0);
      let _756_v10;
      _756_v10 = _dafny.MultiSet.fromElements(_755_v9);
      let _757_v11;
      _757_v11 = _module.D1.create_DC2(p0, (_this).f9, p0, new BigNumber((_756_v10).cardinality()));
      let _758_v12;
      _758_v12 = _dafny.Seq.of(_757_v11);
      let _759_v13;
      let _nw120 = new _module.C1();
      _nw120.__ctor(_dafny.Seq.Concat(_758_v12, _758_v12));
      _759_v13 = _nw120;
      let _760_v14;
      _760_v14 = _dafny.Map.Empty.slice().updateUnsafe(true,p0);
      let _761_v15;
      _761_v15 = new _dafny.CodePoint('o'.codePointAt(0));
      let _762_v16;
      _762_v16 = _dafny.Seq.of(p2, false);
      let _763_v17;
      _763_v17 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('d'.codePointAt(0)),_762_v16);
      let _764_v18;
      let _nw121 = Array((new BigNumber(21)).toNumber());
      _nw121[(_dafny.ZERO).toNumber()] = p2;
      _nw121[(_dafny.ONE).toNumber()] = (_this).f9;
      _nw121[(new BigNumber(2)).toNumber()] = p1;
      _nw121[(new BigNumber(3)).toNumber()] = ((p2) ? (p2) : (p2));
      _nw121[(new BigNumber(4)).toNumber()] = p2;
      _nw121[(new BigNumber(5)).toNumber()] = p1;
      _nw121[(new BigNumber(6)).toNumber()] = !((_this).f9);
      _nw121[(new BigNumber(7)).toNumber()] = !(p0).isEqualTo(new BigNumber(868));
      _nw121[(new BigNumber(8)).toNumber()] = !((_739_v0).update(p0, _module.__default.abs((_dafny.ZERO).minus(new BigNumber((_760_v14).length))))).contains(p0);
      _nw121[(new BigNumber(9)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_761_v15,_dafny.MultiSet.FromArray((((_763_v17).contains(_761_v15)) ? ((_763_v17).get(_761_v15)) : (_762_v16))))).contains(_761_v15);
      _nw121[(new BigNumber(10)).toNumber()] = (p1) || (p1);
      _nw121[(new BigNumber(11)).toNumber()] = false;
      _nw121[(new BigNumber(12)).toNumber()] = (p0).isLessThanOrEqualTo(p0);
      _nw121[(new BigNumber(13)).toNumber()] = (true) === (!(p1));
      _nw121[(new BigNumber(14)).toNumber()] = _module.__default.fm1(_739_v0, p0, globalState);
      _nw121[(new BigNumber(15)).toNumber()] = (_this).f9;
      _nw121[(new BigNumber(16)).toNumber()] = (_this).f9;
      _nw121[(new BigNumber(17)).toNumber()] = (_this).f9;
      _nw121[(new BigNumber(18)).toNumber()] = p2;
      _nw121[(new BigNumber(19)).toNumber()] = (_762_v16)[_module.__default.safeIndex(p0, new BigNumber((_762_v16).length))];
      _nw121[(new BigNumber(20)).toNumber()] = p1;
      _764_v18 = _nw121;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_764_v18).length))) {
        let _765_i6 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_765_i6)) && ((_765_i6).isLessThan(new BigNumber((_764_v18).length))))) {
          (_764_v18)[(_765_i6)] = false;
        }
      }
      let _pat_let_tv20 = p0;
      let _pat_let_tv21 = p0;
      let _pat_let_tv22 = p0;
      let _pat_let_tv23 = p0;
      let _pat_let_tv24 = p0;
      let _source8 = function (_source9) {
        if (_source9.is_DC8) {
          return _module.D1.create_DC1(_pat_let_tv20);
        } else if (_source9.is_DC9) {
          let _766___mcc_h14 = (_source9).cf17;
          let _767___mcc_h15 = (_source9).cf18;
          let _768_cf18 = _767___mcc_h15;
          let _769_cf17 = _766___mcc_h14;
          if (_769_cf17) {
            return _module.D1.create_DC1((_dafny.ZERO).minus(_pat_let_tv21));
          } else {
            return _module.D1.create_DC1(_pat_let_tv22);
          }
        } else if (_source9.is_DC10) {
          return _module.D1.create_DC1(_pat_let_tv23);
        } else {
          let _770___mcc_h16 = (_source9).cf16;
          let _771_cf16 = _770___mcc_h16;
          return _module.D1.create_DC1(_pat_let_tv24);
        }
      }(_module.D3.create_DC8());
      if (_source8.is_DC2) {
        let _772___mcc_h0 = (_source8).cf2;
        let _773___mcc_h1 = (_source8).cf3;
        let _774___mcc_h2 = (_source8).cf4;
        let _775___mcc_h3 = (_source8).cf5;
        let _776_cf5 = _775___mcc_h3;
        let _777_cf4 = _774___mcc_h2;
        let _778_cf3 = _773___mcc_h1;
        let _779_cf2 = _772___mcc_h0;
        let _780_v19;
        _780_v19 = _dafny.Map.Empty.slice().updateUnsafe(p0,false);
        let _781_v20;
        _781_v20 = _module.D1.create_DC4(_776_cf5, p2, (((_780_v19).contains(_777_cf4)) ? ((_780_v19).get(_777_cf4)) : (p2)), true, _778_cf3);
        let _782_v21;
        _782_v21 = _module.D1.create_DC5(_781_v20);
        let _source10 = _782_v21;
        if (_source10.is_DC2) {
          let _783___mcc_h17 = (_source10).cf2;
          let _784___mcc_h18 = (_source10).cf3;
          let _785___mcc_h19 = (_source10).cf4;
          let _786___mcc_h20 = (_source10).cf5;
          let _787_cf5 = _786___mcc_h20;
          let _788_cf4 = _785___mcc_h19;
          let _789_cf3 = _784___mcc_h18;
          let _790_cf2 = _783___mcc_h17;
          let _791_v22;
          let _init20 = ((_792_cf4) => function (_793_i7) {
            return (_793_i7).minus((_dafny.ZERO).minus(_792_cf4));
          })(_788_cf4);
          let _nw122 = Array((new BigNumber(18)).toNumber());
          for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw122.length); _i0_20++) {
            _nw122[_i0_20] = _init20(new BigNumber(_i0_20));
          }
          _791_v22 = _nw122;
          let _index112 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_791_v22).length));
          (_791_v22)[_index112] = new BigNumber((_dafny.Seq.of(_787_cf5)).length);
          let _index113 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_791_v22).length));
          (_791_v22)[_index113] = _790_cf2;
          let _794_v23;
          let _nw123 = new _module.C1();
          _nw123.__ctor(_759_v13.f13);
          _794_v23 = _nw123;
          let _795_v24;
          let _796_v25;
          let _out10;
          let _out11;
          let _outcollector3 = (_759_v13).m6(_794_v23, _764_v18, true, globalState);
          _out10 = _outcollector3[0];
          _out11 = _outcollector3[1];
          _795_v24 = _out10;
          _796_v25 = _out11;
          let _797_v26;
          let _nw124 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _797_v26 = _nw124;
          let _798_v27;
          let _init21 = ((_799_v0) => function (_800_i8) {
            return _dafny.Map.Empty.slice().updateUnsafe(_799_v0,true);
          })(_739_v0);
          let _nw125 = Array((new BigNumber(5)).toNumber());
          for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw125.length); _i0_21++) {
            _nw125[_i0_21] = _init21(new BigNumber(_i0_21));
          }
          _798_v27 = _nw125;
          let _801_v28;
          _801_v28 = _dafny.Map.Empty.slice().updateUnsafe(_739_v0,_796_v25);
          let _index114 = _module.__default.safeIndex(new BigNumber(146), new BigNumber((_798_v27).length));
          (_798_v27)[_index114] = _801_v28;
          let _802_v29;
          _802_v29 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(499)), ((_803_v24, _804_p0) => function (_805_i9) {
            return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_803_v24,_804_p0)).length);
          })(_795_v24, p0)));
          let _806_v31;
          _806_v31 = _module.D4.create_DC12(_795_v24, _796_v25, p0, (_791_v22)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_791_v22).length))], (_this).f9);
          let _807_v32;
          _807_v32 = _module.D7.create_DC18(function () {
  let _coll21 = new _dafny.Map();
  for (const _compr_21 of (_dafny.Seq.of(_739_v0, _739_v0, _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_806_v31).dtor_cf22), _788_cf4))).Elements) {
    let _808_v30 = _compr_21;
    if (_dafny.Seq.contains(_dafny.Seq.of(_739_v0, _739_v0, _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_806_v31).dtor_cf22), _788_cf4)), _808_v30)) {
      _coll21.push([_808_v30,true]);
    }
  }
  return _coll21;
}());
          let _pat_let_tv25 = _781_v20;
          let _index115 = _module.__default.safeIndex(new BigNumber(146), new BigNumber((_798_v27).length));
          let _rhs151 = function (_pat_let22_0) {
            return function (_809_dt__update__tmp_h0) {
              return function (_pat_let23_0) {
                return function (_810_dt__update_hcf14_h0) {
                  return _module.D1.create_DC5(_810_dt__update_hcf14_h0);
                }(_pat_let23_0);
              }(_pat_let_tv25);
            }(_pat_let22_0);
          }(_module.D1.create_DC5(_781_v20));
          let _rhs152 = _797_v26;
          let _rhs153 = ((!((_this).f9)) ? (!_dafny.areEqual(_802_v29, _802_v29)) : (!(_795_v24)));
          let _rhs154 = (_807_v32).dtor_cf32;
          let _lhs96 = _798_v27;
          let _lhs97 = _module.__default.safeIndex(new BigNumber(146), new BigNumber((_798_v27).length));
          _782_v21 = _rhs151;
          _797_v26 = _rhs152;
          _789_cf3 = _rhs153;
          _lhs96[_lhs97] = _rhs154;
          _787_cf5 = (_790_cf2).multipliedBy(new BigNumber((_739_v0).cardinality()));
        } else if (_source10.is_DC3) {
          let _811___mcc_h21 = (_source10).cf6;
          let _812___mcc_h22 = (_source10).cf7;
          let _813___mcc_h23 = (_source10).cf8;
          let _814_cf8 = _813___mcc_h23;
          let _815_cf7 = _812___mcc_h22;
          let _816_cf6 = _811___mcc_h21;
          _module.__default.m0(p1, globalState);
          _761_v15 = _761_v15;
          let _817_v33;
          _817_v33 = _dafny.Set.fromElements(_755_v9, _dafny.Seq.Create(_module.__default.abs(new BigNumber(881)), function (_818_i10) {
            return new BigNumber(561);
          }));
          let _819_v34;
          let _nw126 = new _module.C0();
          _nw126.__ctor(_817_v33, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,_778_cf3)).length));
          _819_v34 = _nw126;
          let _820_v35;
          _820_v35 = _dafny.Map.Empty.slice().updateUnsafe(_816_cf6,_819_v34);
          let _821_v36;
          _821_v36 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(p2,_819_v34), (_820_v35).Merge(_820_v35), (_820_v35).Merge(_dafny.Map.Empty.slice().updateUnsafe(_816_cf6,_819_v34)));
          let _822_v37;
          _822_v37 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(720)), ((_823_v15) => function (_824_i11) {
            return _823_v15;
          })(_761_v15)), _dafny.Seq.UnicodeFromString("unmdd"));
          let _825_v38;
          _825_v38 = _dafny.Seq.of(_822_v37);
          let _826_v39;
          _826_v39 = _dafny.Seq.of(_815_cf7, _815_cf7);
          let _827_v40;
          let _nw127 = Array((new BigNumber(21)).toNumber());
          _nw127[(_dafny.ZERO).toNumber()] = (_822_v37).Union(_822_v37);
          _nw127[(_dafny.ONE).toNumber()] = (_822_v37).Union((_825_v38)[_module.__default.safeIndex(_779_cf2, new BigNumber((_825_v38).length))]);
          _nw127[(new BigNumber(2)).toNumber()] = _822_v37;
          _nw127[(new BigNumber(3)).toNumber()] = _822_v37;
          _nw127[(new BigNumber(4)).toNumber()] = (_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(954)), ((_828_cf7) => function (_829_i12) {
            return _828_cf7;
          })(_815_cf7)))).Union(_822_v37);
          _nw127[(new BigNumber(5)).toNumber()] = (_dafny.MultiSet.FromArray(_826_v39)).Difference(_822_v37);
          _nw127[(new BigNumber(6)).toNumber()] = _module.__default.fm19(p0, globalState);
          _nw127[(new BigNumber(7)).toNumber()] = _822_v37;
          _nw127[(new BigNumber(8)).toNumber()] = (_dafny.MultiSet.FromArray(_826_v39)).Intersect(_822_v37);
          _nw127[(new BigNumber(9)).toNumber()] = _822_v37;
          _nw127[(new BigNumber(10)).toNumber()] = _822_v37;
          _nw127[(new BigNumber(11)).toNumber()] = _822_v37;
          _nw127[(new BigNumber(12)).toNumber()] = _822_v37;
          _nw127[(new BigNumber(13)).toNumber()] = _module.__default.fm19(_777_cf4, globalState);
          _nw127[(new BigNumber(14)).toNumber()] = _822_v37;
          _nw127[(new BigNumber(15)).toNumber()] = _822_v37;
          _nw127[(new BigNumber(16)).toNumber()] = (_822_v37).update(_815_cf7, _module.__default.abs(_819_v34.f12));
          _nw127[(new BigNumber(17)).toNumber()] = _822_v37;
          _nw127[(new BigNumber(18)).toNumber()] = _822_v37;
          _nw127[(new BigNumber(19)).toNumber()] = _822_v37;
          _nw127[(new BigNumber(20)).toNumber()] = _822_v37;
          _827_v40 = _nw127;
          let _index116 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_827_v40).length));
          (_827_v40)[_index116] = _822_v37;
          let _830_v41;
          _830_v41 = _dafny.Map.Empty.slice().updateUnsafe(_819_v34.f12,_819_v34);
          let _831_v42;
          let _init22 = ((_832_p1, _833_cf3) => function (_834_i13) {
            return _dafny.Map.Empty.slice().updateUnsafe(_832_p1,_833_cf3);
          })(p1, _778_cf3);
          let _nw128 = Array((new BigNumber(27)).toNumber());
          for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw128.length); _i0_22++) {
            _nw128[_i0_22] = _init22(new BigNumber(_i0_22));
          }
          _831_v42 = _nw128;
          let _835_v43;
          _835_v43 = _module.D7.create_DC19(_831_v42, p1, _815_cf7);
          let _index117 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_827_v40).length));
          let _rhs155 = new BigNumber(-627);
          let _rhs156 = _dafny.Seq.update(_dafny.Seq.Concat(_821_v36, _821_v36), _module.__default.safeIndex(_776_cf5, new BigNumber((_dafny.Seq.Concat(_821_v36, _821_v36)).length)), _820_v35);
          let _rhs157 = _822_v37;
          let _rhs158 = _830_v41;
          let _rhs159 = new BigNumber((_dafny.Seq.Concat((_835_v43).dtor_cf35, _815_cf7)).length);
          let _lhs98 = _827_v40;
          let _lhs99 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_827_v40).length));
          _777_cf4 = _rhs155;
          _821_v36 = _rhs156;
          _lhs98[_lhs99] = _rhs157;
          _830_v41 = _rhs158;
          _777_cf4 = _rhs159;
          let _836_v44;
          _836_v44 = _module.D1.create_DC4(_777_cf4, true, !(!(_778_cf3)), (_this).f9, false);
          let _837_v45;
          let _nw129 = new _module.C0();
          _nw129.__ctor(_817_v33, (_836_v44).dtor_cf9);
          _837_v45 = _nw129;
        } else if (_source10.is_DC4) {
          let _838___mcc_h24 = (_source10).cf9;
          let _839___mcc_h25 = (_source10).cf10;
          let _840___mcc_h26 = (_source10).cf11;
          let _841___mcc_h27 = (_source10).cf12;
          let _842___mcc_h28 = (_source10).cf13;
          let _843_cf13 = _842___mcc_h28;
          let _844_cf12 = _841___mcc_h27;
          let _845_cf11 = _840___mcc_h26;
          let _846_cf10 = _839___mcc_h25;
          let _847_cf9 = _838___mcc_h24;
          let _848_v46;
          _848_v46 = _dafny.Map.Empty.slice().updateUnsafe((_739_v0).Difference(_739_v0),!((p1) && (p2)));
          _848_v46 = (_848_v46).update((_739_v0).Intersect(_dafny.MultiSet.fromElements(_776_cf5)), !((_762_v16)[_module.__default.safeIndex(_847_cf9, new BigNumber((_762_v16).length))]));
          let _index118 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_764_v18).length));
          (_764_v18)[_index118] = (_778_cf3) === (_778_cf3);
          let _index119 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_764_v18).length));
          (_764_v18)[_index119] = _844_cf12;
          let _index120 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_764_v18).length));
          (_764_v18)[_index120] = (_764_v18)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_764_v18).length))];
          (globalState).f2 = p2;
        } else if (_source10.is_DC1) {
          let _849___mcc_h29 = (_source10).cf1;
          let _850_cf1 = _849___mcc_h29;
          let _851_v47;
          _851_v47 = _dafny.Set.fromElements(_755_v9);
          let _852_v48;
          let _nw130 = new _module.C0();
          _nw130.__ctor((_851_v47).Intersect(_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(534)), ((_853_cf5) => function (_854_i14) {
            return (_dafny.ZERO).minus(_853_cf5);
          })(_776_cf5)))), new BigNumber(197));
          _852_v48 = _nw130;
          _852_v48 = _852_v48;
          let _855_v49;
          _855_v49 = _module.D1.create_DC3((_778_cf3) === (p2), _dafny.Seq.UnicodeFromString("srewo"), _777_cf4);
          _855_v49 = _module.D1.create_DC3(true, _dafny.Seq.UnicodeFromString("lrcaagrw"), _850_cf1);
          let _856_v50;
          let _init23 = ((_857_p0) => function (_858_i15) {
            return (_858_i15).multipliedBy(_857_p0);
          })(p0);
          let _nw131 = Array((new BigNumber(5)).toNumber());
          for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw131.length); _i0_23++) {
            _nw131[_i0_23] = _init23(new BigNumber(_i0_23));
          }
          _856_v50 = _nw131;
          _856_v50 = _856_v50;
        } else {
          let _859___mcc_h30 = (_source10).cf14;
          let _860_cf14 = _859___mcc_h30;
          let _861_v51;
          let _nw132 = new _module.C1();
          _nw132.__ctor(_758_v12);
          _861_v51 = _nw132;
          let _862_v52;
          let _863_v53;
          let _out12;
          let _out13;
          let _outcollector4 = (_759_v13).m6(_861_v51, _764_v18, p1, globalState);
          _out12 = _outcollector4[0];
          _out13 = _outcollector4[1];
          _862_v52 = _out12;
          _863_v53 = _out13;
          let _864_v54;
          _864_v54 = _dafny.Seq.UnicodeFromString("wslaqc");
          _864_v54 = _dafny.Seq.update(_864_v54, _module.__default.safeIndex(_779_cf2, new BigNumber((_864_v54).length)), _761_v15);
          _761_v15 = new _dafny.CodePoint('r'.codePointAt(0));
          _864_v54 = _864_v54;
        }
        let _rhs160 = !(true);
        let _rhs161 = (_dafny.ZERO).minus((_module.__default.safeModuloInt((_dafny.ZERO).minus(_777_cf4), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_755_v9).length),p0)).length))).multipliedBy(_776_cf5));
        let _lhs100 = globalState;
        _lhs100.f2 = _rhs160;
        _777_cf4 = _rhs161;
        _778_cf3 = p1;
        _777_cf4 = (p0).plus(_779_cf2);
      } else if (_source8.is_DC3) {
        let _865___mcc_h4 = (_source8).cf6;
        let _866___mcc_h5 = (_source8).cf7;
        let _867___mcc_h6 = (_source8).cf8;
        let _868_cf8 = _867___mcc_h6;
        let _869_cf7 = _866___mcc_h5;
        let _870_cf6 = _865___mcc_h4;
        let _871_v55;
        let _nw133 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _871_v55 = _nw133;
        let _index121 = _module.__default.safeIndex(new BigNumber(87), new BigNumber((_871_v55).length));
        (_871_v55)[_index121] = _869_cf7;
        let _872_v56;
        let _init24 = ((_873_p0) => function (_874_i16) {
          return _module.__default.safeDivisionInt(_874_i16, _873_p0);
        })(p0);
        let _nw134 = Array((new BigNumber(29)).toNumber());
        for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw134.length); _i0_24++) {
          _nw134[_i0_24] = _init24(new BigNumber(_i0_24));
        }
        _872_v56 = _nw134;
        let _875_v57;
        _875_v57 = _dafny.MultiSet.fromElements(p2);
        let _876_v58;
        _876_v58 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_dafny.Seq.of(p0, new BigNumber((_875_v57).cardinality())));
        let _index122 = _module.__default.safeIndex(new BigNumber(520), new BigNumber((_872_v56).length));
        (_872_v56)[_index122] = _module.__default.safeModuloInt(new BigNumber((_760_v14).length), new BigNumber((_876_v58).length));
        let _index123 = _module.__default.safeIndex(new BigNumber(87), new BigNumber((_871_v55).length));
        let _index124 = _module.__default.safeIndex(new BigNumber(520), new BigNumber((_872_v56).length));
        let _rhs162 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(491)), ((_877_v15) => function (_878_i17) {
          return _877_v15;
        })(_761_v15)), _869_cf7);
        let _rhs163 = p0;
        let _rhs164 = _module.__default.fm20(!(false), p1, _module.__default.fm0(globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(51)), ((_879_v15) => function (_880_i18) {
          return _879_v15;
        })(_761_v15)), globalState);
        let _rhs165 = new BigNumber(812);
        let _lhs101 = _871_v55;
        let _lhs102 = _module.__default.safeIndex(new BigNumber(87), new BigNumber((_871_v55).length));
        let _lhs103 = _872_v56;
        let _lhs104 = _module.__default.safeIndex(new BigNumber(520), new BigNumber((_872_v56).length));
        _lhs101[_lhs102] = _rhs162;
        _lhs103[_lhs104] = _rhs163;
        _762_v16 = _rhs164;
        _868_cf8 = _rhs165;
        let _881_v59;
        _881_v59 = _dafny.Set.fromElements(_dafny.Seq.of(_870_cf6), _762_v16);
        if (!(_881_v59).contains(_762_v16)) {
          let _index125 = _module.__default.safeIndex(new BigNumber(520), new BigNumber((_872_v56).length));
          let _rhs166 = _868_cf8;
          let _rhs167 = _872_v56;
          let _lhs105 = _872_v56;
          let _lhs106 = _module.__default.safeIndex(new BigNumber(520), new BigNumber((_872_v56).length));
          _lhs105[_lhs106] = _rhs166;
          _872_v56 = _rhs167;
          let _index126 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((_764_v18).length));
          (_764_v18)[_index126] = false;
          let _index127 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((_764_v18).length));
          (_764_v18)[_index127] = p2;
          _868_cf8 = (_872_v56)[_module.__default.safeIndex(new BigNumber(520), new BigNumber((_872_v56).length))];
          let _882_v60;
          _882_v60 = _module.D1.create_DC4(p0, false, p1, !(_870_cf6), (_764_v18)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((_764_v18).length))]);
          _882_v60 = _882_v60;
          _870_cf6 = !((((_760_v14).contains((_764_v18)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((_764_v18).length))])) ? ((_760_v14).get((_764_v18)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((_764_v18).length))])) : ((_872_v56)[_module.__default.safeIndex(new BigNumber(520), new BigNumber((_872_v56).length))]))).isEqualTo(_module.__default.safeDivisionInt((_872_v56)[_module.__default.safeIndex(new BigNumber(520), new BigNumber((_872_v56).length))], _868_cf8));
        } else {
          let _index128 = _module.__default.safeIndex(new BigNumber(520), new BigNumber((_872_v56).length));
          (_872_v56)[_index128] = (_module.__default.safeModuloInt(p0, _868_cf8)).multipliedBy((_872_v56)[_module.__default.safeIndex(new BigNumber(520), new BigNumber((_872_v56).length))]);
          _module.__default.m0(_870_cf6, globalState);
          let _883_v61;
          _883_v61 = _dafny.Set.fromElements(_755_v9);
          let _884_v62;
          let _nw135 = new _module.C0();
          _nw135.__ctor(((_870_cf6) ? (_883_v61) : (_dafny.Set.fromElements((_this).fm2(_868_cf8, (_this).f9, _870_cf6, _868_cf8, globalState), _755_v9, _755_v9))), (((_760_v14).contains(p1)) ? ((_760_v14).get(p1)) : (new BigNumber(((_871_v55)[_module.__default.safeIndex(new BigNumber(87), new BigNumber((_871_v55).length))]).length))));
          _884_v62 = _nw135;
          let _885_v63;
          _885_v63 = _dafny.Map.Empty.slice().updateUnsafe(_870_cf6,_872_v56);
          let _886_v64;
          _886_v64 = (((_885_v63).contains(false)) ? ((_885_v63).get(false)) : (_872_v56));
          let _rhs168 = (_886_v64);
          let _rhs169 = (_this).f9;
          let _lhs107 = globalState;
          _872_v56 = _rhs168;
          _lhs107.f2 = _rhs169;
          let _887_v65;
          _887_v65 = _dafny.Map.Empty.slice().updateUnsafe(_868_cf8,_dafny.Seq.UnicodeFromString("sfasfg"));
          _887_v65 = (_887_v65).update(new BigNumber(568), _dafny.Seq.Concat(_869_cf7, _869_cf7));
        }
        _868_cf8 = (_872_v56)[_module.__default.safeIndex(new BigNumber(520), new BigNumber((_872_v56).length))];
        let _index129 = _module.__default.safeIndex(new BigNumber(520), new BigNumber((_872_v56).length));
        (_872_v56)[_index129] = new BigNumber(((_875_v57).Union((_875_v57).Difference(_module.__default.fm21(new BigNumber((_755_v9).length), _870_cf6, globalState)))).cardinality());
      } else if (_source8.is_DC4) {
        let _888___mcc_h7 = (_source8).cf9;
        let _889___mcc_h8 = (_source8).cf10;
        let _890___mcc_h9 = (_source8).cf11;
        let _891___mcc_h10 = (_source8).cf12;
        let _892___mcc_h11 = (_source8).cf13;
        let _893_cf13 = _892___mcc_h11;
        let _894_cf12 = _891___mcc_h10;
        let _895_cf11 = _890___mcc_h9;
        let _896_cf10 = _889___mcc_h8;
        let _897_cf9 = _888___mcc_h7;
        let _898_v66;
        _898_v66 = _dafny.Map.Empty.slice().updateUnsafe(_764_v18,_897_cf9);
        if (!(new BigNumber((_898_v66).length)).isEqualTo(new BigNumber((_dafny.Seq.Concat(_755_v9, _755_v9)).length))) {
          let _899_v67;
          let _init25 = ((_900_v9) => function (_901_i19) {
            return _900_v9;
          })(_755_v9);
          let _nw136 = Array((new BigNumber(19)).toNumber());
          for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw136.length); _i0_25++) {
            _nw136[_i0_25] = _init25(new BigNumber(_i0_25));
          }
          _899_v67 = _nw136;
          let _902_v68;
          _902_v68 = _755_v9;
          let _index130 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_899_v67).length));
          (_899_v67)[_index130] = _902_v68;
          let _index131 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_899_v67).length));
          (_899_v67)[_index131] = _902_v68;
          _760_v14 = (_760_v14).update(_895_cf11, p0);
          let _903_v69;
          _903_v69 = _dafny.Set.fromElements(_dafny.Seq.of(_897_cf9));
          let _904_v70;
          let _nw137 = new _module.C0();
          _nw137.__ctor(_903_v69, new BigNumber(663));
          _904_v70 = _nw137;
          _module.__default.m0(p1, globalState);
          (_904_v70).f12 = p0;
        } else {
          let _905_v71;
          _905_v71 = _dafny.Seq.UnicodeFromString("j");
          _905_v71 = _dafny.Seq.UnicodeFromString("hick");
          let _906_v72;
          let _nw138 = Array((new BigNumber(27)).toNumber()).fill([]);
          _906_v72 = _nw138;
          let _907_v73;
          let _nw139 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Seq.of());
          _907_v73 = _nw139;
          let _index132 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_906_v72).length));
          (_906_v72)[_index132] = _907_v73;
          let _908_v74;
          _908_v74 = _dafny.Seq.of(_907_v73);
          let _index133 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_906_v72).length));
          let _rhs170 = ((_module.__default.fm1(_739_v0, _897_cf9, globalState)) ? (_907_v73) : ((_908_v74)[_module.__default.safeIndex(new BigNumber(522), new BigNumber((_908_v74).length))]));
          let _rhs171 = _764_v18;
          let _rhs172 = (_this).f9;
          let _lhs108 = _906_v72;
          let _lhs109 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_906_v72).length));
          _lhs108[_lhs109] = _rhs170;
          _764_v18 = _rhs171;
          _893_cf13 = _rhs172;
          let _909_v75;
          _909_v75 = _module.D7.create_DC18((_dafny.Map.Empty.slice().updateUnsafe(_739_v0,p1)).Merge(_module.__default.fm22(globalState)));
          let _910_v76;
          _910_v76 = _dafny.Map.Empty.slice().updateUnsafe(_739_v0,_893_cf13);
          _909_v75 = _module.D7.create_DC18(_910_v76);
          let _911_v77;
          _911_v77 = _dafny.Map.Empty.slice().updateUnsafe(p0,_897_cf9);
          let _912_v78;
          _912_v78 = _dafny.MultiSet.fromElements(_911_v77);
          let _913_v79;
          _913_v79 = _dafny.Map.Empty.slice().updateUnsafe(_905_v71,_897_cf9);
          _912_v78 = ((_912_v78).Difference(_912_v78)).update(_911_v77, _module.__default.abs(((_755_v9)[_module.__default.safeIndex(_897_cf9, new BigNumber((_755_v9).length))]).multipliedBy(new BigNumber((_913_v79).length))));
          _896_cf10 = (_896_cf10) === (p1);
        }
        let _914_v80;
        _914_v80 = _dafny.Set.fromElements(_895_cf11, (_this).f9, _893_cf13, p2, _896_cf10);
        let _915_v81;
        _915_v81 = _dafny.Map.Empty.slice().updateUnsafe(_914_v80,_896_cf10);
        _897_cf9 = (p0).minus((_dafny.ZERO).minus(new BigNumber((_915_v81).length)));
        let _916_v82;
        _916_v82 = _dafny.Seq.UnicodeFromString("l");
        let _917_v83;
        _917_v83 = _dafny.Seq.of(_916_v82, _916_v82);
        _897_cf9 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_916_v82), _917_v83)).length);
        let _918_v84;
        let _init26 = ((_919_cf9) => function (_920_i20) {
          return (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-154)), ((_921_cf9) => function (_922_i21) {
            return _921_cf9;
          })(_919_cf9)));
        })(_897_cf9);
        let _nw140 = Array((new BigNumber(27)).toNumber());
        for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw140.length); _i0_26++) {
          _nw140[_i0_26] = _init26(new BigNumber(_i0_26));
        }
        _918_v84 = _nw140;
        _918_v84 = _918_v84;
      } else if (_source8.is_DC1) {
        let _923___mcc_h12 = (_source8).cf1;
        let _924_cf1 = _923___mcc_h12;
        let _925_v85;
        let _nw141 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
        _925_v85 = _nw141;
        let _926_v86;
        _926_v86 = _dafny.Map.Empty.slice().updateUnsafe(p2,!(p2));
        let _index134 = _module.__default.safeIndex(new BigNumber(431), new BigNumber((_925_v85).length));
        (_925_v85)[_index134] = new BigNumber((_926_v86).length);
        let _index135 = _module.__default.safeIndex(new BigNumber(431), new BigNumber((_925_v85).length));
        (_925_v85)[_index135] = _924_cf1;
        let _index136 = _module.__default.safeIndex(new BigNumber(431), new BigNumber((_925_v85).length));
        (_925_v85)[_index136] = _924_cf1;
        let _rhs173 = (_this).f9;
        let _rhs174 = !(!(_module.__default.fm1(_739_v0, _module.__default.safeDivisionInt(new BigNumber(-397), (_925_v85)[_module.__default.safeIndex(new BigNumber(431), new BigNumber((_925_v85).length))]), globalState)));
        let _lhs110 = globalState;
        let _lhs111 = globalState;
        _lhs110.f2 = _rhs173;
        _lhs111.f2 = _rhs174;
        (globalState).f2 = p2;
      } else {
        let _927___mcc_h13 = (_source8).cf14;
        let _928_cf14 = _927___mcc_h13;
        let _929_v87;
        let _nw142 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
        _929_v87 = _nw142;
        let _index137 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_929_v87).length));
        (_929_v87)[_index137] = p0;
        let _index138 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_929_v87).length));
        (_929_v87)[_index138] = p0;
        let _930_v88;
        _930_v88 = _dafny.Seq.UnicodeFromString("j");
        let _931_v90;
        _931_v90 = _dafny.Set.fromElements(new BigNumber(-638));
        let _index139 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_929_v87).length));
        let _rhs175 = (_dafny.ZERO).minus(new BigNumber((_module.__default.fm23(!((function () {
          let _coll22 = new _dafny.Set();
          for (const _compr_22 of _dafny.IntegerRange(new BigNumber(319), new BigNumber(877))) {
            let _932_v89 = _compr_22;
            if (((new BigNumber(319)).isLessThanOrEqualTo(_932_v89)) && ((_932_v89).isLessThan(new BigNumber(877)))) {
              _coll22.add(_module.__default.safeModuloInt(_932_v89, (_929_v87)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_929_v87).length))]));
            }
          }
          return _coll22;
        }()).IsSubsetOf(_931_v90)), !_dafny.areEqual(_930_v88, _dafny.Seq.UnicodeFromString("tnvrovbj")), p0, globalState)).length));
        let _rhs176 = _module.__default.fm1(_dafny.MultiSet.fromElements((_929_v87)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_929_v87).length))], new BigNumber(844)), _module.__default.safeDivisionInt(p0, p0), globalState);
        let _rhs177 = ((_this).f9) && ((_this).f9);
        let _rhs178 = _930_v88;
        let _rhs179 = _930_v88;
        let _lhs112 = _929_v87;
        let _lhs113 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_929_v87).length));
        let _lhs114 = globalState;
        let _lhs115 = globalState;
        _lhs112[_lhs113] = _rhs175;
        _lhs114.f2 = _rhs176;
        _lhs115.f2 = _rhs177;
        _930_v88 = _rhs178;
        _930_v88 = _rhs179;
        let _933_v91;
        let _nw143 = Array((new BigNumber(12)).toNumber());
        _933_v91 = _nw143;
        let _934_v92;
        _934_v92 = _dafny.Map.Empty.slice().updateUnsafe(_755_v9,_739_v0);
        let _935_v93;
        let _nw144 = new _module.C2();
        _nw144.__ctor(new BigNumber(((((_934_v92).contains(_755_v9)) ? ((_934_v92).get(_755_v9)) : (_739_v0))).cardinality()), p1);
        _935_v93 = _nw144;
        let _index140 = _module.__default.safeIndex(new BigNumber(142), new BigNumber((_933_v91).length));
        (_933_v91)[_index140] = _935_v93;
        let _index141 = _module.__default.safeIndex(new BigNumber(142), new BigNumber((_933_v91).length));
        (_933_v91)[_index141] = _935_v93;
        let _index142 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_929_v87).length));
        (_929_v87)[_index142] = (_929_v87)[_module.__default.safeIndex(new BigNumber(175), new BigNumber((_929_v87).length))];
      }
      _755_v9 = _755_v9;
      return;
    }
    m7(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _module.D4.Default();
      let r2 = _dafny.MultiSet.Empty;
      let r3 = false;
      if (!((_this).f9) || (!((_this).f9))) {
        let _936_v0;
        _936_v0 = _dafny.MultiSet.fromElements(new BigNumber(230), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements((_this).f9)).cardinality()),p0)).length));
        let _937_v1;
        _937_v1 = _dafny.Seq.of(p2, p2);
        let _938_v2;
        _938_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_this).f9);
        let _939_v3;
        _939_v3 = _dafny.MultiSet.fromElements((_this).f9);
        let _940_v4;
        _940_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,p0);
        let _941_v5;
        _941_v5 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
        let _942_v6;
        _942_v6 = _dafny.Set.fromElements(_941_v5);
        let _943_v7;
        let _nw145 = Array((new BigNumber(28)).toNumber());
        _nw145[(_dafny.ZERO).toNumber()] = (new BigNumber(261)).multipliedBy(p0);
        _nw145[(_dafny.ONE).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(295), p1);
        _nw145[(new BigNumber(2)).toNumber()] = p1;
        _nw145[(new BigNumber(3)).toNumber()] = p1;
        _nw145[(new BigNumber(4)).toNumber()] = p0;
        _nw145[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((((_936_v0).contains(p0)) ? ((_936_v0).get(p0)) : (p0)), _module.__default.fm0(globalState)));
        _nw145[(new BigNumber(6)).toNumber()] = p1;
        _nw145[(new BigNumber(7)).toNumber()] = p0;
        _nw145[(new BigNumber(8)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.update(_937_v1, _module.__default.safeIndex(p0, new BigNumber((_937_v1).length)), p2)).length), (_dafny.ZERO).minus(p0));
        _nw145[(new BigNumber(9)).toNumber()] = _module.__default.safeModuloInt(p0, p0);
        _nw145[(new BigNumber(10)).toNumber()] = p0;
        _nw145[(new BigNumber(11)).toNumber()] = p0;
        _nw145[(new BigNumber(12)).toNumber()] = p1;
        _nw145[(new BigNumber(13)).toNumber()] = p0;
        _nw145[(new BigNumber(14)).toNumber()] = new BigNumber(898);
        _nw145[(new BigNumber(15)).toNumber()] = p1;
        _nw145[(new BigNumber(16)).toNumber()] = p1;
        _nw145[(new BigNumber(17)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_938_v2).length), new BigNumber((_939_v3).cardinality()));
        _nw145[(new BigNumber(18)).toNumber()] = p0;
        _nw145[(new BigNumber(19)).toNumber()] = (_dafny.ZERO).minus(p1);
        _nw145[(new BigNumber(20)).toNumber()] = p1;
        _nw145[(new BigNumber(21)).toNumber()] = new BigNumber(((_940_v4).Merge(_module.__default.fm23((_this).f9, p2, p0, globalState))).length);
        _nw145[(new BigNumber(22)).toNumber()] = new BigNumber(899);
        _nw145[(new BigNumber(23)).toNumber()] = (p1).minus(new BigNumber(944));
        _nw145[(new BigNumber(24)).toNumber()] = (((_this).f9) ? (new BigNumber(229)) : (p0));
        _nw145[(new BigNumber(25)).toNumber()] = p0;
        _nw145[(new BigNumber(26)).toNumber()] = p1;
        _nw145[(new BigNumber(27)).toNumber()] = (p1).minus(new BigNumber((_942_v6).length));
        _943_v7 = _nw145;
        let _index143 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_943_v7).length));
        (_943_v7)[_index143] = p1;
        let _944_v8;
        _944_v8 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(p0, p1, p0),p2);
        let _945_v9;
        _945_v9 = _module.D7.create_DC18((_944_v8).update(_936_v0, p2));
        let _946_v10;
        let _init27 = function (_947_i0) {
          return _module.D7.create_DC20((_this).f9, new _dafny.CodePoint('q'.codePointAt(0)));
        };
        let _nw146 = Array((new BigNumber(9)).toNumber());
        for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw146.length); _i0_27++) {
          _nw146[_i0_27] = _init27(new BigNumber(_i0_27));
        }
        _946_v10 = _nw146;
        let _948_v11;
        _948_v11 = _dafny.Seq.of((_941_v5).Merge(_941_v5));
        let _949_v12;
        _949_v12 = _dafny.Seq.of(p1);
        let _950_v13;
        _950_v13 = _dafny.Seq.UnicodeFromString("kcofbj");
        let _index144 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_943_v7).length));
        let _rhs180 = new BigNumber((_948_v11).length);
        let _rhs181 = _dafny.Seq.update(_module.__default.fm20(!(p1).isEqualTo(new BigNumber((_937_v1).length)), _module.__default.fm1(_module.__default.fm29(p2, (_this).f9, globalState), (_949_v12)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_949_v12).length))], globalState), (p0).plus(new BigNumber((_937_v1).length)), _950_v13, globalState), _module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((_module.__default.fm20(!(p1).isEqualTo(new BigNumber((_937_v1).length)), _module.__default.fm1(_module.__default.fm29(p2, (_this).f9, globalState), (_949_v12)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_949_v12).length))], globalState), (p0).plus(new BigNumber((_937_v1).length)), _950_v13, globalState)).length)), p2);
        let _rhs182 = _module.D7.create_DC18(_944_v8);
        let _rhs183 = p0;
        let _rhs184 = _946_v10;
        let _lhs116 = _943_v7;
        let _lhs117 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_943_v7).length));
        _lhs116[_lhs117] = _rhs180;
        _937_v1 = _rhs181;
        _945_v9 = _rhs182;
        r0 = _rhs183;
        _946_v10 = _rhs184;
        let _951_v14;
        _951_v14 = _dafny.Seq.of(_dafny.MultiSet.fromElements(new _dafny.CodePoint('c'.codePointAt(0))));
        let _952_v15;
        _952_v15 = _module.D1.create_DC2(new BigNumber((_950_v13).length), (_this).f9, new BigNumber(503), p0);
        let _953_v16;
        _953_v16 = _dafny.Seq.of(function (_pat_let24_0) {
          return function (_954_dt__update__tmp_h0) {
            return function (_pat_let25_0) {
              return function (_955_dt__update_hcf3_h0) {
                return _module.D1.create_DC2((_954_dt__update__tmp_h0).dtor_cf2, _955_dt__update_hcf3_h0, (_954_dt__update__tmp_h0).dtor_cf4, (_954_dt__update__tmp_h0).dtor_cf5);
              }(_pat_let25_0);
            }((_this).f9);
          }(_pat_let24_0);
        }(_952_v15), _952_v15);
        let _956_v17;
        _956_v17 = new _dafny.CodePoint('e'.codePointAt(0));
        let _pat_let_tv26 = _943_v7;
        let _pat_let_tv27 = _943_v7;
        let _pat_let_tv28 = p0;
        let _pat_let_tv29 = _943_v7;
        let _pat_let_tv30 = _943_v7;
        let _pat_let_tv31 = p0;
        let _pat_let_tv32 = p1;
        let _pat_let_tv33 = _943_v7;
        let _pat_let_tv34 = _943_v7;
        let _pat_let_tv35 = p0;
        let _pat_let_tv36 = p0;
        let _pat_let_tv37 = _943_v7;
        let _pat_let_tv38 = _943_v7;
        let _pat_let_tv39 = p0;
        let _pat_let_tv40 = p1;
        let _pat_let_tv41 = _943_v7;
        let _pat_let_tv42 = _943_v7;
        let _pat_let_tv43 = p1;
        let _pat_let_tv44 = _956_v17;
        let _pat_let_tv45 = _936_v0;
        let _pat_let_tv46 = globalState;
        let _957_v18;
        let _nw147 = new _module.C1();
        _nw147.__ctor(((_module.__default.fm1(_dafny.MultiSet.fromElements(new BigNumber(637), new BigNumber(((_951_v14)[_module.__default.safeIndex(p0, new BigNumber((_951_v14).length))]).cardinality()), (_dafny.ZERO).minus((_943_v7)[_module.__default.safeIndex(new BigNumber(262), new BigNumber((_943_v7).length))])), p0, globalState)) ? (_953_v16) : (_dafny.Seq.of(function (_pat_let26_0) {
          return function (_958_dt__update__tmp_h1) {
            return function (_pat_let27_0) {
              return function (_959_dt__update_hcf5_h0) {
                return function (_pat_let28_0) {
                  return function (_962_dt__update_hcf4_h0) {
                    return function (_pat_let29_0) {
                      return function (_964_dt__update_hcf3_h1) {
                        return _module.D1.create_DC2((_958_dt__update__tmp_h1).dtor_cf2, _964_dt__update_hcf3_h1, _962_dt__update_hcf4_h0, _959_dt__update_hcf5_h0);
                      }(_pat_let29_0);
                    }(_module.__default.fm1(_pat_let_tv45, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-302)), function (_963_i1) {
                      return new _dafny.CodePoint('b'.codePointAt(0));
                    })).length), _pat_let_tv46));
                  }(_pat_let28_0);
                }(new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(949)), function (_960_i2) {
                  return new _dafny.CodePoint('c'.codePointAt(0));
                }), _module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_pat_let_tv28, (_pat_let_tv30)[_module.__default.safeIndex(new BigNumber(262), new BigNumber((_pat_let_tv29).length))], _pat_let_tv31, _pat_let_tv32, (_pat_let_tv34)[_module.__default.safeIndex(new BigNumber(262), new BigNumber((_pat_let_tv33).length))]), _module.__default.safeIndex(_pat_let_tv35, new BigNumber((_dafny.Seq.of(_pat_let_tv36, (_pat_let_tv38)[_module.__default.safeIndex(new BigNumber(262), new BigNumber((_pat_let_tv37).length))], _pat_let_tv39, _pat_let_tv40, (_pat_let_tv42)[_module.__default.safeIndex(new BigNumber(262), new BigNumber((_pat_let_tv41).length))])).length)), _pat_let_tv43)).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(949)), function (_961_i2) {
                  return new _dafny.CodePoint('c'.codePointAt(0));
                })).length)), _pat_let_tv44)).length));
              }(_pat_let27_0);
            }((_pat_let_tv27)[_module.__default.safeIndex(new BigNumber(262), new BigNumber((_pat_let_tv26).length))]);
          }(_pat_let26_0);
        }(_952_v15), _952_v15, _952_v15, _952_v15, _952_v15))));
        _957_v18 = _nw147;
        let _965_v19;
        _965_v19 = _dafny.Seq.of(_module.__default.fm30(globalState));
        _938_v2 = (_965_v19)[_module.__default.safeIndex((_943_v7)[_module.__default.safeIndex(new BigNumber(262), new BigNumber((_943_v7).length))], new BigNumber((_965_v19).length))];
        let _966_v20;
        _966_v20 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_943_v7)[_module.__default.safeIndex(new BigNumber(262), new BigNumber((_943_v7).length))]),_956_v17);
        let _967_v21;
        _967_v21 = _dafny.Seq.of(_940_v4, _940_v4, _940_v4);
        let _968_v22;
        _968_v22 = _dafny.Set.fromElements(p2, (_this).f9, p2);
        let _969_v23;
        _969_v23 = _949_v12;
        let _970_v24;
        let _nw148 = Array((new BigNumber(19)).toNumber());
        _nw148[(_dafny.ZERO).toNumber()] = (_939_v3).IsDisjointFrom(_dafny.MultiSet.fromElements((_this).f9));
        _nw148[(_dafny.ONE).toNumber()] = p2;
        _nw148[(new BigNumber(2)).toNumber()] = !((_this).f9);
        _nw148[(new BigNumber(3)).toNumber()] = p2;
        _nw148[(new BigNumber(4)).toNumber()] = (p1).isEqualTo(new BigNumber((_966_v20).length));
        _nw148[(new BigNumber(5)).toNumber()] = (_this).f9;
        _nw148[(new BigNumber(6)).toNumber()] = (_this).f9;
        _nw148[(new BigNumber(7)).toNumber()] = (_this).f9;
        _nw148[(new BigNumber(8)).toNumber()] = p2;
        _nw148[(new BigNumber(9)).toNumber()] = _module.__default.fm1((_dafny.MultiSet.fromElements(p1)).update(p0, _module.__default.abs(new BigNumber(((_967_v21)[_module.__default.safeIndex((_943_v7)[_module.__default.safeIndex(new BigNumber(262), new BigNumber((_943_v7).length))], new BigNumber((_967_v21).length))]).length))), new BigNumber((_968_v22).length), globalState);
        _nw148[(new BigNumber(10)).toNumber()] = (_this).f9;
        _nw148[(new BigNumber(11)).toNumber()] = p2;
        _nw148[(new BigNumber(12)).toNumber()] = p2;
        _nw148[(new BigNumber(13)).toNumber()] = (_this).f9;
        _nw148[(new BigNumber(14)).toNumber()] = (new BigNumber(559)).isEqualTo((_dafny.ZERO).minus(p1));
        _nw148[(new BigNumber(15)).toNumber()] = false;
        _nw148[(new BigNumber(16)).toNumber()] = (p2) === (!(p2));
        _nw148[(new BigNumber(17)).toNumber()] = false;
        _nw148[(new BigNumber(18)).toNumber()] = _dafny.Seq.contains((_dafny.Seq.of((_943_v7)[_module.__default.safeIndex(new BigNumber(262), new BigNumber((_943_v7).length))], (_943_v7)[_module.__default.safeIndex(new BigNumber(262), new BigNumber((_943_v7).length))])), p0);
        _970_v24 = _nw148;
        let _index145 = _module.__default.safeIndex(new BigNumber(745), new BigNumber((_970_v24).length));
        (_970_v24)[_index145] = (_this).f9;
        let _index146 = _module.__default.safeIndex(new BigNumber(91), new BigNumber((_970_v24).length));
        (_970_v24)[_index146] = !(!_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-568)), ((_971_v17) => function (_972_i3) {
          return _971_v17;
        })(_956_v17)), _956_v17));
        let _973_v25;
        _973_v25 = _dafny.Map.Empty.slice().updateUnsafe(false,_956_v17);
        let _index147 = _module.__default.safeIndex(new BigNumber(745), new BigNumber((_970_v24).length));
        let _index148 = _module.__default.safeIndex(new BigNumber(91), new BigNumber((_970_v24).length));
        let _rhs185 = p1;
        let _rhs186 = !(_939_v3).contains(p2);
        let _rhs187 = (p0).isEqualTo(((_943_v7)[_module.__default.safeIndex(new BigNumber(262), new BigNumber((_943_v7).length))]).multipliedBy(new BigNumber((_973_v25).length)));
        let _rhs188 = !(p2) || (false);
        let _rhs189 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Concat(_950_v13, _module.__default.fm28(p1, !((_this).f9), globalState)), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.Concat(_950_v13, _module.__default.fm28(p1, !((_this).f9), globalState))).length)), _956_v17), _950_v13);
        let _lhs118 = globalState;
        let _lhs119 = _970_v24;
        let _lhs120 = _module.__default.safeIndex(new BigNumber(745), new BigNumber((_970_v24).length));
        let _lhs121 = _970_v24;
        let _lhs122 = _module.__default.safeIndex(new BigNumber(91), new BigNumber((_970_v24).length));
        r0 = _rhs185;
        _lhs118.f2 = _rhs186;
        _lhs119[_lhs120] = _rhs187;
        _lhs121[_lhs122] = _rhs188;
        _950_v13 = _rhs189;
        let _974_v26;
        _974_v26 = _module.D10.create_DC27(_940_v4);
        _940_v4 = (((_974_v26).dtor_cf50).Merge(_940_v4)).Merge(_940_v4);
      } else {
        r0 = (p1).multipliedBy(p1);
        let _975_v27;
        _975_v27 = _dafny.Seq.of(p2, true);
        let _976_v28;
        _976_v28 = _dafny.Set.fromElements(p2);
        let _977_v29;
        let _nw149 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Map.Empty);
        _977_v29 = _nw149;
        let _978_v30;
        _978_v30 = _dafny.Seq.UnicodeFromString("n");
        let _979_v31;
        _979_v31 = _module.D7.create_DC19(_977_v29, (_this).f9, _978_v30);
        let _980_v32;
        _980_v32 = _module.D7.create_DC21(_979_v31);
        let _981_v33;
        _981_v33 = _dafny.Set.fromElements(_980_v32, _980_v32);
        let _982_v34;
        _982_v34 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(p0));
        let _983_v35;
        let _nw150 = Array((new BigNumber(19)).toNumber());
        _nw150[(_dafny.ZERO).toNumber()] = (_975_v27)[_module.__default.safeIndex(p1, new BigNumber((_975_v27).length))];
        _nw150[(_dafny.ONE).toNumber()] = (_this).f9;
        _nw150[(new BigNumber(2)).toNumber()] = p2;
        _nw150[(new BigNumber(3)).toNumber()] = p2;
        _nw150[(new BigNumber(4)).toNumber()] = (_976_v28).IsSubsetOf(_976_v28);
        _nw150[(new BigNumber(5)).toNumber()] = (_981_v33).IsProperSubsetOf(_981_v33);
        _nw150[(new BigNumber(6)).toNumber()] = true;
        _nw150[(new BigNumber(7)).toNumber()] = _module.__default.fm1(_982_v34, (_dafny.ZERO).minus(p0), globalState);
        _nw150[(new BigNumber(8)).toNumber()] = (_this).f9;
        _nw150[(new BigNumber(9)).toNumber()] = (_this).f9;
        _nw150[(new BigNumber(10)).toNumber()] = (_this).f9;
        _nw150[(new BigNumber(11)).toNumber()] = p2;
        _nw150[(new BigNumber(12)).toNumber()] = !((_this).f9) || (true);
        _nw150[(new BigNumber(13)).toNumber()] = !(false);
        _nw150[(new BigNumber(14)).toNumber()] = !(false);
        _nw150[(new BigNumber(15)).toNumber()] = !((_976_v28).IsProperSubsetOf(_976_v28));
        _nw150[(new BigNumber(16)).toNumber()] = (_this).f9;
        _nw150[(new BigNumber(17)).toNumber()] = p2;
        _nw150[(new BigNumber(18)).toNumber()] = p2;
        _983_v35 = _nw150;
        let _index149 = _module.__default.safeIndex(new BigNumber(877), new BigNumber((_983_v35).length));
        (_983_v35)[_index149] = (_this).f9;
        let _index150 = _module.__default.safeIndex(new BigNumber(877), new BigNumber((_983_v35).length));
        (_983_v35)[_index150] = (_this).f9;
        let _984_v36;
        _984_v36 = _dafny.Set.fromElements(new BigNumber((_978_v30).length), p1);
        let _985_v37;
        _985_v37 = _dafny.Map.Empty.slice().updateUnsafe(_984_v36,p2);
        let _index151 = _module.__default.safeIndex(new BigNumber(877), new BigNumber((_983_v35).length));
        (_983_v35)[_index151] = !((((_985_v37).contains(_984_v36)) ? ((_985_v37).get(_984_v36)) : (!(!(!((_983_v35)[_module.__default.safeIndex(new BigNumber(877), new BigNumber((_983_v35).length))]))) || (true))));
        let _986_v38;
        _986_v38 = new _dafny.CodePoint('y'.codePointAt(0));
        let _987_v39;
        _987_v39 = _module.D1.create_DC3((_this).f9, _dafny.Seq.update(_dafny.Seq.UnicodeFromString("vwjn"), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.UnicodeFromString("vwjn")).length)), _986_v38), p0);
        let _988_v40;
        _988_v40 = _dafny.Set.fromElements(_986_v38, new _dafny.CodePoint('j'.codePointAt(0)), _986_v38, _986_v38);
        let _989_v41;
        _989_v41 = _dafny.Seq.of(p1);
        let _990_v42;
        _990_v42 = _dafny.Map.Empty.slice().updateUnsafe(p2,_978_v30);
        let _pat_let_tv47 = _978_v30;
        let _pat_let_tv48 = p1;
        let _pat_let_tv49 = p1;
        let _pat_let_tv50 = p1;
        let _pat_let_tv51 = _988_v40;
        let _pat_let_tv52 = p1;
        let _pat_let_tv53 = p2;
        let _pat_let_tv54 = _978_v30;
        let _pat_let_tv55 = p1;
        let _pat_let_tv56 = p1;
        let _pat_let_tv57 = p1;
        let _pat_let_tv58 = _988_v40;
        let _pat_let_tv59 = p1;
        let _pat_let_tv60 = p2;
        let _991_v43;
        let _nw151 = Array((new BigNumber(18)).toNumber());
        _nw151[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_978_v30, _978_v30);
        _nw151[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update((function (_pat_let30_0) {
          return function (_995_dt__update__tmp_h4) {
            return function (_pat_let34_0) {
              return function (_996_dt__update_hcf8_h0) {
                return function (_pat_let35_0) {
                  return function (_997_dt__update_hcf6_h1) {
                    return _module.D1.create_DC3(_997_dt__update_hcf6_h1, (_995_dt__update__tmp_h4).dtor_cf7, _996_dt__update_hcf8_h0);
                  }(_pat_let35_0);
                }(_pat_let_tv53);
              }(_pat_let34_0);
            }(new BigNumber((_dafny.Seq.of(_pat_let_tv48, _pat_let_tv49, _pat_let_tv50, new BigNumber((_pat_let_tv51).length), _pat_let_tv52)).length));
          }(_pat_let30_0);
        }(function (_pat_let31_0) {
          return function (_992_dt__update__tmp_h3) {
            return function (_pat_let32_0) {
              return function (_993_dt__update_hcf7_h0) {
                return function (_pat_let33_0) {
                  return function (_994_dt__update_hcf6_h0) {
                    return _module.D1.create_DC3(_994_dt__update_hcf6_h0, _993_dt__update_hcf7_h0, (_992_dt__update__tmp_h3).dtor_cf8);
                  }(_pat_let33_0);
                }(true);
              }(_pat_let32_0);
            }(_pat_let_tv47);
          }(_pat_let31_0);
        }(_987_v39))).dtor_cf7, _module.__default.safeIndex(new BigNumber((_989_v41).length), new BigNumber(((function (_pat_let36_0) {
          return function (_1001_dt__update__tmp_h6) {
            return function (_pat_let40_0) {
              return function (_1002_dt__update_hcf8_h1) {
                return function (_pat_let41_0) {
                  return function (_1003_dt__update_hcf6_h3) {
                    return _module.D1.create_DC3(_1003_dt__update_hcf6_h3, (_1001_dt__update__tmp_h6).dtor_cf7, _1002_dt__update_hcf8_h1);
                  }(_pat_let41_0);
                }(_pat_let_tv60);
              }(_pat_let40_0);
            }(new BigNumber((_dafny.Seq.of(_pat_let_tv55, _pat_let_tv56, _pat_let_tv57, new BigNumber((_pat_let_tv58).length), _pat_let_tv59)).length));
          }(_pat_let36_0);
        }(function (_pat_let37_0) {
          return function (_998_dt__update__tmp_h5) {
            return function (_pat_let38_0) {
              return function (_999_dt__update_hcf7_h1) {
                return function (_pat_let39_0) {
                  return function (_1000_dt__update_hcf6_h2) {
                    return _module.D1.create_DC3(_1000_dt__update_hcf6_h2, _999_dt__update_hcf7_h1, (_998_dt__update__tmp_h5).dtor_cf8);
                  }(_pat_let39_0);
                }(true);
              }(_pat_let38_0);
            }(_pat_let_tv54);
          }(_pat_let37_0);
        }(_987_v39))).dtor_cf7).length)), _986_v38), _dafny.Seq.UnicodeFromString("ltwca"));
        _nw151[(new BigNumber(2)).toNumber()] = _978_v30;
        _nw151[(new BigNumber(3)).toNumber()] = _978_v30;
        _nw151[(new BigNumber(4)).toNumber()] = _978_v30;
        _nw151[(new BigNumber(5)).toNumber()] = (((_983_v35)[_module.__default.safeIndex(new BigNumber(877), new BigNumber((_983_v35).length))]) ? (_978_v30) : (_dafny.Seq.update(_978_v30, _module.__default.safeIndex(p1, new BigNumber((_978_v30).length)), _986_v38)));
        _nw151[(new BigNumber(6)).toNumber()] = _978_v30;
        _nw151[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_978_v30, _978_v30);
        _nw151[(new BigNumber(8)).toNumber()] = _module.__default.fm28(p1, (_983_v35)[_module.__default.safeIndex(new BigNumber(877), new BigNumber((_983_v35).length))], globalState);
        _nw151[(new BigNumber(9)).toNumber()] = _978_v30;
        _nw151[(new BigNumber(10)).toNumber()] = _dafny.Seq.UnicodeFromString("hrwvirwn");
        _nw151[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_978_v30, _978_v30);
        _nw151[(new BigNumber(12)).toNumber()] = _dafny.Seq.UnicodeFromString("silymvy");
        _nw151[(new BigNumber(13)).toNumber()] = _978_v30;
        _nw151[(new BigNumber(14)).toNumber()] = (((_990_v42).contains(!((_983_v35)[_module.__default.safeIndex(new BigNumber(877), new BigNumber((_983_v35).length))]))) ? ((_990_v42).get(!((_983_v35)[_module.__default.safeIndex(new BigNumber(877), new BigNumber((_983_v35).length))]))) : (_dafny.Seq.UnicodeFromString("jmvllgg")));
        _nw151[(new BigNumber(15)).toNumber()] = _978_v30;
        _nw151[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-884)), ((_1004_v38) => function (_1005_i4) {
          return _1004_v38;
        })(_986_v38)), _978_v30);
        _nw151[(new BigNumber(17)).toNumber()] = _978_v30;
        _991_v43 = _nw151;
        let _index152 = _module.__default.safeIndex(new BigNumber(370), new BigNumber((_991_v43).length));
        (_991_v43)[_index152] = _978_v30;
        let _1006_v44;
        _1006_v44 = _dafny.Map.Empty.slice().updateUnsafe(p0,false);
        let _1007_v45;
        _1007_v45 = _module.D9.create_DC24((((_1006_v44).contains(p0)) ? ((_1006_v44).get(p0)) : (p2)), p0, _dafny.Seq.UnicodeFromString("i"));
        let _index153 = _module.__default.safeIndex(new BigNumber(370), new BigNumber((_991_v43).length));
        (_991_v43)[_index153] = (_1007_v45).dtor_cf43;
        if ((_975_v27)[_module.__default.safeIndex(p0, new BigNumber((_975_v27).length))]) {
          _986_v38 = _module.__default.fm31(globalState);
          let _1008_v46;
          _1008_v46 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber(470), p1, p1, p0),(_this).f9);
          let _1009_v47;
          _1009_v47 = _module.D7.create_DC18((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_989_v41),p2)).Merge(_1008_v46));
          let _index154 = _module.__default.safeIndex(new BigNumber(877), new BigNumber((_983_v35).length));
          let _rhs190 = (_dafny.ZERO).minus((_dafny.ZERO).minus(p1));
          let _rhs191 = p2;
          let _rhs192 = _1009_v47;
          let _lhs123 = _983_v35;
          let _lhs124 = _module.__default.safeIndex(new BigNumber(877), new BigNumber((_983_v35).length));
          r0 = _rhs190;
          _lhs123[_lhs124] = _rhs191;
          _1009_v47 = _rhs192;
          r2 = _module.__default.fm21(new BigNumber((_dafny.Seq.update(_975_v27, _module.__default.safeIndex(p0, new BigNumber((_975_v27).length)), (_this).f9)).length), _module.__default.fm1(_982_v34, p1, globalState), globalState);
          let _1010_v48;
          _1010_v48 = _module.D4.create_DC14(p2, p2, p1);
          r0 = (_1010_v48).dtor_cf28;
          (globalState).f2 = false;
        } else {
          let _1011_v49;
          let _nw152 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
          _1011_v49 = _nw152;
          let _index155 = _module.__default.safeIndex(new BigNumber(555), new BigNumber((_1011_v49).length));
          (_1011_v49)[_index155] = _module.__default.safeDivisionInt(p0, new BigNumber((_dafny.Seq.of((_983_v35)[_module.__default.safeIndex(new BigNumber(877), new BigNumber((_983_v35).length))])).length));
          let _index156 = _module.__default.safeIndex(new BigNumber(555), new BigNumber((_1011_v49).length));
          (_1011_v49)[_index156] = _module.__default.safeModuloInt(new BigNumber(-642), (_dafny.ZERO).minus(p0));
          _978_v30 = _978_v30;
          let _1012_v50;
          let _nw153 = new _module.C2();
          _nw153.__ctor(new BigNumber((_dafny.Seq.Concat(_978_v30, _978_v30)).length), (_983_v35)[_module.__default.safeIndex(new BigNumber(877), new BigNumber((_983_v35).length))]);
          _1012_v50 = _nw153;
          let _index157 = _module.__default.safeIndex(new BigNumber(370), new BigNumber((_991_v43).length));
          (_991_v43)[_index157] = _978_v30;
          let _1013_v51;
          _1013_v51 = _dafny.MultiSet.fromElements(p2);
          let _1014_v52;
          _1014_v52 = _module.D1.create_DC4(new BigNumber(385), (_983_v35)[_module.__default.safeIndex(new BigNumber(877), new BigNumber((_983_v35).length))], p2, (_983_v35)[_module.__default.safeIndex(new BigNumber(877), new BigNumber((_983_v35).length))], p2);
          r2 = (_1013_v51).update((_1014_v52).dtor_cf12, _module.__default.abs(new BigNumber((_975_v27).length)));
        }
      }
      let _source11 = _module.D3.create_DC10();
      if (_source11.is_DC8) {
        let _1015_v53;
        _1015_v53 = _dafny.Seq.UnicodeFromString("ujap");
        r0 = new BigNumber((_1015_v53).length);
        let _1016_v54;
        let _nw154 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
        _1016_v54 = _nw154;
        let _1017_v55;
        _1017_v55 = _1016_v54;
        _1016_v54 = (_1017_v55);
        let _source12 = _module.D1.create_DC3(p2, _1015_v53, p1);
        if (_source12.is_DC2) {
          let _1018___mcc_h3 = (_source12).cf2;
          let _1019___mcc_h4 = (_source12).cf3;
          let _1020___mcc_h5 = (_source12).cf4;
          let _1021___mcc_h6 = (_source12).cf5;
          let _1022_cf5 = _1021___mcc_h6;
          let _1023_cf4 = _1020___mcc_h5;
          let _1024_cf3 = _1019___mcc_h4;
          let _1025_cf2 = _1018___mcc_h3;
          _1023_cf4 = (_1022_cf5).multipliedBy(_1022_cf5);
          let _1026_v56;
          _1026_v56 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_1024_cf3);
          let _1027_v57;
          _1027_v57 = _dafny.MultiSet.fromElements(p0, new BigNumber((_1026_v56).length));
          let _1028_v58;
          _1028_v58 = new _dafny.CodePoint('h'.codePointAt(0));
          let _1029_v59;
          _1029_v59 = _module.D7.create_DC20(false, _1028_v58);
          let _1030_v60;
          _1030_v60 = _module.D4.create_DC15(_module.D4.create_DC12(_module.__default.fm1(_1027_v57, _1023_cf4, globalState), false, _1025_cf2, _1025_cf2, (_1029_v59).dtor_cf36));
          _1030_v60 = _1030_v60;
          let _1031_v61;
          let _init28 = ((_1032_v53, _1033_p0, _1034_v58, _1035_p1) => function (_1036_i5) {
            return !((new BigNumber((_dafny.Seq.update(_1032_v53, _module.__default.safeIndex((_dafny.ZERO).minus(_1033_p0), new BigNumber((_1032_v53).length)), _1034_v58)).length)).isLessThan(_1035_p1));
          })(_1015_v53, p0, _1028_v58, p1);
          let _nw155 = Array((new BigNumber(24)).toNumber());
          for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw155.length); _i0_28++) {
            _nw155[_i0_28] = _init28(new BigNumber(_i0_28));
          }
          _1031_v61 = _nw155;
          _1031_v61 = _1031_v61;
          let _1037_v62;
          let _nw156 = new _module.C2();
          _nw156.__ctor(_1025_cf2, p2);
          _1037_v62 = _nw156;
          let _1038_v63;
          _1038_v63 = _dafny.Map.Empty.slice().updateUnsafe(_1037_v62,_1024_cf3);
          let _1039_v64;
          let _nw157 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Set.Empty);
          _1039_v64 = _nw157;
          let _1040_v65;
          let _nw158 = new _module.C3();
          _nw158.__ctor((((_1038_v63).contains(_1037_v62)) ? ((_1038_v63).get(_1037_v62)) : (!(false))), _1039_v64, _1024_cf3);
          _1040_v65 = _nw158;
          let _1041_v66;
          _1041_v66 = _dafny.Seq.of(_1040_v65, _1040_v65);
          r3 = !(_module.__default.fm1(_1027_v57, new BigNumber((_1041_v66).length), globalState));
        } else if (_source12.is_DC3) {
          let _1042___mcc_h7 = (_source12).cf6;
          let _1043___mcc_h8 = (_source12).cf7;
          let _1044___mcc_h9 = (_source12).cf8;
          let _1045_cf8 = _1044___mcc_h9;
          let _1046_cf7 = _1043___mcc_h8;
          let _1047_cf6 = _1042___mcc_h7;
          (globalState).f2 = _1047_cf6;
          let _1048_v67;
          let _init29 = ((_1049_cf6, _1050_p2) => function (_1051_i6) {
            return ((!(_1049_cf6)) ? (_1049_cf6) : (_1050_p2));
          })(_1047_cf6, p2);
          let _nw159 = Array((new BigNumber(14)).toNumber());
          for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw159.length); _i0_29++) {
            _nw159[_i0_29] = _init29(new BigNumber(_i0_29));
          }
          _1048_v67 = _nw159;
          let _1052_v68;
          _1052_v68 = _dafny.Seq.of(p1, _1045_cf8);
          let _1053_v69;
          _1053_v69 = _dafny.Map.Empty.slice().updateUnsafe(_1052_v68,p0);
          let _1054_v71;
          let _nw160 = new _module.C0();
          _nw160.__ctor(function () {
            let _coll23 = new _dafny.Set();
            for (const _compr_23 of (_1053_v69).Keys.Elements) {
              let _1055_v70 = _compr_23;
              if ((_1053_v69).contains(_1055_v70)) {
                _coll23.add(_1055_v70);
              }
            }
            return _coll23;
          }(), _1045_cf8);
          _1054_v71 = _nw160;
          let _1056_v72;
          _1056_v72 = _dafny.MultiSet.fromElements(_1054_v71, _1054_v71, _1054_v71);
          let _index158 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((_1048_v67).length));
          (_1048_v67)[_index158] = _module.__default.fm1(_dafny.MultiSet.fromElements(p1, p0), new BigNumber((_1056_v72).cardinality()), globalState);
          let _index159 = _module.__default.safeIndex(new BigNumber(497), new BigNumber((_1048_v67).length));
          (_1048_v67)[_index159] = _1047_cf6;
          let _1057_v73;
          _1057_v73 = _dafny.Map.Empty.slice().updateUnsafe(_1054_v71.f12,(_1054_v71.f12).plus((_dafny.ZERO).minus((_dafny.ZERO).minus(p0))));
          _1057_v73 = (_1057_v73).update((_dafny.ZERO).minus(p0), _module.__default.fm0(globalState));
          let _1058_v74;
          let _nw161 = new _module.C2();
          _nw161.__ctor(_1045_cf8, _1047_cf6);
          _1058_v74 = _nw161;
          let _1059_v75;
          _1059_v75 = _dafny.Seq.of(_1058_v74, _1058_v74);
          let _1060_v76;
          _1060_v76 = _dafny.Map.Empty.slice().updateUnsafe(_1059_v75,_1048_v67);
          let _1061_v77;
          let _nw162 = Array((new BigNumber(4)).toNumber());
          _nw162[(_dafny.ZERO).toNumber()] = (((_1060_v76).contains(_1059_v75)) ? ((_1060_v76).get(_1059_v75)) : (_1048_v67));
          _nw162[(_dafny.ONE).toNumber()] = _1048_v67;
          _nw162[(new BigNumber(2)).toNumber()] = _1048_v67;
          _nw162[(new BigNumber(3)).toNumber()] = _1048_v67;
          _1061_v77 = _nw162;
          let _index160 = _module.__default.safeIndex(new BigNumber(469), new BigNumber((_1061_v77).length));
          (_1061_v77)[_index160] = _1048_v67;
          let _1062_v78;
          _1062_v78 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1015_v53).length),_1048_v67);
          let _index161 = _module.__default.safeIndex(new BigNumber(469), new BigNumber((_1061_v77).length));
          (_1061_v77)[_index161] = (((_1062_v78).contains(_1054_v71.f12)) ? ((_1062_v78).get(_1054_v71.f12)) : (_1048_v67));
        } else if (_source12.is_DC4) {
          let _1063___mcc_h10 = (_source12).cf9;
          let _1064___mcc_h11 = (_source12).cf10;
          let _1065___mcc_h12 = (_source12).cf11;
          let _1066___mcc_h13 = (_source12).cf12;
          let _1067___mcc_h14 = (_source12).cf13;
          let _1068_cf13 = _1067___mcc_h14;
          let _1069_cf12 = _1066___mcc_h13;
          let _1070_cf11 = _1065___mcc_h12;
          let _1071_cf10 = _1064___mcc_h11;
          let _1072_cf9 = _1063___mcc_h10;
          let _1073_v79;
          _1073_v79 = new _dafny.CodePoint('d'.codePointAt(0));
          _1073_v79 = _1073_v79;
          let _1074_v80;
          let _nw163 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Set.Empty);
          _1074_v80 = _nw163;
          _1074_v80 = _1074_v80;
          (globalState).f2 = true;
          (globalState).f2 = (true) && ((_this).f9);
        } else if (_source12.is_DC1) {
          let _1075___mcc_h15 = (_source12).cf1;
          let _1076_cf1 = _1075___mcc_h15;
          _1076_cf1 = _module.__default.safeModuloInt(p1, p0);
          let _1077_v81;
          _1077_v81 = new _dafny.CodePoint('x'.codePointAt(0));
          _1015_v53 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1015_v53, _1015_v53), _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("jor"), _module.__default.safeIndex(new BigNumber((_1015_v53).length), new BigNumber((_dafny.Seq.UnicodeFromString("jor")).length)), _1077_v81), _module.__default.safeIndex(_1076_cf1, new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("jor"), _module.__default.safeIndex(new BigNumber((_1015_v53).length), new BigNumber((_dafny.Seq.UnicodeFromString("jor")).length)), _1077_v81)).length)), _module.__default.fm31(globalState)));
          r3 = true;
          r3 = p2;
        } else {
          let _1078___mcc_h16 = (_source12).cf14;
          let _1079_cf14 = _1078___mcc_h16;
          let _1080_v82;
          _1080_v82 = _dafny.Map.Empty.slice().updateUnsafe(!(p2),p2);
          r0 = (p0).multipliedBy((p1).minus(new BigNumber((_1080_v82).length)));
          let _1081_v83;
          let _init30 = function (_1082_i7) {
            return _dafny.Seq.UnicodeFromString("jmokwm");
          };
          let _nw164 = Array((new BigNumber(17)).toNumber());
          for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw164.length); _i0_30++) {
            _nw164[_i0_30] = _init30(new BigNumber(_i0_30));
          }
          _1081_v83 = _nw164;
          _1081_v83 = _1081_v83;
          r0 = _module.__default.safeDivisionInt(p1, new BigNumber(((_1080_v82).Merge(_1080_v82)).length));
          let _1083_v84;
          _1083_v84 = _module.D1.create_DC2(p1, true, p1, _module.__default.fm0(globalState));
          let _1084_v85;
          _1084_v85 = _dafny.Seq.of((_this).f9);
          let _1085_v86;
          let _nw165 = new _module.C1();
          _nw165.__ctor(_dafny.Seq.Concat(_dafny.Seq.of(_1083_v84, _1083_v84, _1083_v84, _1083_v84, _module.D1.create_DC2(new BigNumber(-578), (_this).f9, new BigNumber((_1084_v85).length), p0)), _dafny.Seq.of(_1083_v84, _1083_v84, _1083_v84, _1083_v84, _module.D1.create_DC2(_module.__default.fm0(globalState), p2, p0, p0))));
          _1085_v86 = _nw165;
        }
        let _1086_v87;
        _1086_v87 = _dafny.Map.Empty.slice().updateUnsafe(_1016_v54,(_this).f9);
        _1086_v87 = (_1086_v87).update(_1016_v54, !((_this).f9));
      } else if (_source11.is_DC9) {
        let _1087___mcc_h0 = (_source11).cf17;
        let _1088___mcc_h1 = (_source11).cf18;
        let _1089_cf18 = _1088___mcc_h1;
        let _1090_cf17 = _1087___mcc_h0;
        let _1091_v88;
        let _init31 = ((_1092_p1) => function (_1093_i8) {
          return _module.__default.safeDivisionInt(_1093_i8, _1092_p1);
        })(p1);
        let _nw166 = Array((new BigNumber(11)).toNumber());
        for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw166.length); _i0_31++) {
          _nw166[_i0_31] = _init31(new BigNumber(_i0_31));
        }
        _1091_v88 = _nw166;
        let _index162 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1091_v88).length));
        (_1091_v88)[_index162] = new BigNumber((_1089_cf18).length);
        let _index163 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1091_v88).length));
        (_1091_v88)[_index163] = p1;
        let _1094_v89;
        _1094_v89 = new _dafny.CodePoint('v'.codePointAt(0));
        let _1095_v90;
        _1095_v90 = _dafny.Map.Empty.slice().updateUnsafe(_1094_v89,_1094_v89);
        let _1096_v91;
        _1096_v91 = _dafny.Seq.of(_1094_v89, _1094_v89);
        _1095_v90 = ((_1095_v90).Merge(_1095_v90)).Merge((_1095_v90).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('u'.codePointAt(0)),(_1096_v91)[_module.__default.safeIndex((_1091_v88)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1091_v88).length))], new BigNumber((_1096_v91).length))])));
        let _1097_v92;
        _1097_v92 = _module.D10.create_DC28();
        let _1098_v93;
        _1098_v93 = _dafny.Map.Empty.slice().updateUnsafe(p2,p1);
        let _1099_v94;
        _1099_v94 = _dafny.Map.Empty.slice().updateUnsafe(_1097_v92,_1098_v93);
        let _1100_v95;
        _1100_v95 = _dafny.MultiSet.fromElements(new BigNumber((_1099_v94).length));
        let _rhs193 = (_dafny.ZERO).minus(((_1091_v88)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1091_v88).length))]).multipliedBy((_dafny.ZERO).minus(p0)));
        let _rhs194 = _module.__default.fm1(_1100_v95, (_1091_v88)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1091_v88).length))], globalState);
        let _rhs195 = p1;
        r0 = _rhs193;
        _1090_cf17 = _rhs194;
        r0 = _rhs195;
        let _1101_v96;
        _1101_v96 = _module.D1.create_DC2(p0, false, p0, new BigNumber((_dafny.Seq.UnicodeFromString("chnn")).length));
        let _1102_v97;
        _1102_v97 = _dafny.Seq.of(_1101_v96);
        let _1103_v98;
        let _nw167 = new _module.C1();
        _nw167.__ctor(_dafny.Seq.Concat(_1102_v97, _1102_v97));
        _1103_v98 = _nw167;
        let _1104_v99;
        let _nw168 = new _module.C2();
        _nw168.__ctor((_1091_v88)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1091_v88).length))], !(p2));
        _1104_v99 = _nw168;
        let _1105_v100;
        _1105_v100 = _module.D4.create_DC13((_1104_v99).f14);
        let _1106_v101;
        _1106_v101 = _dafny.Set.fromElements(_1090_cf17, p2);
        let _pat_let_tv61 = _1100_v95;
        let _pat_let_tv62 = _1091_v88;
        let _pat_let_tv63 = _1091_v88;
        let _pat_let_tv64 = p1;
        let _1107_v102;
        let _nw169 = Array((new BigNumber(25)).toNumber());
        _nw169[(_dafny.ZERO).toNumber()] = _1105_v100;
        _nw169[(_dafny.ONE).toNumber()] = _1105_v100;
        _nw169[(new BigNumber(2)).toNumber()] = _1105_v100;
        _nw169[(new BigNumber(3)).toNumber()] = _1105_v100;
        _nw169[(new BigNumber(4)).toNumber()] = _1105_v100;
        _nw169[(new BigNumber(5)).toNumber()] = _1105_v100;
        _nw169[(new BigNumber(6)).toNumber()] = _1105_v100;
        _nw169[(new BigNumber(7)).toNumber()] = function (_pat_let42_0) {
          return function (_1108_dt__update__tmp_h7) {
            return function (_pat_let43_0) {
              return function (_1109_dt__update_hcf25_h0) {
                return _module.D4.create_DC13(_1109_dt__update_hcf25_h0);
              }(_pat_let43_0);
            }(new BigNumber((_pat_let_tv61).cardinality()));
          }(_pat_let42_0);
        }(_1105_v100);
        _nw169[(new BigNumber(8)).toNumber()] = _1105_v100;
        _nw169[(new BigNumber(9)).toNumber()] = _1105_v100;
        _nw169[(new BigNumber(10)).toNumber()] = function (_pat_let44_0) {
          return function (_1110_dt__update__tmp_h8) {
            return function (_pat_let45_0) {
              return function (_1111_dt__update_hcf25_h1) {
                return _module.D4.create_DC13(_1111_dt__update_hcf25_h1);
              }(_pat_let45_0);
            }((_pat_let_tv63)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_pat_let_tv62).length))]);
          }(_pat_let44_0);
        }(_1105_v100);
        _nw169[(new BigNumber(11)).toNumber()] = _1105_v100;
        _nw169[(new BigNumber(12)).toNumber()] = _1105_v100;
        _nw169[(new BigNumber(13)).toNumber()] = _1105_v100;
        _nw169[(new BigNumber(14)).toNumber()] = _1105_v100;
        _nw169[(new BigNumber(15)).toNumber()] = function (_pat_let46_0) {
          return function (_1112_dt__update__tmp_h9) {
            return function (_pat_let47_0) {
              return function (_1113_dt__update_hcf25_h2) {
                return _module.D4.create_DC13(_1113_dt__update_hcf25_h2);
              }(_pat_let47_0);
            }(_pat_let_tv64);
          }(_pat_let46_0);
        }(_1105_v100);
        _nw169[(new BigNumber(16)).toNumber()] = _1105_v100;
        _nw169[(new BigNumber(17)).toNumber()] = _1105_v100;
        _nw169[(new BigNumber(18)).toNumber()] = _1105_v100;
        _nw169[(new BigNumber(19)).toNumber()] = _1105_v100;
        _nw169[(new BigNumber(20)).toNumber()] = _module.D4.create_DC13((_dafny.ZERO).minus(new BigNumber((_1106_v101).length)));
        _nw169[(new BigNumber(21)).toNumber()] = _1105_v100;
        _nw169[(new BigNumber(22)).toNumber()] = _module.D4.create_DC13(p0);
        _nw169[(new BigNumber(23)).toNumber()] = _1105_v100;
        _nw169[(new BigNumber(24)).toNumber()] = _1105_v100;
        _1107_v102 = _nw169;
        let _1114_v103;
        let _nw170 = Array((new BigNumber(13)).toNumber());
        _nw170[(_dafny.ZERO).toNumber()] = _1107_v102;
        _nw170[(_dafny.ONE).toNumber()] = _1107_v102;
        _nw170[(new BigNumber(2)).toNumber()] = _1107_v102;
        _nw170[(new BigNumber(3)).toNumber()] = _1107_v102;
        _nw170[(new BigNumber(4)).toNumber()] = _1107_v102;
        _nw170[(new BigNumber(5)).toNumber()] = _1107_v102;
        _nw170[(new BigNumber(6)).toNumber()] = _1107_v102;
        _nw170[(new BigNumber(7)).toNumber()] = _1107_v102;
        _nw170[(new BigNumber(8)).toNumber()] = _1107_v102;
        _nw170[(new BigNumber(9)).toNumber()] = _1107_v102;
        _nw170[(new BigNumber(10)).toNumber()] = _1107_v102;
        _nw170[(new BigNumber(11)).toNumber()] = _1107_v102;
        _nw170[(new BigNumber(12)).toNumber()] = _1107_v102;
        _1114_v103 = _nw170;
        let _index164 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1091_v88).length));
        let _rhs196 = _1103_v98;
        let _rhs197 = new BigNumber(416);
        let _rhs198 = _1104_v99;
        let _rhs199 = _1114_v103;
        let _lhs125 = _1091_v88;
        let _lhs126 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1091_v88).length));
        _1103_v98 = _rhs196;
        _lhs125[_lhs126] = _rhs197;
        _1104_v99 = _rhs198;
        _1114_v103 = _rhs199;
      } else if (_source11.is_DC10) {
        r3 = (_this).f9;
        let _1115_v104;
        _1115_v104 = new _dafny.CodePoint('k'.codePointAt(0));
        let _1116_v105;
        _1116_v105 = _dafny.Set.fromElements(_1115_v104, _1115_v104);
        let _1117_v106;
        _1117_v106 = _dafny.Seq.of(true);
        let _1118_v107;
        _1118_v107 = _dafny.Set.fromElements(p0);
        let _1119_v108;
        _1119_v108 = _dafny.MultiSet.fromElements(p1);
        let _1120_v109;
        let _nw171 = Array((new BigNumber(29)).toNumber());
        _nw171[(_dafny.ZERO).toNumber()] = (_this).f9;
        _nw171[(_dafny.ONE).toNumber()] = p2;
        _nw171[(new BigNumber(2)).toNumber()] = (_this).f9;
        _nw171[(new BigNumber(3)).toNumber()] = p2;
        _nw171[(new BigNumber(4)).toNumber()] = (_1117_v106)[_module.__default.safeIndex(new BigNumber((_1118_v107).length), new BigNumber((_1117_v106).length))];
        _nw171[(new BigNumber(5)).toNumber()] = p2;
        _nw171[(new BigNumber(6)).toNumber()] = (_this).f9;
        _nw171[(new BigNumber(7)).toNumber()] = false;
        _nw171[(new BigNumber(8)).toNumber()] = p2;
        _nw171[(new BigNumber(9)).toNumber()] = (_this).f9;
        _nw171[(new BigNumber(10)).toNumber()] = true;
        _nw171[(new BigNumber(11)).toNumber()] = (_this).f9;
        _nw171[(new BigNumber(12)).toNumber()] = (_this).f9;
        _nw171[(new BigNumber(13)).toNumber()] = (_this).f9;
        _nw171[(new BigNumber(14)).toNumber()] = !(p2);
        _nw171[(new BigNumber(15)).toNumber()] = _module.__default.fm1(_1119_v108, p1, globalState);
        _nw171[(new BigNumber(16)).toNumber()] = (_this).f9;
        _nw171[(new BigNumber(17)).toNumber()] = p2;
        _nw171[(new BigNumber(18)).toNumber()] = (_this).f9;
        _nw171[(new BigNumber(19)).toNumber()] = true;
        _nw171[(new BigNumber(20)).toNumber()] = (_this).f9;
        _nw171[(new BigNumber(21)).toNumber()] = _module.__default.fm1(_1119_v108, (_dafny.ZERO).minus(p1), globalState);
        _nw171[(new BigNumber(22)).toNumber()] = (_this).f9;
        _nw171[(new BigNumber(23)).toNumber()] = (_this).f9;
        _nw171[(new BigNumber(24)).toNumber()] = (_this).f9;
        _nw171[(new BigNumber(25)).toNumber()] = (_this).f9;
        _nw171[(new BigNumber(26)).toNumber()] = (_this).f9;
        _nw171[(new BigNumber(27)).toNumber()] = !(p2);
        _nw171[(new BigNumber(28)).toNumber()] = !((_this).f9);
        _1120_v109 = _nw171;
        let _1121_v110;
        _1121_v110 = _dafny.Map.Empty.slice().updateUnsafe((_1116_v105).equals(_1116_v105),_1120_v109);
        let _1122_v111;
        _1122_v111 = _dafny.Seq.of(_1121_v110);
        let _1123_v112;
        let _nw172 = new _module.C2();
        _nw172.__ctor(new BigNumber(-876), p2);
        _1123_v112 = _nw172;
        let _1124_v113;
        _1124_v113 = _dafny.Seq.of(_1123_v112);
        _1121_v110 = (((_1122_v111)[_module.__default.safeIndex(new BigNumber((_1124_v113).length), new BigNumber((_1122_v111).length))]).Merge(_1121_v110)).update(_module.__default.fm1(_dafny.MultiSet.fromElements(p1, p0), new BigNumber((_dafny.Set.fromElements((_this).f9)).length), globalState), _1120_v109);
        r0 = (_1123_v112).f14;
        r0 = (_dafny.ZERO).minus((p1).plus(_module.__default.fm0(globalState)));
      } else {
        let _1125___mcc_h2 = (_source11).cf16;
        let _1126_cf16 = _1125___mcc_h2;
        let _1127_v114;
        _1127_v114 = _module.D4.create_DC13((_dafny.ZERO).minus(p0));
        let _1128_v115;
        _1128_v115 = new _dafny.CodePoint('y'.codePointAt(0));
        let _1129_v116;
        let _nw173 = new _module.C1();
        _nw173.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(670)), ((_1130_p0, _1131_p2) => function (_1132_i9) {
          return _module.D1.create_DC2(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f9,_1130_p0)).length), _1131_p2, _1130_p0, _1130_p0);
        })(p0, p2)));
        _1129_v116 = _nw173;
        let _1133_v117;
        _1133_v117 = _dafny.MultiSet.fromElements((_module.D13.create_DC35(p2, _1128_v115, _1129_v116, (_this).f9)).dtor_cf61, !(false));
        let _1134_v118;
        _1134_v118 = _dafny.Map.Empty.slice().updateUnsafe(_1128_v115,(((_1133_v117).contains((_this).f9)) ? ((_1133_v117).get((_this).f9)) : (p0)));
        let _1135_v119;
        _1135_v119 = _dafny.Map.Empty.slice().updateUnsafe(_1127_v114,(((_1134_v118).contains(_1128_v115)) ? ((_1134_v118).get(_1128_v115)) : (p1)));
        _1135_v119 = (_1135_v119).update(_1127_v114, p1);
        let _1136_v120;
        let _nw174 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Map.Empty);
        _1136_v120 = _nw174;
        let _1137_v121;
        _1137_v121 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f9);
        let _index165 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((_1136_v120).length));
        (_1136_v120)[_index165] = (_1137_v121).Merge(_1137_v121);
        let _1138_v123;
        _1138_v123 = _dafny.MultiSet.fromElements(p1);
        let _index166 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((_1136_v120).length));
        (_1136_v120)[_index166] = (_1137_v121).Merge(function () {
          let _coll24 = new _dafny.Map();
          for (const _compr_24 of (_1138_v123).Elements) {
            let _1139_v122 = _compr_24;
            if ((_1138_v123).contains(_1139_v122)) {
              _coll24.push([_module.__default.safeModuloInt(_1139_v122, p0),false]);
            }
          }
          return _coll24;
        }());
        let _1140_v124;
        _1140_v124 = _dafny.Seq.of((_this).f9);
        let _1141_v125;
        _1141_v125 = _module.D4.create_DC12(p2, p2, new BigNumber((_1140_v124).length), p1, (_this).f9);
        let _1142_v126;
        _1142_v126 = _module.D4.create_DC15(((p2) ? (_1141_v125) : (_1141_v125)));
        let _source13 = _1142_v126;
        if (_source13.is_DC12) {
          let _1143___mcc_h17 = (_source13).cf20;
          let _1144___mcc_h18 = (_source13).cf21;
          let _1145___mcc_h19 = (_source13).cf22;
          let _1146___mcc_h20 = (_source13).cf23;
          let _1147___mcc_h21 = (_source13).cf24;
          let _1148_cf24 = _1147___mcc_h21;
          let _1149_cf23 = _1146___mcc_h20;
          let _1150_cf22 = _1145___mcc_h19;
          let _1151_cf21 = _1144___mcc_h18;
          let _1152_cf20 = _1143___mcc_h17;
          let _1153_v127;
          let _init32 = ((_1154_cf23) => function (_1155_i10) {
            return _module.__default.safeDivisionInt(_1155_i10, _1154_cf23);
          })(_1149_cf23);
          let _nw175 = Array((new BigNumber(6)).toNumber());
          for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw175.length); _i0_32++) {
            _nw175[_i0_32] = _init32(new BigNumber(_i0_32));
          }
          _1153_v127 = _nw175;
          r0 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1153_v127,_1153_v127)).length);
          _1153_v127 = _1153_v127;
          let _1156_v128;
          let _nw176 = new _module.C1();
          _nw176.__ctor(_1129_v116.f13);
          _1156_v128 = _nw176;
          let _1157_v129;
          _1157_v129 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1150_cf22,p0)).length));
          let _1158_v130;
          _1158_v130 = _dafny.Set.fromElements(p0);
          let _1159_v131;
          _1159_v131 = _dafny.Map.Empty.slice().updateUnsafe(_1151_cf21,new BigNumber((_dafny.Seq.UnicodeFromString("acs")).length));
          let _1160_v132;
          let _init33 = ((_1161_cf20) => function (_1162_i11) {
            return _1161_cf20;
          })(_1152_cf20);
          let _nw177 = Array((new BigNumber(2)).toNumber());
          for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw177.length); _i0_33++) {
            _nw177[_i0_33] = _init33(new BigNumber(_i0_33));
          }
          _1160_v132 = _nw177;
          let _1163_v133;
          _1163_v133 = _dafny.Map.Empty.slice().updateUnsafe(_1160_v132,_1150_cf22);
          let _1164_v134;
          let _nw178 = Array((new BigNumber(19)).toNumber());
          _nw178[(_dafny.ZERO).toNumber()] = _1158_v130;
          _nw178[(_dafny.ONE).toNumber()] = _1158_v130;
          _nw178[(new BigNumber(2)).toNumber()] = _1158_v130;
          _nw178[(new BigNumber(3)).toNumber()] = _1158_v130;
          _nw178[(new BigNumber(4)).toNumber()] = _1158_v130;
          _nw178[(new BigNumber(5)).toNumber()] = _1158_v130;
          _nw178[(new BigNumber(6)).toNumber()] = _dafny.Set.fromElements(p0);
          _nw178[(new BigNumber(7)).toNumber()] = _1158_v130;
          _nw178[(new BigNumber(8)).toNumber()] = _dafny.Set.fromElements(new BigNumber((_1159_v131).length));
          _nw178[(new BigNumber(9)).toNumber()] = _1158_v130;
          _nw178[(new BigNumber(10)).toNumber()] = _1158_v130;
          _nw178[(new BigNumber(11)).toNumber()] = _1158_v130;
          _nw178[(new BigNumber(12)).toNumber()] = _1158_v130;
          _nw178[(new BigNumber(13)).toNumber()] = _1158_v130;
          _nw178[(new BigNumber(14)).toNumber()] = _dafny.Set.fromElements(new BigNumber((_1163_v133).length), _1149_cf23, new BigNumber(748));
          _nw178[(new BigNumber(15)).toNumber()] = _1158_v130;
          _nw178[(new BigNumber(16)).toNumber()] = _1158_v130;
          _nw178[(new BigNumber(17)).toNumber()] = _1158_v130;
          _nw178[(new BigNumber(18)).toNumber()] = _1158_v130;
          _1164_v134 = _nw178;
          let _1165_v135;
          let _nw179 = new _module.C3();
          _nw179.__ctor(_1151_cf21, _1164_v134, _1152_cf20);
          _1165_v135 = _nw179;
          let _1166_v136;
          _1166_v136 = _dafny.Seq.of(_1165_v135, _1165_v135);
          let _1167_v137;
          _1167_v137 = _dafny.MultiSet.fromElements(_1128_v115);
          let _1168_v138;
          let _nw180 = Array((new BigNumber(24)).toNumber());
          _nw180[(_dafny.ZERO).toNumber()] = !(new BigNumber(62)).isEqualTo((_dafny.ZERO).minus(_1150_cf22));
          _nw180[(_dafny.ONE).toNumber()] = !(_module.__default.fm1(_1138_v123, new BigNumber(186), globalState));
          _nw180[(new BigNumber(2)).toNumber()] = _dafny.Seq.contains(_module.__default.fm28((_1157_v129)[_module.__default.safeIndex(new BigNumber((_1138_v123).cardinality()), new BigNumber((_1157_v129).length))], p2, globalState), _1128_v115);
          _nw180[(new BigNumber(3)).toNumber()] = (new BigNumber(-521)).isLessThan(new BigNumber((_1166_v136).length));
          _nw180[(new BigNumber(4)).toNumber()] = !(!(_1152_cf20));
          _nw180[(new BigNumber(5)).toNumber()] = true;
          _nw180[(new BigNumber(6)).toNumber()] = (_1165_v135).f9;
          _nw180[(new BigNumber(7)).toNumber()] = (_this).f9;
          _nw180[(new BigNumber(8)).toNumber()] = (_1167_v137).IsSubsetOf(_1167_v137);
          _nw180[(new BigNumber(9)).toNumber()] = (_1165_v135).f9;
          _nw180[(new BigNumber(10)).toNumber()] = true;
          _nw180[(new BigNumber(11)).toNumber()] = !(p2) || (!(_1152_cf20));
          _nw180[(new BigNumber(12)).toNumber()] = (_this).f9;
          _nw180[(new BigNumber(13)).toNumber()] = _1148_cf24;
          _nw180[(new BigNumber(14)).toNumber()] = (_this).f9;
          _nw180[(new BigNumber(15)).toNumber()] = p2;
          _nw180[(new BigNumber(16)).toNumber()] = _dafny.Seq.IsPrefixOf((_this).fm2(_1149_cf23, (_1165_v135).f9, !((((_1137_v121).contains(p0)) ? ((_1137_v121).get(p0)) : ((_1165_v135).f9))), _module.__default.fm0(globalState), globalState), _1157_v129);
          _nw180[(new BigNumber(17)).toNumber()] = false;
          _nw180[(new BigNumber(18)).toNumber()] = _1151_cf21;
          _nw180[(new BigNumber(19)).toNumber()] = p2;
          _nw180[(new BigNumber(20)).toNumber()] = _1151_cf21;
          _nw180[(new BigNumber(21)).toNumber()] = _1152_cf20;
          _nw180[(new BigNumber(22)).toNumber()] = _1152_cf20;
          _nw180[(new BigNumber(23)).toNumber()] = _1152_cf20;
          _1168_v138 = _nw180;
          let _index167 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_1160_v132).length));
          (_1160_v132)[_index167] = true;
          let _index168 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_1153_v127).length));
          (_1153_v127)[_index168] = _1150_cf22;
          let _1169_v139;
          _1169_v139 = _dafny.Seq.of(_1159_v131, _1159_v131, _1159_v131, _1159_v131);
          let _index169 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_1153_v127).length));
          (_1153_v127)[_index169] = new BigNumber((_1169_v139).length);
          let _1170_v140;
          _1170_v140 = _dafny.MultiSet.fromElements(_1153_v127, _1153_v127, _1153_v127);
          let _index170 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_1160_v132).length));
          let _index171 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_1153_v127).length));
          let _index172 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_1153_v127).length));
          let _rhs200 = (_1170_v140).IsDisjointFrom(_1170_v140);
          let _rhs201 = p1;
          let _rhs202 = (((_1159_v131).contains((_1133_v117).IsProperSubsetOf(_1133_v117))) ? ((_1159_v131).get((_1133_v117).IsProperSubsetOf(_1133_v117))) : (_1150_cf22));
          let _rhs203 = (_dafny.ZERO).minus(_1149_cf23);
          let _rhs204 = true;
          let _lhs127 = _1160_v132;
          let _lhs128 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_1160_v132).length));
          let _lhs129 = _1153_v127;
          let _lhs130 = _module.__default.safeIndex(new BigNumber(665), new BigNumber((_1153_v127).length));
          let _lhs131 = _1153_v127;
          let _lhs132 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_1153_v127).length));
          let _lhs133 = globalState;
          _lhs127[_lhs128] = _rhs200;
          _lhs129[_lhs130] = _rhs201;
          _lhs131[_lhs132] = _rhs202;
          _1150_cf22 = _rhs203;
          _lhs133.f2 = _rhs204;
        } else if (_source13.is_DC13) {
          let _1171___mcc_h22 = (_source13).cf25;
          let _1172_cf25 = _1171___mcc_h22;
          let _1173_v141;
          let _init34 = ((_1174_cf25) => function (_1175_i12) {
            return (_1175_i12).plus(_1174_cf25);
          })(_1172_cf25);
          let _nw181 = Array((new BigNumber(12)).toNumber());
          for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw181.length); _i0_34++) {
            _nw181[_i0_34] = _init34(new BigNumber(_i0_34));
          }
          _1173_v141 = _nw181;
          let _index173 = _module.__default.safeIndex(new BigNumber(39), new BigNumber((_1173_v141).length));
          (_1173_v141)[_index173] = (((_this).f9) ? (p1) : (_1172_cf25));
          let _index174 = _module.__default.safeIndex(new BigNumber(39), new BigNumber((_1173_v141).length));
          (_1173_v141)[_index174] = _module.__default.safeModuloInt(_1172_cf25, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(34)), ((_1176_v115) => function (_1177_i13) {
            return _1176_v115;
          })(_1128_v115))).length));
          let _1178_v142;
          _1178_v142 = _module.D3.create_DC9(_module.__default.fm1(_dafny.MultiSet.fromElements(new BigNumber(527), _1172_cf25, p1, _1172_cf25), p0, globalState), _dafny.Seq.of(!(true)));
          let _1179_v143;
          _1179_v143 = _dafny.Map.Empty.slice().updateUnsafe(((_dafny.ZERO).minus(p1)).isLessThan((_dafny.ZERO).minus(p1)),(_1178_v142).dtor_cf17);
          _1179_v143 = (_1179_v143).update(p2, !(false) || (p2));
          let _1180_v144;
          let _nw182 = Array((new BigNumber(3)).toNumber()).fill(false);
          _1180_v144 = _nw182;
          let _1181_v145;
          _1181_v145 = _dafny.Map.Empty.slice().updateUnsafe(_1180_v144,(_this).f9);
          let _1182_v146;
          _1182_v146 = _dafny.Seq.of(_1180_v144, _1180_v144);
          let _1183_v147;
          _1183_v147 = _dafny.Seq.UnicodeFromString("n");
          let _1184_v148;
          _1184_v148 = _dafny.MultiSet.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.update(_dafny.Seq.of(_1172_cf25), _module.__default.safeIndex((_1173_v141)[_module.__default.safeIndex(new BigNumber(39), new BigNumber((_1173_v141).length))], new BigNumber((_dafny.Seq.of(_1172_cf25)).length)), _1172_cf25)), _1138_v123);
          _1181_v145 = (_1181_v145).update((_1182_v146)[_module.__default.safeIndex(new BigNumber((_1183_v147).length), new BigNumber((_1182_v146).length))], !(_1184_v148).contains(_1138_v123));
          r0 = _module.__default.safeDivisionInt(p1, new BigNumber(84));
        } else if (_source13.is_DC14) {
          let _1185___mcc_h23 = (_source13).cf26;
          let _1186___mcc_h24 = (_source13).cf27;
          let _1187___mcc_h25 = (_source13).cf28;
          let _1188_cf28 = _1187___mcc_h25;
          let _1189_cf27 = _1186___mcc_h24;
          let _1190_cf26 = _1185___mcc_h23;
          let _1191_v149;
          let _nw183 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
          _1191_v149 = _nw183;
          let _index175 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_1191_v149).length));
          (_1191_v149)[_index175] = p1;
          let _1192_v150;
          _1192_v150 = _dafny.Seq.of(_1138_v123);
          let _1193_v151;
          _1193_v151 = _dafny.Map.Empty.slice().updateUnsafe(_1189_cf27,(_1192_v150)[_module.__default.safeIndex(p1, new BigNumber((_1192_v150).length))]);
          let _1194_v152;
          _1194_v152 = _dafny.Seq.of(p1, new BigNumber((_1140_v124).length));
          let _index176 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_1191_v149).length));
          let _rhs205 = ((((_1193_v151).contains(!((_this).f9))) ? ((_1193_v151).get(!((_this).f9))) : (_dafny.MultiSet.FromArray(_1194_v152)))).Union(_1138_v123);
          let _rhs206 = p1;
          let _lhs134 = _1191_v149;
          let _lhs135 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_1191_v149).length));
          _1138_v123 = _rhs205;
          _lhs134[_lhs135] = _rhs206;
          let _1195_v153;
          _1195_v153 = _dafny.MultiSet.FromArray(_1194_v152);
          let _1196_v154;
          _1196_v154 = _dafny.MultiSet.fromElements(_1195_v153);
          _1188_cf28 = (((_1196_v154).contains(_1195_v153)) ? ((_1196_v154).get(_1195_v153)) : (p0));
          let _1197_v155;
          _1197_v155 = _dafny.Seq.of(_1140_v124, _1140_v124, _1140_v124);
          let _1198_v156;
          _1198_v156 = _dafny.Seq.UnicodeFromString("ifap");
          let _1199_v157;
          let _nw184 = Array((new BigNumber(12)).toNumber());
          _nw184[(_dafny.ZERO).toNumber()] = _1140_v124;
          _nw184[(_dafny.ONE).toNumber()] = (_1197_v155)[_module.__default.safeIndex(p0, new BigNumber((_1197_v155).length))];
          _nw184[(new BigNumber(2)).toNumber()] = _1140_v124;
          _nw184[(new BigNumber(3)).toNumber()] = _module.__default.fm20((_this).f9, _1189_cf27, (_1191_v149)[_module.__default.safeIndex(new BigNumber(857), new BigNumber((_1191_v149).length))], _1198_v156, globalState);
          _nw184[(new BigNumber(4)).toNumber()] = _1140_v124;
          _nw184[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_1190_cf26), _1140_v124);
          _nw184[(new BigNumber(6)).toNumber()] = _1140_v124;
          _nw184[(new BigNumber(7)).toNumber()] = _1140_v124;
          _nw184[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_1140_v124, _1140_v124);
          _nw184[(new BigNumber(9)).toNumber()] = _dafny.Seq.update(_1140_v124, _module.__default.safeIndex(_1188_cf28, new BigNumber((_1140_v124).length)), true);
          _nw184[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_1140_v124, _1140_v124);
          _nw184[(new BigNumber(11)).toNumber()] = _dafny.Seq.update(_1140_v124, _module.__default.safeIndex(p1, new BigNumber((_1140_v124).length)), _1189_cf27);
          _1199_v157 = _nw184;
          let _index177 = _module.__default.safeIndex(new BigNumber(554), new BigNumber((_1199_v157).length));
          (_1199_v157)[_index177] = _dafny.Seq.Concat(_1140_v124, _1140_v124);
          let _1200_v158;
          _1200_v158 = _dafny.Map.Empty.slice().updateUnsafe(_1190_cf26,new BigNumber((_module.__default.fm28(new BigNumber(284), _module.__default.fm1(_1138_v123, p0, globalState), globalState)).length));
          let _1201_v159;
          _1201_v159 = _module.D10.create_DC27(_1200_v158);
          let _index178 = _module.__default.safeIndex(new BigNumber(554), new BigNumber((_1199_v157).length));
          (_1199_v157)[_index178] = _dafny.Seq.Concat(_dafny.Seq.update(_1140_v124, _module.__default.safeIndex(_1188_cf28, new BigNumber((_1140_v124).length)), p2), _module.__default.fm41(_1201_v159, (_1191_v149)[_module.__default.safeIndex(new BigNumber(857), new BigNumber((_1191_v149).length))], _1194_v152, globalState));
          _1190_cf26 = _1189_cf27;
        } else if (_source13.is_DC11) {
          let _1202___mcc_h26 = (_source13).cf19;
          let _1203_cf19 = _1202___mcc_h26;
          let _1204_v160;
          let _init35 = ((_1205_p2) => function (_1206_i14) {
            return _1205_p2;
          })(p2);
          let _nw185 = Array((new BigNumber(11)).toNumber());
          for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw185.length); _i0_35++) {
            _nw185[_i0_35] = _init35(new BigNumber(_i0_35));
          }
          _1204_v160 = _nw185;
          let _1207_v161;
          _1207_v161 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1204_v160);
          let _1208_v162;
          _1208_v162 = _dafny.Map.Empty.slice().updateUnsafe((_1207_v161).Merge(_1207_v161),(_this).f9);
          let _1209_v163;
          _1209_v163 = _module.D9.create_DC24(!(p2), p0, _module.__default.fm28(p1, p2, globalState));
          _1208_v162 = (_1208_v162).update((_dafny.Map.Empty.slice().updateUnsafe(p0,_1204_v160)).Merge(_1207_v161), _module.__default.fm1(_1138_v123, (_1209_v163).dtor_cf42, globalState));
          let _1210_v164;
          let _nw186 = new _module.C2();
          _nw186.__ctor(p1, !(p2));
          _1210_v164 = _nw186;
          _1210_v164 = ((p2) ? (_1210_v164) : (_1210_v164));
          let _1211_v165;
          _1211_v165 = _dafny.Seq.of(_1204_v160);
          (globalState).f2 = ((!(_1133_v117).equals(_1133_v117)) ? ((_this).f9) : (_dafny.Seq.IsProperPrefixOf(_1211_v165, _1211_v165)));
          r0 = new BigNumber((_1140_v124).length);
        } else {
          let _1212___mcc_h27 = (_source13).cf29;
          let _1213_cf29 = _1212___mcc_h27;
          r0 = p1;
          let _1214_v166;
          _1214_v166 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("mgsxcsdpe"));
          _1214_v166 = _1214_v166;
          let _1215_v167;
          let _init36 = ((_1216_p0) => function (_1217_i15) {
            return (_1217_i15).minus(_1216_p0);
          })(p0);
          let _nw187 = Array((new BigNumber(2)).toNumber());
          for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw187.length); _i0_36++) {
            _nw187[_i0_36] = _init36(new BigNumber(_i0_36));
          }
          _1215_v167 = _nw187;
          let _index179 = _module.__default.safeIndex(new BigNumber(168), new BigNumber((_1215_v167).length));
          (_1215_v167)[_index179] = new BigNumber(688);
          let _index180 = _module.__default.safeIndex(new BigNumber(168), new BigNumber((_1215_v167).length));
          (_1215_v167)[_index180] = p1;
          let _1218_v168;
          _1218_v168 = _dafny.Seq.UnicodeFromString("t");
          r0 = new BigNumber((_1218_v168).length);
        }
        if ((_this).f9) {
          _1138_v123 = _1138_v123;
          let _rhs207 = _1133_v117;
          let _rhs208 = (_dafny.ZERO).minus((_dafny.ZERO).minus(p1));
          let _rhs209 = (_module.D14.create_DC38(_1128_v115, p2, _1140_v124)).dtor_cf67;
          _1133_v117 = _rhs207;
          r0 = _rhs208;
          _1128_v115 = _rhs209;
          let _1219_v169;
          _1219_v169 = _dafny.Seq.UnicodeFromString("xqbsdfpq");
          r3 = !_dafny.Seq.contains(_1219_v169, _1128_v115);
          let _1220_v170;
          _1220_v170 = _dafny.Set.fromElements(_1129_v116);
          (globalState).f2 = ((!((_this).f9)) ? (true) : ((_1220_v170).equals(_1220_v170)));
          let _1221_v171;
          let _nw188 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Set.Empty);
          _1221_v171 = _nw188;
          let _1222_v172;
          let _nw189 = Array((new BigNumber(11)).toNumber());
          _nw189[(_dafny.ZERO).toNumber()] = p2;
          _nw189[(_dafny.ONE).toNumber()] = p2;
          _nw189[(new BigNumber(2)).toNumber()] = true;
          _nw189[(new BigNumber(3)).toNumber()] = _module.__default.fm1(_dafny.MultiSet.fromElements(p1), p0, globalState);
          _nw189[(new BigNumber(4)).toNumber()] = p2;
          _nw189[(new BigNumber(5)).toNumber()] = true;
          _nw189[(new BigNumber(6)).toNumber()] = p2;
          _nw189[(new BigNumber(7)).toNumber()] = _module.__default.fm1(_1138_v123, p0, globalState);
          _nw189[(new BigNumber(8)).toNumber()] = (_this).f9;
          _nw189[(new BigNumber(9)).toNumber()] = false;
          _nw189[(new BigNumber(10)).toNumber()] = _module.__default.fm1(_1138_v123, new BigNumber((_dafny.MultiSet.fromElements((_this).f9)).cardinality()), globalState);
          _1222_v172 = _nw189;
          let _index181 = _module.__default.safeIndex(new BigNumber(610), new BigNumber((_1221_v171).length));
          (_1221_v171)[_index181] = _dafny.Set.fromElements(_1222_v172);
          let _1223_v173;
          _1223_v173 = _dafny.Set.fromElements(_1222_v172, _1222_v172, _1222_v172);
          let _index182 = _module.__default.safeIndex(new BigNumber(610), new BigNumber((_1221_v171).length));
          (_1221_v171)[_index182] = _1223_v173;
        } else {
          (globalState).f2 = (_this).f9;
          let _1224_v174;
          _1224_v174 = _dafny.Seq.UnicodeFromString("ie");
          let _1225_v175;
          let _init37 = function (_1226_i17) {
            return _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_this).f9);
          };
          let _nw190 = Array((new BigNumber(4)).toNumber());
          for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw190.length); _i0_37++) {
            _nw190[_i0_37] = _init37(new BigNumber(_i0_37));
          }
          _1225_v175 = _nw190;
          let _1227_v176;
          _1227_v176 = _module.D7.create_DC19(_1225_v175, (_this).f9, _1224_v174);
          let _1228_v177;
          _1228_v177 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(153)), function (_1229_i16) {
            return new _dafny.CodePoint('c'.codePointAt(0));
          }), _1224_v174, (_1227_v176).dtor_cf35);
          let _rhs210 = _1228_v177;
          let _rhs211 = p0;
          let _rhs212 = _dafny.Seq.contains(_1140_v124, (_this).f9);
          let _rhs213 = _1128_v115;
          _1228_v177 = _rhs210;
          r0 = _rhs211;
          r3 = _rhs212;
          _1128_v115 = _rhs213;
          let _1230_v178;
          let _nw191 = Array((new BigNumber(24)).toNumber()).fill(_module.D3.Default());
          _1230_v178 = _nw191;
          let _pat_let_tv65 = _1140_v124;
          let _pat_let_tv66 = p1;
          let _pat_let_tv67 = _1140_v124;
          let _pat_let_tv68 = p2;
          let _index183 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_1230_v178).length));
          (_1230_v178)[_index183] = function (_pat_let48_0) {
            return function (_1231_dt__update__tmp_h10) {
              return function (_pat_let49_0) {
                return function (_1232_dt__update_hcf18_h0) {
                  return _module.D3.create_DC9((_1231_dt__update__tmp_h10).dtor_cf17, _1232_dt__update_hcf18_h0);
                }(_pat_let49_0);
              }(_dafny.Seq.update(_pat_let_tv65, _module.__default.safeIndex(_pat_let_tv66, new BigNumber((_pat_let_tv67).length)), _pat_let_tv68));
            }(_pat_let48_0);
          }(_module.D3.create_DC9((_this).f9, _1140_v124));
          let _1233_v179;
          _1233_v179 = _module.D3.create_DC9((_module.__default.fm0(globalState)).isLessThan(new BigNumber(657)), _dafny.Seq.of((_this).f9, p2));
          let _index184 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_1230_v178).length));
          (_1230_v178)[_index184] = _1233_v179;
          let _1234_v180;
          _1234_v180 = _dafny.MultiSet.fromElements(_1128_v115);
          r0 = new BigNumber((function () {
            let _coll25 = new _dafny.Set();
            for (const _compr_25 of ((_1234_v180).Union(_1234_v180)).Elements) {
              let _1235_v181 = _compr_25;
              if (((_1234_v180).Union(_1234_v180)).contains(_1235_v181)) {
                _coll25.add(_1235_v181);
              }
            }
            return _coll25;
          }()).length);
          let _1236_v182;
          let _init38 = ((_1237_v115, _1238_v117) => function (_1239_i18) {
            return _dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(846)), ((_1240_v115) => function (_1241_i19) {
              return _1240_v115;
            })(_1237_v115))).length), new BigNumber((_1238_v117).cardinality()));
          })(_1128_v115, _1133_v117);
          let _nw192 = Array((new BigNumber(12)).toNumber());
          for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw192.length); _i0_38++) {
            _nw192[_i0_38] = _init38(new BigNumber(_i0_38));
          }
          _1236_v182 = _nw192;
          let _1242_v183;
          _1242_v183 = _dafny.Seq.of(_1236_v182, _1236_v182);
          let _1243_v184;
          _1243_v184 = _dafny.Seq.of((_1242_v183)[_module.__default.safeIndex(_module.__default.fm0(globalState), new BigNumber((_1242_v183).length))]);
          let _1244_v185;
          let _nw193 = new _module.C3();
          _nw193.__ctor(true, (_1243_v184)[_module.__default.safeIndex(p0, new BigNumber((_1243_v184).length))], p2);
          _1244_v185 = _nw193;
          let _1245_v186;
          _1245_v186 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_1244_v185);
          let _1246_v187;
          _1246_v187 = _dafny.Seq.of(p0, p1, (_dafny.ZERO).minus(new BigNumber((_1245_v186).length)));
          _1246_v187 = _1246_v187;
        }
      }
      let _1247_v188;
      _1247_v188 = _dafny.Seq.UnicodeFromString("canqxbdox");
      let _1248_v189;
      _1248_v189 = _module.D1.create_DC2(p1, true, new BigNumber((_1247_v188).length), p0);
      let _1249_v190;
      let _nw194 = new _module.C1();
      _nw194.__ctor(_dafny.Seq.of(_1248_v189));
      _1249_v190 = _nw194;
      let _1250_v191;
      _1250_v191 = _dafny.Set.fromElements(_1249_v190, _1249_v190);
      let _hi3 = (_module.__default.fm0(globalState)).minus(p0);
      for (let _1251_i20 = new BigNumber((_1250_v191).length); _1251_i20.isLessThan(_hi3); _1251_i20 = _1251_i20.plus(_dafny.ONE)) {
        let _1252_v192;
        let _nw195 = new _module.C1();
        _nw195.__ctor(_dafny.Seq.Concat(_1249_v190.f13, _1249_v190.f13));
        _1252_v192 = _nw195;
        let _1253_v193;
        _1253_v193 = _dafny.Seq.of((_this).f9, !((_this).f9) || (!((_this).f9)), true);
        r0 = new BigNumber((_1253_v193).length);
        let _1254_v194;
        _1254_v194 = _dafny.Set.fromElements(_dafny.Seq.of(p0));
        let _1255_v195;
        _1255_v195 = _dafny.MultiSet.fromElements(p0, (_dafny.ZERO).minus((_dafny.ZERO).minus((_1251_i20).multipliedBy(new BigNumber((_1254_v194).length)))));
        _1255_v195 = _1255_v195;
        r0 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(218)), ((_1256_v188, _1257_p0) => function (_1258_i21) {
          return (_1256_v188)[_module.__default.safeIndex(_1257_p0, new BigNumber((_1256_v188).length))];
        })(_1247_v188, p0))).length);
      }
      let _1259_v196;
      _1259_v196 = new _dafny.CodePoint('b'.codePointAt(0));
      let _1260_v197;
      _1260_v197 = _module.D7.create_DC20((_this).f9, _1259_v196);
      let _1261_v198;
      _1261_v198 = _module.D7.create_DC21(_1260_v197);
      let _1262_v199;
      _1262_v199 = _module.D7.create_DC21(_1260_v197);
      let _1263_v200;
      _1263_v200 = _module.D7.create_DC21(_1261_v198);
      let _1264_v201;
      _1264_v201 = _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("lo")).length));
      let _rhs214 = ((_this).f9) === ((_this).f9);
      let _rhs215 = p2;
      let _rhs216 = _1263_v200;
      let _rhs217 = (p0).isEqualTo((((_this).f9) ? (p0) : (new BigNumber((_1264_v201).length))));
      let _rhs218 = _module.__default.safeModuloInt(p0, p0);
      let _lhs136 = globalState;
      let _lhs137 = globalState;
      let _lhs138 = globalState;
      _lhs136.f2 = _rhs214;
      _lhs137.f2 = _rhs215;
      _1263_v200 = _rhs216;
      _lhs138.f2 = _rhs217;
      r0 = _rhs218;
      let _1265_v202;
      let _nw196 = Array((new BigNumber(19)).toNumber()).fill(_module.D15.Default());
      _1265_v202 = _nw196;
      let _1266_v203;
      _1266_v203 = _dafny.Set.fromElements(p0, p1);
      let _1267_v204;
      _1267_v204 = _dafny.Seq.of(_dafny.Set.fromElements(p1));
      let _1268_v206;
      let _nw197 = Array((new BigNumber(15)).toNumber());
      _nw197[(_dafny.ZERO).toNumber()] = _1266_v203;
      _nw197[(_dafny.ONE).toNumber()] = _1266_v203;
      _nw197[(new BigNumber(2)).toNumber()] = _1266_v203;
      _nw197[(new BigNumber(3)).toNumber()] = _1266_v203;
      _nw197[(new BigNumber(4)).toNumber()] = _1266_v203;
      _nw197[(new BigNumber(5)).toNumber()] = _1266_v203;
      _nw197[(new BigNumber(6)).toNumber()] = _1266_v203;
      _nw197[(new BigNumber(7)).toNumber()] = (_1267_v204)[_module.__default.safeIndex(_module.__default.fm0(globalState), new BigNumber((_1267_v204).length))];
      _nw197[(new BigNumber(8)).toNumber()] = _dafny.Set.fromElements(p0, new BigNumber(-854));
      _nw197[(new BigNumber(9)).toNumber()] = _1266_v203;
      _nw197[(new BigNumber(10)).toNumber()] = _1266_v203;
      _nw197[(new BigNumber(11)).toNumber()] = _1266_v203;
      _nw197[(new BigNumber(12)).toNumber()] = _1266_v203;
      _nw197[(new BigNumber(13)).toNumber()] = _1266_v203;
      _nw197[(new BigNumber(14)).toNumber()] = function () {
        let _coll26 = new _dafny.Set();
        for (const _compr_26 of _dafny.IntegerRange(new BigNumber(96), new BigNumber(923))) {
          let _1269_v205 = _compr_26;
          if (((new BigNumber(96)).isLessThanOrEqualTo(_1269_v205)) && ((_1269_v205).isLessThan(new BigNumber(923)))) {
            _coll26.add(_module.__default.safeDivisionInt(_1269_v205, p1));
          }
        }
        return _coll26;
      }();
      _1268_v206 = _nw197;
      let _1270_v207;
      let _nw198 = new _module.C3();
      _nw198.__ctor((_this).f9, _1268_v206, _module.__default.fm1(_module.__default.fm29(!(p2), false, globalState), new BigNumber((_1247_v188).length), globalState));
      _1270_v207 = _nw198;
      let _1271_v208;
      _1271_v208 = _module.D15.create_DC42((_this).f9, p1, _1270_v207, new BigNumber((_1266_v203).length));
      let _index185 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_1265_v202).length));
      (_1265_v202)[_index185] = _1271_v208;
      let _index186 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_1265_v202).length));
      (_1265_v202)[_index186] = _1271_v208;
      (globalState).f2 = p2;
      let _1272_v209;
      _1272_v209 = _dafny.MultiSet.fromElements(p1);
      let _1273_v210;
      _1273_v210 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_1270_v207).f9);
      let _1274_v211;
      _1274_v211 = _module.D1.create_DC4(p1, (_1270_v207).f9, (_this).f9, p2, _module.__default.fm1(_1272_v209, new BigNumber((_1273_v210).length), globalState));
      r0 = (_dafny.ZERO).minus(((_1274_v211).dtor_cf9).minus(p1));
      let _1275_v212;
      _1275_v212 = _module.D4.create_DC13(p0);
      r1 = _1275_v212;
      let _1276_v213;
      _1276_v213 = _dafny.MultiSet.fromElements(_dafny.Seq.IsProperPrefixOf(_1247_v188, _dafny.Seq.update(_1247_v188, _module.__default.safeIndex(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements((_this).f9)).cardinality()))).length), new BigNumber((_1247_v188).length)), _1259_v196)));
      r2 = _1276_v213;
      r3 = p2;
      return [r0, r1, r2, r3];
    }
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f9 = false;
      this._f10 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T1, _module.T2, _module.T0];
    }
    get f9() {
      let _this = this;
      return _this._f9;
    };
    __ctor(f10, f9) {
      let _this = this;
      (_this)._f10 = f10;
      (_this)._f9 = f9;
      return;
    }
    fm2(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_dafny.Seq.of(new BigNumber((_dafny.Seq.of((_this).f9, (_this).f9)).length)));
    };
    fm3(globalState) {
      let _this = this;
      return (_this).f9;
    };
    fm4(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((_this).f10).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_this).f10, (_this).f10),(_dafny.ZERO).minus((_this).f10))).length));
    };
    fm5(p0, p1, p2, globalState) {
      let _this = this;
      return (_dafny.MultiSet.fromElements((_dafny.ZERO).minus((_this).f10), (_this).f10)).IsSubsetOf(_dafny.MultiSet.fromElements((_this).f10, new BigNumber(884)));
    };
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _1277_v0;
      _1277_v0 = _dafny.Seq.UnicodeFromString("m");
      let _1278_v1;
      _1278_v1 = _dafny.Seq.of(_1277_v0, _1277_v0);
      let _hi4 = new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_1278_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(379)), ((_1279_v0) => function (_1280_i1) {
        return _1279_v0;
      })(_1277_v0))))).cardinality());
      for (let _1281_i0 = (_this).f10; _1281_i0.isLessThan(_hi4); _1281_i0 = _1281_i0.plus(_dafny.ONE)) {
        let _1282_v2;
        _1282_v2 = new _dafny.CodePoint('x'.codePointAt(0));
        let _rhs219 = p2;
        let _rhs220 = new _dafny.CodePoint('c'.codePointAt(0));
        let _lhs139 = globalState;
        _lhs139.f2 = _rhs219;
        _1282_v2 = _rhs220;
        let _1283_v3;
        _1283_v3 = new BigNumber(504);
        let _1284_v4;
        let _nw199 = Array((new BigNumber(18)).toNumber()).fill([]);
        _1284_v4 = _nw199;
        let _1285_v5;
        let _nw200 = Array((new BigNumber(29)).toNumber()).fill(false);
        _1285_v5 = _nw200;
        let _index187 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((_1284_v4).length));
        (_1284_v4)[_index187] = _1285_v5;
        let _index188 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((_1284_v4).length));
        let _rhs221 = p2;
        let _rhs222 = p1;
        let _rhs223 = _module.__default.safeDivisionInt(((p2) ? (_1281_i0) : (_1281_i0)), new BigNumber((_1277_v0).length));
        let _rhs224 = _1285_v5;
        let _rhs225 = (_this).f9;
        let _lhs140 = globalState;
        let _lhs141 = globalState;
        let _lhs142 = _1284_v4;
        let _lhs143 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((_1284_v4).length));
        let _lhs144 = globalState;
        _lhs140.f2 = _rhs221;
        _lhs141.f2 = _rhs222;
        _1283_v3 = _rhs223;
        _lhs142[_lhs143] = _rhs224;
        _lhs144.f2 = _rhs225;
        let _1286_v6;
        _1286_v6 = _dafny.Seq.of(p1, p1);
        let _1287_v7;
        _1287_v7 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_1284_v4)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_1284_v4).length))]);
        let _index189 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((_1284_v4).length));
        let _rhs226 = (((_1287_v7).contains(p2)) ? ((_1287_v7).get(p2)) : ((_1284_v4)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_1284_v4).length))]));
        let _rhs227 = p0;
        let _rhs228 = _1286_v6;
        let _lhs145 = _1284_v4;
        let _lhs146 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((_1284_v4).length));
        _lhs145[_lhs146] = _rhs226;
        _1283_v3 = _rhs227;
        _1286_v6 = _rhs228;
        let _1288_v8;
        _1288_v8 = _dafny.MultiSet.fromElements(_1281_i0);
        let _1289_v9;
        _1289_v9 = _dafny.MultiSet.fromElements(p1, p1, p1);
        let _1290_v10;
        _1290_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_module.__default.fm1(_1288_v8, new BigNumber((_1289_v9).cardinality()), globalState));
        let _1291_v11;
        _1291_v11 = _dafny.MultiSet.fromElements(_1290_v10);
        if ((_dafny.MultiSet.fromElements(_1290_v10, _1290_v10)).IsProperSubsetOf(_1291_v11)) {
          let _1292_v12;
          let _init39 = ((_1293_i0) => function (_1294_i2) {
            return (_1294_i2).minus(_1293_i0);
          })(_1281_i0);
          let _nw201 = Array((new BigNumber(18)).toNumber());
          for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw201.length); _i0_39++) {
            _nw201[_i0_39] = _init39(new BigNumber(_i0_39));
          }
          _1292_v12 = _nw201;
          let _index190 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_1292_v12).length));
          (_1292_v12)[_index190] = p0;
          let _index191 = _module.__default.safeIndex(new BigNumber(852), new BigNumber((_1292_v12).length));
          (_1292_v12)[_index191] = _1281_i0;
          _1283_v3 = p0;
          (globalState).f2 = !((true) === ((_this).f9));
          _1290_v10 = (_1290_v10).update(p1, (_this).f9);
          _1290_v10 = (_1290_v10).update(!(p2), p2);
        } else {
          let _arr2 = (_1284_v4)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_1284_v4).length))];
          let _index192 = _module.__default.safeIndex(new BigNumber(931), new BigNumber(((_1284_v4)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_1284_v4).length))]).length));
          _arr2[_index192] = (_this).f9;
          let _1295_v13;
          let _init40 = ((_1296_p0) => function (_1297_i3) {
            return _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_1296_p0);
          })(p0);
          let _nw202 = Array((new BigNumber(9)).toNumber());
          for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw202.length); _i0_40++) {
            _nw202[_i0_40] = _init40(new BigNumber(_i0_40));
          }
          _1295_v13 = _nw202;
          let _1298_v14;
          _1298_v14 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.Concat(_module.__default.fm6(_dafny.Seq.of(p2), globalState), _1277_v0), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_module.__default.fm6(_dafny.Seq.of(p2), globalState), _1277_v0)).length)), _1282_v2),_1295_v13);
          let _arr3 = (_1284_v4)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_1284_v4).length))];
          let _index193 = _module.__default.safeIndex(new BigNumber(931), new BigNumber(((_1284_v4)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_1284_v4).length))]).length));
          let _rhs229 = _1282_v2;
          let _rhs230 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("srhyeyrp"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(574)), ((_1299_v2) => function (_1300_i4) {
            return _1299_v2;
          })(_1282_v2))), _1277_v0);
          let _rhs231 = (((_1298_v14).contains(_dafny.Seq.Concat(_1277_v0, _1277_v0))) ? ((_1298_v14).get(_dafny.Seq.Concat(_1277_v0, _1277_v0))) : (_1295_v13));
          let _lhs147 = (_1284_v4)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_1284_v4).length))];
          let _lhs148 = _module.__default.safeIndex(new BigNumber(931), new BigNumber(((_1284_v4)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_1284_v4).length))]).length));
          _1282_v2 = _rhs229;
          _lhs147[_lhs148] = _rhs230;
          _1295_v13 = _rhs231;
          let _1301_v15;
          _1301_v15 = _dafny.Map.Empty.slice().updateUnsafe((_1286_v6)[_module.__default.safeIndex(p0, new BigNumber((_1286_v6).length))],_1282_v2);
          _1301_v15 = (_1301_v15).update((_1288_v8).IsProperSubsetOf(_1288_v8), _1282_v2);
          let _1302_v16;
          _1302_v16 = _dafny.Map.Empty.slice().updateUnsafe(_1290_v10,_1281_i0);
          let _1303_v17;
          _1303_v17 = _dafny.Seq.of((_this).f10);
          _1283_v3 = (_dafny.ZERO).minus((((_this).f9) ? ((((_1302_v16).contains(_1290_v10)) ? ((_1302_v16).get(_1290_v10)) : ((_this).f10))) : (new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.update(_1303_v17, _module.__default.safeIndex(p0, new BigNumber((_1303_v17).length)), p0))).cardinality()))));
          let _1304_v18;
          _1304_v18 = _module.D1.create_DC1(_1283_v3);
          let _1305_v19;
          _1305_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1277_v0,p2);
          let _1306_v20;
          _1306_v20 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat(_1303_v17, (_this).fm2(new BigNumber((_dafny.Seq.of(new BigNumber(-758), new BigNumber(-196))).length), p2, p1, (_1304_v18).dtor_cf1, globalState))).length),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1305_v19).length),_module.__default.fm0(globalState)));
          let _1307_v21;
          _1307_v21 = _dafny.Set.fromElements((_1286_v6)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(((_1284_v4)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_1284_v4).length))])[_module.__default.safeIndex(new BigNumber(931), new BigNumber(((_1284_v4)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_1284_v4).length))]).length))],p2)).length), new BigNumber((_1286_v6).length))]);
          let _1308_v22;
          let _nw203 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
          _1308_v22 = _nw203;
          let _1309_v23;
          _1309_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1308_v22,_1281_i0);
          let _1310_v24;
          _1310_v24 = _dafny.Map.Empty.slice().updateUnsafe(_1307_v21,p1);
          let _1311_v25;
          _1311_v25 = _dafny.MultiSet.fromElements(_1310_v24);
          let _1312_v26;
          let _nw204 = Array((new BigNumber(27)).toNumber());
          _nw204[(_dafny.ZERO).toNumber()] = (((_1309_v23).contains(_1308_v22)) ? ((_1309_v23).get(_1308_v22)) : (new BigNumber(245)));
          _nw204[(_dafny.ONE).toNumber()] = new BigNumber((_1288_v8).cardinality());
          _nw204[(new BigNumber(2)).toNumber()] = p0;
          _nw204[(new BigNumber(3)).toNumber()] = new BigNumber((_1289_v9).cardinality());
          _nw204[(new BigNumber(4)).toNumber()] = _module.__default.safeDivisionInt(_1283_v3, new BigNumber(457));
          _nw204[(new BigNumber(5)).toNumber()] = _module.__default.fm0(globalState);
          _nw204[(new BigNumber(6)).toNumber()] = p0;
          _nw204[(new BigNumber(7)).toNumber()] = ((_this).f10).minus(_1281_i0);
          _nw204[(new BigNumber(8)).toNumber()] = _1281_i0;
          _nw204[(new BigNumber(9)).toNumber()] = (_1281_i0).plus(new BigNumber((_1277_v0).length));
          _nw204[(new BigNumber(10)).toNumber()] = (_this).f10;
          _nw204[(new BigNumber(11)).toNumber()] = (((_1302_v16).contains(_1290_v10)) ? ((_1302_v16).get(_1290_v10)) : (new BigNumber((_1277_v0).length)));
          _nw204[(new BigNumber(12)).toNumber()] = (_1281_i0).minus(new BigNumber((_1288_v8).cardinality()));
          _nw204[(new BigNumber(13)).toNumber()] = new BigNumber((_1277_v0).length);
          _nw204[(new BigNumber(14)).toNumber()] = (_this).f10;
          _nw204[(new BigNumber(15)).toNumber()] = new BigNumber((_1289_v9).cardinality());
          _nw204[(new BigNumber(16)).toNumber()] = p0;
          _nw204[(new BigNumber(17)).toNumber()] = (_1283_v3).multipliedBy((_this).f10);
          _nw204[(new BigNumber(18)).toNumber()] = (_1283_v3).minus(_1283_v3);
          _nw204[(new BigNumber(19)).toNumber()] = new BigNumber(845);
          _nw204[(new BigNumber(20)).toNumber()] = _module.__default.safeModuloInt(_1283_v3, _1281_i0);
          _nw204[(new BigNumber(21)).toNumber()] = (_dafny.ZERO).minus(p0);
          _nw204[(new BigNumber(22)).toNumber()] = _1283_v3;
          _nw204[(new BigNumber(23)).toNumber()] = (_this).f10;
          _nw204[(new BigNumber(24)).toNumber()] = (_1281_i0).plus(p0);
          _nw204[(new BigNumber(25)).toNumber()] = (new BigNumber((_1277_v0).length)).multipliedBy((_dafny.ZERO).minus((((_1311_v25).contains(_1310_v24)) ? ((_1311_v25).get(_1310_v24)) : (p0))));
          _nw204[(new BigNumber(26)).toNumber()] = ((p1) ? (new BigNumber((_1277_v0).length)) : (new BigNumber(678)));
          _1312_v26 = _nw204;
          let _1313_v27;
          _1313_v27 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-168)), ((_1314_v2) => function (_1315_i5) {
            return _1314_v2;
          })(_1282_v2))).length),((_1284_v4)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_1284_v4).length))])[_module.__default.safeIndex(new BigNumber(931), new BigNumber(((_1284_v4)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_1284_v4).length))]).length))]);
          let _index194 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_1312_v26).length));
          (_1312_v26)[_index194] = new BigNumber((((p1) ? (_1313_v27) : (_1313_v27))).length);
          let _index195 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_1312_v26).length));
          let _rhs232 = (_this).f10;
          let _rhs233 = new BigNumber(854);
          let _rhs234 = ((_1306_v20).Merge(_1306_v20)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1277_v0).length),function () {
            let _coll27 = new _dafny.Map();
            for (const _compr_27 of _dafny.IntegerRange(new BigNumber(644), new BigNumber(494))) {
              let _1316_v28 = _compr_27;
              if (((new BigNumber(644)).isLessThanOrEqualTo(_1316_v28)) && ((_1316_v28).isLessThan(new BigNumber(494)))) {
                _coll27.push([_module.__default.safeModuloInt(_1316_v28, p0),_1283_v3]);
              }
            }
            return _coll27;
          }()));
          let _rhs235 = _1307_v21;
          let _rhs236 = _1281_i0;
          let _lhs149 = _1312_v26;
          let _lhs150 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_1312_v26).length));
          _1283_v3 = _rhs232;
          _1283_v3 = _rhs233;
          _1306_v20 = _rhs234;
          _1307_v21 = _rhs235;
          _lhs149[_lhs150] = _rhs236;
          let _arr4 = (_1284_v4)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_1284_v4).length))];
          let _index196 = _module.__default.safeIndex(new BigNumber(931), new BigNumber(((_1284_v4)[_module.__default.safeIndex(new BigNumber(129), new BigNumber((_1284_v4).length))]).length));
          _arr4[_index196] = (_this).fm3(globalState);
        }
      }
      let _1317_v29;
      _1317_v29 = _dafny.MultiSet.fromElements(p0);
      let _rhs237 = ((((_this).f9) ? (_1317_v29) : (_1317_v29))).IsDisjointFrom(_1317_v29);
      let _rhs238 = _1277_v0;
      let _lhs151 = globalState;
      _lhs151.f2 = _rhs237;
      _1277_v0 = _rhs238;
      let _1318_v30;
      _1318_v30 = _module.D1.create_DC2((_this).f10, (_this).f9, p0, (_module.D1.create_DC4(new BigNumber(347), (_this).f9, p1, false, p2)).dtor_cf9);
      let _1319_v31;
      _1319_v31 = _dafny.Seq.of(p0, (_1318_v30).dtor_cf2);
      (globalState).f2 = (p0).isLessThan((new BigNumber((_1319_v31).length)).multipliedBy(p0));
      let _hi5 = (_this).f10;
      for (let _1320_i6 = p0; _1320_i6.isLessThan(_hi5); _1320_i6 = _1320_i6.plus(_dafny.ONE)) {
        let _1321_v32;
        let _1322_v33;
        let _1323_v34;
        let _out14;
        let _out15;
        let _out16;
        let _outcollector5 = (_this).m5(globalState);
        _out14 = _outcollector5[0];
        _out15 = _outcollector5[1];
        _out16 = _outcollector5[2];
        _1321_v32 = _out14;
        _1322_v33 = _out15;
        _1323_v34 = _out16;
        let _1324_v35;
        _1324_v35 = _dafny.Map.Empty.slice().updateUnsafe(_1321_v32,p1);
        let _1325_v36;
        _1325_v36 = _dafny.Seq.of((_this).f9, !(_1321_v32));
        let _1326_v37;
        _1326_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1324_v35,_1325_v36);
        let _1327_v38;
        _1327_v38 = _dafny.Set.fromElements(_1319_v31);
        let _1328_v39;
        _1328_v39 = _dafny.Map.Empty.slice().updateUnsafe(_1323_v34,_1325_v36);
        _1322_v33 = new BigNumber((_module.__default.fm7(new BigNumber((_1326_v37).length), p2, _module.D1.create_DC4(_1322_v33, p1, !((_this).fm5(_1327_v38, _1328_v39, new BigNumber((_1277_v0).length), globalState)), !(_1321_v32), false), (_1320_i6).isLessThan(new BigNumber((_1325_v36).length)), globalState)).length);
        let _1329_v40;
        _1329_v40 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_this).fm2(new BigNumber((_dafny.Seq.update(_1319_v31, _module.__default.safeIndex(_1323_v34, new BigNumber((_1319_v31).length)), new BigNumber(263))).length), (_this).f9, false, (_this).f10, globalState)).length),_module.D1.create_DC1(new BigNumber(794)));
        let _1330_v41;
        _1330_v41 = _module.D1.create_DC1(_1323_v34);
        _1329_v40 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_1320_i6, _1323_v34),_1330_v41);
        let _1331_v42;
        let _nw205 = new _module.C0();
        _nw205.__ctor(_1327_v38, _1322_v33);
        _1331_v42 = _nw205;
      }
      let _1332_v43;
      _1332_v43 = _module.D1.create_DC4(p0, (_this).f9, true, (_this).f9, (_this).f9);
      let _1333_v44;
      _1333_v44 = _dafny.Seq.of(_1332_v43, _1332_v43, _module.D1.create_DC4((_this).f10, !(p1), !(p1), p2, p2), _1332_v43, _1332_v43);
      let _1334_v45;
      _1334_v45 = _dafny.Seq.of(_1333_v44, _1333_v44, _1333_v44);
      _1333_v44 = _dafny.Seq.Concat((_1334_v45)[_module.__default.safeIndex((_this).f10, new BigNumber((_1334_v45).length))], (_1334_v45)[_module.__default.safeIndex((_this).f10, new BigNumber((_1334_v45).length))]);
      let _1335_v46;
      _1335_v46 = new BigNumber(38);
      let _1336_v47;
      _1336_v47 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(884)), ((_1337_v46) => function (_1338_i7) {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1337_v46,(_this).f9)).length);
      })(_1335_v46));
      let _1339_v48;
      _1339_v48 = _dafny.Set.fromElements(_1319_v31, _1319_v31);
      let _1340_v49;
      let _nw206 = new _module.C0();
      _nw206.__ctor(_1339_v48, _1335_v46);
      _1340_v49 = _nw206;
      let _1341_v50;
      _1341_v50 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_1340_v49);
      let _rhs239 = !(p0).isEqualTo((_dafny.ZERO).minus(_1335_v46));
      let _rhs240 = _module.__default.safeModuloInt(new BigNumber((_1319_v31).length), ((_this).f10).multipliedBy(new BigNumber((_1341_v50).length)));
      let _rhs241 = _module.__default.fm9(_module.__default.safeDivisionInt((_this).f10, p0), _1335_v46, false, globalState);
      let _rhs242 = _dafny.Seq.UnicodeFromString("yuoebdhit");
      let _lhs152 = globalState;
      _lhs152.f2 = _rhs239;
      _1335_v46 = _rhs240;
      _1336_v47 = _rhs241;
      _1277_v0 = _rhs242;
      return;
    }
    m2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let r3 = _dafny.Set.Empty;
      let _1342_i0;
      _1342_i0 = _dafny.ZERO;
      L8: {
        while ((_this).f9) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1342_i0)) {
              break L8;
            }
            _1342_i0 = (_1342_i0).plus(_dafny.ONE);
            let _1343_v0;
            _1343_v0 = _dafny.Set.fromElements(_dafny.Seq.update(p3, _module.__default.safeIndex((_this).f10, new BigNumber((p3).length)), (_this).f10));
            let _1344_v1;
            let _nw207 = new _module.C0();
            _nw207.__ctor(_1343_v0, ((p3)[_module.__default.safeIndex((_this).f10, new BigNumber((p3).length))]).minus((_this).f10));
            _1344_v1 = _nw207;
            let _1345_v2;
            let _nw208 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
            _1345_v2 = _nw208;
            let _index197 = _module.__default.safeIndex(new BigNumber(77), new BigNumber((_1345_v2).length));
            (_1345_v2)[_index197] = _module.__default.safeDivisionInt(_1344_v1.f12, _1344_v1.f12);
            let _index198 = _module.__default.safeIndex(new BigNumber(77), new BigNumber((_1345_v2).length));
            (_1345_v2)[_index198] = _1344_v1.f12;
            r1 = (_this).f9;
            let _index199 = _module.__default.safeIndex(new BigNumber(77), new BigNumber((_1345_v2).length));
            (_1345_v2)[_index199] = _module.__default.safeDivisionInt((_1345_v2)[_module.__default.safeIndex(new BigNumber(77), new BigNumber((_1345_v2).length))], (_1345_v2)[_module.__default.safeIndex(new BigNumber(77), new BigNumber((_1345_v2).length))]);
          }
        }
      }
      let _1346_v3;
      _1346_v3 = _dafny.Seq.UnicodeFromString("qgopc");
      let _1347_v4;
      _1347_v4 = _dafny.Set.fromElements((_this).f10);
      let _1348_v5;
      _1348_v5 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_1347_v4).IsDisjointFrom(_1347_v4));
      let _1349_v6;
      let _nw209 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
      _1349_v6 = _nw209;
      let _index200 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_1349_v6).length));
      (_1349_v6)[_index200] = ((_this).f10).minus((_this).f10);
      let _index201 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_1349_v6).length));
      let _rhs243 = p1;
      let _rhs244 = _1348_v5;
      let _rhs245 = (_dafny.Set.fromElements((_this).f10, new BigNumber(280), (_this).f10, (_this).f10)).contains(_module.__default.fm0(globalState));
      let _rhs246 = (_this).f10;
      let _rhs247 = !((_this).f9) || (((_this).f9) && (false));
      let _lhs153 = _1349_v6;
      let _lhs154 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_1349_v6).length));
      _1346_v3 = _rhs243;
      _1348_v5 = _rhs244;
      r1 = _rhs245;
      _lhs153[_lhs154] = _rhs246;
      r1 = _rhs247;
      if ((_this).f9) {
        let _1350_v7;
        _1350_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,true);
        let _1351_v8;
        _1351_v8 = _module.D1.create_DC3(!(_1350_v7).contains(_module.__default.fm0(globalState)), p1, ((_this).f10).multipliedBy((_this).f10));
        let _1352_v9;
        _1352_v9 = _dafny.Set.fromElements(p2, p2, p2, (((_this).f9) ? (p2) : (p2)));
        let _1353_v10;
        _1353_v10 = _dafny.MultiSet.fromElements((_1349_v6)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_1349_v6).length))]);
        let _1354_v11;
        let _nw210 = Array((new BigNumber(21)).toNumber());
        _nw210[(_dafny.ZERO).toNumber()] = (new BigNumber((_module.__default.fm10((_this).f9, globalState)).length)).isLessThanOrEqualTo((_this).f10);
        _nw210[(_dafny.ONE).toNumber()] = !((_this).f9);
        _nw210[(new BigNumber(2)).toNumber()] = (_this).f9;
        _nw210[(new BigNumber(3)).toNumber()] = (_this).f9;
        _nw210[(new BigNumber(4)).toNumber()] = (_this).f9;
        _nw210[(new BigNumber(5)).toNumber()] = (_1351_v8).dtor_cf6;
        _nw210[(new BigNumber(6)).toNumber()] = (_this).f9;
        _nw210[(new BigNumber(7)).toNumber()] = (_this).f9;
        _nw210[(new BigNumber(8)).toNumber()] = (_this).f9;
        _nw210[(new BigNumber(9)).toNumber()] = (_this).f9;
        _nw210[(new BigNumber(10)).toNumber()] = !_dafny.Seq.contains(_dafny.Seq.update(_1346_v3, _module.__default.safeIndex((_this).f10, new BigNumber((_1346_v3).length)), _module.__default.fm11((_this).f9, globalState)), p2);
        _nw210[(new BigNumber(11)).toNumber()] = (_this).f9;
        _nw210[(new BigNumber(12)).toNumber()] = !((_this).f9);
        _nw210[(new BigNumber(13)).toNumber()] = ((_this).f10).isLessThanOrEqualTo((_module.D1.create_DC1((_this).f10)).dtor_cf1);
        _nw210[(new BigNumber(14)).toNumber()] = (_this).f9;
        _nw210[(new BigNumber(15)).toNumber()] = (_this).f9;
        _nw210[(new BigNumber(16)).toNumber()] = (_this).f9;
        _nw210[(new BigNumber(17)).toNumber()] = _module.__default.fm1(_dafny.MultiSet.FromArray(p3), (_1349_v6)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_1349_v6).length))], globalState);
        _nw210[(new BigNumber(18)).toNumber()] = (_1353_v10).IsProperSubsetOf(_1353_v10);
        _nw210[(new BigNumber(19)).toNumber()] = (_this).f9;
        _nw210[(new BigNumber(20)).toNumber()] = (_this).f9;
        _1354_v11 = _nw210;
        let _index202 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_1354_v11).length));
        (_1354_v11)[_index202] = (_module.__default.fm12((_1351_v8).dtor_cf6, new BigNumber((_1346_v3).length), (_this).f9, new BigNumber((_1346_v3).length), globalState)).IsDisjointFrom(_module.__default.fm12(true, (_this).f10, (_this).f9, (_1349_v6)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_1349_v6).length))], globalState));
        let _1355_v12;
        _1355_v12 = _dafny.Seq.of((_this).f9);
        let _1356_v13;
        _1356_v13 = _module.D1.create_DC4(new BigNumber(144), (_this).f9, (_this).f9, (_this).f9, (_1355_v12)[_module.__default.safeIndex((_this).f10, new BigNumber((_1355_v12).length))]);
        let _pat_let_tv69 = _1349_v6;
        let _pat_let_tv70 = _1346_v3;
        let _index203 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_1354_v11).length));
        let _rhs248 = function (_pat_let50_0) {
          return function (_1357_dt__update__tmp_h0) {
            return function (_pat_let51_0) {
              return function (_1360_dt__update_hcf6_h0) {
                return _module.D1.create_DC3(_1360_dt__update_hcf6_h0, (_1357_dt__update__tmp_h0).dtor_cf7, (_1357_dt__update__tmp_h0).dtor_cf8);
              }(_pat_let51_0);
            }(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(68)), ((_1358_v6) => function (_1359_i1) {
              return (_dafny.ZERO).minus((_1358_v6)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_1358_v6).length))]);
            })(_pat_let_tv69)), _dafny.Seq.of(new BigNumber((_pat_let_tv70).length))));
          }(_pat_let50_0);
        }(_1351_v8);
        let _rhs249 = _1352_v9;
        let _rhs250 = (_this).f9;
        let _rhs251 = (_dafny.ZERO).minus((_1356_v13).dtor_cf9);
        let _rhs252 = (_this).f10;
        let _lhs155 = _1354_v11;
        let _lhs156 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_1354_v11).length));
        _1351_v8 = _rhs248;
        _1352_v9 = _rhs249;
        _lhs155[_lhs156] = _rhs250;
        r2 = _rhs251;
        r2 = _rhs252;
        if (((_1354_v11)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_1354_v11).length))]) || ((_this).f9)) {
          let _1361_v14;
          _1361_v14 = _dafny.Set.fromElements(p3);
          let _1362_v15;
          _1362_v15 = _dafny.Map.Empty.slice().updateUnsafe((_1349_v6)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_1349_v6).length))],_1355_v12);
          (globalState).f2 = (_this).fm5(_1361_v14, _1362_v15, ((_this).f10).multipliedBy((_this).f10), globalState);
          r0 = (_1354_v11)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_1354_v11).length))];
          let _1363_v16;
          _1363_v16 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_dafny.Seq.UnicodeFromString("gjurujc"));
          _1346_v3 = _dafny.Seq.Concat((((_1363_v16).contains((_dafny.ZERO).minus((_1349_v6)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_1349_v6).length))]))) ? ((_1363_v16).get((_dafny.ZERO).minus((_1349_v6)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_1349_v6).length))]))) : (p1)), p1);
          let _1364_v17;
          _1364_v17 = _dafny.MultiSet.fromElements(_dafny.Seq.IsPrefixOf(p1, _dafny.Seq.UnicodeFromString("vfyqtatu")));
          _1364_v17 = _1364_v17;
          _1353_v10 = _1353_v10;
        } else {
          let _index204 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_1349_v6).length));
          let _rhs253 = _module.__default.safeModuloInt((_this).f10, (_this).f10);
          let _rhs254 = (_this).f10;
          let _rhs255 = !((_this).f9);
          let _rhs256 = (_this).f10;
          let _lhs157 = globalState;
          let _lhs158 = _1349_v6;
          let _lhs159 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_1349_v6).length));
          r2 = _rhs253;
          r2 = _rhs254;
          _lhs157.f2 = _rhs255;
          _lhs158[_lhs159] = _rhs256;
          let _1365_v18;
          _1365_v18 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,p2);
          let _1366_v20;
          _1366_v20 = _dafny.Seq.of(function () {
            let _coll28 = new _dafny.Set();
            for (const _compr_28 of _dafny.IntegerRange(new BigNumber(545), new BigNumber(560))) {
              let _1367_v19 = _compr_28;
              if (((new BigNumber(545)).isLessThanOrEqualTo(_1367_v19)) && ((_1367_v19).isLessThan(new BigNumber(560)))) {
                _coll28.add((_1367_v19).plus((_this).f10));
              }
            }
            return _coll28;
          }(), (_1347_v4), _1347_v4);
          let _1368_v21;
          _1368_v21 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
          let _1369_v22;
          let _nw211 = Array((new BigNumber(8)).toNumber());
          _nw211[(_dafny.ZERO).toNumber()] = p2;
          _nw211[(_dafny.ONE).toNumber()] = p2;
          _nw211[(new BigNumber(2)).toNumber()] = p2;
          _nw211[(new BigNumber(3)).toNumber()] = (((_1365_v18).contains(new BigNumber((_1366_v20).length))) ? ((_1365_v18).get(new BigNumber((_1366_v20).length))) : (p2));
          _nw211[(new BigNumber(4)).toNumber()] = p2;
          _nw211[(new BigNumber(5)).toNumber()] = p2;
          _nw211[(new BigNumber(6)).toNumber()] = p2;
          _nw211[(new BigNumber(7)).toNumber()] = (((_1368_v21).contains(p2)) ? ((_1368_v21).get(p2)) : (p2));
          _1369_v22 = _nw211;
          let _index205 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_1369_v22).length));
          (_1369_v22)[_index205] = p2;
          let _1370_v23;
          _1370_v23 = _dafny.Seq.of(new BigNumber((_dafny.Seq.Concat(p1, p1)).length), ((_this).f10).multipliedBy((_1349_v6)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_1349_v6).length))]));
          let _1371_v24;
          let _nw212 = new _module.C0();
          _nw212.__ctor(_dafny.Set.fromElements(_1370_v23), (_this).f10);
          _1371_v24 = _nw212;
          let _1372_v25;
          _1372_v25 = _dafny.Seq.of(_1371_v24);
          let _index206 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_1369_v22).length));
          let _rhs257 = p2;
          let _rhs258 = _dafny.Seq.UnicodeFromString("lg");
          let _rhs259 = _dafny.Seq.UnicodeFromString("lf");
          let _rhs260 = _1370_v23;
          let _rhs261 = _1372_v25;
          let _lhs160 = _1369_v22;
          let _lhs161 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_1369_v22).length));
          _lhs160[_lhs161] = _rhs257;
          _1346_v3 = _rhs258;
          _1346_v3 = _rhs259;
          _1370_v23 = _rhs260;
          _1372_v25 = _rhs261;
          (globalState).f2 = (new BigNumber((p3).length)).isLessThanOrEqualTo((_this).f10);
          let _1373_v26;
          let _init41 = ((_1374_v6, _1375_v10) => function (_1376_i2) {
            return _dafny.Map.Empty.slice().updateUnsafe((_1374_v6)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_1374_v6).length))],_1375_v10);
          })(_1349_v6, _1353_v10);
          let _nw213 = Array((new BigNumber(8)).toNumber());
          for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw213.length); _i0_41++) {
            _nw213[_i0_41] = _init41(new BigNumber(_i0_41));
          }
          _1373_v26 = _nw213;
          let _1377_v27;
          _1377_v27 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1347_v4).length),_dafny.MultiSet.FromArray(p3));
          let _index207 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_1373_v26).length));
          (_1373_v26)[_index207] = (_dafny.Map.Empty.slice().updateUnsafe((_1349_v6)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_1349_v6).length))],_1353_v10)).Merge(_1377_v27);
          let _1378_v28;
          _1378_v28 = _module.D1.create_DC2((_this).f10, false, (_this).f10, (_dafny.ZERO).minus((_this).f10));
          let _1379_v29;
          _1379_v29 = _dafny.Seq.of(_1378_v28, _1378_v28);
          let _1380_v30;
          let _nw214 = new _module.C1();
          _nw214.__ctor(_1379_v29);
          _1380_v30 = _nw214;
          let _1381_v31;
          _1381_v31 = _dafny.Map.Empty.slice().updateUnsafe(true,_1380_v30);
          let _1382_v32;
          _1382_v32 = _module.D1.create_DC2((_1349_v6)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_1349_v6).length))], (_this).f9, new BigNumber((_1381_v31).length), (_this).f10);
          let _index208 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_1373_v26).length));
          (_1373_v26)[_index208] = (_1377_v27).Merge((_module.__default.fm13((_this).fm4((_1354_v11)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_1354_v11).length))], (_this).f9, (_this).f9, (_1382_v32).dtor_cf3, globalState), (_1354_v11)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_1354_v11).length))], (_this).f9, globalState)).Merge(_1377_v27));
          let _1383_v33;
          let _nw215 = new _module.C0();
          _nw215.__ctor(_1371_v24.f11, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(839)), ((_1384_p2) => function (_1385_i3) {
            return _1384_p2;
          })(p2))).length));
          _1383_v33 = _nw215;
        }
        let _1386_v34;
        let _nw216 = Array((new BigNumber(2)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1386_v34 = _nw216;
        let _1387_v35;
        _1387_v35 = _dafny.Seq.of(_1386_v34);
        let _index209 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_1354_v11).length));
        (_1354_v11)[_index209] = _dafny.Seq.contains(_1387_v35, _1386_v34);
        _module.__default.m0((_this).f9, globalState);
        let _1388_v36;
        _1388_v36 = _module.D1.create_DC2((_this).f10, (_this).f9, (_1349_v6)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_1349_v6).length))], (_this).f10);
        let _1389_v37;
        _1389_v37 = _dafny.Seq.of(_1388_v36, _1388_v36);
        let _1390_v38;
        let _nw217 = new _module.C1();
        _nw217.__ctor(_1389_v37);
        _1390_v38 = _nw217;
      } else {
        let _1391_v39;
        let _nw218 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _1391_v39 = _nw218;
        let _index210 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_1391_v39).length));
        (_1391_v39)[_index210] = _1346_v3;
        let _index211 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_1391_v39).length));
        (_1391_v39)[_index211] = _1346_v3;
        let _1392_v40;
        let _nw219 = Array((new BigNumber(21)).toNumber());
        _1392_v40 = _nw219;
        let _1393_v41;
        _1393_v41 = _module.D4.create_DC11(_1392_v40);
        let _1394_v42;
        let _nw220 = Array((new BigNumber(19)).toNumber());
        _nw220[(_dafny.ZERO).toNumber()] = _1392_v40;
        _nw220[(_dafny.ONE).toNumber()] = _1392_v40;
        _nw220[(new BigNumber(2)).toNumber()] = _1392_v40;
        _nw220[(new BigNumber(3)).toNumber()] = _1392_v40;
        _nw220[(new BigNumber(4)).toNumber()] = _1392_v40;
        _nw220[(new BigNumber(5)).toNumber()] = _1392_v40;
        _nw220[(new BigNumber(6)).toNumber()] = _1392_v40;
        _nw220[(new BigNumber(7)).toNumber()] = _1392_v40;
        _nw220[(new BigNumber(8)).toNumber()] = _1392_v40;
        _nw220[(new BigNumber(9)).toNumber()] = _1392_v40;
        _nw220[(new BigNumber(10)).toNumber()] = _1392_v40;
        _nw220[(new BigNumber(11)).toNumber()] = _1392_v40;
        _nw220[(new BigNumber(12)).toNumber()] = _1392_v40;
        _nw220[(new BigNumber(13)).toNumber()] = _1392_v40;
        _nw220[(new BigNumber(14)).toNumber()] = _1392_v40;
        _nw220[(new BigNumber(15)).toNumber()] = _1392_v40;
        _nw220[(new BigNumber(16)).toNumber()] = _1392_v40;
        _nw220[(new BigNumber(17)).toNumber()] = (_1393_v41).dtor_cf19;
        _nw220[(new BigNumber(18)).toNumber()] = _1392_v40;
        _1394_v42 = _nw220;
        let _1395_v43;
        _1395_v43 = _1394_v42;
        _1394_v42 = (_1395_v43);
        let _index212 = _module.__default.safeIndex(new BigNumber(716), new BigNumber((_1349_v6).length));
        (_1349_v6)[_index212] = (_this).f10;
        let _1396_v44;
        _1396_v44 = _module.D1.create_DC2((_this).f10, (_this).f9, (_1349_v6)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_1349_v6).length))], new BigNumber(247));
        let _1397_v45;
        let _nw221 = new _module.C1();
        _nw221.__ctor(_dafny.Seq.of(_1396_v44, _1396_v44));
        _1397_v45 = _nw221;
        let _1398_v46;
        _1398_v46 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_dafny.ZERO).minus((_1349_v6)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_1349_v6).length))]));
        _1398_v46 = (_1398_v46).Merge(_1398_v46);
      }
      r2 = ((_1349_v6)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_1349_v6).length))]).minus((_1349_v6)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_1349_v6).length))]);
      let _1399_v47;
      let _nw222 = Array((new BigNumber(23)).toNumber()).fill(false);
      _1399_v47 = _nw222;
      let _index213 = _module.__default.safeIndex(new BigNumber(305), new BigNumber((_1399_v47).length));
      (_1399_v47)[_index213] = (_this).f9;
      let _1400_v49;
      _1400_v49 = _dafny.Seq.of(_1399_v47);
      let _1401_v50;
      _1401_v50 = _dafny.MultiSet.fromElements(new BigNumber((p1).length));
      let _index214 = _module.__default.safeIndex(new BigNumber(305), new BigNumber((_1399_v47).length));
      let _rhs262 = !((_module.__default.fm12((_this).f9, (_1349_v6)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_1349_v6).length))], (_this).f9, (_this).f10, globalState)).IsSubsetOf(function () {
        let _coll29 = new _dafny.Set();
        for (const _compr_29 of (_1347_v4).Elements) {
          let _1402_v48 = _compr_29;
          if ((_1347_v4).contains(_1402_v48)) {
            _coll29.add((_1402_v48).multipliedBy(new BigNumber(78)));
          }
        }
        return _coll29;
      }()));
      let _rhs263 = _1399_v47;
      let _rhs264 = (_1400_v49)[_module.__default.safeIndex((_this).f10, new BigNumber((_1400_v49).length))];
      let _rhs265 = _module.__default.fm1(_1401_v50, (_1349_v6)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_1349_v6).length))], globalState);
      let _lhs162 = _1399_v47;
      let _lhs163 = _module.__default.safeIndex(new BigNumber(305), new BigNumber((_1399_v47).length));
      _lhs162[_lhs163] = _rhs262;
      _1399_v47 = _rhs263;
      _1399_v47 = _rhs264;
      r1 = _rhs265;
      let _1403_v51;
      _1403_v51 = _module.D3.create_DC10();
      _1403_v51 = (((_this).f9) ? (_1403_v51) : (_1403_v51));
      r0 = (_1399_v47)[_module.__default.safeIndex(new BigNumber(305), new BigNumber((_1399_v47).length))];
      let _1404_v52;
      _1404_v52 = _dafny.Seq.of(p3);
      let _1405_v54;
      _1405_v54 = _dafny.Seq.of((_this).f9);
      let _1406_v55;
      _1406_v55 = _module.D3.create_DC9((_this).f9, _1405_v54);
      let _1407_v56;
      _1407_v56 = _dafny.Map.Empty.slice().updateUnsafe((_1349_v6)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_1349_v6).length))],_dafny.Seq.of(new BigNumber(((_1406_v55).dtor_cf18).length), (_1349_v6)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_1349_v6).length))], (_this).f10));
      let _1408_v57;
      _1408_v57 = _dafny.Set.fromElements((((_1407_v56).contains((_this).f10)) ? ((_1407_v56).get((_this).f10)) : (p3)));
      let _1409_v58;
      _1409_v58 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_1405_v54);
      r1 = (_this).fm5((function () {
        let _coll30 = new _dafny.Set();
        for (const _compr_30 of (_1404_v52).Elements) {
          let _1410_v53 = _compr_30;
          if (_dafny.Seq.contains(_1404_v52, _1410_v53)) {
            _coll30.add(_1410_v53);
          }
        }
        return _coll30;
      }()).Union(_1408_v57), (_1409_v58).Merge(_1409_v58), (_1349_v6)[_module.__default.safeIndex(new BigNumber(716), new BigNumber((_1349_v6).length))], globalState);
      r2 = (_this).f10;
      r3 = (_1347_v4).Intersect(_1347_v4);
      return [r0, r1, r2, r3];
    }
    m3(p0, globalState) {
      let _this = this;
      let r0 = false;
      let _hi6 = (_this).f10;
      for (let _1411_i0 = (_this).f10; _1411_i0.isLessThan(_hi6); _1411_i0 = _1411_i0.plus(_dafny.ONE)) {
        if (p0) {
          let _1412_v0;
          _1412_v0 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-810)), ((_1413_i0) => function (_1414_i1) {
            return _1413_i0;
          })(_1411_i0)));
          let _1415_v1;
          let _nw223 = new _module.C0();
          _nw223.__ctor(_1412_v0, ((_this).f10).plus(_1411_i0));
          _1415_v1 = _nw223;
          let _1416_v2;
          _1416_v2 = _dafny.MultiSet.fromElements((_this).f10, (_this).f10);
          let _1417_v3;
          _1417_v3 = _module.D1.create_DC4((_dafny.ZERO).minus((_dafny.ZERO).minus(((_this).f10).multipliedBy(_1415_v1.f12))), !((_1416_v2).IsProperSubsetOf(_1416_v2)), (_this).f9, (_this).f9, !(p0));
          _1417_v3 = _module.D1.create_DC4(_1411_i0, (_this).f9, p0, (_this).f9, p0);
          let _1418_v4;
          _1418_v4 = _dafny.Seq.of((_this).f10, new BigNumber(267));
          let _1419_v5;
          _1419_v5 = _dafny.Seq.UnicodeFromString("fdwey");
          let _1420_v6;
          _1420_v6 = _module.D1.create_DC1(new BigNumber((_1419_v5).length));
          let _pat_let_tv71 = globalState;
          let _1421_v7;
          let _nw224 = Array((new BigNumber(27)).toNumber());
          _nw224[(_dafny.ZERO).toNumber()] = _module.__default.fm18(new BigNumber((_1418_v4).length), p0, p0, globalState);
          _nw224[(_dafny.ONE).toNumber()] = _1420_v6;
          _nw224[(new BigNumber(2)).toNumber()] = _1420_v6;
          _nw224[(new BigNumber(3)).toNumber()] = _module.__default.fm18(_1411_i0, p0, p0, globalState);
          _nw224[(new BigNumber(4)).toNumber()] = ((p0) ? (_1420_v6) : (_1420_v6));
          _nw224[(new BigNumber(5)).toNumber()] = _module.D1.create_DC1((_dafny.ZERO).minus(_1411_i0));
          _nw224[(new BigNumber(6)).toNumber()] = _1420_v6;
          _nw224[(new BigNumber(7)).toNumber()] = _1420_v6;
          _nw224[(new BigNumber(8)).toNumber()] = _1420_v6;
          _nw224[(new BigNumber(9)).toNumber()] = _1420_v6;
          _nw224[(new BigNumber(10)).toNumber()] = _1420_v6;
          _nw224[(new BigNumber(11)).toNumber()] = function (_pat_let52_0) {
            return function (_1422_dt__update__tmp_h0) {
              return function (_pat_let53_0) {
                return function (_1423_dt__update_hcf1_h0) {
                  return _module.D1.create_DC1(_1423_dt__update_hcf1_h0);
                }(_pat_let53_0);
              }(_1411_i0);
            }(_pat_let52_0);
          }(_1420_v6);
          _nw224[(new BigNumber(12)).toNumber()] = _1420_v6;
          _nw224[(new BigNumber(13)).toNumber()] = _module.D1.create_DC1(_1411_i0);
          _nw224[(new BigNumber(14)).toNumber()] = ((p0) ? (_1420_v6) : (_1420_v6));
          _nw224[(new BigNumber(15)).toNumber()] = _1420_v6;
          _nw224[(new BigNumber(16)).toNumber()] = _1420_v6;
          _nw224[(new BigNumber(17)).toNumber()] = _1420_v6;
          _nw224[(new BigNumber(18)).toNumber()] = _1420_v6;
          _nw224[(new BigNumber(19)).toNumber()] = _1420_v6;
          _nw224[(new BigNumber(20)).toNumber()] = _1420_v6;
          _nw224[(new BigNumber(21)).toNumber()] = (((_this).f9) ? (_1420_v6) : (_1420_v6));
          _nw224[(new BigNumber(22)).toNumber()] = _1420_v6;
          _nw224[(new BigNumber(23)).toNumber()] = _1420_v6;
          _nw224[(new BigNumber(24)).toNumber()] = function (_pat_let54_0) {
            return function (_1424_dt__update__tmp_h1) {
              return function (_pat_let55_0) {
                return function (_1425_dt__update_hcf1_h1) {
                  return _module.D1.create_DC1(_1425_dt__update_hcf1_h1);
                }(_pat_let55_0);
              }(_module.__default.fm0(_pat_let_tv71));
            }(_pat_let54_0);
          }(_1420_v6);
          _nw224[(new BigNumber(25)).toNumber()] = _1420_v6;
          _nw224[(new BigNumber(26)).toNumber()] = _1420_v6;
          _1421_v7 = _nw224;
          let _index215 = _module.__default.safeIndex(new BigNumber(74), new BigNumber((_1421_v7).length));
          (_1421_v7)[_index215] = _1420_v6;
          let _pat_let_tv72 = _1415_v1;
          let _pat_let_tv73 = _1415_v1;
          let _index216 = _module.__default.safeIndex(new BigNumber(74), new BigNumber((_1421_v7).length));
          (_1421_v7)[_index216] = function (_pat_let56_0) {
            return function (_1426_dt__update__tmp_h2) {
              return function (_pat_let57_0) {
                return function (_1427_dt__update_hcf1_h2) {
                  return _module.D1.create_DC1(_1427_dt__update_hcf1_h2);
                }(_pat_let57_0);
              }(_module.__default.safeDivisionInt(_pat_let_tv72.f12, _pat_let_tv73.f12));
            }(_pat_let56_0);
          }(_1420_v6);
          let _1428_v8;
          _1428_v8 = _dafny.Seq.of((_this).f9);
          (_1415_v1).f12 = (_dafny.ZERO).minus(((_1411_i0).multipliedBy(new BigNumber((_1428_v8).length))).plus(_1411_i0));
          let _1429_v9;
          let _nw225 = new _module.C1();
          _nw225.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(468)), ((_1430_p0, _1431_i0, _1432_v4) => function (_1433_i2) {
            return _module.D1.create_DC2((_this).f10, _1430_p0, _1431_i0, new BigNumber((_1432_v4).length));
          })(p0, _1411_i0, _1418_v4)));
          _1429_v9 = _nw225;
          let _1434_v10;
          _1434_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1415_v1.f12,_1429_v9);
          _1429_v9 = (((_1434_v10).contains((_this).f10)) ? ((_1434_v10).get((_this).f10)) : (_1429_v9));
        } else {
          let _1435_v11;
          _1435_v11 = _dafny.Seq.UnicodeFromString("l");
          let _1436_v12;
          _1436_v12 = _module.__default.fm14(_1411_i0, new BigNumber((_1435_v11).length), (_this).f10, globalState);
          _1436_v12 = (((_this).f9) ? (_1436_v12) : (_1436_v12));
          let _1437_v13;
          _1437_v13 = _dafny.MultiSet.fromElements((_this).f9);
          let _1438_v14;
          _1438_v14 = _dafny.Seq.of(_1435_v11);
          let _1439_v15;
          _1439_v15 = _dafny.Seq.of((_this).f9, (_this).f9);
          let _1440_v16;
          _1440_v16 = _dafny.Seq.of(_1411_i0, _1411_i0, new BigNumber((_1439_v15).length), (_this).f10);
          let _1441_v17;
          _1441_v17 = _dafny.Seq.of(_module.__default.fm0(globalState), _1411_i0, new BigNumber((_1440_v16).length));
          let _1442_v18;
          _1442_v18 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1435_v11).length),_1441_v17);
          let _1443_v19;
          _1443_v19 = _dafny.MultiSet.fromElements(_1411_i0);
          let _1444_v20;
          _1444_v20 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1435_v11).length),_1411_i0);
          let _1445_v21;
          _1445_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_dafny.MultiSet.fromElements((_this).f10, new BigNumber(756)));
          let _1446_v22;
          _1446_v22 = _dafny.Set.fromElements(_1443_v19, _dafny.MultiSet.fromElements((_this).f10, new BigNumber((_1444_v20).length)), (((_1445_v21).contains(true)) ? ((_1445_v21).get(true)) : (_1443_v19)));
          let _1447_v23;
          let _nw226 = Array((new BigNumber(21)).toNumber());
          _nw226[(_dafny.ZERO).toNumber()] = (_this).fm4(p0, p0, (_this).f9, p0, globalState);
          _nw226[(_dafny.ONE).toNumber()] = new BigNumber(((_1437_v13).update(!((_this).f9), _module.__default.abs((_this).f10))).cardinality());
          _nw226[(new BigNumber(2)).toNumber()] = (_this).f10;
          _nw226[(new BigNumber(3)).toNumber()] = _module.__default.safeModuloInt(_1411_i0, (_this).f10);
          _nw226[(new BigNumber(4)).toNumber()] = (new BigNumber((_1438_v14).length)).plus((_dafny.ZERO).minus((_this).f10));
          _nw226[(new BigNumber(5)).toNumber()] = (_this).f10;
          _nw226[(new BigNumber(6)).toNumber()] = _1411_i0;
          _nw226[(new BigNumber(7)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(p0))).cardinality()), new BigNumber((_1442_v18).length));
          _nw226[(new BigNumber(8)).toNumber()] = (_this).f10;
          _nw226[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_1441_v17)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_1441_v17).length))], (_this).f10));
          _nw226[(new BigNumber(10)).toNumber()] = (_this).f10;
          _nw226[(new BigNumber(11)).toNumber()] = _1411_i0;
          _nw226[(new BigNumber(12)).toNumber()] = new BigNumber((_1439_v15).length);
          _nw226[(new BigNumber(13)).toNumber()] = (_this).fm4((_this).f9, p0, !(false), !(p0), globalState);
          _nw226[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(807), new BigNumber(9)));
          _nw226[(new BigNumber(15)).toNumber()] = (_this).f10;
          _nw226[(new BigNumber(16)).toNumber()] = (_this).f10;
          _nw226[(new BigNumber(17)).toNumber()] = (_this).f10;
          _nw226[(new BigNumber(18)).toNumber()] = new BigNumber((_1446_v22).length);
          _nw226[(new BigNumber(19)).toNumber()] = _1411_i0;
          _nw226[(new BigNumber(20)).toNumber()] = _1411_i0;
          _1447_v23 = _nw226;
          let _index217 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_1447_v23).length));
          (_1447_v23)[_index217] = _1411_i0;
          let _index218 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_1447_v23).length));
          (_1447_v23)[_index218] = (_dafny.ZERO).minus((_this).f10);
          let _1448_v24;
          let _nw227 = Array((new BigNumber(19)).toNumber()).fill(false);
          _1448_v24 = _nw227;
          let _1449_v25;
          _1449_v25 = _dafny.Seq.of(_1448_v24);
          _1448_v24 = (_1449_v25)[_module.__default.safeIndex((_1447_v23)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_1447_v23).length))], new BigNumber((_1449_v25).length))];
          _1435_v11 = _1435_v11;
          let _1450_v27;
          _1450_v27 = _1443_v19;
          _1444_v20 = function () {
            let _coll31 = new _dafny.Map();
            for (const _compr_31 of ((_1450_v27)).Elements) {
              let _1451_v26 = _compr_31;
              if (((_1450_v27)).contains(_1451_v26)) {
                _coll31.push([(_1451_v26).minus(_1411_i0),(_1447_v23)[_module.__default.safeIndex(new BigNumber(405), new BigNumber((_1447_v23).length))]]);
              }
            }
            return _coll31;
          }();
        }
        let _1452_v28;
        let _nw228 = Array((new BigNumber(7)).toNumber());
        _nw228[(_dafny.ZERO).toNumber()] = (_1411_i0).multipliedBy(_1411_i0);
        _nw228[(_dafny.ONE).toNumber()] = new BigNumber(584);
        _nw228[(new BigNumber(2)).toNumber()] = (_this).f10;
        _nw228[(new BigNumber(3)).toNumber()] = _1411_i0;
        _nw228[(new BigNumber(4)).toNumber()] = (_this).f10;
        _nw228[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus((_this).f10);
        _nw228[(new BigNumber(6)).toNumber()] = ((_this).f10).multipliedBy(_1411_i0);
        _1452_v28 = _nw228;
        let _index219 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_1452_v28).length));
        (_1452_v28)[_index219] = _1411_i0;
        let _1453_v29;
        let _nw229 = new _module.C2();
        _nw229.__ctor(_1411_i0, p0);
        _1453_v29 = _nw229;
        let _1454_v30;
        _1454_v30 = _dafny.Set.fromElements(((p0) ? (_1453_v29) : (_1453_v29)), _1453_v29, _1453_v29, _1453_v29, _1453_v29);
        let _1455_v31;
        _1455_v31 = new BigNumber(485);
        let _index220 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_1452_v28).length));
        let _rhs266 = ((_1455_v31).plus((_this).f10)).plus((_1455_v31).plus(new BigNumber(114)));
        let _rhs267 = ((true) ? ((_1454_v30).Union(_dafny.Set.fromElements(_1453_v29))) : (_1454_v30));
        let _rhs268 = _module.__default.safeDivisionInt((_this).f10, (_this).f10);
        let _lhs164 = _1452_v28;
        let _lhs165 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_1452_v28).length));
        _lhs164[_lhs165] = _rhs266;
        _1454_v30 = _rhs267;
        _1455_v31 = _rhs268;
        let _1456_v32;
        _1456_v32 = _dafny.Seq.UnicodeFromString("jk");
        let _1457_v33;
        _1457_v33 = new _dafny.CodePoint('q'.codePointAt(0));
        let _1458_v34;
        _1458_v34 = _dafny.MultiSet.fromElements(_1457_v33, _1457_v33);
        let _1459_v35;
        _1459_v35 = _dafny.Seq.of(new BigNumber((_1458_v34).cardinality()), _1455_v31);
        let _1460_v36;
        _1460_v36 = _dafny.Set.fromElements(_1459_v35);
        let _1461_v38;
        let _nw230 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Set.Empty);
        _1461_v38 = _nw230;
        let _1462_v39;
        let _nw231 = new _module.C3();
        _nw231.__ctor(p0, _1461_v38, (_this).f9);
        _1462_v39 = _nw231;
        let _1463_v40;
        _1463_v40 = _dafny.Seq.of(_1462_v39);
        let _1464_v41;
        _1464_v41 = _dafny.Seq.of(_1463_v40);
        let _1465_v42;
        _1465_v42 = _dafny.MultiSet.fromElements(new BigNumber((_1464_v41).length), (_this).f10);
        let _1466_v43;
        _1466_v43 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(96)), function (_1467_i3) {
          return new _dafny.CodePoint('s'.codePointAt(0));
        }), _dafny.Seq.update(_dafny.Seq.UnicodeFromString("x"), _module.__default.safeIndex((_this).f10, new BigNumber((_dafny.Seq.UnicodeFromString("x")).length)), _1457_v33), (((_this).fm5(_1460_v36, function () {
          let _coll32 = new _dafny.Map();
          for (const _compr_32 of (_1465_v42).Elements) {
            let _1468_v37 = _compr_32;
            if ((_1465_v42).contains(_1468_v37)) {
              _coll32.push([(_1468_v37).multipliedBy(new BigNumber(173)),_dafny.Seq.of((_1462_v39).f9)]);
            }
          }
          return _coll32;
        }(), (_dafny.ZERO).minus(new BigNumber((_1456_v32).length)), globalState)) ? (_1456_v32) : (_1456_v32)));
        let _1469_v44;
        _1469_v44 = _dafny.Set.fromElements(!((_this).f9), (_1462_v39).f9, (_1453_v29).f9);
        let _1470_v45;
        _1470_v45 = _dafny.Map.Empty.slice().updateUnsafe((_1462_v39).f9,_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("goklh"), _1457_v33));
        let _1471_v46;
        _1471_v46 = _dafny.Seq.of((_1453_v29).f9);
        let _rhs269 = (_1466_v43)[_module.__default.safeIndex((_1459_v35)[_module.__default.safeIndex(new BigNumber((_1469_v44).length), new BigNumber((_1459_v35).length))], new BigNumber((_1466_v43).length))];
        let _rhs270 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-945)), ((_1472_v33) => function (_1473_i4) {
          return _1472_v33;
        })(_1457_v33));
        let _rhs271 = !((((_1470_v45).contains((_1471_v46)[_module.__default.safeIndex((_1452_v28)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_1452_v28).length))], new BigNumber((_1471_v46).length))])) ? ((_1470_v45).get((_1471_v46)[_module.__default.safeIndex((_1452_v28)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_1452_v28).length))], new BigNumber((_1471_v46).length))])) : ((_this).f9)));
        let _rhs272 = (_1462_v39).f9;
        let _lhs166 = globalState;
        let _lhs167 = globalState;
        _1456_v32 = _rhs269;
        _1456_v32 = _rhs270;
        _lhs166.f2 = _rhs271;
        _lhs167.f2 = _rhs272;
        let _1474_v47;
        _1474_v47 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f10);
        let _1475_v48;
        _1475_v48 = _module.D4.create_DC12(!((_this).f9), (_1453_v29).f9, new BigNumber((_1474_v47).length), (_this).f10, !((_1453_v29).f9));
        let _1476_v49;
        _1476_v49 = _module.D4.create_DC15(_1475_v48);
        let _source14 = _1476_v49;
        if (_source14.is_DC12) {
          let _1477___mcc_h0 = (_source14).cf20;
          let _1478___mcc_h1 = (_source14).cf21;
          let _1479___mcc_h2 = (_source14).cf22;
          let _1480___mcc_h3 = (_source14).cf23;
          let _1481___mcc_h4 = (_source14).cf24;
          let _1482_cf24 = _1481___mcc_h4;
          let _1483_cf23 = _1480___mcc_h3;
          let _1484_cf22 = _1479___mcc_h2;
          let _1485_cf21 = _1478___mcc_h1;
          let _1486_cf20 = _1477___mcc_h0;
          let _1487_v50;
          let _init42 = function (_1488_i5) {
            return false;
          };
          let _nw232 = Array((new BigNumber(24)).toNumber());
          for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw232.length); _i0_42++) {
            _nw232[_i0_42] = _init42(new BigNumber(_i0_42));
          }
          _1487_v50 = _nw232;
          let _index221 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_1487_v50).length));
          (_1487_v50)[_index221] = (_1453_v29).f9;
          let _index222 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_1487_v50).length));
          (_1487_v50)[_index222] = (_1462_v39).f9;
          _1483_cf23 = (_1459_v35)[_module.__default.safeIndex((_1452_v28)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_1452_v28).length))], new BigNumber((_1459_v35).length))];
          let _index223 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_1452_v28).length));
          (_1452_v28)[_index223] = (_1459_v35)[_module.__default.safeIndex(_1411_i0, new BigNumber((_1459_v35).length))];
          let _1489_v51;
          let _nw233 = new _module.C2();
          _nw233.__ctor((_dafny.ZERO).minus(new BigNumber((_1469_v44).length)), (_1487_v50)[_module.__default.safeIndex(new BigNumber(949), new BigNumber((_1487_v50).length))]);
          _1489_v51 = _nw233;
        } else if (_source14.is_DC13) {
          let _1490___mcc_h5 = (_source14).cf25;
          let _1491_cf25 = _1490___mcc_h5;
          let _1492_v52;
          _1492_v52 = _dafny.MultiSet.fromElements((_1471_v46)[_module.__default.safeIndex((_1452_v28)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_1452_v28).length))], new BigNumber((_1471_v46).length))]);
          _1492_v52 = ((!((_1462_v39).f9)) ? (_dafny.MultiSet.FromArray(_1471_v46)) : ((_1492_v52).update(p0, _module.__default.abs((_1452_v28)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_1452_v28).length))]))));
          let _1493_v53;
          let _nw234 = Array((new BigNumber(11)).toNumber()).fill([]);
          _1493_v53 = _nw234;
          let _1494_v54;
          _1494_v54 = _module.D14.create_DC38(_1457_v33, (_1453_v29).f9, _1471_v46);
          let _1495_v55;
          let _nw235 = Array((new BigNumber(15)).toNumber());
          _nw235[(_dafny.ZERO).toNumber()] = !(p0);
          _nw235[(_dafny.ONE).toNumber()] = (_1453_v29).f9;
          _nw235[(new BigNumber(2)).toNumber()] = p0;
          _nw235[(new BigNumber(3)).toNumber()] = (_this).f9;
          _nw235[(new BigNumber(4)).toNumber()] = (_1453_v29).f9;
          _nw235[(new BigNumber(5)).toNumber()] = (_this).f9;
          _nw235[(new BigNumber(6)).toNumber()] = (_1494_v54).dtor_cf68;
          _nw235[(new BigNumber(7)).toNumber()] = (_1453_v29).f9;
          _nw235[(new BigNumber(8)).toNumber()] = (_1462_v39).f9;
          _nw235[(new BigNumber(9)).toNumber()] = !((_1453_v29).f9);
          _nw235[(new BigNumber(10)).toNumber()] = (_1462_v39).f9;
          _nw235[(new BigNumber(11)).toNumber()] = (_1462_v39).f9;
          _nw235[(new BigNumber(12)).toNumber()] = (_1453_v29).f9;
          _nw235[(new BigNumber(13)).toNumber()] = (_1462_v39).f9;
          _nw235[(new BigNumber(14)).toNumber()] = (_this).f9;
          _1495_v55 = _nw235;
          let _index224 = _module.__default.safeIndex(new BigNumber(616), new BigNumber((_1493_v53).length));
          (_1493_v53)[_index224] = _1495_v55;
          let _index225 = _module.__default.safeIndex(new BigNumber(616), new BigNumber((_1493_v53).length));
          (_1493_v53)[_index225] = _1495_v55;
          let _1496_v56;
          let _nw236 = new _module.C2();
          _nw236.__ctor((_this).f10, (((_1453_v29).f9) ? ((_1462_v39).f9) : (p0)));
          _1496_v56 = _nw236;
          let _1497_v57;
          _1497_v57 = _dafny.Seq.of(_1496_v56, _1496_v56, _1496_v56, _1496_v56);
          _1496_v56 = (_1497_v57)[_module.__default.safeIndex(_1455_v31, new BigNumber((_1497_v57).length))];
          let _1498_v58;
          _1498_v58 = _dafny.MultiSet.fromElements(_1469_v44);
          let _1499_v59;
          _1499_v59 = _module.D3.create_DC9(true, _1471_v46);
          _1498_v58 = (_module.__default.fm43(_dafny.Set.fromElements(new BigNumber(((_1499_v59).dtor_cf18).length), _1411_i0), new BigNumber((_1474_v47).length), _module.__default.fm0(globalState), globalState)).update(_1469_v44, _module.__default.abs(_1491_cf25));
        } else if (_source14.is_DC14) {
          let _1500___mcc_h6 = (_source14).cf26;
          let _1501___mcc_h7 = (_source14).cf27;
          let _1502___mcc_h8 = (_source14).cf28;
          let _1503_cf28 = _1502___mcc_h8;
          let _1504_cf27 = _1501___mcc_h7;
          let _1505_cf26 = _1500___mcc_h6;
          let _1506_v60;
          let _init43 = ((_1507_cf28, _1508_v33, _1509_v32) => function (_1510_i6) {
            return _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("hksmc"), _module.__default.safeIndex(_1507_cf28, new BigNumber((_dafny.Seq.UnicodeFromString("hksmc")).length)), _1508_v33), _1509_v32);
          })(_1503_cf28, _1457_v33, _1456_v32);
          let _nw237 = Array((new BigNumber(9)).toNumber());
          for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw237.length); _i0_43++) {
            _nw237[_i0_43] = _init43(new BigNumber(_i0_43));
          }
          _1506_v60 = _nw237;
          let _index226 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_1452_v28).length));
          let _rhs273 = _1459_v35;
          let _rhs274 = _module.__default.fm0(globalState);
          let _rhs275 = _1506_v60;
          let _lhs168 = _1452_v28;
          let _lhs169 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_1452_v28).length));
          _1459_v35 = _rhs273;
          _lhs168[_lhs169] = _rhs274;
          _1506_v60 = _rhs275;
          let _1511_v61;
          _1511_v61 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_1455_v31),(_1453_v29).f9);
          let _1512_v62;
          _1512_v62 = _dafny.Set.fromElements((_this).f10, (_this).f10);
          let _1513_v63;
          _1513_v63 = _1512_v62;
          let _1514_v64;
          _1514_v64 = _dafny.Map.Empty.slice().updateUnsafe(_1513_v63,_1503_cf28);
          let _1515_v65;
          _1515_v65 = _dafny.Set.fromElements(_1514_v64);
          let _rhs276 = ((false) ? (p0) : ((((_1511_v61).contains((_1452_v28)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_1452_v28).length))])) ? ((_1511_v61).get((_1452_v28)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_1452_v28).length))])) : (_module.__default.fm1(_dafny.MultiSet.fromElements((_dafny.ZERO).minus(_1503_cf28), (_this).f10), _1503_cf28, globalState)))));
          let _rhs277 = !((((!((_1462_v39).f9)) ? (_1515_v65) : (_1515_v65))).IsProperSubsetOf(_1515_v65));
          let _lhs170 = globalState;
          r0 = _rhs276;
          _lhs170.f2 = _rhs277;
          _1456_v32 = _1456_v32;
          let _1516_v66;
          _1516_v66 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_1455_v31),(_1453_v29).f9);
          let _1517_v67;
          _1517_v67 = _module.D7.create_DC18(_1516_v66);
          let _pat_let_tv74 = _1516_v66;
          let _1518_v68;
          let _nw238 = Array((new BigNumber(9)).toNumber());
          _nw238[(_dafny.ZERO).toNumber()] = _1517_v67;
          _nw238[(_dafny.ONE).toNumber()] = _1517_v67;
          _nw238[(new BigNumber(2)).toNumber()] = _1517_v67;
          _nw238[(new BigNumber(3)).toNumber()] = _1517_v67;
          _nw238[(new BigNumber(4)).toNumber()] = _module.D7.create_DC18(_1516_v66);
          _nw238[(new BigNumber(5)).toNumber()] = _1517_v67;
          _nw238[(new BigNumber(6)).toNumber()] = _1517_v67;
          _nw238[(new BigNumber(7)).toNumber()] = function (_pat_let58_0) {
            return function (_1519_dt__update__tmp_h3) {
              return function (_pat_let59_0) {
                return function (_1520_dt__update_hcf32_h0) {
                  return _module.D7.create_DC18(_1520_dt__update_hcf32_h0);
                }(_pat_let59_0);
              }(_pat_let_tv74);
            }(_pat_let58_0);
          }(_1517_v67);
          _nw238[(new BigNumber(8)).toNumber()] = _module.D7.create_DC18(_1516_v66);
          _1518_v68 = _nw238;
          let _index227 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_1518_v68).length));
          (_1518_v68)[_index227] = _1517_v67;
          let _index228 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_1518_v68).length));
          (_1518_v68)[_index228] = _1517_v67;
        } else if (_source14.is_DC11) {
          let _1521___mcc_h9 = (_source14).cf19;
          let _1522_cf19 = _1521___mcc_h9;
          let _1523_v69;
          let _nw239 = new _module.C4();
          _nw239.__ctor((_1453_v29).f9);
          _1523_v69 = _nw239;
          let _1524_v70;
          _1524_v70 = _dafny.MultiSet.fromElements(_module.__default.fm1(_module.__default.fm29((_this).f9, (_1462_v39).f9, globalState), _1411_i0, globalState), (_1462_v39).f9, (_1453_v29).f9);
          _1455_v31 = (_dafny.ZERO).minus((((_1524_v70).contains((_1469_v44).IsDisjointFrom(_1469_v44))) ? ((_1524_v70).get((_1469_v44).IsDisjointFrom(_1469_v44))) : ((_dafny.ZERO).minus(_1411_i0))));
          _1457_v33 = _1457_v33;
          _1455_v31 = _module.__default.fm0(globalState);
        } else {
          let _1525___mcc_h10 = (_source14).cf29;
          let _1526_cf29 = _1525___mcc_h10;
          let _1527_v71;
          let _nw240 = new _module.C2();
          _nw240.__ctor(new BigNumber((_dafny.Seq.UnicodeFromString("xf")).length), (_this).f9);
          _1527_v71 = _nw240;
          let _1528_v72;
          _1528_v72 = _dafny.Seq.of(_dafny.Seq.of((_1462_v39).f9, true));
          let _index229 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_1452_v28).length));
          (_1452_v28)[_index229] = new BigNumber((_1528_v72).length);
          let _index230 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_1452_v28).length));
          (_1452_v28)[_index230] = (new BigNumber((_1474_v47).length)).plus((_dafny.ZERO).minus((_1527_v71).f14));
          let _index231 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_1452_v28).length));
          (_1452_v28)[_index231] = _1411_i0;
        }
      }
      let _1529_v73;
      _1529_v73 = _module.D1.create_DC4((_this).f10, !(p0), p0, (_this).f9, false);
      let _1530_v74;
      let _nw241 = new _module.C1();
      _nw241.__ctor(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(207)), function (_1531_i7) {
        return _module.D1.create_DC2(new BigNumber(828), (_this).f9, (_this).f10, (_this).f10);
      }), _module.__default.fm7((_this).f10, (_this).f9, _1529_v73, true, globalState)));
      _1530_v74 = _nw241;
      _1530_v74 = _1530_v74;
      let _1532_v75;
      let _nw242 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Set.Empty);
      _1532_v75 = _nw242;
      let _1533_v76;
      let _nw243 = new _module.C3();
      _nw243.__ctor(p0, _1532_v75, true);
      _1533_v76 = _nw243;
      let _1534_v78;
      _1534_v78 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,p0);
      let _1535_v79;
      _1535_v79 = _dafny.Seq.UnicodeFromString("urpoc");
      let _1536_v81;
      _1536_v81 = _dafny.Seq.of(_1534_v78, _1534_v78);
      let _1537_v82;
      let _nw244 = Array((new BigNumber(27)).toNumber());
      _nw244[(_dafny.ZERO).toNumber()] = (function () {
        let _coll33 = new _dafny.Map();
        for (const _compr_33 of _dafny.IntegerRange(new BigNumber(294), new BigNumber(229))) {
          let _1538_v77 = _compr_33;
          if (((new BigNumber(294)).isLessThanOrEqualTo(_1538_v77)) && ((_1538_v77).isLessThan(new BigNumber(229)))) {
            _coll33.push([_module.__default.safeModuloInt(_1538_v77, (_this).f10),p0]);
          }
        }
        return _coll33;
      }()).update((_dafny.ZERO).minus((_this).f10), (_1533_v76).f15);
      _nw244[(_dafny.ONE).toNumber()] = (_module.__default.fm15((_1533_v76).f15, globalState)).Merge(_1534_v78);
      _nw244[(new BigNumber(2)).toNumber()] = _1534_v78;
      _nw244[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_1533_v76).f15);
      _nw244[(new BigNumber(4)).toNumber()] = _1534_v78;
      _nw244[(new BigNumber(5)).toNumber()] = _1534_v78;
      _nw244[(new BigNumber(6)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,p0);
      _nw244[(new BigNumber(7)).toNumber()] = _1534_v78;
      _nw244[(new BigNumber(8)).toNumber()] = _1534_v78;
      _nw244[(new BigNumber(9)).toNumber()] = ((_1534_v78).update((_this).f10, !((_1533_v76).f15))).Merge(_1534_v78);
      _nw244[(new BigNumber(10)).toNumber()] = _1534_v78;
      _nw244[(new BigNumber(11)).toNumber()] = _1534_v78;
      _nw244[(new BigNumber(12)).toNumber()] = _1534_v78;
      _nw244[(new BigNumber(13)).toNumber()] = _1534_v78;
      _nw244[(new BigNumber(14)).toNumber()] = _1534_v78;
      _nw244[(new BigNumber(15)).toNumber()] = _1534_v78;
      _nw244[(new BigNumber(16)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1535_v79).length),p0);
      _nw244[(new BigNumber(17)).toNumber()] = _1534_v78;
      _nw244[(new BigNumber(18)).toNumber()] = _1534_v78;
      _nw244[(new BigNumber(19)).toNumber()] = (_1534_v78).Merge(_1534_v78);
      _nw244[(new BigNumber(20)).toNumber()] = _module.__default.fm15((_this).f9, globalState);
      _nw244[(new BigNumber(21)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,p0);
      _nw244[(new BigNumber(22)).toNumber()] = _1534_v78;
      _nw244[(new BigNumber(23)).toNumber()] = function () {
        let _coll34 = new _dafny.Map();
        for (const _compr_34 of _dafny.IntegerRange(new BigNumber(670), new BigNumber(-668))) {
          let _1539_v80 = _compr_34;
          if (((new BigNumber(670)).isLessThanOrEqualTo(_1539_v80)) && ((_1539_v80).isLessThan(new BigNumber(-668)))) {
            _coll34.push([(_1539_v80).multipliedBy((_this).f10),(_this).f9]);
          }
        }
        return _coll34;
      }();
      _nw244[(new BigNumber(24)).toNumber()] = (_module.__default.fm15(p0, globalState)).Merge((_1536_v81)[_module.__default.safeIndex((_this).f10, new BigNumber((_1536_v81).length))]);
      _nw244[(new BigNumber(25)).toNumber()] = _1534_v78;
      _nw244[(new BigNumber(26)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_this).f10,p0)).Merge(_1534_v78);
      _1537_v82 = _nw244;
      let _index232 = _module.__default.safeIndex(new BigNumber(713), new BigNumber((_1537_v82).length));
      (_1537_v82)[_index232] = _1534_v78;
      let _1540_v83;
      _1540_v83 = _module.D1.create_DC2((_this).f10, true, (_this).f10, (_this).f10);
      let _pat_let_tv75 = _1534_v78;
      let _pat_let_tv76 = _1534_v78;
      let _pat_let_tv77 = _1536_v81;
      let _pat_let_tv78 = _1536_v81;
      let _pat_let_tv79 = _1534_v78;
      let _pat_let_tv80 = _1534_v78;
      let _pat_let_tv81 = _1534_v78;
      let _index233 = _module.__default.safeIndex(new BigNumber(713), new BigNumber((_1537_v82).length));
      (_1537_v82)[_index233] = function (_source15) {
        if (_source15.is_DC2) {
          let _1541___mcc_h11 = (_source15).cf2;
          let _1542___mcc_h12 = (_source15).cf3;
          let _1543___mcc_h13 = (_source15).cf4;
          let _1544___mcc_h14 = (_source15).cf5;
          let _1545_cf5 = _1544___mcc_h14;
          let _1546_cf4 = _1543___mcc_h13;
          let _1547_cf3 = _1542___mcc_h12;
          let _1548_cf2 = _1541___mcc_h11;
          return (_pat_let_tv75).Merge(_pat_let_tv76);
        } else if (_source15.is_DC3) {
          let _1549___mcc_h15 = (_source15).cf6;
          let _1550___mcc_h16 = (_source15).cf7;
          let _1551___mcc_h17 = (_source15).cf8;
          let _1552_cf8 = _1551___mcc_h17;
          let _1553_cf7 = _1550___mcc_h16;
          let _1554_cf6 = _1549___mcc_h15;
          return (_pat_let_tv77)[_module.__default.safeIndex((_this).f10, new BigNumber((_pat_let_tv78).length))];
        } else if (_source15.is_DC4) {
          let _1555___mcc_h18 = (_source15).cf9;
          let _1556___mcc_h19 = (_source15).cf10;
          let _1557___mcc_h20 = (_source15).cf11;
          let _1558___mcc_h21 = (_source15).cf12;
          let _1559___mcc_h22 = (_source15).cf13;
          let _1560_cf13 = _1559___mcc_h22;
          let _1561_cf12 = _1558___mcc_h21;
          let _1562_cf11 = _1557___mcc_h20;
          let _1563_cf10 = _1556___mcc_h19;
          let _1564_cf9 = _1555___mcc_h18;
          return _pat_let_tv79;
        } else if (_source15.is_DC1) {
          let _1565___mcc_h23 = (_source15).cf1;
          let _1566_cf1 = _1565___mcc_h23;
          return _pat_let_tv80;
        } else {
          let _1567___mcc_h24 = (_source15).cf14;
          let _1568_cf14 = _1567___mcc_h24;
          return _pat_let_tv81;
        }
      }(_1540_v83);
      let _1569_v84;
      let _init44 = function (_1570_i8) {
        return (_this).f9;
      };
      let _nw245 = Array((new BigNumber(22)).toNumber());
      for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw245.length); _i0_44++) {
        _nw245[_i0_44] = _init44(new BigNumber(_i0_44));
      }
      _1569_v84 = _nw245;
      let _index234 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_1569_v84).length));
      (_1569_v84)[_index234] = true;
      let _1571_v85;
      _1571_v85 = _dafny.Seq.of(p0);
      let _1572_v86;
      _1572_v86 = _dafny.MultiSet.fromElements(p0, (_1533_v76).f15);
      let _1573_v87;
      _1573_v87 = _dafny.MultiSet.fromElements((_this).f10, new BigNumber((_1572_v86).cardinality()), new BigNumber(771), (_this).f10, (_this).f10);
      let _index235 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_1569_v84).length));
      (_1569_v84)[_index235] = ((true) ? ((_1571_v85)[_module.__default.safeIndex((((_1573_v87).contains((_this).f10)) ? ((_1573_v87).get((_this).f10)) : ((_this).f10)), new BigNumber((_1571_v85).length))]) : (false));
      if (((_this).f10).isLessThanOrEqualTo((_this).f10)) {
        let _1574_v88;
        _1574_v88 = _dafny.Seq.of(_dafny.Set.fromElements(new BigNumber((_1535_v79).length), (_this).f10));
        let _1575_v89;
        _1575_v89 = _dafny.Map.Empty.slice().updateUnsafe((_1533_v76).fm32(_1574_v88, !((_this).f9), (_this).f10, false, globalState),_1535_v79);
        let _1576_v90;
        _1576_v90 = _dafny.Seq.of((_this).f10, new BigNumber(668));
        let _1577_v91;
        _1577_v91 = new _dafny.CodePoint('e'.codePointAt(0));
        _1575_v89 = (_1575_v89).update(_dafny.Seq.IsPrefixOf(_1535_v79, _1535_v79), _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("cgiaxc"), _module.__default.safeIndex((_1576_v90)[_module.__default.safeIndex((_this).f10, new BigNumber((_1576_v90).length))], new BigNumber((_dafny.Seq.UnicodeFromString("cgiaxc")).length)), _1577_v91), _dafny.Seq.UnicodeFromString("ehohxsxbk")));
        let _1578_v92;
        _1578_v92 = new BigNumber(256);
        _1578_v92 = (_this).f10;
        let _index236 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_1569_v84).length));
        (_1569_v84)[_index236] = (_this).f9;
        _1578_v92 = _1578_v92;
        let _1579_v93;
        let _init45 = ((_1580_v79) => function (_1581_i9) {
          return _1580_v79;
        })(_1535_v79);
        let _nw246 = Array((new BigNumber(27)).toNumber());
        for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw246.length); _i0_45++) {
          _nw246[_i0_45] = _init45(new BigNumber(_i0_45));
        }
        _1579_v93 = _nw246;
        let _index237 = _module.__default.safeIndex(new BigNumber(599), new BigNumber((_1579_v93).length));
        (_1579_v93)[_index237] = _dafny.Seq.update(_dafny.Seq.update(_1535_v79, _module.__default.safeIndex(_1578_v92, new BigNumber((_1535_v79).length)), (_module.D13.create_DC35((_1533_v76).f15, _1577_v91, _1530_v74, (_1533_v76).f15)).dtor_cf59), _module.__default.safeIndex((_this).f10, new BigNumber((_dafny.Seq.update(_1535_v79, _module.__default.safeIndex(_1578_v92, new BigNumber((_1535_v79).length)), (_module.D13.create_DC35((_1533_v76).f15, _1577_v91, _1530_v74, (_1533_v76).f15)).dtor_cf59)).length)), _1577_v91);
        let _index238 = _module.__default.safeIndex(new BigNumber(599), new BigNumber((_1579_v93).length));
        (_1579_v93)[_index238] = _dafny.Seq.Concat(_1535_v79, _dafny.Seq.update(_1535_v79, _module.__default.safeIndex((_this).f10, new BigNumber((_1535_v79).length)), _1577_v91));
      } else {
        let _1582_v94;
        let _nw247 = new _module.C4();
        _nw247.__ctor(true);
        _1582_v94 = _nw247;
        let _1583_v95;
        let _nw248 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Seq.of());
        _1583_v95 = _nw248;
        (_1533_v76).m1((_dafny.ZERO).minus(new BigNumber(-225)), p0, p0, _1583_v95, globalState);
        let _1584_v96;
        _1584_v96 = _dafny.Seq.of(_module.D4.create_DC13(new BigNumber(923)));
        let _1585_v97;
        _1585_v97 = _module.D4.create_DC15((_1584_v96)[_module.__default.safeIndex((_this).f10, new BigNumber((_1584_v96).length))]);
        let _1586_v98;
        _1586_v98 = _dafny.Set.fromElements(_module.D4.create_DC15(_module.D4.create_DC13(new BigNumber(863))), _1585_v97, _1585_v97);
        _1586_v98 = _1586_v98;
        let _1587_v99;
        _1587_v99 = _dafny.Seq.of((_this).f10);
        let _1588_v100;
        _1588_v100 = _dafny.Set.fromElements(_1587_v99);
        let _1589_v101;
        let _nw249 = new _module.C0();
        _nw249.__ctor((_1588_v100).Intersect(_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(905)), function (_1590_i10) {
          return new BigNumber(-421);
        }), _dafny.Seq.of((_this).f10))), (_this).f10);
        _1589_v101 = _nw249;
        let _source16 = _module.__default.fm44(globalState);
        if (_source16.is_DC19) {
          let _1591___mcc_h25 = (_source16).cf33;
          let _1592___mcc_h26 = (_source16).cf34;
          let _1593___mcc_h27 = (_source16).cf35;
          let _1594_cf35 = _1593___mcc_h27;
          let _1595_cf34 = _1592___mcc_h26;
          let _1596_cf33 = _1591___mcc_h25;
          (_1589_v101).f12 = _module.__default.safeDivisionInt(((_1595_cf34) ? ((_this).f10) : (new BigNumber((_1571_v85).length))), _1589_v101.f12);
          let _1597_v102;
          let _nw250 = new _module.C3();
          _nw250.__ctor(true, _1532_v75, (_1533_v76).f15);
          _1597_v102 = _nw250;
          let _1598_v103;
          _1598_v103 = _dafny.Seq.of(_1569_v84, _1569_v84, _1569_v84, _1569_v84, _1569_v84);
          _1569_v84 = (_1598_v103)[_module.__default.safeIndex(_1589_v101.f12, new BigNumber((_1598_v103).length))];
          let _1599_v104;
          _1599_v104 = _dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC1(new BigNumber((_dafny.Seq.of(p0)).length)),(_1572_v86).update(false, _module.__default.abs((_dafny.ZERO).minus(_1589_v101.f12))));
          let _pat_let_tv82 = _1589_v101;
          _1599_v104 = (_1599_v104).update(function (_pat_let60_0) {
            return function (_1600_dt__update__tmp_h4) {
              return function (_pat_let61_0) {
                return function (_1601_dt__update_hcf1_h3) {
                  return _module.D1.create_DC1(_1601_dt__update_hcf1_h3);
                }(_pat_let61_0);
              }(_pat_let_tv82.f12);
            }(_pat_let60_0);
          }(_module.D1.create_DC1((((_1572_v86).contains((_1533_v76).f15)) ? ((_1572_v86).get((_1533_v76).f15)) : (_1589_v101.f12)))), _1572_v86);
        } else if (_source16.is_DC20) {
          let _1602___mcc_h28 = (_source16).cf36;
          let _1603___mcc_h29 = (_source16).cf37;
          let _1604_cf37 = _1603___mcc_h29;
          let _1605_cf36 = _1602___mcc_h28;
          let _1606_v105;
          _1606_v105 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(globalState),_1589_v101.f12);
          _1605_cf36 = !(!(_1606_v105).contains((_this).f10));
          let _1607_v106;
          _1607_v106 = _dafny.Map.Empty.slice().updateUnsafe(false,(_dafny.ZERO).minus((new BigNumber((_1587_v99).length)).minus(_1589_v101.f12)));
          _1607_v106 = (_1607_v106).update(!(!((_1569_v84)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_1569_v84).length))])), new BigNumber(414));
          let _1608_v107;
          _1608_v107 = _dafny.Map.Empty.slice().updateUnsafe((_1533_v76).f15,_1605_cf36);
          let _1609_v108;
          _1609_v108 = _dafny.Set.fromElements(p0, p0, p0);
          let _index239 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_1569_v84).length));
          let _rhs278 = (_1589_v101.f12).minus((_this).f10);
          let _rhs279 = (_1572_v86).IsSubsetOf(_dafny.MultiSet.fromElements(_1605_cf36));
          let _rhs280 = _1572_v86;
          let _rhs281 = (((_1608_v107).contains((_dafny.Set.fromElements((_1533_v76).f15)).IsDisjointFrom(_1609_v108))) ? ((_1608_v107).get((_dafny.Set.fromElements((_1533_v76).f15)).IsDisjointFrom(_1609_v108))) : (false));
          let _lhs171 = _1589_v101;
          let _lhs172 = globalState;
          let _lhs173 = _1569_v84;
          let _lhs174 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_1569_v84).length));
          _lhs171.f12 = _rhs278;
          _lhs172.f2 = _rhs279;
          _1572_v86 = _rhs280;
          _lhs173[_lhs174] = _rhs281;
          (_1589_v101).f12 = _module.__default.fm0(globalState);
        } else if (_source16.is_DC18) {
          let _1610___mcc_h30 = (_source16).cf32;
          let _1611_cf32 = _1610___mcc_h30;
          let _1612_v109;
          _1612_v109 = _module.D15.create_DC40(_dafny.Set.fromElements((_1533_v76).f15));
          let _1613_v110;
          _1613_v110 = _dafny.Seq.of((_1612_v109).dtor_cf72, (_1612_v109).dtor_cf72);
          let _1614_v111;
          _1614_v111 = new _dafny.CodePoint('s'.codePointAt(0));
          let _1615_v112;
          _1615_v112 = _dafny.Map.Empty.slice().updateUnsafe(_1589_v101.f12,_1614_v111);
          let _1616_v113;
          _1616_v113 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_1613_v110).length)).plus(new BigNumber((_1535_v79).length)),_1615_v112);
          _1616_v113 = (_1616_v113).update(_1589_v101.f12, _1615_v112);
          let _1617_v114;
          let _nw251 = new _module.C1();
          _nw251.__ctor(_dafny.Seq.Concat(_1530_v74.f13, _1530_v74.f13));
          _1617_v114 = _nw251;
          let _1618_v116;
          _1618_v116 = _module.D14.create_DC38(_1614_v111, (_1571_v85)[_module.__default.safeIndex(new BigNumber((function () {
  let _coll35 = new _dafny.Map();
  for (const _compr_35 of (_module.__default.fm45(globalState)).Elements) {
    let _1619_v115 = _compr_35;
    if (_dafny.Seq.contains(_module.__default.fm45(globalState), _1619_v115)) {
      _coll35.push([_1619_v115,(_this).f9]);
    }
  }
  return _coll35;
}()).length), new BigNumber((_1571_v85).length))], _1571_v85);
          (globalState).f2 = (_1618_v116).dtor_cf68;
          let _1620_v117;
          let _init46 = ((_1621_v101) => function (_1622_i11) {
            return (_1622_i11).plus(_1621_v101.f12);
          })(_1589_v101);
          let _nw252 = Array((new BigNumber(18)).toNumber());
          for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw252.length); _i0_46++) {
            _nw252[_i0_46] = _init46(new BigNumber(_i0_46));
          }
          _1620_v117 = _nw252;
          let _index240 = _module.__default.safeIndex(new BigNumber(734), new BigNumber((_1620_v117).length));
          (_1620_v117)[_index240] = new BigNumber((_dafny.Seq.Concat(_1535_v79, _1535_v79)).length);
          let _index241 = _module.__default.safeIndex(new BigNumber(734), new BigNumber((_1620_v117).length));
          (_1620_v117)[_index241] = (_this).f10;
        } else {
          let _1623___mcc_h31 = (_source16).cf38;
          let _1624_cf38 = _1623___mcc_h31;
          (globalState).f2 = (_1569_v84)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_1569_v84).length))];
          (_1589_v101).f12 = (_this).f10;
          (globalState).f2 = (_1533_v76).f15;
          let _index242 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_1583_v95).length));
          (_1583_v95)[_index242] = _1571_v85;
          let _1625_v118;
          _1625_v118 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,new BigNumber(653));
          let _1626_v119;
          _1626_v119 = _dafny.Map.Empty.slice().updateUnsafe(_1533_v76,_1535_v79);
          let _index243 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_1583_v95).length));
          (_1583_v95)[_index243] = _module.__default.fm41(_module.D10.create_DC27(_1625_v118), (_this).f10, _dafny.Seq.update(_1587_v99, _module.__default.safeIndex(_1589_v101.f12, new BigNumber((_1587_v99).length)), new BigNumber((_1626_v119).length)), globalState);
        }
      }
      r0 = (_1569_v84)[_module.__default.safeIndex(new BigNumber(515), new BigNumber((_1569_v84).length))];
      return r0;
    }
    m4(p0, p1, p2, globalState) {
      let _this = this;
      let _1627_i0;
      _1627_i0 = _dafny.ZERO;
      L9: {
        while (!((_this).f9)) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1627_i0)) {
              break L9;
            }
            _1627_i0 = (_1627_i0).plus(_dafny.ONE);
            let _1628_v0;
            _1628_v0 = new BigNumber(-100);
            let _1629_v1;
            _1629_v1 = new _dafny.CodePoint('r'.codePointAt(0));
            let _1630_v2;
            _1630_v2 = _dafny.Seq.of(_1629_v1);
            let _1631_v3;
            _1631_v3 = _dafny.MultiSet.fromElements(_1629_v1, _1629_v1);
            let _1632_v4;
            _1632_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_dafny.MultiSet.FromArray(_1630_v2)).IsDisjointFrom(_1631_v3));
            _1628_v0 = new BigNumber((_1632_v4).length);
            let _1633_v5;
            let _nw253 = Array((new BigNumber(6)).toNumber()).fill(false);
            _1633_v5 = _nw253;
            let _index244 = _module.__default.safeIndex(new BigNumber(242), new BigNumber((_1633_v5).length));
            (_1633_v5)[_index244] = p2;
            let _index245 = _module.__default.safeIndex(new BigNumber(242), new BigNumber((_1633_v5).length));
            (_1633_v5)[_index245] = (_this).f9;
            _1628_v0 = _module.__default.safeDivisionInt((_this).f10, _1628_v0);
            let _1634_v6;
            _1634_v6 = _dafny.Seq.of((_1633_v5)[_module.__default.safeIndex(new BigNumber(242), new BigNumber((_1633_v5).length))]);
            let _1635_v7;
            _1635_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,new BigNumber((_1634_v6).length));
            let _1636_v8;
            _1636_v8 = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_1635_v7).length)), (_this).f10, _1628_v0, new BigNumber(-171));
            let _index246 = _module.__default.safeIndex(new BigNumber(242), new BigNumber((_1633_v5).length));
            let _rhs282 = true;
            let _rhs283 = (_module.__default.safeModuloInt(_1628_v0, _1628_v0)).plus(new BigNumber(((_dafny.Set.fromElements((_this).f10, new BigNumber(-206))).Difference(_1636_v8)).length));
            let _rhs284 = false;
            let _rhs285 = (_dafny.Map.Empty.slice().updateUnsafe(_1628_v0,(_this).f9)).contains(_1628_v0);
            let _rhs286 = (_this).f10;
            let _lhs175 = _1633_v5;
            let _lhs176 = _module.__default.safeIndex(new BigNumber(242), new BigNumber((_1633_v5).length));
            let _lhs177 = globalState;
            let _lhs178 = globalState;
            _lhs175[_lhs176] = _rhs282;
            _1628_v0 = _rhs283;
            _lhs177.f2 = _rhs284;
            _lhs178.f2 = _rhs285;
            _1628_v0 = _rhs286;
          }
        }
      }
      let _1637_v9;
      _1637_v9 = _dafny.Seq.UnicodeFromString("ua");
      _1637_v9 = _dafny.Seq.UnicodeFromString("iatft");
      let _1638_v10;
      _1638_v10 = new BigNumber(415);
      let _1639_v11;
      let _init47 = function (_1640_i1) {
        return (_this).f9;
      };
      let _nw254 = Array((_dafny.ONE).toNumber());
      for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw254.length); _i0_47++) {
        _nw254[_i0_47] = _init47(new BigNumber(_i0_47));
      }
      _1639_v11 = _nw254;
      let _1641_v12;
      _1641_v12 = _dafny.Map.Empty.slice().updateUnsafe(_1639_v11,(p1).isLessThan(new BigNumber((_dafny.Seq.of(p1)).length)));
      let _1642_v13;
      _1642_v13 = _module.D11.create_DC30();
      let _1643_v14;
      _1643_v14 = _dafny.Seq.of(_module.D11.create_DC30(), _1642_v13, _1642_v13, _1642_v13, _1642_v13);
      let _1644_v15;
      _1644_v15 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_1643_v14);
      let _rhs287 = (_dafny.ZERO).minus((p1).plus((_dafny.ZERO).minus(p1)));
      let _rhs288 = new BigNumber(((_1644_v15).Merge((_dafny.Map.Empty.slice().updateUnsafe((_this).f9,_dafny.Seq.update(_1643_v14, _module.__default.safeIndex((_this).f10, new BigNumber((_1643_v14).length)), _1642_v13))).Merge(_1644_v15))).length);
      let _rhs289 = _1641_v12;
      _1638_v10 = _rhs287;
      _1638_v10 = _rhs288;
      _1641_v12 = _rhs289;
      let _1645_v16;
      _1645_v16 = _dafny.Map.Empty.slice().updateUnsafe(_1637_v9,(_this).f9);
      let _1646_v17;
      _1646_v17 = _module.D9.create_DC23(_1645_v16);
      let _pat_let_tv83 = _1645_v16;
      let _1647_v18;
      _1647_v18 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f9) ? (function (_pat_let62_0) {
        return function (_1648_dt__update__tmp_h0) {
          return function (_pat_let63_0) {
            return function (_1649_dt__update_hcf40_h0) {
              return _module.D9.create_DC23(_1649_dt__update_hcf40_h0);
            }(_pat_let63_0);
          }(_pat_let_tv83);
        }(_pat_let62_0);
      }(_1646_v17)) : (_1646_v17)),_1637_v9);
      _1647_v18 = (_1647_v18).update(_1646_v17, _1637_v9);
      let _1650_v19;
      _1650_v19 = new _dafny.CodePoint('j'.codePointAt(0));
      _1650_v19 = _1650_v19;
      let _1651_v20;
      let _nw255 = Array((new BigNumber(10)).toNumber()).fill([]);
      _1651_v20 = _nw255;
      let _index247 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((_1651_v20).length));
      (_1651_v20)[_index247] = p0;
      let _1652_v21;
      let _nw256 = Array((new BigNumber(18)).toNumber());
      _1652_v21 = _nw256;
      let _1653_v22;
      _1653_v22 = _dafny.Seq.of(_1652_v21);
      let _index248 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((_1651_v20).length));
      let _rhs290 = p0;
      let _rhs291 = (_1653_v22)[_module.__default.safeIndex((_this).f10, new BigNumber((_1653_v22).length))];
      let _lhs179 = _1651_v20;
      let _lhs180 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((_1651_v20).length));
      _lhs179[_lhs180] = _rhs290;
      _1652_v21 = _rhs291;
      return;
    }
    m5(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let _1654_i0;
      _1654_i0 = _dafny.ZERO;
      L10: {
        while ((((_this).f9) ? ((_this).f9) : ((_this).f9))) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1654_i0)) {
              break L10;
            }
            _1654_i0 = (_1654_i0).plus(_dafny.ONE);
            let _1655_v0;
            let _init48 = function (_1656_i1) {
              return new _dafny.CodePoint('x'.codePointAt(0));
            };
            let _nw257 = Array((new BigNumber(16)).toNumber());
            for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw257.length); _i0_48++) {
              _nw257[_i0_48] = _init48(new BigNumber(_i0_48));
            }
            _1655_v0 = _nw257;
            let _1657_v1;
            _1657_v1 = _module.D13.create_DC34(_1655_v0);
            let _1658_v2;
            _1658_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_1657_v1);
            let _1659_v3;
            _1659_v3 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,new BigNumber((_1658_v2).length));
            _1659_v3 = (_1659_v3).Merge(((false) ? (_1659_v3) : ((_1659_v3).update((_this).f9, (_this).f10))));
            let _1660_v4;
            _1660_v4 = _dafny.Seq.of((_this).f9, (_this).f9);
            let _1661_v5;
            _1661_v5 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.Concat(_1660_v4, _1660_v4)).length), (_this).f10, (((_this).f9) ? (new BigNumber((_dafny.Seq.of((_this).f9, (_this).f9)).length)) : ((_this).f10)), new BigNumber(827));
            _1661_v5 = _1661_v5;
            let _1662_v6;
            let _init49 = function (_1663_i2) {
              return _dafny.MultiSet.fromElements(new BigNumber(430), new BigNumber((_dafny.Seq.UnicodeFromString("j")).length), (_this).f10, (_this).f10, (_dafny.ZERO).minus((_this).f10));
            };
            let _nw258 = Array((new BigNumber(16)).toNumber());
            for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw258.length); _i0_49++) {
              _nw258[_i0_49] = _init49(new BigNumber(_i0_49));
            }
            _1662_v6 = _nw258;
            _1662_v6 = _1662_v6;
            let _1664_v7;
            let _nw259 = Array((new BigNumber(13)).toNumber()).fill(_module.D12.Default());
            _1664_v7 = _nw259;
            let _1665_v8;
            _1665_v8 = _dafny.Seq.of(_1664_v7, _1664_v7);
            r0 = ((_this).f9) || (((_this).f10).isLessThanOrEqualTo(new BigNumber((_1665_v8).length)));
          }
        }
      }
      let _1666_v9;
      let _nw260 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Set.Empty);
      _1666_v9 = _nw260;
      let _1667_v10;
      let _nw261 = new _module.C3();
      _nw261.__ctor((_this).f9, _1666_v9, (_this).f9);
      _1667_v10 = _nw261;
      let _1668_v11;
      let _nw262 = Array((new BigNumber(12)).toNumber());
      _nw262[(_dafny.ZERO).toNumber()] = _1667_v10;
      _nw262[(_dafny.ONE).toNumber()] = _1667_v10;
      _nw262[(new BigNumber(2)).toNumber()] = _1667_v10;
      _nw262[(new BigNumber(3)).toNumber()] = _1667_v10;
      _nw262[(new BigNumber(4)).toNumber()] = _1667_v10;
      _nw262[(new BigNumber(5)).toNumber()] = _1667_v10;
      _nw262[(new BigNumber(6)).toNumber()] = _1667_v10;
      _nw262[(new BigNumber(7)).toNumber()] = _1667_v10;
      _nw262[(new BigNumber(8)).toNumber()] = _1667_v10;
      _nw262[(new BigNumber(9)).toNumber()] = _1667_v10;
      _nw262[(new BigNumber(10)).toNumber()] = _1667_v10;
      _nw262[(new BigNumber(11)).toNumber()] = _1667_v10;
      _1668_v11 = _nw262;
      let _index249 = _module.__default.safeIndex(new BigNumber(407), new BigNumber((_1668_v11).length));
      (_1668_v11)[_index249] = _1667_v10;
      let _index250 = _module.__default.safeIndex(new BigNumber(407), new BigNumber((_1668_v11).length));
      (_1668_v11)[_index250] = _1667_v10;
      let _1669_v12;
      _1669_v12 = _dafny.Seq.of((_1667_v10).f15, (_this).fm3(globalState));
      let _hi7 = new BigNumber((_1669_v12).length);
      for (let _1670_i3 = (_dafny.ZERO).minus((_this).f10); _1670_i3.isLessThan(_hi7); _1670_i3 = _1670_i3.plus(_dafny.ONE)) {
        let _1671_v13;
        _1671_v13 = _module.D4.create_DC13((_this).f10);
        let _1672_v14;
        _1672_v14 = _dafny.Seq.of(_dafny.Set.fromElements((_1671_v13).dtor_cf25));
        let _1673_v15;
        _1673_v15 = _dafny.Seq.of((_this).f10);
        let _1674_v16;
        _1674_v16 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_dafny.MultiSet.FromArray(_1673_v15), _1670_i3, globalState),(_1667_v10).f15);
        (globalState).f2 = ((_1667_v10).f15) || ((_1667_v10).fm32(_1672_v14, (_1667_v10).f15, new BigNumber((_1674_v16).length), (_this).f9, globalState));
        let _1675_v17;
        let _nw263 = new _module.C2();
        _nw263.__ctor(_1670_i3, (_this).f9);
        _1675_v17 = _nw263;
        let _1676_v18;
        _1676_v18 = _dafny.Map.Empty.slice().updateUnsafe((_1667_v10).f15,_1675_v17);
        let _1677_v19;
        _1677_v19 = _dafny.Set.fromElements(_1670_i3, new BigNumber((_1676_v18).length));
        let _1678_v20;
        _1678_v20 = _1677_v19;
        let _source17 = _1678_v20;
        let _1679___mcc_h0 = _source17;
        let _1680_cf15 = _1679___mcc_h0;
        let _1681_v21;
        _1681_v21 = _module.D13.create_DC36(_dafny.Seq.UnicodeFromString("hpwdus"), _1670_i3, (_this).f10, (_this).f10);
        r2 = ((_1681_v21).dtor_cf64).multipliedBy(_1670_i3);
        let _1682_v22;
        let _nw264 = Array((new BigNumber(13)).toNumber()).fill(false);
        _1682_v22 = _nw264;
        let _index251 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_1682_v22).length));
        (_1682_v22)[_index251] = (_1667_v10).f15;
        let _1683_v23;
        _1683_v23 = _dafny.Map.Empty.slice().updateUnsafe((_1675_v17).f9,_1670_i3);
        let _1684_v25;
        _1684_v25 = _dafny.Set.fromElements(_1673_v15);
        let _1685_v26;
        let _nw265 = new _module.C0();
        _nw265.__ctor(_1684_v25, _1670_i3);
        _1685_v26 = _nw265;
        let _1686_v27;
        _1686_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1685_v26,(_this).f10);
        let _1687_v28;
        _1687_v28 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-4),_1686_v27);
        let _1688_v29;
        _1688_v29 = _dafny.MultiSet.fromElements((((_1683_v23).contains((_1667_v10).f15)) ? ((_1683_v23).get((_1667_v10).f15)) : (new BigNumber((function () {
          let _coll36 = new _dafny.Set();
          for (const _compr_36 of _dafny.IntegerRange(new BigNumber(-117), new BigNumber(410))) {
            let _1689_v24 = _compr_36;
            if (((new BigNumber(-117)).isLessThanOrEqualTo(_1689_v24)) && ((_1689_v24).isLessThan(new BigNumber(410)))) {
              _coll36.add(_module.__default.safeDivisionInt(_1689_v24, new BigNumber(-521)));
            }
          }
          return _coll36;
        }()).length))), _1670_i3, new BigNumber(((((_1687_v28).contains((_this).f10)) ? ((_1687_v28).get((_this).f10)) : (_1686_v27))).length));
        let _index252 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_1682_v22).length));
        (_1682_v22)[_index252] = (((_1667_v10).f15) ? ((_1667_v10).f15) : (!((_module.__default.fm29((_this).f9, (_this).f9, globalState)).IsSubsetOf(_1688_v29))));
        r1 = new BigNumber(342);
        (_1685_v26).f12 = (_dafny.ZERO).minus(_1670_i3);
        let _1690_v30;
        _1690_v30 = _dafny.Set.fromElements((_1675_v17).f9, (_1667_v10).f15, (_1667_v10).f15);
        let _rhs292 = !((_this).f10).isEqualTo((_this).f10);
        let _rhs293 = _1690_v30;
        let _rhs294 = _module.__default.safeModuloInt(_1670_i3, _module.__default.safeModuloInt((_1673_v15)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f10), new BigNumber((_1673_v15).length))], _1670_i3));
        r0 = _rhs292;
        _1690_v30 = _rhs293;
        r2 = _rhs294;
        let _1691_v31;
        let _nw266 = new _module.C2();
        _nw266.__ctor(_1670_i3, !((_module.__default.fm0(globalState)).isLessThan(new BigNumber((_dafny.Seq.UnicodeFromString("uyy")).length))));
        _1691_v31 = _nw266;
      }
      let _1692_v32;
      _1692_v32 = _dafny.Set.fromElements((_1667_v10).f15);
      let _1693_v33;
      _1693_v33 = _module.D15.create_DC40(_1692_v32);
      let _source18 = _1693_v33;
      if (_source18.is_DC41) {
        r2 = (_this).f10;
        let _1694_v34;
        let _init50 = ((_1695_v10) => function (_1696_i4) {
          return (_1695_v10).f15;
        })(_1667_v10);
        let _nw267 = Array((new BigNumber(15)).toNumber());
        for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw267.length); _i0_50++) {
          _nw267[_i0_50] = _init50(new BigNumber(_i0_50));
        }
        _1694_v34 = _nw267;
        let _index253 = _module.__default.safeIndex(new BigNumber(472), new BigNumber((_1694_v34).length));
        (_1694_v34)[_index253] = (_this).f9;
        let _index254 = _module.__default.safeIndex(new BigNumber(472), new BigNumber((_1694_v34).length));
        (_1694_v34)[_index254] = (_1667_v10).f15;
        let _1697_v35;
        _1697_v35 = _dafny.Seq.of(new BigNumber((_1692_v32).length), (_this).f10);
        let _1698_v36;
        _1698_v36 = _dafny.Set.fromElements(_1697_v35);
        let _1699_v37;
        let _init51 = function (_1700_i5) {
          return _module.__default.safeDivisionInt(_1700_i5, (_this).f10);
        };
        let _nw268 = Array((new BigNumber(11)).toNumber());
        for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw268.length); _i0_51++) {
          _nw268[_i0_51] = _init51(new BigNumber(_i0_51));
        }
        _1699_v37 = _nw268;
        let _1701_v38;
        _1701_v38 = _dafny.Set.fromElements(_1699_v37, _1699_v37, _1699_v37);
        let _1702_v39;
        _1702_v39 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm5(_1698_v36, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1701_v38).length),_1669_v12), new BigNumber((_dafny.Seq.of((_this).f9)).length), globalState),(_this).f10);
        let _1703_v40;
        _1703_v40 = _dafny.MultiSet.fromElements(new BigNumber((_1702_v39).length), (_this).f10);
        let _1704_v41;
        _1704_v41 = _dafny.Map.Empty.slice().updateUnsafe((_1667_v10).f15,_module.__default.fm1(_1703_v40, (_1697_v35)[_module.__default.safeIndex((_this).f10, new BigNumber((_1697_v35).length))], globalState));
        _1704_v41 = (_1704_v41).update((_1667_v10).f15, !((_1694_v34)[_module.__default.safeIndex(new BigNumber(472), new BigNumber((_1694_v34).length))]));
        let _1705_v42;
        _1705_v42 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('x'.codePointAt(0)));
        let _1706_v43;
        let _nw269 = new _module.C0();
        _nw269.__ctor(_1698_v36, new BigNumber((_1705_v42).cardinality()));
        _1706_v43 = _nw269;
        _1706_v43 = _1706_v43;
      } else if (_source18.is_DC42) {
        let _1707___mcc_h1 = (_source18).cf73;
        let _1708___mcc_h2 = (_source18).cf74;
        let _1709___mcc_h3 = (_source18).cf75;
        let _1710___mcc_h4 = (_source18).cf76;
        let _1711_cf76 = _1710___mcc_h4;
        let _1712_cf75 = _1709___mcc_h3;
        let _1713_cf74 = _1708___mcc_h2;
        let _1714_cf73 = _1707___mcc_h1;
        let _1715_v45;
        _1715_v45 = _dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(_1711_cf76, (_this).f10)).cardinality()));
        let _1716_v46;
        _1716_v46 = _dafny.MultiSet.fromElements(new BigNumber(48), _1711_cf76);
        let _1717_v47;
        _1717_v47 = _dafny.MultiSet.fromElements(_dafny.MultiSet.FromArray(_1715_v45), _1716_v46);
        let _1718_v48;
        _1718_v48 = new _dafny.CodePoint('h'.codePointAt(0));
        let _1719_v49;
        _1719_v49 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements((_this).f10, _1713_cf74),_1718_v48);
        let _1720_v50;
        _1720_v50 = _dafny.Map.Empty.slice().updateUnsafe(((!(!(_1714_cf73))) ? (function () {
          let _coll37 = new _dafny.Map();
          for (const _compr_37 of (_1717_v47).Elements) {
            let _1721_v44 = _compr_37;
            if ((_1717_v47).contains(_1721_v44)) {
              _coll37.push([_1721_v44,new _dafny.CodePoint('r'.codePointAt(0))]);
            }
          }
          return _coll37;
        }()) : (_1719_v49)),_module.__default.safeModuloInt((_this).f10, new BigNumber((_1715_v45).length)));
        _1720_v50 = (_1720_v50).update(_1719_v49, (_this).f10);
        let _1722_v51;
        let _nw270 = Array((new BigNumber(9)).toNumber()).fill(_module.D4.Default());
        _1722_v51 = _nw270;
        let _1723_v52;
        _1723_v52 = _dafny.Set.fromElements((((_this).f9) ? (_1722_v51) : (_1722_v51)));
        _1723_v52 = _1723_v52;
        _1713_cf74 = _1713_cf74;
        let _1724_v53;
        _1724_v53 = _dafny.Set.fromElements((_this).f10);
        let _1725_v54;
        _1725_v54 = _dafny.Seq.of(_1724_v53, _dafny.Set.fromElements(_1711_cf76));
        let _1726_v55;
        _1726_v55 = _dafny.MultiSet.fromElements((_1667_v10).f15);
        let _1727_v56;
        _1727_v56 = _dafny.Seq.UnicodeFromString("bxndbo");
        let _1728_v57;
        _1728_v57 = _module.D1.create_DC2(new BigNumber(((_1726_v55).update(true, _module.__default.abs(new BigNumber((_1727_v56).length)))).cardinality()), _1714_cf73, new BigNumber((_1727_v56).length), _1711_cf76);
        let _1729_v58;
        _1729_v58 = _dafny.Map.Empty.slice().updateUnsafe(_1718_v48,_1716_v46);
        let _1730_v59;
        let _nw271 = Array((new BigNumber(25)).toNumber());
        _nw271[(_dafny.ZERO).toNumber()] = (_1667_v10).fm32(_1725_v54, (_this).fm3(globalState), new BigNumber(328), (_1728_v57).dtor_cf3, globalState);
        _nw271[(_dafny.ONE).toNumber()] = (((_1667_v10).f15) ? (true) : ((_1712_cf75).f9));
        _nw271[(new BigNumber(2)).toNumber()] = (_this).f9;
        _nw271[(new BigNumber(3)).toNumber()] = !((((_1712_cf75).f9) ? ((_1712_cf75).f9) : ((_1712_cf75).f9)));
        _nw271[(new BigNumber(4)).toNumber()] = _1714_cf73;
        _nw271[(new BigNumber(5)).toNumber()] = false;
        _nw271[(new BigNumber(6)).toNumber()] = (_1712_cf75).f9;
        _nw271[(new BigNumber(7)).toNumber()] = (_1667_v10).f15;
        _nw271[(new BigNumber(8)).toNumber()] = (_this).f9;
        _nw271[(new BigNumber(9)).toNumber()] = (_this).f9;
        _nw271[(new BigNumber(10)).toNumber()] = _1714_cf73;
        _nw271[(new BigNumber(11)).toNumber()] = !((_1712_cf75).fm3(globalState));
        _nw271[(new BigNumber(12)).toNumber()] = (_1713_cf74).isLessThanOrEqualTo((_this).f10);
        _nw271[(new BigNumber(13)).toNumber()] = (_1712_cf75).f9;
        _nw271[(new BigNumber(14)).toNumber()] = (_this).f9;
        _nw271[(new BigNumber(15)).toNumber()] = (_1667_v10).f15;
        _nw271[(new BigNumber(16)).toNumber()] = !((_1667_v10).f15) || (false);
        _nw271[(new BigNumber(17)).toNumber()] = (_1712_cf75).fm3(globalState);
        _nw271[(new BigNumber(18)).toNumber()] = _1714_cf73;
        _nw271[(new BigNumber(19)).toNumber()] = (_1711_cf76).isLessThan((_this).f10);
        _nw271[(new BigNumber(20)).toNumber()] = (_1712_cf75).f9;
        _nw271[(new BigNumber(21)).toNumber()] = (_1667_v10).f15;
        _nw271[(new BigNumber(22)).toNumber()] = ((_this).f10).isLessThan(_1713_cf74);
        _nw271[(new BigNumber(23)).toNumber()] = !(_module.__default.fm1((((_1729_v58).contains((_1727_v56)[_module.__default.safeIndex(new BigNumber(299), new BigNumber((_1727_v56).length))])) ? ((_1729_v58).get((_1727_v56)[_module.__default.safeIndex(new BigNumber(299), new BigNumber((_1727_v56).length))])) : (_1716_v46)), (_this).f10, globalState));
        _nw271[(new BigNumber(24)).toNumber()] = (_1667_v10).f15;
        _1730_v59 = _nw271;
        let _1731_v60;
        let _nw272 = new _module.C1();
        _nw272.__ctor(_dafny.Seq.of(_1728_v57, _1728_v57));
        _1731_v60 = _nw272;
        let _1732_v61;
        _1732_v61 = _module.D13.create_DC35(_1714_cf73, _1718_v48, _1731_v60, _1714_cf73);
        let _index255 = _module.__default.safeIndex(new BigNumber(10), new BigNumber((_1730_v59).length));
        (_1730_v59)[_index255] = (_1732_v61).dtor_cf58;
        let _1733_v62;
        _1733_v62 = _dafny.Map.Empty.slice().updateUnsafe(_1713_cf74,_1713_cf74);
        let _index256 = _module.__default.safeIndex(new BigNumber(10), new BigNumber((_1730_v59).length));
        let _rhs295 = (_1669_v12)[_module.__default.safeIndex(_module.__default.safeModuloInt(new BigNumber((_1733_v62).length), _1713_cf74), new BigNumber((_1669_v12).length))];
        let _rhs296 = _1669_v12;
        let _rhs297 = _1714_cf73;
        let _lhs181 = _1730_v59;
        let _lhs182 = _module.__default.safeIndex(new BigNumber(10), new BigNumber((_1730_v59).length));
        let _lhs183 = globalState;
        _lhs181[_lhs182] = _rhs295;
        _1669_v12 = _rhs296;
        _lhs183.f2 = _rhs297;
      } else if (_source18.is_DC43) {
        let _1734_v63;
        _1734_v63 = _dafny.Seq.UnicodeFromString("nf");
        let _1735_v64;
        _1735_v64 = _module.D1.create_DC2((_this).f10, (_1667_v10).f15, (_this).f10, new BigNumber((_1734_v63).length));
        let _1736_v65;
        _1736_v65 = _dafny.Seq.of(_1735_v64);
        let _1737_v66;
        let _nw273 = new _module.C1();
        _nw273.__ctor(_1736_v65);
        _1737_v66 = _nw273;
        let _1738_v67;
        _1738_v67 = _dafny.MultiSet.fromElements(_1737_v66, _1737_v66);
        let _1739_v68;
        _1739_v68 = _dafny.Map.Empty.slice().updateUnsafe(_1738_v67,(_1667_v10).f15);
        _1739_v68 = (_1739_v68).update((_1738_v67).update(_1737_v66, _module.__default.abs(new BigNumber(677))), (_1667_v10).f15);
        let _1740_v69;
        _1740_v69 = _dafny.Set.fromElements(new BigNumber((_1734_v63).length));
        let _1741_v70;
        _1741_v70 = _dafny.Seq.of(_dafny.Set.fromElements((_this).f10, (_this).f10), _1740_v69);
        let _1742_v71;
        _1742_v71 = _module.D3.create_DC9((_this).f9, _1669_v12);
        let _1743_v72;
        _1743_v72 = _dafny.Seq.of(new BigNumber(-45), _module.__default.safeModuloInt((_this).fm4(false, (_1667_v10).fm32(_1741_v70, (_1742_v71).dtor_cf17, (_this).f10, (_this).fm3(globalState), globalState), (_1667_v10).f15, (_1667_v10).f15, globalState), (_this).f10));
        _1743_v72 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-392)), function (_1744_i6) {
          return (_dafny.ZERO).minus((_this).f10);
        });
        r0 = (_1669_v12)[_module.__default.safeIndex((_this).f10, new BigNumber((_1669_v12).length))];
        let _1745_v73;
        _1745_v73 = _dafny.Map.Empty.slice().updateUnsafe(_1743_v72,_1734_v63);
        let _1746_v75;
        _1746_v75 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_dafny.Seq.of((_1667_v10).f15));
        let _1747_v76;
        let _nw274 = Array((new BigNumber(11)).toNumber());
        _nw274[(_dafny.ZERO).toNumber()] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(false), _1669_v12);
        _nw274[(_dafny.ONE).toNumber()] = (_this).f9;
        _nw274[(new BigNumber(2)).toNumber()] = (_this).f9;
        _nw274[(new BigNumber(3)).toNumber()] = (_this).fm5(function () {
          let _coll38 = new _dafny.Set();
          for (const _compr_38 of (_1745_v73).Keys.Elements) {
            let _1748_v74 = _compr_38;
            if ((_1745_v73).contains(_1748_v74)) {
              _coll38.add(_1748_v74);
            }
          }
          return _coll38;
        }(), _1746_v75, (_this).f10, globalState);
        _nw274[(new BigNumber(4)).toNumber()] = (_1667_v10).f15;
        _nw274[(new BigNumber(5)).toNumber()] = (_1667_v10).f15;
        _nw274[(new BigNumber(6)).toNumber()] = (_this).f9;
        _nw274[(new BigNumber(7)).toNumber()] = (_1667_v10).f15;
        _nw274[(new BigNumber(8)).toNumber()] = false;
        _nw274[(new BigNumber(9)).toNumber()] = (_1667_v10).fm3(globalState);
        _nw274[(new BigNumber(10)).toNumber()] = true;
        _1747_v76 = _nw274;
        let _index257 = _module.__default.safeIndex(new BigNumber(130), new BigNumber((_1747_v76).length));
        (_1747_v76)[_index257] = (_1667_v10).f15;
        let _index258 = _module.__default.safeIndex(new BigNumber(130), new BigNumber((_1747_v76).length));
        (_1747_v76)[_index258] = !((_1667_v10).f15);
      } else if (_source18.is_DC40) {
        let _1749___mcc_h5 = (_source18).cf72;
        let _1750_cf72 = _1749___mcc_h5;
        let _1751_v77;
        _1751_v77 = _dafny.Seq.UnicodeFromString("lf");
        let _1752_v78;
        _1752_v78 = _module.D1.create_DC2((_this).f10, (_1667_v10).f15, new BigNumber((_1751_v77).length), (_this).f10);
        let _1753_v79;
        _1753_v79 = _dafny.Seq.of(_1752_v78, _1752_v78);
        let _1754_v80;
        let _nw275 = new _module.C1();
        _nw275.__ctor(_1753_v79);
        _1754_v80 = _nw275;
        let _1755_v81;
        _1755_v81 = _dafny.Seq.of((_this).f10);
        _1755_v81 = _1755_v81;
        r1 = (_this).f10;
        _1755_v81 = _1755_v81;
      } else {
        let _1756___mcc_h6 = (_source18).cf77;
        let _1757_cf77 = _1756___mcc_h6;
        let _1758_v82;
        _1758_v82 = _dafny.Seq.UnicodeFromString("dxjy");
        let _1759_v83;
        _1759_v83 = _dafny.Map.Empty.slice().updateUnsafe(_1758_v82,(_dafny.ZERO).minus((_this).f10));
        let _1760_v84;
        _1760_v84 = _dafny.Set.fromElements(((_dafny.ZERO).minus(new BigNumber((_1759_v83).length))).minus((_this).f10), (_this).f10, (_this).f10);
        let _1761_v85;
        _1761_v85 = new _dafny.CodePoint('l'.codePointAt(0));
        let _1762_v86;
        _1762_v86 = _dafny.Map.Empty.slice().updateUnsafe((_module.D7.create_DC20(true, _1761_v85)).dtor_cf36,!(true));
        _1760_v84 = ((_1760_v84).Difference(_dafny.Set.fromElements((_this).f10, new BigNumber((_1669_v12).length), new BigNumber((_1762_v86).length), (_this).f10))).Union(_1760_v84);
        let _1763_v87;
        let _nw276 = new _module.C3();
        _nw276.__ctor((_1667_v10).f15, _1666_v9, (_1667_v10).f15);
        _1763_v87 = _nw276;
        let _1764_v88;
        let _nw277 = Array((new BigNumber(28)).toNumber());
        _nw277[(_dafny.ZERO).toNumber()] = _1763_v87;
        _nw277[(_dafny.ONE).toNumber()] = _1763_v87;
        _nw277[(new BigNumber(2)).toNumber()] = _1763_v87;
        _nw277[(new BigNumber(3)).toNumber()] = _1763_v87;
        _nw277[(new BigNumber(4)).toNumber()] = _1763_v87;
        _nw277[(new BigNumber(5)).toNumber()] = _1763_v87;
        _nw277[(new BigNumber(6)).toNumber()] = _1763_v87;
        _nw277[(new BigNumber(7)).toNumber()] = _1763_v87;
        _nw277[(new BigNumber(8)).toNumber()] = _1763_v87;
        _nw277[(new BigNumber(9)).toNumber()] = _1763_v87;
        _nw277[(new BigNumber(10)).toNumber()] = _1763_v87;
        _nw277[(new BigNumber(11)).toNumber()] = _1763_v87;
        _nw277[(new BigNumber(12)).toNumber()] = _1763_v87;
        _nw277[(new BigNumber(13)).toNumber()] = _1763_v87;
        _nw277[(new BigNumber(14)).toNumber()] = _1763_v87;
        _nw277[(new BigNumber(15)).toNumber()] = _1763_v87;
        _nw277[(new BigNumber(16)).toNumber()] = _1763_v87;
        _nw277[(new BigNumber(17)).toNumber()] = _1763_v87;
        _nw277[(new BigNumber(18)).toNumber()] = _1763_v87;
        _nw277[(new BigNumber(19)).toNumber()] = _1763_v87;
        _nw277[(new BigNumber(20)).toNumber()] = _1763_v87;
        _nw277[(new BigNumber(21)).toNumber()] = _1763_v87;
        _nw277[(new BigNumber(22)).toNumber()] = _1763_v87;
        _nw277[(new BigNumber(23)).toNumber()] = _1763_v87;
        _nw277[(new BigNumber(24)).toNumber()] = _1763_v87;
        _nw277[(new BigNumber(25)).toNumber()] = _1763_v87;
        _nw277[(new BigNumber(26)).toNumber()] = _1763_v87;
        _nw277[(new BigNumber(27)).toNumber()] = _1763_v87;
        _1764_v88 = _nw277;
        let _index259 = _module.__default.safeIndex(new BigNumber(388), new BigNumber((_1764_v88).length));
        (_1764_v88)[_index259] = _1763_v87;
        let _index260 = _module.__default.safeIndex(new BigNumber(388), new BigNumber((_1764_v88).length));
        (_1764_v88)[_index260] = _1763_v87;
        _1758_v82 = _1758_v82;
        if (!(((((_this).f10).isLessThanOrEqualTo((_this).f10)) ? ((((_1667_v10).f15) ? ((_1667_v10).f15) : ((_1763_v87).f9))) : ((_this).f9)))) {
          let _1765_v89;
          _1765_v89 = _dafny.Seq.of(((_this).f10).minus((_this).f10));
          let _1766_v90;
          let _nw278 = Array((new BigNumber(18)).toNumber()).fill(false);
          _1766_v90 = _nw278;
          let _index261 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_1766_v90).length));
          (_1766_v90)[_index261] = true;
          let _1767_v91;
          _1767_v91 = _dafny.MultiSet.fromElements((_this).f10);
          let _index262 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_1766_v90).length));
          let _rhs298 = _dafny.Seq.Concat(_1765_v89, _dafny.Seq.Concat(_dafny.Seq.of((_this).f10), _1765_v89));
          let _rhs299 = _module.__default.fm1(_1767_v91, (_this).f10, globalState);
          let _lhs184 = _1766_v90;
          let _lhs185 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_1766_v90).length));
          _1765_v89 = _rhs298;
          _lhs184[_lhs185] = _rhs299;
          let _1768_v92;
          _1768_v92 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(242),_1758_v82);
          let _1769_v93;
          _1769_v93 = _dafny.Map.Empty.slice().updateUnsafe((_1763_v87).f9,(((_1768_v92).contains(new BigNumber((_dafny.Seq.UnicodeFromString("lm")).length))) ? ((_1768_v92).get(new BigNumber((_dafny.Seq.UnicodeFromString("lm")).length))) : (_1758_v82)));
          let _1770_v94;
          _1770_v94 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-116)), ((_1771_v85) => function (_1772_i9) {
            return _1771_v85;
          })(_1761_v85)));
          let _1773_v95;
          let _nw279 = Array((new BigNumber(25)).toNumber());
          _nw279[(_dafny.ZERO).toNumber()] = _1758_v82;
          _nw279[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_1758_v82, _dafny.Seq.UnicodeFromString("ouulicrh"));
          _nw279[(new BigNumber(2)).toNumber()] = _1758_v82;
          _nw279[(new BigNumber(3)).toNumber()] = _1758_v82;
          _nw279[(new BigNumber(4)).toNumber()] = (((_1667_v10).f15) ? (_dafny.Seq.UnicodeFromString("cc")) : (_1758_v82));
          _nw279[(new BigNumber(5)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(652)), ((_1774_v85) => function (_1775_i7) {
            return _1774_v85;
          })(_1761_v85));
          _nw279[(new BigNumber(6)).toNumber()] = _dafny.Seq.UnicodeFromString("a");
          _nw279[(new BigNumber(7)).toNumber()] = _1758_v82;
          _nw279[(new BigNumber(8)).toNumber()] = _1758_v82;
          _nw279[(new BigNumber(9)).toNumber()] = _1758_v82;
          _nw279[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("anxityp"), _dafny.Seq.UnicodeFromString("j"));
          _nw279[(new BigNumber(11)).toNumber()] = _1758_v82;
          _nw279[(new BigNumber(12)).toNumber()] = (((_1769_v93).contains((_1763_v87).f9)) ? ((_1769_v93).get((_1763_v87).f9)) : (_1758_v82));
          _nw279[(new BigNumber(13)).toNumber()] = _1758_v82;
          _nw279[(new BigNumber(14)).toNumber()] = _module.__default.fm28((_this).f10, !((_this).f9), globalState);
          _nw279[(new BigNumber(15)).toNumber()] = _1758_v82;
          _nw279[(new BigNumber(16)).toNumber()] = _1758_v82;
          _nw279[(new BigNumber(17)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(839)), ((_1776_v85) => function (_1777_i8) {
            return _1776_v85;
          })(_1761_v85));
          _nw279[(new BigNumber(18)).toNumber()] = (_1770_v94)[_module.__default.safeIndex((_this).f10, new BigNumber((_1770_v94).length))];
          _nw279[(new BigNumber(19)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(859)), ((_1778_v85) => function (_1779_i10) {
            return _1778_v85;
          })(_1761_v85));
          _nw279[(new BigNumber(20)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(975)), ((_1780_v85) => function (_1781_i11) {
            return _1780_v85;
          })(_1761_v85));
          _nw279[(new BigNumber(21)).toNumber()] = _1758_v82;
          _nw279[(new BigNumber(22)).toNumber()] = _1758_v82;
          _nw279[(new BigNumber(23)).toNumber()] = _1758_v82;
          _nw279[(new BigNumber(24)).toNumber()] = _dafny.Seq.UnicodeFromString("fc");
          _1773_v95 = _nw279;
          let _1782_v96;
          _1782_v96 = _dafny.Map.Empty.slice().updateUnsafe(_1768_v92,_1773_v95);
          _1773_v95 = (((_1782_v96).contains((_1768_v92).Merge(_1768_v92))) ? ((_1782_v96).get((_1768_v92).Merge(_1768_v92))) : (_1773_v95));
          let _1783_v97;
          _1783_v97 = _dafny.MultiSet.fromElements((_1766_v90)[_module.__default.safeIndex(new BigNumber(590), new BigNumber((_1766_v90).length))]);
          _1783_v97 = _1783_v97;
          let _1784_v98;
          let _nw280 = new _module.C2();
          _nw280.__ctor((_this).f10, (_1763_v87).f9);
          _1784_v98 = _nw280;
          let _1785_v99;
          _1785_v99 = _dafny.Map.Empty.slice().updateUnsafe(!(true),((_1784_v98).f14).plus((_this).f10));
          _1785_v99 = (_1785_v99).update(false, (_1784_v98).f14);
        } else {
          r1 = (_this).f10;
          let _1786_v100;
          _1786_v100 = _dafny.Seq.of((_this).f10, (((_1667_v10).f15) ? ((_this).f10) : ((_this).f10)), (_this).f10, (_this).f10, new BigNumber((_1758_v82).length));
          _1786_v100 = _dafny.Seq.update(_dafny.Seq.of((_this).f10), _module.__default.safeIndex((_this).f10, new BigNumber((_dafny.Seq.of((_this).f10)).length)), (_this).f10);
          let _1787_v101;
          let _nw281 = new _module.C2();
          _nw281.__ctor((_dafny.ZERO).minus((_this).f10), (_1667_v10).f15);
          _1787_v101 = _nw281;
          let _1788_v102;
          _1788_v102 = _dafny.MultiSet.fromElements((_1763_v87).f9, (_1763_v87).f9);
          let _1789_v103;
          let _out17;
          _out17 = (_1667_v10).m3(((_dafny.MultiSet.fromElements((_1667_v10).f15)).update((_1669_v12)[_module.__default.safeIndex((_this).f10, new BigNumber((_1669_v12).length))], _module.__default.abs((_this).f10))).IsDisjointFrom(_1788_v102), globalState);
          _1789_v103 = _out17;
          r0 = ((_dafny.ZERO).minus((_1787_v101).f14)).isLessThan(new BigNumber((_1669_v12).length));
        }
      }
      let _1790_v104;
      let _nw282 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
      _1790_v104 = _nw282;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1790_v104).length))) {
        let _1791_i12 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1791_i12)) && ((_1791_i12).isLessThan(new BigNumber((_1790_v104).length))))) {
          (_1790_v104)[(_1791_i12)] = (_1791_i12).plus((_this).f10);
        }
      }
      let _1792_v105;
      _1792_v105 = _module.D16.create_DC46(((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-44)), function (_1793_i13) {
  return new _dafny.CodePoint('s'.codePointAt(0));
})).length))).minus((_this).f10));
      _1792_v105 = _1792_v105;
      r0 = (_1667_v10).f15;
      r1 = (_this).f10;
      r2 = (_this).f10;
      return [r0, r1, r2];
    }
    get f10() {
      let _this = this;
      return _this._f10;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
