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
    static fm0(p0, p1, globalState) {
      return new BigNumber(991);
    };
    static fm1(p0, p1, p2, p3, globalState) {
      return (function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(409), new BigNumber(-296))) {
          let _0_v0 = _compr_0;
          if (((new BigNumber(409)).isLessThanOrEqualTo(_0_v0)) && ((_0_v0).isLessThan(new BigNumber(-296)))) {
            _coll0.push([(_0_v0).plus(new BigNumber(260)),_module.D0.create_DC3(_module.D0.create_DC3(_module.D0.create_DC3(_module.D0.create_DC2(false, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(202))).length))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-123)), function (_1_i0) {
  return new BigNumber(65);
})).length),!(true))).length))).length), true))))]);
          }
        }
        return _coll0;
      }()).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(944),_module.D0.create_DC3(_module.D0.create_DC2(false, new BigNumber(754), new BigNumber(561), false)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-113))).cardinality()),_module.D0.create_DC3(_module.D0.create_DC3(_module.D0.create_DC2(true, new BigNumber(111), new BigNumber(-539), !(true)))))));
    };
    static fm2(p0, globalState) {
      return _dafny.Seq.of(_module.__default.safeDivisionInt(new BigNumber(466), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('v'.codePointAt(0)),new BigNumber(696))).length)));
    };
    static fm3(p0, globalState) {
      return (_module.D3.create_DC6(_dafny.Seq.UnicodeFromString("xkxmuj"))).dtor_cf10;
    };
    static fm5(globalState) {
      return new _dafny.CodePoint('n'.codePointAt(0));
    };
    static fm6(p0, p1, globalState) {
      return _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(false, true), _dafny.Seq.of(!(false))), _dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of(false))));
    };
    static fm7(globalState) {
      return ((_dafny.Set.fromElements(!(true))).Union(_dafny.Set.fromElements(false))).Difference(_dafny.Set.fromElements(true));
    };
    static fm10(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(true, !(!(false)), false), _dafny.Seq.of(true, false, false, !(false))), _dafny.Seq.of(!(true)));
    };
    static fm11(p0, p1, globalState) {
      return (_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("nw"), _dafny.Seq.UnicodeFromString("fbvst"), _dafny.Seq.UnicodeFromString("dmftxyldf"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(111)), function (_2_i1) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      }))).IsSubsetOf((_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-760)), function (_3_i0) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      })))).Difference(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("hgiqbg"), _dafny.Seq.UnicodeFromString("gxnfl"))));
    };
    static fm12(globalState) {
      if (((!(false)) ? (true) : (false))) {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-835),new BigNumber((_dafny.Seq.of(!(false), !(false))).length));
      } else {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(156),new BigNumber((_dafny.MultiSet.fromElements(!(true), !(true))).cardinality()));
      }
    };
    static fm13(p0, p1, globalState) {
      return _dafny.Set.fromElements(new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)));
    };
    static m2(p0, p1, globalState) {
      if (p0) {
        let _4_v0;
        _4_v0 = new BigNumber(274);
        _4_v0 = _4_v0;
        let _5_v1;
        _5_v1 = _dafny.Set.fromElements(_4_v0);
        if (((_5_v1).IsProperSubsetOf(_5_v1)) || (!(!(p0)) || (p0))) {
          let _6_v2;
          _6_v2 = new _dafny.CodePoint('e'.codePointAt(0));
          (globalState).f18 = _6_v2;
          let _7_v3;
          _7_v3 = _dafny.Seq.of(_dafny.Seq.of(p0));
          let _8_v4;
          let _nw0 = new _module.C0();
          _nw0.__ctor(!(p0), _dafny.Map.Empty.slice().updateUnsafe(_4_v0,(_7_v3)[_module.__default.safeIndex(_4_v0, new BigNumber((_7_v3).length))]));
          _8_v4 = _nw0;
          let _9_v5;
          _9_v5 = _dafny.Set.fromElements(_8_v4);
          (globalState).f19 = !((_9_v5).IsProperSubsetOf(_9_v5));
          let _10_v6;
          _10_v6 = _dafny.MultiSet.fromElements(_8_v4.f24, p1);
          let _11_v7;
          let _nw1 = Array((new BigNumber(22)).toNumber());
          _nw1[(_dafny.ZERO).toNumber()] = p1;
          _nw1[(_dafny.ONE).toNumber()] = true;
          _nw1[(new BigNumber(2)).toNumber()] = p0;
          _nw1[(new BigNumber(3)).toNumber()] = !(true);
          _nw1[(new BigNumber(4)).toNumber()] = p1;
          _nw1[(new BigNumber(5)).toNumber()] = _8_v4.f24;
          _nw1[(new BigNumber(6)).toNumber()] = p0;
          _nw1[(new BigNumber(7)).toNumber()] = !(p0);
          _nw1[(new BigNumber(8)).toNumber()] = false;
          _nw1[(new BigNumber(9)).toNumber()] = p1;
          _nw1[(new BigNumber(10)).toNumber()] = p0;
          _nw1[(new BigNumber(11)).toNumber()] = _8_v4.f24;
          _nw1[(new BigNumber(12)).toNumber()] = _module.__default.fm11(_10_v6, _4_v0, globalState);
          _nw1[(new BigNumber(13)).toNumber()] = p0;
          _nw1[(new BigNumber(14)).toNumber()] = p0;
          _nw1[(new BigNumber(15)).toNumber()] = false;
          _nw1[(new BigNumber(16)).toNumber()] = p0;
          _nw1[(new BigNumber(17)).toNumber()] = _8_v4.f24;
          _nw1[(new BigNumber(18)).toNumber()] = _8_v4.f24;
          _nw1[(new BigNumber(19)).toNumber()] = p0;
          _nw1[(new BigNumber(20)).toNumber()] = p0;
          _nw1[(new BigNumber(21)).toNumber()] = p1;
          _11_v7 = _nw1;
          let _12_v8;
          _12_v8 = _dafny.Map.Empty.slice().updateUnsafe(_11_v7,_11_v7);
          _12_v8 = (_12_v8).update(_11_v7, _11_v7);
          let _13_v9;
          _13_v9 = _dafny.Seq.of(p1, p0);
          let _14_v10;
          _14_v10 = _dafny.Seq.of(new BigNumber(-330), new BigNumber((_13_v9).length));
          let _15_v11;
          _15_v11 = _dafny.Set.fromElements(_14_v10, _dafny.Seq.of((_dafny.ZERO).minus(_4_v0)), _dafny.Seq.of(_4_v0), _14_v10, _dafny.Seq.of(_4_v0, _4_v0, _4_v0));
          let _index0 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_11_v7).length));
          (_11_v7)[_index0] = (_15_v11).IsProperSubsetOf(_15_v11);
          let _16_v12;
          _16_v12 = _dafny.Seq.UnicodeFromString("qlttakuyp");
          let _index1 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_11_v7).length));
          let _rhs0 = p1;
          let _rhs1 = _dafny.Seq.Concat(_dafny.Seq.Concat(_16_v12, _16_v12), _dafny.Seq.UnicodeFromString("tfdjcpnpq"));
          let _rhs2 = (new BigNumber(575)).multipliedBy(_4_v0);
          let _rhs3 = _8_v4;
          let _lhs0 = _11_v7;
          let _lhs1 = _module.__default.safeIndex(new BigNumber(790), new BigNumber((_11_v7).length));
          _lhs0[_lhs1] = _rhs0;
          _16_v12 = _rhs1;
          _4_v0 = _rhs2;
          _8_v4 = _rhs3;
          _8_v4 = (((_8_v4.f24) || ((_11_v7)[_module.__default.safeIndex(new BigNumber(790), new BigNumber((_11_v7).length))])) ? (_8_v4) : (_8_v4));
        } else {
          let _17_v13;
          _17_v13 = _dafny.Seq.of(p1);
          let _18_v14;
          _18_v14 = _dafny.Map.Empty.slice().updateUnsafe(_4_v0,_17_v13);
          let _19_v15;
          let _nw2 = new _module.C0();
          _nw2.__ctor(!(p1) || (p0), _18_v14);
          _19_v15 = _nw2;
          let _20_v16;
          _20_v16 = _dafny.Seq.UnicodeFromString("rqabmwu");
          let _21_v17;
          _21_v17 = new _dafny.CodePoint('u'.codePointAt(0));
          let _22_v18;
          _22_v18 = _dafny.MultiSet.fromElements(p0, p1);
          let _23_v19;
          _23_v19 = _dafny.Map.Empty.slice().updateUnsafe(_21_v17,new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(p0,_4_v0)).update(_module.__default.fm11(_22_v18, new BigNumber(870), globalState), _4_v0)).length));
          _17_v13 = _dafny.Seq.update(_dafny.Seq.Concat(_module.__default.fm10(new BigNumber((_20_v16).length), new BigNumber(751), globalState), _dafny.Seq.Concat(_17_v13, _17_v13)), _module.__default.safeIndex(new BigNumber((_23_v19).length), new BigNumber((_dafny.Seq.Concat(_module.__default.fm10(new BigNumber((_20_v16).length), new BigNumber(751), globalState), _dafny.Seq.Concat(_17_v13, _17_v13))).length)), !(p0));
          _4_v0 = new BigNumber(825);
          let _24_v20;
          let _nw3 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
          _24_v20 = _nw3;
          let _index2 = _module.__default.safeIndex(new BigNumber(390), new BigNumber((_24_v20).length));
          (_24_v20)[_index2] = _4_v0;
          let _index3 = _module.__default.safeIndex(new BigNumber(390), new BigNumber((_24_v20).length));
          (_24_v20)[_index3] = (((p1) ? (_4_v0) : (_4_v0))).multipliedBy(_4_v0);
          (globalState).f19 = (_17_v13)[_module.__default.safeIndex(_4_v0, new BigNumber((_17_v13).length))];
        }
        (globalState).f6 = p1;
        (globalState).f6 = p1;
        let _25_v21;
        _25_v21 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        let _26_v22;
        _26_v22 = _dafny.Map.Empty.slice().updateUnsafe(_4_v0,(((_25_v21).contains(false)) ? ((_25_v21).get(false)) : (p0)));
        let _27_v23;
        _27_v23 = _dafny.Seq.of(p0, p1, p0, p0);
        (globalState).f19 = ((!((((_26_v22).contains((_dafny.ZERO).minus((_dafny.ZERO).minus(_4_v0)))) ? ((_26_v22).get((_dafny.ZERO).minus((_dafny.ZERO).minus(_4_v0)))) : (false)))) ? ((_27_v23)[_module.__default.safeIndex(_module.__default.fm0(p0, p1, globalState), new BigNumber((_27_v23).length))]) : (p1));
      } else {
        let _28_v24;
        _28_v24 = new BigNumber(191);
        let _29_v25;
        _29_v25 = _dafny.Seq.UnicodeFromString("bvmshmd");
        _28_v24 = (new BigNumber((_29_v25).length)).plus(_28_v24);
        let _30_v26;
        let _nw4 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Map.Empty);
        _30_v26 = _nw4;
        _30_v26 = _30_v26;
        let _31_v27;
        _31_v27 = _dafny.Seq.of(false, p0);
        let _32_v28;
        _32_v28 = _module.D0.create_DC1(_31_v27, new BigNumber(-405));
        let _33_v29;
        _33_v29 = _dafny.Map.Empty.slice().updateUnsafe(_28_v24,((p0) ? (_32_v28) : (_32_v28)));
        let _34_v30;
        _34_v30 = _dafny.Map.Empty.slice().updateUnsafe(_28_v24,p1);
        let _35_v31;
        _35_v31 = _dafny.MultiSet.fromElements(_module.__default.fm0(p0, p0, globalState), new BigNumber((_34_v30).length));
        let _36_v32;
        _36_v32 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(p1)).length));
        _33_v29 = (_33_v29).update((((_35_v31).contains(_28_v24)) ? ((_35_v31).get(_28_v24)) : (new BigNumber((_36_v32).length))), _32_v28);
        let _37_v33;
        let _nw5 = Array((new BigNumber(25)).toNumber()).fill(_dafny.MultiSet.Empty);
        _37_v33 = _nw5;
        _37_v33 = _37_v33;
        let _38_v34;
        let _nw6 = Array((new BigNumber(3)).toNumber()).fill(false);
        _38_v34 = _nw6;
        let _39_v35;
        _39_v35 = _dafny.MultiSet.fromElements(_38_v34, _38_v34);
        let _40_v36;
        _40_v36 = ((!(p1)) ? (_39_v35) : (_39_v35));
        let _source0 = _40_v36;
        let _41___mcc_h0 = _source0;
        let _42_cf18 = _41___mcc_h0;
        let _43_v37;
        _43_v37 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(-759));
        let _44_v38;
        _44_v38 = _dafny.Seq.of(_38_v34, _38_v34);
        let _45_v39;
        _45_v39 = _module.D3.create_DC7(_43_v37, p0, p1, p1, (_44_v38)[_module.__default.safeIndex((_32_v28).dtor_cf2, new BigNumber((_44_v38).length))]);
        let _46_v40;
        let _nw7 = new _module.C1();
        _nw7.__ctor(p1, (_45_v39).dtor_cf14);
        _46_v40 = _nw7;
        let _47_v41;
        _47_v41 = new _dafny.CodePoint('t'.codePointAt(0));
        (globalState).f18 = _47_v41;
        (globalState).f19 = !((_46_v40).fm4(new BigNumber(373), new BigNumber(230), _46_v40.f23, globalState)) || (((p1) ? (p0) : ((_45_v39).dtor_cf12)));
        (globalState).f19 = (_46_v40).f22;
      }
      let _48_v42;
      let _nw8 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
      _48_v42 = _nw8;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_48_v42).length))) {
        let _49_i0 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_49_i0)) && ((_49_i0).isLessThan(new BigNumber((_48_v42).length))))) {
          (_48_v42)[(_49_i0)] = _module.__default.safeDivisionInt(_49_i0, new BigNumber(-822));
        }
      }
      let _50_i1;
      _50_i1 = _dafny.ZERO;
      L0: {
        while (!(p0)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_50_i1)) {
              break L0;
            }
            _50_i1 = (_50_i1).plus(_dafny.ONE);
            let _51_v43;
            _51_v43 = new BigNumber(493);
            _51_v43 = _51_v43;
            let _52_v44;
            _52_v44 = _dafny.MultiSet.fromElements(p1);
            let _53_v45;
            let _nw9 = new _module.C1();
            _nw9.__ctor(true, p1);
            _53_v45 = _nw9;
            let _54_v46;
            _54_v46 = _dafny.MultiSet.fromElements(_53_v45);
            let _55_v47;
            _55_v47 = _module.D0.create_DC2(p0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(p1),(_module.D0.create_DC2(_module.__default.fm11(_52_v44, _51_v43, globalState), _51_v43, _51_v43, true)).dtor_cf3)).length), new BigNumber((_54_v46).cardinality()), p0);
            let _56_v48;
            _56_v48 = _dafny.Set.fromElements((_55_v47).dtor_cf3, p1, (_53_v45).f22, _53_v45.f23);
            let _57_v49;
            _57_v49 = _dafny.Map.Empty.slice().updateUnsafe(_53_v45.f23,p0);
            let _58_v50;
            _58_v50 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_56_v48).length),(((_57_v49).contains(_53_v45.f23)) ? ((_57_v49).get(_53_v45.f23)) : (_53_v45.f23)));
            let _59_v51;
            _59_v51 = _dafny.Map.Empty.slice().updateUnsafe(_58_v50,p0);
            _59_v51 = (_59_v51).update(_58_v50, (_53_v45).f22);
            let _60_v52;
            _60_v52 = new _dafny.CodePoint('q'.codePointAt(0));
            let _61_v53;
            _61_v53 = _dafny.Seq.of(_60_v52, _60_v52);
            let _62_v54;
            _62_v54 = _dafny.Map.Empty.slice().updateUnsafe(_53_v45.f23,_61_v53);
            let _63_v55;
            let _nw10 = Array((_dafny.ONE).toNumber());
            _nw10[(_dafny.ZERO).toNumber()] = _module.__default.fm2(_dafny.MultiSet.FromArray((((_62_v54).contains(_53_v45.f23)) ? ((_62_v54).get(_53_v45.f23)) : (_dafny.Seq.of(_60_v52)))), globalState);
            _63_v55 = _nw10;
            let _64_v56;
            _64_v56 = _dafny.Set.fromElements(_module.__default.fm5(globalState));
            let _65_v57;
            _65_v57 = _dafny.Seq.of(_64_v56, _64_v56, _module.__default.fm13(new BigNumber((_61_v53).length), _51_v43, globalState));
            let _66_v58;
            _66_v58 = _dafny.Seq.of(new BigNumber((_65_v57).length));
            let _index4 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_63_v55).length));
            (_63_v55)[_index4] = _66_v58;
            let _67_v59;
            _67_v59 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('m'.codePointAt(0)));
            let _68_v60;
            _68_v60 = _dafny.Seq.of(_dafny.Seq.update(_66_v58, _module.__default.safeIndex(new BigNumber(756), new BigNumber((_66_v58).length)), _51_v43), _dafny.Seq.Create(_module.__default.abs(new BigNumber(3)), function (_69_i2) {
              return new BigNumber(865);
            }));
            let _index5 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_63_v55).length));
            (_63_v55)[_index5] = _dafny.Seq.Concat(_module.__default.fm2(_67_v59, globalState), (_68_v60)[_module.__default.safeIndex(_51_v43, new BigNumber((_68_v60).length))]);
            if ((_53_v45).f22) {
              let _70_v61;
              _70_v61 = _dafny.Map.Empty.slice().updateUnsafe(p1,_51_v43);
              let _71_v62;
              _71_v62 = _module.D3.create_DC8((((_70_v61).contains((_53_v45).f22)) ? ((_70_v61).get((_53_v45).f22)) : (_51_v43)), (((_53_v45).f22) ? (_51_v43) : (_51_v43)));
              let _pat_let_tv0 = _51_v43;
              _71_v62 = (((_53_v45).f22) ? (_71_v62) : (function (_pat_let0_0) {
                return function (_72_dt__update__tmp_h0) {
                  return function (_pat_let1_0) {
                    return function (_73_dt__update_hcf17_h0) {
                      return _module.D3.create_DC8((_72_dt__update__tmp_h0).dtor_cf16, _73_dt__update_hcf17_h0);
                    }(_pat_let1_0);
                  }(_pat_let_tv0);
                }(_pat_let0_0);
              }(_71_v62)));
              let _74_v64;
              let _nw11 = new _module.C0();
              _nw11.__ctor(p1, function () {
                let _coll1 = new _dafny.Map();
                for (const _compr_1 of _dafny.IntegerRange(new BigNumber(-478), new BigNumber(307))) {
                  let _75_v63 = _compr_1;
                  if (((new BigNumber(-478)).isLessThanOrEqualTo(_75_v63)) && ((_75_v63).isLessThan(new BigNumber(307)))) {
                    _coll1.push([_module.__default.safeModuloInt(_75_v63, _51_v43),_dafny.Seq.of(false)]);
                  }
                }
                return _coll1;
              }());
              _74_v64 = _nw11;
              let _76_v65;
              let _nw12 = Array((new BigNumber(4)).toNumber());
              _nw12[(_dafny.ZERO).toNumber()] = _74_v64;
              _nw12[(_dafny.ONE).toNumber()] = _74_v64;
              _nw12[(new BigNumber(2)).toNumber()] = _74_v64;
              _nw12[(new BigNumber(3)).toNumber()] = _74_v64;
              _76_v65 = _nw12;
              _76_v65 = _76_v65;
              let _77_v66;
              let _nw13 = new _module.C1();
              _nw13.__ctor(!(false), p1);
              _77_v66 = _nw13;
              let _rhs4 = _60_v52;
              let _rhs5 = _51_v43;
              let _lhs2 = globalState;
              _lhs2.f18 = _rhs4;
              _51_v43 = _rhs5;
              let _78_v67;
              let _nw14 = Array((new BigNumber(10)).toNumber()).fill(false);
              _78_v67 = _nw14;
              _78_v67 = _78_v67;
            } else {
              _51_v43 = new BigNumber((_dafny.Seq.Concat(_61_v53, _dafny.Seq.UnicodeFromString("q"))).length);
              let _79_v68;
              _79_v68 = _dafny.Seq.of(p0, p0);
              let _80_v69;
              _80_v69 = _dafny.Map.Empty.slice().updateUnsafe(_79_v68,_53_v45.f23);
              _51_v43 = ((_53_v45.f23) ? (_module.__default.safeModuloInt(_51_v43, _51_v43)) : (new BigNumber((_80_v69).length)));
              let _81_v70;
              let _nw15 = Array((new BigNumber(26)).toNumber()).fill(false);
              _81_v70 = _nw15;
              let _index6 = _module.__default.safeIndex(new BigNumber(65), new BigNumber((_81_v70).length));
              (_81_v70)[_index6] = p0;
              let _index7 = _module.__default.safeIndex(new BigNumber(65), new BigNumber((_81_v70).length));
              (_81_v70)[_index7] = ((new BigNumber(-514)).isLessThanOrEqualTo(new BigNumber(861))) || ((_53_v45).f22);
              let _82_v71;
              let _init0 = ((_83_v52) => function (_84_i3) {
                return _83_v52;
              })(_60_v52);
              let _nw16 = Array((new BigNumber(22)).toNumber());
              for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw16.length); _i0_0++) {
                _nw16[_i0_0] = _init0(new BigNumber(_i0_0));
              }
              _82_v71 = _nw16;
              let _index8 = _module.__default.safeIndex(new BigNumber(691), new BigNumber((_82_v71).length));
              (_82_v71)[_index8] = _60_v52;
              let _index9 = _module.__default.safeIndex(new BigNumber(691), new BigNumber((_82_v71).length));
              (_82_v71)[_index9] = _60_v52;
              let _85_v72;
              let _nw17 = Array((new BigNumber(4)).toNumber()).fill([]);
              _85_v72 = _nw17;
              let _86_v73;
              let _nw18 = Array((_dafny.ONE).toNumber());
              _nw18[(_dafny.ZERO).toNumber()] = _85_v72;
              _86_v73 = _nw18;
              let _index10 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_86_v73).length));
              (_86_v73)[_index10] = _85_v72;
              let _index11 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_86_v73).length));
              let _rhs6 = _85_v72;
              let _rhs7 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(((p0) ? (_79_v68) : (_dafny.Seq.of(p0, (_53_v45).f22, (_53_v45).f22))), _79_v68)).length));
              let _rhs8 = _dafny.areEqual((_63_v55)[_module.__default.safeIndex(new BigNumber(816), new BigNumber((_63_v55).length))], _66_v58);
              let _lhs3 = _86_v73;
              let _lhs4 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_86_v73).length));
              let _lhs5 = _53_v45;
              _lhs3[_lhs4] = _rhs6;
              _51_v43 = _rhs7;
              _lhs5.f23 = _rhs8;
            }
          }
        }
      }
      let _87_i4;
      _87_i4 = _dafny.ZERO;
      L1: {
        while (p0) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_87_i4)) {
              break L1;
            }
            _87_i4 = (_87_i4).plus(_dafny.ONE);
            let _88_v74;
            _88_v74 = new BigNumber(120);
            let _index12 = _module.__default.safeIndex(new BigNumber(926), new BigNumber((_48_v42).length));
            (_48_v42)[_index12] = _88_v74;
            let _index13 = _module.__default.safeIndex(new BigNumber(926), new BigNumber((_48_v42).length));
            (_48_v42)[_index13] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_88_v74, new BigNumber(-138)));
            let _89_v75;
            _89_v75 = _dafny.Seq.UnicodeFromString("nqxh");
            let _90_v76;
            _90_v76 = _dafny.Map.Empty.slice().updateUnsafe(p0,_89_v75);
            let _91_v77;
            _91_v77 = _dafny.Map.Empty.slice().updateUnsafe(_88_v74,(_48_v42)[_module.__default.safeIndex(new BigNumber(926), new BigNumber((_48_v42).length))]);
            let _92_v78;
            _92_v78 = new _dafny.CodePoint('h'.codePointAt(0));
            let _93_v79;
            _93_v79 = _module.D5.create_DC10(_92_v78);
            let _94_v80;
            _94_v80 = _dafny.Seq.of(p0);
            let _95_v81;
            let _nw19 = new _module.C0();
            _nw19.__ctor(p0, _dafny.Map.Empty.slice().updateUnsafe((_48_v42)[_module.__default.safeIndex(new BigNumber(926), new BigNumber((_48_v42).length))],_94_v80));
            _95_v81 = _nw19;
            let _96_v82;
            _96_v82 = _dafny.Seq.of(_95_v81);
            let _97_v83;
            let _nw20 = Array((new BigNumber(24)).toNumber());
            _nw20[(_dafny.ZERO).toNumber()] = new BigNumber(-478);
            _nw20[(_dafny.ONE).toNumber()] = new BigNumber((_91_v77).length);
            _nw20[(new BigNumber(2)).toNumber()] = _88_v74;
            _nw20[(new BigNumber(3)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,(_93_v79).dtor_cf19)).length);
            _nw20[(new BigNumber(4)).toNumber()] = (_48_v42)[_module.__default.safeIndex(new BigNumber(926), new BigNumber((_48_v42).length))];
            _nw20[(new BigNumber(5)).toNumber()] = (_48_v42)[_module.__default.safeIndex(new BigNumber(926), new BigNumber((_48_v42).length))];
            _nw20[(new BigNumber(6)).toNumber()] = (_48_v42)[_module.__default.safeIndex(new BigNumber(926), new BigNumber((_48_v42).length))];
            _nw20[(new BigNumber(7)).toNumber()] = (_48_v42)[_module.__default.safeIndex(new BigNumber(926), new BigNumber((_48_v42).length))];
            _nw20[(new BigNumber(8)).toNumber()] = _88_v74;
            _nw20[(new BigNumber(9)).toNumber()] = _88_v74;
            _nw20[(new BigNumber(10)).toNumber()] = (_48_v42)[_module.__default.safeIndex(new BigNumber(926), new BigNumber((_48_v42).length))];
            _nw20[(new BigNumber(11)).toNumber()] = new BigNumber((_96_v82).length);
            _nw20[(new BigNumber(12)).toNumber()] = _88_v74;
            _nw20[(new BigNumber(13)).toNumber()] = new BigNumber((_89_v75).length);
            _nw20[(new BigNumber(14)).toNumber()] = _88_v74;
            _nw20[(new BigNumber(15)).toNumber()] = _88_v74;
            _nw20[(new BigNumber(16)).toNumber()] = _88_v74;
            _nw20[(new BigNumber(17)).toNumber()] = _88_v74;
            _nw20[(new BigNumber(18)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_48_v42)[_module.__default.safeIndex(new BigNumber(926), new BigNumber((_48_v42).length))],p1)).length);
            _nw20[(new BigNumber(19)).toNumber()] = _module.__default.fm0(_95_v81.f24, p0, globalState);
            _nw20[(new BigNumber(20)).toNumber()] = (_48_v42)[_module.__default.safeIndex(new BigNumber(926), new BigNumber((_48_v42).length))];
            _nw20[(new BigNumber(21)).toNumber()] = _88_v74;
            _nw20[(new BigNumber(22)).toNumber()] = _88_v74;
            _nw20[(new BigNumber(23)).toNumber()] = (_48_v42)[_module.__default.safeIndex(new BigNumber(926), new BigNumber((_48_v42).length))];
            _97_v83 = _nw20;
            let _98_v84;
            _98_v84 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("no"),_97_v83);
            let _99_v85;
            _99_v85 = _dafny.Set.fromElements(_88_v74, (_48_v42)[_module.__default.safeIndex(new BigNumber(926), new BigNumber((_48_v42).length))], (_48_v42)[_module.__default.safeIndex(new BigNumber(926), new BigNumber((_48_v42).length))], new BigNumber((_98_v84).length), new BigNumber((_94_v80).length));
            let _100_v86;
            _100_v86 = _dafny.Seq.of(new BigNumber((_99_v85).length), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(4)), ((_101_v75) => function (_102_i5) {
              return (_101_v75)[_module.__default.safeIndex(new BigNumber(293), new BigNumber((_101_v75).length))];
            })(_89_v75))).length), (_dafny.ZERO).minus(new BigNumber((_89_v75).length)))).length));
            let _103_v87;
            let _nw21 = new _module.C0();
            _nw21.__ctor(_95_v81.f24, (_95_v81).f25);
            _103_v87 = _nw21;
            let _104_v88;
            _104_v88 = _dafny.Map.Empty.slice().updateUnsafe((_48_v42)[_module.__default.safeIndex(new BigNumber(926), new BigNumber((_48_v42).length))],_103_v87);
            let _index14 = _module.__default.safeIndex(new BigNumber(926), new BigNumber((_48_v42).length));
            (_48_v42)[_index14] = _module.__default.safeDivisionInt((new BigNumber(((((_90_v76).contains(p0)) ? ((_90_v76).get(p0)) : (_dafny.Seq.UnicodeFromString("ovddoibvv")))).length)).multipliedBy(new BigNumber((_dafny.Seq.update(_100_v86, _module.__default.safeIndex(new BigNumber((_104_v88).length), new BigNumber((_100_v86).length)), _88_v74)).length)), (_48_v42)[_module.__default.safeIndex(new BigNumber(926), new BigNumber((_48_v42).length))]);
            let _105_v89;
            let _nw22 = Array((_dafny.ONE).toNumber()).fill(false);
            _105_v89 = _nw22;
            let _index15 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_105_v89).length));
            (_105_v89)[_index15] = !(p1);
            let _index16 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_105_v89).length));
            (_105_v89)[_index16] = _95_v81.f24;
            let _index17 = _module.__default.safeIndex(new BigNumber(926), new BigNumber((_48_v42).length));
            let _rhs9 = _103_v87;
            let _rhs10 = _95_v81.f24;
            let _rhs11 = (_48_v42)[_module.__default.safeIndex(new BigNumber(926), new BigNumber((_48_v42).length))];
            let _rhs12 = ((_dafny.ZERO).minus(_88_v74)).plus((_48_v42)[_module.__default.safeIndex(new BigNumber(926), new BigNumber((_48_v42).length))]);
            let _lhs6 = globalState;
            let _lhs7 = _48_v42;
            let _lhs8 = _module.__default.safeIndex(new BigNumber(926), new BigNumber((_48_v42).length));
            _103_v87 = _rhs9;
            _lhs6.f19 = _rhs10;
            _lhs7[_lhs8] = _rhs11;
            _88_v74 = _rhs12;
          }
        }
      }
      let _106_v90;
      _106_v90 = new BigNumber(568);
      _106_v90 = _module.__default.safeModuloInt(new BigNumber(-273), _106_v90);
      let _107_v91;
      _107_v91 = _dafny.Seq.UnicodeFromString("obf");
      let _108_v92;
      let _nw23 = Array((new BigNumber(8)).toNumber()).fill(false);
      _108_v92 = _nw23;
      let _109_v93;
      _109_v93 = _dafny.MultiSet.fromElements(p1);
      let _110_v94;
      _110_v94 = _dafny.Seq.of(_module.__default.fm11(_109_v93, _106_v90, globalState));
      let _111_v95;
      let _nw24 = new _module.C0();
      _nw24.__ctor(p1, _dafny.Map.Empty.slice().updateUnsafe(_106_v90,_110_v94));
      _111_v95 = _nw24;
      let _112_v96;
      _112_v96 = _module.D5.create_DC11(p1, p0, _108_v92, _111_v95);
      let _113_v97;
      _113_v97 = _dafny.Seq.of(_112_v96, _module.D5.create_DC11(p1, p0, _108_v92, _111_v95), _112_v96, _module.D5.create_DC11(p0, p1, _108_v92, _111_v95));
      let _114_v98;
      _114_v98 = _dafny.Map.Empty.slice().updateUnsafe(_107_v91,_113_v97);
      let _115_v99;
      let _nw25 = Array((new BigNumber(7)).toNumber()).fill([]);
      _115_v99 = _nw25;
      let _index18 = _module.__default.safeIndex(new BigNumber(278), new BigNumber((_115_v99).length));
      (_115_v99)[_index18] = _48_v42;
      let _116_v100;
      _116_v100 = _48_v42;
      let _index19 = _module.__default.safeIndex(new BigNumber(278), new BigNumber((_115_v99).length));
      let _rhs13 = (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(423)), function (_117_i6) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      }),_113_v97)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("rntowef"),_113_v97));
      let _rhs14 = (_116_v100);
      let _lhs9 = _115_v99;
      let _lhs10 = _module.__default.safeIndex(new BigNumber(278), new BigNumber((_115_v99).length));
      _114_v98 = _rhs13;
      _lhs9[_lhs10] = _rhs14;
      return;
    }
    static Main(__noArgsParameter) {
      let _118_v0;
      _118_v0 = true;
      let _119_v1;
      _119_v1 = _dafny.Seq.of(_118_v0);
      let _120_v2;
      _120_v2 = _dafny.Seq.UnicodeFromString("pvnnsysbg");
      let _121_v3;
      _121_v3 = _module.D0.create_DC1(_119_v1, new BigNumber((_120_v2).length));
      let _122_v4;
      _122_v4 = _dafny.Map.Empty.slice().updateUnsafe(_118_v0,(_121_v3).dtor_cf2);
      let _123_v5;
      _123_v5 = _dafny.Map.Empty.slice().updateUnsafe(_121_v3,_dafny.MultiSet.fromElements(_118_v0));
      let _124_v6;
      _124_v6 = _dafny.Set.fromElements(_118_v0);
      let _125_v7;
      _125_v7 = _dafny.MultiSet.fromElements(_118_v0);
      let _126_v8;
      _126_v8 = _dafny.Seq.of(_120_v2, _120_v2, _120_v2);
      let _127_v9;
      let _init1 = ((_128_v2) => function (_129_i0) {
        return _module.__default.safeModuloInt(_129_i0, new BigNumber((_128_v2).length));
      })(_120_v2);
      let _nw26 = Array((new BigNumber(19)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw26.length); _i0_1++) {
        _nw26[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _127_v9 = _nw26;
      let _130_v11;
      _130_v11 = new _dafny.CodePoint('j'.codePointAt(0));
      let _131_v12;
      _131_v12 = _dafny.Set.fromElements(_130_v11);
      let _132_v13;
      _132_v13 = _dafny.Map.Empty.slice().updateUnsafe(_125_v7,_118_v0);
      let _133_v14;
      _133_v14 = new BigNumber(588);
      let _134_v15;
      _134_v15 = _dafny.Map.Empty.slice().updateUnsafe(_130_v11,(_dafny.Map.Empty.slice().updateUnsafe(_118_v0,_133_v14)).update(_118_v0, _133_v14));
      let _135_globalState;
      let _nw27 = new _module.GlobalState();
      _nw27.__ctor(_122_v4, true, true, (((_123_v5).contains(_module.D0.create_DC1(_119_v1, new BigNumber((_124_v6).length)))) ? ((_123_v5).get(_module.D0.create_DC1(_119_v1, new BigNumber((_124_v6).length)))) : (_125_v7)), true, new BigNumber(869), false, new BigNumber(924), false, new BigNumber(-924), _dafny.Seq.Concat(_126_v8, _126_v8), _120_v2, new BigNumber(380), _127_v9, new BigNumber(486), function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of _dafny.IntegerRange(new BigNumber(729), new BigNumber(212))) {
          let _136_v10 = _compr_2;
          if (((new BigNumber(729)).isLessThanOrEqualTo(_136_v10)) && ((_136_v10).isLessThan(new BigNumber(212)))) {
            _coll2.push([_module.__default.safeDivisionInt(_136_v10, new BigNumber(-701)),_dafny.Seq.of(new BigNumber(-6), (_dafny.ZERO).minus(new BigNumber(-208)))]);
          }
        }
        return _coll2;
      }(), _131_v12, (_132_v13).Merge(_132_v13), new _dafny.CodePoint('y'.codePointAt(0)), false, _134_v15, new BigNumber(964));
      _135_globalState = _nw27;
      let _137_v16;
      _137_v16 = _dafny.Set.fromElements(_119_v1, _119_v1, _dafny.Seq.Concat(_119_v1, _119_v1));
      _137_v16 = (_137_v16).Union((_137_v16).Union(_137_v16));
      let _138_v17;
      _138_v17 = _dafny.Map.Empty.slice().updateUnsafe(((_118_v0) ? (_118_v0) : (!(_118_v0))),_dafny.Seq.Create(_module.__default.abs(new BigNumber(243)), ((_139_v14) => function (_140_i1) {
        return (_dafny.ZERO).minus(_139_v14);
      })(_133_v14)));
      let _141_v18;
      _141_v18 = _dafny.Seq.of(_125_v7);
      let _142_v19;
      _142_v19 = _dafny.Seq.of(_133_v14, new BigNumber(((_141_v18)[_module.__default.safeIndex(_133_v14, new BigNumber((_141_v18).length))]).cardinality()));
      _138_v17 = (_138_v17).update((_118_v0) || (!(_118_v0)), _142_v19);
      let _143_v20;
      let _nw28 = Array((new BigNumber(27)).toNumber()).fill(false);
      _143_v20 = _nw28;
      _143_v20 = _143_v20;
      let _144_v21;
      _144_v21 = _module.D0.create_DC2(_118_v0, _133_v14, new BigNumber((_122_v4).length), _118_v0);
      let _145_v22;
      _145_v22 = _dafny.Map.Empty.slice().updateUnsafe(_133_v14,_module.D0.create_DC3(_144_v21));
      let _146_v23;
      let _nw29 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Map.Empty);
      _146_v23 = _nw29;
      let _147_v24;
      _147_v24 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_130_v11),_module.__default.fm0(_118_v0, _118_v0, _135_globalState));
      let _index20 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_146_v23).length));
      (_146_v23)[_index20] = _147_v24;
      let _index21 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length));
      (_127_v9)[_index21] = _133_v14;
      let _148_v25;
      _148_v25 = _dafny.Set.fromElements(new BigNumber((_120_v2).length));
      let _149_v26;
      _149_v26 = _dafny.MultiSet.fromElements(_130_v11);
      let _index22 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_146_v23).length));
      let _index23 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length));
      let _rhs15 = _dafny.Seq.Concat(_120_v2, _120_v2);
      let _rhs16 = _module.__default.fm1(_118_v0, _118_v0, _module.__default.safeModuloInt(new BigNumber((_148_v25).length), _133_v14), _133_v14, _135_globalState);
      let _rhs17 = _dafny.Map.Empty.slice().updateUnsafe(_149_v26,(_133_v14).plus(_133_v14));
      let _rhs18 = _124_v6;
      let _rhs19 = (_dafny.ZERO).minus((_module.__default.safeModuloInt(_133_v14, _133_v14)).minus(_133_v14));
      let _lhs11 = _146_v23;
      let _lhs12 = _module.__default.safeIndex(new BigNumber(746), new BigNumber((_146_v23).length));
      let _lhs13 = _127_v9;
      let _lhs14 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length));
      _120_v2 = _rhs15;
      _145_v22 = _rhs16;
      _lhs11[_lhs12] = _rhs17;
      _124_v6 = _rhs18;
      _lhs13[_lhs14] = _rhs19;
      _142_v19 = _142_v19;
      _120_v2 = _dafny.Seq.Concat(_120_v2, _dafny.Seq.update(_120_v2, _module.__default.safeIndex(_133_v14, new BigNumber((_120_v2).length)), _130_v11));
      let _index24 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length));
      let _rhs20 = _121_v3;
      let _rhs21 = (_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))];
      let _lhs15 = _127_v9;
      let _lhs16 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length));
      _121_v3 = _rhs20;
      _lhs15[_lhs16] = _rhs21;
      let _150_i2;
      _150_i2 = _dafny.ZERO;
      L2: {
        while (!(_118_v0)) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_150_i2)) {
              break L2;
            }
            _150_i2 = (_150_i2).plus(_dafny.ONE);
            if (!_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(138)), ((_151_v11) => function (_152_i3) {
              return _151_v11;
            })(_130_v11)), _dafny.Seq.Concat(_120_v2, _120_v2))) {
              _142_v19 = _module.__default.fm2((_dafny.MultiSet.FromArray(_120_v2)).Intersect(_149_v26), _135_globalState);
              _120_v2 = _module.__default.fm3(_133_v14, _135_globalState);
              let _index25 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_143_v20).length));
              (_143_v20)[_index25] = _118_v0;
              let _index26 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_143_v20).length));
              (_143_v20)[_index26] = _118_v0;
              let _153_v27;
              _153_v27 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_148_v25).length),_119_v1);
              let _154_v28;
              let _nw30 = new _module.C0();
              _nw30.__ctor(!(!((_143_v20)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_143_v20).length))])), _153_v27);
              _154_v28 = _nw30;
              _module.__default.m2(((_154_v28.f24) ? ((_119_v1)[_module.__default.safeIndex(_133_v14, new BigNumber((_119_v1).length))]) : (_118_v0)), ((_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))]).isLessThan((_dafny.ZERO).minus((_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))])), _135_globalState);
            } else {
              _133_v14 = (_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))];
              (_135_globalState).f19 = !((_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))]).isEqualTo(_133_v14);
              let _155_v29;
              _155_v29 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_142_v19).length),_dafny.Seq.update(_119_v1, _module.__default.safeIndex((((_122_v4).contains(_118_v0)) ? ((_122_v4).get(_118_v0)) : (_133_v14)), new BigNumber((_119_v1).length)), _118_v0));
              let _156_v30;
              let _nw31 = new _module.C0();
              _nw31.__ctor(_118_v0, _155_v29);
              _156_v30 = _nw31;
              let _157_v31;
              _157_v31 = _dafny.Map.Empty.slice().updateUnsafe(_130_v11,_156_v30);
              let _158_v32;
              _158_v32 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(759),_156_v30);
              let _159_v33;
              let _nw32 = Array((new BigNumber(20)).toNumber());
              _nw32[(_dafny.ZERO).toNumber()] = _156_v30;
              _nw32[(_dafny.ONE).toNumber()] = _156_v30;
              _nw32[(new BigNumber(2)).toNumber()] = _156_v30;
              _nw32[(new BigNumber(3)).toNumber()] = _156_v30;
              _nw32[(new BigNumber(4)).toNumber()] = _156_v30;
              _nw32[(new BigNumber(5)).toNumber()] = _156_v30;
              _nw32[(new BigNumber(6)).toNumber()] = _156_v30;
              _nw32[(new BigNumber(7)).toNumber()] = ((_118_v0) ? ((((_157_v31).contains(_130_v11)) ? ((_157_v31).get(_130_v11)) : (_156_v30))) : (_156_v30));
              _nw32[(new BigNumber(8)).toNumber()] = _156_v30;
              _nw32[(new BigNumber(9)).toNumber()] = _156_v30;
              _nw32[(new BigNumber(10)).toNumber()] = _156_v30;
              _nw32[(new BigNumber(11)).toNumber()] = _156_v30;
              _nw32[(new BigNumber(12)).toNumber()] = _156_v30;
              _nw32[(new BigNumber(13)).toNumber()] = ((!(_118_v0)) ? (_156_v30) : (_156_v30));
              _nw32[(new BigNumber(14)).toNumber()] = _156_v30;
              _nw32[(new BigNumber(15)).toNumber()] = _156_v30;
              _nw32[(new BigNumber(16)).toNumber()] = _156_v30;
              _nw32[(new BigNumber(17)).toNumber()] = (((_158_v32).contains((_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))])) ? ((_158_v32).get((_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))])) : (_156_v30));
              _nw32[(new BigNumber(18)).toNumber()] = _156_v30;
              _nw32[(new BigNumber(19)).toNumber()] = _156_v30;
              _159_v33 = _nw32;
              let _index27 = _module.__default.safeIndex(new BigNumber(48), new BigNumber((_159_v33).length));
              (_159_v33)[_index27] = _156_v30;
              let _index28 = _module.__default.safeIndex(new BigNumber(48), new BigNumber((_159_v33).length));
              let _nw33 = new _module.C0();
              _nw33.__ctor(!(_118_v0), (_155_v29).update(new BigNumber((_119_v1).length), _119_v1));
              (_159_v33)[_index28] = _nw33;
              _module.__default.m2(_118_v0, !(_module.__default.fm11(_125_v7, new BigNumber((_dafny.Seq.of((_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))])).length), _135_globalState)), _135_globalState);
              _133_v14 = (_156_v30).fm8(_135_globalState);
            }
            let _index29 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_143_v20).length));
            (_143_v20)[_index29] = !(_module.__default.fm11(_125_v7, new BigNumber((_dafny.Seq.UnicodeFromString("kvyhg")).length), _135_globalState));
            let _index30 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_143_v20).length));
            (_143_v20)[_index30] = _module.__default.fm11((_dafny.MultiSet.fromElements(_module.__default.fm11(_dafny.MultiSet.fromElements(_118_v0, _118_v0), (_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))], _135_globalState), _118_v0)).Union(_125_v7), _133_v14, _135_globalState);
            if (!(!((_143_v20)[_module.__default.safeIndex(new BigNumber(28), new BigNumber((_143_v20).length))]))) {
              let _160_v34;
              _160_v34 = _module.D0.create_DC2(_118_v0, _133_v14, (_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))], _118_v0);
              let _161_v35;
              _161_v35 = _dafny.Map.Empty.slice().updateUnsafe((_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))],(_143_v20)[_module.__default.safeIndex(new BigNumber(28), new BigNumber((_143_v20).length))]);
              let _162_v36;
              let _nw34 = Array((new BigNumber(23)).toNumber());
              _nw34[(_dafny.ZERO).toNumber()] = (_143_v20)[_module.__default.safeIndex(new BigNumber(28), new BigNumber((_143_v20).length))];
              _nw34[(_dafny.ONE).toNumber()] = (_143_v20)[_module.__default.safeIndex(new BigNumber(28), new BigNumber((_143_v20).length))];
              _nw34[(new BigNumber(2)).toNumber()] = (_143_v20)[_module.__default.safeIndex(new BigNumber(28), new BigNumber((_143_v20).length))];
              _nw34[(new BigNumber(3)).toNumber()] = _118_v0;
              _nw34[(new BigNumber(4)).toNumber()] = false;
              _nw34[(new BigNumber(5)).toNumber()] = true;
              _nw34[(new BigNumber(6)).toNumber()] = _118_v0;
              _nw34[(new BigNumber(7)).toNumber()] = _118_v0;
              _nw34[(new BigNumber(8)).toNumber()] = _118_v0;
              _nw34[(new BigNumber(9)).toNumber()] = _118_v0;
              _nw34[(new BigNumber(10)).toNumber()] = (_160_v34).dtor_cf6;
              _nw34[(new BigNumber(11)).toNumber()] = true;
              _nw34[(new BigNumber(12)).toNumber()] = _118_v0;
              _nw34[(new BigNumber(13)).toNumber()] = true;
              _nw34[(new BigNumber(14)).toNumber()] = (_143_v20)[_module.__default.safeIndex(new BigNumber(28), new BigNumber((_143_v20).length))];
              _nw34[(new BigNumber(15)).toNumber()] = _118_v0;
              _nw34[(new BigNumber(16)).toNumber()] = _118_v0;
              _nw34[(new BigNumber(17)).toNumber()] = true;
              _nw34[(new BigNumber(18)).toNumber()] = !(false);
              _nw34[(new BigNumber(19)).toNumber()] = (_160_v34).dtor_cf3;
              _nw34[(new BigNumber(20)).toNumber()] = !(_118_v0);
              _nw34[(new BigNumber(21)).toNumber()] = _118_v0;
              _nw34[(new BigNumber(22)).toNumber()] = (((_161_v35).contains(new BigNumber(293))) ? ((_161_v35).get(new BigNumber(293))) : (_118_v0));
              _162_v36 = _nw34;
              let _163_v37;
              _163_v37 = _dafny.Map.Empty.slice().updateUnsafe(_162_v36,false);
              _163_v37 = (_163_v37).update(_162_v36, (((_161_v35).contains(new BigNumber(788))) ? ((_161_v35).get(new BigNumber(788))) : (_118_v0)));
              let _164_v38;
              _164_v38 = _module.D3.create_DC8(_133_v14, _133_v14);
              let _pat_let_tv1 = _127_v9;
              let _pat_let_tv2 = _127_v9;
              _133_v14 = (((_119_v1)[_module.__default.safeIndex(_133_v14, new BigNumber((_119_v1).length))]) ? (_module.__default.safeModuloInt((_dafny.ZERO).minus((_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))]), (function (_pat_let2_0) {
                return function (_165_dt__update__tmp_h0) {
                  return function (_pat_let3_0) {
                    return function (_166_dt__update_hcf17_h0) {
                      return _module.D3.create_DC8((_165_dt__update__tmp_h0).dtor_cf16, _166_dt__update_hcf17_h0);
                    }(_pat_let3_0);
                  }((_pat_let_tv2)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_pat_let_tv1).length))]);
                }(_pat_let2_0);
              }(_164_v38)).dtor_cf16)) : (new BigNumber(803)));
              (_135_globalState).f6 = ((_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))]).isLessThan(_133_v14);
              let _167_v39;
              _167_v39 = _dafny.Seq.of(_162_v36);
              _162_v36 = (_167_v39)[_module.__default.safeIndex(_module.__default.safeDivisionInt(_133_v14, (_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))]), new BigNumber((_167_v39).length))];
              let _168_v40;
              let _nw35 = new _module.C1();
              _nw35.__ctor(_module.__default.fm11(_125_v7, new BigNumber((_119_v1).length), _135_globalState), _118_v0);
              _168_v40 = _nw35;
            } else {
              _module.__default.m2(((true) ? ((_143_v20)[_module.__default.safeIndex(new BigNumber(28), new BigNumber((_143_v20).length))]) : (_118_v0)), true, _135_globalState);
              (_135_globalState).f19 = _118_v0;
              let _index31 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_143_v20).length));
              (_143_v20)[_index31] = _module.__default.fm11(_dafny.MultiSet.fromElements((_143_v20)[_module.__default.safeIndex(new BigNumber(28), new BigNumber((_143_v20).length))], _118_v0, _118_v0), _133_v14, _135_globalState);
              _module.__default.m2(_118_v0, _118_v0, _135_globalState);
              let _index32 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_143_v20).length));
              (_143_v20)[_index32] = _118_v0;
            }
            _133_v14 = (_module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_133_v14,(_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))])).length), new BigNumber((_120_v2).length))).minus(new BigNumber(((_124_v6).Union(_124_v6)).length));
          }
        }
      }
      let _169_v41;
      _169_v41 = _dafny.Map.Empty.slice().updateUnsafe((_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))],_dafny.Seq.of(_118_v0));
      let _170_v42;
      let _nw36 = new _module.C0();
      _nw36.__ctor(true, _169_v41);
      _170_v42 = _nw36;
      let _171_v43;
      _171_v43 = _dafny.Seq.of(_170_v42, _170_v42, _170_v42);
      (_135_globalState).f19 = _dafny.Seq.IsProperPrefixOf(_171_v43, _171_v43);
      _module.__default.m2(_118_v0, _118_v0, _135_globalState);
      let _172_v44;
      let _nw37 = new _module.C1();
      _nw37.__ctor(((_118_v0) ? (_118_v0) : (_118_v0)), _118_v0);
      _172_v44 = _nw37;
      _172_v44 = _172_v44;
      let _173_v45;
      let _nw38 = Array((new BigNumber(11)).toNumber());
      _nw38[(_dafny.ZERO).toNumber()] = _127_v9;
      _nw38[(_dafny.ONE).toNumber()] = _127_v9;
      _nw38[(new BigNumber(2)).toNumber()] = _127_v9;
      _nw38[(new BigNumber(3)).toNumber()] = _127_v9;
      _nw38[(new BigNumber(4)).toNumber()] = _127_v9;
      _nw38[(new BigNumber(5)).toNumber()] = _127_v9;
      _nw38[(new BigNumber(6)).toNumber()] = _127_v9;
      _nw38[(new BigNumber(7)).toNumber()] = _127_v9;
      _nw38[(new BigNumber(8)).toNumber()] = _127_v9;
      _nw38[(new BigNumber(9)).toNumber()] = _127_v9;
      _nw38[(new BigNumber(10)).toNumber()] = _127_v9;
      _173_v45 = _nw38;
      let _index33 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_173_v45).length));
      (_173_v45)[_index33] = _127_v9;
      let _index34 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_173_v45).length));
      (_173_v45)[_index34] = _127_v9;
      _133_v14 = (((new BigNumber((_module.__default.fm3((((_122_v4).contains((_172_v44).fm4(_133_v14, (_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))], _172_v44.f23, _135_globalState))) ? ((_122_v4).get((_172_v44).fm4(_133_v14, (_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))], _172_v44.f23, _135_globalState))) : ((_dafny.ZERO).minus(_133_v14))), _135_globalState)).length)).isLessThan(new BigNumber(278))) ? ((_dafny.ZERO).minus(_133_v14)) : (_module.__default.fm0(_172_v44.f23, _118_v0, _135_globalState)));
      let _hi0 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_133_v14, (_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))]));
      for (let _174_i4 = (_133_v14).minus((_dafny.ZERO).minus((_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))])); _174_i4.isLessThan(_hi0); _174_i4 = _174_i4.plus(_dafny.ONE)) {
        _133_v14 = new BigNumber(673);
        let _index35 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length));
        let _index36 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length));
        let _rhs22 = !((_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))]).isEqualTo((_174_i4).minus(_174_i4));
        let _rhs23 = _127_v9;
        let _rhs24 = (_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))];
        let _rhs25 = _174_i4;
        let _lhs17 = _127_v9;
        let _lhs18 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length));
        let _lhs19 = _127_v9;
        let _lhs20 = _module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length));
        _118_v0 = _rhs22;
        _127_v9 = _rhs23;
        _lhs17[_lhs18] = _rhs24;
        _lhs19[_lhs20] = _rhs25;
        (_135_globalState).f18 = new _dafny.CodePoint('e'.codePointAt(0));
        _118_v0 = !(!(((_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))]).isLessThanOrEqualTo((_dafny.ZERO).minus((_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))]))) || (_118_v0));
      }
      let _175_v47;
      let _nw39 = new _module.C0();
      _nw39.__ctor(_172_v44.f23, function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(294)), ((_176_v19) => function (_177_i6) {
          return new BigNumber((_176_v19).length);
        })(_142_v19))).Elements) {
          let _178_v46 = _compr_3;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(294)), ((_179_v19) => function (_177_i6) {
            return new BigNumber((_179_v19).length);
          })(_142_v19)), _178_v46)) {
            _coll3.push([(_178_v46).plus(new BigNumber((_122_v4).length)),_119_v1]);
          }
        }
        return _coll3;
      }());
      _175_v47 = _nw39;
      let _180_v48;
      _180_v48 = _dafny.Map.Empty.slice().updateUnsafe(_175_v47,(_172_v44).f22);
      let _hi1 = new BigNumber((_142_v19).length);
      for (let _181_i5 = new BigNumber((_180_v48).length); _181_i5.isLessThan(_hi1); _181_i5 = _181_i5.plus(_dafny.ONE)) {
        let _182_v49;
        let _183_v50;
        let _out0;
        let _out1;
        let _outcollector0 = (_172_v44).m1(_module.__default.safeDivisionInt(new BigNumber((_120_v2).length), _181_i5), _175_v47.f24, (_172_v44).f22, new BigNumber((_dafny.Seq.Concat(_120_v2, _120_v2)).length), _135_globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _182_v49 = _out0;
        _183_v50 = _out1;
        let _184_v51;
        let _185_v52;
        let _out2;
        let _out3;
        let _outcollector1 = (_172_v44).m0(new BigNumber(313), _135_globalState);
        _out2 = _outcollector1[0];
        _out3 = _outcollector1[1];
        _184_v51 = _out2;
        _185_v52 = _out3;
        let _186_v53;
        _186_v53 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_142_v19).length),_143_v20);
        let _187_v54;
        _187_v54 = _dafny.Seq.of(_143_v20, _143_v20);
        let _188_v55;
        _188_v55 = _dafny.MultiSet.fromElements(_143_v20, _143_v20, (((_186_v53).contains(_181_i5)) ? ((_186_v53).get(_181_i5)) : ((_187_v54)[_module.__default.safeIndex(_133_v14, new BigNumber((_187_v54).length))])), _143_v20, _143_v20);
        let _189_v56;
        _189_v56 = _dafny.MultiSet.fromElements(_143_v20);
        let _190_v57;
        _190_v57 = _dafny.Seq.of((_189_v56));
        _188_v55 = (_190_v57)[_module.__default.safeIndex(_module.__default.safeModuloInt(_181_i5, new BigNumber(508)), new BigNumber((_190_v57).length))];
        _125_v7 = _125_v7;
      }
      if ((_172_v44).f22) {
        let _191_v58;
        _191_v58 = _dafny.Map.Empty.slice().updateUnsafe((_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))],new BigNumber((_148_v25).length));
        let _192_v59;
        _192_v59 = _191_v58;
        _192_v59 = _module.__default.fm12(_135_globalState);
        let _193_v60;
        _193_v60 = (_173_v45)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_173_v45).length))];
        let _194_v61;
        _194_v61 = _dafny.Set.fromElements(_193_v60, _193_v60);
        let _195_v62;
        _195_v62 = _dafny.Map.Empty.slice().updateUnsafe(_175_v47,_194_v61);
        (_172_v44).f23 = ((_194_v61).Difference(_194_v61)).IsDisjointFrom((((_195_v62).contains(_175_v47)) ? ((_195_v62).get(_175_v47)) : (_194_v61)));
        _133_v14 = ((_133_v14).multipliedBy(new BigNumber(52))).plus(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_175_v47,_172_v44.f23)).Merge(_180_v48)).length));
        _118_v0 = _module.__default.fm11(_125_v7, _133_v14, _135_globalState);
        (_135_globalState).f19 = (_172_v44).f22;
      } else {
        let _196_v63;
        _196_v63 = _dafny.MultiSet.fromElements(_143_v20, _143_v20);
        let _197_v64;
        _197_v64 = (_196_v63).Union(_196_v63);
        let _source1 = _197_v64;
        let _198___mcc_h0 = _source1;
        let _199_cf18 = _198___mcc_h0;
        let _200_v65;
        _200_v65 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),new BigNumber(938));
        let _201_v66;
        _201_v66 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_200_v65).length),(_173_v45)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_173_v45).length))]);
        _201_v66 = (_201_v66).update(new BigNumber(751), (_173_v45)[_module.__default.safeIndex(new BigNumber(408), new BigNumber((_173_v45).length))]);
        let _nw40 = Array((new BigNumber(14)).toNumber());
        _nw40[(_dafny.ZERO).toNumber()] = (_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))];
        _nw40[(_dafny.ONE).toNumber()] = _module.__default.safeModuloInt((_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))], _133_v14);
        _nw40[(new BigNumber(2)).toNumber()] = (_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))];
        _nw40[(new BigNumber(3)).toNumber()] = _module.__default.fm0(!(true), _172_v44.f23, _135_globalState);
        _nw40[(new BigNumber(4)).toNumber()] = (new BigNumber(338)).multipliedBy(_133_v14);
        _nw40[(new BigNumber(5)).toNumber()] = (((_125_v7).contains(_172_v44.f23)) ? ((_125_v7).get(_172_v44.f23)) : (_133_v14));
        _nw40[(new BigNumber(6)).toNumber()] = ((_dafny.ZERO).minus(_133_v14)).plus(new BigNumber((_120_v2).length));
        _nw40[(new BigNumber(7)).toNumber()] = _133_v14;
        _nw40[(new BigNumber(8)).toNumber()] = (_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))];
        _nw40[(new BigNumber(9)).toNumber()] = (_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))];
        _nw40[(new BigNumber(10)).toNumber()] = _module.__default.safeModuloInt((_142_v19)[_module.__default.safeIndex((_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))], new BigNumber((_142_v19).length))], _133_v14);
        _nw40[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(_133_v14);
        _nw40[(new BigNumber(12)).toNumber()] = (_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))];
        _nw40[(new BigNumber(13)).toNumber()] = (_127_v9)[_module.__default.safeIndex(new BigNumber(562), new BigNumber((_127_v9).length))];
        _127_v9 = _nw40;
        let _index37 = _module.__default.safeIndex(new BigNumber(408), new BigNumber((_173_v45).length));
        let _init2 = ((_202_v14) => function (_203_i7) {
          return (_203_i7).plus(_202_v14);
        })(_133_v14);
        let _nw41 = Array((new BigNumber(6)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw41.length); _i0_2++) {
          _nw41[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        (_173_v45)[_index37] = _nw41;
        let _204_v67;
        _204_v67 = _dafny.Set.fromElements(_172_v44);
        _204_v67 = _204_v67;
        let _205_v68;
        let _nw42 = new _module.C0();
        _nw42.__ctor(_175_v47.f24, (_175_v47).f25);
        _205_v68 = _nw42;
        (_175_v47).f24 = (_module.__default.safeDivisionInt((_170_v42).fm8(_135_globalState), new BigNumber(-823))).isEqualTo(_133_v14);
        (_175_v47).f24 = _172_v44.f23;
        let _206_v69;
        let _207_v70;
        let _out4;
        let _out5;
        let _outcollector2 = (_172_v44).m0(new BigNumber((_142_v19).length), _135_globalState);
        _out4 = _outcollector2[0];
        _out5 = _outcollector2[1];
        _206_v69 = _out4;
        _207_v70 = _out5;
      }
      process.stdout.write(_dafny.toString(_118_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_119_v1, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_120_v2).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_121_v3).dtor_cf1, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_121_v3).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_122_v4).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(9)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_123_v5).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(_dafny.Seq.of(true), new BigNumber(9)),_dafny.MultiSet.fromElements(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v6).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_125_v7).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_126_v8, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("pvnnsysbg"), _dafny.Seq.UnicodeFromString("pvnnsysbg"), _dafny.Seq.UnicodeFromString("pvnnsysbg")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v9)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v9)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v9)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v9)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v9)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v9)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v9)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v9)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v9)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v9)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v9)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v9)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v9)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v9)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v9)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v9)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v9)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v9)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v9)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_130_v11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_v12).equals(_dafny.Set.fromElements(new _dafny.CodePoint('j'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_132_v13).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(true),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_133_v14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_134_v15).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(588))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_globalState.f0).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(36)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_globalState).f3).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_135_globalState.f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_135_globalState).f10, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("pvnnsysbg"), _dafny.Seq.UnicodeFromString("pvnnsysbg"), _dafny.Seq.UnicodeFromString("pvnnsysbg"), _dafny.Seq.UnicodeFromString("pvnnsysbg"), _dafny.Seq.UnicodeFromString("pvnnsysbg"), _dafny.Seq.UnicodeFromString("pvnnsysbg")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_135_globalState).f11).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_globalState.f13)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_globalState.f13)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_globalState.f13)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_globalState.f13)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_globalState.f13)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_globalState.f13)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_globalState.f13)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_globalState.f13)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_globalState.f13)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_globalState.f13)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_globalState.f13)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_globalState.f13)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_globalState.f13)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_globalState.f13)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_globalState.f13)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_globalState.f13)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_globalState.f13)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_globalState.f13)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_globalState.f13)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_globalState).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_globalState).f15).equals(_dafny.Map.Empty.slice())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_globalState).f16).equals(_dafny.Set.fromElements(new _dafny.CodePoint('j'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_globalState).f17).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(true),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_135_globalState.f18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_135_globalState.f19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_globalState).f20).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(588))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_globalState).f21));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_137_v16).equals(_dafny.Set.fromElements(_dafny.Seq.of(true), _dafny.Seq.of(true, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v17).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(new BigNumber(588), _dafny.ONE)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_141_v18, _dafny.Seq.of(_dafny.MultiSet.fromElements(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_142_v19, _dafny.Seq.of(new BigNumber(588), _dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_144_v21).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_144_v21).dtor_cf4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_144_v21).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_144_v21).dtor_cf6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_v22).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(944),_module.D0.create_DC3(_module.D0.create_DC2(false, new BigNumber(754), new BigNumber(561), false))).updateUnsafe(_dafny.ONE,_module.D0.create_DC3(_module.D0.create_DC3(_module.D0.create_DC2(true, new BigNumber(111), new BigNumber(-539), false)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_146_v23)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new _dafny.CodePoint('j'.codePointAt(0))),new BigNumber(1176)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_v24).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new _dafny.CodePoint('j'.codePointAt(0))),new BigNumber(991)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_148_v25).equals(_dafny.Set.fromElements(new BigNumber(9)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_149_v26).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('j'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_150_i2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_169_v41).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(588),_dafny.Seq.of(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_171_v43).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_172_v44).f22));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_172_v44.f23));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ZERO])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ZERO])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ZERO])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ZERO])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ZERO])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ZERO])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ZERO])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ZERO])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ZERO])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ZERO])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ZERO])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ZERO])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ZERO])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ZERO])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ZERO])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ZERO])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ZERO])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ZERO])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ZERO])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ONE])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ONE])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ONE])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ONE])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ONE])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ONE])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ONE])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ONE])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ONE])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ONE])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ONE])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ONE])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ONE])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ONE])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ONE])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ONE])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ONE])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ONE])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[_dafny.ONE])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(2)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(2)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(2)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(2)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(2)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(2)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(2)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(2)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(2)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(2)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(2)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(2)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(2)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(2)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(2)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(2)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(2)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(2)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(2)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(3)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(3)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(3)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(3)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(3)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(3)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(3)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(3)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(3)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(3)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(3)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(3)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(3)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(3)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(3)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(3)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(3)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(3)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(3)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(4)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(4)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(4)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(4)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(4)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(4)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(4)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(4)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(4)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(4)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(4)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(4)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(4)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(4)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(4)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(4)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(4)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(4)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(4)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(5)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(5)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(5)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(5)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(5)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(5)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(5)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(5)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(5)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(5)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(5)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(5)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(5)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(5)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(5)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(5)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(5)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(5)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(5)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(6)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(6)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(6)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(6)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(6)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(6)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(6)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(6)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(6)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(6)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(6)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(6)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(6)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(6)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(6)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(6)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(6)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(6)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(6)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(7)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(7)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(7)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(7)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(7)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(7)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(7)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(7)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(7)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(7)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(7)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(7)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(7)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(7)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(7)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(7)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(7)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(7)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(7)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(8)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(8)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(8)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(8)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(8)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(8)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(8)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(8)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(8)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(8)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(8)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(8)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(8)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(8)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(8)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(8)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(8)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(8)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(8)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(9)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(9)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(9)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(9)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(9)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(9)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(9)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(9)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(9)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(9)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(9)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(9)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(9)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(9)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(9)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(9)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(9)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(9)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(9)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(10)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(10)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(10)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(10)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(10)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(10)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(10)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(10)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(10)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(10)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(10)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(10)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(10)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(10)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(10)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(10)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(10)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(10)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_173_v45)[new BigNumber(10)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_175_v47.f24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_175_v47).f25).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(3),_dafny.Seq.of(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_180_v48).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1, cf2) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      return $dt;
    }
    static create_DC2(cf3, cf4, cf5, cf6) {
      let $dt = new D0(1);
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(2);
      $dt.cf0 = cf0;
      return $dt;
    }
    static create_DC3(cf7) {
      let $dt = new D0(3);
      $dt.cf7 = cf7;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC0() { return this.$tag === 2; }
    get is_DC3() { return this.$tag === 3; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf7() { return this.cf7; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else if (this.$tag === 3) {
        return "D0.DC3" + "(" + _dafny.toString(this.cf7) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1) && _dafny.areEqual(this.cf2, other.cf2);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf3 === other.cf3 && _dafny.areEqual(this.cf4, other.cf4) && _dafny.areEqual(this.cf5, other.cf5) && this.cf6 === other.cf6;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf0, other.cf0);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf7, other.cf7);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(_dafny.Seq.of(), _dafny.ZERO);
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
    static create_DC4(cf8) {
      let $dt = new D1(0);
      $dt.cf8 = cf8;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get dtor_cf8() { return this.cf8; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf8) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf8, other.cf8);
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
          return D1.Default();
        }
      };
    }
  }

  $module.D2 = class D2 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC5(cf9) {
      let $dt = new D2(0);
      $dt.cf9 = cf9;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get dtor_cf9() { return this.cf9; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf9) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf9 === other.cf9;
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
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC7(cf11, cf12, cf13, cf14, cf15) {
      let $dt = new D3(0);
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC8(cf16, cf17) {
      let $dt = new D3(1);
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC6(cf10) {
      let $dt = new D3(2);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC8() { return this.$tag === 1; }
    get is_DC6() { return this.$tag === 2; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC7" + "(" + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC6" + "(" + this.cf10.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf11, other.cf11) && this.cf12 === other.cf12 && this.cf13 === other.cf13 && this.cf14 === other.cf14 && this.cf15 === other.cf15;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf16, other.cf16) && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC7(_dafny.Map.Empty, false, false, false, []);
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
    static create_DC9(cf18) {
      let $dt = new D4(0);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC9" + "(" + _dafny.toString(this.cf18) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf18, other.cf18);
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
          return D4.Default();
        }
      };
    }
  }

  $module.D5 = class D5 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC11(cf20, cf21, cf22, cf23) {
      let $dt = new D5(0);
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC10(cf19) {
      let $dt = new D5(1);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC11" + "(" + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC10" + "(" + _dafny.toString(this.cf19) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf20 === other.cf20 && this.cf21 === other.cf21 && this.cf22 === other.cf22 && this.cf23 === other.cf23;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf19, other.cf19);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC11(false, false, [], null);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D5.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f0 = _dafny.Map.Empty;
      this.f6 = false;
      this.f13 = [];
      this.f18 = new _dafny.CodePoint('D'.codePointAt(0));
      this.f19 = false;
      this._f1 = false;
      this._f2 = false;
      this._f3 = _dafny.MultiSet.Empty;
      this._f4 = false;
      this._f5 = _dafny.ZERO;
      this._f7 = _dafny.ZERO;
      this._f8 = false;
      this._f9 = _dafny.ZERO;
      this._f10 = _dafny.Seq.of();
      this._f11 = _dafny.Seq.UnicodeFromString("");
      this._f12 = _dafny.ZERO;
      this._f14 = _dafny.ZERO;
      this._f15 = _dafny.Map.Empty;
      this._f16 = _dafny.Set.Empty;
      this._f17 = _dafny.Map.Empty;
      this._f20 = _dafny.Map.Empty;
      this._f21 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21) {
      let _this = this;
      (_this).f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this).f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this).f13 = f13;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this).f18 = f18;
      (_this).f19 = f19;
      (_this)._f20 = f20;
      (_this)._f21 = f21;
      return;
    }
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f2() {
      let _this = this;
      return _this._f2;
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
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f10() {
      let _this = this;
      return _this._f10;
    };
    get f11() {
      let _this = this;
      return _this._f11;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
    get f15() {
      let _this = this;
      return _this._f15;
    };
    get f16() {
      let _this = this;
      return _this._f16;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
    get f21() {
      let _this = this;
      return _this._f21;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this.f24 = false;
      this._f25 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f24, f25) {
      let _this = this;
      (_this).f24 = f24;
      (_this)._f25 = f25;
      return;
    }
    fm8(globalState) {
      let _this = this;
      return new BigNumber(694);
    };
    fm9(p0, p1, p2, globalState) {
      let _this = this;
      return new _dafny.CodePoint('h'.codePointAt(0));
    };
    get f25() {
      let _this = this;
      return _this._f25;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this.f23 = false;
      this._f22 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f22, f23) {
      let _this = this;
      (_this)._f22 = f22;
      (_this).f23 = f23;
      return;
    }
    fm4(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f22;
    };
    m0(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Map.Empty;
      let _208_v0;
      _208_v0 = new _dafny.CodePoint('a'.codePointAt(0));
      let _209_v1;
      _209_v1 = _dafny.Map.Empty.slice().updateUnsafe(_208_v0,false);
      let _210_v2;
      _210_v2 = _dafny.Seq.of(_209_v1);
      _210_v2 = _210_v2;
      let _211_v3;
      _211_v3 = _dafny.Seq.of((_this).f22);
      let _212_v4;
      _212_v4 = _dafny.Seq.of(_211_v3, _211_v3);
      let _213_v5;
      _213_v5 = _module.D0.create_DC1((_212_v4)[_module.__default.safeIndex(p0, new BigNumber((_212_v4).length))], p0);
      let _pat_let_tv3 = p0;
      let _source2 = function (_pat_let4_0) {
        return function (_214_dt__update__tmp_h0) {
          return function (_pat_let5_0) {
            return function (_215_dt__update_hcf2_h0) {
              return _module.D0.create_DC1((_214_dt__update__tmp_h0).dtor_cf1, _215_dt__update_hcf2_h0);
            }(_pat_let5_0);
          }(_pat_let_tv3);
        }(_pat_let4_0);
      }(_213_v5);
      if (_source2.is_DC1) {
        let _216___mcc_h0 = (_source2).cf1;
        let _217___mcc_h1 = (_source2).cf2;
        let _218_cf2 = _217___mcc_h1;
        let _219_cf1 = _216___mcc_h0;
        let _220_v6;
        _220_v6 = _dafny.Seq.UnicodeFromString("kgrcsul");
        let _221_v7;
        _221_v7 = _dafny.Map.Empty.slice().updateUnsafe(_220_v6,(_dafny.ZERO).minus(_module.__default.safeDivisionInt(_218_cf2, new BigNumber((_219_cf1).length))));
        _221_v7 = (_221_v7).update(_dafny.Seq.Concat(_220_v6, _220_v6), (_dafny.ZERO).minus(_218_cf2));
        let _222_v8;
        let _init3 = ((_223_cf2) => function (_224_i0) {
          return _module.__default.safeModuloInt(_224_i0, _223_cf2);
        })(_218_cf2);
        let _nw43 = Array((new BigNumber(8)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw43.length); _i0_3++) {
          _nw43[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _222_v8 = _nw43;
        let _225_v9;
        _225_v9 = _dafny.MultiSet.fromElements(_222_v8, _222_v8);
        if (((_225_v9).IsProperSubsetOf(_225_v9)) && (_dafny.Seq.IsPrefixOf(_211_v3, _dafny.Seq.of((_this).fm4(new BigNumber(-705), _218_cf2, _this.f23, globalState))))) {
          (globalState).f19 = !(!((_this).f22) || ((_this).f22));
          let _rhs26 = _222_v8;
          let _rhs27 = (((_221_v7).contains(_220_v6)) ? ((_221_v7).get(_220_v6)) : (_218_cf2));
          let _lhs21 = globalState;
          _lhs21.f13 = _rhs26;
          _218_cf2 = _rhs27;
          _218_cf2 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_module.__default.safeModuloInt(_218_cf2, _218_cf2), p0));
          (globalState).f18 = _208_v0;
          _220_v6 = _dafny.Seq.UnicodeFromString("upfp");
        } else {
          (_this).f23 = (_this).f22;
          (globalState).f6 = !(false);
          let _226_v10;
          _226_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,_219_cf1);
          let _227_v11;
          _227_v11 = _dafny.Set.fromElements(_208_v0, _208_v0, _module.__default.fm5(globalState), _208_v0);
          let _228_v12;
          _228_v12 = _dafny.Seq.of(_227_v11, _227_v11);
          let _rhs28 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Concat(_211_v3, (((_226_v10).contains(_this.f23)) ? ((_226_v10).get(_this.f23)) : (_219_cf1))), _module.__default.safeIndex(new BigNumber((_228_v12).length), new BigNumber((_dafny.Seq.Concat(_211_v3, (((_226_v10).contains(_this.f23)) ? ((_226_v10).get(_this.f23)) : (_219_cf1)))).length)), _this.f23), _dafny.Seq.of(_this.f23, _this.f23))).length));
          let _rhs29 = _this.f23;
          let _lhs22 = globalState;
          _218_cf2 = _rhs28;
          _lhs22.f19 = _rhs29;
          let _229_v13;
          _229_v13 = _dafny.MultiSet.fromElements(false);
          let _230_v14;
          _230_v14 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("vtobwgfoy"));
          let _231_v15;
          _231_v15 = _dafny.Seq.of(_229_v13, ((_229_v13).update((_this).f22, _module.__default.abs(new BigNumber((_230_v14).length)))).Union(_229_v13), _dafny.MultiSet.fromElements(true, (_this).f22));
          let _232_v16;
          _232_v16 = _dafny.Set.fromElements(_218_cf2);
          let _rhs30 = (((_232_v16).contains(_218_cf2)) ? ((_225_v9).IsSubsetOf(_225_v9)) : (((_this.f23) ? (!((_this).f22)) : (_this.f23))));
          let _rhs31 = _dafny.Seq.Concat(_231_v15, _dafny.Seq.Create(_module.__default.abs(new BigNumber(632)), ((_233_v13) => function (_234_i1) {
            return _233_v13;
          })(_229_v13)));
          let _lhs23 = globalState;
          _lhs23.f6 = _rhs30;
          _231_v15 = _rhs31;
          let _235_v17;
          let _init4 = function (_236_i2) {
            return (((_this).f22) ? ((_this).f22) : ((_this).f22));
          };
          let _nw44 = Array((new BigNumber(12)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw44.length); _i0_4++) {
            _nw44[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _235_v17 = _nw44;
          _235_v17 = _235_v17;
        }
        let _237_v18;
        let _nw45 = Array((new BigNumber(24)).toNumber()).fill(false);
        _237_v18 = _nw45;
        let _index38 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_237_v18).length));
        (_237_v18)[_index38] = _this.f23;
        let _238_v19;
        _238_v19 = _dafny.Set.fromElements(_218_cf2, p0);
        let _index39 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_237_v18).length));
        let _rhs32 = new BigNumber((_220_v6).length);
        let _rhs33 = (_211_v3)[_module.__default.safeIndex(new BigNumber((_238_v19).length), new BigNumber((_211_v3).length))];
        let _rhs34 = _this.f23;
        let _rhs35 = _211_v3;
        let _lhs24 = globalState;
        let _lhs25 = _237_v18;
        let _lhs26 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_237_v18).length));
        _218_cf2 = _rhs32;
        _lhs24.f19 = _rhs33;
        _lhs25[_lhs26] = _rhs34;
        _211_v3 = _rhs35;
        let _239_v20;
        _239_v20 = _dafny.Seq.of(p0);
        let _index40 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((_222_v8).length));
        (_222_v8)[_index40] = new BigNumber((_239_v20).length);
        let _index41 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_237_v18).length));
        let _index42 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((_222_v8).length));
        let _rhs36 = (_237_v18)[_module.__default.safeIndex(new BigNumber(192), new BigNumber((_237_v18).length))];
        let _rhs37 = (_dafny.ZERO).minus(_218_cf2);
        let _rhs38 = _dafny.Seq.contains(_dafny.Seq.update(_dafny.Seq.Concat(_219_cf1, _219_cf1), _module.__default.safeIndex(new BigNumber((_220_v6).length), new BigNumber((_dafny.Seq.Concat(_219_cf1, _219_cf1)).length)), _this.f23), (_this).f22);
        let _lhs27 = _237_v18;
        let _lhs28 = _module.__default.safeIndex(new BigNumber(192), new BigNumber((_237_v18).length));
        let _lhs29 = _222_v8;
        let _lhs30 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((_222_v8).length));
        let _lhs31 = globalState;
        _lhs27[_lhs28] = _rhs36;
        _lhs29[_lhs30] = _rhs37;
        _lhs31.f19 = _rhs38;
      } else if (_source2.is_DC2) {
        let _240___mcc_h2 = (_source2).cf3;
        let _241___mcc_h3 = (_source2).cf4;
        let _242___mcc_h4 = (_source2).cf5;
        let _243___mcc_h5 = (_source2).cf6;
        let _244_cf6 = _243___mcc_h5;
        let _245_cf5 = _242___mcc_h4;
        let _246_cf4 = _241___mcc_h3;
        let _247_cf3 = _240___mcc_h2;
        if (_this.f23) {
          let _248_v21;
          let _init5 = ((_249_cf4) => function (_250_i3) {
            return (_250_i3).plus(_249_cf4);
          })(_246_cf4);
          let _nw46 = Array((new BigNumber(9)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw46.length); _i0_5++) {
            _nw46[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _248_v21 = _nw46;
          let _251_v22;
          _251_v22 = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,_248_v21);
          _251_v22 = (_251_v22).update(_244_cf6, _248_v21);
          let _252_v23;
          let _nw47 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Map.Empty);
          _252_v23 = _nw47;
          let _253_v24;
          _253_v24 = _dafny.Map.Empty.slice().updateUnsafe(_244_cf6,p0);
          let _index43 = _module.__default.safeIndex(new BigNumber(962), new BigNumber((_252_v23).length));
          (_252_v23)[_index43] = _253_v24;
          let _254_v25;
          _254_v25 = _dafny.Seq.of(_253_v24);
          let _255_v26;
          _255_v26 = _dafny.Seq.of(new BigNumber((_253_v24).length));
          let _256_v27;
          _256_v27 = _dafny.Seq.of((_255_v26)[_module.__default.safeIndex(_246_cf4, new BigNumber((_255_v26).length))]);
          let _index44 = _module.__default.safeIndex(new BigNumber(962), new BigNumber((_252_v23).length));
          (_252_v23)[_index44] = ((_253_v24).Merge(_253_v24)).Merge((_254_v25)[_module.__default.safeIndex(new BigNumber((_256_v27).length), new BigNumber((_254_v25).length))]);
          let _257_v28;
          let _258_v29;
          let _out6;
          let _out7;
          let _outcollector3 = (_this).m1(new BigNumber(((_252_v23)[_module.__default.safeIndex(new BigNumber(962), new BigNumber((_252_v23).length))]).length), _this.f23, (_this).f22, _246_cf4, globalState);
          _out6 = _outcollector3[0];
          _out7 = _outcollector3[1];
          _257_v28 = _out6;
          _258_v29 = _out7;
          _246_cf4 = _246_cf4;
          _246_cf4 = new BigNumber(669);
        } else {
          _245_cf5 = p0;
          let _259_v30;
          let _nw48 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
          _259_v30 = _nw48;
          let _index45 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_259_v30).length));
          (_259_v30)[_index45] = (_dafny.ZERO).minus((_245_cf5).plus(_245_cf5));
          let _260_v31;
          _260_v31 = _dafny.Map.Empty.slice().updateUnsafe(_244_cf6,new BigNumber(450));
          let _261_v32;
          _261_v32 = _dafny.Seq.of((_dafny.ZERO).minus(_245_cf5), p0);
          let _262_v33;
          _262_v33 = _dafny.Map.Empty.slice().updateUnsafe(_260_v31,(_this).fm4(new BigNumber(-224), (_261_v32)[_module.__default.safeIndex(_246_cf4, new BigNumber((_261_v32).length))], _244_cf6, globalState));
          let _263_v34;
          _263_v34 = _dafny.Map.Empty.slice().updateUnsafe(_246_cf4,_247_cf3);
          let _index46 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_259_v30).length));
          let _rhs39 = _246_cf4;
          let _rhs40 = new BigNumber((_262_v33).length);
          let _rhs41 = (new BigNumber((_263_v34).length)).plus((_dafny.ZERO).minus(_245_cf5));
          let _rhs42 = (new BigNumber(274)).multipliedBy(new BigNumber((_261_v32).length));
          let _rhs43 = !(_244_cf6);
          let _lhs32 = _259_v30;
          let _lhs33 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_259_v30).length));
          let _lhs34 = globalState;
          _245_cf5 = _rhs39;
          _245_cf5 = _rhs40;
          _lhs32[_lhs33] = _rhs41;
          _245_cf5 = _rhs42;
          _lhs34.f19 = _rhs43;
          let _264_v35;
          _264_v35 = _dafny.MultiSet.fromElements(_244_cf6);
          let _265_v36;
          _265_v36 = _module.D0.create_DC2((_this).f22, p0, new BigNumber((_264_v35).cardinality()), _244_cf6);
          let _266_v37;
          _266_v37 = _dafny.Map.Empty.slice().updateUnsafe(_265_v36,(_this).fm4(_245_cf5, (_259_v30)[_module.__default.safeIndex(new BigNumber(828), new BigNumber((_259_v30).length))], _244_cf6, globalState));
          _266_v37 = (_266_v37).update(_265_v36, (_this).f22);
          let _267_v38;
          _267_v38 = _dafny.Seq.UnicodeFromString("nvttpt");
          let _268_v39;
          let _nw49 = Array((new BigNumber(13)).toNumber());
          _nw49[(_dafny.ZERO).toNumber()] = (p0).isEqualTo((_261_v32)[_module.__default.safeIndex((_259_v30)[_module.__default.safeIndex(new BigNumber(828), new BigNumber((_259_v30).length))], new BigNumber((_261_v32).length))]);
          _nw49[(_dafny.ONE).toNumber()] = _244_cf6;
          _nw49[(new BigNumber(2)).toNumber()] = (_247_cf3) === (_244_cf6);
          _nw49[(new BigNumber(3)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_267_v38, _267_v38);
          _nw49[(new BigNumber(4)).toNumber()] = !(_244_cf6);
          _nw49[(new BigNumber(5)).toNumber()] = (_this).f22;
          _nw49[(new BigNumber(6)).toNumber()] = (_this).fm4(_245_cf5, _module.__default.fm0(!(!(_244_cf6)), _this.f23, globalState), false, globalState);
          _nw49[(new BigNumber(7)).toNumber()] = _dafny.areEqual(_211_v3, _211_v3);
          _nw49[(new BigNumber(8)).toNumber()] = _this.f23;
          _nw49[(new BigNumber(9)).toNumber()] = (_this).f22;
          _nw49[(new BigNumber(10)).toNumber()] = _247_cf3;
          _nw49[(new BigNumber(11)).toNumber()] = _this.f23;
          _nw49[(new BigNumber(12)).toNumber()] = _this.f23;
          _268_v39 = _nw49;
          let _269_v40;
          _269_v40 = _module.D0.create_DC0(new BigNumber(554));
          let _270_v41;
          _270_v41 = _dafny.Seq.of(_268_v39, _268_v39);
          let _rhs44 = (_dafny.ZERO).minus((_269_v40).dtor_cf0);
          let _rhs45 = true;
          let _rhs46 = !((((_this).f22) ? ((_this).fm4(_245_cf5, new BigNumber((_dafny.MultiSet.FromArray(_261_v32)).cardinality()), _244_cf6, globalState)) : (_244_cf6)));
          let _rhs47 = (_270_v41)[_module.__default.safeIndex(p0, new BigNumber((_270_v41).length))];
          let _lhs35 = _this;
          let _lhs36 = _this;
          _245_cf5 = _rhs44;
          _lhs35.f23 = _rhs45;
          _lhs36.f23 = _rhs46;
          _268_v39 = _rhs47;
          let _271_v42;
          _271_v42 = _dafny.MultiSet.fromElements(new BigNumber((_211_v3).length), new BigNumber(((_module.__default.fm6(_246_cf4, _247_cf3, globalState)).Difference((_264_v35).update(false, _module.__default.abs(p0)))).cardinality()), _245_cf5);
          _271_v42 = _271_v42;
        }
        let _272_v43;
        _272_v43 = _dafny.Seq.UnicodeFromString("tbjypvnv");
        let _273_v44;
        _273_v44 = _dafny.Seq.of((new BigNumber((_dafny.Seq.update(_272_v43, _module.__default.safeIndex(p0, new BigNumber((_272_v43).length)), _208_v0)).length)).multipliedBy(_246_cf4), p0, _245_cf5, _245_cf5, p0);
        _273_v44 = _dafny.Seq.Concat(_273_v44, _273_v44);
        if (!_dafny.areEqual(_dafny.Seq.Concat(_272_v43, _272_v43), _272_v43)) {
          let _274_v45;
          _274_v45 = _dafny.MultiSet.fromElements(_247_cf3, !(!(_this.f23)));
          let _275_v46;
          _275_v46 = _dafny.Map.Empty.slice().updateUnsafe(_274_v45,_208_v0);
          _275_v46 = (_275_v46).update((_274_v45).Union(_274_v45), _208_v0);
          let _276_v47;
          let _nw50 = Array((new BigNumber(21)).toNumber()).fill(false);
          _276_v47 = _nw50;
          let _index47 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_276_v47).length));
          (_276_v47)[_index47] = (_245_cf5).isEqualTo(p0);
          let _277_v48;
          let _nw51 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Map.Empty);
          _277_v48 = _nw51;
          let _278_v49;
          _278_v49 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_273_v44)).length)),p0);
          let _279_v50;
          _279_v50 = _278_v49;
          let _index48 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_277_v48).length));
          (_277_v48)[_index48] = (_278_v49).Merge((_279_v50));
          let _index49 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_276_v47).length));
          let _index50 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_277_v48).length));
          let _rhs48 = !((_this).f22);
          let _rhs49 = _278_v49;
          let _lhs37 = _276_v47;
          let _lhs38 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_276_v47).length));
          let _lhs39 = _277_v48;
          let _lhs40 = _module.__default.safeIndex(new BigNumber(466), new BigNumber((_277_v48).length));
          _lhs37[_lhs38] = _rhs48;
          _lhs39[_lhs40] = _rhs49;
          r0 = (_this).fm4(new BigNumber(485), (_dafny.ZERO).minus((_245_cf5).multipliedBy(_245_cf5)), _this.f23, globalState);
          let _280_v51;
          _280_v51 = _dafny.Map.Empty.slice().updateUnsafe(_247_cf3,_244_cf6);
          let _281_v52;
          _281_v52 = _dafny.Map.Empty.slice().updateUnsafe(_213_v5,((((_280_v51).contains(_244_cf6)) ? ((_280_v51).get(_244_cf6)) : (!(_244_cf6)))) === (_this.f23));
          r0 = (((_281_v52).contains(_213_v5)) ? ((_281_v52).get(_213_v5)) : (_this.f23));
          let _index51 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_276_v47).length));
          (_276_v47)[_index51] = _244_cf6;
        } else {
          let _282_v53;
          _282_v53 = _dafny.Set.fromElements(new BigNumber(738), _245_cf5, _246_cf4, new BigNumber((_272_v43).length));
          let _283_v54;
          _283_v54 = _module.D0.create_DC0(new BigNumber(818));
          _282_v53 = _dafny.Set.fromElements((_283_v54).dtor_cf0, _246_cf4, p0);
          _245_cf5 = _245_cf5;
          (globalState).f19 = _this.f23;
          _244_cf6 = (_this).fm4(_module.__default.fm0(false, (_this).f22, globalState), _246_cf4, _dafny.Seq.IsProperPrefixOf(_273_v44, _dafny.Seq.Create(_module.__default.abs(new BigNumber(157)), function (_284_i4) {
            return new BigNumber((_dafny.Seq.UnicodeFromString("m")).length);
          })), globalState);
          let _285_v55;
          _285_v55 = _dafny.Set.fromElements((_this).f22);
          let _286_v56;
          _286_v56 = _dafny.MultiSet.fromElements(_247_cf3, _244_cf6);
          let _287_v57;
          _287_v57 = _module.D0.create_DC2(!((_this).f22), p0, new BigNumber(301), (_this).f22);
          let _288_v58;
          let _nw52 = Array((new BigNumber(21)).toNumber());
          _nw52[(_dafny.ZERO).toNumber()] = _245_cf5;
          _nw52[(_dafny.ONE).toNumber()] = _245_cf5;
          _nw52[(new BigNumber(2)).toNumber()] = new BigNumber((_module.__default.fm7(globalState)).length);
          _nw52[(new BigNumber(3)).toNumber()] = _245_cf5;
          _nw52[(new BigNumber(4)).toNumber()] = new BigNumber((_285_v55).length);
          _nw52[(new BigNumber(5)).toNumber()] = p0;
          _nw52[(new BigNumber(6)).toNumber()] = _245_cf5;
          _nw52[(new BigNumber(7)).toNumber()] = _246_cf4;
          _nw52[(new BigNumber(8)).toNumber()] = new BigNumber((_286_v56).cardinality());
          _nw52[(new BigNumber(9)).toNumber()] = p0;
          _nw52[(new BigNumber(10)).toNumber()] = _245_cf5;
          _nw52[(new BigNumber(11)).toNumber()] = (_273_v44)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_273_v44).length))];
          _nw52[(new BigNumber(12)).toNumber()] = _246_cf4;
          _nw52[(new BigNumber(13)).toNumber()] = _245_cf5;
          _nw52[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.Seq.of((_this).fm4(_246_cf4, new BigNumber(923), (_this).f22, globalState))).length);
          _nw52[(new BigNumber(15)).toNumber()] = _245_cf5;
          _nw52[(new BigNumber(16)).toNumber()] = _246_cf4;
          _nw52[(new BigNumber(17)).toNumber()] = _245_cf5;
          _nw52[(new BigNumber(18)).toNumber()] = _module.__default.fm0(!(_this.f23), (_287_v57).dtor_cf6, globalState);
          _nw52[(new BigNumber(19)).toNumber()] = _245_cf5;
          _nw52[(new BigNumber(20)).toNumber()] = _246_cf4;
          _288_v58 = _nw52;
          let _289_v59;
          _289_v59 = _dafny.Map.Empty.slice().updateUnsafe(_288_v58,_286_v56);
          _289_v59 = (_289_v59).update(_288_v58, _286_v56);
        }
        let _290_v60;
        let _nw53 = Array((new BigNumber(4)).toNumber());
        _nw53[(_dafny.ZERO).toNumber()] = _244_cf6;
        _nw53[(_dafny.ONE).toNumber()] = _244_cf6;
        _nw53[(new BigNumber(2)).toNumber()] = (_this).f22;
        _nw53[(new BigNumber(3)).toNumber()] = !(false);
        _290_v60 = _nw53;
        let _index52 = _module.__default.safeIndex(new BigNumber(134), new BigNumber((_290_v60).length));
        (_290_v60)[_index52] = _244_cf6;
        let _index53 = _module.__default.safeIndex(new BigNumber(134), new BigNumber((_290_v60).length));
        (_290_v60)[_index53] = true;
      } else if (_source2.is_DC0) {
        let _291___mcc_h6 = (_source2).cf0;
        let _292_cf0 = _291___mcc_h6;
        let _293_v61;
        _293_v61 = _dafny.Set.fromElements(_211_v3, _211_v3);
        (globalState).f6 = (_293_v61).IsProperSubsetOf(_293_v61);
        if (_this.f23) {
          let _294_v62;
          _294_v62 = _dafny.Seq.of(_292_cf0, _292_cf0);
          let _295_v63;
          _295_v63 = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,_294_v62);
          let _296_v64;
          _296_v64 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(923)), ((_297_cf0) => function (_298_i5) {
            return _297_cf0;
          })(_292_cf0)), _294_v62, _dafny.Seq.update((((_295_v63).contains(true)) ? ((_295_v63).get(true)) : (_294_v62)), _module.__default.safeIndex(_292_cf0, new BigNumber(((((_295_v63).contains(true)) ? ((_295_v63).get(true)) : (_294_v62))).length)), new BigNumber(279)));
          let _299_v65;
          _299_v65 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,_this.f23);
          let _300_v66;
          _300_v66 = _dafny.Map.Empty.slice().updateUnsafe(_296_v64,(_299_v65).Merge(_299_v65));
          _300_v66 = (_300_v66).update((_296_v64).Union(_296_v64), _299_v65);
          let _301_v67;
          _301_v67 = _dafny.Seq.UnicodeFromString("jowiusrj");
          let _302_v68;
          let _nw54 = Array((new BigNumber(15)).toNumber()).fill(false);
          _302_v68 = _nw54;
          let _index54 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_302_v68).length));
          (_302_v68)[_index54] = (_this).f22;
          let _index55 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_302_v68).length));
          let _rhs50 = _dafny.Seq.update(((_this.f23) ? (_301_v67) : (_301_v67)), _module.__default.safeIndex(new BigNumber(909), new BigNumber((((_this.f23) ? (_301_v67) : (_301_v67))).length)), _208_v0);
          let _rhs51 = _208_v0;
          let _rhs52 = _this.f23;
          let _lhs41 = globalState;
          let _lhs42 = _302_v68;
          let _lhs43 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_302_v68).length));
          _301_v67 = _rhs50;
          _lhs41.f18 = _rhs51;
          _lhs42[_lhs43] = _rhs52;
          let _303_v69;
          _303_v69 = _dafny.MultiSet.fromElements(_this.f23);
          r0 = !(new BigNumber((_301_v67).length)).isEqualTo(((((_303_v69).contains(_this.f23)) ? ((_303_v69).get(_this.f23)) : (p0))).plus(_292_cf0));
          _292_cf0 = _292_cf0;
          let _index56 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_302_v68).length));
          (_302_v68)[_index56] = (_this).f22;
        } else {
          let _304_v70;
          _304_v70 = _dafny.Seq.of(p0, p0, _292_cf0, p0, new BigNumber(850));
          _292_cf0 = _module.__default.safeDivisionInt(new BigNumber(700), new BigNumber((_dafny.Seq.Concat(_304_v70, _304_v70)).length));
          let _305_v71;
          let _nw55 = Array((new BigNumber(5)).toNumber()).fill(false);
          _305_v71 = _nw55;
          _305_v71 = _305_v71;
          _292_cf0 = (p0).multipliedBy(_292_cf0);
          _292_cf0 = (p0).multipliedBy(_292_cf0);
          let _306_v72;
          _306_v72 = _dafny.MultiSet.fromElements(_this.f23, true);
          let _307_v73;
          let _nw56 = Array((new BigNumber(2)).toNumber()).fill([]);
          _307_v73 = _nw56;
          let _308_v74;
          let _init6 = ((_309_p0) => function (_310_i6) {
            return (_310_i6).plus(_309_p0);
          })(p0);
          let _nw57 = Array((new BigNumber(2)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw57.length); _i0_6++) {
            _nw57[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _308_v74 = _nw57;
          let _index57 = _module.__default.safeIndex(new BigNumber(517), new BigNumber((_307_v73).length));
          (_307_v73)[_index57] = _308_v74;
          let _index58 = _module.__default.safeIndex(new BigNumber(517), new BigNumber((_307_v73).length));
          let _rhs53 = _306_v72;
          let _rhs54 = _292_cf0;
          let _rhs55 = _213_v5;
          let _rhs56 = _308_v74;
          let _rhs57 = _305_v71;
          let _lhs44 = _307_v73;
          let _lhs45 = _module.__default.safeIndex(new BigNumber(517), new BigNumber((_307_v73).length));
          _306_v72 = _rhs53;
          _292_cf0 = _rhs54;
          _213_v5 = _rhs55;
          _lhs44[_lhs45] = _rhs56;
          _305_v71 = _rhs57;
        }
        (globalState).f19 = (((_this).f22) ? (_this.f23) : ((_this).f22));
        let _311_v75;
        let _nw58 = Array((new BigNumber(6)).toNumber()).fill(false);
        _311_v75 = _nw58;
        let _index59 = _module.__default.safeIndex(new BigNumber(161), new BigNumber((_311_v75).length));
        (_311_v75)[_index59] = false;
        let _index60 = _module.__default.safeIndex(new BigNumber(161), new BigNumber((_311_v75).length));
        (_311_v75)[_index60] = (_module.__default.fm0(_this.f23, (_this).fm4(new BigNumber((_211_v3).length), new BigNumber((_dafny.Set.fromElements((_this).f22, (_this).f22)).length), !((_this).f22), globalState), globalState)).isLessThan(new BigNumber((_dafny.MultiSet.fromElements(_this.f23)).cardinality()));
      } else {
        let _312___mcc_h7 = (_source2).cf7;
        let _313_cf7 = _312___mcc_h7;
        let _314_v76;
        _314_v76 = _dafny.Seq.UnicodeFromString("rfhewcmh");
        if (((!(_this.f23)) ? (!(_dafny.areEqual(_314_v76, _dafny.Seq.Create(_module.__default.abs(new BigNumber(754)), ((_315_v0) => function (_316_i7) {
          return _315_v0;
        })(_208_v0))))) : (_this.f23))) {
          let _317_v77;
          _317_v77 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(945),_211_v3);
          let _318_v78;
          let _nw59 = new _module.C0();
          _nw59.__ctor(false, (_317_v77).Merge(_317_v77));
          _318_v78 = _nw59;
          let _319_v79;
          let _320_v80;
          let _out8;
          let _out9;
          let _outcollector4 = (_this).m1(p0, (_this).f22, (p0).isLessThanOrEqualTo((_dafny.ZERO).minus(p0)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(790)), ((_321_v0) => function (_322_i8) {
            return _321_v0;
          })(_208_v0))).length), globalState);
          _out8 = _outcollector4[0];
          _out9 = _outcollector4[1];
          _319_v79 = _out8;
          _320_v80 = _out9;
          let _323_v81;
          let _nw60 = new _module.C0();
          _nw60.__ctor((_this).fm4(p0, p0, _318_v78.f24, globalState), _317_v77);
          _323_v81 = _nw60;
          let _324_v82;
          _324_v82 = _dafny.MultiSet.fromElements(p0, _module.__default.safeModuloInt(_module.__default.fm0(_323_v81.f24, _320_v80, globalState), p0));
          let _rhs58 = _318_v78;
          let _rhs59 = (_324_v82).update((_213_v5).dtor_cf2, _module.__default.abs(p0));
          let _rhs60 = !((_this).f22);
          let _rhs61 = _208_v0;
          let _rhs62 = (_this).fm4(new BigNumber((_314_v76).length), p0, !(_323_v81.f24), globalState);
          let _lhs46 = globalState;
          let _lhs47 = globalState;
          let _lhs48 = _this;
          _323_v81 = _rhs58;
          _324_v82 = _rhs59;
          _lhs46.f19 = _rhs60;
          _lhs47.f18 = _rhs61;
          _lhs48.f23 = _rhs62;
          _319_v79 = _314_v76;
          let _325_v83;
          _325_v83 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,new BigNumber(271));
          r1 = _dafny.Map.Empty.slice().updateUnsafe(_325_v83,_dafny.Seq.Create(_module.__default.abs(new BigNumber(306)), ((_326_v0) => function (_327_i9) {
            return _326_v0;
          })(_208_v0)));
        } else {
          let _328_v84;
          _328_v84 = new BigNumber(-212);
          _328_v84 = _328_v84;
          let _329_v85;
          let _init7 = ((_330_p0) => function (_331_i10) {
            return (_331_i10).multipliedBy(_330_p0);
          })(p0);
          let _nw61 = Array((new BigNumber(21)).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw61.length); _i0_7++) {
            _nw61[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _329_v85 = _nw61;
          let _332_v86;
          _332_v86 = _dafny.Seq.of(p0);
          let _index61 = _module.__default.safeIndex(new BigNumber(457), new BigNumber((_329_v85).length));
          (_329_v85)[_index61] = _module.__default.safeModuloInt((_dafny.ZERO).minus(p0), new BigNumber((_332_v86).length));
          let _333_v87;
          _333_v87 = _329_v85;
          let _index62 = _module.__default.safeIndex(new BigNumber(457), new BigNumber((_329_v85).length));
          let _rhs63 = (_328_v84).multipliedBy((((_this).f22) ? (p0) : (new BigNumber(407))));
          let _rhs64 = (_333_v87);
          let _rhs65 = new BigNumber(35);
          let _lhs49 = _329_v85;
          let _lhs50 = _module.__default.safeIndex(new BigNumber(457), new BigNumber((_329_v85).length));
          _lhs49[_lhs50] = _rhs63;
          _329_v85 = _rhs64;
          _328_v84 = _rhs65;
          let _334_v88;
          _334_v88 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("gjfewtn"),(_this).f22);
          _334_v88 = (_334_v88).update(_314_v76, (_this).f22);
          let _335_v89;
          let _nw62 = Array((new BigNumber(4)).toNumber());
          _nw62[(_dafny.ZERO).toNumber()] = _this.f23;
          _nw62[(_dafny.ONE).toNumber()] = _this.f23;
          _nw62[(new BigNumber(2)).toNumber()] = !((_this).f22);
          _nw62[(new BigNumber(3)).toNumber()] = _this.f23;
          _335_v89 = _nw62;
          let _336_v90;
          _336_v90 = _dafny.Set.fromElements(_213_v5);
          let _index63 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_335_v89).length));
          (_335_v89)[_index63] = (_336_v90).IsSubsetOf(_336_v90);
          let _index64 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_335_v89).length));
          (_335_v89)[_index64] = (_this).fm4(_328_v84, _module.__default.fm0(!((_this).f22), _this.f23, globalState), (_this).f22, globalState);
          (globalState).f6 = _this.f23;
        }
        let _337_v91;
        _337_v91 = _dafny.Map.Empty.slice().updateUnsafe(((_213_v5).dtor_cf2).isLessThan(p0),_314_v76);
        _337_v91 = _337_v91;
        let _338_v92;
        _338_v92 = _module.D3.create_DC6(_314_v76);
        let _339_v93;
        _339_v93 = _dafny.Seq.of(_314_v76, _314_v76, _314_v76, _dafny.Seq.UnicodeFromString("j"), (_338_v92).dtor_cf10);
        if (_dafny.Seq.IsPrefixOf(_314_v76, (_339_v93)[_module.__default.safeIndex(new BigNumber(-908), new BigNumber((_339_v93).length))])) {
          r0 = !((_this).f22);
          (globalState).f19 = _this.f23;
          let _340_v94;
          _340_v94 = _dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.Create(_module.__default.abs(new BigNumber(20)), ((_341_p0) => function (_342_i11) {
            return _341_p0;
          })(p0)));
          let _343_v95;
          _343_v95 = _dafny.Seq.of(p0);
          let _344_v96;
          _344_v96 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((((_340_v94).contains((_this).f22)) ? ((_340_v94).get((_this).f22)) : (_343_v95))).length),_211_v3);
          let _345_v97;
          let _nw63 = new _module.C0();
          _nw63.__ctor(_this.f23, _344_v96);
          _345_v97 = _nw63;
          (globalState).f6 = _this.f23;
          let _346_v98;
          let _nw64 = Array((new BigNumber(5)).toNumber()).fill(false);
          _346_v98 = _nw64;
          let _347_v99;
          _347_v99 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_346_v98,_345_v97.f24),new BigNumber((_dafny.MultiSet.fromElements(p0)).cardinality()));
          let _348_v100;
          _348_v100 = _dafny.Map.Empty.slice().updateUnsafe(_346_v98,_this.f23);
          _347_v99 = (_347_v99).update(_348_v100, p0);
        } else {
          (globalState).f6 = (_this).f22;
          (globalState).f6 = !(((_this).f22) === (_this.f23));
          let _349_v101;
          _349_v101 = _module.D0.create_DC0(p0);
          (globalState).f19 = (p0).isLessThanOrEqualTo((_349_v101).dtor_cf0);
          (globalState).f6 = (_this).f22;
          let _350_v102;
          _350_v102 = _dafny.Map.Empty.slice().updateUnsafe(p0,_211_v3);
          let _351_v103;
          let _nw65 = new _module.C0();
          _nw65.__ctor(false, _350_v102);
          _351_v103 = _nw65;
          let _352_v104;
          _352_v104 = _dafny.Seq.of(_351_v103, _351_v103, _351_v103, _351_v103);
          let _353_v105;
          _353_v105 = _dafny.MultiSet.fromElements(p0);
          let _354_v106;
          _354_v106 = _dafny.Set.fromElements(p0);
          let _355_v107;
          _355_v107 = _dafny.Map.Empty.slice().updateUnsafe(_351_v103.f24,new BigNumber((_354_v106).length));
          let _356_v108;
          let _nw66 = Array((new BigNumber(29)).toNumber());
          _nw66[(_dafny.ZERO).toNumber()] = p0;
          _nw66[(_dafny.ONE).toNumber()] = new BigNumber(173);
          _nw66[(new BigNumber(2)).toNumber()] = p0;
          _nw66[(new BigNumber(3)).toNumber()] = p0;
          _nw66[(new BigNumber(4)).toNumber()] = new BigNumber(978);
          _nw66[(new BigNumber(5)).toNumber()] = p0;
          _nw66[(new BigNumber(6)).toNumber()] = p0;
          _nw66[(new BigNumber(7)).toNumber()] = p0;
          _nw66[(new BigNumber(8)).toNumber()] = p0;
          _nw66[(new BigNumber(9)).toNumber()] = p0;
          _nw66[(new BigNumber(10)).toNumber()] = new BigNumber((_352_v104).length);
          _nw66[(new BigNumber(11)).toNumber()] = _module.__default.safeModuloInt(p0, p0);
          _nw66[(new BigNumber(12)).toNumber()] = (p0).plus(p0);
          _nw66[(new BigNumber(13)).toNumber()] = p0;
          _nw66[(new BigNumber(14)).toNumber()] = p0;
          _nw66[(new BigNumber(15)).toNumber()] = ((_dafny.ZERO).minus(p0)).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(929)), ((_357_v0) => function (_358_i12) {
            return _357_v0;
          })(_208_v0))).length));
          _nw66[(new BigNumber(16)).toNumber()] = p0;
          _nw66[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus(p0);
          _nw66[(new BigNumber(18)).toNumber()] = new BigNumber(((_353_v105).update(p0, _module.__default.abs(new BigNumber(-355)))).cardinality());
          _nw66[(new BigNumber(19)).toNumber()] = p0;
          _nw66[(new BigNumber(20)).toNumber()] = new BigNumber((_dafny.Seq.update(_211_v3, _module.__default.safeIndex(new BigNumber(-292), new BigNumber((_211_v3).length)), _351_v103.f24)).length);
          _nw66[(new BigNumber(21)).toNumber()] = p0;
          _nw66[(new BigNumber(22)).toNumber()] = p0;
          _nw66[(new BigNumber(23)).toNumber()] = p0;
          _nw66[(new BigNumber(24)).toNumber()] = p0;
          _nw66[(new BigNumber(25)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(-346), p0);
          _nw66[(new BigNumber(26)).toNumber()] = _module.__default.safeModuloInt(p0, p0);
          _nw66[(new BigNumber(27)).toNumber()] = p0;
          _nw66[(new BigNumber(28)).toNumber()] = (((_355_v107).contains(_351_v103.f24)) ? ((_355_v107).get(_351_v103.f24)) : (new BigNumber((_314_v76).length)));
          _356_v108 = _nw66;
          let _index65 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_356_v108).length));
          (_356_v108)[_index65] = p0;
          let _359_v109;
          _359_v109 = _dafny.Seq.of(new BigNumber(-72), p0, (_dafny.ZERO).minus(p0), p0);
          let _index66 = _module.__default.safeIndex(new BigNumber(594), new BigNumber((_356_v108).length));
          (_356_v108)[_index66] = new BigNumber((_359_v109).length);
          let _360_v110;
          _360_v110 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f22);
          let _index67 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_356_v108).length));
          let _index68 = _module.__default.safeIndex(new BigNumber(594), new BigNumber((_356_v108).length));
          let _rhs66 = (((_this).f22) ? (!_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(p0))).length), p0), new BigNumber((_dafny.Seq.update(_314_v76, _module.__default.safeIndex(p0, new BigNumber((_314_v76).length)), _208_v0)).length))) : (_351_v103.f24));
          let _rhs67 = p0;
          let _rhs68 = new BigNumber((_360_v110).length);
          let _rhs69 = (((_this).fm4((_349_v101).dtor_cf0, p0, _this.f23, globalState)) ? (_module.__default.fm2(_dafny.MultiSet.FromArray(_314_v76), globalState)) : (_dafny.Seq.of(p0)));
          let _lhs51 = globalState;
          let _lhs52 = _356_v108;
          let _lhs53 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_356_v108).length));
          let _lhs54 = _356_v108;
          let _lhs55 = _module.__default.safeIndex(new BigNumber(594), new BigNumber((_356_v108).length));
          _lhs51.f19 = _rhs66;
          _lhs52[_lhs53] = _rhs67;
          _lhs54[_lhs55] = _rhs68;
          _359_v109 = _rhs69;
        }
        r0 = (!(true)) && ((p0).isLessThan(new BigNumber(75)));
      }
      _208_v0 = _module.__default.fm5(globalState);
      (globalState).f18 = _208_v0;
      if ((((_this).f22) ? (true) : (_dafny.Seq.IsPrefixOf(_211_v3, _211_v3)))) {
        let _361_v111;
        _361_v111 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0((_this).fm4(p0, p0, (_this).f22, globalState), false, globalState),_dafny.Seq.update(_211_v3, _module.__default.safeIndex(p0, new BigNumber((_211_v3).length)), _this.f23));
        let _362_v112;
        let _nw67 = new _module.C0();
        _nw67.__ctor(false, _361_v111);
        _362_v112 = _nw67;
        let _363_v113;
        _363_v113 = _module.D3.create_DC6(_module.__default.fm3(p0, globalState));
        let _364_v114;
        _364_v114 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,_363_v113);
        _364_v114 = (_364_v114).update(_362_v112.f24, _363_v113);
        (globalState).f6 = _362_v112.f24;
        let _365_v115;
        let _init8 = ((_366_v0) => function (_367_i13) {
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(132)), ((_368_v0) => function (_369_i14) {
            return _368_v0;
          })(_366_v0));
        })(_208_v0);
        let _nw68 = Array((new BigNumber(17)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw68.length); _i0_8++) {
          _nw68[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _365_v115 = _nw68;
        let _index69 = _module.__default.safeIndex(new BigNumber(867), new BigNumber((_365_v115).length));
        (_365_v115)[_index69] = (_363_v113).dtor_cf10;
        let _370_v116;
        _370_v116 = _dafny.Seq.UnicodeFromString("ccagyoi");
        let _index70 = _module.__default.safeIndex(new BigNumber(867), new BigNumber((_365_v115).length));
        (_365_v115)[_index70] = _dafny.Seq.Concat(_370_v116, (_363_v113).dtor_cf10);
        let _371_v117;
        let _nw69 = new _module.C0();
        _nw69.__ctor(!(p0).isEqualTo(p0), (_362_v112).f25);
        _371_v117 = _nw69;
      } else {
        _208_v0 = _208_v0;
        let _372_v118;
        _372_v118 = _dafny.Seq.UnicodeFromString("sbspy");
        (_this).f23 = (new BigNumber((_372_v118).length)).isLessThan(new BigNumber((_372_v118).length));
        let _373_v119;
        _373_v119 = _dafny.MultiSet.fromElements(_372_v118, _372_v118, _dafny.Seq.Concat(_372_v118, _dafny.Seq.Create(_module.__default.abs(new BigNumber(469)), ((_374_v0) => function (_375_i15) {
          return _374_v0;
        })(_208_v0))));
        let _376_v120;
        _376_v120 = _dafny.MultiSet.fromElements(p0);
        let _rhs70 = !((new BigNumber(582)).minus(p0)).isEqualTo((_dafny.ZERO).minus(p0));
        let _rhs71 = false;
        let _rhs72 = (_376_v120).IsProperSubsetOf((_376_v120).Intersect(_dafny.MultiSet.fromElements(p0)));
        let _rhs73 = _373_v119;
        let _lhs56 = globalState;
        let _lhs57 = globalState;
        let _lhs58 = globalState;
        _lhs56.f6 = _rhs70;
        _lhs57.f6 = _rhs71;
        _lhs58.f19 = _rhs72;
        _373_v119 = _rhs73;
        let _377_v121;
        let _nw70 = Array((new BigNumber(9)).toNumber()).fill(false);
        _377_v121 = _nw70;
        let _index71 = _module.__default.safeIndex(new BigNumber(971), new BigNumber((_377_v121).length));
        (_377_v121)[_index71] = (_this).f22;
        let _378_v122;
        let _nw71 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
        _378_v122 = _nw71;
        let _index72 = _module.__default.safeIndex(new BigNumber(971), new BigNumber((_377_v121).length));
        let _rhs74 = _378_v122;
        let _rhs75 = _this.f23;
        let _lhs59 = globalState;
        let _lhs60 = _377_v121;
        let _lhs61 = _module.__default.safeIndex(new BigNumber(971), new BigNumber((_377_v121).length));
        _lhs59.f13 = _rhs74;
        _lhs60[_lhs61] = _rhs75;
        let _379_v123;
        _379_v123 = _dafny.Seq.of(p0);
        let _380_v124;
        _380_v124 = _dafny.Map.Empty.slice().updateUnsafe((_377_v121)[_module.__default.safeIndex(new BigNumber(971), new BigNumber((_377_v121).length))],(_379_v123)[_module.__default.safeIndex(p0, new BigNumber((_379_v123).length))]);
        _380_v124 = (_380_v124).update(_this.f23, (((_this).f22) ? (p0) : (p0)));
      }
      let _381_v125;
      _381_v125 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_211_v3).length),_211_v3);
      let _382_v126;
      let _nw72 = new _module.C0();
      _nw72.__ctor((_this).f22, _381_v125);
      _382_v126 = _nw72;
      let _383_v127;
      _383_v127 = _dafny.Seq.of(_382_v126);
      let _384_v128;
      let _nw73 = Array((new BigNumber(8)).toNumber());
      _nw73[(_dafny.ZERO).toNumber()] = (_this).f22;
      _nw73[(_dafny.ONE).toNumber()] = false;
      _nw73[(new BigNumber(2)).toNumber()] = (_this).f22;
      _nw73[(new BigNumber(3)).toNumber()] = !(((_this).f22) === ((_this).fm4(new BigNumber((_383_v127).length), p0, (_this).f22, globalState)));
      _nw73[(new BigNumber(4)).toNumber()] = _this.f23;
      _nw73[(new BigNumber(5)).toNumber()] = (_this).f22;
      _nw73[(new BigNumber(6)).toNumber()] = _this.f23;
      _nw73[(new BigNumber(7)).toNumber()] = (p0).isEqualTo(p0);
      _384_v128 = _nw73;
      let _index73 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_384_v128).length));
      (_384_v128)[_index73] = (_this).f22;
      let _index74 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_384_v128).length));
      (_384_v128)[_index74] = ((p0).plus((_dafny.ZERO).minus(p0))).isEqualTo(p0);
      r0 = (_this).f22;
      let _385_v129;
      _385_v129 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,p0);
      r1 = _dafny.Map.Empty.slice().updateUnsafe((_385_v129).Merge(_385_v129),_dafny.Seq.Create(_module.__default.abs(new BigNumber(961)), ((_386_v0) => function (_387_i16) {
        return _386_v0;
      })(_208_v0)));
      return [r0, r1];
    }
    m1(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = false;
      (globalState).f0 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
      let _388_v0;
      let _nw74 = Array((new BigNumber(29)).toNumber());
      _388_v0 = _nw74;
      _388_v0 = _388_v0;
      if (_this.f23) {
        let _389_v1;
        _389_v1 = _dafny.Map.Empty.slice().updateUnsafe(_this.f23,p2);
        _389_v1 = _dafny.Map.Empty.slice().updateUnsafe(p2,!(true) || ((_this).f22));
        let _390_v2;
        _390_v2 = _dafny.Seq.UnicodeFromString("ahbq");
        let _391_v3;
        _391_v3 = new _dafny.CodePoint('h'.codePointAt(0));
        let _392_v4;
        _392_v4 = _dafny.Seq.of(_dafny.Seq.IsPrefixOf(_390_v2, _dafny.Seq.update(_390_v2, _module.__default.safeIndex(p3, new BigNumber((_390_v2).length)), _391_v3)));
        _392_v4 = _dafny.Seq.Concat(_392_v4, _392_v4);
        let _393_v5;
        let _nw75 = new _module.C0();
        _nw75.__ctor(p1, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(789),_392_v4));
        _393_v5 = _nw75;
        let _394_v6;
        _394_v6 = new BigNumber(-403);
        let _395_v7;
        _395_v7 = _dafny.Seq.of(_394_v6);
        let _rhs76 = _393_v5;
        let _rhs77 = _module.__default.safeModuloInt(p3, (_dafny.ZERO).minus(new BigNumber((_395_v7).length)));
        _393_v5 = _rhs76;
        _394_v6 = _rhs77;
        (_393_v5).f24 = _this.f23;
        let _396_v9;
        let _nw76 = Array((new BigNumber(25)).toNumber());
        _nw76[(_dafny.ZERO).toNumber()] = (_this).f22;
        _nw76[(_dafny.ONE).toNumber()] = true;
        _nw76[(new BigNumber(2)).toNumber()] = !(!(p1));
        _nw76[(new BigNumber(3)).toNumber()] = _393_v5.f24;
        _nw76[(new BigNumber(4)).toNumber()] = p1;
        _nw76[(new BigNumber(5)).toNumber()] = true;
        _nw76[(new BigNumber(6)).toNumber()] = false;
        _nw76[(new BigNumber(7)).toNumber()] = p2;
        _nw76[(new BigNumber(8)).toNumber()] = p2;
        _nw76[(new BigNumber(9)).toNumber()] = _this.f23;
        _nw76[(new BigNumber(10)).toNumber()] = false;
        _nw76[(new BigNumber(11)).toNumber()] = p1;
        _nw76[(new BigNumber(12)).toNumber()] = _393_v5.f24;
        _nw76[(new BigNumber(13)).toNumber()] = p2;
        _nw76[(new BigNumber(14)).toNumber()] = (_this).f22;
        _nw76[(new BigNumber(15)).toNumber()] = (_this).f22;
        _nw76[(new BigNumber(16)).toNumber()] = _this.f23;
        _nw76[(new BigNumber(17)).toNumber()] = (_this).f22;
        _nw76[(new BigNumber(18)).toNumber()] = _393_v5.f24;
        _nw76[(new BigNumber(19)).toNumber()] = p1;
        _nw76[(new BigNumber(20)).toNumber()] = true;
        _nw76[(new BigNumber(21)).toNumber()] = p1;
        _nw76[(new BigNumber(22)).toNumber()] = p2;
        _nw76[(new BigNumber(23)).toNumber()] = p1;
        _nw76[(new BigNumber(24)).toNumber()] = (_this).fm4(new BigNumber((function () {
          let _coll4 = new _dafny.Map();
          for (const _compr_4 of _dafny.IntegerRange(new BigNumber(36), new BigNumber(816))) {
            let _397_v8 = _compr_4;
            if (((new BigNumber(36)).isLessThanOrEqualTo(_397_v8)) && ((_397_v8).isLessThan(new BigNumber(816)))) {
              _coll4.push([(_397_v8).minus(new BigNumber((_390_v2).length)),true]);
            }
          }
          return _coll4;
        }()).length), p3, !(p2), globalState);
        _396_v9 = _nw76;
        let _398_v10;
        _398_v10 = _dafny.MultiSet.fromElements(_396_v9);
        _394_v6 = new BigNumber(((_398_v10).Union(_398_v10)).cardinality());
      } else {
        let _399_v11;
        _399_v11 = new BigNumber(478);
        let _400_v12;
        _400_v12 = _dafny.MultiSet.fromElements(p3);
        let _401_v13;
        _401_v13 = _dafny.MultiSet.fromElements(_this.f23);
        let _402_v14;
        _402_v14 = _dafny.Map.Empty.slice().updateUnsafe(_401_v13,_this.f23);
        let _403_v15;
        _403_v15 = _dafny.Map.Empty.slice().updateUnsafe(p2,_399_v11);
        let _404_v17;
        _404_v17 = _dafny.Seq.of(new BigNumber((_403_v15).length), new BigNumber((function () {
          let _coll5 = new _dafny.Map();
          for (const _compr_5 of _dafny.IntegerRange(new BigNumber(-789), new BigNumber(579))) {
            let _405_v16 = _compr_5;
            if (((new BigNumber(-789)).isLessThanOrEqualTo(_405_v16)) && ((_405_v16).isLessThan(new BigNumber(579)))) {
              _coll5.push([(_405_v16).multipliedBy(new BigNumber(-899)),p0]);
            }
          }
          return _coll5;
        }()).length), p0, _399_v11);
        let _406_v18;
        _406_v18 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,false)).length), new BigNumber((_402_v14).length), (_404_v17)[_module.__default.safeIndex(p0, new BigNumber((_404_v17).length))]);
        _399_v11 = (new BigNumber((_400_v12).cardinality())).plus((p3).multipliedBy(new BigNumber((_406_v18).length)));
        let _407_v19;
        _407_v19 = _dafny.Set.fromElements(p0, p3, p3, _399_v11, (_dafny.ZERO).minus(_module.__default.safeModuloInt(_399_v11, p0)));
        _407_v19 = _dafny.Set.fromElements(p3);
        let _408_v20;
        _408_v20 = _module.D0.create_DC2((_this).f22, new BigNumber(-817), _399_v11, p1);
        let _409_v21;
        _409_v21 = _dafny.Seq.of(_408_v20, _408_v20, _408_v20, _408_v20);
        let _410_v23;
        _410_v23 = _dafny.Map.Empty.slice().updateUnsafe((_this).f22,p2);
        let _411_v24;
        _411_v24 = _dafny.Seq.of(true);
        let _rhs78 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(999)), ((_412_v20) => function (_413_i0) {
          return _412_v20;
        })(_408_v20));
        let _rhs79 = p2;
        let _rhs80 = new BigNumber((_dafny.Seq.Concat(((false) ? (_dafny.Seq.of((_this).fm4(new BigNumber(((function () {
          let _coll6 = new _dafny.Map();
          for (const _compr_6 of _dafny.IntegerRange(new BigNumber(263), new BigNumber(685))) {
            let _414_v22 = _compr_6;
            if (((new BigNumber(263)).isLessThanOrEqualTo(_414_v22)) && ((_414_v22).isLessThan(new BigNumber(685)))) {
              _coll6.push([_module.__default.safeModuloInt(_414_v22, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_399_v11,_399_v11)).length)),p3]);
            }
          }
          return _coll6;
        }()).update(p0, new BigNumber((_410_v23).length))).length), _399_v11, true, globalState), p2, p2, false)) : (_411_v24)), _dafny.Seq.Concat(_411_v24, _dafny.Seq.of((_this).f22)))).length);
        let _lhs62 = _this;
        _409_v21 = _rhs78;
        _lhs62.f23 = _rhs79;
        _399_v11 = _rhs80;
        let _415_v25;
        _415_v25 = _dafny.Map.Empty.slice().updateUnsafe(p0,_module.__default.fm10(p3, _399_v11, globalState));
        let _416_v26;
        let _nw77 = new _module.C0();
        _nw77.__ctor(p2, _415_v25);
        _416_v26 = _nw77;
        let _417_v27;
        let _init9 = ((_418_p3) => function (_419_i1) {
          return _module.__default.safeDivisionInt(_419_i1, _418_p3);
        })(p3);
        let _nw78 = Array((new BigNumber(17)).toNumber());
        for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw78.length); _i0_9++) {
          _nw78[_i0_9] = _init9(new BigNumber(_i0_9));
        }
        _417_v27 = _nw78;
        let _420_v28;
        let _nw79 = Array((new BigNumber(13)).toNumber());
        _nw79[(_dafny.ZERO).toNumber()] = _417_v27;
        _nw79[(_dafny.ONE).toNumber()] = _417_v27;
        _nw79[(new BigNumber(2)).toNumber()] = _417_v27;
        _nw79[(new BigNumber(3)).toNumber()] = _417_v27;
        _nw79[(new BigNumber(4)).toNumber()] = _417_v27;
        _nw79[(new BigNumber(5)).toNumber()] = _417_v27;
        _nw79[(new BigNumber(6)).toNumber()] = _417_v27;
        _nw79[(new BigNumber(7)).toNumber()] = _417_v27;
        _nw79[(new BigNumber(8)).toNumber()] = _417_v27;
        _nw79[(new BigNumber(9)).toNumber()] = _417_v27;
        _nw79[(new BigNumber(10)).toNumber()] = _417_v27;
        _nw79[(new BigNumber(11)).toNumber()] = _417_v27;
        _nw79[(new BigNumber(12)).toNumber()] = _417_v27;
        _420_v28 = _nw79;
        let _index75 = _module.__default.safeIndex(new BigNumber(390), new BigNumber((_420_v28).length));
        (_420_v28)[_index75] = _417_v27;
        let _index76 = _module.__default.safeIndex(new BigNumber(390), new BigNumber((_420_v28).length));
        (_420_v28)[_index76] = _417_v27;
      }
      let _421_v29;
      let _init10 = ((_422_p0) => function (_423_i2) {
        return _module.D0.create_DC3(_module.D0.create_DC0(_422_p0));
      })(p0);
      let _nw80 = Array((new BigNumber(13)).toNumber());
      for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw80.length); _i0_10++) {
        _nw80[_i0_10] = _init10(new BigNumber(_i0_10));
      }
      _421_v29 = _nw80;
      let _424_v30;
      _424_v30 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),(_dafny.ZERO).minus(p0));
      let _425_v31;
      _425_v31 = _dafny.Map.Empty.slice().updateUnsafe(_424_v30,_421_v29);
      let _426_v32;
      _426_v32 = _dafny.Seq.of(_421_v29, _421_v29, _421_v29);
      let _427_v33;
      _427_v33 = _dafny.Seq.of(_this.f23, p2);
      let _428_v34;
      _428_v34 = _dafny.Map.Empty.slice().updateUnsafe(p0,_427_v33);
      let _429_v35;
      let _nw81 = new _module.C0();
      _nw81.__ctor(p1, _428_v34);
      _429_v35 = _nw81;
      let _430_v36;
      _430_v36 = _dafny.Map.Empty.slice().updateUnsafe(p2,_429_v35);
      _421_v29 = (((_425_v31).contains(_424_v30)) ? ((_425_v31).get(_424_v30)) : ((_426_v32)[_module.__default.safeIndex(new BigNumber((_430_v36).length), new BigNumber((_426_v32).length))]));
      let _431_v37;
      let _nw82 = Array((new BigNumber(2)).toNumber()).fill(false);
      _431_v37 = _nw82;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_431_v37).length))) {
        let _432_i3 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_432_i3)) && ((_432_i3).isLessThan(new BigNumber((_431_v37).length))))) {
          (_431_v37)[(_432_i3)] = _dafny.areEqual(_dafny.Seq.of(_427_v33), _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(true)), _dafny.Seq.of(_427_v33)));
        }
      }
      _431_v37 = _431_v37;
      let _433_v38;
      _433_v38 = _dafny.Seq.UnicodeFromString("crax");
      r0 = _433_v38;
      r1 = (p2) || (p1);
      return [r0, r1];
    }
    get f22() {
      let _this = this;
      return _this._f22;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
