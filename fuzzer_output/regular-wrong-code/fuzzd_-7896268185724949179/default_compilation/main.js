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
    static fm0(p0, p1, p2, p3, globalState) {
      return (new BigNumber((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(-842), new BigNumber(432))) {
          let _0_v0 = _compr_0;
          if (((new BigNumber(-842)).isLessThanOrEqualTo(_0_v0)) && ((_0_v0).isLessThan(new BigNumber(432)))) {
            _coll0.push([_module.__default.safeDivisionInt(_0_v0, new BigNumber(308)),new BigNumber(166)]);
          }
        }
        return _coll0;
      }()).length)).plus(new BigNumber(((_module.D13.create_DC26(_dafny.Seq.UnicodeFromString("gtqau"), true, true)).dtor_cf33).length));
    };
    static fm1(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("juk"), _dafny.Seq.UnicodeFromString("liutef")), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(902)), function (_1_i0) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("olrx")));
    };
    static fm2(p0, globalState) {
      return !(((false) ? (true) : (((false) ? (false) : (true)))));
    };
    static fm3(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("cacr")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length), new BigNumber(792))).length),!(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(581),false))).Merge(function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of _dafny.IntegerRange(new BigNumber(310), new BigNumber(336))) {
          let _2_v0 = _compr_1;
          if (((new BigNumber(310)).isLessThanOrEqualTo(_2_v0)) && ((_2_v0).isLessThan(new BigNumber(336)))) {
            _coll1.push([_module.__default.safeDivisionInt(_2_v0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-856)), function (_4_i0) {
              return new _dafny.CodePoint('w'.codePointAt(0));
            })).length)),(_module.D24.create_DC64(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(468)), function (_3_i1) {
  return new _dafny.CodePoint('x'.codePointAt(0));
})).length), true, false)).dtor_cf93]);
          }
        }
        return _coll1;
      }());
    };
    static fm4(p0, p1, globalState) {
      if (((true) ? (false) : (true))) {
        return _module.D0.create_DC0(!(true));
      } else {
        return _module.D0.create_DC1(_module.D0.create_DC0(false));
      }
    };
    static fm11(p0, globalState) {
      let _source0 = _module.D20.create_DC49(function () {
  let _coll2 = new _dafny.Map();
  for (const _compr_2 of _dafny.IntegerRange(new BigNumber(500), new BigNumber(743))) {
    let _5_v0 = _compr_2;
    if (((new BigNumber(500)).isLessThanOrEqualTo(_5_v0)) && ((_5_v0).isLessThan(new BigNumber(743)))) {
      _coll2.push([_module.__default.safeDivisionInt(_5_v0, new BigNumber(893)),new _dafny.CodePoint('v'.codePointAt(0))]);
    }
  }
  return _coll2;
}());
      if (_source0.is_DC50) {
        let _6___mcc_h0 = (_source0).cf68;
        let _7___mcc_h1 = (_source0).cf69;
        let _8___mcc_h2 = (_source0).cf70;
        let _9_cf70 = _8___mcc_h2;
        let _10_cf69 = _7___mcc_h1;
        let _11_cf68 = _6___mcc_h0;
        return _dafny.Seq.of(_dafny.MultiSet.fromElements(_10_cf69), _dafny.MultiSet.fromElements(false), _dafny.MultiSet.fromElements(_9_cf70));
      } else if (_source0.is_DC51) {
        return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.fromElements(false)), _dafny.Seq.of(_dafny.MultiSet.fromElements(true, true)));
      } else {
        let _12___mcc_h3 = (_source0).cf67;
        let _13_cf67 = _12___mcc_h3;
        return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.MultiSet.fromElements(true), _dafny.MultiSet.fromElements(true)), _dafny.Seq.of(_dafny.MultiSet.fromElements(true)));
      }
    };
    static fm12(p0, p1, p2, globalState) {
      return (_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(872))).cardinality()), new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('v'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)))).cardinality()), new BigNumber(-229), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true, !(true))).length))).length), new BigNumber(358)), _dafny.Seq.of(new BigNumber(636))))).Difference(_dafny.MultiSet.fromElements(new BigNumber(864), new BigNumber((_dafny.Seq.UnicodeFromString("mci")).length)));
    };
    static fm13(p0, p1, p2, p3, globalState) {
      return ((_dafny.MultiSet.fromElements(new _dafny.CodePoint('h'.codePointAt(0)), (_module.D3.create_DC6(_dafny.Seq.Create(_module.__default.abs(new BigNumber(743)), function (_14_i0) {
  return new _dafny.CodePoint('w'.codePointAt(0));
}), new _dafny.CodePoint('c'.codePointAt(0)), false)).dtor_cf7)).Intersect(_dafny.MultiSet.fromElements(new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('q'.codePointAt(0))))).Union(_dafny.MultiSet.fromElements(new _dafny.CodePoint('a'.codePointAt(0))));
    };
    static fm14(p0, p1, globalState) {
      return (((true) ? (_dafny.Set.fromElements(false)) : (_dafny.Set.fromElements(true)))).Difference((_dafny.Set.fromElements(true)).Intersect(_dafny.Set.fromElements(false)));
    };
    static fm15(p0, p1, globalState) {
      return new _dafny.CodePoint('b'.codePointAt(0));
    };
    static fm16(p0, p1, globalState) {
      return _module.D4.create_DC7((_dafny.ZERO).minus((_module.D19.create_DC47(_module.D1.create_DC2(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false)).length),true)), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(352)), function (_15_i0) {
  return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_module.D17.create_DC39(new BigNumber(838), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(353))).length),new BigNumber((function () {
  let _coll3 = new _dafny.Map();
  for (const _compr_3 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("tynvdcxoj")).length)))).cardinality()))).Keys.Elements) {
    let _16_v0 = _compr_3;
    if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("tynvdcxoj")).length)))).cardinality()))).contains(_16_v0)) {
      _coll3.push([_16_v0,true]);
    }
  }
  return _coll3;
}()).length)), !(false), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true)).length),_module.D13.create_DC25(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-51),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(52))).cardinality())))), new BigNumber(358)))).length);
})).length))).cardinality()), false)).dtor_cf63));
    };
    static fm18(p0, p1, p2, globalState) {
      if ((new BigNumber(-845)).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("tfjbncjt")).length))).length))) {
        return _dafny.Seq.of(false);
      } else {
        return _dafny.Seq.Concat(_dafny.Seq.of(true, false), _dafny.Seq.of(true));
      }
    };
    static fm19(p0, p1, p2, p3, globalState) {
      return ((_dafny.Set.fromElements(false)).Union(_dafny.Set.fromElements(true))).Intersect(_dafny.Set.fromElements(false, true));
    };
    static fm20(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,true)).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,true)));
    };
    static fm21(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(171)), function (_17_i0) {
        return new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(-380), new BigNumber((_dafny.Seq.UnicodeFromString("hf")).length)))).cardinality());
      }), _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(41)), function (_18_i1) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      })).length)));
    };
    static fm22(p0, p1, p2, globalState) {
      return _dafny.Seq.of(_dafny.Seq.of(new BigNumber(-136)), _dafny.Seq.of(new BigNumber(-24)), _dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(906))).length), new BigNumber((_dafny.Seq.UnicodeFromString("xtfitcgkw")).length)), _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(44)), function (_19_i0) {
        return new BigNumber(-784);
      })).length))).length))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(844)), function (_20_i1) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("vpemvormm")).length);
      }));
    };
    static fm23(p0, p1, p2, p3, globalState) {
      if (!(_dafny.Seq.IsPrefixOf(_dafny.Seq.of(true), _dafny.Seq.of(false, true)))) {
        return _module.D13.create_DC26(_dafny.Seq.UnicodeFromString("vmuja"), !(!(false)), false);
      } else {
        return _module.D13.create_DC26((_module.D13.create_DC26(_dafny.Seq.UnicodeFromString("lhaqxlf"), true, false)).dtor_cf33, true, true);
      }
    };
    static fm24(p0, p1, p2, p3, globalState) {
      return ((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),!(true)))).Intersect(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('l'.codePointAt(0)),!(true)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),false), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),false)))).Intersect((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),false), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),true), function () {
        let _coll4 = new _dafny.Map();
        for (const _compr_4 of (_dafny.Set.fromElements(new _dafny.CodePoint('c'.codePointAt(0)))).Elements) {
          let _21_v0 = _compr_4;
          if ((_dafny.Set.fromElements(new _dafny.CodePoint('c'.codePointAt(0)))).contains(_21_v0)) {
            _coll4.push([_21_v0,false]);
          }
        }
        return _coll4;
      }(), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),false))).Union(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(922)), function (_22_i0) {
        return function () {
          let _coll5 = new _dafny.Map();
          for (const _compr_5 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('t'.codePointAt(0)),true)).Keys.Elements) {
            let _23_v1 = _compr_5;
            if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('t'.codePointAt(0)),true)).contains(_23_v1)) {
              _coll5.push([_23_v1,false]);
            }
          }
          return _coll5;
        }();
      }))));
    };
    static fm25(p0, p1, p2, p3, globalState) {
      return _module.D3.create_DC6(_dafny.Seq.UnicodeFromString("qs"), new _dafny.CodePoint('p'.codePointAt(0)), (_dafny.MultiSet.fromElements(false)).IsProperSubsetOf(_dafny.MultiSet.fromElements(true)));
    };
    static fm26(p0, p1, globalState) {
      return ((_dafny.Set.fromElements(new BigNumber(8), new BigNumber(496), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(-312), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(582), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.UnicodeFromString("fre")).length))).length), new BigNumber(784))).cardinality())))).cardinality()))).Difference(_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber(((_dafny.Set.fromElements(true, false))).length)), new BigNumber(386)))).Difference(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(false, true)).cardinality()))).length), new BigNumber(190), new BigNumber(-913)));
    };
    static fm27(p0, p1, globalState) {
      if ((_dafny.Set.fromElements(false, false, true)).IsSubsetOf(_dafny.Set.fromElements(true))) {
        return new _dafny.CodePoint('v'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('p'.codePointAt(0));
      }
    };
    static fm28(p0, p1, p2, p3, globalState) {
      let _source1 = _module.D18.create_DC42(_dafny.Map.Empty.slice().updateUnsafe(true,false));
      if (_source1.is_DC43) {
        return _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(642)), function (_24_i0) {
          return new _dafny.CodePoint('n'.codePointAt(0));
        })).length), new BigNumber(-885), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length)), new BigNumber(538), new BigNumber(588))).length));
      } else if (_source1.is_DC44) {
        return _dafny.MultiSet.fromElements(new BigNumber(158), new BigNumber((_dafny.Seq.of(false)).length));
      } else {
        let _25___mcc_h0 = (_source1).cf58;
        let _26_cf58 = _25___mcc_h0;
        return _dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber(-301)), new BigNumber(-858));
      }
    };
    static fm30(p0, p1, p2, p3, globalState) {
      return ((_dafny.Set.fromElements(new BigNumber(915), new BigNumber(103))).Intersect(_dafny.Set.fromElements(new BigNumber(833)))).Intersect(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(368)), function (_27_i0) {
        return new BigNumber((_dafny.Seq.UnicodeFromString("rsvm")).length);
      })).length)));
    };
    static fm31(p0, p1, globalState) {
      return ((_dafny.Set.fromElements(true)).Intersect(_dafny.Set.fromElements(true))).Intersect(_dafny.Set.fromElements(false, true));
    };
    static fm32(p0, p1, p2, globalState) {
      if ((true) === (true)) {
        return _module.D14.create_DC28(_dafny.Set.fromElements(new BigNumber(396)));
      } else if (true) {
        return _module.D14.create_DC28(function () {
  let _coll6 = new _dafny.Set();
  for (const _compr_6 of _dafny.IntegerRange(new BigNumber(88), new BigNumber(346))) {
    let _28_v0 = _compr_6;
    if (((new BigNumber(88)).isLessThanOrEqualTo(_28_v0)) && ((_28_v0).isLessThan(new BigNumber(346)))) {
      _coll6.add((_28_v0).plus(new BigNumber((_dafny.Seq.of(new BigNumber(610))).length)));
    }
  }
  return _coll6;
}());
      } else {
        return _module.D14.create_DC28(function () {
  let _coll7 = new _dafny.Set();
  for (const _compr_7 of _dafny.IntegerRange(new BigNumber(896), new BigNumber(614))) {
    let _29_v1 = _compr_7;
    if (((new BigNumber(896)).isLessThanOrEqualTo(_29_v1)) && ((_29_v1).isLessThan(new BigNumber(614)))) {
      _coll7.add((_29_v1).minus(new BigNumber(-280)));
    }
  }
  return _coll7;
}());
      }
    };
    static fm33(globalState) {
      return (_module.D25.create_DC67(_dafny.MultiSet.fromElements(new BigNumber(-892), new BigNumber(188)))).dtor_cf98;
    };
    static fm34(globalState) {
      return new _dafny.CodePoint('u'.codePointAt(0));
    };
    static fm35(p0, p1, globalState) {
      let _source2 = _module.D21.create_DC54(new BigNumber(262), new BigNumber(529), !(!(true)));
      if (_source2.is_DC53) {
        let _30___mcc_h0 = (_source2).cf72;
        let _31___mcc_h1 = (_source2).cf73;
        let _32___mcc_h2 = (_source2).cf74;
        let _33___mcc_h3 = (_source2).cf75;
        let _34_cf75 = _33___mcc_h3;
        let _35_cf74 = _32___mcc_h2;
        let _36_cf73 = _31___mcc_h1;
        let _37_cf72 = _30___mcc_h0;
        return _dafny.Map.Empty.slice().updateUnsafe(_36_cf73,new BigNumber(449));
      } else if (_source2.is_DC54) {
        let _38___mcc_h4 = (_source2).cf76;
        let _39___mcc_h5 = (_source2).cf77;
        let _40___mcc_h6 = (_source2).cf78;
        let _41_cf78 = _40___mcc_h6;
        let _42_cf77 = _39___mcc_h5;
        let _43_cf76 = _38___mcc_h4;
        return (_dafny.Map.Empty.slice().updateUnsafe(_41_cf78,_42_cf77)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_41_cf78,_42_cf77));
      } else if (_source2.is_DC52) {
        let _44___mcc_h7 = (_source2).cf71;
        let _45_cf71 = _44___mcc_h7;
        if (true) {
          return _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((function () {
            let _coll8 = new _dafny.Map();
            for (const _compr_8 of (_dafny.Seq.of(new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)))).Elements) {
              let _46_v0 = _compr_8;
              if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0))), _46_v0)) {
                _coll8.push([_46_v0,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-442)), function (_47_i0) {
                  return new _dafny.CodePoint('c'.codePointAt(0));
                })).length)]);
              }
            }
            return _coll8;
          }()).length));
        } else {
          return _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(278));
        }
      } else {
        let _48___mcc_h8 = (_source2).cf79;
        let _49_cf79 = _48___mcc_h8;
        return (_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber((_dafny.Seq.UnicodeFromString("abldcs")).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber(542)));
      }
    };
    static fm36(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false),_dafny.Seq.Create(_module.__default.abs(new BigNumber(-663)), function (_50_i0) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      }));
    };
    static fm37(globalState) {
      return _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-504)));
    };
    static fm38(globalState) {
      return (_dafny.Set.fromElements(_dafny.Set.fromElements(new _dafny.CodePoint('i'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('v'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('e'.codePointAt(0))))).Intersect(function () {
        let _coll9 = new _dafny.Set();
        for (const _compr_9 of (function () {
          let _coll10 = new _dafny.Set();
          for (const _compr_10 of (_dafny.Seq.of(_dafny.Set.fromElements(new _dafny.CodePoint('a'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('j'.codePointAt(0))))).Elements) {
            let _51_v0 = _compr_10;
            if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Set.fromElements(new _dafny.CodePoint('a'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('j'.codePointAt(0)))), _51_v0)) {
              _coll10.add(_51_v0);
            }
          }
          return _coll10;
        }()).Elements) {
          let _52_v1 = _compr_9;
          if ((function () {
            let _coll11 = new _dafny.Set();
            for (const _compr_11 of (_dafny.Seq.of(_dafny.Set.fromElements(new _dafny.CodePoint('a'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('j'.codePointAt(0))))).Elements) {
              let _53_v0 = _compr_11;
              if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Set.fromElements(new _dafny.CodePoint('a'.codePointAt(0))), _dafny.Set.fromElements(new _dafny.CodePoint('j'.codePointAt(0)))), _53_v0)) {
                _coll11.add(_53_v0);
              }
            }
            return _coll11;
          }()).contains(_52_v1)) {
            _coll9.add(_52_v1);
          }
        }
        return _coll9;
      }());
    };
    static fm39(p0, p1, p2, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-361)), function (_54_i0) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      })).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.Create(_module.__default.abs(new BigNumber(336)), function (_55_i1) {
        return new BigNumber(476);
      }))).length), new BigNumber((_dafny.Seq.of(new BigNumber(759))).length)),_dafny.Seq.of(!(true)))).Merge(function () {
        let _coll12 = new _dafny.Map();
        for (const _compr_12 of (_dafny.Seq.of(_dafny.Seq.of(new BigNumber(-784)), _dafny.Seq.of(new BigNumber(508), (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)))).length), new BigNumber(486))).length))))).Elements) {
          let _56_v0 = _compr_12;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.of(new BigNumber(-784)), _dafny.Seq.of(new BigNumber(508), (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)))).length), new BigNumber(486))).length)))), _56_v0)) {
            _coll12.push([_56_v0,_dafny.Seq.of(false)]);
          }
        }
        return _coll12;
      }())).Merge((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length)))).length)),_dafny.Seq.of(true, false))).Merge(function () {
        let _coll13 = new _dafny.Map();
        for (const _compr_13 of (_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber(317)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(267)), function (_57_i2) {
          return new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(960))).cardinality()))).cardinality());
        }))).Elements) {
          let _58_v1 = _compr_13;
          if ((_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber(317)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(267)), function (_57_i2) {
            return new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(960))).cardinality()))).cardinality());
          }))).contains(_58_v1)) {
            _coll13.push([_58_v1,_dafny.Seq.of(true)]);
          }
        }
        return _coll13;
      }()));
    };
    static fm40(globalState) {
      return _module.D16.create_DC36(true, (_dafny.MultiSet.fromElements(false, false)).Union(_dafny.MultiSet.fromElements(true)), !(new BigNumber(-574)).isEqualTo(new BigNumber(499)));
    };
    static fm41(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.of(!(true));
    };
    static fm42(globalState) {
      return _module.D4.create_DC10((_dafny.ZERO).minus((_dafny.ZERO).minus((_module.D20.create_DC50(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("uh")).length))).length), false, true)).dtor_cf68)), (new BigNumber((_dafny.Seq.UnicodeFromString("owettajm")).length)).multipliedBy(new BigNumber((_dafny.Seq.of(_dafny.Seq.of(true))).length)), ((!(false)) ? (false) : (!(!(false)))), !(true), new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()));
    };
    static fm43(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(!(true),_module.D4.create_DC10(new BigNumber((_dafny.Seq.UnicodeFromString("memjruqa")).length), new BigNumber(-160), true, true, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(404), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(885),function () {
  let _coll14 = new _dafny.Set();
  for (const _compr_14 of ((_module.D26.create_DC71(_dafny.MultiSet.fromElements(new _dafny.CodePoint('p'.codePointAt(0))))).dtor_cf100).Elements) {
    let _59_v0 = _compr_14;
    if (((_module.D26.create_DC71(_dafny.MultiSet.fromElements(new _dafny.CodePoint('p'.codePointAt(0))))).dtor_cf100).contains(_59_v0)) {
      _coll14.add(_59_v0);
    }
  }
  return _coll14;
}())).length), new BigNumber(955), new BigNumber(236))).cardinality())))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,_module.D4.create_DC10(new BigNumber(285), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("exg")).length)), true, true, new BigNumber(113)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_module.D4.create_DC10(new BigNumber(655), new BigNumber((_dafny.Seq.UnicodeFromString("c")).length), true, true, new BigNumber(849)))));
    };
    static fm44(globalState) {
      if (!(true)) {
        return _module.D13.create_DC27(true);
      } else {
        return _module.D13.create_DC27(true);
      }
    };
    static fm45(globalState) {
      if (false) {
        return _module.D13.create_DC25(function () {
  let _coll15 = new _dafny.Map();
  for (const _compr_15 of _dafny.IntegerRange(new BigNumber(190), new BigNumber(-939))) {
    let _60_v0 = _compr_15;
    if (((new BigNumber(190)).isLessThanOrEqualTo(_60_v0)) && ((_60_v0).isLessThan(new BigNumber(-939)))) {
      _coll15.push([(_60_v0).minus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(743), new BigNumber(-433))).length), new BigNumber((_dafny.Seq.of(new BigNumber(226), new BigNumber(423))).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(622)), function (_61_i0) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      })).length)))).cardinality())),new BigNumber(-545)]);
    }
  }
  return _coll15;
}());
      } else {
        return _module.D13.create_DC25(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length))),new BigNumber(652)));
      }
    };
    static fm46(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(323)), function (_62_i0) {
        return _module.D14.create_DC30(_dafny.Seq.of(new BigNumber(691), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-832)), function (_63_i1) {
  return new _dafny.CodePoint('l'.codePointAt(0));
})).length), new BigNumber((_dafny.Seq.UnicodeFromString("xxqknse")).length)), _dafny.Seq.of(false));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-332)), function (_64_i2) {
        return _module.D14.create_DC30(_dafny.Seq.of(new BigNumber(94)), _dafny.Seq.of(!(false), true, true));
      }));
    };
    static fm47(p0, globalState) {
      return ((true) ? (_dafny.Set.fromElements(true, false)) : (_dafny.Set.fromElements(false)));
    };
    static fm48(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(-648)).minus(new BigNumber(583)),(function () {
        let _coll16 = new _dafny.Map();
        for (const _compr_16 of (_dafny.Seq.of(new _dafny.CodePoint('a'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)))).Elements) {
          let _65_v0 = _compr_16;
          if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('a'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0))), _65_v0)) {
            _coll16.push([_65_v0,true]);
          }
        }
        return _coll16;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),false)));
    };
    static fm49(p0, p1, p2, p3, globalState) {
      return _module.D19.create_DC46(!(_dafny.Map.Empty.slice().updateUnsafe(_module.D23.create_DC60(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Map.Empty.slice().updateUnsafe(false,!(!(false))))),new BigNumber((_dafny.Seq.UnicodeFromString("unbxcg")).length))).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D23.create_DC60(_dafny.Map.Empty.slice().updateUnsafe(!(false),_dafny.Map.Empty.slice().updateUnsafe(!(true),!(false)))),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(366)), function (_66_i0) {
  return new _dafny.CodePoint('o'.codePointAt(0));
})).length))), (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_module.D14.create_DC28(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(true)).length))))).length),false)).length)).multipliedBy(new BigNumber(711)));
    };
    static fm50(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(!(false),_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("obwgo")).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.Create(_module.__default.abs(new BigNumber(40)), function (_67_i0) {
        return new BigNumber((_dafny.Set.fromElements(false, !(true))).length);
      })));
    };
    static fm51(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(new BigNumber(428))).length))).length), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("dgjnixnjt")).length), _dafny.ZERO)).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(532),new BigNumber((_dafny.Seq.UnicodeFromString("umwrjbo")).length))).length),true)).length))).length),true)), _dafny.Seq.of(function () {
        let _coll17 = new _dafny.Map();
        for (const _compr_17 of _dafny.IntegerRange(new BigNumber(620), new BigNumber(233))) {
          let _68_v0 = _compr_17;
          if (((new BigNumber(620)).isLessThanOrEqualTo(_68_v0)) && ((_68_v0).isLessThan(new BigNumber(233)))) {
            _coll17.push([(_68_v0).minus((_dafny.ZERO).minus(new BigNumber(-232))),!(true)]);
          }
        }
        return _coll17;
      }()));
    };
    static fm52(p0, p1, globalState) {
      return (_dafny.MultiSet.FromArray(_dafny.Seq.of(!(!(false)), false, true))).Intersect((_dafny.MultiSet.fromElements(!(false))).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(false))));
    };
    static fm53(p0, globalState) {
      return ((((_module.D23.create_DC61(_dafny.Seq.of(true), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('b'.codePointAt(0)))).length), false)).dtor_cf88) ? (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("yuclgsvm"), _dafny.Seq.UnicodeFromString("l"))) : (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("ndjpjd"), _dafny.Seq.UnicodeFromString("n"), _dafny.Seq.UnicodeFromString("jgu"), _dafny.Seq.UnicodeFromString("p"))))).Intersect(_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(986)), function (_69_i0) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("clwvk")));
    };
    static fm54(globalState) {
      return (_module.D13.create_DC25(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(14)), function (_70_i0) {
  return (_module.D3.create_DC5(new _dafny.CodePoint('n'.codePointAt(0)))).dtor_cf5;
})).length),new BigNumber((_dafny.Seq.of(false)).length)))).dtor_cf32;
    };
    static fm55(p0, p1, p2, globalState) {
      return (_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("fvrqpnj"), _dafny.Seq.UnicodeFromString("db"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(768)), function (_71_i0) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("hs"))).Difference((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("sfwcvbt"))).Intersect(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("dntsofjge"), _dafny.Seq.UnicodeFromString("ctwuqpd"))));
    };
    static fm56(p0, p1, p2, globalState) {
      return ((_dafny.MultiSet.fromElements(_module.D6.create_DC13(new BigNumber(613), new BigNumber(12)))).Difference(_dafny.MultiSet.fromElements(_module.D6.create_DC13(new BigNumber((_dafny.Seq.UnicodeFromString("itggvjgu")).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(36)), function (_72_i0) {
  return new _dafny.CodePoint('v'.codePointAt(0));
})).length))))).Union(_dafny.MultiSet.fromElements(_module.D6.create_DC13(new BigNumber((_dafny.Seq.UnicodeFromString("wmqdtf")).length), new BigNumber(-359)), _module.D6.create_DC13((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(true, false, true)).length)), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(false))).length)))));
    };
    static fm57(globalState) {
      if ((_dafny.MultiSet.fromElements(new BigNumber(564))).IsSubsetOf(_dafny.MultiSet.fromElements(new BigNumber(519), new BigNumber((_dafny.Set.fromElements(new BigNumber(43))).length), new BigNumber((_dafny.Set.fromElements(new BigNumber(934), new BigNumber(879))).length)))) {
        return _module.D16.create_DC35(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(-867))).length), (_module.D22.create_DC57(new BigNumber(-174), new BigNumber((_dafny.Set.fromElements(false)).length), false)).dtor_cf81),_dafny.Seq.of(true)));
      } else {
        return _module.D16.create_DC35(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(553),true)).length)),_dafny.Seq.of(false, false)));
      }
    };
    static fm58(p0, globalState) {
      return _module.D12.create_DC24(_module.D12.create_DC22(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true)).length)))).length))));
    };
    static fm59(p0, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-859)), function (_73_i0) {
        return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("mjaaqd"), _dafny.Seq.UnicodeFromString("bojv"));
      });
    };
    static fm60(globalState) {
      return (function () {
        let _coll18 = new _dafny.Map();
        for (const _compr_18 of (_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("wodnahax")).length),new _dafny.CodePoint('s'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
          let _coll19 = new _dafny.Set();
          for (const _compr_19 of _dafny.IntegerRange(new BigNumber(805), new BigNumber(429))) {
            let _74_v1 = _compr_19;
            if (((new BigNumber(805)).isLessThanOrEqualTo(_74_v1)) && ((_74_v1).isLessThan(new BigNumber(429)))) {
              _coll19.add((_74_v1).plus(new BigNumber(-564)));
            }
          }
          return _coll19;
        }()).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-498),new BigNumber((_dafny.Seq.of(true)).length))).length))).length))).length),new _dafny.CodePoint('q'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(131), new BigNumber(105))).length),new _dafny.CodePoint('f'.codePointAt(0))), function () {
          let _coll20 = new _dafny.Map();
          for (const _compr_20 of _dafny.IntegerRange(new BigNumber(754), new BigNumber(47))) {
            let _75_v2 = _compr_20;
            if (((new BigNumber(754)).isLessThanOrEqualTo(_75_v2)) && ((_75_v2).isLessThan(new BigNumber(47)))) {
              _coll20.push([(_75_v2).minus(new BigNumber((_dafny.Seq.of(false, false)).length)),new _dafny.CodePoint('d'.codePointAt(0))]);
            }
          }
          return _coll20;
        }(), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-540),new _dafny.CodePoint('o'.codePointAt(0))))).Elements) {
          let _76_v0 = _compr_18;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("wodnahax")).length),new _dafny.CodePoint('s'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
            let _coll21 = new _dafny.Set();
            for (const _compr_21 of _dafny.IntegerRange(new BigNumber(805), new BigNumber(429))) {
              let _77_v1 = _compr_21;
              if (((new BigNumber(805)).isLessThanOrEqualTo(_77_v1)) && ((_77_v1).isLessThan(new BigNumber(429)))) {
                _coll21.add((_77_v1).plus(new BigNumber(-564)));
              }
            }
            return _coll21;
          }()).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-498),new BigNumber((_dafny.Seq.of(true)).length))).length))).length))).length),new _dafny.CodePoint('q'.codePointAt(0))), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(131), new BigNumber(105))).length),new _dafny.CodePoint('f'.codePointAt(0))), function () {
            let _coll22 = new _dafny.Map();
            for (const _compr_22 of _dafny.IntegerRange(new BigNumber(754), new BigNumber(47))) {
              let _75_v2 = _compr_22;
              if (((new BigNumber(754)).isLessThanOrEqualTo(_75_v2)) && ((_75_v2).isLessThan(new BigNumber(47)))) {
                _coll22.push([(_75_v2).minus(new BigNumber((_dafny.Seq.of(false, false)).length)),new _dafny.CodePoint('d'.codePointAt(0))]);
              }
            }
            return _coll22;
          }(), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-540),new _dafny.CodePoint('o'.codePointAt(0)))), _76_v0)) {
            _coll18.push([_76_v0,_dafny.Map.Empty.slice().updateUnsafe(false,!(false))]);
          }
        }
        return _coll18;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true, !(true)))).cardinality()),new _dafny.CodePoint('a'.codePointAt(0))),_dafny.Map.Empty.slice().updateUnsafe(false,true)));
    };
    static fm61(p0, p1, p2, globalState) {
      return _dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(455)), function (_78_i0) {
        return (_dafny.Map.Empty.slice().updateUnsafe(false,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,false));
      }));
    };
    static fm62(p0, globalState) {
      if (true) {
        return _module.D18.create_DC43();
      } else {
        return _module.D18.create_DC43();
      }
    };
    static m0(p0, p1, p2, p3, globalState) {
      let _79_v0;
      _79_v0 = _dafny.Seq.of(p0);
      let _80_v1;
      _80_v1 = _dafny.Seq.of(_79_v0);
      let _81_i0;
      _81_i0 = _dafny.ZERO;
      L0: {
        while (_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat((_80_v1)[_module.__default.safeIndex(p1, new BigNumber((_80_v1).length))], _dafny.Seq.of(p0, p0)), _79_v0)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_81_i0)) {
              break L0;
            }
            _81_i0 = (_81_i0).plus(_dafny.ONE);
            let _82_v2;
            _82_v2 = new _dafny.CodePoint('f'.codePointAt(0));
            let _83_v3;
            _83_v3 = _module.D3.create_DC5(_82_v2);
            let _84_v4;
            _84_v4 = _dafny.Map.Empty.slice().updateUnsafe(p1,_83_v3);
            _84_v4 = (_84_v4).update(p3, _83_v3);
            let _85_v5;
            let _nw0 = Array((new BigNumber(13)).toNumber()).fill(false);
            _85_v5 = _nw0;
            let _86_v6;
            let _init0 = function (_87_i1) {
              return _dafny.Seq.Create(_module.__default.abs(new BigNumber(261)), function (_88_i2) {
                return _dafny.Seq.UnicodeFromString("rphheln");
              });
            };
            let _nw1 = Array((new BigNumber(23)).toNumber());
            for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw1.length); _i0_0++) {
              _nw1[_i0_0] = _init0(new BigNumber(_i0_0));
            }
            _86_v6 = _nw1;
            let _89_v7;
            _89_v7 = _dafny.Seq.UnicodeFromString("hne");
            let _index0 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_86_v6).length));
            (_86_v6)[_index0] = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("lgtdubw"), _89_v7, _89_v7);
            let _index1 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_85_v5).length));
            (_85_v5)[_index1] = (_79_v0)[_module.__default.safeIndex(p1, new BigNumber((_79_v0).length))];
            let _90_v8;
            _90_v8 = _dafny.Seq.of(_89_v7);
            let _91_v9;
            _91_v9 = _dafny.MultiSet.fromElements(!(p0), p0);
            let _index2 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_86_v6).length));
            let _index3 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_85_v5).length));
            let _rhs0 = _85_v5;
            let _rhs1 = p1;
            let _rhs2 = (p1).multipliedBy((p1).plus(new BigNumber((_79_v0).length)));
            let _rhs3 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_89_v7), _90_v8), ((p0) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(674)), ((_92_v7) => function (_93_i3) {
              return _92_v7;
            })(_89_v7))) : (_dafny.Seq.update(_90_v8, _module.__default.safeIndex(p1, new BigNumber((_90_v8).length)), _89_v7)))), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_89_v7), _90_v8), ((p0) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(674)), ((_94_v7) => function (_95_i3) {
              return _94_v7;
            })(_89_v7))) : (_dafny.Seq.update(_90_v8, _module.__default.safeIndex(p1, new BigNumber((_90_v8).length)), _89_v7))))).length)), ((p0) ? (_89_v7) : (_89_v7)));
            let _rhs4 = ((_dafny.MultiSet.fromElements(p0)).Difference(_91_v9)).IsSubsetOf(_91_v9);
            let _lhs0 = globalState;
            let _lhs1 = globalState;
            let _lhs2 = _86_v6;
            let _lhs3 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_86_v6).length));
            let _lhs4 = _85_v5;
            let _lhs5 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_85_v5).length));
            _85_v5 = _rhs0;
            _lhs0.f7 = _rhs1;
            _lhs1.f7 = _rhs2;
            _lhs2[_lhs3] = _rhs3;
            _lhs4[_lhs5] = _rhs4;
            (globalState).f13 = ((_85_v5)[_module.__default.safeIndex(new BigNumber(760), new BigNumber((_85_v5).length))]) === ((_79_v0)[_module.__default.safeIndex(p3, new BigNumber((_79_v0).length))]);
            let _96_v10;
            _96_v10 = _dafny.Seq.of(p1);
            (globalState).f16 = (_96_v10)[_module.__default.safeIndex(new BigNumber(947), new BigNumber((_96_v10).length))];
          }
        }
      }
      (globalState).f16 = p1;
      let _97_v11;
      let _nw2 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Map.Empty);
      _97_v11 = _nw2;
      let _98_v13;
      _98_v13 = _dafny.Seq.of(_module.__default.fm1(new BigNumber(712), globalState));
      let _index4 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_97_v11).length));
      (_97_v11)[_index4] = function () {
        let _coll23 = new _dafny.Map();
        for (const _compr_23 of (_98_v13).Elements) {
          let _99_v12 = _compr_23;
          if (_dafny.Seq.contains(_98_v13, _99_v12)) {
            _coll23.push([_99_v12,p1]);
          }
        }
        return _coll23;
      }();
      let _index5 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_97_v11).length));
      (_97_v11)[_index5] = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(209)), function (_100_i4) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      }),p1);
      let _pat_let_tv0 = p1;
      let _pat_let_tv1 = p0;
      let _pat_let_tv2 = p0;
      let _source3 = function (_source4) {
        if (_source4.is_DC0) {
          let _101___mcc_h2 = (_source4).cf0;
          let _102_cf0 = _101___mcc_h2;
          return _module.D1.create_DC2(_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv0,_pat_let_tv1));
        } else {
          let _103___mcc_h3 = (_source4).cf1;
          let _104_cf1 = _103___mcc_h3;
          return _module.D1.create_DC2(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(845),_pat_let_tv2));
        }
      }(_module.D0.create_DC0(p0));
      if (_source3.is_DC3) {
        let _105___mcc_h0 = (_source3).cf3;
        let _106_cf3 = _105___mcc_h0;
        let _107_v14;
        let _nw3 = Array((new BigNumber(18)).toNumber()).fill(false);
        _107_v14 = _nw3;
        let _index6 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_107_v14).length));
        (_107_v14)[_index6] = _106_cf3;
        let _index7 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_107_v14).length));
        (_107_v14)[_index7] = _106_cf3;
        if ((_107_v14)[_module.__default.safeIndex(new BigNumber(612), new BigNumber((_107_v14).length))]) {
          let _108_v15;
          let _nw4 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _108_v15 = _nw4;
          let _109_v16;
          _109_v16 = new _dafny.CodePoint('c'.codePointAt(0));
          let _110_v17;
          _110_v17 = _dafny.Seq.UnicodeFromString("tmcphea");
          let _index8 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_108_v15).length));
          (_108_v15)[_index8] = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(320)), function (_111_i5) {
            return new _dafny.CodePoint('q'.codePointAt(0));
          }), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(320)), function (_112_i5) {
            return new _dafny.CodePoint('q'.codePointAt(0));
          })).length)), _109_v16), _110_v17);
          let _113_v18;
          _113_v18 = _dafny.Set.fromElements(p0);
          let _114_v19;
          _114_v19 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_110_v17).length),new BigNumber((_113_v18).length));
          let _115_v20;
          _115_v20 = _dafny.Map.Empty.slice().updateUnsafe(true,false);
          let _116_v21;
          _116_v21 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(p3,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(_115_v20, globalState),p3)).length)), _dafny.Map.Empty.slice().updateUnsafe(p1,p3));
          let _index9 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_108_v15).length));
          (_108_v15)[_index9] = (((_dafny.Set.fromElements(_114_v19)).equals(_116_v21)) ? (_dafny.Seq.update(_dafny.Seq.UnicodeFromString("on"), _module.__default.safeIndex(p3, new BigNumber((_dafny.Seq.UnicodeFromString("on")).length)), _109_v16)) : (_110_v17));
          _115_v20 = (_115_v20).update(_106_cf3, _106_cf3);
          (globalState).f20 = (((_module.__default.fm0(p1, p3, new BigNumber((_115_v20).length), (_107_v14)[_module.__default.safeIndex(new BigNumber(612), new BigNumber((_107_v14).length))], globalState)).isLessThan(p1)) ? (_dafny.Seq.Concat(_module.__default.fm1(new BigNumber((_dafny.Seq.UnicodeFromString("yootrwdif")).length), globalState), (_108_v15)[_module.__default.safeIndex(new BigNumber(389), new BigNumber((_108_v15).length))])) : (_dafny.Seq.update(_dafny.Seq.Concat(_110_v17, _dafny.Seq.UnicodeFromString("r")), _module.__default.safeIndex(p1, new BigNumber((_dafny.Seq.Concat(_110_v17, _dafny.Seq.UnicodeFromString("r"))).length)), _109_v16)));
          (globalState).f16 = (_dafny.ZERO).minus(p3);
          (globalState).f2 = _106_cf3;
        } else {
          _106_cf3 = ((p0) ? (_106_cf3) : (false));
          (globalState).f0 = p0;
          (globalState).f2 = !(false);
          let _117_v22;
          _117_v22 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(418)).multipliedBy((_dafny.ZERO).minus(p3)),p3);
          let _118_v23;
          let _init1 = ((_119_p0) => function (_120_i6) {
            return _module.D0.create_DC0(_119_p0);
          })(p0);
          let _nw5 = Array((new BigNumber(21)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw5.length); _i0_1++) {
            _nw5[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _118_v23 = _nw5;
          let _121_v24;
          _121_v24 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(p0)).length),!(p0));
          let _122_v25;
          _122_v25 = _module.D0.create_DC0((((_121_v24).contains(new BigNumber(147))) ? ((_121_v24).get(new BigNumber(147))) : (p0)));
          let _pat_let_tv3 = _106_cf3;
          let _index10 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_118_v23).length));
          (_118_v23)[_index10] = function (_pat_let0_0) {
            return function (_123_dt__update__tmp_h0) {
              return function (_pat_let1_0) {
                return function (_124_dt__update_hcf0_h0) {
                  return _module.D0.create_DC0(_124_dt__update_hcf0_h0);
                }(_pat_let1_0);
              }(_pat_let_tv3);
            }(_pat_let0_0);
          }(_122_v25);
          let _125_v27;
          let _init2 = ((_126_v24) => function (_127_i7) {
            return _module.__default.safeModuloInt(_127_i7, new BigNumber((function () {
              let _coll24 = new _dafny.Set();
              for (const _compr_24 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),_126_v24)).Keys.Elements) {
                let _128_v26 = _compr_24;
                if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),_126_v24)).contains(_128_v26)) {
                  _coll24.add(_128_v26);
                }
              }
              return _coll24;
            }()).length));
          })(_121_v24);
          let _nw6 = Array((new BigNumber(3)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw6.length); _i0_2++) {
            _nw6[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _125_v27 = _nw6;
          let _index11 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_125_v27).length));
          (_125_v27)[_index11] = (_dafny.ZERO).minus(p1);
          let _129_v28;
          _129_v28 = _dafny.Seq.UnicodeFromString("ylg");
          let _index12 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_118_v23).length));
          let _index13 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_107_v14).length));
          let _index14 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_125_v27).length));
          let _index15 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_107_v14).length));
          let _rhs5 = (_117_v22).update((p3).minus(new BigNumber(-542)), p3);
          let _rhs6 = _122_v25;
          let _rhs7 = p0;
          let _rhs8 = (p1).minus(p3);
          let _rhs9 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-830)), function (_130_i8) {
            return new _dafny.CodePoint('r'.codePointAt(0));
          }), _129_v28);
          let _lhs6 = _118_v23;
          let _lhs7 = _module.__default.safeIndex(new BigNumber(286), new BigNumber((_118_v23).length));
          let _lhs8 = _107_v14;
          let _lhs9 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_107_v14).length));
          let _lhs10 = _125_v27;
          let _lhs11 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_125_v27).length));
          let _lhs12 = _107_v14;
          let _lhs13 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_107_v14).length));
          _117_v22 = _rhs5;
          _lhs6[_lhs7] = _rhs6;
          _lhs8[_lhs9] = _rhs7;
          _lhs10[_lhs11] = _rhs8;
          _lhs12[_lhs13] = _rhs9;
          let _index16 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_107_v14).length));
          let _index17 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_125_v27).length));
          let _rhs10 = _106_cf3;
          let _rhs11 = (_125_v27)[_module.__default.safeIndex(new BigNumber(224), new BigNumber((_125_v27).length))];
          let _rhs12 = _module.__default.safeModuloInt((_dafny.ZERO).minus(p3), p1);
          let _rhs13 = _125_v27;
          let _lhs14 = _107_v14;
          let _lhs15 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_107_v14).length));
          let _lhs16 = _125_v27;
          let _lhs17 = _module.__default.safeIndex(new BigNumber(224), new BigNumber((_125_v27).length));
          let _lhs18 = globalState;
          _lhs14[_lhs15] = _rhs10;
          _lhs16[_lhs17] = _rhs11;
          _lhs18.f7 = _rhs12;
          _125_v27 = _rhs13;
        }
        let _131_v29;
        let _nw7 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _131_v29 = _nw7;
        let _132_v30;
        _132_v30 = _dafny.Seq.UnicodeFromString("x");
        let _index18 = _module.__default.safeIndex(new BigNumber(916), new BigNumber((_131_v29).length));
        (_131_v29)[_index18] = _dafny.Seq.Concat(_module.__default.fm1(p1, globalState), _132_v30);
        let _index19 = _module.__default.safeIndex(new BigNumber(916), new BigNumber((_131_v29).length));
        (_131_v29)[_index19] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(90)), function (_133_i9) {
          return new _dafny.CodePoint('w'.codePointAt(0));
        }), _132_v30);
        let _index20 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_107_v14).length));
        (_107_v14)[_index20] = true;
      } else {
        let _134___mcc_h1 = (_source3).cf2;
        let _135_cf2 = _134___mcc_h1;
        (globalState).f16 = p3;
        let _136_v31;
        _136_v31 = new _dafny.CodePoint('q'.codePointAt(0));
        let _137_v32;
        _137_v32 = _dafny.Map.Empty.slice().updateUnsafe((p3).plus(p1),_136_v31);
        _136_v31 = (((_137_v32).contains(p3)) ? ((_137_v32).get(p3)) : (_136_v31));
        let _138_v33;
        _138_v33 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        let _139_v34;
        let _nw8 = new _module.C3();
        _nw8.__ctor();
        _139_v34 = _nw8;
        let _140_v35;
        let _nw9 = Array((new BigNumber(18)).toNumber());
        _nw9[(_dafny.ZERO).toNumber()] = _139_v34;
        _nw9[(_dafny.ONE).toNumber()] = _139_v34;
        _nw9[(new BigNumber(2)).toNumber()] = _139_v34;
        _nw9[(new BigNumber(3)).toNumber()] = _139_v34;
        _nw9[(new BigNumber(4)).toNumber()] = _139_v34;
        _nw9[(new BigNumber(5)).toNumber()] = _139_v34;
        _nw9[(new BigNumber(6)).toNumber()] = _139_v34;
        _nw9[(new BigNumber(7)).toNumber()] = _139_v34;
        _nw9[(new BigNumber(8)).toNumber()] = _139_v34;
        _nw9[(new BigNumber(9)).toNumber()] = _139_v34;
        _nw9[(new BigNumber(10)).toNumber()] = _139_v34;
        _nw9[(new BigNumber(11)).toNumber()] = _139_v34;
        _nw9[(new BigNumber(12)).toNumber()] = _139_v34;
        _nw9[(new BigNumber(13)).toNumber()] = _139_v34;
        _nw9[(new BigNumber(14)).toNumber()] = _139_v34;
        _nw9[(new BigNumber(15)).toNumber()] = _139_v34;
        _nw9[(new BigNumber(16)).toNumber()] = _139_v34;
        _nw9[(new BigNumber(17)).toNumber()] = _139_v34;
        _140_v35 = _nw9;
        let _141_v36;
        _141_v36 = _dafny.Map.Empty.slice().updateUnsafe(_140_v35,p0);
        let _142_v37;
        _142_v37 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(_138_v33, globalState),_module.D22.create_DC56(_141_v36));
        let _143_v38;
        _143_v38 = _dafny.Seq.of(_142_v37);
        let _144_v39;
        _144_v39 = _dafny.Seq.UnicodeFromString("frletcac");
        let _145_v40;
        let _nw10 = new _module.C4();
        _nw10.__ctor(!(!((_143_v38)[_module.__default.safeIndex(new BigNumber((_144_v39).length), new BigNumber((_143_v38).length))]).equals(_142_v37)), p0);
        _145_v40 = _nw10;
        _135_cf2 = (_135_cf2).update(((_145_v40.f25) ? (p3) : (_module.__default.fm0(new BigNumber((_79_v0).length), new BigNumber(936), p3, _145_v40.f25, globalState))), false);
      }
      let _146_v41;
      let _nw11 = new _module.C2();
      _nw11.__ctor(p1);
      _146_v41 = _nw11;
      let _147_v42;
      _147_v42 = _module.D18.create_DC43();
      let _148_v43;
      _148_v43 = new _dafny.CodePoint('l'.codePointAt(0));
      let _149_v44;
      _149_v44 = _dafny.Map.Empty.slice().updateUnsafe(p0,_148_v43);
      let _150_v45;
      _150_v45 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_149_v44).length));
      _147_v42 = _module.__default.fm62(new BigNumber(((_150_v45).Merge(_150_v45)).length), globalState);
      return;
    }
    static Main(__noArgsParameter) {
      let _151_v0;
      let _init3 = function (_152_i0) {
        return (_152_i0).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(423),new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(!(false), !(!(true))))).cardinality()))).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-866),new BigNumber(259))).length))).length));
      };
      let _nw12 = Array((new BigNumber(18)).toNumber());
      for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw12.length); _i0_3++) {
        _nw12[_i0_3] = _init3(new BigNumber(_i0_3));
      }
      _151_v0 = _nw12;
      let _153_v1;
      _153_v1 = _dafny.Seq.UnicodeFromString("ucjarmpr");
      let _154_v2;
      _154_v2 = false;
      let _155_v3;
      _155_v3 = _dafny.Set.fromElements(_153_v1, _153_v1);
      let _156_v5;
      _156_v5 = new BigNumber(710);
      let _157_v6;
      _157_v6 = new _dafny.CodePoint('t'.codePointAt(0));
      let _158_v7;
      let _nw13 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Map.Empty);
      _158_v7 = _nw13;
      let _159_globalState;
      let _nw14 = new _module.GlobalState();
      _nw14.__ctor(true, _151_v0, false, new BigNumber(746), true, false, false, new BigNumber(729), _153_v1, new BigNumber(-908), ((_154_v2) ? (_155_v3) : (function () {
        let _coll25 = new _dafny.Set();
        for (const _compr_25 of (_155_v3).Elements) {
          let _160_v4 = _compr_25;
          if ((_155_v3).contains(_160_v4)) {
            _coll25.add(_160_v4);
          }
        }
        return _coll25;
      }())), new BigNumber(-466), new BigNumber(794), true, _dafny.Seq.Concat(_dafny.Seq.update(_153_v1, _module.__default.safeIndex(_156_v5, new BigNumber((_153_v1).length)), _157_v6), _153_v1), new BigNumber(52), new BigNumber(251), false, new BigNumber(953), _153_v1, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("cesqe"), _153_v1), false, _158_v7);
      _159_globalState = _nw14;
      let _161_v8;
      _161_v8 = _dafny.Seq.of(_156_v5);
      _156_v5 = ((_dafny.ZERO).minus(new BigNumber(-504))).minus((_161_v8)[_module.__default.safeIndex((_161_v8)[_module.__default.safeIndex(_156_v5, new BigNumber((_161_v8).length))], new BigNumber((_161_v8).length))]);
      let _162_v9;
      let _nw15 = Array((new BigNumber(7)).toNumber()).fill(false);
      _162_v9 = _nw15;
      let _index21 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length));
      (_162_v9)[_index21] = _154_v2;
      let _index22 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length));
      (_162_v9)[_index22] = _154_v2;
      let _index23 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length));
      (_162_v9)[_index23] = (_154_v2) && ((_162_v9)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length))]);
      if (_154_v2) {
        (_159_globalState).f16 = _156_v5;
        let _163_v10;
        _163_v10 = _dafny.Map.Empty.slice().updateUnsafe(_153_v1,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(new BigNumber(727), new BigNumber((_dafny.Seq.UnicodeFromString("faepiih")).length), _156_v5, _154_v2, _159_globalState),_154_v2)).length));
        let _164_v11;
        _164_v11 = _dafny.Map.Empty.slice().updateUnsafe(_153_v1,new BigNumber((_163_v10).length));
        let _165_v12;
        _165_v12 = _dafny.Seq.of(_154_v2);
        let _166_v13;
        _166_v13 = _dafny.Map.Empty.slice().updateUnsafe(_165_v12,_module.__default.fm1(_156_v5, _159_globalState));
        let _167_v14;
        _167_v14 = _dafny.Seq.of(_165_v12, _165_v12, _165_v12);
        _module.__default.m0((_162_v9)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length))], new BigNumber((_164_v11).length), (_166_v13).update((_167_v14)[_module.__default.safeIndex(_156_v5, new BigNumber((_167_v14).length))], _153_v1), ((_154_v2) ? (_156_v5) : (_156_v5)), _159_globalState);
        (_159_globalState).f7 = (_161_v8)[_module.__default.safeIndex(_156_v5, new BigNumber((_161_v8).length))];
        let _168_v15;
        _168_v15 = _module.D0.create_DC0(_154_v2);
        let _169_v16;
        _169_v16 = _module.D0.create_DC1(_module.D0.create_DC1(_168_v15));
        let _170_v17;
        _170_v17 = _module.D0.create_DC1(_169_v16);
        let _source5 = _170_v17;
        if (_source5.is_DC0) {
          let _171___mcc_h0 = (_source5).cf0;
          let _172_cf0 = _171___mcc_h0;
          let _173_v18;
          _173_v18 = _dafny.Map.Empty.slice().updateUnsafe(_154_v2,_154_v2);
          let _174_v19;
          _174_v19 = _dafny.Map.Empty.slice().updateUnsafe(_172_cf0,(((_173_v18).contains(_154_v2)) ? ((_173_v18).get(_154_v2)) : (_154_v2)));
          let _175_v20;
          _175_v20 = _dafny.Map.Empty.slice().updateUnsafe(_156_v5,(((_174_v19).contains(_154_v2)) ? ((_174_v19).get(_154_v2)) : (_172_cf0)));
          let _176_v21;
          _176_v21 = _dafny.Set.fromElements(_156_v5, new BigNumber((_175_v20).length));
          _module.__default.m0((_176_v21).IsProperSubsetOf(_176_v21), _156_v5, _dafny.Map.Empty.slice().updateUnsafe(_165_v12,_153_v1), _156_v5, _159_globalState);
          let _177_v22;
          let _init4 = ((_178_v20) => function (_179_i1) {
            return (_module.D1.create_DC2(_178_v20)).dtor_cf2;
          })(_175_v20);
          let _nw16 = Array((new BigNumber(2)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw16.length); _i0_4++) {
            _nw16[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _177_v22 = _nw16;
          let _180_v23;
          _180_v23 = _dafny.Map.Empty.slice().updateUnsafe(_156_v5,_156_v5);
          let _181_v24;
          _181_v24 = _dafny.Seq.of(_180_v23);
          let _182_v25;
          _182_v25 = _dafny.Seq.of(_dafny.Seq.Concat(_181_v24, _181_v24));
          let _183_v26;
          _183_v26 = _dafny.MultiSet.fromElements(_151_v0);
          let _184_v27;
          _184_v27 = _dafny.Map.Empty.slice().updateUnsafe(_172_cf0,_177_v22);
          let _rhs14 = new BigNumber((_dafny.MultiSet.FromArray((_182_v25)[_module.__default.safeIndex(_156_v5, new BigNumber((_182_v25).length))])).cardinality());
          let _rhs15 = _module.__default.fm2(_dafny.Map.Empty.slice().updateUnsafe(_154_v2,(_162_v9)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length))]), _159_globalState);
          let _rhs16 = (((_165_v12)[_module.__default.safeIndex((((_183_v26).contains(_151_v0)) ? ((_183_v26).get(_151_v0)) : (_156_v5)), new BigNumber((_165_v12).length))]) ? (_177_v22) : ((((_184_v27).contains(_154_v2)) ? ((_184_v27).get(_154_v2)) : (_177_v22))));
          let _rhs17 = _module.__default.safeDivisionInt((_156_v5).plus(_156_v5), _156_v5);
          let _rhs18 = (_162_v9)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length))];
          let _lhs19 = _159_globalState;
          let _lhs20 = _159_globalState;
          let _lhs21 = _159_globalState;
          let _lhs22 = _159_globalState;
          _lhs19.f7 = _rhs14;
          _lhs20.f2 = _rhs15;
          _177_v22 = _rhs16;
          _lhs21.f16 = _rhs17;
          _lhs22.f2 = _rhs18;
          let _185_v28;
          _185_v28 = _dafny.MultiSet.fromElements(_173_v18, _dafny.Map.Empty.slice().updateUnsafe((_162_v9)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length))],_154_v2));
          let _186_v29;
          _186_v29 = _dafny.MultiSet.fromElements(_156_v5, _156_v5);
          (_159_globalState).f16 = _module.__default.fm0(_156_v5, new BigNumber((_185_v28).cardinality()), (_dafny.ZERO).minus(_module.__default.safeModuloInt(_156_v5, _156_v5)), (_186_v29).IsProperSubsetOf(_186_v29), _159_globalState);
          (_159_globalState).f13 = _dafny.Seq.IsPrefixOf(_153_v1, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("urqdgor"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-451)), ((_187_v6) => function (_188_i2) {
            return _187_v6;
          })(_157_v6))));
        } else {
          let _189___mcc_h1 = (_source5).cf1;
          let _190_cf1 = _189___mcc_h1;
          (_159_globalState).f21 = _154_v2;
          _156_v5 = new BigNumber(439);
          let _191_v30;
          _191_v30 = _dafny.Map.Empty.slice().updateUnsafe((_162_v9)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length))],(_162_v9)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length))]);
          _module.__default.m0(_module.__default.fm2(_191_v30, _159_globalState), _156_v5, _166_v13, new BigNumber(-198), _159_globalState);
          let _192_v31;
          _192_v31 = _module.D1.create_DC3((_154_v2) || ((_162_v9)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length))]));
          _192_v31 = _192_v31;
        }
        let _193_v32;
        _193_v32 = _module.D0.create_DC0((_162_v9)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length))]);
        _193_v32 = _193_v32;
      } else {
        (_159_globalState).f21 = true;
        _153_v1 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("bsi"), _153_v1);
        _154_v2 = false;
        let _index24 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_151_v0).length));
        (_151_v0)[_index24] = (_dafny.ZERO).minus(new BigNumber(-193));
        let _194_v33;
        _194_v33 = _dafny.Set.fromElements(_154_v2, _154_v2);
        let _195_v34;
        _195_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_194_v33).length),(_162_v9)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length))]);
        let _196_v36;
        _196_v36 = _module.D1.create_DC2((_195_v34).Merge(function () {
  let _coll26 = new _dafny.Map();
  for (const _compr_26 of _dafny.IntegerRange(new BigNumber(929), new BigNumber(635))) {
    let _197_v35 = _compr_26;
    if (((new BigNumber(929)).isLessThanOrEqualTo(_197_v35)) && ((_197_v35).isLessThan(new BigNumber(635)))) {
      _coll26.push([_module.__default.safeDivisionInt(_197_v35, _156_v5),(_162_v9)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length))]]);
    }
  }
  return _coll26;
}()));
        let _index25 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_151_v0).length));
        let _rhs19 = (_162_v9)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length))];
        let _rhs20 = ((((_154_v2) ? (_154_v2) : ((_162_v9)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length))]))) ? (new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(_156_v5))).cardinality())) : (_156_v5));
        let _rhs21 = _196_v36;
        let _lhs23 = _159_globalState;
        let _lhs24 = _151_v0;
        let _lhs25 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_151_v0).length));
        _lhs23.f21 = _rhs19;
        _lhs24[_lhs25] = _rhs20;
        _196_v36 = _rhs21;
        let _198_v37;
        let _init5 = ((_199_v1) => function (_200_i3) {
          return _199_v1;
        })(_153_v1);
        let _nw17 = Array((new BigNumber(11)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw17.length); _i0_5++) {
          _nw17[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _198_v37 = _nw17;
        let _201_v38;
        _201_v38 = _dafny.Map.Empty.slice().updateUnsafe(_198_v37,_dafny.Seq.Concat(_153_v1, _153_v1));
        _201_v38 = (_201_v38).update(_198_v37, _153_v1);
      }
      let _202_v39;
      _202_v39 = _module.D0.create_DC0(_154_v2);
      let _203_v40;
      _203_v40 = _dafny.Seq.of((_162_v9)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length))], false, _154_v2, false);
      let _204_v41;
      _204_v41 = _dafny.Map.Empty.slice().updateUnsafe(_203_v40,_dafny.Seq.UnicodeFromString("ncojfqse"));
      _module.__default.m0((_202_v39).dtor_cf0, _156_v5, _204_v41, _156_v5, _159_globalState);
      let _ingredients0 = [];
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_151_v0).length))) {
        let _205_i4 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_205_i4)) && ((_205_i4).isLessThan(new BigNumber((_151_v0).length))))) {
          _ingredients0.push(_dafny.Tuple.of(_151_v0, (_205_i4).toNumber(), (_205_i4).plus((((_162_v9)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length))]) ? (_156_v5) : ((_dafny.ZERO).minus(_156_v5))))));
        }
      }
      for (const _tup0 of _ingredients0) {
        _tup0[0][_tup0[1]] = _tup0[2];
      }
      _module.__default.m0(((false) ? (_154_v2) : (true)), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_156_v5), _161_v8)).length), (_204_v41).Merge(_204_v41), _156_v5, _159_globalState);
      let _206_v42;
      _206_v42 = _dafny.Map.Empty.slice().updateUnsafe(_154_v2,_156_v5);
      let _207_v43;
      _207_v43 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_153_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-461)), ((_208_v6) => function (_209_i5) {
        return _208_v6;
      })(_157_v6))),!((_206_v42).contains((_162_v9)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length))])));
      _207_v43 = _207_v43;
      (_159_globalState).f13 = _154_v2;
      let _210_v44;
      _210_v44 = _dafny.Map.Empty.slice().updateUnsafe(_156_v5,(_162_v9)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length))]);
      let _211_v45;
      _211_v45 = _module.D1.create_DC2(_210_v44);
      let _212_v46;
      _212_v46 = _dafny.Set.fromElements(_156_v5);
      _211_v45 = _module.D1.create_DC2(_module.__default.fm3(_156_v5, false, _153_v1, new BigNumber((_212_v46).length), _159_globalState));
      (_159_globalState).f16 = _156_v5;
      let _213_v47;
      _213_v47 = _dafny.MultiSet.fromElements(_156_v5);
      (_159_globalState).f13 = (_213_v47).IsDisjointFrom(_213_v47);
      let _source6 = _module.D0.create_DC1(_module.__default.fm4(_module.__default.fm0(_156_v5, new BigNumber(593), _156_v5, _154_v2, _159_globalState), _156_v5, _159_globalState));
      if (_source6.is_DC0) {
        let _214___mcc_h2 = (_source6).cf0;
        let _215_cf0 = _214___mcc_h2;
        let _pat_let_tv4 = _210_v44;
        let _source7 = function (_pat_let2_0) {
          return function (_216_dt__update__tmp_h0) {
            return function (_pat_let3_0) {
              return function (_217_dt__update_hcf2_h0) {
                return _module.D1.create_DC2(_217_dt__update_hcf2_h0);
              }(_pat_let3_0);
            }(_pat_let_tv4);
          }(_pat_let2_0);
        }(_211_v45);
        if (_source7.is_DC3) {
          let _218___mcc_h4 = (_source7).cf3;
          let _219_cf3 = _218___mcc_h4;
          _151_v0 = _151_v0;
          _157_v6 = ((_154_v2) ? (_157_v6) : (_157_v6));
          let _nw18 = Array((new BigNumber(4)).toNumber()).fill(false);
          _162_v9 = _nw18;
          let _index26 = _module.__default.safeIndex(new BigNumber(575), new BigNumber((_151_v0).length));
          (_151_v0)[_index26] = _156_v5;
          let _220_v48;
          _220_v48 = _dafny.MultiSet.fromElements((_162_v9)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length))]);
          let _index27 = _module.__default.safeIndex(new BigNumber(575), new BigNumber((_151_v0).length));
          (_151_v0)[_index27] = _module.__default.safeDivisionInt(_156_v5, _module.__default.safeModuloInt(_module.__default.fm0(_156_v5, new BigNumber((_220_v48).cardinality()), _156_v5, (((_210_v44).contains(new BigNumber((_153_v1).length))) ? ((_210_v44).get(new BigNumber((_153_v1).length))) : ((_162_v9)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length))])), _159_globalState), (_dafny.ZERO).minus(_156_v5)));
        } else {
          let _221___mcc_h5 = (_source7).cf2;
          let _222_cf2 = _221___mcc_h5;
          _203_v40 = _dafny.Seq.Concat(_203_v40, _dafny.Seq.of((_162_v9)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length))]));
          _161_v8 = _dafny.Seq.Concat(_dafny.Seq.of(_156_v5, _156_v5), _161_v8);
          (_159_globalState).f13 = !(!(_212_v46).equals(function () {
            let _coll27 = new _dafny.Set();
            for (const _compr_27 of _dafny.IntegerRange(new BigNumber(375), new BigNumber(820))) {
              let _223_v49 = _compr_27;
              if (((new BigNumber(375)).isLessThanOrEqualTo(_223_v49)) && ((_223_v49).isLessThan(new BigNumber(820)))) {
                _coll27.add((_223_v49).minus(new BigNumber((_203_v40).length)));
              }
            }
            return _coll27;
          }()));
          let _224_v50;
          _224_v50 = _dafny.MultiSet.fromElements(_151_v0);
          let _225_v51;
          _225_v51 = _dafny.Map.Empty.slice().updateUnsafe(_215_cf0,(_224_v50));
          let _226_v52;
          _226_v52 = _dafny.MultiSet.fromElements(_151_v0);
          _225_v51 = (_225_v51).update(_154_v2, _226_v52);
        }
        (_159_globalState).f13 = _154_v2;
        _151_v0 = _151_v0;
        _157_v6 = new _dafny.CodePoint('e'.codePointAt(0));
      } else {
        let _227___mcc_h3 = (_source6).cf1;
        let _228_cf1 = _227___mcc_h3;
        let _229_v53;
        let _nw19 = Array((new BigNumber(19)).toNumber()).fill([]);
        _229_v53 = _nw19;
        let _index28 = _module.__default.safeIndex(new BigNumber(229), new BigNumber((_229_v53).length));
        (_229_v53)[_index28] = _151_v0;
        let _index29 = _module.__default.safeIndex(new BigNumber(229), new BigNumber((_229_v53).length));
        (_229_v53)[_index29] = _151_v0;
        _module.__default.m0((_162_v9)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length))], new BigNumber(-852), _204_v41, _module.__default.safeModuloInt(_156_v5, _156_v5), _159_globalState);
        (_159_globalState).f0 = (_162_v9)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length))];
        let _230_v54;
        _230_v54 = _dafny.Map.Empty.slice().updateUnsafe(_157_v6,_154_v2);
        _203_v40 = _dafny.Seq.of((_162_v9)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length))], (((_230_v54).contains(_157_v6)) ? ((_230_v54).get(_157_v6)) : (false)));
      }
      let _231_i6;
      _231_i6 = _dafny.ZERO;
      L1: {
        while ((_162_v9)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length))]) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_231_i6)) {
              break L1;
            }
            _231_i6 = (_231_i6).plus(_dafny.ONE);
            let _pat_let_tv5 = _210_v44;
            _211_v45 = function (_pat_let4_0) {
              return function (_232_dt__update__tmp_h1) {
                return function (_pat_let5_0) {
                  return function (_233_dt__update_hcf2_h1) {
                    return _module.D1.create_DC2(_233_dt__update_hcf2_h1);
                  }(_pat_let5_0);
                }(_pat_let_tv5);
              }(_pat_let4_0);
            }(_211_v45);
            _211_v45 = _211_v45;
            let _234_v55;
            _234_v55 = _dafny.Set.fromElements(_154_v2, (_162_v9)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length))]);
            _156_v5 = new BigNumber((_234_v55).length);
            (_159_globalState).f16 = ((_dafny.ZERO).minus(_156_v5)).minus((new BigNumber(148)).plus((_dafny.ZERO).minus(_156_v5)));
          }
        }
      }
      let _235_v56;
      _235_v56 = _dafny.Map.Empty.slice().updateUnsafe(_162_v9,_156_v5);
      let _236_v57;
      _236_v57 = _module.D3.create_DC6(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("yswqe"), _module.__default.safeIndex(_156_v5, new BigNumber((_dafny.Seq.UnicodeFromString("yswqe")).length)), _157_v6), _157_v6, _154_v2);
      let _237_v58;
      let _nw20 = Array((new BigNumber(21)).toNumber());
      _nw20[(_dafny.ZERO).toNumber()] = _235_v56;
      _nw20[(_dafny.ONE).toNumber()] = _235_v56;
      _nw20[(new BigNumber(2)).toNumber()] = (_235_v56).update(_162_v9, _156_v5);
      _nw20[(new BigNumber(3)).toNumber()] = (_235_v56).Merge(_235_v56);
      _nw20[(new BigNumber(4)).toNumber()] = _235_v56;
      _nw20[(new BigNumber(5)).toNumber()] = _235_v56;
      _nw20[(new BigNumber(6)).toNumber()] = (((_162_v9)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length))]) ? (_235_v56) : (_235_v56));
      _nw20[(new BigNumber(7)).toNumber()] = (_235_v56).Merge((_235_v56).update(_162_v9, new BigNumber((_dafny.MultiSet.fromElements((_236_v57).dtor_cf7)).cardinality())));
      _nw20[(new BigNumber(8)).toNumber()] = (_235_v56).Merge(_235_v56);
      _nw20[(new BigNumber(9)).toNumber()] = _235_v56;
      _nw20[(new BigNumber(10)).toNumber()] = _235_v56;
      _nw20[(new BigNumber(11)).toNumber()] = _235_v56;
      _nw20[(new BigNumber(12)).toNumber()] = _235_v56;
      _nw20[(new BigNumber(13)).toNumber()] = _235_v56;
      _nw20[(new BigNumber(14)).toNumber()] = (_235_v56).update(_162_v9, _156_v5);
      _nw20[(new BigNumber(15)).toNumber()] = _235_v56;
      _nw20[(new BigNumber(16)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_162_v9,_156_v5);
      _nw20[(new BigNumber(17)).toNumber()] = (_235_v56).Merge(_235_v56);
      _nw20[(new BigNumber(18)).toNumber()] = _235_v56;
      _nw20[(new BigNumber(19)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_162_v9,_156_v5);
      _nw20[(new BigNumber(20)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_162_v9,new BigNumber(129));
      _237_v58 = _nw20;
      let _index30 = _module.__default.safeIndex(new BigNumber(421), new BigNumber((_237_v58).length));
      (_237_v58)[_index30] = _235_v56;
      let _index31 = _module.__default.safeIndex(new BigNumber(421), new BigNumber((_237_v58).length));
      let _index32 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length));
      let _rhs22 = _235_v56;
      let _rhs23 = !(_154_v2);
      let _rhs24 = _156_v5;
      let _lhs26 = _237_v58;
      let _lhs27 = _module.__default.safeIndex(new BigNumber(421), new BigNumber((_237_v58).length));
      let _lhs28 = _162_v9;
      let _lhs29 = _module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length));
      _lhs26[_lhs27] = _rhs22;
      _lhs28[_lhs29] = _rhs23;
      _156_v5 = _rhs24;
      _154_v2 = ((_162_v9)[_module.__default.safeIndex(new BigNumber(773), new BigNumber((_162_v9).length))]) && (_154_v2);
      process.stdout.write(_dafny.toString((_151_v0)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_v0)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_v0)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_v0)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_v0)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_v0)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_v0)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_v0)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_v0)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_v0)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_v0)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_v0)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_v0)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_v0)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_v0)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_v0)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_v0)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_v0)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_153_v1).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_154_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_155_v3).equals(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("ucjarmpr")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_156_v5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_157_v6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_159_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_159_globalState).f1)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_159_globalState).f1)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_159_globalState).f1)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_159_globalState).f1)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_159_globalState).f1)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_159_globalState).f1)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_159_globalState).f1)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_159_globalState).f1)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_159_globalState).f1)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_159_globalState).f1)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_159_globalState).f1)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_159_globalState).f1)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_159_globalState).f1)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_159_globalState).f1)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_159_globalState).f1)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_159_globalState).f1)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_159_globalState).f1)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_159_globalState).f1)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_159_globalState.f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_159_globalState.f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_159_globalState.f8).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_globalState.f10).equals(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("ucjarmpr")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_159_globalState.f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_159_globalState).f14).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_globalState).f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_159_globalState.f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_globalState).f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_159_globalState).f18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_159_globalState.f19).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_159_globalState.f20).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_159_globalState.f21));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_161_v8, _dafny.Seq.of(new BigNumber(710)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_162_v9)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_202_v39).dtor_cf0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_203_v40, _dafny.Seq.of(false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_204_v41).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false, false, false, false),_dafny.Seq.UnicodeFromString("ncojfqse")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_206_v42).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-206)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_207_v43).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("bsiucjarmprttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt"),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_210_v44).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-206),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_211_v45).dtor_cf2).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(3),false).updateUnsafe(new BigNumber(581),false).updateUnsafe(_dafny.ZERO,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_212_v46).equals(_dafny.Set.fromElements(new BigNumber(-206)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_213_v47).equals(_dafny.MultiSet.fromElements(new BigNumber(-206)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_231_i6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_235_v56).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_236_v57).dtor_cf6).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_236_v57).dtor_cf7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_236_v57).dtor_cf8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_237_v58)[_dafny.ZERO]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_237_v58)[_dafny.ONE]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_237_v58)[new BigNumber(2)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_237_v58)[new BigNumber(3)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_237_v58)[new BigNumber(4)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_237_v58)[new BigNumber(5)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_237_v58)[new BigNumber(6)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_237_v58)[new BigNumber(7)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_237_v58)[new BigNumber(8)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_237_v58)[new BigNumber(9)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_237_v58)[new BigNumber(10)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_237_v58)[new BigNumber(11)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_237_v58)[new BigNumber(12)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_237_v58)[new BigNumber(13)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_237_v58)[new BigNumber(14)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_237_v58)[new BigNumber(15)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_237_v58)[new BigNumber(16)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_237_v58)[new BigNumber(17)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_237_v58)[new BigNumber(18)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_237_v58)[new BigNumber(19)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_237_v58)[new BigNumber(20)]).length)));
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
    static create_DC1(cf1) {
      let $dt = new D0(1);
      $dt.cf1 = cf1;
      return $dt;
    }
    get is_DC0() { return this.$tag === 0; }
    get is_DC1() { return this.$tag === 1; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf1() { return this.cf1; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf0 === other.cf0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf1, other.cf1);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC0(false);
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
    static create_DC3(cf3) {
      let $dt = new D1(0);
      $dt.cf3 = cf3;
      return $dt;
    }
    static create_DC2(cf2) {
      let $dt = new D1(1);
      $dt.cf2 = cf2;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf2() { return this.cf2; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf2) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf3 === other.cf3;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf2, other.cf2);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC3(false);
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
    static create_DC4(cf4) {
      let $dt = new D2(0);
      $dt.cf4 = cf4;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get dtor_cf4() { return this.cf4; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC4" + "(" + _dafny.toString(this.cf4) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf4, other.cf4);
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
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC6(cf6, cf7, cf8) {
      let $dt = new D3(0);
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC5(cf5) {
      let $dt = new D3(1);
      $dt.cf5 = cf5;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf5() { return this.cf5; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC6" + "(" + this.cf6.toVerbatimString(true) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC5" + "(" + _dafny.toString(this.cf5) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf6, other.cf6) && _dafny.areEqual(this.cf7, other.cf7) && this.cf8 === other.cf8;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf5, other.cf5);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC6(_dafny.Seq.UnicodeFromString(""), new _dafny.CodePoint('D'.codePointAt(0)), false);
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
    static create_DC9(cf11) {
      let $dt = new D4(1);
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC10(cf12, cf13, cf14, cf15, cf16) {
      let $dt = new D4(2);
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC7(cf9) {
      let $dt = new D4(3);
      $dt.cf9 = cf9;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get is_DC10() { return this.$tag === 2; }
    get is_DC7() { return this.$tag === 3; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf9() { return this.cf9; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC8" + "(" + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC9" + "(" + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 3) {
        return "D4.DC7" + "(" + _dafny.toString(this.cf9) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf10 === other.cf10;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf11, other.cf11);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf12, other.cf12) && _dafny.areEqual(this.cf13, other.cf13) && this.cf14 === other.cf14 && this.cf15 === other.cf15 && _dafny.areEqual(this.cf16, other.cf16);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf9, other.cf9);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC8(false);
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
    static create_DC11(cf17) {
      let $dt = new D5(0);
      $dt.cf17 = cf17;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get dtor_cf17() { return this.cf17; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC11" + "(" + _dafny.toString(this.cf17) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf17 === other.cf17;
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
    static create_DC13(cf19, cf20) {
      let $dt = new D6(0);
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC12(cf18) {
      let $dt = new D6(1);
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC14(cf21) {
      let $dt = new D6(2);
      $dt.cf21 = cf21;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get is_DC14() { return this.$tag === 2; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf21() { return this.cf21; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC13" + "(" + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC12" + "(" + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC14" + "(" + _dafny.toString(this.cf21) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf19, other.cf19) && _dafny.areEqual(this.cf20, other.cf20);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf18, other.cf18);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf21, other.cf21);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC13(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC16(cf23) {
      let $dt = new D7(0);
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC15(cf22) {
      let $dt = new D7(1);
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC17(cf24) {
      let $dt = new D7(2);
      $dt.cf24 = cf24;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get is_DC17() { return this.$tag === 2; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf24() { return this.cf24; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC16" + "(" + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC15" + "(" + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC17" + "(" + _dafny.toString(this.cf24) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf23 === other.cf23;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf22, other.cf22);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf24, other.cf24);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC16([]);
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
    static create_DC18(cf25) {
      let $dt = new D8(0);
      $dt.cf25 = cf25;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get dtor_cf25() { return this.cf25; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC18" + "(" + _dafny.toString(this.cf25) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf25, other.cf25);
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
          return D8.Default();
        }
      };
    }
  }

  $module.D9 = class D9 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC19(cf26) {
      let $dt = new D9(0);
      $dt.cf26 = cf26;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get dtor_cf26() { return this.cf26; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC19" + "(" + _dafny.toString(this.cf26) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf26, other.cf26);
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
          return D9.Default();
        }
      };
    }
  }

  $module.D10 = class D10 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC20(cf27) {
      let $dt = new D10(0);
      $dt.cf27 = cf27;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get dtor_cf27() { return this.cf27; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC20" + "(" + _dafny.toString(this.cf27) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf27, other.cf27);
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
          return D10.Default();
        }
      };
    }
  }

  $module.D11 = class D11 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC21(cf28) {
      let $dt = new D11(0);
      $dt.cf28 = cf28;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get dtor_cf28() { return this.cf28; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC21" + "(" + _dafny.toString(this.cf28) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf28, other.cf28);
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
          return D11.Default();
        }
      };
    }
  }

  $module.D12 = class D12 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC23(cf30) {
      let $dt = new D12(0);
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC22(cf29) {
      let $dt = new D12(1);
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC24(cf31) {
      let $dt = new D12(2);
      $dt.cf31 = cf31;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC22() { return this.$tag === 1; }
    get is_DC24() { return this.$tag === 2; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf31() { return this.cf31; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC23" + "(" + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC22" + "(" + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC24" + "(" + _dafny.toString(this.cf31) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf30, other.cf30);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf29, other.cf29);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf31, other.cf31);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC23(_dafny.ZERO);
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
    static create_DC26(cf33, cf34, cf35) {
      let $dt = new D13(0);
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC27(cf36) {
      let $dt = new D13(1);
      $dt.cf36 = cf36;
      return $dt;
    }
    static create_DC25(cf32) {
      let $dt = new D13(2);
      $dt.cf32 = cf32;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get is_DC27() { return this.$tag === 1; }
    get is_DC25() { return this.$tag === 2; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf32() { return this.cf32; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC26" + "(" + this.cf33.toVerbatimString(true) + ", " + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC27" + "(" + _dafny.toString(this.cf36) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC25" + "(" + _dafny.toString(this.cf32) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf33, other.cf33) && this.cf34 === other.cf34 && this.cf35 === other.cf35;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf36 === other.cf36;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf32, other.cf32);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC26(_dafny.Seq.UnicodeFromString(""), false, false);
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
    static create_DC29(cf38, cf39) {
      let $dt = new D14(0);
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC30(cf40, cf41) {
      let $dt = new D14(1);
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC28(cf37) {
      let $dt = new D14(2);
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC31(cf42) {
      let $dt = new D14(3);
      $dt.cf42 = cf42;
      return $dt;
    }
    get is_DC29() { return this.$tag === 0; }
    get is_DC30() { return this.$tag === 1; }
    get is_DC28() { return this.$tag === 2; }
    get is_DC31() { return this.$tag === 3; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf42() { return this.cf42; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC29" + "(" + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC30" + "(" + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 2) {
        return "D14.DC28" + "(" + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 3) {
        return "D14.DC31" + "(" + _dafny.toString(this.cf42) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf38, other.cf38) && _dafny.areEqual(this.cf39, other.cf39);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf40, other.cf40) && _dafny.areEqual(this.cf41, other.cf41);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf37, other.cf37);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf42, other.cf42);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC29(_dafny.ZERO, _dafny.Map.Empty);
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
    static create_DC33(cf44) {
      let $dt = new D15(0);
      $dt.cf44 = cf44;
      return $dt;
    }
    static create_DC32(cf43) {
      let $dt = new D15(1);
      $dt.cf43 = cf43;
      return $dt;
    }
    static create_DC34(cf45) {
      let $dt = new D15(2);
      $dt.cf45 = cf45;
      return $dt;
    }
    get is_DC33() { return this.$tag === 0; }
    get is_DC32() { return this.$tag === 1; }
    get is_DC34() { return this.$tag === 2; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf45() { return this.cf45; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC33" + "(" + _dafny.toString(this.cf44) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC32" + "(" + _dafny.toString(this.cf43) + ")";
      } else if (this.$tag === 2) {
        return "D15.DC34" + "(" + _dafny.toString(this.cf45) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf44, other.cf44);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf43 === other.cf43;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf45, other.cf45);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC33(_dafny.Seq.of());
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
    static create_DC36(cf47, cf48, cf49) {
      let $dt = new D16(0);
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      $dt.cf49 = cf49;
      return $dt;
    }
    static create_DC35(cf46) {
      let $dt = new D16(1);
      $dt.cf46 = cf46;
      return $dt;
    }
    static create_DC37(cf50) {
      let $dt = new D16(2);
      $dt.cf50 = cf50;
      return $dt;
    }
    get is_DC36() { return this.$tag === 0; }
    get is_DC35() { return this.$tag === 1; }
    get is_DC37() { return this.$tag === 2; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf50() { return this.cf50; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC36" + "(" + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ", " + _dafny.toString(this.cf49) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC35" + "(" + _dafny.toString(this.cf46) + ")";
      } else if (this.$tag === 2) {
        return "D16.DC37" + "(" + _dafny.toString(this.cf50) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf47 === other.cf47 && _dafny.areEqual(this.cf48, other.cf48) && this.cf49 === other.cf49;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf46, other.cf46);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf50, other.cf50);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC36(false, _dafny.MultiSet.Empty, false);
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
    static create_DC39(cf52, cf53, cf54, cf55, cf56) {
      let $dt = new D17(0);
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      $dt.cf54 = cf54;
      $dt.cf55 = cf55;
      $dt.cf56 = cf56;
      return $dt;
    }
    static create_DC40() {
      let $dt = new D17(1);
      return $dt;
    }
    static create_DC41(cf57) {
      let $dt = new D17(2);
      $dt.cf57 = cf57;
      return $dt;
    }
    static create_DC38(cf51) {
      let $dt = new D17(3);
      $dt.cf51 = cf51;
      return $dt;
    }
    get is_DC39() { return this.$tag === 0; }
    get is_DC40() { return this.$tag === 1; }
    get is_DC41() { return this.$tag === 2; }
    get is_DC38() { return this.$tag === 3; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf51() { return this.cf51; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC39" + "(" + _dafny.toString(this.cf52) + ", " + _dafny.toString(this.cf53) + ", " + _dafny.toString(this.cf54) + ", " + _dafny.toString(this.cf55) + ", " + _dafny.toString(this.cf56) + ")";
      } else if (this.$tag === 1) {
        return "D17.DC40";
      } else if (this.$tag === 2) {
        return "D17.DC41" + "(" + _dafny.toString(this.cf57) + ")";
      } else if (this.$tag === 3) {
        return "D17.DC38" + "(" + _dafny.toString(this.cf51) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf52, other.cf52) && _dafny.areEqual(this.cf53, other.cf53) && this.cf54 === other.cf54 && _dafny.areEqual(this.cf55, other.cf55) && _dafny.areEqual(this.cf56, other.cf56);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf57, other.cf57);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf51 === other.cf51;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D17.create_DC39(_dafny.ZERO, _dafny.Map.Empty, false, _dafny.Map.Empty, _dafny.ZERO);
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
    static create_DC43() {
      let $dt = new D18(0);
      return $dt;
    }
    static create_DC44() {
      let $dt = new D18(1);
      return $dt;
    }
    static create_DC42(cf58) {
      let $dt = new D18(2);
      $dt.cf58 = cf58;
      return $dt;
    }
    get is_DC43() { return this.$tag === 0; }
    get is_DC44() { return this.$tag === 1; }
    get is_DC42() { return this.$tag === 2; }
    get dtor_cf58() { return this.cf58; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC43";
      } else if (this.$tag === 1) {
        return "D18.DC44";
      } else if (this.$tag === 2) {
        return "D18.DC42" + "(" + _dafny.toString(this.cf58) + ")";
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
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf58, other.cf58);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC43();
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
    static create_DC46(cf60, cf61) {
      let $dt = new D19(0);
      $dt.cf60 = cf60;
      $dt.cf61 = cf61;
      return $dt;
    }
    static create_DC47(cf62, cf63, cf64) {
      let $dt = new D19(1);
      $dt.cf62 = cf62;
      $dt.cf63 = cf63;
      $dt.cf64 = cf64;
      return $dt;
    }
    static create_DC48(cf65, cf66) {
      let $dt = new D19(2);
      $dt.cf65 = cf65;
      $dt.cf66 = cf66;
      return $dt;
    }
    static create_DC45(cf59) {
      let $dt = new D19(3);
      $dt.cf59 = cf59;
      return $dt;
    }
    get is_DC46() { return this.$tag === 0; }
    get is_DC47() { return this.$tag === 1; }
    get is_DC48() { return this.$tag === 2; }
    get is_DC45() { return this.$tag === 3; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf59() { return this.cf59; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC46" + "(" + _dafny.toString(this.cf60) + ", " + _dafny.toString(this.cf61) + ")";
      } else if (this.$tag === 1) {
        return "D19.DC47" + "(" + _dafny.toString(this.cf62) + ", " + _dafny.toString(this.cf63) + ", " + _dafny.toString(this.cf64) + ")";
      } else if (this.$tag === 2) {
        return "D19.DC48" + "(" + _dafny.toString(this.cf65) + ", " + _dafny.toString(this.cf66) + ")";
      } else if (this.$tag === 3) {
        return "D19.DC45" + "(" + _dafny.toString(this.cf59) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf60 === other.cf60 && _dafny.areEqual(this.cf61, other.cf61);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf62, other.cf62) && _dafny.areEqual(this.cf63, other.cf63) && this.cf64 === other.cf64;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf65, other.cf65) && this.cf66 === other.cf66;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf59, other.cf59);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D19.create_DC46(false, _dafny.ZERO);
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
    static create_DC50(cf68, cf69, cf70) {
      let $dt = new D20(0);
      $dt.cf68 = cf68;
      $dt.cf69 = cf69;
      $dt.cf70 = cf70;
      return $dt;
    }
    static create_DC51() {
      let $dt = new D20(1);
      return $dt;
    }
    static create_DC49(cf67) {
      let $dt = new D20(2);
      $dt.cf67 = cf67;
      return $dt;
    }
    get is_DC50() { return this.$tag === 0; }
    get is_DC51() { return this.$tag === 1; }
    get is_DC49() { return this.$tag === 2; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf67() { return this.cf67; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC50" + "(" + _dafny.toString(this.cf68) + ", " + _dafny.toString(this.cf69) + ", " + _dafny.toString(this.cf70) + ")";
      } else if (this.$tag === 1) {
        return "D20.DC51";
      } else if (this.$tag === 2) {
        return "D20.DC49" + "(" + _dafny.toString(this.cf67) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf68, other.cf68) && this.cf69 === other.cf69 && this.cf70 === other.cf70;
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf67, other.cf67);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC50(_dafny.ZERO, false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D20.Default();
        }
      };
    }
  }

  $module.D21 = class D21 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC53(cf72, cf73, cf74, cf75) {
      let $dt = new D21(0);
      $dt.cf72 = cf72;
      $dt.cf73 = cf73;
      $dt.cf74 = cf74;
      $dt.cf75 = cf75;
      return $dt;
    }
    static create_DC54(cf76, cf77, cf78) {
      let $dt = new D21(1);
      $dt.cf76 = cf76;
      $dt.cf77 = cf77;
      $dt.cf78 = cf78;
      return $dt;
    }
    static create_DC52(cf71) {
      let $dt = new D21(2);
      $dt.cf71 = cf71;
      return $dt;
    }
    static create_DC55(cf79) {
      let $dt = new D21(3);
      $dt.cf79 = cf79;
      return $dt;
    }
    get is_DC53() { return this.$tag === 0; }
    get is_DC54() { return this.$tag === 1; }
    get is_DC52() { return this.$tag === 2; }
    get is_DC55() { return this.$tag === 3; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf73() { return this.cf73; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf79() { return this.cf79; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC53" + "(" + _dafny.toString(this.cf72) + ", " + _dafny.toString(this.cf73) + ", " + _dafny.toString(this.cf74) + ", " + _dafny.toString(this.cf75) + ")";
      } else if (this.$tag === 1) {
        return "D21.DC54" + "(" + _dafny.toString(this.cf76) + ", " + _dafny.toString(this.cf77) + ", " + _dafny.toString(this.cf78) + ")";
      } else if (this.$tag === 2) {
        return "D21.DC52" + "(" + _dafny.toString(this.cf71) + ")";
      } else if (this.$tag === 3) {
        return "D21.DC55" + "(" + _dafny.toString(this.cf79) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf72 === other.cf72 && this.cf73 === other.cf73 && _dafny.areEqual(this.cf74, other.cf74) && this.cf75 === other.cf75;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf76, other.cf76) && _dafny.areEqual(this.cf77, other.cf77) && this.cf78 === other.cf78;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf71, other.cf71);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf79, other.cf79);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D21.create_DC53([], false, _dafny.MultiSet.Empty, []);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D21.Default();
        }
      };
    }
  }

  $module.D22 = class D22 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC57(cf81, cf82, cf83) {
      let $dt = new D22(0);
      $dt.cf81 = cf81;
      $dt.cf82 = cf82;
      $dt.cf83 = cf83;
      return $dt;
    }
    static create_DC58() {
      let $dt = new D22(1);
      return $dt;
    }
    static create_DC56(cf80) {
      let $dt = new D22(2);
      $dt.cf80 = cf80;
      return $dt;
    }
    static create_DC59(cf84) {
      let $dt = new D22(3);
      $dt.cf84 = cf84;
      return $dt;
    }
    get is_DC57() { return this.$tag === 0; }
    get is_DC58() { return this.$tag === 1; }
    get is_DC56() { return this.$tag === 2; }
    get is_DC59() { return this.$tag === 3; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf82() { return this.cf82; }
    get dtor_cf83() { return this.cf83; }
    get dtor_cf80() { return this.cf80; }
    get dtor_cf84() { return this.cf84; }
    toString() {
      if (this.$tag === 0) {
        return "D22.DC57" + "(" + _dafny.toString(this.cf81) + ", " + _dafny.toString(this.cf82) + ", " + _dafny.toString(this.cf83) + ")";
      } else if (this.$tag === 1) {
        return "D22.DC58";
      } else if (this.$tag === 2) {
        return "D22.DC56" + "(" + _dafny.toString(this.cf80) + ")";
      } else if (this.$tag === 3) {
        return "D22.DC59" + "(" + _dafny.toString(this.cf84) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf81, other.cf81) && _dafny.areEqual(this.cf82, other.cf82) && this.cf83 === other.cf83;
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf80, other.cf80);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf84, other.cf84);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D22.create_DC57(_dafny.ZERO, _dafny.ZERO, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D22.Default();
        }
      };
    }
  }

  $module.D23 = class D23 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC61(cf86, cf87, cf88) {
      let $dt = new D23(0);
      $dt.cf86 = cf86;
      $dt.cf87 = cf87;
      $dt.cf88 = cf88;
      return $dt;
    }
    static create_DC60(cf85) {
      let $dt = new D23(1);
      $dt.cf85 = cf85;
      return $dt;
    }
    static create_DC62(cf89) {
      let $dt = new D23(2);
      $dt.cf89 = cf89;
      return $dt;
    }
    get is_DC61() { return this.$tag === 0; }
    get is_DC60() { return this.$tag === 1; }
    get is_DC62() { return this.$tag === 2; }
    get dtor_cf86() { return this.cf86; }
    get dtor_cf87() { return this.cf87; }
    get dtor_cf88() { return this.cf88; }
    get dtor_cf85() { return this.cf85; }
    get dtor_cf89() { return this.cf89; }
    toString() {
      if (this.$tag === 0) {
        return "D23.DC61" + "(" + _dafny.toString(this.cf86) + ", " + _dafny.toString(this.cf87) + ", " + _dafny.toString(this.cf88) + ")";
      } else if (this.$tag === 1) {
        return "D23.DC60" + "(" + _dafny.toString(this.cf85) + ")";
      } else if (this.$tag === 2) {
        return "D23.DC62" + "(" + _dafny.toString(this.cf89) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf86, other.cf86) && _dafny.areEqual(this.cf87, other.cf87) && this.cf88 === other.cf88;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf85, other.cf85);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf89, other.cf89);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D23.create_DC61(_dafny.Seq.of(), _dafny.ZERO, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D23.Default();
        }
      };
    }
  }

  $module.D24 = class D24 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC64(cf91, cf92, cf93) {
      let $dt = new D24(0);
      $dt.cf91 = cf91;
      $dt.cf92 = cf92;
      $dt.cf93 = cf93;
      return $dt;
    }
    static create_DC65(cf94, cf95, cf96) {
      let $dt = new D24(1);
      $dt.cf94 = cf94;
      $dt.cf95 = cf95;
      $dt.cf96 = cf96;
      return $dt;
    }
    static create_DC63(cf90) {
      let $dt = new D24(2);
      $dt.cf90 = cf90;
      return $dt;
    }
    static create_DC66(cf97) {
      let $dt = new D24(3);
      $dt.cf97 = cf97;
      return $dt;
    }
    get is_DC64() { return this.$tag === 0; }
    get is_DC65() { return this.$tag === 1; }
    get is_DC63() { return this.$tag === 2; }
    get is_DC66() { return this.$tag === 3; }
    get dtor_cf91() { return this.cf91; }
    get dtor_cf92() { return this.cf92; }
    get dtor_cf93() { return this.cf93; }
    get dtor_cf94() { return this.cf94; }
    get dtor_cf95() { return this.cf95; }
    get dtor_cf96() { return this.cf96; }
    get dtor_cf90() { return this.cf90; }
    get dtor_cf97() { return this.cf97; }
    toString() {
      if (this.$tag === 0) {
        return "D24.DC64" + "(" + _dafny.toString(this.cf91) + ", " + _dafny.toString(this.cf92) + ", " + _dafny.toString(this.cf93) + ")";
      } else if (this.$tag === 1) {
        return "D24.DC65" + "(" + _dafny.toString(this.cf94) + ", " + _dafny.toString(this.cf95) + ", " + _dafny.toString(this.cf96) + ")";
      } else if (this.$tag === 2) {
        return "D24.DC63" + "(" + _dafny.toString(this.cf90) + ")";
      } else if (this.$tag === 3) {
        return "D24.DC66" + "(" + _dafny.toString(this.cf97) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf91, other.cf91) && this.cf92 === other.cf92 && this.cf93 === other.cf93;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf94, other.cf94) && _dafny.areEqual(this.cf95, other.cf95) && _dafny.areEqual(this.cf96, other.cf96);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf90 === other.cf90;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf97, other.cf97);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D24.create_DC64(_dafny.ZERO, false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D24.Default();
        }
      };
    }
  }

  $module.D25 = class D25 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC68() {
      let $dt = new D25(0);
      return $dt;
    }
    static create_DC69() {
      let $dt = new D25(1);
      return $dt;
    }
    static create_DC67(cf98) {
      let $dt = new D25(2);
      $dt.cf98 = cf98;
      return $dt;
    }
    static create_DC70(cf99) {
      let $dt = new D25(3);
      $dt.cf99 = cf99;
      return $dt;
    }
    get is_DC68() { return this.$tag === 0; }
    get is_DC69() { return this.$tag === 1; }
    get is_DC67() { return this.$tag === 2; }
    get is_DC70() { return this.$tag === 3; }
    get dtor_cf98() { return this.cf98; }
    get dtor_cf99() { return this.cf99; }
    toString() {
      if (this.$tag === 0) {
        return "D25.DC68";
      } else if (this.$tag === 1) {
        return "D25.DC69";
      } else if (this.$tag === 2) {
        return "D25.DC67" + "(" + _dafny.toString(this.cf98) + ")";
      } else if (this.$tag === 3) {
        return "D25.DC70" + "(" + _dafny.toString(this.cf99) + ")";
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
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf98, other.cf98);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf99, other.cf99);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D25.create_DC68();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D25.Default();
        }
      };
    }
  }

  $module.D26 = class D26 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC72(cf101) {
      let $dt = new D26(0);
      $dt.cf101 = cf101;
      return $dt;
    }
    static create_DC73(cf102, cf103) {
      let $dt = new D26(1);
      $dt.cf102 = cf102;
      $dt.cf103 = cf103;
      return $dt;
    }
    static create_DC71(cf100) {
      let $dt = new D26(2);
      $dt.cf100 = cf100;
      return $dt;
    }
    static create_DC74(cf104) {
      let $dt = new D26(3);
      $dt.cf104 = cf104;
      return $dt;
    }
    get is_DC72() { return this.$tag === 0; }
    get is_DC73() { return this.$tag === 1; }
    get is_DC71() { return this.$tag === 2; }
    get is_DC74() { return this.$tag === 3; }
    get dtor_cf101() { return this.cf101; }
    get dtor_cf102() { return this.cf102; }
    get dtor_cf103() { return this.cf103; }
    get dtor_cf100() { return this.cf100; }
    get dtor_cf104() { return this.cf104; }
    toString() {
      if (this.$tag === 0) {
        return "D26.DC72" + "(" + _dafny.toString(this.cf101) + ")";
      } else if (this.$tag === 1) {
        return "D26.DC73" + "(" + _dafny.toString(this.cf102) + ", " + _dafny.toString(this.cf103) + ")";
      } else if (this.$tag === 2) {
        return "D26.DC71" + "(" + _dafny.toString(this.cf100) + ")";
      } else if (this.$tag === 3) {
        return "D26.DC74" + "(" + _dafny.toString(this.cf104) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf101, other.cf101);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf102 === other.cf102 && this.cf103 === other.cf103;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf100, other.cf100);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf104, other.cf104);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D26.create_DC72(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D26.Default();
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
      this.f0 = false;
      this.f2 = false;
      this.f7 = _dafny.ZERO;
      this.f8 = _dafny.Seq.UnicodeFromString("");
      this.f10 = _dafny.Set.Empty;
      this.f13 = false;
      this.f16 = _dafny.ZERO;
      this.f19 = _dafny.Seq.UnicodeFromString("");
      this.f20 = _dafny.Seq.UnicodeFromString("");
      this.f21 = false;
      this.f22 = [];
      this._f1 = [];
      this._f3 = _dafny.ZERO;
      this._f4 = false;
      this._f5 = false;
      this._f6 = false;
      this._f9 = _dafny.ZERO;
      this._f11 = _dafny.ZERO;
      this._f12 = _dafny.ZERO;
      this._f14 = _dafny.Seq.UnicodeFromString("");
      this._f15 = _dafny.ZERO;
      this._f17 = false;
      this._f18 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22) {
      let _this = this;
      (_this).f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this).f7 = f7;
      (_this).f8 = f8;
      (_this)._f9 = f9;
      (_this).f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this).f13 = f13;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      (_this).f16 = f16;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      (_this).f19 = f19;
      (_this).f20 = f20;
      (_this).f21 = f21;
      (_this).f22 = f22;
      return;
    }
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
    get f9() {
      let _this = this;
      return _this._f9;
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
      this.f27 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f27) {
      let _this = this;
      (_this).f27 = f27;
      return;
    }
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this.f29 = _dafny.ZERO;
      this._f30 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    __ctor(f29, f30) {
      let _this = this;
      (_this).f29 = f29;
      (_this)._f30 = f30;
      return;
    }
    fm7(p0, p1, p2, globalState) {
      let _this = this;
      return true;
    };
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe((false) === (!(true)),_dafny.Seq.UnicodeFromString("ctdvjyvbi"));
    };
    fm6(p0, globalState) {
      let _this = this;
      return (_this.f29).isEqualTo(_this.f29);
    };
    fm29(globalState) {
      let _this = this;
      return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(false, false, true)).cardinality())),(_dafny.Map.Empty.slice().updateUnsafe(true,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,true)))).length);
    };
    m2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = [];
      if (p1) {
        if (p1) {
          (globalState).f2 = p1;
          let _238_v0;
          let _nw21 = new _module.C0();
          _nw21.__ctor(!(true));
          _238_v0 = _nw21;
          let _239_v1;
          let _nw22 = new _module.C0();
          _nw22.__ctor(p1);
          _239_v1 = _nw22;
          let _240_v2;
          _240_v2 = new _dafny.CodePoint('l'.codePointAt(0));
          let _241_v3;
          _241_v3 = _dafny.Seq.of(_240_v2);
          let _242_v4;
          _242_v4 = _dafny.MultiSet.fromElements(_241_v3);
          let _243_v5;
          _243_v5 = _dafny.Seq.of(_module.__default.safeModuloInt(p0, new BigNumber((_242_v4).cardinality())));
          let _244_v6;
          let _nw23 = Array((new BigNumber(14)).toNumber()).fill([]);
          _244_v6 = _nw23;
          let _245_v7;
          let _nw24 = Array((new BigNumber(3)).toNumber()).fill(false);
          _245_v7 = _nw24;
          let _index33 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_244_v6).length));
          (_244_v6)[_index33] = _245_v7;
          let _246_v8;
          _246_v8 = _dafny.Map.Empty.slice().updateUnsafe(p1,_239_v1.f27);
          let _247_v9;
          let _nw25 = Array((_dafny.ONE).toNumber());
          _nw25[(_dafny.ZERO).toNumber()] = (_module.__default.fm0(p3, new BigNumber(438), new BigNumber((_241_v3).length), p1, globalState)).minus(new BigNumber((_246_v8).length));
          _247_v9 = _nw25;
          let _248_v10;
          _248_v10 = _dafny.Map.Empty.slice().updateUnsafe(p0,_238_v0.f27);
          let _index34 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_244_v6).length));
          let _rhs25 = _dafny.Seq.of(p3, new BigNumber((_248_v10).length), p3);
          let _rhs26 = _245_v7;
          let _rhs27 = p3;
          let _rhs28 = _247_v9;
          let _lhs30 = _244_v6;
          let _lhs31 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_244_v6).length));
          let _lhs32 = globalState;
          _243_v5 = _rhs25;
          _lhs30[_lhs31] = _rhs26;
          _lhs32.f7 = _rhs27;
          _247_v9 = _rhs28;
          (globalState).f21 = _238_v0.f27;
        } else {
          let _249_v11;
          let _init6 = function (_250_i0) {
            return (_250_i0).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("ks")).length));
          };
          let _nw26 = Array((new BigNumber(29)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw26.length); _i0_6++) {
            _nw26[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _249_v11 = _nw26;
          let _nw27 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
          _249_v11 = _nw27;
          (globalState).f0 = p1;
          let _251_v12;
          _251_v12 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
          (globalState).f7 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt(p0, new BigNumber((_251_v12).length)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(p1),true)).length));
          let _252_v13;
          _252_v13 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
          let _253_v14;
          _253_v14 = _dafny.MultiSet.fromElements(_this.f29, (_module.__default.fm0(p0, p3, new BigNumber(758), p1, globalState)).multipliedBy(new BigNumber((_252_v13).length)), p0);
          _253_v14 = _253_v14;
          let _254_v15;
          _254_v15 = _module.D3.create_DC5(new _dafny.CodePoint('m'.codePointAt(0)));
          let _255_v16;
          _255_v16 = _dafny.MultiSet.fromElements(_254_v15);
          let _256_v17;
          _256_v17 = _dafny.Set.fromElements(p3);
          let _257_v18;
          _257_v18 = _dafny.Seq.of(p1, !((_module.__default.fm30(p1, _this.f29, (_dafny.ZERO).minus(_this.f29), _255_v16, globalState)).equals(_256_v17)), p1, p1);
          _257_v18 = _257_v18;
        }
        let _258_v19;
        _258_v19 = _dafny.Map.Empty.slice().updateUnsafe(p1,p3);
        let _rhs29 = new BigNumber((((p1) ? (_258_v19) : (_258_v19))).length);
        let _rhs30 = (_dafny.ZERO).minus(new BigNumber((_module.__default.fm31(p0, !(p1) || (p1), globalState)).length));
        let _lhs33 = globalState;
        _lhs33.f16 = _rhs29;
        r1 = _rhs30;
        if (p1) {
          (globalState).f2 = p1;
          let _259_v20;
          _259_v20 = _dafny.Seq.of(p1);
          let _260_v21;
          _260_v21 = _dafny.Map.Empty.slice().updateUnsafe(p0,_259_v20);
          let _261_v22;
          _261_v22 = _dafny.Map.Empty.slice().updateUnsafe((((_260_v21).contains(new BigNumber(142))) ? ((_260_v21).get(new BigNumber(142))) : (_259_v20)),_module.__default.fm1(_this.f29, globalState));
          _module.__default.m0(p1, p0, _261_v22, (_this).fm29(globalState), globalState);
          let _262_v23;
          _262_v23 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_259_v20, _dafny.Seq.of(p1, p1, p1, p1, p1)),false);
          let _263_v24;
          _263_v24 = _dafny.Set.fromElements(_module.__default.fm0(p3, p3, p3, p1, globalState), new BigNumber((_dafny.Seq.update(_259_v20, _module.__default.safeIndex(p3, new BigNumber((_259_v20).length)), p1)).length), _this.f29, _this.f29);
          let _264_v25;
          _264_v25 = _module.D14.create_DC28(_263_v24);
          (globalState).f2 = (((_262_v23).contains(_259_v20)) ? ((_262_v23).get(_259_v20)) : (((_264_v25).dtor_cf37).IsProperSubsetOf(_263_v24)));
          let _265_v26;
          _265_v26 = _dafny.Seq.UnicodeFromString("ifljbrg");
          let _266_v27;
          _266_v27 = _dafny.Seq.of(_this.f29, new BigNumber((_265_v26).length));
          (globalState).f7 = _module.__default.fm0(p0, (_266_v27)[_module.__default.safeIndex(p0, new BigNumber((_266_v27).length))], new BigNumber(392), p1, globalState);
          let _index35 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((p2).length));
          (p2)[_index35] = p0;
          let _index36 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((p2).length));
          (p2)[_index36] = p3;
        } else {
          let _267_v28;
          _267_v28 = _dafny.Seq.UnicodeFromString("gtrohbu");
          let _268_v29;
          _268_v29 = _dafny.Map.Empty.slice().updateUnsafe(p1,_267_v28);
          let _269_v30;
          _269_v30 = new _dafny.CodePoint('w'.codePointAt(0));
          (globalState).f8 = _dafny.Seq.Concat(_dafny.Seq.update((((_268_v29).contains(p1)) ? ((_268_v29).get(p1)) : (_267_v28)), _module.__default.safeIndex(_this.f29, new BigNumber(((((_268_v29).contains(p1)) ? ((_268_v29).get(p1)) : (_267_v28))).length)), _269_v30), _267_v28);
          let _270_v31;
          let _nw28 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Set.Empty);
          _270_v31 = _nw28;
          let _271_v32;
          _271_v32 = _dafny.Map.Empty.slice().updateUnsafe(p3,_dafny.Seq.of(_this.f29, new BigNumber((_267_v28).length)));
          let _272_v33;
          _272_v33 = _dafny.Seq.of(p1);
          let _273_v34;
          _273_v34 = _module.D12.create_DC23(new BigNumber(528));
          let _274_v35;
          _274_v35 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
          let _275_v36;
          _275_v36 = _dafny.Seq.of(p0, (_273_v34).dtor_cf30, new BigNumber((_274_v35).length));
          let _276_v37;
          _276_v37 = _dafny.Set.fromElements(p0, p0, new BigNumber(((((_271_v32).contains(new BigNumber((_272_v33).length))) ? ((_271_v32).get(new BigNumber((_272_v33).length))) : (_dafny.Seq.of((_275_v36)[_module.__default.safeIndex(p0, new BigNumber((_275_v36).length))], p0)))).length));
          let _index37 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_270_v31).length));
          (_270_v31)[_index37] = _276_v37;
          let _index38 = _module.__default.safeIndex(new BigNumber(53), new BigNumber((_270_v31).length));
          (_270_v31)[_index38] = _276_v37;
          let _277_v38;
          let _nw29 = Array((new BigNumber(3)).toNumber()).fill(_module.D14.Default());
          _277_v38 = _nw29;
          let _278_v39;
          _278_v39 = _module.D14.create_DC28((_270_v31)[_module.__default.safeIndex(new BigNumber(53), new BigNumber((_270_v31).length))]);
          let _index39 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_277_v38).length));
          (_277_v38)[_index39] = _278_v39;
          let _279_v40;
          _279_v40 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements(_274_v35, _274_v35, _274_v35, _dafny.Map.Empty.slice().updateUnsafe((_this).fm6(p0, globalState),true), _274_v35)).IsDisjointFrom(_dafny.Set.fromElements((_274_v35).update(p1, p1), _274_v35)),p2);
          let _280_v41;
          _280_v41 = _dafny.MultiSet.fromElements(p1, (_272_v33)[_module.__default.safeIndex(new BigNumber(522), new BigNumber((_272_v33).length))]);
          let _index40 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_277_v38).length));
          let _rhs31 = p1;
          let _rhs32 = _module.__default.fm32(!_dafny.Seq.contains(_dafny.Seq.of(p1, false), !(p1)), new BigNumber((_280_v41).cardinality()), p3, globalState);
          let _rhs33 = (_279_v40).Merge(_279_v40);
          let _lhs34 = globalState;
          let _lhs35 = _277_v38;
          let _lhs36 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_277_v38).length));
          _lhs34.f2 = _rhs31;
          _lhs35[_lhs36] = _rhs32;
          _279_v40 = _rhs33;
          (globalState).f2 = p1;
          let _281_v42;
          _281_v42 = _dafny.Set.fromElements(p1, p1, (_this).fm7(p1, p3, p1, globalState));
          let _282_v43;
          _282_v43 = _dafny.MultiSet.fromElements(_281_v42);
          let _283_v44;
          _283_v44 = _281_v42;
          let _index41 = _module.__default.safeIndex(new BigNumber(618), new BigNumber((p2).length));
          (p2)[_index41] = (((_282_v43).contains(_283_v44)) ? ((_282_v43).get(_283_v44)) : (p0));
          let _284_v45;
          _284_v45 = _dafny.MultiSet.fromElements((p3).plus(_this.f29), _this.f29, (_this.f29).plus(p0));
          let _index42 = _module.__default.safeIndex(new BigNumber(618), new BigNumber((p2).length));
          (p2)[_index42] = (_dafny.ZERO).minus((((_284_v45).contains(p3)) ? ((_284_v45).get(p3)) : (_module.__default.safeDivisionInt((((_258_v19).contains(p1)) ? ((_258_v19).get(p1)) : (p3)), _this.f29))));
        }
        (globalState).f13 = (new BigNumber((_dafny.Seq.of(p0, new BigNumber((_258_v19).length), _this.f29, p3, _this.f29)).length)).isLessThanOrEqualTo(p0);
        let _285_v46;
        _285_v46 = new _dafny.CodePoint('t'.codePointAt(0));
        let _286_v47;
        _286_v47 = _dafny.Seq.of(_285_v46, _285_v46);
        r1 = _module.__default.safeModuloInt(p0, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(141)), function (_287_i1) {
          return new _dafny.CodePoint('s'.codePointAt(0));
        }), _286_v47)).length));
      } else {
        let _288_v48;
        _288_v48 = _dafny.Seq.UnicodeFromString("n");
        let _289_v49;
        _289_v49 = new _dafny.CodePoint('i'.codePointAt(0));
        let _290_v50;
        _290_v50 = _module.D3.create_DC6(_288_v48, _289_v49, p1);
        (globalState).f7 = new BigNumber((_dafny.Seq.of(_290_v50)).length);
        let _index43 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((p2).length));
        (p2)[_index43] = (_dafny.ZERO).minus((_dafny.ZERO).minus(p3));
        let _291_v51;
        _291_v51 = _module.D4.create_DC9(p3);
        let _292_v52;
        _292_v52 = _dafny.Seq.of(p1, p1, p1);
        let _index44 = _module.__default.safeIndex(new BigNumber(348), new BigNumber((p2).length));
        (p2)[_index44] = (p0).multipliedBy(((_291_v51).dtor_cf11).plus(new BigNumber((_292_v52).length)));
        let _293_v53;
        _293_v53 = _dafny.Map.Empty.slice().updateUnsafe(p1,true);
        (globalState).f21 = _module.__default.fm2((_293_v53).Merge(_293_v53), globalState);
        (globalState).f7 = (_dafny.ZERO).minus((_this.f29).minus((_dafny.ZERO).minus((p0).multipliedBy(new BigNumber(549)))));
        let _294_v54;
        let _nw30 = Array((new BigNumber(21)).toNumber()).fill(false);
        _294_v54 = _nw30;
        let _index45 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_294_v54).length));
        (_294_v54)[_index45] = !(p1);
        let _295_v55;
        _295_v55 = _dafny.Seq.of((_dafny.ZERO).minus((p2)[_module.__default.safeIndex(new BigNumber(348), new BigNumber((p2).length))]));
        let _296_v56;
        _296_v56 = _dafny.Map.Empty.slice().updateUnsafe(true,p3);
        let _297_v57;
        _297_v57 = _dafny.Map.Empty.slice().updateUnsafe(_288_v48,_dafny.MultiSet.fromElements((_295_v55)[_module.__default.safeIndex(new BigNumber((_296_v56).length), new BigNumber((_295_v55).length))], (p2)[_module.__default.safeIndex(new BigNumber(348), new BigNumber((p2).length))]));
        let _298_v58;
        _298_v58 = _dafny.MultiSet.fromElements(p0, new BigNumber((_288_v48).length));
        let _index46 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_294_v54).length));
        let _rhs34 = new BigNumber(((((_297_v57).contains(_288_v48)) ? ((_297_v57).get(_288_v48)) : (_298_v58))).cardinality());
        let _rhs35 = p1;
        let _lhs37 = _this;
        let _lhs38 = _294_v54;
        let _lhs39 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_294_v54).length));
        _lhs37.f29 = _rhs34;
        _lhs38[_lhs39] = _rhs35;
      }
      let _299_v59;
      _299_v59 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
      _299_v59 = (_299_v59).update(p2, p0);
      let _300_v60;
      _300_v60 = _dafny.Seq.of(_this.f29);
      let _301_v61;
      let _nw31 = new _module.C0();
      _nw31.__ctor(p1);
      _301_v61 = _nw31;
      let _302_v62;
      _302_v62 = _dafny.Seq.of(p1, p1, p1, p1);
      let _303_v63;
      _303_v63 = _dafny.Map.Empty.slice().updateUnsafe((_302_v62)[_module.__default.safeIndex(new BigNumber((_302_v62).length), new BigNumber((_302_v62).length))],_301_v61.f27);
      let _304_v64;
      _304_v64 = _dafny.Map.Empty.slice().updateUnsafe(_301_v61,_303_v63);
      let _305_v65;
      _305_v65 = _dafny.Seq.UnicodeFromString("fgkj");
      let _306_v66;
      _306_v66 = _dafny.Map.Empty.slice().updateUnsafe(_302_v62,_305_v65);
      _module.__default.m0(p1, _module.__default.safeDivisionInt(new BigNumber((_300_v60).length), _module.__default.fm0(new BigNumber(-680), new BigNumber(((((_304_v64).contains(_301_v61)) ? ((_304_v64).get(_301_v61)) : (_303_v63))).length), p3, _301_v61.f27, globalState)), _306_v66, new BigNumber((_300_v60).length), globalState);
      let _307_v68;
      _307_v68 = new _dafny.CodePoint('w'.codePointAt(0));
      let _308_v69;
      _308_v69 = _dafny.Map.Empty.slice().updateUnsafe(p3,_307_v68);
      let _309_v70;
      _309_v70 = _dafny.Map.Empty.slice().updateUnsafe(_this.f29,new BigNumber((_308_v69).length));
      let _310_v72;
      _310_v72 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((function () {
        let _coll28 = new _dafny.Set();
        for (const _compr_28 of (_dafny.Seq.of(_309_v70, _309_v70)).Elements) {
          let _311_v71 = _compr_28;
          if (_dafny.Seq.contains(_dafny.Seq.of(_309_v70, _309_v70), _311_v71)) {
            _coll28.add(_311_v71);
          }
        }
        return _coll28;
      }()).length));
      if ((new BigNumber(((function () {
        let _coll29 = new _dafny.Map();
        for (const _compr_29 of (_dafny.Map.Empty.slice().updateUnsafe((((_310_v72).contains(p3)) ? ((_310_v72).get(p3)) : (p0)),p0)).Keys.Elements) {
          let _312_v67 = _compr_29;
          if ((_dafny.Map.Empty.slice().updateUnsafe((((_310_v72).contains(p3)) ? ((_310_v72).get(p3)) : (p0)),p0)).contains(_312_v67)) {
            _coll29.push([_module.__default.safeModuloInt(_312_v67, (_300_v60)[_module.__default.safeIndex(p0, new BigNumber((_300_v60).length))]),p0]);
          }
        }
        return _coll29;
      }()).update(_this.f29, p0)).length)).isLessThanOrEqualTo(p3)) {
        (_this).f29 = p3;
        let _313_v73;
        let _nw32 = Array((_dafny.ONE).toNumber()).fill(false);
        _313_v73 = _nw32;
        let _314_v74;
        _314_v74 = _dafny.MultiSet.fromElements(p3, p0);
        let _index47 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_313_v73).length));
        (_313_v73)[_index47] = (_314_v74).IsProperSubsetOf(_module.__default.fm33(globalState));
        let _index48 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_313_v73).length));
        (_313_v73)[_index48] = (((((_309_v70).contains(p0)) ? ((_309_v70).get(p0)) : (p3))).plus(_this.f29)).isLessThanOrEqualTo((_this).fm29(globalState));
        (globalState).f13 = p1;
        let _315_v75;
        _315_v75 = _dafny.MultiSet.fromElements(_301_v61.f27, (_313_v73)[_module.__default.safeIndex(new BigNumber(758), new BigNumber((_313_v73).length))]);
        if (((_315_v75).update(p1, _module.__default.abs((_dafny.ZERO).minus(p0)))).IsSubsetOf(_315_v75)) {
          r1 = _this.f29;
          let _316_v76;
          let _nw33 = Array((new BigNumber(4)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _316_v76 = _nw33;
          let _index49 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_316_v76).length));
          (_316_v76)[_index49] = _307_v68;
          let _index50 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_316_v76).length));
          (_316_v76)[_index50] = _module.__default.fm34(globalState);
          let _317_v77;
          _317_v77 = _dafny.MultiSet.fromElements(_dafny.Seq.of(true, p1), _302_v62);
          (globalState).f0 = !((_dafny.MultiSet.FromArray(_dafny.Seq.update(_dafny.Seq.of(_302_v62, _302_v62, _dafny.Seq.of(p1, p1, false), _302_v62), _module.__default.safeIndex(new BigNumber((_300_v60).length), new BigNumber((_dafny.Seq.of(_302_v62, _302_v62, _dafny.Seq.of(p1, p1, false), _302_v62)).length)), _302_v62))).IsProperSubsetOf(_317_v77));
          _module.__default.m0(!_dafny.Seq.contains(_305_v65, _307_v68), (p3).plus(_this.f29), _306_v66, new BigNumber((_315_v75).cardinality()), globalState);
          (globalState).f21 = !(!(_this.f29).isEqualTo((_this).fm29(globalState)));
        } else {
          let _nw34 = new _module.C0();
          _nw34.__ctor(p1);
          _301_v61 = _nw34;
          let _318_v78;
          let _nw35 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.of());
          _318_v78 = _nw35;
          let _319_v79;
          _319_v79 = _module.D15.create_DC32(_318_v78);
          _318_v78 = (_319_v79).dtor_cf43;
          let _320_v80;
          _320_v80 = _module.D14.create_DC28(_dafny.Set.fromElements(_this.f29));
          let _321_v81;
          _321_v81 = _module.D7.create_DC15(_300_v60);
          let _322_v82;
          _322_v82 = _module.D7.create_DC17(_321_v81);
          let _323_v83;
          _323_v83 = _module.D7.create_DC17(_322_v82);
          let _324_v84;
          _324_v84 = _module.D7.create_DC17(_322_v82);
          let _325_v85;
          _325_v85 = _module.D7.create_DC17(_322_v82);
          let _326_v86;
          _326_v86 = _module.D7.create_DC17(_323_v83);
          let _327_v87;
          _327_v87 = _dafny.Map.Empty.slice().updateUnsafe(_320_v80,_326_v86);
          let _328_v88;
          _328_v88 = _dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_300_v60)).cardinality()));
          _327_v87 = (_327_v87).update(_module.D14.create_DC28(_328_v88), _module.D7.create_DC17(_324_v84));
          let _329_v89;
          let _nw36 = Array((new BigNumber(21)).toNumber());
          _nw36[(_dafny.ZERO).toNumber()] = _305_v65;
          _nw36[(_dafny.ONE).toNumber()] = _305_v65;
          _nw36[(new BigNumber(2)).toNumber()] = _305_v65;
          _nw36[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_305_v65, _305_v65);
          _nw36[(new BigNumber(4)).toNumber()] = _dafny.Seq.of(_307_v68, _307_v68, _307_v68);
          _nw36[(new BigNumber(5)).toNumber()] = _305_v65;
          _nw36[(new BigNumber(6)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(951)), ((_330_v68) => function (_331_i2) {
            return _330_v68;
          })(_307_v68));
          _nw36[(new BigNumber(7)).toNumber()] = _305_v65;
          _nw36[(new BigNumber(8)).toNumber()] = _dafny.Seq.of(_307_v68, _307_v68, _307_v68, _307_v68, _307_v68);
          _nw36[(new BigNumber(9)).toNumber()] = _305_v65;
          _nw36[(new BigNumber(10)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(631)), ((_332_v68) => function (_333_i3) {
            return _332_v68;
          })(_307_v68));
          _nw36[(new BigNumber(11)).toNumber()] = _305_v65;
          _nw36[(new BigNumber(12)).toNumber()] = _305_v65;
          _nw36[(new BigNumber(13)).toNumber()] = _305_v65;
          _nw36[(new BigNumber(14)).toNumber()] = _305_v65;
          _nw36[(new BigNumber(15)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-368)), ((_334_v68) => function (_335_i4) {
            return _334_v68;
          })(_307_v68));
          _nw36[(new BigNumber(16)).toNumber()] = _305_v65;
          _nw36[(new BigNumber(17)).toNumber()] = _305_v65;
          _nw36[(new BigNumber(18)).toNumber()] = _305_v65;
          _nw36[(new BigNumber(19)).toNumber()] = _dafny.Seq.of(_307_v68);
          _nw36[(new BigNumber(20)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-300)), ((_336_v68) => function (_337_i5) {
            return _336_v68;
          })(_307_v68));
          _329_v89 = _nw36;
          let _index51 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_329_v89).length));
          (_329_v89)[_index51] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(140)), ((_338_v68) => function (_339_i6) {
            return _338_v68;
          })(_307_v68));
          let _340_v90;
          _340_v90 = _module.D3.create_DC5(_307_v68);
          let _index52 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_329_v89).length));
          (_329_v89)[_index52] = ((_301_v61.f27) ? (_305_v65) : (_dafny.Seq.of((_340_v90).dtor_cf5)));
          let _341_v91;
          let _nw37 = new _module.C0();
          _nw37.__ctor(_301_v61.f27);
          _341_v91 = _nw37;
        }
        let _342_v92;
        _342_v92 = _dafny.Seq.of(((_301_v61.f27) ? (_313_v73) : (_313_v73)), _313_v73);
        _313_v73 = (_342_v92)[_module.__default.safeIndex(p0, new BigNumber((_342_v92).length))];
      } else {
        _307_v68 = _307_v68;
        let _343_v93;
        let _nw38 = new _module.C0();
        _nw38.__ctor(_301_v61.f27);
        _343_v93 = _nw38;
        let _344_v94;
        _344_v94 = _dafny.MultiSet.fromElements(p3, p3);
        (globalState).f2 = (_module.__default.fm0(new BigNumber(268), p3, new BigNumber((_344_v94).cardinality()), _301_v61.f27, globalState)).isLessThan((new BigNumber(36)).plus(_this.f29));
        let _345_v95;
        _345_v95 = _dafny.Map.Empty.slice().updateUnsafe(_301_v61.f27,p0);
        let _346_v96;
        _346_v96 = _dafny.Map.Empty.slice().updateUnsafe(((_343_v93.f27) ? (_345_v95) : (_dafny.Map.Empty.slice().updateUnsafe(!(p1),p3))),p3);
        let _347_v97;
        _347_v97 = _dafny.Map.Empty.slice().updateUnsafe(p0,_module.D14.create_DC30(_300_v60, _302_v62));
        _346_v96 = (_346_v96).update((_module.__default.fm35(p1, p3, globalState)).Merge(_345_v95), (_300_v60)[_module.__default.safeIndex(new BigNumber((_347_v97).length), new BigNumber((_300_v60).length))]);
        (globalState).f7 = p3;
      }
      let _348_v98;
      _348_v98 = _dafny.Map.Empty.slice().updateUnsafe(p3,_301_v61.f27);
      let _349_v99;
      _349_v99 = _dafny.Map.Empty.slice().updateUnsafe(_this.f29,(_this).fm7(!((((_348_v98).contains(_this.f29)) ? ((_348_v98).get(_this.f29)) : (_301_v61.f27))), p0, _301_v61.f27, globalState));
      (globalState).f0 = (new BigNumber((_349_v99).length)).isEqualTo(p3);
      let _350_v100;
      _350_v100 = _dafny.Map.Empty.slice().updateUnsafe(_301_v61.f27,p3);
      let _351_v101;
      _351_v101 = _dafny.MultiSet.fromElements(p2);
      let _352_v102;
      _352_v102 = _dafny.MultiSet.fromElements(p0, p3, _this.f29);
      _module.__default.m0(false, p0, _module.__default.fm36(new BigNumber((_350_v100).length), p0, (((_351_v101).contains(p2)) ? ((_351_v101).get(p2)) : (p0)), globalState), new BigNumber((_352_v102).cardinality()), globalState);
      r0 = p1;
      let _353_v103;
      _353_v103 = _dafny.MultiSet.fromElements(_301_v61.f27, p1);
      r1 = (((_353_v103).contains(_301_v61.f27)) ? ((_353_v103).get(_301_v61.f27)) : (new BigNumber(359)));
      let _354_v104;
      let _nw39 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Map.Empty);
      _354_v104 = _nw39;
      r2 = _354_v104;
      return [r0, r1, r2];
    }
    m3(p0, globalState) {
      let _this = this;
      let _355_v0;
      _355_v0 = _dafny.Map.Empty.slice().updateUnsafe(_this.f29,_this.f29);
      let _356_v1;
      _356_v1 = _dafny.Seq.UnicodeFromString("kx");
      let _357_v2;
      _357_v2 = _dafny.Map.Empty.slice().updateUnsafe(_355_v0,_356_v1);
      let _358_v3;
      let _init7 = function (_359_i1) {
        return false;
      };
      let _nw40 = Array((new BigNumber(5)).toNumber());
      for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw40.length); _i0_7++) {
        _nw40[_i0_7] = _init7(new BigNumber(_i0_7));
      }
      _358_v3 = _nw40;
      (globalState).f0 = (_dafny.Map.Empty.slice().updateUnsafe(_this.f29,_358_v3)).contains((new BigNumber(((((_357_v2).contains(_355_v0)) ? ((_357_v2).get(_355_v0)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-463)), function (_360_i0) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      })))).length)).plus(_this.f29));
      let _index53 = _module.__default.safeIndex(new BigNumber(727), new BigNumber((p0).length));
      (p0)[_index53] = _this.f29;
      let _index54 = _module.__default.safeIndex(new BigNumber(727), new BigNumber((p0).length));
      (p0)[_index54] = (new BigNumber((_356_v1).length)).minus(_this.f29);
      let _361_v4;
      _361_v4 = _dafny.MultiSet.fromElements(_358_v3);
      let _362_i2;
      _362_i2 = _dafny.ZERO;
      L2: {
        while ((new BigNumber((_361_v4).cardinality())).isEqualTo(new BigNumber(523))) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_362_i2)) {
              break L2;
            }
            _362_i2 = (_362_i2).plus(_dafny.ONE);
            let _index55 = _module.__default.safeIndex(new BigNumber(727), new BigNumber((p0).length));
            (p0)[_index55] = ((p0)[_module.__default.safeIndex(new BigNumber(727), new BigNumber((p0).length))]).plus(_this.f29);
            let _363_v5;
            _363_v5 = new _dafny.CodePoint('d'.codePointAt(0));
            _363_v5 = _363_v5;
            let _364_v6;
            _364_v6 = true;
            let _365_v7;
            let _nw41 = new _module.C0();
            _nw41.__ctor(!(_364_v6));
            _365_v7 = _nw41;
            let _366_v8;
            _366_v8 = _dafny.MultiSet.fromElements(_365_v7.f27);
            let _367_v9;
            _367_v9 = _dafny.Map.Empty.slice().updateUnsafe((((_366_v8).contains(_364_v6)) ? ((_366_v8).get(_364_v6)) : ((p0)[_module.__default.safeIndex(new BigNumber(727), new BigNumber((p0).length))])),_365_v7.f27);
            let _368_v10;
            _368_v10 = _dafny.Seq.of(_this.f29, ((p0)[_module.__default.safeIndex(new BigNumber(727), new BigNumber((p0).length))]).multipliedBy((p0)[_module.__default.safeIndex(new BigNumber(727), new BigNumber((p0).length))]), (new BigNumber((_367_v9).length)).plus(_this.f29), _this.f29, (p0)[_module.__default.safeIndex(new BigNumber(727), new BigNumber((p0).length))]);
            _368_v10 = _dafny.Seq.Concat(_module.__default.fm37(globalState), _module.__default.fm37(globalState));
          }
        }
      }
      let _369_v11;
      let _nw42 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
      _369_v11 = _nw42;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_369_v11).length))) {
        let _370_i3 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_370_i3)) && ((_370_i3).isLessThan(new BigNumber((_369_v11).length))))) {
          (_369_v11)[(_370_i3)] = (_370_i3).minus((_this.f29).minus(_this.f29));
        }
      }
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_358_v3).length))) {
        let _371_i4 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_371_i4)) && ((_371_i4).isLessThan(new BigNumber((_358_v3).length))))) {
          (_358_v3)[(_371_i4)] = true;
        }
      }
      let _372_v12;
      _372_v12 = false;
      let _373_v13;
      _373_v13 = _dafny.Seq.of(false, _372_v12, _372_v12, (_this).fm6(new BigNumber(-244), globalState), _372_v12);
      let _374_v14;
      _374_v14 = _dafny.Map.Empty.slice().updateUnsafe(!(_372_v12),_373_v13);
      let _hi0 = ((p0)[_module.__default.safeIndex(new BigNumber(727), new BigNumber((p0).length))]).multipliedBy(_this.f29);
      for (let _375_i5 = new BigNumber(((((_374_v14).contains(_372_v12)) ? ((_374_v14).get(_372_v12)) : (_dafny.Seq.of(_372_v12)))).length); _375_i5.isLessThan(_hi0); _375_i5 = _375_i5.plus(_dafny.ONE)) {
        let _376_v15;
        _376_v15 = new _dafny.CodePoint('e'.codePointAt(0));
        let _377_v16;
        _377_v16 = _dafny.Map.Empty.slice().updateUnsafe(_375_i5,_376_v15);
        _377_v16 = (_377_v16).update((_dafny.ZERO).minus((new BigNumber((_373_v13).length)).multipliedBy(new BigNumber(-703))), _376_v15);
        let _378_v17;
        _378_v17 = _module.D12.create_DC23(new BigNumber((_356_v1).length));
        (globalState).f7 = (_378_v17).dtor_cf30;
        let _379_v18;
        _379_v18 = _dafny.Seq.of((p0)[_module.__default.safeIndex(new BigNumber(727), new BigNumber((p0).length))], new BigNumber(651), (p0)[_module.__default.safeIndex(new BigNumber(727), new BigNumber((p0).length))], _this.f29);
        let _380_v19;
        _380_v19 = _dafny.Map.Empty.slice().updateUnsafe(_379_v18,_369_v11);
        let _381_v20;
        _381_v20 = _dafny.Map.Empty.slice().updateUnsafe((((_380_v19).contains(_379_v18)) ? ((_380_v19).get(_379_v18)) : (_369_v11)),_dafny.Seq.of(_379_v18, _379_v18));
        let _382_v21;
        _382_v21 = _dafny.Seq.of(_379_v18);
        _381_v20 = (_381_v20).update(_369_v11, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(880)), ((_383_p0) => function (_384_i6) {
          return _dafny.Seq.of((_383_p0)[_module.__default.safeIndex(new BigNumber(727), new BigNumber((_383_p0).length))], new BigNumber(368), new BigNumber((_dafny.Seq.UnicodeFromString("jcxjkvp")).length));
        })(p0)), _382_v21));
        let _385_v22;
        _385_v22 = _dafny.MultiSet.fromElements(_379_v18, _379_v18, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-816)), ((_386_i5) => function (_387_i7) {
          return _386_i5;
        })(_375_i5)), _dafny.Seq.of(_375_i5, (_dafny.ZERO).minus(_375_i5)), _dafny.Seq.update(_379_v18, _module.__default.safeIndex(new BigNumber(938), new BigNumber((_379_v18).length)), new BigNumber(206)));
        let _388_v23;
        _388_v23 = _dafny.MultiSet.fromElements(_356_v1);
        let _rhs36 = (new BigNumber((_355_v0).length)).isEqualTo(new BigNumber(509));
        let _rhs37 = ((new BigNumber(33)).minus(_this.f29)).isLessThanOrEqualTo((((_385_v22).contains(_379_v18)) ? ((_385_v22).get(_379_v18)) : ((((_388_v23).contains(_356_v1)) ? ((_388_v23).get(_356_v1)) : ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ilycvcw")).length)))))));
        let _lhs40 = globalState;
        let _lhs41 = globalState;
        _lhs40.f21 = _rhs36;
        _lhs41.f13 = _rhs37;
      }
      return;
    }
    m1(p0, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = false;
      let _389_v0;
      _389_v0 = false;
      (globalState).f0 = ((_389_v0) ? (!(p0).isEqualTo(p0)) : (_389_v0));
      if ((_this).fm7(_389_v0, ((_389_v0) ? (_this.f29) : (p0)), (p0).isEqualTo(new BigNumber((_dafny.MultiSet.fromElements(p0)).cardinality())), globalState)) {
        let _390_v1;
        _390_v1 = _dafny.MultiSet.fromElements(_this.f29);
        let _391_v2;
        _391_v2 = _dafny.Map.Empty.slice().updateUnsafe(_390_v1,_389_v0);
        let _392_v3;
        _392_v3 = _dafny.Set.fromElements(false);
        let _393_v4;
        _393_v4 = _dafny.Seq.UnicodeFromString("by");
        let _394_v5;
        _394_v5 = _dafny.Seq.of((((_390_v1).contains(new BigNumber((_392_v3).length))) ? ((_390_v1).get(new BigNumber((_392_v3).length))) : (p0)), p0, new BigNumber((_393_v4).length));
        let _source8 = (_391_v2).update(_dafny.MultiSet.FromArray(_dafny.Seq.update(_394_v5, _module.__default.safeIndex(new BigNumber(143), new BigNumber((_394_v5).length)), _this.f29)), _389_v0);
        let _395___mcc_h0 = _source8;
        let _396_cf27 = _395___mcc_h0;
        (globalState).f13 = _389_v0;
        (globalState).f7 = p0;
        let _397_v6;
        let _nw43 = new _module.C0();
        _nw43.__ctor(_389_v0);
        _397_v6 = _nw43;
        (globalState).f21 = _dafny.areEqual(_dafny.Seq.Concat(_393_v4, _dafny.Seq.UnicodeFromString("sbarm")), _393_v4);
        let _398_v7;
        _398_v7 = _module.D13.create_DC26(_393_v4, _389_v0, _389_v0);
        let _399_v8;
        _399_v8 = _dafny.Set.fromElements(_393_v4, (_398_v7).dtor_cf33, _393_v4);
        (globalState).f10 = (_399_v8).Intersect(_dafny.Set.fromElements(_393_v4, _393_v4));
        let _400_v9;
        let _nw44 = new _module.C0();
        _nw44.__ctor(_389_v0);
        _400_v9 = _nw44;
        (globalState).f7 = _this.f29;
        let _401_v10;
        _401_v10 = new _dafny.CodePoint('l'.codePointAt(0));
        let _402_v11;
        _402_v11 = _dafny.Set.fromElements(_401_v10);
        let _403_v12;
        _403_v12 = _dafny.Set.fromElements((_402_v11).Union(_dafny.Set.fromElements(_401_v10)));
        _403_v12 = _module.__default.fm38(globalState);
      } else {
        if (false) {
          (globalState).f7 = ((_389_v0) ? (new BigNumber(366)) : (p0));
          let _404_v13;
          let _nw45 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Map.Empty);
          _404_v13 = _nw45;
          let _405_v14;
          _405_v14 = new _dafny.CodePoint('b'.codePointAt(0));
          let _406_v15;
          _406_v15 = _dafny.Map.Empty.slice().updateUnsafe(_389_v0,_405_v14);
          let _index56 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_404_v13).length));
          (_404_v13)[_index56] = _406_v15;
          let _407_v16;
          _407_v16 = _dafny.Map.Empty.slice().updateUnsafe(true,_389_v0);
          let _index57 = _module.__default.safeIndex(new BigNumber(938), new BigNumber((_404_v13).length));
          (_404_v13)[_index57] = (_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(_407_v16, globalState),_405_v14)).Merge(_406_v15);
          let _408_v17;
          _408_v17 = _module.D7.create_DC15(_dafny.Seq.Create(_module.__default.abs(new BigNumber(279)), function (_409_i0) {
  return new BigNumber((_dafny.Seq.UnicodeFromString("kueylnvh")).length);
}));
          let _410_v18;
          _410_v18 = _module.D7.create_DC17(_408_v17);
          let _411_v19;
          _411_v19 = _dafny.Seq.of(_410_v18);
          (globalState).f21 = _dafny.Seq.IsProperPrefixOf(_411_v19, _dafny.Seq.Concat(_411_v19, _411_v19));
          let _412_v20;
          let _nw46 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
          _412_v20 = _nw46;
          r0 = _412_v20;
          (globalState).f0 = _389_v0;
        } else {
          (globalState).f16 = new BigNumber(748);
          (globalState).f7 = _this.f29;
          let _413_v21;
          let _nw47 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
          _413_v21 = _nw47;
          let _index58 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_413_v21).length));
          (_413_v21)[_index58] = (_dafny.ZERO).minus(_this.f29);
          let _414_v22;
          _414_v22 = _dafny.Map.Empty.slice().updateUnsafe(_389_v0,p0);
          let _index59 = _module.__default.safeIndex(new BigNumber(810), new BigNumber((_413_v21).length));
          (_413_v21)[_index59] = _module.__default.safeDivisionInt((((_414_v22).contains(_389_v0)) ? ((_414_v22).get(_389_v0)) : (_this.f29)), p0);
          let _415_v23;
          _415_v23 = _dafny.MultiSet.fromElements(_389_v0);
          let _416_v24;
          _416_v24 = _dafny.Seq.of(_415_v23);
          let _417_v25;
          _417_v25 = (_416_v24)[_module.__default.safeIndex(new BigNumber((_414_v22).length), new BigNumber((_416_v24).length))];
          _417_v25 = _417_v25;
          (globalState).f13 = (_389_v0) || (_389_v0);
        }
        let _418_v26;
        let _nw48 = new _module.C0();
        _nw48.__ctor(_389_v0);
        _418_v26 = _nw48;
        (globalState).f7 = p0;
        let _419_v27;
        _419_v27 = _dafny.Map.Empty.slice().updateUnsafe(_418_v26.f27,(_this).fm29(globalState));
        let _420_v28;
        _420_v28 = _dafny.MultiSet.fromElements(_419_v27);
        let _421_v29;
        _421_v29 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_this.f29)).cardinality()),_418_v26.f27);
        let _422_v30;
        _422_v30 = _dafny.Seq.of(new BigNumber((_421_v29).length), _this.f29);
        (globalState).f2 = (_420_v28).IsSubsetOf((_420_v28).Difference(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.MultiSet.FromArray(_422_v30)).cardinality())), _419_v27, _419_v27, _419_v27, _419_v27)));
        (globalState).f7 = p0;
      }
      let _423_v31;
      let _nw49 = new _module.C0();
      _nw49.__ctor(_389_v0);
      _423_v31 = _nw49;
      let _424_i1;
      _424_i1 = _dafny.ZERO;
      L3: {
        while ((_this.f29).isEqualTo((_this).fm29(globalState))) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_424_i1)) {
              break L3;
            }
            _424_i1 = (_424_i1).plus(_dafny.ONE);
            (globalState).f13 = !(_389_v0);
            let _425_v32;
            let _init8 = ((_426_v0, _427_v31) => function (_428_i2) {
              return _module.__default.safeModuloInt(_428_i2, new BigNumber((_dafny.MultiSet.fromElements(_426_v0, !(_427_v31.f27))).cardinality()));
            })(_389_v0, _423_v31);
            let _nw50 = Array((new BigNumber(2)).toNumber());
            for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw50.length); _i0_8++) {
              _nw50[_i0_8] = _init8(new BigNumber(_i0_8));
            }
            _425_v32 = _nw50;
            let _429_v33;
            _429_v33 = _dafny.Set.fromElements(_425_v32, _425_v32);
            let _430_v34;
            _430_v34 = _dafny.Seq.of(_389_v0, _389_v0, _423_v31.f27);
            let _431_v35;
            _431_v35 = _dafny.Seq.UnicodeFromString("xjuxav");
            let _432_v36;
            let _nw51 = Array((new BigNumber(11)).toNumber());
            _nw51[(_dafny.ZERO).toNumber()] = new BigNumber((_module.__default.fm1(new BigNumber((_429_v33).length), globalState)).length);
            _nw51[(_dafny.ONE).toNumber()] = p0;
            _nw51[(new BigNumber(2)).toNumber()] = new BigNumber((_430_v34).length);
            _nw51[(new BigNumber(3)).toNumber()] = (new BigNumber((_431_v35).length)).plus(new BigNumber(-830));
            _nw51[(new BigNumber(4)).toNumber()] = p0;
            _nw51[(new BigNumber(5)).toNumber()] = (new BigNumber((_431_v35).length)).minus(new BigNumber((_431_v35).length));
            _nw51[(new BigNumber(6)).toNumber()] = p0;
            _nw51[(new BigNumber(7)).toNumber()] = _this.f29;
            _nw51[(new BigNumber(8)).toNumber()] = _this.f29;
            _nw51[(new BigNumber(9)).toNumber()] = _module.__default.fm0(new BigNumber(-110), p0, _this.f29, _389_v0, globalState);
            _nw51[(new BigNumber(10)).toNumber()] = new BigNumber(-680);
            _432_v36 = _nw51;
            let _index60 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_425_v32).length));
            (_425_v32)[_index60] = _this.f29;
            let _index61 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((_425_v32).length));
            (_425_v32)[_index61] = _this.f29;
            let _433_v37;
            _433_v37 = _dafny.Map.Empty.slice().updateUnsafe((_425_v32)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((_425_v32).length))],_423_v31.f27);
            let _434_v38;
            _434_v38 = _module.D1.create_DC2((_433_v37).Merge(_433_v37));
            _434_v38 = _module.D1.create_DC2(_433_v37);
            let _435_v39;
            _435_v39 = _dafny.Map.Empty.slice().updateUnsafe(_423_v31.f27,_this.f29);
            _435_v39 = ((_435_v39).Merge(_435_v39)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_389_v0,new BigNumber(811)));
          }
        }
      }
      let _436_v40;
      let _nw52 = new _module.C0();
      _nw52.__ctor(_423_v31.f27);
      _436_v40 = _nw52;
      let _437_v41;
      _437_v41 = _dafny.Map.Empty.slice().updateUnsafe(!(_389_v0),_436_v40.f27);
      let _rhs38 = (_this).fm29(globalState);
      let _rhs39 = _437_v41;
      let _lhs42 = _this;
      _lhs42.f29 = _rhs38;
      _437_v41 = _rhs39;
      let _438_v42;
      let _nw53 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
      _438_v42 = _nw53;
      r0 = _438_v42;
      r1 = _436_v40.f27;
      return [r0, r1];
    }
    get f30() {
      let _this = this;
      return _this._f30;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f28 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    __ctor(f28) {
      let _this = this;
      (_this)._f28 = f28;
      return;
    }
    fm7(p0, p1, p2, globalState) {
      let _this = this;
      return (((_dafny.ZERO).minus((_this).f28)).isLessThan((_this).f28)) === (((!(!(true))) ? (false) : (false)));
    };
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _source9 = _module.D1.create_DC2(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-844),false));
      if (_source9.is_DC3) {
        let _439___mcc_h0 = (_source9).cf3;
        let _440_cf3 = _439___mcc_h0;
        return _dafny.Map.Empty.slice().updateUnsafe(_440_cf3,_dafny.Seq.UnicodeFromString("lcwihbpu"));
      } else {
        let _441___mcc_h1 = (_source9).cf2;
        let _442_cf2 = _441___mcc_h1;
        return (_dafny.Map.Empty.slice().updateUnsafe(!(false),_dafny.Seq.UnicodeFromString("duyckw"))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(true),_dafny.Seq.UnicodeFromString("mtfgtk")));
      }
    };
    fm6(p0, globalState) {
      let _this = this;
      return false;
    };
    fm17(p0, p1, p2, p3, globalState) {
      let _this = this;
      if (!(true) || (false)) {
        return (_this).f28;
      } else {
        return (_this).f28;
      }
    };
    m2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = [];
      let _443_v0;
      _443_v0 = _dafny.Seq.UnicodeFromString("klcc");
      let _hi1 = new BigNumber((_443_v0).length);
      for (let _444_i0 = (_this).f28; _444_i0.isLessThan(_hi1); _444_i0 = _444_i0.plus(_dafny.ONE)) {
        (globalState).f21 = p1;
        let _445_v1;
        let _nw54 = new _module.C0();
        _nw54.__ctor(p1);
        _445_v1 = _nw54;
        let _446_v2;
        _446_v2 = _dafny.MultiSet.fromElements(_445_v1.f27);
        let _447_v3;
        _447_v3 = _dafny.MultiSet.fromElements(_module.__default.safeDivisionInt(new BigNumber((_446_v2).cardinality()), (_this).f28), _444_i0);
        let _448_v4;
        _448_v4 = _module.D4.create_DC9((_dafny.ZERO).minus(p0));
        let _rhs40 = _445_v1.f27;
        let _rhs41 = (((_447_v3).contains((_this).f28)) ? ((_447_v3).get((_this).f28)) : ((_dafny.ZERO).minus((((_447_v3).contains(p0)) ? ((_447_v3).get(p0)) : (p3)))));
        let _rhs42 = ((_448_v4).dtor_cf11).plus(p3);
        let _rhs43 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(p0, _444_i0));
        let _lhs43 = _445_v1;
        let _lhs44 = globalState;
        let _lhs45 = globalState;
        let _lhs46 = globalState;
        _lhs43.f27 = _rhs40;
        _lhs44.f7 = _rhs41;
        _lhs45.f7 = _rhs42;
        _lhs46.f7 = _rhs43;
        let _449_v5;
        _449_v5 = _dafny.Seq.of(p1);
        let _450_v6;
        _450_v6 = _449_v5;
        let _451_v7;
        _451_v7 = new _dafny.CodePoint('q'.codePointAt(0));
        let _452_v8;
        let _nw55 = Array((new BigNumber(19)).toNumber());
        _nw55[(_dafny.ZERO).toNumber()] = _449_v5;
        _nw55[(_dafny.ONE).toNumber()] = _449_v5;
        _nw55[(new BigNumber(2)).toNumber()] = _dafny.Seq.update(_449_v5, _module.__default.safeIndex(new BigNumber(973), new BigNumber((_449_v5).length)), _445_v1.f27);
        _nw55[(new BigNumber(3)).toNumber()] = _449_v5;
        _nw55[(new BigNumber(4)).toNumber()] = (_450_v6);
        _nw55[(new BigNumber(5)).toNumber()] = _dafny.Seq.of(_445_v1.f27);
        _nw55[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_449_v5, _dafny.Seq.update(_module.__default.fm18(_445_v1.f27, new _dafny.CodePoint('p'.codePointAt(0)), _445_v1.f27, globalState), _module.__default.safeIndex(_444_i0, new BigNumber((_module.__default.fm18(_445_v1.f27, new _dafny.CodePoint('p'.codePointAt(0)), _445_v1.f27, globalState)).length)), _445_v1.f27));
        _nw55[(new BigNumber(7)).toNumber()] = _449_v5;
        _nw55[(new BigNumber(8)).toNumber()] = _449_v5;
        _nw55[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_449_v5, _449_v5);
        _nw55[(new BigNumber(10)).toNumber()] = _449_v5;
        _nw55[(new BigNumber(11)).toNumber()] = _dafny.Seq.of(_445_v1.f27);
        _nw55[(new BigNumber(12)).toNumber()] = _449_v5;
        _nw55[(new BigNumber(13)).toNumber()] = _dafny.Seq.of(p1, _445_v1.f27, p1, _445_v1.f27);
        _nw55[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm18(!(_445_v1.f27), _451_v7, _445_v1.f27, globalState), _449_v5);
        _nw55[(new BigNumber(15)).toNumber()] = _dafny.Seq.update(_449_v5, _module.__default.safeIndex(p0, new BigNumber((_449_v5).length)), _445_v1.f27);
        _nw55[(new BigNumber(16)).toNumber()] = _dafny.Seq.of(_445_v1.f27);
        _nw55[(new BigNumber(17)).toNumber()] = _449_v5;
        _nw55[(new BigNumber(18)).toNumber()] = _449_v5;
        _452_v8 = _nw55;
        let _index62 = _module.__default.safeIndex(new BigNumber(977), new BigNumber((_452_v8).length));
        (_452_v8)[_index62] = _449_v5;
        let _index63 = _module.__default.safeIndex(new BigNumber(977), new BigNumber((_452_v8).length));
        (_452_v8)[_index63] = _449_v5;
      }
      let _453_v9;
      let _nw56 = Array((new BigNumber(20)).toNumber()).fill(_module.D3.Default());
      _453_v9 = _nw56;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_453_v9).length))) {
        let _454_i1 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_454_i1)) && ((_454_i1).isLessThan(new BigNumber((_453_v9).length))))) {
          (_453_v9)[(_454_i1)] = _module.D3.create_DC6(_443_v0, new _dafny.CodePoint('v'.codePointAt(0)), p1);
        }
      }
      let _455_v10;
      _455_v10 = _module.D1.create_DC3(p1);
      let _456_v11;
      _456_v11 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_455_v10).dtor_cf3);
      let _457_v12;
      _457_v12 = _module.D4.create_DC7(new BigNumber((_456_v11).length));
      let _458_v13;
      _458_v13 = _dafny.Set.fromElements(false);
      let _459_v14;
      _459_v14 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-94)), function (_460_i2) {
        return new _dafny.CodePoint('w'.codePointAt(0));
      }),_458_v13);
      let _461_v15;
      _461_v15 = _dafny.Seq.of(p0, p0);
      let _462_v16;
      _462_v16 = _dafny.Seq.of(p1, p1, p1);
      let _463_v17;
      let _nw57 = Array((new BigNumber(13)).toNumber());
      _nw57[(_dafny.ZERO).toNumber()] = (((_459_v14).contains(_443_v0)) ? ((_459_v14).get(_443_v0)) : (_458_v13));
      _nw57[(_dafny.ONE).toNumber()] = (_dafny.Set.fromElements(p1)).Intersect(_458_v13);
      _nw57[(new BigNumber(2)).toNumber()] = (_458_v13).Intersect(_458_v13);
      _nw57[(new BigNumber(3)).toNumber()] = _458_v13;
      _nw57[(new BigNumber(4)).toNumber()] = _458_v13;
      _nw57[(new BigNumber(5)).toNumber()] = _458_v13;
      _nw57[(new BigNumber(6)).toNumber()] = _458_v13;
      _nw57[(new BigNumber(7)).toNumber()] = _458_v13;
      _nw57[(new BigNumber(8)).toNumber()] = _module.__default.fm19(new BigNumber((_461_v15).length), (_462_v16)[_module.__default.safeIndex(new BigNumber(702), new BigNumber((_462_v16).length))], p1, new BigNumber(-900), globalState);
      _nw57[(new BigNumber(9)).toNumber()] = ((_458_v13)).Union(_dafny.Set.fromElements(p1));
      _nw57[(new BigNumber(10)).toNumber()] = _458_v13;
      _nw57[(new BigNumber(11)).toNumber()] = (_458_v13).Union(_458_v13);
      _nw57[(new BigNumber(12)).toNumber()] = (_458_v13).Intersect(_dafny.Set.fromElements(p1, p1, p1, p1, p1));
      _463_v17 = _nw57;
      let _index64 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_463_v17).length));
      (_463_v17)[_index64] = _458_v13;
      let _464_v18;
      let _nw58 = new _module.C0();
      _nw58.__ctor((p3).isLessThan((_this).f28));
      _464_v18 = _nw58;
      let _index65 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_463_v17).length));
      let _rhs44 = (((_this).f28).multipliedBy(p3)).multipliedBy(_module.__default.safeDivisionInt(p3, new BigNumber((_461_v15).length)));
      let _rhs45 = p1;
      let _rhs46 = _457_v12;
      let _rhs47 = _458_v13;
      let _rhs48 = _464_v18;
      let _lhs47 = globalState;
      let _lhs48 = globalState;
      let _lhs49 = _463_v17;
      let _lhs50 = _module.__default.safeIndex(new BigNumber(693), new BigNumber((_463_v17).length));
      _lhs47.f7 = _rhs44;
      _lhs48.f21 = _rhs45;
      _457_v12 = _rhs46;
      _lhs49[_lhs50] = _rhs47;
      _464_v18 = _rhs48;
      (_464_v18).f27 = _464_v18.f27;
      let _465_v19;
      let _init9 = ((_466_v18) => function (_467_i3) {
        return ((_466_v18.f27) ? (_466_v18.f27) : (_466_v18.f27));
      })(_464_v18);
      let _nw59 = Array((new BigNumber(17)).toNumber());
      for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw59.length); _i0_9++) {
        _nw59[_i0_9] = _init9(new BigNumber(_i0_9));
      }
      _465_v19 = _nw59;
      let _index66 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_465_v19).length));
      (_465_v19)[_index66] = p1;
      let _468_v20;
      _468_v20 = new _dafny.CodePoint('a'.codePointAt(0));
      let _469_v21;
      _469_v21 = _module.__default.fm18(_464_v18.f27, _468_v20, false, globalState);
      let _index67 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_465_v19).length));
      let _rhs49 = (_this).fm17(_464_v18.f27, _dafny.Seq.Create(_module.__default.abs(new BigNumber(518)), function (_470_i4) {
        return (_this).f28;
      }), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(191)), function (_471_i5) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      })).length), _469_v21, globalState);
      let _rhs50 = (_464_v18.f27) || (_464_v18.f27);
      let _lhs51 = globalState;
      let _lhs52 = _465_v19;
      let _lhs53 = _module.__default.safeIndex(new BigNumber(337), new BigNumber((_465_v19).length));
      _lhs51.f7 = _rhs49;
      _lhs52[_lhs53] = _rhs50;
      let _472_v22;
      let _init10 = ((_473_v10) => function (_474_i7) {
        return _473_v10;
      })(_455_v10);
      let _nw60 = Array((new BigNumber(2)).toNumber());
      for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw60.length); _i0_10++) {
        _nw60[_i0_10] = _init10(new BigNumber(_i0_10));
      }
      _472_v22 = _nw60;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_472_v22).length))) {
        let _475_i6 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_475_i6)) && ((_475_i6).isLessThan(new BigNumber((_472_v22).length))))) {
          (_472_v22)[(_475_i6)] = (((((_456_v11).contains(new BigNumber((_dafny.Set.fromElements(p0, (_this).f28)).length))) ? ((_456_v11).get(new BigNumber((_dafny.Set.fromElements(p0, (_this).f28)).length))) : (_464_v18.f27))) ? (_455_v10) : (_455_v10));
        }
      }
      let _476_v23;
      _476_v23 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
      r0 = !(_module.__default.fm2((_476_v23).update(_464_v18.f27, _464_v18.f27), globalState));
      let _pat_let_tv6 = p3;
      let _pat_let_tv7 = _465_v19;
      let _pat_let_tv8 = _465_v19;
      let _pat_let_tv9 = p3;
      let _pat_let_tv10 = p3;
      r1 = function (_source10) {
        if (_source10.is_DC8) {
          let _477___mcc_h0 = (_source10).cf10;
          let _478_cf10 = _477___mcc_h0;
          return _pat_let_tv6;
        } else if (_source10.is_DC9) {
          let _479___mcc_h1 = (_source10).cf11;
          let _480_cf11 = _479___mcc_h1;
          return new BigNumber(((_module.D12.create_DC22(_dafny.Map.Empty.slice().updateUnsafe((_pat_let_tv8)[_module.__default.safeIndex(new BigNumber(337), new BigNumber((_pat_let_tv7).length))],_pat_let_tv9))).dtor_cf29).length);
        } else if (_source10.is_DC10) {
          let _481___mcc_h2 = (_source10).cf12;
          let _482___mcc_h3 = (_source10).cf13;
          let _483___mcc_h4 = (_source10).cf14;
          let _484___mcc_h5 = (_source10).cf15;
          let _485___mcc_h6 = (_source10).cf16;
          let _486_cf16 = _485___mcc_h6;
          let _487_cf15 = _484___mcc_h5;
          let _488_cf14 = _483___mcc_h4;
          let _489_cf13 = _482___mcc_h3;
          let _490_cf12 = _481___mcc_h2;
          return (_pat_let_tv10).plus(_489_cf13);
        } else {
          let _491___mcc_h7 = (_source10).cf9;
          let _492_cf9 = _491___mcc_h7;
          return (_this).f28;
        }
      }(_module.D4.create_DC7(new BigNumber(406)));
      let _493_v24;
      let _nw61 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Map.Empty);
      _493_v24 = _nw61;
      r2 = _493_v24;
      return [r0, r1, r2];
    }
    m3(p0, globalState) {
      let _this = this;
      (globalState).f16 = new BigNumber(644);
      let _494_v0;
      _494_v0 = _module.D12.create_DC23(new BigNumber(7));
      let _source11 = function (_pat_let6_0) {
        return function (_495_dt__update__tmp_h0) {
          return function (_pat_let7_0) {
            return function (_496_dt__update_hcf30_h0) {
              return _module.D12.create_DC23(_496_dt__update_hcf30_h0);
            }(_pat_let7_0);
          }(((_this).f28).multipliedBy((_this).f28));
        }(_pat_let6_0);
      }(_494_v0);
      if (_source11.is_DC23) {
        let _497___mcc_h0 = (_source11).cf30;
        let _498_cf30 = _497___mcc_h0;
        let _499_v1;
        _499_v1 = true;
        let _500_v2;
        _500_v2 = new _dafny.CodePoint('n'.codePointAt(0));
        let _501_v3;
        _501_v3 = _dafny.Map.Empty.slice().updateUnsafe(!(_499_v1),_500_v2);
        let _502_v4;
        _502_v4 = _dafny.Map.Empty.slice().updateUnsafe(_498_cf30,new BigNumber((_501_v3).length));
        let _503_v5;
        _503_v5 = _dafny.Map.Empty.slice().updateUnsafe(_499_v1,new BigNumber((_502_v4).length));
        if (!((_503_v5).update(_module.__default.fm2(_module.__default.fm20(_498_cf30, _499_v1, _499_v1, _499_v1, globalState), globalState), new BigNumber((_dafny.Seq.UnicodeFromString("bkpcvyger")).length))).equals(_503_v5)) {
          let _504_v6;
          _504_v6 = _dafny.MultiSet.fromElements(_499_v1);
          _499_v1 = ((_504_v6).Difference(_dafny.MultiSet.fromElements(_499_v1))).IsProperSubsetOf(_504_v6);
          let _505_v7;
          let _nw62 = Array((new BigNumber(24)).toNumber()).fill([]);
          _505_v7 = _nw62;
          let _index68 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_505_v7).length));
          (_505_v7)[_index68] = p0;
          let _index69 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_505_v7).length));
          (_505_v7)[_index69] = p0;
          let _506_v8;
          let _nw63 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Set.Empty);
          _506_v8 = _nw63;
          let _507_v9;
          _507_v9 = _dafny.Map.Empty.slice().updateUnsafe(_499_v1,_499_v1);
          let _508_v10;
          _508_v10 = _dafny.Set.fromElements(_507_v9);
          let _index70 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((_506_v8).length));
          (_506_v8)[_index70] = _508_v10;
          let _index71 = _module.__default.safeIndex(new BigNumber(135), new BigNumber((_506_v8).length));
          (_506_v8)[_index71] = _508_v10;
          let _509_v11;
          _509_v11 = _dafny.Set.fromElements(_499_v1);
          let _510_v12;
          _510_v12 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_509_v11).length),_499_v1);
          (globalState).f2 = !(_510_v12).contains(_498_cf30);
          let _511_v13;
          _511_v13 = _dafny.Seq.UnicodeFromString("aaacch");
          let _512_v14;
          _512_v14 = _dafny.Seq.of(_dafny.Seq.Concat(_511_v13, _dafny.Seq.Create(_module.__default.abs(new BigNumber(966)), ((_513_v2) => function (_514_i0) {
            return _513_v2;
          })(_500_v2))), _511_v13, _511_v13);
          let _index72 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_505_v7).length));
          let _rhs51 = _499_v1;
          let _rhs52 = !(_499_v1);
          let _rhs53 = (_505_v7)[_module.__default.safeIndex(new BigNumber(115), new BigNumber((_505_v7).length))];
          let _rhs54 = (_512_v14)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f28), new BigNumber((_512_v14).length))];
          let _rhs55 = _498_cf30;
          let _lhs54 = globalState;
          let _lhs55 = globalState;
          let _lhs56 = _505_v7;
          let _lhs57 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_505_v7).length));
          let _lhs58 = globalState;
          let _lhs59 = globalState;
          _lhs54.f0 = _rhs51;
          _lhs55.f0 = _rhs52;
          _lhs56[_lhs57] = _rhs53;
          _lhs58.f19 = _rhs54;
          _lhs59.f16 = _rhs55;
        } else {
          _500_v2 = new _dafny.CodePoint('a'.codePointAt(0));
          (_this).m8(!(!(!(_499_v1))), _499_v1, _499_v1, globalState);
          let _515_v15;
          _515_v15 = _dafny.Seq.of(_499_v1);
          let _rhs56 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt(_498_cf30, (_this).f28))));
          let _rhs57 = new BigNumber((_515_v15).length);
          let _lhs60 = globalState;
          let _lhs61 = globalState;
          _lhs60.f7 = _rhs56;
          _lhs61.f7 = _rhs57;
          let _516_v16;
          _516_v16 = _dafny.Seq.UnicodeFromString("olkogre");
          let _517_v17;
          _517_v17 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_516_v16, _module.__default.safeIndex(_498_cf30, new BigNumber((_516_v16).length)), _500_v2),((false) ? (!(_499_v1)) : (false)));
          let _rhs58 = _499_v1;
          let _rhs59 = _517_v17;
          let _rhs60 = _499_v1;
          let _lhs62 = globalState;
          let _lhs63 = globalState;
          _lhs62.f0 = _rhs58;
          _517_v17 = _rhs59;
          _lhs63.f2 = _rhs60;
          (globalState).f16 = (_this).f28;
        }
        let _index73 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((p0).length));
        (p0)[_index73] = _498_cf30;
        let _index74 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((p0).length));
        let _rhs61 = (_this).f28;
        let _rhs62 = (new BigNumber(-20)).minus(_498_cf30);
        let _rhs63 = _500_v2;
        let _lhs64 = p0;
        let _lhs65 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((p0).length));
        let _lhs66 = globalState;
        _lhs64[_lhs65] = _rhs61;
        _lhs66.f7 = _rhs62;
        _500_v2 = _rhs63;
        let _518_v18;
        let _nw64 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
        _518_v18 = _nw64;
        let _519_v19;
        _519_v19 = _dafny.MultiSet.fromElements(_518_v18);
        let _520_v20;
        _520_v20 = _519_v19;
        let _521_v21;
        _521_v21 = (_520_v20);
        let _source12 = _520_v20;
        let _522___mcc_h3 = _source12;
        let _523_cf4 = _522___mcc_h3;
        let _524_v22;
        let _nw65 = new _module.C0();
        _nw65.__ctor(!(_499_v1));
        _524_v22 = _nw65;
        (globalState).f16 = _498_cf30;
        (globalState).f16 = (_this).f28;
        let _525_v23;
        _525_v23 = _dafny.Seq.UnicodeFromString("am");
        let _526_v24;
        _526_v24 = _dafny.Set.fromElements(_module.__default.safeDivisionInt(new BigNumber(-250), (_this).f28), (p0)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((p0).length))], (_this).f28, new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_525_v23, _525_v23), _module.__default.safeIndex(_498_cf30, new BigNumber((_dafny.Seq.Concat(_525_v23, _525_v23)).length)), _500_v2)).length));
        _526_v24 = ((_526_v24).Difference(_526_v24)).Difference(_526_v24);
        let _527_v25;
        _527_v25 = _dafny.Set.fromElements(_498_cf30);
        let _528_v26;
        _528_v26 = _dafny.Seq.of(_499_v1, true);
        let _529_v27;
        _529_v27 = _dafny.Seq.of(new BigNumber((_528_v26).length), _498_cf30);
        let _530_v28;
        _530_v28 = _dafny.MultiSet.fromElements(_498_cf30);
        let _531_v29;
        _531_v29 = _dafny.MultiSet.fromElements((p0)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((p0).length))], (_529_v27)[_module.__default.safeIndex(new BigNumber((_530_v28).cardinality()), new BigNumber((_529_v27).length))]);
        _527_v25 = ((_527_v25).Intersect(_dafny.Set.fromElements(_498_cf30, (_this).f28, _498_cf30, new BigNumber((_530_v28).cardinality())))).Union(_527_v25);
      } else if (_source11.is_DC22) {
        let _532___mcc_h1 = (_source11).cf29;
        let _533_cf29 = _532___mcc_h1;
        let _534_v30;
        let _init11 = function (_535_i1) {
          return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("liqnc"), _dafny.Seq.UnicodeFromString("mjrvidhd"));
        };
        let _nw66 = Array((new BigNumber(17)).toNumber());
        for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw66.length); _i0_11++) {
          _nw66[_i0_11] = _init11(new BigNumber(_i0_11));
        }
        _534_v30 = _nw66;
        _534_v30 = _534_v30;
        let _536_v31;
        _536_v31 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f28),(_this).f28);
        let _537_v32;
        _537_v32 = true;
        let _538_v34;
        _538_v34 = _module.D13.create_DC25(function () {
  let _coll30 = new _dafny.Map();
  for (const _compr_30 of (_module.__default.fm21(globalState)).Elements) {
    let _539_v33 = _compr_30;
    if (_dafny.Seq.contains(_module.__default.fm21(globalState), _539_v33)) {
      _coll30.push([_module.__default.safeModuloInt(_539_v33, (_this).f28),(_this).f28]);
    }
  }
  return _coll30;
}());
        let _540_v35;
        _540_v35 = _dafny.Map.Empty.slice().updateUnsafe((false) && (_537_v32),(_538_v34).dtor_cf32);
        _536_v31 = (((_540_v35).contains(_537_v32)) ? ((_540_v35).get(_537_v32)) : (_536_v31));
        let _541_v36;
        let _nw67 = Array((new BigNumber(3)).toNumber()).fill(_dafny.MultiSet.Empty);
        _541_v36 = _nw67;
        let _542_v37;
        _542_v37 = _dafny.MultiSet.fromElements(p0, p0, p0);
        let _543_v38;
        _543_v38 = _542_v37;
        let _index75 = _module.__default.safeIndex(new BigNumber(962), new BigNumber((_541_v36).length));
        (_541_v36)[_index75] = _543_v38;
        let _index76 = _module.__default.safeIndex(new BigNumber(962), new BigNumber((_541_v36).length));
        (_541_v36)[_index76] = _543_v38;
        if (true) {
          let _544_v39;
          _544_v39 = _dafny.Seq.of(new BigNumber(-367));
          let _545_v40;
          _545_v40 = _dafny.Seq.of(_537_v32, _537_v32);
          let _546_v41;
          _546_v41 = _545_v40;
          (globalState).f7 = ((_537_v32) ? ((_this).f28) : (_module.__default.fm0((_this).f28, (_this).fm17(_537_v32, _544_v39, (_this).f28, _546_v41, globalState), (_this).f28, _537_v32, globalState)));
          let _547_v42;
          _547_v42 = new _dafny.CodePoint('n'.codePointAt(0));
          let _548_v43;
          _548_v43 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f28),_547_v42);
          let _549_v44;
          _549_v44 = _dafny.Seq.UnicodeFromString("mg");
          _548_v43 = (_548_v43).update(new BigNumber(366), (_549_v44)[_module.__default.safeIndex((_this).f28, new BigNumber((_549_v44).length))]);
          let _550_v45;
          let _nw68 = new _module.C0();
          _nw68.__ctor(_537_v32);
          _550_v45 = _nw68;
          _545_v40 = _545_v40;
          let _551_v46;
          _551_v46 = _dafny.Map.Empty.slice().updateUnsafe(p0,_537_v32);
          _551_v46 = (_dafny.Map.Empty.slice().updateUnsafe(p0,_537_v32)).Merge(_551_v46);
        } else {
          let _552_v47;
          let _nw69 = new _module.C0();
          _nw69.__ctor(_537_v32);
          _552_v47 = _nw69;
          let _553_v48;
          _553_v48 = _dafny.MultiSet.fromElements(_552_v47);
          (globalState).f7 = (((_this).f28).multipliedBy((_this).f28)).minus(_module.__default.safeDivisionInt((((_553_v48).contains(_552_v47)) ? ((_553_v48).get(_552_v47)) : ((_this).f28)), (_this).f28));
          let _554_v49;
          _554_v49 = _dafny.Seq.of(false);
          let _555_v50;
          _555_v50 = _dafny.Seq.UnicodeFromString("v");
          let _556_v51;
          _556_v51 = new _dafny.CodePoint('s'.codePointAt(0));
          let _557_v52;
          _557_v52 = _module.D3.create_DC6(_555_v50, _556_v51, _552_v47.f27);
          let _558_v53;
          _558_v53 = _dafny.Map.Empty.slice().updateUnsafe(_554_v49,(_557_v52).dtor_cf6);
          _module.__default.m0(_552_v47.f27, (_this).f28, ((_552_v47.f27) ? (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_554_v49, _module.__default.safeIndex((_this).f28, new BigNumber((_554_v49).length)), _552_v47.f27),_555_v50)) : (_558_v53)), (_this).f28, globalState);
          (globalState).f16 = (_this).f28;
          let _559_v54;
          let _nw70 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Map.Empty);
          _559_v54 = _nw70;
          _559_v54 = _559_v54;
          (globalState).f21 = false;
        }
      } else {
        let _560___mcc_h2 = (_source11).cf31;
        let _561_cf31 = _560___mcc_h2;
        let _562_v55;
        _562_v55 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f28);
        let _563_v56;
        _563_v56 = false;
        _562_v55 = (_562_v55).update(_563_v56, ((_563_v56) ? ((_this).f28) : ((_this).f28)));
        (_this).m8(_563_v56, _563_v56, _563_v56, globalState);
        let _564_v57;
        _564_v57 = _dafny.Seq.UnicodeFromString("npq");
        (globalState).f13 = _dafny.Seq.IsProperPrefixOf(_564_v57, _564_v57);
        let _565_v58;
        _565_v58 = new _dafny.CodePoint('d'.codePointAt(0));
        let _566_v59;
        _566_v59 = _dafny.Seq.of(_563_v56);
        let _567_v60;
        _567_v60 = _module.D13.create_DC26(_564_v57, (_module.D3.create_DC6(_dafny.Seq.update(_564_v57, _module.__default.safeIndex((_this).f28, new BigNumber((_564_v57).length)), new _dafny.CodePoint('e'.codePointAt(0))), _565_v58, _563_v56)).dtor_cf8, (_566_v59)[_module.__default.safeIndex((_this).f28, new BigNumber((_566_v59).length))]);
        if ((_567_v60).dtor_cf35) {
          let _568_v61;
          _568_v61 = _dafny.Seq.of((_this).f28);
          let _569_v62;
          _569_v62 = _dafny.Seq.of(_568_v61, _dafny.Seq.Create(_module.__default.abs(new BigNumber(886)), function (_570_i2) {
            return (_this).f28;
          }));
          let _571_v63;
          _571_v63 = _dafny.Seq.of(_569_v62, _569_v62);
          let _572_v64;
          let _nw71 = Array((new BigNumber(7)).toNumber());
          _nw71[(_dafny.ZERO).toNumber()] = (_571_v63)[_module.__default.safeIndex((_this).f28, new BigNumber((_571_v63).length))];
          _nw71[(_dafny.ONE).toNumber()] = _569_v62;
          _nw71[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-179)), ((_573_v58) => function (_574_i3) {
            return _dafny.Seq.update(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-177)), ((_575_v58) => function (_576_i4) {
              return _575_v58;
            })(_573_v58))).length)), _module.__default.safeIndex((_this).f28, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-177)), ((_577_v58) => function (_578_i4) {
              return _577_v58;
            })(_573_v58))).length))).length)), (_this).f28);
          })(_565_v58));
          _nw71[(new BigNumber(3)).toNumber()] = _569_v62;
          _nw71[(new BigNumber(4)).toNumber()] = _569_v62;
          _nw71[(new BigNumber(5)).toNumber()] = _569_v62;
          _nw71[(new BigNumber(6)).toNumber()] = ((_563_v56) ? (_569_v62) : (_dafny.Seq.update(_module.__default.fm22(_563_v56, _563_v56, (_566_v59)[_module.__default.safeIndex((_this).f28, new BigNumber((_566_v59).length))], globalState), _module.__default.safeIndex((_this).f28, new BigNumber((_module.__default.fm22(_563_v56, _563_v56, (_566_v59)[_module.__default.safeIndex((_this).f28, new BigNumber((_566_v59).length))], globalState)).length)), _568_v61)));
          _572_v64 = _nw71;
          let _579_v65;
          _579_v65 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_566_v59).length)),_569_v62);
          let _index77 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_572_v64).length));
          (_572_v64)[_index77] = _dafny.Seq.Concat(_569_v62, (((_579_v65).contains((_this).f28)) ? ((_579_v65).get((_this).f28)) : (_569_v62)));
          let _index78 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_572_v64).length));
          (_572_v64)[_index78] = _569_v62;
          (globalState).f13 = _563_v56;
          let _index79 = _module.__default.safeIndex(new BigNumber(623), new BigNumber((p0).length));
          (p0)[_index79] = _module.__default.safeModuloInt(new BigNumber(976), (_this).f28);
          let _index80 = _module.__default.safeIndex(new BigNumber(623), new BigNumber((p0).length));
          (p0)[_index80] = (_dafny.ZERO).minus((_this).f28);
          let _580_v66;
          _580_v66 = _dafny.MultiSet.fromElements(_565_v58, (_564_v57)[_module.__default.safeIndex(new BigNumber(313), new BigNumber((_564_v57).length))], ((_563_v56) ? (_565_v58) : (_565_v58)), _565_v58);
          (globalState).f16 = (((_580_v66).contains((_564_v57)[_module.__default.safeIndex((_this).f28, new BigNumber((_564_v57).length))])) ? ((_580_v66).get((_564_v57)[_module.__default.safeIndex((_this).f28, new BigNumber((_564_v57).length))])) : (_module.__default.safeDivisionInt((p0)[_module.__default.safeIndex(new BigNumber(623), new BigNumber((p0).length))], (_this).f28)));
          (globalState).f21 = !(!((_this).f28).isEqualTo((_this).f28));
        } else {
          (globalState).f2 = _563_v56;
          let _581_v67;
          _581_v67 = _dafny.Set.fromElements(_565_v58, _565_v58);
          let _582_v68;
          _582_v68 = _dafny.Seq.of(_581_v67);
          (globalState).f2 = ((_582_v68)[_module.__default.safeIndex(new BigNumber(924), new BigNumber((_582_v68).length))]).contains(new _dafny.CodePoint('w'.codePointAt(0)));
          let _583_v69;
          _583_v69 = _dafny.Seq.of((_this).f28);
          let _584_v70;
          _584_v70 = _dafny.Set.fromElements(_563_v56);
          let _585_v71;
          _585_v71 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,_584_v70);
          let _586_v72;
          let _nw72 = Array((new BigNumber(9)).toNumber());
          _nw72[(_dafny.ZERO).toNumber()] = _584_v70;
          _nw72[(_dafny.ONE).toNumber()] = _dafny.Set.fromElements(_563_v56, _563_v56);
          _nw72[(new BigNumber(2)).toNumber()] = _dafny.Set.fromElements(_563_v56);
          _nw72[(new BigNumber(3)).toNumber()] = (((_585_v71).contains((_this).f28)) ? ((_585_v71).get((_this).f28)) : (_584_v70));
          _nw72[(new BigNumber(4)).toNumber()] = _584_v70;
          _nw72[(new BigNumber(5)).toNumber()] = _584_v70;
          _nw72[(new BigNumber(6)).toNumber()] = _584_v70;
          _nw72[(new BigNumber(7)).toNumber()] = _584_v70;
          _nw72[(new BigNumber(8)).toNumber()] = _584_v70;
          _586_v72 = _nw72;
          let _587_v73;
          _587_v73 = _module.D7.create_DC16(_586_v72);
          let _588_v74;
          let _nw73 = Array((new BigNumber(13)).toNumber());
          _nw73[(_dafny.ZERO).toNumber()] = _587_v73;
          _nw73[(_dafny.ONE).toNumber()] = _module.D7.create_DC16(_586_v72);
          _nw73[(new BigNumber(2)).toNumber()] = _587_v73;
          _nw73[(new BigNumber(3)).toNumber()] = _587_v73;
          _nw73[(new BigNumber(4)).toNumber()] = _587_v73;
          _nw73[(new BigNumber(5)).toNumber()] = _587_v73;
          _nw73[(new BigNumber(6)).toNumber()] = _587_v73;
          _nw73[(new BigNumber(7)).toNumber()] = _module.D7.create_DC16(_586_v72);
          _nw73[(new BigNumber(8)).toNumber()] = _587_v73;
          _nw73[(new BigNumber(9)).toNumber()] = _587_v73;
          _nw73[(new BigNumber(10)).toNumber()] = _587_v73;
          _nw73[(new BigNumber(11)).toNumber()] = _module.D7.create_DC16(_586_v72);
          _nw73[(new BigNumber(12)).toNumber()] = _587_v73;
          _588_v74 = _nw73;
          let _589_v75;
          _589_v75 = _dafny.Seq.of(_588_v74, _588_v74);
          let _590_v76;
          _590_v76 = _dafny.Map.Empty.slice().updateUnsafe(_563_v56,_563_v56);
          let _591_v77;
          _591_v77 = _dafny.MultiSet.fromElements(new BigNumber((_590_v76).length), (_this).f28);
          _583_v69 = _dafny.Seq.of((_this).f28, new BigNumber((_589_v75).length), new BigNumber(((_591_v77).Intersect(_dafny.MultiSet.fromElements((_this).f28, new BigNumber((_564_v57).length)))).cardinality()), (_583_v69)[_module.__default.safeIndex((_this).f28, new BigNumber((_583_v69).length))]);
          let _592_v78;
          let _nw74 = new _module.C0();
          _nw74.__ctor(_563_v56);
          _592_v78 = _nw74;
          _583_v69 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_566_v59).length), (_this).f28), _dafny.Seq.of(new BigNumber(-630))), _583_v69);
        }
      }
      if (true) {
        let _593_v79;
        let _nw75 = Array((new BigNumber(27)).toNumber()).fill(false);
        _593_v79 = _nw75;
        let _594_v80;
        _594_v80 = false;
        let _index81 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_593_v79).length));
        (_593_v79)[_index81] = _594_v80;
        let _595_v81;
        _595_v81 = _dafny.Seq.UnicodeFromString("erwiuf");
        let _596_v82;
        _596_v82 = _dafny.Map.Empty.slice().updateUnsafe(_594_v80,_595_v81);
        let _index82 = _module.__default.safeIndex(new BigNumber(864), new BigNumber((_593_v79).length));
        (_593_v79)[_index82] = (_module.__default.fm23(_594_v80, (((_596_v82).contains(true)) ? ((_596_v82).get(true)) : (_595_v81)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(50)), function (_597_i5) {
          return new _dafny.CodePoint('i'.codePointAt(0));
        }), _594_v80, globalState)).dtor_cf34;
        let _598_v83;
        _598_v83 = _dafny.Seq.of(new BigNumber((_595_v81).length), (_this).f28);
        let _599_v84;
        _599_v84 = _dafny.Set.fromElements((_598_v83)[_module.__default.safeIndex((_this).f28, new BigNumber((_598_v83).length))], (_this).f28, (_dafny.ZERO).minus((_this).f28));
        let _index83 = _module.__default.safeIndex(new BigNumber(268), new BigNumber((p0).length));
        (p0)[_index83] = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_599_v84).length)));
        let _index84 = _module.__default.safeIndex(new BigNumber(268), new BigNumber((p0).length));
        (p0)[_index84] = ((_this).f28).plus((_dafny.ZERO).minus((_this).f28));
        let _600_v85;
        _600_v85 = _dafny.Map.Empty.slice().updateUnsafe((_593_v79)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_593_v79).length))],_593_v79);
        let _601_v86;
        _601_v86 = _dafny.Map.Empty.slice().updateUnsafe(_600_v85,new BigNumber(-210));
        _601_v86 = (_601_v86).update((_600_v85).Merge(_600_v85), (p0)[_module.__default.safeIndex(new BigNumber(268), new BigNumber((p0).length))]);
        let _602_v87;
        _602_v87 = _dafny.Map.Empty.slice().updateUnsafe(_594_v80,(_dafny.ZERO).minus((_this).f28));
        let _603_v88;
        _603_v88 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,new BigNumber((_602_v87).length));
        _603_v88 = (_603_v88).update((p0)[_module.__default.safeIndex(new BigNumber(268), new BigNumber((p0).length))], (p0)[_module.__default.safeIndex(new BigNumber(268), new BigNumber((p0).length))]);
        let _604_v89;
        let _nw76 = new _module.C0();
        _nw76.__ctor((_593_v79)[_module.__default.safeIndex(new BigNumber(864), new BigNumber((_593_v79).length))]);
        _604_v89 = _nw76;
      } else {
        let _605_v90;
        _605_v90 = true;
        let _606_v91;
        _606_v91 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_605_v90,(_this).f28)).length),new _dafny.CodePoint('i'.codePointAt(0)));
        let _607_v92;
        _607_v92 = new _dafny.CodePoint('a'.codePointAt(0));
        let _608_v93;
        _608_v93 = _dafny.Map.Empty.slice().updateUnsafe(_605_v90,_605_v90);
        let _609_v94;
        _609_v94 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("cdoyd")).length),_608_v93);
        let _610_v95;
        _610_v95 = _dafny.Map.Empty.slice().updateUnsafe(_607_v92,_605_v90);
        let _611_v96;
        _611_v96 = _dafny.MultiSet.fromElements(_610_v95);
        (globalState).f21 = !((_module.__default.fm24((((_606_v91).contains(new BigNumber(493))) ? ((_606_v91).get(new BigNumber(493))) : (_607_v92)), new BigNumber((_609_v94).length), new BigNumber(342), new BigNumber(724), globalState)).IsDisjointFrom(_611_v96));
        let _index85 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((p0).length));
        (p0)[_index85] = (_this).f28;
        let _612_v98;
        _612_v98 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
          let _coll31 = new _dafny.Map();
          for (const _compr_31 of _dafny.IntegerRange(new BigNumber(-121), new BigNumber(315))) {
            let _613_v97 = _compr_31;
            if (((new BigNumber(-121)).isLessThanOrEqualTo(_613_v97)) && ((_613_v97).isLessThan(new BigNumber(315)))) {
              _coll31.push([(_613_v97).minus((_this).f28),false]);
            }
          }
          return _coll31;
        }()).length),_605_v90);
        let _index86 = _module.__default.safeIndex(new BigNumber(273), new BigNumber((p0).length));
        (p0)[_index86] = _module.__default.safeDivisionInt(new BigNumber(((_612_v98).Merge(_612_v98)).length), (_this).f28);
        (globalState).f21 = _605_v90;
        (globalState).f0 = ((_this).f28).isLessThan((p0)[_module.__default.safeIndex(new BigNumber(273), new BigNumber((p0).length))]);
        let _614_v99;
        _614_v99 = _dafny.Seq.of(_605_v90, !(_605_v90));
        let _615_v100;
        _615_v100 = _dafny.Map.Empty.slice().updateUnsafe(_605_v90,_614_v99);
        _615_v100 = (_615_v100).update(false, _614_v99);
      }
      let _616_v101;
      _616_v101 = _dafny.Seq.UnicodeFromString("jeovgrwpv");
      let _617_v102;
      _617_v102 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(181)), function (_618_i6) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      }), _616_v101, _dafny.Seq.Create(_module.__default.abs(new BigNumber(528)), function (_619_i7) {
        return new _dafny.CodePoint('t'.codePointAt(0));
      }));
      _617_v102 = ((_617_v102).Difference(_617_v102)).Union(_617_v102);
      let _620_v103;
      _620_v103 = true;
      let _621_v104;
      _621_v104 = _dafny.MultiSet.fromElements(_620_v103);
      let _622_v105;
      _622_v105 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(213),!(false));
      let _623_v106;
      _623_v106 = _dafny.Seq.of(_module.__default.fm2(_dafny.Map.Empty.slice().updateUnsafe(_620_v103,_620_v103), globalState));
      let _624_v107;
      _624_v107 = _623_v106;
      let _hi2 = ((_this).f28).plus((_this).fm17(_620_v103, _dafny.Seq.of((_this).f28), new BigNumber((_623_v106).length), _624_v107, globalState));
      for (let _625_i8 = ((((_621_v104).contains(true)) ? ((_621_v104).get(true)) : ((((_621_v104).contains(_620_v103)) ? ((_621_v104).get(_620_v103)) : (new BigNumber((_622_v105).length)))))).plus(new BigNumber(921)); _625_i8.isLessThan(_hi2); _625_i8 = _625_i8.plus(_dafny.ONE)) {
        let _626_v108;
        _626_v108 = new _dafny.CodePoint('o'.codePointAt(0));
        (_this).m8(_dafny.Seq.contains(_module.__default.fm1(_625_i8, globalState), _626_v108), _620_v103, _620_v103, globalState);
        let _627_v109;
        _627_v109 = _dafny.Map.Empty.slice().updateUnsafe(_620_v103,(_this).f28);
        (globalState).f16 = _module.__default.safeModuloInt(new BigNumber(((_627_v109).update(_620_v103, (_this).f28)).length), _module.__default.safeModuloInt(new BigNumber(762), _625_i8));
        let _628_v110;
        _628_v110 = _module.D6.create_DC13((_this).f28, _625_i8);
        let _629_v111;
        _629_v111 = _module.D6.create_DC14(_628_v110);
        let _source13 = _629_v111;
        if (_source13.is_DC13) {
          let _630___mcc_h4 = (_source13).cf19;
          let _631___mcc_h5 = (_source13).cf20;
          let _632_cf20 = _631___mcc_h5;
          let _633_cf19 = _630___mcc_h4;
          let _634_v112;
          _634_v112 = _dafny.Seq.of(p0, p0);
          let _635_v113;
          _635_v113 = _module.D6.create_DC12(_634_v112);
          let _636_v114;
          _636_v114 = _dafny.MultiSet.fromElements(_632_cf20);
          let _pat_let_tv11 = _634_v112;
          let _pat_let_tv12 = _634_v112;
          let _pat_let_tv13 = p0;
          let _rhs64 = _632_cf20;
          let _rhs65 = _632_cf20;
          let _rhs66 = (_620_v103) === (_620_v103);
          let _rhs67 = (_dafny.MultiSet.fromElements(new BigNumber((_616_v101).length))).IsProperSubsetOf(_636_v114);
          let _rhs68 = function (_pat_let8_0) {
            return function (_637_dt__update__tmp_h1) {
              return function (_pat_let9_0) {
                return function (_638_dt__update_hcf18_h0) {
                  return _module.D6.create_DC12(_638_dt__update_hcf18_h0);
                }(_pat_let9_0);
              }(_dafny.Seq.update(_pat_let_tv11, _module.__default.safeIndex(new BigNumber(677), new BigNumber((_pat_let_tv12).length)), _pat_let_tv13));
            }(_pat_let8_0);
          }(_module.D6.create_DC12(_634_v112));
          let _lhs67 = globalState;
          let _lhs68 = globalState;
          let _lhs69 = globalState;
          let _lhs70 = globalState;
          _lhs67.f7 = _rhs64;
          _lhs68.f7 = _rhs65;
          _lhs69.f21 = _rhs66;
          _lhs70.f21 = _rhs67;
          _635_v113 = _rhs68;
          (globalState).f13 = !(!(_620_v103) || ((_620_v103) === (_620_v103)));
          (globalState).f2 = _620_v103;
          _626_v108 = _626_v108;
        } else if (_source13.is_DC12) {
          let _639___mcc_h6 = (_source13).cf18;
          let _640_cf18 = _639___mcc_h6;
          let _641_v115;
          _641_v115 = _dafny.Set.fromElements(_620_v103);
          let _642_v116;
          let _nw77 = new _module.C0();
          _nw77.__ctor((_dafny.Set.fromElements(_620_v103, !(_620_v103), true)).IsSubsetOf(_641_v115));
          _642_v116 = _nw77;
          _623_v106 = _dafny.Seq.Concat(_dafny.Seq.of(_642_v116.f27), _623_v106);
          (_this).m8((_621_v104).IsProperSubsetOf(_621_v104), _620_v103, _642_v116.f27, globalState);
          let _643_v117;
          let _init12 = ((_644_i8, _645_v105, _646_v103) => function (_647_i9) {
            return (((_645_v105).contains(_644_i8)) ? ((_645_v105).get(_644_i8)) : (_646_v103));
          })(_625_i8, _622_v105, _620_v103);
          let _nw78 = Array((new BigNumber(26)).toNumber());
          for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw78.length); _i0_12++) {
            _nw78[_i0_12] = _init12(new BigNumber(_i0_12));
          }
          _643_v117 = _nw78;
          let _648_v118;
          _648_v118 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,_623_v106);
          let _649_v119;
          _649_v119 = _dafny.Seq.of((((_648_v118).contains((_dafny.ZERO).minus((_this).f28))) ? ((_648_v118).get((_dafny.ZERO).minus((_this).f28))) : (_623_v106)));
          let _index87 = _module.__default.safeIndex(new BigNumber(799), new BigNumber((_643_v117).length));
          (_643_v117)[_index87] = _dafny.Seq.contains((_649_v119)[_module.__default.safeIndex(_625_i8, new BigNumber((_649_v119).length))], _620_v103);
          let _index88 = _module.__default.safeIndex(new BigNumber(799), new BigNumber((_643_v117).length));
          (_643_v117)[_index88] = ((_620_v103) ? ((_620_v103) && (_620_v103)) : ((_625_i8).isLessThanOrEqualTo((_this).f28)));
        } else {
          let _650___mcc_h7 = (_source13).cf21;
          let _651_cf21 = _650___mcc_h7;
          (globalState).f2 = (_620_v103) === (_620_v103);
          let _652_v120;
          _652_v120 = _module.D0.create_DC0(_620_v103);
          let _653_v121;
          _653_v121 = _dafny.Seq.of((_dafny.MultiSet.fromElements(!((_652_v120).dtor_cf0))).update(!(_620_v103), _module.__default.abs(_625_i8)));
          _621_v104 = (_653_v121)[_module.__default.safeIndex(_625_i8, new BigNumber((_653_v121).length))];
          let _654_v122;
          let _nw79 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          _654_v122 = _nw79;
          _654_v122 = ((((_this).fm6(new BigNumber(554), globalState)) && (_620_v103)) ? (_654_v122) : (p0));
          let _655_v123;
          let _nw80 = Array((new BigNumber(17)).toNumber()).fill(false);
          _655_v123 = _nw80;
          let _656_v124;
          _656_v124 = _dafny.Map.Empty.slice().updateUnsafe(_655_v123,_625_i8);
          _656_v124 = (_656_v124).update(_655_v123, _625_i8);
        }
        let _657_v125;
        let _nw81 = new _module.C0();
        _nw81.__ctor(_620_v103);
        _657_v125 = _nw81;
      }
      let _658_v126;
      let _nw82 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _658_v126 = _nw82;
      let _index89 = _module.__default.safeIndex(new BigNumber(873), new BigNumber((_658_v126).length));
      (_658_v126)[_index89] = _dafny.Seq.Concat(_616_v101, _module.__default.fm1((_this).f28, globalState));
      let _index90 = _module.__default.safeIndex(new BigNumber(873), new BigNumber((_658_v126).length));
      (_658_v126)[_index90] = _616_v101;
      return;
    }
    m1(p0, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = false;
      let _659_v0;
      _659_v0 = _dafny.Set.fromElements(new BigNumber(939), p0, (_this).f28);
      let _660_v1;
      _660_v1 = true;
      let _661_v2;
      _661_v2 = new _dafny.CodePoint('a'.codePointAt(0));
      let _662_v3;
      _662_v3 = _dafny.Seq.UnicodeFromString("prnrf");
      let _663_v4;
      _663_v4 = _dafny.Seq.of(!(_660_v1), _module.__default.fm2(_dafny.Map.Empty.slice().updateUnsafe(_660_v1,_660_v1), globalState));
      let _664_v5;
      _664_v5 = _module.D1.create_DC3(_660_v1);
      let _665_v6;
      _665_v6 = _dafny.Set.fromElements(true, _660_v1, _660_v1);
      let _666_v7;
      let _nw83 = Array((new BigNumber(16)).toNumber());
      _nw83[(_dafny.ZERO).toNumber()] = !((new BigNumber((_659_v0).length)).isLessThanOrEqualTo((_this).f28));
      _nw83[(_dafny.ONE).toNumber()] = _660_v1;
      _nw83[(new BigNumber(2)).toNumber()] = !_dafny.Seq.contains(_662_v3, _661_v2);
      _nw83[(new BigNumber(3)).toNumber()] = ((_660_v1) ? ((_663_v4)[_module.__default.safeIndex((_this).f28, new BigNumber((_663_v4).length))]) : ((_664_v5).dtor_cf3));
      _nw83[(new BigNumber(4)).toNumber()] = (_660_v1) === (_660_v1);
      _nw83[(new BigNumber(5)).toNumber()] = _660_v1;
      _nw83[(new BigNumber(6)).toNumber()] = _660_v1;
      _nw83[(new BigNumber(7)).toNumber()] = (_this).fm6(new BigNumber((_665_v6).length), globalState);
      _nw83[(new BigNumber(8)).toNumber()] = _660_v1;
      _nw83[(new BigNumber(9)).toNumber()] = !(true) || (_660_v1);
      _nw83[(new BigNumber(10)).toNumber()] = _660_v1;
      _nw83[(new BigNumber(11)).toNumber()] = _660_v1;
      _nw83[(new BigNumber(12)).toNumber()] = _660_v1;
      _nw83[(new BigNumber(13)).toNumber()] = _660_v1;
      _nw83[(new BigNumber(14)).toNumber()] = _660_v1;
      _nw83[(new BigNumber(15)).toNumber()] = _660_v1;
      _666_v7 = _nw83;
      for (const _guard_loop_5 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_666_v7).length))) {
        let _667_i0 = _guard_loop_5;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_667_i0)) && ((_667_i0).isLessThan(new BigNumber((_666_v7).length))))) {
          (_666_v7)[(_667_i0)] = _660_v1;
        }
      }
      let _668_v8;
      _668_v8 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_this).f28);
      let _669_v9;
      let _init13 = ((_670_v3) => function (_671_i1) {
        return (_670_v3)[_module.__default.safeIndex((_this).f28, new BigNumber((_670_v3).length))];
      })(_662_v3);
      let _nw84 = Array((new BigNumber(15)).toNumber());
      for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw84.length); _i0_13++) {
        _nw84[_i0_13] = _init13(new BigNumber(_i0_13));
      }
      _669_v9 = _nw84;
      let _672_v10;
      _672_v10 = _dafny.Map.Empty.slice().updateUnsafe(p0,_669_v9);
      _668_v8 = (_668_v8).update(p0, (new BigNumber((_672_v10).length)).plus((_this).f28));
      if (_660_v1) {
        let _673_v11;
        let _init14 = ((_674_v1, _675_p0) => function (_676_i2) {
          return _module.__default.safeDivisionInt(_676_i2, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_674_v1,_675_p0)).length));
        })(_660_v1, p0);
        let _nw85 = Array((new BigNumber(28)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw85.length); _i0_14++) {
          _nw85[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _673_v11 = _nw85;
        let _index91 = _module.__default.safeIndex(new BigNumber(261), new BigNumber((_673_v11).length));
        (_673_v11)[_index91] = p0;
        let _677_v12;
        _677_v12 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-535)), ((_678_v2) => function (_679_i3) {
          return _678_v2;
        })(_661_v2)), _dafny.Seq.UnicodeFromString("lwkghweu"), _662_v3);
        let _680_v13;
        _680_v13 = _dafny.Seq.of((_677_v12)[_module.__default.safeIndex(new BigNumber(509), new BigNumber((_677_v12).length))], _662_v3, _662_v3, _dafny.Seq.UnicodeFromString("hpd"));
        let _681_v14;
        _681_v14 = _dafny.Map.Empty.slice().updateUnsafe(_663_v4,_666_v7);
        let _682_v15;
        _682_v15 = _dafny.Map.Empty.slice().updateUnsafe(_668_v8,_661_v2);
        let _683_v17;
        _683_v17 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(p0,(_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(_660_v1, _660_v1, false, _660_v1, true)).cardinality()))), _668_v8);
        let _index92 = _module.__default.safeIndex(new BigNumber(261), new BigNumber((_673_v11).length));
        let _rhs69 = _module.__default.safeModuloInt(p0, (_this).f28);
        let _rhs70 = (function () {
          let _coll32 = new _dafny.Set();
          for (const _compr_32 of (_682_v15).Keys.Elements) {
            let _684_v16 = _compr_32;
            if ((_682_v15).contains(_684_v16)) {
              _coll32.add(_684_v16);
            }
          }
          return _coll32;
        }()).IsDisjointFrom(_683_v17);
        let _rhs71 = _677_v12;
        let _rhs72 = ((_681_v14).Merge(_681_v14)).Merge(_681_v14);
        let _lhs71 = _673_v11;
        let _lhs72 = _module.__default.safeIndex(new BigNumber(261), new BigNumber((_673_v11).length));
        let _lhs73 = globalState;
        _lhs71[_lhs72] = _rhs69;
        _lhs73.f21 = _rhs70;
        _677_v12 = _rhs71;
        _681_v14 = _rhs72;
        _660_v1 = false;
        (globalState).f13 = !(_660_v1);
        (globalState).f21 = _660_v1;
        let _685_v18;
        let _nw86 = Array((new BigNumber(10)).toNumber());
        _nw86[(_dafny.ZERO).toNumber()] = _662_v3;
        _nw86[(_dafny.ONE).toNumber()] = _module.__default.fm1(new BigNumber((_662_v3).length), globalState);
        _nw86[(new BigNumber(2)).toNumber()] = _662_v3;
        _nw86[(new BigNumber(3)).toNumber()] = (_module.__default.fm25(_660_v1, (_this).f28, _660_v1, new BigNumber(842), globalState)).dtor_cf6;
        _nw86[(new BigNumber(4)).toNumber()] = _662_v3;
        _nw86[(new BigNumber(5)).toNumber()] = _662_v3;
        _nw86[(new BigNumber(6)).toNumber()] = _662_v3;
        _nw86[(new BigNumber(7)).toNumber()] = _dafny.Seq.update(_662_v3, _module.__default.safeIndex(p0, new BigNumber((_662_v3).length)), _661_v2);
        _nw86[(new BigNumber(8)).toNumber()] = _662_v3;
        _nw86[(new BigNumber(9)).toNumber()] = _662_v3;
        _685_v18 = _nw86;
        let _index93 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_685_v18).length));
        (_685_v18)[_index93] = _662_v3;
        let _index94 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_685_v18).length));
        (_685_v18)[_index94] = _dafny.Seq.Concat(_662_v3, _662_v3);
      } else {
        let _686_v19;
        let _nw87 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
        _686_v19 = _nw87;
        let _687_v20;
        _687_v20 = _dafny.MultiSet.fromElements(_661_v2);
        let _688_v21;
        _688_v21 = _dafny.Map.Empty.slice().updateUnsafe(_660_v1,new BigNumber((_687_v20).cardinality()));
        let _689_v22;
        _689_v22 = _dafny.Map.Empty.slice().updateUnsafe(_688_v21,_660_v1);
        let _690_v23;
        _690_v23 = _dafny.Map.Empty.slice().updateUnsafe(_660_v1,_686_v19);
        let _691_v24;
        let _nw88 = Array((new BigNumber(12)).toNumber());
        _nw88[(_dafny.ZERO).toNumber()] = _686_v19;
        _nw88[(_dafny.ONE).toNumber()] = _686_v19;
        _nw88[(new BigNumber(2)).toNumber()] = _686_v19;
        _nw88[(new BigNumber(3)).toNumber()] = _686_v19;
        _nw88[(new BigNumber(4)).toNumber()] = _686_v19;
        _nw88[(new BigNumber(5)).toNumber()] = _686_v19;
        _nw88[(new BigNumber(6)).toNumber()] = (((_this).fm7(false, p0, (((_689_v22).contains(_688_v21)) ? ((_689_v22).get(_688_v21)) : (_660_v1)), globalState)) ? (_686_v19) : (_686_v19));
        _nw88[(new BigNumber(7)).toNumber()] = _686_v19;
        _nw88[(new BigNumber(8)).toNumber()] = _686_v19;
        _nw88[(new BigNumber(9)).toNumber()] = (((_690_v23).contains(_660_v1)) ? ((_690_v23).get(_660_v1)) : (_686_v19));
        _nw88[(new BigNumber(10)).toNumber()] = _686_v19;
        _nw88[(new BigNumber(11)).toNumber()] = _686_v19;
        _691_v24 = _nw88;
        let _index95 = _module.__default.safeIndex(new BigNumber(141), new BigNumber((_691_v24).length));
        (_691_v24)[_index95] = _686_v19;
        let _692_v25;
        _692_v25 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_663_v4));
        let _693_v26;
        _693_v26 = _dafny.Seq.of(new BigNumber(445));
        let _694_v27;
        _694_v27 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(553),_693_v26);
        let _695_v28;
        _695_v28 = _dafny.MultiSet.fromElements(true);
        let _696_v29;
        _696_v29 = _module.D4.create_DC10(p0, (_this).f28, (_this).fm7(_660_v1, (_this).f28, _660_v1, globalState), _660_v1, p0);
        let _index96 = _module.__default.safeIndex(new BigNumber(141), new BigNumber((_691_v24).length));
        let _rhs73 = ((_660_v1) ? (((_660_v1) ? (_686_v19) : (_686_v19))) : (_686_v19));
        let _rhs74 = _dafny.Seq.update(_dafny.Seq.update(_692_v25, _module.__default.safeIndex(new BigNumber((_694_v27).length), new BigNumber((_692_v25).length)), (_dafny.MultiSet.FromArray(_663_v4)).Difference(_695_v28)), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.update(_692_v25, _module.__default.safeIndex(new BigNumber((_694_v27).length), new BigNumber((_692_v25).length)), (_dafny.MultiSet.FromArray(_663_v4)).Difference(_695_v28))).length)), (_695_v28).update(_660_v1, _module.__default.abs((_this).f28)));
        let _rhs75 = (_696_v29).dtor_cf16;
        let _rhs76 = _660_v1;
        let _lhs74 = _691_v24;
        let _lhs75 = _module.__default.safeIndex(new BigNumber(141), new BigNumber((_691_v24).length));
        let _lhs76 = globalState;
        let _lhs77 = globalState;
        _lhs74[_lhs75] = _rhs73;
        _692_v25 = _rhs74;
        _lhs76.f7 = _rhs75;
        _lhs77.f0 = _rhs76;
        let _697_v30;
        let _nw89 = new _module.C0();
        _nw89.__ctor(true);
        _697_v30 = _nw89;
        let _698_v31;
        let _nw90 = new _module.C0();
        _nw90.__ctor((p0).isLessThanOrEqualTo((_this).f28));
        _698_v31 = _nw90;
        if (_660_v1) {
          let _699_v32;
          _699_v32 = _dafny.Map.Empty.slice().updateUnsafe((p0).isEqualTo((_this).f28),(_this).fm6(new BigNumber((_662_v3).length), globalState));
          _699_v32 = (_699_v32).update((_660_v1) || (_698_v31.f27), _698_v31.f27);
          r1 = _698_v31.f27;
          let _700_v33;
          let _nw91 = Array((new BigNumber(3)).toNumber()).fill([]);
          _700_v33 = _nw91;
          let _701_v34;
          _701_v34 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).fm6((_this).f28, globalState));
          let _702_v37;
          let _nw92 = Array((new BigNumber(7)).toNumber());
          _nw92[(_dafny.ZERO).toNumber()] = _659_v0;
          _nw92[(_dafny.ONE).toNumber()] = _dafny.Set.fromElements(new BigNumber((_701_v34).length));
          _nw92[(new BigNumber(2)).toNumber()] = _659_v0;
          _nw92[(new BigNumber(3)).toNumber()] = function () {
            let _coll33 = new _dafny.Set();
            for (const _compr_33 of _dafny.IntegerRange(new BigNumber(630), new BigNumber(607))) {
              let _703_v35 = _compr_33;
              if (((new BigNumber(630)).isLessThanOrEqualTo(_703_v35)) && ((_703_v35).isLessThan(new BigNumber(607)))) {
                _coll33.add((_703_v35).minus((_this).f28));
              }
            }
            return _coll33;
          }();
          _nw92[(new BigNumber(4)).toNumber()] = _659_v0;
          _nw92[(new BigNumber(5)).toNumber()] = _659_v0;
          _nw92[(new BigNumber(6)).toNumber()] = function () {
            let _coll34 = new _dafny.Set();
            for (const _compr_34 of _dafny.IntegerRange(new BigNumber(910), new BigNumber(974))) {
              let _704_v36 = _compr_34;
              if (((new BigNumber(910)).isLessThanOrEqualTo(_704_v36)) && ((_704_v36).isLessThan(new BigNumber(974)))) {
                _coll34.add((_704_v36).minus((_this).f28));
              }
            }
            return _coll34;
          }();
          _702_v37 = _nw92;
          let _index97 = _module.__default.safeIndex(new BigNumber(650), new BigNumber((_700_v33).length));
          (_700_v33)[_index97] = _702_v37;
          let _705_v38;
          _705_v38 = _dafny.Seq.of(_701_v34, _701_v34);
          let _index98 = _module.__default.safeIndex(new BigNumber(650), new BigNumber((_700_v33).length));
          let _rhs77 = _666_v7;
          let _rhs78 = new BigNumber((_dafny.Seq.Concat(_705_v38, _705_v38)).length);
          let _rhs79 = ((_698_v31.f27) ? (_702_v37) : (_702_v37));
          let _lhs78 = globalState;
          let _lhs79 = _700_v33;
          let _lhs80 = _module.__default.safeIndex(new BigNumber(650), new BigNumber((_700_v33).length));
          _666_v7 = _rhs77;
          _lhs78.f7 = _rhs78;
          _lhs79[_lhs80] = _rhs79;
          _666_v7 = _666_v7;
          let _706_v39;
          _706_v39 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_698_v31.f27),(_this).f28);
          _706_v39 = (_706_v39).update(_dafny.Seq.Concat(_dafny.Seq.of(_698_v31.f27), _663_v4), (_this).f28);
        } else {
          (globalState).f2 = true;
          _693_v26 = _693_v26;
          (globalState).f7 = new BigNumber(519);
          _666_v7 = _666_v7;
          let _707_v40;
          _707_v40 = _dafny.Map.Empty.slice().updateUnsafe(_697_v30.f27,_661_v2);
          let _708_v41;
          _708_v41 = _dafny.Map.Empty.slice().updateUnsafe(_662_v3,new BigNumber((_707_v40).length));
          _708_v41 = (_708_v41).update(_dafny.Seq.UnicodeFromString("vvrdpbh"), p0);
        }
        let _arr0 = (_691_v24)[_module.__default.safeIndex(new BigNumber(141), new BigNumber((_691_v24).length))];
        let _index99 = _module.__default.safeIndex(new BigNumber(365), new BigNumber(((_691_v24)[_module.__default.safeIndex(new BigNumber(141), new BigNumber((_691_v24).length))]).length));
        _arr0[_index99] = _module.__default.safeDivisionInt((_this).f28, p0);
        let _709_v42;
        _709_v42 = _module.D6.create_DC13((_this).f28, new BigNumber((_dafny.MultiSet.fromElements(_698_v31.f27)).cardinality()));
        let _710_v43;
        _710_v43 = _module.D6.create_DC14(_709_v42);
        let _711_v44;
        _711_v44 = _module.D6.create_DC14(_709_v42);
        let _712_v45;
        _712_v45 = _dafny.Seq.of(_662_v3);
        let _713_v46;
        _713_v46 = _dafny.Seq.of((_712_v45)[_module.__default.safeIndex((_this).f28, new BigNumber((_712_v45).length))], _662_v3, _662_v3, _dafny.Seq.UnicodeFromString("ikcvsk"));
        let _arr1 = (_691_v24)[_module.__default.safeIndex(new BigNumber(141), new BigNumber((_691_v24).length))];
        let _index100 = _module.__default.safeIndex(new BigNumber(365), new BigNumber(((_691_v24)[_module.__default.safeIndex(new BigNumber(141), new BigNumber((_691_v24).length))]).length));
        let _rhs80 = (_dafny.ZERO).minus((_this).f28);
        let _rhs81 = _711_v44;
        let _rhs82 = _module.__default.safeModuloInt(p0, p0);
        let _rhs83 = (_713_v46)[_module.__default.safeIndex((_this).f28, new BigNumber((_713_v46).length))];
        let _lhs81 = (_691_v24)[_module.__default.safeIndex(new BigNumber(141), new BigNumber((_691_v24).length))];
        let _lhs82 = _module.__default.safeIndex(new BigNumber(365), new BigNumber(((_691_v24)[_module.__default.safeIndex(new BigNumber(141), new BigNumber((_691_v24).length))]).length));
        let _lhs83 = globalState;
        let _lhs84 = globalState;
        _lhs81[_lhs82] = _rhs80;
        _711_v44 = _rhs81;
        _lhs83.f7 = _rhs82;
        _lhs84.f20 = _rhs83;
      }
      _668_v8 = (_668_v8).update((_this).f28, p0);
      let _hi3 = (_this).f28;
      for (let _714_i4 = p0; _714_i4.isLessThan(_hi3); _714_i4 = _714_i4.plus(_dafny.ONE)) {
        let _715_v47;
        let _nw93 = Array((_dafny.ONE).toNumber()).fill(_dafny.Map.Empty);
        _715_v47 = _nw93;
        let _716_v48;
        _716_v48 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,_660_v1);
        let _index101 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_715_v47).length));
        (_715_v47)[_index101] = (_716_v48).Merge(_716_v48);
        let _index102 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_715_v47).length));
        (_715_v47)[_index102] = (_716_v48).Merge(_dafny.Map.Empty.slice().updateUnsafe(_714_i4,_660_v1));
        (globalState).f21 = (false) && (_660_v1);
        let _717_v49;
        let _nw94 = new _module.C0();
        _nw94.__ctor(_660_v1);
        _717_v49 = _nw94;
        (globalState).f19 = _dafny.Seq.Concat(_662_v3, _662_v3);
      }
      let _718_v50;
      _718_v50 = _dafny.Map.Empty.slice().updateUnsafe(p0,true);
      if ((_659_v0).IsProperSubsetOf(_module.__default.fm26(_661_v2, _718_v50, globalState))) {
        let _719_v51;
        _719_v51 = _dafny.Map.Empty.slice().updateUnsafe(!_dafny.Seq.contains(_662_v3, _661_v2),_module.__default.fm27(!(_660_v1), _660_v1, globalState));
        _719_v51 = (_719_v51).update(_660_v1, ((_660_v1) ? (new _dafny.CodePoint('q'.codePointAt(0))) : (_661_v2)));
        (globalState).f8 = _662_v3;
        let _720_v52;
        _720_v52 = _dafny.MultiSet.fromElements(_660_v1, _660_v1);
        let _721_v53;
        let _nw95 = new _module.C0();
        _nw95.__ctor(((_this).f28).isEqualTo(new BigNumber((_720_v52).cardinality())));
        _721_v53 = _nw95;
        let _722_v54;
        let _nw96 = new _module.C0();
        _nw96.__ctor(!((_this).fm6((_this).f28, globalState)));
        _722_v54 = _nw96;
        let _723_v55;
        let _nw97 = new _module.C0();
        _nw97.__ctor(_721_v53.f27);
        _723_v55 = _nw97;
      } else {
        if (_660_v1) {
          (globalState).f16 = (_dafny.ZERO).minus(p0);
          let _724_v56;
          let _nw98 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
          _724_v56 = _nw98;
          let _725_v57;
          _725_v57 = _dafny.Map.Empty.slice().updateUnsafe(_724_v56,new BigNumber((((_660_v1) ? (_662_v3) : (_662_v3))).length));
          _725_v57 = (_725_v57).update(_724_v56, (new BigNumber(-586)).multipliedBy(p0));
          let _726_v58;
          let _727_v59;
          let _728_v60;
          let _729_v61;
          let _out0;
          let _out1;
          let _out2;
          let _out3;
          let _outcollector0 = (_this).m9(globalState);
          _out0 = _outcollector0[0];
          _out1 = _outcollector0[1];
          _out2 = _outcollector0[2];
          _out3 = _outcollector0[3];
          _726_v58 = _out0;
          _727_v59 = _out1;
          _728_v60 = _out2;
          _729_v61 = _out3;
          _668_v8 = ((_668_v8).Merge(function () {
            let _coll35 = new _dafny.Map();
            for (const _compr_35 of (_659_v0).Elements) {
              let _730_v62 = _compr_35;
              if ((_659_v0).contains(_730_v62)) {
                _coll35.push([(_730_v62).minus((_this).f28),(((_668_v8).contains((_dafny.ZERO).minus((_this).f28))) ? ((_668_v8).get((_dafny.ZERO).minus((_this).f28))) : ((_this).f28))]);
              }
            }
            return _coll35;
          }())).Merge(((false) ? (_668_v8) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(493),new BigNumber((_662_v3).length)))));
          (globalState).f13 = _module.__default.fm2(_dafny.Map.Empty.slice().updateUnsafe(_726_v58,_727_v59), globalState);
        } else {
          let _731_v63;
          _731_v63 = _module.D4.create_DC7((_this).f28);
          let _732_v64;
          _732_v64 = _dafny.MultiSet.fromElements(!(_660_v1));
          let _733_v65;
          let _nw99 = Array((new BigNumber(9)).toNumber());
          _nw99[(_dafny.ZERO).toNumber()] = p0;
          _nw99[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_731_v63).dtor_cf9,true)).length));
          _nw99[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(986)), ((_734_v2) => function (_735_i5) {
            return _734_v2;
          })(_661_v2))).length));
          _nw99[(new BigNumber(3)).toNumber()] = (_this).f28;
          _nw99[(new BigNumber(4)).toNumber()] = p0;
          _nw99[(new BigNumber(5)).toNumber()] = p0;
          _nw99[(new BigNumber(6)).toNumber()] = new BigNumber((_732_v64).cardinality());
          _nw99[(new BigNumber(7)).toNumber()] = p0;
          _nw99[(new BigNumber(8)).toNumber()] = (_this).f28;
          _733_v65 = _nw99;
          r0 = ((_660_v1) ? (_733_v65) : (_733_v65));
          let _736_v66;
          _736_v66 = _dafny.Map.Empty.slice().updateUnsafe(_660_v1,_660_v1);
          _736_v66 = (_736_v66).update(!(_660_v1) || (_660_v1), _660_v1);
          (globalState).f7 = (_this).f28;
          let _737_v67;
          let _nw100 = new _module.C0();
          _nw100.__ctor(_660_v1);
          _737_v67 = _nw100;
          let _738_v68;
          _738_v68 = _dafny.Seq.of(new BigNumber(977), (_this).f28);
          let _739_v69;
          _739_v69 = _dafny.Seq.of(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(989)), function (_740_i8) {
            return new BigNumber((_dafny.Seq.UnicodeFromString("efk")).length);
          }), _module.__default.safeIndex((_this).f28, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(989)), function (_741_i8) {
            return new BigNumber((_dafny.Seq.UnicodeFromString("efk")).length);
          })).length)), (_this).f28), _dafny.Seq.Create(_module.__default.abs(new BigNumber(396)), ((_742_p0) => function (_743_i9) {
            return _742_p0;
          })(p0)));
          let _744_v70;
          _744_v70 = _dafny.Map.Empty.slice().updateUnsafe(_660_v1,(_739_v69)[_module.__default.safeIndex(new BigNumber((_738_v68).length), new BigNumber((_739_v69).length))]);
          let _745_v71;
          let _nw101 = Array((new BigNumber(22)).toNumber());
          _nw101[(_dafny.ZERO).toNumber()] = _738_v68;
          _nw101[(_dafny.ONE).toNumber()] = _738_v68;
          _nw101[(new BigNumber(2)).toNumber()] = ((_660_v1) ? (_738_v68) : (_738_v68));
          _nw101[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_738_v68, _738_v68);
          _nw101[(new BigNumber(4)).toNumber()] = _738_v68;
          _nw101[(new BigNumber(5)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_738_v68, _738_v68), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_738_v68, _738_v68)).length)), (((_668_v8).contains(p0)) ? ((_668_v8).get(p0)) : (new BigNumber((_662_v3).length))));
          _nw101[(new BigNumber(6)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(792)), ((_746_p0) => function (_747_i6) {
            return _746_p0;
          })(p0));
          _nw101[(new BigNumber(7)).toNumber()] = _738_v68;
          _nw101[(new BigNumber(8)).toNumber()] = _738_v68;
          _nw101[(new BigNumber(9)).toNumber()] = _738_v68;
          _nw101[(new BigNumber(10)).toNumber()] = _738_v68;
          _nw101[(new BigNumber(11)).toNumber()] = _dafny.Seq.update(_738_v68, _module.__default.safeIndex((_this).f28, new BigNumber((_738_v68).length)), (_this).f28);
          _nw101[(new BigNumber(12)).toNumber()] = _738_v68;
          _nw101[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_738_v68, _738_v68);
          _nw101[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(471)), ((_748_p0) => function (_749_i7) {
            return _748_p0;
          })(p0)), _738_v68);
          _nw101[(new BigNumber(15)).toNumber()] = _738_v68;
          _nw101[(new BigNumber(16)).toNumber()] = _738_v68;
          _nw101[(new BigNumber(17)).toNumber()] = (((_744_v70).contains(_module.__default.fm2(_736_v66, globalState))) ? ((_744_v70).get(_module.__default.fm2(_736_v66, globalState))) : (_738_v68));
          _nw101[(new BigNumber(18)).toNumber()] = _module.__default.fm21(globalState);
          _nw101[(new BigNumber(19)).toNumber()] = ((_660_v1) ? (_738_v68) : (_738_v68));
          _nw101[(new BigNumber(20)).toNumber()] = _dafny.Seq.Concat(_738_v68, _738_v68);
          _nw101[(new BigNumber(21)).toNumber()] = _738_v68;
          _745_v71 = _nw101;
          let _index103 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_745_v71).length));
          (_745_v71)[_index103] = (_739_v69)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f28), new BigNumber((_739_v69).length))];
          let _index104 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_745_v71).length));
          (_745_v71)[_index104] = _738_v68;
        }
        let _750_v72;
        let _init15 = function (_751_i10) {
          return _module.__default.safeDivisionInt(_751_i10, (_this).f28);
        };
        let _nw102 = Array((new BigNumber(23)).toNumber());
        for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw102.length); _i0_15++) {
          _nw102[_i0_15] = _init15(new BigNumber(_i0_15));
        }
        _750_v72 = _nw102;
        let _752_v73;
        _752_v73 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,_750_v72);
        r0 = (((_752_v73).contains(p0)) ? ((_752_v73).get(p0)) : (_750_v72));
        let _753_v74;
        let _754_v75;
        let _755_v76;
        let _756_v77;
        let _out4;
        let _out5;
        let _out6;
        let _out7;
        let _outcollector1 = (_this).m9(globalState);
        _out4 = _outcollector1[0];
        _out5 = _outcollector1[1];
        _out6 = _outcollector1[2];
        _out7 = _outcollector1[3];
        _753_v74 = _out4;
        _754_v75 = _out5;
        _755_v76 = _out6;
        _756_v77 = _out7;
        if ((_663_v4)[_module.__default.safeIndex(p0, new BigNumber((_663_v4).length))]) {
          _661_v2 = _661_v2;
          let _757_v78;
          let _nw103 = new _module.C0();
          _nw103.__ctor(_756_v77);
          _757_v78 = _nw103;
          _663_v4 = _dafny.Seq.of((_this).fm6(new BigNumber(-80), globalState), _660_v1);
          let _758_v79;
          let _nw104 = Array((new BigNumber(2)).toNumber());
          _nw104[(_dafny.ZERO).toNumber()] = (_659_v0).Difference(_659_v0);
          _nw104[(_dafny.ONE).toNumber()] = _659_v0;
          _758_v79 = _nw104;
          let _index105 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_758_v79).length));
          (_758_v79)[_index105] = _659_v0;
          let _759_v80;
          let _nw105 = Array((new BigNumber(12)).toNumber()).fill([]);
          _759_v80 = _nw105;
          let _index106 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_759_v80).length));
          (_759_v80)[_index106] = _750_v72;
          let _index107 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_758_v79).length));
          let _index108 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_759_v80).length));
          let _rhs84 = _659_v0;
          let _rhs85 = _750_v72;
          let _rhs86 = _755_v76;
          let _lhs85 = _758_v79;
          let _lhs86 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_758_v79).length));
          let _lhs87 = _759_v80;
          let _lhs88 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_759_v80).length));
          let _lhs89 = globalState;
          _lhs85[_lhs86] = _rhs84;
          _lhs87[_lhs88] = _rhs85;
          _lhs89.f2 = _rhs86;
          let _760_v81;
          _760_v81 = _dafny.Map.Empty.slice().updateUnsafe(_660_v1,p0);
          let _761_v82;
          _761_v82 = _dafny.MultiSet.fromElements((((_760_v81).contains(true)) ? ((_760_v81).get(true)) : (p0)));
          let _762_v83;
          _762_v83 = _dafny.MultiSet.fromElements(new BigNumber((_761_v82).cardinality()), (_dafny.ZERO).minus(p0));
          let _763_v84;
          _763_v84 = _dafny.Seq.of(_750_v72);
          let _764_v85;
          _764_v85 = _dafny.Seq.of((_763_v84)[_module.__default.safeIndex(p0, new BigNumber((_763_v84).length))]);
          let _index109 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_759_v80).length));
          let _rhs87 = _666_v7;
          let _rhs88 = (_module.__default.fm28((_758_v79)[_module.__default.safeIndex(new BigNumber(956), new BigNumber((_758_v79).length))], p0, _753_v74, true, globalState)).IsSubsetOf((_761_v82).update(new BigNumber((_662_v3).length), _module.__default.abs(new BigNumber((_dafny.Seq.UnicodeFromString("xnvuuooh")).length))));
          let _rhs89 = (_763_v84)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f28), new BigNumber((_763_v84).length))];
          let _lhs90 = globalState;
          let _lhs91 = _759_v80;
          let _lhs92 = _module.__default.safeIndex(new BigNumber(124), new BigNumber((_759_v80).length));
          _666_v7 = _rhs87;
          _lhs90.f2 = _rhs88;
          _lhs91[_lhs92] = _rhs89;
        } else {
          (globalState).f13 = ((!(!(_755_v76))) ? (((_dafny.ZERO).minus((_this).f28)).isLessThanOrEqualTo(new BigNumber(479))) : (_754_v75));
          let _765_v86;
          _765_v86 = _dafny.Seq.of(_718_v50, _718_v50);
          let _766_v87;
          let _nw106 = new _module.C1();
          _nw106.__ctor(p0, _765_v86);
          _766_v87 = _nw106;
          let _767_v88;
          _767_v88 = _dafny.Set.fromElements(_766_v87);
          let _768_v89;
          _768_v89 = _dafny.Seq.of(p0);
          let _769_v90;
          _769_v90 = _dafny.Seq.update(_dafny.Seq.update(_663_v4, _module.__default.safeIndex(new BigNumber((_668_v8).length), new BigNumber((_663_v4).length)), _753_v74), _module.__default.safeIndex((_this).f28, new BigNumber((_dafny.Seq.update(_663_v4, _module.__default.safeIndex(new BigNumber((_668_v8).length), new BigNumber((_663_v4).length)), _753_v74)).length)), _754_v75);
          let _770_v91;
          _770_v91 = _dafny.MultiSet.fromElements((p0).isLessThan(new BigNumber((_767_v88).length)), !((new BigNumber((_dafny.MultiSet.fromElements(p0)).cardinality())).isLessThanOrEqualTo((_this).fm17(!(_753_v74), _768_v89, p0, _769_v90, globalState))));
          let _771_v92;
          _771_v92 = _dafny.Map.Empty.slice().updateUnsafe(_660_v1,p0);
          (globalState).f16 = (((_770_v91).contains(_660_v1)) ? ((_770_v91).get(_660_v1)) : ((((_771_v92).contains(_753_v74)) ? ((_771_v92).get(_753_v74)) : (p0))));
          _666_v7 = _666_v7;
          let _index110 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((_750_v72).length));
          (_750_v72)[_index110] = _module.__default.fm0(new BigNumber((_662_v3).length), (_this).f28, p0, _755_v76, globalState);
          let _index111 = _module.__default.safeIndex(new BigNumber(526), new BigNumber((_750_v72).length));
          (_750_v72)[_index111] = _module.__default.safeDivisionInt(((!(_754_v75)) ? (new BigNumber((_662_v3).length)) : (new BigNumber((_module.__default.fm35(_755_v76, (_this).f28, globalState)).length))), new BigNumber((_662_v3).length));
          (globalState).f19 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.update(_662_v3, _module.__default.safeIndex((_this).f28, new BigNumber((_662_v3).length)), _661_v2), _662_v3), _dafny.Seq.Concat(_662_v3, _662_v3));
        }
        if (_756_v77) {
          _659_v0 = _659_v0;
          let _772_v93;
          _772_v93 = _dafny.Seq.of(_662_v3);
          (globalState).f20 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-857)), ((_773_v2) => function (_774_i11) {
            return _773_v2;
          })(_661_v2)), (_772_v93)[_module.__default.safeIndex((_this).f28, new BigNumber((_772_v93).length))]);
          let _775_v94;
          _775_v94 = _dafny.Seq.of(_718_v50);
          let _776_v95;
          let _nw107 = new _module.C1();
          _nw107.__ctor((_this).f28, _775_v94);
          _776_v95 = _nw107;
          _666_v7 = _666_v7;
          (globalState).f0 = true;
        } else {
          let _777_v96;
          _777_v96 = _dafny.Map.Empty.slice().updateUnsafe(_756_v77,_755_v76);
          let _778_v98;
          _778_v98 = _dafny.MultiSet.fromElements(_756_v77, _module.__default.fm2(_777_v96, globalState), _755_v76);
          let _779_v99;
          let _nw108 = new _module.C1();
          _nw108.__ctor(new BigNumber((_777_v96).length), _dafny.Seq.of(function () {
            let _coll36 = new _dafny.Map();
            for (const _compr_36 of (_dafny.MultiSet.fromElements((_dafny.ZERO).minus((((_778_v98).contains(_660_v1)) ? ((_778_v98).get(_660_v1)) : (p0))))).Elements) {
              let _780_v97 = _compr_36;
              if ((_dafny.MultiSet.fromElements((_dafny.ZERO).minus((((_778_v98).contains(_660_v1)) ? ((_778_v98).get(_660_v1)) : (p0))))).contains(_780_v97)) {
                _coll36.push([_module.__default.safeDivisionInt(_780_v97, (_this).f28),_660_v1]);
              }
            }
            return _coll36;
          }()));
          _779_v99 = _nw108;
          let _781_v100;
          _781_v100 = _dafny.Map.Empty.slice().updateUnsafe(!(_755_v76),_779_v99);
          _781_v100 = (_781_v100).Merge(_781_v100);
          let _rhs90 = p0;
          let _rhs91 = (_dafny.ZERO).minus(new BigNumber((_662_v3).length));
          let _lhs93 = globalState;
          let _lhs94 = globalState;
          _lhs93.f16 = _rhs90;
          _lhs94.f16 = _rhs91;
          _660_v1 = _756_v77;
          _754_v75 = _660_v1;
          let _782_v101;
          let _nw109 = new _module.C0();
          _nw109.__ctor(_dafny.Seq.IsProperPrefixOf(_662_v3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-673)), ((_783_v2) => function (_784_i12) {
            return _783_v2;
          })(_661_v2))));
          _782_v101 = _nw109;
        }
      }
      let _785_v102;
      let _nw110 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
      _785_v102 = _nw110;
      r0 = _785_v102;
      r1 = _660_v1;
      return [r0, r1];
    }
    m8(p0, p1, p2, globalState) {
      let _this = this;
      let _786_v0;
      _786_v0 = _dafny.Seq.of(p2, p0, p1, p0);
      _786_v0 = _dafny.Seq.Concat(_786_v0, _786_v0);
      let _787_v1;
      _787_v1 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f28);
      let _788_v2;
      _788_v2 = _dafny.Seq.UnicodeFromString("asmgvid");
      let _789_v3;
      let _nw111 = Array((new BigNumber(8)).toNumber());
      _nw111[(_dafny.ZERO).toNumber()] = (_this).f28;
      _nw111[(_dafny.ONE).toNumber()] = (_this).f28;
      _nw111[(new BigNumber(2)).toNumber()] = (((_787_v1).contains(p0)) ? ((_787_v1).get(p0)) : (new BigNumber((_788_v2).length)));
      _nw111[(new BigNumber(3)).toNumber()] = new BigNumber(-74);
      _nw111[(new BigNumber(4)).toNumber()] = (_this).f28;
      _nw111[(new BigNumber(5)).toNumber()] = (_this).f28;
      _nw111[(new BigNumber(6)).toNumber()] = (_this).f28;
      _nw111[(new BigNumber(7)).toNumber()] = (_this).f28;
      _789_v3 = _nw111;
      let _790_v4;
      _790_v4 = _dafny.Seq.of(_789_v3, _789_v3);
      let _791_v5;
      _791_v5 = _module.D6.create_DC12(_790_v4);
      let _pat_let_tv14 = _790_v4;
      let _792_v6;
      _792_v6 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let10_0) {
        return function (_793_dt__update__tmp_h0) {
          return function (_pat_let11_0) {
            return function (_794_dt__update_hcf18_h0) {
              return _module.D6.create_DC12(_794_dt__update_hcf18_h0);
            }(_pat_let11_0);
          }(_pat_let_tv14);
        }(_pat_let10_0);
      }(_791_v5),(_this).f28);
      let _795_v7;
      _795_v7 = _dafny.Seq.of(_792_v6);
      let _796_v8;
      _796_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(177),p1);
      let _797_v9;
      _797_v9 = _dafny.Seq.of(new BigNumber(-739), (_this).f28, (_this).f28);
      let _798_v10;
      let _nw112 = Array((new BigNumber(19)).toNumber());
      _nw112[(_dafny.ZERO).toNumber()] = _792_v6;
      _nw112[(_dafny.ONE).toNumber()] = _792_v6;
      _nw112[(new BigNumber(2)).toNumber()] = ((_795_v7)[_module.__default.safeIndex((_this).f28, new BigNumber((_795_v7).length))]).Merge(_792_v6);
      _nw112[(new BigNumber(3)).toNumber()] = _792_v6;
      _nw112[(new BigNumber(4)).toNumber()] = _792_v6;
      _nw112[(new BigNumber(5)).toNumber()] = (_792_v6).update(_791_v5, (_this).f28);
      _nw112[(new BigNumber(6)).toNumber()] = (_792_v6).Merge(_792_v6);
      _nw112[(new BigNumber(7)).toNumber()] = ((_795_v7)[_module.__default.safeIndex(new BigNumber((_796_v8).length), new BigNumber((_795_v7).length))]).Merge(_792_v6);
      _nw112[(new BigNumber(8)).toNumber()] = (_792_v6).Merge(_792_v6);
      _nw112[(new BigNumber(9)).toNumber()] = _792_v6;
      _nw112[(new BigNumber(10)).toNumber()] = _792_v6;
      _nw112[(new BigNumber(11)).toNumber()] = (_792_v6).update(_791_v5, (_this).f28);
      _nw112[(new BigNumber(12)).toNumber()] = _792_v6;
      _nw112[(new BigNumber(13)).toNumber()] = (_792_v6).update(_791_v5, new BigNumber((_797_v9).length));
      _nw112[(new BigNumber(14)).toNumber()] = _792_v6;
      _nw112[(new BigNumber(15)).toNumber()] = (_792_v6).Merge(_dafny.Map.Empty.slice().updateUnsafe(_791_v5,new BigNumber((_786_v0).length)));
      _nw112[(new BigNumber(16)).toNumber()] = _792_v6;
      _nw112[(new BigNumber(17)).toNumber()] = (_792_v6).Merge(_792_v6);
      _nw112[(new BigNumber(18)).toNumber()] = _792_v6;
      _798_v10 = _nw112;
      _798_v10 = _798_v10;
      let _799_v11;
      let _nw113 = Array((new BigNumber(21)).toNumber());
      _nw113[(_dafny.ZERO).toNumber()] = p2;
      _nw113[(_dafny.ONE).toNumber()] = p1;
      _nw113[(new BigNumber(2)).toNumber()] = ((_dafny.ZERO).minus((_this).f28)).isLessThanOrEqualTo(new BigNumber(511));
      _nw113[(new BigNumber(3)).toNumber()] = p0;
      _nw113[(new BigNumber(4)).toNumber()] = p0;
      _nw113[(new BigNumber(5)).toNumber()] = !(true);
      _nw113[(new BigNumber(6)).toNumber()] = p0;
      _nw113[(new BigNumber(7)).toNumber()] = p1;
      _nw113[(new BigNumber(8)).toNumber()] = p0;
      _nw113[(new BigNumber(9)).toNumber()] = p1;
      _nw113[(new BigNumber(10)).toNumber()] = p1;
      _nw113[(new BigNumber(11)).toNumber()] = p2;
      _nw113[(new BigNumber(12)).toNumber()] = p1;
      _nw113[(new BigNumber(13)).toNumber()] = p1;
      _nw113[(new BigNumber(14)).toNumber()] = p2;
      _nw113[(new BigNumber(15)).toNumber()] = p0;
      _nw113[(new BigNumber(16)).toNumber()] = true;
      _nw113[(new BigNumber(17)).toNumber()] = ((p2) ? (!(p2)) : (p1));
      _nw113[(new BigNumber(18)).toNumber()] = !(!(p2));
      _nw113[(new BigNumber(19)).toNumber()] = p2;
      _nw113[(new BigNumber(20)).toNumber()] = p2;
      _799_v11 = _nw113;
      for (const _guard_loop_6 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_799_v11).length))) {
        let _800_i0 = _guard_loop_6;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_800_i0)) && ((_800_i0).isLessThan(new BigNumber((_799_v11).length))))) {
          (_799_v11)[(_800_i0)] = !(((_this).f28).isLessThan((_this).f28));
        }
      }
      let _801_v12;
      _801_v12 = _module.D12.create_DC23((_this).f28);
      let _802_v13;
      _802_v13 = _module.D12.create_DC24(_801_v12);
      let _803_v14;
      _803_v14 = _module.D12.create_DC24(_802_v13);
      let _804_i1;
      _804_i1 = _dafny.ZERO;
      L4: {
        let _pat_let_tv15 = p1;
        let _pat_let_tv16 = p2;
        let _pat_let_tv17 = p0;
        while (function (_source14) {
          if (_source14.is_DC23) {
            let _820___mcc_h0 = (_source14).cf30;
            let _821_cf30 = _820___mcc_h0;
            return _pat_let_tv15;
          } else if (_source14.is_DC22) {
            let _822___mcc_h1 = (_source14).cf29;
            let _823_cf29 = _822___mcc_h1;
            return _pat_let_tv16;
          } else {
            let _824___mcc_h2 = (_source14).cf31;
            let _825_cf31 = _824___mcc_h2;
            return !(_pat_let_tv17);
          }
        }(_803_v14)) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_804_i1)) {
              break L4;
            }
            _804_i1 = (_804_i1).plus(_dafny.ONE);
            (globalState).f21 = p0;
            if (p1) {
              (globalState).f7 = ((_this).f28).plus(((_this).f28).minus((_this).f28));
              let _index112 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_799_v11).length));
              (_799_v11)[_index112] = p0;
              let _805_v15;
              _805_v15 = _module.D1.create_DC3(p2);
              let _index113 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_799_v11).length));
              (_799_v11)[_index113] = (_805_v15).dtor_cf3;
              (globalState).f21 = false;
              let _index114 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_789_v3).length));
              (_789_v3)[_index114] = _module.__default.safeDivisionInt((_this).f28, (_this).f28);
              let _index115 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_789_v3).length));
              (_789_v3)[_index115] = _module.__default.safeDivisionInt(((_this).f28).multipliedBy((_this).f28), ((_dafny.ZERO).minus((_this).f28)).plus(new BigNumber((_787_v1).length)));
              (globalState).f21 = false;
            } else {
              (globalState).f16 = (_this).f28;
              let _806_v16;
              let _nw114 = new _module.C0();
              _nw114.__ctor(p0);
              _806_v16 = _nw114;
              (globalState).f7 = _module.__default.safeModuloInt((_this).f28, (_this).f28);
              let _807_v17;
              _807_v17 = _dafny.Set.fromElements((_this).f28);
              let _808_v18;
              _808_v18 = _dafny.Map.Empty.slice().updateUnsafe(_807_v17,_dafny.Set.fromElements(p2, p0));
              let _809_v19;
              _809_v19 = _dafny.Set.fromElements(_806_v16.f27, false, p1, p2, p1);
              _808_v18 = (_808_v18).update(_807_v17, _809_v19);
              let _810_v20;
              let _nw115 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.of());
              _810_v20 = _nw115;
              let _index116 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_810_v20).length));
              (_810_v20)[_index116] = _786_v0;
              let _index117 = _module.__default.safeIndex(new BigNumber(965), new BigNumber((_810_v20).length));
              (_810_v20)[_index117] = _dafny.Seq.Concat(_786_v0, _786_v0);
            }
            let _811_v21;
            let _init16 = ((_812_p0, _813_p1) => function (_814_i2) {
              return _dafny.Set.fromElements(_812_p0, _813_p1, !(_812_p0), _813_p1, _812_p0);
            })(p0, p1);
            let _nw116 = Array((new BigNumber(11)).toNumber());
            for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw116.length); _i0_16++) {
              _nw116[_i0_16] = _init16(new BigNumber(_i0_16));
            }
            _811_v21 = _nw116;
            let _815_v22;
            _815_v22 = _module.D7.create_DC16(_811_v21);
            _815_v22 = _module.D7.create_DC16(_811_v21);
            let _816_v23;
            _816_v23 = new _dafny.CodePoint('b'.codePointAt(0));
            let _817_v24;
            _817_v24 = _dafny.Map.Empty.slice().updateUnsafe(p2,false);
            let _818_v25;
            _818_v25 = _dafny.Set.fromElements(p2);
            let _819_v26;
            _819_v26 = _dafny.Map.Empty.slice().updateUnsafe(p2,_818_v25);
            let _rhs92 = _816_v23;
            let _rhs93 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt((_this).f28, _module.__default.fm0(new BigNumber((_817_v24).length), _module.__default.fm0((_this).f28, new BigNumber(350), new BigNumber(((((_819_v26).contains(p2)) ? ((_819_v26).get(p2)) : (_818_v25))).length), p1, globalState), (_this).f28, p0, globalState))));
            let _rhs94 = p2;
            let _rhs95 = (_this).f28;
            let _lhs95 = globalState;
            let _lhs96 = globalState;
            let _lhs97 = globalState;
            _816_v23 = _rhs92;
            _lhs95.f7 = _rhs93;
            _lhs96.f2 = _rhs94;
            _lhs97.f16 = _rhs95;
          }
        }
      }
      let _826_v27;
      let _init17 = ((_827_p2, _828_p0) => function (_829_i3) {
        return _dafny.MultiSet.fromElements(!(_827_p2), _827_p2, _828_p0);
      })(p2, p0);
      let _nw117 = Array((new BigNumber(28)).toNumber());
      for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw117.length); _i0_17++) {
        _nw117[_i0_17] = _init17(new BigNumber(_i0_17));
      }
      _826_v27 = _nw117;
      let _830_v28;
      _830_v28 = _dafny.MultiSet.fromElements(p2);
      let _index118 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_826_v27).length));
      (_826_v27)[_index118] = _830_v28;
      let _831_v29;
      _831_v29 = _dafny.Map.Empty.slice().updateUnsafe(p2,p1);
      let _832_v30;
      _832_v30 = _dafny.MultiSet.fromElements(_module.__default.fm34(globalState));
      let _833_v31;
      _833_v31 = _dafny.Seq.of(_832_v30);
      let _834_v32;
      _834_v32 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,_788_v2);
      let _835_v33;
      _835_v33 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_834_v32).length),new BigNumber((_788_v2).length));
      let _836_v34;
      _836_v34 = _dafny.Set.fromElements((_this).f28, (_this).f28, (new BigNumber((_796_v8).length)).plus(new BigNumber((_833_v31).length)), _module.__default.fm0(new BigNumber((_835_v33).length), (_this).f28, (_this).f28, p0, globalState));
      let _index119 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_826_v27).length));
      let _rhs96 = _830_v28;
      let _rhs97 = (_831_v29).Merge(_831_v29);
      let _rhs98 = (_836_v34).Intersect(_dafny.Set.fromElements((_this).f28, (_this).f28));
      let _lhs98 = _826_v27;
      let _lhs99 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_826_v27).length));
      _lhs98[_lhs99] = _rhs96;
      _831_v29 = _rhs97;
      _836_v34 = _rhs98;
      (globalState).f2 = (_this).fm7(((_this).f28).isLessThanOrEqualTo((_this).f28), (_this).f28, p1, globalState);
      return;
    }
    m9(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = false;
      let r3 = false;
      let _837_v0;
      _837_v0 = _dafny.Seq.of((_this).f28);
      let _838_i0;
      _838_i0 = _dafny.ZERO;
      L5: {
        while (!(_dafny.areEqual(_dafny.Seq.Concat(_837_v0, _dafny.Seq.update(_837_v0, _module.__default.safeIndex((_this).f28, new BigNumber((_837_v0).length)), (_this).f28)), _module.__default.fm37(globalState)))) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_838_i0)) {
              break L5;
            }
            _838_i0 = (_838_i0).plus(_dafny.ONE);
            let _839_v1;
            let _init18 = function (_840_i1) {
              return new _dafny.CodePoint('o'.codePointAt(0));
            };
            let _nw118 = Array((new BigNumber(24)).toNumber());
            for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw118.length); _i0_18++) {
              _nw118[_i0_18] = _init18(new BigNumber(_i0_18));
            }
            _839_v1 = _nw118;
            let _841_v2;
            _841_v2 = new _dafny.CodePoint('h'.codePointAt(0));
            let _index120 = _module.__default.safeIndex(new BigNumber(121), new BigNumber((_839_v1).length));
            (_839_v1)[_index120] = _841_v2;
            let _842_v3;
            _842_v3 = false;
            let _index121 = _module.__default.safeIndex(new BigNumber(121), new BigNumber((_839_v1).length));
            (_839_v1)[_index121] = _module.__default.fm27(((_this).f28).isEqualTo((_this).f28), _842_v3, globalState);
            let _843_v4;
            _843_v4 = _dafny.Map.Empty.slice().updateUnsafe(_842_v3,_842_v3);
            let _844_v5;
            _844_v5 = _dafny.Seq.of((((_843_v4).contains(_842_v3)) ? ((_843_v4).get(_842_v3)) : (_842_v3)), _842_v3);
            let _845_v6;
            _845_v6 = _844_v5;
            let _846_v7;
            _846_v7 = _dafny.Map.Empty.slice().updateUnsafe(_842_v3,_844_v5);
            let _847_v8;
            _847_v8 = _module.D14.create_DC29((_this).fm17(_842_v3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(884)), ((_848_v3) => function (_849_i2) {
  return new BigNumber((_dafny.Seq.of(_848_v3, false)).length);
})(_842_v3)), (_this).f28, _845_v6, globalState), _846_v7);
            (globalState).f16 = (_dafny.ZERO).minus((_847_v8).dtor_cf38);
            let _source15 = _module.D14.create_DC30(_837_v0, _dafny.Seq.update(_844_v5, _module.__default.safeIndex((_this).f28, new BigNumber((_844_v5).length)), !(_842_v3)));
            if (_source15.is_DC29) {
              let _850___mcc_h0 = (_source15).cf38;
              let _851___mcc_h1 = (_source15).cf39;
              let _852_cf39 = _851___mcc_h1;
              let _853_cf38 = _850___mcc_h0;
              let _854_v9;
              _854_v9 = _dafny.MultiSet.fromElements(!(_842_v3) || (_842_v3));
              let _855_v10;
              _855_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_this).f28);
              let _856_v11;
              _856_v11 = _module.D13.create_DC25(_855_v10);
              let _857_v12;
              _857_v12 = _dafny.Set.fromElements(_856_v11);
              let _858_v13;
              _858_v13 = _dafny.Map.Empty.slice().updateUnsafe(_857_v12,_841_v2);
              let _rhs99 = (_this).f28;
              let _rhs100 = new BigNumber(209);
              let _rhs101 = (((_854_v9).contains(_842_v3)) ? ((_854_v9).get(_842_v3)) : (new BigNumber(((_858_v13).update(_857_v12, _841_v2)).length)));
              let _rhs102 = (_dafny.ZERO).minus(_module.__default.fm0((_853_cf38).plus((_this).f28), new BigNumber(-485), ((!(_842_v3)) ? (new BigNumber((_dafny.MultiSet.FromArray(_837_v0)).cardinality())) : (_853_cf38)), _842_v3, globalState));
              let _lhs100 = globalState;
              let _lhs101 = globalState;
              let _lhs102 = globalState;
              _lhs100.f16 = _rhs99;
              _853_cf38 = _rhs100;
              _lhs101.f7 = _rhs101;
              _lhs102.f7 = _rhs102;
              let _859_v14;
              let _init19 = ((_860_v9) => function (_861_i3) {
                return (_861_i3).multipliedBy(new BigNumber((_860_v9).cardinality()));
              })(_854_v9);
              let _nw119 = Array((new BigNumber(11)).toNumber());
              for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw119.length); _i0_19++) {
                _nw119[_i0_19] = _init19(new BigNumber(_i0_19));
              }
              _859_v14 = _nw119;
              let _862_v15;
              _862_v15 = _dafny.MultiSet.fromElements(_853_cf38);
              let _index122 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_859_v14).length));
              (_859_v14)[_index122] = (_853_cf38).minus(new BigNumber((_862_v15).cardinality()));
              let _index123 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_859_v14).length));
              (_859_v14)[_index123] = _module.__default.safeModuloInt(_853_cf38, _853_cf38);
              _837_v0 = _837_v0;
              let _863_v16;
              _863_v16 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.update(_837_v0, _module.__default.safeIndex((_859_v14)[_module.__default.safeIndex(new BigNumber(93), new BigNumber((_859_v14).length))], new BigNumber((_837_v0).length)), _853_cf38), _module.__default.safeIndex((_859_v14)[_module.__default.safeIndex(new BigNumber(93), new BigNumber((_859_v14).length))], new BigNumber((_dafny.Seq.update(_837_v0, _module.__default.safeIndex((_859_v14)[_module.__default.safeIndex(new BigNumber(93), new BigNumber((_859_v14).length))], new BigNumber((_837_v0).length)), _853_cf38)).length)), (_859_v14)[_module.__default.safeIndex(new BigNumber(93), new BigNumber((_859_v14).length))]), _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber(-205)), new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_837_v0, _module.__default.safeIndex((_859_v14)[_module.__default.safeIndex(new BigNumber(93), new BigNumber((_859_v14).length))], new BigNumber((_837_v0).length)), _853_cf38), _module.__default.safeIndex((_859_v14)[_module.__default.safeIndex(new BigNumber(93), new BigNumber((_859_v14).length))], new BigNumber((_dafny.Seq.update(_837_v0, _module.__default.safeIndex((_859_v14)[_module.__default.safeIndex(new BigNumber(93), new BigNumber((_859_v14).length))], new BigNumber((_837_v0).length)), _853_cf38)).length)), (_859_v14)[_module.__default.safeIndex(new BigNumber(93), new BigNumber((_859_v14).length))])).length)), (_this).f28),_844_v5);
              let _864_v17;
              _864_v17 = _dafny.Seq.UnicodeFromString("gsaqxlayd");
              let _865_v18;
              _865_v18 = _dafny.Seq.of(_863_v16, _863_v16, _863_v16, ((_863_v16).update(_837_v0, _844_v5)).update(_837_v0, _844_v5));
              let _866_v19;
              _866_v19 = _dafny.Set.fromElements(_842_v3);
              let _867_v21;
              _867_v21 = _module.D4.create_DC9(new BigNumber((_837_v0).length));
              let _868_v22;
              _868_v22 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm21(globalState),_867_v21);
              let _869_v23;
              _869_v23 = _module.D16.create_DC35(_863_v16);
              let _870_v25;
              _870_v25 = _dafny.Seq.of(_837_v0);
              let _871_v28;
              _871_v28 = _dafny.MultiSet.fromElements(_837_v0);
              let _pat_let_tv18 = _863_v16;
              let _872_v29;
              let _nw120 = Array((new BigNumber(27)).toNumber());
              _nw120[(_dafny.ZERO).toNumber()] = _863_v16;
              _nw120[(_dafny.ONE).toNumber()] = _863_v16;
              _nw120[(new BigNumber(2)).toNumber()] = _863_v16;
              _nw120[(new BigNumber(3)).toNumber()] = _863_v16;
              _nw120[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_837_v0,_dafny.Seq.of(_842_v3, false, _842_v3));
              _nw120[(new BigNumber(5)).toNumber()] = ((_863_v16).update(_dafny.Seq.of(new BigNumber((_864_v17).length), new BigNumber(205)), _844_v5)).Merge(_863_v16);
              _nw120[(new BigNumber(6)).toNumber()] = _863_v16;
              _nw120[(new BigNumber(7)).toNumber()] = (_865_v18)[_module.__default.safeIndex(new BigNumber((_866_v19).length), new BigNumber((_865_v18).length))];
              _nw120[(new BigNumber(8)).toNumber()] = _863_v16;
              _nw120[(new BigNumber(9)).toNumber()] = _863_v16;
              _nw120[(new BigNumber(10)).toNumber()] = _863_v16;
              _nw120[(new BigNumber(11)).toNumber()] = _863_v16;
              _nw120[(new BigNumber(12)).toNumber()] = _863_v16;
              _nw120[(new BigNumber(13)).toNumber()] = _module.__default.fm39(new BigNumber((_854_v9).cardinality()), _842_v3, (_this).f28, globalState);
              _nw120[(new BigNumber(14)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_837_v0,_844_v5);
              _nw120[(new BigNumber(15)).toNumber()] = _863_v16;
              _nw120[(new BigNumber(16)).toNumber()] = (function () {
                let _coll37 = new _dafny.Map();
                for (const _compr_37 of (_868_v22).Keys.Elements) {
                  let _873_v20 = _compr_37;
                  if ((_868_v22).contains(_873_v20)) {
                    _coll37.push([_873_v20,_dafny.Seq.of(_842_v3)]);
                  }
                }
                return _coll37;
              }()).Merge(_863_v16);
              _nw120[(new BigNumber(17)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_837_v0,_844_v5);
              _nw120[(new BigNumber(18)).toNumber()] = (function (_pat_let12_0) {
                return function (_874_dt__update__tmp_h0) {
                  return function (_pat_let13_0) {
                    return function (_875_dt__update_hcf46_h0) {
                      return _module.D16.create_DC35(_875_dt__update_hcf46_h0);
                    }(_pat_let13_0);
                  }(_pat_let_tv18);
                }(_pat_let12_0);
              }(_869_v23)).dtor_cf46;
              _nw120[(new BigNumber(19)).toNumber()] = _863_v16;
              _nw120[(new BigNumber(20)).toNumber()] = (function () {
                let _coll38 = new _dafny.Map();
                for (const _compr_38 of (_870_v25).Elements) {
                  let _876_v24 = _compr_38;
                  if (_dafny.Seq.contains(_870_v25, _876_v24)) {
                    _coll38.push([_876_v24,_844_v5]);
                  }
                }
                return _coll38;
              }()).Merge(_863_v16);
              _nw120[(new BigNumber(21)).toNumber()] = _863_v16;
              _nw120[(new BigNumber(22)).toNumber()] = _863_v16;
              _nw120[(new BigNumber(23)).toNumber()] = _863_v16;
              _nw120[(new BigNumber(24)).toNumber()] = _863_v16;
              _nw120[(new BigNumber(25)).toNumber()] = (function () {
                let _coll39 = new _dafny.Map();
                for (const _compr_39 of (_870_v25).Elements) {
                  let _877_v26 = _compr_39;
                  if (_dafny.Seq.contains(_870_v25, _877_v26)) {
                    _coll39.push([_877_v26,_844_v5]);
                  }
                }
                return _coll39;
              }()).update(_837_v0, _844_v5);
              _nw120[(new BigNumber(26)).toNumber()] = (function () {
                let _coll40 = new _dafny.Map();
                for (const _compr_40 of (_871_v28).Elements) {
                  let _878_v27 = _compr_40;
                  if ((_871_v28).contains(_878_v27)) {
                    _coll40.push([_878_v27,_dafny.Seq.of(_842_v3, false)]);
                  }
                }
                return _coll40;
              }()).Merge(_863_v16);
              _872_v29 = _nw120;
              let _index124 = _module.__default.safeIndex(new BigNumber(867), new BigNumber((_872_v29).length));
              (_872_v29)[_index124] = _863_v16;
              let _879_v30;
              let _nw121 = new _module.C0();
              _nw121.__ctor(false);
              _879_v30 = _nw121;
              let _index125 = _module.__default.safeIndex(new BigNumber(867), new BigNumber((_872_v29).length));
              let _index126 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_859_v14).length));
              let _rhs103 = (_869_v23).dtor_cf46;
              let _rhs104 = _841_v2;
              let _rhs105 = (_859_v14)[_module.__default.safeIndex(new BigNumber(93), new BigNumber((_859_v14).length))];
              let _rhs106 = !((_this).f28).isEqualTo(_module.__default.safeDivisionInt((_this).f28, new BigNumber((_864_v17).length)));
              let _rhs107 = _879_v30;
              let _lhs103 = _872_v29;
              let _lhs104 = _module.__default.safeIndex(new BigNumber(867), new BigNumber((_872_v29).length));
              let _lhs105 = _859_v14;
              let _lhs106 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_859_v14).length));
              _lhs103[_lhs104] = _rhs103;
              _841_v2 = _rhs104;
              _lhs105[_lhs106] = _rhs105;
              r1 = _rhs106;
              _879_v30 = _rhs107;
            } else if (_source15.is_DC30) {
              let _880___mcc_h2 = (_source15).cf40;
              let _881___mcc_h3 = (_source15).cf41;
              let _882_cf41 = _881___mcc_h3;
              let _883_cf40 = _880___mcc_h2;
              let _884_v31;
              _884_v31 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_882_cf41)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f28), new BigNumber((_882_cf41).length))], (_this).fm6((_this).f28, globalState)),_dafny.Seq.Create(_module.__default.abs(new BigNumber(375)), ((_885_v1) => function (_886_i4) {
                return (_885_v1)[_module.__default.safeIndex(new BigNumber(121), new BigNumber((_885_v1).length))];
              })(_839_v1)));
              let _887_v32;
              _887_v32 = _dafny.Seq.UnicodeFromString("bfsvcv");
              let _888_v33;
              _888_v33 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_this).f28);
              _module.__default.m0(!(_842_v3) || (_842_v3), new BigNumber(-369), (_884_v31).Merge((_884_v31).update(_882_cf41, _887_v32)), (new BigNumber(811)).minus(new BigNumber((_888_v33).length)), globalState);
              (globalState).f7 = (_dafny.ZERO).minus((_this).f28);
              (globalState).f16 = (_this).f28;
              let _889_v34;
              let _nw122 = Array((new BigNumber(2)).toNumber()).fill(false);
              _889_v34 = _nw122;
              let _890_v35;
              let _nw123 = Array((new BigNumber(18)).toNumber());
              _nw123[(_dafny.ZERO).toNumber()] = _889_v34;
              _nw123[(_dafny.ONE).toNumber()] = ((_842_v3) ? (_889_v34) : (_889_v34));
              _nw123[(new BigNumber(2)).toNumber()] = _889_v34;
              _nw123[(new BigNumber(3)).toNumber()] = _889_v34;
              _nw123[(new BigNumber(4)).toNumber()] = _889_v34;
              _nw123[(new BigNumber(5)).toNumber()] = _889_v34;
              _nw123[(new BigNumber(6)).toNumber()] = _889_v34;
              _nw123[(new BigNumber(7)).toNumber()] = _889_v34;
              _nw123[(new BigNumber(8)).toNumber()] = _889_v34;
              _nw123[(new BigNumber(9)).toNumber()] = _889_v34;
              _nw123[(new BigNumber(10)).toNumber()] = _889_v34;
              _nw123[(new BigNumber(11)).toNumber()] = _889_v34;
              _nw123[(new BigNumber(12)).toNumber()] = ((_842_v3) ? (_889_v34) : (_889_v34));
              _nw123[(new BigNumber(13)).toNumber()] = _889_v34;
              _nw123[(new BigNumber(14)).toNumber()] = _889_v34;
              _nw123[(new BigNumber(15)).toNumber()] = _889_v34;
              _nw123[(new BigNumber(16)).toNumber()] = _889_v34;
              _nw123[(new BigNumber(17)).toNumber()] = _889_v34;
              _890_v35 = _nw123;
              let _index127 = _module.__default.safeIndex(new BigNumber(926), new BigNumber((_890_v35).length));
              (_890_v35)[_index127] = _889_v34;
              let _index128 = _module.__default.safeIndex(new BigNumber(926), new BigNumber((_890_v35).length));
              (_890_v35)[_index128] = _889_v34;
            } else if (_source15.is_DC28) {
              let _891___mcc_h4 = (_source15).cf37;
              let _892_cf37 = _891___mcc_h4;
              (globalState).f7 = ((_this).f28).multipliedBy(_module.__default.safeModuloInt((_this).f28, new BigNumber(300)));
              let _893_v36;
              _893_v36 = _dafny.Set.fromElements(_842_v3);
              let _894_v37;
              let _nw124 = Array((new BigNumber(2)).toNumber()).fill(false);
              _894_v37 = _nw124;
              let _895_v38;
              _895_v38 = _dafny.Map.Empty.slice().updateUnsafe(_894_v37,_893_v36);
              let _896_v39;
              let _nw125 = Array((new BigNumber(16)).toNumber());
              _nw125[(_dafny.ZERO).toNumber()] = (_893_v36).Difference(_893_v36);
              _nw125[(_dafny.ONE).toNumber()] = _893_v36;
              _nw125[(new BigNumber(2)).toNumber()] = _893_v36;
              _nw125[(new BigNumber(3)).toNumber()] = (_893_v36).Difference(_893_v36);
              _nw125[(new BigNumber(4)).toNumber()] = _893_v36;
              _nw125[(new BigNumber(5)).toNumber()] = _dafny.Set.fromElements((((_843_v4).contains(_842_v3)) ? ((_843_v4).get(_842_v3)) : (_842_v3)));
              _nw125[(new BigNumber(6)).toNumber()] = _module.__default.fm19(new BigNumber(365), _842_v3, _842_v3, (_this).f28, globalState);
              _nw125[(new BigNumber(7)).toNumber()] = (_893_v36).Difference(_893_v36);
              _nw125[(new BigNumber(8)).toNumber()] = ((_842_v3) ? (_893_v36) : (_893_v36));
              _nw125[(new BigNumber(9)).toNumber()] = _893_v36;
              _nw125[(new BigNumber(10)).toNumber()] = _dafny.Set.fromElements(_842_v3, (_this).fm7(_842_v3, (_this).f28, _842_v3, globalState));
              _nw125[(new BigNumber(11)).toNumber()] = _893_v36;
              _nw125[(new BigNumber(12)).toNumber()] = (_893_v36).Union(_893_v36);
              _nw125[(new BigNumber(13)).toNumber()] = _893_v36;
              _nw125[(new BigNumber(14)).toNumber()] = (_893_v36).Difference(_893_v36);
              _nw125[(new BigNumber(15)).toNumber()] = (((_895_v38).contains(_894_v37)) ? ((_895_v38).get(_894_v37)) : (_893_v36));
              _896_v39 = _nw125;
              let _index129 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_896_v39).length));
              (_896_v39)[_index129] = (_893_v36).Intersect(_893_v36);
              let _index130 = _module.__default.safeIndex(new BigNumber(400), new BigNumber((_896_v39).length));
              (_896_v39)[_index130] = _893_v36;
              let _897_v40;
              _897_v40 = _dafny.Seq.of(_837_v0);
              let _898_v41;
              let _nw126 = Array((new BigNumber(4)).toNumber());
              _nw126[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_837_v0, _dafny.Seq.of(new BigNumber((_892_cf37).length), (_this).f28, (_this).f28));
              _nw126[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_837_v0, _837_v0);
              _nw126[(new BigNumber(2)).toNumber()] = _837_v0;
              _nw126[(new BigNumber(3)).toNumber()] = (_897_v40)[_module.__default.safeIndex((_this).f28, new BigNumber((_897_v40).length))];
              _898_v41 = _nw126;
              let _index131 = _module.__default.safeIndex(new BigNumber(791), new BigNumber((_898_v41).length));
              (_898_v41)[_index131] = _dafny.Seq.Concat(_module.__default.fm21(globalState), _837_v0);
              let _index132 = _module.__default.safeIndex(new BigNumber(791), new BigNumber((_898_v41).length));
              (_898_v41)[_index132] = _837_v0;
              let _899_v42;
              let _nw127 = new _module.C0();
              _nw127.__ctor((new BigNumber(78)).isEqualTo((_this).f28));
              _899_v42 = _nw127;
              _899_v42 = _899_v42;
            } else {
              let _900___mcc_h5 = (_source15).cf42;
              let _901_cf42 = _900___mcc_h5;
              let _902_v43;
              _902_v43 = _module.D15.create_DC33(_837_v0);
              let _903_v44;
              _903_v44 = _module.D7.create_DC15((_902_v43).dtor_cf44);
              let _904_v45;
              _904_v45 = _module.D7.create_DC17(_903_v44);
              _904_v45 = _904_v45;
              let _905_v46;
              _905_v46 = _dafny.Map.Empty.slice().updateUnsafe(_842_v3,(_dafny.ZERO).minus((_this).f28));
              let _rhs108 = (_837_v0)[_module.__default.safeIndex((_this).f28, new BigNumber((_837_v0).length))];
              let _rhs109 = (((_905_v46).contains(true)) ? ((_905_v46).get(true)) : ((_this).f28));
              let _lhs107 = globalState;
              let _lhs108 = globalState;
              _lhs107.f7 = _rhs108;
              _lhs108.f16 = _rhs109;
              let _906_v47;
              _906_v47 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_this).f28),_844_v5);
              let _907_v48;
              _907_v48 = _module.D16.create_DC35(_906_v47);
              let _908_v49;
              _908_v49 = _module.D16.create_DC37(_907_v48);
              _908_v49 = _module.D16.create_DC37(_907_v48);
              let _909_v50;
              _909_v50 = _module.D12.create_DC22(_905_v46);
              let _910_v51;
              _910_v51 = _dafny.Set.fromElements(_909_v50);
              let _911_v52;
              _911_v52 = _dafny.Map.Empty.slice().updateUnsafe(_909_v50,_842_v3);
              _910_v51 = ((_dafny.Set.fromElements(_module.D12.create_DC22(_905_v46))).Union(_910_v51)).Difference((function () {
                let _coll41 = new _dafny.Set();
                for (const _compr_41 of (_911_v52).Keys.Elements) {
                  let _912_v53 = _compr_41;
                  if ((_911_v52).contains(_912_v53)) {
                    _coll41.add(_912_v53);
                  }
                }
                return _coll41;
              }()).Difference(_910_v51));
            }
            let _913_v55;
            _913_v55 = _module.D16.create_DC37(_module.D16.create_DC35(function () {
  let _coll42 = new _dafny.Map();
  for (const _compr_42 of (_dafny.Map.Empty.slice().updateUnsafe(_837_v0,_842_v3)).Keys.Elements) {
    let _914_v54 = _compr_42;
    if ((_dafny.Map.Empty.slice().updateUnsafe(_837_v0,_842_v3)).contains(_914_v54)) {
      _coll42.push([_914_v54,_844_v5]);
    }
  }
  return _coll42;
}()));
            let _915_v56;
            _915_v56 = _dafny.MultiSet.fromElements(_913_v55, _913_v55, _913_v55);
            let _916_v57;
            _916_v57 = _dafny.Map.Empty.slice().updateUnsafe(_842_v3,(_915_v56).Union(_915_v56));
            let _917_v58;
            let _init20 = ((_918_v3) => function (_919_i5) {
              return _918_v3;
            })(_842_v3);
            let _nw128 = Array((new BigNumber(8)).toNumber());
            for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw128.length); _i0_20++) {
              _nw128[_i0_20] = _init20(new BigNumber(_i0_20));
            }
            _917_v58 = _nw128;
            let _index133 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_917_v58).length));
            (_917_v58)[_index133] = !(!(!(_842_v3)));
            let _920_v59;
            _920_v59 = _dafny.Seq.UnicodeFromString("ji");
            let _921_v60;
            _921_v60 = _dafny.Map.Empty.slice().updateUnsafe(_842_v3,(_this).f28);
            let _922_v61;
            let _nw129 = Array((new BigNumber(10)).toNumber());
            _nw129[(_dafny.ZERO).toNumber()] = _920_v59;
            _nw129[(_dafny.ONE).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(947)), function (_923_i6) {
              return new _dafny.CodePoint('h'.codePointAt(0));
            });
            _nw129[(new BigNumber(2)).toNumber()] = _920_v59;
            _nw129[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_920_v59, _920_v59);
            _nw129[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(949)), ((_924_v2) => function (_925_i7) {
              return _924_v2;
            })(_841_v2));
            _nw129[(new BigNumber(5)).toNumber()] = _920_v59;
            _nw129[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_920_v59, _dafny.Seq.Create(_module.__default.abs(new BigNumber(31)), ((_926_v2) => function (_927_i8) {
              return _926_v2;
            })(_841_v2)));
            _nw129[(new BigNumber(7)).toNumber()] = _dafny.Seq.update(_dafny.Seq.update(_920_v59, _module.__default.safeIndex(new BigNumber((_921_v60).length), new BigNumber((_920_v59).length)), (_839_v1)[_module.__default.safeIndex(new BigNumber(121), new BigNumber((_839_v1).length))]), _module.__default.safeIndex(new BigNumber((_844_v5).length), new BigNumber((_dafny.Seq.update(_920_v59, _module.__default.safeIndex(new BigNumber((_921_v60).length), new BigNumber((_920_v59).length)), (_839_v1)[_module.__default.safeIndex(new BigNumber(121), new BigNumber((_839_v1).length))])).length)), new _dafny.CodePoint('d'.codePointAt(0)));
            _nw129[(new BigNumber(8)).toNumber()] = _dafny.Seq.UnicodeFromString("xry");
            _nw129[(new BigNumber(9)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(49)), ((_928_v2) => function (_929_i9) {
              return _928_v2;
            })(_841_v2));
            _922_v61 = _nw129;
            let _index134 = _module.__default.safeIndex(new BigNumber(629), new BigNumber((_922_v61).length));
            (_922_v61)[_index134] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(284)), function (_930_i10) {
              return new _dafny.CodePoint('w'.codePointAt(0));
            });
            let _931_v62;
            _931_v62 = _module.D13.create_DC26(_920_v59, _842_v3, _842_v3);
            let _pat_let_tv19 = _842_v3;
            let _index135 = _module.__default.safeIndex(new BigNumber(121), new BigNumber((_839_v1).length));
            let _index136 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_917_v58).length));
            let _index137 = _module.__default.safeIndex(new BigNumber(629), new BigNumber((_922_v61).length));
            let _rhs110 = _dafny.Map.Empty.slice().updateUnsafe((function (_pat_let14_0) {
              return function (_932_dt__update__tmp_h1) {
                return function (_pat_let15_0) {
                  return function (_933_dt__update_hcf34_h0) {
                    return _module.D13.create_DC26((_932_dt__update__tmp_h1).dtor_cf33, _933_dt__update_hcf34_h0, (_932_dt__update__tmp_h1).dtor_cf35);
                  }(_pat_let15_0);
                }(_pat_let_tv19);
              }(_pat_let14_0);
            }(_931_v62)).dtor_cf34,_915_v56);
            let _rhs111 = (((_842_v3) && (_842_v3)) ? ((_839_v1)[_module.__default.safeIndex(new BigNumber(121), new BigNumber((_839_v1).length))]) : ((_839_v1)[_module.__default.safeIndex(new BigNumber(121), new BigNumber((_839_v1).length))]));
            let _rhs112 = (_module.__default.fm2(_843_v4, globalState)) && ((_931_v62).dtor_cf35);
            let _rhs113 = _dafny.Seq.update(_920_v59, _module.__default.safeIndex((_this).f28, new BigNumber((_920_v59).length)), (_839_v1)[_module.__default.safeIndex(new BigNumber(121), new BigNumber((_839_v1).length))]);
            let _lhs109 = _839_v1;
            let _lhs110 = _module.__default.safeIndex(new BigNumber(121), new BigNumber((_839_v1).length));
            let _lhs111 = _917_v58;
            let _lhs112 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_917_v58).length));
            let _lhs113 = _922_v61;
            let _lhs114 = _module.__default.safeIndex(new BigNumber(629), new BigNumber((_922_v61).length));
            _916_v57 = _rhs110;
            _lhs109[_lhs110] = _rhs111;
            _lhs111[_lhs112] = _rhs112;
            _lhs113[_lhs114] = _rhs113;
          }
        }
      }
      let _934_v63;
      _934_v63 = true;
      let _935_v64;
      _935_v64 = _dafny.Seq.of(_934_v63, _934_v63);
      let _936_v65;
      _936_v65 = _dafny.Map.Empty.slice().updateUnsafe(_837_v0,_935_v64);
      let _937_v66;
      _937_v66 = _module.D16.create_DC35(_936_v65);
      let _938_v67;
      let _nw130 = Array((new BigNumber(28)).toNumber()).fill([]);
      _938_v67 = _nw130;
      let _939_v68;
      _939_v68 = _dafny.MultiSet.fromElements(_934_v63, false);
      let _940_v69;
      _940_v69 = _module.D16.create_DC36(true, _939_v68, _934_v63);
      let _pat_let_tv20 = _934_v63;
      let _pat_let_tv21 = _934_v63;
      let _941_v70;
      let _nw131 = Array((new BigNumber(15)).toNumber());
      _nw131[(_dafny.ZERO).toNumber()] = _940_v69;
      _nw131[(_dafny.ONE).toNumber()] = _940_v69;
      _nw131[(new BigNumber(2)).toNumber()] = _940_v69;
      _nw131[(new BigNumber(3)).toNumber()] = _940_v69;
      _nw131[(new BigNumber(4)).toNumber()] = _940_v69;
      _nw131[(new BigNumber(5)).toNumber()] = _module.D16.create_DC36(_934_v63, _939_v68, _934_v63);
      _nw131[(new BigNumber(6)).toNumber()] = _940_v69;
      _nw131[(new BigNumber(7)).toNumber()] = _940_v69;
      _nw131[(new BigNumber(8)).toNumber()] = _module.__default.fm40(globalState);
      _nw131[(new BigNumber(9)).toNumber()] = _940_v69;
      _nw131[(new BigNumber(10)).toNumber()] = function (_pat_let16_0) {
        return function (_942_dt__update__tmp_h2) {
          return function (_pat_let17_0) {
            return function (_943_dt__update_hcf49_h0) {
              return _module.D16.create_DC36((_942_dt__update__tmp_h2).dtor_cf47, (_942_dt__update__tmp_h2).dtor_cf48, _943_dt__update_hcf49_h0);
            }(_pat_let17_0);
          }(_pat_let_tv20);
        }(_pat_let16_0);
      }(_940_v69);
      _nw131[(new BigNumber(11)).toNumber()] = function (_pat_let18_0) {
        return function (_944_dt__update__tmp_h3) {
          return function (_pat_let19_0) {
            return function (_945_dt__update_hcf47_h0) {
              return _module.D16.create_DC36(_945_dt__update_hcf47_h0, (_944_dt__update__tmp_h3).dtor_cf48, (_944_dt__update__tmp_h3).dtor_cf49);
            }(_pat_let19_0);
          }(_pat_let_tv21);
        }(_pat_let18_0);
      }(_940_v69);
      _nw131[(new BigNumber(12)).toNumber()] = _940_v69;
      _nw131[(new BigNumber(13)).toNumber()] = _module.__default.fm40(globalState);
      _nw131[(new BigNumber(14)).toNumber()] = _940_v69;
      _941_v70 = _nw131;
      let _index138 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_938_v67).length));
      (_938_v67)[_index138] = _941_v70;
      let _946_v71;
      _946_v71 = _module.D17.create_DC38(_941_v70);
      let _index139 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_938_v67).length));
      let _rhs114 = _937_v66;
      let _rhs115 = (_this).f28;
      let _rhs116 = (_946_v71).dtor_cf51;
      let _lhs115 = globalState;
      let _lhs116 = _938_v67;
      let _lhs117 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_938_v67).length));
      _937_v66 = _rhs114;
      _lhs115.f7 = _rhs115;
      _lhs116[_lhs117] = _rhs116;
      let _947_v72;
      _947_v72 = new _dafny.CodePoint('p'.codePointAt(0));
      _947_v72 = _947_v72;
      (globalState).f16 = (_this).f28;
      (globalState).f7 = new BigNumber((_dafny.Set.fromElements(_934_v63, ((_934_v63) ? (_934_v63) : (_934_v63)), _934_v63, _934_v63)).length);
      let _948_v73;
      _948_v73 = _module.D15.create_DC33(_dafny.Seq.update(_dafny.Seq.of((_this).f28), _module.__default.safeIndex((_this).f28, new BigNumber((_dafny.Seq.of((_this).f28)).length)), (_this).f28));
      let _949_v74;
      _949_v74 = _module.D15.create_DC34(_948_v73);
      let _950_v75;
      _950_v75 = _dafny.Map.Empty.slice().updateUnsafe(_949_v74,_module.__default.fm0(new BigNumber(-164), (_this).f28, (_this).f28, true, globalState));
      r1 = !(_934_v63) || (!(new BigNumber((_950_v75).length)).isEqualTo((_this).f28));
      r0 = _934_v63;
      r1 = !(_934_v63);
      let _951_v76;
      _951_v76 = _module.D4.create_DC9((_this).f28);
      r2 = (_this).fm6((_951_v76).dtor_cf11, globalState);
      let _952_v77;
      _952_v77 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,_934_v63);
      let _953_v78;
      _953_v78 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f28,_934_v63), _952_v77);
      let _954_v79;
      let _nw132 = new _module.C1();
      _nw132.__ctor((_this).f28, _953_v78);
      _954_v79 = _nw132;
      let _955_v80;
      _955_v80 = _dafny.Seq.of(_954_v79);
      r3 = !(!((new BigNumber((_dafny.Seq.Concat(_955_v80, _955_v80)).length)).isLessThan(_module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber(-779)), (_this).f28))));
      return [r0, r1, r2, r3];
    }
    get f28() {
      let _this = this;
      return _this._f28;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(523)).isEqualTo(new BigNumber((_dafny.Seq.of(_dafny.Seq.of(false))).length)),_dafny.Seq.Create(_module.__default.abs(new BigNumber(992)), function (_956_i0) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      }));
    };
    fm6(p0, globalState) {
      let _this = this;
      return false;
    };
    m1(p0, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = false;
      let _957_v0;
      _957_v0 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      let _958_v1;
      _958_v1 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_957_v0).length),p0);
      let _959_v2;
      _959_v2 = _dafny.MultiSet.fromElements(_957_v0, _957_v0);
      let _960_v4;
      _960_v4 = false;
      let _961_v5;
      _961_v5 = _dafny.MultiSet.fromElements(true, _960_v4);
      let _962_v6;
      _962_v6 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((new BigNumber((function () {
        let _coll43 = new _dafny.Set();
        for (const _compr_43 of (_959_v2).Elements) {
          let _963_v3 = _compr_43;
          if ((_959_v2).contains(_963_v3)) {
            _coll43.add(_963_v3);
          }
        }
        return _coll43;
      }()).length)).multipliedBy(p0)),(_961_v5).Union(_961_v5));
      let _964_v7;
      _964_v7 = _dafny.Seq.of(_960_v4);
      let _965_v8;
      _965_v8 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_964_v7));
      _962_v6 = (_962_v6).update(p0, (_965_v8)[_module.__default.safeIndex(p0, new BigNumber((_965_v8).length))]);
      if (_960_v4) {
        let _966_v9;
        let _init21 = ((_967_p0) => function (_968_i0) {
          return _module.__default.safeModuloInt(_968_i0, _967_p0);
        })(p0);
        let _nw133 = Array((new BigNumber(25)).toNumber());
        for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw133.length); _i0_21++) {
          _nw133[_i0_21] = _init21(new BigNumber(_i0_21));
        }
        _966_v9 = _nw133;
        let _969_v10;
        _969_v10 = _dafny.Seq.of(_966_v9, _966_v9);
        let _970_v11;
        _970_v11 = _dafny.Map.Empty.slice().updateUnsafe(_960_v4,_960_v4);
        let _971_v12;
        _971_v12 = _module.D4.create_DC10(p0, new BigNumber((_970_v11).length), _960_v4, _960_v4, p0);
        let _972_v13;
        _972_v13 = _dafny.Seq.UnicodeFromString("eofajvq");
        let _973_v14;
        _973_v14 = _module.D6.create_DC12(_969_v10);
        let _974_v15;
        _974_v15 = _dafny.Seq.of((_973_v14).dtor_cf18, _969_v10, _dafny.Seq.Concat(_969_v10, _969_v10));
        let _rhs117 = (_971_v12).dtor_cf14;
        let _rhs118 = !(((new BigNumber((_972_v13).length)).multipliedBy(new BigNumber((_970_v11).length))).isLessThanOrEqualTo(p0));
        let _rhs119 = (_974_v15)[_module.__default.safeIndex(p0, new BigNumber((_974_v15).length))];
        let _lhs118 = globalState;
        r1 = _rhs117;
        _lhs118.f0 = _rhs118;
        _969_v10 = _rhs119;
        (globalState).f16 = p0;
        let _975_v16;
        let _nw134 = Array((new BigNumber(14)).toNumber()).fill(false);
        _975_v16 = _nw134;
        let _976_v17;
        _976_v17 = _dafny.Set.fromElements(_975_v16, _975_v16, _975_v16);
        let _rhs120 = p0;
        let _rhs121 = ((_960_v4) ? (p0) : ((_dafny.ZERO).minus((new BigNumber(420)).multipliedBy(p0))));
        let _rhs122 = _964_v7;
        let _rhs123 = _976_v17;
        let _rhs124 = _966_v9;
        let _lhs119 = globalState;
        let _lhs120 = globalState;
        _lhs119.f16 = _rhs120;
        _lhs120.f7 = _rhs121;
        _964_v7 = _rhs122;
        _976_v17 = _rhs123;
        r0 = _rhs124;
        let _977_v18;
        let _nw135 = new _module.C0();
        _nw135.__ctor(_960_v4);
        _977_v18 = _nw135;
        let _pat_let_tv22 = _969_v10;
        let _pat_let_tv23 = _969_v10;
        let _pat_let_tv24 = p0;
        let _pat_let_tv25 = _969_v10;
        let _pat_let_tv26 = _969_v10;
        let _pat_let_tv27 = _966_v9;
        _973_v14 = function (_pat_let20_0) {
          return function (_978_dt__update__tmp_h0) {
            return function (_pat_let21_0) {
              return function (_979_dt__update_hcf18_h0) {
                return _module.D6.create_DC12(_979_dt__update_hcf18_h0);
              }(_pat_let21_0);
            }(_dafny.Seq.update(_dafny.Seq.Concat(_pat_let_tv22, _pat_let_tv23), _module.__default.safeIndex(_pat_let_tv24, new BigNumber((_dafny.Seq.Concat(_pat_let_tv25, _pat_let_tv26)).length)), _pat_let_tv27));
          }(_pat_let20_0);
        }(_973_v14);
      } else {
        (globalState).f7 = new BigNumber(639);
        let _980_v19;
        _980_v19 = _dafny.Map.Empty.slice().updateUnsafe(_960_v4,(_module.__default.fm0(p0, p0, p0, _960_v4, globalState)).isLessThanOrEqualTo(p0));
        _980_v19 = (_980_v19).update(_960_v4, _960_v4);
        (globalState).f7 = (_dafny.ZERO).minus(p0);
        (globalState).f21 = (p0).isLessThan(p0);
        let _981_v20;
        _981_v20 = _dafny.Set.fromElements(_960_v4, !(_960_v4) || (false), _960_v4, ((_960_v4) ? (_960_v4) : (_960_v4)), _960_v4);
        _981_v20 = _dafny.Set.fromElements(_960_v4);
      }
      (globalState).f21 = _960_v4;
      let _982_v21;
      _982_v21 = _module.D0.create_DC0(_960_v4);
      let _pat_let_tv28 = _960_v4;
      let _pat_let_tv29 = _960_v4;
      if (function (_source16) {
        if (_source16.is_DC0) {
          let _983___mcc_h0 = (_source16).cf0;
          let _984_cf0 = _983___mcc_h0;
          if (_pat_let_tv28) {
            return _984_cf0;
          } else {
            return _pat_let_tv29;
          }
        } else {
          let _985___mcc_h1 = (_source16).cf1;
          let _986_cf1 = _985___mcc_h1;
          return true;
        }
      }(_982_v21)) {
        let _987_v22;
        _987_v22 = new _dafny.CodePoint('e'.codePointAt(0));
        _987_v22 = _987_v22;
        let _988_v23;
        let _nw136 = new _module.C0();
        _nw136.__ctor(_960_v4);
        _988_v23 = _nw136;
        (globalState).f2 = (true) === (_988_v23.f27);
        let _989_v24;
        let _nw137 = new _module.C0();
        _nw137.__ctor((_988_v23.f27) === (_988_v23.f27));
        _989_v24 = _nw137;
        let _990_v25;
        let _init22 = ((_991_v23, _992_p0) => function (_993_i1) {
          return _module.__default.safeModuloInt(_993_i1, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_991_v23.f27,_module.D4.create_DC9(_992_p0))).length));
        })(_988_v23, p0);
        let _nw138 = Array((new BigNumber(9)).toNumber());
        for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw138.length); _i0_22++) {
          _nw138[_i0_22] = _init22(new BigNumber(_i0_22));
        }
        _990_v25 = _nw138;
        let _994_v26;
        _994_v26 = _dafny.Seq.of(p0, p0);
        let _index140 = _module.__default.safeIndex(new BigNumber(128), new BigNumber((_990_v25).length));
        (_990_v25)[_index140] = (new BigNumber((_994_v26).length)).minus(p0);
        let _995_v27;
        _995_v27 = _dafny.Set.fromElements(_960_v4);
        let _index141 = _module.__default.safeIndex(new BigNumber(128), new BigNumber((_990_v25).length));
        let _rhs125 = (_995_v27).IsSubsetOf((_995_v27).Difference(_995_v27));
        let _rhs126 = _module.__default.fm0((p0).minus(p0), new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), p0, _989_v24.f27, globalState);
        let _rhs127 = _988_v23.f27;
        let _lhs121 = globalState;
        let _lhs122 = _990_v25;
        let _lhs123 = _module.__default.safeIndex(new BigNumber(128), new BigNumber((_990_v25).length));
        let _lhs124 = globalState;
        _lhs121.f21 = _rhs125;
        _lhs122[_lhs123] = _rhs126;
        _lhs124.f2 = _rhs127;
      } else {
        let _996_v28;
        _996_v28 = _dafny.Seq.UnicodeFromString("tsxxruaee");
        (globalState).f19 = _dafny.Seq.Concat(_996_v28, _dafny.Seq.Concat(_996_v28, _996_v28));
        let _997_v29;
        _997_v29 = _dafny.MultiSet.fromElements(p0, p0, p0, p0, p0);
        let _998_v30;
        _998_v30 = _dafny.Seq.of(p0, p0);
        let _999_v31;
        _999_v31 = _module.D7.create_DC15(_998_v30);
        let _1000_v32;
        _1000_v32 = _module.D7.create_DC15((_999_v31).dtor_cf22);
        let _1001_v33;
        let _nw139 = Array((new BigNumber(22)).toNumber());
        _nw139[(_dafny.ZERO).toNumber()] = _997_v29;
        _nw139[(_dafny.ONE).toNumber()] = _997_v29;
        _nw139[(new BigNumber(2)).toNumber()] = _997_v29;
        _nw139[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.FromArray(_dafny.Seq.of(p0));
        _nw139[(new BigNumber(4)).toNumber()] = _997_v29;
        _nw139[(new BigNumber(5)).toNumber()] = _dafny.MultiSet.fromElements(p0, new BigNumber(790));
        _nw139[(new BigNumber(6)).toNumber()] = _997_v29;
        _nw139[(new BigNumber(7)).toNumber()] = _997_v29;
        _nw139[(new BigNumber(8)).toNumber()] = _dafny.MultiSet.fromElements(p0, (_dafny.ZERO).minus((_dafny.ZERO).minus(p0)), p0, new BigNumber((_996_v28).length));
        _nw139[(new BigNumber(9)).toNumber()] = (_dafny.MultiSet.FromArray((_999_v31).dtor_cf22)).Union(_997_v29);
        _nw139[(new BigNumber(10)).toNumber()] = ((_960_v4) ? (_997_v29) : (_997_v29));
        _nw139[(new BigNumber(11)).toNumber()] = _997_v29;
        _nw139[(new BigNumber(12)).toNumber()] = (_997_v29).Union(_997_v29);
        _nw139[(new BigNumber(13)).toNumber()] = (_997_v29).Difference(_997_v29);
        _nw139[(new BigNumber(14)).toNumber()] = _997_v29;
        _nw139[(new BigNumber(15)).toNumber()] = (_997_v29).Intersect(_997_v29);
        _nw139[(new BigNumber(16)).toNumber()] = _dafny.MultiSet.fromElements(p0);
        _nw139[(new BigNumber(17)).toNumber()] = _997_v29;
        _nw139[(new BigNumber(18)).toNumber()] = _997_v29;
        _nw139[(new BigNumber(19)).toNumber()] = ((_997_v29).update(p0, _module.__default.abs(new BigNumber((_996_v28).length)))).Difference(_997_v29);
        _nw139[(new BigNumber(20)).toNumber()] = _997_v29;
        _nw139[(new BigNumber(21)).toNumber()] = (_module.__default.fm12(_996_v28, new BigNumber(453), _960_v4, globalState)).Union(_dafny.MultiSet.FromArray(_998_v30));
        _1001_v33 = _nw139;
        let _index142 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_1001_v33).length));
        (_1001_v33)[_index142] = (_997_v29).Intersect(_997_v29);
        let _index143 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_1001_v33).length));
        (_1001_v33)[_index143] = _dafny.MultiSet.FromArray(_998_v30);
        (globalState).f2 = _960_v4;
        let _1002_v34;
        let _init23 = ((_1003_p0) => function (_1004_i2) {
          return _module.__default.safeModuloInt(_1004_i2, _1003_p0);
        })(p0);
        let _nw140 = Array((new BigNumber(19)).toNumber());
        for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw140.length); _i0_23++) {
          _nw140[_i0_23] = _init23(new BigNumber(_i0_23));
        }
        _1002_v34 = _nw140;
        let _index144 = _module.__default.safeIndex(new BigNumber(410), new BigNumber((_1002_v34).length));
        (_1002_v34)[_index144] = p0;
        let _index145 = _module.__default.safeIndex(new BigNumber(410), new BigNumber((_1002_v34).length));
        (_1002_v34)[_index145] = new BigNumber(62);
        (globalState).f13 = _960_v4;
      }
      let _1005_v35;
      _1005_v35 = new _dafny.CodePoint('b'.codePointAt(0));
      let _1006_v36;
      _1006_v36 = _dafny.MultiSet.fromElements(_1005_v35, new _dafny.CodePoint('c'.codePointAt(0)), _1005_v35, new _dafny.CodePoint('v'.codePointAt(0)));
      let _1007_v37;
      _1007_v37 = _dafny.Seq.UnicodeFromString("xysyc");
      let _1008_v38;
      _1008_v38 = _module.D4.create_DC10(p0, new BigNumber(30), !(_960_v4), _960_v4, (_dafny.ZERO).minus(p0));
      let _1009_v39;
      _1009_v39 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_1007_v37), _1006_v36);
      let _1010_v40;
      let _nw141 = Array((new BigNumber(12)).toNumber());
      _nw141[(_dafny.ZERO).toNumber()] = _1006_v36;
      _nw141[(_dafny.ONE).toNumber()] = _1006_v36;
      _nw141[(new BigNumber(2)).toNumber()] = _1006_v36;
      _nw141[(new BigNumber(3)).toNumber()] = _1006_v36;
      _nw141[(new BigNumber(4)).toNumber()] = _1006_v36;
      _nw141[(new BigNumber(5)).toNumber()] = _module.__default.fm13(new BigNumber((_1007_v37).length), !(_960_v4), _1008_v38, _960_v4, globalState);
      _nw141[(new BigNumber(6)).toNumber()] = _1006_v36;
      _nw141[(new BigNumber(7)).toNumber()] = _1006_v36;
      _nw141[(new BigNumber(8)).toNumber()] = (_1006_v36).Union((_1009_v39)[_module.__default.safeIndex(p0, new BigNumber((_1009_v39).length))]);
      _nw141[(new BigNumber(9)).toNumber()] = (_1006_v36).Intersect(_1006_v36);
      _nw141[(new BigNumber(10)).toNumber()] = ((_960_v4) ? (_1006_v36) : (_1006_v36));
      _nw141[(new BigNumber(11)).toNumber()] = _1006_v36;
      _1010_v40 = _nw141;
      for (const _guard_loop_7 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1010_v40).length))) {
        let _1011_i3 = _guard_loop_7;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1011_i3)) && ((_1011_i3).isLessThan(new BigNumber((_1010_v40).length))))) {
          (_1010_v40)[(_1011_i3)] = _dafny.MultiSet.fromElements(_1005_v35);
        }
      }
      let _1012_v41;
      let _nw142 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Set.Empty);
      _1012_v41 = _nw142;
      let _1013_v42;
      _1013_v42 = _module.D7.create_DC16(_1012_v41);
      let _1014_v43;
      _1014_v43 = _module.D7.create_DC17(_1013_v42);
      let _1015_v44;
      _1015_v44 = _module.D7.create_DC17(_1014_v43);
      let _source17 = _1015_v44;
      if (_source17.is_DC16) {
        let _1016___mcc_h2 = (_source17).cf23;
        let _1017_cf23 = _1016___mcc_h2;
        let _1018_v45;
        _1018_v45 = _dafny.Map.Empty.slice().updateUnsafe(_1007_v37,p0);
        _1018_v45 = (_1018_v45).update(_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("jyay"), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.UnicodeFromString("jyay")).length)), _1005_v35), _1007_v37), p0);
        let _1019_v46;
        _1019_v46 = _module.D6.create_DC13(p0, p0);
        (globalState).f21 = (_964_v7)[_module.__default.safeIndex((_1019_v46).dtor_cf20, new BigNumber((_964_v7).length))];
        let _1020_v47;
        _1020_v47 = _dafny.Map.Empty.slice().updateUnsafe(_960_v4,_dafny.Seq.UnicodeFromString("eaxlkubrh"));
        let _1021_v48;
        _1021_v48 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.update((((_1020_v47).contains(_960_v4)) ? ((_1020_v47).get(_960_v4)) : (_1007_v37)), _module.__default.safeIndex(new BigNumber(825), new BigNumber(((((_1020_v47).contains(_960_v4)) ? ((_1020_v47).get(_960_v4)) : (_1007_v37))).length)), _1005_v35), _1007_v37),(_964_v7)[_module.__default.safeIndex(_module.__default.fm0(new BigNumber(-147), p0, p0, _960_v4, globalState), new BigNumber((_964_v7).length))]);
        _1021_v48 = (_1021_v48).update(_1007_v37, (p0).isLessThanOrEqualTo(p0));
        if (_960_v4) {
          (globalState).f2 = _960_v4;
          let _1022_v49;
          let _nw143 = Array((new BigNumber(10)).toNumber()).fill([]);
          _1022_v49 = _nw143;
          let _1023_v50;
          let _init24 = ((_1024_p0) => function (_1025_i4) {
            return (_1025_i4).minus(_1024_p0);
          })(p0);
          let _nw144 = Array((new BigNumber(12)).toNumber());
          for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw144.length); _i0_24++) {
            _nw144[_i0_24] = _init24(new BigNumber(_i0_24));
          }
          _1023_v50 = _nw144;
          let _index146 = _module.__default.safeIndex(new BigNumber(964), new BigNumber((_1022_v49).length));
          (_1022_v49)[_index146] = _1023_v50;
          let _index147 = _module.__default.safeIndex(new BigNumber(964), new BigNumber((_1022_v49).length));
          (_1022_v49)[_index147] = _1023_v50;
          (globalState).f21 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("nxdew"), _1007_v37), _1007_v37);
          let _1026_v51;
          _1026_v51 = _dafny.Set.fromElements(_960_v4, _960_v4, _960_v4);
          let _1027_v52;
          _1027_v52 = _dafny.Set.fromElements(!(_1026_v51).equals(_module.__default.fm14((_dafny.ZERO).minus(p0), _960_v4, globalState)), _960_v4, _960_v4, (new BigNumber(580)).isLessThanOrEqualTo(p0));
          _1027_v52 = (_dafny.Set.fromElements(_960_v4, _960_v4, _960_v4)).Union((_dafny.Set.fromElements(!(_960_v4), _960_v4, _960_v4)).Intersect(_dafny.Set.fromElements(_960_v4)));
          let _1028_v53;
          _1028_v53 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(p0, new BigNumber(-127))).cardinality()));
          let _1029_v54;
          _1029_v54 = _dafny.Seq.of(_1023_v50);
          let _1030_v55;
          _1030_v55 = _dafny.Set.fromElements(p0, new BigNumber(318));
          let _1031_v56;
          _1031_v56 = _dafny.Map.Empty.slice().updateUnsafe(_1028_v53,(_dafny.Set.fromElements(p0, new BigNumber((_1029_v54).length), _module.__default.fm0((_dafny.ZERO).minus(p0), p0, p0, _960_v4, globalState), p0, p0)).IsDisjointFrom(_1030_v55));
          let _1032_v57;
          _1032_v57 = _964_v7;
          let _1033_v58;
          let _nw145 = new _module.C0();
          _nw145.__ctor(_960_v4);
          _1033_v58 = _nw145;
          let _1034_v59;
          _1034_v59 = _dafny.Map.Empty.slice().updateUnsafe(_1033_v58,_1028_v53);
          _1031_v56 = (_1031_v56).update(_dafny.MultiSet.fromElements(new BigNumber(((_1032_v57)).length), p0, new BigNumber((_1034_v59).length)), (_this).fm6(new BigNumber((_1007_v37).length), globalState));
        } else {
          let _1035_v60;
          let _init25 = ((_1036_p0) => function (_1037_i5) {
            return _module.__default.safeModuloInt(_1037_i5, (_module.D4.create_DC9(_1036_p0)).dtor_cf11);
          })(p0);
          let _nw146 = Array((new BigNumber(14)).toNumber());
          for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw146.length); _i0_25++) {
            _nw146[_i0_25] = _init25(new BigNumber(_i0_25));
          }
          _1035_v60 = _nw146;
          let _index148 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_1035_v60).length));
          (_1035_v60)[_index148] = p0;
          let _index149 = _module.__default.safeIndex(new BigNumber(556), new BigNumber((_1035_v60).length));
          (_1035_v60)[_index149] = p0;
          let _1038_v61;
          _1038_v61 = _dafny.Seq.of((_1008_v38).dtor_cf16);
          _1005_v35 = _module.__default.fm15(_dafny.Seq.Concat(_1038_v61, _1038_v61), p0, globalState);
          let _1039_v62;
          _1039_v62 = _dafny.Map.Empty.slice().updateUnsafe(_960_v4,_1035_v60);
          let _1040_v63;
          _1040_v63 = _dafny.Map.Empty.slice().updateUnsafe(_1005_v35,p0);
          let _1041_v64;
          _1041_v64 = _dafny.MultiSet.fromElements(new BigNumber(372), (((_1040_v63).contains(_1005_v35)) ? ((_1040_v63).get(_1005_v35)) : ((_1035_v60)[_module.__default.safeIndex(new BigNumber(556), new BigNumber((_1035_v60).length))])));
          let _1042_v65;
          _1042_v65 = _dafny.Map.Empty.slice().updateUnsafe(_960_v4,_1041_v64);
          let _rhs128 = (_module.__default.safeModuloInt(new BigNumber(((_961_v5)).cardinality()), new BigNumber((_1007_v37).length))).isLessThanOrEqualTo(_module.__default.fm0(p0, new BigNumber((_1042_v65).length), (_1035_v60)[_module.__default.safeIndex(new BigNumber(556), new BigNumber((_1035_v60).length))], _960_v4, globalState));
          let _rhs129 = _961_v5;
          let _rhs130 = (_1039_v62).Merge(_1039_v62);
          let _lhs125 = globalState;
          _lhs125.f21 = _rhs128;
          _961_v5 = _rhs129;
          _1039_v62 = _rhs130;
          (globalState).f7 = p0;
          (globalState).f0 = _960_v4;
        }
      } else if (_source17.is_DC15) {
        let _1043___mcc_h3 = (_source17).cf22;
        let _1044_cf22 = _1043___mcc_h3;
        let _1045_v66;
        _1045_v66 = _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(945)), ((_1046_v35) => function (_1047_i6) {
          return _1046_v35;
        })(_1005_v35)), _dafny.Seq.UnicodeFromString("tbsqpray"), _1007_v37);
        (globalState).f7 = new BigNumber((_1045_v66).cardinality());
        (globalState).f0 = (_964_v7)[_module.__default.safeIndex(p0, new BigNumber((_964_v7).length))];
        (globalState).f7 = p0;
        let _1048_v67;
        let _nw147 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
        _1048_v67 = _nw147;
        let _index150 = _module.__default.safeIndex(new BigNumber(517), new BigNumber((_1048_v67).length));
        (_1048_v67)[_index150] = p0;
        let _1049_v68;
        let _nw148 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Map.Empty);
        _1049_v68 = _nw148;
        let _1050_v69;
        _1050_v69 = _dafny.Map.Empty.slice().updateUnsafe(!(_960_v4),_1048_v67);
        let _index151 = _module.__default.safeIndex(new BigNumber(636), new BigNumber((_1049_v68).length));
        (_1049_v68)[_index151] = _1050_v69;
        let _1051_v70;
        _1051_v70 = _dafny.Seq.of(_1050_v69);
        let _index152 = _module.__default.safeIndex(new BigNumber(517), new BigNumber((_1048_v67).length));
        let _index153 = _module.__default.safeIndex(new BigNumber(636), new BigNumber((_1049_v68).length));
        let _rhs131 = (new BigNumber((_module.__default.fm14(p0, _960_v4, globalState)).length)).plus((_dafny.ZERO).minus(p0));
        let _rhs132 = _1044_cf22;
        let _rhs133 = (_1051_v70)[_module.__default.safeIndex(_module.__default.fm0(p0, p0, p0, _960_v4, globalState), new BigNumber((_1051_v70).length))];
        let _lhs126 = _1048_v67;
        let _lhs127 = _module.__default.safeIndex(new BigNumber(517), new BigNumber((_1048_v67).length));
        let _lhs128 = _1049_v68;
        let _lhs129 = _module.__default.safeIndex(new BigNumber(636), new BigNumber((_1049_v68).length));
        _lhs126[_lhs127] = _rhs131;
        _1044_cf22 = _rhs132;
        _lhs128[_lhs129] = _rhs133;
      } else {
        let _1052___mcc_h4 = (_source17).cf24;
        let _1053_cf24 = _1052___mcc_h4;
        if (_960_v4) {
          (globalState).f21 = _dafny.Seq.IsPrefixOf(_1007_v37, _1007_v37);
          let _1054_v71;
          _1054_v71 = _dafny.Map.Empty.slice().updateUnsafe(_960_v4,_960_v4);
          let _1055_v72;
          _1055_v72 = _dafny.Set.fromElements(_1054_v71);
          let _1056_v73;
          let _nw149 = new _module.C0();
          _nw149.__ctor(_960_v4);
          _1056_v73 = _nw149;
          let _1057_v74;
          let _nw150 = Array((new BigNumber(24)).toNumber());
          _nw150[(_dafny.ZERO).toNumber()] = _1056_v73;
          _nw150[(_dafny.ONE).toNumber()] = _1056_v73;
          _nw150[(new BigNumber(2)).toNumber()] = _1056_v73;
          _nw150[(new BigNumber(3)).toNumber()] = _1056_v73;
          _nw150[(new BigNumber(4)).toNumber()] = _1056_v73;
          _nw150[(new BigNumber(5)).toNumber()] = _1056_v73;
          _nw150[(new BigNumber(6)).toNumber()] = _1056_v73;
          _nw150[(new BigNumber(7)).toNumber()] = _1056_v73;
          _nw150[(new BigNumber(8)).toNumber()] = _1056_v73;
          _nw150[(new BigNumber(9)).toNumber()] = _1056_v73;
          _nw150[(new BigNumber(10)).toNumber()] = _1056_v73;
          _nw150[(new BigNumber(11)).toNumber()] = _1056_v73;
          _nw150[(new BigNumber(12)).toNumber()] = _1056_v73;
          _nw150[(new BigNumber(13)).toNumber()] = _1056_v73;
          _nw150[(new BigNumber(14)).toNumber()] = _1056_v73;
          _nw150[(new BigNumber(15)).toNumber()] = _1056_v73;
          _nw150[(new BigNumber(16)).toNumber()] = _1056_v73;
          _nw150[(new BigNumber(17)).toNumber()] = _1056_v73;
          _nw150[(new BigNumber(18)).toNumber()] = _1056_v73;
          _nw150[(new BigNumber(19)).toNumber()] = _1056_v73;
          _nw150[(new BigNumber(20)).toNumber()] = _1056_v73;
          _nw150[(new BigNumber(21)).toNumber()] = _1056_v73;
          _nw150[(new BigNumber(22)).toNumber()] = _1056_v73;
          _nw150[(new BigNumber(23)).toNumber()] = _1056_v73;
          _1057_v74 = _nw150;
          let _1058_v75;
          _1058_v75 = _dafny.Seq.of(_1057_v74, _1057_v74);
          let _rhs134 = (_1055_v72).Intersect((_1055_v72).Intersect(_1055_v72));
          let _rhs135 = _1007_v37;
          let _rhs136 = _960_v4;
          let _rhs137 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_1057_v74, _1057_v74), _1058_v75), _1058_v75)).length);
          let _lhs130 = globalState;
          let _lhs131 = globalState;
          let _lhs132 = globalState;
          _1055_v72 = _rhs134;
          _lhs130.f8 = _rhs135;
          _lhs131.f2 = _rhs136;
          _lhs132.f7 = _rhs137;
          (globalState).f13 = _1056_v73.f27;
          let _1059_v76;
          let _init26 = ((_1060_v73) => function (_1061_i7) {
            return _1060_v73.f27;
          })(_1056_v73);
          let _nw151 = Array((new BigNumber(19)).toNumber());
          for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw151.length); _i0_26++) {
            _nw151[_i0_26] = _init26(new BigNumber(_i0_26));
          }
          _1059_v76 = _nw151;
          let _index154 = _module.__default.safeIndex(new BigNumber(594), new BigNumber((_1059_v76).length));
          (_1059_v76)[_index154] = _module.__default.fm2(_1054_v71, globalState);
          let _1062_v77;
          _1062_v77 = _dafny.Set.fromElements(p0);
          let _1063_v78;
          _1063_v78 = _dafny.MultiSet.fromElements(_1062_v77, (_1062_v77).Union(_dafny.Set.fromElements(p0)), (_dafny.Set.fromElements(p0, p0)).Union(_1062_v77), _1062_v77);
          let _index155 = _module.__default.safeIndex(new BigNumber(594), new BigNumber((_1059_v76).length));
          let _rhs138 = new BigNumber(899);
          let _rhs139 = _dafny.Seq.IsProperPrefixOf(_1007_v37, _1007_v37);
          let _rhs140 = _1063_v78;
          let _lhs133 = globalState;
          let _lhs134 = _1059_v76;
          let _lhs135 = _module.__default.safeIndex(new BigNumber(594), new BigNumber((_1059_v76).length));
          _lhs133.f16 = _rhs138;
          _lhs134[_lhs135] = _rhs139;
          _1063_v78 = _rhs140;
          let _1064_v79;
          let _nw152 = new _module.C0();
          _nw152.__ctor(_960_v4);
          _1064_v79 = _nw152;
        } else {
          _1008_v38 = _1008_v38;
          (globalState).f13 = _960_v4;
          (globalState).f8 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("d"), _1007_v37);
          _1005_v35 = _1005_v35;
          let _1065_v80;
          _1065_v80 = _dafny.Seq.of(_964_v7, _964_v7);
          (globalState).f7 = _module.__default.fm0(new BigNumber((_1065_v80).length), (p0).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(94)), ((_1066_p0) => function (_1067_i8) {
            return _1066_p0;
          })(p0))).length)), (((_961_v5).contains(false)) ? ((_961_v5).get(false)) : (p0)), !((_960_v4) || (_960_v4)), globalState);
        }
        let _1068_v81;
        let _nw153 = new _module.C0();
        _nw153.__ctor((_960_v4) || (_960_v4));
        _1068_v81 = _nw153;
        let _1069_v82;
        let _init27 = ((_1070_p0) => function (_1071_i9) {
          return (_1071_i9).minus(_1070_p0);
        })(p0);
        let _nw154 = Array((new BigNumber(6)).toNumber());
        for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw154.length); _i0_27++) {
          _nw154[_i0_27] = _init27(new BigNumber(_i0_27));
        }
        _1069_v82 = _nw154;
        let _1072_v83;
        _1072_v83 = _dafny.Set.fromElements(_1069_v82);
        let _1073_v84;
        let _nw155 = Array((new BigNumber(5)).toNumber()).fill(false);
        _1073_v84 = _nw155;
        let _rhs141 = _1072_v83;
        let _rhs142 = _1005_v35;
        let _rhs143 = _1073_v84;
        _1072_v83 = _rhs141;
        _1005_v35 = _rhs142;
        _1073_v84 = _rhs143;
        let _1074_v85;
        let _nw156 = Array((new BigNumber(25)).toNumber()).fill(_dafny.MultiSet.Empty);
        _1074_v85 = _nw156;
        let _index156 = _module.__default.safeIndex(new BigNumber(480), new BigNumber((_1074_v85).length));
        (_1074_v85)[_index156] = (_961_v5).Intersect(_961_v5);
        let _index157 = _module.__default.safeIndex(new BigNumber(480), new BigNumber((_1074_v85).length));
        (_1074_v85)[_index157] = _961_v5;
      }
      let _1075_v86;
      let _nw157 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
      _1075_v86 = _nw157;
      r0 = ((_960_v4) ? (_1075_v86) : (_1075_v86));
      r1 = _960_v4;
      return [r0, r1];
    }
    m7(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _module.D1.Default();
      let r1 = _dafny.ZERO;
      let _1076_v0;
      let _nw158 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
      _1076_v0 = _nw158;
      let _1077_v1;
      let _nw159 = Array((new BigNumber(18)).toNumber()).fill([]);
      _1077_v1 = _nw159;
      let _1078_v2;
      _1078_v2 = _dafny.Map.Empty.slice().updateUnsafe(_1076_v0,_1077_v1);
      _1078_v2 = (_1078_v2).update(_1076_v0, _1077_v1);
      let _hi4 = p3;
      for (let _1079_i0 = new BigNumber(896); _1079_i0.isLessThan(_hi4); _1079_i0 = _1079_i0.plus(_dafny.ONE)) {
        let _index158 = _module.__default.safeIndex(new BigNumber(593), new BigNumber((p0).length));
        (p0)[_index158] = new _dafny.CodePoint('k'.codePointAt(0));
        let _1080_v3;
        _1080_v3 = new _dafny.CodePoint('o'.codePointAt(0));
        let _index159 = _module.__default.safeIndex(new BigNumber(593), new BigNumber((p0).length));
        (p0)[_index159] = _1080_v3;
        let _1081_v4;
        _1081_v4 = false;
        if (_1081_v4) {
          (globalState).f2 = _1081_v4;
          (globalState).f16 = new BigNumber(-201);
          let _1082_v5;
          _1082_v5 = _dafny.MultiSet.fromElements(new BigNumber(463), new BigNumber(279), p2);
          _1082_v5 = _1082_v5;
          let _1083_v6;
          let _nw160 = new _module.C0();
          _nw160.__ctor(((!(_1081_v4)) ? (_1081_v4) : (_1081_v4)));
          _1083_v6 = _nw160;
          let _index160 = _module.__default.safeIndex(new BigNumber(593), new BigNumber((p0).length));
          (p0)[_index160] = new _dafny.CodePoint('n'.codePointAt(0));
        } else {
          let _1084_v7;
          let _init28 = ((_1085_p3, _1086_p1, _1087_v4, _1088_p2) => function (_1089_i1) {
            return !(false) || ((_module.D4.create_DC10(_1085_p3, _1086_p1, !(_1087_v4), _1087_v4, (_dafny.ZERO).minus(_1088_p2))).dtor_cf15);
          })(p3, p1, _1081_v4, p2);
          let _nw161 = Array((new BigNumber(29)).toNumber());
          for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw161.length); _i0_28++) {
            _nw161[_i0_28] = _init28(new BigNumber(_i0_28));
          }
          _1084_v7 = _nw161;
          let _index161 = _module.__default.safeIndex(new BigNumber(390), new BigNumber((_1084_v7).length));
          (_1084_v7)[_index161] = (_1079_i0).isLessThan(_1079_i0);
          let _index162 = _module.__default.safeIndex(new BigNumber(390), new BigNumber((_1084_v7).length));
          let _rhs144 = _1081_v4;
          let _rhs145 = (_module.D0.create_DC0(_1081_v4)).dtor_cf0;
          let _rhs146 = _1081_v4;
          let _lhs136 = globalState;
          let _lhs137 = _1084_v7;
          let _lhs138 = _module.__default.safeIndex(new BigNumber(390), new BigNumber((_1084_v7).length));
          let _lhs139 = globalState;
          _lhs136.f2 = _rhs144;
          _lhs137[_lhs138] = _rhs145;
          _lhs139.f0 = _rhs146;
          let _1090_v8;
          _1090_v8 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1084_v7)[_module.__default.safeIndex(new BigNumber(390), new BigNumber((_1084_v7).length))],p3)).length), new BigNumber(879), new BigNumber((_dafny.Set.fromElements(false, _1081_v4)).length), p2, p3);
          let _1091_v9;
          _1091_v9 = _dafny.Seq.of(_1090_v8, _1090_v8);
          let _1092_v10;
          _1092_v10 = _dafny.Map.Empty.slice().updateUnsafe((_1091_v9)[_module.__default.safeIndex(p1, new BigNumber((_1091_v9).length))],_1081_v4);
          let _1093_v11;
          _1093_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1081_v4,_1081_v4);
          let _1094_v12;
          _1094_v12 = _dafny.Map.Empty.slice().updateUnsafe((((_1093_v11).contains(_1081_v4)) ? ((_1093_v11).get(_1081_v4)) : ((_1084_v7)[_module.__default.safeIndex(new BigNumber(390), new BigNumber((_1084_v7).length))])),_1081_v4);
          let _1095_v13;
          let _nw162 = new _module.C0();
          _nw162.__ctor(((((_1092_v10).contains(_dafny.MultiSet.fromElements(p2, p3))) ? ((_1092_v10).get(_dafny.MultiSet.fromElements(p2, p3))) : (_1081_v4))) || (_module.__default.fm2(_1093_v11, globalState)));
          _1095_v13 = _nw162;
          let _1096_v14;
          let _nw163 = Array((new BigNumber(9)).toNumber()).fill(_module.D4.Default());
          _1096_v14 = _nw163;
          let _1097_v15;
          _1097_v15 = _dafny.MultiSet.fromElements(_1095_v13);
          let _1098_v16;
          _1098_v16 = _module.D4.create_DC7(new BigNumber((_1097_v15).cardinality()));
          let _index163 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_1096_v14).length));
          (_1096_v14)[_index163] = _1098_v16;
          let _index164 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_1096_v14).length));
          (_1096_v14)[_index164] = _module.__default.fm16(p1, _1081_v4, globalState);
          (globalState).f7 = p2;
          let _index165 = _module.__default.safeIndex(new BigNumber(390), new BigNumber((_1084_v7).length));
          (_1084_v7)[_index165] = (_1084_v7)[_module.__default.safeIndex(new BigNumber(390), new BigNumber((_1084_v7).length))];
        }
        (globalState).f0 = _1081_v4;
        let _1099_v17;
        _1099_v17 = _dafny.Set.fromElements(p1, p2);
        if ((_1099_v17).IsDisjointFrom((_1099_v17).Union(_1099_v17))) {
          let _1100_v18;
          _1100_v18 = _dafny.Seq.of(_1079_i0);
          r1 = (p1).multipliedBy((new BigNumber((_1100_v18).length)).plus(p2));
          let _1101_v19;
          _1101_v19 = _dafny.Seq.of(true);
          (globalState).f16 = new BigNumber((_1101_v19).length);
          let _1102_v20;
          _1102_v20 = _dafny.MultiSet.fromElements(_module.__default.fm15(_dafny.Seq.Create(_module.__default.abs(new BigNumber(20)), ((_1103_p2) => function (_1104_i2) {
            return _1103_p2;
          })(p2)), (_1100_v18)[_module.__default.safeIndex(p3, new BigNumber((_1100_v18).length))], globalState));
          _1102_v20 = _1102_v20;
          let _1105_v21;
          let _nw164 = Array((new BigNumber(7)).toNumber()).fill(_module.D6.Default());
          _1105_v21 = _nw164;
          let _1106_v22;
          _1106_v22 = _dafny.Set.fromElements(_1081_v4);
          let _1107_v23;
          _1107_v23 = _module.D6.create_DC13(new BigNumber((_1106_v22).length), new BigNumber((_dafny.Set.fromElements(p1)).length));
          let _1108_v24;
          _1108_v24 = _module.D6.create_DC14(_1107_v23);
          let _index166 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_1105_v21).length));
          (_1105_v21)[_index166] = _1108_v24;
          let _index167 = _module.__default.safeIndex(new BigNumber(717), new BigNumber((_1105_v21).length));
          (_1105_v21)[_index167] = _1108_v24;
          (globalState).f16 = p3;
        } else {
          let _1109_v25;
          _1109_v25 = _dafny.Seq.of(_1081_v4, _1081_v4, _1081_v4);
          (globalState).f21 = (_1109_v25)[_module.__default.safeIndex(p1, new BigNumber((_1109_v25).length))];
          let _1110_v26;
          _1110_v26 = _dafny.MultiSet.fromElements(_1079_i0);
          let _1111_v27;
          _1111_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1110_v26,_1081_v4);
          let _1112_v28;
          _1112_v28 = _1111_v27;
          (globalState).f2 = (new BigNumber(((_1111_v27)).length)).isEqualTo(_module.__default.safeModuloInt(p2, p1));
          let _index168 = _module.__default.safeIndex(new BigNumber(593), new BigNumber((p0).length));
          (p0)[_index168] = (p0)[_module.__default.safeIndex(new BigNumber(593), new BigNumber((p0).length))];
          _1080_v3 = _1080_v3;
          (globalState).f0 = false;
        }
      }
      let _1113_v29;
      _1113_v29 = false;
      let _1114_i3;
      _1114_i3 = _dafny.ZERO;
      L6: {
        while (!(_1113_v29)) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1114_i3)) {
              break L6;
            }
            _1114_i3 = (_1114_i3).plus(_dafny.ONE);
            (globalState).f13 = ((_1113_v29) ? (_1113_v29) : (((_1113_v29) ? (_1113_v29) : (true))));
            let _1115_v30;
            let _nw165 = Array((new BigNumber(14)).toNumber()).fill(_module.D4.Default());
            _1115_v30 = _nw165;
            let _1116_v31;
            _1116_v31 = _module.D4.create_DC8(_1113_v29);
            let _index169 = _module.__default.safeIndex(new BigNumber(56), new BigNumber((_1115_v30).length));
            (_1115_v30)[_index169] = _1116_v31;
            let _1117_v32;
            _1117_v32 = new _dafny.CodePoint('d'.codePointAt(0));
            let _index170 = _module.__default.safeIndex(new BigNumber(56), new BigNumber((_1115_v30).length));
            let _rhs147 = _1116_v31;
            let _rhs148 = _1117_v32;
            let _lhs140 = _1115_v30;
            let _lhs141 = _module.__default.safeIndex(new BigNumber(56), new BigNumber((_1115_v30).length));
            _lhs140[_lhs141] = _rhs147;
            _1117_v32 = _rhs148;
            let _1118_v33;
            _1118_v33 = _dafny.Seq.of(_1113_v29);
            let _1119_v34;
            _1119_v34 = _dafny.Set.fromElements(_1118_v33, _1118_v33);
            if ((_1113_v29) || ((_1119_v34).IsSubsetOf(_1119_v34))) {
              let _1120_v35;
              let _nw166 = new _module.C0();
              _nw166.__ctor(((_dafny.ZERO).minus(p2)).isLessThanOrEqualTo(new BigNumber(-32)));
              _1120_v35 = _nw166;
              let _1121_v36;
              let _nw167 = Array((new BigNumber(18)).toNumber()).fill(false);
              _1121_v36 = _nw167;
              _1121_v36 = _1121_v36;
              _1076_v0 = _1076_v0;
              let _index171 = _module.__default.safeIndex(new BigNumber(753), new BigNumber((_1121_v36).length));
              (_1121_v36)[_index171] = (_1120_v35.f27) || (_1120_v35.f27);
              let _index172 = _module.__default.safeIndex(new BigNumber(753), new BigNumber((_1121_v36).length));
              (_1121_v36)[_index172] = !_dafny.areEqual(_1118_v33, _1118_v33);
              _1117_v32 = _1117_v32;
            } else {
              let _1122_v37;
              _1122_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1113_v29,_1113_v29);
              let _1123_v38;
              _1123_v38 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((new BigNumber((_1122_v37).length)).minus(p2)),!(_1113_v29));
              let _1124_v39;
              _1124_v39 = _dafny.Map.Empty.slice().updateUnsafe(_1113_v29,p1);
              let _1125_v40;
              _1125_v40 = _dafny.Seq.UnicodeFromString("vn");
              let _1126_v41;
              _1126_v41 = _dafny.Set.fromElements(new BigNumber(-813));
              let _1127_v43;
              _1127_v43 = _dafny.Seq.of(function () {
                let _coll44 = new _dafny.Map();
                for (const _compr_44 of _dafny.IntegerRange(new BigNumber(316), new BigNumber(965))) {
                  let _1128_v42 = _compr_44;
                  if (((new BigNumber(316)).isLessThanOrEqualTo(_1128_v42)) && ((_1128_v42).isLessThan(new BigNumber(965)))) {
                    _coll44.push([_module.__default.safeModuloInt(_1128_v42, p1),true]);
                  }
                }
                return _coll44;
              }());
              let _1129_v44;
              let _nw168 = new _module.C1();
              _nw168.__ctor(p2, _1127_v43);
              _1129_v44 = _nw168;
              let _1130_v45;
              _1130_v45 = _dafny.Seq.of(_1129_v44);
              let _1131_v46;
              _1131_v46 = _dafny.MultiSet.fromElements(_1129_v44, (_1130_v45)[_module.__default.safeIndex(new BigNumber(-289), new BigNumber((_1130_v45).length))]);
              let _rhs149 = (((_1123_v38).contains(p3)) ? ((_1123_v38).get(p3)) : (_1113_v29));
              let _rhs150 = (_dafny.ZERO).minus((((_1124_v39).contains((_this).fm6(p3, globalState))) ? ((_1124_v39).get((_this).fm6(p3, globalState))) : (_module.__default.fm0(new BigNumber(187), new BigNumber(-322), p3, _1113_v29, globalState))));
              let _rhs151 = _dafny.Seq.update(_1125_v40, _module.__default.safeIndex(p1, new BigNumber((_1125_v40).length)), _1117_v32);
              let _rhs152 = ((new BigNumber((_1126_v41).length)).minus((_dafny.ZERO).minus(p2))).isLessThan(p2);
              let _rhs153 = _module.__default.safeModuloInt((((_1131_v46).contains(_1129_v44)) ? ((_1131_v46).get(_1129_v44)) : (p3)), p1);
              let _lhs142 = globalState;
              let _lhs143 = globalState;
              let _lhs144 = globalState;
              let _lhs145 = globalState;
              _lhs142.f13 = _rhs149;
              r1 = _rhs150;
              _lhs143.f20 = _rhs151;
              _lhs144.f0 = _rhs152;
              _lhs145.f7 = _rhs153;
              r1 = p2;
              let _1132_v47;
              _1132_v47 = _dafny.Set.fromElements(p0);
              let _1133_v48;
              _1133_v48 = _dafny.MultiSet.fromElements(new BigNumber((_1132_v47).length));
              let _rhs154 = ((_1113_v29) ? (new BigNumber(260)) : (_module.__default.safeDivisionInt((((_1133_v48).contains(p2)) ? ((_1133_v48).get(p2)) : (p2)), p1)));
              let _rhs155 = _module.__default.fm0(((_dafny.ZERO).minus(new BigNumber(-766))).plus(p1), _module.__default.fm0(p1, p3, p1, false, globalState), new BigNumber(((_1124_v39).Merge(_1124_v39)).length), !(_1113_v29) || (_1113_v29), globalState);
              let _lhs146 = globalState;
              let _lhs147 = globalState;
              _lhs146.f16 = _rhs154;
              _lhs147.f16 = _rhs155;
              let _1134_v49;
              let _nw169 = new _module.C1();
              _nw169.__ctor(p1, _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.of(_1123_v38, _dafny.Map.Empty.slice().updateUnsafe(p2,_1113_v29), _1123_v38, (_1123_v38).update(p3, _1113_v29)), _module.__default.safeIndex(new BigNumber((_1125_v40).length), new BigNumber((_dafny.Seq.of(_1123_v38, _dafny.Map.Empty.slice().updateUnsafe(p2,_1113_v29), _1123_v38, (_1123_v38).update(p3, _1113_v29))).length)), _1123_v38), _module.__default.safeIndex(p3, new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_1123_v38, _dafny.Map.Empty.slice().updateUnsafe(p2,_1113_v29), _1123_v38, (_1123_v38).update(p3, _1113_v29)), _module.__default.safeIndex(new BigNumber((_1125_v40).length), new BigNumber((_dafny.Seq.of(_1123_v38, _dafny.Map.Empty.slice().updateUnsafe(p2,_1113_v29), _1123_v38, (_1123_v38).update(p3, _1113_v29))).length)), _1123_v38)).length)), _1123_v38));
              _1134_v49 = _nw169;
              let _1135_v50;
              let _init29 = function (_1136_i4) {
                return true;
              };
              let _nw170 = Array((new BigNumber(5)).toNumber());
              for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw170.length); _i0_29++) {
                _nw170[_i0_29] = _init29(new BigNumber(_i0_29));
              }
              _1135_v50 = _nw170;
              let _index173 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_1135_v50).length));
              (_1135_v50)[_index173] = _1113_v29;
              let _index174 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_1135_v50).length));
              (_1135_v50)[_index174] = _1113_v29;
            }
            if (_1113_v29) {
              let _1137_v51;
              _1137_v51 = _dafny.Seq.of(p1);
              let _1138_v52;
              _1138_v52 = _module.D15.create_DC33(_dafny.Seq.Concat(_1137_v51, _1137_v51));
              _1138_v52 = _1138_v52;
              (globalState).f2 = _1113_v29;
              (globalState).f7 = p1;
              let _1139_v53;
              let _init30 = ((_1140_v32) => function (_1141_i5) {
                return _1140_v32;
              })(_1117_v32);
              let _nw171 = Array((new BigNumber(2)).toNumber());
              for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw171.length); _i0_30++) {
                _nw171[_i0_30] = _init30(new BigNumber(_i0_30));
              }
              _1139_v53 = _nw171;
              _1139_v53 = p0;
              (globalState).f16 = p2;
            } else {
              let _1142_v54;
              _1142_v54 = _dafny.Map.Empty.slice().updateUnsafe(p3,_1113_v29);
              let _1143_v55;
              _1143_v55 = _dafny.Seq.of(_1142_v54);
              let _1144_v56;
              let _nw172 = new _module.C1();
              _nw172.__ctor(_module.__default.safeDivisionInt(p3, p3), _dafny.Seq.Concat(_1143_v55, _1143_v55));
              _1144_v56 = _nw172;
              let _1145_v57;
              _1145_v57 = _module.D12.create_DC23(p1);
              let _1146_v58;
              _1146_v58 = _dafny.Seq.of((_1145_v57).dtor_cf30, _1144_v56.f29);
              let _index175 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_1076_v0).length));
              (_1076_v0)[_index175] = new BigNumber((_1146_v58).length);
              let _1147_v59;
              _1147_v59 = _dafny.Seq.UnicodeFromString("ehgphob");
              let _index176 = _module.__default.safeIndex(new BigNumber(485), new BigNumber((_1076_v0).length));
              (_1076_v0)[_index176] = _module.__default.safeModuloInt(new BigNumber((_1147_v59).length), new BigNumber(-723));
              let _1148_v60;
              let _nw173 = new _module.C0();
              _nw173.__ctor(_1113_v29);
              _1148_v60 = _nw173;
              let _1149_v62;
              let _init31 = ((_1150_v56, _1151_p1, _1152_p3, _1153_v59) => function (_1154_i6) {
                return (function () {
                  let _coll45 = new _dafny.Map();
                  for (const _compr_45 of _dafny.IntegerRange(new BigNumber(-631), new BigNumber(508))) {
                    let _1155_v61 = _compr_45;
                    if (((new BigNumber(-631)).isLessThanOrEqualTo(_1155_v61)) && ((_1155_v61).isLessThan(new BigNumber(508)))) {
                      _coll45.push([(_1155_v61).minus(_1151_p1),_dafny.MultiSet.fromElements(_1152_p3)]);
                    }
                  }
                  return _coll45;
                }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1150_v56.f29,_dafny.MultiSet.fromElements(new BigNumber(887), new BigNumber((_1153_v59).length))));
              })(_1144_v56, p1, p3, _1147_v59);
              let _nw174 = Array((new BigNumber(14)).toNumber());
              for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw174.length); _i0_31++) {
                _nw174[_i0_31] = _init31(new BigNumber(_i0_31));
              }
              _1149_v62 = _nw174;
              let _1156_v63;
              _1156_v63 = _dafny.MultiSet.fromElements(new BigNumber((_1146_v58).length));
              let _1157_v64;
              _1157_v64 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(472)), ((_1158_p1) => function (_1159_i7) {
                return _1158_p1;
              })(p1)),(_dafny.Map.Empty.slice().updateUnsafe(p3,_1156_v63)).update(new BigNumber(884), _1156_v63));
              let _1160_v65;
              _1160_v65 = _1118_v33;
              let _1161_v66;
              _1161_v66 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.MultiSet.fromElements(new BigNumber(((_1160_v65)).length)));
              let _index177 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_1149_v62).length));
              (_1149_v62)[_index177] = (((_1157_v64).contains(_1146_v58)) ? ((_1157_v64).get(_1146_v58)) : (_1161_v66));
              let _index178 = _module.__default.safeIndex(new BigNumber(95), new BigNumber((_1149_v62).length));
              (_1149_v62)[_index178] = function () {
                let _coll46 = new _dafny.Map();
                for (const _compr_46 of _dafny.IntegerRange(new BigNumber(3), new BigNumber(524))) {
                  let _1162_v67 = _compr_46;
                  if (((new BigNumber(3)).isLessThanOrEqualTo(_1162_v67)) && ((_1162_v67).isLessThan(new BigNumber(524)))) {
                    _coll46.push([_module.__default.safeDivisionInt(_1162_v67, (_dafny.ZERO).minus(p2)),_1156_v63]);
                  }
                }
                return _coll46;
              }();
              (globalState).f21 = _1148_v60.f27;
            }
          }
        }
      }
      let _1163_v68;
      let _nw175 = new _module.C2();
      _nw175.__ctor(p3);
      _1163_v68 = _nw175;
      let _1164_v69;
      _1164_v69 = _dafny.Map.Empty.slice().updateUnsafe((_1163_v68).f28,_1113_v29);
      let _1165_v70;
      _1165_v70 = _dafny.Seq.UnicodeFromString("dlb");
      let _1166_v71;
      _1166_v71 = _dafny.Seq.of(_1164_v69);
      _1164_v69 = ((_module.__default.fm3(p2, _1113_v29, _1165_v70, (_1163_v68).f28, globalState)).Merge((_1166_v71)[_module.__default.safeIndex((_1163_v68).f28, new BigNumber((_1166_v71).length))])).Merge(_1164_v69);
      for (const _guard_loop_8 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1076_v0).length))) {
        let _1167_i8 = _guard_loop_8;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1167_i8)) && ((_1167_i8).isLessThan(new BigNumber((_1076_v0).length))))) {
          (_1076_v0)[(_1167_i8)] = (_1167_i8).minus(p2);
        }
      }
      r0 = _module.D1.create_DC2(_1164_v69);
      r1 = _module.__default.fm0(p3, (_1163_v68).f28, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(372)), ((_1168_v29) => function (_1169_i9) {
        return _dafny.Seq.of(_1168_v29);
      })(_1113_v29))).length), _1113_v29, globalState);
      return [r0, r1];
    }
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this.f25 = false;
      this._f26 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f25, f26) {
      let _this = this;
      (_this).f25 = f25;
      (_this)._f26 = f26;
      return;
    }
    fm9(p0, globalState) {
      let _this = this;
      let _source18 = _module.D3.create_DC6(_dafny.Seq.UnicodeFromString("tfdl"), new _dafny.CodePoint('r'.codePointAt(0)), _this.f25);
      if (_source18.is_DC6) {
        let _1170___mcc_h0 = (_source18).cf6;
        let _1171___mcc_h1 = (_source18).cf7;
        let _1172___mcc_h2 = (_source18).cf8;
        let _1173_cf8 = _1172___mcc_h2;
        let _1174_cf7 = _1171___mcc_h1;
        let _1175_cf6 = _1170___mcc_h0;
        return true;
      } else {
        let _1176___mcc_h3 = (_source18).cf5;
        let _1177_cf5 = _1176___mcc_h3;
        return (_this).f26;
      }
    };
    fm10(p0, p1, p2, globalState) {
      let _this = this;
      return ((function () {
        let _coll47 = new _dafny.Map();
        for (const _compr_47 of _dafny.IntegerRange(new BigNumber(-511), new BigNumber(-748))) {
          let _1178_v0 = _compr_47;
          if (((new BigNumber(-511)).isLessThanOrEqualTo(_1178_v0)) && ((_1178_v0).isLessThan(new BigNumber(-748)))) {
            _coll47.push([(_1178_v0).multipliedBy(new BigNumber((_dafny.Set.fromElements(new BigNumber(-345))).length)),new _dafny.CodePoint('g'.codePointAt(0))]);
          }
        }
        return _coll47;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(387),new _dafny.CodePoint('s'.codePointAt(0))))).Merge((function () {
        let _coll48 = new _dafny.Map();
        for (const _compr_48 of _dafny.IntegerRange(new BigNumber(-284), new BigNumber(557))) {
          let _1179_v1 = _compr_48;
          if (((new BigNumber(-284)).isLessThanOrEqualTo(_1179_v1)) && ((_1179_v1).isLessThan(new BigNumber(557)))) {
            _coll48.push([(_1179_v1).plus(new BigNumber(882)),new _dafny.CodePoint('a'.codePointAt(0))]);
          }
        }
        return _coll48;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_this.f25)).length),new _dafny.CodePoint('p'.codePointAt(0)))));
    };
    m5(p0, globalState) {
      let _this = this;
      let _1180_v0;
      let _nw176 = Array((new BigNumber(5)).toNumber());
      _nw176[(_dafny.ZERO).toNumber()] = p0;
      _nw176[(_dafny.ONE).toNumber()] = _this.f25;
      _nw176[(new BigNumber(2)).toNumber()] = false;
      _nw176[(new BigNumber(3)).toNumber()] = (_this).f26;
      _nw176[(new BigNumber(4)).toNumber()] = _this.f25;
      _1180_v0 = _nw176;
      let _index179 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_1180_v0).length));
      (_1180_v0)[_index179] = !(p0);
      let _1181_v1;
      _1181_v1 = _dafny.Seq.of(_1180_v0, _1180_v0);
      let _1182_v2;
      _1182_v2 = new BigNumber(683);
      let _1183_v3;
      _1183_v3 = _dafny.Map.Empty.slice().updateUnsafe(((_1181_v1)[_module.__default.safeIndex(_1182_v2, new BigNumber((_1181_v1).length))]),!(true));
      let _1184_v4;
      _1184_v4 = _dafny.Map.Empty.slice().updateUnsafe(_1182_v2,_1182_v2);
      let _1185_v5;
      _1185_v5 = _dafny.MultiSet.fromElements(_1182_v2, new BigNumber(-508), new BigNumber((_1184_v4).length));
      let _1186_v6;
      _1186_v6 = _dafny.Seq.of(p0);
      let _1187_v7;
      _1187_v7 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(768));
      let _1188_v8;
      _1188_v8 = _dafny.MultiSet.fromElements(_this.f25, p0);
      let _1189_v9;
      _1189_v9 = _dafny.Seq.of(_1188_v8, _1188_v8);
      let _index180 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_1180_v0).length));
      let _rhs156 = (((_1183_v3).contains(_1180_v0)) ? ((_1183_v3).get(_1180_v0)) : (!(false)));
      let _rhs157 = _module.__default.safeDivisionInt(((((_1185_v5).contains(new BigNumber((_1186_v6).length))) ? ((_1185_v5).get(new BigNumber((_1186_v6).length))) : (_1182_v2))).plus(new BigNumber((_1187_v7).length)), (((_1185_v5).contains(_1182_v2)) ? ((_1185_v5).get(_1182_v2)) : (_1182_v2)));
      let _rhs158 = (_1186_v6)[_module.__default.safeIndex(_1182_v2, new BigNumber((_1186_v6).length))];
      let _rhs159 = _dafny.Seq.IsPrefixOf(_module.__default.fm11(_1182_v2, globalState), _1189_v9);
      let _lhs148 = _this;
      let _lhs149 = globalState;
      let _lhs150 = _this;
      let _lhs151 = _1180_v0;
      let _lhs152 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_1180_v0).length));
      _lhs148.f25 = _rhs156;
      _lhs149.f7 = _rhs157;
      _lhs150.f25 = _rhs158;
      _lhs151[_lhs152] = _rhs159;
      for (const _guard_loop_9 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1180_v0).length))) {
        let _1190_i0 = _guard_loop_9;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1190_i0)) && ((_1190_i0).isLessThan(new BigNumber((_1180_v0).length))))) {
          (_1180_v0)[(_1190_i0)] = p0;
        }
      }
      let _hi5 = _1182_v2;
      for (let _1191_i1 = _1182_v2; _1191_i1.isLessThan(_hi5); _1191_i1 = _1191_i1.plus(_dafny.ONE)) {
        let _1192_v10;
        let _init32 = ((_1193_i1) => function (_1194_i2) {
          return (_1194_i2).plus(_1193_i1);
        })(_1191_i1);
        let _nw177 = Array((new BigNumber(15)).toNumber());
        for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw177.length); _i0_32++) {
          _nw177[_i0_32] = _init32(new BigNumber(_i0_32));
        }
        _1192_v10 = _nw177;
        let _1195_v11;
        _1195_v11 = _dafny.Seq.of(_1191_i1);
        let _index181 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_1192_v10).length));
        (_1192_v10)[_index181] = ((_1195_v11)[_module.__default.safeIndex(_1182_v2, new BigNumber((_1195_v11).length))]).plus(new BigNumber((_1184_v4).length));
        let _index182 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_1192_v10).length));
        (_1192_v10)[_index182] = (((_1188_v8).contains((_1180_v0)[_module.__default.safeIndex(new BigNumber(513), new BigNumber((_1180_v0).length))])) ? ((_1188_v8).get((_1180_v0)[_module.__default.safeIndex(new BigNumber(513), new BigNumber((_1180_v0).length))])) : (_1191_i1));
        let _1196_v12;
        _1196_v12 = _dafny.Map.Empty.slice().updateUnsafe(_1180_v0,_module.__default.safeModuloInt(_1182_v2, new BigNumber(475)));
        let _1197_v13;
        _1197_v13 = new _dafny.CodePoint('v'.codePointAt(0));
        let _1198_v14;
        _1198_v14 = _dafny.Map.Empty.slice().updateUnsafe((_1191_i1).multipliedBy(_1182_v2),(_1196_v12).update(_1180_v0, _1191_i1));
        let _index183 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_1180_v0).length));
        let _rhs160 = false;
        let _rhs161 = (((_1198_v14).contains((_1192_v10)[_module.__default.safeIndex(new BigNumber(387), new BigNumber((_1192_v10).length))])) ? ((_1198_v14).get((_1192_v10)[_module.__default.safeIndex(new BigNumber(387), new BigNumber((_1192_v10).length))])) : ((((_1198_v14).contains(_1182_v2)) ? ((_1198_v14).get(_1182_v2)) : (_1196_v12))));
        let _rhs162 = ((((_1185_v5).contains(new BigNumber(11))) ? ((_1185_v5).get(new BigNumber(11))) : (new BigNumber(-590)))).multipliedBy(_1182_v2);
        let _rhs163 = (_1186_v6)[_module.__default.safeIndex((_1192_v10)[_module.__default.safeIndex(new BigNumber(387), new BigNumber((_1192_v10).length))], new BigNumber((_1186_v6).length))];
        let _rhs164 = _1197_v13;
        let _lhs153 = _1180_v0;
        let _lhs154 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_1180_v0).length));
        let _lhs155 = globalState;
        let _lhs156 = globalState;
        _lhs153[_lhs154] = _rhs160;
        _1196_v12 = _rhs161;
        _lhs155.f7 = _rhs162;
        _lhs156.f21 = _rhs163;
        _1197_v13 = _rhs164;
        let _index184 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_1192_v10).length));
        (_1192_v10)[_index184] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(28), _1182_v2));
        _1188_v8 = ((_dafny.MultiSet.fromElements((_1180_v0)[_module.__default.safeIndex(new BigNumber(513), new BigNumber((_1180_v0).length))], (_this).fm9(p0, globalState))).update(_this.f25, _module.__default.abs(_1191_i1))).Difference((_1188_v8).Union(_1188_v8));
      }
      let _hi6 = (_dafny.ZERO).minus(_1182_v2);
      for (let _1199_i3 = _module.__default.safeModuloInt(new BigNumber(529), _1182_v2); _1199_i3.isLessThan(_hi6); _1199_i3 = _1199_i3.plus(_dafny.ONE)) {
        let _1200_v15;
        _1200_v15 = _dafny.Map.Empty.slice().updateUnsafe((_1182_v2).plus(_1199_i3),(_1180_v0)[_module.__default.safeIndex(new BigNumber(513), new BigNumber((_1180_v0).length))]);
        _1200_v15 = (_1200_v15).update(_1182_v2, p0);
        let _1201_v16;
        _1201_v16 = _dafny.MultiSet.fromElements(_1180_v0);
        let _1202_v17;
        let _nw178 = new _module.C2();
        _nw178.__ctor(new BigNumber((_1201_v16).cardinality()));
        _1202_v17 = _nw178;
        let _index185 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_1180_v0).length));
        (_1180_v0)[_index185] = (_1180_v0)[_module.__default.safeIndex(new BigNumber(513), new BigNumber((_1180_v0).length))];
        let _1203_v18;
        _1203_v18 = _module.D13.create_DC27(!(_this.f25));
        let _source19 = ((_this.f25) ? (_1203_v18) : (_1203_v18));
        if (_source19.is_DC26) {
          let _1204___mcc_h0 = (_source19).cf33;
          let _1205___mcc_h1 = (_source19).cf34;
          let _1206___mcc_h2 = (_source19).cf35;
          let _1207_cf35 = _1206___mcc_h2;
          let _1208_cf34 = _1205___mcc_h1;
          let _1209_cf33 = _1204___mcc_h0;
          let _1210_v19;
          _1210_v19 = new _dafny.CodePoint('f'.codePointAt(0));
          (globalState).f2 = !((_1186_v6)[_module.__default.safeIndex(_module.__default.safeDivisionInt((_1202_v17).f28, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1202_v17).f28,_1210_v19)).length)), new BigNumber((_1186_v6).length))]);
          let _1211_v20;
          let _nw179 = new _module.C2();
          _nw179.__ctor((new BigNumber(758)).minus(new BigNumber((_dafny.Seq.UnicodeFromString("tmdic")).length)));
          _1211_v20 = _nw179;
          let _1212_v21;
          let _nw180 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
          _1212_v21 = _nw180;
          (_1211_v20).m3(_1212_v21, globalState);
          let _1213_v22;
          _1213_v22 = _dafny.Seq.of(_1182_v2);
          let _1214_v23;
          _1214_v23 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("knj"),_dafny.Seq.Concat(_1213_v22, _1213_v22));
          _1214_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1209_cf33,_1213_v22);
        } else if (_source19.is_DC27) {
          let _1215___mcc_h3 = (_source19).cf36;
          let _1216_cf36 = _1215___mcc_h3;
          (globalState).f13 = _1216_cf36;
          (globalState).f7 = ((_1202_v17).f28).multipliedBy(_1182_v2);
          let _1217_v24;
          _1217_v24 = _dafny.Seq.of(_1182_v2, (_1202_v17).f28);
          let _1218_v25;
          _1218_v25 = _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(794)), ((_1219_i3) => function (_1220_i4) {
            return _1219_i3;
          })(_1199_i3)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-680)), ((_1221_v17) => function (_1222_i5) {
            return (_1221_v17).f28;
          })(_1202_v17))), ((_this.f25) ? (_1217_v24) : (_1217_v24)), _1217_v24);
          let _1223_v26;
          _1223_v26 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1200_v15).length),_1186_v6);
          let _1224_v27;
          _1224_v27 = new _dafny.CodePoint('p'.codePointAt(0));
          let _1225_v28;
          _1225_v28 = _dafny.Map.Empty.slice().updateUnsafe((_1180_v0)[_module.__default.safeIndex(new BigNumber(513), new BigNumber((_1180_v0).length))],_this.f25);
          let _1226_v29;
          let _nw181 = Array((new BigNumber(19)).toNumber());
          _nw181[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_1186_v6, _module.__default.safeIndex(new BigNumber(((_1223_v26).update(_module.__default.fm0(_1182_v2, _1182_v2, (_1202_v17).f28, false, globalState), _1186_v6)).length), new BigNumber((_1186_v6).length)), p0);
          _nw181[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_module.__default.fm18(_1216_cf36, _1224_v27, _this.f25, globalState), _module.__default.fm18((_this).f26, _1224_v27, p0, globalState));
          _nw181[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_1186_v6, _dafny.Seq.of(true, _module.__default.fm2(_1225_v28, globalState), true));
          _nw181[(new BigNumber(3)).toNumber()] = _1186_v6;
          _nw181[(new BigNumber(4)).toNumber()] = _1186_v6;
          _nw181[(new BigNumber(5)).toNumber()] = _1186_v6;
          _nw181[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_1186_v6, _1186_v6);
          _nw181[(new BigNumber(7)).toNumber()] = _1186_v6;
          _nw181[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of((_1180_v0)[_module.__default.safeIndex(new BigNumber(513), new BigNumber((_1180_v0).length))], !((_1180_v0)[_module.__default.safeIndex(new BigNumber(513), new BigNumber((_1180_v0).length))])), _module.__default.fm18(_1216_cf36, _1224_v27, p0, globalState));
          _nw181[(new BigNumber(9)).toNumber()] = _1186_v6;
          _nw181[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of((_1180_v0)[_module.__default.safeIndex(new BigNumber(513), new BigNumber((_1180_v0).length))]), _1186_v6);
          _nw181[(new BigNumber(11)).toNumber()] = _1186_v6;
          _nw181[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_1186_v6, _dafny.Seq.of((_1180_v0)[_module.__default.safeIndex(new BigNumber(513), new BigNumber((_1180_v0).length))]));
          _nw181[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_1186_v6, _1186_v6);
          _nw181[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_1186_v6, _1186_v6);
          _nw181[(new BigNumber(15)).toNumber()] = _dafny.Seq.Concat(_1186_v6, _1186_v6);
          _nw181[(new BigNumber(16)).toNumber()] = _1186_v6;
          _nw181[(new BigNumber(17)).toNumber()] = _dafny.Seq.of(_this.f25);
          _nw181[(new BigNumber(18)).toNumber()] = _dafny.Seq.of(p0);
          _1226_v29 = _nw181;
          let _index186 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_1226_v29).length));
          (_1226_v29)[_index186] = _1186_v6;
          let _1227_v30;
          _1227_v30 = _dafny.Seq.UnicodeFromString("ytqvm");
          let _1228_v31;
          _1228_v31 = _dafny.Set.fromElements(_1227_v30, _1227_v30);
          let _index187 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_1226_v29).length));
          let _rhs165 = false;
          let _rhs166 = _this.f25;
          let _rhs167 = _module.__default.fm22(_module.__default.fm2(_dafny.Map.Empty.slice().updateUnsafe(_1216_cf36,(_1180_v0)[_module.__default.safeIndex(new BigNumber(513), new BigNumber((_1180_v0).length))]), globalState), (_1228_v31).IsDisjointFrom(_1228_v31), p0, globalState);
          let _rhs168 = _dafny.Seq.Concat(_1186_v6, _1186_v6);
          let _lhs157 = _this;
          let _lhs158 = globalState;
          let _lhs159 = _1226_v29;
          let _lhs160 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_1226_v29).length));
          _lhs157.f25 = _rhs165;
          _lhs158.f13 = _rhs166;
          _1218_v25 = _rhs167;
          _lhs159[_lhs160] = _rhs168;
          let _1229_v32;
          let _nw182 = Array((new BigNumber(15)).toNumber());
          _nw182[(_dafny.ZERO).toNumber()] = new BigNumber(770);
          _nw182[(_dafny.ONE).toNumber()] = _1182_v2;
          _nw182[(new BigNumber(2)).toNumber()] = (_1202_v17).f28;
          _nw182[(new BigNumber(3)).toNumber()] = new BigNumber((_1227_v30).length);
          _nw182[(new BigNumber(4)).toNumber()] = (_1202_v17).f28;
          _nw182[(new BigNumber(5)).toNumber()] = (_1202_v17).f28;
          _nw182[(new BigNumber(6)).toNumber()] = (_1202_v17).f28;
          _nw182[(new BigNumber(7)).toNumber()] = _1199_i3;
          _nw182[(new BigNumber(8)).toNumber()] = _1182_v2;
          _nw182[(new BigNumber(9)).toNumber()] = new BigNumber(140);
          _nw182[(new BigNumber(10)).toNumber()] = _1182_v2;
          _nw182[(new BigNumber(11)).toNumber()] = (_1202_v17).f28;
          _nw182[(new BigNumber(12)).toNumber()] = _1199_i3;
          _nw182[(new BigNumber(13)).toNumber()] = _1199_i3;
          _nw182[(new BigNumber(14)).toNumber()] = (_1202_v17).f28;
          _1229_v32 = _nw182;
          let _1230_v33;
          _1230_v33 = _dafny.Set.fromElements(_1229_v32, _1229_v32, _1229_v32);
          (globalState).f13 = ((!((((_1225_v28).contains(true)) ? ((_1225_v28).get(true)) : (_this.f25)))) ? ((_1180_v0)[_module.__default.safeIndex(new BigNumber(513), new BigNumber((_1180_v0).length))]) : ((_1230_v33).contains(_1229_v32)));
        } else {
          let _1231___mcc_h4 = (_source19).cf32;
          let _1232_cf32 = _1231___mcc_h4;
          let _1233_v34;
          let _nw183 = Array((new BigNumber(29)).toNumber()).fill([]);
          _1233_v34 = _nw183;
          let _1234_v35;
          let _nw184 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
          _1234_v35 = _nw184;
          let _index188 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_1233_v34).length));
          (_1233_v34)[_index188] = _1234_v35;
          let _index189 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_1234_v35).length));
          (_1234_v35)[_index189] = (_dafny.ZERO).minus((((_1185_v5).contains((_1202_v17).f28)) ? ((_1185_v5).get((_1202_v17).f28)) : (_1199_i3)));
          let _1235_v36;
          let _init33 = function (_1236_i6) {
            return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(887)), function (_1237_i7) {
              return new _dafny.CodePoint('q'.codePointAt(0));
            }), _dafny.Seq.UnicodeFromString("yxeguqbr"));
          };
          let _nw185 = Array((new BigNumber(8)).toNumber());
          for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw185.length); _i0_33++) {
            _nw185[_i0_33] = _init33(new BigNumber(_i0_33));
          }
          _1235_v36 = _nw185;
          let _1238_v37;
          _1238_v37 = _dafny.Seq.UnicodeFromString("w");
          let _index190 = _module.__default.safeIndex(new BigNumber(540), new BigNumber((_1235_v36).length));
          (_1235_v36)[_index190] = _1238_v37;
          let _1239_v38;
          _1239_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1234_v35,new _dafny.CodePoint('d'.codePointAt(0)));
          let _1240_v39;
          _1240_v39 = _dafny.Map.Empty.slice().updateUnsafe(_this.f25,(_1180_v0)[_module.__default.safeIndex(new BigNumber(513), new BigNumber((_1180_v0).length))]);
          let _1241_v40;
          _1241_v40 = _dafny.Set.fromElements(_1234_v35);
          let _index191 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_1233_v34).length));
          let _index192 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_1234_v35).length));
          let _index193 = _module.__default.safeIndex(new BigNumber(540), new BigNumber((_1235_v36).length));
          let _rhs169 = _1234_v35;
          let _rhs170 = (_dafny.ZERO).minus(new BigNumber(((_1239_v38).Merge(_1239_v38)).length));
          let _rhs171 = _module.__default.fm2(_1240_v39, globalState);
          let _rhs172 = new BigNumber((((_1241_v40).Union(_1241_v40)).Intersect(_1241_v40)).length);
          let _rhs173 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("iaka"), _1238_v37);
          let _lhs161 = _1233_v34;
          let _lhs162 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_1233_v34).length));
          let _lhs163 = globalState;
          let _lhs164 = globalState;
          let _lhs165 = _1234_v35;
          let _lhs166 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_1234_v35).length));
          let _lhs167 = _1235_v36;
          let _lhs168 = _module.__default.safeIndex(new BigNumber(540), new BigNumber((_1235_v36).length));
          _lhs161[_lhs162] = _rhs169;
          _lhs163.f16 = _rhs170;
          _lhs164.f2 = _rhs171;
          _lhs165[_lhs166] = _rhs172;
          _lhs167[_lhs168] = _rhs173;
          let _1242_v41;
          _1242_v41 = _dafny.Set.fromElements((_1180_v0)[_module.__default.safeIndex(new BigNumber(513), new BigNumber((_1180_v0).length))], p0);
          let _1243_v42;
          _1243_v42 = _dafny.Map.Empty.slice().updateUnsafe(_1242_v41,!(_this.f25) || ((_1180_v0)[_module.__default.safeIndex(new BigNumber(513), new BigNumber((_1180_v0).length))]));
          _1243_v42 = (_1243_v42).update((_1242_v41).Intersect(_1242_v41), p0);
          (globalState).f7 = _module.__default.safeDivisionInt((_1202_v17).f28, (_1202_v17).f28);
          let _1244_v43;
          let _nw186 = new _module.C2();
          _nw186.__ctor((_1234_v35)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_1234_v35).length))]);
          _1244_v43 = _nw186;
        }
      }
      _1180_v0 = ((false) ? (_1180_v0) : (_1180_v0));
      let _1245_v44;
      _1245_v44 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      let _rhs174 = (_1186_v6)[_module.__default.safeIndex((new BigNumber(160)).multipliedBy(_1182_v2), new BigNumber((_1186_v6).length))];
      let _rhs175 = _dafny.MultiSet.fromElements(p0, (_1182_v2).isLessThan(new BigNumber((_1245_v44).length)), p0);
      let _lhs169 = globalState;
      _lhs169.f21 = _rhs174;
      _1188_v8 = _rhs175;
      return;
    }
    m6(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let r3 = false;
      let _1246_v0;
      _1246_v0 = _dafny.Seq.of(_module.__default.fm0(p3, new BigNumber(950), p3, _this.f25, globalState));
      let _1247_v1;
      let _init34 = function (_1248_i0) {
        return (((_this).f26) ? (_this.f25) : ((_this).f26));
      };
      let _nw187 = Array((_dafny.ONE).toNumber());
      for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw187.length); _i0_34++) {
        _nw187[_i0_34] = _init34(new BigNumber(_i0_34));
      }
      _1247_v1 = _nw187;
      let _1249_v2;
      _1249_v2 = _dafny.Map.Empty.slice().updateUnsafe(p3,_this.f25);
      let _1250_v5;
      _1250_v5 = _module.D1.create_DC2(function () {
  let _coll49 = new _dafny.Map();
  for (const _compr_49 of (_1246_v0).Elements) {
    let _1251_v4 = _compr_49;
    if (_dafny.Seq.contains(_1246_v0, _1251_v4)) {
      _coll49.push([(_1251_v4).plus(p3),_this.f25]);
    }
  }
  return _coll49;
}());
      let _1252_v7;
      _1252_v7 = _dafny.Seq.of((_this).f26, (_this).f26);
      let _1253_v8;
      _1253_v8 = _dafny.Map.Empty.slice().updateUnsafe(p3,_1249_v2);
      let _1254_v10;
      _1254_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,true);
      let _1255_v13;
      let _nw188 = Array((new BigNumber(23)).toNumber());
      _nw188[(_dafny.ZERO).toNumber()] = _1249_v2;
      _nw188[(_dafny.ONE).toNumber()] = _1249_v2;
      _nw188[(new BigNumber(2)).toNumber()] = _1249_v2;
      _nw188[(new BigNumber(3)).toNumber()] = function () {
        let _coll50 = new _dafny.Map();
        for (const _compr_50 of _dafny.IntegerRange(new BigNumber(306), new BigNumber(956))) {
          let _1256_v3 = _compr_50;
          if (((new BigNumber(306)).isLessThanOrEqualTo(_1256_v3)) && ((_1256_v3).isLessThan(new BigNumber(956)))) {
            _coll50.push([_module.__default.safeModuloInt(_1256_v3, p3),_this.f25]);
          }
        }
        return _coll50;
      }();
      _nw188[(new BigNumber(4)).toNumber()] = (_1250_v5).dtor_cf2;
      _nw188[(new BigNumber(5)).toNumber()] = (_1249_v2).Merge(function () {
        let _coll51 = new _dafny.Map();
        for (const _compr_51 of _dafny.IntegerRange(new BigNumber(766), new BigNumber(466))) {
          let _1257_v6 = _compr_51;
          if (((new BigNumber(766)).isLessThanOrEqualTo(_1257_v6)) && ((_1257_v6).isLessThan(new BigNumber(466)))) {
            _coll51.push([_module.__default.safeModuloInt(_1257_v6, p3),_this.f25]);
          }
        }
        return _coll51;
      }());
      _nw188[(new BigNumber(6)).toNumber()] = (_1249_v2).Merge(_dafny.Map.Empty.slice().updateUnsafe(p3,(_1252_v7)[_module.__default.safeIndex(p3, new BigNumber((_1252_v7).length))]));
      _nw188[(new BigNumber(7)).toNumber()] = (_1249_v2).Merge(_dafny.Map.Empty.slice().updateUnsafe(p3,(_this).f26));
      _nw188[(new BigNumber(8)).toNumber()] = _1249_v2;
      _nw188[(new BigNumber(9)).toNumber()] = _1249_v2;
      _nw188[(new BigNumber(10)).toNumber()] = _1249_v2;
      _nw188[(new BigNumber(11)).toNumber()] = ((((_1253_v8).contains(p3)) ? ((_1253_v8).get(p3)) : ((_dafny.Map.Empty.slice().updateUnsafe(p3,(_this).f26)).update(p3, true)))).Merge(function () {
        let _coll52 = new _dafny.Map();
        for (const _compr_52 of _dafny.IntegerRange(new BigNumber(658), new BigNumber(762))) {
          let _1258_v9 = _compr_52;
          if (((new BigNumber(658)).isLessThanOrEqualTo(_1258_v9)) && ((_1258_v9).isLessThan(new BigNumber(762)))) {
            _coll52.push([_module.__default.safeDivisionInt(_1258_v9, p3),(_this).f26]);
          }
        }
        return _coll52;
      }());
      _nw188[(new BigNumber(12)).toNumber()] = _1249_v2;
      _nw188[(new BigNumber(13)).toNumber()] = _1249_v2;
      _nw188[(new BigNumber(14)).toNumber()] = _1249_v2;
      _nw188[(new BigNumber(15)).toNumber()] = _1249_v2;
      _nw188[(new BigNumber(16)).toNumber()] = _1249_v2;
      _nw188[(new BigNumber(17)).toNumber()] = _1249_v2;
      _nw188[(new BigNumber(18)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1254_v10).length),true);
      _nw188[(new BigNumber(19)).toNumber()] = _1249_v2;
      _nw188[(new BigNumber(20)).toNumber()] = _1249_v2;
      _nw188[(new BigNumber(21)).toNumber()] = function () {
        let _coll53 = new _dafny.Map();
        for (const _compr_53 of _dafny.IntegerRange(new BigNumber(444), new BigNumber(576))) {
          let _1259_v11 = _compr_53;
          if (((new BigNumber(444)).isLessThanOrEqualTo(_1259_v11)) && ((_1259_v11).isLessThan(new BigNumber(576)))) {
            _coll53.push([_module.__default.safeModuloInt(_1259_v11, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f25,(_this).f26)).length)),_this.f25]);
          }
        }
        return _coll53;
      }();
      _nw188[(new BigNumber(22)).toNumber()] = function () {
        let _coll54 = new _dafny.Map();
        for (const _compr_54 of _dafny.IntegerRange(new BigNumber(788), new BigNumber(326))) {
          let _1260_v12 = _compr_54;
          if (((new BigNumber(788)).isLessThanOrEqualTo(_1260_v12)) && ((_1260_v12).isLessThan(new BigNumber(326)))) {
            _coll54.push([_module.__default.safeModuloInt(_1260_v12, p3),_this.f25]);
          }
        }
        return _coll54;
      }();
      _1255_v13 = _nw188;
      let _1261_v14;
      _1261_v14 = _dafny.MultiSet.fromElements(!(_this.f25) || (!((_this).f26)), (_this).fm9(_this.f25, globalState), (_this).f26, !(_module.__default.fm2(_1254_v10, globalState)), ((((_1249_v2).contains(p3)) ? ((_1249_v2).get(p3)) : ((_this).f26))) || (_this.f25));
      let _1262_v15;
      _1262_v15 = _dafny.Seq.of(_1247_v1);
      let _rhs176 = _1246_v0;
      let _rhs177 = (_1262_v15)[_module.__default.safeIndex((p3).plus(p3), new BigNumber((_1262_v15).length))];
      let _rhs178 = _1255_v13;
      let _rhs179 = _1261_v14;
      _1246_v0 = _rhs176;
      _1247_v1 = _rhs177;
      _1255_v13 = _rhs178;
      _1261_v14 = _rhs179;
      let _1263_v16;
      _1263_v16 = _module.D12.create_DC23((_dafny.ZERO).minus(p3));
      let _1264_v17;
      _1264_v17 = _dafny.Map.Empty.slice().updateUnsafe(_this.f25,_1263_v16);
      _1264_v17 = _1264_v17;
      r1 = p3;
      let _1265_v18;
      let _nw189 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
      _1265_v18 = _nw189;
      _1265_v18 = p0;
      let _1266_i1;
      _1266_i1 = _dafny.ZERO;
      L7: {
        while ((_this).fm9((_this.f25) || (_this.f25), globalState)) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1266_i1)) {
              break L7;
            }
            _1266_i1 = (_1266_i1).plus(_dafny.ONE);
            let _1267_v19;
            let _nw190 = new _module.C3();
            _nw190.__ctor();
            _1267_v19 = _nw190;
            if (!_dafny.Seq.contains(p1, new _dafny.CodePoint('u'.codePointAt(0)))) {
              _1247_v1 = _1247_v1;
              let _index194 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_1247_v1).length));
              (_1247_v1)[_index194] = _this.f25;
              let _index195 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_1247_v1).length));
              let _rhs180 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("iongp"), _dafny.Seq.Concat(p1, p1));
              let _rhs181 = (_this).f26;
              let _rhs182 = new BigNumber((_dafny.Seq.Concat(p1, _dafny.Seq.UnicodeFromString("osab"))).length);
              let _lhs170 = _this;
              let _lhs171 = _1247_v1;
              let _lhs172 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_1247_v1).length));
              let _lhs173 = globalState;
              _lhs170.f25 = _rhs180;
              _lhs171[_lhs172] = _rhs181;
              _lhs173.f16 = _rhs182;
              let _1268_v20;
              _1268_v20 = new _dafny.CodePoint('q'.codePointAt(0));
              let _1269_v21;
              _1269_v21 = _dafny.Map.Empty.slice().updateUnsafe(!((_1247_v1)[_module.__default.safeIndex(new BigNumber(813), new BigNumber((_1247_v1).length))]),_module.__default.fm0(p3, new BigNumber(169), p3, false, globalState));
              let _1270_v22;
              _1270_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1268_v20,(_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(((_1269_v21).Merge(_1269_v21)).length)))));
              _1270_v22 = (_1270_v22).update(_1268_v20, (new BigNumber(-25)).minus((((_1269_v21).contains((_1247_v1)[_module.__default.safeIndex(new BigNumber(813), new BigNumber((_1247_v1).length))])) ? ((_1269_v21).get((_1247_v1)[_module.__default.safeIndex(new BigNumber(813), new BigNumber((_1247_v1).length))])) : (p3))));
              let _1271_v23;
              let _init35 = ((_1272_v21) => function (_1273_i2) {
                return _1272_v21;
              })(_1269_v21);
              let _nw191 = Array((new BigNumber(23)).toNumber());
              for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw191.length); _i0_35++) {
                _nw191[_i0_35] = _init35(new BigNumber(_i0_35));
              }
              _1271_v23 = _nw191;
              _1271_v23 = _1271_v23;
              let _1274_v24;
              let _nw192 = Array((_dafny.ONE).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
              _1274_v24 = _nw192;
              _1274_v24 = _1274_v24;
            } else {
              let _1275_v25;
              _1275_v25 = _dafny.Map.Empty.slice().updateUnsafe(p3,(p3).multipliedBy((_dafny.ZERO).minus(p3)));
              let _1276_v27;
              _1276_v27 = _dafny.Map.Empty.slice().updateUnsafe(function () {
                let _coll55 = new _dafny.Set();
                for (const _compr_55 of _dafny.IntegerRange(new BigNumber(-718), new BigNumber(-647))) {
                  let _1277_v26 = _compr_55;
                  if (((new BigNumber(-718)).isLessThanOrEqualTo(_1277_v26)) && ((_1277_v26).isLessThan(new BigNumber(-647)))) {
                    _coll55.add(_module.__default.safeDivisionInt(_1277_v26, p3));
                  }
                }
                return _coll55;
              }(),_this.f25);
              _1275_v25 = (_1275_v25).update(new BigNumber((_1276_v27).length), p3);
              let _1278_v28;
              _1278_v28 = _module.D14.create_DC30(_dafny.Seq.Concat(_1246_v0, _1246_v0), _dafny.Seq.Concat(_1252_v7, _1252_v7));
              _1278_v28 = _1278_v28;
              let _1279_v29;
              let _nw193 = Array((new BigNumber(7)).toNumber());
              _nw193[(_dafny.ZERO).toNumber()] = p2;
              _nw193[(_dafny.ONE).toNumber()] = p2;
              _nw193[(new BigNumber(2)).toNumber()] = p0;
              _nw193[(new BigNumber(3)).toNumber()] = p2;
              _nw193[(new BigNumber(4)).toNumber()] = p2;
              _nw193[(new BigNumber(5)).toNumber()] = p0;
              _nw193[(new BigNumber(6)).toNumber()] = p2;
              _1279_v29 = _nw193;
              let _index196 = _module.__default.safeIndex(new BigNumber(638), new BigNumber((_1279_v29).length));
              (_1279_v29)[_index196] = p2;
              let _index197 = _module.__default.safeIndex(new BigNumber(638), new BigNumber((_1279_v29).length));
              (_1279_v29)[_index197] = _1265_v18;
              let _1280_v30;
              let _nw194 = Array((new BigNumber(11)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
              _1280_v30 = _nw194;
              let _1281_v31;
              _1281_v31 = new _dafny.CodePoint('r'.codePointAt(0));
              let _index198 = _module.__default.safeIndex(new BigNumber(501), new BigNumber((_1280_v30).length));
              (_1280_v30)[_index198] = _1281_v31;
              let _index199 = _module.__default.safeIndex(new BigNumber(501), new BigNumber((_1280_v30).length));
              (_1280_v30)[_index199] = _1281_v31;
              let _1282_v33;
              _1282_v33 = _dafny.Seq.of(function () {
                let _coll56 = new _dafny.Map();
                for (const _compr_56 of _dafny.IntegerRange(new BigNumber(-929), new BigNumber(930))) {
                  let _1283_v32 = _compr_56;
                  if (((new BigNumber(-929)).isLessThanOrEqualTo(_1283_v32)) && ((_1283_v32).isLessThan(new BigNumber(930)))) {
                    _coll56.push([(_1283_v32).multipliedBy(new BigNumber((_dafny.Set.fromElements(p3, p3)).length)),_this.f25]);
                  }
                }
                return _coll56;
              }());
              let _1284_v34;
              let _nw195 = new _module.C1();
              _nw195.__ctor(new BigNumber((p1).length), _1282_v33);
              _1284_v34 = _nw195;
              let _1285_v35;
              _1285_v35 = _dafny.Map.Empty.slice().updateUnsafe(_1284_v34,p3);
              _1285_v35 = (_1285_v35).update(_1284_v34, p3);
            }
            (globalState).f13 = _this.f25;
            let _1286_v36;
            let _nw196 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
            _1286_v36 = _nw196;
            let _index200 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_1286_v36).length));
            (_1286_v36)[_index200] = p1;
            let _1287_v37;
            _1287_v37 = new _dafny.CodePoint('d'.codePointAt(0));
            let _index201 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_1286_v36).length));
            (_1286_v36)[_index201] = _dafny.Seq.Concat(_dafny.Seq.Concat(p1, _dafny.Seq.of(_1287_v37, _1287_v37)), p1);
          }
        }
      }
      let _1288_i3;
      _1288_i3 = _dafny.ZERO;
      L8: {
        while (((_this.f25) ? ((_this).f26) : (true))) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1288_i3)) {
              break L8;
            }
            _1288_i3 = (_1288_i3).plus(_dafny.ONE);
            (globalState).f21 = _this.f25;
            let _1289_v39;
            _1289_v39 = _dafny.Seq.of(function () {
              let _coll57 = new _dafny.Map();
              for (const _compr_57 of _dafny.IntegerRange(new BigNumber(888), new BigNumber(272))) {
                let _1290_v38 = _compr_57;
                if (((new BigNumber(888)).isLessThanOrEqualTo(_1290_v38)) && ((_1290_v38).isLessThan(new BigNumber(272)))) {
                  _coll57.push([_module.__default.safeDivisionInt(_1290_v38, p3),(_this).f26]);
                }
              }
              return _coll57;
            }());
            let _1291_v40;
            let _nw197 = new _module.C1();
            _nw197.__ctor(p3, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(483)), ((_1292_v2) => function (_1293_i4) {
              return _1292_v2;
            })(_1249_v2)), _1289_v39));
            _1291_v40 = _nw197;
            r1 = (new BigNumber(-433)).multipliedBy(_module.__default.safeDivisionInt(p3, p3));
            let _1294_v41;
            _1294_v41 = _dafny.Set.fromElements((_this).f26, _this.f25);
            let _1295_v42;
            _1295_v42 = _module.D4.create_DC7(new BigNumber((_1294_v41).length));
            (_1291_v40).f29 = (_1295_v42).dtor_cf9;
          }
        }
      }
      r0 = !(_this.f25);
      r1 = (p3).multipliedBy((new BigNumber((_1252_v7).length)).plus(p3));
      let _1296_v43;
      _1296_v43 = _module.D18.create_DC42(_1254_v10);
      r2 = _module.__default.fm2((_1296_v43).dtor_cf58, globalState);
      let _1297_v44;
      _1297_v44 = _dafny.Set.fromElements(p2, p2);
      r3 = (_1297_v44).IsProperSubsetOf(_1297_v44);
      return [r0, r1, r2, r3];
    }
    get f26() {
      let _this = this;
      return _this._f26;
    };
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f24 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    __ctor(f24) {
      let _this = this;
      (_this)._f24 = f24;
      return;
    }
    fm7(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f24;
    };
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((_dafny.Map.Empty.slice().updateUnsafe((_this).f24,_dafny.Seq.UnicodeFromString("wx"))).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f24,_dafny.Seq.UnicodeFromString("mjsnrtion")))).Merge((_dafny.Map.Empty.slice().updateUnsafe((_this).f24,_dafny.Seq.Create(_module.__default.abs(new BigNumber(378)), function (_1298_i0) {
        return new _dafny.CodePoint('l'.codePointAt(0));
      }))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("sjvmhb"))));
    };
    fm6(p0, globalState) {
      let _this = this;
      let _source20 = _module.D3.create_DC5(new _dafny.CodePoint('x'.codePointAt(0)));
      if (_source20.is_DC6) {
        let _1299___mcc_h0 = (_source20).cf6;
        let _1300___mcc_h1 = (_source20).cf7;
        let _1301___mcc_h2 = (_source20).cf8;
        let _1302_cf8 = _1301___mcc_h2;
        let _1303_cf7 = _1300___mcc_h1;
        let _1304_cf6 = _1299___mcc_h0;
        return (_this).f24;
      } else {
        let _1305___mcc_h3 = (_source20).cf5;
        let _1306_cf5 = _1305___mcc_h3;
        return ((_this).f24) || ((_this).f24);
      }
    };
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('t'.codePointAt(0)))).cardinality())),(_this).f24)).length)).minus(new BigNumber(-530))).multipliedBy(new BigNumber(86)));
    };
    m2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = [];
      let _1307_v0;
      let _nw198 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _1307_v0 = _nw198;
      for (const _guard_loop_10 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1307_v0).length))) {
        let _1308_i0 = _guard_loop_10;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1308_i0)) && ((_1308_i0).isLessThan(new BigNumber((_1307_v0).length))))) {
          (_1307_v0)[(_1308_i0)] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ckiixp"), _dafny.Seq.UnicodeFromString("ogsr"));
        }
      }
      let _1309_v1;
      _1309_v1 = _module.D4.create_DC7(p3);
      (globalState).f13 = (_this).fm7(p1, (p3).minus(p3), !(new BigNumber(215)).isEqualTo((_1309_v1).dtor_cf9), globalState);
      if (!(new BigNumber(852)).isEqualTo(p0)) {
        let _1310_v2;
        _1310_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,p1);
        (globalState).f0 = _module.__default.fm2(_1310_v2, globalState);
        let _1311_v3;
        _1311_v3 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
        let _1312_v4;
        _1312_v4 = _dafny.Set.fromElements((_this).fm6(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1311_v3,_module.__default.fm0(p3, p0, p3, p1, globalState))).length), globalState));
        r1 = (p3).minus(new BigNumber(((_dafny.Set.fromElements((_this).f24, p1)).Difference(_1312_v4)).length));
        (globalState).f7 = _module.__default.safeModuloInt((new BigNumber(412)).multipliedBy(p3), p0);
        let _1313_v5;
        let _nw199 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
        _1313_v5 = _nw199;
        let _1314_v6;
        _1314_v6 = _dafny.Seq.UnicodeFromString("bqijp");
        let _1315_v7;
        _1315_v7 = new _dafny.CodePoint('s'.codePointAt(0));
        let _1316_v8;
        _1316_v8 = _dafny.Seq.of(!(true));
        let _rhs183 = _module.__default.safeDivisionInt(new BigNumber((_1311_v3).length), p0);
        let _rhs184 = !(((p1) ? (p1) : (p1))) || ((_this).f24);
        let _rhs185 = (_this).fm7(!(p1) || (p1), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("dcbbvmlo"), _dafny.Seq.update(_1314_v6, _module.__default.safeIndex(p0, new BigNumber((_1314_v6).length)), _1315_v7))).length), true, globalState);
        let _rhs186 = (_1316_v8)[_module.__default.safeIndex((_1309_v1).dtor_cf9, new BigNumber((_1316_v8).length))];
        let _rhs187 = _1313_v5;
        let _lhs174 = globalState;
        let _lhs175 = globalState;
        let _lhs176 = globalState;
        _lhs174.f16 = _rhs183;
        _lhs175.f0 = _rhs184;
        r0 = _rhs185;
        _lhs176.f0 = _rhs186;
        _1313_v5 = _rhs187;
        let _1317_v9;
        let _nw200 = new _module.C3();
        _nw200.__ctor();
        _1317_v9 = _nw200;
      } else {
        let _1318_v10;
        let _nw201 = Array((new BigNumber(8)).toNumber());
        _nw201[(_dafny.ZERO).toNumber()] = p2;
        _nw201[(_dafny.ONE).toNumber()] = p2;
        _nw201[(new BigNumber(2)).toNumber()] = p2;
        _nw201[(new BigNumber(3)).toNumber()] = p2;
        _nw201[(new BigNumber(4)).toNumber()] = p2;
        _nw201[(new BigNumber(5)).toNumber()] = p2;
        _nw201[(new BigNumber(6)).toNumber()] = p2;
        _nw201[(new BigNumber(7)).toNumber()] = p2;
        _1318_v10 = _nw201;
        _1318_v10 = _1318_v10;
        let _index202 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((p2).length));
        (p2)[_index202] = p3;
        let _index203 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((p2).length));
        let _rhs188 = (p3).minus(p0);
        let _rhs189 = p3;
        let _rhs190 = p3;
        let _lhs177 = globalState;
        let _lhs178 = globalState;
        let _lhs179 = p2;
        let _lhs180 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((p2).length));
        _lhs177.f16 = _rhs188;
        _lhs178.f16 = _rhs189;
        _lhs179[_lhs180] = _rhs190;
        let _1319_v11;
        _1319_v11 = new _dafny.CodePoint('f'.codePointAt(0));
        _1319_v11 = new _dafny.CodePoint('v'.codePointAt(0));
        let _index204 = _module.__default.safeIndex(new BigNumber(572), new BigNumber((p2).length));
        (p2)[_index204] = p3;
        let _1320_v12;
        let _init36 = ((_1321_p1, _1322_p3) => function (_1323_i1) {
          return _module.D12.create_DC22(_dafny.Map.Empty.slice().updateUnsafe(_1321_p1,_1322_p3));
        })(p1, p3);
        let _nw202 = Array((new BigNumber(11)).toNumber());
        for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw202.length); _i0_36++) {
          _nw202[_i0_36] = _init36(new BigNumber(_i0_36));
        }
        _1320_v12 = _nw202;
        let _1324_v13;
        _1324_v13 = _dafny.Map.Empty.slice().updateUnsafe((p2)[_module.__default.safeIndex(new BigNumber(572), new BigNumber((p2).length))],_1320_v12);
        _1324_v13 = (_1324_v13).update(p0, _1320_v12);
      }
      if (!(p1)) {
        let _1325_v14;
        _1325_v14 = _dafny.Seq.UnicodeFromString("gus");
        (globalState).f19 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1325_v14, _1325_v14), _1325_v14);
        let _1326_v15;
        let _init37 = ((_1327_p0, _1328_p3) => function (_1329_i2) {
          return _dafny.Seq.of(_1327_p0, _1327_p0, _1327_p0, new BigNumber(-190), _1328_p3);
        })(p0, p3);
        let _nw203 = Array((new BigNumber(21)).toNumber());
        for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw203.length); _i0_37++) {
          _nw203[_i0_37] = _init37(new BigNumber(_i0_37));
        }
        _1326_v15 = _nw203;
        let _1330_v16;
        _1330_v16 = _module.D15.create_DC32(_1326_v15);
        let _1331_v17;
        _1331_v17 = _module.D15.create_DC34(_1330_v16);
        let _1332_v18;
        _1332_v18 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1331_v17);
        _1332_v18 = (_1332_v18).update(p1, _1331_v17);
        let _1333_v19;
        _1333_v19 = new _dafny.CodePoint('n'.codePointAt(0));
        let _1334_v20;
        _1334_v20 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Concat(_1325_v14, _dafny.Seq.update(_1325_v14, _module.__default.safeIndex(p3, new BigNumber((_1325_v14).length)), _1333_v19))).length),p3);
        let _1335_v21;
        _1335_v21 = _dafny.Seq.of(_1334_v20);
        _1334_v20 = ((_1334_v20).Merge((_1335_v21)[_module.__default.safeIndex(p0, new BigNumber((_1335_v21).length))])).Merge(_1334_v20);
        r1 = p0;
        (globalState).f13 = p1;
      } else {
        let _index205 = _module.__default.safeIndex(new BigNumber(96), new BigNumber((p2).length));
        (p2)[_index205] = p3;
        let _1336_v22;
        _1336_v22 = _dafny.Set.fromElements(p2, p2, p2);
        let _index206 = _module.__default.safeIndex(new BigNumber(96), new BigNumber((p2).length));
        (p2)[_index206] = new BigNumber((_1336_v22).length);
        let _1337_v23;
        _1337_v23 = _dafny.Seq.UnicodeFromString("coillmrw");
        let _1338_v24;
        _1338_v24 = _dafny.MultiSet.fromElements(_1337_v23);
        let _1339_v25;
        _1339_v25 = _dafny.Map.Empty.slice().updateUnsafe(p0,new _dafny.CodePoint('l'.codePointAt(0)));
        let _1340_v26;
        _1340_v26 = new _dafny.CodePoint('k'.codePointAt(0));
        let _1341_v27;
        _1341_v27 = _dafny.MultiSet.fromElements(((!((_this).f24)) ? (new BigNumber((_1338_v24).cardinality())) : (new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(183)), function (_1342_i3) {
          return new _dafny.CodePoint('c'.codePointAt(0));
        }), _module.__default.safeIndex((_this).fm8((_this).f24, (_1339_v25).update(p3, _1340_v26), _1337_v23, globalState), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(183)), function (_1343_i3) {
          return new _dafny.CodePoint('c'.codePointAt(0));
        })).length)), _module.__default.fm27((_this).f24, !(false), globalState))).length))));
        let _1344_v28;
        _1344_v28 = _dafny.Map.Empty.slice().updateUnsafe((p2)[_module.__default.safeIndex(new BigNumber(96), new BigNumber((p2).length))],p0);
        let _1345_v29;
        _1345_v29 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f24);
        let _index207 = _module.__default.safeIndex(new BigNumber(96), new BigNumber((p2).length));
        (p2)[_index207] = (((_1341_v27).contains(p3)) ? ((_1341_v27).get(p3)) : ((((_1344_v28).contains(new BigNumber(881))) ? ((_1344_v28).get(new BigNumber(881))) : (new BigNumber((_1345_v29).length)))));
        let _1346_v30;
        let _init38 = ((_1347_p0, _1348_p2) => function (_1349_i4) {
          return ((_dafny.ZERO).minus(_1347_p0)).isLessThanOrEqualTo((_1348_p2)[_module.__default.safeIndex(new BigNumber(96), new BigNumber((_1348_p2).length))]);
        })(p0, p2);
        let _nw204 = Array((new BigNumber(28)).toNumber());
        for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw204.length); _i0_38++) {
          _nw204[_i0_38] = _init38(new BigNumber(_i0_38));
        }
        _1346_v30 = _nw204;
        let _index208 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_1346_v30).length));
        (_1346_v30)[_index208] = p1;
        let _1350_v31;
        _1350_v31 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(488)), ((_1351_v26) => function (_1352_i5) {
          return _1351_v26;
        })(_1340_v26)),p0);
        let _1353_v32;
        let _nw205 = new _module.C4();
        _nw205.__ctor(false, p1);
        _1353_v32 = _nw205;
        let _1354_v33;
        _1354_v33 = _dafny.Seq.of((_1353_v32).f26);
        let _1355_v34;
        _1355_v34 = _dafny.Seq.of(p1, (_1354_v33)[_module.__default.safeIndex(p0, new BigNumber((_1354_v33).length))], p1);
        let _1356_v35;
        _1356_v35 = _dafny.Seq.of(p1, _1353_v32.f25);
        let _1357_v36;
        _1357_v36 = _dafny.Map.Empty.slice().updateUnsafe(_1356_v35,_1340_v26);
        let _index209 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_1346_v30).length));
        let _rhs191 = true;
        let _rhs192 = _1346_v30;
        let _rhs193 = ((p1) ? ((new BigNumber((_1350_v31).length)).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_1353_v32, _1353_v32, _1353_v32, _1353_v32)).length)))) : (_module.__default.safeDivisionInt(new BigNumber((_1354_v33).length), new BigNumber((_dafny.Seq.of(p1)).length))));
        let _rhs194 = (((_1357_v36).contains(_module.__default.fm41(!(true), _dafny.Set.fromElements(_module.__default.fm2(_1345_v29, globalState), true, (_1353_v32).f26, _1353_v32.f25), (p2)[_module.__default.safeIndex(new BigNumber(96), new BigNumber((p2).length))], p3, globalState))) ? ((_1357_v36).get(_module.__default.fm41(!(true), _dafny.Set.fromElements(_module.__default.fm2(_1345_v29, globalState), true, (_1353_v32).f26, _1353_v32.f25), (p2)[_module.__default.safeIndex(new BigNumber(96), new BigNumber((p2).length))], p3, globalState))) : (_1340_v26));
        let _lhs181 = _1346_v30;
        let _lhs182 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_1346_v30).length));
        let _lhs183 = globalState;
        _lhs181[_lhs182] = _rhs191;
        _1346_v30 = _rhs192;
        _lhs183.f7 = _rhs193;
        _1340_v26 = _rhs194;
        (globalState).f8 = _1337_v23;
        let _1358_v37;
        let _init39 = ((_1359_v32) => function (_1360_i6) {
          return _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1359_v32.f25);
        })(_1353_v32);
        let _nw206 = Array((new BigNumber(23)).toNumber());
        for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw206.length); _i0_39++) {
          _nw206[_i0_39] = _init39(new BigNumber(_i0_39));
        }
        _1358_v37 = _nw206;
        _1358_v37 = (((_1353_v32).f26) ? (_1358_v37) : (_1358_v37));
      }
      if (!(!(p1)) || (!(p1))) {
        let _1361_v38;
        _1361_v38 = new _dafny.CodePoint('m'.codePointAt(0));
        _1361_v38 = _1361_v38;
        let _1362_v39;
        _1362_v39 = _dafny.Seq.UnicodeFromString("i");
        let _index210 = _module.__default.safeIndex(new BigNumber(453), new BigNumber((_1307_v0).length));
        (_1307_v0)[_index210] = _1362_v39;
        let _index211 = _module.__default.safeIndex(new BigNumber(453), new BigNumber((_1307_v0).length));
        (_1307_v0)[_index211] = _dafny.Seq.UnicodeFromString("woo");
        let _1363_v40;
        let _nw207 = Array((new BigNumber(29)).toNumber()).fill([]);
        _1363_v40 = _nw207;
        _1363_v40 = _1363_v40;
        let _1364_v41;
        _1364_v41 = _dafny.Map.Empty.slice().updateUnsafe(_1361_v38,p0);
        (globalState).f19 = _dafny.Seq.update(_dafny.Seq.update(_1362_v39, _module.__default.safeIndex((((_1364_v41).contains(_1361_v38)) ? ((_1364_v41).get(_1361_v38)) : ((_dafny.ZERO).minus(p3))), new BigNumber((_1362_v39).length)), _1361_v38), _module.__default.safeIndex(p3, new BigNumber((_dafny.Seq.update(_1362_v39, _module.__default.safeIndex((((_1364_v41).contains(_1361_v38)) ? ((_1364_v41).get(_1361_v38)) : ((_dafny.ZERO).minus(p3))), new BigNumber((_1362_v39).length)), _1361_v38)).length)), _1361_v38);
        (globalState).f13 = (_this).f24;
      } else {
        let _1365_v42;
        let _nw208 = Array((_dafny.ONE).toNumber()).fill([]);
        _1365_v42 = _nw208;
        let _index212 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_1365_v42).length));
        (_1365_v42)[_index212] = p2;
        let _index213 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((p2).length));
        (p2)[_index213] = p0;
        let _1366_v43;
        _1366_v43 = _dafny.Seq.UnicodeFromString("krdj");
        let _index214 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_1365_v42).length));
        let _index215 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((p2).length));
        let _rhs195 = (_dafny.ZERO).minus(p0);
        let _rhs196 = p2;
        let _rhs197 = ((p1) ? (new BigNumber((_1366_v43).length)) : (p3));
        let _lhs184 = _1365_v42;
        let _lhs185 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_1365_v42).length));
        let _lhs186 = p2;
        let _lhs187 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((p2).length));
        r1 = _rhs195;
        _lhs184[_lhs185] = _rhs196;
        _lhs186[_lhs187] = _rhs197;
        let _1367_v44;
        let _init40 = ((_1368_p1) => function (_1369_i7) {
          return _1368_p1;
        })(p1);
        let _nw209 = Array((new BigNumber(11)).toNumber());
        for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw209.length); _i0_40++) {
          _nw209[_i0_40] = _init40(new BigNumber(_i0_40));
        }
        _1367_v44 = _nw209;
        let _1370_v45;
        let _nw210 = Array((new BigNumber(11)).toNumber());
        _nw210[(_dafny.ZERO).toNumber()] = _1367_v44;
        _nw210[(_dafny.ONE).toNumber()] = _1367_v44;
        _nw210[(new BigNumber(2)).toNumber()] = _1367_v44;
        _nw210[(new BigNumber(3)).toNumber()] = _1367_v44;
        _nw210[(new BigNumber(4)).toNumber()] = _1367_v44;
        _nw210[(new BigNumber(5)).toNumber()] = _1367_v44;
        _nw210[(new BigNumber(6)).toNumber()] = _1367_v44;
        _nw210[(new BigNumber(7)).toNumber()] = _1367_v44;
        _nw210[(new BigNumber(8)).toNumber()] = _1367_v44;
        _nw210[(new BigNumber(9)).toNumber()] = _1367_v44;
        _nw210[(new BigNumber(10)).toNumber()] = _1367_v44;
        _1370_v45 = _nw210;
        let _1371_v46;
        _1371_v46 = _1367_v44;
        let _index216 = _module.__default.safeIndex(new BigNumber(879), new BigNumber((_1370_v45).length));
        (_1370_v45)[_index216] = (_1371_v46);
        let _1372_v47;
        _1372_v47 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f24);
        let _1373_v48;
        _1373_v48 = _module.D18.create_DC42((_1372_v47).update(p1, (_this).f24));
        let _1374_v49;
        _1374_v49 = _dafny.Seq.of((_this).f24, p1, true, p1);
        let _1375_v50;
        _1375_v50 = _dafny.Seq.of((_1372_v47).Merge((_1373_v48).dtor_cf58), _1372_v47, (_dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f24)).update((_module.__default.fm25(true, p0, p1, new BigNumber((_1374_v49).length), globalState)).dtor_cf8, (_this).f24));
        let _index217 = _module.__default.safeIndex(new BigNumber(879), new BigNumber((_1370_v45).length));
        let _rhs198 = ((p1) ? (_1367_v44) : (_1367_v44));
        let _rhs199 = p1;
        let _rhs200 = p1;
        let _rhs201 = !(_dafny.areEqual(_1366_v43, _1366_v43));
        let _rhs202 = (_1375_v50)[_module.__default.safeIndex(p3, new BigNumber((_1375_v50).length))];
        let _lhs188 = _1370_v45;
        let _lhs189 = _module.__default.safeIndex(new BigNumber(879), new BigNumber((_1370_v45).length));
        let _lhs190 = globalState;
        let _lhs191 = globalState;
        let _lhs192 = globalState;
        _lhs188[_lhs189] = _rhs198;
        _lhs190.f13 = _rhs199;
        _lhs191.f13 = _rhs200;
        _lhs192.f2 = _rhs201;
        _1372_v47 = _rhs202;
        (globalState).f0 = (_this).fm6(p3, globalState);
        let _1376_v51;
        _1376_v51 = _dafny.Map.Empty.slice().updateUnsafe(false,(p2)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((p2).length))]);
        let _1377_v52;
        _1377_v52 = _module.D4.create_DC10(new BigNumber((_1376_v51).length), (p2)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((p2).length))], (_this).f24, p1, (p2)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((p2).length))]);
        let _1378_v53;
        _1378_v53 = _dafny.Seq.of(_1377_v52, _module.__default.fm42(globalState));
        let _1379_v54;
        _1379_v54 = _dafny.MultiSet.fromElements(p1);
        let _1380_v55;
        _1380_v55 = _module.D4.create_DC10(new BigNumber((_1378_v53).length), p3, p1, (_module.D16.create_DC36(!(p1), _1379_v54, p1)).dtor_cf49, (p2)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((p2).length))]);
        let _1381_v56;
        _1381_v56 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1380_v55);
        let _1382_v57;
        _1382_v57 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1381_v56);
        _1382_v57 = _dafny.Map.Empty.slice().updateUnsafe(p0,_module.__default.fm43((_this).f24, p3, new BigNumber((_1366_v43).length), p1, globalState));
        let _1383_v58;
        _1383_v58 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_1365_v42)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_1365_v42).length))]);
        _1383_v58 = (_1383_v58).update((_this).f24, (_1365_v42)[_module.__default.safeIndex(new BigNumber(201), new BigNumber((_1365_v42).length))]);
      }
      (globalState).f16 = p3;
      r0 = !(((((_this).f24) ? (false) : (p1))) === ((_this).f24));
      r1 = (_dafny.ZERO).minus(p0);
      let _init41 = function (_1384_i8) {
        return _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_module.D3.create_DC6(_dafny.Seq.UnicodeFromString("ysrfvc"), new _dafny.CodePoint('w'.codePointAt(0)), false));
      };
      let _nw211 = Array((new BigNumber(6)).toNumber());
      for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw211.length); _i0_41++) {
        _nw211[_i0_41] = _init41(new BigNumber(_i0_41));
      }
      r2 = _nw211;
      return [r0, r1, r2];
    }
    m3(p0, globalState) {
      let _this = this;
      let _1385_v0;
      _1385_v0 = new BigNumber(48);
      let _1386_v1;
      _1386_v1 = _dafny.Set.fromElements(_1385_v0);
      let _1387_v2;
      let _nw212 = new _module.C4();
      _nw212.__ctor(!(_1386_v1).equals(_1386_v1), (_this).f24);
      _1387_v2 = _nw212;
      let _1388_v3;
      _1388_v3 = _dafny.MultiSet.fromElements((_this).f24);
      let _1389_v4;
      _1389_v4 = _module.D16.create_DC36((_1387_v2).f26, _1388_v3, (_1387_v2).f26);
      let _source21 = _1389_v4;
      if (_source21.is_DC36) {
        let _1390___mcc_h0 = (_source21).cf47;
        let _1391___mcc_h1 = (_source21).cf48;
        let _1392___mcc_h2 = (_source21).cf49;
        let _1393_cf49 = _1392___mcc_h2;
        let _1394_cf48 = _1391___mcc_h1;
        let _1395_cf47 = _1390___mcc_h0;
        (globalState).f16 = _1385_v0;
        let _1396_v5;
        _1396_v5 = _dafny.Seq.of(_1393_cf49, _1395_cf47);
        _1394_cf48 = (((_1395_cf47) ? (_1388_v3) : (_1394_cf48))).Difference(_dafny.MultiSet.FromArray(_1396_v5));
        let _1397_v6;
        let _nw213 = new _module.C0();
        _nw213.__ctor((_1396_v5)[_module.__default.safeIndex(_1385_v0, new BigNumber((_1396_v5).length))]);
        _1397_v6 = _nw213;
        let _1398_v7;
        let _nw214 = new _module.C0();
        _nw214.__ctor((_1385_v0).isLessThanOrEqualTo(new BigNumber(74)));
        _1398_v7 = _nw214;
      } else if (_source21.is_DC35) {
        let _1399___mcc_h3 = (_source21).cf46;
        let _1400_cf46 = _1399___mcc_h3;
        let _1401_v8;
        let _nw215 = Array((new BigNumber(13)).toNumber()).fill(_module.D0.Default());
        _1401_v8 = _nw215;
        let _1402_v9;
        _1402_v9 = _module.D0.create_DC1(_module.__default.fm4(new BigNumber(140), _1385_v0, globalState));
        let _1403_v10;
        _1403_v10 = _dafny.Seq.of(_1402_v9, _1402_v9, _module.D0.create_DC0(_1387_v2.f25));
        let _1404_v11;
        _1404_v11 = _module.D0.create_DC1((_1403_v10)[_module.__default.safeIndex(_1385_v0, new BigNumber((_1403_v10).length))]);
        let _index218 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_1401_v8).length));
        (_1401_v8)[_index218] = _1404_v11;
        let _pat_let_tv30 = _1402_v9;
        let _index219 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_1401_v8).length));
        (_1401_v8)[_index219] = function (_pat_let22_0) {
          return function (_1405_dt__update__tmp_h0) {
            return function (_pat_let23_0) {
              return function (_1406_dt__update_hcf1_h0) {
                return _module.D0.create_DC1(_1406_dt__update_hcf1_h0);
              }(_pat_let23_0);
            }(_pat_let_tv30);
          }(_pat_let22_0);
        }(_1404_v11);
        let _1407_v12;
        let _init42 = ((_1408_v0) => function (_1409_i0) {
          return _dafny.Seq.of(_1408_v0);
        })(_1385_v0);
        let _nw216 = Array((new BigNumber(5)).toNumber());
        for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw216.length); _i0_42++) {
          _nw216[_i0_42] = _init42(new BigNumber(_i0_42));
        }
        _1407_v12 = _nw216;
        let _1410_v13;
        _1410_v13 = _module.D15.create_DC32(_1407_v12);
        let _1411_v14;
        _1411_v14 = _dafny.MultiSet.fromElements(_1410_v13);
        if (!(!((_1411_v14).update(_1410_v13, _module.__default.abs(_1385_v0))).equals(_1411_v14))) {
          let _1412_v16;
          _1412_v16 = _dafny.Seq.of((_1387_v2).f26, !((_1387_v2).f26));
          let _1413_v17;
          _1413_v17 = _dafny.Seq.of(p0);
          let _1414_v18;
          let _nw217 = Array((new BigNumber(20)).toNumber());
          _nw217[(_dafny.ZERO).toNumber()] = _1387_v2.f25;
          _nw217[(_dafny.ONE).toNumber()] = ((_1387_v2).f26) === ((_this).f24);
          _nw217[(new BigNumber(2)).toNumber()] = _1387_v2.f25;
          _nw217[(new BigNumber(3)).toNumber()] = (_1387_v2).f26;
          _nw217[(new BigNumber(4)).toNumber()] = (_1387_v2).f26;
          _nw217[(new BigNumber(5)).toNumber()] = (function () {
            let _coll58 = new _dafny.Set();
            for (const _compr_58 of _dafny.IntegerRange(new BigNumber(100), new BigNumber(-738))) {
              let _1415_v15 = _compr_58;
              if (((new BigNumber(100)).isLessThanOrEqualTo(_1415_v15)) && ((_1415_v15).isLessThan(new BigNumber(-738)))) {
                _coll58.add((_1415_v15).minus(_1385_v0));
              }
            }
            return _coll58;
          }()).IsSubsetOf(_1386_v1);
          _nw217[(new BigNumber(6)).toNumber()] = _1387_v2.f25;
          _nw217[(new BigNumber(7)).toNumber()] = _1387_v2.f25;
          _nw217[(new BigNumber(8)).toNumber()] = !((((_this).fm7((_1387_v2).f26, _1385_v0, _1387_v2.f25, globalState)) ? ((_1387_v2).f26) : ((_1387_v2).f26)));
          _nw217[(new BigNumber(9)).toNumber()] = (_1387_v2).f26;
          _nw217[(new BigNumber(10)).toNumber()] = (_1387_v2).f26;
          _nw217[(new BigNumber(11)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_1412_v16, _dafny.Seq.update(_1412_v16, _module.__default.safeIndex(_1385_v0, new BigNumber((_1412_v16).length)), false));
          _nw217[(new BigNumber(12)).toNumber()] = (_1387_v2).f26;
          _nw217[(new BigNumber(13)).toNumber()] = _dafny.Seq.IsPrefixOf(_1413_v17, _1413_v17);
          _nw217[(new BigNumber(14)).toNumber()] = (_1387_v2).f26;
          _nw217[(new BigNumber(15)).toNumber()] = _1387_v2.f25;
          _nw217[(new BigNumber(16)).toNumber()] = (_1387_v2).f26;
          _nw217[(new BigNumber(17)).toNumber()] = _1387_v2.f25;
          _nw217[(new BigNumber(18)).toNumber()] = (_1387_v2).f26;
          _nw217[(new BigNumber(19)).toNumber()] = (_1387_v2).f26;
          _1414_v18 = _nw217;
          _1414_v18 = _1414_v18;
          let _1416_v19;
          _1416_v19 = _dafny.Seq.UnicodeFromString("lb");
          (globalState).f20 = _1416_v19;
          (globalState).f7 = _1385_v0;
          let _1417_v20;
          _1417_v20 = _dafny.MultiSet.fromElements(_1385_v0);
          let _1418_v21;
          _1418_v21 = _dafny.Seq.of(new BigNumber((_1388_v3).cardinality()), _1385_v0);
          (globalState).f16 = (_module.__default.fm0(new BigNumber((_1417_v20).cardinality()), _1385_v0, new BigNumber((_1418_v21).length), (_1387_v2).f26, globalState)).multipliedBy(_1385_v0);
          let _1419_v22;
          let _nw218 = new _module.C4();
          _nw218.__ctor((_1387_v2).f26, !((_1387_v2).f26));
          _1419_v22 = _nw218;
        } else {
          (_1387_v2).m5(_1387_v2.f25, globalState);
          let _1420_v23;
          _1420_v23 = _dafny.Seq.of(_1387_v2.f25);
          let _1421_v24;
          _1421_v24 = _dafny.Set.fromElements(_1387_v2.f25, (_this).f24, _1387_v2.f25);
          let _1422_v25;
          let _nw219 = Array((new BigNumber(14)).toNumber());
          _nw219[(_dafny.ZERO).toNumber()] = _1387_v2.f25;
          _nw219[(_dafny.ONE).toNumber()] = _dafny.Seq.IsProperPrefixOf(_1420_v23, _1420_v23);
          _nw219[(new BigNumber(2)).toNumber()] = false;
          _nw219[(new BigNumber(3)).toNumber()] = !_dafny.Seq.contains(_1420_v23, (_1387_v2).f26);
          _nw219[(new BigNumber(4)).toNumber()] = !((_1387_v2).f26);
          _nw219[(new BigNumber(5)).toNumber()] = (_1387_v2).f26;
          _nw219[(new BigNumber(6)).toNumber()] = (_1387_v2).f26;
          _nw219[(new BigNumber(7)).toNumber()] = (_1387_v2).f26;
          _nw219[(new BigNumber(8)).toNumber()] = ((_1387_v2.f25) ? ((_1387_v2).f26) : (true));
          _nw219[(new BigNumber(9)).toNumber()] = _1387_v2.f25;
          _nw219[(new BigNumber(10)).toNumber()] = (_1421_v24).equals(_1421_v24);
          _nw219[(new BigNumber(11)).toNumber()] = !(_1385_v0).isEqualTo(_1385_v0);
          _nw219[(new BigNumber(12)).toNumber()] = false;
          _nw219[(new BigNumber(13)).toNumber()] = (_this).f24;
          _1422_v25 = _nw219;
          let _index220 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_1422_v25).length));
          (_1422_v25)[_index220] = (_this).f24;
          let _1423_v26;
          _1423_v26 = _dafny.Seq.of(_1385_v0);
          let _index221 = _module.__default.safeIndex(new BigNumber(438), new BigNumber((_1422_v25).length));
          (_1422_v25)[_index221] = (((_dafny.ZERO).minus(_1385_v0)).isEqualTo(_1385_v0)) || (_dafny.Seq.IsPrefixOf(_1423_v26, _1423_v26));
          let _1424_v27;
          _1424_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1385_v0,(_this).f24);
          let _1425_v28;
          _1425_v28 = _dafny.Map.Empty.slice().updateUnsafe((_1387_v2).f26,false);
          let _rhs203 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_dafny.Set.fromElements(_1387_v2.f25)).Intersect(_1421_v24)).length),(_1422_v25)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_1422_v25).length))]);
          let _rhs204 = ((_1425_v28).update(_1387_v2.f25, (_1387_v2).f26)).Merge((_1425_v28).Merge(_1425_v28));
          let _rhs205 = (_1385_v0).minus(new BigNumber(155));
          let _rhs206 = _1385_v0;
          let _rhs207 = (_1387_v2).f26;
          let _lhs193 = globalState;
          let _lhs194 = globalState;
          let _lhs195 = globalState;
          _1424_v27 = _rhs203;
          _1425_v28 = _rhs204;
          _lhs193.f16 = _rhs205;
          _lhs194.f7 = _rhs206;
          _lhs195.f2 = _rhs207;
          let _1426_v29;
          _1426_v29 = _dafny.Seq.UnicodeFromString("udixpjur");
          let _1427_v30;
          _1427_v30 = new _dafny.CodePoint('m'.codePointAt(0));
          (globalState).f8 = _dafny.Seq.update(_1426_v29, _module.__default.safeIndex(new BigNumber((_1421_v24).length), new BigNumber((_1426_v29).length)), _1427_v30);
          let _1428_v31;
          _1428_v31 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm32((_1422_v25)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_1422_v25).length))], _1385_v0, new BigNumber((_1423_v26).length), globalState),_1387_v2.f25);
          let _1429_v32;
          _1429_v32 = _module.D14.create_DC28(_dafny.Set.fromElements(new BigNumber(717), _1385_v0));
          _1428_v31 = (_1428_v31).update((((_1422_v25)[_module.__default.safeIndex(new BigNumber(438), new BigNumber((_1422_v25).length))]) ? (_1429_v32) : (_1429_v32)), (_module.__default.fm44(globalState)).dtor_cf36);
        }
        let _index222 = _module.__default.safeIndex(new BigNumber(774), new BigNumber((p0).length));
        (p0)[_index222] = _1385_v0;
        let _index223 = _module.__default.safeIndex(new BigNumber(774), new BigNumber((p0).length));
        (p0)[_index223] = _1385_v0;
        let _1430_v33;
        _1430_v33 = _dafny.Map.Empty.slice().updateUnsafe((_1387_v2).f26,new BigNumber(262));
        let _1431_v34;
        _1431_v34 = _dafny.Seq.of(_1430_v33);
        let _1432_v35;
        _1432_v35 = _dafny.Seq.of(true);
        let _1433_v36;
        _1433_v36 = _dafny.Set.fromElements(_1387_v2.f25, (_1387_v2).f26);
        let _1434_v37;
        let _nw220 = Array((new BigNumber(9)).toNumber());
        _nw220[(_dafny.ZERO).toNumber()] = (_1387_v2).f26;
        _nw220[(_dafny.ONE).toNumber()] = _1387_v2.f25;
        _nw220[(new BigNumber(2)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(_1430_v33), _1431_v34);
        _nw220[(new BigNumber(3)).toNumber()] = _1387_v2.f25;
        _nw220[(new BigNumber(4)).toNumber()] = _1387_v2.f25;
        _nw220[(new BigNumber(5)).toNumber()] = (_1387_v2.f25) || (_1387_v2.f25);
        _nw220[(new BigNumber(6)).toNumber()] = (_this).f24;
        _nw220[(new BigNumber(7)).toNumber()] = !((_1387_v2).f26) || ((_1387_v2).f26);
        _nw220[(new BigNumber(8)).toNumber()] = !(new BigNumber((_1432_v35).length)).isEqualTo(new BigNumber((_1433_v36).length));
        _1434_v37 = _nw220;
        let _index224 = _module.__default.safeIndex(new BigNumber(566), new BigNumber((_1434_v37).length));
        (_1434_v37)[_index224] = _1387_v2.f25;
        let _index225 = _module.__default.safeIndex(new BigNumber(566), new BigNumber((_1434_v37).length));
        (_1434_v37)[_index225] = (_1387_v2).f26;
      } else {
        let _1435___mcc_h4 = (_source21).cf50;
        let _1436_cf50 = _1435___mcc_h4;
        let _1437_v38;
        _1437_v38 = _dafny.Seq.UnicodeFromString("vuuysgyyb");
        let _1438_v39;
        _1438_v39 = _dafny.Seq.of(_1387_v2.f25);
        let _1439_v40;
        _1439_v40 = _dafny.MultiSet.fromElements((new BigNumber((_1437_v38).length)).plus((((_1388_v3).contains(_1387_v2.f25)) ? ((_1388_v3).get(_1387_v2.f25)) : (_1385_v0))), _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(_1438_v39)).length), new BigNumber((_1437_v38).length)));
        _1439_v40 = _dafny.MultiSet.fromElements(_1385_v0, (_1385_v0).multipliedBy(new BigNumber(306)), _module.__default.fm0(_1385_v0, (_dafny.ZERO).minus(_1385_v0), _1385_v0, !(_1387_v2.f25), globalState), new BigNumber(241));
        (globalState).f13 = _1387_v2.f25;
        (globalState).f13 = (_1387_v2).f26;
        let _nw221 = new _module.C4();
        _nw221.__ctor((_1387_v2).f26, _1387_v2.f25);
        _1387_v2 = _nw221;
      }
      let _1440_v41;
      _1440_v41 = _dafny.Map.Empty.slice().updateUnsafe((_1387_v2).f26,(_this).f24);
      if (_module.__default.fm2(_1440_v41, globalState)) {
        let _1441_v42;
        let _init43 = function (_1442_i1) {
          return new _dafny.CodePoint('a'.codePointAt(0));
        };
        let _nw222 = Array((new BigNumber(6)).toNumber());
        for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw222.length); _i0_43++) {
          _nw222[_i0_43] = _init43(new BigNumber(_i0_43));
        }
        _1441_v42 = _nw222;
        _1441_v42 = _1441_v42;
        let _1443_v43;
        _1443_v43 = _dafny.MultiSet.fromElements(_1385_v0, _1385_v0);
        if (((((_1387_v2).f26) ? (_1385_v0) : ((((_1443_v43).contains(_1385_v0)) ? ((_1443_v43).get(_1385_v0)) : (_1385_v0))))).isLessThanOrEqualTo(new BigNumber(161))) {
          let _1444_v44;
          _1444_v44 = _dafny.Seq.of(_1388_v3);
          (globalState).f16 = _module.__default.safeDivisionInt(new BigNumber((_1444_v44).length), _1385_v0);
          (_1387_v2).f25 = _1387_v2.f25;
          let _1445_v45;
          _1445_v45 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_1385_v0, new BigNumber(380)),_1385_v0);
          let _1446_v46;
          _1446_v46 = _dafny.MultiSet.fromElements(p0);
          _1445_v45 = (_1445_v45).update(_1385_v0, (_1385_v0).minus(new BigNumber((_1446_v46).cardinality())));
          let _1447_v47;
          let _init44 = function (_1448_i2) {
            return _dafny.Seq.UnicodeFromString("hkapgn");
          };
          let _nw223 = Array((new BigNumber(17)).toNumber());
          for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw223.length); _i0_44++) {
            _nw223[_i0_44] = _init44(new BigNumber(_i0_44));
          }
          _1447_v47 = _nw223;
          let _1449_v48;
          _1449_v48 = new _dafny.CodePoint('r'.codePointAt(0));
          let _1450_v49;
          _1450_v49 = _dafny.Map.Empty.slice().updateUnsafe(_1447_v47,_1449_v48);
          _1450_v49 = (_1450_v49).update(_1447_v47, _1449_v48);
          let _index226 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((p0).length));
          (p0)[_index226] = new BigNumber(-806);
          let _index227 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((p0).length));
          (p0)[_index227] = _1385_v0;
        } else {
          let _1451_v50;
          _1451_v50 = _dafny.Map.Empty.slice().updateUnsafe(_1385_v0,(_1387_v2).f26);
          let _1452_v52;
          _1452_v52 = _dafny.Seq.of(_1451_v50, _1451_v50, function () {
            let _coll59 = new _dafny.Map();
            for (const _compr_59 of (_1386_v1).Elements) {
              let _1453_v51 = _compr_59;
              if ((_1386_v1).contains(_1453_v51)) {
                _coll59.push([(_1453_v51).plus(new BigNumber((_dafny.Seq.of(_1387_v2.f25)).length)),_1387_v2.f25]);
              }
            }
            return _coll59;
          }());
          let _1454_v53;
          let _nw224 = new _module.C1();
          _nw224.__ctor((_1385_v0).multipliedBy(_1385_v0), ((_1387_v2.f25) ? (_dafny.Seq.of(_1451_v50, _1451_v50, _1451_v50)) : (_1452_v52)));
          _1454_v53 = _nw224;
          let _1455_v54;
          _1455_v54 = _dafny.Seq.of(p0);
          _1455_v54 = _1455_v54;
          _1388_v3 = _1388_v3;
          (globalState).f0 = (_1387_v2).f26;
          let _1456_v55;
          let _nw225 = new _module.C2();
          _nw225.__ctor(_1385_v0);
          _1456_v55 = _nw225;
        }
        let _source22 = _module.__default.fm45(globalState);
        if (_source22.is_DC26) {
          let _1457___mcc_h5 = (_source22).cf33;
          let _1458___mcc_h6 = (_source22).cf34;
          let _1459___mcc_h7 = (_source22).cf35;
          let _1460_cf35 = _1459___mcc_h7;
          let _1461_cf34 = _1458___mcc_h6;
          let _1462_cf33 = _1457___mcc_h5;
          (globalState).f16 = _1385_v0;
          (globalState).f13 = (_1387_v2).f26;
          let _1463_v56;
          _1463_v56 = new _dafny.CodePoint('w'.codePointAt(0));
          _1463_v56 = (_1462_cf33)[_module.__default.safeIndex((_dafny.ZERO).minus(_1385_v0), new BigNumber((_1462_cf33).length))];
          let _1464_v57;
          _1464_v57 = _dafny.Seq.of((new BigNumber(305)).multipliedBy(_1385_v0));
          _1464_v57 = _dafny.Seq.Concat(_1464_v57, _dafny.Seq.Concat(_1464_v57, _1464_v57));
        } else if (_source22.is_DC27) {
          let _1465___mcc_h8 = (_source22).cf36;
          let _1466_cf36 = _1465___mcc_h8;
          (globalState).f0 = !((_1388_v3).contains((_1387_v2).f26));
          let _1467_v58;
          let _nw226 = Array((new BigNumber(20)).toNumber()).fill(false);
          _1467_v58 = _nw226;
          let _1468_v59;
          _1468_v59 = _dafny.MultiSet.fromElements(_1467_v58, _1467_v58, _1467_v58);
          let _1469_v60;
          _1469_v60 = _dafny.Seq.of(_dafny.MultiSet.fromElements(_1467_v58));
          _1468_v59 = (_1469_v60)[_module.__default.safeIndex(((_1387_v2.f25) ? (_1385_v0) : (_1385_v0)), new BigNumber((_1469_v60).length))];
          let _1470_v61;
          _1470_v61 = _dafny.Map.Empty.slice().updateUnsafe(_1385_v0,(_1387_v2).f26);
          let _1471_v62;
          _1471_v62 = _dafny.Seq.of(_module.__default.fm2(_1440_v41, globalState), (((_1470_v61).contains(_1385_v0)) ? ((_1470_v61).get(_1385_v0)) : ((_this).f24)), (_1385_v0).isLessThan(_1385_v0));
          let _rhs208 = (_1471_v62)[_module.__default.safeIndex(_1385_v0, new BigNumber((_1471_v62).length))];
          let _rhs209 = (_dafny.ZERO).minus(_1385_v0);
          let _rhs210 = _1467_v58;
          let _rhs211 = _1387_v2;
          let _rhs212 = _1385_v0;
          let _lhs196 = _1387_v2;
          let _lhs197 = globalState;
          let _lhs198 = globalState;
          _lhs196.f25 = _rhs208;
          _lhs197.f7 = _rhs209;
          _1467_v58 = _rhs210;
          _1387_v2 = _rhs211;
          _lhs198.f7 = _rhs212;
          let _1472_v63;
          _1472_v63 = new _dafny.CodePoint('c'.codePointAt(0));
          _1472_v63 = _1472_v63;
        } else {
          let _1473___mcc_h9 = (_source22).cf32;
          let _1474_cf32 = _1473___mcc_h9;
          let _1475_v64;
          let _nw227 = Array((new BigNumber(16)).toNumber()).fill([]);
          _1475_v64 = _nw227;
          _1475_v64 = _1475_v64;
          let _1476_v65;
          _1476_v65 = _dafny.Map.Empty.slice().updateUnsafe(_1387_v2.f25,_1474_cf32);
          _1476_v65 = (_1476_v65).update((_1387_v2).f26, _1474_cf32);
          let _1477_v66;
          let _nw228 = Array((new BigNumber(11)).toNumber()).fill(_dafny.MultiSet.Empty);
          _1477_v66 = _nw228;
          let _rhs213 = _1387_v2.f25;
          let _rhs214 = _1477_v66;
          let _rhs215 = (_1387_v2).f26;
          let _rhs216 = _1387_v2.f25;
          let _lhs199 = globalState;
          let _lhs200 = _1387_v2;
          let _lhs201 = globalState;
          _lhs199.f2 = _rhs213;
          _1477_v66 = _rhs214;
          _lhs200.f25 = _rhs215;
          _lhs201.f0 = _rhs216;
          let _1478_v67;
          let _nw229 = new _module.C2();
          _nw229.__ctor((_dafny.ZERO).minus(_1385_v0));
          _1478_v67 = _nw229;
        }
        let _1479_v68;
        _1479_v68 = new _dafny.CodePoint('q'.codePointAt(0));
        let _index228 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_1441_v42).length));
        (_1441_v42)[_index228] = _1479_v68;
        let _index229 = _module.__default.safeIndex(new BigNumber(175), new BigNumber((_1441_v42).length));
        (_1441_v42)[_index229] = _1479_v68;
        let _1480_v69;
        _1480_v69 = _dafny.Seq.of((_1387_v2).f26);
        let _1481_v70;
        _1481_v70 = _module.D14.create_DC30(_dafny.Seq.Create(_module.__default.abs(new BigNumber(116)), function (_1482_i3) {
  return new BigNumber(832);
}), _1480_v69);
        let _1483_v71;
        let _nw230 = Array((new BigNumber(20)).toNumber());
        _nw230[(_dafny.ZERO).toNumber()] = _1480_v69;
        _nw230[(_dafny.ONE).toNumber()] = _1480_v69;
        _nw230[(new BigNumber(2)).toNumber()] = _module.__default.fm18(_1387_v2.f25, new _dafny.CodePoint('p'.codePointAt(0)), (_1387_v2).f26, globalState);
        _nw230[(new BigNumber(3)).toNumber()] = _dafny.Seq.of(_1387_v2.f25);
        _nw230[(new BigNumber(4)).toNumber()] = _dafny.Seq.of(true, !(false), (_this).f24, (_1387_v2).f26, true);
        _nw230[(new BigNumber(5)).toNumber()] = _1480_v69;
        _nw230[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_1387_v2.f25), _1480_v69);
        _nw230[(new BigNumber(7)).toNumber()] = _dafny.Seq.of((_this).f24);
        _nw230[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_1480_v69, _1480_v69);
        _nw230[(new BigNumber(9)).toNumber()] = _1480_v69;
        _nw230[(new BigNumber(10)).toNumber()] = _1480_v69;
        _nw230[(new BigNumber(11)).toNumber()] = _1480_v69;
        _nw230[(new BigNumber(12)).toNumber()] = _1480_v69;
        _nw230[(new BigNumber(13)).toNumber()] = _1480_v69;
        _nw230[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_1480_v69, _dafny.Seq.of(_1387_v2.f25, _1387_v2.f25));
        _nw230[(new BigNumber(15)).toNumber()] = ((true) ? (_module.__default.fm18(_module.__default.fm2(_1440_v41, globalState), _1479_v68, _1387_v2.f25, globalState)) : ((_1481_v70).dtor_cf41));
        _nw230[(new BigNumber(16)).toNumber()] = _dafny.Seq.of(_1387_v2.f25);
        _nw230[(new BigNumber(17)).toNumber()] = _1480_v69;
        _nw230[(new BigNumber(18)).toNumber()] = _1480_v69;
        _nw230[(new BigNumber(19)).toNumber()] = _1480_v69;
        _1483_v71 = _nw230;
        let _index230 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_1483_v71).length));
        (_1483_v71)[_index230] = _1480_v69;
        let _1484_v72;
        _1484_v72 = _dafny.Set.fromElements((_1387_v2).f26);
        let _1485_v73;
        _1485_v73 = (_1484_v72).Difference(_dafny.Set.fromElements((_1387_v2).f26, _1387_v2.f25));
        let _1486_v74;
        _1486_v74 = _dafny.Seq.UnicodeFromString("sgpusvt");
        let _index231 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_1483_v71).length));
        let _rhs217 = _1486_v74;
        let _rhs218 = new BigNumber(-46);
        let _rhs219 = _1480_v69;
        let _rhs220 = _1485_v73;
        let _rhs221 = _1387_v2.f25;
        let _lhs202 = globalState;
        let _lhs203 = globalState;
        let _lhs204 = _1483_v71;
        let _lhs205 = _module.__default.safeIndex(_dafny.ONE, new BigNumber((_1483_v71).length));
        let _lhs206 = globalState;
        _lhs202.f20 = _rhs217;
        _lhs203.f7 = _rhs218;
        _lhs204[_lhs205] = _rhs219;
        _1485_v73 = _rhs220;
        _lhs206.f21 = _rhs221;
      } else {
        let _1487_v75;
        _1487_v75 = _dafny.Seq.UnicodeFromString("bbe");
        let _1488_v76;
        _1488_v76 = _dafny.Map.Empty.slice().updateUnsafe(!_dafny.areEqual(_1487_v75, _1487_v75),_1385_v0);
        let _rhs222 = _module.__default.safeModuloInt(_1385_v0, _1385_v0);
        let _rhs223 = (_1387_v2).f26;
        let _rhs224 = (_1488_v76).Merge(_1488_v76);
        let _lhs207 = globalState;
        let _lhs208 = _1387_v2;
        _lhs207.f7 = _rhs222;
        _lhs208.f25 = _rhs223;
        _1488_v76 = _rhs224;
        let _1489_v77;
        let _init45 = ((_1490_v2) => function (_1491_i4) {
          return ((true) ? (_module.D13.create_DC27(_1490_v2.f25)) : (_module.D13.create_DC27(_1490_v2.f25)));
        })(_1387_v2);
        let _nw231 = Array((new BigNumber(3)).toNumber());
        for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw231.length); _i0_45++) {
          _nw231[_i0_45] = _init45(new BigNumber(_i0_45));
        }
        _1489_v77 = _nw231;
        let _1492_v78;
        _1492_v78 = _module.D13.create_DC27(_1387_v2.f25);
        let _index232 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_1489_v77).length));
        (_1489_v77)[_index232] = _1492_v78;
        let _index233 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_1489_v77).length));
        (_1489_v77)[_index233] = _1492_v78;
        let _1493_v79;
        _1493_v79 = _dafny.Seq.of(_1387_v2.f25);
        let _1494_v80;
        _1494_v80 = _dafny.Seq.of(new BigNumber((_1493_v79).length), _1385_v0);
        let _1495_v81;
        _1495_v81 = _dafny.MultiSet.fromElements(_1385_v0);
        let _1496_v82;
        _1496_v82 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1494_v80).length),_1495_v81);
        let _1497_v83;
        _1497_v83 = _dafny.Map.Empty.slice().updateUnsafe(_1385_v0,(((_1496_v82).contains(_1385_v0)) ? ((_1496_v82).get(_1385_v0)) : (_1495_v81)));
        let _1498_v84;
        _1498_v84 = _dafny.Set.fromElements(_1493_v79);
        let _1499_v85;
        let _nw232 = Array((new BigNumber(15)).toNumber());
        _nw232[(_dafny.ZERO).toNumber()] = (_1495_v81).IsSubsetOf((((_1497_v83).contains(_1385_v0)) ? ((_1497_v83).get(_1385_v0)) : (_1495_v81)));
        _nw232[(_dafny.ONE).toNumber()] = _dafny.Seq.IsPrefixOf(_1493_v79, _dafny.Seq.of((_this).f24));
        _nw232[(new BigNumber(2)).toNumber()] = (_1387_v2).f26;
        _nw232[(new BigNumber(3)).toNumber()] = (_1385_v0).isLessThanOrEqualTo(_1385_v0);
        _nw232[(new BigNumber(4)).toNumber()] = _dafny.areEqual(_dafny.Seq.of(_1387_v2.f25, _1387_v2.f25, _1387_v2.f25), _dafny.Seq.of(_1387_v2.f25, (_1387_v2).f26, _1387_v2.f25));
        _nw232[(new BigNumber(5)).toNumber()] = (!((_1387_v2).f26)) || (false);
        _nw232[(new BigNumber(6)).toNumber()] = (_this).f24;
        _nw232[(new BigNumber(7)).toNumber()] = (_1387_v2).f26;
        _nw232[(new BigNumber(8)).toNumber()] = ((_1387_v2.f25) ? ((_this).f24) : ((_this).f24));
        _nw232[(new BigNumber(9)).toNumber()] = (_1493_v79)[_module.__default.safeIndex(_1385_v0, new BigNumber((_1493_v79).length))];
        _nw232[(new BigNumber(10)).toNumber()] = (_1387_v2).f26;
        _nw232[(new BigNumber(11)).toNumber()] = (_1387_v2).f26;
        _nw232[(new BigNumber(12)).toNumber()] = (_1498_v84).equals(_1498_v84);
        _nw232[(new BigNumber(13)).toNumber()] = (_1387_v2).f26;
        _nw232[(new BigNumber(14)).toNumber()] = (_1387_v2).f26;
        _1499_v85 = _nw232;
        let _index234 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((p0).length));
        (p0)[_index234] = _module.__default.safeModuloInt((_dafny.ZERO).minus(_1385_v0), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-100)), ((_1500_v75, _1501_v2, _1502_v76, _1503_v0) => function (_1504_i5) {
          return (_1500_v75)[_module.__default.safeIndex((((_1502_v76).contains(_1501_v2.f25)) ? ((_1502_v76).get(_1501_v2.f25)) : (_1503_v0)), new BigNumber((_1500_v75).length))];
        })(_1487_v75, _1387_v2, _1488_v76, _1385_v0))).length));
        let _1505_v86;
        _1505_v86 = _dafny.Seq.of(_1487_v75, _dafny.Seq.Create(_module.__default.abs(new BigNumber(441)), function (_1506_i6) {
          return new _dafny.CodePoint('s'.codePointAt(0));
        }));
        let _1507_v87;
        _1507_v87 = _dafny.MultiSet.fromElements(_1499_v85);
        let _index235 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((p0).length));
        let _rhs225 = !(((_1507_v87).Union((_module.D19.create_DC45(_1507_v87)).dtor_cf59)).IsSubsetOf(_1507_v87));
        let _rhs226 = _1499_v85;
        let _rhs227 = _module.__default.safeModuloInt(new BigNumber((_1487_v75).length), _1385_v0);
        let _rhs228 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1505_v86, _1505_v86), _dafny.Seq.of(_1487_v75, _1487_v75, _1487_v75));
        let _lhs209 = globalState;
        let _lhs210 = p0;
        let _lhs211 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((p0).length));
        _lhs209.f13 = _rhs225;
        _1499_v85 = _rhs226;
        _lhs210[_lhs211] = _rhs227;
        _1505_v86 = _rhs228;
        let _index236 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_1499_v85).length));
        (_1499_v85)[_index236] = _1387_v2.f25;
        let _index237 = _module.__default.safeIndex(new BigNumber(383), new BigNumber((_1499_v85).length));
        (_1499_v85)[_index237] = (_1385_v0).isLessThan((_1385_v0).multipliedBy(_1385_v0));
        let _1508_v88;
        _1508_v88 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1487_v75).length),new _dafny.CodePoint('r'.codePointAt(0)));
        let _1509_v89;
        _1509_v89 = _dafny.Set.fromElements(_1387_v2.f25);
        if (((_1385_v0).multipliedBy((((_1495_v81).contains((_this).fm8(true, _1508_v88, _dafny.Seq.update(_1487_v75, _module.__default.safeIndex((p0)[_module.__default.safeIndex(new BigNumber(568), new BigNumber((p0).length))], new BigNumber((_1487_v75).length)), _module.__default.fm34(globalState)), globalState))) ? ((_1495_v81).get((_this).fm8(true, _1508_v88, _dafny.Seq.update(_1487_v75, _module.__default.safeIndex((p0)[_module.__default.safeIndex(new BigNumber(568), new BigNumber((p0).length))], new BigNumber((_1487_v75).length)), _module.__default.fm34(globalState)), globalState))) : (new BigNumber((_1509_v89).length))))).isEqualTo((p0)[_module.__default.safeIndex(new BigNumber(568), new BigNumber((p0).length))])) {
          let _1510_v90;
          _1510_v90 = _module.D4.create_DC10(_1385_v0, (_dafny.ZERO).minus(_1385_v0), _1387_v2.f25, false, new BigNumber((_1488_v76).length));
          (globalState).f2 = (_1510_v90).dtor_cf14;
          let _1511_v91;
          _1511_v91 = _dafny.Map.Empty.slice().updateUnsafe((_1387_v2).f26,_1487_v75);
          let _1512_v92;
          _1512_v92 = new _dafny.CodePoint('y'.codePointAt(0));
          let _1513_v93;
          let _nw233 = Array((new BigNumber(29)).toNumber());
          _nw233[(_dafny.ZERO).toNumber()] = _1487_v75;
          _nw233[(_dafny.ONE).toNumber()] = (((_1511_v91).contains((_1387_v2).f26)) ? ((_1511_v91).get((_1387_v2).f26)) : (_1487_v75));
          _nw233[(new BigNumber(2)).toNumber()] = _module.__default.fm1((p0)[_module.__default.safeIndex(new BigNumber(568), new BigNumber((p0).length))], globalState);
          _nw233[(new BigNumber(3)).toNumber()] = _dafny.Seq.UnicodeFromString("hyp");
          _nw233[(new BigNumber(4)).toNumber()] = _1487_v75;
          _nw233[(new BigNumber(5)).toNumber()] = _1487_v75;
          _nw233[(new BigNumber(6)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-529)), function (_1514_i7) {
            return new _dafny.CodePoint('o'.codePointAt(0));
          });
          _nw233[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("rdnqhsuh"), _dafny.Seq.UnicodeFromString("fuy"));
          _nw233[(new BigNumber(8)).toNumber()] = _1487_v75;
          _nw233[(new BigNumber(9)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-316)), ((_1515_v75, _1516_p0) => function (_1517_i8) {
            return (_1515_v75)[_module.__default.safeIndex((_1516_p0)[_module.__default.safeIndex(new BigNumber(568), new BigNumber((_1516_p0).length))], new BigNumber((_1515_v75).length))];
          })(_1487_v75, p0));
          _nw233[(new BigNumber(10)).toNumber()] = _1487_v75;
          _nw233[(new BigNumber(11)).toNumber()] = _1487_v75;
          _nw233[(new BigNumber(12)).toNumber()] = _dafny.Seq.UnicodeFromString("cepxkgleo");
          _nw233[(new BigNumber(13)).toNumber()] = _dafny.Seq.update(_1487_v75, _module.__default.safeIndex((p0)[_module.__default.safeIndex(new BigNumber(568), new BigNumber((p0).length))], new BigNumber((_1487_v75).length)), new _dafny.CodePoint('d'.codePointAt(0)));
          _nw233[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_1487_v75, (_1505_v86)[_module.__default.safeIndex(_1385_v0, new BigNumber((_1505_v86).length))]);
          _nw233[(new BigNumber(15)).toNumber()] = _1487_v75;
          _nw233[(new BigNumber(16)).toNumber()] = _1487_v75;
          _nw233[(new BigNumber(17)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(116)), function (_1518_i9) {
            return new _dafny.CodePoint('y'.codePointAt(0));
          });
          _nw233[(new BigNumber(18)).toNumber()] = _1487_v75;
          _nw233[(new BigNumber(19)).toNumber()] = _dafny.Seq.update(_1487_v75, _module.__default.safeIndex(_1385_v0, new BigNumber((_1487_v75).length)), _1512_v92);
          _nw233[(new BigNumber(20)).toNumber()] = _dafny.Seq.Concat(_1487_v75, _dafny.Seq.Create(_module.__default.abs(new BigNumber(807)), ((_1519_v92) => function (_1520_i10) {
            return _1519_v92;
          })(_1512_v92)));
          _nw233[(new BigNumber(21)).toNumber()] = _1487_v75;
          _nw233[(new BigNumber(22)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(885)), function (_1521_i11) {
            return new _dafny.CodePoint('c'.codePointAt(0));
          });
          _nw233[(new BigNumber(23)).toNumber()] = _dafny.Seq.Concat((_module.D13.create_DC26(_1487_v75, _1387_v2.f25, _1387_v2.f25)).dtor_cf33, _1487_v75);
          _nw233[(new BigNumber(24)).toNumber()] = _dafny.Seq.UnicodeFromString("ctaxiqtj");
          _nw233[(new BigNumber(25)).toNumber()] = _1487_v75;
          _nw233[(new BigNumber(26)).toNumber()] = _1487_v75;
          _nw233[(new BigNumber(27)).toNumber()] = _1487_v75;
          _nw233[(new BigNumber(28)).toNumber()] = _1487_v75;
          _1513_v93 = _nw233;
          let _index238 = _module.__default.safeIndex(new BigNumber(501), new BigNumber((_1513_v93).length));
          (_1513_v93)[_index238] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("rckcmrpbm"), _1487_v75);
          let _1522_v94;
          _1522_v94 = _dafny.Map.Empty.slice().updateUnsafe(_1385_v0,_1487_v75);
          let _index239 = _module.__default.safeIndex(new BigNumber(501), new BigNumber((_1513_v93).length));
          (_1513_v93)[_index239] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("srrgrldpd"), (((_1522_v94).contains(_1385_v0)) ? ((_1522_v94).get(_1385_v0)) : (_1487_v75)));
          (globalState).f16 = _module.__default.safeModuloInt((p0)[_module.__default.safeIndex(new BigNumber(568), new BigNumber((p0).length))], (p0)[_module.__default.safeIndex(new BigNumber(568), new BigNumber((p0).length))]);
          let _1523_v95;
          let _nw234 = Array((new BigNumber(18)).toNumber()).fill([]);
          _1523_v95 = _nw234;
          let _1524_v96;
          _1524_v96 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1523_v95);
          _1524_v96 = (_1524_v96).update((_1387_v2).f26, _1523_v95);
          let _index240 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((p0).length));
          (p0)[_index240] = (_module.__default.safeModuloInt(_1385_v0, (p0)[_module.__default.safeIndex(new BigNumber(568), new BigNumber((p0).length))])).plus(_1385_v0);
        } else {
          let _1525_v97;
          _1525_v97 = _module.D14.create_DC30(_1494_v80, _1493_v79);
          let _1526_v98;
          _1526_v98 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_module.D14.create_DC30(_1494_v80, _1493_v79), _1525_v97, _1525_v97));
          let _1527_v99;
          _1527_v99 = new _dafny.CodePoint('x'.codePointAt(0));
          _1526_v98 = ((false) ? ((_1526_v98).Difference(_1526_v98)) : (((_1526_v98).update(_module.__default.fm46(_1387_v2.f25, _1527_v99, globalState), _module.__default.abs(new BigNumber(843)))).Difference(_1526_v98)));
          _1488_v76 = (_1488_v76).update(_dafny.Seq.IsProperPrefixOf(_1493_v79, _1493_v79), new BigNumber(931));
          let _1528_v100;
          let _nw235 = new _module.C2();
          _nw235.__ctor((((_1499_v85)[_module.__default.safeIndex(new BigNumber(383), new BigNumber((_1499_v85).length))]) ? (new BigNumber(614)) : ((p0)[_module.__default.safeIndex(new BigNumber(568), new BigNumber((p0).length))])));
          _1528_v100 = _nw235;
          _1528_v100 = _1528_v100;
          (globalState).f13 = _dafny.Seq.IsPrefixOf(_dafny.Seq.update(_1487_v75, _module.__default.safeIndex(new BigNumber((_1493_v79).length), new BigNumber((_1487_v75).length)), _1527_v99), _dafny.Seq.UnicodeFromString("vjgipyft"));
          (_1387_v2).f25 = (_1386_v1).IsSubsetOf(_1386_v1);
        }
      }
      let _1529_v101;
      _1529_v101 = _dafny.Seq.UnicodeFromString("ckhhh");
      let _1530_v102;
      _1530_v102 = _dafny.Map.Empty.slice().updateUnsafe(_1529_v101,(_this).f24);
      let _hi7 = _1385_v0;
      for (let _1531_i12 = new BigNumber((_1530_v102).length); _1531_i12.isLessThan(_hi7); _1531_i12 = _1531_i12.plus(_dafny.ONE)) {
        _1386_v1 = (_1386_v1).Difference(_dafny.Set.fromElements((_dafny.ZERO).minus(_1385_v0)));
        (_1387_v2).f25 = (_1387_v2).f26;
        if (_1387_v2.f25) {
          let _1532_v103;
          _1532_v103 = _dafny.Map.Empty.slice().updateUnsafe(_1385_v0,(_1387_v2).f26);
          let _1533_v104;
          _1533_v104 = _module.D1.create_DC2(_1532_v103);
          _1533_v104 = _1533_v104;
          let _1534_v105;
          let _nw236 = Array((new BigNumber(24)).toNumber()).fill([]);
          _1534_v105 = _nw236;
          let _1535_v106;
          let _nw237 = Array((new BigNumber(27)).toNumber());
          _nw237[(_dafny.ZERO).toNumber()] = _1387_v2.f25;
          _nw237[(_dafny.ONE).toNumber()] = (_1387_v2).f26;
          _nw237[(new BigNumber(2)).toNumber()] = _1387_v2.f25;
          _nw237[(new BigNumber(3)).toNumber()] = _1387_v2.f25;
          _nw237[(new BigNumber(4)).toNumber()] = (_1387_v2).f26;
          _nw237[(new BigNumber(5)).toNumber()] = (_this).f24;
          _nw237[(new BigNumber(6)).toNumber()] = (_1387_v2).f26;
          _nw237[(new BigNumber(7)).toNumber()] = !(_1387_v2.f25);
          _nw237[(new BigNumber(8)).toNumber()] = !(_1387_v2.f25);
          _nw237[(new BigNumber(9)).toNumber()] = (_1387_v2).f26;
          _nw237[(new BigNumber(10)).toNumber()] = (_1387_v2).f26;
          _nw237[(new BigNumber(11)).toNumber()] = (_this).fm7(_1387_v2.f25, _1531_i12, false, globalState);
          _nw237[(new BigNumber(12)).toNumber()] = _1387_v2.f25;
          _nw237[(new BigNumber(13)).toNumber()] = (_1387_v2).f26;
          _nw237[(new BigNumber(14)).toNumber()] = (_1387_v2).f26;
          _nw237[(new BigNumber(15)).toNumber()] = (_1387_v2).f26;
          _nw237[(new BigNumber(16)).toNumber()] = (_this).f24;
          _nw237[(new BigNumber(17)).toNumber()] = (_this).f24;
          _nw237[(new BigNumber(18)).toNumber()] = (_1387_v2).f26;
          _nw237[(new BigNumber(19)).toNumber()] = (_1387_v2).f26;
          _nw237[(new BigNumber(20)).toNumber()] = (_1387_v2).f26;
          _nw237[(new BigNumber(21)).toNumber()] = !(_1387_v2.f25);
          _nw237[(new BigNumber(22)).toNumber()] = (_this).fm6(new BigNumber((_1529_v101).length), globalState);
          _nw237[(new BigNumber(23)).toNumber()] = _1387_v2.f25;
          _nw237[(new BigNumber(24)).toNumber()] = !((_1387_v2).f26);
          _nw237[(new BigNumber(25)).toNumber()] = (_1387_v2).f26;
          _nw237[(new BigNumber(26)).toNumber()] = (_1387_v2).f26;
          _1535_v106 = _nw237;
          let _index241 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_1534_v105).length));
          (_1534_v105)[_index241] = _1535_v106;
          let _1536_v107;
          _1536_v107 = new _dafny.CodePoint('o'.codePointAt(0));
          let _1537_v108;
          _1537_v108 = _module.D13.create_DC26(_1529_v101, (_1387_v2).fm9(_1387_v2.f25, globalState), _1387_v2.f25);
          let _1538_v109;
          _1538_v109 = _module.D19.create_DC46(_1387_v2.f25, _1385_v0);
          let _index242 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_1534_v105).length));
          let _rhs229 = !_dafny.Seq.contains(_dafny.Seq.Concat((_1537_v108).dtor_cf33, _dafny.Seq.update(_1529_v101, _module.__default.safeIndex(_1531_i12, new BigNumber((_1529_v101).length)), _1536_v107)), _1536_v107);
          let _rhs230 = (_1387_v2).f26;
          let _rhs231 = ((_1387_v2.f25) ? ((((_1387_v2).f26) ? (_1531_i12) : ((_1538_v109).dtor_cf61))) : (_1385_v0));
          let _rhs232 = _1535_v106;
          let _lhs212 = globalState;
          let _lhs213 = globalState;
          let _lhs214 = globalState;
          let _lhs215 = _1534_v105;
          let _lhs216 = _module.__default.safeIndex(new BigNumber(805), new BigNumber((_1534_v105).length));
          _lhs212.f2 = _rhs229;
          _lhs213.f0 = _rhs230;
          _lhs214.f7 = _rhs231;
          _lhs215[_lhs216] = _rhs232;
          let _1539_v110;
          _1539_v110 = _dafny.Map.Empty.slice().updateUnsafe((_1387_v2).f26,_1385_v0);
          let _1540_v111;
          _1540_v111 = _dafny.Seq.of(_module.__default.fm1(_1531_i12, globalState), _1529_v101);
          let _1541_v112;
          _1541_v112 = _dafny.Map.Empty.slice().updateUnsafe((_1385_v0).multipliedBy(new BigNumber((_1539_v110).length)),(_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.update(_1540_v111, _module.__default.safeIndex(_1385_v0, new BigNumber((_1540_v111).length)), _dafny.Seq.UnicodeFromString("je")), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("xydhcd"))))).cardinality()))));
          let _1542_v113;
          _1542_v113 = _module.D19.create_DC47(_1533_v104, new BigNumber(-809), true);
          let _1543_v114;
          _1543_v114 = _dafny.Map.Empty.slice().updateUnsafe((_1534_v105)[_module.__default.safeIndex(new BigNumber(805), new BigNumber((_1534_v105).length))],_1387_v2.f25);
          (globalState).f16 = (((_1541_v112).contains((_1542_v113).dtor_cf63)) ? ((_1541_v112).get((_1542_v113).dtor_cf63)) : ((_dafny.ZERO).minus((((_1387_v2).f26) ? ((_dafny.ZERO).minus(new BigNumber((_1440_v41).length))) : (new BigNumber((_1543_v114).length))))));
          let _1544_v115;
          let _nw238 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
          _1544_v115 = _nw238;
          let _init46 = ((_1545_i12) => function (_1546_i13) {
            return (_1546_i13).plus(_1545_i12);
          })(_1531_i12);
          let _nw239 = Array((new BigNumber(27)).toNumber());
          for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw239.length); _i0_46++) {
            _nw239[_i0_46] = _init46(new BigNumber(_i0_46));
          }
          _1544_v115 = _nw239;
          let _1547_v116;
          let _nw240 = new _module.C0();
          _nw240.__ctor(_module.__default.fm2(_1440_v41, globalState));
          _1547_v116 = _nw240;
          let _1548_v117;
          _1548_v117 = _dafny.Map.Empty.slice().updateUnsafe(_1387_v2.f25,_1547_v116);
          _1548_v117 = (_1548_v117).update(_1547_v116.f27, _1547_v116);
        } else {
          let _1549_v118;
          let _nw241 = Array((new BigNumber(23)).toNumber()).fill(_module.D6.Default());
          _1549_v118 = _nw241;
          let _index243 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_1549_v118).length));
          (_1549_v118)[_index243] = _module.D6.create_DC12(_dafny.Seq.of(p0));
          let _1550_v119;
          _1550_v119 = new _dafny.CodePoint('w'.codePointAt(0));
          let _1551_v120;
          _1551_v120 = _dafny.Seq.of(p0, p0);
          let _1552_v121;
          _1552_v121 = _module.D6.create_DC12(_1551_v120);
          let _1553_v122;
          _1553_v122 = _dafny.Seq.of(_1385_v0);
          let _1554_v123;
          _1554_v123 = _dafny.Map.Empty.slice().updateUnsafe((_1553_v122)[_module.__default.safeIndex(_1531_i12, new BigNumber((_1553_v122).length))],_1387_v2.f25);
          let _1555_v124;
          _1555_v124 = _dafny.Set.fromElements(_1550_v119);
          let _1556_v125;
          _1556_v125 = _module.D4.create_DC10(_1385_v0, new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_1555_v124).length)), _1531_i12)).cardinality()), _1387_v2.f25, _1387_v2.f25, _1531_i12);
          let _1557_v126;
          let _nw242 = new _module.C2();
          _nw242.__ctor(new BigNumber((_1554_v123).length));
          _1557_v126 = _nw242;
          let _1558_v127;
          _1558_v127 = _dafny.Seq.of((_1387_v2).f26);
          let _1559_v128;
          _1559_v128 = _dafny.Seq.of(_1558_v127);
          let _pat_let_tv31 = _1385_v0;
          let _index244 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_1549_v118).length));
          let _rhs233 = ((!_dafny.Seq.contains(_1529_v101, _1550_v119)) ? (_1552_v121) : (_1552_v121));
          let _rhs234 = (new BigNumber((_dafny.Seq.UnicodeFromString("ys")).length)).isLessThan((_dafny.ZERO).minus((_this).fm8((((_1554_v123).contains(_1531_i12)) ? ((_1554_v123).get(_1531_i12)) : ((_1387_v2).f26)), _dafny.Map.Empty.slice().updateUnsafe((function (_pat_let24_0) {
            return function (_1560_dt__update__tmp_h1) {
              return function (_pat_let25_0) {
                return function (_1561_dt__update_hcf13_h0) {
                  return _module.D4.create_DC10((_1560_dt__update__tmp_h1).dtor_cf12, _1561_dt__update_hcf13_h0, (_1560_dt__update__tmp_h1).dtor_cf14, (_1560_dt__update__tmp_h1).dtor_cf15, (_1560_dt__update__tmp_h1).dtor_cf16);
                }(_pat_let25_0);
              }(_pat_let_tv31);
            }(_pat_let24_0);
          }(_1556_v125)).dtor_cf12,_1550_v119), _1529_v101, globalState)));
          let _rhs235 = _1531_i12;
          let _rhs236 = (((_1440_v41).contains((_dafny.Set.fromElements(_1557_v126, _1557_v126, _1557_v126)).IsProperSubsetOf(_dafny.Set.fromElements(_1557_v126)))) ? ((_1440_v41).get((_dafny.Set.fromElements(_1557_v126, _1557_v126, _1557_v126)).IsProperSubsetOf(_dafny.Set.fromElements(_1557_v126)))) : (_dafny.Seq.IsPrefixOf(_dafny.Seq.of(_1558_v127, _1558_v127), _1559_v128)));
          let _rhs237 = _1387_v2.f25;
          let _lhs217 = _1549_v118;
          let _lhs218 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_1549_v118).length));
          let _lhs219 = globalState;
          let _lhs220 = globalState;
          let _lhs221 = globalState;
          let _lhs222 = globalState;
          _lhs217[_lhs218] = _rhs233;
          _lhs219.f2 = _rhs234;
          _lhs220.f7 = _rhs235;
          _lhs221.f2 = _rhs236;
          _lhs222.f13 = _rhs237;
          (globalState).f21 = _1387_v2.f25;
          _1385_v0 = _1531_i12;
          let _1562_v129;
          _1562_v129 = _dafny.Set.fromElements((_1387_v2).f26, (_1387_v2).f26);
          let _1563_v130;
          _1563_v130 = _dafny.Seq.of(_1562_v129);
          let _1564_v131;
          _1564_v131 = _dafny.MultiSet.fromElements(((_1563_v130)[_module.__default.safeIndex(_1531_i12, new BigNumber((_1563_v130).length))]).Intersect(_1562_v129), _1562_v129, _dafny.Set.fromElements(!(_module.__default.fm2(_1440_v41, globalState))));
          _1564_v131 = _dafny.MultiSet.fromElements(_1562_v129);
          let _1565_v132;
          let _nw243 = Array((new BigNumber(8)).toNumber()).fill(false);
          _1565_v132 = _nw243;
          let _1566_v133;
          _1566_v133 = _dafny.Map.Empty.slice().updateUnsafe(_1531_i12,new _dafny.CodePoint('c'.codePointAt(0)));
          let _1567_v134;
          _1567_v134 = _dafny.Map.Empty.slice().updateUnsafe(_1565_v132,(_this).fm6((_this).fm8((_1387_v2).f26, _1566_v133, _1529_v101, globalState), globalState));
          _1567_v134 = (_1567_v134).update(_1565_v132, !(_1387_v2.f25));
        }
        (globalState).f2 = _1387_v2.f25;
      }
      let _1568_v135;
      _1568_v135 = _dafny.Set.fromElements((_this).f24, _1387_v2.f25, (_this).f24, true);
      let _1569_v136;
      _1569_v136 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,new BigNumber((_1568_v135).length));
      let _hi8 = new BigNumber(((_1569_v136).Merge(_1569_v136)).length);
      for (let _1570_i14 = (_1385_v0).plus(_1385_v0); _1570_i14.isLessThan(_hi8); _1570_i14 = _1570_i14.plus(_dafny.ONE)) {
        let _1571_v137;
        _1571_v137 = (_1568_v135).Intersect(_module.__default.fm31(_1385_v0, (_1387_v2).f26, globalState));
        _1571_v137 = _1568_v135;
        let _1572_v138;
        _1572_v138 = _module.D13.create_DC27(false);
        _1572_v138 = (((_1387_v2).f26) ? (_1572_v138) : (_1572_v138));
        let _1573_v139;
        let _nw244 = new _module.C3();
        _nw244.__ctor();
        _1573_v139 = _nw244;
        let _nw245 = new _module.C3();
        _nw245.__ctor();
        _1573_v139 = _nw245;
        let _1574_v140;
        let _nw246 = Array((new BigNumber(17)).toNumber()).fill(false);
        _1574_v140 = _nw246;
        let _index245 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_1574_v140).length));
        (_1574_v140)[_index245] = false;
        let _index246 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_1574_v140).length));
        let _rhs238 = (_1387_v2).f26;
        let _rhs239 = _1385_v0;
        let _rhs240 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_1570_i14));
        let _rhs241 = !(false);
        let _lhs223 = globalState;
        let _lhs224 = globalState;
        let _lhs225 = globalState;
        let _lhs226 = _1574_v140;
        let _lhs227 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_1574_v140).length));
        _lhs223.f13 = _rhs238;
        _lhs224.f7 = _rhs239;
        _lhs225.f16 = _rhs240;
        _lhs226[_lhs227] = _rhs241;
      }
      let _hi9 = _1385_v0;
      for (let _1575_i15 = new BigNumber(169); _1575_i15.isLessThan(_hi9); _1575_i15 = _1575_i15.plus(_dafny.ONE)) {
        if ((_1387_v2).f26) {
          let _1576_v141;
          _1576_v141 = _1568_v135;
          let _1577_v142;
          let _nw247 = Array((new BigNumber(2)).toNumber());
          _nw247[(_dafny.ZERO).toNumber()] = (_dafny.Set.fromElements((_1387_v2).f26, true)).IsProperSubsetOf(_1568_v135);
          _nw247[(_dafny.ONE).toNumber()] = (_1387_v2).f26;
          _1577_v142 = _nw247;
          let _1578_v143;
          _1578_v143 = _dafny.Seq.of((_1387_v2).f26, _1387_v2.f25);
          let _index247 = _module.__default.safeIndex(new BigNumber(416), new BigNumber((_1577_v142).length));
          (_1577_v142)[_index247] = _dafny.Seq.IsProperPrefixOf(_1578_v143, _1578_v143);
          let _index248 = _module.__default.safeIndex(new BigNumber(416), new BigNumber((_1577_v142).length));
          let _rhs242 = (((_1575_i15).isLessThanOrEqualTo(_1385_v0)) ? (_1576_v141) : (_1576_v141));
          let _rhs243 = (function () {
            let _coll60 = new _dafny.Set();
            for (const _compr_60 of _dafny.IntegerRange(new BigNumber(872), new BigNumber(204))) {
              let _1579_v144 = _compr_60;
              if (((new BigNumber(872)).isLessThanOrEqualTo(_1579_v144)) && ((_1579_v144).isLessThan(new BigNumber(204)))) {
                _coll60.add((_1579_v144).minus(_1575_i15));
              }
            }
            return _coll60;
          }()).Union((_dafny.Set.fromElements(_1575_i15)).Union(_1386_v1));
          let _rhs244 = (_this).f24;
          let _rhs245 = _1387_v2.f25;
          let _lhs228 = _1387_v2;
          let _lhs229 = _1577_v142;
          let _lhs230 = _module.__default.safeIndex(new BigNumber(416), new BigNumber((_1577_v142).length));
          _1576_v141 = _rhs242;
          _1386_v1 = _rhs243;
          _lhs228.f25 = _rhs244;
          _lhs229[_lhs230] = _rhs245;
          (globalState).f2 = (((_1387_v2).f26) ? ((_this).f24) : ((_1577_v142)[_module.__default.safeIndex(new BigNumber(416), new BigNumber((_1577_v142).length))]));
          let _index249 = _module.__default.safeIndex(new BigNumber(416), new BigNumber((_1577_v142).length));
          (_1577_v142)[_index249] = ((!(_1387_v2.f25)) ? (!(_1385_v0).isEqualTo(new BigNumber((_1386_v1).length))) : (!(new BigNumber((_1568_v135).length)).isEqualTo(_1575_i15)));
          let _1580_v145;
          _1580_v145 = _dafny.Seq.of(new BigNumber((_1529_v101).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1387_v2.f25,_1385_v0)).length), _1575_i15, _1385_v0, _1575_i15);
          _1440_v41 = (_1440_v41).update(_1387_v2.f25, _dafny.Seq.contains(_1580_v145, (_dafny.ZERO).minus(_1385_v0)));
          let _index250 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((p0).length));
          (p0)[_index250] = (_dafny.ZERO).minus(_1575_i15);
          let _index251 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((p0).length));
          (p0)[_index251] = _1385_v0;
        } else {
          let _1581_v146;
          _1581_v146 = _dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(867)));
          let _rhs246 = ((_1568_v135).Union(_1568_v135)).Intersect(_module.__default.fm14(_1575_i15, (_1387_v2).f26, globalState));
          let _rhs247 = (_1581_v146)[_module.__default.safeIndex(_1385_v0, new BigNumber((_1581_v146).length))];
          _1568_v135 = _rhs246;
          _1386_v1 = _rhs247;
          (globalState).f7 = ((((_1569_v136).contains((_1387_v2).f26)) ? ((_1569_v136).get((_1387_v2).f26)) : (_1575_i15))).multipliedBy(((_1387_v2.f25) ? (_1385_v0) : (_1575_i15)));
          (globalState).f2 = _1387_v2.f25;
          let _index252 = _module.__default.safeIndex(new BigNumber(498), new BigNumber((p0).length));
          (p0)[_index252] = _1575_i15;
          let _index253 = _module.__default.safeIndex(new BigNumber(498), new BigNumber((p0).length));
          (p0)[_index253] = _module.__default.safeModuloInt(_1385_v0, _1385_v0);
          (_1387_v2).m5(_module.__default.fm2(_1440_v41, globalState), globalState);
        }
        (globalState).f7 = (_1385_v0).plus(_module.__default.safeDivisionInt(new BigNumber(559), _1385_v0));
        let _1582_v147;
        let _nw248 = Array((new BigNumber(22)).toNumber());
        _nw248[(_dafny.ZERO).toNumber()] = _1440_v41;
        _nw248[(_dafny.ONE).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_1387_v2.f25,true)).Merge(_1440_v41);
        _nw248[(new BigNumber(2)).toNumber()] = (_module.__default.fm20(_1385_v0, (_this).fm6(_1575_i15, globalState), (_1387_v2).f26, _1387_v2.f25, globalState)).update((_1387_v2).f26, (_this).fm7(_1387_v2.f25, (_dafny.ZERO).minus(_1385_v0), !(!((_1387_v2).f26)), globalState));
        _nw248[(new BigNumber(3)).toNumber()] = (_1440_v41).update(!((_this).f24), (_1387_v2).f26);
        _nw248[(new BigNumber(4)).toNumber()] = _1440_v41;
        _nw248[(new BigNumber(5)).toNumber()] = (_1440_v41).Merge(_1440_v41);
        _nw248[(new BigNumber(6)).toNumber()] = _1440_v41;
        _nw248[(new BigNumber(7)).toNumber()] = _1440_v41;
        _nw248[(new BigNumber(8)).toNumber()] = _1440_v41;
        _nw248[(new BigNumber(9)).toNumber()] = _1440_v41;
        _nw248[(new BigNumber(10)).toNumber()] = (_1440_v41).Merge(_1440_v41);
        _nw248[(new BigNumber(11)).toNumber()] = _1440_v41;
        _nw248[(new BigNumber(12)).toNumber()] = _1440_v41;
        _nw248[(new BigNumber(13)).toNumber()] = _1440_v41;
        _nw248[(new BigNumber(14)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_1387_v2).f26,false);
        _nw248[(new BigNumber(15)).toNumber()] = (_1440_v41).Merge(_1440_v41);
        _nw248[(new BigNumber(16)).toNumber()] = (_1440_v41).Merge(_1440_v41);
        _nw248[(new BigNumber(17)).toNumber()] = (_1440_v41).Merge(_1440_v41);
        _nw248[(new BigNumber(18)).toNumber()] = (_1440_v41).Merge(_dafny.Map.Empty.slice().updateUnsafe(!((_1387_v2).f26),_1387_v2.f25));
        _nw248[(new BigNumber(19)).toNumber()] = _1440_v41;
        _nw248[(new BigNumber(20)).toNumber()] = _1440_v41;
        _nw248[(new BigNumber(21)).toNumber()] = (_1440_v41).Merge(_1440_v41);
        _1582_v147 = _nw248;
        let _1583_v148;
        let _nw249 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Set.Empty);
        _1583_v148 = _nw249;
        let _index254 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_1583_v148).length));
        (_1583_v148)[_index254] = (_1568_v135).Union(_1568_v135);
        let _1584_v149;
        _1584_v149 = _dafny.Seq.of(_1575_i15);
        let _1585_v150;
        _1585_v150 = _dafny.Seq.of((_this).f24);
        let _index255 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_1583_v148).length));
        let _rhs248 = _1582_v147;
        let _rhs249 = _1387_v2.f25;
        let _rhs250 = (_1585_v150)[_module.__default.safeIndex(_1575_i15, new BigNumber((_1585_v150).length))];
        let _rhs251 = _1568_v135;
        let _rhs252 = _1584_v149;
        let _lhs231 = globalState;
        let _lhs232 = globalState;
        let _lhs233 = _1583_v148;
        let _lhs234 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_1583_v148).length));
        _1582_v147 = _rhs248;
        _lhs231.f0 = _rhs249;
        _lhs232.f2 = _rhs250;
        _lhs233[_lhs234] = _rhs251;
        _1584_v149 = _rhs252;
        let _index256 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((p0).length));
        (p0)[_index256] = _1385_v0;
        let _index257 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((p0).length));
        let _rhs253 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_1385_v0, new BigNumber(471)));
        let _rhs254 = _module.__default.fm2(_1440_v41, globalState);
        let _lhs235 = p0;
        let _lhs236 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((p0).length));
        let _lhs237 = globalState;
        _lhs235[_lhs236] = _rhs253;
        _lhs237.f21 = _rhs254;
      }
      return;
    }
    m1(p0, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = false;
      let _1586_v0;
      _1586_v0 = _dafny.Seq.of(p0);
      let _1587_v1;
      _1587_v1 = _dafny.Seq.of((_1586_v0)[_module.__default.safeIndex(p0, new BigNumber((_1586_v0).length))]);
      let _hi10 = p0;
      for (let _1588_i0 = (new BigNumber((_1587_v1).length)).minus(p0); _1588_i0.isLessThan(_hi10); _1588_i0 = _1588_i0.plus(_dafny.ONE)) {
        (globalState).f21 = (_this).f24;
        let _1589_v2;
        let _nw250 = new _module.C3();
        _nw250.__ctor();
        _1589_v2 = _nw250;
        let _1590_v3;
        let _nw251 = new _module.C2();
        _nw251.__ctor(p0);
        _1590_v3 = _nw251;
        if (true) {
          (globalState).f0 = true;
          let _1591_v4;
          _1591_v4 = new _dafny.CodePoint('k'.codePointAt(0));
          let _1592_v5;
          _1592_v5 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1591_v4)).length));
          let _1593_v6;
          _1593_v6 = _dafny.Set.fromElements(_module.__default.safeDivisionInt((_1590_v3).f28, _module.__default.fm0((_1590_v3).f28, new BigNumber((_1592_v5).cardinality()), new BigNumber((_1587_v1).length), (_this).f24, globalState)), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_module.__default.fm21(globalState)).length), (_1590_v3).f28), _1586_v0)).length), new BigNumber((_dafny.Seq.Concat(_1586_v0, _1587_v1)).length));
          let _1594_v7;
          _1594_v7 = _dafny.Map.Empty.slice().updateUnsafe((_1590_v3).f28,_1591_v4);
          let _1595_v8;
          _1595_v8 = _dafny.Seq.UnicodeFromString("gffxeber");
          let _1596_v9;
          _1596_v9 = _dafny.Map.Empty.slice().updateUnsafe((_1590_v3).f28,(_this).f24);
          let _1597_v11;
          _1597_v11 = _dafny.Seq.of((_this).f24, (_this).f24);
          let _1598_v12;
          _1598_v12 = _1597_v11;
          let _rhs255 = (_this).f24;
          let _rhs256 = (p0).isLessThanOrEqualTo((p0).minus((_this).fm8(false, _1594_v7, _1595_v8, globalState)));
          let _rhs257 = ((_1593_v6).Union(_1593_v6)).Union(function () {
            let _coll61 = new _dafny.Set();
            for (const _compr_61 of (_1596_v9).Keys.Elements) {
              let _1599_v10 = _compr_61;
              if ((_1596_v9).contains(_1599_v10)) {
                _coll61.add(_module.__default.safeModuloInt(_1599_v10, new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(879)), function (_1600_i1) {
                  return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("cftvrnjd")).length),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(96),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),new BigNumber(379))).length))).length))).length);
                }))).cardinality())));
              }
            }
            return _coll61;
          }());
          let _rhs258 = ((_1590_v3).fm17((_this).f24, _dafny.Seq.Create(_module.__default.abs(new BigNumber(187)), function (_1601_i2) {
            return new BigNumber((_dafny.Seq.of((_this).f24)).length);
          }), _1588_i0, _1598_v12, globalState)).plus(_module.__default.safeModuloInt((_1590_v3).f28, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,false)).length)));
          let _lhs238 = globalState;
          let _lhs239 = globalState;
          _lhs238.f21 = _rhs255;
          r1 = _rhs256;
          _1593_v6 = _rhs257;
          _lhs239.f7 = _rhs258;
          let _1602_v13;
          let _init47 = ((_1603_v8) => function (_1604_i3) {
            return _1603_v8;
          })(_1595_v8);
          let _nw252 = Array((new BigNumber(29)).toNumber());
          for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw252.length); _i0_47++) {
            _nw252[_i0_47] = _init47(new BigNumber(_i0_47));
          }
          _1602_v13 = _nw252;
          let _index258 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_1602_v13).length));
          (_1602_v13)[_index258] = _1595_v8;
          let _index259 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_1602_v13).length));
          (_1602_v13)[_index259] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(777)), ((_1605_v4) => function (_1606_i4) {
            return _1605_v4;
          })(_1591_v4));
          let _1607_v14;
          let _nw253 = Array((new BigNumber(2)).toNumber()).fill(false);
          _1607_v14 = _nw253;
          let _index260 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_1607_v14).length));
          (_1607_v14)[_index260] = (_this).f24;
          let _index261 = _module.__default.safeIndex(new BigNumber(762), new BigNumber((_1607_v14).length));
          (_1607_v14)[_index261] = (new BigNumber((_1586_v0).length)).isLessThanOrEqualTo(new BigNumber(-113));
          let _1608_v16;
          _1608_v16 = _dafny.MultiSet.fromElements(_dafny.Seq.of(p0));
          let _1609_v17;
          _1609_v17 = _module.D16.create_DC35(function () {
  let _coll62 = new _dafny.Map();
  for (const _compr_62 of (_1608_v16).Elements) {
    let _1610_v15 = _compr_62;
    if ((_1608_v16).contains(_1610_v15)) {
      _coll62.push([_1610_v15,_1597_v11]);
    }
  }
  return _coll62;
}());
          let _1611_v18;
          _1611_v18 = _module.D16.create_DC37(_1609_v17);
          let _index262 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_1602_v13).length));
          let _rhs259 = _dafny.Seq.Concat(_1595_v8, _dafny.Seq.Concat(_module.__default.fm1(_1588_i0, globalState), _1595_v8));
          let _rhs260 = _dafny.Seq.Concat(_dafny.Seq.Concat((_1602_v13)[_module.__default.safeIndex(new BigNumber(731), new BigNumber((_1602_v13).length))], _1595_v8), _1595_v8);
          let _rhs261 = _1591_v4;
          let _rhs262 = _1611_v18;
          let _lhs240 = globalState;
          let _lhs241 = _1602_v13;
          let _lhs242 = _module.__default.safeIndex(new BigNumber(731), new BigNumber((_1602_v13).length));
          _lhs240.f8 = _rhs259;
          _lhs241[_lhs242] = _rhs260;
          _1591_v4 = _rhs261;
          _1611_v18 = _rhs262;
        } else {
          let _1612_v19;
          let _nw254 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Map.Empty);
          _1612_v19 = _nw254;
          let _1613_v20;
          _1613_v20 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1612_v19);
          (globalState).f22 = (((_1613_v20).contains(_module.__default.fm0((_1590_v3).f28, _1588_i0, p0, (_this).f24, globalState))) ? ((_1613_v20).get(_module.__default.fm0((_1590_v3).f28, _1588_i0, p0, (_this).f24, globalState))) : (_1612_v19));
          let _1614_v21;
          let _nw255 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.of());
          _1614_v21 = _nw255;
          let _1615_v22;
          _1615_v22 = _dafny.Seq.of((_this).f24);
          let _index263 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_1614_v21).length));
          (_1614_v21)[_index263] = _1615_v22;
          let _1616_v23;
          let _nw256 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
          _1616_v23 = _nw256;
          let _1617_v24;
          _1617_v24 = _dafny.Seq.of(_1616_v23, _1616_v23, _1616_v23, _1616_v23, _1616_v23);
          let _1618_v25;
          _1618_v25 = _dafny.Set.fromElements(_1588_i0);
          let _index264 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_1614_v21).length));
          let _rhs263 = (_1617_v24)[_module.__default.safeIndex(_1588_i0, new BigNumber((_1617_v24).length))];
          let _rhs264 = (_1618_v25).IsDisjointFrom((_module.D14.create_DC28(_1618_v25)).dtor_cf37);
          let _rhs265 = p0;
          let _rhs266 = _1615_v22;
          let _lhs243 = globalState;
          let _lhs244 = globalState;
          let _lhs245 = _1614_v21;
          let _lhs246 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_1614_v21).length));
          r0 = _rhs263;
          _lhs243.f13 = _rhs264;
          _lhs244.f7 = _rhs265;
          _lhs245[_lhs246] = _rhs266;
          let _1619_v26;
          _1619_v26 = _dafny.Map.Empty.slice().updateUnsafe(((_1590_v3).f28).multipliedBy(new BigNumber(899)),!(((_this).f24) || ((_this).f24)));
          _1619_v26 = ((!((_this).f24) || ((_this).f24)) ? (function () {
            let _coll63 = new _dafny.Map();
            for (const _compr_63 of (function () {
              let _coll64 = new _dafny.Set();
              for (const _compr_64 of (_dafny.MultiSet.fromElements((_1590_v3).f28, _1588_i0)).Elements) {
                let _1620_v28 = _compr_64;
                if ((_dafny.MultiSet.fromElements((_1590_v3).f28, _1588_i0)).contains(_1620_v28)) {
                  _coll64.add((_1620_v28).minus(new BigNumber(227)));
                }
              }
              return _coll64;
            }()).Elements) {
              let _1621_v27 = _compr_63;
              if ((function () {
                let _coll65 = new _dafny.Set();
                for (const _compr_65 of (_dafny.MultiSet.fromElements((_1590_v3).f28, _1588_i0)).Elements) {
                  let _1622_v28 = _compr_65;
                  if ((_dafny.MultiSet.fromElements((_1590_v3).f28, _1588_i0)).contains(_1622_v28)) {
                    _coll65.add((_1622_v28).minus(new BigNumber(227)));
                  }
                }
                return _coll65;
              }()).contains(_1621_v27)) {
                _coll63.push([(_1621_v27).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1590_v3).f28,_1588_i0)).length)),(_this).f24]);
              }
            }
            return _coll63;
          }()) : ((_1619_v26).Merge(_dafny.Map.Empty.slice().updateUnsafe((_1590_v3).f28,(_this).f24))));
          let _1623_v29;
          _1623_v29 = _dafny.Seq.UnicodeFromString("ccmntf");
          let _1624_v30;
          _1624_v30 = _dafny.Map.Empty.slice().updateUnsafe((_1614_v21)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_1614_v21).length))],_1623_v29);
          let _1625_v31;
          _1625_v31 = _dafny.MultiSet.fromElements(p0, (_1590_v3).f28);
          _module.__default.m0(((_1614_v21)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_1614_v21).length))])[_module.__default.safeIndex((_1590_v3).f28, new BigNumber(((_1614_v21)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_1614_v21).length))]).length))], _1588_i0, _1624_v30, new BigNumber(((_1625_v31).Difference(_1625_v31)).cardinality()), globalState);
          (globalState).f16 = _1588_i0;
        }
      }
      let _1626_v32;
      _1626_v32 = new _dafny.CodePoint('c'.codePointAt(0));
      _1626_v32 = _1626_v32;
      (globalState).f0 = !(!((_this).f24) || ((_this).f24));
      let _1627_v33;
      _1627_v33 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f24);
      let _rhs267 = p0;
      let _rhs268 = p0;
      let _rhs269 = _module.__default.safeDivisionInt(p0, new BigNumber(((_1627_v33).Merge(_1627_v33)).length));
      let _lhs247 = globalState;
      let _lhs248 = globalState;
      let _lhs249 = globalState;
      _lhs247.f16 = _rhs267;
      _lhs248.f16 = _rhs268;
      _lhs249.f16 = _rhs269;
      let _1628_v34;
      let _init48 = ((_1629_p0) => function (_1630_i5) {
        return (_1630_i5).minus(_1629_p0);
      })(p0);
      let _nw257 = Array((new BigNumber(19)).toNumber());
      for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw257.length); _i0_48++) {
        _nw257[_i0_48] = _init48(new BigNumber(_i0_48));
      }
      _1628_v34 = _nw257;
      r0 = _1628_v34;
      if ((_this).f24) {
        let _index265 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_1628_v34).length));
        (_1628_v34)[_index265] = p0;
        let _1631_v35;
        let _nw258 = Array((new BigNumber(13)).toNumber()).fill(false);
        _1631_v35 = _nw258;
        let _index266 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_1631_v35).length));
        (_1631_v35)[_index266] = (_this).f24;
        let _index267 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_1628_v34).length));
        let _index268 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_1631_v35).length));
        let _rhs270 = (_dafny.ZERO).minus(new BigNumber(-510));
        let _rhs271 = !((_this).f24);
        let _rhs272 = (p0).minus(p0);
        let _rhs273 = (p0).multipliedBy((p0).minus(p0));
        let _lhs250 = _1628_v34;
        let _lhs251 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_1628_v34).length));
        let _lhs252 = _1631_v35;
        let _lhs253 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_1631_v35).length));
        let _lhs254 = globalState;
        let _lhs255 = globalState;
        _lhs250[_lhs251] = _rhs270;
        _lhs252[_lhs253] = _rhs271;
        _lhs254.f16 = _rhs272;
        _lhs255.f16 = _rhs273;
        let _1632_v36;
        _1632_v36 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f24),new BigNumber(-656));
        (globalState).f7 = (((_1632_v36).contains(((_this).f24) === ((_1631_v35)[_module.__default.safeIndex(new BigNumber(633), new BigNumber((_1631_v35).length))]))) ? ((_1632_v36).get(((_this).f24) === ((_1631_v35)[_module.__default.safeIndex(new BigNumber(633), new BigNumber((_1631_v35).length))]))) : (p0));
        let _1633_v37;
        _1633_v37 = _dafny.Set.fromElements(new BigNumber(533));
        let _1634_v38;
        _1634_v38 = _dafny.Seq.UnicodeFromString("ospctjgb");
        let _1635_v39;
        let _nw259 = Array((new BigNumber(19)).toNumber());
        _nw259[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1633_v37).length));
        _nw259[(_dafny.ONE).toNumber()] = new BigNumber((_1634_v38).length);
        _nw259[(new BigNumber(2)).toNumber()] = (_1628_v34)[_module.__default.safeIndex(new BigNumber(748), new BigNumber((_1628_v34).length))];
        _nw259[(new BigNumber(3)).toNumber()] = (_1628_v34)[_module.__default.safeIndex(new BigNumber(748), new BigNumber((_1628_v34).length))];
        _nw259[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(p0);
        _nw259[(new BigNumber(5)).toNumber()] = (_1628_v34)[_module.__default.safeIndex(new BigNumber(748), new BigNumber((_1628_v34).length))];
        _nw259[(new BigNumber(6)).toNumber()] = (_1628_v34)[_module.__default.safeIndex(new BigNumber(748), new BigNumber((_1628_v34).length))];
        _nw259[(new BigNumber(7)).toNumber()] = (_1628_v34)[_module.__default.safeIndex(new BigNumber(748), new BigNumber((_1628_v34).length))];
        _nw259[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus((_1628_v34)[_module.__default.safeIndex(new BigNumber(748), new BigNumber((_1628_v34).length))]);
        _nw259[(new BigNumber(9)).toNumber()] = new BigNumber((_1634_v38).length);
        _nw259[(new BigNumber(10)).toNumber()] = (_1628_v34)[_module.__default.safeIndex(new BigNumber(748), new BigNumber((_1628_v34).length))];
        _nw259[(new BigNumber(11)).toNumber()] = (_this).fm8((_this).f24, (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(692),_1626_v32)).update(new BigNumber(-65), _1626_v32), _1634_v38, globalState);
        _nw259[(new BigNumber(12)).toNumber()] = p0;
        _nw259[(new BigNumber(13)).toNumber()] = new BigNumber(-377);
        _nw259[(new BigNumber(14)).toNumber()] = new BigNumber(143);
        _nw259[(new BigNumber(15)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!((_1631_v35)[_module.__default.safeIndex(new BigNumber(633), new BigNumber((_1631_v35).length))]),(_this).f24)).length);
        _nw259[(new BigNumber(16)).toNumber()] = (_1628_v34)[_module.__default.safeIndex(new BigNumber(748), new BigNumber((_1628_v34).length))];
        _nw259[(new BigNumber(17)).toNumber()] = (_1628_v34)[_module.__default.safeIndex(new BigNumber(748), new BigNumber((_1628_v34).length))];
        _nw259[(new BigNumber(18)).toNumber()] = p0;
        _1635_v39 = _nw259;
        let _1636_v40;
        _1636_v40 = _module.D6.create_DC12(_dafny.Seq.of(_1635_v39));
        let _1637_v42;
        _1637_v42 = _dafny.Map.Empty.slice().updateUnsafe(_1626_v32,(_1631_v35)[_module.__default.safeIndex(new BigNumber(633), new BigNumber((_1631_v35).length))]);
        let _1638_v43;
        _1638_v43 = _dafny.Map.Empty.slice().updateUnsafe((_1628_v34)[_module.__default.safeIndex(new BigNumber(748), new BigNumber((_1628_v34).length))],_1637_v42);
        let _index269 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_1631_v35).length));
        let _rhs274 = (_1631_v35)[_module.__default.safeIndex(new BigNumber(633), new BigNumber((_1631_v35).length))];
        let _rhs275 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(231)), ((_1639_v32) => function (_1640_i6) {
          return _1639_v32;
        })(_1626_v32));
        let _rhs276 = !(!((function () {
          let _coll66 = new _dafny.Map();
          for (const _compr_66 of _dafny.IntegerRange(new BigNumber(-222), new BigNumber(403))) {
            let _1641_v41 = _compr_66;
            if (((new BigNumber(-222)).isLessThanOrEqualTo(_1641_v41)) && ((_1641_v41).isLessThan(new BigNumber(403)))) {
              _coll66.push([_module.__default.safeDivisionInt(_1641_v41, (_1628_v34)[_module.__default.safeIndex(new BigNumber(748), new BigNumber((_1628_v34).length))]),_dafny.Map.Empty.slice().updateUnsafe(_1626_v32,(_1631_v35)[_module.__default.safeIndex(new BigNumber(633), new BigNumber((_1631_v35).length))])]);
            }
          }
          return _coll66;
        }()).Merge(_1638_v43)).equals(_module.__default.fm48(_1626_v32, (_1628_v34)[_module.__default.safeIndex(new BigNumber(748), new BigNumber((_1628_v34).length))], globalState)));
        let _rhs277 = _1636_v40;
        let _rhs278 = !((_this).f24) || ((_1631_v35)[_module.__default.safeIndex(new BigNumber(633), new BigNumber((_1631_v35).length))]);
        let _lhs256 = _1631_v35;
        let _lhs257 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_1631_v35).length));
        let _lhs258 = globalState;
        let _lhs259 = globalState;
        let _lhs260 = globalState;
        _lhs256[_lhs257] = _rhs274;
        _lhs258.f20 = _rhs275;
        _lhs259.f0 = _rhs276;
        _1636_v40 = _rhs277;
        _lhs260.f13 = _rhs278;
        (globalState).f13 = !(true) || ((p0).isLessThanOrEqualTo((_1628_v34)[_module.__default.safeIndex(new BigNumber(748), new BigNumber((_1628_v34).length))]));
        let _1642_v44;
        let _nw260 = new _module.C2();
        _nw260.__ctor((_1628_v34)[_module.__default.safeIndex(new BigNumber(748), new BigNumber((_1628_v34).length))]);
        _1642_v44 = _nw260;
        _1642_v44 = _1642_v44;
      } else {
        let _1643_v45;
        _1643_v45 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm7((_this).f24, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),(_this).f24)).length), (_this).f24, globalState),(_this).f24);
        let _1644_v46;
        _1644_v46 = _dafny.Set.fromElements(_module.__default.fm2(_1643_v45, globalState));
        let _1645_v47;
        _1645_v47 = _dafny.Map.Empty.slice().updateUnsafe(_1644_v46,_module.__default.fm18((_this).f24, _1626_v32, (_this).fm6(p0, globalState), globalState));
        let _1646_v48;
        _1646_v48 = _dafny.Seq.of(true);
        _1645_v47 = (_1645_v47).update(_module.__default.fm14(p0, (_this).f24, globalState), _1646_v48);
        let _1647_v49;
        let _nw261 = Array((new BigNumber(7)).toNumber()).fill([]);
        _1647_v49 = _nw261;
        _1647_v49 = _1647_v49;
        let _index270 = _module.__default.safeIndex(new BigNumber(70), new BigNumber((_1628_v34).length));
        (_1628_v34)[_index270] = p0;
        let _index271 = _module.__default.safeIndex(new BigNumber(70), new BigNumber((_1628_v34).length));
        (_1628_v34)[_index271] = new BigNumber((((_dafny.Set.fromElements((_this).f24, true)).Difference(_1644_v46)).Difference(_1644_v46)).length);
        _1646_v48 = _1646_v48;
        let _nw262 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
        r0 = _nw262;
      }
      r0 = _1628_v34;
      r1 = (_this).f24;
      return [r0, r1];
    }
    get f24() {
      let _this = this;
      return _this._f24;
    };
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm7(p0, p1, p2, globalState) {
      let _this = this;
      return false;
    };
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (((false) ? (_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("xdbnkvr"))) : (_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.Create(_module.__default.abs(new BigNumber(585)), function (_1648_i0) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      }))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.UnicodeFromString("bgtt")));
    };
    fm6(p0, globalState) {
      let _this = this;
      return true;
    };
    m2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = [];
      (globalState).f13 = p1;
      (globalState).f16 = new BigNumber(-429);
      let _hi11 = p3;
      for (let _1649_i0 = p3; _1649_i0.isLessThan(_hi11); _1649_i0 = _1649_i0.plus(_dafny.ONE)) {
        let _1650_v0;
        _1650_v0 = _dafny.MultiSet.fromElements(new BigNumber(307));
        let _1651_v1;
        _1651_v1 = _dafny.MultiSet.fromElements(_1650_v0);
        (globalState).f0 = ((_dafny.MultiSet.fromElements(_1650_v0, _1650_v0, _1650_v0, (_1650_v0).update(_1649_i0, _module.__default.abs(new BigNumber(-346))))).IsProperSubsetOf((_1651_v1).update(_1650_v0, _module.__default.abs(p0)))) === (p1);
        let _1652_v2;
        _1652_v2 = _dafny.Seq.of(!(p1), p1);
        let _1653_v3;
        let _nw263 = new _module.C0();
        _nw263.__ctor((_1652_v2)[_module.__default.safeIndex(p0, new BigNumber((_1652_v2).length))]);
        _1653_v3 = _nw263;
        let _1654_v4;
        _1654_v4 = _dafny.Seq.of(p3);
        if ((((p1) ? (_1649_i0) : (_1649_i0))).isLessThanOrEqualTo(new BigNumber((_1654_v4).length))) {
          (globalState).f7 = p0;
          let _1655_v5;
          let _nw264 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1655_v5 = _nw264;
          let _1656_v6;
          _1656_v6 = _dafny.Seq.UnicodeFromString("odwrbay");
          let _index272 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_1655_v5).length));
          (_1655_v5)[_index272] = _1656_v6;
          let _index273 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_1655_v5).length));
          (_1655_v5)[_index273] = _module.__default.fm1(new BigNumber(151), globalState);
          (globalState).f16 = _1649_i0;
          let _1657_v7;
          let _nw265 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
          _1657_v7 = _nw265;
          _1657_v7 = _1657_v7;
          let _1658_v8;
          _1658_v8 = _dafny.Seq.of((_1655_v5)[_module.__default.safeIndex(new BigNumber(423), new BigNumber((_1655_v5).length))], (_1655_v5)[_module.__default.safeIndex(new BigNumber(423), new BigNumber((_1655_v5).length))], (_1655_v5)[_module.__default.safeIndex(new BigNumber(423), new BigNumber((_1655_v5).length))], (_1655_v5)[_module.__default.safeIndex(new BigNumber(423), new BigNumber((_1655_v5).length))], _module.__default.fm1(_1649_i0, globalState));
          let _1659_v9;
          _1659_v9 = _dafny.MultiSet.fromElements(_1653_v3.f27);
          let _index274 = _module.__default.safeIndex(new BigNumber(423), new BigNumber((_1655_v5).length));
          (_1655_v5)[_index274] = _dafny.Seq.Concat((_1658_v8)[_module.__default.safeIndex((_dafny.ZERO).minus((_1654_v4)[_module.__default.safeIndex(new BigNumber((_1659_v9).cardinality()), new BigNumber((_1654_v4).length))]), new BigNumber((_1658_v8).length))], _1656_v6);
        } else {
          let _index275 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((p2).length));
          (p2)[_index275] = (_1654_v4)[_module.__default.safeIndex(new BigNumber((_1652_v2).length), new BigNumber((_1654_v4).length))];
          let _index276 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((p2).length));
          (p2)[_index276] = p3;
          (_1653_v3).f27 = _1653_v3.f27;
          let _1660_v10;
          _1660_v10 = _dafny.Set.fromElements(_1653_v3.f27, (_this).fm6(p3, globalState));
          let _1661_v11;
          let _nw266 = Array((new BigNumber(15)).toNumber());
          _nw266[(_dafny.ZERO).toNumber()] = _1660_v10;
          _nw266[(_dafny.ONE).toNumber()] = _1660_v10;
          _nw266[(new BigNumber(2)).toNumber()] = _1660_v10;
          _nw266[(new BigNumber(3)).toNumber()] = _1660_v10;
          _nw266[(new BigNumber(4)).toNumber()] = _1660_v10;
          _nw266[(new BigNumber(5)).toNumber()] = _1660_v10;
          _nw266[(new BigNumber(6)).toNumber()] = _1660_v10;
          _nw266[(new BigNumber(7)).toNumber()] = _dafny.Set.fromElements(_1653_v3.f27, _1653_v3.f27, _1653_v3.f27);
          _nw266[(new BigNumber(8)).toNumber()] = _1660_v10;
          _nw266[(new BigNumber(9)).toNumber()] = _1660_v10;
          _nw266[(new BigNumber(10)).toNumber()] = _1660_v10;
          _nw266[(new BigNumber(11)).toNumber()] = _dafny.Set.fromElements(_1653_v3.f27);
          _nw266[(new BigNumber(12)).toNumber()] = _1660_v10;
          _nw266[(new BigNumber(13)).toNumber()] = _1660_v10;
          _nw266[(new BigNumber(14)).toNumber()] = _1660_v10;
          _1661_v11 = _nw266;
          let _1662_v12;
          _1662_v12 = _module.D7.create_DC16(((_1653_v3.f27) ? (_1661_v11) : (_1661_v11)));
          _1662_v12 = _1662_v12;
          let _1663_v13;
          let _nw267 = new _module.C4();
          _nw267.__ctor(p1, _1653_v3.f27);
          _1663_v13 = _nw267;
          let _1664_v14;
          _1664_v14 = _dafny.Map.Empty.slice().updateUnsafe(_1653_v3.f27,_1663_v13);
          let _1665_v15;
          _1665_v15 = _dafny.Map.Empty.slice().updateUnsafe((((_1664_v14).contains(_1663_v13.f25)) ? ((_1664_v14).get(_1663_v13.f25)) : (_1663_v13)),p0);
          _1665_v15 = (_1665_v15).update(_1663_v13, (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.fm0(new BigNumber(511), (p2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((p2).length))], (p2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((p2).length))], p1, globalState))));
          let _1666_v16;
          _1666_v16 = _dafny.Map.Empty.slice().updateUnsafe((p2)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((p2).length))],!(_1653_v3.f27));
          let _1667_v17;
          let _nw268 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Seq.of());
          _1667_v17 = _nw268;
          let _1668_v18;
          _1668_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1666_v16,_1667_v17);
          _1668_v18 = (_1668_v18).update(_1666_v16, _1667_v17);
        }
        let _1669_v19;
        _1669_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1649_i0,_1653_v3.f27);
        let _1670_v20;
        _1670_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1669_v19,p2);
        _1670_v20 = ((_1670_v20).Merge((_1670_v20).update(_dafny.Map.Empty.slice().updateUnsafe(_1649_i0,p1), p2))).Merge(_1670_v20);
      }
      let _1671_v21;
      _1671_v21 = _dafny.Map.Empty.slice().updateUnsafe(p3,p1);
      (globalState).f2 = (_module.D19.create_DC47(_module.D1.create_DC2(_1671_v21), p3, p1)).dtor_cf64;
      (globalState).f2 = (p1) === (p1);
      let _1672_v22;
      let _nw269 = Array((new BigNumber(3)).toNumber());
      _nw269[(_dafny.ZERO).toNumber()] = ((p1) ? (false) : (!(p1)));
      _nw269[(_dafny.ONE).toNumber()] = p1;
      _nw269[(new BigNumber(2)).toNumber()] = p1;
      _1672_v22 = _nw269;
      let _index277 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_1672_v22).length));
      (_1672_v22)[_index277] = !(!(p1)) || (true);
      let _index278 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_1672_v22).length));
      (_1672_v22)[_index278] = p1;
      let _1673_v23;
      _1673_v23 = new _dafny.CodePoint('e'.codePointAt(0));
      let _1674_v24;
      _1674_v24 = _dafny.Seq.of(p0, new BigNumber(325));
      let _1675_v25;
      _1675_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1673_v23,p1);
      let _1676_v26;
      _1676_v26 = _dafny.MultiSet.fromElements((_module.__default.fm49(_1673_v23, _1674_v24, new BigNumber(813), _1673_v23, globalState)).dtor_cf60, (((_1675_v25).contains(_1673_v23)) ? ((_1675_v25).get(_1673_v23)) : (p1)), p1);
      let _index279 = _module.__default.safeIndex(new BigNumber(304), new BigNumber((p2).length));
      (p2)[_index279] = new BigNumber((_1676_v26).cardinality());
      let _index280 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_1672_v22).length));
      let _index281 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_1672_v22).length));
      let _index282 = _module.__default.safeIndex(new BigNumber(304), new BigNumber((p2).length));
      let _rhs279 = p1;
      let _rhs280 = (p0).isLessThan(p0);
      let _rhs281 = _1676_v26;
      let _rhs282 = p3;
      let _lhs261 = _1672_v22;
      let _lhs262 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_1672_v22).length));
      let _lhs263 = _1672_v22;
      let _lhs264 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_1672_v22).length));
      let _lhs265 = p2;
      let _lhs266 = _module.__default.safeIndex(new BigNumber(304), new BigNumber((p2).length));
      _lhs261[_lhs262] = _rhs279;
      _lhs263[_lhs264] = _rhs280;
      _1676_v26 = _rhs281;
      _lhs265[_lhs266] = _rhs282;
      r0 = ((p1) === (p1)) || (p1);
      r1 = _module.__default.safeModuloInt((((_1672_v22)[_module.__default.safeIndex(new BigNumber(13), new BigNumber((_1672_v22).length))]) ? (p0) : ((p2)[_module.__default.safeIndex(new BigNumber(304), new BigNumber((p2).length))])), new BigNumber(133));
      let _1677_v27;
      let _init49 = ((_1678_v22, _1679_v23, _1680_p1) => function (_1681_i1) {
        return (_dafny.Map.Empty.slice().updateUnsafe((_1678_v22)[_module.__default.safeIndex(new BigNumber(13), new BigNumber((_1678_v22).length))],_module.D3.create_DC6(_dafny.Seq.Create(_module.__default.abs(new BigNumber(490)), ((_1682_v23) => function (_1683_i2) {
  return _1682_v23;
})(_1679_v23)), new _dafny.CodePoint('b'.codePointAt(0)), _1680_p1))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1680_p1,_module.D3.create_DC6(_dafny.Seq.UnicodeFromString("whrnsnix"), _1679_v23, (_1678_v22)[_module.__default.safeIndex(new BigNumber(13), new BigNumber((_1678_v22).length))])));
      })(_1672_v22, _1673_v23, p1);
      let _nw270 = Array((new BigNumber(8)).toNumber());
      for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw270.length); _i0_49++) {
        _nw270[_i0_49] = _init49(new BigNumber(_i0_49));
      }
      _1677_v27 = _nw270;
      r2 = _1677_v27;
      return [r0, r1, r2];
    }
    m3(p0, globalState) {
      let _this = this;
      let _1684_v0;
      _1684_v0 = new BigNumber(685);
      let _1685_v1;
      let _nw271 = new _module.C5();
      _nw271.__ctor((_1684_v0).isLessThanOrEqualTo(_1684_v0));
      _1685_v1 = _nw271;
      let _1686_v2;
      let _nw272 = new _module.C3();
      _nw272.__ctor();
      _1686_v2 = _nw272;
      (globalState).f0 = true;
      let _1687_v3;
      _1687_v3 = _dafny.Map.Empty.slice().updateUnsafe((_1685_v1).f24,false);
      let _1688_v4;
      _1688_v4 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2(_1687_v3, globalState),_module.__default.fm0(_1684_v0, new BigNumber((_dafny.Seq.UnicodeFromString("brysun")).length), _1684_v0, (_1685_v1).f24, globalState));
      let _index283 = _module.__default.safeIndex(new BigNumber(277), new BigNumber((p0).length));
      (p0)[_index283] = new BigNumber((_1688_v4).length);
      let _1689_v5;
      _1689_v5 = _module.D19.create_DC48(_1684_v0, (_1685_v1).f24);
      let _1690_v6;
      _1690_v6 = _dafny.Seq.of((_1685_v1).f24, (_1685_v1).f24);
      let _1691_v7;
      _1691_v7 = _dafny.MultiSet.fromElements(_1687_v3);
      let _1692_v8;
      _1692_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1691_v7).cardinality()),_1684_v0);
      let _1693_v9;
      _1693_v9 = _dafny.Set.fromElements(_1692_v8);
      let _1694_v10;
      _1694_v10 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1693_v9).length),(_1685_v1).f24);
      let _1695_v11;
      let _nw273 = Array((new BigNumber(29)).toNumber());
      _nw273[(_dafny.ZERO).toNumber()] = ((!((_1685_v1).f24)) ? (!((_1685_v1).f24)) : ((_1685_v1).f24));
      _nw273[(_dafny.ONE).toNumber()] = (_1689_v5).dtor_cf66;
      _nw273[(new BigNumber(2)).toNumber()] = (_1685_v1).f24;
      _nw273[(new BigNumber(3)).toNumber()] = (_1685_v1).f24;
      _nw273[(new BigNumber(4)).toNumber()] = (_1685_v1).f24;
      _nw273[(new BigNumber(5)).toNumber()] = (_1685_v1).f24;
      _nw273[(new BigNumber(6)).toNumber()] = !((_1685_v1).f24);
      _nw273[(new BigNumber(7)).toNumber()] = (_1685_v1).f24;
      _nw273[(new BigNumber(8)).toNumber()] = (((((_1687_v3).contains((_1685_v1).f24)) ? ((_1687_v3).get((_1685_v1).f24)) : ((_1685_v1).f24))) ? ((_1685_v1).f24) : (false));
      _nw273[(new BigNumber(9)).toNumber()] = (_1685_v1).f24;
      _nw273[(new BigNumber(10)).toNumber()] = (_1685_v1).f24;
      _nw273[(new BigNumber(11)).toNumber()] = (_1685_v1).f24;
      _nw273[(new BigNumber(12)).toNumber()] = (_1685_v1).f24;
      _nw273[(new BigNumber(13)).toNumber()] = (_1685_v1).f24;
      _nw273[(new BigNumber(14)).toNumber()] = (_1685_v1).f24;
      _nw273[(new BigNumber(15)).toNumber()] = (_1685_v1).f24;
      _nw273[(new BigNumber(16)).toNumber()] = false;
      _nw273[(new BigNumber(17)).toNumber()] = !((_1685_v1).f24);
      _nw273[(new BigNumber(18)).toNumber()] = !(_dafny.Seq.IsPrefixOf(_1690_v6, _1690_v6));
      _nw273[(new BigNumber(19)).toNumber()] = (_1685_v1).f24;
      _nw273[(new BigNumber(20)).toNumber()] = (_1685_v1).f24;
      _nw273[(new BigNumber(21)).toNumber()] = (_1685_v1).f24;
      _nw273[(new BigNumber(22)).toNumber()] = (_1685_v1).f24;
      _nw273[(new BigNumber(23)).toNumber()] = ((false) ? ((_1685_v1).f24) : ((_1685_v1).f24));
      _nw273[(new BigNumber(24)).toNumber()] = (!((_1685_v1).f24)) || ((_1685_v1).f24);
      _nw273[(new BigNumber(25)).toNumber()] = (((_1694_v10).contains(_1684_v0)) ? ((_1694_v10).get(_1684_v0)) : (false));
      _nw273[(new BigNumber(26)).toNumber()] = !(false);
      _nw273[(new BigNumber(27)).toNumber()] = false;
      _nw273[(new BigNumber(28)).toNumber()] = (new BigNumber((_dafny.Seq.of((_1685_v1).f24)).length)).isLessThanOrEqualTo(_1684_v0);
      _1695_v11 = _nw273;
      let _1696_v12;
      _1696_v12 = new _dafny.CodePoint('n'.codePointAt(0));
      let _1697_v13;
      _1697_v13 = _module.D20.create_DC49(_dafny.Map.Empty.slice().updateUnsafe(_1684_v0,_1696_v12));
      let _1698_v14;
      _1698_v14 = _dafny.Map.Empty.slice().updateUnsafe(!((_1685_v1).f24),(_1697_v13).dtor_cf67);
      let _1699_v15;
      _1699_v15 = _dafny.Seq.UnicodeFromString("snhrb");
      let _index284 = _module.__default.safeIndex(new BigNumber(277), new BigNumber((p0).length));
      let _rhs283 = _module.__default.safeDivisionInt((_1684_v0).multipliedBy(new BigNumber(229)), (_dafny.ZERO).minus((_1685_v1).fm8((_1690_v6)[_module.__default.safeIndex(_1684_v0, new BigNumber((_1690_v6).length))], (((_1698_v14).contains((_1685_v1).f24)) ? ((_1698_v14).get((_1685_v1).f24)) : (_dafny.Map.Empty.slice().updateUnsafe(_1684_v0,_1696_v12))), _1699_v15, globalState)));
      let _rhs284 = (((_1687_v3).contains((_1685_v1).f24)) ? ((_1687_v3).get((_1685_v1).f24)) : ((_this).fm6(_1684_v0, globalState)));
      let _rhs285 = (_1684_v0).minus(_1684_v0);
      let _rhs286 = _1695_v11;
      let _lhs267 = p0;
      let _lhs268 = _module.__default.safeIndex(new BigNumber(277), new BigNumber((p0).length));
      let _lhs269 = globalState;
      let _lhs270 = globalState;
      _lhs267[_lhs268] = _rhs283;
      _lhs269.f2 = _rhs284;
      _lhs270.f16 = _rhs285;
      _1695_v11 = _rhs286;
      let _1700_i0;
      _1700_i0 = _dafny.ZERO;
      L9: {
        while ((_1685_v1).f24) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1700_i0)) {
              break L9;
            }
            _1700_i0 = (_1700_i0).plus(_dafny.ONE);
            (globalState).f16 = _module.__default.safeModuloInt((((_1685_v1).f24) ? (_1684_v0) : (_1684_v0)), _1684_v0);
            _1694_v10 = (_1694_v10).update(new BigNumber(145), (_1685_v1).f24);
            let _1701_v16;
            let _init50 = ((_1702_v1, _1703_v6) => function (_1704_i1) {
              return _module.D16.create_DC36((_1702_v1).f24, _dafny.MultiSet.FromArray(_1703_v6), (_1702_v1).f24);
            })(_1685_v1, _1690_v6);
            let _nw274 = Array((new BigNumber(10)).toNumber());
            for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw274.length); _i0_50++) {
              _nw274[_i0_50] = _init50(new BigNumber(_i0_50));
            }
            _1701_v16 = _nw274;
            let _1705_v17;
            _1705_v17 = _module.D17.create_DC38(_1701_v16);
            let _1706_v18;
            _1706_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1705_v17,_1688_v4);
            let _1707_v19;
            _1707_v19 = _dafny.Seq.of(_1706_v18, _1706_v18);
            let _1708_v20;
            _1708_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1684_v0,_1688_v4);
            let _1709_v21;
            _1709_v21 = _dafny.Map.Empty.slice().updateUnsafe((_1685_v1).f24,_module.__default.fm35((_1685_v1).f24, new BigNumber((_1699_v15).length), globalState));
            let _1710_v22;
            _1710_v22 = _dafny.MultiSet.fromElements(true, (_1685_v1).f24, (_1685_v1).f24, (_1685_v1).f24, (_1685_v1).f24);
            let _1711_v23;
            let _nw275 = Array((new BigNumber(29)).toNumber());
            _nw275[(_dafny.ZERO).toNumber()] = (_1706_v18).Merge(_1706_v18);
            _nw275[(_dafny.ONE).toNumber()] = _1706_v18;
            _nw275[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1705_v17,(_1688_v4).update((_1685_v1).f24, (p0)[_module.__default.safeIndex(new BigNumber(277), new BigNumber((p0).length))]));
            _nw275[(new BigNumber(3)).toNumber()] = _1706_v18;
            _nw275[(new BigNumber(4)).toNumber()] = (_1706_v18).update(_1705_v17, _1688_v4);
            _nw275[(new BigNumber(5)).toNumber()] = (_1706_v18).Merge(_1706_v18);
            _nw275[(new BigNumber(6)).toNumber()] = (_1707_v19)[_module.__default.safeIndex(_1684_v0, new BigNumber((_1707_v19).length))];
            _nw275[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_module.D17.create_DC38(_1701_v16),_1688_v4);
            _nw275[(new BigNumber(8)).toNumber()] = _1706_v18;
            _nw275[(new BigNumber(9)).toNumber()] = (_1706_v18).Merge(_1706_v18);
            _nw275[(new BigNumber(10)).toNumber()] = _1706_v18;
            _nw275[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1705_v17,_1688_v4);
            _nw275[(new BigNumber(12)).toNumber()] = _1706_v18;
            _nw275[(new BigNumber(13)).toNumber()] = (((_1685_v1).f24) ? (_dafny.Map.Empty.slice().updateUnsafe(_1705_v17,_1688_v4)) : (_dafny.Map.Empty.slice().updateUnsafe(_1705_v17,(((_1708_v20).contains((p0)[_module.__default.safeIndex(new BigNumber(277), new BigNumber((p0).length))])) ? ((_1708_v20).get((p0)[_module.__default.safeIndex(new BigNumber(277), new BigNumber((p0).length))])) : (_1688_v4)))));
            _nw275[(new BigNumber(14)).toNumber()] = _1706_v18;
            _nw275[(new BigNumber(15)).toNumber()] = _1706_v18;
            _nw275[(new BigNumber(16)).toNumber()] = _1706_v18;
            _nw275[(new BigNumber(17)).toNumber()] = _1706_v18;
            _nw275[(new BigNumber(18)).toNumber()] = _1706_v18;
            _nw275[(new BigNumber(19)).toNumber()] = ((_dafny.Map.Empty.slice().updateUnsafe(_1705_v17,_1688_v4)).update(_1705_v17, (_1688_v4).update((_1685_v1).f24, (p0)[_module.__default.safeIndex(new BigNumber(277), new BigNumber((p0).length))]))).Merge(_1706_v18);
            _nw275[(new BigNumber(20)).toNumber()] = (_1706_v18).update(_module.D17.create_DC38(_1701_v16), (((_1709_v21).contains((_1685_v1).f24)) ? ((_1709_v21).get((_1685_v1).f24)) : (_1688_v4)));
            _nw275[(new BigNumber(21)).toNumber()] = _1706_v18;
            _nw275[(new BigNumber(22)).toNumber()] = _1706_v18;
            _nw275[(new BigNumber(23)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1705_v17,_dafny.Map.Empty.slice().updateUnsafe((_1685_v1).f24,(_dafny.ZERO).minus(new BigNumber((_1710_v22).cardinality()))));
            _nw275[(new BigNumber(24)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_1705_v17,_1688_v4)).Merge(_1706_v18);
            _nw275[(new BigNumber(25)).toNumber()] = (_1706_v18).update(_1705_v17, _1688_v4);
            _nw275[(new BigNumber(26)).toNumber()] = (_1706_v18).Merge(_1706_v18);
            _nw275[(new BigNumber(27)).toNumber()] = (_1706_v18).update(_module.D17.create_DC38((_1705_v17).dtor_cf51), _1688_v4);
            _nw275[(new BigNumber(28)).toNumber()] = _1706_v18;
            _1711_v23 = _nw275;
            let _index285 = _module.__default.safeIndex(new BigNumber(907), new BigNumber((_1711_v23).length));
            (_1711_v23)[_index285] = _1706_v18;
            let _index286 = _module.__default.safeIndex(new BigNumber(907), new BigNumber((_1711_v23).length));
            (_1711_v23)[_index286] = (_1707_v19)[_module.__default.safeIndex((p0)[_module.__default.safeIndex(new BigNumber(277), new BigNumber((p0).length))], new BigNumber((_1707_v19).length))];
            (globalState).f19 = _1699_v15;
          }
        }
      }
      let _hi12 = new BigNumber((_1699_v15).length);
      for (let _1712_i2 = ((p0)[_module.__default.safeIndex(new BigNumber(277), new BigNumber((p0).length))]).multipliedBy(new BigNumber(211)); _1712_i2.isLessThan(_hi12); _1712_i2 = _1712_i2.plus(_dafny.ONE)) {
        let _1713_v24;
        _1713_v24 = _dafny.Map.Empty.slice().updateUnsafe((_1685_v1).f24,_dafny.Seq.Create(_module.__default.abs(new BigNumber(318)), ((_1714_v15) => function (_1715_i3) {
          return new BigNumber((_1714_v15).length);
        })(_1699_v15)));
        let _1716_v25;
        _1716_v25 = _dafny.Map.Empty.slice().updateUnsafe((_1685_v1).f24,_1713_v24);
        let _1717_v26;
        _1717_v26 = _dafny.Set.fromElements((_1685_v1).f24);
        let _1718_v27;
        _1718_v27 = _dafny.Seq.of(new BigNumber((_1688_v4).length), new BigNumber((_1717_v26).length), (p0)[_module.__default.safeIndex(new BigNumber(277), new BigNumber((p0).length))]);
        let _1719_v28;
        let _nw276 = Array((new BigNumber(14)).toNumber());
        _nw276[(_dafny.ZERO).toNumber()] = _1713_v24;
        _nw276[(_dafny.ONE).toNumber()] = _1713_v24;
        _nw276[(new BigNumber(2)).toNumber()] = (((_1716_v25).contains((_1685_v1).f24)) ? ((_1716_v25).get((_1685_v1).f24)) : (_1713_v24));
        _nw276[(new BigNumber(3)).toNumber()] = _1713_v24;
        _nw276[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_1685_v1).f24,_1718_v27);
        _nw276[(new BigNumber(5)).toNumber()] = (_1713_v24).Merge(_1713_v24);
        _nw276[(new BigNumber(6)).toNumber()] = (_1713_v24).Merge(_dafny.Map.Empty.slice().updateUnsafe((_1685_v1).f24,_1718_v27));
        _nw276[(new BigNumber(7)).toNumber()] = (_1713_v24).Merge(_1713_v24);
        _nw276[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_1685_v1).f24,_1718_v27);
        _nw276[(new BigNumber(9)).toNumber()] = (_module.__default.fm50(_1712_i2, _1712_i2, globalState)).Merge(_1713_v24);
        _nw276[(new BigNumber(10)).toNumber()] = _1713_v24;
        _nw276[(new BigNumber(11)).toNumber()] = (_1713_v24).Merge(_dafny.Map.Empty.slice().updateUnsafe((_1685_v1).f24,_1718_v27));
        _nw276[(new BigNumber(12)).toNumber()] = _1713_v24;
        _nw276[(new BigNumber(13)).toNumber()] = _1713_v24;
        _1719_v28 = _nw276;
        let _index287 = _module.__default.safeIndex(new BigNumber(498), new BigNumber((_1719_v28).length));
        (_1719_v28)[_index287] = _1713_v24;
        let _1720_v29;
        _1720_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1699_v15,_1692_v8);
        let _index288 = _module.__default.safeIndex(new BigNumber(498), new BigNumber((_1719_v28).length));
        let _rhs287 = _1696_v12;
        let _rhs288 = (_1685_v1).f24;
        let _rhs289 = _dafny.Map.Empty.slice().updateUnsafe(!((_1685_v1).f24),_1718_v27);
        let _rhs290 = !((_1685_v1).f24);
        let _rhs291 = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(((_1720_v29).Merge((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_1699_v15, _module.__default.safeIndex((p0)[_module.__default.safeIndex(new BigNumber(277), new BigNumber((p0).length))], new BigNumber((_1699_v15).length)), new _dafny.CodePoint('i'.codePointAt(0))),_1692_v8)).Merge(_1720_v29))).length)));
        let _lhs271 = globalState;
        let _lhs272 = _1719_v28;
        let _lhs273 = _module.__default.safeIndex(new BigNumber(498), new BigNumber((_1719_v28).length));
        let _lhs274 = globalState;
        let _lhs275 = globalState;
        _1696_v12 = _rhs287;
        _lhs271.f2 = _rhs288;
        _lhs272[_lhs273] = _rhs289;
        _lhs274.f21 = _rhs290;
        _lhs275.f7 = _rhs291;
        let _1721_v30;
        _1721_v30 = _dafny.Seq.of(_1694_v10);
        let _1722_v31;
        let _nw277 = new _module.C1();
        _nw277.__ctor((((_1685_v1).f24) ? (_1712_i2) : (new BigNumber((_1699_v15).length))), _dafny.Seq.Concat(_module.__default.fm51(new _dafny.CodePoint('i'.codePointAt(0)), (p0)[_module.__default.safeIndex(new BigNumber(277), new BigNumber((p0).length))], globalState), _1721_v30));
        _1722_v31 = _nw277;
        let _1723_v32;
        _1723_v32 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm34(globalState),new BigNumber((_1687_v3).length));
        _1723_v32 = (_1723_v32).update(_1696_v12, _1722_v31.f29);
        let _index289 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_1695_v11).length));
        (_1695_v11)[_index289] = (_1685_v1).f24;
        let _1724_v33;
        _1724_v33 = _dafny.Seq.of(p0);
        let _1725_v34;
        _1725_v34 = _module.D6.create_DC12(_dafny.Seq.update(_1724_v33, _module.__default.safeIndex((p0)[_module.__default.safeIndex(new BigNumber(277), new BigNumber((p0).length))], new BigNumber((_1724_v33).length)), p0));
        let _1726_v35;
        _1726_v35 = _module.D6.create_DC14(_1725_v34);
        let _1727_v36;
        _1727_v36 = _module.D6.create_DC14(_1725_v34);
        let _1728_v37;
        _1728_v37 = _module.D6.create_DC14(_1726_v35);
        let _1729_v38;
        _1729_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1722_v31.f29,p0);
        let _index290 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_1695_v11).length));
        let _rhs292 = ((_1722_v31).fm29(globalState)).plus((new BigNumber((_1694_v10).length)).plus(_1712_i2));
        let _rhs293 = (_1729_v38).contains((((_1685_v1).f24) ? ((_dafny.ZERO).minus(_1722_v31.f29)) : (_1684_v0)));
        let _rhs294 = _module.D6.create_DC14(_1725_v34);
        let _rhs295 = _1699_v15;
        let _rhs296 = _1722_v31.f29;
        let _lhs276 = _1722_v31;
        let _lhs277 = _1695_v11;
        let _lhs278 = _module.__default.safeIndex(new BigNumber(760), new BigNumber((_1695_v11).length));
        let _lhs279 = globalState;
        let _lhs280 = globalState;
        _lhs276.f29 = _rhs292;
        _lhs277[_lhs278] = _rhs293;
        _1728_v37 = _rhs294;
        _lhs279.f20 = _rhs295;
        _lhs280.f7 = _rhs296;
      }
      return;
    }
    m1(p0, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = false;
      let _1730_v0;
      _1730_v0 = false;
      let _1731_v1;
      _1731_v1 = _dafny.Map.Empty.slice().updateUnsafe(_1730_v0,_1730_v0);
      _1731_v1 = (_1731_v1).update(_1730_v0, _module.__default.fm2((_1731_v1).update(_1730_v0, !(_1730_v0)), globalState));
      let _1732_i0;
      _1732_i0 = _dafny.ZERO;
      L10: {
        while (_1730_v0) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1732_i0)) {
              break L10;
            }
            _1732_i0 = (_1732_i0).plus(_dafny.ONE);
            let _1733_v2;
            _1733_v2 = new _dafny.CodePoint('t'.codePointAt(0));
            let _rhs297 = _1730_v0;
            let _rhs298 = new _dafny.CodePoint('k'.codePointAt(0));
            let _lhs281 = globalState;
            _lhs281.f0 = _rhs297;
            _1733_v2 = _rhs298;
            if (true) {
              let _1734_v3;
              _1734_v3 = _dafny.MultiSet.fromElements(new BigNumber((_1731_v1).length));
              (globalState).f2 = ((_1734_v3).IsDisjointFrom(_1734_v3)) || (_1730_v0);
              let _1735_v5;
              _1735_v5 = _dafny.Set.fromElements(new BigNumber((function () {
                let _coll67 = new _dafny.Map();
                for (const _compr_67 of (_1734_v3).Elements) {
                  let _1736_v4 = _compr_67;
                  if ((_1734_v3).contains(_1736_v4)) {
                    _coll67.push([(_1736_v4).plus(p0),_dafny.Seq.Create(_module.__default.abs(new BigNumber(351)), ((_1737_p0) => function (_1738_i1) {
                      return _1737_p0;
                    })(p0))]);
                  }
                }
                return _coll67;
              }()).length), p0);
              let _1739_v6;
              _1739_v6 = _dafny.MultiSet.fromElements(_1730_v0, (_1735_v5).IsSubsetOf(_1735_v5));
              let _1740_v7;
              _1740_v7 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1739_v6).update(_1730_v0, _module.__default.abs(p0)));
              let _1741_v8;
              let _nw278 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
              _1741_v8 = _nw278;
              let _1742_v9;
              _1742_v9 = _dafny.Seq.of(_1730_v0);
              let _1743_v10;
              _1743_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1741_v8,_1742_v9);
              _1739_v6 = ((((_1740_v7).contains(p0)) ? ((_1740_v7).get(p0)) : (_1739_v6))).Union(_dafny.MultiSet.FromArray((((_1743_v10).contains(_1741_v8)) ? ((_1743_v10).get(_1741_v8)) : (_1742_v9))));
              let _1744_v11;
              let _nw279 = new _module.C0();
              _nw279.__ctor(_1730_v0);
              _1744_v11 = _nw279;
              _1744_v11 = _1744_v11;
              let _1745_v12;
              let _nw280 = Array((_dafny.ONE).toNumber());
              _1745_v12 = _nw280;
              let _1746_v13;
              _1746_v13 = _dafny.Seq.of(p0, new BigNumber((_dafny.Seq.UnicodeFromString("eoog")).length));
              let _1747_v14;
              _1747_v14 = _dafny.Seq.of(_1746_v13);
              let _1748_v15;
              _1748_v15 = _dafny.Map.Empty.slice().updateUnsafe(_1742_v9,(_1747_v14)[_module.__default.safeIndex(p0, new BigNumber((_1747_v14).length))]);
              let _1749_v16;
              _1749_v16 = _module.D21.create_DC52(_1748_v15);
              let _1750_v17;
              let _nw281 = new _module.C2();
              _nw281.__ctor(new BigNumber(((_1749_v16).dtor_cf71).length));
              _1750_v17 = _nw281;
              let _index291 = _module.__default.safeIndex(new BigNumber(504), new BigNumber((_1745_v12).length));
              (_1745_v12)[_index291] = _1750_v17;
              let _index292 = _module.__default.safeIndex(new BigNumber(504), new BigNumber((_1745_v12).length));
              (_1745_v12)[_index292] = _1750_v17;
              (globalState).f19 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-484)), ((_1751_v2) => function (_1752_i2) {
                return _1751_v2;
              })(_1733_v2));
            } else {
              let _1753_v18;
              let _nw282 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
              _1753_v18 = _nw282;
              r0 = _1753_v18;
              let _1754_v19;
              _1754_v19 = _dafny.Set.fromElements(p0, p0);
              let _1755_v20;
              _1755_v20 = _module.D14.create_DC28(_1754_v19);
              let _1756_v21;
              let _nw283 = Array((new BigNumber(25)).toNumber());
              _nw283[(_dafny.ZERO).toNumber()] = _1730_v0;
              _nw283[(_dafny.ONE).toNumber()] = _1730_v0;
              _nw283[(new BigNumber(2)).toNumber()] = !(_1730_v0);
              _nw283[(new BigNumber(3)).toNumber()] = _1730_v0;
              _nw283[(new BigNumber(4)).toNumber()] = !(_1730_v0);
              _nw283[(new BigNumber(5)).toNumber()] = (_1754_v19).IsProperSubsetOf((_1755_v20).dtor_cf37);
              _nw283[(new BigNumber(6)).toNumber()] = (p0).isLessThan(p0);
              _nw283[(new BigNumber(7)).toNumber()] = _1730_v0;
              _nw283[(new BigNumber(8)).toNumber()] = _1730_v0;
              _nw283[(new BigNumber(9)).toNumber()] = (!(true)) === (false);
              _nw283[(new BigNumber(10)).toNumber()] = true;
              _nw283[(new BigNumber(11)).toNumber()] = !(true);
              _nw283[(new BigNumber(12)).toNumber()] = _1730_v0;
              _nw283[(new BigNumber(13)).toNumber()] = _1730_v0;
              _nw283[(new BigNumber(14)).toNumber()] = !(_1730_v0) || (_1730_v0);
              _nw283[(new BigNumber(15)).toNumber()] = _1730_v0;
              _nw283[(new BigNumber(16)).toNumber()] = _1730_v0;
              _nw283[(new BigNumber(17)).toNumber()] = _1730_v0;
              _nw283[(new BigNumber(18)).toNumber()] = _1730_v0;
              _nw283[(new BigNumber(19)).toNumber()] = false;
              _nw283[(new BigNumber(20)).toNumber()] = _1730_v0;
              _nw283[(new BigNumber(21)).toNumber()] = _1730_v0;
              _nw283[(new BigNumber(22)).toNumber()] = _1730_v0;
              _nw283[(new BigNumber(23)).toNumber()] = _1730_v0;
              _nw283[(new BigNumber(24)).toNumber()] = _1730_v0;
              _1756_v21 = _nw283;
              let _index293 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_1756_v21).length));
              (_1756_v21)[_index293] = _1730_v0;
              let _index294 = _module.__default.safeIndex(new BigNumber(680), new BigNumber((_1756_v21).length));
              (_1756_v21)[_index294] = _module.__default.fm2(_1731_v1, globalState);
              let _1757_v22;
              let _nw284 = new _module.C3();
              _nw284.__ctor();
              _1757_v22 = _nw284;
              let _1758_v23;
              _1758_v23 = _dafny.Map.Empty.slice().updateUnsafe(_1730_v0,_1757_v22);
              let _1759_v24;
              _1759_v24 = _dafny.Seq.of(_1758_v23);
              let _1760_v25;
              _1760_v25 = _dafny.Seq.of((_1759_v24)[_module.__default.safeIndex(p0, new BigNumber((_1759_v24).length))]);
              let _1761_v26;
              _1761_v26 = _dafny.Seq.of(false);
              let _1762_v27;
              _1762_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1761_v26,_dafny.Seq.UnicodeFromString("ulcd"));
              let _1763_v28;
              _1763_v28 = _dafny.Seq.of(_module.__default.fm27((_1756_v21)[_module.__default.safeIndex(new BigNumber(680), new BigNumber((_1756_v21).length))], (_1756_v21)[_module.__default.safeIndex(new BigNumber(680), new BigNumber((_1756_v21).length))], globalState));
              let _1764_v29;
              _1764_v29 = _dafny.Set.fromElements(_1763_v28, _1763_v28, _1763_v28);
              _module.__default.m0((p0).isEqualTo(p0), _module.__default.safeDivisionInt(new BigNumber((_1760_v25).length), p0), _1762_v27, new BigNumber(((_1764_v29).Difference(_1764_v29)).length), globalState);
              let _1765_v33;
              _1765_v33 = _dafny.Seq.of(p0);
              let _1766_v34;
              _1766_v34 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(225));
              let _1767_v35;
              _1767_v35 = _dafny.Seq.of(function () {
                let _coll68 = new _dafny.Map();
                for (const _compr_68 of (_1765_v33).Elements) {
                  let _1768_v32 = _compr_68;
                  if (_dafny.Seq.contains(_1765_v33, _1768_v32)) {
                    _coll68.push([_module.__default.safeModuloInt(_1768_v32, p0),new BigNumber(251)]);
                  }
                }
                return _coll68;
              }(), _1766_v34);
              let _1769_v36;
              _1769_v36 = _dafny.Map.Empty.slice().updateUnsafe(function () {
                let _coll69 = new _dafny.Map();
                for (const _compr_69 of (function () {
                  let _coll70 = new _dafny.Map();
                  for (const _compr_70 of (_1767_v35).Elements) {
                    let _1770_v31 = _compr_70;
                    if (_dafny.Seq.contains(_1767_v35, _1770_v31)) {
                      _coll70.push([_1770_v31,p0]);
                    }
                  }
                  return _coll70;
                }()).Keys.Elements) {
                  let _1771_v30 = _compr_69;
                  if ((function () {
                    let _coll71 = new _dafny.Map();
                    for (const _compr_71 of (_1767_v35).Elements) {
                      let _1770_v31 = _compr_71;
                      if (_dafny.Seq.contains(_1767_v35, _1770_v31)) {
                        _coll71.push([_1770_v31,p0]);
                      }
                    }
                    return _coll71;
                  }()).contains(_1771_v30)) {
                    _coll69.push([_1771_v30,_1730_v0]);
                  }
                }
                return _coll69;
              }(),true);
              let _1772_v38;
              _1772_v38 = _dafny.Map.Empty.slice().updateUnsafe(function () {
                let _coll72 = new _dafny.Map();
                for (const _compr_72 of _dafny.IntegerRange(new BigNumber(83), new BigNumber(288))) {
                  let _1773_v37 = _compr_72;
                  if (((new BigNumber(83)).isLessThanOrEqualTo(_1773_v37)) && ((_1773_v37).isLessThan(new BigNumber(288)))) {
                    _coll72.push([(_1773_v37).plus(p0),p0]);
                  }
                }
                return _coll72;
              }(),(_1756_v21)[_module.__default.safeIndex(new BigNumber(680), new BigNumber((_1756_v21).length))]);
              _1769_v36 = (_1769_v36).update(_1772_v38, _1730_v0);
              let _1774_v39;
              let _nw285 = new _module.C2();
              _nw285.__ctor(p0);
              _1774_v39 = _nw285;
            }
            let _1775_v40;
            _1775_v40 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_1730_v0),_dafny.Seq.UnicodeFromString("l"));
            let _1776_v41;
            _1776_v41 = _dafny.Seq.of(_1775_v40);
            _module.__default.m0(!(_1730_v0) || (_1730_v0), p0, (_1776_v41)[_module.__default.safeIndex(p0, new BigNumber((_1776_v41).length))], (p0).minus(p0), globalState);
            let _1777_v42;
            _1777_v42 = _dafny.Seq.UnicodeFromString("t");
            let _1778_v43;
            _1778_v43 = _dafny.Map.Empty.slice().updateUnsafe(!(_1730_v0),_1777_v42);
            (globalState).f20 = _dafny.Seq.Concat(_1777_v42, (((_1778_v43).contains(_1730_v0)) ? ((_1778_v43).get(_1730_v0)) : (_1777_v42)));
          }
        }
      }
      let _1779_v44;
      _1779_v44 = _dafny.Seq.UnicodeFromString("aapa");
      let _1780_v45;
      _1780_v45 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1730_v0);
      let _1781_v46;
      _1781_v46 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1780_v45).length),p0);
      let _hi13 = p0;
      for (let _1782_i3 = _module.__default.safeDivisionInt(new BigNumber((_1779_v44).length), _module.__default.fm0(p0, new BigNumber((_1781_v46).length), p0, _1730_v0, globalState)); _1782_i3.isLessThan(_hi13); _1782_i3 = _1782_i3.plus(_dafny.ONE)) {
        let _1783_v47;
        let _nw286 = new _module.C3();
        _nw286.__ctor();
        _1783_v47 = _nw286;
        _1783_v47 = _1783_v47;
        let _1784_v48;
        let _init51 = ((_1785_v44) => function (_1786_i4) {
          return _dafny.Seq.contains(_1785_v44, new _dafny.CodePoint('n'.codePointAt(0)));
        })(_1779_v44);
        let _nw287 = Array((new BigNumber(18)).toNumber());
        for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw287.length); _i0_51++) {
          _nw287[_i0_51] = _init51(new BigNumber(_i0_51));
        }
        _1784_v48 = _nw287;
        let _index295 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_1784_v48).length));
        (_1784_v48)[_index295] = _1730_v0;
        let _1787_v49;
        let _nw288 = Array((new BigNumber(19)).toNumber());
        _1787_v49 = _nw288;
        let _1788_v50;
        _1788_v50 = _dafny.Map.Empty.slice().updateUnsafe(_1787_v49,_1730_v0);
        let _index296 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_1784_v48).length));
        let _rhs299 = _1782_i3;
        let _rhs300 = !(_1788_v50).equals((_1788_v50).Merge((_module.D22.create_DC56(_dafny.Map.Empty.slice().updateUnsafe(_1787_v49,false))).dtor_cf80));
        let _lhs282 = globalState;
        let _lhs283 = _1784_v48;
        let _lhs284 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_1784_v48).length));
        _lhs282.f16 = _rhs299;
        _lhs283[_lhs284] = _rhs300;
        let _1789_v51;
        _1789_v51 = _dafny.Map.Empty.slice().updateUnsafe(_1730_v0,_1782_i3);
        let _1790_v52;
        _1790_v52 = _dafny.MultiSet.fromElements(_1782_i3);
        let _1791_v53;
        _1791_v53 = _dafny.Set.fromElements(p0, p0, p0, p0);
        let _1792_v54;
        _1792_v54 = _dafny.Seq.of((_1784_v48)[_module.__default.safeIndex(new BigNumber(567), new BigNumber((_1784_v48).length))], (_1784_v48)[_module.__default.safeIndex(new BigNumber(567), new BigNumber((_1784_v48).length))]);
        let _1793_v55;
        _1793_v55 = _module.D4.create_DC9(_1782_i3);
        let _1794_v56;
        _1794_v56 = _dafny.Seq.of(p0, new BigNumber(515));
        let _1795_v57;
        _1795_v57 = _module.D4.create_DC10(new BigNumber((_1792_v54).length), new BigNumber((_1731_v1).length), _1730_v0, (_1784_v48)[_module.__default.safeIndex(new BigNumber(567), new BigNumber((_1784_v48).length))], _1782_i3);
        let _1796_v58;
        _1796_v58 = new _dafny.CodePoint('k'.codePointAt(0));
        let _1797_v59;
        _1797_v59 = _dafny.MultiSet.fromElements(_1796_v58, _1796_v58, _1796_v58, _1796_v58);
        let _1798_v60;
        _1798_v60 = _dafny.Set.fromElements(_1730_v0, _1730_v0, _1730_v0);
        let _1799_v61;
        let _nw289 = Array((new BigNumber(24)).toNumber());
        _nw289[(_dafny.ZERO).toNumber()] = _1782_i3;
        _nw289[(_dafny.ONE).toNumber()] = _1782_i3;
        _nw289[(new BigNumber(2)).toNumber()] = new BigNumber((_1781_v46).length);
        _nw289[(new BigNumber(3)).toNumber()] = (new BigNumber((_1789_v51).length)).plus(new BigNumber((_1790_v52).cardinality()));
        _nw289[(new BigNumber(4)).toNumber()] = p0;
        _nw289[(new BigNumber(5)).toNumber()] = _module.__default.safeModuloInt(p0, new BigNumber((_1791_v53).length));
        _nw289[(new BigNumber(6)).toNumber()] = (p0).multipliedBy(new BigNumber((_1792_v54).length));
        _nw289[(new BigNumber(7)).toNumber()] = new BigNumber(848);
        _nw289[(new BigNumber(8)).toNumber()] = p0;
        _nw289[(new BigNumber(9)).toNumber()] = (_1793_v55).dtor_cf11;
        _nw289[(new BigNumber(10)).toNumber()] = (new BigNumber((_1781_v46).length)).minus(_1782_i3);
        _nw289[(new BigNumber(11)).toNumber()] = (_1794_v56)[_module.__default.safeIndex(_1782_i3, new BigNumber((_1794_v56).length))];
        _nw289[(new BigNumber(12)).toNumber()] = (_1795_v57).dtor_cf12;
        _nw289[(new BigNumber(13)).toNumber()] = p0;
        _nw289[(new BigNumber(14)).toNumber()] = (p0).minus(new BigNumber((_1797_v59).cardinality()));
        _nw289[(new BigNumber(15)).toNumber()] = (((_1790_v52).contains(_1782_i3)) ? ((_1790_v52).get(_1782_i3)) : (p0));
        _nw289[(new BigNumber(16)).toNumber()] = _1782_i3;
        _nw289[(new BigNumber(17)).toNumber()] = _1782_i3;
        _nw289[(new BigNumber(18)).toNumber()] = p0;
        _nw289[(new BigNumber(19)).toNumber()] = (((_1781_v46).contains(new BigNumber((_1798_v60).length))) ? ((_1781_v46).get(new BigNumber((_1798_v60).length))) : (p0));
        _nw289[(new BigNumber(20)).toNumber()] = _1782_i3;
        _nw289[(new BigNumber(21)).toNumber()] = _1782_i3;
        _nw289[(new BigNumber(22)).toNumber()] = _module.__default.fm0(p0, new BigNumber(827), p0, !((_1784_v48)[_module.__default.safeIndex(new BigNumber(567), new BigNumber((_1784_v48).length))]), globalState);
        _nw289[(new BigNumber(23)).toNumber()] = new BigNumber(-744);
        _1799_v61 = _nw289;
        let _index297 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((_1799_v61).length));
        (_1799_v61)[_index297] = _1782_i3;
        let _index298 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((_1799_v61).length));
        (_1799_v61)[_index298] = new BigNumber(-85);
        _1789_v51 = (_1789_v51).update(!(true), (_1799_v61)[_module.__default.safeIndex(new BigNumber(412), new BigNumber((_1799_v61).length))]);
      }
      let _1800_v62;
      let _nw290 = new _module.C0();
      _nw290.__ctor(_1730_v0);
      _1800_v62 = _nw290;
      (globalState).f7 = p0;
      let _1801_i5;
      _1801_i5 = _dafny.ZERO;
      L11: {
        while (_1800_v62.f27) {
          C11: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1801_i5)) {
              break L11;
            }
            _1801_i5 = (_1801_i5).plus(_dafny.ONE);
            (globalState).f7 = p0;
            let _1802_v63;
            _1802_v63 = _dafny.MultiSet.fromElements(_1800_v62.f27);
            _1802_v63 = _module.__default.fm52(_1800_v62.f27, p0, globalState);
            let _1803_v64;
            let _init52 = ((_1804_v44) => function (_1805_i6) {
              return _dafny.Set.fromElements(_1804_v44, _1804_v44);
            })(_1779_v44);
            let _nw291 = Array((new BigNumber(29)).toNumber());
            for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw291.length); _i0_52++) {
              _nw291[_i0_52] = _init52(new BigNumber(_i0_52));
            }
            _1803_v64 = _nw291;
            let _1806_v65;
            _1806_v65 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("ycoxy"));
            let _index299 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_1803_v64).length));
            (_1803_v64)[_index299] = (_1806_v65).Intersect(function () {
              let _coll73 = new _dafny.Set();
              for (const _compr_73 of (_1806_v65).Elements) {
                let _1807_v66 = _compr_73;
                if ((_1806_v65).contains(_1807_v66)) {
                  _coll73.add(_1807_v66);
                }
              }
              return _coll73;
            }());
            let _index300 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_1803_v64).length));
            (_1803_v64)[_index300] = _1806_v65;
            let _1808_v67;
            _1808_v67 = new _dafny.CodePoint('b'.codePointAt(0));
            _1808_v67 = _1808_v67;
          }
        }
      }
      let _1809_v68;
      let _nw292 = Array((_dafny.ONE).toNumber());
      _nw292[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(p0);
      _1809_v68 = _nw292;
      r0 = _1809_v68;
      r1 = ((_1730_v0) ? (_1800_v62.f27) : (_1800_v62.f27));
      return [r0, r1];
    }
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
      this._f23 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    __ctor(f23) {
      let _this = this;
      (_this)._f23 = f23;
      return;
    }
    fm7(p0, p1, p2, globalState) {
      let _this = this;
      return false;
    };
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe(((false) ? (!(true)) : (true)),_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("uqyvs"), _dafny.Seq.UnicodeFromString("egwtgbjvl")));
    };
    fm6(p0, globalState) {
      let _this = this;
      return !(!(!(true)) || ((_dafny.Set.fromElements(new BigNumber(-747))).IsDisjointFrom(_dafny.Set.fromElements(new BigNumber(158), new BigNumber(-959)))));
    };
    m2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = [];
      let _1810_v0;
      let _nw293 = new _module.C6();
      _nw293.__ctor();
      _1810_v0 = _nw293;
      let _hi14 = p3;
      for (let _1811_i0 = p3; _1811_i0.isLessThan(_hi14); _1811_i0 = _1811_i0.plus(_dafny.ONE)) {
        (globalState).f16 = (new BigNumber(914)).minus((_1811_i0).plus(new BigNumber((_dafny.Seq.of(p3)).length)));
        let _1812_v1;
        let _init53 = function (_1813_i1) {
          return _dafny.Seq.UnicodeFromString("rhc");
        };
        let _nw294 = Array((new BigNumber(5)).toNumber());
        for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw294.length); _i0_53++) {
          _nw294[_i0_53] = _init53(new BigNumber(_i0_53));
        }
        _1812_v1 = _nw294;
        _1812_v1 = (((_1811_i0).isLessThan(_1811_i0)) ? (_1812_v1) : (_1812_v1));
        let _1814_v2;
        let _nw295 = new _module.C5();
        _nw295.__ctor(p1);
        _1814_v2 = _nw295;
        _1814_v2 = _1814_v2;
        let _1815_v4;
        let _init54 = ((_1816_p1) => function (_1817_i2) {
          return _dafny.Seq.of(function () {
            let _coll74 = new _dafny.Map();
            for (const _compr_74 of (_dafny.Set.fromElements(_dafny.Set.fromElements(_1816_p1), _dafny.Set.fromElements(_1816_p1))).Elements) {
              let _1818_v3 = _compr_74;
              if ((_dafny.Set.fromElements(_dafny.Set.fromElements(_1816_p1), _dafny.Set.fromElements(_1816_p1))).contains(_1818_v3)) {
                _coll74.push([_1818_v3,_1816_p1]);
              }
            }
            return _coll74;
          }(), _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_1816_p1),_1816_p1));
        })(p1);
        let _nw296 = Array((new BigNumber(27)).toNumber());
        for (let _i0_54 = 0; _i0_54 < new BigNumber(_nw296.length); _i0_54++) {
          _nw296[_i0_54] = _init54(new BigNumber(_i0_54));
        }
        _1815_v4 = _nw296;
        let _1819_v5;
        _1819_v5 = _dafny.Set.fromElements(false, p1);
        let _1820_v6;
        _1820_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1819_v5,p1);
        let _1821_v7;
        _1821_v7 = _dafny.Seq.of(_1820_v6);
        let _1822_v8;
        _1822_v8 = _1819_v5;
        let _index301 = _module.__default.safeIndex(new BigNumber(54), new BigNumber((_1815_v4).length));
        (_1815_v4)[_index301] = _dafny.Seq.Concat(_1821_v7, _dafny.Seq.of((_dafny.Map.Empty.slice().updateUnsafe(_1822_v8,p1)).update(_1822_v8, false)));
        let _index302 = _module.__default.safeIndex(new BigNumber(54), new BigNumber((_1815_v4).length));
        (_1815_v4)[_index302] = _1821_v7;
      }
      let _1823_v9;
      let _nw297 = new _module.C2();
      _nw297.__ctor((_dafny.ZERO).minus(p0));
      _1823_v9 = _nw297;
      let _1824_v10;
      _1824_v10 = _dafny.Set.fromElements(_1823_v9, _1823_v9, _1823_v9);
      if ((_1824_v10).IsSubsetOf(_dafny.Set.fromElements(_1823_v9, _1823_v9))) {
        (globalState).f16 = p3;
        let _1825_v11;
        _1825_v11 = _dafny.Seq.of((_1823_v9).f28);
        let _1826_v12;
        let _nw298 = new _module.C0();
        _nw298.__ctor(!((_1823_v9).f28).isEqualTo(_module.__default.fm0(p0, new BigNumber((_1825_v11).length), _module.__default.fm0(p3, p3, p0, p1, globalState), p1, globalState)));
        _1826_v12 = _nw298;
        let _1827_v13;
        let _init55 = ((_1828_v12, _1829_p1) => function (_1830_i3) {
          return (_dafny.MultiSet.FromArray(_dafny.Seq.of(_1828_v12.f27))).IsDisjointFrom(_dafny.MultiSet.fromElements(_1829_p1, !(_1828_v12.f27)));
        })(_1826_v12, p1);
        let _nw299 = Array((new BigNumber(20)).toNumber());
        for (let _i0_55 = 0; _i0_55 < new BigNumber(_nw299.length); _i0_55++) {
          _nw299[_i0_55] = _init55(new BigNumber(_i0_55));
        }
        _1827_v13 = _nw299;
        let _1831_v14;
        let _nw300 = new _module.C6();
        _nw300.__ctor();
        _1831_v14 = _nw300;
        let _1832_v15;
        _1832_v15 = _dafny.Seq.of(_1826_v12.f27);
        let _1833_v16;
        _1833_v16 = _dafny.MultiSet.fromElements(((_1826_v12.f27) ? (_1832_v15) : (_1832_v15)));
        let _1834_v17;
        _1834_v17 = _dafny.Seq.UnicodeFromString("mdrtqm");
        let _rhs301 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_1823_v9).f28));
        let _rhs302 = new BigNumber((_1834_v17).length);
        let _rhs303 = _1827_v13;
        let _rhs304 = _1831_v14;
        let _rhs305 = _1833_v16;
        let _lhs285 = globalState;
        r1 = _rhs301;
        _lhs285.f7 = _rhs302;
        _1827_v13 = _rhs303;
        _1831_v14 = _rhs304;
        _1833_v16 = _rhs305;
        let _1835_v18;
        _1835_v18 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
        let _1836_v19;
        let _nw301 = new _module.C1();
        _nw301.__ctor(((_1823_v9).f28).plus(_module.__default.fm0(p3, (_1823_v9).f28, (_1823_v9).f28, true, globalState)), _dafny.Seq.of(_1835_v18, _1835_v18, _1835_v18));
        _1836_v19 = _nw301;
        (globalState).f7 = (_dafny.ZERO).minus((_1823_v9).fm17(false, _1825_v11, p3, _1832_v15, globalState));
      } else {
        let _1837_v20;
        _1837_v20 = _module.D17.create_DC40();
        let _source23 = _1837_v20;
        if (_source23.is_DC39) {
          let _1838___mcc_h0 = (_source23).cf52;
          let _1839___mcc_h1 = (_source23).cf53;
          let _1840___mcc_h2 = (_source23).cf54;
          let _1841___mcc_h3 = (_source23).cf55;
          let _1842___mcc_h4 = (_source23).cf56;
          let _1843_cf56 = _1842___mcc_h4;
          let _1844_cf55 = _1841___mcc_h3;
          let _1845_cf54 = _1840___mcc_h2;
          let _1846_cf53 = _1839___mcc_h1;
          let _1847_cf52 = _1838___mcc_h0;
          _1847_cf52 = new BigNumber(-36);
          let _1848_v21;
          _1848_v21 = _dafny.Map.Empty.slice().updateUnsafe(p1,false);
          let _1849_v22;
          _1849_v22 = _dafny.Seq.of(_1848_v21, _1848_v21);
          (globalState).f13 = _module.__default.fm2((_1849_v22)[_module.__default.safeIndex(_1847_cf52, new BigNumber((_1849_v22).length))], globalState);
          let _1850_v23;
          _1850_v23 = _dafny.MultiSet.fromElements(!(p1), p1);
          let _1851_v24;
          let _nw302 = new _module.C2();
          _nw302.__ctor(new BigNumber((_1850_v23).cardinality()));
          _1851_v24 = _nw302;
          let _1852_v25;
          _1852_v25 = _dafny.Seq.of(_1847_cf52, new BigNumber(955), (_1823_v9).f28, new BigNumber((_1850_v23).cardinality()));
          (globalState).f0 = _dafny.areEqual(_1852_v25, _1852_v25);
        } else if (_source23.is_DC40) {
          let _1853_v26;
          _1853_v26 = _dafny.Seq.of(p1, p1);
          let _1854_v27;
          _1854_v27 = _1853_v26;
          let _index303 = _module.__default.safeIndex(new BigNumber(688), new BigNumber((p2).length));
          (p2)[_index303] = (_1823_v9).fm17(p1, _dafny.Seq.of((_1823_v9).f28), (_1823_v9).f28, _1854_v27, globalState);
          let _1855_v28;
          _1855_v28 = _dafny.Seq.UnicodeFromString("ug");
          let _index304 = _module.__default.safeIndex(new BigNumber(688), new BigNumber((p2).length));
          (p2)[_index304] = new BigNumber((_1855_v28).length);
          (globalState).f16 = p3;
          let _1856_v29;
          let _nw303 = new _module.C5();
          _nw303.__ctor((_this).fm7(p1, (_dafny.ZERO).minus((p2)[_module.__default.safeIndex(new BigNumber(688), new BigNumber((p2).length))]), p1, globalState));
          _1856_v29 = _nw303;
          let _1857_v30;
          _1857_v30 = new _dafny.CodePoint('u'.codePointAt(0));
          let _1858_v31;
          let _nw304 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Map.Empty);
          _1858_v31 = _nw304;
          let _1859_v32;
          _1859_v32 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_1856_v29).f24);
          let _1860_v33;
          _1860_v33 = _dafny.Map.Empty.slice().updateUnsafe((_1856_v29).f24,_1859_v32);
          let _index305 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_1858_v31).length));
          (_1858_v31)[_index305] = (_module.D23.create_DC60(_1860_v33)).dtor_cf85;
          let _1861_v34;
          _1861_v34 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("uvqcohcw"));
          let _1862_v35;
          _1862_v35 = _module.D0.create_DC0(p1);
          let _1863_v36;
          _1863_v36 = _dafny.Set.fromElements(_1857_v30);
          let _1864_v37;
          _1864_v37 = _dafny.MultiSet.fromElements(new BigNumber(648));
          let _1865_v38;
          _1865_v38 = _dafny.Map.Empty.slice().updateUnsafe((_1823_v9).f28,p3);
          let _1866_v39;
          _1866_v39 = _dafny.Map.Empty.slice().updateUnsafe(_1865_v38,p1);
          let _1867_v40;
          _1867_v40 = _dafny.Map.Empty.slice().updateUnsafe((_1856_v29).f24,new BigNumber((_module.__default.fm54(globalState)).length));
          let _1868_v41;
          _1868_v41 = _dafny.MultiSet.fromElements(true);
          let _1869_v42;
          _1869_v42 = _dafny.Seq.of(new BigNumber((_1868_v41).cardinality()));
          let _1870_v43;
          _1870_v43 = _module.D15.create_DC33(_1869_v42);
          let _1871_v44;
          let _nw305 = Array((new BigNumber(29)).toNumber());
          _nw305[(_dafny.ZERO).toNumber()] = (_1856_v29).f24;
          _nw305[(_dafny.ONE).toNumber()] = (_1861_v34).IsDisjointFrom(_module.__default.fm53(p1, globalState));
          _nw305[(new BigNumber(2)).toNumber()] = !((_1856_v29).f24);
          _nw305[(new BigNumber(3)).toNumber()] = p1;
          _nw305[(new BigNumber(4)).toNumber()] = (_1856_v29).f24;
          _nw305[(new BigNumber(5)).toNumber()] = p1;
          _nw305[(new BigNumber(6)).toNumber()] = !(p1);
          _nw305[(new BigNumber(7)).toNumber()] = false;
          _nw305[(new BigNumber(8)).toNumber()] = (_1856_v29).f24;
          _nw305[(new BigNumber(9)).toNumber()] = p1;
          _nw305[(new BigNumber(10)).toNumber()] = !((_1862_v35).dtor_cf0);
          _nw305[(new BigNumber(11)).toNumber()] = p1;
          _nw305[(new BigNumber(12)).toNumber()] = (_1863_v36).IsSubsetOf(_1863_v36);
          _nw305[(new BigNumber(13)).toNumber()] = (_1856_v29).f24;
          _nw305[(new BigNumber(14)).toNumber()] = (_1856_v29).fm6(new BigNumber((_1853_v26).length), globalState);
          _nw305[(new BigNumber(15)).toNumber()] = !((_1864_v37).IsDisjointFrom(_1864_v37));
          _nw305[(new BigNumber(16)).toNumber()] = (_1856_v29).f24;
          _nw305[(new BigNumber(17)).toNumber()] = false;
          _nw305[(new BigNumber(18)).toNumber()] = (_1856_v29).f24;
          _nw305[(new BigNumber(19)).toNumber()] = p1;
          _nw305[(new BigNumber(20)).toNumber()] = (_1856_v29).f24;
          _nw305[(new BigNumber(21)).toNumber()] = (_1856_v29).f24;
          _nw305[(new BigNumber(22)).toNumber()] = (_1866_v39).contains(_1865_v38);
          _nw305[(new BigNumber(23)).toNumber()] = p1;
          _nw305[(new BigNumber(24)).toNumber()] = !(p1);
          _nw305[(new BigNumber(25)).toNumber()] = (_1856_v29).f24;
          _nw305[(new BigNumber(26)).toNumber()] = !(new BigNumber((_1867_v40).length)).isEqualTo((_1823_v9).f28);
          _nw305[(new BigNumber(27)).toNumber()] = (_this).fm6((_1823_v9).fm17(!(p1), (_1870_v43).dtor_cf44, (_1823_v9).f28, _1854_v27, globalState), globalState);
          _nw305[(new BigNumber(28)).toNumber()] = true;
          _1871_v44 = _nw305;
          let _index306 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_1871_v44).length));
          (_1871_v44)[_index306] = (_1856_v29).f24;
          let _1872_v45;
          _1872_v45 = _dafny.Seq.of(_1865_v38);
          let _index307 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_1858_v31).length));
          let _index308 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_1871_v44).length));
          let _rhs306 = (((_1856_v29).f24) ? (_1857_v30) : (_1857_v30));
          let _rhs307 = (_1860_v33).Merge(((p1) ? (_1860_v33) : (_1860_v33)));
          let _rhs308 = !(!((_1856_v29).f24));
          let _rhs309 = _module.__default.safeDivisionInt((_1869_v42)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.FromArray(_1872_v45)).cardinality()), new BigNumber((_1869_v42).length))], _module.__default.fm0((_1869_v42)[_module.__default.safeIndex(new BigNumber((_1855_v28).length), new BigNumber((_1869_v42).length))], new BigNumber(431), p0, _module.__default.fm2(_1859_v32, globalState), globalState));
          let _rhs310 = (_dafny.ZERO).minus((p2)[_module.__default.safeIndex(new BigNumber(688), new BigNumber((p2).length))]);
          let _lhs286 = _1858_v31;
          let _lhs287 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_1858_v31).length));
          let _lhs288 = _1871_v44;
          let _lhs289 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_1871_v44).length));
          let _lhs290 = globalState;
          _1857_v30 = _rhs306;
          _lhs286[_lhs287] = _rhs307;
          _lhs288[_lhs289] = _rhs308;
          r1 = _rhs309;
          _lhs290.f7 = _rhs310;
        } else if (_source23.is_DC41) {
          let _1873___mcc_h5 = (_source23).cf57;
          let _1874_cf57 = _1873___mcc_h5;
          (globalState).f7 = new BigNumber(577);
          let _1875_v46;
          _1875_v46 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-300),p1);
          let _1876_v47;
          _1876_v47 = _dafny.Seq.of(_1875_v46);
          let _1877_v48;
          let _nw306 = new _module.C1();
          _nw306.__ctor(p0, _1876_v47);
          _1877_v48 = _nw306;
          (globalState).f21 = p1;
          let _1878_v49;
          _1878_v49 = _module.D19.create_DC48(p0, p1);
          r1 = (_1878_v49).dtor_cf65;
        } else {
          let _1879___mcc_h6 = (_source23).cf51;
          let _1880_cf51 = _1879___mcc_h6;
          let _1881_v50;
          _1881_v50 = _dafny.Seq.of(p1, p1, p1);
          let _1882_v51;
          _1882_v51 = _dafny.MultiSet.fromElements(p1);
          (globalState).f2 = ((_1881_v50)[_module.__default.safeIndex(new BigNumber((_1882_v51).cardinality()), new BigNumber((_1881_v50).length))]) === (p1);
          let _1883_v52;
          _1883_v52 = _dafny.MultiSet.fromElements(_module.__default.fm0(p3, p0, (_1823_v9).f28, p1, globalState), (_1823_v9).f28, (_1823_v9).f28, (_1823_v9).f28, (_1823_v9).f28);
          (globalState).f2 = (_1883_v52).IsDisjointFrom(_dafny.MultiSet.fromElements((_1823_v9).f28));
          let _1884_v53;
          _1884_v53 = new _dafny.CodePoint('n'.codePointAt(0));
          _1884_v53 = _module.__default.fm34(globalState);
          let _1885_v54;
          let _nw307 = Array((new BigNumber(3)).toNumber()).fill(false);
          _1885_v54 = _nw307;
          let _index309 = _module.__default.safeIndex(new BigNumber(65), new BigNumber((_1885_v54).length));
          (_1885_v54)[_index309] = !(p1) || (p1);
          let _1886_v55;
          _1886_v55 = _dafny.Set.fromElements(true);
          let _index310 = _module.__default.safeIndex(new BigNumber(65), new BigNumber((_1885_v54).length));
          (_1885_v54)[_index310] = ((!(p1)) ? ((_1886_v55).IsSubsetOf(_1886_v55)) : (!(p1)));
        }
        let _1887_v56;
        _1887_v56 = new _dafny.CodePoint('g'.codePointAt(0));
        let _1888_v57;
        let _nw308 = Array((new BigNumber(13)).toNumber());
        _nw308[(_dafny.ZERO).toNumber()] = _1887_v56;
        _nw308[(_dafny.ONE).toNumber()] = _1887_v56;
        _nw308[(new BigNumber(2)).toNumber()] = _1887_v56;
        _nw308[(new BigNumber(3)).toNumber()] = _1887_v56;
        _nw308[(new BigNumber(4)).toNumber()] = _1887_v56;
        _nw308[(new BigNumber(5)).toNumber()] = _1887_v56;
        _nw308[(new BigNumber(6)).toNumber()] = new _dafny.CodePoint('k'.codePointAt(0));
        _nw308[(new BigNumber(7)).toNumber()] = new _dafny.CodePoint('f'.codePointAt(0));
        _nw308[(new BigNumber(8)).toNumber()] = _1887_v56;
        _nw308[(new BigNumber(9)).toNumber()] = _1887_v56;
        _nw308[(new BigNumber(10)).toNumber()] = _1887_v56;
        _nw308[(new BigNumber(11)).toNumber()] = _1887_v56;
        _nw308[(new BigNumber(12)).toNumber()] = _1887_v56;
        _1888_v57 = _nw308;
        let _1889_v58;
        _1889_v58 = _dafny.Map.Empty.slice().updateUnsafe(p3,_1888_v57);
        _1888_v57 = (((_1889_v58).contains((_1823_v9).f28)) ? ((_1889_v58).get((_1823_v9).f28)) : (_1888_v57));
        let _1890_v59;
        _1890_v59 = _dafny.Seq.of(p1);
        let _1891_v60;
        _1891_v60 = _dafny.Seq.UnicodeFromString("sbxhia");
        let _1892_v61;
        _1892_v61 = _dafny.MultiSet.fromElements(new BigNumber((_1891_v60).length), p3);
        let _1893_v62;
        _1893_v62 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_1890_v59).length)).multipliedBy(new BigNumber((_1892_v61).cardinality())),_1887_v56);
        let _rhs311 = (((_1893_v62).contains(_module.__default.safeModuloInt((_1823_v9).f28, new BigNumber((_dafny.Seq.of(new BigNumber((_1891_v60).length), (_1823_v9).f28)).length)))) ? ((_1893_v62).get(_module.__default.safeModuloInt((_1823_v9).f28, new BigNumber((_dafny.Seq.of(new BigNumber((_1891_v60).length), (_1823_v9).f28)).length)))) : (_1887_v56));
        let _rhs312 = (_1823_v9).fm6((_1823_v9).f28, globalState);
        let _lhs291 = globalState;
        _1887_v56 = _rhs311;
        _lhs291.f2 = _rhs312;
        let _index311 = _module.__default.safeIndex(new BigNumber(357), new BigNumber((p2).length));
        (p2)[_index311] = p3;
        let _index312 = _module.__default.safeIndex(new BigNumber(357), new BigNumber((p2).length));
        (p2)[_index312] = new BigNumber((_1891_v60).length);
        r1 = new BigNumber(794);
      }
      let _1894_v63;
      let _nw309 = new _module.C2();
      _nw309.__ctor(((_1823_v9).f28).plus(p0));
      _1894_v63 = _nw309;
      let _1895_v64;
      _1895_v64 = _dafny.Map.Empty.slice().updateUnsafe(false,(_1823_v9).f28);
      let _1896_v65;
      _1896_v65 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(p1,p0), _1895_v64);
      let _1897_v66;
      _1897_v66 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_1896_v65).IsSubsetOf(_1896_v65));
      let _1898_i4;
      _1898_i4 = _dafny.ZERO;
      L12: {
        while (!((((_1897_v66).contains(p1)) ? ((_1897_v66).get(p1)) : (true)))) {
          C12: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1898_i4)) {
              break L12;
            }
            _1898_i4 = (_1898_i4).plus(_dafny.ONE);
            let _1899_v67;
            let _nw310 = new _module.C3();
            _nw310.__ctor();
            _1899_v67 = _nw310;
            (globalState).f16 = new BigNumber((_module.__default.fm1((_1823_v9).f28, globalState)).length);
            let _index313 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((p2).length));
            (p2)[_index313] = p0;
            let _index314 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((p2).length));
            (p2)[_index314] = p0;
            let _1900_v68;
            let _nw311 = Array((new BigNumber(21)).toNumber()).fill(false);
            _1900_v68 = _nw311;
            let _index315 = _module.__default.safeIndex(new BigNumber(629), new BigNumber((_1900_v68).length));
            (_1900_v68)[_index315] = p1;
            let _index316 = _module.__default.safeIndex(new BigNumber(629), new BigNumber((_1900_v68).length));
            (_1900_v68)[_index316] = p1;
          }
        }
      }
      let _1901_v69;
      _1901_v69 = _dafny.Seq.of(p1, true, (p1) && (p1));
      let _1902_v70;
      _1902_v70 = _dafny.Seq.of((_1894_v63).f28, new BigNumber(773));
      let _1903_v71;
      _1903_v71 = _dafny.Seq.UnicodeFromString("ygs");
      let _1904_i5;
      _1904_i5 = _dafny.ZERO;
      L13: {
        while ((_1901_v69)[_module.__default.safeIndex(_module.__default.safeDivisionInt((_dafny.ZERO).minus((_1823_v9).f28), (_1902_v70)[_module.__default.safeIndex(new BigNumber((_1903_v71).length), new BigNumber((_1902_v70).length))]), new BigNumber((_1901_v69).length))]) {
          C13: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1904_i5)) {
              break L13;
            }
            _1904_i5 = (_1904_i5).plus(_dafny.ONE);
            let _1905_v72;
            _1905_v72 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("qxmym"), _1903_v71, _1903_v71, _module.__default.fm1(new BigNumber((_1903_v71).length), globalState), _1903_v71);
            let _1906_v73;
            _1906_v73 = _dafny.MultiSet.fromElements((_1905_v72)[_module.__default.safeIndex((_1823_v9).f28, new BigNumber((_1905_v72).length))], _dafny.Seq.Create(_module.__default.abs(new BigNumber(991)), function (_1907_i6) {
              return new _dafny.CodePoint('d'.codePointAt(0));
            }), _1903_v71);
            let _1908_v74;
            _1908_v74 = _dafny.Map.Empty.slice().updateUnsafe((_1823_v9).f28,_module.__default.fm55((_1894_v63).f28, p1, p1, globalState));
            let _1909_v75;
            _1909_v75 = _dafny.Map.Empty.slice().updateUnsafe((_1901_v69)[_module.__default.safeIndex(p0, new BigNumber((_1901_v69).length))],(((_1908_v74).contains(p3)) ? ((_1908_v74).get(p3)) : (_1906_v73)));
            let _1910_v76;
            _1910_v76 = _dafny.Seq.of((_1906_v73).update(_1903_v71, _module.__default.abs((_1894_v63).f28)), _dafny.MultiSet.FromArray(_dafny.Seq.of(_1903_v71)));
            let _1911_v77;
            let _nw312 = Array((new BigNumber(17)).toNumber());
            _nw312[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.fromElements(_1903_v71, _1903_v71);
            _nw312[(_dafny.ONE).toNumber()] = _1906_v73;
            _nw312[(new BigNumber(2)).toNumber()] = _1906_v73;
            _nw312[(new BigNumber(3)).toNumber()] = ((p1) ? (_1906_v73) : (_dafny.MultiSet.FromArray(_1905_v72)));
            _nw312[(new BigNumber(4)).toNumber()] = (_1906_v73).Intersect(_1906_v73);
            _nw312[(new BigNumber(5)).toNumber()] = _1906_v73;
            _nw312[(new BigNumber(6)).toNumber()] = _dafny.MultiSet.fromElements(_1903_v71, _dafny.Seq.UnicodeFromString("biloj"));
            _nw312[(new BigNumber(7)).toNumber()] = _1906_v73;
            _nw312[(new BigNumber(8)).toNumber()] = _1906_v73;
            _nw312[(new BigNumber(9)).toNumber()] = _1906_v73;
            _nw312[(new BigNumber(10)).toNumber()] = _1906_v73;
            _nw312[(new BigNumber(11)).toNumber()] = (((_1909_v75).contains(p1)) ? ((_1909_v75).get(p1)) : (_1906_v73));
            _nw312[(new BigNumber(12)).toNumber()] = (_1910_v76)[_module.__default.safeIndex(new BigNumber(257), new BigNumber((_1910_v76).length))];
            _nw312[(new BigNumber(13)).toNumber()] = _1906_v73;
            _nw312[(new BigNumber(14)).toNumber()] = _dafny.MultiSet.fromElements(_1903_v71);
            _nw312[(new BigNumber(15)).toNumber()] = (_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-724)), function (_1912_i7) {
              return new _dafny.CodePoint('t'.codePointAt(0));
            }))).update(_1903_v71, _module.__default.abs((_1894_v63).f28));
            _nw312[(new BigNumber(16)).toNumber()] = _1906_v73;
            _1911_v77 = _nw312;
            let _index317 = _module.__default.safeIndex(new BigNumber(233), new BigNumber((_1911_v77).length));
            (_1911_v77)[_index317] = _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(_1903_v71), _1905_v72));
            let _index318 = _module.__default.safeIndex(new BigNumber(233), new BigNumber((_1911_v77).length));
            (_1911_v77)[_index318] = _1906_v73;
            (globalState).f16 = ((_dafny.ZERO).minus((_1894_v63).f28)).plus(((_1894_v63).f28).plus((_1894_v63).f28));
            (globalState).f13 = !(p1);
            let _1913_v78;
            let _nw313 = new _module.C3();
            _nw313.__ctor();
            _1913_v78 = _nw313;
            let _index319 = _module.__default.safeIndex(new BigNumber(344), new BigNumber((p2).length));
            (p2)[_index319] = (_1894_v63).f28;
            let _index320 = _module.__default.safeIndex(new BigNumber(344), new BigNumber((p2).length));
            let _rhs313 = _1913_v78;
            let _rhs314 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of((_1823_v9).f28, p0), _dafny.Seq.of(p3));
            let _rhs315 = p3;
            let _rhs316 = p1;
            let _rhs317 = (_1823_v9).f28;
            let _lhs292 = globalState;
            let _lhs293 = p2;
            let _lhs294 = _module.__default.safeIndex(new BigNumber(344), new BigNumber((p2).length));
            let _lhs295 = globalState;
            let _lhs296 = globalState;
            _1913_v78 = _rhs313;
            _lhs292.f2 = _rhs314;
            _lhs293[_lhs294] = _rhs315;
            _lhs295.f13 = _rhs316;
            _lhs296.f7 = _rhs317;
          }
        }
      }
      let _1914_v79;
      let _nw314 = new _module.C0();
      _nw314.__ctor(p1);
      _1914_v79 = _nw314;
      let _1915_v80;
      _1915_v80 = _dafny.Seq.of(_1914_v79);
      let _1916_v81;
      _1916_v81 = _dafny.Map.Empty.slice().updateUnsafe((_1915_v80)[_module.__default.safeIndex(new BigNumber((_1903_v71).length), new BigNumber((_1915_v80).length))],(_1823_v9).f28);
      r0 = (_1901_v69)[_module.__default.safeIndex(new BigNumber((_1916_v81).length), new BigNumber((_1901_v69).length))];
      r1 = ((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,(_1894_v63).f28)).length))).plus((_1894_v63).f28);
      let _init56 = ((_1917_v79) => function (_1918_i8) {
        return _dafny.Map.Empty.slice().updateUnsafe(_1917_v79.f27,_module.D3.create_DC6(_dafny.Seq.UnicodeFromString("twmsayjd"), new _dafny.CodePoint('c'.codePointAt(0)), _1917_v79.f27));
      })(_1914_v79);
      let _nw315 = Array((new BigNumber(10)).toNumber());
      for (let _i0_56 = 0; _i0_56 < new BigNumber(_nw315.length); _i0_56++) {
        _nw315[_i0_56] = _init56(new BigNumber(_i0_56));
      }
      r2 = _nw315;
      return [r0, r1, r2];
    }
    m3(p0, globalState) {
      let _this = this;
      let _1919_v0;
      _1919_v0 = true;
      let _1920_i0;
      _1920_i0 = _dafny.ZERO;
      L14: {
        while (!(_1919_v0)) {
          C14: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1920_i0)) {
              break L14;
            }
            _1920_i0 = (_1920_i0).plus(_dafny.ONE);
            if (_1919_v0) {
              let _1921_v1;
              let _init57 = ((_1922_v0) => function (_1923_i1) {
                return _1922_v0;
              })(_1919_v0);
              let _nw316 = Array((new BigNumber(6)).toNumber());
              for (let _i0_57 = 0; _i0_57 < new BigNumber(_nw316.length); _i0_57++) {
                _nw316[_i0_57] = _init57(new BigNumber(_i0_57));
              }
              _1921_v1 = _nw316;
              let _1924_v2;
              _1924_v2 = _dafny.Seq.of(_1921_v1);
              let _1925_v3;
              _1925_v3 = new BigNumber(533);
              _1921_v1 = (_1924_v2)[_module.__default.safeIndex(_1925_v3, new BigNumber((_1924_v2).length))];
              let _1926_v4;
              let _init58 = function (_1927_i2) {
                return _module.D15.create_DC33(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("bauwltxil")).length)));
              };
              let _nw317 = Array((new BigNumber(20)).toNumber());
              for (let _i0_58 = 0; _i0_58 < new BigNumber(_nw317.length); _i0_58++) {
                _nw317[_i0_58] = _init58(new BigNumber(_i0_58));
              }
              _1926_v4 = _nw317;
              let _index321 = _module.__default.safeIndex(new BigNumber(559), new BigNumber((_1926_v4).length));
              (_1926_v4)[_index321] = _module.D15.create_DC33(_module.__default.fm21(globalState));
              let _1928_v5;
              _1928_v5 = _dafny.Seq.of(_1919_v0);
              let _1929_v6;
              _1929_v6 = _dafny.Map.Empty.slice().updateUnsafe(_module.D14.create_DC29(_1925_v3, _dafny.Map.Empty.slice().updateUnsafe(false,_1928_v5)),new BigNumber(-262));
              let _1930_v7;
              _1930_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1919_v0,p0);
              let _1931_v8;
              _1931_v8 = _dafny.Seq.of(_1925_v3, new BigNumber((_1929_v6).length), new BigNumber((_1930_v7).length));
              let _1932_v9;
              _1932_v9 = _module.D15.create_DC33(_1931_v8);
              let _pat_let_tv32 = _1931_v8;
              let _index322 = _module.__default.safeIndex(new BigNumber(559), new BigNumber((_1926_v4).length));
              let _rhs318 = _1919_v0;
              let _rhs319 = (_1925_v3).isLessThan(_1925_v3);
              let _rhs320 = function (_pat_let26_0) {
                return function (_1933_dt__update__tmp_h0) {
                  return function (_pat_let27_0) {
                    return function (_1934_dt__update_hcf44_h0) {
                      return _module.D15.create_DC33(_1934_dt__update_hcf44_h0);
                    }(_pat_let27_0);
                  }(_pat_let_tv32);
                }(_pat_let26_0);
              }(_1932_v9);
              let _rhs321 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_1925_v3, _module.__default.safeDivisionInt(_1925_v3, _1925_v3)));
              let _lhs297 = globalState;
              let _lhs298 = globalState;
              let _lhs299 = _1926_v4;
              let _lhs300 = _module.__default.safeIndex(new BigNumber(559), new BigNumber((_1926_v4).length));
              let _lhs301 = globalState;
              _lhs297.f13 = _rhs318;
              _lhs298.f21 = _rhs319;
              _lhs299[_lhs300] = _rhs320;
              _lhs301.f16 = _rhs321;
              let _1935_v10;
              _1935_v10 = _dafny.Set.fromElements(_1925_v3);
              let _1936_v11;
              let _nw318 = new _module.C5();
              _nw318.__ctor((_1935_v10).IsProperSubsetOf(_dafny.Set.fromElements(_1925_v3)));
              _1936_v11 = _nw318;
              _1919_v0 = (_1936_v11).f24;
              let _index323 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_1921_v1).length));
              (_1921_v1)[_index323] = (_1928_v5)[_module.__default.safeIndex(_1925_v3, new BigNumber((_1928_v5).length))];
              let _1937_v12;
              _1937_v12 = _dafny.Set.fromElements(_1931_v8);
              let _1938_v13;
              _1938_v13 = _dafny.Seq.of(_1931_v8, _1931_v8, _1931_v8);
              let _1939_v15;
              _1939_v15 = _dafny.MultiSet.fromElements(_1919_v0, _1919_v0, (_1936_v11).f24);
              let _index324 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_1921_v1).length));
              let _rhs322 = (_1936_v11).f24;
              let _rhs323 = ((_1937_v12).Difference(function () {
                let _coll75 = new _dafny.Set();
                for (const _compr_75 of (_1938_v13).Elements) {
                  let _1940_v14 = _compr_75;
                  if (_dafny.Seq.contains(_1938_v13, _1940_v14)) {
                    _coll75.add(_1940_v14);
                  }
                }
                return _coll75;
              }())).IsProperSubsetOf(_dafny.Set.fromElements(_dafny.Seq.update(_1931_v8, _module.__default.safeIndex(_1925_v3, new BigNumber((_1931_v8).length)), _1925_v3)));
              let _rhs324 = (_1939_v15).IsDisjointFrom(_1939_v15);
              let _lhs302 = globalState;
              let _lhs303 = _1921_v1;
              let _lhs304 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_1921_v1).length));
              let _lhs305 = globalState;
              _lhs302.f2 = _rhs322;
              _lhs303[_lhs304] = _rhs323;
              _lhs305.f13 = _rhs324;
            } else {
              let _1941_v16;
              let _nw319 = new _module.C4();
              _nw319.__ctor(_1919_v0, _1919_v0);
              _1941_v16 = _nw319;
              let _1942_v17;
              let _nw320 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Set.Empty);
              _1942_v17 = _nw320;
              let _1943_v18;
              _1943_v18 = _dafny.Set.fromElements(new BigNumber(738));
              let _index325 = _module.__default.safeIndex(new BigNumber(679), new BigNumber((_1942_v17).length));
              (_1942_v17)[_index325] = _1943_v18;
              let _index326 = _module.__default.safeIndex(new BigNumber(679), new BigNumber((_1942_v17).length));
              (_1942_v17)[_index326] = _1943_v18;
              let _1944_v19;
              let _nw321 = Array((new BigNumber(3)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
              _1944_v19 = _nw321;
              let _1945_v20;
              _1945_v20 = new _dafny.CodePoint('f'.codePointAt(0));
              let _index327 = _module.__default.safeIndex(new BigNumber(577), new BigNumber((_1944_v19).length));
              (_1944_v19)[_index327] = _1945_v20;
              let _index328 = _module.__default.safeIndex(new BigNumber(577), new BigNumber((_1944_v19).length));
              (_1944_v19)[_index328] = _1945_v20;
              let _1946_v21;
              _1946_v21 = new BigNumber(486);
              let _1947_v22;
              _1947_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1946_v21,(_1944_v19)[_module.__default.safeIndex(new BigNumber(577), new BigNumber((_1944_v19).length))]);
              let _1948_v23;
              _1948_v23 = _dafny.Seq.of(_1919_v0);
              let _1949_v24;
              _1949_v24 = _dafny.Seq.UnicodeFromString("nwdg");
              let _1950_v25;
              _1950_v25 = _dafny.Seq.of(_1949_v24, _1949_v24);
              let _1951_v26;
              _1951_v26 = _dafny.Map.Empty.slice().updateUnsafe((_1950_v25)[_module.__default.safeIndex(_1946_v21, new BigNumber((_1950_v25).length))],(_1941_v16).f26);
              let _1952_v27;
              let _nw322 = new _module.C0();
              _nw322.__ctor(_1941_v16.f25);
              _1952_v27 = _nw322;
              let _1953_v28;
              _1953_v28 = _dafny.Map.Empty.slice().updateUnsafe(_1952_v27,_1919_v0);
              let _1954_v29;
              _1954_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1952_v27.f27,!(_1941_v16.f25));
              let _1955_v30;
              let _nw323 = Array((new BigNumber(7)).toNumber());
              _nw323[(_dafny.ZERO).toNumber()] = (new BigNumber((_1947_v22).length)).isLessThanOrEqualTo(_module.__default.fm0(_1946_v21, _1946_v21, _module.__default.fm0(_1946_v21, new BigNumber(((_1942_v17)[_module.__default.safeIndex(new BigNumber(679), new BigNumber((_1942_v17).length))]).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(19)), ((_1956_v21) => function (_1957_i3) {
                return _1956_v21;
              })(_1946_v21))).length), (_1948_v23)[_module.__default.safeIndex(_1946_v21, new BigNumber((_1948_v23).length))], globalState), (_1941_v16).f26, globalState));
              _nw323[(_dafny.ONE).toNumber()] = (((_1951_v26).contains(_1949_v24)) ? ((_1951_v26).get(_1949_v24)) : (_1919_v0));
              _nw323[(new BigNumber(2)).toNumber()] = !((!((((_1953_v28).contains(_1952_v27)) ? ((_1953_v28).get(_1952_v27)) : (_1941_v16.f25)))) === (false));
              _nw323[(new BigNumber(3)).toNumber()] = _module.__default.fm2(_1954_v29, globalState);
              _nw323[(new BigNumber(4)).toNumber()] = (true) && (!(false));
              _nw323[(new BigNumber(5)).toNumber()] = (_1919_v0) === (!(_1919_v0));
              _nw323[(new BigNumber(6)).toNumber()] = false;
              _1955_v30 = _nw323;
              let _index329 = _module.__default.safeIndex(new BigNumber(494), new BigNumber((_1955_v30).length));
              (_1955_v30)[_index329] = !_dafny.Seq.contains(_1949_v24, _1945_v20);
              let _index330 = _module.__default.safeIndex(new BigNumber(494), new BigNumber((_1955_v30).length));
              (_1955_v30)[_index330] = (_1941_v16).f26;
              let _1958_v31;
              let _nw324 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
              _1958_v31 = _nw324;
              _1958_v31 = _1958_v31;
            }
            let _1959_v32;
            _1959_v32 = new BigNumber(678);
            (globalState).f7 = (_dafny.ZERO).minus(_1959_v32);
            let _1960_v33;
            _1960_v33 = new _dafny.CodePoint('p'.codePointAt(0));
            let _1961_v34;
            _1961_v34 = _dafny.Map.Empty.slice().updateUnsafe(_1919_v0,_1919_v0);
            let _1962_v35;
            let _nw325 = Array((new BigNumber(2)).toNumber());
            _nw325[(_dafny.ZERO).toNumber()] = _module.__default.fm18(true, _1960_v33, false, globalState);
            _nw325[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_1919_v0), _module.__default.fm18(true, _1960_v33, _module.__default.fm2(_1961_v34, globalState), globalState));
            _1962_v35 = _nw325;
            let _1963_v36;
            _1963_v36 = _dafny.Seq.of(_1919_v0, !(_1919_v0));
            let _index331 = _module.__default.safeIndex(new BigNumber(847), new BigNumber((_1962_v35).length));
            (_1962_v35)[_index331] = _1963_v36;
            let _index332 = _module.__default.safeIndex(new BigNumber(847), new BigNumber((_1962_v35).length));
            (_1962_v35)[_index332] = _dafny.Seq.of(_1919_v0, _1919_v0);
            if (_1919_v0) {
              let _1964_v37;
              _1964_v37 = _dafny.Seq.UnicodeFromString("dhdknns");
              let _1965_v38;
              _1965_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1919_v0,(_1962_v35)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1962_v35).length))]);
              let _1966_v39;
              _1966_v39 = _dafny.Map.Empty.slice().updateUnsafe(_1959_v32,_1959_v32);
              let _1967_v41;
              _1967_v41 = _module.D13.create_DC25(_1966_v39);
              let _1968_v42;
              _1968_v42 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.update(_1963_v36, _module.__default.safeIndex(new BigNumber((function () {
                let _coll76 = new _dafny.Map();
                for (const _compr_76 of _dafny.IntegerRange(new BigNumber(368), new BigNumber(628))) {
                  let _1969_v40 = _compr_76;
                  if (((new BigNumber(368)).isLessThanOrEqualTo(_1969_v40)) && ((_1969_v40).isLessThan(new BigNumber(628)))) {
                    _coll76.push([(_1969_v40).minus(_1959_v32),_1919_v0]);
                  }
                }
                return _coll76;
              }()).length), new BigNumber((_1963_v36).length)), _1919_v0)).length),_1967_v41);
              let _1970_v43;
              _1970_v43 = _module.D17.create_DC39(new BigNumber(4), _1966_v39, _1919_v0, _1968_v42, _1959_v32);
              let _1971_v44;
              _1971_v44 = _dafny.Map.Empty.slice().updateUnsafe(_1959_v32,_1919_v0);
              let _1972_v45;
              _1972_v45 = _dafny.Seq.of(_1971_v44);
              let _1973_v46;
              let _nw326 = new _module.C1();
              _nw326.__ctor(_module.__default.safeDivisionInt(_module.__default.fm0(new BigNumber((_1964_v37).length), new BigNumber((_1965_v38).length), _1959_v32, (_1970_v43).dtor_cf54, globalState), new BigNumber((_dafny.Seq.UnicodeFromString("cfemxm")).length)), _dafny.Seq.Concat(_dafny.Seq.of(_1971_v44), _1972_v45));
              _1973_v46 = _nw326;
              let _1974_v47;
              _1974_v47 = _dafny.Seq.of(_1959_v32);
              let _1975_v48;
              let _init59 = ((_1976_v0) => function (_1977_i4) {
                return _1976_v0;
              })(_1919_v0);
              let _nw327 = Array((_dafny.ONE).toNumber());
              for (let _i0_59 = 0; _i0_59 < new BigNumber(_nw327.length); _i0_59++) {
                _nw327[_i0_59] = _init59(new BigNumber(_i0_59));
              }
              _1975_v48 = _nw327;
              let _1978_v49;
              _1978_v49 = _1975_v48;
              let _1979_v50;
              let _nw328 = Array((new BigNumber(9)).toNumber());
              _nw328[(_dafny.ZERO).toNumber()] = _1975_v48;
              _nw328[(_dafny.ONE).toNumber()] = _1975_v48;
              _nw328[(new BigNumber(2)).toNumber()] = _1975_v48;
              _nw328[(new BigNumber(3)).toNumber()] = _1975_v48;
              _nw328[(new BigNumber(4)).toNumber()] = (_1978_v49);
              _nw328[(new BigNumber(5)).toNumber()] = _1975_v48;
              _nw328[(new BigNumber(6)).toNumber()] = _1975_v48;
              _nw328[(new BigNumber(7)).toNumber()] = _1975_v48;
              _nw328[(new BigNumber(8)).toNumber()] = _1975_v48;
              _1979_v50 = _nw328;
              let _1980_v51;
              _1980_v51 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(_1973_v46.f29, new BigNumber(310)), _1974_v47),_1979_v50);
              _1980_v51 = (_1980_v51).update((_1973_v46.f29).isLessThanOrEqualTo(new BigNumber(-382)), _1979_v50);
              let _1981_v52;
              _1981_v52 = _module.D18.create_DC43();
              let _1982_v53;
              _1982_v53 = _dafny.Map.Empty.slice().updateUnsafe(_1959_v32,_1981_v52);
              (globalState).f2 = !(_1982_v53).contains(((_1919_v0) ? (new BigNumber(467)) : (_1959_v32)));
              let _1983_v54;
              let _nw329 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _1983_v54 = _nw329;
              _1983_v54 = _1983_v54;
              let _1984_v55;
              let _nw330 = new _module.C3();
              _nw330.__ctor();
              _1984_v55 = _nw330;
            } else {
              let _1985_v56;
              _1985_v56 = _dafny.Seq.UnicodeFromString("ap");
              let _1986_v57;
              _1986_v57 = _module.D13.create_DC26(_1985_v56, _1919_v0, _1919_v0);
              let _1987_v58;
              _1987_v58 = _dafny.Map.Empty.slice().updateUnsafe((_1986_v57).dtor_cf33,_1919_v0);
              _1987_v58 = (_1987_v58).update(_1985_v56, !(true));
              let _1988_v59;
              _1988_v59 = _dafny.Map.Empty.slice().updateUnsafe(_1959_v32,_1919_v0);
              let _1989_v60;
              _1989_v60 = _dafny.Seq.of(_1988_v59);
              let _1990_v61;
              let _nw331 = new _module.C1();
              _nw331.__ctor(new BigNumber((_1985_v56).length), _1989_v60);
              _1990_v61 = _nw331;
              let _1991_v62;
              _1991_v62 = _dafny.Seq.of(_1990_v61, _1990_v61, _1990_v61, _1990_v61, _1990_v61);
              let _1992_v63;
              _1992_v63 = _dafny.Map.Empty.slice().updateUnsafe(_1963_v36,_1985_v56);
              _module.__default.m0(_dafny.areEqual(_1991_v62, _1991_v62), _1959_v32, (_1992_v63).update((_1962_v35)[_module.__default.safeIndex(new BigNumber(847), new BigNumber((_1962_v35).length))], _1985_v56), new BigNumber((_1961_v34).length), globalState);
              (globalState).f19 = _1985_v56;
              let _1993_v64;
              _1993_v64 = _module.D0.create_DC0(_module.__default.fm2(_1961_v34, globalState));
              let _1994_v65;
              _1994_v65 = _module.D0.create_DC1(_1993_v64);
              let _1995_v66;
              _1995_v66 = _dafny.MultiSet.fromElements(_1994_v65, _1994_v65, _1994_v65);
              (globalState).f16 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber(-782)), _1959_v32), (_dafny.ZERO).minus((((_1995_v66).contains(_1994_v65)) ? ((_1995_v66).get(_1994_v65)) : (_1990_v61.f29))));
              let _1996_v67;
              _1996_v67 = _dafny.Map.Empty.slice().updateUnsafe(_1919_v0,_1990_v61.f29);
              let _1997_v68;
              _1997_v68 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm56(_1990_v61.f29, new BigNumber((_1996_v67).length), _1919_v0, globalState),(_1990_v61).fm29(globalState));
              _1997_v68 = _1997_v68;
            }
          }
        }
      }
      let _1998_v69;
      _1998_v69 = new BigNumber(816);
      let _index333 = _module.__default.safeIndex(new BigNumber(213), new BigNumber((p0).length));
      (p0)[_index333] = _1998_v69;
      let _1999_v70;
      let _nw332 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Set.Empty);
      _1999_v70 = _nw332;
      let _2000_v71;
      _2000_v71 = _module.D7.create_DC16(_1999_v70);
      let _2001_v72;
      _2001_v72 = _dafny.MultiSet.fromElements(_2000_v71);
      let _index334 = _module.__default.safeIndex(new BigNumber(213), new BigNumber((p0).length));
      (p0)[_index334] = new BigNumber(((_2001_v72).update(_2000_v71, _module.__default.abs(_1998_v69))).cardinality());
      let _2002_v73;
      let _nw333 = new _module.C5();
      _nw333.__ctor(_1919_v0);
      _2002_v73 = _nw333;
      let _2003_v74;
      _2003_v74 = _dafny.MultiSet.fromElements(p0);
      let _source24 = _2003_v74;
      let _2004___mcc_h0 = _source24;
      let _2005_cf4 = _2004___mcc_h0;
      let _2006_v75;
      _2006_v75 = new _dafny.CodePoint('x'.codePointAt(0));
      let _2007_v76;
      _2007_v76 = _dafny.Seq.of(_2006_v75);
      _1998_v69 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt(_1998_v69, new BigNumber((_2007_v76).length)), _1998_v69);
      let _2008_v77;
      let _nw334 = Array((new BigNumber(22)).toNumber()).fill(_module.D19.Default());
      _2008_v77 = _nw334;
      let _2009_v78;
      let _init60 = ((_2010_v73) => function (_2011_i5) {
        return (_2010_v73).f24;
      })(_2002_v73);
      let _nw335 = Array((new BigNumber(6)).toNumber());
      for (let _i0_60 = 0; _i0_60 < new BigNumber(_nw335.length); _i0_60++) {
        _nw335[_i0_60] = _init60(new BigNumber(_i0_60));
      }
      _2009_v78 = _nw335;
      let _2012_v79;
      _2012_v79 = _module.D19.create_DC45(_dafny.MultiSet.fromElements(_2009_v78));
      let _index335 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_2008_v77).length));
      (_2008_v77)[_index335] = _2012_v79;
      let _index336 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_2008_v77).length));
      (_2008_v77)[_index336] = _2012_v79;
      let _2013_v80;
      _2013_v80 = _dafny.MultiSet.fromElements((_2002_v73).f24);
      let _rhs325 = _2013_v80;
      let _rhs326 = true;
      let _lhs306 = globalState;
      _2013_v80 = _rhs325;
      _lhs306.f2 = _rhs326;
      (globalState).f13 = (_2002_v73).f24;
      let _pat_let_tv33 = _1998_v69;
      let _pat_let_tv34 = _2002_v73;
      let _pat_let_tv35 = _1998_v69;
      let _pat_let_tv36 = _1998_v69;
      let _pat_let_tv37 = _1998_v69;
      (globalState).f7 = function (_source25) {
        if (_source25.is_DC36) {
          let _2014___mcc_h1 = (_source25).cf47;
          let _2015___mcc_h2 = (_source25).cf48;
          let _2016___mcc_h3 = (_source25).cf49;
          let _2017_cf49 = _2016___mcc_h3;
          let _2018_cf48 = _2015___mcc_h2;
          let _2019_cf47 = _2014___mcc_h1;
          return (_pat_let_tv33).multipliedBy(new BigNumber(-220));
        } else if (_source25.is_DC35) {
          let _2020___mcc_h4 = (_source25).cf46;
          let _2021_cf46 = _2020___mcc_h4;
          if ((_pat_let_tv34).f24) {
            return _pat_let_tv35;
          } else {
            return new BigNumber((_dafny.Seq.of(_pat_let_tv36)).length);
          }
        } else {
          let _2022___mcc_h5 = (_source25).cf50;
          let _2023_cf50 = _2022___mcc_h5;
          return (new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(_pat_let_tv37))).length)).minus(new BigNumber(405));
        }
      }(_module.__default.fm57(globalState));
      let _2024_v81;
      _2024_v81 = new _dafny.CodePoint('y'.codePointAt(0));
      _2024_v81 = new _dafny.CodePoint('c'.codePointAt(0));
      return;
    }
    m1(p0, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = false;
      let _2025_v0;
      _2025_v0 = true;
      let _2026_v1;
      _2026_v1 = _dafny.Seq.of(_2025_v0);
      let _2027_v2;
      _2027_v2 = _dafny.MultiSet.fromElements(p0, p0);
      let _2028_v3;
      let _nw336 = new _module.C4();
      _nw336.__ctor(((_2026_v1)[_module.__default.safeIndex(p0, new BigNumber((_2026_v1).length))]) && (_2025_v0), (_dafny.Map.Empty.slice().updateUnsafe(_2027_v2,p0)).contains(_2027_v2));
      _2028_v3 = _nw336;
      if ((_2028_v3).f26) {
        let _2029_v4;
        _2029_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("iuyqqd")).length),p0);
        let _2030_v5;
        _2030_v5 = _dafny.Map.Empty.slice().updateUnsafe(p0,_module.__default.fm45(globalState));
        let _2031_v6;
        _2031_v6 = _dafny.MultiSet.fromElements(true, _2028_v3.f25, _2025_v0);
        let _2032_v7;
        _2032_v7 = _module.D17.create_DC39(p0, (_2029_v4).Merge(_module.__default.fm54(globalState)), !((_2028_v3).f26), _2030_v5, new BigNumber((_2031_v6).cardinality()));
        _2032_v7 = _2032_v7;
        let _2033_v8;
        _2033_v8 = _dafny.Set.fromElements(new BigNumber(-285));
        let _2034_v9;
        _2034_v9 = _dafny.Map.Empty.slice().updateUnsafe(_2033_v8,!(_2028_v3.f25));
        let _2035_v10;
        let _nw337 = Array((new BigNumber(11)).toNumber());
        _nw337[(_dafny.ZERO).toNumber()] = (_2025_v0) || ((_2028_v3).f26);
        _nw337[(_dafny.ONE).toNumber()] = false;
        _nw337[(new BigNumber(2)).toNumber()] = !(_dafny.Map.Empty.slice().updateUnsafe(_2033_v8,p0)).contains(_2033_v8);
        _nw337[(new BigNumber(3)).toNumber()] = (_2028_v3).f26;
        _nw337[(new BigNumber(4)).toNumber()] = false;
        _nw337[(new BigNumber(5)).toNumber()] = _2028_v3.f25;
        _nw337[(new BigNumber(6)).toNumber()] = (((_2034_v9).contains(_2033_v8)) ? ((_2034_v9).get(_2033_v8)) : (_2025_v0));
        _nw337[(new BigNumber(7)).toNumber()] = _2028_v3.f25;
        _nw337[(new BigNumber(8)).toNumber()] = _2028_v3.f25;
        _nw337[(new BigNumber(9)).toNumber()] = (_dafny.Set.fromElements(p0)).IsProperSubsetOf(_2033_v8);
        _nw337[(new BigNumber(10)).toNumber()] = _2028_v3.f25;
        _2035_v10 = _nw337;
        _2035_v10 = _2035_v10;
        let _2036_v11;
        let _nw338 = new _module.C0();
        _nw338.__ctor((_2028_v3).f26);
        _2036_v11 = _nw338;
        let _2037_v12;
        _2037_v12 = _dafny.Map.Empty.slice().updateUnsafe(_2036_v11,_2028_v3.f25);
        let _2038_v13;
        _2038_v13 = _dafny.Map.Empty.slice().updateUnsafe((_2037_v12).Merge(_2037_v12),_2036_v11.f27);
        let _2039_v14;
        _2039_v14 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2037_v12);
        let _2040_v15;
        _2040_v15 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2028_v3.f25);
        let _2041_v16;
        _2041_v16 = _dafny.Seq.of(new BigNumber((_2026_v1).length), new BigNumber((_2040_v15).length));
        let _2042_v17;
        _2042_v17 = _dafny.Map.Empty.slice().updateUnsafe(_2026_v1,_2041_v16);
        let _2043_v18;
        _2043_v18 = _module.D21.create_DC52(_2042_v17);
        let _pat_let_tv38 = _2042_v17;
        let _2044_v19;
        _2044_v19 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let28_0) {
          return function (_2045_dt__update__tmp_h0) {
            return function (_pat_let29_0) {
              return function (_2046_dt__update_hcf71_h0) {
                return _module.D21.create_DC52(_2046_dt__update_hcf71_h0);
              }(_pat_let29_0);
            }(_pat_let_tv38);
          }(_pat_let28_0);
        }(_2043_v18),_2036_v11);
        let _2047_v20;
        _2047_v20 = _dafny.Set.fromElements((((_2044_v19).contains(_2043_v18)) ? ((_2044_v19).get(_2043_v18)) : (_2036_v11)), _2036_v11);
        let _2048_v21;
        _2048_v21 = _dafny.Map.Empty.slice().updateUnsafe(_2036_v11.f27,_2047_v20);
        _2038_v13 = (_2038_v13).update((_2037_v12).Merge((((_2039_v14).contains(p0)) ? ((_2039_v14).get(p0)) : (_2037_v12))), ((((_2048_v21).contains(_2036_v11.f27)) ? ((_2048_v21).get(_2036_v11.f27)) : (_2047_v20))).IsSubsetOf(_2047_v20));
        let _2049_v22;
        let _nw339 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
        _2049_v22 = _nw339;
        let _2050_v23;
        _2050_v23 = _dafny.Seq.of(_2049_v22);
        let _2051_v24;
        _2051_v24 = _module.D6.create_DC12(_2050_v23);
        let _2052_v25;
        _2052_v25 = _module.D6.create_DC14(_2051_v24);
        let _2053_v26;
        _2053_v26 = _module.D0.create_DC0(_2028_v3.f25);
        let _2054_v27;
        _2054_v27 = _dafny.Map.Empty.slice().updateUnsafe(_2053_v26,p0);
        let _2055_v28;
        _2055_v28 = new _dafny.CodePoint('n'.codePointAt(0));
        let _2056_v29;
        _2056_v29 = _dafny.Set.fromElements(_2055_v28);
        let _2057_v30;
        _2057_v30 = _dafny.MultiSet.fromElements(_2056_v29, _dafny.Set.fromElements(new _dafny.CodePoint('x'.codePointAt(0))), _dafny.Set.fromElements(_2055_v28));
        let _2058_v31;
        _2058_v31 = _dafny.MultiSet.fromElements(_2035_v10);
        let _rhs327 = _dafny.Seq.contains(_dafny.Seq.of(new BigNumber(970), p0, new BigNumber((_2054_v27).length), p0, (_dafny.ZERO).minus(_module.__default.fm0(p0, (_dafny.ZERO).minus(p0), p0, (_this).fm7(false, p0, _2028_v3.f25, globalState), globalState))), (p0).plus(new BigNumber(691)));
        let _rhs328 = (_dafny.ZERO).minus((((_2057_v30).contains(_2056_v29)) ? ((_2057_v30).get(_2056_v29)) : ((p0).minus(p0))));
        let _rhs329 = _2052_v25;
        let _rhs330 = p0;
        let _rhs331 = (((_2058_v31).contains(_2035_v10)) ? ((_2058_v31).get(_2035_v10)) : (new BigNumber(384)));
        let _lhs307 = globalState;
        let _lhs308 = globalState;
        let _lhs309 = globalState;
        let _lhs310 = globalState;
        _lhs307.f21 = _rhs327;
        _lhs308.f16 = _rhs328;
        _2052_v25 = _rhs329;
        _lhs309.f7 = _rhs330;
        _lhs310.f7 = _rhs331;
        let _2059_v32;
        _2059_v32 = _dafny.Seq.UnicodeFromString("as");
        (globalState).f0 = ((!(!(_2028_v3.f25) || (_2028_v3.f25))) ? (_dafny.Seq.IsPrefixOf(_2059_v32, _2059_v32)) : ((_2028_v3).f26));
      } else {
        let _2060_v33;
        let _init61 = ((_2061_v2) => function (_2062_i0) {
          return _2061_v2;
        })(_2027_v2);
        let _nw340 = Array((new BigNumber(2)).toNumber());
        for (let _i0_61 = 0; _i0_61 < new BigNumber(_nw340.length); _i0_61++) {
          _nw340[_i0_61] = _init61(new BigNumber(_i0_61));
        }
        _2060_v33 = _nw340;
        _2060_v33 = _2060_v33;
        let _2063_v34;
        let _nw341 = Array((new BigNumber(11)).toNumber()).fill(_module.D12.Default());
        _2063_v34 = _nw341;
        let _2064_v35;
        _2064_v35 = _module.D12.create_DC23(p0);
        let _2065_v36;
        _2065_v36 = _module.D12.create_DC24(_2064_v35);
        let _index337 = _module.__default.safeIndex(new BigNumber(627), new BigNumber((_2063_v34).length));
        (_2063_v34)[_index337] = _2065_v36;
        let _pat_let_tv39 = _2064_v35;
        let _index338 = _module.__default.safeIndex(new BigNumber(627), new BigNumber((_2063_v34).length));
        (_2063_v34)[_index338] = function (_pat_let30_0) {
          return function (_2066_dt__update__tmp_h1) {
            return function (_pat_let31_0) {
              return function (_2067_dt__update_hcf31_h0) {
                return _module.D12.create_DC24(_2067_dt__update_hcf31_h0);
              }(_pat_let31_0);
            }(_pat_let_tv39);
          }(_pat_let30_0);
        }(_module.D12.create_DC24(_2064_v35));
        let _2068_v37;
        let _nw342 = new _module.C6();
        _nw342.__ctor();
        _2068_v37 = _nw342;
        let _2069_v38;
        _2069_v38 = _dafny.Set.fromElements(_2028_v3.f25, (_2028_v3).f26);
        (_2028_v3).m5((_2069_v38).IsDisjointFrom(_2069_v38), globalState);
        let _2070_v39;
        _2070_v39 = _dafny.Map.Empty.slice().updateUnsafe((_2028_v3).f26,(_2028_v3).f26);
        (globalState).f21 = _module.__default.fm2(_2070_v39, globalState);
      }
      let _2071_v40;
      _2071_v40 = new _dafny.CodePoint('c'.codePointAt(0));
      let _2072_v41;
      _2072_v41 = _dafny.Map.Empty.slice().updateUnsafe(_2071_v40,!(_2025_v0));
      _2072_v41 = (_2072_v41).update(_2071_v40, _2025_v0);
      if ((p0).isEqualTo(new BigNumber(306))) {
        (globalState).f16 = new BigNumber(836);
        let _2073_v42;
        _2073_v42 = _dafny.Map.Empty.slice().updateUnsafe(_2027_v2,_2028_v3.f25);
        let _2074_v43;
        _2074_v43 = (_dafny.Map.Empty.slice().updateUnsafe(_2027_v2,(_2028_v3).f26)).Merge(_2073_v42);
        let _2075_v44;
        _2075_v44 = _module.D12.create_DC23(p0);
        let _2076_v45;
        _2076_v45 = _module.D12.create_DC24(_2075_v44);
        let _rhs332 = _2074_v43;
        let _rhs333 = _module.__default.fm58(p0, globalState);
        _2074_v43 = _rhs332;
        _2076_v45 = _rhs333;
        let _2077_v46;
        let _nw343 = Array((new BigNumber(16)).toNumber()).fill(false);
        _2077_v46 = _nw343;
        _2077_v46 = _2077_v46;
        let _2078_v47;
        let _nw344 = new _module.C3();
        _nw344.__ctor();
        _2078_v47 = _nw344;
        let _2079_v48;
        _2079_v48 = _dafny.Seq.of(_2078_v47);
        let _2080_v49;
        _2080_v49 = _dafny.Seq.UnicodeFromString("ggly");
        let _2081_v50;
        _2081_v50 = _module.D24.create_DC63(_2078_v47);
        let _2082_v51;
        let _nw345 = Array((new BigNumber(18)).toNumber());
        _nw345[(_dafny.ZERO).toNumber()] = (_2079_v48)[_module.__default.safeIndex(new BigNumber((_2080_v49).length), new BigNumber((_2079_v48).length))];
        _nw345[(_dafny.ONE).toNumber()] = _2078_v47;
        _nw345[(new BigNumber(2)).toNumber()] = _2078_v47;
        _nw345[(new BigNumber(3)).toNumber()] = _2078_v47;
        _nw345[(new BigNumber(4)).toNumber()] = _2078_v47;
        _nw345[(new BigNumber(5)).toNumber()] = _2078_v47;
        _nw345[(new BigNumber(6)).toNumber()] = _2078_v47;
        _nw345[(new BigNumber(7)).toNumber()] = (_2081_v50).dtor_cf90;
        _nw345[(new BigNumber(8)).toNumber()] = _2078_v47;
        _nw345[(new BigNumber(9)).toNumber()] = ((_2028_v3.f25) ? (_2078_v47) : (_2078_v47));
        _nw345[(new BigNumber(10)).toNumber()] = _2078_v47;
        _nw345[(new BigNumber(11)).toNumber()] = (_2081_v50).dtor_cf90;
        _nw345[(new BigNumber(12)).toNumber()] = _2078_v47;
        _nw345[(new BigNumber(13)).toNumber()] = _2078_v47;
        _nw345[(new BigNumber(14)).toNumber()] = _2078_v47;
        _nw345[(new BigNumber(15)).toNumber()] = _2078_v47;
        _nw345[(new BigNumber(16)).toNumber()] = _2078_v47;
        _nw345[(new BigNumber(17)).toNumber()] = ((!(_2028_v3.f25)) ? (_2078_v47) : (_2078_v47));
        _2082_v51 = _nw345;
        let _index339 = _module.__default.safeIndex(new BigNumber(741), new BigNumber((_2082_v51).length));
        (_2082_v51)[_index339] = _2078_v47;
        let _index340 = _module.__default.safeIndex(new BigNumber(741), new BigNumber((_2082_v51).length));
        (_2082_v51)[_index340] = _2078_v47;
        if (!(!(false))) {
          let _index341 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_2077_v46).length));
          (_2077_v46)[_index341] = !(_2028_v3.f25) || ((_2028_v3).f26);
          let _index342 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_2077_v46).length));
          (_2077_v46)[_index342] = !(!(_2028_v3.f25));
          (globalState).f0 = (_dafny.Seq.IsPrefixOf(_2080_v49, _2080_v49)) === ((_2028_v3).f26);
          let _2083_v52;
          _2083_v52 = _dafny.Map.Empty.slice().updateUnsafe(_2025_v0,p0);
          let _2084_v53;
          _2084_v53 = _dafny.Set.fromElements(new BigNumber((_2080_v49).length));
          let _2085_v54;
          _2085_v54 = _dafny.Set.fromElements(_dafny.Set.fromElements(p0, p0));
          let _rhs334 = ((!(_2084_v53).equals(_2084_v53)) ? (_2026_v1) : (_2026_v1));
          let _rhs335 = _module.__default.safeModuloInt(p0, _module.__default.fm0(p0, p0, p0, (_2028_v3).f26, globalState));
          let _rhs336 = _module.__default.fm35(_2025_v0, new BigNumber((_2085_v54).length), globalState);
          let _rhs337 = p0;
          let _rhs338 = _2080_v49;
          let _lhs311 = globalState;
          let _lhs312 = globalState;
          let _lhs313 = globalState;
          _2026_v1 = _rhs334;
          _lhs311.f16 = _rhs335;
          _2083_v52 = _rhs336;
          _lhs312.f16 = _rhs337;
          _lhs313.f8 = _rhs338;
          (globalState).f7 = (new BigNumber(179)).plus(p0);
          let _2086_v55;
          let _init62 = ((_2087_v49) => function (_2088_i1) {
            return _dafny.Seq.of(_2087_v49, _2087_v49);
          })(_2080_v49);
          let _nw346 = Array((new BigNumber(27)).toNumber());
          for (let _i0_62 = 0; _i0_62 < new BigNumber(_nw346.length); _i0_62++) {
            _nw346[_i0_62] = _init62(new BigNumber(_i0_62));
          }
          _2086_v55 = _nw346;
          let _index343 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_2086_v55).length));
          (_2086_v55)[_index343] = _dafny.Seq.of(_2080_v49);
          let _2089_v56;
          _2089_v56 = _dafny.Seq.of(_2080_v49, _2080_v49, _2080_v49, _2080_v49);
          let _index344 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_2086_v55).length));
          (_2086_v55)[_index344] = _dafny.Seq.Concat(_2089_v56, _module.__default.fm59((_2028_v3).f26, globalState));
        } else {
          let _2090_v57;
          _2090_v57 = _dafny.MultiSet.fromElements(_2025_v0);
          let _2091_v58;
          let _nw347 = Array((_dafny.ONE).toNumber());
          _nw347[(_dafny.ZERO).toNumber()] = p0;
          _2091_v58 = _nw347;
          let _2092_v59;
          _2092_v59 = _module.D21.create_DC53(_2077_v46, false, ((_2090_v57).update(_2025_v0, _module.__default.abs(new BigNumber((_2080_v49).length)))).Difference(_2090_v57), _2091_v58);
          _2092_v59 = ((!((_2028_v3).f26) || (_2028_v3.f25)) ? (_2092_v59) : (_2092_v59));
          (globalState).f16 = p0;
          let _2093_v60;
          let _nw348 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2093_v60 = _nw348;
          _2093_v60 = _2093_v60;
          (globalState).f8 = _dafny.Seq.UnicodeFromString("ydtter");
          let _2094_v61;
          let _nw349 = new _module.C0();
          _nw349.__ctor(_2025_v0);
          _2094_v61 = _nw349;
        }
      } else {
        let _2095_v62;
        let _init63 = ((_2096_p0) => function (_2097_i2) {
          return (_2097_i2).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2096_p0,_2096_p0)).length));
        })(p0);
        let _nw350 = Array((new BigNumber(7)).toNumber());
        for (let _i0_63 = 0; _i0_63 < new BigNumber(_nw350.length); _i0_63++) {
          _nw350[_i0_63] = _init63(new BigNumber(_i0_63));
        }
        _2095_v62 = _nw350;
        let _2098_v63;
        let _out8;
        _out8 = (_this).m4(new BigNumber((_dafny.MultiSet.fromElements(_2095_v62, _2095_v62)).cardinality()), _2025_v0, new BigNumber(934), globalState);
        _2098_v63 = _out8;
        let _2099_v64;
        _2099_v64 = _dafny.Seq.of(_2095_v62);
        r0 = (_2099_v64)[_module.__default.safeIndex(p0, new BigNumber((_2099_v64).length))];
        let _2100_v65;
        _2100_v65 = _dafny.Map.Empty.slice().updateUnsafe(p0,new _dafny.CodePoint('v'.codePointAt(0)));
        let _2101_v66;
        _2101_v66 = _dafny.Map.Empty.slice().updateUnsafe(_2028_v3.f25,_2028_v3.f25);
        let _2102_v67;
        _2102_v67 = _dafny.Map.Empty.slice().updateUnsafe((_2100_v65).Merge(_2100_v65),_2101_v66);
        _2102_v67 = _module.__default.fm60(globalState);
        let _2103_v68;
        _2103_v68 = _dafny.Set.fromElements(_2028_v3.f25);
        let _2104_v69;
        _2104_v69 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_2103_v68).length));
        let _2105_v70;
        _2105_v70 = _dafny.Seq.of(_2104_v69);
        let _rhs339 = (_2028_v3).f26;
        let _rhs340 = (p0).minus(p0);
        let _rhs341 = (new BigNumber((_2105_v70).length)).minus(p0);
        let _lhs314 = _2028_v3;
        let _lhs315 = globalState;
        let _lhs316 = globalState;
        _lhs314.f25 = _rhs339;
        _lhs315.f16 = _rhs340;
        _lhs316.f16 = _rhs341;
        let _index345 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_2095_v62).length));
        (_2095_v62)[_index345] = new BigNumber(((_2027_v2).Intersect(_2027_v2)).cardinality());
        let _index346 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_2095_v62).length));
        (_2095_v62)[_index346] = p0;
      }
      let _2106_v71;
      _2106_v71 = _dafny.MultiSet.fromElements(_2071_v40, _2071_v40, _module.__default.fm15(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-90)), function (_2107_i3) {
        return new BigNumber(214);
      }), p0, globalState));
      let _2108_v72;
      _2108_v72 = _dafny.Seq.UnicodeFromString("h");
      let _2109_v73;
      _2109_v73 = _dafny.MultiSet.fromElements((_2028_v3).f26);
      let _rhs342 = ((_dafny.MultiSet.fromElements((_2108_v72)[_module.__default.safeIndex(new BigNumber((_2109_v73).cardinality()), new BigNumber((_2108_v72).length))])).update(_2071_v40, _module.__default.abs(p0))).Union((_2106_v71).update(_2071_v40, _module.__default.abs(new BigNumber((_module.__default.fm1(new BigNumber(502), globalState)).length))));
      let _rhs343 = _2028_v3.f25;
      let _rhs344 = ((!((new BigNumber(706)).isLessThan(new BigNumber(212)))) ? (new _dafny.CodePoint('x'.codePointAt(0))) : (_2071_v40));
      let _rhs345 = p0;
      let _lhs317 = _2028_v3;
      let _lhs318 = globalState;
      _2106_v71 = _rhs342;
      _lhs317.f25 = _rhs343;
      _2071_v40 = _rhs344;
      _lhs318.f7 = _rhs345;
      let _2110_v74;
      let _nw351 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
      _2110_v74 = _nw351;
      for (const _guard_loop_11 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2110_v74).length))) {
        let _2111_i4 = _guard_loop_11;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2111_i4)) && ((_2111_i4).isLessThan(new BigNumber((_2110_v74).length))))) {
          (_2110_v74)[(_2111_i4)] = (_2111_i4).minus(p0);
        }
      }
      r0 = _2110_v74;
      let _2112_v75;
      _2112_v75 = _dafny.Set.fromElements(_2110_v74, _2110_v74);
      r1 = !(!(p0).isEqualTo(new BigNumber((_2112_v75).length)));
      return [r0, r1];
    }
    m4(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let _2113_v0;
      _2113_v0 = _dafny.Map.Empty.slice().updateUnsafe(false,p1);
      let _2114_v1;
      _2114_v1 = _dafny.Set.fromElements(p1);
      let _2115_v2;
      _2115_v2 = new _dafny.CodePoint('d'.codePointAt(0));
      let _2116_v3;
      _2116_v3 = _dafny.Seq.UnicodeFromString("yjscrou");
      let _2117_v4;
      let _nw352 = Array((new BigNumber(16)).toNumber());
      _nw352[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-831)), function (_2118_i0) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      })).length);
      _nw352[(_dafny.ONE).toNumber()] = p2;
      _nw352[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(p2);
      _nw352[(new BigNumber(3)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(false, true)).length), _module.__default.fm0(p2, p2, p2, p1, globalState));
      _nw352[(new BigNumber(4)).toNumber()] = (new BigNumber((_2113_v0).length)).multipliedBy(p2);
      _nw352[(new BigNumber(5)).toNumber()] = p0;
      _nw352[(new BigNumber(6)).toNumber()] = p2;
      _nw352[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(p0);
      _nw352[(new BigNumber(8)).toNumber()] = new BigNumber(((_dafny.Set.fromElements(false)).Union(_2114_v1)).length);
      _nw352[(new BigNumber(9)).toNumber()] = new BigNumber(359);
      _nw352[(new BigNumber(10)).toNumber()] = p2;
      _nw352[(new BigNumber(11)).toNumber()] = (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2115_v2,_2115_v2)).length)).minus(new BigNumber(303));
      _nw352[(new BigNumber(12)).toNumber()] = p0;
      _nw352[(new BigNumber(13)).toNumber()] = p0;
      _nw352[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_2116_v3).length));
      _nw352[(new BigNumber(15)).toNumber()] = p0;
      _2117_v4 = _nw352;
      let _index347 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length));
      (_2117_v4)[_index347] = p0;
      let _index348 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length));
      (_2117_v4)[_index348] = _module.__default.fm0(p2, p2, p0, ((p1) ? (p1) : (p1)), globalState);
      let _2119_v5;
      _2119_v5 = _dafny.Seq.of(_2117_v4, _2117_v4, _2117_v4, _2117_v4, _2117_v4);
      let _2120_v6;
      _2120_v6 = _module.D6.create_DC12(_2119_v5);
      let _pat_let_tv40 = _2117_v4;
      _2119_v5 = ((p1) ? (_dafny.Seq.Concat(_2119_v5, _2119_v5)) : ((function (_pat_let32_0) {
        return function (_2121_dt__update__tmp_h0) {
          return function (_pat_let33_0) {
            return function (_2122_dt__update_hcf18_h0) {
              return _module.D6.create_DC12(_2122_dt__update_hcf18_h0);
            }(_pat_let33_0);
          }(_dafny.Seq.of(_pat_let_tv40));
        }(_pat_let32_0);
      }(_2120_v6)).dtor_cf18));
      (globalState).f7 = ((p1) ? ((_2117_v4)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length))]) : (p0));
      let _2123_v7;
      let _nw353 = new _module.C6();
      _nw353.__ctor();
      _2123_v7 = _nw353;
      _2123_v7 = _2123_v7;
      for (const _guard_loop_12 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2117_v4).length))) {
        let _2124_i1 = _guard_loop_12;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2124_i1)) && ((_2124_i1).isLessThan(new BigNumber((_2117_v4).length))))) {
          (_2117_v4)[(_2124_i1)] = (_2124_i1).plus(new BigNumber((_2116_v3).length));
        }
      }
      let _pat_let_tv41 = _2119_v5;
      let _source26 = function (_pat_let34_0) {
        return function (_2125_dt__update__tmp_h1) {
          return function (_pat_let35_0) {
            return function (_2126_dt__update_hcf18_h1) {
              return _module.D6.create_DC12(_2126_dt__update_hcf18_h1);
            }(_pat_let35_0);
          }(_pat_let_tv41);
        }(_pat_let34_0);
      }(_2120_v6);
      if (_source26.is_DC13) {
        let _2127___mcc_h0 = (_source26).cf19;
        let _2128___mcc_h1 = (_source26).cf20;
        let _2129_cf20 = _2128___mcc_h1;
        let _2130_cf19 = _2127___mcc_h0;
        let _index349 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length));
        (_2117_v4)[_index349] = (_dafny.ZERO).minus((_dafny.ZERO).minus(p0));
        (globalState).f2 = false;
        let _2131_v8;
        _2131_v8 = _dafny.Seq.of(p1, p1, true);
        let _2132_v9;
        _2132_v9 = _dafny.Seq.of(_2131_v8);
        let _2133_v10;
        _2133_v10 = _dafny.Seq.of(new BigNumber(((_2132_v9)[_module.__default.safeIndex(_2130_cf19, new BigNumber((_2132_v9).length))]).length), new BigNumber(-542), (_2117_v4)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length))]);
        let _2134_v11;
        _2134_v11 = _dafny.Map.Empty.slice().updateUnsafe((_2133_v10)[_module.__default.safeIndex(new BigNumber((_2116_v3).length), new BigNumber((_2133_v10).length))],p2);
        let _2135_v12;
        let _init64 = ((_2136_v2) => function (_2137_i2) {
          return _2136_v2;
        })(_2115_v2);
        let _nw354 = Array((new BigNumber(14)).toNumber());
        for (let _i0_64 = 0; _i0_64 < new BigNumber(_nw354.length); _i0_64++) {
          _nw354[_i0_64] = _init64(new BigNumber(_i0_64));
        }
        _2135_v12 = _nw354;
        let _2138_v13;
        _2138_v13 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2135_v12);
        _2134_v11 = (_2134_v11).update(new BigNumber((_2138_v13).length), p2);
        let _2139_v14;
        _2139_v14 = _dafny.MultiSet.fromElements(p1, p1, p1);
        let _2140_v15;
        _2140_v15 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_2139_v14).Union(_2139_v14));
        _2140_v15 = (_2140_v15).update(((p1) ? (true) : (p1)), _module.__default.fm52(p1, _module.__default.fm0(new BigNumber((_2116_v3).length), (_dafny.ZERO).minus(_2130_cf19), p2, (_2131_v8)[_module.__default.safeIndex(_2130_cf19, new BigNumber((_2131_v8).length))], globalState), globalState));
      } else if (_source26.is_DC12) {
        let _2141___mcc_h2 = (_source26).cf18;
        let _2142_cf18 = _2141___mcc_h2;
        (globalState).f2 = ((p1) ? (p1) : (p1));
        let _2143_v16;
        let _nw355 = Array((new BigNumber(25)).toNumber());
        _nw355[(_dafny.ZERO).toNumber()] = p1;
        _nw355[(_dafny.ONE).toNumber()] = p1;
        _nw355[(new BigNumber(2)).toNumber()] = p1;
        _nw355[(new BigNumber(3)).toNumber()] = p1;
        _nw355[(new BigNumber(4)).toNumber()] = (p2).isLessThanOrEqualTo((_2117_v4)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length))]);
        _nw355[(new BigNumber(5)).toNumber()] = (p1) || (p1);
        _nw355[(new BigNumber(6)).toNumber()] = _module.__default.fm2(_2113_v0, globalState);
        _nw355[(new BigNumber(7)).toNumber()] = p1;
        _nw355[(new BigNumber(8)).toNumber()] = p1;
        _nw355[(new BigNumber(9)).toNumber()] = p1;
        _nw355[(new BigNumber(10)).toNumber()] = p1;
        _nw355[(new BigNumber(11)).toNumber()] = p1;
        _nw355[(new BigNumber(12)).toNumber()] = true;
        _nw355[(new BigNumber(13)).toNumber()] = p1;
        _nw355[(new BigNumber(14)).toNumber()] = p1;
        _nw355[(new BigNumber(15)).toNumber()] = p1;
        _nw355[(new BigNumber(16)).toNumber()] = p1;
        _nw355[(new BigNumber(17)).toNumber()] = p1;
        _nw355[(new BigNumber(18)).toNumber()] = p1;
        _nw355[(new BigNumber(19)).toNumber()] = !(p1);
        _nw355[(new BigNumber(20)).toNumber()] = false;
        _nw355[(new BigNumber(21)).toNumber()] = p1;
        _nw355[(new BigNumber(22)).toNumber()] = p1;
        _nw355[(new BigNumber(23)).toNumber()] = p1;
        _nw355[(new BigNumber(24)).toNumber()] = p1;
        _2143_v16 = _nw355;
        _2143_v16 = _2143_v16;
        let _2144_v17;
        _2144_v17 = _dafny.Seq.of(_2116_v3, _2116_v3);
        let _2145_v18;
        _2145_v18 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(102),p1);
        let _2146_v19;
        _2146_v19 = _dafny.Map.Empty.slice().updateUnsafe((_2117_v4)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length))],_2145_v18);
        let _2147_v20;
        _2147_v20 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber(((((_2146_v19).contains(new BigNumber(170))) ? ((_2146_v19).get(new BigNumber(170))) : (_2145_v18))).length))).length),(_2117_v4)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length))]);
        let _2148_v21;
        _2148_v21 = _module.D3.create_DC6(_2116_v3, _2115_v2, p1);
        let _2149_v22;
        let _nw356 = Array((new BigNumber(26)).toNumber());
        _nw356[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat((_2144_v17)[_module.__default.safeIndex(new BigNumber(331), new BigNumber((_2144_v17).length))], _dafny.Seq.UnicodeFromString("hesfue"));
        _nw356[(_dafny.ONE).toNumber()] = ((p1) ? (_2116_v3) : (_dafny.Seq.update(_2116_v3, _module.__default.safeIndex(p0, new BigNumber((_2116_v3).length)), _2115_v2)));
        _nw356[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-666)), ((_2150_v2) => function (_2151_i3) {
          return _2150_v2;
        })(_2115_v2));
        _nw356[(new BigNumber(3)).toNumber()] = _2116_v3;
        _nw356[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm1(new BigNumber((_2116_v3).length), globalState), _2116_v3);
        _nw356[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_2116_v3, _2116_v3);
        _nw356[(new BigNumber(6)).toNumber()] = _dafny.Seq.UnicodeFromString("thegxd");
        _nw356[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_2116_v3, _2116_v3);
        _nw356[(new BigNumber(8)).toNumber()] = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("npsshuqfe"), _module.__default.safeIndex(new BigNumber((_2147_v20).length), new BigNumber((_dafny.Seq.UnicodeFromString("npsshuqfe")).length)), _2115_v2);
        _nw356[(new BigNumber(9)).toNumber()] = _2116_v3;
        _nw356[(new BigNumber(10)).toNumber()] = _2116_v3;
        _nw356[(new BigNumber(11)).toNumber()] = _dafny.Seq.UnicodeFromString("cshpf");
        _nw356[(new BigNumber(12)).toNumber()] = _2116_v3;
        _nw356[(new BigNumber(13)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-731)), ((_2152_v2) => function (_2153_i4) {
          return _2152_v2;
        })(_2115_v2)), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-731)), ((_2154_v2) => function (_2155_i4) {
          return _2154_v2;
        })(_2115_v2))).length)), _2115_v2);
        _nw356[(new BigNumber(14)).toNumber()] = _2116_v3;
        _nw356[(new BigNumber(15)).toNumber()] = _2116_v3;
        _nw356[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_2116_v3, _2116_v3);
        _nw356[(new BigNumber(17)).toNumber()] = (_2144_v17)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("kiiymoo")).length), new BigNumber((_2144_v17).length))];
        _nw356[(new BigNumber(18)).toNumber()] = _2116_v3;
        _nw356[(new BigNumber(19)).toNumber()] = _dafny.Seq.UnicodeFromString("svnce");
        _nw356[(new BigNumber(20)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_module.__default.fm1(p0, globalState), _module.__default.safeIndex(p0, new BigNumber((_module.__default.fm1(p0, globalState)).length)), new _dafny.CodePoint('d'.codePointAt(0))), _dafny.Seq.UnicodeFromString("xkvbkkxeo"));
        _nw356[(new BigNumber(21)).toNumber()] = _2116_v3;
        _nw356[(new BigNumber(22)).toNumber()] = (_2148_v21).dtor_cf6;
        _nw356[(new BigNumber(23)).toNumber()] = _dafny.Seq.update(_2116_v3, _module.__default.safeIndex(p0, new BigNumber((_2116_v3).length)), _2115_v2);
        _nw356[(new BigNumber(24)).toNumber()] = _2116_v3;
        _nw356[(new BigNumber(25)).toNumber()] = _2116_v3;
        _2149_v22 = _nw356;
        let _index350 = _module.__default.safeIndex(new BigNumber(809), new BigNumber((_2149_v22).length));
        (_2149_v22)[_index350] = _2116_v3;
        let _2156_v23;
        _2156_v23 = _dafny.Map.Empty.slice().updateUnsafe((p0).multipliedBy((_2117_v4)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length))]),_dafny.Seq.Concat(_2116_v3, _dafny.Seq.update(_dafny.Seq.update(_2116_v3, _module.__default.safeIndex(new BigNumber(832), new BigNumber((_2116_v3).length)), _2115_v2), _module.__default.safeIndex(p2, new BigNumber((_dafny.Seq.update(_2116_v3, _module.__default.safeIndex(new BigNumber(832), new BigNumber((_2116_v3).length)), _2115_v2)).length)), _2115_v2)));
        let _index351 = _module.__default.safeIndex(new BigNumber(809), new BigNumber((_2149_v22).length));
        (_2149_v22)[_index351] = (((_2156_v23).contains((_2117_v4)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length))])) ? ((_2156_v23).get((_2117_v4)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length))])) : (_2116_v3));
        (globalState).f2 = p1;
      } else {
        let _2157___mcc_h3 = (_source26).cf21;
        let _2158_cf21 = _2157___mcc_h3;
        if (!(p1)) {
          (globalState).f20 = _dafny.Seq.Concat(_2116_v3, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("isojrib"), _2116_v3));
          (globalState).f16 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt(p2, new BigNumber(-215))));
          let _2159_v24;
          _2159_v24 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
          let _2160_v25;
          _2160_v25 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(p0, p0, (_dafny.ZERO).minus(p2), !(p1), globalState),_2116_v3);
          _2159_v24 = (_2159_v24).update(new BigNumber((_2160_v25).length), p1);
          let _2161_v26;
          _2161_v26 = _dafny.MultiSet.fromElements(p1, p1);
          let _2162_v27;
          _2162_v27 = _dafny.Seq.of(_2161_v26);
          let _2163_v28;
          _2163_v28 = _dafny.Map.Empty.slice().updateUnsafe(((_2162_v27)[_module.__default.safeIndex(p2, new BigNumber((_2162_v27).length))]).IsProperSubsetOf(_2161_v26),_dafny.Seq.of(p1, false));
          let _2164_v29;
          _2164_v29 = _dafny.Seq.of(p1, p1);
          _2163_v28 = (_2163_v28).update(p1, _2164_v29);
          let _2165_v30;
          _2165_v30 = _dafny.Map.Empty.slice().updateUnsafe(_2117_v4,_2123_v7);
          (globalState).f16 = _module.__default.fm0((((_2161_v26).contains(p1)) ? ((_2161_v26).get(p1)) : ((_2117_v4)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length))])), _module.__default.fm0(_module.__default.fm0(p2, p0, (_2117_v4)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length))], true, globalState), new BigNumber((_2159_v24).length), new BigNumber((_2165_v30).length), p1, globalState), _module.__default.fm0((_2117_v4)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length))], (_2117_v4)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length))], (_2117_v4)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length))], p1, globalState), p1, globalState);
        } else {
          let _2166_v31;
          _2166_v31 = _dafny.MultiSet.fromElements(p1, p1);
          let _2167_v32;
          _2167_v32 = _dafny.Map.Empty.slice().updateUnsafe(_2117_v4,!(_2166_v31).contains(p1));
          _2167_v32 = (_2167_v32).update(_2117_v4, p1);
          let _2168_v33;
          _2168_v33 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(267),p2);
          let _2169_v34;
          _2169_v34 = _dafny.Seq.of((((_2168_v33).contains(new BigNumber(((_2168_v33).update((_2117_v4)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length))], p0)).length))) ? ((_2168_v33).get(new BigNumber(((_2168_v33).update((_2117_v4)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length))], p0)).length))) : (p2)), (_2117_v4)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length))], (_2117_v4)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length))]);
          let _2170_v35;
          _2170_v35 = _dafny.Map.Empty.slice().updateUnsafe((_2169_v34)[_module.__default.safeIndex(new BigNumber((_2114_v1).length), new BigNumber((_2169_v34).length))],p1);
          _2170_v35 = (_2170_v35).update((_2117_v4)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length))], p1);
          let _2171_v36;
          let _nw357 = new _module.C4();
          _nw357.__ctor(p1, p1);
          _2171_v36 = _nw357;
          let _2172_v37;
          _2172_v37 = _dafny.Set.fromElements(_2171_v36);
          _2172_v37 = _2172_v37;
          (globalState).f16 = (_2169_v34)[_module.__default.safeIndex((p0).minus(p0), new BigNumber((_2169_v34).length))];
          let _2173_v38;
          let _nw358 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Map.Empty);
          _2173_v38 = _nw358;
          let _2174_v40;
          _2174_v40 = _dafny.Map.Empty.slice().updateUnsafe(_2115_v2,false);
          let _2175_v41;
          _2175_v41 = _module.D22.create_DC59(_module.D22.create_DC57((_2117_v4)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length))], new BigNumber(767), _2171_v36.f25));
          let _2176_v42;
          _2176_v42 = _dafny.Map.Empty.slice().updateUnsafe(function () {
            let _coll77 = new _dafny.Map();
            for (const _compr_77 of (_2174_v40).Keys.Elements) {
              let _2177_v39 = _compr_77;
              if ((_2174_v40).contains(_2177_v39)) {
                _coll77.push([_2177_v39,(_2171_v36).f26]);
              }
            }
            return _coll77;
          }(),_2175_v41);
          let _index352 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_2173_v38).length));
          (_2173_v38)[_index352] = (_2176_v42).Merge(_2176_v42);
          let _2178_v43;
          _2178_v43 = _dafny.Seq.of((_2171_v36).f26);
          let _2179_v44;
          _2179_v44 = _dafny.Map.Empty.slice().updateUnsafe((_2171_v36).f26,_2178_v43);
          let _2180_v45;
          _2180_v45 = _module.D16.create_DC36((_2171_v36).f26, (_2166_v31).Union(_dafny.MultiSet.FromArray((((_2179_v44).contains((_this).fm7(p1, new BigNumber((_2178_v43).length), _2171_v36.f25, globalState))) ? ((_2179_v44).get((_this).fm7(p1, new BigNumber((_2178_v43).length), _2171_v36.f25, globalState))) : (_2178_v43)))), _2171_v36.f25);
          let _2181_v46;
          let _nw359 = Array((new BigNumber(3)).toNumber()).fill(false);
          _2181_v46 = _nw359;
          let _2182_v47;
          _2182_v47 = _dafny.Seq.of(_2181_v46);
          let _2183_v48;
          let _nw360 = Array((new BigNumber(26)).toNumber());
          _nw360[(_dafny.ZERO).toNumber()] = _2181_v46;
          _nw360[(_dafny.ONE).toNumber()] = _2181_v46;
          _nw360[(new BigNumber(2)).toNumber()] = _2181_v46;
          _nw360[(new BigNumber(3)).toNumber()] = _2181_v46;
          _nw360[(new BigNumber(4)).toNumber()] = _2181_v46;
          _nw360[(new BigNumber(5)).toNumber()] = _2181_v46;
          _nw360[(new BigNumber(6)).toNumber()] = _2181_v46;
          _nw360[(new BigNumber(7)).toNumber()] = _2181_v46;
          _nw360[(new BigNumber(8)).toNumber()] = _2181_v46;
          _nw360[(new BigNumber(9)).toNumber()] = _2181_v46;
          _nw360[(new BigNumber(10)).toNumber()] = _2181_v46;
          _nw360[(new BigNumber(11)).toNumber()] = _2181_v46;
          _nw360[(new BigNumber(12)).toNumber()] = _2181_v46;
          _nw360[(new BigNumber(13)).toNumber()] = _2181_v46;
          _nw360[(new BigNumber(14)).toNumber()] = _2181_v46;
          _nw360[(new BigNumber(15)).toNumber()] = _2181_v46;
          _nw360[(new BigNumber(16)).toNumber()] = _2181_v46;
          _nw360[(new BigNumber(17)).toNumber()] = _2181_v46;
          _nw360[(new BigNumber(18)).toNumber()] = (_2182_v47)[_module.__default.safeIndex(p2, new BigNumber((_2182_v47).length))];
          _nw360[(new BigNumber(19)).toNumber()] = _2181_v46;
          _nw360[(new BigNumber(20)).toNumber()] = _2181_v46;
          _nw360[(new BigNumber(21)).toNumber()] = _2181_v46;
          _nw360[(new BigNumber(22)).toNumber()] = _2181_v46;
          _nw360[(new BigNumber(23)).toNumber()] = _2181_v46;
          _nw360[(new BigNumber(24)).toNumber()] = _2181_v46;
          _nw360[(new BigNumber(25)).toNumber()] = _2181_v46;
          _2183_v48 = _nw360;
          let _index353 = _module.__default.safeIndex(new BigNumber(558), new BigNumber((_2183_v48).length));
          (_2183_v48)[_index353] = _2181_v46;
          let _pat_let_tv42 = _2178_v43;
          let _pat_let_tv43 = _2117_v4;
          let _pat_let_tv44 = _2117_v4;
          let _pat_let_tv45 = _2178_v43;
          let _index354 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_2173_v38).length));
          let _index355 = _module.__default.safeIndex(new BigNumber(558), new BigNumber((_2183_v48).length));
          let _rhs346 = (_dafny.Map.Empty.slice().updateUnsafe(_2174_v40,_2175_v41)).Merge(_2176_v42);
          let _rhs347 = _2178_v43;
          let _rhs348 = (_2116_v3)[_module.__default.safeIndex((p0).minus(p2), new BigNumber((_2116_v3).length))];
          let _rhs349 = function (_pat_let36_0) {
            return function (_2184_dt__update__tmp_h2) {
              return function (_pat_let37_0) {
                return function (_2185_dt__update_hcf49_h0) {
                  return _module.D16.create_DC36((_2184_dt__update__tmp_h2).dtor_cf47, (_2184_dt__update__tmp_h2).dtor_cf48, _2185_dt__update_hcf49_h0);
                }(_pat_let37_0);
              }((_pat_let_tv42)[_module.__default.safeIndex((_pat_let_tv44)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_pat_let_tv43).length))], new BigNumber((_pat_let_tv45).length))]);
            }(_pat_let36_0);
          }(_module.D16.create_DC36(!((_2171_v36).f26), _2166_v31, _2171_v36.f25));
          let _rhs350 = _2181_v46;
          let _lhs319 = _2173_v38;
          let _lhs320 = _module.__default.safeIndex(new BigNumber(187), new BigNumber((_2173_v38).length));
          let _lhs321 = _2183_v48;
          let _lhs322 = _module.__default.safeIndex(new BigNumber(558), new BigNumber((_2183_v48).length));
          _lhs319[_lhs320] = _rhs346;
          _2178_v43 = _rhs347;
          _2115_v2 = _rhs348;
          _2180_v45 = _rhs349;
          _lhs321[_lhs322] = _rhs350;
        }
        let _index356 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length));
        (_2117_v4)[_index356] = _module.__default.safeModuloInt((_2117_v4)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length))], _module.__default.safeModuloInt(new BigNumber((function () {
          let _coll78 = new _dafny.Map();
          for (const _compr_78 of (_dafny.Seq.of(_2116_v3)).Elements) {
            let _2186_v49 = _compr_78;
            if (_dafny.Seq.contains(_dafny.Seq.of(_2116_v3), _2186_v49)) {
              _coll78.push([_2186_v49,_2115_v2]);
            }
          }
          return _coll78;
        }()).length), p2));
        let _2187_v50;
        _2187_v50 = _dafny.MultiSet.fromElements(_2117_v4, _2117_v4);
        let _index357 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length));
        (_2117_v4)[_index357] = (_dafny.ZERO).minus(((_2117_v4)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length))]).plus(new BigNumber(((_2187_v50).Difference(_2187_v50)).cardinality())));
        if ((_2123_v7).fm7(!(p1) || (!(!(p1))), p0, p1, globalState)) {
          let _2188_v51;
          let _nw361 = new _module.C3();
          _nw361.__ctor();
          _2188_v51 = _nw361;
          _2188_v51 = _2188_v51;
          let _index358 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length));
          (_2117_v4)[_index358] = p0;
          let _2189_v52;
          _2189_v52 = _dafny.Set.fromElements(p2, p2);
          let _2190_v53;
          _2190_v53 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_dafny.ZERO).minus((_2117_v4)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length))]));
          let _2191_v54;
          _2191_v54 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber((_2116_v3).length));
          let _2192_v55;
          _2192_v55 = _dafny.Seq.of((_2117_v4)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length))], new BigNumber((_2191_v54).length));
          let _2193_v56;
          _2193_v56 = _dafny.Map.Empty.slice().updateUnsafe((((_2190_v53).contains(p1)) ? ((_2190_v53).get(p1)) : ((_2192_v55)[_module.__default.safeIndex(new BigNumber((_2116_v3).length), new BigNumber((_2192_v55).length))])),!(p1));
          _2189_v52 = _module.__default.fm26(_2115_v2, _2193_v56, globalState);
          let _index359 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length));
          (_2117_v4)[_index359] = new BigNumber(((_module.__default.fm61(new BigNumber(748), _2115_v2, _2115_v2, globalState)).update(((_2113_v0).update(true, p1)).Merge((_dafny.Map.Empty.slice().updateUnsafe(p1,p1)).update(true, p1)), _module.__default.abs(p0))).cardinality());
          let _2194_v57;
          _2194_v57 = _dafny.MultiSet.fromElements(p0);
          let _2195_v58;
          _2195_v58 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2113_v0);
          let _2196_v59;
          _2196_v59 = _module.D23.create_DC60(_2195_v58);
          let _2197_v60;
          _2197_v60 = _dafny.Set.fromElements(_2196_v59, _2196_v59, _2196_v59);
          let _2198_v61;
          let _nw362 = new _module.C2();
          _nw362.__ctor(new BigNumber(((_2197_v60).Union(_2197_v60)).length));
          _2198_v61 = _nw362;
          let _rhs351 = (_2194_v57).Difference((_2194_v57).Union(_2194_v57));
          let _rhs352 = !(!(p1));
          let _rhs353 = (_2194_v57).Difference((_module.D25.create_DC67(_2194_v57)).dtor_cf98);
          let _rhs354 = _2198_v61;
          let _lhs323 = globalState;
          _2194_v57 = _rhs351;
          _lhs323.f21 = _rhs352;
          _2194_v57 = _rhs353;
          _2198_v61 = _rhs354;
        } else {
          _2115_v2 = _2115_v2;
          _2113_v0 = (_2113_v0).update(p1, p1);
          let _2199_v62;
          let _nw363 = Array((new BigNumber(21)).toNumber()).fill([]);
          _2199_v62 = _nw363;
          let _index360 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_2199_v62).length));
          (_2199_v62)[_index360] = _2117_v4;
          let _index361 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_2199_v62).length));
          (_2199_v62)[_index361] = _2117_v4;
          (globalState).f13 = p1;
          let _2200_v63;
          let _nw364 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Map.Empty);
          _2200_v63 = _nw364;
          _2200_v63 = _2200_v63;
        }
      }
      r0 = _dafny.Seq.Concat(((p1) ? (_2116_v3) : (_dafny.Seq.update(_2116_v3, _module.__default.safeIndex((_2117_v4)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_2117_v4).length))], new BigNumber((_2116_v3).length)), new _dafny.CodePoint('n'.codePointAt(0))))), _2116_v3);
      return r0;
    }
    get f23() {
      let _this = this;
      return _this._f23;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
