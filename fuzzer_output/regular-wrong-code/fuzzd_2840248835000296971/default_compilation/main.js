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
      return new BigNumber((((false) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(62)), function (_0_i0) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      })) : (_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(584)), function (_1_i1) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("jkumqri"))))).length);
    };
    static fm1(p0, p1, p2, p3, globalState) {
      return !((((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-940)))).plus(new BigNumber(910))).isEqualTo(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("rrv"), _dafny.Seq.UnicodeFromString("ffpcdlj"))).length)));
    };
    static fm2(p0, p1, globalState) {
      return _module.D0.create_DC0(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-974)), function (_2_i0) {
  return new BigNumber(561);
}));
    };
    static fm3(p0, p1, globalState) {
      return function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(224), new BigNumber(957))) {
          let _3_v0 = _compr_0;
          if (((new BigNumber(224)).isLessThanOrEqualTo(_3_v0)) && ((_3_v0).isLessThan(new BigNumber(957)))) {
            _coll0.add(_module.__default.safeModuloInt(_3_v0, new BigNumber(-930)));
          }
        }
        return _coll0;
      }();
    };
    static fm4(p0, globalState) {
      if ((!(false)) || (true)) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('s'.codePointAt(0));
      }
    };
    static fm5(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(true, true), _dafny.Seq.of(false)), _dafny.Seq.Concat(_dafny.Seq.of((_module.D2.create_DC7(true, new BigNumber(-582), true, _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(!(false), true),new BigNumber(370)))).dtor_cf6, false), _dafny.Seq.of(true, true)));
    };
    static fm6(p0, p1, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(767)), function (_4_i0) {
        return new _dafny.CodePoint('w'.codePointAt(0));
      });
    };
    static fm7(globalState) {
      let _source0 = _module.D2.create_DC7(false, new BigNumber(726), true, (_module.D2.create_DC7(false, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(153),function () {
  let _coll1 = new _dafny.Set();
  for (const _compr_1 of (function () {
    let _coll2 = new _dafny.Map();
    for (const _compr_2 of _dafny.IntegerRange(new BigNumber(101), new BigNumber(603))) {
      let _5_v0 = _compr_2;
      if (((new BigNumber(101)).isLessThanOrEqualTo(_5_v0)) && ((_5_v0).isLessThan(new BigNumber(603)))) {
        _coll2.push([_module.__default.safeModuloInt(_5_v0, new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false, false)).length))).length)),(_module.D1.create_DC2(false)).dtor_cf2]);
      }
    }
    return _coll2;
  }()).Keys.Elements) {
    let _6_v1 = _compr_1;
    if ((function () {
      let _coll3 = new _dafny.Map();
      for (const _compr_3 of _dafny.IntegerRange(new BigNumber(101), new BigNumber(603))) {
        let _5_v0 = _compr_3;
        if (((new BigNumber(101)).isLessThanOrEqualTo(_5_v0)) && ((_5_v0).isLessThan(new BigNumber(603)))) {
          _coll3.push([_module.__default.safeModuloInt(_5_v0, new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false, false)).length))).length)),(_module.D1.create_DC2(false)).dtor_cf2]);
        }
      }
      return _coll3;
    }()).contains(_6_v1)) {
      _coll1.add(_module.__default.safeDivisionInt(_6_v1, new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(-790))).length), _dafny.ONE)).length))).length)));
    }
  }
  return _coll1;
}())).length), false, _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(!(false)),new BigNumber((_dafny.Seq.UnicodeFromString("gcxukgk")).length)))).dtor_cf9);
      if (_source0.is_DC5) {
        let _7___mcc_h0 = (_source0).cf4;
        let _8___mcc_h1 = (_source0).cf5;
        let _9_cf5 = _8___mcc_h1;
        let _10_cf4 = _7___mcc_h0;
        return _dafny.Seq.of(new BigNumber(859));
      } else if (_source0.is_DC6) {
        return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(989)), function (_11_i0) {
          return new BigNumber((function () {
            let _coll4 = new _dafny.Set();
            for (const _compr_4 of (_dafny.Seq.of(_dafny.MultiSet.fromElements(true))).Elements) {
              let _12_v2 = _compr_4;
              if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.MultiSet.fromElements(true)), _12_v2)) {
                _coll4.add(_12_v2);
              }
            }
            return _coll4;
          }()).length);
        }), _dafny.Seq.of(new BigNumber(71)));
      } else if (_source0.is_DC7) {
        let _13___mcc_h2 = (_source0).cf6;
        let _14___mcc_h3 = (_source0).cf7;
        let _15___mcc_h4 = (_source0).cf8;
        let _16___mcc_h5 = (_source0).cf9;
        let _17_cf9 = _16___mcc_h5;
        let _18_cf8 = _15___mcc_h4;
        let _19_cf7 = _14___mcc_h3;
        let _20_cf6 = _13___mcc_h2;
        return _dafny.Seq.of(_19_cf7);
      } else {
        let _21___mcc_h6 = (_source0).cf3;
        let _22_cf3 = _21___mcc_h6;
        return _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,!(!(false)))).length), _22_cf3, _22_cf3, _22_cf3);
      }
    };
    static fm8(p0, p1, p2, p3, globalState) {
      return (_dafny.Set.fromElements(true, true, !(true))).Difference((_dafny.Set.fromElements(false, true)).Difference(_dafny.Set.fromElements(!(false))));
    };
    static fm9(globalState) {
      return (_dafny.MultiSet.fromElements(false)).Union((_dafny.MultiSet.fromElements(false, false)).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of(true, true))));
    };
    static m0(globalState) {
      let r0 = [];
      let r1 = _dafny.ZERO;
      r1 = _module.__default.fm0(globalState);
      let _23_v0;
      _23_v0 = new BigNumber(727);
      let _24_v1;
      _24_v1 = _dafny.Set.fromElements(_23_v0);
      let _25_v2;
      _25_v2 = false;
      let _26_v3;
      _26_v3 = _dafny.MultiSet.fromElements(_25_v2);
      let _27_v4;
      _27_v4 = _dafny.Seq.of(_26_v3, _26_v3);
      let _28_v5;
      _28_v5 = new _dafny.CodePoint('k'.codePointAt(0));
      let _29_v6;
      _29_v6 = _dafny.Map.Empty.slice().updateUnsafe((_27_v4)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_28_v5)).length), new BigNumber((_27_v4).length))],new BigNumber((_26_v3).cardinality()));
      let _30_v7;
      _30_v7 = _module.D2.create_DC7((_24_v1).IsProperSubsetOf(_24_v1), _23_v0, _25_v2, _29_v6);
      let _source1 = _30_v7;
      if (_source1.is_DC5) {
        let _31___mcc_h0 = (_source1).cf4;
        let _32___mcc_h1 = (_source1).cf5;
        let _33_cf5 = _32___mcc_h1;
        let _34_cf4 = _31___mcc_h0;
        let _35_v8;
        let _init0 = function (_36_i0) {
          return _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("sldfemqf"), _dafny.Seq.UnicodeFromString("e"));
        };
        let _nw0 = Array((new BigNumber(17)).toNumber());
        for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw0.length); _i0_0++) {
          _nw0[_i0_0] = _init0(new BigNumber(_i0_0));
        }
        _35_v8 = _nw0;
        let _37_v9;
        _37_v9 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("adcvwjlwm"));
        let _index0 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_35_v8).length));
        (_35_v8)[_index0] = (_37_v9).Union(_37_v9);
        let _38_v10;
        _38_v10 = _dafny.Seq.UnicodeFromString("rbff");
        let _39_v11;
        _39_v11 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("oktgtfhvg"), _38_v10);
        let _index1 = _module.__default.safeIndex(new BigNumber(144), new BigNumber((_35_v8).length));
        (_35_v8)[_index1] = ((_37_v9).Union(function () {
          let _coll5 = new _dafny.Set();
          for (const _compr_5 of (_39_v11).Elements) {
            let _40_v12 = _compr_5;
            if (_dafny.Seq.contains(_39_v11, _40_v12)) {
              _coll5.add(_40_v12);
            }
          }
          return _coll5;
        }())).Intersect(_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(589)), ((_41_v5) => function (_42_i1) {
          return _41_v5;
        })(_28_v5)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(792)), ((_43_v5) => function (_44_i2) {
          return _43_v5;
        })(_28_v5))));
        let _45_v13;
        _45_v13 = _dafny.Map.Empty.slice().updateUnsafe(((_25_v2) ? (_28_v5) : (_28_v5)),!(_25_v2));
        (globalState).f1 = (((_45_v13).contains(_module.__default.fm4(_25_v2, globalState))) ? ((_45_v13).get(_module.__default.fm4(_25_v2, globalState))) : (_25_v2));
        let _46_v14;
        _46_v14 = _dafny.Map.Empty.slice().updateUnsafe(_33_cf5,true);
        let _47_v15;
        _47_v15 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_33_cf5),(((_46_v14).contains(new BigNumber((_module.__default.fm5(_33_cf5, _34_cf4, globalState)).length))) ? ((_46_v14).get(new BigNumber((_module.__default.fm5(_33_cf5, _34_cf4, globalState)).length))) : (_25_v2)));
        let _48_v16;
        _48_v16 = _dafny.Seq.of(true, _25_v2, _25_v2);
        let _49_v17;
        _49_v17 = _dafny.Seq.of(_48_v16, _48_v16, _dafny.Seq.of(_25_v2, true, _25_v2, _25_v2, _25_v2));
        let _50_v18;
        let _nw1 = Array((new BigNumber(21)).toNumber());
        _nw1[(_dafny.ZERO).toNumber()] = _dafny.Seq.of((((_47_v15).contains(_33_cf5)) ? ((_47_v15).get(_33_cf5)) : (!(true))), _25_v2);
        _nw1[(_dafny.ONE).toNumber()] = _48_v16;
        _nw1[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_48_v16, _48_v16);
        _nw1[(new BigNumber(3)).toNumber()] = _dafny.Seq.of(_module.__default.fm1(_23_v0, _25_v2, _25_v2, true, globalState));
        _nw1[(new BigNumber(4)).toNumber()] = _48_v16;
        _nw1[(new BigNumber(5)).toNumber()] = _dafny.Seq.of(false, !(_25_v2));
        _nw1[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_48_v16, _48_v16);
        _nw1[(new BigNumber(7)).toNumber()] = _48_v16;
        _nw1[(new BigNumber(8)).toNumber()] = _48_v16;
        _nw1[(new BigNumber(9)).toNumber()] = _48_v16;
        _nw1[(new BigNumber(10)).toNumber()] = (_49_v17)[_module.__default.safeIndex(_23_v0, new BigNumber((_49_v17).length))];
        _nw1[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_48_v16, _48_v16);
        _nw1[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_48_v16, _48_v16);
        _nw1[(new BigNumber(13)).toNumber()] = _48_v16;
        _nw1[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_48_v16, _48_v16);
        _nw1[(new BigNumber(15)).toNumber()] = _48_v16;
        _nw1[(new BigNumber(16)).toNumber()] = _48_v16;
        _nw1[(new BigNumber(17)).toNumber()] = _48_v16;
        _nw1[(new BigNumber(18)).toNumber()] = _48_v16;
        _nw1[(new BigNumber(19)).toNumber()] = _48_v16;
        _nw1[(new BigNumber(20)).toNumber()] = _48_v16;
        _50_v18 = _nw1;
        let _index2 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_50_v18).length));
        (_50_v18)[_index2] = _48_v16;
        let _index3 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_50_v18).length));
        (_50_v18)[_index3] = _48_v16;
        let _rhs0 = ((true) ? (_25_v2) : ((_25_v2) || (_25_v2)));
        let _rhs1 = !(_25_v2);
        let _rhs2 = new BigNumber((_dafny.Seq.UnicodeFromString("wmqcbpj")).length);
        let _lhs0 = globalState;
        _lhs0.f1 = _rhs0;
        _25_v2 = _rhs1;
        r1 = _rhs2;
      } else if (_source1.is_DC6) {
        let _51_v19;
        _51_v19 = _dafny.Set.fromElements(_25_v2, _25_v2, _25_v2);
        let _52_v20;
        let _nw2 = new _module.C0();
        _nw2.__ctor(new BigNumber((_dafny.MultiSet.fromElements(_25_v2, _25_v2, _25_v2)).cardinality()));
        _52_v20 = _nw2;
        let _53_v21;
        _53_v21 = _dafny.Set.fromElements(_52_v20, _52_v20, _52_v20, _52_v20, _52_v20);
        let _54_v22;
        let _nw3 = Array((new BigNumber(11)).toNumber());
        _nw3[(_dafny.ZERO).toNumber()] = _25_v2;
        _nw3[(_dafny.ONE).toNumber()] = !(_25_v2) || (!(!(false)));
        _nw3[(new BigNumber(2)).toNumber()] = _25_v2;
        _nw3[(new BigNumber(3)).toNumber()] = _25_v2;
        _nw3[(new BigNumber(4)).toNumber()] = _25_v2;
        _nw3[(new BigNumber(5)).toNumber()] = (_dafny.Set.fromElements(_25_v2, true, _25_v2, false, _25_v2)).IsProperSubsetOf(_51_v19);
        _nw3[(new BigNumber(6)).toNumber()] = _25_v2;
        _nw3[(new BigNumber(7)).toNumber()] = !(_23_v0).isEqualTo(_23_v0);
        _nw3[(new BigNumber(8)).toNumber()] = _25_v2;
        _nw3[(new BigNumber(9)).toNumber()] = _25_v2;
        _nw3[(new BigNumber(10)).toNumber()] = (_53_v21).IsProperSubsetOf(_53_v21);
        _54_v22 = _nw3;
        let _index4 = _module.__default.safeIndex(new BigNumber(912), new BigNumber((_54_v22).length));
        (_54_v22)[_index4] = _25_v2;
        let _55_v23;
        _55_v23 = _dafny.Seq.UnicodeFromString("h");
        let _56_v24;
        _56_v24 = _module.D2.create_DC4(_52_v20.f5);
        let _pat_let_tv0 = _52_v20;
        let _pat_let_tv1 = _52_v20;
        let _index5 = _module.__default.safeIndex(new BigNumber(912), new BigNumber((_54_v22).length));
        (_54_v22)[_index5] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_55_v23, _55_v23), _dafny.Seq.update(_module.__default.fm6(_23_v0, function (_pat_let0_0) {
          return function (_57_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_58_dt__update_hcf3_h0) {
                return _module.D2.create_DC4(_58_dt__update_hcf3_h0);
              }(_pat_let1_0);
            }(_pat_let_tv0.f5);
          }(_pat_let0_0);
        }(_56_v24), globalState), _module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(_25_v2, _25_v2, _25_v2, _25_v2, _25_v2)).length), new BigNumber((_module.__default.fm6(_23_v0, function (_pat_let2_0) {
          return function (_59_dt__update__tmp_h1) {
            return function (_pat_let3_0) {
              return function (_60_dt__update_hcf3_h1) {
                return _module.D2.create_DC4(_60_dt__update_hcf3_h1);
              }(_pat_let3_0);
            }(_pat_let_tv1.f5);
          }(_pat_let2_0);
        }(_56_v24), globalState)).length)), _28_v5));
        let _61_v25;
        let _nw4 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.of());
        _61_v25 = _nw4;
        let _62_v26;
        let _nw5 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
        _62_v26 = _nw5;
        let _63_v27;
        _63_v27 = _dafny.Seq.of(_62_v26);
        let _64_v28;
        _64_v28 = _63_v27;
        let _index6 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_61_v25).length));
        (_61_v25)[_index6] = _dafny.Seq.of(_64_v28);
        let _65_v29;
        _65_v29 = _module.D3.create_DC10((_54_v22)[_module.__default.safeIndex(new BigNumber(912), new BigNumber((_54_v22).length))], (_54_v22)[_module.__default.safeIndex(new BigNumber(912), new BigNumber((_54_v22).length))], _23_v0, _23_v0, _52_v20.f5);
        let _66_v30;
        _66_v30 = _dafny.Seq.of(_63_v27, (((_65_v29).dtor_cf13) ? (_64_v28) : (_64_v28)), _64_v28, _64_v28);
        let _index7 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_61_v25).length));
        (_61_v25)[_index7] = _66_v30;
        let _67_v31;
        _67_v31 = _dafny.MultiSet.fromElements(_52_v20);
        let _68_v32;
        _68_v32 = _dafny.Seq.of(new BigNumber((_67_v31).cardinality()), _module.__default.fm0(globalState));
        let _69_v33;
        _69_v33 = _dafny.MultiSet.fromElements(_68_v32);
        _69_v33 = (_69_v33).Intersect(_69_v33);
        if ((_module.D3.create_DC10(_25_v2, (_54_v22)[_module.__default.safeIndex(new BigNumber(912), new BigNumber((_54_v22).length))], (_68_v32)[_module.__default.safeIndex(_23_v0, new BigNumber((_68_v32).length))], _52_v20.f5, _23_v0)).dtor_cf12) {
          let _70_v34;
          _70_v34 = _dafny.Map.Empty.slice().updateUnsafe(_28_v5,_52_v20.f5);
          let _71_v35;
          _71_v35 = _dafny.MultiSet.fromElements(new BigNumber(-939), new BigNumber((_70_v34).length));
          (globalState).f1 = ((_26_v3).IsProperSubsetOf(_dafny.MultiSet.fromElements(_module.__default.fm1((((_71_v35).contains(_23_v0)) ? ((_71_v35).get(_23_v0)) : (_23_v0)), _module.__default.fm1(_52_v20.f5, _25_v2, (_54_v22)[_module.__default.safeIndex(new BigNumber(912), new BigNumber((_54_v22).length))], _25_v2, globalState), (_54_v22)[_module.__default.safeIndex(new BigNumber(912), new BigNumber((_54_v22).length))], true, globalState)))) === (!(_25_v2));
          let _72_v36;
          let _nw6 = Array((new BigNumber(10)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _72_v36 = _nw6;
          let _73_v37;
          _73_v37 = _dafny.MultiSet.fromElements(_72_v36);
          let _74_v38;
          _74_v38 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(new BigNumber((_73_v37).cardinality()), new BigNumber(-224)),(_54_v22)[_module.__default.safeIndex(new BigNumber(912), new BigNumber((_54_v22).length))]);
          _74_v38 = (_74_v38).update(_52_v20.f5, (_54_v22)[_module.__default.safeIndex(new BigNumber(912), new BigNumber((_54_v22).length))]);
          let _75_v39;
          let _nw7 = new _module.C0();
          _nw7.__ctor(_module.__default.safeModuloInt(_52_v20.f5, new BigNumber(844)));
          _75_v39 = _nw7;
          let _nw8 = new _module.C0();
          _nw8.__ctor((_75_v39.f5).multipliedBy(_23_v0));
          _75_v39 = _nw8;
          (_52_v20).f5 = _module.__default.fm0(globalState);
        } else {
          _68_v32 = _68_v32;
          (globalState).f1 = _25_v2;
          let _76_v40;
          let _nw9 = Array((new BigNumber(27)).toNumber());
          _nw9[(_dafny.ZERO).toNumber()] = _module.__default.fm4(false, globalState);
          _nw9[(_dafny.ONE).toNumber()] = _28_v5;
          _nw9[(new BigNumber(2)).toNumber()] = _28_v5;
          _nw9[(new BigNumber(3)).toNumber()] = _28_v5;
          _nw9[(new BigNumber(4)).toNumber()] = _28_v5;
          _nw9[(new BigNumber(5)).toNumber()] = new _dafny.CodePoint('a'.codePointAt(0));
          _nw9[(new BigNumber(6)).toNumber()] = _28_v5;
          _nw9[(new BigNumber(7)).toNumber()] = _28_v5;
          _nw9[(new BigNumber(8)).toNumber()] = _28_v5;
          _nw9[(new BigNumber(9)).toNumber()] = _28_v5;
          _nw9[(new BigNumber(10)).toNumber()] = _28_v5;
          _nw9[(new BigNumber(11)).toNumber()] = _28_v5;
          _nw9[(new BigNumber(12)).toNumber()] = _28_v5;
          _nw9[(new BigNumber(13)).toNumber()] = _28_v5;
          _nw9[(new BigNumber(14)).toNumber()] = _28_v5;
          _nw9[(new BigNumber(15)).toNumber()] = _28_v5;
          _nw9[(new BigNumber(16)).toNumber()] = _28_v5;
          _nw9[(new BigNumber(17)).toNumber()] = new _dafny.CodePoint('n'.codePointAt(0));
          _nw9[(new BigNumber(18)).toNumber()] = _28_v5;
          _nw9[(new BigNumber(19)).toNumber()] = _28_v5;
          _nw9[(new BigNumber(20)).toNumber()] = _28_v5;
          _nw9[(new BigNumber(21)).toNumber()] = ((_module.__default.fm1(_52_v20.f5, !((_54_v22)[_module.__default.safeIndex(new BigNumber(912), new BigNumber((_54_v22).length))]), _25_v2, (_54_v22)[_module.__default.safeIndex(new BigNumber(912), new BigNumber((_54_v22).length))], globalState)) ? (_28_v5) : (new _dafny.CodePoint('c'.codePointAt(0))));
          _nw9[(new BigNumber(22)).toNumber()] = _28_v5;
          _nw9[(new BigNumber(23)).toNumber()] = new _dafny.CodePoint('n'.codePointAt(0));
          _nw9[(new BigNumber(24)).toNumber()] = _28_v5;
          _nw9[(new BigNumber(25)).toNumber()] = _28_v5;
          _nw9[(new BigNumber(26)).toNumber()] = _28_v5;
          _76_v40 = _nw9;
          let _index8 = _module.__default.safeIndex(new BigNumber(672), new BigNumber((_76_v40).length));
          (_76_v40)[_index8] = _28_v5;
          let _77_v41;
          _77_v41 = _dafny.Map.Empty.slice().updateUnsafe(_55_v23,_dafny.Seq.Concat(_68_v32, _68_v32));
          let _index9 = _module.__default.safeIndex(new BigNumber(672), new BigNumber((_76_v40).length));
          let _rhs3 = _52_v20.f5;
          let _rhs4 = new _dafny.CodePoint('v'.codePointAt(0));
          let _rhs5 = ((_52_v20.f5).multipliedBy(_52_v20.f5)).isEqualTo(_23_v0);
          let _rhs6 = (((_77_v41).contains(_dafny.Seq.Concat(_55_v23, _55_v23))) ? ((_77_v41).get(_dafny.Seq.Concat(_55_v23, _55_v23))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(715)), ((_78_v20) => function (_79_i3) {
            return _78_v20.f5;
          })(_52_v20))));
          let _lhs1 = _52_v20;
          let _lhs2 = _76_v40;
          let _lhs3 = _module.__default.safeIndex(new BigNumber(672), new BigNumber((_76_v40).length));
          let _lhs4 = globalState;
          _lhs1.f5 = _rhs3;
          _lhs2[_lhs3] = _rhs4;
          _lhs4.f1 = _rhs5;
          _68_v32 = _rhs6;
          let _80_v42;
          let _nw10 = new _module.C0();
          _nw10.__ctor(_52_v20.f5);
          _80_v42 = _nw10;
          let _nw11 = Array((new BigNumber(16)).toNumber()).fill(false);
          _54_v22 = _nw11;
        }
      } else if (_source1.is_DC7) {
        let _81___mcc_h2 = (_source1).cf6;
        let _82___mcc_h3 = (_source1).cf7;
        let _83___mcc_h4 = (_source1).cf8;
        let _84___mcc_h5 = (_source1).cf9;
        let _85_cf9 = _84___mcc_h5;
        let _86_cf8 = _83___mcc_h4;
        let _87_cf7 = _82___mcc_h3;
        let _88_cf6 = _81___mcc_h2;
        let _89_v43;
        _89_v43 = _module.D2.create_DC4(_23_v0);
        let _90_v44;
        _90_v44 = _dafny.Seq.UnicodeFromString("xautmy");
        let _91_v45;
        _91_v45 = _module.D5.create_DC12(_dafny.Seq.update(_dafny.Seq.of(_89_v43, _89_v43, _89_v43, _89_v43, _89_v43), _module.__default.safeIndex(new BigNumber((_90_v44).length), new BigNumber((_dafny.Seq.of(_89_v43, _89_v43, _89_v43, _89_v43, _89_v43)).length)), _module.D2.create_DC4(_23_v0)));
        let _92_v46;
        _92_v46 = _dafny.Seq.of(_89_v43);
        let _93_v47;
        _93_v47 = _dafny.Seq.of((_dafny.ZERO).minus(_87_cf7));
        let _pat_let_tv2 = _89_v43;
        let _94_v48;
        let _nw12 = Array((new BigNumber(28)).toNumber());
        _nw12[(_dafny.ZERO).toNumber()] = ((_86_cf8) ? (_86_cf8) : (_25_v2));
        _nw12[(_dafny.ONE).toNumber()] = _88_cf6;
        _nw12[(new BigNumber(2)).toNumber()] = _25_v2;
        _nw12[(new BigNumber(3)).toNumber()] = true;
        _nw12[(new BigNumber(4)).toNumber()] = _88_cf6;
        _nw12[(new BigNumber(5)).toNumber()] = !(_86_cf8);
        _nw12[(new BigNumber(6)).toNumber()] = _25_v2;
        _nw12[(new BigNumber(7)).toNumber()] = _25_v2;
        _nw12[(new BigNumber(8)).toNumber()] = (_87_cf7).isEqualTo(new BigNumber((_dafny.Set.fromElements(_87_cf7)).length));
        _nw12[(new BigNumber(9)).toNumber()] = _88_cf6;
        _nw12[(new BigNumber(10)).toNumber()] = (_dafny.MultiSet.FromArray(_92_v46)).IsProperSubsetOf(_dafny.MultiSet.FromArray((function (_pat_let4_0) {
          return function (_95_dt__update__tmp_h2) {
            return function (_pat_let5_0) {
              return function (_96_dt__update_hcf18_h0) {
                return _module.D5.create_DC12(_96_dt__update_hcf18_h0);
              }(_pat_let5_0);
            }(_dafny.Seq.of(_pat_let_tv2));
          }(_pat_let4_0);
        }(_91_v45)).dtor_cf18));
        _nw12[(new BigNumber(11)).toNumber()] = ((!(_25_v2)) ? (_86_cf8) : (_25_v2));
        _nw12[(new BigNumber(12)).toNumber()] = _86_cf8;
        _nw12[(new BigNumber(13)).toNumber()] = _25_v2;
        _nw12[(new BigNumber(14)).toNumber()] = !(_88_cf6);
        _nw12[(new BigNumber(15)).toNumber()] = (new BigNumber((_90_v44).length)).isLessThanOrEqualTo(new BigNumber((_module.__default.fm7(globalState)).length));
        _nw12[(new BigNumber(16)).toNumber()] = _86_cf8;
        _nw12[(new BigNumber(17)).toNumber()] = !((_dafny.ZERO).minus(_23_v0)).isEqualTo((_dafny.ZERO).minus(_23_v0));
        _nw12[(new BigNumber(18)).toNumber()] = _25_v2;
        _nw12[(new BigNumber(19)).toNumber()] = _25_v2;
        _nw12[(new BigNumber(20)).toNumber()] = _module.__default.fm1(new BigNumber((_93_v47).length), true, false, _86_cf8, globalState);
        _nw12[(new BigNumber(21)).toNumber()] = (_26_v3).IsProperSubsetOf(_26_v3);
        _nw12[(new BigNumber(22)).toNumber()] = _88_cf6;
        _nw12[(new BigNumber(23)).toNumber()] = _86_cf8;
        _nw12[(new BigNumber(24)).toNumber()] = !((_30_v7).dtor_cf8);
        _nw12[(new BigNumber(25)).toNumber()] = _86_cf8;
        _nw12[(new BigNumber(26)).toNumber()] = !_dafny.areEqual(_90_v44, _90_v44);
        _nw12[(new BigNumber(27)).toNumber()] = _25_v2;
        _94_v48 = _nw12;
        let _index10 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_94_v48).length));
        (_94_v48)[_index10] = _25_v2;
        let _97_v49;
        _97_v49 = _dafny.Map.Empty.slice().updateUnsafe(_23_v0,_module.__default.fm0(globalState));
        let _98_v50;
        let _nw13 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
        _98_v50 = _nw13;
        let _index11 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_94_v48).length));
        let _rhs7 = ((true) ? (_98_v50) : (_98_v50));
        let _rhs8 = _98_v50;
        let _rhs9 = ((_86_cf8) ? (_25_v2) : (_88_cf6));
        let _rhs10 = _dafny.Map.Empty.slice().updateUnsafe(_87_cf7,(_93_v47)[_module.__default.safeIndex((_dafny.ZERO).minus(_23_v0), new BigNumber((_93_v47).length))]);
        let _rhs11 = _94_v48;
        let _lhs5 = _94_v48;
        let _lhs6 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_94_v48).length));
        r0 = _rhs7;
        r0 = _rhs8;
        _lhs5[_lhs6] = _rhs9;
        _97_v49 = _rhs10;
        _94_v48 = _rhs11;
        let _99_v51;
        let _nw14 = new _module.C0();
        _nw14.__ctor(_23_v0);
        _99_v51 = _nw14;
        let _100_v52;
        let _nw15 = Array((new BigNumber(6)).toNumber());
        _nw15[(_dafny.ZERO).toNumber()] = _99_v51;
        _nw15[(_dafny.ONE).toNumber()] = _99_v51;
        _nw15[(new BigNumber(2)).toNumber()] = _99_v51;
        _nw15[(new BigNumber(3)).toNumber()] = ((_module.__default.fm1(_23_v0, false, false, _88_cf6, globalState)) ? (_99_v51) : (_99_v51));
        _nw15[(new BigNumber(4)).toNumber()] = _99_v51;
        _nw15[(new BigNumber(5)).toNumber()] = _99_v51;
        _100_v52 = _nw15;
        let _index12 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_100_v52).length));
        (_100_v52)[_index12] = _99_v51;
        let _101_v53;
        _101_v53 = _dafny.Map.Empty.slice().updateUnsafe(_87_cf7,(((_94_v48)[_module.__default.safeIndex(new BigNumber(203), new BigNumber((_94_v48).length))]) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(216)), ((_102_v5) => function (_103_i4) {
          return _102_v5;
        })(_28_v5))) : (_90_v44)));
        let _index13 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_100_v52).length));
        let _rhs12 = _99_v51;
        let _rhs13 = (_101_v53).Merge(_101_v53);
        let _lhs7 = _100_v52;
        let _lhs8 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_100_v52).length));
        _lhs7[_lhs8] = _rhs12;
        _101_v53 = _rhs13;
        if (_25_v2) {
          _97_v49 = (_97_v49).update((_dafny.ZERO).minus(_99_v51.f5), _module.__default.fm0(globalState));
          let _104_v54;
          _104_v54 = _dafny.MultiSet.fromElements(_98_v50);
          _104_v54 = _104_v54;
          let _105_v55;
          _105_v55 = _dafny.Set.fromElements(!(!(false)));
          let _106_v56;
          let _nw16 = Array((new BigNumber(14)).toNumber());
          _nw16[(_dafny.ZERO).toNumber()] = _98_v50;
          _nw16[(_dafny.ONE).toNumber()] = _98_v50;
          _nw16[(new BigNumber(2)).toNumber()] = _98_v50;
          _nw16[(new BigNumber(3)).toNumber()] = _98_v50;
          _nw16[(new BigNumber(4)).toNumber()] = _98_v50;
          _nw16[(new BigNumber(5)).toNumber()] = _98_v50;
          _nw16[(new BigNumber(6)).toNumber()] = _98_v50;
          _nw16[(new BigNumber(7)).toNumber()] = _98_v50;
          _nw16[(new BigNumber(8)).toNumber()] = _98_v50;
          _nw16[(new BigNumber(9)).toNumber()] = _98_v50;
          _nw16[(new BigNumber(10)).toNumber()] = _98_v50;
          _nw16[(new BigNumber(11)).toNumber()] = _98_v50;
          _nw16[(new BigNumber(12)).toNumber()] = _98_v50;
          _nw16[(new BigNumber(13)).toNumber()] = _98_v50;
          _106_v56 = _nw16;
          let _index14 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_106_v56).length));
          (_106_v56)[_index14] = _98_v50;
          let _107_v57;
          let _init1 = ((_108_v47) => function (_109_i5) {
            return _108_v47;
          })(_93_v47);
          let _nw17 = Array((new BigNumber(21)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw17.length); _i0_1++) {
            _nw17[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _107_v57 = _nw17;
          let _index15 = _module.__default.safeIndex(new BigNumber(719), new BigNumber((_107_v57).length));
          (_107_v57)[_index15] = _dafny.Seq.of(new BigNumber((_90_v44).length), _87_cf7);
          let _110_v58;
          _110_v58 = _dafny.Seq.of((_94_v48)[_module.__default.safeIndex(new BigNumber(203), new BigNumber((_94_v48).length))]);
          let _index16 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_106_v56).length));
          let _index17 = _module.__default.safeIndex(new BigNumber(719), new BigNumber((_107_v57).length));
          let _rhs14 = _module.__default.fm8(_88_cf6, _86_cf8, _dafny.Seq.Concat(_110_v58, _110_v58), (_105_v55).equals(_105_v55), globalState);
          let _rhs15 = _98_v50;
          let _rhs16 = _module.__default.fm7(globalState);
          let _lhs9 = _106_v56;
          let _lhs10 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_106_v56).length));
          let _lhs11 = _107_v57;
          let _lhs12 = _module.__default.safeIndex(new BigNumber(719), new BigNumber((_107_v57).length));
          _105_v55 = _rhs14;
          _lhs9[_lhs10] = _rhs15;
          _lhs11[_lhs12] = _rhs16;
          let _111_v59;
          _111_v59 = _dafny.MultiSet.fromElements(_26_v3, _26_v3);
          _111_v59 = (_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(422)), ((_112_v3) => function (_113_i6) {
            return _112_v3;
          })(_26_v3)), _dafny.Seq.of(_26_v3, _dafny.MultiSet.FromArray(_110_v58), _module.__default.fm9(globalState), _26_v3)))).Intersect(_111_v59);
          let _index18 = _module.__default.safeIndex(new BigNumber(165), new BigNumber((_98_v50).length));
          (_98_v50)[_index18] = _99_v51.f5;
          let _index19 = _module.__default.safeIndex(new BigNumber(165), new BigNumber((_98_v50).length));
          (_98_v50)[_index19] = _87_cf7;
        } else {
          _94_v48 = _94_v48;
          let _rhs17 = _module.__default.safeDivisionInt(_23_v0, (_dafny.ZERO).minus(_87_cf7));
          let _rhs18 = _88_cf6;
          let _lhs13 = _99_v51;
          let _lhs14 = globalState;
          _lhs13.f5 = _rhs17;
          _lhs14.f1 = _rhs18;
          _86_cf8 = _25_v2;
          _28_v5 = _28_v5;
          let _114_v60;
          let _nw18 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _114_v60 = _nw18;
          let _index20 = _module.__default.safeIndex(new BigNumber(452), new BigNumber((_114_v60).length));
          (_114_v60)[_index20] = _90_v44;
          let _index21 = _module.__default.safeIndex(new BigNumber(452), new BigNumber((_114_v60).length));
          let _rhs19 = _114_v60;
          let _rhs20 = _dafny.Seq.Concat(_dafny.Seq.Concat(_90_v44, _dafny.Seq.Create(_module.__default.abs(new BigNumber(621)), ((_115_v44, _116_v51) => function (_117_i7) {
            return (_115_v44)[_module.__default.safeIndex(_116_v51.f5, new BigNumber((_115_v44).length))];
          })(_90_v44, _99_v51))), _90_v44);
          let _lhs15 = _114_v60;
          let _lhs16 = _module.__default.safeIndex(new BigNumber(452), new BigNumber((_114_v60).length));
          _114_v60 = _rhs19;
          _lhs15[_lhs16] = _rhs20;
        }
        let _118_v61;
        _118_v61 = _dafny.Seq.of(_98_v50);
        let _119_v62;
        _119_v62 = _dafny.MultiSet.fromElements(new BigNumber(-259));
        let _120_v63;
        _120_v63 = _dafny.Seq.of(_98_v50);
        _118_v61 = (((_119_v62).IsProperSubsetOf(_119_v62)) ? (_120_v63) : (_120_v63));
      } else {
        let _121___mcc_h6 = (_source1).cf3;
        let _122_cf3 = _121___mcc_h6;
        let _123_v64;
        let _nw19 = Array((new BigNumber(22)).toNumber()).fill(false);
        _123_v64 = _nw19;
        let _index22 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_123_v64).length));
        (_123_v64)[_index22] = _25_v2;
        let _124_v65;
        _124_v65 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(_23_v0, _23_v0),true);
        let _125_v66;
        _125_v66 = _module.D2.create_DC4(_23_v0);
        let _126_v67;
        _126_v67 = _dafny.Seq.of(_25_v2);
        let _index23 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_123_v64).length));
        let _rhs21 = _25_v2;
        let _rhs22 = _module.__default.fm1(_122_cf3, ((!(_module.__default.fm1(_122_cf3, _25_v2, true, _25_v2, globalState))) ? (_25_v2) : ((_126_v67)[_module.__default.safeIndex(_122_cf3, new BigNumber((_126_v67).length))])), (_122_cf3).isEqualTo(_23_v0), !(_25_v2), globalState);
        let _rhs23 = (_124_v65).update(_23_v0, _25_v2);
        let _rhs24 = _module.D2.create_DC4(_23_v0);
        let _rhs25 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_122_cf3, _122_cf3));
        let _lhs17 = _123_v64;
        let _lhs18 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_123_v64).length));
        let _lhs19 = globalState;
        _lhs17[_lhs18] = _rhs21;
        _lhs19.f1 = _rhs22;
        _124_v65 = _rhs23;
        _125_v66 = _rhs24;
        _122_cf3 = _rhs25;
        (globalState).f1 = (_123_v64)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_123_v64).length))];
        _28_v5 = _28_v5;
        let _127_v68;
        _127_v68 = _dafny.Map.Empty.slice().updateUnsafe(_28_v5,new BigNumber(382));
        let _128_v69;
        let _nw20 = new _module.C0();
        _nw20.__ctor(_23_v0);
        _128_v69 = _nw20;
        let _129_v70;
        _129_v70 = _dafny.Map.Empty.slice().updateUnsafe((_123_v64)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_123_v64).length))],(_123_v64)[_module.__default.safeIndex(new BigNumber(910), new BigNumber((_123_v64).length))]);
        let _130_v71;
        _130_v71 = _dafny.MultiSet.fromElements(_28_v5);
        let _131_v72;
        let _nw21 = Array((new BigNumber(13)).toNumber());
        _nw21[(_dafny.ZERO).toNumber()] = _122_cf3;
        _nw21[(_dafny.ONE).toNumber()] = new BigNumber((_129_v70).length);
        _nw21[(new BigNumber(2)).toNumber()] = new BigNumber((_130_v71).cardinality());
        _nw21[(new BigNumber(3)).toNumber()] = _128_v69.f5;
        _nw21[(new BigNumber(4)).toNumber()] = new BigNumber(679);
        _nw21[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(_122_cf3);
        _nw21[(new BigNumber(6)).toNumber()] = _128_v69.f5;
        _nw21[(new BigNumber(7)).toNumber()] = _23_v0;
        _nw21[(new BigNumber(8)).toNumber()] = new BigNumber(821);
        _nw21[(new BigNumber(9)).toNumber()] = _128_v69.f5;
        _nw21[(new BigNumber(10)).toNumber()] = _23_v0;
        _nw21[(new BigNumber(11)).toNumber()] = _23_v0;
        _nw21[(new BigNumber(12)).toNumber()] = _128_v69.f5;
        _131_v72 = _nw21;
        let _132_v73;
        _132_v73 = _dafny.Seq.of(_131_v72, _131_v72);
        let _133_v74;
        _133_v74 = _132_v73;
        let _134_v75;
        _134_v75 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_128_v69)).length),_133_v74);
        _127_v68 = (_127_v68).update(_28_v5, new BigNumber((_134_v75).length));
      }
      let _135_v76;
      let _nw22 = new _module.C0();
      _nw22.__ctor((_23_v0).multipliedBy(_23_v0));
      _135_v76 = _nw22;
      let _136_v77;
      _136_v77 = _dafny.MultiSet.fromElements((_135_v76.f5).multipliedBy(_135_v76.f5), _23_v0);
      _136_v77 = (_136_v77).Intersect(_136_v77);
      let _pat_let_tv3 = _29_v6;
      let _pat_let_tv4 = _23_v0;
      _30_v7 = function (_pat_let6_0) {
        return function (_137_dt__update__tmp_h4) {
          return function (_pat_let7_0) {
            return function (_138_dt__update_hcf9_h0) {
              return function (_pat_let8_0) {
                return function (_139_dt__update_hcf7_h0) {
                  return _module.D2.create_DC7((_137_dt__update__tmp_h4).dtor_cf6, _139_dt__update_hcf7_h0, (_137_dt__update__tmp_h4).dtor_cf8, _138_dt__update_hcf9_h0);
                }(_pat_let8_0);
              }(_pat_let_tv4);
            }(_pat_let7_0);
          }(_pat_let_tv3);
        }(_pat_let6_0);
      }(_30_v7);
      let _140_v78;
      _140_v78 = _module.D3.create_DC10(_25_v2, _25_v2, _23_v0, _135_v76.f5, _23_v0);
      let _141_v79;
      _141_v79 = _dafny.Set.fromElements(_25_v2);
      let _142_v80;
      _142_v80 = _dafny.Seq.of(_module.__default.fm1(_135_v76.f5, _25_v2, _25_v2, _25_v2, globalState), _25_v2, false);
      let _143_v81;
      _143_v81 = _module.D3.create_DC9(_136_v77);
      let _144_v82;
      _144_v82 = _dafny.MultiSet.fromElements(_143_v81, _143_v81);
      let _145_v83;
      _145_v83 = _dafny.Set.fromElements(_24_v1);
      let _146_v84;
      let _nw23 = Array((new BigNumber(20)).toNumber());
      _nw23[(_dafny.ZERO).toNumber()] = (_140_v78).dtor_cf12;
      _nw23[(_dafny.ONE).toNumber()] = _25_v2;
      _nw23[(new BigNumber(2)).toNumber()] = _25_v2;
      _nw23[(new BigNumber(3)).toNumber()] = ((!(!(_25_v2))) ? (_25_v2) : (_25_v2));
      _nw23[(new BigNumber(4)).toNumber()] = ((_dafny.ZERO).minus(new BigNumber((_141_v79).length))).isLessThanOrEqualTo(_135_v76.f5);
      _nw23[(new BigNumber(5)).toNumber()] = (_142_v80)[_module.__default.safeIndex(_135_v76.f5, new BigNumber((_142_v80).length))];
      _nw23[(new BigNumber(6)).toNumber()] = _25_v2;
      _nw23[(new BigNumber(7)).toNumber()] = _25_v2;
      _nw23[(new BigNumber(8)).toNumber()] = _25_v2;
      _nw23[(new BigNumber(9)).toNumber()] = _25_v2;
      _nw23[(new BigNumber(10)).toNumber()] = _25_v2;
      _nw23[(new BigNumber(11)).toNumber()] = _25_v2;
      _nw23[(new BigNumber(12)).toNumber()] = _25_v2;
      _nw23[(new BigNumber(13)).toNumber()] = _25_v2;
      _nw23[(new BigNumber(14)).toNumber()] = true;
      _nw23[(new BigNumber(15)).toNumber()] = _25_v2;
      _nw23[(new BigNumber(16)).toNumber()] = _25_v2;
      _nw23[(new BigNumber(17)).toNumber()] = _25_v2;
      _nw23[(new BigNumber(18)).toNumber()] = !(_144_v82).equals(_dafny.MultiSet.fromElements(_143_v81));
      _nw23[(new BigNumber(19)).toNumber()] = (_145_v83).IsSubsetOf(_145_v83);
      _146_v84 = _nw23;
      let _index24 = _module.__default.safeIndex(new BigNumber(382), new BigNumber((_146_v84).length));
      (_146_v84)[_index24] = _25_v2;
      let _index25 = _module.__default.safeIndex(new BigNumber(382), new BigNumber((_146_v84).length));
      (_146_v84)[_index25] = (_140_v78).dtor_cf12;
      let _nw24 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
      r0 = _nw24;
      let _147_v85;
      _147_v85 = _dafny.Seq.UnicodeFromString("matixu");
      let _148_v86;
      _148_v86 = _dafny.Map.Empty.slice().updateUnsafe(_147_v85,_dafny.Seq.of(_28_v5));
      let _149_v87;
      _149_v87 = _dafny.Seq.of(new BigNumber((_dafny.Seq.Concat((((_148_v86).contains(_147_v85)) ? ((_148_v86).get(_147_v85)) : (_147_v85)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(814)), ((_150_v5) => function (_151_i8) {
        return _150_v5;
      })(_28_v5)))).length), _135_v76.f5, (_23_v0).multipliedBy(_135_v76.f5));
      r1 = (_149_v87)[_module.__default.safeIndex(_23_v0, new BigNumber((_149_v87).length))];
      return [r0, r1];
    }
    static Main(__noArgsParameter) {
      let _152_v0;
      _152_v0 = new BigNumber(-560);
      let _153_v1;
      _153_v1 = _dafny.Seq.of(_152_v0, _152_v0);
      let _154_v2;
      let _init2 = function (_155_i0) {
        return _dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0)));
      };
      let _nw25 = Array((new BigNumber(19)).toNumber());
      for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw25.length); _i0_2++) {
        _nw25[_i0_2] = _init2(new BigNumber(_i0_2));
      }
      _154_v2 = _nw25;
      let _156_globalState;
      let _nw26 = new _module.GlobalState();
      _nw26.__ctor(new BigNumber(791), false, _153_v1, _dafny.Seq.UnicodeFromString("kpno"), _154_v2);
      _156_globalState = _nw26;
      let _157_v3;
      _157_v3 = true;
      let _158_i1;
      _158_i1 = _dafny.ZERO;
      L0: {
        while (_157_v3) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_158_i1)) {
              break L0;
            }
            _158_i1 = (_158_i1).plus(_dafny.ONE);
            let _159_v4;
            let _nw27 = Array((new BigNumber(19)).toNumber()).fill(false);
            _159_v4 = _nw27;
            let _index26 = _module.__default.safeIndex(new BigNumber(258), new BigNumber((_159_v4).length));
            (_159_v4)[_index26] = false;
            let _index27 = _module.__default.safeIndex(new BigNumber(258), new BigNumber((_159_v4).length));
            (_159_v4)[_index27] = !(_157_v3) || (_157_v3);
            if (!((_159_v4)[_module.__default.safeIndex(new BigNumber(258), new BigNumber((_159_v4).length))])) {
              let _160_v5;
              let _161_v6;
              let _out0;
              let _out1;
              let _outcollector0 = _module.__default.m0(_156_globalState);
              _out0 = _outcollector0[0];
              _out1 = _outcollector0[1];
              _160_v5 = _out0;
              _161_v6 = _out1;
              (_156_globalState).f1 = !(_157_v3) || (_157_v3);
              _153_v1 = _dafny.Seq.update((_module.D0.create_DC0(_153_v1)).dtor_cf0, _module.__default.safeIndex(_module.__default.fm0(_156_globalState), new BigNumber(((_module.D0.create_DC0(_153_v1)).dtor_cf0).length)), _152_v0);
              (_156_globalState).f1 = (_159_v4)[_module.__default.safeIndex(new BigNumber(258), new BigNumber((_159_v4).length))];
              _152_v0 = _module.__default.safeModuloInt(_module.__default.fm0(_156_globalState), _module.__default.safeDivisionInt(_161_v6, _152_v0));
            } else {
              let _162_v7;
              let _163_v8;
              let _out2;
              let _out3;
              let _outcollector1 = _module.__default.m0(_156_globalState);
              _out2 = _outcollector1[0];
              _out3 = _outcollector1[1];
              _162_v7 = _out2;
              _163_v8 = _out3;
              _157_v3 = (_159_v4)[_module.__default.safeIndex(new BigNumber(258), new BigNumber((_159_v4).length))];
              let _164_v9;
              _164_v9 = _module.D1.create_DC2((_159_v4)[_module.__default.safeIndex(new BigNumber(258), new BigNumber((_159_v4).length))]);
              let _index28 = _module.__default.safeIndex(new BigNumber(258), new BigNumber((_159_v4).length));
              (_159_v4)[_index28] = (_164_v9).dtor_cf2;
              let _165_v10;
              let _nw28 = new _module.C0();
              _nw28.__ctor(_152_v0);
              _165_v10 = _nw28;
              let _rhs26 = _163_v8;
              let _rhs27 = _module.__default.safeModuloInt(new BigNumber(90), _152_v0);
              let _lhs20 = _165_v10;
              _152_v0 = _rhs26;
              _lhs20.f5 = _rhs27;
            }
            (_156_globalState).f1 = true;
            let _166_v11;
            _166_v11 = _dafny.MultiSet.fromElements(_157_v3, (_159_v4)[_module.__default.safeIndex(new BigNumber(258), new BigNumber((_159_v4).length))]);
            let _167_v12;
            _167_v12 = _module.D0.create_DC0(((_157_v3) ? (_153_v1) : (_dafny.Seq.of(new BigNumber((_166_v11).cardinality())))));
            let _rhs28 = ((_157_v3) ? (new BigNumber(612)) : (_152_v0));
            let _rhs29 = _167_v12;
            _152_v0 = _rhs28;
            _167_v12 = _rhs29;
          }
        }
      }
      let _168_v13;
      let _nw29 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
      _168_v13 = _nw29;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_168_v13).length))) {
        let _169_i2 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_169_i2)) && ((_169_i2).isLessThan(new BigNumber((_168_v13).length))))) {
          (_168_v13)[(_169_i2)] = _module.__default.safeModuloInt(_169_i2, (_module.D2.create_DC4(_152_v0)).dtor_cf3);
        }
      }
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_168_v13).length))) {
        let _170_i3 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_170_i3)) && ((_170_i3).isLessThan(new BigNumber((_168_v13).length))))) {
          (_168_v13)[(_170_i3)] = (_170_i3).multipliedBy(_152_v0);
        }
      }
      (_156_globalState).f3 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-708)), function (_171_i4) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      });
      let _172_v14;
      let _nw30 = new _module.C0();
      _nw30.__ctor(new BigNumber(190));
      _172_v14 = _nw30;
      let _173_v15;
      _173_v15 = _dafny.Seq.UnicodeFromString("xqlvtuvhq");
      let _174_i5;
      _174_i5 = _dafny.ZERO;
      L1: {
        while ((((_152_v0).isLessThanOrEqualTo(_152_v0)) ? (_157_v3) : (!_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(671)), function (_178_i6) {
          return new _dafny.CodePoint('g'.codePointAt(0));
        }), _173_v15)))) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_174_i5)) {
              break L1;
            }
            _174_i5 = (_174_i5).plus(_dafny.ONE);
            (_172_v14).f5 = new BigNumber((_173_v15).length);
            let _175_v16;
            _175_v16 = _dafny.Map.Empty.slice().updateUnsafe(_157_v3,_172_v14.f5);
            let _176_v17;
            _176_v17 = _dafny.Map.Empty.slice().updateUnsafe(_152_v0,_175_v16);
            let _177_v18;
            _177_v18 = _dafny.Map.Empty.slice().updateUnsafe(_176_v17,_172_v14.f5);
            _177_v18 = (_177_v18).update(_176_v17, _172_v14.f5);
            _153_v1 = _153_v1;
            _152_v0 = _172_v14.f5;
          }
        }
      }
      let _179_v19;
      _179_v19 = new _dafny.CodePoint('q'.codePointAt(0));
      _179_v19 = (_173_v15)[_module.__default.safeIndex(new BigNumber(764), new BigNumber((_173_v15).length))];
      let _180_v20;
      let _init3 = ((_181_v0, _182_v14) => function (_183_i7) {
        return _dafny.Map.Empty.slice().updateUnsafe(_181_v0,_182_v14.f5);
      })(_152_v0, _172_v14);
      let _nw31 = Array((new BigNumber(17)).toNumber());
      for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw31.length); _i0_3++) {
        _nw31[_i0_3] = _init3(new BigNumber(_i0_3));
      }
      _180_v20 = _nw31;
      let _184_v21;
      _184_v21 = _dafny.Map.Empty.slice().updateUnsafe(_157_v3,_152_v0);
      let _185_v22;
      _185_v22 = _dafny.Map.Empty.slice().updateUnsafe(_172_v14.f5,new BigNumber((_184_v21).length));
      let _index29 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_180_v20).length));
      (_180_v20)[_index29] = _185_v22;
      let _index30 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_180_v20).length));
      (_180_v20)[_index30] = ((_185_v22).Merge(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_172_v14.f5),_172_v14.f5))).update(_152_v0, _172_v14.f5);
      let _hi0 = _152_v0;
      for (let _186_i8 = new BigNumber(232); _186_i8.isLessThan(_hi0); _186_i8 = _186_i8.plus(_dafny.ONE)) {
        let _187_v23;
        let _nw32 = new _module.C0();
        _nw32.__ctor(_152_v0);
        _187_v23 = _nw32;
        let _188_v24;
        let _189_v25;
        let _out4;
        let _out5;
        let _outcollector2 = _module.__default.m0(_156_globalState);
        _out4 = _outcollector2[0];
        _out5 = _outcollector2[1];
        _188_v24 = _out4;
        _189_v25 = _out5;
        let _index31 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_188_v24).length));
        (_188_v24)[_index31] = _186_i8;
        let _190_v26;
        let _nw33 = Array((new BigNumber(13)).toNumber());
        _nw33[(_dafny.ZERO).toNumber()] = !(_157_v3);
        _nw33[(_dafny.ONE).toNumber()] = false;
        _nw33[(new BigNumber(2)).toNumber()] = !(_157_v3);
        _nw33[(new BigNumber(3)).toNumber()] = _157_v3;
        _nw33[(new BigNumber(4)).toNumber()] = _157_v3;
        _nw33[(new BigNumber(5)).toNumber()] = _157_v3;
        _nw33[(new BigNumber(6)).toNumber()] = _157_v3;
        _nw33[(new BigNumber(7)).toNumber()] = _157_v3;
        _nw33[(new BigNumber(8)).toNumber()] = _157_v3;
        _nw33[(new BigNumber(9)).toNumber()] = _157_v3;
        _nw33[(new BigNumber(10)).toNumber()] = _157_v3;
        _nw33[(new BigNumber(11)).toNumber()] = _157_v3;
        _nw33[(new BigNumber(12)).toNumber()] = _157_v3;
        _190_v26 = _nw33;
        let _191_v27;
        _191_v27 = _dafny.Map.Empty.slice().updateUnsafe(_190_v26,_module.__default.fm1(_189_v25, false, _157_v3, _157_v3, _156_globalState));
        let _index32 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_188_v24).length));
        let _rhs30 = (_dafny.ZERO).minus(_module.__default.fm0(_156_globalState));
        let _rhs31 = _191_v27;
        let _rhs32 = (_dafny.ZERO).minus(_172_v14.f5);
        let _lhs21 = _188_v24;
        let _lhs22 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_188_v24).length));
        _lhs21[_lhs22] = _rhs30;
        _191_v27 = _rhs31;
        _189_v25 = _rhs32;
        let _source2 = _module.__default.fm2(_157_v3, new BigNumber(321), _156_globalState);
        if (_source2.is_DC1) {
          let _192___mcc_h0 = (_source2).cf1;
          let _193_cf1 = _192___mcc_h0;
          let _index33 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_188_v24).length));
          (_188_v24)[_index33] = (_172_v14.f5).minus((_172_v14.f5).multipliedBy(_module.__default.fm0(_156_globalState)));
          let _194_v28;
          let _nw34 = new _module.C0();
          _nw34.__ctor(_152_v0);
          _194_v28 = _nw34;
          let _195_v29;
          _195_v29 = _module.D2.create_DC5(_172_v14.f5, _186_i8);
          let _pat_let_tv5 = _152_v0;
          let _196_v30;
          _196_v30 = _dafny.Map.Empty.slice().updateUnsafe(_173_v15,function (_pat_let9_0) {
            return function (_197_dt__update__tmp_h0) {
              return function (_pat_let10_0) {
                return function (_198_dt__update_hcf5_h0) {
                  return _module.D2.create_DC5((_197_dt__update__tmp_h0).dtor_cf4, _198_dt__update_hcf5_h0);
                }(_pat_let10_0);
              }(_pat_let_tv5);
            }(_pat_let9_0);
          }(_195_v29));
          _196_v30 = _196_v30;
          let _199_v31;
          _199_v31 = _dafny.Seq.of(_190_v26);
          let _200_v32;
          _200_v32 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(!(_157_v3),_dafny.Seq.of(_157_v3)));
          let _201_v33;
          let _nw35 = Array((new BigNumber(25)).toNumber());
          _nw35[(_dafny.ZERO).toNumber()] = _190_v26;
          _nw35[(_dafny.ONE).toNumber()] = _190_v26;
          _nw35[(new BigNumber(2)).toNumber()] = _190_v26;
          _nw35[(new BigNumber(3)).toNumber()] = _190_v26;
          _nw35[(new BigNumber(4)).toNumber()] = _190_v26;
          _nw35[(new BigNumber(5)).toNumber()] = _190_v26;
          _nw35[(new BigNumber(6)).toNumber()] = _190_v26;
          _nw35[(new BigNumber(7)).toNumber()] = _190_v26;
          _nw35[(new BigNumber(8)).toNumber()] = _190_v26;
          _nw35[(new BigNumber(9)).toNumber()] = _190_v26;
          _nw35[(new BigNumber(10)).toNumber()] = _190_v26;
          _nw35[(new BigNumber(11)).toNumber()] = _190_v26;
          _nw35[(new BigNumber(12)).toNumber()] = _190_v26;
          _nw35[(new BigNumber(13)).toNumber()] = (_199_v31)[_module.__default.safeIndex(new BigNumber(((_200_v32)[_module.__default.safeIndex(_187_v23.f5, new BigNumber((_200_v32).length))]).length), new BigNumber((_199_v31).length))];
          _nw35[(new BigNumber(14)).toNumber()] = _190_v26;
          _nw35[(new BigNumber(15)).toNumber()] = _190_v26;
          _nw35[(new BigNumber(16)).toNumber()] = _190_v26;
          _nw35[(new BigNumber(17)).toNumber()] = _190_v26;
          _nw35[(new BigNumber(18)).toNumber()] = _190_v26;
          _nw35[(new BigNumber(19)).toNumber()] = _190_v26;
          _nw35[(new BigNumber(20)).toNumber()] = _190_v26;
          _nw35[(new BigNumber(21)).toNumber()] = _190_v26;
          _nw35[(new BigNumber(22)).toNumber()] = _190_v26;
          _nw35[(new BigNumber(23)).toNumber()] = _190_v26;
          _nw35[(new BigNumber(24)).toNumber()] = _190_v26;
          _201_v33 = _nw35;
          let _202_v34;
          _202_v34 = _dafny.Map.Empty.slice().updateUnsafe(_172_v14,_190_v26);
          let _index34 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((_201_v33).length));
          (_201_v33)[_index34] = (((_202_v34).contains(_187_v23)) ? ((_202_v34).get(_187_v23)) : (_190_v26));
          let _203_v35;
          _203_v35 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_180_v20)[_module.__default.safeIndex(new BigNumber(590), new BigNumber((_180_v20).length))]).length),_157_v3);
          let _204_v36;
          _204_v36 = _dafny.MultiSet.fromElements(_203_v35);
          let _205_v37;
          _205_v37 = _dafny.MultiSet.fromElements(new BigNumber((_204_v36).cardinality()), _186_i8, (_dafny.ZERO).minus((_188_v24)[_module.__default.safeIndex(new BigNumber(212), new BigNumber((_188_v24).length))]), _194_v28.f5, _189_v25);
          let _206_v38;
          _206_v38 = _dafny.Seq.of(_194_v28);
          let _index35 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((_201_v33).length));
          let _rhs33 = ((_205_v37).update(new BigNumber((_206_v38).length), _module.__default.abs(new BigNumber((_173_v15).length)))).IsDisjointFrom(_205_v37);
          let _rhs34 = _190_v26;
          let _lhs23 = _156_globalState;
          let _lhs24 = _201_v33;
          let _lhs25 = _module.__default.safeIndex(new BigNumber(372), new BigNumber((_201_v33).length));
          _lhs23.f1 = _rhs33;
          _lhs24[_lhs25] = _rhs34;
        } else {
          let _207___mcc_h1 = (_source2).cf0;
          let _208_cf0 = _207___mcc_h1;
          let _209_v39;
          let _nw36 = new _module.C0();
          _nw36.__ctor((_188_v24)[_module.__default.safeIndex(new BigNumber(212), new BigNumber((_188_v24).length))]);
          _209_v39 = _nw36;
          let _210_v40;
          _210_v40 = _dafny.MultiSet.fromElements((_187_v23.f5).isLessThanOrEqualTo(_172_v14.f5));
          let _211_v41;
          _211_v41 = _dafny.Seq.of(_157_v3, _157_v3);
          _210_v40 = _dafny.MultiSet.fromElements(((_157_v3) ? (_157_v3) : (false)), (_211_v41)[_module.__default.safeIndex(_module.__default.fm0(_156_globalState), new BigNumber((_211_v41).length))], (_210_v40).IsProperSubsetOf(_210_v40));
          let _212_v42;
          let _nw37 = Array((new BigNumber(8)).toNumber()).fill([]);
          _212_v42 = _nw37;
          let _213_v43;
          _213_v43 = _module.D0.create_DC1(_212_v42);
          let _214_v44;
          _214_v44 = _dafny.Map.Empty.slice().updateUnsafe(_213_v43,false);
          let _215_v45;
          _215_v45 = _module.D3.create_DC8(_214_v44);
          let _216_v46;
          _216_v46 = _dafny.Seq.of((_214_v44).Merge((_215_v45).dtor_cf10), _214_v44, _214_v44, (_214_v44).update(_module.D0.create_DC1(_212_v42), _157_v3));
          _189_v25 = new BigNumber((_216_v46).length);
          (_209_v39).f5 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_189_v25));
        }
      }
      if (_157_v3) {
        let _217_v47;
        let _init4 = function (_218_i9) {
          return false;
        };
        let _nw38 = Array((new BigNumber(25)).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw38.length); _i0_4++) {
          _nw38[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _217_v47 = _nw38;
        let _index36 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_217_v47).length));
        (_217_v47)[_index36] = _157_v3;
        let _index37 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_217_v47).length));
        (_217_v47)[_index37] = !(_157_v3);
        (_172_v14).f5 = _152_v0;
        let _index38 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((_168_v13).length));
        (_168_v13)[_index38] = ((((_180_v20)[_module.__default.safeIndex(new BigNumber(590), new BigNumber((_180_v20).length))]).contains((_dafny.ZERO).minus(_172_v14.f5))) ? (((_180_v20)[_module.__default.safeIndex(new BigNumber(590), new BigNumber((_180_v20).length))]).get((_dafny.ZERO).minus(_172_v14.f5))) : (_152_v0));
        let _219_v48;
        _219_v48 = _dafny.Seq.of(_168_v13, _168_v13, _168_v13, _168_v13);
        let _220_v50;
        _220_v50 = _dafny.Set.fromElements(new BigNumber((_153_v1).length), _152_v0);
        let _221_v51;
        _221_v51 = _dafny.Map.Empty.slice().updateUnsafe(_172_v14,_168_v13);
        let _index39 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_217_v47).length));
        let _index40 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((_168_v13).length));
        let _rhs35 = ((_220_v50).Union(_module.__default.fm3((_217_v47)[_module.__default.safeIndex(new BigNumber(360), new BigNumber((_217_v47).length))], !(_157_v3), _156_globalState))).IsSubsetOf(function () {
          let _coll6 = new _dafny.Set();
          for (const _compr_6 of _dafny.IntegerRange(new BigNumber(803), new BigNumber(238))) {
            let _222_v49 = _compr_6;
            if (((new BigNumber(803)).isLessThanOrEqualTo(_222_v49)) && ((_222_v49).isLessThan(new BigNumber(238)))) {
              _coll6.add(_module.__default.safeDivisionInt(_222_v49, _172_v14.f5));
            }
          }
          return _coll6;
        }());
        let _rhs36 = _module.__default.fm0(_156_globalState);
        let _rhs37 = _dafny.Seq.Concat(_dafny.Seq.Concat(_219_v48, (_dafny.Seq.of((((_221_v51).contains(_172_v14)) ? ((_221_v51).get(_172_v14)) : (_168_v13))))), _dafny.Seq.of(_168_v13));
        let _rhs38 = _172_v14;
        let _lhs26 = _217_v47;
        let _lhs27 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_217_v47).length));
        let _lhs28 = _168_v13;
        let _lhs29 = _module.__default.safeIndex(new BigNumber(587), new BigNumber((_168_v13).length));
        _lhs26[_lhs27] = _rhs35;
        _lhs28[_lhs29] = _rhs36;
        _219_v48 = _rhs37;
        _172_v14 = _rhs38;
        (_156_globalState).f3 = _173_v15;
        let _223_v52;
        let _224_v53;
        let _out6;
        let _out7;
        let _outcollector3 = _module.__default.m0(_156_globalState);
        _out6 = _outcollector3[0];
        _out7 = _outcollector3[1];
        _223_v52 = _out6;
        _224_v53 = _out7;
      } else {
        _152_v0 = (_dafny.ZERO).minus(_172_v14.f5);
        let _225_v54;
        _225_v54 = _dafny.MultiSet.fromElements(_152_v0);
        if (_module.__default.fm1(_172_v14.f5, !(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(930)), ((_226_v14) => function (_227_i10) {
          return _226_v14.f5;
        })(_172_v14))).length)).isEqualTo(_152_v0), (new BigNumber((_225_v54).cardinality())).isLessThan(_152_v0), _157_v3, _156_globalState)) {
          let _228_v55;
          let _nw39 = Array((new BigNumber(12)).toNumber()).fill(false);
          _228_v55 = _nw39;
          let _index41 = _module.__default.safeIndex(new BigNumber(358), new BigNumber((_228_v55).length));
          (_228_v55)[_index41] = !(((_dafny.ZERO).minus(new BigNumber((_173_v15).length))).isLessThanOrEqualTo(_172_v14.f5));
          let _index42 = _module.__default.safeIndex(new BigNumber(358), new BigNumber((_228_v55).length));
          (_228_v55)[_index42] = _157_v3;
          (_156_globalState).f1 = _157_v3;
          let _229_v56;
          _229_v56 = _dafny.Seq.of((_228_v55)[_module.__default.safeIndex(new BigNumber(358), new BigNumber((_228_v55).length))], (_228_v55)[_module.__default.safeIndex(new BigNumber(358), new BigNumber((_228_v55).length))]);
          let _230_v57;
          _230_v57 = _dafny.Seq.of(_229_v56, _229_v56);
          let _231_v58;
          _231_v58 = _dafny.Set.fromElements(_152_v0, _module.__default.safeModuloInt(_152_v0, _172_v14.f5), new BigNumber((_dafny.Seq.update(_173_v15, _module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements(_module.__default.fm1(new BigNumber(((_230_v57)[_module.__default.safeIndex(_172_v14.f5, new BigNumber((_230_v57).length))]).length), (_228_v55)[_module.__default.safeIndex(new BigNumber(358), new BigNumber((_228_v55).length))], true, _157_v3, _156_globalState), (_228_v55)[_module.__default.safeIndex(new BigNumber(358), new BigNumber((_228_v55).length))])).cardinality()), new BigNumber((_173_v15).length)), _module.__default.fm4((_228_v55)[_module.__default.safeIndex(new BigNumber(358), new BigNumber((_228_v55).length))], _156_globalState))).length), new BigNumber(321));
          _231_v58 = ((_231_v58).Union(function () {
            let _coll7 = new _dafny.Set();
            for (const _compr_7 of (_153_v1).Elements) {
              let _232_v59 = _compr_7;
              if (_dafny.Seq.contains(_153_v1, _232_v59)) {
                _coll7.add(_module.__default.safeModuloInt(_232_v59, new BigNumber(-437)));
              }
            }
            return _coll7;
          }())).Intersect(_dafny.Set.fromElements(_152_v0));
          _179_v19 = ((true) ? (_179_v19) : (_179_v19));
          let _233_v60;
          let _init5 = ((_234_v19) => function (_235_i11) {
            return _234_v19;
          })(_179_v19);
          let _nw40 = Array((new BigNumber(16)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw40.length); _i0_5++) {
            _nw40[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _233_v60 = _nw40;
          let _236_v61;
          _236_v61 = _dafny.Map.Empty.slice().updateUnsafe(_152_v0,_233_v60);
          _233_v60 = (((_236_v61).contains(_152_v0)) ? ((_236_v61).get(_152_v0)) : ((((_228_v55)[_module.__default.safeIndex(new BigNumber(358), new BigNumber((_228_v55).length))]) ? (_233_v60) : (_233_v60))));
        } else {
          let _237_v62;
          _237_v62 = _module.D2.create_DC5(_172_v14.f5, _172_v14.f5);
          _237_v62 = _237_v62;
          let _238_v63;
          _238_v63 = _dafny.MultiSet.fromElements(_157_v3);
          let _239_v64;
          _239_v64 = _dafny.Map.Empty.slice().updateUnsafe(_172_v14,_172_v14);
          let _240_v65;
          _240_v65 = _dafny.Seq.of(_157_v3);
          let _nw41 = Array((new BigNumber(8)).toNumber());
          _nw41[(_dafny.ZERO).toNumber()] = _172_v14.f5;
          _nw41[(_dafny.ONE).toNumber()] = _172_v14.f5;
          _nw41[(new BigNumber(2)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_238_v63).cardinality())), new BigNumber((_173_v15).length));
          _nw41[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt(_172_v14.f5, _172_v14.f5);
          _nw41[(new BigNumber(4)).toNumber()] = new BigNumber((_239_v64).length);
          _nw41[(new BigNumber(5)).toNumber()] = _module.__default.fm0(_156_globalState);
          _nw41[(new BigNumber(6)).toNumber()] = new BigNumber((_240_v65).length);
          _nw41[(new BigNumber(7)).toNumber()] = _172_v14.f5;
          _168_v13 = _nw41;
          _152_v0 = (_172_v14.f5).minus(_172_v14.f5);
          (_172_v14).f5 = _152_v0;
          let _241_v66;
          let _242_v67;
          let _out8;
          let _out9;
          let _outcollector4 = _module.__default.m0(_156_globalState);
          _out8 = _outcollector4[0];
          _out9 = _outcollector4[1];
          _241_v66 = _out8;
          _242_v67 = _out9;
        }
        (_156_globalState).f1 = _157_v3;
        let _243_v68;
        _243_v68 = _module.D0.create_DC0(_153_v1);
        let _pat_let_tv6 = _153_v1;
        _243_v68 = function (_pat_let11_0) {
          return function (_244_dt__update__tmp_h1) {
            return function (_pat_let12_0) {
              return function (_245_dt__update_hcf0_h0) {
                return _module.D0.create_DC0(_245_dt__update_hcf0_h0);
              }(_pat_let12_0);
            }(_pat_let_tv6);
          }(_pat_let11_0);
        }(_243_v68);
        let _246_v69;
        _246_v69 = _dafny.Set.fromElements(_157_v3);
        let _247_v70;
        _247_v70 = _dafny.Set.fromElements(_246_v69, _246_v69);
        _247_v70 = (_247_v70).Union(_247_v70);
      }
      let _248_v71;
      let _init6 = ((_249_v3) => function (_250_i12) {
        return _249_v3;
      })(_157_v3);
      let _nw42 = Array((_dafny.ONE).toNumber());
      for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw42.length); _i0_6++) {
        _nw42[_i0_6] = _init6(new BigNumber(_i0_6));
      }
      _248_v71 = _nw42;
      _248_v71 = _248_v71;
      let _251_v72;
      _251_v72 = _dafny.Map.Empty.slice().updateUnsafe(_157_v3,_248_v71);
      _251_v72 = (_251_v72).update(false, _248_v71);
      (_172_v14).f5 = _module.__default.safeModuloInt(_172_v14.f5, _module.__default.safeModuloInt((_153_v1)[_module.__default.safeIndex(_172_v14.f5, new BigNumber((_153_v1).length))], _module.__default.fm0(_156_globalState)));
      let _252_v73;
      let _253_v74;
      let _out10;
      let _out11;
      let _outcollector5 = _module.__default.m0(_156_globalState);
      _out10 = _outcollector5[0];
      _out11 = _outcollector5[1];
      _252_v73 = _out10;
      _253_v74 = _out11;
      (_156_globalState).f1 = _157_v3;
      let _hi1 = _172_v14.f5;
      for (let _254_i13 = _172_v14.f5; _254_i13.isLessThan(_hi1); _254_i13 = _254_i13.plus(_dafny.ONE)) {
        let _255_v75;
        let _nw43 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Map.Empty);
        _255_v75 = _nw43;
        let _256_v76;
        let _nw44 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.of());
        _256_v76 = _nw44;
        let _257_v77;
        _257_v77 = _dafny.Map.Empty.slice().updateUnsafe(_256_v76,_172_v14.f5);
        let _index43 = _module.__default.safeIndex(new BigNumber(843), new BigNumber((_255_v75).length));
        (_255_v75)[_index43] = ((_157_v3) ? (_257_v77) : (_257_v77));
        let _index44 = _module.__default.safeIndex(new BigNumber(843), new BigNumber((_255_v75).length));
        (_255_v75)[_index44] = _257_v77;
        _157_v3 = ((_153_v1)[_module.__default.safeIndex(_152_v0, new BigNumber((_153_v1).length))]).isLessThan(_172_v14.f5);
        _152_v0 = (new BigNumber((_173_v15).length)).plus(_172_v14.f5);
        (_172_v14).f5 = (_152_v0).multipliedBy(_254_i13);
      }
      process.stdout.write(_dafny.toString(_152_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_153_v1, _dafny.Seq.of(new BigNumber(-560), new BigNumber(-560)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v2)[_dafny.ZERO]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v2)[_dafny.ONE]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v2)[new BigNumber(2)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v2)[new BigNumber(3)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v2)[new BigNumber(4)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v2)[new BigNumber(5)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v2)[new BigNumber(6)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v2)[new BigNumber(7)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v2)[new BigNumber(8)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v2)[new BigNumber(9)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v2)[new BigNumber(10)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v2)[new BigNumber(11)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v2)[new BigNumber(12)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v2)[new BigNumber(13)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v2)[new BigNumber(14)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v2)[new BigNumber(15)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v2)[new BigNumber(16)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v2)[new BigNumber(17)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_154_v2)[new BigNumber(18)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_156_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_156_globalState.f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_156_globalState).f2, _dafny.Seq.of(new BigNumber(-560), new BigNumber(-560)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_156_globalState.f3).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_156_globalState).f4)[_dafny.ZERO]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_156_globalState).f4)[_dafny.ONE]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_156_globalState).f4)[new BigNumber(2)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_156_globalState).f4)[new BigNumber(3)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_156_globalState).f4)[new BigNumber(4)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_156_globalState).f4)[new BigNumber(5)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_156_globalState).f4)[new BigNumber(6)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_156_globalState).f4)[new BigNumber(7)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_156_globalState).f4)[new BigNumber(8)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_156_globalState).f4)[new BigNumber(9)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_156_globalState).f4)[new BigNumber(10)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_156_globalState).f4)[new BigNumber(11)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_156_globalState).f4)[new BigNumber(12)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_156_globalState).f4)[new BigNumber(13)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_156_globalState).f4)[new BigNumber(14)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_156_globalState).f4)[new BigNumber(15)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_156_globalState).f4)[new BigNumber(16)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_156_globalState).f4)[new BigNumber(17)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_156_globalState).f4)[new BigNumber(18)]).equals(_dafny.MultiSet.fromElements(new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_157_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_158_i1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v13)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v13)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v13)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v13)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v13)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_172_v14.f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_173_v15).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_174_i5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_179_v19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_180_v20)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(9),new BigNumber(9)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_180_v20)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(9),new BigNumber(9)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_180_v20)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(9),new BigNumber(9)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_180_v20)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(9),new BigNumber(9)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_180_v20)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(9),new BigNumber(9)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_180_v20)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(9),new BigNumber(9)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_180_v20)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(9),new BigNumber(9)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_180_v20)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(9),new BigNumber(9)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_180_v20)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(9),new BigNumber(9)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_180_v20)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(9),new BigNumber(9)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_180_v20)[new BigNumber(10)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(9),new BigNumber(9)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_180_v20)[new BigNumber(11)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(9),new BigNumber(9)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_180_v20)[new BigNumber(12)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(9),new BigNumber(9)).updateUnsafe(new BigNumber(-9),new BigNumber(9)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_180_v20)[new BigNumber(13)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(9),new BigNumber(9)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_180_v20)[new BigNumber(14)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(9),new BigNumber(9)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_180_v20)[new BigNumber(15)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(9),new BigNumber(9)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_180_v20)[new BigNumber(16)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(9),new BigNumber(9)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_184_v21).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(9)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_185_v22).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(9),_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_248_v71)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_251_v72).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_253_v74));
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
    get is_DC1() { return this.$tag === 0; }
    get is_DC0() { return this.$tag === 1; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ")";
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
        return other.$tag === 0 && this.cf1 === other.cf1;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1([]);
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
    static create_DC3() {
      let $dt = new D1(0);
      return $dt;
    }
    static create_DC2(cf2) {
      let $dt = new D1(1);
      $dt.cf2 = cf2;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get dtor_cf2() { return this.cf2; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC3";
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
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf2 === other.cf2;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC3();
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
    static create_DC5(cf4, cf5) {
      let $dt = new D2(0);
      $dt.cf4 = cf4;
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC6() {
      let $dt = new D2(1);
      return $dt;
    }
    static create_DC7(cf6, cf7, cf8, cf9) {
      let $dt = new D2(2);
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC4(cf3) {
      let $dt = new D2(3);
      $dt.cf3 = cf3;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get is_DC7() { return this.$tag === 2; }
    get is_DC4() { return this.$tag === 3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf3() { return this.cf3; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf4) + ", " + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC6";
      } else if (this.$tag === 2) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 3) {
        return "D2.DC4" + "(" + _dafny.toString(this.cf3) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf4, other.cf4) && _dafny.areEqual(this.cf5, other.cf5);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf6 === other.cf6 && _dafny.areEqual(this.cf7, other.cf7) && this.cf8 === other.cf8 && _dafny.areEqual(this.cf9, other.cf9);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf3, other.cf3);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC5(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC9(cf11) {
      let $dt = new D3(0);
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC10(cf12, cf13, cf14, cf15, cf16) {
      let $dt = new D3(1);
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC8(cf10) {
      let $dt = new D3(2);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get is_DC8() { return this.$tag === 2; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf10) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf11, other.cf11);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf12 === other.cf12 && this.cf13 === other.cf13 && _dafny.areEqual(this.cf14, other.cf14) && _dafny.areEqual(this.cf15, other.cf15) && _dafny.areEqual(this.cf16, other.cf16);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC9(_dafny.MultiSet.Empty);
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
    get is_DC11() { return this.$tag === 0; }
    get dtor_cf17() { return this.cf17; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf17) + ")";
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
      return _dafny.Seq.of();
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
    static create_DC13(cf19, cf20, cf21, cf22) {
      let $dt = new D5(0);
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC12(cf18) {
      let $dt = new D5(1);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC12" + "(" + _dafny.toString(this.cf18) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf19, other.cf19) && _dafny.areEqual(this.cf20, other.cf20) && _dafny.areEqual(this.cf21, other.cf21) && _dafny.areEqual(this.cf22, other.cf22);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf18, other.cf18);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC13(_module.D0.Default(), _dafny.ZERO, _dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D5.Default();
        }
      };
    }
  }

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f1 = false;
      this.f3 = _dafny.Seq.UnicodeFromString("");
      this._f0 = _dafny.ZERO;
      this._f2 = _dafny.Seq.of();
      this._f4 = [];
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4) {
      let _this = this;
      (_this)._f0 = f0;
      (_this).f1 = f1;
      (_this)._f2 = f2;
      (_this).f3 = f3;
      (_this)._f4 = f4;
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
    get f4() {
      let _this = this;
      return _this._f4;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this.f5 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f5) {
      let _this = this;
      (_this).f5 = f5;
      return;
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
