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
    static fm1(globalState) {
      return ((new BigNumber((_dafny.MultiSet.fromElements(!(true), true)).cardinality())).plus(new BigNumber((_dafny.Seq.UnicodeFromString("y")).length))).isLessThan((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-257)))));
    };
    static fm2(p0, globalState) {
      return ((new BigNumber((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(98), new BigNumber(163))) {
          let _0_v0 = _compr_0;
          if (((new BigNumber(98)).isLessThanOrEqualTo(_0_v0)) && ((_0_v0).isLessThan(new BigNumber(163)))) {
            _coll0.push([_module.__default.safeDivisionInt(_0_v0, new BigNumber(131)),false]);
          }
        }
        return _coll0;
      }()).length)).minus(new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-722),true), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(393),false), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-391)), function (_1_i0) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      })).length),false), function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(122),new BigNumber(265))).Keys.Elements) {
          let _2_v1 = _compr_1;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(122),new BigNumber(265))).contains(_2_v1)) {
            _coll1.push([(_2_v1).minus(new BigNumber(798)),true]);
          }
        }
        return _coll1;
      }(), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-223)), function (_3_i1) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      })).length),true))).length))).minus(new BigNumber(117));
    };
    static fm5(p0, p1, p2, globalState) {
      return (function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of (_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false)).length),true)).length))).cardinality()))).Elements) {
          let _4_v0 = _compr_2;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false)).length),true)).length))).cardinality())), _4_v0)) {
            _coll2.push([(_4_v0).multipliedBy(new BigNumber(802)),new BigNumber(-596)]);
          }
        }
        return _coll2;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(567),new BigNumber((_dafny.Seq.of(_dafny.Seq.of(true, true), _dafny.Seq.of(false, false), _dafny.Seq.of(true))).length)));
    };
    static fm6(globalState) {
      if ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(837)), function (_5_i1) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      })).length), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber(47))).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("okfva")).length))))).cardinality()), new BigNumber(-567), new BigNumber(792), new BigNumber(-57))).IsSubsetOf(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-769)), function (_6_i0) {
        return new BigNumber(985);
      })).length), new BigNumber(-713)))) {
        return new _dafny.CodePoint('l'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('k'.codePointAt(0));
      }
    };
    static fm7(p0, p1, p2, globalState) {
      return _module.D1.create_DC4(!(false));
    };
    static fm8(p0, p1, p2, p3, globalState) {
      return function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of (_dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC9(new _dafny.CodePoint('k'.codePointAt(0))),false)).Keys.Elements) {
          let _7_v0 = _compr_3;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC9(new _dafny.CodePoint('k'.codePointAt(0))),false)).contains(_7_v0)) {
            _coll3.push([_7_v0,new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ev"), _dafny.Seq.UnicodeFromString("sbujxi"))).length)]);
          }
        }
        return _coll3;
      }();
    };
    static fm9(p0, p1, globalState) {
      if (true) {
        return (_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,true))).Union(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,true), _dafny.Map.Empty.slice().updateUnsafe(false,(_module.D2.create_DC10(false, !(false))).dtor_cf20)));
      } else {
        return (_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,false))).Union(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,true), _dafny.Map.Empty.slice().updateUnsafe(false,false)));
      }
    };
    static fm10(p0, p1, globalState) {
      return _dafny.MultiSet.fromElements(!((new BigNumber(772)).isLessThan(new BigNumber(-677))), (_dafny.Set.fromElements(_dafny.MultiSet.fromElements(false, !(false)))).IsProperSubsetOf((_module.D6.create_DC21(_dafny.Set.fromElements(_dafny.MultiSet.fromElements(false, !(true))))).dtor_cf42));
    };
    static Main(__noArgsParameter) {
      let _8_globalState;
      let _nw0 = new _module.GlobalState();
      _nw0.__ctor(new BigNumber(470), new BigNumber(-748), true, new BigNumber(220), new BigNumber(551), false);
      _8_globalState = _nw0;
      let _9_v0;
      _9_v0 = new _dafny.CodePoint('e'.codePointAt(0));
      let _10_v1;
      let _nw1 = new _module.C1();
      _nw1.__ctor(_9_v0, false);
      _10_v1 = _nw1;
      (_8_globalState).f1 = new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-275)), ((_11_v1) => function (_12_i0) {
        return _dafny.Seq.of((_11_v1).f7);
      })(_10_v1)))).cardinality());
      let _13_v2;
      _13_v2 = _module.D1.create_DC6(new _dafny.CodePoint('o'.codePointAt(0)), _module.__default.fm6(_8_globalState));
      let _14_v3;
      _14_v3 = _dafny.Seq.of(_module.D1.create_DC6(_9_v0, (_10_v1).f6), ((false) ? (_13_v2) : (_13_v2)));
      _14_v3 = _dafny.Seq.of(_13_v2, _13_v2, _13_v2, _13_v2, _13_v2);
      let _15_v4;
      _15_v4 = new BigNumber(22);
      let _16_v5;
      _16_v5 = _dafny.Map.Empty.slice().updateUnsafe(_15_v4,_15_v4);
      let _17_v6;
      _17_v6 = _dafny.Map.Empty.slice().updateUnsafe(!((_10_v1).f7),true);
      let _18_v7;
      let _init0 = ((_19_v1) => function (_20_i1) {
        return (_19_v1).f7;
      })(_10_v1);
      let _nw2 = Array((new BigNumber(16)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw2.length); _i0_0++) {
        _nw2[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _18_v7 = _nw2;
      let _21_v8;
      _21_v8 = _module.D0.create_DC2(_16_v5, _17_v6, _18_v7);
      let _pat_let_tv0 = _18_v7;
      let _source0 = function (_pat_let0_0) {
        return function (_22_dt__update__tmp_h0) {
          return function (_pat_let1_0) {
            return function (_23_dt__update_hcf8_h0) {
              return _module.D0.create_DC2((_22_dt__update__tmp_h0).dtor_cf6, (_22_dt__update__tmp_h0).dtor_cf7, _23_dt__update_hcf8_h0);
            }(_pat_let1_0);
          }(_pat_let_tv0);
        }(_pat_let0_0);
      }(_21_v8);
      if (_source0.is_DC1) {
        let _24___mcc_h0 = (_source0).cf1;
        let _25___mcc_h1 = (_source0).cf2;
        let _26___mcc_h2 = (_source0).cf3;
        let _27___mcc_h3 = (_source0).cf4;
        let _28___mcc_h4 = (_source0).cf5;
        let _29_cf5 = _28___mcc_h4;
        let _30_cf4 = _27___mcc_h3;
        let _31_cf3 = _26___mcc_h2;
        let _32_cf2 = _25___mcc_h1;
        let _33_cf1 = _24___mcc_h0;
        let _nw3 = Array((new BigNumber(2)).toNumber());
        _nw3[(_dafny.ZERO).toNumber()] = (_10_v1).f7;
        _nw3[(_dafny.ONE).toNumber()] = (_module.__default.fm7(new BigNumber(758), _15_v4, false, _8_globalState)).dtor_cf10;
        _18_v7 = _nw3;
        let _34_v9;
        _34_v9 = _dafny.Map.Empty.slice().updateUnsafe(false,(((_10_v1).f7) ? (_10_v1) : (_10_v1)));
        _34_v9 = (_34_v9).update((_10_v1).f7, _10_v1);
        _16_v5 = (_16_v5).update(_15_v4, _15_v4);
        let _35_v10;
        _35_v10 = _dafny.MultiSet.fromElements(_33_cf1);
        (_8_globalState).f5 = ((_35_v10).IsSubsetOf(_35_v10)) === (_module.__default.fm1(_8_globalState));
      } else if (_source0.is_DC2) {
        let _36___mcc_h5 = (_source0).cf6;
        let _37___mcc_h6 = (_source0).cf7;
        let _38___mcc_h7 = (_source0).cf8;
        let _39_cf8 = _38___mcc_h7;
        let _40_cf7 = _37___mcc_h6;
        let _41_cf6 = _36___mcc_h5;
        let _42_v11;
        _42_v11 = _dafny.Seq.of(_15_v4, _15_v4, _15_v4);
        let _43_v12;
        _43_v12 = _dafny.Map.Empty.slice().updateUnsafe(_42_v11,_15_v4);
        let _44_v13;
        let _45_v14;
        let _46_v15;
        let _out0;
        let _out1;
        let _out2;
        let _outcollector0 = (_10_v1).m0((((_43_v12).contains(_42_v11)) ? ((_43_v12).get(_42_v11)) : (new BigNumber(105))), _15_v4, _8_globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _out2 = _outcollector0[2];
        _44_v13 = _out0;
        _45_v14 = _out1;
        _46_v15 = _out2;
        let _47_v16;
        _47_v16 = _dafny.Seq.UnicodeFromString("olipx");
        _47_v16 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(301)), ((_48_v1) => function (_49_i2) {
          return (_48_v1).f6;
        })(_10_v1)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(434)), ((_50_v1) => function (_51_i3) {
          return (_50_v1).f6;
        })(_10_v1)));
        (_8_globalState).f5 = (_10_v1).f7;
        let _52_v17;
        _52_v17 = _module.D0.create_DC3(_45_v14);
        let _53_v18;
        _53_v18 = _module.D1.create_DC5((_52_v17).dtor_cf9, _41_cf6, new BigNumber((_47_v16).length));
        let _index0 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_46_v15).length));
        (_46_v15)[_index0] = (_53_v18).dtor_cf11;
        let _54_v19;
        _54_v19 = _dafny.Set.fromElements((_10_v1).f7);
        let _55_v20;
        let _init1 = ((_56_v1) => function (_57_i4) {
          return _dafny.Seq.of((_56_v1).f7);
        })(_10_v1);
        let _nw4 = Array((new BigNumber(23)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw4.length); _i0_1++) {
          _nw4[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _55_v20 = _nw4;
        let _58_v21;
        _58_v21 = _module.D0.create_DC1(_15_v4, _54_v19, _45_v14, _55_v20, _9_v0);
        let _59_v22;
        _59_v22 = _dafny.Map.Empty.slice().updateUnsafe((_58_v21).dtor_cf5,!((_10_v1).f7));
        let _60_v23;
        _60_v23 = _dafny.Seq.of((_59_v22).update((_10_v1).f6, (_10_v1).f7));
        let _index1 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_46_v15).length));
        (_46_v15)[_index1] = new BigNumber((_60_v23).length);
      } else if (_source0.is_DC3) {
        let _61___mcc_h8 = (_source0).cf9;
        let _62_cf9 = _61___mcc_h8;
        let _63_v24;
        let _nw5 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
        _63_v24 = _nw5;
        let _64_v25;
        _64_v25 = _dafny.Seq.of(_63_v24);
        _64_v25 = _dafny.Seq.Concat(_64_v25, _dafny.Seq.update(_64_v25, _module.__default.safeIndex(_15_v4, new BigNumber((_64_v25).length)), _63_v24));
        (_8_globalState).f5 = !((_10_v1).f7);
        let _65_v26;
        _65_v26 = _dafny.Set.fromElements((_module.D2.create_DC8(_10_v1)).dtor_cf17);
        if ((_65_v26).IsDisjointFrom((_65_v26).Union(_65_v26))) {
          let _66_v27;
          let _init2 = ((_67_v1, _68_v4) => function (_69_i5) {
            return _dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC9((_67_v1).f6),_68_v4);
          })(_10_v1, _15_v4);
          let _nw6 = Array((new BigNumber(16)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw6.length); _i0_2++) {
            _nw6[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _66_v27 = _nw6;
          let _70_v28;
          _70_v28 = _dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC9((_10_v1).f6),_15_v4);
          let _index2 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_66_v27).length));
          (_66_v27)[_index2] = ((!((_10_v1).f7)) ? (_70_v28) : (_70_v28));
          let _index3 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_66_v27).length));
          (_66_v27)[_index3] = (_70_v28).Merge(_module.__default.fm8(_15_v4, (_10_v1).f7, (_10_v1).f7, _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_10_v1).f7),_62_cf9), _8_globalState));
          let _rhs0 = _10_v1;
          let _rhs1 = (_10_v1).f7;
          let _lhs0 = _8_globalState;
          _10_v1 = _rhs0;
          _lhs0.f5 = _rhs1;
          let _index4 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_18_v7).length));
          (_18_v7)[_index4] = (_10_v1).f7;
          let _71_v29;
          _71_v29 = _dafny.Set.fromElements(!((_10_v1).f7), ((_10_v1).f7) || ((_10_v1).f7), (_10_v1).f7);
          let _72_v30;
          _72_v30 = _dafny.Map.Empty.slice().updateUnsafe(_16_v5,!((_10_v1).f7));
          let _index5 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_18_v7).length));
          let _rhs2 = (_15_v4).isLessThanOrEqualTo(new BigNumber((_72_v30).length));
          let _rhs3 = _71_v29;
          let _lhs1 = _18_v7;
          let _lhs2 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_18_v7).length));
          _lhs1[_lhs2] = _rhs2;
          _71_v29 = _rhs3;
          let _73_v31;
          let _nw7 = new _module.C0();
          _nw7.__ctor(_16_v5, (_18_v7)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_18_v7).length))]);
          _73_v31 = _nw7;
          _73_v31 = _73_v31;
          let _74_v32;
          let _nw8 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _74_v32 = _nw8;
          let _75_v33;
          _75_v33 = _dafny.Seq.UnicodeFromString("bhphawf");
          let _index6 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_74_v32).length));
          (_74_v32)[_index6] = _75_v33;
          let _index7 = _module.__default.safeIndex(new BigNumber(990), new BigNumber((_74_v32).length));
          (_74_v32)[_index7] = _dafny.Seq.Concat(_75_v33, _75_v33);
        } else {
          let _index8 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_18_v7).length));
          (_18_v7)[_index8] = (_10_v1).f7;
          let _index9 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_18_v7).length));
          (_18_v7)[_index9] = (((_10_v1).f7) ? ((_62_cf9).isEqualTo(new BigNumber(-190))) : (_module.__default.fm1(_8_globalState)));
          _10_v1 = _10_v1;
          let _76_v34;
          let _nw9 = new _module.C1();
          _nw9.__ctor((_10_v1).f6, (_10_v1).f7);
          _76_v34 = _nw9;
          let _77_v35;
          _77_v35 = _dafny.Seq.of(_15_v4, _62_cf9);
          (_8_globalState).f5 = !(!((_10_v1).f7) || (_dafny.Seq.IsPrefixOf(_77_v35, _77_v35)));
          let _78_v36;
          _78_v36 = _module.D2.create_DC8(_10_v1);
          let _79_v37;
          let _nw10 = Array((new BigNumber(14)).toNumber());
          _nw10[(_dafny.ZERO).toNumber()] = (_78_v36).dtor_cf17;
          _nw10[(_dafny.ONE).toNumber()] = _10_v1;
          _nw10[(new BigNumber(2)).toNumber()] = _10_v1;
          _nw10[(new BigNumber(3)).toNumber()] = _76_v34;
          _nw10[(new BigNumber(4)).toNumber()] = _76_v34;
          _nw10[(new BigNumber(5)).toNumber()] = _76_v34;
          _nw10[(new BigNumber(6)).toNumber()] = _10_v1;
          _nw10[(new BigNumber(7)).toNumber()] = _76_v34;
          _nw10[(new BigNumber(8)).toNumber()] = _76_v34;
          _nw10[(new BigNumber(9)).toNumber()] = _10_v1;
          _nw10[(new BigNumber(10)).toNumber()] = (((_10_v1).f7) ? (_10_v1) : (_10_v1));
          _nw10[(new BigNumber(11)).toNumber()] = _10_v1;
          _nw10[(new BigNumber(12)).toNumber()] = _76_v34;
          _nw10[(new BigNumber(13)).toNumber()] = _76_v34;
          _79_v37 = _nw10;
          let _index10 = _module.__default.safeIndex(new BigNumber(895), new BigNumber((_79_v37).length));
          (_79_v37)[_index10] = _10_v1;
          let _index11 = _module.__default.safeIndex(new BigNumber(895), new BigNumber((_79_v37).length));
          (_79_v37)[_index11] = _10_v1;
        }
        let _80_v38;
        _80_v38 = _dafny.Seq.UnicodeFromString("qsq");
        _15_v4 = (_15_v4).minus(new BigNumber((_80_v38).length));
      } else {
        let _81___mcc_h9 = (_source0).cf0;
        let _82_cf0 = _81___mcc_h9;
        let _83_v39;
        let _init3 = ((_84_v4) => function (_85_i6) {
          return (_85_i6).plus(_84_v4);
        })(_15_v4);
        let _nw11 = Array((new BigNumber(24)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw11.length); _i0_3++) {
          _nw11[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _83_v39 = _nw11;
        let _index12 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_83_v39).length));
        (_83_v39)[_index12] = _15_v4;
        let _index13 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_83_v39).length));
        let _rhs4 = _15_v4;
        let _rhs5 = new BigNumber(280);
        let _lhs3 = _83_v39;
        let _lhs4 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_83_v39).length));
        let _lhs5 = _8_globalState;
        _lhs3[_lhs4] = _rhs4;
        _lhs5.f1 = _rhs5;
        let _86_v40;
        _86_v40 = _module.D0.create_DC0(_18_v7);
        let _87_v41;
        _87_v41 = _dafny.Map.Empty.slice().updateUnsafe(_86_v40,(_83_v39)[_module.__default.safeIndex(new BigNumber(624), new BigNumber((_83_v39).length))]);
        let _88_v42;
        _88_v42 = _dafny.Seq.UnicodeFromString("htil");
        let _89_v43;
        _89_v43 = _module.D1.create_DC5(new BigNumber((_88_v42).length), _16_v5, (_83_v39)[_module.__default.safeIndex(new BigNumber(624), new BigNumber((_83_v39).length))]);
        let _index14 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_83_v39).length));
        (_83_v39)[_index14] = (((_10_v1).f7) ? (_module.__default.safeModuloInt((((_87_v41).contains(_86_v40)) ? ((_87_v41).get(_86_v40)) : ((_83_v39)[_module.__default.safeIndex(new BigNumber(624), new BigNumber((_83_v39).length))])), new BigNumber(623))) : ((_dafny.ZERO).minus((_89_v43).dtor_cf13)));
        _18_v7 = _18_v7;
        _17_v6 = (_17_v6).update((_10_v1).f7, (_10_v1).f7);
      }
      let _hi0 = _15_v4;
      for (let _90_i7 = _15_v4; _90_i7.isLessThan(_hi0); _90_i7 = _90_i7.plus(_dafny.ONE)) {
        (_8_globalState).f1 = (_90_i7).minus(_module.__default.fm2(false, _8_globalState));
        let _91_v44;
        let _nw12 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
        _91_v44 = _nw12;
        let _92_v45;
        _92_v45 = _dafny.Seq.UnicodeFromString("a");
        let _index15 = _module.__default.safeIndex(new BigNumber(413), new BigNumber((_91_v44).length));
        (_91_v44)[_index15] = (_15_v4).plus((_dafny.ZERO).minus(new BigNumber((_92_v45).length)));
        let _index16 = _module.__default.safeIndex(new BigNumber(413), new BigNumber((_91_v44).length));
        (_91_v44)[_index16] = _15_v4;
        let _93_v46;
        _93_v46 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("ekadg"),_15_v4);
        _93_v46 = (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("tpsq"),(_91_v44)[_module.__default.safeIndex(new BigNumber(413), new BigNumber((_91_v44).length))])).Merge(_93_v46);
        let _index17 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_18_v7).length));
        (_18_v7)[_index17] = !(false);
        let _index18 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_18_v7).length));
        let _rhs6 = (_10_v1).f7;
        let _rhs7 = (_10_v1).f7;
        let _rhs8 = (_10_v1).f6;
        let _lhs6 = _18_v7;
        let _lhs7 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_18_v7).length));
        let _lhs8 = _8_globalState;
        _lhs6[_lhs7] = _rhs6;
        _lhs8.f5 = _rhs7;
        _9_v0 = _rhs8;
      }
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_18_v7).length))) {
        let _94_i8 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_94_i8)) && ((_94_i8).isLessThan(new BigNumber((_18_v7).length))))) {
          (_18_v7)[(_94_i8)] = (_10_v1).f7;
        }
      }
      let _95_v47;
      _95_v47 = _dafny.Seq.UnicodeFromString("fyyyndfmr");
      _95_v47 = (((_10_v1).f7) ? (_95_v47) : (_dafny.Seq.UnicodeFromString("vintro")));
      let _hi1 = new BigNumber(320);
      for (let _96_i9 = _15_v4; _96_i9.isLessThan(_hi1); _96_i9 = _96_i9.plus(_dafny.ONE)) {
        let _97_v48;
        let _98_v49;
        let _99_v50;
        let _out3;
        let _out4;
        let _out5;
        let _outcollector1 = (_10_v1).m0(_15_v4, _15_v4, _8_globalState);
        _out3 = _outcollector1[0];
        _out4 = _outcollector1[1];
        _out5 = _outcollector1[2];
        _97_v48 = _out3;
        _98_v49 = _out4;
        _99_v50 = _out5;
        let _100_v51;
        _100_v51 = _module.D2.create_DC8(_10_v1);
        let _101_v52;
        _101_v52 = _module.D2.create_DC12(_100_v51);
        let _source1 = _101_v52;
        if (_source1.is_DC9) {
          let _102___mcc_h10 = (_source1).cf18;
          let _103_cf18 = _102___mcc_h10;
          let _index19 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_99_v50).length));
          (_99_v50)[_index19] = (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_10_v1).f7,_96_i9)).length));
          let _index20 = _module.__default.safeIndex(new BigNumber(687), new BigNumber((_99_v50).length));
          (_99_v50)[_index20] = _96_i9;
          let _104_v53;
          let _nw13 = Array((new BigNumber(4)).toNumber()).fill([]);
          _104_v53 = _nw13;
          let _index21 = _module.__default.safeIndex(new BigNumber(613), new BigNumber((_104_v53).length));
          (_104_v53)[_index21] = _18_v7;
          let _index22 = _module.__default.safeIndex(new BigNumber(613), new BigNumber((_104_v53).length));
          (_104_v53)[_index22] = _18_v7;
          let _105_v54;
          _105_v54 = _module.D1.create_DC7((_10_v1).f6);
          let _106_v55;
          _106_v55 = _dafny.Map.Empty.slice().updateUnsafe(_105_v54,(_dafny.ZERO).minus(_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("pb")).length)), (_dafny.ZERO).minus(_98_v49))));
          _106_v55 = _106_v55;
          _16_v5 = (_16_v5).update((_15_v4).plus(_96_i9), new BigNumber(321));
        } else if (_source1.is_DC10) {
          let _107___mcc_h11 = (_source1).cf19;
          let _108___mcc_h12 = (_source1).cf20;
          let _109_cf20 = _108___mcc_h12;
          let _110_cf19 = _107___mcc_h11;
          let _111_v56;
          _111_v56 = _dafny.Map.Empty.slice().updateUnsafe(true,_97_v48);
          let _index23 = _module.__default.safeIndex(new BigNumber(536), new BigNumber((_99_v50).length));
          (_99_v50)[_index23] = new BigNumber(((_111_v56).update((_10_v1).f7, _97_v48)).length);
          let _index24 = _module.__default.safeIndex(new BigNumber(536), new BigNumber((_99_v50).length));
          (_99_v50)[_index24] = _module.__default.safeDivisionInt(new BigNumber((_95_v47).length), _98_v49);
          let _index25 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_18_v7).length));
          (_18_v7)[_index25] = _110_cf19;
          let _index26 = _module.__default.safeIndex(new BigNumber(6), new BigNumber((_18_v7).length));
          (_18_v7)[_index26] = !(true) || (!_dafny.Seq.contains(_95_v47, _9_v0));
          let _112_v57;
          _112_v57 = _dafny.Map.Empty.slice().updateUnsafe(!(_module.__default.fm2((_10_v1).f7, _8_globalState)).isEqualTo((_99_v50)[_module.__default.safeIndex(new BigNumber(536), new BigNumber((_99_v50).length))]),_dafny.Set.fromElements((_18_v7)[_module.__default.safeIndex(new BigNumber(6), new BigNumber((_18_v7).length))]));
          let _113_v58;
          _113_v58 = _module.D1.create_DC4(_110_cf19);
          let _114_v59;
          _114_v59 = _dafny.Map.Empty.slice().updateUnsafe(_113_v58,_dafny.Seq.UnicodeFromString("g"));
          let _115_v60;
          _115_v60 = _dafny.Seq.of(_98_v49);
          let _116_v61;
          _116_v61 = _dafny.Set.fromElements(_109_cf20, true, _module.__default.fm1(_8_globalState), _110_cf19, _110_cf19);
          let _rhs9 = !_dafny.areEqual((((_114_v59).contains(_module.D1.create_DC4(_110_cf19))) ? ((_114_v59).get(_module.D1.create_DC4(_110_cf19))) : (_95_v47)), _95_v47);
          let _rhs10 = (_dafny.ZERO).minus(_98_v49);
          let _rhs11 = (_112_v57).update(_dafny.Seq.contains(_115_v60, (_99_v50)[_module.__default.safeIndex(new BigNumber(536), new BigNumber((_99_v50).length))]), (_dafny.Set.fromElements(_110_cf19, _109_cf20, _110_cf19)).Intersect(_116_v61));
          let _rhs12 = _10_v1;
          let _lhs9 = _8_globalState;
          _110_cf19 = _rhs9;
          _lhs9.f1 = _rhs10;
          _112_v57 = _rhs11;
          _10_v1 = _rhs12;
          let _117_v62;
          _117_v62 = _dafny.Map.Empty.slice().updateUnsafe(_98_v49,_dafny.MultiSet.fromElements(_10_v1, _10_v1));
          let _118_v63;
          _118_v63 = _module.D2.create_DC8(_10_v1);
          let _119_v64;
          _119_v64 = _dafny.MultiSet.fromElements(_10_v1, _10_v1, _10_v1, (_118_v63).dtor_cf17, _10_v1);
          let _120_v65;
          let _121_v66;
          let _122_v67;
          let _out6;
          let _out7;
          let _out8;
          let _outcollector2 = (_10_v1).m0(new BigNumber((_115_v60).length), new BigNumber(((((_117_v62).contains(_98_v49)) ? ((_117_v62).get(_98_v49)) : (_119_v64))).cardinality()), _8_globalState);
          _out6 = _outcollector2[0];
          _out7 = _outcollector2[1];
          _out8 = _outcollector2[2];
          _120_v65 = _out6;
          _121_v66 = _out7;
          _122_v67 = _out8;
        } else if (_source1.is_DC11) {
          let _123___mcc_h13 = (_source1).cf21;
          let _124___mcc_h14 = (_source1).cf22;
          let _125_cf22 = _124___mcc_h14;
          let _126_cf21 = _123___mcc_h13;
          let _rhs13 = _15_v4;
          let _rhs14 = _99_v50;
          _126_cf21 = _rhs13;
          _99_v50 = _rhs14;
          (_8_globalState).f5 = true;
          (_8_globalState).f5 = (_10_v1).f7;
          let _127_v68;
          _127_v68 = _dafny.Map.Empty.slice().updateUnsafe(_98_v49,true);
          _127_v68 = (_127_v68).update(new BigNumber(-12), false);
        } else if (_source1.is_DC8) {
          let _128___mcc_h15 = (_source1).cf17;
          let _129_cf17 = _128___mcc_h15;
          _18_v7 = _18_v7;
          (_8_globalState).f1 = _97_v48;
          let _130_v69;
          _130_v69 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_95_v47).length),!((_10_v1).f7));
          let _131_v70;
          let _132_v71;
          let _133_v72;
          let _out9;
          let _out10;
          let _out11;
          let _outcollector3 = (_10_v1).m0(_98_v49, _module.__default.safeModuloInt(new BigNumber((_130_v69).length), new BigNumber(440)), _8_globalState);
          _out9 = _outcollector3[0];
          _out10 = _outcollector3[1];
          _out11 = _outcollector3[2];
          _131_v70 = _out9;
          _132_v71 = _out10;
          _133_v72 = _out11;
          (_8_globalState).f5 = (_10_v1).f7;
        } else {
          let _134___mcc_h16 = (_source1).cf23;
          let _135_cf23 = _134___mcc_h16;
          (_8_globalState).f5 = !(_module.__default.fm1(_8_globalState));
          let _136_v73;
          let _nw14 = new _module.C0();
          _nw14.__ctor(_16_v5, (_10_v1).f7);
          _136_v73 = _nw14;
          let _137_v74;
          _137_v74 = _dafny.MultiSet.fromElements(_136_v73, _136_v73, _136_v73, _136_v73);
          let _138_v75;
          _138_v75 = _dafny.Seq.of(new BigNumber((_137_v74).cardinality()));
          let _139_v76;
          _139_v76 = _dafny.Seq.of(_138_v75);
          let _140_v77;
          _140_v77 = _dafny.Set.fromElements(_97_v48, new BigNumber((_139_v76).length));
          let _rhs15 = (_10_v1).f7;
          let _rhs16 = _140_v77;
          let _lhs10 = _8_globalState;
          _lhs10.f5 = _rhs15;
          _140_v77 = _rhs16;
          _9_v0 = _9_v0;
          let _141_v78;
          _141_v78 = _dafny.MultiSet.fromElements(_18_v7, _18_v7);
          let _142_v79;
          _142_v79 = _dafny.Map.Empty.slice().updateUnsafe(_98_v49,(_dafny.MultiSet.fromElements(_18_v7)).update(_18_v7, _module.__default.abs(_98_v49)));
          let _143_v80;
          let _nw15 = Array((new BigNumber(14)).toNumber());
          _nw15[(_dafny.ZERO).toNumber()] = _141_v78;
          _nw15[(_dafny.ONE).toNumber()] = _141_v78;
          _nw15[(new BigNumber(2)).toNumber()] = _141_v78;
          _nw15[(new BigNumber(3)).toNumber()] = _141_v78;
          _nw15[(new BigNumber(4)).toNumber()] = (_141_v78).update(_18_v7, _module.__default.abs(_97_v48));
          _nw15[(new BigNumber(5)).toNumber()] = _141_v78;
          _nw15[(new BigNumber(6)).toNumber()] = (_141_v78).Difference(_141_v78);
          _nw15[(new BigNumber(7)).toNumber()] = (_141_v78).Union(_141_v78);
          _nw15[(new BigNumber(8)).toNumber()] = _141_v78;
          _nw15[(new BigNumber(9)).toNumber()] = _141_v78;
          _nw15[(new BigNumber(10)).toNumber()] = _141_v78;
          _nw15[(new BigNumber(11)).toNumber()] = (((_142_v79).contains(_97_v48)) ? ((_142_v79).get(_97_v48)) : (_141_v78));
          _nw15[(new BigNumber(12)).toNumber()] = (_141_v78).Difference(_141_v78);
          _nw15[(new BigNumber(13)).toNumber()] = _dafny.MultiSet.fromElements(_18_v7);
          _143_v80 = _nw15;
          let _index27 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_143_v80).length));
          (_143_v80)[_index27] = ((_dafny.MultiSet.fromElements(_18_v7)).update(_18_v7, _module.__default.abs(_96_i9))).update(_18_v7, _module.__default.abs(new BigNumber((_95_v47).length)));
          let _144_v81;
          _144_v81 = _dafny.Seq.of(_18_v7);
          let _145_v82;
          _145_v82 = _dafny.Seq.of(_141_v78, _141_v78, _141_v78, _141_v78, _dafny.MultiSet.fromElements(_18_v7, _18_v7, (_144_v81)[_module.__default.safeIndex(_15_v4, new BigNumber((_144_v81).length))], _18_v7, _18_v7));
          let _146_v83;
          _146_v83 = _dafny.Map.Empty.slice().updateUnsafe((_10_v1).f7,_15_v4);
          let _147_v84;
          _147_v84 = _dafny.Map.Empty.slice().updateUnsafe(_140_v77,_146_v83);
          let _index28 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_143_v80).length));
          (_143_v80)[_index28] = ((_145_v82)[_module.__default.safeIndex(new BigNumber(((((_147_v84).contains(_140_v77)) ? ((_147_v84).get(_140_v77)) : (_146_v83))).length), new BigNumber((_145_v82).length))]).Difference(_141_v78);
        }
        let _index29 = _module.__default.safeIndex(new BigNumber(859), new BigNumber((_18_v7).length));
        (_18_v7)[_index29] = (_10_v1).f7;
        let _index30 = _module.__default.safeIndex(new BigNumber(859), new BigNumber((_18_v7).length));
        (_18_v7)[_index30] = true;
        let _148_v85;
        let _nw16 = new _module.C1();
        _nw16.__ctor(new _dafny.CodePoint('b'.codePointAt(0)), false);
        _148_v85 = _nw16;
      }
      let _149_i10;
      _149_i10 = _dafny.ZERO;
      L0: {
        while ((_10_v1).f7) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_149_i10)) {
              break L0;
            }
            _149_i10 = (_149_i10).plus(_dafny.ONE);
            if ((_10_v1).f7) {
              (_8_globalState).f1 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_15_v4, _15_v4));
              let _150_v86;
              _150_v86 = _module.D2.create_DC9((_10_v1).f6);
              _150_v86 = _150_v86;
              let _151_v87;
              let _152_v88;
              let _153_v89;
              let _out12;
              let _out13;
              let _out14;
              let _outcollector4 = (_10_v1).m0(_15_v4, _15_v4, _8_globalState);
              _out12 = _outcollector4[0];
              _out13 = _outcollector4[1];
              _out14 = _outcollector4[2];
              _151_v87 = _out12;
              _152_v88 = _out13;
              _153_v89 = _out14;
              let _154_v90;
              _154_v90 = _dafny.Map.Empty.slice().updateUnsafe(_18_v7,_15_v4);
              _154_v90 = (_154_v90).update(_18_v7, _15_v4);
              _95_v47 = (_10_v1).fm0(_dafny.Seq.UnicodeFromString("arwyvlusb"), (_10_v1).f6, _8_globalState);
            } else {
              _9_v0 = (_95_v47)[_module.__default.safeIndex(_15_v4, new BigNumber((_95_v47).length))];
              let _index31 = _module.__default.safeIndex(new BigNumber(41), new BigNumber((_18_v7).length));
              (_18_v7)[_index31] = (_10_v1).f7;
              let _index32 = _module.__default.safeIndex(new BigNumber(41), new BigNumber((_18_v7).length));
              let _rhs17 = _dafny.Seq.Concat(_dafny.Seq.Concat(_95_v47, _95_v47), _dafny.Seq.UnicodeFromString("hgpygk"));
              let _rhs18 = (_10_v1).f7;
              let _rhs19 = (_10_v1).f7;
              let _lhs11 = _8_globalState;
              let _lhs12 = _18_v7;
              let _lhs13 = _module.__default.safeIndex(new BigNumber(41), new BigNumber((_18_v7).length));
              _95_v47 = _rhs17;
              _lhs11.f5 = _rhs18;
              _lhs12[_lhs13] = _rhs19;
              (_8_globalState).f5 = !((_15_v4).isLessThan(_15_v4));
              let _155_v91;
              let _nw17 = Array((new BigNumber(2)).toNumber());
              _nw17[(_dafny.ZERO).toNumber()] = _9_v0;
              _nw17[(_dafny.ONE).toNumber()] = _9_v0;
              _155_v91 = _nw17;
              let _156_v92;
              _156_v92 = _dafny.Seq.of(_155_v91);
              let _157_v93;
              _157_v93 = _module.D3.create_DC13(_155_v91);
              let _158_v94;
              let _nw18 = Array((new BigNumber(26)).toNumber());
              _nw18[(_dafny.ZERO).toNumber()] = _155_v91;
              _nw18[(_dafny.ONE).toNumber()] = _155_v91;
              _nw18[(new BigNumber(2)).toNumber()] = _155_v91;
              _nw18[(new BigNumber(3)).toNumber()] = _155_v91;
              _nw18[(new BigNumber(4)).toNumber()] = _155_v91;
              _nw18[(new BigNumber(5)).toNumber()] = _155_v91;
              _nw18[(new BigNumber(6)).toNumber()] = _155_v91;
              _nw18[(new BigNumber(7)).toNumber()] = _155_v91;
              _nw18[(new BigNumber(8)).toNumber()] = _155_v91;
              _nw18[(new BigNumber(9)).toNumber()] = (((_10_v1).f7) ? (_155_v91) : (_155_v91));
              _nw18[(new BigNumber(10)).toNumber()] = _155_v91;
              _nw18[(new BigNumber(11)).toNumber()] = (_156_v92)[_module.__default.safeIndex(_15_v4, new BigNumber((_156_v92).length))];
              _nw18[(new BigNumber(12)).toNumber()] = _155_v91;
              _nw18[(new BigNumber(13)).toNumber()] = _155_v91;
              _nw18[(new BigNumber(14)).toNumber()] = _155_v91;
              _nw18[(new BigNumber(15)).toNumber()] = _155_v91;
              _nw18[(new BigNumber(16)).toNumber()] = _155_v91;
              _nw18[(new BigNumber(17)).toNumber()] = _155_v91;
              _nw18[(new BigNumber(18)).toNumber()] = _155_v91;
              _nw18[(new BigNumber(19)).toNumber()] = (_157_v93).dtor_cf24;
              _nw18[(new BigNumber(20)).toNumber()] = _155_v91;
              _nw18[(new BigNumber(21)).toNumber()] = _155_v91;
              _nw18[(new BigNumber(22)).toNumber()] = _155_v91;
              _nw18[(new BigNumber(23)).toNumber()] = _155_v91;
              _nw18[(new BigNumber(24)).toNumber()] = _155_v91;
              _nw18[(new BigNumber(25)).toNumber()] = _155_v91;
              _158_v94 = _nw18;
              _158_v94 = _158_v94;
              let _159_v95;
              _159_v95 = _dafny.MultiSet.fromElements(_15_v4);
              let _160_v96;
              let _nw19 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Seq.of());
              _160_v96 = _nw19;
              let _161_v97;
              let _init4 = ((_162_v1) => function (_163_i11) {
                return _module.D2.create_DC9((_162_v1).f6);
              })(_10_v1);
              let _nw20 = Array((new BigNumber(3)).toNumber());
              for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw20.length); _i0_4++) {
                _nw20[_i0_4] = _init4(new BigNumber(_i0_4));
              }
              _161_v97 = _nw20;
              let _164_v98;
              _164_v98 = _dafny.Seq.of(_161_v97, _161_v97);
              let _index33 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_160_v96).length));
              (_160_v96)[_index33] = _164_v98;
              let _165_v99;
              _165_v99 = _dafny.Set.fromElements((_18_v7)[_module.__default.safeIndex(new BigNumber(41), new BigNumber((_18_v7).length))]);
              let _index34 = _module.__default.safeIndex(new BigNumber(41), new BigNumber((_18_v7).length));
              let _index35 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_160_v96).length));
              let _rhs20 = (_18_v7)[_module.__default.safeIndex(new BigNumber(41), new BigNumber((_18_v7).length))];
              let _rhs21 = _159_v95;
              let _rhs22 = _module.__default.safeModuloInt(_15_v4, _15_v4);
              let _rhs23 = _dafny.Seq.of(_161_v97);
              let _rhs24 = (_dafny.Set.fromElements((_10_v1).f7)).IsProperSubsetOf((((_10_v1).f7) ? (_165_v99) : (_165_v99)));
              let _lhs14 = _18_v7;
              let _lhs15 = _module.__default.safeIndex(new BigNumber(41), new BigNumber((_18_v7).length));
              let _lhs16 = _8_globalState;
              let _lhs17 = _160_v96;
              let _lhs18 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_160_v96).length));
              let _lhs19 = _8_globalState;
              _lhs14[_lhs15] = _rhs20;
              _159_v95 = _rhs21;
              _lhs16.f1 = _rhs22;
              _lhs17[_lhs18] = _rhs23;
              _lhs19.f5 = _rhs24;
            }
            if (_dafny.areEqual(_95_v47, _95_v47)) {
              let _166_v100;
              _166_v100 = _dafny.Map.Empty.slice().updateUnsafe(_15_v4,(_10_v1).f7);
              _166_v100 = (_166_v100).update(_module.__default.safeDivisionInt(_15_v4, new BigNumber((_dafny.MultiSet.fromElements(_95_v47)).cardinality())), (_10_v1).f7);
              let _index36 = _module.__default.safeIndex(new BigNumber(823), new BigNumber((_18_v7).length));
              (_18_v7)[_index36] = (_10_v1).f7;
              let _167_v101;
              let _nw21 = Array((new BigNumber(7)).toNumber());
              _167_v101 = _nw21;
              let _168_v102;
              _168_v102 = _dafny.Map.Empty.slice().updateUnsafe(_167_v101,(_10_v1).f7);
              let _169_v103;
              _169_v103 = _dafny.MultiSet.fromElements(_15_v4, _15_v4, _15_v4);
              let _170_v104;
              _170_v104 = _dafny.Set.fromElements(!((_10_v1).f7), (_10_v1).f7);
              let _index37 = _module.__default.safeIndex(new BigNumber(823), new BigNumber((_18_v7).length));
              let _rhs25 = (_10_v1).f7;
              let _rhs26 = (_dafny.ZERO).minus(new BigNumber(-735));
              let _rhs27 = new BigNumber(((_168_v102).update(_167_v101, (_dafny.MultiSet.fromElements(_15_v4, _15_v4, new BigNumber((_170_v104).length), _15_v4, _15_v4)).IsProperSubsetOf(_169_v103))).length);
              let _lhs20 = _18_v7;
              let _lhs21 = _module.__default.safeIndex(new BigNumber(823), new BigNumber((_18_v7).length));
              let _lhs22 = _8_globalState;
              let _lhs23 = _8_globalState;
              _lhs20[_lhs21] = _rhs25;
              _lhs22.f1 = _rhs26;
              _lhs23.f0 = _rhs27;
              let _171_v105;
              let _nw22 = Array((new BigNumber(7)).toNumber()).fill([]);
              _171_v105 = _nw22;
              let _172_v106;
              _172_v106 = _dafny.Map.Empty.slice().updateUnsafe(_171_v105,(_10_v1).f7);
              _172_v106 = (_172_v106).update(_171_v105, (_10_v1).f7);
              let _173_v107;
              _173_v107 = _dafny.Map.Empty.slice().updateUnsafe(_10_v1,new BigNumber(806));
              let _174_v108;
              _174_v108 = _dafny.Seq.of((_10_v1).f7, _module.__default.fm1(_8_globalState), (_10_v1).f7, !((_10_v1).f7), (_18_v7)[_module.__default.safeIndex(new BigNumber(823), new BigNumber((_18_v7).length))]);
              let _175_v109;
              let _nw23 = new _module.C0();
              _nw23.__ctor((_16_v5).update((((_173_v107).contains(_10_v1)) ? ((_173_v107).get(_10_v1)) : (new BigNumber((_174_v108).length))), _15_v4), (_18_v7)[_module.__default.safeIndex(new BigNumber(823), new BigNumber((_18_v7).length))]);
              _175_v109 = _nw23;
              _15_v4 = _15_v4;
            } else {
              let _index38 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_18_v7).length));
              (_18_v7)[_index38] = _module.__default.fm1(_8_globalState);
              let _index39 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_18_v7).length));
              (_18_v7)[_index39] = (_10_v1).f7;
              let _176_v110;
              _176_v110 = _dafny.Map.Empty.slice().updateUnsafe(_15_v4,_95_v47);
              let _177_v111;
              _177_v111 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-45),_176_v110);
              _177_v111 = (_177_v111).update(_15_v4, _176_v110);
              let _178_v112;
              _178_v112 = _module.D1.create_DC4((_10_v1).f7);
              let _pat_let_tv1 = _18_v7;
              let _pat_let_tv2 = _18_v7;
              _178_v112 = function (_pat_let2_0) {
                return function (_179_dt__update__tmp_h1) {
                  return function (_pat_let3_0) {
                    return function (_180_dt__update_hcf10_h0) {
                      return _module.D1.create_DC4(_180_dt__update_hcf10_h0);
                    }(_pat_let3_0);
                  }((_pat_let_tv2)[_module.__default.safeIndex(new BigNumber(315), new BigNumber((_pat_let_tv1).length))]);
                }(_pat_let2_0);
              }(_178_v112);
              (_8_globalState).f0 = (_dafny.ZERO).minus(_15_v4);
              let _181_v113;
              let _nw24 = new _module.C1();
              _nw24.__ctor(new _dafny.CodePoint('j'.codePointAt(0)), (_18_v7)[_module.__default.safeIndex(new BigNumber(315), new BigNumber((_18_v7).length))]);
              _181_v113 = _nw24;
            }
            let _182_v114;
            _182_v114 = _dafny.Seq.of(_module.__default.fm2((_10_v1).f7, _8_globalState), new BigNumber(555));
            let _183_v116;
            _183_v116 = _dafny.Set.fromElements(_15_v4);
            (_8_globalState).f5 = (_183_v116).IsProperSubsetOf(function () {
              let _coll4 = new _dafny.Set();
              for (const _compr_4 of (_182_v114).Elements) {
                let _184_v115 = _compr_4;
                if (_dafny.Seq.contains(_182_v114, _184_v115)) {
                  _coll4.add((_184_v115).minus(new BigNumber(-668)));
                }
              }
              return _coll4;
            }());
            let _185_v117;
            let _nw25 = Array((new BigNumber(2)).toNumber());
            _nw25[(_dafny.ZERO).toNumber()] = _95_v47;
            _nw25[(_dafny.ONE).toNumber()] = ((!((_10_v1).f7)) ? (_dafny.Seq.UnicodeFromString("tdaboijva")) : (_95_v47));
            _185_v117 = _nw25;
            let _index40 = _module.__default.safeIndex(new BigNumber(996), new BigNumber((_185_v117).length));
            (_185_v117)[_index40] = _95_v47;
            let _index41 = _module.__default.safeIndex(new BigNumber(996), new BigNumber((_185_v117).length));
            (_185_v117)[_index41] = (_10_v1).fm0((_10_v1).fm0(_dafny.Seq.UnicodeFromString("gqains"), _9_v0, _8_globalState), (_10_v1).f6, _8_globalState);
          }
        }
      }
      let _186_v118;
      _186_v118 = _dafny.MultiSet.fromElements(_17_v6);
      let _187_v119;
      let _nw26 = Array((new BigNumber(11)).toNumber());
      _nw26[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-948)), ((_188_v1) => function (_189_i12) {
        return (_188_v1).f6;
      })(_10_v1))).length);
      _nw26[(_dafny.ONE).toNumber()] = _15_v4;
      _nw26[(new BigNumber(2)).toNumber()] = _15_v4;
      _nw26[(new BigNumber(3)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(621), new BigNumber(665));
      _nw26[(new BigNumber(4)).toNumber()] = new BigNumber((_17_v6).length);
      _nw26[(new BigNumber(5)).toNumber()] = _module.__default.fm2((((_17_v6).contains(false)) ? ((_17_v6).get(false)) : ((_10_v1).f7)), _8_globalState);
      _nw26[(new BigNumber(6)).toNumber()] = _15_v4;
      _nw26[(new BigNumber(7)).toNumber()] = _15_v4;
      _nw26[(new BigNumber(8)).toNumber()] = _15_v4;
      _nw26[(new BigNumber(9)).toNumber()] = _15_v4;
      _nw26[(new BigNumber(10)).toNumber()] = new BigNumber(((_module.__default.fm9((_10_v1).f7, (_10_v1).f7, _8_globalState)).Difference(_186_v118)).cardinality());
      _187_v119 = _nw26;
      let _index42 = _module.__default.safeIndex(new BigNumber(229), new BigNumber((_187_v119).length));
      (_187_v119)[_index42] = _15_v4;
      let _190_v120;
      _190_v120 = _dafny.MultiSet.fromElements(_15_v4, _15_v4);
      let _191_v121;
      _191_v121 = _dafny.Seq.of(new BigNumber((_95_v47).length));
      let _index43 = _module.__default.safeIndex(new BigNumber(229), new BigNumber((_187_v119).length));
      let _rhs28 = _15_v4;
      let _rhs29 = (((_190_v120).contains(new BigNumber(-694))) ? ((_190_v120).get(new BigNumber(-694))) : ((_191_v121)[_module.__default.safeIndex(_module.__default.fm2((_10_v1).f7, _8_globalState), new BigNumber((_191_v121).length))]));
      let _rhs30 = _15_v4;
      let _rhs31 = !(!((_10_v1).f7) || (!((_10_v1).f7)));
      let _lhs24 = _187_v119;
      let _lhs25 = _module.__default.safeIndex(new BigNumber(229), new BigNumber((_187_v119).length));
      let _lhs26 = _8_globalState;
      _lhs24[_lhs25] = _rhs28;
      _15_v4 = _rhs29;
      _15_v4 = _rhs30;
      _lhs26.f5 = _rhs31;
      let _index44 = _module.__default.safeIndex(new BigNumber(229), new BigNumber((_187_v119).length));
      (_187_v119)[_index44] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(872)), ((_192_v0) => function (_193_i13) {
        return _192_v0;
      })(_9_v0))).length);
      let _194_v122;
      _194_v122 = _dafny.Seq.of(!((_10_v1).f7) || ((_10_v1).f7));
      (_8_globalState).f5 = (_194_v122)[_module.__default.safeIndex(_15_v4, new BigNumber((_194_v122).length))];
      let _195_v123;
      _195_v123 = _dafny.Map.Empty.slice().updateUnsafe((_10_v1).f7,(_187_v119)[_module.__default.safeIndex(new BigNumber(229), new BigNumber((_187_v119).length))]);
      let _196_v124;
      _196_v124 = _module.D0.create_DC3(_15_v4);
      let _197_v125;
      _197_v125 = _dafny.Map.Empty.slice().updateUnsafe((_194_v122)[_module.__default.safeIndex(new BigNumber(((_190_v120).update((((_195_v123).contains(!(!((_10_v1).f7)))) ? ((_195_v123).get(!(!((_10_v1).f7)))) : (_15_v4)), _module.__default.abs((_187_v119)[_module.__default.safeIndex(new BigNumber(229), new BigNumber((_187_v119).length))]))).cardinality()), new BigNumber((_194_v122).length))],_196_v124);
      _197_v125 = (_197_v125).update((_10_v1).f7, _196_v124);
      if ((new BigNumber((_95_v47).length)).isLessThanOrEqualTo(_15_v4)) {
        let _198_v126;
        _198_v126 = _dafny.Set.fromElements(new BigNumber(132), _15_v4);
        _198_v126 = (_198_v126).Intersect((function () {
          let _coll5 = new _dafny.Set();
          for (const _compr_5 of _dafny.IntegerRange(new BigNumber(928), new BigNumber(795))) {
            let _199_v127 = _compr_5;
            if (((new BigNumber(928)).isLessThanOrEqualTo(_199_v127)) && ((_199_v127).isLessThan(new BigNumber(795)))) {
              _coll5.add((_199_v127).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(347)), ((_200_v119) => function (_201_i14) {
                return (_200_v119)[_module.__default.safeIndex(new BigNumber(229), new BigNumber((_200_v119).length))];
              })(_187_v119))).length)));
            }
          }
          return _coll5;
        }()).Difference(_198_v126));
        (_8_globalState).f5 = !((_10_v1).f7);
        let _202_v128;
        _202_v128 = _dafny.Seq.of(_16_v5);
        let _203_v130;
        _203_v130 = _dafny.Map.Empty.slice().updateUnsafe((_187_v119)[_module.__default.safeIndex(new BigNumber(229), new BigNumber((_187_v119).length))],(_10_v1).f7);
        let _204_v131;
        _204_v131 = _dafny.Seq.of(_203_v130);
        let _205_v132;
        let _nw27 = new _module.C0();
        _nw27.__ctor(((_202_v128)[_module.__default.safeIndex(new BigNumber((_191_v121).length), new BigNumber((_202_v128).length))]).update(_15_v4, new BigNumber((function () {
          let _coll6 = new _dafny.Map();
          for (const _compr_6 of (_204_v131).Elements) {
            let _206_v129 = _compr_6;
            if (_dafny.Seq.contains(_204_v131, _206_v129)) {
              _coll6.push([_206_v129,(((_17_v6).contains(true)) ? ((_17_v6).get(true)) : ((_10_v1).f7))]);
            }
          }
          return _coll6;
        }()).length)), (_10_v1).f7);
        _205_v132 = _nw27;
        let _207_v133;
        _207_v133 = _module.D1.create_DC5(_module.__default.safeDivisionInt((_187_v119)[_module.__default.safeIndex(new BigNumber(229), new BigNumber((_187_v119).length))], _15_v4), (_205_v132.f8).Merge(_16_v5), (((_205_v132.f8).contains(_15_v4)) ? ((_205_v132.f8).get(_15_v4)) : (_15_v4)));
        let _source2 = _207_v133;
        if (_source2.is_DC5) {
          let _208___mcc_h17 = (_source2).cf11;
          let _209___mcc_h18 = (_source2).cf12;
          let _210___mcc_h19 = (_source2).cf13;
          let _211_cf13 = _210___mcc_h19;
          let _212_cf12 = _209___mcc_h18;
          let _213_cf11 = _208___mcc_h17;
          let _214_v134;
          let _215_v135;
          let _216_v136;
          let _out15;
          let _out16;
          let _out17;
          let _outcollector5 = (_10_v1).m0((_187_v119)[_module.__default.safeIndex(new BigNumber(229), new BigNumber((_187_v119).length))], _213_cf11, _8_globalState);
          _out15 = _outcollector5[0];
          _out16 = _outcollector5[1];
          _out17 = _outcollector5[2];
          _214_v134 = _out15;
          _215_v135 = _out16;
          _216_v136 = _out17;
          let _217_v137;
          let _nw28 = new _module.C0();
          _nw28.__ctor((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_191_v121).length),new BigNumber((_95_v47).length))).Merge(_212_cf12), true);
          _217_v137 = _nw28;
          let _218_v138;
          _218_v138 = _dafny.MultiSet.fromElements((_10_v1).f6);
          _218_v138 = _218_v138;
          _9_v0 = (_95_v47)[_module.__default.safeIndex(_213_cf11, new BigNumber((_95_v47).length))];
        } else if (_source2.is_DC6) {
          let _219___mcc_h20 = (_source2).cf14;
          let _220___mcc_h21 = (_source2).cf15;
          let _221_cf15 = _220___mcc_h21;
          let _222_cf14 = _219___mcc_h20;
          let _rhs32 = (_205_v132).f9;
          let _rhs33 = _module.__default.safeDivisionInt(new BigNumber(596), _15_v4);
          let _lhs27 = _8_globalState;
          let _lhs28 = _8_globalState;
          _lhs27.f5 = _rhs32;
          _lhs28.f1 = _rhs33;
          let _223_v139;
          let _nw29 = Array((new BigNumber(12)).toNumber());
          _nw29[(_dafny.ZERO).toNumber()] = _205_v132;
          _nw29[(_dafny.ONE).toNumber()] = _205_v132;
          _nw29[(new BigNumber(2)).toNumber()] = _205_v132;
          _nw29[(new BigNumber(3)).toNumber()] = _205_v132;
          _nw29[(new BigNumber(4)).toNumber()] = _205_v132;
          _nw29[(new BigNumber(5)).toNumber()] = _205_v132;
          _nw29[(new BigNumber(6)).toNumber()] = ((true) ? (_205_v132) : (_205_v132));
          _nw29[(new BigNumber(7)).toNumber()] = _205_v132;
          _nw29[(new BigNumber(8)).toNumber()] = _205_v132;
          _nw29[(new BigNumber(9)).toNumber()] = _205_v132;
          _nw29[(new BigNumber(10)).toNumber()] = _205_v132;
          _nw29[(new BigNumber(11)).toNumber()] = _205_v132;
          _223_v139 = _nw29;
          let _index45 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_223_v139).length));
          (_223_v139)[_index45] = _205_v132;
          let _224_v140;
          let _nw30 = Array((new BigNumber(5)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _224_v140 = _nw30;
          let _225_v141;
          let _nw31 = Array((new BigNumber(17)).toNumber());
          _nw31[(_dafny.ZERO).toNumber()] = _224_v140;
          _nw31[(_dafny.ONE).toNumber()] = _224_v140;
          _nw31[(new BigNumber(2)).toNumber()] = _224_v140;
          _nw31[(new BigNumber(3)).toNumber()] = _224_v140;
          _nw31[(new BigNumber(4)).toNumber()] = _224_v140;
          _nw31[(new BigNumber(5)).toNumber()] = _224_v140;
          _nw31[(new BigNumber(6)).toNumber()] = _224_v140;
          _nw31[(new BigNumber(7)).toNumber()] = _224_v140;
          _nw31[(new BigNumber(8)).toNumber()] = _224_v140;
          _nw31[(new BigNumber(9)).toNumber()] = (((_10_v1).f7) ? (_224_v140) : (_224_v140));
          _nw31[(new BigNumber(10)).toNumber()] = _224_v140;
          _nw31[(new BigNumber(11)).toNumber()] = _224_v140;
          _nw31[(new BigNumber(12)).toNumber()] = _224_v140;
          _nw31[(new BigNumber(13)).toNumber()] = _224_v140;
          _nw31[(new BigNumber(14)).toNumber()] = _224_v140;
          _nw31[(new BigNumber(15)).toNumber()] = _224_v140;
          _nw31[(new BigNumber(16)).toNumber()] = _224_v140;
          _225_v141 = _nw31;
          let _index46 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_223_v139).length));
          let _rhs34 = _205_v132;
          let _rhs35 = (_187_v119)[_module.__default.safeIndex(new BigNumber(229), new BigNumber((_187_v119).length))];
          let _rhs36 = _225_v141;
          let _rhs37 = (_187_v119)[_module.__default.safeIndex(new BigNumber(229), new BigNumber((_187_v119).length))];
          let _rhs38 = !((_10_v1).f7);
          let _lhs29 = _223_v139;
          let _lhs30 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_223_v139).length));
          let _lhs31 = _8_globalState;
          let _lhs32 = _8_globalState;
          let _lhs33 = _8_globalState;
          _lhs29[_lhs30] = _rhs34;
          _lhs31.f1 = _rhs35;
          _225_v141 = _rhs36;
          _lhs32.f1 = _rhs37;
          _lhs33.f5 = _rhs38;
          _221_cf15 = (_10_v1).f6;
          (_8_globalState).f5 = ((_205_v132).f9) || ((_10_v1).f7);
        } else if (_source2.is_DC7) {
          let _226___mcc_h22 = (_source2).cf16;
          let _227_cf16 = _226___mcc_h22;
          let _228_v142;
          let _229_v143;
          let _230_v144;
          let _out18;
          let _out19;
          let _out20;
          let _outcollector6 = (_10_v1).m0((_187_v119)[_module.__default.safeIndex(new BigNumber(229), new BigNumber((_187_v119).length))], _15_v4, _8_globalState);
          _out18 = _outcollector6[0];
          _out19 = _outcollector6[1];
          _out20 = _outcollector6[2];
          _228_v142 = _out18;
          _229_v143 = _out19;
          _230_v144 = _out20;
          let _231_v145;
          _231_v145 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements((_10_v1).f7, (_10_v1).f7, (_10_v1).f7, (_10_v1).f7),_15_v4);
          _231_v145 = (_231_v145).update(_module.__default.fm10(new BigNumber(105), _dafny.Set.fromElements((_10_v1).f7, false, (_10_v1).f7, (_205_v132).f9, !(!(!((_10_v1).f7)))), _8_globalState), ((!((_10_v1).f7)) ? (_15_v4) : ((_187_v119)[_module.__default.safeIndex(new BigNumber(229), new BigNumber((_187_v119).length))])));
          let _232_v146;
          _232_v146 = _module.D2.create_DC10((_205_v132).f9, (_10_v1).f7);
          _232_v146 = _module.D2.create_DC10((_205_v132).f9, (_10_v1).f7);
          let _233_v147;
          _233_v147 = _dafny.MultiSet.fromElements((_10_v1).f7, (_205_v132).f9);
          let _index47 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_18_v7).length));
          (_18_v7)[_index47] = (_15_v4).isLessThanOrEqualTo((_dafny.ZERO).minus(new BigNumber((_233_v147).cardinality())));
          let _index48 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_18_v7).length));
          let _rhs39 = !((_187_v119)[_module.__default.safeIndex(new BigNumber(229), new BigNumber((_187_v119).length))]).isEqualTo((_187_v119)[_module.__default.safeIndex(new BigNumber(229), new BigNumber((_187_v119).length))]);
          let _rhs40 = (_205_v132).f9;
          let _rhs41 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_228_v142, (_228_v142).minus(new BigNumber((_195_v123).length))));
          let _lhs34 = _8_globalState;
          let _lhs35 = _18_v7;
          let _lhs36 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_18_v7).length));
          let _lhs37 = _8_globalState;
          _lhs34.f5 = _rhs39;
          _lhs35[_lhs36] = _rhs40;
          _lhs37.f0 = _rhs41;
        } else {
          let _234___mcc_h23 = (_source2).cf10;
          let _235_cf10 = _234___mcc_h23;
          let _index49 = _module.__default.safeIndex(new BigNumber(137), new BigNumber((_18_v7).length));
          (_18_v7)[_index49] = _235_cf10;
          let _index50 = _module.__default.safeIndex(new BigNumber(137), new BigNumber((_18_v7).length));
          (_18_v7)[_index50] = _235_cf10;
          let _236_v148;
          _236_v148 = _module.D2.create_DC10(true, (_205_v132).fm4(_8_globalState));
          let _index51 = _module.__default.safeIndex(new BigNumber(137), new BigNumber((_18_v7).length));
          (_18_v7)[_index51] = (_236_v148).dtor_cf20;
          _9_v0 = (_10_v1).f6;
          (_8_globalState).f5 = (_235_cf10) === ((_205_v132).f9);
        }
        let _237_v149;
        let _nw32 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _237_v149 = _nw32;
        let _index52 = _module.__default.safeIndex(new BigNumber(774), new BigNumber((_237_v149).length));
        (_237_v149)[_index52] = _dafny.Seq.Concat(_95_v47, _95_v47);
        let _238_v150;
        _238_v150 = _module.D2.create_DC10((_10_v1).f7, (_10_v1).f7);
        let _239_v151;
        _239_v151 = _module.D3.create_DC14(_10_v1, (_10_v1).f6, _10_v1, _205_v132, (_205_v132).f9);
        let _240_v152;
        _240_v152 = _dafny.Set.fromElements(_module.D3.create_DC14(_10_v1, (_10_v1).f6, _10_v1, _205_v132, (_205_v132).f9), _239_v151);
        let _index53 = _module.__default.safeIndex(new BigNumber(229), new BigNumber((_187_v119).length));
        let _index54 = _module.__default.safeIndex(new BigNumber(774), new BigNumber((_237_v149).length));
        let _rhs42 = !((!(!((_238_v150).dtor_cf19))) && ((_240_v152).IsDisjointFrom(_240_v152)));
        let _rhs43 = (_187_v119)[_module.__default.safeIndex(new BigNumber(229), new BigNumber((_187_v119).length))];
        let _rhs44 = _95_v47;
        let _lhs38 = _8_globalState;
        let _lhs39 = _187_v119;
        let _lhs40 = _module.__default.safeIndex(new BigNumber(229), new BigNumber((_187_v119).length));
        let _lhs41 = _237_v149;
        let _lhs42 = _module.__default.safeIndex(new BigNumber(774), new BigNumber((_237_v149).length));
        _lhs38.f5 = _rhs42;
        _lhs39[_lhs40] = _rhs43;
        _lhs41[_lhs42] = _rhs44;
      } else {
        let _241_v153;
        _241_v153 = _dafny.Seq.of(_95_v47, _dafny.Seq.Concat(_95_v47, _95_v47), _95_v47);
        let _242_v154;
        _242_v154 = _module.D4.create_DC16(_241_v153);
        _241_v153 = _dafny.Seq.Concat(_dafny.Seq.Concat(_241_v153, _dafny.Seq.update((_242_v154).dtor_cf31, _module.__default.safeIndex((_dafny.ZERO).minus((_187_v119)[_module.__default.safeIndex(new BigNumber(229), new BigNumber((_187_v119).length))]), new BigNumber(((_242_v154).dtor_cf31).length)), _95_v47)), _241_v153);
        let _243_v155;
        _243_v155 = _module.D5.create_DC18(_187_v119);
        _187_v119 = (_243_v155).dtor_cf37;
        let _244_v156;
        let _245_v157;
        let _246_v158;
        let _out21;
        let _out22;
        let _out23;
        let _outcollector7 = (_10_v1).m0((_187_v119)[_module.__default.safeIndex(new BigNumber(229), new BigNumber((_187_v119).length))], (((_195_v123).contains((_10_v1).f7)) ? ((_195_v123).get((_10_v1).f7)) : (_15_v4)), _8_globalState);
        _out21 = _outcollector7[0];
        _out22 = _outcollector7[1];
        _out23 = _outcollector7[2];
        _244_v156 = _out21;
        _245_v157 = _out22;
        _246_v158 = _out23;
        (_8_globalState).f5 = (_10_v1).f7;
        let _index55 = _module.__default.safeIndex(new BigNumber(704), new BigNumber((_18_v7).length));
        (_18_v7)[_index55] = (_10_v1).f7;
        let _247_v159;
        let _nw33 = new _module.C0();
        _nw33.__ctor(_dafny.Map.Empty.slice().updateUnsafe((_187_v119)[_module.__default.safeIndex(new BigNumber(229), new BigNumber((_187_v119).length))],_245_v157), (_10_v1).f7);
        _247_v159 = _nw33;
        let _248_v160;
        _248_v160 = _dafny.Seq.of(_247_v159, _247_v159, _247_v159);
        let _index56 = _module.__default.safeIndex(new BigNumber(704), new BigNumber((_18_v7).length));
        let _index57 = _module.__default.safeIndex(new BigNumber(229), new BigNumber((_187_v119).length));
        let _rhs45 = _245_v157;
        let _rhs46 = _module.__default.fm6(_8_globalState);
        let _rhs47 = ((_dafny.ZERO).minus(new BigNumber((_248_v160).length))).isLessThan(_15_v4);
        let _rhs48 = (_245_v157).minus(_245_v157);
        let _rhs49 = new BigNumber(943);
        let _lhs43 = _8_globalState;
        let _lhs44 = _18_v7;
        let _lhs45 = _module.__default.safeIndex(new BigNumber(704), new BigNumber((_18_v7).length));
        let _lhs46 = _187_v119;
        let _lhs47 = _module.__default.safeIndex(new BigNumber(229), new BigNumber((_187_v119).length));
        _lhs43.f1 = _rhs45;
        _9_v0 = _rhs46;
        _lhs44[_lhs45] = _rhs47;
        _244_v156 = _rhs48;
        _lhs46[_lhs47] = _rhs49;
      }
      let _249_v161;
      _249_v161 = _dafny.Set.fromElements(new BigNumber(110), (_187_v119)[_module.__default.safeIndex(new BigNumber(229), new BigNumber((_187_v119).length))]);
      let _250_v162;
      _250_v162 = _dafny.MultiSet.fromElements(_95_v47);
      let _251_v163;
      let _252_v164;
      let _253_v165;
      let _out24;
      let _out25;
      let _out26;
      let _outcollector8 = (_10_v1).m0((new BigNumber((_249_v161).length)).minus(new BigNumber((_250_v162).cardinality())), (((_10_v1).f7) ? ((_187_v119)[_module.__default.safeIndex(new BigNumber(229), new BigNumber((_187_v119).length))]) : (_15_v4)), _8_globalState);
      _out24 = _outcollector8[0];
      _out25 = _outcollector8[1];
      _out26 = _outcollector8[2];
      _251_v163 = _out24;
      _252_v164 = _out25;
      _253_v165 = _out26;
      _252_v164 = new BigNumber(749);
      process.stdout.write(_dafny.toString(_8_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_8_globalState.f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_8_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_8_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_8_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_8_globalState.f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_9_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_10_v1).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_10_v1).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_13_v2).dtor_cf14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_13_v2).dtor_cf15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_14_v3, _dafny.Seq.of(_module.D1.create_DC6(new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0))), _module.D1.create_DC6(new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0))), _module.D1.create_DC6(new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0))), _module.D1.create_DC6(new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0))), _module.D1.create_DC6(new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_15_v4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_16_v5).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(22),new BigNumber(22)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_17_v6).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_v7)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_v7)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_v7)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_v7)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_v7)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_v7)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_v7)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_v7)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_v7)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_v7)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_v7)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_v7)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_v7)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_v7)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_v7)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_18_v7)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_21_v8).dtor_cf6).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(22),new BigNumber(22)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_21_v8).dtor_cf7).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_21_v8).dtor_cf8)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_21_v8).dtor_cf8)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_21_v8).dtor_cf8)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_21_v8).dtor_cf8)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_21_v8).dtor_cf8)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_21_v8).dtor_cf8)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_21_v8).dtor_cf8)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_21_v8).dtor_cf8)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_21_v8).dtor_cf8)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_21_v8).dtor_cf8)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_21_v8).dtor_cf8)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_21_v8).dtor_cf8)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_21_v8).dtor_cf8)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_21_v8).dtor_cf8)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_21_v8).dtor_cf8)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_21_v8).dtor_cf8)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_95_v47).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_149_i10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v118).equals(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_187_v119)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_187_v119)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_187_v119)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_187_v119)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_187_v119)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_187_v119)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_187_v119)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_187_v119)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_187_v119)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_187_v119)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_187_v119)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_190_v120).equals(_dafny.MultiSet.fromElements(new BigNumber(22), new BigNumber(22)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_191_v121, _dafny.Seq.of(new BigNumber(6)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_194_v122, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_195_v123).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(872)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_196_v124).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_197_v125).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_module.D0.create_DC3(new BigNumber(22))).updateUnsafe(false,_module.D0.create_DC3(new BigNumber(22))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_249_v161).equals(_dafny.Set.fromElements(new BigNumber(110), new BigNumber(872)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_250_v162).equals(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("vintro")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_251_v163));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_252_v164));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1, cf2, cf3, cf4, cf5) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC2(cf6, cf7, cf8) {
      let $dt = new D0(1);
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC3(cf9) {
      let $dt = new D0(2);
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(3);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get is_DC0() { return this.$tag === 3; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC3" + "(" + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 3) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1) && _dafny.areEqual(this.cf2, other.cf2) && _dafny.areEqual(this.cf3, other.cf3) && this.cf4 === other.cf4 && _dafny.areEqual(this.cf5, other.cf5);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf6, other.cf6) && _dafny.areEqual(this.cf7, other.cf7) && this.cf8 === other.cf8;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf9, other.cf9);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf0 === other.cf0;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(_dafny.ZERO, _dafny.Set.Empty, _dafny.ZERO, [], new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC5(cf11, cf12, cf13) {
      let $dt = new D1(0);
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC6(cf14, cf15) {
      let $dt = new D1(1);
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC7(cf16) {
      let $dt = new D1(2);
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC4(cf10) {
      let $dt = new D1(3);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get is_DC7() { return this.$tag === 2; }
    get is_DC4() { return this.$tag === 3; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC6" + "(" + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC7" + "(" + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 3) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf10) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf11, other.cf11) && _dafny.areEqual(this.cf12, other.cf12) && _dafny.areEqual(this.cf13, other.cf13);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf14, other.cf14) && _dafny.areEqual(this.cf15, other.cf15);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf16, other.cf16);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf10 === other.cf10;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC5(_dafny.ZERO, _dafny.Map.Empty, _dafny.ZERO);
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
    static create_DC9(cf18) {
      let $dt = new D2(0);
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC10(cf19, cf20) {
      let $dt = new D2(1);
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC11(cf21, cf22) {
      let $dt = new D2(2);
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC8(cf17) {
      let $dt = new D2(3);
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC12(cf23) {
      let $dt = new D2(4);
      $dt.cf23 = cf23;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get is_DC11() { return this.$tag === 2; }
    get is_DC8() { return this.$tag === 3; }
    get is_DC12() { return this.$tag === 4; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf23() { return this.cf23; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC9" + "(" + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC10" + "(" + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC11" + "(" + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 3) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 4) {
        return "D2.DC12" + "(" + _dafny.toString(this.cf23) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf18, other.cf18);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf19 === other.cf19 && this.cf20 === other.cf20;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf21, other.cf21) && _dafny.areEqual(this.cf22, other.cf22);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf17 === other.cf17;
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf23, other.cf23);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC9(new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC14(cf25, cf26, cf27, cf28, cf29) {
      let $dt = new D3(0);
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC13(cf24) {
      let $dt = new D3(1);
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC15(cf30) {
      let $dt = new D3(2);
      $dt.cf30 = cf30;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get is_DC15() { return this.$tag === 2; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf30() { return this.cf30; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC14" + "(" + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC13" + "(" + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC15" + "(" + _dafny.toString(this.cf30) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf25 === other.cf25 && _dafny.areEqual(this.cf26, other.cf26) && this.cf27 === other.cf27 && this.cf28 === other.cf28 && this.cf29 === other.cf29;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf24 === other.cf24;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf30, other.cf30);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC14(null, new _dafny.CodePoint('D'.codePointAt(0)), null, null, false);
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
    static create_DC17(cf32, cf33, cf34, cf35, cf36) {
      let $dt = new D4(0);
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      return $dt;
    }
    static create_DC16(cf31) {
      let $dt = new D4(1);
      $dt.cf31 = cf31;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf31() { return this.cf31; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC17" + "(" + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC16" + "(" + _dafny.toString(this.cf31) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf32, other.cf32) && this.cf33 === other.cf33 && _dafny.areEqual(this.cf34, other.cf34) && _dafny.areEqual(this.cf35, other.cf35) && this.cf36 === other.cf36;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf31, other.cf31);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC17(_dafny.ZERO, false, _dafny.ZERO, _dafny.ZERO, null);
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
    static create_DC19(cf38, cf39, cf40) {
      let $dt = new D5(0);
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      $dt.cf40 = cf40;
      return $dt;
    }
    static create_DC18(cf37) {
      let $dt = new D5(1);
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC20(cf41) {
      let $dt = new D5(2);
      $dt.cf41 = cf41;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get is_DC18() { return this.$tag === 1; }
    get is_DC20() { return this.$tag === 2; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf41() { return this.cf41; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC19" + "(" + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ", " + _dafny.toString(this.cf40) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC18" + "(" + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC20" + "(" + _dafny.toString(this.cf41) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf38 === other.cf38 && this.cf39 === other.cf39 && this.cf40 === other.cf40;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf37 === other.cf37;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf41, other.cf41);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC19(false, false, false);
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
    static create_DC22(cf43) {
      let $dt = new D6(0);
      $dt.cf43 = cf43;
      return $dt;
    }
    static create_DC23(cf44, cf45, cf46) {
      let $dt = new D6(1);
      $dt.cf44 = cf44;
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      return $dt;
    }
    static create_DC21(cf42) {
      let $dt = new D6(2);
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC24(cf47) {
      let $dt = new D6(3);
      $dt.cf47 = cf47;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get is_DC23() { return this.$tag === 1; }
    get is_DC21() { return this.$tag === 2; }
    get is_DC24() { return this.$tag === 3; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf47() { return this.cf47; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC22" + "(" + _dafny.toString(this.cf43) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC23" + "(" + _dafny.toString(this.cf44) + ", " + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC21" + "(" + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 3) {
        return "D6.DC24" + "(" + _dafny.toString(this.cf47) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf43 === other.cf43;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf44, other.cf44) && _dafny.areEqual(this.cf45, other.cf45) && _dafny.areEqual(this.cf46, other.cf46);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf42, other.cf42);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf47, other.cf47);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC22(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D6.Default();
        }
      };
    }
  }

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f0 = _dafny.ZERO;
      this.f1 = _dafny.ZERO;
      this.f5 = false;
      this._f2 = false;
      this._f3 = _dafny.ZERO;
      this._f4 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5) {
      let _this = this;
      (_this).f0 = f0;
      (_this).f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this).f5 = f5;
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
    get f4() {
      let _this = this;
      return _this._f4;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this.f8 = _dafny.Map.Empty;
      this._f9 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f8, f9) {
      let _this = this;
      (_this).f8 = f8;
      (_this)._f9 = f9;
      return;
    }
    fm3(p0, p1, p2, globalState) {
      let _this = this;
      return ((_dafny.ZERO).minus(new BigNumber(-152))).minus((new BigNumber(851)).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f9,new BigNumber(940))).length)));
    };
    fm4(globalState) {
      let _this = this;
      return ((_dafny.MultiSet.fromElements(_dafny.Seq.of(true), _dafny.Seq.of((_this).f9), _dafny.Seq.of((_this).f9, (_this).f9))).Union(_dafny.MultiSet.fromElements(_dafny.Seq.of(false, (_this).f9), _dafny.Seq.of((_this).f9)))).IsDisjointFrom((_dafny.MultiSet.fromElements(_dafny.Seq.of((_this).f9), _dafny.Seq.of((_this).f9))).Difference(_dafny.MultiSet.fromElements(_dafny.Seq.of((_this).f9), _dafny.Seq.of((_this).f9), _dafny.Seq.of((_this).f9, false, (_this).f9))));
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f6 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f7 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f6, f7) {
      let _this = this;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      return;
    }
    fm0(p0, p1, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-188)), function (_254_i0) {
        return (_this).f6;
      }), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qkbbvoi"), _dafny.Seq.UnicodeFromString("dcitg")));
    };
    m0(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = [];
      if ((_this).f7) {
        let _255_v0;
        let _init5 = ((_256_p1) => function (_257_i0) {
          return (((_this).f7) ? (_dafny.Map.Empty.slice().updateUnsafe(_256_p1,_dafny.Seq.of((_this).f7))) : (_dafny.Map.Empty.slice().updateUnsafe(_256_p1,_dafny.Seq.of((_this).f7, (_this).f7))));
        })(p1);
        let _nw34 = Array((new BigNumber(21)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw34.length); _i0_5++) {
          _nw34[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _255_v0 = _nw34;
        let _258_v1;
        _258_v1 = _dafny.Seq.of((_this).f7);
        let _259_v2;
        _259_v2 = _dafny.Map.Empty.slice().updateUnsafe(p0,_258_v1);
        let _index58 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_255_v0).length));
        (_255_v0)[_index58] = _259_v2;
        let _260_v3;
        let _nw35 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
        _260_v3 = _nw35;
        let _index59 = _module.__default.safeIndex(new BigNumber(365), new BigNumber((_260_v3).length));
        (_260_v3)[_index59] = p1;
        let _index60 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_255_v0).length));
        let _index61 = _module.__default.safeIndex(new BigNumber(365), new BigNumber((_260_v3).length));
        let _rhs50 = _module.__default.safeModuloInt(p1, new BigNumber((_258_v1).length));
        let _rhs51 = ((_module.__default.fm1(globalState)) ? (_dafny.Map.Empty.slice().updateUnsafe(p1,_258_v1)) : ((_259_v2).Merge(_259_v2)));
        let _rhs52 = new BigNumber(48);
        let _rhs53 = new BigNumber(682);
        let _lhs48 = globalState;
        let _lhs49 = _255_v0;
        let _lhs50 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_255_v0).length));
        let _lhs51 = _260_v3;
        let _lhs52 = _module.__default.safeIndex(new BigNumber(365), new BigNumber((_260_v3).length));
        _lhs48.f1 = _rhs50;
        _lhs49[_lhs50] = _rhs51;
        r1 = _rhs52;
        _lhs51[_lhs52] = _rhs53;
        (globalState).f5 = ((false) ? ((_this).f7) : ((_this).f7));
        let _261_v4;
        _261_v4 = _dafny.Seq.UnicodeFromString("cfkjko");
        let _262_v5;
        _262_v5 = _dafny.Map.Empty.slice().updateUnsafe(_261_v4,(_260_v3)[_module.__default.safeIndex(new BigNumber(365), new BigNumber((_260_v3).length))]);
        _262_v5 = (_262_v5).update(_261_v4, (_260_v3)[_module.__default.safeIndex(new BigNumber(365), new BigNumber((_260_v3).length))]);
        let _263_v6;
        _263_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,p1);
        let _264_v7;
        _264_v7 = _dafny.Map.Empty.slice().updateUnsafe(_261_v4,_dafny.Map.Empty.slice().updateUnsafe((_this).f7,(_this).f7));
        let _265_v8;
        _265_v8 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f7),(_this).f7);
        (globalState).f1 = ((!(_263_v6).contains((_this).f7)) ? (_module.__default.fm2(_module.__default.fm1(globalState), globalState)) : ((((_this).f7) ? (new BigNumber(((((_264_v7).contains(_261_v4)) ? ((_264_v7).get(_261_v4)) : (_265_v8))).length)) : ((_260_v3)[_module.__default.safeIndex(new BigNumber(365), new BigNumber((_260_v3).length))]))));
        let _266_v9;
        _266_v9 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,((_this).f7) || ((_this).f7));
        _266_v9 = (_266_v9).update((_this).f6, false);
      } else {
        let _267_v10;
        _267_v10 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2((_this).f7, globalState),(_this).f7);
        let _268_v11;
        _268_v11 = _dafny.MultiSet.fromElements((_this).f7);
        let _269_v12;
        let _nw36 = new _module.C0();
        _nw36.__ctor(_module.__default.fm5(new BigNumber((_267_v10).length), (_this).f7, (_this).f7, globalState), (_dafny.MultiSet.fromElements((_this).f7, (_this).f7)).equals(_268_v11));
        _269_v12 = _nw36;
        let _270_v13;
        let _nw37 = Array((new BigNumber(28)).toNumber()).fill(false);
        _270_v13 = _nw37;
        let _271_v14;
        _271_v14 = _module.D0.create_DC0(_270_v13);
        _270_v13 = (_271_v14).dtor_cf0;
        let _272_v15;
        _272_v15 = _dafny.Map.Empty.slice().updateUnsafe((_269_v12).f9,p0);
        _272_v15 = (_272_v15).update((_269_v12).f9, _module.__default.fm2((_this).f7, globalState));
        let _273_v16;
        _273_v16 = new _dafny.CodePoint('q'.codePointAt(0));
        _273_v16 = ((!((_269_v12).f9)) ? ((_this).f6) : ((_this).f6));
        let _274_v17;
        let _nw38 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Set.Empty);
        _274_v17 = _nw38;
        let _275_v18;
        _275_v18 = _dafny.Set.fromElements(_270_v13, _270_v13);
        let _276_v19;
        _276_v19 = _dafny.Map.Empty.slice().updateUnsafe(_275_v18,false);
        _274_v17 = (((((_269_v12).f9) ? ((((_276_v19).contains(_275_v18)) ? ((_276_v19).get(_275_v18)) : ((_269_v12).fm4(globalState)))) : ((_this).f7))) ? (_274_v17) : (_274_v17));
      }
      let _277_v20;
      _277_v20 = _dafny.Seq.UnicodeFromString("swfjwfa");
      let _278_v21;
      _278_v21 = _module.D1.create_DC4(false);
      let _279_v22;
      _279_v22 = _dafny.Set.fromElements(p0, new BigNumber((_277_v20).length), p1);
      let _280_v23;
      _280_v23 = _dafny.MultiSet.fromElements(new BigNumber((_279_v22).length));
      let _281_v24;
      let _nw39 = Array((new BigNumber(15)).toNumber());
      _nw39[(_dafny.ZERO).toNumber()] = (_this).f7;
      _nw39[(_dafny.ONE).toNumber()] = !((_this).f7);
      _nw39[(new BigNumber(2)).toNumber()] = (p0).isLessThan(p1);
      _nw39[(new BigNumber(3)).toNumber()] = !(_dafny.areEqual(_277_v20, _277_v20));
      _nw39[(new BigNumber(4)).toNumber()] = !((_278_v21).dtor_cf10);
      _nw39[(new BigNumber(5)).toNumber()] = (_this).f7;
      _nw39[(new BigNumber(6)).toNumber()] = !(_module.__default.fm1(globalState));
      _nw39[(new BigNumber(7)).toNumber()] = (_this).f7;
      _nw39[(new BigNumber(8)).toNumber()] = !((_this).f7);
      _nw39[(new BigNumber(9)).toNumber()] = (_280_v23).equals(_280_v23);
      _nw39[(new BigNumber(10)).toNumber()] = _module.__default.fm1(globalState);
      _nw39[(new BigNumber(11)).toNumber()] = (_this).f7;
      _nw39[(new BigNumber(12)).toNumber()] = (_this).f7;
      _nw39[(new BigNumber(13)).toNumber()] = (_this).f7;
      _nw39[(new BigNumber(14)).toNumber()] = (_this).f7;
      _281_v24 = _nw39;
      let _index62 = _module.__default.safeIndex(new BigNumber(448), new BigNumber((_281_v24).length));
      (_281_v24)[_index62] = (p0).isLessThan(_module.__default.fm2((_this).f7, globalState));
      let _index63 = _module.__default.safeIndex(new BigNumber(448), new BigNumber((_281_v24).length));
      (_281_v24)[_index63] = (_this).f7;
      (globalState).f5 = (_this).f7;
      let _282_v25;
      _282_v25 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-673),p1);
      let _283_v26;
      _283_v26 = _dafny.Set.fromElements(_dafny.Seq.update(_dafny.Seq.Concat(_277_v20, _277_v20), _module.__default.safeIndex(new BigNumber((_282_v25).length), new BigNumber((_dafny.Seq.Concat(_277_v20, _277_v20)).length)), new _dafny.CodePoint('h'.codePointAt(0))));
      r1 = new BigNumber((_283_v26).length);
      let _284_v27;
      _284_v27 = _dafny.Set.fromElements((_this).f7, (_this).f7, (_this).f7);
      let _285_v28;
      let _nw40 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Seq.of());
      _285_v28 = _nw40;
      let _286_v29;
      _286_v29 = _module.D0.create_DC1(p1, _284_v27, p0, _285_v28, (_this).f6);
      let _287_v30;
      let _nw41 = Array((new BigNumber(19)).toNumber());
      _nw41[(_dafny.ZERO).toNumber()] = _286_v29;
      _nw41[(_dafny.ONE).toNumber()] = _286_v29;
      _nw41[(new BigNumber(2)).toNumber()] = _286_v29;
      _nw41[(new BigNumber(3)).toNumber()] = _286_v29;
      _nw41[(new BigNumber(4)).toNumber()] = _286_v29;
      _nw41[(new BigNumber(5)).toNumber()] = _286_v29;
      _nw41[(new BigNumber(6)).toNumber()] = _286_v29;
      _nw41[(new BigNumber(7)).toNumber()] = _286_v29;
      _nw41[(new BigNumber(8)).toNumber()] = _286_v29;
      _nw41[(new BigNumber(9)).toNumber()] = _286_v29;
      _nw41[(new BigNumber(10)).toNumber()] = _286_v29;
      _nw41[(new BigNumber(11)).toNumber()] = _286_v29;
      _nw41[(new BigNumber(12)).toNumber()] = _286_v29;
      _nw41[(new BigNumber(13)).toNumber()] = _286_v29;
      _nw41[(new BigNumber(14)).toNumber()] = _286_v29;
      _nw41[(new BigNumber(15)).toNumber()] = _286_v29;
      _nw41[(new BigNumber(16)).toNumber()] = _286_v29;
      _nw41[(new BigNumber(17)).toNumber()] = _286_v29;
      _nw41[(new BigNumber(18)).toNumber()] = _286_v29;
      _287_v30 = _nw41;
      let _288_v31;
      _288_v31 = _dafny.Map.Empty.slice().updateUnsafe(_287_v30,(_this).f7);
      _288_v31 = (_288_v31).update(_287_v30, (_281_v24)[_module.__default.safeIndex(new BigNumber(448), new BigNumber((_281_v24).length))]);
      (globalState).f5 = (_281_v24)[_module.__default.safeIndex(new BigNumber(448), new BigNumber((_281_v24).length))];
      let _289_v32;
      _289_v32 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(192));
      r0 = _module.__default.safeModuloInt((((_281_v24)[_module.__default.safeIndex(new BigNumber(448), new BigNumber((_281_v24).length))]) ? (new BigNumber((_dafny.Set.fromElements(true)).length)) : (p1)), new BigNumber((_289_v32).length));
      let _290_v33;
      let _nw42 = new _module.C0();
      _nw42.__ctor(_282_v25, !((_this).f7));
      _290_v33 = _nw42;
      let _291_v34;
      let _nw43 = Array((new BigNumber(14)).toNumber());
      _nw43[(_dafny.ZERO).toNumber()] = _290_v33;
      _nw43[(_dafny.ONE).toNumber()] = _290_v33;
      _nw43[(new BigNumber(2)).toNumber()] = _290_v33;
      _nw43[(new BigNumber(3)).toNumber()] = _290_v33;
      _nw43[(new BigNumber(4)).toNumber()] = _290_v33;
      _nw43[(new BigNumber(5)).toNumber()] = _290_v33;
      _nw43[(new BigNumber(6)).toNumber()] = _290_v33;
      _nw43[(new BigNumber(7)).toNumber()] = _290_v33;
      _nw43[(new BigNumber(8)).toNumber()] = _290_v33;
      _nw43[(new BigNumber(9)).toNumber()] = _290_v33;
      _nw43[(new BigNumber(10)).toNumber()] = _290_v33;
      _nw43[(new BigNumber(11)).toNumber()] = _290_v33;
      _nw43[(new BigNumber(12)).toNumber()] = _290_v33;
      _nw43[(new BigNumber(13)).toNumber()] = _290_v33;
      _291_v34 = _nw43;
      let _292_v35;
      _292_v35 = _dafny.Map.Empty.slice().updateUnsafe(true,false);
      r1 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((((_this).f7) ? (_291_v34) : (_291_v34)),(((((_292_v35).contains((_this).f7)) ? ((_292_v35).get((_this).f7)) : ((_this).f7))) ? ((_this).f7) : (true)))).length);
      let _293_v36;
      let _nw44 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
      _293_v36 = _nw44;
      r2 = _293_v36;
      return [r0, r1, r2];
    }
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
