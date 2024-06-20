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
    static fm4(p0, p1, globalState) {
      if ((new BigNumber(104)).isEqualTo(new BigNumber((_dafny.Set.fromElements(new BigNumber(-395))).length))) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('p'.codePointAt(0));
      }
    };
    static fm5(p0, p1, p2, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(33)), function (_0_i0) {
        return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(412)), function (_1_i1) {
          return _dafny.Map.Empty.slice().updateUnsafe(true,false);
        }), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,false), _dafny.Map.Empty.slice().updateUnsafe(false,false), _dafny.Map.Empty.slice().updateUnsafe(true,!(false))))).length);
      });
    };
    static fm6(globalState) {
      return _module.__default.safeModuloInt((new BigNumber((_dafny.Seq.UnicodeFromString("nasbdl")).length)).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)))).length))), new BigNumber(((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(-139), new BigNumber(404))) {
          let _2_v0 = _compr_0;
          if (((new BigNumber(-139)).isLessThanOrEqualTo(_2_v0)) && ((_2_v0).isLessThan(new BigNumber(404)))) {
            _coll0.push([_module.__default.safeModuloInt(_2_v0, new BigNumber(153)),function () {
              let _coll1 = new _dafny.Map();
              for (const _compr_1 of (_dafny.Seq.of(new BigNumber(741))).Elements) {
                let _3_v1 = _compr_1;
                if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(741)), _3_v1)) {
                  _coll1.push([_module.__default.safeDivisionInt(_3_v1, new BigNumber(293)),!(true)]);
                }
              }
              return _coll1;
            }()]);
          }
        }
        return _coll0;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("u")).length)))).length),function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(22),new BigNumber(203))).Keys.Elements) {
          let _4_v2 = _compr_2;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(22),new BigNumber(203))).contains(_4_v2)) {
            _coll2.push([(_4_v2).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length)),false]);
          }
        }
        return _coll2;
      }()))).length));
    };
    static fm9(p0, p1, globalState) {
      return function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of (_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-545)), function (_5_i0) {
          return _dafny.Seq.of(true, false, false);
        }), _dafny.Seq.of(_dafny.Seq.of(true, true, false)))).Elements) {
          let _6_v0 = _compr_3;
          if (_dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-545)), function (_5_i0) {
            return _dafny.Seq.of(true, false, false);
          }), _dafny.Seq.of(_dafny.Seq.of(true, true, false))), _6_v0)) {
            _coll3.push([_6_v0,!((!(false)) === (true))]);
          }
        }
        return _coll3;
      }();
    };
    static fm10(globalState) {
      return _module.D1.create_DC3(_dafny.Set.fromElements(_dafny.Set.fromElements(new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('e'.codePointAt(0)))));
    };
    static fm11(p0, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(555)), function (_7_i0) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      });
    };
    static fm12(p0, p1, p2, p3, globalState) {
      return _dafny.areEqual(_dafny.Seq.of(new _dafny.CodePoint('a'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0))), _dafny.Seq.of(new _dafny.CodePoint('k'.codePointAt(0))));
    };
    static fm13(p0, p1, p2, p3, globalState) {
      return (_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("jlwcc"))).Difference((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("dulq"), _dafny.Seq.UnicodeFromString("ouob"), _dafny.Seq.UnicodeFromString("wntsnxr"))).Difference(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("oikjsmub"))));
    };
    static fm14(p0, p1, globalState) {
      return function () {
        let _coll4 = new _dafny.Set();
        for (const _compr_4 of _dafny.IntegerRange(new BigNumber(-141), new BigNumber(493))) {
          let _8_v0 = _compr_4;
          if (((new BigNumber(-141)).isLessThanOrEqualTo(_8_v0)) && ((_8_v0).isLessThan(new BigNumber(493)))) {
            _coll4.add((_8_v0).minus(new BigNumber((_dafny.Seq.of(_dafny.Seq.of(false, true), _dafny.Seq.of(true))).length)));
          }
        }
        return _coll4;
      }();
    };
    static fm15(globalState) {
      return (_dafny.Set.fromElements(new _dafny.CodePoint('a'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0)), new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0)))).Union((_dafny.Set.fromElements(new _dafny.CodePoint('j'.codePointAt(0)))).Difference(function () {
        let _coll5 = new _dafny.Set();
        for (const _compr_5 of (_dafny.Set.fromElements(new _dafny.CodePoint('r'.codePointAt(0)))).Elements) {
          let _9_v0 = _compr_5;
          if ((_dafny.Set.fromElements(new _dafny.CodePoint('r'.codePointAt(0)))).contains(_9_v0)) {
            _coll5.add(_9_v0);
          }
        }
        return _coll5;
      }()));
    };
    static fm16(p0, p1, p2, globalState) {
      return _dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(!(false))).length), new BigNumber((_dafny.MultiSet.fromElements(_dafny.Set.fromElements(!(false)))).cardinality())));
    };
    static fm17(p0, p1, globalState) {
      return (_dafny.MultiSet.fromElements(!((_module.D9.create_DC20(_dafny.Seq.UnicodeFromString("lt"), false, new BigNumber(130))).dtor_cf38), false, true, false)).Difference((_dafny.MultiSet.fromElements(!(true))).Union((_module.D3.create_DC7(_dafny.MultiSet.FromArray(_dafny.Seq.of(!(true), true)))).dtor_cf12));
    };
    static fm18(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("rkakqfc"), _dafny.Seq.UnicodeFromString("uhejkofsg")),new BigNumber((((false) ? (_dafny.Seq.UnicodeFromString("olxnmwt")) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(461)), function (_10_i0) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      })))).length));
    };
    static fm19(globalState) {
      return function () {
        let _coll6 = new _dafny.Map();
        for (const _compr_6 of ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber(-997)),_module.D3.create_DC7(_dafny.MultiSet.fromElements(!(true), true)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber(-629)),_module.D3.create_DC7(_dafny.MultiSet.fromElements(true))))).Keys.Elements) {
          let _11_v0 = _compr_6;
          if (((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber(-997)),_module.D3.create_DC7(_dafny.MultiSet.fromElements(!(true), true)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber(-629)),_module.D3.create_DC7(_dafny.MultiSet.fromElements(true))))).contains(_11_v0)) {
            _coll6.push([_11_v0,((_dafny.ZERO).minus(new BigNumber(-516))).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.UnicodeFromString("bsjog")).length))]);
          }
        }
        return _coll6;
      }();
    };
    static m0(p0, p1, globalState) {
      let r0 = false;
      let _12_v0;
      _12_v0 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-251),p1);
      let _13_v1;
      _13_v1 = new BigNumber(-477);
      (globalState).f3 = (new BigNumber((_12_v0).length)).minus(_13_v1);
      let _14_v2;
      _14_v2 = _dafny.Seq.UnicodeFromString("wqmiaea");
      let _15_v3;
      _15_v3 = _dafny.Map.Empty.slice().updateUnsafe(_14_v2,_13_v1);
      let _16_v4;
      _16_v4 = _dafny.Seq.of(new BigNumber(382), _13_v1);
      _15_v3 = ((_dafny.Map.Empty.slice().updateUnsafe(_14_v2,_13_v1)).Merge(_15_v3)).Merge((_module.__default.fm18(new BigNumber(878), _16_v4, _13_v1, _13_v1, globalState)).Merge(_15_v3));
      let _17_v5;
      _17_v5 = _dafny.MultiSet.fromElements((_16_v4)[_module.__default.safeIndex(new BigNumber(-426), new BigNumber((_16_v4).length))]);
      let _18_v6;
      _18_v6 = new _dafny.CodePoint('t'.codePointAt(0));
      let _hi0 = ((p1) ? (new BigNumber((_14_v2).length)) : (_13_v1));
      for (let _19_i0 = (_dafny.ZERO).minus((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_17_v5).cardinality()),_18_v6)).length)).multipliedBy(new BigNumber(582))); _19_i0.isLessThan(_hi0); _19_i0 = _19_i0.plus(_dafny.ONE)) {
        if (p1) {
          let _index0 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((p0).length));
          (p0)[_index0] = !((p1) === (p1));
          let _index1 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((p0).length));
          (p0)[_index1] = (p1) || (!(!(new BigNumber((_module.__default.fm19(globalState)).length)).isEqualTo(_19_i0)));
          let _20_v7;
          _20_v7 = _dafny.Set.fromElements(_14_v2, _14_v2, _14_v2, _14_v2);
          let _21_v8;
          let _nw0 = new _module.C0();
          _nw0.__ctor(_20_v7);
          _21_v8 = _nw0;
          let _22_v9;
          _22_v9 = _dafny.Seq.of(_21_v8);
          let _23_v10;
          _23_v10 = _dafny.MultiSet.fromElements(_21_v8, _21_v8);
          let _24_v11;
          let _nw1 = new _module.C3();
          _nw1.__ctor((p0)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p0).length))], p1, _14_v2, p1);
          _24_v11 = _nw1;
          let _25_v12;
          _25_v12 = _module.D7.create_DC16(p1, false, _24_v11, (_24_v11).f27);
          let _26_v13;
          _26_v13 = _dafny.Map.Empty.slice().updateUnsafe(_25_v12,_18_v6);
          let _27_v14;
          _27_v14 = _dafny.MultiSet.fromElements((_24_v11).f27);
          let _28_v15;
          let _nw2 = Array((new BigNumber(17)).toNumber()).fill(false);
          _28_v15 = _nw2;
          let _29_v16;
          _29_v16 = _dafny.Map.Empty.slice().updateUnsafe((_16_v4)[_module.__default.safeIndex(_13_v1, new BigNumber((_16_v4).length))],_28_v15);
          let _30_v17;
          _30_v17 = _29_v16;
          let _31_v18;
          _31_v18 = _dafny.Map.Empty.slice().updateUnsafe(_13_v1,_13_v1);
          let _32_v19;
          let _nw3 = Array((new BigNumber(20)).toNumber());
          _nw3[(_dafny.ZERO).toNumber()] = (_13_v1).plus((_dafny.ZERO).minus(_19_i0));
          _nw3[(_dafny.ONE).toNumber()] = _19_i0;
          _nw3[(new BigNumber(2)).toNumber()] = (_13_v1).plus(_19_i0);
          _nw3[(new BigNumber(3)).toNumber()] = _19_i0;
          _nw3[(new BigNumber(4)).toNumber()] = _module.__default.fm6(globalState);
          _nw3[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_22_v9, _22_v9)).length);
          _nw3[(new BigNumber(6)).toNumber()] = (((_23_v10).contains(_21_v8)) ? ((_23_v10).get(_21_v8)) : (_13_v1));
          _nw3[(new BigNumber(7)).toNumber()] = _19_i0;
          _nw3[(new BigNumber(8)).toNumber()] = _13_v1;
          _nw3[(new BigNumber(9)).toNumber()] = (_13_v1).minus(new BigNumber((_26_v13).length));
          _nw3[(new BigNumber(10)).toNumber()] = new BigNumber(628);
          _nw3[(new BigNumber(11)).toNumber()] = (((_17_v5).contains(_13_v1)) ? ((_17_v5).get(_13_v1)) : (new BigNumber((_27_v14).cardinality())));
          _nw3[(new BigNumber(12)).toNumber()] = new BigNumber(((_29_v16).Merge((_30_v17))).length);
          _nw3[(new BigNumber(13)).toNumber()] = _19_i0;
          _nw3[(new BigNumber(14)).toNumber()] = _13_v1;
          _nw3[(new BigNumber(15)).toNumber()] = _13_v1;
          _nw3[(new BigNumber(16)).toNumber()] = _module.__default.safeDivisionInt(_13_v1, _19_i0);
          _nw3[(new BigNumber(17)).toNumber()] = (((_31_v18).contains(_13_v1)) ? ((_31_v18).get(_13_v1)) : (_13_v1));
          _nw3[(new BigNumber(18)).toNumber()] = (_13_v1).multipliedBy(new BigNumber((_21_v8.f26).length));
          _nw3[(new BigNumber(19)).toNumber()] = _19_i0;
          _32_v19 = _nw3;
          let _index2 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_32_v19).length));
          (_32_v19)[_index2] = _19_i0;
          let _33_v20;
          let _nw4 = Array((new BigNumber(9)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _33_v20 = _nw4;
          let _index3 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_33_v20).length));
          (_33_v20)[_index3] = _18_v6;
          let _index4 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_32_v19).length));
          let _index5 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_33_v20).length));
          let _index6 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((p0).length));
          let _index7 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((p0).length));
          let _rhs0 = (_19_i0).plus(((_24_v11.f28) ? (_13_v1) : (new BigNumber(468))));
          let _rhs1 = _module.__default.safeModuloInt(_19_i0, new BigNumber((_14_v2).length));
          let _rhs2 = _18_v6;
          let _rhs3 = !(_dafny.Seq.IsPrefixOf(_dafny.Seq.update(_dafny.Seq.Concat(_14_v2, _14_v2), _module.__default.safeIndex(_13_v1, new BigNumber((_dafny.Seq.Concat(_14_v2, _14_v2)).length)), _18_v6), _14_v2));
          let _rhs4 = (_24_v11).f27;
          let _lhs0 = _32_v19;
          let _lhs1 = _module.__default.safeIndex(new BigNumber(378), new BigNumber((_32_v19).length));
          let _lhs2 = globalState;
          let _lhs3 = _33_v20;
          let _lhs4 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_33_v20).length));
          let _lhs5 = p0;
          let _lhs6 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((p0).length));
          let _lhs7 = p0;
          let _lhs8 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((p0).length));
          _lhs0[_lhs1] = _rhs0;
          _lhs2.f20 = _rhs1;
          _lhs3[_lhs4] = _rhs2;
          _lhs5[_lhs6] = _rhs3;
          _lhs7[_lhs8] = _rhs4;
          let _34_v21;
          let _nw5 = new _module.C3();
          _nw5.__ctor(p1, true, _14_v2, (p0)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p0).length))]);
          _34_v21 = _nw5;
          let _35_v22;
          _35_v22 = _dafny.Map.Empty.slice().updateUnsafe(_34_v21,p1);
          let _36_v23;
          _36_v23 = _dafny.Seq.of(_34_v21.f22);
          let _37_v24;
          _37_v24 = _dafny.Map.Empty.slice().updateUnsafe(_24_v11.f28,_13_v1);
          let _38_v25;
          _38_v25 = _module.D1.create_DC5(_34_v21.f22, _37_v24, _13_v1, (_dafny.MultiSet.fromElements(new BigNumber(379))).update((_32_v19)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_32_v19).length))], _module.__default.abs(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(627)), ((_39_v20) => function (_40_i1) {
  return (_39_v20)[_module.__default.safeIndex(new BigNumber(84), new BigNumber((_39_v20).length))];
})(_33_v20))).length))));
          let _41_v26;
          _41_v26 = _module.D3.create_DC8(new BigNumber((_36_v23).length), _dafny.Map.Empty.slice().updateUnsafe(_27_v14,new BigNumber((_36_v23).length)), _38_v25, p1);
          let _42_v27;
          _42_v27 = _module.D1.create_DC5((_41_v26).dtor_cf16, _37_v24, (_dafny.ZERO).minus(_13_v1), _dafny.MultiSet.fromElements(_19_i0, (_32_v19)[_module.__default.safeIndex(new BigNumber(378), new BigNumber((_32_v19).length))]));
          let _43_v28;
          _43_v28 = _dafny.Seq.of(_16_v4);
          let _44_v29;
          let _nw6 = new _module.C2();
          _nw6.__ctor(_43_v28, _dafny.Seq.Create(_module.__default.abs(new BigNumber(604)), ((_45_v20) => function (_46_i2) {
            return (_45_v20)[_module.__default.safeIndex(new BigNumber(84), new BigNumber((_45_v20).length))];
          })(_33_v20)), _24_v11.f28);
          _44_v29 = _nw6;
          let _47_v30;
          _47_v30 = _dafny.Seq.of(_44_v29, _44_v29);
          let _48_v31;
          _48_v31 = _dafny.Map.Empty.slice().updateUnsafe(!(_35_v22).equals(_35_v22),(((_42_v27).dtor_cf7) ? (_47_v30) : (_47_v30)));
          _48_v31 = (_48_v31).update(!(_34_v21.f22), _dafny.Seq.Concat(_47_v30, _47_v30));
          _44_v29 = _44_v29;
          let _49_v32;
          _49_v32 = _dafny.Set.fromElements(_44_v29);
          _49_v32 = (_49_v32).Union(_49_v32);
        } else {
          let _50_v33;
          _50_v33 = _dafny.MultiSet.fromElements(p1, false);
          let _51_v34;
          _51_v34 = _dafny.MultiSet.fromElements(_50_v33, _50_v33, _50_v33, _50_v33, _50_v33);
          let _52_v35;
          _52_v35 = _dafny.Seq.of(_module.__default.fm12(_13_v1, _module.__default.fm17(_13_v1, _module.__default.fm12(_13_v1, _50_v33, p1, _51_v34, globalState), globalState), p1, _51_v34, globalState));
          r0 = (_52_v35)[_module.__default.safeIndex(_module.__default.fm6(globalState), new BigNumber((_52_v35).length))];
          r0 = _dafny.areEqual(_14_v2, _14_v2);
          let _53_v36;
          let _init0 = ((_54_v33) => function (_55_i3) {
            return _54_v33;
          })(_50_v33);
          let _nw7 = Array((new BigNumber(12)).toNumber());
          for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw7.length); _i0_0++) {
            _nw7[_i0_0] = _init0(new BigNumber(_i0_0));
          }
          _53_v36 = _nw7;
          let _56_v37;
          let _nw8 = new _module.C1();
          _nw8.__ctor(_53_v36, p1);
          _56_v37 = _nw8;
          _56_v37 = _56_v37;
          let _57_v38;
          _57_v38 = _dafny.Set.fromElements((_56_v37).f25);
          let _58_v39;
          _58_v39 = _dafny.Set.fromElements(_57_v38, _dafny.Set.fromElements(true, (_56_v37).f25));
          let _59_v40;
          _59_v40 = _dafny.Map.Empty.slice().updateUnsafe((_module.D9.create_DC18(_58_v39)).dtor_cf31,p1);
          _59_v40 = (_59_v40).update(_58_v39, p1);
          let _60_v41;
          _60_v41 = _dafny.Map.Empty.slice().updateUnsafe(_19_i0,_18_v6);
          let _61_v42;
          _61_v42 = _dafny.Seq.of(_50_v33, _50_v33);
          r0 = _dafny.Seq.contains(_14_v2, (((_60_v41).contains(new BigNumber(((_61_v42)[_module.__default.safeIndex(_13_v1, new BigNumber((_61_v42).length))]).cardinality()))) ? ((_60_v41).get(new BigNumber(((_61_v42)[_module.__default.safeIndex(_13_v1, new BigNumber((_61_v42).length))]).cardinality()))) : (_18_v6)));
        }
        let _62_v43;
        _62_v43 = _dafny.Map.Empty.slice().updateUnsafe(p0,_13_v1);
        _13_v1 = new BigNumber((_62_v43).length);
        _13_v1 = new BigNumber(-286);
        let _index8 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((p0).length));
        (p0)[_index8] = p1;
        let _63_v44;
        let _nw9 = Array((new BigNumber(10)).toNumber()).fill(false);
        _63_v44 = _nw9;
        let _64_v45;
        _64_v45 = _dafny.Seq.of(p0);
        let _index9 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((p0).length));
        let _rhs5 = _dafny.Seq.contains(_16_v4, (_13_v1).minus(_19_i0));
        let _rhs6 = ((!(true)) ? (p0) : ((_64_v45)[_module.__default.safeIndex(new BigNumber(513), new BigNumber((_64_v45).length))]));
        let _lhs9 = p0;
        let _lhs10 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((p0).length));
        _lhs9[_lhs10] = _rhs5;
        _63_v44 = _rhs6;
      }
      let _65_v46;
      _65_v46 = _dafny.MultiSet.fromElements(p1);
      let _66_v47;
      _66_v47 = _dafny.Map.Empty.slice().updateUnsafe(_65_v46,_module.__default.fm6(globalState));
      let _67_v48;
      _67_v48 = _module.D1.create_DC5(p1, _dafny.Map.Empty.slice().updateUnsafe(p1,(_dafny.ZERO).minus(_13_v1)), _13_v1, _17_v5);
      let _pat_let_tv0 = _16_v4;
      let _pat_let_tv1 = _16_v4;
      let _pat_let_tv2 = _16_v4;
      let _rhs7 = function (_source0) {
        if (_source0.is_DC8) {
          let _68___mcc_h0 = (_source0).cf13;
          let _69___mcc_h1 = (_source0).cf14;
          let _70___mcc_h2 = (_source0).cf15;
          let _71___mcc_h3 = (_source0).cf16;
          let _72_cf16 = _71___mcc_h3;
          let _73_cf15 = _70___mcc_h2;
          let _74_cf14 = _69___mcc_h1;
          let _75_cf13 = _68___mcc_h0;
          return _pat_let_tv0;
        } else if (_source0.is_DC9) {
          let _76___mcc_h4 = (_source0).cf17;
          let _77___mcc_h5 = (_source0).cf18;
          let _78___mcc_h6 = (_source0).cf19;
          let _79_cf19 = _78___mcc_h6;
          let _80_cf18 = _77___mcc_h5;
          let _81_cf17 = _76___mcc_h4;
          return _pat_let_tv1;
        } else {
          let _82___mcc_h7 = (_source0).cf12;
          let _83_cf12 = _82___mcc_h7;
          return _pat_let_tv2;
        }
      }(_module.D3.create_DC8(_13_v1, _66_v47, _67_v48, !(p1)));
      let _rhs8 = p1;
      let _rhs9 = p1;
      _16_v4 = _rhs7;
      r0 = _rhs8;
      r0 = _rhs9;
      let _84_v49;
      _84_v49 = _dafny.Set.fromElements(new BigNumber((_12_v0).length));
      let _85_v50;
      _85_v50 = _84_v49;
      let _86_v51;
      _86_v51 = _dafny.Seq.of(_dafny.Seq.update(_16_v4, _module.__default.safeIndex(_13_v1, new BigNumber((_16_v4).length)), _13_v1), _16_v4, _16_v4);
      let _87_v52;
      let _nw10 = new _module.C4();
      _nw10.__ctor(_85_v50, p1, _86_v51, _14_v2, (new BigNumber((_84_v49).length)).isEqualTo(_13_v1));
      _87_v52 = _nw10;
      let _88_v53;
      _88_v53 = _dafny.Set.fromElements(false);
      let _89_v54;
      _89_v54 = _dafny.Map.Empty.slice().updateUnsafe(_88_v53,p1);
      let _90_v56;
      _90_v56 = _module.D9.create_DC18(function () {
  let _coll7 = new _dafny.Set();
  for (const _compr_7 of (_89_v54).Keys.Elements) {
    let _91_v55 = _compr_7;
    if ((_89_v54).contains(_91_v55)) {
      _coll7.add(_91_v55);
    }
  }
  return _coll7;
}());
      let _92_v57;
      _92_v57 = _dafny.Map.Empty.slice().updateUnsafe(true,p1);
      let _93_v58;
      _93_v58 = _module.D9.create_DC20(_87_v52.f21, p1, new BigNumber((_92_v57).length));
      let _94_v59;
      _94_v59 = _module.D9.create_DC20((_93_v58).dtor_cf37, true, _13_v1);
      let _pat_let_tv3 = _13_v1;
      let _pat_let_tv4 = _67_v48;
      let _pat_let_tv5 = _13_v1;
      let _pat_let_tv6 = _13_v1;
      let _pat_let_tv7 = _65_v46;
      let _pat_let_tv8 = _65_v46;
      let _pat_let_tv9 = p1;
      let _pat_let_tv10 = p1;
      let _pat_let_tv11 = p1;
      let _pat_let_tv12 = _65_v46;
      let _pat_let_tv13 = _65_v46;
      let _rhs10 = function (_source1) {
        if (_source1.is_DC19) {
          let _95___mcc_h8 = (_source1).cf32;
          let _96___mcc_h9 = (_source1).cf33;
          let _97___mcc_h10 = (_source1).cf34;
          let _98___mcc_h11 = (_source1).cf35;
          let _99___mcc_h12 = (_source1).cf36;
          let _100_cf36 = _99___mcc_h12;
          let _101_cf35 = _98___mcc_h11;
          let _102_cf34 = _97___mcc_h10;
          let _103_cf33 = _96___mcc_h9;
          let _104_cf32 = _95___mcc_h8;
          return _pat_let_tv3;
        } else if (_source1.is_DC20) {
          let _105___mcc_h13 = (_source1).cf37;
          let _106___mcc_h14 = (_source1).cf38;
          let _107___mcc_h15 = (_source1).cf39;
          let _108_cf39 = _107___mcc_h15;
          let _109_cf38 = _106___mcc_h14;
          let _110_cf37 = _105___mcc_h13;
          return new BigNumber((((_pat_let_tv4).dtor_cf8).Merge(_dafny.Map.Empty.slice().updateUnsafe(_109_cf38,_pat_let_tv5))).length);
        } else if (_source1.is_DC21) {
          let _111___mcc_h16 = (_source1).cf40;
          let _112___mcc_h17 = (_source1).cf41;
          let _113_cf41 = _112___mcc_h17;
          let _114_cf40 = _111___mcc_h16;
          return new BigNumber(995);
        } else {
          let _115___mcc_h18 = (_source1).cf31;
          let _116_cf31 = _115___mcc_h18;
          return _pat_let_tv6;
        }
      }(_90_v56);
      let _rhs11 = _87_v52;
      let _rhs12 = function (_source2) {
        if (_source2.is_DC19) {
          let _117___mcc_h19 = (_source2).cf32;
          let _118___mcc_h20 = (_source2).cf33;
          let _119___mcc_h21 = (_source2).cf34;
          let _120___mcc_h22 = (_source2).cf35;
          let _121___mcc_h23 = (_source2).cf36;
          let _122_cf36 = _121___mcc_h23;
          let _123_cf35 = _120___mcc_h22;
          let _124_cf34 = _119___mcc_h21;
          let _125_cf33 = _118___mcc_h20;
          let _126_cf32 = _117___mcc_h19;
          return _pat_let_tv7;
        } else if (_source2.is_DC20) {
          let _127___mcc_h24 = (_source2).cf37;
          let _128___mcc_h25 = (_source2).cf38;
          let _129___mcc_h26 = (_source2).cf39;
          let _130_cf39 = _129___mcc_h26;
          let _131_cf38 = _128___mcc_h25;
          let _132_cf37 = _127___mcc_h24;
          return _pat_let_tv8;
        } else if (_source2.is_DC21) {
          let _133___mcc_h27 = (_source2).cf40;
          let _134___mcc_h28 = (_source2).cf41;
          let _135_cf41 = _134___mcc_h28;
          let _136_cf40 = _133___mcc_h27;
          return _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(_pat_let_tv9, _pat_let_tv10), _dafny.Seq.of(_135_cf41, !(_pat_let_tv11), false, true, _135_cf41)));
        } else {
          let _137___mcc_h29 = (_source2).cf31;
          let _138_cf31 = _137___mcc_h29;
          return (_pat_let_tv12).Union(_pat_let_tv13);
        }
      }(_94_v59);
      let _lhs11 = globalState;
      _lhs11.f3 = _rhs10;
      _87_v52 = _rhs11;
      _65_v46 = _rhs12;
      (_87_v52).f22 = p1;
      let _139_v60;
      _139_v60 = _dafny.Seq.of(_87_v52.f22, _87_v52.f22, _87_v52.f22, _87_v52.f22, true);
      r0 = (!(!(!_dafny.Seq.contains(_139_v60, _87_v52.f22)))) && (_87_v52.f22);
      return r0;
    }
    static Main(__noArgsParameter) {
      let _140_v1;
      _140_v1 = _dafny.Seq.UnicodeFromString("truxk");
      let _141_v2;
      _141_v2 = _dafny.Set.fromElements(_140_v1, _140_v1);
      let _142_v3;
      _142_v3 = new BigNumber(-610);
      let _143_v4;
      _143_v4 = new _dafny.CodePoint('g'.codePointAt(0));
      let _144_v5;
      let _nw11 = Array((new BigNumber(29)).toNumber()).fill(false);
      _144_v5 = _nw11;
      let _145_v6;
      let _nw12 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
      _145_v6 = _nw12;
      let _146_v7;
      _146_v7 = _dafny.MultiSet.fromElements(_142_v3);
      let _147_globalState;
      let _nw13 = new _module.GlobalState();
      _nw13.__ctor(false, function () {
        let _coll8 = new _dafny.Map();
        for (const _compr_8 of (_141_v2).Elements) {
          let _148_v0 = _compr_8;
          if ((_141_v2).contains(_148_v0)) {
            _coll8.push([_148_v0,new BigNumber((_dafny.Seq.of(new BigNumber(281))).length)]);
          }
        }
        return _coll8;
      }(), new BigNumber(-658), new BigNumber(2), new BigNumber(-712), true, _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("v"), _module.__default.safeIndex((_dafny.ZERO).minus(_142_v3), new BigNumber((_dafny.Seq.UnicodeFromString("v")).length)), _143_v4), _module.__default.safeIndex(_142_v3, new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("v"), _module.__default.safeIndex((_dafny.ZERO).minus(_142_v3), new BigNumber((_dafny.Seq.UnicodeFromString("v")).length)), _143_v4)).length)), _143_v4), true, _140_v1, _140_v1, new BigNumber(896), true, new _dafny.CodePoint('c'.codePointAt(0)), new BigNumber(-158), new BigNumber(316), new BigNumber(-654), _144_v5, _145_v6, false, _146_v7, new BigNumber(308));
      _147_globalState = _nw13;
      let _149_v8;
      _149_v8 = false;
      let _150_i0;
      _150_i0 = _dafny.ZERO;
      L0: {
        while (!(!(_149_v8))) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_150_i0)) {
              break L0;
            }
            _150_i0 = (_150_i0).plus(_dafny.ONE);
            let _151_v9;
            let _nw14 = Array((new BigNumber(27)).toNumber()).fill([]);
            _151_v9 = _nw14;
            let _index10 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_151_v9).length));
            (_151_v9)[_index10] = _144_v5;
            let _index11 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_151_v9).length));
            let _rhs13 = (_142_v3).minus(((_149_v8) ? (_142_v3) : (_142_v3)));
            let _rhs14 = _144_v5;
            let _rhs15 = _142_v3;
            let _lhs12 = _151_v9;
            let _lhs13 = _module.__default.safeIndex(new BigNumber(512), new BigNumber((_151_v9).length));
            let _lhs14 = _147_globalState;
            _142_v3 = _rhs13;
            _lhs12[_lhs13] = _rhs14;
            _lhs14.f20 = _rhs15;
            (_147_globalState).f20 = _142_v3;
            _142_v3 = _142_v3;
            let _152_v10;
            let _out0;
            _out0 = _module.__default.m0(_144_v5, _149_v8, _147_globalState);
            _152_v10 = _out0;
          }
        }
      }
      if (_149_v8) {
        let _153_v11;
        let _init1 = function (_154_i1) {
          return _dafny.MultiSet.fromElements(true);
        };
        let _nw15 = Array((new BigNumber(23)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw15.length); _i0_1++) {
          _nw15[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _153_v11 = _nw15;
        let _155_v12;
        let _nw16 = new _module.C1();
        _nw16.__ctor(_153_v11, !(false) || (_149_v8));
        _155_v12 = _nw16;
        let _156_v13;
        _156_v13 = _dafny.MultiSet.fromElements(true);
        let _157_v14;
        _157_v14 = _dafny.MultiSet.fromElements(_156_v13);
        let _158_v15;
        _158_v15 = _dafny.Set.fromElements((_155_v12).f25, false);
        let _159_v16;
        let _nw17 = new _module.C2();
        _nw17.__ctor(_module.__default.fm16(_158_v15, (_155_v12).f25, _149_v8, _147_globalState), _140_v1, _149_v8);
        _159_v16 = _nw17;
        let _160_v17;
        _160_v17 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm12(_142_v3, _156_v13, (_155_v12).f25, (_157_v14).update(_156_v13, _module.__default.abs(new BigNumber((_140_v1).length))), _147_globalState),_159_v16);
        let _161_v18;
        _161_v18 = _dafny.Map.Empty.slice().updateUnsafe(_140_v1,_160_v17);
        let _162_v19;
        _162_v19 = _dafny.Map.Empty.slice().updateUnsafe(_149_v8,false);
        let _163_v20;
        _163_v20 = _dafny.Seq.of(_162_v19, _162_v19, _dafny.Map.Empty.slice().updateUnsafe(_149_v8,false));
        let _rhs16 = _142_v3;
        let _rhs17 = _161_v18;
        let _rhs18 = _module.__default.safeDivisionInt(_142_v3, _142_v3);
        let _rhs19 = !_dafny.Seq.contains(_163_v20, _162_v19);
        let _lhs15 = _147_globalState;
        let _lhs16 = _147_globalState;
        _lhs15.f3 = _rhs16;
        _161_v18 = _rhs17;
        _lhs16.f3 = _rhs18;
        _149_v8 = _rhs19;
        let _164_v21;
        let _out1;
        _out1 = _module.__default.m0(_144_v5, !((_155_v12).f25), _147_globalState);
        _164_v21 = _out1;
        _149_v8 = ((_155_v12).f25) || ((_155_v12).f25);
        _142_v3 = (((_159_v16).fm3(_147_globalState)) ? (_142_v3) : (_142_v3));
      } else {
        let _index12 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_145_v6).length));
        (_145_v6)[_index12] = new BigNumber(-385);
        let _165_v22;
        _165_v22 = _dafny.Seq.of(_142_v3);
        let _index13 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_145_v6).length));
        (_145_v6)[_index13] = (_dafny.ZERO).minus(((_165_v22)[_module.__default.safeIndex(_142_v3, new BigNumber((_165_v22).length))]).multipliedBy(_142_v3));
        let _166_v23;
        _166_v23 = _dafny.MultiSet.fromElements(_149_v8, _149_v8);
        let _167_v24;
        _167_v24 = _dafny.Map.Empty.slice().updateUnsafe((_145_v6)[_module.__default.safeIndex(new BigNumber(347), new BigNumber((_145_v6).length))],_149_v8);
        let _168_v25;
        _168_v25 = _dafny.MultiSet.fromElements(_166_v23);
        let _169_v26;
        let _nw18 = Array((new BigNumber(27)).toNumber());
        _nw18[(_dafny.ZERO).toNumber()] = _166_v23;
        _nw18[(_dafny.ONE).toNumber()] = _dafny.MultiSet.fromElements(true);
        _nw18[(new BigNumber(2)).toNumber()] = _166_v23;
        _nw18[(new BigNumber(3)).toNumber()] = _166_v23;
        _nw18[(new BigNumber(4)).toNumber()] = _166_v23;
        _nw18[(new BigNumber(5)).toNumber()] = _166_v23;
        _nw18[(new BigNumber(6)).toNumber()] = _166_v23;
        _nw18[(new BigNumber(7)).toNumber()] = _166_v23;
        _nw18[(new BigNumber(8)).toNumber()] = _166_v23;
        _nw18[(new BigNumber(9)).toNumber()] = _166_v23;
        _nw18[(new BigNumber(10)).toNumber()] = _166_v23;
        _nw18[(new BigNumber(11)).toNumber()] = _166_v23;
        _nw18[(new BigNumber(12)).toNumber()] = _166_v23;
        _nw18[(new BigNumber(13)).toNumber()] = _dafny.MultiSet.fromElements(_149_v8, _149_v8, _149_v8);
        _nw18[(new BigNumber(14)).toNumber()] = _166_v23;
        _nw18[(new BigNumber(15)).toNumber()] = _166_v23;
        _nw18[(new BigNumber(16)).toNumber()] = _dafny.MultiSet.fromElements(_149_v8);
        _nw18[(new BigNumber(17)).toNumber()] = _dafny.MultiSet.fromElements(_149_v8, _149_v8);
        _nw18[(new BigNumber(18)).toNumber()] = _dafny.MultiSet.fromElements(_149_v8, _149_v8);
        _nw18[(new BigNumber(19)).toNumber()] = _dafny.MultiSet.fromElements(_149_v8, _149_v8, _module.__default.fm12(_142_v3, _dafny.MultiSet.fromElements(_149_v8, true, (((_167_v24).contains(_142_v3)) ? ((_167_v24).get(_142_v3)) : (_149_v8))), !(_module.__default.fm12(_142_v3, _dafny.MultiSet.fromElements(_149_v8, _149_v8), _149_v8, _168_v25, _147_globalState)), _168_v25, _147_globalState));
        _nw18[(new BigNumber(20)).toNumber()] = _166_v23;
        _nw18[(new BigNumber(21)).toNumber()] = _module.__default.fm17((_145_v6)[_module.__default.safeIndex(new BigNumber(347), new BigNumber((_145_v6).length))], _149_v8, _147_globalState);
        _nw18[(new BigNumber(22)).toNumber()] = _dafny.MultiSet.fromElements(_149_v8, _149_v8, _149_v8);
        _nw18[(new BigNumber(23)).toNumber()] = _166_v23;
        _nw18[(new BigNumber(24)).toNumber()] = _166_v23;
        _nw18[(new BigNumber(25)).toNumber()] = _dafny.MultiSet.fromElements(_149_v8);
        _nw18[(new BigNumber(26)).toNumber()] = _166_v23;
        _169_v26 = _nw18;
        let _170_v27;
        _170_v27 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm12(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_149_v8,_142_v3)).length), _166_v23, _149_v8, _168_v25, _147_globalState),_142_v3);
        let _171_v28;
        let _nw19 = new _module.C1();
        _nw19.__ctor(_169_v26, (new BigNumber((_170_v27).length)).isLessThanOrEqualTo((_145_v6)[_module.__default.safeIndex(new BigNumber(347), new BigNumber((_145_v6).length))]));
        _171_v28 = _nw19;
        let _172_v29;
        _172_v29 = _dafny.Seq.of(true, _149_v8);
        let _173_v30;
        _173_v30 = _module.D3.create_DC7(_dafny.MultiSet.FromArray(_172_v29));
        let _index14 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_145_v6).length));
        let _rhs20 = _171_v28;
        let _rhs21 = (((_142_v3).isLessThanOrEqualTo(_142_v3)) ? (new BigNumber(((_141_v2).Union(_141_v2)).length)) : (new BigNumber(-794)));
        let _rhs22 = _173_v30;
        let _rhs23 = (_142_v3).isLessThanOrEqualTo(_module.__default.safeModuloInt(new BigNumber((_165_v22).length), _142_v3));
        let _lhs17 = _145_v6;
        let _lhs18 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_145_v6).length));
        _171_v28 = _rhs20;
        _lhs17[_lhs18] = _rhs21;
        _173_v30 = _rhs22;
        _149_v8 = _rhs23;
        let _174_v31;
        let _init2 = function (_175_i2) {
          return _dafny.Set.fromElements(true);
        };
        let _nw20 = Array((new BigNumber(24)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw20.length); _i0_2++) {
          _nw20[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _174_v31 = _nw20;
        let _176_v32;
        let _nw21 = Array((new BigNumber(4)).toNumber());
        _nw21[(_dafny.ZERO).toNumber()] = _174_v31;
        _nw21[(_dafny.ONE).toNumber()] = _174_v31;
        _nw21[(new BigNumber(2)).toNumber()] = _174_v31;
        _nw21[(new BigNumber(3)).toNumber()] = _174_v31;
        _176_v32 = _nw21;
        let _index15 = _module.__default.safeIndex(new BigNumber(134), new BigNumber((_176_v32).length));
        (_176_v32)[_index15] = _174_v31;
        let _index16 = _module.__default.safeIndex(new BigNumber(134), new BigNumber((_176_v32).length));
        (_176_v32)[_index16] = _174_v31;
        let _177_v33;
        _177_v33 = _dafny.Seq.of(_165_v22, _module.__default.fm5(_142_v3, _142_v3, _142_v3, _147_globalState));
        let _178_v34;
        let _nw22 = new _module.C2();
        _nw22.__ctor(_177_v33, _dafny.Seq.UnicodeFromString("m"), _149_v8);
        _178_v34 = _nw22;
        let _179_v35;
        _179_v35 = _dafny.Set.fromElements(_178_v34);
        _179_v35 = _dafny.Set.fromElements(_178_v34, _178_v34, _178_v34, _178_v34, _178_v34);
        let _180_v36;
        _180_v36 = _module.D0.create_DC1((_171_v28).f25);
        let _181_v37;
        let _nw23 = new _module.C1();
        _nw23.__ctor(_171_v28.f24, (_180_v36).dtor_cf1);
        _181_v37 = _nw23;
      }
      _149_v8 = _149_v8;
      let _182_v38;
      _182_v38 = _dafny.MultiSet.fromElements(false);
      let _183_v39;
      _183_v39 = _dafny.MultiSet.fromElements(_182_v38, _182_v38);
      let _184_v40;
      _184_v40 = _module.D0.create_DC0(_144_v5);
      let _185_v41;
      _185_v41 = _module.D0.create_DC2(_184_v40);
      let _source3 = ((_module.__default.fm12(_142_v3, _182_v38, _149_v8, _183_v39, _147_globalState)) ? (_185_v41) : (_185_v41));
      if (_source3.is_DC1) {
        let _186___mcc_h0 = (_source3).cf1;
        let _187_cf1 = _186___mcc_h0;
        let _index17 = _module.__default.safeIndex(new BigNumber(881), new BigNumber((_145_v6).length));
        (_145_v6)[_index17] = new BigNumber(203);
        let _188_v42;
        _188_v42 = _dafny.Seq.of(_142_v3, new BigNumber(555), _142_v3, _142_v3);
        let _189_v43;
        _189_v43 = _dafny.Seq.of(_149_v8, _149_v8);
        let _190_v44;
        let _nw24 = Array((new BigNumber(19)).toNumber());
        _nw24[(_dafny.ZERO).toNumber()] = _182_v38;
        _nw24[(_dafny.ONE).toNumber()] = _182_v38;
        _nw24[(new BigNumber(2)).toNumber()] = _182_v38;
        _nw24[(new BigNumber(3)).toNumber()] = _182_v38;
        _nw24[(new BigNumber(4)).toNumber()] = _182_v38;
        _nw24[(new BigNumber(5)).toNumber()] = _182_v38;
        _nw24[(new BigNumber(6)).toNumber()] = _182_v38;
        _nw24[(new BigNumber(7)).toNumber()] = _dafny.MultiSet.fromElements(_187_cf1, _187_cf1, _187_cf1);
        _nw24[(new BigNumber(8)).toNumber()] = _182_v38;
        _nw24[(new BigNumber(9)).toNumber()] = _module.__default.fm17(new BigNumber((_188_v42).length), _149_v8, _147_globalState);
        _nw24[(new BigNumber(10)).toNumber()] = _182_v38;
        _nw24[(new BigNumber(11)).toNumber()] = _182_v38;
        _nw24[(new BigNumber(12)).toNumber()] = _182_v38;
        _nw24[(new BigNumber(13)).toNumber()] = _182_v38;
        _nw24[(new BigNumber(14)).toNumber()] = _dafny.MultiSet.FromArray(_189_v43);
        _nw24[(new BigNumber(15)).toNumber()] = _182_v38;
        _nw24[(new BigNumber(16)).toNumber()] = _182_v38;
        _nw24[(new BigNumber(17)).toNumber()] = _182_v38;
        _nw24[(new BigNumber(18)).toNumber()] = _module.__default.fm17(_142_v3, true, _147_globalState);
        _190_v44 = _nw24;
        let _191_v45;
        _191_v45 = _dafny.Seq.of(_190_v44);
        let _192_v46;
        let _nw25 = new _module.C1();
        _nw25.__ctor((_191_v45)[_module.__default.safeIndex(new BigNumber(616), new BigNumber((_191_v45).length))], _149_v8);
        _192_v46 = _nw25;
        let _193_v47;
        _193_v47 = _dafny.Seq.of(_192_v46, _192_v46, _192_v46);
        let _index18 = _module.__default.safeIndex(new BigNumber(881), new BigNumber((_145_v6).length));
        let _rhs24 = (_142_v3).plus(new BigNumber((_193_v47).length));
        let _rhs25 = (_142_v3).minus(new BigNumber(-786));
        let _lhs19 = _145_v6;
        let _lhs20 = _module.__default.safeIndex(new BigNumber(881), new BigNumber((_145_v6).length));
        let _lhs21 = _147_globalState;
        _lhs19[_lhs20] = _rhs24;
        _lhs21.f20 = _rhs25;
        let _index19 = _module.__default.safeIndex(new BigNumber(881), new BigNumber((_145_v6).length));
        (_145_v6)[_index19] = new BigNumber(723);
        let _194_v48;
        let _nw26 = Array((new BigNumber(19)).toNumber());
        _194_v48 = _nw26;
        let _195_v49;
        let _nw27 = new _module.C3();
        _nw27.__ctor((_192_v46).f25, _149_v8, _dafny.Seq.update(_dafny.Seq.UnicodeFromString("r"), _module.__default.safeIndex((_dafny.ZERO).minus(_142_v3), new BigNumber((_dafny.Seq.UnicodeFromString("r")).length)), _143_v4), false);
        _195_v49 = _nw27;
        let _index20 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_194_v48).length));
        (_194_v48)[_index20] = _195_v49;
        let _index21 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_194_v48).length));
        (_194_v48)[_index21] = _195_v49;
        let _196_v50;
        let _nw28 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _196_v50 = _nw28;
        _196_v50 = _196_v50;
      } else if (_source3.is_DC0) {
        let _197___mcc_h1 = (_source3).cf0;
        let _198_cf0 = _197___mcc_h1;
        let _199_v51;
        _199_v51 = _dafny.Seq.of(_142_v3);
        let _200_v52;
        _200_v52 = _dafny.Seq.of(_199_v51);
        let _201_v53;
        let _nw29 = new _module.C2();
        _nw29.__ctor(_200_v52, _140_v1, ((_149_v8) ? (_149_v8) : (_149_v8)));
        _201_v53 = _nw29;
        _201_v53 = _201_v53;
        let _202_v54;
        _202_v54 = _dafny.Map.Empty.slice().updateUnsafe(_140_v1,_142_v3);
        let _203_v56;
        let _nw30 = new _module.C0();
        _nw30.__ctor(function () {
          let _coll9 = new _dafny.Set();
          for (const _compr_9 of (_202_v54).Keys.Elements) {
            let _204_v55 = _compr_9;
            if ((_202_v54).contains(_204_v55)) {
              _coll9.add(_204_v55);
            }
          }
          return _coll9;
        }());
        _203_v56 = _nw30;
        (_147_globalState).f8 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(140)), ((_205_v4) => function (_206_i3) {
          return _205_v4;
        })(_143_v4));
        let _207_v57;
        _207_v57 = _dafny.Seq.of(!(_201_v53.f22), true, _201_v53.f22);
        if (_dafny.Seq.IsProperPrefixOf(_207_v57, _207_v57)) {
          let _index22 = _module.__default.safeIndex(new BigNumber(657), new BigNumber((_145_v6).length));
          (_145_v6)[_index22] = new BigNumber(842);
          let _index23 = _module.__default.safeIndex(new BigNumber(483), new BigNumber((_144_v5).length));
          (_144_v5)[_index23] = !(_149_v8);
          let _208_v58;
          _208_v58 = _dafny.Seq.of(_201_v53);
          let _209_v59;
          _209_v59 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('y'.codePointAt(0))));
          let _210_v60;
          _210_v60 = _dafny.Map.Empty.slice().updateUnsafe(_209_v59,_201_v53);
          let _211_v61;
          let _nw31 = new _module.C3();
          _nw31.__ctor(_201_v53.f22, _149_v8, _201_v53.f21, false);
          _211_v61 = _nw31;
          let _212_v62;
          _212_v62 = _dafny.MultiSet.fromElements(_211_v61);
          let _213_v63;
          _213_v63 = _211_v61;
          let _214_v64;
          let _nw32 = Array((new BigNumber(28)).toNumber());
          _nw32[(_dafny.ZERO).toNumber()] = _201_v53;
          _nw32[(_dafny.ONE).toNumber()] = _201_v53;
          _nw32[(new BigNumber(2)).toNumber()] = _201_v53;
          _nw32[(new BigNumber(3)).toNumber()] = _201_v53;
          _nw32[(new BigNumber(4)).toNumber()] = _201_v53;
          _nw32[(new BigNumber(5)).toNumber()] = (_208_v58)[_module.__default.safeIndex((_dafny.ZERO).minus(_142_v3), new BigNumber((_208_v58).length))];
          _nw32[(new BigNumber(6)).toNumber()] = _201_v53;
          _nw32[(new BigNumber(7)).toNumber()] = _201_v53;
          _nw32[(new BigNumber(8)).toNumber()] = _201_v53;
          _nw32[(new BigNumber(9)).toNumber()] = _201_v53;
          _nw32[(new BigNumber(10)).toNumber()] = (((_210_v60).contains(_209_v59)) ? ((_210_v60).get(_209_v59)) : (_201_v53));
          _nw32[(new BigNumber(11)).toNumber()] = _201_v53;
          _nw32[(new BigNumber(12)).toNumber()] = _201_v53;
          _nw32[(new BigNumber(13)).toNumber()] = ((_201_v53.f22) ? (_201_v53) : (_201_v53));
          _nw32[(new BigNumber(14)).toNumber()] = _201_v53;
          _nw32[(new BigNumber(15)).toNumber()] = _201_v53;
          _nw32[(new BigNumber(16)).toNumber()] = _201_v53;
          _nw32[(new BigNumber(17)).toNumber()] = _201_v53;
          _nw32[(new BigNumber(18)).toNumber()] = (_208_v58)[_module.__default.safeIndex((((_212_v62).contains((_213_v63))) ? ((_212_v62).get((_213_v63))) : (_142_v3)), new BigNumber((_208_v58).length))];
          _nw32[(new BigNumber(19)).toNumber()] = (_208_v58)[_module.__default.safeIndex(new BigNumber((_201_v53.f21).length), new BigNumber((_208_v58).length))];
          _nw32[(new BigNumber(20)).toNumber()] = _201_v53;
          _nw32[(new BigNumber(21)).toNumber()] = _201_v53;
          _nw32[(new BigNumber(22)).toNumber()] = _201_v53;
          _nw32[(new BigNumber(23)).toNumber()] = _201_v53;
          _nw32[(new BigNumber(24)).toNumber()] = _201_v53;
          _nw32[(new BigNumber(25)).toNumber()] = _201_v53;
          _nw32[(new BigNumber(26)).toNumber()] = _201_v53;
          _nw32[(new BigNumber(27)).toNumber()] = _201_v53;
          _214_v64 = _nw32;
          let _index24 = _module.__default.safeIndex(new BigNumber(357), new BigNumber((_214_v64).length));
          (_214_v64)[_index24] = _201_v53;
          let _215_v65;
          _215_v65 = _dafny.Seq.of(_201_v53.f21, _140_v1);
          let _216_v66;
          _216_v66 = _dafny.Map.Empty.slice().updateUnsafe(_211_v61.f28,_142_v3);
          let _217_v67;
          _217_v67 = _dafny.Map.Empty.slice().updateUnsafe(_216_v66,_142_v3);
          let _218_v68;
          let _nw33 = Array((new BigNumber(14)).toNumber());
          _nw33[(_dafny.ZERO).toNumber()] = (_215_v65)[_module.__default.safeIndex(new BigNumber((_217_v67).length), new BigNumber((_215_v65).length))];
          _nw33[(_dafny.ONE).toNumber()] = _140_v1;
          _nw33[(new BigNumber(2)).toNumber()] = _201_v53.f21;
          _nw33[(new BigNumber(3)).toNumber()] = _140_v1;
          _nw33[(new BigNumber(4)).toNumber()] = _201_v53.f21;
          _nw33[(new BigNumber(5)).toNumber()] = _module.__default.fm11((_211_v61).f27, _147_globalState);
          _nw33[(new BigNumber(6)).toNumber()] = _201_v53.f21;
          _nw33[(new BigNumber(7)).toNumber()] = _201_v53.f21;
          _nw33[(new BigNumber(8)).toNumber()] = _201_v53.f21;
          _nw33[(new BigNumber(9)).toNumber()] = _140_v1;
          _nw33[(new BigNumber(10)).toNumber()] = _module.__default.fm11(_211_v61.f28, _147_globalState);
          _nw33[(new BigNumber(11)).toNumber()] = _201_v53.f21;
          _nw33[(new BigNumber(12)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(434)), ((_219_v4) => function (_220_i4) {
            return _219_v4;
          })(_143_v4));
          _nw33[(new BigNumber(13)).toNumber()] = _201_v53.f21;
          _218_v68 = _nw33;
          let _221_v69;
          _221_v69 = _dafny.MultiSet.fromElements(_218_v68);
          let _222_v70;
          _222_v70 = _dafny.Set.fromElements(_145_v6, _145_v6, _145_v6, _145_v6, _145_v6);
          let _index25 = _module.__default.safeIndex(new BigNumber(657), new BigNumber((_145_v6).length));
          let _index26 = _module.__default.safeIndex(new BigNumber(483), new BigNumber((_144_v5).length));
          let _index27 = _module.__default.safeIndex(new BigNumber(357), new BigNumber((_214_v64).length));
          let _rhs26 = (((_221_v69).contains(_218_v68)) ? ((_221_v69).get(_218_v68)) : (_142_v3));
          let _rhs27 = (_222_v70).IsDisjointFrom((_222_v70).Intersect(_dafny.Set.fromElements(_145_v6)));
          let _rhs28 = _201_v53;
          let _lhs22 = _145_v6;
          let _lhs23 = _module.__default.safeIndex(new BigNumber(657), new BigNumber((_145_v6).length));
          let _lhs24 = _144_v5;
          let _lhs25 = _module.__default.safeIndex(new BigNumber(483), new BigNumber((_144_v5).length));
          let _lhs26 = _214_v64;
          let _lhs27 = _module.__default.safeIndex(new BigNumber(357), new BigNumber((_214_v64).length));
          _lhs22[_lhs23] = _rhs26;
          _lhs24[_lhs25] = _rhs27;
          _lhs26[_lhs27] = _rhs28;
          let _223_v71;
          let _nw34 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
          _223_v71 = _nw34;
          _223_v71 = _223_v71;
          let _224_v72;
          _224_v72 = _module.D1.create_DC5(_149_v8, _216_v66, _142_v3, _146_v7);
          let _225_v73;
          let _nw35 = new _module.C3();
          _nw35.__ctor((_142_v3).isEqualTo(_142_v3), _211_v61.f28, _201_v53.f21, (_142_v3).isLessThanOrEqualTo((_224_v72).dtor_cf9));
          _225_v73 = _nw35;
          let _226_v74;
          _226_v74 = _dafny.Set.fromElements(_142_v3);
          let _227_v75;
          _227_v75 = _226_v74;
          let _228_v76;
          let _nw36 = new _module.C4();
          _nw36.__ctor(_227_v75, (_211_v61).f27, _200_v52, _140_v1, (_142_v3).isLessThan(new BigNumber((_207_v57).length)));
          _228_v76 = _nw36;
          let _rhs29 = _module.__default.fm11((_225_v73).fm1(_140_v1, (_144_v5)[_module.__default.safeIndex(new BigNumber(483), new BigNumber((_144_v5).length))], (_211_v61).f27, _147_globalState), _147_globalState);
          let _rhs30 = _225_v73;
          let _rhs31 = _228_v76;
          let _lhs28 = _201_v53;
          _lhs28.f21 = _rhs29;
          _225_v73 = _rhs30;
          _228_v76 = _rhs31;
          let _229_v77;
          _229_v77 = _dafny.Map.Empty.slice().updateUnsafe((_145_v6)[_module.__default.safeIndex(new BigNumber(657), new BigNumber((_145_v6).length))],_225_v73.f22);
          let _230_v78;
          _230_v78 = _dafny.Map.Empty.slice().updateUnsafe(_229_v77,_198_cf0);
          let _231_v79;
          let _out2;
          _out2 = (_201_v53).m1(_225_v73.f22, _230_v78, _147_globalState);
          _231_v79 = _out2;
          let _232_v80;
          let _nw37 = Array((new BigNumber(19)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _232_v80 = _nw37;
          let _index28 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_232_v80).length));
          (_232_v80)[_index28] = _143_v4;
          let _233_v81;
          _233_v81 = _dafny.Set.fromElements(!((_211_v61).f27), _149_v8);
          let _234_v82;
          let _nw38 = new _module.C4();
          _nw38.__ctor(_227_v75, !(true) || (_149_v8), _module.__default.fm16(_233_v81, _211_v61.f28, _149_v8, _147_globalState), _dafny.Seq.update(_dafny.Seq.UnicodeFromString("lpt"), _module.__default.safeIndex(_142_v3, new BigNumber((_dafny.Seq.UnicodeFromString("lpt")).length)), _143_v4), _201_v53.f22);
          _234_v82 = _nw38;
          let _index29 = _module.__default.safeIndex(new BigNumber(483), new BigNumber((_144_v5).length));
          let _index30 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_232_v80).length));
          let _rhs32 = (_234_v82).fm1(_140_v1, !((_234_v82).f30), _211_v61.f28, _147_globalState);
          let _rhs33 = _143_v4;
          let _rhs34 = _234_v82;
          let _lhs29 = _144_v5;
          let _lhs30 = _module.__default.safeIndex(new BigNumber(483), new BigNumber((_144_v5).length));
          let _lhs31 = _232_v80;
          let _lhs32 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_232_v80).length));
          _lhs29[_lhs30] = _rhs32;
          _lhs31[_lhs32] = _rhs33;
          _234_v82 = _rhs34;
        } else {
          (_201_v53).f22 = (_207_v57)[_module.__default.safeIndex(new BigNumber((((_201_v53.f22) ? (_199_v51) : (_module.__default.fm5((_dafny.ZERO).minus(_module.__default.fm6(_147_globalState)), _142_v3, new BigNumber(109), _147_globalState)))).length), new BigNumber((_207_v57).length))];
          (_147_globalState).f20 = _142_v3;
          _149_v8 = _149_v8;
          let _235_v83;
          let _nw39 = Array((new BigNumber(29)).toNumber()).fill(_dafny.MultiSet.Empty);
          _235_v83 = _nw39;
          let _236_v84;
          let _nw40 = new _module.C1();
          _nw40.__ctor(_235_v83, false);
          _236_v84 = _nw40;
          let _237_v85;
          _237_v85 = _dafny.Map.Empty.slice().updateUnsafe(_201_v53.f22,_dafny.MultiSet.FromArray(_dafny.Seq.of(false)));
          let _238_v86;
          _238_v86 = _dafny.Map.Empty.slice().updateUnsafe(_236_v84,_module.__default.fm12(_142_v3, (((_237_v85).contains(false)) ? ((_237_v85).get(false)) : (_dafny.MultiSet.fromElements(_201_v53.f22))), _201_v53.f22, _183_v39, _147_globalState));
          _238_v86 = (_238_v86).update(_236_v84, (_207_v57)[_module.__default.safeIndex(_module.__default.fm6(_147_globalState), new BigNumber((_207_v57).length))]);
          _182_v38 = _dafny.MultiSet.fromElements(!(_201_v53.f22));
        }
      } else {
        let _239___mcc_h2 = (_source3).cf2;
        let _240_cf2 = _239___mcc_h2;
        let _index31 = _module.__default.safeIndex(new BigNumber(775), new BigNumber((_145_v6).length));
        (_145_v6)[_index31] = _142_v3;
        let _index32 = _module.__default.safeIndex(new BigNumber(775), new BigNumber((_145_v6).length));
        (_145_v6)[_index32] = _142_v3;
        (_147_globalState).f20 = _module.__default.fm6(_147_globalState);
        let _241_v87;
        _241_v87 = _dafny.Seq.of((_145_v6)[_module.__default.safeIndex(new BigNumber(775), new BigNumber((_145_v6).length))], _142_v3);
        let _rhs35 = (_142_v3).plus(_142_v3);
        let _rhs36 = (_241_v87)[_module.__default.safeIndex(_142_v3, new BigNumber((_241_v87).length))];
        let _lhs33 = _147_globalState;
        let _lhs34 = _147_globalState;
        _lhs33.f20 = _rhs35;
        _lhs34.f3 = _rhs36;
        let _242_v88;
        _242_v88 = _dafny.Set.fromElements((_145_v6)[_module.__default.safeIndex(new BigNumber(775), new BigNumber((_145_v6).length))]);
        let _243_v89;
        _243_v89 = _dafny.Seq.of(_242_v88);
        let _244_v90;
        _244_v90 = (_243_v89)[_module.__default.safeIndex(_142_v3, new BigNumber((_243_v89).length))];
        let _245_v91;
        _245_v91 = _dafny.Seq.of(_241_v87, _241_v87);
        let _246_v92;
        let _nw41 = new _module.C4();
        _nw41.__ctor(_dafny.Set.fromElements((_145_v6)[_module.__default.safeIndex(new BigNumber(775), new BigNumber((_145_v6).length))]), _149_v8, _245_v91, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-37)), ((_247_v4) => function (_248_i5) {
          return _247_v4;
        })(_143_v4)), _dafny.Seq.contains(_241_v87, _module.__default.fm6(_147_globalState)));
        _246_v92 = _nw41;
        let _rhs37 = _246_v92.f22;
        let _rhs38 = ((!(_149_v8) || (_149_v8)) ? (_246_v92) : (_246_v92));
        let _rhs39 = _142_v3;
        let _lhs35 = _147_globalState;
        _149_v8 = _rhs37;
        _246_v92 = _rhs38;
        _lhs35.f3 = _rhs39;
      }
      let _249_v93;
      _249_v93 = _dafny.MultiSet.fromElements(_144_v5, _144_v5);
      let _250_v94;
      _250_v94 = _dafny.Map.Empty.slice().updateUnsafe(_142_v3,_149_v8);
      let _251_v95;
      _251_v95 = _dafny.Seq.of(_149_v8, _149_v8);
      let _252_v96;
      let _nw42 = Array((new BigNumber(25)).toNumber());
      _nw42[(_dafny.ZERO).toNumber()] = _249_v93;
      _nw42[(_dafny.ONE).toNumber()] = _249_v93;
      _nw42[(new BigNumber(2)).toNumber()] = _249_v93;
      _nw42[(new BigNumber(3)).toNumber()] = (_249_v93).Intersect(((_module.D7.create_DC15(_249_v93)).dtor_cf25).update(_144_v5, _module.__default.abs(_142_v3)));
      _nw42[(new BigNumber(4)).toNumber()] = _249_v93;
      _nw42[(new BigNumber(5)).toNumber()] = _dafny.MultiSet.fromElements(_144_v5, _144_v5);
      _nw42[(new BigNumber(6)).toNumber()] = (_249_v93).Difference(_249_v93);
      _nw42[(new BigNumber(7)).toNumber()] = _249_v93;
      _nw42[(new BigNumber(8)).toNumber()] = _249_v93;
      _nw42[(new BigNumber(9)).toNumber()] = _249_v93;
      _nw42[(new BigNumber(10)).toNumber()] = _249_v93;
      _nw42[(new BigNumber(11)).toNumber()] = (_249_v93).Union(_249_v93);
      _nw42[(new BigNumber(12)).toNumber()] = _249_v93;
      _nw42[(new BigNumber(13)).toNumber()] = (((((_250_v94).contains(_142_v3)) ? ((_250_v94).get(_142_v3)) : (!(_149_v8)))) ? (_249_v93) : (_249_v93));
      _nw42[(new BigNumber(14)).toNumber()] = _249_v93;
      _nw42[(new BigNumber(15)).toNumber()] = _dafny.MultiSet.fromElements(_144_v5);
      _nw42[(new BigNumber(16)).toNumber()] = _249_v93;
      _nw42[(new BigNumber(17)).toNumber()] = ((_module.__default.fm12(new BigNumber(572), _dafny.MultiSet.fromElements(_149_v8, _149_v8, _149_v8, _149_v8, _149_v8), (_251_v95)[_module.__default.safeIndex(new BigNumber((_140_v1).length), new BigNumber((_251_v95).length))], _dafny.MultiSet.fromElements(_182_v38), _147_globalState)) ? (_249_v93) : (_249_v93));
      _nw42[(new BigNumber(18)).toNumber()] = (_module.D7.create_DC15(_249_v93)).dtor_cf25;
      _nw42[(new BigNumber(19)).toNumber()] = _249_v93;
      _nw42[(new BigNumber(20)).toNumber()] = _249_v93;
      _nw42[(new BigNumber(21)).toNumber()] = (_249_v93).update(_144_v5, _module.__default.abs(_142_v3));
      _nw42[(new BigNumber(22)).toNumber()] = _249_v93;
      _nw42[(new BigNumber(23)).toNumber()] = (_249_v93).Union(_249_v93);
      _nw42[(new BigNumber(24)).toNumber()] = (_dafny.MultiSet.fromElements(_144_v5)).Union(_249_v93);
      _252_v96 = _nw42;
      let _index33 = _module.__default.safeIndex(new BigNumber(543), new BigNumber((_252_v96).length));
      (_252_v96)[_index33] = (_dafny.MultiSet.fromElements(_144_v5, _144_v5)).Union(_249_v93);
      let _index34 = _module.__default.safeIndex(new BigNumber(543), new BigNumber((_252_v96).length));
      (_252_v96)[_index34] = (((_249_v93).update(_144_v5, _module.__default.abs(new BigNumber((_250_v94).length)))).Union(_249_v93)).Difference(_249_v93);
      let _253_v97;
      _253_v97 = _dafny.Set.fromElements(_142_v3, _142_v3, new BigNumber((_dafny.Seq.UnicodeFromString("ly")).length));
      let _254_v98;
      _254_v98 = _253_v97;
      let _255_v99;
      _255_v99 = _dafny.Seq.of(_142_v3);
      let _256_v100;
      let _nw43 = new _module.C4();
      _nw43.__ctor(_254_v98, _149_v8, _dafny.Seq.of(_dafny.Seq.of(_142_v3), _255_v99, _dafny.Seq.update(_dafny.Seq.of(_142_v3), _module.__default.safeIndex(_142_v3, new BigNumber((_dafny.Seq.of(_142_v3)).length)), _142_v3)), _140_v1, true);
      _256_v100 = _nw43;
      let _257_v101;
      _257_v101 = _dafny.Set.fromElements(_256_v100, _256_v100, _256_v100);
      let _258_v102;
      _258_v102 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm12(new BigNumber((_140_v1).length), _dafny.MultiSet.fromElements(_149_v8), _149_v8, _183_v39, _147_globalState),_257_v101);
      let _hi1 = (_dafny.ZERO).minus(new BigNumber((_258_v102).length));
      for (let _259_i6 = _142_v3; _259_i6.isLessThan(_hi1); _259_i6 = _259_i6.plus(_dafny.ONE)) {
        let _260_v103;
        _260_v103 = _dafny.Seq.of(_255_v99);
        let _261_v104;
        _261_v104 = _dafny.Seq.of((_260_v103)[_module.__default.safeIndex(_142_v3, new BigNumber((_260_v103).length))]);
        let _262_v105;
        let _nw44 = new _module.C2();
        _nw44.__ctor(_260_v103, _dafny.Seq.Concat(_140_v1, _dafny.Seq.update(_140_v1, _module.__default.safeIndex(_259_i6, new BigNumber((_140_v1).length)), new _dafny.CodePoint('u'.codePointAt(0)))), (_142_v3).isLessThanOrEqualTo(_142_v3));
        _262_v105 = _nw44;
        let _263_v106;
        _263_v106 = _dafny.Seq.of(_250_v94);
        let _264_v107;
        let _out3;
        _out3 = (_262_v105).m1((_256_v100).f30, _dafny.Map.Empty.slice().updateUnsafe((_263_v106)[_module.__default.safeIndex(new BigNumber(840), new BigNumber((_263_v106).length))],_144_v5), _147_globalState);
        _264_v107 = _out3;
        _149_v8 = (_256_v100).f30;
        _255_v99 = _255_v99;
      }
      _149_v8 = (_256_v100).f30;
      let _265_v108;
      let _nw45 = Array((new BigNumber(8)).toNumber());
      _nw45[(_dafny.ZERO).toNumber()] = _143_v4;
      _nw45[(_dafny.ONE).toNumber()] = ((_149_v8) ? (_143_v4) : (_143_v4));
      _nw45[(new BigNumber(2)).toNumber()] = _143_v4;
      _nw45[(new BigNumber(3)).toNumber()] = _143_v4;
      _nw45[(new BigNumber(4)).toNumber()] = _143_v4;
      _nw45[(new BigNumber(5)).toNumber()] = _143_v4;
      _nw45[(new BigNumber(6)).toNumber()] = _143_v4;
      _nw45[(new BigNumber(7)).toNumber()] = _143_v4;
      _265_v108 = _nw45;
      _265_v108 = _265_v108;
      _255_v99 = _255_v99;
      let _266_v109;
      _266_v109 = _module.D3.create_DC7(_182_v38);
      let _pat_let_tv14 = _182_v38;
      let _267_v110;
      let _nw46 = Array((new BigNumber(25)).toNumber());
      _nw46[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.fromElements(false);
      _nw46[(_dafny.ONE).toNumber()] = _module.__default.fm17(_142_v3, (_256_v100).f30, _147_globalState);
      _nw46[(new BigNumber(2)).toNumber()] = _182_v38;
      _nw46[(new BigNumber(3)).toNumber()] = _182_v38;
      _nw46[(new BigNumber(4)).toNumber()] = _182_v38;
      _nw46[(new BigNumber(5)).toNumber()] = (function (_pat_let0_0) {
        return function (_268_dt__update__tmp_h1) {
          return function (_pat_let1_0) {
            return function (_269_dt__update_hcf12_h0) {
              return _module.D3.create_DC7(_269_dt__update_hcf12_h0);
            }(_pat_let1_0);
          }(_pat_let_tv14);
        }(_pat_let0_0);
      }(_266_v109)).dtor_cf12;
      _nw46[(new BigNumber(6)).toNumber()] = _182_v38;
      _nw46[(new BigNumber(7)).toNumber()] = _182_v38;
      _nw46[(new BigNumber(8)).toNumber()] = _182_v38;
      _nw46[(new BigNumber(9)).toNumber()] = _dafny.MultiSet.fromElements(_149_v8, (_256_v100).f30);
      _nw46[(new BigNumber(10)).toNumber()] = _182_v38;
      _nw46[(new BigNumber(11)).toNumber()] = _182_v38;
      _nw46[(new BigNumber(12)).toNumber()] = _182_v38;
      _nw46[(new BigNumber(13)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.update(_251_v95, _module.__default.safeIndex(_142_v3, new BigNumber((_251_v95).length)), (_256_v100).fm1(_140_v1, _149_v8, (_256_v100).f30, _147_globalState)));
      _nw46[(new BigNumber(14)).toNumber()] = _dafny.MultiSet.fromElements(!((_256_v100).f30));
      _nw46[(new BigNumber(15)).toNumber()] = _182_v38;
      _nw46[(new BigNumber(16)).toNumber()] = _182_v38;
      _nw46[(new BigNumber(17)).toNumber()] = _182_v38;
      _nw46[(new BigNumber(18)).toNumber()] = _dafny.MultiSet.fromElements((_256_v100).f30);
      _nw46[(new BigNumber(19)).toNumber()] = _182_v38;
      _nw46[(new BigNumber(20)).toNumber()] = (_182_v38).update(_149_v8, _module.__default.abs((_dafny.ZERO).minus(_142_v3)));
      _nw46[(new BigNumber(21)).toNumber()] = _182_v38;
      _nw46[(new BigNumber(22)).toNumber()] = _182_v38;
      _nw46[(new BigNumber(23)).toNumber()] = _module.__default.fm17(new BigNumber(889), _149_v8, _147_globalState);
      _nw46[(new BigNumber(24)).toNumber()] = _dafny.MultiSet.FromArray(_251_v95);
      _267_v110 = _nw46;
      let _270_v111;
      let _nw47 = new _module.C1();
      _nw47.__ctor(_267_v110, (_256_v100).f30);
      _270_v111 = _nw47;
      let _271_v112;
      _271_v112 = _dafny.Map.Empty.slice().updateUnsafe(_142_v3,_142_v3);
      let _272_v113;
      _272_v113 = _dafny.Map.Empty.slice().updateUnsafe(_182_v38,_142_v3);
      let _273_v114;
      _273_v114 = _module.D1.create_DC5(false, _dafny.Map.Empty.slice().updateUnsafe((_270_v111).f25,_142_v3), _142_v3, _146_v7);
      _142_v3 = _module.__default.safeModuloInt(new BigNumber((_271_v112).length), (_module.D3.create_DC8(_142_v3, _272_v113, _273_v114, _149_v8)).dtor_cf13);
      let _hi2 = _142_v3;
      for (let _274_i7 = _module.__default.safeDivisionInt(new BigNumber(86), _142_v3); _274_i7.isLessThan(_hi2); _274_i7 = _274_i7.plus(_dafny.ONE)) {
        let _275_v115;
        _275_v115 = _dafny.Seq.of(_module.__default.fm11((_270_v111).f25, _147_globalState), _140_v1, _module.__default.fm11(false, _147_globalState));
        _140_v1 = _dafny.Seq.Concat((_275_v115)[_module.__default.safeIndex(_142_v3, new BigNumber((_275_v115).length))], _140_v1);
        (_147_globalState).f20 = new BigNumber(761);
        let _index35 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_265_v108).length));
        (_265_v108)[_index35] = _143_v4;
        let _276_v116;
        _276_v116 = _dafny.MultiSet.fromElements(_140_v1);
        let _index36 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_265_v108).length));
        (_265_v108)[_index36] = (_140_v1)[_module.__default.safeIndex(new BigNumber(((_dafny.MultiSet.fromElements(_140_v1)).Intersect(_276_v116)).cardinality()), new BigNumber((_140_v1).length))];
        _142_v3 = new BigNumber(((_182_v38).update((_270_v111).f25, _module.__default.abs((_dafny.ZERO).minus((_142_v3).plus(new BigNumber(984)))))).cardinality());
      }
      let _277_v117;
      let _nw48 = new _module.C3();
      _nw48.__ctor((_270_v111).f25, _149_v8, _140_v1, (_270_v111).f25);
      _277_v117 = _nw48;
      _142_v3 = _142_v3;
      (_147_globalState).f12 = (((_277_v117).f27) ? (_module.__default.fm4(_277_v117.f28, _182_v38, _147_globalState)) : (new _dafny.CodePoint('q'.codePointAt(0))));
      (_277_v117).f28 = (_270_v111).f25;
      process.stdout.write((_140_v1).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_141_v2).equals(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("truxk")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_142_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_143_v4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_v6)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_145_v6)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_146_v7).equals(_dafny.MultiSet.fromElements(new BigNumber(-610)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_147_globalState).f1).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("truxk"),_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_147_globalState.f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_147_globalState.f6).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_147_globalState.f8).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_147_globalState).f9).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_147_globalState.f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_globalState).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_globalState).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_globalState).f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_147_globalState).f17)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_147_globalState).f17)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_globalState).f18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_globalState.f19).equals(_dafny.MultiSet.fromElements(new BigNumber(-610)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_147_globalState.f20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_149_v8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_150_i0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_v38).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_183_v39).equals(_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(false), _dafny.MultiSet.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_249_v93).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_250_v94).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-610),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_251_v95, _dafny.Seq.of(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_252_v96)[_dafny.ZERO]).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_252_v96)[_dafny.ONE]).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_252_v96)[new BigNumber(2)]).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_252_v96)[new BigNumber(3)]).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_252_v96)[new BigNumber(4)]).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_252_v96)[new BigNumber(5)]).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_252_v96)[new BigNumber(6)]).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_252_v96)[new BigNumber(7)]).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_252_v96)[new BigNumber(8)]).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_252_v96)[new BigNumber(9)]).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_252_v96)[new BigNumber(10)]).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_252_v96)[new BigNumber(11)]).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_252_v96)[new BigNumber(12)]).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_252_v96)[new BigNumber(13)]).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_252_v96)[new BigNumber(14)]).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_252_v96)[new BigNumber(15)]).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_252_v96)[new BigNumber(16)]).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_252_v96)[new BigNumber(17)]).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_252_v96)[new BigNumber(18)]).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_252_v96)[new BigNumber(19)]).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_252_v96)[new BigNumber(20)]).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_252_v96)[new BigNumber(21)]).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_252_v96)[new BigNumber(22)]).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_252_v96)[new BigNumber(23)]).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_252_v96)[new BigNumber(24)]).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_253_v97).equals(_dafny.Set.fromElements(new BigNumber(-610), new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_254_v98)).equals(_dafny.Set.fromElements(new BigNumber(-610), new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_255_v99, _dafny.Seq.of(new BigNumber(-610)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_256_v100.f29)).equals(_dafny.Set.fromElements(new BigNumber(-610), new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_256_v100).f30));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_257_v101).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_258_v102).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_265_v108)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_265_v108)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_265_v108)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_265_v108)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_265_v108)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_265_v108)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_265_v108)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_265_v108)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_266_v109).dtor_cf12).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_267_v110)[_dafny.ZERO]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_267_v110)[_dafny.ONE]).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_267_v110)[new BigNumber(2)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_267_v110)[new BigNumber(3)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_267_v110)[new BigNumber(4)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_267_v110)[new BigNumber(5)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_267_v110)[new BigNumber(6)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_267_v110)[new BigNumber(7)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_267_v110)[new BigNumber(8)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_267_v110)[new BigNumber(9)]).equals(_dafny.MultiSet.fromElements(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_267_v110)[new BigNumber(10)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_267_v110)[new BigNumber(11)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_267_v110)[new BigNumber(12)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_267_v110)[new BigNumber(13)]).equals(_dafny.MultiSet.fromElements(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_267_v110)[new BigNumber(14)]).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_267_v110)[new BigNumber(15)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_267_v110)[new BigNumber(16)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_267_v110)[new BigNumber(17)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_267_v110)[new BigNumber(18)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_267_v110)[new BigNumber(19)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_267_v110)[new BigNumber(20)]).equals(_dafny.MultiSet.fromElements(false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_267_v110)[new BigNumber(21)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_267_v110)[new BigNumber(22)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_267_v110)[new BigNumber(23)]).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_267_v110)[new BigNumber(24)]).equals(_dafny.MultiSet.fromElements(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_270_v111.f24)[_dafny.ZERO]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_270_v111.f24)[_dafny.ONE]).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_270_v111.f24)[new BigNumber(2)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_270_v111.f24)[new BigNumber(3)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_270_v111.f24)[new BigNumber(4)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_270_v111.f24)[new BigNumber(5)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_270_v111.f24)[new BigNumber(6)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_270_v111.f24)[new BigNumber(7)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_270_v111.f24)[new BigNumber(8)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_270_v111.f24)[new BigNumber(9)]).equals(_dafny.MultiSet.fromElements(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_270_v111.f24)[new BigNumber(10)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_270_v111.f24)[new BigNumber(11)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_270_v111.f24)[new BigNumber(12)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_270_v111.f24)[new BigNumber(13)]).equals(_dafny.MultiSet.fromElements(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_270_v111.f24)[new BigNumber(14)]).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_270_v111.f24)[new BigNumber(15)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_270_v111.f24)[new BigNumber(16)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_270_v111.f24)[new BigNumber(17)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_270_v111.f24)[new BigNumber(18)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_270_v111.f24)[new BigNumber(19)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_270_v111.f24)[new BigNumber(20)]).equals(_dafny.MultiSet.fromElements(false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_270_v111.f24)[new BigNumber(21)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_270_v111.f24)[new BigNumber(22)]).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_270_v111.f24)[new BigNumber(23)]).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_270_v111.f24)[new BigNumber(24)]).equals(_dafny.MultiSet.fromElements(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_270_v111).f25));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_271_v112).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-610),new BigNumber(-610)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_272_v113).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(false),new BigNumber(-610)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_273_v114).dtor_cf7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_273_v114).dtor_cf8).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-610)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_273_v114).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_273_v114).dtor_cf10).equals(_dafny.MultiSet.fromElements(new BigNumber(-610)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_277_v117).f27));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_277_v117.f28));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(1);
      $dt.cf0 = cf0;
      return $dt;
    }
    static create_DC2(cf2) {
      let $dt = new D0(2);
      $dt.cf2 = cf2;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC0() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf2() { return this.cf2; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf2) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf1 === other.cf1;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf0 === other.cf0;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf2, other.cf2);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(false);
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
    static create_DC4(cf4, cf5, cf6) {
      let $dt = new D1(0);
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC5(cf7, cf8, cf9, cf10) {
      let $dt = new D1(1);
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC3(cf3) {
      let $dt = new D1(2);
      $dt.cf3 = cf3;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf3() { return this.cf3; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf3) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf4, other.cf4) && _dafny.areEqual(this.cf5, other.cf5) && _dafny.areEqual(this.cf6, other.cf6);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf7 === other.cf7 && _dafny.areEqual(this.cf8, other.cf8) && _dafny.areEqual(this.cf9, other.cf9) && _dafny.areEqual(this.cf10, other.cf10);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf3, other.cf3);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(new _dafny.CodePoint('D'.codePointAt(0)), _dafny.Map.Empty, _dafny.Seq.of());
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
    static create_DC6(cf11) {
      let $dt = new D2(0);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf11, other.cf11);
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
    static create_DC8(cf13, cf14, cf15, cf16) {
      let $dt = new D3(0);
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC9(cf17, cf18, cf19) {
      let $dt = new D3(1);
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      return $dt;
    }
    static create_DC7(cf12) {
      let $dt = new D3(2);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get is_DC7() { return this.$tag === 2; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC7" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf13, other.cf13) && _dafny.areEqual(this.cf14, other.cf14) && _dafny.areEqual(this.cf15, other.cf15) && this.cf16 === other.cf16;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf17, other.cf17) && _dafny.areEqual(this.cf18, other.cf18) && _dafny.areEqual(this.cf19, other.cf19);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf12, other.cf12);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC8(_dafny.ZERO, _dafny.Map.Empty, _module.D1.Default(), false);
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
    static create_DC11(cf21) {
      let $dt = new D4(0);
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC10(cf20) {
      let $dt = new D4(1);
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC12(cf22) {
      let $dt = new D4(2);
      $dt.cf22 = cf22;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get is_DC12() { return this.$tag === 2; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf22() { return this.cf22; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf22) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf21, other.cf21);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf20 === other.cf20;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf22, other.cf22);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC11(_dafny.Set.Empty);
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
    static create_DC13(cf23) {
      let $dt = new D5(0);
      $dt.cf23 = cf23;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get dtor_cf23() { return this.cf23; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf23) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf23 === other.cf23;
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
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC14(cf24) {
      let $dt = new D6(0);
      $dt.cf24 = cf24;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get dtor_cf24() { return this.cf24; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC14" + "(" + _dafny.toString(this.cf24) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf24 === other.cf24;
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
          return D6.Default();
        }
      };
    }
  }

  $module.D7 = class D7 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC16(cf26, cf27, cf28, cf29) {
      let $dt = new D7(0);
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC15(cf25) {
      let $dt = new D7(1);
      $dt.cf25 = cf25;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf25() { return this.cf25; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC16" + "(" + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC15" + "(" + _dafny.toString(this.cf25) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf26 === other.cf26 && this.cf27 === other.cf27 && this.cf28 === other.cf28 && this.cf29 === other.cf29;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf25, other.cf25);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC16(false, false, null, false);
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
    static create_DC17(cf30) {
      let $dt = new D8(0);
      $dt.cf30 = cf30;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get dtor_cf30() { return this.cf30; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC17" + "(" + _dafny.toString(this.cf30) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf30, other.cf30);
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
          return D8.Default();
        }
      };
    }
  }

  $module.D9 = class D9 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC19(cf32, cf33, cf34, cf35, cf36) {
      let $dt = new D9(0);
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      return $dt;
    }
    static create_DC20(cf37, cf38, cf39) {
      let $dt = new D9(1);
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC21(cf40, cf41) {
      let $dt = new D9(2);
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC18(cf31) {
      let $dt = new D9(3);
      $dt.cf31 = cf31;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get is_DC20() { return this.$tag === 1; }
    get is_DC21() { return this.$tag === 2; }
    get is_DC18() { return this.$tag === 3; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf31() { return this.cf31; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC19" + "(" + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC20" + "(" + this.cf37.toVerbatimString(true) + ", " + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC21" + "(" + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 3) {
        return "D9.DC18" + "(" + _dafny.toString(this.cf31) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf32, other.cf32) && _dafny.areEqual(this.cf33, other.cf33) && this.cf34 === other.cf34 && _dafny.areEqual(this.cf35, other.cf35) && _dafny.areEqual(this.cf36, other.cf36);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf37, other.cf37) && this.cf38 === other.cf38 && _dafny.areEqual(this.cf39, other.cf39);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf40 === other.cf40 && this.cf41 === other.cf41;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf31, other.cf31);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC19(_dafny.ZERO, _dafny.ZERO, false, _dafny.ZERO, _dafny.Set.Empty);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D9.Default();
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
      this.f3 = _dafny.ZERO;
      this.f6 = _dafny.Seq.UnicodeFromString("");
      this.f8 = _dafny.Seq.UnicodeFromString("");
      this.f12 = new _dafny.CodePoint('D'.codePointAt(0));
      this.f19 = _dafny.MultiSet.Empty;
      this.f20 = _dafny.ZERO;
      this._f0 = false;
      this._f1 = _dafny.Map.Empty;
      this._f2 = _dafny.ZERO;
      this._f4 = _dafny.ZERO;
      this._f5 = false;
      this._f7 = false;
      this._f9 = _dafny.Seq.UnicodeFromString("");
      this._f10 = _dafny.ZERO;
      this._f11 = false;
      this._f13 = _dafny.ZERO;
      this._f14 = _dafny.ZERO;
      this._f15 = _dafny.ZERO;
      this._f16 = [];
      this._f17 = [];
      this._f18 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this).f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this).f6 = f6;
      (_this)._f7 = f7;
      (_this).f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this).f12 = f12;
      (_this)._f13 = f13;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      (_this).f19 = f19;
      (_this).f20 = f20;
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
    get f16() {
      let _this = this;
      return _this._f16;
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

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this.f26 = _dafny.Set.Empty;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f26) {
      let _this = this;
      (_this).f26 = f26;
      return;
    }
    fm7(p0, globalState) {
      let _this = this;
      return !(false) || ((false) && (true));
    };
    fm8(globalState) {
      let _this = this;
      return (_module.D1.create_DC3(_dafny.Set.fromElements(_dafny.Set.fromElements(new _dafny.CodePoint('r'.codePointAt(0)))))).dtor_cf3;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this.f24 = [];
      this._f25 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f24, f25) {
      let _this = this;
      (_this).f24 = f24;
      (_this)._f25 = f25;
      return;
    }
    m3(p0, p1, p2, globalState) {
      let _this = this;
      let _278_v0;
      _278_v0 = true;
      let _279_v1;
      _279_v1 = new BigNumber(976);
      let _280_v2;
      _280_v2 = _dafny.Set.fromElements(_279_v1, _279_v1, _279_v1);
      _278_v0 = (p1) || (!(_280_v2).equals(_280_v2));
      let _281_v3;
      let _nw49 = Array((new BigNumber(21)).toNumber()).fill(false);
      _281_v3 = _nw49;
      let _282_v4;
      _282_v4 = _module.D0.create_DC0(_281_v3);
      let _pat_let_tv15 = _281_v3;
      _282_v4 = function (_pat_let2_0) {
        return function (_283_dt__update__tmp_h0) {
          return function (_pat_let3_0) {
            return function (_284_dt__update_hcf0_h0) {
              return _module.D0.create_DC0(_284_dt__update_hcf0_h0);
            }(_pat_let3_0);
          }(_pat_let_tv15);
        }(_pat_let2_0);
      }(_282_v4);
      let _285_v5;
      _285_v5 = _dafny.Seq.of(p0, (_this).f25);
      let _286_v6;
      _286_v6 = _dafny.MultiSet.fromElements(!(!(_278_v0)), (_this).f25, p0, (_285_v5)[_module.__default.safeIndex(_module.__default.fm6(globalState), new BigNumber((_285_v5).length))], (_this).f25);
      let _287_v7;
      _287_v7 = _dafny.Seq.UnicodeFromString("ekkjfqs");
      let _index37 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_281_v3).length));
      (_281_v3)[_index37] = _dafny.Seq.IsProperPrefixOf(_287_v7, _287_v7);
      let _288_v8;
      _288_v8 = _dafny.Seq.of(_285_v5, _285_v5, _285_v5);
      let _index38 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_281_v3).length));
      let _rhs40 = _286_v6;
      let _rhs41 = _dafny.areEqual(_285_v5, (_288_v8)[_module.__default.safeIndex(new BigNumber((_287_v7).length), new BigNumber((_288_v8).length))]);
      let _rhs42 = p0;
      let _lhs36 = _281_v3;
      let _lhs37 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_281_v3).length));
      _286_v6 = _rhs40;
      _278_v0 = _rhs41;
      _lhs36[_lhs37] = _rhs42;
      let _289_i0;
      _289_i0 = _dafny.ZERO;
      L1: {
        while ((_281_v3)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_281_v3).length))]) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_289_i0)) {
              break L1;
            }
            _289_i0 = (_289_i0).plus(_dafny.ONE);
            let _290_v9;
            let _nw50 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Map.Empty);
            _290_v9 = _nw50;
            _290_v9 = _290_v9;
            _281_v3 = _281_v3;
            let _291_v10;
            _291_v10 = _module.D0.create_DC1((_this).f25);
            if ((_291_v10).dtor_cf1) {
              (globalState).f3 = (new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(_279_v1))).length)).plus(new BigNumber((_dafny.MultiSet.fromElements(_279_v1)).cardinality()));
              _278_v0 = p0;
              let _292_v11;
              _292_v11 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("u"), _dafny.Seq.UnicodeFromString("px"));
              let _293_v12;
              let _nw51 = new _module.C0();
              _nw51.__ctor(_292_v11);
              _293_v12 = _nw51;
              (globalState).f3 = new BigNumber((_dafny.Seq.Concat(_287_v7, _dafny.Seq.Create(_module.__default.abs(new BigNumber(687)), function (_294_i1) {
                return new _dafny.CodePoint('c'.codePointAt(0));
              }))).length);
              let _295_v13;
              _295_v13 = _dafny.Map.Empty.slice().updateUnsafe(p1,_279_v1);
              let _296_v14;
              _296_v14 = _dafny.Seq.of(_279_v1, new BigNumber((_module.__default.fm9(_278_v0, p2, globalState)).length));
              let _297_v15;
              _297_v15 = _dafny.Map.Empty.slice().updateUnsafe(_279_v1,(_dafny.MultiSet.FromArray(_296_v14)).update(_279_v1, _module.__default.abs(_279_v1)));
              let _298_v16;
              _298_v16 = _dafny.MultiSet.fromElements(new BigNumber((_285_v5).length));
              let _299_v17;
              _299_v17 = _module.D1.create_DC5((_291_v10).dtor_cf1, _295_v13, (_dafny.ZERO).minus(new BigNumber((_287_v7).length)), (((_297_v15).contains(_279_v1)) ? ((_297_v15).get(_279_v1)) : (_298_v16)));
              _299_v17 = _299_v17;
            } else {
              let _index39 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_281_v3).length));
              (_281_v3)[_index39] = !(p0) || (p0);
              _279_v1 = ((p1) ? (new BigNumber(515)) : (_279_v1));
              let _300_v18;
              _300_v18 = _dafny.Set.fromElements(_287_v7);
              let _301_v19;
              let _nw52 = new _module.C0();
              _nw52.__ctor(_300_v18);
              _301_v19 = _nw52;
              (globalState).f20 = _module.__default.safeModuloInt(new BigNumber(219), _279_v1);
              _278_v0 = _278_v0;
            }
            let _source4 = _module.__default.fm10(globalState);
            if (_source4.is_DC4) {
              let _302___mcc_h0 = (_source4).cf4;
              let _303___mcc_h1 = (_source4).cf5;
              let _304___mcc_h2 = (_source4).cf6;
              let _305_cf6 = _304___mcc_h2;
              let _306_cf5 = _303___mcc_h1;
              let _307_cf4 = _302___mcc_h0;
              let _308_v20;
              let _init3 = ((_309_v10, _310_p1) => function (_311_i2) {
                return function (_pat_let4_0) {
                  return function (_312_dt__update__tmp_h1) {
                    return function (_pat_let5_0) {
                      return function (_313_dt__update_hcf1_h0) {
                        return _module.D0.create_DC1(_313_dt__update_hcf1_h0);
                      }(_pat_let5_0);
                    }(_310_p1);
                  }(_pat_let4_0);
                }(_309_v10);
              })(_291_v10, p1);
              let _nw53 = Array((new BigNumber(20)).toNumber());
              for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw53.length); _i0_3++) {
                _nw53[_i0_3] = _init3(new BigNumber(_i0_3));
              }
              _308_v20 = _nw53;
              let _314_v21;
              _314_v21 = _dafny.Map.Empty.slice().updateUnsafe(_308_v20,_279_v1);
              let _315_v22;
              _315_v22 = _dafny.Map.Empty.slice().updateUnsafe((_281_v3)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_281_v3).length))],_314_v21);
              _315_v22 = (_315_v22).update((_this).f25, _314_v21);
              let _316_v23;
              let _nw54 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
              _316_v23 = _nw54;
              let _317_v24;
              _317_v24 = _dafny.Map.Empty.slice().updateUnsafe(false,_316_v23);
              _317_v24 = (_317_v24).update(p0, _316_v23);
              let _318_v25;
              let _init4 = ((_319_cf4) => function (_320_i3) {
                return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-865)), ((_321_cf4) => function (_322_i4) {
                  return _321_cf4;
                })(_319_cf4));
              })(_307_cf4);
              let _nw55 = Array((new BigNumber(8)).toNumber());
              for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw55.length); _i0_4++) {
                _nw55[_i0_4] = _init4(new BigNumber(_i0_4));
              }
              _318_v25 = _nw55;
              let _index40 = _module.__default.safeIndex(new BigNumber(96), new BigNumber((_318_v25).length));
              (_318_v25)[_index40] = _module.__default.fm11(false, globalState);
              let _index41 = _module.__default.safeIndex(new BigNumber(96), new BigNumber((_318_v25).length));
              (_318_v25)[_index41] = _287_v7;
              _279_v1 = (_dafny.ZERO).minus(((_279_v1).minus(_279_v1)).minus((_305_cf6)[_module.__default.safeIndex(_279_v1, new BigNumber((_305_cf6).length))]));
            } else if (_source4.is_DC5) {
              let _323___mcc_h3 = (_source4).cf7;
              let _324___mcc_h4 = (_source4).cf8;
              let _325___mcc_h5 = (_source4).cf9;
              let _326___mcc_h6 = (_source4).cf10;
              let _327_cf10 = _326___mcc_h6;
              let _328_cf9 = _325___mcc_h5;
              let _329_cf8 = _324___mcc_h4;
              let _330_cf7 = _323___mcc_h3;
              let _331_v26;
              _331_v26 = new _dafny.CodePoint('t'.codePointAt(0));
              _278_v0 = ((p1) ? (!_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("bdncgjj"), _331_v26)) : ((_281_v3)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_281_v3).length))]));
              let _index42 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_281_v3).length));
              (_281_v3)[_index42] = ((_280_v2).Difference(_280_v2)).IsProperSubsetOf(_280_v2);
              let _332_v27;
              _332_v27 = _dafny.Seq.of(_328_cf9, _328_cf9);
              let _333_v28;
              _333_v28 = _dafny.MultiSet.fromElements(_286_v6);
              let _334_v29;
              _334_v29 = _dafny.Map.Empty.slice().updateUnsafe(_279_v1,_328_cf9);
              let _index43 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_281_v3).length));
              (_281_v3)[_index43] = _module.__default.fm12(new BigNumber((_dafny.Seq.Concat(_332_v27, _332_v27)).length), ((_286_v6).update(_278_v0, _module.__default.abs(_328_cf9))).Intersect(_286_v6), (_281_v3)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_281_v3).length))], _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(_module.__default.fm12(_328_cf9, _286_v6, p2, _333_v28, globalState), p2), _dafny.MultiSet.FromArray(_285_v5), _dafny.MultiSet.fromElements((_285_v5)[_module.__default.safeIndex((_332_v27)[_module.__default.safeIndex(new BigNumber((_334_v29).length), new BigNumber((_332_v27).length))], new BigNumber((_285_v5).length))], _278_v0), _dafny.MultiSet.fromElements(p0, (_this).f25)), globalState);
              let _335_v30;
              _335_v30 = _dafny.Map.Empty.slice().updateUnsafe(_328_cf9,p2);
              let _336_v31;
              _336_v31 = _dafny.Map.Empty.slice().updateUnsafe(_327_cf10,_module.__default.fm11((((_335_v30).contains(_279_v1)) ? ((_335_v30).get(_279_v1)) : (_278_v0)), globalState));
              _336_v31 = (_dafny.Map.Empty.slice().updateUnsafe(_327_cf10,_dafny.Seq.Create(_module.__default.abs(new BigNumber(786)), ((_337_v26) => function (_338_i5) {
                return _337_v26;
              })(_331_v26)))).Merge(_336_v31);
            } else {
              let _339___mcc_h7 = (_source4).cf3;
              let _340_cf3 = _339___mcc_h7;
              _278_v0 = ((_279_v1).multipliedBy(_279_v1)).isLessThanOrEqualTo(_279_v1);
              (globalState).f20 = _module.__default.fm6(globalState);
              (globalState).f3 = _279_v1;
              let _341_v32;
              _341_v32 = _dafny.Set.fromElements(_287_v7);
              let _342_v33;
              let _nw56 = new _module.C0();
              _nw56.__ctor((_341_v32).Intersect(_341_v32));
              _342_v33 = _nw56;
            }
          }
        }
      }
      let _343_v34;
      _343_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_285_v5).length),new BigNumber((_287_v7).length));
      let _344_v35;
      _344_v35 = _dafny.Seq.of(_343_v34);
      let _index44 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_281_v3).length));
      (_281_v3)[_index44] = !(function () {
        let _coll10 = new _dafny.Set();
        for (const _compr_10 of ((_344_v35)[_module.__default.safeIndex(_279_v1, new BigNumber((_344_v35).length))]).Keys.Elements) {
          let _345_v36 = _compr_10;
          if (((_344_v35)[_module.__default.safeIndex(_279_v1, new BigNumber((_344_v35).length))]).contains(_345_v36)) {
            _coll10.add((_345_v36).minus(new BigNumber((_dafny.Set.fromElements((_module.D0.create_DC1(true)).dtor_cf1)).length)));
          }
        }
        return _coll10;
      }()).equals((_280_v2).Union(_280_v2));
      let _346_v37;
      _346_v37 = _dafny.Map.Empty.slice().updateUnsafe(p2,_279_v1);
      let _347_v38;
      _347_v38 = _dafny.MultiSet.fromElements(new BigNumber((_280_v2).length));
      let _348_v39;
      _348_v39 = _module.D1.create_DC5((_281_v3)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_281_v3).length))], _346_v37, _279_v1, _347_v38);
      let _pat_let_tv16 = _281_v3;
      let _pat_let_tv17 = _281_v3;
      let _pat_let_tv18 = _279_v1;
      let _pat_let_tv19 = _279_v1;
      (globalState).f3 = function (_source5) {
        if (_source5.is_DC4) {
          let _349___mcc_h8 = (_source5).cf4;
          let _350___mcc_h9 = (_source5).cf5;
          let _351___mcc_h10 = (_source5).cf6;
          let _352_cf6 = _351___mcc_h10;
          let _353_cf5 = _350___mcc_h9;
          let _354_cf4 = _349___mcc_h8;
          return new BigNumber(((((_pat_let_tv17)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_pat_let_tv16).length))]) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(478)), ((_355_cf6) => function (_356_i6) {
            return _355_cf6;
          })(_352_cf6))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(437)), ((_357_cf4) => function (_358_i7) {
            return _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(_357_cf4)).length));
          })(_354_cf4))))).length);
        } else if (_source5.is_DC5) {
          let _359___mcc_h11 = (_source5).cf7;
          let _360___mcc_h12 = (_source5).cf8;
          let _361___mcc_h13 = (_source5).cf9;
          let _362___mcc_h14 = (_source5).cf10;
          let _363_cf10 = _362___mcc_h14;
          let _364_cf9 = _361___mcc_h13;
          let _365_cf8 = _360___mcc_h12;
          let _366_cf7 = _359___mcc_h11;
          return _pat_let_tv18;
        } else {
          let _367___mcc_h15 = (_source5).cf3;
          let _368_cf3 = _367___mcc_h15;
          return (_pat_let_tv19).minus(new BigNumber(585));
        }
      }(_348_v39);
      return;
    }
    get f25() {
      let _this = this;
      return _this._f25;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f21 = _dafny.Seq.UnicodeFromString("");
      this._f22 = false;
      this._f23 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [_module.T2, _module.T1, _module.T0];
    }
    get f21() {
      let _this = this;
      return _this._f21;
    };
    set f21(value) {
      let _this = this;
      _this._f21 = value;
    };
    get f22() {
      let _this = this;
      return _this._f22;
    };
    set f22(value) {
      let _this = this;
      _this._f22 = value;
    };
    get f23() {
      let _this = this;
      return _this._f23;
    };
    set f23(value) {
      let _this = this;
      _this._f23 = value;
    };
    __ctor(f23, f21, f22) {
      let _this = this;
      (_this)._f23 = f23;
      (_this)._f21 = f21;
      (_this)._f22 = f22;
      return;
    }
    fm2(p0, globalState) {
      let _this = this;
      return !(((false) ? (_this.f22) : (_this.f22))) || (!(false));
    };
    fm0(globalState) {
      let _this = this;
      return _dafny.Set.fromElements(_this.f22, _this.f22, false, _this.f22, _this.f22);
    };
    fm1(p0, p1, p2, globalState) {
      let _this = this;
      return ((new BigNumber(99)).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(_this.f22)).cardinality()))).length))).isEqualTo(new BigNumber(((_dafny.Set.fromElements(_this.f22, _this.f22)).Union(_dafny.Set.fromElements(true))).length));
    };
    fm3(globalState) {
      let _this = this;
      return !((((_this.f22) ? (new BigNumber(-141)) : (new BigNumber(419)))).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.of(_this.f21)).length)));
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _369_v0;
      _369_v0 = new BigNumber(191);
      let _370_v1;
      _370_v1 = _dafny.Seq.of(_369_v0, _369_v0);
      let _371_v2;
      _371_v2 = _dafny.MultiSet.fromElements((_this.f23)[_module.__default.safeIndex(_369_v0, new BigNumber((_this.f23).length))], _370_v1);
      let _372_v3;
      _372_v3 = _dafny.Set.fromElements((((_371_v2).contains(_370_v1)) ? ((_371_v2).get(_370_v1)) : (_369_v0)));
      let _373_v5;
      _373_v5 = _dafny.Set.fromElements(_372_v3, _372_v3, _372_v3, (function () {
        let _coll11 = new _dafny.Set();
        for (const _compr_11 of _dafny.IntegerRange(new BigNumber(132), new BigNumber(989))) {
          let _374_v4 = _compr_11;
          if (((new BigNumber(132)).isLessThanOrEqualTo(_374_v4)) && ((_374_v4).isLessThan(new BigNumber(989)))) {
            _coll11.add((_374_v4).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_369_v0,p0)).length)));
          }
        }
        return _coll11;
      }()).Difference(_372_v3));
      _373_v5 = (((p0) ? (_dafny.Set.fromElements(_372_v3)) : (_373_v5))).Union(_373_v5);
      if (true) {
        (globalState).f20 = _369_v0;
        let _375_v6;
        _375_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_369_v0, new BigNumber((_dafny.Seq.of(_this.f22)).length), _369_v0))).cardinality()),_this.f21);
        let _376_v7;
        _376_v7 = new _dafny.CodePoint('u'.codePointAt(0));
        let _377_v8;
        _377_v8 = _dafny.Seq.of(_this.f21, (((_375_v6).contains(_369_v0)) ? ((_375_v6).get(_369_v0)) : (_this.f21)), _dafny.Seq.of(_376_v7));
        (_this).f21 = _dafny.Seq.update((_377_v8)[_module.__default.safeIndex(_369_v0, new BigNumber((_377_v8).length))], _module.__default.safeIndex(_369_v0, new BigNumber(((_377_v8)[_module.__default.safeIndex(_369_v0, new BigNumber((_377_v8).length))]).length)), _376_v7);
        let _378_v9;
        let _init5 = function (_379_i0) {
          return true;
        };
        let _nw57 = Array((new BigNumber(19)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw57.length); _i0_5++) {
          _nw57[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _378_v9 = _nw57;
        let _380_v10;
        _380_v10 = _module.D0.create_DC0(_378_v9);
        let _381_v11;
        _381_v11 = _dafny.Seq.of(_378_v9);
        let _382_v12;
        let _nw58 = Array((new BigNumber(6)).toNumber());
        _nw58[(_dafny.ZERO).toNumber()] = _378_v9;
        _nw58[(_dafny.ONE).toNumber()] = _378_v9;
        _nw58[(new BigNumber(2)).toNumber()] = _378_v9;
        _nw58[(new BigNumber(3)).toNumber()] = (_380_v10).dtor_cf0;
        _nw58[(new BigNumber(4)).toNumber()] = _378_v9;
        _nw58[(new BigNumber(5)).toNumber()] = (_381_v11)[_module.__default.safeIndex(_369_v0, new BigNumber((_381_v11).length))];
        _382_v12 = _nw58;
        let _383_v13;
        _383_v13 = _module.D0.create_DC1(_this.f22);
        let _384_v14;
        _384_v14 = _dafny.MultiSet.fromElements(_this.f22);
        let _385_v15;
        _385_v15 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm4(_this.f22, _384_v14, globalState),_382_v12);
        _382_v12 = (((function (_pat_let6_0) {
          return function (_386_dt__update__tmp_h0) {
            return function (_pat_let7_0) {
              return function (_387_dt__update_hcf1_h0) {
                return _module.D0.create_DC1(_387_dt__update_hcf1_h0);
              }(_pat_let7_0);
            }(false);
          }(_pat_let6_0);
        }(_383_v13)).dtor_cf1) ? ((((_385_v15).contains(new _dafny.CodePoint('x'.codePointAt(0)))) ? ((_385_v15).get(new _dafny.CodePoint('x'.codePointAt(0)))) : (_382_v12))) : (_382_v12));
        let _index45 = _module.__default.safeIndex(new BigNumber(623), new BigNumber((_382_v12).length));
        (_382_v12)[_index45] = _378_v9;
        let _index46 = _module.__default.safeIndex(new BigNumber(623), new BigNumber((_382_v12).length));
        (_382_v12)[_index46] = _378_v9;
        let _388_v16;
        _388_v16 = _dafny.MultiSet.fromElements(_369_v0);
        r0 = (new BigNumber((_388_v16).cardinality())).multipliedBy((_370_v1)[_module.__default.safeIndex(new BigNumber(849), new BigNumber((_370_v1).length))]);
      } else {
        let _389_v17;
        _389_v17 = _dafny.Map.Empty.slice().updateUnsafe(p0,_369_v0);
        let _390_v18;
        _390_v18 = _dafny.Map.Empty.slice().updateUnsafe(_389_v17,!(_this.f22) || (_this.f22));
        _390_v18 = _dafny.Map.Empty.slice().updateUnsafe(_389_v17,true);
        (_this).f22 = _this.f22;
        let _391_v19;
        let _nw59 = Array((new BigNumber(16)).toNumber()).fill(false);
        _391_v19 = _nw59;
        let _392_v20;
        let _out4;
        _out4 = _module.__default.m0(_391_v19, (true) && (p0), globalState);
        _392_v20 = _out4;
        (_this).f22 = _392_v20;
        let _393_v21;
        _393_v21 = _dafny.Seq.of(_372_v3, _372_v3);
        (globalState).f20 = new BigNumber((_dafny.Seq.Concat(_393_v21, _dafny.Seq.update(_393_v21, _module.__default.safeIndex(new BigNumber(182), new BigNumber((_393_v21).length)), _dafny.Set.fromElements(_369_v0, (_dafny.ZERO).minus(_369_v0))))).length);
      }
      let _394_v22;
      _394_v22 = _module.D0.create_DC1(true);
      let _395_v23;
      _395_v23 = _module.D0.create_DC2(_394_v22);
      let _396_v24;
      _396_v24 = _dafny.Seq.of(_this.f21);
      let _397_v25;
      _397_v25 = _dafny.Seq.of(_this.f22, _this.f22, !(p0));
      let _398_v26;
      let _nw60 = Array((new BigNumber(15)).toNumber());
      _nw60[(_dafny.ZERO).toNumber()] = new BigNumber((_396_v24).length);
      _nw60[(_dafny.ONE).toNumber()] = _369_v0;
      _nw60[(new BigNumber(2)).toNumber()] = _369_v0;
      _nw60[(new BigNumber(3)).toNumber()] = _369_v0;
      _nw60[(new BigNumber(4)).toNumber()] = new BigNumber(426);
      _nw60[(new BigNumber(5)).toNumber()] = _module.__default.safeModuloInt(_369_v0, _369_v0);
      _nw60[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(_369_v0);
      _nw60[(new BigNumber(7)).toNumber()] = new BigNumber(834);
      _nw60[(new BigNumber(8)).toNumber()] = _369_v0;
      _nw60[(new BigNumber(9)).toNumber()] = new BigNumber((_396_v24).length);
      _nw60[(new BigNumber(10)).toNumber()] = _369_v0;
      _nw60[(new BigNumber(11)).toNumber()] = new BigNumber(-412);
      _nw60[(new BigNumber(12)).toNumber()] = _369_v0;
      _nw60[(new BigNumber(13)).toNumber()] = _369_v0;
      _nw60[(new BigNumber(14)).toNumber()] = new BigNumber((_397_v25).length);
      _398_v26 = _nw60;
      let _index47 = _module.__default.safeIndex(new BigNumber(925), new BigNumber((_398_v26).length));
      (_398_v26)[_index47] = (new BigNumber((_370_v1).length)).multipliedBy(_369_v0);
      let _index48 = _module.__default.safeIndex(new BigNumber(925), new BigNumber((_398_v26).length));
      let _rhs43 = new BigNumber((_this.f23).length);
      let _rhs44 = _369_v0;
      let _rhs45 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_369_v0));
      let _rhs46 = _395_v23;
      let _rhs47 = _module.__default.safeDivisionInt(_369_v0, (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_397_v25, _397_v25)).length))));
      let _lhs38 = globalState;
      let _lhs39 = globalState;
      let _lhs40 = _398_v26;
      let _lhs41 = _module.__default.safeIndex(new BigNumber(925), new BigNumber((_398_v26).length));
      _lhs38.f20 = _rhs43;
      _369_v0 = _rhs44;
      _lhs39.f20 = _rhs45;
      _395_v23 = _rhs46;
      _lhs40[_lhs41] = _rhs47;
      r0 = (_398_v26)[_module.__default.safeIndex(new BigNumber(925), new BigNumber((_398_v26).length))];
      let _399_i1;
      _399_i1 = _dafny.ZERO;
      L2: {
        while (false) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_399_i1)) {
              break L2;
            }
            _399_i1 = (_399_i1).plus(_dafny.ONE);
            let _400_v27;
            _400_v27 = new _dafny.CodePoint('i'.codePointAt(0));
            let _401_v28;
            _401_v28 = _dafny.Map.Empty.slice().updateUnsafe((_398_v26)[_module.__default.safeIndex(new BigNumber(925), new BigNumber((_398_v26).length))],_400_v27);
            let _402_v29;
            _402_v29 = _dafny.Map.Empty.slice().updateUnsafe(_this.f22,new BigNumber((_401_v28).length));
            let _403_v30;
            _403_v30 = _dafny.Set.fromElements(_module.__default.fm5((_398_v26)[_module.__default.safeIndex(new BigNumber(925), new BigNumber((_398_v26).length))], new BigNumber((_dafny.Seq.of(true)).length), _369_v0, globalState));
            let _index49 = _module.__default.safeIndex(new BigNumber(925), new BigNumber((_398_v26).length));
            (_398_v26)[_index49] = ((new BigNumber((_402_v29).length)).minus(new BigNumber(914))).minus(new BigNumber((((false) ? (_403_v30) : (_dafny.Set.fromElements(_370_v1, _dafny.Seq.of(_369_v0, (_398_v26)[_module.__default.safeIndex(new BigNumber(925), new BigNumber((_398_v26).length))], _369_v0))))).length));
            let _404_v31;
            let _init6 = ((_405_p0) => function (_406_i2) {
              return _dafny.MultiSet.fromElements(!(_this.f22), _405_p0);
            })(p0);
            let _nw61 = Array((new BigNumber(25)).toNumber());
            for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw61.length); _i0_6++) {
              _nw61[_i0_6] = _init6(new BigNumber(_i0_6));
            }
            _404_v31 = _nw61;
            let _407_v32;
            let _nw62 = new _module.C1();
            _nw62.__ctor(_404_v31, true);
            _407_v32 = _nw62;
            let _408_v33;
            _408_v33 = _dafny.Set.fromElements(_400_v27);
            let _409_v34;
            _409_v34 = _module.D1.create_DC3(_dafny.Set.fromElements(_408_v33));
            let _410_v35;
            _410_v35 = _dafny.Set.fromElements(_409_v34, _409_v34);
            let _411_v36;
            _411_v36 = _dafny.Set.fromElements((_410_v35).Difference(_410_v35), _dafny.Set.fromElements(_module.__default.fm10(globalState), _409_v34, _module.__default.fm10(globalState), _409_v34));
            _411_v36 = _411_v36;
            let _412_v37;
            let _nw63 = Array((new BigNumber(20)).toNumber()).fill(_dafny.MultiSet.Empty);
            _412_v37 = _nw63;
            let _413_v38;
            _413_v38 = _dafny.MultiSet.fromElements(_this.f21, _dafny.Seq.update(_this.f21, _module.__default.safeIndex(new BigNumber((_397_v25).length), new BigNumber((_this.f21).length)), _400_v27), _this.f21, _this.f21);
            let _index50 = _module.__default.safeIndex(new BigNumber(214), new BigNumber((_412_v37).length));
            (_412_v37)[_index50] = _413_v38;
            let _414_v39;
            _414_v39 = _module.D0.create_DC1((_407_v32).f25);
            let _415_v40;
            _415_v40 = _dafny.Set.fromElements((_414_v39).dtor_cf1, false, !(_this.f22), p0);
            let _416_v41;
            _416_v41 = _dafny.Map.Empty.slice().updateUnsafe(_407_v32,_415_v40);
            let _index51 = _module.__default.safeIndex(new BigNumber(214), new BigNumber((_412_v37).length));
            (_412_v37)[_index51] = _module.__default.fm13(_this.f22, _415_v40, _module.__default.safeDivisionInt(new BigNumber((_this.f21).length), new BigNumber((_dafny.MultiSet.fromElements((((_416_v41).contains(_407_v32)) ? ((_416_v41).get(_407_v32)) : (_415_v40)), _415_v40, _dafny.Set.fromElements(false, p0, (_407_v32).f25, (_407_v32).f25, (_407_v32).f25))).cardinality())), p0, globalState);
          }
        }
      }
      let _417_v43;
      _417_v43 = _module.D1.create_DC3(_dafny.Set.fromElements(function () {
  let _coll12 = new _dafny.Set();
  for (const _compr_12 of (_dafny.MultiSet.FromArray(_this.f21)).Elements) {
    let _418_v42 = _compr_12;
    if ((_dafny.MultiSet.FromArray(_this.f21)).contains(_418_v42)) {
      _coll12.add(_418_v42);
    }
  }
  return _coll12;
}()));
      let _419_v44;
      _419_v44 = _dafny.Map.Empty.slice().updateUnsafe(_369_v0,p0);
      let _pat_let_tv20 = _398_v26;
      let _pat_let_tv21 = _398_v26;
      let _pat_let_tv22 = _369_v0;
      let _pat_let_tv23 = _369_v0;
      let _pat_let_tv24 = _398_v26;
      let _pat_let_tv25 = _398_v26;
      let _rhs48 = _dafny.Seq.Concat(_this.f21, _this.f21);
      let _rhs49 = function (_source6) {
        if (_source6.is_DC4) {
          let _420___mcc_h0 = (_source6).cf4;
          let _421___mcc_h1 = (_source6).cf5;
          let _422___mcc_h2 = (_source6).cf6;
          let _423_cf6 = _422___mcc_h2;
          let _424_cf5 = _421___mcc_h1;
          let _425_cf4 = _420___mcc_h0;
          return ((_pat_let_tv21)[_module.__default.safeIndex(new BigNumber(925), new BigNumber((_pat_let_tv20).length))]).minus((_module.D1.create_DC5(_this.f22, _dafny.Map.Empty.slice().updateUnsafe(!(_this.f22),new BigNumber((_423_cf6).length)), new BigNumber((_dafny.Set.fromElements(_this.f22, false)).length), _dafny.MultiSet.fromElements(_pat_let_tv22, new BigNumber((_this.f21).length)))).dtor_cf9);
        } else if (_source6.is_DC5) {
          let _426___mcc_h3 = (_source6).cf7;
          let _427___mcc_h4 = (_source6).cf8;
          let _428___mcc_h5 = (_source6).cf9;
          let _429___mcc_h6 = (_source6).cf10;
          let _430_cf10 = _429___mcc_h6;
          let _431_cf9 = _428___mcc_h5;
          let _432_cf8 = _427___mcc_h4;
          let _433_cf7 = _426___mcc_h3;
          return _pat_let_tv23;
        } else {
          let _434___mcc_h7 = (_source6).cf3;
          let _435_cf3 = _434___mcc_h7;
          return ((_pat_let_tv25)[_module.__default.safeIndex(new BigNumber(925), new BigNumber((_pat_let_tv24).length))]).plus(new BigNumber(443));
        }
      }(_417_v43);
      let _rhs50 = _398_v26;
      let _rhs51 = ((_398_v26)[_module.__default.safeIndex(new BigNumber(925), new BigNumber((_398_v26).length))]).minus(_369_v0);
      let _rhs52 = new BigNumber(((_419_v44).Merge(_419_v44)).length);
      let _lhs42 = _this;
      let _lhs43 = globalState;
      let _lhs44 = globalState;
      _lhs42.f21 = _rhs48;
      r0 = _rhs49;
      _398_v26 = _rhs50;
      _lhs43.f20 = _rhs51;
      _lhs44.f20 = _rhs52;
      let _436_v45;
      _436_v45 = new _dafny.CodePoint('s'.codePointAt(0));
      let _437_v46;
      _437_v46 = _dafny.Set.fromElements(_this.f22, _this.f22, p0);
      let _438_v47;
      _438_v47 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_437_v46).length),_module.__default.fm6(globalState));
      let _pat_let_tv26 = p0;
      let _pat_let_tv27 = p0;
      let _pat_let_tv28 = p0;
      let _pat_let_tv29 = p0;
      let _pat_let_tv30 = p0;
      r0 = new BigNumber((function (_source7) {
        if (_source7.is_DC4) {
          let _439___mcc_h8 = (_source7).cf4;
          let _440___mcc_h9 = (_source7).cf5;
          let _441___mcc_h10 = (_source7).cf6;
          let _442_cf6 = _441___mcc_h10;
          let _443_cf5 = _440___mcc_h9;
          let _444_cf4 = _439___mcc_h8;
          return (_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv26,_this.f22)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv27,false));
        } else if (_source7.is_DC5) {
          let _445___mcc_h11 = (_source7).cf7;
          let _446___mcc_h12 = (_source7).cf8;
          let _447___mcc_h13 = (_source7).cf9;
          let _448___mcc_h14 = (_source7).cf10;
          let _449_cf10 = _448___mcc_h14;
          let _450_cf9 = _447___mcc_h13;
          let _451_cf8 = _446___mcc_h12;
          let _452_cf7 = _445___mcc_h11;
          return (_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv28,_452_cf7)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_this.f22));
        } else {
          let _453___mcc_h15 = (_source7).cf3;
          let _454_cf3 = _453___mcc_h15;
          return (_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv29,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f22,_pat_let_tv30));
        }
      }(_module.D1.create_DC4(_436_v45, _438_v47, _370_v1))).length);
      return r0;
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = false;
      let _455_v0;
      let _nw64 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
      _455_v0 = _nw64;
      let _index52 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_455_v0).length));
      (_455_v0)[_index52] = new BigNumber((_this.f21).length);
      let _index53 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_455_v0).length));
      (_455_v0)[_index53] = (p0).minus(p0);
      let _456_v1;
      _456_v1 = _dafny.Seq.of(_this.f22);
      let _457_v2;
      _457_v2 = _dafny.Set.fromElements(true);
      _456_v1 = _dafny.Seq.of(!(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(350),new BigNumber((_457_v2).length))).contains((_455_v0)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_455_v0).length))]), _this.f22, (_this.f22) && (_this.f22));
      if (_this.f22) {
        let _458_v3;
        _458_v3 = _dafny.MultiSet.fromElements(_this.f22);
        (globalState).f3 = ((!(((!(_this.f22)) ? (_this.f22) : (_this.f22)))) ? (_module.__default.safeModuloInt(_module.__default.fm6(globalState), (_455_v0)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_455_v0).length))])) : (new BigNumber((_458_v3).cardinality())));
        let _459_v4;
        _459_v4 = new _dafny.CodePoint('u'.codePointAt(0));
        let _460_v5;
        _460_v5 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_458_v3).cardinality()),(_455_v0)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_455_v0).length))]);
        let _461_v6;
        _461_v6 = _dafny.Seq.of((_455_v0)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_455_v0).length))]);
        let _462_v7;
        _462_v7 = _module.D1.create_DC4(_459_v4, _460_v5, _461_v6);
        _462_v7 = _462_v7;
        let _463_v8;
        _463_v8 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm12((_455_v0)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_455_v0).length))], _dafny.MultiSet.FromArray(_dafny.Seq.of(_this.f22)), _this.f22, _dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(false, !(true)), _458_v3, _458_v3, _dafny.MultiSet.fromElements(_this.f22, _this.f22, _this.f22), _458_v3), globalState),(_dafny.ZERO).minus((_455_v0)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_455_v0).length))]));
        let _464_v9;
        _464_v9 = _dafny.Map.Empty.slice().updateUnsafe(_this.f22,false);
        let _465_v10;
        _465_v10 = _module.D0.create_DC1((((_464_v9).contains(true)) ? ((_464_v9).get(true)) : (_this.f22)));
        _463_v8 = (_463_v8).update((_465_v10).dtor_cf1, (_dafny.ZERO).minus(p0));
        let _466_v11;
        let _nw65 = Array((new BigNumber(27)).toNumber());
        _nw65[(_dafny.ZERO).toNumber()] = _this.f22;
        _nw65[(_dafny.ONE).toNumber()] = _this.f22;
        _nw65[(new BigNumber(2)).toNumber()] = _this.f22;
        _nw65[(new BigNumber(3)).toNumber()] = _this.f22;
        _nw65[(new BigNumber(4)).toNumber()] = _this.f22;
        _nw65[(new BigNumber(5)).toNumber()] = _this.f22;
        _nw65[(new BigNumber(6)).toNumber()] = _this.f22;
        _nw65[(new BigNumber(7)).toNumber()] = _this.f22;
        _nw65[(new BigNumber(8)).toNumber()] = false;
        _nw65[(new BigNumber(9)).toNumber()] = _this.f22;
        _nw65[(new BigNumber(10)).toNumber()] = _this.f22;
        _nw65[(new BigNumber(11)).toNumber()] = true;
        _nw65[(new BigNumber(12)).toNumber()] = _this.f22;
        _nw65[(new BigNumber(13)).toNumber()] = _this.f22;
        _nw65[(new BigNumber(14)).toNumber()] = _this.f22;
        _nw65[(new BigNumber(15)).toNumber()] = _this.f22;
        _nw65[(new BigNumber(16)).toNumber()] = _this.f22;
        _nw65[(new BigNumber(17)).toNumber()] = true;
        _nw65[(new BigNumber(18)).toNumber()] = _this.f22;
        _nw65[(new BigNumber(19)).toNumber()] = false;
        _nw65[(new BigNumber(20)).toNumber()] = _this.f22;
        _nw65[(new BigNumber(21)).toNumber()] = _this.f22;
        _nw65[(new BigNumber(22)).toNumber()] = !(_this.f22);
        _nw65[(new BigNumber(23)).toNumber()] = _this.f22;
        _nw65[(new BigNumber(24)).toNumber()] = _this.f22;
        _nw65[(new BigNumber(25)).toNumber()] = false;
        _nw65[(new BigNumber(26)).toNumber()] = _this.f22;
        _466_v11 = _nw65;
        let _467_v12;
        _467_v12 = _module.D0.create_DC0(_466_v11);
        let _468_v13;
        _468_v13 = _module.D0.create_DC2(_467_v12);
        let _469_v14;
        _469_v14 = _dafny.Map.Empty.slice().updateUnsafe(_468_v13,new BigNumber((_this.f21).length));
        _469_v14 = (_469_v14).update(_468_v13, p0);
        r1 = _this.f22;
      } else {
        (globalState).f20 = ((_455_v0)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_455_v0).length))]).minus(new BigNumber((_dafny.Seq.UnicodeFromString("etxvn")).length));
        (globalState).f12 = new _dafny.CodePoint('d'.codePointAt(0));
        let _470_v15;
        let _init7 = ((_471_v0) => function (_472_i0) {
          return _dafny.Set.fromElements((_471_v0)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_471_v0).length))]);
        })(_455_v0);
        let _nw66 = Array((new BigNumber(15)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw66.length); _i0_7++) {
          _nw66[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _470_v15 = _nw66;
        let _473_v16;
        _473_v16 = _dafny.Set.fromElements((_455_v0)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_455_v0).length))], p0);
        let _index54 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_470_v15).length));
        (_470_v15)[_index54] = (_473_v16);
        let _474_v17;
        _474_v17 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm1(_this.f21, true, false, globalState),_this.f22);
        let _index55 = _module.__default.safeIndex(new BigNumber(983), new BigNumber((_470_v15).length));
        (_470_v15)[_index55] = _module.__default.fm14(_474_v17, p0, globalState);
        let _source8 = _473_v16;
        let _475___mcc_h0 = _source8;
        let _476_cf11 = _475___mcc_h0;
        (globalState).f20 = (_module.__default.fm6(globalState)).plus(_module.__default.safeDivisionInt(new BigNumber((_this.f21).length), (_455_v0)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_455_v0).length))]));
        _473_v16 = _473_v16;
        _474_v17 = _474_v17;
        (_this).f22 = (((_this.f22) ? (_this.f22) : (_this.f22))) && ((_this).fm2(_this.f22, globalState));
        let _477_v18;
        let _nw67 = Array((new BigNumber(4)).toNumber()).fill(_dafny.MultiSet.Empty);
        _477_v18 = _nw67;
        let _478_v19;
        _478_v19 = _dafny.Map.Empty.slice().updateUnsafe(_477_v18,(_dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_this.f21).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_this.f21).length),p0)));
        let _479_v20;
        _479_v20 = _dafny.Map.Empty.slice().updateUnsafe((_455_v0)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_455_v0).length))],p0);
        _478_v19 = (_478_v19).update(_477_v18, _479_v20);
      }
      if (true) {
        let _480_v21;
        _480_v21 = _dafny.Set.fromElements((_455_v0)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_455_v0).length))]);
        let _481_v22;
        _481_v22 = (_480_v21).Intersect(_480_v21);
        _481_v22 = _481_v22;
        let _482_v23;
        _482_v23 = _dafny.Set.fromElements(_this.f21, _this.f21);
        let _483_v24;
        let _nw68 = new _module.C0();
        _nw68.__ctor(_482_v23);
        _483_v24 = _nw68;
        _483_v24 = _483_v24;
        let _484_v25;
        let _init8 = function (_485_i1) {
          return ((_this.f22) ? (_this.f22) : (_this.f22));
        };
        let _nw69 = Array((new BigNumber(15)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw69.length); _i0_8++) {
          _nw69[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _484_v25 = _nw69;
        let _index56 = _module.__default.safeIndex(new BigNumber(184), new BigNumber((_484_v25).length));
        (_484_v25)[_index56] = _this.f22;
        let _486_v26;
        _486_v26 = _dafny.MultiSet.fromElements(_this.f22, (_this).fm1(_this.f21, _this.f22, _this.f22, globalState));
        let _index57 = _module.__default.safeIndex(new BigNumber(184), new BigNumber((_484_v25).length));
        (_484_v25)[_index57] = !(((new BigNumber(211)).isEqualTo((_455_v0)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_455_v0).length))])) === ((_486_v26).IsDisjointFrom(_dafny.MultiSet.FromArray(_456_v1))));
        r1 = _this.f22;
        let _487_v27;
        let _nw70 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Map.Empty);
        _487_v27 = _nw70;
        let _488_v28;
        _488_v28 = _dafny.MultiSet.fromElements(_this.f21, _module.__default.fm11(false, globalState), _dafny.Seq.UnicodeFromString("bdbdchmi"));
        let _index58 = _module.__default.safeIndex(new BigNumber(184), new BigNumber((_484_v25).length));
        let _rhs53 = (((_488_v28).contains(_this.f21)) ? ((_488_v28).get(_this.f21)) : ((_455_v0)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_455_v0).length))]));
        let _rhs54 = (_484_v25)[_module.__default.safeIndex(new BigNumber(184), new BigNumber((_484_v25).length))];
        let _rhs55 = (_484_v25)[_module.__default.safeIndex(new BigNumber(184), new BigNumber((_484_v25).length))];
        let _rhs56 = _487_v27;
        let _lhs45 = globalState;
        let _lhs46 = _484_v25;
        let _lhs47 = _module.__default.safeIndex(new BigNumber(184), new BigNumber((_484_v25).length));
        let _lhs48 = _this;
        _lhs45.f20 = _rhs53;
        _lhs46[_lhs47] = _rhs54;
        _lhs48.f22 = _rhs55;
        _487_v27 = _rhs56;
      } else {
        let _489_v29;
        let _nw71 = Array((new BigNumber(8)).toNumber()).fill(false);
        _489_v29 = _nw71;
        _489_v29 = _489_v29;
        let _490_v30;
        _490_v30 = _dafny.MultiSet.fromElements(_this.f22, false);
        let _491_v32;
        _491_v32 = _dafny.Set.fromElements(_490_v30);
        let _492_v33;
        _492_v33 = _dafny.Seq.of(function () {
          let _coll13 = new _dafny.Map();
          for (const _compr_13 of (_491_v32).Elements) {
            let _493_v31 = _compr_13;
            if ((_491_v32).contains(_493_v31)) {
              _coll13.push([_493_v31,(_dafny.ZERO).minus(p0)]);
            }
          }
          return _coll13;
        }());
        let _494_v34;
        _494_v34 = _dafny.Map.Empty.slice().updateUnsafe(_490_v30,(_492_v33)[_module.__default.safeIndex((_dafny.ZERO).minus((_455_v0)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_455_v0).length))]), new BigNumber((_492_v33).length))]);
        _494_v34 = (_494_v34).update((_dafny.MultiSet.fromElements(_this.f22)).Intersect(_490_v30), (_492_v33)[_module.__default.safeIndex((_455_v0)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_455_v0).length))], new BigNumber((_492_v33).length))]);
        _456_v1 = _dafny.Seq.Concat(_456_v1, _456_v1);
        if (!(_this.f22)) {
          let _index59 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_455_v0).length));
          let _rhs57 = (_dafny.ZERO).minus(((_this.f22) ? (((((_490_v30).contains(_this.f22)) ? ((_490_v30).get(_this.f22)) : (new BigNumber(-936)))).plus((_455_v0)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_455_v0).length))])) : (p0)));
          let _rhs58 = _this.f22;
          let _rhs59 = _455_v0;
          let _lhs49 = _455_v0;
          let _lhs50 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_455_v0).length));
          let _lhs51 = _this;
          _lhs49[_lhs50] = _rhs57;
          _lhs51.f22 = _rhs58;
          _455_v0 = _rhs59;
          let _495_v35;
          _495_v35 = _dafny.Map.Empty.slice().updateUnsafe(_this.f22,_489_v29);
          let _496_v36;
          _496_v36 = _module.D0.create_DC1(_this.f22);
          let _497_v37;
          _497_v37 = _dafny.Map.Empty.slice().updateUnsafe((((_495_v35).contains(_this.f22)) ? ((_495_v35).get(_this.f22)) : (_489_v29)),_496_v36);
          (globalState).f20 = new BigNumber(((_497_v37).Merge(_497_v37)).length);
          r1 = (_456_v1)[_module.__default.safeIndex(p0, new BigNumber((_456_v1).length))];
          let _498_v38;
          _498_v38 = new _dafny.CodePoint('x'.codePointAt(0));
          let _499_v39;
          _499_v39 = _dafny.MultiSet.fromElements(_498_v38, _498_v38);
          let _rhs60 = (_499_v39).Difference(_499_v39);
          let _rhs61 = _this.f22;
          let _lhs52 = _this;
          _499_v39 = _rhs60;
          _lhs52.f22 = _rhs61;
          let _500_v40;
          _500_v40 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm3(globalState),_this.f22);
          let _501_v41;
          _501_v41 = _dafny.Seq.of(p0);
          let _502_v42;
          _502_v42 = _dafny.Set.fromElements(new BigNumber((_501_v41).length));
          _500_v40 = (_500_v40).update((_dafny.Set.fromElements(new BigNumber(28), p0, new BigNumber((_490_v30).cardinality()))).IsDisjointFrom(_502_v42), (_this).fm3(globalState));
        } else {
          let _503_v43;
          _503_v43 = _dafny.Map.Empty.slice().updateUnsafe((_455_v0)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_455_v0).length))],_this.f22);
          (_this).f22 = (new BigNumber((_this.f21).length)).isLessThanOrEqualTo(_module.__default.safeModuloInt(new BigNumber((_503_v43).length), (_455_v0)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_455_v0).length))]));
          r1 = !(((_this.f22) ? (!(_this.f22)) : ((((((_503_v43).contains(new BigNumber(643))) ? ((_503_v43).get(new BigNumber(643))) : (_this.f22))) ? (_this.f22) : (_this.f22)))));
          let _504_v44;
          _504_v44 = new _dafny.CodePoint('g'.codePointAt(0));
          let _505_v45;
          _505_v45 = _dafny.Set.fromElements(_504_v44, _504_v44);
          let _506_v46;
          _506_v46 = _module.D1.create_DC3(_dafny.Set.fromElements(_505_v45, _dafny.Set.fromElements(_504_v44)));
          let _507_v47;
          _507_v47 = _dafny.Seq.of(_506_v46, _module.__default.fm10(globalState));
          (globalState).f20 = new BigNumber((_507_v47).length);
          let _508_v48;
          let _nw72 = Array((new BigNumber(28)).toNumber()).fill([]);
          _508_v48 = _nw72;
          let _509_v49;
          let _init9 = ((_510_v30) => function (_511_i2) {
            return _510_v30;
          })(_490_v30);
          let _nw73 = Array((new BigNumber(6)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw73.length); _i0_9++) {
            _nw73[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _509_v49 = _nw73;
          let _index60 = _module.__default.safeIndex(new BigNumber(40), new BigNumber((_508_v48).length));
          (_508_v48)[_index60] = _509_v49;
          let _index61 = _module.__default.safeIndex(new BigNumber(40), new BigNumber((_508_v48).length));
          (_508_v48)[_index61] = _509_v49;
          _489_v29 = _489_v29;
        }
        let _512_v50;
        let _init10 = ((_513_v30, _514_p0) => function (_515_i3) {
          return (_513_v30).update(_this.f22, _module.__default.abs(_514_p0));
        })(_490_v30, p0);
        let _nw74 = Array((new BigNumber(2)).toNumber());
        for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw74.length); _i0_10++) {
          _nw74[_i0_10] = _init10(new BigNumber(_i0_10));
        }
        _512_v50 = _nw74;
        let _516_v51;
        let _nw75 = new _module.C1();
        _nw75.__ctor(_512_v50, _this.f22);
        _516_v51 = _nw75;
        let _517_v52;
        _517_v52 = _dafny.Seq.of(_516_v51, _516_v51);
        r1 = (!(_this.f22)) === ((_dafny.Set.fromElements(_517_v52, _517_v52)).IsProperSubsetOf(_dafny.Set.fromElements(_517_v52, _517_v52, _517_v52)));
      }
      let _518_v53;
      _518_v53 = _dafny.Set.fromElements(_this.f21, _this.f21);
      let _519_v54;
      let _nw76 = new _module.C0();
      _nw76.__ctor(_518_v53);
      _519_v54 = _nw76;
      let _520_i4;
      _520_i4 = _dafny.ZERO;
      L3: {
        while ((p0).isLessThanOrEqualTo(p0)) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_520_i4)) {
              break L3;
            }
            _520_i4 = (_520_i4).plus(_dafny.ONE);
            let _521_v55;
            _521_v55 = _dafny.MultiSet.fromElements(_this.f22);
            let _522_v56;
            _522_v56 = _module.D3.create_DC7(_dafny.MultiSet.fromElements(_this.f22));
            _521_v55 = (_522_v56).dtor_cf12;
            let _index62 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_455_v0).length));
            (_455_v0)[_index62] = p0;
            (globalState).f20 = _module.__default.fm6(globalState);
            let _523_v57;
            _523_v57 = _dafny.Set.fromElements(p0, (_455_v0)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_455_v0).length))], p0);
            let _524_v58;
            _524_v58 = _dafny.Map.Empty.slice().updateUnsafe(((_this.f22) ? (_this.f22) : (_this.f22)),_this.f22);
            let _525_v59;
            _525_v59 = _module.D0.create_DC1(true);
            let _526_v60;
            _526_v60 = _module.D0.create_DC2(_525_v59);
            let _rhs62 = (_module.__default.fm14(_dafny.Map.Empty.slice().updateUnsafe((((_524_v58).contains(!(_this.f22))) ? ((_524_v58).get(!(_this.f22))) : (_this.f22)),_this.f22), new BigNumber((_this.f21).length), globalState)).Difference(_523_v57);
            let _rhs63 = (_456_v1)[_module.__default.safeIndex(new BigNumber((_module.__default.fm15(globalState)).length), new BigNumber((_456_v1).length))];
            let _rhs64 = ((_this.f22) ? (_524_v58) : (_524_v58));
            let _rhs65 = _526_v60;
            _523_v57 = _rhs62;
            r1 = _rhs63;
            _524_v58 = _rhs64;
            _526_v60 = _rhs65;
          }
        }
      }
      let _527_v62;
      _527_v62 = function () {
        let _coll14 = new _dafny.Set();
        for (const _compr_14 of _dafny.IntegerRange(new BigNumber(-25), new BigNumber(428))) {
          let _528_v61 = _compr_14;
          if (((new BigNumber(-25)).isLessThanOrEqualTo(_528_v61)) && ((_528_v61).isLessThan(new BigNumber(428)))) {
            _coll14.add(_module.__default.safeModuloInt(_528_v61, (_455_v0)[_module.__default.safeIndex(new BigNumber(931), new BigNumber((_455_v0).length))]));
          }
        }
        return _coll14;
      }();
      r0 = function (_source9) {
        let _529___mcc_h1 = _source9;
        let _530_cf11 = _529___mcc_h1;
        return _this.f21;
      }(_527_v62);
      let _531_v63;
      _531_v63 = _dafny.MultiSet.fromElements(_this.f22, _this.f22, true, _this.f22, _this.f22);
      let _532_v64;
      _532_v64 = _dafny.Map.Empty.slice().updateUnsafe(_456_v1,p0);
      let _533_v65;
      let _nw77 = Array((_dafny.ONE).toNumber()).fill(false);
      _533_v65 = _nw77;
      let _534_v66;
      _534_v66 = _dafny.Seq.of(_533_v65, _533_v65);
      let _535_v67;
      _535_v67 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_532_v64).length),(_534_v66)[_module.__default.safeIndex(p0, new BigNumber((_534_v66).length))]);
      let _536_v68;
      _536_v68 = _dafny.MultiSet.fromElements(_dafny.MultiSet.FromArray(_456_v1));
      r1 = _module.__default.fm12(new BigNumber((_dafny.Seq.UnicodeFromString("hsldowxme")).length), ((!(_this.f22)) ? (_531_v63) : (_531_v63)), (_535_v67).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(88),_533_v65)), _536_v68, globalState);
      return [r0, r1];
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f21 = _dafny.Seq.UnicodeFromString("");
      this._f22 = false;
      this.f28 = false;
      this._f27 = false;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f21() {
      let _this = this;
      return _this._f21;
    };
    set f21(value) {
      let _this = this;
      _this._f21 = value;
    };
    get f22() {
      let _this = this;
      return _this._f22;
    };
    set f22(value) {
      let _this = this;
      _this._f22 = value;
    };
    __ctor(f27, f28, f21, f22) {
      let _this = this;
      (_this)._f27 = f27;
      (_this).f28 = f28;
      (_this)._f21 = f21;
      (_this)._f22 = f22;
      return;
    }
    fm0(globalState) {
      let _this = this;
      return (((false) ? (_dafny.Set.fromElements(_this.f22, true)) : (_dafny.Set.fromElements((_this).f27, _this.f28)))).Union(_dafny.Set.fromElements((_this).f27));
    };
    fm1(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f27;
    };
    m4(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = [];
      (_this).f22 = false;
      let _hi3 = p1;
      for (let _537_i0 = new BigNumber((_this.f21).length); _537_i0.isLessThan(_hi3); _537_i0 = _537_i0.plus(_dafny.ONE)) {
        let _538_v0;
        let _nw78 = Array((new BigNumber(3)).toNumber());
        _538_v0 = _nw78;
        let _539_v1;
        _539_v1 = _dafny.MultiSet.fromElements(_538_v0, _538_v0);
        _539_v1 = ((_539_v1).Intersect(_539_v1)).Union((_539_v1).Intersect(_539_v1));
        (globalState).f6 = _this.f21;
        if ((_this).f27) {
          let _540_v2;
          _540_v2 = _dafny.Map.Empty.slice().updateUnsafe(_this.f22,p1);
          let _541_v3;
          _541_v3 = _dafny.Set.fromElements(_540_v2, _540_v2);
          (_this).f22 = ((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_this.f28,p3), _540_v2, _540_v2, _dafny.Map.Empty.slice().updateUnsafe(_this.f28,p2))).Union(_541_v3)).IsDisjointFrom(_541_v3);
          let _542_v4;
          _542_v4 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("cafmfqmaf"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-524)), function (_543_i1) {
            return new _dafny.CodePoint('m'.codePointAt(0));
          }));
          let _544_v5;
          let _nw79 = new _module.C0();
          _nw79.__ctor(_542_v4);
          _544_v5 = _nw79;
          let _545_v6;
          _545_v6 = _module.D4.create_DC10(_544_v5);
          let _rhs66 = p3;
          let _rhs67 = (_545_v6).dtor_cf20;
          let _lhs53 = globalState;
          _lhs53.f20 = _rhs66;
          _544_v5 = _rhs67;
          let _546_v7;
          _546_v7 = _dafny.Seq.of(p1, _module.__default.fm6(globalState), p2);
          _546_v7 = _546_v7;
          let _547_v8;
          let _nw80 = Array((new BigNumber(10)).toNumber()).fill(false);
          _547_v8 = _nw80;
          let _rhs68 = _547_v8;
          let _rhs69 = ((true) ? (_module.__default.safeDivisionInt(_537_i0, _537_i0)) : (_module.__default.safeDivisionInt(_537_i0, p3)));
          let _lhs54 = globalState;
          _547_v8 = _rhs68;
          _lhs54.f3 = _rhs69;
          (_this).f22 = false;
        } else {
          let _548_v9;
          _548_v9 = _dafny.MultiSet.fromElements(p3, p0);
          let _549_v10;
          _549_v10 = _dafny.Map.Empty.slice().updateUnsafe(_this.f22,new BigNumber((_548_v9).cardinality()));
          let _550_v11;
          _550_v11 = _dafny.Seq.of((_dafny.Map.Empty.slice().updateUnsafe((_this).f27,p0)).Merge(_549_v10));
          _550_v11 = _550_v11;
          let _551_v12;
          _551_v12 = _dafny.Set.fromElements(p1, new BigNumber(-257), new BigNumber(555), new BigNumber(879), _537_i0);
          let _552_v14;
          _552_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("mkm")).length),function () {
            let _coll15 = new _dafny.Set();
            for (const _compr_15 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(502)), function (_553_i2) {
              return new BigNumber(711);
            })).Elements) {
              let _554_v13 = _compr_15;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(502)), function (_553_i2) {
                return new BigNumber(711);
              }), _554_v13)) {
                _coll15.add((_554_v13).plus(new BigNumber(-762)));
              }
            }
            return _coll15;
          }());
          let _555_v15;
          _555_v15 = _dafny.Set.fromElements(_551_v12, (((_552_v14).contains(p1)) ? ((_552_v14).get(p1)) : (_551_v12)));
          (globalState).f3 = new BigNumber(((_555_v15).Union(_dafny.Set.fromElements(_dafny.Set.fromElements(p0, new BigNumber((_551_v12).length), p2)))).length);
          let _556_v16;
          _556_v16 = _module.D3.create_DC7(_dafny.MultiSet.fromElements((_this).f27, _this.f28));
          let _557_v17;
          _557_v17 = _dafny.Map.Empty.slice().updateUnsafe(p3,_556_v16);
          _557_v17 = (_557_v17).Merge(_557_v17);
          (_this).f22 = ((_this).f27) || ((_this).f27);
          let _558_v18;
          let _nw81 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _558_v18 = _nw81;
          let _index63 = _module.__default.safeIndex(new BigNumber(391), new BigNumber((_558_v18).length));
          (_558_v18)[_index63] = _this.f21;
          let _index64 = _module.__default.safeIndex(new BigNumber(391), new BigNumber((_558_v18).length));
          (_558_v18)[_index64] = _this.f21;
        }
        let _559_v19;
        let _init11 = function (_560_i3) {
          return _dafny.areEqual(_this.f21, _this.f21);
        };
        let _nw82 = Array((new BigNumber(5)).toNumber());
        for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw82.length); _i0_11++) {
          _nw82[_i0_11] = _init11(new BigNumber(_i0_11));
        }
        _559_v19 = _nw82;
        let _index65 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_559_v19).length));
        (_559_v19)[_index65] = (_this).f27;
        let _index66 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_559_v19).length));
        (_559_v19)[_index66] = (_this.f22) || ((_this).f27);
      }
      let _561_v20;
      let _nw83 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
      _561_v20 = _nw83;
      let _562_v21;
      let _nw84 = Array((new BigNumber(23)).toNumber()).fill(false);
      _562_v21 = _nw84;
      let _563_v22;
      _563_v22 = _dafny.Seq.of(_this.f28, false, true, _this.f22, _this.f22);
      let _index67 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_562_v21).length));
      (_562_v21)[_index67] = (_563_v22)[_module.__default.safeIndex(p0, new BigNumber((_563_v22).length))];
      let _index68 = _module.__default.safeIndex(new BigNumber(791), new BigNumber((_562_v21).length));
      (_562_v21)[_index68] = (_this).f27;
      let _index69 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_562_v21).length));
      let _index70 = _module.__default.safeIndex(new BigNumber(791), new BigNumber((_562_v21).length));
      let _rhs70 = _561_v20;
      let _rhs71 = _this.f22;
      let _rhs72 = (_563_v22)[_module.__default.safeIndex(new BigNumber(-531), new BigNumber((_563_v22).length))];
      let _rhs73 = ((_this.f22) ? ((_this.f28) || (_this.f22)) : (true));
      let _lhs55 = _562_v21;
      let _lhs56 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_562_v21).length));
      let _lhs57 = _562_v21;
      let _lhs58 = _module.__default.safeIndex(new BigNumber(791), new BigNumber((_562_v21).length));
      let _lhs59 = _this;
      _561_v20 = _rhs70;
      _lhs55[_lhs56] = _rhs71;
      _lhs57[_lhs58] = _rhs72;
      _lhs59.f28 = _rhs73;
      let _564_v23;
      _564_v23 = _dafny.Set.fromElements(new BigNumber(((_dafny.MultiSet.FromArray(_dafny.Seq.of(p2, p3))).update(p2, _module.__default.abs(p0))).cardinality()), p1);
      let _565_v24;
      _565_v24 = _564_v23;
      let _source10 = _565_v24;
      let _566___mcc_h0 = _source10;
      let _567_cf11 = _566___mcc_h0;
      (globalState).f3 = p3;
      let _568_v25;
      _568_v25 = _dafny.Seq.of(_561_v20, _561_v20, ((_this.f28) ? (_561_v20) : (_561_v20)), _561_v20, _561_v20);
      _568_v25 = _568_v25;
      (_this).f28 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(336)), function (_569_i4) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      }), _this.f21);
      let _570_v26;
      _570_v26 = new _dafny.CodePoint('t'.codePointAt(0));
      let _571_v27;
      _571_v27 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(468)), ((_572_v26) => function (_573_i5) {
        return _572_v26;
      })(_570_v26)));
      let _574_v28;
      _574_v28 = _dafny.Map.Empty.slice().updateUnsafe(_570_v26,_571_v27);
      let _575_v29;
      let _nw85 = new _module.C0();
      _nw85.__ctor((((_574_v28).contains(_570_v26)) ? ((_574_v28).get(_570_v26)) : (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("ce")))));
      _575_v29 = _nw85;
      let _576_v31;
      _576_v31 = _dafny.Map.Empty.slice().updateUnsafe(p0,false);
      if (_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(526)), ((_577_p2) => function (_578_i6) {
        return function () {
          let _coll16 = new _dafny.Map();
          for (const _compr_16 of _dafny.IntegerRange(new BigNumber(680), new BigNumber(613))) {
            let _579_v30 = _compr_16;
            if (((new BigNumber(680)).isLessThanOrEqualTo(_579_v30)) && ((_579_v30).isLessThan(new BigNumber(613)))) {
              _coll16.push([(_579_v30).plus(_577_p2),(_this).f27]);
            }
          }
          return _coll16;
        }();
      })(p2)), _dafny.Seq.of(_576_v31, _576_v31, _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f27), _576_v31))) {
        let _580_v32;
        let _nw86 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Seq.of());
        _580_v32 = _nw86;
        let _index71 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_580_v32).length));
        (_580_v32)[_index71] = _dafny.Seq.Concat(_563_v22, _dafny.Seq.of(_this.f22, (_562_v21)[_module.__default.safeIndex(new BigNumber(690), new BigNumber((_562_v21).length))], (_562_v21)[_module.__default.safeIndex(new BigNumber(690), new BigNumber((_562_v21).length))], (_562_v21)[_module.__default.safeIndex(new BigNumber(690), new BigNumber((_562_v21).length))]));
        let _index72 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_580_v32).length));
        (_580_v32)[_index72] = _dafny.Seq.Concat(((!(_this.f22)) ? (_563_v22) : (_563_v22)), _563_v22);
        let _581_v33;
        let _init12 = function (_582_i7) {
          return _dafny.MultiSet.fromElements(!(_this.f22), _this.f22);
        };
        let _nw87 = Array((new BigNumber(26)).toNumber());
        for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw87.length); _i0_12++) {
          _nw87[_i0_12] = _init12(new BigNumber(_i0_12));
        }
        _581_v33 = _nw87;
        let _583_v34;
        let _nw88 = new _module.C1();
        _nw88.__ctor(_581_v33, (_this).f27);
        _583_v34 = _nw88;
        let _index73 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_561_v20).length));
        (_561_v20)[_index73] = p2;
        let _584_v35;
        _584_v35 = _dafny.Seq.of(new BigNumber((_563_v22).length), new BigNumber((_this.f21).length));
        let _585_v36;
        _585_v36 = _dafny.Seq.of(new BigNumber(244), p3, new BigNumber((_dafny.MultiSet.FromArray(_584_v35)).cardinality()), new BigNumber((_563_v22).length), p2);
        let _586_v37;
        _586_v37 = _dafny.MultiSet.fromElements(new BigNumber(886));
        let _587_v38;
        _587_v38 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_586_v37).cardinality()),p0);
        let _588_v39;
        _588_v39 = _dafny.MultiSet.fromElements((_562_v21)[_module.__default.safeIndex(new BigNumber(690), new BigNumber((_562_v21).length))]);
        let _589_v40;
        _589_v40 = _dafny.MultiSet.fromElements(_588_v39, _588_v39, _588_v39, _588_v39, _588_v39);
        let _590_v41;
        _590_v41 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,false);
        let _index74 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_561_v20).length));
        let _rhs74 = _583_v34;
        let _rhs75 = ((_dafny.Seq.IsProperPrefixOf(_585_v36, _module.__default.fm5(p3, new BigNumber((_587_v38).length), p3, globalState))) ? ((_this.f21)[_module.__default.safeIndex(p0, new BigNumber((_this.f21).length))]) : (new _dafny.CodePoint('k'.codePointAt(0))));
        let _rhs76 = (_module.__default.fm6(globalState)).minus(new BigNumber(((_580_v32)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_580_v32).length))]).length));
        let _rhs77 = ((_module.__default.fm12(p1, _588_v39, _this.f28, _589_v40, globalState)) ? (_module.__default.fm6(globalState)) : (((_585_v36)[_module.__default.safeIndex((((_588_v39).contains((_562_v21)[_module.__default.safeIndex(new BigNumber(690), new BigNumber((_562_v21).length))])) ? ((_588_v39).get((_562_v21)[_module.__default.safeIndex(new BigNumber(690), new BigNumber((_562_v21).length))])) : ((_dafny.ZERO).minus(new BigNumber((_590_v41).length)))), new BigNumber((_585_v36).length))]).plus(p3)));
        let _lhs60 = globalState;
        let _lhs61 = _561_v20;
        let _lhs62 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_561_v20).length));
        let _lhs63 = globalState;
        _583_v34 = _rhs74;
        _lhs60.f12 = _rhs75;
        _lhs61[_lhs62] = _rhs76;
        _lhs63.f3 = _rhs77;
        let _591_v42;
        _591_v42 = new _dafny.CodePoint('e'.codePointAt(0));
        let _592_v43;
        _592_v43 = _dafny.Set.fromElements(new _dafny.CodePoint('e'.codePointAt(0)), _591_v42, _591_v42, _591_v42);
        let _593_v44;
        _593_v44 = _dafny.Set.fromElements(_592_v43);
        let _594_v45;
        _594_v45 = _module.D1.create_DC3(_593_v44);
        let _source11 = _594_v45;
        if (_source11.is_DC4) {
          let _595___mcc_h1 = (_source11).cf4;
          let _596___mcc_h2 = (_source11).cf5;
          let _597___mcc_h3 = (_source11).cf6;
          let _598_cf6 = _597___mcc_h3;
          let _599_cf5 = _596___mcc_h2;
          let _600_cf4 = _595___mcc_h1;
          let _index75 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_562_v21).length));
          (_562_v21)[_index75] = (_583_v34).f25;
          _600_cf4 = _591_v42;
          let _pat_let_tv31 = _593_v44;
          _594_v45 = function (_pat_let8_0) {
            return function (_601_dt__update__tmp_h0) {
              return function (_pat_let9_0) {
                return function (_602_dt__update_hcf3_h0) {
                  return _module.D1.create_DC3(_602_dt__update_hcf3_h0);
                }(_pat_let9_0);
              }(_pat_let_tv31);
            }(_pat_let8_0);
          }(_594_v45);
          let _603_v46;
          _603_v46 = _583_v34;
          let _rhs78 = (_603_v46);
          let _rhs79 = p3;
          let _rhs80 = ((_this.f28) ? ((_this).f27) : (_dafny.Seq.IsProperPrefixOf(_this.f21, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-738)), ((_604_v42) => function (_605_i8) {
            return _604_v42;
          })(_591_v42)))));
          let _lhs64 = globalState;
          let _lhs65 = _this;
          _583_v34 = _rhs78;
          _lhs64.f3 = _rhs79;
          _lhs65.f22 = _rhs80;
        } else if (_source11.is_DC5) {
          let _606___mcc_h4 = (_source11).cf7;
          let _607___mcc_h5 = (_source11).cf8;
          let _608___mcc_h6 = (_source11).cf9;
          let _609___mcc_h7 = (_source11).cf10;
          let _610_cf10 = _609___mcc_h7;
          let _611_cf9 = _608___mcc_h6;
          let _612_cf8 = _607___mcc_h5;
          let _613_cf7 = _606___mcc_h4;
          (_this).f22 = (_588_v39).IsProperSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.Concat((_580_v32)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_580_v32).length))], (_580_v32)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_580_v32).length))])));
          let _index76 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_562_v21).length));
          (_562_v21)[_index76] = _this.f22;
          (_this).f28 = _613_cf7;
          let _614_v47;
          _614_v47 = _dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_563_v22,_this.f22)).length)), _585_v36, _585_v36);
          let _615_v48;
          let _nw89 = new _module.C2();
          _nw89.__ctor(_614_v47, _this.f21, (new BigNumber(645)).isLessThanOrEqualTo((_585_v36)[_module.__default.safeIndex(new BigNumber(((_580_v32)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_580_v32).length))]).length), new BigNumber((_585_v36).length))]));
          _615_v48 = _nw89;
        } else {
          let _616___mcc_h8 = (_source11).cf3;
          let _617_cf3 = _616___mcc_h8;
          let _618_v49;
          _618_v49 = _dafny.Set.fromElements(!((_this).f27), false, !((_583_v34).f25));
          let _index77 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_561_v20).length));
          (_561_v20)[_index77] = ((new BigNumber((_618_v49).length)).plus(new BigNumber(-765))).minus(p1);
          let _index78 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_562_v21).length));
          let _rhs81 = (_dafny.ZERO).minus((((_561_v20)[_module.__default.safeIndex(new BigNumber(331), new BigNumber((_561_v20).length))]).multipliedBy(new BigNumber(876))).minus(_module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_563_v22).length))).length), _module.__default.fm6(globalState))));
          let _rhs82 = (_562_v21)[_module.__default.safeIndex(new BigNumber(690), new BigNumber((_562_v21).length))];
          let _rhs83 = _module.__default.fm4(!_dafny.Seq.contains(_dafny.Seq.of(p1), (_dafny.ZERO).minus((_561_v20)[_module.__default.safeIndex(new BigNumber(331), new BigNumber((_561_v20).length))])), _dafny.MultiSet.fromElements((_583_v34).f25, !(_this.f28), (_this).f27, (_this).f27, (_583_v34).f25), globalState);
          let _lhs66 = globalState;
          let _lhs67 = _562_v21;
          let _lhs68 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_562_v21).length));
          _lhs66.f20 = _rhs81;
          _lhs67[_lhs68] = _rhs82;
          _591_v42 = _rhs83;
          (_this).f22 = !((_583_v34).f25);
          (globalState).f3 = p0;
        }
        let _619_v50;
        _619_v50 = _dafny.Map.Empty.slice().updateUnsafe(_591_v42,new BigNumber((_this.f21).length));
        _619_v50 = (_619_v50).update(_591_v42, _module.__default.safeModuloInt(p0, p0));
        (_this).f22 = _this.f28;
      } else {
        let _620_v51;
        _620_v51 = new _dafny.CodePoint('r'.codePointAt(0));
        let _621_v52;
        _621_v52 = _dafny.Map.Empty.slice().updateUnsafe(p1,_620_v51);
        let _622_v53;
        let _nw90 = Array((new BigNumber(19)).toNumber());
        _nw90[(_dafny.ZERO).toNumber()] = _620_v51;
        _nw90[(_dafny.ONE).toNumber()] = _620_v51;
        _nw90[(new BigNumber(2)).toNumber()] = _620_v51;
        _nw90[(new BigNumber(3)).toNumber()] = _620_v51;
        _nw90[(new BigNumber(4)).toNumber()] = _620_v51;
        _nw90[(new BigNumber(5)).toNumber()] = _620_v51;
        _nw90[(new BigNumber(6)).toNumber()] = _620_v51;
        _nw90[(new BigNumber(7)).toNumber()] = _620_v51;
        _nw90[(new BigNumber(8)).toNumber()] = _620_v51;
        _nw90[(new BigNumber(9)).toNumber()] = _620_v51;
        _nw90[(new BigNumber(10)).toNumber()] = _620_v51;
        _nw90[(new BigNumber(11)).toNumber()] = _620_v51;
        _nw90[(new BigNumber(12)).toNumber()] = _620_v51;
        _nw90[(new BigNumber(13)).toNumber()] = _620_v51;
        _nw90[(new BigNumber(14)).toNumber()] = _620_v51;
        _nw90[(new BigNumber(15)).toNumber()] = ((_this.f28) ? (new _dafny.CodePoint('v'.codePointAt(0))) : ((((_621_v52).contains(_module.__default.fm6(globalState))) ? ((_621_v52).get(_module.__default.fm6(globalState))) : (_620_v51))));
        _nw90[(new BigNumber(16)).toNumber()] = _620_v51;
        _nw90[(new BigNumber(17)).toNumber()] = _620_v51;
        _nw90[(new BigNumber(18)).toNumber()] = new _dafny.CodePoint('b'.codePointAt(0));
        _622_v53 = _nw90;
        let _index79 = _module.__default.safeIndex(new BigNumber(376), new BigNumber((_622_v53).length));
        (_622_v53)[_index79] = _620_v51;
        let _index80 = _module.__default.safeIndex(new BigNumber(376), new BigNumber((_622_v53).length));
        (_622_v53)[_index80] = _620_v51;
        let _623_v54;
        _623_v54 = _dafny.Seq.of((_dafny.ZERO).minus(p1));
        let _624_v55;
        _624_v55 = _dafny.Seq.of((_623_v54)[_module.__default.safeIndex(p0, new BigNumber((_623_v54).length))]);
        let _625_v56;
        _625_v56 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,p2);
        let _626_v57;
        _626_v57 = _module.D1.create_DC5(_this.f22, _625_v56, p1, _dafny.MultiSet.FromArray(_623_v54));
        let _627_v58;
        _627_v58 = _dafny.Map.Empty.slice().updateUnsafe(_620_v51,p2);
        let _628_v59;
        _628_v59 = _dafny.Seq.of(_module.__default.fm5(p1, (((_627_v58).contains((_622_v53)[_module.__default.safeIndex(new BigNumber(376), new BigNumber((_622_v53).length))])) ? ((_627_v58).get((_622_v53)[_module.__default.safeIndex(new BigNumber(376), new BigNumber((_622_v53).length))])) : (p0)), p2, globalState), _dafny.Seq.of(p1), _624_v55);
        let _629_v60;
        let _nw91 = Array((new BigNumber(17)).toNumber()).fill(_dafny.MultiSet.Empty);
        _629_v60 = _nw91;
        let _630_v61;
        let _nw92 = new _module.C1();
        _nw92.__ctor(_629_v60, _this.f28);
        _630_v61 = _nw92;
        let _631_v62;
        _631_v62 = _dafny.Map.Empty.slice().updateUnsafe((_562_v21)[_module.__default.safeIndex(new BigNumber(690), new BigNumber((_562_v21).length))],_630_v61);
        let _632_v63;
        _632_v63 = _dafny.MultiSet.fromElements((_630_v61).f25);
        let _633_v64;
        _633_v64 = _dafny.MultiSet.fromElements(_632_v63);
        let _634_v65;
        _634_v65 = _dafny.Map.Empty.slice().updateUnsafe(_563_v22,(((_576_v31).contains(p1)) ? ((_576_v31).get(p1)) : (_module.__default.fm12(new BigNumber((_631_v62).length), _632_v63, false, _633_v64, globalState))));
        let _635_v66;
        let _nw93 = Array((new BigNumber(27)).toNumber());
        _nw93[(_dafny.ZERO).toNumber()] = _623_v54;
        _nw93[(_dafny.ONE).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-664)), ((_636_p2) => function (_637_i9) {
          return _636_p2;
        })(p2));
        _nw93[(new BigNumber(2)).toNumber()] = _624_v55;
        _nw93[(new BigNumber(3)).toNumber()] = _dafny.Seq.of((_626_v57).dtor_cf9, p2, p0, new BigNumber((_623_v54).length));
        _nw93[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_623_v54, (_628_v59)[_module.__default.safeIndex(p1, new BigNumber((_628_v59).length))]);
        _nw93[(new BigNumber(5)).toNumber()] = _624_v55;
        _nw93[(new BigNumber(6)).toNumber()] = _623_v54;
        _nw93[(new BigNumber(7)).toNumber()] = _module.__default.fm5(p3, p0, new BigNumber(-369), globalState);
        _nw93[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_623_v54, _623_v54);
        _nw93[(new BigNumber(9)).toNumber()] = _624_v55;
        _nw93[(new BigNumber(10)).toNumber()] = _module.__default.fm5(p3, p3, new BigNumber(-379), globalState);
        _nw93[(new BigNumber(11)).toNumber()] = _dafny.Seq.of(new BigNumber(919));
        _nw93[(new BigNumber(12)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(588)), ((_638_p1) => function (_639_i10) {
          return _638_p1;
        })(p1));
        _nw93[(new BigNumber(13)).toNumber()] = _624_v55;
        _nw93[(new BigNumber(14)).toNumber()] = _623_v54;
        _nw93[(new BigNumber(15)).toNumber()] = _623_v54;
        _nw93[(new BigNumber(16)).toNumber()] = _624_v55;
        _nw93[(new BigNumber(17)).toNumber()] = _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(766)), ((_640_p2) => function (_641_i11) {
          return _640_p2;
        })(p2)), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(766)), ((_642_p2) => function (_643_i11) {
          return _642_p2;
        })(p2))).length)), new BigNumber((_634_v65).length)), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(766)), ((_644_p2) => function (_645_i11) {
          return _644_p2;
        })(p2)), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(766)), ((_646_p2) => function (_647_i11) {
          return _646_p2;
        })(p2))).length)), new BigNumber((_634_v65).length))).length)), p1), _module.__default.safeIndex(new BigNumber((_this.f21).length), new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(766)), ((_648_p2) => function (_649_i11) {
          return _648_p2;
        })(p2)), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(766)), ((_650_p2) => function (_651_i11) {
          return _650_p2;
        })(p2))).length)), new BigNumber((_634_v65).length)), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(766)), ((_652_p2) => function (_653_i11) {
          return _652_p2;
        })(p2)), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(766)), ((_654_p2) => function (_655_i11) {
          return _654_p2;
        })(p2))).length)), new BigNumber((_634_v65).length))).length)), p1)).length)), p0);
        _nw93[(new BigNumber(18)).toNumber()] = _623_v54;
        _nw93[(new BigNumber(19)).toNumber()] = _623_v54;
        _nw93[(new BigNumber(20)).toNumber()] = _624_v55;
        _nw93[(new BigNumber(21)).toNumber()] = _623_v54;
        _nw93[(new BigNumber(22)).toNumber()] = _dafny.Seq.of(p1, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f27,p1)).length), p3, p2);
        _nw93[(new BigNumber(23)).toNumber()] = _dafny.Seq.of(new BigNumber((_625_v56).length), p1);
        _nw93[(new BigNumber(24)).toNumber()] = _dafny.Seq.Concat(_623_v54, _624_v55);
        _nw93[(new BigNumber(25)).toNumber()] = _623_v54;
        _nw93[(new BigNumber(26)).toNumber()] = _dafny.Seq.Concat(_623_v54, _624_v55);
        _635_v66 = _nw93;
        let _index81 = _module.__default.safeIndex(new BigNumber(543), new BigNumber((_635_v66).length));
        (_635_v66)[_index81] = _dafny.Seq.of(p0);
        let _index82 = _module.__default.safeIndex(new BigNumber(543), new BigNumber((_635_v66).length));
        (_635_v66)[_index82] = _dafny.Seq.Concat(_624_v55, _624_v55);
        let _656_v67;
        _656_v67 = _module.D3.create_DC9((p3).plus(p1), p0, p0);
        let _source12 = _656_v67;
        if (_source12.is_DC8) {
          let _657___mcc_h9 = (_source12).cf13;
          let _658___mcc_h10 = (_source12).cf14;
          let _659___mcc_h11 = (_source12).cf15;
          let _660___mcc_h12 = (_source12).cf16;
          let _661_cf16 = _660___mcc_h12;
          let _662_cf15 = _659___mcc_h11;
          let _663_cf14 = _658___mcc_h10;
          let _664_cf13 = _657___mcc_h9;
          (globalState).f3 = (p3).multipliedBy(_module.__default.fm6(globalState));
          let _665_v68;
          _665_v68 = _dafny.Map.Empty.slice().updateUnsafe(p1,_664_cf13);
          let _666_v69;
          _666_v69 = _dafny.Set.fromElements(_this.f21);
          let _667_v70;
          let _nw94 = new _module.C0();
          _nw94.__ctor(_666_v69);
          _667_v70 = _nw94;
          let _668_v71;
          _668_v71 = _dafny.Seq.of(_667_v70, _667_v70);
          let _669_v72;
          _669_v72 = _dafny.Seq.of(_668_v71, _668_v71, _dafny.Seq.of(_667_v70, _667_v70));
          let _670_v73;
          _670_v73 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_669_v72)).cardinality()), _664_cf13, _664_cf13);
          let _pat_let_tv32 = _625_v56;
          let _pat_let_tv33 = _625_v56;
          let _pat_let_tv34 = p1;
          let _pat_let_tv35 = _625_v56;
          let _pat_let_tv36 = _670_v73;
          let _pat_let_tv37 = p2;
          let _rhs84 = (((_665_v68).contains((new BigNumber((_this.f21).length)).multipliedBy(p1))) ? ((_665_v68).get((new BigNumber((_this.f21).length)).multipliedBy(p1))) : ((_dafny.ZERO).minus(p3)));
          let _rhs85 = _this.f22;
          let _rhs86 = function (_pat_let10_0) {
            return function (_671_dt__update__tmp_h1) {
              return function (_pat_let11_0) {
                return function (_672_dt__update_hcf8_h0) {
                  return function (_pat_let12_0) {
                    return function (_673_dt__update_hcf10_h0) {
                      return function (_pat_let13_0) {
                        return function (_674_dt__update_hcf9_h0) {
                          return _module.D1.create_DC5((_671_dt__update__tmp_h1).dtor_cf7, _672_dt__update_hcf8_h0, _674_dt__update_hcf9_h0, _673_dt__update_hcf10_h0);
                        }(_pat_let13_0);
                      }(_pat_let_tv37);
                    }(_pat_let12_0);
                  }(_pat_let_tv36);
                }(_pat_let11_0);
              }((_pat_let_tv32).update(_this.f28, (((_pat_let_tv35).contains(false)) ? ((_pat_let_tv33).get(false)) : (_pat_let_tv34))));
            }(_pat_let10_0);
          }(_662_cf15);
          let _rhs87 = (_562_v21)[_module.__default.safeIndex(new BigNumber(690), new BigNumber((_562_v21).length))];
          let _rhs88 = _this.f22;
          let _lhs69 = globalState;
          let _lhs70 = _this;
          _lhs69.f20 = _rhs84;
          _661_cf16 = _rhs85;
          _626_v57 = _rhs86;
          _661_cf16 = _rhs87;
          _lhs70.f22 = _rhs88;
          let _675_v74;
          _675_v74 = _dafny.Map.Empty.slice().updateUnsafe(!(_661_cf16),!(new BigNumber((_665_v68).length)).isEqualTo(p1));
          let _rhs89 = _675_v74;
          let _rhs90 = _667_v70;
          let _rhs91 = _661_cf16;
          let _rhs92 = (_675_v74).Merge(_675_v74);
          let _lhs71 = _this;
          _675_v74 = _rhs89;
          _667_v70 = _rhs90;
          _lhs71.f28 = _rhs91;
          _675_v74 = _rhs92;
          let _rhs93 = ((_562_v21)[_module.__default.safeIndex(new BigNumber(690), new BigNumber((_562_v21).length))]) && (((_this.f28) ? ((_630_v61).f25) : (_this.f28)));
          let _rhs94 = (_this).f27;
          let _lhs72 = _this;
          let _lhs73 = _this;
          _lhs72.f28 = _rhs93;
          _lhs73.f22 = _rhs94;
        } else if (_source12.is_DC9) {
          let _676___mcc_h13 = (_source12).cf17;
          let _677___mcc_h14 = (_source12).cf18;
          let _678___mcc_h15 = (_source12).cf19;
          let _679_cf19 = _678___mcc_h15;
          let _680_cf18 = _677___mcc_h14;
          let _681_cf17 = _676___mcc_h13;
          let _682_v75;
          _682_v75 = _dafny.Seq.of(_629_v60);
          let _683_v76;
          let _nw95 = new _module.C1();
          _nw95.__ctor((_682_v75)[_module.__default.safeIndex(_680_cf18, new BigNumber((_682_v75).length))], ((_562_v21)[_module.__default.safeIndex(new BigNumber(690), new BigNumber((_562_v21).length))]) || ((_this).f27));
          _683_v76 = _nw95;
          let _684_v77;
          let _out5;
          _out5 = _module.__default.m0(_562_v21, _this.f22, globalState);
          _684_v77 = _out5;
          let _685_v78;
          let _out6;
          _out6 = _module.__default.m0(_562_v21, (_630_v61).f25, globalState);
          _685_v78 = _out6;
          (_this).f22 = _684_v77;
        } else {
          let _686___mcc_h16 = (_source12).cf12;
          let _687_cf12 = _686___mcc_h16;
          (globalState).f20 = p3;
          let _688_v79;
          _688_v79 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,((_this.f22) ? (true) : ((_630_v61).f25)));
          _688_v79 = (_688_v79).update((_this).f27, (_630_v61).f25);
          let _689_v80;
          let _nw96 = Array((new BigNumber(11)).toNumber());
          _nw96[(_dafny.ZERO).toNumber()] = _this.f21;
          _nw96[(_dafny.ONE).toNumber()] = _dafny.Seq.update(_module.__default.fm11(_this.f22, globalState), _module.__default.safeIndex(new BigNumber(668), new BigNumber((_module.__default.fm11(_this.f22, globalState)).length)), _620_v51);
          _nw96[(new BigNumber(2)).toNumber()] = _dafny.Seq.UnicodeFromString("jnsquovwx");
          _nw96[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_this.f21, _dafny.Seq.UnicodeFromString("ittgfr"));
          _nw96[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("cqonwvi"), _this.f21);
          _nw96[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("xwfxaqtqq");
          _nw96[(new BigNumber(6)).toNumber()] = _dafny.Seq.UnicodeFromString("mt");
          _nw96[(new BigNumber(7)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(242)), function (_690_i12) {
            return new _dafny.CodePoint('v'.codePointAt(0));
          });
          _nw96[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-696)), ((_691_v51) => function (_692_i13) {
            return _691_v51;
          })(_620_v51)), _this.f21);
          _nw96[(new BigNumber(9)).toNumber()] = _this.f21;
          _nw96[(new BigNumber(10)).toNumber()] = _this.f21;
          _689_v80 = _nw96;
          let _index83 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_689_v80).length));
          (_689_v80)[_index83] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-122)), ((_693_v51) => function (_694_i14) {
            return _693_v51;
          })(_620_v51)), _dafny.Seq.UnicodeFromString("kmvw"));
          let _index84 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_689_v80).length));
          (_689_v80)[_index84] = _this.f21;
          _624_v55 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(301)), function (_695_i15) {
            return new BigNumber(222);
          });
        }
        let _index85 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_562_v21).length));
        (_562_v21)[_index85] = true;
        let _696_v81;
        _696_v81 = _dafny.Seq.of(_dafny.MultiSet.fromElements(_632_v63));
        _631_v62 = (_631_v62).update(_module.__default.fm12(p2, _dafny.MultiSet.fromElements((_this).f27, _this.f22, (_630_v61).f25, _module.__default.fm12(p1, _632_v63, (_this).f27, _dafny.MultiSet.fromElements(_632_v63), globalState), _module.__default.fm12(new BigNumber((_this.f21).length), _dafny.MultiSet.fromElements(_this.f22, _this.f28), _this.f22, _633_v64, globalState)), _this.f22, (_696_v81)[_module.__default.safeIndex(new BigNumber(-979), new BigNumber((_696_v81).length))], globalState), _630_v61);
      }
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_561_v20).length))) {
        let _697_i16 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_697_i16)) && ((_697_i16).isLessThan(new BigNumber((_561_v20).length))))) {
          (_561_v20)[(_697_i16)] = (_697_i16).multipliedBy(p3);
        }
      }
      let _698_v82;
      let _nw97 = Array((new BigNumber(24)).toNumber()).fill([]);
      _698_v82 = _nw97;
      r0 = _698_v82;
      return r0;
    }
    get f27() {
      let _this = this;
      return _this._f27;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f21 = _dafny.Seq.UnicodeFromString("");
      this._f22 = false;
      this._f23 = _dafny.Seq.of();
      this.f29 = _dafny.Set.Empty;
      this._f30 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f21() {
      let _this = this;
      return _this._f21;
    };
    set f21(value) {
      let _this = this;
      _this._f21 = value;
    };
    get f22() {
      let _this = this;
      return _this._f22;
    };
    set f22(value) {
      let _this = this;
      _this._f22 = value;
    };
    get f23() {
      let _this = this;
      return _this._f23;
    };
    set f23(value) {
      let _this = this;
      _this._f23 = value;
    };
    __ctor(f29, f30, f23, f21, f22) {
      let _this = this;
      (_this).f29 = f29;
      (_this)._f30 = f30;
      (_this)._f23 = f23;
      (_this)._f21 = f21;
      (_this)._f22 = f22;
      return;
    }
    fm2(p0, globalState) {
      let _this = this;
      if ((_this).f30) {
        return false;
      } else {
        return _this.f22;
      }
    };
    fm0(globalState) {
      let _this = this;
      return (_dafny.Set.fromElements(false)).Intersect((_dafny.Set.fromElements(_this.f22, (_this).f30)).Intersect(_dafny.Set.fromElements((_this).f30)));
    };
    fm1(p0, p1, p2, globalState) {
      let _this = this;
      return _this.f22;
    };
    get f30() {
      let _this = this;
      return _this._f30;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
