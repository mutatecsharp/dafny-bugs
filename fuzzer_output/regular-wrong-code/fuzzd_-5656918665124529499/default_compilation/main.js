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
      return (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(934)), function (_0_i0) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("goxbqw"))).length)).isLessThan(_module.__default.safeModuloInt(new BigNumber(763), new BigNumber(294)));
    };
    static fm1(p0, p1, p2, globalState) {
      if ((true) === (true)) {
        return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(882)), function (_1_i0) {
          return new _dafny.CodePoint('y'.codePointAt(0));
        }), _dafny.Seq.UnicodeFromString("ov"))).length);
      } else {
        return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0))), _dafny.Seq.of(new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0))))).length);
      }
    };
    static fm5(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.of((new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(710),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(816),new _dafny.CodePoint('l'.codePointAt(0)))).length))).length),true)).length))).length)).isLessThan(new BigNumber((_dafny.Seq.UnicodeFromString("irm")).length)), (_module.D9.create_DC25(_dafny.Seq.UnicodeFromString("ap"), false, new BigNumber(214))).dtor_cf60, false);
    };
    static fm7(globalState) {
      if (!((new BigNumber(306)).isLessThanOrEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.UnicodeFromString("sfmw"))).length)))) {
        return _dafny.Seq.of(true, false);
      } else {
        return _dafny.Seq.of(true, !(!(true)));
      }
    };
    static fm9(p0, p1, globalState) {
      return _dafny.Seq.UnicodeFromString("x");
    };
    static fm10(p0, p1, p2, globalState) {
      if ((true) || (false)) {
        return _dafny.Seq.of(false);
      } else {
        return _dafny.Seq.of(true);
      }
    };
    static fm13(p0, p1, globalState) {
      return ((_dafny.MultiSet.fromElements(new BigNumber((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(-834), new BigNumber(-776))) {
          let _2_v0 = _compr_0;
          if (((new BigNumber(-834)).isLessThanOrEqualTo(_2_v0)) && ((_2_v0).isLessThan(new BigNumber(-776)))) {
            _coll0.push([_module.__default.safeModuloInt(_2_v0, new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality())),new BigNumber(54)]);
          }
        }
        return _coll0;
      }()).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(false)).length))).length))).length)))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false),false)).length), new BigNumber((_dafny.Set.fromElements(!((_module.D8.create_DC20(true, false, false)).dtor_cf48))).length))).Intersect(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(false, true)).length), _dafny.ZERO, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_module.D7.create_DC18(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true)).length),new BigNumber(552))).length), _dafny.Seq.UnicodeFromString("dvkfi"), true, !(false), function () {
  let _coll1 = new _dafny.Set();
  for (const _compr_1 of (_dafny.Set.fromElements(new BigNumber(-683))).Elements) {
    let _3_v1 = _compr_1;
    if ((_dafny.Set.fromElements(new BigNumber(-683))).contains(_3_v1)) {
      _coll1.add(_module.__default.safeDivisionInt(_3_v1, new BigNumber((_dafny.Seq.UnicodeFromString("mqnecwgoe")).length)));
    }
  }
  return _coll1;
}())).dtor_cf43,new _dafny.CodePoint('r'.codePointAt(0)))).length))).length)), new BigNumber((_dafny.Seq.of(new BigNumber(951), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-863),true)).length), new BigNumber(159), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),true)).length), new BigNumber((_dafny.Seq.of(false)).length))).length)))).Union((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(-677)))).Intersect(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true)).length), new BigNumber(861))).length), new BigNumber((_dafny.Seq.UnicodeFromString("od")).length), new BigNumber((_dafny.Seq.of(_dafny.ONE)).length))));
    };
    static fm14(p0, p1, globalState) {
      return new _dafny.CodePoint('n'.codePointAt(0));
    };
    static fm15(p0, p1, p2, globalState) {
      return _module.D1.create_DC2((_module.D9.create_DC25(_dafny.Seq.UnicodeFromString("klphfov"), !(true), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("arkhat")).length),false)).length)))).dtor_cf60, true, false, ((true) ? (new _dafny.CodePoint('k'.codePointAt(0))) : (new _dafny.CodePoint('g'.codePointAt(0)))), new _dafny.CodePoint('k'.codePointAt(0)));
    };
    static fm16(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(!(false)),true));
    };
    static fm17(p0, p1, p2, p3, globalState) {
      return _dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(172)), function (_4_i0) {
        return _dafny.Set.fromElements(_module.D1.create_DC3(_module.D1.create_DC1(_dafny.Set.fromElements(true))));
      }));
    };
    static fm18(p0, p1, globalState) {
      return ((_dafny.Set.fromElements(false, false, !(true), true, false)).Union(_dafny.Set.fromElements(true, !(false)))).Union(_dafny.Set.fromElements(false, true));
    };
    static fm19(p0, globalState) {
      return _dafny.MultiSet.fromElements((_module.D10.create_DC27(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(143)))).dtor_cf63, (_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(176))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(694))));
    };
    static fm20(p0, p1, p2, p3, globalState) {
      return (function () {
        let _coll2 = new _dafny.Set();
        for (const _compr_2 of _dafny.IntegerRange(new BigNumber(514), new BigNumber(-236))) {
          let _5_v0 = _compr_2;
          if (((new BigNumber(514)).isLessThanOrEqualTo(_5_v0)) && ((_5_v0).isLessThan(new BigNumber(-236)))) {
            _coll2.add((_5_v0).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-183)), function (_6_i0) {
              return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber((function () {
                let _coll3 = new _dafny.Map();
                for (const _compr_3 of _dafny.IntegerRange(new BigNumber(740), new BigNumber(303))) {
                  let _7_v1 = _compr_3;
                  if (((new BigNumber(740)).isLessThanOrEqualTo(_7_v1)) && ((_7_v1).isLessThan(new BigNumber(303)))) {
                    _coll3.push([(_7_v1).plus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(278))).cardinality())),new BigNumber((_dafny.Seq.UnicodeFromString("wag")).length)]);
                  }
                }
                return _coll3;
              }()).length))).length);
            })).length)));
          }
        }
        return _coll2;
      }()).Difference((function () {
        let _coll4 = new _dafny.Set();
        for (const _compr_4 of _dafny.IntegerRange(new BigNumber(-46), new BigNumber(53))) {
          let _8_v2 = _compr_4;
          if (((new BigNumber(-46)).isLessThanOrEqualTo(_8_v2)) && ((_8_v2).isLessThan(new BigNumber(53)))) {
            _coll4.add((_8_v2).multipliedBy(new BigNumber(370)));
          }
        }
        return _coll4;
      }()).Union(_dafny.Set.fromElements(new BigNumber(-604), (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements((_module.D1.create_DC2(!(true), true, false, new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)))).dtor_cf2)).length)), new BigNumber(806))));
    };
    static fm21(p0, p1, p2, globalState) {
      return (_module.D15.create_DC36(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),new BigNumber(123)))).dtor_cf78;
    };
    static fm22(p0, p1, globalState) {
      return _dafny.MultiSet.fromElements((new BigNumber(492)).isLessThanOrEqualTo(new BigNumber(754)));
    };
    static fm23(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("oot")).length),false)).length));
    };
    static fm24(globalState) {
      return _dafny.Seq.of(new BigNumber(120), (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length))).length), new BigNumber(712))), new BigNumber(-813), ((!(true)) ? (new BigNumber(941)) : (new BigNumber((_dafny.Seq.UnicodeFromString("ldpqylslw")).length))));
    };
    static fm25(p0, p1, p2, globalState) {
      return ((function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("aua"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(66)), function (_9_i0) {
          return new _dafny.CodePoint('c'.codePointAt(0));
        }))).Elements) {
          let _10_v0 = _compr_5;
          if ((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("aua"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(66)), function (_9_i0) {
            return new _dafny.CodePoint('c'.codePointAt(0));
          }))).contains(_10_v0)) {
            _coll5.push([_10_v0,_dafny.Seq.UnicodeFromString("jhdbecap")]);
          }
        }
        return _coll5;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-65)), function (_11_i1) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      }),_dafny.Seq.Create(_module.__default.abs(new BigNumber(993)), function (_12_i2) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      })))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("iqvwwh"),_dafny.Seq.UnicodeFromString("oudm")));
    };
    static fm26(p0, p1, globalState) {
      return _module.D8.create_DC20(((false) ? (!(true)) : (true)), true, _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("qoxnse"), _dafny.Seq.UnicodeFromString("rjxfs")));
    };
    static fm27(p0, p1, globalState) {
      return _dafny.Set.fromElements(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(396))).length)), _dafny.Seq.of(new BigNumber(410), new BigNumber(166))));
    };
    static fm28(globalState) {
      if (false) {
        return true;
      } else {
        return true;
      }
    };
    static fm29(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-549),_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-220)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-451),_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(540))));
    };
    static m0(globalState) {
      let _13_v0;
      _13_v0 = false;
      (globalState).f3 = _13_v0;
      let _14_v1;
      _14_v1 = new _dafny.CodePoint('f'.codePointAt(0));
      let _15_v2;
      _15_v2 = _dafny.Map.Empty.slice().updateUnsafe(_13_v0,_14_v1);
      let _16_v3;
      _16_v3 = new BigNumber(692);
      _15_v2 = (_15_v2).update((new BigNumber(390)).isLessThanOrEqualTo(_16_v3), _14_v1);
      let _17_v4;
      _17_v4 = _dafny.Seq.of(_16_v3, new BigNumber(164));
      let _18_i0;
      _18_i0 = _dafny.ZERO;
      L0: {
        while ((_16_v3).isLessThanOrEqualTo(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(new BigNumber((_17_v4).length))).length), _16_v3))) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_18_i0)) {
              break L0;
            }
            _18_i0 = (_18_i0).plus(_dafny.ONE);
            (globalState).f9 = _16_v3;
            if (!(!(!(!(_13_v0))))) {
              let _19_v5;
              let _init0 = ((_20_v3) => function (_21_i1) {
                return !(_20_v3).isEqualTo(_20_v3);
              })(_16_v3);
              let _nw0 = Array((new BigNumber(18)).toNumber());
              for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw0.length); _i0_0++) {
                _nw0[_i0_0] = _init0(new BigNumber(_i0_0));
              }
              _19_v5 = _nw0;
              let _22_v6;
              _22_v6 = _dafny.Seq.UnicodeFromString("luog");
              let _index0 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_19_v5).length));
              (_19_v5)[_index0] = !_dafny.areEqual(_22_v6, _22_v6);
              let _index1 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_19_v5).length));
              (_19_v5)[_index1] = _module.__default.fm0(_14_v1, globalState);
              let _23_v7;
              let _nw1 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
              _23_v7 = _nw1;
              let _index2 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_19_v5).length));
              let _rhs0 = (_dafny.Set.fromElements(_23_v7)).IsProperSubsetOf(_dafny.Set.fromElements(_23_v7));
              let _rhs1 = (((_19_v5)[_module.__default.safeIndex(new BigNumber(830), new BigNumber((_19_v5).length))]) ? (true) : ((_19_v5)[_module.__default.safeIndex(new BigNumber(830), new BigNumber((_19_v5).length))]));
              let _lhs0 = globalState;
              let _lhs1 = _19_v5;
              let _lhs2 = _module.__default.safeIndex(new BigNumber(830), new BigNumber((_19_v5).length));
              _lhs0.f3 = _rhs0;
              _lhs1[_lhs2] = _rhs1;
              (globalState).f7 = false;
              (globalState).f9 = _16_v3;
              (globalState).f9 = _16_v3;
            } else {
              let _24_v8;
              _24_v8 = _dafny.MultiSet.fromElements(_13_v0, false);
              let _25_v9;
              _25_v9 = _dafny.Seq.of(_13_v0, _13_v0);
              let _26_v10;
              _26_v10 = _dafny.Set.fromElements(_13_v0, _13_v0);
              let _27_v11;
              _27_v11 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber(130), _16_v3),_module.__default.fm1(_13_v0, new BigNumber((_26_v10).length), _16_v3, globalState));
              let _28_v12;
              _28_v12 = _dafny.Set.fromElements(_16_v3);
              let _29_v13;
              _29_v13 = _dafny.Seq.UnicodeFromString("jlgqgkipu");
              _16_v3 = ((((_24_v8).contains((_25_v9)[_module.__default.safeIndex(_16_v3, new BigNumber((_25_v9).length))])) ? ((_24_v8).get((_25_v9)[_module.__default.safeIndex(_16_v3, new BigNumber((_25_v9).length))])) : (_16_v3))).plus((((_27_v11).contains(_28_v12)) ? ((_27_v11).get(_28_v12)) : (new BigNumber((_29_v13).length))));
              let _30_v14;
              let _nw2 = new _module.C2();
              _nw2.__ctor(_13_v0, _24_v8, _module.__default.fm0((_29_v13)[_module.__default.safeIndex(_16_v3, new BigNumber((_29_v13).length))], globalState));
              _30_v14 = _nw2;
              let _rhs2 = (_16_v3).minus(((_30_v14.f15) ? (new BigNumber((_29_v13).length)) : (_16_v3)));
              let _rhs3 = new BigNumber(-264);
              let _rhs4 = (_dafny.ZERO).minus(_16_v3);
              let _rhs5 = _13_v0;
              let _lhs3 = globalState;
              _lhs3.f9 = _rhs2;
              _16_v3 = _rhs3;
              _16_v3 = _rhs4;
              _13_v0 = _rhs5;
              let _31_v15;
              _31_v15 = _dafny.MultiSet.fromElements(_16_v3, _16_v3, _16_v3, _16_v3, new BigNumber((_29_v13).length));
              (globalState).f9 = (_module.__default.safeModuloInt((_dafny.ZERO).minus((_dafny.ZERO).minus((((_31_v15).contains(_16_v3)) ? ((_31_v15).get(_16_v3)) : (new BigNumber((_17_v4).length))))), new BigNumber((_dafny.Set.fromElements(_30_v14.f15)).length))).minus(_16_v3);
              (globalState).f7 = _30_v14.f15;
            }
            let _32_v16;
            _32_v16 = _dafny.Map.Empty.slice().updateUnsafe(_13_v0,_16_v3);
            let _33_v17;
            _33_v17 = _module.D10.create_DC27(_32_v16);
            let _pat_let_tv0 = _13_v0;
            let _pat_let_tv1 = _16_v3;
            let _pat_let_tv2 = _32_v16;
            let _source0 = function (_pat_let0_0) {
              return function (_34_dt__update__tmp_h0) {
                return function (_pat_let1_0) {
                  return function (_35_dt__update_hcf63_h0) {
                    return _module.D10.create_DC27(_35_dt__update_hcf63_h0);
                  }(_pat_let1_0);
                }((_dafny.Map.Empty.slice().updateUnsafe(!(_pat_let_tv0),_pat_let_tv1)).Merge(_pat_let_tv2));
              }(_pat_let0_0);
            }(_33_v17);
            if (_source0.is_DC28) {
              let _36___mcc_h0 = (_source0).cf64;
              let _37___mcc_h1 = (_source0).cf65;
              let _38___mcc_h2 = (_source0).cf66;
              let _39___mcc_h3 = (_source0).cf67;
              let _40_cf67 = _39___mcc_h3;
              let _41_cf66 = _38___mcc_h2;
              let _42_cf65 = _37___mcc_h1;
              let _43_cf64 = _36___mcc_h0;
              let _44_v18;
              _44_v18 = _dafny.Seq.of(_41_cf66);
              let _45_v19;
              _45_v19 = _dafny.Map.Empty.slice().updateUnsafe(_14_v1,_44_v18);
              _45_v19 = (_45_v19).update(_14_v1, _dafny.Seq.of(_40_cf67, _13_v0));
              let _46_v20;
              _46_v20 = _dafny.Set.fromElements(_13_v0);
              _40_cf67 = (_46_v20).IsProperSubsetOf((_46_v20).Difference(_46_v20));
              let _47_v21;
              _47_v21 = _dafny.Set.fromElements(_module.__default.fm24(globalState));
              let _48_v22;
              _48_v22 = _dafny.MultiSet.fromElements(_14_v1);
              let _49_v23;
              _49_v23 = _48_v22;
              let _50_v24;
              _50_v24 = _dafny.Seq.UnicodeFromString("kjr");
              let _51_v26;
              _51_v26 = _dafny.Map.Empty.slice().updateUnsafe(_16_v3,_47_v21);
              let _52_v28;
              _52_v28 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm24(globalState),_14_v1);
              let _53_v30;
              let _nw3 = Array((new BigNumber(15)).toNumber());
              _nw3[(_dafny.ZERO).toNumber()] = _47_v21;
              _nw3[(_dafny.ONE).toNumber()] = _module.__default.fm27(_49_v23, _50_v24, globalState);
              _nw3[(new BigNumber(2)).toNumber()] = _47_v21;
              _nw3[(new BigNumber(3)).toNumber()] = _dafny.Set.fromElements(_17_v4, _17_v4);
              _nw3[(new BigNumber(4)).toNumber()] = _module.__default.fm27(_48_v22, _50_v24, globalState);
              _nw3[(new BigNumber(5)).toNumber()] = _47_v21;
              _nw3[(new BigNumber(6)).toNumber()] = _47_v21;
              _nw3[(new BigNumber(7)).toNumber()] = function () {
                let _coll6 = new _dafny.Set();
                for (const _compr_6 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-30)), ((_54_v4) => function (_55_i2) {
                  return _54_v4;
                })(_17_v4))).Elements) {
                  let _56_v25 = _compr_6;
                  if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-30)), ((_57_v4) => function (_55_i2) {
                    return _57_v4;
                  })(_17_v4)), _56_v25)) {
                    _coll6.add(_56_v25);
                  }
                }
                return _coll6;
              }();
              _nw3[(new BigNumber(8)).toNumber()] = _47_v21;
              _nw3[(new BigNumber(9)).toNumber()] = ((((_51_v26).contains(_16_v3)) ? ((_51_v26).get(_16_v3)) : (_dafny.Set.fromElements(_17_v4, _17_v4)))).Difference(_dafny.Set.fromElements(_dafny.Seq.update(_17_v4, _module.__default.safeIndex(_16_v3, new BigNumber((_17_v4).length)), _16_v3)));
              _nw3[(new BigNumber(10)).toNumber()] = _47_v21;
              _nw3[(new BigNumber(11)).toNumber()] = function () {
                let _coll7 = new _dafny.Set();
                for (const _compr_7 of (_47_v21).Elements) {
                  let _58_v27 = _compr_7;
                  if ((_47_v21).contains(_58_v27)) {
                    _coll7.add(_58_v27);
                  }
                }
                return _coll7;
              }();
              _nw3[(new BigNumber(12)).toNumber()] = _47_v21;
              _nw3[(new BigNumber(13)).toNumber()] = function () {
                let _coll8 = new _dafny.Set();
                for (const _compr_8 of (_52_v28).Keys.Elements) {
                  let _59_v29 = _compr_8;
                  if ((_52_v28).contains(_59_v29)) {
                    _coll8.add(_59_v29);
                  }
                }
                return _coll8;
              }();
              _nw3[(new BigNumber(14)).toNumber()] = _47_v21;
              _53_v30 = _nw3;
              let _index3 = _module.__default.safeIndex(new BigNumber(210), new BigNumber((_53_v30).length));
              (_53_v30)[_index3] = _module.__default.fm27(_48_v22, _50_v24, globalState);
              let _index4 = _module.__default.safeIndex(new BigNumber(210), new BigNumber((_53_v30).length));
              (_53_v30)[_index4] = _47_v21;
              (globalState).f9 = new BigNumber(162);
            } else if (_source0.is_DC29) {
              let _60___mcc_h4 = (_source0).cf68;
              let _61_cf68 = _60___mcc_h4;
              _14_v1 = _14_v1;
              (globalState).f9 = (_17_v4)[_module.__default.safeIndex((_16_v3).plus(_16_v3), new BigNumber((_17_v4).length))];
              let _62_v31;
              let _init1 = ((_63_v3, _64_v0, _65_cf68) => function (_66_i3) {
                return (_dafny.Map.Empty.slice().updateUnsafe(_module.D5.create_DC14(),_63_v3)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.D5.create_DC14(),new BigNumber((_dafny.MultiSet.fromElements(_64_v0, _65_cf68)).cardinality())));
              })(_16_v3, _13_v0, _61_cf68);
              let _nw4 = Array((new BigNumber(11)).toNumber());
              for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw4.length); _i0_1++) {
                _nw4[_i0_1] = _init1(new BigNumber(_i0_1));
              }
              _62_v31 = _nw4;
              let _67_v32;
              _67_v32 = _module.D5.create_DC14();
              let _68_v33;
              _68_v33 = _dafny.Map.Empty.slice().updateUnsafe(_67_v32,_16_v3);
              let _index5 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_62_v31).length));
              (_62_v31)[_index5] = _68_v33;
              let _index6 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_62_v31).length));
              (_62_v31)[_index6] = (_68_v33).Merge(_68_v33);
              let _69_v34;
              _69_v34 = _dafny.Seq.of(true);
              let _70_v35;
              _70_v35 = _module.D1.create_DC2(_61_cf68, _13_v0, _61_cf68, _14_v1, _14_v1);
              let _71_v36;
              _71_v36 = _module.D1.create_DC2((new BigNumber((_69_v34).length)).isEqualTo(_16_v3), _13_v0, _61_cf68, _module.__default.fm14(_70_v35, new BigNumber((_69_v34).length), globalState), _14_v1);
              let _pat_let_tv3 = _13_v0;
              let _pat_let_tv4 = _14_v1;
              let _pat_let_tv5 = _61_cf68;
              _70_v35 = function (_pat_let2_0) {
                return function (_75_dt__update__tmp_h3) {
                  return function (_pat_let6_0) {
                    return function (_76_dt__update_hcf4_h0) {
                      return function (_pat_let7_0) {
                        return function (_77_dt__update_hcf3_h1) {
                          return _module.D1.create_DC2((_75_dt__update__tmp_h3).dtor_cf2, _77_dt__update_hcf3_h1, _76_dt__update_hcf4_h0, (_75_dt__update__tmp_h3).dtor_cf5, (_75_dt__update__tmp_h3).dtor_cf6);
                        }(_pat_let7_0);
                      }(_pat_let_tv5);
                    }(_pat_let6_0);
                  }(false);
                }(_pat_let2_0);
              }(function (_pat_let3_0) {
                return function (_72_dt__update__tmp_h2) {
                  return function (_pat_let4_0) {
                    return function (_73_dt__update_hcf3_h0) {
                      return function (_pat_let5_0) {
                        return function (_74_dt__update_hcf6_h0) {
                          return _module.D1.create_DC2((_72_dt__update__tmp_h2).dtor_cf2, _73_dt__update_hcf3_h0, (_72_dt__update__tmp_h2).dtor_cf4, (_72_dt__update__tmp_h2).dtor_cf5, _74_dt__update_hcf6_h0);
                        }(_pat_let5_0);
                      }(_pat_let_tv4);
                    }(_pat_let4_0);
                  }(_pat_let_tv3);
                }(_pat_let3_0);
              }(_70_v35));
            } else {
              let _78___mcc_h5 = (_source0).cf63;
              let _79_cf63 = _78___mcc_h5;
              let _80_v37;
              let _nw5 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
              _80_v37 = _nw5;
              (globalState).f1 = ((_13_v0) ? (_80_v37) : (_80_v37));
              let _81_v38;
              _81_v38 = _dafny.MultiSet.fromElements(_13_v0);
              let _82_v39;
              let _nw6 = new _module.C2();
              _nw6.__ctor(true, _81_v38, _13_v0);
              _82_v39 = _nw6;
              let _83_v40;
              _83_v40 = _dafny.Map.Empty.slice().updateUnsafe(_82_v39,_16_v3);
              let _84_v41;
              _84_v41 = _dafny.Set.fromElements(((_13_v0) ? (new BigNumber(((_83_v40).update(_82_v39, _16_v3)).length)) : (_16_v3)), (_16_v3).plus(_16_v3));
              _84_v41 = (_84_v41).Union(_84_v41);
              (globalState).f9 = (_16_v3).plus(_16_v3);
              let _85_v42;
              _85_v42 = _dafny.Seq.UnicodeFromString("bdgl");
              _85_v42 = _85_v42;
            }
            let _86_v43;
            _86_v43 = _dafny.MultiSet.fromElements(_13_v0, false, _13_v0, false);
            let _87_v44;
            let _nw7 = new _module.C2();
            _nw7.__ctor(_13_v0, _86_v43, _13_v0);
            _87_v44 = _nw7;
            let _88_v45;
            _88_v45 = _dafny.MultiSet.fromElements(_87_v44, _87_v44);
            let _89_v46;
            _89_v46 = _dafny.Map.Empty.slice().updateUnsafe(_16_v3,_87_v44.f12);
            _16_v3 = _module.__default.fm1(_module.__default.fm0(_14_v1, globalState), (_dafny.ZERO).minus(new BigNumber((_88_v45).cardinality())), new BigNumber((_89_v46).length), globalState);
          }
        }
      }
      let _90_i4;
      _90_i4 = _dafny.ZERO;
      L1: {
        while (_13_v0) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_90_i4)) {
              break L1;
            }
            _90_i4 = (_90_i4).plus(_dafny.ONE);
            let _91_v47;
            _91_v47 = _dafny.Seq.of((_13_v0) === (_13_v0), _13_v0, _13_v0);
            if ((_91_v47)[_module.__default.safeIndex(_16_v3, new BigNumber((_91_v47).length))]) {
              let _92_v48;
              _92_v48 = _dafny.MultiSet.fromElements(_13_v0, _13_v0);
              let _93_v49;
              let _nw8 = new _module.C2();
              _nw8.__ctor(_13_v0, _92_v48, false);
              _93_v49 = _nw8;
              _92_v48 = ((!((_16_v3).isLessThan(_16_v3))) ? (_92_v48) : (_92_v48));
              (globalState).f9 = _16_v3;
              let _94_v50;
              _94_v50 = _dafny.Seq.UnicodeFromString("s");
              let _95_v51;
              _95_v51 = _dafny.Map.Empty.slice().updateUnsafe(_16_v3,_93_v49.f15);
              let _96_v52;
              _96_v52 = _dafny.Seq.of(_95_v51);
              let _97_v53;
              _97_v53 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_94_v50, _94_v50), _module.__default.safeIndex(_16_v3, new BigNumber((_dafny.Seq.Concat(_94_v50, _94_v50)).length)), _14_v1)).length)),(_16_v3).minus(new BigNumber(((_96_v52)[_module.__default.safeIndex(_16_v3, new BigNumber((_96_v52).length))]).length)));
              _97_v53 = (_97_v53).update(_16_v3, _16_v3);
              let _98_v54;
              let _nw9 = new _module.C1();
              _nw9.__ctor(_16_v3, _92_v48, _93_v49.f15);
              _98_v54 = _nw9;
              let _99_v55;
              _99_v55 = _dafny.Map.Empty.slice().updateUnsafe(_98_v54,_17_v4);
              let _100_v56;
              _100_v56 = _dafny.Set.fromElements(_16_v3, _16_v3, new BigNumber((_dafny.MultiSet.FromArray((((_99_v55).contains(_98_v54)) ? ((_99_v55).get(_98_v54)) : (_17_v4)))).cardinality()), new BigNumber((_dafny.Set.fromElements(_14_v1, _14_v1)).length));
              let _101_v57;
              _101_v57 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('a'.codePointAt(0)),_92_v48);
              let _102_v58;
              let _nw10 = new _module.C3();
              _nw10.__ctor(new BigNumber((_dafny.MultiSet.fromElements(_13_v0, _13_v0)).cardinality()), (_100_v56).IsProperSubsetOf(_100_v56), (((_101_v57).contains(new _dafny.CodePoint('y'.codePointAt(0)))) ? ((_101_v57).get(new _dafny.CodePoint('y'.codePointAt(0)))) : (_92_v48)), (_dafny.Set.fromElements(_16_v3)).IsProperSubsetOf(_dafny.Set.fromElements(_16_v3)));
              _102_v58 = _nw10;
            } else {
              let _103_v59;
              _103_v59 = !(_13_v0) || (true);
              _103_v59 = !(_13_v0) || (_13_v0);
              (globalState).f7 = _13_v0;
              let _104_v60;
              let _nw11 = new _module.C0();
              _nw11.__ctor(((_13_v0) ? (_13_v0) : (_13_v0)), _16_v3, _dafny.MultiSet.fromElements(true, _13_v0), !_dafny.Seq.contains(_17_v4, _16_v3));
              _104_v60 = _nw11;
              _104_v60 = _104_v60;
              let _105_v61;
              _105_v61 = _dafny.MultiSet.fromElements(_13_v0);
              let _106_v62;
              let _nw12 = new _module.C2();
              _nw12.__ctor(_13_v0, (_105_v61).Union(_dafny.MultiSet.fromElements((_104_v60).f17)), _13_v0);
              _106_v62 = _nw12;
              let _107_v63;
              _107_v63 = _dafny.Seq.UnicodeFromString("r");
              let _108_v64;
              _108_v64 = _dafny.MultiSet.fromElements(_module.__default.fm1((_104_v60).f17, (_104_v60).f18, new BigNumber((_15_v2).length), globalState));
              let _109_v65;
              _109_v65 = _dafny.Map.Empty.slice().updateUnsafe(_13_v0,_16_v3);
              let _rhs6 = _module.__default.fm9(_16_v3, !((_104_v60).f17) || ((_104_v60).fm2(_108_v64, (_dafny.ZERO).minus(_16_v3), _109_v65, _13_v0, globalState)), globalState);
              let _rhs7 = _16_v3;
              let _rhs8 = ((new BigNumber(350)).plus(_16_v3)).multipliedBy(_16_v3);
              let _lhs4 = globalState;
              _107_v63 = _rhs6;
              _lhs4.f9 = _rhs7;
              _16_v3 = _rhs8;
            }
            let _110_v66;
            let _init2 = ((_111_v0) => function (_112_i5) {
              return ((false) ? (_111_v0) : (_111_v0));
            })(_13_v0);
            let _nw13 = Array((new BigNumber(7)).toNumber());
            for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw13.length); _i0_2++) {
              _nw13[_i0_2] = _init2(new BigNumber(_i0_2));
            }
            _110_v66 = _nw13;
            let _index7 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_110_v66).length));
            (_110_v66)[_index7] = !(_13_v0);
            let _113_v67;
            _113_v67 = _dafny.Map.Empty.slice().updateUnsafe(_13_v0,_16_v3);
            let _index8 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_110_v66).length));
            (_110_v66)[_index8] = ((_113_v67).Merge(_113_v67)).contains(true);
            (globalState).f6 = _91_v47;
            let _114_v68;
            _114_v68 = _dafny.MultiSet.fromElements(_14_v1);
            let _115_v69;
            _115_v69 = _dafny.Seq.of(_14_v1);
            let _index9 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_110_v66).length));
            (_110_v66)[_index9] = (_dafny.MultiSet.FromArray(_115_v69)).IsProperSubsetOf(_114_v68);
          }
        }
      }
      let _116_v70;
      _116_v70 = _dafny.Seq.UnicodeFromString("dslsfkrf");
      let _117_v71;
      _117_v71 = _dafny.Map.Empty.slice().updateUnsafe(_116_v70,new BigNumber((_116_v70).length));
      let _118_v72;
      _118_v72 = _dafny.Set.fromElements(_module.__default.fm0(_14_v1, globalState));
      let _119_v73;
      _119_v73 = _dafny.MultiSet.fromElements(_13_v0);
      let _120_v74;
      let _nw14 = Array((new BigNumber(26)).toNumber());
      _nw14[(_dafny.ZERO).toNumber()] = _13_v0;
      _nw14[(_dafny.ONE).toNumber()] = _13_v0;
      _nw14[(new BigNumber(2)).toNumber()] = (true) === (_13_v0);
      _nw14[(new BigNumber(3)).toNumber()] = _13_v0;
      _nw14[(new BigNumber(4)).toNumber()] = _13_v0;
      _nw14[(new BigNumber(5)).toNumber()] = _13_v0;
      _nw14[(new BigNumber(6)).toNumber()] = !(_117_v71).equals((_117_v71).update(_116_v70, _16_v3));
      _nw14[(new BigNumber(7)).toNumber()] = _13_v0;
      _nw14[(new BigNumber(8)).toNumber()] = !(_118_v72).contains(_13_v0);
      _nw14[(new BigNumber(9)).toNumber()] = false;
      _nw14[(new BigNumber(10)).toNumber()] = _13_v0;
      _nw14[(new BigNumber(11)).toNumber()] = !(_13_v0);
      _nw14[(new BigNumber(12)).toNumber()] = (_118_v72).IsDisjointFrom(_118_v72);
      _nw14[(new BigNumber(13)).toNumber()] = true;
      _nw14[(new BigNumber(14)).toNumber()] = _module.__default.fm0(_14_v1, globalState);
      _nw14[(new BigNumber(15)).toNumber()] = true;
      _nw14[(new BigNumber(16)).toNumber()] = ((_13_v0) ? (_13_v0) : (false));
      _nw14[(new BigNumber(17)).toNumber()] = _13_v0;
      _nw14[(new BigNumber(18)).toNumber()] = _13_v0;
      _nw14[(new BigNumber(19)).toNumber()] = _13_v0;
      _nw14[(new BigNumber(20)).toNumber()] = (new BigNumber((_119_v73).cardinality())).isEqualTo(new BigNumber((_116_v70).length));
      _nw14[(new BigNumber(21)).toNumber()] = (_13_v0);
      _nw14[(new BigNumber(22)).toNumber()] = _13_v0;
      _nw14[(new BigNumber(23)).toNumber()] = _13_v0;
      _nw14[(new BigNumber(24)).toNumber()] = _13_v0;
      _nw14[(new BigNumber(25)).toNumber()] = _13_v0;
      _120_v74 = _nw14;
      let _121_v75;
      _121_v75 = _dafny.Seq.of(_120_v74, _120_v74, ((_13_v0) ? (_120_v74) : (_120_v74)));
      let _122_v76;
      _122_v76 = _dafny.MultiSet.fromElements(_16_v3, _16_v3, _16_v3, _module.__default.fm1(_13_v0, _16_v3, new BigNumber(496), globalState));
      let _rhs9 = (_121_v75)[_module.__default.safeIndex(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(792)), ((_123_v70) => function (_124_i6) {
        return (_123_v70)[_module.__default.safeIndex(new BigNumber(877), new BigNumber((_123_v70).length))];
      })(_116_v70))).length), _16_v3), new BigNumber((_121_v75).length))];
      let _rhs10 = _120_v74;
      let _rhs11 = ((_dafny.Seq.IsProperPrefixOf(_116_v70, _116_v70)) ? (!(_13_v0)) : ((_122_v76).contains(_16_v3)));
      _120_v74 = _rhs9;
      _120_v74 = _rhs10;
      _13_v0 = _rhs11;
      let _125_v77;
      let _nw15 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _125_v77 = _nw15;
      let _126_v78;
      _126_v78 = _module.D5.create_DC13(_16_v3, _13_v0, _16_v3, _125_v77, _13_v0);
      let _hi0 = (_126_v78).dtor_cf30;
      for (let _127_i7 = _16_v3; _127_i7.isLessThan(_hi0); _127_i7 = _127_i7.plus(_dafny.ONE)) {
        let _128_v79;
        _128_v79 = (_dafny.MultiSet.fromElements(_14_v1)).Intersect(_dafny.MultiSet.fromElements(new _dafny.CodePoint('g'.codePointAt(0))));
        let _source1 = _128_v79;
        let _129___mcc_h6 = _source1;
        let _130_cf77 = _129___mcc_h6;
        let _131_v80;
        _131_v80 = _dafny.Seq.of(true);
        let _132_v81;
        _132_v81 = _dafny.Seq.of(_13_v0, _13_v0, false, (_131_v80)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("jcvbfjwrl")).length)), new BigNumber((_131_v80).length))]);
        (globalState).f9 = (((_122_v76).contains(_127_i7)) ? (_16_v3) : ((_127_i7).plus(new BigNumber((_132_v81).length))));
        let _133_v82;
        let _init3 = ((_134_v3) => function (_135_i8) {
          return (_135_i8).plus(_134_v3);
        })(_16_v3);
        let _nw16 = Array((_dafny.ONE).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw16.length); _i0_3++) {
          _nw16[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _133_v82 = _nw16;
        let _index10 = _module.__default.safeIndex(new BigNumber(618), new BigNumber((_133_v82).length));
        (_133_v82)[_index10] = (_16_v3).multipliedBy(_16_v3);
        let _136_v83;
        _136_v83 = _dafny.MultiSet.fromElements(_133_v82, _133_v82);
        let _137_v84;
        _137_v84 = _dafny.Set.fromElements(_127_i7, _127_i7, _16_v3);
        let _138_v85;
        _138_v85 = _module.D2.create_DC5((_dafny.ZERO).minus(_127_i7), new BigNumber((_136_v83).cardinality()), new BigNumber((_137_v84).length), new BigNumber(((_dafny.MultiSet.FromArray(_131_v80)).update(_13_v0, _module.__default.abs(_module.__default.fm1(_13_v0, _127_i7, _16_v3, globalState)))).cardinality()), _132_v81);
        let _index11 = _module.__default.safeIndex(new BigNumber(618), new BigNumber((_133_v82).length));
        let _rhs12 = _133_v82;
        let _rhs13 = _module.__default.fm1(_13_v0, (_16_v3).multipliedBy(_16_v3), _16_v3, globalState);
        let _rhs14 = _138_v85;
        let _lhs5 = globalState;
        let _lhs6 = _133_v82;
        let _lhs7 = _module.__default.safeIndex(new BigNumber(618), new BigNumber((_133_v82).length));
        _lhs5.f4 = _rhs12;
        _lhs6[_lhs7] = _rhs13;
        _138_v85 = _rhs14;
        _14_v1 = ((true) ? (_14_v1) : (_14_v1));
        let _139_v86;
        _139_v86 = _dafny.Map.Empty.slice().updateUnsafe(_14_v1,_16_v3);
        _139_v86 = (_139_v86).update(_14_v1, _module.__default.fm1(_13_v0, (_133_v82)[_module.__default.safeIndex(new BigNumber(618), new BigNumber((_133_v82).length))], (_dafny.ZERO).minus((_133_v82)[_module.__default.safeIndex(new BigNumber(618), new BigNumber((_133_v82).length))]), globalState));
        _13_v0 = false;
        let _140_v87;
        _140_v87 = _dafny.Seq.of(_dafny.Seq.of(_13_v0));
        let _141_v88;
        _141_v88 = _dafny.Map.Empty.slice().updateUnsafe((_140_v87)[_module.__default.safeIndex(_16_v3, new BigNumber((_140_v87).length))],_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(529)), ((_142_v1) => function (_143_i9) {
          return _142_v1;
        })(_14_v1)), _116_v70, _116_v70, _dafny.Seq.UnicodeFromString("bultv"), _dafny.Seq.UnicodeFromString("yvdibk")), _dafny.Seq.of(_116_v70, _116_v70))));
        let _144_v89;
        _144_v89 = _dafny.Seq.of(_13_v0);
        let _145_v90;
        _145_v90 = _dafny.MultiSet.fromElements(_116_v70, _116_v70, _116_v70);
        _141_v88 = (_141_v88).update(_144_v89, _145_v90);
        if (!(!(_13_v0))) {
          let _146_v91;
          let _nw17 = new _module.C4();
          _nw17.__ctor((_dafny.MultiSet.FromArray(_144_v89)).Union(_119_v73), !(!(_13_v0)));
          _146_v91 = _nw17;
          (globalState).f9 = _module.__default.safeDivisionInt(_127_i7, _127_i7);
          _116_v70 = _116_v70;
          let _147_v92;
          _147_v92 = _dafny.Set.fromElements(_127_i7);
          let _148_v93;
          _148_v93 = _module.D3.create_DC8((_dafny.ZERO).minus((_dafny.ZERO).minus(_16_v3)), _127_i7);
          let _149_v94;
          _149_v94 = _dafny.Map.Empty.slice().updateUnsafe((_147_v92).contains(new BigNumber((_147_v92).length)),_148_v93);
          _149_v94 = (_149_v94).update(_13_v0, _148_v93);
          let _150_v95;
          _150_v95 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("o"),_dafny.Seq.UnicodeFromString("uli"));
          let _151_v96;
          _151_v96 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm29((_144_v89)[_module.__default.safeIndex(new BigNumber((_147_v92).length), new BigNumber((_144_v89).length))], _16_v3, globalState),new BigNumber((_dafny.Set.fromElements(new BigNumber(917), new BigNumber((_17_v4).length), (_146_v91).fm3(_13_v0, _150_v95, _13_v0, _127_i7, globalState))).length));
          let _152_v97;
          _152_v97 = _dafny.Map.Empty.slice().updateUnsafe(_13_v0,_127_i7);
          let _153_v98;
          _153_v98 = _dafny.Map.Empty.slice().updateUnsafe(_127_i7,_152_v97);
          let _154_v99;
          _154_v99 = _dafny.Map.Empty.slice().updateUnsafe(_13_v0,_13_v0);
          _151_v96 = (_151_v96).update((_153_v98).Merge(_dafny.Map.Empty.slice().updateUnsafe(_16_v3,_152_v97)), (new BigNumber((_154_v99).length)).multipliedBy(_127_i7));
        } else {
          let _155_v100;
          let _nw18 = Array((new BigNumber(2)).toNumber());
          _155_v100 = _nw18;
          let _156_v101;
          _156_v101 = _module.D9.create_DC23(_dafny.Map.Empty.slice().updateUnsafe(_13_v0,_155_v100));
          let _157_v102;
          _157_v102 = _dafny.Map.Empty.slice().updateUnsafe(_16_v3,_dafny.Map.Empty.slice().updateUnsafe(_13_v0,_155_v100));
          let _158_v103;
          _158_v103 = _dafny.Map.Empty.slice().updateUnsafe(_13_v0,_155_v100);
          let _pat_let_tv6 = _157_v102;
          let _pat_let_tv7 = _16_v3;
          let _pat_let_tv8 = _158_v103;
          let _pat_let_tv9 = _157_v102;
          let _pat_let_tv10 = _16_v3;
          let _159_v104;
          _159_v104 = _dafny.Set.fromElements(function (_pat_let8_0) {
            return function (_160_dt__update__tmp_h6) {
              return function (_pat_let9_0) {
                return function (_161_dt__update_hcf55_h0) {
                  return _module.D9.create_DC23(_161_dt__update_hcf55_h0);
                }(_pat_let9_0);
              }((((_pat_let_tv9).contains(_pat_let_tv10)) ? ((_pat_let_tv6).get(_pat_let_tv7)) : (_pat_let_tv8)));
            }(_pat_let8_0);
          }(_156_v101), _156_v101, ((true) ? (_156_v101) : (_module.D9.create_DC23(_158_v103))), _156_v101, _156_v101);
          _159_v104 = (_159_v104).Union(_159_v104);
          let _index12 = _module.__default.safeIndex(new BigNumber(705), new BigNumber((_120_v74).length));
          (_120_v74)[_index12] = _13_v0;
          let _162_v105;
          _162_v105 = _dafny.Map.Empty.slice().updateUnsafe(!(((_13_v0) ? (_13_v0) : (!(_13_v0)))),!((_13_v0) && (_13_v0)));
          let _index13 = _module.__default.safeIndex(new BigNumber(705), new BigNumber((_120_v74).length));
          (_120_v74)[_index13] = (((_162_v105).contains(((!(_13_v0)) ? (_13_v0) : (_13_v0)))) ? ((_162_v105).get(((!(_13_v0)) ? (_13_v0) : (_13_v0)))) : (_13_v0));
          let _163_v106;
          _163_v106 = _module.D9.create_DC24(_116_v70, _13_v0, _16_v3);
          let _164_v107;
          let _nw19 = Array((new BigNumber(20)).toNumber());
          _nw19[(_dafny.ZERO).toNumber()] = (_120_v74)[_module.__default.safeIndex(new BigNumber(705), new BigNumber((_120_v74).length))];
          _nw19[(_dafny.ONE).toNumber()] = _13_v0;
          _nw19[(new BigNumber(2)).toNumber()] = (_120_v74)[_module.__default.safeIndex(new BigNumber(705), new BigNumber((_120_v74).length))];
          _nw19[(new BigNumber(3)).toNumber()] = false;
          _nw19[(new BigNumber(4)).toNumber()] = (_120_v74)[_module.__default.safeIndex(new BigNumber(705), new BigNumber((_120_v74).length))];
          _nw19[(new BigNumber(5)).toNumber()] = (_163_v106).dtor_cf57;
          _nw19[(new BigNumber(6)).toNumber()] = (_120_v74)[_module.__default.safeIndex(new BigNumber(705), new BigNumber((_120_v74).length))];
          _nw19[(new BigNumber(7)).toNumber()] = !(_13_v0);
          _nw19[(new BigNumber(8)).toNumber()] = (_144_v89)[_module.__default.safeIndex(_16_v3, new BigNumber((_144_v89).length))];
          _nw19[(new BigNumber(9)).toNumber()] = (_120_v74)[_module.__default.safeIndex(new BigNumber(705), new BigNumber((_120_v74).length))];
          _nw19[(new BigNumber(10)).toNumber()] = (_120_v74)[_module.__default.safeIndex(new BigNumber(705), new BigNumber((_120_v74).length))];
          _nw19[(new BigNumber(11)).toNumber()] = (_120_v74)[_module.__default.safeIndex(new BigNumber(705), new BigNumber((_120_v74).length))];
          _nw19[(new BigNumber(12)).toNumber()] = (_120_v74)[_module.__default.safeIndex(new BigNumber(705), new BigNumber((_120_v74).length))];
          _nw19[(new BigNumber(13)).toNumber()] = (_120_v74)[_module.__default.safeIndex(new BigNumber(705), new BigNumber((_120_v74).length))];
          _nw19[(new BigNumber(14)).toNumber()] = _13_v0;
          _nw19[(new BigNumber(15)).toNumber()] = _13_v0;
          _nw19[(new BigNumber(16)).toNumber()] = _13_v0;
          _nw19[(new BigNumber(17)).toNumber()] = _13_v0;
          _nw19[(new BigNumber(18)).toNumber()] = _13_v0;
          _nw19[(new BigNumber(19)).toNumber()] = _13_v0;
          _164_v107 = _nw19;
          let _165_v108;
          _165_v108 = _dafny.Set.fromElements(_164_v107, _164_v107);
          let _166_v109;
          _166_v109 = _dafny.Map.Empty.slice().updateUnsafe(_127_i7,new BigNumber((_162_v105).length));
          let _167_v110;
          _167_v110 = _dafny.Map.Empty.slice().updateUnsafe(_13_v0,_166_v109);
          let _168_v111;
          _168_v111 = _dafny.Map.Empty.slice().updateUnsafe(!(new BigNumber((_162_v105).length)).isEqualTo(new BigNumber((_165_v108).length)),((_167_v110).update((_120_v74)[_module.__default.safeIndex(new BigNumber(705), new BigNumber((_120_v74).length))], _dafny.Map.Empty.slice().updateUnsafe(_16_v3,new BigNumber((_dafny.Seq.of(new BigNumber(305))).length)))).Merge(_167_v110));
          _168_v111 = (_168_v111).update((((_144_v89)[_module.__default.safeIndex(_127_i7, new BigNumber((_144_v89).length))]) ? ((_120_v74)[_module.__default.safeIndex(new BigNumber(705), new BigNumber((_120_v74).length))]) : (_13_v0)), (_167_v110).Merge(_dafny.Map.Empty.slice().updateUnsafe(_13_v0,_166_v109)));
          let _169_v112;
          let _nw20 = new _module.C1();
          _nw20.__ctor(new BigNumber(300), _dafny.MultiSet.fromElements((_120_v74)[_module.__default.safeIndex(new BigNumber(705), new BigNumber((_120_v74).length))]), _13_v0);
          _169_v112 = _nw20;
          let _170_v113;
          _170_v113 = _dafny.Map.Empty.slice().updateUnsafe(true,_16_v3);
          let _171_v114;
          let _nw21 = Array((new BigNumber(10)).toNumber());
          _nw21[(_dafny.ZERO).toNumber()] = new BigNumber((_170_v113).length);
          _nw21[(_dafny.ONE).toNumber()] = _16_v3;
          _nw21[(new BigNumber(2)).toNumber()] = _16_v3;
          _nw21[(new BigNumber(3)).toNumber()] = new BigNumber((_144_v89).length);
          _nw21[(new BigNumber(4)).toNumber()] = _127_i7;
          _nw21[(new BigNumber(5)).toNumber()] = new BigNumber(93);
          _nw21[(new BigNumber(6)).toNumber()] = (_17_v4)[_module.__default.safeIndex(_169_v112.f16, new BigNumber((_17_v4).length))];
          _nw21[(new BigNumber(7)).toNumber()] = _169_v112.f16;
          _nw21[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(_16_v3);
          _nw21[(new BigNumber(9)).toNumber()] = new BigNumber((_17_v4).length);
          _171_v114 = _nw21;
          let _172_v115;
          _172_v115 = _dafny.Seq.of(_171_v114);
          let _173_v116;
          let _nw22 = Array((new BigNumber(17)).toNumber());
          _nw22[(_dafny.ZERO).toNumber()] = _164_v107;
          _nw22[(_dafny.ONE).toNumber()] = _164_v107;
          _nw22[(new BigNumber(2)).toNumber()] = _164_v107;
          _nw22[(new BigNumber(3)).toNumber()] = _164_v107;
          _nw22[(new BigNumber(4)).toNumber()] = _164_v107;
          _nw22[(new BigNumber(5)).toNumber()] = _164_v107;
          _nw22[(new BigNumber(6)).toNumber()] = _164_v107;
          _nw22[(new BigNumber(7)).toNumber()] = _164_v107;
          _nw22[(new BigNumber(8)).toNumber()] = _164_v107;
          _nw22[(new BigNumber(9)).toNumber()] = _164_v107;
          _nw22[(new BigNumber(10)).toNumber()] = _164_v107;
          _nw22[(new BigNumber(11)).toNumber()] = _164_v107;
          _nw22[(new BigNumber(12)).toNumber()] = _164_v107;
          _nw22[(new BigNumber(13)).toNumber()] = _164_v107;
          _nw22[(new BigNumber(14)).toNumber()] = _164_v107;
          _nw22[(new BigNumber(15)).toNumber()] = _164_v107;
          _nw22[(new BigNumber(16)).toNumber()] = _164_v107;
          _173_v116 = _nw22;
          let _174_v117;
          _174_v117 = _module.D8.create_DC21(!(_13_v0), _173_v116, _16_v3, _164_v107, _16_v3);
          let _rhs15 = _169_v112;
          let _rhs16 = ((_dafny.Seq.IsPrefixOf(_dafny.Seq.update(_172_v115, _module.__default.safeIndex((_dafny.ZERO).minus(_169_v112.f16), new BigNumber((_172_v115).length)), _171_v114), _172_v115)) ? ((_174_v117).dtor_cf52) : (_164_v107));
          let _rhs17 = (_120_v74)[_module.__default.safeIndex(new BigNumber(705), new BigNumber((_120_v74).length))];
          let _lhs8 = globalState;
          _169_v112 = _rhs15;
          _164_v107 = _rhs16;
          _lhs8.f8 = _rhs17;
          let _index14 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_171_v114).length));
          (_171_v114)[_index14] = _169_v112.f16;
          let _index15 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_171_v114).length));
          (_171_v114)[_index15] = _127_i7;
        }
      }
      return;
    }
    static Main(__noArgsParameter) {
      let _175_v1;
      _175_v1 = new BigNumber(-351);
      let _176_v2;
      _176_v2 = _dafny.Seq.of(_175_v1, new BigNumber(231), _175_v1);
      let _177_v3;
      _177_v3 = new _dafny.CodePoint('p'.codePointAt(0));
      let _178_v4;
      _178_v4 = _dafny.Map.Empty.slice().updateUnsafe(function () {
        let _coll9 = new _dafny.Map();
        for (const _compr_9 of (_176_v2).Elements) {
          let _179_v0 = _compr_9;
          if (_dafny.Seq.contains(_176_v2, _179_v0)) {
            _coll9.push([(_179_v0).minus(_175_v1),false]);
          }
        }
        return _coll9;
      }(),_177_v3);
      let _180_v5;
      let _init4 = ((_181_v1) => function (_182_i0) {
        return _module.__default.safeModuloInt(_182_i0, _181_v1);
      })(_175_v1);
      let _nw23 = Array((new BigNumber(29)).toNumber());
      for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw23.length); _i0_4++) {
        _nw23[_i0_4] = _init4(new BigNumber(_i0_4));
      }
      _180_v5 = _nw23;
      let _183_v6;
      _183_v6 = false;
      let _184_v7;
      _184_v7 = _dafny.Seq.of(_183_v6, true);
      let _185_globalState;
      let _nw24 = new _module.GlobalState();
      _nw24.__ctor(_178_v4, _180_v5, true, false, _180_v5, new BigNumber(724), _184_v7, true, true, new BigNumber(-663), false);
      _185_globalState = _nw24;
      _module.__default.m0(_185_globalState);
      let _186_i1;
      _186_i1 = _dafny.ZERO;
      L2: {
        while (((false) ? (_183_v6) : (!(!(!(_183_v6)))))) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_186_i1)) {
              break L2;
            }
            _186_i1 = (_186_i1).plus(_dafny.ONE);
            (_185_globalState).f8 = _module.__default.fm0(_177_v3, _185_globalState);
            let _187_v8;
            _187_v8 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(_175_v1, _175_v1),_module.__default.fm0(_177_v3, _185_globalState));
            let _188_v10;
            _188_v10 = _dafny.Map.Empty.slice().updateUnsafe(_183_v6,_175_v1);
            _187_v8 = function () {
              let _coll10 = new _dafny.Map();
              for (const _compr_10 of (_dafny.Seq.of(_module.__default.fm1(_183_v6, _175_v1, new BigNumber((_188_v10).length), _185_globalState))).Elements) {
                let _189_v9 = _compr_10;
                if (_dafny.Seq.contains(_dafny.Seq.of(_module.__default.fm1(_183_v6, _175_v1, new BigNumber((_188_v10).length), _185_globalState)), _189_v9)) {
                  _coll10.push([_module.__default.safeDivisionInt(_189_v9, new BigNumber((_dafny.MultiSet.fromElements(_175_v1, new BigNumber((_dafny.Seq.UnicodeFromString("kengs")).length), new BigNumber((function () {
                    let _coll11 = new _dafny.Set();
                    for (const _compr_11 of (_187_v8).Keys.Elements) {
                      let _190_v11 = _compr_11;
                      if ((_187_v8).contains(_190_v11)) {
                        _coll11.add((_190_v11).minus(new BigNumber(206)));
                      }
                    }
                    return _coll11;
                  }()).length))).cardinality())),(_dafny.MultiSet.fromElements(_183_v6, _183_v6)).IsProperSubsetOf(_dafny.MultiSet.fromElements(_183_v6, _183_v6))]);
                }
              }
              return _coll10;
            }();
            let _191_v12;
            _191_v12 = _dafny.MultiSet.fromElements(_183_v6);
            let _192_v13;
            let _nw25 = new _module.C4();
            _nw25.__ctor(_191_v12, _183_v6);
            _192_v13 = _nw25;
            if (_183_v6) {
              let _index16 = _module.__default.safeIndex(new BigNumber(968), new BigNumber((_180_v5).length));
              (_180_v5)[_index16] = (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(284)), ((_193_v3) => function (_194_i2) {
                return _193_v3;
              })(_177_v3))).length)).minus(_175_v1);
              let _index17 = _module.__default.safeIndex(new BigNumber(968), new BigNumber((_180_v5).length));
              (_180_v5)[_index17] = _175_v1;
              let _index18 = _module.__default.safeIndex(new BigNumber(968), new BigNumber((_180_v5).length));
              (_180_v5)[_index18] = _175_v1;
              (_185_globalState).f9 = _175_v1;
              let _195_v14;
              let _nw26 = new _module.C4();
              _nw26.__ctor(_191_v12, _183_v6);
              _195_v14 = _nw26;
              (_185_globalState).f7 = _183_v6;
            } else {
              let _196_v15;
              _196_v15 = _dafny.Seq.UnicodeFromString("fuaink");
              let _197_v16;
              _197_v16 = _dafny.MultiSet.fromElements(_196_v15, ((_183_v6) ? (_196_v15) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(217)), ((_198_v3) => function (_199_i3) {
                return _198_v3;
              })(_177_v3)))), _dafny.Seq.Concat(_196_v15, _dafny.Seq.UnicodeFromString("tdvl")), _196_v15, (((((_187_v8).contains((_dafny.ZERO).minus(_175_v1))) ? ((_187_v8).get((_dafny.ZERO).minus(_175_v1))) : (_183_v6))) ? (_196_v15) : ((_module.D9.create_DC25(_196_v15, _183_v6, new BigNumber((_dafny.Seq.UnicodeFromString("gsmf")).length))).dtor_cf59)));
              (_185_globalState).f9 = new BigNumber((_197_v16).cardinality());
              _188_v10 = (_188_v10).update(_183_v6, (_192_v13).fm3(_183_v6, _module.__default.fm25(_177_v3, _177_v3, !(_183_v6), _185_globalState), !(_183_v6), _175_v1, _185_globalState));
              let _200_v17;
              let _nw27 = new _module.C2();
              _nw27.__ctor(_183_v6, _191_v12, _183_v6);
              _200_v17 = _nw27;
              let _201_v18;
              let _nw28 = new _module.C4();
              _nw28.__ctor(_dafny.MultiSet.fromElements(_200_v17.f15), _200_v17.f15);
              _201_v18 = _nw28;
              let _202_v19;
              _202_v19 = _dafny.Map.Empty.slice().updateUnsafe(_175_v1,_201_v18);
              let _203_v20;
              _203_v20 = _module.D10.create_DC28(_200_v17, _200_v17.f15, _201_v18.f12, _201_v18.f12);
              let _204_v21;
              let _nw29 = Array((new BigNumber(8)).toNumber());
              _nw29[(_dafny.ZERO).toNumber()] = true;
              _nw29[(_dafny.ONE).toNumber()] = _200_v17.f15;
              _nw29[(new BigNumber(2)).toNumber()] = _200_v17.f15;
              _nw29[(new BigNumber(3)).toNumber()] = !(_175_v1).isEqualTo(new BigNumber((_202_v19).length));
              _nw29[(new BigNumber(4)).toNumber()] = (_203_v20).dtor_cf67;
              _nw29[(new BigNumber(5)).toNumber()] = true;
              _nw29[(new BigNumber(6)).toNumber()] = _200_v17.f15;
              _nw29[(new BigNumber(7)).toNumber()] = _201_v18.f12;
              _204_v21 = _nw29;
              let _index19 = _module.__default.safeIndex(new BigNumber(591), new BigNumber((_204_v21).length));
              (_204_v21)[_index19] = (new BigNumber(489)).isLessThanOrEqualTo(_175_v1);
              let _index20 = _module.__default.safeIndex(new BigNumber(591), new BigNumber((_204_v21).length));
              (_204_v21)[_index20] = _201_v18.f12;
              (_185_globalState).f7 = _200_v17.f15;
            }
          }
        }
      }
      let _205_i4;
      _205_i4 = _dafny.ZERO;
      L3: {
        while (_183_v6) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_205_i4)) {
              break L3;
            }
            _205_i4 = (_205_i4).plus(_dafny.ONE);
            let _206_v22;
            _206_v22 = _dafny.Seq.UnicodeFromString("lrpi");
            let _207_v23;
            _207_v23 = _dafny.MultiSet.fromElements(_183_v6);
            let _208_v24;
            let _nw30 = new _module.C3();
            _nw30.__ctor(new BigNumber((_206_v22).length), !(_183_v6), _207_v23, _183_v6);
            _208_v24 = _nw30;
            _208_v24 = ((_208_v24.f12) ? (_208_v24) : (_208_v24));
            let _209_v25;
            let _init5 = ((_210_v24) => function (_211_i5) {
              return _210_v24.f11;
            })(_208_v24);
            let _nw31 = Array((_dafny.ONE).toNumber());
            for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw31.length); _i0_5++) {
              _nw31[_i0_5] = _init5(new BigNumber(_i0_5));
            }
            _209_v25 = _nw31;
            let _index21 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_209_v25).length));
            (_209_v25)[_index21] = _dafny.MultiSet.fromElements(!(true), _208_v24.f12);
            let _212_v26;
            let _nw32 = new _module.C2();
            _nw32.__ctor(_208_v24.f12, _module.__default.fm22(_175_v1, !(_208_v24.f12), _185_globalState), _183_v6);
            _212_v26 = _nw32;
            let _213_v27;
            _213_v27 = _dafny.Seq.of(_212_v26, _212_v26, _212_v26);
            let _214_v28;
            let _nw33 = new _module.C0();
            _nw33.__ctor(_183_v6, _175_v1, _dafny.MultiSet.fromElements(_212_v26.f15, _208_v24.f12), _208_v24.f12);
            _214_v28 = _nw33;
            let _215_v29;
            _215_v29 = _dafny.Set.fromElements((_214_v28).f18);
            let _216_v30;
            _216_v30 = _dafny.Map.Empty.slice().updateUnsafe(_215_v29,_214_v28);
            let _217_v31;
            let _nw34 = Array((new BigNumber(22)).toNumber());
            _nw34[(_dafny.ZERO).toNumber()] = _214_v28;
            _nw34[(_dafny.ONE).toNumber()] = _214_v28;
            _nw34[(new BigNumber(2)).toNumber()] = _214_v28;
            _nw34[(new BigNumber(3)).toNumber()] = _214_v28;
            _nw34[(new BigNumber(4)).toNumber()] = _214_v28;
            _nw34[(new BigNumber(5)).toNumber()] = _214_v28;
            _nw34[(new BigNumber(6)).toNumber()] = _214_v28;
            _nw34[(new BigNumber(7)).toNumber()] = _214_v28;
            _nw34[(new BigNumber(8)).toNumber()] = (((_216_v30).contains(_dafny.Set.fromElements(_175_v1, (_214_v28).f18))) ? ((_216_v30).get(_dafny.Set.fromElements(_175_v1, (_214_v28).f18))) : (_214_v28));
            _nw34[(new BigNumber(9)).toNumber()] = _214_v28;
            _nw34[(new BigNumber(10)).toNumber()] = _214_v28;
            _nw34[(new BigNumber(11)).toNumber()] = _214_v28;
            _nw34[(new BigNumber(12)).toNumber()] = _214_v28;
            _nw34[(new BigNumber(13)).toNumber()] = _214_v28;
            _nw34[(new BigNumber(14)).toNumber()] = _214_v28;
            _nw34[(new BigNumber(15)).toNumber()] = _214_v28;
            _nw34[(new BigNumber(16)).toNumber()] = _214_v28;
            _nw34[(new BigNumber(17)).toNumber()] = _214_v28;
            _nw34[(new BigNumber(18)).toNumber()] = _214_v28;
            _nw34[(new BigNumber(19)).toNumber()] = _214_v28;
            _nw34[(new BigNumber(20)).toNumber()] = _214_v28;
            _nw34[(new BigNumber(21)).toNumber()] = _214_v28;
            _217_v31 = _nw34;
            let _218_v32;
            _218_v32 = _dafny.Map.Empty.slice().updateUnsafe(_183_v6,_217_v31);
            let _219_v33;
            _219_v33 = _dafny.Map.Empty.slice().updateUnsafe(_module.D9.create_DC23(_218_v32),_208_v24.f12);
            let _220_v34;
            _220_v34 = _dafny.Map.Empty.slice().updateUnsafe(_212_v26.f15,_218_v32);
            let _221_v35;
            _221_v35 = _dafny.Map.Empty.slice().updateUnsafe(_212_v26.f15,_175_v1);
            let _222_v36;
            _222_v36 = _module.D9.create_DC24(_206_v22, (_214_v28).f17, _175_v1);
            let _223_v37;
            _223_v37 = _dafny.Map.Empty.slice().updateUnsafe((_214_v28).fm2(_dafny.MultiSet.fromElements(new BigNumber((_208_v24.f11).cardinality())), _175_v1, _221_v35, false, _185_globalState),(_222_v36).dtor_cf57);
            let _224_v38;
            _224_v38 = _dafny.MultiSet.fromElements(_module.__default.fm16((_214_v28).f18, _185_globalState), _223_v37);
            let _225_v39;
            _225_v39 = _dafny.MultiSet.fromElements(_177_v3, _177_v3, _177_v3, _177_v3, _177_v3);
            let _226_v40;
            _226_v40 = _dafny.MultiSet.fromElements((_214_v28).f18, new BigNumber(80), (_214_v28).f18, _175_v1, (((_225_v39).contains(_177_v3)) ? ((_225_v39).get(_177_v3)) : ((_214_v28).f18)));
            let _index22 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_209_v25).length));
            let _rhs18 = (((_219_v33).contains(_module.D9.create_DC23((((_220_v34).contains((_214_v28).f17)) ? ((_220_v34).get((_214_v28).f17)) : (_218_v32))))) ? ((_219_v33).get(_module.D9.create_DC23((((_220_v34).contains((_214_v28).f17)) ? ((_220_v34).get((_214_v28).f17)) : (_218_v32))))) : ((_183_v6) === ((_214_v28).f17)));
            let _rhs19 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_206_v22, _dafny.Seq.UnicodeFromString("svi")), _206_v22);
            let _rhs20 = _module.__default.fm22((((_224_v38).contains(_dafny.Map.Empty.slice().updateUnsafe((_214_v28).f17,(_214_v28).f17))) ? ((_224_v38).get(_dafny.Map.Empty.slice().updateUnsafe((_214_v28).f17,(_214_v28).f17))) : ((_214_v28).f18)), _208_v24.f12, _185_globalState);
            let _rhs21 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_212_v26), _213_v27), _213_v27);
            let _rhs22 = !(!((_208_v24).fm2(_226_v40, (_214_v28).f18, _module.__default.fm23((_214_v28).f18, _185_globalState), _212_v26.f15, _185_globalState)));
            let _lhs9 = _185_globalState;
            let _lhs10 = _185_globalState;
            let _lhs11 = _209_v25;
            let _lhs12 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_209_v25).length));
            let _lhs13 = _185_globalState;
            _lhs9.f3 = _rhs18;
            _lhs10.f7 = _rhs19;
            _lhs11[_lhs12] = _rhs20;
            _213_v27 = _rhs21;
            _lhs13.f7 = _rhs22;
            let _227_v41;
            _227_v41 = _module.D6.create_DC16(_175_v1, (_225_v39).IsProperSubsetOf((_225_v39).update(_177_v3, _module.__default.abs((((_208_v24.f11).contains(_212_v26.f15)) ? ((_208_v24.f11).get(_212_v26.f15)) : (_175_v1))))), !(_183_v6));
            let _source2 = _227_v41;
            if (_source2.is_DC16) {
              let _228___mcc_h0 = (_source2).cf36;
              let _229___mcc_h1 = (_source2).cf37;
              let _230___mcc_h2 = (_source2).cf38;
              let _231_cf38 = _230___mcc_h2;
              let _232_cf37 = _229___mcc_h1;
              let _233_cf36 = _228___mcc_h0;
              _175_v1 = (_214_v28).f18;
              let _234_v42;
              _234_v42 = _dafny.Map.Empty.slice().updateUnsafe(_232_cf37,_223_v37);
              let _235_v43;
              _235_v43 = _223_v37;
              let _rhs23 = _208_v24;
              let _rhs24 = (((_module.__default.fm0(_177_v3, _185_globalState)) ? ((((((_234_v42).contains((_214_v28).f17)) ? ((_234_v42).get((_214_v28).f17)) : (_223_v37))).update(_212_v26.f15, _208_v24.f12)).update(_232_cf37, _208_v24.f12)) : (_223_v37))).Merge((_235_v43));
              _208_v24 = _rhs23;
              _223_v37 = _rhs24;
              let _236_v44;
              let _nw35 = Array((new BigNumber(19)).toNumber());
              _nw35[(_dafny.ZERO).toNumber()] = (_214_v28).f17;
              _nw35[(_dafny.ONE).toNumber()] = (_214_v28).f17;
              _nw35[(new BigNumber(2)).toNumber()] = _231_cf38;
              _nw35[(new BigNumber(3)).toNumber()] = _208_v24.f12;
              _nw35[(new BigNumber(4)).toNumber()] = _212_v26.f15;
              _nw35[(new BigNumber(5)).toNumber()] = _212_v26.f15;
              _nw35[(new BigNumber(6)).toNumber()] = (_214_v28).f17;
              _nw35[(new BigNumber(7)).toNumber()] = _232_cf37;
              _nw35[(new BigNumber(8)).toNumber()] = _212_v26.f15;
              _nw35[(new BigNumber(9)).toNumber()] = _212_v26.f15;
              _nw35[(new BigNumber(10)).toNumber()] = (_214_v28).f17;
              _nw35[(new BigNumber(11)).toNumber()] = (_214_v28).f17;
              _nw35[(new BigNumber(12)).toNumber()] = _208_v24.f12;
              _nw35[(new BigNumber(13)).toNumber()] = (_214_v28).f17;
              _nw35[(new BigNumber(14)).toNumber()] = !(_212_v26.f15);
              _nw35[(new BigNumber(15)).toNumber()] = _208_v24.f12;
              _nw35[(new BigNumber(16)).toNumber()] = (_214_v28).f17;
              _nw35[(new BigNumber(17)).toNumber()] = _231_cf38;
              _nw35[(new BigNumber(18)).toNumber()] = (_212_v26).fm2(_dafny.MultiSet.fromElements((_214_v28).f18, _175_v1), (_214_v28).f18, _221_v35, _212_v26.f15, _185_globalState);
              _236_v44 = _nw35;
              let _237_v45;
              _237_v45 = _dafny.Seq.of(_236_v44, _236_v44, _236_v44, _236_v44, _236_v44);
              let _238_v46;
              _238_v46 = _dafny.Set.fromElements(_236_v44, _236_v44, (_237_v45)[_module.__default.safeIndex(new BigNumber(849), new BigNumber((_237_v45).length))]);
              _238_v46 = (_238_v46).Intersect(_238_v46);
              let _239_v47;
              _239_v47 = _module.D1.create_DC2(_231_cf38, _212_v26.f15, (_214_v28).f17, _177_v3, _177_v3);
              let _240_v48;
              _240_v48 = _dafny.Map.Empty.slice().updateUnsafe((_214_v28).f18,_module.__default.fm14(_239_v47, new BigNumber((_208_v24.f11).cardinality()), _185_globalState));
              let _241_v49;
              _241_v49 = _module.D1.create_DC2(_231_cf38, (((_223_v37).contains(_208_v24.f12)) ? ((_223_v37).get(_208_v24.f12)) : (true)), false, _177_v3, (((_240_v48).contains((_214_v28).f18)) ? ((_240_v48).get((_214_v28).f18)) : (_177_v3)));
              _177_v3 = _module.__default.fm14(_241_v49, _175_v1, _185_globalState);
            } else {
              let _242___mcc_h3 = (_source2).cf35;
              let _243_cf35 = _242___mcc_h3;
              let _244_v50;
              let _245_v51;
              let _out0;
              let _out1;
              let _outcollector0 = (_212_v26).m5(_module.__default.safeModuloInt((_214_v28).f18, (_214_v28).f18), _206_v22, _185_globalState);
              _out0 = _outcollector0[0];
              _out1 = _outcollector0[1];
              _244_v50 = _out0;
              _245_v51 = _out1;
              _206_v22 = _dafny.Seq.Concat(_206_v22, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("alxsdoim"), _206_v22));
              _245_v51 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_184_v7, _dafny.Seq.of(_183_v6, _212_v26.f15)), _dafny.Seq.Concat(_dafny.Seq.of(_208_v24.f12, _208_v24.f12, _212_v26.f15), _184_v7))).length);
              (_185_globalState).f3 = !((_214_v28).f18).isEqualTo(new BigNumber((_206_v22).length));
            }
            let _246_v52;
            let _247_v53;
            let _out2;
            let _out3;
            let _outcollector1 = (_212_v26).m5(_module.__default.fm1(_212_v26.f15, _175_v1, (_214_v28).f18, _185_globalState), _206_v22, _185_globalState);
            _out2 = _outcollector1[0];
            _out3 = _outcollector1[1];
            _246_v52 = _out2;
            _247_v53 = _out3;
          }
        }
      }
      _176_v2 = _dafny.Seq.Concat(_176_v2, _176_v2);
      (_185_globalState).f3 = _183_v6;
      let _248_v54;
      _248_v54 = _dafny.Seq.UnicodeFromString("eyxck");
      let _249_v55;
      _249_v55 = _dafny.Set.fromElements(_175_v1);
      let _250_v56;
      _250_v56 = _module.D7.create_DC18(new BigNumber((_184_v7).length), _248_v54, _183_v6, _183_v6, _249_v55);
      let _251_v57;
      _251_v57 = _dafny.Seq.of((_250_v56).dtor_cf41);
      if ((_175_v1).isLessThanOrEqualTo(new BigNumber(((_251_v57)[_module.__default.safeIndex(_175_v1, new BigNumber((_251_v57).length))]).length))) {
        let _252_v58;
        _252_v58 = _dafny.MultiSet.fromElements(_183_v6);
        let _253_v59;
        let _nw36 = new _module.C2();
        _nw36.__ctor(_183_v6, _252_v58, _183_v6);
        _253_v59 = _nw36;
        let _254_v60;
        _254_v60 = _dafny.MultiSet.fromElements(_253_v59, _253_v59, _253_v59, _253_v59, _253_v59);
        let _255_v61;
        _255_v61 = _dafny.Map.Empty.slice().updateUnsafe(_175_v1,_183_v6);
        let _256_v62;
        _256_v62 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_175_v1),_175_v1);
        let _257_v63;
        _257_v63 = _dafny.MultiSet.fromElements(_175_v1, new BigNumber((_256_v62).length), _175_v1);
        let _258_v64;
        let _nw37 = new _module.C1();
        _nw37.__ctor(new BigNumber((_dafny.Set.fromElements(_175_v1, _175_v1, new BigNumber(466), new BigNumber((_254_v60).cardinality()), new BigNumber((_255_v61).length))).length), _dafny.MultiSet.FromArray(_module.__default.fm10(_253_v59.f12, _253_v59.f12, _183_v6, _185_globalState)), (new BigNumber((_257_v63).cardinality())).isLessThan(_module.__default.fm1(_183_v6, new BigNumber((_249_v55).length), _175_v1, _185_globalState)));
        _258_v64 = _nw37;
        let _259_v66;
        let _nw38 = Array((new BigNumber(14)).toNumber());
        _nw38[(_dafny.ZERO).toNumber()] = _248_v54;
        _nw38[(_dafny.ONE).toNumber()] = _248_v54;
        _nw38[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(60)), ((_260_v54, _261_v64, _262_v1) => function (_263_i6) {
          return (_260_v54)[_module.__default.safeIndex(new BigNumber((function () {
            let _coll12 = new _dafny.Map();
            for (const _compr_12 of (_dafny.Seq.of(_261_v64.f16)).Elements) {
              let _264_v65 = _compr_12;
              if (_dafny.Seq.contains(_dafny.Seq.of(_261_v64.f16), _264_v65)) {
                _coll12.push([(_264_v65).minus(_262_v1),new BigNumber(-244)]);
              }
            }
            return _coll12;
          }()).length), new BigNumber((_260_v54).length))];
        })(_248_v54, _258_v64, _175_v1));
        _nw38[(new BigNumber(3)).toNumber()] = _248_v54;
        _nw38[(new BigNumber(4)).toNumber()] = _248_v54;
        _nw38[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("r");
        _nw38[(new BigNumber(6)).toNumber()] = _dafny.Seq.update(_248_v54, _module.__default.safeIndex(_258_v64.f16, new BigNumber((_248_v54).length)), _177_v3);
        _nw38[(new BigNumber(7)).toNumber()] = _248_v54;
        _nw38[(new BigNumber(8)).toNumber()] = _248_v54;
        _nw38[(new BigNumber(9)).toNumber()] = _248_v54;
        _nw38[(new BigNumber(10)).toNumber()] = _248_v54;
        _nw38[(new BigNumber(11)).toNumber()] = _dafny.Seq.UnicodeFromString("foxyq");
        _nw38[(new BigNumber(12)).toNumber()] = _248_v54;
        _nw38[(new BigNumber(13)).toNumber()] = _248_v54;
        _259_v66 = _nw38;
        let _265_v67;
        _265_v67 = _module.D5.create_DC13(_258_v64.f16, _183_v6, (_dafny.ZERO).minus(_258_v64.f16), _259_v66, _253_v59.f12);
        _175_v1 = (_258_v64.f16).multipliedBy((_265_v67).dtor_cf30);
        let _rhs25 = (new BigNumber((_dafny.Seq.UnicodeFromString("qtkfwgn")).length)).minus((((_257_v63).contains(new BigNumber((_255_v61).length))) ? ((_257_v63).get(new BigNumber((_255_v61).length))) : (_175_v1)));
        let _rhs26 = _175_v1;
        let _lhs14 = _258_v64;
        _lhs14.f16 = _rhs25;
        _175_v1 = _rhs26;
        let _266_v68;
        _266_v68 = _module.D9.create_DC25(_dafny.Seq.Concat(_248_v54, _dafny.Seq.UnicodeFromString("ak")), _253_v59.f12, _258_v64.f16);
        let _source3 = _266_v68;
        if (_source3.is_DC24) {
          let _267___mcc_h4 = (_source3).cf56;
          let _268___mcc_h5 = (_source3).cf57;
          let _269___mcc_h6 = (_source3).cf58;
          let _270_cf58 = _269___mcc_h6;
          let _271_cf57 = _268___mcc_h5;
          let _272_cf56 = _267___mcc_h4;
          let _273_v69;
          _273_v69 = _dafny.Map.Empty.slice().updateUnsafe(_253_v59.f12,false);
          _273_v69 = (_273_v69).update(((_dafny.ZERO).minus(new BigNumber(-496))).isLessThan(new BigNumber((_dafny.Set.fromElements(_271_cf57, _253_v59.f12)).length)), ((_253_v59.f12) ? (_253_v59.f12) : (_253_v59.f12)));
          let _274_v70;
          _274_v70 = _dafny.Map.Empty.slice().updateUnsafe(_253_v59.f12,_258_v64.f16);
          let _275_v71;
          _275_v71 = _dafny.Map.Empty.slice().updateUnsafe(_258_v64.f16,_274_v70);
          _275_v71 = (_275_v71).update(new BigNumber(12), _274_v70);
          let _index23 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_180_v5).length));
          (_180_v5)[_index23] = new BigNumber((_dafny.Seq.Concat(_272_cf56, _272_cf56)).length);
          let _index24 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_180_v5).length));
          (_180_v5)[_index24] = _258_v64.f16;
          _177_v3 = _177_v3;
        } else if (_source3.is_DC25) {
          let _276___mcc_h7 = (_source3).cf59;
          let _277___mcc_h8 = (_source3).cf60;
          let _278___mcc_h9 = (_source3).cf61;
          let _279_cf61 = _278___mcc_h9;
          let _280_cf60 = _277___mcc_h8;
          let _281_cf59 = _276___mcc_h7;
          let _282_v72;
          let _nw39 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Seq.of());
          _282_v72 = _nw39;
          let _283_v73;
          _283_v73 = _dafny.Map.Empty.slice().updateUnsafe((((_255_v61).contains(_258_v64.f16)) ? ((_255_v61).get(_258_v64.f16)) : (_253_v59.f12)),_282_v72);
          _283_v73 = (_283_v73).update(_253_v59.f12, _282_v72);
          let _284_v74;
          _284_v74 = _dafny.Map.Empty.slice().updateUnsafe(_253_v59.f12,_175_v1);
          let _285_v75;
          _285_v75 = _dafny.Map.Empty.slice().updateUnsafe((((_256_v62).contains(_258_v64.f16)) ? ((_256_v62).get(_258_v64.f16)) : (new BigNumber((_284_v74).length))),_dafny.Seq.of(_280_cf60));
          let _286_v76;
          let _nw40 = new _module.C3();
          _nw40.__ctor(new BigNumber(((((_285_v75).contains(new BigNumber(679))) ? ((_285_v75).get(new BigNumber(679))) : (_184_v7))).length), _183_v6, _module.__default.fm22(_258_v64.f16, _253_v59.f12, _185_globalState), _253_v59.f12);
          _286_v76 = _nw40;
          let _287_v77;
          _287_v77 = _dafny.Map.Empty.slice().updateUnsafe(_286_v76,(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(283)), ((_288_v3) => function (_289_i7) {
            return _288_v3;
          })(_177_v3))).length)).minus(_258_v64.f16));
          (_185_globalState).f9 = new BigNumber((_287_v77).length);
          (_185_globalState).f8 = _253_v59.f12;
          let _290_v78;
          _290_v78 = _module.D6.create_DC16(_258_v64.f16, _280_cf60, _183_v6);
          (_253_v59).f12 = (_290_v78).dtor_cf38;
        } else if (_source3.is_DC26) {
          let _291___mcc_h10 = (_source3).cf62;
          let _292_cf62 = _291___mcc_h10;
          (_185_globalState).f8 = _253_v59.f12;
          let _293_v79;
          let _nw41 = new _module.C0();
          _nw41.__ctor(!(_253_v59.f12), _module.__default.safeModuloInt(_258_v64.f16, _258_v64.f16), _252_v58, !(_183_v6));
          _293_v79 = _nw41;
          _292_cf62 = new BigNumber(-697);
          let _294_v80;
          let _nw42 = Array((new BigNumber(4)).toNumber());
          _294_v80 = _nw42;
          let _295_v81;
          _295_v81 = _dafny.Seq.of(_253_v59.f11);
          let _296_v82;
          let _nw43 = new _module.C3();
          _nw43.__ctor(_258_v64.f16, _183_v6, (_295_v81)[_module.__default.safeIndex(_175_v1, new BigNumber((_295_v81).length))], _253_v59.f12);
          _296_v82 = _nw43;
          let _index25 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_294_v80).length));
          (_294_v80)[_index25] = _296_v82;
          let _index26 = _module.__default.safeIndex(new BigNumber(648), new BigNumber((_294_v80).length));
          (_294_v80)[_index26] = _296_v82;
        } else {
          let _297___mcc_h11 = (_source3).cf55;
          let _298_cf55 = _297___mcc_h11;
          let _299_v83;
          let _nw44 = new _module.C4();
          _nw44.__ctor(_253_v59.f11, _253_v59.f12);
          _299_v83 = _nw44;
          let _300_v84;
          let _nw45 = Array((new BigNumber(21)).toNumber()).fill([]);
          _300_v84 = _nw45;
          let _301_v85;
          _301_v85 = _dafny.Map.Empty.slice().updateUnsafe(_183_v6,_300_v84);
          let _302_v86;
          _302_v86 = _dafny.Seq.of(_252_v58);
          let _303_v87;
          _303_v87 = _dafny.Map.Empty.slice().updateUnsafe((((_301_v85).contains(_183_v6)) ? ((_301_v85).get(_183_v6)) : (_300_v84)),_302_v86);
          _303_v87 = (_303_v87).update(_300_v84, _302_v86);
          let _304_v88;
          _304_v88 = _module.D8.create_DC20(_253_v59.f12, _253_v59.f12, false);
          _304_v88 = _module.D8.create_DC20(_253_v59.f12, _183_v6, (_299_v83).fm4(_258_v64.f16, _258_v64.f16, _185_globalState));
          let _305_v89;
          _305_v89 = _dafny.Set.fromElements(_253_v59.f12, true, _253_v59.f12, _183_v6, !(_253_v59.f12));
          let _306_v90;
          let _307_v91;
          let _308_v92;
          let _out4;
          let _out5;
          let _out6;
          let _outcollector2 = (_299_v83).m1(_305_v89, _185_globalState);
          _out4 = _outcollector2[0];
          _out5 = _outcollector2[1];
          _out6 = _outcollector2[2];
          _306_v90 = _out4;
          _307_v91 = _out5;
          _308_v92 = _out6;
        }
        let _rhs27 = _258_v64.f16;
        let _rhs28 = (((_258_v64.f16).isLessThanOrEqualTo(_175_v1)) ? (_258_v64.f16) : (_175_v1));
        let _rhs29 = _175_v1;
        let _rhs30 = _183_v6;
        let _rhs31 = _183_v6;
        let _lhs15 = _258_v64;
        let _lhs16 = _185_globalState;
        let _lhs17 = _185_globalState;
        let _lhs18 = _185_globalState;
        let _lhs19 = _185_globalState;
        _lhs15.f16 = _rhs27;
        _lhs16.f9 = _rhs28;
        _lhs17.f9 = _rhs29;
        _lhs18.f3 = _rhs30;
        _lhs19.f8 = _rhs31;
      } else {
        let _309_v93;
        _309_v93 = _dafny.MultiSet.fromElements(!(_183_v6), false, _183_v6);
        let _310_v94;
        _310_v94 = _dafny.Map.Empty.slice().updateUnsafe(!(_dafny.MultiSet.FromArray(_dafny.Seq.of(true))).equals(_309_v93),_176_v2);
        _310_v94 = (_310_v94).update((_184_v7)[_module.__default.safeIndex(_175_v1, new BigNumber((_184_v7).length))], _dafny.Seq.Create(_module.__default.abs(new BigNumber(226)), ((_311_v1) => function (_312_i8) {
          return _311_v1;
        })(_175_v1)));
        (_185_globalState).f8 = (_175_v1).isLessThan(_175_v1);
        let _index27 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_180_v5).length));
        (_180_v5)[_index27] = (_175_v1).multipliedBy(_175_v1);
        let _index28 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_180_v5).length));
        let _rhs32 = ((_183_v6) ? (_175_v1) : (_175_v1));
        let _rhs33 = (_176_v2)[_module.__default.safeIndex(new BigNumber(927), new BigNumber((_176_v2).length))];
        let _lhs20 = _185_globalState;
        let _lhs21 = _180_v5;
        let _lhs22 = _module.__default.safeIndex(new BigNumber(537), new BigNumber((_180_v5).length));
        _lhs20.f9 = _rhs32;
        _lhs21[_lhs22] = _rhs33;
        _module.__default.m0(_185_globalState);
        let _313_v95;
        _313_v95 = _dafny.Map.Empty.slice().updateUnsafe(true,(_180_v5)[_module.__default.safeIndex(new BigNumber(537), new BigNumber((_180_v5).length))]);
        let _314_v96;
        _314_v96 = _dafny.Map.Empty.slice().updateUnsafe((_180_v5)[_module.__default.safeIndex(new BigNumber(537), new BigNumber((_180_v5).length))],(_180_v5)[_module.__default.safeIndex(new BigNumber(537), new BigNumber((_180_v5).length))]);
        (_185_globalState).f9 = _module.__default.safeDivisionInt((((_313_v95).contains(!(true))) ? ((_313_v95).get(!(true))) : (_module.__default.fm1(_183_v6, (_180_v5)[_module.__default.safeIndex(new BigNumber(537), new BigNumber((_180_v5).length))], (_dafny.ZERO).minus(new BigNumber((_314_v96).length)), _185_globalState))), (_180_v5)[_module.__default.safeIndex(new BigNumber(537), new BigNumber((_180_v5).length))]);
      }
      (_185_globalState).f3 = (new BigNumber((((_183_v6) ? (_248_v54) : (_dafny.Seq.update(_dafny.Seq.UnicodeFromString("gul"), _module.__default.safeIndex(_175_v1, new BigNumber((_dafny.Seq.UnicodeFromString("gul")).length)), new _dafny.CodePoint('d'.codePointAt(0)))))).length)).isEqualTo(_175_v1);
      if (_183_v6) {
        let _315_v97;
        _315_v97 = _dafny.Set.fromElements(_183_v6, _183_v6);
        if ((_315_v97).IsProperSubsetOf(_315_v97)) {
          let _316_v98;
          _316_v98 = _dafny.Map.Empty.slice().updateUnsafe(_183_v6,_module.__default.fm0(_177_v3, _185_globalState));
          let _317_v99;
          _317_v99 = _dafny.Map.Empty.slice().updateUnsafe(_176_v2,_316_v98);
          let _318_v100;
          _318_v100 = _module.D10.create_DC29(_183_v6);
          let _pat_let_tv11 = _183_v6;
          let _319_v101;
          _319_v101 = _dafny.Set.fromElements(_315_v97, _module.__default.fm18(new BigNumber(((_317_v99).update(_dafny.Seq.of(new BigNumber((_251_v57).length)), _316_v98)).length), (function (_pat_let10_0) {
            return function (_320_dt__update__tmp_h0) {
              return function (_pat_let11_0) {
                return function (_321_dt__update_hcf68_h0) {
                  return _module.D10.create_DC29(_321_dt__update_hcf68_h0);
                }(_pat_let11_0);
              }(_pat_let_tv11);
            }(_pat_let10_0);
          }(_318_v100)).dtor_cf68, _185_globalState), _dafny.Set.fromElements(_183_v6, _183_v6));
          let _rhs34 = !(_module.__default.fm0(_177_v3, _185_globalState));
          let _rhs35 = (_175_v1).plus(new BigNumber((_dafny.Seq.UnicodeFromString("vomrh")).length));
          let _rhs36 = _319_v101;
          let _rhs37 = _183_v6;
          let _rhs38 = ((new BigNumber(-312)).multipliedBy(_175_v1)).isLessThan(_175_v1);
          let _lhs23 = _185_globalState;
          let _lhs24 = _185_globalState;
          let _lhs25 = _185_globalState;
          _lhs23.f3 = _rhs34;
          _175_v1 = _rhs35;
          _319_v101 = _rhs36;
          _lhs24.f7 = _rhs37;
          _lhs25.f7 = _rhs38;
          let _322_v102;
          let _init6 = ((_323_v55) => function (_324_i9) {
            return _323_v55;
          })(_249_v55);
          let _nw46 = Array((new BigNumber(20)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw46.length); _i0_6++) {
            _nw46[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _322_v102 = _nw46;
          let _index29 = _module.__default.safeIndex(new BigNumber(838), new BigNumber((_322_v102).length));
          (_322_v102)[_index29] = _249_v55;
          let _index30 = _module.__default.safeIndex(new BigNumber(838), new BigNumber((_322_v102).length));
          (_322_v102)[_index30] = _249_v55;
          let _325_v103;
          _325_v103 = _dafny.Map.Empty.slice().updateUnsafe(_175_v1,_183_v6);
          let _326_v104;
          _326_v104 = _dafny.Map.Empty.slice().updateUnsafe(_325_v103,_175_v1);
          (_185_globalState).f9 = new BigNumber(((_326_v104).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_175_v1,_183_v6),(_dafny.ZERO).minus(new BigNumber(((_322_v102)[_module.__default.safeIndex(new BigNumber(838), new BigNumber((_322_v102).length))]).length))))).length);
          _175_v1 = _175_v1;
          (_185_globalState).f7 = _183_v6;
        } else {
          let _327_v105;
          _327_v105 = _dafny.Map.Empty.slice().updateUnsafe(_315_v97,_175_v1);
          let _index31 = _module.__default.safeIndex(new BigNumber(579), new BigNumber((_180_v5).length));
          (_180_v5)[_index31] = (((_327_v105).contains(_315_v97)) ? ((_327_v105).get(_315_v97)) : (_175_v1));
          let _index32 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_180_v5).length));
          (_180_v5)[_index32] = _175_v1;
          let _328_v106;
          let _nw47 = new _module.C4();
          _nw47.__ctor(_dafny.MultiSet.fromElements(_183_v6, _183_v6), _183_v6);
          _328_v106 = _nw47;
          let _329_v107;
          _329_v107 = _dafny.MultiSet.fromElements(_328_v106, _328_v106);
          let _index33 = _module.__default.safeIndex(new BigNumber(579), new BigNumber((_180_v5).length));
          let _index34 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_180_v5).length));
          let _rhs39 = _183_v6;
          let _rhs40 = _175_v1;
          let _rhs41 = _module.__default.safeDivisionInt(new BigNumber(574), _175_v1);
          let _rhs42 = _module.__default.safeDivisionInt(_175_v1, new BigNumber(-498));
          let _rhs43 = new BigNumber(((((_dafny.Map.Empty.slice().updateUnsafe(_329_v107,_183_v6)).equals(_dafny.Map.Empty.slice().updateUnsafe(_329_v107,_183_v6))) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(543)), ((_330_v1) => function (_331_i10) {
            return _330_v1;
          })(_175_v1))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(817)), ((_332_v1) => function (_333_i11) {
            return _332_v1;
          })(_175_v1))))).length);
          let _lhs26 = _185_globalState;
          let _lhs27 = _180_v5;
          let _lhs28 = _module.__default.safeIndex(new BigNumber(579), new BigNumber((_180_v5).length));
          let _lhs29 = _180_v5;
          let _lhs30 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_180_v5).length));
          let _lhs31 = _185_globalState;
          let _lhs32 = _185_globalState;
          _lhs26.f3 = _rhs39;
          _lhs27[_lhs28] = _rhs40;
          _lhs29[_lhs30] = _rhs41;
          _lhs31.f9 = _rhs42;
          _lhs32.f9 = _rhs43;
          (_185_globalState).f9 = new BigNumber(760);
          let _334_v108;
          _334_v108 = _dafny.MultiSet.fromElements(_183_v6, _183_v6, false, _183_v6);
          let _335_v109;
          _335_v109 = _dafny.Map.Empty.slice().updateUnsafe(_175_v1,_334_v108);
          let _336_v110;
          _336_v110 = _module.D1.create_DC2(true, _183_v6, _183_v6, _177_v3, _177_v3);
          let _pat_let_tv12 = _177_v3;
          let _337_v111;
          let _nw48 = new _module.C0();
          _nw48.__ctor(_183_v6, _175_v1, (((_335_v109).contains((_180_v5)[_module.__default.safeIndex(new BigNumber(579), new BigNumber((_180_v5).length))])) ? ((_335_v109).get((_180_v5)[_module.__default.safeIndex(new BigNumber(579), new BigNumber((_180_v5).length))])) : (_dafny.MultiSet.fromElements(_183_v6, _183_v6, (function (_pat_let12_0) {
            return function (_338_dt__update__tmp_h1) {
              return function (_pat_let13_0) {
                return function (_339_dt__update_hcf6_h0) {
                  return function (_pat_let14_0) {
                    return function (_340_dt__update_hcf2_h0) {
                      return _module.D1.create_DC2(_340_dt__update_hcf2_h0, (_338_dt__update__tmp_h1).dtor_cf3, (_338_dt__update__tmp_h1).dtor_cf4, (_338_dt__update__tmp_h1).dtor_cf5, _339_dt__update_hcf6_h0);
                    }(_pat_let14_0);
                  }(false);
                }(_pat_let13_0);
              }(_pat_let_tv12);
            }(_pat_let12_0);
          }(_336_v110)).dtor_cf4))), _183_v6);
          _337_v111 = _nw48;
          let _341_v112;
          _341_v112 = _dafny.Map.Empty.slice().updateUnsafe(_337_v111,_177_v3);
          _341_v112 = (_341_v112).update(_337_v111, new _dafny.CodePoint('d'.codePointAt(0)));
          let _index35 = _module.__default.safeIndex(new BigNumber(579), new BigNumber((_180_v5).length));
          let _index36 = _module.__default.safeIndex(new BigNumber(579), new BigNumber((_180_v5).length));
          let _rhs44 = ((_337_v111).f18).multipliedBy(((_337_v111).f18).multipliedBy(new BigNumber(55)));
          let _rhs45 = _module.__default.safeDivisionInt((_180_v5)[_module.__default.safeIndex(new BigNumber(579), new BigNumber((_180_v5).length))], (((_334_v108).contains(_183_v6)) ? ((_334_v108).get(_183_v6)) : (new BigNumber(542))));
          let _rhs46 = (_dafny.ZERO).minus((_337_v111).f18);
          let _lhs33 = _180_v5;
          let _lhs34 = _module.__default.safeIndex(new BigNumber(579), new BigNumber((_180_v5).length));
          let _lhs35 = _180_v5;
          let _lhs36 = _module.__default.safeIndex(new BigNumber(579), new BigNumber((_180_v5).length));
          let _lhs37 = _185_globalState;
          _lhs33[_lhs34] = _rhs44;
          _lhs35[_lhs36] = _rhs45;
          _lhs37.f9 = _rhs46;
          let _342_v113;
          let _343_v114;
          let _344_v115;
          let _out7;
          let _out8;
          let _out9;
          let _outcollector3 = (_328_v106).m1(_module.__default.fm18(new BigNumber(269), (_337_v111).f17, _185_globalState), _185_globalState);
          _out7 = _outcollector3[0];
          _out8 = _outcollector3[1];
          _out9 = _outcollector3[2];
          _342_v113 = _out7;
          _343_v114 = _out8;
          _344_v115 = _out9;
        }
        (_185_globalState).f3 = !(_175_v1).isEqualTo(_module.__default.safeDivisionInt(_175_v1, _175_v1));
        let _345_v116;
        let _nw49 = Array((new BigNumber(9)).toNumber());
        _345_v116 = _nw49;
        let _346_v117;
        _346_v117 = _dafny.Seq.of(_345_v116);
        let _347_v118;
        let _nw50 = Array((new BigNumber(14)).toNumber());
        _nw50[(_dafny.ZERO).toNumber()] = _345_v116;
        _nw50[(_dafny.ONE).toNumber()] = _345_v116;
        _nw50[(new BigNumber(2)).toNumber()] = _345_v116;
        _nw50[(new BigNumber(3)).toNumber()] = _345_v116;
        _nw50[(new BigNumber(4)).toNumber()] = _345_v116;
        _nw50[(new BigNumber(5)).toNumber()] = (_346_v117)[_module.__default.safeIndex((_dafny.ZERO).minus(_175_v1), new BigNumber((_346_v117).length))];
        _nw50[(new BigNumber(6)).toNumber()] = _345_v116;
        _nw50[(new BigNumber(7)).toNumber()] = _345_v116;
        _nw50[(new BigNumber(8)).toNumber()] = _345_v116;
        _nw50[(new BigNumber(9)).toNumber()] = _345_v116;
        _nw50[(new BigNumber(10)).toNumber()] = _345_v116;
        _nw50[(new BigNumber(11)).toNumber()] = _345_v116;
        _nw50[(new BigNumber(12)).toNumber()] = _345_v116;
        _nw50[(new BigNumber(13)).toNumber()] = _345_v116;
        _347_v118 = _nw50;
        let _index37 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_347_v118).length));
        (_347_v118)[_index37] = ((_183_v6) ? (_345_v116) : (_345_v116));
        let _index38 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_347_v118).length));
        (_347_v118)[_index38] = _345_v116;
        let _348_v119;
        _348_v119 = _dafny.Map.Empty.slice().updateUnsafe(_177_v3,_175_v1);
        let _349_v120;
        _349_v120 = _dafny.Set.fromElements(_348_v119);
        let _350_v121;
        _350_v121 = _dafny.Map.Empty.slice().updateUnsafe((_349_v120).Intersect(_349_v120),(_249_v55).Difference(_249_v55));
        _249_v55 = (((_350_v121).contains(_dafny.Set.fromElements(_348_v119, _348_v119, _dafny.Map.Empty.slice().updateUnsafe(_177_v3,_175_v1)))) ? ((_350_v121).get(_dafny.Set.fromElements(_348_v119, _348_v119, _dafny.Map.Empty.slice().updateUnsafe(_177_v3,_175_v1)))) : (((_250_v56).dtor_cf44).Difference(_249_v55)));
        if (_module.__default.fm0(_177_v3, _185_globalState)) {
          (_185_globalState).f9 = new BigNumber(678);
          let _351_v122;
          let _nw51 = Array((new BigNumber(28)).toNumber()).fill(false);
          _351_v122 = _nw51;
          _351_v122 = _351_v122;
          let _352_v123;
          let _nw52 = Array((new BigNumber(29)).toNumber()).fill(_module.D8.Default());
          _352_v123 = _nw52;
          let _353_v124;
          let _nw53 = Array((new BigNumber(8)).toNumber());
          _nw53[(_dafny.ZERO).toNumber()] = _352_v123;
          _nw53[(_dafny.ONE).toNumber()] = _352_v123;
          _nw53[(new BigNumber(2)).toNumber()] = _352_v123;
          _nw53[(new BigNumber(3)).toNumber()] = _352_v123;
          _nw53[(new BigNumber(4)).toNumber()] = _352_v123;
          _nw53[(new BigNumber(5)).toNumber()] = _352_v123;
          _nw53[(new BigNumber(6)).toNumber()] = _352_v123;
          _nw53[(new BigNumber(7)).toNumber()] = _352_v123;
          _353_v124 = _nw53;
          let _index39 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((_353_v124).length));
          (_353_v124)[_index39] = _352_v123;
          let _index40 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((_353_v124).length));
          (_353_v124)[_index40] = _352_v123;
          let _354_v125;
          _354_v125 = _dafny.Map.Empty.slice().updateUnsafe(_177_v3,_183_v6);
          _354_v125 = (_354_v125).update(_177_v3, _183_v6);
          let _355_v126;
          let _nw54 = Array((new BigNumber(8)).toNumber()).fill([]);
          _355_v126 = _nw54;
          let _356_v127;
          _356_v127 = _module.D8.create_DC21(_183_v6, _355_v126, _175_v1, _351_v122, new BigNumber(325));
          let _357_v128;
          _357_v128 = _dafny.MultiSet.fromElements(!(_module.__default.fm0(_177_v3, _185_globalState)) || (!((_356_v127).dtor_cf49)), _183_v6, _183_v6, _183_v6);
          _357_v128 = ((_module.__default.fm22(_175_v1, _183_v6, _185_globalState)).Union(_357_v128)).Intersect((_357_v128).Intersect(_357_v128));
        } else {
          (_185_globalState).f3 = _183_v6;
          (_185_globalState).f9 = new BigNumber(-58);
          _module.__default.m0(_185_globalState);
          let _358_v129;
          let _nw55 = Array((new BigNumber(16)).toNumber());
          _358_v129 = _nw55;
          let _359_v130;
          _359_v130 = _dafny.Map.Empty.slice().updateUnsafe(_358_v129,_183_v6);
          _359_v130 = (_359_v130).update(_358_v129, _183_v6);
          _176_v2 = _dafny.Seq.of((_175_v1).multipliedBy(_175_v1), new BigNumber(975), _175_v1);
        }
      } else {
        (_185_globalState).f7 = _183_v6;
        let _360_v131;
        _360_v131 = _dafny.MultiSet.fromElements(_183_v6);
        let _361_v132;
        let _nw56 = new _module.C4();
        _nw56.__ctor(_360_v131, _183_v6);
        _361_v132 = _nw56;
        let _362_v133;
        let _nw57 = new _module.C0();
        _nw57.__ctor(_183_v6, _175_v1, (_360_v131).update(_183_v6, _module.__default.abs(_175_v1)), _183_v6);
        _362_v133 = _nw57;
        _362_v133 = _362_v133;
        let _363_v134;
        _363_v134 = _module.D8.create_DC20(_183_v6, _183_v6, _362_v133.f12);
        let _364_v135;
        let _nw58 = new _module.C0();
        _nw58.__ctor(!(_362_v133.f12), _175_v1, _362_v133.f11, (_363_v134).dtor_cf48);
        _364_v135 = _nw58;
        let _365_v136;
        _365_v136 = _dafny.Map.Empty.slice().updateUnsafe(_183_v6,_175_v1);
        let _366_v137;
        _366_v137 = _dafny.Set.fromElements(_365_v136);
        let _367_v138;
        _367_v138 = _dafny.Set.fromElements(true, (_364_v135).f17);
        let _368_v139;
        let _nw59 = new _module.C0();
        _nw59.__ctor(_183_v6, new BigNumber((_366_v137).length), _362_v133.f11, (new BigNumber((_367_v138).length)).isLessThanOrEqualTo((_364_v135).f18));
        _368_v139 = _nw59;
      }
      let _369_v140;
      _369_v140 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("muwivve"), _module.__default.safeIndex(_175_v1, new BigNumber((_dafny.Seq.UnicodeFromString("muwivve")).length)), _177_v3),(_module.__default.fm26(_175_v1, !(false), _185_globalState)).dtor_cf46);
      if ((((_369_v140).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(820)), ((_373_v3) => function (_374_i12) {
        return _373_v3;
      })(_177_v3)))) ? ((_369_v140).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(820)), ((_370_v3) => function (_371_i12) {
        return _370_v3;
      })(_177_v3)))) : ((_249_v55).IsProperSubsetOf(function () {
        let _coll13 = new _dafny.Set();
        for (const _compr_13 of _dafny.IntegerRange(new BigNumber(567), new BigNumber(942))) {
          let _372_v141 = _compr_13;
          if (((new BigNumber(567)).isLessThanOrEqualTo(_372_v141)) && ((_372_v141).isLessThan(new BigNumber(942)))) {
            _coll13.add((_372_v141).plus(_175_v1));
          }
        }
        return _coll13;
      }())))) {
        let _375_v142;
        let _init7 = ((_376_v6) => function (_377_i13) {
          return _376_v6;
        })(_183_v6);
        let _nw60 = Array((new BigNumber(7)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw60.length); _i0_7++) {
          _nw60[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _375_v142 = _nw60;
        let _378_v143;
        _378_v143 = _dafny.Seq.of(_375_v142);
        let _379_v144;
        let _nw61 = Array((new BigNumber(23)).toNumber());
        _nw61[(_dafny.ZERO).toNumber()] = _375_v142;
        _nw61[(_dafny.ONE).toNumber()] = _375_v142;
        _nw61[(new BigNumber(2)).toNumber()] = _375_v142;
        _nw61[(new BigNumber(3)).toNumber()] = (_378_v143)[_module.__default.safeIndex(_module.__default.fm1(_183_v6, _module.__default.fm1(_183_v6, _175_v1, _175_v1, _185_globalState), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(810)), ((_380_v3) => function (_381_i14) {
          return _380_v3;
        })(_177_v3))).length), _185_globalState), new BigNumber((_378_v143).length))];
        _nw61[(new BigNumber(4)).toNumber()] = _375_v142;
        _nw61[(new BigNumber(5)).toNumber()] = _375_v142;
        _nw61[(new BigNumber(6)).toNumber()] = _375_v142;
        _nw61[(new BigNumber(7)).toNumber()] = _375_v142;
        _nw61[(new BigNumber(8)).toNumber()] = _375_v142;
        _nw61[(new BigNumber(9)).toNumber()] = _375_v142;
        _nw61[(new BigNumber(10)).toNumber()] = _375_v142;
        _nw61[(new BigNumber(11)).toNumber()] = _375_v142;
        _nw61[(new BigNumber(12)).toNumber()] = _375_v142;
        _nw61[(new BigNumber(13)).toNumber()] = _375_v142;
        _nw61[(new BigNumber(14)).toNumber()] = _375_v142;
        _nw61[(new BigNumber(15)).toNumber()] = _375_v142;
        _nw61[(new BigNumber(16)).toNumber()] = _375_v142;
        _nw61[(new BigNumber(17)).toNumber()] = _375_v142;
        _nw61[(new BigNumber(18)).toNumber()] = _375_v142;
        _nw61[(new BigNumber(19)).toNumber()] = _375_v142;
        _nw61[(new BigNumber(20)).toNumber()] = _375_v142;
        _nw61[(new BigNumber(21)).toNumber()] = _375_v142;
        _nw61[(new BigNumber(22)).toNumber()] = _375_v142;
        _379_v144 = _nw61;
        let _382_v145;
        _382_v145 = _module.D8.create_DC21(_183_v6, _379_v144, _175_v1, _375_v142, _module.__default.fm1(_183_v6, new BigNumber(759), _175_v1, _185_globalState));
        let _383_v146;
        _383_v146 = _dafny.MultiSet.fromElements(_183_v6, (_382_v145).dtor_cf49);
        let _384_v147;
        let _nw62 = new _module.C2();
        _nw62.__ctor(true, _383_v146, !(_183_v6));
        _384_v147 = _nw62;
        let _385_v148;
        _385_v148 = _dafny.Map.Empty.slice().updateUnsafe(_384_v147,_375_v142);
        let _pat_let_tv13 = _379_v144;
        let _386_v149;
        let _nw63 = Array((new BigNumber(23)).toNumber());
        _nw63[(_dafny.ZERO).toNumber()] = _375_v142;
        _nw63[(_dafny.ONE).toNumber()] = _375_v142;
        _nw63[(new BigNumber(2)).toNumber()] = _375_v142;
        _nw63[(new BigNumber(3)).toNumber()] = _375_v142;
        _nw63[(new BigNumber(4)).toNumber()] = (((_385_v148).contains(_384_v147)) ? ((_385_v148).get(_384_v147)) : (_375_v142));
        _nw63[(new BigNumber(5)).toNumber()] = _375_v142;
        _nw63[(new BigNumber(6)).toNumber()] = _375_v142;
        _nw63[(new BigNumber(7)).toNumber()] = _375_v142;
        _nw63[(new BigNumber(8)).toNumber()] = (function (_pat_let15_0) {
          return function (_387_dt__update__tmp_h2) {
            return function (_pat_let16_0) {
              return function (_388_dt__update_hcf50_h0) {
                return _module.D8.create_DC21((_387_dt__update__tmp_h2).dtor_cf49, _388_dt__update_hcf50_h0, (_387_dt__update__tmp_h2).dtor_cf51, (_387_dt__update__tmp_h2).dtor_cf52, (_387_dt__update__tmp_h2).dtor_cf53);
              }(_pat_let16_0);
            }(_pat_let_tv13);
          }(_pat_let15_0);
        }(_382_v145)).dtor_cf52;
        _nw63[(new BigNumber(9)).toNumber()] = _375_v142;
        _nw63[(new BigNumber(10)).toNumber()] = _375_v142;
        _nw63[(new BigNumber(11)).toNumber()] = _375_v142;
        _nw63[(new BigNumber(12)).toNumber()] = _375_v142;
        _nw63[(new BigNumber(13)).toNumber()] = _375_v142;
        _nw63[(new BigNumber(14)).toNumber()] = _375_v142;
        _nw63[(new BigNumber(15)).toNumber()] = _375_v142;
        _nw63[(new BigNumber(16)).toNumber()] = _375_v142;
        _nw63[(new BigNumber(17)).toNumber()] = _375_v142;
        _nw63[(new BigNumber(18)).toNumber()] = _375_v142;
        _nw63[(new BigNumber(19)).toNumber()] = _375_v142;
        _nw63[(new BigNumber(20)).toNumber()] = _375_v142;
        _nw63[(new BigNumber(21)).toNumber()] = _375_v142;
        _nw63[(new BigNumber(22)).toNumber()] = _375_v142;
        _386_v149 = _nw63;
        let _index41 = _module.__default.safeIndex(new BigNumber(586), new BigNumber((_379_v144).length));
        (_379_v144)[_index41] = _375_v142;
        let _389_v150;
        _389_v150 = _dafny.Map.Empty.slice().updateUnsafe(_175_v1,_384_v147.f15);
        let _390_v151;
        _390_v151 = _dafny.Map.Empty.slice().updateUnsafe((((_389_v150).contains(_175_v1)) ? ((_389_v150).get(_175_v1)) : (_384_v147.f15)),_375_v142);
        let _391_v152;
        _391_v152 = _dafny.MultiSet.fromElements(_175_v1, _175_v1, new BigNumber((_184_v7).length));
        let _392_v153;
        _392_v153 = _dafny.Map.Empty.slice().updateUnsafe(_183_v6,_175_v1);
        let _index42 = _module.__default.safeIndex(new BigNumber(586), new BigNumber((_379_v144).length));
        (_379_v144)[_index42] = (((_390_v151).contains((_384_v147).fm2(_391_v152, _175_v1, _392_v153, !(_183_v6), _185_globalState))) ? ((_390_v151).get((_384_v147).fm2(_391_v152, _175_v1, _392_v153, !(_183_v6), _185_globalState))) : (_375_v142));
        let _arr0 = (_379_v144)[_module.__default.safeIndex(new BigNumber(586), new BigNumber((_379_v144).length))];
        let _index43 = _module.__default.safeIndex(new BigNumber(94), new BigNumber(((_379_v144)[_module.__default.safeIndex(new BigNumber(586), new BigNumber((_379_v144).length))]).length));
        _arr0[_index43] = _384_v147.f15;
        let _arr1 = (_379_v144)[_module.__default.safeIndex(new BigNumber(586), new BigNumber((_379_v144).length))];
        let _index44 = _module.__default.safeIndex(new BigNumber(94), new BigNumber(((_379_v144)[_module.__default.safeIndex(new BigNumber(586), new BigNumber((_379_v144).length))]).length));
        _arr1[_index44] = _183_v6;
        let _393_v154;
        let _nw64 = new _module.C0();
        _nw64.__ctor(((_379_v144)[_module.__default.safeIndex(new BigNumber(586), new BigNumber((_379_v144).length))])[_module.__default.safeIndex(new BigNumber(94), new BigNumber(((_379_v144)[_module.__default.safeIndex(new BigNumber(586), new BigNumber((_379_v144).length))]).length))], _175_v1, _383_v146, _183_v6);
        _393_v154 = _nw64;
        if (_384_v147.f15) {
          _183_v6 = ((_393_v154).f18).isLessThanOrEqualTo((new BigNumber((_392_v153).length)).multipliedBy((_393_v154).f18));
          (_185_globalState).f8 = ((_379_v144)[_module.__default.safeIndex(new BigNumber(586), new BigNumber((_379_v144).length))])[_module.__default.safeIndex(new BigNumber(94), new BigNumber(((_379_v144)[_module.__default.safeIndex(new BigNumber(586), new BigNumber((_379_v144).length))]).length))];
          let _394_v155;
          _394_v155 = _dafny.Set.fromElements(_177_v3, _177_v3);
          let _395_v156;
          let _nw65 = Array((new BigNumber(27)).toNumber());
          _nw65[(_dafny.ZERO).toNumber()] = _248_v54;
          _nw65[(_dafny.ONE).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-298)), ((_396_v3) => function (_397_i15) {
            return _396_v3;
          })(_177_v3));
          _nw65[(new BigNumber(2)).toNumber()] = _dafny.Seq.UnicodeFromString("ecbkysbx");
          _nw65[(new BigNumber(3)).toNumber()] = _dafny.Seq.UnicodeFromString("mekb");
          _nw65[(new BigNumber(4)).toNumber()] = _248_v54;
          _nw65[(new BigNumber(5)).toNumber()] = _248_v54;
          _nw65[(new BigNumber(6)).toNumber()] = _248_v54;
          _nw65[(new BigNumber(7)).toNumber()] = _248_v54;
          _nw65[(new BigNumber(8)).toNumber()] = _248_v54;
          _nw65[(new BigNumber(9)).toNumber()] = _248_v54;
          _nw65[(new BigNumber(10)).toNumber()] = _248_v54;
          _nw65[(new BigNumber(11)).toNumber()] = _248_v54;
          _nw65[(new BigNumber(12)).toNumber()] = _248_v54;
          _nw65[(new BigNumber(13)).toNumber()] = _248_v54;
          _nw65[(new BigNumber(14)).toNumber()] = _dafny.Seq.UnicodeFromString("cnflkujmg");
          _nw65[(new BigNumber(15)).toNumber()] = _248_v54;
          _nw65[(new BigNumber(16)).toNumber()] = _248_v54;
          _nw65[(new BigNumber(17)).toNumber()] = _248_v54;
          _nw65[(new BigNumber(18)).toNumber()] = _module.__default.fm9(new BigNumber((_394_v155).length), true, _185_globalState);
          _nw65[(new BigNumber(19)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(432)), function (_398_i16) {
            return new _dafny.CodePoint('f'.codePointAt(0));
          });
          _nw65[(new BigNumber(20)).toNumber()] = _248_v54;
          _nw65[(new BigNumber(21)).toNumber()] = _248_v54;
          _nw65[(new BigNumber(22)).toNumber()] = _248_v54;
          _nw65[(new BigNumber(23)).toNumber()] = _dafny.Seq.UnicodeFromString("c");
          _nw65[(new BigNumber(24)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(325)), ((_399_v3) => function (_400_i17) {
            return _399_v3;
          })(_177_v3));
          _nw65[(new BigNumber(25)).toNumber()] = _248_v54;
          _nw65[(new BigNumber(26)).toNumber()] = _248_v54;
          _395_v156 = _nw65;
          let _401_v157;
          _401_v157 = _module.D5.create_DC13(new BigNumber((_392_v153).length), true, _175_v1, _395_v156, (_393_v154).f17);
          let _402_v158;
          _402_v158 = _dafny.Set.fromElements(_401_v157, _401_v157);
          let _403_v159;
          _403_v159 = _dafny.Seq.of(_402_v158);
          let _404_v160;
          _404_v160 = _module.D12.create_DC31(_403_v159);
          let _405_v161;
          _405_v161 = _dafny.Seq.of(_403_v159, _dafny.Seq.of(_402_v158));
          (_384_v147).f15 = _dafny.areEqual((_404_v160).dtor_cf70, (_405_v161)[_module.__default.safeIndex((_393_v154).f18, new BigNumber((_405_v161).length))]);
          let _406_v162;
          let _nw66 = new _module.C1();
          _nw66.__ctor((_dafny.ZERO).minus((_393_v154).f18), _dafny.MultiSet.fromElements(_183_v6), ((_393_v154).f18).isEqualTo((_393_v154).f18));
          _406_v162 = _nw66;
          let _407_v163;
          _407_v163 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("ygwxn"));
          let _408_v164;
          _408_v164 = _dafny.Map.Empty.slice().updateUnsafe((_393_v154).f17,_407_v163);
          _407_v163 = ((((_408_v164).contains(_183_v6)) ? ((_408_v164).get(_183_v6)) : (_407_v163))).Difference(_407_v163);
        } else {
          _module.__default.m0(_185_globalState);
          let _index45 = _module.__default.safeIndex(new BigNumber(527), new BigNumber((_180_v5).length));
          (_180_v5)[_index45] = (_393_v154).f18;
          let _index46 = _module.__default.safeIndex(new BigNumber(527), new BigNumber((_180_v5).length));
          (_180_v5)[_index46] = (((_383_v146).contains((_384_v147).fm2(_391_v152, (_393_v154).f18, _dafny.Map.Empty.slice().updateUnsafe(true,(_393_v154).f18), _183_v6, _185_globalState))) ? ((_383_v146).get((_384_v147).fm2(_391_v152, (_393_v154).f18, _dafny.Map.Empty.slice().updateUnsafe(true,(_393_v154).f18), _183_v6, _185_globalState))) : ((_393_v154).f18));
          let _409_v165;
          _409_v165 = _dafny.MultiSet.fromElements(_177_v3);
          let _410_v166;
          _410_v166 = _dafny.Map.Empty.slice().updateUnsafe(((_379_v144)[_module.__default.safeIndex(new BigNumber(586), new BigNumber((_379_v144).length))])[_module.__default.safeIndex(new BigNumber(94), new BigNumber(((_379_v144)[_module.__default.safeIndex(new BigNumber(586), new BigNumber((_379_v144).length))]).length))],_409_v165);
          let _411_v167;
          _411_v167 = _dafny.Set.fromElements(!(_183_v6));
          let _412_v168;
          let _nw67 = Array((new BigNumber(14)).toNumber());
          _nw67[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus((_393_v154).f18);
          _nw67[(_dafny.ONE).toNumber()] = (((_393_v154).f17) ? (new BigNumber(-420)) : (new BigNumber(((((_410_v166).contains(false)) ? ((_410_v166).get(false)) : (_409_v165))).cardinality())));
          _nw67[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus((new BigNumber((_dafny.MultiSet.fromElements((_393_v154).f17, _183_v6)).cardinality())).minus((_393_v154).f18));
          _nw67[(new BigNumber(3)).toNumber()] = _175_v1;
          _nw67[(new BigNumber(4)).toNumber()] = new BigNumber((_411_v167).length);
          _nw67[(new BigNumber(5)).toNumber()] = _175_v1;
          _nw67[(new BigNumber(6)).toNumber()] = (_175_v1).minus(new BigNumber(844));
          _nw67[(new BigNumber(7)).toNumber()] = ((_393_v154).fm12(_185_globalState)).multipliedBy((_393_v154).f18);
          _nw67[(new BigNumber(8)).toNumber()] = new BigNumber(620);
          _nw67[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(610),(_393_v154).f18)).length);
          _nw67[(new BigNumber(10)).toNumber()] = (_393_v154).f18;
          _nw67[(new BigNumber(11)).toNumber()] = _module.__default.safeModuloInt(_175_v1, (_393_v154).f18);
          _nw67[(new BigNumber(12)).toNumber()] = (_180_v5)[_module.__default.safeIndex(new BigNumber(527), new BigNumber((_180_v5).length))];
          _nw67[(new BigNumber(13)).toNumber()] = _module.__default.safeDivisionInt(_175_v1, new BigNumber((_248_v54).length));
          _412_v168 = _nw67;
          (_185_globalState).f1 = _412_v168;
          let _413_v169;
          _413_v169 = _dafny.Set.fromElements(_177_v3, _177_v3);
          let _414_v170;
          _414_v170 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_413_v169).length),_391_v152);
          let _415_v171;
          _415_v171 = _dafny.Map.Empty.slice().updateUnsafe(_175_v1,_383_v146);
          let _416_v172;
          let _nw68 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _416_v172 = _nw68;
          let _417_v173;
          _417_v173 = _module.D5.create_DC13(_175_v1, false, new BigNumber((_415_v171).length), _416_v172, (_393_v154).f17);
          let _index47 = _module.__default.safeIndex(new BigNumber(527), new BigNumber((_180_v5).length));
          let _rhs47 = _dafny.Seq.Concat(_248_v54, _248_v54);
          let _rhs48 = (_393_v154).f18;
          let _rhs49 = (_384_v147).fm2(_dafny.MultiSet.fromElements((_393_v154).f18, new BigNumber(((((_414_v170).contains(new BigNumber((_249_v55).length))) ? ((_414_v170).get(new BigNumber((_249_v55).length))) : (_391_v152))).cardinality())), (_417_v173).dtor_cf32, (_module.__default.fm23((_dafny.ZERO).minus((_180_v5)[_module.__default.safeIndex(new BigNumber(527), new BigNumber((_180_v5).length))]), _185_globalState)).update(false, _175_v1), ((_379_v144)[_module.__default.safeIndex(new BigNumber(586), new BigNumber((_379_v144).length))])[_module.__default.safeIndex(new BigNumber(94), new BigNumber(((_379_v144)[_module.__default.safeIndex(new BigNumber(586), new BigNumber((_379_v144).length))]).length))], _185_globalState);
          let _rhs50 = ((_dafny.MultiSet.FromArray(_176_v2)).Intersect(_module.__default.fm13((_248_v54)[_module.__default.safeIndex((_393_v154).f18, new BigNumber((_248_v54).length))], false, _185_globalState))).Union(_391_v152);
          let _lhs38 = _180_v5;
          let _lhs39 = _module.__default.safeIndex(new BigNumber(527), new BigNumber((_180_v5).length));
          let _lhs40 = _185_globalState;
          _248_v54 = _rhs47;
          _lhs38[_lhs39] = _rhs48;
          _lhs40.f7 = _rhs49;
          _391_v152 = _rhs50;
          (_185_globalState).f9 = (_393_v154).fm12(_185_globalState);
        }
        let _418_v174;
        _418_v174 = _module.D5.create_DC14();
        let _source4 = _418_v174;
        if (_source4.is_DC13) {
          let _419___mcc_h12 = (_source4).cf30;
          let _420___mcc_h13 = (_source4).cf31;
          let _421___mcc_h14 = (_source4).cf32;
          let _422___mcc_h15 = (_source4).cf33;
          let _423___mcc_h16 = (_source4).cf34;
          let _424_cf34 = _423___mcc_h16;
          let _425_cf33 = _422___mcc_h15;
          let _426_cf32 = _421___mcc_h14;
          let _427_cf31 = _420___mcc_h13;
          let _428_cf30 = _419___mcc_h12;
          let _429_v175;
          let _nw69 = new _module.C4();
          _nw69.__ctor(_383_v146, (_249_v55).contains((_393_v154).f18));
          _429_v175 = _nw69;
          let _index48 = _module.__default.safeIndex(new BigNumber(539), new BigNumber((_180_v5).length));
          (_180_v5)[_index48] = new BigNumber(135);
          let _430_v176;
          let _nw70 = new _module.C3();
          _nw70.__ctor(_426_cf32, (_393_v154).f17, _383_v146, false);
          _430_v176 = _nw70;
          let _431_v177;
          _431_v177 = _dafny.Map.Empty.slice().updateUnsafe((_430_v176).f13,_428_cf30);
          let _432_v178;
          _432_v178 = _dafny.Map.Empty.slice().updateUnsafe(_430_v176,new BigNumber((_431_v177).length));
          let _433_v179;
          _433_v179 = _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements((((_432_v178).contains(_430_v176)) ? ((_432_v178).get(_430_v176)) : ((_393_v154).f18)), _module.__default.fm1(!(_384_v147.f15), _175_v1, _175_v1, _185_globalState), _175_v1), (_391_v152).Difference(_391_v152), _391_v152, _391_v152);
          let _434_v180;
          _434_v180 = _module.D9.create_DC24(_248_v54, _384_v147.f15, _175_v1);
          let _index49 = _module.__default.safeIndex(new BigNumber(539), new BigNumber((_180_v5).length));
          let _rhs51 = _177_v3;
          let _rhs52 = _429_v175;
          let _rhs53 = (_249_v55).Intersect(_249_v55);
          let _rhs54 = new BigNumber(((_434_v180).dtor_cf56).length);
          let _rhs55 = _433_v179;
          let _lhs41 = _180_v5;
          let _lhs42 = _module.__default.safeIndex(new BigNumber(539), new BigNumber((_180_v5).length));
          _177_v3 = _rhs51;
          _429_v175 = _rhs52;
          _249_v55 = _rhs53;
          _lhs41[_lhs42] = _rhs54;
          _433_v179 = _rhs55;
          (_185_globalState).f9 = (_393_v154).f18;
          let _435_v181;
          _435_v181 = _dafny.Seq.of(_383_v146);
          let _436_v182;
          let _nw71 = new _module.C4();
          _nw71.__ctor((_435_v181)[_module.__default.safeIndex((_180_v5)[_module.__default.safeIndex(new BigNumber(539), new BigNumber((_180_v5).length))], new BigNumber((_435_v181).length))], !(_424_cf34));
          _436_v182 = _nw71;
          let _437_v183;
          _437_v183 = _dafny.Map.Empty.slice().updateUnsafe(_436_v182,_dafny.Seq.update(_248_v54, _module.__default.safeIndex((_dafny.ZERO).minus((_430_v176).f13), new BigNumber((_248_v54).length)), _177_v3));
          let _438_v184;
          _438_v184 = _dafny.Seq.of(_module.__default.fm23(new BigNumber((_176_v2).length), _185_globalState), _module.__default.fm23((_180_v5)[_module.__default.safeIndex(new BigNumber(539), new BigNumber((_180_v5).length))], _185_globalState));
          let _439_v185;
          _439_v185 = _dafny.MultiSet.fromElements(_436_v182, _436_v182);
          let _440_v186;
          _440_v186 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(467),_386_v149);
          let _441_v187;
          _441_v187 = _dafny.Set.fromElements(_436_v182.f12);
          let _442_v188;
          _442_v188 = _dafny.Set.fromElements(_441_v187);
          let _arr2 = (_379_v144)[_module.__default.safeIndex(new BigNumber(586), new BigNumber((_379_v144).length))];
          let _index50 = _module.__default.safeIndex(new BigNumber(94), new BigNumber(((_379_v144)[_module.__default.safeIndex(new BigNumber(586), new BigNumber((_379_v144).length))]).length));
          let _rhs56 = _437_v183;
          let _rhs57 = (((_439_v185).contains(_436_v182)) ? ((_439_v185).get(_436_v182)) : (_module.__default.safeDivisionInt((_393_v154).f18, new BigNumber(222))));
          let _rhs58 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_248_v54, _248_v54), _dafny.Seq.Create(_module.__default.abs(new BigNumber(805)), ((_443_v3) => function (_444_i18) {
            return _443_v3;
          })(_177_v3)));
          let _rhs59 = _module.D8.create_DC21(_384_v147.f15, (((_440_v186).contains((_180_v5)[_module.__default.safeIndex(new BigNumber(539), new BigNumber((_180_v5).length))])) ? ((_440_v186).get((_180_v5)[_module.__default.safeIndex(new BigNumber(539), new BigNumber((_180_v5).length))])) : (_379_v144)), _module.__default.safeModuloInt((_180_v5)[_module.__default.safeIndex(new BigNumber(539), new BigNumber((_180_v5).length))], (_393_v154).f18), (_379_v144)[_module.__default.safeIndex(new BigNumber(586), new BigNumber((_379_v144).length))], (new BigNumber((_442_v188).length)).multipliedBy(_module.__default.fm1(_427_cf31, _428_cf30, _428_cf30, _185_globalState)));
          let _rhs60 = _dafny.Seq.Concat(_438_v184, _438_v184);
          let _lhs43 = (_379_v144)[_module.__default.safeIndex(new BigNumber(586), new BigNumber((_379_v144).length))];
          let _lhs44 = _module.__default.safeIndex(new BigNumber(94), new BigNumber(((_379_v144)[_module.__default.safeIndex(new BigNumber(586), new BigNumber((_379_v144).length))]).length));
          _437_v183 = _rhs56;
          _428_cf30 = _rhs57;
          _lhs43[_lhs44] = _rhs58;
          _382_v145 = _rhs59;
          _438_v184 = _rhs60;
          let _445_v189;
          let _446_v190;
          let _out10;
          let _out11;
          let _outcollector4 = (_430_v176).m2(_dafny.Seq.UnicodeFromString("jb"), (_430_v176).f14, _425_cf33, _185_globalState);
          _out10 = _outcollector4[0];
          _out11 = _outcollector4[1];
          _445_v189 = _out10;
          _446_v190 = _out11;
        } else if (_source4.is_DC14) {
          (_384_v147).f15 = true;
          (_185_globalState).f3 = _384_v147.f15;
          _392_v153 = (_392_v153).update((_393_v154).fm2(_391_v152, new BigNumber((_378_v143).length), _392_v153, _183_v6, _185_globalState), (_393_v154).f18);
          let _arr3 = (_379_v144)[_module.__default.safeIndex(new BigNumber(586), new BigNumber((_379_v144).length))];
          let _index51 = _module.__default.safeIndex(new BigNumber(94), new BigNumber(((_379_v144)[_module.__default.safeIndex(new BigNumber(586), new BigNumber((_379_v144).length))]).length));
          let _rhs61 = !(!(false));
          let _rhs62 = (_391_v152).Difference((_391_v152).Union(_391_v152));
          let _rhs63 = (_393_v154).f17;
          let _rhs64 = !(((_379_v144)[_module.__default.safeIndex(new BigNumber(586), new BigNumber((_379_v144).length))])[_module.__default.safeIndex(new BigNumber(94), new BigNumber(((_379_v144)[_module.__default.safeIndex(new BigNumber(586), new BigNumber((_379_v144).length))]).length))]);
          let _lhs45 = _185_globalState;
          let _lhs46 = (_379_v144)[_module.__default.safeIndex(new BigNumber(586), new BigNumber((_379_v144).length))];
          let _lhs47 = _module.__default.safeIndex(new BigNumber(94), new BigNumber(((_379_v144)[_module.__default.safeIndex(new BigNumber(586), new BigNumber((_379_v144).length))]).length));
          _183_v6 = _rhs61;
          _391_v152 = _rhs62;
          _lhs45.f3 = _rhs63;
          _lhs46[_lhs47] = _rhs64;
        } else {
          let _447___mcc_h17 = (_source4).cf29;
          let _448_cf29 = _447___mcc_h17;
          let _449_v191;
          let _nw72 = new _module.C2();
          _nw72.__ctor(_384_v147.f15, _383_v146, ((_379_v144)[_module.__default.safeIndex(new BigNumber(586), new BigNumber((_379_v144).length))])[_module.__default.safeIndex(new BigNumber(94), new BigNumber(((_379_v144)[_module.__default.safeIndex(new BigNumber(586), new BigNumber((_379_v144).length))]).length))]);
          _449_v191 = _nw72;
          (_185_globalState).f9 = new BigNumber((_248_v54).length);
          let _450_v192;
          let _451_v193;
          let _out12;
          let _out13;
          let _outcollector5 = (_449_v191).m5(new BigNumber(765), _dafny.Seq.Concat(_248_v54, _248_v54), _185_globalState);
          _out12 = _outcollector5[0];
          _out13 = _outcollector5[1];
          _450_v192 = _out12;
          _451_v193 = _out13;
          _451_v193 = (_dafny.ZERO).minus(_175_v1);
        }
      } else {
        let _452_v194;
        let _nw73 = Array((new BigNumber(6)).toNumber()).fill(_module.D6.Default());
        _452_v194 = _nw73;
        let _453_v195;
        _453_v195 = _dafny.Map.Empty.slice().updateUnsafe(_183_v6,_183_v6);
        let _454_v196;
        let _nw74 = new _module.C1();
        _nw74.__ctor(_175_v1, (_dafny.MultiSet.FromArray(_184_v7)).update(_183_v6, _module.__default.abs(new BigNumber((_453_v195).length))), _183_v6);
        _454_v196 = _nw74;
        let _455_v197;
        _455_v197 = _dafny.MultiSet.fromElements(_175_v1, _454_v196.f16);
        let _456_v198;
        _456_v198 = _dafny.Map.Empty.slice().updateUnsafe(_454_v196,_455_v197);
        let _457_v199;
        _457_v199 = _module.D6.create_DC15((((_456_v198).contains(_454_v196)) ? ((_456_v198).get(_454_v196)) : (_455_v197)));
        let _index52 = _module.__default.safeIndex(new BigNumber(601), new BigNumber((_452_v194).length));
        (_452_v194)[_index52] = _457_v199;
        let _index53 = _module.__default.safeIndex(new BigNumber(601), new BigNumber((_452_v194).length));
        (_452_v194)[_index53] = _457_v199;
        let _458_v200;
        _458_v200 = _module.D1.create_DC2(false, _183_v6, _183_v6, _177_v3, _177_v3);
        let _459_v201;
        _459_v201 = _module.D1.create_DC3(_458_v200);
        let _460_v202;
        _460_v202 = _module.D1.create_DC3(_459_v201);
        let _461_v203;
        _461_v203 = _module.D1.create_DC3(_460_v202);
        let _462_v204;
        _462_v204 = _module.D1.create_DC3(_461_v203);
        _462_v204 = _462_v204;
        let _463_v206;
        _463_v206 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_453_v195).length),_453_v195);
        _175_v1 = _module.__default.safeDivisionInt(new BigNumber(((function () {
          let _coll14 = new _dafny.Map();
          for (const _compr_14 of _dafny.IntegerRange(new BigNumber(359), new BigNumber(-879))) {
            let _464_v205 = _compr_14;
            if (((new BigNumber(359)).isLessThanOrEqualTo(_464_v205)) && ((_464_v205).isLessThan(new BigNumber(-879)))) {
              _coll14.push([_module.__default.safeDivisionInt(_464_v205, new BigNumber((_dafny.MultiSet.FromArray(_184_v7)).cardinality())),(_453_v195).update(!(_183_v6), _183_v6)]);
            }
          }
          return _coll14;
        }()).Merge((_463_v206).update(new BigNumber((_248_v54).length), _453_v195))).length), new BigNumber((_248_v54).length));
        let _465_v207;
        let _init8 = ((_466_v7, _467_v1) => function (_468_i19) {
          return (_466_v7)[_module.__default.safeIndex(_467_v1, new BigNumber((_466_v7).length))];
        })(_184_v7, _175_v1);
        let _nw75 = Array((new BigNumber(15)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw75.length); _i0_8++) {
          _nw75[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _465_v207 = _nw75;
        let _index54 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_465_v207).length));
        (_465_v207)[_index54] = _183_v6;
        let _469_v208;
        _469_v208 = _dafny.Set.fromElements(true);
        let _index55 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_465_v207).length));
        (_465_v207)[_index55] = !(!(_dafny.Set.fromElements(_183_v6, _183_v6, _183_v6)).equals(_469_v208)) || (((_183_v6) ? (_183_v6) : (_183_v6)));
        let _470_v209;
        _470_v209 = _dafny.Set.fromElements((_251_v57)[_module.__default.safeIndex(_454_v196.f16, new BigNumber((_251_v57).length))]);
        if (((_470_v209).Intersect(function () {
          let _coll15 = new _dafny.Set();
          for (const _compr_15 of (_dafny.Map.Empty.slice().updateUnsafe(_248_v54,_183_v6)).Keys.Elements) {
            let _471_v210 = _compr_15;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_248_v54,_183_v6)).contains(_471_v210)) {
              _coll15.add(_471_v210);
            }
          }
          return _coll15;
        }())).IsProperSubsetOf(_470_v209)) {
          _184_v7 = _dafny.Seq.update(_dafny.Seq.Concat(_184_v7, _184_v7), _module.__default.safeIndex(_175_v1, new BigNumber((_dafny.Seq.Concat(_184_v7, _184_v7)).length)), _183_v6);
          _248_v54 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(32)), ((_472_v3) => function (_473_i20) {
            return _472_v3;
          })(_177_v3));
          let _474_v211;
          _474_v211 = _dafny.MultiSet.fromElements((_465_v207)[_module.__default.safeIndex(new BigNumber(717), new BigNumber((_465_v207).length))], _183_v6, (_465_v207)[_module.__default.safeIndex(new BigNumber(717), new BigNumber((_465_v207).length))]);
          let _475_v212;
          let _nw76 = new _module.C4();
          _nw76.__ctor(_474_v211, false);
          _475_v212 = _nw76;
          let _476_v213;
          let _nw77 = new _module.C0();
          _nw77.__ctor((_465_v207)[_module.__default.safeIndex(new BigNumber(717), new BigNumber((_465_v207).length))], new BigNumber(449), _474_v211, (_465_v207)[_module.__default.safeIndex(new BigNumber(717), new BigNumber((_465_v207).length))]);
          _476_v213 = _nw77;
          _183_v6 = (_476_v213).f17;
        } else {
          let _477_v214;
          _477_v214 = _dafny.MultiSet.fromElements((_465_v207)[_module.__default.safeIndex(new BigNumber(717), new BigNumber((_465_v207).length))]);
          let _478_v215;
          let _nw78 = Array((new BigNumber(20)).toNumber());
          _nw78[(_dafny.ZERO).toNumber()] = _465_v207;
          _nw78[(_dafny.ONE).toNumber()] = _465_v207;
          _nw78[(new BigNumber(2)).toNumber()] = _465_v207;
          _nw78[(new BigNumber(3)).toNumber()] = _465_v207;
          _nw78[(new BigNumber(4)).toNumber()] = _465_v207;
          _nw78[(new BigNumber(5)).toNumber()] = _465_v207;
          _nw78[(new BigNumber(6)).toNumber()] = _465_v207;
          _nw78[(new BigNumber(7)).toNumber()] = _465_v207;
          _nw78[(new BigNumber(8)).toNumber()] = _465_v207;
          _nw78[(new BigNumber(9)).toNumber()] = _465_v207;
          _nw78[(new BigNumber(10)).toNumber()] = _465_v207;
          _nw78[(new BigNumber(11)).toNumber()] = _465_v207;
          _nw78[(new BigNumber(12)).toNumber()] = _465_v207;
          _nw78[(new BigNumber(13)).toNumber()] = _465_v207;
          _nw78[(new BigNumber(14)).toNumber()] = _465_v207;
          _nw78[(new BigNumber(15)).toNumber()] = _465_v207;
          _nw78[(new BigNumber(16)).toNumber()] = _465_v207;
          _nw78[(new BigNumber(17)).toNumber()] = _465_v207;
          _nw78[(new BigNumber(18)).toNumber()] = _465_v207;
          _nw78[(new BigNumber(19)).toNumber()] = _465_v207;
          _478_v215 = _nw78;
          let _479_v216;
          _479_v216 = _module.D8.create_DC21(false, _478_v215, _175_v1, _465_v207, new BigNumber((_dafny.Set.fromElements((_465_v207)[_module.__default.safeIndex(new BigNumber(717), new BigNumber((_465_v207).length))], (_465_v207)[_module.__default.safeIndex(new BigNumber(717), new BigNumber((_465_v207).length))])).length));
          let _480_v217;
          let _nw79 = new _module.C4();
          _nw79.__ctor(_477_v214, (_479_v216).dtor_cf49);
          _480_v217 = _nw79;
          let _481_v218;
          _481_v218 = _dafny.Map.Empty.slice().updateUnsafe(true,_480_v217);
          let _482_v219;
          _482_v219 = _dafny.Map.Empty.slice().updateUnsafe(_454_v196.f16,(_465_v207)[_module.__default.safeIndex(new BigNumber(717), new BigNumber((_465_v207).length))]);
          _481_v218 = (_481_v218).update((!(_183_v6)) && ((((_482_v219).contains(_454_v196.f16)) ? ((_482_v219).get(_454_v196.f16)) : (_183_v6))), _480_v217);
          _477_v214 = _477_v214;
          (_185_globalState).f1 = (((_480_v217).fm4(_454_v196.f16, _454_v196.f16, _185_globalState)) ? (_180_v5) : (_180_v5));
          (_185_globalState).f9 = _454_v196.f16;
          (_454_v196).f16 = _175_v1;
        }
      }
      let _483_v220;
      _483_v220 = _dafny.Set.fromElements(_183_v6);
      let _484_v221;
      let _nw80 = new _module.C0();
      _nw80.__ctor(_183_v6, new BigNumber((_483_v220).length), _dafny.MultiSet.fromElements(_183_v6, _183_v6), _183_v6);
      _484_v221 = _nw80;
      let _485_v222;
      _485_v222 = _dafny.MultiSet.fromElements((_484_v221).f17, _183_v6, _183_v6);
      let _486_v223;
      let _nw81 = new _module.C3();
      _nw81.__ctor(_175_v1, (_484_v221).f17, _485_v222, _183_v6);
      _486_v223 = _nw81;
      let _487_v224;
      _487_v224 = _dafny.Map.Empty.slice().updateUnsafe(_175_v1,_486_v223);
      let _488_v225;
      _488_v225 = _dafny.Seq.of(_485_v222, _485_v222, _485_v222, _module.__default.fm22(new BigNumber((_487_v224).length), _183_v6, _185_globalState));
      let _489_v226;
      _489_v226 = _dafny.Map.Empty.slice().updateUnsafe(_484_v221,((_488_v225)[_module.__default.safeIndex((_484_v221).f18, new BigNumber((_488_v225).length))]).IsDisjointFrom(_module.__default.fm22((_486_v223).f13, (_486_v223).f14, _185_globalState)));
      _489_v226 = (_489_v226).update(_484_v221, !((_486_v223).f14));
      let _490_v227;
      let _init9 = ((_491_v54) => function (_492_i22) {
        return _491_v54;
      })(_248_v54);
      let _nw82 = Array((new BigNumber(27)).toNumber());
      for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw82.length); _i0_9++) {
        _nw82[_i0_9] = _init9(new BigNumber(_i0_9));
      }
      _490_v227 = _nw82;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_490_v227).length))) {
        let _493_i21 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_493_i21)) && ((_493_i21).isLessThan(new BigNumber((_490_v227).length))))) {
          (_490_v227)[(_493_i21)] = _248_v54;
        }
      }
      let _494_v228;
      let _nw83 = new _module.C1();
      _nw83.__ctor((_486_v223).f13, _485_v222, (_484_v221).f17);
      _494_v228 = _nw83;
      let _495_v229;
      _495_v229 = _dafny.Seq.of(_494_v228, _494_v228);
      let _496_i23;
      _496_i23 = _dafny.ZERO;
      L4: {
        while (!(!(_dafny.areEqual(_dafny.Seq.update(_dafny.Seq.Concat(_495_v229, _dafny.Seq.of(_494_v228)), _module.__default.safeIndex(_175_v1, new BigNumber((_dafny.Seq.Concat(_495_v229, _dafny.Seq.of(_494_v228))).length)), _494_v228), _495_v229)))) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_496_i23)) {
              break L4;
            }
            _496_i23 = (_496_i23).plus(_dafny.ONE);
            let _497_v230;
            let _nw84 = new _module.C1();
            _nw84.__ctor((_486_v223).f13, _dafny.MultiSet.fromElements(_183_v6), (_484_v221).f17);
            _497_v230 = _nw84;
            let _498_v231;
            _498_v231 = _dafny.Set.fromElements(_497_v230);
            _251_v57 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(485)), ((_499_v3) => function (_500_i24) {
              return _499_v3;
            })(_177_v3)), _248_v54, _dafny.Seq.update(_dafny.Seq.UnicodeFromString("ot"), _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_498_v231).length)), new BigNumber((_dafny.Seq.UnicodeFromString("ot")).length)), new _dafny.CodePoint('r'.codePointAt(0))), _248_v54);
            let _501_v233;
            let _init10 = ((_502_v221, _503_v1, _504_v223, _505_v2, _506_v228) => function (_507_i25) {
              return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_502_v221).f18, _503_v1),(_504_v223).f13)).Merge(function () {
                let _coll16 = new _dafny.Map();
                for (const _compr_16 of (_dafny.MultiSet.fromElements(_dafny.Seq.update(_505_v2, _module.__default.safeIndex(_506_v228.f16, new BigNumber((_505_v2).length)), _503_v1), _505_v2)).Elements) {
                  let _508_v232 = _compr_16;
                  if ((_dafny.MultiSet.fromElements(_dafny.Seq.update(_505_v2, _module.__default.safeIndex(_506_v228.f16, new BigNumber((_505_v2).length)), _503_v1), _505_v2)).contains(_508_v232)) {
                    _coll16.push([_508_v232,(_504_v223).f13]);
                  }
                }
                return _coll16;
              }());
            })(_484_v221, _175_v1, _486_v223, _176_v2, _494_v228);
            let _nw85 = Array((new BigNumber(8)).toNumber());
            for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw85.length); _i0_10++) {
              _nw85[_i0_10] = _init10(new BigNumber(_i0_10));
            }
            _501_v233 = _nw85;
            let _509_v234;
            _509_v234 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_248_v54).length)),_175_v1);
            let _510_v235;
            _510_v235 = _509_v234;
            let _index56 = _module.__default.safeIndex(new BigNumber(824), new BigNumber((_501_v233).length));
            (_501_v233)[_index56] = (_510_v235);
            let _index57 = _module.__default.safeIndex(new BigNumber(824), new BigNumber((_501_v233).length));
            (_501_v233)[_index57] = ((_509_v234).Merge(_dafny.Map.Empty.slice().updateUnsafe(_176_v2,_175_v1))).Merge((_509_v234).Merge(function () {
              let _coll17 = new _dafny.Map();
              for (const _compr_17 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(82)), ((_511_v2) => function (_512_i26) {
                return _511_v2;
              })(_176_v2))).Elements) {
                let _513_v236 = _compr_17;
                if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(82)), ((_514_v2) => function (_512_i26) {
                  return _514_v2;
                })(_176_v2)), _513_v236)) {
                  _coll17.push([_513_v236,(_484_v221).f18]);
                }
              }
              return _coll17;
            }()));
            let _515_v237;
            _515_v237 = _dafny.MultiSet.fromElements(_494_v228.f16, new BigNumber(813), (_484_v221).f18, new BigNumber((_184_v7).length));
            let _516_v238;
            _516_v238 = _dafny.Map.Empty.slice().updateUnsafe(_497_v230.f12,(_486_v223).f13);
            let _517_v239;
            _517_v239 = _dafny.Seq.of(_176_v2, (((_494_v228).fm2(_515_v237, new BigNumber(673), _516_v238, (_486_v223).fm6((_486_v223).f14, _175_v1, _185_globalState), _185_globalState)) ? (_176_v2) : (_176_v2)), _dafny.Seq.Concat(_176_v2, _176_v2));
            let _518_v240;
            _518_v240 = _dafny.Seq.of(_517_v239, _dafny.Seq.Concat(_517_v239, _517_v239));
            let _519_v241;
            _519_v241 = _module.D9.create_DC24(_248_v54, true, new BigNumber(707));
            _517_v239 = _dafny.Seq.update((_518_v240)[_module.__default.safeIndex((_486_v223).f13, new BigNumber((_518_v240).length))], _module.__default.safeIndex((_484_v221).f18, new BigNumber(((_518_v240)[_module.__default.safeIndex((_486_v223).f13, new BigNumber((_518_v240).length))]).length)), _dafny.Seq.Concat(_dafny.Seq.of((_519_v241).dtor_cf58), _176_v2));
            (_497_v230).f12 = !(_183_v6);
          }
        }
      }
      _module.__default.m0(_185_globalState);
      if (_183_v6) {
        let _rhs65 = _248_v54;
        let _rhs66 = _176_v2;
        _248_v54 = _rhs65;
        _176_v2 = _rhs66;
        let _520_v242;
        _520_v242 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_248_v54).length),(_484_v221).f17);
        let _521_v243;
        let _nw86 = new _module.C0();
        _nw86.__ctor((_486_v223).f14, (_484_v221).f18, _485_v222, _183_v6);
        _521_v243 = _nw86;
        let _522_v244;
        _522_v244 = _dafny.Map.Empty.slice().updateUnsafe(_177_v3,_521_v243);
        let _523_v245;
        _523_v245 = _dafny.Seq.of((_522_v244).update(_177_v3, _521_v243));
        let _524_v246;
        let _nw87 = new _module.C1();
        _nw87.__ctor(_module.__default.safeModuloInt(new BigNumber(847), new BigNumber((_248_v54).length)), _dafny.MultiSet.fromElements((((_520_v242).contains(new BigNumber((_523_v245).length))) ? ((_520_v242).get(new BigNumber((_523_v245).length))) : (_521_v243.f12)), (_486_v223).f14), !((_484_v221).f17));
        _524_v246 = _nw87;
        if (!((_486_v223).fm6(!(!(_183_v6)), (_484_v221).fm12(_185_globalState), _185_globalState)) || (_521_v243.f12)) {
          let _525_v247;
          let _nw88 = Array((new BigNumber(10)).toNumber()).fill(false);
          _525_v247 = _nw88;
          let _526_v248;
          _526_v248 = _dafny.Seq.of(_525_v247);
          let _527_v249;
          _527_v249 = _dafny.MultiSet.fromElements(_494_v228.f16, _175_v1, new BigNumber((_526_v248).length), (_dafny.ZERO).minus(_494_v228.f16), _494_v228.f16);
          let _528_v250;
          _528_v250 = _dafny.Map.Empty.slice().updateUnsafe(_175_v1,_527_v249);
          let _529_v251;
          let _nw89 = Array((new BigNumber(12)).toNumber());
          _nw89[(_dafny.ZERO).toNumber()] = _module.__default.fm13(_177_v3, !((_484_v221).f17), _185_globalState);
          _nw89[(_dafny.ONE).toNumber()] = _527_v249;
          _nw89[(new BigNumber(2)).toNumber()] = _527_v249;
          _nw89[(new BigNumber(3)).toNumber()] = _527_v249;
          _nw89[(new BigNumber(4)).toNumber()] = (((_528_v250).contains((_484_v221).f18)) ? ((_528_v250).get((_484_v221).f18)) : (_dafny.MultiSet.fromElements(_494_v228.f16)));
          _nw89[(new BigNumber(5)).toNumber()] = _527_v249;
          _nw89[(new BigNumber(6)).toNumber()] = (_527_v249).update((_486_v223).f13, _module.__default.abs((_486_v223).f13));
          _nw89[(new BigNumber(7)).toNumber()] = _527_v249;
          _nw89[(new BigNumber(8)).toNumber()] = (_dafny.MultiSet.fromElements(new BigNumber((_176_v2).length))).Union(_527_v249);
          _nw89[(new BigNumber(9)).toNumber()] = _dafny.MultiSet.fromElements(_494_v228.f16, _524_v246.f16, new BigNumber((_dafny.Seq.UnicodeFromString("co")).length));
          _nw89[(new BigNumber(10)).toNumber()] = ((_521_v243.f12) ? (_527_v249) : (_527_v249));
          _nw89[(new BigNumber(11)).toNumber()] = (_dafny.MultiSet.fromElements(new BigNumber(162), _494_v228.f16)).update((_dafny.ZERO).minus(_175_v1), _module.__default.abs(_494_v228.f16));
          _529_v251 = _nw89;
          let _index58 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_529_v251).length));
          (_529_v251)[_index58] = _module.__default.fm13(new _dafny.CodePoint('h'.codePointAt(0)), (_486_v223).f14, _185_globalState);
          let _index59 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_529_v251).length));
          (_529_v251)[_index59] = _527_v249;
          let _530_v252;
          _530_v252 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_176_v2, _module.__default.safeIndex(_524_v246.f16, new BigNumber((_176_v2).length)), (_484_v221).f18),new BigNumber(354));
          let _531_v253;
          _531_v253 = _530_v252;
          let _532_v254;
          _532_v254 = _dafny.Map.Empty.slice().updateUnsafe(_531_v253,(_486_v223).f14);
          let _533_v255;
          let _nw90 = Array((new BigNumber(10)).toNumber()).fill(_module.D8.Default());
          _533_v255 = _nw90;
          let _534_v256;
          _534_v256 = _dafny.Map.Empty.slice().updateUnsafe(_532_v254,_533_v255);
          let _index60 = _module.__default.safeIndex(new BigNumber(120), new BigNumber((_180_v5).length));
          (_180_v5)[_index60] = _module.__default.safeModuloInt((_484_v221).f18, _494_v228.f16);
          let _index61 = _module.__default.safeIndex(new BigNumber(120), new BigNumber((_180_v5).length));
          let _rhs67 = _534_v256;
          let _rhs68 = !((_486_v223).f13).isEqualTo((_484_v221).f18);
          let _rhs69 = (((_184_v7)[_module.__default.safeIndex((_484_v221).f18, new BigNumber((_184_v7).length))]) ? ((((_484_v221).f17) ? (_486_v223) : (_486_v223))) : (_486_v223));
          let _rhs70 = _494_v228.f16;
          let _lhs48 = _521_v243;
          let _lhs49 = _180_v5;
          let _lhs50 = _module.__default.safeIndex(new BigNumber(120), new BigNumber((_180_v5).length));
          _534_v256 = _rhs67;
          _lhs48.f12 = _rhs68;
          _486_v223 = _rhs69;
          _lhs49[_lhs50] = _rhs70;
          let _535_v257;
          let _nw91 = new _module.C2();
          _nw91.__ctor(_521_v243.f12, _521_v243.f11, _521_v243.f12);
          _535_v257 = _nw91;
          let _536_v258;
          _536_v258 = _module.D12.create_DC32(_module.__default.safeDivisionInt((_484_v221).f18, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-861),_524_v246)).length)), _535_v257, _484_v221, (_486_v223).f13);
          let _537_v259;
          _537_v259 = _dafny.Map.Empty.slice().updateUnsafe(_524_v246.f16,_484_v221);
          let _538_v260;
          _538_v260 = _dafny.Seq.of(_484_v221, (((_537_v259).contains((_484_v221).f18)) ? ((_537_v259).get((_484_v221).f18)) : (_484_v221)), _484_v221);
          _536_v258 = _module.D12.create_DC32(_494_v228.f16, _535_v257, (_538_v260)[_module.__default.safeIndex(new BigNumber((_483_v220).length), new BigNumber((_538_v260).length))], (_dafny.ZERO).minus(new BigNumber((_248_v54).length)));
          let _539_v261;
          let _nw92 = new _module.C3();
          _nw92.__ctor((_dafny.ZERO).minus((_486_v223).f13), _521_v243.f12, (_dafny.MultiSet.fromElements(false, (_484_v221).f17)).Union(_485_v222), (_484_v221).f17);
          _539_v261 = _nw92;
          let _540_v262;
          let _nw93 = Array((new BigNumber(13)).toNumber());
          _nw93[(_dafny.ZERO).toNumber()] = ((_486_v223).f13).plus(_494_v228.f16);
          _nw93[(_dafny.ONE).toNumber()] = (_486_v223).f13;
          _nw93[(new BigNumber(2)).toNumber()] = (_494_v228.f16).plus((_dafny.ZERO).minus(_524_v246.f16));
          _nw93[(new BigNumber(3)).toNumber()] = _module.__default.fm1(_535_v257.f15, (_486_v223).f13, new BigNumber((_249_v55).length), _185_globalState);
          _nw93[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(((_486_v223).f13).multipliedBy((_484_v221).f18));
          _nw93[(new BigNumber(5)).toNumber()] = _494_v228.f16;
          _nw93[(new BigNumber(6)).toNumber()] = (_486_v223).f13;
          _nw93[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(_494_v228.f16);
          _nw93[(new BigNumber(8)).toNumber()] = (_486_v223).f13;
          _nw93[(new BigNumber(9)).toNumber()] = _494_v228.f16;
          _nw93[(new BigNumber(10)).toNumber()] = ((!(!(_535_v257.f15))) ? (new BigNumber((_248_v54).length)) : (_175_v1));
          _nw93[(new BigNumber(11)).toNumber()] = new BigNumber((_184_v7).length);
          _nw93[(new BigNumber(12)).toNumber()] = (new BigNumber((_248_v54).length)).minus((_484_v221).f18);
          _540_v262 = _nw93;
          let _rhs71 = !(!(true) || (_521_v243.f12));
          let _rhs72 = _540_v262;
          let _rhs73 = ((((_539_v261).f13).isLessThanOrEqualTo(new BigNumber(-876))) ? ((_175_v1).plus((_484_v221).f18)) : (new BigNumber(979)));
          let _lhs51 = _185_globalState;
          let _lhs52 = _185_globalState;
          _lhs51.f3 = _rhs71;
          _lhs52.f1 = _rhs72;
          _175_v1 = _rhs73;
        } else {
          _183_v6 = _521_v243.f12;
          let _541_v263;
          _541_v263 = _dafny.Map.Empty.slice().updateUnsafe(_176_v2,_494_v228.f16);
          let _542_v264;
          _542_v264 = _dafny.Map.Empty.slice().updateUnsafe(_176_v2,_524_v246.f16);
          let _543_v265;
          let _nw94 = Array((new BigNumber(4)).toNumber());
          _nw94[(_dafny.ZERO).toNumber()] = _541_v263;
          _nw94[(_dafny.ONE).toNumber()] = _542_v264;
          _nw94[(new BigNumber(2)).toNumber()] = _541_v263;
          _nw94[(new BigNumber(3)).toNumber()] = _541_v263;
          _543_v265 = _nw94;
          let _544_v266;
          _544_v266 = _dafny.Map.Empty.slice().updateUnsafe((_486_v223).f14,_543_v265);
          let _545_v267;
          let _nw95 = Array((new BigNumber(10)).toNumber());
          _nw95[(_dafny.ZERO).toNumber()] = (_dafny.Set.fromElements(new BigNumber((_483_v220).length), _175_v1, _494_v228.f16)).IsProperSubsetOf(_249_v55);
          _nw95[(_dafny.ONE).toNumber()] = (new BigNumber(533)).isLessThan((_486_v223).f13);
          _nw95[(new BigNumber(2)).toNumber()] = ((_484_v221).f17) && ((_486_v223).f14);
          _nw95[(new BigNumber(3)).toNumber()] = _183_v6;
          _nw95[(new BigNumber(4)).toNumber()] = (_486_v223).f14;
          _nw95[(new BigNumber(5)).toNumber()] = (_486_v223).f14;
          _nw95[(new BigNumber(6)).toNumber()] = _521_v243.f12;
          _nw95[(new BigNumber(7)).toNumber()] = !(_544_v266).contains(!(_521_v243.f12));
          _nw95[(new BigNumber(8)).toNumber()] = (_184_v7)[_module.__default.safeIndex(_module.__default.fm1(_183_v6, _175_v1, _494_v228.f16, _185_globalState), new BigNumber((_184_v7).length))];
          _nw95[(new BigNumber(9)).toNumber()] = (_484_v221).f17;
          _545_v267 = _nw95;
          let _index62 = _module.__default.safeIndex(new BigNumber(622), new BigNumber((_545_v267).length));
          (_545_v267)[_index62] = (_486_v223).f14;
          let _index63 = _module.__default.safeIndex(new BigNumber(622), new BigNumber((_545_v267).length));
          let _rhs74 = (_dafny.Seq.IsPrefixOf(_176_v2, _176_v2)) || ((_484_v221).f17);
          let _rhs75 = (_484_v221).f18;
          let _rhs76 = (_484_v221).f17;
          let _rhs77 = _177_v3;
          let _lhs53 = _185_globalState;
          let _lhs54 = _185_globalState;
          let _lhs55 = _545_v267;
          let _lhs56 = _module.__default.safeIndex(new BigNumber(622), new BigNumber((_545_v267).length));
          _lhs53.f7 = _rhs74;
          _lhs54.f9 = _rhs75;
          _lhs55[_lhs56] = _rhs76;
          _177_v3 = _rhs77;
          let _546_v268;
          let _out14;
          _out14 = (_494_v228).m6(_185_globalState);
          _546_v268 = _out14;
          let _547_v269;
          _547_v269 = _dafny.Map.Empty.slice().updateUnsafe(false,(_486_v223).f13);
          let _548_v270;
          _548_v270 = _dafny.Map.Empty.slice().updateUnsafe((_486_v223).f14,(_486_v223).f14);
          let _549_v271;
          _549_v271 = _dafny.MultiSet.fromElements(new BigNumber((_176_v2).length), (_484_v221).fm12(_185_globalState), new BigNumber((_548_v270).length), _494_v228.f16, (_484_v221).f18);
          let _550_v272;
          _550_v272 = _dafny.Map.Empty.slice().updateUnsafe(_180_v5,true);
          let _551_v273;
          _551_v273 = _dafny.Map.Empty.slice().updateUnsafe((((_547_v269).contains((_521_v243).fm2(_549_v271, (_dafny.ZERO).minus(new BigNumber((_548_v270).length)), _547_v269, (_545_v267)[_module.__default.safeIndex(new BigNumber(622), new BigNumber((_545_v267).length))], _185_globalState))) ? ((_547_v269).get((_521_v243).fm2(_549_v271, (_dafny.ZERO).minus(new BigNumber((_548_v270).length)), _547_v269, (_545_v267)[_module.__default.safeIndex(new BigNumber(622), new BigNumber((_545_v267).length))], _185_globalState))) : (new BigNumber((_550_v272).length))),_545_v267);
          let _552_v274;
          _552_v274 = _dafny.Set.fromElements(_177_v3);
          let _553_v275;
          _553_v275 = _dafny.Map.Empty.slice().updateUnsafe((_552_v274).equals(_552_v274),(((_486_v223).f14) ? (_545_v267) : (_545_v267)));
          let _rhs78 = _551_v273;
          let _rhs79 = (((_486_v223).f14) ? ((_545_v267)[_module.__default.safeIndex(new BigNumber(622), new BigNumber((_545_v267).length))]) : (_521_v243.f12));
          let _rhs80 = new BigNumber((_248_v54).length);
          let _rhs81 = ((_484_v221).f18).plus((_dafny.ZERO).minus(_494_v228.f16));
          let _rhs82 = (_553_v275).Merge(_553_v275);
          let _lhs57 = _185_globalState;
          let _lhs58 = _494_v228;
          let _lhs59 = _494_v228;
          _551_v273 = _rhs78;
          _lhs57.f3 = _rhs79;
          _lhs58.f16 = _rhs80;
          _lhs59.f16 = _rhs81;
          _553_v275 = _rhs82;
          _548_v270 = (_548_v270).update((_484_v221).f17, (_486_v223).f14);
        }
        _486_v223 = _486_v223;
        let _554_v276;
        _554_v276 = _module.D1.create_DC2((_484_v221).f17, (_486_v223).f14, (_484_v221).f17, _177_v3, _177_v3);
        let _pat_let_tv14 = _486_v223;
        let _555_v277;
        let _nw96 = Array((new BigNumber(12)).toNumber());
        _nw96[(_dafny.ZERO).toNumber()] = _module.__default.fm14(_554_v276, _524_v246.f16, _185_globalState);
        _nw96[(_dafny.ONE).toNumber()] = _177_v3;
        _nw96[(new BigNumber(2)).toNumber()] = _177_v3;
        _nw96[(new BigNumber(3)).toNumber()] = _177_v3;
        _nw96[(new BigNumber(4)).toNumber()] = _module.__default.fm14(function (_pat_let17_0) {
          return function (_556_dt__update__tmp_h3) {
            return function (_pat_let18_0) {
              return function (_557_dt__update_hcf3_h0) {
                return _module.D1.create_DC2((_556_dt__update__tmp_h3).dtor_cf2, _557_dt__update_hcf3_h0, (_556_dt__update__tmp_h3).dtor_cf4, (_556_dt__update__tmp_h3).dtor_cf5, (_556_dt__update__tmp_h3).dtor_cf6);
              }(_pat_let18_0);
            }((_pat_let_tv14).f14);
          }(_pat_let17_0);
        }(_554_v276), (_484_v221).f18, _185_globalState);
        _nw96[(new BigNumber(5)).toNumber()] = _177_v3;
        _nw96[(new BigNumber(6)).toNumber()] = _177_v3;
        _nw96[(new BigNumber(7)).toNumber()] = _177_v3;
        _nw96[(new BigNumber(8)).toNumber()] = _177_v3;
        _nw96[(new BigNumber(9)).toNumber()] = _177_v3;
        _nw96[(new BigNumber(10)).toNumber()] = ((!(!((_486_v223).f14))) ? ((_248_v54)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_486_v223).fm6(_183_v6, (_486_v223).f13, _185_globalState),_175_v1)).length), new BigNumber((_248_v54).length))]) : (_177_v3));
        _nw96[(new BigNumber(11)).toNumber()] = _177_v3;
        _555_v277 = _nw96;
        let _index64 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_555_v277).length));
        (_555_v277)[_index64] = _module.__default.fm14(_554_v276, (_484_v221).f18, _185_globalState);
        let _558_v278;
        let _nw97 = new _module.C2();
        _nw97.__ctor((_520_v242).equals(_520_v242), _521_v243.f11, _521_v243.f12);
        _558_v278 = _nw97;
        let _559_v279;
        _559_v279 = _dafny.MultiSet.fromElements(new BigNumber((_184_v7).length), new BigNumber((_176_v2).length), _494_v228.f16, _524_v246.f16);
        let _560_v280;
        _560_v280 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-467),_559_v279);
        let _561_v281;
        let _nw98 = Array((new BigNumber(12)).toNumber());
        _nw98[(_dafny.ZERO).toNumber()] = (_559_v279).update(_494_v228.f16, _module.__default.abs(_175_v1));
        _nw98[(_dafny.ONE).toNumber()] = (_559_v279).Intersect(_559_v279);
        _nw98[(new BigNumber(2)).toNumber()] = _559_v279;
        _nw98[(new BigNumber(3)).toNumber()] = _559_v279;
        _nw98[(new BigNumber(4)).toNumber()] = _559_v279;
        _nw98[(new BigNumber(5)).toNumber()] = _dafny.MultiSet.FromArray(_176_v2);
        _nw98[(new BigNumber(6)).toNumber()] = (((_560_v280).contains((_484_v221).f18)) ? ((_560_v280).get((_484_v221).f18)) : (_559_v279));
        _nw98[(new BigNumber(7)).toNumber()] = _559_v279;
        _nw98[(new BigNumber(8)).toNumber()] = _559_v279;
        _nw98[(new BigNumber(9)).toNumber()] = (_module.__default.fm13(_177_v3, (_484_v221).f17, _185_globalState)).update((_484_v221).f18, _module.__default.abs((_dafny.ZERO).minus(new BigNumber((_248_v54).length))));
        _nw98[(new BigNumber(10)).toNumber()] = (_dafny.MultiSet.fromElements(_524_v246.f16)).Union(_559_v279);
        _nw98[(new BigNumber(11)).toNumber()] = _559_v279;
        _561_v281 = _nw98;
        let _index65 = _module.__default.safeIndex(new BigNumber(81), new BigNumber((_561_v281).length));
        (_561_v281)[_index65] = _559_v279;
        let _index66 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_555_v277).length));
        let _index67 = _module.__default.safeIndex(new BigNumber(81), new BigNumber((_561_v281).length));
        let _rhs83 = _494_v228.f16;
        let _rhs84 = _177_v3;
        let _rhs85 = _558_v278;
        let _rhs86 = _559_v279;
        let _lhs60 = _185_globalState;
        let _lhs61 = _555_v277;
        let _lhs62 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_555_v277).length));
        let _lhs63 = _561_v281;
        let _lhs64 = _module.__default.safeIndex(new BigNumber(81), new BigNumber((_561_v281).length));
        _lhs60.f9 = _rhs83;
        _lhs61[_lhs62] = _rhs84;
        _558_v278 = _rhs85;
        _lhs63[_lhs64] = _rhs86;
      } else {
        let _562_v282;
        let _nw99 = new _module.C1();
        _nw99.__ctor((_dafny.ZERO).minus(new BigNumber((_485_v222).cardinality())), _485_v222, true);
        _562_v282 = _nw99;
        let _563_v283;
        _563_v283 = _dafny.Map.Empty.slice().updateUnsafe(_562_v282,(((_484_v221).f17) ? (new BigNumber((_dafny.Seq.update((_251_v57)[_module.__default.safeIndex(new BigNumber((_249_v55).length), new BigNumber((_251_v57).length))], _module.__default.safeIndex((_484_v221).f18, new BigNumber(((_251_v57)[_module.__default.safeIndex(new BigNumber((_249_v55).length), new BigNumber((_251_v57).length))]).length)), _177_v3)).length)) : (new BigNumber((_184_v7).length))));
        _563_v283 = (_563_v283).update(_562_v282, _module.__default.fm1((_184_v7)[_module.__default.safeIndex((_484_v221).f18, new BigNumber((_184_v7).length))], (_dafny.ZERO).minus((_484_v221).f18), (_486_v223).f13, _185_globalState));
        let _564_v284;
        _564_v284 = _dafny.Map.Empty.slice().updateUnsafe(!((_486_v223).f14),false);
        let _565_v285;
        _565_v285 = _dafny.Seq.of(_564_v284, (_564_v284).update(_562_v282.f12, !((_486_v223).fm6((_486_v223).f14, (((_485_v222).contains((_486_v223).f14)) ? ((_485_v222).get((_486_v223).f14)) : (_175_v1)), _185_globalState))), _564_v284);
        _565_v285 = _dafny.Seq.of((_564_v284).Merge((_564_v284).update((_484_v221).f17, _183_v6)), _564_v284, _564_v284, (_564_v284).Merge(_564_v284));
        let _566_v286;
        let _init11 = function (_567_i27) {
          return true;
        };
        let _nw100 = Array((new BigNumber(25)).toNumber());
        for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw100.length); _i0_11++) {
          _nw100[_i0_11] = _init11(new BigNumber(_i0_11));
        }
        _566_v286 = _nw100;
        let _568_v287;
        _568_v287 = _dafny.Map.Empty.slice().updateUnsafe(false,_566_v286);
        let _569_v288;
        let _nw101 = Array((new BigNumber(22)).toNumber());
        _nw101[(_dafny.ZERO).toNumber()] = _566_v286;
        _nw101[(_dafny.ONE).toNumber()] = _566_v286;
        _nw101[(new BigNumber(2)).toNumber()] = _566_v286;
        _nw101[(new BigNumber(3)).toNumber()] = _566_v286;
        _nw101[(new BigNumber(4)).toNumber()] = _566_v286;
        _nw101[(new BigNumber(5)).toNumber()] = _566_v286;
        _nw101[(new BigNumber(6)).toNumber()] = _566_v286;
        _nw101[(new BigNumber(7)).toNumber()] = _566_v286;
        _nw101[(new BigNumber(8)).toNumber()] = _566_v286;
        _nw101[(new BigNumber(9)).toNumber()] = _566_v286;
        _nw101[(new BigNumber(10)).toNumber()] = _566_v286;
        _nw101[(new BigNumber(11)).toNumber()] = _566_v286;
        _nw101[(new BigNumber(12)).toNumber()] = _566_v286;
        _nw101[(new BigNumber(13)).toNumber()] = _566_v286;
        _nw101[(new BigNumber(14)).toNumber()] = _566_v286;
        _nw101[(new BigNumber(15)).toNumber()] = (((_568_v287).contains(_183_v6)) ? ((_568_v287).get(_183_v6)) : (_566_v286));
        _nw101[(new BigNumber(16)).toNumber()] = _566_v286;
        _nw101[(new BigNumber(17)).toNumber()] = _566_v286;
        _nw101[(new BigNumber(18)).toNumber()] = _566_v286;
        _nw101[(new BigNumber(19)).toNumber()] = _566_v286;
        _nw101[(new BigNumber(20)).toNumber()] = _566_v286;
        _nw101[(new BigNumber(21)).toNumber()] = _566_v286;
        _569_v288 = _nw101;
        let _570_v289;
        _570_v289 = _module.D8.create_DC21((_484_v221).f17, _569_v288, (_484_v221).f18, _566_v286, _module.__default.safeModuloInt((_486_v223).f13, (_484_v221).f18));
        let _source5 = _570_v289;
        if (_source5.is_DC20) {
          let _571___mcc_h18 = (_source5).cf46;
          let _572___mcc_h19 = (_source5).cf47;
          let _573___mcc_h20 = (_source5).cf48;
          let _574_cf48 = _573___mcc_h20;
          let _575_cf47 = _572___mcc_h19;
          let _576_cf46 = _571___mcc_h18;
          let _577_v290;
          _577_v290 = _dafny.Seq.of(_483_v220);
          _175_v1 = (_dafny.ZERO).minus(new BigNumber((((_577_v290)[_module.__default.safeIndex(_494_v228.f16, new BigNumber((_577_v290).length))]).Difference((_483_v220).Intersect(_483_v220))).length));
          let _578_v291;
          _578_v291 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(803),_177_v3);
          _177_v3 = (((_578_v291).contains(_494_v228.f16)) ? ((_578_v291).get(_494_v228.f16)) : (_177_v3));
          let _579_v292;
          _579_v292 = _dafny.MultiSet.fromElements(_177_v3, _177_v3, _177_v3, (_248_v54)[_module.__default.safeIndex((_484_v221).f18, new BigNumber((_248_v54).length))]);
          let _580_v293;
          _580_v293 = _dafny.MultiSet.fromElements(_177_v3);
          let _581_v294;
          _581_v294 = _dafny.Seq.of(_579_v292);
          _575_cf47 = ((_581_v294)[_module.__default.safeIndex(new BigNumber((_248_v54).length), new BigNumber((_581_v294).length))]).IsSubsetOf(((true) ? (_579_v292) : ((_580_v293))));
          (_185_globalState).f3 = (_486_v223).f14;
        } else if (_source5.is_DC21) {
          let _582___mcc_h21 = (_source5).cf49;
          let _583___mcc_h22 = (_source5).cf50;
          let _584___mcc_h23 = (_source5).cf51;
          let _585___mcc_h24 = (_source5).cf52;
          let _586___mcc_h25 = (_source5).cf53;
          let _587_cf53 = _586___mcc_h25;
          let _588_cf52 = _585___mcc_h24;
          let _589_cf51 = _584___mcc_h23;
          let _590_cf50 = _583___mcc_h22;
          let _591_cf49 = _582___mcc_h21;
          _589_cf51 = new BigNumber(((_251_v57)[_module.__default.safeIndex(((_562_v282.f12) ? (_175_v1) : ((_484_v221).f18)), new BigNumber((_251_v57).length))]).length);
          (_562_v282).f12 = _591_cf49;
          _588_cf52 = _566_v286;
          let _592_v295;
          _592_v295 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(235)), ((_593_v228) => function (_594_i28) {
            return _593_v228.f16;
          })(_494_v228)));
          _589_cf51 = (((_592_v295).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(899)), function (_596_i29) {
            return (_dafny.ZERO).minus(new BigNumber(-198));
          }))) ? ((_592_v295).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(899)), function (_595_i29) {
            return (_dafny.ZERO).minus(new BigNumber(-198));
          }))) : (_module.__default.safeDivisionInt((_484_v221).f18, (_484_v221).f18)));
        } else if (_source5.is_DC19) {
          let _597___mcc_h26 = (_source5).cf45;
          let _598_cf45 = _597___mcc_h26;
          _module.__default.m0(_185_globalState);
          (_185_globalState).f9 = ((_562_v282.f12) ? ((_486_v223).f13) : ((_486_v223).f13));
          (_185_globalState).f9 = ((_484_v221).f18).multipliedBy(_494_v228.f16);
          _562_v282 = _562_v282;
        } else {
          let _599___mcc_h27 = (_source5).cf54;
          let _600_cf54 = _599___mcc_h27;
          (_562_v282).f12 = _562_v282.f12;
          let _601_v296;
          _601_v296 = _module.D1.create_DC2(_562_v282.f12, (_484_v221).f17, (_486_v223).f14, _177_v3, _177_v3);
          _177_v3 = _module.__default.fm14(_601_v296, _175_v1, _185_globalState);
          let _602_v297;
          _602_v297 = _dafny.Map.Empty.slice().updateUnsafe((_486_v223).f14,(_484_v221).f18);
          let _603_v298;
          let _nw102 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.of());
          _603_v298 = _nw102;
          let _604_v299;
          _604_v299 = _dafny.Map.Empty.slice().updateUnsafe(_603_v298,_602_v297);
          let _605_v300;
          _605_v300 = _dafny.Seq.of(_602_v297, _602_v297);
          let _606_v301;
          _606_v301 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_248_v54).length),(_484_v221).f18);
          let _607_v302;
          _607_v302 = _dafny.MultiSet.fromElements(_494_v228.f16, (_486_v223).f13);
          let _608_v303;
          _608_v303 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements((((_606_v301).contains(new BigNumber((_607_v302).cardinality()))) ? ((_606_v301).get(new BigNumber((_607_v302).cardinality()))) : (_494_v228.f16)), _175_v1),(_484_v221).f18);
          let _609_v304;
          let _nw103 = Array((new BigNumber(19)).toNumber());
          _nw103[(_dafny.ZERO).toNumber()] = _602_v297;
          _nw103[(_dafny.ONE).toNumber()] = _602_v297;
          _nw103[(new BigNumber(2)).toNumber()] = (((_604_v299).contains(_603_v298)) ? ((_604_v299).get(_603_v298)) : (_602_v297));
          _nw103[(new BigNumber(3)).toNumber()] = _602_v297;
          _nw103[(new BigNumber(4)).toNumber()] = _module.__default.fm23(_494_v228.f16, _185_globalState);
          _nw103[(new BigNumber(5)).toNumber()] = _602_v297;
          _nw103[(new BigNumber(6)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_484_v221).f17,(_486_v223).f13);
          _nw103[(new BigNumber(7)).toNumber()] = _602_v297;
          _nw103[(new BigNumber(8)).toNumber()] = (_602_v297).Merge(_602_v297);
          _nw103[(new BigNumber(9)).toNumber()] = (_602_v297).Merge(_602_v297);
          _nw103[(new BigNumber(10)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_484_v221).f17,(_484_v221).f18);
          _nw103[(new BigNumber(11)).toNumber()] = _602_v297;
          _nw103[(new BigNumber(12)).toNumber()] = _602_v297;
          _nw103[(new BigNumber(13)).toNumber()] = _602_v297;
          _nw103[(new BigNumber(14)).toNumber()] = _602_v297;
          _nw103[(new BigNumber(15)).toNumber()] = (_602_v297).Merge(_602_v297);
          _nw103[(new BigNumber(16)).toNumber()] = _602_v297;
          _nw103[(new BigNumber(17)).toNumber()] = _602_v297;
          _nw103[(new BigNumber(18)).toNumber()] = (_602_v297).Merge((_605_v300)[_module.__default.safeIndex(new BigNumber((_608_v303).length), new BigNumber((_605_v300).length))]);
          _609_v304 = _nw103;
          let _610_v305;
          let _init12 = ((_611_v1) => function (_612_i30) {
            return _module.D9.create_DC26(_611_v1);
          })(_175_v1);
          let _nw104 = Array((new BigNumber(19)).toNumber());
          for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw104.length); _i0_12++) {
            _nw104[_i0_12] = _init12(new BigNumber(_i0_12));
          }
          _610_v305 = _nw104;
          let _rhs87 = new BigNumber((_dafny.MultiSet.fromElements(_610_v305)).cardinality());
          let _rhs88 = _494_v228;
          let _rhs89 = (_484_v221).f18;
          let _rhs90 = (_486_v223).f14;
          let _rhs91 = _609_v304;
          let _lhs65 = _494_v228;
          let _lhs66 = _494_v228;
          let _lhs67 = _185_globalState;
          _lhs65.f16 = _rhs87;
          _494_v228 = _rhs88;
          _lhs66.f16 = _rhs89;
          _lhs67.f8 = _rhs90;
          _609_v304 = _rhs91;
          let _613_v306;
          _613_v306 = _dafny.Map.Empty.slice().updateUnsafe(true,_494_v228);
          let _614_v307;
          let _nw105 = Array((new BigNumber(4)).toNumber());
          _nw105[(_dafny.ZERO).toNumber()] = (((_613_v306).contains(_183_v6)) ? ((_613_v306).get(_183_v6)) : (_494_v228));
          _nw105[(_dafny.ONE).toNumber()] = _494_v228;
          _nw105[(new BigNumber(2)).toNumber()] = _494_v228;
          _nw105[(new BigNumber(3)).toNumber()] = _494_v228;
          _614_v307 = _nw105;
          _614_v307 = _614_v307;
        }
        _module.__default.m0(_185_globalState);
        _566_v286 = _566_v286;
      }
      (_185_globalState).f7 = _183_v6;
      let _615_i31;
      _615_i31 = _dafny.ZERO;
      L5: {
        while ((_484_v221).f17) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_615_i31)) {
              break L5;
            }
            _615_i31 = (_615_i31).plus(_dafny.ONE);
            (_494_v228).f16 = _175_v1;
            let _616_v308;
            let _nw106 = new _module.C2();
            _nw106.__ctor((_484_v221).f17, _485_v222, (_486_v223).f14);
            _616_v308 = _nw106;
            let _617_v309;
            _617_v309 = _module.D12.create_DC32(new BigNumber((_248_v54).length), _616_v308, _484_v221, (_486_v223).f13);
            let _618_v310;
            _618_v310 = _dafny.Map.Empty.slice().updateUnsafe(_176_v2,_617_v309);
            let _619_v311;
            _619_v311 = _dafny.Map.Empty.slice().updateUnsafe(_494_v228.f16,new BigNumber((_dafny.Set.fromElements(_616_v308.f15)).length));
            let _620_v312;
            _620_v312 = _dafny.Map.Empty.slice().updateUnsafe((_486_v223).f13,new BigNumber((_619_v311).length));
            _618_v310 = (_618_v310).update(_176_v2, _module.D12.create_DC32(_494_v228.f16, _616_v308, _484_v221, (((_619_v311).contains(new BigNumber(-252))) ? ((_619_v311).get(new BigNumber(-252))) : ((_484_v221).f18))));
            let _621_v313;
            _621_v313 = _dafny.Seq.of(_484_v221);
            let _622_v314;
            _622_v314 = _dafny.MultiSet.fromElements(new BigNumber((_621_v313).length));
            let _623_v315;
            _623_v315 = _module.D9.create_DC25(_248_v54, _616_v308.f15, (_484_v221).f18);
            let _624_v316;
            _624_v316 = _dafny.Map.Empty.slice().updateUnsafe((_494_v228).fm2(_622_v314, _494_v228.f16, _dafny.Map.Empty.slice().updateUnsafe((_623_v315).dtor_cf60,_494_v228.f16), (_486_v223).f14, _185_globalState),new BigNumber(373));
            let _625_v317;
            let _nw107 = new _module.C4();
            _nw107.__ctor(_485_v222, (_484_v221).f17);
            _625_v317 = _nw107;
            let _626_v318;
            _626_v318 = _dafny.Map.Empty.slice().updateUnsafe((((_624_v316).contains(_616_v308.f15)) ? ((_624_v316).get(_616_v308.f15)) : (new BigNumber(312))),_625_v317);
            _626_v318 = (_626_v318).update(new BigNumber(615), (((_486_v223).f14) ? (_625_v317) : (_625_v317)));
            let _627_v319;
            let _628_v320;
            let _out15;
            let _out16;
            let _outcollector6 = (_486_v223).m2(_248_v54, (_486_v223).f14, _490_v227, _185_globalState);
            _out15 = _outcollector6[0];
            _out16 = _outcollector6[1];
            _627_v319 = _out15;
            _628_v320 = _out16;
          }
        }
      }
      process.stdout.write(_dafny.toString(_175_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_176_v2, _dafny.Seq.of(new BigNumber(-351), new BigNumber(231), new BigNumber(-351), new BigNumber(-351), new BigNumber(231), new BigNumber(-351)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_177_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_178_v4).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,false).updateUnsafe(new BigNumber(582),false),new _dafny.CodePoint('p'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v5)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v5)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v5)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v5)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v5)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v5)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v5)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v5)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v5)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v5)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v5)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v5)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v5)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v5)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v5)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v5)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v5)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v5)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v5)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v5)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v5)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v5)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v5)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v5)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v5)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v5)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v5)[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v5)[new BigNumber(27)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v5)[new BigNumber(28)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_183_v6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_184_v7, _dafny.Seq.of(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_185_globalState).f0).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,false).updateUnsafe(new BigNumber(582),false),new _dafny.CodePoint('p'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f1)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f1)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f1)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f1)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f1)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f1)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f1)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f1)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f1)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f1)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f1)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f1)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f1)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f1)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f1)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f1)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f1)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f1)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f1)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f1)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f1)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f1)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f1)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f1)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f1)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f1)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f1)[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f1)[new BigNumber(27)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f1)[new BigNumber(28)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_185_globalState.f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f4)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f4)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f4)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f4)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f4)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f4)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f4)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f4)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f4)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f4)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f4)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f4)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f4)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f4)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f4)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f4)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f4)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f4)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f4)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f4)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f4)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f4)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f4)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f4)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f4)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f4)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f4)[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f4)[new BigNumber(27)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState.f4)[new BigNumber(28)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_185_globalState.f6, _dafny.Seq.of(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_185_globalState.f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_185_globalState.f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_185_globalState.f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_186_i1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_205_i4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_248_v54).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_249_v55).equals(_dafny.Set.fromElements(new BigNumber(-351)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_250_v56).dtor_cf40));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_250_v56).dtor_cf41).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_250_v56).dtor_cf42));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_250_v56).dtor_cf43));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_250_v56).dtor_cf44).equals(_dafny.Set.fromElements(new BigNumber(-351)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_251_v57, _dafny.Seq.of(_dafny.Seq.UnicodeFromString("eyxck")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_369_v140).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("mupivve"),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_483_v220).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_484_v221).f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_484_v221).f18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_485_v222).equals(_dafny.MultiSet.fromElements(false, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_486_v223).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_486_v223).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_487_v224).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_488_v225, _dafny.Seq.of(_dafny.MultiSet.fromElements(false, false, false), _dafny.MultiSet.fromElements(false, false, false), _dafny.MultiSet.fromElements(false, false, false), _dafny.MultiSet.fromElements(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_489_v226).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_490_v227)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_490_v227)[_dafny.ONE]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_490_v227)[new BigNumber(2)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_490_v227)[new BigNumber(3)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_490_v227)[new BigNumber(4)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_490_v227)[new BigNumber(5)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_490_v227)[new BigNumber(6)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_490_v227)[new BigNumber(7)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_490_v227)[new BigNumber(8)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_490_v227)[new BigNumber(9)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_490_v227)[new BigNumber(10)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_490_v227)[new BigNumber(11)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_490_v227)[new BigNumber(12)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_490_v227)[new BigNumber(13)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_490_v227)[new BigNumber(14)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_490_v227)[new BigNumber(15)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_490_v227)[new BigNumber(16)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_490_v227)[new BigNumber(17)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_490_v227)[new BigNumber(18)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_490_v227)[new BigNumber(19)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_490_v227)[new BigNumber(20)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_490_v227)[new BigNumber(21)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_490_v227)[new BigNumber(22)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_490_v227)[new BigNumber(23)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_490_v227)[new BigNumber(24)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_490_v227)[new BigNumber(25)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_490_v227)[new BigNumber(26)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_494_v228.f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_495_v229).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_496_i23));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_615_i31));
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
      return false;
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
    static create_DC2(cf2, cf3, cf4, cf5, cf6) {
      let $dt = new D1(0);
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC1(cf1) {
      let $dt = new D1(1);
      $dt.cf1 = cf1;
      return $dt;
    }
    static create_DC3(cf7) {
      let $dt = new D1(2);
      $dt.cf7 = cf7;
      return $dt;
    }
    get is_DC2() { return this.$tag === 0; }
    get is_DC1() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf7() { return this.cf7; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf7) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf2 === other.cf2 && this.cf3 === other.cf3 && this.cf4 === other.cf4 && _dafny.areEqual(this.cf5, other.cf5) && _dafny.areEqual(this.cf6, other.cf6);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf1, other.cf1);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf7, other.cf7);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC2(false, false, false, new _dafny.CodePoint('D'.codePointAt(0)), new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC5(cf9, cf10, cf11, cf12, cf13) {
      let $dt = new D2(0);
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC4(cf8) {
      let $dt = new D2(1);
      $dt.cf8 = cf8;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC4() { return this.$tag === 1; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf8() { return this.cf8; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC4" + "(" + _dafny.toString(this.cf8) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf9, other.cf9) && _dafny.areEqual(this.cf10, other.cf10) && _dafny.areEqual(this.cf11, other.cf11) && _dafny.areEqual(this.cf12, other.cf12) && _dafny.areEqual(this.cf13, other.cf13);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf8, other.cf8);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC5(_dafny.ZERO, _dafny.ZERO, _dafny.ZERO, _dafny.ZERO, _dafny.Seq.of());
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
    static create_DC7(cf15, cf16, cf17) {
      let $dt = new D3(0);
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC8(cf18, cf19) {
      let $dt = new D3(1);
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      return $dt;
    }
    static create_DC9(cf20, cf21, cf22, cf23) {
      let $dt = new D3(2);
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC6(cf14) {
      let $dt = new D3(3);
      $dt.cf14 = cf14;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC8() { return this.$tag === 1; }
    get is_DC9() { return this.$tag === 2; }
    get is_DC6() { return this.$tag === 3; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf14() { return this.cf14; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC7" + "(" + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 3) {
        return "D3.DC6" + "(" + _dafny.toString(this.cf14) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf15, other.cf15) && _dafny.areEqual(this.cf16, other.cf16) && this.cf17 === other.cf17;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf18, other.cf18) && _dafny.areEqual(this.cf19, other.cf19);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf20, other.cf20) && _dafny.areEqual(this.cf21, other.cf21) && this.cf22 === other.cf22 && _dafny.areEqual(this.cf23, other.cf23);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf14 === other.cf14;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC7(_dafny.ZERO, _dafny.ZERO, null);
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
    static create_DC11(cf25, cf26, cf27, cf28) {
      let $dt = new D4(0);
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC10(cf24) {
      let $dt = new D4(1);
      $dt.cf24 = cf24;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf24() { return this.cf24; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf24) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf25 === other.cf25 && _dafny.areEqual(this.cf26, other.cf26) && _dafny.areEqual(this.cf27, other.cf27) && _dafny.areEqual(this.cf28, other.cf28);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf24 === other.cf24;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC11(false, _dafny.ZERO, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC13(cf30, cf31, cf32, cf33, cf34) {
      let $dt = new D5(0);
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      return $dt;
    }
    static create_DC14() {
      let $dt = new D5(1);
      return $dt;
    }
    static create_DC12(cf29) {
      let $dt = new D5(2);
      $dt.cf29 = cf29;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get is_DC12() { return this.$tag === 2; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf29() { return this.cf29; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC14";
      } else if (this.$tag === 2) {
        return "D5.DC12" + "(" + _dafny.toString(this.cf29) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf30, other.cf30) && this.cf31 === other.cf31 && _dafny.areEqual(this.cf32, other.cf32) && this.cf33 === other.cf33 && this.cf34 === other.cf34;
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf29, other.cf29);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC13(_dafny.ZERO, false, _dafny.ZERO, [], false);
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
    static create_DC16(cf36, cf37, cf38) {
      let $dt = new D6(0);
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC15(cf35) {
      let $dt = new D6(1);
      $dt.cf35 = cf35;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf35() { return this.cf35; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC15" + "(" + _dafny.toString(this.cf35) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf36, other.cf36) && this.cf37 === other.cf37 && this.cf38 === other.cf38;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf35, other.cf35);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC16(_dafny.ZERO, false, false);
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
    static create_DC18(cf40, cf41, cf42, cf43, cf44) {
      let $dt = new D7(0);
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      $dt.cf42 = cf42;
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      return $dt;
    }
    static create_DC17(cf39) {
      let $dt = new D7(1);
      $dt.cf39 = cf39;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get is_DC17() { return this.$tag === 1; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf39() { return this.cf39; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC18" + "(" + _dafny.toString(this.cf40) + ", " + this.cf41.toVerbatimString(true) + ", " + _dafny.toString(this.cf42) + ", " + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC17" + "(" + _dafny.toString(this.cf39) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf40, other.cf40) && _dafny.areEqual(this.cf41, other.cf41) && this.cf42 === other.cf42 && this.cf43 === other.cf43 && _dafny.areEqual(this.cf44, other.cf44);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf39, other.cf39);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC18(_dafny.ZERO, _dafny.Seq.UnicodeFromString(""), false, false, _dafny.Set.Empty);
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
    static create_DC20(cf46, cf47, cf48) {
      let $dt = new D8(0);
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      return $dt;
    }
    static create_DC21(cf49, cf50, cf51, cf52, cf53) {
      let $dt = new D8(1);
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      return $dt;
    }
    static create_DC19(cf45) {
      let $dt = new D8(2);
      $dt.cf45 = cf45;
      return $dt;
    }
    static create_DC22(cf54) {
      let $dt = new D8(3);
      $dt.cf54 = cf54;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get is_DC21() { return this.$tag === 1; }
    get is_DC19() { return this.$tag === 2; }
    get is_DC22() { return this.$tag === 3; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf54() { return this.cf54; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC20" + "(" + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC21" + "(" + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ", " + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ", " + _dafny.toString(this.cf53) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC19" + "(" + _dafny.toString(this.cf45) + ")";
      } else if (this.$tag === 3) {
        return "D8.DC22" + "(" + _dafny.toString(this.cf54) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf46 === other.cf46 && this.cf47 === other.cf47 && this.cf48 === other.cf48;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf49 === other.cf49 && this.cf50 === other.cf50 && _dafny.areEqual(this.cf51, other.cf51) && this.cf52 === other.cf52 && _dafny.areEqual(this.cf53, other.cf53);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf45, other.cf45);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf54, other.cf54);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC20(false, false, false);
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
    static create_DC24(cf56, cf57, cf58) {
      let $dt = new D9(0);
      $dt.cf56 = cf56;
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      return $dt;
    }
    static create_DC25(cf59, cf60, cf61) {
      let $dt = new D9(1);
      $dt.cf59 = cf59;
      $dt.cf60 = cf60;
      $dt.cf61 = cf61;
      return $dt;
    }
    static create_DC26(cf62) {
      let $dt = new D9(2);
      $dt.cf62 = cf62;
      return $dt;
    }
    static create_DC23(cf55) {
      let $dt = new D9(3);
      $dt.cf55 = cf55;
      return $dt;
    }
    get is_DC24() { return this.$tag === 0; }
    get is_DC25() { return this.$tag === 1; }
    get is_DC26() { return this.$tag === 2; }
    get is_DC23() { return this.$tag === 3; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf55() { return this.cf55; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC24" + "(" + this.cf56.toVerbatimString(true) + ", " + _dafny.toString(this.cf57) + ", " + _dafny.toString(this.cf58) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC25" + "(" + this.cf59.toVerbatimString(true) + ", " + _dafny.toString(this.cf60) + ", " + _dafny.toString(this.cf61) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC26" + "(" + _dafny.toString(this.cf62) + ")";
      } else if (this.$tag === 3) {
        return "D9.DC23" + "(" + _dafny.toString(this.cf55) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf56, other.cf56) && this.cf57 === other.cf57 && _dafny.areEqual(this.cf58, other.cf58);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf59, other.cf59) && this.cf60 === other.cf60 && _dafny.areEqual(this.cf61, other.cf61);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf62, other.cf62);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf55, other.cf55);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC24(_dafny.Seq.UnicodeFromString(""), false, _dafny.ZERO);
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
    static create_DC28(cf64, cf65, cf66, cf67) {
      let $dt = new D10(0);
      $dt.cf64 = cf64;
      $dt.cf65 = cf65;
      $dt.cf66 = cf66;
      $dt.cf67 = cf67;
      return $dt;
    }
    static create_DC29(cf68) {
      let $dt = new D10(1);
      $dt.cf68 = cf68;
      return $dt;
    }
    static create_DC27(cf63) {
      let $dt = new D10(2);
      $dt.cf63 = cf63;
      return $dt;
    }
    get is_DC28() { return this.$tag === 0; }
    get is_DC29() { return this.$tag === 1; }
    get is_DC27() { return this.$tag === 2; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf63() { return this.cf63; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC28" + "(" + _dafny.toString(this.cf64) + ", " + _dafny.toString(this.cf65) + ", " + _dafny.toString(this.cf66) + ", " + _dafny.toString(this.cf67) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC29" + "(" + _dafny.toString(this.cf68) + ")";
      } else if (this.$tag === 2) {
        return "D10.DC27" + "(" + _dafny.toString(this.cf63) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf64 === other.cf64 && this.cf65 === other.cf65 && this.cf66 === other.cf66 && this.cf67 === other.cf67;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf68 === other.cf68;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf63, other.cf63);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC28(null, false, false, false);
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
    static create_DC30(cf69) {
      let $dt = new D11(0);
      $dt.cf69 = cf69;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get dtor_cf69() { return this.cf69; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC30" + "(" + _dafny.toString(this.cf69) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf69, other.cf69);
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
          return D11.Default();
        }
      };
    }
  }

  $module.D12 = class D12 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC32(cf71, cf72, cf73, cf74) {
      let $dt = new D12(0);
      $dt.cf71 = cf71;
      $dt.cf72 = cf72;
      $dt.cf73 = cf73;
      $dt.cf74 = cf74;
      return $dt;
    }
    static create_DC31(cf70) {
      let $dt = new D12(1);
      $dt.cf70 = cf70;
      return $dt;
    }
    static create_DC33(cf75) {
      let $dt = new D12(2);
      $dt.cf75 = cf75;
      return $dt;
    }
    get is_DC32() { return this.$tag === 0; }
    get is_DC31() { return this.$tag === 1; }
    get is_DC33() { return this.$tag === 2; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf73() { return this.cf73; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf75() { return this.cf75; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC32" + "(" + _dafny.toString(this.cf71) + ", " + _dafny.toString(this.cf72) + ", " + _dafny.toString(this.cf73) + ", " + _dafny.toString(this.cf74) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC31" + "(" + _dafny.toString(this.cf70) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC33" + "(" + _dafny.toString(this.cf75) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf71, other.cf71) && this.cf72 === other.cf72 && this.cf73 === other.cf73 && _dafny.areEqual(this.cf74, other.cf74);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf70, other.cf70);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf75, other.cf75);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC32(_dafny.ZERO, null, null, _dafny.ZERO);
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
    static create_DC34(cf76) {
      let $dt = new D13(0);
      $dt.cf76 = cf76;
      return $dt;
    }
    get is_DC34() { return this.$tag === 0; }
    get dtor_cf76() { return this.cf76; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC34" + "(" + _dafny.toString(this.cf76) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf76, other.cf76);
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
          return D13.Default();
        }
      };
    }
  }

  $module.D14 = class D14 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC35(cf77) {
      let $dt = new D14(0);
      $dt.cf77 = cf77;
      return $dt;
    }
    get is_DC35() { return this.$tag === 0; }
    get dtor_cf77() { return this.cf77; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC35" + "(" + _dafny.toString(this.cf77) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf77, other.cf77);
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
          return D14.Default();
        }
      };
    }
  }

  $module.D15 = class D15 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC37(cf79, cf80, cf81, cf82) {
      let $dt = new D15(0);
      $dt.cf79 = cf79;
      $dt.cf80 = cf80;
      $dt.cf81 = cf81;
      $dt.cf82 = cf82;
      return $dt;
    }
    static create_DC36(cf78) {
      let $dt = new D15(1);
      $dt.cf78 = cf78;
      return $dt;
    }
    get is_DC37() { return this.$tag === 0; }
    get is_DC36() { return this.$tag === 1; }
    get dtor_cf79() { return this.cf79; }
    get dtor_cf80() { return this.cf80; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf82() { return this.cf82; }
    get dtor_cf78() { return this.cf78; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC37" + "(" + _dafny.toString(this.cf79) + ", " + _dafny.toString(this.cf80) + ", " + _dafny.toString(this.cf81) + ", " + _dafny.toString(this.cf82) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC36" + "(" + _dafny.toString(this.cf78) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf79, other.cf79) && this.cf80 === other.cf80 && this.cf81 === other.cf81 && _dafny.areEqual(this.cf82, other.cf82);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf78, other.cf78);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC37(_dafny.Seq.of(), false, false, new _dafny.CodePoint('D'.codePointAt(0)));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D15.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f1 = [];
      this.f3 = false;
      this.f4 = [];
      this.f6 = _dafny.Seq.of();
      this.f7 = false;
      this.f8 = false;
      this.f9 = _dafny.ZERO;
      this._f0 = _dafny.Map.Empty;
      this._f2 = false;
      this._f5 = _dafny.ZERO;
      this._f10 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10) {
      let _this = this;
      (_this)._f0 = f0;
      (_this).f1 = f1;
      (_this)._f2 = f2;
      (_this).f3 = f3;
      (_this).f4 = f4;
      (_this)._f5 = f5;
      (_this).f6 = f6;
      (_this).f7 = f7;
      (_this).f8 = f8;
      (_this).f9 = f9;
      (_this)._f10 = f10;
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
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f10() {
      let _this = this;
      return _this._f10;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f11 = _dafny.MultiSet.Empty;
      this._f12 = false;
      this._f17 = false;
      this._f18 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
    set f11(value) {
      let _this = this;
      _this._f11 = value;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
    set f12(value) {
      let _this = this;
      _this._f12 = value;
    };
    __ctor(f17, f18, f11, f12) {
      let _this = this;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      return;
    }
    fm2(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _this.f12;
    };
    fm11(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt((_this).f18, (_this).f18),_module.D1.create_DC2(_this.f12, false, (_this).f17, new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0))));
    };
    fm12(globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt((_this).f18, new BigNumber(291));
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
    get f18() {
      let _this = this;
      return _this._f18;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f11 = _dafny.MultiSet.Empty;
      this._f12 = false;
      this.f16 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
    set f11(value) {
      let _this = this;
      _this._f11 = value;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
    set f12(value) {
      let _this = this;
      _this._f12 = value;
    };
    __ctor(f16, f11, f12) {
      let _this = this;
      (_this).f16 = f16;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      return;
    }
    fm2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _source6 = _module.D2.create_DC5(_this.f16, _this.f16, _this.f16, new BigNumber(-458), _dafny.Seq.of(_this.f12));
      if (_source6.is_DC5) {
        let _629___mcc_h0 = (_source6).cf9;
        let _630___mcc_h1 = (_source6).cf10;
        let _631___mcc_h2 = (_source6).cf11;
        let _632___mcc_h3 = (_source6).cf12;
        let _633___mcc_h4 = (_source6).cf13;
        let _634_cf13 = _633___mcc_h4;
        let _635_cf12 = _632___mcc_h3;
        let _636_cf11 = _631___mcc_h2;
        let _637_cf10 = _630___mcc_h1;
        let _638_cf9 = _629___mcc_h0;
        return _this.f12;
      } else {
        let _639___mcc_h5 = (_source6).cf8;
        let _640_cf8 = _639___mcc_h5;
        return !(((_this.f12) ? (_this.f12) : (false)));
      }
    };
    m6(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _641_i0;
      _641_i0 = _dafny.ZERO;
      L6: {
        while (_this.f12) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_641_i0)) {
              break L6;
            }
            _641_i0 = (_641_i0).plus(_dafny.ONE);
            if ((((_this.f12) ? (_this.f16) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f12,_this.f12)).length)))).isEqualTo(_module.__default.safeDivisionInt(_this.f16, _this.f16))) {
              let _642_v0;
              _642_v0 = _module.D2.create_DC4(new BigNumber(43));
              (globalState).f9 = _module.__default.safeModuloInt((_642_v0).dtor_cf8, _this.f16);
              let _643_v1;
              let _init13 = function (_644_i1) {
                return (_644_i1).plus(_this.f16);
              };
              let _nw108 = Array((new BigNumber(27)).toNumber());
              for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw108.length); _i0_13++) {
                _nw108[_i0_13] = _init13(new BigNumber(_i0_13));
              }
              _643_v1 = _nw108;
              let _645_v2;
              _645_v2 = _dafny.Seq.of(_643_v1);
              let _646_v3;
              _646_v3 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_645_v2, _645_v2), _module.__default.safeIndex(_this.f16, new BigNumber((_dafny.Seq.Concat(_645_v2, _645_v2)).length)), (_645_v2)[_module.__default.safeIndex(_this.f16, new BigNumber((_645_v2).length))])).length));
              _646_v3 = (_646_v3).Merge(_646_v3);
              let _index68 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_643_v1).length));
              (_643_v1)[_index68] = (_dafny.ZERO).minus(_this.f16);
              let _index69 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_643_v1).length));
              (_643_v1)[_index69] = new BigNumber(276);
              let _647_v4;
              let _nw109 = Array((new BigNumber(3)).toNumber()).fill(false);
              _647_v4 = _nw109;
              _647_v4 = _647_v4;
              let _648_v5;
              let _nw110 = Array((new BigNumber(27)).toNumber()).fill([]);
              _648_v5 = _nw110;
              let _index70 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_648_v5).length));
              (_648_v5)[_index70] = _647_v4;
              let _index71 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_647_v4).length));
              (_647_v4)[_index71] = false;
              let _649_v6;
              let _nw111 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _649_v6 = _nw111;
              let _650_v7;
              _650_v7 = new _dafny.CodePoint('t'.codePointAt(0));
              let _651_v8;
              _651_v8 = _module.D1.create_DC2(_this.f12, _this.f12, true, _650_v7, _650_v7);
              let _index72 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_649_v6).length));
              (_649_v6)[_index72] = _module.__default.fm9((_643_v1)[_module.__default.safeIndex(new BigNumber(524), new BigNumber((_643_v1).length))], (_651_v8).dtor_cf3, globalState);
              let _652_v9;
              _652_v9 = _dafny.Seq.UnicodeFromString("xcdmomb");
              let _653_v10;
              _653_v10 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(248)), ((_654_v7) => function (_655_i2) {
                return _654_v7;
              })(_650_v7)), _dafny.Seq.UnicodeFromString("cqudam"), _652_v9);
              let _656_v11;
              _656_v11 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,_this.f12);
              let _657_v12;
              _657_v12 = _dafny.Seq.of(new BigNumber((_656_v11).length), (_643_v1)[_module.__default.safeIndex(new BigNumber(524), new BigNumber((_643_v1).length))], new BigNumber((_652_v9).length), (_643_v1)[_module.__default.safeIndex(new BigNumber(524), new BigNumber((_643_v1).length))]);
              let _658_v13;
              _658_v13 = _dafny.Seq.of(_652_v9);
              let _659_v14;
              _659_v14 = _dafny.MultiSet.fromElements(_this.f16);
              let _660_v15;
              _660_v15 = _dafny.Map.Empty.slice().updateUnsafe((_643_v1)[_module.__default.safeIndex(new BigNumber(524), new BigNumber((_643_v1).length))],_this.f16);
              let _index73 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_648_v5).length));
              let _index74 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_647_v4).length));
              let _index75 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_649_v6).length));
              let _rhs92 = _647_v4;
              let _rhs93 = (_653_v10).IsProperSubsetOf((_653_v10).Difference(_653_v10));
              let _rhs94 = _dafny.areEqual(_dafny.Seq.of(new BigNumber((_656_v11).length)), _657_v12);
              let _rhs95 = _dafny.Seq.Concat(_652_v9, (_658_v13)[_module.__default.safeIndex(_module.__default.fm1((_this).fm2(_659_v14, _this.f16, _dafny.Map.Empty.slice().updateUnsafe(_this.f12,new BigNumber(-762)), _this.f12, globalState), (_643_v1)[_module.__default.safeIndex(new BigNumber(524), new BigNumber((_643_v1).length))], new BigNumber((_660_v15).length), globalState), new BigNumber((_658_v13).length))]);
              let _lhs68 = _648_v5;
              let _lhs69 = _module.__default.safeIndex(new BigNumber(596), new BigNumber((_648_v5).length));
              let _lhs70 = _this;
              let _lhs71 = _647_v4;
              let _lhs72 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_647_v4).length));
              let _lhs73 = _649_v6;
              let _lhs74 = _module.__default.safeIndex(new BigNumber(789), new BigNumber((_649_v6).length));
              _lhs68[_lhs69] = _rhs92;
              _lhs70.f12 = _rhs93;
              _lhs71[_lhs72] = _rhs94;
              _lhs73[_lhs74] = _rhs95;
            } else {
              let _661_v16;
              _661_v16 = new _dafny.CodePoint('t'.codePointAt(0));
              let _662_v17;
              _662_v17 = _dafny.Seq.of(_661_v16);
              let _663_v18;
              _663_v18 = _dafny.Seq.of(_this.f12, !_dafny.Seq.contains(_662_v17, new _dafny.CodePoint('p'.codePointAt(0))), _this.f12, _this.f12);
              (globalState).f6 = _dafny.Seq.update(_663_v18, _module.__default.safeIndex((_this.f16).multipliedBy(_this.f16), new BigNumber((_663_v18).length)), true);
              let _664_v19;
              _664_v19 = _dafny.MultiSet.fromElements(new BigNumber((_662_v17).length));
              let _665_v20;
              let _nw112 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
              _665_v20 = _nw112;
              let _666_v21;
              _666_v21 = _dafny.Map.Empty.slice().updateUnsafe(_663_v18,_665_v20);
              (globalState).f9 = ((!((_663_v18)[_module.__default.safeIndex(_this.f16, new BigNumber((_663_v18).length))])) ? ((((_664_v19).contains(new BigNumber((_666_v21).length))) ? ((_664_v19).get(new BigNumber((_666_v21).length))) : (new BigNumber(354)))) : (new BigNumber(-844)));
              (globalState).f3 = !(_this.f12);
              (globalState).f9 = _this.f16;
              let _667_v22;
              _667_v22 = _module.D1.create_DC2(_this.f12, _this.f12, _this.f12, _661_v16, _661_v16);
              (globalState).f9 = (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(((_this.f12) ? (_this.f12) : (_this.f12)),(_667_v22).dtor_cf5)).length));
            }
            let _668_v23;
            _668_v23 = new _dafny.CodePoint('q'.codePointAt(0));
            let _669_v24;
            _669_v24 = _dafny.Seq.of(_module.__default.fm0(_668_v23, globalState));
            let _670_v25;
            _670_v25 = _dafny.Seq.UnicodeFromString("yfreelwss");
            let _671_v26;
            let _nw113 = Array((new BigNumber(23)).toNumber());
            _nw113[(_dafny.ZERO).toNumber()] = _this.f12;
            _nw113[(_dafny.ONE).toNumber()] = ((_this.f12) ? (_this.f12) : (false));
            _nw113[(new BigNumber(2)).toNumber()] = _this.f12;
            _nw113[(new BigNumber(3)).toNumber()] = _this.f12;
            _nw113[(new BigNumber(4)).toNumber()] = _this.f12;
            _nw113[(new BigNumber(5)).toNumber()] = true;
            _nw113[(new BigNumber(6)).toNumber()] = (_this.f12) === (!(_this.f12));
            _nw113[(new BigNumber(7)).toNumber()] = _dafny.areEqual(_669_v24, _669_v24);
            _nw113[(new BigNumber(8)).toNumber()] = _this.f12;
            _nw113[(new BigNumber(9)).toNumber()] = !(_this.f12) || (_this.f12);
            _nw113[(new BigNumber(10)).toNumber()] = (_this.f12) || (!(_this.f12));
            _nw113[(new BigNumber(11)).toNumber()] = _this.f12;
            _nw113[(new BigNumber(12)).toNumber()] = _this.f12;
            _nw113[(new BigNumber(13)).toNumber()] = (_this.f16).isLessThan(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f12,_this.f12)).length));
            _nw113[(new BigNumber(14)).toNumber()] = _this.f12;
            _nw113[(new BigNumber(15)).toNumber()] = _this.f12;
            _nw113[(new BigNumber(16)).toNumber()] = (_this.f16).isLessThan(new BigNumber((_670_v25).length));
            _nw113[(new BigNumber(17)).toNumber()] = !(_this.f12);
            _nw113[(new BigNumber(18)).toNumber()] = _this.f12;
            _nw113[(new BigNumber(19)).toNumber()] = ((_this.f12) ? (_this.f12) : (_this.f12));
            _nw113[(new BigNumber(20)).toNumber()] = _this.f12;
            _nw113[(new BigNumber(21)).toNumber()] = _this.f12;
            _nw113[(new BigNumber(22)).toNumber()] = _this.f12;
            _671_v26 = _nw113;
            let _index76 = _module.__default.safeIndex(new BigNumber(19), new BigNumber((_671_v26).length));
            (_671_v26)[_index76] = _this.f12;
            let _672_v27;
            _672_v27 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f12),_dafny.areEqual(_670_v25, _670_v25));
            let _673_v28;
            _673_v28 = _dafny.Seq.of(_669_v24, _669_v24, _dafny.Seq.of(_this.f12, _this.f12), _669_v24, _module.__default.fm10(_this.f12, _this.f12, true, globalState));
            let _674_v29;
            _674_v29 = _dafny.MultiSet.fromElements(new BigNumber(((_673_v28)[_module.__default.safeIndex(_this.f16, new BigNumber((_673_v28).length))]).length));
            let _675_v30;
            _675_v30 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,_this.f16);
            let _index77 = _module.__default.safeIndex(new BigNumber(19), new BigNumber((_671_v26).length));
            let _rhs96 = (_this.f16).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ddcprsn")).length));
            let _rhs97 = (((_672_v27).contains(_this.f12)) ? ((_672_v27).get(_this.f12)) : (!(_this.f12) || ((_this).fm2(_674_v29, _this.f16, _675_v30, _this.f12, globalState))));
            let _lhs75 = globalState;
            let _lhs76 = _671_v26;
            let _lhs77 = _module.__default.safeIndex(new BigNumber(19), new BigNumber((_671_v26).length));
            _lhs75.f9 = _rhs96;
            _lhs76[_lhs77] = _rhs97;
            let _676_v31;
            let _nw114 = Array((new BigNumber(2)).toNumber()).fill([]);
            _676_v31 = _nw114;
            let _index78 = _module.__default.safeIndex(new BigNumber(811), new BigNumber((_676_v31).length));
            (_676_v31)[_index78] = _671_v26;
            let _index79 = _module.__default.safeIndex(new BigNumber(811), new BigNumber((_676_v31).length));
            let _rhs98 = _674_v29;
            let _rhs99 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("nvme"), _670_v25), _module.__default.safeIndex((_this.f16).minus(_this.f16), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("nvme"), _670_v25)).length)), _668_v23);
            let _rhs100 = _671_v26;
            let _rhs101 = _671_v26;
            let _lhs78 = _676_v31;
            let _lhs79 = _module.__default.safeIndex(new BigNumber(811), new BigNumber((_676_v31).length));
            _674_v29 = _rhs98;
            _670_v25 = _rhs99;
            _lhs78[_lhs79] = _rhs100;
            _671_v26 = _rhs101;
            let _677_v32;
            _677_v32 = _dafny.MultiSet.fromElements(_668_v23, _668_v23, _668_v23);
            let _678_v33;
            _678_v33 = _dafny.Seq.of(_this.f16, _this.f16);
            let _679_v34;
            _679_v34 = _dafny.Map.Empty.slice().updateUnsafe(_677_v32,_678_v33);
            _679_v34 = (_679_v34).update(_677_v32, _dafny.Seq.update(_678_v33, _module.__default.safeIndex(_this.f16, new BigNumber((_678_v33).length)), _this.f16));
          }
        }
      }
      let _680_v35;
      let _nw115 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
      _680_v35 = _nw115;
      let _index80 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_680_v35).length));
      (_680_v35)[_index80] = _this.f16;
      let _681_v36;
      _681_v36 = _dafny.Seq.UnicodeFromString("bso");
      let _index81 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_680_v35).length));
      (_680_v35)[_index81] = _module.__default.safeDivisionInt(new BigNumber((_681_v36).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(52)), function (_682_i3) {
        return _module.D1.create_DC3(_module.D1.create_DC2(_this.f12, _this.f12, _this.f12, new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0))));
      })).length));
      let _683_v37;
      _683_v37 = _dafny.Map.Empty.slice().updateUnsafe(true,_this.f12);
      _683_v37 = _683_v37;
      let _684_v38;
      _684_v38 = _module.D2.create_DC4((_dafny.ZERO).minus(_this.f16));
      let _685_v39;
      let _init14 = function (_686_i4) {
        return false;
      };
      let _nw116 = Array((new BigNumber(28)).toNumber());
      for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw116.length); _i0_14++) {
        _nw116[_i0_14] = _init14(new BigNumber(_i0_14));
      }
      _685_v39 = _nw116;
      let _687_v40;
      _687_v40 = _dafny.Map.Empty.slice().updateUnsafe(((_this.f12) ? (_684_v38) : (_684_v38)),_685_v39);
      let _index82 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_680_v35).length));
      let _rhs102 = (_this.f16).plus((_this.f16).multipliedBy((_680_v35)[_module.__default.safeIndex(new BigNumber(829), new BigNumber((_680_v35).length))]));
      let _rhs103 = _687_v40;
      let _lhs80 = _680_v35;
      let _lhs81 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_680_v35).length));
      _lhs80[_lhs81] = _rhs102;
      _687_v40 = _rhs103;
      let _688_v41;
      _688_v41 = _dafny.Seq.of(_this.f12);
      let _689_v42;
      _689_v42 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,_688_v41);
      let _690_v43;
      let _nw117 = Array((new BigNumber(10)).toNumber());
      _nw117[(_dafny.ZERO).toNumber()] = _688_v41;
      _nw117[(_dafny.ONE).toNumber()] = _688_v41;
      _nw117[(new BigNumber(2)).toNumber()] = _688_v41;
      _nw117[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_688_v41, _module.__default.safeIndex((_680_v35)[_module.__default.safeIndex(new BigNumber(829), new BigNumber((_680_v35).length))], new BigNumber((_688_v41).length)), _this.f12);
      _nw117[(new BigNumber(4)).toNumber()] = _688_v41;
      _nw117[(new BigNumber(5)).toNumber()] = _688_v41;
      _nw117[(new BigNumber(6)).toNumber()] = _module.__default.fm10(!(_this.f12), _this.f12, (_688_v41)[_module.__default.safeIndex(_this.f16, new BigNumber((_688_v41).length))], globalState);
      _nw117[(new BigNumber(7)).toNumber()] = _688_v41;
      _nw117[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat((((_689_v42).contains(_module.__default.fm1(_this.f12, _this.f16, _this.f16, globalState))) ? ((_689_v42).get(_module.__default.fm1(_this.f12, _this.f16, _this.f16, globalState))) : (_688_v41)), _688_v41);
      _nw117[(new BigNumber(9)).toNumber()] = _dafny.Seq.of(_this.f12);
      _690_v43 = _nw117;
      let _691_v44;
      _691_v44 = _this.f12;
      let _692_v45;
      _692_v45 = _dafny.Set.fromElements(_this.f12);
      let _693_v46;
      _693_v46 = _dafny.Seq.of(_688_v41, _688_v41, _dafny.Seq.update(_dafny.Seq.of(_this.f12, (_691_v44), _this.f12), _module.__default.safeIndex(new BigNumber((_692_v45).length), new BigNumber((_dafny.Seq.of(_this.f12, (_691_v44), _this.f12)).length)), _this.f12), _688_v41, _688_v41);
      let _694_v47;
      _694_v47 = _dafny.Map.Empty.slice().updateUnsafe((_680_v35)[_module.__default.safeIndex(new BigNumber(829), new BigNumber((_680_v35).length))],_this.f12);
      let _index83 = _module.__default.safeIndex(new BigNumber(21), new BigNumber((_690_v43).length));
      (_690_v43)[_index83] = (_693_v46)[_module.__default.safeIndex(new BigNumber((_694_v47).length), new BigNumber((_693_v46).length))];
      let _index84 = _module.__default.safeIndex(new BigNumber(21), new BigNumber((_690_v43).length));
      (_690_v43)[_index84] = _688_v41;
      _691_v44 = _691_v44;
      r0 = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(_this.f12),(_680_v35)[_module.__default.safeIndex(new BigNumber(829), new BigNumber((_680_v35).length))])).length);
      return r0;
    }
    m7(p0, p1, globalState) {
      let _this = this;
      let r0 = [];
      let _695_v0;
      _695_v0 = _this.f12;
      let _696_v1;
      _696_v1 = new _dafny.CodePoint('s'.codePointAt(0));
      let _697_v2;
      _697_v2 = _module.D1.create_DC2((_695_v0), _this.f12, _this.f12, _696_v1, new _dafny.CodePoint('w'.codePointAt(0)));
      let _source7 = _697_v2;
      if (_source7.is_DC2) {
        let _698___mcc_h0 = (_source7).cf2;
        let _699___mcc_h1 = (_source7).cf3;
        let _700___mcc_h2 = (_source7).cf4;
        let _701___mcc_h3 = (_source7).cf5;
        let _702___mcc_h4 = (_source7).cf6;
        let _703_cf6 = _702___mcc_h4;
        let _704_cf5 = _701___mcc_h3;
        let _705_cf4 = _700___mcc_h2;
        let _706_cf3 = _699___mcc_h1;
        let _707_cf2 = _698___mcc_h0;
        (globalState).f9 = _this.f16;
        let _708_v3;
        _708_v3 = _dafny.Seq.of(_this.f16, _this.f16, _this.f16);
        (_this).f16 = (_708_v3)[_module.__default.safeIndex(_this.f16, new BigNumber((_708_v3).length))];
        let _rhs104 = _706_cf3;
        let _rhs105 = _dafny.Seq.Concat(_708_v3, _dafny.Seq.Concat(_708_v3, _708_v3));
        let _rhs106 = false;
        let _lhs82 = globalState;
        let _lhs83 = globalState;
        _lhs82.f7 = _rhs104;
        _708_v3 = _rhs105;
        _lhs83.f3 = _rhs106;
        if (((p0).Difference(_dafny.MultiSet.fromElements(_this.f16))).IsProperSubsetOf(p0)) {
          (globalState).f7 = ((_706_cf3) ? (_705_cf4) : (!(_this.f12) || (_this.f12)));
          let _709_v4;
          _709_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm9(_this.f16, _706_cf3, globalState)).length),_706_cf3);
          let _710_v5;
          let _init15 = ((_711_cf4) => function (_712_i0) {
            return _711_cf4;
          })(_705_cf4);
          let _nw118 = Array((new BigNumber(6)).toNumber());
          for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw118.length); _i0_15++) {
            _nw118[_i0_15] = _init15(new BigNumber(_i0_15));
          }
          _710_v5 = _nw118;
          let _713_v6;
          _713_v6 = _dafny.Map.Empty.slice().updateUnsafe(_710_v5,_705_cf4);
          _709_v4 = (_709_v4).update((_dafny.ZERO).minus(_this.f16), (((_713_v6).contains(_710_v5)) ? ((_713_v6).get(_710_v5)) : (_705_cf4)));
          let _714_v7;
          _714_v7 = _dafny.Seq.of(_705_cf4);
          let _rhs107 = new BigNumber((_module.__default.fm9(new BigNumber(-404), _706_cf3, globalState)).length);
          let _rhs108 = _module.__default.safeDivisionInt(_this.f16, new BigNumber(302));
          let _rhs109 = _707_cf2;
          let _rhs110 = (_dafny.MultiSet.FromArray(_714_v7)).Union(((_705_cf4) ? (_this.f11) : (_dafny.MultiSet.fromElements(_707_cf2, _706_cf3, _706_cf3))));
          let _rhs111 = _706_cf3;
          let _lhs84 = globalState;
          let _lhs85 = globalState;
          let _lhs86 = _this;
          let _lhs87 = globalState;
          _lhs84.f9 = _rhs107;
          _lhs85.f9 = _rhs108;
          _706_cf3 = _rhs109;
          _lhs86.f11 = _rhs110;
          _lhs87.f3 = _rhs111;
          (globalState).f9 = _this.f16;
          let _715_v8;
          _715_v8 = _dafny.Seq.of(p1);
          let _716_v9;
          _716_v9 = _dafny.Set.fromElements(_this.f16);
          let _717_v10;
          let _nw119 = Array((new BigNumber(23)).toNumber());
          _nw119[(_dafny.ZERO).toNumber()] = p1;
          _nw119[(_dafny.ONE).toNumber()] = p1;
          _nw119[(new BigNumber(2)).toNumber()] = (_715_v8)[_module.__default.safeIndex(_this.f16, new BigNumber((_715_v8).length))];
          _nw119[(new BigNumber(3)).toNumber()] = _dafny.Seq.UnicodeFromString("exlcjc");
          _nw119[(new BigNumber(4)).toNumber()] = p1;
          _nw119[(new BigNumber(5)).toNumber()] = p1;
          _nw119[(new BigNumber(6)).toNumber()] = p1;
          _nw119[(new BigNumber(7)).toNumber()] = _module.__default.fm9(new BigNumber((_716_v9).length), _706_cf3, globalState);
          _nw119[(new BigNumber(8)).toNumber()] = p1;
          _nw119[(new BigNumber(9)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(527)), ((_718_v1) => function (_719_i1) {
            return _718_v1;
          })(_696_v1));
          _nw119[(new BigNumber(10)).toNumber()] = p1;
          _nw119[(new BigNumber(11)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(110)), ((_720_cf6) => function (_721_i2) {
            return _720_cf6;
          })(_703_cf6));
          _nw119[(new BigNumber(12)).toNumber()] = p1;
          _nw119[(new BigNumber(13)).toNumber()] = p1;
          _nw119[(new BigNumber(14)).toNumber()] = p1;
          _nw119[(new BigNumber(15)).toNumber()] = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("mw"), _module.__default.safeIndex(new BigNumber((_714_v7).length), new BigNumber((_dafny.Seq.UnicodeFromString("mw")).length)), _703_cf6);
          _nw119[(new BigNumber(16)).toNumber()] = p1;
          _nw119[(new BigNumber(17)).toNumber()] = p1;
          _nw119[(new BigNumber(18)).toNumber()] = p1;
          _nw119[(new BigNumber(19)).toNumber()] = p1;
          _nw119[(new BigNumber(20)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(793)), ((_722_v1) => function (_723_i3) {
            return _722_v1;
          })(_696_v1));
          _nw119[(new BigNumber(21)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(233)), ((_724_cf6) => function (_725_i4) {
            return _724_cf6;
          })(_703_cf6));
          _nw119[(new BigNumber(22)).toNumber()] = p1;
          _717_v10 = _nw119;
          let _726_v11;
          _726_v11 = _dafny.Map.Empty.slice().updateUnsafe(_704_cf5,_717_v10);
          let _727_v12;
          _727_v12 = _dafny.Map.Empty.slice().updateUnsafe(_707_cf2,_717_v10);
          let _728_v13;
          _728_v13 = _dafny.Seq.of(_717_v10);
          let _729_v14;
          let _nw120 = Array((new BigNumber(19)).toNumber());
          _nw120[(_dafny.ZERO).toNumber()] = _717_v10;
          _nw120[(_dafny.ONE).toNumber()] = _717_v10;
          _nw120[(new BigNumber(2)).toNumber()] = _717_v10;
          _nw120[(new BigNumber(3)).toNumber()] = (((_726_v11).contains(_696_v1)) ? ((_726_v11).get(_696_v1)) : (_717_v10));
          _nw120[(new BigNumber(4)).toNumber()] = ((_this.f12) ? (_717_v10) : (_717_v10));
          _nw120[(new BigNumber(5)).toNumber()] = _717_v10;
          _nw120[(new BigNumber(6)).toNumber()] = _717_v10;
          _nw120[(new BigNumber(7)).toNumber()] = _717_v10;
          _nw120[(new BigNumber(8)).toNumber()] = _717_v10;
          _nw120[(new BigNumber(9)).toNumber()] = (((_727_v12).contains(_705_cf4)) ? ((_727_v12).get(_705_cf4)) : (_717_v10));
          _nw120[(new BigNumber(10)).toNumber()] = _717_v10;
          _nw120[(new BigNumber(11)).toNumber()] = _717_v10;
          _nw120[(new BigNumber(12)).toNumber()] = _717_v10;
          _nw120[(new BigNumber(13)).toNumber()] = _717_v10;
          _nw120[(new BigNumber(14)).toNumber()] = ((_705_cf4) ? ((_728_v13)[_module.__default.safeIndex(_this.f16, new BigNumber((_728_v13).length))]) : (_717_v10));
          _nw120[(new BigNumber(15)).toNumber()] = _717_v10;
          _nw120[(new BigNumber(16)).toNumber()] = _717_v10;
          _nw120[(new BigNumber(17)).toNumber()] = (((_727_v12).contains(_707_cf2)) ? ((_727_v12).get(_707_cf2)) : (_717_v10));
          _nw120[(new BigNumber(18)).toNumber()] = _717_v10;
          _729_v14 = _nw120;
          let _index85 = _module.__default.safeIndex(new BigNumber(160), new BigNumber((_729_v14).length));
          (_729_v14)[_index85] = _717_v10;
          let _index86 = _module.__default.safeIndex(new BigNumber(160), new BigNumber((_729_v14).length));
          (_729_v14)[_index86] = _717_v10;
        } else {
          let _730_v15;
          _730_v15 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,_this.f16);
          (globalState).f7 = (_this).fm2(((p0).update(_this.f16, _module.__default.abs(_this.f16))).Union(p0), _module.__default.safeDivisionInt(new BigNumber(768), (_dafny.ZERO).minus(_this.f16)), _730_v15, (_this.f16).isEqualTo(_this.f16), globalState);
          _696_v1 = _696_v1;
          (globalState).f9 = (_this.f16).minus(_this.f16);
          let _731_v16;
          _731_v16 = _dafny.Seq.of(_705_cf4, _707_cf2);
          let _732_v17;
          _732_v17 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(374)), ((_733_v1) => function (_734_i5) {
            return _733_v1;
          })(_696_v1)),_dafny.Seq.Concat(_731_v16, _731_v16));
          let _735_v18;
          let _nw121 = new _module.C0();
          _nw121.__ctor(_706_cf3, _this.f16, _this.f11, _707_cf2);
          _735_v18 = _nw121;
          let _736_v19;
          _736_v19 = _dafny.Map.Empty.slice().updateUnsafe(_735_v18,_731_v16);
          _732_v17 = (_732_v17).update(((_706_cf3) ? (p1) : (p1)), (((_736_v19).contains(_735_v18)) ? ((_736_v19).get(_735_v18)) : (_731_v16)));
          let _737_v20;
          let _nw122 = Array((new BigNumber(26)).toNumber());
          _737_v20 = _nw122;
          let _738_v21;
          _738_v21 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,_this.f16);
          let _739_v22;
          _739_v22 = _dafny.Map.Empty.slice().updateUnsafe(_703_cf6,p1);
          let _rhs112 = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Concat(_731_v16, _731_v16)).length), _this.f16);
          let _rhs113 = (_731_v16)[_module.__default.safeIndex((_this.f16).multipliedBy(_this.f16), new BigNumber((_731_v16).length))];
          let _rhs114 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_705_cf4)).length)), new BigNumber((_dafny.Seq.of(_707_cf2)).length)), (_this.f16).plus(_this.f16));
          let _rhs115 = _737_v20;
          let _rhs116 = ((_707_cf2) ? (_this.f16) : ((_dafny.ZERO).minus((_this.f16).minus((((_738_v21).contains(_this.f16)) ? ((_738_v21).get(_this.f16)) : (new BigNumber((_739_v22).length)))))));
          let _lhs88 = globalState;
          let _lhs89 = globalState;
          let _lhs90 = _this;
          let _lhs91 = globalState;
          _lhs88.f9 = _rhs112;
          _lhs89.f7 = _rhs113;
          _lhs90.f16 = _rhs114;
          _737_v20 = _rhs115;
          _lhs91.f9 = _rhs116;
        }
      } else if (_source7.is_DC1) {
        let _740___mcc_h5 = (_source7).cf1;
        let _741_cf1 = _740___mcc_h5;
        let _742_v23;
        _742_v23 = _module.D1.create_DC1(_741_cf1);
        let _743_v24;
        _743_v24 = _dafny.Map.Empty.slice().updateUnsafe(_742_v23,_741_cf1);
        let _744_v25;
        let _nw123 = Array((new BigNumber(8)).toNumber());
        _nw123[(_dafny.ZERO).toNumber()] = new BigNumber((p0).cardinality());
        _nw123[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.Concat(p1, _dafny.Seq.UnicodeFromString("q"))).length);
        _nw123[(new BigNumber(2)).toNumber()] = new BigNumber(135);
        _nw123[(new BigNumber(3)).toNumber()] = _this.f16;
        _nw123[(new BigNumber(4)).toNumber()] = (_this.f16).multipliedBy(_module.__default.fm1(_this.f12, new BigNumber(467), _this.f16, globalState));
        _nw123[(new BigNumber(5)).toNumber()] = _this.f16;
        _nw123[(new BigNumber(6)).toNumber()] = _this.f16;
        _nw123[(new BigNumber(7)).toNumber()] = new BigNumber(((((_743_v24).contains(_742_v23)) ? ((_743_v24).get(_742_v23)) : (_741_cf1))).length);
        _744_v25 = _nw123;
        (globalState).f1 = _744_v25;
        let _745_v27;
        _745_v27 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,_dafny.Map.Empty.slice().updateUnsafe((function () {
          let _coll18 = new _dafny.Map();
          for (const _compr_18 of _dafny.IntegerRange(new BigNumber(893), new BigNumber(501))) {
            let _746_v26 = _compr_18;
            if (((new BigNumber(893)).isLessThanOrEqualTo(_746_v26)) && ((_746_v26).isLessThan(new BigNumber(501)))) {
              _coll18.push([(_746_v26).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f16,_696_v1)).length)),_this.f12]);
            }
          }
          return _coll18;
        }()).update(new BigNumber(73), _this.f12),_this.f16));
        let _747_v29;
        _747_v29 = _dafny.Set.fromElements(_this.f16);
        if (((_this.f12) ? (_this.f12) : ((_this.f16).isLessThan(new BigNumber(((((_745_v27).contains(new BigNumber(-670))) ? ((_745_v27).get(new BigNumber(-670))) : (_dafny.Map.Empty.slice().updateUnsafe(function () {
          let _coll19 = new _dafny.Map();
          for (const _compr_19 of (_747_v29).Elements) {
            let _748_v28 = _compr_19;
            if ((_747_v29).contains(_748_v28)) {
              _coll19.push([_module.__default.safeModuloInt(_748_v28, new BigNumber(553)),_this.f12]);
            }
          }
          return _coll19;
        }(),_this.f16)))).length))))) {
          (globalState).f7 = _this.f12;
          (_this).f12 = _this.f12;
          let _rhs117 = _module.__default.fm1(_this.f12, _this.f16, _this.f16, globalState);
          let _rhs118 = false;
          let _lhs92 = _this;
          let _lhs93 = globalState;
          _lhs92.f16 = _rhs117;
          _lhs93.f3 = _rhs118;
          (globalState).f9 = _module.__default.safeModuloInt(_module.__default.fm1(_this.f12, new BigNumber(834), (((_this.f11).contains(_this.f12)) ? ((_this.f11).get(_this.f12)) : (new BigNumber(849))), globalState), (_this.f16).plus(new BigNumber(945)));
          let _749_v30;
          let _nw124 = Array((new BigNumber(21)).toNumber()).fill([]);
          _749_v30 = _nw124;
          let _750_v31;
          _750_v31 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_this.f16),_749_v30);
          let _751_v32;
          _751_v32 = _module.D2.create_DC4(_this.f16);
          let _752_v33;
          _752_v33 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,_751_v32);
          _750_v31 = (_750_v31).update((p0).Intersect((p0).update(new BigNumber((_752_v33).length), _module.__default.abs(_this.f16))), _749_v30);
        } else {
          let _753_v34;
          _753_v34 = _dafny.Seq.UnicodeFromString("pekwurnd");
          let _754_v35;
          _754_v35 = _dafny.Map.Empty.slice().updateUnsafe(_696_v1,_dafny.Seq.UnicodeFromString("tfxhr"));
          _753_v34 = (((_754_v35).contains(new _dafny.CodePoint('w'.codePointAt(0)))) ? ((_754_v35).get(new _dafny.CodePoint('w'.codePointAt(0)))) : (p1));
          let _755_v36;
          let _init16 = function (_756_i6) {
            return !(((true) ? (_this.f12) : (true)));
          };
          let _nw125 = Array((new BigNumber(14)).toNumber());
          for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw125.length); _i0_16++) {
            _nw125[_i0_16] = _init16(new BigNumber(_i0_16));
          }
          _755_v36 = _nw125;
          let _index87 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_755_v36).length));
          (_755_v36)[_index87] = (_this.f16).isLessThan(new BigNumber((_dafny.Seq.UnicodeFromString("gdkv")).length));
          let _index88 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_755_v36).length));
          (_755_v36)[_index88] = _this.f12;
          let _757_v37;
          _757_v37 = _module.D3.create_DC6(_755_v36);
          _755_v36 = (_757_v37).dtor_cf14;
          let _758_v38;
          let _nw126 = Array((new BigNumber(9)).toNumber()).fill(_dafny.MultiSet.Empty);
          _758_v38 = _nw126;
          let _index89 = _module.__default.safeIndex(new BigNumber(284), new BigNumber((_758_v38).length));
          (_758_v38)[_index89] = _module.__default.fm13(_696_v1, _this.f12, globalState);
          let _index90 = _module.__default.safeIndex(new BigNumber(284), new BigNumber((_758_v38).length));
          (_758_v38)[_index90] = (p0).Intersect(p0);
          let _index91 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_755_v36).length));
          (_755_v36)[_index91] = _this.f12;
        }
        let _759_v39;
        _759_v39 = _dafny.Map.Empty.slice().updateUnsafe(!(false),_this.f16);
        let _760_v40;
        _760_v40 = _dafny.Seq.of(_759_v39, _759_v39);
        _760_v40 = ((true) ? (_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.of(_759_v39), _module.__default.safeIndex(new BigNumber((p1).length), new BigNumber((_dafny.Seq.of(_759_v39)).length)), _759_v39), _dafny.Seq.of(_759_v39))) : (_760_v40));
        let _761_v41;
        let _nw127 = Array((new BigNumber(26)).toNumber()).fill(false);
        _761_v41 = _nw127;
        _761_v41 = _761_v41;
      } else {
        let _762___mcc_h6 = (_source7).cf7;
        let _763_cf7 = _762___mcc_h6;
        if (!(_module.__default.fm0(_696_v1, globalState))) {
          let _764_v42;
          let _nw128 = Array((new BigNumber(24)).toNumber());
          _nw128[(_dafny.ZERO).toNumber()] = _696_v1;
          _nw128[(_dafny.ONE).toNumber()] = _696_v1;
          _nw128[(new BigNumber(2)).toNumber()] = _696_v1;
          _nw128[(new BigNumber(3)).toNumber()] = _696_v1;
          _nw128[(new BigNumber(4)).toNumber()] = new _dafny.CodePoint('f'.codePointAt(0));
          _nw128[(new BigNumber(5)).toNumber()] = new _dafny.CodePoint('w'.codePointAt(0));
          _nw128[(new BigNumber(6)).toNumber()] = _696_v1;
          _nw128[(new BigNumber(7)).toNumber()] = _696_v1;
          _nw128[(new BigNumber(8)).toNumber()] = _696_v1;
          _nw128[(new BigNumber(9)).toNumber()] = _696_v1;
          _nw128[(new BigNumber(10)).toNumber()] = _696_v1;
          _nw128[(new BigNumber(11)).toNumber()] = new _dafny.CodePoint('c'.codePointAt(0));
          _nw128[(new BigNumber(12)).toNumber()] = _696_v1;
          _nw128[(new BigNumber(13)).toNumber()] = _696_v1;
          _nw128[(new BigNumber(14)).toNumber()] = _696_v1;
          _nw128[(new BigNumber(15)).toNumber()] = new _dafny.CodePoint('r'.codePointAt(0));
          _nw128[(new BigNumber(16)).toNumber()] = _696_v1;
          _nw128[(new BigNumber(17)).toNumber()] = _696_v1;
          _nw128[(new BigNumber(18)).toNumber()] = _696_v1;
          _nw128[(new BigNumber(19)).toNumber()] = _696_v1;
          _nw128[(new BigNumber(20)).toNumber()] = _696_v1;
          _nw128[(new BigNumber(21)).toNumber()] = _696_v1;
          _nw128[(new BigNumber(22)).toNumber()] = _696_v1;
          _nw128[(new BigNumber(23)).toNumber()] = new _dafny.CodePoint('w'.codePointAt(0));
          _764_v42 = _nw128;
          let _765_v43;
          _765_v43 = _dafny.Seq.of(_764_v42);
          _765_v43 = _765_v43;
          let _766_v44;
          _766_v44 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(588),_this.f12);
          _766_v44 = (_766_v44).update(_this.f16, _this.f12);
          let _767_v45;
          _767_v45 = _dafny.Set.fromElements(_this.f12, _this.f12, _this.f12, _this.f12);
          (globalState).f8 = (_this.f16).isEqualTo(new BigNumber((_767_v45).length));
          let _768_v46;
          let _nw129 = Array((new BigNumber(11)).toNumber()).fill(false);
          _768_v46 = _nw129;
          _768_v46 = _768_v46;
          let _769_v47;
          _769_v47 = _dafny.Seq.UnicodeFromString("njgjri");
          let _770_v48;
          let _init17 = function (_771_i7) {
            return (_771_i7).multipliedBy((_dafny.ZERO).minus(_this.f16));
          };
          let _nw130 = Array((new BigNumber(27)).toNumber());
          for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw130.length); _i0_17++) {
            _nw130[_i0_17] = _init17(new BigNumber(_i0_17));
          }
          _770_v48 = _nw130;
          let _772_v49;
          _772_v49 = _module.D4.create_DC10(_770_v48);
          let _773_v50;
          _773_v50 = _dafny.Seq.of(_770_v48);
          let _pat_let_tv15 = _770_v48;
          let _774_v51;
          let _nw131 = Array((new BigNumber(11)).toNumber());
          _nw131[(_dafny.ZERO).toNumber()] = _770_v48;
          _nw131[(_dafny.ONE).toNumber()] = _770_v48;
          _nw131[(new BigNumber(2)).toNumber()] = ((_this.f12) ? (_770_v48) : (_770_v48));
          _nw131[(new BigNumber(3)).toNumber()] = (function (_pat_let19_0) {
            return function (_775_dt__update__tmp_h0) {
              return function (_pat_let20_0) {
                return function (_776_dt__update_hcf24_h0) {
                  return _module.D4.create_DC10(_776_dt__update_hcf24_h0);
                }(_pat_let20_0);
              }(_pat_let_tv15);
            }(_pat_let19_0);
          }(_772_v49)).dtor_cf24;
          _nw131[(new BigNumber(4)).toNumber()] = (_773_v50)[_module.__default.safeIndex(_this.f16, new BigNumber((_773_v50).length))];
          _nw131[(new BigNumber(5)).toNumber()] = _770_v48;
          _nw131[(new BigNumber(6)).toNumber()] = _770_v48;
          _nw131[(new BigNumber(7)).toNumber()] = _770_v48;
          _nw131[(new BigNumber(8)).toNumber()] = _770_v48;
          _nw131[(new BigNumber(9)).toNumber()] = _770_v48;
          _nw131[(new BigNumber(10)).toNumber()] = _770_v48;
          _774_v51 = _nw131;
          let _777_v52;
          _777_v52 = _dafny.Map.Empty.slice().updateUnsafe(false,_696_v1);
          let _778_v53;
          _778_v53 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,_this.f16);
          let _index92 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_768_v46).length));
          (_768_v46)[_index92] = ((_this).fm2(_dafny.MultiSet.fromElements((((_this.f11).contains(!(_this.f12))) ? ((_this.f11).get(!(_this.f12))) : (new BigNumber((_777_v52).length)))), _this.f16, _778_v53, _this.f12, globalState)) && (_this.f12);
          let _index93 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_768_v46).length));
          let _rhs119 = _dafny.Seq.update(_769_v47, _module.__default.safeIndex(_this.f16, new BigNumber((_769_v47).length)), _module.__default.fm14(_697_v2, _this.f16, globalState));
          let _rhs120 = _774_v51;
          let _rhs121 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.update(_769_v47, _module.__default.safeIndex(_this.f16, new BigNumber((_769_v47).length)), new _dafny.CodePoint('c'.codePointAt(0))), _dafny.Seq.Concat(_dafny.Seq.update(_769_v47, _module.__default.safeIndex(_this.f16, new BigNumber((_769_v47).length)), _696_v1), p1));
          let _lhs94 = _768_v46;
          let _lhs95 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_768_v46).length));
          _769_v47 = _rhs119;
          _774_v51 = _rhs120;
          _lhs94[_lhs95] = _rhs121;
        } else {
          let _779_v54;
          _779_v54 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,_this.f12);
          _696_v1 = (((_779_v54).equals(_779_v54)) ? (_696_v1) : (_696_v1));
          let _780_v55;
          let _init18 = function (_781_i8) {
            return _dafny.Seq.UnicodeFromString("nol");
          };
          let _nw132 = Array((_dafny.ONE).toNumber());
          for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw132.length); _i0_18++) {
            _nw132[_i0_18] = _init18(new BigNumber(_i0_18));
          }
          _780_v55 = _nw132;
          let _index94 = _module.__default.safeIndex(new BigNumber(481), new BigNumber((_780_v55).length));
          (_780_v55)[_index94] = p1;
          let _index95 = _module.__default.safeIndex(new BigNumber(481), new BigNumber((_780_v55).length));
          (_780_v55)[_index95] = p1;
          (globalState).f9 = (_this.f16).multipliedBy(_this.f16);
          let _782_v56;
          _782_v56 = _dafny.Seq.of(_this.f16, new BigNumber(278));
          let _783_v57;
          _783_v57 = _module.D5.create_DC12(_782_v56);
          _782_v56 = _dafny.Seq.update((_783_v57).dtor_cf29, _module.__default.safeIndex(((_this.f12) ? (new BigNumber((_dafny.MultiSet.fromElements(_this.f12, _this.f12)).cardinality())) : ((_dafny.ZERO).minus(new BigNumber(((_780_v55)[_module.__default.safeIndex(new BigNumber(481), new BigNumber((_780_v55).length))]).length)))), new BigNumber(((_783_v57).dtor_cf29).length)), _this.f16);
          let _784_v58;
          _784_v58 = _dafny.Seq.of(_this.f12, _this.f12);
          (globalState).f9 = _module.__default.safeDivisionInt(_this.f16, (_this.f16).minus(new BigNumber((_784_v58).length)));
        }
        let _785_v59;
        let _nw133 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _785_v59 = _nw133;
        r0 = _785_v59;
        let _source8 = _695_v0;
        let _786___mcc_h7 = _source8;
        let _787_cf0 = _786___mcc_h7;
        let _788_v60;
        let _nw134 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
        _788_v60 = _nw134;
        let _index96 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_788_v60).length));
        (_788_v60)[_index96] = _this.f16;
        let _789_v61;
        _789_v61 = _dafny.Seq.UnicodeFromString("jljlfkl");
        let _790_v62;
        let _nw135 = Array((new BigNumber(21)).toNumber()).fill(_module.D1.Default());
        _790_v62 = _nw135;
        let _791_v63;
        _791_v63 = _module.D1.create_DC3(_module.__default.fm15(new BigNumber((_789_v61).length), _this.f16, _787_cf0, globalState));
        let _792_v64;
        _792_v64 = _module.D1.create_DC3(_791_v63);
        let _index97 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_790_v62).length));
        (_790_v62)[_index97] = _792_v64;
        let _793_v65;
        _793_v65 = _module.D3.create_DC8(_this.f16, _module.__default.fm1(true, _this.f16, _this.f16, globalState));
        let _794_v66;
        _794_v66 = _dafny.Seq.of(true, _this.f12);
        let _795_v67;
        _795_v67 = _dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC8(_this.f16, new BigNumber((_dafny.Seq.of(_this.f16, new BigNumber((_794_v66).length))).length)),_this.f16);
        let _796_v68;
        _796_v68 = _dafny.Seq.of(_this.f16);
        let _797_v69;
        _797_v69 = _dafny.Set.fromElements(_module.__default.fm9(_this.f16, _this.f12, globalState));
        let _798_v70;
        _798_v70 = _dafny.MultiSet.fromElements(p1, _dafny.Seq.UnicodeFromString("sarchm"), p1);
        let _799_v71;
        let _nw136 = Array((new BigNumber(25)).toNumber());
        _nw136[(_dafny.ZERO).toNumber()] = _696_v1;
        _nw136[(_dafny.ONE).toNumber()] = _696_v1;
        _nw136[(new BigNumber(2)).toNumber()] = _696_v1;
        _nw136[(new BigNumber(3)).toNumber()] = _696_v1;
        _nw136[(new BigNumber(4)).toNumber()] = _696_v1;
        _nw136[(new BigNumber(5)).toNumber()] = _696_v1;
        _nw136[(new BigNumber(6)).toNumber()] = _module.__default.fm14(_697_v2, _this.f16, globalState);
        _nw136[(new BigNumber(7)).toNumber()] = _696_v1;
        _nw136[(new BigNumber(8)).toNumber()] = _696_v1;
        _nw136[(new BigNumber(9)).toNumber()] = _696_v1;
        _nw136[(new BigNumber(10)).toNumber()] = _696_v1;
        _nw136[(new BigNumber(11)).toNumber()] = _696_v1;
        _nw136[(new BigNumber(12)).toNumber()] = _696_v1;
        _nw136[(new BigNumber(13)).toNumber()] = _696_v1;
        _nw136[(new BigNumber(14)).toNumber()] = _696_v1;
        _nw136[(new BigNumber(15)).toNumber()] = _module.__default.fm14(_module.__default.fm15(_this.f16, _this.f16, _this.f12, globalState), new BigNumber(651), globalState);
        _nw136[(new BigNumber(16)).toNumber()] = new _dafny.CodePoint('o'.codePointAt(0));
        _nw136[(new BigNumber(17)).toNumber()] = _696_v1;
        _nw136[(new BigNumber(18)).toNumber()] = _696_v1;
        _nw136[(new BigNumber(19)).toNumber()] = new _dafny.CodePoint('p'.codePointAt(0));
        _nw136[(new BigNumber(20)).toNumber()] = _696_v1;
        _nw136[(new BigNumber(21)).toNumber()] = _696_v1;
        _nw136[(new BigNumber(22)).toNumber()] = _696_v1;
        _nw136[(new BigNumber(23)).toNumber()] = (_789_v61)[_module.__default.safeIndex(_this.f16, new BigNumber((_789_v61).length))];
        _nw136[(new BigNumber(24)).toNumber()] = _696_v1;
        _799_v71 = _nw136;
        let _800_v72;
        _800_v72 = _dafny.Map.Empty.slice().updateUnsafe(_799_v71,false);
        let _801_v73;
        _801_v73 = _dafny.Set.fromElements(false, _787_cf0, !(_787_cf0));
        let _802_v74;
        _802_v74 = _dafny.Set.fromElements(_this.f16);
        let _803_v75;
        _803_v75 = _module.D4.create_DC11(_787_cf0, _this.f16, _module.__default.fm1(_this.f12, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_this.f16)).length)), _this.f16, globalState), new BigNumber((_802_v74).length));
        let _804_v76;
        _804_v76 = _dafny.Map.Empty.slice().updateUnsafe(_789_v61,new BigNumber((_796_v68).length));
        let _805_v77;
        let _nw137 = new _module.C0();
        _nw137.__ctor(_787_cf0, new BigNumber((_804_v76).length), _this.f11, (_794_v66)[_module.__default.safeIndex((((p0).contains(new BigNumber((_this.f11).cardinality()))) ? ((p0).get(new BigNumber((_this.f11).cardinality()))) : (_this.f16)), new BigNumber((_794_v66).length))]);
        _805_v77 = _nw137;
        let _806_v78;
        _806_v78 = _dafny.Set.fromElements(_805_v77, _805_v77, _805_v77, _805_v77);
        let _807_v79;
        let _nw138 = Array((new BigNumber(11)).toNumber());
        _nw138[(_dafny.ZERO).toNumber()] = (_795_v67).contains(_793_v65);
        _nw138[(_dafny.ONE).toNumber()] = (_this.f16).isLessThanOrEqualTo(new BigNumber((_796_v68).length));
        _nw138[(new BigNumber(2)).toNumber()] = (new BigNumber(-92)).isEqualTo(new BigNumber((_797_v69).length));
        _nw138[(new BigNumber(3)).toNumber()] = _this.f12;
        _nw138[(new BigNumber(4)).toNumber()] = (_798_v70).contains(_789_v61);
        _nw138[(new BigNumber(5)).toNumber()] = (((_800_v72).contains(_799_v71)) ? ((_800_v72).get(_799_v71)) : (_this.f12));
        _nw138[(new BigNumber(6)).toNumber()] = _787_cf0;
        _nw138[(new BigNumber(7)).toNumber()] = (_801_v73).IsSubsetOf(_dafny.Set.fromElements((_803_v75).dtor_cf25));
        _nw138[(new BigNumber(8)).toNumber()] = _787_cf0;
        _nw138[(new BigNumber(9)).toNumber()] = _787_cf0;
        _nw138[(new BigNumber(10)).toNumber()] = ((_805_v77).f18).isLessThan(new BigNumber((_806_v78).length));
        _807_v79 = _nw138;
        let _808_v80;
        _808_v80 = _dafny.Map.Empty.slice().updateUnsafe((_805_v77).f17,new BigNumber(53));
        let _index98 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_807_v79).length));
        (_807_v79)[_index98] = (_805_v77).fm2(p0, _this.f16, _808_v80, (_805_v77).f17, globalState);
        let _809_v81;
        _809_v81 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_805_v77).f17);
        let _index99 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_788_v60).length));
        let _index100 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_790_v62).length));
        let _index101 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_807_v79).length));
        let _rhs122 = _this.f16;
        let _rhs123 = _789_v61;
        let _rhs124 = _792_v64;
        let _rhs125 = _807_v79;
        let _rhs126 = !((((_809_v81).contains(p1)) ? ((_809_v81).get(p1)) : (_787_cf0)));
        let _lhs96 = _788_v60;
        let _lhs97 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_788_v60).length));
        let _lhs98 = _790_v62;
        let _lhs99 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_790_v62).length));
        let _lhs100 = _807_v79;
        let _lhs101 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_807_v79).length));
        _lhs96[_lhs97] = _rhs122;
        _789_v61 = _rhs123;
        _lhs98[_lhs99] = _rhs124;
        _807_v79 = _rhs125;
        _lhs100[_lhs101] = _rhs126;
        let _index102 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_788_v60).length));
        (_788_v60)[_index102] = new BigNumber(-18);
        let _index103 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_785_v59).length));
        (_785_v59)[_index103] = p1;
        let _index104 = _module.__default.safeIndex(new BigNumber(736), new BigNumber((_785_v59).length));
        (_785_v59)[_index104] = _module.__default.fm9((_788_v60)[_module.__default.safeIndex(new BigNumber(693), new BigNumber((_788_v60).length))], ((((_808_v80).contains((_805_v77).f17)) ? ((_808_v80).get((_805_v77).f17)) : (_this.f16))).isLessThanOrEqualTo((_dafny.ZERO).minus((_788_v60)[_module.__default.safeIndex(new BigNumber(693), new BigNumber((_788_v60).length))])), globalState);
        _697_v2 = _697_v2;
        let _810_v82;
        let _nw139 = Array((new BigNumber(11)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _810_v82 = _nw139;
        let _index105 = _module.__default.safeIndex(new BigNumber(113), new BigNumber((_810_v82).length));
        (_810_v82)[_index105] = _696_v1;
        let _index106 = _module.__default.safeIndex(new BigNumber(113), new BigNumber((_810_v82).length));
        (_810_v82)[_index106] = ((_this.f12) ? (_696_v1) : (_696_v1));
      }
      let _hi1 = _this.f16;
      for (let _811_i9 = _this.f16; _811_i9.isLessThan(_hi1); _811_i9 = _811_i9.plus(_dafny.ONE)) {
        let _812_v83;
        _812_v83 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(434)), function (_813_i10) {
          return _this.f16;
        })).length),_this.f12);
        let _814_v84;
        _814_v84 = _module.D4.create_DC11(!(false), _811_i9, _811_i9, new BigNumber((_812_v83).length));
        let _815_v85;
        let _nw140 = new _module.C0();
        _nw140.__ctor(((_814_v84).dtor_cf25) === (_this.f12), (_dafny.ZERO).minus(_811_i9), _this.f11, false);
        _815_v85 = _nw140;
        let _816_v86;
        let _nw141 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Map.Empty);
        _816_v86 = _nw141;
        let _817_v87;
        _817_v87 = _dafny.Seq.of(_this.f16, (_815_v85).f18);
        let _818_v88;
        _818_v88 = _dafny.Seq.of(new BigNumber((_817_v87).length));
        let _819_v89;
        _819_v89 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,_818_v88);
        let _index107 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_816_v86).length));
        (_816_v86)[_index107] = (_819_v89).Merge(_819_v89);
        let _820_v90;
        let _nw142 = Array((new BigNumber(2)).toNumber()).fill(false);
        _820_v90 = _nw142;
        let _821_v91;
        _821_v91 = _dafny.Seq.of(_820_v90, _820_v90);
        let _index108 = _module.__default.safeIndex(new BigNumber(225), new BigNumber((_816_v86).length));
        (_816_v86)[_index108] = ((_dafny.areEqual(_821_v91, _821_v91)) ? (((_this.f12) ? (_819_v89) : (_819_v89))) : ((_819_v89).Merge(_819_v89)));
        (_this).f12 = ((_815_v85).f17) && (((_815_v85).f17) && (_this.f12));
        let _822_v92;
        _822_v92 = _dafny.Map.Empty.slice().updateUnsafe((_815_v85).f18,_dafny.Seq.Create(_module.__default.abs(new BigNumber(66)), ((_823_v85) => function (_824_i11) {
          return (_823_v85).f18;
        })(_815_v85)));
        let _825_v93;
        _825_v93 = _dafny.Set.fromElements(_817_v87, _817_v87, (((_822_v92).contains(_811_i9)) ? ((_822_v92).get(_811_i9)) : (_818_v88)));
        let _826_v94;
        _826_v94 = _dafny.Map.Empty.slice().updateUnsafe(_825_v93,_this.f12);
        _826_v94 = (_826_v94).update(_825_v93, (_815_v85).f17);
      }
      let _827_v95;
      let _nw143 = Array((new BigNumber(20)).toNumber()).fill(false);
      _827_v95 = _nw143;
      let _index109 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length));
      (_827_v95)[_index109] = _this.f12;
      let _index110 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length));
      (_827_v95)[_index110] = (_this.f11).IsProperSubsetOf((_this.f11).Union(_dafny.MultiSet.fromElements(_this.f12, true)));
      let _828_v96;
      let _nw144 = Array((new BigNumber(16)).toNumber());
      _nw144[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('b'.codePointAt(0));
      _nw144[(_dafny.ONE).toNumber()] = _696_v1;
      _nw144[(new BigNumber(2)).toNumber()] = _696_v1;
      _nw144[(new BigNumber(3)).toNumber()] = _696_v1;
      _nw144[(new BigNumber(4)).toNumber()] = _696_v1;
      _nw144[(new BigNumber(5)).toNumber()] = _696_v1;
      _nw144[(new BigNumber(6)).toNumber()] = _696_v1;
      _nw144[(new BigNumber(7)).toNumber()] = _module.__default.fm14(_697_v2, (_dafny.ZERO).minus(_this.f16), globalState);
      _nw144[(new BigNumber(8)).toNumber()] = _696_v1;
      _nw144[(new BigNumber(9)).toNumber()] = _696_v1;
      _nw144[(new BigNumber(10)).toNumber()] = _696_v1;
      _nw144[(new BigNumber(11)).toNumber()] = _696_v1;
      _nw144[(new BigNumber(12)).toNumber()] = _696_v1;
      _nw144[(new BigNumber(13)).toNumber()] = _696_v1;
      _nw144[(new BigNumber(14)).toNumber()] = _696_v1;
      _nw144[(new BigNumber(15)).toNumber()] = _696_v1;
      _828_v96 = _nw144;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_828_v96).length))) {
        let _829_i12 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_829_i12)) && ((_829_i12).isLessThan(new BigNumber((_828_v96).length))))) {
          (_828_v96)[(_829_i12)] = _696_v1;
        }
      }
      let _830_v97;
      let _init19 = ((_831_p1) => function (_832_i13) {
        return _831_p1;
      })(p1);
      let _nw145 = Array((new BigNumber(3)).toNumber());
      for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw145.length); _i0_19++) {
        _nw145[_i0_19] = _init19(new BigNumber(_i0_19));
      }
      _830_v97 = _nw145;
      let _833_v98;
      _833_v98 = _module.D5.create_DC13(_this.f16, (_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))], _this.f16, _830_v97, (_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))]);
      let _834_v99;
      _834_v99 = _dafny.Seq.of((_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))], (_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))], _this.f12, false);
      let _835_v100;
      let _nw146 = Array((new BigNumber(3)).toNumber());
      _nw146[(_dafny.ZERO).toNumber()] = p0;
      _nw146[(_dafny.ONE).toNumber()] = p0;
      _nw146[(new BigNumber(2)).toNumber()] = ((p0).update((_833_v98).dtor_cf32, _module.__default.abs((_dafny.ZERO).minus(_this.f16)))).update(new BigNumber((_dafny.Seq.update(_834_v99, _module.__default.safeIndex(new BigNumber(142), new BigNumber((_834_v99).length)), (_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))])).length), _module.__default.abs(_this.f16));
      _835_v100 = _nw146;
      let _index111 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_835_v100).length));
      (_835_v100)[_index111] = (p0).Difference(p0);
      let _index112 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_835_v100).length));
      (_835_v100)[_index112] = (p0).Intersect(p0);
      if ((_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))]) {
        (_this).f12 = (_module.__default.safeDivisionInt(new BigNumber((_dafny.MultiSet.fromElements((_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))])).cardinality()), _this.f16)).isLessThanOrEqualTo(_this.f16);
        (globalState).f9 = _this.f16;
        (globalState).f7 = _module.__default.fm0(_696_v1, globalState);
        let _836_v101;
        _836_v101 = _module.D3.create_DC6(_827_v95);
        let _source9 = _836_v101;
        if (_source9.is_DC7) {
          let _837___mcc_h8 = (_source9).cf15;
          let _838___mcc_h9 = (_source9).cf16;
          let _839___mcc_h10 = (_source9).cf17;
          let _840_cf17 = _839___mcc_h10;
          let _841_cf16 = _838___mcc_h9;
          let _842_cf15 = _837___mcc_h8;
          (globalState).f7 = true;
          let _rhs127 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("eokqyem")).length), _841_cf16));
          let _rhs128 = _module.__default.safeModuloInt((_841_cf16).plus(_842_cf15), _this.f16);
          let _rhs129 = _840_cf17;
          let _rhs130 = (_dafny.ZERO).minus((_840_cf17).f18);
          let _lhs102 = globalState;
          let _lhs103 = globalState;
          _842_cf15 = _rhs127;
          _lhs102.f9 = _rhs128;
          _840_cf17 = _rhs129;
          _lhs103.f9 = _rhs130;
          let _843_v102;
          _843_v102 = _dafny.Map.Empty.slice().updateUnsafe(_842_cf15,_this.f16);
          let _844_v103;
          _844_v103 = _dafny.Set.fromElements(_842_cf15, (_840_cf17).f18, _841_cf16, (new BigNumber(228)).multipliedBy(new BigNumber((_843_v102).length)), (_840_cf17).fm12(globalState));
          _844_v103 = _844_v103;
          let _845_v104;
          _845_v104 = _dafny.Set.fromElements((_840_cf17).f17);
          let _846_v105;
          _846_v105 = _dafny.Map.Empty.slice().updateUnsafe(true,_this.f12);
          let _847_v106;
          _847_v106 = _dafny.Set.fromElements(_this.f12, (_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))], !(_this.f12), !(_module.__default.fm16(new BigNumber((_845_v104).length), globalState)).equals(_846_v105));
          let _848_v107;
          _848_v107 = _dafny.Seq.of(_842_cf15, new BigNumber(((_835_v100)[_module.__default.safeIndex(new BigNumber(505), new BigNumber((_835_v100).length))]).cardinality()));
          let _index113 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length));
          let _rhs131 = ((_844_v103).Difference(_844_v103)).IsProperSubsetOf(function () {
            let _coll20 = new _dafny.Set();
            for (const _compr_20 of (_dafny.Seq.update(_848_v107, _module.__default.safeIndex(_this.f16, new BigNumber((_848_v107).length)), _842_cf15)).Elements) {
              let _849_v108 = _compr_20;
              if (_dafny.Seq.contains(_dafny.Seq.update(_848_v107, _module.__default.safeIndex(_this.f16, new BigNumber((_848_v107).length)), _842_cf15), _849_v108)) {
                _coll20.add((_849_v108).minus(new BigNumber(992)));
              }
            }
            return _coll20;
          }());
          let _rhs132 = new BigNumber((_843_v102).length);
          let _rhs133 = (_847_v106).Intersect(_dafny.Set.fromElements((_840_cf17).f17));
          let _rhs134 = (_845_v104).IsSubsetOf(_dafny.Set.fromElements(_this.f12));
          let _lhs104 = _827_v95;
          let _lhs105 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length));
          let _lhs106 = globalState;
          _lhs104[_lhs105] = _rhs131;
          _841_cf16 = _rhs132;
          _847_v106 = _rhs133;
          _lhs106.f8 = _rhs134;
        } else if (_source9.is_DC8) {
          let _850___mcc_h11 = (_source9).cf18;
          let _851___mcc_h12 = (_source9).cf19;
          let _852_cf19 = _851___mcc_h12;
          let _853_cf18 = _850___mcc_h11;
          let _index114 = _module.__default.safeIndex(new BigNumber(474), new BigNumber((_830_v97).length));
          (_830_v97)[_index114] = p1;
          let _pat_let_tv16 = _827_v95;
          let _index115 = _module.__default.safeIndex(new BigNumber(474), new BigNumber((_830_v97).length));
          let _rhs135 = p1;
          let _rhs136 = _dafny.Seq.Concat(((!(!(false))) ? (_dafny.Seq.of((_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))], _this.f12)) : (_834_v99)), _834_v99);
          let _rhs137 = !(_module.__default.fm0(_696_v1, globalState));
          let _rhs138 = _this.f12;
          let _rhs139 = (function (_pat_let21_0) {
            return function (_854_dt__update__tmp_h1) {
              return function (_pat_let22_0) {
                return function (_855_dt__update_hcf14_h0) {
                  return _module.D3.create_DC6(_855_dt__update_hcf14_h0);
                }(_pat_let22_0);
              }(_pat_let_tv16);
            }(_pat_let21_0);
          }(_836_v101)).dtor_cf14;
          let _lhs107 = _830_v97;
          let _lhs108 = _module.__default.safeIndex(new BigNumber(474), new BigNumber((_830_v97).length));
          let _lhs109 = globalState;
          let _lhs110 = globalState;
          let _lhs111 = globalState;
          _lhs107[_lhs108] = _rhs135;
          _lhs109.f6 = _rhs136;
          _lhs110.f7 = _rhs137;
          _lhs111.f3 = _rhs138;
          _827_v95 = _rhs139;
          let _856_v109;
          let _nw147 = new _module.C0();
          _nw147.__ctor(false, _this.f16, _this.f11, (_834_v99)[_module.__default.safeIndex(_852_cf19, new BigNumber((_834_v99).length))]);
          _856_v109 = _nw147;
          let _857_v110;
          _857_v110 = _dafny.Map.Empty.slice().updateUnsafe(_852_cf19,_856_v109);
          let _858_v111;
          _858_v111 = _dafny.Map.Empty.slice().updateUnsafe((_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))],_this.f16);
          (globalState).f7 = (_this).fm2(_module.__default.fm13(_696_v1, (_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))], globalState), _module.__default.safeModuloInt(new BigNumber((_857_v110).length), _853_cf18), (_858_v111).Merge(_858_v111), (_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))], globalState);
          let _859_v112;
          _859_v112 = _dafny.Seq.of(_853_cf18, _853_cf18);
          let _860_v113;
          let _nw148 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
          _860_v113 = _nw148;
          let _861_v114;
          _861_v114 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_dafny.Seq.update(((_this.f12) ? (_859_v112) : (_859_v112)), _module.__default.safeIndex(_this.f16, new BigNumber((((_this.f12) ? (_859_v112) : (_859_v112))).length)), _this.f16)),_860_v113);
          _861_v114 = (_861_v114).update((p0).Difference((_835_v100)[_module.__default.safeIndex(new BigNumber(505), new BigNumber((_835_v100).length))]), _860_v113);
          let _862_v115;
          let _nw149 = new _module.C0();
          _nw149.__ctor((_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))], _this.f16, (_856_v109.f11).Union(_this.f11), ((_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))]) || ((_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))]));
          _862_v115 = _nw149;
          _862_v115 = _862_v115;
        } else if (_source9.is_DC9) {
          let _863___mcc_h13 = (_source9).cf20;
          let _864___mcc_h14 = (_source9).cf21;
          let _865___mcc_h15 = (_source9).cf22;
          let _866___mcc_h16 = (_source9).cf23;
          let _867_cf23 = _866___mcc_h16;
          let _868_cf22 = _865___mcc_h15;
          let _869_cf21 = _864___mcc_h14;
          let _870_cf20 = _863___mcc_h13;
          let _871_v116;
          let _nw150 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
          _871_v116 = _nw150;
          let _index116 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_871_v116).length));
          (_871_v116)[_index116] = _this.f16;
          let _872_v117;
          _872_v117 = _dafny.Map.Empty.slice().updateUnsafe(p1,_this.f12);
          let _index117 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_871_v116).length));
          (_871_v116)[_index117] = _module.__default.safeModuloInt(new BigNumber((_872_v117).length), (_868_cf22).f18);
          let _873_v118;
          _873_v118 = _module.D1.create_DC2(_this.f12, (_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))], _this.f12, _696_v1, _696_v1);
          let _874_v119;
          _874_v119 = _module.D1.create_DC3(_873_v118);
          let _875_v120;
          _875_v120 = _module.D1.create_DC3(_874_v119);
          let _876_v121;
          _876_v121 = _module.D1.create_DC3(_873_v118);
          let _pat_let_tv17 = _873_v118;
          let _877_v122;
          _877_v122 = _dafny.Set.fromElements(function (_pat_let23_0) {
            return function (_878_dt__update__tmp_h2) {
              return function (_pat_let24_0) {
                return function (_879_dt__update_hcf7_h0) {
                  return _module.D1.create_DC3(_879_dt__update_hcf7_h0);
                }(_pat_let24_0);
              }(_pat_let_tv17);
            }(_pat_let23_0);
          }(_module.D1.create_DC3(_875_v120)), _876_v121);
          let _880_v123;
          _880_v123 = _dafny.MultiSet.fromElements((_877_v122).Union(_877_v122), _877_v122);
          _880_v123 = (_880_v123).Difference((_880_v123).Difference(_module.__default.fm17((_868_cf22).f17, !(false), (_868_cf22).f18, new BigNumber(190), globalState)));
          _870_cf20 = (_dafny.ZERO).minus(_870_cf20);
          let _881_v124;
          _881_v124 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,_827_v95);
          _881_v124 = (_881_v124).update((_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))], _827_v95);
        } else {
          let _882___mcc_h17 = (_source9).cf14;
          let _883_cf14 = _882___mcc_h17;
          let _884_v125;
          _884_v125 = _dafny.Set.fromElements(_this.f16);
          let _885_v126;
          _885_v126 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,_this.f16);
          let _886_v127;
          _886_v127 = _dafny.Set.fromElements(_this.f12, (_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))]);
          let _887_v128;
          _887_v128 = _module.D2.create_DC4(_this.f16);
          let _888_v129;
          let _nw151 = Array((new BigNumber(6)).toNumber());
          _nw151[(_dafny.ZERO).toNumber()] = _this.f16;
          _nw151[(_dafny.ONE).toNumber()] = _this.f16;
          _nw151[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_module.D2.create_DC5(new BigNumber((_834_v99).length), _this.f16, _this.f16, _this.f16, _834_v99)).dtor_cf9,_this.f12)).length));
          _nw151[(new BigNumber(3)).toNumber()] = new BigNumber(((_884_v125).Difference(_dafny.Set.fromElements((_dafny.ZERO).minus(_this.f16)))).length);
          _nw151[(new BigNumber(4)).toNumber()] = (new BigNumber(-623)).minus(new BigNumber(((_885_v126).update(!(_this.f12), _module.__default.fm1(!((_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))]), new BigNumber((_886_v127).length), _this.f16, globalState))).length));
          _nw151[(new BigNumber(5)).toNumber()] = _module.__default.fm1(_this.f12, (_887_v128).dtor_cf8, new BigNumber(797), globalState);
          _888_v129 = _nw151;
          let _index118 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((_888_v129).length));
          (_888_v129)[_index118] = _this.f16;
          let _889_v130;
          _889_v130 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),(_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))]);
          let _890_v131;
          _890_v131 = _dafny.MultiSet.fromElements(_696_v1, (_697_v2).dtor_cf5, _696_v1, (p1)[_module.__default.safeIndex(_this.f16, new BigNumber((p1).length))]);
          let _891_v132;
          _891_v132 = _dafny.Seq.of(_this.f16, new BigNumber(((_889_v130).Merge(_889_v130)).length), _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.update(_834_v99, _module.__default.safeIndex(new BigNumber(671), new BigNumber((_834_v99).length)), (_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))])).length), _this.f16), _this.f16, new BigNumber(((_890_v131).Difference(_890_v131)).cardinality()));
          let _index119 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((_888_v129).length));
          (_888_v129)[_index119] = (_891_v132)[_module.__default.safeIndex(_this.f16, new BigNumber((_891_v132).length))];
          let _892_v133;
          let _nw152 = new _module.C0();
          _nw152.__ctor(_this.f12, _this.f16, (_dafny.MultiSet.fromElements((_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))], _this.f12)).Union(_this.f11), (_this.f16).isEqualTo(_this.f16));
          _892_v133 = _nw152;
          let _index120 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((_888_v129).length));
          (_888_v129)[_index120] = _this.f16;
          let _index121 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((_888_v129).length));
          (_888_v129)[_index121] = (_888_v129)[_module.__default.safeIndex(new BigNumber(273), new BigNumber((_888_v129).length))];
        }
        let _893_v135;
        _893_v135 = _dafny.Set.fromElements(_this.f16);
        let _894_v136;
        let _nw153 = new _module.C0();
        _nw153.__ctor(false, _this.f16, _dafny.MultiSet.FromArray(_dafny.Seq.update(_dafny.Seq.of((_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))]), _module.__default.safeIndex(_this.f16, new BigNumber((_dafny.Seq.of((_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))])).length)), _this.f12)), (_893_v135).IsProperSubsetOf(function () {
          let _coll21 = new _dafny.Set();
          for (const _compr_21 of _dafny.IntegerRange(new BigNumber(447), new BigNumber(623))) {
            let _895_v134 = _compr_21;
            if (((new BigNumber(447)).isLessThanOrEqualTo(_895_v134)) && ((_895_v134).isLessThan(new BigNumber(623)))) {
              _coll21.add(_module.__default.safeModuloInt(_895_v134, _this.f16));
            }
          }
          return _coll21;
        }()));
        _894_v136 = _nw153;
      } else {
        (globalState).f9 = _this.f16;
        let _896_v137;
        _896_v137 = _dafny.Map.Empty.slice().updateUnsafe(!((_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))]),_this.f16);
        (globalState).f3 = !(_module.__default.safeModuloInt(new BigNumber(-201), (((_896_v137).contains((_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))])) ? ((_896_v137).get((_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))])) : (_this.f16)))).isEqualTo(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(606)), function (_897_i14) {
          return (_dafny.ZERO).minus(_this.f16);
        })).length));
        if (!(_dafny.Map.Empty.slice().updateUnsafe(_this.f12,_this.f16)).contains((_this.f16).isEqualTo(_module.__default.fm1(_this.f12, _this.f16, _this.f16, globalState)))) {
          let _898_v138;
          _898_v138 = _dafny.Set.fromElements(_this.f12, _this.f12);
          let _899_v139;
          let _nw154 = new _module.C0();
          _nw154.__ctor(_this.f12, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f16,new BigNumber((_898_v138).length))).length), (_this.f11).update(_this.f12, _module.__default.abs(new BigNumber((p1).length))), _this.f12);
          _899_v139 = _nw154;
          let _900_v140;
          _900_v140 = _module.D3.create_DC9(_this.f16, new BigNumber((_module.__default.fm18(_this.f16, _this.f12, globalState)).length), _899_v139, (_899_v139).f18);
          let _901_v141;
          _901_v141 = _dafny.Map.Empty.slice().updateUnsafe(_this.f16,_module.__default.safeDivisionInt(_this.f16, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_900_v140,(_899_v139).f17)).length)));
          _901_v141 = ((_901_v141).Merge(_901_v141)).Merge(_901_v141);
          let _index122 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length));
          (_827_v95)[_index122] = true;
          let _902_v142;
          let _init20 = function (_903_i15) {
            return (_903_i15).multipliedBy(new BigNumber(263));
          };
          let _nw155 = Array((new BigNumber(10)).toNumber());
          for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw155.length); _i0_20++) {
            _nw155[_i0_20] = _init20(new BigNumber(_i0_20));
          }
          _902_v142 = _nw155;
          let _index123 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_902_v142).length));
          (_902_v142)[_index123] = _this.f16;
          let _904_v143;
          _904_v143 = _module.D3.create_DC6(_827_v95);
          let _index124 = _module.__default.safeIndex(new BigNumber(712), new BigNumber((_902_v142).length));
          (_902_v142)[_index124] = _this.f16;
          let _index125 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_902_v142).length));
          let _index126 = _module.__default.safeIndex(new BigNumber(712), new BigNumber((_902_v142).length));
          let _rhs140 = _696_v1;
          let _rhs141 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_this.f16, (_this.f16).plus(new BigNumber((_this.f11).cardinality()))));
          let _rhs142 = _904_v143;
          let _rhs143 = _this.f16;
          let _lhs112 = _902_v142;
          let _lhs113 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_902_v142).length));
          let _lhs114 = _902_v142;
          let _lhs115 = _module.__default.safeIndex(new BigNumber(712), new BigNumber((_902_v142).length));
          _696_v1 = _rhs140;
          _lhs112[_lhs113] = _rhs141;
          _904_v143 = _rhs142;
          _lhs114[_lhs115] = _rhs143;
          (globalState).f3 = (_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))];
          let _index127 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length));
          let _index128 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_835_v100).length));
          let _rhs144 = _this.f12;
          let _rhs145 = p0;
          let _rhs146 = (_899_v139).fm2(_module.__default.fm13(_696_v1, false, globalState), (_dafny.ZERO).minus(_this.f16), _896_v137, !(_this.f12) || ((_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))]), globalState);
          let _rhs147 = (_899_v139).f17;
          let _lhs116 = _827_v95;
          let _lhs117 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length));
          let _lhs118 = _835_v100;
          let _lhs119 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_835_v100).length));
          let _lhs120 = globalState;
          let _lhs121 = _this;
          _lhs116[_lhs117] = _rhs144;
          _lhs118[_lhs119] = _rhs145;
          _lhs120.f7 = _rhs146;
          _lhs121.f12 = _rhs147;
        } else {
          let _905_v144;
          let _nw156 = Array((new BigNumber(24)).toNumber());
          _905_v144 = _nw156;
          let _906_v145;
          _906_v145 = _module.D3.create_DC6(_827_v95);
          let _pat_let_tv18 = _827_v95;
          let _pat_let_tv19 = _827_v95;
          let _index129 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_905_v144).length));
          (_905_v144)[_index129] = function (_pat_let25_0) {
            return function (_909_dt__update__tmp_h4) {
              return function (_pat_let28_0) {
                return function (_910_dt__update_hcf14_h2) {
                  return _module.D3.create_DC6(_910_dt__update_hcf14_h2);
                }(_pat_let28_0);
              }(_pat_let_tv19);
            }(_pat_let25_0);
          }(function (_pat_let26_0) {
            return function (_907_dt__update__tmp_h3) {
              return function (_pat_let27_0) {
                return function (_908_dt__update_hcf14_h1) {
                  return _module.D3.create_DC6(_908_dt__update_hcf14_h1);
                }(_pat_let27_0);
              }(_pat_let_tv18);
            }(_pat_let26_0);
          }(_906_v145));
          let _index130 = _module.__default.safeIndex(new BigNumber(190), new BigNumber((_905_v144).length));
          (_905_v144)[_index130] = _906_v145;
          let _911_v146;
          let _nw157 = new _module.C0();
          _nw157.__ctor(_this.f12, new BigNumber(((_835_v100)[_module.__default.safeIndex(new BigNumber(505), new BigNumber((_835_v100).length))]).cardinality()), (_this.f11).Intersect(_this.f11), !(!((_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))]) || ((_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))])));
          _911_v146 = _nw157;
          let _912_v147;
          let _nw158 = new _module.C0();
          _nw158.__ctor((((_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))]) ? ((_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))]) : (_this.f12)), ((_911_v146).f18).multipliedBy(_this.f16), (_this.f11).update(_this.f12, _module.__default.abs((_911_v146).f18)), _this.f12);
          _912_v147 = _nw158;
          _912_v147 = _912_v147;
          let _913_v148;
          let _nw159 = new _module.C0();
          _nw159.__ctor((_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))], (_911_v146).f18, (_912_v147.f11).Union(_912_v147.f11), _this.f12);
          _913_v148 = _nw159;
          let _914_v149;
          _914_v149 = _dafny.Seq.UnicodeFromString("ggsc");
          _914_v149 = p1;
        }
        let _915_v150;
        _915_v150 = _dafny.Seq.of(_827_v95);
        let _916_v151;
        _916_v151 = _dafny.MultiSet.fromElements(_696_v1);
        let _917_v152;
        _917_v152 = _dafny.Map.Empty.slice().updateUnsafe(false,_916_v151);
        let _918_v153;
        let _nw160 = Array((new BigNumber(10)).toNumber());
        _nw160[(_dafny.ZERO).toNumber()] = _827_v95;
        _nw160[(_dafny.ONE).toNumber()] = _827_v95;
        _nw160[(new BigNumber(2)).toNumber()] = _827_v95;
        _nw160[(new BigNumber(3)).toNumber()] = _827_v95;
        _nw160[(new BigNumber(4)).toNumber()] = (_915_v150)[_module.__default.safeIndex(new BigNumber((_917_v152).length), new BigNumber((_915_v150).length))];
        _nw160[(new BigNumber(5)).toNumber()] = _827_v95;
        _nw160[(new BigNumber(6)).toNumber()] = _827_v95;
        _nw160[(new BigNumber(7)).toNumber()] = (_915_v150)[_module.__default.safeIndex(_this.f16, new BigNumber((_915_v150).length))];
        _nw160[(new BigNumber(8)).toNumber()] = _827_v95;
        _nw160[(new BigNumber(9)).toNumber()] = _827_v95;
        _918_v153 = _nw160;
        let _index131 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_918_v153).length));
        (_918_v153)[_index131] = _827_v95;
        let _index132 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_918_v153).length));
        (_918_v153)[_index132] = _827_v95;
        (globalState).f3 = (_827_v95)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_827_v95).length))];
      }
      r0 = _830_v97;
      return r0;
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f11 = _dafny.MultiSet.Empty;
      this._f12 = false;
      this.f15 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
    set f11(value) {
      let _this = this;
      _this._f11 = value;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
    set f12(value) {
      let _this = this;
      _this._f12 = value;
    };
    __ctor(f15, f11, f12) {
      let _this = this;
      (_this).f15 = f15;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      return;
    }
    fm2(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _this.f12;
    };
    fm8(p0, p1, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(622)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(392)), function (_919_i0) {
        return new BigNumber(959);
      }));
    };
    m4(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Seq.UnicodeFromString("");
      let r2 = _dafny.ZERO;
      let r3 = false;
      if ((new BigNumber((_dafny.Seq.of(p0)).length)).isEqualTo(p0)) {
        if (p1) {
          let _920_v0;
          let _nw161 = Array((new BigNumber(16)).toNumber()).fill(false);
          _920_v0 = _nw161;
          _920_v0 = (((_this.f12) && (p1)) ? (_920_v0) : (_920_v0));
          r2 = _module.__default.safeDivisionInt(p0, p0);
          let _921_v1;
          _921_v1 = _dafny.Seq.UnicodeFromString("nbffrd");
          (globalState).f7 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(5)), function (_922_i0) {
            return new _dafny.CodePoint('j'.codePointAt(0));
          }), _921_v1), _dafny.Seq.UnicodeFromString("xycvusgt"));
          let _923_v2;
          let _nw162 = Array((new BigNumber(17)).toNumber()).fill([]);
          _923_v2 = _nw162;
          let _924_v3;
          _924_v3 = _dafny.Map.Empty.slice().updateUnsafe(_923_v2,_this.f15);
          _924_v3 = (_924_v3).update(_923_v2, _this.f15);
          let _925_v4;
          _925_v4 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),_this.f12);
          let _926_v5;
          _926_v5 = _dafny.Set.fromElements(_920_v0);
          let _927_v6;
          _927_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_925_v4).length),(_926_v5).IsProperSubsetOf(_926_v5));
          let _rhs148 = _dafny.Seq.Concat(_921_v1, _921_v1);
          let _rhs149 = _927_v6;
          let _rhs150 = _920_v0;
          let _rhs151 = _this.f12;
          let _rhs152 = !(_this.f15) || (_this.f12);
          let _lhs122 = globalState;
          let _lhs123 = globalState;
          _921_v1 = _rhs148;
          _927_v6 = _rhs149;
          _920_v0 = _rhs150;
          _lhs122.f7 = _rhs151;
          _lhs123.f3 = _rhs152;
        } else {
          let _928_v7;
          _928_v7 = _dafny.Seq.of(p0, p0);
          _928_v7 = _928_v7;
          let _929_v8;
          let _nw163 = Array((new BigNumber(28)).toNumber()).fill(false);
          _929_v8 = _nw163;
          let _index133 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_929_v8).length));
          (_929_v8)[_index133] = !_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(23)), function (_930_i1) {
            return new _dafny.CodePoint('u'.codePointAt(0));
          }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(924)), function (_931_i2) {
            return new _dafny.CodePoint('h'.codePointAt(0));
          }));
          let _index134 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_929_v8).length));
          (_929_v8)[_index134] = _this.f12;
          let _932_v9;
          let _nw164 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
          _932_v9 = _nw164;
          let _index135 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_932_v9).length));
          (_932_v9)[_index135] = _module.__default.fm1(_this.f15, p0, p0, globalState);
          let _933_v10;
          _933_v10 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f12);
          let _index136 = _module.__default.safeIndex(new BigNumber(459), new BigNumber((_932_v9).length));
          (_932_v9)[_index136] = (_module.__default.safeModuloInt(p0, p0)).minus(new BigNumber((_933_v10).length));
          _module.__default.m0(globalState);
          let _934_v11;
          _934_v11 = _dafny.Map.Empty.slice().updateUnsafe(_929_v8,_932_v9);
          _934_v11 = (_934_v11).update(_929_v8, _932_v9);
        }
        (globalState).f9 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("xswkq")).length));
        (globalState).f9 = p0;
        r2 = (((_this.f11).contains(!(_this.f12) || (_this.f15))) ? ((_this.f11).get(!(_this.f12) || (_this.f15))) : (p0));
        let _935_v12;
        _935_v12 = _dafny.Set.fromElements(new BigNumber(-479), p0, p0);
        let _936_v13;
        _936_v13 = _dafny.Map.Empty.slice().updateUnsafe(!(p1) || (_this.f12),_dafny.Set.fromElements(new BigNumber((_935_v12).length)));
        _935_v12 = (((_936_v13).contains(false)) ? ((_936_v13).get(false)) : (_935_v12));
      } else {
        let _937_v14;
        _937_v14 = _dafny.Seq.of(p0);
        let _938_v16;
        let _nw165 = new _module.C0();
        _nw165.__ctor(_this.f12, p0, _this.f11, !(_dafny.Set.fromElements(new BigNumber((_937_v14).length))).equals(function () {
          let _coll22 = new _dafny.Set();
          for (const _compr_22 of _dafny.IntegerRange(new BigNumber(279), new BigNumber(299))) {
            let _939_v15 = _compr_22;
            if (((new BigNumber(279)).isLessThanOrEqualTo(_939_v15)) && ((_939_v15).isLessThan(new BigNumber(299)))) {
              _coll22.add(_module.__default.safeModuloInt(_939_v15, p0));
            }
          }
          return _coll22;
        }()));
        _938_v16 = _nw165;
        let _940_v17;
        _940_v17 = _dafny.Map.Empty.slice().updateUnsafe((_938_v16).f17,p0);
        _940_v17 = (_940_v17).update(p1, p0);
        let _941_v18;
        _941_v18 = _dafny.Seq.of(_this.f15);
        let _942_v19;
        let _nw166 = Array((new BigNumber(2)).toNumber());
        _nw166[(_dafny.ZERO).toNumber()] = (_938_v16).f17;
        _nw166[(_dafny.ONE).toNumber()] = !(((_938_v16).f18).isLessThan(new BigNumber((_941_v18).length)));
        _942_v19 = _nw166;
        let _943_v20;
        _943_v20 = _dafny.Map.Empty.slice().updateUnsafe((_938_v16).f18,_this.f12);
        let _944_v21;
        _944_v21 = _dafny.Map.Empty.slice().updateUnsafe(_this.f15,(((_943_v20).contains(new BigNumber(694))) ? ((_943_v20).get(new BigNumber(694))) : (p1)));
        let _index137 = _module.__default.safeIndex(new BigNumber(21), new BigNumber((_942_v19).length));
        (_942_v19)[_index137] = (((_944_v21).contains(!(!(true)))) ? ((_944_v21).get(!(!(true)))) : ((_938_v16).f17));
        let _index138 = _module.__default.safeIndex(new BigNumber(21), new BigNumber((_942_v19).length));
        (_942_v19)[_index138] = _this.f15;
        (globalState).f3 = (p0).isEqualTo(_module.__default.safeDivisionInt(p0, p0));
        _940_v17 = (_940_v17).update((_942_v19)[_module.__default.safeIndex(new BigNumber(21), new BigNumber((_942_v19).length))], (_938_v16).f18);
      }
      let _945_i3;
      _945_i3 = _dafny.ZERO;
      L7: {
        while (p1) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_945_i3)) {
              break L7;
            }
            _945_i3 = (_945_i3).plus(_dafny.ONE);
            let _946_v22;
            _946_v22 = _dafny.Seq.of(p0);
            let _947_v23;
            _947_v23 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),true);
            _946_v22 = (_this).fm8((((_947_v23).contains(p0)) ? ((_947_v23).get(p0)) : (_this.f12)), p1, globalState);
            let _948_v24;
            _948_v24 = _dafny.Set.fromElements(_this.f12);
            let _949_v25;
            _949_v25 = new _dafny.CodePoint('q'.codePointAt(0));
            let _950_v26;
            _950_v26 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_948_v24).length)).minus(p0),_949_v25);
            _950_v26 = _950_v26;
            let _951_v27;
            _951_v27 = _dafny.MultiSet.fromElements(new BigNumber(100), p0, p0, new BigNumber(494), p0);
            (globalState).f9 = (((!(_this.f15)) ? (new BigNumber((_this.f11).cardinality())) : ((((_951_v27).contains(p0)) ? ((_951_v27).get(p0)) : (new BigNumber(316)))))).multipliedBy(p0);
            let _952_v28;
            _952_v28 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_946_v22, _module.__default.safeIndex(new BigNumber(-740), new BigNumber((_946_v22).length)), p0),new BigNumber(954));
            let _953_v30;
            _953_v30 = _dafny.MultiSet.fromElements(_946_v22);
            let _954_v31;
            _954_v31 = _dafny.Seq.of(_952_v28);
            _952_v28 = ((_952_v28).Merge(function () {
              let _coll23 = new _dafny.Map();
              for (const _compr_23 of (_953_v30).Elements) {
                let _955_v29 = _compr_23;
                if ((_953_v30).contains(_955_v29)) {
                  _coll23.push([_955_v29,p0]);
                }
              }
              return _coll23;
            }())).Merge((_954_v31)[_module.__default.safeIndex(p0, new BigNumber((_954_v31).length))]);
          }
        }
      }
      (_this).f15 = _this.f15;
      let _956_v32;
      _956_v32 = _dafny.Seq.of(_this.f12);
      let _957_v33;
      _957_v33 = _dafny.Seq.of(_module.__default.fm1(p1, p0, p0, globalState), p0, p0);
      let _958_v34;
      _958_v34 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_this.f11).cardinality())).multipliedBy((((_this.f11).contains((_956_v32)[_module.__default.safeIndex(new BigNumber((_957_v33).length), new BigNumber((_956_v32).length))])) ? ((_this.f11).get((_956_v32)[_module.__default.safeIndex(new BigNumber((_957_v33).length), new BigNumber((_956_v32).length))])) : (p0))),p0);
      let _959_v35;
      _959_v35 = _dafny.Seq.UnicodeFromString("hptlhvxc");
      let _960_v36;
      _960_v36 = new _dafny.CodePoint('c'.codePointAt(0));
      _958_v34 = (_958_v34).update(((_this.f12) ? (new BigNumber((_dafny.Seq.update(_959_v35, _module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_959_v35).length)), _960_v36)).length)) : (p0)), new BigNumber((_959_v35).length));
      let _961_v37;
      let _init21 = function (_962_i5) {
        return (_962_i5).plus(new BigNumber(758));
      };
      let _nw167 = Array((new BigNumber(29)).toNumber());
      for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw167.length); _i0_21++) {
        _nw167[_i0_21] = _init21(new BigNumber(_i0_21));
      }
      _961_v37 = _nw167;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_961_v37).length))) {
        let _963_i4 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_963_i4)) && ((_963_i4).isLessThan(new BigNumber((_961_v37).length))))) {
          (_961_v37)[(_963_i4)] = (_963_i4).multipliedBy((_module.D2.create_DC4(new BigNumber(-115))).dtor_cf8);
        }
      }
      r2 = ((_dafny.ZERO).minus(p0)).minus((p0).plus((_dafny.ZERO).minus(p0)));
      let _964_v38;
      _964_v38 = _dafny.MultiSet.fromElements(p0);
      let _965_v39;
      _965_v39 = _module.D6.create_DC15(_964_v38);
      let _966_v40;
      _966_v40 = _dafny.Set.fromElements(p1, true);
      let _967_v41;
      _967_v41 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,p0);
      r0 = (_this).fm2(((_965_v39).dtor_cf35).Intersect(_dafny.MultiSet.fromElements(p0, new BigNumber((_959_v35).length), new BigNumber((_966_v40).length))), p0, _967_v41, _this.f15, globalState);
      r1 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(880)), ((_968_v36) => function (_969_i6) {
        return _968_v36;
      })(_960_v36));
      r2 = _module.__default.safeModuloInt(p0, _module.__default.fm1(p1, (((_this.f11).contains(_this.f15)) ? ((_this.f11).get(_this.f15)) : (new BigNumber(-55))), new BigNumber((_966_v40).length), globalState));
      r3 = !(_this.f12);
      return [r0, r1, r2, r3];
    }
    m5(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      r1 = p0;
      let _970_v0;
      _970_v0 = new _dafny.CodePoint('a'.codePointAt(0));
      let _971_v1;
      _971_v1 = _module.D1.create_DC2(_this.f15, _this.f12, true, _970_v0, _970_v0);
      let _972_v2;
      _972_v2 = _dafny.Map.Empty.slice().updateUnsafe(_this.f15,_module.__default.fm14(_971_v1, new BigNumber(344), globalState));
      _972_v2 = (_972_v2).update(_this.f15, _970_v0);
      let _973_v3;
      let _nw168 = Array((new BigNumber(3)).toNumber());
      _nw168[(_dafny.ZERO).toNumber()] = _this.f15;
      _nw168[(_dafny.ONE).toNumber()] = _this.f12;
      _nw168[(new BigNumber(2)).toNumber()] = _this.f15;
      _973_v3 = _nw168;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_973_v3).length))) {
        let _974_i0 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_974_i0)) && ((_974_i0).isLessThan(new BigNumber((_973_v3).length))))) {
          (_973_v3)[(_974_i0)] = !(_this.f15);
        }
      }
      let _975_v4;
      _975_v4 = _dafny.Set.fromElements(_this.f12, _this.f15);
      let _976_v5;
      _976_v5 = _module.D1.create_DC1(_975_v4);
      let _977_v6;
      _977_v6 = _dafny.Set.fromElements(_976_v5);
      let _978_v7;
      _978_v7 = _dafny.Map.Empty.slice().updateUnsafe(_this.f15,(_977_v6).IsProperSubsetOf(_dafny.Set.fromElements(_module.D1.create_DC1(_975_v4))));
      if ((((_978_v7).contains(_this.f15)) ? ((_978_v7).get(_this.f15)) : (_this.f15))) {
        (globalState).f9 = p0;
        (_this).f12 = _this.f12;
        let _979_v8;
        _979_v8 = _dafny.Map.Empty.slice().updateUnsafe(false,p0);
        let _980_v9;
        _980_v9 = _dafny.Seq.of(p1, _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(591)), function (_981_i1) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        }), _module.__default.safeIndex((((_979_v8).contains(_this.f15)) ? ((_979_v8).get(_this.f15)) : (p0)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(591)), function (_982_i1) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        })).length)), _970_v0), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(591)), function (_983_i1) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        }), _module.__default.safeIndex((((_979_v8).contains(_this.f15)) ? ((_979_v8).get(_this.f15)) : (p0)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(591)), function (_984_i1) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        })).length)), _970_v0)).length)), _970_v0));
        let _985_v10;
        _985_v10 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((p1).length)), (new BigNumber((_980_v9).length)).minus(p0), ((_this.f15) ? (p0) : (p0)), p0, p0);
        let _986_v11;
        let _nw169 = new _module.C0();
        _nw169.__ctor(_this.f12, p0, _this.f11, _this.f15);
        _986_v11 = _nw169;
        (globalState).f9 = (_985_v10)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_986_v11, _986_v11)).length), new BigNumber((_985_v10).length))];
        let _987_v12;
        _987_v12 = _dafny.Seq.of(_dafny.Seq.Concat(_985_v10, _985_v10), _985_v10);
        _985_v10 = (_987_v12)[_module.__default.safeIndex((_986_v11).f18, new BigNumber((_987_v12).length))];
        let _988_v13;
        _988_v13 = _dafny.Set.fromElements(_970_v0);
        let _989_v14;
        _989_v14 = _dafny.Map.Empty.slice().updateUnsafe(_module.D5.create_DC14(),!(false));
        let _990_v15;
        _990_v15 = _module.D7.create_DC17(_989_v14);
        let _991_v16;
        let _nw170 = Array((new BigNumber(26)).toNumber());
        _nw170[(_dafny.ZERO).toNumber()] = p0;
        _nw170[(_dafny.ONE).toNumber()] = p0;
        _nw170[(new BigNumber(2)).toNumber()] = (_986_v11).f18;
        _nw170[(new BigNumber(3)).toNumber()] = p0;
        _nw170[(new BigNumber(4)).toNumber()] = (new BigNumber(300)).multipliedBy((_986_v11).f18);
        _nw170[(new BigNumber(5)).toNumber()] = (new BigNumber(631)).minus((_986_v11).f18);
        _nw170[(new BigNumber(6)).toNumber()] = (p0).plus(p0);
        _nw170[(new BigNumber(7)).toNumber()] = p0;
        _nw170[(new BigNumber(8)).toNumber()] = p0;
        _nw170[(new BigNumber(9)).toNumber()] = ((_986_v11).f18).plus(p0);
        _nw170[(new BigNumber(10)).toNumber()] = new BigNumber((_975_v4).length);
        _nw170[(new BigNumber(11)).toNumber()] = _dafny.ZERO;
        _nw170[(new BigNumber(12)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_988_v13).length), p0);
        _nw170[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(((_990_v15).dtor_cf39).length));
        _nw170[(new BigNumber(14)).toNumber()] = p0;
        _nw170[(new BigNumber(15)).toNumber()] = p0;
        _nw170[(new BigNumber(16)).toNumber()] = _module.__default.fm1(_this.f12, (_986_v11).f18, (_986_v11).f18, globalState);
        _nw170[(new BigNumber(17)).toNumber()] = (_986_v11).f18;
        _nw170[(new BigNumber(18)).toNumber()] = (((_this.f11).contains(_this.f12)) ? ((_this.f11).get(_this.f12)) : ((_986_v11).f18));
        _nw170[(new BigNumber(19)).toNumber()] = (_986_v11).f18;
        _nw170[(new BigNumber(20)).toNumber()] = (_986_v11).f18;
        _nw170[(new BigNumber(21)).toNumber()] = new BigNumber((p1).length);
        _nw170[(new BigNumber(22)).toNumber()] = ((_dafny.ZERO).minus((_986_v11).f18)).multipliedBy((_986_v11).f18);
        _nw170[(new BigNumber(23)).toNumber()] = ((_986_v11).f18).plus(new BigNumber((_this.f11).cardinality()));
        _nw170[(new BigNumber(24)).toNumber()] = (_986_v11).f18;
        _nw170[(new BigNumber(25)).toNumber()] = (_986_v11).f18;
        _991_v16 = _nw170;
        let _index139 = _module.__default.safeIndex(new BigNumber(371), new BigNumber((_991_v16).length));
        (_991_v16)[_index139] = (_985_v10)[_module.__default.safeIndex(p0, new BigNumber((_985_v10).length))];
        let _index140 = _module.__default.safeIndex(new BigNumber(371), new BigNumber((_991_v16).length));
        (_991_v16)[_index140] = (_dafny.ZERO).minus((p0).plus(new BigNumber(635)));
      } else {
        let _992_v17;
        let _nw171 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.of());
        _992_v17 = _nw171;
        let _index141 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_992_v17).length));
        (_992_v17)[_index141] = (_this).fm8(_this.f15, _this.f12, globalState);
        let _993_v18;
        _993_v18 = _dafny.Seq.of(p0);
        let _994_v19;
        _994_v19 = _module.D5.create_DC12(_993_v18);
        let _index142 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_992_v17).length));
        (_992_v17)[_index142] = _dafny.Seq.Concat((_module.D5.create_DC12(_993_v18)).dtor_cf29, (_994_v19).dtor_cf29);
        let _995_v20;
        let _init22 = ((_996_p0) => function (_997_i2) {
          return _module.__default.safeModuloInt(_997_i2, _996_p0);
        })(p0);
        let _nw172 = Array((new BigNumber(25)).toNumber());
        for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw172.length); _i0_22++) {
          _nw172[_i0_22] = _init22(new BigNumber(_i0_22));
        }
        _995_v20 = _nw172;
        let _998_v21;
        _998_v21 = _dafny.Map.Empty.slice().updateUnsafe(_this.f15,_995_v20);
        let _999_v22;
        _999_v22 = _dafny.Seq.of((((_998_v21).contains(_this.f15)) ? ((_998_v21).get(_this.f15)) : (_995_v20)));
        let _index143 = _module.__default.safeIndex(new BigNumber(260), new BigNumber((_973_v3).length));
        (_973_v3)[_index143] = _this.f12;
        let _1000_v23;
        _1000_v23 = _dafny.Seq.UnicodeFromString("muarlkhi");
        let _index144 = _module.__default.safeIndex(new BigNumber(260), new BigNumber((_973_v3).length));
        let _rhs153 = _dafny.Seq.Concat(_dafny.Seq.Concat(_999_v22, _999_v22), _999_v22);
        let _rhs154 = (_978_v7).equals(_978_v7);
        let _rhs155 = p1;
        let _lhs124 = _973_v3;
        let _lhs125 = _module.__default.safeIndex(new BigNumber(260), new BigNumber((_973_v3).length));
        _999_v22 = _rhs153;
        _lhs124[_lhs125] = _rhs154;
        _1000_v23 = _rhs155;
        let _1001_v24;
        _1001_v24 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1000_v23);
        let _1002_v25;
        _1002_v25 = _dafny.MultiSet.fromElements(p0);
        let _rhs156 = _dafny.Seq.IsPrefixOf(p1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(370)), function (_1003_i3) {
          return new _dafny.CodePoint('a'.codePointAt(0));
        }));
        let _rhs157 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt((((_1002_v25).contains(p0)) ? ((_1002_v25).get(p0)) : (p0)), p0), _module.__default.safeDivisionInt((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(p0))), new BigNumber((p1).length)));
        let _rhs158 = ((_1001_v24).update(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(108)), function (_1004_i4) {
          return new _dafny.CodePoint('f'.codePointAt(0));
        })).length), p1)).Merge(_1001_v24);
        let _lhs126 = globalState;
        let _lhs127 = globalState;
        _lhs126.f3 = _rhs156;
        _lhs127.f9 = _rhs157;
        _1001_v24 = _rhs158;
        let _1005_v26;
        _1005_v26 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_993_v18)[_module.__default.safeIndex(p0, new BigNumber((_993_v18).length))]);
        let _1006_v27;
        _1006_v27 = _dafny.Set.fromElements(_1005_v26, _1005_v26);
        let _1007_v28;
        _1007_v28 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,_module.__default.fm1(!(_this.f12), p0, new BigNumber((_1006_v27).length), globalState));
        _1007_v28 = _1007_v28;
        let _1008_v29;
        _1008_v29 = _dafny.Map.Empty.slice().updateUnsafe(_module.D6.create_DC15(_1002_v25),_1002_v25);
        let _1009_v30;
        _1009_v30 = _module.D6.create_DC15(_1002_v25);
        let _1010_v31;
        _1010_v31 = _dafny.Set.fromElements(new BigNumber((_1007_v28).length));
        let _1011_v32;
        _1011_v32 = _module.D4.create_DC11((_this).fm2((((_1008_v29).contains(_1009_v30)) ? ((_1008_v29).get(_1009_v30)) : (_1002_v25)), p0, _1007_v28, (_973_v3)[_module.__default.safeIndex(new BigNumber(260), new BigNumber((_973_v3).length))], globalState), p0, p0, _module.__default.fm1((_973_v3)[_module.__default.safeIndex(new BigNumber(260), new BigNumber((_973_v3).length))], (_dafny.ZERO).minus(new BigNumber((_1010_v31).length)), new BigNumber((_978_v7).length), globalState));
        let _1012_v33;
        _1012_v33 = _dafny.Map.Empty.slice().updateUnsafe(_995_v20,(_1011_v32).dtor_cf25);
        _1012_v33 = (_1012_v33).update(_995_v20, _this.f12);
      }
      let _1013_v34;
      _1013_v34 = _module.D8.create_DC19(_this.f11);
      let _1014_v35;
      _1014_v35 = _dafny.Seq.of(_this.f15);
      let _1015_v36;
      let _nw173 = new _module.C1();
      _nw173.__ctor(new BigNumber((p1).length), ((_1013_v34).dtor_cf45).update((_1014_v35)[_module.__default.safeIndex(new BigNumber(-857), new BigNumber((_1014_v35).length))], _module.__default.abs(new BigNumber(-462))), _this.f12);
      _1015_v36 = _nw173;
      (globalState).f3 = _module.__default.fm0(_970_v0, globalState);
      let _1016_v37;
      _1016_v37 = _dafny.Set.fromElements(p0);
      let _1017_v38;
      _1017_v38 = false;
      r0 = (((_dafny.Set.fromElements(p0, _1015_v36.f16)).IsSubsetOf(_1016_v37)) ? (_1017_v38) : (_1017_v38));
      r1 = (_1015_v36.f16).plus(p0);
      return [r0, r1];
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f11 = _dafny.MultiSet.Empty;
      this._f12 = false;
      this._f13 = _dafny.ZERO;
      this._f14 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
    set f11(value) {
      let _this = this;
      _this._f11 = value;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
    set f12(value) {
      let _this = this;
      _this._f12 = value;
    };
    __ctor(f13, f14, f11, f12) {
      let _this = this;
      (_this)._f13 = f13;
      (_this)._f14 = f14;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      return;
    }
    fm2(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !((_this).f14) || (_this.f12);
    };
    fm6(p0, p1, globalState) {
      let _this = this;
      if (((true) ? (_this.f12) : (_this.f12))) {
        if (_this.f12) {
          return (_this).f14;
        } else {
          return !(_this.f12);
        }
      } else {
        return ((_this).f13).isLessThan((_this).f13);
      }
    };
    m2(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      if ((_this).f14) {
        let _1018_v0;
        let _nw174 = Array((new BigNumber(29)).toNumber()).fill(_dafny.MultiSet.Empty);
        _1018_v0 = _nw174;
        let _1019_v1;
        _1019_v1 = new _dafny.CodePoint('d'.codePointAt(0));
        let _1020_v2;
        _1020_v2 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('a'.codePointAt(0)), _1019_v1);
        let _1021_v3;
        _1021_v3 = _dafny.Seq.of(_1020_v2);
        let _index145 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_1018_v0).length));
        (_1018_v0)[_index145] = (_1021_v3)[_module.__default.safeIndex((_this).f13, new BigNumber((_1021_v3).length))];
        let _index146 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_1018_v0).length));
        (_1018_v0)[_index146] = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(p0, _dafny.Seq.Concat(p0, _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.of(_1019_v1, _1019_v1), _module.__default.safeIndex((_this).f13, new BigNumber((_dafny.Seq.of(_1019_v1, _1019_v1)).length)), _1019_v1), _module.__default.safeIndex((_this).f13, new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_1019_v1, _1019_v1), _module.__default.safeIndex((_this).f13, new BigNumber((_dafny.Seq.of(_1019_v1, _1019_v1)).length)), _1019_v1)).length)), new _dafny.CodePoint('i'.codePointAt(0))))));
        if (p1) {
          let _1022_v4;
          _1022_v4 = _dafny.Set.fromElements((_this).f13, (_this).f13);
          let _1023_v5;
          let _nw175 = Array((new BigNumber(5)).toNumber());
          _nw175[(_dafny.ZERO).toNumber()] = _this.f12;
          _nw175[(_dafny.ONE).toNumber()] = !((_this).f13).isEqualTo((_this).f13);
          _nw175[(new BigNumber(2)).toNumber()] = (_1022_v4).IsDisjointFrom(_1022_v4);
          _nw175[(new BigNumber(3)).toNumber()] = (_this).f14;
          _nw175[(new BigNumber(4)).toNumber()] = p1;
          _1023_v5 = _nw175;
          let _index147 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_1023_v5).length));
          (_1023_v5)[_index147] = (p1) === ((_this).f14);
          let _1024_v6;
          _1024_v6 = _dafny.Seq.of(_1023_v5, _1023_v5, _1023_v5);
          let _1025_v7;
          _1025_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1024_v6,_this.f12);
          let _index148 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_1023_v5).length));
          (_1023_v5)[_index148] = (((_1025_v7).contains(_1024_v6)) ? ((_1025_v7).get(_1024_v6)) : (((p1) ? (false) : (p1))));
          let _1026_v8;
          let _nw176 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
          _1026_v8 = _nw176;
          let _index149 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((_1026_v8).length));
          (_1026_v8)[_index149] = (_this).f13;
          let _index150 = _module.__default.safeIndex(new BigNumber(557), new BigNumber((_1026_v8).length));
          (_1026_v8)[_index150] = (_this).f13;
          let _1027_v9;
          _1027_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_1026_v8, _1026_v8, _1026_v8)).length),p0);
          _1027_v9 = (_1027_v9).update(((_1026_v8)[_module.__default.safeIndex(new BigNumber(557), new BigNumber((_1026_v8).length))]).multipliedBy((_this).f13), p0);
          let _1028_v10;
          _1028_v10 = _module.D1.create_DC2(true, (_this).f14, (_1023_v5)[_module.__default.safeIndex(_dafny.ONE, new BigNumber((_1023_v5).length))], _1019_v1, new _dafny.CodePoint('t'.codePointAt(0)));
          let _1029_v11;
          _1029_v11 = _module.D1.create_DC3(_1028_v10);
          let _1030_v12;
          _1030_v12 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-76),_1029_v11);
          _1030_v12 = (_1030_v12).update((_this).f13, _module.D1.create_DC3(_1028_v10));
          let _1031_v13;
          _1031_v13 = p1;
          let _rhs159 = _1031_v13;
          let _rhs160 = p1;
          let _lhs128 = globalState;
          _1031_v13 = _rhs159;
          _lhs128.f7 = _rhs160;
        } else {
          let _1032_v14;
          _1032_v14 = _dafny.Seq.of(_this.f12);
          let _1033_v15;
          _1033_v15 = _module.D2.create_DC5((_this).f13, (_this).f13, (_this).f13, (_this).f13, _1032_v14);
          let _pat_let_tv20 = globalState;
          r0 = (new BigNumber((p0).length)).multipliedBy((function (_pat_let29_0) {
            return function (_1034_dt__update__tmp_h0) {
              return function (_pat_let30_0) {
                return function (_1035_dt__update_hcf13_h0) {
                  return function (_pat_let31_0) {
                    return function (_1036_dt__update_hcf10_h0) {
                      return _module.D2.create_DC5((_1034_dt__update__tmp_h0).dtor_cf9, _1036_dt__update_hcf10_h0, (_1034_dt__update__tmp_h0).dtor_cf11, (_1034_dt__update__tmp_h0).dtor_cf12, _1035_dt__update_hcf13_h0);
                    }(_pat_let31_0);
                  }((_this).f13);
                }(_pat_let30_0);
              }(_module.__default.fm7(_pat_let_tv20));
            }(_pat_let29_0);
          }(_1033_v15)).dtor_cf10);
          (globalState).f9 = (_this).f13;
          let _1037_v16;
          let _nw177 = new _module.C1();
          _nw177.__ctor((_this).f13, _this.f11, p1);
          _1037_v16 = _nw177;
          let _1038_v17;
          let _out17;
          _out17 = (_1037_v16).m6(globalState);
          _1038_v17 = _out17;
          let _1039_v18;
          let _nw178 = new _module.C1();
          _nw178.__ctor((_this).f13, _this.f11, (_this).f14);
          _1039_v18 = _nw178;
        }
        if (!((_this).f14)) {
          let _1040_v19;
          _1040_v19 = _dafny.Seq.of(new _dafny.CodePoint('k'.codePointAt(0)));
          let _1041_v20;
          let _nw179 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
          _1041_v20 = _nw179;
          let _index151 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_1041_v20).length));
          (_1041_v20)[_index151] = (_this).f13;
          let _1042_v21;
          _1042_v21 = _dafny.Set.fromElements((_this).f13, (_this).f13, _dafny.ZERO);
          let _1043_v23;
          _1043_v23 = _module.D7.create_DC18((_this).f13, _dafny.Seq.Create(_module.__default.abs(new BigNumber(787)), ((_1044_v1) => function (_1045_i0) {
  return _1044_v1;
})(_1019_v1)), _this.f12, (_this).f14, function () {
  let _coll24 = new _dafny.Set();
  for (const _compr_24 of (_1042_v21).Elements) {
    let _1046_v22 = _compr_24;
    if ((_1042_v21).contains(_1046_v22)) {
      _coll24.add((_1046_v22).plus(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality())));
    }
  }
  return _coll24;
}());
          let _index152 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_1041_v20).length));
          let _rhs161 = ((p1) ? (new BigNumber(43)) : ((_this).f13));
          let _rhs162 = (_1043_v23).dtor_cf41;
          let _rhs163 = ((_this).f13).multipliedBy((_this).f13);
          let _rhs164 = !(_this.f12);
          let _lhs129 = _1041_v20;
          let _lhs130 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_1041_v20).length));
          let _lhs131 = globalState;
          r0 = _rhs161;
          _1040_v19 = _rhs162;
          _lhs129[_lhs130] = _rhs163;
          _lhs131.f7 = _rhs164;
          let _1047_v24;
          _1047_v24 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f14);
          let _1048_v25;
          let _nw180 = Array((new BigNumber(10)).toNumber());
          _1048_v25 = _nw180;
          let _1049_v26;
          _1049_v26 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.MultiSet.fromElements(p1)).IsProperSubsetOf(_this.f11),_1048_v25);
          let _1050_v27;
          _1050_v27 = _module.D4.create_DC11(false, (_1041_v20)[_module.__default.safeIndex(new BigNumber(265), new BigNumber((_1041_v20).length))], (_this).f13, (_this).f13);
          let _1051_v28;
          _1051_v28 = _dafny.Map.Empty.slice().updateUnsafe((_1041_v20)[_module.__default.safeIndex(new BigNumber(265), new BigNumber((_1041_v20).length))],_1041_v20);
          let _rhs165 = _1047_v24;
          let _rhs166 = _1049_v26;
          let _rhs167 = _module.D4.create_DC11((_module.D8.create_DC20(p1, _this.f12, p1)).dtor_cf48, (_1041_v20)[_module.__default.safeIndex(new BigNumber(265), new BigNumber((_1041_v20).length))], (_1041_v20)[_module.__default.safeIndex(new BigNumber(265), new BigNumber((_1041_v20).length))], new BigNumber(((((_this).f14) ? (_1051_v28) : (_1051_v28))).length));
          _1047_v24 = _rhs165;
          _1049_v26 = _rhs166;
          _1050_v27 = _rhs167;
          let _1052_v29;
          let _nw181 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.of());
          _1052_v29 = _nw181;
          let _1053_v30;
          let _nw182 = Array((new BigNumber(2)).toNumber()).fill(_module.D6.Default());
          _1053_v30 = _nw182;
          let _index153 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_1052_v29).length));
          (_1052_v29)[_index153] = _dafny.Seq.of(_1053_v30, _1053_v30, _1053_v30, _1053_v30, _1053_v30);
          let _index154 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_1052_v29).length));
          (_1052_v29)[_index154] = _dafny.Seq.of(_1053_v30, _1053_v30);
          let _1054_v31;
          _1054_v31 = _this.f12;
          let _1055_v32;
          _1055_v32 = _dafny.Map.Empty.slice().updateUnsafe((_1054_v31),_1019_v1);
          (_this).m3(((_this).f13).minus(new BigNumber((_1055_v32).length)), globalState);
          let _index155 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_1041_v20).length));
          let _rhs168 = (_this).f13;
          let _rhs169 = new BigNumber(386);
          let _rhs170 = _this.f11;
          let _rhs171 = ((_1041_v20)[_module.__default.safeIndex(new BigNumber(265), new BigNumber((_1041_v20).length))]).plus(_module.__default.safeModuloInt(new BigNumber(918), (_1050_v27).dtor_cf28));
          let _lhs132 = globalState;
          let _lhs133 = _this;
          let _lhs134 = _1041_v20;
          let _lhs135 = _module.__default.safeIndex(new BigNumber(265), new BigNumber((_1041_v20).length));
          r0 = _rhs168;
          _lhs132.f9 = _rhs169;
          _lhs133.f11 = _rhs170;
          _lhs134[_lhs135] = _rhs171;
        } else {
          r1 = false;
          (globalState).f9 = (_this).f13;
          let _1056_v33;
          let _init23 = ((_1057_p0) => function (_1058_i1) {
            return (_dafny.Set.fromElements(_dafny.Set.fromElements(new BigNumber((_1057_p0).length)))).Union(_dafny.Set.fromElements(_dafny.Set.fromElements(new BigNumber((_1057_p0).length), (_this).f13)));
          })(p0);
          let _nw183 = Array((new BigNumber(14)).toNumber());
          for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw183.length); _i0_23++) {
            _nw183[_i0_23] = _init23(new BigNumber(_i0_23));
          }
          _1056_v33 = _nw183;
          let _1059_v34;
          _1059_v34 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,_this.f12);
          let _1060_v35;
          _1060_v35 = _dafny.Set.fromElements((_this).f13, new BigNumber((_1059_v34).length));
          let _1061_v36;
          _1061_v36 = _dafny.Seq.of((_this).f14);
          let _1062_v37;
          let _nw184 = new _module.C1();
          _nw184.__ctor((_this).f13, _dafny.MultiSet.fromElements(_this.f12, true), true);
          _1062_v37 = _nw184;
          let _1063_v38;
          _1063_v38 = _dafny.Seq.of(_1062_v37);
          let _1064_v39;
          _1064_v39 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_1063_v38)[_module.__default.safeIndex((_this).f13, new BigNumber((_1063_v38).length))]);
          let _1065_v40;
          _1065_v40 = _module.D2.create_DC5(new BigNumber((_1061_v36).length), new BigNumber((_1064_v39).length), (_this).f13, (_this).f13, _dafny.Seq.of(_1062_v37.f12, _1062_v37.f12));
          let _1066_v41;
          _1066_v41 = _dafny.Set.fromElements(_1060_v35, _1060_v35, _dafny.Set.fromElements((_1065_v40).dtor_cf12, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(706)), ((_1067_v1) => function (_1068_i2) {
            return _1067_v1;
          })(_1019_v1))).length)));
          let _index156 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_1056_v33).length));
          (_1056_v33)[_index156] = _1066_v41;
          let _1069_v42;
          let _nw185 = Array((new BigNumber(10)).toNumber()).fill(false);
          _1069_v42 = _nw185;
          let _index157 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_1056_v33).length));
          let _rhs172 = _1066_v41;
          let _rhs173 = _1069_v42;
          let _rhs174 = ((p1) ? (p1) : (_this.f12));
          let _lhs136 = _1056_v33;
          let _lhs137 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_1056_v33).length));
          let _lhs138 = _1062_v37;
          _lhs136[_lhs137] = _rhs172;
          _1069_v42 = _rhs173;
          _lhs138.f12 = _rhs174;
          let _index158 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_1069_v42).length));
          (_1069_v42)[_index158] = _1062_v37.f12;
          let _index159 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_1069_v42).length));
          (_1069_v42)[_index159] = p1;
          let _1070_v43;
          _1070_v43 = _dafny.MultiSet.fromElements(_1065_v40, _1065_v40);
          let _1071_v44;
          _1071_v44 = _dafny.Seq.of(_dafny.MultiSet.fromElements(_1065_v40, _1065_v40), (_1070_v43).update(_1065_v40, _module.__default.abs((_this).f13)), _1070_v43);
          _1071_v44 = _dafny.Seq.Concat(_1071_v44, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-912)), ((_1072_v43) => function (_1073_i3) {
            return _1072_v43;
          })(_1070_v43)));
        }
        let _1074_v45;
        let _nw186 = Array((new BigNumber(27)).toNumber());
        _1074_v45 = _nw186;
        let _1075_v46;
        _1075_v46 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1074_v45);
        let _1076_v47;
        _1076_v47 = _dafny.Seq.of((_this).f14);
        let _1077_v48;
        _1077_v48 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1((_this).f14, (_this).f13, new BigNumber((_1076_v47).length), globalState),_dafny.Map.Empty.slice().updateUnsafe((_this).f14,_1074_v45));
        let _1078_v49;
        _1078_v49 = _module.D9.create_DC23(_1075_v46);
        let _1079_v50;
        let _nw187 = Array((new BigNumber(25)).toNumber());
        _nw187[(_dafny.ZERO).toNumber()] = _1075_v46;
        _nw187[(_dafny.ONE).toNumber()] = ((_this.f12) ? (_1075_v46) : (_1075_v46));
        _nw187[(new BigNumber(2)).toNumber()] = (_1075_v46).update(true, _1074_v45);
        _nw187[(new BigNumber(3)).toNumber()] = _1075_v46;
        _nw187[(new BigNumber(4)).toNumber()] = (_1075_v46).Merge(_1075_v46);
        _nw187[(new BigNumber(5)).toNumber()] = _1075_v46;
        _nw187[(new BigNumber(6)).toNumber()] = _1075_v46;
        _nw187[(new BigNumber(7)).toNumber()] = _1075_v46;
        _nw187[(new BigNumber(8)).toNumber()] = _1075_v46;
        _nw187[(new BigNumber(9)).toNumber()] = (((_1077_v48).contains((((_this.f11).contains(true)) ? ((_this.f11).get(true)) : ((_this).f13)))) ? ((_1077_v48).get((((_this.f11).contains(true)) ? ((_this.f11).get(true)) : ((_this).f13)))) : (_1075_v46));
        _nw187[(new BigNumber(10)).toNumber()] = (_1078_v49).dtor_cf55;
        _nw187[(new BigNumber(11)).toNumber()] = ((_this.f12) ? (_1075_v46) : (_1075_v46));
        _nw187[(new BigNumber(12)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,_1074_v45);
        _nw187[(new BigNumber(13)).toNumber()] = _1075_v46;
        _nw187[(new BigNumber(14)).toNumber()] = _1075_v46;
        _nw187[(new BigNumber(15)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_this).f14,_1074_v45)).update(_this.f12, _1074_v45);
        _nw187[(new BigNumber(16)).toNumber()] = _1075_v46;
        _nw187[(new BigNumber(17)).toNumber()] = _1075_v46;
        _nw187[(new BigNumber(18)).toNumber()] = _1075_v46;
        _nw187[(new BigNumber(19)).toNumber()] = (_1075_v46).Merge(_1075_v46);
        _nw187[(new BigNumber(20)).toNumber()] = _1075_v46;
        _nw187[(new BigNumber(21)).toNumber()] = _1075_v46;
        _nw187[(new BigNumber(22)).toNumber()] = _1075_v46;
        _nw187[(new BigNumber(23)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(true,_1074_v45);
        _nw187[(new BigNumber(24)).toNumber()] = (_1075_v46).update(_this.f12, _1074_v45);
        _1079_v50 = _nw187;
        let _index160 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_1079_v50).length));
        (_1079_v50)[_index160] = ((p1) ? (_1075_v46) : (_1075_v46));
        let _1080_v51;
        _1080_v51 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_this).f13));
        let _1081_v52;
        _1081_v52 = _dafny.Seq.of(_1080_v51, _1080_v51);
        let _1082_v53;
        _1082_v53 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,(_this).f13);
        let _index161 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_1079_v50).length));
        let _rhs175 = (_this).f13;
        let _rhs176 = (_1075_v46).Merge((_1078_v49).dtor_cf55);
        let _rhs177 = !((_this).fm2((_1081_v52)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_1076_v47).length)), new BigNumber((_1081_v52).length))], (_this).f13, (_1082_v53).update(false, (_this).f13), (_this).f14, globalState));
        let _rhs178 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt((_this).f13, (_this).f13), (new BigNumber((_this.f11).cardinality())).plus((_this).f13));
        let _lhs139 = _1079_v50;
        let _lhs140 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_1079_v50).length));
        let _lhs141 = globalState;
        r0 = _rhs175;
        _lhs139[_lhs140] = _rhs176;
        _lhs141.f3 = _rhs177;
        r0 = _rhs178;
        let _1083_v54;
        let _nw188 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
        _1083_v54 = _nw188;
        let _index162 = _module.__default.safeIndex(new BigNumber(68), new BigNumber((_1083_v54).length));
        (_1083_v54)[_index162] = (_this).f13;
        let _index163 = _module.__default.safeIndex(new BigNumber(68), new BigNumber((_1083_v54).length));
        (_1083_v54)[_index163] = (_this).f13;
      } else {
        let _1084_v55;
        _1084_v55 = new _dafny.CodePoint('y'.codePointAt(0));
        let _1085_v56;
        _1085_v56 = _dafny.MultiSet.fromElements(_1084_v55, _1084_v55);
        let _1086_v57;
        let _nw189 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.of());
        _1086_v57 = _nw189;
        let _1087_v58;
        _1087_v58 = _dafny.Seq.of((_this).f13);
        let _index164 = _module.__default.safeIndex(new BigNumber(204), new BigNumber((_1086_v57).length));
        (_1086_v57)[_index164] = _1087_v58;
        let _index165 = _module.__default.safeIndex(new BigNumber(204), new BigNumber((_1086_v57).length));
        let _rhs179 = _dafny.MultiSet.fromElements(_1084_v55);
        let _rhs180 = (_this).f14;
        let _rhs181 = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_this).f13), (_this).f13);
        let _rhs182 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1087_v58, _1087_v58), _1087_v58);
        let _lhs142 = globalState;
        let _lhs143 = _1086_v57;
        let _lhs144 = _module.__default.safeIndex(new BigNumber(204), new BigNumber((_1086_v57).length));
        _1085_v56 = _rhs179;
        r1 = _rhs180;
        _lhs142.f9 = _rhs181;
        _lhs143[_lhs144] = _rhs182;
        let _1088_v59;
        _1088_v59 = _dafny.MultiSet.fromElements((_this).f13);
        let _1089_v60;
        _1089_v60 = _module.D1.create_DC2(_this.f12, _this.f12, _this.f12, _1084_v55, _1084_v55);
        let _1090_v61;
        _1090_v61 = _module.D1.create_DC2(p1, false, !(p1), _1084_v55, _module.__default.fm14(_1089_v60, (_this).f13, globalState));
        let _1091_v62;
        let _nw190 = Array((new BigNumber(23)).toNumber());
        _nw190[(_dafny.ZERO).toNumber()] = p0;
        _nw190[(_dafny.ONE).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(833)), ((_1092_v55) => function (_1093_i4) {
          return _1092_v55;
        })(_1084_v55));
        _nw190[(new BigNumber(2)).toNumber()] = p0;
        _nw190[(new BigNumber(3)).toNumber()] = p0;
        _nw190[(new BigNumber(4)).toNumber()] = p0;
        _nw190[(new BigNumber(5)).toNumber()] = p0;
        _nw190[(new BigNumber(6)).toNumber()] = p0;
        _nw190[(new BigNumber(7)).toNumber()] = p0;
        _nw190[(new BigNumber(8)).toNumber()] = p0;
        _nw190[(new BigNumber(9)).toNumber()] = p0;
        _nw190[(new BigNumber(10)).toNumber()] = p0;
        _nw190[(new BigNumber(11)).toNumber()] = p0;
        _nw190[(new BigNumber(12)).toNumber()] = p0;
        _nw190[(new BigNumber(13)).toNumber()] = _dafny.Seq.UnicodeFromString("ppym");
        _nw190[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(p0, p0);
        _nw190[(new BigNumber(15)).toNumber()] = p0;
        _nw190[(new BigNumber(16)).toNumber()] = p0;
        _nw190[(new BigNumber(17)).toNumber()] = _dafny.Seq.Concat(p0, p0);
        _nw190[(new BigNumber(18)).toNumber()] = (((_this).f14) ? (p0) : (_dafny.Seq.update(p0, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber(((_1088_v59).update((_this).f13, _module.__default.abs((_dafny.ZERO).minus((((_1088_v59).contains((_this).f13)) ? ((_1088_v59).get((_this).f13)) : ((_this).f13)))))).cardinality())), new BigNumber((p0).length)), (_1090_v61).dtor_cf6)));
        _nw190[(new BigNumber(19)).toNumber()] = p0;
        _nw190[(new BigNumber(20)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-542)), ((_1094_v55) => function (_1095_i5) {
          return _1094_v55;
        })(_1084_v55));
        _nw190[(new BigNumber(21)).toNumber()] = _dafny.Seq.UnicodeFromString("drfnfyqmg");
        _nw190[(new BigNumber(22)).toNumber()] = p0;
        _1091_v62 = _nw190;
        let _1096_v63;
        _1096_v63 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,new BigNumber((_dafny.Seq.UnicodeFromString("whjni")).length));
        let _1097_v64;
        _1097_v64 = _dafny.Seq.of(true);
        let _1098_v65;
        _1098_v65 = _module.D2.create_DC5((_this).f13, new BigNumber((_1085_v56).cardinality()), (_this).f13, (((_1096_v63).contains((_this).f14)) ? ((_1096_v63).get((_this).f14)) : ((_this).f13)), _1097_v64);
        let _pat_let_tv21 = _1097_v64;
        let _1099_v66;
        _1099_v66 = _dafny.Map.Empty.slice().updateUnsafe(((_this.f12) ? (_1098_v65) : (function (_pat_let32_0) {
          return function (_1100_dt__update__tmp_h1) {
            return function (_pat_let33_0) {
              return function (_1101_dt__update_hcf13_h1) {
                return function (_pat_let34_0) {
                  return function (_1102_dt__update_hcf9_h0) {
                    return _module.D2.create_DC5(_1102_dt__update_hcf9_h0, (_1100_dt__update__tmp_h1).dtor_cf10, (_1100_dt__update__tmp_h1).dtor_cf11, (_1100_dt__update__tmp_h1).dtor_cf12, _1101_dt__update_hcf13_h1);
                  }(_pat_let34_0);
                }((_this).f13);
              }(_pat_let33_0);
            }(_pat_let_tv21);
          }(_pat_let32_0);
        }(_1098_v65))),_1091_v62);
        let _1103_v67;
        let _nw191 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
        _1103_v67 = _nw191;
        let _1104_v68;
        _1104_v68 = _dafny.Seq.of(_1103_v67, _1103_v67);
        let _1105_v70;
        _1105_v70 = _dafny.Set.fromElements((_this).f13, (_this).f13);
        let _1106_v71;
        _1106_v71 = _dafny.Seq.of(function () {
          let _coll25 = new _dafny.Set();
          for (const _compr_25 of (_1088_v59).Elements) {
            let _1107_v69 = _compr_25;
            if ((_1088_v59).contains(_1107_v69)) {
              _coll25.add((_1107_v69).multipliedBy(new BigNumber(-388)));
            }
          }
          return _coll25;
        }(), _1105_v70, _dafny.Set.fromElements(new BigNumber((_1097_v64).length), (_this).f13));
        let _1108_v72;
        _1108_v72 = _dafny.Seq.of(_1103_v67, _1103_v67, (_1104_v68)[_module.__default.safeIndex(((_1086_v57)[_module.__default.safeIndex(new BigNumber(204), new BigNumber((_1086_v57).length))])[_module.__default.safeIndex(_module.__default.fm1(p1, (_this).f13, new BigNumber((_1106_v71).length), globalState), new BigNumber(((_1086_v57)[_module.__default.safeIndex(new BigNumber(204), new BigNumber((_1086_v57).length))]).length))], new BigNumber((_1104_v68).length))]);
        let _1109_v73;
        _1109_v73 = _module.D9.create_DC24(_dafny.Seq.UnicodeFromString("c"), _this.f12, (_this).f13);
        let _index166 = _module.__default.safeIndex(new BigNumber(204), new BigNumber((_1086_v57).length));
        let _rhs183 = new BigNumber(-676);
        let _rhs184 = (((_1099_v66).contains(function (_pat_let37_0) {
          return function (_1112_dt__update__tmp_h2) {
            return function (_pat_let38_0) {
              return function (_1113_dt__update_hcf9_h1) {
                return _module.D2.create_DC5(_1113_dt__update_hcf9_h1, (_1112_dt__update__tmp_h2).dtor_cf10, (_1112_dt__update__tmp_h2).dtor_cf11, (_1112_dt__update__tmp_h2).dtor_cf12, (_1112_dt__update__tmp_h2).dtor_cf13);
              }(_pat_let38_0);
            }(new BigNumber(697));
          }(_pat_let37_0);
        }(_1098_v65))) ? ((_1099_v66).get(function (_pat_let35_0) {
          return function (_1110_dt__update__tmp_h3) {
            return function (_pat_let36_0) {
              return function (_1111_dt__update_hcf9_h2) {
                return _module.D2.create_DC5(_1111_dt__update_hcf9_h2, (_1110_dt__update__tmp_h3).dtor_cf10, (_1110_dt__update__tmp_h3).dtor_cf11, (_1110_dt__update__tmp_h3).dtor_cf12, (_1110_dt__update__tmp_h3).dtor_cf13);
              }(_pat_let36_0);
            }(new BigNumber(697));
          }(_pat_let35_0);
        }(_1098_v65))) : (p2));
        let _rhs185 = _dafny.Seq.Concat(_1087_v58, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-730)), function (_1114_i6) {
          return (_this).f13;
        }));
        let _rhs186 = (_1104_v68)[_module.__default.safeIndex(_module.__default.safeModuloInt((_this).f13, (_this).f13), new BigNumber((_1104_v68).length))];
        let _rhs187 = (function (_pat_let39_0) {
          return function (_1115_dt__update__tmp_h4) {
            return function (_pat_let40_0) {
              return function (_1116_dt__update_hcf58_h0) {
                return _module.D9.create_DC24((_1115_dt__update__tmp_h4).dtor_cf56, (_1115_dt__update__tmp_h4).dtor_cf57, _1116_dt__update_hcf58_h0);
              }(_pat_let40_0);
            }((_this).f13);
          }(_pat_let39_0);
        }(_1109_v73)).dtor_cf57;
        let _lhs145 = globalState;
        let _lhs146 = _1086_v57;
        let _lhs147 = _module.__default.safeIndex(new BigNumber(204), new BigNumber((_1086_v57).length));
        let _lhs148 = globalState;
        let _lhs149 = globalState;
        _lhs145.f9 = _rhs183;
        _1091_v62 = _rhs184;
        _lhs146[_lhs147] = _rhs185;
        _lhs148.f4 = _rhs186;
        _lhs149.f7 = _rhs187;
        let _1117_v74;
        _1117_v74 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,_this.f12);
        _1117_v74 = (_1117_v74).update(_module.__default.safeDivisionInt((_this).f13, (_this).f13), !(p1) || (_this.f12));
        r0 = (_dafny.ZERO).minus((_this).f13);
        let _1118_v75;
        let _nw192 = new _module.C0();
        _nw192.__ctor((_this).f14, (_this).f13, (_this.f11).update(_this.f12, _module.__default.abs((_this).f13)), (_this).f14);
        _1118_v75 = _nw192;
        let _1119_v76;
        _1119_v76 = _module.D3.create_DC9((_this).f13, _module.__default.safeModuloInt((((_1088_v59).contains(_module.__default.fm1(_this.f12, (_this).f13, (_this).f13, globalState))) ? ((_1088_v59).get(_module.__default.fm1(_this.f12, (_this).f13, (_this).f13, globalState))) : ((_this).f13)), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update((_1086_v57)[_module.__default.safeIndex(new BigNumber(204), new BigNumber((_1086_v57).length))], _module.__default.safeIndex(new BigNumber(((_1086_v57)[_module.__default.safeIndex(new BigNumber(204), new BigNumber((_1086_v57).length))]).length), new BigNumber(((_1086_v57)[_module.__default.safeIndex(new BigNumber(204), new BigNumber((_1086_v57).length))]).length)), new BigNumber(-968))).length))), _1118_v75, (_1118_v75).f18);
        _1119_v76 = _1119_v76;
      }
      let _1120_v77;
      _1120_v77 = _dafny.Seq.of(!((_this).f14), _this.f12, p1);
      let _rhs188 = (_this).f14;
      let _rhs189 = (_1120_v77)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f13,(_this).f13)).length), new BigNumber((_1120_v77).length))];
      let _lhs150 = globalState;
      r1 = _rhs188;
      _lhs150.f3 = _rhs189;
      let _1121_v78;
      _1121_v78 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f12),(_this).f13);
      let _1122_v79;
      _1122_v79 = _dafny.MultiSet.fromElements(_1121_v78);
      let _1123_v80;
      _1123_v80 = _module.D10.create_DC27(_1121_v78);
      let _1124_v81;
      _1124_v81 = _dafny.Seq.of(new BigNumber(((_1121_v78).update(p1, (_dafny.ZERO).minus((_this).f13))).length), (_this).f13, (_this).f13);
      let _1125_v82;
      let _nw193 = Array((new BigNumber(13)).toNumber());
      _nw193[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_this).f14,new BigNumber((p0).length)), _1121_v78, _dafny.Map.Empty.slice().updateUnsafe(_this.f12,(_this).f13));
      _nw193[(_dafny.ONE).toNumber()] = _dafny.MultiSet.fromElements(_1121_v78, _1121_v78);
      _nw193[(new BigNumber(2)).toNumber()] = (_1122_v79).Intersect(_1122_v79);
      _nw193[(new BigNumber(3)).toNumber()] = _1122_v79;
      _nw193[(new BigNumber(4)).toNumber()] = _1122_v79;
      _nw193[(new BigNumber(5)).toNumber()] = _dafny.MultiSet.fromElements((_1123_v80).dtor_cf63);
      _nw193[(new BigNumber(6)).toNumber()] = _1122_v79;
      _nw193[(new BigNumber(7)).toNumber()] = _1122_v79;
      _nw193[(new BigNumber(8)).toNumber()] = _1122_v79;
      _nw193[(new BigNumber(9)).toNumber()] = _1122_v79;
      _nw193[(new BigNumber(10)).toNumber()] = _1122_v79;
      _nw193[(new BigNumber(11)).toNumber()] = (_1122_v79).update(_1121_v78, _module.__default.abs(new BigNumber((_1124_v81).length)));
      _nw193[(new BigNumber(12)).toNumber()] = (_1122_v79).Difference(_module.__default.fm19((_this).f13, globalState));
      _1125_v82 = _nw193;
      let _index167 = _module.__default.safeIndex(new BigNumber(599), new BigNumber((_1125_v82).length));
      (_1125_v82)[_index167] = _1122_v79;
      let _index168 = _module.__default.safeIndex(new BigNumber(599), new BigNumber((_1125_v82).length));
      (_1125_v82)[_index168] = (_1122_v79).Difference((_1122_v79).Union(_1122_v79));
      let _1126_v83;
      _1126_v83 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_this).f13),(_this).f13);
      if ((_1126_v83).contains(_1124_v81)) {
        let _1127_v84;
        _1127_v84 = _dafny.Map.Empty.slice().updateUnsafe(_1120_v77,p1);
        _1127_v84 = (_1127_v84).update(_dafny.Seq.Concat(_1120_v77, _dafny.Seq.update(_1120_v77, _module.__default.safeIndex(new BigNumber((_this.f11).cardinality()), new BigNumber((_1120_v77).length)), !(p1))), _this.f12);
        let _1128_v85;
        let _nw194 = new _module.C2();
        _nw194.__ctor((_this).f14, _this.f11, (_this).f14);
        _1128_v85 = _nw194;
        _1128_v85 = _1128_v85;
        let _1129_v86;
        _1129_v86 = _dafny.Seq.UnicodeFromString("kfcc");
        _1129_v86 = _dafny.Seq.Concat(_1129_v86, p0);
        let _1130_v87;
        _1130_v87 = _module.D5.create_DC14();
        let _1131_v88;
        _1131_v88 = _module.D7.create_DC17(_dafny.Map.Empty.slice().updateUnsafe(_1130_v87,p1));
        let _1132_v89;
        _1132_v89 = _dafny.Map.Empty.slice().updateUnsafe(_1130_v87,_this.f12);
        let _pat_let_tv22 = _1132_v89;
        let _pat_let_tv23 = _1132_v89;
        _1131_v88 = function (_pat_let41_0) {
          return function (_1133_dt__update__tmp_h5) {
            return function (_pat_let42_0) {
              return function (_1134_dt__update_hcf39_h0) {
                return _module.D7.create_DC17(_1134_dt__update_hcf39_h0);
              }(_pat_let42_0);
            }((_pat_let_tv22).Merge(_pat_let_tv23));
          }(_pat_let41_0);
        }(_1131_v88);
        (globalState).f7 = ((_1128_v85.f15) ? (_1128_v85.f15) : (_this.f12));
      } else {
        (globalState).f8 = (_this).f14;
        let _1135_v90;
        _1135_v90 = new _dafny.CodePoint('s'.codePointAt(0));
        let _1136_v91;
        _1136_v91 = _module.D1.create_DC2(_this.f12, false, (_this).f14, _1135_v90, _1135_v90);
        let _1137_v92;
        _1137_v92 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1124_v81).length),_module.__default.fm14(_1136_v91, (_this).f13, globalState));
        let _1138_v93;
        _1138_v93 = _dafny.Set.fromElements((_this).f14);
        let _1139_v94;
        _1139_v94 = _dafny.MultiSet.fromElements((_this).f13);
        let _1140_v95;
        _1140_v95 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f13),(_this).f13);
        let _1141_v96;
        _1141_v96 = _dafny.Set.fromElements((_this).f13, new BigNumber((_1140_v95).length), (_this).f13);
        let _1142_v97;
        _1142_v97 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm2(_1139_v94, (_this).f13, _1121_v78, p1, globalState),_1141_v96);
        let _1143_v98;
        let _nw195 = Array((new BigNumber(24)).toNumber());
        _nw195[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1124_v81).length));
        _nw195[(_dafny.ONE).toNumber()] = _module.__default.fm1((_this).f14, (_this).f13, new BigNumber(375), globalState);
        _nw195[(new BigNumber(2)).toNumber()] = new BigNumber((_1137_v92).length);
        _nw195[(new BigNumber(3)).toNumber()] = (_this).f13;
        _nw195[(new BigNumber(4)).toNumber()] = (_this).f13;
        _nw195[(new BigNumber(5)).toNumber()] = (_this).f13;
        _nw195[(new BigNumber(6)).toNumber()] = (_this).f13;
        _nw195[(new BigNumber(7)).toNumber()] = new BigNumber((_1138_v93).length);
        _nw195[(new BigNumber(8)).toNumber()] = (_this).f13;
        _nw195[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("mwn")).length);
        _nw195[(new BigNumber(10)).toNumber()] = new BigNumber((p0).length);
        _nw195[(new BigNumber(11)).toNumber()] = (_this).f13;
        _nw195[(new BigNumber(12)).toNumber()] = (_this).f13;
        _nw195[(new BigNumber(13)).toNumber()] = (_this).f13;
        _nw195[(new BigNumber(14)).toNumber()] = (_this).f13;
        _nw195[(new BigNumber(15)).toNumber()] = new BigNumber((_1142_v97).length);
        _nw195[(new BigNumber(16)).toNumber()] = (_this).f13;
        _nw195[(new BigNumber(17)).toNumber()] = (_this).f13;
        _nw195[(new BigNumber(18)).toNumber()] = _module.__default.fm1(p1, new BigNumber((_1121_v78).length), (_this).f13, globalState);
        _nw195[(new BigNumber(19)).toNumber()] = (_this).f13;
        _nw195[(new BigNumber(20)).toNumber()] = (_dafny.ZERO).minus((_this).f13);
        _nw195[(new BigNumber(21)).toNumber()] = (_this).f13;
        _nw195[(new BigNumber(22)).toNumber()] = (_this).f13;
        _nw195[(new BigNumber(23)).toNumber()] = (_this).f13;
        _1143_v98 = _nw195;
        let _1144_v99;
        _1144_v99 = _dafny.Set.fromElements(_1143_v98);
        _1144_v99 = (((_this).f14) ? ((_1144_v99).Difference(_1144_v99)) : ((_1144_v99).Union(_1144_v99)));
        let _1145_v100;
        _1145_v100 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,_1143_v98);
        _1145_v100 = (_1145_v100).update(_this.f12, _1143_v98);
        (_this).f12 = !_dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(558)), function (_1146_i7) {
          return (_this).f13;
        }), _1124_v81), _1124_v81);
        let _1147_v101;
        _1147_v101 = _dafny.Seq.UnicodeFromString("m");
        _1147_v101 = p0;
      }
      let _1148_v102;
      let _nw196 = new _module.C0();
      _nw196.__ctor(!((_this).f14), (_this).f13, _this.f11, p1);
      _1148_v102 = _nw196;
      (_this).f12 = _this.f12;
      r0 = (_dafny.ZERO).minus(_module.__default.fm1(_this.f12, (_this).f13, _module.__default.safeModuloInt((_this).f13, new BigNumber(120)), globalState));
      let _1149_v103;
      _1149_v103 = _module.D5.create_DC12(_dafny.Seq.Create(_module.__default.abs(new BigNumber(179)), function (_1150_i8) {
  return new BigNumber(382);
}));
      let _1151_v104;
      _1151_v104 = _dafny.MultiSet.fromElements(_1149_v103, _1149_v103, _1149_v103);
      r1 = (_1151_v104).IsSubsetOf(_1151_v104);
      return [r0, r1];
    }
    m3(p0, globalState) {
      let _this = this;
      let _1152_v0;
      let _nw197 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
      _1152_v0 = _nw197;
      let _index169 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length));
      (_1152_v0)[_index169] = p0;
      let _index170 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length));
      (_1152_v0)[_index170] = p0;
      let _1153_v1;
      let _nw198 = Array((new BigNumber(13)).toNumber()).fill(_module.D4.Default());
      _1153_v1 = _nw198;
      let _1154_v2;
      _1154_v2 = _module.D4.create_DC10(_1152_v0);
      let _pat_let_tv24 = _1152_v0;
      let _index171 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_1153_v1).length));
      (_1153_v1)[_index171] = function (_pat_let43_0) {
        return function (_1155_dt__update__tmp_h0) {
          return function (_pat_let44_0) {
            return function (_1156_dt__update_hcf24_h0) {
              return _module.D4.create_DC10(_1156_dt__update_hcf24_h0);
            }(_pat_let44_0);
          }(_pat_let_tv24);
        }(_pat_let43_0);
      }(_1154_v2);
      let _index172 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_1153_v1).length));
      (_1153_v1)[_index172] = _module.D4.create_DC10(_1152_v0);
      if ((_this).f14) {
        let _1157_v3;
        _1157_v3 = new _dafny.CodePoint('m'.codePointAt(0));
        let _1158_v4;
        _1158_v4 = _module.D1.create_DC2((_this).f14, (_this).f14, _this.f12, _1157_v3, _1157_v3);
        let _1159_v5;
        _1159_v5 = _dafny.Map.Empty.slice().updateUnsafe(p0,_module.__default.fm14(_1158_v4, p0, globalState));
        let _1160_v6;
        _1160_v6 = _dafny.Set.fromElements(p0, (_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))]);
        let _rhs190 = !((_1160_v6).IsSubsetOf(_module.__default.fm20(_1157_v3, new BigNumber(678), new BigNumber(388), _1159_v5, globalState)));
        let _rhs191 = new BigNumber(-307);
        let _lhs151 = _this;
        let _lhs152 = globalState;
        _lhs151.f12 = _rhs190;
        _lhs152.f9 = _rhs191;
        let _index173 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length));
        (_1152_v0)[_index173] = (_this).f13;
        let _1161_v7;
        _1161_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1152_v0,new BigNumber((_1160_v6).length));
        let _1162_v8;
        _1162_v8 = _dafny.Seq.of(_1152_v0, _1152_v0);
        let _1163_v9;
        _1163_v9 = _dafny.Set.fromElements(_this.f12);
        let _1164_v10;
        _1164_v10 = _dafny.Seq.of(new BigNumber((_1163_v9).length), p0, (_this).f13, (_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))]);
        _1161_v7 = (_1161_v7).update((_1162_v8)[_module.__default.safeIndex(new BigNumber((_1164_v10).length), new BigNumber((_1162_v8).length))], p0);
        let _1165_v11;
        _1165_v11 = _dafny.Seq.UnicodeFromString("hoftygcb");
        _1165_v11 = _1165_v11;
        (globalState).f9 = (p0).multipliedBy((_this).f13);
      } else {
        (globalState).f9 = new BigNumber(701);
        let _1166_v12;
        let _init24 = function (_1167_i0) {
          return _dafny.Seq.UnicodeFromString("ipdww");
        };
        let _nw199 = Array((new BigNumber(28)).toNumber());
        for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw199.length); _i0_24++) {
          _nw199[_i0_24] = _init24(new BigNumber(_i0_24));
        }
        _1166_v12 = _nw199;
        let _1168_v13;
        _1168_v13 = _module.D5.create_DC13((_this).f13, _this.f12, (p0).minus(p0), _1166_v12, (_this).f14);
        let _source10 = _1168_v13;
        if (_source10.is_DC13) {
          let _1169___mcc_h0 = (_source10).cf30;
          let _1170___mcc_h1 = (_source10).cf31;
          let _1171___mcc_h2 = (_source10).cf32;
          let _1172___mcc_h3 = (_source10).cf33;
          let _1173___mcc_h4 = (_source10).cf34;
          let _1174_cf34 = _1173___mcc_h4;
          let _1175_cf33 = _1172___mcc_h3;
          let _1176_cf32 = _1171___mcc_h2;
          let _1177_cf31 = _1170___mcc_h1;
          let _1178_cf30 = _1169___mcc_h0;
          let _1179_v14;
          let _nw200 = new _module.C2();
          _nw200.__ctor(_this.f12, (_this.f11).Difference(_dafny.MultiSet.fromElements((_this).f14)), _1174_cf34);
          _1179_v14 = _nw200;
          _1179_v14 = _1179_v14;
          _1176_cf32 = _module.__default.safeModuloInt((_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))], (new BigNumber((_dafny.Seq.UnicodeFromString("klkgrol")).length)).multipliedBy(_1176_cf32));
          let _1180_v15;
          let _init25 = function (_1181_i1) {
            return _this.f12;
          };
          let _nw201 = Array((new BigNumber(8)).toNumber());
          for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw201.length); _i0_25++) {
            _nw201[_i0_25] = _init25(new BigNumber(_i0_25));
          }
          _1180_v15 = _nw201;
          let _index174 = _module.__default.safeIndex(new BigNumber(803), new BigNumber((_1180_v15).length));
          (_1180_v15)[_index174] = _this.f12;
          let _index175 = _module.__default.safeIndex(new BigNumber(803), new BigNumber((_1180_v15).length));
          (_1180_v15)[_index175] = (_this).f14;
          let _1182_v16;
          _1182_v16 = _dafny.Map.Empty.slice().updateUnsafe(_1179_v14.f15,_1180_v15);
          (globalState).f3 = ((_1182_v16).Merge(_1182_v16)).equals(_1182_v16);
        } else if (_source10.is_DC14) {
          (globalState).f9 = _module.__default.fm1((_this).f14, new BigNumber(12), (new BigNumber((_dafny.Seq.of(new BigNumber(86), new BigNumber(617), new BigNumber((_this.f11).cardinality()))).length)).minus((_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))]), globalState);
          let _1183_v17;
          _1183_v17 = _dafny.Seq.of(_this.f12);
          let _1184_v18;
          _1184_v18 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(false,false),p0);
          let _1185_v19;
          _1185_v19 = _dafny.Seq.UnicodeFromString("x");
          let _1186_v20;
          _1186_v20 = _dafny.Set.fromElements(new BigNumber((_1185_v19).length));
          let _1187_v21;
          let _nw202 = new _module.C0();
          _nw202.__ctor(_this.f12, p0, (_this.f11).Difference((_dafny.MultiSet.FromArray(_1183_v17)).update((_this).f14, _module.__default.abs(new BigNumber((_1183_v17).length)))), (_dafny.Set.fromElements((_dafny.ZERO).minus((((_1184_v18).contains(_dafny.Map.Empty.slice().updateUnsafe((false),(_this).f14))) ? ((_1184_v18).get(_dafny.Map.Empty.slice().updateUnsafe((false),(_this).f14))) : (new BigNumber(-628)))), (_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))])).IsSubsetOf(_1186_v20));
          _1187_v21 = _nw202;
          let _nw203 = new _module.C0();
          _nw203.__ctor((_this).f14, (_this).f13, _dafny.MultiSet.fromElements((_1187_v21).f17), (_this).f14);
          _1187_v21 = _nw203;
          let _1188_v22;
          _1188_v22 = _module.D9.create_DC25(_1185_v19, (_1187_v21).f17, p0);
          let _1189_v23;
          _1189_v23 = new _dafny.CodePoint('y'.codePointAt(0));
          _1185_v19 = _dafny.Seq.update((_1188_v22).dtor_cf59, _module.__default.safeIndex((_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))], new BigNumber(((_1188_v22).dtor_cf59).length)), _1189_v23);
          let _1190_v24;
          _1190_v24 = _module.D2.create_DC5(p0, (_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))], new BigNumber(223), p0, _1183_v17);
          let _rhs192 = _this.f12;
          let _rhs193 = _1189_v23;
          let _rhs194 = (_module.__default.safeModuloInt((_1187_v21).f18, (_1187_v21).fm12(globalState))).isLessThan((new BigNumber(((_1190_v24).dtor_cf13).length)).plus((_this).f13));
          let _rhs195 = _module.__default.safeModuloInt(new BigNumber((_1186_v20).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1189_v23,_this.f12)).length));
          let _lhs153 = globalState;
          let _lhs154 = globalState;
          let _lhs155 = globalState;
          _lhs153.f8 = _rhs192;
          _1189_v23 = _rhs193;
          _lhs154.f7 = _rhs194;
          _lhs155.f9 = _rhs195;
        } else {
          let _1191___mcc_h5 = (_source10).cf29;
          let _1192_cf29 = _1191___mcc_h5;
          let _1193_v25;
          let _nw204 = new _module.C2();
          _nw204.__ctor((_this).f14, _this.f11, _this.f12);
          _1193_v25 = _nw204;
          _1193_v25 = _1193_v25;
          let _1194_v26;
          _1194_v26 = new _dafny.CodePoint('c'.codePointAt(0));
          let _1195_v27;
          _1195_v27 = _dafny.Seq.of((_this).f14);
          let _1196_v28;
          _1196_v28 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1195_v27).length),(_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))]);
          let _1197_v29;
          let _nw205 = new _module.C0();
          _nw205.__ctor(_module.__default.fm0(_1194_v26, globalState), _module.__default.fm1(_1193_v25.f12, (_this).f13, (_this).f13, globalState), _dafny.MultiSet.fromElements(!((_this).f14), (_this).f14, _this.f12, _this.f12, !((_this).f14)), !(((((_1196_v28).contains((_this).f13)) ? ((_1196_v28).get((_this).f13)) : ((_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))]))).isEqualTo(new BigNumber(-424))));
          _1197_v29 = _nw205;
          let _1198_v30;
          _1198_v30 = _dafny.Set.fromElements(_this.f12, (_this).f14);
          let _1199_v31;
          let _nw206 = new _module.C0();
          _nw206.__ctor((_dafny.Set.fromElements(_this.f12, (_this).f14)).IsProperSubsetOf(_1198_v30), new BigNumber(502), _dafny.MultiSet.fromElements((_this).f14, (_1197_v29).f17), !((_1197_v29).f17));
          _1199_v31 = _nw206;
          let _1200_v32;
          let _init26 = ((_1201_v31, _1202_v25) => function (_1203_i2) {
            return !_dafny.Seq.contains(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_1202_v25.f12,(_1201_v31).f17)), _dafny.Map.Empty.slice().updateUnsafe((_1201_v31).f17,(_1201_v31).f17));
          })(_1199_v31, _1193_v25);
          let _nw207 = Array((new BigNumber(6)).toNumber());
          for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw207.length); _i0_26++) {
            _nw207[_i0_26] = _init26(new BigNumber(_i0_26));
          }
          _1200_v32 = _nw207;
          let _index176 = _module.__default.safeIndex(new BigNumber(186), new BigNumber((_1200_v32).length));
          (_1200_v32)[_index176] = (_1199_v31).f17;
          let _index177 = _module.__default.safeIndex(new BigNumber(186), new BigNumber((_1200_v32).length));
          (_1200_v32)[_index177] = !(_module.__default.fm0(_1194_v26, globalState)) || (!((_1197_v29).f17));
        }
        let _index178 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length));
        (_1152_v0)[_index178] = (_module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f13), (_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))])).plus((_this).f13);
        let _1204_v33;
        _1204_v33 = _dafny.Seq.of(_1152_v0);
        let _1205_v34;
        _1205_v34 = new _dafny.CodePoint('l'.codePointAt(0));
        let _1206_v35;
        _1206_v35 = _module.D1.create_DC2((_this).f14, !((_this).f14), _this.f12, new _dafny.CodePoint('c'.codePointAt(0)), _1205_v34);
        let _source11 = ((_this.f12) ? (_module.D5.create_DC13(new BigNumber((_this.f11).cardinality()), _this.f12, new BigNumber((_1204_v33).length), _1166_v12, (_1206_v35).dtor_cf2)) : (_1168_v13));
        if (_source11.is_DC13) {
          let _1207___mcc_h6 = (_source11).cf30;
          let _1208___mcc_h7 = (_source11).cf31;
          let _1209___mcc_h8 = (_source11).cf32;
          let _1210___mcc_h9 = (_source11).cf33;
          let _1211___mcc_h10 = (_source11).cf34;
          let _1212_cf34 = _1211___mcc_h10;
          let _1213_cf33 = _1210___mcc_h9;
          let _1214_cf32 = _1209___mcc_h8;
          let _1215_cf31 = _1208___mcc_h7;
          let _1216_cf30 = _1207___mcc_h6;
          (globalState).f3 = _1212_cf34;
          let _1217_v36;
          let _init27 = ((_1218_cf34) => function (_1219_i3) {
            return _1218_cf34;
          })(_1212_cf34);
          let _nw208 = Array((new BigNumber(19)).toNumber());
          for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw208.length); _i0_27++) {
            _nw208[_i0_27] = _init27(new BigNumber(_i0_27));
          }
          _1217_v36 = _nw208;
          let _index179 = _module.__default.safeIndex(new BigNumber(763), new BigNumber((_1217_v36).length));
          (_1217_v36)[_index179] = !(_1212_cf34);
          let _1220_v38;
          _1220_v38 = _dafny.Set.fromElements(new _dafny.CodePoint('c'.codePointAt(0)), _1205_v34, _1205_v34);
          let _index180 = _module.__default.safeIndex(new BigNumber(763), new BigNumber((_1217_v36).length));
          (_1217_v36)[_index180] = (function () {
            let _coll26 = new _dafny.Set();
            for (const _compr_26 of (_module.__default.fm21(_1212_cf34, true, _1215_cf31, globalState)).Keys.Elements) {
              let _1221_v37 = _compr_26;
              if ((_module.__default.fm21(_1212_cf34, true, _1215_cf31, globalState)).contains(_1221_v37)) {
                _coll26.add(_1221_v37);
              }
            }
            return _coll26;
          }()).IsSubsetOf(_1220_v38);
          let _1222_v39;
          _1222_v39 = _dafny.Seq.of(_1217_v36, _1217_v36, _1217_v36, _1217_v36, _1217_v36);
          _1222_v39 = _1222_v39;
          let _1223_v40;
          _1223_v40 = _dafny.Map.Empty.slice().updateUnsafe(_1152_v0,(_this).f13);
          _1205_v34 = (((_1223_v40).equals(_1223_v40)) ? (_1205_v34) : (_1205_v34));
        } else if (_source11.is_DC14) {
          let _1224_v41;
          _1224_v41 = _dafny.Seq.of(_this.f12);
          let _index181 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length));
          let _index182 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_1153_v1).length));
          let _rhs196 = (_this).f13;
          let _rhs197 = _module.__default.safeDivisionInt(new BigNumber((_1224_v41).length), _module.__default.fm1(_this.f12, p0, new BigNumber((_dafny.Seq.UnicodeFromString("scarno")).length), globalState));
          let _rhs198 = _module.D4.create_DC10(_1152_v0);
          let _lhs156 = _1152_v0;
          let _lhs157 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length));
          let _lhs158 = globalState;
          let _lhs159 = _1153_v1;
          let _lhs160 = _module.__default.safeIndex(new BigNumber(525), new BigNumber((_1153_v1).length));
          _lhs156[_lhs157] = _rhs196;
          _lhs158.f9 = _rhs197;
          _lhs159[_lhs160] = _rhs198;
          let _1225_v42;
          _1225_v42 = _dafny.Seq.of(_1224_v41, _1224_v41, _dafny.Seq.update(_1224_v41, _module.__default.safeIndex(new BigNumber(815), new BigNumber((_1224_v41).length)), _this.f12));
          let _1226_v43;
          _1226_v43 = _dafny.Map.Empty.slice().updateUnsafe(false,_1225_v42);
          _1226_v43 = (_1226_v43).update((_this).f14, _dafny.Seq.Concat(_1225_v42, _1225_v42));
          let _1227_v44;
          let _nw209 = Array((new BigNumber(28)).toNumber()).fill(false);
          _1227_v44 = _nw209;
          let _index183 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_1227_v44).length));
          (_1227_v44)[_index183] = (_this).f14;
          let _index184 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_1227_v44).length));
          let _rhs199 = p0;
          let _rhs200 = _this.f12;
          let _rhs201 = _this.f12;
          let _lhs161 = globalState;
          let _lhs162 = _1227_v44;
          let _lhs163 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_1227_v44).length));
          let _lhs164 = globalState;
          _lhs161.f9 = _rhs199;
          _lhs162[_lhs163] = _rhs200;
          _lhs164.f8 = _rhs201;
          let _1228_v45;
          let _nw210 = Array((new BigNumber(15)).toNumber()).fill([]);
          _1228_v45 = _nw210;
          _1228_v45 = _1228_v45;
        } else {
          let _1229___mcc_h11 = (_source11).cf29;
          let _1230_cf29 = _1229___mcc_h11;
          let _1231_v46;
          _1231_v46 = _dafny.Seq.of(_this.f12);
          let _1232_v47;
          let _nw211 = new _module.C0();
          _nw211.__ctor(true, (_this).f13, _dafny.MultiSet.FromArray(_1231_v46), false);
          _1232_v47 = _nw211;
          let _1233_v48;
          _1233_v48 = _dafny.Seq.of(_1232_v47);
          let _1234_v49;
          _1234_v49 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.Concat(_1233_v48, _1233_v48));
          _1234_v49 = (_1234_v49).update(p0, _dafny.Seq.Concat(_1233_v48, _1233_v48));
          let _1235_v50;
          let _nw212 = new _module.C2();
          _nw212.__ctor(_1232_v47.f12, _this.f11, (_this).f14);
          _1235_v50 = _nw212;
          let _1236_v51;
          let _nw213 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Set.Empty);
          _1236_v51 = _nw213;
          let _1237_v52;
          _1237_v52 = _dafny.Seq.UnicodeFromString("tdx");
          let _1238_v53;
          _1238_v53 = _dafny.Set.fromElements(new BigNumber((_1237_v52).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(583)), ((_1239_v34) => function (_1240_i4) {
            return _1239_v34;
          })(_1205_v34))).length));
          let _index185 = _module.__default.safeIndex(new BigNumber(668), new BigNumber((_1236_v51).length));
          (_1236_v51)[_index185] = _1238_v53;
          let _index186 = _module.__default.safeIndex(new BigNumber(668), new BigNumber((_1236_v51).length));
          (_1236_v51)[_index186] = function () {
            let _coll27 = new _dafny.Set();
            for (const _compr_27 of _dafny.IntegerRange(new BigNumber(188), new BigNumber(25))) {
              let _1241_v54 = _compr_27;
              if (((new BigNumber(188)).isLessThanOrEqualTo(_1241_v54)) && ((_1241_v54).isLessThan(new BigNumber(25)))) {
                _coll27.add((_1241_v54).minus((_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))]));
              }
            }
            return _coll27;
          }();
          (globalState).f8 = ((_module.__default.fm0(_1205_v34, globalState)) ? (_1235_v50.f15) : (false));
        }
        (globalState).f7 = (_this).fm6(!(_this.f12), (_this).f13, globalState);
      }
      let _1242_i5;
      _1242_i5 = _dafny.ZERO;
      L8: {
        while (_this.f12) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1242_i5)) {
              break L8;
            }
            _1242_i5 = (_1242_i5).plus(_dafny.ONE);
            _module.__default.m0(globalState);
            let _index187 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length));
            (_1152_v0)[_index187] = (new BigNumber(-310)).multipliedBy(_module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f13), p0));
            let _1243_v55;
            _1243_v55 = new _dafny.CodePoint('p'.codePointAt(0));
            let _1244_v56;
            _1244_v56 = _dafny.Map.Empty.slice().updateUnsafe((_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))],_1243_v55);
            if (_module.__default.fm0((((_1244_v56).contains(p0)) ? ((_1244_v56).get(p0)) : (_1243_v55)), globalState)) {
              let _1245_v57;
              let _init28 = function (_1246_i6) {
                return ((_this).f13).isLessThan((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ad")).length))));
              };
              let _nw214 = Array((new BigNumber(25)).toNumber());
              for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw214.length); _i0_28++) {
                _nw214[_i0_28] = _init28(new BigNumber(_i0_28));
              }
              _1245_v57 = _nw214;
              let _1247_v58;
              _1247_v58 = _dafny.Seq.UnicodeFromString("xmboh");
              let _1248_v59;
              _1248_v59 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(134)), ((_1249_v55) => function (_1250_i7) {
                return _1249_v55;
              })(_1243_v55))).length));
              let _1251_v60;
              _1251_v60 = _module.D7.create_DC18(p0, _1247_v58, (_this).f14, (_this).f14, _1248_v59);
              let _1252_v61;
              _1252_v61 = _dafny.Seq.of(_this.f12);
              let _pat_let_tv25 = _1248_v59;
              let _nw215 = Array((new BigNumber(11)).toNumber());
              _nw215[(_dafny.ZERO).toNumber()] = !((_this).f14);
              _nw215[(_dafny.ONE).toNumber()] = (_this).f14;
              _nw215[(new BigNumber(2)).toNumber()] = (_this).f14;
              _nw215[(new BigNumber(3)).toNumber()] = false;
              _nw215[(new BigNumber(4)).toNumber()] = _this.f12;
              _nw215[(new BigNumber(5)).toNumber()] = !_dafny.Seq.contains(_1247_v58, _1243_v55);
              _nw215[(new BigNumber(6)).toNumber()] = (_this).f14;
              _nw215[(new BigNumber(7)).toNumber()] = !((function (_pat_let45_0) {
                return function (_1253_dt__update__tmp_h1) {
                  return function (_pat_let46_0) {
                    return function (_1254_dt__update_hcf44_h0) {
                      return _module.D7.create_DC18((_1253_dt__update__tmp_h1).dtor_cf40, (_1253_dt__update__tmp_h1).dtor_cf41, (_1253_dt__update__tmp_h1).dtor_cf42, (_1253_dt__update__tmp_h1).dtor_cf43, _1254_dt__update_hcf44_h0);
                    }(_pat_let46_0);
                  }(_pat_let_tv25);
                }(_pat_let45_0);
              }(_1251_v60)).dtor_cf43);
              _nw215[(new BigNumber(8)).toNumber()] = (_this).f14;
              _nw215[(new BigNumber(9)).toNumber()] = _this.f12;
              _nw215[(new BigNumber(10)).toNumber()] = !_dafny.areEqual(_1252_v61, _1252_v61);
              _1245_v57 = _nw215;
              let _1255_v62;
              let _nw216 = new _module.C2();
              _nw216.__ctor(_this.f12, (_dafny.MultiSet.fromElements((_this).f14)).update(_module.__default.fm0(new _dafny.CodePoint('c'.codePointAt(0)), globalState), _module.__default.abs((_this).f13)), (_this).f14);
              _1255_v62 = _nw216;
              let _1256_v63;
              _1256_v63 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1245_v57);
              (globalState).f8 = !((_1256_v63).equals(_1256_v63)) || (_1255_v62.f15);
              let _1257_v64;
              let _nw217 = new _module.C2();
              _nw217.__ctor(_this.f12, _this.f11, (_this).f14);
              _1257_v64 = _nw217;
              (globalState).f9 = (_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))];
            } else {
              let _1258_v65;
              _1258_v65 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,(_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))]);
              let _1259_v66;
              _1259_v66 = _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("tpnnpacy")).length));
              let _1260_v67;
              _1260_v67 = _dafny.Seq.of(_1259_v66, _1259_v66);
              let _1261_v68;
              _1261_v68 = _dafny.Map.Empty.slice().updateUnsafe(_1258_v65,_dafny.Seq.Concat((_1260_v67)[_module.__default.safeIndex((_1259_v66)[_module.__default.safeIndex(p0, new BigNumber((_1259_v66).length))], new BigNumber((_1260_v67).length))], _1259_v66));
              _1261_v68 = _1261_v68;
              let _1262_v69;
              let _nw218 = Array((new BigNumber(4)).toNumber()).fill([]);
              _1262_v69 = _nw218;
              let _index188 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_1262_v69).length));
              (_1262_v69)[_index188] = _1152_v0;
              let _index189 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_1262_v69).length));
              (_1262_v69)[_index189] = _1152_v0;
              let _1263_v70;
              let _nw219 = new _module.C2();
              _nw219.__ctor(_module.__default.fm0(_1243_v55, globalState), (_module.__default.fm22(p0, _this.f12, globalState)).update((_this).f14, _module.__default.abs((_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))])), (_this).f14);
              _1263_v70 = _nw219;
              let _1264_v71;
              _1264_v71 = _dafny.Map.Empty.slice().updateUnsafe((_this.f11).update(_1263_v70.f15, _module.__default.abs(p0)),(_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))]);
              let _1265_v72;
              _1265_v72 = _dafny.Map.Empty.slice().updateUnsafe(_1243_v55,(_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))]);
              let _1266_v74;
              let _nw220 = Array((new BigNumber(8)).toNumber());
              _nw220[(_dafny.ZERO).toNumber()] = (_1265_v72).update(_1243_v55, (_this).f13);
              _nw220[(_dafny.ONE).toNumber()] = _1265_v72;
              _nw220[(new BigNumber(2)).toNumber()] = _1265_v72;
              _nw220[(new BigNumber(3)).toNumber()] = _1265_v72;
              _nw220[(new BigNumber(4)).toNumber()] = function () {
                let _coll28 = new _dafny.Map();
                for (const _compr_28 of (_dafny.Set.fromElements(_1243_v55)).Elements) {
                  let _1267_v73 = _compr_28;
                  if ((_dafny.Set.fromElements(_1243_v55)).contains(_1267_v73)) {
                    _coll28.push([_1267_v73,(_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))]]);
                  }
                }
                return _coll28;
              }();
              _nw220[(new BigNumber(5)).toNumber()] = _1265_v72;
              _nw220[(new BigNumber(6)).toNumber()] = _1265_v72;
              _nw220[(new BigNumber(7)).toNumber()] = _1265_v72;
              _1266_v74 = _nw220;
              let _index190 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_1266_v74).length));
              (_1266_v74)[_index190] = _dafny.Map.Empty.slice().updateUnsafe(_1243_v55,(_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))]);
              let _1268_v76;
              _1268_v76 = _dafny.MultiSet.fromElements(_this.f11);
              let _1269_v78;
              _1269_v78 = _dafny.Seq.of(_this.f11, _this.f11);
              let _1270_v79;
              _1270_v79 = _dafny.Set.fromElements(_this.f11, (_1269_v78)[_module.__default.safeIndex(p0, new BigNumber((_1269_v78).length))]);
              let _index191 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length));
              let _index192 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length));
              let _index193 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_1266_v74).length));
              let _rhs202 = _1263_v70;
              let _rhs203 = (_module.__default.safeDivisionInt(_module.__default.fm1((_this).f14, (_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))], (_this).f13, globalState), (_dafny.ZERO).minus((_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))]))).multipliedBy(p0);
              let _rhs204 = ((!((_this).f13).isEqualTo(new BigNumber(411))) ? ((function () {
                let _coll29 = new _dafny.Map();
                for (const _compr_29 of (_1268_v76).Elements) {
                  let _1271_v75 = _compr_29;
                  if ((_1268_v76).contains(_1271_v75)) {
                    _coll29.push([_1271_v75,new BigNumber((_1258_v65).length)]);
                  }
                }
                return _coll29;
              }()).Merge(function () {
                let _coll30 = new _dafny.Map();
                for (const _compr_30 of (_1270_v79).Elements) {
                  let _1272_v77 = _compr_30;
                  if ((_1270_v79).contains(_1272_v77)) {
                    _coll30.push([_1272_v77,(_this).f13]);
                  }
                }
                return _coll30;
              }())) : ((_dafny.Map.Empty.slice().updateUnsafe(_this.f11,(_this).f13)).Merge(_1264_v71)));
              let _rhs205 = _module.__default.safeDivisionInt(new BigNumber(815), new BigNumber((_1258_v65).length));
              let _rhs206 = _1265_v72;
              let _lhs165 = _1152_v0;
              let _lhs166 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length));
              let _lhs167 = _1152_v0;
              let _lhs168 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length));
              let _lhs169 = _1266_v74;
              let _lhs170 = _module.__default.safeIndex(new BigNumber(354), new BigNumber((_1266_v74).length));
              _1263_v70 = _rhs202;
              _lhs165[_lhs166] = _rhs203;
              _1264_v71 = _rhs204;
              _lhs167[_lhs168] = _rhs205;
              _lhs169[_lhs170] = _rhs206;
              let _1273_v80;
              let _nw221 = new _module.C0();
              _nw221.__ctor(_this.f12, (_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))], _this.f11, _this.f12);
              _1273_v80 = _nw221;
              let _1274_v81;
              _1274_v81 = _dafny.Seq.UnicodeFromString("xcisubu");
              let _1275_v82;
              _1275_v82 = _dafny.Map.Empty.slice().updateUnsafe(_1273_v80,_1274_v81);
              _1275_v82 = (_1275_v82).update(_1273_v80, _1274_v81);
              (globalState).f3 = _module.__default.fm0(_1243_v55, globalState);
            }
            let _1276_v83;
            _1276_v83 = _dafny.Set.fromElements(p0);
            let _1277_v84;
            let _nw222 = Array((new BigNumber(2)).toNumber());
            _nw222[(_dafny.ZERO).toNumber()] = !(_this.f12) || ((_this).f14);
            _nw222[(_dafny.ONE).toNumber()] = (_1276_v83).IsSubsetOf(_1276_v83);
            _1277_v84 = _nw222;
            let _index194 = _module.__default.safeIndex(new BigNumber(705), new BigNumber((_1277_v84).length));
            (_1277_v84)[_index194] = ((_this.f12) ? (_module.__default.fm0(_1243_v55, globalState)) : ((_this).f14));
            let _1278_v85;
            _1278_v85 = _dafny.Seq.UnicodeFromString("hbi");
            let _index195 = _module.__default.safeIndex(new BigNumber(705), new BigNumber((_1277_v84).length));
            (_1277_v84)[_index195] = _dafny.areEqual((((_this).fm6((_this).f14, p0, globalState)) ? (_module.__default.fm9(p0, false, globalState)) : (_1278_v85)), _1278_v85);
          }
        }
      }
      let _1279_v86;
      _1279_v86 = _dafny.Set.fromElements(_this.f12, false, (_this).f14);
      let _1280_v87;
      _1280_v87 = _dafny.Seq.UnicodeFromString("kjje");
      let _1281_v88;
      _1281_v88 = _dafny.Set.fromElements((_this).f13);
      let _hi2 = (_module.D7.create_DC18(new BigNumber((_1279_v86).length), _1280_v87, false, (_this).f14, _1281_v88)).dtor_cf40;
      for (let _1282_i8 = (_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))]; _1282_i8.isLessThan(_hi2); _1282_i8 = _1282_i8.plus(_dafny.ONE)) {
        if ((_this).f14) {
          let _1283_v89;
          let _init29 = function (_1284_i9) {
            return (_this).f14;
          };
          let _nw223 = Array((new BigNumber(21)).toNumber());
          for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw223.length); _i0_29++) {
            _nw223[_i0_29] = _init29(new BigNumber(_i0_29));
          }
          _1283_v89 = _nw223;
          let _1285_v90;
          _1285_v90 = _dafny.Map.Empty.slice().updateUnsafe(_1282_i8,_1283_v89);
          _1285_v90 = ((_1285_v90).Merge(_1285_v90)).Merge(_1285_v90);
          let _1286_v91;
          _1286_v91 = _dafny.Seq.of(_1282_i8);
          (globalState).f9 = ((_this.f12) ? ((_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))]) : (((_1286_v91)[_module.__default.safeIndex(p0, new BigNumber((_1286_v91).length))]).minus((_this).f13)));
          _1285_v90 = _1285_v90;
          let _index196 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length));
          (_1152_v0)[_index196] = _module.__default.safeModuloInt((_this).f13, (_this).f13);
          let _1287_v92;
          _1287_v92 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,p0);
          let _1288_v93;
          _1288_v93 = _dafny.Seq.of(_1287_v92);
          let _1289_v94;
          _1289_v94 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,_1280_v87);
          let _1290_v95;
          _1290_v95 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1287_v92);
          let _1291_v96;
          _1291_v96 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,_this.f12);
          let _1292_v97;
          let _nw224 = Array((new BigNumber(19)).toNumber());
          _nw224[(_dafny.ZERO).toNumber()] = _1287_v92;
          _nw224[(_dafny.ONE).toNumber()] = _1287_v92;
          _nw224[(new BigNumber(2)).toNumber()] = ((_1288_v93)[_module.__default.safeIndex(new BigNumber((_1289_v94).length), new BigNumber((_1288_v93).length))]).Merge(_1287_v92);
          _nw224[(new BigNumber(3)).toNumber()] = _1287_v92;
          _nw224[(new BigNumber(4)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_this.f12,(_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))])).Merge(_1287_v92);
          _nw224[(new BigNumber(5)).toNumber()] = _1287_v92;
          _nw224[(new BigNumber(6)).toNumber()] = (_1287_v92).Merge(_1287_v92);
          _nw224[(new BigNumber(7)).toNumber()] = (_1287_v92).update(false, (_this).f13);
          _nw224[(new BigNumber(8)).toNumber()] = (_module.__default.fm23(p0, globalState)).Merge(_1287_v92);
          _nw224[(new BigNumber(9)).toNumber()] = _1287_v92;
          _nw224[(new BigNumber(10)).toNumber()] = _1287_v92;
          _nw224[(new BigNumber(11)).toNumber()] = (_module.__default.fm23((_this).f13, globalState)).update((_this).f14, (_this).f13);
          _nw224[(new BigNumber(12)).toNumber()] = (_1287_v92).Merge(_1287_v92);
          _nw224[(new BigNumber(13)).toNumber()] = _1287_v92;
          _nw224[(new BigNumber(14)).toNumber()] = (_1287_v92).Merge(_1287_v92);
          _nw224[(new BigNumber(15)).toNumber()] = (_1287_v92).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f14,p0));
          _nw224[(new BigNumber(16)).toNumber()] = (((_1290_v95).contains(new BigNumber((_1291_v96).length))) ? ((_1290_v95).get(new BigNumber((_1291_v96).length))) : (_dafny.Map.Empty.slice().updateUnsafe(false,_1282_i8)));
          _nw224[(new BigNumber(17)).toNumber()] = (_1287_v92).Merge(_1287_v92);
          _nw224[(new BigNumber(18)).toNumber()] = _1287_v92;
          _1292_v97 = _nw224;
          let _rhs207 = _1292_v97;
          let _rhs208 = _dafny.Seq.UnicodeFromString("l");
          _1292_v97 = _rhs207;
          _1280_v87 = _rhs208;
        } else {
          let _1293_v98;
          let _init30 = function (_1294_i10) {
            return true;
          };
          let _nw225 = Array((new BigNumber(2)).toNumber());
          for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw225.length); _i0_30++) {
            _nw225[_i0_30] = _init30(new BigNumber(_i0_30));
          }
          _1293_v98 = _nw225;
          let _index197 = _module.__default.safeIndex(new BigNumber(627), new BigNumber((_1293_v98).length));
          (_1293_v98)[_index197] = _this.f12;
          let _index198 = _module.__default.safeIndex(new BigNumber(627), new BigNumber((_1293_v98).length));
          (_1293_v98)[_index198] = (_this).f14;
          (globalState).f1 = _1152_v0;
          let _1295_v99;
          _1295_v99 = _dafny.Seq.of(_1279_v86, (_1279_v86).Difference(_1279_v86));
          let _1296_v100;
          _1296_v100 = new _dafny.CodePoint('q'.codePointAt(0));
          let _1297_v101;
          _1297_v101 = _dafny.Seq.of(_1280_v87);
          let _index199 = _module.__default.safeIndex(new BigNumber(627), new BigNumber((_1293_v98).length));
          let _rhs209 = new BigNumber(((_1295_v99)[_module.__default.safeIndex((_dafny.ZERO).minus(_module.__default.fm1(true, _1282_i8, (_dafny.ZERO).minus(_1282_i8), globalState)), new BigNumber((_1295_v99).length))]).length);
          let _rhs210 = (((_this).f14) ? (_this.f12) : ((_1293_v98)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_1293_v98).length))]));
          let _rhs211 = _dafny.Seq.update(_1280_v87, _module.__default.safeIndex(new BigNumber(((_this.f11).Union(_this.f11)).cardinality()), new BigNumber((_1280_v87).length)), _1296_v100);
          let _rhs212 = _dafny.Seq.contains((_1297_v101)[_module.__default.safeIndex((_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))], new BigNumber((_1297_v101).length))], new _dafny.CodePoint('d'.codePointAt(0)));
          let _rhs213 = !(_this.f12);
          let _lhs171 = globalState;
          let _lhs172 = _this;
          let _lhs173 = globalState;
          let _lhs174 = _1293_v98;
          let _lhs175 = _module.__default.safeIndex(new BigNumber(627), new BigNumber((_1293_v98).length));
          _lhs171.f9 = _rhs209;
          _lhs172.f12 = _rhs210;
          _1280_v87 = _rhs211;
          _lhs173.f8 = _rhs212;
          _lhs174[_lhs175] = _rhs213;
          _module.__default.m0(globalState);
          let _1298_v102;
          let _nw226 = Array((new BigNumber(10)).toNumber());
          _1298_v102 = _nw226;
          let _1299_v103;
          _1299_v103 = _dafny.Map.Empty.slice().updateUnsafe((_1293_v98)[_module.__default.safeIndex(new BigNumber(627), new BigNumber((_1293_v98).length))],_1298_v102);
          let _1300_v104;
          _1300_v104 = _module.D9.create_DC23(_1299_v103);
          let _pat_let_tv26 = _1299_v103;
          let _index200 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length));
          let _index201 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length));
          let _rhs214 = function (_pat_let47_0) {
            return function (_1301_dt__update__tmp_h2) {
              return function (_pat_let48_0) {
                return function (_1302_dt__update_hcf55_h0) {
                  return _module.D9.create_DC23(_1302_dt__update_hcf55_h0);
                }(_pat_let48_0);
              }(_pat_let_tv26);
            }(_pat_let47_0);
          }(_1300_v104);
          let _rhs215 = (_this).f13;
          let _rhs216 = p0;
          let _rhs217 = new BigNumber(-637);
          let _lhs176 = _1152_v0;
          let _lhs177 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length));
          let _lhs178 = globalState;
          let _lhs179 = _1152_v0;
          let _lhs180 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length));
          _1300_v104 = _rhs214;
          _lhs176[_lhs177] = _rhs215;
          _lhs178.f9 = _rhs216;
          _lhs179[_lhs180] = _rhs217;
        }
        let _index202 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length));
        (_1152_v0)[_index202] = (_this).f13;
        let _1303_v105;
        _1303_v105 = _dafny.Seq.of(_this.f12, _this.f12);
        let _1304_v106;
        _1304_v106 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f14);
        let _1305_v107;
        let _nw227 = Array((new BigNumber(8)).toNumber());
        _nw227[(_dafny.ZERO).toNumber()] = (_this).f14;
        _nw227[(_dafny.ONE).toNumber()] = !((((_this).f14) ? (true) : (_this.f12)));
        _nw227[(new BigNumber(2)).toNumber()] = !(((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_1303_v105, _module.__default.safeIndex(p0, new BigNumber((_1303_v105).length)), _this.f12)).length))).isEqualTo((_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))]));
        _nw227[(new BigNumber(3)).toNumber()] = (((_1304_v106).contains((_this).f14)) ? ((_1304_v106).get((_this).f14)) : (_this.f12));
        _nw227[(new BigNumber(4)).toNumber()] = _this.f12;
        _nw227[(new BigNumber(5)).toNumber()] = (_this).f14;
        _nw227[(new BigNumber(6)).toNumber()] = _this.f12;
        _nw227[(new BigNumber(7)).toNumber()] = _this.f12;
        _1305_v107 = _nw227;
        let _index203 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((_1305_v107).length));
        (_1305_v107)[_index203] = _dafny.areEqual(_1280_v87, _1280_v87);
        let _1306_v108;
        let _nw228 = Array((new BigNumber(9)).toNumber()).fill(_module.D8.Default());
        _1306_v108 = _nw228;
        let _1307_v109;
        _1307_v109 = _module.D8.create_DC19(_this.f11);
        let _1308_v110;
        _1308_v110 = _module.D8.create_DC22(_1307_v109);
        let _index204 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((_1306_v108).length));
        (_1306_v108)[_index204] = _1308_v110;
        let _index205 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((_1305_v107).length));
        let _index206 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length));
        let _index207 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((_1306_v108).length));
        let _rhs218 = !((_1281_v88).IsSubsetOf(_1281_v88));
        let _rhs219 = ((_this.f12) ? (((_this).f13).plus(new BigNumber(466))) : (_1282_i8));
        let _rhs220 = _1308_v110;
        let _rhs221 = (_this).f13;
        let _rhs222 = p0;
        let _lhs181 = _1305_v107;
        let _lhs182 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((_1305_v107).length));
        let _lhs183 = _1152_v0;
        let _lhs184 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length));
        let _lhs185 = _1306_v108;
        let _lhs186 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((_1306_v108).length));
        let _lhs187 = globalState;
        let _lhs188 = globalState;
        _lhs181[_lhs182] = _rhs218;
        _lhs183[_lhs184] = _rhs219;
        _lhs185[_lhs186] = _rhs220;
        _lhs187.f9 = _rhs221;
        _lhs188.f9 = _rhs222;
        let _1309_v111;
        _1309_v111 = _dafny.Map.Empty.slice().updateUnsafe(_1152_v0,_this.f12);
        _1309_v111 = (_1309_v111).update(_1152_v0, _dafny.areEqual(_1280_v87, _1280_v87));
      }
      let _1310_v112;
      _1310_v112 = new _dafny.CodePoint('x'.codePointAt(0));
      let _1311_v113;
      _1311_v113 = _dafny.Seq.of(!(_this.f12), _module.__default.fm0(_1310_v112, globalState));
      let _1312_v114;
      _1312_v114 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f12),new BigNumber((_1311_v113).length));
      let _1313_v115;
      _1313_v115 = _dafny.Seq.of(p0, (_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))], new BigNumber((_dafny.Seq.UnicodeFromString("tyfnjxj")).length), (_this).f13, (((_1312_v114).contains(_this.f12)) ? ((_1312_v114).get(_this.f12)) : ((_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))])));
      let _index208 = _module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length));
      (_1152_v0)[_index208] = (((_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))]).plus(new BigNumber((_1313_v115).length))).plus((_1152_v0)[_module.__default.safeIndex(new BigNumber(3), new BigNumber((_1152_v0).length))]);
      return;
    }
    get f13() {
      let _this = this;
      return _this._f13;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f11 = _dafny.MultiSet.Empty;
      this._f12 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
    set f11(value) {
      let _this = this;
      _this._f11 = value;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
    set f12(value) {
      let _this = this;
      _this._f12 = value;
    };
    __ctor(f11, f12) {
      let _this = this;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      return;
    }
    fm2(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !(new BigNumber(896)).isEqualTo(((_dafny.ZERO).minus(new BigNumber(-394))).multipliedBy(new BigNumber(46)));
    };
    fm3(p0, p1, p2, p3, globalState) {
      let _this = this;
      return new BigNumber(9);
    };
    fm4(p0, p1, globalState) {
      let _this = this;
      return (((_this.f12) ? (new BigNumber(-538)) : (new BigNumber(-626)))).isEqualTo(new BigNumber(((_dafny.MultiSet.fromElements(new BigNumber(908))).Union(_dafny.MultiSet.fromElements(new BigNumber(-317), new BigNumber(163)))).cardinality()));
    };
    m1(p0, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = [];
      let r2 = _dafny.ZERO;
      let _1314_v0;
      _1314_v0 = new BigNumber(-835);
      let _1315_v1;
      _1315_v1 = _dafny.Seq.UnicodeFromString("fahubdda");
      let _rhs223 = (_dafny.ZERO).minus((_module.__default.fm1(_this.f12, _1314_v0, _1314_v0, globalState)).multipliedBy(_1314_v0));
      let _rhs224 = _this.f12;
      let _rhs225 = (_this.f11).update(_this.f12, _module.__default.abs((_dafny.ZERO).minus((new BigNumber(36)).plus(_1314_v0))));
      let _rhs226 = _dafny.Seq.Concat(_module.__default.fm5(_this.f12, _this.f12, _this.f12, _1314_v0, globalState), _module.__default.fm5(_this.f12, _this.f12, _this.f12, new BigNumber((_1315_v1).length), globalState));
      let _lhs189 = globalState;
      let _lhs190 = globalState;
      let _lhs191 = _this;
      let _lhs192 = globalState;
      _lhs189.f9 = _rhs223;
      _lhs190.f7 = _rhs224;
      _lhs191.f11 = _rhs225;
      _lhs192.f6 = _rhs226;
      let _1316_v2;
      let _nw229 = Array((new BigNumber(3)).toNumber()).fill(false);
      _1316_v2 = _nw229;
      let _1317_v3;
      _1317_v3 = _dafny.Set.fromElements(_1316_v2, _1316_v2, _1316_v2, _1316_v2);
      let _1318_v4;
      _1318_v4 = _dafny.Seq.of(_1314_v0, _module.__default.fm1(_this.f12, _1314_v0, _1314_v0, globalState), _1314_v0, (new BigNumber((_dafny.Seq.UnicodeFromString("kd")).length)).plus(_1314_v0));
      let _1319_v5;
      _1319_v5 = _dafny.Seq.of(p0, p0);
      let _1320_v6;
      _1320_v6 = _module.D1.create_DC1(p0);
      let _1321_v7;
      let _nw230 = Array((new BigNumber(29)).toNumber());
      _nw230[(_dafny.ZERO).toNumber()] = p0;
      _nw230[(_dafny.ONE).toNumber()] = p0;
      _nw230[(new BigNumber(2)).toNumber()] = (p0).Intersect(p0);
      _nw230[(new BigNumber(3)).toNumber()] = (p0).Union(p0);
      _nw230[(new BigNumber(4)).toNumber()] = _dafny.Set.fromElements(_this.f12, (_this.f12), _this.f12, !(_this.f12));
      _nw230[(new BigNumber(5)).toNumber()] = (p0).Union(p0);
      _nw230[(new BigNumber(6)).toNumber()] = p0;
      _nw230[(new BigNumber(7)).toNumber()] = (p0).Union(p0);
      _nw230[(new BigNumber(8)).toNumber()] = p0;
      _nw230[(new BigNumber(9)).toNumber()] = (p0).Difference((_1319_v5)[_module.__default.safeIndex(_1314_v0, new BigNumber((_1319_v5).length))]);
      _nw230[(new BigNumber(10)).toNumber()] = p0;
      _nw230[(new BigNumber(11)).toNumber()] = (_1320_v6).dtor_cf1;
      _nw230[(new BigNumber(12)).toNumber()] = p0;
      _nw230[(new BigNumber(13)).toNumber()] = p0;
      _nw230[(new BigNumber(14)).toNumber()] = p0;
      _nw230[(new BigNumber(15)).toNumber()] = p0;
      _nw230[(new BigNumber(16)).toNumber()] = p0;
      _nw230[(new BigNumber(17)).toNumber()] = (_1319_v5)[_module.__default.safeIndex(new BigNumber(-249), new BigNumber((_1319_v5).length))];
      _nw230[(new BigNumber(18)).toNumber()] = p0;
      _nw230[(new BigNumber(19)).toNumber()] = (p0).Intersect(p0);
      _nw230[(new BigNumber(20)).toNumber()] = p0;
      _nw230[(new BigNumber(21)).toNumber()] = (p0).Difference(_dafny.Set.fromElements(_this.f12, _this.f12));
      _nw230[(new BigNumber(22)).toNumber()] = ((_this.f12) ? (p0) : (p0));
      _nw230[(new BigNumber(23)).toNumber()] = (_dafny.Set.fromElements(!(false))).Union(p0);
      _nw230[(new BigNumber(24)).toNumber()] = (p0).Intersect(_dafny.Set.fromElements(_this.f12));
      _nw230[(new BigNumber(25)).toNumber()] = p0;
      _nw230[(new BigNumber(26)).toNumber()] = _dafny.Set.fromElements(_this.f12, _this.f12);
      _nw230[(new BigNumber(27)).toNumber()] = (p0).Difference(p0);
      _nw230[(new BigNumber(28)).toNumber()] = p0;
      _1321_v7 = _nw230;
      let _index209 = _module.__default.safeIndex(new BigNumber(185), new BigNumber((_1321_v7).length));
      (_1321_v7)[_index209] = (p0).Difference(p0);
      let _1322_v8;
      _1322_v8 = _dafny.Seq.of(_1317_v3);
      let _index210 = _module.__default.safeIndex(new BigNumber(185), new BigNumber((_1321_v7).length));
      let _rhs227 = (_1322_v8)[_module.__default.safeIndex(_module.__default.safeModuloInt(new BigNumber(662), (_dafny.ZERO).minus(_1314_v0)), new BigNumber((_1322_v8).length))];
      let _rhs228 = _1318_v4;
      let _rhs229 = p0;
      let _rhs230 = _this.f12;
      let _rhs231 = new BigNumber(922);
      let _lhs193 = _1321_v7;
      let _lhs194 = _module.__default.safeIndex(new BigNumber(185), new BigNumber((_1321_v7).length));
      let _lhs195 = globalState;
      _1317_v3 = _rhs227;
      _1318_v4 = _rhs228;
      _lhs193[_lhs194] = _rhs229;
      _lhs195.f3 = _rhs230;
      _1314_v0 = _rhs231;
      _1316_v2 = _1316_v2;
      let _1323_i0;
      _1323_i0 = _dafny.ZERO;
      L9: {
        while (_this.f12) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1323_i0)) {
              break L9;
            }
            _1323_i0 = (_1323_i0).plus(_dafny.ONE);
            (globalState).f9 = (_1314_v0).minus(_1314_v0);
            (globalState).f3 = !(_this.f12);
            if (_this.f12) {
              let _1324_v9;
              _1324_v9 = _module.D1.create_DC1((_1321_v7)[_module.__default.safeIndex(new BigNumber(185), new BigNumber((_1321_v7).length))]);
              let _1325_v10;
              _1325_v10 = _dafny.MultiSet.fromElements(_module.D1.create_DC3(_1324_v9));
              let _1326_v11;
              _1326_v11 = _module.D1.create_DC3(_1324_v9);
              (globalState).f3 = ((_1325_v10).Difference(_1325_v10)).IsProperSubsetOf((_1325_v10).update(_1326_v11, _module.__default.abs(_1314_v0)));
              let _1327_v12;
              let _nw231 = new _module.C3();
              _nw231.__ctor((_dafny.ZERO).minus(((_this.f12) ? (_1314_v0) : (_1314_v0))), _this.f12, (_dafny.MultiSet.fromElements(_this.f12, _this.f12)).Difference(_dafny.MultiSet.fromElements(!(_this.f12), _this.f12)), (_this.f11).IsDisjointFrom(_this.f11));
              _1327_v12 = _nw231;
              let _1328_v13;
              _1328_v13 = _dafny.Map.Empty.slice().updateUnsafe((_1327_v12).f13,(_dafny.ZERO).minus((_dafny.ZERO).minus((_1327_v12).f13)));
              let _1329_v14;
              _1329_v14 = _dafny.Set.fromElements((_1327_v12).f13, (_1327_v12).f13, (((_1328_v13).contains(new BigNumber(-708))) ? ((_1328_v13).get(new BigNumber(-708))) : ((_1327_v12).f13)));
              let _rhs232 = _1329_v14;
              let _rhs233 = _1314_v0;
              let _rhs234 = _1314_v0;
              let _rhs235 = (((_1327_v12).f14) ? (_1318_v4) : (_dafny.Seq.Concat(_1318_v4, _1318_v4)));
              let _lhs196 = globalState;
              let _lhs197 = globalState;
              _1329_v14 = _rhs232;
              _lhs196.f9 = _rhs233;
              _lhs197.f9 = _rhs234;
              _1318_v4 = _rhs235;
              let _1330_v15;
              let _nw232 = Array((new BigNumber(19)).toNumber()).fill([]);
              _1330_v15 = _nw232;
              let _index211 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_1330_v15).length));
              (_1330_v15)[_index211] = ((_this.f12) ? (_1316_v2) : (_1316_v2));
              let _index212 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_1330_v15).length));
              let _rhs236 = (_1314_v0).minus((((_1327_v12).f14) ? ((_1327_v12).f13) : (_1314_v0)));
              let _rhs237 = _1316_v2;
              let _lhs198 = _1330_v15;
              let _lhs199 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_1330_v15).length));
              _1314_v0 = _rhs236;
              _lhs198[_lhs199] = _rhs237;
              let _1331_v16;
              _1331_v16 = _dafny.Map.Empty.slice().updateUnsafe((_1318_v4)[_module.__default.safeIndex((_1327_v12).f13, new BigNumber((_1318_v4).length))],_1315_v1);
              let _rhs238 = (_1327_v12).f13;
              let _rhs239 = (_1327_v12).f14;
              let _rhs240 = (_1327_v12).f14;
              let _rhs241 = (((_1331_v16).contains(_1314_v0)) ? ((_1331_v16).get(_1314_v0)) : (_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("tghqo"), _1315_v1)));
              let _lhs200 = globalState;
              let _lhs201 = globalState;
              let _lhs202 = globalState;
              _lhs200.f9 = _rhs238;
              _lhs201.f8 = _rhs239;
              _lhs202.f8 = _rhs240;
              _1315_v1 = _rhs241;
            } else {
              let _1332_v17;
              _1332_v17 = _module.D9.create_DC25(_1315_v1, _this.f12, _1314_v0);
              let _1333_v18;
              _1333_v18 = _dafny.Seq.of(_1332_v17, _1332_v17);
              let _1334_v19;
              let _nw233 = new _module.C3();
              _nw233.__ctor(_1314_v0, _this.f12, _this.f11, _this.f12);
              _1334_v19 = _nw233;
              let _1335_v20;
              _1335_v20 = _dafny.Seq.of(_1334_v19, _1334_v19);
              let _1336_v21;
              _1336_v21 = _dafny.Map.Empty.slice().updateUnsafe(!((_1334_v19).f14),_this.f12);
              let _1337_v22;
              _1337_v22 = _dafny.Seq.of(true, _this.f12, (_1334_v19).f14);
              let _1338_v23;
              _1338_v23 = new _dafny.CodePoint('v'.codePointAt(0));
              let _1339_v24;
              _1339_v24 = _module.D1.create_DC2((_1334_v19).f14, _this.f12, false, _1338_v23, _1338_v23);
              let _1340_v25;
              let _nw234 = new _module.C1();
              _nw234.__ctor(new BigNumber((_1337_v22).length), _dafny.MultiSet.fromElements((_1339_v24).dtor_cf4), !(true));
              _1340_v25 = _nw234;
              let _1341_v26;
              _1341_v26 = _dafny.Set.fromElements(_1314_v0);
              let _1342_v27;
              let _init31 = ((_1343_v1) => function (_1344_i1) {
                return _1343_v1;
              })(_1315_v1);
              let _nw235 = Array((new BigNumber(7)).toNumber());
              for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw235.length); _i0_31++) {
                _nw235[_i0_31] = _init31(new BigNumber(_i0_31));
              }
              _1342_v27 = _nw235;
              let _1345_v28;
              _1345_v28 = _module.D5.create_DC13((_1334_v19).f13, true, (_1334_v19).f13, _1342_v27, _1340_v25.f12);
              let _1346_v29;
              _1346_v29 = _dafny.Map.Empty.slice().updateUnsafe((_1345_v28).dtor_cf34,(_1334_v19).f13);
              let _1347_v30;
              let _nw236 = new _module.C0();
              _nw236.__ctor(true, _1314_v0, _this.f11, _1340_v25.f12);
              _1347_v30 = _nw236;
              let _1348_v31;
              _1348_v31 = _dafny.Seq.of(_1347_v30);
              let _1349_v32;
              let _nw237 = Array((new BigNumber(19)).toNumber());
              _nw237[(_dafny.ZERO).toNumber()] = _1314_v0;
              _nw237[(_dafny.ONE).toNumber()] = _1314_v0;
              _nw237[(new BigNumber(2)).toNumber()] = ((((_this.f11).contains(_this.f12)) ? ((_this.f11).get(_this.f12)) : (_1314_v0))).plus(_1314_v0);
              _nw237[(new BigNumber(3)).toNumber()] = new BigNumber((_1333_v18).length);
              _nw237[(new BigNumber(4)).toNumber()] = _1314_v0;
              _nw237[(new BigNumber(5)).toNumber()] = _1314_v0;
              _nw237[(new BigNumber(6)).toNumber()] = _1314_v0;
              _nw237[(new BigNumber(7)).toNumber()] = _1314_v0;
              _nw237[(new BigNumber(8)).toNumber()] = new BigNumber((_1335_v20).length);
              _nw237[(new BigNumber(9)).toNumber()] = new BigNumber((_1336_v21).length);
              _nw237[(new BigNumber(10)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1334_v19).f13,_1340_v25)).length);
              _nw237[(new BigNumber(11)).toNumber()] = _module.__default.safeDivisionInt((_1334_v19).f13, (_1334_v19).f13);
              _nw237[(new BigNumber(12)).toNumber()] = new BigNumber(((_1341_v26).Difference(_1341_v26)).length);
              _nw237[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus((_1334_v19).f13);
              _nw237[(new BigNumber(14)).toNumber()] = (((_1346_v29).contains(_1340_v25.f12)) ? ((_1346_v29).get(_1340_v25.f12)) : (new BigNumber((_1348_v31).length)));
              _nw237[(new BigNumber(15)).toNumber()] = (_1347_v30).f18;
              _nw237[(new BigNumber(16)).toNumber()] = new BigNumber(278);
              _nw237[(new BigNumber(17)).toNumber()] = (_1334_v19).f13;
              _nw237[(new BigNumber(18)).toNumber()] = (_1334_v19).f13;
              _1349_v32 = _nw237;
              let _index213 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_1349_v32).length));
              (_1349_v32)[_index213] = _module.__default.safeDivisionInt(_1314_v0, _1314_v0);
              let _index214 = _module.__default.safeIndex(new BigNumber(355), new BigNumber((_1349_v32).length));
              (_1349_v32)[_index214] = _1314_v0;
              let _1350_v33;
              let _nw238 = new _module.C0();
              _nw238.__ctor((_1334_v19).f14, _module.__default.safeModuloInt(new BigNumber((p0).length), (_1334_v19).f13), _this.f11, !(new BigNumber((_dafny.MultiSet.fromElements(_1318_v4)).cardinality())).isEqualTo(_1314_v0));
              _1350_v33 = _nw238;
              let _index215 = _module.__default.safeIndex(new BigNumber(162), new BigNumber((_1316_v2).length));
              (_1316_v2)[_index215] = (_1347_v30).f17;
              let _index216 = _module.__default.safeIndex(new BigNumber(162), new BigNumber((_1316_v2).length));
              (_1316_v2)[_index216] = (_1347_v30).f17;
              (globalState).f8 = ((((_1334_v19).f13).isEqualTo(new BigNumber(591))) ? (_1340_v25.f12) : (_this.f12));
              let _1351_v34;
              _1351_v34 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(((_1349_v32)[_module.__default.safeIndex(new BigNumber(355), new BigNumber((_1349_v32).length))]).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(792)), ((_1352_v23) => function (_1353_i2) {
                return _1352_v23;
              })(_1338_v23))).length))),_dafny.Seq.Create(_module.__default.abs(new BigNumber(336)), ((_1354_v19) => function (_1355_i3) {
                return (_1354_v19).f13;
              })(_1334_v19)));
              _1351_v34 = ((((!((_1316_v2)[_module.__default.safeIndex(new BigNumber(162), new BigNumber((_1316_v2).length))])) ? ((_1350_v33).f17) : ((_1350_v33).f17))) ? (_1351_v34) : ((_dafny.Map.Empty.slice().updateUnsafe((_1350_v33).f18,_1318_v4)).Merge((_dafny.Map.Empty.slice().updateUnsafe((_1349_v32)[_module.__default.safeIndex(new BigNumber(355), new BigNumber((_1349_v32).length))],_1318_v4)).update(new BigNumber(907), _1318_v4))));
            }
            _1315_v1 = _module.__default.fm9(new BigNumber((_dafny.MultiSet.fromElements(_1314_v0)).cardinality()), _this.f12, globalState);
          }
        }
      }
      let _1356_v35;
      _1356_v35 = new _dafny.CodePoint('m'.codePointAt(0));
      let _1357_v36;
      let _nw239 = Array((new BigNumber(5)).toNumber());
      _nw239[(_dafny.ZERO).toNumber()] = _1356_v35;
      _nw239[(_dafny.ONE).toNumber()] = _1356_v35;
      _nw239[(new BigNumber(2)).toNumber()] = _1356_v35;
      _nw239[(new BigNumber(3)).toNumber()] = _1356_v35;
      _nw239[(new BigNumber(4)).toNumber()] = _1356_v35;
      _1357_v36 = _nw239;
      let _1358_v37;
      _1358_v37 = _module.D1.create_DC2(_this.f12, _this.f12, _this.f12, _1356_v35, _1356_v35);
      let _1359_v38;
      let _nw240 = Array((new BigNumber(19)).toNumber()).fill([]);
      _1359_v38 = _nw240;
      let _1360_v39;
      let _nw241 = new _module.C3();
      _nw241.__ctor(_1314_v0, _this.f12, _this.f11, _this.f12);
      _1360_v39 = _nw241;
      let _1361_v40;
      _1361_v40 = _dafny.Set.fromElements(_1360_v39, _1360_v39);
      let _1362_v41;
      _1362_v41 = _module.D8.create_DC21(true, _1359_v38, new BigNumber((_1315_v1).length), _1316_v2, new BigNumber((_1361_v40).length));
      let _index217 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_1357_v36).length));
      (_1357_v36)[_index217] = _module.__default.fm14(_1358_v37, (_dafny.ZERO).minus((_1362_v41).dtor_cf53), globalState);
      let _index218 = _module.__default.safeIndex(new BigNumber(185), new BigNumber((_1321_v7).length));
      let _index219 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_1357_v36).length));
      let _rhs242 = new BigNumber((_1318_v4).length);
      let _rhs243 = (p0).Intersect(_dafny.Set.fromElements(true));
      let _rhs244 = _1356_v35;
      let _rhs245 = _module.__default.fm24(globalState);
      let _lhs203 = globalState;
      let _lhs204 = _1321_v7;
      let _lhs205 = _module.__default.safeIndex(new BigNumber(185), new BigNumber((_1321_v7).length));
      let _lhs206 = _1357_v36;
      let _lhs207 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_1357_v36).length));
      _lhs203.f9 = _rhs242;
      _lhs204[_lhs205] = _rhs243;
      _lhs206[_lhs207] = _rhs244;
      _1318_v4 = _rhs245;
      let _hi3 = _1314_v0;
      for (let _1363_i4 = _module.__default.safeDivisionInt(_1314_v0, _1314_v0); _1363_i4.isLessThan(_hi3); _1363_i4 = _1363_i4.plus(_dafny.ONE)) {
        let _index220 = _module.__default.safeIndex(new BigNumber(185), new BigNumber((_1321_v7).length));
        (_1321_v7)[_index220] = p0;
        let _1364_v42;
        let _nw242 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _1364_v42 = _nw242;
        let _index221 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_1364_v42).length));
        (_1364_v42)[_index221] = _1315_v1;
        let _index222 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_1364_v42).length));
        (_1364_v42)[_index222] = _1315_v1;
        let _1365_v43;
        _1365_v43 = _dafny.Seq.of(_this.f12);
        let _1366_v44;
        _1366_v44 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_dafny.Seq.of(_1360_v39.f12, _this.f12)),_1365_v43);
        let _1367_v45;
        _1367_v45 = _dafny.Map.Empty.slice().updateUnsafe((_1357_v36)[_module.__default.safeIndex(new BigNumber(164), new BigNumber((_1357_v36).length))],_1360_v39.f12);
        let _rhs246 = (((_1366_v44).contains(_1360_v39.f11)) ? ((_1366_v44).get(_1360_v39.f11)) : (_1365_v43));
        let _rhs247 = _1314_v0;
        let _rhs248 = ((_1367_v45).Merge(_1367_v45)).contains((_1357_v36)[_module.__default.safeIndex(new BigNumber(164), new BigNumber((_1357_v36).length))]);
        let _lhs208 = globalState;
        let _lhs209 = globalState;
        _lhs208.f6 = _rhs246;
        _1314_v0 = _rhs247;
        _lhs209.f7 = _rhs248;
        let _1368_v46;
        _1368_v46 = _dafny.Map.Empty.slice().updateUnsafe(_1360_v39,_1360_v39);
        _1368_v46 = _dafny.Map.Empty.slice().updateUnsafe(_1360_v39,_1360_v39);
      }
      let _1369_v47;
      let _nw243 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Map.Empty);
      _1369_v47 = _nw243;
      r0 = _1369_v47;
      let _1370_v48;
      let _nw244 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.of());
      _1370_v48 = _nw244;
      r1 = _1370_v48;
      r2 = _1314_v0;
      return [r0, r1, r2];
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
