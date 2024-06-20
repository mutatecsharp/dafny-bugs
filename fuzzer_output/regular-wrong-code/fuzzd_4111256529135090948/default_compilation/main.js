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
      if (true) {
        return _module.__default.safeModuloInt(new BigNumber(-138), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length));
      } else {
        return new BigNumber(541);
      }
    };
    static fm1(p0, p1, globalState) {
      return true;
    };
    static fm2(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ko"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(296)), function (_0_i0) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      })), _dafny.Seq.Create(_module.__default.abs(new BigNumber(78)), function (_1_i1) {
        return new _dafny.CodePoint('i'.codePointAt(0));
      }));
    };
    static fm3(p0, globalState) {
      return function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of (_dafny.Seq.of(new _dafny.CodePoint('s'.codePointAt(0)))).Elements) {
          let _2_v0 = _compr_0;
          if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('s'.codePointAt(0))), _2_v0)) {
            _coll0.push([_2_v0,_dafny.areEqual(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("ac")), _dafny.Seq.of(_dafny.Seq.UnicodeFromString("n"), _dafny.Seq.UnicodeFromString("dubvfe"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(152)), function (_3_i0) {
              return new _dafny.CodePoint('w'.codePointAt(0));
            }), _dafny.Seq.UnicodeFromString("pt"), _dafny.Seq.UnicodeFromString("k")))]);
          }
        }
        return _coll0;
      }();
    };
    static fm4(p0, p1, p2, p3, globalState) {
      return _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(false, false), _dafny.Seq.of(false)), _dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(true, false, false, true, false))));
    };
    static fm5(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(_dafny.Set.fromElements(!(!(false))), _dafny.Set.fromElements(true, false, true, true, (_module.D9.create_DC25(true, _dafny.Set.fromElements(new _dafny.CodePoint('d'.codePointAt(0))))).dtor_cf41), _dafny.Set.fromElements(false), _dafny.Set.fromElements(true))).length),!(_dafny.MultiSet.fromElements(!(true), false)).equals(_dafny.MultiSet.FromArray(_dafny.Seq.of(true))));
    };
    static fm6(p0, globalState) {
      return new _dafny.CodePoint('b'.codePointAt(0));
    };
    static fm7(p0, p1, p2, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,!(true)))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(true),!(false))));
    };
    static fm9(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((_dafny.MultiSet.fromElements(false, false, false, false, true)).Intersect(_dafny.MultiSet.fromElements(false, false)),new _dafny.CodePoint('l'.codePointAt(0)));
    };
    static fm13(p0, p1, p2, p3, globalState) {
      if (false) {
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(671)), function (_4_i0) {
          return new _dafny.CodePoint('p'.codePointAt(0));
        })).length),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(273),new BigNumber((_dafny.MultiSet.fromElements(!(false), false)).cardinality())))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-33),function () {
          let _coll1 = new _dafny.Map();
          for (const _compr_1 of _dafny.IntegerRange(new BigNumber(228), new BigNumber(940))) {
            let _5_v0 = _compr_1;
            if (((new BigNumber(228)).isLessThanOrEqualTo(_5_v0)) && ((_5_v0).isLessThan(new BigNumber(940)))) {
              _coll1.push([_module.__default.safeModuloInt(_5_v0, new BigNumber(160)),new BigNumber(791)]);
            }
          }
          return _coll1;
        }()));
      } else {
        return _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(-777)),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length),new BigNumber(-658)));
      }
    };
    static fm14(p0, p1, p2, globalState) {
      return (_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(459)), _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-480)))).Union(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Set.fromElements(true)).length))), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(207))))));
    };
    static fm15(p0, p1, globalState) {
      return (((true) ? (_dafny.MultiSet.fromElements(_dafny.Seq.of(true))) : (_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.of(true)))))).Difference(_dafny.MultiSet.fromElements(_dafny.Seq.of(false)));
    };
    static fm16(p0, p1, p2, globalState) {
      if (!(new BigNumber(988)).isEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length))) {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(478),new BigNumber(394));
      } else {
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(105),new BigNumber(451))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(251),new BigNumber(-748)));
      }
    };
    static fm17(p0, p1, globalState) {
      if ((true) || (false)) {
        return _dafny.MultiSet.fromElements(new _dafny.CodePoint('s'.codePointAt(0)));
      } else {
        return _dafny.MultiSet.fromElements(new _dafny.CodePoint('c'.codePointAt(0)), new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)));
      }
    };
    static fm19(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('q'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0)))).cardinality())), _dafny.Seq.Concat(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("kwrvsluv"), _dafny.Seq.UnicodeFromString("mtfd"))).length)), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((function () {
        let _coll2 = new _dafny.Set();
        for (const _compr_2 of (_dafny.Seq.of(_module.D0.create_DC2(_module.D0.create_DC0(_dafny.Seq.Create(_module.__default.abs(new BigNumber(633)), function (_6_i0) {
  return new _dafny.CodePoint('x'.codePointAt(0));
}))))).Elements) {
          let _7_v0 = _compr_2;
          if (_dafny.Seq.contains(_dafny.Seq.of(_module.D0.create_DC2(_module.D0.create_DC0(_dafny.Seq.Create(_module.__default.abs(new BigNumber(633)), function (_6_i0) {
  return new _dafny.CodePoint('x'.codePointAt(0));
})))), _7_v0)) {
            _coll2.add(_7_v0);
          }
        }
        return _coll2;
      }()).length)))).cardinality()), new BigNumber((_dafny.Set.fromElements(false)).length)), _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("yrcmuol")).length))));
    };
    static fm21(p0, globalState) {
      return (_dafny.Set.fromElements(new BigNumber(175), new BigNumber(652), new BigNumber(-976), new BigNumber(896))).Union(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("b")).length)));
    };
    static fm22(p0, globalState) {
      return _module.D3.create_DC12(_dafny.Seq.Concat(_dafny.Seq.of(!(true), false), _dafny.Seq.of(!(true))));
    };
    static fm23(globalState) {
      return _module.D1.create_DC5((_dafny.ZERO).minus((new BigNumber(182)).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(357)), function (_8_i0) {
  return _dafny.Seq.of(true, !(!(false)));
})).length))), !(new BigNumber(-715)).isEqualTo(new BigNumber((_dafny.Seq.UnicodeFromString("lucmkk")).length)), new BigNumber(-424), new BigNumber(((function () {
  let _coll3 = new _dafny.Set();
  for (const _compr_3 of _dafny.IntegerRange(new BigNumber(673), new BigNumber(245))) {
    let _9_v0 = _compr_3;
    if (((new BigNumber(673)).isLessThanOrEqualTo(_9_v0)) && ((_9_v0).isLessThan(new BigNumber(245)))) {
      _coll3.add((_9_v0).minus(new BigNumber(937)));
    }
  }
  return _coll3;
}()).Union(_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,false))).length))).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(329)), function (_10_i1) {
  return new _dafny.CodePoint('l'.codePointAt(0));
})).length))).cardinality())))).length));
    };
    static fm24(p0, p1, globalState) {
      return _module.D9.create_DC26((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("yglvl"), _dafny.Seq.UnicodeFromString("emwcyv"))).length)));
    };
    static fm25(p0, p1, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(new BigNumber(-269)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(773)));
    };
    static fm26(p0, globalState) {
      return _module.D0.create_DC2(_module.D0.create_DC1(new BigNumber(873), !(true), new BigNumber((_dafny.Seq.of(new BigNumber(380))).length)));
    };
    static fm27(globalState) {
      if (!((_dafny.MultiSet.FromArray(_dafny.Seq.of(!(false), false, !(true)))).IsProperSubsetOf(_dafny.MultiSet.fromElements(false, false)))) {
        return _dafny.Map.Empty.slice().updateUnsafe(!(false),_dafny.Seq.UnicodeFromString("ivfd"));
      } else {
        return (_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.UnicodeFromString("jbqnyc"))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("jjxglfc")));
      }
    };
    static fm28(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false),new BigNumber(868))).Merge((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true, false),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(false, false)).length)))).Merge(function () {
        let _coll4 = new _dafny.Map();
        for (const _compr_4 of (_dafny.Set.fromElements(_dafny.Seq.of(true))).Elements) {
          let _11_v0 = _compr_4;
          if ((_dafny.Set.fromElements(_dafny.Seq.of(true))).contains(_11_v0)) {
            _coll4.push([_11_v0,new BigNumber(200)]);
          }
        }
        return _coll4;
      }()));
    };
    static fm29(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(171)), function (_12_i0) {
        return _dafny.MultiSet.fromElements(false, false, true);
      }), _dafny.Seq.of(_dafny.MultiSet.fromElements(true))), _dafny.Seq.of(_dafny.MultiSet.fromElements(!(false)), _dafny.MultiSet.FromArray(_dafny.Seq.of(true, false)), _dafny.MultiSet.fromElements(true), _dafny.MultiSet.fromElements(true), _dafny.MultiSet.fromElements(true, false)));
    };
    static fm30(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(false)), _dafny.Seq.Concat(_dafny.Seq.of(!(false)), _dafny.Seq.of(false)));
    };
    static fm31(p0, globalState) {
      if (!(!(true))) {
        return _module.D12.create_DC36(true, false, new BigNumber(125), new BigNumber(710), _dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.Create(_module.__default.abs(new BigNumber(674)), function (_13_i0) {
  return new BigNumber(-438);
})));
      } else if (!(true)) {
        return _module.D12.create_DC36(false, true, new BigNumber(770), new BigNumber((_dafny.Seq.of(new BigNumber(988))).length), _dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.of(new BigNumber(735), new BigNumber((_dafny.Seq.UnicodeFromString("p")).length), new BigNumber(-224))));
      } else {
        return _module.D12.create_DC36(false, !(false), new BigNumber(977), new BigNumber((_dafny.Seq.of(false, false)).length), _dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(false)).length), new BigNumber(-67), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true)).length)))).length))).length))));
      }
    };
    static fm32(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_module.D11.create_DC31(new BigNumber(-732))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(681)), function (_14_i0) {
        return _module.D11.create_DC31(new BigNumber(582));
      })), _dafny.Seq.of(_module.D11.create_DC31(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of((_module.D0.create_DC1((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(true, false)).length)), true, new BigNumber(333))).dtor_cf2, false))).cardinality())), _module.D11.create_DC31(new BigNumber((_dafny.Set.fromElements(!(true))).length)), _module.D11.create_DC31(new BigNumber((_dafny.Seq.of(true)).length)), _module.D11.create_DC31(new BigNumber((_dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('l'.codePointAt(0)))).length))))).length))));
    };
    static fm33(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-70)), function (_15_i0) {
        return _module.D11.create_DC31(new BigNumber(807));
      }),_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("aguorpcjs")).length))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_module.D11.create_DC31(new BigNumber((_dafny.Seq.UnicodeFromString("habj")).length))),_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("nv")).length)), new BigNumber(200))))).Merge((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_module.D11.create_DC31(new BigNumber((_dafny.Seq.UnicodeFromString("kdusdcx")).length))),_dafny.Seq.of(new BigNumber(582)))).Merge(function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_module.D11.create_DC31(new BigNumber((_dafny.Seq.UnicodeFromString("to")).length))),new BigNumber((_dafny.Seq.UnicodeFromString("bffhds")).length))).Keys.Elements) {
          let _16_v0 = _compr_5;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_module.D11.create_DC31(new BigNumber((_dafny.Seq.UnicodeFromString("to")).length))),new BigNumber((_dafny.Seq.UnicodeFromString("bffhds")).length))).contains(_16_v0)) {
            _coll5.push([_16_v0,_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-691))).length))]);
          }
        }
        return _coll5;
      }()));
    };
    static fm34(p0, globalState) {
      return _module.D0.create_DC0(_dafny.Seq.UnicodeFromString("mlwjdjdmf"));
    };
    static fm35(p0, p1, p2, globalState) {
      return ((_dafny.Set.fromElements(!(!(true)), false)).Union(_dafny.Set.fromElements(true))).Union(_dafny.Set.fromElements(false, true));
    };
    static m0(p0, p1, p2, globalState) {
      let r0 = _module.D0.Default();
      let r1 = _dafny.Map.Empty;
      let r2 = _dafny.ZERO;
      r2 = p2;
      let _17_v0;
      let _nw0 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
      _17_v0 = _nw0;
      let _18_v1;
      _18_v1 = false;
      let _19_v2;
      _19_v2 = _dafny.Map.Empty.slice().updateUnsafe(_17_v0,_18_v1);
      let _20_v3;
      _20_v3 = _dafny.Seq.of(p2);
      let _21_v4;
      _21_v4 = _module.D8.create_DC20(true, (_20_v3)[_module.__default.safeIndex(p2, new BigNumber((_20_v3).length))]);
      let _22_v5;
      let _nw1 = new _module.C5();
      _nw1.__ctor(new BigNumber((_19_v2).length), _module.__default.safeModuloInt(p2, p2), (_21_v4).dtor_cf34);
      _22_v5 = _nw1;
      let _23_v6;
      _23_v6 = _dafny.Seq.of(p1);
      let _24_v7;
      _24_v7 = new _dafny.CodePoint('x'.codePointAt(0));
      let _25_v8;
      let _nw2 = new _module.C4();
      _nw2.__ctor(_17_v0, (_22_v5).f25);
      _25_v8 = _nw2;
      let _26_v9;
      _26_v9 = _dafny.Seq.of(_25_v8, _25_v8);
      let _27_v10;
      _27_v10 = _dafny.MultiSet.fromElements((_26_v9)[_module.__default.safeIndex((_22_v5).f25, new BigNumber((_26_v9).length))]);
      let _28_v11;
      _28_v11 = _dafny.Seq.of(_module.__default.fm1(new BigNumber(895), _18_v1, globalState));
      let _29_v12;
      _29_v12 = _dafny.Set.fromElements(p2);
      let _30_v13;
      _30_v13 = _dafny.Map.Empty.slice().updateUnsafe((_22_v5).f25,(_22_v5).f26);
      let _rhs0 = (new BigNumber((_dafny.Seq.update(p1, _module.__default.safeIndex((_module.D8.create_DC21(_18_v1, _18_v1, (_dafny.ZERO).minus(new BigNumber(((_23_v6)[_module.__default.safeIndex(p2, new BigNumber((_23_v6).length))]).length)))).dtor_cf37, new BigNumber((p1).length)), _24_v7)).length)).isLessThan(_module.__default.safeModuloInt((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_27_v10).cardinality()))), new BigNumber((_28_v11).length)));
      let _rhs1 = (_22_v5).f25;
      let _rhs2 = (_29_v12).IsSubsetOf(_29_v12);
      let _rhs3 = _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_30_v13).length)), p2);
      let _lhs0 = globalState;
      let _lhs1 = globalState;
      _lhs0.f7 = _rhs0;
      r2 = _rhs1;
      _18_v1 = _rhs2;
      _lhs1.f8 = _rhs3;
      let _31_v14;
      _31_v14 = _dafny.Map.Empty.slice().updateUnsafe(_18_v1,_18_v1);
      let _32_v15;
      let _nw3 = new _module.C0();
      _nw3.__ctor();
      _32_v15 = _nw3;
      let _33_v16;
      _33_v16 = _dafny.Map.Empty.slice().updateUnsafe(_32_v15,new BigNumber(797));
      let _34_v17;
      _34_v17 = _dafny.Map.Empty.slice().updateUnsafe(_33_v16,_31_v14);
      let _35_v18;
      let _nw4 = Array((new BigNumber(7)).toNumber());
      _nw4[(_dafny.ZERO).toNumber()] = _31_v14;
      _nw4[(_dafny.ONE).toNumber()] = _31_v14;
      _nw4[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(!(_18_v1),_18_v1);
      _nw4[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(!(_18_v1),_18_v1);
      _nw4[(new BigNumber(4)).toNumber()] = (((_34_v17).contains(_dafny.Map.Empty.slice().updateUnsafe(_32_v15,(_22_v5).f26))) ? ((_34_v17).get(_dafny.Map.Empty.slice().updateUnsafe(_32_v15,(_22_v5).f26))) : (_31_v14));
      _nw4[(new BigNumber(5)).toNumber()] = _31_v14;
      _nw4[(new BigNumber(6)).toNumber()] = _31_v14;
      _35_v18 = _nw4;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_35_v18).length))) {
        let _36_i0 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_36_i0)) && ((_36_i0).isLessThan(new BigNumber((_35_v18).length))))) {
          (_35_v18)[(_36_i0)] = (_dafny.Map.Empty.slice().updateUnsafe(false,_18_v1)).Merge((_module.D19.create_DC52(_31_v14)).dtor_cf85);
        }
      }
      let _hi0 = p2;
      for (let _37_i1 = (_22_v5).f25; _37_i1.isLessThan(_hi0); _37_i1 = _37_i1.plus(_dafny.ONE)) {
        let _38_v19;
        let _nw5 = Array((new BigNumber(16)).toNumber()).fill(false);
        _38_v19 = _nw5;
        _38_v19 = _38_v19;
        r2 = _module.__default.fm0((_22_v5).f25, _18_v1, globalState);
        (globalState).f7 = _18_v1;
        (globalState).f6 = p1;
      }
      _18_v1 = _18_v1;
      let _39_v20;
      let _nw6 = new _module.C2();
      _nw6.__ctor(_30_v13, _18_v1, _18_v1, _20_v3, p2);
      _39_v20 = _nw6;
      let _40_v21;
      let _nw7 = Array((_dafny.ONE).toNumber());
      _nw7[(_dafny.ZERO).toNumber()] = _39_v20;
      _40_v21 = _nw7;
      let _41_v22;
      _41_v22 = _dafny.MultiSet.fromElements(_40_v21, _40_v21, _40_v21, _40_v21);
      let _42_v23;
      _42_v23 = _module.D20.create_DC55(_41_v22);
      let _43_v24;
      _43_v24 = _module.D0.create_DC1(_module.__default.fm0(p2, false, globalState), (_41_v22).IsSubsetOf((_42_v23).dtor_cf89), (_22_v5).f26);
      r0 = _43_v24;
      let _44_v25;
      _44_v25 = _dafny.Map.Empty.slice().updateUnsafe(_18_v1,_20_v3);
      let _45_v26;
      _45_v26 = _module.D12.create_DC36(!((_39_v20).f32), _18_v1, new BigNumber((p1).length), p2, _44_v25);
      let _pat_let_tv0 = _39_v20;
      let _pat_let_tv1 = _22_v5;
      let _pat_let_tv2 = _39_v20;
      let _pat_let_tv3 = _24_v7;
      let _pat_let_tv4 = _24_v7;
      let _pat_let_tv5 = _39_v20;
      let _pat_let_tv6 = _22_v5;
      let _pat_let_tv7 = p1;
      let _pat_let_tv8 = _18_v1;
      let _pat_let_tv9 = _22_v5;
      let _pat_let_tv10 = _22_v5;
      let _pat_let_tv11 = _18_v1;
      r1 = function (_source0) {
        if (_source0.is_DC35) {
          return (_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(((_pat_let_tv0).f31).length)),true)).update((_pat_let_tv1).f26, (_pat_let_tv2).f32);
        } else if (_source0.is_DC36) {
          let _46___mcc_h0 = (_source0).cf56;
          let _47___mcc_h1 = (_source0).cf57;
          let _48___mcc_h2 = (_source0).cf58;
          let _49___mcc_h3 = (_source0).cf59;
          let _50___mcc_h4 = (_source0).cf60;
          let _51_cf60 = _50___mcc_h4;
          let _52_cf59 = _49___mcc_h3;
          let _53_cf58 = _48___mcc_h2;
          let _54_cf57 = _47___mcc_h1;
          let _55_cf56 = _46___mcc_h0;
          return (function () {
            let _coll6 = new _dafny.Map();
            for (const _compr_6 of (_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(842)), ((_56_v7) => function (_57_i2) {
              return _56_v7;
            })(_pat_let_tv3))).length), new BigNumber((_dafny.Seq.of(new BigNumber(-950))).length))).Elements) {
              let _58_v27 = _compr_6;
              if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(842)), ((_59_v7) => function (_57_i2) {
                return _59_v7;
              })(_pat_let_tv4))).length), new BigNumber((_dafny.Seq.of(new BigNumber(-950))).length)), _58_v27)) {
                _coll6.push([(_58_v27).multipliedBy((_pat_let_tv6).f26),(_pat_let_tv5).f32]);
              }
            }
            return _coll6;
          }()).Merge((function () {
            let _coll7 = new _dafny.Map();
            for (const _compr_7 of _dafny.IntegerRange(new BigNumber(-214), new BigNumber(460))) {
              let _60_v28 = _compr_7;
              if (((new BigNumber(-214)).isLessThanOrEqualTo(_60_v28)) && ((_60_v28).isLessThan(new BigNumber(460)))) {
                _coll7.push([_module.__default.safeDivisionInt(_60_v28, _53_cf58),_55_cf56]);
              }
            }
            return _coll7;
          }()).update(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_pat_let_tv7).length))).length), _54_cf57));
        } else if (_source0.is_DC34) {
          let _61___mcc_h5 = (_source0).cf55;
          let _62_cf55 = _61___mcc_h5;
          return function () {
            let _coll8 = new _dafny.Map();
            for (const _compr_8 of _dafny.IntegerRange(new BigNumber(989), new BigNumber(320))) {
              let _63_v29 = _compr_8;
              if (((new BigNumber(989)).isLessThanOrEqualTo(_63_v29)) && ((_63_v29).isLessThan(new BigNumber(320)))) {
                _coll8.push([_module.__default.safeModuloInt(_63_v29, (_pat_let_tv9).f25),_pat_let_tv8]);
              }
            }
            return _coll8;
          }();
        } else {
          let _64___mcc_h6 = (_source0).cf61;
          let _65_cf61 = _64___mcc_h6;
          return _dafny.Map.Empty.slice().updateUnsafe((_pat_let_tv10).f25,_pat_let_tv11);
        }
      }(_45_v26);
      r2 = (_22_v5).f25;
      return [r0, r1, r2];
    }
    static Main(__noArgsParameter) {
      let _66_v0;
      _66_v0 = _dafny.Seq.UnicodeFromString("hg");
      let _67_v1;
      _67_v1 = _module.D0.create_DC0(_66_v0);
      let _68_v2;
      _68_v2 = new _dafny.CodePoint('t'.codePointAt(0));
      let _69_v3;
      let _nw8 = Array((new BigNumber(12)).toNumber());
      _nw8[(_dafny.ZERO).toNumber()] = _66_v0;
      _nw8[(_dafny.ONE).toNumber()] = _dafny.Seq.UnicodeFromString("mexrhnq");
      _nw8[(new BigNumber(2)).toNumber()] = _66_v0;
      _nw8[(new BigNumber(3)).toNumber()] = _66_v0;
      _nw8[(new BigNumber(4)).toNumber()] = _66_v0;
      _nw8[(new BigNumber(5)).toNumber()] = (_67_v1).dtor_cf0;
      _nw8[(new BigNumber(6)).toNumber()] = _66_v0;
      _nw8[(new BigNumber(7)).toNumber()] = _dafny.Seq.UnicodeFromString("qu");
      _nw8[(new BigNumber(8)).toNumber()] = _dafny.Seq.UnicodeFromString("wggsbwf");
      _nw8[(new BigNumber(9)).toNumber()] = _dafny.Seq.UnicodeFromString("rnjar");
      _nw8[(new BigNumber(10)).toNumber()] = _66_v0;
      _nw8[(new BigNumber(11)).toNumber()] = _dafny.Seq.update(_66_v0, _module.__default.safeIndex(new BigNumber(-282), new BigNumber((_66_v0).length)), _68_v2);
      _69_v3 = _nw8;
      let _70_v4;
      _70_v4 = true;
      let _71_v5;
      _71_v5 = _dafny.MultiSet.fromElements(_70_v4, true, _70_v4);
      let _72_v6;
      _72_v6 = _dafny.MultiSet.fromElements(_71_v5);
      let _73_v7;
      let _nw9 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Set.Empty);
      _73_v7 = _nw9;
      let _74_v8;
      let _nw10 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
      _74_v8 = _nw10;
      let _75_v9;
      _75_v9 = new BigNumber(-20);
      let _76_v10;
      _76_v10 = _dafny.Map.Empty.slice().updateUnsafe(_74_v8,_75_v9);
      let _77_globalState;
      let _nw11 = new _module.GlobalState();
      _nw11.__ctor(true, new BigNumber(907), _69_v3, new BigNumber(-486), false, _72_v6, _66_v0, false, new BigNumber(-977), false, new BigNumber(632), new _dafny.CodePoint('f'.codePointAt(0)), new BigNumber(179), _73_v7, new BigNumber(435), new BigNumber(814), true, new BigNumber(558), (_76_v10).Merge(_76_v10), false, _74_v8, false, new BigNumber(634));
      _77_globalState = _nw11;
      let _78_i0;
      _78_i0 = _dafny.ZERO;
      L0: {
        while (_70_v4) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_78_i0)) {
              break L0;
            }
            _78_i0 = (_78_i0).plus(_dafny.ONE);
            let _79_v11;
            let _nw12 = Array((new BigNumber(16)).toNumber()).fill(false);
            _79_v11 = _nw12;
            let _80_v12;
            let _81_v13;
            let _82_v14;
            let _out0;
            let _out1;
            let _out2;
            let _outcollector0 = _module.__default.m0(_79_v11, _66_v0, _75_v9, _77_globalState);
            _out0 = _outcollector0[0];
            _out1 = _outcollector0[1];
            _out2 = _outcollector0[2];
            _80_v12 = _out0;
            _81_v13 = _out1;
            _82_v14 = _out2;
            let _83_v15;
            _83_v15 = _dafny.Map.Empty.slice().updateUnsafe(_75_v9,new BigNumber(921));
            let _index0 = _module.__default.safeIndex(new BigNumber(91), new BigNumber((_74_v8).length));
            (_74_v8)[_index0] = (((_83_v15).contains(new BigNumber((_66_v0).length))) ? ((_83_v15).get(new BigNumber((_66_v0).length))) : (new BigNumber(973)));
            let _index1 = _module.__default.safeIndex(new BigNumber(91), new BigNumber((_74_v8).length));
            (_74_v8)[_index1] = _module.__default.fm0(_75_v9, _70_v4, _77_globalState);
            let _84_v16;
            let _nw13 = Array((new BigNumber(22)).toNumber());
            _nw13[(_dafny.ZERO).toNumber()] = _69_v3;
            _nw13[(_dafny.ONE).toNumber()] = _69_v3;
            _nw13[(new BigNumber(2)).toNumber()] = _69_v3;
            _nw13[(new BigNumber(3)).toNumber()] = _69_v3;
            _nw13[(new BigNumber(4)).toNumber()] = _69_v3;
            _nw13[(new BigNumber(5)).toNumber()] = ((_70_v4) ? (_69_v3) : (_69_v3));
            _nw13[(new BigNumber(6)).toNumber()] = ((_70_v4) ? (_69_v3) : (_69_v3));
            _nw13[(new BigNumber(7)).toNumber()] = _69_v3;
            _nw13[(new BigNumber(8)).toNumber()] = _69_v3;
            _nw13[(new BigNumber(9)).toNumber()] = _69_v3;
            _nw13[(new BigNumber(10)).toNumber()] = _69_v3;
            _nw13[(new BigNumber(11)).toNumber()] = _69_v3;
            _nw13[(new BigNumber(12)).toNumber()] = _69_v3;
            _nw13[(new BigNumber(13)).toNumber()] = _69_v3;
            _nw13[(new BigNumber(14)).toNumber()] = _69_v3;
            _nw13[(new BigNumber(15)).toNumber()] = _69_v3;
            _nw13[(new BigNumber(16)).toNumber()] = _69_v3;
            _nw13[(new BigNumber(17)).toNumber()] = _69_v3;
            _nw13[(new BigNumber(18)).toNumber()] = _69_v3;
            _nw13[(new BigNumber(19)).toNumber()] = _69_v3;
            _nw13[(new BigNumber(20)).toNumber()] = _69_v3;
            _nw13[(new BigNumber(21)).toNumber()] = _69_v3;
            _84_v16 = _nw13;
            let _index2 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_84_v16).length));
            (_84_v16)[_index2] = _69_v3;
            let _index3 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_84_v16).length));
            (_84_v16)[_index3] = _69_v3;
            let _85_v17;
            _85_v17 = _dafny.Seq.of(_70_v4, (_82_v14).isLessThan(new BigNumber((_66_v0).length)), _70_v4);
            _85_v17 = _85_v17;
          }
        }
      }
      let _86_v18;
      let _nw14 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Map.Empty);
      _86_v18 = _nw14;
      let _87_v19;
      _87_v19 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(378),new BigNumber((_dafny.MultiSet.fromElements(_75_v9)).cardinality()));
      let _index4 = _module.__default.safeIndex(new BigNumber(957), new BigNumber((_86_v18).length));
      (_86_v18)[_index4] = _87_v19;
      let _88_v21;
      _88_v21 = _dafny.Seq.of(new BigNumber(467));
      let _index5 = _module.__default.safeIndex(new BigNumber(957), new BigNumber((_86_v18).length));
      (_86_v18)[_index5] = ((_87_v19).update(new BigNumber(688), _75_v9)).Merge((function () {
        let _coll9 = new _dafny.Map();
        for (const _compr_9 of (_88_v21).Elements) {
          let _89_v20 = _compr_9;
          if (_dafny.Seq.contains(_88_v21, _89_v20)) {
            _coll9.push([(_89_v20).minus(new BigNumber(732)),_75_v9]);
          }
        }
        return _coll9;
      }()).Merge(_87_v19));
      let _90_v22;
      let _nw15 = Array((new BigNumber(20)).toNumber()).fill(false);
      _90_v22 = _nw15;
      let _91_v23;
      _91_v23 = _dafny.Set.fromElements(_70_v4, _70_v4);
      let _index6 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length));
      (_90_v22)[_index6] = (_dafny.Set.fromElements(_70_v4, _70_v4, _70_v4, _70_v4)).IsProperSubsetOf(_91_v23);
      let _92_v24;
      _92_v24 = _dafny.Map.Empty.slice().updateUnsafe(_68_v2,_module.__default.fm1(_75_v9, _70_v4, _77_globalState));
      let _93_v25;
      _93_v25 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("jxuh"), _module.__default.fm2(_77_globalState), _66_v0);
      let _index7 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length));
      let _rhs4 = !(_70_v4) || (!(_dafny.Map.Empty.slice().updateUnsafe(_68_v2,!(!(_70_v4)))).equals(_92_v24));
      let _rhs5 = (_75_v9).multipliedBy((((_93_v25).contains(_66_v0)) ? ((_93_v25).get(_66_v0)) : (_75_v9)));
      let _rhs6 = _module.__default.safeModuloInt(_75_v9, _75_v9);
      let _rhs7 = !(true);
      let _lhs2 = _77_globalState;
      let _lhs3 = _77_globalState;
      let _lhs4 = _77_globalState;
      let _lhs5 = _90_v22;
      let _lhs6 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length));
      _lhs2.f7 = _rhs4;
      _lhs3.f14 = _rhs5;
      _lhs4.f14 = _rhs6;
      _lhs5[_lhs6] = _rhs7;
      let _94_v26;
      _94_v26 = _dafny.Seq.of(!(_70_v4));
      let _hi1 = new BigNumber((((true) ? (_dafny.Seq.update(_94_v26, _module.__default.safeIndex(_75_v9, new BigNumber((_94_v26).length)), (_90_v22)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length))])) : (_94_v26))).length);
      for (let _95_i1 = _75_v9; _95_i1.isLessThan(_hi1); _95_i1 = _95_i1.plus(_dafny.ONE)) {
        (_77_globalState).f14 = _75_v9;
        let _96_v27;
        let _nw16 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Seq.of());
        _96_v27 = _nw16;
        let _97_v28;
        _97_v28 = _dafny.Map.Empty.slice().updateUnsafe(_95_i1,_96_v27);
        _97_v28 = (_97_v28).update(_module.__default.safeModuloInt(_95_i1, _75_v9), _96_v27);
        _70_v4 = (_90_v22)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length))];
        let _index8 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_74_v8).length));
        (_74_v8)[_index8] = _75_v9;
        let _index9 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_74_v8).length));
        (_74_v8)[_index9] = _module.__default.safeDivisionInt(new BigNumber((_dafny.Set.fromElements(_70_v4)).length), (new BigNumber((_88_v21).length)).plus(_75_v9));
      }
      let _98_i2;
      _98_i2 = _dafny.ZERO;
      L1: {
        while (!((_75_v9).multipliedBy(_75_v9)).isEqualTo(_module.__default.safeDivisionInt(_75_v9, (_dafny.ZERO).minus(_75_v9)))) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_98_i2)) {
              break L1;
            }
            _98_i2 = (_98_i2).plus(_dafny.ONE);
            (_77_globalState).f8 = new BigNumber((_dafny.Seq.of((_91_v23).Union(_91_v23), _91_v23, _91_v23, _91_v23)).length);
            let _99_v29;
            let _100_v30;
            let _101_v31;
            let _out3;
            let _out4;
            let _out5;
            let _outcollector1 = _module.__default.m0(_90_v22, _66_v0, _75_v9, _77_globalState);
            _out3 = _outcollector1[0];
            _out4 = _outcollector1[1];
            _out5 = _outcollector1[2];
            _99_v29 = _out3;
            _100_v30 = _out4;
            _101_v31 = _out5;
            let _102_v32;
            _102_v32 = _dafny.Seq.of(_module.__default.fm3(_75_v9, _77_globalState), _92_v24, _dafny.Map.Empty.slice().updateUnsafe(_68_v2,_70_v4), _92_v24, _dafny.Map.Empty.slice().updateUnsafe(_68_v2,_70_v4));
            let _103_v33;
            let _104_v34;
            let _105_v35;
            let _out6;
            let _out7;
            let _out8;
            let _outcollector2 = _module.__default.m0(_90_v22, _dafny.Seq.Create(_module.__default.abs(new BigNumber(751)), ((_106_v2) => function (_107_i3) {
              return _106_v2;
            })(_68_v2)), (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_102_v32).length),_101_v31)).length)).multipliedBy(new BigNumber((_module.__default.fm4(_module.D0.create_DC0(_dafny.Seq.UnicodeFromString("k")), new BigNumber(998), _94_v26, _68_v2, _77_globalState)).cardinality())), _77_globalState);
            _out6 = _outcollector2[0];
            _out7 = _outcollector2[1];
            _out8 = _outcollector2[2];
            _103_v33 = _out6;
            _104_v34 = _out7;
            _105_v35 = _out8;
            let _index10 = _module.__default.safeIndex(new BigNumber(279), new BigNumber((_74_v8).length));
            (_74_v8)[_index10] = new BigNumber(-24);
            let _108_v36;
            _108_v36 = _dafny.Map.Empty.slice().updateUnsafe((_90_v22)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length))],_74_v8);
            let _index11 = _module.__default.safeIndex(new BigNumber(279), new BigNumber((_74_v8).length));
            (_74_v8)[_index11] = _module.__default.fm0(_101_v31, (_75_v9).isLessThan(new BigNumber((_108_v36).length)), _77_globalState);
          }
        }
      }
      (_77_globalState).f6 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ahdth"), _dafny.Seq.UnicodeFromString("syrkxpe"));
      (_77_globalState).f14 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(((_70_v4) ? (_75_v9) : (_75_v9)), (_dafny.ZERO).minus(new BigNumber((_module.__default.fm2(_77_globalState)).length))));
      if ((_90_v22)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length))]) {
        (_77_globalState).f7 = _70_v4;
        (_77_globalState).f8 = _module.__default.safeDivisionInt(new BigNumber((_66_v0).length), (_75_v9).plus(new BigNumber(30)));
        let _109_v37;
        _109_v37 = _dafny.Set.fromElements(_75_v9);
        let _110_v38;
        _110_v38 = _dafny.Seq.of(_module.__default.fm5(_109_v37, _77_globalState));
        (_77_globalState).f8 = _module.__default.safeModuloInt(_75_v9, _module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_66_v0).length)), new BigNumber((_110_v38).length)));
        let _111_v39;
        _111_v39 = _dafny.Map.Empty.slice().updateUnsafe(_70_v4,_75_v9);
        let _112_v40;
        _112_v40 = _module.D0.create_DC1(_75_v9, (_90_v22)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length))], _75_v9);
        _111_v39 = ((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.update(_66_v0, _module.__default.safeIndex(_75_v9, new BigNumber((_66_v0).length)), _68_v2)).length))).update((_112_v40).dtor_cf2, new BigNumber(-355))).Merge(_111_v39);
        _88_v21 = _88_v21;
      } else {
        let _113_v41;
        _113_v41 = _dafny.MultiSet.fromElements(_75_v9, new BigNumber((_88_v21).length));
        let _index12 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length));
        let _rhs8 = _75_v9;
        let _rhs9 = _dafny.Seq.update(_dafny.Seq.update(_66_v0, _module.__default.safeIndex(_75_v9, new BigNumber((_66_v0).length)), _68_v2), _module.__default.safeIndex(_75_v9, new BigNumber((_dafny.Seq.update(_66_v0, _module.__default.safeIndex(_75_v9, new BigNumber((_66_v0).length)), _68_v2)).length)), _68_v2);
        let _rhs10 = _70_v4;
        let _rhs11 = (_113_v41).IsSubsetOf(_113_v41);
        let _rhs12 = (_94_v26)[_module.__default.safeIndex((new BigNumber((_94_v26).length)).multipliedBy(new BigNumber((_dafny.Seq.update(_module.__default.fm2(_77_globalState), _module.__default.safeIndex(new BigNumber((_88_v21).length), new BigNumber((_module.__default.fm2(_77_globalState)).length)), _68_v2)).length)), new BigNumber((_94_v26).length))];
        let _lhs7 = _77_globalState;
        let _lhs8 = _77_globalState;
        let _lhs9 = _90_v22;
        let _lhs10 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length));
        let _lhs11 = _77_globalState;
        _lhs7.f14 = _rhs8;
        _lhs8.f6 = _rhs9;
        _lhs9[_lhs10] = _rhs10;
        _70_v4 = _rhs11;
        _lhs11.f21 = _rhs12;
        (_77_globalState).f11 = _module.__default.fm6(_75_v9, _77_globalState);
        (_77_globalState).f9 = (_module.D0.create_DC1(_75_v9, (_90_v22)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length))], _75_v9)).dtor_cf2;
        let _index13 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length));
        let _index14 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length));
        let _rhs13 = (((_90_v22)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length))]) ? ((_90_v22)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length))]) : (_70_v4));
        let _rhs14 = _70_v4;
        let _rhs15 = (_90_v22)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length))];
        let _lhs12 = _77_globalState;
        let _lhs13 = _90_v22;
        let _lhs14 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length));
        let _lhs15 = _90_v22;
        let _lhs16 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length));
        _lhs12.f7 = _rhs13;
        _lhs13[_lhs14] = _rhs14;
        _lhs15[_lhs16] = _rhs15;
        let _114_v42;
        let _115_v43;
        let _116_v44;
        let _out9;
        let _out10;
        let _out11;
        let _outcollector3 = _module.__default.m0(_90_v22, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-195)), ((_117_v2) => function (_118_i4) {
          return _117_v2;
        })(_68_v2)), (_dafny.ZERO).minus(_75_v9), _77_globalState);
        _out9 = _outcollector3[0];
        _out10 = _outcollector3[1];
        _out11 = _outcollector3[2];
        _114_v42 = _out9;
        _115_v43 = _out10;
        _116_v44 = _out11;
      }
      let _hi2 = _75_v9;
      for (let _119_i5 = _75_v9; _119_i5.isLessThan(_hi2); _119_i5 = _119_i5.plus(_dafny.ONE)) {
        let _index15 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length));
        (_90_v22)[_index15] = !_dafny.areEqual(_66_v0, _66_v0);
        (_77_globalState).f21 = _70_v4;
        (_77_globalState).f8 = (_119_i5).plus(_module.__default.fm0(new BigNumber(892), _70_v4, _77_globalState));
        let _120_v45;
        let _121_v46;
        let _122_v47;
        let _out12;
        let _out13;
        let _out14;
        let _outcollector4 = _module.__default.m0((((_90_v22)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length))]) ? (_90_v22) : (_90_v22)), _66_v0, _119_i5, _77_globalState);
        _out12 = _outcollector4[0];
        _out13 = _outcollector4[1];
        _out14 = _outcollector4[2];
        _120_v45 = _out12;
        _121_v46 = _out13;
        _122_v47 = _out14;
      }
      let _123_i6;
      _123_i6 = _dafny.ZERO;
      L2: {
        while ((new BigNumber(-886)).isLessThanOrEqualTo(_75_v9)) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_123_i6)) {
              break L2;
            }
            _123_i6 = (_123_i6).plus(_dafny.ONE);
            let _124_v48;
            _124_v48 = _dafny.Map.Empty.slice().updateUnsafe(_75_v9,_70_v4);
            let _rhs16 = _74_v8;
            let _rhs17 = (new BigNumber((_124_v48).length)).minus((_dafny.ZERO).minus(_75_v9));
            let _rhs18 = (_module.__default.fm0(_75_v9, _70_v4, _77_globalState)).isEqualTo(new BigNumber(743));
            let _rhs19 = ((_70_v4) ? (_70_v4) : (_70_v4));
            let _rhs20 = _75_v9;
            let _lhs17 = _77_globalState;
            let _lhs18 = _77_globalState;
            let _lhs19 = _77_globalState;
            let _lhs20 = _77_globalState;
            _74_v8 = _rhs16;
            _lhs17.f14 = _rhs17;
            _lhs18.f21 = _rhs18;
            _lhs19.f7 = _rhs19;
            _lhs20.f14 = _rhs20;
            (_77_globalState).f9 = (_90_v22)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length))];
            let _125_v49;
            _125_v49 = _dafny.Seq.of(_dafny.Seq.of((_90_v22)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length))]), _94_v26);
            (_77_globalState).f14 = _module.__default.safeModuloInt(_75_v9, (_75_v9).minus(new BigNumber(((_125_v49)[_module.__default.safeIndex(_75_v9, new BigNumber((_125_v49).length))]).length)));
            let _126_v50;
            let _nw17 = Array((new BigNumber(13)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
            _126_v50 = _nw17;
            let _index16 = _module.__default.safeIndex(new BigNumber(963), new BigNumber((_126_v50).length));
            (_126_v50)[_index16] = _68_v2;
            let _127_v51;
            _127_v51 = _dafny.MultiSet.fromElements(_75_v9);
            let _index17 = _module.__default.safeIndex(new BigNumber(963), new BigNumber((_126_v50).length));
            let _rhs21 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ghjmshpn"), _66_v0);
            let _rhs22 = (_66_v0)[_module.__default.safeIndex(new BigNumber((_127_v51).cardinality()), new BigNumber((_66_v0).length))];
            let _rhs23 = _68_v2;
            let _lhs21 = _77_globalState;
            let _lhs22 = _77_globalState;
            let _lhs23 = _126_v50;
            let _lhs24 = _module.__default.safeIndex(new BigNumber(963), new BigNumber((_126_v50).length));
            _lhs21.f6 = _rhs21;
            _lhs22.f11 = _rhs22;
            _lhs23[_lhs24] = _rhs23;
          }
        }
      }
      let _128_v52;
      let _129_v53;
      let _130_v54;
      let _out15;
      let _out16;
      let _out17;
      let _outcollector5 = _module.__default.m0(_90_v22, _module.__default.fm2(_77_globalState), _75_v9, _77_globalState);
      _out15 = _outcollector5[0];
      _out16 = _outcollector5[1];
      _out17 = _outcollector5[2];
      _128_v52 = _out15;
      _129_v53 = _out16;
      _130_v54 = _out17;
      let _131_v55;
      _131_v55 = _dafny.Set.fromElements(new BigNumber((_66_v0).length), _75_v9);
      let _hi3 = new BigNumber((_131_v55).length);
      for (let _132_i7 = new BigNumber(((_71_v5).Difference(_71_v5)).cardinality()); _132_i7.isLessThan(_hi3); _132_i7 = _132_i7.plus(_dafny.ONE)) {
        let _rhs24 = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_module.__default.fm2(_77_globalState)).length)));
        let _rhs25 = _module.__default.safeDivisionInt(new BigNumber(-737), _132_i7);
        let _rhs26 = _74_v8;
        let _lhs25 = _77_globalState;
        _lhs25.f8 = _rhs24;
        _130_v54 = _rhs25;
        _74_v8 = _rhs26;
        let _133_v56;
        let _nw18 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Map.Empty);
        _133_v56 = _nw18;
        _133_v56 = _133_v56;
        let _134_v57;
        let _nw19 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Seq.of());
        _134_v57 = _nw19;
        let _135_v58;
        _135_v58 = _dafny.Seq.of(_74_v8);
        let _136_v59;
        _136_v59 = _module.D1.create_DC3(_135_v58);
        let _index18 = _module.__default.safeIndex(new BigNumber(996), new BigNumber((_134_v57).length));
        (_134_v57)[_index18] = (_136_v59).dtor_cf5;
        let _index19 = _module.__default.safeIndex(new BigNumber(996), new BigNumber((_134_v57).length));
        (_134_v57)[_index19] = _135_v58;
        let _137_v60;
        _137_v60 = _module.D1.create_DC5(_130_v54, true, new BigNumber((_131_v55).length), _module.__default.fm0(_75_v9, _70_v4, _77_globalState));
        let _index20 = _module.__default.safeIndex(new BigNumber(957), new BigNumber((_86_v18).length));
        (_86_v18)[_index20] = ((_86_v18)[_module.__default.safeIndex(new BigNumber(957), new BigNumber((_86_v18).length))]).update((_module.__default.fm0(_module.__default.fm0(_130_v54, _70_v4, _77_globalState), (_90_v22)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length))], _77_globalState)).plus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_75_v9, _132_i7, (_137_v60).dtor_cf12, _132_i7)).length)))), _75_v9);
      }
      let _138_v61;
      _138_v61 = _dafny.MultiSet.fromElements(_75_v9, new BigNumber((_dafny.Seq.update(_dafny.Seq.of(new BigNumber(75)), _module.__default.safeIndex(_75_v9, new BigNumber((_dafny.Seq.of(new BigNumber(75))).length)), _130_v54)).length));
      if ((_70_v4) === (((_138_v61).update(new BigNumber((_dafny.Set.fromElements(_67_v1, _module.D0.create_DC0(_66_v0))).length), _module.__default.abs(new BigNumber((_66_v0).length)))).IsDisjointFrom(_138_v61))) {
        (_77_globalState).f9 = _70_v4;
        let _index21 = _module.__default.safeIndex(new BigNumber(268), new BigNumber((_74_v8).length));
        (_74_v8)[_index21] = (_dafny.ZERO).minus(_75_v9);
        let _index22 = _module.__default.safeIndex(new BigNumber(268), new BigNumber((_74_v8).length));
        (_74_v8)[_index22] = new BigNumber(-38);
        let _139_v62;
        let _nw20 = new _module.C0();
        _nw20.__ctor();
        _139_v62 = _nw20;
        let _140_v63;
        _140_v63 = _dafny.Seq.of(_71_v5, _dafny.MultiSet.fromElements((_90_v22)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length))]));
        let _141_v64;
        let _nw21 = new _module.C3();
        _nw21.__ctor(_140_v63, (_74_v8)[_module.__default.safeIndex(new BigNumber(268), new BigNumber((_74_v8).length))], !(_70_v4), _dafny.Seq.of(_75_v9));
        _141_v64 = _nw21;
        let _142_v65;
        _142_v65 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_139_v62,_141_v64));
        let _143_v66;
        _143_v66 = _dafny.Map.Empty.slice().updateUnsafe(_139_v62,_141_v64);
        let _144_v67;
        let _nw22 = new _module.C3();
        _nw22.__ctor(_module.__default.fm29(_75_v9, _dafny.Set.fromElements(_130_v54, _75_v9), _77_globalState), (_74_v8)[_module.__default.safeIndex(new BigNumber(268), new BigNumber((_74_v8).length))], !_dafny.Seq.contains(_88_v21, (_74_v8)[_module.__default.safeIndex(new BigNumber(268), new BigNumber((_74_v8).length))]), _dafny.Seq.update(_88_v21, _module.__default.safeIndex((((_142_v65).contains(_143_v66)) ? ((_142_v65).get(_143_v66)) : (_130_v54)), new BigNumber((_88_v21).length)), _75_v9));
        _144_v67 = _nw22;
        let _index23 = _module.__default.safeIndex(new BigNumber(268), new BigNumber((_74_v8).length));
        let _rhs27 = _90_v22;
        let _rhs28 = _130_v54;
        let _rhs29 = _module.__default.fm1(_75_v9, ((false) ? (!(true)) : ((_90_v22)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length))])), _77_globalState);
        let _lhs26 = _74_v8;
        let _lhs27 = _module.__default.safeIndex(new BigNumber(268), new BigNumber((_74_v8).length));
        let _lhs28 = _77_globalState;
        _90_v22 = _rhs27;
        _lhs26[_lhs27] = _rhs28;
        _lhs28.f7 = _rhs29;
        let _index24 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length));
        (_90_v22)[_index24] = _dafny.Seq.IsProperPrefixOf(_66_v0, _66_v0);
      } else {
        let _145_v68;
        _145_v68 = _dafny.Map.Empty.slice().updateUnsafe(_70_v4,!((_90_v22)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length))]));
        let _146_v69;
        _146_v69 = _dafny.Map.Empty.slice().updateUnsafe((_90_v22)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length))],new BigNumber((_145_v68).length));
        _75_v9 = (((_71_v5).contains((_90_v22)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length))])) ? ((_71_v5).get((_90_v22)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length))])) : ((((_146_v69).contains((_90_v22)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length))])) ? ((_146_v69).get((_90_v22)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length))])) : (_75_v9))));
        let _147_v70;
        _147_v70 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(347)), ((_148_v2) => function (_149_i8) {
          return _148_v2;
        })(_68_v2)),(_90_v22)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length))]);
        _147_v70 = (_147_v70).update(_66_v0, _70_v4);
        let _150_v71;
        _150_v71 = _dafny.Seq.of(_94_v26, _94_v26);
        _94_v26 = _dafny.Seq.Concat(_94_v26, _dafny.Seq.Concat((_150_v71)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_66_v0).length)), new BigNumber((_150_v71).length))], _94_v26));
        let _151_v72;
        _151_v72 = _dafny.Seq.of(_145_v68);
        _151_v72 = _151_v72;
        let _index25 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_74_v8).length));
        (_74_v8)[_index25] = _130_v54;
        let _index26 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_74_v8).length));
        let _rhs30 = _70_v4;
        let _rhs31 = _75_v9;
        let _rhs32 = _75_v9;
        let _rhs33 = _75_v9;
        let _lhs29 = _77_globalState;
        let _lhs30 = _74_v8;
        let _lhs31 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_74_v8).length));
        _lhs29.f7 = _rhs30;
        _lhs30[_lhs31] = _rhs31;
        _75_v9 = _rhs32;
        _75_v9 = _rhs33;
      }
      let _152_v73;
      _152_v73 = _dafny.Map.Empty.slice().updateUnsafe(!(_70_v4),(_90_v22)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length))]);
      _152_v73 = (_152_v73).update(_70_v4, _70_v4);
      let _153_v74;
      let _nw23 = Array((new BigNumber(4)).toNumber());
      _153_v74 = _nw23;
      let _154_v75;
      _154_v75 = _dafny.Map.Empty.slice().updateUnsafe((_90_v22)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length))],_153_v74);
      let _155_v76;
      let _nw24 = new _module.C2();
      _nw24.__ctor((_86_v18)[_module.__default.safeIndex(new BigNumber(957), new BigNumber((_86_v18).length))], (_90_v22)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length))], _70_v4, _module.__default.fm19(_75_v9, (_90_v22)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_90_v22).length))], _77_globalState), new BigNumber((_154_v75).length));
      _155_v76 = _nw24;
      let _156_v77;
      let _nw25 = Array((new BigNumber(11)).toNumber()).fill(_dafny.MultiSet.Empty);
      _156_v77 = _nw25;
      let _157_v78;
      _157_v78 = _dafny.Map.Empty.slice().updateUnsafe(_155_v76,_156_v77);
      _157_v78 = (_157_v78).update(_155_v76, _156_v77);
      (_77_globalState).f7 = (_module.__default.safeModuloInt(_130_v54, new BigNumber((_dafny.Seq.update(_66_v0, _module.__default.safeIndex(_75_v9, new BigNumber((_66_v0).length)), _68_v2)).length))).isLessThanOrEqualTo(_75_v9);
      process.stdout.write((_66_v0).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_67_v1).dtor_cf0).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_68_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_69_v3)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_69_v3)[_dafny.ONE]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_69_v3)[new BigNumber(2)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_69_v3)[new BigNumber(3)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_69_v3)[new BigNumber(4)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_69_v3)[new BigNumber(5)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_69_v3)[new BigNumber(6)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_69_v3)[new BigNumber(7)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_69_v3)[new BigNumber(8)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_69_v3)[new BigNumber(9)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_69_v3)[new BigNumber(10)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_69_v3)[new BigNumber(11)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_70_v4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_71_v5).equals(_dafny.MultiSet.fromElements(true, true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_72_v6).equals(_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(true, true, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_74_v8)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_74_v8)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_74_v8)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_74_v8)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_75_v9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_76_v10).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_77_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_77_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_77_globalState).f2)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_77_globalState).f2)[_dafny.ONE]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_77_globalState).f2)[new BigNumber(2)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_77_globalState).f2)[new BigNumber(3)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_77_globalState).f2)[new BigNumber(4)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_77_globalState).f2)[new BigNumber(5)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_77_globalState).f2)[new BigNumber(6)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_77_globalState).f2)[new BigNumber(7)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_77_globalState).f2)[new BigNumber(8)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_77_globalState).f2)[new BigNumber(9)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_77_globalState).f2)[new BigNumber(10)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((((_77_globalState).f2)[new BigNumber(11)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_77_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_77_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_77_globalState).f5).equals(_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(true, true, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_77_globalState.f6).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_77_globalState.f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_77_globalState.f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_77_globalState.f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_77_globalState).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_77_globalState.f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_77_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_77_globalState.f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_77_globalState).f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_77_globalState).f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_77_globalState).f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_77_globalState).f18).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_77_globalState).f19));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_77_globalState).f20)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_77_globalState).f20)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_77_globalState).f20)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_77_globalState).f20)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_77_globalState.f21));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_77_globalState).f22));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_78_i0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_86_v18)[new BigNumber(17)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(378),_dafny.ONE).updateUnsafe(new BigNumber(688),new BigNumber(-20)).updateUnsafe(new BigNumber(-265),new BigNumber(-20)).updateUnsafe(new BigNumber(4),new BigNumber(-20)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_87_v19).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(378),_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_88_v21, _dafny.Seq.of(new BigNumber(467)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_90_v22)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_91_v23).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_92_v24).equals(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('t'.codePointAt(0)),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_93_v25).equals(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("jxuh"), _dafny.Seq.UnicodeFromString("koddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii"), _dafny.Seq.UnicodeFromString("hg")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_94_v26, _dafny.Seq.of(false, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_98_i2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_123_i6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_128_v52).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_128_v52).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_128_v52).dtor_cf3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_129_v53).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,false).updateUnsafe(new BigNumber(-1),false).updateUnsafe(_dafny.ONE,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_130_v54));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_v55).equals(_dafny.Set.fromElements(new BigNumber(2), new BigNumber(-20)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_138_v61).equals(_dafny.MultiSet.fromElements(new BigNumber(-20), _dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_152_v73).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false).updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_154_v75).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_155_v76).f28));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_155_v76.f29, _dafny.Seq.of(new BigNumber(3), new BigNumber(-2), _dafny.ONE, _dafny.ONE, new BigNumber(7)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_155_v76).f23));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_157_v78).length)));
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
    static create_DC0(cf0) {
      let $dt = new D0(1);
      $dt.cf0 = cf0;
      return $dt;
    }
    static create_DC2(cf4) {
      let $dt = new D0(2);
      $dt.cf4 = cf4;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC0() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf4() { return this.cf4; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC0" + "(" + this.cf0.toVerbatimString(true) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf4) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1) && this.cf2 === other.cf2 && _dafny.areEqual(this.cf3, other.cf3);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf0, other.cf0);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf4, other.cf4);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(_dafny.ZERO, false, _dafny.ZERO);
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
    static create_DC4(cf6, cf7, cf8) {
      let $dt = new D1(0);
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC5(cf9, cf10, cf11, cf12) {
      let $dt = new D1(1);
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC6(cf13) {
      let $dt = new D1(2);
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC3(cf5) {
      let $dt = new D1(3);
      $dt.cf5 = cf5;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC6() { return this.$tag === 2; }
    get is_DC3() { return this.$tag === 3; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf5() { return this.cf5; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC6" + "(" + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 3) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf5) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf6, other.cf6) && this.cf7 === other.cf7 && this.cf8 === other.cf8;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf9, other.cf9) && this.cf10 === other.cf10 && _dafny.areEqual(this.cf11, other.cf11) && _dafny.areEqual(this.cf12, other.cf12);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf13 === other.cf13;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf5, other.cf5);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(_dafny.Map.Empty, false, false);
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
    static create_DC8(cf15, cf16, cf17, cf18, cf19) {
      let $dt = new D2(0);
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      return $dt;
    }
    static create_DC7(cf14) {
      let $dt = new D2(1);
      $dt.cf14 = cf14;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf14() { return this.cf14; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf14) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf15, other.cf15) && _dafny.areEqual(this.cf16, other.cf16) && _dafny.areEqual(this.cf17, other.cf17) && _dafny.areEqual(this.cf18, other.cf18) && this.cf19 === other.cf19;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf14, other.cf14);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC8(_dafny.ZERO, _dafny.ZERO, new _dafny.CodePoint('D'.codePointAt(0)), _dafny.ZERO, false);
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
    static create_DC10(cf21) {
      let $dt = new D3(0);
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC11(cf22) {
      let $dt = new D3(1);
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC12(cf23) {
      let $dt = new D3(2);
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC9(cf20) {
      let $dt = new D3(3);
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC13(cf24) {
      let $dt = new D3(4);
      $dt.cf24 = cf24;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get is_DC12() { return this.$tag === 2; }
    get is_DC9() { return this.$tag === 3; }
    get is_DC13() { return this.$tag === 4; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf24() { return this.cf24; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC12" + "(" + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 3) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 4) {
        return "D3.DC13" + "(" + _dafny.toString(this.cf24) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf21 === other.cf21;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf22, other.cf22);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf23, other.cf23);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf20, other.cf20);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf24, other.cf24);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC10(false);
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
    static create_DC15(cf26, cf27, cf28) {
      let $dt = new D4(0);
      $dt.cf26 = cf26;
      $dt.cf27 = cf27;
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC14(cf25) {
      let $dt = new D4(1);
      $dt.cf25 = cf25;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf25() { return this.cf25; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC15" + "(" + _dafny.toString(this.cf26) + ", " + _dafny.toString(this.cf27) + ", " + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC14" + "(" + _dafny.toString(this.cf25) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf26, other.cf26) && _dafny.areEqual(this.cf27, other.cf27) && _dafny.areEqual(this.cf28, other.cf28);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf25 === other.cf25;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC15(_dafny.ZERO, new _dafny.CodePoint('D'.codePointAt(0)), _dafny.ZERO);
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
    static create_DC16(cf29) {
      let $dt = new D5(0);
      $dt.cf29 = cf29;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get dtor_cf29() { return this.cf29; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC16" + "(" + _dafny.toString(this.cf29) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf29 === other.cf29;
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
    static create_DC17(cf30) {
      let $dt = new D6(0);
      $dt.cf30 = cf30;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get dtor_cf30() { return this.cf30; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf30) + ")";
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
      return _dafny.Seq.of();
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
    static create_DC18(cf31) {
      let $dt = new D7(0);
      $dt.cf31 = cf31;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get dtor_cf31() { return this.cf31; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC18" + "(" + _dafny.toString(this.cf31) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf31, other.cf31);
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
    static create_DC20(cf33, cf34) {
      let $dt = new D8(0);
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      return $dt;
    }
    static create_DC21(cf35, cf36, cf37) {
      let $dt = new D8(1);
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC22(cf38) {
      let $dt = new D8(2);
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC19(cf32) {
      let $dt = new D8(3);
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC23(cf39) {
      let $dt = new D8(4);
      $dt.cf39 = cf39;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get is_DC21() { return this.$tag === 1; }
    get is_DC22() { return this.$tag === 2; }
    get is_DC19() { return this.$tag === 3; }
    get is_DC23() { return this.$tag === 4; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf39() { return this.cf39; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC20" + "(" + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC21" + "(" + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC22" + "(" + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 3) {
        return "D8.DC19" + "(" + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 4) {
        return "D8.DC23" + "(" + _dafny.toString(this.cf39) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf33 === other.cf33 && _dafny.areEqual(this.cf34, other.cf34);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf35 === other.cf35 && this.cf36 === other.cf36 && _dafny.areEqual(this.cf37, other.cf37);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf38 === other.cf38;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf32, other.cf32);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf39, other.cf39);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC20(false, _dafny.ZERO);
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
    static create_DC25(cf41, cf42) {
      let $dt = new D9(0);
      $dt.cf41 = cf41;
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC26(cf43) {
      let $dt = new D9(1);
      $dt.cf43 = cf43;
      return $dt;
    }
    static create_DC27(cf44, cf45, cf46, cf47, cf48) {
      let $dt = new D9(2);
      $dt.cf44 = cf44;
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      return $dt;
    }
    static create_DC24(cf40) {
      let $dt = new D9(3);
      $dt.cf40 = cf40;
      return $dt;
    }
    static create_DC28(cf49) {
      let $dt = new D9(4);
      $dt.cf49 = cf49;
      return $dt;
    }
    get is_DC25() { return this.$tag === 0; }
    get is_DC26() { return this.$tag === 1; }
    get is_DC27() { return this.$tag === 2; }
    get is_DC24() { return this.$tag === 3; }
    get is_DC28() { return this.$tag === 4; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf49() { return this.cf49; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC25" + "(" + _dafny.toString(this.cf41) + ", " + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC26" + "(" + _dafny.toString(this.cf43) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC27" + "(" + _dafny.toString(this.cf44) + ", " + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ")";
      } else if (this.$tag === 3) {
        return "D9.DC24" + "(" + _dafny.toString(this.cf40) + ")";
      } else if (this.$tag === 4) {
        return "D9.DC28" + "(" + _dafny.toString(this.cf49) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf41 === other.cf41 && _dafny.areEqual(this.cf42, other.cf42);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf43, other.cf43);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf44 === other.cf44 && _dafny.areEqual(this.cf45, other.cf45) && this.cf46 === other.cf46 && this.cf47 === other.cf47 && this.cf48 === other.cf48;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf40 === other.cf40;
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf49, other.cf49);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC25(false, _dafny.Set.Empty);
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
    static create_DC29(cf50) {
      let $dt = new D10(0);
      $dt.cf50 = cf50;
      return $dt;
    }
    get is_DC29() { return this.$tag === 0; }
    get dtor_cf50() { return this.cf50; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC29" + "(" + _dafny.toString(this.cf50) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf50, other.cf50);
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
    static create_DC31(cf52) {
      let $dt = new D11(0);
      $dt.cf52 = cf52;
      return $dt;
    }
    static create_DC32(cf53) {
      let $dt = new D11(1);
      $dt.cf53 = cf53;
      return $dt;
    }
    static create_DC30(cf51) {
      let $dt = new D11(2);
      $dt.cf51 = cf51;
      return $dt;
    }
    static create_DC33(cf54) {
      let $dt = new D11(3);
      $dt.cf54 = cf54;
      return $dt;
    }
    get is_DC31() { return this.$tag === 0; }
    get is_DC32() { return this.$tag === 1; }
    get is_DC30() { return this.$tag === 2; }
    get is_DC33() { return this.$tag === 3; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf54() { return this.cf54; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC31" + "(" + _dafny.toString(this.cf52) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC32" + "(" + this.cf53.toVerbatimString(true) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC30" + "(" + _dafny.toString(this.cf51) + ")";
      } else if (this.$tag === 3) {
        return "D11.DC33" + "(" + _dafny.toString(this.cf54) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf52, other.cf52);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf53, other.cf53);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf51 === other.cf51;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf54, other.cf54);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC31(_dafny.ZERO);
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
    static create_DC35() {
      let $dt = new D12(0);
      return $dt;
    }
    static create_DC36(cf56, cf57, cf58, cf59, cf60) {
      let $dt = new D12(1);
      $dt.cf56 = cf56;
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      $dt.cf60 = cf60;
      return $dt;
    }
    static create_DC34(cf55) {
      let $dt = new D12(2);
      $dt.cf55 = cf55;
      return $dt;
    }
    static create_DC37(cf61) {
      let $dt = new D12(3);
      $dt.cf61 = cf61;
      return $dt;
    }
    get is_DC35() { return this.$tag === 0; }
    get is_DC36() { return this.$tag === 1; }
    get is_DC34() { return this.$tag === 2; }
    get is_DC37() { return this.$tag === 3; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf61() { return this.cf61; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC35";
      } else if (this.$tag === 1) {
        return "D12.DC36" + "(" + _dafny.toString(this.cf56) + ", " + _dafny.toString(this.cf57) + ", " + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ", " + _dafny.toString(this.cf60) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC34" + "(" + _dafny.toString(this.cf55) + ")";
      } else if (this.$tag === 3) {
        return "D12.DC37" + "(" + _dafny.toString(this.cf61) + ")";
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
        return other.$tag === 1 && this.cf56 === other.cf56 && this.cf57 === other.cf57 && _dafny.areEqual(this.cf58, other.cf58) && _dafny.areEqual(this.cf59, other.cf59) && _dafny.areEqual(this.cf60, other.cf60);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf55, other.cf55);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf61, other.cf61);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC35();
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
    static create_DC39(cf63) {
      let $dt = new D13(0);
      $dt.cf63 = cf63;
      return $dt;
    }
    static create_DC38(cf62) {
      let $dt = new D13(1);
      $dt.cf62 = cf62;
      return $dt;
    }
    static create_DC40(cf64) {
      let $dt = new D13(2);
      $dt.cf64 = cf64;
      return $dt;
    }
    get is_DC39() { return this.$tag === 0; }
    get is_DC38() { return this.$tag === 1; }
    get is_DC40() { return this.$tag === 2; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf64() { return this.cf64; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC39" + "(" + _dafny.toString(this.cf63) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC38" + "(" + _dafny.toString(this.cf62) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC40" + "(" + _dafny.toString(this.cf64) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf63, other.cf63);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf62, other.cf62);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf64, other.cf64);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC39(_dafny.ZERO);
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
    static create_DC42(cf66) {
      let $dt = new D14(0);
      $dt.cf66 = cf66;
      return $dt;
    }
    static create_DC41(cf65) {
      let $dt = new D14(1);
      $dt.cf65 = cf65;
      return $dt;
    }
    get is_DC42() { return this.$tag === 0; }
    get is_DC41() { return this.$tag === 1; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf65() { return this.cf65; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC42" + "(" + _dafny.toString(this.cf66) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC41" + "(" + _dafny.toString(this.cf65) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf66, other.cf66);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf65, other.cf65);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC42(_dafny.Map.Empty);
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
    static create_DC44(cf68, cf69, cf70, cf71, cf72) {
      let $dt = new D15(0);
      $dt.cf68 = cf68;
      $dt.cf69 = cf69;
      $dt.cf70 = cf70;
      $dt.cf71 = cf71;
      $dt.cf72 = cf72;
      return $dt;
    }
    static create_DC43(cf67) {
      let $dt = new D15(1);
      $dt.cf67 = cf67;
      return $dt;
    }
    static create_DC45(cf73) {
      let $dt = new D15(2);
      $dt.cf73 = cf73;
      return $dt;
    }
    get is_DC44() { return this.$tag === 0; }
    get is_DC43() { return this.$tag === 1; }
    get is_DC45() { return this.$tag === 2; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf73() { return this.cf73; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC44" + "(" + _dafny.toString(this.cf68) + ", " + _dafny.toString(this.cf69) + ", " + _dafny.toString(this.cf70) + ", " + _dafny.toString(this.cf71) + ", " + _dafny.toString(this.cf72) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC43" + "(" + _dafny.toString(this.cf67) + ")";
      } else if (this.$tag === 2) {
        return "D15.DC45" + "(" + _dafny.toString(this.cf73) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf68 === other.cf68 && _dafny.areEqual(this.cf69, other.cf69) && _dafny.areEqual(this.cf70, other.cf70) && this.cf71 === other.cf71 && this.cf72 === other.cf72;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf67, other.cf67);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf73, other.cf73);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC44(false, _dafny.Seq.of(), _dafny.ZERO, false, false);
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
    static create_DC47(cf75, cf76, cf77, cf78) {
      let $dt = new D16(0);
      $dt.cf75 = cf75;
      $dt.cf76 = cf76;
      $dt.cf77 = cf77;
      $dt.cf78 = cf78;
      return $dt;
    }
    static create_DC46(cf74) {
      let $dt = new D16(1);
      $dt.cf74 = cf74;
      return $dt;
    }
    get is_DC47() { return this.$tag === 0; }
    get is_DC46() { return this.$tag === 1; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf74() { return this.cf74; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC47" + "(" + _dafny.toString(this.cf75) + ", " + _dafny.toString(this.cf76) + ", " + _dafny.toString(this.cf77) + ", " + _dafny.toString(this.cf78) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC46" + "(" + _dafny.toString(this.cf74) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf75, other.cf75) && _dafny.areEqual(this.cf76, other.cf76) && this.cf77 === other.cf77 && _dafny.areEqual(this.cf78, other.cf78);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf74, other.cf74);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC47(_dafny.ZERO, _dafny.ZERO, null, _dafny.ZERO);
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
    static create_DC49(cf80, cf81, cf82, cf83) {
      let $dt = new D17(0);
      $dt.cf80 = cf80;
      $dt.cf81 = cf81;
      $dt.cf82 = cf82;
      $dt.cf83 = cf83;
      return $dt;
    }
    static create_DC48(cf79) {
      let $dt = new D17(1);
      $dt.cf79 = cf79;
      return $dt;
    }
    get is_DC49() { return this.$tag === 0; }
    get is_DC48() { return this.$tag === 1; }
    get dtor_cf80() { return this.cf80; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf82() { return this.cf82; }
    get dtor_cf83() { return this.cf83; }
    get dtor_cf79() { return this.cf79; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC49" + "(" + _dafny.toString(this.cf80) + ", " + _dafny.toString(this.cf81) + ", " + _dafny.toString(this.cf82) + ", " + _dafny.toString(this.cf83) + ")";
      } else if (this.$tag === 1) {
        return "D17.DC48" + "(" + _dafny.toString(this.cf79) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf80, other.cf80) && this.cf81 === other.cf81 && this.cf82 === other.cf82 && this.cf83 === other.cf83;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf79, other.cf79);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D17.create_DC49(_dafny.ZERO, false, false, false);
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
    static create_DC51() {
      let $dt = new D18(0);
      return $dt;
    }
    static create_DC50(cf84) {
      let $dt = new D18(1);
      $dt.cf84 = cf84;
      return $dt;
    }
    get is_DC51() { return this.$tag === 0; }
    get is_DC50() { return this.$tag === 1; }
    get dtor_cf84() { return this.cf84; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC51";
      } else if (this.$tag === 1) {
        return "D18.DC50" + "(" + _dafny.toString(this.cf84) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf84, other.cf84);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC51();
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
    static create_DC53(cf86, cf87) {
      let $dt = new D19(0);
      $dt.cf86 = cf86;
      $dt.cf87 = cf87;
      return $dt;
    }
    static create_DC52(cf85) {
      let $dt = new D19(1);
      $dt.cf85 = cf85;
      return $dt;
    }
    static create_DC54(cf88) {
      let $dt = new D19(2);
      $dt.cf88 = cf88;
      return $dt;
    }
    get is_DC53() { return this.$tag === 0; }
    get is_DC52() { return this.$tag === 1; }
    get is_DC54() { return this.$tag === 2; }
    get dtor_cf86() { return this.cf86; }
    get dtor_cf87() { return this.cf87; }
    get dtor_cf85() { return this.cf85; }
    get dtor_cf88() { return this.cf88; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC53" + "(" + _dafny.toString(this.cf86) + ", " + _dafny.toString(this.cf87) + ")";
      } else if (this.$tag === 1) {
        return "D19.DC52" + "(" + _dafny.toString(this.cf85) + ")";
      } else if (this.$tag === 2) {
        return "D19.DC54" + "(" + _dafny.toString(this.cf88) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf86, other.cf86) && _dafny.areEqual(this.cf87, other.cf87);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf85, other.cf85);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf88, other.cf88);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D19.create_DC53(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC56(cf90, cf91, cf92, cf93, cf94) {
      let $dt = new D20(0);
      $dt.cf90 = cf90;
      $dt.cf91 = cf91;
      $dt.cf92 = cf92;
      $dt.cf93 = cf93;
      $dt.cf94 = cf94;
      return $dt;
    }
    static create_DC57() {
      let $dt = new D20(1);
      return $dt;
    }
    static create_DC55(cf89) {
      let $dt = new D20(2);
      $dt.cf89 = cf89;
      return $dt;
    }
    static create_DC58(cf95) {
      let $dt = new D20(3);
      $dt.cf95 = cf95;
      return $dt;
    }
    get is_DC56() { return this.$tag === 0; }
    get is_DC57() { return this.$tag === 1; }
    get is_DC55() { return this.$tag === 2; }
    get is_DC58() { return this.$tag === 3; }
    get dtor_cf90() { return this.cf90; }
    get dtor_cf91() { return this.cf91; }
    get dtor_cf92() { return this.cf92; }
    get dtor_cf93() { return this.cf93; }
    get dtor_cf94() { return this.cf94; }
    get dtor_cf89() { return this.cf89; }
    get dtor_cf95() { return this.cf95; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC56" + "(" + _dafny.toString(this.cf90) + ", " + _dafny.toString(this.cf91) + ", " + _dafny.toString(this.cf92) + ", " + _dafny.toString(this.cf93) + ", " + _dafny.toString(this.cf94) + ")";
      } else if (this.$tag === 1) {
        return "D20.DC57";
      } else if (this.$tag === 2) {
        return "D20.DC55" + "(" + _dafny.toString(this.cf89) + ")";
      } else if (this.$tag === 3) {
        return "D20.DC58" + "(" + _dafny.toString(this.cf95) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf90 === other.cf90 && _dafny.areEqual(this.cf91, other.cf91) && this.cf92 === other.cf92 && _dafny.areEqual(this.cf93, other.cf93) && _dafny.areEqual(this.cf94, other.cf94);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf89, other.cf89);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf95, other.cf95);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC56(false, _dafny.ZERO, [], new _dafny.CodePoint('D'.codePointAt(0)), _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D20.Default();
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
      this.f6 = _dafny.Seq.UnicodeFromString("");
      this.f7 = false;
      this.f8 = _dafny.ZERO;
      this.f9 = false;
      this.f11 = new _dafny.CodePoint('D'.codePointAt(0));
      this.f14 = _dafny.ZERO;
      this.f21 = false;
      this._f0 = false;
      this._f1 = _dafny.ZERO;
      this._f2 = [];
      this._f3 = _dafny.ZERO;
      this._f4 = false;
      this._f5 = _dafny.MultiSet.Empty;
      this._f10 = _dafny.ZERO;
      this._f12 = _dafny.ZERO;
      this._f13 = [];
      this._f15 = _dafny.ZERO;
      this._f16 = false;
      this._f17 = _dafny.ZERO;
      this._f18 = _dafny.Map.Empty;
      this._f19 = false;
      this._f20 = [];
      this._f22 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this).f6 = f6;
      (_this).f7 = f7;
      (_this).f8 = f8;
      (_this).f9 = f9;
      (_this)._f10 = f10;
      (_this).f11 = f11;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      (_this).f14 = f14;
      (_this)._f15 = f15;
      (_this)._f16 = f16;
      (_this)._f17 = f17;
      (_this)._f18 = f18;
      (_this)._f19 = f19;
      (_this)._f20 = f20;
      (_this).f21 = f21;
      (_this)._f22 = f22;
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
    get f10() {
      let _this = this;
      return _this._f10;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f13() {
      let _this = this;
      return _this._f13;
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
    get f19() {
      let _this = this;
      return _this._f19;
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
    get f22() {
      let _this = this;
      return _this._f22;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm20(p0, p1, p2, p3, globalState) {
      let _this = this;
      return false;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f29 = _dafny.Seq.of();
      this._f23 = _dafny.ZERO;
      this._f28 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f29() {
      let _this = this;
      return _this._f29;
    };
    set f29(value) {
      let _this = this;
      _this._f29 = value;
    };
    get f23() {
      let _this = this;
      return _this._f23;
    };
    get f28() {
      let _this = this;
      return _this._f28;
    };
    __ctor(f28, f29, f23) {
      let _this = this;
      (_this)._f28 = f28;
      (_this)._f29 = f29;
      (_this)._f23 = f23;
      return;
    }
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Set.Empty;
      let _index27 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((p0).length));
      (p0)[_index27] = p1;
      let _index28 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((p0).length));
      (p0)[_index28] = (_this).f28;
      let _158_v0;
      _158_v0 = _dafny.Seq.UnicodeFromString("pilipchus");
      let _159_v1;
      _159_v1 = _dafny.Set.fromElements(_158_v0);
      (_this).f29 = _module.__default.fm19(new BigNumber((_158_v0).length), !(_159_v1).contains(_158_v0), globalState);
      let _160_v2;
      _160_v2 = new _dafny.CodePoint('n'.codePointAt(0));
      let _161_v3;
      _161_v3 = _dafny.MultiSet.fromElements(_160_v2);
      let _index29 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((p0).length));
      let _rhs34 = _module.__default.fm0((_this).f23, (p0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((p0).length))], globalState);
      let _rhs35 = (_this).f28;
      let _rhs36 = (_this).f23;
      let _rhs37 = _161_v3;
      let _rhs38 = p1;
      let _lhs32 = globalState;
      let _lhs33 = globalState;
      let _lhs34 = p0;
      let _lhs35 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((p0).length));
      _lhs32.f14 = _rhs34;
      _lhs33.f7 = _rhs35;
      r0 = _rhs36;
      _161_v3 = _rhs37;
      _lhs34[_lhs35] = _rhs38;
      let _162_v5;
      _162_v5 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f23);
      let _hi4 = (_this).f23;
      for (let _163_i0 = new BigNumber((_dafny.Seq.of(new BigNumber((function () {
        let _coll10 = new _dafny.Set();
        for (const _compr_10 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(965)), function (_191_i1) {
          return (_this).f23;
        })).Elements) {
          let _192_v4 = _compr_10;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(965)), function (_191_i1) {
            return (_this).f23;
          }), _192_v4)) {
            _coll10.add((_192_v4).minus(new BigNumber(783)));
          }
        }
        return _coll10;
      }()).length), (_this).f23, (((_162_v5).contains(p0)) ? ((_162_v5).get(p0)) : ((_this).f23)))).length); _163_i0.isLessThan(_hi4); _163_i0 = _163_i0.plus(_dafny.ONE)) {
        if (p1) {
          (globalState).f14 = ((_dafny.Seq.IsPrefixOf(_158_v0, _158_v0)) ? ((_163_i0).multipliedBy((_this).f23)) : (_163_i0));
          let _164_v6;
          _164_v6 = _dafny.Seq.of(_dafny.Seq.update(_dafny.Seq.of(p1, (_this).f28, (p0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((p0).length))], _module.__default.fm1((_this).f23, (_this).f28, globalState), p1), _module.__default.safeIndex(_163_i0, new BigNumber((_dafny.Seq.of(p1, (_this).f28, (p0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((p0).length))], _module.__default.fm1((_this).f23, (_this).f28, globalState), p1)).length)), (_this).f28));
          let _165_v7;
          _165_v7 = _dafny.Map.Empty.slice().updateUnsafe(_163_i0,(_this).f23);
          let _166_v8;
          _166_v8 = _module.D1.create_DC4(_165_v7, (p0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((p0).length))], (_this).f28);
          let _167_v9;
          _167_v9 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_164_v6, _dafny.Seq.update(_164_v6, _module.__default.safeIndex((_this).f23, new BigNumber((_164_v6).length)), _dafny.Seq.of(p1, (p0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((p0).length))]))),(_166_v8).dtor_cf8);
          _167_v9 = (_167_v9).update(_164_v6, (_this).f28);
          let _168_v10;
          _168_v10 = _dafny.Map.Empty.slice().updateUnsafe((p0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((p0).length))],new BigNumber(-375));
          let _169_v11;
          _169_v11 = _dafny.Seq.of((_this).f28);
          let _rhs39 = ((_this).f23).plus((((_168_v10).contains(_module.__default.fm1(new BigNumber((_169_v11).length), (_this).f28, globalState))) ? ((_168_v10).get(_module.__default.fm1(new BigNumber((_169_v11).length), (_this).f28, globalState))) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-960)), function (_170_i2) {
            return new _dafny.CodePoint('v'.codePointAt(0));
          })).length))));
          let _rhs40 = ((_this).f23).isLessThan((_this).f23);
          let _lhs36 = globalState;
          let _lhs37 = globalState;
          _lhs36.f14 = _rhs39;
          _lhs37.f9 = _rhs40;
          let _171_v12;
          let _init0 = function (_172_i3) {
            return (_172_i3).multipliedBy(new BigNumber(635));
          };
          let _nw26 = Array((new BigNumber(13)).toNumber());
          for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw26.length); _i0_0++) {
            _nw26[_i0_0] = _init0(new BigNumber(_i0_0));
          }
          _171_v12 = _nw26;
          let _173_v13;
          _173_v13 = _dafny.Seq.of(_171_v12);
          let _174_v14;
          _174_v14 = _module.D1.create_DC3(_dafny.Seq.of(_171_v12, _171_v12, _171_v12));
          let _175_v15;
          let _nw27 = Array((new BigNumber(10)).toNumber());
          _nw27[(_dafny.ZERO).toNumber()] = _module.D1.create_DC3(_173_v13);
          _nw27[(_dafny.ONE).toNumber()] = _174_v14;
          _nw27[(new BigNumber(2)).toNumber()] = _174_v14;
          _nw27[(new BigNumber(3)).toNumber()] = _174_v14;
          _nw27[(new BigNumber(4)).toNumber()] = _module.D1.create_DC3(_173_v13);
          _nw27[(new BigNumber(5)).toNumber()] = _174_v14;
          _nw27[(new BigNumber(6)).toNumber()] = _174_v14;
          _nw27[(new BigNumber(7)).toNumber()] = _174_v14;
          _nw27[(new BigNumber(8)).toNumber()] = _174_v14;
          _nw27[(new BigNumber(9)).toNumber()] = _174_v14;
          _175_v15 = _nw27;
          let _176_v16;
          _176_v16 = _dafny.MultiSet.fromElements(_175_v15, _175_v15);
          let _177_v17;
          _177_v17 = _module.D3.create_DC9(_176_v16);
          let _178_v18;
          _178_v18 = _dafny.Seq.of(_177_v17, _177_v17, _177_v17, _177_v17);
          _178_v18 = _dafny.Seq.Concat((((_this).f28) ? (_178_v18) : (_178_v18)), _178_v18);
          (globalState).f21 = (_this).f28;
        } else {
          let _179_v19;
          let _180_v20;
          let _181_v21;
          let _out18;
          let _out19;
          let _out20;
          let _outcollector6 = _module.__default.m0(p0, _dafny.Seq.Concat(_module.__default.fm2(globalState), _158_v0), (_this.f29)[_module.__default.safeIndex((_this).f23, new BigNumber((_this.f29).length))], globalState);
          _out18 = _outcollector6[0];
          _out19 = _outcollector6[1];
          _out20 = _outcollector6[2];
          _179_v19 = _out18;
          _180_v20 = _out19;
          _181_v21 = _out20;
          let _182_v22;
          let _nw28 = Array((new BigNumber(25)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _182_v22 = _nw28;
          let _nw29 = Array((new BigNumber(24)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _182_v22 = _nw29;
          let _183_v23;
          _183_v23 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_158_v0).length),_181_v21);
          let _184_v24;
          _184_v24 = _dafny.Set.fromElements(new BigNumber((_183_v23).length), (_this).f23);
          let _index30 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((p0).length));
          let _rhs41 = p1;
          let _rhs42 = !((((_184_v24).IsProperSubsetOf(_184_v24)) ? (true) : (!_dafny.areEqual(_158_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(196)), ((_185_v2) => function (_186_i4) {
            return _185_v2;
          })(_160_v2))))));
          let _rhs43 = new BigNumber(595);
          let _lhs38 = p0;
          let _lhs39 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((p0).length));
          let _lhs40 = globalState;
          let _lhs41 = globalState;
          _lhs38[_lhs39] = _rhs41;
          _lhs40.f7 = _rhs42;
          _lhs41.f8 = _rhs43;
          let _187_v25;
          _187_v25 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_this.f29).length),_160_v2);
          (globalState).f11 = (((_187_v25).contains((_this).f23)) ? ((_187_v25).get((_this).f23)) : (_module.__default.fm6(_163_i0, globalState)));
          let _188_v26;
          let _nw30 = Array((new BigNumber(17)).toNumber()).fill(false);
          _188_v26 = _nw30;
          _188_v26 = _188_v26;
        }
        let _189_v27;
        _189_v27 = _dafny.Seq.of((_this).f28);
        let _190_v28;
        _190_v28 = _dafny.Seq.of(_189_v27, _189_v27);
        let _index31 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((p0).length));
        let _rhs44 = p1;
        let _rhs45 = _dafny.Seq.Concat(_190_v28, _dafny.Seq.update(_190_v28, _module.__default.safeIndex(_163_i0, new BigNumber((_190_v28).length)), _189_v27));
        let _rhs46 = (_189_v27)[_module.__default.safeIndex(_module.__default.safeDivisionInt((_this).f23, _163_i0), new BigNumber((_189_v27).length))];
        let _lhs42 = p0;
        let _lhs43 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((p0).length));
        let _lhs44 = globalState;
        _lhs42[_lhs43] = _rhs44;
        _190_v28 = _rhs45;
        _lhs44.f21 = _rhs46;
        (globalState).f8 = (((p0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((p0).length))]) ? ((new BigNumber(950)).plus((_this).f23)) : (_163_i0));
        (globalState).f21 = (_this).f28;
      }
      let _193_v29;
      _193_v29 = _dafny.Seq.of(false);
      let _194_v30;
      _194_v30 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_this).f23);
      let _195_v31;
      _195_v31 = _dafny.Seq.of(true, !(((_this).f23).isLessThanOrEqualTo(new BigNumber((_193_v29).length))), p1, (_this).f28, (_194_v30).contains(false));
      if ((_195_v31)[_module.__default.safeIndex((_this).f23, new BigNumber((_195_v31).length))]) {
        let _196_v32;
        let _nw31 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
        _196_v32 = _nw31;
        let _index32 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_196_v32).length));
        (_196_v32)[_index32] = new BigNumber((_this.f29).length);
        let _197_v33;
        _197_v33 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,_196_v32);
        let _198_v34;
        _198_v34 = _dafny.Map.Empty.slice().updateUnsafe(_197_v33,true);
        let _index33 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_196_v32).length));
        let _rhs47 = (_dafny.ZERO).minus((_this).f23);
        let _rhs48 = _dafny.Seq.Concat(_this.f29, _dafny.Seq.Create(_module.__default.abs(new BigNumber(288)), function (_199_i5) {
          return new BigNumber(715);
        }));
        let _rhs49 = (_dafny.ZERO).minus((_this).f23);
        let _rhs50 = (((_198_v34).contains(_197_v33)) ? ((_198_v34).get(_197_v33)) : (_module.__default.fm1((_this).f23, p1, globalState)));
        let _rhs51 = ((!(!(!((p0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((p0).length))]))) || ((_this).f28)) ? (new BigNumber((_195_v31).length)) : ((_this).f23));
        let _lhs45 = globalState;
        let _lhs46 = _this;
        let _lhs47 = _196_v32;
        let _lhs48 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_196_v32).length));
        let _lhs49 = globalState;
        _lhs45.f14 = _rhs47;
        _lhs46.f29 = _rhs48;
        _lhs47[_lhs48] = _rhs49;
        _lhs49.f21 = _rhs50;
        r0 = _rhs51;
        (globalState).f14 = (_196_v32)[_module.__default.safeIndex(new BigNumber(532), new BigNumber((_196_v32).length))];
        let _200_v35;
        _200_v35 = _dafny.Seq.of(_158_v0, _module.__default.fm2(globalState), _158_v0, _158_v0, _158_v0);
        let _201_v37;
        _201_v37 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(831),(_this).f28);
        _159_v1 = function () {
          let _coll11 = new _dafny.Set();
          for (const _compr_11 of ((function () {
            let _coll12 = new _dafny.Set();
            for (const _compr_12 of (_200_v35).Elements) {
              let _202_v36 = _compr_12;
              if (_dafny.Seq.contains(_200_v35, _202_v36)) {
                _coll12.add(_202_v36);
              }
            }
            return _coll12;
          }()).Difference(function () {
            let _coll13 = new _dafny.Set();
            for (const _compr_13 of (_dafny.Map.Empty.slice().updateUnsafe(_158_v0,(((_201_v37).contains(new BigNumber((_194_v30).length))) ? ((_201_v37).get(new BigNumber((_194_v30).length))) : (p1)))).Keys.Elements) {
              let _203_v38 = _compr_13;
              if ((_dafny.Map.Empty.slice().updateUnsafe(_158_v0,(((_201_v37).contains(new BigNumber((_194_v30).length))) ? ((_201_v37).get(new BigNumber((_194_v30).length))) : (p1)))).contains(_203_v38)) {
                _coll13.add(_203_v38);
              }
            }
            return _coll13;
          }())).Elements) {
            let _204_v39 = _compr_11;
            if (((function () {
              let _coll14 = new _dafny.Set();
              for (const _compr_14 of (_200_v35).Elements) {
                let _205_v36 = _compr_14;
                if (_dafny.Seq.contains(_200_v35, _205_v36)) {
                  _coll14.add(_205_v36);
                }
              }
              return _coll14;
            }()).Difference(function () {
              let _coll15 = new _dafny.Set();
              for (const _compr_15 of (_dafny.Map.Empty.slice().updateUnsafe(_158_v0,(((_201_v37).contains(new BigNumber((_194_v30).length))) ? ((_201_v37).get(new BigNumber((_194_v30).length))) : (p1)))).Keys.Elements) {
                let _206_v38 = _compr_15;
                if ((_dafny.Map.Empty.slice().updateUnsafe(_158_v0,(((_201_v37).contains(new BigNumber((_194_v30).length))) ? ((_201_v37).get(new BigNumber((_194_v30).length))) : (p1)))).contains(_206_v38)) {
                  _coll15.add(_206_v38);
                }
              }
              return _coll15;
            }())).contains(_204_v39)) {
              _coll11.add(_204_v39);
            }
          }
          return _coll11;
        }();
        let _207_v41;
        _207_v41 = _dafny.Set.fromElements(new BigNumber((function () {
          let _coll16 = new _dafny.Set();
          for (const _compr_16 of _dafny.IntegerRange(new BigNumber(784), new BigNumber(480))) {
            let _208_v40 = _compr_16;
            if (((new BigNumber(784)).isLessThanOrEqualTo(_208_v40)) && ((_208_v40).isLessThan(new BigNumber(480)))) {
              _coll16.add(_module.__default.safeDivisionInt(_208_v40, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f23,_dafny.MultiSet.FromArray(_this.f29))).length)));
            }
          }
          return _coll16;
        }()).length), (_196_v32)[_module.__default.safeIndex(new BigNumber(532), new BigNumber((_196_v32).length))]);
        r0 = _module.__default.safeDivisionInt(_module.__default.fm0(new BigNumber((_module.__default.fm5(_207_v41, globalState)).length), (p0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((p0).length))], globalState), ((_module.__default.fm1((_196_v32)[_module.__default.safeIndex(new BigNumber(532), new BigNumber((_196_v32).length))], p1, globalState)) ? ((_196_v32)[_module.__default.safeIndex(new BigNumber(532), new BigNumber((_196_v32).length))]) : ((_this).f23)));
        if ((p0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((p0).length))]) {
          let _index34 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((p0).length));
          (p0)[_index34] = (p0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((p0).length))];
          let _209_v42;
          let _nw32 = new _module.C0();
          _nw32.__ctor();
          _209_v42 = _nw32;
          let _210_v43;
          let _nw33 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _210_v43 = _nw33;
          let _211_v44;
          _211_v44 = _dafny.Seq.of(_193_v29, _195_v31, _dafny.Seq.of(p1));
          let _rhs52 = !_dafny.Seq.contains(_dafny.Seq.Concat((_211_v44)[_module.__default.safeIndex((_196_v32)[_module.__default.safeIndex(new BigNumber(532), new BigNumber((_196_v32).length))], new BigNumber((_211_v44).length))], _193_v29), true);
          let _rhs53 = _210_v43;
          let _lhs50 = globalState;
          _lhs50.f9 = _rhs52;
          _210_v43 = _rhs53;
          (globalState).f9 = p1;
          let _212_v45;
          _212_v45 = _dafny.Set.fromElements(_160_v2, _160_v2);
          let _213_v46;
          _213_v46 = _module.D2.create_DC7(_212_v45);
          let _214_v47;
          _214_v47 = _dafny.Map.Empty.slice().updateUnsafe(_213_v46,_158_v0);
          let _215_v48;
          _215_v48 = _dafny.MultiSet.fromElements((_209_v42).fm20(_214_v47, (_196_v32)[_module.__default.safeIndex(new BigNumber(532), new BigNumber((_196_v32).length))], new BigNumber((_158_v0).length), (_196_v32)[_module.__default.safeIndex(new BigNumber(532), new BigNumber((_196_v32).length))], globalState));
          let _216_v49;
          _216_v49 = _dafny.Map.Empty.slice().updateUnsafe(_215_v48,_dafny.Seq.UnicodeFromString("dux"));
          let _217_v50;
          _217_v50 = _module.D0.create_DC0(_158_v0);
          let _218_v51;
          _218_v51 = _dafny.Seq.of(_module.__default.fm4(_217_v50, (_this).f23, _dafny.Seq.of((_this).f28), _160_v2, globalState), _215_v48);
          _216_v49 = (_216_v49).update((_218_v51)[_module.__default.safeIndex((_this).f23, new BigNumber((_218_v51).length))], _158_v0);
        } else {
          let _219_v52;
          _219_v52 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,new BigNumber(275));
          _219_v52 = (_219_v52).update((_196_v32)[_module.__default.safeIndex(new BigNumber(532), new BigNumber((_196_v32).length))], (_196_v32)[_module.__default.safeIndex(new BigNumber(532), new BigNumber((_196_v32).length))]);
          _207_v41 = _module.__default.fm21(_200_v35, globalState);
          let _220_v53;
          let _221_v54;
          let _222_v55;
          let _out21;
          let _out22;
          let _out23;
          let _outcollector7 = _module.__default.m0(p0, _dafny.Seq.Concat(_dafny.Seq.update(_158_v0, _module.__default.safeIndex((_this).f23, new BigNumber((_158_v0).length)), _160_v2), _158_v0), (_this).f23, globalState);
          _out21 = _outcollector7[0];
          _out22 = _outcollector7[1];
          _out23 = _outcollector7[2];
          _220_v53 = _out21;
          _221_v54 = _out22;
          _222_v55 = _out23;
          let _index35 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_196_v32).length));
          (_196_v32)[_index35] = (_this).f23;
          (globalState).f9 = !((!((_196_v32)[_module.__default.safeIndex(new BigNumber(532), new BigNumber((_196_v32).length))]).isEqualTo(_222_v55)) === (((_220_v53).dtor_cf3).isEqualTo((_196_v32)[_module.__default.safeIndex(new BigNumber(532), new BigNumber((_196_v32).length))])));
        }
      } else {
        let _223_v56;
        let _nw34 = Array((new BigNumber(2)).toNumber()).fill([]);
        _223_v56 = _nw34;
        let _224_v57;
        let _init1 = function (_225_i6) {
          return _this.f29;
        };
        let _nw35 = Array((new BigNumber(19)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw35.length); _i0_1++) {
          _nw35[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _224_v57 = _nw35;
        let _226_v58;
        _226_v58 = _dafny.Seq.of(_224_v57, _224_v57);
        let _index36 = _module.__default.safeIndex(new BigNumber(242), new BigNumber((_223_v56).length));
        (_223_v56)[_index36] = (_226_v58)[_module.__default.safeIndex((_this).f23, new BigNumber((_226_v58).length))];
        let _index37 = _module.__default.safeIndex(new BigNumber(242), new BigNumber((_223_v56).length));
        (_223_v56)[_index37] = _224_v57;
        let _227_v59;
        let _nw36 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
        _227_v59 = _nw36;
        let _index38 = _module.__default.safeIndex(new BigNumber(461), new BigNumber((_227_v59).length));
        (_227_v59)[_index38] = (_this).f23;
        let _index39 = _module.__default.safeIndex(new BigNumber(461), new BigNumber((_227_v59).length));
        (_227_v59)[_index39] = (_this).f23;
        let _228_v60;
        _228_v60 = _dafny.Set.fromElements(!((_this).f28), (p0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((p0).length))], false, p1);
        let _index40 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((p0).length));
        let _index41 = _module.__default.safeIndex(new BigNumber(461), new BigNumber((_227_v59).length));
        let _rhs54 = _module.__default.fm1(_module.__default.fm0((_dafny.ZERO).minus(new BigNumber((_228_v60).length)), p1, globalState), (_this).f28, globalState);
        let _rhs55 = !((_this).f28);
        let _rhs56 = (_this).f23;
        let _rhs57 = (p0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((p0).length))];
        let _rhs58 = _module.__default.safeModuloInt(_module.__default.safeModuloInt((_227_v59)[_module.__default.safeIndex(new BigNumber(461), new BigNumber((_227_v59).length))], new BigNumber((_158_v0).length)), _module.__default.safeDivisionInt((_this).f23, _module.__default.fm0((_227_v59)[_module.__default.safeIndex(new BigNumber(461), new BigNumber((_227_v59).length))], p1, globalState)));
        let _lhs51 = globalState;
        let _lhs52 = p0;
        let _lhs53 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((p0).length));
        let _lhs54 = globalState;
        let _lhs55 = _227_v59;
        let _lhs56 = _module.__default.safeIndex(new BigNumber(461), new BigNumber((_227_v59).length));
        _lhs51.f9 = _rhs54;
        _lhs52[_lhs53] = _rhs55;
        r0 = _rhs56;
        _lhs54.f9 = _rhs57;
        _lhs55[_lhs56] = _rhs58;
        let _229_v61;
        _229_v61 = _module.D3.create_DC12(_195_v31);
        let _230_v62;
        _230_v62 = _module.D3.create_DC10(true);
        _229_v61 = _module.__default.fm22(_230_v62, globalState);
        let _index42 = _module.__default.safeIndex(new BigNumber(461), new BigNumber((_227_v59).length));
        (_227_v59)[_index42] = (_this.f29)[_module.__default.safeIndex((_227_v59)[_module.__default.safeIndex(new BigNumber(461), new BigNumber((_227_v59).length))], new BigNumber((_this.f29).length))];
      }
      let _hi5 = (_this).f23;
      for (let _231_i7 = (_this).f23; _231_i7.isLessThan(_hi5); _231_i7 = _231_i7.plus(_dafny.ONE)) {
        let _232_v63;
        let _nw37 = Array((_dafny.ONE).toNumber());
        _nw37[(_dafny.ZERO).toNumber()] = (p0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((p0).length))];
        _232_v63 = _nw37;
        _232_v63 = _232_v63;
        (globalState).f11 = _160_v2;
        _194_v30 = (_194_v30).Merge(_194_v30);
        let _index43 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((p0).length));
        let _rhs59 = _module.__default.fm1((_this).f23, (p0)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((p0).length))], globalState);
        let _rhs60 = (_this).f28;
        let _rhs61 = (_193_v29)[_module.__default.safeIndex(_231_i7, new BigNumber((_193_v29).length))];
        let _lhs57 = p0;
        let _lhs58 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((p0).length));
        let _lhs59 = globalState;
        let _lhs60 = globalState;
        _lhs57[_lhs58] = _rhs59;
        _lhs59.f7 = _rhs60;
        _lhs60.f7 = _rhs61;
      }
      r0 = (_this).f23;
      let _233_v64;
      _233_v64 = _dafny.Set.fromElements((_dafny.ZERO).minus(((_this).f23).multipliedBy((_this).f23)), (_dafny.ZERO).minus(new BigNumber((_158_v0).length)));
      r1 = _233_v64;
      return [r0, r1];
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Seq.of();
      let _rhs62 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("hoiflwc"), p0);
      let _rhs63 = _dafny.Seq.Concat(p0, p0);
      let _rhs64 = new BigNumber(641);
      let _rhs65 = ((_this).f23).plus((_this).f23);
      let _lhs61 = globalState;
      let _lhs62 = globalState;
      _lhs61.f6 = _rhs62;
      _lhs62.f6 = _rhs63;
      r0 = _rhs64;
      r0 = _rhs65;
      let _234_v0;
      _234_v0 = _dafny.MultiSet.fromElements((_this).f28);
      let _235_v1;
      _235_v1 = _dafny.Map.Empty.slice().updateUnsafe((_234_v0).IsSubsetOf(_dafny.MultiSet.fromElements((_this).f28)),(_dafny.ZERO).minus((_this).f23));
      _235_v1 = (_235_v1).update(((_this).f28) || ((_this).f28), new BigNumber(922));
      (globalState).f14 = (_dafny.ZERO).minus((_this).f23);
      let _236_v2;
      _236_v2 = _dafny.MultiSet.fromElements((_this).f23);
      let _237_v3;
      _237_v3 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(308),_236_v2));
      let _238_v4;
      _238_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,_236_v2);
      let _239_v5;
      _239_v5 = _dafny.Seq.of((_237_v3)[_module.__default.safeIndex((_this).f23, new BigNumber((_237_v3).length))], (_238_v4).update(new BigNumber(324), _236_v2), _238_v4);
      let _240_v6;
      _240_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,new BigNumber(((_239_v5)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f23), new BigNumber((_239_v5).length))]).length));
      let _241_v7;
      _241_v7 = new _dafny.CodePoint('n'.codePointAt(0));
      let _242_v8;
      _242_v8 = _module.D2.create_DC8((_this).f23, new BigNumber((_240_v6).length), _241_v7, (_this).f23, (_this).f28);
      _235_v1 = (_235_v1).update((_242_v8).dtor_cf19, (_dafny.ZERO).minus((_this).f23));
      let _243_v9;
      let _nw38 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
      _243_v9 = _nw38;
      let _hi6 = (_this).f23;
      for (let _244_i0 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_243_v9)).length)); _244_i0.isLessThan(_hi6); _244_i0 = _244_i0.plus(_dafny.ONE)) {
        (globalState).f14 = (_244_i0).multipliedBy((_this).f23);
        if (!(_module.__default.fm1(_244_i0, (_this).f28, globalState))) {
          (globalState).f21 = !((_this).f28);
          let _245_v10;
          _245_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_this).f28);
          let _246_v11;
          _246_v11 = _dafny.Seq.of((_this).f28, (_this).f28);
          _245_v10 = (_245_v10).update((_this).f28, (_246_v11)[_module.__default.safeIndex(_244_i0, new BigNumber((_246_v11).length))]);
          let _247_v12;
          let _nw39 = Array((new BigNumber(19)).toNumber()).fill([]);
          _247_v12 = _nw39;
          let _248_v13;
          let _nw40 = Array((new BigNumber(3)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _248_v13 = _nw40;
          let _index44 = _module.__default.safeIndex(new BigNumber(536), new BigNumber((_247_v12).length));
          (_247_v12)[_index44] = _248_v13;
          let _index45 = _module.__default.safeIndex(new BigNumber(536), new BigNumber((_247_v12).length));
          let _nw41 = Array((new BigNumber(23)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          (_247_v12)[_index45] = _nw41;
          r0 = (_this).f23;
          let _249_v14;
          let _init2 = ((_250_p0, _251_i0, _252_v7) => function (_253_i1) {
            return _dafny.Seq.update(_250_p0, _module.__default.safeIndex(_251_i0, new BigNumber((_250_p0).length)), _252_v7);
          })(p0, _244_i0, _241_v7);
          let _nw42 = Array((new BigNumber(20)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw42.length); _i0_2++) {
            _nw42[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _249_v14 = _nw42;
          let _index46 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_249_v14).length));
          (_249_v14)[_index46] = _dafny.Seq.Concat(_dafny.Seq.update(p0, _module.__default.safeIndex(_244_i0, new BigNumber((p0).length)), _241_v7), p0);
          let _254_v15;
          _254_v15 = _dafny.Seq.of(_dafny.Seq.update(_dafny.Seq.update(p0, _module.__default.safeIndex(_244_i0, new BigNumber((p0).length)), _241_v7), _module.__default.safeIndex(_244_i0, new BigNumber((_dafny.Seq.update(p0, _module.__default.safeIndex(_244_i0, new BigNumber((p0).length)), _241_v7)).length)), _241_v7));
          let _index47 = _module.__default.safeIndex(new BigNumber(363), new BigNumber((_249_v14).length));
          (_249_v14)[_index47] = _dafny.Seq.Concat((_254_v15)[_module.__default.safeIndex(_244_i0, new BigNumber((_254_v15).length))], p0);
          let _255_v16;
          _255_v16 = _dafny.Seq.of(_234_v0);
          let _256_v17;
          let _nw43 = Array((new BigNumber(21)).toNumber());
          _nw43[(_dafny.ZERO).toNumber()] = _234_v0;
          _nw43[(_dafny.ONE).toNumber()] = _234_v0;
          _nw43[(new BigNumber(2)).toNumber()] = _dafny.MultiSet.fromElements((_this).f28, (_this).f28, (_this).f28, (_this).f28);
          _nw43[(new BigNumber(3)).toNumber()] = _234_v0;
          _nw43[(new BigNumber(4)).toNumber()] = _234_v0;
          _nw43[(new BigNumber(5)).toNumber()] = _234_v0;
          _nw43[(new BigNumber(6)).toNumber()] = (((_234_v0).update((_this).f28, _module.__default.abs((_this).f23))).update((_this).f28, _module.__default.abs((_this).f23))).Union(_dafny.MultiSet.fromElements((_this).f28, true));
          _nw43[(new BigNumber(7)).toNumber()] = _234_v0;
          _nw43[(new BigNumber(8)).toNumber()] = _234_v0;
          _nw43[(new BigNumber(9)).toNumber()] = _234_v0;
          _nw43[(new BigNumber(10)).toNumber()] = _234_v0;
          _nw43[(new BigNumber(11)).toNumber()] = (_234_v0).Intersect(_234_v0);
          _nw43[(new BigNumber(12)).toNumber()] = (_234_v0).update((_this).f28, _module.__default.abs((_dafny.ZERO).minus((_this).f23)));
          _nw43[(new BigNumber(13)).toNumber()] = (_234_v0).Union(_dafny.MultiSet.fromElements((_this).f28));
          _nw43[(new BigNumber(14)).toNumber()] = _234_v0;
          _nw43[(new BigNumber(15)).toNumber()] = _234_v0;
          _nw43[(new BigNumber(16)).toNumber()] = _234_v0;
          _nw43[(new BigNumber(17)).toNumber()] = (_234_v0).Difference(_234_v0);
          _nw43[(new BigNumber(18)).toNumber()] = (_255_v16)[_module.__default.safeIndex(_244_i0, new BigNumber((_255_v16).length))];
          _nw43[(new BigNumber(19)).toNumber()] = _234_v0;
          _nw43[(new BigNumber(20)).toNumber()] = (_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f28))).Intersect(_234_v0);
          _256_v17 = _nw43;
          let _index48 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_256_v17).length));
          (_256_v17)[_index48] = _234_v0;
          let _257_v18;
          _257_v18 = _dafny.Set.fromElements(!((_this).f28) || (!((_this).f28)));
          let _index49 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_249_v14).length));
          let _index50 = _module.__default.safeIndex(new BigNumber(363), new BigNumber((_249_v14).length));
          let _index51 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_256_v17).length));
          let _rhs66 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("bshdxkt"), p0);
          let _rhs67 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(128)), ((_258_v7) => function (_259_i2) {
            return _258_v7;
          })(_241_v7)), _dafny.Seq.UnicodeFromString("sgs"));
          let _rhs68 = _234_v0;
          let _rhs69 = _257_v18;
          let _rhs70 = p0;
          let _lhs63 = _249_v14;
          let _lhs64 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_249_v14).length));
          let _lhs65 = _249_v14;
          let _lhs66 = _module.__default.safeIndex(new BigNumber(363), new BigNumber((_249_v14).length));
          let _lhs67 = _256_v17;
          let _lhs68 = _module.__default.safeIndex(new BigNumber(314), new BigNumber((_256_v17).length));
          let _lhs69 = globalState;
          _lhs63[_lhs64] = _rhs66;
          _lhs65[_lhs66] = _rhs67;
          _lhs67[_lhs68] = _rhs68;
          _257_v18 = _rhs69;
          _lhs69.f6 = _rhs70;
        } else {
          (globalState).f21 = (_this).f28;
          let _index52 = _module.__default.safeIndex(new BigNumber(258), new BigNumber((_243_v9).length));
          (_243_v9)[_index52] = _244_i0;
          let _index53 = _module.__default.safeIndex(new BigNumber(258), new BigNumber((_243_v9).length));
          let _rhs71 = _dafny.Seq.UnicodeFromString("celkv");
          let _rhs72 = new BigNumber((_dafny.Seq.of((_this).f28)).length);
          let _lhs70 = globalState;
          let _lhs71 = _243_v9;
          let _lhs72 = _module.__default.safeIndex(new BigNumber(258), new BigNumber((_243_v9).length));
          _lhs70.f6 = _rhs71;
          _lhs71[_lhs72] = _rhs72;
          let _260_v19;
          _260_v19 = _module.D0.create_DC1((_this).f23, (_this).f28, (((_236_v2).contains((_this).f23)) ? ((_236_v2).get((_this).f23)) : (new BigNumber((_236_v2).cardinality()))));
          (globalState).f9 = (_260_v19).dtor_cf2;
          (globalState).f21 = (_this).f28;
          let _261_v20;
          _261_v20 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,((_module.__default.fm1((_dafny.ZERO).minus((_243_v9)[_module.__default.safeIndex(new BigNumber(258), new BigNumber((_243_v9).length))]), (_this).f28, globalState)) ? ((_this).f28) : ((_this).f28)));
          _261_v20 = (_261_v20).update((_this).f23, (_this).f28);
        }
        let _262_v21;
        _262_v21 = _module.D4.create_DC15(new BigNumber(-165), _241_v7, _module.__default.safeModuloInt((_this.f29)[_module.__default.safeIndex(new BigNumber((_236_v2).cardinality()), new BigNumber((_this.f29).length))], (_this).f23));
        _262_v21 = _262_v21;
        let _263_v22;
        let _nw44 = Array((new BigNumber(22)).toNumber()).fill(false);
        _263_v22 = _nw44;
        let _index54 = _module.__default.safeIndex(new BigNumber(424), new BigNumber((_263_v22).length));
        (_263_v22)[_index54] = (_this).f28;
        let _index55 = _module.__default.safeIndex(new BigNumber(424), new BigNumber((_263_v22).length));
        (_263_v22)[_index55] = (_242_v8).dtor_cf19;
      }
      let _264_i3;
      _264_i3 = _dafny.ZERO;
      L3: {
        while (!((((_this).f28) ? (!(new BigNumber(628)).isEqualTo((_module.D1.create_DC5((_this).f23, (_this).f28, (_this).f23, (_this).f23)).dtor_cf9)) : ((_this).f28)))) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_264_i3)) {
              break L3;
            }
            _264_i3 = (_264_i3).plus(_dafny.ONE);
            let _265_v23;
            _265_v23 = _dafny.Seq.of((_this).f28);
            let _266_v24;
            _266_v24 = _module.D3.create_DC10((_265_v23)[_module.__default.safeIndex((_this).f23, new BigNumber((_265_v23).length))]);
            let _267_v25;
            _267_v25 = _dafny.Set.fromElements((_this).f28);
            let _268_v26;
            _268_v26 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(329),(_this).f28);
            let _269_v27;
            _269_v27 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,(((_268_v26).contains((_this).f23)) ? ((_268_v26).get((_this).f23)) : ((_this).f28)));
            let _270_v28;
            let _nw45 = Array((new BigNumber(29)).toNumber());
            _nw45[(_dafny.ZERO).toNumber()] = (_this).f28;
            _nw45[(_dafny.ONE).toNumber()] = _module.__default.fm1((_this).f23, (_this).f28, globalState);
            _nw45[(new BigNumber(2)).toNumber()] = (_this).f28;
            _nw45[(new BigNumber(3)).toNumber()] = (((_this).f28) ? ((_this).f28) : ((_this).f28));
            _nw45[(new BigNumber(4)).toNumber()] = (_this).f28;
            _nw45[(new BigNumber(5)).toNumber()] = (_this).f28;
            _nw45[(new BigNumber(6)).toNumber()] = (_266_v24).dtor_cf21;
            _nw45[(new BigNumber(7)).toNumber()] = (_this).f28;
            _nw45[(new BigNumber(8)).toNumber()] = (_this).f28;
            _nw45[(new BigNumber(9)).toNumber()] = (_this).f28;
            _nw45[(new BigNumber(10)).toNumber()] = (_this).f28;
            _nw45[(new BigNumber(11)).toNumber()] = (_this).f28;
            _nw45[(new BigNumber(12)).toNumber()] = !((_265_v23)[_module.__default.safeIndex((_this).f23, new BigNumber((_265_v23).length))]) || (_module.__default.fm1((_this).f23, (_this).f28, globalState));
            _nw45[(new BigNumber(13)).toNumber()] = (((_this).f28) ? ((_this).f28) : ((_this).f28));
            _nw45[(new BigNumber(14)).toNumber()] = (_267_v25).equals(_dafny.Set.fromElements((_this).f28));
            _nw45[(new BigNumber(15)).toNumber()] = (_this).f28;
            _nw45[(new BigNumber(16)).toNumber()] = (_this).f28;
            _nw45[(new BigNumber(17)).toNumber()] = (_this).f28;
            _nw45[(new BigNumber(18)).toNumber()] = (_this).f28;
            _nw45[(new BigNumber(19)).toNumber()] = (((_269_v27).contains((_this).f28)) ? ((_269_v27).get((_this).f28)) : ((_this).f28));
            _nw45[(new BigNumber(20)).toNumber()] = !_dafny.Seq.contains(_dafny.Seq.update(_this.f29, _module.__default.safeIndex((_this).f23, new BigNumber((_this.f29).length)), new BigNumber(128)), (_this).f23);
            _nw45[(new BigNumber(21)).toNumber()] = !(((_dafny.ZERO).minus((_this).f23)).isLessThan((_this).f23));
            _nw45[(new BigNumber(22)).toNumber()] = (_this).f28;
            _nw45[(new BigNumber(23)).toNumber()] = (_this).f28;
            _nw45[(new BigNumber(24)).toNumber()] = (new BigNumber(-581)).isLessThan((((_235_v1).contains((_this).f28)) ? ((_235_v1).get((_this).f28)) : ((_this).f23)));
            _nw45[(new BigNumber(25)).toNumber()] = (_this).f28;
            _nw45[(new BigNumber(26)).toNumber()] = (_265_v23)[_module.__default.safeIndex((_this).f23, new BigNumber((_265_v23).length))];
            _nw45[(new BigNumber(27)).toNumber()] = !((_this).f28);
            _nw45[(new BigNumber(28)).toNumber()] = ((_this).f23).isLessThan((_this).f23);
            _270_v28 = _nw45;
            let _index56 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_270_v28).length));
            (_270_v28)[_index56] = ((_module.__default.fm1((_this).f23, (_this).f28, globalState)) ? ((_this).f28) : ((_this).f28));
            let _271_v29;
            _271_v29 = _dafny.MultiSet.fromElements(_243_v9, _243_v9, _243_v9);
            let _index57 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_270_v28).length));
            let _rhs73 = (_dafny.MultiSet.fromElements(_243_v9, _243_v9)).IsSubsetOf(_271_v29);
            let _rhs74 = (_this).f23;
            let _rhs75 = _module.__default.fm1((_this).f23, (_this).f28, globalState);
            let _rhs76 = !((_this).f28) || ((_this).f28);
            let _rhs77 = _module.__default.fm1((_this).f23, (_this).f28, globalState);
            let _lhs73 = _270_v28;
            let _lhs74 = _module.__default.safeIndex(new BigNumber(547), new BigNumber((_270_v28).length));
            let _lhs75 = globalState;
            let _lhs76 = globalState;
            let _lhs77 = globalState;
            _lhs73[_lhs74] = _rhs73;
            r0 = _rhs74;
            _lhs75.f21 = _rhs75;
            _lhs76.f7 = _rhs76;
            _lhs77.f21 = _rhs77;
            let _272_v30;
            _272_v30 = _dafny.Map.Empty.slice().updateUnsafe(((!(true)) ? (_265_v23) : (_dafny.Seq.of((_this).f28))),(_this).f28);
            let _273_v31;
            _273_v31 = _dafny.Map.Empty.slice().updateUnsafe(_241_v7,p0);
            _272_v30 = (_272_v30).update(_265_v23, !((_270_v28)[_module.__default.safeIndex(new BigNumber(547), new BigNumber((_270_v28).length))]) || ((_265_v23)[_module.__default.safeIndex(new BigNumber((_273_v31).length), new BigNumber((_265_v23).length))]));
            let _274_v32;
            let _nw46 = new _module.C0();
            _nw46.__ctor();
            _274_v32 = _nw46;
            let _275_v33;
            let _nw47 = new _module.C0();
            _nw47.__ctor();
            _275_v33 = _nw47;
          }
        }
      }
      r0 = (((_this).f23).minus((_this).f23)).plus(new BigNumber(712));
      r1 = _dafny.Seq.Concat(_this.f29, ((_module.__default.fm1((_this).f23, !((_this).f28), globalState)) ? (_this.f29) : (_this.f29)));
      return [r0, r1];
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f29 = _dafny.Seq.of();
      this._f23 = _dafny.ZERO;
      this._f28 = false;
      this._f31 = _dafny.Map.Empty;
      this._f32 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f29() {
      let _this = this;
      return _this._f29;
    };
    set f29(value) {
      let _this = this;
      _this._f29 = value;
    };
    get f23() {
      let _this = this;
      return _this._f23;
    };
    get f28() {
      let _this = this;
      return _this._f28;
    };
    __ctor(f31, f32, f28, f29, f23) {
      let _this = this;
      (_this)._f31 = f31;
      (_this)._f32 = f32;
      (_this)._f28 = f28;
      (_this)._f29 = f29;
      (_this)._f23 = f23;
      return;
    }
    fm18(p0, p1, p2, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(new BigNumber((((((_this).f28) ? (_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(779)), function (_276_i0) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      }))) : (_dafny.MultiSet.fromElements(new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)))))).Intersect((_dafny.MultiSet.fromElements(new _dafny.CodePoint('s'.codePointAt(0)))).Intersect(_dafny.MultiSet.fromElements(new _dafny.CodePoint('n'.codePointAt(0)))))).cardinality()));
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Set.Empty;
      let _277_v0;
      _277_v0 = _dafny.Seq.UnicodeFromString("rovoqj");
      let _278_v1;
      _278_v1 = _dafny.Map.Empty.slice().updateUnsafe(_277_v0,true);
      let _279_v2;
      _279_v2 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_277_v0, _277_v0),(_278_v1).Merge(_278_v1));
      _279_v2 = (_279_v2).update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-392)), function (_280_i0) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      }), (_278_v1).Merge(_278_v1));
      let _281_v3;
      let _init3 = ((_282_p1) => function (_283_i1) {
        return _dafny.Seq.Concat(_dafny.Seq.of((_this).f32, _282_p1), _dafny.Seq.of(_282_p1));
      })(p1);
      let _nw48 = Array((new BigNumber(26)).toNumber());
      for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw48.length); _i0_3++) {
        _nw48[_i0_3] = _init3(new BigNumber(_i0_3));
      }
      _281_v3 = _nw48;
      _281_v3 = _281_v3;
      let _284_i2;
      _284_i2 = _dafny.ZERO;
      L4: {
        while (((_this).f23).isLessThanOrEqualTo((_this).f23)) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_284_i2)) {
              break L4;
            }
            _284_i2 = (_284_i2).plus(_dafny.ONE);
            let _285_v4;
            _285_v4 = _dafny.MultiSet.fromElements(_277_v0, _277_v0);
            let _286_v5;
            let _nw49 = new _module.C1();
            _nw49.__ctor(false, _this.f29, new BigNumber((_277_v0).length));
            _286_v5 = _nw49;
            let _287_v6;
            _287_v6 = _dafny.MultiSet.fromElements(_286_v5, _286_v5, _286_v5);
            let _288_v7;
            _288_v7 = new _dafny.CodePoint('f'.codePointAt(0));
            let _289_v8;
            let _nw50 = Array((new BigNumber(11)).toNumber());
            _nw50[(_dafny.ZERO).toNumber()] = (_this).f23;
            _nw50[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus((((_285_v4).contains(_277_v0)) ? ((_285_v4).get(_277_v0)) : ((_this).f23)));
            _nw50[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.Set.fromElements(new BigNumber(-334), new BigNumber((_287_v6).cardinality()))).length);
            _nw50[(new BigNumber(3)).toNumber()] = new BigNumber(118);
            _nw50[(new BigNumber(4)).toNumber()] = (_286_v5).f23;
            _nw50[(new BigNumber(5)).toNumber()] = (_this.f29)[_module.__default.safeIndex(new BigNumber(248), new BigNumber((_this.f29).length))];
            _nw50[(new BigNumber(6)).toNumber()] = (_this).f23;
            _nw50[(new BigNumber(7)).toNumber()] = (_this).f23;
            _nw50[(new BigNumber(8)).toNumber()] = (_this).f23;
            _nw50[(new BigNumber(9)).toNumber()] = (_this).f23;
            _nw50[(new BigNumber(10)).toNumber()] = new BigNumber((_dafny.Seq.of(_288_v7)).length);
            _289_v8 = _nw50;
            let _290_v9;
            _290_v9 = _dafny.Seq.of(_289_v8, _289_v8, _289_v8, _289_v8, _289_v8);
            let _291_v10;
            _291_v10 = _dafny.Map.Empty.slice().updateUnsafe(_290_v9,(_286_v5).f23);
            _291_v10 = (_291_v10).update(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of(_289_v8), _290_v9), _module.__default.safeIndex((_286_v5).f23, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(_289_v8), _290_v9)).length)), _289_v8), (_this).f23);
            let _292_v11;
            let _nw51 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.of());
            _292_v11 = _nw51;
            let _293_v12;
            let _nw52 = new _module.C0();
            _nw52.__ctor();
            _293_v12 = _nw52;
            let _294_v13;
            _294_v13 = _dafny.Seq.of(_293_v12);
            let _index58 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_292_v11).length));
            (_292_v11)[_index58] = _294_v13;
            let _index59 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_292_v11).length));
            (_292_v11)[_index59] = (((_this).f28) ? (_294_v13) : (_294_v13));
            let _295_v14;
            _295_v14 = _module.D1.create_DC4((_this).f31, p1, true);
            if ((_295_v14).dtor_cf8) {
              let _296_v15;
              _296_v15 = _module.D1.create_DC3(_290_v9);
              _296_v15 = _296_v15;
              (globalState).f9 = p1;
              let _297_v16;
              _297_v16 = _dafny.MultiSet.fromElements((_this).f28);
              (globalState).f9 = ((_297_v16).IsSubsetOf(_297_v16)) || (p1);
              let _rhs78 = !((_module.__default.safeModuloInt((_286_v5).f23, (_286_v5).f23)).isLessThanOrEqualTo((_286_v5).f23));
              let _rhs79 = (_dafny.ZERO).minus((_this).f23);
              let _rhs80 = (_286_v5).f23;
              let _lhs78 = globalState;
              let _lhs79 = globalState;
              let _lhs80 = globalState;
              _lhs78.f7 = _rhs78;
              _lhs79.f8 = _rhs79;
              _lhs80.f8 = _rhs80;
              (globalState).f7 = ((_this).f32) === (_dafny.Seq.IsProperPrefixOf(_277_v0, _dafny.Seq.update(_277_v0, _module.__default.safeIndex((_286_v5).f23, new BigNumber((_277_v0).length)), _288_v7)));
            } else {
              let _298_v18;
              _298_v18 = _dafny.Seq.of(_277_v0, _dafny.Seq.UnicodeFromString("tocc"));
              let _299_v19;
              _299_v19 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt((_286_v5).f23, (_286_v5).f23),function () {
                let _coll17 = new _dafny.Map();
                for (const _compr_17 of (_dafny.Seq.update(_298_v18, _module.__default.safeIndex(new BigNumber(474), new BigNumber((_298_v18).length)), _277_v0)).Elements) {
                  let _300_v17 = _compr_17;
                  if (_dafny.Seq.contains(_dafny.Seq.update(_298_v18, _module.__default.safeIndex(new BigNumber(474), new BigNumber((_298_v18).length)), _277_v0), _300_v17)) {
                    _coll17.push([_300_v17,new BigNumber((_dafny.MultiSet.fromElements((_this).f28)).cardinality())]);
                  }
                }
                return _coll17;
              }());
              let _301_v20;
              _301_v20 = _dafny.Map.Empty.slice().updateUnsafe(_277_v0,(_dafny.ZERO).minus((_286_v5).f23));
              _299_v19 = (_299_v19).update((_this).f23, _301_v20);
              let _302_v21;
              _302_v21 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0((_286_v5).f23, p1, globalState),p1);
              _302_v21 = (_302_v21).update(((_this).f23).multipliedBy((_286_v5).f23), _module.__default.fm1((_286_v5).f23, (_this).f28, globalState));
              (globalState).f14 = (_module.__default.fm0((_286_v5).f23, p1, globalState)).multipliedBy((new BigNumber(-677)).multipliedBy((_286_v5).f23));
              let _303_v22;
              _303_v22 = _dafny.Map.Empty.slice().updateUnsafe((_this).f32,new BigNumber(368));
              _303_v22 = (_303_v22).update(p1, (new BigNumber((_module.__default.fm2(globalState)).length)).minus((_286_v5).f23));
              let _rhs81 = _dafny.Seq.Concat(_dafny.Seq.Concat(_277_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-846)), ((_304_v7) => function (_305_i3) {
                return _304_v7;
              })(_288_v7))), _277_v0);
              let _rhs82 = (_this).f23;
              let _lhs81 = globalState;
              _277_v0 = _rhs81;
              _lhs81.f14 = _rhs82;
            }
            let _306_v23;
            let _nw53 = Array((new BigNumber(2)).toNumber()).fill(_module.D0.Default());
            _306_v23 = _nw53;
            let _307_v24;
            _307_v24 = _module.D0.create_DC1((_this).f23, (_this).f32, (_this).f23);
            let _index60 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_306_v23).length));
            (_306_v23)[_index60] = _307_v24;
            let _index61 = _module.__default.safeIndex(new BigNumber(326), new BigNumber((_306_v23).length));
            (_306_v23)[_index61] = _307_v24;
          }
        }
      }
      let _308_v25;
      _308_v25 = _dafny.Set.fromElements(p1);
      _308_v25 = _308_v25;
      let _309_i4;
      _309_i4 = _dafny.ZERO;
      L5: {
        while ((_this).f28) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_309_i4)) {
              break L5;
            }
            _309_i4 = (_309_i4).plus(_dafny.ONE);
            let _310_v26;
            _310_v26 = _dafny.Seq.of((_this).f32, (_this).f32);
            let _311_v27;
            _311_v27 = _dafny.MultiSet.fromElements(_module.__default.fm0((_this).f23, p1, globalState), (_this).f23);
            let _312_v28;
            _312_v28 = _dafny.Seq.of(p0);
            let _313_v29;
            _313_v29 = _dafny.MultiSet.fromElements(_277_v0);
            let _314_v30;
            let _nw54 = Array((new BigNumber(26)).toNumber());
            _nw54[(_dafny.ZERO).toNumber()] = _module.__default.fm0((_this).f23, (_this).f32, globalState);
            _nw54[(_dafny.ONE).toNumber()] = (_this).f23;
            _nw54[(new BigNumber(2)).toNumber()] = new BigNumber(((((_310_v26)[_module.__default.safeIndex((_this).f23, new BigNumber((_310_v26).length))]) ? (_277_v0) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(11)), function (_315_i5) {
              return new _dafny.CodePoint('k'.codePointAt(0));
            })))).length);
            _nw54[(new BigNumber(3)).toNumber()] = (_this).f23;
            _nw54[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("hpcxes")).length);
            _nw54[(new BigNumber(5)).toNumber()] = (_this).f23;
            _nw54[(new BigNumber(6)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(481)), function (_316_i6) {
              return new _dafny.CodePoint('g'.codePointAt(0));
            }), _277_v0)).length);
            _nw54[(new BigNumber(7)).toNumber()] = (_this).f23;
            _nw54[(new BigNumber(8)).toNumber()] = (_this).f23;
            _nw54[(new BigNumber(9)).toNumber()] = (_this).f23;
            _nw54[(new BigNumber(10)).toNumber()] = _module.__default.safeDivisionInt((_this).f23, (_this).f23);
            _nw54[(new BigNumber(11)).toNumber()] = (_this).f23;
            _nw54[(new BigNumber(12)).toNumber()] = ((p1) ? ((_this).f23) : (new BigNumber(222)));
            _nw54[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus((_this).f23);
            _nw54[(new BigNumber(14)).toNumber()] = new BigNumber(583);
            _nw54[(new BigNumber(15)).toNumber()] = (_this).f23;
            _nw54[(new BigNumber(16)).toNumber()] = (new BigNumber((_277_v0).length)).multipliedBy(new BigNumber(((_dafny.MultiSet.fromElements(p1)).update(p1, _module.__default.abs((_this).f23))).cardinality()));
            _nw54[(new BigNumber(17)).toNumber()] = (_this).f23;
            _nw54[(new BigNumber(18)).toNumber()] = (_this).f23;
            _nw54[(new BigNumber(19)).toNumber()] = ((_this).f23).multipliedBy((_this).f23);
            _nw54[(new BigNumber(20)).toNumber()] = new BigNumber((_308_v25).length);
            _nw54[(new BigNumber(21)).toNumber()] = (_this).f23;
            _nw54[(new BigNumber(22)).toNumber()] = (((_311_v27).contains(new BigNumber((_312_v28).length))) ? ((_311_v27).get(new BigNumber((_312_v28).length))) : ((_this).f23));
            _nw54[(new BigNumber(23)).toNumber()] = new BigNumber((_313_v29).cardinality());
            _nw54[(new BigNumber(24)).toNumber()] = (_this).f23;
            _nw54[(new BigNumber(25)).toNumber()] = (_this).f23;
            _314_v30 = _nw54;
            _314_v30 = _314_v30;
            let _317_v31;
            _317_v31 = _dafny.MultiSet.fromElements((_this).f32, false);
            let _318_v32;
            _318_v32 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("thg"), _dafny.Seq.update(_dafny.Seq.UnicodeFromString("suocs"), _module.__default.safeIndex(new BigNumber((_317_v31).cardinality()), new BigNumber((_dafny.Seq.UnicodeFromString("suocs")).length)), new _dafny.CodePoint('p'.codePointAt(0))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(247)), ((_319_v0) => function (_320_i7) {
              return (_module.D4.create_DC15((_this).f23, new _dafny.CodePoint('f'.codePointAt(0)), new BigNumber((_319_v0).length))).dtor_cf27;
            })(_277_v0)), _277_v0, _277_v0);
            let _321_v33;
            _321_v33 = _dafny.Seq.of(_310_v26);
            let _322_v34;
            _322_v34 = _dafny.Map.Empty.slice().updateUnsafe((_318_v32)[_module.__default.safeIndex(_module.__default.fm0(new BigNumber(348), !((_this).f28), globalState), new BigNumber((_318_v32).length))],_321_v33);
            _322_v34 = (_322_v34).update(_277_v0, _321_v33);
            let _323_v35;
            _323_v35 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,_314_v30);
            let _324_v36;
            let _nw55 = Array((new BigNumber(13)).toNumber());
            _nw55[(_dafny.ZERO).toNumber()] = (((_this).f32) ? (_314_v30) : (_314_v30));
            _nw55[(_dafny.ONE).toNumber()] = _314_v30;
            _nw55[(new BigNumber(2)).toNumber()] = _314_v30;
            _nw55[(new BigNumber(3)).toNumber()] = _314_v30;
            _nw55[(new BigNumber(4)).toNumber()] = _314_v30;
            _nw55[(new BigNumber(5)).toNumber()] = (((_323_v35).contains((_this).f23)) ? ((_323_v35).get((_this).f23)) : (_314_v30));
            _nw55[(new BigNumber(6)).toNumber()] = _314_v30;
            _nw55[(new BigNumber(7)).toNumber()] = _314_v30;
            _nw55[(new BigNumber(8)).toNumber()] = _314_v30;
            _nw55[(new BigNumber(9)).toNumber()] = _314_v30;
            _nw55[(new BigNumber(10)).toNumber()] = _314_v30;
            _nw55[(new BigNumber(11)).toNumber()] = _314_v30;
            _nw55[(new BigNumber(12)).toNumber()] = _314_v30;
            _324_v36 = _nw55;
            let _index62 = _module.__default.safeIndex(new BigNumber(502), new BigNumber((_324_v36).length));
            (_324_v36)[_index62] = _314_v30;
            let _index63 = _module.__default.safeIndex(new BigNumber(502), new BigNumber((_324_v36).length));
            (_324_v36)[_index63] = _314_v30;
            let _325_v37;
            let _nw56 = new _module.C0();
            _nw56.__ctor();
            _325_v37 = _nw56;
          }
        }
      }
      let _326_v38;
      _326_v38 = _dafny.Seq.of((_this).f32);
      let _327_v39;
      _327_v39 = _dafny.MultiSet.fromElements((_this).f23);
      _326_v38 = _dafny.Seq.Concat(_dafny.Seq.update(_326_v38, _module.__default.safeIndex(_module.__default.fm0(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_327_v39).cardinality()),_dafny.MultiSet.FromArray(_326_v38))).length), !(p1), globalState), new BigNumber((_326_v38).length)), _module.__default.fm1((_this).f23, (_this).f32, globalState)), _326_v38);
      r0 = (_this).f23;
      let _328_v40;
      _328_v40 = _dafny.Set.fromElements((_this).f23, new BigNumber((_this.f29).length));
      r1 = (_328_v40).Union(_328_v40);
      return [r0, r1];
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Seq.of();
      let _329_v0;
      let _nw57 = Array((new BigNumber(29)).toNumber()).fill(false);
      _329_v0 = _nw57;
      let _330_v1;
      _330_v1 = _dafny.Seq.of(p0, p0, _dafny.Seq.Concat(p0, p0));
      let _331_v2;
      _331_v2 = _module.D0.create_DC0(p0);
      let _332_v3;
      _332_v3 = _module.D1.create_DC5((_this).f23, (_this).f32, (_this).f23, (_this).f23);
      let _pat_let_tv12 = p0;
      let _pat_let_tv13 = p0;
      let _pat_let_tv14 = p0;
      let _pat_let_tv15 = p0;
      let _pat_let_tv16 = p0;
      let _pat_let_tv17 = p0;
      let _pat_let_tv18 = p0;
      let _pat_let_tv19 = p0;
      let _pat_let_tv20 = _330_v1;
      let _pat_let_tv21 = p0;
      let _pat_let_tv22 = p0;
      let _pat_let_tv23 = p0;
      let _pat_let_tv24 = p0;
      let _pat_let_tv25 = _330_v1;
      let _rhs83 = _dafny.areEqual(_dafny.Seq.UnicodeFromString("wcapto"), p0);
      let _rhs84 = _329_v0;
      let _rhs85 = function (_source1) {
        if (_source1.is_DC1) {
          let _333___mcc_h0 = (_source1).cf1;
          let _334___mcc_h1 = (_source1).cf2;
          let _335___mcc_h2 = (_source1).cf3;
          let _336_cf3 = _335___mcc_h2;
          let _337_cf2 = _334___mcc_h1;
          let _338_cf1 = _333___mcc_h0;
          if (_337_cf2) {
            return _337_cf2;
          } else {
            return (_this).f32;
          }
        } else if (_source1.is_DC0) {
          let _339___mcc_h3 = (_source1).cf0;
          let _340_cf0 = _339___mcc_h3;
          return (_this).f28;
        } else {
          let _341___mcc_h4 = (_source1).cf4;
          let _342_cf4 = _341___mcc_h4;
          return !_dafny.areEqual(_dafny.Seq.update(_pat_let_tv12, _module.__default.safeIndex((_this).f23, new BigNumber((_pat_let_tv13).length)), (_module.D2.create_DC8((_this).f23, (_this).f23, new _dafny.CodePoint('o'.codePointAt(0)), new BigNumber((_pat_let_tv14).length), (_this).f28)).dtor_cf17), _dafny.Seq.update(_dafny.Seq.update(_pat_let_tv15, _module.__default.safeIndex((_this).f23, new BigNumber((_pat_let_tv16).length)), new _dafny.CodePoint('i'.codePointAt(0))), _module.__default.safeIndex((_this).f23, new BigNumber((_dafny.Seq.update(_pat_let_tv17, _module.__default.safeIndex((_this).f23, new BigNumber((_pat_let_tv18).length)), new _dafny.CodePoint('i'.codePointAt(0)))).length)), new _dafny.CodePoint('b'.codePointAt(0))));
        }
      }(_331_v2);
      let _rhs86 = function (_source2) {
        if (_source2.is_DC4) {
          let _343___mcc_h5 = (_source2).cf6;
          let _344___mcc_h6 = (_source2).cf7;
          let _345___mcc_h7 = (_source2).cf8;
          let _346_cf8 = _345___mcc_h7;
          let _347_cf7 = _344___mcc_h6;
          let _348_cf6 = _343___mcc_h5;
          return _dafny.Seq.of(_pat_let_tv19);
        } else if (_source2.is_DC5) {
          let _349___mcc_h8 = (_source2).cf9;
          let _350___mcc_h9 = (_source2).cf10;
          let _351___mcc_h10 = (_source2).cf11;
          let _352___mcc_h11 = (_source2).cf12;
          let _353_cf12 = _352___mcc_h11;
          let _354_cf11 = _351___mcc_h10;
          let _355_cf10 = _350___mcc_h9;
          let _356_cf9 = _349___mcc_h8;
          return _dafny.Seq.Concat((_pat_let_tv20), _dafny.Seq.of(_pat_let_tv21));
        } else if (_source2.is_DC6) {
          let _357___mcc_h12 = (_source2).cf13;
          let _358_cf13 = _357___mcc_h12;
          return _dafny.Seq.of(_pat_let_tv22, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-44)), function (_359_i0) {
            return new _dafny.CodePoint('f'.codePointAt(0));
          }), _pat_let_tv23, _pat_let_tv24);
        } else {
          let _360___mcc_h13 = (_source2).cf5;
          let _361_cf5 = _360___mcc_h13;
          return _pat_let_tv25;
        }
      }(_332_v3);
      let _lhs82 = globalState;
      let _lhs83 = globalState;
      _lhs82.f7 = _rhs83;
      _329_v0 = _rhs84;
      _lhs83.f9 = _rhs85;
      _330_v1 = _rhs86;
      let _362_v4;
      _362_v4 = _module.D1.create_DC6(_329_v0);
      let _363_v5;
      _363_v5 = _dafny.MultiSet.fromElements((_362_v4).dtor_cf13);
      _363_v5 = (_363_v5).Union(_363_v5);
      (globalState).f21 = !((_this).f32) || ((_this).f28);
      let _hi7 = (_this).f23;
      for (let _364_i1 = ((_this).f23).minus((_this).f23); _364_i1.isLessThan(_hi7); _364_i1 = _364_i1.plus(_dafny.ONE)) {
        let _365_v6;
        _365_v6 = _dafny.MultiSet.fromElements((_this).f23);
        if (((_365_v6).Difference(_365_v6)).contains(_364_i1)) {
          (globalState).f7 = true;
          let _366_v7;
          _366_v7 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f23),(_this).f32);
          _366_v7 = ((!(false)) ? (_366_v7) : (_366_v7));
          let _367_v8;
          _367_v8 = _dafny.Map.Empty.slice().updateUnsafe(_364_i1,(_dafny.ZERO).minus(((_this).f23).multipliedBy((_this).f23)));
          _367_v8 = (_367_v8).update(_364_i1, new BigNumber(((_this).f31).length));
          (globalState).f8 = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(p0, _dafny.Seq.Concat(p0, p0)), _module.__default.safeIndex((_364_i1).minus(new BigNumber(((_363_v5).update(_329_v0, _module.__default.abs((_this).f23))).cardinality())), new BigNumber((_dafny.Seq.Concat(p0, _dafny.Seq.Concat(p0, p0))).length)), new _dafny.CodePoint('b'.codePointAt(0)))).length);
          let _368_v9;
          _368_v9 = _dafny.MultiSet.fromElements((_this).f32, (_this).f28);
          let _369_v10;
          _369_v10 = _dafny.Seq.of((((_366_v7).contains((_this).f23)) ? ((_366_v7).get((_this).f23)) : ((_this).f28)), (_this).f32);
          let _370_v11;
          _370_v11 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f28),(_this).f28);
          let _371_v12;
          _371_v12 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_this).f28), _370_v11);
          let _372_v13;
          let _nw58 = Array((new BigNumber(24)).toNumber());
          _nw58[(_dafny.ZERO).toNumber()] = _module.__default.fm0(_364_i1, (_this).f32, globalState);
          _nw58[(_dafny.ONE).toNumber()] = (_this).f23;
          _nw58[(new BigNumber(2)).toNumber()] = _364_i1;
          _nw58[(new BigNumber(3)).toNumber()] = _364_i1;
          _nw58[(new BigNumber(4)).toNumber()] = _364_i1;
          _nw58[(new BigNumber(5)).toNumber()] = (_this).f23;
          _nw58[(new BigNumber(6)).toNumber()] = new BigNumber((_369_v10).length);
          _nw58[(new BigNumber(7)).toNumber()] = new BigNumber(362);
          _nw58[(new BigNumber(8)).toNumber()] = _364_i1;
          _nw58[(new BigNumber(9)).toNumber()] = new BigNumber(((_371_v12)[_module.__default.safeIndex(_364_i1, new BigNumber((_371_v12).length))]).length);
          _nw58[(new BigNumber(10)).toNumber()] = new BigNumber((p0).length);
          _nw58[(new BigNumber(11)).toNumber()] = (_this).f23;
          _nw58[(new BigNumber(12)).toNumber()] = new BigNumber(((_368_v9).update(true, _module.__default.abs(_364_i1))).cardinality());
          _nw58[(new BigNumber(13)).toNumber()] = _module.__default.fm0((_this).f23, (_this).f28, globalState);
          _nw58[(new BigNumber(14)).toNumber()] = (_this).f23;
          _nw58[(new BigNumber(15)).toNumber()] = new BigNumber(384);
          _nw58[(new BigNumber(16)).toNumber()] = new BigNumber((p0).length);
          _nw58[(new BigNumber(17)).toNumber()] = _364_i1;
          _nw58[(new BigNumber(18)).toNumber()] = new BigNumber(501);
          _nw58[(new BigNumber(19)).toNumber()] = new BigNumber((_368_v9).cardinality());
          _nw58[(new BigNumber(20)).toNumber()] = (_this).f23;
          _nw58[(new BigNumber(21)).toNumber()] = new BigNumber((_369_v10).length);
          _nw58[(new BigNumber(22)).toNumber()] = _364_i1;
          _nw58[(new BigNumber(23)).toNumber()] = _364_i1;
          _372_v13 = _nw58;
          let _373_v14;
          let _nw59 = new _module.C1();
          _nw59.__ctor((_this).f32, _this.f29, (((_368_v9).contains(!((_this).f28))) ? ((_368_v9).get(!((_this).f28))) : ((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f32,_372_v13)).length)))));
          _373_v14 = _nw59;
        } else {
          let _374_v15;
          let _init4 = function (_375_i2) {
            return _module.__default.safeModuloInt(_375_i2, (_dafny.ZERO).minus((_this).f23));
          };
          let _nw60 = Array((new BigNumber(27)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw60.length); _i0_4++) {
            _nw60[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _374_v15 = _nw60;
          let _index64 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_374_v15).length));
          (_374_v15)[_index64] = _364_i1;
          let _index65 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_374_v15).length));
          (_374_v15)[_index65] = (_this).f23;
          (globalState).f21 = ((_this).f32) || (!((_this).f32));
          (globalState).f7 = !_dafny.areEqual(_dafny.Seq.Concat(p0, p0), _module.__default.fm2(globalState));
          let _376_v16;
          _376_v16 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f32) === ((_this).f28),_374_v15);
          _376_v16 = (_376_v16).Merge(_376_v16);
          let _377_v17;
          _377_v17 = _dafny.Map.Empty.slice().updateUnsafe(_364_i1,(_dafny.MultiSet.fromElements(true, (_this).f32, (_this).f32)).IsDisjointFrom(_dafny.MultiSet.fromElements(true)));
          (globalState).f9 = !((((_377_v17).contains(_364_i1)) ? ((_377_v17).get(_364_i1)) : ((_this).f32)));
        }
        let _378_v18;
        _378_v18 = _dafny.Map.Empty.slice().updateUnsafe((_this).f32,_dafny.Seq.Create(_module.__default.abs(new BigNumber(-469)), function (_379_i3) {
          return new _dafny.CodePoint('d'.codePointAt(0));
        }));
        let _380_v19;
        let _nw61 = new _module.C1();
        _nw61.__ctor((_this).f32, _this.f29, new BigNumber(((((_378_v18).contains((_this).f32)) ? ((_378_v18).get((_this).f32)) : (p0))).length));
        _380_v19 = _nw61;
        let _381_v20;
        _381_v20 = _dafny.Seq.of((_this).f28, true, (_this).f32, (_this).f32, _module.__default.fm1(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(131)), function (_382_i4) {
          return new _dafny.CodePoint('t'.codePointAt(0));
        })).length), (_this).f32, globalState));
        let _383_v21;
        _383_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,p0);
        let _384_v22;
        _384_v22 = _dafny.Map.Empty.slice().updateUnsafe(_381_v20,_dafny.Seq.Concat(p0, (((_383_v21).contains(_364_i1)) ? ((_383_v21).get(_364_i1)) : (_dafny.Seq.UnicodeFromString("hkctqqp")))));
        _384_v22 = (_384_v22).update(_381_v20, p0);
        let _385_v23;
        let _nw62 = Array((new BigNumber(15)).toNumber()).fill(_module.D1.Default());
        _385_v23 = _nw62;
        let _index66 = _module.__default.safeIndex(new BigNumber(72), new BigNumber((_385_v23).length));
        (_385_v23)[_index66] = _332_v3;
        let _index67 = _module.__default.safeIndex(new BigNumber(72), new BigNumber((_385_v23).length));
        (_385_v23)[_index67] = ((!((_this).f32) || ((_this).f32)) ? (_module.__default.fm23(globalState)) : (_332_v3));
      }
      (globalState).f14 = (_this).f23;
      let _386_v24;
      let _nw63 = Array((new BigNumber(11)).toNumber()).fill(_module.D1.Default());
      _386_v24 = _nw63;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_386_v24).length))) {
        let _387_i5 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_387_i5)) && ((_387_i5).isLessThan(new BigNumber((_386_v24).length))))) {
          (_386_v24)[(_387_i5)] = _module.D1.create_DC4(_dafny.Map.Empty.slice().updateUnsafe((_this).f23,(_this).f23), (_this).f28, !((_this).f32));
        }
      }
      r0 = (_this).f23;
      r1 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-922)), ((_388_p0) => function (_389_i6) {
        return new BigNumber((_388_p0).length);
      })(p0)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(454)), function (_390_i7) {
        return (_this).f23;
      })), _module.__default.safeIndex((_this).f23, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-922)), ((_391_p0) => function (_392_i6) {
        return new BigNumber((_391_p0).length);
      })(p0)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(454)), function (_393_i7) {
        return (_this).f23;
      }))).length)), (_this).f23);
      return [r0, r1];
    }
    m6(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Map.Empty;
      let r2 = _module.D1.Default();
      let r3 = [];
      let _394_v0;
      _394_v0 = _dafny.MultiSet.fromElements((_this).f28, (_this).f28, p0, true);
      let _395_v1;
      _395_v1 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_394_v0).cardinality()));
      let _396_v2;
      _396_v2 = _dafny.Seq.UnicodeFromString("fnnsolgl");
      let _397_v3;
      _397_v3 = _module.D2.create_DC8((_this).fm18(_395_v1, _396_v2, (_this).f28, globalState), (_this).f23, new _dafny.CodePoint('a'.codePointAt(0)), new BigNumber(812), false);
      let _398_v4;
      _398_v4 = _module.D2.create_DC8((_this).f23, (_this).f23, new _dafny.CodePoint('t'.codePointAt(0)), (_397_v3).dtor_cf16, (_397_v3).dtor_cf19);
      let _399_v5;
      _399_v5 = _dafny.Set.fromElements(_398_v4, _397_v3);
      _399_v5 = _399_v5;
      let _400_v6;
      let _nw64 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
      _400_v6 = _nw64;
      let _index68 = _module.__default.safeIndex(new BigNumber(669), new BigNumber((_400_v6).length));
      (_400_v6)[_index68] = new BigNumber(528);
      let _index69 = _module.__default.safeIndex(new BigNumber(669), new BigNumber((_400_v6).length));
      (_400_v6)[_index69] = (_this).f23;
      let _hi8 = ((_400_v6)[_module.__default.safeIndex(new BigNumber(669), new BigNumber((_400_v6).length))]).minus((_400_v6)[_module.__default.safeIndex(new BigNumber(669), new BigNumber((_400_v6).length))]);
      for (let _401_i0 = (new BigNumber((_this.f29).length)).minus((_this).f23); _401_i0.isLessThan(_hi8); _401_i0 = _401_i0.plus(_dafny.ONE)) {
        (globalState).f8 = (_400_v6)[_module.__default.safeIndex(new BigNumber(669), new BigNumber((_400_v6).length))];
        let _402_v7;
        _402_v7 = new _dafny.CodePoint('w'.codePointAt(0));
        let _403_v8;
        _403_v8 = _dafny.Map.Empty.slice().updateUnsafe(p0,_402_v7);
        let _index70 = _module.__default.safeIndex(new BigNumber(669), new BigNumber((_400_v6).length));
        (_400_v6)[_index70] = ((!((_403_v8).update((_this).f32, _402_v7)).equals(_403_v8)) ? ((((_395_v1).contains(p0)) ? ((_395_v1).get(p0)) : ((_this).f23))) : (((_this).f23).multipliedBy(new BigNumber(-826))));
        let _404_v9;
        let _nw65 = new _module.C1();
        _nw65.__ctor((_this).f28, _dafny.Seq.Concat(_dafny.Seq.of(_401_i0), _this.f29), _401_i0);
        _404_v9 = _nw65;
        (_this).f29 = _this.f29;
      }
      let _405_v10;
      _405_v10 = _dafny.Seq.of(p0, (_this).f28);
      let _406_v11;
      _406_v11 = _module.D3.create_DC12(_405_v10);
      _406_v11 = (((p0) === (p0)) ? (_406_v11) : (_406_v11));
      let _407_v12;
      let _nw66 = Array((new BigNumber(25)).toNumber()).fill(false);
      _407_v12 = _nw66;
      let _index71 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_407_v12).length));
      (_407_v12)[_index71] = p0;
      let _index72 = _module.__default.safeIndex(new BigNumber(779), new BigNumber((_407_v12).length));
      (_407_v12)[_index72] = (_this).f32;
      let _408_i1;
      _408_i1 = _dafny.ZERO;
      L6: {
        while ((_407_v12)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_407_v12).length))]) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_408_i1)) {
              break L6;
            }
            _408_i1 = (_408_i1).plus(_dafny.ONE);
            r0 = !((_407_v12)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_407_v12).length))]);
            let _409_v13;
            _409_v13 = _dafny.Seq.of(_396_v2, _396_v2);
            (globalState).f21 = !_dafny.Seq.contains(_409_v13, _396_v2);
            let _410_v14;
            _410_v14 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(((_395_v1)).length)));
            (globalState).f14 = (_this).fm18((_410_v14)[_module.__default.safeIndex((_this).f23, new BigNumber((_410_v14).length))], _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(339)), function (_411_i2) {
              return new _dafny.CodePoint('h'.codePointAt(0));
            }), _396_v2), (_407_v12)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_407_v12).length))], globalState);
            _395_v1 = (_395_v1).update((_this).f32, (_this).f23);
          }
        }
      }
      r0 = (_this).f28;
      let _412_v15;
      _412_v15 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f23).multipliedBy((_this).f23),_407_v12);
      r1 = _412_v15;
      let _413_v16;
      _413_v16 = _module.D1.create_DC4((_this).f31, (_407_v12)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_407_v12).length))], (_407_v12)[_module.__default.safeIndex(new BigNumber(779), new BigNumber((_407_v12).length))]);
      r2 = _413_v16;
      r3 = _407_v12;
      return [r0, r1, r2, r3];
    }
    get f31() {
      let _this = this;
      return _this._f31;
    };
    get f32() {
      let _this = this;
      return _this._f32;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f29 = _dafny.Seq.of();
      this._f23 = _dafny.ZERO;
      this._f28 = false;
      this._f30 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f29() {
      let _this = this;
      return _this._f29;
    };
    set f29(value) {
      let _this = this;
      _this._f29 = value;
    };
    get f23() {
      let _this = this;
      return _this._f23;
    };
    get f28() {
      let _this = this;
      return _this._f28;
    };
    __ctor(f30, f23, f28, f29) {
      let _this = this;
      (_this)._f30 = f30;
      (_this)._f23 = f23;
      (_this)._f28 = f28;
      (_this)._f29 = f29;
      return;
    }
    fm12(p0, p1, p2, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(((_this).f23).plus(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-585)), function (_414_i0) {
        return new _dafny.CodePoint('r'.codePointAt(0));
      })).length), new BigNumber(974))));
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Set.Empty;
      let _415_v0;
      let _nw67 = Array((new BigNumber(22)).toNumber()).fill(false);
      _415_v0 = _nw67;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_415_v0).length))) {
        let _416_i0 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_416_i0)) && ((_416_i0).isLessThan(new BigNumber((_415_v0).length))))) {
          (_415_v0)[(_416_i0)] = (_module.D2.create_DC8(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f23,(_this).f23)))).cardinality()), (_this).f23, new _dafny.CodePoint('p'.codePointAt(0)), new BigNumber((_dafny.Seq.UnicodeFromString("nasts")).length), true)).dtor_cf19;
        }
      }
      if ((_module.__default.safeDivisionInt((_this).f23, (_this).f23)).isLessThan(((_this).f23).multipliedBy((_this).f23))) {
        let _417_v1;
        _417_v1 = _dafny.Set.fromElements(p1, p1);
        let _rhs87 = (_417_v1).IsProperSubsetOf((_417_v1).Difference(_dafny.Set.fromElements((_this).f28)));
        let _rhs88 = (_this).f28;
        let _lhs84 = globalState;
        let _lhs85 = globalState;
        _lhs84.f21 = _rhs87;
        _lhs85.f21 = _rhs88;
        let _418_v2;
        _418_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,false);
        let _index73 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((p0).length));
        (p0)[_index73] = (((_418_v2).contains(p1)) ? ((_418_v2).get(p1)) : (p1));
        let _index74 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((p0).length));
        (p0)[_index74] = (_this).f28;
        let _419_v3;
        let _nw68 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Map.Empty);
        _419_v3 = _nw68;
        let _420_v4;
        _420_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,true);
        let _index75 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_419_v3).length));
        (_419_v3)[_index75] = _420_v4;
        let _index76 = _module.__default.safeIndex(new BigNumber(243), new BigNumber((_419_v3).length));
        (_419_v3)[_index76] = _420_v4;
        let _index77 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((p0).length));
        (p0)[_index77] = !(_418_v2).contains((p0)[_module.__default.safeIndex(new BigNumber(868), new BigNumber((p0).length))]);
        let _421_v5;
        _421_v5 = _module.D1.create_DC5((_this).f23, (p0)[_module.__default.safeIndex(new BigNumber(868), new BigNumber((p0).length))], (_this).f23, (_this).f23);
        let _422_v6;
        _422_v6 = _dafny.Seq.of(_421_v5);
        let _423_v7;
        _423_v7 = _dafny.Map.Empty.slice().updateUnsafe(_422_v6,p1);
        let _424_v8;
        _424_v8 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-225),(_this).f23));
        let _425_v9;
        _425_v9 = _dafny.MultiSet.fromElements((_this).f28);
        let _426_v10;
        _426_v10 = _module.D3.create_DC11(_module.__default.fm0(new BigNumber((_425_v9).cardinality()), (_this).f28, globalState));
        let _427_v11;
        _427_v11 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_425_v9).cardinality()),(_426_v10).dtor_cf22);
        let _428_v12;
        _428_v12 = _dafny.Set.fromElements((_this).f23);
        let _429_v13;
        _429_v13 = _module.D1.create_DC4(_427_v11, ((_this).f28) || (p1), !(new BigNumber((_428_v12).length)).isEqualTo((_this).f23));
        let _rhs89 = _423_v7;
        let _rhs90 = (_this).f23;
        let _rhs91 = _module.__default.fm13(new _dafny.CodePoint('q'.codePointAt(0)), _module.__default.fm1((_this).f23, (_this).f28, globalState), ((_module.__default.fm1((_this).f23, p1, globalState)) ? ((_this).f23) : ((_this).f23)), ((p1) ? ((p0)[_module.__default.safeIndex(new BigNumber(868), new BigNumber((p0).length))]) : (_module.__default.fm1((_this).f23, (p0)[_module.__default.safeIndex(new BigNumber(868), new BigNumber((p0).length))], globalState))), globalState);
        let _rhs92 = _429_v13;
        let _rhs93 = new BigNumber(473);
        let _lhs86 = globalState;
        let _lhs87 = globalState;
        _423_v7 = _rhs89;
        _lhs86.f14 = _rhs90;
        _424_v8 = _rhs91;
        _429_v13 = _rhs92;
        _lhs87.f14 = _rhs93;
      } else {
        (globalState).f14 = (_this).f23;
        let _430_v14;
        _430_v14 = _dafny.Set.fromElements(p1);
        let _431_v15;
        _431_v15 = new _dafny.CodePoint('h'.codePointAt(0));
        let _432_v16;
        _432_v16 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f23);
        let _433_v17;
        _433_v17 = _dafny.MultiSet.fromElements((_this).f23, (_this).fm12((_this).f23, (_this).f28, p1, globalState), (_this).f23);
        let _434_v18;
        _434_v18 = _dafny.Seq.UnicodeFromString("htcv");
        let _435_v19;
        let _nw69 = Array((new BigNumber(28)).toNumber());
        _nw69[(_dafny.ZERO).toNumber()] = (_this).f23;
        _nw69[(_dafny.ONE).toNumber()] = _module.__default.safeDivisionInt((_this).f23, (_this).f23);
        _nw69[(new BigNumber(2)).toNumber()] = (_this).f23;
        _nw69[(new BigNumber(3)).toNumber()] = (_this).f23;
        _nw69[(new BigNumber(4)).toNumber()] = (_this).f23;
        _nw69[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("n")).length);
        _nw69[(new BigNumber(6)).toNumber()] = (new BigNumber((_430_v14).length)).multipliedBy((_this).f23);
        _nw69[(new BigNumber(7)).toNumber()] = (_this).f23;
        _nw69[(new BigNumber(8)).toNumber()] = new BigNumber((_module.__default.fm14((_this).f23, _431_v15, (_this).f28, globalState)).cardinality());
        _nw69[(new BigNumber(9)).toNumber()] = (_this).f23;
        _nw69[(new BigNumber(10)).toNumber()] = (_this).f23;
        _nw69[(new BigNumber(11)).toNumber()] = (_this).f23;
        _nw69[(new BigNumber(12)).toNumber()] = ((_this).fm12((_this).f23, _module.__default.fm1((_this).f23, p1, globalState), !(p1), globalState)).plus(new BigNumber(447));
        _nw69[(new BigNumber(13)).toNumber()] = (_this).f23;
        _nw69[(new BigNumber(14)).toNumber()] = new BigNumber(629);
        _nw69[(new BigNumber(15)).toNumber()] = (_this).fm12(_module.__default.fm0((_this).f23, p1, globalState), (_this).f28, (_this).f28, globalState);
        _nw69[(new BigNumber(16)).toNumber()] = (_this).f23;
        _nw69[(new BigNumber(17)).toNumber()] = (new BigNumber((_432_v16).length)).multipliedBy((((_433_v17).contains(new BigNumber(876))) ? ((_433_v17).get(new BigNumber(876))) : ((_this).f23)));
        _nw69[(new BigNumber(18)).toNumber()] = ((_this).f23).plus((_this).fm12(new BigNumber((_434_v18).length), (_this).f28, p1, globalState));
        _nw69[(new BigNumber(19)).toNumber()] = (new BigNumber(547)).plus((_this).f23);
        _nw69[(new BigNumber(20)).toNumber()] = new BigNumber((_434_v18).length);
        _nw69[(new BigNumber(21)).toNumber()] = ((p1) ? (new BigNumber((_this.f29).length)) : ((_this).f23));
        _nw69[(new BigNumber(22)).toNumber()] = (_this).f23;
        _nw69[(new BigNumber(23)).toNumber()] = (_this).f23;
        _nw69[(new BigNumber(24)).toNumber()] = _module.__default.safeModuloInt((_this).f23, (_this).f23);
        _nw69[(new BigNumber(25)).toNumber()] = new BigNumber(((_433_v17).update((_this).f23, _module.__default.abs((_this).fm12((_this).f23, (_this).f28, (_this).f28, globalState)))).cardinality());
        _nw69[(new BigNumber(26)).toNumber()] = new BigNumber(966);
        _nw69[(new BigNumber(27)).toNumber()] = (_this).f23;
        _435_v19 = _nw69;
        let _index78 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_435_v19).length));
        (_435_v19)[_index78] = (_this).f23;
        let _436_v20;
        _436_v20 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,((_this).f23).isLessThanOrEqualTo(new BigNumber(141)));
        let _index79 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_435_v19).length));
        (_435_v19)[_index79] = (_dafny.ZERO).minus(new BigNumber((_436_v20).length));
        (globalState).f8 = _module.__default.fm0((_this).f23, p1, globalState);
        if ((((_this).f28) ? ((_this).f28) : (p1))) {
          let _437_v21;
          _437_v21 = _dafny.Seq.of(false, p1, (_this).f28);
          let _438_v22;
          _438_v22 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_437_v21, _437_v21));
          let _439_v23;
          _439_v23 = _dafny.MultiSet.fromElements(p1);
          let _440_v24;
          _440_v24 = _dafny.Map.Empty.slice().updateUnsafe((_435_v19)[_module.__default.safeIndex(new BigNumber(726), new BigNumber((_435_v19).length))],_439_v23);
          let _pat_let_tv26 = p1;
          _438_v22 = _module.__default.fm15((function (_pat_let0_0) {
            return function (_441_dt__update__tmp_h0) {
              return function (_pat_let1_0) {
                return function (_442_dt__update_hcf21_h0) {
                  return _module.D3.create_DC10(_442_dt__update_hcf21_h0);
                }(_pat_let1_0);
              }(_pat_let_tv26);
            }(_pat_let0_0);
          }(_module.D3.create_DC10(p1))).dtor_cf21, _440_v24, globalState);
          let _443_v25;
          _443_v25 = _dafny.Seq.of(_415_v0, _415_v0);
          let _444_v26;
          let _nw70 = Array((new BigNumber(14)).toNumber());
          _nw70[(_dafny.ZERO).toNumber()] = _415_v0;
          _nw70[(_dafny.ONE).toNumber()] = p0;
          _nw70[(new BigNumber(2)).toNumber()] = p0;
          _nw70[(new BigNumber(3)).toNumber()] = _415_v0;
          _nw70[(new BigNumber(4)).toNumber()] = (((_this).f28) ? (p0) : (_415_v0));
          _nw70[(new BigNumber(5)).toNumber()] = _415_v0;
          _nw70[(new BigNumber(6)).toNumber()] = (_443_v25)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.FromArray(_this.f29)).cardinality()), new BigNumber((_443_v25).length))];
          _nw70[(new BigNumber(7)).toNumber()] = _415_v0;
          _nw70[(new BigNumber(8)).toNumber()] = ((p1) ? (p0) : (_415_v0));
          _nw70[(new BigNumber(9)).toNumber()] = p0;
          _nw70[(new BigNumber(10)).toNumber()] = p0;
          _nw70[(new BigNumber(11)).toNumber()] = _415_v0;
          _nw70[(new BigNumber(12)).toNumber()] = p0;
          _nw70[(new BigNumber(13)).toNumber()] = p0;
          _444_v26 = _nw70;
          let _index80 = _module.__default.safeIndex(new BigNumber(281), new BigNumber((_444_v26).length));
          (_444_v26)[_index80] = p0;
          let _index81 = _module.__default.safeIndex(new BigNumber(281), new BigNumber((_444_v26).length));
          (_444_v26)[_index81] = p0;
          _430_v14 = _430_v14;
          let _445_v27;
          let _nw71 = Array((new BigNumber(15)).toNumber()).fill([]);
          _445_v27 = _nw71;
          let _index82 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_445_v27).length));
          (_445_v27)[_index82] = _435_v19;
          let _446_v28;
          _446_v28 = _dafny.Map.Empty.slice().updateUnsafe((_435_v19)[_module.__default.safeIndex(new BigNumber(726), new BigNumber((_435_v19).length))],p1);
          let _447_v29;
          _447_v29 = _dafny.Map.Empty.slice().updateUnsafe(_434_v18,((p1) ? (_dafny.Map.Empty.slice().updateUnsafe((_437_v21)[_module.__default.safeIndex((_435_v19)[_module.__default.safeIndex(new BigNumber(726), new BigNumber((_435_v19).length))], new BigNumber((_437_v21).length))],new BigNumber((_439_v23).cardinality()))) : (_dafny.Map.Empty.slice().updateUnsafe((((_446_v28).contains(new BigNumber((_434_v18).length))) ? ((_446_v28).get(new BigNumber((_434_v18).length))) : ((_this).f28)),(_435_v19)[_module.__default.safeIndex(new BigNumber(726), new BigNumber((_435_v19).length))]))));
          let _index83 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_445_v27).length));
          let _index84 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_435_v19).length));
          let _rhs94 = (((_447_v29).contains(_434_v18)) ? ((_447_v29).get(_434_v18)) : (_432_v16));
          let _rhs95 = _435_v19;
          let _rhs96 = (_this).fm12((_435_v19)[_module.__default.safeIndex(new BigNumber(726), new BigNumber((_435_v19).length))], false, p1, globalState);
          let _rhs97 = (_this).f28;
          let _rhs98 = ((_435_v19)[_module.__default.safeIndex(new BigNumber(726), new BigNumber((_435_v19).length))]).minus((_this).f23);
          let _lhs88 = _445_v27;
          let _lhs89 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_445_v27).length));
          let _lhs90 = _435_v19;
          let _lhs91 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_435_v19).length));
          let _lhs92 = globalState;
          let _lhs93 = globalState;
          _432_v16 = _rhs94;
          _lhs88[_lhs89] = _rhs95;
          _lhs90[_lhs91] = _rhs96;
          _lhs92.f9 = _rhs97;
          _lhs93.f14 = _rhs98;
          (globalState).f8 = (_435_v19)[_module.__default.safeIndex(new BigNumber(726), new BigNumber((_435_v19).length))];
        } else {
          let _448_v30;
          _448_v30 = _dafny.Map.Empty.slice().updateUnsafe(_415_v0,((_this).f23).multipliedBy(new BigNumber(434)));
          _448_v30 = (_448_v30).update(_415_v0, (_this).f23);
          let _index85 = _module.__default.safeIndex(new BigNumber(726), new BigNumber((_435_v19).length));
          (_435_v19)[_index85] = (_435_v19)[_module.__default.safeIndex(new BigNumber(726), new BigNumber((_435_v19).length))];
          let _449_v31;
          _449_v31 = _dafny.Seq.of(_dafny.Seq.Concat(_this.f29, _this.f29), _this.f29);
          let _450_v32;
          _450_v32 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(391),_this.f29);
          _449_v31 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.update(_449_v31, _module.__default.safeIndex((((_433_v17).contains((_this).f23)) ? ((_433_v17).get((_this).f23)) : ((_this).f23)), new BigNumber((_449_v31).length)), _this.f29), _module.__default.safeIndex(new BigNumber((_dafny.Seq.of(p1)).length), new BigNumber((_dafny.Seq.update(_449_v31, _module.__default.safeIndex((((_433_v17).contains((_this).f23)) ? ((_433_v17).get((_this).f23)) : ((_this).f23)), new BigNumber((_449_v31).length)), _this.f29)).length)), (((_450_v32).contains((_435_v19)[_module.__default.safeIndex(new BigNumber(726), new BigNumber((_435_v19).length))])) ? ((_450_v32).get((_435_v19)[_module.__default.safeIndex(new BigNumber(726), new BigNumber((_435_v19).length))])) : (_dafny.Seq.of((_435_v19)[_module.__default.safeIndex(new BigNumber(726), new BigNumber((_435_v19).length))])))), _dafny.Seq.Concat(_449_v31, _449_v31));
          (_this).m5(globalState);
          let _451_v33;
          _451_v33 = _module.D4.create_DC14(_435_v19);
          _435_v19 = (_451_v33).dtor_cf25;
        }
        let _452_v34;
        _452_v34 = _dafny.Map.Empty.slice().updateUnsafe((_module.D0.create_DC1((_435_v19)[_module.__default.safeIndex(new BigNumber(726), new BigNumber((_435_v19).length))], (_this).f28, (_435_v19)[_module.__default.safeIndex(new BigNumber(726), new BigNumber((_435_v19).length))])).dtor_cf1,(_this).f23);
        let _453_v35;
        _453_v35 = _dafny.Seq.of(_module.__default.fm16((_this).f28, (_this).f23, (_this).f23, globalState));
        _452_v34 = (_452_v34).update((_this).f23, new BigNumber((((_453_v35)[_module.__default.safeIndex((_435_v19)[_module.__default.safeIndex(new BigNumber(726), new BigNumber((_435_v19).length))], new BigNumber((_453_v35).length))]).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_434_v18).length),(_this).f23))).length));
      }
      (globalState).f14 = (_this).f23;
      let _454_v36;
      _454_v36 = _dafny.Seq.UnicodeFromString("tjvi");
      let _455_v37;
      _455_v37 = new _dafny.CodePoint('y'.codePointAt(0));
      if (!(_dafny.areEqual(_dafny.Seq.UnicodeFromString("ii"), _dafny.Seq.update(_454_v36, _module.__default.safeIndex((_this).f23, new BigNumber((_454_v36).length)), _455_v37)))) {
        let _index86 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_415_v0).length));
        (_415_v0)[_index86] = (_this).f28;
        let _index87 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_415_v0).length));
        (_415_v0)[_index87] = !((_module.D3.create_DC10((_this).f28)).dtor_cf21);
        let _456_v38;
        let _init5 = ((_457_v0) => function (_458_i1) {
          return _dafny.MultiSet.fromElements((_457_v0)[_module.__default.safeIndex(new BigNumber(956), new BigNumber((_457_v0).length))]);
        })(_415_v0);
        let _nw72 = Array((new BigNumber(14)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw72.length); _i0_5++) {
          _nw72[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _456_v38 = _nw72;
        _456_v38 = _456_v38;
        let _459_v40;
        _459_v40 = _module.D1.create_DC4(function () {
  let _coll18 = new _dafny.Map();
  for (const _compr_18 of _dafny.IntegerRange(new BigNumber(57), new BigNumber(984))) {
    let _460_v39 = _compr_18;
    if (((new BigNumber(57)).isLessThanOrEqualTo(_460_v39)) && ((_460_v39).isLessThan(new BigNumber(984)))) {
      _coll18.push([(_460_v39).plus((_this).f23),(_this).f23]);
    }
  }
  return _coll18;
}(), (_415_v0)[_module.__default.safeIndex(new BigNumber(956), new BigNumber((_415_v0).length))], p1);
        let _index88 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_415_v0).length));
        (_415_v0)[_index88] = (_module.__default.fm1((_this).f23, false, globalState)) === ((_459_v40).dtor_cf7);
        (globalState).f6 = _dafny.Seq.Concat(_454_v36, _dafny.Seq.UnicodeFromString("l"));
        (globalState).f8 = (_this.f29)[_module.__default.safeIndex((_this).f23, new BigNumber((_this.f29).length))];
      } else {
        let _index89 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_415_v0).length));
        (_415_v0)[_index89] = p1;
        let _index90 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_415_v0).length));
        (_415_v0)[_index90] = p1;
        let _461_v41;
        let _init6 = function (_462_i2) {
          return (_462_i2).plus((_this).f23);
        };
        let _nw73 = Array((new BigNumber(11)).toNumber());
        for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw73.length); _i0_6++) {
          _nw73[_i0_6] = _init6(new BigNumber(_i0_6));
        }
        _461_v41 = _nw73;
        let _463_v42;
        _463_v42 = _dafny.Seq.of((_415_v0)[_module.__default.safeIndex(new BigNumber(296), new BigNumber((_415_v0).length))], p1);
        let _index91 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_461_v41).length));
        (_461_v41)[_index91] = new BigNumber((_463_v42).length);
        let _index92 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_461_v41).length));
        (_461_v41)[_index92] = (_this).fm12((_this).f23, (_415_v0)[_module.__default.safeIndex(new BigNumber(296), new BigNumber((_415_v0).length))], (_415_v0)[_module.__default.safeIndex(new BigNumber(296), new BigNumber((_415_v0).length))], globalState);
        let _464_v43;
        let _init7 = ((_465_v36) => function (_466_i3) {
          return _dafny.Seq.Concat(_465_v36, _465_v36);
        })(_454_v36);
        let _nw74 = Array((new BigNumber(18)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw74.length); _i0_7++) {
          _nw74[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _464_v43 = _nw74;
        let _index93 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_464_v43).length));
        (_464_v43)[_index93] = _454_v36;
        let _467_v44;
        _467_v44 = _dafny.MultiSet.fromElements(_455_v37);
        let _468_v45;
        _468_v45 = _dafny.Map.Empty.slice().updateUnsafe(true,(_461_v41)[_module.__default.safeIndex(new BigNumber(509), new BigNumber((_461_v41).length))]);
        let _469_v46;
        _469_v46 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("qkbs"), _dafny.Seq.UnicodeFromString("uxecrdjl"));
        let _470_v47;
        _470_v47 = _dafny.Set.fromElements(true, (_this).f28);
        let _471_v48;
        _471_v48 = _dafny.MultiSet.fromElements(_470_v47);
        let _472_v49;
        _472_v49 = _module.D4.create_DC15(new BigNumber((_471_v48).cardinality()), _455_v37, (_this).f23);
        let _index94 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_461_v41).length));
        let _index95 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_461_v41).length));
        let _index96 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_464_v43).length));
        let _index97 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_461_v41).length));
        let _rhs99 = (_461_v41)[_module.__default.safeIndex(new BigNumber(509), new BigNumber((_461_v41).length))];
        let _rhs100 = _module.__default.fm0((_this).f23, (_module.__default.fm17((_this).f28, new BigNumber((_454_v36).length), globalState)).IsProperSubsetOf((_467_v44).update(_455_v37, _module.__default.abs(new BigNumber((_468_v45).length)))), globalState);
        let _rhs101 = _dafny.Seq.Concat(_454_v36, _454_v36);
        let _rhs102 = ((_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(39)), ((_473_v36) => function (_474_i4) {
          return _473_v36;
        })(_454_v36)), _469_v46)) ? ((((_415_v0)[_module.__default.safeIndex(new BigNumber(296), new BigNumber((_415_v0).length))]) ? ((_472_v49).dtor_cf28) : ((_this).fm12((_this).f23, (_this).f28, _module.__default.fm1((_this).f23, p1, globalState), globalState)))) : (((_461_v41)[_module.__default.safeIndex(new BigNumber(509), new BigNumber((_461_v41).length))]).minus((_this).f23)));
        let _lhs94 = _461_v41;
        let _lhs95 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_461_v41).length));
        let _lhs96 = _461_v41;
        let _lhs97 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_461_v41).length));
        let _lhs98 = _464_v43;
        let _lhs99 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_464_v43).length));
        let _lhs100 = _461_v41;
        let _lhs101 = _module.__default.safeIndex(new BigNumber(509), new BigNumber((_461_v41).length));
        _lhs94[_lhs95] = _rhs99;
        _lhs96[_lhs97] = _rhs100;
        _lhs98[_lhs99] = _rhs101;
        _lhs100[_lhs101] = _rhs102;
        (_this).m5(globalState);
        let _index98 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_415_v0).length));
        (_415_v0)[_index98] = ((_461_v41)[_module.__default.safeIndex(new BigNumber(509), new BigNumber((_461_v41).length))]).isEqualTo(new BigNumber((_468_v45).length));
      }
      let _hi9 = new BigNumber(581);
      for (let _475_i5 = (_dafny.ZERO).minus((_this).f23); _475_i5.isLessThan(_hi9); _475_i5 = _475_i5.plus(_dafny.ONE)) {
        let _476_v50;
        _476_v50 = _dafny.Map.Empty.slice().updateUnsafe(p1,!(p1));
        let _477_v51;
        _477_v51 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_476_v50).length),p0);
        let _478_v52;
        let _init8 = function (_479_i6) {
          return (_479_i6).plus(new BigNumber((_dafny.Set.fromElements(_this.f29)).length));
        };
        let _nw75 = Array((new BigNumber(17)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw75.length); _i0_8++) {
          _nw75[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _478_v52 = _nw75;
        let _480_v53;
        _480_v53 = _dafny.Map.Empty.slice().updateUnsafe((((_477_v51).contains((_this).f23)) ? ((_477_v51).get((_this).f23)) : (p0)),_478_v52);
        let _481_v54;
        _481_v54 = _module.D4.create_DC14((((_480_v53).contains(p0)) ? ((_480_v53).get(p0)) : (_478_v52)));
        _481_v54 = _module.D4.create_DC14(_478_v52);
        let _482_v55;
        _482_v55 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,(_this).f23);
        let _483_v56;
        _483_v56 = _module.D1.create_DC4(_482_v55, p1, p1);
        let _source3 = _483_v56;
        if (_source3.is_DC4) {
          let _484___mcc_h0 = (_source3).cf6;
          let _485___mcc_h1 = (_source3).cf7;
          let _486___mcc_h2 = (_source3).cf8;
          let _487_cf8 = _486___mcc_h2;
          let _488_cf7 = _485___mcc_h1;
          let _489_cf6 = _484___mcc_h0;
          (globalState).f14 = (new BigNumber((_454_v36).length)).multipliedBy(((_this).f23).multipliedBy((_this).f23));
          let _490_v57;
          _490_v57 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_this).f23);
          let _491_v58;
          _491_v58 = _dafny.MultiSet.fromElements(new BigNumber((_490_v57).length));
          let _rhs103 = (_491_v58).IsSubsetOf(_491_v58);
          let _rhs104 = _481_v54;
          let _rhs105 = _455_v37;
          let _rhs106 = (_475_i5).plus((_this).f23);
          let _rhs107 = _481_v54;
          let _lhs102 = globalState;
          _488_cf7 = _rhs103;
          _481_v54 = _rhs104;
          _455_v37 = _rhs105;
          _lhs102.f14 = _rhs106;
          _481_v54 = _rhs107;
          let _492_v59;
          _492_v59 = _dafny.Map.Empty.slice().updateUnsafe(!(_module.__default.fm1(new BigNumber(631), _487_cf8, globalState)) || (!(_487_cf8)),_476_v50);
          let _493_v60;
          _493_v60 = _dafny.Seq.of(_476_v50);
          _492_v59 = (_492_v59).update((_this).f28, ((_476_v50).update(p1, _487_cf8)).Merge((_493_v60)[_module.__default.safeIndex((_this).f23, new BigNumber((_493_v60).length))]));
          let _494_v61;
          _494_v61 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_478_v52,_dafny.Set.fromElements(_455_v37, _455_v37)),(_dafny.ZERO).minus((_this).f23));
          let _495_v62;
          _495_v62 = _dafny.Set.fromElements(_455_v37, new _dafny.CodePoint('e'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0)), _455_v37);
          let _496_v63;
          _496_v63 = _dafny.Map.Empty.slice().updateUnsafe(_478_v52,_495_v62);
          _494_v61 = (_494_v61).update((_496_v63).Merge(_496_v63), new BigNumber(568));
        } else if (_source3.is_DC5) {
          let _497___mcc_h3 = (_source3).cf9;
          let _498___mcc_h4 = (_source3).cf10;
          let _499___mcc_h5 = (_source3).cf11;
          let _500___mcc_h6 = (_source3).cf12;
          let _501_cf12 = _500___mcc_h6;
          let _502_cf11 = _499___mcc_h5;
          let _503_cf10 = _498___mcc_h4;
          let _504_cf9 = _497___mcc_h3;
          let _505_v64;
          let _nw76 = Array((new BigNumber(21)).toNumber());
          _505_v64 = _nw76;
          let _506_v65;
          let _nw77 = new _module.C2();
          _nw77.__ctor(_482_v55, false, _503_cf10, _dafny.Seq.of(new BigNumber(970)), (_this.f29)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f23), new BigNumber((_this.f29).length))]);
          _506_v65 = _nw77;
          let _507_v66;
          _507_v66 = _506_v65;
          let _index99 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((_505_v64).length));
          (_505_v64)[_index99] = (_507_v66);
          let _index100 = _module.__default.safeIndex(new BigNumber(821), new BigNumber((_505_v64).length));
          (_505_v64)[_index100] = _506_v65;
          let _508_v67;
          _508_v67 = _dafny.Seq.of((_this).f28, p1, _503_cf10, (_this).f28);
          _502_cf11 = (_502_cf11).minus(_module.__default.safeDivisionInt(new BigNumber((_508_v67).length), (_this).f23));
          (globalState).f14 = _501_cf12;
          let _index101 = _module.__default.safeIndex(new BigNumber(836), new BigNumber((_415_v0).length));
          (_415_v0)[_index101] = !(true);
          let _index102 = _module.__default.safeIndex(new BigNumber(836), new BigNumber((_415_v0).length));
          (_415_v0)[_index102] = _503_cf10;
        } else if (_source3.is_DC6) {
          let _509___mcc_h7 = (_source3).cf13;
          let _510_cf13 = _509___mcc_h7;
          let _511_v68;
          let _nw78 = new _module.C1();
          _nw78.__ctor((_this).f28, _this.f29, _module.__default.fm0((_this).f23, (_this).f28, globalState));
          _511_v68 = _nw78;
          _511_v68 = _511_v68;
          _476_v50 = (_476_v50).update(p1, !(p1));
          let _512_v69;
          _512_v69 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,_455_v37);
          _512_v69 = (_512_v69).update((_this).f28, _455_v37);
          let _513_v70;
          _513_v70 = _dafny.Map.Empty.slice().updateUnsafe(false,_476_v50);
          (globalState).f9 = (_513_v70).contains((_this).f28);
        } else {
          let _514___mcc_h8 = (_source3).cf5;
          let _515_cf5 = _514___mcc_h8;
          let _rhs108 = p1;
          let _rhs109 = new BigNumber(408);
          let _lhs103 = globalState;
          _lhs103.f7 = _rhs108;
          r0 = _rhs109;
          let _index103 = _module.__default.safeIndex(new BigNumber(936), new BigNumber((_478_v52).length));
          (_478_v52)[_index103] = (_this).f23;
          let _index104 = _module.__default.safeIndex(new BigNumber(936), new BigNumber((_478_v52).length));
          (_478_v52)[_index104] = (_this).fm12((_this).f23, p1, (_this).f28, globalState);
          let _516_v71;
          let _517_v72;
          let _518_v73;
          let _out24;
          let _out25;
          let _out26;
          let _outcollector8 = _module.__default.m0((((_this).f28) ? (p0) : (_415_v0)), _dafny.Seq.UnicodeFromString("nplonay"), _475_i5, globalState);
          _out24 = _outcollector8[0];
          _out25 = _outcollector8[1];
          _out26 = _outcollector8[2];
          _516_v71 = _out24;
          _517_v72 = _out25;
          _518_v73 = _out26;
          (globalState).f14 = (_475_i5).plus(_518_v73);
        }
        (globalState).f14 = (_475_i5).minus(new BigNumber((_454_v36).length));
        let _index105 = _module.__default.safeIndex(new BigNumber(107), new BigNumber((_415_v0).length));
        (_415_v0)[_index105] = p1;
        let _index106 = _module.__default.safeIndex(new BigNumber(107), new BigNumber((_415_v0).length));
        (_415_v0)[_index106] = _module.__default.fm1((_this).f23, !(p1), globalState);
      }
      let _519_v74;
      _519_v74 = _dafny.MultiSet.fromElements(!(true));
      let _520_v75;
      _520_v75 = _dafny.Set.fromElements((_this).f23, _module.__default.fm0((_this).f23, (_this).f28, globalState));
      if (((_dafny.ZERO).minus((_dafny.ZERO).minus((new BigNumber((_519_v74).cardinality())).minus(new BigNumber((_520_v75).length))))).isEqualTo((_this).f23)) {
        if ((_520_v75).IsProperSubsetOf(_520_v75)) {
          (globalState).f7 = false;
          let _521_v76;
          let _init9 = function (_522_i7) {
            return _module.__default.safeDivisionInt(_522_i7, (_this).f23);
          };
          let _nw79 = Array((new BigNumber(21)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw79.length); _i0_9++) {
            _nw79[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _521_v76 = _nw79;
          let _523_v77;
          let _nw80 = Array((new BigNumber(20)).toNumber());
          _nw80[(_dafny.ZERO).toNumber()] = _521_v76;
          _nw80[(_dafny.ONE).toNumber()] = _521_v76;
          _nw80[(new BigNumber(2)).toNumber()] = _521_v76;
          _nw80[(new BigNumber(3)).toNumber()] = _521_v76;
          _nw80[(new BigNumber(4)).toNumber()] = _521_v76;
          _nw80[(new BigNumber(5)).toNumber()] = _521_v76;
          _nw80[(new BigNumber(6)).toNumber()] = _521_v76;
          _nw80[(new BigNumber(7)).toNumber()] = _521_v76;
          _nw80[(new BigNumber(8)).toNumber()] = _521_v76;
          _nw80[(new BigNumber(9)).toNumber()] = _521_v76;
          _nw80[(new BigNumber(10)).toNumber()] = _521_v76;
          _nw80[(new BigNumber(11)).toNumber()] = _521_v76;
          _nw80[(new BigNumber(12)).toNumber()] = _521_v76;
          _nw80[(new BigNumber(13)).toNumber()] = _521_v76;
          _nw80[(new BigNumber(14)).toNumber()] = _521_v76;
          _nw80[(new BigNumber(15)).toNumber()] = (((_this).f28) ? (_521_v76) : (_521_v76));
          _nw80[(new BigNumber(16)).toNumber()] = _521_v76;
          _nw80[(new BigNumber(17)).toNumber()] = _521_v76;
          _nw80[(new BigNumber(18)).toNumber()] = _521_v76;
          _nw80[(new BigNumber(19)).toNumber()] = _521_v76;
          _523_v77 = _nw80;
          let _index107 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_523_v77).length));
          (_523_v77)[_index107] = _521_v76;
          let _index108 = _module.__default.safeIndex(new BigNumber(71), new BigNumber((_523_v77).length));
          (_523_v77)[_index108] = _521_v76;
          let _524_v78;
          _524_v78 = _dafny.MultiSet.fromElements(_module.__default.safeDivisionInt((_this).f23, (_this).f23));
          (globalState).f8 = (((_524_v78).contains((_this).f23)) ? ((_524_v78).get((_this).f23)) : ((_this).f23));
          (globalState).f21 = (_this).f28;
          let _525_v79;
          _525_v79 = _dafny.MultiSet.fromElements(_521_v76, _521_v76, _521_v76);
          _525_v79 = _525_v79;
        } else {
          let _526_v80;
          _526_v80 = _module.D8.create_DC19(_dafny.Seq.update(_this.f29, _module.__default.safeIndex((_this).f23, new BigNumber((_this.f29).length)), (_this).f23));
          let _527_v81;
          _527_v81 = _dafny.MultiSet.fromElements(_this.f29, (_526_v80).dtor_cf32);
          (globalState).f21 = (_527_v81).equals(_527_v81);
          (globalState).f6 = _module.__default.fm2(globalState);
          _415_v0 = p0;
          let _528_v82;
          let _nw81 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
          _528_v82 = _nw81;
          let _index109 = _module.__default.safeIndex(new BigNumber(127), new BigNumber((_528_v82).length));
          (_528_v82)[_index109] = (_this).f23;
          let _index110 = _module.__default.safeIndex(new BigNumber(127), new BigNumber((_528_v82).length));
          (_528_v82)[_index110] = (_this).f23;
          let _529_v83;
          _529_v83 = _module.D3.create_DC11((_this).f23);
          let _530_v84;
          _530_v84 = _dafny.Map.Empty.slice().updateUnsafe(_520_v75,_529_v83);
          let _531_v86;
          _531_v86 = _dafny.Seq.of(_520_v75);
          let _index111 = _module.__default.safeIndex(new BigNumber(127), new BigNumber((_528_v82).length));
          let _rhs110 = (function () {
            let _coll19 = new _dafny.Map();
            for (const _compr_19 of (_531_v86).Elements) {
              let _532_v85 = _compr_19;
              if (_dafny.Seq.contains(_531_v86, _532_v85)) {
                _coll19.push([_532_v85,_529_v83]);
              }
            }
            return _coll19;
          }()).Merge(_530_v84);
          let _rhs111 = _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_454_v36, _dafny.Seq.UnicodeFromString("udxyoam"))).length)), (_528_v82)[_module.__default.safeIndex(new BigNumber(127), new BigNumber((_528_v82).length))]);
          let _lhs104 = _528_v82;
          let _lhs105 = _module.__default.safeIndex(new BigNumber(127), new BigNumber((_528_v82).length));
          _530_v84 = _rhs110;
          _lhs104[_lhs105] = _rhs111;
        }
        (globalState).f21 = ((_this).f23).isEqualTo(new BigNumber(805));
        let _533_v87;
        _533_v87 = _dafny.Seq.of(p1);
        let _index112 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((p0).length));
        (p0)[_index112] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.update(_533_v87, _module.__default.safeIndex((_this).f23, new BigNumber((_533_v87).length)), (_this).f28), _dafny.Seq.of(true, (_this).f28));
        let _index113 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((p0).length));
        (p0)[_index113] = p1;
        let _534_v89;
        _534_v89 = _dafny.Map.Empty.slice().updateUnsafe(true,function () {
          let _coll20 = new _dafny.Map();
          for (const _compr_20 of (_454_v36).Elements) {
            let _535_v88 = _compr_20;
            if (_dafny.Seq.contains(_454_v36, _535_v88)) {
              _coll20.push([_535_v88,(_this).f23]);
            }
          }
          return _coll20;
        }());
        let _536_v90;
        _536_v90 = _dafny.Set.fromElements((p0)[_module.__default.safeIndex(new BigNumber(347), new BigNumber((p0).length))]);
        let _537_v91;
        _537_v91 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_455_v37,new BigNumber((_534_v89).length))).length), true, globalState),_dafny.Seq.update(_dafny.Seq.update(_this.f29, _module.__default.safeIndex(new BigNumber((_536_v90).length), new BigNumber((_this.f29).length)), (_this).f23), _module.__default.safeIndex((_this).f23, new BigNumber((_dafny.Seq.update(_this.f29, _module.__default.safeIndex(new BigNumber((_536_v90).length), new BigNumber((_this.f29).length)), (_this).f23)).length)), new BigNumber(943)));
        let _538_v92;
        let _nw82 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
        _538_v92 = _nw82;
        let _539_v93;
        _539_v93 = _dafny.Map.Empty.slice().updateUnsafe(_538_v92,_this.f29);
        let _540_v94;
        let _nw83 = new _module.C1();
        _nw83.__ctor(false, (((_539_v93).contains(_538_v92)) ? ((_539_v93).get(_538_v92)) : (_this.f29)), (_this).f23);
        _540_v94 = _nw83;
        let _541_v95;
        _541_v95 = _dafny.Map.Empty.slice().updateUnsafe((((_537_v91).contains((_this).f23)) ? ((_537_v91).get((_this).f23)) : (_dafny.Seq.of(new BigNumber((_537_v91).length)))),_540_v94);
        _541_v95 = (_541_v95).update(_dafny.Seq.Concat(_this.f29, _dafny.Seq.of((_this).f23)), _540_v94);
        (globalState).f21 = true;
      } else {
        let _542_v96;
        _542_v96 = _dafny.Seq.of((_this).f28, false, p1);
        let _rhs112 = (_this).f23;
        let _rhs113 = _542_v96;
        let _rhs114 = p1;
        let _lhs106 = globalState;
        r0 = _rhs112;
        _542_v96 = _rhs113;
        _lhs106.f9 = _rhs114;
        let _543_v97;
        _543_v97 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,(_this).f28);
        _543_v97 = (_543_v97).update(new BigNumber((_454_v36).length), !(true));
        let _544_v98;
        let _nw84 = Array((new BigNumber(27)).toNumber()).fill([]);
        _544_v98 = _nw84;
        let _545_v99;
        let _init10 = ((_546_v37) => function (_547_i8) {
          return _546_v37;
        })(_455_v37);
        let _nw85 = Array((new BigNumber(28)).toNumber());
        for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw85.length); _i0_10++) {
          _nw85[_i0_10] = _init10(new BigNumber(_i0_10));
        }
        _545_v99 = _nw85;
        let _index114 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_544_v98).length));
        (_544_v98)[_index114] = _545_v99;
        let _index115 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_544_v98).length));
        let _rhs115 = (_dafny.ZERO).minus((_this.f29)[_module.__default.safeIndex((_this).f23, new BigNumber((_this.f29).length))]);
        let _rhs116 = _545_v99;
        let _rhs117 = !(p1);
        let _rhs118 = (_this).f23;
        let _lhs107 = globalState;
        let _lhs108 = _544_v98;
        let _lhs109 = _module.__default.safeIndex(new BigNumber(139), new BigNumber((_544_v98).length));
        let _lhs110 = globalState;
        let _lhs111 = globalState;
        _lhs107.f14 = _rhs115;
        _lhs108[_lhs109] = _rhs116;
        _lhs110.f9 = _rhs117;
        _lhs111.f14 = _rhs118;
        let _548_v100;
        _548_v100 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,(_this).f23);
        let _549_v101;
        let _nw86 = new _module.C2();
        _nw86.__ctor(_548_v100, p1, (!(p1)) || (p1), _dafny.Seq.Concat(_dafny.Seq.of((_this).f23, (_this).f23), _this.f29), (((_this).f28) ? ((_this).f23) : ((_this).f23)));
        _549_v101 = _nw86;
        (globalState).f8 = (_this).f23;
      }
      r0 = (_this).f23;
      let _550_v102;
      _550_v102 = _dafny.MultiSet.fromElements((_this).f23);
      let _551_v103;
      _551_v103 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,(_this).f23);
      let _552_v104;
      _552_v104 = _dafny.Map.Empty.slice().updateUnsafe(_455_v37,(((_550_v102).contains(new BigNumber((_551_v103).length))) ? ((_550_v102).get(new BigNumber((_551_v103).length))) : ((_this).f23)));
      r1 = (_520_v75).Difference(_dafny.Set.fromElements(new BigNumber((function () {
        let _coll21 = new _dafny.Set();
        for (const _compr_21 of (_552_v104).Keys.Elements) {
          let _553_v105 = _compr_21;
          if ((_552_v104).contains(_553_v105)) {
            _coll21.add(_553_v105);
          }
        }
        return _coll21;
      }()).length), (_this).f23, new BigNumber((_454_v36).length)));
      return [r0, r1];
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Seq.of();
      let _554_v0;
      let _init11 = function (_555_i0) {
        return _module.__default.safeDivisionInt(_555_i0, (_this).f23);
      };
      let _nw87 = Array((new BigNumber(26)).toNumber());
      for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw87.length); _i0_11++) {
        _nw87[_i0_11] = _init11(new BigNumber(_i0_11));
      }
      _554_v0 = _nw87;
      let _556_v1;
      _556_v1 = _dafny.Seq.of(_554_v0);
      let _557_v2;
      _557_v2 = _module.D1.create_DC3(_556_v1);
      let _558_v3;
      _558_v3 = _dafny.MultiSet.fromElements(_557_v2, _557_v2, _557_v2);
      let _559_v4;
      _559_v4 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false),_558_v3);
      let _560_v5;
      _560_v5 = _dafny.Seq.of((_this).f28);
      _559_v4 = (_559_v4).update(_560_v5, _558_v3);
      let _index116 = _module.__default.safeIndex(new BigNumber(260), new BigNumber((_554_v0).length));
      (_554_v0)[_index116] = (_this).f23;
      let _561_v6;
      _561_v6 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(317)), function (_562_i1) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      }));
      let _563_v7;
      _563_v7 = _561_v6;
      let _index117 = _module.__default.safeIndex(new BigNumber(260), new BigNumber((_554_v0).length));
      (_554_v0)[_index117] = function (_source4) {
        let _564___mcc_h0 = _source4;
        let _565_cf30 = _564___mcc_h0;
        return (_this).f23;
      }((((_this).f28) ? (_563_v7) : (_561_v6)));
      let _566_i2;
      _566_i2 = _dafny.ZERO;
      L7: {
        while (!((_this).f28) || (_dafny.Seq.contains(_this.f29, (_this).f23))) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_566_i2)) {
              break L7;
            }
            _566_i2 = (_566_i2).plus(_dafny.ONE);
            let _567_v8;
            _567_v8 = _dafny.MultiSet.fromElements((_this).f28);
            let _568_v9;
            _568_v9 = new _dafny.CodePoint('g'.codePointAt(0));
            let _569_v10;
            _569_v10 = _module.D2.create_DC8((_this).f23, new BigNumber((_567_v8).cardinality()), _568_v9, new BigNumber((p0).length), (_this).f28);
            let _rhs119 = !(!((_569_v10).dtor_cf19));
            let _rhs120 = ((_this).f23).plus((new BigNumber(692)).plus((_554_v0)[_module.__default.safeIndex(new BigNumber(260), new BigNumber((_554_v0).length))]));
            let _rhs121 = !((_567_v8).update((_this).f28, _module.__default.abs((_554_v0)[_module.__default.safeIndex(new BigNumber(260), new BigNumber((_554_v0).length))]))).equals(_567_v8);
            let _lhs112 = globalState;
            let _lhs113 = globalState;
            let _lhs114 = globalState;
            _lhs112.f7 = _rhs119;
            _lhs113.f8 = _rhs120;
            _lhs114.f21 = _rhs121;
            (globalState).f21 = false;
            (globalState).f11 = _568_v9;
            (globalState).f8 = (_554_v0)[_module.__default.safeIndex(new BigNumber(260), new BigNumber((_554_v0).length))];
          }
        }
      }
      let _570_v11;
      let _init12 = function (_571_i3) {
        return new _dafny.CodePoint('p'.codePointAt(0));
      };
      let _nw88 = Array((new BigNumber(15)).toNumber());
      for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw88.length); _i0_12++) {
        _nw88[_i0_12] = _init12(new BigNumber(_i0_12));
      }
      _570_v11 = _nw88;
      let _572_v12;
      _572_v12 = _dafny.Map.Empty.slice().updateUnsafe(_570_v11,(_this).f28);
      _572_v12 = (_572_v12).update(_570_v11, (_this).f28);
      if (false) {
        let _573_v14;
        _573_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_560_v5).length),(_this).f28);
        let _574_v15;
        _574_v15 = _dafny.Seq.of(_dafny.Seq.of(new BigNumber((_573_v14).length)));
        let _575_v16;
        let _nw89 = new _module.C2();
        _nw89.__ctor(function () {
          let _coll22 = new _dafny.Map();
          for (const _compr_22 of ((_574_v15)[_module.__default.safeIndex((_this).f23, new BigNumber((_574_v15).length))]).Elements) {
            let _576_v13 = _compr_22;
            if (_dafny.Seq.contains((_574_v15)[_module.__default.safeIndex((_this).f23, new BigNumber((_574_v15).length))], _576_v13)) {
              _coll22.push([(_576_v13).minus((_this).f23),(_this).f23]);
            }
          }
          return _coll22;
        }(), (_this).f28, (_this).f28, _this.f29, (_554_v0)[_module.__default.safeIndex(new BigNumber(260), new BigNumber((_554_v0).length))]);
        _575_v16 = _nw89;
        let _577_v17;
        _577_v17 = _dafny.MultiSet.fromElements((_575_v16).f32, (_575_v16).f32, (_this).f28);
        (globalState).f6 = (_561_v6)[_module.__default.safeIndex(new BigNumber((_577_v17).cardinality()), new BigNumber((_561_v6).length))];
        let _578_v18;
        let _init13 = ((_579_p0) => function (_580_i4) {
          return _579_p0;
        })(p0);
        let _nw90 = Array((new BigNumber(17)).toNumber());
        for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw90.length); _i0_13++) {
          _nw90[_i0_13] = _init13(new BigNumber(_i0_13));
        }
        _578_v18 = _nw90;
        _578_v18 = _578_v18;
        let _581_v19;
        let _nw91 = Array((new BigNumber(21)).toNumber()).fill(false);
        _581_v19 = _nw91;
        let _582_v20;
        _582_v20 = _dafny.Map.Empty.slice().updateUnsafe(_581_v19,((_dafny.ZERO).minus(new BigNumber(-555))).isLessThanOrEqualTo(new BigNumber(945)));
        (globalState).f21 = (((_582_v20).contains(_581_v19)) ? ((_582_v20).get(_581_v19)) : ((_this).f28));
        let _index118 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_581_v19).length));
        (_581_v19)[_index118] = (_575_v16).f32;
        let _583_v21;
        _583_v21 = _dafny.Map.Empty.slice().updateUnsafe((_575_v16).f32,(_575_v16).f32);
        let _584_v22;
        _584_v22 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,new BigNumber((_583_v21).length));
        let _585_v23;
        _585_v23 = _dafny.MultiSet.fromElements(new BigNumber(39), (_554_v0)[_module.__default.safeIndex(new BigNumber(260), new BigNumber((_554_v0).length))]);
        let _index119 = _module.__default.safeIndex(new BigNumber(409), new BigNumber((_581_v19).length));
        (_581_v19)[_index119] = !((new BigNumber((_584_v22).length)).minus((_554_v0)[_module.__default.safeIndex(new BigNumber(260), new BigNumber((_554_v0).length))])).isEqualTo((((_585_v23).contains((_this).f23)) ? ((_585_v23).get((_this).f23)) : (new BigNumber((_560_v5).length))));
      } else {
        let _586_v24;
        _586_v24 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_561_v6).length)),new BigNumber(63));
        let _587_v25;
        let _nw92 = new _module.C2();
        _nw92.__ctor(_586_v24, (_this).f28, (_this).f28, _this.f29, (_554_v0)[_module.__default.safeIndex(new BigNumber(260), new BigNumber((_554_v0).length))]);
        _587_v25 = _nw92;
        let _588_v26;
        _588_v26 = _587_v25;
        let _source5 = _588_v26;
        let _589___mcc_h1 = _source5;
        let _590_cf29 = _589___mcc_h1;
        let _591_v27;
        _591_v27 = _dafny.MultiSet.fromElements((_587_v25).f23, (_590_cf29).f23, (_dafny.ZERO).minus(new BigNumber((_560_v5).length)));
        let _592_v28;
        _592_v28 = _dafny.MultiSet.fromElements(_591_v27, _dafny.MultiSet.FromArray(_this.f29), (_591_v27).Union(_591_v27), (_591_v27).update((_554_v0)[_module.__default.safeIndex(new BigNumber(260), new BigNumber((_554_v0).length))], _module.__default.abs((_587_v25).f23)), _591_v27);
        let _593_v29;
        _593_v29 = _dafny.MultiSet.fromElements(_570_v11, _570_v11, _570_v11);
        let _rhs122 = _module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f23), ((_this).f23).plus(new BigNumber((p0).length)));
        let _rhs123 = _dafny.MultiSet.fromElements((_dafny.MultiSet.fromElements((_587_v25).f23, new BigNumber((_593_v29).cardinality()), (_587_v25).f23, (_554_v0)[_module.__default.safeIndex(new BigNumber(260), new BigNumber((_554_v0).length))], new BigNumber((p0).length))).update((_587_v25).f23, _module.__default.abs((_590_cf29).f23)), _591_v27, (_591_v27).Intersect(_591_v27), _591_v27);
        r0 = _rhs122;
        _592_v28 = _rhs123;
        let _594_v30;
        let _init14 = ((_595_v25, _596_cf29) => function (_597_i5) {
          return (_dafny.MultiSet.fromElements(false, !((_595_v25).f28), (_596_cf29).f28, (_596_cf29).f28)).Union(_dafny.MultiSet.fromElements((_595_v25).f28, false, !(false), (_595_v25).f28));
        })(_587_v25, _590_cf29);
        let _nw93 = Array((new BigNumber(27)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw93.length); _i0_14++) {
          _nw93[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _594_v30 = _nw93;
        let _rhs124 = (_587_v25).f28;
        let _rhs125 = _594_v30;
        let _lhs115 = globalState;
        _lhs115.f21 = _rhs124;
        _594_v30 = _rhs125;
        let _598_v31;
        _598_v31 = _module.D3.create_DC11((_554_v0)[_module.__default.safeIndex(new BigNumber(260), new BigNumber((_554_v0).length))]);
        let _pat_let_tv27 = _587_v25;
        let _599_v32;
        _599_v32 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.of(function (_pat_let2_0) {
          return function (_600_dt__update__tmp_h0) {
            return function (_pat_let3_0) {
              return function (_601_dt__update_hcf22_h0) {
                return _module.D3.create_DC11(_601_dt__update_hcf22_h0);
              }(_pat_let3_0);
            }((_pat_let_tv27).f23);
          }(_pat_let2_0);
        }(_598_v31)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(490)), ((_602_v31) => function (_603_i6) {
          return _602_v31;
        })(_598_v31))),new BigNumber(-173));
        let _604_v33;
        _604_v33 = _dafny.Seq.of(_module.D3.create_DC11((_590_cf29).f23), _598_v31);
        _599_v32 = (_599_v32).update((((_590_cf29).f28) ? (_604_v33) : (_604_v33)), (_554_v0)[_module.__default.safeIndex(new BigNumber(260), new BigNumber((_554_v0).length))]);
        let _605_v34;
        _605_v34 = _dafny.MultiSet.fromElements(_554_v0, _554_v0, _554_v0);
        (globalState).f7 = ((_dafny.MultiSet.fromElements(_554_v0)).update((_556_v1)[_module.__default.safeIndex((_590_cf29).f23, new BigNumber((_556_v1).length))], _module.__default.abs((_587_v25).f23))).IsProperSubsetOf(_605_v34);
        let _606_v35;
        _606_v35 = _dafny.Map.Empty.slice().updateUnsafe(_554_v0,_587_v25.f29);
        let _607_v36;
        _607_v36 = _dafny.Map.Empty.slice().updateUnsafe(false,(_606_v35).update(_554_v0, _this.f29));
        _607_v36 = (_607_v36).update((_587_v25).f28, _606_v35);
        if (!(_module.__default.fm1((_587_v25).f23, (_587_v25).f28, globalState))) {
          (globalState).f7 = (_587_v25).f28;
          _570_v11 = _570_v11;
          (globalState).f9 = false;
          _587_v25 = _587_v25;
          let _608_v37;
          let _nw94 = Array((new BigNumber(5)).toNumber());
          _nw94[(_dafny.ZERO).toNumber()] = (_587_v25).f28;
          _nw94[(_dafny.ONE).toNumber()] = (_587_v25).f28;
          _nw94[(new BigNumber(2)).toNumber()] = _module.__default.fm1((_dafny.ZERO).minus((_554_v0)[_module.__default.safeIndex(new BigNumber(260), new BigNumber((_554_v0).length))]), (_587_v25).f28, globalState);
          _nw94[(new BigNumber(3)).toNumber()] = (_587_v25).f28;
          _nw94[(new BigNumber(4)).toNumber()] = _module.__default.fm1((_554_v0)[_module.__default.safeIndex(new BigNumber(260), new BigNumber((_554_v0).length))], (_587_v25).f28, globalState);
          _608_v37 = _nw94;
          let _index120 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_608_v37).length));
          (_608_v37)[_index120] = (_this).f28;
          let _609_v38;
          _609_v38 = new _dafny.CodePoint('a'.codePointAt(0));
          let _610_v39;
          _610_v39 = _module.D9.create_DC24(_570_v11);
          let _index121 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_608_v37).length));
          let _rhs126 = !_dafny.Seq.contains(p0, _609_v38);
          let _rhs127 = new BigNumber(-741);
          let _rhs128 = (_610_v39).dtor_cf40;
          let _rhs129 = _dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(729)), ((_611_v38) => function (_612_i7) {
            return _611_v38;
          })(_609_v38)), p0);
          let _rhs130 = _608_v37;
          let _lhs116 = globalState;
          let _lhs117 = globalState;
          let _lhs118 = _608_v37;
          let _lhs119 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_608_v37).length));
          _lhs116.f21 = _rhs126;
          _lhs117.f14 = _rhs127;
          _570_v11 = _rhs128;
          _lhs118[_lhs119] = _rhs129;
          _608_v37 = _rhs130;
        } else {
          let _613_v40;
          _613_v40 = _dafny.Map.Empty.slice().updateUnsafe(_560_v5,(_this).f28);
          let _614_v42;
          _614_v42 = new _dafny.CodePoint('w'.codePointAt(0));
          let _615_v43;
          _615_v43 = _dafny.Map.Empty.slice().updateUnsafe(_614_v42,(_587_v25).f28);
          _613_v40 = (((function () {
            let _coll23 = new _dafny.Map();
            for (const _compr_23 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(770)), function (_616_i8) {
              return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(177)), function (_617_i9) {
                return new _dafny.CodePoint('y'.codePointAt(0));
              })).length);
            })).Elements) {
              let _618_v41 = _compr_23;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(770)), function (_616_i8) {
                return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(177)), function (_617_i9) {
                  return new _dafny.CodePoint('y'.codePointAt(0));
                })).length);
              }), _618_v41)) {
                _coll23.push([_module.__default.safeModuloInt(_618_v41, (_587_v25).f23),(_554_v0)[_module.__default.safeIndex(new BigNumber(260), new BigNumber((_554_v0).length))]]);
              }
            }
            return _coll23;
          }()).equals(_module.__default.fm16((_this).f28, new BigNumber((_615_v43).length), (_this).f23, globalState))) ? (_dafny.Map.Empty.slice().updateUnsafe(_560_v5,(_this).f28)) : (_613_v40));
          let _619_v44;
          _619_v44 = _module.D9.create_DC26(new BigNumber(548));
          let _pat_let_tv28 = _554_v0;
          let _pat_let_tv29 = _554_v0;
          let _620_v45;
          let _nw95 = Array((new BigNumber(22)).toNumber());
          _nw95[(_dafny.ZERO).toNumber()] = _module.D9.create_DC26((_554_v0)[_module.__default.safeIndex(new BigNumber(260), new BigNumber((_554_v0).length))]);
          _nw95[(_dafny.ONE).toNumber()] = _619_v44;
          _nw95[(new BigNumber(2)).toNumber()] = _619_v44;
          _nw95[(new BigNumber(3)).toNumber()] = _619_v44;
          _nw95[(new BigNumber(4)).toNumber()] = _619_v44;
          _nw95[(new BigNumber(5)).toNumber()] = _619_v44;
          _nw95[(new BigNumber(6)).toNumber()] = _619_v44;
          _nw95[(new BigNumber(7)).toNumber()] = _619_v44;
          _nw95[(new BigNumber(8)).toNumber()] = _619_v44;
          _nw95[(new BigNumber(9)).toNumber()] = _619_v44;
          _nw95[(new BigNumber(10)).toNumber()] = _619_v44;
          _nw95[(new BigNumber(11)).toNumber()] = _619_v44;
          _nw95[(new BigNumber(12)).toNumber()] = _619_v44;
          _nw95[(new BigNumber(13)).toNumber()] = _module.D9.create_DC26(new BigNumber(604));
          _nw95[(new BigNumber(14)).toNumber()] = function (_pat_let4_0) {
            return function (_621_dt__update__tmp_h1) {
              return function (_pat_let5_0) {
                return function (_622_dt__update_hcf43_h0) {
                  return _module.D9.create_DC26(_622_dt__update_hcf43_h0);
                }(_pat_let5_0);
              }((_this).f23);
            }(_pat_let4_0);
          }(_619_v44);
          _nw95[(new BigNumber(15)).toNumber()] = _619_v44;
          _nw95[(new BigNumber(16)).toNumber()] = _619_v44;
          _nw95[(new BigNumber(17)).toNumber()] = _619_v44;
          _nw95[(new BigNumber(18)).toNumber()] = function (_pat_let6_0) {
            return function (_623_dt__update__tmp_h2) {
              return function (_pat_let7_0) {
                return function (_624_dt__update_hcf43_h1) {
                  return _module.D9.create_DC26(_624_dt__update_hcf43_h1);
                }(_pat_let7_0);
              }((_pat_let_tv29)[_module.__default.safeIndex(new BigNumber(260), new BigNumber((_pat_let_tv28).length))]);
            }(_pat_let6_0);
          }(_619_v44);
          _nw95[(new BigNumber(19)).toNumber()] = _619_v44;
          _nw95[(new BigNumber(20)).toNumber()] = _619_v44;
          _nw95[(new BigNumber(21)).toNumber()] = _619_v44;
          _620_v45 = _nw95;
          let _index122 = _module.__default.safeIndex(new BigNumber(94), new BigNumber((_620_v45).length));
          (_620_v45)[_index122] = _619_v44;
          let _index123 = _module.__default.safeIndex(new BigNumber(94), new BigNumber((_620_v45).length));
          (_620_v45)[_index123] = _619_v44;
          let _index124 = _module.__default.safeIndex(new BigNumber(260), new BigNumber((_554_v0).length));
          (_554_v0)[_index124] = new BigNumber(721);
          let _625_v46;
          let _nw96 = Array((new BigNumber(9)).toNumber()).fill([]);
          _625_v46 = _nw96;
          let _626_v47;
          _626_v47 = _dafny.MultiSet.fromElements(true);
          let _627_v48;
          _627_v48 = _dafny.Seq.of(_dafny.Set.fromElements((_this).f28));
          let _628_v49;
          _628_v49 = _module.D0.create_DC0(p0);
          let _629_v50;
          _629_v50 = _dafny.MultiSet.fromElements(new BigNumber((p0).length));
          let _630_v51;
          let _nw97 = Array((new BigNumber(22)).toNumber());
          _nw97[(_dafny.ZERO).toNumber()] = _626_v47;
          _nw97[(_dafny.ONE).toNumber()] = _626_v47;
          _nw97[(new BigNumber(2)).toNumber()] = _626_v47;
          _nw97[(new BigNumber(3)).toNumber()] = _626_v47;
          _nw97[(new BigNumber(4)).toNumber()] = _626_v47;
          _nw97[(new BigNumber(5)).toNumber()] = _626_v47;
          _nw97[(new BigNumber(6)).toNumber()] = _626_v47;
          _nw97[(new BigNumber(7)).toNumber()] = _dafny.MultiSet.fromElements((_587_v25).f28, (_this).f28);
          _nw97[(new BigNumber(8)).toNumber()] = _dafny.MultiSet.fromElements((_587_v25).f28, (_this).f28);
          _nw97[(new BigNumber(9)).toNumber()] = _626_v47;
          _nw97[(new BigNumber(10)).toNumber()] = _626_v47;
          _nw97[(new BigNumber(11)).toNumber()] = _626_v47;
          _nw97[(new BigNumber(12)).toNumber()] = (((_this).f30)[_module.__default.safeIndex((_this).f23, new BigNumber(((_this).f30).length))]).update((_this).f28, _module.__default.abs(new BigNumber((_627_v48).length)));
          _nw97[(new BigNumber(13)).toNumber()] = _626_v47;
          _nw97[(new BigNumber(14)).toNumber()] = _626_v47;
          _nw97[(new BigNumber(15)).toNumber()] = (_626_v47).update((_587_v25).f28, _module.__default.abs((_587_v25).f23));
          _nw97[(new BigNumber(16)).toNumber()] = _dafny.MultiSet.fromElements((_587_v25).f28, (_587_v25).f28);
          _nw97[(new BigNumber(17)).toNumber()] = _module.__default.fm4(_628_v49, (_this).f23, _560_v5, _614_v42, globalState);
          _nw97[(new BigNumber(18)).toNumber()] = _dafny.MultiSet.FromArray(_560_v5);
          _nw97[(new BigNumber(19)).toNumber()] = _dafny.MultiSet.fromElements((_587_v25).f28, (_this).f28);
          _nw97[(new BigNumber(20)).toNumber()] = _626_v47;
          _nw97[(new BigNumber(21)).toNumber()] = _module.__default.fm4(_628_v49, (((_629_v50).contains(new BigNumber((_626_v47).cardinality()))) ? ((_629_v50).get(new BigNumber((_626_v47).cardinality()))) : ((_587_v25).f23)), _560_v5, _614_v42, globalState);
          _630_v51 = _nw97;
          let _index125 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_625_v46).length));
          (_625_v46)[_index125] = _630_v51;
          let _631_v52;
          _631_v52 = _dafny.Map.Empty.slice().updateUnsafe((_587_v25).f23,(_this).f28);
          let _632_v53;
          let _init15 = function (_633_i10) {
            return false;
          };
          let _nw98 = Array((new BigNumber(27)).toNumber());
          for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw98.length); _i0_15++) {
            _nw98[_i0_15] = _init15(new BigNumber(_i0_15));
          }
          _632_v53 = _nw98;
          let _634_v54;
          _634_v54 = _dafny.Map.Empty.slice().updateUnsafe(_632_v53,_module.__default.fm1((_587_v25).f23, (_this).f28, globalState));
          let _index126 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_625_v46).length));
          let _rhs131 = (new BigNumber((_634_v54).length)).isLessThanOrEqualTo(new BigNumber((_631_v52).length));
          let _rhs132 = (_587_v25).f23;
          let _rhs133 = _630_v51;
          let _lhs120 = globalState;
          let _lhs121 = globalState;
          let _lhs122 = _625_v46;
          let _lhs123 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_625_v46).length));
          _lhs120.f21 = _rhs131;
          _lhs121.f8 = _rhs132;
          _lhs122[_lhs123] = _rhs133;
          let _635_v55;
          let _636_v56;
          let _637_v57;
          let _out27;
          let _out28;
          let _out29;
          let _outcollector9 = _module.__default.m0(_632_v53, _dafny.Seq.Concat(p0, _dafny.Seq.update(p0, _module.__default.safeIndex((_this).f23, new BigNumber((p0).length)), _614_v42)), (_587_v25).f23, globalState);
          _out27 = _outcollector9[0];
          _out28 = _outcollector9[1];
          _out29 = _outcollector9[2];
          _635_v55 = _out27;
          _636_v56 = _out28;
          _637_v57 = _out29;
        }
        let _638_v58;
        _638_v58 = _module.D8.create_DC22((_this).f28);
        (globalState).f8 = (((_638_v58).dtor_cf38) ? ((_this).f23) : ((_554_v0)[_module.__default.safeIndex(new BigNumber(260), new BigNumber((_554_v0).length))]));
        (globalState).f14 = (_this).f23;
      }
      let _639_v59;
      let _nw99 = new _module.C0();
      _nw99.__ctor();
      _639_v59 = _nw99;
      let _640_v60;
      _640_v60 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("rtk"),(_this).f23);
      let _641_v61;
      _641_v61 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(499),_dafny.Seq.UnicodeFromString("gun"));
      let _642_v62;
      _642_v62 = new _dafny.CodePoint('a'.codePointAt(0));
      let _643_v63;
      _643_v63 = _dafny.Set.fromElements((_this).f23, new BigNumber(-620), (_this).f23, (_554_v0)[_module.__default.safeIndex(new BigNumber(260), new BigNumber((_554_v0).length))], (_this).f23);
      let _644_v64;
      _644_v64 = _dafny.Map.Empty.slice().updateUnsafe((((_640_v60).contains(_dafny.Seq.update((((_641_v61).contains(new BigNumber(-191))) ? ((_641_v61).get(new BigNumber(-191))) : (p0)), _module.__default.safeIndex((_this).f23, new BigNumber(((((_641_v61).contains(new BigNumber(-191))) ? ((_641_v61).get(new BigNumber(-191))) : (p0))).length)), _642_v62))) ? ((_640_v60).get(_dafny.Seq.update((((_641_v61).contains(new BigNumber(-191))) ? ((_641_v61).get(new BigNumber(-191))) : (p0)), _module.__default.safeIndex((_this).f23, new BigNumber(((((_641_v61).contains(new BigNumber(-191))) ? ((_641_v61).get(new BigNumber(-191))) : (p0))).length)), _642_v62))) : (new BigNumber((_643_v63).length))),(_dafny.ZERO).minus((_554_v0)[_module.__default.safeIndex(new BigNumber(260), new BigNumber((_554_v0).length))]));
      r0 = (((_644_v64).contains((_dafny.ZERO).minus((_this).f23))) ? ((_644_v64).get((_dafny.ZERO).minus((_this).f23))) : ((_554_v0)[_module.__default.safeIndex(new BigNumber(260), new BigNumber((_554_v0).length))]));
      r1 = _dafny.Seq.update(_dafny.Seq.Concat(_this.f29, _this.f29), _module.__default.safeIndex((_554_v0)[_module.__default.safeIndex(new BigNumber(260), new BigNumber((_554_v0).length))], new BigNumber((_dafny.Seq.Concat(_this.f29, _this.f29)).length)), (_554_v0)[_module.__default.safeIndex(new BigNumber(260), new BigNumber((_554_v0).length))]);
      return [r0, r1];
    }
    m5(globalState) {
      let _this = this;
      let _645_v0;
      let _nw100 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
      _645_v0 = _nw100;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_645_v0).length))) {
        let _646_i0 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_646_i0)) && ((_646_i0).isLessThan(new BigNumber((_645_v0).length))))) {
          (_645_v0)[(_646_i0)] = _module.__default.safeModuloInt(_646_i0, (new BigNumber((_this.f29).length)).minus(new BigNumber((_dafny.Seq.UnicodeFromString("fgr")).length)));
        }
      }
      let _647_i1;
      _647_i1 = _dafny.ZERO;
      L8: {
        while (!((_this).f28)) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_647_i1)) {
              break L8;
            }
            _647_i1 = (_647_i1).plus(_dafny.ONE);
            let _648_v1;
            let _nw101 = new _module.C0();
            _nw101.__ctor();
            _648_v1 = _nw101;
            let _649_v2;
            let _nw102 = Array((new BigNumber(23)).toNumber()).fill(false);
            _649_v2 = _nw102;
            let _650_v3;
            _650_v3 = _dafny.Set.fromElements((_dafny.ZERO).minus((_this).f23));
            let _651_v4;
            _651_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,(_this).f23);
            let _index127 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((_649_v2).length));
            (_649_v2)[_index127] = (_650_v3).IsProperSubsetOf(_dafny.Set.fromElements((((_651_v4).contains((_this.f29)[_module.__default.safeIndex((_this).f23, new BigNumber((_this.f29).length))])) ? ((_651_v4).get((_this.f29)[_module.__default.safeIndex((_this).f23, new BigNumber((_this.f29).length))])) : ((_this).f23))));
            let _index128 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((_649_v2).length));
            (_649_v2)[_index128] = !(((!((_this).f28)) ? ((function () {
              let _coll24 = new _dafny.Set();
              for (const _compr_24 of _dafny.IntegerRange(new BigNumber(979), new BigNumber(687))) {
                let _652_v5 = _compr_24;
                if (((new BigNumber(979)).isLessThanOrEqualTo(_652_v5)) && ((_652_v5).isLessThan(new BigNumber(687)))) {
                  _coll24.add((_652_v5).plus((_this).f23));
                }
              }
              return _coll24;
            }()).IsDisjointFrom(_650_v3)) : ((((_this).f28) ? ((_this).f28) : ((_this).f28)))));
            let _653_v6;
            _653_v6 = _dafny.Set.fromElements(false, (_this).f28);
            let _654_v7;
            _654_v7 = _dafny.Seq.of(_653_v6);
            let _rhs134 = (_this).f23;
            let _rhs135 = (_654_v7)[_module.__default.safeIndex((_this).f23, new BigNumber((_654_v7).length))];
            let _rhs136 = (_this.f29)[_module.__default.safeIndex((_this).f23, new BigNumber((_this.f29).length))];
            let _lhs124 = globalState;
            let _lhs125 = globalState;
            _lhs124.f14 = _rhs134;
            _653_v6 = _rhs135;
            _lhs125.f8 = _rhs136;
            let _index129 = _module.__default.safeIndex(new BigNumber(174), new BigNumber((_649_v2).length));
            (_649_v2)[_index129] = (_this).f28;
          }
        }
      }
      let _655_v8;
      _655_v8 = _dafny.Seq.UnicodeFromString("kgw");
      let _656_v9;
      _656_v9 = _dafny.MultiSet.fromElements(new BigNumber((_655_v8).length), (_this).f23);
      let _657_v10;
      _657_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,_this.f29);
      let _hi10 = (((_656_v9).contains(new BigNumber(685))) ? ((_656_v9).get(new BigNumber(685))) : (new BigNumber(((((_657_v10).contains(new BigNumber(260))) ? ((_657_v10).get(new BigNumber(260))) : (_dafny.Seq.of((_this).f23, (_this).f23)))).length)));
      for (let _658_i2 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_module.__default.fm17((_this).f28, _module.__default.fm0((_this).f23, (_this).f28, globalState), globalState)).cardinality()), (_this).f23), _this.f29)).length); _658_i2.isLessThan(_hi10); _658_i2 = _658_i2.plus(_dafny.ONE)) {
        let _659_v11;
        _659_v11 = _dafny.Seq.of((_this).f28);
        let _660_v12;
        _660_v12 = _dafny.MultiSet.fromElements(!((_this).f28), (_this).f28);
        let _661_v13;
        _661_v13 = _dafny.Map.Empty.slice().updateUnsafe(false,true);
        let _662_v14;
        _662_v14 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,_661_v13);
        let _663_v15;
        _663_v15 = _dafny.Set.fromElements(_658_i2, _module.__default.safeModuloInt((_this).f23, new BigNumber((_656_v9).cardinality())), (_658_i2).minus(new BigNumber((_659_v11).length)), (new BigNumber(((_660_v12).update((_this).f28, _module.__default.abs(_658_i2))).cardinality())).multipliedBy((_this).f23), new BigNumber((_662_v14).length));
        _663_v15 = ((false) ? (_663_v15) : (_663_v15));
        let _664_v16;
        _664_v16 = new _dafny.CodePoint('e'.codePointAt(0));
        let _665_v17;
        _665_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_this.f29).length),new BigNumber((_dafny.Seq.of(_dafny.Seq.update(_dafny.Seq.update(_655_v8, _module.__default.safeIndex((_this).f23, new BigNumber((_655_v8).length)), _664_v16), _module.__default.safeIndex((_this).f23, new BigNumber((_dafny.Seq.update(_655_v8, _module.__default.safeIndex((_this).f23, new BigNumber((_655_v8).length)), _664_v16)).length)), _664_v16), _655_v8)).length));
        let _666_v18;
        _666_v18 = _module.D4.create_DC14(_645_v0);
        let _667_v19;
        let _nw103 = new _module.C2();
        _nw103.__ctor(_665_v17, _module.__default.fm1(_658_i2, (_this).f28, globalState), (_this).f28, _dafny.Seq.Concat(_dafny.Seq.update(_this.f29, _module.__default.safeIndex(_658_i2, new BigNumber((_this.f29).length)), _658_i2), _this.f29), new BigNumber((_dafny.Seq.of(_666_v18, _666_v18)).length));
        _667_v19 = _nw103;
        _665_v17 = (_665_v17).update(new BigNumber((_655_v8).length), (_module.__default.fm24((_667_v19).f32, _658_i2, globalState)).dtor_cf43);
        if ((_this).f28) {
          let _index130 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_645_v0).length));
          (_645_v0)[_index130] = _658_i2;
          let _668_v20;
          let _nw104 = Array((new BigNumber(6)).toNumber());
          _nw104[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('n'.codePointAt(0));
          _nw104[(_dafny.ONE).toNumber()] = _664_v16;
          _nw104[(new BigNumber(2)).toNumber()] = _664_v16;
          _nw104[(new BigNumber(3)).toNumber()] = _664_v16;
          _nw104[(new BigNumber(4)).toNumber()] = _664_v16;
          _nw104[(new BigNumber(5)).toNumber()] = _664_v16;
          _668_v20 = _nw104;
          let _669_v21;
          _669_v21 = _module.D9.create_DC24(_668_v20);
          let _670_v22;
          let _nw105 = Array((new BigNumber(19)).toNumber());
          _nw105[(_dafny.ZERO).toNumber()] = _668_v20;
          _nw105[(_dafny.ONE).toNumber()] = _668_v20;
          _nw105[(new BigNumber(2)).toNumber()] = _668_v20;
          _nw105[(new BigNumber(3)).toNumber()] = _668_v20;
          _nw105[(new BigNumber(4)).toNumber()] = _668_v20;
          _nw105[(new BigNumber(5)).toNumber()] = _668_v20;
          _nw105[(new BigNumber(6)).toNumber()] = _668_v20;
          _nw105[(new BigNumber(7)).toNumber()] = _668_v20;
          _nw105[(new BigNumber(8)).toNumber()] = _668_v20;
          _nw105[(new BigNumber(9)).toNumber()] = _668_v20;
          _nw105[(new BigNumber(10)).toNumber()] = _668_v20;
          _nw105[(new BigNumber(11)).toNumber()] = _668_v20;
          _nw105[(new BigNumber(12)).toNumber()] = _668_v20;
          _nw105[(new BigNumber(13)).toNumber()] = _668_v20;
          _nw105[(new BigNumber(14)).toNumber()] = _668_v20;
          _nw105[(new BigNumber(15)).toNumber()] = _668_v20;
          _nw105[(new BigNumber(16)).toNumber()] = (_669_v21).dtor_cf40;
          _nw105[(new BigNumber(17)).toNumber()] = _668_v20;
          _nw105[(new BigNumber(18)).toNumber()] = _668_v20;
          _670_v22 = _nw105;
          let _671_v23;
          let _nw106 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Seq.of());
          _671_v23 = _nw106;
          let _index131 = _module.__default.safeIndex(new BigNumber(742), new BigNumber((_671_v23).length));
          (_671_v23)[_index131] = _this.f29;
          let _index132 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_645_v0).length));
          let _index133 = _module.__default.safeIndex(new BigNumber(742), new BigNumber((_671_v23).length));
          let _rhs137 = new BigNumber(938);
          let _rhs138 = _670_v22;
          let _rhs139 = _this.f29;
          let _rhs140 = (_this).f28;
          let _lhs126 = _645_v0;
          let _lhs127 = _module.__default.safeIndex(new BigNumber(888), new BigNumber((_645_v0).length));
          let _lhs128 = _671_v23;
          let _lhs129 = _module.__default.safeIndex(new BigNumber(742), new BigNumber((_671_v23).length));
          let _lhs130 = globalState;
          _lhs126[_lhs127] = _rhs137;
          _670_v22 = _rhs138;
          _lhs128[_lhs129] = _rhs139;
          _lhs130.f21 = _rhs140;
          let _672_v24;
          _672_v24 = _module.D1.create_DC5(_658_i2, (_667_v19).f32, new BigNumber(284), (_this).f23);
          (globalState).f9 = (_672_v24).dtor_cf10;
          (globalState).f11 = _664_v16;
          let _673_v25;
          _673_v25 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,(_667_v19).f32);
          let _674_v26;
          _674_v26 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,new BigNumber((_673_v25).length));
          let _pat_let_tv30 = _655_v8;
          let _pat_let_tv31 = _645_v0;
          let _pat_let_tv32 = _645_v0;
          let _pat_let_tv33 = _655_v8;
          let _pat_let_tv34 = _664_v16;
          (globalState).f14 = (_667_v19).fm18((((_this).f28) ? (_674_v26) : (_module.__default.fm25(new BigNumber((_dafny.MultiSet.fromElements(_664_v16)).cardinality()), (_667_v19).f32, globalState))), _dafny.Seq.Concat((function (_pat_let8_0) {
            return function (_675_dt__update__tmp_h0) {
              return function (_pat_let9_0) {
                return function (_676_dt__update_hcf0_h0) {
                  return _module.D0.create_DC0(_676_dt__update_hcf0_h0);
                }(_pat_let9_0);
              }(_dafny.Seq.update(_pat_let_tv30, _module.__default.safeIndex((_dafny.ZERO).minus((_pat_let_tv32)[_module.__default.safeIndex(new BigNumber(888), new BigNumber((_pat_let_tv31).length))]), new BigNumber((_pat_let_tv33).length)), _pat_let_tv34));
            }(_pat_let8_0);
          }(_module.D0.create_DC0(_dafny.Seq.UnicodeFromString("lvnjskw")))).dtor_cf0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(387)), ((_677_v16) => function (_678_i3) {
            return _677_v16;
          })(_664_v16))), (_667_v19).f32, globalState);
          (globalState).f7 = (_667_v19).f32;
        } else {
          let _679_v27;
          _679_v27 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,_658_i2);
          let _680_v28;
          _680_v28 = _679_v27;
          _680_v28 = _679_v27;
          (_this).f29 = _dafny.Seq.of((_667_v19).fm18(_679_v27, _dafny.Seq.UnicodeFromString("epl"), (_659_v11)[_module.__default.safeIndex(_658_i2, new BigNumber((_659_v11).length))], globalState), _module.__default.fm0(_658_i2, (_667_v19).f32, globalState), (_this).f23);
          let _index134 = _module.__default.safeIndex(new BigNumber(661), new BigNumber((_645_v0).length));
          (_645_v0)[_index134] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("pnttsu")).length));
          let _index135 = _module.__default.safeIndex(new BigNumber(661), new BigNumber((_645_v0).length));
          let _rhs141 = _module.__default.safeModuloInt(new BigNumber((_655_v8).length), (_this).f23);
          let _rhs142 = _module.__default.safeModuloInt((_this).f23, _module.__default.safeDivisionInt((_this).f23, new BigNumber((_dafny.Seq.UnicodeFromString("bqnxnel")).length)));
          let _rhs143 = (_this).f23;
          let _rhs144 = ((_module.__default.fm1(new BigNumber(((_667_v19).f31).length), (_667_v19).f32, globalState)) ? (_module.__default.safeDivisionInt((_this).f23, _658_i2)) : (_658_i2));
          let _lhs131 = globalState;
          let _lhs132 = _645_v0;
          let _lhs133 = _module.__default.safeIndex(new BigNumber(661), new BigNumber((_645_v0).length));
          let _lhs134 = globalState;
          let _lhs135 = globalState;
          _lhs131.f8 = _rhs141;
          _lhs132[_lhs133] = _rhs142;
          _lhs134.f14 = _rhs143;
          _lhs135.f14 = _rhs144;
          let _681_v29;
          let _nw107 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Map.Empty);
          _681_v29 = _nw107;
          let _682_v30;
          _682_v30 = _dafny.Map.Empty.slice().updateUnsafe(_655_v8,(_667_v19).f32);
          let _index136 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_681_v29).length));
          (_681_v29)[_index136] = (((_667_v19).f32) ? (_682_v30) : (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("pinimjpwv"),(_this).f28)));
          let _683_v31;
          let _init16 = ((_684_v0) => function (_685_i4) {
            return (_685_i4).multipliedBy((_684_v0)[_module.__default.safeIndex(new BigNumber(661), new BigNumber((_684_v0).length))]);
          })(_645_v0);
          let _nw108 = Array((new BigNumber(25)).toNumber());
          for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw108.length); _i0_16++) {
            _nw108[_i0_16] = _init16(new BigNumber(_i0_16));
          }
          _683_v31 = _nw108;
          let _index137 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_681_v29).length));
          let _rhs145 = (_682_v30).update(_655_v8, !((_667_v19).f32));
          let _rhs146 = _683_v31;
          let _rhs147 = _module.__default.safeDivisionInt(_658_i2, (_dafny.ZERO).minus((_645_v0)[_module.__default.safeIndex(new BigNumber(661), new BigNumber((_645_v0).length))]));
          let _lhs136 = _681_v29;
          let _lhs137 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_681_v29).length));
          let _lhs138 = globalState;
          _lhs136[_lhs137] = _rhs145;
          _683_v31 = _rhs146;
          _lhs138.f14 = _rhs147;
          let _686_v32;
          _686_v32 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(new BigNumber(564), (_this).f28, globalState),_664_v16);
          let _687_v33;
          _687_v33 = _dafny.Map.Empty.slice().updateUnsafe((_667_v19).f32,_664_v16);
          let _688_v34;
          _688_v34 = _dafny.MultiSet.fromElements((_686_v32).Merge(_686_v32), (_687_v33), ((_module.__default.fm1(_658_i2, (_667_v19).f32, globalState)) ? (_686_v32) : (_686_v32)));
          _688_v34 = _688_v34;
        }
      }
      (globalState).f14 = (_this).f23;
      (globalState).f14 = (_this).f23;
      (globalState).f9 = (((_this).f28) ? ((_this).f28) : ((_this).f28));
      return;
    }
    get f30() {
      let _this = this;
      return _this._f30;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f23 = _dafny.ZERO;
      this._f27 = [];
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f23() {
      let _this = this;
      return _this._f23;
    };
    __ctor(f27, f23) {
      let _this = this;
      (_this)._f27 = f27;
      (_this)._f23 = f23;
      return;
    }
    fm10(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((_dafny.Set.fromElements(false)).Difference(_dafny.Set.fromElements(true, !(true), false))).IsSubsetOf((_dafny.Set.fromElements(false, false)).Difference(_dafny.Set.fromElements(true, false)));
    };
    fm11(globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.of(true), ((true) ? (_dafny.Seq.of(false)) : (_dafny.Seq.of(false))));
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Set.Empty;
      let _689_v0;
      let _nw109 = Array((new BigNumber(15)).toNumber()).fill(_module.D1.Default());
      _689_v0 = _nw109;
      let _690_v1;
      _690_v1 = _dafny.Seq.of(_dafny.MultiSet.fromElements(_689_v0));
      let _691_v2;
      _691_v2 = _dafny.Seq.UnicodeFromString("ru");
      let _692_v3;
      _692_v3 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_690_v1)[_module.__default.safeIndex(new BigNumber((_691_v2).length), new BigNumber((_690_v1).length))]);
      let _693_v4;
      _693_v4 = _dafny.MultiSet.fromElements(_689_v0);
      let _694_v5;
      _694_v5 = _module.D3.create_DC9(_693_v4);
      let _rhs148 = new BigNumber(((((_692_v3).contains(p1)) ? ((_692_v3).get(p1)) : ((_694_v5).dtor_cf20))).cardinality());
      let _rhs149 = (_this).f23;
      let _rhs150 = (_this).f23;
      let _rhs151 = _module.__default.safeModuloInt((_this).f23, ((p1) ? (_module.__default.fm0((_this).f23, p1, globalState)) : ((_this).f23)));
      let _rhs152 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("taruujebt"), _691_v2), _691_v2);
      let _lhs139 = globalState;
      let _lhs140 = globalState;
      let _lhs141 = globalState;
      r0 = _rhs148;
      _lhs139.f8 = _rhs149;
      r0 = _rhs150;
      _lhs140.f14 = _rhs151;
      _lhs141.f6 = _rhs152;
      let _695_v6;
      _695_v6 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f23);
      let _696_v7;
      _696_v7 = _695_v6;
      let _697_v8;
      _697_v8 = _dafny.MultiSet.fromElements(_696_v7, _696_v7, _696_v7);
      let _698_v9;
      _698_v9 = _dafny.Seq.of(new BigNumber((_697_v8).cardinality()));
      let _699_v10;
      let _nw110 = new _module.C1();
      _nw110.__ctor(_module.__default.fm1((_this).f23, true, globalState), _dafny.Seq.Concat(_698_v9, _698_v9), (_this).f23);
      _699_v10 = _nw110;
      let _700_v11;
      let _init17 = ((_701_v9) => function (_702_i0) {
        return _701_v9;
      })(_698_v9);
      let _nw111 = Array((new BigNumber(15)).toNumber());
      for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw111.length); _i0_17++) {
        _nw111[_i0_17] = _init17(new BigNumber(_i0_17));
      }
      _700_v11 = _nw111;
      let _index138 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_700_v11).length));
      (_700_v11)[_index138] = _dafny.Seq.Concat(_dafny.Seq.of((_this).f23), _698_v9);
      let _index139 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length));
      (p0)[_index139] = ((_this).f23).isLessThanOrEqualTo((_this).f23);
      let _index140 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_700_v11).length));
      let _index141 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length));
      let _rhs153 = (_module.D9.create_DC27(p1, _698_v9, p0, p1, p1)).dtor_cf45;
      let _rhs154 = p1;
      let _rhs155 = p1;
      let _rhs156 = p1;
      let _rhs157 = (((_this).f23).plus((_this).f23)).isLessThan((_this).f23);
      let _lhs142 = _700_v11;
      let _lhs143 = _module.__default.safeIndex(new BigNumber(122), new BigNumber((_700_v11).length));
      let _lhs144 = globalState;
      let _lhs145 = globalState;
      let _lhs146 = p0;
      let _lhs147 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length));
      let _lhs148 = globalState;
      _lhs142[_lhs143] = _rhs153;
      _lhs144.f7 = _rhs154;
      _lhs145.f21 = _rhs155;
      _lhs146[_lhs147] = _rhs156;
      _lhs148.f21 = _rhs157;
      let _703_i1;
      _703_i1 = _dafny.ZERO;
      L9: {
        while ((p0)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length))]) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_703_i1)) {
              break L9;
            }
            _703_i1 = (_703_i1).plus(_dafny.ONE);
            (globalState).f8 = (_this).f23;
            let _704_v12;
            let _nw112 = Array((new BigNumber(11)).toNumber());
            _704_v12 = _nw112;
            let _705_v13;
            _705_v13 = _module.D3.create_DC10((p0)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length))]);
            let _rhs158 = _704_v12;
            let _rhs159 = (_this).f23;
            let _rhs160 = function (_pat_let10_0) {
              return function (_706_dt__update__tmp_h0) {
                return function (_pat_let11_0) {
                  return function (_707_dt__update_hcf21_h0) {
                    return _module.D3.create_DC10(_707_dt__update_hcf21_h0);
                  }(_pat_let11_0);
                }((new BigNumber(856)).isLessThan(new BigNumber(-625)));
              }(_pat_let10_0);
            }(_module.D3.create_DC10((p0)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length))]));
            let _lhs149 = globalState;
            _704_v12 = _rhs158;
            _lhs149.f8 = _rhs159;
            _705_v13 = _rhs160;
            let _708_v14;
            _708_v14 = _dafny.Set.fromElements((_this).f27);
            let _709_v15;
            _709_v15 = _dafny.Map.Empty.slice().updateUnsafe(_708_v14,p1);
            (globalState).f21 = (((_709_v15).contains(_708_v14)) ? ((_709_v15).get(_708_v14)) : (true));
            let _710_v16;
            let _nw113 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
            _710_v16 = _nw113;
            _710_v16 = _710_v16;
          }
        }
      }
      let _711_v17;
      let _init18 = ((_712_p1, _713_p0) => function (_714_i2) {
        return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(false, _712_p1, (_713_p0)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((_713_p0).length))]),(_this).f23)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(!((_713_p0)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((_713_p0).length))]), (_713_p0)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((_713_p0).length))], true, (_713_p0)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((_713_p0).length))], _712_p1),(_this).f23));
      })(p1, p0);
      let _nw114 = Array((new BigNumber(16)).toNumber());
      for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw114.length); _i0_18++) {
        _nw114[_i0_18] = _init18(new BigNumber(_i0_18));
      }
      _711_v17 = _nw114;
      let _715_v18;
      _715_v18 = _dafny.Set.fromElements(!((p0)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length))]));
      let _index142 = _module.__default.safeIndex(new BigNumber(476), new BigNumber((_711_v17).length));
      (_711_v17)[_index142] = _dafny.Map.Empty.slice().updateUnsafe(_715_v18,(_this).f23);
      let _716_v20;
      _716_v20 = _dafny.Map.Empty.slice().updateUnsafe(_715_v18,_698_v9);
      let _index143 = _module.__default.safeIndex(new BigNumber(476), new BigNumber((_711_v17).length));
      (_711_v17)[_index143] = function () {
        let _coll25 = new _dafny.Map();
        for (const _compr_25 of (_716_v20).Keys.Elements) {
          let _717_v19 = _compr_25;
          if ((_716_v20).contains(_717_v19)) {
            _coll25.push([_717_v19,(_this).f23]);
          }
        }
        return _coll25;
      }();
      if ((p0)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length))]) {
        let _718_v21;
        _718_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,(_this).f23);
        r0 = _module.__default.fm0((((_718_v21).contains(new BigNumber((_691_v2).length))) ? ((_718_v21).get(new BigNumber((_691_v2).length))) : ((_this).f23)), false, globalState);
        (globalState).f9 = (p0)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length))];
        let _719_v22;
        _719_v22 = _dafny.MultiSet.fromElements((p0)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length))], !(p1), !(p1));
        let _720_v23;
        _720_v23 = _dafny.Seq.of(_719_v22);
        let _721_v24;
        let _nw115 = new _module.C3();
        _nw115.__ctor(_720_v23, _module.__default.fm0((_this).f23, p1, globalState), (p0)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length))], (_700_v11)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_700_v11).length))]);
        _721_v24 = _nw115;
        let _722_v25;
        _722_v25 = _dafny.Map.Empty.slice().updateUnsafe(_721_v24,(p0)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length))]);
        let _723_v26;
        let _nw116 = new _module.C3();
        _nw116.__ctor((((p0)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length))]) ? (_720_v23) : (_720_v23)), ((_dafny.ZERO).minus((_this).f23)).minus(new BigNumber((_722_v25).length)), (p0)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length))], (_700_v11)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_700_v11).length))]);
        _723_v26 = _nw116;
        let _724_v29;
        _724_v29 = _dafny.Seq.of(_718_v21, _718_v21);
        let _725_v30;
        _725_v30 = _dafny.Set.fromElements((_this).f23, new BigNumber((_718_v21).length));
        let _726_v31;
        let _nw117 = Array((new BigNumber(16)).toNumber());
        _nw117[(_dafny.ZERO).toNumber()] = _718_v21;
        _nw117[(_dafny.ONE).toNumber()] = _718_v21;
        _nw117[(new BigNumber(2)).toNumber()] = _718_v21;
        _nw117[(new BigNumber(3)).toNumber()] = _718_v21;
        _nw117[(new BigNumber(4)).toNumber()] = ((_module.__default.fm1((((_719_v22).contains((p0)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length))])) ? ((_719_v22).get((p0)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length))])) : (new BigNumber(651))), (p0)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length))], globalState)) ? (_718_v21) : (_718_v21));
        _nw117[(new BigNumber(5)).toNumber()] = _718_v21;
        _nw117[(new BigNumber(6)).toNumber()] = _718_v21;
        _nw117[(new BigNumber(7)).toNumber()] = _718_v21;
        _nw117[(new BigNumber(8)).toNumber()] = _718_v21;
        _nw117[(new BigNumber(9)).toNumber()] = _718_v21;
        _nw117[(new BigNumber(10)).toNumber()] = (_718_v21).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_718_v21).length),(_721_v24).f23));
        _nw117[(new BigNumber(11)).toNumber()] = function () {
          let _coll26 = new _dafny.Map();
          for (const _compr_26 of (function () {
            let _coll27 = new _dafny.Map();
            for (const _compr_27 of _dafny.IntegerRange(new BigNumber(118), new BigNumber(492))) {
              let _727_v28 = _compr_27;
              if (((new BigNumber(118)).isLessThanOrEqualTo(_727_v28)) && ((_727_v28).isLessThan(new BigNumber(492)))) {
                _coll27.push([_module.__default.safeDivisionInt(_727_v28, new BigNumber((_dafny.Set.fromElements(_module.D3.create_DC10(!(p1)))).length)),_dafny.MultiSet.fromElements(new BigNumber((_715_v18).length))]);
              }
            }
            return _coll27;
          }()).Keys.Elements) {
            let _728_v27 = _compr_26;
            if ((function () {
              let _coll28 = new _dafny.Map();
              for (const _compr_28 of _dafny.IntegerRange(new BigNumber(118), new BigNumber(492))) {
                let _727_v28 = _compr_28;
                if (((new BigNumber(118)).isLessThanOrEqualTo(_727_v28)) && ((_727_v28).isLessThan(new BigNumber(492)))) {
                  _coll28.push([_module.__default.safeDivisionInt(_727_v28, new BigNumber((_dafny.Set.fromElements(_module.D3.create_DC10(!(p1)))).length)),_dafny.MultiSet.fromElements(new BigNumber((_715_v18).length))]);
                }
              }
              return _coll28;
            }()).contains(_728_v27)) {
              _coll26.push([(_728_v27).minus((_721_v24).f23),(_721_v24).f23]);
            }
          }
          return _coll26;
        }();
        _nw117[(new BigNumber(12)).toNumber()] = (_724_v29)[_module.__default.safeIndex(new BigNumber((_691_v2).length), new BigNumber((_724_v29).length))];
        _nw117[(new BigNumber(13)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_721_v24).f23),(_721_v24).f23)).update(new BigNumber((_725_v30).length), (_this).f23);
        _nw117[(new BigNumber(14)).toNumber()] = _718_v21;
        _nw117[(new BigNumber(15)).toNumber()] = _718_v21;
        _726_v31 = _nw117;
        let _729_v32;
        _729_v32 = _dafny.Seq.of((_this).fm10(_dafny.Seq.Create(_module.__default.abs(new BigNumber(868)), function (_730_i3) {
          return new _dafny.CodePoint('e'.codePointAt(0));
        }), false, _695_v6, (p0)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length))], globalState), p1, p1, p1, !((p0)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length))]));
        let _index144 = _module.__default.safeIndex(new BigNumber(32), new BigNumber(((_this).f27).length));
        ((_this).f27)[_index144] = (_dafny.ZERO).minus(_module.__default.fm0(((_700_v11)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_700_v11).length))])[_module.__default.safeIndex(new BigNumber((_729_v32).length), new BigNumber(((_700_v11)[_module.__default.safeIndex(new BigNumber(122), new BigNumber((_700_v11).length))]).length))], (p0)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length))], globalState));
        let _731_v33;
        _731_v33 = _module.D11.create_DC30(_726_v31);
        let _index145 = _module.__default.safeIndex(new BigNumber(32), new BigNumber(((_this).f27).length));
        let _rhs161 = (_721_v24).f23;
        let _rhs162 = (_731_v33).dtor_cf51;
        let _rhs163 = (_this).f23;
        let _rhs164 = (_this).f23;
        let _lhs150 = globalState;
        let _lhs151 = (_this).f27;
        let _lhs152 = _module.__default.safeIndex(new BigNumber(32), new BigNumber(((_this).f27).length));
        let _lhs153 = globalState;
        _lhs150.f14 = _rhs161;
        _726_v31 = _rhs162;
        _lhs151[_lhs152] = _rhs163;
        _lhs153.f14 = _rhs164;
        (globalState).f14 = new BigNumber((_698_v9).length);
      } else {
        let _732_v34;
        _732_v34 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0((_this).f23, (p0)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length))], globalState),_module.__default.fm6((_dafny.ZERO).minus((_this).f23), globalState));
        (globalState).f8 = (_dafny.ZERO).minus(new BigNumber((_732_v34).length));
        let _733_v35;
        _733_v35 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,(_this).f23);
        let _734_v36;
        let _nw118 = new _module.C2();
        _nw118.__ctor(_733_v35, (p0)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length))], !(new BigNumber((_691_v2).length)).isEqualTo((_dafny.ZERO).minus((_this).f23)), _dafny.Seq.Concat(_698_v9, _dafny.Seq.Create(_module.__default.abs(new BigNumber(397)), function (_735_i4) {
          return (_this).f23;
        })), (_this).f23);
        _734_v36 = _nw118;
        let _736_v37;
        let _nw119 = Array((new BigNumber(21)).toNumber()).fill(false);
        _736_v37 = _nw119;
        _736_v37 = p0;
        let _737_v38;
        _737_v38 = _dafny.Set.fromElements(_734_v36);
        let _738_v39;
        _738_v39 = _dafny.Seq.of(p1, (p0)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length))], _module.__default.fm1((_this).f23, (_734_v36).f32, globalState), p1, (_737_v38).IsDisjointFrom(_737_v38));
        if ((_738_v39)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f23), new BigNumber((_738_v39).length))]) {
          let _index146 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length));
          (p0)[_index146] = ((_this).f23).isLessThan((_this).f23);
          let _739_v40;
          let _nw120 = Array((new BigNumber(24)).toNumber());
          _nw120[(_dafny.ZERO).toNumber()] = _695_v6;
          _nw120[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_734_v36).f32,(_this).f23);
          _nw120[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((p0)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length))],(_this).f23);
          _nw120[(new BigNumber(3)).toNumber()] = (_695_v6).Merge(_695_v6);
          _nw120[(new BigNumber(4)).toNumber()] = _695_v6;
          _nw120[(new BigNumber(5)).toNumber()] = (_695_v6).update((_734_v36).f32, (_this).f23);
          _nw120[(new BigNumber(6)).toNumber()] = _695_v6;
          _nw120[(new BigNumber(7)).toNumber()] = _695_v6;
          _nw120[(new BigNumber(8)).toNumber()] = (((_734_v36).f32) ? (_695_v6) : (_dafny.Map.Empty.slice().updateUnsafe((_734_v36).f32,(_this).f23)));
          _nw120[(new BigNumber(9)).toNumber()] = _695_v6;
          _nw120[(new BigNumber(10)).toNumber()] = (_695_v6).Merge(_695_v6);
          _nw120[(new BigNumber(11)).toNumber()] = (_695_v6).Merge(_695_v6);
          _nw120[(new BigNumber(12)).toNumber()] = ((_695_v6).update(false, (_this).f23)).Merge(_695_v6);
          _nw120[(new BigNumber(13)).toNumber()] = _695_v6;
          _nw120[(new BigNumber(14)).toNumber()] = (_695_v6).Merge(_695_v6);
          _nw120[(new BigNumber(15)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f23)).update((_734_v36).f32, new BigNumber((_691_v2).length));
          _nw120[(new BigNumber(16)).toNumber()] = _695_v6;
          _nw120[(new BigNumber(17)).toNumber()] = _695_v6;
          _nw120[(new BigNumber(18)).toNumber()] = _695_v6;
          _nw120[(new BigNumber(19)).toNumber()] = (_695_v6).update((p0)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length))], new BigNumber((_695_v6).length));
          _nw120[(new BigNumber(20)).toNumber()] = _695_v6;
          _nw120[(new BigNumber(21)).toNumber()] = _695_v6;
          _nw120[(new BigNumber(22)).toNumber()] = _695_v6;
          _nw120[(new BigNumber(23)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((p0)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length))],(_this).f23);
          _739_v40 = _nw120;
          let _index147 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_739_v40).length));
          (_739_v40)[_index147] = _695_v6;
          let _740_v41;
          _740_v41 = new _dafny.CodePoint('v'.codePointAt(0));
          let _741_v42;
          _741_v42 = _dafny.Map.Empty.slice().updateUnsafe(_740_v41,p1);
          let _742_v43;
          _742_v43 = _module.D12.create_DC34(_741_v42);
          let _index148 = _module.__default.safeIndex(new BigNumber(449), new BigNumber((_739_v40).length));
          (_739_v40)[_index148] = (_dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber(((_742_v43).dtor_cf55).length))).Merge(_695_v6);
          let _index149 = _module.__default.safeIndex(new BigNumber(259), new BigNumber(((_this).f27).length));
          ((_this).f27)[_index149] = (_this).f23;
          let _743_v44;
          _743_v44 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,(p0)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length))]);
          let _744_v45;
          _744_v45 = _module.D8.create_DC21((_734_v36).f32, (p0)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length))], new BigNumber((_743_v44).length));
          let _index150 = _module.__default.safeIndex(new BigNumber(259), new BigNumber(((_this).f27).length));
          ((_this).f27)[_index150] = _module.__default.safeDivisionInt(new BigNumber(798), ((_744_v45).dtor_cf37).multipliedBy((_this).f23));
          let _745_v46;
          let _746_v47;
          let _747_v48;
          let _748_v49;
          let _out30;
          let _out31;
          let _out32;
          let _out33;
          let _outcollector10 = (_734_v36).m6((_dafny.Set.fromElements(true, true, (_734_v36).f32)).IsSubsetOf(_715_v18), globalState);
          _out30 = _outcollector10[0];
          _out31 = _outcollector10[1];
          _out32 = _outcollector10[2];
          _out33 = _outcollector10[3];
          _745_v46 = _out30;
          _746_v47 = _out31;
          _747_v48 = _out32;
          _748_v49 = _out33;
          let _index151 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length));
          (p0)[_index151] = p1;
        } else {
          let _index152 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length));
          (p0)[_index152] = ((new BigNumber(33)).multipliedBy(new BigNumber(127))).isLessThan((_this).f23);
          let _index153 = _module.__default.safeIndex(new BigNumber(914), new BigNumber(((_this).f27).length));
          ((_this).f27)[_index153] = (_this).f23;
          let _index154 = _module.__default.safeIndex(new BigNumber(914), new BigNumber(((_this).f27).length));
          ((_this).f27)[_index154] = _module.__default.safeDivisionInt((_this).f23, _module.__default.safeModuloInt((_this).f23, (_this).f23));
          let _749_v50;
          _749_v50 = new _dafny.CodePoint('n'.codePointAt(0));
          (globalState).f21 = _dafny.Seq.IsPrefixOf(_dafny.Seq.update(_691_v2, _module.__default.safeIndex(((_this).f27)[_module.__default.safeIndex(new BigNumber(914), new BigNumber(((_this).f27).length))], new BigNumber((_691_v2).length)), _749_v50), _dafny.Seq.Concat(_691_v2, _691_v2));
          let _750_v51;
          _750_v51 = _module.D3.create_DC11(new BigNumber(453));
          let _index155 = _module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length));
          (p0)[_index155] = (_738_v39)[_module.__default.safeIndex((_dafny.ZERO).minus((((_734_v36).f32) ? (((_this).f27)[_module.__default.safeIndex(new BigNumber(914), new BigNumber(((_this).f27).length))]) : ((_750_v51).dtor_cf22))), new BigNumber((_738_v39).length))];
          (globalState).f21 = false;
        }
        let _751_v52;
        _751_v52 = _dafny.Set.fromElements((_this).f23);
        let _752_v53;
        let _nw121 = new _module.C2();
        _nw121.__ctor(((_734_v36).f31).update((_this).f23, new BigNumber(797)), false, (_this).fm10(_691_v2, p1, (_695_v6).update((p0)[_module.__default.safeIndex(new BigNumber(428), new BigNumber((p0).length))], (_this).f23), _module.__default.fm1((_dafny.ZERO).minus((_this).f23), (_734_v36).f32, globalState), globalState), _698_v9, _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_751_v52).length)), (_dafny.ZERO).minus((_this).f23)));
        _752_v53 = _nw121;
      }
      r0 = (_this).f23;
      let _753_v54;
      _753_v54 = _dafny.Set.fromElements(new BigNumber((_695_v6).length), (_this).f23);
      r1 = (_753_v54).Union((_753_v54).Intersect(_753_v54));
      return [r0, r1];
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Seq.of();
      (globalState).f8 = (_this).f23;
      let _754_v0;
      _754_v0 = false;
      let _755_i0;
      _755_i0 = _dafny.ZERO;
      L10: {
        while ((((false) ? ((_this).fm10(p0, true, _dafny.Map.Empty.slice().updateUnsafe(_754_v0,(_this).f23), _754_v0, globalState)) : (_754_v0))) || (_754_v0)) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_755_i0)) {
              break L10;
            }
            _755_i0 = (_755_i0).plus(_dafny.ONE);
            let _index156 = _module.__default.safeIndex(new BigNumber(562), new BigNumber(((_this).f27).length));
            ((_this).f27)[_index156] = _module.__default.safeDivisionInt((_this).f23, (_this).f23);
            let _756_v1;
            _756_v1 = _dafny.Seq.of(_754_v0);
            let _757_v2;
            let _nw122 = Array((new BigNumber(3)).toNumber());
            _nw122[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("uovdktobe"), p0)).length);
            _nw122[(_dafny.ONE).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_756_v1).length), (_this).f23);
            _nw122[(new BigNumber(2)).toNumber()] = (_this).f23;
            _757_v2 = _nw122;
            let _758_v3;
            _758_v3 = _dafny.Seq.of((_this).f23);
            let _759_v4;
            _759_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,new BigNumber((_758_v3).length));
            let _760_v5;
            _760_v5 = _dafny.Set.fromElements(_754_v0);
            let _761_v6;
            _761_v6 = _dafny.Map.Empty.slice().updateUnsafe(!(_754_v0),new BigNumber(((_module.D13.create_DC38(_760_v5)).dtor_cf62).length));
            let _762_v7;
            let _nw123 = new _module.C2();
            _nw123.__ctor(_759_v4, (_this).fm10(p0, _754_v0, _761_v6, true, globalState), _754_v0, _758_v3, new BigNumber(465));
            _762_v7 = _nw123;
            let _index157 = _module.__default.safeIndex(new BigNumber(562), new BigNumber(((_this).f27).length));
            let _rhs165 = ((_this).f23).minus((_this).f23);
            let _rhs166 = _754_v0;
            let _rhs167 = (_this).f27;
            let _rhs168 = ((_754_v0) ? (_762_v7) : (_762_v7));
            let _lhs154 = (_this).f27;
            let _lhs155 = _module.__default.safeIndex(new BigNumber(562), new BigNumber(((_this).f27).length));
            let _lhs156 = globalState;
            _lhs154[_lhs155] = _rhs165;
            _lhs156.f21 = _rhs166;
            _757_v2 = _rhs167;
            _762_v7 = _rhs168;
            let _763_v8;
            _763_v8 = new _dafny.CodePoint('n'.codePointAt(0));
            let _764_v9;
            _764_v9 = _dafny.Set.fromElements((_module.D8.create_DC20(_754_v0, ((_this).f27)[_module.__default.safeIndex(new BigNumber(562), new BigNumber(((_this).f27).length))])).dtor_cf34, new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("dxmrtsdjn"), _module.__default.safeIndex(((_this).f27)[_module.__default.safeIndex(new BigNumber(562), new BigNumber(((_this).f27).length))], new BigNumber((_dafny.Seq.UnicodeFromString("dxmrtsdjn")).length)), _763_v8), _module.__default.safeIndex(new BigNumber(512), new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("dxmrtsdjn"), _module.__default.safeIndex(((_this).f27)[_module.__default.safeIndex(new BigNumber(562), new BigNumber(((_this).f27).length))], new BigNumber((_dafny.Seq.UnicodeFromString("dxmrtsdjn")).length)), _763_v8)).length)), _763_v8)).length));
            let _765_v10;
            _765_v10 = _dafny.Map.Empty.slice().updateUnsafe(_754_v0,_754_v0);
            let _766_v11;
            _766_v11 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_764_v9).length),(((_765_v10).contains((_762_v7).f32)) ? ((_765_v10).get((_762_v7).f32)) : (_754_v0)));
            let _767_v12;
            let _nw124 = Array((new BigNumber(12)).toNumber());
            _nw124[(_dafny.ZERO).toNumber()] = true;
            _nw124[(_dafny.ONE).toNumber()] = (_762_v7).f32;
            _nw124[(new BigNumber(2)).toNumber()] = true;
            _nw124[(new BigNumber(3)).toNumber()] = (_762_v7).f32;
            _nw124[(new BigNumber(4)).toNumber()] = _754_v0;
            _nw124[(new BigNumber(5)).toNumber()] = ((_762_v7).f32) || (_754_v0);
            _nw124[(new BigNumber(6)).toNumber()] = !_dafny.areEqual(_756_v1, _756_v1);
            _nw124[(new BigNumber(7)).toNumber()] = _754_v0;
            _nw124[(new BigNumber(8)).toNumber()] = (_754_v0) === (_754_v0);
            _nw124[(new BigNumber(9)).toNumber()] = (((_762_v7).f32) ? ((_762_v7).f32) : ((_762_v7).f32));
            _nw124[(new BigNumber(10)).toNumber()] = false;
            _nw124[(new BigNumber(11)).toNumber()] = (((_this).f27)[_module.__default.safeIndex(new BigNumber(562), new BigNumber(((_this).f27).length))]).isEqualTo(new BigNumber((_dafny.Seq.of(_754_v0, _754_v0, (_762_v7).f32, (((_766_v11).contains((_dafny.ZERO).minus((_this).f23))) ? ((_766_v11).get((_dafny.ZERO).minus((_this).f23))) : ((_762_v7).f32)))).length));
            _767_v12 = _nw124;
            _767_v12 = _767_v12;
            let _768_v13;
            _768_v13 = _module.D9.create_DC27(!(new BigNumber((_764_v9).length)).isEqualTo((_this).f23), _758_v3, _767_v12, _754_v0, (_762_v7).f32);
            let _769_v14;
            _769_v14 = _dafny.Map.Empty.slice().updateUnsafe((_762_v7).f32,new BigNumber(957));
            let _index158 = _module.__default.safeIndex(new BigNumber(562), new BigNumber(((_this).f27).length));
            let _rhs169 = _module.__default.fm0(new BigNumber((p0).length), (((_766_v11).contains(((_this).f27)[_module.__default.safeIndex(new BigNumber(562), new BigNumber(((_this).f27).length))])) ? ((_766_v11).get(((_this).f27)[_module.__default.safeIndex(new BigNumber(562), new BigNumber(((_this).f27).length))])) : ((((_765_v10).contains((_762_v7).f32)) ? ((_765_v10).get((_762_v7).f32)) : ((_this).fm10(p0, _754_v0, (_769_v14), _754_v0, globalState))))), globalState);
            let _rhs170 = _768_v13;
            let _lhs157 = (_this).f27;
            let _lhs158 = _module.__default.safeIndex(new BigNumber(562), new BigNumber(((_this).f27).length));
            _lhs157[_lhs158] = _rhs169;
            _768_v13 = _rhs170;
            let _770_v15;
            _770_v15 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f27)[_module.__default.safeIndex(new BigNumber(562), new BigNumber(((_this).f27).length))],_dafny.Seq.update(_dafny.Seq.Concat(_756_v1, _756_v1), _module.__default.safeIndex(new BigNumber((p0).length), new BigNumber((_dafny.Seq.Concat(_756_v1, _756_v1)).length)), (_762_v7).f32));
            let _rhs171 = ((_this).f23).isEqualTo((new BigNumber((_dafny.MultiSet.fromElements(((_this).f27)[_module.__default.safeIndex(new BigNumber(562), new BigNumber(((_this).f27).length))])).cardinality())).plus(((_this).f27)[_module.__default.safeIndex(new BigNumber(562), new BigNumber(((_this).f27).length))]));
            let _rhs172 = ((_dafny.Map.Empty.slice().updateUnsafe(((_this).f27)[_module.__default.safeIndex(new BigNumber(562), new BigNumber(((_this).f27).length))],_756_v1)).Merge(_770_v15)).Merge((_770_v15).update(new BigNumber((_756_v1).length), (_this).fm11(globalState)));
            let _rhs173 = (_dafny.ZERO).minus((_this).f23);
            let _lhs159 = globalState;
            _lhs159.f9 = _rhs171;
            _770_v15 = _rhs172;
            r0 = _rhs173;
          }
        }
      }
      let _771_v16;
      _771_v16 = _dafny.Set.fromElements(!(_754_v0), !(!(_754_v0)));
      (globalState).f8 = (_dafny.ZERO).minus(function (_source6) {
        if (_source6.is_DC1) {
          let _772___mcc_h0 = (_source6).cf1;
          let _773___mcc_h1 = (_source6).cf2;
          let _774___mcc_h2 = (_source6).cf3;
          let _775_cf3 = _774___mcc_h2;
          let _776_cf2 = _773___mcc_h1;
          let _777_cf1 = _772___mcc_h0;
          return (_775_cf3).minus(_777_cf1);
        } else if (_source6.is_DC0) {
          let _778___mcc_h3 = (_source6).cf0;
          let _779_cf0 = _778___mcc_h3;
          return (new BigNumber(11)).minus((_this).f23);
        } else {
          let _780___mcc_h4 = (_source6).cf4;
          let _781_cf4 = _780___mcc_h4;
          return (_module.D9.create_DC26((_this).f23)).dtor_cf43;
        }
      }(_module.__default.fm26(_771_v16, globalState)));
      let _782_i1;
      _782_i1 = _dafny.ZERO;
      L11: {
        while (!(_754_v0)) {
          C11: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_782_i1)) {
              break L11;
            }
            _782_i1 = (_782_i1).plus(_dafny.ONE);
            if (_754_v0) {
              let _783_v17;
              _783_v17 = new _dafny.CodePoint('s'.codePointAt(0));
              let _784_v18;
              _784_v18 = _dafny.MultiSet.fromElements(_754_v0, _754_v0, false, (_754_v0) || (_module.__default.fm1(new BigNumber((_dafny.Seq.of(_783_v17)).length), _754_v0, globalState)), ((_this).f23).isLessThanOrEqualTo((_this).f23));
              _784_v18 = (_dafny.MultiSet.fromElements(_754_v0, _754_v0)).Difference(_784_v18);
              let _785_v19;
              _785_v19 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f23).isEqualTo((_this).f23),_783_v17);
              let _786_v20;
              _786_v20 = _dafny.Seq.of((_this).f23, new BigNumber(419), (_this).f23);
              let _787_v21;
              _787_v21 = _dafny.Map.Empty.slice().updateUnsafe(_754_v0,_dafny.Seq.Create(_module.__default.abs(new BigNumber(-174)), function (_788_i2) {
                return (_this).f23;
              }));
              let _789_v22;
              _789_v22 = _module.D12.create_DC36(_754_v0, true, (_this).f23, (_this).f23, _787_v21);
              let _790_v23;
              let _nw125 = Array((new BigNumber(26)).toNumber());
              _nw125[(_dafny.ZERO).toNumber()] = ((_this).f23).isLessThanOrEqualTo(new BigNumber(597));
              _nw125[(_dafny.ONE).toNumber()] = _754_v0;
              _nw125[(new BigNumber(2)).toNumber()] = _754_v0;
              _nw125[(new BigNumber(3)).toNumber()] = _754_v0;
              _nw125[(new BigNumber(4)).toNumber()] = !(_754_v0) || (_754_v0);
              _nw125[(new BigNumber(5)).toNumber()] = true;
              _nw125[(new BigNumber(6)).toNumber()] = !(_754_v0) || (_754_v0);
              _nw125[(new BigNumber(7)).toNumber()] = _754_v0;
              _nw125[(new BigNumber(8)).toNumber()] = true;
              _nw125[(new BigNumber(9)).toNumber()] = _754_v0;
              _nw125[(new BigNumber(10)).toNumber()] = ((_dafny.ZERO).minus((_this).f23)).isLessThan((_this).f23);
              _nw125[(new BigNumber(11)).toNumber()] = _754_v0;
              _nw125[(new BigNumber(12)).toNumber()] = !(_754_v0);
              _nw125[(new BigNumber(13)).toNumber()] = _754_v0;
              _nw125[(new BigNumber(14)).toNumber()] = _754_v0;
              _nw125[(new BigNumber(15)).toNumber()] = !_dafny.Seq.contains(p0, _783_v17);
              _nw125[(new BigNumber(16)).toNumber()] = _754_v0;
              _nw125[(new BigNumber(17)).toNumber()] = _754_v0;
              _nw125[(new BigNumber(18)).toNumber()] = !(_dafny.areEqual(_786_v20, _786_v20));
              _nw125[(new BigNumber(19)).toNumber()] = (_754_v0) || (_754_v0);
              _nw125[(new BigNumber(20)).toNumber()] = (_789_v22).dtor_cf57;
              _nw125[(new BigNumber(21)).toNumber()] = _754_v0;
              _nw125[(new BigNumber(22)).toNumber()] = !(true);
              _nw125[(new BigNumber(23)).toNumber()] = _754_v0;
              _nw125[(new BigNumber(24)).toNumber()] = _754_v0;
              _nw125[(new BigNumber(25)).toNumber()] = _754_v0;
              _790_v23 = _nw125;
              let _index159 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_790_v23).length));
              (_790_v23)[_index159] = !(_754_v0);
              let _791_v24;
              let _init19 = function (_792_i3) {
                return new _dafny.CodePoint('o'.codePointAt(0));
              };
              let _nw126 = Array((new BigNumber(6)).toNumber());
              for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw126.length); _i0_19++) {
                _nw126[_i0_19] = _init19(new BigNumber(_i0_19));
              }
              _791_v24 = _nw126;
              let _index160 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_791_v24).length));
              (_791_v24)[_index160] = _783_v17;
              let _index161 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_790_v23).length));
              let _index162 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_791_v24).length));
              let _rhs174 = (_785_v19).Merge(_785_v19);
              let _rhs175 = !(_754_v0) || (!(_754_v0));
              let _rhs176 = _790_v23;
              let _rhs177 = _783_v17;
              let _lhs160 = _790_v23;
              let _lhs161 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_790_v23).length));
              let _lhs162 = _791_v24;
              let _lhs163 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_791_v24).length));
              _785_v19 = _rhs174;
              _lhs160[_lhs161] = _rhs175;
              _790_v23 = _rhs176;
              _lhs162[_lhs163] = _rhs177;
              let _index163 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_790_v23).length));
              (_790_v23)[_index163] = !(_module.__default.fm1((_this).f23, !((_790_v23)[_module.__default.safeIndex(new BigNumber(36), new BigNumber((_790_v23).length))]), globalState)) || (_754_v0);
              let _793_v25;
              _793_v25 = _dafny.Map.Empty.slice().updateUnsafe(_754_v0,false);
              let _794_v26;
              _794_v26 = _dafny.Seq.of(_784_v18, _784_v18, ((_784_v18).update((_790_v23)[_module.__default.safeIndex(new BigNumber(36), new BigNumber((_790_v23).length))], _module.__default.abs(new BigNumber((_793_v25).length)))).update((_790_v23)[_module.__default.safeIndex(new BigNumber(36), new BigNumber((_790_v23).length))], _module.__default.abs(new BigNumber(678))), _784_v18, _784_v18);
              let _795_v27;
              let _nw127 = new _module.C3();
              _nw127.__ctor(_794_v26, (_this).f23, !((_790_v23)[_module.__default.safeIndex(new BigNumber(36), new BigNumber((_790_v23).length))]), _786_v20);
              _795_v27 = _nw127;
              (globalState).f21 = _754_v0;
            } else {
              let _796_v28;
              _796_v28 = new _dafny.CodePoint('x'.codePointAt(0));
              (globalState).f11 = _796_v28;
              let _797_v29;
              let _nw128 = Array((new BigNumber(15)).toNumber()).fill([]);
              _797_v29 = _nw128;
              _797_v29 = _797_v29;
              let _798_v31;
              _798_v31 = _dafny.Map.Empty.slice().updateUnsafe(_754_v0,(_this).f23);
              let _799_v32;
              _799_v32 = _dafny.Set.fromElements(_module.__default.fm0((_this).f23, _754_v0, globalState), (_this).f23, (_this).f23, _module.__default.fm0((_dafny.ZERO).minus((_this).f23), (_this).fm10(_dafny.Seq.Create(_module.__default.abs(new BigNumber(532)), ((_800_v28) => function (_801_i4) {
                return _800_v28;
              })(_796_v28)), _754_v0, _798_v31, _754_v0, globalState), globalState), (_this).f23);
              let _802_v33;
              let _nw129 = Array((new BigNumber(2)).toNumber());
              _nw129[(_dafny.ZERO).toNumber()] = (function () {
                let _coll29 = new _dafny.Set();
                for (const _compr_29 of _dafny.IntegerRange(new BigNumber(397), new BigNumber(-231))) {
                  let _803_v30 = _compr_29;
                  if (((new BigNumber(397)).isLessThanOrEqualTo(_803_v30)) && ((_803_v30).isLessThan(new BigNumber(-231)))) {
                    _coll29.add((_803_v30).minus((_this).f23));
                  }
                }
                return _coll29;
              }()).Union(_dafny.Set.fromElements((_this).f23));
              _nw129[(_dafny.ONE).toNumber()] = _799_v32;
              _802_v33 = _nw129;
              let _index164 = _module.__default.safeIndex(new BigNumber(754), new BigNumber((_802_v33).length));
              (_802_v33)[_index164] = _799_v32;
              let _index165 = _module.__default.safeIndex(new BigNumber(754), new BigNumber((_802_v33).length));
              (_802_v33)[_index165] = _799_v32;
              let _804_v34;
              _804_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(!(_754_v0))).length),_754_v0);
              _804_v34 = (_804_v34).Merge(_804_v34);
              let _805_v35;
              _805_v35 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,(_this).f23);
              _805_v35 = (_805_v35).update(_module.__default.fm0(new BigNumber(-241), !(_754_v0), globalState), (_this).f23);
            }
            let _806_v36;
            _806_v36 = _dafny.Seq.of((_this).f23);
            let _807_v37;
            _807_v37 = _module.D8.create_DC20(_754_v0, (_this).f23);
            let _808_v38;
            _808_v38 = _dafny.Seq.of(_754_v0, _754_v0, (_807_v37).dtor_cf33, _754_v0, _754_v0);
            let _809_v39;
            let _nw130 = Array((new BigNumber(18)).toNumber());
            _nw130[(_dafny.ZERO).toNumber()] = _module.__default.fm1((_this).f23, true, globalState);
            _nw130[(_dafny.ONE).toNumber()] = _754_v0;
            _nw130[(new BigNumber(2)).toNumber()] = (_771_v16).IsSubsetOf(_771_v16);
            _nw130[(new BigNumber(3)).toNumber()] = _754_v0;
            _nw130[(new BigNumber(4)).toNumber()] = !(_754_v0);
            _nw130[(new BigNumber(5)).toNumber()] = ((_this).f23).isLessThanOrEqualTo(new BigNumber((_806_v36).length));
            _nw130[(new BigNumber(6)).toNumber()] = _754_v0;
            _nw130[(new BigNumber(7)).toNumber()] = (_808_v38)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("hm")).length), new BigNumber((_808_v38).length))];
            _nw130[(new BigNumber(8)).toNumber()] = _754_v0;
            _nw130[(new BigNumber(9)).toNumber()] = _754_v0;
            _nw130[(new BigNumber(10)).toNumber()] = _754_v0;
            _nw130[(new BigNumber(11)).toNumber()] = _754_v0;
            _nw130[(new BigNumber(12)).toNumber()] = false;
            _nw130[(new BigNumber(13)).toNumber()] = _754_v0;
            _nw130[(new BigNumber(14)).toNumber()] = _module.__default.fm1(new BigNumber((p0).length), _754_v0, globalState);
            _nw130[(new BigNumber(15)).toNumber()] = !(!(!(_754_v0)));
            _nw130[(new BigNumber(16)).toNumber()] = !(_754_v0) || (_754_v0);
            _nw130[(new BigNumber(17)).toNumber()] = _754_v0;
            _809_v39 = _nw130;
            let _index166 = _module.__default.safeIndex(new BigNumber(808), new BigNumber((_809_v39).length));
            (_809_v39)[_index166] = _754_v0;
            let _index167 = _module.__default.safeIndex(new BigNumber(808), new BigNumber((_809_v39).length));
            (_809_v39)[_index167] = !(_754_v0);
            if (_754_v0) {
              let _810_v40;
              _810_v40 = _dafny.Map.Empty.slice().updateUnsafe(_806_v36,_754_v0);
              _810_v40 = ((_810_v40).Merge(_810_v40)).Merge(_810_v40);
              let _811_v41;
              _811_v41 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1((_this).f23, (_809_v39)[_module.__default.safeIndex(new BigNumber(808), new BigNumber((_809_v39).length))], globalState),((_module.D0.create_DC1(new BigNumber(967), (_809_v39)[_module.__default.safeIndex(new BigNumber(808), new BigNumber((_809_v39).length))], (_this).f23)).dtor_cf3).isLessThanOrEqualTo(new BigNumber(872)));
              (globalState).f7 = (((_811_v41).contains(((_dafny.ZERO).minus((_this).f23)).isLessThanOrEqualTo((_this).f23))) ? ((_811_v41).get(((_dafny.ZERO).minus((_this).f23)).isLessThanOrEqualTo((_this).f23))) : (((_this).f23).isLessThanOrEqualTo(new BigNumber(341))));
              (globalState).f14 = _module.__default.fm0((_this).f23, (new BigNumber(80)).isLessThanOrEqualTo((_this).f23), globalState);
              let _812_v42;
              _812_v42 = _dafny.Map.Empty.slice().updateUnsafe((_809_v39)[_module.__default.safeIndex(new BigNumber(808), new BigNumber((_809_v39).length))],(_this).f23);
              let _813_v43;
              _813_v43 = (_812_v42).Merge(_812_v42);
              _813_v43 = _813_v43;
              (globalState).f21 = !(true) || (_754_v0);
            } else {
              let _814_v44;
              let _815_v45;
              let _816_v46;
              let _out34;
              let _out35;
              let _out36;
              let _outcollector11 = _module.__default.m0(_809_v39, p0, (_this).f23, globalState);
              _out34 = _outcollector11[0];
              _out35 = _outcollector11[1];
              _out36 = _outcollector11[2];
              _814_v44 = _out34;
              _815_v45 = _out35;
              _816_v46 = _out36;
              let _817_v47;
              let _nw131 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.of());
              _817_v47 = _nw131;
              let _818_v48;
              let _nw132 = new _module.C1();
              _nw132.__ctor((_809_v39)[_module.__default.safeIndex(new BigNumber(808), new BigNumber((_809_v39).length))], _806_v36, (_this).f23);
              _818_v48 = _nw132;
              let _819_v49;
              _819_v49 = _dafny.Map.Empty.slice().updateUnsafe(_818_v48,_817_v47);
              let _820_v50;
              let _nw133 = Array((new BigNumber(23)).toNumber());
              _nw133[(_dafny.ZERO).toNumber()] = _817_v47;
              _nw133[(_dafny.ONE).toNumber()] = _817_v47;
              _nw133[(new BigNumber(2)).toNumber()] = _817_v47;
              _nw133[(new BigNumber(3)).toNumber()] = _817_v47;
              _nw133[(new BigNumber(4)).toNumber()] = _817_v47;
              _nw133[(new BigNumber(5)).toNumber()] = (((_809_v39)[_module.__default.safeIndex(new BigNumber(808), new BigNumber((_809_v39).length))]) ? (_817_v47) : ((((_819_v49).contains(_818_v48)) ? ((_819_v49).get(_818_v48)) : (_817_v47))));
              _nw133[(new BigNumber(6)).toNumber()] = _817_v47;
              _nw133[(new BigNumber(7)).toNumber()] = _817_v47;
              _nw133[(new BigNumber(8)).toNumber()] = _817_v47;
              _nw133[(new BigNumber(9)).toNumber()] = _817_v47;
              _nw133[(new BigNumber(10)).toNumber()] = _817_v47;
              _nw133[(new BigNumber(11)).toNumber()] = _817_v47;
              _nw133[(new BigNumber(12)).toNumber()] = _817_v47;
              _nw133[(new BigNumber(13)).toNumber()] = _817_v47;
              _nw133[(new BigNumber(14)).toNumber()] = _817_v47;
              _nw133[(new BigNumber(15)).toNumber()] = _817_v47;
              _nw133[(new BigNumber(16)).toNumber()] = _817_v47;
              _nw133[(new BigNumber(17)).toNumber()] = _817_v47;
              _nw133[(new BigNumber(18)).toNumber()] = _817_v47;
              _nw133[(new BigNumber(19)).toNumber()] = _817_v47;
              _nw133[(new BigNumber(20)).toNumber()] = _817_v47;
              _nw133[(new BigNumber(21)).toNumber()] = _817_v47;
              _nw133[(new BigNumber(22)).toNumber()] = _817_v47;
              _820_v50 = _nw133;
              let _index168 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_820_v50).length));
              (_820_v50)[_index168] = _817_v47;
              let _821_v51;
              _821_v51 = new _dafny.CodePoint('r'.codePointAt(0));
              let _822_v52;
              _822_v52 = _dafny.MultiSet.fromElements(_821_v51, new _dafny.CodePoint('o'.codePointAt(0)), _821_v51);
              let _index169 = _module.__default.safeIndex(new BigNumber(136), new BigNumber(((_this).f27).length));
              ((_this).f27)[_index169] = (((_822_v52).contains(_821_v51)) ? ((_822_v52).get(_821_v51)) : (new BigNumber((_dafny.Seq.of((_809_v39)[_module.__default.safeIndex(new BigNumber(808), new BigNumber((_809_v39).length))])).length)));
              let _index170 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_820_v50).length));
              let _index171 = _module.__default.safeIndex(new BigNumber(136), new BigNumber(((_this).f27).length));
              let _rhs178 = _817_v47;
              let _rhs179 = ((_dafny.ZERO).minus(_816_v46)).plus((((_809_v39)[_module.__default.safeIndex(new BigNumber(808), new BigNumber((_809_v39).length))]) ? ((_this).f23) : ((_this).f23)));
              let _lhs164 = _820_v50;
              let _lhs165 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_820_v50).length));
              let _lhs166 = (_this).f27;
              let _lhs167 = _module.__default.safeIndex(new BigNumber(136), new BigNumber(((_this).f27).length));
              _lhs164[_lhs165] = _rhs178;
              _lhs166[_lhs167] = _rhs179;
              let _823_v53;
              _823_v53 = _dafny.MultiSet.fromElements((_this).f23);
              let _824_v54;
              _824_v54 = _dafny.Seq.of(p0, _dafny.Seq.UnicodeFromString("sxaeugrwk"));
              let _825_v55;
              _825_v55 = _dafny.Map.Empty.slice().updateUnsafe(((((_823_v53).contains(new BigNumber((_dafny.MultiSet.fromElements(_816_v46, ((_this).f27)[_module.__default.safeIndex(new BigNumber(136), new BigNumber(((_this).f27).length))])).cardinality()))) ? ((_823_v53).get(new BigNumber((_dafny.MultiSet.fromElements(_816_v46, ((_this).f27)[_module.__default.safeIndex(new BigNumber(136), new BigNumber(((_this).f27).length))])).cardinality()))) : ((_this).f23))).minus(((_this).f27)[_module.__default.safeIndex(new BigNumber(136), new BigNumber(((_this).f27).length))]),new BigNumber(((_824_v54)[_module.__default.safeIndex(((_this).f27)[_module.__default.safeIndex(new BigNumber(136), new BigNumber(((_this).f27).length))], new BigNumber((_824_v54).length))]).length));
              _825_v55 = _825_v55;
              let _826_v56;
              let _init20 = ((_827_v45, _828_v0) => function (_829_i5) {
                return _dafny.Map.Empty.slice().updateUnsafe(_827_v45,_828_v0);
              })(_815_v45, _754_v0);
              let _nw134 = Array((new BigNumber(25)).toNumber());
              for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw134.length); _i0_20++) {
                _nw134[_i0_20] = _init20(new BigNumber(_i0_20));
              }
              _826_v56 = _nw134;
              let _830_v57;
              _830_v57 = _dafny.Map.Empty.slice().updateUnsafe(_815_v45,false);
              let _index172 = _module.__default.safeIndex(new BigNumber(452), new BigNumber((_826_v56).length));
              (_826_v56)[_index172] = _830_v57;
              let _831_v58;
              _831_v58 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(975)), function (_832_i6) {
                return _dafny.Seq.UnicodeFromString("o");
              }));
              let _833_v59;
              _833_v59 = _dafny.Seq.of(_module.__default.fm21(_dafny.Seq.update((_831_v58)[_module.__default.safeIndex(_816_v46, new BigNumber((_831_v58).length))], _module.__default.safeIndex(_816_v46, new BigNumber(((_831_v58)[_module.__default.safeIndex(_816_v46, new BigNumber((_831_v58).length))]).length)), p0), globalState));
              let _index173 = _module.__default.safeIndex(new BigNumber(452), new BigNumber((_826_v56).length));
              let _rhs180 = _dafny.Map.Empty.slice().updateUnsafe(_815_v45,_754_v0);
              let _rhs181 = false;
              let _rhs182 = _833_v59;
              let _rhs183 = (_this).f23;
              let _lhs168 = _826_v56;
              let _lhs169 = _module.__default.safeIndex(new BigNumber(452), new BigNumber((_826_v56).length));
              let _lhs170 = globalState;
              let _lhs171 = globalState;
              _lhs168[_lhs169] = _rhs180;
              _lhs170.f21 = _rhs181;
              _833_v59 = _rhs182;
              _lhs171.f8 = _rhs183;
              (globalState).f14 = _module.__default.fm0((_module.D0.create_DC1(_816_v46, !((_809_v39)[_module.__default.safeIndex(new BigNumber(808), new BigNumber((_809_v39).length))]), _816_v46)).dtor_cf3, !((((_this).f27)[_module.__default.safeIndex(new BigNumber(136), new BigNumber(((_this).f27).length))]).isLessThanOrEqualTo(_816_v46)), globalState);
            }
            if (true) {
              let _834_v60;
              let _nw135 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Map.Empty);
              _834_v60 = _nw135;
              let _835_v61;
              _835_v61 = _module.D11.create_DC33(_module.D11.create_DC30(_834_v60));
              let _836_v62;
              _836_v62 = _module.D11.create_DC30(_834_v60);
              let _837_v63;
              let _nw136 = Array((new BigNumber(18)).toNumber());
              _nw136[(_dafny.ZERO).toNumber()] = _835_v61;
              _nw136[(_dafny.ONE).toNumber()] = _835_v61;
              _nw136[(new BigNumber(2)).toNumber()] = _835_v61;
              _nw136[(new BigNumber(3)).toNumber()] = _835_v61;
              _nw136[(new BigNumber(4)).toNumber()] = _835_v61;
              _nw136[(new BigNumber(5)).toNumber()] = _835_v61;
              _nw136[(new BigNumber(6)).toNumber()] = _835_v61;
              _nw136[(new BigNumber(7)).toNumber()] = _835_v61;
              _nw136[(new BigNumber(8)).toNumber()] = _module.D11.create_DC33(_836_v62);
              _nw136[(new BigNumber(9)).toNumber()] = _835_v61;
              _nw136[(new BigNumber(10)).toNumber()] = _835_v61;
              _nw136[(new BigNumber(11)).toNumber()] = _835_v61;
              _nw136[(new BigNumber(12)).toNumber()] = _835_v61;
              _nw136[(new BigNumber(13)).toNumber()] = _835_v61;
              _nw136[(new BigNumber(14)).toNumber()] = _835_v61;
              _nw136[(new BigNumber(15)).toNumber()] = _835_v61;
              _nw136[(new BigNumber(16)).toNumber()] = _835_v61;
              _nw136[(new BigNumber(17)).toNumber()] = _835_v61;
              _837_v63 = _nw136;
              let _838_v64;
              _838_v64 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_806_v36).length), (_dafny.ZERO).minus((_806_v36)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_806_v36).length))]))),_837_v63);
              _838_v64 = (_838_v64).update(new BigNumber(885), _837_v63);
              let _839_v65;
              let _nw137 = new _module.C1();
              _nw137.__ctor((_809_v39)[_module.__default.safeIndex(new BigNumber(808), new BigNumber((_809_v39).length))], _806_v36, (_this).f23);
              _839_v65 = _nw137;
              let _840_v66;
              _840_v66 = _dafny.Seq.of(_839_v65);
              let _841_v67;
              _841_v67 = _dafny.Map.Empty.slice().updateUnsafe((_809_v39)[_module.__default.safeIndex(new BigNumber(808), new BigNumber((_809_v39).length))],_module.__default.fm1(new BigNumber((_840_v66).length), _module.__default.fm1(new BigNumber((_771_v16).length), (_809_v39)[_module.__default.safeIndex(new BigNumber(808), new BigNumber((_809_v39).length))], globalState), globalState));
              let _842_v68;
              _842_v68 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_841_v67).length),_771_v16);
              let _843_v69;
              _843_v69 = _dafny.Map.Empty.slice().updateUnsafe(((_842_v68).update((_this).f23, _771_v16)).Merge(_842_v68),(_this).f23);
              (globalState).f14 = (((_843_v69).contains((((_809_v39)[_module.__default.safeIndex(new BigNumber(808), new BigNumber((_809_v39).length))]) ? (_842_v68) : (_dafny.Map.Empty.slice().updateUnsafe((_this).f23,_771_v16))))) ? ((_843_v69).get((((_809_v39)[_module.__default.safeIndex(new BigNumber(808), new BigNumber((_809_v39).length))]) ? (_842_v68) : (_dafny.Map.Empty.slice().updateUnsafe((_this).f23,_771_v16))))) : ((_this).f23));
              let _844_v70;
              let _nw138 = Array((new BigNumber(10)).toNumber());
              _844_v70 = _nw138;
              let _845_v71;
              _845_v71 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,(_this).f23);
              let _846_v72;
              _846_v72 = _dafny.Map.Empty.slice().updateUnsafe((_809_v39)[_module.__default.safeIndex(new BigNumber(808), new BigNumber((_809_v39).length))],new BigNumber((_dafny.Seq.UnicodeFromString("eys")).length));
              let _847_v73;
              _847_v73 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-251)), function (_848_i7) {
                return new BigNumber(-445);
              })).length));
              let _849_v74;
              let _nw139 = new _module.C2();
              _nw139.__ctor(_845_v71, _754_v0, (_809_v39)[_module.__default.safeIndex(new BigNumber(808), new BigNumber((_809_v39).length))], _806_v36, (_dafny.ZERO).minus((((_846_v72).contains((_this).fm10(_dafny.Seq.UnicodeFromString("bms"), _754_v0, _846_v72, true, globalState))) ? ((_846_v72).get((_this).fm10(_dafny.Seq.UnicodeFromString("bms"), _754_v0, _846_v72, true, globalState))) : (new BigNumber((_847_v73).cardinality())))));
              _849_v74 = _nw139;
              let _index174 = _module.__default.safeIndex(new BigNumber(156), new BigNumber((_844_v70).length));
              (_844_v70)[_index174] = _849_v74;
              let _850_v75;
              _850_v75 = new _dafny.CodePoint('o'.codePointAt(0));
              let _851_v76;
              _851_v76 = _dafny.Map.Empty.slice().updateUnsafe(_849_v74,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(_754_v0),_850_v75)).length));
              let _index175 = _module.__default.safeIndex(new BigNumber(156), new BigNumber((_844_v70).length));
              let _rhs184 = !(_754_v0);
              let _rhs185 = (_849_v74).f32;
              let _rhs186 = _module.__default.safeDivisionInt((_this).f23, (((_851_v76).contains(_849_v74)) ? ((_851_v76).get(_849_v74)) : ((_this).f23)));
              let _rhs187 = _849_v74;
              let _lhs172 = globalState;
              let _lhs173 = globalState;
              let _lhs174 = globalState;
              let _lhs175 = _844_v70;
              let _lhs176 = _module.__default.safeIndex(new BigNumber(156), new BigNumber((_844_v70).length));
              _lhs172.f9 = _rhs184;
              _lhs173.f21 = _rhs185;
              _lhs174.f14 = _rhs186;
              _lhs175[_lhs176] = _rhs187;
              let _index176 = _module.__default.safeIndex(new BigNumber(808), new BigNumber((_809_v39).length));
              (_809_v39)[_index176] = _754_v0;
              let _852_v77;
              _852_v77 = _dafny.Map.Empty.slice().updateUnsafe(true,_839_v65);
              _852_v77 = _852_v77;
            } else {
              let _853_v78;
              _853_v78 = _dafny.Map.Empty.slice().updateUnsafe((((_809_v39)[_module.__default.safeIndex(new BigNumber(808), new BigNumber((_809_v39).length))]) ? ((_809_v39)[_module.__default.safeIndex(new BigNumber(808), new BigNumber((_809_v39).length))]) : ((_809_v39)[_module.__default.safeIndex(new BigNumber(808), new BigNumber((_809_v39).length))])),(_this).f23);
              let _854_v79;
              _854_v79 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt((_this).f23, (_this).f23),(_this).f23);
              let _rhs188 = _754_v0;
              let _rhs189 = (((_854_v79).contains((_this).f23)) ? ((_854_v79).get((_this).f23)) : ((_this).f23));
              let _rhs190 = new BigNumber(((_853_v78).Merge((_853_v78).update(_754_v0, new BigNumber(437)))).length);
              let _rhs191 = (_853_v78).Merge(_853_v78);
              let _lhs177 = globalState;
              let _lhs178 = globalState;
              _lhs177.f7 = _rhs188;
              r0 = _rhs189;
              _lhs178.f8 = _rhs190;
              _853_v78 = _rhs191;
              let _855_v80;
              _855_v80 = _dafny.Map.Empty.slice().updateUnsafe((_809_v39)[_module.__default.safeIndex(new BigNumber(808), new BigNumber((_809_v39).length))],p0);
              let _856_v81;
              _856_v81 = _module.D11.create_DC32(p0);
              let _857_v82;
              let _nw140 = Array((new BigNumber(11)).toNumber());
              _nw140[(_dafny.ZERO).toNumber()] = _855_v80;
              _nw140[(_dafny.ONE).toNumber()] = _855_v80;
              _nw140[(new BigNumber(2)).toNumber()] = _855_v80;
              _nw140[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(!(_module.__default.fm1(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_809_v39)[_module.__default.safeIndex(new BigNumber(808), new BigNumber((_809_v39).length))],_754_v0)).update((_809_v39)[_module.__default.safeIndex(new BigNumber(808), new BigNumber((_809_v39).length))], _754_v0)).length), _754_v0, globalState)),_dafny.Seq.UnicodeFromString("lqs"));
              _nw140[(new BigNumber(4)).toNumber()] = _855_v80;
              _nw140[(new BigNumber(5)).toNumber()] = _855_v80;
              _nw140[(new BigNumber(6)).toNumber()] = _module.__default.fm27(globalState);
              _nw140[(new BigNumber(7)).toNumber()] = _855_v80;
              _nw140[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_754_v0,(_856_v81).dtor_cf53);
              _nw140[(new BigNumber(9)).toNumber()] = (_855_v80).Merge(_855_v80);
              _nw140[(new BigNumber(10)).toNumber()] = _855_v80;
              _857_v82 = _nw140;
              let _index177 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_857_v82).length));
              (_857_v82)[_index177] = _855_v80;
              let _858_v83;
              _858_v83 = _dafny.Map.Empty.slice().updateUnsafe(_808_v38,(_855_v80).Merge((_855_v80).update(_754_v0, p0)));
              let _index178 = _module.__default.safeIndex(new BigNumber(933), new BigNumber((_857_v82).length));
              (_857_v82)[_index178] = (((_858_v83).contains(_dafny.Seq.Concat(_808_v38, _808_v38))) ? ((_858_v83).get(_dafny.Seq.Concat(_808_v38, _808_v38))) : (_module.__default.fm27(globalState)));
              let _859_v84;
              _859_v84 = _module.D3.create_DC12(_808_v38);
              let _860_v85;
              let _init21 = ((_861_p0) => function (_862_i8) {
                return _dafny.Seq.of(_861_p0, _861_p0, _861_p0);
              })(p0);
              let _nw141 = Array((new BigNumber(6)).toNumber());
              for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw141.length); _i0_21++) {
                _nw141[_i0_21] = _init21(new BigNumber(_i0_21));
              }
              _860_v85 = _nw141;
              let _863_v86;
              _863_v86 = _dafny.Seq.of(_module.__default.fm2(globalState));
              let _864_v87;
              _864_v87 = _863_v86;
              let _index179 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_860_v85).length));
              (_860_v85)[_index179] = _864_v87;
              let _index180 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_860_v85).length));
              let _rhs192 = _859_v84;
              let _rhs193 = _864_v87;
              let _rhs194 = (_this).f23;
              let _lhs179 = _860_v85;
              let _lhs180 = _module.__default.safeIndex(new BigNumber(292), new BigNumber((_860_v85).length));
              let _lhs181 = globalState;
              _859_v84 = _rhs192;
              _lhs179[_lhs180] = _rhs193;
              _lhs181.f8 = _rhs194;
              (globalState).f14 = (_this).f23;
              let _865_v88;
              _865_v88 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,(_809_v39)[_module.__default.safeIndex(new BigNumber(808), new BigNumber((_809_v39).length))]);
              _854_v79 = (_854_v79).update(new BigNumber((_865_v88).length), new BigNumber(933));
            }
          }
        }
      }
      let _866_v89;
      _866_v89 = _module.D13.create_DC38(_771_v16);
      let _867_v90;
      _867_v90 = _module.D13.create_DC40(_866_v89);
      let _pat_let_tv35 = _866_v89;
      let _pat_let_tv36 = _866_v89;
      let _source7 = function (_pat_let12_0) {
        return function (_870_dt__update__tmp_h1) {
          return function (_pat_let15_0) {
            return function (_871_dt__update_hcf64_h1) {
              return _module.D13.create_DC40(_871_dt__update_hcf64_h1);
            }(_pat_let15_0);
          }(_pat_let_tv36);
        }(_pat_let12_0);
      }(function (_pat_let13_0) {
        return function (_868_dt__update__tmp_h0) {
          return function (_pat_let14_0) {
            return function (_869_dt__update_hcf64_h0) {
              return _module.D13.create_DC40(_869_dt__update_hcf64_h0);
            }(_pat_let14_0);
          }(_module.D13.create_DC40(_pat_let_tv35));
        }(_pat_let13_0);
      }(_867_v90));
      if (_source7.is_DC39) {
        let _872___mcc_h5 = (_source7).cf63;
        let _873_cf63 = _872___mcc_h5;
        let _874_v91;
        let _nw142 = Array((new BigNumber(22)).toNumber()).fill(false);
        _874_v91 = _nw142;
        let _875_v92;
        _875_v92 = _dafny.Set.fromElements(_874_v91);
        let _876_v93;
        _876_v93 = _dafny.Seq.of(_874_v91);
        let _877_v94;
        _877_v94 = _dafny.Seq.of(_875_v92, _dafny.Set.fromElements(_874_v91));
        let _878_v95;
        let _nw143 = Array((new BigNumber(19)).toNumber());
        _nw143[(_dafny.ZERO).toNumber()] = ((_754_v0) ? (_875_v92) : (_875_v92));
        _nw143[(_dafny.ONE).toNumber()] = _875_v92;
        _nw143[(new BigNumber(2)).toNumber()] = _875_v92;
        _nw143[(new BigNumber(3)).toNumber()] = _dafny.Set.fromElements(_874_v91, (_876_v93)[_module.__default.safeIndex(_873_cf63, new BigNumber((_876_v93).length))]);
        _nw143[(new BigNumber(4)).toNumber()] = _875_v92;
        _nw143[(new BigNumber(5)).toNumber()] = _875_v92;
        _nw143[(new BigNumber(6)).toNumber()] = (_875_v92).Union(_875_v92);
        _nw143[(new BigNumber(7)).toNumber()] = _875_v92;
        _nw143[(new BigNumber(8)).toNumber()] = _875_v92;
        _nw143[(new BigNumber(9)).toNumber()] = _875_v92;
        _nw143[(new BigNumber(10)).toNumber()] = _875_v92;
        _nw143[(new BigNumber(11)).toNumber()] = (_877_v94)[_module.__default.safeIndex(_873_cf63, new BigNumber((_877_v94).length))];
        _nw143[(new BigNumber(12)).toNumber()] = _875_v92;
        _nw143[(new BigNumber(13)).toNumber()] = _875_v92;
        _nw143[(new BigNumber(14)).toNumber()] = _875_v92;
        _nw143[(new BigNumber(15)).toNumber()] = _875_v92;
        _nw143[(new BigNumber(16)).toNumber()] = _875_v92;
        _nw143[(new BigNumber(17)).toNumber()] = _875_v92;
        _nw143[(new BigNumber(18)).toNumber()] = _875_v92;
        _878_v95 = _nw143;
        let _index181 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_878_v95).length));
        (_878_v95)[_index181] = _875_v92;
        let _879_v96;
        let _nw144 = Array((new BigNumber(23)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _879_v96 = _nw144;
        let _880_v97;
        _880_v97 = _dafny.Seq.of(_module.D9.create_DC24(_879_v96));
        let _881_v98;
        _881_v98 = _dafny.Seq.of(new BigNumber((_880_v97).length));
        let _882_v99;
        _882_v99 = _module.D9.create_DC27(_754_v0, _881_v98, _874_v91, false, true);
        let _index182 = _module.__default.safeIndex(new BigNumber(446), new BigNumber((_878_v95).length));
        (_878_v95)[_index182] = _dafny.Set.fromElements((_882_v99).dtor_cf46, _874_v91);
        let _index183 = _module.__default.safeIndex(new BigNumber(197), new BigNumber(((_this).f27).length));
        ((_this).f27)[_index183] = (_this).f23;
        let _883_v100;
        let _nw145 = new _module.C1();
        _nw145.__ctor(!(_754_v0), _881_v98, (_this).f23);
        _883_v100 = _nw145;
        let _884_v101;
        _884_v101 = _dafny.Map.Empty.slice().updateUnsafe(_754_v0,_883_v100);
        let _885_v102;
        _885_v102 = _dafny.Seq.of(_754_v0);
        let _index184 = _module.__default.safeIndex(new BigNumber(197), new BigNumber(((_this).f27).length));
        ((_this).f27)[_index184] = _module.__default.fm0(new BigNumber((_884_v101).length), (_885_v102)[_module.__default.safeIndex((_this).f23, new BigNumber((_885_v102).length))], globalState);
        let _886_v103;
        _886_v103 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(((_this).f27)[_module.__default.safeIndex(new BigNumber(197), new BigNumber(((_this).f27).length))], _754_v0, globalState),_754_v0);
        let _887_v104;
        _887_v104 = _dafny.MultiSet.fromElements((_this).f23, new BigNumber((_886_v103).length), (_this).f23);
        if (((new BigNumber((_887_v104).cardinality())).plus(((_this).f27)[_module.__default.safeIndex(new BigNumber(197), new BigNumber(((_this).f27).length))])).isLessThan(((_this).f27)[_module.__default.safeIndex(new BigNumber(197), new BigNumber(((_this).f27).length))])) {
          _885_v102 = _dafny.Seq.Concat(_dafny.Seq.update(_885_v102, _module.__default.safeIndex(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(878)), ((_888_v102) => function (_889_i9) {
            return new BigNumber((_888_v102).length);
          })(_885_v102))).length), new BigNumber((_885_v102).length)), _754_v0), _dafny.Seq.Concat(_dafny.Seq.of(_754_v0), _885_v102));
          let _index185 = _module.__default.safeIndex(new BigNumber(197), new BigNumber(((_this).f27).length));
          ((_this).f27)[_index185] = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(p0, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("chtuw"), p0)), _module.__default.safeIndex(_module.__default.safeModuloInt(((_this).f27)[_module.__default.safeIndex(new BigNumber(197), new BigNumber(((_this).f27).length))], ((_this).f27)[_module.__default.safeIndex(new BigNumber(197), new BigNumber(((_this).f27).length))]), new BigNumber((_dafny.Seq.Concat(p0, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("chtuw"), p0))).length)), (p0)[_module.__default.safeIndex(((_this).f27)[_module.__default.safeIndex(new BigNumber(197), new BigNumber(((_this).f27).length))], new BigNumber((p0).length))])).length);
          let _890_v105;
          _890_v105 = _module.D3.create_DC11((((_887_v104).contains(_873_cf63)) ? ((_887_v104).get(_873_cf63)) : (((_this).f27)[_module.__default.safeIndex(new BigNumber(197), new BigNumber(((_this).f27).length))])));
          _890_v105 = _890_v105;
          let _891_v106;
          _891_v106 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f27)[_module.__default.safeIndex(new BigNumber(197), new BigNumber(((_this).f27).length))],(_this).f23);
          _891_v106 = (function () {
            let _coll30 = new _dafny.Map();
            for (const _compr_30 of _dafny.IntegerRange(new BigNumber(311), new BigNumber(520))) {
              let _892_v107 = _compr_30;
              if (((new BigNumber(311)).isLessThanOrEqualTo(_892_v107)) && ((_892_v107).isLessThan(new BigNumber(520)))) {
                _coll30.push([(_892_v107).plus((_this).f23),_873_cf63]);
              }
            }
            return _coll30;
          }()).Merge(_891_v106);
          let _893_v108;
          let _nw146 = new _module.C0();
          _nw146.__ctor();
          _893_v108 = _nw146;
        } else {
          let _894_v109;
          _894_v109 = new _dafny.CodePoint('n'.codePointAt(0));
          let _895_v110;
          _895_v110 = _dafny.MultiSet.fromElements(_894_v109);
          let _896_v112;
          _896_v112 = _module.D9.create_DC25(_754_v0, function () {
  let _coll31 = new _dafny.Set();
  for (const _compr_31 of (_895_v110).Elements) {
    let _897_v111 = _compr_31;
    if ((_895_v110).contains(_897_v111)) {
      _coll31.add(_897_v111);
    }
  }
  return _coll31;
}());
          let _898_v113;
          _898_v113 = _module.D9.create_DC28(_896_v112);
          _898_v113 = _898_v113;
          let _899_v114;
          _899_v114 = _dafny.Map.Empty.slice().updateUnsafe(_754_v0,((_this).f27)[_module.__default.safeIndex(new BigNumber(197), new BigNumber(((_this).f27).length))]);
          let _900_v115;
          _900_v115 = _module.D9.create_DC26((_this).f23);
          let _rhs195 = p0;
          let _rhs196 = (_module.__default.safeDivisionInt((_this).f23, _873_cf63)).isLessThanOrEqualTo((((_this).f27)[_module.__default.safeIndex(new BigNumber(197), new BigNumber(((_this).f27).length))]).plus(new BigNumber((_899_v114).length)));
          let _rhs197 = _754_v0;
          let _rhs198 = _754_v0;
          let _rhs199 = _module.__default.safeModuloInt(new BigNumber(-409), (_900_v115).dtor_cf43);
          let _lhs182 = globalState;
          let _lhs183 = globalState;
          let _lhs184 = globalState;
          let _lhs185 = globalState;
          let _lhs186 = globalState;
          _lhs182.f6 = _rhs195;
          _lhs183.f9 = _rhs196;
          _lhs184.f9 = _rhs197;
          _lhs185.f9 = _rhs198;
          _lhs186.f8 = _rhs199;
          _754_v0 = !((_this).fm10(_dafny.Seq.UnicodeFromString("vrkt"), _754_v0, _dafny.Map.Empty.slice().updateUnsafe(_754_v0,((_this).f27)[_module.__default.safeIndex(new BigNumber(197), new BigNumber(((_this).f27).length))]), _754_v0, globalState)) || (_754_v0);
          let _index186 = _module.__default.safeIndex(new BigNumber(197), new BigNumber(((_this).f27).length));
          ((_this).f27)[_index186] = ((_this).f27)[_module.__default.safeIndex(new BigNumber(197), new BigNumber(((_this).f27).length))];
          let _901_v116;
          _901_v116 = _dafny.MultiSet.fromElements(_754_v0);
          let _902_v117;
          _902_v117 = _dafny.Seq.of(_901_v116, _901_v116);
          let _903_v118;
          let _nw147 = new _module.C3();
          _nw147.__ctor(_dafny.Seq.Concat(_902_v117, _902_v117), _873_cf63, _module.__default.fm1(_873_cf63, _754_v0, globalState), _881_v98);
          _903_v118 = _nw147;
        }
        let _904_v119;
        _904_v119 = _module.D9.create_DC26((_this).f23);
        let _905_v120;
        _905_v120 = _module.D9.create_DC28(_904_v119);
        let _906_v121;
        _906_v121 = _dafny.Seq.of(((_754_v0) ? (_905_v120) : (_905_v120)));
        let _907_v122;
        _907_v122 = _module.D14.create_DC41(_906_v121);
        _906_v121 = _dafny.Seq.Concat(_906_v121, (_907_v122).dtor_cf65);
      } else if (_source7.is_DC38) {
        let _908___mcc_h6 = (_source7).cf62;
        let _909_cf62 = _908___mcc_h6;
        let _910_v123;
        let _init22 = ((_911_v0) => function (_912_i10) {
          return _911_v0;
        })(_754_v0);
        let _nw148 = Array((new BigNumber(21)).toNumber());
        for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw148.length); _i0_22++) {
          _nw148[_i0_22] = _init22(new BigNumber(_i0_22));
        }
        _910_v123 = _nw148;
        let _913_v124;
        let _914_v125;
        let _915_v126;
        let _out37;
        let _out38;
        let _out39;
        let _outcollector12 = _module.__default.m0(_910_v123, _dafny.Seq.Create(_module.__default.abs(new BigNumber(779)), function (_916_i11) {
          return new _dafny.CodePoint('k'.codePointAt(0));
        }), new BigNumber((_dafny.Seq.UnicodeFromString("lwigmupbn")).length), globalState);
        _out37 = _outcollector12[0];
        _out38 = _outcollector12[1];
        _out39 = _outcollector12[2];
        _913_v124 = _out37;
        _914_v125 = _out38;
        _915_v126 = _out39;
        if (_754_v0) {
          let _917_v127;
          _917_v127 = new _dafny.CodePoint('v'.codePointAt(0));
          (globalState).f11 = _917_v127;
          let _918_v128;
          _918_v128 = _dafny.Seq.of(p0, _dafny.Seq.UnicodeFromString("fq"));
          let _919_v129;
          _919_v129 = _918_v128;
          _918_v128 = _dafny.Seq.Concat(_918_v128, _dafny.Seq.Concat((_919_v129), _918_v128));
          let _920_v130;
          _920_v130 = _dafny.Map.Empty.slice().updateUnsafe(((_754_v0) ? (_910_v123) : (_910_v123)),((_754_v0) ? (_754_v0) : (_754_v0)));
          _920_v130 = (_920_v130).update(_910_v123, !_dafny.areEqual(p0, p0));
          let _921_v131;
          let _nw149 = new _module.C1();
          _nw149.__ctor(_754_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(771)), function (_922_i12) {
            return new BigNumber(-205);
          }), new BigNumber((_dafny.Seq.UnicodeFromString("vfnpyft")).length));
          _921_v131 = _nw149;
          _921_v131 = _921_v131;
          let _index187 = _module.__default.safeIndex(new BigNumber(913), new BigNumber(((_this).f27).length));
          ((_this).f27)[_index187] = (_this).f23;
          let _index188 = _module.__default.safeIndex(new BigNumber(913), new BigNumber(((_this).f27).length));
          ((_this).f27)[_index188] = (_this).f23;
        } else {
          let _923_v132;
          _923_v132 = _dafny.Map.Empty.slice().updateUnsafe(_915_v126,_module.__default.fm6(new BigNumber(962), globalState));
          let _924_v133;
          _924_v133 = new _dafny.CodePoint('j'.codePointAt(0));
          let _925_v134;
          _925_v134 = _dafny.Map.Empty.slice().updateUnsafe(_924_v133,!(_754_v0));
          let _926_v135;
          _926_v135 = _dafny.Seq.of((_this).f23);
          let _927_v136;
          _927_v136 = _dafny.Seq.of(new BigNumber((_925_v134).length), (_this).f23, (_926_v135)[_module.__default.safeIndex(_915_v126, new BigNumber((_926_v135).length))]);
          let _rhs200 = (((_923_v132).contains((_this).f23)) ? ((_923_v132).get((_this).f23)) : (_924_v133));
          let _rhs201 = _dafny.Seq.Concat(_926_v135, _dafny.Seq.Create(_module.__default.abs(new BigNumber(336)), function (_928_i13) {
            return (_this).f23;
          }));
          let _lhs187 = globalState;
          _lhs187.f11 = _rhs200;
          r1 = _rhs201;
          let _929_v137;
          _929_v137 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f23,new BigNumber(396)));
          let _930_v138;
          _930_v138 = _dafny.Map.Empty.slice().updateUnsafe(_929_v137,_dafny.Seq.Create(_module.__default.abs(new BigNumber(-831)), ((_931_v133) => function (_932_i14) {
            return _931_v133;
          })(_924_v133)));
          _930_v138 = (_930_v138).update(_dafny.Seq.Concat(_929_v137, _929_v137), _dafny.Seq.Concat(p0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(271)), ((_933_v133) => function (_934_i15) {
            return _933_v133;
          })(_924_v133))));
          let _index189 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_910_v123).length));
          (_910_v123)[_index189] = !(((_dafny.ZERO).minus(_915_v126)).isLessThanOrEqualTo(new BigNumber(495)));
          let _index190 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_910_v123).length));
          (_910_v123)[_index190] = !(_754_v0);
          let _935_v139;
          _935_v139 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('w'.codePointAt(0)), _924_v133, _924_v133);
          let _rhs202 = p0;
          let _rhs203 = _module.__default.fm1((_926_v135)[_module.__default.safeIndex((((_935_v139).contains(_924_v133)) ? ((_935_v139).get(_924_v133)) : (new BigNumber((_dafny.Seq.UnicodeFromString("ew")).length))), new BigNumber((_926_v135).length))], _module.__default.fm1(new BigNumber((function () {
            let _coll32 = new _dafny.Set();
            for (const _compr_32 of (_925_v134).Keys.Elements) {
              let _936_v140 = _compr_32;
              if ((_925_v134).contains(_936_v140)) {
                _coll32.add(_936_v140);
              }
            }
            return _coll32;
          }()).length), true, globalState), globalState);
          let _rhs204 = _915_v126;
          let _lhs188 = globalState;
          let _lhs189 = globalState;
          let _lhs190 = globalState;
          _lhs188.f6 = _rhs202;
          _lhs189.f7 = _rhs203;
          _lhs190.f8 = _rhs204;
          let _937_v141;
          _937_v141 = _dafny.Seq.of((_915_v126).isEqualTo(new BigNumber(-100)));
          let _index191 = _module.__default.safeIndex(new BigNumber(403), new BigNumber((_910_v123).length));
          (_910_v123)[_index191] = (_937_v141)[_module.__default.safeIndex(new BigNumber((p0).length), new BigNumber((_937_v141).length))];
        }
        if (_754_v0) {
          let _938_v142;
          _938_v142 = _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("ufjwpiq")).length));
          (globalState).f11 = _module.__default.fm6((_dafny.ZERO).minus((_938_v142)[_module.__default.safeIndex((_this).f23, new BigNumber((_938_v142).length))]), globalState);
          let _939_v143;
          let _nw150 = new _module.C1();
          _nw150.__ctor(!(_module.__default.fm1((_this).f23, !(_754_v0), globalState)) || (_754_v0), _938_v142, (_this).f23);
          _939_v143 = _nw150;
          (globalState).f8 = _915_v126;
          let _940_v144;
          _940_v144 = _dafny.MultiSet.fromElements(!(_754_v0), !(_754_v0) || (_754_v0), _754_v0, _754_v0);
          _940_v144 = _940_v144;
          (globalState).f7 = !(_754_v0);
        } else {
          let _941_v145;
          _941_v145 = _dafny.Seq.of(_915_v126);
          (globalState).f8 = ((_941_v145)[_module.__default.safeIndex(_915_v126, new BigNumber((_941_v145).length))]).minus((_this).f23);
          let _index192 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_910_v123).length));
          (_910_v123)[_index192] = _754_v0;
          let _942_v146;
          _942_v146 = _dafny.Seq.of(_754_v0, _754_v0);
          let _943_v147;
          _943_v147 = _dafny.Map.Empty.slice().updateUnsafe(_942_v146,new BigNumber(267));
          let _index193 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_910_v123).length));
          (_910_v123)[_index193] = !(((true) ? (_module.__default.fm28(_915_v126, globalState)) : (_943_v147))).equals(_943_v147);
          let _944_v148;
          _944_v148 = new _dafny.CodePoint('e'.codePointAt(0));
          let _945_v149;
          _945_v149 = _module.D4.create_DC15((_this).f23, (p0)[_module.__default.safeIndex((_this).f23, new BigNumber((p0).length))], (_this).f23);
          let _946_v150;
          _946_v150 = _dafny.Set.fromElements(_944_v148);
          let _947_v151;
          _947_v151 = _dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC7(_946_v150),_944_v148);
          let _948_v152;
          _948_v152 = _module.D2.create_DC7(_946_v150);
          let _949_v153;
          _949_v153 = _module.D2.create_DC8(_915_v126, new BigNumber((_914_v125).length), _944_v148, (_dafny.ZERO).minus((_this).f23), (((_914_v125).contains(_915_v126)) ? ((_914_v125).get(_915_v126)) : ((_910_v123)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_910_v123).length))])));
          let _950_v154;
          let _nw151 = Array((new BigNumber(19)).toNumber());
          _nw151[(_dafny.ZERO).toNumber()] = _944_v148;
          _nw151[(_dafny.ONE).toNumber()] = _944_v148;
          _nw151[(new BigNumber(2)).toNumber()] = _944_v148;
          _nw151[(new BigNumber(3)).toNumber()] = _module.__default.fm6(new BigNumber(531), globalState);
          _nw151[(new BigNumber(4)).toNumber()] = _944_v148;
          _nw151[(new BigNumber(5)).toNumber()] = (_945_v149).dtor_cf27;
          _nw151[(new BigNumber(6)).toNumber()] = _944_v148;
          _nw151[(new BigNumber(7)).toNumber()] = new _dafny.CodePoint('y'.codePointAt(0));
          _nw151[(new BigNumber(8)).toNumber()] = _944_v148;
          _nw151[(new BigNumber(9)).toNumber()] = _944_v148;
          _nw151[(new BigNumber(10)).toNumber()] = _944_v148;
          _nw151[(new BigNumber(11)).toNumber()] = _944_v148;
          _nw151[(new BigNumber(12)).toNumber()] = _module.__default.fm6(new BigNumber((_942_v146).length), globalState);
          _nw151[(new BigNumber(13)).toNumber()] = (((_947_v151).contains(_948_v152)) ? ((_947_v151).get(_948_v152)) : (_944_v148));
          _nw151[(new BigNumber(14)).toNumber()] = _944_v148;
          _nw151[(new BigNumber(15)).toNumber()] = new _dafny.CodePoint('u'.codePointAt(0));
          _nw151[(new BigNumber(16)).toNumber()] = _944_v148;
          _nw151[(new BigNumber(17)).toNumber()] = (_949_v153).dtor_cf17;
          _nw151[(new BigNumber(18)).toNumber()] = new _dafny.CodePoint('y'.codePointAt(0));
          _950_v154 = _nw151;
          _950_v154 = _950_v154;
          let _951_v155;
          _951_v155 = _dafny.Map.Empty.slice().updateUnsafe(((_910_v123)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_910_v123).length))]) === (_754_v0),p0);
          _951_v155 = (_951_v155).update((_910_v123)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_910_v123).length))], p0);
          let _952_v157;
          let _nw152 = new _module.C3();
          _nw152.__ctor(_module.__default.fm29(_915_v126, function () {
            let _coll33 = new _dafny.Set();
            for (const _compr_33 of _dafny.IntegerRange(new BigNumber(899), new BigNumber(229))) {
              let _953_v156 = _compr_33;
              if (((new BigNumber(899)).isLessThanOrEqualTo(_953_v156)) && ((_953_v156).isLessThan(new BigNumber(229)))) {
                _coll33.add((_953_v156).minus((_this).f23));
              }
            }
            return _coll33;
          }(), globalState), (_this).f23, (_910_v123)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_910_v123).length))], _dafny.Seq.Concat(_941_v145, _941_v145));
          _952_v157 = _nw152;
          _952_v157 = _952_v157;
        }
        _754_v0 = !(!(false)) || (_754_v0);
      } else {
        let _954___mcc_h7 = (_source7).cf64;
        let _955_cf64 = _954___mcc_h7;
        let _956_v158;
        let _nw153 = Array((new BigNumber(25)).toNumber()).fill(_module.D1.Default());
        _956_v158 = _nw153;
        let _957_v159;
        _957_v159 = _dafny.MultiSet.fromElements(_956_v158);
        let _958_v160;
        _958_v160 = _dafny.Seq.of(_module.D3.create_DC9(_957_v159));
        let _959_v161;
        _959_v161 = _dafny.MultiSet.fromElements(new BigNumber(393));
        let _960_v162;
        _960_v162 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements((_this).f23),_module.__default.fm1((_this).f23, _754_v0, globalState));
        let _961_v163;
        _961_v163 = _dafny.Seq.of(!(_754_v0));
        let _962_v164;
        _962_v164 = _dafny.Set.fromElements((_dafny.ZERO).minus((_this).f23));
        let _963_v165;
        _963_v165 = new _dafny.CodePoint('p'.codePointAt(0));
        let _964_v166;
        _964_v166 = _dafny.Map.Empty.slice().updateUnsafe(_963_v165,(_this).f23);
        let _965_v167;
        _965_v167 = _dafny.MultiSet.fromElements(_754_v0, false);
        let _966_v168;
        _966_v168 = _dafny.Map.Empty.slice().updateUnsafe(_754_v0,_754_v0);
        let _967_v169;
        _967_v169 = _dafny.Map.Empty.slice().updateUnsafe(_754_v0,new BigNumber((_966_v168).length));
        let _968_v170;
        let _nw154 = Array((new BigNumber(27)).toNumber());
        _nw154[(_dafny.ZERO).toNumber()] = _754_v0;
        _nw154[(_dafny.ONE).toNumber()] = _dafny.Seq.IsProperPrefixOf(_958_v160, _958_v160);
        _nw154[(new BigNumber(2)).toNumber()] = (_959_v161).IsSubsetOf((_959_v161).update((_this).f23, _module.__default.abs((_this).f23)));
        _nw154[(new BigNumber(3)).toNumber()] = _754_v0;
        _nw154[(new BigNumber(4)).toNumber()] = (((_960_v162).contains(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_961_v163).length))))) ? ((_960_v162).get(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_961_v163).length))))) : (_754_v0));
        _nw154[(new BigNumber(5)).toNumber()] = _754_v0;
        _nw154[(new BigNumber(6)).toNumber()] = (_962_v164).IsProperSubsetOf(_962_v164);
        _nw154[(new BigNumber(7)).toNumber()] = !(new BigNumber(-888)).isEqualTo((_dafny.ZERO).minus((_this).f23));
        _nw154[(new BigNumber(8)).toNumber()] = true;
        _nw154[(new BigNumber(9)).toNumber()] = ((_dafny.ZERO).minus((((_964_v166).contains(_963_v165)) ? ((_964_v166).get(_963_v165)) : (new BigNumber((_959_v161).cardinality()))))).isLessThan((_this).f23);
        _nw154[(new BigNumber(10)).toNumber()] = ((_this).f23).isLessThan((_this).f23);
        _nw154[(new BigNumber(11)).toNumber()] = _754_v0;
        _nw154[(new BigNumber(12)).toNumber()] = !((_754_v0) === (_754_v0));
        _nw154[(new BigNumber(13)).toNumber()] = ((_this).f23).isLessThanOrEqualTo((_this).f23);
        _nw154[(new BigNumber(14)).toNumber()] = _754_v0;
        _nw154[(new BigNumber(15)).toNumber()] = !(_754_v0);
        _nw154[(new BigNumber(16)).toNumber()] = _754_v0;
        _nw154[(new BigNumber(17)).toNumber()] = _754_v0;
        _nw154[(new BigNumber(18)).toNumber()] = _754_v0;
        _nw154[(new BigNumber(19)).toNumber()] = !(_754_v0);
        _nw154[(new BigNumber(20)).toNumber()] = !(new BigNumber((_965_v167).cardinality())).isEqualTo((_this).f23);
        _nw154[(new BigNumber(21)).toNumber()] = (_this).fm10(p0, _754_v0, _967_v169, _754_v0, globalState);
        _nw154[(new BigNumber(22)).toNumber()] = (_this).fm10(p0, _754_v0, _967_v169, _754_v0, globalState);
        _nw154[(new BigNumber(23)).toNumber()] = _754_v0;
        _nw154[(new BigNumber(24)).toNumber()] = _754_v0;
        _nw154[(new BigNumber(25)).toNumber()] = _754_v0;
        _nw154[(new BigNumber(26)).toNumber()] = _754_v0;
        _968_v170 = _nw154;
        _968_v170 = _968_v170;
        let _969_v171;
        let _init23 = function (_970_i16) {
          return (_970_i16).multipliedBy((_this).f23);
        };
        let _nw155 = Array((new BigNumber(6)).toNumber());
        for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw155.length); _i0_23++) {
          _nw155[_i0_23] = _init23(new BigNumber(_i0_23));
        }
        _969_v171 = _nw155;
        let _971_v172;
        _971_v172 = _dafny.Map.Empty.slice().updateUnsafe(_754_v0,(_this).f23);
        let _rhs205 = _module.__default.fm0((_dafny.ZERO).minus((_this).f23), true, globalState);
        let _rhs206 = _969_v171;
        let _rhs207 = _967_v169;
        let _rhs208 = (_this).f23;
        r0 = _rhs205;
        _969_v171 = _rhs206;
        _971_v172 = _rhs207;
        r0 = _rhs208;
        let _972_v173;
        let _nw156 = Array((new BigNumber(16)).toNumber()).fill([]);
        _972_v173 = _nw156;
        _972_v173 = _972_v173;
        let _index194 = _module.__default.safeIndex(new BigNumber(618), new BigNumber((_968_v170).length));
        (_968_v170)[_index194] = ((!(_754_v0)) ? (true) : (_754_v0));
        let _index195 = _module.__default.safeIndex(new BigNumber(618), new BigNumber((_968_v170).length));
        (_968_v170)[_index195] = _754_v0;
      }
      let _973_v174;
      _973_v174 = _module.D8.create_DC22(_754_v0);
      let _pat_let_tv37 = _754_v0;
      (globalState).f7 = (function (_pat_let16_0) {
        return function (_974_dt__update__tmp_h3) {
          return function (_pat_let17_0) {
            return function (_975_dt__update_hcf38_h0) {
              return _module.D8.create_DC22(_975_dt__update_hcf38_h0);
            }(_pat_let17_0);
          }(_pat_let_tv37);
        }(_pat_let16_0);
      }(_973_v174)).dtor_cf38;
      r0 = (_this).f23;
      r1 = _dafny.Seq.of(new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus((_this).f23), (_this).f23)).length), ((_this).f23).multipliedBy((_this).f23), (_this).f23);
      return [r0, r1];
    }
    get f27() {
      let _this = this;
      return _this._f27;
    };
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f23 = _dafny.ZERO;
      this._f25 = _dafny.ZERO;
      this._f26 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f23() {
      let _this = this;
      return _this._f23;
    };
    __ctor(f25, f26, f23) {
      let _this = this;
      (_this)._f25 = f25;
      (_this)._f26 = f26;
      (_this)._f23 = f23;
      return;
    }
    fm8(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Map.Empty.slice().updateUnsafe((function () {
        let _coll34 = new _dafny.Map();
        for (const _compr_34 of (function () {
          let _coll35 = new _dafny.Map();
          for (const _compr_35 of (_dafny.MultiSet.fromElements((_this).f26, (_this).f26)).Elements) {
            let _976_v1 = _compr_35;
            if ((_dafny.MultiSet.fromElements((_this).f26, (_this).f26)).contains(_976_v1)) {
              _coll35.push([(_976_v1).minus(new BigNumber((function () {
                let _coll36 = new _dafny.Map();
                for (const _compr_36 of ((_module.D2.create_DC7(function () {
  let _coll37 = new _dafny.Set();
  for (const _compr_37 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(547)), function (_977_i0) {
    return new _dafny.CodePoint('u'.codePointAt(0));
  })).Elements) {
    let _978_v3 = _compr_37;
    if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(547)), function (_977_i0) {
      return new _dafny.CodePoint('u'.codePointAt(0));
    }), _978_v3)) {
      _coll37.add(_978_v3);
    }
  }
  return _coll37;
}())).dtor_cf14).Elements) {
                  let _979_v2 = _compr_36;
                  if (((_module.D2.create_DC7(function () {
  let _coll38 = new _dafny.Set();
  for (const _compr_38 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(547)), function (_977_i0) {
    return new _dafny.CodePoint('u'.codePointAt(0));
  })).Elements) {
    let _980_v3 = _compr_38;
    if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(547)), function (_977_i0) {
      return new _dafny.CodePoint('u'.codePointAt(0));
    }), _980_v3)) {
      _coll38.add(_980_v3);
    }
  }
  return _coll38;
}())).dtor_cf14).contains(_979_v2)) {
                    _coll36.push([_979_v2,true]);
                  }
                }
                return _coll36;
              }()).length)),false]);
            }
          }
          return _coll35;
        }()).Keys.Elements) {
          let _981_v0 = _compr_34;
          if ((function () {
            let _coll39 = new _dafny.Map();
            for (const _compr_39 of (_dafny.MultiSet.fromElements((_this).f26, (_this).f26)).Elements) {
              let _976_v1 = _compr_39;
              if ((_dafny.MultiSet.fromElements((_this).f26, (_this).f26)).contains(_976_v1)) {
                _coll39.push([(_976_v1).minus(new BigNumber((function () {
                  let _coll40 = new _dafny.Map();
                  for (const _compr_40 of ((_module.D2.create_DC7(function () {
  let _coll41 = new _dafny.Set();
  for (const _compr_41 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(547)), function (_977_i0) {
    return new _dafny.CodePoint('u'.codePointAt(0));
  })).Elements) {
    let _982_v3 = _compr_41;
    if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(547)), function (_977_i0) {
      return new _dafny.CodePoint('u'.codePointAt(0));
    }), _982_v3)) {
      _coll41.add(_982_v3);
    }
  }
  return _coll41;
}())).dtor_cf14).Elements) {
                    let _979_v2 = _compr_40;
                    if (((_module.D2.create_DC7(function () {
  let _coll42 = new _dafny.Set();
  for (const _compr_42 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(547)), function (_977_i0) {
    return new _dafny.CodePoint('u'.codePointAt(0));
  })).Elements) {
    let _983_v3 = _compr_42;
    if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(547)), function (_977_i0) {
      return new _dafny.CodePoint('u'.codePointAt(0));
    }), _983_v3)) {
      _coll42.add(_983_v3);
    }
  }
  return _coll42;
}())).dtor_cf14).contains(_979_v2)) {
                      _coll40.push([_979_v2,true]);
                    }
                  }
                  return _coll40;
                }()).length)),false]);
              }
            }
            return _coll39;
          }()).contains(_981_v0)) {
            _coll34.push([_module.__default.safeModuloInt(_981_v0, (_this).f25),false]);
          }
        }
        return _coll34;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("vdal")).length),false)),(true) && (true));
    };
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Set.Empty;
      let _984_v0;
      _984_v0 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber(-317));
      let _rhs209 = p1;
      let _rhs210 = new BigNumber((_module.__default.fm9(p1, (_dafny.ZERO).minus((_this).f25), (_dafny.Map.Empty.slice().updateUnsafe(p1,(_this).f26)).Merge(_984_v0), globalState)).length);
      let _lhs191 = globalState;
      _lhs191.f21 = _rhs209;
      r0 = _rhs210;
      let _985_v1;
      _985_v1 = _dafny.Seq.of((_this).f23);
      let _986_v2;
      let _nw157 = new _module.C1();
      _nw157.__ctor((_module.__default.fm1((_this).f26, p1, globalState)) || (p1), _dafny.Seq.Concat(_985_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(12)), function (_987_i0) {
        return new BigNumber(-818);
      })), (_dafny.ZERO).minus((_this).f26));
      _986_v2 = _nw157;
      let _988_v3;
      _988_v3 = new _dafny.CodePoint('m'.codePointAt(0));
      let _989_v4;
      _989_v4 = _dafny.Map.Empty.slice().updateUnsafe(_988_v3,p0);
      let _990_v5;
      _990_v5 = _module.D2.create_DC8(new BigNumber(340), new BigNumber(-216), _988_v3, (_this).f26, p1);
      _989_v4 = (_989_v4).update((_990_v5).dtor_cf17, p0);
      let _991_v6;
      _991_v6 = _module.D1.create_DC4(_module.__default.fm16(p1, (_this).f23, _module.__default.fm0((_this).f26, false, globalState), globalState), p1, p1);
      let _992_i1;
      _992_i1 = _dafny.ZERO;
      L12: {
        while ((_991_v6).dtor_cf7) {
          C12: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_992_i1)) {
              break L12;
            }
            _992_i1 = (_992_i1).plus(_dafny.ONE);
            (globalState).f21 = p1;
            let _993_v7;
            let _nw158 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
            _993_v7 = _nw158;
            let _994_v8;
            _994_v8 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,(_this).f23);
            let _index196 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_993_v7).length));
            (_993_v7)[_index196] = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_this).f23), new BigNumber((_994_v8).length));
            let _index197 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_993_v7).length));
            (_993_v7)[_index197] = (_dafny.ZERO).minus(((((_this).f25).isLessThan((_this).f26)) ? ((_this).f23) : ((_this).f25)));
            let _index198 = _module.__default.safeIndex(new BigNumber(281), new BigNumber((p0).length));
            (p0)[_index198] = p1;
            let _index199 = _module.__default.safeIndex(new BigNumber(281), new BigNumber((p0).length));
            (p0)[_index199] = ((_this).f26).isLessThanOrEqualTo((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.fm0((_this).f25, p1, globalState)))));
            let _995_v9;
            _995_v9 = _dafny.Seq.UnicodeFromString("jyx");
            let _996_v10;
            _996_v10 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(762),(p0)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((p0).length))]);
            let _997_v11;
            _997_v11 = _dafny.Map.Empty.slice().updateUnsafe(_995_v9,_996_v10);
            let _index200 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_993_v7).length));
            let _rhs211 = ((_997_v11).update(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("gxiuegtec"), _module.__default.safeIndex(new BigNumber(-119), new BigNumber((_dafny.Seq.UnicodeFromString("gxiuegtec")).length)), _988_v3), _996_v10)).update(_module.__default.fm2(globalState), _996_v10);
            let _rhs212 = (_dafny.ZERO).minus(new BigNumber((_995_v9).length));
            let _lhs192 = _993_v7;
            let _lhs193 = _module.__default.safeIndex(new BigNumber(507), new BigNumber((_993_v7).length));
            _997_v11 = _rhs211;
            _lhs192[_lhs193] = _rhs212;
          }
        }
      }
      _984_v0 = (_984_v0).update(true, (_this).f26);
      let _hi11 = (_this).f25;
      for (let _998_i2 = new BigNumber((_985_v1).length); _998_i2.isLessThan(_hi11); _998_i2 = _998_i2.plus(_dafny.ONE)) {
        let _999_v12;
        _999_v12 = _dafny.Seq.of(((_dafny.ZERO).minus(new BigNumber(-840))).isLessThan((_this).f23));
        let _1000_v13;
        _1000_v13 = _module.D8.create_DC19(_985_v1);
        let _1001_v14;
        _1001_v14 = _dafny.Seq.UnicodeFromString("nxqooq");
        let _rhs213 = !(!(new BigNumber(4)).isEqualTo(((_dafny.ZERO).minus((_this).f23)).multipliedBy((_this).f25)));
        let _rhs214 = !(_module.__default.safeModuloInt((_this).f26, (_this).f23)).isEqualTo(_module.__default.safeDivisionInt((_this).f26, (_this).f26));
        let _rhs215 = _dafny.Seq.Concat(_dafny.Seq.Concat(_module.__default.fm30((_this).f23, _1000_v13, _1001_v14, (_this).f25, globalState), _999_v12), _dafny.Seq.Concat(_999_v12, _999_v12));
        let _rhs216 = ((_this).f26).isLessThan((_this).f23);
        let _lhs194 = globalState;
        let _lhs195 = globalState;
        let _lhs196 = globalState;
        _lhs194.f7 = _rhs213;
        _lhs195.f21 = _rhs214;
        _999_v12 = _rhs215;
        _lhs196.f21 = _rhs216;
        let _1002_v15;
        _1002_v15 = _module.D3.create_DC10(p1);
        let _1003_v16;
        _1003_v16 = _dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC10(p1),_module.__default.fm0((_this).f25, p1, globalState));
        if ((_1003_v16).contains(_1002_v15)) {
          let _1004_v17;
          _1004_v17 = _dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.update(_985_v1, _module.__default.safeIndex(_998_i2, new BigNumber((_985_v1).length)), (_this).f26));
          let _1005_v18;
          _1005_v18 = _module.D12.create_DC36(false, p1, (_this).f25, (_this).f25, _1004_v17);
          let _1006_v19;
          _1006_v19 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,_1001_v14)).length),_988_v3);
          let _1007_v20;
          _1007_v20 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_this).f26,_988_v3), _1006_v19);
          let _1008_v21;
          _1008_v21 = _dafny.Set.fromElements((_this).f26);
          let _rhs217 = (_this).f23;
          let _rhs218 = p1;
          let _rhs219 = _1001_v14;
          let _rhs220 = _module.__default.safeModuloInt(new BigNumber(((_1007_v20)[_module.__default.safeIndex((_this).f26, new BigNumber((_1007_v20).length))]).length), new BigNumber(((_1008_v21).Union(_1008_v21)).length));
          let _rhs221 = _module.__default.fm31((_this).f25, globalState);
          let _lhs197 = globalState;
          let _lhs198 = globalState;
          let _lhs199 = globalState;
          let _lhs200 = globalState;
          _lhs197.f8 = _rhs217;
          _lhs198.f7 = _rhs218;
          _lhs199.f6 = _rhs219;
          _lhs200.f8 = _rhs220;
          _1005_v18 = _rhs221;
          let _1009_v22;
          _1009_v22 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(344),p1);
          let _1010_v23;
          _1010_v23 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
          let _1011_v24;
          _1011_v24 = _module.D0.create_DC1((_this).f26, p1, new BigNumber((_1010_v23).length));
          _1009_v22 = (_1009_v22).update((_this).f25, (_1011_v24).dtor_cf2);
          let _1012_v25;
          let _nw159 = new _module.C0();
          _nw159.__ctor();
          _1012_v25 = _nw159;
          let _index201 = _module.__default.safeIndex(new BigNumber(200), new BigNumber((p0).length));
          (p0)[_index201] = p1;
          let _index202 = _module.__default.safeIndex(new BigNumber(200), new BigNumber((p0).length));
          (p0)[_index202] = p1;
          let _1013_v26;
          let _nw160 = Array((new BigNumber(9)).toNumber()).fill([]);
          _1013_v26 = _nw160;
          let _1014_v27;
          let _nw161 = Array((new BigNumber(15)).toNumber()).fill(false);
          _1014_v27 = _nw161;
          let _index203 = _module.__default.safeIndex(new BigNumber(402), new BigNumber((_1013_v26).length));
          (_1013_v26)[_index203] = _1014_v27;
          let _index204 = _module.__default.safeIndex(new BigNumber(402), new BigNumber((_1013_v26).length));
          (_1013_v26)[_index204] = _1014_v27;
        } else {
          let _1015_v28;
          _1015_v28 = _dafny.MultiSet.fromElements(new BigNumber((_1001_v14).length));
          (globalState).f7 = (p1) === ((_1015_v28).IsSubsetOf(_dafny.MultiSet.fromElements((_this).f26)));
          let _1016_v29;
          let _init24 = ((_1017_i2) => function (_1018_i3) {
            return (_1018_i3).multipliedBy(_1017_i2);
          })(_998_i2);
          let _nw162 = Array((new BigNumber(27)).toNumber());
          for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw162.length); _i0_24++) {
            _nw162[_i0_24] = _init24(new BigNumber(_i0_24));
          }
          _1016_v29 = _nw162;
          let _1019_v30;
          _1019_v30 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1016_v29);
          let _1020_v31;
          _1020_v31 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_this).f26, _998_i2),_1019_v30);
          _1020_v31 = (_1020_v31).update((_this).f25, _1019_v30);
          let _1021_v32;
          _1021_v32 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_module.__default.fm0((_this).f25, _module.__default.fm1((_this).f23, p1, globalState), globalState)),(_this).f26);
          let _1022_v33;
          _1022_v33 = _dafny.Map.Empty.slice().updateUnsafe(_985_v1,p1);
          let _1023_v34;
          let _nw163 = new _module.C2();
          _nw163.__ctor(_1021_v32, (_1022_v33).equals(_1022_v33), p1, _dafny.Seq.Concat(_dafny.Seq.of((_this).f25, (_this).f26), _985_v1), (_this).f25);
          _1023_v34 = _nw163;
          let _1024_v35;
          let _init25 = function (_1025_i4) {
            return _module.D12.create_DC37(_module.D12.create_DC35());
          };
          let _nw164 = Array((new BigNumber(18)).toNumber());
          for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw164.length); _i0_25++) {
            _nw164[_i0_25] = _init25(new BigNumber(_i0_25));
          }
          _1024_v35 = _nw164;
          let _1026_v36;
          _1026_v36 = _module.D12.create_DC35();
          let _1027_v37;
          _1027_v37 = _module.D12.create_DC37(_1026_v36);
          let _index205 = _module.__default.safeIndex(new BigNumber(391), new BigNumber((_1024_v35).length));
          (_1024_v35)[_index205] = _1027_v37;
          let _index206 = _module.__default.safeIndex(new BigNumber(391), new BigNumber((_1024_v35).length));
          (_1024_v35)[_index206] = _1027_v37;
          let _index207 = _module.__default.safeIndex(new BigNumber(528), new BigNumber((_1016_v29).length));
          (_1016_v29)[_index207] = (_this).f25;
          let _index208 = _module.__default.safeIndex(new BigNumber(528), new BigNumber((_1016_v29).length));
          (_1016_v29)[_index208] = ((_this).f23).minus((_this).f26);
        }
        r0 = ((_dafny.ZERO).minus((_this).f26)).multipliedBy((_this).f25);
        let _1028_v38;
        let _init26 = function (_1029_i5) {
          return (_1029_i5).minus((_this).f25);
        };
        let _nw165 = Array((new BigNumber(7)).toNumber());
        for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw165.length); _i0_26++) {
          _nw165[_i0_26] = _init26(new BigNumber(_i0_26));
        }
        _1028_v38 = _nw165;
        _1028_v38 = _1028_v38;
      }
      r0 = (_this).f23;
      let _1030_v39;
      _1030_v39 = _dafny.Set.fromElements((_this).f23);
      r1 = ((_1030_v39).Intersect(_1030_v39)).Intersect(_1030_v39);
      return [r0, r1];
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Seq.of();
      let _1031_v0;
      _1031_v0 = false;
      let _1032_v1;
      _1032_v1 = _dafny.Set.fromElements(_1031_v0, _1031_v0);
      let _1033_v2;
      _1033_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f25);
      let _1034_v3;
      _1034_v3 = _module.D13.create_DC39((((_1033_v2).contains((_this).f23)) ? ((_1033_v2).get((_this).f23)) : ((_this).f23)));
      let _1035_v4;
      _1035_v4 = _dafny.Set.fromElements(new BigNumber((_1032_v1).length));
      let _1036_v5;
      let _nw166 = Array((new BigNumber(12)).toNumber());
      _nw166[(_dafny.ZERO).toNumber()] = (_this).f23;
      _nw166[(_dafny.ONE).toNumber()] = new BigNumber((p0).length);
      _nw166[(new BigNumber(2)).toNumber()] = (_this).f23;
      _nw166[(new BigNumber(3)).toNumber()] = (_this).f23;
      _nw166[(new BigNumber(4)).toNumber()] = new BigNumber(700);
      _nw166[(new BigNumber(5)).toNumber()] = (_this).f25;
      _nw166[(new BigNumber(6)).toNumber()] = (_1034_v3).dtor_cf63;
      _nw166[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus((_this).f25);
      _nw166[(new BigNumber(8)).toNumber()] = (_this).f26;
      _nw166[(new BigNumber(9)).toNumber()] = (_this).f25;
      _nw166[(new BigNumber(10)).toNumber()] = new BigNumber((_1035_v4).length);
      _nw166[(new BigNumber(11)).toNumber()] = new BigNumber(947);
      _1036_v5 = _nw166;
      let _1037_v6;
      _1037_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1032_v1,_1036_v5);
      let _1038_v7;
      _1038_v7 = _module.D13.create_DC38(_1032_v1);
      let _1039_v8;
      _1039_v8 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1((_this).f23, !(_1031_v0), globalState),(_1038_v7).dtor_cf62);
      let _1040_v9;
      _1040_v9 = _module.D0.create_DC0(p0);
      let _rhs222 = _dafny.Map.Empty.slice().updateUnsafe((((_1039_v8).contains(false)) ? ((_1039_v8).get(false)) : (_1032_v1)),_1036_v5);
      let _rhs223 = (_1032_v1).Difference(_1032_v1);
      let _rhs224 = _dafny.Seq.Concat((_1040_v9).dtor_cf0, p0);
      let _lhs201 = globalState;
      _1037_v6 = _rhs222;
      _1032_v1 = _rhs223;
      _lhs201.f6 = _rhs224;
      _1032_v1 = _1032_v1;
      let _1041_v10;
      _1041_v10 = new _dafny.CodePoint('k'.codePointAt(0));
      (globalState).f11 = _1041_v10;
      r0 = _module.__default.safeModuloInt(((_this).f26).minus((_this).f25), (_this).f26);
      let _1042_v11;
      let _nw167 = Array((new BigNumber(3)).toNumber());
      _nw167[(_dafny.ZERO).toNumber()] = _1041_v10;
      _nw167[(_dafny.ONE).toNumber()] = _1041_v10;
      _nw167[(new BigNumber(2)).toNumber()] = _1041_v10;
      _1042_v11 = _nw167;
      let _1043_v12;
      _1043_v12 = _module.D9.create_DC24(_1042_v11);
      let _1044_v13;
      _1044_v13 = _dafny.MultiSet.fromElements(_1043_v12);
      let _1045_v14;
      _1045_v14 = _dafny.Seq.of(_module.D9.create_DC24(_1042_v11), _1043_v12, _1043_v12);
      (globalState).f8 = new BigNumber((((_1044_v13).Difference(_1044_v13)).Difference(_dafny.MultiSet.FromArray(_1045_v14))).cardinality());
      let _1046_v15;
      _1046_v15 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),_1031_v0);
      _1046_v15 = (_1046_v15).update(_1041_v10, _1031_v0);
      r0 = _module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f25), new BigNumber((p0).length));
      r1 = _dafny.Seq.of((_this).f25);
      return [r0, r1];
    }
    m3(globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Seq.of();
      let r2 = _dafny.Map.Empty;
      let _1047_v0;
      _1047_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,(_this).f26);
      let _1048_v1;
      _1048_v1 = false;
      let _1049_v2;
      _1049_v2 = _module.D1.create_DC5(new BigNumber(954), _1048_v1, (_this).f23, (_this).f26);
      let _1050_v3;
      _1050_v3 = _dafny.Seq.of((_this).f25);
      let _1051_v4;
      _1051_v4 = _dafny.Map.Empty.slice().updateUnsafe(_1048_v1,_1048_v1);
      let _1052_v5;
      let _nw168 = new _module.C2();
      _nw168.__ctor(_1047_v0, _1048_v1, (_1049_v2).dtor_cf10, _1050_v3, new BigNumber((_1051_v4).length));
      _1052_v5 = _nw168;
      let _1053_v6;
      _1053_v6 = _dafny.Seq.of(_1052_v5);
      let _1054_v7;
      _1054_v7 = _dafny.Seq.UnicodeFromString("uckqsq");
      let _1055_v8;
      _1055_v8 = _dafny.Seq.of(_1054_v7);
      let _hi12 = ((false) ? (new BigNumber(((_1055_v8)[_module.__default.safeIndex((_this).f23, new BigNumber((_1055_v8).length))]).length)) : ((_this).f26));
      for (let _1056_i0 = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_1053_v6, _dafny.Seq.update(_1053_v6, _module.__default.safeIndex(new BigNumber(-858), new BigNumber((_1053_v6).length)), _1052_v5)), _module.__default.safeIndex((_1052_v5).f23, new BigNumber((_dafny.Seq.Concat(_1053_v6, _dafny.Seq.update(_1053_v6, _module.__default.safeIndex(new BigNumber(-858), new BigNumber((_1053_v6).length)), _1052_v5))).length)), _1052_v5)).length); _1056_i0.isLessThan(_hi12); _1056_i0 = _1056_i0.plus(_dafny.ONE)) {
        (globalState).f9 = !((_this).f25).isEqualTo((_this).f26);
        let _1057_v9;
        let _nw169 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
        _1057_v9 = _nw169;
        let _init27 = function (_1058_i1) {
          return _module.__default.safeDivisionInt(_1058_i1, (_this).f23);
        };
        let _nw170 = Array((new BigNumber(28)).toNumber());
        for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw170.length); _i0_27++) {
          _nw170[_i0_27] = _init27(new BigNumber(_i0_27));
        }
        _1057_v9 = _nw170;
        r0 = ((true) ? ((_1052_v5).f28) : ((_1052_v5).f28));
        _1057_v9 = _1057_v9;
      }
      if (_1048_v1) {
        let _1059_v10;
        _1059_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1054_v7,(_1052_v5).f28);
        _1059_v10 = (_1059_v10).update(_1054_v7, _1048_v1);
        (_1052_v5).f29 = _1050_v3;
        (globalState).f7 = false;
        let _1060_v11;
        _1060_v11 = _dafny.MultiSet.fromElements((_1052_v5).f28, (_1052_v5).f28);
        let _1061_v12;
        _1061_v12 = _dafny.Seq.of(_dafny.MultiSet.fromElements((_1052_v5).f28));
        let _1062_v13;
        let _nw171 = new _module.C3();
        _nw171.__ctor(_dafny.Seq.Concat(_dafny.Seq.of(_1060_v11, _1060_v11), _1061_v12), (_this).f25, (_1052_v5).f28, _1050_v3);
        _1062_v13 = _nw171;
        let _1063_v14;
        _1063_v14 = new _dafny.CodePoint('a'.codePointAt(0));
        let _1064_v15;
        _1064_v15 = _dafny.Seq.of(!(_1048_v1));
        let _1065_v16;
        let _nw172 = Array((new BigNumber(27)).toNumber());
        _nw172[(_dafny.ZERO).toNumber()] = (_1052_v5).f28;
        _nw172[(_dafny.ONE).toNumber()] = (_1052_v5).f28;
        _nw172[(new BigNumber(2)).toNumber()] = (_1052_v5).f28;
        _nw172[(new BigNumber(3)).toNumber()] = true;
        _nw172[(new BigNumber(4)).toNumber()] = (_1052_v5).f28;
        _nw172[(new BigNumber(5)).toNumber()] = (_1052_v5).f28;
        _nw172[(new BigNumber(6)).toNumber()] = (_1052_v5).f28;
        _nw172[(new BigNumber(7)).toNumber()] = !(_1048_v1);
        _nw172[(new BigNumber(8)).toNumber()] = true;
        _nw172[(new BigNumber(9)).toNumber()] = (_1052_v5).f28;
        _nw172[(new BigNumber(10)).toNumber()] = (_1052_v5).f28;
        _nw172[(new BigNumber(11)).toNumber()] = _1048_v1;
        _nw172[(new BigNumber(12)).toNumber()] = _1048_v1;
        _nw172[(new BigNumber(13)).toNumber()] = (_1064_v15)[_module.__default.safeIndex((_this).f26, new BigNumber((_1064_v15).length))];
        _nw172[(new BigNumber(14)).toNumber()] = _1048_v1;
        _nw172[(new BigNumber(15)).toNumber()] = _1048_v1;
        _nw172[(new BigNumber(16)).toNumber()] = _module.__default.fm1(new BigNumber(-246), _1048_v1, globalState);
        _nw172[(new BigNumber(17)).toNumber()] = _1048_v1;
        _nw172[(new BigNumber(18)).toNumber()] = (_1052_v5).f28;
        _nw172[(new BigNumber(19)).toNumber()] = (_1052_v5).f28;
        _nw172[(new BigNumber(20)).toNumber()] = false;
        _nw172[(new BigNumber(21)).toNumber()] = (_1052_v5).f28;
        _nw172[(new BigNumber(22)).toNumber()] = _1048_v1;
        _nw172[(new BigNumber(23)).toNumber()] = (_1064_v15)[_module.__default.safeIndex(new BigNumber(356), new BigNumber((_1064_v15).length))];
        _nw172[(new BigNumber(24)).toNumber()] = true;
        _nw172[(new BigNumber(25)).toNumber()] = (_1052_v5).f28;
        _nw172[(new BigNumber(26)).toNumber()] = _1048_v1;
        _1065_v16 = _nw172;
        let _1066_v17;
        _1066_v17 = _module.D9.create_DC27((_1052_v5).f28, _dafny.Seq.Create(_module.__default.abs(new BigNumber(19)), function (_1067_i2) {
  return (_this).f23;
}), _1065_v16, (_1052_v5).f28, (_1064_v15)[_module.__default.safeIndex((_this).f23, new BigNumber((_1064_v15).length))]);
        let _1068_v18;
        _1068_v18 = _module.D9.create_DC28(_1066_v17);
        let _1069_v19;
        _1069_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1063_v14,_1068_v18);
        _1069_v19 = (_1069_v19).update(_1063_v14, _1068_v18);
      } else {
        let _1070_v20;
        let _nw173 = new _module.C0();
        _nw173.__ctor();
        _1070_v20 = _nw173;
        (globalState).f14 = (_dafny.ZERO).minus(((_1052_v5).f23).plus((_1052_v5).f23));
        let _1071_v21;
        _1071_v21 = _module.D11.create_DC32(_1054_v7);
        let _source8 = _1071_v21;
        if (_source8.is_DC31) {
          let _1072___mcc_h0 = (_source8).cf52;
          let _1073_cf52 = _1072___mcc_h0;
          let _1074_v22;
          _1074_v22 = _dafny.Set.fromElements((_1052_v5).f23, new BigNumber(335), new BigNumber((_1051_v4).length), (_this).f23, (_1050_v3)[_module.__default.safeIndex(_1073_cf52, new BigNumber((_1050_v3).length))]);
          let _1075_v23;
          let _nw174 = new _module.C3();
          _nw174.__ctor(_module.__default.fm29(_module.__default.fm0(new BigNumber(383), (_1052_v5).f28, globalState), _1074_v22, globalState), (_1052_v5).f23, (_1052_v5).f28, _dafny.Seq.Concat(_dafny.Seq.of((_this).f23), _1052_v5.f29));
          _1075_v23 = _nw174;
          _1051_v4 = (_1051_v4).update(_module.__default.fm1(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("mnssog"),new BigNumber(387))).length), (_1052_v5).f28, globalState), _1048_v1);
          let _1076_v24;
          _1076_v24 = _dafny.Set.fromElements((_1052_v5).f28, _1048_v1);
          _1076_v24 = _1076_v24;
          let _1077_v25;
          _1077_v25 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm24((_1052_v5).f28, new BigNumber((_1054_v7).length), globalState),(_1052_v5).f23);
          let _1078_v26;
          _1078_v26 = _module.D9.create_DC26((_1052_v5).f23);
          _1077_v25 = (_1077_v25).update(_1078_v26, (_1052_v5).f23);
        } else if (_source8.is_DC32) {
          let _1079___mcc_h1 = (_source8).cf53;
          let _1080_cf53 = _1079___mcc_h1;
          let _rhs225 = (_1052_v5).f23;
          let _rhs226 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(695)), function (_1081_i3) {
            return new _dafny.CodePoint('f'.codePointAt(0));
          });
          let _lhs202 = globalState;
          let _lhs203 = globalState;
          _lhs202.f14 = _rhs225;
          _lhs203.f6 = _rhs226;
          let _1082_v27;
          _1082_v27 = _dafny.Seq.of((_1052_v5).f28);
          _1048_v1 = !_dafny.Seq.contains(_1082_v27, (_1052_v5).f28);
          let _1083_v28;
          _1083_v28 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(915)), function (_1084_i4) {
            return new _dafny.CodePoint('h'.codePointAt(0));
          }));
          (globalState).f7 = !(_1083_v28).equals(_dafny.Set.fromElements(_1080_cf53, _1054_v7));
          _1050_v3 = _1052_v5.f29;
        } else if (_source8.is_DC30) {
          let _1085___mcc_h2 = (_source8).cf51;
          let _1086_cf51 = _1085___mcc_h2;
          (globalState).f8 = (_module.__default.safeModuloInt((_this).f26, (_dafny.ZERO).minus((_1052_v5).f23))).multipliedBy(new BigNumber((_1054_v7).length));
          let _1087_v29;
          _1087_v29 = _dafny.Set.fromElements((_1052_v5).f23);
          let _1088_v30;
          _1088_v30 = _dafny.Map.Empty.slice().updateUnsafe(_1048_v1,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1070_v20,(_1052_v5).f28)).length));
          let _1089_v31;
          _1089_v31 = _dafny.Map.Empty.slice().updateUnsafe(false,_1088_v30);
          (globalState).f21 = (_1089_v31).contains(!((_dafny.ZERO).minus(_module.__default.fm0(new BigNumber((_1087_v29).length), !((_1052_v5).f28), globalState))).isEqualTo((_this).f23));
          (globalState).f7 = (_1052_v5).f28;
          let _1090_v32;
          _1090_v32 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("r")).length),false);
          let _1091_v33;
          let _nw175 = new _module.C2();
          _nw175.__ctor(_1047_v0, !((_1052_v5).f28), _1048_v1, _1052_v5.f29, new BigNumber(((_1090_v32).update((_1052_v5).f23, (_1052_v5).f28)).length));
          _1091_v33 = _nw175;
          let _1092_v34;
          _1092_v34 = _dafny.MultiSet.fromElements(!(true));
          let _1093_v35;
          _1093_v35 = _module.D9.create_DC26((_1052_v5).f23);
          let _1094_v36;
          _1094_v36 = _dafny.Set.fromElements(false);
          let _1095_v37;
          let _nw176 = Array((new BigNumber(17)).toNumber());
          _nw176[(_dafny.ZERO).toNumber()] = _module.D9.create_DC26((_1052_v5).f23);
          _nw176[(_dafny.ONE).toNumber()] = _module.D9.create_DC26((((_1092_v34).contains((_1052_v5).f28)) ? ((_1092_v34).get((_1052_v5).f28)) : (new BigNumber(418))));
          _nw176[(new BigNumber(2)).toNumber()] = _1093_v35;
          _nw176[(new BigNumber(3)).toNumber()] = _1093_v35;
          _nw176[(new BigNumber(4)).toNumber()] = _module.__default.fm24((_1052_v5).f28, new BigNumber(159), globalState);
          _nw176[(new BigNumber(5)).toNumber()] = _module.D9.create_DC26((_this).f25);
          _nw176[(new BigNumber(6)).toNumber()] = _1093_v35;
          _nw176[(new BigNumber(7)).toNumber()] = _1093_v35;
          _nw176[(new BigNumber(8)).toNumber()] = _1093_v35;
          _nw176[(new BigNumber(9)).toNumber()] = _1093_v35;
          _nw176[(new BigNumber(10)).toNumber()] = _1093_v35;
          _nw176[(new BigNumber(11)).toNumber()] = _1093_v35;
          _nw176[(new BigNumber(12)).toNumber()] = _module.D9.create_DC26(new BigNumber((_1094_v36).length));
          _nw176[(new BigNumber(13)).toNumber()] = _1093_v35;
          _nw176[(new BigNumber(14)).toNumber()] = _1093_v35;
          _nw176[(new BigNumber(15)).toNumber()] = _1093_v35;
          _nw176[(new BigNumber(16)).toNumber()] = _1093_v35;
          _1095_v37 = _nw176;
          let _1096_v38;
          _1096_v38 = _dafny.Map.Empty.slice().updateUnsafe(_1091_v33,_1095_v37);
          _1096_v38 = (_1096_v38).update(_1091_v33, _1095_v37);
        } else {
          let _1097___mcc_h3 = (_source8).cf54;
          let _1098_cf54 = _1097___mcc_h3;
          let _1099_v39;
          _1099_v39 = _module.D9.create_DC25(_1048_v1, _dafny.Set.fromElements(new _dafny.CodePoint('v'.codePointAt(0))));
          let _1100_v40;
          _1100_v40 = _module.D2.create_DC7((_1099_v39).dtor_cf42);
          let _1101_v41;
          _1101_v41 = _dafny.Map.Empty.slice().updateUnsafe(_1100_v40,_1054_v7);
          let _1102_v42;
          _1102_v42 = _dafny.Set.fromElements(_1070_v20);
          let _1103_v43;
          _1103_v43 = _dafny.Set.fromElements(_1048_v1, (_1052_v5).f28, (_1052_v5).f28);
          let _1104_v44;
          _1104_v44 = _dafny.Seq.of(!((_1070_v20).fm20(_1101_v41, (_1052_v5).f23, (_this).f26, new BigNumber((_1102_v42).length), globalState)) || ((_1052_v5).f28), (_1052_v5).f28, (((_1052_v5).f28) ? ((_1052_v5).f28) : ((_1052_v5).f28)), ((_1052_v5).f23).isLessThanOrEqualTo(new BigNumber((_1103_v43).length)));
          (globalState).f21 = (_1104_v44)[_module.__default.safeIndex((((_1052_v5).f28) ? ((_this).f25) : (new BigNumber(414))), new BigNumber((_1104_v44).length))];
          let _1105_v45;
          let _nw177 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Seq.of());
          _1105_v45 = _nw177;
          let _1106_v46;
          _1106_v46 = _dafny.Seq.of(_dafny.Seq.of((_1052_v5).f23, (_dafny.ZERO).minus((_1052_v5).f23)));
          let _index209 = _module.__default.safeIndex(new BigNumber(694), new BigNumber((_1105_v45).length));
          (_1105_v45)[_index209] = _1106_v46;
          let _index210 = _module.__default.safeIndex(new BigNumber(694), new BigNumber((_1105_v45).length));
          (_1105_v45)[_index210] = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(506)), ((_1107_v5) => function (_1108_i5) {
            return _1107_v5.f29;
          })(_1052_v5)), _module.__default.safeIndex(((_1048_v1) ? ((_1052_v5).f23) : (new BigNumber((_1055_v8).length))), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(506)), ((_1109_v5) => function (_1110_i5) {
            return _1109_v5.f29;
          })(_1052_v5))).length)), _1052_v5.f29);
          let _1111_v47;
          let _nw178 = Array((new BigNumber(26)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1111_v47 = _nw178;
          let _1112_v48;
          _1112_v48 = _module.D9.create_DC24(_1111_v47);
          let _1113_v49;
          _1113_v49 = _module.D9.create_DC28(_1112_v48);
          _1113_v49 = _1113_v49;
          let _1114_v50;
          let _nw179 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Set.Empty);
          _1114_v50 = _nw179;
          let _1115_v51;
          _1115_v51 = new _dafny.CodePoint('l'.codePointAt(0));
          let _1116_v52;
          _1116_v52 = _dafny.Map.Empty.slice().updateUnsafe(_1115_v51,(_1052_v5).f28);
          let _1117_v53;
          _1117_v53 = _dafny.Set.fromElements(_1116_v52);
          let _1118_v54;
          _1118_v54 = _dafny.Map.Empty.slice().updateUnsafe(_1116_v52,(_1052_v5).f23);
          let _index211 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_1114_v50).length));
          (_1114_v50)[_index211] = (_1117_v53).Union(function () {
            let _coll43 = new _dafny.Set();
            for (const _compr_43 of (_1118_v54).Keys.Elements) {
              let _1119_v55 = _compr_43;
              if ((_1118_v54).contains(_1119_v55)) {
                _coll43.add(_1119_v55);
              }
            }
            return _coll43;
          }());
          let _index212 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_1114_v50).length));
          let _rhs227 = (_1052_v5).f23;
          let _rhs228 = (_1117_v53).Difference(_1117_v53);
          let _rhs229 = _dafny.Seq.update(_1054_v7, _module.__default.safeIndex(new BigNumber((_1051_v4).length), new BigNumber((_1054_v7).length)), _1115_v51);
          let _lhs204 = globalState;
          let _lhs205 = _1114_v50;
          let _lhs206 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_1114_v50).length));
          let _lhs207 = globalState;
          _lhs204.f14 = _rhs227;
          _lhs205[_lhs206] = _rhs228;
          _lhs207.f6 = _rhs229;
        }
        let _1120_v56;
        _1120_v56 = _dafny.Seq.of(_1048_v1, (_1052_v5).f28, (_1052_v5).f28);
        let _1121_v57;
        let _nw180 = new _module.C1();
        _nw180.__ctor(!(!((_1120_v56)[_module.__default.safeIndex((_1052_v5).f23, new BigNumber((_1120_v56).length))])), _1052_v5.f29, ((_this).f23).plus((_this).f23));
        _1121_v57 = _nw180;
        let _1122_v58;
        _1122_v58 = new _dafny.CodePoint('y'.codePointAt(0));
        let _1123_v59;
        let _init28 = ((_1124_v5) => function (_1125_i6) {
          return (_1125_i6).multipliedBy((_1124_v5).f23);
        })(_1052_v5);
        let _nw181 = Array((new BigNumber(17)).toNumber());
        for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw181.length); _i0_28++) {
          _nw181[_i0_28] = _init28(new BigNumber(_i0_28));
        }
        _1123_v59 = _nw181;
        let _1126_v60;
        _1126_v60 = _dafny.Seq.of(_1123_v59, _1123_v59);
        let _1127_v61;
        _1127_v61 = _module.D1.create_DC3(_1126_v60);
        let _1128_v62;
        _1128_v62 = _dafny.Map.Empty.slice().updateUnsafe(_1122_v58,_1127_v61);
        let _1129_v63;
        _1129_v63 = _dafny.Set.fromElements((_1052_v5).f28, !((_1052_v5).f28), (_1052_v5).f28);
        let _1130_v64;
        _1130_v64 = _dafny.Map.Empty.slice().updateUnsafe(_1129_v63,(_1128_v62).Merge(_1128_v62));
        _1128_v62 = (((_1130_v64).contains((_1129_v63).Difference(_1129_v63))) ? ((_1130_v64).get((_1129_v63).Difference(_1129_v63))) : (_1128_v62));
      }
      (globalState).f7 = (_1052_v5).f28;
      let _hi13 = ((_1052_v5).f23).minus((_1052_v5).f23);
      for (let _1131_i7 = (_this).f23; _1131_i7.isLessThan(_hi13); _1131_i7 = _1131_i7.plus(_dafny.ONE)) {
        let _1132_v65;
        let _init29 = function (_1133_i8) {
          return _module.__default.safeDivisionInt(_1133_i8, (_this).f23);
        };
        let _nw182 = Array((new BigNumber(18)).toNumber());
        for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw182.length); _i0_29++) {
          _nw182[_i0_29] = _init29(new BigNumber(_i0_29));
        }
        _1132_v65 = _nw182;
        let _1134_v66;
        let _nw183 = new _module.C4();
        _nw183.__ctor(_1132_v65, new BigNumber((_1054_v7).length));
        _1134_v66 = _nw183;
        let _1135_v67;
        _1135_v67 = _dafny.MultiSet.fromElements(_1134_v66);
        let _1136_v68;
        _1136_v68 = _module.D11.create_DC32(_1054_v7);
        let _pat_let_tv38 = _1054_v7;
        let _1137_v69;
        _1137_v69 = _dafny.MultiSet.fromElements((function (_pat_let18_0) {
          return function (_1138_dt__update__tmp_h0) {
            return function (_pat_let19_0) {
              return function (_1139_dt__update_hcf53_h0) {
                return _module.D11.create_DC32(_1139_dt__update_hcf53_h0);
              }(_pat_let19_0);
            }(_pat_let_tv38);
          }(_pat_let18_0);
        }(_1136_v68)).dtor_cf53, _1054_v7);
        if (_module.__default.fm1((((_1135_v67).contains(_1134_v66)) ? ((_1135_v67).get(_1134_v66)) : ((_this).f23)), (_1137_v69).IsSubsetOf(_1137_v69), globalState)) {
          let _1140_v70;
          let _init30 = ((_1141_v5) => function (_1142_i9) {
            return ((_1141_v5).f28) === (!((_1141_v5).f28));
          })(_1052_v5);
          let _nw184 = Array((new BigNumber(21)).toNumber());
          for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw184.length); _i0_30++) {
            _nw184[_i0_30] = _init30(new BigNumber(_i0_30));
          }
          _1140_v70 = _nw184;
          let _index213 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_1140_v70).length));
          (_1140_v70)[_index213] = (_1052_v5).f28;
          let _index214 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_1140_v70).length));
          (_1140_v70)[_index214] = (_1052_v5).f28;
          let _1143_v71;
          let _nw185 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Map.Empty);
          _1143_v71 = _nw185;
          let _1144_v72;
          _1144_v72 = _dafny.Map.Empty.slice().updateUnsafe((_1052_v5).f28,_1143_v71);
          _1144_v72 = (_1144_v72).update(!(!(true)) || (true), _1143_v71);
          let _1145_v73;
          _1145_v73 = new _dafny.CodePoint('b'.codePointAt(0));
          (globalState).f11 = _1145_v73;
          let _1146_v74;
          let _nw186 = Array((new BigNumber(7)).toNumber());
          _nw186[(_dafny.ZERO).toNumber()] = _1054_v7;
          _nw186[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_1054_v7, _1054_v7);
          _nw186[(new BigNumber(2)).toNumber()] = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("vrrj"), _module.__default.safeIndex(_module.__default.fm0((_this).f26, (_1052_v5).f28, globalState), new BigNumber((_dafny.Seq.UnicodeFromString("vrrj")).length)), _1145_v73);
          _nw186[(new BigNumber(3)).toNumber()] = _1054_v7;
          _nw186[(new BigNumber(4)).toNumber()] = _module.__default.fm2(globalState);
          _nw186[(new BigNumber(5)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(379)), ((_1147_v73) => function (_1148_i10) {
            return _1147_v73;
          })(_1145_v73));
          _nw186[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("etxm"), _1054_v7);
          _1146_v74 = _nw186;
          let _1149_v75;
          _1149_v75 = _dafny.Map.Empty.slice().updateUnsafe((_1052_v5).f28,_1131_i7);
          let _index215 = _module.__default.safeIndex(new BigNumber(926), new BigNumber((_1146_v74).length));
          (_1146_v74)[_index215] = _dafny.Seq.update(_dafny.Seq.Concat(_1054_v7, _1054_v7), _module.__default.safeIndex((((_1149_v75).contains(_1048_v1)) ? ((_1149_v75).get(_1048_v1)) : (new BigNumber((_1149_v75).length))), new BigNumber((_dafny.Seq.Concat(_1054_v7, _1054_v7)).length)), _1145_v73);
          let _index216 = _module.__default.safeIndex(new BigNumber(926), new BigNumber((_1146_v74).length));
          (_1146_v74)[_index216] = _1054_v7;
          let _1150_v76;
          _1150_v76 = _dafny.MultiSet.fromElements(_1145_v73);
          let _1151_v77;
          _1151_v77 = _module.D8.create_DC20((_1052_v5).f28, _1131_i7);
          let _1152_v78;
          let _nw187 = new _module.C1();
          _nw187.__ctor((_1052_v5).f28, _1050_v3, (_1052_v5).f23);
          _1152_v78 = _nw187;
          let _1153_v79;
          let _nw188 = Array((new BigNumber(18)).toNumber());
          _nw188[(_dafny.ZERO).toNumber()] = _1152_v78;
          _nw188[(_dafny.ONE).toNumber()] = _1152_v78;
          _nw188[(new BigNumber(2)).toNumber()] = _1152_v78;
          _nw188[(new BigNumber(3)).toNumber()] = _1152_v78;
          _nw188[(new BigNumber(4)).toNumber()] = _1152_v78;
          _nw188[(new BigNumber(5)).toNumber()] = _1152_v78;
          _nw188[(new BigNumber(6)).toNumber()] = _1152_v78;
          _nw188[(new BigNumber(7)).toNumber()] = _1152_v78;
          _nw188[(new BigNumber(8)).toNumber()] = _1152_v78;
          _nw188[(new BigNumber(9)).toNumber()] = _1152_v78;
          _nw188[(new BigNumber(10)).toNumber()] = _1152_v78;
          _nw188[(new BigNumber(11)).toNumber()] = _1152_v78;
          _nw188[(new BigNumber(12)).toNumber()] = _1152_v78;
          _nw188[(new BigNumber(13)).toNumber()] = _1152_v78;
          _nw188[(new BigNumber(14)).toNumber()] = _1152_v78;
          _nw188[(new BigNumber(15)).toNumber()] = _1152_v78;
          _nw188[(new BigNumber(16)).toNumber()] = _1152_v78;
          _nw188[(new BigNumber(17)).toNumber()] = _1152_v78;
          _1153_v79 = _nw188;
          let _1154_v80;
          _1154_v80 = _dafny.Map.Empty.slice().updateUnsafe((_1052_v5).f28,_1153_v79);
          let _1155_v81;
          _1155_v81 = _dafny.Map.Empty.slice().updateUnsafe(!((_1140_v70)[_module.__default.safeIndex(new BigNumber(470), new BigNumber((_1140_v70).length))]),_1052_v5.f29);
          let _1156_v82;
          _1156_v82 = _dafny.MultiSet.fromElements(new BigNumber((_1155_v81).length));
          (globalState).f14 = _module.__default.fm0((new BigNumber(((_dafny.MultiSet.FromArray(_1052_v5.f29)).update(new BigNumber((_1150_v76).cardinality()), _module.__default.abs((_1151_v77).dtor_cf34))).cardinality())).minus((_dafny.ZERO).minus(new BigNumber((_module.__default.fm2(globalState)).length))), !((_dafny.MultiSet.fromElements(new BigNumber((_1154_v80).length), (_1052_v5).f23)).IsProperSubsetOf(_1156_v82)), globalState);
        } else {
          let _index217 = _module.__default.safeIndex(new BigNumber(610), new BigNumber(((_1134_v66).f27).length));
          ((_1134_v66).f27)[_index217] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("hoqqdqjyb"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-545)), function (_1157_i11) {
            return new _dafny.CodePoint('t'.codePointAt(0));
          }))).length);
          let _index218 = _module.__default.safeIndex(new BigNumber(610), new BigNumber(((_1134_v66).f27).length));
          ((_1134_v66).f27)[_index218] = ((_1052_v5).f23).plus(new BigNumber((_1054_v7).length));
          let _1158_v83;
          let _nw189 = new _module.C0();
          _nw189.__ctor();
          _1158_v83 = _nw189;
          let _1159_v84;
          _1159_v84 = _dafny.Seq.of(false);
          let _1160_v85;
          _1160_v85 = new _dafny.CodePoint('l'.codePointAt(0));
          let _1161_v86;
          _1161_v86 = _dafny.Set.fromElements(_1160_v85);
          let _1162_v87;
          _1162_v87 = _module.D9.create_DC25(true, _1161_v86);
          let _1163_v88;
          _1163_v88 = _dafny.Map.Empty.slice().updateUnsafe((_1052_v5).f28,(_this).f26);
          let _1164_v89;
          _1164_v89 = _module.D8.create_DC21((_1134_v66).fm10(_dafny.Seq.Create(_module.__default.abs(new BigNumber(224)), ((_1165_v85) => function (_1166_i12) {
  return _1165_v85;
})(_1160_v85)), (_1052_v5).f28, _1163_v88, (_1159_v84)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("ixco")).length), new BigNumber((_1159_v84).length))], globalState), (_1052_v5).f28, (_1052_v5).f23);
          let _1167_v90;
          _1167_v90 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('q'.codePointAt(0)));
          let _1168_v92;
          _1168_v92 = _module.D2.create_DC7(function () {
  let _coll44 = new _dafny.Set();
  for (const _compr_44 of (_1167_v90).Elements) {
    let _1169_v91 = _compr_44;
    if ((_1167_v90).contains(_1169_v91)) {
      _coll44.add(_1169_v91);
    }
  }
  return _coll44;
}());
          let _1170_v93;
          _1170_v93 = _dafny.Map.Empty.slice().updateUnsafe(_1168_v92,_1054_v7);
          let _1171_v94;
          let _init31 = ((_1172_v1) => function (_1173_i13) {
            return _1172_v1;
          })(_1048_v1);
          let _nw190 = Array((new BigNumber(21)).toNumber());
          for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw190.length); _i0_31++) {
            _nw190[_i0_31] = _init31(new BigNumber(_i0_31));
          }
          _1171_v94 = _nw190;
          let _1174_v95;
          _1174_v95 = _dafny.Set.fromElements(_1050_v3, _1052_v5.f29, _1050_v3);
          let _1175_v96;
          let _nw191 = Array((new BigNumber(27)).toNumber());
          _nw191[(_dafny.ZERO).toNumber()] = _1048_v1;
          _nw191[(_dafny.ONE).toNumber()] = _1048_v1;
          _nw191[(new BigNumber(2)).toNumber()] = (_1159_v84)[_module.__default.safeIndex(new BigNumber((_1054_v7).length), new BigNumber((_1159_v84).length))];
          _nw191[(new BigNumber(3)).toNumber()] = _1048_v1;
          _nw191[(new BigNumber(4)).toNumber()] = false;
          _nw191[(new BigNumber(5)).toNumber()] = (_1162_v87).dtor_cf41;
          _nw191[(new BigNumber(6)).toNumber()] = (((_1052_v5).f28) ? (_1048_v1) : (false));
          _nw191[(new BigNumber(7)).toNumber()] = (_1052_v5).f28;
          _nw191[(new BigNumber(8)).toNumber()] = (_1052_v5).f28;
          _nw191[(new BigNumber(9)).toNumber()] = (_1052_v5).f28;
          _nw191[(new BigNumber(10)).toNumber()] = (_1052_v5).f28;
          _nw191[(new BigNumber(11)).toNumber()] = (_1052_v5).f28;
          _nw191[(new BigNumber(12)).toNumber()] = (_1052_v5).f28;
          _nw191[(new BigNumber(13)).toNumber()] = false;
          _nw191[(new BigNumber(14)).toNumber()] = !_dafny.areEqual(_1054_v7, _1054_v7);
          _nw191[(new BigNumber(15)).toNumber()] = true;
          _nw191[(new BigNumber(16)).toNumber()] = (_1052_v5).f28;
          _nw191[(new BigNumber(17)).toNumber()] = (_1164_v89).dtor_cf36;
          _nw191[(new BigNumber(18)).toNumber()] = _1048_v1;
          _nw191[(new BigNumber(19)).toNumber()] = (_1052_v5).f28;
          _nw191[(new BigNumber(20)).toNumber()] = !(!(_dafny.Map.Empty.slice().updateUnsafe((_1052_v5).f28,_1171_v94)).contains((_1158_v83).fm20(_1170_v93, new BigNumber((_1054_v7).length), (_1052_v5).f23, (_1052_v5).f23, globalState)));
          _nw191[(new BigNumber(21)).toNumber()] = (_1052_v5).f28;
          _nw191[(new BigNumber(22)).toNumber()] = (_1052_v5).f28;
          _nw191[(new BigNumber(23)).toNumber()] = (_1052_v5).f28;
          _nw191[(new BigNumber(24)).toNumber()] = true;
          _nw191[(new BigNumber(25)).toNumber()] = (_1052_v5).f28;
          _nw191[(new BigNumber(26)).toNumber()] = (_1174_v95).IsSubsetOf(_1174_v95);
          _1175_v96 = _nw191;
          let _1176_v97;
          _1176_v97 = _dafny.Seq.of(_1175_v96, _1175_v96, _1175_v96);
          _1175_v96 = (_1176_v97)[_module.__default.safeIndex((_1052_v5).f23, new BigNumber((_1176_v97).length))];
          (globalState).f8 = _module.__default.fm0(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(545)), function (_1177_i14) {
            return (_this).f26;
          })).length), _1048_v1, globalState);
          let _1178_v98;
          _1178_v98 = _dafny.MultiSet.fromElements(_1159_v84);
          (globalState).f8 = ((((_1178_v98).contains(_1159_v84)) ? ((_1178_v98).get(_1159_v84)) : ((_this).f23))).plus(((_1134_v66).f27)[_module.__default.safeIndex(new BigNumber(610), new BigNumber(((_1134_v66).f27).length))]);
        }
        r0 = _1048_v1;
        let _1179_v99;
        _1179_v99 = new _dafny.CodePoint('g'.codePointAt(0));
        (globalState).f21 = _dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(931)), ((_1180_v99, _1181_v7) => function (_1182_i15) {
          return (_module.D4.create_DC15((_this).f26, _1180_v99, new BigNumber((_1181_v7).length))).dtor_cf27;
        })(_1179_v99, _1054_v7)), _1179_v99);
        let _source9 = _module.D8.create_DC22(_1048_v1);
        if (_source9.is_DC20) {
          let _1183___mcc_h4 = (_source9).cf33;
          let _1184___mcc_h5 = (_source9).cf34;
          let _1185_cf34 = _1184___mcc_h5;
          let _1186_cf33 = _1183___mcc_h4;
          (globalState).f21 = _module.__default.fm1((_this).f26, _module.__default.fm1((_1052_v5).f23, !(_1186_cf33), globalState), globalState);
          let _index219 = _module.__default.safeIndex(new BigNumber(182), new BigNumber(((_1134_v66).f27).length));
          ((_1134_v66).f27)[_index219] = new BigNumber(953);
          let _1187_v100;
          _1187_v100 = _dafny.Set.fromElements(_1048_v1);
          let _index220 = _module.__default.safeIndex(new BigNumber(182), new BigNumber(((_1134_v66).f27).length));
          let _rhs230 = ((_1186_cf33) ? ((_1054_v7)[_module.__default.safeIndex(_1131_i7, new BigNumber((_1054_v7).length))]) : (new _dafny.CodePoint('b'.codePointAt(0))));
          let _rhs231 = _module.__default.fm0(_module.__default.safeDivisionInt(new BigNumber((_1187_v100).length), new BigNumber((_1052_v5.f29).length)), _1186_cf33, globalState);
          let _rhs232 = _module.__default.fm6((_this).f25, globalState);
          let _lhs208 = (_1134_v66).f27;
          let _lhs209 = _module.__default.safeIndex(new BigNumber(182), new BigNumber(((_1134_v66).f27).length));
          let _lhs210 = globalState;
          _1179_v99 = _rhs230;
          _lhs208[_lhs209] = _rhs231;
          _lhs210.f11 = _rhs232;
          let _1188_v101;
          let _init32 = ((_1189_v5) => function (_1190_i16) {
            return (_1189_v5).f28;
          })(_1052_v5);
          let _nw192 = Array((new BigNumber(5)).toNumber());
          for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw192.length); _i0_32++) {
            _nw192[_i0_32] = _init32(new BigNumber(_i0_32));
          }
          _1188_v101 = _nw192;
          let _1191_v102;
          _1191_v102 = _dafny.Map.Empty.slice().updateUnsafe(_1188_v101,(((_1134_v66).f27)[_module.__default.safeIndex(new BigNumber(182), new BigNumber(((_1134_v66).f27).length))]).isLessThanOrEqualTo(new BigNumber((_dafny.MultiSet.fromElements(_1048_v1, (_1052_v5).f28)).cardinality())));
          let _1192_v103;
          _1192_v103 = _dafny.Map.Empty.slice().updateUnsafe((_1052_v5).f28,((_1134_v66).f27)[_module.__default.safeIndex(new BigNumber(182), new BigNumber(((_1134_v66).f27).length))]);
          _1191_v102 = (_1191_v102).update(_1188_v101, (_1134_v66).fm10(_1054_v7, (_1052_v5).f28, _1192_v103, _1048_v1, globalState));
          let _rhs233 = new _dafny.CodePoint('v'.codePointAt(0));
          let _rhs234 = (_dafny.ZERO).minus((_this).f26);
          let _lhs211 = globalState;
          let _lhs212 = globalState;
          _lhs211.f11 = _rhs233;
          _lhs212.f8 = _rhs234;
        } else if (_source9.is_DC21) {
          let _1193___mcc_h6 = (_source9).cf35;
          let _1194___mcc_h7 = (_source9).cf36;
          let _1195___mcc_h8 = (_source9).cf37;
          let _1196_cf37 = _1195___mcc_h8;
          let _1197_cf36 = _1194___mcc_h7;
          let _1198_cf35 = _1193___mcc_h6;
          let _index221 = _module.__default.safeIndex(new BigNumber(129), new BigNumber(((_1134_v66).f27).length));
          ((_1134_v66).f27)[_index221] = _1196_cf37;
          let _index222 = _module.__default.safeIndex(new BigNumber(129), new BigNumber(((_1134_v66).f27).length));
          ((_1134_v66).f27)[_index222] = (_dafny.ZERO).minus(_1131_i7);
          let _1199_v104;
          let _init33 = ((_1200_v5) => function (_1201_i17) {
            return (_1200_v5).f28;
          })(_1052_v5);
          let _nw193 = Array((new BigNumber(6)).toNumber());
          for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw193.length); _i0_33++) {
            _nw193[_i0_33] = _init33(new BigNumber(_i0_33));
          }
          _1199_v104 = _nw193;
          let _index223 = _module.__default.safeIndex(new BigNumber(284), new BigNumber((_1199_v104).length));
          (_1199_v104)[_index223] = (_1052_v5).f28;
          let _index224 = _module.__default.safeIndex(new BigNumber(284), new BigNumber((_1199_v104).length));
          (_1199_v104)[_index224] = true;
          let _index225 = _module.__default.safeIndex(new BigNumber(284), new BigNumber((_1199_v104).length));
          (_1199_v104)[_index225] = !(!_dafny.Seq.contains(_1054_v7, _1179_v99));
          let _1202_v105;
          _1202_v105 = _dafny.Map.Empty.slice().updateUnsafe(_1199_v104,_1198_cf35);
          _1202_v105 = _1202_v105;
        } else if (_source9.is_DC22) {
          let _1203___mcc_h9 = (_source9).cf38;
          let _1204_cf38 = _1203___mcc_h9;
          let _1205_v106;
          let _init34 = ((_1206_v5) => function (_1207_i18) {
            return (_1206_v5).f28;
          })(_1052_v5);
          let _nw194 = Array((new BigNumber(8)).toNumber());
          for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw194.length); _i0_34++) {
            _nw194[_i0_34] = _init34(new BigNumber(_i0_34));
          }
          _1205_v106 = _nw194;
          let _index226 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_1205_v106).length));
          (_1205_v106)[_index226] = _1048_v1;
          let _index227 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_1205_v106).length));
          (_1205_v106)[_index227] = _1204_cf38;
          (globalState).f11 = new _dafny.CodePoint('l'.codePointAt(0));
          let _1208_v107;
          let _nw195 = new _module.C0();
          _nw195.__ctor();
          _1208_v107 = _nw195;
          _1051_v4 = (_1051_v4).update(!(true), _dafny.Seq.contains(_1052_v5.f29, _1131_i7));
        } else if (_source9.is_DC19) {
          let _1209___mcc_h10 = (_source9).cf32;
          let _1210_cf32 = _1209___mcc_h10;
          let _1211_v108;
          _1211_v108 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("vflk"),(_1052_v5).f28);
          let _1212_v109;
          let _nw196 = Array((new BigNumber(21)).toNumber());
          _nw196[(_dafny.ZERO).toNumber()] = _1048_v1;
          _nw196[(_dafny.ONE).toNumber()] = (_1052_v5).f28;
          _nw196[(new BigNumber(2)).toNumber()] = (_1052_v5).f28;
          _nw196[(new BigNumber(3)).toNumber()] = (_1052_v5).f28;
          _nw196[(new BigNumber(4)).toNumber()] = (_1052_v5).f28;
          _nw196[(new BigNumber(5)).toNumber()] = (_1052_v5).f28;
          _nw196[(new BigNumber(6)).toNumber()] = _1048_v1;
          _nw196[(new BigNumber(7)).toNumber()] = (_1052_v5).f28;
          _nw196[(new BigNumber(8)).toNumber()] = (_1052_v5).f28;
          _nw196[(new BigNumber(9)).toNumber()] = (_1052_v5).f28;
          _nw196[(new BigNumber(10)).toNumber()] = (_module.D8.create_DC20((_1052_v5).f28, (_this).f23)).dtor_cf33;
          _nw196[(new BigNumber(11)).toNumber()] = false;
          _nw196[(new BigNumber(12)).toNumber()] = _1048_v1;
          _nw196[(new BigNumber(13)).toNumber()] = true;
          _nw196[(new BigNumber(14)).toNumber()] = _1048_v1;
          _nw196[(new BigNumber(15)).toNumber()] = (_1052_v5).f28;
          _nw196[(new BigNumber(16)).toNumber()] = (((_1211_v108).contains(_1054_v7)) ? ((_1211_v108).get(_1054_v7)) : ((_1052_v5).f28));
          _nw196[(new BigNumber(17)).toNumber()] = _module.__default.fm1(_1131_i7, _1048_v1, globalState);
          _nw196[(new BigNumber(18)).toNumber()] = _1048_v1;
          _nw196[(new BigNumber(19)).toNumber()] = (_1052_v5).f28;
          _nw196[(new BigNumber(20)).toNumber()] = !((_1052_v5).f28);
          _1212_v109 = _nw196;
          let _1213_v110;
          _1213_v110 = _dafny.Map.Empty.slice().updateUnsafe(_1212_v109,_1179_v99);
          (globalState).f11 = (((_1213_v110).contains(_1212_v109)) ? ((_1213_v110).get(_1212_v109)) : (_1179_v99));
          let _1214_v111;
          let _1215_v112;
          let _1216_v113;
          let _out40;
          let _out41;
          let _out42;
          let _outcollector13 = _module.__default.m0(_1212_v109, _dafny.Seq.Concat(_1054_v7, _1054_v7), _1131_i7, globalState);
          _out40 = _outcollector13[0];
          _out41 = _outcollector13[1];
          _out42 = _outcollector13[2];
          _1214_v111 = _out40;
          _1215_v112 = _out41;
          _1216_v113 = _out42;
          (globalState).f14 = ((_this).f25).minus(_1131_i7);
          let _1217_v114;
          let _init35 = ((_1218_v99) => function (_1219_i19) {
            return _1218_v99;
          })(_1179_v99);
          let _nw197 = Array((new BigNumber(5)).toNumber());
          for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw197.length); _i0_35++) {
            _nw197[_i0_35] = _init35(new BigNumber(_i0_35));
          }
          _1217_v114 = _nw197;
          let _1220_v115;
          _1220_v115 = _dafny.Seq.of((_1052_v5).f28);
          let _1221_v116;
          _1221_v116 = _module.D3.create_DC12(_1220_v115);
          let _1222_v117;
          _1222_v117 = _dafny.Map.Empty.slice().updateUnsafe(_module.D9.create_DC24(_1217_v114),_1221_v116);
          _1222_v117 = _1222_v117;
        } else {
          let _1223___mcc_h11 = (_source9).cf39;
          let _1224_cf39 = _1223___mcc_h11;
          let _1225_v118;
          _1225_v118 = _dafny.Map.Empty.slice().updateUnsafe((_1052_v5).f28,(_1052_v5).f23);
          let _1226_v119;
          _1226_v119 = _dafny.Seq.of(_1225_v118);
          _1051_v4 = (_1051_v4).update((_1134_v66).fm10(_1054_v7, _1048_v1, (_dafny.Map.Empty.slice().updateUnsafe((_1134_v66).fm10(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-438)), ((_1227_v99) => function (_1228_i20) {
            return _1227_v99;
          })(_1179_v99)), (_1052_v5).f28, _1225_v118, true, globalState),new BigNumber(922))).update((_1052_v5).f28, _module.__default.fm0(new BigNumber(((_1226_v119)[_module.__default.safeIndex((_this).f25, new BigNumber((_1226_v119).length))]).length), (_1052_v5).f28, globalState)), (_1052_v5).f28, globalState), _1048_v1);
          let _1229_v120;
          let _nw198 = Array((new BigNumber(20)).toNumber()).fill([]);
          _1229_v120 = _nw198;
          let _1230_v121;
          let _nw199 = Array((new BigNumber(25)).toNumber()).fill(false);
          _1230_v121 = _nw199;
          let _index228 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_1229_v120).length));
          (_1229_v120)[_index228] = _1230_v121;
          let _1231_v122;
          _1231_v122 = _dafny.MultiSet.fromElements(_1179_v99);
          let _1232_v123;
          _1232_v123 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,(_1052_v5).f28);
          let _index229 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_1229_v120).length));
          let _rhs235 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt(new BigNumber((_1055_v8).length), (_this).f25), new BigNumber(491));
          let _rhs236 = (((_1051_v4).contains(!((_1052_v5).f28))) ? ((_1051_v4).get(!((_1052_v5).f28))) : (_dafny.Seq.contains(_dafny.Seq.update(_1054_v7, _module.__default.safeIndex((_1052_v5).f23, new BigNumber((_1054_v7).length)), _1179_v99), _1179_v99)));
          let _rhs237 = _1230_v121;
          let _rhs238 = (((_1231_v122).contains(_1179_v99)) ? ((_1231_v122).get(_1179_v99)) : (new BigNumber(645)));
          let _rhs239 = (((_1232_v123).update(new BigNumber(-753), (_1052_v5).f28)).Merge((_dafny.Map.Empty.slice().updateUnsafe(_1131_i7,(_1052_v5).f28)).update((_this).f25, _1048_v1))).update(_module.__default.safeDivisionInt(new BigNumber((function () {
            let _coll45 = new _dafny.Map();
            for (const _compr_45 of _dafny.IntegerRange(new BigNumber(-391), new BigNumber(524))) {
              let _1233_v124 = _compr_45;
              if (((new BigNumber(-391)).isLessThanOrEqualTo(_1233_v124)) && ((_1233_v124).isLessThan(new BigNumber(524)))) {
                _coll45.push([_module.__default.safeModuloInt(_1233_v124, new BigNumber(909)),(_dafny.ZERO).minus((_1052_v5).f23)]);
              }
            }
            return _coll45;
          }()).length), (_1052_v5).f23), true);
          let _lhs213 = globalState;
          let _lhs214 = _1229_v120;
          let _lhs215 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_1229_v120).length));
          let _lhs216 = globalState;
          _lhs213.f14 = _rhs235;
          _1048_v1 = _rhs236;
          _lhs214[_lhs215] = _rhs237;
          _lhs216.f14 = _rhs238;
          r2 = _rhs239;
          _1232_v123 = (_1232_v123).update(_1131_i7, (_1052_v5).f28);
          let _1234_v125;
          _1234_v125 = _dafny.Seq.of((_1052_v5).f28);
          let _1235_v126;
          _1235_v126 = _dafny.Map.Empty.slice().updateUnsafe(_1234_v125,((_this).f23).isLessThan((_this).f23));
          let _rhs240 = _1054_v7;
          let _rhs241 = (_1052_v5).f28;
          let _rhs242 = (_1235_v126).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1234_v125,(_1052_v5).f28));
          let _rhs243 = !((_1052_v5).f28);
          let _rhs244 = ((new BigNumber(-332)).plus((_this).f23)).plus((_1052_v5).f23);
          let _lhs217 = globalState;
          let _lhs218 = globalState;
          let _lhs219 = globalState;
          _lhs217.f6 = _rhs240;
          r0 = _rhs241;
          _1235_v126 = _rhs242;
          _lhs218.f7 = _rhs243;
          _lhs219.f8 = _rhs244;
        }
      }
      let _hi14 = _module.__default.safeDivisionInt((_1052_v5).f23, (_this).f26);
      for (let _1236_i21 = (_dafny.ZERO).minus((_1052_v5).f23); _1236_i21.isLessThan(_hi14); _1236_i21 = _1236_i21.plus(_dafny.ONE)) {
        (globalState).f14 = _module.__default.fm0((_this).f25, (((_1051_v4).contains((_1052_v5).f28)) ? ((_1051_v4).get((_1052_v5).f28)) : ((_1052_v5).f28)), globalState);
        let _1237_v127;
        _1237_v127 = _module.D12.create_DC35();
        let _1238_v128;
        let _nw200 = Array((new BigNumber(5)).toNumber()).fill(false);
        _1238_v128 = _nw200;
        let _1239_v129;
        _1239_v129 = _dafny.Map.Empty.slice().updateUnsafe(_1237_v127,_1238_v128);
        let _1240_v130;
        let _1241_v131;
        let _out43;
        let _out44;
        let _outcollector14 = (_1052_v5).m1((((_1239_v129).contains(_1237_v127)) ? ((_1239_v129).get(_1237_v127)) : (_1238_v128)), ((_1052_v5).f28) && ((_1052_v5).f28), globalState);
        _out43 = _outcollector14[0];
        _out44 = _outcollector14[1];
        _1240_v130 = _out43;
        _1241_v131 = _out44;
        let _1242_v132;
        let _init36 = ((_1243_v130) => function (_1244_i22) {
          return _module.__default.safeDivisionInt(_1244_i22, _1243_v130);
        })(_1240_v130);
        let _nw201 = Array((new BigNumber(15)).toNumber());
        for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw201.length); _i0_36++) {
          _nw201[_i0_36] = _init36(new BigNumber(_i0_36));
        }
        _1242_v132 = _nw201;
        let _1245_v133;
        _1245_v133 = _dafny.Seq.of(_1242_v132);
        let _1246_v134;
        _1246_v134 = _module.D1.create_DC3(_1245_v133);
        let _source10 = _1246_v134;
        if (_source10.is_DC4) {
          let _1247___mcc_h12 = (_source10).cf6;
          let _1248___mcc_h13 = (_source10).cf7;
          let _1249___mcc_h14 = (_source10).cf8;
          let _1250_cf8 = _1249___mcc_h14;
          let _1251_cf7 = _1248___mcc_h13;
          let _1252_cf6 = _1247___mcc_h12;
          let _1253_v135;
          let _nw202 = new _module.C3();
          _nw202.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(67)), function (_1254_i23) {
            return _dafny.MultiSet.fromElements(true, true);
          }), (_1052_v5).f23, (_1052_v5).f28, _1052_v5.f29);
          _1253_v135 = _nw202;
          (globalState).f14 = (_this).f23;
          let _1255_v136;
          let _nw203 = new _module.C4();
          _nw203.__ctor(_1242_v132, _1236_i21);
          _1255_v136 = _nw203;
          let _1256_v137;
          _1256_v137 = _module.D3.create_DC12(_dafny.Seq.of(_1048_v1));
          let _1257_v138;
          _1257_v138 = _dafny.Map.Empty.slice().updateUnsafe(_1256_v137,(_1052_v5).f28);
          let _1258_v139;
          _1258_v139 = _dafny.Seq.of(!((_1052_v5).f28), !(_1250_cf8));
          let _pat_let_tv39 = _1258_v139;
          let _pat_let_tv40 = _1258_v139;
          (globalState).f21 = (((_1257_v138).contains(function (_pat_let22_0) {
            return function (_1261_dt__update__tmp_h1) {
              return function (_pat_let23_0) {
                return function (_1262_dt__update_hcf23_h0) {
                  return _module.D3.create_DC12(_1262_dt__update_hcf23_h0);
                }(_pat_let23_0);
              }(_pat_let_tv40);
            }(_pat_let22_0);
          }(_1256_v137))) ? ((_1257_v138).get(function (_pat_let20_0) {
            return function (_1259_dt__update__tmp_h2) {
              return function (_pat_let21_0) {
                return function (_1260_dt__update_hcf23_h1) {
                  return _module.D3.create_DC12(_1260_dt__update_hcf23_h1);
                }(_pat_let21_0);
              }(_pat_let_tv39);
            }(_pat_let20_0);
          }(_1256_v137))) : (true));
        } else if (_source10.is_DC5) {
          let _1263___mcc_h15 = (_source10).cf9;
          let _1264___mcc_h16 = (_source10).cf10;
          let _1265___mcc_h17 = (_source10).cf11;
          let _1266___mcc_h18 = (_source10).cf12;
          let _1267_cf12 = _1266___mcc_h18;
          let _1268_cf11 = _1265___mcc_h17;
          let _1269_cf10 = _1264___mcc_h16;
          let _1270_cf9 = _1263___mcc_h15;
          (globalState).f14 = (_1236_i21).minus((_1052_v5).f23);
          let _1271_v140;
          _1271_v140 = _dafny.Seq.of(_1048_v1, (_1052_v5).f28);
          (globalState).f21 = ((_dafny.Seq.IsPrefixOf(_1271_v140, _1271_v140)) ? (_1269_cf10) : (((_this).f26).isLessThanOrEqualTo(_1236_i21)));
          _1055_v8 = _1055_v8;
          let _1272_v141;
          let _nw204 = new _module.C0();
          _nw204.__ctor();
          _1272_v141 = _nw204;
        } else if (_source10.is_DC6) {
          let _1273___mcc_h19 = (_source10).cf13;
          let _1274_cf13 = _1273___mcc_h19;
          (globalState).f8 = (_dafny.ZERO).minus((_this).f23);
          (globalState).f9 = _1048_v1;
          _1274_cf13 = (((_1052_v5).f28) ? (_1274_cf13) : (_1274_cf13));
          let _1275_v142;
          _1275_v142 = _dafny.Map.Empty.slice().updateUnsafe(_1052_v5.f29,_1236_i21);
          _1275_v142 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.of(_module.__default.fm0(new BigNumber(455), _1048_v1, globalState), (_this).f25), _dafny.Seq.of((_1052_v5).f23)),new BigNumber(993));
        } else {
          let _1276___mcc_h20 = (_source10).cf5;
          let _1277_cf5 = _1276___mcc_h20;
          let _1278_v143;
          _1278_v143 = _dafny.Map.Empty.slice().updateUnsafe((_1052_v5).f28,_1240_v130);
          let _1279_v144;
          _1279_v144 = _dafny.Seq.of((_1052_v5).f28);
          _1278_v143 = (_1278_v143).update(_1048_v1, new BigNumber((_1279_v144).length));
          let _1280_v145;
          _1280_v145 = _dafny.MultiSet.fromElements((_1052_v5).f28);
          (globalState).f14 = _module.__default.safeDivisionInt(new BigNumber((_1280_v145).cardinality()), ((_dafny.ZERO).minus((_this).f23)).plus(new BigNumber((_dafny.MultiSet.fromElements((_1052_v5).f23)).cardinality())));
          let _1281_v146;
          let _init37 = ((_1282_v7) => function (_1283_i24) {
            return _1282_v7;
          })(_1054_v7);
          let _nw205 = Array((new BigNumber(13)).toNumber());
          for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw205.length); _i0_37++) {
            _nw205[_i0_37] = _init37(new BigNumber(_i0_37));
          }
          _1281_v146 = _nw205;
          let _index230 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_1281_v146).length));
          (_1281_v146)[_index230] = _1054_v7;
          let _index231 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_1281_v146).length));
          (_1281_v146)[_index231] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-84)), function (_1284_i25) {
            return new _dafny.CodePoint('x'.codePointAt(0));
          }), _1054_v7);
          _1054_v7 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1054_v7, (_1281_v146)[_module.__default.safeIndex(new BigNumber(404), new BigNumber((_1281_v146).length))]), _1054_v7);
        }
        if (_1048_v1) {
          let _1285_v147;
          let _init38 = ((_1286_v131) => function (_1287_i26) {
            return _1286_v131;
          })(_1241_v131);
          let _nw206 = Array((new BigNumber(18)).toNumber());
          for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw206.length); _i0_38++) {
            _nw206[_i0_38] = _init38(new BigNumber(_i0_38));
          }
          _1285_v147 = _nw206;
          let _index232 = _module.__default.safeIndex(new BigNumber(867), new BigNumber((_1285_v147).length));
          (_1285_v147)[_index232] = function () {
            let _coll46 = new _dafny.Set();
            for (const _compr_46 of (function () {
              let _coll47 = new _dafny.Set();
              for (const _compr_47 of _dafny.IntegerRange(new BigNumber(836), new BigNumber(653))) {
                let _1288_v148 = _compr_47;
                if (((new BigNumber(836)).isLessThanOrEqualTo(_1288_v148)) && ((_1288_v148).isLessThan(new BigNumber(653)))) {
                  _coll47.add((_1288_v148).multipliedBy(new BigNumber(25)));
                }
              }
              return _coll47;
            }()).Elements) {
              let _1289_v149 = _compr_46;
              if ((function () {
                let _coll48 = new _dafny.Set();
                for (const _compr_48 of _dafny.IntegerRange(new BigNumber(836), new BigNumber(653))) {
                  let _1290_v148 = _compr_48;
                  if (((new BigNumber(836)).isLessThanOrEqualTo(_1290_v148)) && ((_1290_v148).isLessThan(new BigNumber(653)))) {
                    _coll48.add((_1290_v148).multipliedBy(new BigNumber(25)));
                  }
                }
                return _coll48;
              }()).contains(_1289_v149)) {
                _coll46.add(_module.__default.safeModuloInt(_1289_v149, new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('b'.codePointAt(0)), new _dafny.CodePoint('e'.codePointAt(0)))).cardinality())));
              }
            }
            return _coll46;
          }();
          let _index233 = _module.__default.safeIndex(new BigNumber(867), new BigNumber((_1285_v147).length));
          (_1285_v147)[_index233] = _1241_v131;
          let _1291_v150;
          let _1292_v151;
          let _out45;
          let _out46;
          let _outcollector15 = (_1052_v5).m2(_dafny.Seq.Concat(_1054_v7, _1054_v7), globalState);
          _out45 = _outcollector15[0];
          _out46 = _outcollector15[1];
          _1291_v150 = _out45;
          _1292_v151 = _out46;
          let _1293_v152;
          _1293_v152 = _module.D0.create_DC1((_dafny.ZERO).minus(_1240_v130), _1048_v1, (_this).f26);
          let _rhs245 = (_1052_v5).f23;
          let _rhs246 = (_1052_v5).f28;
          let _rhs247 = (_1293_v152).dtor_cf2;
          let _lhs220 = globalState;
          let _lhs221 = globalState;
          _lhs220.f14 = _rhs245;
          _1048_v1 = _rhs246;
          _lhs221.f21 = _rhs247;
          let _1294_v153;
          _1294_v153 = _dafny.Map.Empty.slice().updateUnsafe((_1052_v5).f23,(_1052_v5).f28);
          let _1295_v154;
          _1295_v154 = _dafny.Map.Empty.slice().updateUnsafe(_1048_v1,_1240_v130);
          let _index234 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_1238_v128).length));
          (_1238_v128)[_index234] = !(((_dafny.ZERO).minus((((_1295_v154).contains((_1052_v5).f28)) ? ((_1295_v154).get((_1052_v5).f28)) : ((_1052_v5).f23)))).isLessThan(new BigNumber((_1294_v153).length)));
          let _1296_v155;
          _1296_v155 = new _dafny.CodePoint('u'.codePointAt(0));
          let _1297_v156;
          _1297_v156 = _module.D11.create_DC32(_dafny.Seq.UnicodeFromString("mxdqyef"));
          let _1298_v157;
          let _nw207 = Array((new BigNumber(26)).toNumber());
          _nw207[(_dafny.ZERO).toNumber()] = _1054_v7;
          _nw207[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_1054_v7, _1054_v7);
          _nw207[(new BigNumber(2)).toNumber()] = _1054_v7;
          _nw207[(new BigNumber(3)).toNumber()] = (_1055_v8)[_module.__default.safeIndex(_module.__default.fm0((_1052_v5).f23, (_1052_v5).f28, globalState), new BigNumber((_1055_v8).length))];
          _nw207[(new BigNumber(4)).toNumber()] = _dafny.Seq.update(_1054_v7, _module.__default.safeIndex(new BigNumber(((_1055_v8)[_module.__default.safeIndex((_this).f23, new BigNumber((_1055_v8).length))]).length), new BigNumber((_1054_v7).length)), _1296_v155);
          _nw207[(new BigNumber(5)).toNumber()] = _1054_v7;
          _nw207[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_1054_v7, _dafny.Seq.update(_1054_v7, _module.__default.safeIndex(_1291_v150, new BigNumber((_1054_v7).length)), _1296_v155));
          _nw207[(new BigNumber(7)).toNumber()] = _1054_v7;
          _nw207[(new BigNumber(8)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(525)), function (_1299_i27) {
            return new _dafny.CodePoint('n'.codePointAt(0));
          });
          _nw207[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("tfynwpxj"), _1054_v7);
          _nw207[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_1054_v7, _1054_v7);
          _nw207[(new BigNumber(11)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(62)), ((_1300_v155) => function (_1301_i28) {
            return _1300_v155;
          })(_1296_v155));
          _nw207[(new BigNumber(12)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-834)), ((_1302_v155) => function (_1303_i29) {
            return _1302_v155;
          })(_1296_v155));
          _nw207[(new BigNumber(13)).toNumber()] = (_1297_v156).dtor_cf53;
          _nw207[(new BigNumber(14)).toNumber()] = _dafny.Seq.UnicodeFromString("u");
          _nw207[(new BigNumber(15)).toNumber()] = _module.__default.fm2(globalState);
          _nw207[(new BigNumber(16)).toNumber()] = _1054_v7;
          _nw207[(new BigNumber(17)).toNumber()] = _dafny.Seq.Concat(_1054_v7, _1054_v7);
          _nw207[(new BigNumber(18)).toNumber()] = _1054_v7;
          _nw207[(new BigNumber(19)).toNumber()] = _dafny.Seq.UnicodeFromString("g");
          _nw207[(new BigNumber(20)).toNumber()] = _1054_v7;
          _nw207[(new BigNumber(21)).toNumber()] = _1054_v7;
          _nw207[(new BigNumber(22)).toNumber()] = _1054_v7;
          _nw207[(new BigNumber(23)).toNumber()] = _1054_v7;
          _nw207[(new BigNumber(24)).toNumber()] = _dafny.Seq.UnicodeFromString("esqnlyvg");
          _nw207[(new BigNumber(25)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("tkq"), _1054_v7);
          _1298_v157 = _nw207;
          let _index235 = _module.__default.safeIndex(new BigNumber(881), new BigNumber((_1298_v157).length));
          (_1298_v157)[_index235] = (((_1052_v5).f28) ? (_dafny.Seq.UnicodeFromString("nxukkyffw")) : (_1054_v7));
          let _1304_v158;
          let _nw208 = Array((_dafny.ONE).toNumber());
          _nw208[(_dafny.ZERO).toNumber()] = _1047_v0;
          _1304_v158 = _nw208;
          let _1305_v159;
          _1305_v159 = _module.D11.create_DC30(_1304_v158);
          let _1306_v160;
          _1306_v160 = _dafny.Map.Empty.slice().updateUnsafe(((true) ? (_1296_v155) : (new _dafny.CodePoint('e'.codePointAt(0)))),(_1052_v5).f28);
          let _1307_v161;
          _1307_v161 = _dafny.Map.Empty.slice().updateUnsafe(_1054_v7,_1054_v7);
          let _index236 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_1238_v128).length));
          let _index237 = _module.__default.safeIndex(new BigNumber(881), new BigNumber((_1298_v157).length));
          let _rhs248 = !(((_1052_v5).f23).isLessThan((_1052_v5).f23));
          let _rhs249 = _dafny.Seq.Concat(_dafny.Seq.Concat((((_1307_v161).contains(_dafny.Seq.UnicodeFromString("culpysq"))) ? ((_1307_v161).get(_dafny.Seq.UnicodeFromString("culpysq"))) : (_1054_v7)), _1054_v7), _dafny.Seq.Concat(_1054_v7, _1054_v7));
          let _rhs250 = _1305_v159;
          let _rhs251 = (_1306_v160).Merge(_1306_v160);
          let _lhs222 = _1238_v128;
          let _lhs223 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_1238_v128).length));
          let _lhs224 = _1298_v157;
          let _lhs225 = _module.__default.safeIndex(new BigNumber(881), new BigNumber((_1298_v157).length));
          _lhs222[_lhs223] = _rhs248;
          _lhs224[_lhs225] = _rhs249;
          _1305_v159 = _rhs250;
          _1306_v160 = _rhs251;
          _1054_v7 = (_1298_v157)[_module.__default.safeIndex(new BigNumber(881), new BigNumber((_1298_v157).length))];
        } else {
          _1055_v8 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(217)), function (_1308_i30) {
            return _dafny.Seq.UnicodeFromString("tp");
          });
          (globalState).f8 = ((_this).f23).plus((_1052_v5).f23);
          let _1309_v162;
          _1309_v162 = _dafny.Map.Empty.slice().updateUnsafe((_1052_v5).f28,(_this).f26);
          let _1310_v163;
          let _nw209 = new _module.C1();
          _nw209.__ctor((_1052_v5).f28, _dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(_1309_v162)).cardinality())), (_1052_v5).f23);
          _1310_v163 = _nw209;
          let _index238 = _module.__default.safeIndex(new BigNumber(872), new BigNumber((_1242_v132).length));
          (_1242_v132)[_index238] = (_this).f25;
          let _1311_v164;
          _1311_v164 = new _dafny.CodePoint('i'.codePointAt(0));
          let _1312_v165;
          _1312_v165 = _dafny.MultiSet.fromElements(_1311_v164, _1311_v164);
          let _index239 = _module.__default.safeIndex(new BigNumber(872), new BigNumber((_1242_v132).length));
          (_1242_v132)[_index239] = (((_this).f26).plus((((_1312_v165).contains(_1311_v164)) ? ((_1312_v165).get(_1311_v164)) : ((_this).f26)))).plus((_this).f26);
          let _1313_v166;
          _1313_v166 = _module.D11.create_DC31(_1236_i21);
          let _1314_v167;
          _1314_v167 = _dafny.Seq.of(_1313_v166, _1313_v166, _1313_v166);
          let _1315_v168;
          _1315_v168 = _dafny.Map.Empty.slice().updateUnsafe(_1314_v167,_1052_v5.f29);
          let _1316_v170;
          _1316_v170 = _dafny.Set.fromElements(_module.__default.fm32(globalState));
          let _1317_v172;
          _1317_v172 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_1313_v166), _1314_v167, _dafny.Seq.update(_dafny.Seq.of(_1313_v166), _module.__default.safeIndex((_1052_v5).f23, new BigNumber((_dafny.Seq.of(_1313_v166)).length)), _1313_v166), _1314_v167);
          let _1318_v173;
          let _nw210 = Array((new BigNumber(26)).toNumber());
          _nw210[(_dafny.ZERO).toNumber()] = (_1315_v168).update(_1314_v167, _1052_v5.f29);
          _nw210[(_dafny.ONE).toNumber()] = _1315_v168;
          _nw210[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1314_v167,_dafny.Seq.Create(_module.__default.abs(new BigNumber(428)), ((_1319_v5, _1320_v130, _1321_i21, _1322_v132) => function (_1323_i31) {
            return new BigNumber((_dafny.Seq.of(_dafny.MultiSet.fromElements((_this).f26), _dafny.MultiSet.FromArray(_1319_v5.f29), _dafny.MultiSet.fromElements(_1320_v130, new BigNumber(383)), _dafny.MultiSet.fromElements(_1321_i21, (_1322_v132)[_module.__default.safeIndex(new BigNumber(872), new BigNumber((_1322_v132).length))]))).length);
          })(_1052_v5, _1240_v130, _1236_i21, _1242_v132)));
          _nw210[(new BigNumber(3)).toNumber()] = (_1315_v168).Merge((_module.D15.create_DC43(_1315_v168)).dtor_cf67);
          _nw210[(new BigNumber(4)).toNumber()] = function () {
            let _coll49 = new _dafny.Map();
            for (const _compr_49 of (_1316_v170).Elements) {
              let _1324_v169 = _compr_49;
              if ((_1316_v170).contains(_1324_v169)) {
                _coll49.push([_1324_v169,_1052_v5.f29]);
              }
            }
            return _coll49;
          }();
          _nw210[(new BigNumber(5)).toNumber()] = _1315_v168;
          _nw210[(new BigNumber(6)).toNumber()] = _1315_v168;
          _nw210[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1314_v167,_1052_v5.f29);
          _nw210[(new BigNumber(8)).toNumber()] = _1315_v168;
          _nw210[(new BigNumber(9)).toNumber()] = _1315_v168;
          _nw210[(new BigNumber(10)).toNumber()] = _1315_v168;
          _nw210[(new BigNumber(11)).toNumber()] = _1315_v168;
          _nw210[(new BigNumber(12)).toNumber()] = function () {
            let _coll50 = new _dafny.Map();
            for (const _compr_50 of (_1317_v172).Elements) {
              let _1325_v171 = _compr_50;
              if ((_1317_v172).contains(_1325_v171)) {
                _coll50.push([_1325_v171,_1050_v3]);
              }
            }
            return _coll50;
          }();
          _nw210[(new BigNumber(13)).toNumber()] = _1315_v168;
          _nw210[(new BigNumber(14)).toNumber()] = (((_1052_v5).f28) ? (_1315_v168) : (_1315_v168));
          _nw210[(new BigNumber(15)).toNumber()] = _module.__default.fm33(new BigNumber((_dafny.Set.fromElements(new BigNumber((_1052_v5.f29).length), _1236_i21, (_dafny.ZERO).minus((_this).f23))).length), (_1052_v5).f28, globalState);
          _nw210[(new BigNumber(16)).toNumber()] = _1315_v168;
          _nw210[(new BigNumber(17)).toNumber()] = _1315_v168;
          _nw210[(new BigNumber(18)).toNumber()] = _1315_v168;
          _nw210[(new BigNumber(19)).toNumber()] = _1315_v168;
          _nw210[(new BigNumber(20)).toNumber()] = _1315_v168;
          _nw210[(new BigNumber(21)).toNumber()] = (_1315_v168).Merge(_1315_v168);
          _nw210[(new BigNumber(22)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1314_v167,_1050_v3);
          _nw210[(new BigNumber(23)).toNumber()] = _1315_v168;
          _nw210[(new BigNumber(24)).toNumber()] = (_1315_v168).Merge((_1315_v168).update(_1314_v167, _1052_v5.f29));
          _nw210[(new BigNumber(25)).toNumber()] = _1315_v168;
          _1318_v173 = _nw210;
          let _index240 = _module.__default.safeIndex(new BigNumber(41), new BigNumber((_1318_v173).length));
          (_1318_v173)[_index240] = _1315_v168;
          let _1326_v174;
          _1326_v174 = _dafny.Seq.of((_1052_v5).f28);
          let _index241 = _module.__default.safeIndex(new BigNumber(41), new BigNumber((_1318_v173).length));
          (_1318_v173)[_index241] = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(836)), ((_1327_v166) => function (_1328_i32) {
            return _1327_v166;
          })(_1313_v166)), _module.__default.safeIndex(new BigNumber((_1326_v174).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(836)), ((_1329_v166) => function (_1330_i32) {
            return _1329_v166;
          })(_1313_v166))).length)), _1313_v166),_dafny.Seq.Concat(_1052_v5.f29, _dafny.Seq.of((_1242_v132)[_module.__default.safeIndex(new BigNumber(872), new BigNumber((_1242_v132).length))])));
        }
      }
      let _1331_v175;
      _1331_v175 = _dafny.Set.fromElements((((_1052_v5).f28) ? (_1054_v7) : (_dafny.Seq.UnicodeFromString("seef"))));
      let _1332_v176;
      _1332_v176 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_1331_v175);
      let _1333_v177;
      _1333_v177 = _module.D16.create_DC46(_1331_v175);
      _1331_v175 = (((_1332_v176).contains((_this).f26)) ? ((_1332_v176).get((_this).f26)) : ((((_1052_v5).f28) ? (_1331_v175) : ((_1333_v177).dtor_cf74))));
      let _1334_v178;
      _1334_v178 = _dafny.Map.Empty.slice().updateUnsafe(_1054_v7,(_1052_v5).f28);
      r0 = !((((_1334_v178).contains(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ulwvqx"), _1054_v7))) ? ((_1334_v178).get(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ulwvqx"), _1054_v7))) : ((_1052_v5).f28)));
      let _1335_v179;
      _1335_v179 = _dafny.Map.Empty.slice().updateUnsafe(_1048_v1,(_this).f23);
      let _1336_v180;
      _1336_v180 = _dafny.Seq.of(_1335_v179, _1335_v179);
      let _1337_v181;
      _1337_v181 = _dafny.Set.fromElements(new BigNumber((_1055_v8).length));
      let _1338_v182;
      _1338_v182 = _dafny.Set.fromElements(_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_1337_v181).length)), new BigNumber(957), new BigNumber(843), (_this).f26));
      r1 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_1336_v180, _dafny.Seq.Create(_module.__default.abs(new BigNumber(681)), ((_1339_v179) => function (_1340_i33) {
        return _1339_v179;
      })(_1335_v179))), _dafny.Seq.Concat(_1336_v180, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_1052_v5).f28,(_this).f25), _1335_v179, _1335_v179, _1335_v179))), _module.__default.safeIndex(new BigNumber((_1338_v182).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_1336_v180, _dafny.Seq.Create(_module.__default.abs(new BigNumber(681)), ((_1341_v179) => function (_1342_i33) {
        return _1341_v179;
      })(_1335_v179))), _dafny.Seq.Concat(_1336_v180, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_1052_v5).f28,(_this).f25), _1335_v179, _1335_v179, _1335_v179)))).length)), _1335_v179);
      r2 = _module.__default.fm5(_1337_v181, globalState);
      return [r0, r1, r2];
    }
    m4(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      (globalState).f7 = p0;
      let _1343_v0;
      _1343_v0 = _module.D11.create_DC31((_this).f25);
      let _1344_v1;
      _1344_v1 = _dafny.Seq.of(_1343_v0, _1343_v0);
      let _1345_v2;
      _1345_v2 = _dafny.Seq.UnicodeFromString("fvpo");
      let _1346_v3;
      _1346_v3 = _dafny.Seq.of(new BigNumber((_1345_v2).length), (_this).f25);
      let _1347_v4;
      _1347_v4 = _dafny.Map.Empty.slice().updateUnsafe(_1344_v1,_1346_v3);
      let _1348_v5;
      _1348_v5 = _module.D15.create_DC43((_1347_v4).Merge(_1347_v4));
      let _source11 = _1348_v5;
      if (_source11.is_DC44) {
        let _1349___mcc_h0 = (_source11).cf68;
        let _1350___mcc_h1 = (_source11).cf69;
        let _1351___mcc_h2 = (_source11).cf70;
        let _1352___mcc_h3 = (_source11).cf71;
        let _1353___mcc_h4 = (_source11).cf72;
        let _1354_cf72 = _1353___mcc_h4;
        let _1355_cf71 = _1352___mcc_h3;
        let _1356_cf70 = _1351___mcc_h2;
        let _1357_cf69 = _1350___mcc_h1;
        let _1358_cf68 = _1349___mcc_h0;
        let _1359_v6;
        let _nw211 = Array((new BigNumber(5)).toNumber()).fill([]);
        _1359_v6 = _nw211;
        let _1360_v7;
        _1360_v7 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('u'.codePointAt(0)),p0);
        let _1361_v8;
        _1361_v8 = _module.D12.create_DC34(_1360_v7);
        let _pat_let_tv41 = _1360_v7;
        let _pat_let_tv42 = _1360_v7;
        let _pat_let_tv43 = _1360_v7;
        let _1362_v9;
        let _nw212 = Array((new BigNumber(21)).toNumber());
        _nw212[(_dafny.ZERO).toNumber()] = _module.D12.create_DC34(_1360_v7);
        _nw212[(_dafny.ONE).toNumber()] = _module.D12.create_DC34(_1360_v7);
        _nw212[(new BigNumber(2)).toNumber()] = _1361_v8;
        _nw212[(new BigNumber(3)).toNumber()] = _1361_v8;
        _nw212[(new BigNumber(4)).toNumber()] = function (_pat_let24_0) {
          return function (_1363_dt__update__tmp_h0) {
            return function (_pat_let25_0) {
              return function (_1364_dt__update_hcf55_h0) {
                return _module.D12.create_DC34(_1364_dt__update_hcf55_h0);
              }(_pat_let25_0);
            }(_pat_let_tv41);
          }(_pat_let24_0);
        }(_1361_v8);
        _nw212[(new BigNumber(5)).toNumber()] = _1361_v8;
        _nw212[(new BigNumber(6)).toNumber()] = _1361_v8;
        _nw212[(new BigNumber(7)).toNumber()] = function (_pat_let26_0) {
          return function (_1367_dt__update__tmp_h2) {
            return function (_pat_let29_0) {
              return function (_1368_dt__update_hcf55_h2) {
                return _module.D12.create_DC34(_1368_dt__update_hcf55_h2);
              }(_pat_let29_0);
            }(_pat_let_tv43);
          }(_pat_let26_0);
        }(function (_pat_let27_0) {
          return function (_1365_dt__update__tmp_h1) {
            return function (_pat_let28_0) {
              return function (_1366_dt__update_hcf55_h1) {
                return _module.D12.create_DC34(_1366_dt__update_hcf55_h1);
              }(_pat_let28_0);
            }(_pat_let_tv42);
          }(_pat_let27_0);
        }(_1361_v8));
        _nw212[(new BigNumber(8)).toNumber()] = _1361_v8;
        _nw212[(new BigNumber(9)).toNumber()] = _module.D12.create_DC34(_1360_v7);
        _nw212[(new BigNumber(10)).toNumber()] = _1361_v8;
        _nw212[(new BigNumber(11)).toNumber()] = _1361_v8;
        _nw212[(new BigNumber(12)).toNumber()] = _1361_v8;
        _nw212[(new BigNumber(13)).toNumber()] = _1361_v8;
        _nw212[(new BigNumber(14)).toNumber()] = _module.D12.create_DC34(_1360_v7);
        _nw212[(new BigNumber(15)).toNumber()] = _1361_v8;
        _nw212[(new BigNumber(16)).toNumber()] = _1361_v8;
        _nw212[(new BigNumber(17)).toNumber()] = _1361_v8;
        _nw212[(new BigNumber(18)).toNumber()] = _1361_v8;
        _nw212[(new BigNumber(19)).toNumber()] = _1361_v8;
        _nw212[(new BigNumber(20)).toNumber()] = _1361_v8;
        _1362_v9 = _nw212;
        let _index242 = _module.__default.safeIndex(new BigNumber(643), new BigNumber((_1359_v6).length));
        (_1359_v6)[_index242] = _1362_v9;
        let _index243 = _module.__default.safeIndex(new BigNumber(643), new BigNumber((_1359_v6).length));
        (_1359_v6)[_index243] = _1362_v9;
        let _1369_v10;
        _1369_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1354_cf72,_1358_cf68);
        let _1370_v11;
        _1370_v11 = _dafny.Seq.of(_1345_v2);
        let _1371_v12;
        _1371_v12 = _dafny.Seq.of(_1370_v11);
        let _1372_v13;
        _1372_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).f25,_1355_cf71);
        let _1373_v14;
        _1373_v14 = _dafny.Set.fromElements(_1354_cf72);
        let _1374_v15;
        let _nw213 = Array((new BigNumber(27)).toNumber());
        _nw213[(_dafny.ZERO).toNumber()] = new BigNumber((_1369_v10).length);
        _nw213[(_dafny.ONE).toNumber()] = (_this).f26;
        _nw213[(new BigNumber(2)).toNumber()] = new BigNumber(((_1371_v12)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements(!(_1355_cf71))).cardinality()), new BigNumber((_1371_v12).length))]).length);
        _nw213[(new BigNumber(3)).toNumber()] = new BigNumber((_1372_v13).length);
        _nw213[(new BigNumber(4)).toNumber()] = (_this).f25;
        _nw213[(new BigNumber(5)).toNumber()] = (_this).f23;
        _nw213[(new BigNumber(6)).toNumber()] = (_this).f26;
        _nw213[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_1373_v14).length)))).length),(_this).f26)).length);
        _nw213[(new BigNumber(8)).toNumber()] = new BigNumber(353);
        _nw213[(new BigNumber(9)).toNumber()] = (_this).f25;
        _nw213[(new BigNumber(10)).toNumber()] = (_this).f25;
        _nw213[(new BigNumber(11)).toNumber()] = new BigNumber(-646);
        _nw213[(new BigNumber(12)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("rocyxmwas")).length);
        _nw213[(new BigNumber(13)).toNumber()] = _1356_cf70;
        _nw213[(new BigNumber(14)).toNumber()] = new BigNumber((_1372_v13).length);
        _nw213[(new BigNumber(15)).toNumber()] = new BigNumber(390);
        _nw213[(new BigNumber(16)).toNumber()] = (_this).f25;
        _nw213[(new BigNumber(17)).toNumber()] = (_this).f23;
        _nw213[(new BigNumber(18)).toNumber()] = _1356_cf70;
        _nw213[(new BigNumber(19)).toNumber()] = (_this).f23;
        _nw213[(new BigNumber(20)).toNumber()] = (_this).f23;
        _nw213[(new BigNumber(21)).toNumber()] = (_this).f26;
        _nw213[(new BigNumber(22)).toNumber()] = (_this).f25;
        _nw213[(new BigNumber(23)).toNumber()] = (_this).f23;
        _nw213[(new BigNumber(24)).toNumber()] = (_this).f23;
        _nw213[(new BigNumber(25)).toNumber()] = (_this).f23;
        _nw213[(new BigNumber(26)).toNumber()] = (_this).f23;
        _1374_v15 = _nw213;
        let _1375_v16;
        _1375_v16 = _dafny.Map.Empty.slice().updateUnsafe(_1374_v15,_1355_cf71);
        let _1376_v17;
        let _init39 = ((_1377_cf71) => function (_1378_i0) {
          return _1377_cf71;
        })(_1355_cf71);
        let _nw214 = Array((new BigNumber(4)).toNumber());
        for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw214.length); _i0_39++) {
          _nw214[_i0_39] = _init39(new BigNumber(_i0_39));
        }
        _1376_v17 = _nw214;
        let _index244 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_1376_v17).length));
        (_1376_v17)[_index244] = _1355_cf71;
        let _index245 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_1376_v17).length));
        let _rhs252 = _1375_v16;
        let _rhs253 = (((_1369_v10).contains(_1354_cf72)) ? ((_1369_v10).get(_1354_cf72)) : (!(((_this).f25).isLessThanOrEqualTo((_this).f23))));
        let _lhs226 = _1376_v17;
        let _lhs227 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_1376_v17).length));
        _1375_v16 = _rhs252;
        _lhs226[_lhs227] = _rhs253;
        let _1379_v18;
        _1379_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1346_v3,_1358_cf68);
        let _rhs254 = _1354_cf72;
        let _rhs255 = _dafny.Map.Empty.slice().updateUnsafe(_1346_v3,_1354_cf72);
        let _rhs256 = _1356_cf70;
        _1358_cf68 = _rhs254;
        _1379_v18 = _rhs255;
        r1 = _rhs256;
        (globalState).f6 = _module.__default.fm2(globalState);
      } else if (_source11.is_DC43) {
        let _1380___mcc_h5 = (_source11).cf67;
        let _1381_cf67 = _1380___mcc_h5;
        let _1382_v19;
        _1382_v19 = _dafny.MultiSet.fromElements(p0, p0);
        let _1383_v20;
        _1383_v20 = _dafny.Seq.of(_1382_v19);
        let _1384_v21;
        let _nw215 = new _module.C3();
        _nw215.__ctor(_dafny.Seq.Concat(_dafny.Seq.of(_1382_v19), _1383_v20), (new BigNumber(983)).plus((_this).f23), p0, _1346_v3);
        _1384_v21 = _nw215;
        let _1385_v22;
        _1385_v22 = _dafny.Set.fromElements((_1382_v19).update(p0, _module.__default.abs(new BigNumber(190))), _1382_v19);
        if ((_1385_v22).equals(function () {
          let _coll51 = new _dafny.Set();
          for (const _compr_51 of (_dafny.Map.Empty.slice().updateUnsafe(_1382_v19,p0)).Keys.Elements) {
            let _1386_v23 = _compr_51;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_1382_v19,p0)).contains(_1386_v23)) {
              _coll51.add(_1386_v23);
            }
          }
          return _coll51;
        }())) {
          let _1387_v24;
          let _nw216 = Array((new BigNumber(9)).toNumber());
          _nw216[(_dafny.ZERO).toNumber()] = p0;
          _nw216[(_dafny.ONE).toNumber()] = true;
          _nw216[(new BigNumber(2)).toNumber()] = _module.__default.fm1((_this).f26, p0, globalState);
          _nw216[(new BigNumber(3)).toNumber()] = p0;
          _nw216[(new BigNumber(4)).toNumber()] = p0;
          _nw216[(new BigNumber(5)).toNumber()] = p0;
          _nw216[(new BigNumber(6)).toNumber()] = _module.__default.fm1((_this).f25, p0, globalState);
          _nw216[(new BigNumber(7)).toNumber()] = p0;
          _nw216[(new BigNumber(8)).toNumber()] = p0;
          _1387_v24 = _nw216;
          let _1388_v25;
          _1388_v25 = _dafny.Seq.of(_1387_v24, _1387_v24, _1387_v24, _1387_v24);
          _1388_v25 = _1388_v25;
          let _1389_v26;
          _1389_v26 = new _dafny.CodePoint('j'.codePointAt(0));
          let _1390_v27;
          _1390_v27 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f26),_1345_v2);
          let _1391_v28;
          _1391_v28 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _1392_v29;
          _1392_v29 = _dafny.Seq.of(_1345_v2);
          let _1393_v30;
          let _nw217 = Array((new BigNumber(23)).toNumber());
          _nw217[(_dafny.ZERO).toNumber()] = _1345_v2;
          _nw217[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("pfojm"), _dafny.Seq.UnicodeFromString("mmwijrbtt"));
          _nw217[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(172)), function (_1394_i1) {
            return new _dafny.CodePoint('s'.codePointAt(0));
          });
          _nw217[(new BigNumber(3)).toNumber()] = _dafny.Seq.UnicodeFromString("vabaqredp");
          _nw217[(new BigNumber(4)).toNumber()] = _1345_v2;
          _nw217[(new BigNumber(5)).toNumber()] = _dafny.Seq.update(_1345_v2, _module.__default.safeIndex((_this).f25, new BigNumber((_1345_v2).length)), _1389_v26);
          _nw217[(new BigNumber(6)).toNumber()] = (((_1390_v27).contains(new BigNumber((_1391_v28).length))) ? ((_1390_v27).get(new BigNumber((_1391_v28).length))) : (_1345_v2));
          _nw217[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat((_1392_v29)[_module.__default.safeIndex(new BigNumber(195), new BigNumber((_1392_v29).length))], _1345_v2);
          _nw217[(new BigNumber(8)).toNumber()] = _1345_v2;
          _nw217[(new BigNumber(9)).toNumber()] = _dafny.Seq.UnicodeFromString("untl");
          _nw217[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("vfldre"), _1345_v2);
          _nw217[(new BigNumber(11)).toNumber()] = _1345_v2;
          _nw217[(new BigNumber(12)).toNumber()] = _1345_v2;
          _nw217[(new BigNumber(13)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(349)), ((_1395_v26) => function (_1396_i2) {
            return _1395_v26;
          })(_1389_v26));
          _nw217[(new BigNumber(14)).toNumber()] = _1345_v2;
          _nw217[(new BigNumber(15)).toNumber()] = _1345_v2;
          _nw217[(new BigNumber(16)).toNumber()] = _1345_v2;
          _nw217[(new BigNumber(17)).toNumber()] = _1345_v2;
          _nw217[(new BigNumber(18)).toNumber()] = _1345_v2;
          _nw217[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_1345_v2, (_module.__default.fm34(new BigNumber(500), globalState)).dtor_cf0);
          _nw217[(new BigNumber(20)).toNumber()] = _dafny.Seq.UnicodeFromString("k");
          _nw217[(new BigNumber(21)).toNumber()] = ((p0) ? (_1345_v2) : (_1345_v2));
          _nw217[(new BigNumber(22)).toNumber()] = _dafny.Seq.Concat(_1345_v2, _dafny.Seq.update(_1345_v2, _module.__default.safeIndex((_this).f23, new BigNumber((_1345_v2).length)), _1389_v26));
          _1393_v30 = _nw217;
          let _index246 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_1393_v30).length));
          (_1393_v30)[_index246] = _1345_v2;
          let _index247 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_1393_v30).length));
          (_1393_v30)[_index247] = _1345_v2;
          (globalState).f6 = _dafny.Seq.Concat((_1393_v30)[_module.__default.safeIndex(new BigNumber(464), new BigNumber((_1393_v30).length))], _dafny.Seq.UnicodeFromString("qwammv"));
          let _1397_v31;
          let _nw218 = new _module.C0();
          _nw218.__ctor();
          _1397_v31 = _nw218;
          let _1398_v32;
          _1398_v32 = _dafny.Seq.of(_1397_v31, _1397_v31, _1397_v31);
          let _1399_v33;
          _1399_v33 = _module.D17.create_DC48(_1382_v19);
          let _1400_v34;
          _1400_v34 = _dafny.Set.fromElements(p0);
          _1398_v32 = (((_module.__default.fm35(p0, false, (_1399_v33).dtor_cf79, globalState)).IsDisjointFrom(_1400_v34)) ? (_dafny.Seq.Concat(_1398_v32, _1398_v32)) : (_1398_v32));
          let _index248 = _module.__default.safeIndex(new BigNumber(464), new BigNumber((_1393_v30).length));
          (_1393_v30)[_index248] = _1345_v2;
        } else {
          r0 = _module.__default.fm1((_this).f23, p0, globalState);
          let _1401_v35;
          _1401_v35 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,(_this).f26);
          let _1402_v36;
          let _nw219 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
          _1402_v36 = _nw219;
          let _1403_v37;
          _1403_v37 = _dafny.Seq.of(_1402_v36, _1402_v36, _1402_v36);
          let _1404_v38;
          let _nw220 = new _module.C1();
          _nw220.__ctor(p0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-756)), function (_1405_i3) {
            return (_this).f26;
          }), _module.__default.safeModuloInt((((_1401_v35).contains(new BigNumber((_1403_v37).length))) ? ((_1401_v35).get(new BigNumber((_1403_v37).length))) : ((_this).f26)), (_this).f23));
          _1404_v38 = _nw220;
          r2 = new BigNumber(-370);
          (globalState).f9 = p0;
          let _1406_v39;
          _1406_v39 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _1407_v40;
          _1407_v40 = _dafny.Seq.of(p0, p0, p0);
          let _1408_v41;
          _1408_v41 = _dafny.Seq.of(p0, (_1407_v40)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements((_this).f25)).cardinality()), new BigNumber((_1407_v40).length))], p0, p0);
          _1406_v39 = (_1406_v39).update(p0, (_1407_v40)[_module.__default.safeIndex((_this).f25, new BigNumber((_1407_v40).length))]);
        }
        r2 = _module.__default.safeDivisionInt((_this).f25, (_this).f25);
        let _1409_v42;
        let _init40 = function (_1410_i4) {
          return false;
        };
        let _nw221 = Array((new BigNumber(10)).toNumber());
        for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw221.length); _i0_40++) {
          _nw221[_i0_40] = _init40(new BigNumber(_i0_40));
        }
        _1409_v42 = _nw221;
        let _1411_v43;
        _1411_v43 = _module.D1.create_DC4(_dafny.Map.Empty.slice().updateUnsafe((_this).f25,(_this).f25), p0, p0);
        let _index249 = _module.__default.safeIndex(new BigNumber(604), new BigNumber((_1409_v42).length));
        (_1409_v42)[_index249] = (_1411_v43).dtor_cf8;
        let _index250 = _module.__default.safeIndex(new BigNumber(604), new BigNumber((_1409_v42).length));
        (_1409_v42)[_index250] = p0;
      } else {
        let _1412___mcc_h6 = (_source11).cf73;
        let _1413_cf73 = _1412___mcc_h6;
        (globalState).f14 = ((_dafny.ZERO).minus((_this).f25)).minus((_this).f25);
        let _1414_v44;
        let _nw222 = Array((new BigNumber(2)).toNumber());
        _nw222[(_dafny.ZERO).toNumber()] = (_this).f25;
        _nw222[(_dafny.ONE).toNumber()] = new BigNumber(79);
        _1414_v44 = _nw222;
        let _1415_v45;
        _1415_v45 = _dafny.Seq.of(_1414_v44, _1414_v44, _1414_v44);
        _1415_v45 = _1415_v45;
        let _1416_v46;
        _1416_v46 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,_module.__default.fm0((_this).f23, false, globalState));
        let _1417_v47;
        let _nw223 = new _module.C2();
        _nw223.__ctor(_1416_v46, ((p0) ? (p0) : (p0)), ((true) ? (p0) : (p0)), _dafny.Seq.Concat(_1346_v3, _1346_v3), (_this).f23);
        _1417_v47 = _nw223;
        let _1418_v48;
        let _nw224 = new _module.C0();
        _nw224.__ctor();
        _1418_v48 = _nw224;
      }
      (globalState).f7 = (p0) && (p0);
      (globalState).f9 = !(p0) || (p0);
      let _hi15 = (_this).f23;
      for (let _1419_i5 = (_this).f23; _1419_i5.isLessThan(_hi15); _1419_i5 = _1419_i5.plus(_dafny.ONE)) {
        let _1420_v49;
        _1420_v49 = _dafny.Seq.of(p0);
        let _1421_v50;
        _1421_v50 = _dafny.Seq.of(_1420_v49, _dafny.Seq.update(_dafny.Seq.of(_module.__default.fm1(new BigNumber((_dafny.MultiSet.FromArray(_1346_v3)).cardinality()), p0, globalState)), _module.__default.safeIndex(_module.__default.fm0((_this).f26, p0, globalState), new BigNumber((_dafny.Seq.of(_module.__default.fm1(new BigNumber((_dafny.MultiSet.FromArray(_1346_v3)).cardinality()), p0, globalState))).length)), p0), _dafny.Seq.update(_1420_v49, _module.__default.safeIndex((_this).f25, new BigNumber((_1420_v49).length)), p0));
        let _1422_v51;
        _1422_v51 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0((_this).f23, p0, globalState),(_this).f26);
        let _1423_v52;
        _1423_v52 = _module.D8.create_DC19(_1346_v3);
        let _1424_v53;
        let _nw225 = Array((new BigNumber(27)).toNumber());
        _nw225[(_dafny.ZERO).toNumber()] = _1420_v49;
        _nw225[(_dafny.ONE).toNumber()] = (_1421_v50)[_module.__default.safeIndex((_this).f25, new BigNumber((_1421_v50).length))];
        _nw225[(new BigNumber(2)).toNumber()] = _1420_v49;
        _nw225[(new BigNumber(3)).toNumber()] = _1420_v49;
        _nw225[(new BigNumber(4)).toNumber()] = _1420_v49;
        _nw225[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_1420_v49, _1420_v49);
        _nw225[(new BigNumber(6)).toNumber()] = _1420_v49;
        _nw225[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_1420_v49, _1420_v49);
        _nw225[(new BigNumber(8)).toNumber()] = _1420_v49;
        _nw225[(new BigNumber(9)).toNumber()] = ((p0) ? (_dafny.Seq.of(true)) : (_1420_v49));
        _nw225[(new BigNumber(10)).toNumber()] = _1420_v49;
        _nw225[(new BigNumber(11)).toNumber()] = _module.__default.fm30((((_1422_v51).contains(new BigNumber((_dafny.Seq.update(_1346_v3, _module.__default.safeIndex((_dafny.ZERO).minus((_this).f23), new BigNumber((_1346_v3).length)), new BigNumber(612))).length))) ? ((_1422_v51).get(new BigNumber((_dafny.Seq.update(_1346_v3, _module.__default.safeIndex((_dafny.ZERO).minus((_this).f23), new BigNumber((_1346_v3).length)), new BigNumber(612))).length))) : ((_this).f26)), _1423_v52, _1345_v2, _1419_i5, globalState);
        _nw225[(new BigNumber(12)).toNumber()] = _1420_v49;
        _nw225[(new BigNumber(13)).toNumber()] = _1420_v49;
        _nw225[(new BigNumber(14)).toNumber()] = _module.__default.fm30((_this).f23, _1423_v52, _1345_v2, (_this).f25, globalState);
        _nw225[(new BigNumber(15)).toNumber()] = _1420_v49;
        _nw225[(new BigNumber(16)).toNumber()] = ((true) ? (_1420_v49) : (_1420_v49));
        _nw225[(new BigNumber(17)).toNumber()] = _1420_v49;
        _nw225[(new BigNumber(18)).toNumber()] = _dafny.Seq.update(_1420_v49, _module.__default.safeIndex(new BigNumber(261), new BigNumber((_1420_v49).length)), p0);
        _nw225[(new BigNumber(19)).toNumber()] = _1420_v49;
        _nw225[(new BigNumber(20)).toNumber()] = _dafny.Seq.of(p0, p0);
        _nw225[(new BigNumber(21)).toNumber()] = _1420_v49;
        _nw225[(new BigNumber(22)).toNumber()] = _dafny.Seq.Concat(_1420_v49, _1420_v49);
        _nw225[(new BigNumber(23)).toNumber()] = _1420_v49;
        _nw225[(new BigNumber(24)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm30((_this).f23, _1423_v52, _1345_v2, _1419_i5, globalState), _dafny.Seq.of(p0, true, !(p0)));
        _nw225[(new BigNumber(25)).toNumber()] = _dafny.Seq.update(_1420_v49, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_1345_v2).length)), new BigNumber((_1420_v49).length)), true);
        _nw225[(new BigNumber(26)).toNumber()] = _1420_v49;
        _1424_v53 = _nw225;
        _1424_v53 = _1424_v53;
        let _rhs257 = (_this).f26;
        let _rhs258 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1345_v2, _1345_v2), _1345_v2);
        let _rhs259 = p0;
        let _lhs228 = globalState;
        let _lhs229 = globalState;
        let _lhs230 = globalState;
        _lhs228.f8 = _rhs257;
        _lhs229.f6 = _rhs258;
        _lhs230.f9 = _rhs259;
        let _1425_v54;
        let _nw226 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
        _1425_v54 = _nw226;
        let _1426_v55;
        _1426_v55 = _module.D8.create_DC20(true, _1419_i5);
        let _1427_v56;
        _1427_v56 = _module.D8.create_DC23(_1426_v55);
        let _1428_v57;
        let _nw227 = new _module.C2();
        _nw227.__ctor(_1422_v51, p0, !(p0), _1346_v3, (_this).f25);
        _1428_v57 = _nw227;
        let _1429_v58;
        _1429_v58 = _dafny.Seq.of(_1428_v57);
        let _1430_v59;
        _1430_v59 = _dafny.Seq.of(_1425_v54, _1425_v54);
        let _1431_v60;
        _1431_v60 = _dafny.Set.fromElements((_this).f26, (_this).f26);
        let _1432_v61;
        _1432_v61 = _dafny.Map.Empty.slice().updateUnsafe((_1428_v57).f28,_1428_v57);
        let _1433_v62;
        _1433_v62 = _1428_v57;
        let _1434_v63;
        let _nw228 = Array((new BigNumber(28)).toNumber());
        _nw228[(_dafny.ZERO).toNumber()] = _1428_v57;
        _nw228[(_dafny.ONE).toNumber()] = (_1429_v58)[_module.__default.safeIndex(new BigNumber((_1430_v59).length), new BigNumber((_1429_v58).length))];
        _nw228[(new BigNumber(2)).toNumber()] = _1428_v57;
        _nw228[(new BigNumber(3)).toNumber()] = _1428_v57;
        _nw228[(new BigNumber(4)).toNumber()] = _1428_v57;
        _nw228[(new BigNumber(5)).toNumber()] = _1428_v57;
        _nw228[(new BigNumber(6)).toNumber()] = _1428_v57;
        _nw228[(new BigNumber(7)).toNumber()] = _1428_v57;
        _nw228[(new BigNumber(8)).toNumber()] = _1428_v57;
        _nw228[(new BigNumber(9)).toNumber()] = (_1429_v58)[_module.__default.safeIndex(_1419_i5, new BigNumber((_1429_v58).length))];
        _nw228[(new BigNumber(10)).toNumber()] = _1428_v57;
        _nw228[(new BigNumber(11)).toNumber()] = (_1429_v58)[_module.__default.safeIndex(new BigNumber((_1431_v60).length), new BigNumber((_1429_v58).length))];
        _nw228[(new BigNumber(12)).toNumber()] = (_1429_v58)[_module.__default.safeIndex((_this).f25, new BigNumber((_1429_v58).length))];
        _nw228[(new BigNumber(13)).toNumber()] = (((_1432_v61).contains(!((_1428_v57).f28))) ? ((_1432_v61).get(!((_1428_v57).f28))) : (_1428_v57));
        _nw228[(new BigNumber(14)).toNumber()] = _1428_v57;
        _nw228[(new BigNumber(15)).toNumber()] = (_1433_v62);
        _nw228[(new BigNumber(16)).toNumber()] = _1428_v57;
        _nw228[(new BigNumber(17)).toNumber()] = _1428_v57;
        _nw228[(new BigNumber(18)).toNumber()] = _1428_v57;
        _nw228[(new BigNumber(19)).toNumber()] = _1428_v57;
        _nw228[(new BigNumber(20)).toNumber()] = _1428_v57;
        _nw228[(new BigNumber(21)).toNumber()] = (_1433_v62);
        _nw228[(new BigNumber(22)).toNumber()] = _1428_v57;
        _nw228[(new BigNumber(23)).toNumber()] = _1428_v57;
        _nw228[(new BigNumber(24)).toNumber()] = _1428_v57;
        _nw228[(new BigNumber(25)).toNumber()] = _1428_v57;
        _nw228[(new BigNumber(26)).toNumber()] = _1428_v57;
        _nw228[(new BigNumber(27)).toNumber()] = _1428_v57;
        _1434_v63 = _nw228;
        let _index251 = _module.__default.safeIndex(new BigNumber(638), new BigNumber((_1434_v63).length));
        (_1434_v63)[_index251] = _1428_v57;
        let _index252 = _module.__default.safeIndex(new BigNumber(638), new BigNumber((_1434_v63).length));
        let _rhs260 = ((_this).f25).multipliedBy((_1428_v57).f23);
        let _rhs261 = _1425_v54;
        let _rhs262 = p0;
        let _rhs263 = _1427_v56;
        let _rhs264 = _1428_v57;
        let _lhs231 = globalState;
        let _lhs232 = globalState;
        let _lhs233 = _1434_v63;
        let _lhs234 = _module.__default.safeIndex(new BigNumber(638), new BigNumber((_1434_v63).length));
        _lhs231.f14 = _rhs260;
        _1425_v54 = _rhs261;
        _lhs232.f9 = _rhs262;
        _1427_v56 = _rhs263;
        _lhs233[_lhs234] = _rhs264;
        (globalState).f7 = ((_this).f26).isEqualTo((((_1428_v57).f28) ? (new BigNumber((_1420_v49).length)) : ((_this).f25)));
      }
      let _1435_i6;
      _1435_i6 = _dafny.ZERO;
      L13: {
        while (((_this).f23).isLessThan((_this).f26)) {
          C13: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1435_i6)) {
              break L13;
            }
            _1435_i6 = (_1435_i6).plus(_dafny.ONE);
            let _1436_v64;
            _1436_v64 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f26);
            let _1437_v65;
            _1437_v65 = (_1436_v64).Merge(_1436_v64);
            let _source12 = _1437_v65;
            let _1438___mcc_h7 = _source12;
            let _1439_cf31 = _1438___mcc_h7;
            let _1440_v66;
            _1440_v66 = _dafny.Seq.of(true);
            let _1441_v67;
            let _nw229 = Array((new BigNumber(14)).toNumber()).fill(false);
            _1441_v67 = _nw229;
            let _1442_v68;
            _1442_v68 = _dafny.Map.Empty.slice().updateUnsafe(_1440_v66,_1441_v67);
            let _1443_v69;
            _1443_v69 = _dafny.Seq.of(_1442_v68);
            let _1444_v70;
            _1444_v70 = _dafny.Map.Empty.slice().updateUnsafe((_1443_v69)[_module.__default.safeIndex((_this).f23, new BigNumber((_1443_v69).length))],(_dafny.ZERO).minus((new BigNumber(-861)).minus((_this).f26)));
            _1444_v70 = (_1444_v70).update(_1442_v68, (_this).f23);
            (globalState).f21 = true;
            let _1445_v71;
            let _nw230 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Map.Empty);
            _1445_v71 = _nw230;
            let _1446_v72;
            _1446_v72 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(300),!(p0));
            let _index253 = _module.__default.safeIndex(new BigNumber(978), new BigNumber((_1445_v71).length));
            (_1445_v71)[_index253] = (_1446_v72).update((_this).f23, p0);
            let _index254 = _module.__default.safeIndex(new BigNumber(978), new BigNumber((_1445_v71).length));
            (_1445_v71)[_index254] = _1446_v72;
            let _1447_v73;
            _1447_v73 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(((_this).f25).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(705)), function (_1448_i7) {
              return (_this).f25;
            })).length))),_dafny.Map.Empty.slice().updateUnsafe((_this).f23,!(p0)));
            _1447_v73 = (_1447_v73).update(new BigNumber((_dafny.Seq.of(new BigNumber(-201), new BigNumber(598))).length), _dafny.Map.Empty.slice().updateUnsafe((_this).f25,p0));
            let _1449_v74;
            _1449_v74 = _dafny.Seq.of(_1348_v5, _1348_v5, _1348_v5);
            _1449_v74 = _1449_v74;
            let _1450_v75;
            let _nw231 = new _module.C0();
            _nw231.__ctor();
            _1450_v75 = _nw231;
            let _1451_v76;
            _1451_v76 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,_1450_v75);
            _1451_v76 = (_1451_v76).update(_module.__default.safeModuloInt((_this).f23, (_this).f25), _1450_v75);
            (globalState).f8 = ((_1346_v3)[_module.__default.safeIndex((_this).f26, new BigNumber((_1346_v3).length))]).plus(((p0) ? ((_this).f23) : (new BigNumber(((_1436_v64).update(p0, (_this).f26)).length))));
          }
        }
      }
      r0 = p0;
      r1 = (_this).f25;
      let _1452_v77;
      _1452_v77 = _module.D0.create_DC0(_1345_v2);
      let _1453_v78;
      _1453_v78 = _module.D0.create_DC2(_1452_v77);
      let _pat_let_tv44 = _1346_v3;
      let _pat_let_tv45 = _1346_v3;
      let _pat_let_tv46 = _1346_v3;
      let _pat_let_tv47 = _1346_v3;
      r2 = function (_source13) {
        if (_source13.is_DC1) {
          let _1454___mcc_h8 = (_source13).cf1;
          let _1455___mcc_h9 = (_source13).cf2;
          let _1456___mcc_h10 = (_source13).cf3;
          let _1457_cf3 = _1456___mcc_h10;
          let _1458_cf2 = _1455___mcc_h9;
          let _1459_cf1 = _1454___mcc_h8;
          return ((_this).f26).plus((_this).f23);
        } else if (_source13.is_DC0) {
          let _1460___mcc_h11 = (_source13).cf0;
          let _1461_cf0 = _1460___mcc_h11;
          return _module.__default.safeModuloInt(new BigNumber((function () {
            let _coll52 = new _dafny.Map();
            for (const _compr_52 of (function () {
              let _coll53 = new _dafny.Set();
              for (const _compr_53 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-298)), ((_1462_v3) => function (_1463_i8) {
                return _1462_v3;
              })(_pat_let_tv44))).Elements) {
                let _1464_v80 = _compr_53;
                if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-298)), ((_1465_v3) => function (_1463_i8) {
                  return _1465_v3;
                })(_pat_let_tv45)), _1464_v80)) {
                  _coll53.add(_1464_v80);
                }
              }
              return _coll53;
            }()).Elements) {
              let _1466_v79 = _compr_52;
              if ((function () {
                let _coll54 = new _dafny.Set();
                for (const _compr_54 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-298)), ((_1467_v3) => function (_1463_i8) {
                  return _1467_v3;
                })(_pat_let_tv46))).Elements) {
                  let _1468_v80 = _compr_54;
                  if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-298)), ((_1469_v3) => function (_1463_i8) {
                    return _1469_v3;
                  })(_pat_let_tv47)), _1468_v80)) {
                    _coll54.add(_1468_v80);
                  }
                }
                return _coll54;
              }()).contains(_1466_v79)) {
                _coll52.push([_1466_v79,(_this).f25]);
              }
            }
            return _coll52;
          }()).length), new BigNumber(362));
        } else {
          let _1470___mcc_h12 = (_source13).cf4;
          let _1471_cf4 = _1470___mcc_h12;
          return (_this).f26;
        }
      }(_1453_v78);
      return [r0, r1, r2];
    }
    get f25() {
      let _this = this;
      return _this._f25;
    };
    get f26() {
      let _this = this;
      return _this._f26;
    };
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this._f23 = _dafny.ZERO;
      this._f24 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f23() {
      let _this = this;
      return _this._f23;
    };
    __ctor(f24, f23) {
      let _this = this;
      (_this)._f24 = f24;
      (_this)._f23 = f23;
      return;
    }
    m1(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Set.Empty;
      let _1472_v0;
      _1472_v0 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of((_this).f24, (_this).f23)).length),p0);
      let _hi16 = new BigNumber((_1472_v0).length);
      for (let _1473_i0 = (_this).f23; _1473_i0.isLessThan(_hi16); _1473_i0 = _1473_i0.plus(_dafny.ONE)) {
        r0 = (_this).f24;
        let _1474_v1;
        _1474_v1 = _dafny.Map.Empty.slice().updateUnsafe(p1,!(p1));
        let _1475_v2;
        _1475_v2 = _dafny.Set.fromElements(((p1) ? (_1474_v1) : (_dafny.Map.Empty.slice().updateUnsafe(!(p1),p1))), _1474_v1, _1474_v1, _1474_v1);
        let _1476_v3;
        _1476_v3 = new _dafny.CodePoint('j'.codePointAt(0));
        let _1477_v4;
        _1477_v4 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('n'.codePointAt(0)), _1476_v3);
        _1475_v2 = _dafny.Set.fromElements((_1474_v1).update(_module.__default.fm1(new BigNumber((_1477_v4).cardinality()), p1, globalState), p1), (_module.__default.fm7((_this).f24, p1, p1, globalState)).Merge(_1474_v1), _1474_v1, (_dafny.Map.Empty.slice().updateUnsafe(p1,p1)).Merge(_1474_v1));
        let _index255 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((p0).length));
        (p0)[_index255] = p1;
        let _1478_v5;
        _1478_v5 = _dafny.MultiSet.fromElements(p1, p1);
        let _1479_v6;
        _1479_v6 = _dafny.Seq.of(p1);
        let _index256 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((p0).length));
        (p0)[_index256] = (_1478_v5).IsDisjointFrom(_module.__default.fm4(_module.D0.create_DC0(_dafny.Seq.Create(_module.__default.abs(new BigNumber(591)), ((_1480_v3) => function (_1481_i1) {
  return _1480_v3;
})(_1476_v3))), _1473_i0, _1479_v6, _1476_v3, globalState));
        let _1482_v7;
        _1482_v7 = _dafny.Seq.UnicodeFromString("cduaj");
        let _1483_v8;
        _1483_v8 = _dafny.Seq.of(_1482_v7);
        (globalState).f14 = (_dafny.ZERO).minus(new BigNumber((_1483_v8).length));
      }
      let _1484_v9;
      _1484_v9 = _dafny.Seq.of((_this).f23);
      (globalState).f14 = (_1484_v9)[_module.__default.safeIndex(new BigNumber(439), new BigNumber((_1484_v9).length))];
      let _index257 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((p0).length));
      (p0)[_index257] = !(_module.__default.fm0((_this).f23, p1, globalState)).isEqualTo((_this).f24);
      let _1485_v10;
      _1485_v10 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,(_this).f24);
      let _1486_v11;
      _1486_v11 = _dafny.Seq.UnicodeFromString("l");
      let _1487_v12;
      _1487_v12 = _dafny.Set.fromElements((_this).f24, (((_1485_v10).contains((_this).f24)) ? ((_1485_v10).get((_this).f24)) : (new BigNumber((_1486_v11).length))));
      let _1488_v13;
      _1488_v13 = _dafny.MultiSet.fromElements(_module.__default.fm1(new BigNumber((_1487_v12).length), false, globalState), p1);
      let _1489_v14;
      _1489_v14 = _dafny.Set.fromElements(p1, p1);
      let _1490_v15;
      let _nw232 = new _module.C3();
      _nw232.__ctor(_dafny.Seq.of((_1488_v13).update(true, _module.__default.abs(new BigNumber((_1489_v14).length))), _dafny.MultiSet.fromElements(p1, p1), _1488_v13, _1488_v13, _dafny.MultiSet.fromElements(p1)), (_this).f23, p1, _1484_v9);
      _1490_v15 = _nw232;
      let _1491_v16;
      _1491_v16 = _dafny.Seq.of(_1490_v15);
      let _index258 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((p0).length));
      (p0)[_index258] = !_dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.update(_1491_v16, _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p1,p1)).length), new BigNumber((_1491_v16).length)), _1490_v15), _1491_v16), _1490_v15);
      let _1492_v17;
      _1492_v17 = _dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber(904));
      let _1493_v18;
      _1493_v18 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((p0)[_module.__default.safeIndex(new BigNumber(487), new BigNumber((p0).length))],(_1490_v15).f23), _1492_v17);
      r0 = ((_dafny.Seq.contains(_1493_v18, _1492_v17)) ? (((_this).f23).multipliedBy((_this).f24)) : (new BigNumber(-954)));
      let _hi17 = _module.__default.safeDivisionInt((_this).f24, (_this).f23);
      for (let _1494_i2 = (_1490_v15).f23; _1494_i2.isLessThan(_hi17); _1494_i2 = _1494_i2.plus(_dafny.ONE)) {
        let _1495_v19;
        _1495_v19 = _dafny.Map.Empty.slice().updateUnsafe(false,!(!(p1)));
        _1495_v19 = (_1495_v19).update(p1, _module.__default.fm1((_1490_v15).f23, p1, globalState));
        let _1496_v20;
        _1496_v20 = _dafny.Seq.of(_module.__default.fm1((_this).f23, (p0)[_module.__default.safeIndex(new BigNumber(487), new BigNumber((p0).length))], globalState), p1, (_1488_v13).IsDisjointFrom(_1488_v13), p1, ((p1) ? (!(p1)) : (true)));
        (globalState).f7 = (_1496_v20)[_module.__default.safeIndex(new BigNumber((_module.__default.fm2(globalState)).length), new BigNumber((_1496_v20).length))];
        let _1497_v21;
        _1497_v21 = new _dafny.CodePoint('n'.codePointAt(0));
        let _1498_v22;
        _1498_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1497_v21,_dafny.Seq.Concat(_1484_v9, _dafny.Seq.of((_1490_v15).f23)));
        _1498_v22 = (_1498_v22).update((_1486_v11)[_module.__default.safeIndex((_1490_v15).f23, new BigNumber((_1486_v11).length))], _1484_v9);
        _1496_v20 = _dafny.Seq.Concat(_1496_v20, _1496_v20);
      }
      let _1499_v23;
      let _nw233 = Array((new BigNumber(15)).toNumber()).fill(false);
      _1499_v23 = _nw233;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1499_v23).length))) {
        let _1500_i3 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1500_i3)) && ((_1500_i3).isLessThan(new BigNumber((_1499_v23).length))))) {
          (_1499_v23)[(_1500_i3)] = _dafny.Seq.IsPrefixOf(_1486_v11, _1486_v11);
        }
      }
      r0 = (_this).f24;
      let _1501_v24;
      _1501_v24 = _dafny.MultiSet.fromElements((_this).f23);
      let _1502_v25;
      _1502_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1501_v24,new BigNumber((_1488_v13).cardinality()));
      let _1503_v26;
      _1503_v26 = _module.D18.create_DC50(_1502_v25);
      r1 = _dafny.Set.fromElements((_1490_v15).f23, new BigNumber((((_1503_v26).dtor_cf84).Merge(_1502_v25)).length), (_this).f23);
      return [r0, r1];
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Seq.of();
      let _1504_v0;
      _1504_v0 = true;
      let _1505_v1;
      _1505_v1 = _dafny.Map.Empty.slice().updateUnsafe(_1504_v0,(_this).f23);
      let _1506_v2;
      _1506_v2 = new _dafny.CodePoint('u'.codePointAt(0));
      let _1507_v3;
      _1507_v3 = _dafny.Set.fromElements(_1506_v2);
      let _1508_v4;
      _1508_v4 = _dafny.Map.Empty.slice().updateUnsafe(_1507_v3,p0);
      let _1509_v5;
      _1509_v5 = _dafny.Set.fromElements(new BigNumber(((((_1508_v4).contains(_1507_v3)) ? ((_1508_v4).get(_1507_v3)) : (_dafny.Seq.update(_dafny.Seq.update(p0, _module.__default.safeIndex((_this).f23, new BigNumber((p0).length)), _1506_v2), _module.__default.safeIndex((_this).f23, new BigNumber((_dafny.Seq.update(p0, _module.__default.safeIndex((_this).f23, new BigNumber((p0).length)), _1506_v2)).length)), _1506_v2)))).length), (_this).f23);
      _1505_v1 = (_1505_v1).update((_1509_v5).IsProperSubsetOf(_dafny.Set.fromElements((_this).f23)), (_this).f24);
      _1506_v2 = _1506_v2;
      if (_1504_v0) {
        let _1510_v6;
        _1510_v6 = _module.D3.create_DC10(!(_1504_v0));
        let _source14 = _1510_v6;
        if (_source14.is_DC10) {
          let _1511___mcc_h0 = (_source14).cf21;
          let _1512_cf21 = _1511___mcc_h0;
          let _1513_v7;
          let _nw234 = Array((new BigNumber(22)).toNumber()).fill([]);
          _1513_v7 = _nw234;
          let _1514_v8;
          let _init41 = ((_1515_cf21) => function (_1516_i0) {
            return _1515_cf21;
          })(_1512_cf21);
          let _nw235 = Array((new BigNumber(13)).toNumber());
          for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw235.length); _i0_41++) {
            _nw235[_i0_41] = _init41(new BigNumber(_i0_41));
          }
          _1514_v8 = _nw235;
          let _index259 = _module.__default.safeIndex(new BigNumber(540), new BigNumber((_1513_v7).length));
          (_1513_v7)[_index259] = _1514_v8;
          let _1517_v9;
          _1517_v9 = _dafny.Seq.of((_this).f24);
          let _1518_v10;
          let _nw236 = new _module.C3();
          _nw236.__ctor(_module.__default.fm29(new BigNumber((_1517_v9).length), _dafny.Set.fromElements((_this).f23, (_this).f23), globalState), (_this).f24, false, _dafny.Seq.Concat(_1517_v9, _1517_v9));
          _1518_v10 = _nw236;
          let _index260 = _module.__default.safeIndex(new BigNumber(540), new BigNumber((_1513_v7).length));
          let _rhs265 = (_dafny.ZERO).minus((_this).f23);
          let _rhs266 = _1506_v2;
          let _rhs267 = _1514_v8;
          let _rhs268 = _1518_v10;
          let _rhs269 = _1514_v8;
          let _lhs235 = globalState;
          let _lhs236 = _1513_v7;
          let _lhs237 = _module.__default.safeIndex(new BigNumber(540), new BigNumber((_1513_v7).length));
          _lhs235.f14 = _rhs265;
          _1506_v2 = _rhs266;
          _lhs236[_lhs237] = _rhs267;
          _1518_v10 = _rhs268;
          _1514_v8 = _rhs269;
          (globalState).f6 = _module.__default.fm2(globalState);
          (globalState).f21 = _1504_v0;
          let _index261 = _module.__default.safeIndex(new BigNumber(682), new BigNumber((_1514_v8).length));
          (_1514_v8)[_index261] = _1512_cf21;
          let _index262 = _module.__default.safeIndex(new BigNumber(682), new BigNumber((_1514_v8).length));
          (_1514_v8)[_index262] = _1512_cf21;
        } else if (_source14.is_DC11) {
          let _1519___mcc_h1 = (_source14).cf22;
          let _1520_cf22 = _1519___mcc_h1;
          let _1521_v11;
          _1521_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1509_v5,_dafny.Seq.IsPrefixOf(p0, _dafny.Seq.update(p0, _module.__default.safeIndex((_this).f23, new BigNumber((p0).length)), _1506_v2)));
          _1521_v11 = (_1521_v11).update(_1509_v5, _1504_v0);
          let _1522_v12;
          _1522_v12 = _dafny.MultiSet.fromElements(_1504_v0);
          let _1523_v13;
          let _nw237 = Array((new BigNumber(8)).toNumber());
          _nw237[(_dafny.ZERO).toNumber()] = (_this).f24;
          _nw237[(_dafny.ONE).toNumber()] = _1520_cf22;
          _nw237[(new BigNumber(2)).toNumber()] = _1520_cf22;
          _nw237[(new BigNumber(3)).toNumber()] = new BigNumber((_1505_v1).length);
          _nw237[(new BigNumber(4)).toNumber()] = (_this).f24;
          _nw237[(new BigNumber(5)).toNumber()] = (_this).f24;
          _nw237[(new BigNumber(6)).toNumber()] = new BigNumber((_1522_v12).cardinality());
          _nw237[(new BigNumber(7)).toNumber()] = (_this).f24;
          _1523_v13 = _nw237;
          let _1524_v14;
          _1524_v14 = _dafny.Map.Empty.slice().updateUnsafe(_1523_v13,_1504_v0);
          let _1525_v15;
          _1525_v15 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1504_v0,(((_1524_v14).contains(_1523_v13)) ? ((_1524_v14).get(_1523_v13)) : (_1504_v0)))).length), new BigNumber((p0).length), (_this).f23, _1520_cf22, _1520_cf22);
          let _1526_v16;
          let _nw238 = Array((new BigNumber(11)).toNumber()).fill(false);
          _1526_v16 = _nw238;
          let _1527_v17;
          _1527_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1525_v15,_1526_v16);
          let _1528_v18;
          _1528_v18 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f24),_1527_v17);
          _1528_v18 = (_1528_v18).update((_this).f24, _1527_v17);
          let _1529_v19;
          _1529_v19 = _module.D2.create_DC7(_1507_v3);
          let _1530_v20;
          _1530_v20 = _dafny.Map.Empty.slice().updateUnsafe(_1529_v19,_1504_v0);
          let _1531_v21;
          _1531_v21 = _dafny.Map.Empty.slice().updateUnsafe((_1530_v20).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1529_v19,_1504_v0)),(_this).f24);
          _1531_v21 = (_1531_v21).update(_1530_v20, (_this).f24);
          let _1532_v22;
          let _nw239 = new _module.C5();
          _nw239.__ctor(new BigNumber((p0).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("r")).length)), (_this).f24);
          _1532_v22 = _nw239;
          _1532_v22 = _1532_v22;
        } else if (_source14.is_DC12) {
          let _1533___mcc_h2 = (_source14).cf23;
          let _1534_cf23 = _1533___mcc_h2;
          let _1535_v23;
          _1535_v23 = _dafny.Map.Empty.slice().updateUnsafe((_this).f23,(_this).f24);
          let _1536_v24;
          let _nw240 = Array((new BigNumber(12)).toNumber());
          _nw240[(_dafny.ZERO).toNumber()] = (_this).f23;
          _nw240[(_dafny.ONE).toNumber()] = _module.__default.fm0((_this).f24, _1504_v0, globalState);
          _nw240[(new BigNumber(2)).toNumber()] = (_this).f24;
          _nw240[(new BigNumber(3)).toNumber()] = ((false) ? ((_this).f24) : ((_this).f24));
          _nw240[(new BigNumber(4)).toNumber()] = (((_1535_v23).contains((_this).f24)) ? ((_1535_v23).get((_this).f24)) : ((_this).f23));
          _nw240[(new BigNumber(5)).toNumber()] = (_this).f24;
          _nw240[(new BigNumber(6)).toNumber()] = ((_this).f23).minus(new BigNumber((p0).length));
          _nw240[(new BigNumber(7)).toNumber()] = (new BigNumber((_1534_cf23).length)).plus((_dafny.ZERO).minus(_module.__default.fm0((_this).f23, _1504_v0, globalState)));
          _nw240[(new BigNumber(8)).toNumber()] = _module.__default.fm0((_dafny.ZERO).minus(new BigNumber((_1535_v23).length)), _1504_v0, globalState);
          _nw240[(new BigNumber(9)).toNumber()] = (_this).f24;
          _nw240[(new BigNumber(10)).toNumber()] = (_this).f24;
          _nw240[(new BigNumber(11)).toNumber()] = (_this).f24;
          _1536_v24 = _nw240;
          _1536_v24 = _1536_v24;
          let _1537_v25;
          _1537_v25 = _dafny.Seq.of(p0, p0);
          let _1538_v26;
          _1538_v26 = _dafny.Seq.of(_dafny.Seq.update(p0, _module.__default.safeIndex(new BigNumber((_1537_v25).length), new BigNumber((p0).length)), _1506_v2));
          let _1539_v27;
          _1539_v27 = _module.D11.create_DC32((_1538_v26)[_module.__default.safeIndex((_this).f24, new BigNumber((_1538_v26).length))]);
          let _1540_v28;
          _1540_v28 = _dafny.MultiSet.fromElements(_1539_v27, _1539_v27);
          let _1541_v29;
          _1541_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1540_v28,_dafny.Seq.update(_dafny.Seq.of((_this).f23, (_this).f24, new BigNumber((_1534_cf23).length), (_this).f24, (_this).f24), _module.__default.safeIndex((_this).f24, new BigNumber((_dafny.Seq.of((_this).f23, (_this).f24, new BigNumber((_1534_cf23).length), (_this).f24, (_this).f24)).length)), _module.__default.fm0((_this).f24, _1504_v0, globalState)));
          let _1542_v30;
          _1542_v30 = _dafny.Seq.of((_this).f23, (_this).f24, (_this).f23, (_this).f23, (_this).f24);
          _1541_v29 = (_1541_v29).update(_1540_v28, _dafny.Seq.Concat(_1542_v30, _1542_v30));
          let _1543_v31;
          _1543_v31 = _dafny.Map.Empty.slice().updateUnsafe(_1504_v0,_dafny.Seq.UnicodeFromString("bxnn"));
          _1543_v31 = ((!(_1504_v0)) ? ((_1543_v31).Merge(_1543_v31)) : (_1543_v31));
          (globalState).f8 = (_dafny.ZERO).minus((_this).f23);
        } else if (_source14.is_DC9) {
          let _1544___mcc_h3 = (_source14).cf20;
          let _1545_cf20 = _1544___mcc_h3;
          (globalState).f9 = _1504_v0;
          (globalState).f8 = (_this).f23;
          (globalState).f6 = p0;
          let _1546_v32;
          _1546_v32 = _dafny.MultiSet.fromElements(_1504_v0, true, true);
          let _1547_v33;
          _1547_v33 = _dafny.Seq.of(_1504_v0, _1504_v0, _1504_v0);
          let _1548_v34;
          let _nw241 = Array((new BigNumber(21)).toNumber());
          _nw241[(_dafny.ZERO).toNumber()] = _1504_v0;
          _nw241[(_dafny.ONE).toNumber()] = ((_this).f23).isLessThanOrEqualTo((_dafny.ZERO).minus(new BigNumber((p0).length)));
          _nw241[(new BigNumber(2)).toNumber()] = !(_1504_v0);
          _nw241[(new BigNumber(3)).toNumber()] = !(_1504_v0);
          _nw241[(new BigNumber(4)).toNumber()] = _1504_v0;
          _nw241[(new BigNumber(5)).toNumber()] = !(_1504_v0);
          _nw241[(new BigNumber(6)).toNumber()] = _module.__default.fm1((_this).f23, _1504_v0, globalState);
          _nw241[(new BigNumber(7)).toNumber()] = _dafny.Seq.contains(p0, _1506_v2);
          _nw241[(new BigNumber(8)).toNumber()] = _1504_v0;
          _nw241[(new BigNumber(9)).toNumber()] = (_1504_v0) === (!(_1504_v0));
          _nw241[(new BigNumber(10)).toNumber()] = true;
          _nw241[(new BigNumber(11)).toNumber()] = _module.__default.fm1((_this).f23, _1504_v0, globalState);
          _nw241[(new BigNumber(12)).toNumber()] = _1504_v0;
          _nw241[(new BigNumber(13)).toNumber()] = (_1546_v32).IsProperSubsetOf(_1546_v32);
          _nw241[(new BigNumber(14)).toNumber()] = _1504_v0;
          _nw241[(new BigNumber(15)).toNumber()] = _1504_v0;
          _nw241[(new BigNumber(16)).toNumber()] = _1504_v0;
          _nw241[(new BigNumber(17)).toNumber()] = _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("thw"), p0);
          _nw241[(new BigNumber(18)).toNumber()] = _1504_v0;
          _nw241[(new BigNumber(19)).toNumber()] = (_1547_v33)[_module.__default.safeIndex((_this).f23, new BigNumber((_1547_v33).length))];
          _nw241[(new BigNumber(20)).toNumber()] = _1504_v0;
          _1548_v34 = _nw241;
          _1548_v34 = _1548_v34;
        } else {
          let _1549___mcc_h4 = (_source14).cf24;
          let _1550_cf24 = _1549___mcc_h4;
          let _1551_v35;
          _1551_v35 = _dafny.MultiSet.fromElements(_1504_v0);
          let _1552_v36;
          _1552_v36 = _dafny.Set.fromElements(_module.__default.fm35(_1504_v0, !(_1504_v0), _1551_v35, globalState));
          let _1553_v37;
          _1553_v37 = _dafny.Seq.of((_this).f24);
          let _1554_v38;
          _1554_v38 = _dafny.Seq.of(_1553_v37);
          let _1555_v39;
          _1555_v39 = _dafny.Map.Empty.slice().updateUnsafe(_1504_v0,_dafny.Seq.of((_this).f23, (_this).f24));
          let _1556_v40;
          _1556_v40 = _module.D12.create_DC36(true, _1504_v0, (_this).f23, (_this).f24, _1555_v39);
          let _1557_v41;
          _1557_v41 = _dafny.MultiSet.fromElements((_this).f24, (_this).f24, (_dafny.ZERO).minus((_this).f24), (_this).f24);
          let _1558_v42;
          let _nw242 = Array((new BigNumber(17)).toNumber());
          _nw242[(_dafny.ZERO).toNumber()] = (_this).f24;
          _nw242[(_dafny.ONE).toNumber()] = (_this).f23;
          _nw242[(new BigNumber(2)).toNumber()] = ((_this).f23).plus((_this).f23);
          _nw242[(new BigNumber(3)).toNumber()] = (_this).f24;
          _nw242[(new BigNumber(4)).toNumber()] = (_this).f23;
          _nw242[(new BigNumber(5)).toNumber()] = (_this).f23;
          _nw242[(new BigNumber(6)).toNumber()] = ((_this).f23).multipliedBy((_this).f24);
          _nw242[(new BigNumber(7)).toNumber()] = (_this).f24;
          _nw242[(new BigNumber(8)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(196), (_this).f24);
          _nw242[(new BigNumber(9)).toNumber()] = (_this).f24;
          _nw242[(new BigNumber(10)).toNumber()] = (_this).f23;
          _nw242[(new BigNumber(11)).toNumber()] = new BigNumber((_1554_v38).length);
          _nw242[(new BigNumber(12)).toNumber()] = (_1556_v40).dtor_cf59;
          _nw242[(new BigNumber(13)).toNumber()] = _module.__default.fm0((((_1557_v41).contains((_this).f24)) ? ((_1557_v41).get((_this).f24)) : (new BigNumber(-104))), _1504_v0, globalState);
          _nw242[(new BigNumber(14)).toNumber()] = (_this).f23;
          _nw242[(new BigNumber(15)).toNumber()] = ((_this).f23).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements((_this).f23, (_this).f24, (_this).f23, new BigNumber((p0).length))).cardinality()));
          _nw242[(new BigNumber(16)).toNumber()] = (_this).f23;
          _1558_v42 = _nw242;
          let _index263 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_1558_v42).length));
          (_1558_v42)[_index263] = (_this).f24;
          let _1559_v43;
          _1559_v43 = _dafny.Set.fromElements(_1504_v0);
          let _index264 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_1558_v42).length));
          let _rhs270 = _dafny.Set.fromElements(_1559_v43);
          let _rhs271 = new BigNumber((p0).length);
          let _rhs272 = new BigNumber(((_dafny.Set.fromElements(_1504_v0, _1504_v0)).Intersect(_1559_v43)).length);
          let _lhs238 = _1558_v42;
          let _lhs239 = _module.__default.safeIndex(new BigNumber(645), new BigNumber((_1558_v42).length));
          _1552_v36 = _rhs270;
          _lhs238[_lhs239] = _rhs271;
          r0 = _rhs272;
          let _1560_v44;
          let _init42 = function (_1561_i1) {
            return true;
          };
          let _nw243 = Array((new BigNumber(15)).toNumber());
          for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw243.length); _i0_42++) {
            _nw243[_i0_42] = _init42(new BigNumber(_i0_42));
          }
          _1560_v44 = _nw243;
          let _1562_v45;
          let _1563_v46;
          let _1564_v47;
          let _out47;
          let _out48;
          let _out49;
          let _outcollector16 = _module.__default.m0(_1560_v44, _dafny.Seq.UnicodeFromString("hqminyem"), (_this).f23, globalState);
          _out47 = _outcollector16[0];
          _out48 = _outcollector16[1];
          _out49 = _outcollector16[2];
          _1562_v45 = _out47;
          _1563_v46 = _out48;
          _1564_v47 = _out49;
          _1504_v0 = !((_this).f23).isEqualTo(((_1558_v42)[_module.__default.safeIndex(new BigNumber(645), new BigNumber((_1558_v42).length))]).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-520)), ((_1565_v2) => function (_1566_i2) {
            return _1565_v2;
          })(_1506_v2))).length)));
          _1505_v1 = (((_1557_v41).IsProperSubsetOf(_1557_v41)) ? (((_1504_v0) ? (_1505_v1) : (_1505_v1))) : (_1505_v1));
        }
        let _1567_v48;
        let _nw244 = Array((new BigNumber(2)).toNumber());
        _nw244[(_dafny.ZERO).toNumber()] = _1504_v0;
        _nw244[(_dafny.ONE).toNumber()] = false;
        _1567_v48 = _nw244;
        _1567_v48 = _1567_v48;
        let _1568_v49;
        let _nw245 = Array((new BigNumber(25)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1568_v49 = _nw245;
        let _1569_v50;
        _1569_v50 = _dafny.Map.Empty.slice().updateUnsafe(_1504_v0,_1506_v2);
        let _nw246 = Array((new BigNumber(6)).toNumber());
        _nw246[(_dafny.ZERO).toNumber()] = _1506_v2;
        _nw246[(_dafny.ONE).toNumber()] = _1506_v2;
        _nw246[(new BigNumber(2)).toNumber()] = _1506_v2;
        _nw246[(new BigNumber(3)).toNumber()] = (((_1569_v50).contains(false)) ? ((_1569_v50).get(false)) : (_module.__default.fm6((_this).f23, globalState)));
        _nw246[(new BigNumber(4)).toNumber()] = (((_1569_v50).contains(_1504_v0)) ? ((_1569_v50).get(_1504_v0)) : (_1506_v2));
        _nw246[(new BigNumber(5)).toNumber()] = _module.__default.fm6((_this).f23, globalState);
        _1568_v49 = _nw246;
        (globalState).f14 = new BigNumber(558);
        let _rhs273 = true;
        let _rhs274 = false;
        let _lhs240 = globalState;
        let _lhs241 = globalState;
        _lhs240.f9 = _rhs273;
        _lhs241.f7 = _rhs274;
      } else {
        (globalState).f8 = (_this).f23;
        let _1570_v51;
        _1570_v51 = _module.D13.create_DC39((_dafny.ZERO).minus(_module.__default.fm0((_this).f24, _1504_v0, globalState)));
        _1570_v51 = ((_1504_v0) ? (_1570_v51) : (_1570_v51));
        let _1571_v52;
        _1571_v52 = _dafny.MultiSet.fromElements(true, _1504_v0);
        let _1572_v53;
        let _nw247 = new _module.C3();
        _nw247.__ctor(_dafny.Seq.of(_1571_v52, _1571_v52, _1571_v52, _1571_v52), (_this).f23, (_1504_v0) === (_1504_v0), _dafny.Seq.Create(_module.__default.abs(new BigNumber(205)), ((_1573_p0) => function (_1574_i3) {
          return new BigNumber((_1573_p0).length);
        })(p0)));
        _1572_v53 = _nw247;
        let _rhs275 = _1506_v2;
        let _rhs276 = (_1572_v53).f23;
        let _rhs277 = _module.__default.fm1((_this).f23, (_1572_v53).f28, globalState);
        let _rhs278 = (((_1572_v53).f28) ? (_1572_v53) : (_1572_v53));
        let _lhs242 = globalState;
        let _lhs243 = globalState;
        let _lhs244 = globalState;
        _lhs242.f11 = _rhs275;
        _lhs243.f14 = _rhs276;
        _lhs244.f21 = _rhs277;
        _1572_v53 = _rhs278;
        (globalState).f7 = (_1572_v53).f28;
        let _1575_v54;
        _1575_v54 = _dafny.Map.Empty.slice().updateUnsafe((_1572_v53).f28,p0);
        (globalState).f6 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(350)), ((_1576_v2) => function (_1577_i4) {
          return _1576_v2;
        })(_1506_v2)), p0), _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(635)), function (_1578_i5) {
          return new _dafny.CodePoint('k'.codePointAt(0));
        }), (((_1575_v54).contains(_1504_v0)) ? ((_1575_v54).get(_1504_v0)) : (p0))), _module.__default.safeIndex((_1572_v53).f23, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(635)), function (_1579_i5) {
          return new _dafny.CodePoint('k'.codePointAt(0));
        }), (((_1575_v54).contains(_1504_v0)) ? ((_1575_v54).get(_1504_v0)) : (p0)))).length)), _module.__default.fm6(new BigNumber(-66), globalState)));
      }
      r0 = ((_this).f23).multipliedBy((_this).f23);
      let _1580_v55;
      let _nw248 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
      _1580_v55 = _nw248;
      for (const _guard_loop_5 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1580_v55).length))) {
        let _1581_i6 = _guard_loop_5;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1581_i6)) && ((_1581_i6).isLessThan(new BigNumber((_1580_v55).length))))) {
          (_1580_v55)[(_1581_i6)] = _module.__default.safeModuloInt(_1581_i6, ((_this).f24).multipliedBy((_this).f24));
        }
      }
      let _1582_v56;
      _1582_v56 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_1504_v0);
      let _1583_v57;
      _1583_v57 = _dafny.Map.Empty.slice().updateUnsafe(_1582_v56,(_this).f24);
      (globalState).f9 = (_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("r"), _1506_v2)) || (((_this).f23).isLessThan(new BigNumber((_1583_v57).length)));
      r0 = (new BigNumber(166)).minus((_this).f23);
      let _1584_v58;
      _1584_v58 = _dafny.Seq.of((_this).f24, (_this).f23);
      r1 = _1584_v58;
      return [r0, r1];
    }
    get f24() {
      let _this = this;
      return _this._f24;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
