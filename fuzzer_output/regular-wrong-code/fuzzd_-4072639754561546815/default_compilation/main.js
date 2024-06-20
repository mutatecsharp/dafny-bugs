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
    static fm0(p0, globalState) {
      if (!(false)) {
        return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(863)), function (_0_i0) {
          return new _dafny.CodePoint('e'.codePointAt(0));
        }), _dafny.Seq.UnicodeFromString("qy"));
      } else {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(108)), function (_1_i1) {
          return new _dafny.CodePoint('d'.codePointAt(0));
        });
      }
    };
    static fm1(p0, globalState) {
      return ((new BigNumber(716)).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,(_module.D9.create_DC25((_dafny.ZERO).minus(new BigNumber(-341)), false, true, !(true))).dtor_cf44)).length),new _dafny.CodePoint('b'.codePointAt(0)))).length)))).plus(new BigNumber(28));
    };
    static fm2(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.of(_dafny.areEqual(_dafny.Seq.of(true, true), _dafny.Seq.of(false, false, false, false, !(true))), !(!_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(74)), new BigNumber(-680))), !(false));
    };
    static fm3(globalState) {
      return _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("e"), _dafny.Seq.UnicodeFromString("hipjlan")), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("csbbvc"), _dafny.Seq.UnicodeFromString("ejenrea")));
    };
    static fm4(p0, p1, globalState) {
      return _dafny.Seq.of(_dafny.MultiSet.fromElements(true, false, true), _dafny.MultiSet.fromElements(true, false, true, true));
    };
    static fm5(p0, p1, globalState) {
      if ((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(528)), function (_2_i0) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      })).length)).isLessThanOrEqualTo(new BigNumber((_dafny.Set.fromElements(false, false)).length))) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('v'.codePointAt(0));
      }
    };
    static fm7(p0, p1, globalState) {
      return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(986)), _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),true)).length)))).length);
    };
    static fm8(p0, p1, globalState) {
      return (new BigNumber(-717)).isEqualTo(new BigNumber(45));
    };
    static fm10(p0, p1, globalState) {
      return !(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("g"), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("wueqr"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(631)), function (_3_i0) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      }))));
    };
    static fm11(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(640)), _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(-931), new BigNumber(501)), _dafny.Seq.of(new BigNumber(97))));
    };
    static fm14(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(false, false, !(false), true, true), _dafny.Seq.of(true)), _dafny.Seq.Concat(_dafny.Seq.of(true, !(false)), _dafny.Seq.of(!(true), false)));
    };
    static fm15(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(((_module.D17.create_DC46(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-338)), function (_4_i0) {
  return new _dafny.CodePoint('d'.codePointAt(0));
})).length),new BigNumber((_dafny.MultiSet.fromElements(true, false)).cardinality())))).dtor_cf80).length))).length)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(861),true)).Keys.Elements) {
          let _5_v0 = _compr_0;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(861),true)).contains(_5_v0)) {
            _coll0.add(_module.__default.safeModuloInt(_5_v0, new BigNumber(63)));
          }
        }
        return _coll0;
      }()).length),!(true))).length), new BigNumber(190), new BigNumber((function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-462),new _dafny.CodePoint('g'.codePointAt(0)))).Keys.Elements) {
          let _6_v1 = _compr_1;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-462),new _dafny.CodePoint('g'.codePointAt(0)))).contains(_6_v1)) {
            _coll1.push([(_6_v1).multipliedBy(new BigNumber((_dafny.Seq.of(new BigNumber(103))).length)),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
              let _coll2 = new _dafny.Map();
              for (const _compr_2 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),false)).Keys.Elements) {
                let _7_v2 = _compr_2;
                if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),false)).contains(_7_v2)) {
                  _coll2.push([_7_v2,new BigNumber((_dafny.Seq.UnicodeFromString("xgonjaa")).length)]);
                }
              }
              return _coll2;
            }()).length),new BigNumber(106))]);
          }
        }
        return _coll1;
      }()).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(867)), function (_8_i1) {
        return new _dafny.CodePoint('r'.codePointAt(0));
      })).length)), _dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new _dafny.CodePoint('p'.codePointAt(0))))).cardinality()), new BigNumber(62), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(705)), function (_9_i2) {
        return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(369)), function (_10_i3) {
          return new _dafny.CodePoint('l'.codePointAt(0));
        })).length);
      })).length), new BigNumber((_dafny.Seq.UnicodeFromString("jbdseslwf")).length), new BigNumber((_dafny.Seq.UnicodeFromString("v")).length)));
    };
    static fm18(p0, p1, globalState) {
      return (_module.D17.create_DC47(!(false), (_dafny.ZERO).minus(new BigNumber(-432)))).dtor_cf82;
    };
    static fm19(p0, p1, globalState) {
      return (_module.D18.create_DC48(_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber(-652)))).dtor_cf83;
    };
    static fm20(p0, p1, p2, p3, globalState) {
      return new _dafny.CodePoint('m'.codePointAt(0));
    };
    static fm21(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(707),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(248),true));
    };
    static fm22(p0, globalState) {
      if (true) {
        return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(414)), function (_11_i0) {
          return _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-540),false)).length));
        }), _dafny.Seq.of(function () {
          let _coll3 = new _dafny.Set();
          for (const _compr_3 of _dafny.IntegerRange(new BigNumber(680), new BigNumber(4))) {
            let _12_v0 = _compr_3;
            if (((new BigNumber(680)).isLessThanOrEqualTo(_12_v0)) && ((_12_v0).isLessThan(new BigNumber(4)))) {
              _coll3.add((_12_v0).plus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("dqm")).length))).cardinality())));
            }
          }
          return _coll3;
        }()));
      } else {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(783)), function (_13_i1) {
          return _dafny.Set.fromElements(new BigNumber(322), (_dafny.ZERO).minus(new BigNumber(-322)), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-173))).cardinality()), (_dafny.ZERO).minus(new BigNumber(-996)), new BigNumber(-776));
        });
      }
    };
    static fm23(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(-110), new BigNumber(-249)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-74)), function (_14_i0) {
        return new BigNumber((_dafny.Seq.of(new BigNumber(916))).length);
      })), _dafny.Seq.of(new BigNumber(378), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(11),_dafny.Seq.UnicodeFromString("ucvrgkh"))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(178)), function (_15_i1) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      })).length),new _dafny.CodePoint('t'.codePointAt(0)))).length), new BigNumber(933)));
    };
    static fm24(p0, p1, p2, globalState) {
      return _dafny.MultiSet.fromElements(new BigNumber(32));
    };
    static fm25(p0, p1, p2, p3, globalState) {
      return (_dafny.Set.fromElements(_module.D2.create_DC6((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(!(false)))).cardinality()))))).Intersect(_dafny.Set.fromElements(_module.D2.create_DC6(new BigNumber((_dafny.Seq.UnicodeFromString("msuhavx")).length)), _module.D2.create_DC6(new BigNumber(-218)), _module.D2.create_DC6(new BigNumber((function () {
  let _coll4 = new _dafny.Map();
  for (const _compr_4 of _dafny.IntegerRange(new BigNumber(783), new BigNumber(-172))) {
    let _16_v0 = _compr_4;
    if (((new BigNumber(783)).isLessThanOrEqualTo(_16_v0)) && ((_16_v0).isLessThan(new BigNumber(-172)))) {
      _coll4.push([_module.__default.safeDivisionInt(_16_v0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(208)), function (_17_i0) {
        return new BigNumber((_dafny.Seq.of(new BigNumber(682), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("jlnaqhjk")).length)))).length);
      })).length)),false]);
    }
  }
  return _coll4;
}()).length)), _module.D2.create_DC6(new BigNumber(531))));
    };
    static fm26(p0, p1, p2, p3, globalState) {
      return (new BigNumber(855)).isEqualTo(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.UnicodeFromString("bfehftfco")).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)))).length))));
    };
    static fm27(p0, p1, p2, globalState) {
      if ((_dafny.MultiSet.fromElements(false, true, true)).IsDisjointFrom(_dafny.MultiSet.fromElements(true, false, false, !(!(false))))) {
        return _module.D9.create_DC23(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(984)), function (_18_i0) {
  return new _dafny.CodePoint('w'.codePointAt(0));
})));
      } else {
        return _module.D9.create_DC23(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("nttb")));
      }
    };
    static fm28(p0, p1, p2, globalState) {
      return _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-747))).length), new BigNumber(-663), _module.__default.safeDivisionInt(new BigNumber(-818), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.D9.create_DC25(new BigNumber(933), false, true, false),new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-790),false)).length), new BigNumber(80)))).cardinality()))).length)));
    };
    static fm29(p0, p1, globalState) {
      if (false) {
        return _module.D0.create_DC1(new _dafny.CodePoint('c'.codePointAt(0)), false, true);
      } else {
        return _module.D0.create_DC1(new _dafny.CodePoint('j'.codePointAt(0)), false, false);
      }
    };
    static fm30(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(!(true),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,true));
    };
    static fm31(p0, globalState) {
      return _module.D0.create_DC3(new _dafny.CodePoint('u'.codePointAt(0)));
    };
    static fm32(p0, p1, p2, p3, globalState) {
      return (function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of _dafny.IntegerRange(new BigNumber(-687), new BigNumber(640))) {
          let _19_v0 = _compr_5;
          if (((new BigNumber(-687)).isLessThanOrEqualTo(_19_v0)) && ((_19_v0).isLessThan(new BigNumber(640)))) {
            _coll5.push([(_19_v0).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(379)), function (_20_i0) {
              return new BigNumber((_dafny.Seq.of(new BigNumber(604))).length);
            })).length)),new BigNumber(-983)]);
          }
        }
        return _coll5;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-585))).cardinality()),new BigNumber(-836)));
    };
    static fm33(p0, p1, p2, globalState) {
      return _module.D8.create_DC21(((true) ? (!(false)) : (true)), !(!((true) && (true))), new BigNumber(((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("k")).length))).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("td")).length))))).cardinality()));
    };
    static fm34(p0, globalState) {
      return (_module.D19.create_DC52(_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(!(false), true, true)).length)), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
  let _coll6 = new _dafny.Map();
  for (const _compr_6 of (_dafny.Seq.of(new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)))).Elements) {
    let _21_v0 = _compr_6;
    if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0))), _21_v0)) {
      _coll6.push([_21_v0,false]);
    }
  }
  return _coll6;
}()).length),new BigNumber(965))).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(719))).cardinality()), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(461)), function (_22_i0) {
  return new _dafny.CodePoint('q'.codePointAt(0));
})).length))), _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("vo")).length)), _dafny.MultiSet.fromElements(new BigNumber(69)), _dafny.MultiSet.fromElements(new BigNumber(856), new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()))))).dtor_cf91;
    };
    static fm35(p0, p1, p2, globalState) {
      let _source0 = _module.D17.create_DC47(false, new BigNumber(897));
      if (_source0.is_DC47) {
        let _23___mcc_h0 = (_source0).cf81;
        let _24___mcc_h1 = (_source0).cf82;
        let _25_cf82 = _24___mcc_h1;
        let _26_cf81 = _23___mcc_h0;
        return (_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_26_cf81,new BigNumber((_dafny.Seq.of(_26_cf81)).length)), _dafny.Map.Empty.slice().updateUnsafe(_26_cf81,_25_cf82))).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_26_cf81,_25_cf82))));
      } else {
        let _27___mcc_h2 = (_source0).cf80;
        let _28_cf80 = _27___mcc_h2;
        return (_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((function () {
          let _coll7 = new _dafny.Map();
          for (const _compr_7 of (_dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll8 = new _dafny.Set();
            for (const _compr_8 of _dafny.IntegerRange(new BigNumber(869), new BigNumber(577))) {
              let _29_v1 = _compr_8;
              if (((new BigNumber(869)).isLessThanOrEqualTo(_29_v1)) && ((_29_v1).isLessThan(new BigNumber(577)))) {
                _coll8.add(_module.__default.safeModuloInt(_29_v1, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(569))).length)));
              }
            }
            return _coll8;
          }(),new BigNumber(89))).Keys.Elements) {
            let _30_v0 = _compr_7;
            if ((_dafny.Map.Empty.slice().updateUnsafe(function () {
              let _coll9 = new _dafny.Set();
              for (const _compr_9 of _dafny.IntegerRange(new BigNumber(869), new BigNumber(577))) {
                let _31_v1 = _compr_9;
                if (((new BigNumber(869)).isLessThanOrEqualTo(_31_v1)) && ((_31_v1).isLessThan(new BigNumber(577)))) {
                  _coll9.add(_module.__default.safeModuloInt(_31_v1, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(569))).length)));
                }
              }
              return _coll9;
            }(),new BigNumber(89))).contains(_30_v0)) {
              _coll7.push([_30_v0,!(false)]);
            }
          }
          return _coll7;
        }()).length)))).Intersect(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length))));
      }
    };
    static fm36(globalState) {
      return (_dafny.Set.fromElements(true)).Intersect((_dafny.Set.fromElements(true)).Union(_dafny.Set.fromElements(false, false, true)));
    };
    static m0(p0, globalState) {
      let _32_v0;
      _32_v0 = _module.D2.create_DC6(p0);
      let _33_v1;
      _33_v1 = _module.D2.create_DC8(_32_v0);
      let _34_v2;
      _34_v2 = _module.D2.create_DC8(_33_v1);
      let _35_v3;
      _35_v3 = _module.D2.create_DC8(_33_v1);
      (globalState).f16 = function (_source1) {
        if (_source1.is_DC7) {
          let _36___mcc_h0 = (_source1).cf9;
          let _37___mcc_h1 = (_source1).cf10;
          let _38_cf10 = _37___mcc_h1;
          let _39_cf9 = _36___mcc_h0;
          return (_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.of(false), _dafny.Seq.of(false)))).IsProperSubsetOf(_dafny.MultiSet.fromElements(_dafny.Seq.of(!(false))));
        } else if (_source1.is_DC6) {
          let _40___mcc_h2 = (_source1).cf8;
          let _41_cf8 = _40___mcc_h2;
          return false;
        } else {
          let _42___mcc_h3 = (_source1).cf11;
          let _43_cf11 = _42___mcc_h3;
          return false;
        }
      }(_35_v3);
      let _44_v4;
      _44_v4 = _dafny.Seq.UnicodeFromString("m");
      let _45_v5;
      _45_v5 = true;
      let _46_v6;
      _46_v6 = _module.D3.create_DC10(_44_v4, p0, _45_v5);
      let _47_v7;
      _47_v7 = _dafny.Seq.of(_module.D3.create_DC10(_44_v4, p0, _45_v5), _46_v6, _46_v6, _46_v6, _46_v6);
      let _48_i0;
      _48_i0 = _dafny.ZERO;
      L0: {
        while (!_dafny.areEqual(_47_v7, _dafny.Seq.of(_module.D3.create_DC10(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-571)), function (_112_i1) {
  return new _dafny.CodePoint('m'.codePointAt(0));
}), p0, _45_v5)))) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_48_i0)) {
              break L0;
            }
            _48_i0 = (_48_i0).plus(_dafny.ONE);
            let _49_v8;
            _49_v8 = new _dafny.CodePoint('k'.codePointAt(0));
            let _50_v9;
            _50_v9 = _dafny.MultiSet.fromElements(_49_v8);
            let _51_v10;
            _51_v10 = _module.D14.create_DC40(_45_v5, _45_v5, new BigNumber((_50_v9).cardinality()));
            let _52_v11;
            _52_v11 = _module.D14.create_DC42(_51_v10);
            let _53_v12;
            _53_v12 = _module.D14.create_DC42(_52_v11);
            let _source2 = _53_v12;
            if (_source2.is_DC40) {
              let _54___mcc_h4 = (_source2).cf69;
              let _55___mcc_h5 = (_source2).cf70;
              let _56___mcc_h6 = (_source2).cf71;
              let _57_cf71 = _56___mcc_h6;
              let _58_cf70 = _55___mcc_h5;
              let _59_cf69 = _54___mcc_h4;
              _49_v8 = _49_v8;
              let _60_v13;
              let _nw0 = new _module.C0();
              _nw0.__ctor(true);
              _60_v13 = _nw0;
              let _61_v14;
              let _nw1 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
              _61_v14 = _nw1;
              let _index0 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_61_v14).length));
              (_61_v14)[_index0] = p0;
              let _62_v15;
              _62_v15 = _dafny.MultiSet.fromElements(p0);
              let _index1 = _module.__default.safeIndex(new BigNumber(837), new BigNumber((_61_v14).length));
              (_61_v14)[_index1] = (((_dafny.ZERO).minus(new BigNumber((_62_v15).cardinality()))).minus(p0)).plus((_dafny.ZERO).minus(p0));
              let _63_v16;
              _63_v16 = _dafny.Map.Empty.slice().updateUnsafe(!(_58_cf70),_45_v5);
              (globalState).f28 = (_57_cf71).isLessThanOrEqualTo((_dafny.ZERO).minus(new BigNumber((_63_v16).length)));
            } else if (_source2.is_DC41) {
              let _64___mcc_h7 = (_source2).cf72;
              let _65___mcc_h8 = (_source2).cf73;
              let _66_cf73 = _65___mcc_h8;
              let _67_cf72 = _64___mcc_h7;
              let _68_v17;
              let _init0 = ((_69_v4) => function (_70_i2) {
                return _69_v4;
              })(_44_v4);
              let _nw2 = Array((new BigNumber(8)).toNumber());
              for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw2.length); _i0_0++) {
                _nw2[_i0_0] = _init0(new BigNumber(_i0_0));
              }
              _68_v17 = _nw2;
              let _71_v18;
              _71_v18 = _dafny.MultiSet.fromElements(_44_v4);
              let _72_v19;
              _72_v19 = _dafny.Map.Empty.slice().updateUnsafe(_68_v17,_71_v18);
              _72_v19 = (_72_v19).update(_68_v17, (_71_v18).update(_44_v4, _module.__default.abs(new BigNumber(-699))));
              _66_cf73 = new BigNumber(312);
              let _73_v20;
              let _nw3 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
              _73_v20 = _nw3;
              let _index2 = _module.__default.safeIndex(new BigNumber(636), new BigNumber((_73_v20).length));
              (_73_v20)[_index2] = new BigNumber(-670);
              let _index3 = _module.__default.safeIndex(new BigNumber(636), new BigNumber((_73_v20).length));
              (_73_v20)[_index3] = _66_cf73;
              let _74_v21;
              _74_v21 = _dafny.Map.Empty.slice().updateUnsafe(_45_v5,!(_45_v5) || (_45_v5));
              _74_v21 = (_74_v21).update(_45_v5, _45_v5);
            } else if (_source2.is_DC39) {
              let _75___mcc_h9 = (_source2).cf68;
              let _76_cf68 = _75___mcc_h9;
              _44_v4 = _dafny.Seq.Concat(_44_v4, _44_v4);
              let _77_v22;
              _77_v22 = _dafny.Map.Empty.slice().updateUnsafe(p0,_45_v5);
              let _78_v23;
              _78_v23 = _dafny.Map.Empty.slice().updateUnsafe(p0,_77_v22);
              let _79_v24;
              _79_v24 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
              let _80_v25;
              _80_v25 = _dafny.Seq.of(_44_v4);
              let _81_v26;
              let _nw4 = new _module.C3();
              _nw4.__ctor(_76_cf68, _78_v23, _module.__default.fm26(_45_v5, _45_v5, _45_v5, (((_79_v24).contains(p0)) ? ((_79_v24).get(p0)) : (new BigNumber(((_80_v25)[_module.__default.safeIndex(p0, new BigNumber((_80_v25).length))]).length))), globalState));
              _81_v26 = _nw4;
              let _82_v27;
              _82_v27 = _dafny.Set.fromElements(p0);
              let _83_v28;
              _83_v28 = _dafny.Map.Empty.slice().updateUnsafe(_45_v5,p0);
              let _84_v29;
              _84_v29 = _dafny.Map.Empty.slice().updateUnsafe(_81_v26,(new BigNumber((_82_v27).length)).minus(new BigNumber((_83_v28).length)));
              _84_v29 = (_84_v29).update(_81_v26, p0);
              let _85_v30;
              _85_v30 = _module.D0.create_DC1(_49_v8, _45_v5, _45_v5);
              _85_v30 = _85_v30;
              let _86_v31;
              let _nw5 = Array((_dafny.ONE).toNumber());
              _nw5[(_dafny.ZERO).toNumber()] = _45_v5;
              _86_v31 = _nw5;
              let _87_v32;
              _87_v32 = _dafny.Map.Empty.slice().updateUnsafe(_86_v31,p0);
              let _88_v33;
              _88_v33 = _dafny.Map.Empty.slice().updateUnsafe(_86_v31,(p0).isLessThan(new BigNumber((_87_v32).length)));
              _88_v33 = (_88_v33).update(_86_v31, (_module.__default.fm35(p0, _44_v4, (_dafny.ZERO).minus(p0), globalState)).contains(_83_v28));
            } else {
              let _89___mcc_h10 = (_source2).cf74;
              let _90_cf74 = _89___mcc_h10;
              let _91_v34;
              _91_v34 = _dafny.Set.fromElements(_45_v5, _45_v5, _45_v5, false, _45_v5);
              let _92_v35;
              _92_v35 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_91_v34).length),_45_v5);
              let _93_v36;
              _93_v36 = _dafny.Seq.of(p0, p0, p0, p0, p0);
              let _94_v37;
              _94_v37 = _module.D16.create_DC44(_93_v36);
              let _95_v38;
              let _nw6 = Array((new BigNumber(23)).toNumber());
              _nw6[(_dafny.ZERO).toNumber()] = new BigNumber(-58);
              _nw6[(_dafny.ONE).toNumber()] = p0;
              _nw6[(new BigNumber(2)).toNumber()] = new BigNumber((_module.__default.fm36(globalState)).length);
              _nw6[(new BigNumber(3)).toNumber()] = p0;
              _nw6[(new BigNumber(4)).toNumber()] = p0;
              _nw6[(new BigNumber(5)).toNumber()] = p0;
              _nw6[(new BigNumber(6)).toNumber()] = new BigNumber((_92_v35).length);
              _nw6[(new BigNumber(7)).toNumber()] = p0;
              _nw6[(new BigNumber(8)).toNumber()] = new BigNumber(((_94_v37).dtor_cf76).length);
              _nw6[(new BigNumber(9)).toNumber()] = p0;
              _nw6[(new BigNumber(10)).toNumber()] = new BigNumber((_44_v4).length);
              _nw6[(new BigNumber(11)).toNumber()] = p0;
              _nw6[(new BigNumber(12)).toNumber()] = p0;
              _nw6[(new BigNumber(13)).toNumber()] = p0;
              _nw6[(new BigNumber(14)).toNumber()] = new BigNumber((_44_v4).length);
              _nw6[(new BigNumber(15)).toNumber()] = p0;
              _nw6[(new BigNumber(16)).toNumber()] = p0;
              _nw6[(new BigNumber(17)).toNumber()] = new BigNumber((_93_v36).length);
              _nw6[(new BigNumber(18)).toNumber()] = p0;
              _nw6[(new BigNumber(19)).toNumber()] = p0;
              _nw6[(new BigNumber(20)).toNumber()] = p0;
              _nw6[(new BigNumber(21)).toNumber()] = p0;
              _nw6[(new BigNumber(22)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(635)), function (_96_i3) {
                return new _dafny.CodePoint('g'.codePointAt(0));
              })).length));
              _95_v38 = _nw6;
              let _97_v39;
              _97_v39 = _95_v38;
              (globalState).f24 = (_97_v39);
              let _98_v40;
              _98_v40 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_93_v36, _93_v36),false);
              _98_v40 = (_98_v40).update(_93_v36, (p0).isLessThan(new BigNumber(29)));
              let _99_v41;
              let _nw7 = new _module.C0();
              _nw7.__ctor(!(_45_v5));
              _99_v41 = _nw7;
              let _100_v42;
              _100_v42 = _dafny.Map.Empty.slice().updateUnsafe(_45_v5,_99_v41);
              let _101_v43;
              _101_v43 = _dafny.Set.fromElements(p0);
              let _102_v44;
              _102_v44 = _module.D12.create_DC32(true, (_99_v41).f31, _99_v41, _101_v43, p0);
              let _103_v45;
              _103_v45 = _dafny.Map.Empty.slice().updateUnsafe(p0,_99_v41);
              let _104_v46;
              let _nw8 = Array((new BigNumber(26)).toNumber());
              _nw8[(_dafny.ZERO).toNumber()] = _99_v41;
              _nw8[(_dafny.ONE).toNumber()] = _99_v41;
              _nw8[(new BigNumber(2)).toNumber()] = _99_v41;
              _nw8[(new BigNumber(3)).toNumber()] = _99_v41;
              _nw8[(new BigNumber(4)).toNumber()] = _99_v41;
              _nw8[(new BigNumber(5)).toNumber()] = _99_v41;
              _nw8[(new BigNumber(6)).toNumber()] = _99_v41;
              _nw8[(new BigNumber(7)).toNumber()] = (((_100_v42).contains((_99_v41).f31)) ? ((_100_v42).get((_99_v41).f31)) : (_99_v41));
              _nw8[(new BigNumber(8)).toNumber()] = _99_v41;
              _nw8[(new BigNumber(9)).toNumber()] = _99_v41;
              _nw8[(new BigNumber(10)).toNumber()] = _99_v41;
              _nw8[(new BigNumber(11)).toNumber()] = _99_v41;
              _nw8[(new BigNumber(12)).toNumber()] = _99_v41;
              _nw8[(new BigNumber(13)).toNumber()] = _99_v41;
              _nw8[(new BigNumber(14)).toNumber()] = _99_v41;
              _nw8[(new BigNumber(15)).toNumber()] = _99_v41;
              _nw8[(new BigNumber(16)).toNumber()] = _99_v41;
              _nw8[(new BigNumber(17)).toNumber()] = _99_v41;
              _nw8[(new BigNumber(18)).toNumber()] = _99_v41;
              _nw8[(new BigNumber(19)).toNumber()] = _99_v41;
              _nw8[(new BigNumber(20)).toNumber()] = (_102_v44).dtor_cf57;
              _nw8[(new BigNumber(21)).toNumber()] = _99_v41;
              _nw8[(new BigNumber(22)).toNumber()] = _99_v41;
              _nw8[(new BigNumber(23)).toNumber()] = (((_103_v45).contains(p0)) ? ((_103_v45).get(p0)) : (_99_v41));
              _nw8[(new BigNumber(24)).toNumber()] = _99_v41;
              _nw8[(new BigNumber(25)).toNumber()] = _99_v41;
              _104_v46 = _nw8;
              let _index4 = _module.__default.safeIndex(new BigNumber(10), new BigNumber((_104_v46).length));
              (_104_v46)[_index4] = _99_v41;
              let _index5 = _module.__default.safeIndex(new BigNumber(10), new BigNumber((_104_v46).length));
              (_104_v46)[_index5] = ((!((_99_v41).f31)) ? (_99_v41) : (_99_v41));
              let _105_v47;
              _105_v47 = new BigNumber(-327);
              _105_v47 = new BigNumber(774);
            }
            let _106_v48;
            _106_v48 = new BigNumber(662);
            _106_v48 = _106_v48;
            let _107_v49;
            let _nw9 = new _module.C5();
            _nw9.__ctor(!(!(_45_v5)), (p0).isLessThan(_106_v48));
            _107_v49 = _nw9;
            let _108_v50;
            _108_v50 = _dafny.Map.Empty.slice().updateUnsafe(_45_v5,_106_v48);
            let _109_v51;
            _109_v51 = _dafny.Seq.of(_108_v50);
            let _110_v52;
            _110_v52 = _dafny.MultiSet.fromElements(_108_v50);
            let _111_v53;
            let _nw10 = new _module.C5();
            _nw10.__ctor((_107_v49).f30, (_110_v52).IsSubsetOf(_dafny.MultiSet.FromArray(_109_v51)));
            _111_v53 = _nw10;
          }
        }
      }
      let _113_v54;
      let _nw11 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
      _113_v54 = _nw11;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_113_v54).length))) {
        let _114_i4 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_114_i4)) && ((_114_i4).isLessThan(new BigNumber((_113_v54).length))))) {
          (_113_v54)[(_114_i4)] = (_114_i4).minus(new BigNumber(57));
        }
      }
      let _115_v55;
      let _init1 = function (_116_i5) {
        return new _dafny.CodePoint('r'.codePointAt(0));
      };
      let _nw12 = Array((new BigNumber(8)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw12.length); _i0_1++) {
        _nw12[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _115_v55 = _nw12;
      let _117_v56;
      _117_v56 = new _dafny.CodePoint('m'.codePointAt(0));
      let _index6 = _module.__default.safeIndex(new BigNumber(887), new BigNumber((_115_v55).length));
      (_115_v55)[_index6] = _117_v56;
      let _118_v57;
      _118_v57 = _dafny.MultiSet.fromElements(_45_v5);
      let _index7 = _module.__default.safeIndex(new BigNumber(887), new BigNumber((_115_v55).length));
      (_115_v55)[_index7] = _module.__default.fm5(_45_v5, (new BigNumber((_118_v57).cardinality())).multipliedBy(new BigNumber(479)), globalState);
      (globalState).f28 = !_dafny.areEqual(_44_v4, _44_v4);
      (globalState).f16 = (((_45_v5) ? (_45_v5) : (_45_v5))) === (_45_v5);
      return;
    }
    static Main(__noArgsParameter) {
      let _119_v0;
      _119_v0 = new _dafny.CodePoint('v'.codePointAt(0));
      let _120_v1;
      let _nw13 = Array((new BigNumber(3)).toNumber());
      _nw13[(_dafny.ZERO).toNumber()] = _119_v0;
      _nw13[(_dafny.ONE).toNumber()] = _119_v0;
      _nw13[(new BigNumber(2)).toNumber()] = _119_v0;
      _120_v1 = _nw13;
      let _121_v2;
      _121_v2 = _dafny.Seq.UnicodeFromString("mau");
      let _122_v3;
      _122_v3 = false;
      let _123_v4;
      _123_v4 = _dafny.Map.Empty.slice().updateUnsafe(_121_v2,_122_v3);
      let _124_v5;
      let _nw14 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
      _124_v5 = _nw14;
      let _125_v6;
      _125_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-850),_124_v5);
      let _126_v7;
      _126_v7 = new BigNumber(159);
      let _127_v8;
      _127_v8 = _dafny.Seq.of((((_125_v6).contains(_126_v7)) ? ((_125_v6).get(_126_v7)) : (_124_v5)));
      let _128_v9;
      _128_v9 = _dafny.Map.Empty.slice().updateUnsafe(_122_v3,_126_v7);
      let _129_v10;
      _129_v10 = _dafny.Seq.of(!(_122_v3), _122_v3);
      let _130_v11;
      _130_v11 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_126_v7, (((_128_v9).contains(_122_v3)) ? ((_128_v9).get(_122_v3)) : (new BigNumber((_128_v9).length)))),new BigNumber((_dafny.MultiSet.FromArray(_129_v10)).cardinality()));
      let _131_v12;
      let _init2 = ((_132_v10) => function (_133_i0) {
        return _132_v10;
      })(_129_v10);
      let _nw15 = Array((new BigNumber(6)).toNumber());
      for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw15.length); _i0_2++) {
        _nw15[_i0_2] = _init2(new BigNumber(_i0_2));
      }
      _131_v12 = _nw15;
      let _134_v13;
      _134_v13 = _module.D0.create_DC0(_131_v12);
      let _135_v14;
      let _nw16 = Array((new BigNumber(12)).toNumber());
      _nw16[(_dafny.ZERO).toNumber()] = _122_v3;
      _nw16[(_dafny.ONE).toNumber()] = !(_122_v3);
      _nw16[(new BigNumber(2)).toNumber()] = false;
      _nw16[(new BigNumber(3)).toNumber()] = _122_v3;
      _nw16[(new BigNumber(4)).toNumber()] = _122_v3;
      _nw16[(new BigNumber(5)).toNumber()] = _122_v3;
      _nw16[(new BigNumber(6)).toNumber()] = false;
      _nw16[(new BigNumber(7)).toNumber()] = false;
      _nw16[(new BigNumber(8)).toNumber()] = _122_v3;
      _nw16[(new BigNumber(9)).toNumber()] = _122_v3;
      _nw16[(new BigNumber(10)).toNumber()] = _122_v3;
      _nw16[(new BigNumber(11)).toNumber()] = _122_v3;
      _135_v14 = _nw16;
      let _136_v15;
      _136_v15 = _124_v5;
      let _137_v16;
      _137_v16 = _dafny.Map.Empty.slice().updateUnsafe(_122_v3,_dafny.Set.fromElements(_122_v3, _122_v3, _122_v3, false, _122_v3));
      let _138_globalState;
      let _nw17 = new _module.GlobalState();
      _nw17.__ctor(_120_v1, false, false, _123_v4, false, true, _121_v2, new BigNumber(267), false, true, true, _127_v8, _130_v11, new BigNumber(175), true, new BigNumber(87), false, false, (_134_v13).dtor_cf0, true, _135_v14, (_136_v15), false, (_137_v16).Merge(_137_v16), _124_v5, new BigNumber(21), new BigNumber(-144), _135_v14, true);
      _138_globalState = _nw17;
      let _139_v17;
      _139_v17 = _module.D0.create_DC3(_119_v0);
      let _140_v18;
      _140_v18 = _dafny.Set.fromElements(_139_v17);
      let _141_v19;
      _141_v19 = _dafny.Seq.of(_126_v7, _126_v7, (_126_v7).multipliedBy(new BigNumber((_140_v18).length)), _126_v7);
      _126_v7 = (_141_v19)[_module.__default.safeIndex(_126_v7, new BigNumber((_141_v19).length))];
      let _hi0 = (_126_v7).multipliedBy(new BigNumber((_121_v2).length));
      for (let _142_i1 = _module.__default.safeModuloInt(_126_v7, _126_v7); _142_i1.isLessThan(_hi0); _142_i1 = _142_i1.plus(_dafny.ONE)) {
        let _143_v20;
        _143_v20 = _dafny.Set.fromElements(_122_v3, true);
        let _144_v21;
        _144_v21 = _dafny.Map.Empty.slice().updateUnsafe(_126_v7,new BigNumber((_143_v20).length));
        _module.__default.m0((((_144_v21).contains(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_135_v14,_142_i1)).length))) ? ((_144_v21).get(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_135_v14,_142_i1)).length))) : (_126_v7)), _138_globalState);
        let _145_v22;
        _145_v22 = _module.D0.create_DC2(_122_v3);
        let _pat_let_tv0 = _129_v10;
        let _pat_let_tv1 = _129_v10;
        let _rhs0 = function (_pat_let0_0) {
          return function (_146_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_147_dt__update_hcf4_h0) {
                return _module.D0.create_DC2(_147_dt__update_hcf4_h0);
              }(_pat_let1_0);
            }(_dafny.Seq.IsProperPrefixOf(_pat_let_tv0, _pat_let_tv1));
          }(_pat_let0_0);
        }(_145_v22);
        let _rhs1 = _122_v3;
        let _rhs2 = _135_v14;
        let _lhs0 = _138_globalState;
        let _lhs1 = _138_globalState;
        _145_v22 = _rhs0;
        _lhs0.f16 = _rhs1;
        _lhs1.f20 = _rhs2;
        let _source3 = ((_122_v3) ? (_136_v15) : (_124_v5));
        let _148___mcc_h0 = _source3;
        let _149_cf7 = _148___mcc_h0;
        _122_v3 = ((new BigNumber((_128_v9).length)).multipliedBy(_142_i1)).isEqualTo(_126_v7);
        _134_v13 = _134_v13;
        let _index8 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_149_cf7).length));
        (_149_cf7)[_index8] = _126_v7;
        let _index9 = _module.__default.safeIndex(new BigNumber(697), new BigNumber((_149_cf7).length));
        (_149_cf7)[_index9] = _126_v7;
        _122_v3 = !(_122_v3) || (_122_v3);
        let _150_v23;
        _150_v23 = _module.D2.create_DC6(_126_v7);
        (_138_globalState).f28 = _dafny.Seq.IsPrefixOf(_dafny.Seq.update(_127_v8, _module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_121_v2, _module.__default.safeIndex(_126_v7, new BigNumber((_121_v2).length)), new _dafny.CodePoint('q'.codePointAt(0))), _module.__default.safeIndex((_150_v23).dtor_cf8, new BigNumber((_dafny.Seq.update(_121_v2, _module.__default.safeIndex(_126_v7, new BigNumber((_121_v2).length)), new _dafny.CodePoint('q'.codePointAt(0)))).length)), _119_v0)).length), new BigNumber((_127_v8).length)), _124_v5), _dafny.Seq.Concat(_127_v8, _127_v8));
      }
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_135_v14).length))) {
        let _151_i2 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_151_i2)) && ((_151_i2).isLessThan(new BigNumber((_135_v14).length))))) {
          (_135_v14)[(_151_i2)] = false;
        }
      }
      let _152_v24;
      _152_v24 = _dafny.Map.Empty.slice().updateUnsafe(_126_v7,_126_v7);
      let _153_v25;
      _153_v25 = _dafny.Map.Empty.slice().updateUnsafe(_126_v7,(((_152_v24).contains(_126_v7)) ? ((_152_v24).get(_126_v7)) : (_126_v7)));
      (_138_globalState).f16 = !((_153_v25).equals(_153_v25));
      if (false) {
        let _154_v26;
        _154_v26 = _dafny.Map.Empty.slice().updateUnsafe(_122_v3,_122_v3);
        _126_v7 = new BigNumber((_154_v26).length);
        let _index10 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_124_v5).length));
        (_124_v5)[_index10] = _126_v7;
        let _index11 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_124_v5).length));
        let _rhs3 = !(_122_v3) || (_122_v3);
        let _rhs4 = (_dafny.ZERO).minus(_126_v7);
        let _rhs5 = _129_v10;
        let _rhs6 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_126_v7));
        let _lhs2 = _124_v5;
        let _lhs3 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_124_v5).length));
        _122_v3 = _rhs3;
        _126_v7 = _rhs4;
        _129_v10 = _rhs5;
        _lhs2[_lhs3] = _rhs6;
        let _155_v27;
        _155_v27 = _dafny.Map.Empty.slice().updateUnsafe(((_124_v5)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_124_v5).length))]).isLessThanOrEqualTo(_126_v7),new _dafny.CodePoint('u'.codePointAt(0)));
        _155_v27 = (_155_v27).update(_122_v3, _119_v0);
        _126_v7 = (_module.__default.safeDivisionInt(new BigNumber((_module.__default.fm0(_dafny.MultiSet.FromArray(_dafny.Seq.of(false, !(_122_v3))), _138_globalState)).length), new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()))).plus(((_122_v3) ? (new BigNumber(-660)) : (_126_v7)));
        let _156_v28;
        let _nw18 = Array((new BigNumber(13)).toNumber()).fill([]);
        _156_v28 = _nw18;
        let _157_v29;
        let _nw19 = Array((new BigNumber(7)).toNumber());
        _nw19[(_dafny.ZERO).toNumber()] = _141_v19;
        _nw19[(_dafny.ONE).toNumber()] = _141_v19;
        _nw19[(new BigNumber(2)).toNumber()] = _141_v19;
        _nw19[(new BigNumber(3)).toNumber()] = _141_v19;
        _nw19[(new BigNumber(4)).toNumber()] = _141_v19;
        _nw19[(new BigNumber(5)).toNumber()] = _141_v19;
        _nw19[(new BigNumber(6)).toNumber()] = _dafny.Seq.of(_126_v7, (_124_v5)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_124_v5).length))]);
        _157_v29 = _nw19;
        let _index12 = _module.__default.safeIndex(new BigNumber(363), new BigNumber((_156_v28).length));
        (_156_v28)[_index12] = _157_v29;
        let _index13 = _module.__default.safeIndex(new BigNumber(363), new BigNumber((_156_v28).length));
        (_156_v28)[_index13] = _157_v29;
      } else {
        _122_v3 = _122_v3;
        let _158_v30;
        _158_v30 = _dafny.Set.fromElements(_141_v19);
        let _159_v31;
        _159_v31 = _dafny.Set.fromElements(new BigNumber((_158_v30).length), _126_v7);
        if (!((_159_v31).IsProperSubsetOf(_dafny.Set.fromElements(_126_v7)))) {
          _121_v2 = _121_v2;
          let _pat_let_tv2 = _119_v0;
          let _160_v32;
          let _nw20 = Array((new BigNumber(26)).toNumber());
          _nw20[(_dafny.ZERO).toNumber()] = _module.D0.create_DC3(_119_v0);
          _nw20[(_dafny.ONE).toNumber()] = _139_v17;
          _nw20[(new BigNumber(2)).toNumber()] = function (_pat_let2_0) {
            return function (_161_dt__update__tmp_h1) {
              return function (_pat_let3_0) {
                return function (_162_dt__update_hcf5_h0) {
                  return _module.D0.create_DC3(_162_dt__update_hcf5_h0);
                }(_pat_let3_0);
              }(_pat_let_tv2);
            }(_pat_let2_0);
          }(_139_v17);
          _nw20[(new BigNumber(3)).toNumber()] = _139_v17;
          _nw20[(new BigNumber(4)).toNumber()] = ((_122_v3) ? (_139_v17) : (_module.D0.create_DC3(new _dafny.CodePoint('n'.codePointAt(0)))));
          _nw20[(new BigNumber(5)).toNumber()] = _139_v17;
          _nw20[(new BigNumber(6)).toNumber()] = _139_v17;
          _nw20[(new BigNumber(7)).toNumber()] = _139_v17;
          _nw20[(new BigNumber(8)).toNumber()] = _module.D0.create_DC3(new _dafny.CodePoint('p'.codePointAt(0)));
          _nw20[(new BigNumber(9)).toNumber()] = _139_v17;
          _nw20[(new BigNumber(10)).toNumber()] = _139_v17;
          _nw20[(new BigNumber(11)).toNumber()] = _139_v17;
          _nw20[(new BigNumber(12)).toNumber()] = _139_v17;
          _nw20[(new BigNumber(13)).toNumber()] = _139_v17;
          _nw20[(new BigNumber(14)).toNumber()] = _139_v17;
          _nw20[(new BigNumber(15)).toNumber()] = _139_v17;
          _nw20[(new BigNumber(16)).toNumber()] = _139_v17;
          _nw20[(new BigNumber(17)).toNumber()] = _139_v17;
          _nw20[(new BigNumber(18)).toNumber()] = _139_v17;
          _nw20[(new BigNumber(19)).toNumber()] = _139_v17;
          _nw20[(new BigNumber(20)).toNumber()] = _module.D0.create_DC3(_119_v0);
          _nw20[(new BigNumber(21)).toNumber()] = _139_v17;
          _nw20[(new BigNumber(22)).toNumber()] = _139_v17;
          _nw20[(new BigNumber(23)).toNumber()] = _139_v17;
          _nw20[(new BigNumber(24)).toNumber()] = _139_v17;
          _nw20[(new BigNumber(25)).toNumber()] = _139_v17;
          _160_v32 = _nw20;
          _160_v32 = _160_v32;
          let _163_v33;
          _163_v33 = _dafny.MultiSet.fromElements(_126_v7, (_126_v7).minus(_126_v7));
          let _rhs7 = _122_v3;
          let _rhs8 = !(_122_v3) || (_122_v3);
          let _rhs9 = _dafny.MultiSet.fromElements(new BigNumber(-665));
          let _lhs4 = _138_globalState;
          let _lhs5 = _138_globalState;
          _lhs4.f16 = _rhs7;
          _lhs5.f28 = _rhs8;
          _163_v33 = _rhs9;
          _126_v7 = _126_v7;
          let _164_v34;
          _164_v34 = _dafny.Map.Empty.slice().updateUnsafe(_122_v3,new _dafny.CodePoint('g'.codePointAt(0)));
          let _165_v35;
          _165_v35 = _dafny.MultiSet.fromElements(_module.D0.create_DC1((((_164_v34).contains(_122_v3)) ? ((_164_v34).get(_122_v3)) : (_119_v0)), _122_v3, false));
          let _166_v36;
          _166_v36 = _dafny.Seq.of(_165_v35);
          let _167_v37;
          _167_v37 = _module.D0.create_DC1(_119_v0, _122_v3, _122_v3);
          let _index14 = _module.__default.safeIndex(new BigNumber(33), new BigNumber((_135_v14).length));
          (_135_v14)[_index14] = (_dafny.MultiSet.FromArray(_dafny.Seq.of(_167_v37))).IsSubsetOf((_166_v36)[_module.__default.safeIndex(_126_v7, new BigNumber((_166_v36).length))]);
          let _index15 = _module.__default.safeIndex(new BigNumber(33), new BigNumber((_135_v14).length));
          (_135_v14)[_index15] = (_module.__default.safeDivisionInt(_126_v7, _126_v7)).isLessThanOrEqualTo(_126_v7);
        } else {
          let _168_v38;
          let _nw21 = new _module.C2();
          _nw21.__ctor(_129_v10, true);
          _168_v38 = _nw21;
          let _169_v39;
          _169_v39 = _dafny.Map.Empty.slice().updateUnsafe(_122_v3,_122_v3);
          let _170_v40;
          let _nw22 = Array((new BigNumber(26)).toNumber());
          _nw22[(_dafny.ZERO).toNumber()] = _169_v39;
          _nw22[(_dafny.ONE).toNumber()] = _169_v39;
          _nw22[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(!(_122_v3),_122_v3);
          _nw22[(new BigNumber(3)).toNumber()] = _module.__default.fm30(_119_v0, _139_v17, _138_globalState);
          _nw22[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_122_v3,!(_122_v3));
          _nw22[(new BigNumber(5)).toNumber()] = _169_v39;
          _nw22[(new BigNumber(6)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_122_v3,_122_v3);
          _nw22[(new BigNumber(7)).toNumber()] = _169_v39;
          _nw22[(new BigNumber(8)).toNumber()] = _169_v39;
          _nw22[(new BigNumber(9)).toNumber()] = _169_v39;
          _nw22[(new BigNumber(10)).toNumber()] = _169_v39;
          _nw22[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_122_v3,_122_v3);
          _nw22[(new BigNumber(12)).toNumber()] = (_169_v39).update(_122_v3, _122_v3);
          _nw22[(new BigNumber(13)).toNumber()] = _169_v39;
          _nw22[(new BigNumber(14)).toNumber()] = _169_v39;
          _nw22[(new BigNumber(15)).toNumber()] = _169_v39;
          _nw22[(new BigNumber(16)).toNumber()] = _169_v39;
          _nw22[(new BigNumber(17)).toNumber()] = _169_v39;
          _nw22[(new BigNumber(18)).toNumber()] = _169_v39;
          _nw22[(new BigNumber(19)).toNumber()] = _169_v39;
          _nw22[(new BigNumber(20)).toNumber()] = _169_v39;
          _nw22[(new BigNumber(21)).toNumber()] = _169_v39;
          _nw22[(new BigNumber(22)).toNumber()] = _169_v39;
          _nw22[(new BigNumber(23)).toNumber()] = _169_v39;
          _nw22[(new BigNumber(24)).toNumber()] = _169_v39;
          _nw22[(new BigNumber(25)).toNumber()] = _169_v39;
          _170_v40 = _nw22;
          let _171_v41;
          _171_v41 = _dafny.Map.Empty.slice().updateUnsafe(_126_v7,_122_v3);
          let _172_v42;
          let _nw23 = new _module.C3();
          _nw23.__ctor(_170_v40, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-190),_171_v41), _122_v3);
          _172_v42 = _nw23;
          _172_v42 = _172_v42;
          let _index16 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_124_v5).length));
          (_124_v5)[_index16] = ((_122_v3) ? (_126_v7) : (new BigNumber((_121_v2).length)));
          let _index17 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_124_v5).length));
          (_124_v5)[_index17] = (_168_v38).fm13(_126_v7, _138_globalState);
          (_138_globalState).f28 = false;
          _122_v3 = (_172_v42).f31;
        }
        let _173_v43;
        _173_v43 = _dafny.Seq.of(_129_v10);
        let _174_v44;
        _174_v44 = _dafny.Seq.of(_129_v10, (_173_v43)[_module.__default.safeIndex(_126_v7, new BigNumber((_173_v43).length))], _129_v10, _dafny.Seq.of(_122_v3));
        _126_v7 = _module.__default.safeModuloInt(new BigNumber((_174_v44).length), _126_v7);
        if (_122_v3) {
          _128_v9 = (_128_v9).update((_122_v3) && (_122_v3), new BigNumber(62));
          _module.__default.m0(_126_v7, _138_globalState);
          _126_v7 = _module.__default.safeDivisionInt(_126_v7, _126_v7);
          _module.__default.m0(_126_v7, _138_globalState);
          let _175_v45;
          let _nw24 = Array((_dafny.ONE).toNumber());
          _nw24[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_121_v2, _module.__default.safeIndex(_126_v7, new BigNumber((_121_v2).length)), (_121_v2)[_module.__default.safeIndex(new BigNumber((_121_v2).length), new BigNumber((_121_v2).length))]);
          _175_v45 = _nw24;
          _175_v45 = _175_v45;
        } else {
          _126_v7 = _126_v7;
          _139_v17 = _module.__default.fm31(_126_v7, _138_globalState);
          let _176_v46;
          _176_v46 = _module.D5.create_DC14(_159_v31, _126_v7, true, _135_v14);
          let _177_v47;
          let _nw25 = Array((new BigNumber(7)).toNumber());
          _nw25[(_dafny.ZERO).toNumber()] = _135_v14;
          _nw25[(_dafny.ONE).toNumber()] = _135_v14;
          _nw25[(new BigNumber(2)).toNumber()] = (_176_v46).dtor_cf22;
          _nw25[(new BigNumber(3)).toNumber()] = _135_v14;
          _nw25[(new BigNumber(4)).toNumber()] = _135_v14;
          _nw25[(new BigNumber(5)).toNumber()] = _135_v14;
          _nw25[(new BigNumber(6)).toNumber()] = _135_v14;
          _177_v47 = _nw25;
          let _178_v48;
          let _nw26 = new _module.C1();
          _nw26.__ctor(_126_v7, _177_v47, _122_v3);
          _178_v48 = _nw26;
          let _179_v49;
          _179_v49 = _dafny.Map.Empty.slice().updateUnsafe(_126_v7,_122_v3);
          let _180_v50;
          _180_v50 = _module.D9.create_DC26(_122_v3, _121_v2, _126_v7, _179_v49, _122_v3);
          let _181_v51;
          _181_v51 = _dafny.Map.Empty.slice().updateUnsafe(_180_v50,_dafny.MultiSet.fromElements(new BigNumber(-674), new BigNumber((_141_v19).length)));
          let _182_v52;
          _182_v52 = _dafny.MultiSet.fromElements(_126_v7);
          let _183_v53;
          let _nw27 = new _module.C5();
          _nw27.__ctor(!_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("cf"), _module.__default.fm20(_121_v2, (((_181_v51).contains(_180_v50)) ? ((_181_v51).get(_180_v50)) : (_182_v52)), !(!((_129_v10)[_module.__default.safeIndex(_126_v7, new BigNumber((_129_v10).length))])), new BigNumber((_182_v52).cardinality()), _138_globalState)), _122_v3);
          _183_v53 = _nw27;
          _126_v7 = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_183_v53).f29,_119_v0)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_122_v3,_119_v0))).length);
        }
        let _184_v54;
        _184_v54 = _dafny.MultiSet.fromElements(_126_v7);
        let _index18 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_135_v14).length));
        (_135_v14)[_index18] = (_184_v54).contains(_126_v7);
        let _index19 = _module.__default.safeIndex(new BigNumber(858), new BigNumber((_135_v14).length));
        (_135_v14)[_index19] = _122_v3;
        let _185_v55;
        let _nw28 = new _module.C5();
        _nw28.__ctor(_122_v3, _122_v3);
        _185_v55 = _nw28;
        let _186_v56;
        _186_v56 = _dafny.Map.Empty.slice().updateUnsafe(_122_v3,_185_v55);
        let _index20 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_135_v14).length));
        let _index21 = _module.__default.safeIndex(new BigNumber(858), new BigNumber((_135_v14).length));
        let _rhs10 = ((_126_v7).minus(new BigNumber((_121_v2).length))).isLessThan(_module.__default.safeModuloInt(_126_v7, _126_v7));
        let _rhs11 = new BigNumber((((_186_v56).Merge((_186_v56).update((_185_v55).f30, _185_v55))).Merge((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm10(_119_v0, (_185_v55).f30, _138_globalState),_185_v55)).Merge(_186_v56))).length);
        let _rhs12 = (_185_v55).f29;
        let _rhs13 = _135_v14;
        let _lhs6 = _135_v14;
        let _lhs7 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_135_v14).length));
        let _lhs8 = _135_v14;
        let _lhs9 = _module.__default.safeIndex(new BigNumber(858), new BigNumber((_135_v14).length));
        let _lhs10 = _138_globalState;
        _lhs6[_lhs7] = _rhs10;
        _126_v7 = _rhs11;
        _lhs8[_lhs9] = _rhs12;
        _lhs10.f20 = _rhs13;
      }
      _module.__default.m0(_module.__default.safeDivisionInt(_126_v7, _126_v7), _138_globalState);
      let _187_i3;
      _187_i3 = _dafny.ZERO;
      L1: {
        while (_122_v3) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_187_i3)) {
              break L1;
            }
            _187_i3 = (_187_i3).plus(_dafny.ONE);
            let _188_v57;
            let _nw29 = new _module.C0();
            _nw29.__ctor((_126_v7).isLessThan(_126_v7));
            _188_v57 = _nw29;
            _126_v7 = _126_v7;
            _126_v7 = new BigNumber(796);
            _module.__default.m0(_126_v7, _138_globalState);
          }
        }
      }
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_124_v5).length))) {
        let _189_i4 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_189_i4)) && ((_189_i4).isLessThan(new BigNumber((_124_v5).length))))) {
          (_124_v5)[(_189_i4)] = (_189_i4).multipliedBy(_126_v7);
        }
      }
      let _190_i5;
      _190_i5 = _dafny.ZERO;
      L2: {
        while (_122_v3) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_190_i5)) {
              break L2;
            }
            _190_i5 = (_190_i5).plus(_dafny.ONE);
            let _191_v58;
            _191_v58 = _dafny.MultiSet.fromElements(new BigNumber(49));
            let _192_v59;
            _192_v59 = _dafny.Seq.of(_dafny.MultiSet.fromElements(_126_v7), _191_v58, _191_v58);
            let _193_v60;
            _193_v60 = _module.D12.create_DC31(_129_v10);
            let _pat_let_tv3 = _122_v3;
            let _rhs14 = (_dafny.MultiSet.fromElements(_126_v7)).IsSubsetOf((_192_v59)[_module.__default.safeIndex(_126_v7, new BigNumber((_192_v59).length))]);
            let _rhs15 = _124_v5;
            let _rhs16 = _122_v3;
            let _rhs17 = (function (_pat_let4_0) {
              return function (_194_dt__update__tmp_h2) {
                return function (_pat_let5_0) {
                  return function (_195_dt__update_hcf54_h0) {
                    return _module.D12.create_DC31(_195_dt__update_hcf54_h0);
                  }(_pat_let5_0);
                }(_dafny.Seq.of(!(_pat_let_tv3)));
              }(_pat_let4_0);
            }(_193_v60)).dtor_cf54;
            let _rhs18 = _122_v3;
            let _lhs11 = _138_globalState;
            let _lhs12 = _138_globalState;
            let _lhs13 = _138_globalState;
            _lhs11.f16 = _rhs14;
            _lhs12.f24 = _rhs15;
            _122_v3 = _rhs16;
            _129_v10 = _rhs17;
            _lhs13.f16 = _rhs18;
            _126_v7 = _126_v7;
            let _index22 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_135_v14).length));
            (_135_v14)[_index22] = (_126_v7).isLessThan(_126_v7);
            let _index23 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_135_v14).length));
            let _rhs19 = _dafny.Seq.Concat(_129_v10, _dafny.Seq.Concat(_129_v10, _dafny.Seq.of(_122_v3, true)));
            let _rhs20 = _122_v3;
            let _rhs21 = _122_v3;
            let _lhs14 = _135_v14;
            let _lhs15 = _module.__default.safeIndex(new BigNumber(259), new BigNumber((_135_v14).length));
            let _lhs16 = _138_globalState;
            _129_v10 = _rhs19;
            _lhs14[_lhs15] = _rhs20;
            _lhs16.f28 = _rhs21;
            _module.__default.m0(_126_v7, _138_globalState);
          }
        }
      }
      let _196_v61;
      _196_v61 = _dafny.MultiSet.fromElements(_122_v3);
      let _197_v62;
      _197_v62 = _dafny.Seq.of(_dafny.MultiSet.fromElements(_122_v3, _122_v3, _122_v3, _122_v3), (_196_v61).update(_122_v3, _module.__default.abs(_126_v7)));
      let _198_v63;
      _198_v63 = _module.D5.create_DC13((_197_v62)[_module.__default.safeIndex(_126_v7, new BigNumber((_197_v62).length))]);
      let _source4 = _198_v63;
      if (_source4.is_DC14) {
        let _199___mcc_h1 = (_source4).cf19;
        let _200___mcc_h2 = (_source4).cf20;
        let _201___mcc_h3 = (_source4).cf21;
        let _202___mcc_h4 = (_source4).cf22;
        let _203_cf22 = _202___mcc_h4;
        let _204_cf21 = _201___mcc_h3;
        let _205_cf20 = _200___mcc_h2;
        let _206_cf19 = _199___mcc_h1;
        _205_cf20 = (new BigNumber((_141_v19).length)).multipliedBy(_module.__default.fm7(_204_cf21, _122_v3, _138_globalState));
        let _207_v64;
        let _nw30 = new _module.C2();
        _nw30.__ctor(_dafny.Seq.update(_129_v10, _module.__default.safeIndex(new BigNumber(703), new BigNumber((_129_v10).length)), _122_v3), false);
        _207_v64 = _nw30;
        _207_v64 = _207_v64;
        let _208_v65;
        _208_v65 = _module.D0.create_DC1(_119_v0, !(_204_cf21), _module.__default.fm26(_122_v3, _122_v3, _204_cf21, _205_cf20, _138_globalState));
        let _209_v66;
        _209_v66 = _dafny.Seq.of(_208_v65);
        let _210_v67;
        _210_v67 = _dafny.Map.Empty.slice().updateUnsafe(_119_v0,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_122_v3,new BigNumber((_dafny.Seq.UnicodeFromString("ks")).length))).length));
        let _211_v68;
        _211_v68 = _dafny.Map.Empty.slice().updateUnsafe(_126_v7,_119_v0);
        let _212_v69;
        let _nw31 = Array((new BigNumber(12)).toNumber()).fill([]);
        _212_v69 = _nw31;
        let _213_v70;
        let _nw32 = new _module.C4();
        _nw32.__ctor(_209_v66, _122_v3, (_dafny.ZERO).minus((((_210_v67).contains((((_211_v68).contains(new BigNumber((_121_v2).length))) ? ((_211_v68).get(new BigNumber((_121_v2).length))) : (_119_v0)))) ? ((_210_v67).get((((_211_v68).contains(new BigNumber((_121_v2).length))) ? ((_211_v68).get(new BigNumber((_121_v2).length))) : (_119_v0)))) : (new BigNumber((_121_v2).length)))), _212_v69, _122_v3);
        _213_v70 = _nw32;
        _213_v70 = _213_v70;
        let _214_v71;
        _214_v71 = _dafny.Set.fromElements(_213_v70.f35);
        (_138_globalState).f16 = !((_214_v71).IsSubsetOf(_214_v71)) || (_204_cf21);
      } else if (_source4.is_DC13) {
        let _215___mcc_h5 = (_source4).cf18;
        let _216_cf18 = _215___mcc_h5;
        _124_v5 = _124_v5;
        let _217_v72;
        let _nw33 = Array((new BigNumber(5)).toNumber());
        _217_v72 = _nw33;
        let _218_v73;
        let _nw34 = Array((new BigNumber(22)).toNumber()).fill([]);
        _218_v73 = _nw34;
        let _219_v74;
        let _nw35 = new _module.C1();
        _nw35.__ctor(_126_v7, _218_v73, _122_v3);
        _219_v74 = _nw35;
        let _220_v75;
        _220_v75 = _dafny.Seq.of(_219_v74, _219_v74);
        let _index24 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_217_v72).length));
        (_217_v72)[_index24] = (_220_v75)[_module.__default.safeIndex((_141_v19)[_module.__default.safeIndex(_126_v7, new BigNumber((_141_v19).length))], new BigNumber((_220_v75).length))];
        let _index25 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_217_v72).length));
        (_217_v72)[_index25] = _219_v74;
        _129_v10 = _129_v10;
        let _index26 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_218_v73).length));
        (_218_v73)[_index26] = _135_v14;
        let _index27 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_218_v73).length));
        (_218_v73)[_index27] = _135_v14;
      } else {
        let _221___mcc_h6 = (_source4).cf23;
        let _222_cf23 = _221___mcc_h6;
        let _223_v76;
        _223_v76 = _module.D0.create_DC2(_122_v3);
        let _224_v77;
        _224_v77 = _module.D0.create_DC1(_119_v0, (_223_v76).dtor_cf4, _122_v3);
        let _225_v78;
        _225_v78 = _dafny.Seq.of(_224_v77);
        let _226_v79;
        _226_v79 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_126_v7),_dafny.Seq.Create(_module.__default.abs(new BigNumber(106)), ((_227_v19, _228_v3, _229_v7, _230_v0) => function (_231_i6) {
          return new BigNumber((_dafny.Seq.of((_227_v19)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC2(_228_v3),_229_v7)).length), new BigNumber((_227_v19).length))], new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(971)), ((_232_v0) => function (_233_i7) {
            return _232_v0;
          })(_230_v0))).length))).length);
        })(_141_v19, _122_v3, _126_v7, _119_v0)));
        let _234_v80;
        _234_v80 = _dafny.Seq.of(_141_v19);
        let _235_v81;
        let _nw36 = Array((new BigNumber(25)).toNumber()).fill([]);
        _235_v81 = _nw36;
        let _236_v82;
        let _nw37 = new _module.C4();
        _nw37.__ctor(_dafny.Seq.Concat(_225_v78, _225_v78), _122_v3, new BigNumber(((((_226_v79).contains((_234_v80)[_module.__default.safeIndex(_126_v7, new BigNumber((_234_v80).length))])) ? ((_226_v79).get((_234_v80)[_module.__default.safeIndex(_126_v7, new BigNumber((_234_v80).length))])) : (_141_v19))).length), _235_v81, !(false) || (_122_v3));
        _236_v82 = _nw37;
        let _237_v83;
        _237_v83 = _dafny.Set.fromElements(_module.__default.fm1(_119_v0, _138_globalState));
        let _238_v84;
        _238_v84 = _dafny.Map.Empty.slice().updateUnsafe(_122_v3,_236_v82.f35);
        let _rhs22 = _module.__default.fm7(!(_236_v82.f35) || (_122_v3), (((_238_v84).contains(!(_122_v3))) ? ((_238_v84).get(!(_122_v3))) : (_236_v82.f35)), _138_globalState);
        let _rhs23 = ((_122_v3) ? (_236_v82) : (_236_v82));
        let _rhs24 = _236_v82.f35;
        let _rhs25 = _237_v83;
        let _rhs26 = ((_122_v3) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(674)), ((_239_v0) => function (_240_i8) {
          return _239_v0;
        })(_119_v0))) : (_121_v2));
        let _lhs17 = _138_globalState;
        _126_v7 = _rhs22;
        _236_v82 = _rhs23;
        _lhs17.f28 = _rhs24;
        _237_v83 = _rhs25;
        _121_v2 = _rhs26;
        let _index28 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_135_v14).length));
        (_135_v14)[_index28] = !_dafny.areEqual(_121_v2, _dafny.Seq.update(_dafny.Seq.UnicodeFromString("w"), _module.__default.safeIndex(new BigNumber((_module.__default.fm32(_126_v7, _236_v82.f35, _236_v82.f35, _236_v82.f35, _138_globalState)).length), new BigNumber((_dafny.Seq.UnicodeFromString("w")).length)), _119_v0));
        let _index29 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_135_v14).length));
        (_135_v14)[_index29] = _236_v82.f35;
        let _241_v85;
        let _242_v86;
        let _243_v87;
        let _244_v88;
        let _out0;
        let _out1;
        let _out2;
        let _out3;
        let _outcollector0 = (_236_v82).m2((((_135_v14)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_135_v14).length))]) ? (_236_v82.f35) : (_122_v3)), (_196_v61).Union(_196_v61), _196_v61, _138_globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _out2 = _outcollector0[2];
        _out3 = _outcollector0[3];
        _241_v85 = _out0;
        _242_v86 = _out1;
        _243_v87 = _out2;
        _244_v88 = _out3;
        let _245_v89;
        _245_v89 = _module.D3.create_DC10(_dafny.Seq.UnicodeFromString("eavwb"), _241_v85, false);
        let _246_v90;
        _246_v90 = _module.D3.create_DC11(_245_v89);
        (_138_globalState).f16 = (_module.__default.fm33(_244_v88, _246_v90, _121_v2, _138_globalState)).dtor_cf32;
      }
      let _247_v91;
      _247_v91 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_121_v2).length),_122_v3);
      let _248_v92;
      _248_v92 = _dafny.Map.Empty.slice().updateUnsafe((_247_v91).update(_126_v7, _122_v3),_126_v7);
      let _249_v93;
      _249_v93 = _module.D13.create_DC34(_248_v92);
      let _250_v94;
      _250_v94 = _dafny.Map.Empty.slice().updateUnsafe(_122_v3,_135_v14);
      let _251_v95;
      _251_v95 = _dafny.Seq.of((((_250_v94).contains(_122_v3)) ? ((_250_v94).get(_122_v3)) : (_135_v14)));
      let _252_v96;
      _252_v96 = _dafny.MultiSet.fromElements(new BigNumber((_251_v95).length));
      let _253_v97;
      _253_v97 = _dafny.Seq.of(_252_v96);
      let _254_v98;
      _254_v98 = _dafny.MultiSet.fromElements(_252_v96, _252_v96, _252_v96);
      (_138_globalState).f16 = ((_dafny.MultiSet.FromArray(_253_v97)).Difference(_254_v98)).IsProperSubsetOf(_module.__default.fm34((_249_v93).dtor_cf61, _138_globalState));
      let _255_v99;
      _255_v99 = _dafny.Map.Empty.slice().updateUnsafe(_122_v3,_141_v19);
      _module.__default.m0((((_153_v25).contains(new BigNumber(((_255_v99).update(_122_v3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(608)), ((_258_v3) => function (_259_i9) {
        return new BigNumber((_dafny.Seq.of(_module.D13.create_DC36(_258_v3, _258_v3))).length);
      })(_122_v3)))).length))) ? ((_153_v25).get(new BigNumber(((_255_v99).update(_122_v3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(608)), ((_256_v3) => function (_257_i9) {
        return new BigNumber((_dafny.Seq.of(_module.D13.create_DC36(_256_v3, _256_v3))).length);
      })(_122_v3)))).length))) : (_126_v7)), _138_globalState);
      let _hi1 = _126_v7;
      for (let _260_i10 = _126_v7; _260_i10.isLessThan(_hi1); _260_i10 = _260_i10.plus(_dafny.ONE)) {
        let _261_v100;
        let _nw38 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Map.Empty);
        _261_v100 = _nw38;
        let _262_v101;
        _262_v101 = _module.D14.create_DC39(_261_v100);
        let _263_v102;
        _263_v102 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_129_v10)).cardinality()),_247_v91);
        let _264_v103;
        let _nw39 = new _module.C3();
        _nw39.__ctor((_262_v101).dtor_cf68, _263_v102, _122_v3);
        _264_v103 = _nw39;
        let _265_v104;
        _265_v104 = _dafny.Set.fromElements(_264_v103);
        _265_v104 = ((_265_v104).Intersect(_265_v104)).Difference((_265_v104).Difference(_265_v104));
        let _266_v105;
        _266_v105 = _dafny.Set.fromElements(_124_v5);
        let _rhs27 = (_dafny.ZERO).minus(_126_v7);
        let _rhs28 = (_260_i10).plus(_126_v7);
        let _rhs29 = _126_v7;
        let _rhs30 = (_266_v105).Difference(_266_v105);
        _126_v7 = _rhs27;
        _126_v7 = _rhs28;
        _126_v7 = _rhs29;
        _266_v105 = _rhs30;
        let _index30 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_124_v5).length));
        (_124_v5)[_index30] = _126_v7;
        let _267_v106;
        let _nw40 = Array((new BigNumber(6)).toNumber());
        _nw40[(_dafny.ZERO).toNumber()] = _135_v14;
        _nw40[(_dafny.ONE).toNumber()] = _135_v14;
        _nw40[(new BigNumber(2)).toNumber()] = _135_v14;
        _nw40[(new BigNumber(3)).toNumber()] = _135_v14;
        _nw40[(new BigNumber(4)).toNumber()] = _135_v14;
        _nw40[(new BigNumber(5)).toNumber()] = _135_v14;
        _267_v106 = _nw40;
        let _268_v107;
        let _nw41 = new _module.C1();
        _nw41.__ctor(_126_v7, _267_v106, _122_v3);
        _268_v107 = _nw41;
        let _index31 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_124_v5).length));
        (_124_v5)[_index31] = new BigNumber((_dafny.Set.fromElements(_268_v107)).length);
        let _index32 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_124_v5).length));
        (_124_v5)[_index32] = new BigNumber((_121_v2).length);
      }
      let _269_v108;
      let _nw42 = new _module.C5();
      _nw42.__ctor((_129_v10)[_module.__default.safeIndex(_126_v7, new BigNumber((_129_v10).length))], _122_v3);
      _269_v108 = _nw42;
      let _270_v109;
      _270_v109 = _module.D6.create_DC16(_120_v1);
      let _source5 = _270_v109;
      if (_source5.is_DC17) {
        let _271___mcc_h7 = (_source5).cf25;
        let _272___mcc_h8 = (_source5).cf26;
        let _273___mcc_h9 = (_source5).cf27;
        let _274_cf27 = _273___mcc_h9;
        let _275_cf26 = _272___mcc_h8;
        let _276_cf25 = _271___mcc_h7;
        let _277_v110;
        _277_v110 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm10(_module.__default.fm20(_121_v2, _252_v96, (_269_v108).f30, new BigNumber((_196_v61).cardinality()), _138_globalState), (_269_v108).f29, _138_globalState),(_269_v108).f30);
        let _278_v111;
        _278_v111 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm18((((_277_v110).contains(false)) ? ((_277_v110).get(false)) : (_276_cf25)), _122_v3, _138_globalState),_277_v110);
        _278_v111 = (_278_v111).update(_126_v7, _277_v110);
        _274_cf27 = _module.__default.safeDivisionInt((_126_v7).minus(new BigNumber((_129_v10).length)), _274_cf27);
        _module.__default.m0(_126_v7, _138_globalState);
        _252_v96 = (((_269_v108).f30) ? (_dafny.MultiSet.fromElements(new BigNumber((_121_v2).length))) : (_252_v96));
      } else if (_source5.is_DC16) {
        let _279___mcc_h10 = (_source5).cf24;
        let _280_cf24 = _279___mcc_h10;
        let _281_v112;
        _281_v112 = _module.D0.create_DC1(_119_v0, (_269_v108).f30, (_269_v108).f30);
        let _282_v113;
        let _283_v114;
        let _out4;
        let _out5;
        let _outcollector1 = (_269_v108).m1(_141_v19, new BigNumber(-33), _281_v112, _138_globalState);
        _out4 = _outcollector1[0];
        _out5 = _outcollector1[1];
        _282_v113 = _out4;
        _283_v114 = _out5;
        if (_283_v114) {
          _122_v3 = !_dafny.areEqual(_dafny.Seq.UnicodeFromString("ei"), _module.__default.fm0(_196_v61, _138_globalState));
          _121_v2 = _121_v2;
          _module.__default.m0(_126_v7, _138_globalState);
          let _284_v115;
          let _285_v116;
          let _out6;
          let _out7;
          let _outcollector2 = (_269_v108).m1(_141_v19, _126_v7, _281_v112, _138_globalState);
          _out6 = _outcollector2[0];
          _out7 = _outcollector2[1];
          _284_v115 = _out6;
          _285_v116 = _out7;
          _251_v95 = _dafny.Seq.Concat(_251_v95, _dafny.Seq.Concat(_251_v95, _251_v95));
        } else {
          (_138_globalState).f20 = _135_v14;
          let _pat_let_tv4 = _119_v0;
          let _286_v117;
          let _nw43 = Array((new BigNumber(21)).toNumber());
          _nw43[(_dafny.ZERO).toNumber()] = function (_pat_let6_0) {
            return function (_287_dt__update__tmp_h3) {
              return function (_pat_let7_0) {
                return function (_288_dt__update_hcf5_h1) {
                  return _module.D0.create_DC3(_288_dt__update_hcf5_h1);
                }(_pat_let7_0);
              }(_pat_let_tv4);
            }(_pat_let6_0);
          }(_139_v17);
          _nw43[(_dafny.ONE).toNumber()] = _139_v17;
          _nw43[(new BigNumber(2)).toNumber()] = _139_v17;
          _nw43[(new BigNumber(3)).toNumber()] = _module.D0.create_DC3(_119_v0);
          _nw43[(new BigNumber(4)).toNumber()] = _139_v17;
          _nw43[(new BigNumber(5)).toNumber()] = _139_v17;
          _nw43[(new BigNumber(6)).toNumber()] = _139_v17;
          _nw43[(new BigNumber(7)).toNumber()] = _139_v17;
          _nw43[(new BigNumber(8)).toNumber()] = _139_v17;
          _nw43[(new BigNumber(9)).toNumber()] = _139_v17;
          _nw43[(new BigNumber(10)).toNumber()] = _139_v17;
          _nw43[(new BigNumber(11)).toNumber()] = _module.D0.create_DC3(_119_v0);
          _nw43[(new BigNumber(12)).toNumber()] = _139_v17;
          _nw43[(new BigNumber(13)).toNumber()] = _139_v17;
          _nw43[(new BigNumber(14)).toNumber()] = _139_v17;
          _nw43[(new BigNumber(15)).toNumber()] = _139_v17;
          _nw43[(new BigNumber(16)).toNumber()] = _139_v17;
          _nw43[(new BigNumber(17)).toNumber()] = _139_v17;
          _nw43[(new BigNumber(18)).toNumber()] = _139_v17;
          _nw43[(new BigNumber(19)).toNumber()] = _139_v17;
          _nw43[(new BigNumber(20)).toNumber()] = _139_v17;
          _286_v117 = _nw43;
          let _289_v118;
          _289_v118 = _dafny.Seq.of(_286_v117, _286_v117, _286_v117);
          let _290_v119;
          _290_v119 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm7(_282_v113, (_269_v108).f30, _138_globalState),_289_v118);
          _290_v119 = (_290_v119).update(_126_v7, _289_v118);
          let _291_v120;
          _291_v120 = _module.D3.create_DC10(_dafny.Seq.UnicodeFromString("soglnufp"), _126_v7, (_269_v108).f29);
          let _292_v121;
          _292_v121 = _dafny.Map.Empty.slice().updateUnsafe(_291_v120,(_269_v108).f30);
          _292_v121 = _dafny.Map.Empty.slice().updateUnsafe(_291_v120,(_122_v3) || (_283_v114));
          let _293_v122;
          _293_v122 = _module.D9.create_DC25(_126_v7, (_269_v108).f30, _module.__default.fm3(_138_globalState), _283_v114);
          let _294_v123;
          let _nw44 = new _module.C5();
          _nw44.__ctor(!(!(false)), !(false) || ((_293_v122).dtor_cf44));
          _294_v123 = _nw44;
          let _295_v124;
          let _nw45 = new _module.C2();
          _nw45.__ctor(_module.__default.fm2(_282_v113, _126_v7, (_269_v108).f29, (((_247_v91).contains(new BigNumber((_121_v2).length))) ? ((_247_v91).get(new BigNumber((_121_v2).length))) : ((_269_v108).f30)), _138_globalState), (new BigNumber(125)).isEqualTo(_126_v7));
          _295_v124 = _nw45;
          _295_v124 = _295_v124;
        }
        _283_v114 = (_269_v108).f29;
        let _296_v125;
        let _init3 = ((_297_v3, _298_v108) => function (_299_i11) {
          return _dafny.Map.Empty.slice().updateUnsafe(_297_v3,(_298_v108).f30);
        })(_122_v3, _269_v108);
        let _nw46 = Array((new BigNumber(25)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw46.length); _i0_3++) {
          _nw46[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _296_v125 = _nw46;
        let _300_v126;
        _300_v126 = _dafny.Map.Empty.slice().updateUnsafe(_126_v7,_247_v91);
        let _301_v127;
        let _nw47 = new _module.C3();
        _nw47.__ctor(_296_v125, _300_v126, (_269_v108).f30);
        _301_v127 = _nw47;
        let _302_v128;
        _302_v128 = _dafny.MultiSet.fromElements(_301_v127, _301_v127, _301_v127, _301_v127);
        _283_v114 = _module.__default.fm8((_302_v128).IsSubsetOf(_dafny.MultiSet.fromElements(_301_v127, _301_v127)), (_126_v7).minus(_126_v7), _138_globalState);
      } else {
        let _303___mcc_h11 = (_source5).cf28;
        let _304_cf28 = _303___mcc_h11;
        let _305_v129;
        let _nw48 = new _module.C0();
        _nw48.__ctor((_269_v108).f30);
        _305_v129 = _nw48;
        let _306_v130;
        _306_v130 = _305_v129;
        let _307_v131;
        _307_v131 = _dafny.Seq.of(_305_v129, _305_v129);
        let _308_v132;
        let _nw49 = Array((new BigNumber(29)).toNumber());
        _nw49[(_dafny.ZERO).toNumber()] = _305_v129;
        _nw49[(_dafny.ONE).toNumber()] = _305_v129;
        _nw49[(new BigNumber(2)).toNumber()] = _305_v129;
        _nw49[(new BigNumber(3)).toNumber()] = _305_v129;
        _nw49[(new BigNumber(4)).toNumber()] = _305_v129;
        _nw49[(new BigNumber(5)).toNumber()] = _305_v129;
        _nw49[(new BigNumber(6)).toNumber()] = _305_v129;
        _nw49[(new BigNumber(7)).toNumber()] = _305_v129;
        _nw49[(new BigNumber(8)).toNumber()] = _305_v129;
        _nw49[(new BigNumber(9)).toNumber()] = _305_v129;
        _nw49[(new BigNumber(10)).toNumber()] = _305_v129;
        _nw49[(new BigNumber(11)).toNumber()] = _305_v129;
        _nw49[(new BigNumber(12)).toNumber()] = (_306_v130);
        _nw49[(new BigNumber(13)).toNumber()] = _305_v129;
        _nw49[(new BigNumber(14)).toNumber()] = _305_v129;
        _nw49[(new BigNumber(15)).toNumber()] = _305_v129;
        _nw49[(new BigNumber(16)).toNumber()] = _305_v129;
        _nw49[(new BigNumber(17)).toNumber()] = _305_v129;
        _nw49[(new BigNumber(18)).toNumber()] = _305_v129;
        _nw49[(new BigNumber(19)).toNumber()] = (_307_v131)[_module.__default.safeIndex((_dafny.ZERO).minus(_126_v7), new BigNumber((_307_v131).length))];
        _nw49[(new BigNumber(20)).toNumber()] = _305_v129;
        _nw49[(new BigNumber(21)).toNumber()] = _305_v129;
        _nw49[(new BigNumber(22)).toNumber()] = _305_v129;
        _nw49[(new BigNumber(23)).toNumber()] = _305_v129;
        _nw49[(new BigNumber(24)).toNumber()] = (_307_v131)[_module.__default.safeIndex(_126_v7, new BigNumber((_307_v131).length))];
        _nw49[(new BigNumber(25)).toNumber()] = (_306_v130);
        _nw49[(new BigNumber(26)).toNumber()] = _305_v129;
        _nw49[(new BigNumber(27)).toNumber()] = _305_v129;
        _nw49[(new BigNumber(28)).toNumber()] = _305_v129;
        _308_v132 = _nw49;
        _126_v7 = new BigNumber((_dafny.Seq.of(_308_v132, _308_v132, _308_v132, _308_v132, _308_v132)).length);
        let _rhs31 = (_129_v10)[_module.__default.safeIndex(_126_v7, new BigNumber((_129_v10).length))];
        let _rhs32 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt(_126_v7, _126_v7), (_126_v7).plus(new BigNumber((_129_v10).length)));
        let _rhs33 = _module.__default.fm1(_119_v0, _138_globalState);
        let _lhs18 = _138_globalState;
        _lhs18.f16 = _rhs31;
        _126_v7 = _rhs32;
        _126_v7 = _rhs33;
        if (_122_v3) {
          let _309_v133;
          let _nw50 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Seq.of());
          _309_v133 = _nw50;
          let _index33 = _module.__default.safeIndex(new BigNumber(629), new BigNumber((_309_v133).length));
          (_309_v133)[_index33] = _dafny.Seq.of(_135_v14);
          let _index34 = _module.__default.safeIndex(new BigNumber(963), new BigNumber((_124_v5).length));
          (_124_v5)[_index34] = _module.__default.safeDivisionInt(new BigNumber((_252_v96).cardinality()), new BigNumber(-566));
          let _310_v134;
          _310_v134 = _dafny.Set.fromElements((_269_v108).f30);
          let _311_v135;
          _311_v135 = _dafny.Map.Empty.slice().updateUnsafe(_126_v7,_310_v134);
          let _index35 = _module.__default.safeIndex(new BigNumber(629), new BigNumber((_309_v133).length));
          let _index36 = _module.__default.safeIndex(new BigNumber(963), new BigNumber((_124_v5).length));
          let _rhs34 = _dafny.Seq.Concat(_251_v95, _251_v95);
          let _rhs35 = (_module.__default.safeModuloInt(_126_v7, new BigNumber(((((_311_v135).contains(_126_v7)) ? ((_311_v135).get(_126_v7)) : (_dafny.Set.fromElements(true)))).length))).plus((_126_v7).plus((((_128_v9).contains((_269_v108).f30)) ? ((_128_v9).get((_269_v108).f30)) : (new BigNumber(885)))));
          let _rhs36 = _module.__default.fm7((_269_v108).f29, _122_v3, _138_globalState);
          let _rhs37 = new BigNumber(643);
          let _rhs38 = (_269_v108).f29;
          let _lhs19 = _309_v133;
          let _lhs20 = _module.__default.safeIndex(new BigNumber(629), new BigNumber((_309_v133).length));
          let _lhs21 = _124_v5;
          let _lhs22 = _module.__default.safeIndex(new BigNumber(963), new BigNumber((_124_v5).length));
          _lhs19[_lhs20] = _rhs34;
          _126_v7 = _rhs35;
          _lhs21[_lhs22] = _rhs36;
          _126_v7 = _rhs37;
          _122_v3 = _rhs38;
          _126_v7 = (_dafny.ZERO).minus(_module.__default.fm1(new _dafny.CodePoint('r'.codePointAt(0)), _138_globalState));
          let _312_v136;
          _312_v136 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_126_v7,(_269_v108).f30));
          (_138_globalState).f16 = !_dafny.Seq.contains(_312_v136, _247_v91);
          (_138_globalState).f16 = !(!(((((_269_v108).f30) === ((_269_v108).f30)) ? (((_269_v108).f29) && ((_269_v108).f30)) : (!(!((_269_v108).f30))))));
          let _rhs39 = _126_v7;
          let _rhs40 = ((_309_v133)[_module.__default.safeIndex(new BigNumber(629), new BigNumber((_309_v133).length))])[_module.__default.safeIndex(_126_v7, new BigNumber(((_309_v133)[_module.__default.safeIndex(new BigNumber(629), new BigNumber((_309_v133).length))]).length))];
          let _rhs41 = _121_v2;
          let _rhs42 = _dafny.Seq.update(_121_v2, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_128_v9).length)), new BigNumber((_121_v2).length)), _119_v0);
          let _lhs23 = _138_globalState;
          _126_v7 = _rhs39;
          _lhs23.f27 = _rhs40;
          _121_v2 = _rhs41;
          _121_v2 = _rhs42;
        } else {
          let _313_v137;
          _313_v137 = _module.D3.create_DC10(_121_v2, _126_v7, _122_v3);
          _126_v7 = (_313_v137).dtor_cf14;
          _126_v7 = new BigNumber(915);
          _module.__default.m0(_126_v7, _138_globalState);
          let _314_v138;
          _314_v138 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-779)), ((_315_v0) => function (_316_i12) {
            return _315_v0;
          })(_119_v0)));
          _126_v7 = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat((((_269_v108).f29) ? (_314_v138) : (_314_v138)), _314_v138), _module.__default.safeIndex(_126_v7, new BigNumber((_dafny.Seq.Concat((((_269_v108).f29) ? (_314_v138) : (_314_v138)), _314_v138)).length)), _dafny.Seq.Concat(_121_v2, _121_v2))).length);
          (_138_globalState).f16 = (_269_v108).f29;
        }
        _126_v7 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_126_v7));
      }
      let _317_i13;
      _317_i13 = _dafny.ZERO;
      L3: {
        while (_122_v3) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_317_i13)) {
              break L3;
            }
            _317_i13 = (_317_i13).plus(_dafny.ONE);
            if ((_129_v10)[_module.__default.safeIndex(_126_v7, new BigNumber((_129_v10).length))]) {
              let _index37 = _module.__default.safeIndex(new BigNumber(148), new BigNumber((_124_v5).length));
              (_124_v5)[_index37] = _module.__default.safeDivisionInt(_126_v7, _126_v7);
              let _318_v139;
              let _nw51 = Array((new BigNumber(2)).toNumber()).fill([]);
              _318_v139 = _nw51;
              let _index38 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_318_v139).length));
              (_318_v139)[_index38] = _135_v14;
              let _319_v140;
              _319_v140 = _dafny.Seq.of(_module.__default.fm14(_126_v7, _126_v7, _126_v7, (_269_v108).f29, _138_globalState), _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.of((_269_v108).f29), _module.__default.safeIndex(new BigNumber(-53), new BigNumber((_dafny.Seq.of((_269_v108).f29)).length)), (_269_v108).f30), _129_v10));
              let _index39 = _module.__default.safeIndex(new BigNumber(148), new BigNumber((_124_v5).length));
              let _index40 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_318_v139).length));
              let _rhs43 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_141_v19,(_269_v108).f30)).length);
              let _rhs44 = _dafny.Seq.Concat(_121_v2, _121_v2);
              let _rhs45 = _135_v14;
              let _rhs46 = _dafny.Seq.Concat(_319_v140, _319_v140);
              let _rhs47 = new BigNumber(101);
              let _lhs24 = _124_v5;
              let _lhs25 = _module.__default.safeIndex(new BigNumber(148), new BigNumber((_124_v5).length));
              let _lhs26 = _318_v139;
              let _lhs27 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_318_v139).length));
              _lhs24[_lhs25] = _rhs43;
              _121_v2 = _rhs44;
              _lhs26[_lhs27] = _rhs45;
              _319_v140 = _rhs46;
              _126_v7 = _rhs47;
              let _index41 = _module.__default.safeIndex(new BigNumber(148), new BigNumber((_124_v5).length));
              (_124_v5)[_index41] = _126_v7;
              (_138_globalState).f16 = (_269_v108).f30;
              let _320_v141;
              let _nw52 = new _module.C1();
              _nw52.__ctor((_124_v5)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_124_v5).length))], _318_v139, !((_269_v108).f30));
              _320_v141 = _nw52;
              _320_v141 = _320_v141;
              let _index42 = _module.__default.safeIndex(new BigNumber(148), new BigNumber((_124_v5).length));
              (_124_v5)[_index42] = (_126_v7).multipliedBy(((_124_v5)[_module.__default.safeIndex(new BigNumber(148), new BigNumber((_124_v5).length))]).minus((_141_v19)[_module.__default.safeIndex(_126_v7, new BigNumber((_141_v19).length))]));
            } else {
              let _321_v142;
              let _nw53 = new _module.C5();
              _nw53.__ctor((_269_v108).f29, (_269_v108).f30);
              _321_v142 = _nw53;
              let _322_v143;
              _322_v143 = _module.D0.create_DC1(_119_v0, (_269_v108).f30, true);
              let _323_v144;
              let _324_v145;
              let _out8;
              let _out9;
              let _outcollector3 = (_321_v142).m1(_141_v19, _module.__default.safeDivisionInt(_126_v7, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_129_v10,_126_v7)).length)), _322_v143, _138_globalState);
              _out8 = _outcollector3[0];
              _out9 = _outcollector3[1];
              _323_v144 = _out8;
              _324_v145 = _out9;
              let _325_v146;
              let _326_v147;
              let _out10;
              let _out11;
              let _outcollector4 = (_321_v142).m1(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(877)), ((_327_v7) => function (_328_i14) {
                return _327_v7;
              })(_126_v7)), _141_v19), _126_v7, _322_v143, _138_globalState);
              _out10 = _outcollector4[0];
              _out11 = _outcollector4[1];
              _325_v146 = _out10;
              _326_v147 = _out11;
              _325_v146 = (_269_v108).f30;
              let _329_v148;
              let _nw54 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Seq.of());
              _329_v148 = _nw54;
              let _330_v149;
              let _nw55 = Array((new BigNumber(17)).toNumber());
              _nw55[(_dafny.ZERO).toNumber()] = _329_v148;
              _nw55[(_dafny.ONE).toNumber()] = _329_v148;
              _nw55[(new BigNumber(2)).toNumber()] = _329_v148;
              _nw55[(new BigNumber(3)).toNumber()] = _329_v148;
              _nw55[(new BigNumber(4)).toNumber()] = _329_v148;
              _nw55[(new BigNumber(5)).toNumber()] = _329_v148;
              _nw55[(new BigNumber(6)).toNumber()] = _329_v148;
              _nw55[(new BigNumber(7)).toNumber()] = _329_v148;
              _nw55[(new BigNumber(8)).toNumber()] = _329_v148;
              _nw55[(new BigNumber(9)).toNumber()] = _329_v148;
              _nw55[(new BigNumber(10)).toNumber()] = _329_v148;
              _nw55[(new BigNumber(11)).toNumber()] = _329_v148;
              _nw55[(new BigNumber(12)).toNumber()] = _329_v148;
              _nw55[(new BigNumber(13)).toNumber()] = _329_v148;
              _nw55[(new BigNumber(14)).toNumber()] = _329_v148;
              _nw55[(new BigNumber(15)).toNumber()] = _329_v148;
              _nw55[(new BigNumber(16)).toNumber()] = _329_v148;
              _330_v149 = _nw55;
              let _index43 = _module.__default.safeIndex(new BigNumber(558), new BigNumber((_330_v149).length));
              (_330_v149)[_index43] = _329_v148;
              let _index44 = _module.__default.safeIndex(new BigNumber(558), new BigNumber((_330_v149).length));
              (_330_v149)[_index44] = _329_v148;
            }
            if (((_129_v10)[_module.__default.safeIndex(new BigNumber((_248_v92).length), new BigNumber((_129_v10).length))]) && (!(_module.__default.fm26((_269_v108).f29, (_269_v108).f29, (_129_v10)[_module.__default.safeIndex(_126_v7, new BigNumber((_129_v10).length))], (((_153_v25).contains(_module.__default.fm1(_119_v0, _138_globalState))) ? ((_153_v25).get(_module.__default.fm1(_119_v0, _138_globalState))) : (_126_v7)), _138_globalState)))) {
              let _331_v150;
              _331_v150 = _dafny.Set.fromElements(_119_v0);
              let _332_v151;
              _332_v151 = _dafny.Map.Empty.slice().updateUnsafe(_331_v150,_121_v2);
              _332_v151 = (_332_v151).update((_331_v150).Intersect(_331_v150), _121_v2);
              let _index45 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_135_v14).length));
              (_135_v14)[_index45] = (_269_v108).f29;
              let _index46 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_135_v14).length));
              (_135_v14)[_index46] = ((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_129_v10).length), _126_v7))).isLessThan(_126_v7);
              _126_v7 = new BigNumber(929);
              let _333_v152;
              _333_v152 = _dafny.Map.Empty.slice().updateUnsafe(_124_v5,_124_v5);
              _333_v152 = (_333_v152).update(_124_v5, _124_v5);
              let _334_v153;
              _334_v153 = _module.D0.create_DC1(_119_v0, (_269_v108).f29, (_269_v108).f30);
              let _335_v154;
              let _336_v155;
              let _out12;
              let _out13;
              let _outcollector5 = (_269_v108).m1(_141_v19, _126_v7, _334_v153, _138_globalState);
              _out12 = _outcollector5[0];
              _out13 = _outcollector5[1];
              _335_v154 = _out12;
              _336_v155 = _out13;
            } else {
              let _337_v156;
              _337_v156 = _dafny.Set.fromElements(new BigNumber(324));
              _126_v7 = new BigNumber(((_dafny.Set.fromElements(new BigNumber(-408))).Intersect(_337_v156)).length);
              _247_v91 = (_247_v91).update(new BigNumber((_247_v91).length), (_269_v108).f30);
              _126_v7 = _module.__default.safeDivisionInt(_126_v7, (new BigNumber((_121_v2).length)).multipliedBy(_126_v7));
              _126_v7 = (_dafny.ZERO).minus(_126_v7);
              let _338_v157;
              let _nw56 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Map.Empty);
              _338_v157 = _nw56;
              let _339_v158;
              _339_v158 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_119_v0, _138_globalState),_247_v91);
              let _340_v159;
              let _nw57 = new _module.C3();
              _nw57.__ctor(_338_v157, _339_v158, (_269_v108).f29);
              _340_v159 = _nw57;
            }
            let _rhs48 = (_269_v108).f29;
            let _rhs49 = ((_122_v3) ? (_121_v2) : (_dafny.Seq.update(_module.__default.fm0(_196_v61, _138_globalState), _module.__default.safeIndex(_126_v7, new BigNumber((_module.__default.fm0(_196_v61, _138_globalState)).length)), _119_v0)));
            let _rhs50 = _module.__default.fm10(_119_v0, (_126_v7).isEqualTo(_126_v7), _138_globalState);
            let _rhs51 = (((_269_v108).f30) ? (new _dafny.CodePoint('k'.codePointAt(0))) : (_119_v0));
            let _lhs28 = _138_globalState;
            let _lhs29 = _138_globalState;
            _lhs28.f28 = _rhs48;
            _121_v2 = _rhs49;
            _lhs29.f28 = _rhs50;
            _119_v0 = _rhs51;
            let _341_v160;
            let _init4 = ((_342_v108, _343_v7) => function (_344_i15) {
              return _module.D14.create_DC40((_342_v108).f29, (_342_v108).f29, _343_v7);
            })(_269_v108, _126_v7);
            let _nw58 = Array((new BigNumber(22)).toNumber());
            for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw58.length); _i0_4++) {
              _nw58[_i0_4] = _init4(new BigNumber(_i0_4));
            }
            _341_v160 = _nw58;
            let _345_v161;
            _345_v161 = _module.D14.create_DC40(!((_269_v108).f30), _122_v3, _126_v7);
            let _index47 = _module.__default.safeIndex(new BigNumber(601), new BigNumber((_341_v160).length));
            (_341_v160)[_index47] = _345_v161;
            let _index48 = _module.__default.safeIndex(new BigNumber(601), new BigNumber((_341_v160).length));
            (_341_v160)[_index48] = (((_269_v108).f29) ? (_345_v161) : (_345_v161));
          }
        }
      }
      process.stdout.write(_dafny.toString(_119_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_120_v1)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_120_v1)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_120_v1)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_121_v2).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_122_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_123_v4).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("mau"),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v5)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v5)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v5)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v5)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v5)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v5)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v5)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v5)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v5)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v5)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v5)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v5)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v5)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v5)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v5)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v5)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v5)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v5)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v5)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v5)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v5)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v5)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v5)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v5)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v5)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v5)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v5)[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_125_v6).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_126_v7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_127_v8).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_128_v9).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(159)).updateUnsafe(true,new BigNumber(62)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_129_v10, _dafny.Seq.of(false, false, true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v11).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(159), new BigNumber(159)),new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_131_v12)[_dafny.ZERO], _dafny.Seq.of(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_131_v12)[_dafny.ONE], _dafny.Seq.of(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_131_v12)[new BigNumber(2)], _dafny.Seq.of(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_131_v12)[new BigNumber(3)], _dafny.Seq.of(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_131_v12)[new BigNumber(4)], _dafny.Seq.of(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_131_v12)[new BigNumber(5)], _dafny.Seq.of(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_134_v13).dtor_cf0)[_dafny.ZERO], _dafny.Seq.of(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_134_v13).dtor_cf0)[_dafny.ONE], _dafny.Seq.of(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_134_v13).dtor_cf0)[new BigNumber(2)], _dafny.Seq.of(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_134_v13).dtor_cf0)[new BigNumber(3)], _dafny.Seq.of(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_134_v13).dtor_cf0)[new BigNumber(4)], _dafny.Seq.of(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_134_v13).dtor_cf0)[new BigNumber(5)], _dafny.Seq.of(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_v14)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_v14)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_v14)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_v14)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_v14)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_v14)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_v14)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_v14)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_v14)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_v14)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_v14)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_135_v14)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_136_v15))[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_136_v15))[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_136_v15))[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_136_v15))[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_136_v15))[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_136_v15))[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_136_v15))[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_136_v15))[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_136_v15))[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_136_v15))[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_136_v15))[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_136_v15))[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_136_v15))[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_136_v15))[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_136_v15))[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_136_v15))[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_136_v15))[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_136_v15))[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_136_v15))[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_136_v15))[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_136_v15))[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_136_v15))[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_136_v15))[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_136_v15))[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_136_v15))[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_136_v15))[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_136_v15))[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_137_v16).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f0)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f0)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f0)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f3).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("mau"),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_138_globalState).f6).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_138_globalState).f11).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f12).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(159), new BigNumber(159)),new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState).f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_138_globalState.f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState).f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_138_globalState).f18)[_dafny.ZERO], _dafny.Seq.of(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_138_globalState).f18)[_dafny.ONE], _dafny.Seq.of(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_138_globalState).f18)[new BigNumber(2)], _dafny.Seq.of(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_138_globalState).f18)[new BigNumber(3)], _dafny.Seq.of(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_138_globalState).f18)[new BigNumber(4)], _dafny.Seq.of(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_138_globalState).f18)[new BigNumber(5)], _dafny.Seq.of(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState).f19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f20)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f20)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f20)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f20)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f20)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f20)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f20)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f20)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f20)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f20)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f20)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f20)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f21)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f21)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f21)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f21)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f21)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f21)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f21)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f21)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f21)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f21)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f21)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f21)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f21)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f21)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f21)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f21)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f21)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f21)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f21)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f21)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f21)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f21)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f21)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f21)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f21)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f21)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f21)[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState).f22));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_138_globalState).f23).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f24)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f24)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f24)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f24)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f24)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f24)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f24)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f24)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f24)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f24)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f24)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f24)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f24)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f24)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f24)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f24)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f24)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f24)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f24)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f24)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f24)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f24)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f24)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState).f25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState).f26));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f27)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f27)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f27)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f27)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f27)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f27)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f27)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f27)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f27)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f27)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f27)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_globalState.f27)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_138_globalState.f28));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_139_v17).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_140_v18).equals(_dafny.Set.fromElements(_module.D0.create_DC3(new _dafny.CodePoint('v'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_141_v19, _dafny.Seq.of(new BigNumber(159), new BigNumber(159), new BigNumber(159), new BigNumber(159)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_152_v24).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(159),new BigNumber(159)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_153_v25).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(159),new BigNumber(159)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_187_i3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_190_i5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_196_v61).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_197_v62, _dafny.Seq.of(_dafny.MultiSet.fromElements(true, true, true, true), _dafny.MultiSet.fromElements(true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_198_v63).dtor_cf18).equals(_dafny.MultiSet.fromElements(true, true, true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_247_v91).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(3),true).updateUnsafe(_dafny.ONE,true).updateUnsafe(new BigNumber(2),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_248_v92).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(3),true).updateUnsafe(new BigNumber(796),true),new BigNumber(796)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_249_v93).dtor_cf61).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(3),true).updateUnsafe(new BigNumber(796),true),new BigNumber(796)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_250_v94).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_251_v95).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_252_v96).equals(_dafny.MultiSet.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_253_v97, _dafny.Seq.of(_dafny.MultiSet.fromElements(_dafny.ONE)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_254_v98).equals(_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(_dafny.ONE), _dafny.MultiSet.fromElements(_dafny.ONE), _dafny.MultiSet.fromElements(_dafny.ONE)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_255_v99).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(new BigNumber(159), new BigNumber(159), new BigNumber(159), new BigNumber(159))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_269_v108).f29));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_269_v108).f30));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_270_v109).dtor_cf24)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_270_v109).dtor_cf24)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_270_v109).dtor_cf24)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_317_i13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1, cf2, cf3) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC2(cf4) {
      let $dt = new D0(1);
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC3(cf5) {
      let $dt = new D0(2);
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(3);
      $dt.cf0 = cf0;
      return $dt;
    }
    static create_DC4(cf6) {
      let $dt = new D0(4);
      $dt.cf6 = cf6;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get is_DC0() { return this.$tag === 3; }
    get is_DC4() { return this.$tag === 4; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf6() { return this.cf6; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC3" + "(" + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 3) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else if (this.$tag === 4) {
        return "D0.DC4" + "(" + _dafny.toString(this.cf6) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1) && this.cf2 === other.cf2 && this.cf3 === other.cf3;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf4 === other.cf4;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf5, other.cf5);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf0 === other.cf0;
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf6, other.cf6);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(new _dafny.CodePoint('D'.codePointAt(0)), false, false);
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
    static create_DC5(cf7) {
      let $dt = new D1(0);
      $dt.cf7 = cf7;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get dtor_cf7() { return this.cf7; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf7) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf7 === other.cf7;
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
          return D1.Default();
        }
      };
    }
  }

  $module.D2 = class D2 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC7(cf9, cf10) {
      let $dt = new D2(0);
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC6(cf8) {
      let $dt = new D2(1);
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC8(cf11) {
      let $dt = new D2(2);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf9, other.cf9) && _dafny.areEqual(this.cf10, other.cf10);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf8, other.cf8);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf11, other.cf11);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC7(new _dafny.CodePoint('D'.codePointAt(0)), _dafny.Set.Empty);
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
    static create_DC10(cf13, cf14, cf15) {
      let $dt = new D3(0);
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC9(cf12) {
      let $dt = new D3(1);
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC11(cf16) {
      let $dt = new D3(2);
      $dt.cf16 = cf16;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get is_DC11() { return this.$tag === 2; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf16() { return this.cf16; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC10" + "(" + this.cf13.toVerbatimString(true) + ", " + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf16) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf13, other.cf13) && _dafny.areEqual(this.cf14, other.cf14) && this.cf15 === other.cf15;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf12, other.cf12);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf16, other.cf16);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC10(_dafny.Seq.UnicodeFromString(""), _dafny.ZERO, false);
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
    static create_DC12(cf17) {
      let $dt = new D4(0);
      $dt.cf17 = cf17;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get dtor_cf17() { return this.cf17; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf17) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf17, other.cf17);
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
          return D4.Default();
        }
      };
    }
  }

  $module.D5 = class D5 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC14(cf19, cf20, cf21, cf22) {
      let $dt = new D5(0);
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC13(cf18) {
      let $dt = new D5(1);
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC15(cf23) {
      let $dt = new D5(2);
      $dt.cf23 = cf23;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get is_DC15() { return this.$tag === 2; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf23() { return this.cf23; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf23) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf19, other.cf19) && _dafny.areEqual(this.cf20, other.cf20) && this.cf21 === other.cf21 && this.cf22 === other.cf22;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf18, other.cf18);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf23, other.cf23);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC14(_dafny.Set.Empty, _dafny.ZERO, false, []);
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
    static create_DC17(cf25, cf26, cf27) {
      let $dt = new D6(0);
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC16(cf24) {
      let $dt = new D6(1);
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC18(cf28) {
      let $dt = new D6(2);
      $dt.cf28 = cf28;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get is_DC18() { return this.$tag === 2; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf28() { return this.cf28; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC18" + "(" + _dafny.toString(this.cf28) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf25 === other.cf25 && this.cf26 === other.cf26 && _dafny.areEqual(this.cf27, other.cf27);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf24 === other.cf24;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf28, other.cf28);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC17(false, [], _dafny.ZERO);
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
    static create_DC19(cf29) {
      let $dt = new D7(0);
      $dt.cf29 = cf29;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get dtor_cf29() { return this.cf29; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC19" + "(" + _dafny.toString(this.cf29) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf29, other.cf29);
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
          return D7.Default();
        }
      };
    }
  }

  $module.D8 = class D8 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC21(cf31, cf32, cf33) {
      let $dt = new D8(0);
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC22(cf34, cf35, cf36, cf37, cf38) {
      let $dt = new D8(1);
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC20(cf30) {
      let $dt = new D8(2);
      $dt.cf30 = cf30;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get is_DC22() { return this.$tag === 1; }
    get is_DC20() { return this.$tag === 2; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf30() { return this.cf30; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC21" + "(" + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC22" + "(" + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC20" + "(" + _dafny.toString(this.cf30) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf31 === other.cf31 && this.cf32 === other.cf32 && _dafny.areEqual(this.cf33, other.cf33);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf34 === other.cf34 && this.cf35 === other.cf35 && _dafny.areEqual(this.cf36, other.cf36) && this.cf37 === other.cf37 && this.cf38 === other.cf38;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf30, other.cf30);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC21(false, false, _dafny.ZERO);
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
    static create_DC24(cf40) {
      let $dt = new D9(0);
      $dt.cf40 = cf40;
      return $dt;
    }
    static create_DC25(cf41, cf42, cf43, cf44) {
      let $dt = new D9(1);
      $dt.cf41 = cf41;
      $dt.cf42 = cf42;
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      return $dt;
    }
    static create_DC26(cf45, cf46, cf47, cf48, cf49) {
      let $dt = new D9(2);
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      $dt.cf49 = cf49;
      return $dt;
    }
    static create_DC23(cf39) {
      let $dt = new D9(3);
      $dt.cf39 = cf39;
      return $dt;
    }
    get is_DC24() { return this.$tag === 0; }
    get is_DC25() { return this.$tag === 1; }
    get is_DC26() { return this.$tag === 2; }
    get is_DC23() { return this.$tag === 3; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf39() { return this.cf39; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC24" + "(" + _dafny.toString(this.cf40) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC25" + "(" + _dafny.toString(this.cf41) + ", " + _dafny.toString(this.cf42) + ", " + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC26" + "(" + _dafny.toString(this.cf45) + ", " + this.cf46.toVerbatimString(true) + ", " + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ", " + _dafny.toString(this.cf49) + ")";
      } else if (this.$tag === 3) {
        return "D9.DC23" + "(" + _dafny.toString(this.cf39) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf40, other.cf40);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf41, other.cf41) && this.cf42 === other.cf42 && this.cf43 === other.cf43 && this.cf44 === other.cf44;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf45 === other.cf45 && _dafny.areEqual(this.cf46, other.cf46) && _dafny.areEqual(this.cf47, other.cf47) && _dafny.areEqual(this.cf48, other.cf48) && this.cf49 === other.cf49;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf39, other.cf39);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC24(_dafny.MultiSet.Empty);
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
    static create_DC28(cf51) {
      let $dt = new D10(0);
      $dt.cf51 = cf51;
      return $dt;
    }
    static create_DC27(cf50) {
      let $dt = new D10(1);
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC29(cf52) {
      let $dt = new D10(2);
      $dt.cf52 = cf52;
      return $dt;
    }
    get is_DC28() { return this.$tag === 0; }
    get is_DC27() { return this.$tag === 1; }
    get is_DC29() { return this.$tag === 2; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf52() { return this.cf52; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC28" + "(" + _dafny.toString(this.cf51) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC27" + "(" + _dafny.toString(this.cf50) + ")";
      } else if (this.$tag === 2) {
        return "D10.DC29" + "(" + _dafny.toString(this.cf52) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf51, other.cf51);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf50, other.cf50);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf52, other.cf52);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC28(_dafny.ZERO);
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
    static create_DC30(cf53) {
      let $dt = new D11(0);
      $dt.cf53 = cf53;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get dtor_cf53() { return this.cf53; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC30" + "(" + _dafny.toString(this.cf53) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf53 === other.cf53;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return null;
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
    static create_DC32(cf55, cf56, cf57, cf58, cf59) {
      let $dt = new D12(0);
      $dt.cf55 = cf55;
      $dt.cf56 = cf56;
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      return $dt;
    }
    static create_DC33(cf60) {
      let $dt = new D12(1);
      $dt.cf60 = cf60;
      return $dt;
    }
    static create_DC31(cf54) {
      let $dt = new D12(2);
      $dt.cf54 = cf54;
      return $dt;
    }
    get is_DC32() { return this.$tag === 0; }
    get is_DC33() { return this.$tag === 1; }
    get is_DC31() { return this.$tag === 2; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf54() { return this.cf54; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC32" + "(" + _dafny.toString(this.cf55) + ", " + _dafny.toString(this.cf56) + ", " + _dafny.toString(this.cf57) + ", " + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC33" + "(" + this.cf60.toVerbatimString(true) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC31" + "(" + _dafny.toString(this.cf54) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf55 === other.cf55 && this.cf56 === other.cf56 && this.cf57 === other.cf57 && _dafny.areEqual(this.cf58, other.cf58) && _dafny.areEqual(this.cf59, other.cf59);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf60, other.cf60);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf54, other.cf54);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC32(false, false, null, _dafny.Set.Empty, _dafny.ZERO);
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
    static create_DC35(cf62) {
      let $dt = new D13(0);
      $dt.cf62 = cf62;
      return $dt;
    }
    static create_DC36(cf63, cf64) {
      let $dt = new D13(1);
      $dt.cf63 = cf63;
      $dt.cf64 = cf64;
      return $dt;
    }
    static create_DC37(cf65, cf66) {
      let $dt = new D13(2);
      $dt.cf65 = cf65;
      $dt.cf66 = cf66;
      return $dt;
    }
    static create_DC34(cf61) {
      let $dt = new D13(3);
      $dt.cf61 = cf61;
      return $dt;
    }
    static create_DC38(cf67) {
      let $dt = new D13(4);
      $dt.cf67 = cf67;
      return $dt;
    }
    get is_DC35() { return this.$tag === 0; }
    get is_DC36() { return this.$tag === 1; }
    get is_DC37() { return this.$tag === 2; }
    get is_DC34() { return this.$tag === 3; }
    get is_DC38() { return this.$tag === 4; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf67() { return this.cf67; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC35" + "(" + _dafny.toString(this.cf62) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC36" + "(" + _dafny.toString(this.cf63) + ", " + _dafny.toString(this.cf64) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC37" + "(" + _dafny.toString(this.cf65) + ", " + _dafny.toString(this.cf66) + ")";
      } else if (this.$tag === 3) {
        return "D13.DC34" + "(" + _dafny.toString(this.cf61) + ")";
      } else if (this.$tag === 4) {
        return "D13.DC38" + "(" + _dafny.toString(this.cf67) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf62, other.cf62);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf63 === other.cf63 && this.cf64 === other.cf64;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf65 === other.cf65 && _dafny.areEqual(this.cf66, other.cf66);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf61, other.cf61);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf67, other.cf67);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC35(_dafny.ZERO);
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
    static create_DC40(cf69, cf70, cf71) {
      let $dt = new D14(0);
      $dt.cf69 = cf69;
      $dt.cf70 = cf70;
      $dt.cf71 = cf71;
      return $dt;
    }
    static create_DC41(cf72, cf73) {
      let $dt = new D14(1);
      $dt.cf72 = cf72;
      $dt.cf73 = cf73;
      return $dt;
    }
    static create_DC39(cf68) {
      let $dt = new D14(2);
      $dt.cf68 = cf68;
      return $dt;
    }
    static create_DC42(cf74) {
      let $dt = new D14(3);
      $dt.cf74 = cf74;
      return $dt;
    }
    get is_DC40() { return this.$tag === 0; }
    get is_DC41() { return this.$tag === 1; }
    get is_DC39() { return this.$tag === 2; }
    get is_DC42() { return this.$tag === 3; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf73() { return this.cf73; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf74() { return this.cf74; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC40" + "(" + _dafny.toString(this.cf69) + ", " + _dafny.toString(this.cf70) + ", " + _dafny.toString(this.cf71) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC41" + "(" + _dafny.toString(this.cf72) + ", " + _dafny.toString(this.cf73) + ")";
      } else if (this.$tag === 2) {
        return "D14.DC39" + "(" + _dafny.toString(this.cf68) + ")";
      } else if (this.$tag === 3) {
        return "D14.DC42" + "(" + _dafny.toString(this.cf74) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf69 === other.cf69 && this.cf70 === other.cf70 && _dafny.areEqual(this.cf71, other.cf71);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf72, other.cf72) && _dafny.areEqual(this.cf73, other.cf73);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf68 === other.cf68;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf74, other.cf74);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC40(false, false, _dafny.ZERO);
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
    static create_DC43(cf75) {
      let $dt = new D15(0);
      $dt.cf75 = cf75;
      return $dt;
    }
    get is_DC43() { return this.$tag === 0; }
    get dtor_cf75() { return this.cf75; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC43" + "(" + _dafny.toString(this.cf75) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf75 === other.cf75;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return null;
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
    static create_DC45(cf77, cf78, cf79) {
      let $dt = new D16(0);
      $dt.cf77 = cf77;
      $dt.cf78 = cf78;
      $dt.cf79 = cf79;
      return $dt;
    }
    static create_DC44(cf76) {
      let $dt = new D16(1);
      $dt.cf76 = cf76;
      return $dt;
    }
    get is_DC45() { return this.$tag === 0; }
    get is_DC44() { return this.$tag === 1; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf79() { return this.cf79; }
    get dtor_cf76() { return this.cf76; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC45" + "(" + _dafny.toString(this.cf77) + ", " + _dafny.toString(this.cf78) + ", " + _dafny.toString(this.cf79) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC44" + "(" + _dafny.toString(this.cf76) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf77, other.cf77) && this.cf78 === other.cf78 && _dafny.areEqual(this.cf79, other.cf79);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf76, other.cf76);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC45(new _dafny.CodePoint('D'.codePointAt(0)), false, _dafny.ZERO);
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
    static create_DC47(cf81, cf82) {
      let $dt = new D17(0);
      $dt.cf81 = cf81;
      $dt.cf82 = cf82;
      return $dt;
    }
    static create_DC46(cf80) {
      let $dt = new D17(1);
      $dt.cf80 = cf80;
      return $dt;
    }
    get is_DC47() { return this.$tag === 0; }
    get is_DC46() { return this.$tag === 1; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf82() { return this.cf82; }
    get dtor_cf80() { return this.cf80; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC47" + "(" + _dafny.toString(this.cf81) + ", " + _dafny.toString(this.cf82) + ")";
      } else if (this.$tag === 1) {
        return "D17.DC46" + "(" + _dafny.toString(this.cf80) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf81 === other.cf81 && _dafny.areEqual(this.cf82, other.cf82);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf80, other.cf80);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D17.create_DC47(false, _dafny.ZERO);
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
    static create_DC49() {
      let $dt = new D18(0);
      return $dt;
    }
    static create_DC50(cf84, cf85) {
      let $dt = new D18(1);
      $dt.cf84 = cf84;
      $dt.cf85 = cf85;
      return $dt;
    }
    static create_DC51(cf86, cf87, cf88, cf89, cf90) {
      let $dt = new D18(2);
      $dt.cf86 = cf86;
      $dt.cf87 = cf87;
      $dt.cf88 = cf88;
      $dt.cf89 = cf89;
      $dt.cf90 = cf90;
      return $dt;
    }
    static create_DC48(cf83) {
      let $dt = new D18(3);
      $dt.cf83 = cf83;
      return $dt;
    }
    get is_DC49() { return this.$tag === 0; }
    get is_DC50() { return this.$tag === 1; }
    get is_DC51() { return this.$tag === 2; }
    get is_DC48() { return this.$tag === 3; }
    get dtor_cf84() { return this.cf84; }
    get dtor_cf85() { return this.cf85; }
    get dtor_cf86() { return this.cf86; }
    get dtor_cf87() { return this.cf87; }
    get dtor_cf88() { return this.cf88; }
    get dtor_cf89() { return this.cf89; }
    get dtor_cf90() { return this.cf90; }
    get dtor_cf83() { return this.cf83; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC49";
      } else if (this.$tag === 1) {
        return "D18.DC50" + "(" + _dafny.toString(this.cf84) + ", " + _dafny.toString(this.cf85) + ")";
      } else if (this.$tag === 2) {
        return "D18.DC51" + "(" + _dafny.toString(this.cf86) + ", " + _dafny.toString(this.cf87) + ", " + _dafny.toString(this.cf88) + ", " + _dafny.toString(this.cf89) + ", " + _dafny.toString(this.cf90) + ")";
      } else if (this.$tag === 3) {
        return "D18.DC48" + "(" + _dafny.toString(this.cf83) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf84, other.cf84) && _dafny.areEqual(this.cf85, other.cf85);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf86, other.cf86) && this.cf87 === other.cf87 && this.cf88 === other.cf88 && _dafny.areEqual(this.cf89, other.cf89) && this.cf90 === other.cf90;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf83, other.cf83);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC49();
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
    static create_DC53(cf92) {
      let $dt = new D19(0);
      $dt.cf92 = cf92;
      return $dt;
    }
    static create_DC54() {
      let $dt = new D19(1);
      return $dt;
    }
    static create_DC52(cf91) {
      let $dt = new D19(2);
      $dt.cf91 = cf91;
      return $dt;
    }
    get is_DC53() { return this.$tag === 0; }
    get is_DC54() { return this.$tag === 1; }
    get is_DC52() { return this.$tag === 2; }
    get dtor_cf92() { return this.cf92; }
    get dtor_cf91() { return this.cf91; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC53" + "(" + _dafny.toString(this.cf92) + ")";
      } else if (this.$tag === 1) {
        return "D19.DC54";
      } else if (this.$tag === 2) {
        return "D19.DC52" + "(" + _dafny.toString(this.cf91) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf92 === other.cf92;
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf91, other.cf91);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D19.create_DC53(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D19.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.T1 = class T1 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f12 = _dafny.Map.Empty;
      this.f16 = false;
      this.f20 = [];
      this.f24 = [];
      this.f27 = [];
      this.f28 = false;
      this._f0 = [];
      this._f1 = false;
      this._f2 = false;
      this._f3 = _dafny.Map.Empty;
      this._f4 = false;
      this._f5 = false;
      this._f6 = _dafny.Seq.UnicodeFromString("");
      this._f7 = _dafny.ZERO;
      this._f8 = false;
      this._f9 = false;
      this._f10 = false;
      this._f11 = _dafny.Seq.of();
      this._f13 = _dafny.ZERO;
      this._f14 = false;
      this._f15 = _dafny.ZERO;
      this._f17 = false;
      this._f18 = [];
      this._f19 = false;
      this._f21 = [];
      this._f22 = false;
      this._f23 = _dafny.Map.Empty;
      this._f25 = _dafny.ZERO;
      this._f26 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25, f26, f27, f28) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this).f12 = f12;
      (_this)._f13 = f13;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      (_this).f16 = f16;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      (_this)._f19 = f19;
      (_this).f20 = f20;
      (_this)._f21 = f21;
      (_this)._f22 = f22;
      (_this)._f23 = f23;
      (_this).f24 = f24;
      (_this)._f25 = f25;
      (_this)._f26 = f26;
      (_this).f27 = f27;
      (_this).f28 = f28;
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
    get f13() {
      let _this = this;
      return _this._f13;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
    get f15() {
      let _this = this;
      return _this._f15;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
    get f18() {
      let _this = this;
      return _this._f18;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
    get f21() {
      let _this = this;
      return _this._f21;
    };
    get f22() {
      let _this = this;
      return _this._f22;
    };
    get f23() {
      let _this = this;
      return _this._f23;
    };
    get f25() {
      let _this = this;
      return _this._f25;
    };
    get f26() {
      let _this = this;
      return _this._f26;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f31 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f31() {
      let _this = this;
      return _this._f31;
    };
    __ctor(f31) {
      let _this = this;
      (_this)._f31 = f31;
      return;
    }
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f31 = false;
      this._f32 = _dafny.ZERO;
      this._f33 = [];
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f31() {
      let _this = this;
      return _this._f31;
    };
    get f32() {
      let _this = this;
      return _this._f32;
    };
    get f33() {
      let _this = this;
      return _this._f33;
    };
    __ctor(f32, f33, f31) {
      let _this = this;
      (_this)._f32 = f32;
      (_this)._f33 = f33;
      (_this)._f31 = f31;
      return;
    }
    fm6(p0, globalState) {
      let _this = this;
      return _dafny.Seq.of((_this).f31, false, (_this).f31, true, !(((_this).f32).isEqualTo((_this).f32)));
    };
    fm16(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(531)), function (_346_i0) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      }), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(432)), function (_347_i1) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("lriggcwfc")));
    };
    fm17(p0, globalState) {
      let _this = this;
      return (_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f32))).IsSubsetOf((((_this).f31) ? (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ndqhhq")).length), (_this).f32, new BigNumber(679))) : (_dafny.MultiSet.fromElements(new BigNumber(-370), (_this).f32, new BigNumber((_dafny.MultiSet.fromElements((_this).f31, !((_this).f31), true, (_this).f31, (_this).f31)).cardinality())))));
    };
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = [];
      let r2 = _dafny.Seq.UnicodeFromString("");
      let r3 = _dafny.ZERO;
      let _348_v0;
      _348_v0 = _dafny.Seq.of(!(p0), (_this).f31, !(false), (_this).f31);
      let _349_v1;
      _349_v1 = _dafny.Seq.of((_this).f32, (_this).f32, new BigNumber((_348_v0).length), _module.__default.fm18((_this).f31, p0, globalState), (_this).f32);
      let _350_v2;
      _350_v2 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(269),_349_v1);
      _350_v2 = (_350_v2).update((_this).f32, _dafny.Seq.Concat(_dafny.Seq.of((_this).f32, (_this).f32), _349_v1));
      r0 = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of((_348_v0)[_module.__default.safeIndex(new BigNumber((_349_v1).length), new BigNumber((_348_v0).length))]), _dafny.Seq.Concat(_dafny.Seq.of((_this).f31, (_this).f31, (_this).f31), _348_v0)), _module.__default.safeIndex(((_this).f32).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-955)), function (_351_i0) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      })).length)), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of((_348_v0)[_module.__default.safeIndex(new BigNumber((_349_v1).length), new BigNumber((_348_v0).length))]), _dafny.Seq.Concat(_dafny.Seq.of((_this).f31, (_this).f31, (_this).f31), _348_v0))).length)), p0)).length);
      let _352_v3;
      let _nw59 = Array((new BigNumber(8)).toNumber()).fill([]);
      _352_v3 = _nw59;
      _352_v3 = _352_v3;
      let _353_i1;
      _353_i1 = _dafny.ZERO;
      L4: {
        while (false) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_353_i1)) {
              break L4;
            }
            _353_i1 = (_353_i1).plus(_dafny.ONE);
            let _354_v5;
            _354_v5 = _dafny.Map.Empty.slice().updateUnsafe(p0,true);
            let _355_v6;
            _355_v6 = _dafny.Seq.of(_354_v5);
            let _356_v7;
            let _nw60 = new _module.C0();
            _nw60.__ctor(p0);
            _356_v7 = _nw60;
            let _357_v8;
            _357_v8 = _dafny.MultiSet.fromElements(_356_v7, _356_v7);
            let _358_v9;
            _358_v9 = _dafny.Map.Empty.slice().updateUnsafe(_354_v5,new BigNumber((_357_v8).cardinality()));
            r3 = new BigNumber(((function () {
              let _coll10 = new _dafny.Map();
              for (const _compr_10 of (_355_v6).Elements) {
                let _359_v4 = _compr_10;
                if (_dafny.Seq.contains(_355_v6, _359_v4)) {
                  _coll10.push([_359_v4,(_this).f32]);
                }
              }
              return _coll10;
            }()).Merge(_358_v9)).length);
            r0 = (_this).f32;
            let _360_v10;
            _360_v10 = new _dafny.CodePoint('m'.codePointAt(0));
            if ((_this).fm17(_360_v10, globalState)) {
              let _361_v11;
              _361_v11 = _dafny.Set.fromElements((_this).f32);
              (globalState).f16 = ((false) ? ((_361_v11).IsProperSubsetOf(_361_v11)) : (((_this).f32).isLessThan((_this).f32)));
              let _362_v12;
              _362_v12 = _dafny.MultiSet.fromElements(!((_this).f32).isEqualTo((_this).f32), ((_this).f32).isLessThan(new BigNumber((_348_v0).length)));
              _362_v12 = ((p0) ? (p1) : (p2));
              (globalState).f28 = p0;
              let _363_v13;
              _363_v13 = _dafny.Set.fromElements((_this).f31, p0);
              let _364_v14;
              _364_v14 = _module.D2.create_DC7(_360_v10, _363_v13);
              let _365_v15;
              _365_v15 = _dafny.Seq.UnicodeFromString("hxuqqi");
              let _366_v16;
              _366_v16 = _dafny.Map.Empty.slice().updateUnsafe((_364_v14).dtor_cf10,_dafny.areEqual(_365_v15, _dafny.Seq.Create(_module.__default.abs(new BigNumber(726)), ((_367_v10) => function (_368_i2) {
                return _367_v10;
              })(_360_v10))));
              _366_v16 = _366_v16;
              let _369_v17;
              let _nw61 = new _module.C0();
              _nw61.__ctor((_this).f31);
              _369_v17 = _nw61;
            } else {
              let _370_v18;
              _370_v18 = _dafny.Seq.of(p1);
              let _371_v19;
              let _nw62 = Array((new BigNumber(5)).toNumber());
              _nw62[(_dafny.ZERO).toNumber()] = _360_v10;
              _nw62[(_dafny.ONE).toNumber()] = _360_v10;
              _nw62[(new BigNumber(2)).toNumber()] = _360_v10;
              _nw62[(new BigNumber(3)).toNumber()] = _360_v10;
              _nw62[(new BigNumber(4)).toNumber()] = new _dafny.CodePoint('v'.codePointAt(0));
              _371_v19 = _nw62;
              let _372_v20;
              _372_v20 = _module.D6.create_DC16(_371_v19);
              let _373_v21;
              _373_v21 = _dafny.Set.fromElements(_371_v19, (_372_v20).dtor_cf24, _371_v19, _371_v19, _371_v19);
              let _374_v22;
              _374_v22 = _dafny.Map.Empty.slice().updateUnsafe((_370_v18)[_module.__default.safeIndex(new BigNumber((_373_v21).length), new BigNumber((_370_v18).length))],(_this).f31);
              let _375_v24;
              _375_v24 = _dafny.Set.fromElements(p1, p2);
              let _376_v25;
              let _nw63 = Array((new BigNumber(10)).toNumber());
              _nw63[(_dafny.ZERO).toNumber()] = _374_v22;
              _nw63[(_dafny.ONE).toNumber()] = (_374_v22).Merge(_374_v22);
              _nw63[(new BigNumber(2)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(p2,p0)).Merge(_374_v22);
              _nw63[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f31);
              _nw63[(new BigNumber(4)).toNumber()] = _374_v22;
              _nw63[(new BigNumber(5)).toNumber()] = _374_v22;
              _nw63[(new BigNumber(6)).toNumber()] = function () {
                let _coll11 = new _dafny.Map();
                for (const _compr_11 of (_375_v24).Elements) {
                  let _377_v23 = _compr_11;
                  if ((_375_v24).contains(_377_v23)) {
                    _coll11.push([_377_v23,p0]);
                  }
                }
                return _coll11;
              }();
              _nw63[(new BigNumber(7)).toNumber()] = _374_v22;
              _nw63[(new BigNumber(8)).toNumber()] = (_374_v22).Merge(_374_v22);
              _nw63[(new BigNumber(9)).toNumber()] = (_374_v22).Merge(_374_v22);
              _376_v25 = _nw63;
              let _index49 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_376_v25).length));
              (_376_v25)[_index49] = function () {
                let _coll12 = new _dafny.Map();
                for (const _compr_12 of (_dafny.Seq.of(p2)).Elements) {
                  let _378_v26 = _compr_12;
                  if (_dafny.Seq.contains(_dafny.Seq.of(p2), _378_v26)) {
                    _coll12.push([_378_v26,!(p0)]);
                  }
                }
                return _coll12;
              }();
              let _379_v27;
              _379_v27 = _374_v22;
              let _index50 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_376_v25).length));
              (_376_v25)[_index50] = (((p0) ? (_379_v27) : (_379_v27)));
              let _380_v28;
              _380_v28 = _dafny.Seq.UnicodeFromString("kyqocej");
              let _381_v29;
              _381_v29 = _module.D3.create_DC10(_380_v28, new BigNumber(398), false);
              let _382_v30;
              _382_v30 = _dafny.Map.Empty.slice().updateUnsafe((_381_v29).dtor_cf13,p0);
              let _383_v31;
              let _nw64 = Array((new BigNumber(18)).toNumber());
              _nw64[(_dafny.ZERO).toNumber()] = (_this).f31;
              _nw64[(_dafny.ONE).toNumber()] = false;
              _nw64[(new BigNumber(2)).toNumber()] = (_this).f31;
              _nw64[(new BigNumber(3)).toNumber()] = p0;
              _nw64[(new BigNumber(4)).toNumber()] = (_this).f31;
              _nw64[(new BigNumber(5)).toNumber()] = (((_382_v30).contains(_380_v28)) ? ((_382_v30).get(_380_v28)) : (!(p0)));
              _nw64[(new BigNumber(6)).toNumber()] = (_this).f31;
              _nw64[(new BigNumber(7)).toNumber()] = (_this).f31;
              _nw64[(new BigNumber(8)).toNumber()] = (_this).f31;
              _nw64[(new BigNumber(9)).toNumber()] = (_this).f31;
              _nw64[(new BigNumber(10)).toNumber()] = (_this).f31;
              _nw64[(new BigNumber(11)).toNumber()] = false;
              _nw64[(new BigNumber(12)).toNumber()] = (_dafny.MultiSet.FromArray(_349_v1)).IsDisjointFrom((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(360)), ((_384_p0) => function (_385_i3) {
                return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f32,_384_p0)).length);
              })(p0)))).update((_this).f32, _module.__default.abs((_this).f32)));
              _nw64[(new BigNumber(13)).toNumber()] = p0;
              _nw64[(new BigNumber(14)).toNumber()] = ((_this).f31) === (p0);
              _nw64[(new BigNumber(15)).toNumber()] = p0;
              _nw64[(new BigNumber(16)).toNumber()] = p0;
              _nw64[(new BigNumber(17)).toNumber()] = p0;
              _383_v31 = _nw64;
              let _index51 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_383_v31).length));
              (_383_v31)[_index51] = ((_this).f31) && ((_this).f31);
              let _index52 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_383_v31).length));
              (_383_v31)[_index52] = false;
              (globalState).f16 = (_383_v31)[_module.__default.safeIndex(new BigNumber(928), new BigNumber((_383_v31).length))];
              let _index53 = _module.__default.safeIndex(new BigNumber(928), new BigNumber((_383_v31).length));
              (_383_v31)[_index53] = (((_this).f31) ? (((_this).f32).isLessThanOrEqualTo(new BigNumber(297))) : ((new BigNumber((_349_v1).length)).isEqualTo(new BigNumber((p1).cardinality()))));
              r0 = (_this).f32;
            }
            let _386_v32;
            _386_v32 = _dafny.Map.Empty.slice().updateUnsafe((_this).f32,(_this).f32);
            _386_v32 = (_386_v32).update(new BigNumber(982), new BigNumber(652));
          }
        }
      }
      let _387_v33;
      let _init5 = function (_388_i4) {
        return (_388_i4).multipliedBy((_this).f32);
      };
      let _nw65 = Array((new BigNumber(14)).toNumber());
      for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw65.length); _i0_5++) {
        _nw65[_i0_5] = _init5(new BigNumber(_i0_5));
      }
      _387_v33 = _nw65;
      let _index54 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_387_v33).length));
      (_387_v33)[_index54] = (_this).f32;
      let _389_v34;
      _389_v34 = _dafny.Map.Empty.slice().updateUnsafe((_this).f32,!((_this).f31));
      let _index55 = _module.__default.safeIndex(new BigNumber(109), new BigNumber((_387_v33).length));
      (_387_v33)[_index55] = new BigNumber(((_389_v34).Merge((_389_v34).Merge(_389_v34))).length);
      let _390_v35;
      _390_v35 = _dafny.MultiSet.fromElements((_this).f32);
      let _391_v36;
      let _nw66 = Array((new BigNumber(4)).toNumber());
      _nw66[(_dafny.ZERO).toNumber()] = (_this).f31;
      _nw66[(_dafny.ONE).toNumber()] = ((_387_v33)[_module.__default.safeIndex(new BigNumber(109), new BigNumber((_387_v33).length))]).isLessThanOrEqualTo((_this).f32);
      _nw66[(new BigNumber(2)).toNumber()] = (_this).f31;
      _nw66[(new BigNumber(3)).toNumber()] = p0;
      _391_v36 = _nw66;
      let _index56 = _module.__default.safeIndex(new BigNumber(564), new BigNumber((_391_v36).length));
      (_391_v36)[_index56] = false;
      let _392_v37;
      _392_v37 = _module.D3.create_DC10(_dafny.Seq.UnicodeFromString("cjaamgms"), (_this).f32, (_this).f31);
      let _393_v38;
      _393_v38 = _dafny.Seq.of(_390_v35, _390_v35);
      let _index57 = _module.__default.safeIndex(new BigNumber(564), new BigNumber((_391_v36).length));
      let _rhs52 = (_392_v37).dtor_cf13;
      let _rhs53 = ((_393_v38)[_module.__default.safeIndex(new BigNumber((_349_v1).length), new BigNumber((_393_v38).length))]).Union((_390_v35).update((_this).f32, _module.__default.abs((_this).f32)));
      let _rhs54 = true;
      let _lhs30 = _391_v36;
      let _lhs31 = _module.__default.safeIndex(new BigNumber(564), new BigNumber((_391_v36).length));
      r2 = _rhs52;
      _390_v35 = _rhs53;
      _lhs30[_lhs31] = _rhs54;
      r0 = (_this).f32;
      let _nw67 = Array((new BigNumber(14)).toNumber()).fill(_dafny.MultiSet.Empty);
      r1 = _nw67;
      let _394_v39;
      _394_v39 = _dafny.Seq.UnicodeFromString("kip");
      r2 = _394_v39;
      r3 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_this).f32, ((((p2).contains((_this).f31)) ? ((p2).get((_this).f31)) : ((_this).f32))).plus((_387_v33)[_module.__default.safeIndex(new BigNumber(109), new BigNumber((_387_v33).length))])));
      return [r0, r1, r2, r3];
    }
    m5(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = _dafny.MultiSet.Empty;
      let r3 = false;
      r0 = (_this).f32;
      let _395_v0;
      _395_v0 = _dafny.MultiSet.fromElements((_this).f32);
      let _396_v1;
      _396_v1 = _dafny.Seq.of(new BigNumber((_395_v0).cardinality()), (_this).f32);
      let _397_v2;
      _397_v2 = _dafny.Seq.of(_396_v1, _dafny.Seq.Concat(_dafny.Seq.of((_this).f32), _396_v1), _dafny.Seq.Concat(_396_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(595)), function (_398_i0) {
        return (_this).f32;
      })));
      let _399_v3;
      let _nw68 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
      _399_v3 = _nw68;
      let _index58 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_399_v3).length));
      (_399_v3)[_index58] = new BigNumber(291);
      let _400_v4;
      _400_v4 = new _dafny.CodePoint('o'.codePointAt(0));
      let _401_v5;
      _401_v5 = _dafny.Seq.of((_this).f31);
      let _402_v6;
      _402_v6 = _dafny.Set.fromElements(_401_v5, _dafny.Seq.of((_this).f31, (_this).f31), _401_v5, _401_v5, _dafny.Seq.of((_this).f31));
      let _index59 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_399_v3).length));
      let _rhs55 = (_this).f32;
      let _rhs56 = _397_v2;
      let _rhs57 = (_this).f32;
      let _rhs58 = _module.__default.fm18((_this).fm17(_400_v4, globalState), (_this).f31, globalState);
      let _rhs59 = (new BigNumber(((_402_v6).Union(_402_v6)).length)).multipliedBy((_this).f32);
      let _lhs32 = _399_v3;
      let _lhs33 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_399_v3).length));
      r0 = _rhs55;
      _397_v2 = _rhs56;
      r0 = _rhs57;
      r0 = _rhs58;
      _lhs32[_lhs33] = _rhs59;
      r1 = (_this).f31;
      let _403_i1;
      _403_i1 = _dafny.ZERO;
      L5: {
        while ((_this).f31) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_403_i1)) {
              break L5;
            }
            _403_i1 = (_403_i1).plus(_dafny.ONE);
            let _404_v7;
            _404_v7 = _dafny.Map.Empty.slice().updateUnsafe((_399_v3)[_module.__default.safeIndex(new BigNumber(373), new BigNumber((_399_v3).length))],_401_v5);
            let _405_v10;
            let _nw69 = Array((new BigNumber(14)).toNumber());
            _nw69[(_dafny.ZERO).toNumber()] = _404_v7;
            _nw69[(_dafny.ONE).toNumber()] = (_404_v7).Merge(_404_v7);
            _nw69[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f32,_dafny.Seq.of((_this).f31));
            _nw69[(new BigNumber(3)).toNumber()] = ((_404_v7).update((_399_v3)[_module.__default.safeIndex(new BigNumber(373), new BigNumber((_399_v3).length))], _401_v5)).Merge((_dafny.Map.Empty.slice().updateUnsafe((_this).f32,_401_v5)).update((_399_v3)[_module.__default.safeIndex(new BigNumber(373), new BigNumber((_399_v3).length))], _401_v5));
            _nw69[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-777),_401_v5);
            _nw69[(new BigNumber(5)).toNumber()] = _404_v7;
            _nw69[(new BigNumber(6)).toNumber()] = function () {
              let _coll13 = new _dafny.Map();
              for (const _compr_13 of _dafny.IntegerRange(new BigNumber(189), new BigNumber(37))) {
                let _406_v8 = _compr_13;
                if (((new BigNumber(189)).isLessThanOrEqualTo(_406_v8)) && ((_406_v8).isLessThan(new BigNumber(37)))) {
                  _coll13.push([_module.__default.safeModuloInt(_406_v8, new BigNumber(58)),_401_v5]);
                }
              }
              return _coll13;
            }();
            _nw69[(new BigNumber(7)).toNumber()] = _404_v7;
            _nw69[(new BigNumber(8)).toNumber()] = function () {
              let _coll14 = new _dafny.Map();
              for (const _compr_14 of _dafny.IntegerRange(new BigNumber(719), new BigNumber(649))) {
                let _407_v9 = _compr_14;
                if (((new BigNumber(719)).isLessThanOrEqualTo(_407_v9)) && ((_407_v9).isLessThan(new BigNumber(649)))) {
                  _coll14.push([(_407_v9).plus((_399_v3)[_module.__default.safeIndex(new BigNumber(373), new BigNumber((_399_v3).length))]),_401_v5]);
                }
              }
              return _coll14;
            }();
            _nw69[(new BigNumber(9)).toNumber()] = _404_v7;
            _nw69[(new BigNumber(10)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_399_v3)[_module.__default.safeIndex(new BigNumber(373), new BigNumber((_399_v3).length))],_401_v5);
            _nw69[(new BigNumber(11)).toNumber()] = (((_this).f31) ? (_404_v7) : (_404_v7));
            _nw69[(new BigNumber(12)).toNumber()] = _404_v7;
            _nw69[(new BigNumber(13)).toNumber()] = _404_v7;
            _405_v10 = _nw69;
            let _index60 = _module.__default.safeIndex(new BigNumber(778), new BigNumber((_405_v10).length));
            (_405_v10)[_index60] = _404_v7;
            let _index61 = _module.__default.safeIndex(new BigNumber(778), new BigNumber((_405_v10).length));
            (_405_v10)[_index61] = _404_v7;
            _396_v1 = _396_v1;
            let _408_v11;
            _408_v11 = _dafny.Seq.UnicodeFromString("smv");
            _408_v11 = _408_v11;
            let _rhs60 = _400_v4;
            let _rhs61 = _module.__default.fm18((((_this).f31) ? ((_this).f31) : ((_this).f31)), (_this).f31, globalState);
            let _rhs62 = _399_v3;
            _400_v4 = _rhs60;
            r0 = _rhs61;
            _399_v3 = _rhs62;
          }
        }
      }
      let _409_v12;
      _409_v12 = _dafny.Seq.UnicodeFromString("cxidk");
      let _hi2 = new BigNumber((_409_v12).length);
      for (let _410_i2 = (_399_v3)[_module.__default.safeIndex(new BigNumber(373), new BigNumber((_399_v3).length))]; _410_i2.isLessThan(_hi2); _410_i2 = _410_i2.plus(_dafny.ONE)) {
        let _411_v13;
        let _412_v14;
        let _413_v15;
        let _out14;
        let _out15;
        let _out16;
        let _outcollector6 = (_this).m6((_410_i2).isEqualTo(new BigNumber((_module.__default.fm19((_this).f32, new BigNumber((_409_v12).length), globalState)).length)), globalState);
        _out14 = _outcollector6[0];
        _out15 = _outcollector6[1];
        _out16 = _outcollector6[2];
        _411_v13 = _out14;
        _412_v14 = _out15;
        _413_v15 = _out16;
        let _source6 = function (_pat_let8_0) {
          return function (_414_dt__update__tmp_h0) {
            return function (_pat_let9_0) {
              return function (_415_dt__update_hcf15_h0) {
                return function (_pat_let10_0) {
                  return function (_416_dt__update_hcf14_h0) {
                    return _module.D3.create_DC10((_414_dt__update__tmp_h0).dtor_cf13, _416_dt__update_hcf14_h0, _415_dt__update_hcf15_h0);
                  }(_pat_let10_0);
                }((_this).f32);
              }(_pat_let9_0);
            }(true);
          }(_pat_let8_0);
        }(_module.D3.create_DC10(_409_v12, new BigNumber(377), (_this).f31));
        if (_source6.is_DC10) {
          let _417___mcc_h0 = (_source6).cf13;
          let _418___mcc_h1 = (_source6).cf14;
          let _419___mcc_h2 = (_source6).cf15;
          let _420_cf15 = _419___mcc_h2;
          let _421_cf14 = _418___mcc_h1;
          let _422_cf13 = _417___mcc_h0;
          r0 = _421_cf14;
          let _index62 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_399_v3).length));
          (_399_v3)[_index62] = (_this).f32;
          let _423_v16;
          let _nw70 = Array((new BigNumber(15)).toNumber()).fill([]);
          _423_v16 = _nw70;
          let _index63 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_423_v16).length));
          (_423_v16)[_index63] = _399_v3;
          let _index64 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_423_v16).length));
          (_423_v16)[_index64] = _399_v3;
          let _424_v18;
          _424_v18 = _dafny.Seq.of(function () {
            let _coll15 = new _dafny.Map();
            for (const _compr_15 of _dafny.IntegerRange(new BigNumber(-16), new BigNumber(-893))) {
              let _425_v17 = _compr_15;
              if (((new BigNumber(-16)).isLessThanOrEqualTo(_425_v17)) && ((_425_v17).isLessThan(new BigNumber(-893)))) {
                _coll15.push([(_425_v17).multipliedBy(new BigNumber(507)),(_this).f31]);
              }
            }
            return _coll15;
          }());
          let _426_v19;
          _426_v19 = _dafny.MultiSet.fromElements(_411_v13);
          let _427_v20;
          _427_v20 = _dafny.Set.fromElements(_410_i2, new BigNumber((_426_v19).cardinality()), (_this).f32);
          let _428_v21;
          let _nw71 = new _module.C0();
          _nw71.__ctor((_427_v20).contains(new BigNumber((_424_v18).length)));
          _428_v21 = _nw71;
        } else if (_source6.is_DC9) {
          let _429___mcc_h3 = (_source6).cf12;
          let _430_cf12 = _429___mcc_h3;
          r0 = (((_413_v15) && (_413_v15)) ? ((_399_v3)[_module.__default.safeIndex(new BigNumber(373), new BigNumber((_399_v3).length))]) : (_410_i2));
          let _431_v22;
          let _nw72 = Array((new BigNumber(12)).toNumber()).fill(false);
          _431_v22 = _nw72;
          let _index65 = _module.__default.safeIndex(new BigNumber(865), new BigNumber((_431_v22).length));
          (_431_v22)[_index65] = _413_v15;
          let _432_v23;
          _432_v23 = _dafny.Map.Empty.slice().updateUnsafe(_413_v15,(_399_v3)[_module.__default.safeIndex(new BigNumber(373), new BigNumber((_399_v3).length))]);
          let _index66 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_399_v3).length));
          let _index67 = _module.__default.safeIndex(new BigNumber(865), new BigNumber((_431_v22).length));
          let _index68 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_399_v3).length));
          let _rhs63 = (_this).f32;
          let _rhs64 = _411_v13;
          let _rhs65 = ((!(_432_v23).contains(_412_v14)) ? (_411_v13) : ((_this).f31));
          let _rhs66 = (_this).f32;
          let _rhs67 = (_this).f31;
          let _lhs34 = _399_v3;
          let _lhs35 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_399_v3).length));
          let _lhs36 = globalState;
          let _lhs37 = _431_v22;
          let _lhs38 = _module.__default.safeIndex(new BigNumber(865), new BigNumber((_431_v22).length));
          let _lhs39 = _399_v3;
          let _lhs40 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_399_v3).length));
          _lhs34[_lhs35] = _rhs63;
          _lhs36.f16 = _rhs64;
          _lhs37[_lhs38] = _rhs65;
          _lhs39[_lhs40] = _rhs66;
          _411_v13 = _rhs67;
          let _433_v24;
          let _nw73 = Array((new BigNumber(6)).toNumber());
          _nw73[(_dafny.ZERO).toNumber()] = _401_v5;
          _nw73[(_dafny.ONE).toNumber()] = _401_v5;
          _nw73[(new BigNumber(2)).toNumber()] = _401_v5;
          _nw73[(new BigNumber(3)).toNumber()] = _401_v5;
          _nw73[(new BigNumber(4)).toNumber()] = _dafny.Seq.of(_411_v13);
          _nw73[(new BigNumber(5)).toNumber()] = _401_v5;
          _433_v24 = _nw73;
          let _index69 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_433_v24).length));
          (_433_v24)[_index69] = _dafny.Seq.Concat(_401_v5, _401_v5);
          let _index70 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_433_v24).length));
          let _rhs68 = (_399_v3)[_module.__default.safeIndex(new BigNumber(373), new BigNumber((_399_v3).length))];
          let _rhs69 = _401_v5;
          let _rhs70 = _412_v14;
          let _rhs71 = _413_v15;
          let _rhs72 = _module.__default.safeDivisionInt((_this).f32, _410_i2);
          let _lhs41 = _433_v24;
          let _lhs42 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_433_v24).length));
          let _lhs43 = globalState;
          r0 = _rhs68;
          _lhs41[_lhs42] = _rhs69;
          _lhs43.f16 = _rhs70;
          _411_v13 = _rhs71;
          r0 = _rhs72;
          let _434_v25;
          let _435_v26;
          let _436_v27;
          let _out17;
          let _out18;
          let _out19;
          let _outcollector7 = (_this).m6(_412_v14, globalState);
          _out17 = _outcollector7[0];
          _out18 = _outcollector7[1];
          _out19 = _outcollector7[2];
          _434_v25 = _out17;
          _435_v26 = _out18;
          _436_v27 = _out19;
        } else {
          let _437___mcc_h4 = (_source6).cf16;
          let _438_cf16 = _437___mcc_h4;
          let _439_v28;
          _439_v28 = _dafny.Set.fromElements((_399_v3)[_module.__default.safeIndex(new BigNumber(373), new BigNumber((_399_v3).length))]);
          let _440_v29;
          let _nw74 = Array((new BigNumber(20)).toNumber()).fill(false);
          _440_v29 = _nw74;
          let _441_v30;
          _441_v30 = _module.D5.create_DC14(_439_v28, _410_i2, true, _440_v29);
          r0 = (_441_v30).dtor_cf20;
          let _index71 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_399_v3).length));
          (_399_v3)[_index71] = _410_i2;
          let _index72 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_399_v3).length));
          (_399_v3)[_index72] = (_dafny.ZERO).minus((_this).f32);
          let _442_v31;
          let _nw75 = new _module.C0();
          _nw75.__ctor(!(_411_v13));
          _442_v31 = _nw75;
        }
        let _index73 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_399_v3).length));
        (_399_v3)[_index73] = (_399_v3)[_module.__default.safeIndex(new BigNumber(373), new BigNumber((_399_v3).length))];
        if (_411_v13) {
          _409_v12 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(468)), ((_443_v4) => function (_444_i3) {
            return _443_v4;
          })(_400_v4));
          r0 = (_this).f32;
          let _445_v32;
          _445_v32 = _dafny.Set.fromElements(_400_v4);
          _445_v32 = _445_v32;
          let _446_v33;
          _446_v33 = _module.D0.create_DC1(_400_v4, (_this).f31, false);
          let _pat_let_tv5 = _400_v4;
          let _pat_let_tv6 = _399_v3;
          let _pat_let_tv7 = _399_v3;
          let _pat_let_tv8 = _409_v12;
          _446_v33 = function (_pat_let11_0) {
            return function (_447_dt__update__tmp_h1) {
              return function (_pat_let12_0) {
                return function (_448_dt__update_hcf1_h0) {
                  return function (_pat_let13_0) {
                    return function (_449_dt__update_hcf3_h0) {
                      return _module.D0.create_DC1(_448_dt__update_hcf1_h0, (_447_dt__update__tmp_h1).dtor_cf2, _449_dt__update_hcf3_h0);
                    }(_pat_let13_0);
                  }(((_pat_let_tv7)[_module.__default.safeIndex(new BigNumber(373), new BigNumber((_pat_let_tv6).length))]).isLessThan(new BigNumber((_pat_let_tv8).length)));
                }(_pat_let12_0);
              }(_pat_let_tv5);
            }(_pat_let11_0);
          }(_446_v33);
          let _450_v34;
          let _nw76 = new _module.C0();
          _nw76.__ctor(_413_v15);
          _450_v34 = _nw76;
        } else {
          let _451_v35;
          _451_v35 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_395_v0).cardinality()),(_this).f31);
          let _452_v36;
          let _init6 = ((_453_i2) => function (_454_i4) {
            return _dafny.Set.fromElements(_453_i2, _453_i2);
          })(_410_i2);
          let _nw77 = Array((new BigNumber(14)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw77.length); _i0_6++) {
            _nw77[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _452_v36 = _nw77;
          let _455_v37;
          _455_v37 = _module.D6.create_DC17(_413_v15, _452_v36, _410_i2);
          let _456_v38;
          _456_v38 = _dafny.Map.Empty.slice().updateUnsafe(_451_v35,_dafny.Seq.of((_401_v5)[_module.__default.safeIndex((_this).f32, new BigNumber((_401_v5).length))], (_455_v37).dtor_cf25, _411_v13));
          let _457_v39;
          let _init7 = function (_458_i5) {
            return new _dafny.CodePoint('r'.codePointAt(0));
          };
          let _nw78 = Array((new BigNumber(5)).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw78.length); _i0_7++) {
            _nw78[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _457_v39 = _nw78;
          let _index74 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_457_v39).length));
          (_457_v39)[_index74] = _400_v4;
          let _459_v40;
          _459_v40 = _module.D0.create_DC2((_this).f31);
          let _460_v41;
          _460_v41 = _dafny.Map.Empty.slice().updateUnsafe(_459_v40,_400_v4);
          let _index75 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_457_v39).length));
          let _rhs73 = _module.__default.fm20(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("gswdyire"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(564)), ((_461_v4) => function (_462_i6) {
            return _461_v4;
          })(_400_v4))), _395_v0, (_this).f31, (_this).f32, globalState);
          let _rhs74 = (_412_v14) === (!(true) || (_411_v13));
          let _rhs75 = (_456_v38).update(_451_v35, _401_v5);
          let _rhs76 = (((_460_v41).contains(_459_v40)) ? ((_460_v41).get(_459_v40)) : (new _dafny.CodePoint('o'.codePointAt(0))));
          let _lhs44 = _457_v39;
          let _lhs45 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_457_v39).length));
          _400_v4 = _rhs73;
          _411_v13 = _rhs74;
          _456_v38 = _rhs75;
          _lhs44[_lhs45] = _rhs76;
          let _rhs77 = _400_v4;
          let _rhs78 = _411_v13;
          let _lhs46 = globalState;
          _400_v4 = _rhs77;
          _lhs46.f16 = _rhs78;
          let _pat_let_tv9 = _452_v36;
          _452_v36 = (((true) ? (function (_pat_let14_0) {
            return function (_463_dt__update__tmp_h2) {
              return function (_pat_let15_0) {
                return function (_464_dt__update_hcf26_h0) {
                  return _module.D6.create_DC17((_463_dt__update__tmp_h2).dtor_cf25, _464_dt__update_hcf26_h0, (_463_dt__update__tmp_h2).dtor_cf27);
                }(_pat_let15_0);
              }(_pat_let_tv9);
            }(_pat_let14_0);
          }(_455_v37)) : (_455_v37))).dtor_cf26;
          let _465_v42;
          _465_v42 = _module.D3.create_DC10(_dafny.Seq.UnicodeFromString("rpccqjlr"), new BigNumber((_dafny.Seq.update(_409_v12, _module.__default.safeIndex(new BigNumber(847), new BigNumber((_409_v12).length)), (_457_v39)[_module.__default.safeIndex(new BigNumber(392), new BigNumber((_457_v39).length))])).length), _411_v13);
          _409_v12 = (_465_v42).dtor_cf13;
          let _index76 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_399_v3).length));
          let _rhs79 = _module.__default.fm21(_395_v0, _412_v14, !_dafny.areEqual(_dafny.Seq.update(_409_v12, _module.__default.safeIndex((_this).f32, new BigNumber((_409_v12).length)), _module.__default.fm20((_this).fm16(_396_v1, _413_v15, _411_v13, _412_v14, globalState), _395_v0, (_this).f31, new BigNumber((_395_v0).cardinality()), globalState)), _409_v12), globalState);
          let _rhs80 = _410_i2;
          let _lhs47 = _399_v3;
          let _lhs48 = _module.__default.safeIndex(new BigNumber(373), new BigNumber((_399_v3).length));
          _451_v35 = _rhs79;
          _lhs47[_lhs48] = _rhs80;
        }
      }
      (globalState).f28 = ((_399_v3)[_module.__default.safeIndex(new BigNumber(373), new BigNumber((_399_v3).length))]).isEqualTo(((_this).f32).minus(new BigNumber(141)));
      r0 = (_this).f32;
      r1 = ((_this).f32).isEqualTo(_module.__default.fm18((_this).f31, !((_this).f31), globalState));
      let _466_v43;
      _466_v43 = _dafny.Set.fromElements((_this).f32, new BigNumber((_409_v12).length));
      let _467_v44;
      _467_v44 = _dafny.MultiSet.fromElements(_466_v43);
      let _468_v45;
      _468_v45 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f32),(_this).f31);
      let _469_v46;
      _469_v46 = _dafny.Map.Empty.slice().updateUnsafe((_this).f31,_468_v45);
      r2 = (((_this).f31) ? ((_dafny.MultiSet.FromArray(_module.__default.fm22((_this).f32, globalState))).Difference(_467_v44)) : ((_467_v44).Union((_467_v44).update(_466_v43, _module.__default.abs(new BigNumber((_module.__default.fm23(_469_v46, _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_409_v12).length)), _dafny.Seq.UnicodeFromString("lqta"), globalState)).length))))));
      r3 = (_this).f31;
      return [r0, r1, r2, r3];
    }
    m6(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = false;
      let _470_v0;
      let _nw79 = Array((new BigNumber(6)).toNumber());
      _nw79[(_dafny.ZERO).toNumber()] = (_this).f32;
      _nw79[(_dafny.ONE).toNumber()] = (_this).f32;
      _nw79[(new BigNumber(2)).toNumber()] = (_this).f32;
      _nw79[(new BigNumber(3)).toNumber()] = (_this).f32;
      _nw79[(new BigNumber(4)).toNumber()] = (_this).f32;
      _nw79[(new BigNumber(5)).toNumber()] = (_this).f32;
      _470_v0 = _nw79;
      (globalState).f24 = _470_v0;
      let _471_v1;
      _471_v1 = _dafny.Seq.UnicodeFromString("jymesegv");
      let _472_v2;
      _472_v2 = _dafny.MultiSet.fromElements((_this).f32);
      let _473_v3;
      _473_v3 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f32),p0);
      let _474_v4;
      _474_v4 = _dafny.Seq.of(_module.__default.fm21(_472_v2, (_this).f31, (_this).f31, globalState), _473_v3, _473_v3);
      let _475_v5;
      _475_v5 = _module.D3.create_DC10(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(162)), function (_476_i0) {
  return new _dafny.CodePoint('q'.codePointAt(0));
}), _471_v1), new BigNumber((_dafny.MultiSet.FromArray(_474_v4)).cardinality()), (_this).f31);
      let _source7 = _475_v5;
      if (_source7.is_DC10) {
        let _477___mcc_h0 = (_source7).cf13;
        let _478___mcc_h1 = (_source7).cf14;
        let _479___mcc_h2 = (_source7).cf15;
        let _480_cf15 = _479___mcc_h2;
        let _481_cf14 = _478___mcc_h1;
        let _482_cf13 = _477___mcc_h0;
        _481_cf14 = _481_cf14;
        let _483_v6;
        _483_v6 = _dafny.Seq.of(_480_cf15, (_this).f31);
        let _484_v7;
        let _nw80 = Array((new BigNumber(17)).toNumber());
        _nw80[(_dafny.ZERO).toNumber()] = _483_v6;
        _nw80[(_dafny.ONE).toNumber()] = _483_v6;
        _nw80[(new BigNumber(2)).toNumber()] = _483_v6;
        _nw80[(new BigNumber(3)).toNumber()] = _483_v6;
        _nw80[(new BigNumber(4)).toNumber()] = _483_v6;
        _nw80[(new BigNumber(5)).toNumber()] = _483_v6;
        _nw80[(new BigNumber(6)).toNumber()] = _483_v6;
        _nw80[(new BigNumber(7)).toNumber()] = _483_v6;
        _nw80[(new BigNumber(8)).toNumber()] = _483_v6;
        _nw80[(new BigNumber(9)).toNumber()] = _dafny.Seq.of(_480_cf15, false);
        _nw80[(new BigNumber(10)).toNumber()] = _483_v6;
        _nw80[(new BigNumber(11)).toNumber()] = _483_v6;
        _nw80[(new BigNumber(12)).toNumber()] = _483_v6;
        _nw80[(new BigNumber(13)).toNumber()] = _483_v6;
        _nw80[(new BigNumber(14)).toNumber()] = _483_v6;
        _nw80[(new BigNumber(15)).toNumber()] = _483_v6;
        _nw80[(new BigNumber(16)).toNumber()] = _dafny.Seq.of(p0);
        _484_v7 = _nw80;
        let _485_v8;
        _485_v8 = _module.D0.create_DC0(_484_v7);
        let _486_v9;
        _486_v9 = _module.D0.create_DC4(_485_v8);
        let _487_v10;
        let _nw81 = Array((new BigNumber(5)).toNumber());
        _nw81[(_dafny.ZERO).toNumber()] = _486_v9;
        _nw81[(_dafny.ONE).toNumber()] = _486_v9;
        _nw81[(new BigNumber(2)).toNumber()] = _486_v9;
        _nw81[(new BigNumber(3)).toNumber()] = _486_v9;
        _nw81[(new BigNumber(4)).toNumber()] = _486_v9;
        _487_v10 = _nw81;
        _487_v10 = _487_v10;
        let _488_v11;
        _488_v11 = _module.D0.create_DC2((_this).f31);
        let _source8 = _488_v11;
        if (_source8.is_DC1) {
          let _489___mcc_h5 = (_source8).cf1;
          let _490___mcc_h6 = (_source8).cf2;
          let _491___mcc_h7 = (_source8).cf3;
          let _492_cf3 = _491___mcc_h7;
          let _493_cf2 = _490___mcc_h6;
          let _494_cf1 = _489___mcc_h5;
          let _495_v12;
          _495_v12 = _dafny.Map.Empty.slice().updateUnsafe(_494_cf1,(_this).fm17(_494_cf1, globalState));
          _495_v12 = _495_v12;
          let _496_v13;
          let _nw82 = Array((new BigNumber(11)).toNumber()).fill(_module.D3.Default());
          _496_v13 = _nw82;
          let _497_v14;
          _497_v14 = _dafny.Seq.of(_496_v13);
          let _498_v15;
          let _nw83 = Array((new BigNumber(4)).toNumber());
          _nw83[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of(_496_v13), _dafny.Seq.of(_496_v13, _496_v13)), _module.__default.safeIndex(_481_cf14, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_496_v13), _dafny.Seq.of(_496_v13, _496_v13))).length)), _496_v13);
          _nw83[(_dafny.ONE).toNumber()] = _497_v14;
          _nw83[(new BigNumber(2)).toNumber()] = _497_v14;
          _nw83[(new BigNumber(3)).toNumber()] = _497_v14;
          _498_v15 = _nw83;
          let _index77 = _module.__default.safeIndex(new BigNumber(97), new BigNumber((_498_v15).length));
          (_498_v15)[_index77] = _497_v14;
          let _index78 = _module.__default.safeIndex(new BigNumber(97), new BigNumber((_498_v15).length));
          (_498_v15)[_index78] = _497_v14;
          let _499_v16;
          _499_v16 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(512),new BigNumber((_471_v1).length));
          let _500_v17;
          _500_v17 = _dafny.Set.fromElements(_499_v16);
          let _501_v18;
          _501_v18 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_482_cf13).length)).isEqualTo((_this).f32),(_500_v17).Union(_500_v17));
          let _502_v21;
          _502_v21 = _dafny.Seq.of(function () {
            let _coll16 = new _dafny.Set();
            for (const _compr_16 of (function () {
              let _coll17 = new _dafny.Set();
              for (const _compr_17 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(682)), ((_503_cf14, _504_cf13) => function (_505_i1) {
                return _dafny.Map.Empty.slice().updateUnsafe(_503_cf14,new BigNumber((_504_cf13).length));
              })(_481_cf14, _482_cf13))).Elements) {
                let _506_v19 = _compr_17;
                if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(682)), ((_507_cf14, _508_cf13) => function (_505_i1) {
                  return _dafny.Map.Empty.slice().updateUnsafe(_507_cf14,new BigNumber((_508_cf13).length));
                })(_481_cf14, _482_cf13)), _506_v19)) {
                  _coll17.add(_506_v19);
                }
              }
              return _coll17;
            }()).Elements) {
              let _509_v20 = _compr_16;
              if ((function () {
                let _coll18 = new _dafny.Set();
                for (const _compr_18 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(682)), ((_510_cf14, _511_cf13) => function (_505_i1) {
                  return _dafny.Map.Empty.slice().updateUnsafe(_510_cf14,new BigNumber((_511_cf13).length));
                })(_481_cf14, _482_cf13))).Elements) {
                  let _512_v19 = _compr_18;
                  if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(682)), ((_513_cf14, _514_cf13) => function (_505_i1) {
                    return _dafny.Map.Empty.slice().updateUnsafe(_513_cf14,new BigNumber((_514_cf13).length));
                  })(_481_cf14, _482_cf13)), _512_v19)) {
                    _coll18.add(_512_v19);
                  }
                }
                return _coll18;
              }()).contains(_509_v20)) {
                _coll16.add(_509_v20);
              }
            }
            return _coll16;
          }(), (((_501_v18).contains(p0)) ? ((_501_v18).get(p0)) : (_500_v17)));
          _501_v18 = (_501_v18).update(p0, (_502_v21)[_module.__default.safeIndex((_this).f32, new BigNumber((_502_v21).length))]);
          _482_cf13 = _482_cf13;
        } else if (_source8.is_DC2) {
          let _515___mcc_h8 = (_source8).cf4;
          let _516_cf4 = _515___mcc_h8;
          let _517_v22;
          _517_v22 = new _dafny.CodePoint('w'.codePointAt(0));
          let _518_v23;
          _518_v23 = _dafny.MultiSet.fromElements(_517_v22, _module.__default.fm20(_471_v1, _dafny.MultiSet.fromElements((_this).f32), (_this).f31, (_this).f32, globalState));
          _module.__default.m0((((_518_v23).contains(_517_v22)) ? ((_518_v23).get(_517_v22)) : ((_this).f32)), globalState);
          let _519_v24;
          _519_v24 = _dafny.Map.Empty.slice().updateUnsafe((_this).f31,_481_cf14);
          let _520_v25;
          _520_v25 = _dafny.Map.Empty.slice().updateUnsafe((((_519_v24).contains(p0)) ? ((_519_v24).get(p0)) : ((_this).f32)),_module.__default.fm18(p0, (_this).f31, globalState));
          _520_v25 = (_520_v25).update((new BigNumber((_471_v1).length)).plus(_481_cf14), _481_cf14);
          let _521_v26;
          _521_v26 = _dafny.MultiSet.fromElements((_module.D3.create_DC10(_dafny.Seq.UnicodeFromString("hiecrh"), (_this).f32, _516_cf4)).dtor_cf13);
          let _522_v27;
          _522_v27 = _dafny.Set.fromElements(_516_cf4);
          let _523_v28;
          _523_v28 = _dafny.Map.Empty.slice().updateUnsafe(_521_v26,(_522_v27).Intersect(_522_v27));
          _523_v28 = (_523_v28).update(_dafny.MultiSet.fromElements(_471_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(735)), ((_524_v22) => function (_525_i2) {
            return _524_v22;
          })(_517_v22))), _522_v27);
          let _526_v29;
          let _nw84 = new _module.C0();
          _nw84.__ctor(((_this).f31) === ((_this).fm17(_517_v22, globalState)));
          _526_v29 = _nw84;
        } else if (_source8.is_DC3) {
          let _527___mcc_h9 = (_source8).cf5;
          let _528_cf5 = _527___mcc_h9;
          let _529_v30;
          _529_v30 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_482_cf13).length)).minus((_this).f32),new BigNumber((_472_v2).cardinality()));
          _529_v30 = _529_v30;
          let _530_v31;
          let _nw85 = Array((new BigNumber(9)).toNumber());
          _nw85[(_dafny.ZERO).toNumber()] = !(!((_this).f31)) || ((_this).f31);
          _nw85[(_dafny.ONE).toNumber()] = !(false) || ((_this).f31);
          _nw85[(new BigNumber(2)).toNumber()] = !(p0);
          _nw85[(new BigNumber(3)).toNumber()] = (_this).f31;
          _nw85[(new BigNumber(4)).toNumber()] = (p0) === (_480_cf15);
          _nw85[(new BigNumber(5)).toNumber()] = (_481_cf14).isLessThan((_dafny.ZERO).minus((_this).f32));
          _nw85[(new BigNumber(6)).toNumber()] = p0;
          _nw85[(new BigNumber(7)).toNumber()] = (_this).f31;
          _nw85[(new BigNumber(8)).toNumber()] = !(_480_cf15);
          _530_v31 = _nw85;
          let _index79 = _module.__default.safeIndex(new BigNumber(865), new BigNumber((_530_v31).length));
          (_530_v31)[_index79] = _480_cf15;
          let _531_v32;
          _531_v32 = _dafny.Map.Empty.slice().updateUnsafe(_486_v9,_481_cf14);
          let _index80 = _module.__default.safeIndex(new BigNumber(865), new BigNumber((_530_v31).length));
          (_530_v31)[_index80] = (_483_v6)[_module.__default.safeIndex((((_531_v32).contains(_486_v9)) ? ((_531_v32).get(_486_v9)) : (new BigNumber(759))), new BigNumber((_483_v6).length))];
          let _index81 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_470_v0).length));
          (_470_v0)[_index81] = new BigNumber((_dafny.Seq.Concat(_482_cf13, _482_cf13)).length);
          let _532_v33;
          _532_v33 = _dafny.Seq.of(_482_cf13, _dafny.Seq.UnicodeFromString("ufldtem"), _482_cf13);
          let _533_v34;
          let _nw86 = Array((new BigNumber(21)).toNumber());
          _nw86[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("eohrqs");
          _nw86[(_dafny.ONE).toNumber()] = _471_v1;
          _nw86[(new BigNumber(2)).toNumber()] = _471_v1;
          _nw86[(new BigNumber(3)).toNumber()] = _471_v1;
          _nw86[(new BigNumber(4)).toNumber()] = _482_cf13;
          _nw86[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat((_475_v5).dtor_cf13, _482_cf13);
          _nw86[(new BigNumber(6)).toNumber()] = _471_v1;
          _nw86[(new BigNumber(7)).toNumber()] = _dafny.Seq.UnicodeFromString("vet");
          _nw86[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_471_v1, _471_v1);
          _nw86[(new BigNumber(9)).toNumber()] = _dafny.Seq.UnicodeFromString("et");
          _nw86[(new BigNumber(10)).toNumber()] = _482_cf13;
          _nw86[(new BigNumber(11)).toNumber()] = (_532_v33)[_module.__default.safeIndex(new BigNumber(407), new BigNumber((_532_v33).length))];
          _nw86[(new BigNumber(12)).toNumber()] = _482_cf13;
          _nw86[(new BigNumber(13)).toNumber()] = _471_v1;
          _nw86[(new BigNumber(14)).toNumber()] = _471_v1;
          _nw86[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_471_v1, _471_v1);
          _nw86[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_482_cf13, _dafny.Seq.UnicodeFromString("fvmafgib"));
          _nw86[(new BigNumber(17)).toNumber()] = _482_cf13;
          _nw86[(new BigNumber(18)).toNumber()] = _471_v1;
          _nw86[(new BigNumber(19)).toNumber()] = (_475_v5).dtor_cf13;
          _nw86[(new BigNumber(20)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(902)), ((_534_cf5) => function (_535_i3) {
            return _534_cf5;
          })(_528_cf5));
          _533_v34 = _nw86;
          let _index82 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_533_v34).length));
          (_533_v34)[_index82] = _471_v1;
          let _index83 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_470_v0).length));
          let _index84 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_533_v34).length));
          let _rhs81 = (_this).f32;
          let _rhs82 = (_this).f32;
          let _rhs83 = (_481_cf14).multipliedBy(new BigNumber(-720));
          let _rhs84 = _482_cf13;
          let _lhs49 = _470_v0;
          let _lhs50 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_470_v0).length));
          let _lhs51 = _533_v34;
          let _lhs52 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_533_v34).length));
          _481_cf14 = _rhs81;
          _481_cf14 = _rhs82;
          _lhs49[_lhs50] = _rhs83;
          _lhs51[_lhs52] = _rhs84;
          let _536_v35;
          _536_v35 = _module.D2.create_DC6((_this).f32);
          let _537_v36;
          _537_v36 = _dafny.Set.fromElements(_536_v35);
          let _538_v37;
          _538_v37 = _dafny.Map.Empty.slice().updateUnsafe(_481_cf14,_528_cf5);
          let _539_v38;
          let _nw87 = Array((new BigNumber(13)).toNumber());
          _nw87[(_dafny.ZERO).toNumber()] = _528_cf5;
          _nw87[(_dafny.ONE).toNumber()] = new _dafny.CodePoint('u'.codePointAt(0));
          _nw87[(new BigNumber(2)).toNumber()] = _528_cf5;
          _nw87[(new BigNumber(3)).toNumber()] = _528_cf5;
          _nw87[(new BigNumber(4)).toNumber()] = (((_538_v37).contains((_470_v0)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_470_v0).length))])) ? ((_538_v37).get((_470_v0)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_470_v0).length))])) : (_528_cf5));
          _nw87[(new BigNumber(5)).toNumber()] = _528_cf5;
          _nw87[(new BigNumber(6)).toNumber()] = _528_cf5;
          _nw87[(new BigNumber(7)).toNumber()] = _528_cf5;
          _nw87[(new BigNumber(8)).toNumber()] = _528_cf5;
          _nw87[(new BigNumber(9)).toNumber()] = _528_cf5;
          _nw87[(new BigNumber(10)).toNumber()] = new _dafny.CodePoint('v'.codePointAt(0));
          _nw87[(new BigNumber(11)).toNumber()] = _528_cf5;
          _nw87[(new BigNumber(12)).toNumber()] = _528_cf5;
          _539_v38 = _nw87;
          let _540_v39;
          _540_v39 = _dafny.Map.Empty.slice().updateUnsafe(_537_v36,_539_v38);
          _540_v39 = (_540_v39).update(_537_v36, _539_v38);
        } else if (_source8.is_DC0) {
          let _541___mcc_h10 = (_source8).cf0;
          let _542_cf0 = _541___mcc_h10;
          _481_cf14 = (_this).f32;
          _481_cf14 = _481_cf14;
          let _543_v40;
          _543_v40 = _dafny.Map.Empty.slice().updateUnsafe(_480_cf15,p0);
          let _544_v41;
          _544_v41 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_543_v40).update((_this).f31, (((_543_v40).contains((_module.D0.create_DC2((_this).f31)).dtor_cf4)) ? ((_543_v40).get((_module.D0.create_DC2((_this).f31)).dtor_cf4)) : (p0)))).length),_module.__default.fm24(_480_cf15, _480_cf15, new BigNumber(681), globalState));
          let _545_v42;
          _545_v42 = _dafny.Map.Empty.slice().updateUnsafe((_this).f31,new BigNumber(((((_544_v41).contains(_module.__default.fm18(_480_cf15, _480_cf15, globalState))) ? ((_544_v41).get(_module.__default.fm18(_480_cf15, _480_cf15, globalState))) : (_472_v2))).cardinality()));
          _545_v42 = (_545_v42).update(p0, (_this).f32);
          let _546_v43;
          let _nw88 = Array((new BigNumber(16)).toNumber());
          _546_v43 = _nw88;
          let _547_v44;
          let _nw89 = new _module.C0();
          _nw89.__ctor(p0);
          _547_v44 = _nw89;
          let _index85 = _module.__default.safeIndex(new BigNumber(339), new BigNumber((_546_v43).length));
          (_546_v43)[_index85] = _547_v44;
          let _index86 = _module.__default.safeIndex(new BigNumber(339), new BigNumber((_546_v43).length));
          (_546_v43)[_index86] = _547_v44;
        } else {
          let _548___mcc_h11 = (_source8).cf6;
          let _549_cf6 = _548___mcc_h11;
          let _550_v45;
          _550_v45 = _dafny.MultiSet.fromElements(p0);
          let _551_v46;
          _551_v46 = _dafny.Map.Empty.slice().updateUnsafe((_483_v6)[_module.__default.safeIndex(new BigNumber((_550_v45).cardinality()), new BigNumber((_483_v6).length))],(_this).f31);
          let _552_v47;
          let _nw90 = Array((new BigNumber(4)).toNumber());
          _nw90[(_dafny.ZERO).toNumber()] = _480_cf15;
          _nw90[(_dafny.ONE).toNumber()] = _480_cf15;
          _nw90[(new BigNumber(2)).toNumber()] = (((_551_v46).contains(true)) ? ((_551_v46).get(true)) : ((_488_v11).dtor_cf4));
          _nw90[(new BigNumber(3)).toNumber()] = (_this).f31;
          _552_v47 = _nw90;
          let _553_v48;
          _553_v48 = _dafny.Map.Empty.slice().updateUnsafe(_552_v47,(new BigNumber(693)).isLessThanOrEqualTo(_module.__default.fm18((_483_v6)[_module.__default.safeIndex(new BigNumber(494), new BigNumber((_483_v6).length))], _480_cf15, globalState)));
          _553_v48 = (_553_v48).update(_552_v47, _480_cf15);
          let _554_v49;
          let _nw91 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Map.Empty);
          _554_v49 = _nw91;
          let _555_v50;
          _555_v50 = _dafny.Seq.of(_554_v49);
          let _556_v51;
          _556_v51 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm20(_471_v1, _472_v2, !(true), (_this).f32, globalState),new BigNumber(-279));
          let _557_v52;
          _557_v52 = new _dafny.CodePoint('u'.codePointAt(0));
          let _rhs85 = (_555_v50)[_module.__default.safeIndex(new BigNumber((_556_v51).length), new BigNumber((_555_v50).length))];
          let _rhs86 = _551_v46;
          let _rhs87 = ((_550_v45).Difference(_550_v45)).IsProperSubsetOf((_550_v45).Difference(_550_v45));
          let _rhs88 = (_this).fm17(_557_v52, globalState);
          let _lhs53 = globalState;
          _554_v49 = _rhs85;
          _551_v46 = _rhs86;
          _lhs53.f28 = _rhs87;
          r0 = _rhs88;
          _481_cf14 = new BigNumber(-842);
          _480_cf15 = (_475_v5).dtor_cf15;
        }
        let _558_v53;
        _558_v53 = _dafny.Seq.of(_484_v7, _484_v7);
        let _559_v54;
        _559_v54 = _module.D0.create_DC0((_558_v53)[_module.__default.safeIndex(_481_cf14, new BigNumber((_558_v53).length))]);
        let _560_v55;
        _560_v55 = _dafny.Map.Empty.slice().updateUnsafe(_559_v54,(_this).f31);
        _560_v55 = _560_v55;
      } else if (_source7.is_DC9) {
        let _561___mcc_h3 = (_source7).cf12;
        let _562_cf12 = _561___mcc_h3;
        _module.__default.m0(((_this).f32).plus((_dafny.ZERO).minus((_this).f32)), globalState);
        let _563_v56;
        _563_v56 = _module.__default.fm25(false, new BigNumber(-932), (_this).f32, (_this).f31, globalState);
        let _564_v57;
        _564_v57 = _dafny.MultiSet.fromElements((_this).f31, (_this).f31, p0);
        let _565_v58;
        _565_v58 = _module.D2.create_DC6(new BigNumber((_564_v57).cardinality()));
        let _566_v59;
        _566_v59 = _dafny.Set.fromElements(_565_v58, _565_v58, _565_v58, _565_v58);
        _563_v56 = _566_v59;
        (globalState).f16 = p0;
        let _567_v60;
        _567_v60 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f32);
        let _568_v61;
        _568_v61 = _dafny.Seq.of((_this).f32, new BigNumber((_567_v60).length), (_this).f32, (_this).f32, (_this).f32);
        let _569_v62;
        _569_v62 = _dafny.Set.fromElements((_this).f32, new BigNumber((_568_v61).length), (_this).f32, (_this).f32, new BigNumber((_472_v2).cardinality()));
        _569_v62 = (_569_v62).Difference(_dafny.Set.fromElements((_this).f32, (_this).f32));
      } else {
        let _570___mcc_h4 = (_source7).cf16;
        let _571_cf16 = _570___mcc_h4;
        let _572_v63;
        _572_v63 = new BigNumber(592);
        _572_v63 = _572_v63;
        let _573_v64;
        _573_v64 = _dafny.Seq.of(p0);
        let _574_v65;
        _574_v65 = _dafny.Set.fromElements(new BigNumber((_573_v64).length));
        let _575_v66;
        _575_v66 = _dafny.Map.Empty.slice().updateUnsafe(_470_v0,new BigNumber(((_574_v65).Intersect(_574_v65)).length));
        let _576_v67;
        let _nw92 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _576_v67 = _nw92;
        let _577_v68;
        _577_v68 = new _dafny.CodePoint('d'.codePointAt(0));
        let _index87 = _module.__default.safeIndex(new BigNumber(803), new BigNumber((_576_v67).length));
        (_576_v67)[_index87] = _dafny.Seq.of(_577_v68);
        let _578_v69;
        _578_v69 = _dafny.Seq.of((_this).f32, (_this).f32);
        let _index88 = _module.__default.safeIndex(new BigNumber(803), new BigNumber((_576_v67).length));
        let _rhs89 = ((_575_v66).Merge(_575_v66)).Merge(((_575_v66).update(_470_v0, (_this).f32)).Merge(_575_v66));
        let _rhs90 = (_572_v63).isLessThanOrEqualTo((_578_v69)[_module.__default.safeIndex(new BigNumber(873), new BigNumber((_578_v69).length))]);
        let _rhs91 = _471_v1;
        let _rhs92 = p0;
        let _lhs54 = _576_v67;
        let _lhs55 = _module.__default.safeIndex(new BigNumber(803), new BigNumber((_576_v67).length));
        _575_v66 = _rhs89;
        r2 = _rhs90;
        _lhs54[_lhs55] = _rhs91;
        r2 = _rhs92;
        let _579_v70;
        _579_v70 = _dafny.Map.Empty.slice().updateUnsafe((_this).f31,(_dafny.ZERO).minus((_this).f32));
        let _index89 = _module.__default.safeIndex(new BigNumber(531), new BigNumber((_470_v0).length));
        (_470_v0)[_index89] = (((_this).f31) ? (_572_v63) : (new BigNumber((_579_v70).length)));
        let _index90 = _module.__default.safeIndex(new BigNumber(531), new BigNumber((_470_v0).length));
        (_470_v0)[_index90] = (_this).f32;
        _module.__default.m0((_470_v0)[_module.__default.safeIndex(new BigNumber(531), new BigNumber((_470_v0).length))], globalState);
      }
      let _580_v71;
      _580_v71 = _dafny.Set.fromElements(_470_v0, _470_v0);
      let _581_v72;
      _581_v72 = _dafny.Map.Empty.slice().updateUnsafe(_580_v71,(_this).f32);
      _581_v72 = (_581_v72).update(_580_v71, (_this).f32);
      let _582_v73;
      _582_v73 = new BigNumber(355);
      _582_v73 = (_582_v73).minus(new BigNumber((_dafny.Seq.of(p0, p0)).length));
      let _583_v74;
      _583_v74 = new _dafny.CodePoint('b'.codePointAt(0));
      let _584_i4;
      _584_i4 = _dafny.ZERO;
      L6: {
        while (!(!(p0) || ((_this).fm17(_583_v74, globalState)))) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_584_i4)) {
              break L6;
            }
            _584_i4 = (_584_i4).plus(_dafny.ONE);
            _471_v1 = _471_v1;
            let _585_v75;
            let _init8 = ((_586_v74) => function (_587_i5) {
              return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-543)), ((_588_v74) => function (_589_i6) {
                return _588_v74;
              })(_586_v74));
            })(_583_v74);
            let _nw93 = Array((_dafny.ONE).toNumber());
            for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw93.length); _i0_8++) {
              _nw93[_i0_8] = _init8(new BigNumber(_i0_8));
            }
            _585_v75 = _nw93;
            let _590_v76;
            let _nw94 = Array((new BigNumber(6)).toNumber());
            _nw94[(_dafny.ZERO).toNumber()] = _585_v75;
            _nw94[(_dafny.ONE).toNumber()] = _585_v75;
            _nw94[(new BigNumber(2)).toNumber()] = _585_v75;
            _nw94[(new BigNumber(3)).toNumber()] = _585_v75;
            _nw94[(new BigNumber(4)).toNumber()] = _585_v75;
            _nw94[(new BigNumber(5)).toNumber()] = _585_v75;
            _590_v76 = _nw94;
            let _index91 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_590_v76).length));
            (_590_v76)[_index91] = _585_v75;
            let _index92 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_590_v76).length));
            (_590_v76)[_index92] = _585_v75;
            let _591_v77;
            let _nw95 = new _module.C0();
            _nw95.__ctor(_dafny.areEqual(_471_v1, _471_v1));
            _591_v77 = _nw95;
            let _rhs93 = (_this).f32;
            let _rhs94 = !((_this).f31);
            let _lhs56 = globalState;
            _582_v73 = _rhs93;
            _lhs56.f16 = _rhs94;
          }
        }
      }
      let _592_v78;
      _592_v78 = _dafny.Seq.of(_582_v73);
      let _593_v79;
      _593_v79 = _dafny.Seq.of(p0, _dafny.Seq.IsProperPrefixOf(_592_v78, _592_v78));
      _593_v79 = _593_v79;
      let _594_v80;
      _594_v80 = _dafny.Map.Empty.slice().updateUnsafe(p0,_582_v73);
      let _595_v81;
      _595_v81 = _dafny.Seq.of(_594_v80);
      r0 = !_dafny.areEqual(_595_v81, _595_v81);
      r1 = true;
      r2 = p0;
      return [r0, r1, r2];
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f31 = false;
      this._f38 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f31() {
      let _this = this;
      return _this._f31;
    };
    __ctor(f38, f31) {
      let _this = this;
      (_this)._f38 = f38;
      (_this)._f31 = f31;
      return;
    }
    fm12(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_module.D2.create_DC6(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("my")).length))).length))).dtor_cf8;
    };
    fm13(p0, globalState) {
      let _this = this;
      return ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("xwst")).length))).minus(new BigNumber(84));
    };
    m4(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = [];
      let _596_v0;
      _596_v0 = _dafny.Seq.UnicodeFromString("fnnicgsqf");
      let _597_v1;
      _597_v1 = new _dafny.CodePoint('r'.codePointAt(0));
      let _598_v2;
      _598_v2 = _module.D0.create_DC1(_597_v1, (_this).f31, p2);
      let _599_v3;
      _599_v3 = _dafny.Map.Empty.slice().updateUnsafe((_598_v2).dtor_cf2,p1);
      let _hi3 = new BigNumber(((_599_v3).Merge(_599_v3)).length);
      for (let _600_i0 = (new BigNumber((_596_v0).length)).multipliedBy(p1); _600_i0.isLessThan(_hi3); _600_i0 = _600_i0.plus(_dafny.ONE)) {
        let _601_v4;
        let _nw96 = Array((new BigNumber(26)).toNumber()).fill(false);
        _601_v4 = _nw96;
        let _602_v5;
        _602_v5 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(910),_601_v4);
        let _603_v6;
        _603_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm14(_600_i0, _600_i0, p1, true, globalState)).length),p0);
        let _604_v7;
        _604_v7 = _dafny.Set.fromElements(p3, false);
        let _605_v8;
        _605_v8 = _dafny.Set.fromElements(new BigNumber((_604_v7).length), p0, new BigNumber(151), p0, p1);
        let _606_v9;
        _606_v9 = _module.D5.create_DC14(_605_v8, _600_i0, p2, _601_v4);
        let _607_v10;
        let _nw97 = Array((new BigNumber(23)).toNumber());
        _nw97[(_dafny.ZERO).toNumber()] = (((_602_v5).contains(new BigNumber(((_603_v6).update(p1, new BigNumber((_596_v0).length))).length))) ? ((_602_v5).get(new BigNumber(((_603_v6).update(p1, new BigNumber((_596_v0).length))).length))) : (_601_v4));
        _nw97[(_dafny.ONE).toNumber()] = _601_v4;
        _nw97[(new BigNumber(2)).toNumber()] = _601_v4;
        _nw97[(new BigNumber(3)).toNumber()] = (((_this).f31) ? (_601_v4) : (_601_v4));
        _nw97[(new BigNumber(4)).toNumber()] = _601_v4;
        _nw97[(new BigNumber(5)).toNumber()] = _601_v4;
        _nw97[(new BigNumber(6)).toNumber()] = _601_v4;
        _nw97[(new BigNumber(7)).toNumber()] = _601_v4;
        _nw97[(new BigNumber(8)).toNumber()] = _601_v4;
        _nw97[(new BigNumber(9)).toNumber()] = (_606_v9).dtor_cf22;
        _nw97[(new BigNumber(10)).toNumber()] = _601_v4;
        _nw97[(new BigNumber(11)).toNumber()] = _601_v4;
        _nw97[(new BigNumber(12)).toNumber()] = _601_v4;
        _nw97[(new BigNumber(13)).toNumber()] = _601_v4;
        _nw97[(new BigNumber(14)).toNumber()] = _601_v4;
        _nw97[(new BigNumber(15)).toNumber()] = _601_v4;
        _nw97[(new BigNumber(16)).toNumber()] = _601_v4;
        _nw97[(new BigNumber(17)).toNumber()] = _601_v4;
        _nw97[(new BigNumber(18)).toNumber()] = _601_v4;
        _nw97[(new BigNumber(19)).toNumber()] = _601_v4;
        _nw97[(new BigNumber(20)).toNumber()] = _601_v4;
        _nw97[(new BigNumber(21)).toNumber()] = _601_v4;
        _nw97[(new BigNumber(22)).toNumber()] = _601_v4;
        _607_v10 = _nw97;
        _607_v10 = _607_v10;
        if (p2) {
          let _608_v11;
          _608_v11 = new BigNumber(-699);
          let _rhs95 = ((p2) ? (p2) : ((_604_v7).IsDisjointFrom(_604_v7)));
          let _rhs96 = (_module.D2.create_DC6(p0)).dtor_cf8;
          let _rhs97 = new BigNumber((_603_v6).length);
          let _rhs98 = (new BigNumber(-292)).minus(_608_v11);
          let _lhs57 = globalState;
          _lhs57.f28 = _rhs95;
          _608_v11 = _rhs96;
          _608_v11 = _rhs97;
          _608_v11 = _rhs98;
          let _609_v12;
          _609_v12 = _dafny.Set.fromElements(_597_v1, _597_v1, _597_v1, _597_v1, _597_v1);
          let _610_v13;
          let _nw98 = new _module.C0();
          _nw98.__ctor((_609_v12).equals(_609_v12));
          _610_v13 = _nw98;
          _597_v1 = _597_v1;
          _608_v11 = (_608_v11).minus(p1);
          let _611_v14;
          _611_v14 = _dafny.MultiSet.fromElements(true, p3, (_606_v9).dtor_cf21);
          (globalState).f28 = (_611_v14).IsProperSubsetOf(_dafny.MultiSet.fromElements(p2, !(p2), p2));
        } else {
          let _612_v15;
          _612_v15 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
          let _613_v16;
          _613_v16 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),p2);
          (globalState).f28 = !((((_612_v15).contains(p3)) ? ((_612_v15).get(p3)) : ((((_613_v16).contains(_597_v1)) ? ((_613_v16).get(_597_v1)) : ((_this).f31)))));
          let _614_v17;
          _614_v17 = new BigNumber(100);
          _614_v17 = p1;
          (globalState).f16 = p2;
          let _615_v18;
          let _nw99 = new _module.C0();
          _nw99.__ctor(true);
          _615_v18 = _nw99;
          let _616_v19;
          let _nw100 = Array((new BigNumber(6)).toNumber()).fill([]);
          _616_v19 = _nw100;
          let _617_v20;
          _617_v20 = _dafny.Map.Empty.slice().updateUnsafe(_616_v19,_614_v17);
          let _618_v21;
          _618_v21 = _dafny.Map.Empty.slice().updateUnsafe(_600_i0,_616_v19);
          _617_v20 = (_617_v20).update((((_618_v21).contains(new BigNumber(-299))) ? ((_618_v21).get(new BigNumber(-299))) : (_616_v19)), (_this).fm13((_dafny.ZERO).minus(_614_v17), globalState));
        }
        (globalState).f28 = p3;
        let _619_v22;
        _619_v22 = _dafny.Seq.of(p0, p0, p1);
        let _620_v23;
        _620_v23 = _dafny.Map.Empty.slice().updateUnsafe((_this).f31,_619_v22);
        _620_v23 = (_620_v23).update(p2, _dafny.Seq.Concat(_619_v22, _dafny.Seq.of((_this).fm12(p1, p0, _600_i0, _605_v8, globalState))));
      }
      let _hi4 = p1;
      for (let _621_i1 = p1; _621_i1.isLessThan(_hi4); _621_i1 = _621_i1.plus(_dafny.ONE)) {
        let _622_v24;
        _622_v24 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(p1, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_621_i1,_621_i1)).length)),(_this).f31);
        _622_v24 = (_622_v24).update(p1, (_this).f31);
        let _623_v25;
        let _nw101 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Seq.of());
        _623_v25 = _nw101;
        let _624_v26;
        _624_v26 = _dafny.Map.Empty.slice().updateUnsafe(_623_v25,p1);
        _624_v26 = (_624_v26).update(_623_v25, _module.__default.safeModuloInt(p0, (_dafny.ZERO).minus(new BigNumber(-690))));
        let _625_v27;
        let _nw102 = new _module.C0();
        _nw102.__ctor((_this).f31);
        _625_v27 = _nw102;
        let _626_v28;
        let _init9 = ((_627_p1) => function (_628_i2) {
          return _module.__default.safeDivisionInt(_628_i2, _627_p1);
        })(p1);
        let _nw103 = Array((new BigNumber(8)).toNumber());
        for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw103.length); _i0_9++) {
          _nw103[_i0_9] = _init9(new BigNumber(_i0_9));
        }
        _626_v28 = _nw103;
        let _629_v29;
        _629_v29 = _626_v28;
        let _source9 = _629_v29;
        let _630___mcc_h0 = _source9;
        let _631_cf7 = _630___mcc_h0;
        let _632_v30;
        _632_v30 = _dafny.Map.Empty.slice().updateUnsafe((_this).f31,p3);
        let _633_v31;
        _633_v31 = _dafny.Seq.of(_632_v30);
        let _634_v32;
        _634_v32 = _dafny.MultiSet.fromElements(_621_i1, (_this).fm13(p1, globalState), p1);
        let _635_v33;
        _635_v33 = new BigNumber(465);
        let _636_v34;
        _636_v34 = _dafny.Seq.of((_this).f38);
        let _637_v35;
        _637_v35 = _dafny.Set.fromElements(new BigNumber((_636_v34).length), new BigNumber((_dafny.Seq.of((_this).f31)).length), _621_i1);
        let _rhs99 = _dafny.Seq.update(_633_v31, _module.__default.safeIndex(new BigNumber((_module.__default.fm15(p2, globalState)).length), new BigNumber((_633_v31).length)), (_632_v30).Merge(_632_v30));
        let _rhs100 = _634_v32;
        let _rhs101 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_596_v0, _596_v0), _596_v0)).length));
        let _rhs102 = (_this).fm12(new BigNumber((_dafny.Seq.of(true, (_this).f31, p2, (_this).f31, p2)).length), (_this).fm13(p0, globalState), p0, (_dafny.Set.fromElements(_621_i1)).Difference(_637_v35), globalState);
        let _rhs103 = p2;
        let _lhs58 = globalState;
        _633_v31 = _rhs99;
        _634_v32 = _rhs100;
        _635_v33 = _rhs101;
        _635_v33 = _rhs102;
        _lhs58.f28 = _rhs103;
        let _638_v36;
        _638_v36 = _dafny.MultiSet.fromElements((_this).f31);
        let _639_v37;
        _639_v37 = _dafny.Map.Empty.slice().updateUnsafe((_this).f31,_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jhpbetysj"), _module.__default.fm0(_638_v36, globalState)), _module.__default.safeIndex((_this).fm13(p1, globalState), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jhpbetysj"), _module.__default.fm0(_638_v36, globalState))).length)), new _dafny.CodePoint('p'.codePointAt(0))));
        _639_v37 = (_639_v37).update(true, _596_v0);
        let _rhs104 = (_this).fm13(p1, globalState);
        let _rhs105 = p1;
        let _rhs106 = (_629_v29);
        _635_v33 = _rhs104;
        _635_v33 = _rhs105;
        _631_cf7 = _rhs106;
        let _640_v38;
        let _nw104 = new _module.C0();
        _nw104.__ctor((_this).f31);
        _640_v38 = _nw104;
      }
      let _641_v39;
      _641_v39 = _dafny.Set.fromElements(p0);
      _module.__default.m0(new BigNumber(((_641_v39).Difference(_641_v39)).length), globalState);
      r0 = ((p3) ? ((_this).f31) : (p3));
      _596_v0 = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(891)), ((_642_p3, _643_v1) => function (_644_i3) {
        return ((_642_p3) ? (_643_v1) : (_643_v1));
      })(p3, _597_v1)), _module.__default.safeIndex(new BigNumber((_596_v0).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(891)), ((_645_p3, _646_v1) => function (_647_i3) {
        return ((_645_p3) ? (_646_v1) : (_646_v1));
      })(p3, _597_v1))).length)), _597_v1);
      let _648_i4;
      _648_i4 = _dafny.ZERO;
      L7: {
        while ((_this).f31) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_648_i4)) {
              break L7;
            }
            _648_i4 = (_648_i4).plus(_dafny.ONE);
            let _649_v40;
            let _nw105 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
            _649_v40 = _nw105;
            let _650_v41;
            _650_v41 = _dafny.Map.Empty.slice().updateUnsafe(_649_v40,(_this).fm13(p0, globalState));
            let _651_v42;
            _651_v42 = _dafny.Seq.of(p0, new BigNumber((((p3) ? (_596_v0) : (_596_v0))).length), _module.__default.safeDivisionInt(new BigNumber((_650_v41).length), p0), p0);
            _651_v42 = _651_v42;
            let _652_v43;
            let _init10 = function (_653_i5) {
              return true;
            };
            let _nw106 = Array((new BigNumber(13)).toNumber());
            for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw106.length); _i0_10++) {
              _nw106[_i0_10] = _init10(new BigNumber(_i0_10));
            }
            _652_v43 = _nw106;
            (globalState).f20 = _652_v43;
            r0 = !(p2) || (p2);
            if (!_dafny.Seq.contains((_this).f38, !(_dafny.MultiSet.FromArray(_651_v42)).contains(p1))) {
              let _654_v44;
              _654_v44 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(46),_597_v1);
              _654_v44 = (_654_v44).update(p0, _597_v1);
              let _655_v45;
              _655_v45 = new BigNumber(825);
              let _656_v46;
              _656_v46 = _dafny.Map.Empty.slice().updateUnsafe(p2,_597_v1);
              _655_v45 = (_this).fm12(p1, p1, (_this).fm12(new BigNumber((_656_v46).length), p0, new BigNumber((_596_v0).length), _641_v39, globalState), _641_v39, globalState);
              let _657_v47;
              _657_v47 = _dafny.MultiSet.fromElements(!(p2));
              _596_v0 = _module.__default.fm0(_657_v47, globalState);
              let _658_v48;
              _658_v48 = _dafny.Seq.of(_649_v40, _649_v40, _649_v40);
              let _659_v49;
              let _nw107 = Array((new BigNumber(7)).toNumber());
              _nw107[(_dafny.ZERO).toNumber()] = _652_v43;
              _nw107[(_dafny.ONE).toNumber()] = _652_v43;
              _nw107[(new BigNumber(2)).toNumber()] = _652_v43;
              _nw107[(new BigNumber(3)).toNumber()] = _652_v43;
              _nw107[(new BigNumber(4)).toNumber()] = _652_v43;
              _nw107[(new BigNumber(5)).toNumber()] = _652_v43;
              _nw107[(new BigNumber(6)).toNumber()] = _652_v43;
              _659_v49 = _nw107;
              let _660_v50;
              let _nw108 = new _module.C1();
              _nw108.__ctor(p1, _659_v49, (_this).f31);
              _660_v50 = _nw108;
              let _661_v51;
              _661_v51 = _dafny.Map.Empty.slice().updateUnsafe(_660_v50,_655_v45);
              _655_v45 = new BigNumber((_dafny.Seq.update(_658_v48, _module.__default.safeIndex((((_661_v51).contains(_660_v50)) ? ((_661_v51).get(_660_v50)) : (p1)), new BigNumber((_658_v48).length)), _649_v40)).length);
              let _662_v52;
              _662_v52 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_660_v50).f32),_module.__default.fm26((_this).f31, p2, p3, (_dafny.ZERO).minus((_660_v50).f32), globalState));
              _662_v52 = (_662_v52).update((_660_v50).f32, p2);
            } else {
              let _663_v53;
              _663_v53 = _dafny.Seq.of(_596_v0);
              _596_v0 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("xyaujnuy"), _596_v0), (_663_v53)[_module.__default.safeIndex(p0, new BigNumber((_663_v53).length))]);
              let _664_v54;
              _664_v54 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(new BigNumber((_599_v3).length), (_dafny.ZERO).minus(p0)),_652_v43);
              _664_v54 = (_664_v54).Merge((_module.D8.create_DC20(_dafny.Map.Empty.slice().updateUnsafe(p0,_652_v43))).dtor_cf30);
              (globalState).f28 = p2;
              let _665_v55;
              _665_v55 = new BigNumber(590);
              let _666_v56;
              _666_v56 = _dafny.Map.Empty.slice().updateUnsafe((_this).f31,_596_v0);
              _665_v55 = new BigNumber((_666_v56).length);
              let _667_v57;
              let _nw109 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.of());
              _667_v57 = _nw109;
              _667_v57 = (((_this).f31) ? (_667_v57) : (_667_v57));
            }
          }
        }
      }
      r0 = p3;
      let _668_v59;
      let _init11 = ((_669_v1, _670_p2, _671_p3) => function (_672_i6) {
        return function () {
          let _coll19 = new _dafny.Map();
          for (const _compr_19 of (_dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC7(_669_v1, _dafny.Set.fromElements(_670_p2)),false)).Keys.Elements) {
            let _673_v58 = _compr_19;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC7(_669_v1, _dafny.Set.fromElements(_670_p2)),false)).contains(_673_v58)) {
              _coll19.push([_673_v58,_671_p3]);
            }
          }
          return _coll19;
        }();
      })(_597_v1, p2, p3);
      let _nw110 = Array((new BigNumber(19)).toNumber());
      for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw110.length); _i0_11++) {
        _nw110[_i0_11] = _init11(new BigNumber(_i0_11));
      }
      _668_v59 = _nw110;
      r1 = ((p2) ? (((false) ? (_668_v59) : (_668_v59))) : (_668_v59));
      return [r0, r1];
    }
    get f38() {
      let _this = this;
      return _this._f38;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f31 = false;
      this._f36 = [];
      this._f37 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f31() {
      let _this = this;
      return _this._f31;
    };
    __ctor(f36, f37, f31) {
      let _this = this;
      (_this)._f36 = f36;
      (_this)._f37 = f37;
      (_this)._f31 = f31;
      return;
    }
    fm9(p0, p1, p2, p3, globalState) {
      let _this = this;
      return new BigNumber((function (_source10) {
        if (_source10.is_DC7) {
          let _674___mcc_h0 = (_source10).cf9;
          let _675___mcc_h1 = (_source10).cf10;
          let _676_cf10 = _675___mcc_h1;
          let _677_cf9 = _674___mcc_h0;
          if ((_this).f31) {
            return _dafny.Seq.UnicodeFromString("tj");
          } else {
            return _dafny.Seq.UnicodeFromString("puwxcvuv");
          }
        } else if (_source10.is_DC6) {
          let _678___mcc_h2 = (_source10).cf8;
          let _679_cf8 = _678___mcc_h2;
          return _dafny.Seq.UnicodeFromString("rasephr");
        } else {
          let _680___mcc_h3 = (_source10).cf11;
          let _681_cf11 = _680___mcc_h3;
          return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("wlyxgncg"), _dafny.Seq.UnicodeFromString("mit"));
        }
      }(_module.D2.create_DC6(new BigNumber((function () {
  let _coll20 = new _dafny.Map();
  for (const _compr_20 of _dafny.IntegerRange(new BigNumber(691), new BigNumber(135))) {
    let _682_v0 = _compr_20;
    if (((new BigNumber(691)).isLessThanOrEqualTo(_682_v0)) && ((_682_v0).isLessThan(new BigNumber(135)))) {
      _coll20.push([(_682_v0).minus(new BigNumber((_dafny.Set.fromElements((_this).f31, (_this).f31, (_this).f31, (_this).f31, (_this).f31)).length)),new BigNumber(640)]);
    }
  }
  return _coll20;
}()).length)))).length);
    };
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _683_v0;
      _683_v0 = new BigNumber(-523);
      _683_v0 = _module.__default.safeModuloInt(p2, p2);
      if (p3) {
        let _684_v1;
        _684_v1 = _dafny.Seq.UnicodeFromString("uda");
        let _685_v2;
        let _nw111 = Array((new BigNumber(10)).toNumber());
        _nw111[(_dafny.ZERO).toNumber()] = _684_v1;
        _nw111[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_684_v1, _684_v1);
        _nw111[(new BigNumber(2)).toNumber()] = _684_v1;
        _nw111[(new BigNumber(3)).toNumber()] = _684_v1;
        _nw111[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(60)), function (_686_i0) {
          return new _dafny.CodePoint('y'.codePointAt(0));
        }), _684_v1);
        _nw111[(new BigNumber(5)).toNumber()] = _684_v1;
        _nw111[(new BigNumber(6)).toNumber()] = _684_v1;
        _nw111[(new BigNumber(7)).toNumber()] = _684_v1;
        _nw111[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_684_v1, _dafny.Seq.UnicodeFromString("j"));
        _nw111[(new BigNumber(9)).toNumber()] = _dafny.Seq.UnicodeFromString("qviqoktq");
        _685_v2 = _nw111;
        let _index93 = _module.__default.safeIndex(new BigNumber(410), new BigNumber((_685_v2).length));
        (_685_v2)[_index93] = _684_v1;
        let _index94 = _module.__default.safeIndex(new BigNumber(410), new BigNumber((_685_v2).length));
        (_685_v2)[_index94] = _dafny.Seq.UnicodeFromString("ldgnc");
        let _687_v4;
        _687_v4 = _dafny.Map.Empty.slice().updateUnsafe(_683_v0,_683_v0);
        let _688_v5;
        _688_v5 = _dafny.MultiSet.fromElements(new BigNumber((function () {
          let _coll21 = new _dafny.Map();
          for (const _compr_21 of (_687_v4).Keys.Elements) {
            let _689_v3 = _compr_21;
            if ((_687_v4).contains(_689_v3)) {
              _coll21.push([(_689_v3).multipliedBy(new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("h"), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.UnicodeFromString("h")).length)), new _dafny.CodePoint('o'.codePointAt(0)))).length)),(_this).f31]);
            }
          }
          return _coll21;
        }()).length));
        let _690_v6;
        _690_v6 = _dafny.Map.Empty.slice().updateUnsafe((_688_v5).IsDisjointFrom(_688_v5),(_this).f31);
        _690_v6 = (_690_v6).update(p3, p3);
        let _691_v7;
        _691_v7 = _dafny.MultiSet.fromElements((_this).f31, p3);
        _684_v1 = _module.__default.fm0(_691_v7, globalState);
        let _692_v8;
        let _nw112 = Array((new BigNumber(25)).toNumber()).fill(_dafny.MultiSet.Empty);
        _692_v8 = _nw112;
        let _index95 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_692_v8).length));
        (_692_v8)[_index95] = _691_v7;
        let _693_v9;
        let _nw113 = Array((new BigNumber(7)).toNumber()).fill([]);
        _693_v9 = _nw113;
        let _694_v10;
        _694_v10 = _dafny.Set.fromElements((_this).f31);
        let _695_v11;
        _695_v11 = new _dafny.CodePoint('d'.codePointAt(0));
        let _696_v12;
        _696_v12 = _module.D0.create_DC1(_695_v11, (_this).f31, p3);
        let _697_v13;
        let _nw114 = Array((new BigNumber(25)).toNumber());
        _nw114[(_dafny.ZERO).toNumber()] = _module.__default.fm10(_695_v11, (_this).f31, globalState);
        _nw114[(_dafny.ONE).toNumber()] = false;
        _nw114[(new BigNumber(2)).toNumber()] = (_this).f31;
        _nw114[(new BigNumber(3)).toNumber()] = p3;
        _nw114[(new BigNumber(4)).toNumber()] = (_this).f31;
        _nw114[(new BigNumber(5)).toNumber()] = true;
        _nw114[(new BigNumber(6)).toNumber()] = p3;
        _nw114[(new BigNumber(7)).toNumber()] = (_this).f31;
        _nw114[(new BigNumber(8)).toNumber()] = p3;
        _nw114[(new BigNumber(9)).toNumber()] = (_this).f31;
        _nw114[(new BigNumber(10)).toNumber()] = true;
        _nw114[(new BigNumber(11)).toNumber()] = p3;
        _nw114[(new BigNumber(12)).toNumber()] = (_this).f31;
        _nw114[(new BigNumber(13)).toNumber()] = (_this).f31;
        _nw114[(new BigNumber(14)).toNumber()] = (_696_v12).dtor_cf3;
        _nw114[(new BigNumber(15)).toNumber()] = p3;
        _nw114[(new BigNumber(16)).toNumber()] = !((_this).f31);
        _nw114[(new BigNumber(17)).toNumber()] = (_this).f31;
        _nw114[(new BigNumber(18)).toNumber()] = _module.__default.fm10(_695_v11, (_this).f31, globalState);
        _nw114[(new BigNumber(19)).toNumber()] = true;
        _nw114[(new BigNumber(20)).toNumber()] = p3;
        _nw114[(new BigNumber(21)).toNumber()] = true;
        _nw114[(new BigNumber(22)).toNumber()] = p3;
        _nw114[(new BigNumber(23)).toNumber()] = p3;
        _nw114[(new BigNumber(24)).toNumber()] = !((_this).f31);
        _697_v13 = _nw114;
        let _698_v14;
        _698_v14 = _dafny.Map.Empty.slice().updateUnsafe(p3,_697_v13);
        let _699_v15;
        let _nw115 = Array((new BigNumber(5)).toNumber());
        _nw115[(_dafny.ZERO).toNumber()] = new BigNumber((_694_v10).length);
        _nw115[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Set.fromElements(_683_v0)).length);
        _nw115[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus((((_691_v7).contains((_this).f31)) ? ((_691_v7).get((_this).f31)) : (new BigNumber((_698_v14).length))));
        _nw115[(new BigNumber(3)).toNumber()] = _683_v0;
        _nw115[(new BigNumber(4)).toNumber()] = new BigNumber((p1).length);
        _699_v15 = _nw115;
        let _index96 = _module.__default.safeIndex(new BigNumber(270), new BigNumber((_693_v9).length));
        (_693_v9)[_index96] = _699_v15;
        let _700_v16;
        _700_v16 = _module.D5.create_DC13(_691_v7);
        let _index97 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_692_v8).length));
        let _index98 = _module.__default.safeIndex(new BigNumber(270), new BigNumber((_693_v9).length));
        let _rhs107 = (_700_v16).dtor_cf18;
        let _rhs108 = _699_v15;
        let _lhs59 = _692_v8;
        let _lhs60 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_692_v8).length));
        let _lhs61 = _693_v9;
        let _lhs62 = _module.__default.safeIndex(new BigNumber(270), new BigNumber((_693_v9).length));
        _lhs59[_lhs60] = _rhs107;
        _lhs61[_lhs62] = _rhs108;
        let _arr0 = (_693_v9)[_module.__default.safeIndex(new BigNumber(270), new BigNumber((_693_v9).length))];
        let _index99 = _module.__default.safeIndex(new BigNumber(843), new BigNumber(((_693_v9)[_module.__default.safeIndex(new BigNumber(270), new BigNumber((_693_v9).length))]).length));
        _arr0[_index99] = (_dafny.ZERO).minus(p2);
        let _701_v17;
        _701_v17 = _dafny.Seq.of(p2);
        let _702_v18;
        _702_v18 = _dafny.Map.Empty.slice().updateUnsafe(_699_v15,_dafny.Seq.Concat(_module.__default.fm11(globalState), _701_v17));
        let _703_v19;
        _703_v19 = _dafny.Map.Empty.slice().updateUnsafe((_this).f31,p2);
        let _arr1 = (_693_v9)[_module.__default.safeIndex(new BigNumber(270), new BigNumber((_693_v9).length))];
        let _index100 = _module.__default.safeIndex(new BigNumber(843), new BigNumber(((_693_v9)[_module.__default.safeIndex(new BigNumber(270), new BigNumber((_693_v9).length))]).length));
        let _index101 = _module.__default.safeIndex(new BigNumber(410), new BigNumber((_685_v2).length));
        let _rhs109 = new BigNumber(((((_702_v18).contains(((p3) ? (_699_v15) : (_699_v15)))) ? ((_702_v18).get(((p3) ? (_699_v15) : (_699_v15)))) : (((p3) ? (_701_v17) : (_701_v17))))).length);
        let _rhs110 = p2;
        let _rhs111 = (_683_v0).multipliedBy(_683_v0);
        let _rhs112 = _dafny.Seq.update((_685_v2)[_module.__default.safeIndex(new BigNumber(410), new BigNumber((_685_v2).length))], _module.__default.safeIndex(_module.__default.safeDivisionInt(p2, _683_v0), new BigNumber(((_685_v2)[_module.__default.safeIndex(new BigNumber(410), new BigNumber((_685_v2).length))]).length)), ((_685_v2)[_module.__default.safeIndex(new BigNumber(410), new BigNumber((_685_v2).length))])[_module.__default.safeIndex(new BigNumber((_703_v19).length), new BigNumber(((_685_v2)[_module.__default.safeIndex(new BigNumber(410), new BigNumber((_685_v2).length))]).length))]);
        let _lhs63 = (_693_v9)[_module.__default.safeIndex(new BigNumber(270), new BigNumber((_693_v9).length))];
        let _lhs64 = _module.__default.safeIndex(new BigNumber(843), new BigNumber(((_693_v9)[_module.__default.safeIndex(new BigNumber(270), new BigNumber((_693_v9).length))]).length));
        let _lhs65 = _685_v2;
        let _lhs66 = _module.__default.safeIndex(new BigNumber(410), new BigNumber((_685_v2).length));
        _683_v0 = _rhs109;
        _683_v0 = _rhs110;
        _lhs63[_lhs64] = _rhs111;
        _lhs65[_lhs66] = _rhs112;
      } else {
        let _704_v20;
        _704_v20 = _dafny.Set.fromElements(_683_v0, p2);
        let _705_v21;
        _705_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this).f31,new BigNumber((_704_v20).length));
        let _706_v22;
        _706_v22 = _dafny.MultiSet.fromElements(p2, p2, (((_705_v21).contains(p3)) ? ((_705_v21).get(p3)) : (_683_v0)));
        let _707_v23;
        _707_v23 = _dafny.Seq.UnicodeFromString("eys");
        let _708_v24;
        _708_v24 = _dafny.Seq.of(_683_v0);
        let _709_v25;
        let _nw116 = Array((new BigNumber(10)).toNumber());
        _nw116[(_dafny.ZERO).toNumber()] = (_this).f31;
        _nw116[(_dafny.ONE).toNumber()] = !(_706_v22).equals(_706_v22);
        _nw116[(new BigNumber(2)).toNumber()] = (p2).isLessThan(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(918)), ((_710_p2) => function (_711_i1) {
          return (_dafny.ZERO).minus(_710_p2);
        })(p2)))).cardinality()));
        _nw116[(new BigNumber(3)).toNumber()] = ((_706_v22).update(_683_v0, _module.__default.abs(_683_v0))).IsSubsetOf(_706_v22);
        _nw116[(new BigNumber(4)).toNumber()] = false;
        _nw116[(new BigNumber(5)).toNumber()] = true;
        _nw116[(new BigNumber(6)).toNumber()] = (_683_v0).isLessThan(_683_v0);
        _nw116[(new BigNumber(7)).toNumber()] = (_this).f31;
        _nw116[(new BigNumber(8)).toNumber()] = _dafny.areEqual(_707_v23, _707_v23);
        _nw116[(new BigNumber(9)).toNumber()] = _dafny.areEqual(_708_v24, _708_v24);
        _709_v25 = _nw116;
        let _index102 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_709_v25).length));
        (_709_v25)[_index102] = !(p3);
        let _712_v26;
        let _init12 = function (_713_i2) {
          return new _dafny.CodePoint('w'.codePointAt(0));
        };
        let _nw117 = Array((new BigNumber(16)).toNumber());
        for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw117.length); _i0_12++) {
          _nw117[_i0_12] = _init12(new BigNumber(_i0_12));
        }
        _712_v26 = _nw117;
        let _714_v27;
        _714_v27 = _dafny.Map.Empty.slice().updateUnsafe(true,_712_v26);
        let _715_v28;
        _715_v28 = new _dafny.CodePoint('h'.codePointAt(0));
        let _716_v29;
        _716_v29 = _dafny.Map.Empty.slice().updateUnsafe((_this).f31,_715_v28);
        let _717_v30;
        let _nw118 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
        _717_v30 = _nw118;
        let _718_v31;
        _718_v31 = _dafny.Map.Empty.slice().updateUnsafe(_717_v30,p2);
        let _719_v32;
        _719_v32 = _dafny.Map.Empty.slice().updateUnsafe(p2,_717_v30);
        let _720_v33;
        let _nw119 = Array((new BigNumber(7)).toNumber());
        _nw119[(_dafny.ZERO).toNumber()] = new BigNumber((_714_v27).length);
        _nw119[(_dafny.ONE).toNumber()] = (_this).fm9(new BigNumber(977), _708_v24, _683_v0, (_this).f31, globalState);
        _nw119[(new BigNumber(2)).toNumber()] = (_this).fm9((_this).fm9(new BigNumber(557), _708_v24, _683_v0, p3, globalState), _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-475)), function (_721_i3) {
          return new BigNumber(-745);
        }), _module.__default.safeIndex(new BigNumber((_716_v29).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-475)), function (_722_i3) {
          return new BigNumber(-745);
        })).length)), new BigNumber((_718_v31).length)), p2, (_this).f31, globalState);
        _nw119[(new BigNumber(3)).toNumber()] = _683_v0;
        _nw119[(new BigNumber(4)).toNumber()] = new BigNumber(((_719_v32).Merge(_dafny.Map.Empty.slice().updateUnsafe(p2,_717_v30))).length);
        _nw119[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((p1).length));
        _nw119[(new BigNumber(6)).toNumber()] = _683_v0;
        _720_v33 = _nw119;
        let _index103 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_717_v30).length));
        (_717_v30)[_index103] = _683_v0;
        let _index104 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_709_v25).length));
        let _index105 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_717_v30).length));
        let _rhs113 = _709_v25;
        let _rhs114 = p2;
        let _rhs115 = (((_this).f31) ? ((_this).f31) : ((_this).f31));
        let _rhs116 = (_683_v0).minus((_this).fm9(_683_v0, _708_v24, new BigNumber(927), (_this).f31, globalState));
        let _lhs67 = globalState;
        let _lhs68 = _709_v25;
        let _lhs69 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_709_v25).length));
        let _lhs70 = _717_v30;
        let _lhs71 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_717_v30).length));
        _lhs67.f27 = _rhs113;
        _683_v0 = _rhs114;
        _lhs68[_lhs69] = _rhs115;
        _lhs70[_lhs71] = _rhs116;
        let _723_v34;
        let _nw120 = new _module.C0();
        _nw120.__ctor((_709_v25)[_module.__default.safeIndex(new BigNumber(758), new BigNumber((_709_v25).length))]);
        _723_v34 = _nw120;
        let _724_v35;
        let _nw121 = Array((_dafny.ONE).toNumber());
        _nw121[(_dafny.ZERO).toNumber()] = _720_v33;
        _724_v35 = _nw121;
        let _index106 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_724_v35).length));
        (_724_v35)[_index106] = _720_v33;
        let _725_v36;
        _725_v36 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_717_v30)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_717_v30).length))]);
        let _726_v37;
        _726_v37 = _dafny.Seq.of((_706_v22).Intersect(_dafny.MultiSet.fromElements(new BigNumber(((_725_v36).update(p2, new BigNumber(975))).length))), _dafny.MultiSet.FromArray(_708_v24), _706_v22);
        let _727_v38;
        _727_v38 = _dafny.MultiSet.fromElements((_this).f31);
        let _728_v39;
        _728_v39 = _module.D8.create_DC21(p3, (_this).f31, (_717_v30)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_717_v30).length))]);
        let _pat_let_tv10 = _709_v25;
        let _pat_let_tv11 = _709_v25;
        let _index107 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_724_v35).length));
        let _index108 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_717_v30).length));
        let _rhs117 = (_dafny.ZERO).minus(new BigNumber((_726_v37).length));
        let _rhs118 = ((p3) ? (_720_v33) : (_720_v33));
        let _rhs119 = (((_709_v25)[_module.__default.safeIndex(new BigNumber(758), new BigNumber((_709_v25).length))]) ? (_715_v28) : (_715_v28));
        let _rhs120 = p2;
        let _rhs121 = _module.__default.fm0((_727_v38).update((_709_v25)[_module.__default.safeIndex(new BigNumber(758), new BigNumber((_709_v25).length))], _module.__default.abs((function (_pat_let16_0) {
          return function (_729_dt__update__tmp_h0) {
            return function (_pat_let17_0) {
              return function (_730_dt__update_hcf31_h0) {
                return _module.D8.create_DC21(_730_dt__update_hcf31_h0, (_729_dt__update__tmp_h0).dtor_cf32, (_729_dt__update__tmp_h0).dtor_cf33);
              }(_pat_let17_0);
            }((_pat_let_tv11)[_module.__default.safeIndex(new BigNumber(758), new BigNumber((_pat_let_tv10).length))]);
          }(_pat_let16_0);
        }(_728_v39)).dtor_cf33)), globalState);
        let _lhs72 = _724_v35;
        let _lhs73 = _module.__default.safeIndex(new BigNumber(8), new BigNumber((_724_v35).length));
        let _lhs74 = _717_v30;
        let _lhs75 = _module.__default.safeIndex(new BigNumber(905), new BigNumber((_717_v30).length));
        _683_v0 = _rhs117;
        _lhs72[_lhs73] = _rhs118;
        _715_v28 = _rhs119;
        _lhs74[_lhs75] = _rhs120;
        _707_v23 = _rhs121;
        let _731_v40;
        let _nw122 = Array((new BigNumber(9)).toNumber()).fill([]);
        _731_v40 = _nw122;
        let _732_v41;
        _732_v41 = _module.D5.create_DC14(_704_v20, _683_v0, p3, _709_v25);
        let _index109 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_731_v40).length));
        (_731_v40)[_index109] = ((p3) ? ((_732_v41).dtor_cf22) : (_709_v25));
        let _index110 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_731_v40).length));
        (_731_v40)[_index110] = _709_v25;
        let _733_v42;
        _733_v42 = _dafny.Seq.of(_707_v23);
        let _734_v43;
        _734_v43 = _dafny.Seq.of(_dafny.Seq.update(_708_v24, _module.__default.safeIndex((_717_v30)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_717_v30).length))], new BigNumber((_708_v24).length)), (_717_v30)[_module.__default.safeIndex(new BigNumber(905), new BigNumber((_717_v30).length))]), _708_v24);
        let _735_v44;
        let _nw123 = new _module.C1();
        _nw123.__ctor(new BigNumber((_734_v43).length), _731_v40, (_this).f31);
        _735_v44 = _nw123;
        let _736_v45;
        _736_v45 = _dafny.Map.Empty.slice().updateUnsafe(_683_v0,_735_v44);
        _733_v42 = (_module.__default.fm27(p3, new BigNumber((_736_v45).length), (_this).f31, globalState)).dtor_cf39;
      }
      let _737_v46;
      _737_v46 = _dafny.MultiSet.fromElements(p3, (_this).f31, false);
      let _738_v47;
      _738_v47 = _dafny.Seq.of((_737_v46).IsSubsetOf(_737_v46));
      let _739_v48;
      _739_v48 = _dafny.Map.Empty.slice().updateUnsafe(p2,p3);
      let _740_v49;
      _740_v49 = _dafny.Seq.UnicodeFromString("ugtivxrgy");
      let _741_v50;
      _741_v50 = _module.D9.create_DC26(_module.__default.fm26(p3, p3, (((_739_v48).contains(_683_v0)) ? ((_739_v48).get(_683_v0)) : (p3)), p2, globalState), _740_v49, _683_v0, _739_v48, p3);
      let _pat_let_tv12 = p1;
      let _pat_let_tv13 = p1;
      let _pat_let_tv14 = _738_v47;
      let _pat_let_tv15 = _738_v47;
      let _pat_let_tv16 = p1;
      _738_v47 = function (_source11) {
        if (_source11.is_DC24) {
          let _742___mcc_h0 = (_source11).cf40;
          let _743_cf40 = _742___mcc_h0;
          return _pat_let_tv12;
        } else if (_source11.is_DC25) {
          let _744___mcc_h1 = (_source11).cf41;
          let _745___mcc_h2 = (_source11).cf42;
          let _746___mcc_h3 = (_source11).cf43;
          let _747___mcc_h4 = (_source11).cf44;
          let _748_cf44 = _747___mcc_h4;
          let _749_cf43 = _746___mcc_h3;
          let _750_cf42 = _745___mcc_h2;
          let _751_cf41 = _744___mcc_h1;
          return _pat_let_tv13;
        } else if (_source11.is_DC26) {
          let _752___mcc_h5 = (_source11).cf45;
          let _753___mcc_h6 = (_source11).cf46;
          let _754___mcc_h7 = (_source11).cf47;
          let _755___mcc_h8 = (_source11).cf48;
          let _756___mcc_h9 = (_source11).cf49;
          let _757_cf49 = _756___mcc_h9;
          let _758_cf48 = _755___mcc_h8;
          let _759_cf47 = _754___mcc_h7;
          let _760_cf46 = _753___mcc_h6;
          let _761_cf45 = _752___mcc_h5;
          return _pat_let_tv14;
        } else {
          let _762___mcc_h10 = (_source11).cf39;
          let _763_cf39 = _762___mcc_h10;
          return _dafny.Seq.Concat(_pat_let_tv15, _pat_let_tv16);
        }
      }(_741_v50);
      _683_v0 = p2;
      let _764_v51;
      let _nw124 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
      _764_v51 = _nw124;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_764_v51).length))) {
        let _765_i4 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_765_i4)) && ((_765_i4).isLessThan(new BigNumber((_764_v51).length))))) {
          (_764_v51)[(_765_i4)] = (_765_i4).multipliedBy((new BigNumber((p1).length)).multipliedBy(new BigNumber(47)));
        }
      }
      if (((_683_v0).multipliedBy((_dafny.ZERO).minus(p2))).isLessThanOrEqualTo(p2)) {
        _module.__default.m0(p2, globalState);
        let _766_v52;
        _766_v52 = _dafny.MultiSet.fromElements(p1, _738_v47, _738_v47);
        (globalState).f16 = (_766_v52).IsDisjointFrom(_766_v52);
        let _767_v53;
        _767_v53 = _dafny.Map.Empty.slice().updateUnsafe(p3,_683_v0);
        _767_v53 = _767_v53;
        let _768_v54;
        _768_v54 = _module.D2.create_DC6(p2);
        let _769_v55;
        _769_v55 = _dafny.Set.fromElements(_768_v54);
        let _770_v56;
        _770_v56 = _769_v55;
        let _source12 = _770_v56;
        let _771___mcc_h11 = _source12;
        let _772_cf17 = _771___mcc_h11;
        let _773_v57;
        _773_v57 = _dafny.Seq.of(_741_v50);
        let _774_v59;
        _774_v59 = _dafny.MultiSet.fromElements(function () {
          let _coll22 = new _dafny.Set();
          for (const _compr_22 of (_773_v57).Elements) {
            let _775_v58 = _compr_22;
            if (_dafny.Seq.contains(_773_v57, _775_v58)) {
              _coll22.add(_775_v58);
            }
          }
          return _coll22;
        }());
        let _776_v60;
        _776_v60 = _dafny.Map.Empty.slice().updateUnsafe(_774_v59,(_dafny.ZERO).minus((p2).multipliedBy((_dafny.ZERO).minus(p2))));
        let _777_v61;
        _777_v61 = _dafny.Seq.of(_774_v59);
        _683_v0 = (((_776_v60).contains(((_777_v61)[_module.__default.safeIndex(_683_v0, new BigNumber((_777_v61).length))]).Union(_774_v59))) ? ((_776_v60).get(((_777_v61)[_module.__default.safeIndex(_683_v0, new BigNumber((_777_v61).length))]).Union(_774_v59))) : (_683_v0));
        let _778_v62;
        _778_v62 = _dafny.Set.fromElements(true);
        let _779_v63;
        let _nw125 = Array((_dafny.ONE).toNumber());
        _nw125[(_dafny.ZERO).toNumber()] = (_dafny.Set.fromElements(p3)).Intersect(_778_v62);
        _779_v63 = _nw125;
        _779_v63 = _779_v63;
        _683_v0 = p2;
        let _780_v64;
        _780_v64 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_740_v49).length),_683_v0);
        let _781_v65;
        _781_v65 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_780_v64).length),_740_v49);
        let _782_v66;
        _782_v66 = _dafny.Map.Empty.slice().updateUnsafe((_this).f31,_740_v49);
        let _783_v67;
        _783_v67 = new _dafny.CodePoint('k'.codePointAt(0));
        let _784_v68;
        let _nw126 = Array((new BigNumber(27)).toNumber());
        _nw126[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_740_v49, (((_781_v65).contains(p2)) ? ((_781_v65).get(p2)) : (_740_v49)));
        _nw126[(_dafny.ONE).toNumber()] = _dafny.Seq.UnicodeFromString("lyh");
        _nw126[(new BigNumber(2)).toNumber()] = _740_v49;
        _nw126[(new BigNumber(3)).toNumber()] = _740_v49;
        _nw126[(new BigNumber(4)).toNumber()] = (((_782_v66).contains((_738_v47)[_module.__default.safeIndex(p2, new BigNumber((_738_v47).length))])) ? ((_782_v66).get((_738_v47)[_module.__default.safeIndex(p2, new BigNumber((_738_v47).length))])) : (_740_v49));
        _nw126[(new BigNumber(5)).toNumber()] = _740_v49;
        _nw126[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_740_v49, _740_v49);
        _nw126[(new BigNumber(7)).toNumber()] = _740_v49;
        _nw126[(new BigNumber(8)).toNumber()] = ((p3) ? (_740_v49) : (_dafny.Seq.UnicodeFromString("icxd")));
        _nw126[(new BigNumber(9)).toNumber()] = _740_v49;
        _nw126[(new BigNumber(10)).toNumber()] = _740_v49;
        _nw126[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm0(_737_v46, globalState), _740_v49);
        _nw126[(new BigNumber(12)).toNumber()] = _740_v49;
        _nw126[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_740_v49, _740_v49);
        _nw126[(new BigNumber(14)).toNumber()] = _740_v49;
        _nw126[(new BigNumber(15)).toNumber()] = _740_v49;
        _nw126[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_740_v49, _740_v49);
        _nw126[(new BigNumber(17)).toNumber()] = _740_v49;
        _nw126[(new BigNumber(18)).toNumber()] = _dafny.Seq.UnicodeFromString("eghuqqi");
        _nw126[(new BigNumber(19)).toNumber()] = _740_v49;
        _nw126[(new BigNumber(20)).toNumber()] = _dafny.Seq.update(_740_v49, _module.__default.safeIndex(_683_v0, new BigNumber((_740_v49).length)), _783_v67);
        _nw126[(new BigNumber(21)).toNumber()] = _dafny.Seq.Concat(_740_v49, _740_v49);
        _nw126[(new BigNumber(22)).toNumber()] = _740_v49;
        _nw126[(new BigNumber(23)).toNumber()] = _740_v49;
        _nw126[(new BigNumber(24)).toNumber()] = _dafny.Seq.Concat(_740_v49, _740_v49);
        _nw126[(new BigNumber(25)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("s"), _module.__default.safeIndex(new BigNumber(573), new BigNumber((_dafny.Seq.UnicodeFromString("s")).length)), _783_v67), _740_v49);
        _nw126[(new BigNumber(26)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(493)), ((_785_v67) => function (_786_i5) {
          return _785_v67;
        })(_783_v67));
        _784_v68 = _nw126;
        _784_v68 = _784_v68;
        let _787_v69;
        let _init13 = function (_788_i6) {
          return (_this).f31;
        };
        let _nw127 = Array((new BigNumber(20)).toNumber());
        for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw127.length); _i0_13++) {
          _nw127[_i0_13] = _init13(new BigNumber(_i0_13));
        }
        _787_v69 = _nw127;
        let _789_v70;
        _789_v70 = _dafny.Set.fromElements(new BigNumber((_740_v49).length), new BigNumber((_dafny.Seq.of(p3)).length));
        let _index111 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_787_v69).length));
        (_787_v69)[_index111] = (_789_v70).IsSubsetOf(_789_v70);
        let _index112 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_787_v69).length));
        (_787_v69)[_index112] = (_this).f31;
      } else {
        let _790_v71;
        let _init14 = ((_791_v0, _792_p2) => function (_793_i7) {
          return (_dafny.MultiSet.FromArray(_dafny.Seq.of(_791_v0, _792_p2))).Union(_dafny.MultiSet.fromElements(_791_v0, new BigNumber(213), _792_p2));
        })(_683_v0, p2);
        let _nw128 = Array((new BigNumber(2)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw128.length); _i0_14++) {
          _nw128[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _790_v71 = _nw128;
        let _794_v72;
        _794_v72 = _dafny.MultiSet.fromElements(p2, p2);
        let _index113 = _module.__default.safeIndex(new BigNumber(849), new BigNumber((_790_v71).length));
        (_790_v71)[_index113] = (_794_v72).Union(_dafny.MultiSet.fromElements(new BigNumber((_738_v47).length), _683_v0));
        let _index114 = _module.__default.safeIndex(new BigNumber(849), new BigNumber((_790_v71).length));
        (_790_v71)[_index114] = (_794_v72).update(p2, _module.__default.abs(new BigNumber(955)));
        _683_v0 = (p2).minus(new BigNumber(-367));
        _740_v49 = _dafny.Seq.Concat(_dafny.Seq.Concat(_740_v49, _740_v49), _740_v49);
        let _795_v73;
        _795_v73 = _dafny.Map.Empty.slice().updateUnsafe((_740_v49)[_module.__default.safeIndex(_683_v0, new BigNumber((_740_v49).length))],_683_v0);
        let _796_v74;
        _796_v74 = new _dafny.CodePoint('n'.codePointAt(0));
        _795_v73 = (_795_v73).update(_796_v74, (_this).fm9((_741_v50).dtor_cf47, _dafny.Seq.of(_683_v0, _683_v0), new BigNumber((_738_v47).length), (_this).f31, globalState));
        _module.__default.m0(p2, globalState);
      }
      return;
    }
    get f36() {
      let _this = this;
      return _this._f36;
    };
    get f37() {
      let _this = this;
      return _this._f37;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f31 = false;
      this._f32 = _dafny.ZERO;
      this._f33 = [];
      this.f35 = false;
      this._f34 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f31() {
      let _this = this;
      return _this._f31;
    };
    get f32() {
      let _this = this;
      return _this._f32;
    };
    get f33() {
      let _this = this;
      return _this._f33;
    };
    __ctor(f34, f35, f32, f33, f31) {
      let _this = this;
      (_this)._f34 = f34;
      (_this).f35 = f35;
      (_this)._f32 = f32;
      (_this)._f33 = f33;
      (_this)._f31 = f31;
      return;
    }
    fm6(p0, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_this.f35, (_this).f31), _dafny.Seq.of(true)), _dafny.Seq.of(!(_this.f35)));
    };
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = [];
      let r2 = _dafny.Seq.UnicodeFromString("");
      let r3 = _dafny.ZERO;
      let _797_i0;
      _797_i0 = _dafny.ZERO;
      L8: {
        while (!((_module.__default.safeModuloInt((_this).f32, (_this).f32)).isLessThan((_this).f32))) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_797_i0)) {
              break L8;
            }
            _797_i0 = (_797_i0).plus(_dafny.ONE);
            (globalState).f28 = (_this).f31;
            r0 = new BigNumber(462);
            let _798_v1;
            _798_v1 = _dafny.Seq.of((_this).f32, (_dafny.ZERO).minus((new BigNumber((function () {
              let _coll23 = new _dafny.Set();
              for (const _compr_23 of _dafny.IntegerRange(new BigNumber(180), new BigNumber(462))) {
                let _799_v0 = _compr_23;
                if (((new BigNumber(180)).isLessThanOrEqualTo(_799_v0)) && ((_799_v0).isLessThan(new BigNumber(462)))) {
                  _coll23.add((_799_v0).minus(new BigNumber((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_this).f31,p0), _dafny.Map.Empty.slice().updateUnsafe(true,_this.f35))).length)));
                }
              }
              return _coll23;
            }()).length)).minus((_this).f32)), (_this).f32);
            _798_v1 = _798_v1;
            r3 = _module.__default.fm7((_this).f31, (_this).f31, globalState);
          }
        }
      }
      let _800_v2;
      _800_v2 = _dafny.Seq.UnicodeFromString("w");
      let _801_v3;
      _801_v3 = _dafny.MultiSet.fromElements(_800_v2);
      let _802_v4;
      _802_v4 = _dafny.Seq.of(p0);
      let _803_v5;
      _803_v5 = _dafny.Seq.of((_this).f32, (_this).f32);
      let _804_v6;
      _804_v6 = _dafny.Seq.of((((_801_v3).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(876)), function (_806_i1) {
        return new _dafny.CodePoint('w'.codePointAt(0));
      }))) ? ((_801_v3).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(876)), function (_805_i1) {
        return new _dafny.CodePoint('w'.codePointAt(0));
      }))) : (new BigNumber((_802_v4).length))), new BigNumber((_803_v5).length), (_this).f32, (_this).f32, new BigNumber(214));
      let _807_v7;
      _807_v7 = _dafny.Set.fromElements(_803_v5, _804_v6);
      if (((_807_v7).Difference(_807_v7)).IsSubsetOf((_807_v7).Difference(_807_v7))) {
        (globalState).f16 = _this.f35;
        let _808_v8;
        let _init15 = ((_809_v4) => function (_810_i2) {
          return _809_v4;
        })(_802_v4);
        let _nw129 = Array((new BigNumber(9)).toNumber());
        for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw129.length); _i0_15++) {
          _nw129[_i0_15] = _init15(new BigNumber(_i0_15));
        }
        _808_v8 = _nw129;
        let _811_v9;
        _811_v9 = _module.D0.create_DC0(_808_v8);
        let _812_v10;
        _812_v10 = _dafny.Set.fromElements(_this.f35);
        let _pat_let_tv17 = _808_v8;
        let _pat_let_tv18 = _808_v8;
        let _rhs122 = new BigNumber((_812_v10).length);
        let _rhs123 = function (_pat_let18_0) {
          return function (_815_dt__update__tmp_h1) {
            return function (_pat_let21_0) {
              return function (_816_dt__update_hcf0_h1) {
                return _module.D0.create_DC0(_816_dt__update_hcf0_h1);
              }(_pat_let21_0);
            }(_pat_let_tv18);
          }(_pat_let18_0);
        }(function (_pat_let19_0) {
          return function (_813_dt__update__tmp_h0) {
            return function (_pat_let20_0) {
              return function (_814_dt__update_hcf0_h0) {
                return _module.D0.create_DC0(_814_dt__update_hcf0_h0);
              }(_pat_let20_0);
            }(_pat_let_tv17);
          }(_pat_let19_0);
        }(_811_v9));
        r3 = _rhs122;
        _811_v9 = _rhs123;
        _802_v4 = _802_v4;
        let _817_v11;
        let _nw130 = Array((new BigNumber(21)).toNumber()).fill(false);
        _817_v11 = _nw130;
        let _index115 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_817_v11).length));
        (_817_v11)[_index115] = _module.__default.fm8(p0, new BigNumber(989), globalState);
        let _index116 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_817_v11).length));
        (_817_v11)[_index116] = (_this).f31;
        let _818_v12;
        _818_v12 = _dafny.MultiSet.fromElements(new BigNumber(-823), (_this).f32, (_this).f32);
        let _819_v13;
        let _nw131 = Array((new BigNumber(5)).toNumber());
        _nw131[(_dafny.ZERO).toNumber()] = new BigNumber(((_818_v12).Intersect(_818_v12)).cardinality());
        _nw131[(_dafny.ONE).toNumber()] = (_this).f32;
        _nw131[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus((_this).f32);
        _nw131[(new BigNumber(3)).toNumber()] = new BigNumber(592);
        _nw131[(new BigNumber(4)).toNumber()] = new BigNumber(-12);
        _819_v13 = _nw131;
        let _index117 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_819_v13).length));
        (_819_v13)[_index117] = ((_this).f32).minus(new BigNumber(439));
        let _index118 = _module.__default.safeIndex(new BigNumber(592), new BigNumber((_819_v13).length));
        (_819_v13)[_index118] = (_module.__default.fm7(p0, (_817_v11)[_module.__default.safeIndex(new BigNumber(888), new BigNumber((_817_v11).length))], globalState)).plus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_this.f35, p0), _802_v4)).length));
      } else {
        (globalState).f28 = (_this).f31;
        let _820_v14;
        _820_v14 = _module.D3.create_DC10(_dafny.Seq.Create(_module.__default.abs(new BigNumber(455)), function (_821_i3) {
  return new _dafny.CodePoint('b'.codePointAt(0));
}), (_this).f32, p0);
        r3 = ((_this.f35) ? ((_820_v14).dtor_cf14) : ((_this).f32));
        let _822_v15;
        _822_v15 = new _dafny.CodePoint('y'.codePointAt(0));
        (globalState).f28 = (_dafny.MultiSet.fromElements(_822_v15)).IsSubsetOf(_dafny.MultiSet.fromElements(_822_v15, _822_v15));
        _802_v4 = _802_v4;
        r3 = new BigNumber((_800_v2).length);
      }
      let _823_v16;
      let _nw132 = Array((new BigNumber(22)).toNumber());
      _nw132[(_dafny.ZERO).toNumber()] = !(p0);
      _nw132[(_dafny.ONE).toNumber()] = (_this).f31;
      _nw132[(new BigNumber(2)).toNumber()] = p0;
      _nw132[(new BigNumber(3)).toNumber()] = p0;
      _nw132[(new BigNumber(4)).toNumber()] = _this.f35;
      _nw132[(new BigNumber(5)).toNumber()] = !(p0);
      _nw132[(new BigNumber(6)).toNumber()] = p0;
      _nw132[(new BigNumber(7)).toNumber()] = (_this).f31;
      _nw132[(new BigNumber(8)).toNumber()] = !(false);
      _nw132[(new BigNumber(9)).toNumber()] = (_this).f31;
      _nw132[(new BigNumber(10)).toNumber()] = (_this).f31;
      _nw132[(new BigNumber(11)).toNumber()] = _this.f35;
      _nw132[(new BigNumber(12)).toNumber()] = _this.f35;
      _nw132[(new BigNumber(13)).toNumber()] = true;
      _nw132[(new BigNumber(14)).toNumber()] = !((_this).f31);
      _nw132[(new BigNumber(15)).toNumber()] = _module.__default.fm8(p0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(79)), function (_824_i4) {
        return new _dafny.CodePoint('p'.codePointAt(0));
      })).length), globalState);
      _nw132[(new BigNumber(16)).toNumber()] = p0;
      _nw132[(new BigNumber(17)).toNumber()] = p0;
      _nw132[(new BigNumber(18)).toNumber()] = p0;
      _nw132[(new BigNumber(19)).toNumber()] = _this.f35;
      _nw132[(new BigNumber(20)).toNumber()] = p0;
      _nw132[(new BigNumber(21)).toNumber()] = (_this).f31;
      _823_v16 = _nw132;
      let _825_v17;
      _825_v17 = _dafny.Seq.of(_823_v16, _823_v16);
      (globalState).f20 = (_825_v17)[_module.__default.safeIndex((_dafny.ZERO).minus(_module.__default.safeModuloInt((_this).f32, (_this).f32)), new BigNumber((_825_v17).length))];
      let _826_v18;
      _826_v18 = _dafny.Map.Empty.slice().updateUnsafe((_this).f31,(_this).f31);
      let _827_v19;
      let _nw133 = Array((_dafny.ONE).toNumber());
      _nw133[(_dafny.ZERO).toNumber()] = _826_v18;
      _827_v19 = _nw133;
      let _828_v20;
      _828_v20 = _module.D10.create_DC27(_dafny.Map.Empty.slice().updateUnsafe((_this).f32,_dafny.Map.Empty.slice().updateUnsafe((_this).f32,(_this).f31)));
      let _829_v21;
      let _nw134 = new _module.C3();
      _nw134.__ctor(_827_v19, (_828_v20).dtor_cf50, (_this).f31);
      _829_v21 = _nw134;
      _829_v21 = _829_v21;
      let _830_v22;
      _830_v22 = _dafny.Map.Empty.slice().updateUnsafe((_this).f32,_802_v4);
      _802_v4 = _dafny.Seq.Concat((((_830_v22).contains((_this).f32)) ? ((_830_v22).get((_this).f32)) : (_802_v4)), _802_v4);
      (globalState).f16 = (((_826_v18).contains(((_this).f32).isLessThanOrEqualTo((_this).f32))) ? ((_826_v18).get(((_this).f32).isLessThanOrEqualTo((_this).f32))) : (_this.f35));
      r0 = (_this).f32;
      let _831_v23;
      let _nw135 = Array((new BigNumber(8)).toNumber()).fill(_dafny.MultiSet.Empty);
      _831_v23 = _nw135;
      r1 = _831_v23;
      r2 = _dafny.Seq.UnicodeFromString("lsi");
      r3 = (_803_v5)[_module.__default.safeIndex((_this).f32, new BigNumber((_803_v5).length))];
      return [r0, r1, r2, r3];
    }
    get f34() {
      let _this = this;
      return _this._f34;
    };
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f29 = false;
      this._f30 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f29, f30) {
      let _this = this;
      (_this)._f29 = f29;
      (_this)._f30 = f30;
      return;
    }
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      if ((_this).f29) {
        let _832_v0;
        _832_v0 = new BigNumber(351);
        let _833_v1;
        _833_v1 = new _dafny.CodePoint('r'.codePointAt(0));
        _832_v0 = (_dafny.ZERO).minus(_module.__default.fm1(_833_v1, globalState));
        let _834_v2;
        _834_v2 = _dafny.Seq.UnicodeFromString("s");
        let _835_v3;
        _835_v3 = _dafny.Seq.of(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.update(_834_v2, _module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((_834_v2).length)), _833_v1), _834_v2), _module.__default.safeIndex(new BigNumber((_dafny.Seq.of(new BigNumber(826))).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_834_v2, _module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((_834_v2).length)), _833_v1), _834_v2)).length)), _833_v1));
        _835_v3 = _dafny.Seq.Concat(_835_v3, _835_v3);
        let _836_v4;
        let _init16 = ((_837_v0) => function (_838_i0) {
          return (_838_i0).plus(_837_v0);
        })(_832_v0);
        let _nw136 = Array((new BigNumber(14)).toNumber());
        for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw136.length); _i0_16++) {
          _nw136[_i0_16] = _init16(new BigNumber(_i0_16));
        }
        _836_v4 = _nw136;
        let _index119 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_836_v4).length));
        (_836_v4)[_index119] = _832_v0;
        let _index120 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_836_v4).length));
        (_836_v4)[_index120] = p1;
        _832_v0 = (_836_v4)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_836_v4).length))];
        let _index121 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_836_v4).length));
        (_836_v4)[_index121] = p1;
      } else {
        let _839_v5;
        _839_v5 = _dafny.MultiSet.fromElements(p1);
        _module.__default.m0((((_this).f30) ? (new BigNumber((_839_v5).cardinality())) : (p1)), globalState);
        let _840_v6;
        _840_v6 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f30);
        if (!((((_840_v6).contains(p1)) ? ((_840_v6).get(p1)) : ((_this).f29))) || ((_this).f30)) {
          let _841_v7;
          _841_v7 = _dafny.MultiSet.fromElements((_this).f29);
          let _842_v8;
          _842_v8 = _dafny.Seq.of(_841_v7, _841_v7);
          let _843_v9;
          _843_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_842_v8).length),new BigNumber((_module.__default.fm2((_this).f29, p1, (_this).f30, (_this).f29, globalState)).length));
          let _844_v10;
          _844_v10 = _dafny.Seq.of(_843_v9);
          _844_v10 = _844_v10;
          (globalState).f28 = (p1).isEqualTo(p1);
          let _845_v11;
          _845_v11 = new BigNumber(497);
          let _846_v12;
          _846_v12 = _dafny.Set.fromElements((_this).f30, (_this).f29);
          let _847_v13;
          _847_v13 = _dafny.Set.fromElements(new BigNumber(410), new BigNumber((_846_v12).length), _845_v11);
          let _848_v14;
          _848_v14 = _dafny.Seq.of((_this).f29, true);
          _845_v11 = (new BigNumber((_847_v13).length)).plus(new BigNumber((_848_v14).length));
          let _849_v15;
          _849_v15 = _dafny.Map.Empty.slice().updateUnsafe(((_module.__default.fm3(globalState)) ? (true) : (_module.__default.fm3(globalState))),p2);
          _849_v15 = (_849_v15).update((_this).f30, p2);
          let _850_v16;
          let _nw137 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
          _850_v16 = _nw137;
          let _index122 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_850_v16).length));
          (_850_v16)[_index122] = p1;
          let _index123 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_850_v16).length));
          (_850_v16)[_index123] = (_dafny.ZERO).minus(p1);
        } else {
          (globalState).f28 = (p1).isEqualTo(p1);
          let _851_v17;
          _851_v17 = _dafny.MultiSet.fromElements(_module.__default.fm3(globalState), (_this).f30, (_this).f29, (_this).f29, (_this).f29);
          let _852_v18;
          _852_v18 = _dafny.Seq.of(_851_v17);
          let _853_v19;
          _853_v19 = _module.D3.create_DC9(_852_v18);
          let _854_v20;
          _854_v20 = _module.D0.create_DC2(true);
          let _855_v21;
          let _nw138 = Array((new BigNumber(12)).toNumber());
          _nw138[(_dafny.ZERO).toNumber()] = _852_v18;
          _nw138[(_dafny.ONE).toNumber()] = _852_v18;
          _nw138[(new BigNumber(2)).toNumber()] = _852_v18;
          _nw138[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_852_v18, _module.__default.fm4((_this).f30, new BigNumber(700), globalState));
          _nw138[(new BigNumber(4)).toNumber()] = _dafny.Seq.update(_module.__default.fm4(!((_this).f30), p1, globalState), _module.__default.safeIndex(p1, new BigNumber((_module.__default.fm4(!((_this).f30), p1, globalState)).length)), _851_v17);
          _nw138[(new BigNumber(5)).toNumber()] = _module.__default.fm4((_this).f30, p1, globalState);
          _nw138[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(654)), ((_856_v17) => function (_857_i1) {
            return _856_v17;
          })(_851_v17)), _852_v18);
          _nw138[(new BigNumber(7)).toNumber()] = (_853_v19).dtor_cf12;
          _nw138[(new BigNumber(8)).toNumber()] = _module.__default.fm4((_854_v20).dtor_cf4, new BigNumber(-70), globalState);
          _nw138[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_851_v17, _851_v17), _852_v18);
          _nw138[(new BigNumber(10)).toNumber()] = _852_v18;
          _nw138[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_852_v18, _852_v18);
          _855_v21 = _nw138;
          let _index124 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_855_v21).length));
          (_855_v21)[_index124] = _dafny.Seq.update(_852_v18, _module.__default.safeIndex(p1, new BigNumber((_852_v18).length)), _851_v17);
          let _index125 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_855_v21).length));
          (_855_v21)[_index125] = _dafny.Seq.of(_851_v17, _851_v17, (_851_v17).Intersect(_851_v17), _851_v17, _851_v17);
          let _858_v22;
          _858_v22 = _dafny.Seq.UnicodeFromString("fudavpasx");
          let _859_v23;
          _859_v23 = _dafny.Set.fromElements((_this).f30);
          let _860_v24;
          _860_v24 = _module.D2.create_DC7((_858_v22)[_module.__default.safeIndex(p1, new BigNumber((_858_v22).length))], _859_v23);
          let _861_v25;
          let _nw139 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
          _861_v25 = _nw139;
          let _rhs124 = ((_this).f30) || ((_this).f29);
          let _rhs125 = _860_v24;
          let _rhs126 = (new BigNumber(-849)).isLessThan((p1).plus((_dafny.ZERO).minus(p1)));
          let _rhs127 = !(_module.__default.safeModuloInt(p1, p1)).isEqualTo(p1);
          let _rhs128 = _861_v25;
          let _lhs76 = globalState;
          let _lhs77 = globalState;
          let _lhs78 = globalState;
          _lhs76.f28 = _rhs124;
          _860_v24 = _rhs125;
          _lhs77.f16 = _rhs126;
          r0 = _rhs127;
          _lhs78.f24 = _rhs128;
          let _862_v26;
          _862_v26 = _module.D2.create_DC6((((_this).f30) ? (p1) : (p1)));
          _862_v26 = _862_v26;
          let _863_v27;
          _863_v27 = new BigNumber(-488);
          _863_v27 = _863_v27;
        }
        let _864_v28;
        _864_v28 = new _dafny.CodePoint('d'.codePointAt(0));
        let _865_v29;
        let _nw140 = Array((new BigNumber(13)).toNumber()).fill(false);
        _865_v29 = _nw140;
        let _866_v30;
        _866_v30 = _dafny.Map.Empty.slice().updateUnsafe(_864_v28,_865_v29);
        let _867_v31;
        _867_v31 = _dafny.Seq.of(_865_v29);
        _866_v30 = (_866_v30).update(_864_v28, (_867_v31)[_module.__default.safeIndex(p1, new BigNumber((_867_v31).length))]);
        (globalState).f16 = (_this).f29;
        let _868_v32;
        _868_v32 = new BigNumber(277);
        _868_v32 = p1;
      }
      let _869_v33;
      let _nw141 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.of());
      _869_v33 = _nw141;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_869_v33).length))) {
        let _870_i2 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_870_i2)) && ((_870_i2).isLessThan(new BigNumber((_869_v33).length))))) {
          (_869_v33)[(_870_i2)] = _dafny.Seq.of((_this).f30, (_this).f29, false, (_dafny.MultiSet.fromElements((_this).f29, (_this).f29)).IsSubsetOf(_dafny.MultiSet.fromElements((_this).f30)), ((_dafny.Set.fromElements(_module.D2.create_DC6(p1)))).IsDisjointFrom(function () {
            let _coll24 = new _dafny.Set();
            for (const _compr_24 of (_dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC6(new BigNumber((_dafny.Seq.UnicodeFromString("ho")).length)),p1)).Keys.Elements) {
              let _871_v34 = _compr_24;
              if ((_dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC6(new BigNumber((_dafny.Seq.UnicodeFromString("ho")).length)),p1)).contains(_871_v34)) {
                _coll24.add(_871_v34);
              }
            }
            return _coll24;
          }()));
        }
      }
      let _872_v35;
      _872_v35 = new _dafny.CodePoint('u'.codePointAt(0));
      let _873_v36;
      _873_v36 = _dafny.Set.fromElements((_this).f30, true);
      let _874_v37;
      _874_v37 = _module.D2.create_DC7(_872_v35, _873_v36);
      let _pat_let_tv19 = _873_v36;
      _872_v35 = (function (_pat_let22_0) {
        return function (_875_dt__update__tmp_h0) {
          return function (_pat_let23_0) {
            return function (_876_dt__update_hcf10_h0) {
              return _module.D2.create_DC7((_875_dt__update__tmp_h0).dtor_cf9, _876_dt__update_hcf10_h0);
            }(_pat_let23_0);
          }(_pat_let_tv19);
        }(_pat_let22_0);
      }(_874_v37)).dtor_cf9;
      let _877_v38;
      _877_v38 = _dafny.Seq.UnicodeFromString("vjubcpq");
      let _878_v39;
      _878_v39 = _module.D3.create_DC10(_877_v38, p1, (_this).f30);
      let _879_v40;
      _879_v40 = _module.D3.create_DC11(_878_v39);
      let _880_v41;
      _880_v41 = _module.D3.create_DC11(_879_v40);
      let _source13 = _880_v41;
      if (_source13.is_DC10) {
        let _881___mcc_h0 = (_source13).cf13;
        let _882___mcc_h1 = (_source13).cf14;
        let _883___mcc_h2 = (_source13).cf15;
        let _884_cf15 = _883___mcc_h2;
        let _885_cf14 = _882___mcc_h1;
        let _886_cf13 = _881___mcc_h0;
        _885_cf14 = _885_cf14;
        let _887_v42;
        let _nw142 = Array((_dafny.ONE).toNumber()).fill(_module.D0.Default());
        _887_v42 = _nw142;
        let _888_v43;
        _888_v43 = _dafny.Set.fromElements(_887_v42);
        _884_cf15 = (_888_v43).IsDisjointFrom(_888_v43);
        let _889_v44;
        let _init17 = ((_890_cf14) => function (_891_i3) {
          return _module.__default.safeDivisionInt(_891_i3, _890_cf14);
        })(_885_cf14);
        let _nw143 = Array((new BigNumber(28)).toNumber());
        for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw143.length); _i0_17++) {
          _nw143[_i0_17] = _init17(new BigNumber(_i0_17));
        }
        _889_v44 = _nw143;
        let _892_v45;
        _892_v45 = _889_v44;
        let _source14 = _892_v45;
        let _893___mcc_h5 = _source14;
        let _894_cf7 = _893___mcc_h5;
        let _895_v46;
        let _init18 = ((_896_cf15) => function (_897_i4) {
          return !(!(_896_cf15));
        })(_884_cf15);
        let _nw144 = Array((new BigNumber(4)).toNumber());
        for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw144.length); _i0_18++) {
          _nw144[_i0_18] = _init18(new BigNumber(_i0_18));
        }
        _895_v46 = _nw144;
        (globalState).f27 = _895_v46;
        let _898_v47;
        _898_v47 = _dafny.Seq.of(_877_v38);
        let _899_v48;
        _899_v48 = _module.D3.create_DC10((_898_v47)[_module.__default.safeIndex(p1, new BigNumber((_898_v47).length))], p1, _884_cf15);
        let _900_v49;
        _900_v49 = _dafny.Seq.of(true);
        (globalState).f16 = (((_884_cf15) ? (_899_v48) : (_module.D3.create_DC10(_877_v38, new BigNumber((_900_v49).length), _884_cf15)))).dtor_cf15;
        let _901_v50;
        _901_v50 = _dafny.Map.Empty.slice().updateUnsafe(_877_v38,false);
        let _902_v51;
        _902_v51 = _dafny.MultiSet.fromElements((_this).f29);
        let _903_v52;
        _903_v52 = _dafny.Map.Empty.slice().updateUnsafe(_884_cf15,_module.__default.fm0(_902_v51, globalState));
        _901_v50 = (_901_v50).update((((_903_v52).contains((_this).f29)) ? ((_903_v52).get((_this).f29)) : (_877_v38)), (_this).f29);
        _884_cf15 = !(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(46)), ((_904_p1) => function (_905_i5) {
          return (_dafny.ZERO).minus((_dafny.ZERO).minus(_904_p1));
        })(p1)), p0)).length)).isEqualTo((_dafny.ZERO).minus(new BigNumber((_877_v38).length)));
        let _906_v53;
        _906_v53 = _dafny.MultiSet.fromElements(p0);
        let _907_v54;
        _907_v54 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm5((_this).f29, _885_cf14, globalState),(_this).f29);
        let _908_v55;
        _908_v55 = _dafny.MultiSet.fromElements((_this).f30);
        let _909_v57;
        _909_v57 = _dafny.Set.fromElements(_872_v35);
        let _rhs129 = (_this).f29;
        let _rhs130 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(p0, p0));
        let _rhs131 = (p1).minus(_module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_908_v55,(_this).f29)).length), (p0)[_module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus(p1)), new BigNumber((p0).length))]));
        let _rhs132 = function () {
          let _coll25 = new _dafny.Map();
          for (const _compr_25 of (_909_v57).Elements) {
            let _910_v56 = _compr_25;
            if ((_909_v57).contains(_910_v56)) {
              _coll25.push([_910_v56,false]);
            }
          }
          return _coll25;
        }();
        let _rhs133 = (new BigNumber(387)).isLessThan(_885_cf14);
        let _lhs79 = globalState;
        let _lhs80 = globalState;
        _lhs79.f28 = _rhs129;
        _906_v53 = _rhs130;
        _885_cf14 = _rhs131;
        _907_v54 = _rhs132;
        _lhs80.f28 = _rhs133;
      } else if (_source13.is_DC9) {
        let _911___mcc_h3 = (_source13).cf12;
        let _912_cf12 = _911___mcc_h3;
        let _913_v58;
        _913_v58 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f30);
        let _914_v59;
        _914_v59 = _dafny.MultiSet.fromElements((_this).f30, false);
        let _915_v60;
        _915_v60 = _dafny.MultiSet.fromElements(p1, new BigNumber((_913_v58).length), p1, new BigNumber((_914_v59).cardinality()), new BigNumber(364));
        let _916_v61;
        _916_v61 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_915_v60).cardinality()),(_this).f29);
        let _917_v62;
        _917_v62 = _module.D2.create_DC6(p1);
        let _918_v63;
        _918_v63 = _dafny.Set.fromElements(_917_v62, _917_v62);
        let _919_v64;
        _919_v64 = _918_v63;
        let _920_v65;
        _920_v65 = _dafny.Seq.of(_919_v64, _919_v64, _919_v64, _919_v64, _dafny.Set.fromElements(_917_v62));
        let _921_v66;
        _921_v66 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_916_v61).Merge(_916_v61)).length),_920_v65);
        let _922_v67;
        _922_v67 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p1),_920_v65)).update(new BigNumber((p0).length), _920_v65));
        _921_v66 = (((_922_v67).contains((_this).f30)) ? ((_922_v67).get((_this).f30)) : (_921_v66));
        (globalState).f28 = !((_this).f29) || ((_this).f30);
        r0 = !((new BigNumber(453)).plus(p1)).isEqualTo(p1);
        let _rhs134 = (_this).f30;
        let _rhs135 = (_this).f29;
        let _lhs81 = globalState;
        let _lhs82 = globalState;
        _lhs81.f16 = _rhs134;
        _lhs82.f28 = _rhs135;
      } else {
        let _923___mcc_h4 = (_source13).cf16;
        let _924_cf16 = _923___mcc_h4;
        let _925_v68;
        _925_v68 = new BigNumber(435);
        let _926_v69;
        _926_v69 = _dafny.MultiSet.fromElements((_this).f30);
        let _927_v70;
        _927_v70 = _dafny.Map.Empty.slice().updateUnsafe(_926_v69,(p0)[_module.__default.safeIndex(_925_v68, new BigNumber((p0).length))]);
        _925_v68 = (_dafny.ZERO).minus((((_this).f30) ? (_925_v68) : ((((_927_v70).contains(_926_v69)) ? ((_927_v70).get(_926_v69)) : (p1)))));
        let _928_v71;
        _928_v71 = _dafny.Seq.of(p2, p2);
        let _929_v72;
        _929_v72 = _dafny.Map.Empty.slice().updateUnsafe(p1,true);
        let _930_v73;
        let _nw145 = Array((new BigNumber(22)).toNumber());
        _nw145[(_dafny.ZERO).toNumber()] = (_this).f29;
        _nw145[(_dafny.ONE).toNumber()] = false;
        _nw145[(new BigNumber(2)).toNumber()] = false;
        _nw145[(new BigNumber(3)).toNumber()] = (_this).f30;
        _nw145[(new BigNumber(4)).toNumber()] = (_this).f30;
        _nw145[(new BigNumber(5)).toNumber()] = !(true);
        _nw145[(new BigNumber(6)).toNumber()] = (_this).f29;
        _nw145[(new BigNumber(7)).toNumber()] = (_this).f29;
        _nw145[(new BigNumber(8)).toNumber()] = (_this).f30;
        _nw145[(new BigNumber(9)).toNumber()] = (_this).f29;
        _nw145[(new BigNumber(10)).toNumber()] = (_this).f29;
        _nw145[(new BigNumber(11)).toNumber()] = !((_this).f29);
        _nw145[(new BigNumber(12)).toNumber()] = (_this).f30;
        _nw145[(new BigNumber(13)).toNumber()] = (((_929_v72).contains(_925_v68)) ? ((_929_v72).get(_925_v68)) : ((_this).f30));
        _nw145[(new BigNumber(14)).toNumber()] = (_this).f29;
        _nw145[(new BigNumber(15)).toNumber()] = (_this).f29;
        _nw145[(new BigNumber(16)).toNumber()] = (((_929_v72).contains(_925_v68)) ? ((_929_v72).get(_925_v68)) : ((_this).f29));
        _nw145[(new BigNumber(17)).toNumber()] = (_this).f30;
        _nw145[(new BigNumber(18)).toNumber()] = (_this).f29;
        _nw145[(new BigNumber(19)).toNumber()] = (_this).f29;
        _nw145[(new BigNumber(20)).toNumber()] = true;
        _nw145[(new BigNumber(21)).toNumber()] = true;
        _930_v73 = _nw145;
        let _931_v74;
        _931_v74 = _dafny.Seq.of((_this).f30);
        let _932_v75;
        _932_v75 = _dafny.Map.Empty.slice().updateUnsafe(_931_v74,_930_v73);
        let _933_v76;
        _933_v76 = _dafny.Set.fromElements(_925_v68);
        let _934_v77;
        _934_v77 = _module.D5.create_DC14(_933_v76, p1, (_this).f30, _930_v73);
        let _935_v78;
        let _nw146 = Array((new BigNumber(29)).toNumber());
        _nw146[(_dafny.ZERO).toNumber()] = _930_v73;
        _nw146[(_dafny.ONE).toNumber()] = _930_v73;
        _nw146[(new BigNumber(2)).toNumber()] = _930_v73;
        _nw146[(new BigNumber(3)).toNumber()] = _930_v73;
        _nw146[(new BigNumber(4)).toNumber()] = _930_v73;
        _nw146[(new BigNumber(5)).toNumber()] = _930_v73;
        _nw146[(new BigNumber(6)).toNumber()] = _930_v73;
        _nw146[(new BigNumber(7)).toNumber()] = _930_v73;
        _nw146[(new BigNumber(8)).toNumber()] = _930_v73;
        _nw146[(new BigNumber(9)).toNumber()] = _930_v73;
        _nw146[(new BigNumber(10)).toNumber()] = _930_v73;
        _nw146[(new BigNumber(11)).toNumber()] = _930_v73;
        _nw146[(new BigNumber(12)).toNumber()] = _930_v73;
        _nw146[(new BigNumber(13)).toNumber()] = _930_v73;
        _nw146[(new BigNumber(14)).toNumber()] = _930_v73;
        _nw146[(new BigNumber(15)).toNumber()] = (((_932_v75).contains(_931_v74)) ? ((_932_v75).get(_931_v74)) : (_930_v73));
        _nw146[(new BigNumber(16)).toNumber()] = _930_v73;
        _nw146[(new BigNumber(17)).toNumber()] = (_934_v77).dtor_cf22;
        _nw146[(new BigNumber(18)).toNumber()] = _930_v73;
        _nw146[(new BigNumber(19)).toNumber()] = _930_v73;
        _nw146[(new BigNumber(20)).toNumber()] = _930_v73;
        _nw146[(new BigNumber(21)).toNumber()] = _930_v73;
        _nw146[(new BigNumber(22)).toNumber()] = _930_v73;
        _nw146[(new BigNumber(23)).toNumber()] = _930_v73;
        _nw146[(new BigNumber(24)).toNumber()] = _930_v73;
        _nw146[(new BigNumber(25)).toNumber()] = _930_v73;
        _nw146[(new BigNumber(26)).toNumber()] = _930_v73;
        _nw146[(new BigNumber(27)).toNumber()] = _930_v73;
        _nw146[(new BigNumber(28)).toNumber()] = _930_v73;
        _935_v78 = _nw146;
        let _936_v79;
        _936_v79 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,_925_v68);
        let _937_v81;
        let _nw147 = new _module.C4();
        _nw147.__ctor(_928_v71, (_this).f30, p1, _935_v78, !((function () {
          let _coll26 = new _dafny.Set();
          for (const _compr_26 of (_dafny.MultiSet.fromElements(_936_v79, _936_v79)).Elements) {
            let _938_v80 = _compr_26;
            if ((_dafny.MultiSet.fromElements(_936_v79, _936_v79)).contains(_938_v80)) {
              _coll26.add(_938_v80);
            }
          }
          return _coll26;
        }()).IsDisjointFrom(_dafny.Set.fromElements(_936_v79))));
        _937_v81 = _nw147;
        let _index126 = _module.__default.safeIndex(new BigNumber(848), new BigNumber((_930_v73).length));
        (_930_v73)[_index126] = !((_this).f30);
        let _index127 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_935_v78).length));
        (_935_v78)[_index127] = _930_v73;
        let _939_v82;
        _939_v82 = _dafny.MultiSet.fromElements(new BigNumber(280), _925_v68, _925_v68);
        let _940_v83;
        let _nw148 = Array((new BigNumber(5)).toNumber());
        _nw148[(_dafny.ZERO).toNumber()] = _925_v68;
        _nw148[(_dafny.ONE).toNumber()] = p1;
        _nw148[(new BigNumber(2)).toNumber()] = new BigNumber((_877_v38).length);
        _nw148[(new BigNumber(3)).toNumber()] = p1;
        _nw148[(new BigNumber(4)).toNumber()] = new BigNumber((_877_v38).length);
        _940_v83 = _nw148;
        let _941_v84;
        _941_v84 = _dafny.Map.Empty.slice().updateUnsafe(_940_v83,_930_v73);
        let _index128 = _module.__default.safeIndex(new BigNumber(848), new BigNumber((_930_v73).length));
        let _index129 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_935_v78).length));
        let _rhs136 = !_dafny.Seq.contains(_877_v38, _872_v35);
        let _rhs137 = (_939_v82).IsDisjointFrom(_939_v82);
        let _rhs138 = (((_941_v84).contains(_940_v83)) ? ((_941_v84).get(_940_v83)) : (_930_v73));
        let _rhs139 = ((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_925_v68, new BigNumber(484)))).isLessThan(_module.__default.safeModuloInt(new BigNumber((p0).length), new BigNumber(6)));
        let _rhs140 = _877_v38;
        let _lhs83 = _930_v73;
        let _lhs84 = _module.__default.safeIndex(new BigNumber(848), new BigNumber((_930_v73).length));
        let _lhs85 = globalState;
        let _lhs86 = _935_v78;
        let _lhs87 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_935_v78).length));
        let _lhs88 = _937_v81;
        _lhs83[_lhs84] = _rhs136;
        _lhs85.f16 = _rhs137;
        _lhs86[_lhs87] = _rhs138;
        _lhs88.f35 = _rhs139;
        _877_v38 = _rhs140;
        _925_v68 = p1;
      }
      let _942_v85;
      _942_v85 = _dafny.Map.Empty.slice().updateUnsafe(p1,_module.__default.fm3(globalState));
      let _943_v86;
      _943_v86 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,(_this).f29);
      let _944_i6;
      _944_i6 = _dafny.ZERO;
      L9: {
        while ((((_942_v85).contains(p1)) ? ((_942_v85).get(p1)) : ((new BigNumber((_877_v38).length)).isLessThanOrEqualTo(new BigNumber((_943_v86).length))))) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_944_i6)) {
              break L9;
            }
            _944_i6 = (_944_i6).plus(_dafny.ONE);
            let _945_v87;
            let _nw149 = Array((new BigNumber(7)).toNumber()).fill(_module.D6.Default());
            _945_v87 = _nw149;
            let _946_v88;
            let _init19 = ((_947_p1) => function (_948_i7) {
              return _dafny.Set.fromElements(_947_p1, _947_p1);
            })(p1);
            let _nw150 = Array((new BigNumber(18)).toNumber());
            for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw150.length); _i0_19++) {
              _nw150[_i0_19] = _init19(new BigNumber(_i0_19));
            }
            _946_v88 = _nw150;
            let _949_v89;
            _949_v89 = _dafny.MultiSet.fromElements(p1);
            let _950_v90;
            _950_v90 = _module.D6.create_DC17((_this).f30, _946_v88, new BigNumber((_949_v89).cardinality()));
            let _index130 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_945_v87).length));
            (_945_v87)[_index130] = _950_v90;
            let _index131 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_945_v87).length));
            (_945_v87)[_index131] = _950_v90;
            let _951_v91;
            _951_v91 = _module.D0.create_DC0(_869_v33);
            let _952_v92;
            _952_v92 = _module.D0.create_DC4(_951_v91);
            let _953_v93;
            _953_v93 = _dafny.Map.Empty.slice().updateUnsafe(_952_v92,_module.__default.fm1(_872_v35, globalState));
            (globalState).f28 = !(_953_v93).contains((((_this).f29) ? (_952_v92) : (_952_v92)));
            let _954_v94;
            _954_v94 = _dafny.MultiSet.fromElements((_this).f30);
            let _955_v95;
            _955_v95 = _dafny.Seq.of((_this).f30, (_this).f29, (_this).f30, (_this).f30);
            let _956_v96;
            _956_v96 = _dafny.Seq.of(_942_v85, (_942_v85).Merge(_942_v85), _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f30), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_954_v94).cardinality()),(_955_v95)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_955_v95).length))]), _942_v85);
            _942_v85 = (_956_v96)[_module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((_956_v96).length))];
            let _957_v97;
            let _nw151 = Array((new BigNumber(28)).toNumber()).fill(false);
            _957_v97 = _nw151;
            let _index132 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_957_v97).length));
            (_957_v97)[_index132] = _module.__default.fm10(_872_v35, false, globalState);
            let _index133 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_957_v97).length));
            (_957_v97)[_index133] = (_this).f29;
          }
        }
      }
      let _958_v98;
      _958_v98 = _dafny.Seq.of(_877_v38, _877_v38);
      let _959_v99;
      _959_v99 = _module.D9.create_DC23(_958_v98);
      let _source15 = _959_v99;
      if (_source15.is_DC24) {
        let _960___mcc_h6 = (_source15).cf40;
        let _961_cf40 = _960___mcc_h6;
        let _962_v100;
        _962_v100 = _dafny.Seq.of(p2);
        let _963_v101;
        let _nw152 = Array((new BigNumber(26)).toNumber());
        _nw152[(_dafny.ZERO).toNumber()] = (_this).f29;
        _nw152[(_dafny.ONE).toNumber()] = (_this).f30;
        _nw152[(new BigNumber(2)).toNumber()] = (_this).f29;
        _nw152[(new BigNumber(3)).toNumber()] = (_this).f29;
        _nw152[(new BigNumber(4)).toNumber()] = (_this).f30;
        _nw152[(new BigNumber(5)).toNumber()] = (_this).f30;
        _nw152[(new BigNumber(6)).toNumber()] = (_this).f30;
        _nw152[(new BigNumber(7)).toNumber()] = (_this).f29;
        _nw152[(new BigNumber(8)).toNumber()] = true;
        _nw152[(new BigNumber(9)).toNumber()] = (_this).f29;
        _nw152[(new BigNumber(10)).toNumber()] = (_this).f29;
        _nw152[(new BigNumber(11)).toNumber()] = (_this).f30;
        _nw152[(new BigNumber(12)).toNumber()] = (_this).f30;
        _nw152[(new BigNumber(13)).toNumber()] = (_this).f29;
        _nw152[(new BigNumber(14)).toNumber()] = (_this).f29;
        _nw152[(new BigNumber(15)).toNumber()] = (_this).f30;
        _nw152[(new BigNumber(16)).toNumber()] = (_this).f30;
        _nw152[(new BigNumber(17)).toNumber()] = (_this).f29;
        _nw152[(new BigNumber(18)).toNumber()] = (_this).f29;
        _nw152[(new BigNumber(19)).toNumber()] = (_this).f30;
        _nw152[(new BigNumber(20)).toNumber()] = (_this).f29;
        _nw152[(new BigNumber(21)).toNumber()] = (_this).f30;
        _nw152[(new BigNumber(22)).toNumber()] = !((_this).f29);
        _nw152[(new BigNumber(23)).toNumber()] = true;
        _nw152[(new BigNumber(24)).toNumber()] = false;
        _nw152[(new BigNumber(25)).toNumber()] = (_this).f29;
        _963_v101 = _nw152;
        let _964_v102;
        _964_v102 = _dafny.Seq.of(_963_v101);
        let _965_v103;
        let _nw153 = Array((new BigNumber(28)).toNumber());
        _nw153[(_dafny.ZERO).toNumber()] = _963_v101;
        _nw153[(_dafny.ONE).toNumber()] = _963_v101;
        _nw153[(new BigNumber(2)).toNumber()] = _963_v101;
        _nw153[(new BigNumber(3)).toNumber()] = _963_v101;
        _nw153[(new BigNumber(4)).toNumber()] = _963_v101;
        _nw153[(new BigNumber(5)).toNumber()] = _963_v101;
        _nw153[(new BigNumber(6)).toNumber()] = _963_v101;
        _nw153[(new BigNumber(7)).toNumber()] = (_964_v102)[_module.__default.safeIndex(new BigNumber((p0).length), new BigNumber((_964_v102).length))];
        _nw153[(new BigNumber(8)).toNumber()] = _963_v101;
        _nw153[(new BigNumber(9)).toNumber()] = _963_v101;
        _nw153[(new BigNumber(10)).toNumber()] = _963_v101;
        _nw153[(new BigNumber(11)).toNumber()] = _963_v101;
        _nw153[(new BigNumber(12)).toNumber()] = _963_v101;
        _nw153[(new BigNumber(13)).toNumber()] = _963_v101;
        _nw153[(new BigNumber(14)).toNumber()] = _963_v101;
        _nw153[(new BigNumber(15)).toNumber()] = _963_v101;
        _nw153[(new BigNumber(16)).toNumber()] = _963_v101;
        _nw153[(new BigNumber(17)).toNumber()] = _963_v101;
        _nw153[(new BigNumber(18)).toNumber()] = _963_v101;
        _nw153[(new BigNumber(19)).toNumber()] = _963_v101;
        _nw153[(new BigNumber(20)).toNumber()] = _963_v101;
        _nw153[(new BigNumber(21)).toNumber()] = _963_v101;
        _nw153[(new BigNumber(22)).toNumber()] = (_964_v102)[_module.__default.safeIndex(p1, new BigNumber((_964_v102).length))];
        _nw153[(new BigNumber(23)).toNumber()] = _963_v101;
        _nw153[(new BigNumber(24)).toNumber()] = _963_v101;
        _nw153[(new BigNumber(25)).toNumber()] = _963_v101;
        _nw153[(new BigNumber(26)).toNumber()] = _963_v101;
        _nw153[(new BigNumber(27)).toNumber()] = _963_v101;
        _965_v103 = _nw153;
        let _966_v104;
        let _nw154 = new _module.C4();
        _nw154.__ctor(_962_v100, (new BigNumber((_961_cf40).cardinality())).isEqualTo((_dafny.ZERO).minus(p1)), new BigNumber((_877_v38).length), _965_v103, (_this).f30);
        _966_v104 = _nw154;
        let _967_v105;
        _967_v105 = new BigNumber(570);
        _967_v105 = _967_v105;
        let _968_v106;
        let _nw155 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
        _968_v106 = _nw155;
        let _969_v107;
        _969_v107 = _dafny.Set.fromElements(_968_v106);
        _969_v107 = _dafny.Set.fromElements(_968_v106, _968_v106);
        let _970_v108;
        let _nw156 = new _module.C2();
        _nw156.__ctor(_module.__default.fm14(p1, _967_v105, new BigNumber(372), (_this).f29, globalState), (_this).f29);
        _970_v108 = _nw156;
      } else if (_source15.is_DC25) {
        let _971___mcc_h7 = (_source15).cf41;
        let _972___mcc_h8 = (_source15).cf42;
        let _973___mcc_h9 = (_source15).cf43;
        let _974___mcc_h10 = (_source15).cf44;
        let _975_cf44 = _974___mcc_h10;
        let _976_cf43 = _973___mcc_h9;
        let _977_cf42 = _972___mcc_h8;
        let _978_cf41 = _971___mcc_h7;
        let _979_v109;
        _979_v109 = _module.D10.create_DC28(p1);
        let _source16 = _979_v109;
        if (_source16.is_DC28) {
          let _980___mcc_h17 = (_source16).cf51;
          let _981_cf51 = _980___mcc_h17;
          let _982_v110;
          let _nw157 = Array((new BigNumber(13)).toNumber()).fill([]);
          _982_v110 = _nw157;
          _982_v110 = _982_v110;
          let _983_v111;
          let _nw158 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
          _983_v111 = _nw158;
          let _984_v112;
          _984_v112 = _983_v111;
          _984_v112 = _984_v112;
          let _985_v113;
          _985_v113 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,_module.__default.safeDivisionInt(_981_cf51, _981_cf51));
          _985_v113 = (_985_v113).update((_this).f29, _981_cf51);
          _877_v38 = _877_v38;
        } else if (_source16.is_DC27) {
          let _986___mcc_h18 = (_source16).cf50;
          let _987_cf50 = _986___mcc_h18;
          let _988_v114;
          let _nw159 = Array((new BigNumber(3)).toNumber());
          _nw159[(_dafny.ZERO).toNumber()] = new BigNumber(430);
          _nw159[(_dafny.ONE).toNumber()] = new BigNumber(-353);
          _nw159[(new BigNumber(2)).toNumber()] = new BigNumber(-342);
          _988_v114 = _nw159;
          let _index134 = _module.__default.safeIndex(new BigNumber(357), new BigNumber((_988_v114).length));
          (_988_v114)[_index134] = (((_this).f30) ? (_978_cf41) : (_978_cf41));
          let _index135 = _module.__default.safeIndex(new BigNumber(357), new BigNumber((_988_v114).length));
          (_988_v114)[_index135] = new BigNumber((_dafny.Seq.Concat(_877_v38, _dafny.Seq.Concat((_958_v98)[_module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((_958_v98).length))], _877_v38))).length);
          let _989_v115;
          _989_v115 = _dafny.Seq.of(true);
          _989_v115 = _989_v115;
          (globalState).f16 = _975_cf44;
          let _990_v116;
          _990_v116 = _dafny.Map.Empty.slice().updateUnsafe((_this).f29,p1);
          let _index136 = _module.__default.safeIndex(new BigNumber(357), new BigNumber((_988_v114).length));
          (_988_v114)[_index136] = (((_976_cf43) ? (new BigNumber((_990_v116).length)) : (new BigNumber((_877_v38).length)))).multipliedBy((_978_cf41).multipliedBy(_module.__default.fm18((((_942_v85).contains(p1)) ? ((_942_v85).get(p1)) : ((_this).f30)), (_this).f29, globalState)));
        } else {
          let _991___mcc_h19 = (_source16).cf52;
          let _992_cf52 = _991___mcc_h19;
          _975_cf44 = _module.__default.fm8((_978_cf41).isLessThanOrEqualTo(_978_cf41), new BigNumber(-349), globalState);
          _978_cf41 = _978_cf41;
          let _993_v117;
          _993_v117 = _dafny.Seq.of(_978_cf41);
          _993_v117 = _dafny.Seq.of(_978_cf41, _module.__default.fm7(false, (_this).f30, globalState), new BigNumber(661));
          let _994_v118;
          let _nw160 = Array((new BigNumber(2)).toNumber());
          _nw160[(_dafny.ZERO).toNumber()] = !(_977_cf42);
          _nw160[(_dafny.ONE).toNumber()] = (_this).f30;
          _994_v118 = _nw160;
          let _995_v119;
          _995_v119 = _dafny.Map.Empty.slice().updateUnsafe(p1,_994_v118);
          let _996_v120;
          _996_v120 = _module.D8.create_DC20(_995_v119);
          let _rhs141 = (new BigNumber(-520)).isLessThan((_dafny.ZERO).minus(_module.__default.fm1(_872_v35, globalState)));
          let _rhs142 = _996_v120;
          let _rhs143 = (_this).f30;
          let _lhs89 = globalState;
          _lhs89.f28 = _rhs141;
          _996_v120 = _rhs142;
          _976_cf43 = _rhs143;
        }
        let _997_v122;
        let _init20 = ((_998_cf43) => function (_999_i8) {
          return _998_cf43;
        })(_976_cf43);
        let _nw161 = Array((new BigNumber(23)).toNumber());
        for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw161.length); _i0_20++) {
          _nw161[_i0_20] = _init20(new BigNumber(_i0_20));
        }
        _997_v122 = _nw161;
        let _1000_v123;
        _1000_v123 = _module.D5.create_DC14(function () {
  let _coll27 = new _dafny.Set();
  for (const _compr_27 of _dafny.IntegerRange(new BigNumber(-870), new BigNumber(990))) {
    let _1001_v121 = _compr_27;
    if (((new BigNumber(-870)).isLessThanOrEqualTo(_1001_v121)) && ((_1001_v121).isLessThan(new BigNumber(990)))) {
      _coll27.add(_module.__default.safeModuloInt(_1001_v121, p1));
    }
  }
  return _coll27;
}(), _978_cf41, !(_975_cf44), _997_v122);
        let _1002_v124;
        _1002_v124 = _dafny.Map.Empty.slice().updateUnsafe(_1000_v123,_872_v35);
        let _1003_v125;
        _1003_v125 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ouwdevy")).length));
        let _rhs144 = (((_1002_v124).contains(_module.D5.create_DC14(_1003_v125, p1, !((_this).f29), _997_v122))) ? ((_1002_v124).get(_module.D5.create_DC14(_1003_v125, p1, !((_this).f29), _997_v122))) : (_872_v35));
        let _rhs145 = (_978_cf41).minus(_module.__default.fm1(_872_v35, globalState));
        let _rhs146 = _module.__default.fm3(globalState);
        let _rhs147 = true;
        _872_v35 = _rhs144;
        _978_cf41 = _rhs145;
        _976_cf43 = _rhs146;
        _977_cf42 = _rhs147;
        let _1004_v126;
        _1004_v126 = _dafny.Map.Empty.slice().updateUnsafe(_978_cf41,_dafny.Set.fromElements(_978_cf41, _978_cf41));
        let _1005_v130;
        let _nw162 = Array((new BigNumber(29)).toNumber());
        _nw162[(_dafny.ZERO).toNumber()] = _1003_v125;
        _nw162[(_dafny.ONE).toNumber()] = _1003_v125;
        _nw162[(new BigNumber(2)).toNumber()] = _1003_v125;
        _nw162[(new BigNumber(3)).toNumber()] = _1003_v125;
        _nw162[(new BigNumber(4)).toNumber()] = _1003_v125;
        _nw162[(new BigNumber(5)).toNumber()] = _dafny.Set.fromElements(p1, p1);
        _nw162[(new BigNumber(6)).toNumber()] = _1003_v125;
        _nw162[(new BigNumber(7)).toNumber()] = _dafny.Set.fromElements(_978_cf41, p1, _978_cf41);
        _nw162[(new BigNumber(8)).toNumber()] = _1003_v125;
        _nw162[(new BigNumber(9)).toNumber()] = _1003_v125;
        _nw162[(new BigNumber(10)).toNumber()] = _1003_v125;
        _nw162[(new BigNumber(11)).toNumber()] = _1003_v125;
        _nw162[(new BigNumber(12)).toNumber()] = _dafny.Set.fromElements(_978_cf41, p1);
        _nw162[(new BigNumber(13)).toNumber()] = _1003_v125;
        _nw162[(new BigNumber(14)).toNumber()] = (_1000_v123).dtor_cf19;
        _nw162[(new BigNumber(15)).toNumber()] = _1003_v125;
        _nw162[(new BigNumber(16)).toNumber()] = _1003_v125;
        _nw162[(new BigNumber(17)).toNumber()] = _dafny.Set.fromElements(p1, (_dafny.ZERO).minus(_978_cf41));
        _nw162[(new BigNumber(18)).toNumber()] = _dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber(-73)));
        _nw162[(new BigNumber(19)).toNumber()] = (((_1004_v126).contains(_978_cf41)) ? ((_1004_v126).get(_978_cf41)) : (function () {
          let _coll28 = new _dafny.Set();
          for (const _compr_28 of (function () {
            let _coll29 = new _dafny.Set();
            for (const _compr_29 of _dafny.IntegerRange(new BigNumber(464), new BigNumber(420))) {
              let _1006_v127 = _compr_29;
              if (((new BigNumber(464)).isLessThanOrEqualTo(_1006_v127)) && ((_1006_v127).isLessThan(new BigNumber(420)))) {
                _coll29.add(_module.__default.safeDivisionInt(_1006_v127, p1));
              }
            }
            return _coll29;
          }()).Elements) {
            let _1007_v128 = _compr_28;
            if ((function () {
              let _coll30 = new _dafny.Set();
              for (const _compr_30 of _dafny.IntegerRange(new BigNumber(464), new BigNumber(420))) {
                let _1008_v127 = _compr_30;
                if (((new BigNumber(464)).isLessThanOrEqualTo(_1008_v127)) && ((_1008_v127).isLessThan(new BigNumber(420)))) {
                  _coll30.add(_module.__default.safeDivisionInt(_1008_v127, p1));
                }
              }
              return _coll30;
            }()).contains(_1007_v128)) {
              _coll28.add(_module.__default.safeModuloInt(_1007_v128, new BigNumber((_dafny.Seq.UnicodeFromString("r")).length)));
            }
          }
          return _coll28;
        }()));
        _nw162[(new BigNumber(20)).toNumber()] = _1003_v125;
        _nw162[(new BigNumber(21)).toNumber()] = _dafny.Set.fromElements(_978_cf41);
        _nw162[(new BigNumber(22)).toNumber()] = _dafny.Set.fromElements(p1);
        _nw162[(new BigNumber(23)).toNumber()] = _1003_v125;
        _nw162[(new BigNumber(24)).toNumber()] = _1003_v125;
        _nw162[(new BigNumber(25)).toNumber()] = _module.__default.fm28(new BigNumber((_943_v86).length), false, _975_cf44, globalState);
        _nw162[(new BigNumber(26)).toNumber()] = _1003_v125;
        _nw162[(new BigNumber(27)).toNumber()] = _1003_v125;
        _nw162[(new BigNumber(28)).toNumber()] = function () {
          let _coll31 = new _dafny.Set();
          for (const _compr_31 of _dafny.IntegerRange(new BigNumber(160), new BigNumber(-306))) {
            let _1009_v129 = _compr_31;
            if (((new BigNumber(160)).isLessThanOrEqualTo(_1009_v129)) && ((_1009_v129).isLessThan(new BigNumber(-306)))) {
              _coll31.add((_1009_v129).multipliedBy(_978_cf41));
            }
          }
          return _coll31;
        }();
        _1005_v130 = _nw162;
        let _1010_v131;
        _1010_v131 = _module.D6.create_DC17(_975_cf44, _1005_v130, p1);
        _975_cf44 = (_1010_v131).dtor_cf25;
        let _1011_v132;
        _1011_v132 = _dafny.Map.Empty.slice().updateUnsafe(_872_v35,(p0)[_module.__default.safeIndex(p1, new BigNumber((p0).length))]);
        let _1012_v133;
        _1012_v133 = _dafny.MultiSet.fromElements(p1);
        _1011_v132 = (_1011_v132).update(_module.__default.fm20(_877_v38, _1012_v133, false, new BigNumber(-38), globalState), _module.__default.safeModuloInt(_978_cf41, new BigNumber(150)));
      } else if (_source15.is_DC26) {
        let _1013___mcc_h11 = (_source15).cf45;
        let _1014___mcc_h12 = (_source15).cf46;
        let _1015___mcc_h13 = (_source15).cf47;
        let _1016___mcc_h14 = (_source15).cf48;
        let _1017___mcc_h15 = (_source15).cf49;
        let _1018_cf49 = _1017___mcc_h15;
        let _1019_cf48 = _1016___mcc_h14;
        let _1020_cf47 = _1015___mcc_h13;
        let _1021_cf46 = _1014___mcc_h12;
        let _1022_cf45 = _1013___mcc_h11;
        if (((_873_v36).Union(_873_v36)).equals(_dafny.Set.fromElements(_1018_cf49))) {
          let _1023_v134;
          _1023_v134 = _dafny.MultiSet.fromElements(!(_1018_cf49));
          let _1024_v135;
          let _nw163 = Array((new BigNumber(17)).toNumber());
          _nw163[(_dafny.ZERO).toNumber()] = true;
          _nw163[(_dafny.ONE).toNumber()] = _1018_cf49;
          _nw163[(new BigNumber(2)).toNumber()] = true;
          _nw163[(new BigNumber(3)).toNumber()] = ((_1018_cf49) ? ((_this).f29) : ((_this).f30));
          _nw163[(new BigNumber(4)).toNumber()] = (_1022_cf45) === ((((_943_v86).contains((_this).f30)) ? ((_943_v86).get((_this).f30)) : (_1022_cf45)));
          _nw163[(new BigNumber(5)).toNumber()] = (_this).f30;
          _nw163[(new BigNumber(6)).toNumber()] = _1018_cf49;
          _nw163[(new BigNumber(7)).toNumber()] = (_this).f30;
          _nw163[(new BigNumber(8)).toNumber()] = _1022_cf45;
          _nw163[(new BigNumber(9)).toNumber()] = _1022_cf45;
          _nw163[(new BigNumber(10)).toNumber()] = (_this).f29;
          _nw163[(new BigNumber(11)).toNumber()] = (_dafny.MultiSet.fromElements(_1020_cf47, p1)).IsSubsetOf(_dafny.MultiSet.fromElements(_1020_cf47));
          _nw163[(new BigNumber(12)).toNumber()] = (((_this).f29) ? (!((_this).f29)) : (_1022_cf45));
          _nw163[(new BigNumber(13)).toNumber()] = (p1).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.update(_1021_cf46, _module.__default.safeIndex(new BigNumber((_1023_v134).cardinality()), new BigNumber((_1021_cf46).length)), _872_v35)).length));
          _nw163[(new BigNumber(14)).toNumber()] = !_dafny.areEqual(_dafny.Seq.UnicodeFromString("vu"), _877_v38);
          _nw163[(new BigNumber(15)).toNumber()] = _1018_cf49;
          _nw163[(new BigNumber(16)).toNumber()] = false;
          _1024_v135 = _nw163;
          let _index137 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_1024_v135).length));
          (_1024_v135)[_index137] = _1018_cf49;
          let _index138 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_1024_v135).length));
          (_1024_v135)[_index138] = ((_1022_cf45) ? ((_1020_cf47).isLessThanOrEqualTo(_1020_cf47)) : (_1018_cf49));
          let _1025_v136;
          let _nw164 = Array((new BigNumber(5)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1025_v136 = _nw164;
          let _index139 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_1025_v136).length));
          (_1025_v136)[_index139] = _872_v35;
          let _1026_v137;
          let _nw165 = Array((new BigNumber(2)).toNumber()).fill([]);
          _1026_v137 = _nw165;
          let _1027_v138;
          let _nw166 = new _module.C1();
          _nw166.__ctor(new BigNumber((_module.__default.fm28(p1, (_this).f29, _1018_cf49, globalState)).length), _1026_v137, _1022_cf45);
          _1027_v138 = _nw166;
          let _1028_v139;
          _1028_v139 = _1027_v138;
          let _1029_v140;
          let _nw167 = Array((new BigNumber(13)).toNumber());
          _nw167[(_dafny.ZERO).toNumber()] = _1027_v138;
          _nw167[(_dafny.ONE).toNumber()] = _1027_v138;
          _nw167[(new BigNumber(2)).toNumber()] = _1027_v138;
          _nw167[(new BigNumber(3)).toNumber()] = (_1028_v139);
          _nw167[(new BigNumber(4)).toNumber()] = _1027_v138;
          _nw167[(new BigNumber(5)).toNumber()] = _1027_v138;
          _nw167[(new BigNumber(6)).toNumber()] = _1027_v138;
          _nw167[(new BigNumber(7)).toNumber()] = _1027_v138;
          _nw167[(new BigNumber(8)).toNumber()] = _1027_v138;
          _nw167[(new BigNumber(9)).toNumber()] = _1027_v138;
          _nw167[(new BigNumber(10)).toNumber()] = _1027_v138;
          _nw167[(new BigNumber(11)).toNumber()] = (_1028_v139);
          _nw167[(new BigNumber(12)).toNumber()] = _1027_v138;
          _1029_v140 = _nw167;
          let _index140 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_1029_v140).length));
          (_1029_v140)[_index140] = _1027_v138;
          let _1030_v141;
          let _nw168 = Array((new BigNumber(13)).toNumber());
          _nw168[(_dafny.ZERO).toNumber()] = (_943_v86).update(_1022_cf45, _1018_cf49);
          _nw168[(_dafny.ONE).toNumber()] = _943_v86;
          _nw168[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1018_cf49,_1022_cf45);
          _nw168[(new BigNumber(3)).toNumber()] = _943_v86;
          _nw168[(new BigNumber(4)).toNumber()] = _943_v86;
          _nw168[(new BigNumber(5)).toNumber()] = _943_v86;
          _nw168[(new BigNumber(6)).toNumber()] = _943_v86;
          _nw168[(new BigNumber(7)).toNumber()] = _943_v86;
          _nw168[(new BigNumber(8)).toNumber()] = _943_v86;
          _nw168[(new BigNumber(9)).toNumber()] = _943_v86;
          _nw168[(new BigNumber(10)).toNumber()] = _943_v86;
          _nw168[(new BigNumber(11)).toNumber()] = _943_v86;
          _nw168[(new BigNumber(12)).toNumber()] = _943_v86;
          _1030_v141 = _nw168;
          let _1031_v142;
          _1031_v142 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_1020_cf47),_dafny.Map.Empty.slice().updateUnsafe(p1,_1018_cf49));
          let _1032_v143;
          let _nw169 = new _module.C3();
          _nw169.__ctor(_1030_v141, _1031_v142, _1022_cf45);
          _1032_v143 = _nw169;
          let _1033_v144;
          _1033_v144 = _dafny.Set.fromElements(_1032_v143, _1032_v143);
          let _1034_v145;
          _1034_v145 = _dafny.Map.Empty.slice().updateUnsafe(_943_v86,_1033_v144);
          let _1035_v146;
          let _nw170 = Array((new BigNumber(4)).toNumber());
          _nw170[(_dafny.ZERO).toNumber()] = _1033_v144;
          _nw170[(_dafny.ONE).toNumber()] = _1033_v144;
          _nw170[(new BigNumber(2)).toNumber()] = (((_1034_v145).contains(_943_v86)) ? ((_1034_v145).get(_943_v86)) : (_dafny.Set.fromElements(_1032_v143, _1032_v143, _1032_v143, _1032_v143, _1032_v143)));
          _nw170[(new BigNumber(3)).toNumber()] = _dafny.Set.fromElements(_1032_v143);
          _1035_v146 = _nw170;
          let _1036_v147;
          _1036_v147 = _dafny.Map.Empty.slice().updateUnsafe(_1035_v146,(_this).f29);
          let _1037_v148;
          _1037_v148 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm29(_1020_cf47, _1018_cf49, globalState),(p0)[_module.__default.safeIndex(_1020_cf47, new BigNumber((p0).length))]);
          let _index141 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_1024_v135).length));
          let _index142 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_1025_v136).length));
          let _index143 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_1029_v140).length));
          let _rhs148 = (((_1036_v147).contains(_1035_v146)) ? ((_1036_v147).get(_1035_v146)) : (!(new BigNumber((p0).length)).isEqualTo(p1)));
          let _rhs149 = (((_1037_v148).contains(p2)) ? ((_1037_v148).get(p2)) : (_1020_cf47));
          let _rhs150 = _1021_cf46;
          let _rhs151 = _872_v35;
          let _rhs152 = ((_1018_cf49) ? (_1027_v138) : ((((_this).f29) ? (_1027_v138) : (_1027_v138))));
          let _lhs90 = _1024_v135;
          let _lhs91 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_1024_v135).length));
          let _lhs92 = _1025_v136;
          let _lhs93 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_1025_v136).length));
          let _lhs94 = _1029_v140;
          let _lhs95 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_1029_v140).length));
          _lhs90[_lhs91] = _rhs148;
          _1020_cf47 = _rhs149;
          _877_v38 = _rhs150;
          _lhs92[_lhs93] = _rhs151;
          _lhs94[_lhs95] = _rhs152;
          let _1038_v149;
          _1038_v149 = _module.D3.create_DC10(_877_v38, _1020_cf47, _1018_cf49);
          let _1039_v150;
          let _init21 = ((_1040_cf47) => function (_1041_i9) {
            return (_1041_i9).plus(_1040_cf47);
          })(_1020_cf47);
          let _nw171 = Array((new BigNumber(25)).toNumber());
          for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw171.length); _i0_21++) {
            _nw171[_i0_21] = _init21(new BigNumber(_i0_21));
          }
          _1039_v150 = _nw171;
          let _rhs153 = _1038_v149;
          let _rhs154 = _1039_v150;
          let _rhs155 = p1;
          let _rhs156 = (_this).f30;
          let _lhs96 = globalState;
          _1038_v149 = _rhs153;
          _lhs96.f24 = _rhs154;
          _1020_cf47 = _rhs155;
          _1022_cf45 = _rhs156;
          let _1042_v151;
          let _nw172 = new _module.C2();
          _nw172.__ctor(_dafny.Seq.of(true), _1018_cf49);
          _1042_v151 = _nw172;
          let _1043_v152;
          _1043_v152 = _dafny.Map.Empty.slice().updateUnsafe(_1020_cf47,_1042_v151);
          let _rhs157 = (_this).f29;
          let _rhs158 = (((_1043_v152).contains((_1032_v143).fm9(_1020_cf47, p0, p1, (_this).f30, globalState))) ? ((_1043_v152).get((_1032_v143).fm9(_1020_cf47, p0, p1, (_this).f30, globalState))) : (_1042_v151));
          _1022_cf45 = _rhs157;
          _1042_v151 = _rhs158;
          let _1044_v153;
          _1044_v153 = _dafny.Set.fromElements(_1020_cf47);
          let _1045_v154;
          _1045_v154 = _dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Set.fromElements(new BigNumber((p0).length)));
          let _1046_v155;
          _1046_v155 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1044_v153).length),(((_1045_v154).contains(_1022_cf45)) ? ((_1045_v154).get(_1022_cf45)) : (_1044_v153)));
          let _1047_v156;
          _1047_v156 = _dafny.Map.Empty.slice().updateUnsafe(_877_v38,new BigNumber((_1023_v134).cardinality()));
          let _1048_v157;
          _1048_v157 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(new BigNumber(607), new BigNumber((_1047_v156).length)),_1046_v155);
          let _index144 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_1024_v135).length));
          let _rhs159 = !(!_dafny.Seq.contains((_1042_v151).f38, (_1018_cf49) || ((_1024_v135)[_module.__default.safeIndex(new BigNumber(338), new BigNumber((_1024_v135).length))])));
          let _rhs160 = (((_1048_v157).contains(new BigNumber((_dafny.Seq.Concat(_1021_cf46, _1021_cf46)).length))) ? ((_1048_v157).get(new BigNumber((_dafny.Seq.Concat(_1021_cf46, _1021_cf46)).length))) : (_1046_v155));
          let _rhs161 = (_1023_v134).update(!(_1022_cf45), _module.__default.abs(_module.__default.safeModuloInt((_dafny.ZERO).minus((_dafny.ZERO).minus(p1)), p1)));
          let _rhs162 = _1024_v135;
          let _lhs97 = _1024_v135;
          let _lhs98 = _module.__default.safeIndex(new BigNumber(338), new BigNumber((_1024_v135).length));
          let _lhs99 = globalState;
          _lhs97[_lhs98] = _rhs159;
          _1046_v155 = _rhs160;
          _1023_v134 = _rhs161;
          _lhs99.f20 = _rhs162;
        } else {
          let _1049_v158;
          let _init22 = ((_1050_cf45) => function (_1051_i10) {
            return _1050_cf45;
          })(_1022_cf45);
          let _nw173 = Array((new BigNumber(7)).toNumber());
          for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw173.length); _i0_22++) {
            _nw173[_i0_22] = _init22(new BigNumber(_i0_22));
          }
          _1049_v158 = _nw173;
          (globalState).f20 = _1049_v158;
          _module.__default.m0(_module.__default.fm1(_872_v35, globalState), globalState);
          let _1052_v159;
          _1052_v159 = _dafny.Seq.of((_this).f29, (_this).f29, _module.__default.fm26((_this).f29, (_this).f29, (_this).f30, p1, globalState));
          let _1053_v160;
          _1053_v160 = _dafny.Map.Empty.slice().updateUnsafe(_1052_v159,p1);
          _1053_v160 = _1053_v160;
          let _1054_v161;
          _1054_v161 = _dafny.Seq.of(_1049_v158, _1049_v158, _1049_v158);
          let _index145 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_1049_v158).length));
          (_1049_v158)[_index145] = (_this).f30;
          let _1055_v162;
          let _nw174 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Set.Empty);
          _1055_v162 = _nw174;
          let _1056_v163;
          _1056_v163 = _dafny.Set.fromElements(_872_v35);
          let _index146 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_1055_v162).length));
          (_1055_v162)[_index146] = (_1056_v163).Union(_1056_v163);
          let _1057_v164;
          _1057_v164 = _dafny.Seq.of(_dafny.Set.fromElements(new _dafny.CodePoint('k'.codePointAt(0)), _872_v35, _872_v35));
          let _index147 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_1049_v158).length));
          let _index148 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_1055_v162).length));
          let _rhs163 = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_1052_v159, (((_this).f29) ? (_1052_v159) : (_1052_v159))), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.Concat(_1052_v159, (((_this).f29) ? (_1052_v159) : (_1052_v159)))).length)), !((_this).f30) || (_1018_cf49))).length);
          let _rhs164 = (new BigNumber(-756)).plus(p1);
          let _rhs165 = _1054_v161;
          let _rhs166 = (_this).f29;
          let _rhs167 = (_1057_v164)[_module.__default.safeIndex(new BigNumber((p0).length), new BigNumber((_1057_v164).length))];
          let _lhs100 = _1049_v158;
          let _lhs101 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_1049_v158).length));
          let _lhs102 = _1055_v162;
          let _lhs103 = _module.__default.safeIndex(new BigNumber(386), new BigNumber((_1055_v162).length));
          _1020_cf47 = _rhs163;
          _1020_cf47 = _rhs164;
          _1054_v161 = _rhs165;
          _lhs100[_lhs101] = _rhs166;
          _lhs102[_lhs103] = _rhs167;
          let _index149 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_1049_v158).length));
          let _rhs168 = (((((_943_v86).contains((_this).f29)) ? ((_943_v86).get((_this).f29)) : ((_1049_v158)[_module.__default.safeIndex(new BigNumber(132), new BigNumber((_1049_v158).length))]))) ? (true) : ((_this).f29));
          let _rhs169 = (_1020_cf47).plus((_dafny.ZERO).minus((_1020_cf47).minus(_1020_cf47)));
          let _lhs104 = _1049_v158;
          let _lhs105 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_1049_v158).length));
          _lhs104[_lhs105] = _rhs168;
          _1020_cf47 = _rhs169;
        }
        let _1058_v165;
        let _nw175 = Array((new BigNumber(18)).toNumber()).fill(false);
        _1058_v165 = _nw175;
        let _index150 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_1058_v165).length));
        (_1058_v165)[_index150] = _1018_cf49;
        let _index151 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_1058_v165).length));
        (_1058_v165)[_index151] = (_this).f29;
        let _1059_v166;
        _1059_v166 = _dafny.Seq.of((_this).f30);
        let _index152 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_1058_v165).length));
        let _index153 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_1058_v165).length));
        let _rhs170 = !(_873_v36).contains((_1059_v166)[_module.__default.safeIndex(p1, new BigNumber((_1059_v166).length))]);
        let _rhs171 = (_1058_v165)[_module.__default.safeIndex(new BigNumber(590), new BigNumber((_1058_v165).length))];
        let _rhs172 = _1059_v166;
        let _rhs173 = _module.__default.fm3(globalState);
        let _lhs106 = _1058_v165;
        let _lhs107 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_1058_v165).length));
        let _lhs108 = _1058_v165;
        let _lhs109 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_1058_v165).length));
        let _lhs110 = globalState;
        _lhs106[_lhs107] = _rhs170;
        _lhs108[_lhs109] = _rhs171;
        _1059_v166 = _rhs172;
        _lhs110.f16 = _rhs173;
        let _1060_v167;
        _1060_v167 = _dafny.Map.Empty.slice().updateUnsafe(_1018_cf49,_872_v35);
        let _1061_v168;
        _1061_v168 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p0).length),_872_v35);
        let _1062_v169;
        let _nw176 = Array((new BigNumber(7)).toNumber());
        _nw176[(_dafny.ZERO).toNumber()] = (((_1060_v167).contains(_1022_cf45)) ? ((_1060_v167).get(_1022_cf45)) : (_872_v35));
        _nw176[(_dafny.ONE).toNumber()] = _872_v35;
        _nw176[(new BigNumber(2)).toNumber()] = new _dafny.CodePoint('k'.codePointAt(0));
        _nw176[(new BigNumber(3)).toNumber()] = _872_v35;
        _nw176[(new BigNumber(4)).toNumber()] = _872_v35;
        _nw176[(new BigNumber(5)).toNumber()] = (((_1061_v168).contains(_1020_cf47)) ? ((_1061_v168).get(_1020_cf47)) : (_872_v35));
        _nw176[(new BigNumber(6)).toNumber()] = _872_v35;
        _1062_v169 = _nw176;
        let _index154 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_1062_v169).length));
        (_1062_v169)[_index154] = _872_v35;
        let _index155 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_1062_v169).length));
        (_1062_v169)[_index155] = _872_v35;
      } else {
        let _1063___mcc_h16 = (_source15).cf39;
        let _1064_cf39 = _1063___mcc_h16;
        let _1065_v170;
        _1065_v170 = new BigNumber(18);
        let _1066_v171;
        _1066_v171 = _dafny.Set.fromElements(p1);
        _1065_v170 = new BigNumber((_1066_v171).length);
        _877_v38 = _877_v38;
        let _1067_v172;
        let _nw177 = Array((new BigNumber(6)).toNumber()).fill(false);
        _1067_v172 = _nw177;
        let _index156 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_1067_v172).length));
        (_1067_v172)[_index156] = !(_873_v36).contains((_this).f29);
        let _index157 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_1067_v172).length));
        (_1067_v172)[_index157] = (_this).f30;
        (globalState).f16 = (_1065_v170).isLessThanOrEqualTo(_module.__default.safeDivisionInt(new BigNumber(-612), _1065_v170));
      }
      r0 = (_this).f29;
      r1 = !(!(_module.__default.safeDivisionInt(p1, p1)).isEqualTo(p1));
      return [r0, r1];
    }
    get f29() {
      let _this = this;
      return _this._f29;
    };
    get f30() {
      let _this = this;
      return _this._f30;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
