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
      return (_dafny.Set.fromElements(true, true)).Union(_dafny.Set.fromElements(true));
    };
    static fm1(p0, p1, p2, globalState) {
      return true;
    };
    static fm2(p0, globalState) {
      return _module.D0.create_DC1();
    };
    static fm3(globalState) {
      let _source0 = _module.D0.create_DC0(true);
      if (_source0.is_DC1) {
        return _dafny.Seq.UnicodeFromString("ttioiues");
      } else {
        let _0___mcc_h0 = (_source0).cf0;
        let _1_cf0 = _0___mcc_h0;
        return _dafny.Seq.UnicodeFromString("p");
      }
    };
    static fm4(p0, p1, p2, p3, globalState) {
      return new BigNumber(-413);
    };
    static fm6(p0, p1, p2, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(857)), function (_2_i0) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("meafobh")).length);
      });
    };
    static fm7(p0, globalState) {
      let _source1 = _module.D0.create_DC1();
      if (_source1.is_DC1) {
        return _module.D0.create_DC0(true);
      } else {
        let _3___mcc_h0 = (_source1).cf0;
        let _4_cf0 = _3___mcc_h0;
        return _module.D0.create_DC0(_4_cf0);
      }
    };
    static m0(p0, p1, globalState) {
      let r0 = _dafny.Seq.of();
      let r1 = _dafny.ZERO;
      let r2 = false;
      let _5_v0;
      let _nw0 = Array((new BigNumber(7)).toNumber()).fill(false);
      _5_v0 = _nw0;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_5_v0).length))) {
        let _6_i0 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_6_i0)) && ((_6_i0).isLessThan(new BigNumber((_5_v0).length))))) {
          (_5_v0)[(_6_i0)] = ((!(new BigNumber((_dafny.Seq.UnicodeFromString("icxytex")).length)).isEqualTo(new BigNumber(410))) ? (p0) : (p0));
        }
      }
      let _7_v1;
      _7_v1 = new BigNumber(14);
      let _8_v2;
      _8_v2 = _dafny.Seq.of(_7_v1);
      let _9_v3;
      let _nw1 = new _module.C0();
      _nw1.__ctor(_8_v2);
      _9_v3 = _nw1;
      let _10_v4;
      _10_v4 = _dafny.Map.Empty.slice().updateUnsafe(_9_v3,p1);
      let _11_v5;
      _11_v5 = _dafny.Map.Empty.slice().updateUnsafe(_7_v1,_10_v4);
      let _12_v6;
      _12_v6 = _dafny.Seq.of(_11_v5);
      let _13_v7;
      _13_v7 = _module.D1.create_DC4((_12_v6)[_module.__default.safeIndex(_7_v1, new BigNumber((_12_v6).length))], _7_v1);
      let _source2 = _module.__default.fm2((_13_v7).dtor_cf5, globalState);
      if (_source2.is_DC1) {
        r2 = false;
        let _14_v8;
        _14_v8 = _dafny.Seq.of(p0, p0);
        let _15_v9;
        _15_v9 = _dafny.Map.Empty.slice().updateUnsafe((_14_v8)[_module.__default.safeIndex(_7_v1, new BigNumber((_14_v8).length))],_7_v1);
        let _16_v10;
        let _nw2 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
        _16_v10 = _nw2;
        let _17_v11;
        _17_v11 = _dafny.MultiSet.fromElements(_7_v1, _module.__default.fm4(_7_v1, _7_v1, p0, p0, globalState), new BigNumber((_15_v9).length), new BigNumber((_dafny.Seq.of(_16_v10, _16_v10, _16_v10)).length));
        let _18_v12;
        _18_v12 = _module.D0.create_DC0((_dafny.MultiSet.fromElements(_7_v1)).equals(_17_v11));
        let _19_v13;
        _19_v13 = _dafny.Seq.UnicodeFromString("avymisae");
        let _20_v14;
        _20_v14 = _dafny.Map.Empty.slice().updateUnsafe(p1,_19_v13);
        let _21_v15;
        _21_v15 = _dafny.Map.Empty.slice().updateUnsafe(!(p1),(((_20_v14).contains(p0)) ? ((_20_v14).get(p0)) : (_19_v13)));
        let _22_v16;
        let _nw3 = Array((new BigNumber(16)).toNumber());
        _nw3[(_dafny.ZERO).toNumber()] = _19_v13;
        _nw3[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_19_v13, _19_v13);
        _nw3[(new BigNumber(2)).toNumber()] = _19_v13;
        _nw3[(new BigNumber(3)).toNumber()] = _19_v13;
        _nw3[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_19_v13, _19_v13);
        _nw3[(new BigNumber(5)).toNumber()] = (((_21_v15).contains(p1)) ? ((_21_v15).get(p1)) : (_dafny.Seq.update(_19_v13, _module.__default.safeIndex(_7_v1, new BigNumber((_19_v13).length)), new _dafny.CodePoint('w'.codePointAt(0)))));
        _nw3[(new BigNumber(6)).toNumber()] = _19_v13;
        _nw3[(new BigNumber(7)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(490)), function (_23_i1) {
          return new _dafny.CodePoint('m'.codePointAt(0));
        });
        _nw3[(new BigNumber(8)).toNumber()] = _19_v13;
        _nw3[(new BigNumber(9)).toNumber()] = ((p0) ? (_19_v13) : (_19_v13));
        _nw3[(new BigNumber(10)).toNumber()] = _19_v13;
        _nw3[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_19_v13, _19_v13);
        _nw3[(new BigNumber(12)).toNumber()] = _dafny.Seq.UnicodeFromString("rwsgvn");
        _nw3[(new BigNumber(13)).toNumber()] = _module.__default.fm3(globalState);
        _nw3[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_19_v13, _19_v13);
        _nw3[(new BigNumber(15)).toNumber()] = _19_v13;
        _22_v16 = _nw3;
        let _24_v17;
        _24_v17 = new _dafny.CodePoint('v'.codePointAt(0));
        let _25_v18;
        let _nw4 = Array((new BigNumber(4)).toNumber());
        _nw4[(_dafny.ZERO).toNumber()] = _24_v17;
        _nw4[(_dafny.ONE).toNumber()] = _24_v17;
        _nw4[(new BigNumber(2)).toNumber()] = _24_v17;
        _nw4[(new BigNumber(3)).toNumber()] = _24_v17;
        _25_v18 = _nw4;
        let _26_v19;
        _26_v19 = _dafny.Seq.of(_9_v3);
        let _27_v20;
        let _nw5 = Array((new BigNumber(10)).toNumber());
        _nw5[(_dafny.ZERO).toNumber()] = (_26_v19)[_module.__default.safeIndex(_7_v1, new BigNumber((_26_v19).length))];
        _nw5[(_dafny.ONE).toNumber()] = _9_v3;
        _nw5[(new BigNumber(2)).toNumber()] = _9_v3;
        _nw5[(new BigNumber(3)).toNumber()] = _9_v3;
        _nw5[(new BigNumber(4)).toNumber()] = _9_v3;
        _nw5[(new BigNumber(5)).toNumber()] = _9_v3;
        _nw5[(new BigNumber(6)).toNumber()] = _9_v3;
        _nw5[(new BigNumber(7)).toNumber()] = _9_v3;
        _nw5[(new BigNumber(8)).toNumber()] = _9_v3;
        _nw5[(new BigNumber(9)).toNumber()] = _9_v3;
        _27_v20 = _nw5;
        let _28_v21;
        _28_v21 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_14_v8).length),_27_v20);
        let _29_v22;
        _29_v22 = _dafny.Map.Empty.slice().updateUnsafe(_5_v0,(((_28_v21).contains(_7_v1)) ? ((_28_v21).get(_7_v1)) : (_27_v20)));
        let _30_v23;
        _30_v23 = _dafny.Map.Empty.slice().updateUnsafe(_25_v18,_29_v22);
        let _31_v24;
        _31_v24 = _dafny.Seq.of(_29_v22);
        let _32_v25;
        let _nw6 = Array((new BigNumber(23)).toNumber());
        _nw6[(_dafny.ZERO).toNumber()] = (((_30_v23).contains(_25_v18)) ? ((_30_v23).get(_25_v18)) : (_dafny.Map.Empty.slice().updateUnsafe(_5_v0,_27_v20)));
        _nw6[(_dafny.ONE).toNumber()] = _29_v22;
        _nw6[(new BigNumber(2)).toNumber()] = (_29_v22).Merge(_29_v22);
        _nw6[(new BigNumber(3)).toNumber()] = _29_v22;
        _nw6[(new BigNumber(4)).toNumber()] = _29_v22;
        _nw6[(new BigNumber(5)).toNumber()] = _29_v22;
        _nw6[(new BigNumber(6)).toNumber()] = _29_v22;
        _nw6[(new BigNumber(7)).toNumber()] = _29_v22;
        _nw6[(new BigNumber(8)).toNumber()] = (_29_v22).Merge(_dafny.Map.Empty.slice().updateUnsafe(_5_v0,_27_v20));
        _nw6[(new BigNumber(9)).toNumber()] = (_module.D3.create_DC6(_dafny.Map.Empty.slice().updateUnsafe(_5_v0,_27_v20))).dtor_cf7;
        _nw6[(new BigNumber(10)).toNumber()] = _29_v22;
        _nw6[(new BigNumber(11)).toNumber()] = _29_v22;
        _nw6[(new BigNumber(12)).toNumber()] = _29_v22;
        _nw6[(new BigNumber(13)).toNumber()] = _29_v22;
        _nw6[(new BigNumber(14)).toNumber()] = _29_v22;
        _nw6[(new BigNumber(15)).toNumber()] = (_29_v22).Merge(_dafny.Map.Empty.slice().updateUnsafe(_5_v0,_27_v20));
        _nw6[(new BigNumber(16)).toNumber()] = _29_v22;
        _nw6[(new BigNumber(17)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_5_v0,_27_v20);
        _nw6[(new BigNumber(18)).toNumber()] = (_31_v24)[_module.__default.safeIndex(_7_v1, new BigNumber((_31_v24).length))];
        _nw6[(new BigNumber(19)).toNumber()] = (_29_v22).Merge(_29_v22);
        _nw6[(new BigNumber(20)).toNumber()] = (_29_v22).Merge(_29_v22);
        _nw6[(new BigNumber(21)).toNumber()] = _29_v22;
        _nw6[(new BigNumber(22)).toNumber()] = (_29_v22).update(_5_v0, _27_v20);
        _32_v25 = _nw6;
        let _index0 = _module.__default.safeIndex(new BigNumber(548), new BigNumber((_32_v25).length));
        (_32_v25)[_index0] = _29_v22;
        let _index1 = _module.__default.safeIndex(new BigNumber(548), new BigNumber((_32_v25).length));
        let _rhs0 = p0;
        let _rhs1 = _18_v12;
        let _rhs2 = _22_v16;
        let _rhs3 = (_29_v22).Merge(_29_v22);
        let _lhs0 = _32_v25;
        let _lhs1 = _module.__default.safeIndex(new BigNumber(548), new BigNumber((_32_v25).length));
        r2 = _rhs0;
        _18_v12 = _rhs1;
        _22_v16 = _rhs2;
        _lhs0[_lhs1] = _rhs3;
        let _33_v26;
        let _nw7 = Array((new BigNumber(25)).toNumber()).fill([]);
        _33_v26 = _nw7;
        let _index2 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_33_v26).length));
        (_33_v26)[_index2] = _16_v10;
        let _34_v27;
        _34_v27 = _dafny.Map.Empty.slice().updateUnsafe(p1,_17_v11);
        let _35_v28;
        _35_v28 = _dafny.MultiSet.fromElements(p1, p0);
        let _36_v29;
        _36_v29 = _dafny.Seq.of(_16_v10);
        let _37_v30;
        _37_v30 = _dafny.Map.Empty.slice().updateUnsafe((_17_v11).IsProperSubsetOf((((_34_v27).contains(p0)) ? ((_34_v27).get(p0)) : (_dafny.MultiSet.fromElements(_7_v1, new BigNumber((_35_v28).cardinality()), new BigNumber(57), _7_v1)))),(_36_v29)[_module.__default.safeIndex(_7_v1, new BigNumber((_36_v29).length))]);
        let _38_v31;
        _38_v31 = _16_v10;
        let _index3 = _module.__default.safeIndex(new BigNumber(132), new BigNumber((_33_v26).length));
        (_33_v26)[_index3] = (((_37_v30).contains(p1)) ? ((_37_v30).get(p1)) : ((_38_v31)));
        r2 = !(!(!(p0)));
      } else {
        let _39___mcc_h0 = (_source2).cf0;
        let _40_cf0 = _39___mcc_h0;
        let _41_v32;
        _41_v32 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-50)), ((_42_v1) => function (_43_i2) {
          return _42_v1;
        })(_7_v1));
        _8_v2 = _dafny.Seq.Concat((_9_v3).f0, (_41_v32));
        let _44_v33;
        _44_v33 = new _dafny.CodePoint('y'.codePointAt(0));
        _44_v33 = _44_v33;
        r1 = (_7_v1).minus(new BigNumber(-38));
        let _45_v34;
        _45_v34 = _dafny.Set.fromElements(_13_v7);
        let _46_v35;
        _46_v35 = _dafny.Map.Empty.slice().updateUnsafe(_9_v3,(_45_v34).Intersect(_45_v34));
        let _47_v36;
        _47_v36 = _dafny.Seq.of(_45_v34, _45_v34, _45_v34, _45_v34, _45_v34);
        _46_v35 = (_46_v35).update(_9_v3, (_47_v36)[_module.__default.safeIndex(_7_v1, new BigNumber((_47_v36).length))]);
      }
      let _index4 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_5_v0).length));
      (_5_v0)[_index4] = p0;
      let _48_v37;
      _48_v37 = _dafny.Seq.UnicodeFromString("mexshnevu");
      let _49_v38;
      _49_v38 = _dafny.MultiSet.fromElements(_module.__default.fm1(p1, p0, new BigNumber((_dafny.Seq.of(_8_v2)).length), globalState), p0);
      let _50_v39;
      _50_v39 = _dafny.Seq.of(p0, p0);
      let _51_v40;
      _51_v40 = _dafny.MultiSet.fromElements(_9_v3);
      let _index5 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_5_v0).length));
      let _rhs4 = p0;
      let _rhs5 = !_dafny.Seq.contains(_50_v39, (_49_v38).IsSubsetOf(_49_v38));
      let _rhs6 = (_51_v40).IsProperSubsetOf(_51_v40);
      let _rhs7 = _48_v37;
      let _lhs2 = _5_v0;
      let _lhs3 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_5_v0).length));
      r2 = _rhs4;
      r2 = _rhs5;
      _lhs2[_lhs3] = _rhs6;
      _48_v37 = _rhs7;
      let _52_v41;
      let _nw8 = new _module.C0();
      _nw8.__ctor(_dafny.Seq.Concat(_8_v2, (_9_v3).f0));
      _52_v41 = _nw8;
      let _53_v42;
      let _init0 = ((_54_v0) => function (_55_i3) {
        return ((false) ? (_dafny.Map.Empty.slice().updateUnsafe((_54_v0)[_module.__default.safeIndex(new BigNumber(314), new BigNumber((_54_v0).length))],_module.D0.create_DC1())) : (_dafny.Map.Empty.slice().updateUnsafe((_54_v0)[_module.__default.safeIndex(new BigNumber(314), new BigNumber((_54_v0).length))],_module.D0.create_DC1())));
      })(_5_v0);
      let _nw9 = Array((new BigNumber(9)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw9.length); _i0_0++) {
        _nw9[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _53_v42 = _nw9;
      let _56_v43;
      _56_v43 = _module.D0.create_DC1();
      let _57_v44;
      _57_v44 = _dafny.Map.Empty.slice().updateUnsafe(p0,_56_v43);
      let _index6 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((_53_v42).length));
      (_53_v42)[_index6] = _57_v44;
      let _pat_let_tv0 = _57_v44;
      let _pat_let_tv1 = _57_v44;
      let _index7 = _module.__default.safeIndex(new BigNumber(646), new BigNumber((_53_v42).length));
      (_53_v42)[_index7] = function (_source3) {
        if (_source3.is_DC1) {
          return _pat_let_tv0;
        } else {
          let _58___mcc_h1 = (_source3).cf0;
          let _59_cf0 = _58___mcc_h1;
          return _pat_let_tv1;
        }
      }(_module.__default.fm7((_5_v0)[_module.__default.safeIndex(new BigNumber(314), new BigNumber((_5_v0).length))], globalState));
      let _60_v45;
      let _nw10 = Array((new BigNumber(7)).toNumber()).fill([]);
      _60_v45 = _nw10;
      let _61_v46;
      let _init1 = ((_62_v1) => function (_63_i4) {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_62_v1, new BigNumber((_dafny.MultiSet.fromElements(_62_v1)).cardinality()))).length),_62_v1);
      })(_7_v1);
      let _nw11 = Array((new BigNumber(26)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw11.length); _i0_1++) {
        _nw11[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _61_v46 = _nw11;
      let _index8 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_60_v45).length));
      (_60_v45)[_index8] = _61_v46;
      let _index9 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_60_v45).length));
      (_60_v45)[_index9] = _61_v46;
      r0 = _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(139), _7_v1, _7_v1), _8_v2);
      r1 = _7_v1;
      r2 = p0;
      return [r0, r1, r2];
    }
    static Main(__noArgsParameter) {
      let _64_globalState;
      let _nw12 = new _module.GlobalState();
      _nw12.__ctor();
      _64_globalState = _nw12;
      let _65_v0;
      let _nw13 = Array((new BigNumber(7)).toNumber()).fill(false);
      _65_v0 = _nw13;
      let _66_v1;
      _66_v1 = _dafny.Set.fromElements(_65_v0);
      _66_v1 = _66_v1;
      let _67_v2;
      _67_v2 = new BigNumber(98);
      let _68_v3;
      _68_v3 = true;
      let _69_v4;
      _69_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_68_v3, _68_v3, _68_v3)).length),_67_v2);
      let _70_v6;
      _70_v6 = _module.D0.create_DC0(false);
      let _rhs8 = _67_v2;
      let _rhs9 = (function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(-725), new BigNumber(-518))) {
          let _71_v5 = _compr_0;
          if (((new BigNumber(-725)).isLessThanOrEqualTo(_71_v5)) && ((_71_v5).isLessThan(new BigNumber(-518)))) {
            _coll0.push([_module.__default.safeModuloInt(_71_v5, _67_v2),(_dafny.ZERO).minus((_dafny.ZERO).minus(_67_v2))]);
          }
        }
        return _coll0;
      }()).Merge(_69_v4);
      let _rhs10 = (_68_v3) === ((_70_v6).dtor_cf0);
      _67_v2 = _rhs8;
      _69_v4 = _rhs9;
      _68_v3 = _rhs10;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_65_v0).length))) {
        let _72_i0 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_72_i0)) && ((_72_i0).isLessThan(new BigNumber((_65_v0).length))))) {
          (_65_v0)[(_72_i0)] = _68_v3;
        }
      }
      let _73_v7;
      _73_v7 = new _dafny.CodePoint('y'.codePointAt(0));
      let _74_v8;
      _74_v8 = _dafny.Map.Empty.slice().updateUnsafe(_73_v7,_68_v3);
      let _hi0 = _67_v2;
      for (let _75_i1 = new BigNumber((_74_v8).length); _75_i1.isLessThan(_hi0); _75_i1 = _75_i1.plus(_dafny.ONE)) {
        let _76_v9;
        _76_v9 = _module.D0.create_DC1();
        _76_v9 = _76_v9;
        let _77_v10;
        _77_v10 = _dafny.Set.fromElements(_68_v3, _68_v3);
        let _78_v11;
        let _79_v12;
        let _80_v13;
        let _out0;
        let _out1;
        let _out2;
        let _outcollector0 = _module.__default.m0(!((_77_v10).IsSubsetOf(_module.__default.fm0(_64_globalState))), !(true) || (true), _64_globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _out2 = _outcollector0[2];
        _78_v11 = _out0;
        _79_v12 = _out1;
        _80_v13 = _out2;
        let _rhs11 = _module.__default.fm1(_80_v13, _80_v13, _79_v12, _64_globalState);
        let _rhs12 = _73_v7;
        let _rhs13 = _80_v13;
        _68_v3 = _rhs11;
        _73_v7 = _rhs12;
        _68_v3 = _rhs13;
        _77_v10 = _77_v10;
      }
      let _81_v14;
      _81_v14 = _dafny.Seq.UnicodeFromString("braqfma");
      let _82_v15;
      _82_v15 = _module.D0.create_DC1();
      let _83_v16;
      let _nw14 = Array((new BigNumber(12)).toNumber());
      _nw14[(_dafny.ZERO).toNumber()] = _82_v15;
      _nw14[(_dafny.ONE).toNumber()] = _82_v15;
      _nw14[(new BigNumber(2)).toNumber()] = _82_v15;
      _nw14[(new BigNumber(3)).toNumber()] = _82_v15;
      _nw14[(new BigNumber(4)).toNumber()] = _82_v15;
      _nw14[(new BigNumber(5)).toNumber()] = _module.__default.fm2(_67_v2, _64_globalState);
      _nw14[(new BigNumber(6)).toNumber()] = _82_v15;
      _nw14[(new BigNumber(7)).toNumber()] = _module.D0.create_DC1();
      _nw14[(new BigNumber(8)).toNumber()] = _82_v15;
      _nw14[(new BigNumber(9)).toNumber()] = _82_v15;
      _nw14[(new BigNumber(10)).toNumber()] = _82_v15;
      _nw14[(new BigNumber(11)).toNumber()] = _82_v15;
      _83_v16 = _nw14;
      let _84_v17;
      _84_v17 = _dafny.Map.Empty.slice().updateUnsafe(_83_v16,_68_v3);
      let _rhs14 = _dafny.Seq.Concat(_dafny.Seq.Concat(_81_v14, _81_v14), _dafny.Seq.UnicodeFromString("dvtinys"));
      let _rhs15 = _68_v3;
      let _rhs16 = _68_v3;
      let _rhs17 = (_84_v17).Merge(_84_v17);
      let _rhs18 = _70_v6;
      _81_v14 = _rhs14;
      _68_v3 = _rhs15;
      _68_v3 = _rhs16;
      _84_v17 = _rhs17;
      _70_v6 = _rhs18;
      let _85_v18;
      _85_v18 = _dafny.MultiSet.fromElements(_81_v14, _81_v14, _81_v14);
      let _rhs19 = true;
      let _rhs20 = (_67_v2).plus((new BigNumber((_85_v18).cardinality())).plus(_67_v2));
      _68_v3 = _rhs19;
      _67_v2 = _rhs20;
      let _86_i2;
      _86_i2 = _dafny.ZERO;
      L0: {
        while ((_68_v3) === (!(_68_v3))) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_86_i2)) {
              break L0;
            }
            _86_i2 = (_86_i2).plus(_dafny.ONE);
            _67_v2 = (new BigNumber(175)).plus(_67_v2);
            let _87_v19;
            let _88_v20;
            let _89_v21;
            let _out3;
            let _out4;
            let _out5;
            let _outcollector1 = _module.__default.m0(false, _68_v3, _64_globalState);
            _out3 = _outcollector1[0];
            _out4 = _outcollector1[1];
            _out5 = _outcollector1[2];
            _87_v19 = _out3;
            _88_v20 = _out4;
            _89_v21 = _out5;
            let _90_v22;
            _90_v22 = _dafny.MultiSet.fromElements(_89_v21);
            _89_v21 = !(_67_v2).isEqualTo(_module.__default.safeModuloInt(_88_v20, new BigNumber((_90_v22).cardinality())));
            _89_v21 = _89_v21;
          }
        }
      }
      _81_v14 = _dafny.Seq.Concat(_module.__default.fm3(_64_globalState), _dafny.Seq.UnicodeFromString("ttjivi"));
      if (!(_68_v3)) {
        _68_v3 = _68_v3;
        let _91_v23;
        _91_v23 = _dafny.Map.Empty.slice().updateUnsafe(_67_v2,_68_v3);
        let _92_v24;
        _92_v24 = _dafny.Map.Empty.slice().updateUnsafe(_68_v3,_67_v2);
        let _93_v25;
        _93_v25 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_68_v3,new BigNumber((_91_v23).length)), _92_v24);
        let _94_v26;
        _94_v26 = _dafny.Map.Empty.slice().updateUnsafe(false,(_module.__default.fm4(_67_v2, new BigNumber((_93_v25).length), _68_v3, _68_v3, _64_globalState)).minus(new BigNumber(-707)));
        _92_v24 = (_92_v24).update(_68_v3, (_67_v2).minus(_67_v2));
        let _95_v27;
        _95_v27 = _dafny.Seq.of(_68_v3);
        _68_v3 = (_95_v27)[_module.__default.safeIndex(_67_v2, new BigNumber((_95_v27).length))];
        _67_v2 = new BigNumber(640);
        _67_v2 = _67_v2;
      } else {
        if (_68_v3) {
          let _96_v28;
          let _nw15 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
          _96_v28 = _nw15;
          let _index10 = _module.__default.safeIndex(new BigNumber(611), new BigNumber((_96_v28).length));
          (_96_v28)[_index10] = _67_v2;
          let _97_v29;
          _97_v29 = _dafny.MultiSet.fromElements(_73_v7, new _dafny.CodePoint('k'.codePointAt(0)), _73_v7);
          let _98_v30;
          _98_v30 = _dafny.Map.Empty.slice().updateUnsafe(_68_v3,new BigNumber(-34));
          let _99_v31;
          _99_v31 = _dafny.MultiSet.fromElements(_68_v3, _68_v3, _68_v3);
          let _100_v32;
          _100_v32 = _dafny.MultiSet.fromElements((((_97_v29).contains(_73_v7)) ? ((_97_v29).get(_73_v7)) : (_module.__default.fm4(_67_v2, _67_v2, false, _68_v3, _64_globalState))), new BigNumber((_dafny.MultiSet.fromElements(_68_v3)).cardinality()), _67_v2, (new BigNumber(((_98_v30).update(_module.__default.fm1(_68_v3, _68_v3, _67_v2, _64_globalState), _67_v2)).length)).plus(new BigNumber((_99_v31).cardinality())), _67_v2);
          let _index11 = _module.__default.safeIndex(new BigNumber(611), new BigNumber((_96_v28).length));
          (_96_v28)[_index11] = (((_100_v32).contains(_67_v2)) ? ((_100_v32).get(_67_v2)) : (_module.__default.safeDivisionInt(new BigNumber(850), (_dafny.ZERO).minus(_67_v2))));
          let _101_v33;
          _101_v33 = _dafny.Seq.of(false);
          let _102_v34;
          let _nw16 = new _module.C0();
          _nw16.__ctor(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray(_101_v33)).cardinality())));
          _102_v34 = _nw16;
          _81_v14 = _81_v14;
          _73_v7 = (_81_v14)[_module.__default.safeIndex(_67_v2, new BigNumber((_81_v14).length))];
          let _103_v35;
          let _nw17 = Array((new BigNumber(9)).toNumber());
          _103_v35 = _nw17;
          let _104_v36;
          _104_v36 = _dafny.Map.Empty.slice().updateUnsafe(_103_v35,_68_v3);
          let _105_v37;
          let _106_v38;
          let _107_v39;
          let _out6;
          let _out7;
          let _out8;
          let _outcollector2 = _module.__default.m0(_68_v3, !(((_dafny.ZERO).minus(new BigNumber((_104_v36).length))).isLessThan((_102_v34).fm5(_68_v3, _68_v3, _68_v3, (_96_v28)[_module.__default.safeIndex(new BigNumber(611), new BigNumber((_96_v28).length))], _64_globalState))), _64_globalState);
          _out6 = _outcollector2[0];
          _out7 = _outcollector2[1];
          _out8 = _outcollector2[2];
          _105_v37 = _out6;
          _106_v38 = _out7;
          _107_v39 = _out8;
        } else {
          let _108_v40;
          _108_v40 = _dafny.Set.fromElements(!(_68_v3));
          let _109_v41;
          _109_v41 = _dafny.Map.Empty.slice().updateUnsafe(_70_v6,_108_v40);
          _68_v3 = ((_68_v3) ? (_68_v3) : ((_109_v41).equals(_109_v41)));
          _67_v2 = (_67_v2).plus((_dafny.ZERO).minus(_67_v2));
          _67_v2 = _67_v2;
          let _110_v42;
          _110_v42 = _dafny.Map.Empty.slice().updateUnsafe(_68_v3,_68_v3);
          _110_v42 = (_110_v42).Merge(_110_v42);
          let _111_v43;
          _111_v43 = _dafny.Seq.of(true, _68_v3, false, _68_v3, _68_v3);
          let _112_v44;
          _112_v44 = _dafny.MultiSet.fromElements(_module.D0.create_DC1());
          let _113_v45;
          _113_v45 = _dafny.Map.Empty.slice().updateUnsafe((_111_v43)[_module.__default.safeIndex(_67_v2, new BigNumber((_111_v43).length))],(((_112_v44).contains(_module.D0.create_DC1())) ? ((_112_v44).get(_module.D0.create_DC1())) : (_67_v2)));
          _67_v2 = (_module.__default.safeModuloInt(new BigNumber((_113_v45).length), _67_v2)).plus(_67_v2);
        }
        let _114_v46;
        _114_v46 = _dafny.Map.Empty.slice().updateUnsafe(_68_v3,_68_v3);
        _114_v46 = (_114_v46).update(_68_v3, (_70_v6).dtor_cf0);
        let _115_v47;
        _115_v47 = _dafny.MultiSet.fromElements(_68_v3, _68_v3, _68_v3);
        let _116_v48;
        let _nw18 = new _module.C0();
        _nw18.__ctor(_dafny.Seq.of(new BigNumber((_81_v14).length), new BigNumber(706), new BigNumber((_115_v47).cardinality())));
        _116_v48 = _nw18;
        let _117_v49;
        _117_v49 = _dafny.Map.Empty.slice().updateUnsafe(false,_116_v48);
        let _118_v50;
        _118_v50 = _module.D1.create_DC2((((_117_v49).contains(_68_v3)) ? ((_117_v49).get(_68_v3)) : (_116_v48)));
        let _119_v51;
        let _nw19 = Array((new BigNumber(16)).toNumber());
        _nw19[(_dafny.ZERO).toNumber()] = _116_v48;
        _nw19[(_dafny.ONE).toNumber()] = _116_v48;
        _nw19[(new BigNumber(2)).toNumber()] = _116_v48;
        _nw19[(new BigNumber(3)).toNumber()] = _116_v48;
        _nw19[(new BigNumber(4)).toNumber()] = ((true) ? (_116_v48) : (_116_v48));
        _nw19[(new BigNumber(5)).toNumber()] = _116_v48;
        _nw19[(new BigNumber(6)).toNumber()] = _116_v48;
        _nw19[(new BigNumber(7)).toNumber()] = _116_v48;
        _nw19[(new BigNumber(8)).toNumber()] = _116_v48;
        _nw19[(new BigNumber(9)).toNumber()] = _116_v48;
        _nw19[(new BigNumber(10)).toNumber()] = _116_v48;
        _nw19[(new BigNumber(11)).toNumber()] = _116_v48;
        _nw19[(new BigNumber(12)).toNumber()] = (_118_v50).dtor_cf1;
        _nw19[(new BigNumber(13)).toNumber()] = _116_v48;
        _nw19[(new BigNumber(14)).toNumber()] = _116_v48;
        _nw19[(new BigNumber(15)).toNumber()] = _116_v48;
        _119_v51 = _nw19;
        let _index12 = _module.__default.safeIndex(new BigNumber(539), new BigNumber((_119_v51).length));
        (_119_v51)[_index12] = _116_v48;
        let _index13 = _module.__default.safeIndex(new BigNumber(539), new BigNumber((_119_v51).length));
        (_119_v51)[_index13] = _116_v48;
        _68_v3 = _68_v3;
        let _index14 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_65_v0).length));
        (_65_v0)[_index14] = (_68_v3) === (_68_v3);
        let _index15 = _module.__default.safeIndex(new BigNumber(866), new BigNumber((_65_v0).length));
        (_65_v0)[_index15] = _68_v3;
      }
      if ((_70_v6).dtor_cf0) {
        let _120_v52;
        _120_v52 = _dafny.Seq.of(_68_v3);
        let _121_v53;
        _121_v53 = _dafny.MultiSet.fromElements(_68_v3, _68_v3);
        _67_v2 = (((_121_v53).IsProperSubsetOf(_dafny.MultiSet.FromArray(_120_v52))) ? (_67_v2) : (((_dafny.ZERO).minus(new BigNumber((_120_v52).length))).multipliedBy(_67_v2)));
        let _122_v54;
        _122_v54 = _dafny.Set.fromElements(true);
        _122_v54 = ((_68_v3) ? ((_dafny.Set.fromElements(false, _68_v3)).Difference(_dafny.Set.fromElements(!(_68_v3)))) : (_dafny.Set.fromElements(_68_v3, _68_v3, _68_v3)));
        let _123_v55;
        let _init2 = ((_124_v2) => function (_125_i3) {
          return (_125_i3).minus(_124_v2);
        })(_67_v2);
        let _nw20 = Array((new BigNumber(25)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw20.length); _i0_2++) {
          _nw20[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _123_v55 = _nw20;
        let _index16 = _module.__default.safeIndex(new BigNumber(637), new BigNumber((_123_v55).length));
        (_123_v55)[_index16] = _67_v2;
        let _index17 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_123_v55).length));
        (_123_v55)[_index17] = new BigNumber(-266);
        let _126_v56;
        _126_v56 = _dafny.Seq.of(_67_v2);
        let _127_v57;
        _127_v57 = _dafny.Seq.of(new BigNumber((_126_v56).length), new BigNumber((_126_v56).length));
        let _128_v58;
        let _nw21 = new _module.C0();
        _nw21.__ctor(_127_v57);
        _128_v58 = _nw21;
        let _129_v59;
        _129_v59 = _dafny.Map.Empty.slice().updateUnsafe(_128_v58,true);
        let _130_v60;
        _130_v60 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_81_v14).length),_129_v59);
        let _index18 = _module.__default.safeIndex(new BigNumber(637), new BigNumber((_123_v55).length));
        let _index19 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_123_v55).length));
        let _rhs21 = (_dafny.ZERO).minus(_67_v2);
        let _rhs22 = (_70_v6).dtor_cf0;
        let _rhs23 = _67_v2;
        let _rhs24 = new BigNumber(246);
        let _rhs25 = ((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_81_v14).length), (_module.D1.create_DC4(_130_v60, _67_v2)).dtor_cf5))).minus((_128_v58).fm5(false, _68_v3, _68_v3, _module.__default.fm4(_67_v2, _67_v2, _68_v3, _68_v3, _64_globalState), _64_globalState));
        let _lhs4 = _123_v55;
        let _lhs5 = _module.__default.safeIndex(new BigNumber(637), new BigNumber((_123_v55).length));
        let _lhs6 = _123_v55;
        let _lhs7 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_123_v55).length));
        _67_v2 = _rhs21;
        _68_v3 = _rhs22;
        _lhs4[_lhs5] = _rhs23;
        _67_v2 = _rhs24;
        _lhs6[_lhs7] = _rhs25;
        _68_v3 = (_121_v53).IsSubsetOf(_dafny.MultiSet.FromArray(_120_v52));
        _68_v3 = true;
      } else {
        _69_v4 = (_69_v4).update(_67_v2, (_dafny.ZERO).minus(_67_v2));
        _67_v2 = new BigNumber(935);
        _67_v2 = (_dafny.ZERO).minus(_67_v2);
        _68_v3 = !(_68_v3);
        let _131_v61;
        _131_v61 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_68_v3),_67_v2);
        let _132_v62;
        _132_v62 = _dafny.MultiSet.fromElements(_68_v3);
        _68_v3 = ((((_131_v61).contains(_132_v62)) ? ((_131_v61).get(_132_v62)) : (_67_v2))).isLessThanOrEqualTo(_67_v2);
      }
      let _133_i4;
      _133_i4 = _dafny.ZERO;
      L1: {
        while (_68_v3) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_133_i4)) {
              break L1;
            }
            _133_i4 = (_133_i4).plus(_dafny.ONE);
            _67_v2 = _67_v2;
            let _134_v63;
            _134_v63 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(_67_v2)),_68_v3);
            let _135_v64;
            let _nw22 = new _module.C0();
            _nw22.__ctor((_module.__default.fm6(_67_v2, _68_v3, (((_134_v63).contains(new BigNumber(834))) ? ((_134_v63).get(new BigNumber(834))) : (_68_v3)), _64_globalState)));
            _135_v64 = _nw22;
            _81_v14 = _81_v14;
            _67_v2 = ((_68_v3) ? (_67_v2) : (_67_v2));
          }
        }
      }
      _68_v3 = _68_v3;
      _68_v3 = (_module.__default.fm4(_67_v2, (_dafny.ZERO).minus(_67_v2), false, true, _64_globalState)).isLessThanOrEqualTo(_67_v2);
      let _136_v65;
      _136_v65 = _dafny.Set.fromElements(new BigNumber((_81_v14).length));
      let _137_v66;
      _137_v66 = _dafny.Seq.of(_67_v2);
      let _138_v67;
      let _nw23 = new _module.C0();
      _nw23.__ctor(_137_v66);
      _138_v67 = _nw23;
      let _139_v68;
      _139_v68 = _dafny.Map.Empty.slice().updateUnsafe(_138_v67,_68_v3);
      let _140_v69;
      _140_v69 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(232),_139_v68);
      let _141_v70;
      _141_v70 = _module.D1.create_DC4(_140_v69, new BigNumber((_dafny.MultiSet.fromElements(!(_68_v3), _68_v3)).cardinality()));
      let _142_v71;
      _142_v71 = _dafny.Map.Empty.slice().updateUnsafe(_136_v65,_141_v70);
      let _143_v72;
      _143_v72 = _dafny.Seq.of(new BigNumber((_142_v71).length), (_dafny.ZERO).minus(new BigNumber(-911)), new BigNumber((_81_v14).length));
      let _144_v73;
      let _nw24 = new _module.C0();
      _nw24.__ctor(_137_v66);
      _144_v73 = _nw24;
      let _145_v74;
      let _146_v75;
      let _147_v76;
      let _out9;
      let _out10;
      let _out11;
      let _outcollector3 = _module.__default.m0(_68_v3, (_70_v6).dtor_cf0, _64_globalState);
      _out9 = _outcollector3[0];
      _out10 = _outcollector3[1];
      _out11 = _outcollector3[2];
      _145_v74 = _out9;
      _146_v75 = _out10;
      _147_v76 = _out11;
      _146_v75 = ((_68_v3) ? (new BigNumber(-550)) : (_module.__default.safeModuloInt((_dafny.ZERO).minus(_67_v2), (_dafny.ZERO).minus(_146_v75))));
      process.stdout.write(_dafny.toString((_65_v0)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_65_v0)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_65_v0)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_65_v0)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_65_v0)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_65_v0)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_65_v0)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_66_v1).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_67_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_68_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v4).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(59),new BigNumber(98)).updateUnsafe(new BigNumber(60),new BigNumber(98)).updateUnsafe(new BigNumber(61),new BigNumber(98)).updateUnsafe(new BigNumber(62),new BigNumber(98)).updateUnsafe(new BigNumber(63),new BigNumber(98)).updateUnsafe(new BigNumber(64),new BigNumber(98)).updateUnsafe(new BigNumber(65),new BigNumber(98)).updateUnsafe(new BigNumber(66),new BigNumber(98)).updateUnsafe(new BigNumber(67),new BigNumber(98)).updateUnsafe(new BigNumber(68),new BigNumber(98)).updateUnsafe(new BigNumber(69),new BigNumber(98)).updateUnsafe(new BigNumber(70),new BigNumber(98)).updateUnsafe(new BigNumber(71),new BigNumber(98)).updateUnsafe(new BigNumber(72),new BigNumber(98)).updateUnsafe(new BigNumber(73),new BigNumber(98)).updateUnsafe(new BigNumber(74),new BigNumber(98)).updateUnsafe(new BigNumber(75),new BigNumber(98)).updateUnsafe(new BigNumber(76),new BigNumber(98)).updateUnsafe(new BigNumber(77),new BigNumber(98)).updateUnsafe(new BigNumber(78),new BigNumber(98)).updateUnsafe(new BigNumber(79),new BigNumber(98)).updateUnsafe(new BigNumber(80),new BigNumber(98)).updateUnsafe(new BigNumber(81),new BigNumber(98)).updateUnsafe(new BigNumber(82),new BigNumber(98)).updateUnsafe(new BigNumber(83),new BigNumber(98)).updateUnsafe(new BigNumber(84),new BigNumber(98)).updateUnsafe(new BigNumber(85),new BigNumber(98)).updateUnsafe(new BigNumber(86),new BigNumber(98)).updateUnsafe(new BigNumber(87),new BigNumber(98)).updateUnsafe(new BigNumber(88),new BigNumber(98)).updateUnsafe(new BigNumber(89),new BigNumber(98)).updateUnsafe(new BigNumber(90),new BigNumber(98)).updateUnsafe(new BigNumber(91),new BigNumber(98)).updateUnsafe(new BigNumber(92),new BigNumber(98)).updateUnsafe(new BigNumber(93),new BigNumber(98)).updateUnsafe(new BigNumber(94),new BigNumber(98)).updateUnsafe(new BigNumber(95),new BigNumber(98)).updateUnsafe(new BigNumber(96),new BigNumber(98)).updateUnsafe(new BigNumber(97),new BigNumber(98)).updateUnsafe(_dafny.ZERO,new BigNumber(98)).updateUnsafe(_dafny.ONE,new BigNumber(98)).updateUnsafe(new BigNumber(2),new BigNumber(98)).updateUnsafe(new BigNumber(3),new BigNumber(98)).updateUnsafe(new BigNumber(4),new BigNumber(98)).updateUnsafe(new BigNumber(5),new BigNumber(98)).updateUnsafe(new BigNumber(6),new BigNumber(98)).updateUnsafe(new BigNumber(7),new BigNumber(98)).updateUnsafe(new BigNumber(8),new BigNumber(98)).updateUnsafe(new BigNumber(9),new BigNumber(98)).updateUnsafe(new BigNumber(10),new BigNumber(98)).updateUnsafe(new BigNumber(11),new BigNumber(98)).updateUnsafe(new BigNumber(12),new BigNumber(98)).updateUnsafe(new BigNumber(13),new BigNumber(98)).updateUnsafe(new BigNumber(14),new BigNumber(98)).updateUnsafe(new BigNumber(15),new BigNumber(98)).updateUnsafe(new BigNumber(16),new BigNumber(98)).updateUnsafe(new BigNumber(17),new BigNumber(98)).updateUnsafe(new BigNumber(18),new BigNumber(98)).updateUnsafe(new BigNumber(19),new BigNumber(98)).updateUnsafe(new BigNumber(20),new BigNumber(98)).updateUnsafe(new BigNumber(21),new BigNumber(98)).updateUnsafe(new BigNumber(22),new BigNumber(98)).updateUnsafe(new BigNumber(23),new BigNumber(98)).updateUnsafe(new BigNumber(24),new BigNumber(98)).updateUnsafe(new BigNumber(25),new BigNumber(98)).updateUnsafe(new BigNumber(26),new BigNumber(98)).updateUnsafe(new BigNumber(27),new BigNumber(98)).updateUnsafe(new BigNumber(28),new BigNumber(98)).updateUnsafe(new BigNumber(29),new BigNumber(98)).updateUnsafe(new BigNumber(30),new BigNumber(98)).updateUnsafe(new BigNumber(31),new BigNumber(98)).updateUnsafe(new BigNumber(32),new BigNumber(98)).updateUnsafe(new BigNumber(33),new BigNumber(98)).updateUnsafe(new BigNumber(34),new BigNumber(98)).updateUnsafe(new BigNumber(35),new BigNumber(98)).updateUnsafe(new BigNumber(36),new BigNumber(98)).updateUnsafe(new BigNumber(37),new BigNumber(98)).updateUnsafe(new BigNumber(38),new BigNumber(98)).updateUnsafe(new BigNumber(39),new BigNumber(98)).updateUnsafe(new BigNumber(40),new BigNumber(98)).updateUnsafe(new BigNumber(41),new BigNumber(98)).updateUnsafe(new BigNumber(42),new BigNumber(98)).updateUnsafe(new BigNumber(43),new BigNumber(98)).updateUnsafe(new BigNumber(44),new BigNumber(98)).updateUnsafe(new BigNumber(45),new BigNumber(98)).updateUnsafe(new BigNumber(46),new BigNumber(98)).updateUnsafe(new BigNumber(47),new BigNumber(98)).updateUnsafe(new BigNumber(48),new BigNumber(98)).updateUnsafe(new BigNumber(49),new BigNumber(98)).updateUnsafe(new BigNumber(50),new BigNumber(98)).updateUnsafe(new BigNumber(51),new BigNumber(98)).updateUnsafe(new BigNumber(52),new BigNumber(98)).updateUnsafe(new BigNumber(53),new BigNumber(98)).updateUnsafe(new BigNumber(54),new BigNumber(98)).updateUnsafe(new BigNumber(55),new BigNumber(98)).updateUnsafe(new BigNumber(56),new BigNumber(98)).updateUnsafe(new BigNumber(57),new BigNumber(98)).updateUnsafe(new BigNumber(58),new BigNumber(98)).updateUnsafe(new BigNumber(199),new BigNumber(-199)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_70_v6).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_73_v7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_74_v8).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('y'.codePointAt(0)),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_81_v14).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_84_v17).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v18).equals(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("braqfmabraqfmadvtinys"), _dafny.Seq.UnicodeFromString("braqfmabraqfmadvtinys"), _dafny.Seq.UnicodeFromString("braqfmabraqfmadvtinys")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_86_i2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_133_i4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_v65).equals(_dafny.Set.fromElements(new BigNumber(7)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_137_v66, _dafny.Seq.of(new BigNumber(-935)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_138_v67).f0, _dafny.Seq.of(new BigNumber(-935)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_139_v68).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_140_v69).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_141_v70).dtor_cf4).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_141_v70).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_142_v71).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_143_v72, _dafny.Seq.of(_dafny.ONE, new BigNumber(911), new BigNumber(7)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_144_v73).f0, _dafny.Seq.of(new BigNumber(-935)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_145_v74, _dafny.Seq.of(new BigNumber(139), new BigNumber(14), new BigNumber(14), new BigNumber(14)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_146_v75));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_147_v76));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1() {
      let $dt = new D0(0);
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(1);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC0() { return this.$tag === 1; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1";
      } else if (this.$tag === 1) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
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
        return other.$tag === 1 && this.cf0 === other.cf0;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1();
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
    static create_DC3(cf2, cf3) {
      let $dt = new D1(0);
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC4(cf4, cf5) {
      let $dt = new D1(1);
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC2(cf1) {
      let $dt = new D1(2);
      $dt.cf1 = cf1;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get is_DC4() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf1() { return this.cf1; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf1) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf2, other.cf2) && _dafny.areEqual(this.cf3, other.cf3);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf4, other.cf4) && _dafny.areEqual(this.cf5, other.cf5);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf1 === other.cf1;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC3(new _dafny.CodePoint('D'.codePointAt(0)), _dafny.ZERO);
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
    static create_DC5(cf6) {
      let $dt = new D2(0);
      $dt.cf6 = cf6;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get dtor_cf6() { return this.cf6; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf6) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf6, other.cf6);
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
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC7(cf8, cf9) {
      let $dt = new D3(0);
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC6(cf7) {
      let $dt = new D3(1);
      $dt.cf7 = cf7;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf7() { return this.cf7; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC7" + "(" + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC6" + "(" + _dafny.toString(this.cf7) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf8 === other.cf8 && this.cf9 === other.cf9;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf7, other.cf7);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC7(false, false);
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
    static create_DC8(cf10) {
      let $dt = new D4(0);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC8" + "(" + _dafny.toString(this.cf10) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf10 === other.cf10;
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
          return D4.Default();
        }
      };
    }
  }

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f0 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0) {
      let _this = this;
      (_this)._f0 = f0;
      return;
    }
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return new BigNumber(-344);
    };
    get f0() {
      let _this = this;
      return _this._f0;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
