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
      return (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-469),false)).length)).multipliedBy(_module.__default.safeDivisionInt(new BigNumber(-411), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(94)), function (_0_i0) {
        return new BigNumber(366);
      }))).cardinality()))).length)));
    };
    static fm1(p0, p1, p2, p3, globalState) {
      return ((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(850), new BigNumber(668))) {
          let _1_v0 = _compr_0;
          if (((new BigNumber(850)).isLessThanOrEqualTo(_1_v0)) && ((_1_v0).isLessThan(new BigNumber(668)))) {
            _coll0.push([_module.__default.safeModuloInt(_1_v0, new BigNumber(-868)),new BigNumber((function () {
              let _coll1 = new _dafny.Set();
              for (const _compr_1 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('b'.codePointAt(0)),false)).Keys.Elements) {
                let _2_v1 = _compr_1;
                if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('b'.codePointAt(0)),false)).contains(_2_v1)) {
                  _coll1.add(_2_v1);
                }
              }
              return _coll1;
            }()).length)]);
          }
        }
        return _coll0;
      }())).length))).Union(_dafny.Set.fromElements(new BigNumber(176)))).Union((_dafny.Set.fromElements(new BigNumber(166))).Union(_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("henpihgvp")).length)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(91)), function (_3_i0) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      })).length), new BigNumber(936))).length)))));
    };
    static fm3(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(76)), function (_4_i0) {
        return _module.D1.create_DC4(_module.D1.create_DC3(true, true, false, !(!(!(true)))));
      });
    };
    static fm4(p0, globalState) {
      return _dafny.MultiSet.fromElements((_module.D4.create_DC11(new _dafny.CodePoint('b'.codePointAt(0)), true, _dafny.MultiSet.fromElements(new _dafny.CodePoint('t'.codePointAt(0))), true, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-835)), function (_5_i0) {
  return _dafny.MultiSet.fromElements(new BigNumber(493));
}))).dtor_cf20);
    };
    static fm5(globalState) {
      return new _dafny.CodePoint('g'.codePointAt(0));
    };
    static fm6(p0, p1, p2, globalState) {
      let _source0 = _module.D3.create_DC7();
      if (_source0.is_DC7) {
        return false;
      } else if (_source0.is_DC8) {
        let _6___mcc_h0 = (_source0).cf14;
        let _7_cf14 = _6___mcc_h0;
        return false;
      } else {
        let _8___mcc_h1 = (_source0).cf13;
        let _9_cf13 = _8___mcc_h1;
        return !(!(false));
      }
    };
    static fm7(p0, p1, p2, globalState) {
      return _dafny.Seq.of(new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)));
    };
    static fm8(p0, globalState) {
      return _dafny.Seq.Concat(((false) ? (_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(false, false)).length))) : (_dafny.Seq.of(new BigNumber((function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of _dafny.IntegerRange(new BigNumber(939), new BigNumber(466))) {
          let _10_v0 = _compr_2;
          if (((new BigNumber(939)).isLessThanOrEqualTo(_10_v0)) && ((_10_v0).isLessThan(new BigNumber(466)))) {
            _coll2.push([(_10_v0).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length)),false]);
          }
        }
        return _coll2;
      }()).length), new BigNumber(844), new BigNumber(518)))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(681)), function (_11_i0) {
        return new BigNumber(244);
      }));
    };
    static fm9(p0, p1, globalState) {
      return _module.D3.create_DC7();
    };
    static fm10(p0, p1, p2, p3, globalState) {
      let _source1 = _module.D4.create_DC11(new _dafny.CodePoint('h'.codePointAt(0)), false, _dafny.MultiSet.fromElements(new _dafny.CodePoint('m'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0))), true, _dafny.Seq.Create(_module.__default.abs(new BigNumber(271)), function (_12_i0) {
  return _dafny.MultiSet.fromElements(new BigNumber((function () {
    let _coll3 = new _dafny.Map();
    for (const _compr_3 of _dafny.IntegerRange(new BigNumber(421), new BigNumber(551))) {
      let _13_v0 = _compr_3;
      if (((new BigNumber(421)).isLessThanOrEqualTo(_13_v0)) && ((_13_v0).isLessThan(new BigNumber(551)))) {
        _coll3.push([(_13_v0).minus(new BigNumber((_dafny.Seq.UnicodeFromString("dxbtxxe")).length)),new BigNumber((_dafny.Seq.UnicodeFromString("gi")).length)]);
      }
    }
    return _coll3;
  }()).length));
}));
      if (_source1.is_DC10) {
        let _14___mcc_h0 = (_source1).cf16;
        let _15___mcc_h1 = (_source1).cf17;
        let _16___mcc_h2 = (_source1).cf18;
        let _17_cf18 = _16___mcc_h2;
        let _18_cf17 = _15___mcc_h1;
        let _19_cf16 = _14___mcc_h0;
        return _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_18_cf17,true));
      } else if (_source1.is_DC11) {
        let _20___mcc_h3 = (_source1).cf19;
        let _21___mcc_h4 = (_source1).cf20;
        let _22___mcc_h5 = (_source1).cf21;
        let _23___mcc_h6 = (_source1).cf22;
        let _24___mcc_h7 = (_source1).cf23;
        let _25_cf23 = _24___mcc_h7;
        let _26_cf22 = _23___mcc_h6;
        let _27_cf21 = _22___mcc_h5;
        let _28_cf20 = _21___mcc_h4;
        let _29_cf19 = _20___mcc_h3;
        return (_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_26_cf22,_26_cf22), _dafny.Map.Empty.slice().updateUnsafe(_28_cf20,_28_cf20))).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,_26_cf22), _dafny.Map.Empty.slice().updateUnsafe(_28_cf20,_26_cf22))));
      } else if (_source1.is_DC12) {
        let _30___mcc_h8 = (_source1).cf24;
        let _31___mcc_h9 = (_source1).cf25;
        let _32___mcc_h10 = (_source1).cf26;
        let _33___mcc_h11 = (_source1).cf27;
        let _34_cf27 = _33___mcc_h11;
        let _35_cf26 = _32___mcc_h10;
        let _36_cf25 = _31___mcc_h9;
        let _37_cf24 = _30___mcc_h8;
        return _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_35_cf26,_35_cf26), _dafny.Map.Empty.slice().updateUnsafe(_35_cf26,(_module.D4.create_DC10(_37_cf24, false, _35_cf26)).dtor_cf18));
      } else {
        let _38___mcc_h12 = (_source1).cf15;
        let _39_cf15 = _38___mcc_h12;
        return _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,true)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(632)), function (_40_i1) {
          return _dafny.Map.Empty.slice().updateUnsafe(true,false);
        })));
      }
    };
    static m0(p0, p1, p2, p3, globalState) {
      let r0 = _dafny.Set.Empty;
      let r1 = false;
      let r2 = _dafny.MultiSet.Empty;
      if (p1) {
        (globalState).f15 = _module.__default.safeModuloInt(p0, _module.__default.safeDivisionInt(p0, new BigNumber(639)));
        let _41_v0;
        _41_v0 = _dafny.Seq.of(p0, p0, (_dafny.ZERO).minus(_module.__default.fm0(p1, p3, p2, p0, globalState)), (p0).minus(p0));
        let _42_v1;
        _42_v1 = new _dafny.CodePoint('y'.codePointAt(0));
        let _43_v2;
        let _nw0 = new _module.C0();
        _nw0.__ctor(p0, _42_v1);
        _43_v2 = _nw0;
        let _44_v3;
        _44_v3 = _dafny.Seq.UnicodeFromString("deyue");
        let _45_v4;
        _45_v4 = _module.D0.create_DC0(_44_v3);
        let _rhs0 = _module.__default.fm8(!(p2) || (p2), globalState);
        let _rhs1 = _dafny.areEqual(_dafny.Seq.Concat(_module.__default.fm7(false, (_43_v2).f19, new BigNumber((_41_v0).length), globalState), _44_v3), ((!(p2)) ? (_44_v3) : ((_45_v4).dtor_cf0)));
        let _rhs2 = (_43_v2).f19;
        let _rhs3 = _43_v2;
        let _rhs4 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("hkykdidt"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(74)), ((_46_v1) => function (_47_i0) {
          return _46_v1;
        })(_42_v1)));
        let _lhs0 = globalState;
        let _lhs1 = globalState;
        _41_v0 = _rhs0;
        _lhs0.f6 = _rhs1;
        _lhs1.f13 = _rhs2;
        _43_v2 = _rhs3;
        _44_v3 = _rhs4;
        let _48_v5;
        let _init0 = ((_49_v2) => function (_50_i1) {
          return _module.__default.safeModuloInt(_50_i1, (_49_v2).f19);
        })(_43_v2);
        let _nw1 = Array((new BigNumber(8)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw1.length); _i0_0++) {
          _nw1[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _48_v5 = _nw1;
        let _51_v6;
        _51_v6 = _dafny.Map.Empty.slice().updateUnsafe(p2,_48_v5);
        let _52_v7;
        _52_v7 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
        let _53_v8;
        _53_v8 = _dafny.Map.Empty.slice().updateUnsafe((((_51_v6).contains(true)) ? ((_51_v6).get(true)) : (_48_v5)),_52_v7);
        _53_v8 = (_53_v8).update(_48_v5, _52_v7);
        let _54_v9;
        _54_v9 = _dafny.Map.Empty.slice().updateUnsafe(_41_v0,_48_v5);
        _54_v9 = (_54_v9).Merge(_54_v9);
        let _55_v10;
        let _nw2 = Array((new BigNumber(14)).toNumber());
        _nw2[(_dafny.ZERO).toNumber()] = _42_v1;
        _nw2[(_dafny.ONE).toNumber()] = _42_v1;
        _nw2[(new BigNumber(2)).toNumber()] = (p3)[_module.__default.safeIndex((_43_v2).f19, new BigNumber((p3).length))];
        _nw2[(new BigNumber(3)).toNumber()] = (_43_v2).f20;
        _nw2[(new BigNumber(4)).toNumber()] = (_43_v2).f20;
        _nw2[(new BigNumber(5)).toNumber()] = (_43_v2).f20;
        _nw2[(new BigNumber(6)).toNumber()] = _42_v1;
        _nw2[(new BigNumber(7)).toNumber()] = new _dafny.CodePoint('c'.codePointAt(0));
        _nw2[(new BigNumber(8)).toNumber()] = _42_v1;
        _nw2[(new BigNumber(9)).toNumber()] = _42_v1;
        _nw2[(new BigNumber(10)).toNumber()] = (_43_v2).f20;
        _nw2[(new BigNumber(11)).toNumber()] = (_43_v2).f20;
        _nw2[(new BigNumber(12)).toNumber()] = _42_v1;
        _nw2[(new BigNumber(13)).toNumber()] = (_43_v2).f20;
        _55_v10 = _nw2;
        let _index0 = _module.__default.safeIndex(new BigNumber(614), new BigNumber((_55_v10).length));
        (_55_v10)[_index0] = _42_v1;
        let _index1 = _module.__default.safeIndex(new BigNumber(614), new BigNumber((_55_v10).length));
        (_55_v10)[_index1] = (_43_v2).f20;
      } else {
        let _56_v11;
        _56_v11 = new _dafny.CodePoint('a'.codePointAt(0));
        (globalState).f6 = _dafny.Seq.contains(p3, ((p2) ? (_56_v11) : (_module.__default.fm5(globalState))));
        let _57_v12;
        _57_v12 = _module.D3.create_DC7();
        _57_v12 = _57_v12;
        let _58_v13;
        _58_v13 = _dafny.MultiSet.fromElements(((p1) ? (!(p1)) : (p1)), (false) && (p1));
        _58_v13 = _module.__default.fm4(p2, globalState);
        (globalState).f13 = (_dafny.ZERO).minus(p0);
        let _59_v14;
        let _init1 = ((_60_p0) => function (_61_i2) {
          return _module.__default.safeModuloInt(_61_i2, _60_p0);
        })(p0);
        let _nw3 = Array((new BigNumber(18)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw3.length); _i0_1++) {
          _nw3[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _59_v14 = _nw3;
        let _62_v15;
        let _nw4 = Array((new BigNumber(19)).toNumber());
        _nw4[(_dafny.ZERO).toNumber()] = p2;
        _nw4[(_dafny.ONE).toNumber()] = !(p1);
        _nw4[(new BigNumber(2)).toNumber()] = p2;
        _nw4[(new BigNumber(3)).toNumber()] = p2;
        _nw4[(new BigNumber(4)).toNumber()] = p1;
        _nw4[(new BigNumber(5)).toNumber()] = p1;
        _nw4[(new BigNumber(6)).toNumber()] = p1;
        _nw4[(new BigNumber(7)).toNumber()] = p2;
        _nw4[(new BigNumber(8)).toNumber()] = p2;
        _nw4[(new BigNumber(9)).toNumber()] = true;
        _nw4[(new BigNumber(10)).toNumber()] = p2;
        _nw4[(new BigNumber(11)).toNumber()] = !(p1);
        _nw4[(new BigNumber(12)).toNumber()] = p2;
        _nw4[(new BigNumber(13)).toNumber()] = false;
        _nw4[(new BigNumber(14)).toNumber()] = p2;
        _nw4[(new BigNumber(15)).toNumber()] = p1;
        _nw4[(new BigNumber(16)).toNumber()] = p2;
        _nw4[(new BigNumber(17)).toNumber()] = p2;
        _nw4[(new BigNumber(18)).toNumber()] = p1;
        _62_v15 = _nw4;
        let _63_v16;
        _63_v16 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(691)), ((_64_v11) => function (_65_i3) {
          return _64_v11;
        })(_56_v11)), p3);
        let _66_v17;
        _66_v17 = _module.D0.create_DC1(_59_v14, _62_v15, _dafny.MultiSet.FromArray(_dafny.Seq.update(_63_v16, _module.__default.safeIndex(p0, new BigNumber((_63_v16).length)), _dafny.Seq.of(_56_v11))), p0, p0);
        let _index2 = _module.__default.safeIndex(new BigNumber(248), new BigNumber((_59_v14).length));
        (_59_v14)[_index2] = (_66_v17).dtor_cf4;
        let _67_v18;
        _67_v18 = _module.D1.create_DC3(p2, p1, true, _module.__default.fm6(p0, _module.__default.fm5(globalState), (_dafny.ZERO).minus(p0), globalState));
        let _index3 = _module.__default.safeIndex(new BigNumber(248), new BigNumber((_59_v14).length));
        (_59_v14)[_index3] = (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(p1, (_67_v18).dtor_cf8)).cardinality()));
      }
      let _68_v19;
      let _nw5 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
      _68_v19 = _nw5;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_68_v19).length))) {
        let _69_i4 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_69_i4)) && ((_69_i4).isLessThan(new BigNumber((_68_v19).length))))) {
          (_68_v19)[(_69_i4)] = _module.__default.safeModuloInt(_69_i4, p0);
        }
      }
      let _hi0 = p0;
      for (let _70_i5 = p0; _70_i5.isLessThan(_hi0); _70_i5 = _70_i5.plus(_dafny.ONE)) {
        let _71_v20;
        let _nw6 = Array((new BigNumber(13)).toNumber()).fill(false);
        _71_v20 = _nw6;
        let _72_v21;
        _72_v21 = _dafny.Set.fromElements(_71_v20);
        _72_v21 = _72_v21;
        let _73_v22;
        let _nw7 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Map.Empty);
        _73_v22 = _nw7;
        _73_v22 = _73_v22;
        let _74_v23;
        _74_v23 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        let _source2 = _module.__default.fm9(true, new BigNumber(((_74_v23).Merge(_74_v23)).length), globalState);
        if (_source2.is_DC7) {
          (globalState).f6 = true;
          let _75_v24;
          _75_v24 = _dafny.Seq.of(p1);
          let _76_v25;
          _76_v25 = _dafny.Seq.of(_dafny.Seq.update(_dafny.Seq.update(p3, _module.__default.safeIndex(p0, new BigNumber((p3).length)), new _dafny.CodePoint('i'.codePointAt(0))), _module.__default.safeIndex(_70_i5, new BigNumber((_dafny.Seq.update(p3, _module.__default.safeIndex(p0, new BigNumber((p3).length)), new _dafny.CodePoint('i'.codePointAt(0)))).length)), _module.__default.fm5(globalState)), p3, _dafny.Seq.Concat(p3, p3), _module.__default.fm7(!((_75_v24)[_module.__default.safeIndex(_70_i5, new BigNumber((_75_v24).length))]), p0, p0, globalState));
          _76_v25 = _dafny.Seq.Concat(_76_v25, _76_v25);
          let _77_v26;
          _77_v26 = _dafny.Map.Empty.slice().updateUnsafe(p3,p0);
          let _78_v27;
          _78_v27 = _dafny.Map.Empty.slice().updateUnsafe((p2) || (p1),_77_v26);
          _78_v27 = (_78_v27).update(p1, (_77_v26).Merge(_dafny.Map.Empty.slice().updateUnsafe(p3,p0)));
          let _79_v28;
          _79_v28 = new _dafny.CodePoint('w'.codePointAt(0));
          let _80_v29;
          let _nw8 = new _module.C0();
          _nw8.__ctor(p0, _79_v28);
          _80_v29 = _nw8;
          let _81_v30;
          _81_v30 = _module.D4.create_DC9(_80_v29);
          let _82_v31;
          let _nw9 = Array((new BigNumber(19)).toNumber());
          _nw9[(_dafny.ZERO).toNumber()] = _80_v29;
          _nw9[(_dafny.ONE).toNumber()] = _80_v29;
          _nw9[(new BigNumber(2)).toNumber()] = _80_v29;
          _nw9[(new BigNumber(3)).toNumber()] = _80_v29;
          _nw9[(new BigNumber(4)).toNumber()] = (_81_v30).dtor_cf15;
          _nw9[(new BigNumber(5)).toNumber()] = _80_v29;
          _nw9[(new BigNumber(6)).toNumber()] = _80_v29;
          _nw9[(new BigNumber(7)).toNumber()] = _80_v29;
          _nw9[(new BigNumber(8)).toNumber()] = _80_v29;
          _nw9[(new BigNumber(9)).toNumber()] = _80_v29;
          _nw9[(new BigNumber(10)).toNumber()] = _80_v29;
          _nw9[(new BigNumber(11)).toNumber()] = _80_v29;
          _nw9[(new BigNumber(12)).toNumber()] = _80_v29;
          _nw9[(new BigNumber(13)).toNumber()] = _80_v29;
          _nw9[(new BigNumber(14)).toNumber()] = _80_v29;
          _nw9[(new BigNumber(15)).toNumber()] = _80_v29;
          _nw9[(new BigNumber(16)).toNumber()] = _80_v29;
          _nw9[(new BigNumber(17)).toNumber()] = _80_v29;
          _nw9[(new BigNumber(18)).toNumber()] = _80_v29;
          _82_v31 = _nw9;
          let _83_v32;
          _83_v32 = _dafny.Set.fromElements(_82_v31, _82_v31, _82_v31, _82_v31, _82_v31);
          _83_v32 = _dafny.Set.fromElements(((p2) ? (_82_v31) : (_82_v31)));
        } else if (_source2.is_DC8) {
          let _84___mcc_h0 = (_source2).cf14;
          let _85_cf14 = _84___mcc_h0;
          let _86_v33;
          _86_v33 = new _dafny.CodePoint('c'.codePointAt(0));
          let _87_v34;
          _87_v34 = _dafny.Map.Empty.slice().updateUnsafe(_86_v33,p1);
          let _88_v35;
          _88_v35 = _dafny.Map.Empty.slice().updateUnsafe(_71_v20,_87_v34);
          let _89_v36;
          _89_v36 = _dafny.Seq.of(((p1) ? (_dafny.Map.Empty.slice().updateUnsafe(_71_v20,_87_v34)) : (_88_v35)));
          let _90_v37;
          let _nw10 = new _module.C0();
          _nw10.__ctor(_85_cf14, _86_v33);
          _90_v37 = _nw10;
          let _91_v38;
          _91_v38 = _dafny.Map.Empty.slice().updateUnsafe(_90_v37,_70_i5);
          let _rhs5 = (_89_v36)[_module.__default.safeIndex(_85_cf14, new BigNumber((_89_v36).length))];
          let _rhs6 = (new BigNumber(((_91_v38).update(_90_v37, _85_cf14)).length)).isLessThan(new BigNumber(908));
          _88_v35 = _rhs5;
          r1 = _rhs6;
          let _92_v39;
          _92_v39 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(470),new BigNumber(((_module.D0.create_DC0(p3)).dtor_cf0).length));
          r1 = (((p1) === (true)) ? ((new BigNumber((_92_v39).length)).isLessThan((_90_v37).f19)) : (p2));
          _71_v20 = _71_v20;
          (globalState).f15 = _85_cf14;
        } else {
          let _93___mcc_h1 = (_source2).cf13;
          let _94_cf13 = _93___mcc_h1;
          let _95_v40;
          _95_v40 = new _dafny.CodePoint('t'.codePointAt(0));
          let _96_v41;
          let _nw11 = new _module.C0();
          _nw11.__ctor(p0, _95_v40);
          _96_v41 = _nw11;
          let _97_v42;
          _97_v42 = _dafny.Seq.of(p0, new BigNumber((p3).length), _70_i5);
          let _98_v43;
          _98_v43 = _dafny.Map.Empty.slice().updateUnsafe(_97_v42,_96_v41);
          let _99_v44;
          _99_v44 = _dafny.MultiSet.fromElements((_96_v41).f20, (_96_v41).f20);
          let _100_v45;
          _100_v45 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(!(p2))).length), new BigNumber(328));
          let _101_v46;
          _101_v46 = _dafny.Seq.of(_100_v45, _100_v45);
          let _102_v47;
          _102_v47 = _module.D4.create_DC11((_96_v41).f20, p1, _99_v44, p1, _101_v46);
          let _rhs7 = (_102_v47).dtor_cf22;
          let _rhs8 = _dafny.Map.Empty.slice().updateUnsafe(_97_v42,_96_v41);
          r1 = _rhs7;
          _98_v43 = _rhs8;
          (globalState).f5 = (_dafny.ZERO).minus((((_96_v41).f19).multipliedBy((_96_v41).f19)).plus(_70_i5));
          _97_v42 = _dafny.Seq.Concat(_dafny.Seq.Concat(_97_v42, _97_v42), ((p2) ? (_97_v42) : (_dafny.Seq.of((_96_v41).f19))));
        }
        let _103_v48;
        _103_v48 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
        let _104_v49;
        _104_v49 = new _dafny.CodePoint('c'.codePointAt(0));
        let _index4 = _module.__default.safeIndex(new BigNumber(518), new BigNumber((_68_v19).length));
        (_68_v19)[_index4] = _module.__default.safeDivisionInt(_70_i5, _module.__default.fm0(p1, p3, _module.__default.fm6(new BigNumber((_103_v48).length), _104_v49, new BigNumber((_dafny.Seq.of(p2)).length), globalState), p0, globalState));
        let _index5 = _module.__default.safeIndex(new BigNumber(518), new BigNumber((_68_v19).length));
        (_68_v19)[_index5] = (_dafny.ZERO).minus(p0);
      }
      let _105_v50;
      let _init2 = ((_106_p3) => function (_107_i6) {
        return _106_p3;
      })(p3);
      let _nw12 = Array((new BigNumber(25)).toNumber());
      for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw12.length); _i0_2++) {
        _nw12[_i0_2] = _init2(new BigNumber(_i0_2));
      }
      _105_v50 = _nw12;
      let _index6 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_105_v50).length));
      (_105_v50)[_index6] = _dafny.Seq.UnicodeFromString("d");
      let _index7 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_105_v50).length));
      (_105_v50)[_index7] = _dafny.Seq.UnicodeFromString("arixobu");
      let _108_v51;
      let _nw13 = Array((new BigNumber(3)).toNumber()).fill(false);
      _108_v51 = _nw13;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_108_v51).length))) {
        let _109_i7 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_109_i7)) && ((_109_i7).isLessThan(new BigNumber((_108_v51).length))))) {
          (_108_v51)[(_109_i7)] = p1;
        }
      }
      let _110_v52;
      _110_v52 = new _dafny.CodePoint('x'.codePointAt(0));
      let _111_v53;
      let _nw14 = new _module.C0();
      _nw14.__ctor(p0, _110_v52);
      _111_v53 = _nw14;
      _111_v53 = _111_v53;
      let _112_v54;
      _112_v54 = _dafny.Set.fromElements((_111_v53).f19, p0);
      let _113_v56;
      _113_v56 = _dafny.Set.fromElements(_112_v54, (_112_v54).Union(function () {
        let _coll4 = new _dafny.Set();
        for (const _compr_4 of _dafny.IntegerRange(new BigNumber(-145), new BigNumber(464))) {
          let _114_v55 = _compr_4;
          if (((new BigNumber(-145)).isLessThanOrEqualTo(_114_v55)) && ((_114_v55).isLessThan(new BigNumber(464)))) {
            _coll4.add(_module.__default.safeDivisionInt(_114_v55, new BigNumber((_dafny.Seq.of(p2, p1, true, p1, p1)).length)));
          }
        }
        return _coll4;
      }()), _module.__default.fm1(p1, (_111_v53).f19, p0, new BigNumber(503), globalState));
      r0 = _113_v56;
      r1 = p1;
      let _115_v57;
      _115_v57 = _module.D4.create_DC10(p0, p2, p1);
      let _116_v58;
      _116_v58 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(828));
      let _117_v60;
      _117_v60 = _dafny.MultiSet.fromElements(_116_v58, function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of _dafny.IntegerRange(new BigNumber(335), new BigNumber(915))) {
          let _118_v59 = _compr_5;
          if (((new BigNumber(335)).isLessThanOrEqualTo(_118_v59)) && ((_118_v59).isLessThan(new BigNumber(915)))) {
            _coll5.push([(_118_v59).minus(new BigNumber((_112_v54).length)),(_111_v53).f19]);
          }
        }
        return _coll5;
      }());
      r2 = _module.__default.fm10(p0, (new BigNumber(435)).minus((_115_v57).dtor_cf16), _117_v60, p0, globalState);
      return [r0, r1, r2];
    }
    static Main(__noArgsParameter) {
      let _119_v0;
      _119_v0 = new BigNumber(327);
      let _120_v1;
      _120_v1 = _dafny.Seq.UnicodeFromString("qljvqw");
      let _121_v2;
      _121_v2 = true;
      let _122_v3;
      _122_v3 = _dafny.Seq.of(_121_v2);
      let _123_v4;
      _123_v4 = _dafny.MultiSet.fromElements(_120_v1, _120_v1, _120_v1);
      let _124_v5;
      let _init3 = ((_125_v2) => function (_126_i0) {
        return _125_v2;
      })(_121_v2);
      let _nw15 = Array((new BigNumber(29)).toNumber());
      for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw15.length); _i0_3++) {
        _nw15[_i0_3] = _init3(new BigNumber(_i0_3));
      }
      _124_v5 = _nw15;
      let _127_globalState;
      let _nw16 = new _module.GlobalState();
      _nw16.__ctor(new BigNumber(624), new BigNumber(743), false, new BigNumber(-990), false, new BigNumber(913), false, _dafny.Seq.of(_119_v0, new BigNumber((_120_v1).length), _119_v0, (_dafny.ZERO).minus(_119_v0)), _dafny.Map.Empty.slice().updateUnsafe(_122_v3,_121_v2), new BigNumber(580), new BigNumber(647), new _dafny.CodePoint('r'.codePointAt(0)), _123_v4, new BigNumber(481), _dafny.Seq.of(_124_v5, _124_v5), new BigNumber(31), true, new BigNumber(-830), false);
      _127_globalState = _nw16;
      let _128_v6;
      _128_v6 = new _dafny.CodePoint('m'.codePointAt(0));
      let _129_v7;
      _129_v7 = _dafny.Seq.of(new BigNumber(-349));
      let _130_v8;
      _130_v8 = _dafny.Set.fromElements(_119_v0, (_129_v7)[_module.__default.safeIndex(new BigNumber(-857), new BigNumber((_129_v7).length))], _119_v0, _119_v0, _119_v0);
      let _131_v9;
      _131_v9 = _dafny.MultiSet.fromElements(_module.__default.fm1(_121_v2, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true, _121_v2, true, _121_v2, _121_v2)).length)), _119_v0, (_dafny.ZERO).minus(_119_v0), _127_globalState), _130_v8);
      let _132_v11;
      _132_v11 = _dafny.Map.Empty.slice().updateUnsafe(_119_v0,(_dafny.ZERO).minus(_119_v0));
      let _133_v12;
      let _nw17 = Array((new BigNumber(14)).toNumber());
      _nw17[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_119_v0,_121_v2)).length);
      _nw17[(_dafny.ONE).toNumber()] = _module.__default.safeModuloInt(_119_v0, _119_v0);
      _nw17[(new BigNumber(2)).toNumber()] = _119_v0;
      _nw17[(new BigNumber(3)).toNumber()] = new BigNumber(-498);
      _nw17[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Seq.update(_120_v1, _module.__default.safeIndex(_module.__default.fm0(_121_v2, _120_v1, _121_v2, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_121_v2,_128_v6)).length), _127_globalState), new BigNumber((_120_v1).length)), new _dafny.CodePoint('q'.codePointAt(0)))).length);
      _nw17[(new BigNumber(5)).toNumber()] = _119_v0;
      _nw17[(new BigNumber(6)).toNumber()] = _119_v0;
      _nw17[(new BigNumber(7)).toNumber()] = (((_131_v9).contains(function () {
        let _coll7 = new _dafny.Set();
        for (const _compr_7 of _dafny.IntegerRange(new BigNumber(218), new BigNumber(451))) {
          let _135_v10 = _compr_7;
          if (((new BigNumber(218)).isLessThanOrEqualTo(_135_v10)) && ((_135_v10).isLessThan(new BigNumber(451)))) {
            _coll7.add(_module.__default.safeDivisionInt(_135_v10, (_dafny.ZERO).minus(new BigNumber((_130_v8).length))));
          }
        }
        return _coll7;
      }())) ? ((_131_v9).get(function () {
        let _coll6 = new _dafny.Set();
        for (const _compr_6 of _dafny.IntegerRange(new BigNumber(218), new BigNumber(451))) {
          let _134_v10 = _compr_6;
          if (((new BigNumber(218)).isLessThanOrEqualTo(_134_v10)) && ((_134_v10).isLessThan(new BigNumber(451)))) {
            _coll6.add(_module.__default.safeDivisionInt(_134_v10, (_dafny.ZERO).minus(new BigNumber((_130_v8).length))));
          }
        }
        return _coll6;
      }())) : (new BigNumber(157)));
      _nw17[(new BigNumber(8)).toNumber()] = _119_v0;
      _nw17[(new BigNumber(9)).toNumber()] = (_119_v0).minus(_119_v0);
      _nw17[(new BigNumber(10)).toNumber()] = _119_v0;
      _nw17[(new BigNumber(11)).toNumber()] = _module.__default.safeDivisionInt(_119_v0, _119_v0);
      _nw17[(new BigNumber(12)).toNumber()] = _module.__default.safeModuloInt((_dafny.ZERO).minus(_119_v0), _119_v0);
      _nw17[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus(((_121_v2) ? (new BigNumber((_132_v11).length)) : ((_dafny.ZERO).minus(_119_v0))));
      _133_v12 = _nw17;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_133_v12).length))) {
        let _136_i1 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_136_i1)) && ((_136_i1).isLessThan(new BigNumber((_133_v12).length))))) {
          (_133_v12)[(_136_i1)] = _module.__default.safeDivisionInt(_136_i1, _119_v0);
        }
      }
      let _137_v13;
      _137_v13 = _dafny.MultiSet.fromElements(_119_v0, _119_v0);
      let _rhs9 = _128_v6;
      let _rhs10 = _119_v0;
      let _rhs11 = (_137_v13).IsSubsetOf(_137_v13);
      let _rhs12 = _133_v12;
      let _lhs2 = _127_globalState;
      let _lhs3 = _127_globalState;
      _128_v6 = _rhs9;
      _lhs2.f5 = _rhs10;
      _lhs3.f6 = _rhs11;
      _133_v12 = _rhs12;
      let _138_v14;
      let _139_v15;
      let _140_v16;
      let _out0;
      let _out1;
      let _out2;
      let _outcollector0 = _module.__default.m0(_119_v0, false, _121_v2, _120_v1, _127_globalState);
      _out0 = _outcollector0[0];
      _out1 = _outcollector0[1];
      _out2 = _outcollector0[2];
      _138_v14 = _out0;
      _139_v15 = _out1;
      _140_v16 = _out2;
      (_127_globalState).f5 = _119_v0;
      let _index8 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length));
      (_124_v5)[_index8] = !(_139_v15);
      let _index9 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length));
      (_124_v5)[_index9] = _139_v15;
      let _141_v17;
      let _init4 = ((_142_v1) => function (_143_i2) {
        return _142_v1;
      })(_120_v1);
      let _nw18 = Array((new BigNumber(10)).toNumber());
      for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw18.length); _i0_4++) {
        _nw18[_i0_4] = _init4(new BigNumber(_i0_4));
      }
      _141_v17 = _nw18;
      let _index10 = _module.__default.safeIndex(new BigNumber(564), new BigNumber((_141_v17).length));
      (_141_v17)[_index10] = _120_v1;
      let _index11 = _module.__default.safeIndex(new BigNumber(564), new BigNumber((_141_v17).length));
      (_141_v17)[_index11] = (_module.D0.create_DC0(_120_v1)).dtor_cf0;
      let _144_i3;
      _144_i3 = _dafny.ZERO;
      L0: {
        while (!(_139_v15)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_144_i3)) {
              break L0;
            }
            _144_i3 = (_144_i3).plus(_dafny.ONE);
            let _145_v18;
            _145_v18 = _module.D1.create_DC2(_128_v6);
            let _146_v19;
            let _nw19 = new _module.C0();
            _nw19.__ctor(_119_v0, (_145_v18).dtor_cf6);
            _146_v19 = _nw19;
            let _index12 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_133_v12).length));
            (_133_v12)[_index12] = (new BigNumber((_dafny.Seq.UnicodeFromString("ofwu")).length)).plus(new BigNumber((_120_v1).length));
            let _147_v20;
            _147_v20 = _dafny.Map.Empty.slice().updateUnsafe(_121_v2,new BigNumber((_dafny.Seq.UnicodeFromString("oijvun")).length));
            let _index13 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_133_v12).length));
            (_133_v12)[_index13] = (_dafny.ZERO).minus(new BigNumber((_module.__default.fm3(((false) ? ((((_147_v20).contains(_121_v2)) ? ((_147_v20).get(_121_v2)) : (new BigNumber(-524)))) : ((_146_v19).f19)), ((_124_v5)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length))]) || ((_124_v5)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length))]), _122_v3, _147_v20, _127_globalState)).length));
            let _148_v21;
            _148_v21 = _dafny.Map.Empty.slice().updateUnsafe((_124_v5)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length))],!(_139_v15));
            let _source3 = _module.D1.create_DC3(_139_v15, !_dafny.Seq.contains(_122_v3, true), _139_v15, _dafny.Seq.IsPrefixOf(_dafny.Seq.update(_129_v7, _module.__default.safeIndex(_module.__default.fm0(_139_v15, (_141_v17)[_module.__default.safeIndex(new BigNumber(564), new BigNumber((_141_v17).length))], _139_v15, _119_v0, _127_globalState), new BigNumber((_129_v7).length)), (_146_v19).f19), _dafny.Seq.of(new BigNumber((_148_v21).length))));
            if (_source3.is_DC3) {
              let _149___mcc_h0 = (_source3).cf7;
              let _150___mcc_h1 = (_source3).cf8;
              let _151___mcc_h2 = (_source3).cf9;
              let _152___mcc_h3 = (_source3).cf10;
              let _153_cf10 = _152___mcc_h3;
              let _154_cf9 = _151___mcc_h2;
              let _155_cf8 = _150___mcc_h1;
              let _156_cf7 = _149___mcc_h0;
              (_127_globalState).f5 = (_146_v19).f19;
              let _157_v22;
              let _158_v23;
              let _159_v24;
              let _out3;
              let _out4;
              let _out5;
              let _outcollector1 = _module.__default.m0(_module.__default.safeDivisionInt((_133_v12)[_module.__default.safeIndex(new BigNumber(686), new BigNumber((_133_v12).length))], (_133_v12)[_module.__default.safeIndex(new BigNumber(686), new BigNumber((_133_v12).length))]), (_124_v5)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length))], (_124_v5)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length))], _dafny.Seq.Concat((_141_v17)[_module.__default.safeIndex(new BigNumber(564), new BigNumber((_141_v17).length))], _120_v1), _127_globalState);
              _out3 = _outcollector1[0];
              _out4 = _outcollector1[1];
              _out5 = _outcollector1[2];
              _157_v22 = _out3;
              _158_v23 = _out4;
              _159_v24 = _out5;
              let _160_v25;
              let _nw20 = new _module.C0();
              _nw20.__ctor(_119_v0, _128_v6);
              _160_v25 = _nw20;
              let _161_v26;
              _161_v26 = _dafny.MultiSet.fromElements(_146_v19);
              let _162_v27;
              let _nw21 = Array((new BigNumber(27)).toNumber());
              _nw21[(_dafny.ZERO).toNumber()] = (_160_v25).f19;
              _nw21[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Set.fromElements(_158_v23)).length);
              _nw21[(new BigNumber(2)).toNumber()] = (_146_v19).f19;
              _nw21[(new BigNumber(3)).toNumber()] = (_133_v12)[_module.__default.safeIndex(new BigNumber(686), new BigNumber((_133_v12).length))];
              _nw21[(new BigNumber(4)).toNumber()] = (_146_v19).f19;
              _nw21[(new BigNumber(5)).toNumber()] = (_146_v19).f19;
              _nw21[(new BigNumber(6)).toNumber()] = (_146_v19).f19;
              _nw21[(new BigNumber(7)).toNumber()] = (_133_v12)[_module.__default.safeIndex(new BigNumber(686), new BigNumber((_133_v12).length))];
              _nw21[(new BigNumber(8)).toNumber()] = new BigNumber(179);
              _nw21[(new BigNumber(9)).toNumber()] = _119_v0;
              _nw21[(new BigNumber(10)).toNumber()] = _119_v0;
              _nw21[(new BigNumber(11)).toNumber()] = (_146_v19).f19;
              _nw21[(new BigNumber(12)).toNumber()] = (_146_v19).f19;
              _nw21[(new BigNumber(13)).toNumber()] = new BigNumber((_148_v21).length);
              _nw21[(new BigNumber(14)).toNumber()] = _119_v0;
              _nw21[(new BigNumber(15)).toNumber()] = (_160_v25).f19;
              _nw21[(new BigNumber(16)).toNumber()] = new BigNumber(((_141_v17)[_module.__default.safeIndex(new BigNumber(564), new BigNumber((_141_v17).length))]).length);
              _nw21[(new BigNumber(17)).toNumber()] = (_146_v19).f19;
              _nw21[(new BigNumber(18)).toNumber()] = (_146_v19).f19;
              _nw21[(new BigNumber(19)).toNumber()] = _119_v0;
              _nw21[(new BigNumber(20)).toNumber()] = new BigNumber(-124);
              _nw21[(new BigNumber(21)).toNumber()] = (_146_v19).f19;
              _nw21[(new BigNumber(22)).toNumber()] = (_146_v19).f19;
              _nw21[(new BigNumber(23)).toNumber()] = (_133_v12)[_module.__default.safeIndex(new BigNumber(686), new BigNumber((_133_v12).length))];
              _nw21[(new BigNumber(24)).toNumber()] = (((_137_v13).contains(new BigNumber((_161_v26).cardinality()))) ? ((_137_v13).get(new BigNumber((_161_v26).cardinality()))) : ((_146_v19).f19));
              _nw21[(new BigNumber(25)).toNumber()] = (_146_v19).f19;
              _nw21[(new BigNumber(26)).toNumber()] = new BigNumber(-897);
              _162_v27 = _nw21;
              let _163_v28;
              _163_v28 = _dafny.Map.Empty.slice().updateUnsafe((_146_v19).f20,_162_v27);
              let _164_v29;
              let _165_v30;
              let _166_v31;
              let _out6;
              let _out7;
              let _out8;
              let _outcollector2 = _module.__default.m0(_119_v0, !(_163_v28).equals((_163_v28)), _139_v15, (_141_v17)[_module.__default.safeIndex(new BigNumber(564), new BigNumber((_141_v17).length))], _127_globalState);
              _out6 = _outcollector2[0];
              _out7 = _outcollector2[1];
              _out8 = _outcollector2[2];
              _164_v29 = _out6;
              _165_v30 = _out7;
              _166_v31 = _out8;
            } else if (_source3.is_DC2) {
              let _167___mcc_h4 = (_source3).cf6;
              let _168_cf6 = _167___mcc_h4;
              let _169_v32;
              _169_v32 = _dafny.Seq.of(_129_v7);
              let _170_v33;
              _170_v33 = _dafny.Map.Empty.slice().updateUnsafe((_169_v32)[_module.__default.safeIndex((_133_v12)[_module.__default.safeIndex(new BigNumber(686), new BigNumber((_133_v12).length))], new BigNumber((_169_v32).length))],(_141_v17)[_module.__default.safeIndex(new BigNumber(564), new BigNumber((_141_v17).length))]);
              let _171_v34;
              _171_v34 = _dafny.Set.fromElements(!((_170_v33).equals(_170_v33)));
              let _172_v35;
              _172_v35 = _dafny.Map.Empty.slice().updateUnsafe(_146_v19,_119_v0);
              let _173_v36;
              _173_v36 = _dafny.MultiSet.fromElements((_172_v35).contains(_146_v19), false);
              let _174_v37;
              _174_v37 = _dafny.Seq.of(_146_v19);
              let _175_v38;
              _175_v38 = _dafny.Seq.of(_171_v34);
              let _index14 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_133_v12).length));
              let _index15 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_133_v12).length));
              let _rhs13 = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.update(_174_v37, _module.__default.safeIndex((_133_v12)[_module.__default.safeIndex(new BigNumber(686), new BigNumber((_133_v12).length))], new BigNumber((_174_v37).length)), _146_v19)).length), _119_v0);
              let _rhs14 = (_175_v38)[_module.__default.safeIndex(new BigNumber((_122_v3).length), new BigNumber((_175_v38).length))];
              let _rhs15 = _module.__default.fm4(_139_v15, _127_globalState);
              let _rhs16 = (_146_v19).f19;
              let _lhs4 = _133_v12;
              let _lhs5 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_133_v12).length));
              let _lhs6 = _133_v12;
              let _lhs7 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_133_v12).length));
              _lhs4[_lhs5] = _rhs13;
              _171_v34 = _rhs14;
              _173_v36 = _rhs15;
              _lhs6[_lhs7] = _rhs16;
              _119_v0 = _119_v0;
              _130_v8 = function () {
                let _coll8 = new _dafny.Set();
                for (const _compr_8 of (_129_v7).Elements) {
                  let _176_v39 = _compr_8;
                  if (_dafny.Seq.contains(_129_v7, _176_v39)) {
                    _coll8.add(_module.__default.safeDivisionInt(_176_v39, new BigNumber(-557)));
                  }
                }
                return _coll8;
              }();
              let _177_v40;
              let _nw22 = Array((new BigNumber(4)).toNumber());
              _177_v40 = _nw22;
              _177_v40 = _177_v40;
            } else {
              let _178___mcc_h5 = (_source3).cf11;
              let _179_cf11 = _178___mcc_h5;
              (_127_globalState).f15 = (_133_v12)[_module.__default.safeIndex(new BigNumber(686), new BigNumber((_133_v12).length))];
              let _180_v41;
              let _init5 = ((_181_v0) => function (_182_i4) {
                return (_182_i4).multipliedBy(_181_v0);
              })(_119_v0);
              let _nw23 = Array((new BigNumber(23)).toNumber());
              for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw23.length); _i0_5++) {
                _nw23[_i0_5] = _init5(new BigNumber(_i0_5));
              }
              _180_v41 = _nw23;
              let _183_v42;
              _183_v42 = _dafny.Map.Empty.slice().updateUnsafe((_146_v19).f20,_180_v41);
              let _184_v43;
              _184_v43 = _183_v42;
              _184_v43 = _183_v42;
              _128_v6 = (_146_v19).f20;
              _184_v43 = _184_v43;
            }
            let _185_v44;
            let _186_v45;
            let _187_v46;
            let _out9;
            let _out10;
            let _out11;
            let _outcollector3 = _module.__default.m0((_146_v19).f19, _121_v2, _139_v15, _dafny.Seq.Concat((_141_v17)[_module.__default.safeIndex(new BigNumber(564), new BigNumber((_141_v17).length))], _120_v1), _127_globalState);
            _out9 = _outcollector3[0];
            _out10 = _outcollector3[1];
            _out11 = _outcollector3[2];
            _185_v44 = _out9;
            _186_v45 = _out10;
            _187_v46 = _out11;
          }
        }
      }
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_133_v12).length))) {
        let _188_i5 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_188_i5)) && ((_188_i5).isLessThan(new BigNumber((_133_v12).length))))) {
          (_133_v12)[(_188_i5)] = (_188_i5).minus(((_139_v15) ? (_119_v0) : (_119_v0)));
        }
      }
      let _189_v47;
      _189_v47 = _module.D1.create_DC3((_124_v5)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length))], _121_v2, _121_v2, _139_v15);
      (_127_globalState).f6 = !((_189_v47).dtor_cf7);
      let _hi1 = new BigNumber((_dafny.Set.fromElements(_119_v0, _119_v0)).length);
      for (let _190_i6 = _119_v0; _190_i6.isLessThan(_hi1); _190_i6 = _190_i6.plus(_dafny.ONE)) {
        let _191_v48;
        _191_v48 = _module.D0.create_DC1(_133_v12, _124_v5, _123_v4, _190_i6, _190_i6);
        let _source4 = _191_v48;
        if (_source4.is_DC1) {
          let _192___mcc_h6 = (_source4).cf1;
          let _193___mcc_h7 = (_source4).cf2;
          let _194___mcc_h8 = (_source4).cf3;
          let _195___mcc_h9 = (_source4).cf4;
          let _196___mcc_h10 = (_source4).cf5;
          let _197_cf5 = _196___mcc_h10;
          let _198_cf4 = _195___mcc_h9;
          let _199_cf3 = _194___mcc_h8;
          let _200_cf2 = _193___mcc_h7;
          let _201_cf1 = _192___mcc_h6;
          _128_v6 = (_120_v1)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(603)), ((_202_v6) => function (_203_i7) {
            return _202_v6;
          })(_128_v6))).length), new BigNumber((_120_v1).length))];
          let _204_v49;
          let _205_v50;
          let _206_v51;
          let _out12;
          let _out13;
          let _out14;
          let _outcollector4 = _module.__default.m0(_198_cf4, true, _139_v15, _120_v1, _127_globalState);
          _out12 = _outcollector4[0];
          _out13 = _outcollector4[1];
          _out14 = _outcollector4[2];
          _204_v49 = _out12;
          _205_v50 = _out13;
          _206_v51 = _out14;
          (_127_globalState).f5 = _197_cf5;
          _197_cf5 = _198_cf4;
        } else {
          let _207___mcc_h11 = (_source4).cf0;
          let _208_cf0 = _207___mcc_h11;
          _122_v3 = _122_v3;
          let _209_v52;
          let _init6 = ((_210_v1) => function (_211_i8) {
            return _module.D0.create_DC0(_210_v1);
          })(_120_v1);
          let _nw24 = Array((new BigNumber(11)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw24.length); _i0_6++) {
            _nw24[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _209_v52 = _nw24;
          let _212_v53;
          _212_v53 = _module.D0.create_DC0(_120_v1);
          let _index16 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_209_v52).length));
          (_209_v52)[_index16] = _212_v53;
          let _index17 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_209_v52).length));
          let _rhs17 = (_141_v17)[_module.__default.safeIndex(new BigNumber(564), new BigNumber((_141_v17).length))];
          let _rhs18 = _119_v0;
          let _rhs19 = (_124_v5)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length))];
          let _rhs20 = !(false);
          let _rhs21 = _module.D0.create_DC0(_dafny.Seq.Concat(_208_cf0, _dafny.Seq.UnicodeFromString("fy")));
          let _lhs8 = _127_globalState;
          let _lhs9 = _127_globalState;
          let _lhs10 = _127_globalState;
          let _lhs11 = _209_v52;
          let _lhs12 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_209_v52).length));
          _208_cf0 = _rhs17;
          _lhs8.f13 = _rhs18;
          _lhs9.f6 = _rhs19;
          _lhs10.f6 = _rhs20;
          _lhs11[_lhs12] = _rhs21;
          let _213_v54;
          _213_v54 = _dafny.Map.Empty.slice().updateUnsafe(_128_v6,_133_v12);
          let _214_v55;
          _214_v55 = _213_v54;
          let _215_v56;
          _215_v56 = _dafny.Map.Empty.slice().updateUnsafe(_119_v0,_214_v55);
          _121_v2 = (new BigNumber(((_215_v56).update(_190_i6, _214_v55)).length)).isLessThanOrEqualTo(_190_i6);
          let _216_v57;
          let _217_v58;
          let _218_v59;
          let _out15;
          let _out16;
          let _out17;
          let _outcollector5 = _module.__default.m0(new BigNumber(849), ((_124_v5)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length))]) === ((_124_v5)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length))]), (_124_v5)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length))], _120_v1, _127_globalState);
          _out15 = _outcollector5[0];
          _out16 = _outcollector5[1];
          _out17 = _outcollector5[2];
          _216_v57 = _out15;
          _217_v58 = _out16;
          _218_v59 = _out17;
        }
        let _219_v60;
        let _init7 = ((_220_v7, _221_i6) => function (_222_i9) {
          return _dafny.Seq.update(_220_v7, _module.__default.safeIndex(new BigNumber(180), new BigNumber((_220_v7).length)), _221_i6);
        })(_129_v7, _190_i6);
        let _nw25 = Array((new BigNumber(5)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw25.length); _i0_7++) {
          _nw25[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _219_v60 = _nw25;
        _219_v60 = _219_v60;
        (_127_globalState).f6 = !(!(!(_139_v15)));
        let _223_v61;
        let _nw26 = new _module.C0();
        _nw26.__ctor(new BigNumber(204), _128_v6);
        _223_v61 = _nw26;
        let _224_v62;
        let _init8 = ((_225_v5) => function (_226_i10) {
          return _module.D1.create_DC3((_225_v5)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_225_v5).length))], (_225_v5)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_225_v5).length))], true, (_225_v5)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_225_v5).length))]);
        })(_124_v5);
        let _nw27 = Array((new BigNumber(29)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw27.length); _i0_8++) {
          _nw27[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _224_v62 = _nw27;
        let _227_v63;
        _227_v63 = _dafny.Map.Empty.slice().updateUnsafe(_223_v61,_224_v62);
        let _index18 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length));
        let _rhs22 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt(_119_v0, new BigNumber((_227_v63).length)), _190_i6);
        let _rhs23 = _121_v2;
        let _rhs24 = _119_v0;
        let _rhs25 = _141_v17;
        let _rhs26 = new BigNumber(827);
        let _lhs13 = _124_v5;
        let _lhs14 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length));
        let _lhs15 = _127_globalState;
        _119_v0 = _rhs22;
        _lhs13[_lhs14] = _rhs23;
        _119_v0 = _rhs24;
        _141_v17 = _rhs25;
        _lhs15.f15 = _rhs26;
      }
      let _hi2 = _119_v0;
      for (let _228_i11 = _module.__default.safeDivisionInt(_119_v0, _119_v0); _228_i11.isLessThan(_hi2); _228_i11 = _228_i11.plus(_dafny.ONE)) {
        if (_121_v2) {
          let _229_v64;
          let _nw28 = new _module.C0();
          _nw28.__ctor(_119_v0, _128_v6);
          _229_v64 = _nw28;
          let _230_v65;
          _230_v65 = _dafny.Map.Empty.slice().updateUnsafe(_121_v2,_229_v64);
          let _231_v66;
          let _nw29 = Array((new BigNumber(3)).toNumber());
          _nw29[(_dafny.ZERO).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_139_v15,_229_v64)).Merge(_230_v65);
          _nw29[(_dafny.ONE).toNumber()] = _230_v65;
          _nw29[(new BigNumber(2)).toNumber()] = (_230_v65).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(_139_v15),_229_v64));
          _231_v66 = _nw29;
          _231_v66 = _231_v66;
          let _232_v67;
          let _init9 = ((_233_v3, _234_v15) => function (_235_i12) {
            return (_dafny.MultiSet.FromArray(_233_v3)).Difference(_dafny.MultiSet.fromElements(_234_v15));
          })(_122_v3, _139_v15);
          let _nw30 = Array((new BigNumber(3)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw30.length); _i0_9++) {
            _nw30[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _232_v67 = _nw30;
          let _236_v68;
          _236_v68 = _dafny.MultiSet.fromElements(_139_v15);
          let _index19 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_232_v67).length));
          (_232_v67)[_index19] = (_236_v68).Difference(_dafny.MultiSet.fromElements((_124_v5)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length))], (_124_v5)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length))], _139_v15, _121_v2));
          let _index20 = _module.__default.safeIndex(new BigNumber(456), new BigNumber((_232_v67).length));
          (_232_v67)[_index20] = _236_v68;
          let _index21 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_133_v12).length));
          (_133_v12)[_index21] = (_229_v64).f19;
          let _237_v69;
          _237_v69 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_236_v68).cardinality()),_139_v15);
          let _index22 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_133_v12).length));
          let _rhs27 = _module.__default.safeModuloInt(new BigNumber((_237_v69).length), _228_i11);
          let _rhs28 = _121_v2;
          let _lhs16 = _133_v12;
          let _lhs17 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_133_v12).length));
          let _lhs18 = _127_globalState;
          _lhs16[_lhs17] = _rhs27;
          _lhs18.f6 = _rhs28;
          let _238_v70;
          _238_v70 = _dafny.Seq.of(_124_v5);
          _238_v70 = _dafny.Seq.update(_238_v70, _module.__default.safeIndex((_133_v12)[_module.__default.safeIndex(new BigNumber(807), new BigNumber((_133_v12).length))], new BigNumber((_238_v70).length)), _124_v5);
          (_127_globalState).f5 = _module.__default.safeDivisionInt(_228_i11, _228_i11);
        } else {
          (_127_globalState).f6 = (_124_v5)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length))];
          let _239_v71;
          _239_v71 = _module.D0.create_DC1(_133_v12, _124_v5, _123_v4, _119_v0, (_129_v7)[_module.__default.safeIndex(new BigNumber((_120_v1).length), new BigNumber((_129_v7).length))]);
          let _240_v72;
          _240_v72 = _dafny.Seq.of(_239_v71);
          let _index23 = _module.__default.safeIndex(new BigNumber(564), new BigNumber((_141_v17).length));
          let _rhs29 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("vw"), _module.__default.safeIndex(_228_i11, new BigNumber((_dafny.Seq.UnicodeFromString("vw")).length)), _128_v6), _dafny.Seq.UnicodeFromString("pmpe"));
          let _rhs30 = !(_139_v15) || (_139_v15);
          let _rhs31 = _dafny.Seq.Concat(_dafny.Seq.Concat(_240_v72, _240_v72), _240_v72);
          let _rhs32 = _module.__default.fm5(_127_globalState);
          let _lhs19 = _141_v17;
          let _lhs20 = _module.__default.safeIndex(new BigNumber(564), new BigNumber((_141_v17).length));
          let _lhs21 = _127_globalState;
          _lhs19[_lhs20] = _rhs29;
          _lhs21.f6 = _rhs30;
          _240_v72 = _rhs31;
          _128_v6 = _rhs32;
          (_127_globalState).f15 = _228_i11;
          (_127_globalState).f6 = _121_v2;
          let _241_v73;
          _241_v73 = _dafny.Map.Empty.slice().updateUnsafe(_239_v71,(new BigNumber((_137_v13).cardinality())).plus(_228_i11));
          _241_v73 = (_241_v73).update(_239_v71, (_dafny.ZERO).minus(_119_v0));
        }
        let _242_v74;
        _242_v74 = _dafny.Seq.of(_120_v1);
        _242_v74 = _242_v74;
        let _243_v75;
        _243_v75 = _dafny.Set.fromElements(true);
        let _244_v76;
        _244_v76 = _dafny.Set.fromElements(_243_v75);
        let _245_v77;
        _245_v77 = _dafny.Map.Empty.slice().updateUnsafe(_128_v6,new BigNumber((_244_v76).length));
        _245_v77 = (_245_v77).update(_128_v6, (_dafny.ZERO).minus(_119_v0));
        (_127_globalState).f15 = (_119_v0).multipliedBy(_228_i11);
      }
      let _246_i13;
      _246_i13 = _dafny.ZERO;
      L1: {
        while (!(_139_v15)) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_246_i13)) {
              break L1;
            }
            _246_i13 = (_246_i13).plus(_dafny.ONE);
            let _247_v78;
            let _248_v79;
            let _249_v80;
            let _out18;
            let _out19;
            let _out20;
            let _outcollector6 = _module.__default.m0((((_132_v11).contains(new BigNumber(((_141_v17)[_module.__default.safeIndex(new BigNumber(564), new BigNumber((_141_v17).length))]).length))) ? ((_132_v11).get(new BigNumber(((_141_v17)[_module.__default.safeIndex(new BigNumber(564), new BigNumber((_141_v17).length))]).length))) : (_119_v0)), !(_139_v15), !(_119_v0).isEqualTo(_119_v0), _120_v1, _127_globalState);
            _out18 = _outcollector6[0];
            _out19 = _outcollector6[1];
            _out20 = _outcollector6[2];
            _247_v78 = _out18;
            _248_v79 = _out19;
            _249_v80 = _out20;
            let _250_v81;
            let _nw31 = new _module.C0();
            _nw31.__ctor(_119_v0, _128_v6);
            _250_v81 = _nw31;
            let _251_v82;
            _251_v82 = _dafny.Seq.of(_130_v8, _module.__default.fm1((_124_v5)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length))], new BigNumber((_dafny.Seq.UnicodeFromString("vityokux")).length), new BigNumber((_129_v7).length), _119_v0, _127_globalState));
            let _252_v83;
            _252_v83 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsPrefixOf(_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(719), (_250_v81).f19), _130_v8), _251_v82),!((_119_v0).isLessThan(_119_v0)));
            _252_v83 = (_252_v83).update(_248_v79, (_124_v5)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length))]);
            let _index24 = _module.__default.safeIndex(new BigNumber(550), new BigNumber((_133_v12).length));
            (_133_v12)[_index24] = _119_v0;
            let _index25 = _module.__default.safeIndex(new BigNumber(550), new BigNumber((_133_v12).length));
            (_133_v12)[_index25] = _119_v0;
          }
        }
      }
      let _253_v84;
      _253_v84 = _dafny.Set.fromElements(_139_v15, (_124_v5)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length))]);
      let _hi3 = _module.__default.safeModuloInt(_119_v0, _119_v0);
      for (let _254_i14 = new BigNumber(((_253_v84).Union(_253_v84)).length); _254_i14.isLessThan(_hi3); _254_i14 = _254_i14.plus(_dafny.ONE)) {
        if (_139_v15) {
          _189_v47 = _189_v47;
          (_127_globalState).f15 = (_dafny.ZERO).minus((_254_i14).minus(new BigNumber(881)));
          let _index26 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length));
          (_124_v5)[_index26] = !(!(_139_v15));
          (_127_globalState).f13 = _254_i14;
          let _255_v85;
          _255_v85 = _dafny.Seq.of(_124_v5);
          let _256_v86;
          _256_v86 = _module.D3.create_DC6(_255_v85);
          (_127_globalState).f5 = new BigNumber(((_256_v86).dtor_cf13).length);
        } else {
          (_127_globalState).f6 = _121_v2;
          _189_v47 = _189_v47;
          (_127_globalState).f6 = (new BigNumber(-512)).isLessThan(_119_v0);
          let _index27 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length));
          (_124_v5)[_index27] = (_253_v84).IsDisjointFrom(_253_v84);
          let _257_v87;
          let _nw32 = Array((new BigNumber(14)).toNumber());
          _257_v87 = _nw32;
          let _258_v88;
          _258_v88 = _dafny.Map.Empty.slice().updateUnsafe(_119_v0,_121_v2);
          let _259_v89;
          _259_v89 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm6((_dafny.ZERO).minus(new BigNumber((_258_v88).length)), _128_v6, (_dafny.ZERO).minus(_119_v0), _127_globalState),_257_v87);
          _257_v87 = (((_259_v89).contains((_dafny.Set.fromElements(_121_v2)).IsProperSubsetOf(_253_v84))) ? ((_259_v89).get((_dafny.Set.fromElements(_121_v2)).IsProperSubsetOf(_253_v84))) : (_257_v87));
        }
        (_127_globalState).f15 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt(_119_v0, _254_i14), new BigNumber((_122_v3).length));
        let _260_v90;
        _260_v90 = _dafny.Map.Empty.slice().updateUnsafe(true,(_module.D0.create_DC1(_133_v12, _124_v5, _123_v4, _119_v0, _254_i14)).dtor_cf2);
        let _261_v91;
        _261_v91 = _dafny.Map.Empty.slice().updateUnsafe(_128_v6,(_141_v17)[_module.__default.safeIndex(new BigNumber(564), new BigNumber((_141_v17).length))]);
        let _262_v92;
        _262_v92 = _dafny.Map.Empty.slice().updateUnsafe(_254_i14,_139_v15);
        let _263_v93;
        _263_v93 = _module.D0.create_DC1(_133_v12, (((_260_v90).contains((_124_v5)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length))])) ? ((_260_v90).get((_124_v5)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length))])) : (_124_v5)), _dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(592)), ((_264_v6) => function (_265_i15) {
  return _264_v6;
})(_128_v6)), (((_261_v91).contains(_128_v6)) ? ((_261_v91).get(_128_v6)) : (_module.__default.fm7((_124_v5)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length))], _254_i14, _module.__default.fm0((((_262_v92).contains(_module.__default.fm0((_124_v5)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length))], _120_v1, (_124_v5)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length))], _254_i14, _127_globalState))) ? ((_262_v92).get(_module.__default.fm0((_124_v5)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length))], _120_v1, (_124_v5)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length))], _254_i14, _127_globalState))) : (!(_121_v2))), _dafny.Seq.UnicodeFromString("ytn"), _139_v15, _119_v0, _127_globalState), _127_globalState))), (_141_v17)[_module.__default.safeIndex(new BigNumber(564), new BigNumber((_141_v17).length))], (_141_v17)[_module.__default.safeIndex(new BigNumber(564), new BigNumber((_141_v17).length))]), _254_i14, _119_v0);
        _263_v93 = _263_v93;
        let _266_v94;
        let _267_v95;
        let _268_v96;
        let _out21;
        let _out22;
        let _out23;
        let _outcollector7 = _module.__default.m0((_dafny.ZERO).minus(_module.__default.safeDivisionInt(_254_i14, _119_v0)), _139_v15, !(_130_v8).contains(new BigNumber(232)), _dafny.Seq.UnicodeFromString("ckhhe"), _127_globalState);
        _out21 = _outcollector7[0];
        _out22 = _outcollector7[1];
        _out23 = _outcollector7[2];
        _266_v94 = _out21;
        _267_v95 = _out22;
        _268_v96 = _out23;
      }
      let _269_v97;
      let _nw33 = Array((new BigNumber(4)).toNumber());
      _269_v97 = _nw33;
      let _270_v98;
      let _nw34 = new _module.C0();
      _nw34.__ctor(_119_v0, _128_v6);
      _270_v98 = _nw34;
      let _index28 = _module.__default.safeIndex(new BigNumber(411), new BigNumber((_269_v97).length));
      (_269_v97)[_index28] = _270_v98;
      let _index29 = _module.__default.safeIndex(new BigNumber(411), new BigNumber((_269_v97).length));
      (_269_v97)[_index29] = _270_v98;
      let _271_v99;
      _271_v99 = _dafny.Seq.of((_269_v97)[_module.__default.safeIndex(new BigNumber(411), new BigNumber((_269_v97).length))]);
      let _272_i16;
      _272_i16 = _dafny.ZERO;
      L2: {
        while (!_dafny.areEqual(_271_v99, _271_v99)) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_272_i16)) {
              break L2;
            }
            _272_i16 = (_272_i16).plus(_dafny.ONE);
            let _index30 = _module.__default.safeIndex(new BigNumber(564), new BigNumber((_141_v17).length));
            (_141_v17)[_index30] = _dafny.Seq.Concat((_141_v17)[_module.__default.safeIndex(new BigNumber(564), new BigNumber((_141_v17).length))], (_141_v17)[_module.__default.safeIndex(new BigNumber(564), new BigNumber((_141_v17).length))]);
            let _273_v100;
            let _274_v101;
            let _275_v102;
            let _out24;
            let _out25;
            let _out26;
            let _outcollector8 = _module.__default.m0((((_137_v13).contains(_119_v0)) ? ((_137_v13).get(_119_v0)) : ((_270_v98).f19)), !(_139_v15) || (!(true)), _121_v2, _dafny.Seq.update(_dafny.Seq.Concat(_120_v1, _dafny.Seq.update(_120_v1, _module.__default.safeIndex((_270_v98).f19, new BigNumber((_120_v1).length)), _128_v6)), _module.__default.safeIndex((_270_v98).f19, new BigNumber((_dafny.Seq.Concat(_120_v1, _dafny.Seq.update(_120_v1, _module.__default.safeIndex((_270_v98).f19, new BigNumber((_120_v1).length)), _128_v6))).length)), _128_v6), _127_globalState);
            _out24 = _outcollector8[0];
            _out25 = _outcollector8[1];
            _out26 = _outcollector8[2];
            _273_v100 = _out24;
            _274_v101 = _out25;
            _275_v102 = _out26;
            _274_v101 = _121_v2;
            if (_121_v2) {
              let _276_v103;
              let _nw35 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Map.Empty);
              _276_v103 = _nw35;
              let _277_v104;
              _277_v104 = _dafny.Seq.of(_124_v5);
              let _278_v105;
              _278_v105 = _module.D3.create_DC6(_277_v104);
              let _279_v106;
              _279_v106 = _dafny.Map.Empty.slice().updateUnsafe(_278_v105,(_270_v98).f19);
              let _280_v107;
              _280_v107 = _dafny.Map.Empty.slice().updateUnsafe(_279_v106,_139_v15);
              let _index31 = _module.__default.safeIndex(new BigNumber(196), new BigNumber((_276_v103).length));
              (_276_v103)[_index31] = _280_v107;
              let _281_v108;
              _281_v108 = _dafny.Map.Empty.slice().updateUnsafe((_124_v5)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length))],_280_v107);
              let _index32 = _module.__default.safeIndex(new BigNumber(196), new BigNumber((_276_v103).length));
              let _rhs33 = ((((_281_v108).contains(false)) ? ((_281_v108).get(false)) : (_280_v107))).Merge(_280_v107);
              let _rhs34 = (new BigNumber(138)).isEqualTo(new BigNumber(62));
              let _lhs22 = _276_v103;
              let _lhs23 = _module.__default.safeIndex(new BigNumber(196), new BigNumber((_276_v103).length));
              let _lhs24 = _127_globalState;
              _lhs22[_lhs23] = _rhs33;
              _lhs24.f6 = _rhs34;
              _141_v17 = _141_v17;
              (_127_globalState).f5 = _module.__default.safeModuloInt(((_270_v98).f19).plus((_270_v98).f19), (_270_v98).f19);
              let _282_v109;
              _282_v109 = _module.D1.create_DC2((_270_v98).f20);
              let _pat_let_tv0 = _128_v6;
              _282_v109 = function (_pat_let0_0) {
                return function (_283_dt__update__tmp_h1) {
                  return function (_pat_let1_0) {
                    return function (_284_dt__update_hcf6_h0) {
                      return _module.D1.create_DC2(_284_dt__update_hcf6_h0);
                    }(_pat_let1_0);
                  }(_pat_let_tv0);
                }(_pat_let0_0);
              }(_282_v109);
              (_127_globalState).f15 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt(((_dafny.ZERO).minus((_270_v98).f19)).plus(new BigNumber(259)), (_270_v98).fm2(_119_v0, (_270_v98).f19, (_270_v98).f19, _127_globalState))));
            } else {
              _132_v11 = (_132_v11).update(new BigNumber((_122_v3).length), (_270_v98).f19);
              _129_v7 = _dafny.Seq.Concat(_module.__default.fm8((_124_v5)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length))], _127_globalState), _129_v7);
              _120_v1 = _dafny.Seq.Concat(_120_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(368)), ((_285_v98) => function (_286_i17) {
                return (_285_v98).f20;
              })(_270_v98)));
              let _287_v110;
              let _init10 = ((_288_v2, _289_v101) => function (_290_i18) {
                return _dafny.Map.Empty.slice().updateUnsafe(_288_v2,_289_v101);
              })(_121_v2, _274_v101);
              let _nw36 = Array((new BigNumber(21)).toNumber());
              for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw36.length); _i0_10++) {
                _nw36[_i0_10] = _init10(new BigNumber(_i0_10));
              }
              _287_v110 = _nw36;
              let _index33 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length));
              let _rhs35 = ((_270_v98).f19).isLessThan(_119_v0);
              let _rhs36 = _287_v110;
              let _lhs25 = _124_v5;
              let _lhs26 = _module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length));
              _lhs25[_lhs26] = _rhs35;
              _287_v110 = _rhs36;
              let _291_v111;
              let _292_v112;
              let _293_v113;
              let _out27;
              let _out28;
              let _out29;
              let _outcollector9 = _module.__default.m0((_270_v98).f19, _274_v101, !((_124_v5)[_module.__default.safeIndex(new BigNumber(514), new BigNumber((_124_v5).length))]), _dafny.Seq.UnicodeFromString("svi"), _127_globalState);
              _out27 = _outcollector9[0];
              _out28 = _outcollector9[1];
              _out29 = _outcollector9[2];
              _291_v111 = _out27;
              _292_v112 = _out28;
              _293_v113 = _out29;
            }
          }
        }
      }
      let _index34 = _module.__default.safeIndex(new BigNumber(411), new BigNumber((_269_v97).length));
      (_269_v97)[_index34] = (_269_v97)[_module.__default.safeIndex(new BigNumber(411), new BigNumber((_269_v97).length))];
      process.stdout.write(_dafny.toString(_119_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_120_v1).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_121_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_122_v3, _dafny.Seq.of(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_123_v4).equals(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("qljvqw"), _dafny.Seq.UnicodeFromString("qljvqw"), _dafny.Seq.UnicodeFromString("qljvqw")))));
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
      process.stdout.write(_dafny.toString((_124_v5)[new BigNumber(27)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_124_v5)[new BigNumber(28)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_127_globalState.f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_127_globalState.f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_127_globalState).f7, _dafny.Seq.of(new BigNumber(327), new BigNumber(6), new BigNumber(327), new BigNumber(-327)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_globalState).f8).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_globalState).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_127_globalState).f12).equals(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("qljvqw"), _dafny.Seq.UnicodeFromString("qljvqw"), _dafny.Seq.UnicodeFromString("qljvqw")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_127_globalState.f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_127_globalState).f14).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_127_globalState.f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_globalState).f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_globalState).f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_globalState).f18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_128_v6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_129_v7, _dafny.Seq.of(new BigNumber(-349)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v8).equals(_dafny.Set.fromElements(new BigNumber(327), new BigNumber(-349)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_v9).equals(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(_dafny.ONE, new BigNumber(176), new BigNumber(166), new BigNumber(-3)), _dafny.Set.fromElements(new BigNumber(327), new BigNumber(-349))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_132_v11).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(327),new BigNumber(-327)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_133_v12)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_133_v12)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_133_v12)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_133_v12)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_133_v12)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_133_v12)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_133_v12)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_133_v12)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_133_v12)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_133_v12)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_133_v12)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_133_v12)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_133_v12)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_133_v12)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_137_v13).equals(_dafny.MultiSet.fromElements(new BigNumber(327), new BigNumber(327)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v14).equals(_dafny.Set.fromElements(_dafny.Set.fromElements(new BigNumber(327)), _dafny.Set.fromElements(new BigNumber(327), new BigNumber(-29), new BigNumber(-28), new BigNumber(-27), new BigNumber(-26), new BigNumber(-25), new BigNumber(-24), new BigNumber(-23), new BigNumber(-22), new BigNumber(-21), new BigNumber(-20), new BigNumber(-19), new BigNumber(-18), new BigNumber(-17), new BigNumber(-16), new BigNumber(-15), new BigNumber(-14), new BigNumber(-13), new BigNumber(-12), new BigNumber(-11), new BigNumber(-10), new BigNumber(-9), new BigNumber(-8), new BigNumber(-7), new BigNumber(-6), new BigNumber(-5), new BigNumber(-4), new BigNumber(-3), new BigNumber(-2), new BigNumber(-1), _dafny.ZERO, _dafny.ONE, new BigNumber(2), new BigNumber(3), new BigNumber(4), new BigNumber(5), new BigNumber(6), new BigNumber(7), new BigNumber(8), new BigNumber(9), new BigNumber(10), new BigNumber(11), new BigNumber(12), new BigNumber(13), new BigNumber(14), new BigNumber(15), new BigNumber(16), new BigNumber(17), new BigNumber(18), new BigNumber(19), new BigNumber(20), new BigNumber(21), new BigNumber(22), new BigNumber(23), new BigNumber(24), new BigNumber(25), new BigNumber(26), new BigNumber(27), new BigNumber(28), new BigNumber(29), new BigNumber(30), new BigNumber(31), new BigNumber(32), new BigNumber(33), new BigNumber(34), new BigNumber(35), new BigNumber(36), new BigNumber(37), new BigNumber(38), new BigNumber(39), new BigNumber(40), new BigNumber(41), new BigNumber(42), new BigNumber(43), new BigNumber(44), new BigNumber(45), new BigNumber(46), new BigNumber(47), new BigNumber(48), new BigNumber(49), new BigNumber(50), new BigNumber(51), new BigNumber(52), new BigNumber(53), new BigNumber(54), new BigNumber(55), new BigNumber(56), new BigNumber(57), new BigNumber(58), new BigNumber(59), new BigNumber(60), new BigNumber(61), new BigNumber(62), new BigNumber(63), new BigNumber(64), new BigNumber(65), new BigNumber(66), new BigNumber(67), new BigNumber(68), new BigNumber(69), new BigNumber(70), new BigNumber(71), new BigNumber(72), new BigNumber(73), new BigNumber(74), new BigNumber(75), new BigNumber(76), new BigNumber(77), new BigNumber(78), new BigNumber(79), new BigNumber(80), new BigNumber(81), new BigNumber(82), new BigNumber(83), new BigNumber(84), new BigNumber(85), new BigNumber(86), new BigNumber(87), new BigNumber(88), new BigNumber(89), new BigNumber(90), new BigNumber(91), new BigNumber(92)), _dafny.Set.fromElements(_dafny.ONE, new BigNumber(176), new BigNumber(166), new BigNumber(-3))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_139_v15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_140_v16).equals(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_141_v17)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_141_v17)[_dafny.ONE]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_141_v17)[new BigNumber(2)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_141_v17)[new BigNumber(3)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_141_v17)[new BigNumber(4)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_141_v17)[new BigNumber(5)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_141_v17)[new BigNumber(6)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_141_v17)[new BigNumber(7)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_141_v17)[new BigNumber(8)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_141_v17)[new BigNumber(9)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_144_i3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v47).dtor_cf7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v47).dtor_cf8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v47).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_189_v47).dtor_cf10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_246_i13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_253_v84).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_269_v97)[new BigNumber(3)]).f19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_269_v97)[new BigNumber(3)]).f20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_270_v98).f19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_270_v98).f20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_271_v99).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_272_i16));
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
    static create_DC0(cf0) {
      let $dt = new D0(1);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC0() { return this.$tag === 1; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC0" + "(" + this.cf0.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf1 === other.cf1 && this.cf2 === other.cf2 && _dafny.areEqual(this.cf3, other.cf3) && _dafny.areEqual(this.cf4, other.cf4) && _dafny.areEqual(this.cf5, other.cf5);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1([], [], _dafny.MultiSet.Empty, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC3(cf7, cf8, cf9, cf10) {
      let $dt = new D1(0);
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC2(cf6) {
      let $dt = new D1(1);
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC4(cf11) {
      let $dt = new D1(2);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC4() { return this.$tag === 2; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf7 === other.cf7 && this.cf8 === other.cf8 && this.cf9 === other.cf9 && this.cf10 === other.cf10;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf6, other.cf6);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf11, other.cf11);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC3(false, false, false, false);
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
    static create_DC5(cf12) {
      let $dt = new D2(0);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf12, other.cf12);
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
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC7() {
      let $dt = new D3(0);
      return $dt;
    }
    static create_DC8(cf14) {
      let $dt = new D3(1);
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC6(cf13) {
      let $dt = new D3(2);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC8() { return this.$tag === 1; }
    get is_DC6() { return this.$tag === 2; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC7";
      } else if (this.$tag === 1) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC6" + "(" + _dafny.toString(this.cf13) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf14, other.cf14);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf13, other.cf13);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC7();
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
    static create_DC10(cf16, cf17, cf18) {
      let $dt = new D4(0);
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC11(cf19, cf20, cf21, cf22, cf23) {
      let $dt = new D4(1);
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC12(cf24, cf25, cf26, cf27) {
      let $dt = new D4(2);
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC9(cf15) {
      let $dt = new D4(3);
      $dt.cf15 = cf15;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get is_DC12() { return this.$tag === 2; }
    get is_DC9() { return this.$tag === 3; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf15() { return this.cf15; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 3) {
        return "D4.DC9" + "(" + _dafny.toString(this.cf15) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf16, other.cf16) && this.cf17 === other.cf17 && this.cf18 === other.cf18;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf19, other.cf19) && this.cf20 === other.cf20 && _dafny.areEqual(this.cf21, other.cf21) && this.cf22 === other.cf22 && _dafny.areEqual(this.cf23, other.cf23);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf24, other.cf24) && this.cf25 === other.cf25 && this.cf26 === other.cf26 && _dafny.areEqual(this.cf27, other.cf27);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf15 === other.cf15;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC10(_dafny.ZERO, false, false);
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
      this.f5 = _dafny.ZERO;
      this.f6 = false;
      this.f13 = _dafny.ZERO;
      this.f15 = _dafny.ZERO;
      this._f0 = _dafny.ZERO;
      this._f1 = _dafny.ZERO;
      this._f2 = false;
      this._f3 = _dafny.ZERO;
      this._f4 = false;
      this._f7 = _dafny.Seq.of();
      this._f8 = _dafny.Map.Empty;
      this._f9 = _dafny.ZERO;
      this._f10 = _dafny.ZERO;
      this._f11 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f12 = _dafny.MultiSet.Empty;
      this._f14 = _dafny.Seq.of();
      this._f16 = false;
      this._f17 = _dafny.ZERO;
      this._f18 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this).f5 = f5;
      (_this).f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this).f13 = f13;
      (_this)._f14 = f14;
      (_this).f15 = f15;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
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
      this._f19 = _dafny.ZERO;
      this._f20 = new _dafny.CodePoint('D'.codePointAt(0));
    }
    _parentTraits() {
      return [];
    }
    __ctor(f19, f20) {
      let _this = this;
      (_this)._f19 = f19;
      (_this)._f20 = f20;
      return;
    }
    fm2(p0, p1, p2, globalState) {
      let _this = this;
      return (_this).f19;
    };
    get f19() {
      let _this = this;
      return _this._f19;
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
