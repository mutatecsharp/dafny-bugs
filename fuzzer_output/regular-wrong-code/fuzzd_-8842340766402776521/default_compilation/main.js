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
      return (_module.D1.create_DC2(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(898))).length), new BigNumber((_dafny.Seq.UnicodeFromString("mcftfnqg")).length), new BigNumber(-391), new BigNumber((function () {
  let _coll0 = new _dafny.Map();
  for (const _compr_0 of _dafny.IntegerRange(new BigNumber(194), new BigNumber(317))) {
    let _0_v0 = _compr_0;
    if (((new BigNumber(194)).isLessThanOrEqualTo(_0_v0)) && ((_0_v0).isLessThan(new BigNumber(317)))) {
      _coll0.push([(_0_v0).minus(new BigNumber(-949)),_dafny.Set.fromElements(false)]);
    }
  }
  return _coll0;
}()).length))).dtor_cf3;
    };
    static fm1(p0, p1, p2, p3, globalState) {
      return (_dafny.Set.fromElements(false)).IsDisjointFrom((_dafny.Set.fromElements(true)).Difference(_dafny.Set.fromElements(true, false, false)));
    };
    static fm2(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("p")).length),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(false, true, !(!(false)), !(false))).length),false));
    };
    static fm3(p0, p1, p2, globalState) {
      if ((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(807)), function (_1_i0) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      })).length)).isEqualTo(new BigNumber((_dafny.Set.fromElements(!(!(!(false))), !(!(!(!(true)))), !(true), false)).length))) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('g'.codePointAt(0));
      }
    };
    static fm6(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber(759)),new BigNumber((_dafny.Set.fromElements(true, false)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(function () {
        let _coll1 = new _dafny.Set();
        for (const _compr_1 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-685)), function (_2_i0) {
          return new BigNumber(-994);
        })).length),true)).Keys.Elements) {
          let _3_v0 = _compr_1;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-685)), function (_2_i0) {
            return new BigNumber(-994);
          })).length),true)).contains(_3_v0)) {
            _coll1.add((_3_v0).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(79)), function (_4_i1) {
              return _dafny.Seq.of(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)));
            })).length)));
          }
        }
        return _coll1;
      }(),new BigNumber(930)));
    };
    static fm7(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("j"), _dafny.Seq.UnicodeFromString("iptpuc")), _dafny.Seq.Create(_module.__default.abs(new BigNumber(838)), function (_5_i0) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      }));
    };
    static fm8(p0, globalState) {
      let _source0 = _module.D1.create_DC2(new BigNumber(243), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("adlori"))).length),_module.D3.create_DC9(true))).length), new BigNumber((_dafny.Seq.UnicodeFromString("omdgnphda")).length), new BigNumber((_dafny.Seq.UnicodeFromString("gtoykysed")).length));
      if (_source0.is_DC2) {
        let _6___mcc_h0 = (_source0).cf2;
        let _7___mcc_h1 = (_source0).cf3;
        let _8___mcc_h2 = (_source0).cf4;
        let _9___mcc_h3 = (_source0).cf5;
        let _10_cf5 = _9___mcc_h3;
        let _11_cf4 = _8___mcc_h2;
        let _12_cf3 = _7___mcc_h1;
        let _13_cf2 = _6___mcc_h0;
        return _dafny.Seq.Concat(_dafny.Seq.of(!(false), true), _dafny.Seq.of(!(true)));
      } else if (_source0.is_DC1) {
        let _14___mcc_h4 = (_source0).cf1;
        let _15_cf1 = _14___mcc_h4;
        return _dafny.Seq.Concat(_dafny.Seq.of(false, false), _dafny.Seq.of(!(true), true));
      } else {
        let _16___mcc_h5 = (_source0).cf6;
        let _17_cf6 = _16___mcc_h5;
        return _dafny.Seq.of(true, false, false);
      }
    };
    static m0(p0, p1, globalState) {
      (globalState).f10 = (p1).multipliedBy((p1).multipliedBy(p1));
      let _18_v0;
      let _nw0 = new _module.C0();
      _nw0.__ctor(p1);
      _18_v0 = _nw0;
      let _19_v1;
      _19_v1 = _dafny.Map.Empty.slice().updateUnsafe(_18_v0,(_18_v0).f14);
      let _20_v2;
      _20_v2 = _dafny.Seq.of(p0);
      let _21_v3;
      _21_v3 = _dafny.Seq.UnicodeFromString("iaqmfjp");
      let _22_v4;
      _22_v4 = _dafny.MultiSet.fromElements((_20_v2)[_module.__default.safeIndex((_18_v0).f14, new BigNumber((_20_v2).length))], p0, p0);
      let _23_v5;
      let _nw1 = Array((new BigNumber(12)).toNumber());
      _nw1[(_dafny.ZERO).toNumber()] = (_19_v1).contains(_18_v0);
      _nw1[(_dafny.ONE).toNumber()] = (_20_v2)[_module.__default.safeIndex(new BigNumber((_21_v3).length), new BigNumber((_20_v2).length))];
      _nw1[(new BigNumber(2)).toNumber()] = p0;
      _nw1[(new BigNumber(3)).toNumber()] = !((p0) || (p0));
      _nw1[(new BigNumber(4)).toNumber()] = p0;
      _nw1[(new BigNumber(5)).toNumber()] = p0;
      _nw1[(new BigNumber(6)).toNumber()] = p0;
      _nw1[(new BigNumber(7)).toNumber()] = p0;
      _nw1[(new BigNumber(8)).toNumber()] = !((_dafny.ZERO).minus(new BigNumber((_21_v3).length))).isEqualTo(new BigNumber((_dafny.Seq.UnicodeFromString("rpjgtvwk")).length));
      _nw1[(new BigNumber(9)).toNumber()] = p0;
      _nw1[(new BigNumber(10)).toNumber()] = (_22_v4).IsProperSubsetOf(_22_v4);
      _nw1[(new BigNumber(11)).toNumber()] = (_18_v0).fm4(p0, globalState);
      _23_v5 = _nw1;
      let _index0 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_23_v5).length));
      (_23_v5)[_index0] = true;
      let _index1 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_23_v5).length));
      (_23_v5)[_index1] = p0;
      let _24_v6;
      let _nw2 = new _module.C0();
      _nw2.__ctor(_module.__default.safeDivisionInt(p1, p1));
      _24_v6 = _nw2;
      let _25_v7;
      _25_v7 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_20_v2).length));
      (globalState).f10 = ((((_25_v7).contains(p0)) ? ((_25_v7).get(p0)) : (p1))).minus((_18_v0).f14);
      let _26_v8;
      _26_v8 = _dafny.Seq.of(_24_v6, _18_v0);
      _26_v8 = _26_v8;
      if (!((_23_v5)[_module.__default.safeIndex(new BigNumber(212), new BigNumber((_23_v5).length))])) {
        (globalState).f6 = (_24_v6).f14;
        (globalState).f10 = p1;
        _25_v7 = (_25_v7).update(p0, new BigNumber(800));
        let _27_v9;
        let _nw3 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
        _27_v9 = _nw3;
        let _index2 = _module.__default.safeIndex(new BigNumber(881), new BigNumber((_27_v9).length));
        (_27_v9)[_index2] = p1;
        let _28_v10;
        _28_v10 = _dafny.Map.Empty.slice().updateUnsafe(p0,_23_v5);
        let _index3 = _module.__default.safeIndex(new BigNumber(881), new BigNumber((_27_v9).length));
        let _index4 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_23_v5).length));
        let _rhs0 = _21_v3;
        let _rhs1 = (_24_v6).f14;
        let _rhs2 = new BigNumber((_28_v10).length);
        let _rhs3 = !(p0);
        let _lhs0 = globalState;
        let _lhs1 = _27_v9;
        let _lhs2 = _module.__default.safeIndex(new BigNumber(881), new BigNumber((_27_v9).length));
        let _lhs3 = _23_v5;
        let _lhs4 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_23_v5).length));
        _21_v3 = _rhs0;
        _lhs0.f10 = _rhs1;
        _lhs1[_lhs2] = _rhs2;
        _lhs3[_lhs4] = _rhs3;
        let _29_v11;
        _29_v11 = _dafny.Map.Empty.slice().updateUnsafe((_23_v5)[_module.__default.safeIndex(new BigNumber(212), new BigNumber((_23_v5).length))],_dafny.Seq.Create(_module.__default.abs(new BigNumber(823)), ((_30_v5) => function (_31_i0) {
          return _dafny.Map.Empty.slice().updateUnsafe((_30_v5)[_module.__default.safeIndex(new BigNumber(212), new BigNumber((_30_v5).length))],(_30_v5)[_module.__default.safeIndex(new BigNumber(212), new BigNumber((_30_v5).length))]);
        })(_23_v5)));
        let _32_v12;
        _32_v12 = _dafny.Map.Empty.slice().updateUnsafe((_23_v5)[_module.__default.safeIndex(new BigNumber(212), new BigNumber((_23_v5).length))],p0);
        let _33_v13;
        _33_v13 = _dafny.Seq.of(_32_v12, _32_v12);
        _29_v11 = (_29_v11).update((_23_v5)[_module.__default.safeIndex(new BigNumber(212), new BigNumber((_23_v5).length))], _33_v13);
      } else {
        let _34_v14;
        _34_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(713),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_23_v5)[_module.__default.safeIndex(new BigNumber(212), new BigNumber((_23_v5).length))],(_23_v5)[_module.__default.safeIndex(new BigNumber(212), new BigNumber((_23_v5).length))])).length));
        let _35_v15;
        _35_v15 = _module.D1.create_DC1(_34_v14);
        let _pat_let_tv0 = _34_v14;
        let _source1 = function (_pat_let0_0) {
          return function (_36_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_37_dt__update_hcf1_h0) {
                return _module.D1.create_DC1(_37_dt__update_hcf1_h0);
              }(_pat_let1_0);
            }(_pat_let_tv0);
          }(_pat_let0_0);
        }(_35_v15);
        if (_source1.is_DC2) {
          let _38___mcc_h0 = (_source1).cf2;
          let _39___mcc_h1 = (_source1).cf3;
          let _40___mcc_h2 = (_source1).cf4;
          let _41___mcc_h3 = (_source1).cf5;
          let _42_cf5 = _41___mcc_h3;
          let _43_cf4 = _40___mcc_h2;
          let _44_cf3 = _39___mcc_h1;
          let _45_cf2 = _38___mcc_h0;
          let _46_v16;
          _46_v16 = _dafny.MultiSet.fromElements(_21_v3);
          let _47_v17;
          _47_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_46_v16).cardinality()),_21_v3);
          _21_v3 = (((_47_v17).contains((_18_v0).f14)) ? ((_47_v17).get((_18_v0).f14)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(166)), function (_48_i1) {
            return new _dafny.CodePoint('h'.codePointAt(0));
          })));
          let _49_v18;
          let _nw4 = new _module.C0();
          _nw4.__ctor((_18_v0).f14);
          _49_v18 = _nw4;
          let _index5 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_23_v5).length));
          (_23_v5)[_index5] = (_43_cf4).isLessThan(_43_cf4);
          _44_cf3 = (_24_v6).f14;
        } else if (_source1.is_DC1) {
          let _50___mcc_h4 = (_source1).cf1;
          let _51_cf1 = _50___mcc_h4;
          let _index6 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_23_v5).length));
          (_23_v5)[_index6] = (((_dafny.ZERO).minus((_24_v6).f14)).plus((_dafny.ZERO).minus(new BigNumber((_21_v3).length)))).isLessThanOrEqualTo((((_25_v7).contains(p0)) ? ((_25_v7).get(p0)) : ((_24_v6).f14)));
          let _52_v19;
          let _nw5 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
          _52_v19 = _nw5;
          let _index7 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_52_v19).length));
          (_52_v19)[_index7] = (_18_v0).f14;
          let _index8 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_52_v19).length));
          let _rhs4 = _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_21_v3).length)), (_24_v6).f14);
          let _rhs5 = (new BigNumber(74)).minus(p1);
          let _rhs6 = _module.__default.safeModuloInt((_18_v0).f14, p1);
          let _lhs5 = globalState;
          let _lhs6 = _52_v19;
          let _lhs7 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_52_v19).length));
          let _lhs8 = globalState;
          _lhs5.f10 = _rhs4;
          _lhs6[_lhs7] = _rhs5;
          _lhs8.f10 = _rhs6;
          let _53_v20;
          _53_v20 = new _dafny.CodePoint('c'.codePointAt(0));
          _53_v20 = _53_v20;
          _52_v19 = _52_v19;
        } else {
          let _54___mcc_h5 = (_source1).cf6;
          let _55_cf6 = _54___mcc_h5;
          (globalState).f10 = new BigNumber(568);
          let _index9 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_23_v5).length));
          (_23_v5)[_index9] = (_24_v6).fm5(globalState);
          let _56_v21;
          let _nw6 = new _module.C0();
          _nw6.__ctor((_18_v0).f14);
          _56_v21 = _nw6;
          let _57_v22;
          let _nw7 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Map.Empty);
          _57_v22 = _nw7;
          let _58_v23;
          _58_v23 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),_dafny.Seq.update(_20_v2, _module.__default.safeIndex((_56_v21).f14, new BigNumber((_20_v2).length)), (_23_v5)[_module.__default.safeIndex(new BigNumber(212), new BigNumber((_23_v5).length))]));
          let _index10 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_57_v22).length));
          (_57_v22)[_index10] = _58_v23;
          let _59_v24;
          _59_v24 = new _dafny.CodePoint('d'.codePointAt(0));
          let _60_v25;
          _60_v25 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_dafny.Seq.of(p0)).length)),_56_v21);
          let _61_v26;
          _61_v26 = _dafny.Map.Empty.slice().updateUnsafe(_59_v24,(((_60_v25).contains(_25_v7)) ? ((_60_v25).get(_25_v7)) : (_56_v21)));
          let _index11 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_57_v22).length));
          let _rhs7 = (_24_v6).f14;
          let _rhs8 = _58_v23;
          let _rhs9 = (_61_v26).Merge(_dafny.Map.Empty.slice().updateUnsafe(_59_v24,_18_v0));
          let _lhs9 = globalState;
          let _lhs10 = _57_v22;
          let _lhs11 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_57_v22).length));
          _lhs9.f10 = _rhs7;
          _lhs10[_lhs11] = _rhs8;
          _61_v26 = _rhs9;
        }
        let _62_v27;
        _62_v27 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_23_v5)[_module.__default.safeIndex(new BigNumber(212), new BigNumber((_23_v5).length))]);
        (globalState).f10 = (new BigNumber((_62_v27).length)).minus((_dafny.ZERO).minus(p1));
        if (!(_dafny.areEqual(_21_v3, _21_v3))) {
          let _index12 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_23_v5).length));
          (_23_v5)[_index12] = false;
          let _63_v28;
          _63_v28 = _dafny.Seq.of(p1, (_18_v0).f14);
          let _64_v29;
          _64_v29 = _module.D2.create_DC5(p0, p0);
          let _65_v30;
          _65_v30 = _dafny.MultiSet.fromElements(_64_v29, _64_v29);
          let _66_v31;
          _66_v31 = _dafny.MultiSet.fromElements(new BigNumber(568), new BigNumber((_65_v30).cardinality()), (_24_v6).f14);
          let _67_v32;
          _67_v32 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Seq.update(_63_v28, _module.__default.safeIndex((_18_v0).f14, new BigNumber((_63_v28).length)), new BigNumber((_66_v31).cardinality())));
          let _68_v33;
          _68_v33 = _dafny.Set.fromElements((_18_v0).f14, new BigNumber(831));
          let _index13 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_23_v5).length));
          (_23_v5)[_index13] = (((p0) ? (_68_v33) : (_68_v33))).contains(new BigNumber((_67_v32).length));
          let _69_v34;
          _69_v34 = new _dafny.CodePoint('l'.codePointAt(0));
          let _70_v35;
          _70_v35 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(p0,(_23_v5)[_module.__default.safeIndex(new BigNumber(212), new BigNumber((_23_v5).length))]),_69_v34);
          (globalState).f10 = (_63_v28)[_module.__default.safeIndex((new BigNumber((_70_v35).length)).minus((_24_v6).f14), new BigNumber((_63_v28).length))];
          let _71_v36;
          _71_v36 = _module.D1.create_DC1(_34_v14);
          let _72_v37;
          _72_v37 = _module.D1.create_DC3(_71_v36);
          _72_v37 = ((p0) ? (_module.D1.create_DC3(_71_v36)) : (_72_v37));
          (globalState).f10 = p1;
        } else {
          (globalState).f6 = (_18_v0).f14;
          let _73_v38;
          _73_v38 = new _dafny.CodePoint('g'.codePointAt(0));
          let _index14 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_23_v5).length));
          (_23_v5)[_index14] = !((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_21_v3, _21_v3), _module.__default.safeIndex((_18_v0).f14, new BigNumber((_dafny.Seq.Concat(_21_v3, _21_v3)).length)), _73_v38)).length))).isEqualTo((_18_v0).f14);
          let _74_v39;
          _74_v39 = _module.D2.create_DC6(!((_24_v6).fm5(globalState)), ((_24_v6).f14).isLessThan(new BigNumber((_21_v3).length)), _module.__default.fm3((_18_v0).f14, (_24_v6).f14, p0, globalState));
          let _pat_let_tv1 = _20_v2;
          let _pat_let_tv2 = p0;
          _74_v39 = function (_pat_let2_0) {
            return function (_75_dt__update__tmp_h1) {
              return function (_pat_let3_0) {
                return function (_76_dt__update_hcf11_h0) {
                  return _module.D2.create_DC6((_75_dt__update__tmp_h1).dtor_cf10, _76_dt__update_hcf11_h0, (_75_dt__update__tmp_h1).dtor_cf12);
                }(_pat_let3_0);
              }(!_dafny.Seq.contains(_pat_let_tv1, _pat_let_tv2));
            }(_pat_let2_0);
          }(_74_v39);
          let _77_v40;
          _77_v40 = _module.D1.create_DC2((_18_v0).f14, (_18_v0).f14, (_24_v6).f14, _module.__default.fm0(globalState));
          let _78_v41;
          _78_v41 = _module.D1.create_DC3(_77_v40);
          let _79_v42;
          let _nw8 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.of());
          _79_v42 = _nw8;
          let _80_v43;
          _80_v43 = _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_25_v7,(_24_v6).f14)).length), (_24_v6).f14);
          let _index15 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_79_v42).length));
          (_79_v42)[_index15] = _dafny.Seq.Concat(_80_v43, _80_v43);
          let _index16 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_23_v5).length));
          let _index17 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_79_v42).length));
          let _rhs10 = _78_v41;
          let _rhs11 = (p0) === ((_23_v5)[_module.__default.safeIndex(new BigNumber(212), new BigNumber((_23_v5).length))]);
          let _rhs12 = _dafny.Seq.Concat(_80_v43, _80_v43);
          let _lhs12 = _23_v5;
          let _lhs13 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_23_v5).length));
          let _lhs14 = _79_v42;
          let _lhs15 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_79_v42).length));
          _78_v41 = _rhs10;
          _lhs12[_lhs13] = _rhs11;
          _lhs14[_lhs15] = _rhs12;
          let _index18 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_23_v5).length));
          (_23_v5)[_index18] = false;
        }
        let _81_v44;
        _81_v44 = false;
        _81_v44 = _81_v44;
        let _82_v45;
        _82_v45 = _module.D1.create_DC2(new BigNumber(-506), new BigNumber((_34_v14).length), (_24_v6).f14, (_24_v6).f14);
        let _83_v46;
        _83_v46 = _dafny.Seq.of((_dafny.ZERO).minus((_24_v6).f14), (_82_v45).dtor_cf4);
        let _index19 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_23_v5).length));
        (_23_v5)[_index19] = (new BigNumber((_83_v46).length)).isLessThan(((_24_v6).f14).minus((_dafny.ZERO).minus((_83_v46)[_module.__default.safeIndex((_18_v0).f14, new BigNumber((_83_v46).length))])));
      }
      return;
    }
    static Main(__noArgsParameter) {
      let _84_v0;
      _84_v0 = new BigNumber(457);
      let _85_v1;
      _85_v1 = false;
      let _86_v2;
      _86_v2 = _dafny.Map.Empty.slice().updateUnsafe(_84_v0,_85_v1);
      let _87_v3;
      _87_v3 = _dafny.Seq.of((((_86_v2).contains(new BigNumber((_dafny.Seq.UnicodeFromString("koq")).length))) ? ((_86_v2).get(new BigNumber((_dafny.Seq.UnicodeFromString("koq")).length))) : (_85_v1)));
      let _88_v4;
      _88_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(924),_87_v3);
      let _89_v5;
      _89_v5 = _dafny.Seq.UnicodeFromString("rlnyx");
      let _90_v6;
      let _nw9 = Array((new BigNumber(14)).toNumber()).fill(false);
      _90_v6 = _nw9;
      let _91_globalState;
      let _nw10 = new _module.GlobalState();
      _nw10.__ctor(true, true, true, true, (_88_v4).Merge((_88_v4).update(_84_v0, _87_v3)), new BigNumber(-8), new BigNumber(256), false, _89_v5, _89_v5, new BigNumber(133), _90_v6, true, new BigNumber(191));
      _91_globalState = _nw10;
      if (!((_module.__default.fm0(_91_globalState)).isLessThan((_dafny.ZERO).minus((_dafny.ZERO).minus(_84_v0)))) || (_85_v1)) {
        let _92_v7;
        let _nw11 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Map.Empty);
        _92_v7 = _nw11;
        let _93_v8;
        _93_v8 = _dafny.Map.Empty.slice().updateUnsafe(_92_v7,_85_v1);
        _93_v8 = (_93_v8).update(_92_v7, _85_v1);
        _84_v0 = _84_v0;
        (_91_globalState).f6 = new BigNumber(-930);
        let _94_v9;
        let _nw12 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
        _94_v9 = _nw12;
        _94_v9 = _94_v9;
        let _95_v10;
        _95_v10 = _dafny.Seq.of(_84_v0);
        let _96_v11;
        _96_v11 = _dafny.MultiSet.fromElements(_84_v0);
        let _97_v12;
        _97_v12 = _dafny.Set.fromElements(_96_v11, _96_v11);
        let _98_v13;
        _98_v13 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_97_v12).length),_84_v0);
        let _99_v14;
        _99_v14 = _dafny.Set.fromElements(new BigNumber(320), new BigNumber((_89_v5).length), new BigNumber((_95_v10).length), _84_v0, new BigNumber((_98_v13).length));
        _85_v1 = ((_99_v14).Union(_dafny.Set.fromElements(_84_v0, new BigNumber((_dafny.Seq.UnicodeFromString("ryggvk")).length)))).IsProperSubsetOf(_dafny.Set.fromElements(_84_v0, new BigNumber(-755), _84_v0, _84_v0));
      } else {
        (_91_globalState).f10 = (_dafny.ZERO).minus(_84_v0);
        let _100_v15;
        _100_v15 = _dafny.Map.Empty.slice().updateUnsafe(_85_v1,_84_v0);
        let _101_v16;
        _101_v16 = _dafny.Seq.of(new BigNumber((_86_v2).length), new BigNumber((_100_v15).length));
        _85_v1 = _module.__default.fm1((_101_v16)[_module.__default.safeIndex(new BigNumber(139), new BigNumber((_101_v16).length))], false, _85_v1, _dafny.Seq.Concat(_dafny.Seq.of(_86_v2, _module.__default.fm2(_84_v0, _85_v1, _85_v1, _91_globalState)), _dafny.Seq.of(_86_v2)), _91_globalState);
        let _102_v17;
        _102_v17 = _dafny.Map.Empty.slice().updateUnsafe(_84_v0,_module.__default.fm0(_91_globalState));
        _102_v17 = (_102_v17).update(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_89_v5, _module.__default.safeIndex(_84_v0, new BigNumber((_89_v5).length)), _module.__default.fm3(_84_v0, _84_v0, _85_v1, _91_globalState)), _89_v5)).length), _84_v0);
        (_91_globalState).f6 = _module.__default.safeDivisionInt(_84_v0, _84_v0);
        _85_v1 = _85_v1;
      }
      let _103_v18;
      _103_v18 = _dafny.Map.Empty.slice().updateUnsafe(_84_v0,_84_v0);
      let _104_v20;
      _104_v20 = _dafny.Seq.of(function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of _dafny.IntegerRange(new BigNumber(920), new BigNumber(-802))) {
          let _105_v19 = _compr_2;
          if (((new BigNumber(920)).isLessThanOrEqualTo(_105_v19)) && ((_105_v19).isLessThan(new BigNumber(-802)))) {
            _coll2.push([(_105_v19).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_85_v1,_85_v1)).length)),_85_v1]);
          }
        }
        return _coll2;
      }());
      let _106_v21;
      _106_v21 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_103_v18).length),_104_v20);
      let _index20 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length));
      (_90_v6)[_index20] = _module.__default.fm1(_84_v0, true, false, (((_106_v21).contains(_84_v0)) ? ((_106_v21).get(_84_v0)) : (_104_v20)), _91_globalState);
      let _index21 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length));
      (_90_v6)[_index21] = _85_v1;
      let _107_v22;
      _107_v22 = new _dafny.CodePoint('i'.codePointAt(0));
      let _108_v23;
      _108_v23 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("c"), _module.__default.safeIndex(_84_v0, new BigNumber((_dafny.Seq.UnicodeFromString("c")).length)), _107_v22),(_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))]);
      _85_v1 = !((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))]) || ((((_108_v23).contains(_89_v5)) ? ((_108_v23).get(_89_v5)) : (true)));
      let _109_v24;
      let _init0 = ((_110_v0) => function (_111_i0) {
        return _module.__default.safeDivisionInt(_111_i0, _110_v0);
      })(_84_v0);
      let _nw13 = Array((new BigNumber(4)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw13.length); _i0_0++) {
        _nw13[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _109_v24 = _nw13;
      let _index22 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length));
      (_109_v24)[_index22] = _84_v0;
      let _index23 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_109_v24).length));
      (_109_v24)[_index23] = new BigNumber((_dafny.MultiSet.fromElements((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))], _85_v1)).cardinality());
      let _112_v25;
      _112_v25 = _dafny.Seq.of(_84_v0, _84_v0, _84_v0);
      let _index24 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length));
      let _index25 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_109_v24).length));
      let _rhs13 = _module.__default.safeModuloInt(_84_v0, (_dafny.ZERO).minus(_module.__default.fm0(_91_globalState)));
      let _rhs14 = ((_84_v0).plus(new BigNumber((_103_v18).length))).multipliedBy(_84_v0);
      let _rhs15 = (_89_v5)[_module.__default.safeIndex(_84_v0, new BigNumber((_89_v5).length))];
      let _rhs16 = (_dafny.ZERO).minus(_84_v0);
      let _rhs17 = _112_v25;
      let _lhs16 = _109_v24;
      let _lhs17 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length));
      let _lhs18 = _91_globalState;
      let _lhs19 = _109_v24;
      let _lhs20 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_109_v24).length));
      _lhs16[_lhs17] = _rhs13;
      _lhs18.f6 = _rhs14;
      _107_v22 = _rhs15;
      _lhs19[_lhs20] = _rhs16;
      _112_v25 = _rhs17;
      let _hi0 = (_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))];
      for (let _113_i1 = (_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))]; _113_i1.isLessThan(_hi0); _113_i1 = _113_i1.plus(_dafny.ONE)) {
        let _114_v26;
        _114_v26 = _dafny.Set.fromElements(_85_v1);
        let _115_v27;
        _115_v27 = _dafny.Map.Empty.slice().updateUnsafe(_85_v1,new BigNumber((_114_v26).length));
        _115_v27 = (_115_v27).update(true, new BigNumber(482));
        _module.__default.m0((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))], (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_112_v25, _112_v25)).length))), _91_globalState);
        _87_v3 = _dafny.Seq.Concat(_87_v3, _87_v3);
        let _116_v28;
        _116_v28 = _dafny.Map.Empty.slice().updateUnsafe(_87_v3,_109_v24);
        let _117_v29;
        _117_v29 = _dafny.Seq.of(_109_v24, _109_v24, (((_116_v28).contains(_dafny.Seq.of(_85_v1))) ? ((_116_v28).get(_dafny.Seq.of(_85_v1))) : (_109_v24)));
        _109_v24 = (_117_v29)[_module.__default.safeIndex(_84_v0, new BigNumber((_117_v29).length))];
      }
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_90_v6).length))) {
        let _118_i2 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_118_i2)) && ((_118_i2).isLessThan(new BigNumber((_90_v6).length))))) {
          (_90_v6)[(_118_i2)] = _85_v1;
        }
      }
      let _119_v30;
      _119_v30 = _dafny.Map.Empty.slice().updateUnsafe((_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))],_107_v22);
      _119_v30 = (_119_v30).update(_84_v0, _107_v22);
      if (_85_v1) {
        (_91_globalState).f6 = ((_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))]).multipliedBy((_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))]);
        if ((!(false)) || (_module.__default.fm1(new BigNumber((_89_v5).length), (_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))], (_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))], _104_v20, _91_globalState))) {
          _85_v1 = (_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))];
          _90_v6 = _90_v6;
          _89_v5 = _89_v5;
          let _120_v31;
          let _nw14 = new _module.C0();
          _nw14.__ctor((_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))]);
          _120_v31 = _nw14;
          let _121_v32;
          _121_v32 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_87_v3, _87_v3),!((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))]) || ((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))]));
          _121_v32 = (_121_v32).update(_dafny.Seq.update(_87_v3, _module.__default.safeIndex(_84_v0, new BigNumber((_87_v3).length)), (_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))]), !(((_87_v3)[_module.__default.safeIndex(_84_v0, new BigNumber((_87_v3).length))]) === ((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))])));
        } else {
          let _122_v33;
          _122_v33 = _dafny.Seq.of(_90_v6);
          _module.__default.m0((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))], new BigNumber((_122_v33).length), _91_globalState);
          let _index26 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length));
          (_109_v24)[_index26] = _module.__default.fm0(_91_globalState);
          _85_v1 = _85_v1;
          let _123_v34;
          _123_v34 = _dafny.Seq.of(_89_v5, _89_v5);
          let _124_v35;
          _124_v35 = _dafny.Map.Empty.slice().updateUnsafe((_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))],_dafny.Seq.of(_84_v0, (_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))]));
          let _index27 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length));
          let _rhs18 = (_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))];
          let _rhs19 = new BigNumber(602);
          let _rhs20 = !(!((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))]));
          let _rhs21 = _module.__default.safeDivisionInt(_84_v0, new BigNumber((_dafny.Seq.update(_123_v34, _module.__default.safeIndex(new BigNumber(790), new BigNumber((_123_v34).length)), _dafny.Seq.UnicodeFromString("c"))).length));
          let _rhs22 = (_124_v35).contains(((_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))]).multipliedBy(new BigNumber((_87_v3).length)));
          let _lhs21 = _91_globalState;
          let _lhs22 = _91_globalState;
          let _lhs23 = _90_v6;
          let _lhs24 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length));
          _85_v1 = _rhs18;
          _lhs21.f10 = _rhs19;
          _85_v1 = _rhs20;
          _lhs22.f6 = _rhs21;
          _lhs23[_lhs24] = _rhs22;
          let _125_v36;
          let _nw15 = new _module.C0();
          _nw15.__ctor(new BigNumber(478));
          _125_v36 = _nw15;
        }
        let _index28 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length));
        (_90_v6)[_index28] = _85_v1;
        let _126_v37;
        _126_v37 = _dafny.Map.Empty.slice().updateUnsafe(_89_v5,_90_v6);
        _90_v6 = (((_126_v37).contains(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(76)), ((_131_v22) => function (_132_i3) {
          return _131_v22;
        })(_107_v22)), _89_v5), _module.__default.safeIndex(_84_v0, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(76)), ((_133_v22) => function (_134_i3) {
          return _133_v22;
        })(_107_v22)), _89_v5)).length)), new _dafny.CodePoint('c'.codePointAt(0))))) ? ((_126_v37).get(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(76)), ((_127_v22) => function (_128_i3) {
          return _127_v22;
        })(_107_v22)), _89_v5), _module.__default.safeIndex(_84_v0, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(76)), ((_129_v22) => function (_130_i3) {
          return _129_v22;
        })(_107_v22)), _89_v5)).length)), new _dafny.CodePoint('c'.codePointAt(0))))) : (_90_v6));
        let _135_v38;
        _135_v38 = _dafny.Set.fromElements(_112_v25);
        _module.__default.m0((_135_v38).IsProperSubsetOf(_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(-144)))), (_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))], _91_globalState);
      } else {
        _107_v22 = new _dafny.CodePoint('y'.codePointAt(0));
        _89_v5 = _89_v5;
        if ((new BigNumber((_89_v5).length)).isEqualTo(_84_v0)) {
          let _136_v39;
          let _nw16 = Array((new BigNumber(25)).toNumber()).fill(_dafny.MultiSet.Empty);
          _136_v39 = _nw16;
          let _137_v40;
          let _nw17 = new _module.C0();
          _nw17.__ctor((_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))]);
          _137_v40 = _nw17;
          let _138_v41;
          _138_v41 = _dafny.Seq.of(_137_v40);
          let _139_v42;
          _139_v42 = _dafny.MultiSet.fromElements(_138_v41, _138_v41);
          let _140_v43;
          _140_v43 = _dafny.Map.Empty.slice().updateUnsafe(_109_v24,_139_v42);
          let _141_v44;
          _141_v44 = _dafny.Seq.of(_138_v41);
          let _142_v45;
          _142_v45 = _dafny.Map.Empty.slice().updateUnsafe((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))],_85_v1);
          let _143_v46;
          _143_v46 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_84_v0),_142_v45);
          let _index29 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_136_v39).length));
          (_136_v39)[_index29] = ((((_140_v43).contains(_109_v24)) ? ((_140_v43).get(_109_v24)) : (_dafny.MultiSet.FromArray(_141_v44)))).update(_138_v41, _module.__default.abs(new BigNumber((_143_v46).length)));
          let _index30 = _module.__default.safeIndex(new BigNumber(793), new BigNumber((_136_v39).length));
          (_136_v39)[_index30] = _139_v42;
          _module.__default.m0((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))], (_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))], _91_globalState);
          _112_v25 = _dafny.Seq.Concat(_112_v25, _dafny.Seq.Concat(_112_v25, _112_v25));
          _module.__default.m0(_85_v1, _84_v0, _91_globalState);
          _103_v18 = (_module.D1.create_DC1(_103_v18)).dtor_cf1;
        } else {
          let _index31 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length));
          (_90_v6)[_index31] = false;
          let _144_v47;
          let _nw18 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Map.Empty);
          _144_v47 = _nw18;
          let _145_v48;
          _145_v48 = _dafny.Map.Empty.slice().updateUnsafe(_144_v47,new BigNumber((_112_v25).length));
          _145_v48 = (_145_v48).update(_144_v47, (new BigNumber((_87_v3).length)).minus((_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))]));
          let _index32 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length));
          (_109_v24)[_index32] = new BigNumber(878);
          _88_v4 = _88_v4;
          let _146_v49;
          _146_v49 = _dafny.Set.fromElements((_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))], (_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))], (_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))]);
          let _147_v50;
          _147_v50 = _dafny.Map.Empty.slice().updateUnsafe((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))],(_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))]);
          let _148_v51;
          _148_v51 = _dafny.Map.Empty.slice().updateUnsafe(_146_v49,new BigNumber(((_147_v50).update((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))], (_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))])).length));
          let _149_v52;
          let _nw19 = new _module.C0();
          _nw19.__ctor((_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))]);
          _149_v52 = _nw19;
          let _150_v53;
          _150_v53 = _dafny.Seq.of(_149_v52, _149_v52);
          let _151_v54;
          _151_v54 = _dafny.Seq.of(_148_v51, _148_v51);
          let _152_v55;
          let _nw20 = Array((new BigNumber(19)).toNumber());
          _nw20[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_84_v0),_84_v0);
          _nw20[(_dafny.ONE).toNumber()] = (_148_v51).update(_146_v49, new BigNumber((_150_v53).length));
          _nw20[(new BigNumber(2)).toNumber()] = (_148_v51).Merge(_148_v51);
          _nw20[(new BigNumber(3)).toNumber()] = _148_v51;
          _nw20[(new BigNumber(4)).toNumber()] = (_151_v54)[_module.__default.safeIndex(_84_v0, new BigNumber((_151_v54).length))];
          _nw20[(new BigNumber(5)).toNumber()] = (_148_v51).Merge(_148_v51);
          _nw20[(new BigNumber(6)).toNumber()] = _148_v51;
          _nw20[(new BigNumber(7)).toNumber()] = _148_v51;
          _nw20[(new BigNumber(8)).toNumber()] = _module.__default.fm6((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))], _85_v1, (_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))], _85_v1, _91_globalState);
          _nw20[(new BigNumber(9)).toNumber()] = _148_v51;
          _nw20[(new BigNumber(10)).toNumber()] = (_148_v51).Merge(_148_v51);
          _nw20[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_146_v49,new BigNumber(735));
          _nw20[(new BigNumber(12)).toNumber()] = _148_v51;
          _nw20[(new BigNumber(13)).toNumber()] = _148_v51;
          _nw20[(new BigNumber(14)).toNumber()] = (_148_v51).Merge(_148_v51);
          _nw20[(new BigNumber(15)).toNumber()] = _148_v51;
          _nw20[(new BigNumber(16)).toNumber()] = (_148_v51).update(_146_v49, (_dafny.ZERO).minus((_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))]));
          _nw20[(new BigNumber(17)).toNumber()] = _148_v51;
          _nw20[(new BigNumber(18)).toNumber()] = _148_v51;
          _152_v55 = _nw20;
          let _index33 = _module.__default.safeIndex(new BigNumber(728), new BigNumber((_152_v55).length));
          (_152_v55)[_index33] = _148_v51;
          let _153_v56;
          _153_v56 = _dafny.Seq.of(_89_v5);
          let _index34 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length));
          let _index35 = _module.__default.safeIndex(new BigNumber(728), new BigNumber((_152_v55).length));
          let _rhs23 = new BigNumber((_153_v56).length);
          let _rhs24 = (_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))];
          let _rhs25 = _module.__default.fm0(_91_globalState);
          let _rhs26 = (_148_v51).Merge(_148_v51);
          let _lhs25 = _109_v24;
          let _lhs26 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length));
          let _lhs27 = _91_globalState;
          let _lhs28 = _152_v55;
          let _lhs29 = _module.__default.safeIndex(new BigNumber(728), new BigNumber((_152_v55).length));
          _lhs25[_lhs26] = _rhs23;
          _85_v1 = _rhs24;
          _lhs27.f6 = _rhs25;
          _lhs28[_lhs29] = _rhs26;
        }
        let _154_v57;
        _154_v57 = _dafny.Map.Empty.slice().updateUnsafe(_85_v1,_84_v0);
        _154_v57 = _154_v57;
        let _index36 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length));
        (_90_v6)[_index36] = (_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))];
      }
      let _hi1 = new BigNumber(998);
      for (let _155_i4 = (_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))]; _155_i4.isLessThan(_hi1); _155_i4 = _155_i4.plus(_dafny.ONE)) {
        let _156_v58;
        let _nw21 = Array((new BigNumber(6)).toNumber());
        _156_v58 = _nw21;
        let _157_v59;
        let _nw22 = new _module.C0();
        _nw22.__ctor(new BigNumber((_dafny.Seq.UnicodeFromString("ook")).length));
        _157_v59 = _nw22;
        let _index37 = _module.__default.safeIndex(new BigNumber(728), new BigNumber((_156_v58).length));
        (_156_v58)[_index37] = _157_v59;
        let _index38 = _module.__default.safeIndex(new BigNumber(728), new BigNumber((_156_v58).length));
        (_156_v58)[_index38] = _157_v59;
        let _158_v60;
        _158_v60 = _dafny.MultiSet.fromElements(_85_v1);
        let _159_v61;
        _159_v61 = _module.D1.create_DC2((_157_v59).f14, new BigNumber((_158_v60).cardinality()), new BigNumber(250), _155_i4);
        let _160_v62;
        _160_v62 = _module.D1.create_DC3(_159_v61);
        let _161_v63;
        _161_v63 = _module.D1.create_DC3(_159_v61);
        let _pat_let_tv3 = _159_v61;
        _161_v63 = function (_pat_let4_0) {
          return function (_162_dt__update__tmp_h0) {
            return function (_pat_let5_0) {
              return function (_163_dt__update_hcf6_h0) {
                return _module.D1.create_DC3(_163_dt__update_hcf6_h0);
              }(_pat_let5_0);
            }(_pat_let_tv3);
          }(_pat_let4_0);
        }(_161_v63);
        let _164_v64;
        _164_v64 = _dafny.Map.Empty.slice().updateUnsafe((_156_v58)[_module.__default.safeIndex(new BigNumber(728), new BigNumber((_156_v58).length))],(_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))]);
        if ((((_164_v64).contains((_156_v58)[_module.__default.safeIndex(new BigNumber(728), new BigNumber((_156_v58).length))])) ? ((_164_v64).get((_156_v58)[_module.__default.safeIndex(new BigNumber(728), new BigNumber((_156_v58).length))])) : ((_85_v1) && ((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))])))) {
          let _165_v65;
          _165_v65 = _dafny.Set.fromElements((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))]);
          let _166_v66;
          _166_v66 = _dafny.MultiSet.fromElements(_90_v6);
          let _nw23 = Array((new BigNumber(14)).toNumber());
          _nw23[(_dafny.ZERO).toNumber()] = _155_i4;
          _nw23[(_dafny.ONE).toNumber()] = (_157_v59).f14;
          _nw23[(new BigNumber(2)).toNumber()] = _155_i4;
          _nw23[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_158_v60).cardinality()), _155_i4);
          _nw23[(new BigNumber(4)).toNumber()] = new BigNumber(-490);
          _nw23[(new BigNumber(5)).toNumber()] = _module.__default.safeModuloInt((_157_v59).f14, new BigNumber(461));
          _nw23[(new BigNumber(6)).toNumber()] = _84_v0;
          _nw23[(new BigNumber(7)).toNumber()] = (_157_v59).f14;
          _nw23[(new BigNumber(8)).toNumber()] = (_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))];
          _nw23[(new BigNumber(9)).toNumber()] = new BigNumber((_165_v65).length);
          _nw23[(new BigNumber(10)).toNumber()] = (_157_v59).f14;
          _nw23[(new BigNumber(11)).toNumber()] = (_157_v59).f14;
          _nw23[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus((_157_v59).f14);
          _nw23[(new BigNumber(13)).toNumber()] = (new BigNumber((_166_v66).cardinality())).multipliedBy(_84_v0);
          _109_v24 = _nw23;
          _module.__default.m0(!(_85_v1), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(556)), ((_167_v22) => function (_168_i5) {
            return _167_v22;
          })(_107_v22))).length), _91_globalState);
          _module.__default.m0((!(_85_v1)) && (_85_v1), (_157_v59).f14, _91_globalState);
          _157_v59 = _157_v59;
          let _169_v67;
          _169_v67 = _dafny.MultiSet.fromElements(_107_v22, _107_v22);
          let _170_v68;
          let _nw24 = new _module.C0();
          _nw24.__ctor(((_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))]).multipliedBy((((_169_v67).contains(_107_v22)) ? ((_169_v67).get(_107_v22)) : (_155_i4))));
          _170_v68 = _nw24;
        } else {
          let _171_v69;
          _171_v69 = _module.D1.create_DC1((_103_v18).Merge(_103_v18));
          let _172_v70;
          _172_v70 = _dafny.Set.fromElements((_156_v58)[_module.__default.safeIndex(new BigNumber(728), new BigNumber((_156_v58).length))]);
          let _173_v71;
          _173_v71 = _module.D2.create_DC4(_172_v70);
          let _index39 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length));
          let _index40 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length));
          let _index41 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length));
          let _rhs27 = !(_module.__default.fm1((_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))], (_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))], (_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))], _dafny.Seq.of(_86_v2), _91_globalState));
          let _rhs28 = ((((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))]) ? (_172_v70) : ((_173_v71).dtor_cf7))).IsDisjointFrom(_172_v70);
          let _rhs29 = _85_v1;
          let _rhs30 = _171_v69;
          let _rhs31 = (_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))];
          let _lhs30 = _90_v6;
          let _lhs31 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length));
          let _lhs32 = _90_v6;
          let _lhs33 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length));
          let _lhs34 = _90_v6;
          let _lhs35 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length));
          _lhs30[_lhs31] = _rhs27;
          _85_v1 = _rhs28;
          _lhs32[_lhs33] = _rhs29;
          _171_v69 = _rhs30;
          _lhs34[_lhs35] = _rhs31;
          let _index42 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length));
          (_109_v24)[_index42] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_84_v0, new BigNumber(371)));
          let _init1 = ((_174_v6, _175_v1) => function (_176_i6) {
            return _module.__default.safeDivisionInt(_176_i6, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of((_174_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_174_v6).length))], _175_v1, (_174_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_174_v6).length))], true)).length)));
          })(_90_v6, _85_v1);
          let _nw25 = Array((new BigNumber(7)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw25.length); _i0_1++) {
            _nw25[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _109_v24 = _nw25;
          let _177_v72;
          let _nw26 = new _module.C0();
          _nw26.__ctor((_dafny.ZERO).minus((_84_v0).minus((_157_v59).f14)));
          _177_v72 = _nw26;
          _module.__default.m0(_85_v1, _module.__default.safeDivisionInt((_157_v59).f14, (_157_v59).f14), _91_globalState);
        }
        _85_v1 = (_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))];
      }
      let _178_v73;
      let _nw27 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _178_v73 = _nw27;
      let _index43 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length));
      let _rhs32 = _178_v73;
      let _rhs33 = _85_v1;
      let _lhs36 = _90_v6;
      let _lhs37 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length));
      _178_v73 = _rhs32;
      _lhs36[_lhs37] = _rhs33;
      let _index44 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length));
      (_109_v24)[_index44] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))]));
      let _179_v74;
      let _nw28 = new _module.C0();
      _nw28.__ctor(_84_v0);
      _179_v74 = _nw28;
      let _180_v75;
      _180_v75 = _dafny.Seq.of(_179_v74);
      let _181_v76;
      _181_v76 = _dafny.MultiSet.fromElements(_179_v74);
      let _182_v77;
      _182_v77 = _dafny.Set.fromElements(new BigNumber((_180_v75).length), _84_v0, (_dafny.ZERO).minus((_179_v74).f14), new BigNumber((_87_v3).length), new BigNumber((_181_v76).cardinality()));
      _85_v1 = ((_182_v77).Union(_182_v77)).IsDisjointFrom((_182_v77).Union(_182_v77));
      if (_85_v1) {
        let _183_v78;
        _183_v78 = _dafny.Map.Empty.slice().updateUnsafe(_85_v1,(_179_v74).f14);
        let _184_v79;
        _184_v79 = _dafny.Seq.of(_183_v78);
        let _185_v80;
        let _init2 = ((_186_v3) => function (_187_i7) {
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(700)), ((_188_v3) => function (_189_i8) {
            return _188_v3;
          })(_186_v3));
        })(_87_v3);
        let _nw29 = Array((new BigNumber(10)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw29.length); _i0_2++) {
          _nw29[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _185_v80 = _nw29;
        let _190_v81;
        _190_v81 = _dafny.Seq.of(_dafny.Seq.of(_85_v1, _85_v1));
        let _index45 = _module.__default.safeIndex(new BigNumber(593), new BigNumber((_185_v80).length));
        (_185_v80)[_index45] = _190_v81;
        let _191_v82;
        _191_v82 = _dafny.Map.Empty.slice().updateUnsafe(_107_v22,_107_v22);
        let _192_v83;
        _192_v83 = _dafny.Set.fromElements(_191_v82);
        let _193_v84;
        _193_v84 = _module.D2.create_DC6(!((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))]), _85_v1, _107_v22);
        let _194_v86;
        _194_v86 = _dafny.Set.fromElements(_107_v22);
        let _195_v87;
        _195_v87 = _module.D3.create_DC8(_194_v86);
        let _index46 = _module.__default.safeIndex(new BigNumber(593), new BigNumber((_185_v80).length));
        let _rhs34 = _dafny.Seq.update(_184_v79, _module.__default.safeIndex((_179_v74).f14, new BigNumber((_184_v79).length)), _183_v78);
        let _rhs35 = _190_v81;
        let _rhs36 = (_192_v83).Intersect(function () {
          let _coll3 = new _dafny.Set();
          for (const _compr_3 of (_dafny.MultiSet.fromElements(_191_v82, function () {
            let _coll4 = new _dafny.Map();
            for (const _compr_4 of ((_195_v87).dtor_cf14).Elements) {
              let _196_v85 = _compr_4;
              if (((_195_v87).dtor_cf14).contains(_196_v85)) {
                _coll4.push([_196_v85,_107_v22]);
              }
            }
            return _coll4;
          }())).Elements) {
            let _197_v88 = _compr_3;
            if ((_dafny.MultiSet.fromElements(_191_v82, function () {
              let _coll5 = new _dafny.Map();
              for (const _compr_5 of ((_195_v87).dtor_cf14).Elements) {
                let _196_v85 = _compr_5;
                if (((_195_v87).dtor_cf14).contains(_196_v85)) {
                  _coll5.push([_196_v85,_107_v22]);
                }
              }
              return _coll5;
            }())).contains(_197_v88)) {
              _coll3.add(_197_v88);
            }
          }
          return _coll3;
        }());
        let _rhs37 = _193_v84;
        let _lhs38 = _185_v80;
        let _lhs39 = _module.__default.safeIndex(new BigNumber(593), new BigNumber((_185_v80).length));
        _184_v79 = _rhs34;
        _lhs38[_lhs39] = _rhs35;
        _192_v83 = _rhs36;
        _193_v84 = _rhs37;
        _module.__default.m0((_module.D2.create_DC5(_85_v1, (_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))])).dtor_cf9, _module.__default.safeDivisionInt((_179_v74).f14, (_179_v74).f14), _91_globalState);
        let _index47 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length));
        (_109_v24)[_index47] = (_179_v74).f14;
        let _index48 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length));
        let _rhs38 = !(new BigNumber((_89_v5).length)).isEqualTo(new BigNumber((_103_v18).length));
        let _rhs39 = function () {
          let _coll6 = new _dafny.Map();
          for (const _compr_6 of _dafny.IntegerRange(new BigNumber(-729), new BigNumber(42))) {
            let _198_v89 = _compr_6;
            if (((new BigNumber(-729)).isLessThanOrEqualTo(_198_v89)) && ((_198_v89).isLessThan(new BigNumber(42)))) {
              _coll6.push([_module.__default.safeModuloInt(_198_v89, (_179_v74).f14),_85_v1]);
            }
          }
          return _coll6;
        }();
        let _rhs40 = new BigNumber(516);
        let _rhs41 = (new BigNumber((_89_v5).length)).multipliedBy((_179_v74).f14);
        let _rhs42 = (_dafny.ZERO).minus(_84_v0);
        let _lhs40 = _109_v24;
        let _lhs41 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length));
        let _lhs42 = _91_globalState;
        let _lhs43 = _91_globalState;
        _85_v1 = _rhs38;
        _86_v2 = _rhs39;
        _lhs40[_lhs41] = _rhs40;
        _lhs42.f10 = _rhs41;
        _lhs43.f6 = _rhs42;
        _86_v2 = (_86_v2).update((_179_v74).f14, _85_v1);
      } else {
        if (_85_v1) {
          let _199_v90;
          _199_v90 = _dafny.Seq.of(_182_v77);
          let _index49 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length));
          let _rhs43 = (_179_v74).fm4(_85_v1, _91_globalState);
          let _rhs44 = _85_v1;
          let _rhs45 = (_199_v90)[_module.__default.safeIndex(((_179_v74).f14).plus((_179_v74).f14), new BigNumber((_199_v90).length))];
          let _rhs46 = _107_v22;
          let _lhs44 = _90_v6;
          let _lhs45 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length));
          _lhs44[_lhs45] = _rhs43;
          _85_v1 = _rhs44;
          _182_v77 = _rhs45;
          _107_v22 = _rhs46;
          let _200_v91;
          _200_v91 = _dafny.Seq.of(_dafny.Set.fromElements(_107_v22));
          let _201_v92;
          _201_v92 = _dafny.Map.Empty.slice().updateUnsafe(_179_v74,(_200_v91)[_module.__default.safeIndex((_179_v74).f14, new BigNumber((_200_v91).length))]);
          let _202_v93;
          _202_v93 = _module.D3.create_DC8((((_201_v92).contains(_179_v74)) ? ((_201_v92).get(_179_v74)) : (_dafny.Set.fromElements(_107_v22))));
          let _203_v94;
          _203_v94 = _dafny.MultiSet.fromElements(_202_v93);
          let _204_v95;
          _204_v95 = _dafny.MultiSet.fromElements(_85_v1, (_85_v1) || (_85_v1), (_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))], _85_v1, (_203_v94).equals(_203_v94));
          _204_v95 = ((_204_v95).Intersect(_204_v95)).Difference(_204_v95);
          _module.__default.m0(_85_v1, (_179_v74).f14, _91_globalState);
          let _205_v96;
          let _nw30 = Array((new BigNumber(28)).toNumber());
          _205_v96 = _nw30;
          _205_v96 = _205_v96;
          let _206_v97;
          _206_v97 = _dafny.MultiSet.fromElements(new BigNumber(235));
          (_91_globalState).f10 = (_112_v25)[_module.__default.safeIndex((((_206_v97).contains((_179_v74).f14)) ? ((_206_v97).get((_179_v74).f14)) : ((_179_v74).f14)), new BigNumber((_112_v25).length))];
        } else {
          let _nw31 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
          _109_v24 = _nw31;
          let _207_v98;
          _207_v98 = _dafny.MultiSet.fromElements(_107_v22);
          let _index50 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length));
          (_109_v24)[_index50] = (new BigNumber(((_207_v98).Union(_dafny.MultiSet.FromArray(_89_v5))).cardinality())).minus((_179_v74).f14);
          let _208_v99;
          _208_v99 = _dafny.Set.fromElements((_180_v75)[_module.__default.safeIndex(_84_v0, new BigNumber((_180_v75).length))]);
          let _209_v100;
          _209_v100 = _module.D2.create_DC4((_208_v99).Intersect(_208_v99));
          _209_v100 = _209_v100;
          let _210_v101;
          _210_v101 = _module.D2.create_DC5(!((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))]), _85_v1);
          let _211_v102;
          _211_v102 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_109_v24, _109_v24, _109_v24),(((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))]) ? (_210_v101) : (_210_v101)));
          _211_v102 = (_211_v102).update(_dafny.Set.fromElements(_109_v24, _109_v24), _210_v101);
          _module.__default.m0((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))], (_179_v74).f14, _91_globalState);
        }
        _109_v24 = _109_v24;
        let _index51 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length));
        (_109_v24)[_index51] = (_dafny.ZERO).minus((_179_v74).f14);
        let _212_v103;
        _212_v103 = _dafny.MultiSet.fromElements(_85_v1);
        let _213_v104;
        _213_v104 = _dafny.Map.Empty.slice().updateUnsafe(_86_v2,!((_87_v3)[_module.__default.safeIndex((((_212_v103).contains((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))])) ? ((_212_v103).get((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))])) : ((_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))])), new BigNumber((_87_v3).length))]));
        let _index52 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length));
        let _rhs47 = (_212_v103).IsProperSubsetOf(_212_v103);
        let _rhs48 = _module.__default.safeModuloInt(_84_v0, (_179_v74).f14);
        let _rhs49 = (_213_v104).Merge(_213_v104);
        let _lhs46 = _90_v6;
        let _lhs47 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length));
        let _lhs48 = _91_globalState;
        _lhs46[_lhs47] = _rhs47;
        _lhs48.f6 = _rhs48;
        _213_v104 = _rhs49;
        if (_85_v1) {
          let _index53 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length));
          (_90_v6)[_index53] = !(_84_v0).isEqualTo((_179_v74).f14);
          let _index54 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length));
          (_109_v24)[_index54] = (_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))];
          _85_v1 = (_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))];
          let _214_v105;
          _214_v105 = _dafny.Seq.of(_112_v25);
          _214_v105 = _214_v105;
          _85_v1 = (_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))];
        } else {
          _85_v1 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(new BigNumber(611), _84_v0), _112_v25);
          _85_v1 = !((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))]) || (_85_v1);
          let _215_v106;
          let _nw32 = Array((new BigNumber(9)).toNumber()).fill(_dafny.MultiSet.Empty);
          _215_v106 = _nw32;
          let _index55 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_215_v106).length));
          (_215_v106)[_index55] = _181_v76;
          let _index56 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_215_v106).length));
          (_215_v106)[_index56] = ((_dafny.MultiSet.fromElements(_179_v74)).Union(_dafny.MultiSet.fromElements(_179_v74))).Intersect((_181_v76).update(_179_v74, _module.__default.abs((_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))])));
          (_91_globalState).f10 = ((_179_v74).f14).minus((_179_v74).f14);
          let _index57 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length));
          (_90_v6)[_index57] = false;
        }
      }
      let _index58 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length));
      (_90_v6)[_index58] = (_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))];
      let _216_v107;
      _216_v107 = _module.D2.create_DC5(_85_v1, (_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))]);
      let _217_v108;
      _217_v108 = _module.D2.create_DC7(_216_v107);
      let _source2 = ((_85_v1) ? (_217_v108) : (_217_v108));
      if (_source2.is_DC5) {
        let _218___mcc_h0 = (_source2).cf8;
        let _219___mcc_h1 = (_source2).cf9;
        let _220_cf9 = _219___mcc_h1;
        let _221_cf8 = _218___mcc_h0;
        let _222_v109;
        _222_v109 = _dafny.Map.Empty.slice().updateUnsafe(_107_v22,_179_v74);
        _222_v109 = (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('l'.codePointAt(0)),_179_v74)).update(_107_v22, _179_v74);
        let _223_v110;
        _223_v110 = _dafny.Map.Empty.slice().updateUnsafe(_84_v0,_module.D3.create_DC10(_module.D2.create_DC5(_220_cf9, _221_cf8), _107_v22, new BigNumber((_89_v5).length)));
        let _224_v111;
        _224_v111 = _module.D2.create_DC5(_220_cf9, _221_cf8);
        let _225_v112;
        _225_v112 = _module.D3.create_DC10(_224_v111, _107_v22, (_179_v74).f14);
        _223_v110 = (_223_v110).update((_179_v74).f14, _225_v112);
        _220_cf9 = !(!_dafny.areEqual(_89_v5, _dafny.Seq.UnicodeFromString("fglj")));
        _85_v1 = (_179_v74).fm5(_91_globalState);
      } else if (_source2.is_DC6) {
        let _226___mcc_h2 = (_source2).cf10;
        let _227___mcc_h3 = (_source2).cf11;
        let _228___mcc_h4 = (_source2).cf12;
        let _229_cf12 = _228___mcc_h4;
        let _230_cf11 = _227___mcc_h3;
        let _231_cf10 = _226___mcc_h2;
        let _init3 = ((_232_v25) => function (_233_i9) {
          return _module.__default.safeDivisionInt(_233_i9, new BigNumber((_232_v25).length));
        })(_112_v25);
        let _nw33 = Array((new BigNumber(29)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw33.length); _i0_3++) {
          _nw33[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _109_v24 = _nw33;
        let _234_v113;
        let _nw34 = new _module.C0();
        _nw34.__ctor((_179_v74).f14);
        _234_v113 = _nw34;
        let _235_v114;
        _235_v114 = _dafny.Map.Empty.slice().updateUnsafe(_85_v1,_230_cf11);
        let _index59 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length));
        (_90_v6)[_index59] = (((_235_v114).contains((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))])) ? ((_235_v114).get((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))])) : (false));
        _module.__default.m0(_85_v1, new BigNumber(644), _91_globalState);
      } else if (_source2.is_DC4) {
        let _236___mcc_h5 = (_source2).cf7;
        let _237_cf7 = _236___mcc_h5;
        if (_85_v1) {
          let _238_v115;
          _238_v115 = _module.D3.create_DC9((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))]);
          let _239_v116;
          _239_v116 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(((_179_v74).f14).plus(_84_v0)),_238_v115);
          let _pat_let_tv4 = _85_v1;
          _239_v116 = (_239_v116).update((_179_v74).f14, function (_pat_let6_0) {
            return function (_240_dt__update__tmp_h1) {
              return function (_pat_let7_0) {
                return function (_241_dt__update_hcf15_h0) {
                  return _module.D3.create_DC9(_241_dt__update_hcf15_h0);
                }(_pat_let7_0);
              }(!(_pat_let_tv4));
            }(_pat_let6_0);
          }(_238_v115));
          let _242_v117;
          let _nw35 = new _module.C0();
          _nw35.__ctor((_179_v74).f14);
          _242_v117 = _nw35;
          _85_v1 = _85_v1;
          let _243_v118;
          _243_v118 = _dafny.Map.Empty.slice().updateUnsafe(_84_v0,_109_v24);
          _243_v118 = ((_243_v118).update((_242_v117).f14, _109_v24)).Merge((_243_v118).Merge(_243_v118));
          _179_v74 = _179_v74;
        } else {
          _103_v18 = (_103_v18).update(new BigNumber((_dafny.Seq.of(_85_v1, (_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))])).length), (_179_v74).f14);
          let _244_v119;
          let _nw36 = Array((new BigNumber(2)).toNumber());
          _nw36[(_dafny.ZERO).toNumber()] = _178_v73;
          _nw36[(_dafny.ONE).toNumber()] = _178_v73;
          _244_v119 = _nw36;
          let _index60 = _module.__default.safeIndex(new BigNumber(43), new BigNumber((_244_v119).length));
          (_244_v119)[_index60] = (((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))]) ? (_178_v73) : (_178_v73));
          let _245_v120;
          _245_v120 = _module.D2.create_DC4(_237_cf7);
          let _246_v121;
          _246_v121 = _dafny.Set.fromElements(_245_v120, _245_v120, _245_v120, _245_v120, _module.D2.create_DC4(_237_cf7));
          let _index61 = _module.__default.safeIndex(new BigNumber(43), new BigNumber((_244_v119).length));
          let _index62 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length));
          let _rhs50 = _178_v73;
          let _rhs51 = (_179_v74).f14;
          let _rhs52 = ((_179_v74).f14).isEqualTo(_84_v0);
          let _rhs53 = new BigNumber(((_246_v121).Union(_246_v121)).length);
          let _rhs54 = _module.__default.safeDivisionInt((_179_v74).f14, new BigNumber((_89_v5).length));
          let _lhs49 = _244_v119;
          let _lhs50 = _module.__default.safeIndex(new BigNumber(43), new BigNumber((_244_v119).length));
          let _lhs51 = _109_v24;
          let _lhs52 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length));
          let _lhs53 = _91_globalState;
          let _lhs54 = _91_globalState;
          _lhs49[_lhs50] = _rhs50;
          _lhs51[_lhs52] = _rhs51;
          _85_v1 = _rhs52;
          _lhs53.f6 = _rhs53;
          _lhs54.f6 = _rhs54;
          _84_v0 = _module.__default.fm0(_91_globalState);
          _107_v22 = _107_v22;
          let _247_v122;
          let _nw37 = new _module.C0();
          _nw37.__ctor((_179_v74).f14);
          _247_v122 = _nw37;
        }
        let _248_v123;
        let _nw38 = new _module.C0();
        _nw38.__ctor((_dafny.ZERO).minus(_module.__default.safeModuloInt((_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))], new BigNumber(806))));
        _248_v123 = _nw38;
        if (((_dafny.MultiSet.fromElements(_179_v74)).Intersect(_181_v76)).IsProperSubsetOf(_181_v76)) {
          let _index63 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length));
          (_109_v24)[_index63] = (new BigNumber((_89_v5).length)).minus((_179_v74).f14);
          (_91_globalState).f6 = (_179_v74).f14;
          (_91_globalState).f6 = (_179_v74).f14;
          let _249_v124;
          let _nw39 = Array((new BigNumber(28)).toNumber()).fill(_module.D3.Default());
          _249_v124 = _nw39;
          let _250_v125;
          _250_v125 = _dafny.MultiSet.fromElements(_249_v124, _249_v124, _249_v124, _249_v124, _249_v124);
          let _rhs55 = _90_v6;
          let _rhs56 = (((_250_v125).contains((_module.D4.create_DC11(_249_v124)).dtor_cf19)) ? ((_250_v125).get((_module.D4.create_DC11(_249_v124)).dtor_cf19)) : (((_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))]).plus((_179_v74).f14)));
          let _rhs57 = (_248_v123).f14;
          let _rhs58 = (_179_v74).f14;
          let _lhs55 = _91_globalState;
          let _lhs56 = _91_globalState;
          let _lhs57 = _91_globalState;
          _90_v6 = _rhs55;
          _lhs55.f6 = _rhs56;
          _lhs56.f6 = _rhs57;
          _lhs57.f6 = _rhs58;
          let _251_v126;
          _251_v126 = _dafny.MultiSet.fromElements(_90_v6);
          (_91_globalState).f6 = _module.__default.safeDivisionInt(((_179_v74).f14).plus(new BigNumber((_251_v126).cardinality())), (_dafny.ZERO).minus((_dafny.ZERO).minus(((_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))]).multipliedBy(_84_v0))));
        } else {
          _module.__default.m0(_85_v1, (_dafny.ZERO).minus(((_248_v123).f14).plus(new BigNumber(254))), _91_globalState);
          _89_v5 = _89_v5;
          _85_v1 = (new BigNumber((_89_v5).length)).isLessThanOrEqualTo((_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))]);
          let _252_v127;
          _252_v127 = _dafny.MultiSet.fromElements((_179_v74).f14);
          let _rhs59 = ((((_252_v127).contains((_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))])) ? ((_252_v127).get((_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))])) : ((_248_v123).f14))).plus(((_179_v74).f14).multipliedBy(new BigNumber((_module.__default.fm7((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))], _107_v22, _91_globalState)).length)));
          let _rhs60 = (_dafny.ZERO).minus((_248_v123).f14);
          let _lhs58 = _91_globalState;
          let _lhs59 = _91_globalState;
          _lhs58.f6 = _rhs59;
          _lhs59.f6 = _rhs60;
          let _nw40 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
          _109_v24 = _nw40;
        }
        let _253_v128;
        _253_v128 = _dafny.Map.Empty.slice().updateUnsafe(_179_v74,_112_v25);
        let _254_v129;
        _254_v129 = _dafny.Seq.of((_253_v128).Merge(_253_v128));
        _253_v128 = (_254_v129)[_module.__default.safeIndex((_179_v74).f14, new BigNumber((_254_v129).length))];
      } else {
        let _255___mcc_h6 = (_source2).cf13;
        let _256_cf13 = _255___mcc_h6;
        let _nw41 = new _module.C0();
        _nw41.__ctor(_module.__default.safeDivisionInt((_179_v74).f14, _84_v0));
        _179_v74 = _nw41;
        let _index64 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length));
        let _rhs61 = ((_179_v74).f14).isLessThanOrEqualTo((_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))]);
        let _rhs62 = (_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))];
        let _lhs60 = _90_v6;
        let _lhs61 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length));
        let _lhs62 = _91_globalState;
        _lhs60[_lhs61] = _rhs61;
        _lhs62.f6 = _rhs62;
        _85_v1 = _dafny.areEqual(_dafny.Seq.of(_85_v1), _module.__default.fm8((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))], _91_globalState));
        if ((_87_v3)[_module.__default.safeIndex((_179_v74).f14, new BigNumber((_87_v3).length))]) {
          let _257_v130;
          _257_v130 = _module.D3.create_DC8(_dafny.Set.fromElements(_107_v22, _107_v22, _107_v22, _107_v22, new _dafny.CodePoint('v'.codePointAt(0))));
          let _258_v131;
          _258_v131 = _dafny.Set.fromElements(_107_v22, _107_v22);
          let _pat_let_tv5 = _258_v131;
          _257_v130 = function (_pat_let8_0) {
            return function (_259_dt__update__tmp_h2) {
              return function (_pat_let9_0) {
                return function (_260_dt__update_hcf14_h0) {
                  return _module.D3.create_DC8(_260_dt__update_hcf14_h0);
                }(_pat_let9_0);
              }(_pat_let_tv5);
            }(_pat_let8_0);
          }(_257_v130);
          _257_v130 = (((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))]) ? (_257_v130) : (_257_v130));
          let _index65 = _module.__default.safeIndex(new BigNumber(437), new BigNumber((_178_v73).length));
          (_178_v73)[_index65] = _89_v5;
          let _index66 = _module.__default.safeIndex(new BigNumber(437), new BigNumber((_178_v73).length));
          (_178_v73)[_index66] = _dafny.Seq.Concat(_89_v5, _89_v5);
          let _261_v132;
          let _nw42 = Array((new BigNumber(15)).toNumber());
          _261_v132 = _nw42;
          _261_v132 = _261_v132;
          (_91_globalState).f6 = (_179_v74).f14;
        } else {
          (_91_globalState).f6 = (_dafny.ZERO).minus((_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))]);
          let _262_v133;
          _262_v133 = _dafny.Seq.of(_109_v24, _109_v24, _109_v24, _109_v24);
          let _263_v134;
          _263_v134 = _dafny.MultiSet.fromElements((_262_v133)[_module.__default.safeIndex((_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))], new BigNumber((_262_v133).length))]);
          let _index67 = _module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length));
          (_90_v6)[_index67] = (_dafny.MultiSet.fromElements(_109_v24)).IsProperSubsetOf(_263_v134);
          let _index68 = _module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length));
          (_109_v24)[_index68] = new BigNumber((_dafny.Seq.Concat(_89_v5, _module.__default.fm7((_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))], _module.__default.fm3((_109_v24)[_module.__default.safeIndex(new BigNumber(76), new BigNumber((_109_v24).length))], new BigNumber(-30), (_90_v6)[_module.__default.safeIndex(new BigNumber(506), new BigNumber((_90_v6).length))], _91_globalState), _91_globalState))).length);
          (_91_globalState).f10 = (_dafny.ZERO).minus(_84_v0);
          let _264_v135;
          _264_v135 = _dafny.Map.Empty.slice().updateUnsafe(_182_v77,_module.__default.fm0(_91_globalState));
          _264_v135 = (_264_v135).update(_182_v77, (_179_v74).f14);
        }
      }
      _module.__default.m0(false, (_179_v74).f14, _91_globalState);
      process.stdout.write(_dafny.toString(_84_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_85_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_86_v2).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(457),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_87_v3, _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_88_v4).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(924),_dafny.Seq.of(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_89_v5).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_90_v6)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_90_v6)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_90_v6)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_90_v6)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_90_v6)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_90_v6)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_90_v6)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_90_v6)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_90_v6)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_90_v6)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_90_v6)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_90_v6)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_90_v6)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_90_v6)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_91_globalState).f4).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(924),_dafny.Seq.of(false)).updateUnsafe(new BigNumber(457),_dafny.Seq.of(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_91_globalState.f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_91_globalState).f8).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_91_globalState).f9).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_91_globalState.f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_91_globalState).f11)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_91_globalState).f11)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_91_globalState).f11)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_91_globalState).f11)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_91_globalState).f11)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_91_globalState).f11)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_91_globalState).f11)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_91_globalState).f11)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_91_globalState).f11)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_91_globalState).f11)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_91_globalState).f11)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_91_globalState).f11)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_91_globalState).f11)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_91_globalState).f11)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_globalState).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_103_v18).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(457),new BigNumber(457)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_104_v20, _dafny.Seq.of(_dafny.Map.Empty.slice()))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v21).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,_dafny.Seq.of(_dafny.Map.Empty.slice())))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_107_v22));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_108_v23).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("i"),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_109_v24)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_112_v25, _dafny.Seq.of(new BigNumber(457), new BigNumber(457), new BigNumber(457)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_119_v30).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new _dafny.CodePoint('n'.codePointAt(0))).updateUnsafe(new BigNumber(457),new _dafny.CodePoint('n'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_179_v74).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_180_v75).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_181_v76).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_v77).equals(_dafny.Set.fromElements(_dafny.ONE, new BigNumber(457), new BigNumber(-457)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_216_v107).dtor_cf8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_216_v107).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_217_v108).dtor_cf13).dtor_cf8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_217_v108).dtor_cf13).dtor_cf9));
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
    static create_DC2(cf2, cf3, cf4, cf5) {
      let $dt = new D1(0);
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC1(cf1) {
      let $dt = new D1(1);
      $dt.cf1 = cf1;
      return $dt;
    }
    static create_DC3(cf6) {
      let $dt = new D1(2);
      $dt.cf6 = cf6;
      return $dt;
    }
    get is_DC2() { return this.$tag === 0; }
    get is_DC1() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf6() { return this.cf6; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf6) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf2, other.cf2) && _dafny.areEqual(this.cf3, other.cf3) && _dafny.areEqual(this.cf4, other.cf4) && _dafny.areEqual(this.cf5, other.cf5);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf1, other.cf1);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf6, other.cf6);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC2(_dafny.ZERO, _dafny.ZERO, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC5(cf8, cf9) {
      let $dt = new D2(0);
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC6(cf10, cf11, cf12) {
      let $dt = new D2(1);
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC4(cf7) {
      let $dt = new D2(2);
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC7(cf13) {
      let $dt = new D2(3);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get is_DC4() { return this.$tag === 2; }
    get is_DC7() { return this.$tag === 3; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC4" + "(" + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 3) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf13) + ")";
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
        return other.$tag === 1 && this.cf10 === other.cf10 && this.cf11 === other.cf11 && _dafny.areEqual(this.cf12, other.cf12);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf7, other.cf7);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf13, other.cf13);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC5(false, false);
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
    static create_DC10(cf16, cf17, cf18) {
      let $dt = new D3(1);
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC8(cf14) {
      let $dt = new D3(2);
      $dt.cf14 = cf14;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf14() { return this.cf14; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf14) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf15 === other.cf15;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf16, other.cf16) && _dafny.areEqual(this.cf17, other.cf17) && _dafny.areEqual(this.cf18, other.cf18);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf14, other.cf14);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC9(false);
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
    static create_DC12(cf20, cf21, cf22, cf23) {
      let $dt = new D4(0);
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC11(cf19) {
      let $dt = new D4(1);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf19) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf20 === other.cf20 && this.cf21 === other.cf21 && this.cf22 === other.cf22 && _dafny.areEqual(this.cf23, other.cf23);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf19 === other.cf19;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC12(false, false, false, _dafny.Seq.of());
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
      this.f6 = _dafny.ZERO;
      this.f10 = _dafny.ZERO;
      this._f0 = false;
      this._f1 = false;
      this._f2 = false;
      this._f3 = false;
      this._f4 = _dafny.Map.Empty;
      this._f5 = _dafny.ZERO;
      this._f7 = false;
      this._f8 = _dafny.Seq.UnicodeFromString("");
      this._f9 = _dafny.Seq.UnicodeFromString("");
      this._f11 = [];
      this._f12 = false;
      this._f13 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this).f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this).f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
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
    get f11() {
      let _this = this;
      return _this._f11;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f14 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f14) {
      let _this = this;
      (_this)._f14 = f14;
      return;
    }
    fm4(p0, globalState) {
      let _this = this;
      return (true);
    };
    fm5(globalState) {
      let _this = this;
      return (_dafny.MultiSet.fromElements(!(true), false, !(false), true)).IsDisjointFrom((_dafny.MultiSet.fromElements(true, false)).Union(_dafny.MultiSet.fromElements(true)));
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
