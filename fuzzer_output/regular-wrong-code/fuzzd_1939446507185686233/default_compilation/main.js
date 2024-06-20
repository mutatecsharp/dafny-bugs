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
      return !(new BigNumber(864)).isEqualTo((new BigNumber((_dafny.Set.fromElements(new BigNumber(572))).length)).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-483)), function (_0_i0) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      })).length)));
    };
    static fm1(p0, p1, p2, p3, globalState) {
      return ((function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(587),false)).Keys.Elements) {
          let _1_v0 = _compr_0;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(587),false)).contains(_1_v0)) {
            _coll0.add((_1_v0).minus(new BigNumber((_dafny.MultiSet.fromElements(!(true))).cardinality())));
          }
        }
        return _coll0;
      }()).Union(function () {
        let _coll1 = new _dafny.Set();
        for (const _compr_1 of _dafny.IntegerRange(new BigNumber(182), new BigNumber(877))) {
          let _2_v1 = _compr_1;
          if (((new BigNumber(182)).isLessThanOrEqualTo(_2_v1)) && ((_2_v1).isLessThan(new BigNumber(877)))) {
            _coll1.add((_2_v1).plus(new BigNumber((_dafny.Set.fromElements(false)).length)));
          }
        }
        return _coll1;
      }())).Intersect(_dafny.Set.fromElements(new BigNumber(-216)));
    };
    static fm2(p0, p1, p2, p3, globalState) {
      return _module.__default.safeDivisionInt(new BigNumber(((_dafny.Set.fromElements(false)).Intersect(_dafny.Set.fromElements(false))).length), new BigNumber(81));
    };
    static fm4(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.UnicodeFromString("nargnu");
    };
    static fm5(p0, globalState) {
      let _source0 = _module.D2.create_DC2(new _dafny.CodePoint('a'.codePointAt(0)));
      if (_source0.is_DC3) {
        let _3___mcc_h0 = (_source0).cf3;
        let _4___mcc_h1 = (_source0).cf4;
        let _5_cf4 = _4___mcc_h1;
        let _6_cf3 = _3___mcc_h0;
        return new _dafny.CodePoint('x'.codePointAt(0));
      } else if (_source0.is_DC4) {
        let _7___mcc_h2 = (_source0).cf5;
        let _8___mcc_h3 = (_source0).cf6;
        let _9_cf6 = _8___mcc_h3;
        let _10_cf5 = _7___mcc_h2;
        return (_module.D2.create_DC2(new _dafny.CodePoint('l'.codePointAt(0)))).dtor_cf2;
      } else {
        let _11___mcc_h4 = (_source0).cf2;
        let _12_cf2 = _11___mcc_h4;
        return _12_cf2;
      }
    };
    static fm6(p0, p1, globalState) {
      return ((_dafny.MultiSet.fromElements(!(false), true, !(true))).Difference(_dafny.MultiSet.fromElements(false, false))).Difference((_dafny.MultiSet.FromArray(_dafny.Seq.of(true, true, true))).Difference(_dafny.MultiSet.fromElements(false, false, !(!(!(true))))));
    };
    static fm7(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-644),new BigNumber(320))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true, false))).cardinality()),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(-380))).length))))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-569),new BigNumber(-889))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(432),new BigNumber(210))));
    };
    static fm8(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(601)), function (_13_i0) {
        return new BigNumber(531);
      }), _dafny.Seq.of(new BigNumber(238))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(779)), function (_14_i1) {
        return new BigNumber(209);
      }));
    };
    static fm9(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber(827))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length))).cardinality()),_dafny.Set.fromElements(_module.D2.create_DC3(false, true)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-47),function () {
        let _coll2 = new _dafny.Set();
        for (const _compr_2 of (_dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC3(!(!(false)), true),_dafny.Set.fromElements(false))).Keys.Elements) {
          let _15_v0 = _compr_2;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC3(!(!(false)), true),_dafny.Set.fromElements(false))).contains(_15_v0)) {
            _coll2.add(_15_v0);
          }
        }
        return _coll2;
      }()))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-487),_dafny.Set.fromElements(_module.D2.create_DC3(false, true)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-654),!(true))).length),_dafny.Set.fromElements(_module.D2.create_DC3(true, true)))));
    };
    static fm10(globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(220)), function (_16_i0) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(493)), function (_17_i1) {
          return new _dafny.CodePoint('n'.codePointAt(0));
        });
      });
    };
    static fm11(p0, globalState) {
      return _dafny.Seq.of(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("vwsaa"), _dafny.Seq.UnicodeFromString("dmyeoj")), !(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length),false)).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((function () {
        let _coll3 = new _dafny.Set();
        for (const _compr_3 of _dafny.IntegerRange(new BigNumber(808), new BigNumber(-796))) {
          let _18_v0 = _compr_3;
          if (((new BigNumber(808)).isLessThanOrEqualTo(_18_v0)) && ((_18_v0).isLessThan(new BigNumber(-796)))) {
            _coll3.add((_18_v0).plus(new BigNumber(10)));
          }
        }
        return _coll3;
      }()).length)), new BigNumber(754), new BigNumber(462), new BigNumber((_dafny.Seq.UnicodeFromString("tkpurq")).length))).length),!(false))), !_dafny.areEqual(_dafny.Seq.of(new BigNumber(898), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("agfup")).length))).length)), _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-752))).length), new BigNumber((_dafny.Seq.of(false)).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(!(true))).length), (_module.D4.create_DC7(new BigNumber(102), true)).dtor_cf9))).length))).cardinality()), new BigNumber(86), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(741))).length))));
    };
    static fm12(p0, globalState) {
      if (!(((_module.D8.create_DC17(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(977)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(101))))).dtor_cf28).IsDisjointFrom(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((function () {
        let _coll4 = new _dafny.Map();
        for (const _compr_4 of _dafny.IntegerRange(new BigNumber(857), new BigNumber(355))) {
          let _19_v0 = _compr_4;
          if (((new BigNumber(857)).isLessThanOrEqualTo(_19_v0)) && ((_19_v0).isLessThan(new BigNumber(355)))) {
            _coll4.push([_module.__default.safeDivisionInt(_19_v0, new BigNumber(886)),_dafny.Seq.of(true)]);
          }
        }
        return _coll4;
      }()).length)))))) {
        return _module.D2.create_DC3(!(true), false);
      } else {
        return _module.D2.create_DC3(!(true), false);
      }
    };
    static fm13(p0, p1, p2, p3, globalState) {
      if (true) {
        return (_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("b"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-551)), function (_20_i0) {
          return new _dafny.CodePoint('q'.codePointAt(0));
        }), _dafny.Seq.UnicodeFromString("w"))).Union(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("l")));
      } else {
        return _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("ll"));
      }
    };
    static fm14(p0, p1, p2, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(true),!(!(false)))).Merge(function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of (_dafny.Seq.of(_dafny.MultiSet.fromElements(!(true)))).Elements) {
          let _21_v0 = _compr_5;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.MultiSet.fromElements(!(true))), _21_v0)) {
            _coll5.push([_21_v0,false]);
          }
        }
        return _coll5;
      }())).Merge((function () {
        let _coll6 = new _dafny.Map();
        for (const _compr_6 of (_dafny.MultiSet.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.of(false)), _dafny.MultiSet.fromElements(true), _dafny.MultiSet.fromElements(true, false), _dafny.MultiSet.fromElements(false))).Elements) {
          let _22_v1 = _compr_6;
          if ((_dafny.MultiSet.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.of(false)), _dafny.MultiSet.fromElements(true), _dafny.MultiSet.fromElements(true, false), _dafny.MultiSet.fromElements(false))).contains(_22_v1)) {
            _coll6.push([_22_v1,false]);
          }
        }
        return _coll6;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements((_module.D8.create_DC18((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber(-998))).length)), true)).dtor_cf30),!(true))));
    };
    static fm15(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,false))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,true)));
    };
    static fm16(globalState) {
      let _source1 = _module.D6.create_DC13();
      if (_source1.is_DC11) {
        let _23___mcc_h0 = (_source1).cf17;
        let _24___mcc_h1 = (_source1).cf18;
        let _25___mcc_h2 = (_source1).cf19;
        let _26___mcc_h3 = (_source1).cf20;
        let _27___mcc_h4 = (_source1).cf21;
        let _28_cf21 = _27___mcc_h4;
        let _29_cf20 = _26___mcc_h3;
        let _30_cf19 = _25___mcc_h2;
        let _31_cf18 = _24___mcc_h1;
        let _32_cf17 = _23___mcc_h0;
        return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(true, true),_28_cf21);
      } else if (_source1.is_DC12) {
        let _33___mcc_h5 = (_source1).cf22;
        let _34_cf22 = _33___mcc_h5;
        return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_34_cf22),new BigNumber(-495));
      } else if (_source1.is_DC13) {
        return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(!(false), true, true, true, false),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("hw")).length)));
      } else if (_source1.is_DC10) {
        let _35___mcc_h6 = (_source1).cf16;
        let _36_cf16 = _35___mcc_h6;
        return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(true),new BigNumber(414));
      } else {
        let _37___mcc_h7 = (_source1).cf23;
        let _38_cf23 = _37___mcc_h7;
        return _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(true),new BigNumber(-560));
      }
    };
    static fm17(p0, p1, p2, globalState) {
      return _dafny.Set.fromElements(true);
    };
    static m1(p0, p1, globalState) {
      let r0 = [];
      let _39_v0;
      _39_v0 = false;
      let _index0 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length));
      (p1)[_index0] = _39_v0;
      let _40_v1;
      _40_v1 = _dafny.Seq.UnicodeFromString("xdfamt");
      let _41_v2;
      _41_v2 = _dafny.MultiSet.fromElements(p0, (_dafny.ZERO).minus(p0));
      let _42_v3;
      _42_v3 = _dafny.Seq.of(p0, p0, p0, new BigNumber((_40_v1).length), (((_41_v2).contains(p0)) ? ((_41_v2).get(p0)) : (p0)));
      let _index1 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length));
      (p1)[_index1] = !(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(p0, p0, new BigNumber((_dafny.Seq.UnicodeFromString("benlgnvi")).length), p0), _42_v3)) || (_39_v0);
      let _43_v4;
      _43_v4 = _module.D2.create_DC3((p1)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length))], (p1)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length))]);
      let _44_v5;
      let _nw0 = Array((new BigNumber(3)).toNumber());
      _nw0[(_dafny.ZERO).toNumber()] = _module.D2.create_DC3((p1)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length))], (p1)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length))]);
      _nw0[(_dafny.ONE).toNumber()] = _43_v4;
      _nw0[(new BigNumber(2)).toNumber()] = _43_v4;
      _44_v5 = _nw0;
      let _45_v6;
      _45_v6 = _dafny.Seq.of(_44_v5);
      _45_v6 = _45_v6;
      let _46_v7;
      _46_v7 = _dafny.Seq.of((p1)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length))]);
      if ((!_dafny.Seq.contains(_46_v7, (p1)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length))])) && (!(!(_module.__default.fm0(_39_v0, true, new BigNumber(444), p0, globalState)) || (!((p1)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length))]))))) {
        let _47_v8;
        _47_v8 = (p1)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length))];
        let _48_v9;
        let _nw1 = new _module.C1();
        _nw1.__ctor(_47_v8);
        _48_v9 = _nw1;
        let _49_v10;
        _49_v10 = new _dafny.CodePoint('i'.codePointAt(0));
        let _50_v11;
        _50_v11 = _module.D5.create_DC9(_40_v1, _39_v0, _48_v9, _49_v10);
        let _pat_let_tv0 = _48_v9;
        let _source2 = function (_pat_let0_0) {
          return function (_51_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_52_dt__update_hcf14_h0) {
                return _module.D5.create_DC9((_51_dt__update__tmp_h0).dtor_cf12, (_51_dt__update__tmp_h0).dtor_cf13, _52_dt__update_hcf14_h0, (_51_dt__update__tmp_h0).dtor_cf15);
              }(_pat_let1_0);
            }(_pat_let_tv0);
          }(_pat_let0_0);
        }(((!(_39_v0)) ? (_module.D5.create_DC9(_dafny.Seq.UnicodeFromString("j"), _39_v0, _48_v9, _49_v10)) : (_50_v11)));
        if (_source2.is_DC9) {
          let _53___mcc_h0 = (_source2).cf12;
          let _54___mcc_h1 = (_source2).cf13;
          let _55___mcc_h2 = (_source2).cf14;
          let _56___mcc_h3 = (_source2).cf15;
          let _57_cf15 = _56___mcc_h3;
          let _58_cf14 = _55___mcc_h2;
          let _59_cf13 = _54___mcc_h1;
          let _60_cf12 = _53___mcc_h0;
          let _61_v12;
          _61_v12 = _dafny.Set.fromElements(_59_cf13);
          let _62_v13;
          _62_v13 = _dafny.Map.Empty.slice().updateUnsafe(_61_v12,new BigNumber((_dafny.MultiSet.fromElements((p1)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length))])).cardinality()));
          let _63_v14;
          let _nw2 = new _module.C0();
          _nw2.__ctor((_60_cf12)[_module.__default.safeIndex(p0, new BigNumber((_60_cf12).length))]);
          _63_v14 = _nw2;
          let _64_v15;
          _64_v15 = _dafny.Map.Empty.slice().updateUnsafe(!(_module.__default.fm16(globalState)).equals(_62_v13),_63_v14);
          let _rhs0 = !(function () {
            let _coll7 = new _dafny.Set();
            for (const _compr_7 of _dafny.IntegerRange(new BigNumber(294), new BigNumber(485))) {
              let _65_v16 = _compr_7;
              if (((new BigNumber(294)).isLessThanOrEqualTo(_65_v16)) && ((_65_v16).isLessThan(new BigNumber(485)))) {
                _coll7.add((_65_v16).minus(new BigNumber((_dafny.Seq.of(_40_v1, _60_cf12)).length)));
              }
            }
            return _coll7;
          }()).contains(new BigNumber(((((p1)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length))]) ? (_46_v7) : (_46_v7))).length));
          let _rhs1 = (_39_v0) === (_module.__default.fm0(_39_v0, _39_v0, p0, p0, globalState));
          let _rhs2 = _64_v15;
          let _lhs0 = globalState;
          _39_v0 = _rhs0;
          _lhs0.f9 = _rhs1;
          _64_v15 = _rhs2;
          let _66_v17;
          let _nw3 = new _module.C0();
          _nw3.__ctor(_49_v10);
          _66_v17 = _nw3;
          let _67_v18;
          _67_v18 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_41_v2).cardinality()),p1);
          let _68_v19;
          _68_v19 = _module.D5.create_DC8((_67_v18).Merge(_67_v18));
          _68_v19 = ((_39_v0) ? (_68_v19) : (_68_v19));
          let _69_v20;
          _69_v20 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_module.__default.fm2(_39_v0, p0, p0, _60_cf12, globalState)),_49_v10);
          let _70_v21;
          _70_v21 = _dafny.Map.Empty.slice().updateUnsafe(_39_v0,p0);
          let _71_v22;
          _71_v22 = _dafny.Map.Empty.slice().updateUnsafe(p0,_70_v21);
          let _72_v23;
          let _nw4 = Array((new BigNumber(28)).toNumber());
          _nw4[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('k'.codePointAt(0));
          _nw4[(_dafny.ONE).toNumber()] = _57_cf15;
          _nw4[(new BigNumber(2)).toNumber()] = _57_cf15;
          _nw4[(new BigNumber(3)).toNumber()] = _57_cf15;
          _nw4[(new BigNumber(4)).toNumber()] = (((_69_v20).contains(new BigNumber(285))) ? ((_69_v20).get(new BigNumber(285))) : (_49_v10));
          _nw4[(new BigNumber(5)).toNumber()] = _57_cf15;
          _nw4[(new BigNumber(6)).toNumber()] = new _dafny.CodePoint('e'.codePointAt(0));
          _nw4[(new BigNumber(7)).toNumber()] = (_63_v14).f12;
          _nw4[(new BigNumber(8)).toNumber()] = (_66_v17).f12;
          _nw4[(new BigNumber(9)).toNumber()] = _49_v10;
          _nw4[(new BigNumber(10)).toNumber()] = (_66_v17).f12;
          _nw4[(new BigNumber(11)).toNumber()] = _49_v10;
          _nw4[(new BigNumber(12)).toNumber()] = (_66_v17).f12;
          _nw4[(new BigNumber(13)).toNumber()] = _57_cf15;
          _nw4[(new BigNumber(14)).toNumber()] = (_40_v1)[_module.__default.safeIndex(new BigNumber(-569), new BigNumber((_40_v1).length))];
          _nw4[(new BigNumber(15)).toNumber()] = _49_v10;
          _nw4[(new BigNumber(16)).toNumber()] = (_63_v14).f12;
          _nw4[(new BigNumber(17)).toNumber()] = (_63_v14).f12;
          _nw4[(new BigNumber(18)).toNumber()] = _module.__default.fm5(new BigNumber((_71_v22).length), globalState);
          _nw4[(new BigNumber(19)).toNumber()] = _49_v10;
          _nw4[(new BigNumber(20)).toNumber()] = _49_v10;
          _nw4[(new BigNumber(21)).toNumber()] = (_66_v17).f12;
          _nw4[(new BigNumber(22)).toNumber()] = _49_v10;
          _nw4[(new BigNumber(23)).toNumber()] = _57_cf15;
          _nw4[(new BigNumber(24)).toNumber()] = _49_v10;
          _nw4[(new BigNumber(25)).toNumber()] = (_63_v14).f12;
          _nw4[(new BigNumber(26)).toNumber()] = new _dafny.CodePoint('f'.codePointAt(0));
          _nw4[(new BigNumber(27)).toNumber()] = (_63_v14).f12;
          _72_v23 = _nw4;
          let _73_v24;
          _73_v24 = _dafny.Map.Empty.slice().updateUnsafe(_59_cf13,_72_v23);
          let _74_v25;
          _74_v25 = _dafny.Seq.of((((_73_v24).contains((p1)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length))])) ? ((_73_v24).get((p1)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length))])) : (_72_v23)));
          _39_v0 = !(true) || (!_dafny.Seq.contains(_dafny.Seq.update(_74_v25, _module.__default.safeIndex(p0, new BigNumber((_74_v25).length)), _72_v23), _72_v23));
        } else {
          let _75___mcc_h4 = (_source2).cf11;
          let _76_cf11 = _75___mcc_h4;
          let _index2 = _module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length));
          (p1)[_index2] = (p1)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length))];
          (globalState).f4 = (p0).multipliedBy(p0);
          let _77_v26;
          let _nw5 = Array((new BigNumber(20)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _77_v26 = _nw5;
          let _index3 = _module.__default.safeIndex(new BigNumber(104), new BigNumber((_77_v26).length));
          (_77_v26)[_index3] = _49_v10;
          let _index4 = _module.__default.safeIndex(new BigNumber(104), new BigNumber((_77_v26).length));
          let _rhs3 = new BigNumber(-951);
          let _rhs4 = _40_v1;
          let _rhs5 = _49_v10;
          let _lhs1 = globalState;
          let _lhs2 = _77_v26;
          let _lhs3 = _module.__default.safeIndex(new BigNumber(104), new BigNumber((_77_v26).length));
          _lhs1.f4 = _rhs3;
          _40_v1 = _rhs4;
          _lhs2[_lhs3] = _rhs5;
          let _78_v27;
          _78_v27 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_77_v26)[_module.__default.safeIndex(new BigNumber(104), new BigNumber((_77_v26).length))]);
          let _79_v28;
          _79_v28 = _dafny.Map.Empty.slice().updateUnsafe(p0,_78_v27);
          (globalState).f4 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt(p0, p0), (((p1)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length))]) ? (p0) : ((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(((((_79_v28).contains(p0)) ? ((_79_v28).get(p0)) : (_dafny.Map.Empty.slice().updateUnsafe(p0,_49_v10)))).length))))));
        }
        let _80_v29;
        _80_v29 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("hqp"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(12)), ((_81_v10) => function (_82_i0) {
          return _81_v10;
        })(_49_v10)));
        let _83_v30;
        _83_v30 = _dafny.MultiSet.fromElements(_39_v0);
        let _84_v31;
        _84_v31 = _dafny.Map.Empty.slice().updateUnsafe(p0,_83_v30);
        let _85_v32;
        _85_v32 = _dafny.Map.Empty.slice().updateUnsafe((p1)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length))],_39_v0);
        let _86_v33;
        _86_v33 = _dafny.Map.Empty.slice().updateUnsafe(p0,(p1)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length))]);
        let _87_v34;
        _87_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_86_v33).length),p0);
        let _88_v35;
        let _nw6 = Array((new BigNumber(21)).toNumber());
        _nw6[(_dafny.ZERO).toNumber()] = p0;
        _nw6[(_dafny.ONE).toNumber()] = (p0).multipliedBy(p0);
        _nw6[(new BigNumber(2)).toNumber()] = p0;
        _nw6[(new BigNumber(3)).toNumber()] = (((_80_v29).contains(_40_v1)) ? ((_80_v29).get(_40_v1)) : (new BigNumber((_40_v1).length)));
        _nw6[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(255)), ((_89_v1) => function (_90_i1) {
          return _89_v1;
        })(_40_v1))).length);
        _nw6[(new BigNumber(5)).toNumber()] = p0;
        _nw6[(new BigNumber(6)).toNumber()] = p0;
        _nw6[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(p0);
        _nw6[(new BigNumber(8)).toNumber()] = (p0).minus(p0);
        _nw6[(new BigNumber(9)).toNumber()] = p0;
        _nw6[(new BigNumber(10)).toNumber()] = new BigNumber((_84_v31).length);
        _nw6[(new BigNumber(11)).toNumber()] = new BigNumber(936);
        _nw6[(new BigNumber(12)).toNumber()] = p0;
        _nw6[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(((_85_v32).update(_39_v0, true)).length));
        _nw6[(new BigNumber(14)).toNumber()] = p0;
        _nw6[(new BigNumber(15)).toNumber()] = p0;
        _nw6[(new BigNumber(16)).toNumber()] = p0;
        _nw6[(new BigNumber(17)).toNumber()] = p0;
        _nw6[(new BigNumber(18)).toNumber()] = (p0).multipliedBy((((_83_v30).contains(_39_v0)) ? ((_83_v30).get(_39_v0)) : (new BigNumber((_87_v34).length))));
        _nw6[(new BigNumber(19)).toNumber()] = new BigNumber(405);
        _nw6[(new BigNumber(20)).toNumber()] = ((_39_v0) ? (p0) : (p0));
        _88_v35 = _nw6;
        let _index5 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_88_v35).length));
        (_88_v35)[_index5] = p0;
        let _index6 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_88_v35).length));
        (_88_v35)[_index6] = (((_87_v34).contains((_module.D4.create_DC7((((_83_v30).contains(_39_v0)) ? ((_83_v30).get(_39_v0)) : (p0)), (p1)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length))])).dtor_cf9)) ? ((_87_v34).get((_module.D4.create_DC7((((_83_v30).contains(_39_v0)) ? ((_83_v30).get(_39_v0)) : (p0)), (p1)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length))])).dtor_cf9)) : (p0));
        let _91_v36;
        let _nw7 = new _module.C0();
        _nw7.__ctor(((_module.__default.fm0((p1)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length))], true, (_dafny.ZERO).minus(p0), (_88_v35)[_module.__default.safeIndex(new BigNumber(878), new BigNumber((_88_v35).length))], globalState)) ? (_module.__default.fm5(new BigNumber((_dafny.Seq.update(_40_v1, _module.__default.safeIndex(new BigNumber(768), new BigNumber((_40_v1).length)), _49_v10)).length), globalState)) : (_49_v10)));
        _91_v36 = _nw7;
        (globalState).f10 = !((p1)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length))]);
        let _92_v37;
        _92_v37 = _dafny.Set.fromElements(_40_v1, _40_v1);
        let _nw8 = Array((new BigNumber(12)).toNumber());
        _nw8[(_dafny.ZERO).toNumber()] = (_88_v35)[_module.__default.safeIndex(new BigNumber(878), new BigNumber((_88_v35).length))];
        _nw8[(_dafny.ONE).toNumber()] = (_88_v35)[_module.__default.safeIndex(new BigNumber(878), new BigNumber((_88_v35).length))];
        _nw8[(new BigNumber(2)).toNumber()] = p0;
        _nw8[(new BigNumber(3)).toNumber()] = p0;
        _nw8[(new BigNumber(4)).toNumber()] = p0;
        _nw8[(new BigNumber(5)).toNumber()] = (_88_v35)[_module.__default.safeIndex(new BigNumber(878), new BigNumber((_88_v35).length))];
        _nw8[(new BigNumber(6)).toNumber()] = (_88_v35)[_module.__default.safeIndex(new BigNumber(878), new BigNumber((_88_v35).length))];
        _nw8[(new BigNumber(7)).toNumber()] = new BigNumber((_92_v37).length);
        _nw8[(new BigNumber(8)).toNumber()] = (_88_v35)[_module.__default.safeIndex(new BigNumber(878), new BigNumber((_88_v35).length))];
        _nw8[(new BigNumber(9)).toNumber()] = _module.__default.safeModuloInt(p0, p0);
        _nw8[(new BigNumber(10)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_46_v7, _46_v7)).length);
        _nw8[(new BigNumber(11)).toNumber()] = (_88_v35)[_module.__default.safeIndex(new BigNumber(878), new BigNumber((_88_v35).length))];
        _88_v35 = _nw8;
      } else {
        let _93_v38;
        _93_v38 = !(_39_v0);
        let _94_v39;
        let _nw9 = new _module.C1();
        _nw9.__ctor(_93_v38);
        _94_v39 = _nw9;
        let _95_v40;
        _95_v40 = _dafny.Map.Empty.slice().updateUnsafe(_94_v39,_39_v0);
        let _96_v41;
        _96_v41 = _module.D6.create_DC10(_95_v40);
        let _97_v42;
        _97_v42 = _dafny.Map.Empty.slice().updateUnsafe(_41_v2,(_96_v41).dtor_cf16);
        let _98_v43;
        _98_v43 = _dafny.Seq.of(_95_v40);
        let _99_v44;
        _99_v44 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-989),_95_v40);
        let _100_v45;
        _100_v45 = _dafny.Map.Empty.slice().updateUnsafe((p1)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length))],_94_v39);
        let _101_v46;
        _101_v46 = _dafny.Seq.of(_94_v39);
        let _102_v47;
        let _nw10 = Array((new BigNumber(25)).toNumber());
        _nw10[(_dafny.ZERO).toNumber()] = (((_97_v42).contains(_41_v2)) ? ((_97_v42).get(_41_v2)) : (_95_v40));
        _nw10[(_dafny.ONE).toNumber()] = _95_v40;
        _nw10[(new BigNumber(2)).toNumber()] = _95_v40;
        _nw10[(new BigNumber(3)).toNumber()] = _95_v40;
        _nw10[(new BigNumber(4)).toNumber()] = _95_v40;
        _nw10[(new BigNumber(5)).toNumber()] = (_95_v40).Merge(_95_v40);
        _nw10[(new BigNumber(6)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_94_v39,(p1)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length))]);
        _nw10[(new BigNumber(7)).toNumber()] = _95_v40;
        _nw10[(new BigNumber(8)).toNumber()] = (_98_v43)[_module.__default.safeIndex(p0, new BigNumber((_98_v43).length))];
        _nw10[(new BigNumber(9)).toNumber()] = _95_v40;
        _nw10[(new BigNumber(10)).toNumber()] = (_95_v40).Merge(_95_v40);
        _nw10[(new BigNumber(11)).toNumber()] = _95_v40;
        _nw10[(new BigNumber(12)).toNumber()] = _95_v40;
        _nw10[(new BigNumber(13)).toNumber()] = ((_95_v40).update(_94_v39, true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_94_v39,_39_v0));
        _nw10[(new BigNumber(14)).toNumber()] = _95_v40;
        _nw10[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_94_v39,_module.__default.fm0(_39_v0, _module.__default.fm0(true, _39_v0, p0, p0, globalState), p0, p0, globalState));
        _nw10[(new BigNumber(16)).toNumber()] = _95_v40;
        _nw10[(new BigNumber(17)).toNumber()] = ((((_99_v44).contains(new BigNumber((_42_v3).length))) ? ((_99_v44).get(new BigNumber((_42_v3).length))) : (_95_v40))).Merge(_95_v40);
        _nw10[(new BigNumber(18)).toNumber()] = _95_v40;
        _nw10[(new BigNumber(19)).toNumber()] = _95_v40;
        _nw10[(new BigNumber(20)).toNumber()] = _95_v40;
        _nw10[(new BigNumber(21)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((((_100_v45).contains(_39_v0)) ? ((_100_v45).get(_39_v0)) : (_94_v39)),_39_v0);
        _nw10[(new BigNumber(22)).toNumber()] = _95_v40;
        _nw10[(new BigNumber(23)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_101_v46)[_module.__default.safeIndex(p0, new BigNumber((_101_v46).length))],(p1)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length))]);
        _nw10[(new BigNumber(24)).toNumber()] = (_95_v40).Merge(_dafny.Map.Empty.slice().updateUnsafe(_94_v39,_39_v0));
        _102_v47 = _nw10;
        let _103_v48;
        _103_v48 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_dafny.ZERO).minus(p0));
        let _104_v49;
        let _nw11 = Array((new BigNumber(3)).toNumber());
        _nw11[(_dafny.ZERO).toNumber()] = p0;
        _nw11[(_dafny.ONE).toNumber()] = (((_103_v48).contains(p0)) ? ((_103_v48).get(p0)) : (p0));
        _nw11[(new BigNumber(2)).toNumber()] = p0;
        _104_v49 = _nw11;
        let _index7 = _module.__default.safeIndex(new BigNumber(242), new BigNumber((_104_v49).length));
        (_104_v49)[_index7] = _module.__default.fm2(_39_v0, p0, p0, _40_v1, globalState);
        let _105_v50;
        let _nw12 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Set.Empty);
        _105_v50 = _nw12;
        let _106_v51;
        _106_v51 = _dafny.Set.fromElements(_39_v0);
        let _index8 = _module.__default.safeIndex(new BigNumber(267), new BigNumber((_105_v50).length));
        (_105_v50)[_index8] = _106_v51;
        let _107_v52;
        _107_v52 = _dafny.Map.Empty.slice().updateUnsafe(_41_v2,_102_v47);
        let _108_v53;
        _108_v53 = new _dafny.CodePoint('h'.codePointAt(0));
        let _109_v54;
        let _nw13 = new _module.C0();
        _nw13.__ctor(_108_v53);
        _109_v54 = _nw13;
        let _110_v55;
        _110_v55 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Map.Empty.slice().updateUnsafe(p0,(p1)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length))]));
        let _index9 = _module.__default.safeIndex(new BigNumber(242), new BigNumber((_104_v49).length));
        let _index10 = _module.__default.safeIndex(new BigNumber(267), new BigNumber((_105_v50).length));
        let _rhs6 = (((_107_v52).contains(_dafny.MultiSet.fromElements(new BigNumber(((_module.D6.create_DC11(_40_v1, _109_v54, new BigNumber((_dafny.Seq.of(_109_v54)).length), _dafny.Seq.of(p0, p0), p0)).dtor_cf20).length), p0, p0, p0, new BigNumber(-582)))) ? ((_107_v52).get(_dafny.MultiSet.fromElements(new BigNumber(((_module.D6.create_DC11(_40_v1, _109_v54, new BigNumber((_dafny.Seq.of(_109_v54)).length), _dafny.Seq.of(p0, p0), p0)).dtor_cf20).length), p0, p0, p0, new BigNumber(-582)))) : (_102_v47));
        let _rhs7 = p0;
        let _rhs8 = (_dafny.ZERO).minus(p0);
        let _rhs9 = _39_v0;
        let _rhs10 = _module.__default.fm17(_110_v55, _40_v1, p0, globalState);
        let _lhs4 = _104_v49;
        let _lhs5 = _module.__default.safeIndex(new BigNumber(242), new BigNumber((_104_v49).length));
        let _lhs6 = globalState;
        let _lhs7 = globalState;
        let _lhs8 = _105_v50;
        let _lhs9 = _module.__default.safeIndex(new BigNumber(267), new BigNumber((_105_v50).length));
        _102_v47 = _rhs6;
        _lhs4[_lhs5] = _rhs7;
        _lhs6.f4 = _rhs8;
        _lhs7.f9 = _rhs9;
        _lhs8[_lhs9] = _rhs10;
        (globalState).f9 = _39_v0;
        (globalState).f4 = (p0).plus(p0);
        let _111_v56;
        let _init0 = ((_112_v7) => function (_113_i2) {
          return _112_v7;
        })(_46_v7);
        let _nw14 = Array((new BigNumber(16)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw14.length); _i0_0++) {
          _nw14[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _111_v56 = _nw14;
        let _index11 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_111_v56).length));
        (_111_v56)[_index11] = _dafny.Seq.Concat(_46_v7, _46_v7);
        let _114_v57;
        _114_v57 = _module.D7.create_DC15(_46_v7);
        let _index12 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_111_v56).length));
        (_111_v56)[_index12] = (_114_v57).dtor_cf24;
        let _115_v58;
        let _nw15 = Array((new BigNumber(17)).toNumber());
        _nw15[(_dafny.ZERO).toNumber()] = p1;
        _nw15[(_dafny.ONE).toNumber()] = ((_39_v0) ? (p1) : (p1));
        _nw15[(new BigNumber(2)).toNumber()] = p1;
        _nw15[(new BigNumber(3)).toNumber()] = p1;
        _nw15[(new BigNumber(4)).toNumber()] = p1;
        _nw15[(new BigNumber(5)).toNumber()] = p1;
        _nw15[(new BigNumber(6)).toNumber()] = p1;
        _nw15[(new BigNumber(7)).toNumber()] = p1;
        _nw15[(new BigNumber(8)).toNumber()] = p1;
        _nw15[(new BigNumber(9)).toNumber()] = p1;
        _nw15[(new BigNumber(10)).toNumber()] = p1;
        _nw15[(new BigNumber(11)).toNumber()] = p1;
        _nw15[(new BigNumber(12)).toNumber()] = p1;
        _nw15[(new BigNumber(13)).toNumber()] = p1;
        _nw15[(new BigNumber(14)).toNumber()] = p1;
        _nw15[(new BigNumber(15)).toNumber()] = p1;
        _nw15[(new BigNumber(16)).toNumber()] = p1;
        _115_v58 = _nw15;
        let _index13 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_115_v58).length));
        (_115_v58)[_index13] = p1;
        let _index14 = _module.__default.safeIndex(new BigNumber(85), new BigNumber((_115_v58).length));
        (_115_v58)[_index14] = p1;
      }
      let _116_v59;
      let _init1 = ((_117_v3) => function (_118_i3) {
        return _117_v3;
      })(_42_v3);
      let _nw16 = Array((new BigNumber(10)).toNumber());
      for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw16.length); _i0_1++) {
        _nw16[_i0_1] = _init1(new BigNumber(_i0_1));
      }
      _116_v59 = _nw16;
      let _index15 = _module.__default.safeIndex(new BigNumber(21), new BigNumber((_116_v59).length));
      (_116_v59)[_index15] = _42_v3;
      let _119_v60;
      let _nw17 = Array((new BigNumber(18)).toNumber());
      _119_v60 = _nw17;
      let _120_v61;
      _120_v61 = false;
      let _121_v62;
      let _nw18 = new _module.C1();
      _nw18.__ctor(_120_v61);
      _121_v62 = _nw18;
      let _index16 = _module.__default.safeIndex(new BigNumber(430), new BigNumber((_119_v60).length));
      (_119_v60)[_index16] = _121_v62;
      let _122_v63;
      _122_v63 = _dafny.Set.fromElements(p0, new BigNumber((_41_v2).cardinality()));
      let _123_v64;
      _123_v64 = _dafny.MultiSet.fromElements(_122_v63);
      let _index17 = _module.__default.safeIndex(new BigNumber(21), new BigNumber((_116_v59).length));
      let _index18 = _module.__default.safeIndex(new BigNumber(430), new BigNumber((_119_v60).length));
      let _rhs11 = _dafny.Seq.Concat(_42_v3, _dafny.Seq.Concat(_42_v3, _42_v3));
      let _rhs12 = _121_v62;
      let _rhs13 = _121_v62;
      let _rhs14 = _123_v64;
      let _lhs10 = _116_v59;
      let _lhs11 = _module.__default.safeIndex(new BigNumber(21), new BigNumber((_116_v59).length));
      let _lhs12 = _119_v60;
      let _lhs13 = _module.__default.safeIndex(new BigNumber(430), new BigNumber((_119_v60).length));
      _lhs10[_lhs11] = _rhs11;
      _lhs12[_lhs13] = _rhs12;
      _121_v62 = _rhs13;
      _123_v64 = _rhs14;
      let _124_v65;
      _124_v65 = _dafny.Set.fromElements((p1)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length))]);
      if ((_124_v65).IsSubsetOf(_124_v65)) {
        let _125_v66;
        _125_v66 = _dafny.Map.Empty.slice().updateUnsafe(_120_v61,!((p1)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length))]));
        _125_v66 = (_125_v66).update((_121_v62).f11, _39_v0);
        _40_v1 = ((_39_v0) ? (((true) ? (_40_v1) : (_40_v1))) : (_dafny.Seq.Concat(_40_v1, _40_v1)));
        let _126_v67;
        _126_v67 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_46_v7, _46_v7),_39_v0);
        _126_v67 = (_126_v67).update(_dafny.Seq.Concat(_dafny.Seq.of(!((p1)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length))])), _46_v7), (p1)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length))]);
        (globalState).f3 = p1;
        (globalState).f4 = new BigNumber(31);
      } else {
        let _127_v68;
        let _nw19 = new _module.C1();
        _nw19.__ctor((_121_v62).f11);
        _127_v68 = _nw19;
        let _128_v69;
        _128_v69 = _dafny.Map.Empty.slice().updateUnsafe(p0,_39_v0);
        let _129_v70;
        _129_v70 = _dafny.Set.fromElements(_128_v69, _128_v69);
        _129_v70 = _129_v70;
        let _130_v71;
        _130_v71 = _module.D6.create_DC13();
        let _131_v72;
        _131_v72 = _dafny.Seq.of(_130_v71, _130_v71, _130_v71);
        let _132_v73;
        _132_v73 = _module.D6.create_DC14((_131_v72)[_module.__default.safeIndex(new BigNumber((_40_v1).length), new BigNumber((_131_v72).length))]);
        _132_v73 = _132_v73;
        _40_v1 = _dafny.Seq.Concat(_40_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(771)), function (_133_i4) {
          return new _dafny.CodePoint('y'.codePointAt(0));
        }));
        (globalState).f9 = (p1)[_module.__default.safeIndex(new BigNumber(666), new BigNumber((p1).length))];
      }
      let _134_v74;
      let _out0;
      _out0 = (_121_v62).m0(new BigNumber((_40_v1).length), _40_v1, _39_v0, p0, globalState);
      _134_v74 = _out0;
      r0 = p1;
      return r0;
    }
    static Main(__noArgsParameter) {
      let _135_v0;
      let _nw20 = Array((new BigNumber(14)).toNumber()).fill(false);
      _135_v0 = _nw20;
      let _136_globalState;
      let _nw21 = new _module.GlobalState();
      _nw21.__ctor(_135_v0, new BigNumber(319), false, _135_v0, new BigNumber(250), new BigNumber(888), false, new BigNumber(746), false, false, true);
      _136_globalState = _nw21;
      let _137_v1;
      _137_v1 = true;
      let _138_v2;
      _138_v2 = _dafny.MultiSet.fromElements(_137_v1);
      let _139_v3;
      _139_v3 = _dafny.Map.Empty.slice().updateUnsafe((_138_v2).Union(_138_v2),_137_v1);
      _139_v3 = (_139_v3).Merge(_139_v3);
      let _140_v4;
      _140_v4 = new BigNumber(384);
      let _hi0 = ((_dafny.ZERO).minus(_140_v4)).minus(_140_v4);
      for (let _141_i0 = _140_v4; _141_i0.isLessThan(_hi0); _141_i0 = _141_i0.plus(_dafny.ONE)) {
        let _142_v5;
        _142_v5 = _dafny.Seq.of(_137_v1);
        _137_v1 = ((_module.__default.fm0(_137_v1, _137_v1, new BigNumber((_dafny.Seq.UnicodeFromString("bfjvu")).length), new BigNumber((_142_v5).length), _136_globalState)) ? (_137_v1) : (_module.__default.fm0(_137_v1, !(true), _140_v4, _141_i0, _136_globalState)));
        let _143_v6;
        _143_v6 = _dafny.Seq.of(_135_v0);
        _143_v6 = _dafny.Seq.of(_135_v0, _135_v0);
        let _144_v7;
        let _init2 = ((_145_i0) => function (_146_i1) {
          return _dafny.Map.Empty.slice().updateUnsafe(_145_i0,new _dafny.CodePoint('g'.codePointAt(0)));
        })(_141_i0);
        let _nw22 = Array((new BigNumber(23)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw22.length); _i0_2++) {
          _nw22[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _144_v7 = _nw22;
        let _147_v8;
        _147_v8 = new _dafny.CodePoint('o'.codePointAt(0));
        let _148_v9;
        _148_v9 = _dafny.Map.Empty.slice().updateUnsafe(_140_v4,_147_v8);
        let _index19 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_144_v7).length));
        (_144_v7)[_index19] = _148_v9;
        let _149_v10;
        _149_v10 = false;
        let _index20 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_144_v7).length));
        (_144_v7)[_index20] = _dafny.Map.Empty.slice().updateUnsafe((_141_i0).minus(new BigNumber((_module.__default.fm1(_dafny.Map.Empty.slice().updateUnsafe(_141_i0,_140_v4), (_149_v10), _141_i0, _140_v4, _136_globalState)).length)),_147_v8);
        let _150_v11;
        _150_v11 = _dafny.Seq.UnicodeFromString("ewk");
        _150_v11 = _150_v11;
      }
      let _151_v12;
      _151_v12 = _dafny.Map.Empty.slice().updateUnsafe((_140_v4).minus(new BigNumber(261)),!(_module.__default.fm0(_137_v1, _137_v1, _140_v4, new BigNumber(789), _136_globalState)));
      _151_v12 = _151_v12;
      let _152_v13;
      _152_v13 = _dafny.Map.Empty.slice().updateUnsafe(_137_v1,_dafny.Seq.UnicodeFromString("du"));
      let _153_v14;
      _153_v14 = _dafny.Seq.of(_140_v4, _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_152_v13).length)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(671)), function (_154_i2) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      })).length)));
      let _155_v15;
      _155_v15 = _137_v1;
      let _156_v16;
      _156_v16 = _dafny.Map.Empty.slice().updateUnsafe(_137_v1,_140_v4);
      let _157_v17;
      _157_v17 = _dafny.Seq.UnicodeFromString("ajftdewh");
      (_136_globalState).f4 = (_153_v14)[_module.__default.safeIndex(_module.__default.fm2(_137_v1, (((_156_v16).contains(_155_v15)) ? ((_156_v16).get(_155_v15)) : (_140_v4)), _140_v4, _157_v17, _136_globalState), new BigNumber((_153_v14).length))];
      let _158_v18;
      _158_v18 = _dafny.Map.Empty.slice().updateUnsafe(_137_v1,((_137_v1) ? (false) : (_137_v1)));
      _158_v18 = (_158_v18).update(_137_v1, _137_v1);
      let _source3 = _155_v15;
      let _159___mcc_h0 = _source3;
      let _160_cf0 = _159___mcc_h0;
      let _161_v19;
      _161_v19 = new _dafny.CodePoint('t'.codePointAt(0));
      let _162_v20;
      let _nw23 = new _module.C0();
      _nw23.__ctor(_161_v19);
      _162_v20 = _nw23;
      let _163_v21;
      _163_v21 = _dafny.Seq.of(_157_v17);
      let _164_v22;
      let _nw24 = Array((new BigNumber(18)).toNumber());
      _nw24[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_157_v17, _157_v17);
      _nw24[(_dafny.ONE).toNumber()] = _dafny.Seq.UnicodeFromString("kwmdrfli");
      _nw24[(new BigNumber(2)).toNumber()] = _157_v17;
      _nw24[(new BigNumber(3)).toNumber()] = (_163_v21)[_module.__default.safeIndex(_140_v4, new BigNumber((_163_v21).length))];
      _nw24[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(101)), ((_165_v19) => function (_166_i3) {
        return _165_v19;
      })(_161_v19));
      _nw24[(new BigNumber(5)).toNumber()] = _157_v17;
      _nw24[(new BigNumber(6)).toNumber()] = (((_152_v13).contains(_module.__default.fm0(_137_v1, _137_v1, _140_v4, _140_v4, _136_globalState))) ? ((_152_v13).get(_module.__default.fm0(_137_v1, _137_v1, _140_v4, _140_v4, _136_globalState))) : (_157_v17));
      _nw24[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(195)), ((_167_v19) => function (_168_i4) {
        return _167_v19;
      })(_161_v19)), _157_v17);
      _nw24[(new BigNumber(8)).toNumber()] = _dafny.Seq.UnicodeFromString("krisnrwt");
      _nw24[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_157_v17, _dafny.Seq.Create(_module.__default.abs(new BigNumber(496)), ((_169_v20) => function (_170_i5) {
        return (_169_v20).f12;
      })(_162_v20)));
      _nw24[(new BigNumber(10)).toNumber()] = _157_v17;
      _nw24[(new BigNumber(11)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(649)), ((_171_v20) => function (_172_i6) {
        return (_171_v20).f12;
      })(_162_v20));
      _nw24[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_157_v17, _157_v17);
      _nw24[(new BigNumber(13)).toNumber()] = _157_v17;
      _nw24[(new BigNumber(14)).toNumber()] = _157_v17;
      _nw24[(new BigNumber(15)).toNumber()] = _dafny.Seq.UnicodeFromString("dkm");
      _nw24[(new BigNumber(16)).toNumber()] = _module.__default.fm4((_dafny.ZERO).minus(_140_v4), _160_cf0, _160_cf0, (_dafny.ZERO).minus(_140_v4), _136_globalState);
      _nw24[(new BigNumber(17)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_157_v17, _dafny.Seq.UnicodeFromString("qfkthhxah")), _module.__default.safeIndex(_140_v4, new BigNumber((_dafny.Seq.Concat(_157_v17, _dafny.Seq.UnicodeFromString("qfkthhxah"))).length)), _161_v19);
      _164_v22 = _nw24;
      let _index21 = _module.__default.safeIndex(new BigNumber(280), new BigNumber((_164_v22).length));
      (_164_v22)[_index21] = _157_v17;
      let _index22 = _module.__default.safeIndex(new BigNumber(280), new BigNumber((_164_v22).length));
      (_164_v22)[_index22] = _157_v17;
      let _173_v23;
      let _nw25 = Array((new BigNumber(10)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _173_v23 = _nw25;
      let _index23 = _module.__default.safeIndex(new BigNumber(16), new BigNumber((_173_v23).length));
      (_173_v23)[_index23] = (_162_v20).f12;
      let _index24 = _module.__default.safeIndex(new BigNumber(16), new BigNumber((_173_v23).length));
      (_173_v23)[_index24] = _161_v19;
      (_136_globalState).f10 = (!(!(_137_v1)) || (_137_v1)) && (_137_v1);
      (_136_globalState).f10 = _137_v1;
      let _174_v24;
      let _init3 = ((_175_v17) => function (_176_i8) {
        return (_176_i8).minus(new BigNumber((_175_v17).length));
      })(_157_v17);
      let _nw26 = Array((new BigNumber(29)).toNumber());
      for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw26.length); _i0_3++) {
        _nw26[_i0_3] = _init3(new BigNumber(_i0_3));
      }
      _174_v24 = _nw26;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_174_v24).length))) {
        let _177_i7 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_177_i7)) && ((_177_i7).isLessThan(new BigNumber((_174_v24).length))))) {
          (_174_v24)[(_177_i7)] = (_177_i7).multipliedBy(_140_v4);
        }
      }
      if (_137_v1) {
        let _index25 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_174_v24).length));
        (_174_v24)[_index25] = (_dafny.ZERO).minus(_140_v4);
        let _178_v25;
        _178_v25 = _dafny.Seq.of(_137_v1, _137_v1);
        let _index26 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_174_v24).length));
        let _rhs15 = (_140_v4).minus(_140_v4);
        let _rhs16 = !_dafny.Seq.contains(_dafny.Seq.update(_178_v25, _module.__default.safeIndex(_140_v4, new BigNumber((_178_v25).length)), _137_v1), _137_v1);
        let _rhs17 = _module.__default.safeDivisionInt(_140_v4, new BigNumber((_module.__default.fm11((_dafny.ZERO).minus(_140_v4), _136_globalState)).length));
        let _rhs18 = _137_v1;
        let _lhs14 = _174_v24;
        let _lhs15 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_174_v24).length));
        let _lhs16 = _136_globalState;
        let _lhs17 = _136_globalState;
        _lhs14[_lhs15] = _rhs15;
        _137_v1 = _rhs16;
        _lhs16.f4 = _rhs17;
        _lhs17.f9 = _rhs18;
        let _source4 = _module.__default.fm12(_140_v4, _136_globalState);
        if (_source4.is_DC3) {
          let _179___mcc_h1 = (_source4).cf3;
          let _180___mcc_h2 = (_source4).cf4;
          let _181_cf4 = _180___mcc_h2;
          let _182_cf3 = _179___mcc_h1;
          let _183_v26;
          _183_v26 = _dafny.Set.fromElements((((_151_v12).contains(_140_v4)) ? ((_151_v12).get(_140_v4)) : (_182_cf3)));
          (_136_globalState).f9 = !(_183_v26).equals(_183_v26);
          _153_v14 = _153_v14;
          let _index27 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_174_v24).length));
          (_174_v24)[_index27] = new BigNumber((_module.__default.fm11((_174_v24)[_module.__default.safeIndex(new BigNumber(891), new BigNumber((_174_v24).length))], _136_globalState)).length);
          let _184_v27;
          _184_v27 = _module.D2.create_DC3((_178_v25)[_module.__default.safeIndex((_174_v24)[_module.__default.safeIndex(new BigNumber(891), new BigNumber((_174_v24).length))], new BigNumber((_178_v25).length))], false);
          let _185_v28;
          let _nw27 = new _module.C0();
          _nw27.__ctor((((_184_v27).dtor_cf4) ? (new _dafny.CodePoint('l'.codePointAt(0))) : (new _dafny.CodePoint('k'.codePointAt(0)))));
          _185_v28 = _nw27;
        } else if (_source4.is_DC4) {
          let _186___mcc_h3 = (_source4).cf5;
          let _187___mcc_h4 = (_source4).cf6;
          let _188_cf6 = _187___mcc_h4;
          let _189_cf5 = _186___mcc_h3;
          let _190_v29;
          let _nw28 = new _module.C1();
          _nw28.__ctor(_155_v15);
          _190_v29 = _nw28;
          _190_v29 = _190_v29;
          let _191_v30;
          let _nw29 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
          _191_v30 = _nw29;
          let _rhs19 = _191_v30;
          let _rhs20 = _137_v1;
          let _rhs21 = _178_v25;
          let _rhs22 = _137_v1;
          let _lhs18 = _136_globalState;
          let _lhs19 = _136_globalState;
          _191_v30 = _rhs19;
          _lhs18.f9 = _rhs20;
          _178_v25 = _rhs21;
          _lhs19.f9 = _rhs22;
          _137_v1 = _137_v1;
          let _192_v31;
          _192_v31 = _dafny.Set.fromElements(_138_v2, _138_v2);
          let _193_v32;
          _193_v32 = _dafny.MultiSet.fromElements(_138_v2, _138_v2);
          _192_v31 = function () {
            let _coll8 = new _dafny.Set();
            for (const _compr_8 of (_193_v32).Elements) {
              let _194_v33 = _compr_8;
              if ((_193_v32).contains(_194_v33)) {
                _coll8.add(_194_v33);
              }
            }
            return _coll8;
          }();
        } else {
          let _195___mcc_h5 = (_source4).cf2;
          let _196_cf2 = _195___mcc_h5;
          _157_v17 = _dafny.Seq.Concat(_157_v17, _157_v17);
          let _197_v34;
          let _nw30 = new _module.C1();
          _nw30.__ctor(_155_v15);
          _197_v34 = _nw30;
          _197_v34 = _197_v34;
          let _198_v35;
          let _nw31 = new _module.C0();
          _nw31.__ctor(_196_cf2);
          _198_v35 = _nw31;
          _198_v35 = _198_v35;
          let _199_v36;
          _199_v36 = _dafny.Map.Empty.slice().updateUnsafe(_138_v2,_140_v4);
          let _200_v37;
          let _out1;
          _out1 = (_197_v34).m0(new BigNumber((_199_v36).length), _dafny.Seq.Concat(_157_v17, _dafny.Seq.Create(_module.__default.abs(new BigNumber(918)), ((_201_cf2) => function (_202_i9) {
            return _201_cf2;
          })(_196_cf2))), !(_137_v1) || (_137_v1), (_174_v24)[_module.__default.safeIndex(new BigNumber(891), new BigNumber((_174_v24).length))], _136_globalState);
          _200_v37 = _out1;
        }
        let _index28 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_174_v24).length));
        (_174_v24)[_index28] = (_174_v24)[_module.__default.safeIndex(new BigNumber(891), new BigNumber((_174_v24).length))];
        let _203_v38;
        _203_v38 = new _dafny.CodePoint('f'.codePointAt(0));
        let _204_v39;
        _204_v39 = _module.D2.create_DC3(_dafny.Seq.IsPrefixOf(_157_v17, _dafny.Seq.update(_157_v17, _module.__default.safeIndex(_140_v4, new BigNumber((_157_v17).length)), _203_v38)), _dafny.Seq.IsProperPrefixOf(_157_v17, _dafny.Seq.update(_157_v17, _module.__default.safeIndex(_140_v4, new BigNumber((_157_v17).length)), _203_v38)));
        _204_v39 = _204_v39;
        let _205_v40;
        let _nw32 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
        _205_v40 = _nw32;
        _205_v40 = _205_v40;
      } else {
        let _206_v41;
        _206_v41 = _dafny.Set.fromElements(_157_v17, _157_v17, _157_v17);
        _206_v41 = _module.__default.fm13(new BigNumber(839), _140_v4, _157_v17, _137_v1, _136_globalState);
        let _207_v42;
        _207_v42 = new _dafny.CodePoint('w'.codePointAt(0));
        let _208_v43;
        let _nw33 = new _module.C0();
        _nw33.__ctor(_207_v42);
        _208_v43 = _nw33;
        let _209_v44;
        _209_v44 = _dafny.Seq.of(_208_v43);
        let _210_v45;
        let _nw34 = Array((new BigNumber(21)).toNumber());
        _nw34[(_dafny.ZERO).toNumber()] = _208_v43;
        _nw34[(_dafny.ONE).toNumber()] = _208_v43;
        _nw34[(new BigNumber(2)).toNumber()] = _208_v43;
        _nw34[(new BigNumber(3)).toNumber()] = _208_v43;
        _nw34[(new BigNumber(4)).toNumber()] = _208_v43;
        _nw34[(new BigNumber(5)).toNumber()] = _208_v43;
        _nw34[(new BigNumber(6)).toNumber()] = (_209_v44)[_module.__default.safeIndex(_140_v4, new BigNumber((_209_v44).length))];
        _nw34[(new BigNumber(7)).toNumber()] = _208_v43;
        _nw34[(new BigNumber(8)).toNumber()] = _208_v43;
        _nw34[(new BigNumber(9)).toNumber()] = _208_v43;
        _nw34[(new BigNumber(10)).toNumber()] = _208_v43;
        _nw34[(new BigNumber(11)).toNumber()] = _208_v43;
        _nw34[(new BigNumber(12)).toNumber()] = _208_v43;
        _nw34[(new BigNumber(13)).toNumber()] = _208_v43;
        _nw34[(new BigNumber(14)).toNumber()] = _208_v43;
        _nw34[(new BigNumber(15)).toNumber()] = _208_v43;
        _nw34[(new BigNumber(16)).toNumber()] = _208_v43;
        _nw34[(new BigNumber(17)).toNumber()] = _208_v43;
        _nw34[(new BigNumber(18)).toNumber()] = _208_v43;
        _nw34[(new BigNumber(19)).toNumber()] = _208_v43;
        _nw34[(new BigNumber(20)).toNumber()] = _208_v43;
        _210_v45 = _nw34;
        _210_v45 = _210_v45;
        (_136_globalState).f4 = (_module.__default.fm2(_137_v1, new BigNumber(942), _140_v4, _dafny.Seq.UnicodeFromString("p"), _136_globalState)).plus(_140_v4);
        let _211_v46;
        let _out2;
        _out2 = _module.__default.m1(_140_v4, _135_v0, _136_globalState);
        _211_v46 = _out2;
        let _212_v47;
        _212_v47 = _module.D4.create_DC6(_211_v46);
        let _213_v48;
        let _out3;
        _out3 = _module.__default.m1(_module.__default.safeDivisionInt((_dafny.ZERO).minus(_140_v4), _140_v4), (_212_v47).dtor_cf8, _136_globalState);
        _213_v48 = _out3;
      }
      let _214_v49;
      let _nw35 = new _module.C1();
      _nw35.__ctor(_155_v15);
      _214_v49 = _nw35;
      _214_v49 = _214_v49;
      let _215_v50;
      _215_v50 = _dafny.Seq.of(_137_v1, _137_v1);
      _137_v1 = !_dafny.areEqual(_215_v50, _215_v50);
      if (_137_v1) {
        (_136_globalState).f10 = _137_v1;
        let _216_v52;
        _216_v52 = _dafny.Set.fromElements((new BigNumber((function () {
          let _coll9 = new _dafny.Map();
          for (const _compr_9 of _dafny.IntegerRange(new BigNumber(-52), new BigNumber(-65))) {
            let _217_v51 = _compr_9;
            if (((new BigNumber(-52)).isLessThanOrEqualTo(_217_v51)) && ((_217_v51).isLessThan(new BigNumber(-65)))) {
              _coll9.push([(_217_v51).plus(_140_v4),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_157_v17).length))).length)]);
            }
          }
          return _coll9;
        }()).length)).multipliedBy(_140_v4));
        _216_v52 = (_216_v52).Intersect(_216_v52);
        let _218_v53;
        _218_v53 = new _dafny.CodePoint('y'.codePointAt(0));
        let _219_v54;
        _219_v54 = _module.D2.create_DC4(_140_v4, _218_v53);
        let _220_v55;
        _220_v55 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber(14)),((_137_v1) ? (_140_v4) : ((_219_v54).dtor_cf5)));
        _220_v55 = (_220_v55).update(_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.of(_140_v4), _module.__default.safeIndex(_140_v4, new BigNumber((_dafny.Seq.of(_140_v4)).length)), _140_v4), _module.__default.safeIndex(_140_v4, new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_140_v4), _module.__default.safeIndex(_140_v4, new BigNumber((_dafny.Seq.of(_140_v4)).length)), _140_v4)).length)), _140_v4), _module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_137_v1)).length), new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.of(_140_v4), _module.__default.safeIndex(_140_v4, new BigNumber((_dafny.Seq.of(_140_v4)).length)), _140_v4), _module.__default.safeIndex(_140_v4, new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_140_v4), _module.__default.safeIndex(_140_v4, new BigNumber((_dafny.Seq.of(_140_v4)).length)), _140_v4)).length)), _140_v4)).length)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_137_v1,_137_v1)).length)), new BigNumber(-114));
        let _221_v56;
        let _nw36 = new _module.C0();
        _nw36.__ctor(_218_v53);
        _221_v56 = _nw36;
        let _222_v57;
        let _init4 = ((_223_v14, _224_v4) => function (_225_i10) {
          return _dafny.Seq.update(_223_v14, _module.__default.safeIndex(_224_v4, new BigNumber((_223_v14).length)), _224_v4);
        })(_153_v14, _140_v4);
        let _nw37 = Array((new BigNumber(11)).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw37.length); _i0_4++) {
          _nw37[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _222_v57 = _nw37;
        let _226_v58;
        _226_v58 = _dafny.Seq.of(_222_v57);
        let _227_v59;
        let _nw38 = Array((new BigNumber(18)).toNumber());
        _nw38[(_dafny.ZERO).toNumber()] = _222_v57;
        _nw38[(_dafny.ONE).toNumber()] = _222_v57;
        _nw38[(new BigNumber(2)).toNumber()] = _222_v57;
        _nw38[(new BigNumber(3)).toNumber()] = (_226_v58)[_module.__default.safeIndex(new BigNumber(316), new BigNumber((_226_v58).length))];
        _nw38[(new BigNumber(4)).toNumber()] = _222_v57;
        _nw38[(new BigNumber(5)).toNumber()] = _222_v57;
        _nw38[(new BigNumber(6)).toNumber()] = _222_v57;
        _nw38[(new BigNumber(7)).toNumber()] = _222_v57;
        _nw38[(new BigNumber(8)).toNumber()] = _222_v57;
        _nw38[(new BigNumber(9)).toNumber()] = _222_v57;
        _nw38[(new BigNumber(10)).toNumber()] = _222_v57;
        _nw38[(new BigNumber(11)).toNumber()] = _222_v57;
        _nw38[(new BigNumber(12)).toNumber()] = _222_v57;
        _nw38[(new BigNumber(13)).toNumber()] = ((!(_137_v1)) ? (_222_v57) : (_222_v57));
        _nw38[(new BigNumber(14)).toNumber()] = _222_v57;
        _nw38[(new BigNumber(15)).toNumber()] = _222_v57;
        _nw38[(new BigNumber(16)).toNumber()] = _222_v57;
        _nw38[(new BigNumber(17)).toNumber()] = _222_v57;
        _227_v59 = _nw38;
        let _index29 = _module.__default.safeIndex(new BigNumber(757), new BigNumber((_227_v59).length));
        (_227_v59)[_index29] = _222_v57;
        let _228_v60;
        let _nw39 = Array((new BigNumber(15)).toNumber()).fill([]);
        _228_v60 = _nw39;
        let _229_v61;
        _229_v61 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_140_v4),_221_v56);
        let _index30 = _module.__default.safeIndex(new BigNumber(757), new BigNumber((_227_v59).length));
        let _rhs23 = (((_229_v61).contains(new BigNumber((_module.__default.fm8(_137_v1, _157_v17, _136_globalState)).length))) ? ((_229_v61).get(new BigNumber((_module.__default.fm8(_137_v1, _157_v17, _136_globalState)).length))) : (((_137_v1) ? (_221_v56) : (_221_v56))));
        let _rhs24 = (((_158_v18).contains(((_137_v1) ? (_137_v1) : (_137_v1)))) ? ((_158_v18).get(((_137_v1) ? (_137_v1) : (_137_v1)))) : (_137_v1));
        let _rhs25 = _222_v57;
        let _rhs26 = _228_v60;
        let _rhs27 = _140_v4;
        let _lhs20 = _136_globalState;
        let _lhs21 = _227_v59;
        let _lhs22 = _module.__default.safeIndex(new BigNumber(757), new BigNumber((_227_v59).length));
        let _lhs23 = _136_globalState;
        _221_v56 = _rhs23;
        _lhs20.f10 = _rhs24;
        _lhs21[_lhs22] = _rhs25;
        _228_v60 = _rhs26;
        _lhs23.f4 = _rhs27;
        let _230_v62;
        let _nw40 = new _module.C1();
        _nw40.__ctor((_214_v49).f11);
        _230_v62 = _nw40;
      } else {
        let _231_v63;
        let _nw41 = Array((new BigNumber(13)).toNumber());
        _nw41[(_dafny.ZERO).toNumber()] = _135_v0;
        _nw41[(_dafny.ONE).toNumber()] = _135_v0;
        _nw41[(new BigNumber(2)).toNumber()] = _135_v0;
        _nw41[(new BigNumber(3)).toNumber()] = _135_v0;
        _nw41[(new BigNumber(4)).toNumber()] = _135_v0;
        _nw41[(new BigNumber(5)).toNumber()] = _135_v0;
        _nw41[(new BigNumber(6)).toNumber()] = _135_v0;
        _nw41[(new BigNumber(7)).toNumber()] = _135_v0;
        _nw41[(new BigNumber(8)).toNumber()] = _135_v0;
        _nw41[(new BigNumber(9)).toNumber()] = _135_v0;
        _nw41[(new BigNumber(10)).toNumber()] = _135_v0;
        _nw41[(new BigNumber(11)).toNumber()] = _135_v0;
        _nw41[(new BigNumber(12)).toNumber()] = _135_v0;
        _231_v63 = _nw41;
        _231_v63 = _231_v63;
        let _232_v64;
        _232_v64 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("jfb"));
        let _233_v65;
        _233_v65 = _151_v12;
        if (_module.__default.fm0(_137_v1, _dafny.Seq.IsProperPrefixOf(_157_v17, (_232_v64)[_module.__default.safeIndex(_140_v4, new BigNumber((_232_v64).length))]), new BigNumber(-756), _module.__default.fm2(_module.__default.fm0(_137_v1, _137_v1, _140_v4, _140_v4, _136_globalState), new BigNumber((_dafny.Seq.of(_233_v65)).length), new BigNumber((_215_v50).length), _157_v17, _136_globalState), _136_globalState)) {
          let _234_v66;
          _234_v66 = new _dafny.CodePoint('w'.codePointAt(0));
          let _235_v67;
          let _nw42 = new _module.C0();
          _nw42.__ctor(_234_v66);
          _235_v67 = _nw42;
          (_136_globalState).f4 = _140_v4;
          let _236_v69;
          let _init5 = ((_237_v4) => function (_238_i11) {
            return (_dafny.Map.Empty.slice().updateUnsafe(_237_v4,_237_v4)).Merge(function () {
              let _coll10 = new _dafny.Map();
              for (const _compr_10 of _dafny.IntegerRange(new BigNumber(-874), new BigNumber(754))) {
                let _239_v68 = _compr_10;
                if (((new BigNumber(-874)).isLessThanOrEqualTo(_239_v68)) && ((_239_v68).isLessThan(new BigNumber(754)))) {
                  _coll10.push([_module.__default.safeModuloInt(_239_v68, _237_v4),_237_v4]);
                }
              }
              return _coll10;
            }());
          })(_140_v4);
          let _nw43 = Array((new BigNumber(20)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw43.length); _i0_5++) {
            _nw43[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _236_v69 = _nw43;
          let _240_v70;
          _240_v70 = _dafny.Map.Empty.slice().updateUnsafe(_140_v4,_140_v4);
          let _index31 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_236_v69).length));
          (_236_v69)[_index31] = _240_v70;
          let _241_v71;
          _241_v71 = _dafny.Seq.of(_214_v49);
          let _242_v72;
          _242_v72 = _dafny.Map.Empty.slice().updateUnsafe(_140_v4,_135_v0);
          let _243_v73;
          _243_v73 = _module.D5.create_DC8(_242_v72);
          let _244_v74;
          _244_v74 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_241_v71),new BigNumber((((_243_v73).dtor_cf11).Merge(_242_v72)).length));
          let _245_v77;
          _245_v77 = _dafny.Map.Empty.slice().updateUnsafe((_235_v67).f12,_137_v1);
          let _246_v78;
          _246_v78 = _dafny.Map.Empty.slice().updateUnsafe(_138_v2,function () {
            let _coll11 = new _dafny.Map();
            for (const _compr_11 of (_245_v77).Keys.Elements) {
              let _247_v76 = _compr_11;
              if ((_245_v77).contains(_247_v76)) {
                _coll11.push([_247_v76,_140_v4]);
              }
            }
            return _coll11;
          }());
          let _248_v79;
          _248_v79 = _dafny.Set.fromElements(new BigNumber((_157_v17).length), _140_v4, _140_v4);
          let _249_v80;
          let _nw44 = Array((new BigNumber(5)).toNumber());
          _249_v80 = _nw44;
          let _250_v81;
          _250_v81 = _dafny.Map.Empty.slice().updateUnsafe(_137_v1,_249_v80);
          let _251_v82;
          _251_v82 = _dafny.Map.Empty.slice().updateUnsafe(_250_v81,_244_v74);
          let _252_v83;
          _252_v83 = _dafny.MultiSet.fromElements(_214_v49);
          let _index32 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_236_v69).length));
          let _rhs28 = _240_v70;
          let _rhs29 = (_dafny.Map.Empty.slice().updateUnsafe(_138_v2,_137_v1)).Merge((function () {
            let _coll12 = new _dafny.Map();
            for (const _compr_12 of (_246_v78).Keys.Elements) {
              let _253_v75 = _compr_12;
              if ((_246_v78).contains(_253_v75)) {
                _coll12.push([_253_v75,_137_v1]);
              }
            }
            return _coll12;
          }()).Merge(_module.__default.fm14(true, _140_v4, _248_v79, _136_globalState)));
          let _rhs30 = (((_251_v82).contains(_250_v81)) ? ((_251_v82).get(_250_v81)) : ((_244_v74).update(_252_v83, _140_v4)));
          let _lhs24 = _236_v69;
          let _lhs25 = _module.__default.safeIndex(new BigNumber(565), new BigNumber((_236_v69).length));
          _lhs24[_lhs25] = _rhs28;
          _139_v3 = _rhs29;
          _244_v74 = _rhs30;
          let _254_v84;
          let _out4;
          _out4 = _module.__default.m1(_140_v4, _135_v0, _136_globalState);
          _254_v84 = _out4;
          let _255_v85;
          _255_v85 = _module.D2.create_DC2((_235_v67).f12);
          let _256_v86;
          _256_v86 = _dafny.MultiSet.fromElements(_255_v85);
          (_136_globalState).f10 = !((_256_v86).IsProperSubsetOf((_256_v86).Intersect((_256_v86).update(_255_v85, _module.__default.abs(_140_v4)))));
        } else {
          let _257_v87;
          _257_v87 = _dafny.Map.Empty.slice().updateUnsafe(((_137_v1) ? (_137_v1) : (_137_v1)),_140_v4);
          let _258_v88;
          _258_v88 = new _dafny.CodePoint('n'.codePointAt(0));
          let _rhs31 = !(_137_v1);
          let _rhs32 = _140_v4;
          let _rhs33 = _dafny.areEqual(_157_v17, _dafny.Seq.update(_157_v17, _module.__default.safeIndex(_140_v4, new BigNumber((_157_v17).length)), _258_v88));
          let _rhs34 = _174_v24;
          let _rhs35 = (_257_v87).Merge(_dafny.Map.Empty.slice().updateUnsafe(_137_v1,_140_v4));
          let _lhs26 = _136_globalState;
          let _lhs27 = _136_globalState;
          let _lhs28 = _136_globalState;
          _lhs26.f10 = _rhs31;
          _lhs27.f4 = _rhs32;
          _lhs28.f10 = _rhs33;
          _174_v24 = _rhs34;
          _257_v87 = _rhs35;
          let _259_v89;
          let _nw45 = Array((new BigNumber(4)).toNumber());
          _259_v89 = _nw45;
          let _index33 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_259_v89).length));
          (_259_v89)[_index33] = ((_137_v1) ? (_214_v49) : (_214_v49));
          let _index34 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_174_v24).length));
          (_174_v24)[_index34] = new BigNumber((_module.__default.fm15(_136_globalState)).length);
          let _260_v90;
          _260_v90 = _dafny.Set.fromElements(_135_v0);
          let _index35 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_259_v89).length));
          let _index36 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_174_v24).length));
          let _rhs36 = _214_v49;
          let _rhs37 = _module.__default.fm0(_137_v1, _137_v1, _140_v4, _140_v4, _136_globalState);
          let _rhs38 = _module.__default.safeModuloInt(_module.__default.fm2(_137_v1, _140_v4, _140_v4, _157_v17, _136_globalState), (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-324)), ((_261_v4) => function (_262_i12) {
            return _261_v4;
          })(_140_v4))).length)).multipliedBy(_140_v4));
          let _rhs39 = _260_v90;
          let _rhs40 = true;
          let _lhs29 = _259_v89;
          let _lhs30 = _module.__default.safeIndex(new BigNumber(264), new BigNumber((_259_v89).length));
          let _lhs31 = _174_v24;
          let _lhs32 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_174_v24).length));
          let _lhs33 = _136_globalState;
          _lhs29[_lhs30] = _rhs36;
          _137_v1 = _rhs37;
          _lhs31[_lhs32] = _rhs38;
          _260_v90 = _rhs39;
          _lhs33.f10 = _rhs40;
          let _263_v91;
          let _out5;
          _out5 = _module.__default.m1((_174_v24)[_module.__default.safeIndex(new BigNumber(183), new BigNumber((_174_v24).length))], _135_v0, _136_globalState);
          _263_v91 = _out5;
          let _264_v92;
          _264_v92 = _dafny.Seq.of(_135_v0, _263_v91, _263_v91, _135_v0);
          let _265_v93;
          let _out6;
          _out6 = _module.__default.m1(new BigNumber((_157_v17).length), (_264_v92)[_module.__default.safeIndex(_140_v4, new BigNumber((_264_v92).length))], _136_globalState);
          _265_v93 = _out6;
          let _index37 = _module.__default.safeIndex(new BigNumber(183), new BigNumber((_174_v24).length));
          (_174_v24)[_index37] = new BigNumber(412);
        }
        _140_v4 = (_153_v14)[_module.__default.safeIndex(_140_v4, new BigNumber((_153_v14).length))];
        let _266_v94;
        _266_v94 = _dafny.MultiSet.fromElements(_140_v4, _140_v4);
        let _267_v95;
        _267_v95 = _dafny.Seq.of(_266_v94);
        let _268_v96;
        _268_v96 = _dafny.Map.Empty.slice().updateUnsafe(_267_v95,!(!(((_137_v1) ? (_137_v1) : (_137_v1)))));
        _268_v96 = (_268_v96).update(_267_v95, _137_v1);
        _140_v4 = _140_v4;
      }
      _140_v4 = _140_v4;
      (_136_globalState).f4 = _140_v4;
      if (_137_v1) {
        let _269_v97;
        let _out7;
        _out7 = _module.__default.m1(_140_v4, _135_v0, _136_globalState);
        _269_v97 = _out7;
        let _index38 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_174_v24).length));
        (_174_v24)[_index38] = _140_v4;
        let _index39 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_174_v24).length));
        (_174_v24)[_index39] = _140_v4;
        (_136_globalState).f10 = _137_v1;
        let _270_v98;
        let _nw46 = new _module.C1();
        _nw46.__ctor((_214_v49).f11);
        _270_v98 = _nw46;
        (_136_globalState).f4 = ((_137_v1) ? (_140_v4) : (_140_v4));
      } else {
        let _271_v99;
        _271_v99 = new _dafny.CodePoint('r'.codePointAt(0));
        let _272_v100;
        _272_v100 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm5(_140_v4, _136_globalState),new BigNumber(692));
        let _273_v101;
        _273_v101 = _dafny.Map.Empty.slice().updateUnsafe(_135_v0,_158_v18);
        let _274_v102;
        _274_v102 = _dafny.Map.Empty.slice().updateUnsafe(((_137_v1) ? (_dafny.Map.Empty.slice().updateUnsafe(_271_v99,new BigNumber(-509))) : (_272_v100)),new BigNumber((_273_v101).length));
        _274_v102 = (_274_v102).update(_272_v100, _140_v4);
        (_136_globalState).f4 = _140_v4;
        (_136_globalState).f3 = _135_v0;
        let _275_v103;
        let _out8;
        _out8 = (_214_v49).m0(_140_v4, _dafny.Seq.Concat(_157_v17, _157_v17), !(_137_v1), _140_v4, _136_globalState);
        _275_v103 = _out8;
        (_136_globalState).f10 = ((_214_v49).f11);
      }
      (_136_globalState).f4 = _module.__default.fm2(_137_v1, _140_v4, _140_v4, _157_v17, _136_globalState);
      process.stdout.write(_dafny.toString((_135_v0)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_136_globalState).f0)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_globalState).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_globalState.f3)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_136_globalState.f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_136_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_136_globalState.f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_136_globalState.f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_137_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v2).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_139_v3).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(true),false).updateUnsafe(_dafny.MultiSet.fromElements(false),false).updateUnsafe(_dafny.MultiSet.fromElements(true, false),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_140_v4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_151_v12).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(123),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_152_v13).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.UnicodeFromString("du")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_153_v14, _dafny.Seq.of(new BigNumber(384), new BigNumber(670)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_155_v15)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_v16).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(384)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_157_v17).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_158_v18).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v24)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v24)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v24)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v24)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v24)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v24)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v24)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v24)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v24)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v24)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v24)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v24)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v24)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v24)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v24)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v24)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v24)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v24)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v24)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v24)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v24)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v24)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v24)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v24)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v24)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v24)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v24)[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v24)[new BigNumber(27)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_174_v24)[new BigNumber(28)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_214_v49).f11)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_215_v50, _dafny.Seq.of(false, false))));
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
    static create_DC1(cf1) {
      let $dt = new D1(0);
      $dt.cf1 = cf1;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get dtor_cf1() { return this.cf1; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1);
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
    static create_DC3(cf3, cf4) {
      let $dt = new D2(0);
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC4(cf5, cf6) {
      let $dt = new D2(1);
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC2(cf2) {
      let $dt = new D2(2);
      $dt.cf2 = cf2;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get is_DC4() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf2() { return this.cf2; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC3" + "(" + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC4" + "(" + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC2" + "(" + _dafny.toString(this.cf2) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf3 === other.cf3 && this.cf4 === other.cf4;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf5, other.cf5) && _dafny.areEqual(this.cf6, other.cf6);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf2, other.cf2);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC3(false, false);
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
    static create_DC5(cf7) {
      let $dt = new D3(0);
      $dt.cf7 = cf7;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get dtor_cf7() { return this.cf7; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC5" + "(" + _dafny.toString(this.cf7) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf7, other.cf7);
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
          return D3.Default();
        }
      };
    }
  }

  $module.D4 = class D4 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC7(cf9, cf10) {
      let $dt = new D4(0);
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC6(cf8) {
      let $dt = new D4(1);
      $dt.cf8 = cf8;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf8() { return this.cf8; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC7" + "(" + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC6" + "(" + _dafny.toString(this.cf8) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf9, other.cf9) && this.cf10 === other.cf10;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf8 === other.cf8;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC7(_dafny.ZERO, false);
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
    static create_DC9(cf12, cf13, cf14, cf15) {
      let $dt = new D5(0);
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC8(cf11) {
      let $dt = new D5(1);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC8() { return this.$tag === 1; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC9" + "(" + this.cf12.toVerbatimString(true) + ", " + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC8" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf12, other.cf12) && this.cf13 === other.cf13 && this.cf14 === other.cf14 && _dafny.areEqual(this.cf15, other.cf15);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf11, other.cf11);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC9(_dafny.Seq.UnicodeFromString(""), false, null, new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC11(cf17, cf18, cf19, cf20, cf21) {
      let $dt = new D6(0);
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC12(cf22) {
      let $dt = new D6(1);
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC13() {
      let $dt = new D6(2);
      return $dt;
    }
    static create_DC10(cf16) {
      let $dt = new D6(3);
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC14(cf23) {
      let $dt = new D6(4);
      $dt.cf23 = cf23;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get is_DC13() { return this.$tag === 2; }
    get is_DC10() { return this.$tag === 3; }
    get is_DC14() { return this.$tag === 4; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf23() { return this.cf23; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC11" + "(" + this.cf17.toVerbatimString(true) + ", " + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC12" + "(" + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC13";
      } else if (this.$tag === 3) {
        return "D6.DC10" + "(" + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 4) {
        return "D6.DC14" + "(" + _dafny.toString(this.cf23) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf17, other.cf17) && this.cf18 === other.cf18 && _dafny.areEqual(this.cf19, other.cf19) && _dafny.areEqual(this.cf20, other.cf20) && _dafny.areEqual(this.cf21, other.cf21);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf22 === other.cf22;
      } else if (this.$tag === 2) {
        return other.$tag === 2;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf16, other.cf16);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf23, other.cf23);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC11(_dafny.Seq.UnicodeFromString(""), null, _dafny.ZERO, _dafny.Seq.of(), _dafny.ZERO);
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
    static create_DC16(cf25, cf26, cf27) {
      let $dt = new D7(0);
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC15(cf24) {
      let $dt = new D7(1);
      $dt.cf24 = cf24;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf24() { return this.cf24; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC16" + "(" + this.cf25.toVerbatimString(true) + ", " + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC15" + "(" + _dafny.toString(this.cf24) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf25, other.cf25) && _dafny.areEqual(this.cf26, other.cf26) && this.cf27 === other.cf27;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf24, other.cf24);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC16(_dafny.Seq.UnicodeFromString(""), _dafny.MultiSet.Empty, false);
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
    static create_DC18(cf29, cf30) {
      let $dt = new D8(0);
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC17(cf28) {
      let $dt = new D8(1);
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC19(cf31) {
      let $dt = new D8(2);
      $dt.cf31 = cf31;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get is_DC17() { return this.$tag === 1; }
    get is_DC19() { return this.$tag === 2; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf31() { return this.cf31; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC18" + "(" + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC17" + "(" + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC19" + "(" + _dafny.toString(this.cf31) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf29, other.cf29) && this.cf30 === other.cf30;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf28, other.cf28);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf31, other.cf31);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC18(_dafny.ZERO, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D8.Default();
        }
      };
    }
  }

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f3 = [];
      this.f4 = _dafny.ZERO;
      this.f9 = false;
      this.f10 = false;
      this._f0 = [];
      this._f1 = _dafny.ZERO;
      this._f2 = false;
      this._f5 = _dafny.ZERO;
      this._f6 = false;
      this._f7 = _dafny.ZERO;
      this._f8 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this).f3 = f3;
      (_this).f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this).f9 = f9;
      (_this).f10 = f10;
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
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f12 = new _dafny.CodePoint('D'.codePointAt(0));
    }
    _parentTraits() {
      return [];
    }
    __ctor(f12) {
      let _this = this;
      (_this)._f12 = f12;
      return;
    }
    fm3(p0, globalState) {
      let _this = this;
      if ((((!(true))) ? (false) : (true))) {
        return false;
      } else {
        return false;
      }
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f11 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f11) {
      let _this = this;
      (_this)._f11 = f11;
      return;
    }
    m0(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let _276_v0;
      _276_v0 = _dafny.Seq.of(new BigNumber((p1).length), p3);
      let _277_i0;
      _277_i0 = _dafny.ZERO;
      L0: {
        while (_dafny.Seq.IsProperPrefixOf(_276_v0, _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-33))))) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_277_i0)) {
              break L0;
            }
            _277_i0 = (_277_i0).plus(_dafny.ONE);
            (globalState).f4 = (_dafny.ZERO).minus(p3);
            let _278_v1;
            _278_v1 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-768),p1);
            _278_v1 = (_278_v1).update(p0, p1);
            (globalState).f10 = _module.__default.fm0(p2, ((p2) ? (p2) : (true)), p3, _module.__default.safeDivisionInt(p3, p3), globalState);
            let _279_v2;
            let _init6 = ((_280_p1) => function (_281_i1) {
              return _280_p1;
            })(p1);
            let _nw47 = Array((new BigNumber(26)).toNumber());
            for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw47.length); _i0_6++) {
              _nw47[_i0_6] = _init6(new BigNumber(_i0_6));
            }
            _279_v2 = _nw47;
            let _index40 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_279_v2).length));
            (_279_v2)[_index40] = p1;
            let _index41 = _module.__default.safeIndex(new BigNumber(105), new BigNumber((_279_v2).length));
            (_279_v2)[_index41] = _dafny.Seq.UnicodeFromString("tqgbnrs");
          }
        }
      }
      (globalState).f10 = p2;
      let _hi1 = p3;
      for (let _282_i2 = p3; _282_i2.isLessThan(_hi1); _282_i2 = _282_i2.plus(_dafny.ONE)) {
        let _283_v3;
        _283_v3 = new _dafny.CodePoint('p'.codePointAt(0));
        let _284_v4;
        _284_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_283_v3,p2)).length),p2);
        let _285_v5;
        _285_v5 = _dafny.Seq.of(p1);
        let _286_v6;
        let _nw48 = Array((new BigNumber(6)).toNumber());
        _nw48[(_dafny.ZERO).toNumber()] = p3;
        _nw48[(_dafny.ONE).toNumber()] = _282_i2;
        _nw48[(new BigNumber(2)).toNumber()] = _282_i2;
        _nw48[(new BigNumber(3)).toNumber()] = _module.__default.fm2(p2, new BigNumber((_dafny.MultiSet.fromElements(p2, p2, p2)).cardinality()), new BigNumber((_284_v4).length), (_285_v5)[_module.__default.safeIndex(new BigNumber((p1).length), new BigNumber((_285_v5).length))], globalState);
        _nw48[(new BigNumber(4)).toNumber()] = p3;
        _nw48[(new BigNumber(5)).toNumber()] = (p0).minus(p3);
        _286_v6 = _nw48;
        let _index42 = _module.__default.safeIndex(new BigNumber(323), new BigNumber((_286_v6).length));
        (_286_v6)[_index42] = _282_i2;
        let _287_v7;
        _287_v7 = _dafny.Seq.of(p2);
        let _index43 = _module.__default.safeIndex(new BigNumber(323), new BigNumber((_286_v6).length));
        (_286_v6)[_index43] = (_dafny.ZERO).minus((new BigNumber((_287_v7).length)).minus(_282_i2));
        if (p2) {
          let _288_v8;
          _288_v8 = _dafny.MultiSet.fromElements(new BigNumber(866), _282_i2);
          (globalState).f4 = ((_286_v6)[_module.__default.safeIndex(new BigNumber(323), new BigNumber((_286_v6).length))]).multipliedBy((_dafny.ZERO).minus(((_dafny.ZERO).minus((((_288_v8).contains(new BigNumber((_dafny.Seq.update(p1, _module.__default.safeIndex(p0, new BigNumber((p1).length)), _283_v3)).length))) ? ((_288_v8).get(new BigNumber((_dafny.Seq.update(p1, _module.__default.safeIndex(p0, new BigNumber((p1).length)), _283_v3)).length))) : (_282_i2)))).multipliedBy((_dafny.ZERO).minus(p3))));
          let _index44 = _module.__default.safeIndex(new BigNumber(323), new BigNumber((_286_v6).length));
          (_286_v6)[_index44] = p3;
          (globalState).f9 = !(p2) || (_module.__default.fm0(p2, false, (_276_v0)[_module.__default.safeIndex((_286_v6)[_module.__default.safeIndex(new BigNumber(323), new BigNumber((_286_v6).length))], new BigNumber((_276_v0).length))], (_dafny.ZERO).minus(new BigNumber((_287_v7).length)), globalState));
          let _index45 = _module.__default.safeIndex(new BigNumber(323), new BigNumber((_286_v6).length));
          (_286_v6)[_index45] = ((((_288_v8).contains(new BigNumber((p1).length))) ? ((_288_v8).get(new BigNumber((p1).length))) : ((_dafny.ZERO).minus(p3)))).multipliedBy((_276_v0)[_module.__default.safeIndex(p3, new BigNumber((_276_v0).length))]);
          let _289_v9;
          _289_v9 = _dafny.Map.Empty.slice().updateUnsafe(p2,_276_v0);
          (globalState).f10 = _dafny.areEqual((((_289_v9).contains(p2)) ? ((_289_v9).get(p2)) : (_276_v0)), _276_v0);
        } else {
          let _290_v10;
          let _nw49 = new _module.C0();
          _nw49.__ctor(_283_v3);
          _290_v10 = _nw49;
          let _291_v11;
          let _nw50 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Map.Empty);
          _291_v11 = _nw50;
          let _index46 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_291_v11).length));
          (_291_v11)[_index46] = _dafny.Map.Empty.slice().updateUnsafe(_282_i2,p2);
          let _292_v12;
          _292_v12 = _dafny.Map.Empty.slice().updateUnsafe(_282_i2,p2);
          let _293_v13;
          _293_v13 = _dafny.Seq.of(_284_v4, ((_292_v12)).Merge(_dafny.Map.Empty.slice().updateUnsafe(p3,p2)), (_284_v4).Merge(_284_v4), _284_v4, _284_v4);
          let _index47 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_291_v11).length));
          (_291_v11)[_index47] = (_293_v13)[_module.__default.safeIndex(((_286_v6)[_module.__default.safeIndex(new BigNumber(323), new BigNumber((_286_v6).length))]).multipliedBy(p3), new BigNumber((_293_v13).length))];
          let _294_v14;
          let _nw51 = new _module.C0();
          _nw51.__ctor((_290_v10).f12);
          _294_v14 = _nw51;
          let _295_v15;
          let _init7 = ((_296_p1) => function (_297_i3) {
            return _296_p1;
          })(p1);
          let _nw52 = Array((new BigNumber(13)).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw52.length); _i0_7++) {
            _nw52[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _295_v15 = _nw52;
          let _index48 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_295_v15).length));
          (_295_v15)[_index48] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-706)), ((_298_p3, _299_v3) => function (_300_i4) {
            return (_module.D2.create_DC4(_298_p3, _299_v3)).dtor_cf6;
          })(p3, _283_v3));
          let _301_v16;
          _301_v16 = _dafny.Set.fromElements(p3);
          let _302_v17;
          _302_v17 = _dafny.Seq.of(_286_v6);
          let _303_v18;
          _303_v18 = _dafny.Map.Empty.slice().updateUnsafe((_302_v17)[_module.__default.safeIndex(p0, new BigNumber((_302_v17).length))],new BigNumber((_276_v0).length));
          let _304_v19;
          _304_v19 = _dafny.MultiSet.fromElements(p2);
          let _305_v20;
          _305_v20 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(615),_282_i2);
          let _index49 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_295_v15).length));
          let _rhs41 = p1;
          let _rhs42 = (_301_v16).IsSubsetOf(_301_v16);
          let _rhs43 = ((!(_303_v18).equals(_303_v18)) ? ((((_304_v19).contains((_287_v7)[_module.__default.safeIndex(new BigNumber((_276_v0).length), new BigNumber((_287_v7).length))])) ? ((_304_v19).get((_287_v7)[_module.__default.safeIndex(new BigNumber((_276_v0).length), new BigNumber((_287_v7).length))])) : (_282_i2))) : ((((_305_v20).contains(p0)) ? ((_305_v20).get(p0)) : (p3))));
          let _rhs44 = !(p2);
          let _lhs34 = _295_v15;
          let _lhs35 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_295_v15).length));
          let _lhs36 = globalState;
          let _lhs37 = globalState;
          let _lhs38 = globalState;
          _lhs34[_lhs35] = _rhs41;
          _lhs36.f9 = _rhs42;
          _lhs37.f4 = _rhs43;
          _lhs38.f9 = _rhs44;
          let _306_v21;
          _306_v21 = _dafny.MultiSet.fromElements(_290_v10, _290_v10);
          let _index50 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_295_v15).length));
          (_295_v15)[_index50] = _module.__default.fm4((_dafny.ZERO).minus(new BigNumber((_306_v21).cardinality())), p2, !(p2), ((_286_v6)[_module.__default.safeIndex(new BigNumber(323), new BigNumber((_286_v6).length))]).multipliedBy((_286_v6)[_module.__default.safeIndex(new BigNumber(323), new BigNumber((_286_v6).length))]), globalState);
        }
        let _307_v22;
        _307_v22 = _284_v4;
        let _308_v23;
        _308_v23 = _dafny.Seq.of(_307_v22, _307_v22);
        (globalState).f10 = _dafny.Seq.IsPrefixOf(((p2) ? (_308_v23) : (_308_v23)), _308_v23);
        (globalState).f4 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt(_module.__default.fm2(p2, (_286_v6)[_module.__default.safeIndex(new BigNumber(323), new BigNumber((_286_v6).length))], new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(199)), ((_309_v3) => function (_310_i5) {
          return _309_v3;
        })(_283_v3))).length), p1, globalState), _282_i2), (_286_v6)[_module.__default.safeIndex(new BigNumber(323), new BigNumber((_286_v6).length))]);
      }
      if (p2) {
        let _311_v24;
        let _nw53 = new _module.C0();
        _nw53.__ctor(_module.__default.fm5((_dafny.ZERO).minus(new BigNumber((_module.__default.fm4(p3, p2, p2, new BigNumber((_dafny.Seq.UnicodeFromString("aroliv")).length), globalState)).length)), globalState));
        _311_v24 = _nw53;
        let _312_v25;
        _312_v25 = _dafny.Seq.of(_276_v0);
        _312_v25 = _312_v25;
        (globalState).f10 = !((false) === ((p0).isLessThan(new BigNumber(636))));
        let _313_v26;
        let _nw54 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
        _313_v26 = _nw54;
        let _index51 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_313_v26).length));
        (_313_v26)[_index51] = p0;
        let _314_v27;
        _314_v27 = _dafny.Set.fromElements(p2, p2, p2, !(p2));
        let _315_v28;
        let _nw55 = Array((new BigNumber(15)).toNumber()).fill(false);
        _315_v28 = _nw55;
        let _316_v29;
        _316_v29 = _dafny.Seq.of(p2);
        let _317_v30;
        _317_v30 = _module.D2.create_DC3(p2, p2);
        let _index52 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_313_v26).length));
        let _rhs45 = p3;
        let _rhs46 = (_314_v27).IsSubsetOf(_314_v27);
        let _rhs47 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(p0, p0));
        let _rhs48 = _module.__default.safeDivisionInt(p0, p0);
        let _rhs49 = ((((_317_v30).dtor_cf4) ? (p3) : (p0))).isLessThan(new BigNumber((_module.__default.fm4(new BigNumber((_dafny.MultiSet.fromElements(_315_v28)).cardinality()), p2, (_316_v29)[_module.__default.safeIndex(p0, new BigNumber((_316_v29).length))], p0, globalState)).length));
        let _lhs39 = _313_v26;
        let _lhs40 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_313_v26).length));
        let _lhs41 = globalState;
        let _lhs42 = globalState;
        let _lhs43 = globalState;
        let _lhs44 = globalState;
        _lhs39[_lhs40] = _rhs45;
        _lhs41.f9 = _rhs46;
        _lhs42.f4 = _rhs47;
        _lhs43.f4 = _rhs48;
        _lhs44.f9 = _rhs49;
        let _318_v31;
        _318_v31 = _dafny.MultiSet.fromElements(p2, p2, p2, p2, !(p2));
        let _319_v32;
        _319_v32 = _dafny.Map.Empty.slice().updateUnsafe(_318_v31,_311_v24);
        let _320_v33;
        _320_v33 = _319_v32;
        let _321_v34;
        _321_v34 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_320_v33));
        let _322_v35;
        let _nw56 = Array((new BigNumber(8)).toNumber());
        _nw56[(_dafny.ZERO).toNumber()] = (((_321_v34).contains(p2)) ? ((_321_v34).get(p2)) : (_319_v32));
        _nw56[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm6(true, (_313_v26)[_module.__default.safeIndex(new BigNumber(933), new BigNumber((_313_v26).length))], globalState),_311_v24);
        _nw56[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_318_v31,_311_v24);
        _nw56[(new BigNumber(3)).toNumber()] = _319_v32;
        _nw56[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_318_v31,_311_v24);
        _nw56[(new BigNumber(5)).toNumber()] = _319_v32;
        _nw56[(new BigNumber(6)).toNumber()] = _319_v32;
        _nw56[(new BigNumber(7)).toNumber()] = _319_v32;
        _322_v35 = _nw56;
        let _index53 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_322_v35).length));
        (_322_v35)[_index53] = _319_v32;
        let _index54 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_322_v35).length));
        (_322_v35)[_index54] = (_320_v33);
      } else {
        (globalState).f10 = p2;
        let _323_v36;
        _323_v36 = _module.D2.create_DC3(p2, p2);
        let _source5 = _323_v36;
        if (_source5.is_DC3) {
          let _324___mcc_h0 = (_source5).cf3;
          let _325___mcc_h1 = (_source5).cf4;
          let _326_cf4 = _325___mcc_h1;
          let _327_cf3 = _324___mcc_h0;
          let _328_v37;
          _328_v37 = new _dafny.CodePoint('b'.codePointAt(0));
          let _329_v38;
          let _init8 = ((_330_p2) => function (_331_i6) {
            return _330_p2;
          })(p2);
          let _nw57 = Array((new BigNumber(22)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw57.length); _i0_8++) {
            _nw57[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _329_v38 = _nw57;
          let _332_v39;
          _332_v39 = _dafny.Map.Empty.slice().updateUnsafe(_328_v37,_329_v38);
          _332_v39 = _dafny.Map.Empty.slice().updateUnsafe(((p2) ? (_module.__default.fm5(new BigNumber((_276_v0).length), globalState)) : (_328_v37)),_329_v38);
          let _333_v40;
          _333_v40 = _dafny.Seq.UnicodeFromString("ovmhyj");
          let _index55 = _module.__default.safeIndex(new BigNumber(415), new BigNumber((_329_v38).length));
          (_329_v38)[_index55] = _327_cf3;
          let _334_v41;
          _334_v41 = _dafny.MultiSet.fromElements(p2);
          let _index56 = _module.__default.safeIndex(new BigNumber(415), new BigNumber((_329_v38).length));
          let _rhs50 = _326_cf4;
          let _rhs51 = _dafny.Seq.Concat(p1, _dafny.Seq.Concat(_333_v40, p1));
          let _rhs52 = ((_334_v41).contains(false)) === (p2);
          let _lhs45 = globalState;
          let _lhs46 = _329_v38;
          let _lhs47 = _module.__default.safeIndex(new BigNumber(415), new BigNumber((_329_v38).length));
          _lhs45.f10 = _rhs50;
          _333_v40 = _rhs51;
          _lhs46[_lhs47] = _rhs52;
          let _335_v42;
          let _nw58 = Array((new BigNumber(14)).toNumber());
          _nw58[(_dafny.ZERO).toNumber()] = _module.__default.fm2(!(p2), (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(p0, p0)).length)), p0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-331)), ((_336_v37) => function (_337_i7) {
            return _336_v37;
          })(_328_v37)), globalState);
          _nw58[(_dafny.ONE).toNumber()] = new BigNumber(-31);
          _nw58[(new BigNumber(2)).toNumber()] = p3;
          _nw58[(new BigNumber(3)).toNumber()] = p3;
          _nw58[(new BigNumber(4)).toNumber()] = p3;
          _nw58[(new BigNumber(5)).toNumber()] = p3;
          _nw58[(new BigNumber(6)).toNumber()] = p3;
          _nw58[(new BigNumber(7)).toNumber()] = p0;
          _nw58[(new BigNumber(8)).toNumber()] = p3;
          _nw58[(new BigNumber(9)).toNumber()] = p0;
          _nw58[(new BigNumber(10)).toNumber()] = p0;
          _nw58[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(p0);
          _nw58[(new BigNumber(12)).toNumber()] = p3;
          _nw58[(new BigNumber(13)).toNumber()] = p3;
          _335_v42 = _nw58;
          let _338_v43;
          _338_v43 = _dafny.Map.Empty.slice().updateUnsafe(p2,_335_v42);
          let _339_v44;
          _339_v44 = _dafny.Seq.of(_327_cf3);
          let _340_v45;
          _340_v45 = _dafny.Map.Empty.slice().updateUnsafe(_328_v37,new BigNumber((_dafny.Seq.UnicodeFromString("w")).length));
          let _341_v46;
          _341_v46 = _dafny.MultiSet.fromElements(new BigNumber(43));
          let _342_v47;
          _342_v47 = _dafny.Seq.of(false, _module.__default.fm0(_module.__default.fm0(_327_cf3, (_339_v44)[_module.__default.safeIndex(p0, new BigNumber((_339_v44).length))], new BigNumber(808), p0, globalState), p2, (((_340_v45).contains(new _dafny.CodePoint('r'.codePointAt(0)))) ? ((_340_v45).get(new _dafny.CodePoint('r'.codePointAt(0)))) : (new BigNumber(550))), new BigNumber((_341_v46).cardinality()), globalState), _326_cf4, _327_cf3, p2);
          _338_v43 = (_338_v43).update((_339_v44)[_module.__default.safeIndex(new BigNumber(-217), new BigNumber((_339_v44).length))], _335_v42);
          (globalState).f9 = _327_cf3;
        } else if (_source5.is_DC4) {
          let _343___mcc_h2 = (_source5).cf5;
          let _344___mcc_h3 = (_source5).cf6;
          let _345_cf6 = _344___mcc_h3;
          let _346_cf5 = _343___mcc_h2;
          let _347_v48;
          _347_v48 = _module.D2.create_DC4(_346_cf5, _345_cf6);
          let _348_v49;
          _348_v49 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let2_0) {
            return function (_349_dt__update__tmp_h0) {
              return function (_pat_let3_0) {
                return function (_350_dt__update_hcf6_h0) {
                  return _module.D2.create_DC4((_349_dt__update__tmp_h0).dtor_cf5, _350_dt__update_hcf6_h0);
                }(_pat_let3_0);
              }(new _dafny.CodePoint('d'.codePointAt(0)));
            }(_pat_let2_0);
          }(_347_v48),true);
          _348_v49 = (_348_v49).update(_347_v48, _dafny.Seq.IsProperPrefixOf(p1, p1));
          let _351_v50;
          let _nw59 = new _module.C0();
          _nw59.__ctor(_345_cf6);
          _351_v50 = _nw59;
          let _352_v51;
          _352_v51 = _dafny.MultiSet.fromElements(_351_v50, _351_v50, _351_v50, _351_v50);
          _352_v51 = (_352_v51).update(_351_v50, _module.__default.abs(_346_cf5));
          _345_cf6 = (_351_v50).f12;
          let _353_v52;
          _353_v52 = _dafny.Set.fromElements(!(p2), p2, p2, _module.__default.fm0(true, p2, new BigNumber(585), _346_cf5, globalState), true);
          let _354_v53;
          _354_v53 = _dafny.MultiSet.fromElements(p3);
          let _355_v54;
          _355_v54 = _dafny.Map.Empty.slice().updateUnsafe(_351_v50,p3);
          let _356_v55;
          _356_v55 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(p2),_351_v50);
          let _357_v56;
          _357_v56 = _356_v55;
          let _358_v57;
          _358_v57 = _dafny.Map.Empty.slice().updateUnsafe((_323_v36).dtor_cf3,_357_v56);
          let _359_v58;
          _359_v58 = _dafny.Seq.of(p2);
          let _360_v59;
          let _nw60 = Array((new BigNumber(17)).toNumber());
          _nw60[(_dafny.ZERO).toNumber()] = new BigNumber((_276_v0).length);
          _nw60[(_dafny.ONE).toNumber()] = p0;
          _nw60[(new BigNumber(2)).toNumber()] = p3;
          _nw60[(new BigNumber(3)).toNumber()] = p0;
          _nw60[(new BigNumber(4)).toNumber()] = new BigNumber((_353_v52).length);
          _nw60[(new BigNumber(5)).toNumber()] = new BigNumber(608);
          _nw60[(new BigNumber(6)).toNumber()] = new BigNumber((_354_v53).cardinality());
          _nw60[(new BigNumber(7)).toNumber()] = new BigNumber((p1).length);
          _nw60[(new BigNumber(8)).toNumber()] = (((_354_v53).contains(_346_cf5)) ? ((_354_v53).get(_346_cf5)) : ((_dafny.ZERO).minus(_module.__default.fm2(true, p0, p0, p1, globalState))));
          _nw60[(new BigNumber(9)).toNumber()] = _346_cf5;
          _nw60[(new BigNumber(10)).toNumber()] = new BigNumber((_355_v54).length);
          _nw60[(new BigNumber(11)).toNumber()] = new BigNumber(((_358_v57).update(p2, _357_v56)).length);
          _nw60[(new BigNumber(12)).toNumber()] = new BigNumber((_359_v58).length);
          _nw60[(new BigNumber(13)).toNumber()] = p0;
          _nw60[(new BigNumber(14)).toNumber()] = p3;
          _nw60[(new BigNumber(15)).toNumber()] = _346_cf5;
          _nw60[(new BigNumber(16)).toNumber()] = _module.__default.fm2(p2, p0, _346_cf5, p1, globalState);
          _360_v59 = _nw60;
          let _361_v60;
          let _nw61 = Array((new BigNumber(7)).toNumber());
          _nw61[(_dafny.ZERO).toNumber()] = new BigNumber((p1).length);
          _nw61[(_dafny.ONE).toNumber()] = _346_cf5;
          _nw61[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_351_v50),_360_v59)).length);
          _nw61[(new BigNumber(3)).toNumber()] = p3;
          _nw61[(new BigNumber(4)).toNumber()] = new BigNumber(671);
          _nw61[(new BigNumber(5)).toNumber()] = _346_cf5;
          _nw61[(new BigNumber(6)).toNumber()] = p3;
          _361_v60 = _nw61;
          let _index57 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_361_v60).length));
          (_361_v60)[_index57] = (new BigNumber(-38)).plus(p3);
          let _index58 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_361_v60).length));
          (_361_v60)[_index58] = p0;
        } else {
          let _362___mcc_h4 = (_source5).cf2;
          let _363_cf2 = _362___mcc_h4;
          let _364_v61;
          _364_v61 = _dafny.Set.fromElements(p3);
          let _365_v62;
          _365_v62 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,p3);
          let _366_v63;
          _366_v63 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
          let _367_v64;
          let _nw62 = new _module.C0();
          _nw62.__ctor(_363_cf2);
          _367_v64 = _nw62;
          let _368_v65;
          _368_v65 = _dafny.Map.Empty.slice().updateUnsafe(_367_v64,p2);
          let _369_v66;
          _369_v66 = _dafny.MultiSet.fromElements(p0, p3);
          let _370_v67;
          _370_v67 = _dafny.MultiSet.fromElements(p2, p2);
          let _371_v68;
          let _nw63 = Array((new BigNumber(21)).toNumber());
          _nw63[(_dafny.ZERO).toNumber()] = _module.__default.fm2(p2, p3, new BigNumber((p1).length), p1, globalState);
          _nw63[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.fm2(p2, p3, p0, p1, globalState)));
          _nw63[(new BigNumber(2)).toNumber()] = ((p2) ? ((_dafny.ZERO).minus(p0)) : (p3));
          _nw63[(new BigNumber(3)).toNumber()] = p0;
          _nw63[(new BigNumber(4)).toNumber()] = ((!(p2)) ? (new BigNumber((_364_v61).length)) : (p0));
          _nw63[(new BigNumber(5)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(807), p0);
          _nw63[(new BigNumber(6)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((p1).length), p3);
          _nw63[(new BigNumber(7)).toNumber()] = p3;
          _nw63[(new BigNumber(8)).toNumber()] = p3;
          _nw63[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(750)), ((_372_p0) => function (_373_i8) {
            return _372_p0;
          })(p0)))).cardinality());
          _nw63[(new BigNumber(10)).toNumber()] = p3;
          _nw63[(new BigNumber(11)).toNumber()] = p3;
          _nw63[(new BigNumber(12)).toNumber()] = ((p2) ? (p0) : ((((_365_v62).contains((_this).f11)) ? ((_365_v62).get((_this).f11)) : (p3))));
          _nw63[(new BigNumber(13)).toNumber()] = p0;
          _nw63[(new BigNumber(14)).toNumber()] = new BigNumber((p1).length);
          _nw63[(new BigNumber(15)).toNumber()] = new BigNumber(-629);
          _nw63[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber(-111), new BigNumber((_366_v63).length)));
          _nw63[(new BigNumber(17)).toNumber()] = new BigNumber((_368_v65).length);
          _nw63[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_369_v66).cardinality()));
          _nw63[(new BigNumber(19)).toNumber()] = p3;
          _nw63[(new BigNumber(20)).toNumber()] = ((p2) ? (p3) : (new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements((((_370_v67).contains(p2)) ? ((_370_v67).get(p2)) : (p0)))).cardinality()))).length)));
          _371_v68 = _nw63;
          let _374_v69;
          let _init9 = ((_375_v64) => function (_376_i9) {
            return (_375_v64).f12;
          })(_367_v64);
          let _nw64 = Array((new BigNumber(21)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw64.length); _i0_9++) {
            _nw64[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _374_v69 = _nw64;
          let _index59 = _module.__default.safeIndex(new BigNumber(650), new BigNumber((_374_v69).length));
          (_374_v69)[_index59] = _363_cf2;
          let _index60 = _module.__default.safeIndex(new BigNumber(650), new BigNumber((_374_v69).length));
          let _rhs53 = _371_v68;
          let _rhs54 = _363_cf2;
          let _rhs55 = p2;
          let _rhs56 = p2;
          let _lhs48 = _374_v69;
          let _lhs49 = _module.__default.safeIndex(new BigNumber(650), new BigNumber((_374_v69).length));
          let _lhs50 = globalState;
          let _lhs51 = globalState;
          _371_v68 = _rhs53;
          _lhs48[_lhs49] = _rhs54;
          _lhs50.f10 = _rhs55;
          _lhs51.f9 = _rhs56;
          let _nw65 = new _module.C0();
          _nw65.__ctor((_367_v64).f12);
          _367_v64 = _nw65;
          let _377_v70;
          _377_v70 = _dafny.Map.Empty.slice().updateUnsafe(p3,_module.__default.fm0(p2, p2, p3, p3, globalState));
          _377_v70 = function () {
            let _coll13 = new _dafny.Map();
            for (const _compr_13 of _dafny.IntegerRange(new BigNumber(531), new BigNumber(273))) {
              let _378_v71 = _compr_13;
              if (((new BigNumber(531)).isLessThanOrEqualTo(_378_v71)) && ((_378_v71).isLessThan(new BigNumber(273)))) {
                _coll13.push([(_378_v71).minus(new BigNumber((_276_v0).length)),!(_dafny.areEqual(_dafny.Seq.of(p2), _dafny.Seq.of(p2, p2)))]);
              }
            }
            return _coll13;
          }();
          let _379_v72;
          let _nw66 = new _module.C0();
          _nw66.__ctor((_367_v64).f12);
          _379_v72 = _nw66;
        }
        let _380_v73;
        _380_v73 = _dafny.Map.Empty.slice().updateUnsafe(p2,_dafny.Seq.UnicodeFromString("nqmai"));
        let _381_v74;
        _381_v74 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
        let _382_v75;
        _382_v75 = _dafny.Set.fromElements(new BigNumber((_381_v74).length));
        let _383_v76;
        _383_v76 = new _dafny.CodePoint('o'.codePointAt(0));
        let _384_v77;
        _384_v77 = _dafny.MultiSet.fromElements(p2);
        let _385_v78;
        _385_v78 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("twnjp"), p1);
        let _386_v79;
        let _nw67 = Array((new BigNumber(28)).toNumber());
        _nw67[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Set.fromElements(p3)).length);
        _nw67[(_dafny.ONE).toNumber()] = p3;
        _nw67[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(p3, p3));
        _nw67[(new BigNumber(3)).toNumber()] = p0;
        _nw67[(new BigNumber(4)).toNumber()] = ((p2) ? (p0) : ((_dafny.ZERO).minus(p3)));
        _nw67[(new BigNumber(5)).toNumber()] = (new BigNumber((_module.__default.fm7(p2, p0, p2, new BigNumber((_380_v73).length), globalState)).length)).minus((_dafny.ZERO).minus(p3));
        _nw67[(new BigNumber(6)).toNumber()] = p0;
        _nw67[(new BigNumber(7)).toNumber()] = p3;
        _nw67[(new BigNumber(8)).toNumber()] = p3;
        _nw67[(new BigNumber(9)).toNumber()] = p3;
        _nw67[(new BigNumber(10)).toNumber()] = new BigNumber((_382_v75).length);
        _nw67[(new BigNumber(11)).toNumber()] = p0;
        _nw67[(new BigNumber(12)).toNumber()] = _module.__default.safeDivisionInt(p0, new BigNumber((_381_v74).length));
        _nw67[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus(p0);
        _nw67[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("m"), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.UnicodeFromString("m")).length)), _383_v76)).length);
        _nw67[(new BigNumber(15)).toNumber()] = (_276_v0)[_module.__default.safeIndex(new BigNumber(250), new BigNumber((_276_v0).length))];
        _nw67[(new BigNumber(16)).toNumber()] = p3;
        _nw67[(new BigNumber(17)).toNumber()] = new BigNumber((_module.__default.fm8(p2, p1, globalState)).length);
        _nw67[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus(p3);
        _nw67[(new BigNumber(19)).toNumber()] = p3;
        _nw67[(new BigNumber(20)).toNumber()] = (new BigNumber(-652)).minus(p0);
        _nw67[(new BigNumber(21)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("xlpfc")).length);
        _nw67[(new BigNumber(22)).toNumber()] = new BigNumber(739);
        _nw67[(new BigNumber(23)).toNumber()] = new BigNumber((((_dafny.MultiSet.fromElements(p2)).update(p2, _module.__default.abs(p3))).Union(_384_v77)).cardinality());
        _nw67[(new BigNumber(24)).toNumber()] = p3;
        _nw67[(new BigNumber(25)).toNumber()] = p3;
        _nw67[(new BigNumber(26)).toNumber()] = new BigNumber(655);
        _nw67[(new BigNumber(27)).toNumber()] = new BigNumber((_dafny.Seq.Concat(p1, (_385_v78)[_module.__default.safeIndex(p0, new BigNumber((_385_v78).length))])).length);
        _386_v79 = _nw67;
        let _387_v80;
        _387_v80 = _dafny.MultiSet.fromElements(new BigNumber((p1).length), p3);
        let _388_v81;
        let _nw68 = new _module.C0();
        _nw68.__ctor(new _dafny.CodePoint('j'.codePointAt(0)));
        _388_v81 = _nw68;
        let _389_v82;
        _389_v82 = _dafny.Map.Empty.slice().updateUnsafe(_384_v77,_388_v81);
        let _390_v83;
        _390_v83 = _389_v82;
        let _391_v84;
        _391_v84 = _dafny.Set.fromElements(_390_v83);
        let _index61 = _module.__default.safeIndex(new BigNumber(566), new BigNumber((_386_v79).length));
        (_386_v79)[_index61] = (((_387_v80).contains(p0)) ? ((_387_v80).get(p0)) : ((_dafny.ZERO).minus(new BigNumber((_391_v84).length))));
        let _392_v85;
        _392_v85 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        let _393_v86;
        _393_v86 = _dafny.Map.Empty.slice().updateUnsafe(p3,(((_392_v85).contains(p0)) ? ((_392_v85).get(p0)) : (p3)));
        let _394_v87;
        _394_v87 = _dafny.Map.Empty.slice().updateUnsafe(p3,_dafny.Set.fromElements(_module.D2.create_DC3(p2, p2), _323_v36, _323_v36, _323_v36));
        let _index62 = _module.__default.safeIndex(new BigNumber(566), new BigNumber((_386_v79).length));
        let _rhs57 = new BigNumber((_module.__default.fm1(_392_v85, !(p2), new BigNumber((p1).length), (_dafny.ZERO).minus(p0), globalState)).length);
        let _rhs58 = (p1)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((p1).length))];
        let _rhs59 = !(_module.__default.fm9(p2, globalState)).equals(_394_v87);
        let _lhs52 = _386_v79;
        let _lhs53 = _module.__default.safeIndex(new BigNumber(566), new BigNumber((_386_v79).length));
        let _lhs54 = globalState;
        _lhs52[_lhs53] = _rhs57;
        _383_v76 = _rhs58;
        _lhs54.f10 = _rhs59;
        (globalState).f9 = p2;
        let _395_v88;
        _395_v88 = _dafny.Map.Empty.slice().updateUnsafe(p2,p0);
        _395_v88 = (_395_v88).update(p2, (new BigNumber(-902)).plus(new BigNumber((_395_v88).length)));
      }
      let _396_v89;
      _396_v89 = _dafny.MultiSet.fromElements(p2);
      let _397_v90;
      _397_v90 = new _dafny.CodePoint('y'.codePointAt(0));
      let _398_v91;
      let _nw69 = new _module.C0();
      _nw69.__ctor(_397_v90);
      _398_v91 = _nw69;
      let _399_v92;
      _399_v92 = _dafny.Map.Empty.slice().updateUnsafe(_396_v89,_398_v91);
      let _400_v93;
      _400_v93 = _399_v92;
      let _401_v94;
      _401_v94 = _dafny.Seq.of(_400_v93);
      let _hi2 = _module.__default.fm2(p2, (_dafny.ZERO).minus(new BigNumber(-285)), p0, p1, globalState);
      for (let _402_i10 = _module.__default.safeDivisionInt(new BigNumber((_dafny.MultiSet.FromArray(_401_v94)).cardinality()), p0); _402_i10.isLessThan(_hi2); _402_i10 = _402_i10.plus(_dafny.ONE)) {
        let _403_v95;
        _403_v95 = _dafny.Map.Empty.slice().updateUnsafe(p3,p2);
        let _404_v96;
        _404_v96 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_403_v95).length));
        let _405_v97;
        _405_v97 = _dafny.Map.Empty.slice().updateUnsafe((((_404_v96).contains(p0)) ? ((_404_v96).get(p0)) : (p3)),p2);
        let _406_v98;
        _406_v98 = _dafny.Seq.of((p2) === (!((((_405_v97).contains(p3)) ? ((_405_v97).get(p3)) : (p2)))));
        (globalState).f10 = (_406_v98)[_module.__default.safeIndex(_402_i10, new BigNumber((_406_v98).length))];
        let _407_v99;
        let _nw70 = Array((new BigNumber(21)).toNumber());
        _nw70[(_dafny.ZERO).toNumber()] = _400_v93;
        _nw70[(_dafny.ONE).toNumber()] = ((p2) ? (_400_v93) : (_400_v93));
        _nw70[(new BigNumber(2)).toNumber()] = _400_v93;
        _nw70[(new BigNumber(3)).toNumber()] = _400_v93;
        _nw70[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_396_v89,_398_v91);
        _nw70[(new BigNumber(5)).toNumber()] = _400_v93;
        _nw70[(new BigNumber(6)).toNumber()] = _400_v93;
        _nw70[(new BigNumber(7)).toNumber()] = _400_v93;
        _nw70[(new BigNumber(8)).toNumber()] = _400_v93;
        _nw70[(new BigNumber(9)).toNumber()] = _400_v93;
        _nw70[(new BigNumber(10)).toNumber()] = _400_v93;
        _nw70[(new BigNumber(11)).toNumber()] = _400_v93;
        _nw70[(new BigNumber(12)).toNumber()] = _400_v93;
        _nw70[(new BigNumber(13)).toNumber()] = _399_v92;
        _nw70[(new BigNumber(14)).toNumber()] = _399_v92;
        _nw70[(new BigNumber(15)).toNumber()] = _400_v93;
        _nw70[(new BigNumber(16)).toNumber()] = _400_v93;
        _nw70[(new BigNumber(17)).toNumber()] = _399_v92;
        _nw70[(new BigNumber(18)).toNumber()] = _399_v92;
        _nw70[(new BigNumber(19)).toNumber()] = _399_v92;
        _nw70[(new BigNumber(20)).toNumber()] = (_399_v92);
        _407_v99 = _nw70;
        let _index63 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_407_v99).length));
        (_407_v99)[_index63] = _400_v93;
        let _index64 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_407_v99).length));
        (_407_v99)[_index64] = _399_v92;
        let _408_v100;
        _408_v100 = _dafny.Map.Empty.slice().updateUnsafe(p2,_402_i10);
        let _409_v101;
        let _nw71 = Array((new BigNumber(11)).toNumber()).fill(false);
        _409_v101 = _nw71;
        let _410_v102;
        _410_v102 = _dafny.Map.Empty.slice().updateUnsafe(_408_v100,_409_v101);
        _410_v102 = (_410_v102).update(_408_v100, _409_v101);
        let _411_v103;
        let _nw72 = new _module.C0();
        _nw72.__ctor(_397_v90);
        _411_v103 = _nw72;
      }
      let _412_v104;
      _412_v104 = _dafny.Seq.of(p2);
      let _413_v105;
      let _nw73 = Array((new BigNumber(26)).toNumber());
      _nw73[(_dafny.ZERO).toNumber()] = p2;
      _nw73[(_dafny.ONE).toNumber()] = p2;
      _nw73[(new BigNumber(2)).toNumber()] = (_this).f11;
      _nw73[(new BigNumber(3)).toNumber()] = (_this).f11;
      _nw73[(new BigNumber(4)).toNumber()] = (_this).f11;
      _nw73[(new BigNumber(5)).toNumber()] = (_this).f11;
      _nw73[(new BigNumber(6)).toNumber()] = (_this).f11;
      _nw73[(new BigNumber(7)).toNumber()] = (_this).f11;
      _nw73[(new BigNumber(8)).toNumber()] = (_this).f11;
      _nw73[(new BigNumber(9)).toNumber()] = true;
      _nw73[(new BigNumber(10)).toNumber()] = (_this).f11;
      _nw73[(new BigNumber(11)).toNumber()] = (_this).f11;
      _nw73[(new BigNumber(12)).toNumber()] = (_this).f11;
      _nw73[(new BigNumber(13)).toNumber()] = p2;
      _nw73[(new BigNumber(14)).toNumber()] = (_this).f11;
      _nw73[(new BigNumber(15)).toNumber()] = (_this).f11;
      _nw73[(new BigNumber(16)).toNumber()] = (_this).f11;
      _nw73[(new BigNumber(17)).toNumber()] = (_this).f11;
      _nw73[(new BigNumber(18)).toNumber()] = (_this).f11;
      _nw73[(new BigNumber(19)).toNumber()] = (_this).f11;
      _nw73[(new BigNumber(20)).toNumber()] = (_412_v104)[_module.__default.safeIndex(p3, new BigNumber((_412_v104).length))];
      _nw73[(new BigNumber(21)).toNumber()] = (_this).f11;
      _nw73[(new BigNumber(22)).toNumber()] = (_this).f11;
      _nw73[(new BigNumber(23)).toNumber()] = (_this).f11;
      _nw73[(new BigNumber(24)).toNumber()] = p2;
      _nw73[(new BigNumber(25)).toNumber()] = (_this).f11;
      _413_v105 = _nw73;
      let _index65 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_413_v105).length));
      (_413_v105)[_index65] = (_this).f11;
      let _index66 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_413_v105).length));
      (_413_v105)[_index66] = (_this).f11;
      r0 = _module.__default.fm10(globalState);
      return r0;
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
