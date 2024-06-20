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
      let _source0 = _module.D0.create_DC1(new BigNumber(801), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(672), new BigNumber(917), new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()))).cardinality()), new BigNumber(458));
      if (_source0.is_DC1) {
        let _0___mcc_h0 = (_source0).cf1;
        let _1___mcc_h1 = (_source0).cf2;
        let _2___mcc_h2 = (_source0).cf3;
        let _3_cf3 = _2___mcc_h2;
        let _4_cf2 = _1___mcc_h1;
        let _5_cf1 = _0___mcc_h0;
        return _dafny.areEqual(_dafny.Seq.UnicodeFromString("bielrkk"), _dafny.Seq.UnicodeFromString("w"));
      } else if (_source0.is_DC2) {
        let _6___mcc_h3 = (_source0).cf4;
        let _7___mcc_h4 = (_source0).cf5;
        let _8___mcc_h5 = (_source0).cf6;
        let _9_cf6 = _8___mcc_h5;
        let _10_cf5 = _7___mcc_h4;
        let _11_cf4 = _6___mcc_h3;
        return !(false);
      } else if (_source0.is_DC0) {
        let _12___mcc_h6 = (_source0).cf0;
        let _13_cf0 = _12___mcc_h6;
        return (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(263)), function (_14_i0) {
          return new _dafny.CodePoint('e'.codePointAt(0));
        })).length)).isLessThan(new BigNumber((_dafny.Seq.UnicodeFromString("phiihuc")).length));
      } else {
        let _15___mcc_h7 = (_source0).cf7;
        let _16_cf7 = _15___mcc_h7;
        return true;
      }
    };
    static fm1(p0, p1, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-879)), function (_17_i0) {
        return ((true) ? (new _dafny.CodePoint('p'.codePointAt(0))) : (new _dafny.CodePoint('n'.codePointAt(0))));
      });
    };
    static fm2(p0, p1, p2, globalState) {
      return (((true) ? (_module.D5.create_DC13(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(641),false)))) : (_module.D5.create_DC13(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(219),false)))))).dtor_cf20;
    };
    static fm3(p0, p1, globalState) {
      if (!(true) || (false)) {
        if (false) {
          return new _dafny.CodePoint('j'.codePointAt(0));
        } else {
          return new _dafny.CodePoint('g'.codePointAt(0));
        }
      } else {
        return new _dafny.CodePoint('y'.codePointAt(0));
      }
    };
    static fm4(p0, p1, p2, globalState) {
      return ((_dafny.Set.fromElements(true)).Union(_dafny.Set.fromElements(!(!(!(!(true))))))).Union(_dafny.Set.fromElements(false));
    };
    static fm5(p0, globalState) {
      return new BigNumber(521);
    };
    static fm6(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(!(!((_module.D1.create_DC5(!(false), !(true), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true)).length),new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(!(false)))).cardinality()))).length))).dtor_cf10)),true)).Merge(((false) ? (_dafny.Map.Empty.slice().updateUnsafe(false,false)) : (_dafny.Map.Empty.slice().updateUnsafe(true,!(true)))));
    };
    static fm7(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(523)).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber(185))).length)),new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.of(false, !(true), !(false)))).cardinality())))).length));
    };
    static fm8(p0, globalState) {
      let _source1 = _module.D0.create_DC0(false);
      if (_source1.is_DC1) {
        let _18___mcc_h0 = (_source1).cf1;
        let _19___mcc_h1 = (_source1).cf2;
        let _20___mcc_h2 = (_source1).cf3;
        let _21_cf3 = _20___mcc_h2;
        let _22_cf2 = _19___mcc_h1;
        let _23_cf1 = _18___mcc_h0;
        return function () {
          let _coll0 = new _dafny.Set();
          for (const _compr_0 of (_dafny.Seq.of(_module.D0.create_DC0(true), _module.D0.create_DC0(true), _module.D0.create_DC0(false))).Elements) {
            let _24_v0 = _compr_0;
            if (_dafny.Seq.contains(_dafny.Seq.of(_module.D0.create_DC0(true), _module.D0.create_DC0(true), _module.D0.create_DC0(false)), _24_v0)) {
              _coll0.add(_24_v0);
            }
          }
          return _coll0;
        }();
      } else if (_source1.is_DC2) {
        let _25___mcc_h3 = (_source1).cf4;
        let _26___mcc_h4 = (_source1).cf5;
        let _27___mcc_h5 = (_source1).cf6;
        let _28_cf6 = _27___mcc_h5;
        let _29_cf5 = _26___mcc_h4;
        let _30_cf4 = _25___mcc_h3;
        if (!(false)) {
          return function () {
            let _coll1 = new _dafny.Set();
            for (const _compr_1 of (_dafny.Set.fromElements(_module.D0.create_DC0(true))).Elements) {
              let _31_v1 = _compr_1;
              if ((_dafny.Set.fromElements(_module.D0.create_DC0(true))).contains(_31_v1)) {
                _coll1.add(_31_v1);
              }
            }
            return _coll1;
          }();
        } else {
          return function () {
            let _coll2 = new _dafny.Set();
            for (const _compr_2 of (_dafny.Seq.of(_module.D0.create_DC0(false), _module.D0.create_DC0(false))).Elements) {
              let _32_v2 = _compr_2;
              if (_dafny.Seq.contains(_dafny.Seq.of(_module.D0.create_DC0(false), _module.D0.create_DC0(false)), _32_v2)) {
                _coll2.add(_32_v2);
              }
            }
            return _coll2;
          }();
        }
      } else if (_source1.is_DC0) {
        let _33___mcc_h6 = (_source1).cf0;
        let _34_cf0 = _33___mcc_h6;
        return function () {
          let _coll3 = new _dafny.Set();
          for (const _compr_3 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(226)), ((_35_cf0) => function (_36_i0) {
            return _module.D0.create_DC0(_35_cf0);
          })(_34_cf0))).Elements) {
            let _37_v3 = _compr_3;
            if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(226)), ((_38_cf0) => function (_36_i0) {
              return _module.D0.create_DC0(_38_cf0);
            })(_34_cf0)), _37_v3)) {
              _coll3.add(_37_v3);
            }
          }
          return _coll3;
        }();
      } else {
        let _39___mcc_h7 = (_source1).cf7;
        let _40_cf7 = _39___mcc_h7;
        return _dafny.Set.fromElements(_module.D0.create_DC0(false), _module.D0.create_DC0(!(true)));
      }
    };
    static fm9(globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('x'.codePointAt(0)),new BigNumber(697))).Merge(function () {
        let _coll4 = new _dafny.Map();
        for (const _compr_4 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(732)), function (_41_i0) {
          return new _dafny.CodePoint('w'.codePointAt(0));
        })).Elements) {
          let _42_v0 = _compr_4;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(732)), function (_41_i0) {
            return new _dafny.CodePoint('w'.codePointAt(0));
          }), _42_v0)) {
            _coll4.push([_42_v0,new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(_module.D5.create_DC13(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(808),false))))).length))).cardinality()), new BigNumber(720))).cardinality())]);
          }
        }
        return _coll4;
      }());
    };
    static fm10(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('r'.codePointAt(0)),new BigNumber(498))), _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),new BigNumber(330))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(515)), function (_43_i0) {
        return function () {
          let _coll5 = new _dafny.Map();
          for (const _compr_5 of (_dafny.Set.fromElements(new _dafny.CodePoint('t'.codePointAt(0)))).Elements) {
            let _44_v0 = _compr_5;
            if ((_dafny.Set.fromElements(new _dafny.CodePoint('t'.codePointAt(0)))).contains(_44_v0)) {
              _coll5.push([_44_v0,new BigNumber(357)]);
            }
          }
          return _coll5;
        }();
      })));
    };
    static fm11(p0, globalState) {
      return _module.D2.create_DC8(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(328)), function (_45_i0) {
  return new _dafny.CodePoint('p'.codePointAt(0));
}), _dafny.Seq.Create(_module.__default.abs(new BigNumber(387)), function (_46_i1) {
  return new _dafny.CodePoint('j'.codePointAt(0));
})), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new _dafny.CodePoint('q'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0))), _dafny.Seq.of(new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0))))).length)));
    };
    static fm12(p0, p1, p2, p3, globalState) {
      return (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("yo")).length))).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(-126))));
    };
    static fm13(p0, globalState) {
      let _source2 = _module.D1.create_DC6();
      if (_source2.is_DC5) {
        let _47___mcc_h0 = (_source2).cf9;
        let _48___mcc_h1 = (_source2).cf10;
        let _49___mcc_h2 = (_source2).cf11;
        let _50_cf11 = _49___mcc_h2;
        let _51_cf10 = _48___mcc_h1;
        let _52_cf9 = _47___mcc_h0;
        if (_52_cf9) {
          return _module.D1.create_DC6();
        } else {
          return _module.D1.create_DC6();
        }
      } else if (_source2.is_DC6) {
        return _module.D1.create_DC6();
      } else {
        let _53___mcc_h3 = (_source2).cf8;
        let _54_cf8 = _53___mcc_h3;
        return _module.D1.create_DC6();
      }
    };
    static fm14(p0, p1, p2, p3, globalState) {
      return _module.D4.create_DC10(_dafny.Seq.of(_module.D0.create_DC3(_module.D0.create_DC2(_dafny.Seq.UnicodeFromString("givcgweu"), new BigNumber(-147), new BigNumber(324)))));
    };
    static Main(__noArgsParameter) {
      let _55_globalState;
      let _nw0 = new _module.GlobalState();
      _nw0.__ctor(new BigNumber(825), _dafny.ONE, true);
      _55_globalState = _nw0;
      let _56_v0;
      _56_v0 = true;
      let _57_v1;
      let _nw1 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
      _57_v1 = _nw1;
      let _58_v2;
      _58_v2 = new BigNumber(891);
      let _59_v3;
      _59_v3 = _dafny.Map.Empty.slice().updateUnsafe(((_56_v0) ? (_57_v1) : (_57_v1)),(_dafny.ZERO).minus(_58_v2));
      _59_v3 = _59_v3;
      _56_v0 = !(_module.__default.fm0(_58_v2, new BigNumber((_dafny.Seq.UnicodeFromString("if")).length), _55_globalState));
      let _60_v4;
      let _init0 = ((_61_v2, _62_v0) => function (_63_i0) {
        return (_dafny.Map.Empty.slice().updateUnsafe(_61_v2,_62_v0)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_61_v2,_62_v0));
      })(_58_v2, _56_v0);
      let _nw2 = Array((new BigNumber(23)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw2.length); _i0_0++) {
        _nw2[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _60_v4 = _nw2;
      let _64_v5;
      _64_v5 = _dafny.Map.Empty.slice().updateUnsafe(_58_v2,_56_v0);
      let _index0 = _module.__default.safeIndex(new BigNumber(111), new BigNumber((_60_v4).length));
      (_60_v4)[_index0] = _64_v5;
      let _index1 = _module.__default.safeIndex(new BigNumber(111), new BigNumber((_60_v4).length));
      (_60_v4)[_index1] = _64_v5;
      let _65_v6;
      _65_v6 = _dafny.MultiSet.fromElements(false);
      _56_v0 = !(((_58_v2).plus(new BigNumber((_65_v6).cardinality()))).isLessThan((_58_v2).minus(_58_v2)));
      let _66_v7;
      _66_v7 = _module.D0.create_DC0(_56_v0);
      let _67_v8;
      let _nw3 = new _module.C0();
      _nw3.__ctor(_66_v7, _dafny.Map.Empty.slice().updateUnsafe(_56_v0,_57_v1), _57_v1, _58_v2);
      _67_v8 = _nw3;
      let _68_v9;
      _68_v9 = new _dafny.CodePoint('q'.codePointAt(0));
      let _69_v10;
      let _nw4 = new _module.C0();
      _nw4.__ctor((_67_v8).f7, (_67_v8).f8, _57_v1, _58_v2);
      _69_v10 = _nw4;
      let _70_v11;
      _70_v11 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('w'.codePointAt(0)),_69_v10);
      let _71_v12;
      _71_v12 = _dafny.Map.Empty.slice().updateUnsafe(_64_v5,!(!(_70_v11).contains(_68_v9)));
      _71_v12 = _71_v12;
      let _72_v13;
      let _nw5 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Set.Empty);
      _72_v13 = _nw5;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_72_v13).length))) {
        let _73_i1 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_73_i1)) && ((_73_i1).isLessThan(new BigNumber((_72_v13).length))))) {
          (_72_v13)[(_73_i1)] = ((_dafny.Set.fromElements(false, _56_v0)).Intersect(_dafny.Set.fromElements(_56_v0, true, _56_v0))).Union(_dafny.Set.fromElements(_56_v0));
        }
      }
      let _rhs0 = new BigNumber(551);
      let _rhs1 = new BigNumber(-189);
      let _lhs0 = _55_globalState;
      _lhs0.f1 = _rhs0;
      _58_v2 = _rhs1;
      _56_v0 = _56_v0;
      let _index2 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length));
      (_57_v1)[_index2] = _58_v2;
      let _index3 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length));
      (_57_v1)[_index3] = (((_65_v6).contains((_58_v2).isLessThan(_58_v2))) ? ((_65_v6).get((_58_v2).isLessThan(_58_v2))) : ((_69_v10).f4));
      let _74_v14;
      let _nw6 = Array((new BigNumber(7)).toNumber());
      _74_v14 = _nw6;
      let _75_v15;
      _75_v15 = _dafny.Map.Empty.slice().updateUnsafe(_74_v14,_56_v0);
      let _76_v16;
      _76_v16 = _dafny.Map.Empty.slice().updateUnsafe((_57_v1)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length))],(_69_v10).f4);
      let _77_v17;
      _77_v17 = _dafny.Seq.of(_56_v0);
      let _78_v18;
      _78_v18 = _module.D0.create_DC1(new BigNumber((_77_v17).length), (_69_v10).f4, new BigNumber((_dafny.Seq.of(_58_v2)).length));
      let _79_v19;
      _79_v19 = _dafny.Map.Empty.slice().updateUnsafe(_69_v10,_78_v18);
      _75_v15 = (_75_v15).update(_74_v14, _module.__default.fm0(new BigNumber((_76_v16).length), new BigNumber((_79_v19).length), _55_globalState));
      let _80_v20;
      let _nw7 = Array((new BigNumber(25)).toNumber());
      _80_v20 = _nw7;
      let _81_v21;
      _81_v21 = _dafny.Map.Empty.slice().updateUnsafe(_69_v10,_80_v20);
      _81_v21 = (_81_v21).update(_69_v10, _80_v20);
      let _82_v22;
      _82_v22 = _module.D1.create_DC4(_57_v1);
      _82_v22 = _82_v22;
      let _83_v23;
      _83_v23 = _dafny.Seq.UnicodeFromString("gpuqdrl");
      let _84_v24;
      let _nw8 = Array((new BigNumber(2)).toNumber()).fill(false);
      _84_v24 = _nw8;
      let _85_v25;
      let _nw9 = new _module.C1();
      _nw9.__ctor((new _dafny.CodePoint('y'.codePointAt(0))), _84_v24, (_69_v10).f3, _58_v2);
      _85_v25 = _nw9;
      let _rhs2 = _dafny.Seq.Concat(_83_v23, _83_v23);
      let _rhs3 = _85_v25;
      _83_v23 = _rhs2;
      _85_v25 = _rhs3;
      let _86_v26;
      _86_v26 = _module.D0.create_DC3(_module.D0.create_DC2(_83_v23, _module.__default.fm5(false, _55_globalState), _58_v2));
      let _source3 = _86_v26;
      if (_source3.is_DC1) {
        let _87___mcc_h0 = (_source3).cf1;
        let _88___mcc_h1 = (_source3).cf2;
        let _89___mcc_h2 = (_source3).cf3;
        let _90_cf3 = _89___mcc_h2;
        let _91_cf2 = _88___mcc_h1;
        let _92_cf1 = _87___mcc_h0;
        let _93_v27;
        _93_v27 = _dafny.MultiSet.fromElements(_90_cf3, (_69_v10).f4);
        if ((_93_v27).IsSubsetOf(_dafny.MultiSet.fromElements((_69_v10).f4))) {
          let _94_v28;
          _94_v28 = _dafny.Map.Empty.slice().updateUnsafe(_90_cf3,_85_v25);
          let _95_v29;
          _95_v29 = _dafny.Map.Empty.slice().updateUnsafe((_94_v28).Merge(_dafny.Map.Empty.slice().updateUnsafe((_57_v1)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length))],_85_v25)),!(_56_v0));
          _95_v29 = _95_v29;
          let _96_v30;
          _96_v30 = _dafny.Seq.of(_86_v26);
          let _97_v31;
          _97_v31 = _module.D4.create_DC10(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-50)), ((_98_v10, _99_v0, _100_v1) => function (_101_i2) {
  return _module.D0.create_DC3(_module.D0.create_DC1((_98_v10).f4, (_98_v10).f4, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_99_v0,(_100_v1)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_100_v1).length))])).length)));
})(_69_v10, _56_v0, _57_v1)));
          let _102_v32;
          _102_v32 = _dafny.Map.Empty.slice().updateUnsafe(_56_v0,_92_cf1);
          let _103_v33;
          _103_v33 = _dafny.Seq.of(new BigNumber((_83_v23).length));
          let _104_v34;
          _104_v34 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_96_v30, _dafny.Seq.update((_97_v31).dtor_cf16, _module.__default.safeIndex((((_102_v32).contains(_56_v0)) ? ((_102_v32).get(_56_v0)) : (new BigNumber((_103_v33).length))), new BigNumber(((_97_v31).dtor_cf16).length)), _86_v26)), _96_v30);
          let _105_v35;
          _105_v35 = _module.D2.create_DC7(_77_v17);
          let _rhs4 = _56_v0;
          let _rhs5 = (_69_v10).f4;
          let _rhs6 = _104_v34;
          let _rhs7 = _dafny.MultiSet.FromArray((_105_v35).dtor_cf12);
          _56_v0 = _rhs4;
          _91_cf2 = _rhs5;
          _104_v34 = _rhs6;
          _65_v6 = _rhs7;
          _56_v0 = _56_v0;
          let _106_v36;
          let _107_v37;
          let _108_v38;
          let _out0;
          let _out1;
          let _out2;
          let _outcollector0 = (_85_v25).m0(_55_globalState);
          _out0 = _outcollector0[0];
          _out1 = _outcollector0[1];
          _out2 = _outcollector0[2];
          _106_v36 = _out0;
          _107_v37 = _out1;
          _108_v38 = _out2;
          _108_v38 = _106_v36;
        } else {
          let _109_v39;
          _109_v39 = _dafny.Seq.of(_69_v10, _69_v10);
          _69_v10 = ((_56_v0) ? ((_109_v39)[_module.__default.safeIndex(new BigNumber(-846), new BigNumber((_109_v39).length))]) : (_69_v10));
          let _110_v40;
          _110_v40 = _dafny.Seq.of(new BigNumber((_93_v27).cardinality()));
          let _111_v41;
          _111_v41 = _dafny.Seq.of(_93_v27, _dafny.MultiSet.fromElements(new BigNumber((_110_v40).length)));
          let _112_v42;
          _112_v42 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_91_cf2),(_111_v41)[_module.__default.safeIndex(new BigNumber(737), new BigNumber((_111_v41).length))]);
          _93_v27 = (((_112_v42).contains(new BigNumber(855))) ? ((_112_v42).get(new BigNumber(855))) : (_93_v27));
          let _113_v43;
          _113_v43 = _module.D2.create_DC8(_83_v23, _92_cf1);
          let _pat_let_tv0 = _83_v23;
          let _index4 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length));
          let _rhs8 = _module.__default.fm0((_69_v10).f4, _90_cf3, _55_globalState);
          let _rhs9 = function (_pat_let0_0) {
            return function (_114_dt__update__tmp_h0) {
              return function (_pat_let1_0) {
                return function (_115_dt__update_hcf13_h0) {
                  return _module.D2.create_DC8(_115_dt__update_hcf13_h0, (_114_dt__update__tmp_h0).dtor_cf14);
                }(_pat_let1_0);
              }(_pat_let_tv0);
            }(_pat_let0_0);
          }(_module.__default.fm11(_56_v0, _55_globalState));
          let _rhs10 = _module.__default.fm3(_68_v9, _92_cf1, _55_globalState);
          let _rhs11 = (_69_v10).f4;
          let _lhs1 = _57_v1;
          let _lhs2 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length));
          _56_v0 = _rhs8;
          _113_v43 = _rhs9;
          _68_v9 = _rhs10;
          _lhs1[_lhs2] = _rhs11;
          _84_v24 = (_85_v25).f6;
          let _index5 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length));
          (_57_v1)[_index5] = (_dafny.ZERO).minus((_69_v10).f4);
        }
        let _index6 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length));
        (_57_v1)[_index6] = ((_69_v10).f4).minus((_57_v1)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length))]);
        if (!_dafny.Seq.contains(_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.of(_56_v0), _module.__default.safeIndex(_58_v2, new BigNumber((_dafny.Seq.of(_56_v0)).length)), _56_v0), _module.__default.safeIndex((_69_v10).f4, new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_56_v0), _module.__default.safeIndex(_58_v2, new BigNumber((_dafny.Seq.of(_56_v0)).length)), _56_v0)).length)), _56_v0), _56_v0)) {
          let _index7 = _module.__default.safeIndex(new BigNumber(631), new BigNumber(((_85_v25).f6).length));
          ((_85_v25).f6)[_index7] = _56_v0;
          let _index8 = _module.__default.safeIndex(new BigNumber(631), new BigNumber(((_85_v25).f6).length));
          ((_85_v25).f6)[_index8] = _module.__default.fm0(_91_cf2, ((_69_v10).f4).minus(_91_cf2), _55_globalState);
          _76_v16 = (_76_v16).update(new BigNumber(((_60_v4)[_module.__default.safeIndex(new BigNumber(111), new BigNumber((_60_v4).length))]).length), _92_cf1);
          let _index9 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length));
          (_57_v1)[_index9] = new BigNumber(656);
          let _116_v44;
          let _out3;
          _out3 = (_85_v25).m1(false, _58_v2, _module.__default.fm5(!(_56_v0), _55_globalState), _55_globalState);
          _116_v44 = _out3;
          let _117_v45;
          _117_v45 = _dafny.MultiSet.fromElements((_69_v10).f3);
          let _index10 = _module.__default.safeIndex(new BigNumber(631), new BigNumber(((_85_v25).f6).length));
          let _index11 = _module.__default.safeIndex(new BigNumber(631), new BigNumber(((_85_v25).f6).length));
          let _index12 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length));
          let _rhs12 = !(_module.__default.fm0(new BigNumber((_83_v23).length), (_57_v1)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length))], _55_globalState));
          let _rhs13 = !(!((_117_v45).Difference(_117_v45)).contains((_69_v10).f3));
          let _rhs14 = new BigNumber((_116_v44).length);
          let _lhs3 = (_85_v25).f6;
          let _lhs4 = _module.__default.safeIndex(new BigNumber(631), new BigNumber(((_85_v25).f6).length));
          let _lhs5 = (_85_v25).f6;
          let _lhs6 = _module.__default.safeIndex(new BigNumber(631), new BigNumber(((_85_v25).f6).length));
          let _lhs7 = _57_v1;
          let _lhs8 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length));
          _lhs3[_lhs4] = _rhs12;
          _lhs5[_lhs6] = _rhs13;
          _lhs7[_lhs8] = _rhs14;
        } else {
          let _index13 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length));
          (_57_v1)[_index13] = _module.__default.safeDivisionInt(_90_cf3, (_57_v1)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length))]);
          _56_v0 = _56_v0;
          let _118_v46;
          _118_v46 = _dafny.Seq.of(_83_v23, _83_v23, _83_v23, _83_v23);
          let _119_v47;
          _119_v47 = _dafny.Map.Empty.slice().updateUnsafe(_56_v0,(_69_v10).f4);
          let _120_v48;
          let _nw10 = Array((new BigNumber(18)).toNumber());
          _nw10[(_dafny.ZERO).toNumber()] = ((_56_v0) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-279)), ((_121_v25) => function (_122_i3) {
            return (_121_v25).f5;
          })(_85_v25))) : (_dafny.Seq.UnicodeFromString("avwjo")));
          _nw10[(_dafny.ONE).toNumber()] = _dafny.Seq.UnicodeFromString("ntbyypb");
          _nw10[(new BigNumber(2)).toNumber()] = _83_v23;
          _nw10[(new BigNumber(3)).toNumber()] = (_118_v46)[_module.__default.safeIndex((((_119_v47).contains(_56_v0)) ? ((_119_v47).get(_56_v0)) : (_module.__default.fm5(_56_v0, _55_globalState))), new BigNumber((_118_v46).length))];
          _nw10[(new BigNumber(4)).toNumber()] = _module.__default.fm1(_56_v0, _56_v0, _55_globalState);
          _nw10[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("dddi");
          _nw10[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_83_v23, _83_v23);
          _nw10[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("gnfds"), _83_v23);
          _nw10[(new BigNumber(8)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-336)), ((_123_v25) => function (_124_i4) {
            return (_123_v25).f5;
          })(_85_v25));
          _nw10[(new BigNumber(9)).toNumber()] = _dafny.Seq.UnicodeFromString("c");
          _nw10[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_83_v23, _83_v23);
          _nw10[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_83_v23, _dafny.Seq.UnicodeFromString("ddtjgujl"));
          _nw10[(new BigNumber(12)).toNumber()] = _83_v23;
          _nw10[(new BigNumber(13)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(308)), ((_125_v25) => function (_126_i5) {
            return (_125_v25).f5;
          })(_85_v25));
          _nw10[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_83_v23, _module.__default.fm1(_56_v0, !(_56_v0), _55_globalState));
          _nw10[(new BigNumber(15)).toNumber()] = _83_v23;
          _nw10[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("my"), _83_v23);
          _nw10[(new BigNumber(17)).toNumber()] = _83_v23;
          _120_v48 = _nw10;
          let _index14 = _module.__default.safeIndex(new BigNumber(636), new BigNumber((_120_v48).length));
          (_120_v48)[_index14] = _dafny.Seq.Concat(_83_v23, _83_v23);
          let _index15 = _module.__default.safeIndex(new BigNumber(636), new BigNumber((_120_v48).length));
          (_120_v48)[_index15] = _dafny.Seq.update(_dafny.Seq.Concat(_83_v23, _dafny.Seq.Create(_module.__default.abs(new BigNumber(237)), function (_127_i6) {
            return new _dafny.CodePoint('a'.codePointAt(0));
          })), _module.__default.safeIndex(_91_cf2, new BigNumber((_dafny.Seq.Concat(_83_v23, _dafny.Seq.Create(_module.__default.abs(new BigNumber(237)), function (_128_i6) {
            return new _dafny.CodePoint('a'.codePointAt(0));
          }))).length)), (_85_v25).f5);
          _69_v10 = _69_v10;
          let _129_v49;
          _129_v49 = _dafny.Seq.of(_92_cf1);
          let _130_v50;
          let _out4;
          _out4 = (_85_v25).m1(_56_v0, (_dafny.ZERO).minus((_57_v1)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length))]), (_129_v49)[_module.__default.safeIndex(_91_cf2, new BigNumber((_129_v49).length))], _55_globalState);
          _130_v50 = _out4;
        }
        let _131_v51;
        _131_v51 = _dafny.Map.Empty.slice().updateUnsafe(_77_v17,!_dafny.areEqual(_83_v23, _dafny.Seq.of((_83_v23)[_module.__default.safeIndex(_90_cf3, new BigNumber((_83_v23).length))])));
        let _pat_let_tv1 = _56_v0;
        let _index16 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length));
        let _rhs15 = (((_131_v51).contains(_dafny.Seq.of(_56_v0))) ? ((_131_v51).get(_dafny.Seq.of(_56_v0))) : ((function (_pat_let2_0) {
          return function (_132_dt__update__tmp_h1) {
            return function (_pat_let3_0) {
              return function (_133_dt__update_hcf0_h0) {
                return _module.D0.create_DC0(_133_dt__update_hcf0_h0);
              }(_pat_let3_0);
            }(_pat_let_tv1);
          }(_pat_let2_0);
        }(_66_v7)).dtor_cf0));
        let _rhs16 = _module.__default.fm5(_56_v0, _55_globalState);
        let _lhs9 = _57_v1;
        let _lhs10 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length));
        _56_v0 = _rhs15;
        _lhs9[_lhs10] = _rhs16;
      } else if (_source3.is_DC2) {
        let _134___mcc_h3 = (_source3).cf4;
        let _135___mcc_h4 = (_source3).cf5;
        let _136___mcc_h5 = (_source3).cf6;
        let _137_cf6 = _136___mcc_h5;
        let _138_cf5 = _135___mcc_h4;
        let _139_cf4 = _134___mcc_h3;
        let _140_v52;
        _140_v52 = _dafny.Seq.of((_69_v10).f3);
        _57_v1 = ((((_67_v8).f8).contains(!(_56_v0))) ? (((_67_v8).f8).get(!(_56_v0))) : ((_140_v52)[_module.__default.safeIndex((_dafny.ZERO).minus((_69_v10).f4), new BigNumber((_140_v52).length))]));
        if (_56_v0) {
          let _141_v53;
          let _nw11 = new _module.C1();
          _nw11.__ctor((_85_v25).f5, (_85_v25).f6, (_69_v10).f3, new BigNumber((_77_v17).length));
          _141_v53 = _nw11;
          _57_v1 = (_69_v10).f3;
          _139_cf4 = _83_v23;
          let _142_v54;
          let _init1 = ((_143_v25) => function (_144_i7) {
            return _dafny.Seq.of((_143_v25).f5);
          })(_85_v25);
          let _nw12 = Array((new BigNumber(27)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw12.length); _i0_1++) {
            _nw12[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _142_v54 = _nw12;
          let _rhs17 = (_dafny.ZERO).minus(_138_cf5);
          let _rhs18 = _142_v54;
          let _rhs19 = (_69_v10).f4;
          let _rhs20 = _56_v0;
          let _rhs21 = ((_57_v1)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length))]).isLessThan((_69_v10).f4);
          let _lhs11 = _55_globalState;
          _lhs11.f1 = _rhs17;
          _142_v54 = _rhs18;
          _58_v2 = _rhs19;
          _56_v0 = _rhs20;
          _56_v0 = _rhs21;
          let _145_v55;
          _145_v55 = _module.D2.create_DC8(_83_v23, _138_cf5);
          let _146_v56;
          _146_v56 = _dafny.MultiSet.fromElements((_145_v55).dtor_cf14);
          let _147_v57;
          _147_v57 = _dafny.Map.Empty.slice().updateUnsafe(_141_v53,_146_v56);
          _56_v0 = !(_56_v0) || ((_dafny.Set.fromElements(new BigNumber((_147_v57).length))).IsDisjointFrom(_dafny.Set.fromElements((_57_v1)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length))], _138_cf5)));
        } else {
          _69_v10 = _69_v10;
          _77_v17 = _dafny.Seq.Concat(_77_v17, _77_v17);
          let _148_v58;
          let _149_v59;
          let _150_v60;
          let _out5;
          let _out6;
          let _out7;
          let _outcollector1 = (_85_v25).m0(_55_globalState);
          _out5 = _outcollector1[0];
          _out6 = _outcollector1[1];
          _out7 = _outcollector1[2];
          _148_v58 = _out5;
          _149_v59 = _out6;
          _150_v60 = _out7;
          _150_v60 = ((_dafny.ZERO).minus(_58_v2)).isEqualTo(((_module.__default.fm0((_69_v10).f4, (_57_v1)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length))], _55_globalState)) ? (_137_cf6) : ((_69_v10).f4)));
          let _151_v61;
          _151_v61 = _dafny.Seq.of((_69_v10).f4, (_69_v10).f4, _138_cf5);
          let _152_v62;
          _152_v62 = _dafny.Map.Empty.slice().updateUnsafe(_151_v61,_137_cf6);
          let _153_v64;
          let _nw13 = Array((new BigNumber(16)).toNumber());
          _nw13[(_dafny.ZERO).toNumber()] = _module.__default.fm3(_68_v9, new BigNumber((function () {
            let _coll6 = new _dafny.Set();
            for (const _compr_6 of (_152_v62).Keys.Elements) {
              let _154_v63 = _compr_6;
              if ((_152_v62).contains(_154_v63)) {
                _coll6.add(_154_v63);
              }
            }
            return _coll6;
          }()).length), _55_globalState);
          _nw13[(_dafny.ONE).toNumber()] = _68_v9;
          _nw13[(new BigNumber(2)).toNumber()] = (_85_v25).f5;
          _nw13[(new BigNumber(3)).toNumber()] = (_85_v25).f5;
          _nw13[(new BigNumber(4)).toNumber()] = (_85_v25).f5;
          _nw13[(new BigNumber(5)).toNumber()] = _68_v9;
          _nw13[(new BigNumber(6)).toNumber()] = (_85_v25).f5;
          _nw13[(new BigNumber(7)).toNumber()] = new _dafny.CodePoint('q'.codePointAt(0));
          _nw13[(new BigNumber(8)).toNumber()] = _68_v9;
          _nw13[(new BigNumber(9)).toNumber()] = (_85_v25).f5;
          _nw13[(new BigNumber(10)).toNumber()] = ((_149_v59) ? ((_85_v25).f5) : (new _dafny.CodePoint('o'.codePointAt(0))));
          _nw13[(new BigNumber(11)).toNumber()] = (_85_v25).f5;
          _nw13[(new BigNumber(12)).toNumber()] = _68_v9;
          _nw13[(new BigNumber(13)).toNumber()] = (_85_v25).f5;
          _nw13[(new BigNumber(14)).toNumber()] = (_85_v25).f5;
          _nw13[(new BigNumber(15)).toNumber()] = _module.__default.fm3(new _dafny.CodePoint('o'.codePointAt(0)), _58_v2, _55_globalState);
          _153_v64 = _nw13;
          let _index17 = _module.__default.safeIndex(new BigNumber(179), new BigNumber((_153_v64).length));
          (_153_v64)[_index17] = new _dafny.CodePoint('t'.codePointAt(0));
          let _index18 = _module.__default.safeIndex(new BigNumber(179), new BigNumber((_153_v64).length));
          (_153_v64)[_index18] = _68_v9;
        }
        if (false) {
          let _155_v65;
          _155_v65 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0((_57_v1)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length))], _138_cf5, _55_globalState),new BigNumber((_dafny.Seq.UnicodeFromString("bjmglnlgb")).length));
          _137_cf6 = ((((_56_v0) ? (_56_v0) : (true))) ? ((_dafny.ZERO).minus(new BigNumber((_140_v52).length))) : (((((_155_v65).contains(_56_v0)) ? ((_155_v65).get(_56_v0)) : (_137_cf6))).plus((_69_v10).f4)));
          let _156_v66;
          _156_v66 = _dafny.Seq.of((_69_v10).f4, (_69_v10).f4);
          let _157_v67;
          _157_v67 = _dafny.MultiSet.fromElements(new BigNumber((_156_v66).length));
          _56_v0 = !((_157_v67).Difference(_157_v67)).equals((_157_v67).Difference(_157_v67));
          let _158_v68;
          _158_v68 = (_85_v25).f5;
          let _159_v69;
          _159_v69 = _dafny.Map.Empty.slice().updateUnsafe(_158_v68,true);
          _159_v69 = (_159_v69).update(_68_v9, _56_v0);
          _56_v0 = (_module.__default.fm0(_138_cf5, (_57_v1)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length))], _55_globalState)) === (_56_v0);
          let _160_v70;
          _160_v70 = _module.D1.create_DC6();
          let _161_v71;
          _161_v71 = _dafny.Map.Empty.slice().updateUnsafe((_83_v23)[_module.__default.safeIndex((_69_v10).f4, new BigNumber((_83_v23).length))],_67_v8);
          let _162_v72;
          _162_v72 = _dafny.Seq.of(_67_v8, _67_v8, (((_161_v71).contains((_85_v25).f5)) ? ((_161_v71).get((_85_v25).f5)) : (_67_v8)), _67_v8, _67_v8);
          let _rhs22 = _160_v70;
          let _rhs23 = _162_v72;
          let _rhs24 = _56_v0;
          _160_v70 = _rhs22;
          _162_v72 = _rhs23;
          _56_v0 = _rhs24;
        } else {
          let _163_v73;
          let _nw14 = new _module.C0();
          _nw14.__ctor(_66_v7, _dafny.Map.Empty.slice().updateUnsafe(_56_v0,_57_v1), _57_v1, (_69_v10).f4);
          _163_v73 = _nw14;
          let _164_v74;
          let _nw15 = new _module.C0();
          _nw15.__ctor((_67_v8).f7, (_67_v8).f8, (_69_v10).f3, new BigNumber(931));
          _164_v74 = _nw15;
          let _165_v75;
          _165_v75 = _dafny.Map.Empty.slice().updateUnsafe(_65_v6,_56_v0);
          let _166_v76;
          _166_v76 = _dafny.Map.Empty.slice().updateUnsafe(_56_v0,_56_v0);
          _165_v75 = (_165_v75).update(_65_v6, !(new BigNumber((_166_v76).length)).isEqualTo((_57_v1)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length))]));
          _56_v0 = !(_module.__default.fm0((_69_v10).f4, _137_cf6, _55_globalState));
          let _167_v77;
          let _nw16 = new _module.C1();
          _nw16.__ctor(_68_v9, (_85_v25).f6, (_140_v52)[_module.__default.safeIndex(_58_v2, new BigNumber((_140_v52).length))], (_69_v10).f4);
          _167_v77 = _nw16;
        }
        _83_v23 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("mo"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-834)), function (_168_i8) {
          return new _dafny.CodePoint('d'.codePointAt(0));
        })), _module.__default.safeIndex((_69_v10).f4, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("mo"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-834)), function (_169_i8) {
          return new _dafny.CodePoint('d'.codePointAt(0));
        }))).length)), (_85_v25).f5);
      } else if (_source3.is_DC0) {
        let _170___mcc_h6 = (_source3).cf0;
        let _171_cf0 = _170___mcc_h6;
        let _172_v78;
        let _nw17 = new _module.C0();
        _nw17.__ctor((_67_v8).f7, ((_67_v8).f8).Merge(_dafny.Map.Empty.slice().updateUnsafe(_171_cf0,(_69_v10).f3)), (_69_v10).f3, new BigNumber((_83_v23).length));
        _172_v78 = _nw17;
        let _173_v79;
        _173_v79 = _dafny.MultiSet.fromElements(new BigNumber(-708), new BigNumber(399), _58_v2, (_57_v1)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length))], new BigNumber(153));
        let _174_v80;
        _174_v80 = _dafny.Seq.of((_69_v10).f4);
        if (((_dafny.MultiSet.FromArray(_174_v80)).Intersect(_173_v79)).IsSubsetOf(_173_v79)) {
          let _175_v81;
          let _out8;
          _out8 = (_85_v25).m1(_56_v0, (_69_v10).f4, (_dafny.ZERO).minus((_69_v10).f4), _55_globalState);
          _175_v81 = _out8;
          (_55_globalState).f1 = new BigNumber(90);
          let _176_v82;
          _176_v82 = _dafny.Seq.of(_86_v26);
          let _177_v83;
          _177_v83 = _module.D4.create_DC10(_176_v82);
          let _178_v84;
          _178_v84 = _dafny.Seq.of(_177_v83, _module.D4.create_DC10(_176_v82), _177_v83);
          let _179_v85;
          _179_v85 = _dafny.Map.Empty.slice().updateUnsafe(_173_v79,_dafny.Seq.Concat(_178_v84, _178_v84));
          _179_v85 = (_179_v85).update(_module.__default.fm12(new BigNumber((_module.__default.fm12(new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("ijxylmpsk"), _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("dk")).length), new BigNumber((_dafny.Seq.UnicodeFromString("ijxylmpsk")).length)), (_85_v25).f5)).length), (_69_v10).f4, (_69_v10).f4, _dafny.Seq.update(_175_v81, _module.__default.safeIndex((_69_v10).f4, new BigNumber((_175_v81).length)), (_85_v25).f5), _55_globalState)).cardinality()), _58_v2, (_57_v1)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length))], _175_v81, _55_globalState), _dafny.Seq.of(_177_v83));
          let _180_v86;
          _180_v86 = _module.D0.create_DC2(_83_v23, _58_v2, new BigNumber(785));
          let _rhs25 = _58_v2;
          let _rhs26 = _180_v86;
          let _lhs12 = _55_globalState;
          _lhs12.f1 = _rhs25;
          _180_v86 = _rhs26;
          _171_cf0 = _56_v0;
        } else {
          (_55_globalState).f1 = _module.__default.safeModuloInt((_module.__default.fm5(_56_v0, _55_globalState)).plus((_69_v10).f4), _module.__default.safeDivisionInt((_69_v10).f4, (_57_v1)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length))]));
          let _181_v87;
          _181_v87 = _dafny.Seq.of(_83_v23, _83_v23, _dafny.Seq.update(_83_v23, _module.__default.safeIndex((_69_v10).f4, new BigNumber((_83_v23).length)), _68_v9));
          _83_v23 = (_181_v87)[_module.__default.safeIndex((_69_v10).f4, new BigNumber((_181_v87).length))];
          _57_v1 = (_69_v10).f3;
          let _182_v88;
          let _183_v89;
          let _184_v90;
          let _out9;
          let _out10;
          let _out11;
          let _outcollector2 = (_85_v25).m0(_55_globalState);
          _out9 = _outcollector2[0];
          _out10 = _outcollector2[1];
          _out11 = _outcollector2[2];
          _182_v88 = _out9;
          _183_v89 = _out10;
          _184_v90 = _out11;
          let _185_v91;
          _185_v91 = _module.D1.create_DC6();
          let _186_v92;
          let _nw18 = Array((new BigNumber(18)).toNumber());
          _nw18[(_dafny.ZERO).toNumber()] = _185_v91;
          _nw18[(_dafny.ONE).toNumber()] = _185_v91;
          _nw18[(new BigNumber(2)).toNumber()] = _185_v91;
          _nw18[(new BigNumber(3)).toNumber()] = _module.__default.fm13((_dafny.ZERO).minus((_69_v10).f4), _55_globalState);
          _nw18[(new BigNumber(4)).toNumber()] = _185_v91;
          _nw18[(new BigNumber(5)).toNumber()] = _185_v91;
          _nw18[(new BigNumber(6)).toNumber()] = ((_183_v89) ? (_185_v91) : (_185_v91));
          _nw18[(new BigNumber(7)).toNumber()] = _185_v91;
          _nw18[(new BigNumber(8)).toNumber()] = _185_v91;
          _nw18[(new BigNumber(9)).toNumber()] = _module.D1.create_DC6();
          _nw18[(new BigNumber(10)).toNumber()] = _185_v91;
          _nw18[(new BigNumber(11)).toNumber()] = _185_v91;
          _nw18[(new BigNumber(12)).toNumber()] = _185_v91;
          _nw18[(new BigNumber(13)).toNumber()] = _module.__default.fm13(new BigNumber((_76_v16).length), _55_globalState);
          _nw18[(new BigNumber(14)).toNumber()] = _185_v91;
          _nw18[(new BigNumber(15)).toNumber()] = _185_v91;
          _nw18[(new BigNumber(16)).toNumber()] = _185_v91;
          _nw18[(new BigNumber(17)).toNumber()] = _185_v91;
          _186_v92 = _nw18;
          let _index19 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_186_v92).length));
          (_186_v92)[_index19] = _185_v91;
          let _index20 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_186_v92).length));
          (_186_v92)[_index20] = _185_v91;
        }
        let _187_v93;
        let _nw19 = new _module.C0();
        _nw19.__ctor((_67_v8).f7, ((_172_v78).f8).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0((_57_v1)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length))], _58_v2, _55_globalState),(_69_v10).f3)), _57_v1, (new BigNumber((_83_v23).length)).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-931)), ((_188_v25) => function (_189_i9) {
          return (_188_v25).f5;
        })(_85_v25))).length))));
        _187_v93 = _nw19;
        let _190_v94;
        _190_v94 = _dafny.Map.Empty.slice().updateUnsafe((_85_v25).f5,_171_cf0);
        let _source4 = _module.__default.fm14(((_56_v0) ? (_190_v94) : (_190_v94)), ((_69_v10).f4).multipliedBy((_69_v10).f4), _module.__default.safeModuloInt(new BigNumber(197), (_dafny.ZERO).minus((_69_v10).f4)), _78_v18, _55_globalState);
        if (_source4.is_DC11) {
          let _191___mcc_h8 = (_source4).cf17;
          let _192_cf17 = _191___mcc_h8;
          let _193_v95;
          let _init2 = ((_194_v25) => function (_195_i10) {
            return _dafny.Seq.Create(_module.__default.abs(new BigNumber(813)), ((_196_v25) => function (_197_i11) {
              return (_196_v25).f5;
            })(_194_v25));
          })(_85_v25);
          let _nw20 = Array((new BigNumber(9)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw20.length); _i0_2++) {
            _nw20[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _193_v95 = _nw20;
          let _index21 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_193_v95).length));
          (_193_v95)[_index21] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(38)), ((_198_v9) => function (_199_i12) {
            return (_198_v9);
          })(_68_v9)), _83_v23);
          let _index22 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_193_v95).length));
          let _index23 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length));
          let _rhs27 = new BigNumber(849);
          let _rhs28 = _83_v23;
          let _rhs29 = ((_dafny.ZERO).minus((_69_v10).f4)).multipliedBy(_module.__default.safeDivisionInt((_69_v10).f4, (_69_v10).f4));
          let _rhs30 = ((_dafny.ZERO).minus(((_57_v1)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length))]).plus((_57_v1)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length))]))).minus((_69_v10).f4);
          let _lhs13 = _55_globalState;
          let _lhs14 = _193_v95;
          let _lhs15 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_193_v95).length));
          let _lhs16 = _55_globalState;
          let _lhs17 = _57_v1;
          let _lhs18 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length));
          _lhs13.f1 = _rhs27;
          _lhs14[_lhs15] = _rhs28;
          _lhs16.f1 = _rhs29;
          _lhs17[_lhs18] = _rhs30;
          let _200_v96;
          let _nw21 = Array((new BigNumber(6)).toNumber()).fill([]);
          _200_v96 = _nw21;
          let _index24 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_200_v96).length));
          (_200_v96)[_index24] = (_85_v25).f6;
          let _index25 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_200_v96).length));
          let _nw22 = Array((new BigNumber(9)).toNumber());
          _nw22[(_dafny.ZERO).toNumber()] = false;
          _nw22[(_dafny.ONE).toNumber()] = _171_cf0;
          _nw22[(new BigNumber(2)).toNumber()] = _56_v0;
          _nw22[(new BigNumber(3)).toNumber()] = _56_v0;
          _nw22[(new BigNumber(4)).toNumber()] = false;
          _nw22[(new BigNumber(5)).toNumber()] = _56_v0;
          _nw22[(new BigNumber(6)).toNumber()] = _171_cf0;
          _nw22[(new BigNumber(7)).toNumber()] = _171_cf0;
          _nw22[(new BigNumber(8)).toNumber()] = _56_v0;
          (_200_v96)[_index25] = _nw22;
          let _pat_let_tv2 = _171_cf0;
          let _201_v97;
          let _nw23 = new _module.C0();
          _nw23.__ctor(function (_pat_let4_0) {
            return function (_202_dt__update__tmp_h2) {
              return function (_pat_let5_0) {
                return function (_203_dt__update_hcf0_h1) {
                  return _module.D0.create_DC0(_203_dt__update_hcf0_h1);
                }(_pat_let5_0);
              }(_pat_let_tv2);
            }(_pat_let4_0);
          }((_187_v93).f7), _dafny.Map.Empty.slice().updateUnsafe(_56_v0,(_69_v10).f3), (_69_v10).f3, new BigNumber((_174_v80).length));
          _201_v97 = _nw23;
          let _index26 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length));
          (_57_v1)[_index26] = (_69_v10).f4;
        } else if (_source4.is_DC12) {
          let _204___mcc_h9 = (_source4).cf18;
          let _205___mcc_h10 = (_source4).cf19;
          let _206_cf19 = _205___mcc_h10;
          let _207_cf18 = _204___mcc_h9;
          let _208_v98;
          _208_v98 = _dafny.Set.fromElements(_171_cf0);
          (_55_globalState).f1 = _module.__default.safeDivisionInt(_module.__default.fm5(!(_module.__default.fm0((_57_v1)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length))], (_69_v10).f4, _55_globalState)), _55_globalState), new BigNumber((_208_v98).length));
          let _index27 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length));
          (_57_v1)[_index27] = ((_69_v10).f4).multipliedBy(_58_v2);
          let _209_v99;
          _209_v99 = _dafny.Set.fromElements(_83_v23, _dafny.Seq.Concat(_83_v23, _dafny.Seq.Create(_module.__default.abs(new BigNumber(696)), ((_210_v23, _211_v2) => function (_212_i13) {
            return (_210_v23)[_module.__default.safeIndex(_211_v2, new BigNumber((_210_v23).length))];
          })(_83_v23, _58_v2))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(974)), ((_213_v25) => function (_214_i14) {
            return (_213_v25).f5;
          })(_85_v25)), _83_v23, _83_v23);
          _209_v99 = (_dafny.Set.fromElements(_83_v23)).Intersect(_209_v99);
          _173_v79 = _173_v79;
        } else {
          let _215___mcc_h11 = (_source4).cf16;
          let _216_cf16 = _215___mcc_h11;
          let _217_v100;
          _217_v100 = _module.D0.create_DC2(_83_v23, (_69_v10).f4, (_69_v10).f4);
          (_55_globalState).f1 = (_217_v100).dtor_cf5;
          _187_v93 = _67_v8;
          let _index28 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length));
          (_57_v1)[_index28] = _58_v2;
          let _218_v101;
          let _nw24 = new _module.C1();
          _nw24.__ctor((_85_v25).f5, _84_v24, (_69_v10).f3, _module.__default.safeDivisionInt((_69_v10).f4, _58_v2));
          _218_v101 = _nw24;
        }
      } else {
        let _219___mcc_h7 = (_source3).cf7;
        let _220_cf7 = _219___mcc_h7;
        (_55_globalState).f1 = (_69_v10).f4;
        let _221_v102;
        let _222_v103;
        let _223_v104;
        let _out12;
        let _out13;
        let _out14;
        let _outcollector3 = (_85_v25).m0(_55_globalState);
        _out12 = _outcollector3[0];
        _out13 = _outcollector3[1];
        _out14 = _outcollector3[2];
        _221_v102 = _out12;
        _222_v103 = _out13;
        _223_v104 = _out14;
        let _224_v105;
        _224_v105 = _dafny.Seq.of(new BigNumber((_83_v23).length), (_69_v10).f4);
        _56_v0 = _module.__default.fm0((new BigNumber(918)).minus(_module.__default.fm5(_module.__default.fm0(new BigNumber((_224_v105).length), _58_v2, _55_globalState), _55_globalState)), (_58_v2).multipliedBy(new BigNumber(-630)), _55_globalState);
        let _225_v106;
        _225_v106 = _dafny.Seq.of(_69_v10, _69_v10);
        let _rhs31 = !(_module.__default.fm0((_57_v1)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length))], (_69_v10).f4, _55_globalState));
        let _rhs32 = (_dafny.ZERO).minus(_58_v2);
        let _rhs33 = _83_v23;
        let _rhs34 = (_225_v106)[_module.__default.safeIndex(_58_v2, new BigNumber((_225_v106).length))];
        let _rhs35 = (_module.D1.create_DC5(_223_v104, _223_v104, (_69_v10).f4)).dtor_cf10;
        let _lhs19 = _55_globalState;
        _56_v0 = _rhs31;
        _lhs19.f1 = _rhs32;
        _83_v23 = _rhs33;
        _69_v10 = _rhs34;
        _222_v103 = _rhs35;
      }
      let _226_v107;
      _226_v107 = _dafny.Map.Empty.slice().updateUnsafe(_69_v10,_56_v0);
      let _227_v108;
      _227_v108 = _dafny.Seq.of(new BigNumber((_226_v107).length));
      let _index29 = _module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length));
      (_57_v1)[_index29] = _module.__default.safeDivisionInt((_57_v1)[_module.__default.safeIndex(new BigNumber(634), new BigNumber((_57_v1).length))], _module.__default.safeModuloInt(_58_v2, (_dafny.ZERO).minus((_227_v108)[_module.__default.safeIndex(_58_v2, new BigNumber((_227_v108).length))])));
      process.stdout.write(_dafny.toString((_55_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_55_globalState.f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_55_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_56_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_57_v1)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_57_v1)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_57_v1)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_57_v1)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_58_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_59_v3).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v4)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(891),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v4)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(891),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v4)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(891),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v4)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(891),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v4)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(891),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v4)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(891),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v4)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(891),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v4)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(891),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v4)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(891),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v4)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(891),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v4)[new BigNumber(10)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(891),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v4)[new BigNumber(11)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(891),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v4)[new BigNumber(12)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(891),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v4)[new BigNumber(13)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(891),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v4)[new BigNumber(14)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(891),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v4)[new BigNumber(15)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(891),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v4)[new BigNumber(16)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(891),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v4)[new BigNumber(17)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(891),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v4)[new BigNumber(18)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(891),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v4)[new BigNumber(19)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(891),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v4)[new BigNumber(20)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(891),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v4)[new BigNumber(21)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(891),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_60_v4)[new BigNumber(22)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(891),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_64_v5).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(891),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_65_v6).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_66_v7).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_67_v8).f7).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_67_v8).f8).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_68_v9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_69_v10).f3)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_69_v10).f3)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_69_v10).f3)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_69_v10).f3)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v10).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_70_v11).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v12).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(891),true),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_72_v13)[_dafny.ZERO]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_72_v13)[_dafny.ONE]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_72_v13)[new BigNumber(2)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_72_v13)[new BigNumber(3)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_72_v13)[new BigNumber(4)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_72_v13)[new BigNumber(5)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_72_v13)[new BigNumber(6)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_72_v13)[new BigNumber(7)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_72_v13)[new BigNumber(8)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_72_v13)[new BigNumber(9)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_72_v13)[new BigNumber(10)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_72_v13)[new BigNumber(11)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_72_v13)[new BigNumber(12)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_72_v13)[new BigNumber(13)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_72_v13)[new BigNumber(14)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_72_v13)[new BigNumber(15)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_75_v15).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_76_v16).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(891)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_77_v17, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_78_v18).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_78_v18).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_78_v18).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_79_v19).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_81_v21).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_82_v22).dtor_cf8)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_82_v22).dtor_cf8)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_82_v22).dtor_cf8)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_82_v22).dtor_cf8)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_83_v23).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_84_v24)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_84_v24)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v25).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_v25).f6)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_85_v25).f6)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_86_v26).dtor_cf7).dtor_cf4).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_86_v26).dtor_cf7).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_86_v26).dtor_cf7).dtor_cf6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_226_v107).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_227_v108, _dafny.Seq.of(_dafny.ONE))));
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
    static create_DC2(cf4, cf5, cf6) {
      let $dt = new D0(1);
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
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + this.cf4.toVerbatimString(true) + ", " + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ")";
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
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1) && _dafny.areEqual(this.cf2, other.cf2) && _dafny.areEqual(this.cf3, other.cf3);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf4, other.cf4) && _dafny.areEqual(this.cf5, other.cf5) && _dafny.areEqual(this.cf6, other.cf6);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf0 === other.cf0;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf7, other.cf7);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(_dafny.ZERO, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC5(cf9, cf10, cf11) {
      let $dt = new D1(0);
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC6() {
      let $dt = new D1(1);
      return $dt;
    }
    static create_DC4(cf8) {
      let $dt = new D1(2);
      $dt.cf8 = cf8;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get is_DC4() { return this.$tag === 2; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf8() { return this.cf8; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC6";
      } else if (this.$tag === 2) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf8) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf9 === other.cf9 && this.cf10 === other.cf10 && _dafny.areEqual(this.cf11, other.cf11);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf8 === other.cf8;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC5(false, false, _dafny.ZERO);
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
    static create_DC8(cf13, cf14) {
      let $dt = new D2(0);
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC7(cf12) {
      let $dt = new D2(1);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC8" + "(" + this.cf13.toVerbatimString(true) + ", " + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf13, other.cf13) && _dafny.areEqual(this.cf14, other.cf14);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf12, other.cf12);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC8(_dafny.Seq.UnicodeFromString(""), _dafny.ZERO);
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
    static create_DC9(cf15) {
      let $dt = new D3(0);
      $dt.cf15 = cf15;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get dtor_cf15() { return this.cf15; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf15) + ")";
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
      return new _dafny.CodePoint('D'.codePointAt(0));
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
    static create_DC11(cf17) {
      let $dt = new D4(0);
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC12(cf18, cf19) {
      let $dt = new D4(1);
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      return $dt;
    }
    static create_DC10(cf16) {
      let $dt = new D4(2);
      $dt.cf16 = cf16;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get is_DC10() { return this.$tag === 2; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf16() { return this.cf16; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf16) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf18 === other.cf18 && this.cf19 === other.cf19;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf16, other.cf16);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC11(new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC14(cf21, cf22, cf23, cf24) {
      let $dt = new D5(0);
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC13(cf20) {
      let $dt = new D5(1);
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC15(cf25) {
      let $dt = new D5(2);
      $dt.cf25 = cf25;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get is_DC15() { return this.$tag === 2; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf25() { return this.cf25; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf25) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf21, other.cf21) && _dafny.areEqual(this.cf22, other.cf22) && this.cf23 === other.cf23 && _dafny.areEqual(this.cf24, other.cf24);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf20, other.cf20);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf25, other.cf25);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC14(_dafny.ZERO, _dafny.ZERO, false, _dafny.ZERO);
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
      this.f1 = _dafny.ZERO;
      this._f0 = _dafny.ZERO;
      this._f2 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2) {
      let _this = this;
      (_this)._f0 = f0;
      (_this).f1 = f1;
      (_this)._f2 = f2;
      return;
    }
    get f0() {
      let _this = this;
      return _this._f0;
    };
    get f2() {
      let _this = this;
      return _this._f2;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f3 = [];
      this._f4 = _dafny.ZERO;
      this._f7 = _module.D0.Default();
      this._f8 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f3() {
      let _this = this;
      return _this._f3;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    __ctor(f7, f8, f3, f4) {
      let _this = this;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      return;
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f3 = [];
      this._f4 = _dafny.ZERO;
      this._f5 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f6 = [];
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f3() {
      let _this = this;
      return _this._f3;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    __ctor(f5, f6, f3, f4) {
      let _this = this;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      return;
    }
    m0(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = false;
      let _228_v0;
      _228_v0 = _dafny.Seq.UnicodeFromString("lwcdnhtm");
      let _229_v1;
      _229_v1 = _dafny.Set.fromElements(_228_v0);
      let _230_v2;
      _230_v2 = false;
      let _231_v3;
      _231_v3 = _module.D0.create_DC0(_230_v2);
      let _232_v4;
      _232_v4 = _dafny.Map.Empty.slice().updateUnsafe(_229_v1,(_231_v3).dtor_cf0);
      _232_v4 = (_232_v4).update((_229_v1).Union(_229_v1), _230_v2);
      let _233_v5;
      _233_v5 = _dafny.Seq.of((_this).f4, new BigNumber(930), (_this).f4, (_this).f4, (_this).f4);
      let _234_v6;
      _234_v6 = _dafny.Map.Empty.slice().updateUnsafe(_230_v2,(_dafny.ZERO).minus((_this).f4));
      let _235_v7;
      _235_v7 = _dafny.MultiSet.fromElements((_this).f4, new BigNumber((_234_v6).length), new BigNumber((_228_v0).length), (_this).f4);
      (globalState).f1 = ((!(!(_230_v2))) ? ((_dafny.ZERO).minus(new BigNumber((_233_v5).length))) : (_module.__default.safeModuloInt((_this).f4, new BigNumber((_235_v7).cardinality()))));
      let _index30 = _module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length));
      ((_this).f6)[_index30] = _230_v2;
      let _236_v8;
      _236_v8 = _dafny.Map.Empty.slice().updateUnsafe(_230_v2,!(_230_v2));
      let _index31 = _module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length));
      ((_this).f6)[_index31] = !_dafny.Seq.contains(_233_v5, new BigNumber((((_230_v2) ? (_dafny.Map.Empty.slice().updateUnsafe(_230_v2,_230_v2)) : (_236_v8))).length));
      let _237_v9;
      _237_v9 = _module.D0.create_DC2(_228_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))],_dafny.Map.Empty.slice().updateUnsafe(true,_230_v2))).length), new BigNumber((_228_v0).length));
      let _238_v10;
      _238_v10 = _dafny.Seq.of((_237_v9).dtor_cf4, _228_v0);
      let _hi0 = new BigNumber(((_238_v10)[_module.__default.safeIndex((_this).f4, new BigNumber((_238_v10).length))]).length);
      for (let _239_i0 = new BigNumber(-768); _239_i0.isLessThan(_hi0); _239_i0 = _239_i0.plus(_dafny.ONE)) {
        let _240_v11;
        _240_v11 = _dafny.Seq.of(_230_v2);
        let _index32 = _module.__default.safeIndex(new BigNumber(159), new BigNumber(((_this).f3).length));
        ((_this).f3)[_index32] = (new BigNumber((_235_v7).cardinality())).plus(new BigNumber((_240_v11).length));
        let _index33 = _module.__default.safeIndex(new BigNumber(159), new BigNumber(((_this).f3).length));
        let _rhs36 = (_237_v9).dtor_cf6;
        let _rhs37 = (_module.__default.safeModuloInt((_this).f4, (_this).f4)).isEqualTo(new BigNumber(288));
        let _rhs38 = _230_v2;
        let _rhs39 = new BigNumber((_228_v0).length);
        let _lhs20 = globalState;
        let _lhs21 = (_this).f3;
        let _lhs22 = _module.__default.safeIndex(new BigNumber(159), new BigNumber(((_this).f3).length));
        _lhs20.f1 = _rhs36;
        r1 = _rhs37;
        r1 = _rhs38;
        _lhs21[_lhs22] = _rhs39;
        (globalState).f1 = _module.__default.safeModuloInt((new BigNumber(-468)).minus(((_this).f3)[_module.__default.safeIndex(new BigNumber(159), new BigNumber(((_this).f3).length))]), (_this).f4);
        let _241_v12;
        _241_v12 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))],_dafny.Seq.update(_module.__default.fm1(((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))], _230_v2, globalState), _module.__default.safeIndex(((_this).f3)[_module.__default.safeIndex(new BigNumber(159), new BigNumber(((_this).f3).length))], new BigNumber((_module.__default.fm1(((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))], _230_v2, globalState)).length)), (_this).f5));
        let _index34 = _module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length));
        ((_this).f6)[_index34] = (((_236_v8).contains(((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))])) ? ((_236_v8).get(((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))])) : (!((_241_v12).contains(_230_v2))));
        let _242_v13;
        let _nw25 = Array((new BigNumber(11)).toNumber()).fill([]);
        _242_v13 = _nw25;
        let _243_v14;
        let _init3 = ((_244_v3) => function (_245_i1) {
          return _dafny.Seq.of((_244_v3).dtor_cf0);
        })(_231_v3);
        let _nw26 = Array((new BigNumber(5)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw26.length); _i0_3++) {
          _nw26[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _243_v14 = _nw26;
        let _index35 = _module.__default.safeIndex(new BigNumber(11), new BigNumber((_242_v13).length));
        (_242_v13)[_index35] = _243_v14;
        let _index36 = _module.__default.safeIndex(new BigNumber(11), new BigNumber((_242_v13).length));
        (_242_v13)[_index36] = _243_v14;
      }
      if (!(((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))])) {
        r2 = false;
        let _246_v15;
        _246_v15 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f4).plus((_this).f4),((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))]);
        _246_v15 = (_246_v15).update((_this).f4, ((_this).f4).isLessThanOrEqualTo((_this).f4));
        let _247_v16;
        _247_v16 = _dafny.Seq.of(((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))], _230_v2);
        let _248_v17;
        _248_v17 = _dafny.Seq.of((((_247_v16)[_module.__default.safeIndex((((_234_v6).contains(((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))])) ? ((_234_v6).get(((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))])) : ((_this).f4)), new BigNumber((_247_v16).length))]) ? (_247_v16) : (_dafny.Seq.of(false))));
        _248_v17 = _dafny.Seq.Concat(_dafny.Seq.Concat(_248_v17, _248_v17), _dafny.Seq.update(_248_v17, _module.__default.safeIndex((_this).f4, new BigNumber((_248_v17).length)), _247_v16));
        (globalState).f1 = (_this).f4;
        if (_230_v2) {
          (globalState).f1 = (_this).f4;
          let _249_v18;
          _249_v18 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))],_246_v15);
          _249_v18 = (_249_v18).Merge((_249_v18).Merge(_module.__default.fm2((_this).f4, !(((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))]), _230_v2, globalState)));
          _235_v7 = _235_v7;
          let _index37 = _module.__default.safeIndex(new BigNumber(821), new BigNumber(((_this).f3).length));
          ((_this).f3)[_index37] = (_this).f4;
          let _index38 = _module.__default.safeIndex(new BigNumber(821), new BigNumber(((_this).f3).length));
          ((_this).f3)[_index38] = (_this).f4;
          (globalState).f1 = ((_this).f3)[_module.__default.safeIndex(new BigNumber(821), new BigNumber(((_this).f3).length))];
        } else {
          let _250_v19;
          _250_v19 = _dafny.Set.fromElements(_230_v2);
          let _index39 = _module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length));
          ((_this).f6)[_index39] = (_dafny.Set.fromElements(false)).IsProperSubsetOf(_250_v19);
          let _251_v20;
          let _nw27 = Array((new BigNumber(24)).toNumber()).fill(_dafny.MultiSet.Empty);
          _251_v20 = _nw27;
          let _252_v21;
          _252_v21 = _dafny.MultiSet.fromElements(_module.D0.create_DC0(((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))]));
          let _index40 = _module.__default.safeIndex(new BigNumber(836), new BigNumber((_251_v20).length));
          (_251_v20)[_index40] = _252_v21;
          let _253_v22;
          _253_v22 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,_246_v15);
          let _254_v23;
          _254_v23 = _dafny.MultiSet.fromElements(_module.__default.fm3((_this).f5, (_this).f4, globalState), _module.__default.fm3((_this).f5, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(601)), function (_255_i2) {
            return (_this).f5;
          })).length), globalState), (_this).f5, (_this).f5);
          let _256_v24;
          _256_v24 = _dafny.Map.Empty.slice().updateUnsafe((_253_v22).equals(_253_v22),_254_v23);
          let _257_v25;
          _257_v25 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_this).f4);
          let _index41 = _module.__default.safeIndex(new BigNumber(836), new BigNumber((_251_v20).length));
          let _rhs40 = (_dafny.ZERO).minus(new BigNumber(((((_256_v24).contains(!(_230_v2))) ? ((_256_v24).get(!(_230_v2))) : (_254_v23))).cardinality()));
          let _rhs41 = (((_257_v25).contains((_this).f4)) ? ((_257_v25).get((_this).f4)) : (_module.__default.safeDivisionInt((_this).f4, new BigNumber(-768))));
          let _rhs42 = _252_v21;
          let _lhs23 = globalState;
          let _lhs24 = globalState;
          let _lhs25 = _251_v20;
          let _lhs26 = _module.__default.safeIndex(new BigNumber(836), new BigNumber((_251_v20).length));
          _lhs23.f1 = _rhs40;
          _lhs24.f1 = _rhs41;
          _lhs25[_lhs26] = _rhs42;
          (globalState).f1 = new BigNumber((_228_v0).length);
          let _258_v26;
          let _nw28 = Array((new BigNumber(21)).toNumber()).fill(false);
          _258_v26 = _nw28;
          _258_v26 = _258_v26;
          let _259_v27;
          _259_v27 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))],(_this).f3);
          let _260_v28;
          let _nw29 = new _module.C0();
          _nw29.__ctor(_module.D0.create_DC0(true), (_259_v27).Merge(_259_v27), (_this).f3, (_this).f4);
          _260_v28 = _nw29;
        }
      } else {
        (globalState).f1 = (_module.__default.safeDivisionInt((_this).f4, new BigNumber((_dafny.Seq.of(_230_v2)).length))).minus((_dafny.ZERO).minus((_this).f4));
        _228_v0 = _dafny.Seq.UnicodeFromString("blty");
        r1 = (true) && (!((_231_v3).dtor_cf0));
        let _261_v29;
        _261_v29 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))],(_this).f3);
        let _262_v30;
        let _nw30 = new _module.C0();
        _nw30.__ctor(_231_v3, _261_v29, (_this).f3, (_this).f4);
        _262_v30 = _nw30;
        r0 = ((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))];
      }
      let _hi1 = new BigNumber((_dafny.Seq.UnicodeFromString("hki")).length);
      for (let _263_i3 = new BigNumber((_module.__default.fm4((((_236_v8).contains(_230_v2)) ? ((_236_v8).get(_230_v2)) : (_230_v2)), new BigNumber((_dafny.Seq.of(_230_v2, _230_v2, _230_v2)).length), (_this).f4, globalState)).length); _263_i3.isLessThan(_hi1); _263_i3 = _263_i3.plus(_dafny.ONE)) {
        let _264_v31;
        let _nw31 = Array((_dafny.ONE).toNumber()).fill(false);
        _264_v31 = _nw31;
        let _265_v32;
        _265_v32 = _dafny.Seq.of(_264_v31);
        let _rhs43 = !(((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))]);
        let _rhs44 = _228_v0;
        let _rhs45 = _dafny.Seq.Concat(_265_v32, _265_v32);
        r0 = _rhs43;
        _228_v0 = _rhs44;
        _265_v32 = _rhs45;
        let _266_v33;
        _266_v33 = _dafny.Set.fromElements(_264_v31, _264_v31, _264_v31, _264_v31, _264_v31);
        let _267_v34;
        let _init4 = ((_268_v0) => function (_269_i4) {
          return _268_v0;
        })(_228_v0);
        let _nw32 = Array((new BigNumber(16)).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw32.length); _i0_4++) {
          _nw32[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _267_v34 = _nw32;
        let _index42 = _module.__default.safeIndex(new BigNumber(767), new BigNumber((_267_v34).length));
        (_267_v34)[_index42] = _dafny.Seq.UnicodeFromString("vcbn");
        let _index43 = _module.__default.safeIndex(new BigNumber(767), new BigNumber((_267_v34).length));
        let _rhs46 = _230_v2;
        let _rhs47 = _266_v33;
        let _rhs48 = ((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))];
        let _rhs49 = _228_v0;
        let _lhs27 = _267_v34;
        let _lhs28 = _module.__default.safeIndex(new BigNumber(767), new BigNumber((_267_v34).length));
        r0 = _rhs46;
        _266_v33 = _rhs47;
        r1 = _rhs48;
        _lhs27[_lhs28] = _rhs49;
        r0 = _230_v2;
        let _source5 = _237_v9;
        if (_source5.is_DC1) {
          let _270___mcc_h0 = (_source5).cf1;
          let _271___mcc_h1 = (_source5).cf2;
          let _272___mcc_h2 = (_source5).cf3;
          let _273_cf3 = _272___mcc_h2;
          let _274_cf2 = _271___mcc_h1;
          let _275_cf1 = _270___mcc_h0;
          let _276_v35;
          let _nw33 = Array((new BigNumber(17)).toNumber()).fill([]);
          _276_v35 = _nw33;
          let _index44 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_276_v35).length));
          (_276_v35)[_index44] = _264_v31;
          let _index45 = _module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length));
          let _index46 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_276_v35).length));
          let _rhs50 = _dafny.Seq.contains(_228_v0, (_this).f5);
          let _rhs51 = (_265_v32)[_module.__default.safeIndex((_this).f4, new BigNumber((_265_v32).length))];
          let _rhs52 = _230_v2;
          let _lhs29 = (_this).f6;
          let _lhs30 = _module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length));
          let _lhs31 = _276_v35;
          let _lhs32 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_276_v35).length));
          _lhs29[_lhs30] = _rhs50;
          _lhs31[_lhs32] = _rhs51;
          r2 = _rhs52;
          let _277_v36;
          _277_v36 = _dafny.Set.fromElements((_this).f4, _274_cf2);
          _277_v36 = (_277_v36).Union(_277_v36);
          let _278_v37;
          _278_v37 = _dafny.Seq.of(_231_v3, _231_v3, _module.D0.create_DC0(_230_v2));
          let _279_v38;
          _279_v38 = new _dafny.CodePoint('h'.codePointAt(0));
          let _280_v39;
          _280_v39 = _module.D1.create_DC4((_this).f3);
          let _281_v40;
          _281_v40 = _dafny.Seq.of((_this).f3);
          let _282_v41;
          let _nw34 = Array((new BigNumber(28)).toNumber());
          _nw34[(_dafny.ZERO).toNumber()] = (_this).f3;
          _nw34[(_dafny.ONE).toNumber()] = (_280_v39).dtor_cf8;
          _nw34[(new BigNumber(2)).toNumber()] = (_this).f3;
          _nw34[(new BigNumber(3)).toNumber()] = (_this).f3;
          _nw34[(new BigNumber(4)).toNumber()] = (_this).f3;
          _nw34[(new BigNumber(5)).toNumber()] = (_this).f3;
          _nw34[(new BigNumber(6)).toNumber()] = (_this).f3;
          _nw34[(new BigNumber(7)).toNumber()] = (_this).f3;
          _nw34[(new BigNumber(8)).toNumber()] = (_this).f3;
          _nw34[(new BigNumber(9)).toNumber()] = (_this).f3;
          _nw34[(new BigNumber(10)).toNumber()] = (_this).f3;
          _nw34[(new BigNumber(11)).toNumber()] = (_this).f3;
          _nw34[(new BigNumber(12)).toNumber()] = (_this).f3;
          _nw34[(new BigNumber(13)).toNumber()] = (_this).f3;
          _nw34[(new BigNumber(14)).toNumber()] = (_this).f3;
          _nw34[(new BigNumber(15)).toNumber()] = (_this).f3;
          _nw34[(new BigNumber(16)).toNumber()] = (_this).f3;
          _nw34[(new BigNumber(17)).toNumber()] = (_this).f3;
          _nw34[(new BigNumber(18)).toNumber()] = (_this).f3;
          _nw34[(new BigNumber(19)).toNumber()] = (_this).f3;
          _nw34[(new BigNumber(20)).toNumber()] = (_this).f3;
          _nw34[(new BigNumber(21)).toNumber()] = (_this).f3;
          _nw34[(new BigNumber(22)).toNumber()] = (_281_v40)[_module.__default.safeIndex(new BigNumber(347), new BigNumber((_281_v40).length))];
          _nw34[(new BigNumber(23)).toNumber()] = (_this).f3;
          _nw34[(new BigNumber(24)).toNumber()] = (_this).f3;
          _nw34[(new BigNumber(25)).toNumber()] = (_this).f3;
          _nw34[(new BigNumber(26)).toNumber()] = (_281_v40)[_module.__default.safeIndex(new BigNumber((_233_v5).length), new BigNumber((_281_v40).length))];
          _nw34[(new BigNumber(27)).toNumber()] = (_this).f3;
          _282_v41 = _nw34;
          let _283_v42;
          _283_v42 = _dafny.Seq.of(_282_v41, _282_v41, _282_v41, _282_v41, _282_v41);
          let _284_v43;
          let _nw35 = Array((new BigNumber(18)).toNumber());
          _nw35[(_dafny.ZERO).toNumber()] = _282_v41;
          _nw35[(_dafny.ONE).toNumber()] = _282_v41;
          _nw35[(new BigNumber(2)).toNumber()] = _282_v41;
          _nw35[(new BigNumber(3)).toNumber()] = _282_v41;
          _nw35[(new BigNumber(4)).toNumber()] = _282_v41;
          _nw35[(new BigNumber(5)).toNumber()] = ((_module.__default.fm0((_this).f4, _274_cf2, globalState)) ? (_282_v41) : (_282_v41));
          _nw35[(new BigNumber(6)).toNumber()] = _282_v41;
          _nw35[(new BigNumber(7)).toNumber()] = _282_v41;
          _nw35[(new BigNumber(8)).toNumber()] = (_283_v42)[_module.__default.safeIndex(_274_cf2, new BigNumber((_283_v42).length))];
          _nw35[(new BigNumber(9)).toNumber()] = _282_v41;
          _nw35[(new BigNumber(10)).toNumber()] = _282_v41;
          _nw35[(new BigNumber(11)).toNumber()] = _282_v41;
          _nw35[(new BigNumber(12)).toNumber()] = _282_v41;
          _nw35[(new BigNumber(13)).toNumber()] = _282_v41;
          _nw35[(new BigNumber(14)).toNumber()] = ((((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))]) ? (_282_v41) : (_282_v41));
          _nw35[(new BigNumber(15)).toNumber()] = _282_v41;
          _nw35[(new BigNumber(16)).toNumber()] = _282_v41;
          _nw35[(new BigNumber(17)).toNumber()] = _282_v41;
          _284_v43 = _nw35;
          let _index47 = _module.__default.safeIndex(new BigNumber(787), new BigNumber((_284_v43).length));
          (_284_v43)[_index47] = _282_v41;
          let _index48 = _module.__default.safeIndex(new BigNumber(787), new BigNumber((_284_v43).length));
          let _rhs53 = _dafny.Seq.update(_dafny.Seq.of(_module.D0.create_DC0(((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))]), function (_pat_let6_0) {
            return function (_285_dt__update__tmp_h0) {
              return function (_pat_let7_0) {
                return function (_286_dt__update_hcf0_h0) {
                  return _module.D0.create_DC0(_286_dt__update_hcf0_h0);
                }(_pat_let7_0);
              }(((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))]);
            }(_pat_let6_0);
          }(_module.D0.create_DC0(((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))]))), _module.__default.safeIndex(_275_cf1, new BigNumber((_dafny.Seq.of(_module.D0.create_DC0(((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))]), function (_pat_let8_0) {
            return function (_287_dt__update__tmp_h1) {
              return function (_pat_let9_0) {
                return function (_288_dt__update_hcf0_h1) {
                  return _module.D0.create_DC0(_288_dt__update_hcf0_h1);
                }(_pat_let9_0);
              }(((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))]);
            }(_pat_let8_0);
          }(_module.D0.create_DC0(((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))])))).length)), _231_v3);
          let _rhs54 = _279_v38;
          let _rhs55 = _282_v41;
          let _lhs33 = _284_v43;
          let _lhs34 = _module.__default.safeIndex(new BigNumber(787), new BigNumber((_284_v43).length));
          _278_v37 = _rhs53;
          _279_v38 = _rhs54;
          _lhs33[_lhs34] = _rhs55;
          let _index49 = _module.__default.safeIndex(new BigNumber(499), new BigNumber(((_this).f3).length));
          ((_this).f3)[_index49] = _263_i3;
          let _index50 = _module.__default.safeIndex(new BigNumber(499), new BigNumber(((_this).f3).length));
          ((_this).f3)[_index50] = (_dafny.ZERO).minus(_275_cf1);
        } else if (_source5.is_DC2) {
          let _289___mcc_h3 = (_source5).cf4;
          let _290___mcc_h4 = (_source5).cf5;
          let _291___mcc_h5 = (_source5).cf6;
          let _292_cf6 = _291___mcc_h5;
          let _293_cf5 = _290___mcc_h4;
          let _294_cf4 = _289___mcc_h3;
          (globalState).f1 = _293_cf5;
          let _295_v44;
          _295_v44 = _dafny.Map.Empty.slice().updateUnsafe(_230_v2,(_this).f3);
          let _296_v45;
          let _nw36 = new _module.C0();
          _nw36.__ctor(_231_v3, _295_v44, (_this).f3, _module.__default.safeDivisionInt((_this).f4, _263_i3));
          _296_v45 = _nw36;
          _296_v45 = _296_v45;
          let _297_v46;
          let _out15;
          _out15 = (_this).m1(((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))], new BigNumber(484), (_dafny.ZERO).minus(_263_i3), globalState);
          _297_v46 = _out15;
          let _298_v47;
          let _nw37 = new _module.C0();
          _nw37.__ctor((_296_v45).f7, (_295_v44).Merge((_296_v45).f8), (_this).f3, _263_i3);
          _298_v47 = _nw37;
        } else if (_source5.is_DC0) {
          let _299___mcc_h6 = (_source5).cf0;
          let _300_cf0 = _299___mcc_h6;
          let _301_v48;
          _301_v48 = _dafny.Map.Empty.slice().updateUnsafe(_300_cf0,(_this).f3);
          let _302_v49;
          let _nw38 = new _module.C0();
          _nw38.__ctor(_231_v3, _301_v48, (_this).f3, ((_230_v2) ? (new BigNumber((_228_v0).length)) : ((_this).f4)));
          _302_v49 = _nw38;
          let _303_v50;
          _303_v50 = _module.D1.create_DC6();
          let _304_v51;
          _304_v51 = _module.D0.create_DC1(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_230_v2,_263_i3)).length), new BigNumber((_233_v5).length), new BigNumber(852));
          let _305_v52;
          _305_v52 = _dafny.Map.Empty.slice().updateUnsafe(_303_v50,_304_v51);
          _305_v52 = (_305_v52).update(_module.D1.create_DC6(), _304_v51);
          let _306_v53;
          _306_v53 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_267_v34)[_module.__default.safeIndex(new BigNumber(767), new BigNumber((_267_v34).length))]).length),true);
          _236_v8 = (_236_v8).update((((((_236_v8).contains(_300_cf0)) ? ((_236_v8).get(_300_cf0)) : (_300_cf0))) ? ((((_306_v53).contains(_263_i3)) ? ((_306_v53).get(_263_i3)) : (_module.__default.fm0((_dafny.ZERO).minus(_module.__default.fm5(((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))], globalState)), (_this).f4, globalState)))) : (_230_v2)), _300_cf0);
          let _307_v54;
          _307_v54 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,new BigNumber((_dafny.Set.fromElements(_230_v2, _230_v2)).length));
          let _308_v55;
          _308_v55 = _dafny.Map.Empty.slice().updateUnsafe(_230_v2,_236_v8);
          let _rhs56 = _module.__default.fm5(false, globalState);
          let _rhs57 = _307_v54;
          let _rhs58 = ((_263_i3).isLessThanOrEqualTo(new BigNumber((_module.__default.fm6(true, ((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))], _263_i3, globalState)).length))) || (_230_v2);
          let _rhs59 = ((((_this).f4).isLessThanOrEqualTo(_263_i3)) ? (_308_v55) : ((_308_v55).Merge(_308_v55)));
          let _rhs60 = _231_v3;
          let _lhs35 = globalState;
          _lhs35.f1 = _rhs56;
          _307_v54 = _rhs57;
          _300_cf0 = _rhs58;
          _308_v55 = _rhs59;
          _231_v3 = _rhs60;
        } else {
          let _309___mcc_h7 = (_source5).cf7;
          let _310_cf7 = _309___mcc_h7;
          let _311_v56;
          let _nw39 = Array((new BigNumber(29)).toNumber()).fill([]);
          _311_v56 = _nw39;
          let _312_v57;
          let _nw40 = Array((new BigNumber(22)).toNumber()).fill([]);
          _312_v57 = _nw40;
          let _index51 = _module.__default.safeIndex(new BigNumber(27), new BigNumber((_311_v56).length));
          (_311_v56)[_index51] = _312_v57;
          let _index52 = _module.__default.safeIndex(new BigNumber(27), new BigNumber((_311_v56).length));
          (_311_v56)[_index52] = _312_v57;
          let _313_v58;
          _313_v58 = new _dafny.CodePoint('r'.codePointAt(0));
          _313_v58 = _313_v58;
          (globalState).f1 = _module.__default.safeDivisionInt((_this).f4, new BigNumber((_module.__default.fm7(((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))], globalState)).length));
          (globalState).f1 = new BigNumber(-569);
        }
      }
      let _314_v59;
      _314_v59 = _dafny.Seq.of(((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))], ((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))], _230_v2);
      r0 = (_314_v59)[_module.__default.safeIndex((_this).f4, new BigNumber((_314_v59).length))];
      r1 = _230_v2;
      r2 = ((_this).f6)[_module.__default.safeIndex(new BigNumber(592), new BigNumber(((_this).f6).length))];
      return [r0, r1, r2];
    }
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let _315_i0;
      _315_i0 = _dafny.ZERO;
      L0: {
        while (p0) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_315_i0)) {
              break L0;
            }
            _315_i0 = (_315_i0).plus(_dafny.ONE);
            let _316_v0;
            let _nw41 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Map.Empty);
            _316_v0 = _nw41;
            _316_v0 = _316_v0;
            let _317_v1;
            _317_v1 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_this).f4);
            let _318_v2;
            _318_v2 = _dafny.Seq.of(_317_v1, _317_v1, _317_v1);
            let _319_v3;
            _319_v3 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-227)),(_318_v2)[_module.__default.safeIndex((_this).f4, new BigNumber((_318_v2).length))]);
            let _index53 = _module.__default.safeIndex(new BigNumber(803), new BigNumber(((_this).f3).length));
            ((_this).f3)[_index53] = new BigNumber(((((_319_v3).contains(p1)) ? ((_319_v3).get(p1)) : (_317_v1))).length);
            let _index54 = _module.__default.safeIndex(new BigNumber(439), new BigNumber(((_this).f6).length));
            ((_this).f6)[_index54] = p0;
            let _index55 = _module.__default.safeIndex(new BigNumber(803), new BigNumber(((_this).f3).length));
            let _index56 = _module.__default.safeIndex(new BigNumber(439), new BigNumber(((_this).f6).length));
            let _rhs61 = p1;
            let _rhs62 = !(p0);
            let _lhs36 = (_this).f3;
            let _lhs37 = _module.__default.safeIndex(new BigNumber(803), new BigNumber(((_this).f3).length));
            let _lhs38 = (_this).f6;
            let _lhs39 = _module.__default.safeIndex(new BigNumber(439), new BigNumber(((_this).f6).length));
            _lhs36[_lhs37] = _rhs61;
            _lhs38[_lhs39] = _rhs62;
            let _320_v4;
            _320_v4 = _dafny.Seq.UnicodeFromString("k");
            let _321_v5;
            _321_v5 = _dafny.Seq.of(_320_v4, _320_v4);
            let _322_v6;
            _322_v6 = _module.D0.create_DC0(p0);
            let _323_v7;
            let _init5 = function (_324_i1) {
              return (_324_i1).multipliedBy((_this).f4);
            };
            let _nw42 = Array((new BigNumber(19)).toNumber());
            for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw42.length); _i0_5++) {
              _nw42[_i0_5] = _init5(new BigNumber(_i0_5));
            }
            _323_v7 = _nw42;
            let _325_v8;
            _325_v8 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f6)[_module.__default.safeIndex(new BigNumber(439), new BigNumber(((_this).f6).length))],_323_v7);
            let _326_v9;
            _326_v9 = _dafny.Seq.of(p0);
            let _327_v10;
            let _nw43 = new _module.C0();
            _nw43.__ctor(_322_v6, _325_v8, _323_v7, new BigNumber((_326_v9).length));
            _327_v10 = _nw43;
            let _328_v11;
            _328_v11 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.Concat(_321_v5, _321_v5), _module.__default.safeIndex(((_this).f3)[_module.__default.safeIndex(new BigNumber(803), new BigNumber(((_this).f3).length))], new BigNumber((_dafny.Seq.Concat(_321_v5, _321_v5)).length)), _320_v4),_327_v10);
            let _329_v12;
            _329_v12 = _dafny.Map.Empty.slice().updateUnsafe(p2,_321_v5);
            let _330_v13;
            _330_v13 = _dafny.Map.Empty.slice().updateUnsafe(p0,_327_v10);
            _328_v11 = (_328_v11).update((((_329_v12).contains((_327_v10).f4)) ? ((_329_v12).get((_327_v10).f4)) : (_321_v5)), (((_330_v13).contains(!(p0))) ? ((_330_v13).get(!(p0))) : (_327_v10)));
            let _331_v14;
            _331_v14 = _dafny.Map.Empty.slice().updateUnsafe(_320_v4,_dafny.Seq.UnicodeFromString("hcftnkrjh"));
            _331_v14 = (_331_v14).update(_320_v4, _320_v4);
          }
        }
      }
      let _332_v15;
      let _nw44 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.of());
      _332_v15 = _nw44;
      _332_v15 = _332_v15;
      let _333_v16;
      let _init6 = function (_334_i3) {
        return _module.__default.safeDivisionInt(_334_i3, new BigNumber((_dafny.Seq.UnicodeFromString("noju")).length));
      };
      let _nw45 = Array((new BigNumber(17)).toNumber());
      for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw45.length); _i0_6++) {
        _nw45[_i0_6] = _init6(new BigNumber(_i0_6));
      }
      _333_v16 = _nw45;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_333_v16).length))) {
        let _335_i2 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_335_i2)) && ((_335_i2).isLessThan(new BigNumber((_333_v16).length))))) {
          (_333_v16)[(_335_i2)] = (_335_i2).plus((_dafny.ZERO).minus((_this).f4));
        }
      }
      if (p0) {
        let _336_v17;
        _336_v17 = _module.D1.create_DC4((_this).f3);
        let _337_v18;
        _337_v18 = _dafny.MultiSet.fromElements((_336_v17).dtor_cf8);
        let _338_v19;
        _338_v19 = _dafny.Seq.of(_dafny.MultiSet.fromElements(_333_v16), _337_v18, _337_v18);
        let _339_v20;
        _339_v20 = _module.D0.create_DC0(p0);
        let _340_v21;
        _340_v21 = _dafny.Map.Empty.slice().updateUnsafe(p0,_333_v16);
        let _341_v22;
        let _nw46 = new _module.C0();
        _nw46.__ctor(_339_v20, _340_v21, (_this).f3, p2);
        _341_v22 = _nw46;
        let _342_v23;
        _342_v23 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.MultiSet.fromElements((_336_v17).dtor_cf8)).IsProperSubsetOf((_338_v19)[_module.__default.safeIndex(p1, new BigNumber((_338_v19).length))]),_341_v22);
        let _343_v24;
        _343_v24 = false;
        let _344_v25;
        _344_v25 = new _dafny.CodePoint('p'.codePointAt(0));
        let _345_v26;
        _345_v26 = _dafny.Map.Empty.slice().updateUnsafe(_344_v25,(_this).f5);
        let _346_v27;
        _346_v27 = _dafny.Seq.of((((_345_v26).contains(_344_v25)) ? ((_345_v26).get(_344_v25)) : (_344_v25)));
        let _347_v28;
        _347_v28 = _dafny.Set.fromElements(p1);
        let _348_v29;
        let _nw47 = Array((new BigNumber(23)).toNumber());
        _nw47[(_dafny.ZERO).toNumber()] = _346_v27;
        _nw47[(_dafny.ONE).toNumber()] = _346_v27;
        _nw47[(new BigNumber(2)).toNumber()] = _dafny.Seq.update(((_343_v24) ? (_346_v27) : (_dafny.Seq.of(_344_v25, (_this).f5))), _module.__default.safeIndex(p2, new BigNumber((((_343_v24) ? (_346_v27) : (_dafny.Seq.of(_344_v25, (_this).f5)))).length)), new _dafny.CodePoint('l'.codePointAt(0)));
        _nw47[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(940)), ((_349_v25) => function (_350_i4) {
          return _349_v25;
        })(_344_v25)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(92)), ((_351_v25) => function (_352_i5) {
          return _351_v25;
        })(_344_v25)));
        _nw47[(new BigNumber(4)).toNumber()] = _346_v27;
        _nw47[(new BigNumber(5)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_346_v27, _346_v27), _module.__default.safeIndex((_dafny.ZERO).minus(p2), new BigNumber((_dafny.Seq.Concat(_346_v27, _346_v27)).length)), (_346_v27)[_module.__default.safeIndex(p1, new BigNumber((_346_v27).length))]);
        _nw47[(new BigNumber(6)).toNumber()] = _346_v27;
        _nw47[(new BigNumber(7)).toNumber()] = _346_v27;
        _nw47[(new BigNumber(8)).toNumber()] = _346_v27;
        _nw47[(new BigNumber(9)).toNumber()] = _346_v27;
        _nw47[(new BigNumber(10)).toNumber()] = _346_v27;
        _nw47[(new BigNumber(11)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(924)), function (_353_i6) {
          return (_this).f5;
        });
        _nw47[(new BigNumber(12)).toNumber()] = _346_v27;
        _nw47[(new BigNumber(13)).toNumber()] = _346_v27;
        _nw47[(new BigNumber(14)).toNumber()] = _346_v27;
        _nw47[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_346_v27, _346_v27);
        _nw47[(new BigNumber(16)).toNumber()] = _346_v27;
        _nw47[(new BigNumber(17)).toNumber()] = ((p0) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(938)), function (_354_i7) {
          return (_this).f5;
        })) : (_dafny.Seq.of((_this).f5)));
        _nw47[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_346_v27, _346_v27);
        _nw47[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_346_v27, _module.__default.safeIndex(new BigNumber((_347_v28).length), new BigNumber((_346_v27).length)), (_this).f5), _346_v27);
        _nw47[(new BigNumber(20)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(516)), ((_355_v25) => function (_356_i8) {
          return _355_v25;
        })(_344_v25));
        _nw47[(new BigNumber(21)).toNumber()] = _346_v27;
        _nw47[(new BigNumber(22)).toNumber()] = _dafny.Seq.of((_this).f5);
        _348_v29 = _nw47;
        let _357_v30;
        _357_v30 = _dafny.Seq.of(_346_v27);
        let _index57 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_348_v29).length));
        (_348_v29)[_index57] = (_357_v30)[_module.__default.safeIndex(p2, new BigNumber((_357_v30).length))];
        let _358_v31;
        _358_v31 = _dafny.Map.Empty.slice().updateUnsafe(true,p1);
        let _index58 = _module.__default.safeIndex(new BigNumber(381), new BigNumber(((_this).f3).length));
        ((_this).f3)[_index58] = (((_358_v31).contains(p0)) ? ((_358_v31).get(p0)) : (p1));
        let _359_v32;
        _359_v32 = _dafny.Seq.of(_343_v24, p0);
        let _index59 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_348_v29).length));
        let _index60 = _module.__default.safeIndex(new BigNumber(381), new BigNumber(((_this).f3).length));
        let _rhs63 = (_342_v23).update(_343_v24, _341_v22);
        let _rhs64 = p0;
        let _rhs65 = _344_v25;
        let _rhs66 = _module.__default.fm1(p0, (new BigNumber((_359_v32).length)).isLessThanOrEqualTo(_module.__default.fm5(_343_v24, globalState)), globalState);
        let _rhs67 = new BigNumber(157);
        let _lhs40 = _348_v29;
        let _lhs41 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_348_v29).length));
        let _lhs42 = (_this).f3;
        let _lhs43 = _module.__default.safeIndex(new BigNumber(381), new BigNumber(((_this).f3).length));
        _342_v23 = _rhs63;
        _343_v24 = _rhs64;
        _344_v25 = _rhs65;
        _lhs40[_lhs41] = _rhs66;
        _lhs42[_lhs43] = _rhs67;
        if (p0) {
          _343_v24 = _343_v24;
          let _360_v33;
          let _init7 = ((_361_p0, _362_v24) => function (_363_i9) {
            return ((_361_p0) ? (_361_p0) : (_362_v24));
          })(p0, _343_v24);
          let _nw48 = Array((new BigNumber(17)).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw48.length); _i0_7++) {
            _nw48[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _360_v33 = _nw48;
          let _364_v34;
          let _init8 = ((_365_v24) => function (_366_i10) {
            return _dafny.Seq.of(_365_v24);
          })(_343_v24);
          let _nw49 = Array((new BigNumber(16)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw49.length); _i0_8++) {
            _nw49[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _364_v34 = _nw49;
          let _index61 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_348_v29).length));
          let _rhs68 = (_this).f6;
          let _rhs69 = (_this).f5;
          let _rhs70 = _364_v34;
          let _rhs71 = (_348_v29)[_module.__default.safeIndex(new BigNumber(266), new BigNumber((_348_v29).length))];
          let _rhs72 = _module.__default.fm0(p2, p2, globalState);
          let _lhs44 = _348_v29;
          let _lhs45 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_348_v29).length));
          _360_v33 = _rhs68;
          _344_v25 = _rhs69;
          _364_v34 = _rhs70;
          _lhs44[_lhs45] = _rhs71;
          _343_v24 = _rhs72;
          let _index62 = _module.__default.safeIndex(new BigNumber(381), new BigNumber(((_this).f3).length));
          ((_this).f3)[_index62] = (((p0) || (_343_v24)) ? (_module.__default.fm5(!(_343_v24), globalState)) : (new BigNumber(469)));
          (globalState).f1 = new BigNumber(683);
          let _367_v35;
          _367_v35 = _dafny.MultiSet.fromElements(p0);
          let _368_v36;
          _368_v36 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0((_this).f4, ((_this).f3)[_module.__default.safeIndex(new BigNumber(381), new BigNumber(((_this).f3).length))], globalState),_367_v35);
          _368_v36 = _dafny.Map.Empty.slice().updateUnsafe(_343_v24,_367_v35);
        } else {
          (globalState).f1 = _module.__default.fm5(p0, globalState);
          let _369_v37;
          _369_v37 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _index63 = _module.__default.safeIndex(new BigNumber(696), new BigNumber(((_this).f6).length));
          ((_this).f6)[_index63] = (((_369_v37).contains(p0)) ? ((_369_v37).get(p0)) : (p0));
          let _index64 = _module.__default.safeIndex(new BigNumber(696), new BigNumber(((_this).f6).length));
          ((_this).f6)[_index64] = (new BigNumber(175)).isEqualTo(p1);
          let _index65 = _module.__default.safeIndex(new BigNumber(381), new BigNumber(((_this).f3).length));
          ((_this).f3)[_index65] = p1;
          let _nw50 = new _module.C0();
          _nw50.__ctor(_339_v20, ((_341_v22).f8).update(((_this).f6)[_module.__default.safeIndex(new BigNumber(696), new BigNumber(((_this).f6).length))], _333_v16), _333_v16, new BigNumber(-771));
          _341_v22 = _nw50;
          (globalState).f1 = (_dafny.ZERO).minus(new BigNumber((((((_this).f6)[_module.__default.safeIndex(new BigNumber(696), new BigNumber(((_this).f6).length))]) ? (_dafny.Seq.UnicodeFromString("i")) : (_346_v27))).length));
        }
        if (p0) {
          let _370_v38;
          _370_v38 = _module.D0.create_DC2(_346_v27, ((_this).f3)[_module.__default.safeIndex(new BigNumber(381), new BigNumber(((_this).f3).length))], p1);
          let _index66 = _module.__default.safeIndex(new BigNumber(266), new BigNumber((_348_v29).length));
          (_348_v29)[_index66] = _dafny.Seq.Concat(_dafny.Seq.Concat((_348_v29)[_module.__default.safeIndex(new BigNumber(266), new BigNumber((_348_v29).length))], (_370_v38).dtor_cf4), (_348_v29)[_module.__default.safeIndex(new BigNumber(266), new BigNumber((_348_v29).length))]);
          let _371_v39;
          let _nw51 = new _module.C0();
          _nw51.__ctor(_module.D0.create_DC0(_343_v24), ((_341_v22).f8).Merge((_341_v22).f8), _333_v16, (new BigNumber(-897)).minus(p1));
          _371_v39 = _nw51;
          let _372_v40;
          let _nw52 = new _module.C0();
          _nw52.__ctor(_module.D0.create_DC0(_343_v24), (_371_v39).f8, _333_v16, _module.__default.safeDivisionInt(p2, p1));
          _372_v40 = _nw52;
          (globalState).f1 = p1;
          let _index67 = _module.__default.safeIndex(new BigNumber(381), new BigNumber(((_this).f3).length));
          ((_this).f3)[_index67] = p1;
        } else {
          _358_v31 = ((_358_v31).Merge(_358_v31)).Merge(_358_v31);
          let _373_v41;
          _373_v41 = _dafny.Map.Empty.slice().updateUnsafe(_343_v24,_343_v24);
          let _374_v42;
          _374_v42 = _dafny.Set.fromElements((_341_v22).f7, _339_v20, (_341_v22).f7);
          let _375_v43;
          let _nw53 = Array((new BigNumber(19)).toNumber());
          _nw53[(_dafny.ZERO).toNumber()] = p0;
          _nw53[(_dafny.ONE).toNumber()] = _343_v24;
          _nw53[(new BigNumber(2)).toNumber()] = p0;
          _nw53[(new BigNumber(3)).toNumber()] = _343_v24;
          _nw53[(new BigNumber(4)).toNumber()] = p0;
          _nw53[(new BigNumber(5)).toNumber()] = _343_v24;
          _nw53[(new BigNumber(6)).toNumber()] = _343_v24;
          _nw53[(new BigNumber(7)).toNumber()] = (((_this).f3)[_module.__default.safeIndex(new BigNumber(381), new BigNumber(((_this).f3).length))]).isLessThan((_this).f4);
          _nw53[(new BigNumber(8)).toNumber()] = ((_this).f4).isLessThanOrEqualTo(p2);
          _nw53[(new BigNumber(9)).toNumber()] = p0;
          _nw53[(new BigNumber(10)).toNumber()] = !(_343_v24);
          _nw53[(new BigNumber(11)).toNumber()] = _343_v24;
          _nw53[(new BigNumber(12)).toNumber()] = (((_373_v41).contains(p0)) ? ((_373_v41).get(p0)) : (p0));
          _nw53[(new BigNumber(13)).toNumber()] = (_374_v42).IsProperSubsetOf(_module.__default.fm8(_343_v24, globalState));
          _nw53[(new BigNumber(14)).toNumber()] = p0;
          _nw53[(new BigNumber(15)).toNumber()] = _343_v24;
          _nw53[(new BigNumber(16)).toNumber()] = p0;
          _nw53[(new BigNumber(17)).toNumber()] = (((_this).f3)[_module.__default.safeIndex(new BigNumber(381), new BigNumber(((_this).f3).length))]).isLessThan((((_358_v31).contains(_module.__default.fm0(((_this).f3)[_module.__default.safeIndex(new BigNumber(381), new BigNumber(((_this).f3).length))], new BigNumber((_347_v28).length), globalState))) ? ((_358_v31).get(_module.__default.fm0(((_this).f3)[_module.__default.safeIndex(new BigNumber(381), new BigNumber(((_this).f3).length))], new BigNumber((_347_v28).length), globalState))) : (p1)));
          _nw53[(new BigNumber(18)).toNumber()] = !(true);
          _375_v43 = _nw53;
          let _init9 = ((_376_p0) => function (_377_i11) {
            return _376_p0;
          })(p0);
          let _nw54 = Array((_dafny.ONE).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw54.length); _i0_9++) {
            _nw54[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _375_v43 = _nw54;
          let _378_v44;
          _378_v44 = _dafny.Seq.of(new BigNumber(247), (_this).f4);
          let _379_v45;
          _379_v45 = _dafny.Set.fromElements(!(!(_343_v24)), true);
          let _380_v46;
          _380_v46 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm5(p0, globalState),_379_v45);
          let _381_v47;
          _381_v47 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_346_v27, _346_v27),_dafny.Seq.Concat(_378_v44, _dafny.Seq.update(_378_v44, _module.__default.safeIndex(new BigNumber((_380_v46).length), new BigNumber((_378_v44).length)), new BigNumber(32))));
          let _rhs73 = _343_v24;
          let _rhs74 = (_358_v31).update(!(p0), ((_this).f3)[_module.__default.safeIndex(new BigNumber(381), new BigNumber(((_this).f3).length))]);
          let _rhs75 = (((_381_v47).contains(_346_v27)) ? ((_381_v47).get(_346_v27)) : (_378_v44));
          _343_v24 = _rhs73;
          _358_v31 = _rhs74;
          _378_v44 = _rhs75;
          let _index68 = _module.__default.safeIndex(new BigNumber(381), new BigNumber(((_this).f3).length));
          ((_this).f3)[_index68] = ((_this).f4).plus(p2);
          let _382_v48;
          _382_v48 = _module.D1.create_DC5((p0) && (_343_v24), true, ((_this).f3)[_module.__default.safeIndex(new BigNumber(381), new BigNumber(((_this).f3).length))]);
          let _pat_let_tv3 = _382_v48;
          let _pat_let_tv4 = _343_v24;
          _382_v48 = function (_pat_let10_0) {
            return function (_383_dt__update__tmp_h0) {
              return function (_pat_let11_0) {
                return function (_384_dt__update_hcf9_h0) {
                  return function (_pat_let12_0) {
                    return function (_385_dt__update_hcf10_h0) {
                      return _module.D1.create_DC5(_384_dt__update_hcf9_h0, _385_dt__update_hcf10_h0, (_383_dt__update__tmp_h0).dtor_cf11);
                    }(_pat_let12_0);
                  }(_pat_let_tv4);
                }(_pat_let11_0);
              }((_pat_let_tv3).dtor_cf9);
            }(_pat_let10_0);
          }(_module.D1.create_DC5(p0, p0, new BigNumber(-408)));
        }
        _343_v24 = (p2).isLessThanOrEqualTo(new BigNumber(-255));
        _346_v27 = (_348_v29)[_module.__default.safeIndex(new BigNumber(266), new BigNumber((_348_v29).length))];
      } else {
        let _386_v49;
        _386_v49 = _dafny.Seq.UnicodeFromString("uaymukwrs");
        let _387_v50;
        let _nw55 = Array((new BigNumber(13)).toNumber());
        _nw55[(_dafny.ZERO).toNumber()] = (_this).f5;
        _nw55[(_dafny.ONE).toNumber()] = (_386_v49)[_module.__default.safeIndex((_this).f4, new BigNumber((_386_v49).length))];
        _nw55[(new BigNumber(2)).toNumber()] = (_this).f5;
        _nw55[(new BigNumber(3)).toNumber()] = (_this).f5;
        _nw55[(new BigNumber(4)).toNumber()] = _module.__default.fm3((_this).f5, p2, globalState);
        _nw55[(new BigNumber(5)).toNumber()] = (_this).f5;
        _nw55[(new BigNumber(6)).toNumber()] = (_this).f5;
        _nw55[(new BigNumber(7)).toNumber()] = (_this).f5;
        _nw55[(new BigNumber(8)).toNumber()] = (_this).f5;
        _nw55[(new BigNumber(9)).toNumber()] = (_this).f5;
        _nw55[(new BigNumber(10)).toNumber()] = (_this).f5;
        _nw55[(new BigNumber(11)).toNumber()] = new _dafny.CodePoint('x'.codePointAt(0));
        _nw55[(new BigNumber(12)).toNumber()] = (_this).f5;
        _387_v50 = _nw55;
        let _index69 = _module.__default.safeIndex(new BigNumber(496), new BigNumber((_387_v50).length));
        (_387_v50)[_index69] = (_this).f5;
        let _index70 = _module.__default.safeIndex(new BigNumber(496), new BigNumber((_387_v50).length));
        (_387_v50)[_index70] = ((p0) ? ((_this).f5) : ((_this).f5));
        let _388_v51;
        _388_v51 = false;
        _388_v51 = p0;
        let _389_v52;
        _389_v52 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(553),_386_v49);
        let _390_v53;
        _390_v53 = _dafny.Set.fromElements(new BigNumber((_389_v52).length));
        let _391_v54;
        _391_v54 = _dafny.Seq.of(_388_v51, _388_v51);
        let _392_v55;
        _392_v55 = _module.D0.create_DC2(_386_v49, new BigNumber((_390_v53).length), (_dafny.ZERO).minus(new BigNumber((_391_v54).length)));
        let _393_v56;
        _393_v56 = _module.D0.create_DC3(_392_v55);
        let _394_v57;
        _394_v57 = _module.D0.create_DC3(_393_v56);
        let _source6 = _394_v57;
        if (_source6.is_DC1) {
          let _395___mcc_h0 = (_source6).cf1;
          let _396___mcc_h1 = (_source6).cf2;
          let _397___mcc_h2 = (_source6).cf3;
          let _398_cf3 = _397___mcc_h2;
          let _399_cf2 = _396___mcc_h1;
          let _400_cf1 = _395___mcc_h0;
          let _401_v58;
          _401_v58 = _dafny.Seq.of(_module.D0.create_DC2(_dafny.Seq.Create(_module.__default.abs(new BigNumber(783)), function (_402_i12) {
  return (_this).f5;
}), (_this).f4, p2));
          let _403_v59;
          _403_v59 = _dafny.Seq.of(new BigNumber((_401_v58).length));
          let _404_v60;
          let _nw56 = Array((new BigNumber(10)).toNumber());
          _nw56[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_391_v54, _391_v54);
          _nw56[(_dafny.ONE).toNumber()] = _391_v54;
          _nw56[(new BigNumber(2)).toNumber()] = _391_v54;
          _nw56[(new BigNumber(3)).toNumber()] = _391_v54;
          _nw56[(new BigNumber(4)).toNumber()] = _dafny.Seq.of(p0, _388_v51, p0, true);
          _nw56[(new BigNumber(5)).toNumber()] = _391_v54;
          _nw56[(new BigNumber(6)).toNumber()] = _391_v54;
          _nw56[(new BigNumber(7)).toNumber()] = ((p0) ? (_391_v54) : (_391_v54));
          _nw56[(new BigNumber(8)).toNumber()] = _dafny.Seq.of(p0);
          _nw56[(new BigNumber(9)).toNumber()] = _391_v54;
          _404_v60 = _nw56;
          let _405_v61;
          _405_v61 = _dafny.MultiSet.fromElements(_403_v59);
          let _rhs76 = ((true) ? (_403_v59) : (_dafny.Seq.update(_403_v59, _module.__default.safeIndex(new BigNumber(-229), new BigNumber((_403_v59).length)), new BigNumber((_405_v61).cardinality()))));
          let _rhs77 = _398_cf3;
          let _rhs78 = _404_v60;
          let _rhs79 = (_dafny.ZERO).minus(p1);
          let _lhs46 = globalState;
          _403_v59 = _rhs76;
          _398_cf3 = _rhs77;
          _404_v60 = _rhs78;
          _lhs46.f1 = _rhs79;
          let _406_v62;
          _406_v62 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f3);
          let _407_v63;
          let _nw57 = new _module.C0();
          _nw57.__ctor(_module.D0.create_DC0(false), _406_v62, _333_v16, _398_cf3);
          _407_v63 = _nw57;
          let _408_v64;
          _408_v64 = _dafny.Seq.of(_407_v63);
          _386_v49 = _dafny.Seq.update(_386_v49, _module.__default.safeIndex(new BigNumber((_dafny.Seq.Concat(_408_v64, _408_v64)).length), new BigNumber((_386_v49).length)), new _dafny.CodePoint('b'.codePointAt(0)));
          _394_v57 = _394_v57;
          let _nw58 = Array((new BigNumber(13)).toNumber());
          _nw58[(_dafny.ZERO).toNumber()] = (_386_v49)[_module.__default.safeIndex(_398_cf3, new BigNumber((_386_v49).length))];
          _nw58[(_dafny.ONE).toNumber()] = (_this).f5;
          _nw58[(new BigNumber(2)).toNumber()] = (_this).f5;
          _nw58[(new BigNumber(3)).toNumber()] = (_this).f5;
          _nw58[(new BigNumber(4)).toNumber()] = (_this).f5;
          _nw58[(new BigNumber(5)).toNumber()] = (_this).f5;
          _nw58[(new BigNumber(6)).toNumber()] = (_this).f5;
          _nw58[(new BigNumber(7)).toNumber()] = (_387_v50)[_module.__default.safeIndex(new BigNumber(496), new BigNumber((_387_v50).length))];
          _nw58[(new BigNumber(8)).toNumber()] = (_387_v50)[_module.__default.safeIndex(new BigNumber(496), new BigNumber((_387_v50).length))];
          _nw58[(new BigNumber(9)).toNumber()] = (_387_v50)[_module.__default.safeIndex(new BigNumber(496), new BigNumber((_387_v50).length))];
          _nw58[(new BigNumber(10)).toNumber()] = (_387_v50)[_module.__default.safeIndex(new BigNumber(496), new BigNumber((_387_v50).length))];
          _nw58[(new BigNumber(11)).toNumber()] = new _dafny.CodePoint('b'.codePointAt(0));
          _nw58[(new BigNumber(12)).toNumber()] = (_387_v50)[_module.__default.safeIndex(new BigNumber(496), new BigNumber((_387_v50).length))];
          _387_v50 = _nw58;
        } else if (_source6.is_DC2) {
          let _409___mcc_h3 = (_source6).cf4;
          let _410___mcc_h4 = (_source6).cf5;
          let _411___mcc_h5 = (_source6).cf6;
          let _412_cf6 = _411___mcc_h5;
          let _413_cf5 = _410___mcc_h4;
          let _414_cf4 = _409___mcc_h3;
          let _415_v65;
          _415_v65 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f3);
          let _416_v66;
          let _nw59 = new _module.C0();
          _nw59.__ctor(_module.D0.create_DC0(p0), _415_v65, _333_v16, p1);
          _416_v66 = _nw59;
          let _417_v67;
          let _init10 = ((_418_v57) => function (_419_i13) {
            return ((true) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(10)), ((_420_v57) => function (_421_i14) {
              return _420_v57;
            })(_418_v57))) : (_dafny.Seq.of(_418_v57, _418_v57)));
          })(_394_v57);
          let _nw60 = Array((new BigNumber(13)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw60.length); _i0_10++) {
            _nw60[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _417_v67 = _nw60;
          let _422_v68;
          _422_v68 = _module.D1.create_DC5(_388_v51, _388_v51, p1);
          let _rhs80 = p1;
          let _rhs81 = _417_v67;
          let _rhs82 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(p1), new BigNumber((_dafny.Seq.Concat(_414_cf4, _dafny.Seq.update(_414_cf4, _module.__default.safeIndex(p2, new BigNumber((_414_cf4).length)), (_387_v50)[_module.__default.safeIndex(new BigNumber(496), new BigNumber((_387_v50).length))]))).length));
          let _rhs83 = ((_422_v68).dtor_cf11).isLessThanOrEqualTo(_module.__default.safeModuloInt(p1, new BigNumber(856)));
          let _lhs47 = globalState;
          let _lhs48 = globalState;
          _lhs47.f1 = _rhs80;
          _417_v67 = _rhs81;
          _lhs48.f1 = _rhs82;
          _388_v51 = _rhs83;
          let _423_v69;
          _423_v69 = _dafny.Seq.of(_module.__default.fm5(_388_v51, globalState));
          let _424_v70;
          _424_v70 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm5(true, globalState),new BigNumber(-217));
          let _425_v71;
          let _nw61 = Array((new BigNumber(19)).toNumber());
          _nw61[(_dafny.ZERO).toNumber()] = _388_v51;
          _nw61[(_dafny.ONE).toNumber()] = ((_this).f4).isLessThan(_412_cf6);
          _nw61[(new BigNumber(2)).toNumber()] = _388_v51;
          _nw61[(new BigNumber(3)).toNumber()] = _dafny.Seq.IsPrefixOf(_423_v69, _423_v69);
          _nw61[(new BigNumber(4)).toNumber()] = (p1).isLessThanOrEqualTo(new BigNumber((_390_v53).length));
          _nw61[(new BigNumber(5)).toNumber()] = p0;
          _nw61[(new BigNumber(6)).toNumber()] = (_390_v53).IsProperSubsetOf(_390_v53);
          _nw61[(new BigNumber(7)).toNumber()] = p0;
          _nw61[(new BigNumber(8)).toNumber()] = (_388_v51) === (_388_v51);
          _nw61[(new BigNumber(9)).toNumber()] = p0;
          _nw61[(new BigNumber(10)).toNumber()] = p0;
          _nw61[(new BigNumber(11)).toNumber()] = _388_v51;
          _nw61[(new BigNumber(12)).toNumber()] = (_391_v54)[_module.__default.safeIndex(_412_cf6, new BigNumber((_391_v54).length))];
          _nw61[(new BigNumber(13)).toNumber()] = true;
          _nw61[(new BigNumber(14)).toNumber()] = _388_v51;
          _nw61[(new BigNumber(15)).toNumber()] = !((((_424_v70).contains((_dafny.ZERO).minus(_412_cf6))) ? ((_424_v70).get((_dafny.ZERO).minus(_412_cf6))) : ((_dafny.ZERO).minus((_this).f4)))).isEqualTo((_this).f4);
          _nw61[(new BigNumber(16)).toNumber()] = _dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(752)), ((_426_v50) => function (_427_i15) {
            return (_426_v50)[_module.__default.safeIndex(new BigNumber(496), new BigNumber((_426_v50).length))];
          })(_387_v50)), _386_v49);
          _nw61[(new BigNumber(17)).toNumber()] = !(_388_v51);
          _nw61[(new BigNumber(18)).toNumber()] = p0;
          _425_v71 = _nw61;
          _425_v71 = _425_v71;
          (globalState).f1 = p1;
        } else if (_source6.is_DC0) {
          let _428___mcc_h6 = (_source6).cf0;
          let _429_cf0 = _428___mcc_h6;
          let _430_v72;
          _430_v72 = _module.D1.create_DC4(_333_v16);
          let _431_v73;
          _431_v73 = _dafny.Map.Empty.slice().updateUnsafe(!(_429_cf0),_module.__default.fm5(true, globalState));
          let _rhs84 = (((p0) ? ((_dafny.ZERO).minus(new BigNumber((_431_v73).length))) : ((_this).f4))).plus((_this).f4);
          let _rhs85 = _430_v72;
          let _lhs49 = globalState;
          _lhs49.f1 = _rhs84;
          _430_v72 = _rhs85;
          let _432_v74;
          _432_v74 = _module.D0.create_DC0(_388_v51);
          let _433_v75;
          _433_v75 = _dafny.Map.Empty.slice().updateUnsafe(_429_cf0,(_this).f3);
          let _434_v76;
          let _nw62 = new _module.C0();
          _nw62.__ctor(_432_v74, _433_v75, (_this).f3, p2);
          _434_v76 = _nw62;
          (globalState).f1 = (_this).f4;
          let _rhs86 = p1;
          let _rhs87 = _module.__default.safeModuloInt(p1, p1);
          let _lhs50 = globalState;
          let _lhs51 = globalState;
          _lhs50.f1 = _rhs86;
          _lhs51.f1 = _rhs87;
        } else {
          let _435___mcc_h7 = (_source6).cf7;
          let _436_cf7 = _435___mcc_h7;
          (globalState).f1 = p2;
          let _437_v77;
          _437_v77 = _dafny.Map.Empty.slice().updateUnsafe((_this).f4,(_this).f4);
          let _438_v78;
          _438_v78 = _dafny.Map.Empty.slice().updateUnsafe((_388_v51) && (false),(_386_v49)[_module.__default.safeIndex(new BigNumber((_437_v77).length), new BigNumber((_386_v49).length))]);
          _438_v78 = (_438_v78).update(_388_v51, (_this).f5);
          let _439_v79;
          _439_v79 = _dafny.Map.Empty.slice().updateUnsafe((_this).f3,p1);
          _439_v79 = (_439_v79).update(_333_v16, p2);
          let _440_v80;
          let _nw63 = Array((new BigNumber(18)).toNumber()).fill(false);
          _440_v80 = _nw63;
          _440_v80 = (_this).f6;
        }
        let _441_v81;
        _441_v81 = _dafny.Seq.of(_387_v50, _387_v50);
        let _rhs88 = _dafny.Seq.IsProperPrefixOf(_441_v81, _dafny.Seq.Concat(_441_v81, _dafny.Seq.of((_441_v81)[_module.__default.safeIndex(new BigNumber((_391_v54).length), new BigNumber((_441_v81).length))], _387_v50, _387_v50)));
        let _rhs89 = (_this).f4;
        let _rhs90 = !(!_dafny.areEqual(_386_v49, _386_v49));
        let _lhs52 = globalState;
        _388_v51 = _rhs88;
        _lhs52.f1 = _rhs89;
        _388_v51 = _rhs90;
        _388_v51 = _388_v51;
      }
      let _442_v82;
      _442_v82 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_dafny.ZERO).minus(p2));
      let _443_v83;
      _443_v83 = _module.D0.create_DC1((_this).f4, (_dafny.ZERO).minus(p2), new BigNumber((_442_v82).length));
      let _444_v84;
      _444_v84 = _dafny.Seq.UnicodeFromString("uohkfjswt");
      let _hi2 = _module.__default.safeModuloInt((_this).f4, _module.__default.fm5(p0, globalState));
      for (let _445_i16 = _module.__default.safeModuloInt((_443_v83).dtor_cf2, new BigNumber((_444_v84).length)); _445_i16.isLessThan(_hi2); _445_i16 = _445_i16.plus(_dafny.ONE)) {
        let _446_v85;
        _446_v85 = _dafny.Seq.of(p1);
        let _447_v86;
        let _nw64 = Array((new BigNumber(22)).toNumber());
        _nw64[(_dafny.ZERO).toNumber()] = _module.__default.fm3((_this).f5, (_this).f4, globalState);
        _nw64[(_dafny.ONE).toNumber()] = new _dafny.CodePoint('f'.codePointAt(0));
        _nw64[(new BigNumber(2)).toNumber()] = (_this).f5;
        _nw64[(new BigNumber(3)).toNumber()] = (_this).f5;
        _nw64[(new BigNumber(4)).toNumber()] = _module.__default.fm3((_this).f5, p1, globalState);
        _nw64[(new BigNumber(5)).toNumber()] = (_this).f5;
        _nw64[(new BigNumber(6)).toNumber()] = (_this).f5;
        _nw64[(new BigNumber(7)).toNumber()] = (_this).f5;
        _nw64[(new BigNumber(8)).toNumber()] = new _dafny.CodePoint('a'.codePointAt(0));
        _nw64[(new BigNumber(9)).toNumber()] = (_this).f5;
        _nw64[(new BigNumber(10)).toNumber()] = (_this).f5;
        _nw64[(new BigNumber(11)).toNumber()] = (_this).f5;
        _nw64[(new BigNumber(12)).toNumber()] = (_this).f5;
        _nw64[(new BigNumber(13)).toNumber()] = (_this).f5;
        _nw64[(new BigNumber(14)).toNumber()] = new _dafny.CodePoint('j'.codePointAt(0));
        _nw64[(new BigNumber(15)).toNumber()] = (_this).f5;
        _nw64[(new BigNumber(16)).toNumber()] = (_this).f5;
        _nw64[(new BigNumber(17)).toNumber()] = (_this).f5;
        _nw64[(new BigNumber(18)).toNumber()] = (_444_v84)[_module.__default.safeIndex(new BigNumber((_446_v85).length), new BigNumber((_444_v84).length))];
        _nw64[(new BigNumber(19)).toNumber()] = (_this).f5;
        _nw64[(new BigNumber(20)).toNumber()] = (_this).f5;
        _nw64[(new BigNumber(21)).toNumber()] = (_this).f5;
        _447_v86 = _nw64;
        let _448_v87;
        _448_v87 = _dafny.Seq.of(_447_v86, _447_v86);
        let _449_v88;
        _449_v88 = _dafny.Set.fromElements((_this).f5);
        let _source7 = _module.D0.create_DC2(_444_v84, new BigNumber((_448_v87).length), _module.__default.safeDivisionInt(p2, new BigNumber((_449_v88).length)));
        if (_source7.is_DC1) {
          let _450___mcc_h8 = (_source7).cf1;
          let _451___mcc_h9 = (_source7).cf2;
          let _452___mcc_h10 = (_source7).cf3;
          let _453_cf3 = _452___mcc_h10;
          let _454_cf2 = _451___mcc_h9;
          let _455_cf1 = _450___mcc_h8;
          let _456_v89;
          _456_v89 = false;
          _456_v89 = !(!(!(_456_v89)));
          let _457_v90;
          _457_v90 = _dafny.Set.fromElements(_456_v89, _456_v89, true);
          _457_v90 = (_457_v90).Union(_457_v90);
          let _458_v91;
          _458_v91 = _module.D0.create_DC0(p0);
          let _459_v92;
          _459_v92 = _dafny.Map.Empty.slice().updateUnsafe(_456_v89,_333_v16);
          let _460_v93;
          let _nw65 = new _module.C0();
          _nw65.__ctor(_458_v91, _459_v92, _333_v16, _453_cf3);
          _460_v93 = _nw65;
          let _461_v94;
          let _nw66 = Array((new BigNumber(21)).toNumber());
          _nw66[(_dafny.ZERO).toNumber()] = _460_v93;
          _nw66[(_dafny.ONE).toNumber()] = _460_v93;
          _nw66[(new BigNumber(2)).toNumber()] = _460_v93;
          _nw66[(new BigNumber(3)).toNumber()] = _460_v93;
          _nw66[(new BigNumber(4)).toNumber()] = _460_v93;
          _nw66[(new BigNumber(5)).toNumber()] = _460_v93;
          _nw66[(new BigNumber(6)).toNumber()] = _460_v93;
          _nw66[(new BigNumber(7)).toNumber()] = _460_v93;
          _nw66[(new BigNumber(8)).toNumber()] = _460_v93;
          _nw66[(new BigNumber(9)).toNumber()] = _460_v93;
          _nw66[(new BigNumber(10)).toNumber()] = _460_v93;
          _nw66[(new BigNumber(11)).toNumber()] = _460_v93;
          _nw66[(new BigNumber(12)).toNumber()] = _460_v93;
          _nw66[(new BigNumber(13)).toNumber()] = _460_v93;
          _nw66[(new BigNumber(14)).toNumber()] = _460_v93;
          _nw66[(new BigNumber(15)).toNumber()] = _460_v93;
          _nw66[(new BigNumber(16)).toNumber()] = _460_v93;
          _nw66[(new BigNumber(17)).toNumber()] = _460_v93;
          _nw66[(new BigNumber(18)).toNumber()] = _460_v93;
          _nw66[(new BigNumber(19)).toNumber()] = _460_v93;
          _nw66[(new BigNumber(20)).toNumber()] = _460_v93;
          _461_v94 = _nw66;
          let _index71 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_461_v94).length));
          (_461_v94)[_index71] = _460_v93;
          let _index72 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_461_v94).length));
          (_461_v94)[_index72] = _460_v93;
          let _462_v95;
          _462_v95 = _dafny.Map.Empty.slice().updateUnsafe(_456_v89,_456_v89);
          let _463_v96;
          let _nw67 = new _module.C0();
          _nw67.__ctor(_458_v91, _459_v92, (_this).f3, new BigNumber((_462_v95).length));
          _463_v96 = _nw67;
          let _464_v97;
          _464_v97 = _dafny.Map.Empty.slice().updateUnsafe(_463_v96,(_460_v93).f3);
          let _465_v98;
          let _nw68 = new _module.C0();
          _nw68.__ctor(_module.D0.create_DC0(p0), _459_v92, (((_464_v97).contains(_463_v96)) ? ((_464_v97).get(_463_v96)) : (_333_v16)), new BigNumber(-693));
          _465_v98 = _nw68;
        } else if (_source7.is_DC2) {
          let _466___mcc_h11 = (_source7).cf4;
          let _467___mcc_h12 = (_source7).cf5;
          let _468___mcc_h13 = (_source7).cf6;
          let _469_cf6 = _468___mcc_h13;
          let _470_cf5 = _467___mcc_h12;
          let _471_cf4 = _466___mcc_h11;
          let _472_v99;
          let _init11 = function (_473_i17) {
            return _dafny.MultiSet.FromArray((_module.D2.create_DC7(_dafny.Seq.of(true))).dtor_cf12);
          };
          let _nw69 = Array((new BigNumber(27)).toNumber());
          for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw69.length); _i0_11++) {
            _nw69[_i0_11] = _init11(new BigNumber(_i0_11));
          }
          _472_v99 = _nw69;
          let _474_v100;
          _474_v100 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_module.__default.safeModuloInt(_469_cf6, _470_cf5)),_472_v99);
          _474_v100 = (_474_v100).update(_469_cf6, _472_v99);
          let _475_v101;
          _475_v101 = _dafny.MultiSet.fromElements(((!(p0)) ? (new BigNumber((_471_cf4).length)) : (_470_cf5)), _module.__default.safeModuloInt(_445_i16, _470_cf5), _469_cf6, (_470_cf5).plus(p2));
          (globalState).f1 = new BigNumber((_475_v101).cardinality());
          let _476_v102;
          _476_v102 = true;
          let _477_v103;
          _477_v103 = _dafny.Map.Empty.slice().updateUnsafe((_this).f5,_476_v102);
          _476_v102 = !(!(!((_477_v103).Merge(_477_v103)).equals((_477_v103).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),p0)))));
          let _478_v104;
          _478_v104 = _module.D1.create_DC4((_this).f3);
          let _479_v105;
          _479_v105 = _dafny.Seq.of(_476_v102);
          let _480_v106;
          _480_v106 = _module.D1.create_DC5(false, p0, new BigNumber((_479_v105).length));
          _478_v104 = _module.D1.create_DC4((((_480_v106).dtor_cf10) ? (_333_v16) : (_333_v16)));
        } else if (_source7.is_DC0) {
          let _481___mcc_h14 = (_source7).cf0;
          let _482_cf0 = _481___mcc_h14;
          _482_cf0 = true;
          (globalState).f1 = _445_i16;
          _444_v84 = _444_v84;
          let _483_v107;
          _483_v107 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(993)), function (_484_i18) {
            return (_this).f5;
          }),_445_i16);
          _483_v107 = (_483_v107).update(_444_v84, _module.__default.fm5(p0, globalState));
        } else {
          let _485___mcc_h15 = (_source7).cf7;
          let _486_cf7 = _485___mcc_h15;
          let _487_v108;
          let _nw70 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.of());
          _487_v108 = _nw70;
          let _488_v109;
          _488_v109 = _module.D1.create_DC4(_333_v16);
          let _489_v110;
          _489_v110 = _dafny.Seq.of((_488_v109).dtor_cf8);
          let _index73 = _module.__default.safeIndex(new BigNumber(211), new BigNumber((_487_v108).length));
          (_487_v108)[_index73] = _489_v110;
          let _490_v111;
          _490_v111 = false;
          let _491_v112;
          _491_v112 = _dafny.Seq.of((_this).f6, (_this).f6, (_this).f6, (_this).f6, (_this).f6);
          let _index74 = _module.__default.safeIndex(new BigNumber(211), new BigNumber((_487_v108).length));
          let _rhs91 = new BigNumber(957);
          let _rhs92 = new BigNumber((_dafny.Seq.update(_491_v112, _module.__default.safeIndex(p2, new BigNumber((_491_v112).length)), (_this).f6)).length);
          let _rhs93 = _489_v110;
          let _rhs94 = !(((!(_490_v111) || (_490_v111)) ? (_module.__default.fm0(new BigNumber((_446_v85).length), _445_i16, globalState)) : (_490_v111)));
          let _lhs53 = globalState;
          let _lhs54 = globalState;
          let _lhs55 = _487_v108;
          let _lhs56 = _module.__default.safeIndex(new BigNumber(211), new BigNumber((_487_v108).length));
          _lhs53.f1 = _rhs91;
          _lhs54.f1 = _rhs92;
          _lhs55[_lhs56] = _rhs93;
          _490_v111 = _rhs94;
          _490_v111 = (_490_v111) === ((_module.D1.create_DC5(_490_v111, _490_v111, new BigNumber((_442_v82).length))).dtor_cf9);
          _490_v111 = _490_v111;
          let _492_v113;
          _492_v113 = new _dafny.CodePoint('j'.codePointAt(0));
          _492_v113 = (_this).f5;
        }
        let _493_v114;
        _493_v114 = _module.D0.create_DC0(p0);
        let _494_v115;
        _494_v115 = _dafny.Map.Empty.slice().updateUnsafe(p0,_333_v16);
        let _495_v116;
        _495_v116 = _dafny.Seq.of(p0, p0);
        let _pat_let_tv5 = p0;
        let _496_v117;
        let _nw71 = new _module.C0();
        _nw71.__ctor(function (_pat_let13_0) {
          return function (_497_dt__update__tmp_h1) {
            return function (_pat_let14_0) {
              return function (_498_dt__update_hcf0_h0) {
                return _module.D0.create_DC0(_498_dt__update_hcf0_h0);
              }(_pat_let14_0);
            }(_pat_let_tv5);
          }(_pat_let13_0);
        }(_493_v114), ((_494_v115).update(p0, _333_v16)).Merge((_494_v115).update((_495_v116)[_module.__default.safeIndex(_module.__default.fm5(p0, globalState), new BigNumber((_495_v116).length))], (_this).f3)), (_this).f3, new BigNumber((_444_v84).length));
        _496_v117 = _nw71;
        let _499_v118;
        _499_v118 = true;
        let _500_v119;
        _500_v119 = _dafny.Seq.of(_444_v84);
        _499_v118 = _dafny.areEqual((_500_v119)[_module.__default.safeIndex(new BigNumber(198), new BigNumber((_500_v119).length))], _dafny.Seq.Concat(_444_v84, _dafny.Seq.UnicodeFromString("hjsywnkxd")));
        let _501_v120;
        _501_v120 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('c'.codePointAt(0)),_445_i16);
        let _502_v121;
        _502_v121 = _dafny.MultiSet.fromElements(_501_v120, _module.__default.fm9(globalState), _501_v120);
        _502_v121 = _dafny.MultiSet.FromArray(_dafny.Seq.update(_module.__default.fm10(_module.D1.create_DC6(), (_444_v84)[_module.__default.safeIndex(_445_i16, new BigNumber((_444_v84).length))], (_dafny.ZERO).minus(_445_i16), !(_499_v118), globalState), _module.__default.safeIndex(p2, new BigNumber((_module.__default.fm10(_module.D1.create_DC6(), (_444_v84)[_module.__default.safeIndex(_445_i16, new BigNumber((_444_v84).length))], (_dafny.ZERO).minus(_445_i16), !(_499_v118), globalState)).length)), _501_v120));
      }
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_333_v16).length))) {
        let _503_i19 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_503_i19)) && ((_503_i19).isLessThan(new BigNumber((_333_v16).length))))) {
          (_333_v16)[(_503_i19)] = _module.__default.safeModuloInt(_503_i19, ((p0) ? ((_this).f4) : (p2)));
        }
      }
      let _504_v122;
      _504_v122 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(870)), function (_505_i20) {
        return (_this).f5;
      }));
      let _506_v123;
      _506_v123 = _module.D0.create_DC2(_444_v84, (_this).f4, (_dafny.ZERO).minus(p2));
      r0 = _dafny.Seq.Concat((_504_v122)[_module.__default.safeIndex((_this).f4, new BigNumber((_504_v122).length))], (_506_v123).dtor_cf4);
      return r0;
    }
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
