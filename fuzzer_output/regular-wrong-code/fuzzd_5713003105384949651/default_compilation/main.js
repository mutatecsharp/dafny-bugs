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
      return false;
    };
    static fm2(p0, globalState) {
      return (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(-490), _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(_dafny.Seq.UnicodeFromString("guwwkcoj"))).length), new BigNumber(-612))));
    };
    static fm3(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(309)), function (_0_i0) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("dwtkddnh")), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("kq"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(218)), function (_1_i1) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      })));
    };
    static fm4(globalState) {
      return (_dafny.ZERO).minus(((!((true) && (false))) ? ((new BigNumber(512)).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('s'.codePointAt(0)))).length))) : (_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(268),true)).length), new BigNumber(589), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length), new BigNumber((_dafny.Seq.UnicodeFromString("ddwiwxcq")).length))).length), new BigNumber(-555)))));
    };
    static fm5(p0, p1, p2, p3, globalState) {
      return _module.D1.create_DC3(new _dafny.CodePoint('d'.codePointAt(0)));
    };
    static fm6(p0, p1, p2, p3, globalState) {
      return new _dafny.CodePoint('s'.codePointAt(0));
    };
    static fm9(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(795)), function (_2_i0) {
        return new _dafny.CodePoint('p'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("l"));
    };
    static fm10(globalState) {
      return (function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(144), new BigNumber((function () {
          let _coll1 = new _dafny.Set();
          for (const _compr_1 of (function () {
            let _coll2 = new _dafny.Set();
            for (const _compr_2 of _dafny.IntegerRange(new BigNumber(66), new BigNumber(817))) {
              let _3_v0 = _compr_2;
              if (((new BigNumber(66)).isLessThanOrEqualTo(_3_v0)) && ((_3_v0).isLessThan(new BigNumber(817)))) {
                _coll2.add(_module.__default.safeModuloInt(_3_v0, new BigNumber(401)));
              }
            }
            return _coll2;
          }()).Elements) {
            let _4_v1 = _compr_1;
            if ((function () {
              let _coll3 = new _dafny.Set();
              for (const _compr_3 of _dafny.IntegerRange(new BigNumber(66), new BigNumber(817))) {
                let _5_v0 = _compr_3;
                if (((new BigNumber(66)).isLessThanOrEqualTo(_5_v0)) && ((_5_v0).isLessThan(new BigNumber(817)))) {
                  _coll3.add(_module.__default.safeModuloInt(_5_v0, new BigNumber(401)));
                }
              }
              return _coll3;
            }()).contains(_4_v1)) {
              _coll1.add(_module.__default.safeModuloInt(_4_v1, new BigNumber((_dafny.Set.fromElements(true, true)).length)));
            }
          }
          return _coll1;
        }()).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true)).length)))).cardinality()),true)).Keys.Elements) {
          let _6_v2 = _compr_0;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(144), new BigNumber((function () {
            let _coll4 = new _dafny.Set();
            for (const _compr_4 of (function () {
              let _coll5 = new _dafny.Set();
              for (const _compr_5 of _dafny.IntegerRange(new BigNumber(66), new BigNumber(817))) {
                let _7_v0 = _compr_5;
                if (((new BigNumber(66)).isLessThanOrEqualTo(_7_v0)) && ((_7_v0).isLessThan(new BigNumber(817)))) {
                  _coll5.add(_module.__default.safeModuloInt(_7_v0, new BigNumber(401)));
                }
              }
              return _coll5;
            }()).Elements) {
              let _8_v1 = _compr_4;
              if ((function () {
                let _coll6 = new _dafny.Set();
                for (const _compr_6 of _dafny.IntegerRange(new BigNumber(66), new BigNumber(817))) {
                  let _9_v0 = _compr_6;
                  if (((new BigNumber(66)).isLessThanOrEqualTo(_9_v0)) && ((_9_v0).isLessThan(new BigNumber(817)))) {
                    _coll6.add(_module.__default.safeModuloInt(_9_v0, new BigNumber(401)));
                  }
                }
                return _coll6;
              }()).contains(_8_v1)) {
                _coll4.add(_module.__default.safeModuloInt(_8_v1, new BigNumber((_dafny.Set.fromElements(true, true)).length)));
              }
            }
            return _coll4;
          }()).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true)).length)))).cardinality()),true)).contains(_6_v2)) {
            _coll0.add(_module.__default.safeDivisionInt(_6_v2, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("vup")).length))));
          }
        }
        return _coll0;
      }()).Intersect(((!(true)) ? (function () {
        let _coll7 = new _dafny.Set();
        for (const _compr_7 of _dafny.IntegerRange(new BigNumber(704), new BigNumber(-431))) {
          let _10_v3 = _compr_7;
          if (((new BigNumber(704)).isLessThanOrEqualTo(_10_v3)) && ((_10_v3).isLessThan(new BigNumber(-431)))) {
            _coll7.add((_10_v3).minus(new BigNumber(-676)));
          }
        }
        return _coll7;
      }()) : (_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("qdem")).length), new BigNumber((_dafny.Seq.of(true, true, true, !((_module.D0.create_DC0(true)).dtor_cf0), false)).length)))));
    };
    static fm11(p0, p1, globalState) {
      return ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(true, false, false, !(!(!(false))), true)).length))).Union(_dafny.MultiSet.fromElements(new BigNumber(-769), new BigNumber((_dafny.Seq.of(new BigNumber(798))).length)))).Difference((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(-619)))).Union(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(752)), function (_11_i0) {
        return new _dafny.CodePoint('v'.codePointAt(0));
      })).length), new BigNumber(803), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_module.D1.create_DC5(false, false)).dtor_cf8,true)).length), new BigNumber(648))));
    };
    static fm12(p0, globalState) {
      if (((true) ? (true) : (true))) {
        return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(-402)), _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("mbafox")).length), new BigNumber((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(!(false),!(!(false))), _dafny.Map.Empty.slice().updateUnsafe(false,false), _dafny.Map.Empty.slice().updateUnsafe(true,true), _dafny.Map.Empty.slice().updateUnsafe(true,true))).length)));
      } else {
        return _dafny.Seq.of(new BigNumber(-643));
      }
    };
    static fm13(p0, p1, globalState) {
      return _module.D1.create_DC5(true, (_dafny.Set.fromElements(_dafny.Seq.of(false), _dafny.Seq.of(true))).IsSubsetOf(_dafny.Set.fromElements(_dafny.Seq.of(false), _dafny.Seq.of(!(true), true, false, true, true))));
    };
    static fm14(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(false,function () {
        let _coll8 = new _dafny.Map();
        for (const _compr_8 of _dafny.IntegerRange(new BigNumber(920), new BigNumber(101))) {
          let _12_v0 = _compr_8;
          if (((new BigNumber(920)).isLessThanOrEqualTo(_12_v0)) && ((_12_v0).isLessThan(new BigNumber(101)))) {
            _coll8.push([(_12_v0).plus(new BigNumber((_dafny.Seq.of(false)).length)),new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(621)), function (_13_i0) {
              return new BigNumber((_dafny.Seq.UnicodeFromString("yfuahe")).length);
            }))).cardinality())]);
          }
        }
        return _coll8;
      }())).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
        let _coll9 = new _dafny.Set();
        for (const _compr_9 of (_dafny.Seq.of(new BigNumber((_dafny.Seq.of(_module.D11.create_DC25(), _module.D11.create_DC25(), _module.D11.create_DC25())).length), new BigNumber(-411))).Elements) {
          let _14_v1 = _compr_9;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(_module.D11.create_DC25(), _module.D11.create_DC25(), _module.D11.create_DC25())).length), new BigNumber(-411)), _14_v1)) {
            _coll9.add(_module.__default.safeModuloInt(_14_v1, new BigNumber(280)));
          }
        }
        return _coll9;
      }()).length),new BigNumber(114)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(424),new BigNumber(154)))));
    };
    static fm15(p0, p1, p2, globalState) {
      return ((_dafny.Set.fromElements(true)).Difference(_dafny.Set.fromElements(false))).Difference((_dafny.Set.fromElements(false)).Intersect(_dafny.Set.fromElements(false)));
    };
    static fm16(p0, p1, globalState) {
      return _dafny.Seq.of(!(_dafny.areEqual(_dafny.Seq.of(new _dafny.CodePoint('g'.codePointAt(0))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(258)), function (_15_i0) {
        return new _dafny.CodePoint('w'.codePointAt(0));
      }))), true, (false) || (false));
    };
    static fm17(p0, p1, p2, p3, globalState) {
      return _module.D1.create_DC4((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(216)), function (_16_i0) {
  return _dafny.Seq.UnicodeFromString("nrkvxhch");
}))).Intersect(_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.UnicodeFromString("ugtro")))), new BigNumber(149));
    };
    static fm18(p0, globalState) {
      return ((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("odvbie"))).Intersect(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("efm"), _dafny.Seq.UnicodeFromString("ctebj")))).Difference((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("wmdxukk"), _dafny.Seq.UnicodeFromString("m"), _dafny.Seq.UnicodeFromString("dlv"))).Difference(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("affp"), _dafny.Seq.UnicodeFromString("v"))));
    };
    static fm19(globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("cjrbft")).length)),new BigNumber(587))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-296),new BigNumber(464))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(537),new BigNumber((_dafny.Seq.of(new BigNumber((function () {
        let _coll10 = new _dafny.Map();
        for (const _compr_10 of _dafny.IntegerRange(new BigNumber(932), new BigNumber(302))) {
          let _17_v0 = _compr_10;
          if (((new BigNumber(932)).isLessThanOrEqualTo(_17_v0)) && ((_17_v0).isLessThan(new BigNumber(302)))) {
            _coll10.push([_module.__default.safeDivisionInt(_17_v0, (_dafny.ZERO).minus(new BigNumber(-106))),false]);
          }
        }
        return _coll10;
      }()).length))).length))));
    };
    static fm20(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(!(false), true)).length),false)).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality())), _dafny.Seq.of(new BigNumber(795)), _dafny.Seq.of(new BigNumber(521)))).length),!(false))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(505),true)).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),new BigNumber((_dafny.Seq.of(!(true))).length))).length), new BigNumber(265))).cardinality()),!(false))));
    };
    static fm21(p0, p1, p2, p3, globalState) {
      let _source0 = _module.D0.create_DC1((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(908)), function (_18_i0) {
  return new _dafny.CodePoint('y'.codePointAt(0));
})).length)), new BigNumber(-506), new BigNumber(562));
      if (_source0.is_DC1) {
        let _19___mcc_h0 = (_source0).cf1;
        let _20___mcc_h1 = (_source0).cf2;
        let _21___mcc_h2 = (_source0).cf3;
        let _22_cf3 = _21___mcc_h2;
        let _23_cf2 = _20___mcc_h1;
        let _24_cf1 = _19___mcc_h0;
        return function () {
          let _coll11 = new _dafny.Map();
          for (const _compr_11 of (_dafny.Seq.of(_module.D1.create_DC6(_module.D1.create_DC3(new _dafny.CodePoint('f'.codePointAt(0)))), _module.D1.create_DC6(_module.D1.create_DC5(false, true)))).Elements) {
            let _25_v0 = _compr_11;
            if (_dafny.Seq.contains(_dafny.Seq.of(_module.D1.create_DC6(_module.D1.create_DC3(new _dafny.CodePoint('f'.codePointAt(0)))), _module.D1.create_DC6(_module.D1.create_DC5(false, true))), _25_v0)) {
              _coll11.push([_25_v0,new _dafny.CodePoint('y'.codePointAt(0))]);
            }
          }
          return _coll11;
        }();
      } else if (_source0.is_DC0) {
        let _26___mcc_h3 = (_source0).cf0;
        let _27_cf0 = _26___mcc_h3;
        return _dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC6(_module.D1.create_DC4(_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(888)), function (_28_i1) {
  return new _dafny.CodePoint('r'.codePointAt(0));
}))), new BigNumber(720))),new _dafny.CodePoint('k'.codePointAt(0)));
      } else {
        let _29___mcc_h4 = (_source0).cf4;
        let _30_cf4 = _29___mcc_h4;
        return (_module.D15.create_DC31(_dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC6(_module.D1.create_DC6(_module.D1.create_DC3(new _dafny.CodePoint('g'.codePointAt(0))))),new _dafny.CodePoint('r'.codePointAt(0))))).dtor_cf43;
      }
    };
    static m0(p0, p1, p2, p3, globalState) {
      let _31_v0;
      _31_v0 = false;
      let _32_v1;
      _32_v1 = _dafny.Seq.of(_31_v0, _31_v0, _31_v0);
      let _33_v2;
      _33_v2 = _dafny.Map.Empty.slice().updateUnsafe(_31_v0,_32_v1);
      _33_v2 = (_33_v2).update(true, _32_v1);
      if ((_32_v1)[_module.__default.safeIndex(p3, new BigNumber((_32_v1).length))]) {
        let _34_v3;
        _34_v3 = _module.D3.create_DC10(_31_v0, p3, _31_v0, p0, _dafny.Map.Empty.slice().updateUnsafe(p1,!(_31_v0)));
        let _35_v4;
        _35_v4 = _module.D3.create_DC11(_34_v3);
        _35_v4 = _35_v4;
        let _36_v5;
        _36_v5 = _module.D1.create_DC5(_31_v0, _31_v0);
        _31_v0 = (_36_v5).dtor_cf8;
        let _37_v6;
        _37_v6 = _module.D1.create_DC5(_31_v0, _31_v0);
        let _38_v7;
        _38_v7 = _module.D1.create_DC6(_37_v6);
        let _39_v8;
        let _nw0 = new _module.C0();
        _nw0.__ctor();
        _39_v8 = _nw0;
        let _40_v9;
        _40_v9 = _dafny.MultiSet.fromElements(_31_v0, true);
        let _41_v10;
        _41_v10 = _dafny.Seq.of(p1);
        let _42_v11;
        _42_v11 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_40_v9).cardinality())), p3, new BigNumber((_41_v10).length));
        let _43_v12;
        _43_v12 = _dafny.Seq.of(p1, new BigNumber((_dafny.Seq.update(_42_v11, _module.__default.safeIndex(p3, new BigNumber((_42_v11).length)), p1)).length));
        let _44_v13;
        let _nw1 = new _module.C2();
        _nw1.__ctor(p1);
        _44_v13 = _nw1;
        let _45_v14;
        _45_v14 = _dafny.Map.Empty.slice().updateUnsafe(p1,_44_v13);
        let _46_v15;
        _46_v15 = _dafny.Map.Empty.slice().updateUnsafe(_31_v0,new BigNumber((_45_v14).length));
        let _47_v16;
        let _nw2 = new _module.C1();
        _nw2.__ctor(false);
        _47_v16 = _nw2;
        let _48_v17;
        _48_v17 = _dafny.Map.Empty.slice().updateUnsafe(_31_v0,_47_v16);
        let _49_v18;
        let _nw3 = Array((new BigNumber(24)).toNumber());
        _nw3[(_dafny.ZERO).toNumber()] = p1;
        _nw3[(_dafny.ONE).toNumber()] = p3;
        _nw3[(new BigNumber(2)).toNumber()] = p3;
        _nw3[(new BigNumber(3)).toNumber()] = (_43_v12)[_module.__default.safeIndex(p1, new BigNumber((_43_v12).length))];
        _nw3[(new BigNumber(4)).toNumber()] = p1;
        _nw3[(new BigNumber(5)).toNumber()] = p3;
        _nw3[(new BigNumber(6)).toNumber()] = p1;
        _nw3[(new BigNumber(7)).toNumber()] = p3;
        _nw3[(new BigNumber(8)).toNumber()] = p1;
        _nw3[(new BigNumber(9)).toNumber()] = p3;
        _nw3[(new BigNumber(10)).toNumber()] = new BigNumber((_41_v10).length);
        _nw3[(new BigNumber(11)).toNumber()] = p3;
        _nw3[(new BigNumber(12)).toNumber()] = p3;
        _nw3[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus(p1);
        _nw3[(new BigNumber(14)).toNumber()] = _dafny.ONE;
        _nw3[(new BigNumber(15)).toNumber()] = p1;
        _nw3[(new BigNumber(16)).toNumber()] = (((_46_v15).contains(_31_v0)) ? ((_46_v15).get(_31_v0)) : (p3));
        _nw3[(new BigNumber(17)).toNumber()] = new BigNumber(597);
        _nw3[(new BigNumber(18)).toNumber()] = p3;
        _nw3[(new BigNumber(19)).toNumber()] = p1;
        _nw3[(new BigNumber(20)).toNumber()] = new BigNumber(-469);
        _nw3[(new BigNumber(21)).toNumber()] = _44_v13.f13;
        _nw3[(new BigNumber(22)).toNumber()] = new BigNumber(896);
        _nw3[(new BigNumber(23)).toNumber()] = new BigNumber((_48_v17).length);
        _49_v18 = _nw3;
        let _50_v19;
        _50_v19 = _dafny.Seq.of(_49_v18, _49_v18);
        let _51_v20;
        let _nw4 = Array((new BigNumber(22)).toNumber());
        _nw4[(_dafny.ZERO).toNumber()] = _49_v18;
        _nw4[(_dafny.ONE).toNumber()] = _49_v18;
        _nw4[(new BigNumber(2)).toNumber()] = _49_v18;
        _nw4[(new BigNumber(3)).toNumber()] = _49_v18;
        _nw4[(new BigNumber(4)).toNumber()] = _49_v18;
        _nw4[(new BigNumber(5)).toNumber()] = _49_v18;
        _nw4[(new BigNumber(6)).toNumber()] = _49_v18;
        _nw4[(new BigNumber(7)).toNumber()] = _49_v18;
        _nw4[(new BigNumber(8)).toNumber()] = (_50_v19)[_module.__default.safeIndex(p3, new BigNumber((_50_v19).length))];
        _nw4[(new BigNumber(9)).toNumber()] = _49_v18;
        _nw4[(new BigNumber(10)).toNumber()] = _49_v18;
        _nw4[(new BigNumber(11)).toNumber()] = _49_v18;
        _nw4[(new BigNumber(12)).toNumber()] = _49_v18;
        _nw4[(new BigNumber(13)).toNumber()] = _49_v18;
        _nw4[(new BigNumber(14)).toNumber()] = _49_v18;
        _nw4[(new BigNumber(15)).toNumber()] = _49_v18;
        _nw4[(new BigNumber(16)).toNumber()] = _49_v18;
        _nw4[(new BigNumber(17)).toNumber()] = _49_v18;
        _nw4[(new BigNumber(18)).toNumber()] = _49_v18;
        _nw4[(new BigNumber(19)).toNumber()] = _49_v18;
        _nw4[(new BigNumber(20)).toNumber()] = _49_v18;
        _nw4[(new BigNumber(21)).toNumber()] = _49_v18;
        _51_v20 = _nw4;
        let _rhs0 = _38_v7;
        let _rhs1 = _49_v18;
        let _rhs2 = _31_v0;
        let _rhs3 = _39_v8;
        let _rhs4 = _51_v20;
        let _lhs0 = globalState;
        _38_v7 = _rhs0;
        _lhs0.f10 = _rhs1;
        _31_v0 = _rhs2;
        _39_v8 = _rhs3;
        _51_v20 = _rhs4;
        let _52_v21;
        _52_v21 = _32_v1;
        let _53_v22;
        let _nw5 = Array((new BigNumber(13)).toNumber());
        _nw5[(_dafny.ZERO).toNumber()] = ((_47_v16.f14) ? (_32_v1) : (_32_v1));
        _nw5[(_dafny.ONE).toNumber()] = _32_v1;
        _nw5[(new BigNumber(2)).toNumber()] = _32_v1;
        _nw5[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_47_v16.f14, _47_v16.f14, _47_v16.f14, _47_v16.f14), _32_v1);
        _nw5[(new BigNumber(4)).toNumber()] = _dafny.Seq.of(_47_v16.f14);
        _nw5[(new BigNumber(5)).toNumber()] = _32_v1;
        _nw5[(new BigNumber(6)).toNumber()] = _32_v1;
        _nw5[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_32_v1, _32_v1);
        _nw5[(new BigNumber(8)).toNumber()] = _dafny.Seq.of(_31_v0, _47_v16.f14);
        _nw5[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_32_v1, _32_v1);
        _nw5[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_32_v1, _32_v1);
        _nw5[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat((_52_v21), _32_v1);
        _nw5[(new BigNumber(12)).toNumber()] = _32_v1;
        _53_v22 = _nw5;
        let _54_v23;
        _54_v23 = new _dafny.CodePoint('l'.codePointAt(0));
        let _rhs5 = _53_v22;
        let _rhs6 = _54_v23;
        let _rhs7 = ((_dafny.ZERO).minus(_44_v13.f13)).isLessThanOrEqualTo(p3);
        let _lhs1 = _47_v16;
        _53_v22 = _rhs5;
        _54_v23 = _rhs6;
        _lhs1.f14 = _rhs7;
        let _55_v24;
        _55_v24 = _dafny.Set.fromElements(_32_v1);
        _55_v24 = (_55_v24).Union(_55_v24);
      } else {
        let _index0 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((p0).length));
        (p0)[_index0] = _31_v0;
        let _56_v25;
        let _nw6 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.of());
        _56_v25 = _nw6;
        let _57_v26;
        _57_v26 = _dafny.Seq.of(p1, p3);
        let _index1 = _module.__default.safeIndex(new BigNumber(855), new BigNumber((_56_v25).length));
        (_56_v25)[_index1] = _57_v26;
        let _58_v27;
        _58_v27 = _dafny.Seq.UnicodeFromString("wwcctfdwq");
        let _59_v28;
        _59_v28 = _dafny.Map.Empty.slice().updateUnsafe(p3,!(_31_v0));
        let _60_v29;
        _60_v29 = _module.D7.create_DC17(_31_v0);
        let _61_v30;
        _61_v30 = _module.D11.create_DC23(_57_v26);
        let _index2 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((p0).length));
        let _index3 = _module.__default.safeIndex(new BigNumber(855), new BigNumber((_56_v25).length));
        let _rhs8 = (_dafny.Seq.contains(_57_v26, new BigNumber(233))) && (_dafny.Seq.IsPrefixOf(_58_v27, _dafny.Seq.UnicodeFromString("g")));
        let _rhs9 = new BigNumber(((_59_v28).Merge(_module.__default.fm20(_31_v0, (_dafny.ZERO).minus(p3), _module.__default.fm0(true, _31_v0, (_60_v29).dtor_cf25, _31_v0, globalState), globalState))).length);
        let _rhs10 = ((((!(_module.__default.fm0(_31_v0, _31_v0, _31_v0, _31_v0, globalState))) ? (_31_v0) : (false))) ? (_31_v0) : (((_31_v0) ? (_31_v0) : (_31_v0))));
        let _rhs11 = _dafny.Seq.update((_61_v30).dtor_cf34, _module.__default.safeIndex((p1).multipliedBy(p3), new BigNumber(((_61_v30).dtor_cf34).length)), (_dafny.ZERO).minus(p1));
        let _lhs2 = p0;
        let _lhs3 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((p0).length));
        let _lhs4 = globalState;
        let _lhs5 = _56_v25;
        let _lhs6 = _module.__default.safeIndex(new BigNumber(855), new BigNumber((_56_v25).length));
        _lhs2[_lhs3] = _rhs8;
        _lhs4.f2 = _rhs9;
        _31_v0 = _rhs10;
        _lhs5[_lhs6] = _rhs11;
        let _index4 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((p0).length));
        (p0)[_index4] = _31_v0;
        let _62_v31;
        let _nw7 = new _module.C1();
        _nw7.__ctor(_31_v0);
        _62_v31 = _nw7;
        let _63_v32;
        _63_v32 = _dafny.Map.Empty.slice().updateUnsafe((p1).isLessThanOrEqualTo(p1),_62_v31);
        let _64_v33;
        _64_v33 = _dafny.Map.Empty.slice().updateUnsafe((_57_v26)[_module.__default.safeIndex(p1, new BigNumber((_57_v26).length))],_62_v31);
        _62_v31 = (((_63_v32).contains((p0)[_module.__default.safeIndex(new BigNumber(491), new BigNumber((p0).length))])) ? ((_63_v32).get((p0)[_module.__default.safeIndex(new BigNumber(491), new BigNumber((p0).length))])) : ((((_64_v33).contains(p3)) ? ((_64_v33).get(p3)) : (_62_v31))));
        (globalState).f2 = ((_56_v25)[_module.__default.safeIndex(new BigNumber(855), new BigNumber((_56_v25).length))])[_module.__default.safeIndex(new BigNumber((_58_v27).length), new BigNumber(((_56_v25)[_module.__default.safeIndex(new BigNumber(855), new BigNumber((_56_v25).length))]).length))];
        let _65_v34;
        _65_v34 = _dafny.MultiSet.fromElements(_31_v0, _31_v0);
        let _66_v35;
        let _nw8 = new _module.C2();
        _nw8.__ctor(new BigNumber((_65_v34).cardinality()));
        _66_v35 = _nw8;
        let _67_v36;
        _67_v36 = new _dafny.CodePoint('x'.codePointAt(0));
        let _68_v37;
        _68_v37 = _dafny.Map.Empty.slice().updateUnsafe(_66_v35,_67_v36);
        (globalState).f2 = (p1).minus(new BigNumber((_68_v37).length));
      }
      if (_31_v0) {
        _31_v0 = _31_v0;
        let _69_v38;
        let _nw9 = new _module.C0();
        _nw9.__ctor();
        _69_v38 = _nw9;
        let _70_v39;
        _70_v39 = _dafny.Seq.of(_69_v38, _69_v38);
        let _71_v40;
        _71_v40 = _70_v39;
        let _72_v41;
        _72_v41 = _dafny.Map.Empty.slice().updateUnsafe((_71_v40),(_dafny.ZERO).minus(p1));
        let _73_v42;
        _73_v42 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(718),_module.__default.fm4(globalState));
        let _74_v43;
        _74_v43 = _dafny.Map.Empty.slice().updateUnsafe(_31_v0,_31_v0);
        let _75_v44;
        _75_v44 = _dafny.Seq.of(_74_v43);
        _72_v41 = (_72_v41).update(_dafny.Seq.update(_70_v39, _module.__default.safeIndex(p1, new BigNumber((_70_v39).length)), _69_v38), ((((_73_v42).contains(new BigNumber((_75_v44).length))) ? ((_73_v42).get(new BigNumber((_75_v44).length))) : (p3))).minus(p1));
        let _76_v45;
        _76_v45 = _dafny.Set.fromElements(p3);
        let _77_v46;
        _77_v46 = _dafny.Map.Empty.slice().updateUnsafe(!((_76_v45).IsProperSubsetOf(_dafny.Set.fromElements(p1))),p3);
        _77_v46 = (_77_v46).update(_module.__default.fm0(_31_v0, _31_v0, !(_31_v0), _31_v0, globalState), p3);
        let _78_v47;
        _78_v47 = _module.D11.create_DC24(p3);
        _78_v47 = _78_v47;
        if (!(_31_v0)) {
          let _79_v48;
          let _nw10 = new _module.C0();
          _nw10.__ctor();
          _79_v48 = _nw10;
          _31_v0 = _31_v0;
          let _80_v49;
          let _nw11 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          _80_v49 = _nw11;
          let _81_v50;
          let _nw12 = new _module.C2();
          _nw12.__ctor(p3);
          _81_v50 = _nw12;
          let _82_v51;
          _82_v51 = _dafny.Map.Empty.slice().updateUnsafe(_81_v50,_80_v49);
          let _83_v52;
          let _nw13 = Array((new BigNumber(25)).toNumber());
          _nw13[(_dafny.ZERO).toNumber()] = _80_v49;
          _nw13[(_dafny.ONE).toNumber()] = _80_v49;
          _nw13[(new BigNumber(2)).toNumber()] = _80_v49;
          _nw13[(new BigNumber(3)).toNumber()] = _80_v49;
          _nw13[(new BigNumber(4)).toNumber()] = _80_v49;
          _nw13[(new BigNumber(5)).toNumber()] = _80_v49;
          _nw13[(new BigNumber(6)).toNumber()] = _80_v49;
          _nw13[(new BigNumber(7)).toNumber()] = _80_v49;
          _nw13[(new BigNumber(8)).toNumber()] = _80_v49;
          _nw13[(new BigNumber(9)).toNumber()] = _80_v49;
          _nw13[(new BigNumber(10)).toNumber()] = _80_v49;
          _nw13[(new BigNumber(11)).toNumber()] = _80_v49;
          _nw13[(new BigNumber(12)).toNumber()] = _80_v49;
          _nw13[(new BigNumber(13)).toNumber()] = _80_v49;
          _nw13[(new BigNumber(14)).toNumber()] = (((_82_v51).contains(_81_v50)) ? ((_82_v51).get(_81_v50)) : (_80_v49));
          _nw13[(new BigNumber(15)).toNumber()] = _80_v49;
          _nw13[(new BigNumber(16)).toNumber()] = _80_v49;
          _nw13[(new BigNumber(17)).toNumber()] = _80_v49;
          _nw13[(new BigNumber(18)).toNumber()] = _80_v49;
          _nw13[(new BigNumber(19)).toNumber()] = _80_v49;
          _nw13[(new BigNumber(20)).toNumber()] = _80_v49;
          _nw13[(new BigNumber(21)).toNumber()] = _80_v49;
          _nw13[(new BigNumber(22)).toNumber()] = _80_v49;
          _nw13[(new BigNumber(23)).toNumber()] = _80_v49;
          _nw13[(new BigNumber(24)).toNumber()] = _80_v49;
          _83_v52 = _nw13;
          let _84_v53;
          _84_v53 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),_80_v49);
          let _85_v54;
          _85_v54 = new _dafny.CodePoint('g'.codePointAt(0));
          let _index5 = _module.__default.safeIndex(new BigNumber(704), new BigNumber((_83_v52).length));
          (_83_v52)[_index5] = (((_84_v53).contains(_85_v54)) ? ((_84_v53).get(_85_v54)) : (_80_v49));
          let _index6 = _module.__default.safeIndex(new BigNumber(704), new BigNumber((_83_v52).length));
          (_83_v52)[_index6] = _80_v49;
          let _86_v55;
          _86_v55 = _dafny.Map.Empty.slice().updateUnsafe(_69_v38,new BigNumber((_77_v46).length));
          _86_v55 = (_86_v55).update(_69_v38, p3);
          (globalState).f2 = _81_v50.f13;
        } else {
          let _87_v56;
          let _init0 = ((_88_p3) => function (_89_i0) {
            return _module.__default.safeModuloInt(_89_i0, _88_p3);
          })(p3);
          let _nw14 = Array((new BigNumber(12)).toNumber());
          for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw14.length); _i0_0++) {
            _nw14[_i0_0] = _init0(new BigNumber(_i0_0));
          }
          _87_v56 = _nw14;
          let _90_v57;
          _90_v57 = _dafny.Map.Empty.slice().updateUnsafe(_87_v56,p1);
          let _91_v58;
          _91_v58 = _dafny.MultiSet.fromElements(p3, p3);
          let _92_v59;
          _92_v59 = _dafny.Set.fromElements(_31_v0, _31_v0, _31_v0);
          _73_v42 = (_73_v42).update((((_90_v57).contains(_87_v56)) ? ((_90_v57).get(_87_v56)) : ((_69_v38).fm8(new BigNumber((_91_v58).cardinality()), _92_v59, p1, globalState))), p1);
          let _93_v60;
          let _init1 = function (_94_i1) {
            return _module.D1.create_DC6(_module.D1.create_DC6(_module.D1.create_DC3(new _dafny.CodePoint('h'.codePointAt(0)))));
          };
          let _nw15 = Array((new BigNumber(2)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw15.length); _i0_1++) {
            _nw15[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _93_v60 = _nw15;
          let _95_v61;
          _95_v61 = _module.D2.create_DC7(_93_v60);
          let _96_v62;
          _96_v62 = new _dafny.CodePoint('m'.codePointAt(0));
          let _rhs12 = (p1).minus(_module.__default.safeDivisionInt(p3, new BigNumber((_module.__default.fm11(_31_v0, _96_v62, globalState)).cardinality())));
          let _rhs13 = ((_dafny.ZERO).minus(p1)).isLessThanOrEqualTo((_dafny.ZERO).minus(p3));
          let _rhs14 = _95_v61;
          let _lhs7 = globalState;
          _lhs7.f2 = _rhs12;
          _31_v0 = _rhs13;
          _95_v61 = _rhs14;
          (globalState).f2 = (_dafny.ZERO).minus((_dafny.ZERO).minus(p1));
          let _97_v63;
          let _nw16 = new _module.C2();
          _nw16.__ctor(p1);
          _97_v63 = _nw16;
          let _index7 = _module.__default.safeIndex(new BigNumber(895), new BigNumber((_87_v56).length));
          (_87_v56)[_index7] = _module.__default.fm2(_31_v0, globalState);
          let _index8 = _module.__default.safeIndex(new BigNumber(895), new BigNumber((_87_v56).length));
          (_87_v56)[_index8] = _97_v63.f13;
        }
      } else {
        let _98_v64;
        _98_v64 = _module.D7.create_DC18(_module.D7.create_DC17(_31_v0));
        let _source1 = _98_v64;
        if (_source1.is_DC17) {
          let _99___mcc_h0 = (_source1).cf25;
          let _100_cf25 = _99___mcc_h0;
          _100_cf25 = (!(p1).isEqualTo(p3)) && (!(_31_v0));
          let _101_v65;
          _101_v65 = _dafny.Map.Empty.slice().updateUnsafe(p1,_100_cf25);
          let _102_v66;
          _102_v66 = _dafny.Seq.UnicodeFromString("dpwnr");
          let _103_v67;
          _103_v67 = _dafny.Map.Empty.slice().updateUnsafe(!((_module.D3.create_DC10(_31_v0, p3, _31_v0, p0, _101_v65)).dtor_cf16),_102_v66);
          let _104_v68;
          _104_v68 = _dafny.Seq.of(new BigNumber((_103_v67).length));
          let _rhs15 = (_104_v68)[_module.__default.safeIndex(new BigNumber(808), new BigNumber((_104_v68).length))];
          let _rhs16 = p1;
          let _lhs8 = globalState;
          let _lhs9 = globalState;
          _lhs8.f2 = _rhs15;
          _lhs9.f2 = _rhs16;
          let _105_v69;
          let _nw17 = new _module.C2();
          _nw17.__ctor(((_31_v0) ? (p1) : (p1)));
          _105_v69 = _nw17;
          _105_v69 = _105_v69;
          let _106_v70;
          let _nw18 = new _module.C3();
          _nw18.__ctor(_100_cf25, new BigNumber(7));
          _106_v70 = _nw18;
        } else if (_source1.is_DC16) {
          let _107___mcc_h1 = (_source1).cf24;
          let _108_cf24 = _107___mcc_h1;
          let _109_v71;
          let _init2 = function (_110_i2) {
            return _dafny.Seq.of(new _dafny.CodePoint('q'.codePointAt(0)), new _dafny.CodePoint('q'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)));
          };
          let _nw19 = Array((new BigNumber(23)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw19.length); _i0_2++) {
            _nw19[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _109_v71 = _nw19;
          let _111_v72;
          let _init3 = ((_112_p3) => function (_113_i3) {
            return (_113_i3).minus((_dafny.ZERO).minus(_112_p3));
          })(p3);
          let _nw20 = Array((new BigNumber(14)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw20.length); _i0_3++) {
            _nw20[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _111_v72 = _nw20;
          let _114_v73;
          _114_v73 = _dafny.Seq.UnicodeFromString("joaugkus");
          let _index9 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_111_v72).length));
          (_111_v72)[_index9] = new BigNumber((_114_v73).length);
          let _115_v74;
          _115_v74 = _dafny.Map.Empty.slice().updateUnsafe(!(p1).isEqualTo(p3),_dafny.Seq.Concat(_114_v73, _114_v73));
          let _116_v75;
          let _nw21 = Array((new BigNumber(16)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _116_v75 = _nw21;
          let _index10 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_111_v72).length));
          let _rhs17 = ((_31_v0) ? (_109_v71) : (_109_v71));
          let _rhs18 = p1;
          let _rhs19 = _109_v71;
          let _rhs20 = (_115_v74).Merge(_115_v74);
          let _rhs21 = _116_v75;
          let _lhs10 = _111_v72;
          let _lhs11 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_111_v72).length));
          _109_v71 = _rhs17;
          _lhs10[_lhs11] = _rhs18;
          _109_v71 = _rhs19;
          _115_v74 = _rhs20;
          _116_v75 = _rhs21;
          let _117_v76;
          _117_v76 = _dafny.Set.fromElements(_31_v0);
          let _index11 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((p0).length));
          (p0)[_index11] = (_117_v76).IsDisjointFrom(_117_v76);
          let _118_v77;
          let _init4 = ((_119_p3, _120_p1, _121_v72) => function (_122_i4) {
            return (_dafny.Map.Empty.slice().updateUnsafe(_119_p3,_120_p1)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_121_v72)[_module.__default.safeIndex(new BigNumber(235), new BigNumber((_121_v72).length))],_119_p3));
          })(p3, p1, _111_v72);
          let _nw22 = Array((new BigNumber(2)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw22.length); _i0_4++) {
            _nw22[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _118_v77 = _nw22;
          let _123_v78;
          _123_v78 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("nufq")).length),p1);
          let _index12 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_118_v77).length));
          (_118_v77)[_index12] = _123_v78;
          let _124_v79;
          _124_v79 = _dafny.MultiSet.fromElements(p1);
          let _index13 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((p0).length));
          let _index14 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_118_v77).length));
          let _index15 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_111_v72).length));
          let _rhs22 = (p1).minus(p3);
          let _rhs23 = (_124_v79).IsSubsetOf(_124_v79);
          let _rhs24 = _123_v78;
          let _rhs25 = p1;
          let _lhs12 = globalState;
          let _lhs13 = p0;
          let _lhs14 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((p0).length));
          let _lhs15 = _118_v77;
          let _lhs16 = _module.__default.safeIndex(new BigNumber(296), new BigNumber((_118_v77).length));
          let _lhs17 = _111_v72;
          let _lhs18 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_111_v72).length));
          _lhs12.f2 = _rhs22;
          _lhs13[_lhs14] = _rhs23;
          _lhs15[_lhs16] = _rhs24;
          _lhs17[_lhs18] = _rhs25;
          let _125_v80;
          let _nw23 = new _module.C1();
          _nw23.__ctor((_module.D7.create_DC17(false)).dtor_cf25);
          _125_v80 = _nw23;
          let _126_v81;
          _126_v81 = _dafny.Map.Empty.slice().updateUnsafe((p0)[_module.__default.safeIndex(new BigNumber(979), new BigNumber((p0).length))],_125_v80);
          let _index16 = _module.__default.safeIndex(new BigNumber(979), new BigNumber((p0).length));
          (p0)[_index16] = !(!((new BigNumber(((_126_v81).Merge(_126_v81)).length)).isLessThan(new BigNumber((_dafny.Seq.Concat(_114_v73, _114_v73)).length))));
          (globalState).f2 = (_dafny.ZERO).minus((((p0)[_module.__default.safeIndex(new BigNumber(979), new BigNumber((p0).length))]) ? ((_dafny.ZERO).minus(((_111_v72)[_module.__default.safeIndex(new BigNumber(235), new BigNumber((_111_v72).length))]).plus(new BigNumber((_114_v73).length)))) : ((_dafny.ZERO).minus(new BigNumber(((_117_v76).Union(_117_v76)).length)))));
        } else {
          let _127___mcc_h2 = (_source1).cf26;
          let _128_cf26 = _127___mcc_h2;
          let _129_v82;
          _129_v82 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("gwvsosvy"));
          let _130_v83;
          _130_v83 = _module.D1.create_DC4(_129_v82, p1);
          let _131_v84;
          _131_v84 = _dafny.Seq.UnicodeFromString("qcm");
          let _pat_let_tv0 = _129_v82;
          let _pat_let_tv1 = _131_v84;
          let _pat_let_tv2 = _131_v84;
          _130_v83 = function (_pat_let0_0) {
            return function (_132_dt__update__tmp_h0) {
              return function (_pat_let1_0) {
                return function (_133_dt__update_hcf6_h0) {
                  return _module.D1.create_DC4(_133_dt__update_hcf6_h0, (_132_dt__update__tmp_h0).dtor_cf7);
                }(_pat_let1_0);
              }((_pat_let_tv0).Intersect(_dafny.MultiSet.fromElements(_pat_let_tv1, _pat_let_tv2)));
            }(_pat_let0_0);
          }(_130_v83);
          (globalState).f2 = _module.__default.safeDivisionInt(p1, p1);
          let _134_v85;
          let _nw24 = new _module.C3();
          _nw24.__ctor(_dafny.Seq.IsPrefixOf(_module.__default.fm9(_31_v0, globalState), _131_v84), p1);
          _134_v85 = _nw24;
          _31_v0 = _31_v0;
        }
        let _135_v86;
        _135_v86 = _dafny.MultiSet.fromElements(false);
        let _136_v87;
        _136_v87 = _dafny.MultiSet.fromElements(new BigNumber((_135_v86).cardinality()), p3);
        let _137_v88;
        _137_v88 = _dafny.Seq.UnicodeFromString("fpbqp");
        let _138_v89;
        _138_v89 = _dafny.Map.Empty.slice().updateUnsafe(_137_v88,_module.__default.fm0(_31_v0, _31_v0, _31_v0, _31_v0, globalState));
        let _139_v90;
        _139_v90 = _dafny.Set.fromElements(((((_136_v87).contains(p3)) ? ((_136_v87).get(p3)) : (p3))).isLessThan(new BigNumber(383)), _31_v0, (((_138_v89).contains(_137_v88)) ? ((_138_v89).get(_137_v88)) : (_31_v0)));
        _139_v90 = (_dafny.Set.fromElements(_31_v0, _31_v0, false)).Intersect((_139_v90).Intersect(_139_v90));
        (globalState).f2 = (((_136_v87).contains(new BigNumber(530))) ? ((_136_v87).get(new BigNumber(530))) : (new BigNumber((_dafny.MultiSet.fromElements(_31_v0, _31_v0)).cardinality())));
        let _140_v91;
        let _nw25 = new _module.C1();
        _nw25.__ctor(_31_v0);
        _140_v91 = _nw25;
        _140_v91 = _140_v91;
        let _141_v92;
        _141_v92 = _dafny.MultiSet.fromElements(_137_v88, _137_v88);
        let _142_v93;
        _142_v93 = _dafny.Set.fromElements((((_141_v92).contains(_137_v88)) ? ((_141_v92).get(_137_v88)) : (p1)), new BigNumber(311), p1, p1);
        _142_v93 = _142_v93;
      }
      let _143_v94;
      _143_v94 = _module.D1.create_DC5(_31_v0, _module.__default.fm0(_31_v0, _31_v0, false, _31_v0, globalState));
      _31_v0 = !((_143_v94).dtor_cf9);
      let _144_v95;
      _144_v95 = _dafny.Seq.UnicodeFromString("jfta");
      let _145_v96;
      _145_v96 = _dafny.Map.Empty.slice().updateUnsafe(_31_v0,_144_v95);
      let _146_v97;
      _146_v97 = _dafny.MultiSet.fromElements(p3);
      if ((_dafny.MultiSet.fromElements(new BigNumber((_145_v96).length))).IsProperSubsetOf(_146_v97)) {
        let _147_v98;
        let _nw26 = Array((new BigNumber(3)).toNumber()).fill(false);
        _147_v98 = _nw26;
        _147_v98 = _147_v98;
        let _148_v99;
        let _nw27 = new _module.C1();
        _nw27.__ctor(_31_v0);
        _148_v99 = _nw27;
        let _149_v100;
        let _nw28 = new _module.C3();
        _nw28.__ctor(_31_v0, p1);
        _149_v100 = _nw28;
        let _150_v101;
        _150_v101 = _dafny.Seq.of(_149_v100);
        let _rhs26 = _dafny.Seq.Concat(_32_v1, _32_v1);
        let _rhs27 = _148_v99.f14;
        let _rhs28 = _32_v1;
        let _rhs29 = p3;
        let _rhs30 = (p3).plus(new BigNumber((_150_v101).length));
        let _lhs19 = globalState;
        let _lhs20 = globalState;
        _32_v1 = _rhs26;
        _31_v0 = _rhs27;
        _32_v1 = _rhs28;
        _lhs19.f2 = _rhs29;
        _lhs20.f2 = _rhs30;
        let _151_v102;
        _151_v102 = new _dafny.CodePoint('v'.codePointAt(0));
        let _152_v103;
        let _nw29 = Array((new BigNumber(21)).toNumber());
        _nw29[(_dafny.ZERO).toNumber()] = _151_v102;
        _nw29[(_dafny.ONE).toNumber()] = _151_v102;
        _nw29[(new BigNumber(2)).toNumber()] = _151_v102;
        _nw29[(new BigNumber(3)).toNumber()] = _151_v102;
        _nw29[(new BigNumber(4)).toNumber()] = ((_module.__default.fm0(_148_v99.f14, (_149_v100).f11, _148_v99.f14, _148_v99.f14, globalState)) ? (_151_v102) : (_151_v102));
        _nw29[(new BigNumber(5)).toNumber()] = _151_v102;
        _nw29[(new BigNumber(6)).toNumber()] = _151_v102;
        _nw29[(new BigNumber(7)).toNumber()] = ((_31_v0) ? (_151_v102) : (_151_v102));
        _nw29[(new BigNumber(8)).toNumber()] = _151_v102;
        _nw29[(new BigNumber(9)).toNumber()] = _151_v102;
        _nw29[(new BigNumber(10)).toNumber()] = _151_v102;
        _nw29[(new BigNumber(11)).toNumber()] = _151_v102;
        _nw29[(new BigNumber(12)).toNumber()] = ((_31_v0) ? (_151_v102) : (new _dafny.CodePoint('x'.codePointAt(0))));
        _nw29[(new BigNumber(13)).toNumber()] = _151_v102;
        _nw29[(new BigNumber(14)).toNumber()] = _151_v102;
        _nw29[(new BigNumber(15)).toNumber()] = _151_v102;
        _nw29[(new BigNumber(16)).toNumber()] = _151_v102;
        _nw29[(new BigNumber(17)).toNumber()] = _151_v102;
        _nw29[(new BigNumber(18)).toNumber()] = _151_v102;
        _nw29[(new BigNumber(19)).toNumber()] = _151_v102;
        _nw29[(new BigNumber(20)).toNumber()] = new _dafny.CodePoint('y'.codePointAt(0));
        _152_v103 = _nw29;
        let _nw30 = Array((new BigNumber(8)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _152_v103 = _nw30;
        (_148_v99).f14 = (p1).isEqualTo(new BigNumber((_32_v1).length));
      } else {
        let _153_v104;
        let _nw31 = new _module.C0();
        _nw31.__ctor();
        _153_v104 = _nw31;
        (globalState).f2 = p3;
        let _154_v105;
        let _nw32 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
        _154_v105 = _nw32;
        let _155_v106;
        _155_v106 = _dafny.Seq.of(_154_v105, _154_v105);
        let _156_v107;
        _156_v107 = _dafny.Set.fromElements(_31_v0);
        let _157_v108;
        let _nw33 = Array((new BigNumber(19)).toNumber());
        _nw33[(_dafny.ZERO).toNumber()] = _154_v105;
        _nw33[(_dafny.ONE).toNumber()] = _154_v105;
        _nw33[(new BigNumber(2)).toNumber()] = _154_v105;
        _nw33[(new BigNumber(3)).toNumber()] = _154_v105;
        _nw33[(new BigNumber(4)).toNumber()] = _154_v105;
        _nw33[(new BigNumber(5)).toNumber()] = _154_v105;
        _nw33[(new BigNumber(6)).toNumber()] = (_155_v106)[_module.__default.safeIndex(new BigNumber((_156_v107).length), new BigNumber((_155_v106).length))];
        _nw33[(new BigNumber(7)).toNumber()] = _154_v105;
        _nw33[(new BigNumber(8)).toNumber()] = _154_v105;
        _nw33[(new BigNumber(9)).toNumber()] = _154_v105;
        _nw33[(new BigNumber(10)).toNumber()] = _154_v105;
        _nw33[(new BigNumber(11)).toNumber()] = _154_v105;
        _nw33[(new BigNumber(12)).toNumber()] = _154_v105;
        _nw33[(new BigNumber(13)).toNumber()] = _154_v105;
        _nw33[(new BigNumber(14)).toNumber()] = _154_v105;
        _nw33[(new BigNumber(15)).toNumber()] = _154_v105;
        _nw33[(new BigNumber(16)).toNumber()] = _154_v105;
        _nw33[(new BigNumber(17)).toNumber()] = _154_v105;
        _nw33[(new BigNumber(18)).toNumber()] = _154_v105;
        _157_v108 = _nw33;
        let _index17 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_157_v108).length));
        (_157_v108)[_index17] = _154_v105;
        let _158_v109;
        _158_v109 = _dafny.MultiSet.fromElements(_144_v95);
        let _159_v110;
        _159_v110 = _module.D1.create_DC6(_module.D1.create_DC4(_158_v109, p1));
        let _160_v111;
        _160_v111 = _module.D1.create_DC3(new _dafny.CodePoint('q'.codePointAt(0)));
        let _161_v112;
        _161_v112 = new _dafny.CodePoint('a'.codePointAt(0));
        let _pat_let_tv3 = _160_v111;
        let _162_v113;
        _162_v113 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let2_0) {
          return function (_163_dt__update__tmp_h1) {
            return function (_pat_let3_0) {
              return function (_164_dt__update_hcf10_h0) {
                return _module.D1.create_DC6(_164_dt__update_hcf10_h0);
              }(_pat_let3_0);
            }(_pat_let_tv3);
          }(_pat_let2_0);
        }(_159_v110),_161_v112);
        let _index18 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_157_v108).length));
        let _rhs31 = _154_v105;
        let _rhs32 = _154_v105;
        let _rhs33 = ((_module.__default.fm21(_31_v0, _31_v0, new BigNumber(-627), _31_v0, globalState)).Merge(_162_v113)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_159_v110,new _dafny.CodePoint('w'.codePointAt(0))));
        let _lhs21 = globalState;
        let _lhs22 = _157_v108;
        let _lhs23 = _module.__default.safeIndex(new BigNumber(744), new BigNumber((_157_v108).length));
        _lhs21.f10 = _rhs31;
        _lhs22[_lhs23] = _rhs32;
        _162_v113 = _rhs33;
        _31_v0 = !(_31_v0) || (_31_v0);
        _31_v0 = ((_31_v0) ? (_31_v0) : (!(_31_v0)));
      }
      if (!_dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(789)), function (_165_i5) {
        return new _dafny.CodePoint('k'.codePointAt(0));
      }), _144_v95), _144_v95)) {
        let _166_v114;
        _166_v114 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        let _167_v115;
        _167_v115 = _dafny.Map.Empty.slice().updateUnsafe(p3,_31_v0);
        _31_v0 = _module.__default.fm0(_31_v0, ((_dafny.ZERO).minus(p3)).isLessThanOrEqualTo(p3), !(p1).isEqualTo(new BigNumber((_166_v114).length)), !((((_167_v115).contains(new BigNumber(61))) ? ((_167_v115).get(new BigNumber(61))) : (_31_v0))), globalState);
        _31_v0 = _31_v0;
        let _168_v116;
        let _nw34 = new _module.C2();
        _nw34.__ctor(p3);
        _168_v116 = _nw34;
        let _169_v117;
        _169_v117 = _module.D8.create_DC19(_168_v116);
        _169_v117 = (((_168_v116.f13).isLessThanOrEqualTo(p3)) ? (_169_v117) : (_169_v117));
        let _170_v118;
        let _init5 = ((_171_v95, _172_v116) => function (_173_i6) {
          return _dafny.Seq.Concat(_dafny.Seq.update(_171_v95, _module.__default.safeIndex(_172_v116.f13, new BigNumber((_171_v95).length)), new _dafny.CodePoint('o'.codePointAt(0))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(429)), function (_174_i7) {
            return new _dafny.CodePoint('v'.codePointAt(0));
          }));
        })(_144_v95, _168_v116);
        let _nw35 = Array((new BigNumber(23)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw35.length); _i0_5++) {
          _nw35[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _170_v118 = _nw35;
        let _index19 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((_170_v118).length));
        (_170_v118)[_index19] = _dafny.Seq.Concat(_144_v95, _module.__default.fm9(_31_v0, globalState));
        let _index20 = _module.__default.safeIndex(new BigNumber(301), new BigNumber((_170_v118).length));
        (_170_v118)[_index20] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-589)), function (_175_i8) {
          return new _dafny.CodePoint('v'.codePointAt(0));
        });
        let _176_v119;
        _176_v119 = _dafny.Set.fromElements(_168_v116.f13);
        let _177_v120;
        _177_v120 = _dafny.Seq.of(_176_v119);
        _31_v0 = (_176_v119).IsProperSubsetOf((_177_v120)[_module.__default.safeIndex((_dafny.ZERO).minus(_168_v116.f13), new BigNumber((_177_v120).length))]);
      } else {
        let _178_v121;
        let _nw36 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Map.Empty);
        _178_v121 = _nw36;
        let _179_v122;
        _179_v122 = _dafny.Map.Empty.slice().updateUnsafe(_31_v0,p1);
        let _index21 = _module.__default.safeIndex(new BigNumber(978), new BigNumber((_178_v121).length));
        (_178_v121)[_index21] = (_dafny.Map.Empty.slice().updateUnsafe(_31_v0,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(107)), function (_180_i9) {
          return new _dafny.CodePoint('n'.codePointAt(0));
        })).length))).Merge(_179_v122);
        let _181_v123;
        _181_v123 = _dafny.Seq.of(p1);
        let _182_v124;
        _182_v124 = _dafny.Set.fromElements(_31_v0);
        let _183_v125;
        _183_v125 = _dafny.Seq.of(_182_v124, _182_v124);
        let _184_v127;
        let _nw37 = Array((new BigNumber(17)).toNumber());
        _nw37[(_dafny.ZERO).toNumber()] = p3;
        _nw37[(_dafny.ONE).toNumber()] = p3;
        _nw37[(new BigNumber(2)).toNumber()] = p3;
        _nw37[(new BigNumber(3)).toNumber()] = p3;
        _nw37[(new BigNumber(4)).toNumber()] = new BigNumber((_181_v123).length);
        _nw37[(new BigNumber(5)).toNumber()] = new BigNumber((_182_v124).length);
        _nw37[(new BigNumber(6)).toNumber()] = p3;
        _nw37[(new BigNumber(7)).toNumber()] = new BigNumber((_146_v97).cardinality());
        _nw37[(new BigNumber(8)).toNumber()] = new BigNumber((function () {
          let _coll12 = new _dafny.Map();
          for (const _compr_12 of _dafny.IntegerRange(new BigNumber(175), new BigNumber(637))) {
            let _185_v126 = _compr_12;
            if (((new BigNumber(175)).isLessThanOrEqualTo(_185_v126)) && ((_185_v126).isLessThan(new BigNumber(637)))) {
              _coll12.push([(_185_v126).minus(p1),_31_v0]);
            }
          }
          return _coll12;
        }()).length);
        _nw37[(new BigNumber(9)).toNumber()] = p3;
        _nw37[(new BigNumber(10)).toNumber()] = new BigNumber((_146_v97).cardinality());
        _nw37[(new BigNumber(11)).toNumber()] = p1;
        _nw37[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus(p1);
        _nw37[(new BigNumber(13)).toNumber()] = p3;
        _nw37[(new BigNumber(14)).toNumber()] = p3;
        _nw37[(new BigNumber(15)).toNumber()] = p1;
        _nw37[(new BigNumber(16)).toNumber()] = p3;
        _184_v127 = _nw37;
        let _186_v128;
        _186_v128 = _dafny.Seq.of(_184_v127, _184_v127);
        let _187_v129;
        _187_v129 = _module.D11.create_DC26(p3, (_186_v128)[_module.__default.safeIndex(p1, new BigNumber((_186_v128).length))], _31_v0);
        let _188_v130;
        let _nw38 = Array((new BigNumber(15)).toNumber());
        _nw38[(_dafny.ZERO).toNumber()] = _module.__default.fm15(p3, new BigNumber((_181_v123).length), _module.__default.fm0(_module.__default.fm0(_31_v0, _31_v0, _31_v0, !(_31_v0), globalState), _31_v0, _31_v0, !(_31_v0), globalState), globalState);
        _nw38[(_dafny.ONE).toNumber()] = _182_v124;
        _nw38[(new BigNumber(2)).toNumber()] = (_183_v125)[_module.__default.safeIndex((_181_v123)[_module.__default.safeIndex(p3, new BigNumber((_181_v123).length))], new BigNumber((_183_v125).length))];
        _nw38[(new BigNumber(3)).toNumber()] = _182_v124;
        _nw38[(new BigNumber(4)).toNumber()] = (_182_v124).Intersect(_182_v124);
        _nw38[(new BigNumber(5)).toNumber()] = _182_v124;
        _nw38[(new BigNumber(6)).toNumber()] = ((_31_v0) ? (_182_v124) : (_module.__default.fm15(p1, p1, _31_v0, globalState)));
        _nw38[(new BigNumber(7)).toNumber()] = _182_v124;
        _nw38[(new BigNumber(8)).toNumber()] = (_182_v124).Union(_dafny.Set.fromElements(!(_31_v0), (_187_v129).dtor_cf38));
        _nw38[(new BigNumber(9)).toNumber()] = _182_v124;
        _nw38[(new BigNumber(10)).toNumber()] = _module.__default.fm15(p1, p1, _31_v0, globalState);
        _nw38[(new BigNumber(11)).toNumber()] = _182_v124;
        _nw38[(new BigNumber(12)).toNumber()] = _dafny.Set.fromElements(false);
        _nw38[(new BigNumber(13)).toNumber()] = _182_v124;
        _nw38[(new BigNumber(14)).toNumber()] = _182_v124;
        _188_v130 = _nw38;
        let _index22 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_188_v130).length));
        (_188_v130)[_index22] = _182_v124;
        let _189_v131;
        _189_v131 = _dafny.Set.fromElements(p1, p3);
        let _index23 = _module.__default.safeIndex(new BigNumber(978), new BigNumber((_178_v121).length));
        let _index24 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_188_v130).length));
        let _rhs34 = _179_v122;
        let _rhs35 = _dafny.Set.fromElements((_189_v131).IsDisjointFrom(_189_v131));
        let _lhs24 = _178_v121;
        let _lhs25 = _module.__default.safeIndex(new BigNumber(978), new BigNumber((_178_v121).length));
        let _lhs26 = _188_v130;
        let _lhs27 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_188_v130).length));
        _lhs24[_lhs25] = _rhs34;
        _lhs26[_lhs27] = _rhs35;
        let _190_v132;
        _190_v132 = _module.D0.create_DC0(_31_v0);
        _31_v0 = _module.__default.fm0(!((_190_v132).dtor_cf0) || (_31_v0), _module.__default.fm0(_31_v0, _31_v0, _31_v0, _31_v0, globalState), ((_dafny.ZERO).minus(p3)).isLessThan(p3), _module.__default.fm0(_31_v0, _module.__default.fm0(_31_v0, _31_v0, _31_v0, _31_v0, globalState), _31_v0, _31_v0, globalState), globalState);
        _31_v0 = !(true);
        let _191_v133;
        let _nw39 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.of());
        _191_v133 = _nw39;
        let _index25 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_191_v133).length));
        (_191_v133)[_index25] = _181_v123;
        let _index26 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_191_v133).length));
        (_191_v133)[_index26] = _181_v123;
        let _index27 = _module.__default.safeIndex(new BigNumber(45), new BigNumber((_184_v127).length));
        (_184_v127)[_index27] = p3;
        let _index28 = _module.__default.safeIndex(new BigNumber(45), new BigNumber((_184_v127).length));
        (_184_v127)[_index28] = (_dafny.ZERO).minus(p1);
      }
      return;
    }
    static Main(__noArgsParameter) {
      let _192_v0;
      _192_v0 = true;
      let _193_v1;
      _193_v1 = _dafny.Set.fromElements(!(!(_192_v0)));
      let _194_v2;
      _194_v2 = _dafny.Set.fromElements(_dafny.Set.fromElements(_192_v0, !(!(_192_v0))), _193_v1);
      let _195_v3;
      _195_v3 = new BigNumber(848);
      let _196_v4;
      _196_v4 = _dafny.Seq.UnicodeFromString("bk");
      let _197_v5;
      _197_v5 = _dafny.Map.Empty.slice().updateUnsafe(_192_v0,_196_v4);
      let _198_v6;
      _198_v6 = _dafny.Seq.of(new BigNumber((_197_v5).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_195_v3,_195_v3)).length), _195_v3);
      let _199_v7;
      _199_v7 = _dafny.Seq.of(_193_v1, _193_v1, _193_v1);
      let _200_v8;
      let _nw40 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
      _200_v8 = _nw40;
      let _201_v9;
      _201_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_199_v7)[_module.__default.safeIndex(_195_v3, new BigNumber((_199_v7).length))]).length),_200_v8);
      let _202_v10;
      _202_v10 = _dafny.MultiSet.fromElements(!(_192_v0));
      let _203_globalState;
      let _nw41 = new _module.GlobalState();
      _nw41.__ctor(new BigNumber(245), false, new BigNumber(862), false, (_dafny.Set.fromElements(_193_v1, _193_v1)).Difference(_194_v2), ((_192_v0) ? (_dafny.Map.Empty.slice().updateUnsafe(_195_v3,_192_v0)) : (_dafny.Map.Empty.slice().updateUnsafe(_195_v3,!(_192_v0)))), _198_v6, _198_v6, false, true, (((_201_v9).contains((((_202_v10).contains(_192_v0)) ? ((_202_v10).get(_192_v0)) : (_195_v3)))) ? ((_201_v9).get((((_202_v10).contains(_192_v0)) ? ((_202_v10).get(_192_v0)) : (_195_v3)))) : (_200_v8)));
      _203_globalState = _nw41;
      (_203_globalState).f2 = (_module.__default.safeModuloInt(new BigNumber(-661), (_dafny.ZERO).minus(_195_v3))).multipliedBy(_195_v3);
      let _204_v11;
      _204_v11 = _dafny.Map.Empty.slice().updateUnsafe(_192_v0,(_195_v3).isLessThanOrEqualTo(new BigNumber(758)));
      let _205_v12;
      _205_v12 = _dafny.Map.Empty.slice().updateUnsafe(_192_v0,_195_v3);
      _204_v11 = (_204_v11).update((_195_v3).isLessThanOrEqualTo(_195_v3), (_195_v3).isEqualTo(new BigNumber((_205_v12).length)));
      let _206_v13;
      let _init6 = ((_207_v0) => function (_208_i0) {
        return !(_207_v0);
      })(_192_v0);
      let _nw42 = Array((new BigNumber(15)).toNumber());
      for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw42.length); _i0_6++) {
        _nw42[_i0_6] = _init6(new BigNumber(_i0_6));
      }
      _206_v13 = _nw42;
      let _209_v14;
      let _init7 = ((_210_v11) => function (_211_i1) {
        return _210_v11;
      })(_204_v11);
      let _nw43 = Array((new BigNumber(14)).toNumber());
      for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw43.length); _i0_7++) {
        _nw43[_i0_7] = _init7(new BigNumber(_i0_7));
      }
      _209_v14 = _nw43;
      _module.__default.m0(_206_v13, _195_v3, _209_v14, _195_v3, _203_globalState);
      if (!(_192_v0)) {
        let _212_v15;
        _212_v15 = _module.D0.create_DC0(true);
        let _213_v16;
        _213_v16 = _dafny.Seq.of(_module.__default.fm0(_192_v0, _192_v0, _192_v0, (_212_v15).dtor_cf0, _203_globalState));
        _module.__default.m0(_206_v13, (new BigNumber((_213_v16).length)).multipliedBy(_195_v3), _209_v14, new BigNumber((_193_v1).length), _203_globalState);
        _192_v0 = _192_v0;
        let _214_v17;
        let _nw44 = new _module.C2();
        _nw44.__ctor(_195_v3);
        _214_v17 = _nw44;
        _192_v0 = _192_v0;
        _192_v0 = _192_v0;
      } else {
        (_203_globalState).f2 = _195_v3;
        _192_v0 = _192_v0;
        let _215_v18;
        _215_v18 = new _dafny.CodePoint('t'.codePointAt(0));
        let _216_v19;
        _216_v19 = _dafny.Seq.of(_192_v0, _192_v0);
        let _source2 = _module.__default.fm13(_215_v18, _dafny.MultiSet.FromArray(_216_v19), _203_globalState);
        if (_source2.is_DC4) {
          let _217___mcc_h0 = (_source2).cf6;
          let _218___mcc_h1 = (_source2).cf7;
          let _219_cf7 = _218___mcc_h1;
          let _220_cf6 = _217___mcc_h0;
          let _rhs36 = _192_v0;
          let _rhs37 = _module.__default.safeDivisionInt(_195_v3, _219_cf7);
          let _rhs38 = _206_v13;
          let _rhs39 = (_dafny.ZERO).minus(_219_cf7);
          let _lhs28 = _203_globalState;
          _192_v0 = _rhs36;
          _lhs28.f2 = _rhs37;
          _206_v13 = _rhs38;
          _195_v3 = _rhs39;
          (_203_globalState).f2 = _219_cf7;
          let _221_v20;
          let _nw45 = new _module.C1();
          _nw45.__ctor(true);
          _221_v20 = _nw45;
          _221_v20 = _221_v20;
          let _222_v21;
          let _223_v22;
          let _224_v23;
          let _out0;
          let _out1;
          let _out2;
          let _outcollector0 = (_221_v20).m4(_205_v12, _203_globalState);
          _out0 = _outcollector0[0];
          _out1 = _outcollector0[1];
          _out2 = _outcollector0[2];
          _222_v21 = _out0;
          _223_v22 = _out1;
          _224_v23 = _out2;
        } else if (_source2.is_DC5) {
          let _225___mcc_h2 = (_source2).cf8;
          let _226___mcc_h3 = (_source2).cf9;
          let _227_cf9 = _226___mcc_h3;
          let _228_cf8 = _225___mcc_h2;
          let _229_v24;
          _229_v24 = _module.D4.create_DC13(_215_v18);
          let _230_v25;
          _230_v25 = _dafny.Map.Empty.slice().updateUnsafe(_227_cf9,_229_v24);
          let _231_v26;
          _231_v26 = _dafny.Set.fromElements(_195_v3, _195_v3);
          let _232_v27;
          _232_v27 = _module.D3.create_DC10(_192_v0, _195_v3, _227_cf9, _206_v13, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_231_v26).length),_227_cf9));
          _230_v25 = (_230_v25).update((_232_v27).dtor_cf16, _229_v24);
          _204_v11 = (_204_v11).update(_192_v0, _192_v0);
          _227_cf9 = _227_cf9;
          let _233_v28;
          let _nw46 = new _module.C3();
          _nw46.__ctor(_227_cf9, ((_198_v6)[_module.__default.safeIndex(_195_v3, new BigNumber((_198_v6).length))]).multipliedBy(new BigNumber((_216_v19).length)));
          _233_v28 = _nw46;
        } else if (_source2.is_DC3) {
          let _234___mcc_h4 = (_source2).cf5;
          let _235_cf5 = _234___mcc_h4;
          (_203_globalState).f2 = _195_v3;
          _192_v0 = _192_v0;
          (_203_globalState).f2 = _module.__default.safeDivisionInt(_195_v3, _195_v3);
          _195_v3 = _module.__default.fm2(!(_192_v0), _203_globalState);
        } else {
          let _236___mcc_h5 = (_source2).cf10;
          let _237_cf10 = _236___mcc_h5;
          let _238_v29;
          _238_v29 = _module.D1.create_DC3(_215_v18);
          let _239_v30;
          _239_v30 = _dafny.Seq.of(_module.__default.fm5(new BigNumber(454), new BigNumber(47), new BigNumber(921), _216_v19, _203_globalState), _238_v29);
          let _240_v31;
          _240_v31 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("mawdmqm"), _196_v4),new BigNumber((_239_v30).length));
          _240_v31 = (((true) ? (_240_v31) : (_240_v31))).Merge(_240_v31);
          _module.__default.m0(_206_v13, ((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.FromArray(_216_v19)).cardinality()))).multipliedBy(_195_v3), _209_v14, _195_v3, _203_globalState);
          let _241_v32;
          _241_v32 = _dafny.Set.fromElements(new BigNumber((_216_v19).length), _195_v3);
          let _index29 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_200_v8).length));
          (_200_v8)[_index29] = new BigNumber((_241_v32).length);
          let _index30 = _module.__default.safeIndex(new BigNumber(188), new BigNumber((_200_v8).length));
          (_200_v8)[_index30] = _195_v3;
          (_203_globalState).f2 = _195_v3;
        }
        (_203_globalState).f2 = _195_v3;
        _202_v10 = _202_v10;
      }
      let _hi0 = _195_v3;
      for (let _242_i2 = (_195_v3).multipliedBy(_195_v3); _242_i2.isLessThan(_hi0); _242_i2 = _242_i2.plus(_dafny.ONE)) {
        let _243_v33;
        let _nw47 = new _module.C2();
        _nw47.__ctor(_195_v3);
        _243_v33 = _nw47;
        let _244_v34;
        let _nw48 = new _module.C1();
        _nw48.__ctor(_192_v0);
        _244_v34 = _nw48;
        let _245_v35;
        _245_v35 = _dafny.Map.Empty.slice().updateUnsafe(_195_v3,_244_v34);
        let _246_v36;
        let _nw49 = new _module.C0();
        _nw49.__ctor();
        _246_v36 = _nw49;
        let _247_v37;
        _247_v37 = _dafny.Map.Empty.slice().updateUnsafe(_245_v35,_246_v36);
        let _248_v38;
        _248_v38 = _dafny.Map.Empty.slice().updateUnsafe((((_247_v37).contains(_245_v35)) ? ((_247_v37).get(_245_v35)) : (_246_v36)),_243_v33);
        let _249_v39;
        _249_v39 = _dafny.Seq.of(!(_244_v34.f14));
        let _250_v40;
        let _nw50 = Array((new BigNumber(25)).toNumber());
        _nw50[(_dafny.ZERO).toNumber()] = _243_v33;
        _nw50[(_dafny.ONE).toNumber()] = _243_v33;
        _nw50[(new BigNumber(2)).toNumber()] = _243_v33;
        _nw50[(new BigNumber(3)).toNumber()] = _243_v33;
        _nw50[(new BigNumber(4)).toNumber()] = _243_v33;
        _nw50[(new BigNumber(5)).toNumber()] = _243_v33;
        _nw50[(new BigNumber(6)).toNumber()] = _243_v33;
        _nw50[(new BigNumber(7)).toNumber()] = _243_v33;
        _nw50[(new BigNumber(8)).toNumber()] = _243_v33;
        _nw50[(new BigNumber(9)).toNumber()] = _243_v33;
        _nw50[(new BigNumber(10)).toNumber()] = _243_v33;
        _nw50[(new BigNumber(11)).toNumber()] = _243_v33;
        _nw50[(new BigNumber(12)).toNumber()] = _243_v33;
        _nw50[(new BigNumber(13)).toNumber()] = ((_192_v0) ? (_243_v33) : (_243_v33));
        _nw50[(new BigNumber(14)).toNumber()] = _243_v33;
        _nw50[(new BigNumber(15)).toNumber()] = _243_v33;
        _nw50[(new BigNumber(16)).toNumber()] = _243_v33;
        _nw50[(new BigNumber(17)).toNumber()] = _243_v33;
        _nw50[(new BigNumber(18)).toNumber()] = _243_v33;
        _nw50[(new BigNumber(19)).toNumber()] = _243_v33;
        _nw50[(new BigNumber(20)).toNumber()] = (((_248_v38).contains(_246_v36)) ? ((_248_v38).get(_246_v36)) : (_243_v33));
        _nw50[(new BigNumber(21)).toNumber()] = (((_249_v39)[_module.__default.safeIndex(_243_v33.f13, new BigNumber((_249_v39).length))]) ? (_243_v33) : (_243_v33));
        _nw50[(new BigNumber(22)).toNumber()] = _243_v33;
        _nw50[(new BigNumber(23)).toNumber()] = _243_v33;
        _nw50[(new BigNumber(24)).toNumber()] = _243_v33;
        _250_v40 = _nw50;
        let _index31 = _module.__default.safeIndex(new BigNumber(930), new BigNumber((_250_v40).length));
        (_250_v40)[_index31] = _243_v33;
        let _index32 = _module.__default.safeIndex(new BigNumber(930), new BigNumber((_250_v40).length));
        let _rhs40 = !(((_195_v3).multipliedBy(new BigNumber(-807))).isLessThanOrEqualTo(_243_v33.f13));
        let _rhs41 = _243_v33;
        let _lhs29 = _250_v40;
        let _lhs30 = _module.__default.safeIndex(new BigNumber(930), new BigNumber((_250_v40).length));
        _192_v0 = _rhs40;
        _lhs29[_lhs30] = _rhs41;
        _196_v4 = _dafny.Seq.Concat(_196_v4, _196_v4);
        let _251_v42;
        _251_v42 = _dafny.MultiSet.fromElements(_242_i2, _243_v33.f13);
        let _252_v43;
        _252_v43 = _dafny.Set.fromElements(new BigNumber((_251_v42).cardinality()));
        (_244_v34).f14 = !((function () {
          let _coll13 = new _dafny.Map();
          for (const _compr_13 of (_252_v43).Elements) {
            let _253_v41 = _compr_13;
            if ((_252_v43).contains(_253_v41)) {
              _coll13.push([_module.__default.safeModuloInt(_253_v41, _195_v3),_243_v33.f13]);
            }
          }
          return _coll13;
        }()).update((_dafny.ZERO).minus(_243_v33.f13), new BigNumber((_252_v43).length))).contains(_195_v3);
        let _254_v44;
        _254_v44 = _dafny.MultiSet.fromElements(_200_v8, _200_v8, _200_v8, _200_v8);
        if ((_254_v44).contains(_200_v8)) {
          (_203_globalState).f2 = (_dafny.ZERO).minus(((_198_v6)[_module.__default.safeIndex(_243_v33.f13, new BigNumber((_198_v6).length))]).multipliedBy(_242_i2));
          _196_v4 = _dafny.Seq.Concat(_dafny.Seq.Concat(_196_v4, _196_v4), _196_v4);
          let _255_v45;
          _255_v45 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_196_v4, _dafny.Seq.UnicodeFromString("ux")), _196_v4, _dafny.Seq.Concat(_196_v4, _196_v4), _196_v4, _196_v4);
          _255_v45 = _255_v45;
          let _256_v46;
          let _init8 = function (_257_i3) {
            return _dafny.Seq.UnicodeFromString("k");
          };
          let _nw51 = Array((new BigNumber(15)).toNumber());
          for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw51.length); _i0_8++) {
            _nw51[_i0_8] = _init8(new BigNumber(_i0_8));
          }
          _256_v46 = _nw51;
          let _258_v47;
          _258_v47 = new _dafny.CodePoint('d'.codePointAt(0));
          let _index33 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_256_v46).length));
          (_256_v46)[_index33] = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("dfldjy"), _dafny.Seq.UnicodeFromString("e")), _module.__default.safeIndex(_243_v33.f13, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("dfldjy"), _dafny.Seq.UnicodeFromString("e"))).length)), _258_v47);
          let _259_v48;
          _259_v48 = _module.D2.create_DC8(new BigNumber((_249_v39).length));
          let _index34 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_256_v46).length));
          let _rhs42 = (_243_v33.f13).multipliedBy((_259_v48).dtor_cf12);
          let _rhs43 = _dafny.Seq.UnicodeFromString("cqdmq");
          let _rhs44 = _243_v33.f13;
          let _rhs45 = (_243_v33.f13).plus(_243_v33.f13);
          let _lhs31 = _243_v33;
          let _lhs32 = _256_v46;
          let _lhs33 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_256_v46).length));
          let _lhs34 = _203_globalState;
          let _lhs35 = _203_globalState;
          _lhs31.f13 = _rhs42;
          _lhs32[_lhs33] = _rhs43;
          _lhs34.f2 = _rhs44;
          _lhs35.f2 = _rhs45;
          let _260_v49;
          _260_v49 = _dafny.Map.Empty.slice().updateUnsafe(_205_v12,_dafny.Map.Empty.slice().updateUnsafe(_195_v3,_244_v34));
          _260_v49 = _260_v49;
        } else {
          _204_v11 = _204_v11;
          let _261_v50;
          _261_v50 = new _dafny.CodePoint('h'.codePointAt(0));
          let _262_v51;
          _262_v51 = _module.D1.create_DC3(_261_v50);
          let _263_v52;
          let _nw52 = new _module.C3();
          _nw52.__ctor(!((_246_v36).fm7(false, _262_v51, _196_v4, _244_v34.f14, _203_globalState)), _195_v3);
          _263_v52 = _nw52;
          _263_v52 = _263_v52;
          let _264_v53;
          let _nw53 = Array((_dafny.ONE).toNumber()).fill(_module.D4.Default());
          _264_v53 = _nw53;
          let _265_v54;
          _265_v54 = _module.D4.create_DC12(_193_v1);
          let _index35 = _module.__default.safeIndex(new BigNumber(230), new BigNumber((_264_v53).length));
          (_264_v53)[_index35] = _265_v54;
          let _index36 = _module.__default.safeIndex(new BigNumber(230), new BigNumber((_264_v53).length));
          (_264_v53)[_index36] = _265_v54;
          let _266_v55;
          _266_v55 = _module.D1.create_DC5(_192_v0, _244_v34.f14);
          let _pat_let_tv4 = _244_v34;
          _192_v0 = (function (_pat_let4_0) {
            return function (_267_dt__update__tmp_h0) {
              return function (_pat_let5_0) {
                return function (_268_dt__update_hcf9_h0) {
                  return _module.D1.create_DC5((_267_dt__update__tmp_h0).dtor_cf8, _268_dt__update_hcf9_h0);
                }(_pat_let5_0);
              }(_pat_let_tv4.f14);
            }(_pat_let4_0);
          }(_266_v55)).dtor_cf9;
          (_244_v34).f14 = (_244_v34.f14) === (_244_v34.f14);
        }
      }
      let _269_v56;
      _269_v56 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("sharml"),_195_v3);
      let _270_v57;
      _270_v57 = new _dafny.CodePoint('c'.codePointAt(0));
      let _271_v58;
      let _nw54 = new _module.C3();
      _nw54.__ctor(_192_v0, _195_v3);
      _271_v58 = _nw54;
      let _272_v59;
      _272_v59 = _dafny.Map.Empty.slice().updateUnsafe(_270_v57,_271_v58);
      let _273_v60;
      _273_v60 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm4(_203_globalState),(((_272_v59).contains(_270_v57)) ? ((_272_v59).get(_270_v57)) : (_271_v58)));
      let _274_v61;
      _274_v61 = _dafny.Map.Empty.slice().updateUnsafe(_273_v60,_200_v8);
      let _275_v62;
      _275_v62 = _dafny.Seq.of((new BigNumber((_274_v61).length)).isLessThan(new BigNumber((_269_v56).length)));
      _275_v62 = _275_v62;
      let _276_v63;
      let _277_v64;
      let _278_v65;
      let _out3;
      let _out4;
      let _out5;
      let _outcollector1 = (_271_v58).m2(_192_v0, (_271_v58).f11, _203_globalState);
      _out3 = _outcollector1[0];
      _out4 = _outcollector1[1];
      _out5 = _outcollector1[2];
      _276_v63 = _out3;
      _277_v64 = _out4;
      _278_v65 = _out5;
      let _279_v66;
      _279_v66 = _dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1((_271_v58).f12, (_271_v58).f12, (_271_v58).f12),(_271_v58).f11);
      let _280_v67;
      _280_v67 = _dafny.Map.Empty.slice().updateUnsafe((_271_v58).f12,_195_v3);
      let _281_v68;
      _281_v68 = _dafny.Map.Empty.slice().updateUnsafe(_192_v0,_280_v67);
      let _282_v69;
      _282_v69 = _201_v9;
      let _283_v70;
      _283_v70 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_282_v69,(_271_v58).f12));
      let _index37 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length));
      (_200_v8)[_index37] = new BigNumber((_283_v70).cardinality());
      let _index38 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length));
      let _rhs46 = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_271_v58).f12, new BigNumber(((_204_v11).Merge(_204_v11)).length)));
      let _rhs47 = _279_v66;
      let _rhs48 = _module.__default.fm14(_dafny.Seq.Concat(_196_v4, _196_v4), _203_globalState);
      let _rhs49 = (new BigNumber((_196_v4).length)).plus((_271_v58).f12);
      let _lhs36 = _200_v8;
      let _lhs37 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length));
      _195_v3 = _rhs46;
      _279_v66 = _rhs47;
      _281_v68 = _rhs48;
      _lhs36[_lhs37] = _rhs49;
      let _284_i4;
      _284_i4 = _dafny.ZERO;
      L0: {
        while ((new BigNumber((_202_v10).cardinality())).isLessThanOrEqualTo((_271_v58).f12)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_284_i4)) {
              break L0;
            }
            _284_i4 = (_284_i4).plus(_dafny.ONE);
            let _285_v71;
            _285_v71 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_200_v8)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length))], (_271_v58).f12),_196_v4);
            let _286_v72;
            _286_v72 = _dafny.Seq.of(_271_v58);
            let _287_v73;
            _287_v73 = _286_v72;
            _285_v71 = (_285_v71).update(new BigNumber(((_287_v73)).length), _dafny.Seq.of(_module.__default.fm6((_271_v58).f12, _192_v0, _278_v65, new _dafny.CodePoint('n'.codePointAt(0)), _203_globalState)));
            _192_v0 = ((_202_v10).Intersect(_dafny.MultiSet.fromElements(true))).equals(_202_v10);
            let _288_v74;
            _288_v74 = _module.D4.create_DC13((((_271_v58).f11) ? (_270_v57) : (_270_v57)));
            _288_v74 = _288_v74;
            let _289_v75;
            _289_v75 = _module.D1.create_DC5((_271_v58).f11, _192_v0);
            let _290_v76;
            _290_v76 = _dafny.Seq.of(_289_v75);
            _290_v76 = _290_v76;
          }
        }
      }
      let _291_v77;
      let _nw55 = new _module.C3();
      _nw55.__ctor(!((_271_v58).f11), new BigNumber(651));
      _291_v77 = _nw55;
      let _292_v78;
      _292_v78 = _dafny.Map.Empty.slice().updateUnsafe(_291_v77,(_291_v77).f12);
      _292_v78 = (_292_v78).update(_291_v77, (_291_v77).f12);
      _195_v3 = (_271_v58).f12;
      (_291_v77).m1(!(_276_v63), _203_globalState);
      let _293_v79;
      let _nw56 = Array((new BigNumber(20)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _293_v79 = _nw56;
      let _294_v80;
      _294_v80 = _dafny.MultiSet.fromElements(_293_v79, _293_v79, _293_v79);
      _294_v80 = _294_v80;
      let _source3 = _module.__default.fm13(_270_v57, _dafny.MultiSet.fromElements((_271_v58).f11, _278_v65, !((_291_v77).f11)), _203_globalState);
      if (_source3.is_DC4) {
        let _295___mcc_h6 = (_source3).cf6;
        let _296___mcc_h7 = (_source3).cf7;
        let _297_cf7 = _296___mcc_h7;
        let _298_cf6 = _295___mcc_h6;
        let _299_v81;
        _299_v81 = _module.D1.create_DC3(_270_v57);
        let _source4 = _299_v81;
        if (_source4.is_DC4) {
          let _300___mcc_h12 = (_source4).cf6;
          let _301___mcc_h13 = (_source4).cf7;
          let _302_cf7 = _301___mcc_h13;
          let _303_cf6 = _300___mcc_h12;
          let _304_v82;
          _304_v82 = _dafny.Set.fromElements(new BigNumber(832), (_291_v77).f12, (_291_v77).f12);
          let _305_v83;
          let _nw57 = new _module.C3();
          _nw57.__ctor(!(_192_v0), new BigNumber((_304_v82).length));
          _305_v83 = _nw57;
          let _306_v84;
          _306_v84 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("nuwybql")).length),((_305_v83).f11) || (_module.__default.fm0(false, (_271_v58).f11, _192_v0, (_291_v77).f11, _203_globalState)));
          let _307_v85;
          _307_v85 = _dafny.MultiSet.fromElements(_204_v11);
          _306_v84 = (_306_v84).update((_dafny.ZERO).minus(new BigNumber(((_307_v85).Difference(_dafny.MultiSet.fromElements(_204_v11, _204_v11, _204_v11))).cardinality())), (_291_v77).f11);
          let _index39 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length));
          (_200_v8)[_index39] = (_271_v58).f12;
          _276_v63 = _278_v65;
        } else if (_source4.is_DC5) {
          let _308___mcc_h14 = (_source4).cf8;
          let _309___mcc_h15 = (_source4).cf9;
          let _310_cf9 = _309___mcc_h15;
          let _311_cf8 = _308___mcc_h14;
          let _312_v86;
          _312_v86 = _dafny.MultiSet.fromElements((_271_v58).fm1(_196_v4, (_271_v58).f12, _203_globalState), (_291_v77).f12, (_291_v77).f12, new BigNumber((_196_v4).length), new BigNumber((_197_v5).length));
          _module.__default.m0(_206_v13, ((((_312_v86).contains(new BigNumber(-408))) ? ((_312_v86).get(new BigNumber(-408))) : ((_291_v77).f12))).plus((_291_v77).f12), _209_v14, _195_v3, _203_globalState);
          _196_v4 = _dafny.Seq.Concat(_196_v4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(889)), function (_313_i5) {
            return new _dafny.CodePoint('j'.codePointAt(0));
          }));
          let _314_v87;
          let _315_v88;
          let _316_v89;
          let _out6;
          let _out7;
          let _out8;
          let _outcollector2 = (_271_v58).m2(!(_195_v3).isEqualTo(_297_cf7), _276_v63, _203_globalState);
          _out6 = _outcollector2[0];
          _out7 = _outcollector2[1];
          _out8 = _outcollector2[2];
          _314_v87 = _out6;
          _315_v88 = _out7;
          _316_v89 = _out8;
          let _index40 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length));
          (_200_v8)[_index40] = (_dafny.ZERO).minus((_291_v77).f12);
        } else if (_source4.is_DC3) {
          let _317___mcc_h16 = (_source4).cf5;
          let _318_cf5 = _317___mcc_h16;
          _276_v63 = !(!(!(((false) ? ((_271_v58).f11) : ((_271_v58).f11))) || (!((_271_v58).f11))));
          _297_cf7 = (_297_cf7).minus((new BigNumber(10)).minus(new BigNumber((function () {
            let _coll14 = new _dafny.Set();
            for (const _compr_14 of _dafny.IntegerRange(new BigNumber(393), new BigNumber(611))) {
              let _319_v90 = _compr_14;
              if (((new BigNumber(393)).isLessThanOrEqualTo(_319_v90)) && ((_319_v90).isLessThan(new BigNumber(611)))) {
                _coll14.add(_module.__default.safeDivisionInt(_319_v90, (_271_v58).f12));
              }
            }
            return _coll14;
          }()).length)));
          _204_v11 = (_204_v11).update((_271_v58).f11, _276_v63);
          let _320_v91;
          let _nw58 = new _module.C0();
          _nw58.__ctor();
          _320_v91 = _nw58;
        } else {
          let _321___mcc_h17 = (_source4).cf10;
          let _322_cf10 = _321___mcc_h17;
          let _323_v92;
          let _nw59 = new _module.C1();
          _nw59.__ctor((_291_v77).f11);
          _323_v92 = _nw59;
          let _324_v93;
          _324_v93 = _dafny.Set.fromElements(_323_v92, _323_v92, _323_v92);
          (_271_v58).m1((_324_v93).IsProperSubsetOf(_324_v93), _203_globalState);
          let _325_v94;
          let _nw60 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _325_v94 = _nw60;
          let _index41 = _module.__default.safeIndex(new BigNumber(269), new BigNumber((_325_v94).length));
          (_325_v94)[_index41] = _196_v4;
          let _index42 = _module.__default.safeIndex(new BigNumber(269), new BigNumber((_325_v94).length));
          (_325_v94)[_index42] = _dafny.Seq.UnicodeFromString("adpwq");
          let _326_v95;
          _326_v95 = _dafny.MultiSet.fromElements(_271_v58);
          let _index43 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length));
          let _rhs50 = _195_v3;
          let _rhs51 = _326_v95;
          let _rhs52 = (_module.__default.safeModuloInt((_291_v77).f12, (_200_v8)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length))])).isEqualTo(((_291_v77).f12).multipliedBy(_195_v3));
          let _lhs38 = _200_v8;
          let _lhs39 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length));
          _lhs38[_lhs39] = _rhs50;
          _326_v95 = _rhs51;
          _276_v63 = _rhs52;
          _195_v3 = (_271_v58).f12;
        }
        _206_v13 = _206_v13;
        if (_192_v0) {
          (_203_globalState).f2 = (((_271_v58).f11) ? ((_dafny.ZERO).minus((_291_v77).f12)) : (new BigNumber(-201)));
          let _327_v96;
          _327_v96 = _module.D7.create_DC17(_192_v0);
          let _328_v97;
          _328_v97 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(!(_276_v63), (_327_v96).dtor_cf25)).length));
          let _rhs53 = (_dafny.Set.fromElements(_297_cf7)).IsSubsetOf(_328_v97);
          let _rhs54 = new BigNumber(((((_291_v77).f11) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-994)), ((_329_v57) => function (_330_i6) {
            return _329_v57;
          })(_270_v57))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-963)), ((_331_v57) => function (_332_i7) {
            return _331_v57;
          })(_270_v57))))).length);
          let _lhs40 = _203_globalState;
          _192_v0 = _rhs53;
          _lhs40.f2 = _rhs54;
          let _333_v98;
          _333_v98 = _module.D0.create_DC1((_271_v58).f12, new BigNumber((_198_v6).length), (_291_v77).f12);
          let _334_v99;
          let _nw61 = new _module.C1();
          _nw61.__ctor(((_333_v98).dtor_cf1).isEqualTo(_297_cf7));
          _334_v99 = _nw61;
          let _335_v100;
          let _init9 = function (_336_i8) {
            return _dafny.Seq.UnicodeFromString("n");
          };
          let _nw62 = Array((new BigNumber(8)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw62.length); _i0_9++) {
            _nw62[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _335_v100 = _nw62;
          let _index44 = _module.__default.safeIndex(new BigNumber(469), new BigNumber((_335_v100).length));
          (_335_v100)[_index44] = _dafny.Seq.Concat(_196_v4, _196_v4);
          let _index45 = _module.__default.safeIndex(new BigNumber(469), new BigNumber((_335_v100).length));
          (_335_v100)[_index45] = _196_v4;
          let _337_v101;
          let _nw63 = new _module.C1();
          _nw63.__ctor((_291_v77).f11);
          _337_v101 = _nw63;
        } else {
          _275_v62 = _dafny.Seq.Concat(_275_v62, _dafny.Seq.Concat(_275_v62, _275_v62));
          _module.__default.m0(_206_v13, (_271_v58).f12, _209_v14, (_291_v77).f12, _203_globalState);
          let _338_v102;
          _338_v102 = _dafny.Map.Empty.slice().updateUnsafe(_292_v78,(_291_v77).fm1(_196_v4, _195_v3, _203_globalState));
          _297_cf7 = (((!((_271_v58).f11)) === (_module.__default.fm0(true, (_271_v58).f11, (_271_v58).f11, (_271_v58).f11, _203_globalState))) ? ((_271_v58).f12) : ((((_338_v102).contains(_292_v78)) ? ((_338_v102).get(_292_v78)) : ((_200_v8)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length))]))));
          (_203_globalState).f2 = (_297_cf7).minus((_291_v77).f12);
          _206_v13 = (((_291_v77).f11) ? (_206_v13) : (_206_v13));
        }
        let _339_v103;
        _339_v103 = _dafny.MultiSet.fromElements(_200_v8, (((_291_v77).f11) ? (_200_v8) : (_200_v8)));
        let _340_v104;
        _340_v104 = _dafny.Seq.of(_291_v77, _271_v58);
        let _341_v105;
        _341_v105 = _340_v104;
        let _342_v106;
        let _nw64 = Array((new BigNumber(9)).toNumber());
        _nw64[(_dafny.ZERO).toNumber()] = _340_v104;
        _nw64[(_dafny.ONE).toNumber()] = _341_v105;
        _nw64[(new BigNumber(2)).toNumber()] = _341_v105;
        _nw64[(new BigNumber(3)).toNumber()] = _340_v104;
        _nw64[(new BigNumber(4)).toNumber()] = _340_v104;
        _nw64[(new BigNumber(5)).toNumber()] = _341_v105;
        _nw64[(new BigNumber(6)).toNumber()] = _341_v105;
        _nw64[(new BigNumber(7)).toNumber()] = _dafny.Seq.update(_340_v104, _module.__default.safeIndex((_291_v77).f12, new BigNumber((_340_v104).length)), _271_v58);
        _nw64[(new BigNumber(8)).toNumber()] = _341_v105;
        _342_v106 = _nw64;
        let _343_v107;
        _343_v107 = _dafny.Map.Empty.slice().updateUnsafe(_342_v106,_dafny.MultiSet.fromElements(_200_v8));
        _339_v103 = (((_343_v107).contains(_342_v106)) ? ((_343_v107).get(_342_v106)) : (_339_v103));
      } else if (_source3.is_DC5) {
        let _344___mcc_h8 = (_source3).cf8;
        let _345___mcc_h9 = (_source3).cf9;
        let _346_cf9 = _345___mcc_h9;
        let _347_cf8 = _344___mcc_h8;
        let _348_v108;
        _348_v108 = _dafny.MultiSet.fromElements(_module.__default.fm9((_291_v77).f11, _203_globalState));
        let _349_v109;
        _349_v109 = _module.D1.create_DC4(_348_v108, (_291_v77).f12);
        _349_v109 = _349_v109;
        let _index46 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_206_v13).length));
        (_206_v13)[_index46] = (_271_v58).f11;
        let _index47 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_206_v13).length));
        (_206_v13)[_index47] = false;
        let _350_v110;
        _350_v110 = _dafny.Map.Empty.slice().updateUnsafe(_347_cf8,_275_v62);
        let _351_v111;
        _351_v111 = _dafny.Map.Empty.slice().updateUnsafe((((_350_v110).contains(_346_cf9)) ? ((_350_v110).get(_346_cf9)) : (_dafny.Seq.update(_275_v62, _module.__default.safeIndex(_195_v3, new BigNumber((_275_v62).length)), _276_v63))),new BigNumber((_dafny.Seq.update(_275_v62, _module.__default.safeIndex(new BigNumber(-519), new BigNumber((_275_v62).length)), !(!(_347_cf8)))).length));
        let _352_v112;
        _352_v112 = _dafny.MultiSet.fromElements(new BigNumber((_351_v111).length));
        _195_v3 = new BigNumber((_dafny.Set.fromElements(((_271_v58).f12).minus(new BigNumber((_352_v112).cardinality())))).length);
        let _353_v113;
        _353_v113 = _module.D1.create_DC5(_346_cf9, _347_cf8);
        if ((_353_v113).dtor_cf9) {
          _196_v4 = _dafny.Seq.Concat(_196_v4, _196_v4);
          let _354_v114;
          let _nw65 = new _module.C3();
          _nw65.__ctor((_291_v77).f11, _module.__default.safeDivisionInt(new BigNumber((_module.__default.fm15((_291_v77).f12, (_291_v77).f12, _192_v0, _203_globalState)).length), _195_v3));
          _354_v114 = _nw65;
          let _355_v115;
          _355_v115 = _dafny.Seq.of(_200_v8, _200_v8);
          let _356_v116;
          _356_v116 = _dafny.Map.Empty.slice().updateUnsafe(_270_v57,_200_v8);
          let _357_v117;
          let _nw66 = Array((new BigNumber(20)).toNumber());
          _nw66[(_dafny.ZERO).toNumber()] = (((_201_v9).contains(new BigNumber(83))) ? ((_201_v9).get(new BigNumber(83))) : ((_355_v115)[_module.__default.safeIndex(new BigNumber(-561), new BigNumber((_355_v115).length))]));
          _nw66[(_dafny.ONE).toNumber()] = _200_v8;
          _nw66[(new BigNumber(2)).toNumber()] = (((_291_v77).f11) ? (_200_v8) : (_200_v8));
          _nw66[(new BigNumber(3)).toNumber()] = (((_356_v116).contains(_270_v57)) ? ((_356_v116).get(_270_v57)) : (_200_v8));
          _nw66[(new BigNumber(4)).toNumber()] = _200_v8;
          _nw66[(new BigNumber(5)).toNumber()] = _200_v8;
          _nw66[(new BigNumber(6)).toNumber()] = _200_v8;
          _nw66[(new BigNumber(7)).toNumber()] = _200_v8;
          _nw66[(new BigNumber(8)).toNumber()] = _200_v8;
          _nw66[(new BigNumber(9)).toNumber()] = _200_v8;
          _nw66[(new BigNumber(10)).toNumber()] = _200_v8;
          _nw66[(new BigNumber(11)).toNumber()] = _200_v8;
          _nw66[(new BigNumber(12)).toNumber()] = _200_v8;
          _nw66[(new BigNumber(13)).toNumber()] = _200_v8;
          _nw66[(new BigNumber(14)).toNumber()] = _200_v8;
          _nw66[(new BigNumber(15)).toNumber()] = _200_v8;
          _nw66[(new BigNumber(16)).toNumber()] = _200_v8;
          _nw66[(new BigNumber(17)).toNumber()] = _200_v8;
          _nw66[(new BigNumber(18)).toNumber()] = _200_v8;
          _nw66[(new BigNumber(19)).toNumber()] = _200_v8;
          _357_v117 = _nw66;
          let _index48 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_357_v117).length));
          (_357_v117)[_index48] = _200_v8;
          let _index49 = _module.__default.safeIndex(new BigNumber(20), new BigNumber((_357_v117).length));
          (_357_v117)[_index49] = _200_v8;
          _354_v114 = _271_v58;
          let _index50 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length));
          (_200_v8)[_index50] = ((_dafny.ZERO).minus((_354_v114).f12)).minus(new BigNumber(500));
        } else {
          let _358_v118;
          let _nw67 = Array((new BigNumber(5)).toNumber()).fill(false);
          _358_v118 = _nw67;
          _module.__default.m0(_358_v118, _module.__default.safeDivisionInt(_195_v3, (_271_v58).f12), _209_v14, (_271_v58).f12, _203_globalState);
          let _359_v119;
          let _360_v120;
          let _361_v121;
          let _out9;
          let _out10;
          let _out11;
          let _outcollector3 = (_271_v58).m2(_347_cf8, ((_271_v58).f12).isLessThanOrEqualTo((_291_v77).f12), _203_globalState);
          _out9 = _outcollector3[0];
          _out10 = _outcollector3[1];
          _out11 = _outcollector3[2];
          _359_v119 = _out9;
          _360_v120 = _out10;
          _361_v121 = _out11;
          let _rhs55 = (_198_v6)[_module.__default.safeIndex(((_192_v0) ? ((_291_v77).f12) : (_195_v3)), new BigNumber((_198_v6).length))];
          let _rhs56 = (_275_v62)[_module.__default.safeIndex(new BigNumber(433), new BigNumber((_275_v62).length))];
          let _rhs57 = _dafny.Seq.contains(_196_v4, _270_v57);
          let _lhs41 = _203_globalState;
          _lhs41.f2 = _rhs55;
          _359_v119 = _rhs56;
          _276_v63 = _rhs57;
          let _index51 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_206_v13).length));
          let _rhs58 = (((_202_v10).contains(_347_cf8)) ? ((_202_v10).get(_347_cf8)) : ((_291_v77).f12));
          let _rhs59 = _dafny.Seq.Concat(_275_v62, _module.__default.fm16(new BigNumber(345), (_353_v113).dtor_cf9, _203_globalState));
          let _rhs60 = true;
          let _rhs61 = ((!(_278_v65)) ? (_200_v8) : (_200_v8));
          let _rhs62 = !(_276_v63);
          let _lhs42 = _203_globalState;
          let _lhs43 = _206_v13;
          let _lhs44 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_206_v13).length));
          let _lhs45 = _203_globalState;
          _lhs42.f2 = _rhs58;
          _275_v62 = _rhs59;
          _lhs43[_lhs44] = _rhs60;
          _lhs45.f10 = _rhs61;
          _346_cf9 = _rhs62;
          let _362_v122;
          let _nw68 = new _module.C2();
          _nw68.__ctor(_195_v3);
          _362_v122 = _nw68;
        }
      } else if (_source3.is_DC3) {
        let _363___mcc_h10 = (_source3).cf5;
        let _364_cf5 = _363___mcc_h10;
        let _365_v123;
        let _nw69 = new _module.C0();
        _nw69.__ctor();
        _365_v123 = _nw69;
        let _366_v124;
        _366_v124 = _module.D7.create_DC16(_365_v123);
        _366_v124 = _366_v124;
        let _367_v125;
        let _nw70 = new _module.C0();
        _nw70.__ctor();
        _367_v125 = _nw70;
        _module.__default.m0(_206_v13, (_291_v77).f12, _209_v14, (_291_v77).f12, _203_globalState);
        let _index52 = _module.__default.safeIndex(new BigNumber(775), new BigNumber((_206_v13).length));
        (_206_v13)[_index52] = _276_v63;
        let _index53 = _module.__default.safeIndex(new BigNumber(775), new BigNumber((_206_v13).length));
        (_206_v13)[_index53] = _276_v63;
      } else {
        let _368___mcc_h11 = (_source3).cf10;
        let _369_cf10 = _368___mcc_h11;
        let _index54 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_206_v13).length));
        (_206_v13)[_index54] = true;
        let _index55 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_206_v13).length));
        (_206_v13)[_index55] = !(_276_v63);
        (_291_v77).m1((_291_v77).f11, _203_globalState);
        let _370_v126;
        let _nw71 = Array((new BigNumber(9)).toNumber()).fill(false);
        _370_v126 = _nw71;
        let _371_v127;
        _371_v127 = _dafny.Map.Empty.slice().updateUnsafe((_291_v77).f12,_192_v0);
        let _372_v128;
        _372_v128 = _module.D3.create_DC10((_206_v13)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_206_v13).length))], (_200_v8)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length))], true, _370_v126, _dafny.Map.Empty.slice().updateUnsafe((_271_v58).f12,(_291_v77).f11));
        let _pat_let_tv5 = _370_v126;
        let _pat_let_tv6 = _271_v58;
        let _pat_let_tv7 = _192_v0;
        let _source5 = function (_pat_let6_0) {
          return function (_373_dt__update__tmp_h4) {
            return function (_pat_let7_0) {
              return function (_374_dt__update_hcf17_h0) {
                return function (_pat_let8_0) {
                  return function (_375_dt__update_hcf14_h0) {
                    return function (_pat_let9_0) {
                      return function (_376_dt__update_hcf16_h0) {
                        return _module.D3.create_DC10(_375_dt__update_hcf14_h0, (_373_dt__update__tmp_h4).dtor_cf15, _376_dt__update_hcf16_h0, _374_dt__update_hcf17_h0, (_373_dt__update__tmp_h4).dtor_cf18);
                      }(_pat_let9_0);
                    }(_pat_let_tv7);
                  }(_pat_let8_0);
                }((_pat_let_tv6).f11);
              }(_pat_let7_0);
            }(_pat_let_tv5);
          }(_pat_let6_0);
        }((((_291_v77).f11) ? (_module.D3.create_DC10(_278_v65, (_200_v8)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length))], (_271_v58).f11, _370_v126, _371_v127)) : (_372_v128)));
        if (_source5.is_DC10) {
          let _377___mcc_h18 = (_source5).cf14;
          let _378___mcc_h19 = (_source5).cf15;
          let _379___mcc_h20 = (_source5).cf16;
          let _380___mcc_h21 = (_source5).cf17;
          let _381___mcc_h22 = (_source5).cf18;
          let _382_cf18 = _381___mcc_h22;
          let _383_cf17 = _380___mcc_h21;
          let _384_cf16 = _379___mcc_h20;
          let _385_cf15 = _378___mcc_h19;
          let _386_cf14 = _377___mcc_h18;
          let _index56 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length));
          (_200_v8)[_index56] = (_271_v58).fm1(_196_v4, new BigNumber((_280_v67).length), _203_globalState);
          _270_v57 = new _dafny.CodePoint('w'.codePointAt(0));
          _195_v3 = ((_dafny.ZERO).minus((_271_v58).f12)).multipliedBy((_271_v58).f12);
          _277_v64 = (_277_v64).update(_275_v62, _module.__default.fm0(_192_v0, (_271_v58).f11, _276_v63, (_271_v58).f11, _203_globalState));
        } else if (_source5.is_DC9) {
          let _387___mcc_h23 = (_source5).cf13;
          let _388_cf13 = _387___mcc_h23;
          let _index57 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_293_v79).length));
          (_293_v79)[_index57] = (((_271_v58).f11) ? (_270_v57) : (_270_v57));
          let _index58 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_293_v79).length));
          (_293_v79)[_index58] = ((!_dafny.Seq.contains(_275_v62, _276_v63)) ? (_270_v57) : (_270_v57));
          let _389_v129;
          _389_v129 = _dafny.MultiSet.fromElements((_271_v58).f12);
          _389_v129 = _389_v129;
          let _390_v130;
          _390_v130 = _module.D8.create_DC20((_291_v77).f11, (_206_v13)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_206_v13).length))], (((_371_v127).contains(new BigNumber((_204_v11).length))) ? ((_371_v127).get(new BigNumber((_204_v11).length))) : ((_291_v77).f11)), (_271_v58).f11);
          let _391_v131;
          _391_v131 = _dafny.Set.fromElements(_196_v4);
          let _392_v132;
          let _nw72 = new _module.C1();
          _nw72.__ctor((_291_v77).f11);
          _392_v132 = _nw72;
          let _393_v133;
          _393_v133 = _dafny.Map.Empty.slice().updateUnsafe(_392_v132,_278_v65);
          _278_v65 = _module.__default.fm0((_390_v130).dtor_cf28, _276_v63, (_391_v131).IsDisjointFrom(_391_v131), (((_393_v133).contains(_392_v132)) ? ((_393_v133).get(_392_v132)) : (false)), _203_globalState);
          let _394_v134;
          let _init10 = ((_395_v6) => function (_396_i9) {
            return _395_v6;
          })(_198_v6);
          let _nw73 = Array((new BigNumber(17)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw73.length); _i0_10++) {
            _nw73[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _394_v134 = _nw73;
          let _397_v135;
          _397_v135 = _dafny.Map.Empty.slice().updateUnsafe(_394_v134,true);
          let _index59 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_206_v13).length));
          (_206_v13)[_index59] = (((_397_v135).contains(_394_v134)) ? ((_397_v135).get(_394_v134)) : ((_291_v77).f11));
        } else {
          let _398___mcc_h24 = (_source5).cf19;
          let _399_cf19 = _398___mcc_h24;
          let _index60 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_206_v13).length));
          let _rhs63 = _module.__default.safeModuloInt(_module.__default.fm2(_278_v65, _203_globalState), (_291_v77).f12);
          let _rhs64 = !(_192_v0) || (true);
          let _lhs46 = _203_globalState;
          let _lhs47 = _206_v13;
          let _lhs48 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_206_v13).length));
          _lhs46.f2 = _rhs63;
          _lhs47[_lhs48] = _rhs64;
          let _400_v136;
          _400_v136 = _dafny.MultiSet.fromElements(_195_v3, (_291_v77).f12, (_291_v77).f12);
          (_203_globalState).f2 = ((((_271_v58).f12).isLessThan(new BigNumber((_400_v136).cardinality()))) ? ((((_202_v10).contains((_291_v77).f11)) ? ((_202_v10).get((_291_v77).f11)) : ((_291_v77).f12))) : ((_dafny.ZERO).minus((_271_v58).f12)));
          _278_v65 = (_372_v128).dtor_cf16;
          _270_v57 = (((_dafny.Set.fromElements(!((_291_v77).f11), _278_v65, (_291_v77).f11, (_271_v58).f11)).IsSubsetOf(_dafny.Set.fromElements((_206_v13)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_206_v13).length))], !(_276_v63), _276_v63))) ? (new _dafny.CodePoint('c'.codePointAt(0))) : (_270_v57));
        }
        _280_v67 = (_280_v67).update(_195_v3, (_dafny.ZERO).minus(new BigNumber((_205_v12).length)));
      }
      let _source6 = _module.D0.create_DC0(_192_v0);
      if (_source6.is_DC1) {
        let _401___mcc_h25 = (_source6).cf1;
        let _402___mcc_h26 = (_source6).cf2;
        let _403___mcc_h27 = (_source6).cf3;
        let _404_cf3 = _403___mcc_h27;
        let _405_cf2 = _402___mcc_h26;
        let _406_cf1 = _401___mcc_h25;
        let _source7 = _module.__default.fm17((_291_v77).f11, _dafny.Seq.Concat(_198_v6, _198_v6), (_275_v62)[_module.__default.safeIndex(_405_cf2, new BigNumber((_275_v62).length))], _192_v0, _203_globalState);
        if (_source7.is_DC4) {
          let _407___mcc_h30 = (_source7).cf6;
          let _408___mcc_h31 = (_source7).cf7;
          let _409_cf7 = _408___mcc_h31;
          let _410_cf6 = _407___mcc_h30;
          let _411_v137;
          _411_v137 = _module.D11.create_DC23(_198_v6);
          _198_v6 = _dafny.Seq.Concat((_411_v137).dtor_cf34, _dafny.Seq.update((((_271_v58).f11) ? (_198_v6) : (_198_v6)), _module.__default.safeIndex(new BigNumber((_196_v4).length), new BigNumber(((((_271_v58).f11) ? (_198_v6) : (_198_v6))).length)), new BigNumber((_197_v5).length)));
          let _412_v138;
          let _413_v139;
          let _414_v140;
          let _out12;
          let _out13;
          let _out14;
          let _outcollector4 = (_291_v77).m2((_271_v58).f11, ((_dafny.ZERO).minus(_406_cf1)).isLessThanOrEqualTo((_291_v77).f12), _203_globalState);
          _out12 = _outcollector4[0];
          _out13 = _outcollector4[1];
          _out14 = _outcollector4[2];
          _412_v138 = _out12;
          _413_v139 = _out13;
          _414_v140 = _out14;
          let _415_v141;
          let _416_v142;
          let _417_v143;
          let _out15;
          let _out16;
          let _out17;
          let _outcollector5 = (_271_v58).m2(((_291_v77).f11) || (_276_v63), !((_271_v58).f12).isEqualTo((_291_v77).f12), _203_globalState);
          _out15 = _outcollector5[0];
          _out16 = _outcollector5[1];
          _out17 = _outcollector5[2];
          _415_v141 = _out15;
          _416_v142 = _out16;
          _417_v143 = _out17;
          let _418_v144;
          _418_v144 = _dafny.Seq.of(((_278_v65) ? (_dafny.Seq.of((_271_v58).f11, (_291_v77).f11)) : (_275_v62)), _275_v62, _module.__default.fm16((_291_v77).f12, (_291_v77).f11, _203_globalState));
          _418_v144 = _418_v144;
        } else if (_source7.is_DC5) {
          let _419___mcc_h32 = (_source7).cf8;
          let _420___mcc_h33 = (_source7).cf9;
          let _421_cf9 = _420___mcc_h33;
          let _422_cf8 = _419___mcc_h32;
          let _423_v145;
          _423_v145 = _dafny.Map.Empty.slice().updateUnsafe(_404_cf3,((_271_v58).f12).isLessThanOrEqualTo((_200_v8)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length))]));
          _423_v145 = _423_v145;
          (_203_globalState).f2 = (_291_v77).f12;
          let _424_v146;
          _424_v146 = _module.D0.create_DC1(new BigNumber(287), (_200_v8)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length))], (_271_v58).f12);
          _195_v3 = (_198_v6)[_module.__default.safeIndex(_module.__default.safeDivisionInt((_424_v146).dtor_cf2, (_dafny.ZERO).minus((((_202_v10).contains((_271_v58).f11)) ? ((_202_v10).get((_271_v58).f11)) : ((_271_v58).f12)))), new BigNumber((_198_v6).length))];
          _196_v4 = _196_v4;
        } else if (_source7.is_DC3) {
          let _425___mcc_h34 = (_source7).cf5;
          let _426_cf5 = _425___mcc_h34;
          _275_v62 = _275_v62;
          let _427_v147;
          _427_v147 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_196_v4).length),_206_v13);
          _206_v13 = (((_427_v147).contains((_200_v8)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length))])) ? ((_427_v147).get((_200_v8)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length))])) : (_206_v13));
          _196_v4 = _dafny.Seq.Concat(_196_v4, _196_v4);
          _276_v63 = (_module.__default.safeDivisionInt((_271_v58).f12, (_dafny.ZERO).minus((_271_v58).f12))).isLessThanOrEqualTo((_291_v77).f12);
        } else {
          let _428___mcc_h35 = (_source7).cf10;
          let _429_cf10 = _428___mcc_h35;
          let _430_v148;
          let _nw74 = Array((new BigNumber(15)).toNumber());
          _nw74[(_dafny.ZERO).toNumber()] = _206_v13;
          _nw74[(_dafny.ONE).toNumber()] = _206_v13;
          _nw74[(new BigNumber(2)).toNumber()] = _206_v13;
          _nw74[(new BigNumber(3)).toNumber()] = _206_v13;
          _nw74[(new BigNumber(4)).toNumber()] = _206_v13;
          _nw74[(new BigNumber(5)).toNumber()] = _206_v13;
          _nw74[(new BigNumber(6)).toNumber()] = _206_v13;
          _nw74[(new BigNumber(7)).toNumber()] = _206_v13;
          _nw74[(new BigNumber(8)).toNumber()] = _206_v13;
          _nw74[(new BigNumber(9)).toNumber()] = _206_v13;
          _nw74[(new BigNumber(10)).toNumber()] = _206_v13;
          _nw74[(new BigNumber(11)).toNumber()] = _206_v13;
          _nw74[(new BigNumber(12)).toNumber()] = _206_v13;
          _nw74[(new BigNumber(13)).toNumber()] = _206_v13;
          _nw74[(new BigNumber(14)).toNumber()] = _206_v13;
          _430_v148 = _nw74;
          let _431_v149;
          _431_v149 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(673),_430_v148);
          let _432_v150;
          _432_v150 = _dafny.Map.Empty.slice().updateUnsafe((_271_v58).f11,(((_431_v149).contains((_271_v58).f12)) ? ((_431_v149).get((_271_v58).f12)) : (_430_v148)));
          _432_v150 = (_432_v150).update(true, _430_v148);
          _270_v57 = _270_v57;
          (_203_globalState).f2 = (_dafny.ZERO).minus(_404_cf3);
          let _433_v151;
          let _init11 = ((_434_v58) => function (_435_i10) {
            return (((_434_v58).f11) ? (_module.D11.create_DC25()) : (_module.D11.create_DC25()));
          })(_271_v58);
          let _nw75 = Array((new BigNumber(29)).toNumber());
          for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw75.length); _i0_11++) {
            _nw75[_i0_11] = _init11(new BigNumber(_i0_11));
          }
          _433_v151 = _nw75;
          let _436_v152;
          _436_v152 = _module.D11.create_DC25();
          let _index61 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_433_v151).length));
          (_433_v151)[_index61] = _436_v152;
          let _437_v153;
          _437_v153 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_404_cf3, (_271_v58).fm1(_196_v4, (_200_v8)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length))], _203_globalState)),_dafny.Seq.of((_291_v77).f12));
          let _index62 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length));
          let _index63 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length));
          let _index64 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_433_v151).length));
          let _index65 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length));
          let _rhs65 = new BigNumber(((((_437_v153).contains(_406_cf1)) ? ((_437_v153).get(_406_cf1)) : (_198_v6))).length);
          let _rhs66 = false;
          let _rhs67 = ((_dafny.Seq.contains(_196_v4, _270_v57)) ? (_406_cf1) : ((_291_v77).f12));
          let _rhs68 = _436_v152;
          let _rhs69 = _module.__default.safeModuloInt(_195_v3, new BigNumber((_196_v4).length));
          let _lhs49 = _200_v8;
          let _lhs50 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length));
          let _lhs51 = _200_v8;
          let _lhs52 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length));
          let _lhs53 = _433_v151;
          let _lhs54 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_433_v151).length));
          let _lhs55 = _200_v8;
          let _lhs56 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length));
          _lhs49[_lhs50] = _rhs65;
          _278_v65 = _rhs66;
          _lhs51[_lhs52] = _rhs67;
          _lhs53[_lhs54] = _rhs68;
          _lhs55[_lhs56] = _rhs69;
        }
        let _438_v154;
        let _nw76 = new _module.C3();
        _nw76.__ctor(_276_v63, _module.__default.safeModuloInt(new BigNumber((_280_v67).length), _404_cf3));
        _438_v154 = _nw76;
        let _source8 = _module.__default.fm17((_271_v58).f11, _198_v6, !((new BigNumber((_dafny.Seq.of((_271_v58).f11)).length)).isLessThanOrEqualTo((_291_v77).f12)), _dafny.Seq.IsPrefixOf(_196_v4, _196_v4), _203_globalState);
        if (_source8.is_DC4) {
          let _439___mcc_h36 = (_source8).cf6;
          let _440___mcc_h37 = (_source8).cf7;
          let _441_cf7 = _440___mcc_h37;
          let _442_cf6 = _439___mcc_h36;
          _278_v65 = _dafny.Seq.IsProperPrefixOf(_196_v4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(379)), ((_443_v57) => function (_444_i11) {
            return _443_v57;
          })(_270_v57)));
          _278_v65 = _module.__default.fm0(!(_205_v12).equals(_205_v12), (_291_v77).f11, _192_v0, !((_438_v154).f11), _203_globalState);
          _276_v63 = (((_271_v58).f11) ? ((_291_v77).f11) : ((_438_v154).f11));
          let _445_v155;
          let _nw77 = new _module.C1();
          _nw77.__ctor((!((_291_v77).f11)) === ((_271_v58).f11));
          _445_v155 = _nw77;
        } else if (_source8.is_DC5) {
          let _446___mcc_h38 = (_source8).cf8;
          let _447___mcc_h39 = (_source8).cf9;
          let _448_cf9 = _447___mcc_h39;
          let _449_cf8 = _446___mcc_h38;
          _196_v4 = _196_v4;
          (_291_v77).m1(((_438_v154).f12).isLessThanOrEqualTo((_271_v58).f12), _203_globalState);
          let _450_v156;
          let _nw78 = new _module.C1();
          _nw78.__ctor(_dafny.Seq.IsProperPrefixOf(_275_v62, _275_v62));
          _450_v156 = _nw78;
          let _451_v157;
          let _nw79 = new _module.C3();
          _nw79.__ctor((_291_v77).f11, _module.__default.safeModuloInt(_404_cf3, new BigNumber((_module.__default.fm18(_405_cf2, _203_globalState)).length)));
          _451_v157 = _nw79;
        } else if (_source8.is_DC3) {
          let _452___mcc_h40 = (_source8).cf5;
          let _453_cf5 = _452___mcc_h40;
          let _454_v158;
          _454_v158 = _module.D1.create_DC3(_453_cf5);
          _453_cf5 = (_454_v158).dtor_cf5;
          let _455_v159;
          _455_v159 = _dafny.Seq.of((_202_v10).update((_271_v58).f11, _module.__default.abs((_291_v77).f12)));
          _270_v57 = _module.__default.fm6((_438_v154).f12, !_dafny.areEqual(_455_v159, _dafny.Seq.Create(_module.__default.abs(new BigNumber(712)), ((_456_v77) => function (_457_i12) {
            return _dafny.MultiSet.fromElements((_456_v77).f11);
          })(_291_v77))), (_291_v77).f11, _453_cf5, _203_globalState);
          (_203_globalState).f2 = new BigNumber((_198_v6).length);
          let _index66 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_206_v13).length));
          (_206_v13)[_index66] = true;
          let _index67 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_206_v13).length));
          (_206_v13)[_index67] = _276_v63;
        } else {
          let _458___mcc_h41 = (_source8).cf10;
          let _459_cf10 = _458___mcc_h41;
          let _460_v160;
          _460_v160 = _dafny.Map.Empty.slice().updateUnsafe(_200_v8,(_271_v58).f11);
          _460_v160 = (_460_v160).update(_200_v8, !_dafny.Seq.contains(_196_v4, _module.__default.fm6(_195_v3, (_291_v77).f11, _192_v0, _270_v57, _203_globalState)));
          _404_cf3 = (_module.__default.safeDivisionInt((_dafny.ZERO).minus(_405_cf2), (_438_v154).f12)).plus((_438_v154).fm1(_dafny.Seq.UnicodeFromString("sn"), _406_cf1, _203_globalState));
          let _461_v161;
          let _nw80 = new _module.C0();
          _nw80.__ctor();
          _461_v161 = _nw80;
          let _462_v162;
          _462_v162 = _module.D7.create_DC16(_461_v161);
          let _463_v163;
          _463_v163 = _dafny.Map.Empty.slice().updateUnsafe((_462_v162).dtor_cf24,new BigNumber((_198_v6).length));
          let _464_v164;
          _464_v164 = _dafny.Map.Empty.slice().updateUnsafe(_463_v163,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(376)), function (_465_i13) {
            return _module.D11.create_DC25();
          })).length)));
          _464_v164 = _464_v164;
          (_203_globalState).f2 = _195_v3;
        }
        let _466_v165;
        _466_v165 = _module.D3.create_DC9(((true) ? (_206_v13) : (_206_v13)));
        let _source9 = _466_v165;
        if (_source9.is_DC10) {
          let _467___mcc_h42 = (_source9).cf14;
          let _468___mcc_h43 = (_source9).cf15;
          let _469___mcc_h44 = (_source9).cf16;
          let _470___mcc_h45 = (_source9).cf17;
          let _471___mcc_h46 = (_source9).cf18;
          let _472_cf18 = _471___mcc_h46;
          let _473_cf17 = _470___mcc_h45;
          let _474_cf16 = _469___mcc_h44;
          let _475_cf15 = _468___mcc_h43;
          let _476_cf14 = _467___mcc_h42;
          let _rhs70 = (_271_v58).f11;
          let _rhs71 = (_196_v4)[_module.__default.safeIndex(_475_cf15, new BigNumber((_196_v4).length))];
          let _rhs72 = !(!(false));
          _474_cf16 = _rhs70;
          _270_v57 = _rhs71;
          _276_v63 = _rhs72;
          let _477_v166;
          _477_v166 = _dafny.Seq.of(_dafny.Seq.Concat(_275_v62, _275_v62));
          _477_v166 = _477_v166;
          _475_cf15 = new BigNumber(((_280_v67).Merge((_280_v67).Merge(_280_v67))).length);
          _204_v11 = _204_v11;
        } else if (_source9.is_DC9) {
          let _478___mcc_h47 = (_source9).cf13;
          let _479_cf13 = _478___mcc_h47;
          _195_v3 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_196_v4, _dafny.Seq.UnicodeFromString("qgqwbmtnb")), _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(58)), function (_480_i14) {
            return new _dafny.CodePoint('r'.codePointAt(0));
          }), _module.__default.safeIndex(new BigNumber(156), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(58)), function (_481_i14) {
            return new _dafny.CodePoint('r'.codePointAt(0));
          })).length)), _module.__default.fm6(_195_v3, (_291_v77).f11, _278_v65, new _dafny.CodePoint('b'.codePointAt(0)), _203_globalState)), _196_v4))).length);
          let _index68 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length));
          (_200_v8)[_index68] = (_271_v58).f12;
          let _482_v167;
          _482_v167 = _module.D0.create_DC1((_271_v58).f12, (_dafny.ZERO).minus((_271_v58).fm1(_dafny.Seq.Create(_module.__default.abs(new BigNumber(637)), ((_483_v57) => function (_484_i15) {
  return _483_v57;
})(_270_v57)), (_271_v58).f12, _203_globalState)), (_438_v154).f12);
          _195_v3 = (_482_v167).dtor_cf3;
          let _index69 = _module.__default.safeIndex(new BigNumber(714), new BigNumber((_479_cf13).length));
          (_479_cf13)[_index69] = !(true) || (_278_v65);
          let _485_v168;
          _485_v168 = _module.D11.create_DC26((_291_v77).f12, _200_v8, (_438_v154).f11);
          let _486_v169;
          _486_v169 = _module.D0.create_DC0((_485_v168).dtor_cf38);
          let _index70 = _module.__default.safeIndex(new BigNumber(714), new BigNumber((_479_cf13).length));
          (_479_cf13)[_index70] = (_486_v169).dtor_cf0;
        } else {
          let _487___mcc_h48 = (_source9).cf19;
          let _488_cf19 = _487___mcc_h48;
          _405_cf2 = (new BigNumber((_193_v1).length)).minus(_405_cf2);
          let _489_v170;
          let _nw81 = Array((new BigNumber(9)).toNumber());
          _489_v170 = _nw81;
          let _490_v171;
          let _nw82 = new _module.C1();
          _nw82.__ctor((_291_v77).f11);
          _490_v171 = _nw82;
          let _index71 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_489_v170).length));
          (_489_v170)[_index71] = _490_v171;
          let _491_v172;
          _491_v172 = _490_v171;
          let _index72 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_489_v170).length));
          (_489_v170)[_index72] = (_491_v172);
          let _index73 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length));
          (_200_v8)[_index73] = ((_200_v8)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length))]).multipliedBy((_271_v58).f12);
          let _index74 = _module.__default.safeIndex(new BigNumber(32), new BigNumber((_206_v13).length));
          (_206_v13)[_index74] = _278_v65;
          let _index75 = _module.__default.safeIndex(new BigNumber(32), new BigNumber((_206_v13).length));
          (_206_v13)[_index75] = (_438_v154).f11;
        }
      } else if (_source6.is_DC0) {
        let _492___mcc_h28 = (_source6).cf0;
        let _493_cf0 = _492___mcc_h28;
        let _494_v173;
        _494_v173 = _dafny.Map.Empty.slice().updateUnsafe(_278_v65,_270_v57);
        let _495_v174;
        let _496_v175;
        let _497_v176;
        let _out18;
        let _out19;
        let _out20;
        let _outcollector6 = (_291_v77).m2((new BigNumber((_494_v173).length)).isEqualTo((_200_v8)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length))]), (_271_v58).f11, _203_globalState);
        _out18 = _outcollector6[0];
        _out19 = _outcollector6[1];
        _out20 = _outcollector6[2];
        _495_v174 = _out18;
        _496_v175 = _out19;
        _497_v176 = _out20;
        (_271_v58).m1(_278_v65, _203_globalState);
        _495_v174 = !_dafny.areEqual(_198_v6, _dafny.Seq.Concat(_198_v6, _198_v6));
        let _498_v177;
        _498_v177 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_204_v11).length),_495_v174);
        let _499_v178;
        _499_v178 = _module.D3.create_DC10(_495_v174, (_271_v58).f12, _276_v63, _206_v13, _498_v177);
        _204_v11 = (_204_v11).update(((_200_v8)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length))]).isLessThanOrEqualTo((_499_v178).dtor_cf15), _276_v63);
      } else {
        let _500___mcc_h29 = (_source6).cf4;
        let _501_cf4 = _500___mcc_h29;
        let _502_v179;
        _502_v179 = _dafny.Seq.of(_196_v4, _196_v4, _196_v4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(570)), ((_503_v57) => function (_504_i16) {
          return _503_v57;
        })(_270_v57)));
        if (_dafny.Seq.IsProperPrefixOf(_502_v179, _dafny.Seq.Create(_module.__default.abs(new BigNumber(24)), ((_505_v4) => function (_506_i17) {
          return _505_v4;
        })(_196_v4)))) {
          let _507_v180;
          _507_v180 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_271_v58).f12),(_271_v58).f11);
          _507_v180 = (_507_v180).update(_195_v3, true);
          let _index76 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_206_v13).length));
          (_206_v13)[_index76] = (_291_v77).f11;
          let _index77 = _module.__default.safeIndex(new BigNumber(578), new BigNumber((_206_v13).length));
          (_206_v13)[_index77] = (_module.D3.create_DC10((_291_v77).f11, (_291_v77).f12, _278_v65, _206_v13, _507_v180)).dtor_cf16;
          (_291_v77).m1(((_200_v8)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length))]).isLessThanOrEqualTo((_271_v58).f12), _203_globalState);
          _196_v4 = _196_v4;
          (_291_v77).m1((_291_v77).f11, _203_globalState);
        } else {
          let _508_v181;
          let _nw83 = Array((new BigNumber(17)).toNumber()).fill([]);
          _508_v181 = _nw83;
          let _index78 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_508_v181).length));
          (_508_v181)[_index78] = _206_v13;
          let _index79 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_508_v181).length));
          (_508_v181)[_index79] = _206_v13;
          let _509_v182;
          let _nw84 = new _module.C3();
          _nw84.__ctor((_291_v77).f11, (_291_v77).f12);
          _509_v182 = _nw84;
          let _index80 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length));
          (_200_v8)[_index80] = _module.__default.fm4(_203_globalState);
          let _510_v183;
          _510_v183 = _196_v4;
          (_203_globalState).f2 = new BigNumber(((_510_v183)).length);
          let _511_v184;
          _511_v184 = _module.D11.create_DC24((_509_v182).f12);
          let _index81 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length));
          (_200_v8)[_index81] = (_511_v184).dtor_cf35;
        }
        let _512_v185;
        _512_v185 = _module.D0.create_DC1(_195_v3, (_291_v77).f12, (_271_v58).f12);
        let _rhs73 = ((_291_v77).f12).multipliedBy((_198_v6)[_module.__default.safeIndex(new BigNumber((_275_v62).length), new BigNumber((_198_v6).length))]);
        let _rhs74 = ((_512_v185).dtor_cf2).isLessThan(new BigNumber(694));
        let _rhs75 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_291_v77).f12),(_271_v58).f12);
        let _rhs76 = _module.__default.safeModuloInt(_module.__default.safeModuloInt((_271_v58).f12, new BigNumber(697)), (_291_v77).f12);
        _195_v3 = _rhs73;
        _278_v65 = _rhs74;
        _280_v67 = _rhs75;
        _195_v3 = _rhs76;
        let _513_v186;
        _513_v186 = _dafny.Map.Empty.slice().updateUnsafe(_196_v4,true);
        let _514_v187;
        _514_v187 = _module.D0.create_DC0((((_513_v186).contains(_dafny.Seq.UnicodeFromString("ddnbjifo"))) ? ((_513_v186).get(_dafny.Seq.UnicodeFromString("ddnbjifo"))) : (_278_v65)));
        let _source10 = _514_v187;
        if (_source10.is_DC1) {
          let _515___mcc_h49 = (_source10).cf1;
          let _516___mcc_h50 = (_source10).cf2;
          let _517___mcc_h51 = (_source10).cf3;
          let _518_cf3 = _517___mcc_h51;
          let _519_cf2 = _516___mcc_h50;
          let _520_cf1 = _515___mcc_h49;
          let _521_v188;
          _521_v188 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm19(_203_globalState),_195_v3);
          _519_cf2 = (((_271_v58).f11) ? (_module.__default.safeModuloInt((_291_v77).f12, _519_cf2)) : (new BigNumber((_521_v188).length)));
          _278_v65 = (_291_v77).f11;
          _192_v0 = ((_291_v77).f12).isLessThanOrEqualTo(_518_cf3);
          _192_v0 = (_291_v77).f11;
        } else if (_source10.is_DC0) {
          let _522___mcc_h52 = (_source10).cf0;
          let _523_cf0 = _522___mcc_h52;
          let _524_v189;
          let _nw85 = new _module.C1();
          _nw85.__ctor(false);
          _524_v189 = _nw85;
          _523_cf0 = _523_cf0;
          _513_v186 = (_513_v186).update(_196_v4, true);
          let _525_v190;
          let _526_v191;
          let _527_v192;
          let _out21;
          let _out22;
          let _out23;
          let _outcollector7 = (_291_v77).m2(((_278_v65) ? ((_291_v77).f11) : ((_271_v58).f11)), !(_276_v63), _203_globalState);
          _out21 = _outcollector7[0];
          _out22 = _outcollector7[1];
          _out23 = _outcollector7[2];
          _525_v190 = _out21;
          _526_v191 = _out22;
          _527_v192 = _out23;
        } else {
          let _528___mcc_h53 = (_source10).cf4;
          let _529_cf4 = _528___mcc_h53;
          let _530_v193;
          let _531_v194;
          let _532_v195;
          let _out24;
          let _out25;
          let _out26;
          let _outcollector8 = (_291_v77).m2(!((_271_v58).f11), _278_v65, _203_globalState);
          _out24 = _outcollector8[0];
          _out25 = _outcollector8[1];
          _out26 = _outcollector8[2];
          _530_v193 = _out24;
          _531_v194 = _out25;
          _532_v195 = _out26;
          let _533_v196;
          _533_v196 = _module.D7.create_DC17((_291_v77).f11);
          _278_v65 = _module.__default.fm0((_532_v195) === ((_291_v77).f11), (_533_v196).dtor_cf25, _530_v193, _276_v63, _203_globalState);
          _532_v195 = _192_v0;
          let _534_v197;
          let _init12 = ((_535_v4) => function (_536_i18) {
            return (_535_v4);
          })(_196_v4);
          let _nw86 = Array((new BigNumber(19)).toNumber());
          for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw86.length); _i0_12++) {
            _nw86[_i0_12] = _init12(new BigNumber(_i0_12));
          }
          _534_v197 = _nw86;
          let _index82 = _module.__default.safeIndex(new BigNumber(130), new BigNumber((_534_v197).length));
          (_534_v197)[_index82] = _196_v4;
          let _index83 = _module.__default.safeIndex(new BigNumber(130), new BigNumber((_534_v197).length));
          let _rhs77 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("c"), _196_v4);
          let _rhs78 = _532_v195;
          let _lhs57 = _534_v197;
          let _lhs58 = _module.__default.safeIndex(new BigNumber(130), new BigNumber((_534_v197).length));
          _lhs57[_lhs58] = _rhs77;
          _530_v193 = _rhs78;
        }
        let _537_v198;
        _537_v198 = _dafny.Set.fromElements((_291_v77).f12, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-9)), ((_538_v1) => function (_539_i19) {
          return (_dafny.ZERO).minus(new BigNumber((_538_v1).length));
        })(_193_v1))).length), (_200_v8)[_module.__default.safeIndex(new BigNumber(595), new BigNumber((_200_v8).length))]);
        let _540_v199;
        _540_v199 = _dafny.Map.Empty.slice().updateUnsafe(_537_v198,_192_v0);
        let _541_v200;
        _541_v200 = _dafny.Seq.of((_540_v199).Merge(_540_v199), _540_v199, _540_v199, _dafny.Map.Empty.slice().updateUnsafe(_537_v198,false));
        let _rhs79 = (((_271_v58).f11) ? ((_dafny.ZERO).minus(((_271_v58).f12).minus(_195_v3))) : (_module.__default.safeDivisionInt((_dafny.ZERO).minus((_271_v58).f12), new BigNumber((_196_v4).length))));
        let _rhs80 = (_541_v200)[_module.__default.safeIndex(_195_v3, new BigNumber((_541_v200).length))];
        _195_v3 = _rhs79;
        _540_v199 = _rhs80;
      }
      process.stdout.write(_dafny.toString(_192_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_193_v1).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_194_v2).equals(_dafny.Set.fromElements(_dafny.Set.fromElements(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_195_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_196_v4).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_197_v5).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.UnicodeFromString("bk")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_198_v6, _dafny.Seq.of(_dafny.ONE, _dafny.ONE, new BigNumber(848)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_199_v7, _dafny.Seq.of(_dafny.Set.fromElements(true), _dafny.Set.fromElements(true), _dafny.Set.fromElements(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_200_v8)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_201_v9).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_202_v10).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_203_globalState.f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_203_globalState).f4).equals(_dafny.Set.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_203_globalState).f5).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(848),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_203_globalState).f6, _dafny.Seq.of(_dafny.ONE, _dafny.ONE, new BigNumber(848)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_203_globalState.f7, _dafny.Seq.of(new BigNumber(490)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState).f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState.f10)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState.f10)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState.f10)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState.f10)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState.f10)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState.f10)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState.f10)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState.f10)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState.f10)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState.f10)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState.f10)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState.f10)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState.f10)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState.f10)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState.f10)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState.f10)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState.f10)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState.f10)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState.f10)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState.f10)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState.f10)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState.f10)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState.f10)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState.f10)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState.f10)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState.f10)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState.f10)[new BigNumber(26)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState.f10)[new BigNumber(27)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_203_globalState.f10)[new BigNumber(28)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_204_v11).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_205_v12).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(848)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_206_v13)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_206_v13)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_206_v13)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_206_v13)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_206_v13)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_206_v13)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_206_v13)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_206_v13)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_206_v13)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_206_v13)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_206_v13)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_206_v13)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_206_v13)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_206_v13)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_206_v13)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_209_v14)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_209_v14)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_209_v14)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_209_v14)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_209_v14)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_209_v14)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_209_v14)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_209_v14)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_209_v14)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_209_v14)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_209_v14)[new BigNumber(10)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_209_v14)[new BigNumber(11)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_209_v14)[new BigNumber(12)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_209_v14)[new BigNumber(13)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_269_v56).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("sharml"),new BigNumber(848)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_270_v57));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_271_v58).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_271_v58).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_272_v59).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_273_v60).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_274_v61).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_275_v62, _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_276_v63));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_277_v64).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false, true, true),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_278_v65));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_279_v66).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC1(new BigNumber(848), new BigNumber(848), new BigNumber(848)),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_280_v67).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(848),new BigNumber(848)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_281_v68).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(424),new BigNumber(154))).updateUnsafe(true,_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(2),new BigNumber(114))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_282_v69)).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_283_v70).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_284_i4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_291_v77).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_291_v77).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_292_v78).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_294_v80).cardinality())));
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
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
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
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1) && _dafny.areEqual(this.cf2, other.cf2) && _dafny.areEqual(this.cf3, other.cf3);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf0 === other.cf0;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf4, other.cf4);
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
    static create_DC4(cf6, cf7) {
      let $dt = new D1(0);
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC5(cf8, cf9) {
      let $dt = new D1(1);
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC3(cf5) {
      let $dt = new D1(2);
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC6(cf10) {
      let $dt = new D1(3);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get is_DC6() { return this.$tag === 3; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC3" + "(" + _dafny.toString(this.cf5) + ")";
      } else if (this.$tag === 3) {
        return "D1.DC6" + "(" + _dafny.toString(this.cf10) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf6, other.cf6) && _dafny.areEqual(this.cf7, other.cf7);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf8 === other.cf8 && this.cf9 === other.cf9;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf5, other.cf5);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC4(_dafny.MultiSet.Empty, _dafny.ZERO);
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
    static create_DC8(cf12) {
      let $dt = new D2(0);
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC7(cf11) {
      let $dt = new D2(1);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC7" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf12, other.cf12);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf11 === other.cf11;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC8(_dafny.ZERO);
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
    static create_DC10(cf14, cf15, cf16, cf17, cf18) {
      let $dt = new D3(0);
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC9(cf13) {
      let $dt = new D3(1);
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC11(cf19) {
      let $dt = new D3(2);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get is_DC11() { return this.$tag === 2; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC11" + "(" + _dafny.toString(this.cf19) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf14 === other.cf14 && _dafny.areEqual(this.cf15, other.cf15) && this.cf16 === other.cf16 && this.cf17 === other.cf17 && _dafny.areEqual(this.cf18, other.cf18);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf13 === other.cf13;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf19, other.cf19);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC10(false, _dafny.ZERO, false, [], _dafny.Map.Empty);
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
    static create_DC13(cf21) {
      let $dt = new D4(0);
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC12(cf20) {
      let $dt = new D4(1);
      $dt.cf20 = cf20;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf20() { return this.cf20; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf20) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf20, other.cf20);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC13(new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC14(cf22) {
      let $dt = new D5(0);
      $dt.cf22 = cf22;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get dtor_cf22() { return this.cf22; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf22) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf22, other.cf22);
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
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC15(cf23) {
      let $dt = new D6(0);
      $dt.cf23 = cf23;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get dtor_cf23() { return this.cf23; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC15" + "(" + _dafny.toString(this.cf23) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf23, other.cf23);
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
          return D6.Default();
        }
      };
    }
  }

  $module.D7 = class D7 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC17(cf25) {
      let $dt = new D7(0);
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC16(cf24) {
      let $dt = new D7(1);
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC18(cf26) {
      let $dt = new D7(2);
      $dt.cf26 = cf26;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get is_DC18() { return this.$tag === 2; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf26() { return this.cf26; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC17" + "(" + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC16" + "(" + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC18" + "(" + _dafny.toString(this.cf26) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf25 === other.cf25;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf24 === other.cf24;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf26, other.cf26);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC17(false);
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
    static create_DC20(cf28, cf29, cf30, cf31) {
      let $dt = new D8(0);
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      return $dt;
    }
    static create_DC19(cf27) {
      let $dt = new D8(1);
      $dt.cf27 = cf27;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get is_DC19() { return this.$tag === 1; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf27() { return this.cf27; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC20" + "(" + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC19" + "(" + _dafny.toString(this.cf27) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf28 === other.cf28 && this.cf29 === other.cf29 && this.cf30 === other.cf30 && this.cf31 === other.cf31;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf27 === other.cf27;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC20(false, false, false, false);
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
    static create_DC21(cf32) {
      let $dt = new D9(0);
      $dt.cf32 = cf32;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get dtor_cf32() { return this.cf32; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC21" + "(" + this.cf32.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf32, other.cf32);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Seq.UnicodeFromString("");
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
    static create_DC22(cf33) {
      let $dt = new D10(0);
      $dt.cf33 = cf33;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get dtor_cf33() { return this.cf33; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC22" + "(" + _dafny.toString(this.cf33) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf33, other.cf33);
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
          return D10.Default();
        }
      };
    }
  }

  $module.D11 = class D11 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC24(cf35) {
      let $dt = new D11(0);
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC25() {
      let $dt = new D11(1);
      return $dt;
    }
    static create_DC26(cf36, cf37, cf38) {
      let $dt = new D11(2);
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC23(cf34) {
      let $dt = new D11(3);
      $dt.cf34 = cf34;
      return $dt;
    }
    static create_DC27(cf39) {
      let $dt = new D11(4);
      $dt.cf39 = cf39;
      return $dt;
    }
    get is_DC24() { return this.$tag === 0; }
    get is_DC25() { return this.$tag === 1; }
    get is_DC26() { return this.$tag === 2; }
    get is_DC23() { return this.$tag === 3; }
    get is_DC27() { return this.$tag === 4; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf39() { return this.cf39; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC24" + "(" + _dafny.toString(this.cf35) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC25";
      } else if (this.$tag === 2) {
        return "D11.DC26" + "(" + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 3) {
        return "D11.DC23" + "(" + _dafny.toString(this.cf34) + ")";
      } else if (this.$tag === 4) {
        return "D11.DC27" + "(" + _dafny.toString(this.cf39) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf35, other.cf35);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf36, other.cf36) && this.cf37 === other.cf37 && this.cf38 === other.cf38;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf34, other.cf34);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf39, other.cf39);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC24(_dafny.ZERO);
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
    static create_DC28(cf40) {
      let $dt = new D12(0);
      $dt.cf40 = cf40;
      return $dt;
    }
    get is_DC28() { return this.$tag === 0; }
    get dtor_cf40() { return this.cf40; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC28" + "(" + _dafny.toString(this.cf40) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf40 === other.cf40;
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
          return D12.Default();
        }
      };
    }
  }

  $module.D13 = class D13 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC29(cf41) {
      let $dt = new D13(0);
      $dt.cf41 = cf41;
      return $dt;
    }
    get is_DC29() { return this.$tag === 0; }
    get dtor_cf41() { return this.cf41; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC29" + "(" + _dafny.toString(this.cf41) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf41, other.cf41);
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
          return D13.Default();
        }
      };
    }
  }

  $module.D14 = class D14 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC30(cf42) {
      let $dt = new D14(0);
      $dt.cf42 = cf42;
      return $dt;
    }
    get is_DC30() { return this.$tag === 0; }
    get dtor_cf42() { return this.cf42; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC30" + "(" + _dafny.toString(this.cf42) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf42, other.cf42);
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
          return D14.Default();
        }
      };
    }
  }

  $module.D15 = class D15 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC32(cf44) {
      let $dt = new D15(0);
      $dt.cf44 = cf44;
      return $dt;
    }
    static create_DC31(cf43) {
      let $dt = new D15(1);
      $dt.cf43 = cf43;
      return $dt;
    }
    get is_DC32() { return this.$tag === 0; }
    get is_DC31() { return this.$tag === 1; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf43() { return this.cf43; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC32" + "(" + _dafny.toString(this.cf44) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC31" + "(" + _dafny.toString(this.cf43) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf44 === other.cf44;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf43, other.cf43);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC32(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D15.Default();
        }
      };
    }
  }

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f2 = _dafny.ZERO;
      this.f7 = _dafny.Seq.of();
      this.f10 = [];
      this._f0 = _dafny.ZERO;
      this._f1 = false;
      this._f3 = false;
      this._f4 = _dafny.Set.Empty;
      this._f5 = _dafny.Map.Empty;
      this._f6 = _dafny.Seq.of();
      this._f8 = false;
      this._f9 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this).f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
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
    get f8() {
      let _this = this;
      return _this._f8;
    };
    get f9() {
      let _this = this;
      return _this._f9;
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
    fm7(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((new BigNumber(107)).minus(new BigNumber((_dafny.Seq.UnicodeFromString("wt")).length))).isLessThan(new BigNumber(-225));
    };
    fm8(p0, p1, p2, globalState) {
      let _this = this;
      return ((new BigNumber(341)).minus(new BigNumber(-256))).plus(new BigNumber(((_dafny.Set.fromElements(new BigNumber(789), new BigNumber((_dafny.Seq.UnicodeFromString("lan")).length), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber(150))).length), new BigNumber((_dafny.Seq.of(false)).length))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ghqr")).length))).length),false)).length))).Intersect(_dafny.Set.fromElements(new BigNumber(330)))).length));
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this.f14 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f14) {
      let _this = this;
      (_this).f14 = f14;
      return;
    }
    m4(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = _module.D1.Default();
      let r2 = _dafny.ZERO;
      let _542_v0;
      _542_v0 = new BigNumber(-769);
      r2 = _542_v0;
      _542_v0 = (_542_v0).plus(new BigNumber(358));
      let _543_v1;
      _543_v1 = _dafny.Seq.of(new BigNumber(341), _542_v0);
      let _544_v2;
      _544_v2 = new _dafny.CodePoint('q'.codePointAt(0));
      let _545_v3;
      _545_v3 = _dafny.MultiSet.fromElements(_542_v0);
      let _hi1 = (new BigNumber((_545_v3).cardinality())).multipliedBy(_542_v0);
      for (let _546_i0 = (_543_v1)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("n"), _module.__default.safeIndex(_542_v0, new BigNumber((_dafny.Seq.UnicodeFromString("n")).length)), _544_v2)).length), new BigNumber((_543_v1).length))]; _546_i0.isLessThan(_hi1); _546_i0 = _546_i0.plus(_dafny.ONE)) {
        let _547_v4;
        let _nw87 = Array((new BigNumber(9)).toNumber()).fill([]);
        _547_v4 = _nw87;
        let _548_v5;
        let _nw88 = Array((new BigNumber(9)).toNumber()).fill(false);
        _548_v5 = _nw88;
        let _549_v6;
        _549_v6 = _dafny.MultiSet.fromElements(_548_v5, _548_v5);
        let _550_v7;
        _550_v7 = _dafny.Map.Empty.slice().updateUnsafe(true,_this.f14);
        let _551_v8;
        _551_v8 = _dafny.Seq.of(_550_v7, _550_v7, _550_v7, _550_v7);
        let _552_v9;
        _552_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_551_v8).length),_542_v0);
        let _553_v10;
        let _nw89 = Array((new BigNumber(19)).toNumber());
        _nw89[(_dafny.ZERO).toNumber()] = new BigNumber((_543_v1).length);
        _nw89[(_dafny.ONE).toNumber()] = new BigNumber(917);
        _nw89[(new BigNumber(2)).toNumber()] = _546_i0;
        _nw89[(new BigNumber(3)).toNumber()] = _542_v0;
        _nw89[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_this.f14, _this.f14)).length));
        _nw89[(new BigNumber(5)).toNumber()] = new BigNumber((_549_v6).cardinality());
        _nw89[(new BigNumber(6)).toNumber()] = _542_v0;
        _nw89[(new BigNumber(7)).toNumber()] = _546_i0;
        _nw89[(new BigNumber(8)).toNumber()] = new BigNumber(691);
        _nw89[(new BigNumber(9)).toNumber()] = new BigNumber((_550_v7).length);
        _nw89[(new BigNumber(10)).toNumber()] = _546_i0;
        _nw89[(new BigNumber(11)).toNumber()] = _542_v0;
        _nw89[(new BigNumber(12)).toNumber()] = _546_i0;
        _nw89[(new BigNumber(13)).toNumber()] = _546_i0;
        _nw89[(new BigNumber(14)).toNumber()] = new BigNumber(731);
        _nw89[(new BigNumber(15)).toNumber()] = _546_i0;
        _nw89[(new BigNumber(16)).toNumber()] = (((_552_v9).contains(_546_i0)) ? ((_552_v9).get(_546_i0)) : (_module.__default.fm4(globalState)));
        _nw89[(new BigNumber(17)).toNumber()] = _542_v0;
        _nw89[(new BigNumber(18)).toNumber()] = _546_i0;
        _553_v10 = _nw89;
        let _index84 = _module.__default.safeIndex(new BigNumber(770), new BigNumber((_547_v4).length));
        (_547_v4)[_index84] = _553_v10;
        let _554_v11;
        _554_v11 = _dafny.Map.Empty.slice().updateUnsafe(_546_i0,_this.f14);
        let _index85 = _module.__default.safeIndex(new BigNumber(770), new BigNumber((_547_v4).length));
        let _rhs81 = _553_v10;
        let _rhs82 = _this.f14;
        let _rhs83 = (!(_this.f14)) === ((((_554_v11).contains(_546_i0)) ? ((_554_v11).get(_546_i0)) : (_this.f14)));
        let _lhs59 = _547_v4;
        let _lhs60 = _module.__default.safeIndex(new BigNumber(770), new BigNumber((_547_v4).length));
        let _lhs61 = _this;
        let _lhs62 = _this;
        _lhs59[_lhs60] = _rhs81;
        _lhs61.f14 = _rhs82;
        _lhs62.f14 = _rhs83;
        let _555_v12;
        _555_v12 = _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p0).length),_this.f14)).length), new BigNumber((p0).length));
        let _556_v13;
        _556_v13 = _dafny.Map.Empty.slice().updateUnsafe(_555_v12,_542_v0);
        let _557_v14;
        _557_v14 = _dafny.MultiSet.fromElements((new BigNumber((_556_v13).length)).isLessThanOrEqualTo(_546_i0));
        _557_v14 = _557_v14;
        (globalState).f2 = _546_i0;
        (_this).f14 = ((_this.f14) ? (_this.f14) : (_this.f14));
      }
      let _558_v15;
      _558_v15 = _dafny.Map.Empty.slice().updateUnsafe(_542_v0,_this.f14);
      _558_v15 = (_558_v15).update((new BigNumber((_543_v1).length)).plus(_542_v0), _this.f14);
      let _559_v16;
      let _nw90 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Map.Empty);
      _559_v16 = _nw90;
      let _560_v17;
      _560_v17 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,!(_this.f14));
      let _index86 = _module.__default.safeIndex(new BigNumber(345), new BigNumber((_559_v16).length));
      (_559_v16)[_index86] = _560_v17;
      let _index87 = _module.__default.safeIndex(new BigNumber(345), new BigNumber((_559_v16).length));
      (_559_v16)[_index87] = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_this.f14);
      let _561_v18;
      _561_v18 = _dafny.Seq.of(!(_this.f14), _this.f14, _this.f14);
      let _562_v19;
      let _init13 = function (_563_i1) {
        return _this.f14;
      };
      let _nw91 = Array((new BigNumber(11)).toNumber());
      for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw91.length); _i0_13++) {
        _nw91[_i0_13] = _init13(new BigNumber(_i0_13));
      }
      _562_v19 = _nw91;
      let _564_v20;
      _564_v20 = _module.D3.create_DC10(true, _542_v0, _this.f14, _562_v19, _558_v15);
      let _pat_let_tv8 = _558_v15;
      let _pat_let_tv9 = _564_v20;
      let _pat_let_tv10 = globalState;
      let _source11 = function (_pat_let10_0) {
        return function (_565_dt__update__tmp_h0) {
          return function (_pat_let11_0) {
            return function (_566_dt__update_hcf5_h0) {
              return _module.D1.create_DC3(_566_dt__update_hcf5_h0);
            }(_pat_let11_0);
          }(_module.__default.fm6(new BigNumber((_pat_let_tv8).length), (_pat_let_tv9).dtor_cf14, !(_this.f14), new _dafny.CodePoint('a'.codePointAt(0)), _pat_let_tv10));
        }(_pat_let10_0);
      }(_module.__default.fm5(_542_v0, _542_v0, _542_v0, _561_v18, globalState));
      if (_source11.is_DC4) {
        let _567___mcc_h0 = (_source11).cf6;
        let _568___mcc_h1 = (_source11).cf7;
        let _569_cf7 = _568___mcc_h1;
        let _570_cf6 = _567___mcc_h0;
        let _571_v21;
        let _nw92 = new _module.C0();
        _nw92.__ctor();
        _571_v21 = _nw92;
        let _572_v22;
        _572_v22 = _dafny.Seq.of(_562_v19);
        let _573_v23;
        let _nw93 = Array((new BigNumber(23)).toNumber());
        _nw93[(_dafny.ZERO).toNumber()] = _562_v19;
        _nw93[(_dafny.ONE).toNumber()] = ((_this.f14) ? (_562_v19) : (_562_v19));
        _nw93[(new BigNumber(2)).toNumber()] = _562_v19;
        _nw93[(new BigNumber(3)).toNumber()] = _562_v19;
        _nw93[(new BigNumber(4)).toNumber()] = _562_v19;
        _nw93[(new BigNumber(5)).toNumber()] = _562_v19;
        _nw93[(new BigNumber(6)).toNumber()] = _562_v19;
        _nw93[(new BigNumber(7)).toNumber()] = (_564_v20).dtor_cf17;
        _nw93[(new BigNumber(8)).toNumber()] = _562_v19;
        _nw93[(new BigNumber(9)).toNumber()] = _562_v19;
        _nw93[(new BigNumber(10)).toNumber()] = _562_v19;
        _nw93[(new BigNumber(11)).toNumber()] = _562_v19;
        _nw93[(new BigNumber(12)).toNumber()] = (_572_v22)[_module.__default.safeIndex(_569_cf7, new BigNumber((_572_v22).length))];
        _nw93[(new BigNumber(13)).toNumber()] = _562_v19;
        _nw93[(new BigNumber(14)).toNumber()] = _562_v19;
        _nw93[(new BigNumber(15)).toNumber()] = _562_v19;
        _nw93[(new BigNumber(16)).toNumber()] = _562_v19;
        _nw93[(new BigNumber(17)).toNumber()] = _562_v19;
        _nw93[(new BigNumber(18)).toNumber()] = _562_v19;
        _nw93[(new BigNumber(19)).toNumber()] = _562_v19;
        _nw93[(new BigNumber(20)).toNumber()] = _562_v19;
        _nw93[(new BigNumber(21)).toNumber()] = _562_v19;
        _nw93[(new BigNumber(22)).toNumber()] = _562_v19;
        _573_v23 = _nw93;
        let _index88 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_573_v23).length));
        (_573_v23)[_index88] = _562_v19;
        let _index89 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_573_v23).length));
        (_573_v23)[_index89] = _562_v19;
        let _574_v24;
        _574_v24 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(_this.f14, _this.f14, _this.f14, _this.f14, globalState),_562_v19);
        _562_v19 = (((_574_v24).contains(_this.f14)) ? ((_574_v24).get(_this.f14)) : (_562_v19));
        _545_v3 = _545_v3;
      } else if (_source11.is_DC5) {
        let _575___mcc_h2 = (_source11).cf8;
        let _576___mcc_h3 = (_source11).cf9;
        let _577_cf9 = _576___mcc_h3;
        let _578_cf8 = _575___mcc_h2;
        let _579_v25;
        _579_v25 = _dafny.Map.Empty.slice().updateUnsafe(_542_v0,_dafny.Seq.UnicodeFromString("vsixj"));
        let _580_v26;
        _580_v26 = _dafny.Seq.UnicodeFromString("oukgvej");
        let _581_v27;
        _581_v27 = _dafny.Map.Empty.slice().updateUnsafe(_542_v0,((_578_cf8) ? (new BigNumber(883)) : (new BigNumber(((((_579_v25).contains(_542_v0)) ? ((_579_v25).get(_542_v0)) : (_580_v26))).length))));
        _581_v27 = (_581_v27).update(_542_v0, new BigNumber((_dafny.Seq.UnicodeFromString("lysyusnd")).length));
        let _582_v28;
        _582_v28 = _dafny.Map.Empty.slice().updateUnsafe(((_578_cf8) ? (false) : (_this.f14)),_542_v0);
        _582_v28 = (_582_v28).update((_this.f14) || (_578_cf8), new BigNumber(177));
        let _583_v29;
        let _nw94 = Array((new BigNumber(15)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _583_v29 = _nw94;
        let _index90 = _module.__default.safeIndex(new BigNumber(725), new BigNumber((_583_v29).length));
        (_583_v29)[_index90] = _544_v2;
        let _index91 = _module.__default.safeIndex(new BigNumber(725), new BigNumber((_583_v29).length));
        (_583_v29)[_index91] = _544_v2;
        let _584_v30;
        let _nw95 = new _module.C0();
        _nw95.__ctor();
        _584_v30 = _nw95;
      } else if (_source11.is_DC3) {
        let _585___mcc_h4 = (_source11).cf5;
        let _586_cf5 = _585___mcc_h4;
        let _587_v31;
        let _nw96 = new _module.C0();
        _nw96.__ctor();
        _587_v31 = _nw96;
        (_this).f14 = (_561_v18)[_module.__default.safeIndex((((_545_v3).contains(new BigNumber(((_559_v16)[_module.__default.safeIndex(new BigNumber(345), new BigNumber((_559_v16).length))]).length))) ? ((_545_v3).get(new BigNumber(((_559_v16)[_module.__default.safeIndex(new BigNumber(345), new BigNumber((_559_v16).length))]).length))) : (_542_v0)), new BigNumber((_561_v18).length))];
        let _index92 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_562_v19).length));
        (_562_v19)[_index92] = false;
        let _588_v32;
        _588_v32 = _dafny.Map.Empty.slice().updateUnsafe(_562_v19,_542_v0);
        let _589_v33;
        _589_v33 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_562_v19,_542_v0), _dafny.Map.Empty.slice().updateUnsafe(_562_v19,(_dafny.ZERO).minus(_542_v0)));
        let _index93 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_562_v19).length));
        let _rhs84 = _this.f14;
        let _rhs85 = (((false) ? (_588_v32) : (_588_v32))).Merge((_589_v33)[_module.__default.safeIndex((_dafny.ZERO).minus(_542_v0), new BigNumber((_589_v33).length))]);
        let _rhs86 = (_this.f14) === (true);
        let _rhs87 = _this.f14;
        let _rhs88 = _module.__default.fm9(_this.f14, globalState);
        let _lhs63 = _562_v19;
        let _lhs64 = _module.__default.safeIndex(new BigNumber(870), new BigNumber((_562_v19).length));
        let _lhs65 = _this;
        let _lhs66 = _this;
        _lhs63[_lhs64] = _rhs84;
        _588_v32 = _rhs85;
        _lhs65.f14 = _rhs86;
        _lhs66.f14 = _rhs87;
        r0 = _rhs88;
        let _590_v34;
        _590_v34 = _dafny.Seq.UnicodeFromString("ihdetle");
        let _591_v35;
        _591_v35 = _dafny.Set.fromElements((_dafny.ZERO).minus(_542_v0), _542_v0);
        let _592_v36;
        _592_v36 = _dafny.MultiSet.fromElements(_590_v34);
        let _593_v37;
        _593_v37 = _module.D1.create_DC4(_592_v36, _542_v0);
        let _594_v38;
        _594_v38 = _dafny.Map.Empty.slice().updateUnsafe((_562_v19)[_module.__default.safeIndex(new BigNumber(870), new BigNumber((_562_v19).length))],_561_v18);
        let _595_v39;
        let _nw97 = Array((new BigNumber(17)).toNumber());
        _nw97[(_dafny.ZERO).toNumber()] = _542_v0;
        _nw97[(_dafny.ONE).toNumber()] = _542_v0;
        _nw97[(new BigNumber(2)).toNumber()] = _542_v0;
        _nw97[(new BigNumber(3)).toNumber()] = _542_v0;
        _nw97[(new BigNumber(4)).toNumber()] = new BigNumber(137);
        _nw97[(new BigNumber(5)).toNumber()] = _module.__default.safeDivisionInt(_542_v0, _542_v0);
        _nw97[(new BigNumber(6)).toNumber()] = (new BigNumber((_590_v34).length)).plus(new BigNumber((_591_v35).length));
        _nw97[(new BigNumber(7)).toNumber()] = _542_v0;
        _nw97[(new BigNumber(8)).toNumber()] = (_593_v37).dtor_cf7;
        _nw97[(new BigNumber(9)).toNumber()] = (_542_v0).plus(new BigNumber((_594_v38).length));
        _nw97[(new BigNumber(10)).toNumber()] = (((_545_v3).contains(_542_v0)) ? ((_545_v3).get(_542_v0)) : (new BigNumber(745)));
        _nw97[(new BigNumber(11)).toNumber()] = _542_v0;
        _nw97[(new BigNumber(12)).toNumber()] = (_542_v0).plus(_542_v0);
        _nw97[(new BigNumber(13)).toNumber()] = _module.__default.fm4(globalState);
        _nw97[(new BigNumber(14)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(386), _542_v0);
        _nw97[(new BigNumber(15)).toNumber()] = _542_v0;
        _nw97[(new BigNumber(16)).toNumber()] = (_543_v1)[_module.__default.safeIndex(new BigNumber((_591_v35).length), new BigNumber((_543_v1).length))];
        _595_v39 = _nw97;
        let _index94 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_595_v39).length));
        (_595_v39)[_index94] = _542_v0;
        let _index95 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_595_v39).length));
        (_595_v39)[_index95] = _542_v0;
      } else {
        let _596___mcc_h5 = (_source11).cf10;
        let _597_cf10 = _596___mcc_h5;
        let _598_v40;
        let _init14 = ((_599_v2) => function (_600_i2) {
          return _module.D1.create_DC6(_module.D1.create_DC3(_599_v2));
        })(_544_v2);
        let _nw98 = Array((new BigNumber(8)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw98.length); _i0_14++) {
          _nw98[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _598_v40 = _nw98;
        let _601_v41;
        _601_v41 = _module.D2.create_DC7(_598_v40);
        _601_v41 = _601_v41;
        (_this).f14 = ((_this.f14) ? (_this.f14) : (_this.f14));
        let _602_v42;
        _602_v42 = _module.D2.create_DC8(new BigNumber(675));
        let _603_v43;
        _603_v43 = _dafny.MultiSet.fromElements(_602_v42);
        (globalState).f2 = new BigNumber((((_this.f14) ? ((_dafny.MultiSet.fromElements(_602_v42)).Difference(_dafny.MultiSet.fromElements(_602_v42, _602_v42))) : (((_dafny.MultiSet.fromElements(_602_v42)).update(_602_v42, _module.__default.abs(_542_v0))).Union(_603_v43)))).cardinality());
        let _604_v44;
        let _nw99 = new _module.C0();
        _nw99.__ctor();
        _604_v44 = _nw99;
      }
      let _605_v45;
      _605_v45 = _dafny.Seq.UnicodeFromString("nwhkcx");
      let _606_v46;
      _606_v46 = _dafny.Seq.of(_605_v45, _dafny.Seq.update(_605_v45, _module.__default.safeIndex((_dafny.ZERO).minus(_542_v0), new BigNumber((_605_v45).length)), _544_v2), _605_v45, _605_v45);
      r0 = (_606_v46)[_module.__default.safeIndex((new BigNumber((_543_v1).length)).multipliedBy(_542_v0), new BigNumber((_606_v46).length))];
      let _607_v47;
      _607_v47 = _module.D1.create_DC5(_this.f14, _dafny.Seq.IsProperPrefixOf(_605_v45, _605_v45));
      r1 = _607_v47;
      let _608_v48;
      _608_v48 = _dafny.Seq.of(_562_v19, _562_v19, _562_v19, _562_v19);
      r2 = _module.__default.safeModuloInt(new BigNumber((_608_v48).length), (_542_v0).multipliedBy(new BigNumber((_605_v45).length)));
      return [r0, r1, r2];
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this.f13 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f13) {
      let _this = this;
      (_this).f13 = f13;
      return;
    }
    m3(p0, p1, p2, globalState) {
      let _this = this;
      let _hi2 = _this.f13;
      for (let _609_i0 = _this.f13; _609_i0.isLessThan(_hi2); _609_i0 = _609_i0.plus(_dafny.ONE)) {
        let _610_v0;
        _610_v0 = true;
        let _611_v1;
        _611_v1 = new _dafny.CodePoint('e'.codePointAt(0));
        let _612_v2;
        _612_v2 = _dafny.Map.Empty.slice().updateUnsafe(_611_v1,_610_v0);
        let _613_v3;
        _613_v3 = _dafny.Map.Empty.slice().updateUnsafe(_609_i0,_610_v0);
        let _614_v4;
        _614_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_613_v3).length),_this.f13);
        (_this).f13 = (((_610_v0) ? (_this.f13) : (new BigNumber((_612_v2).length)))).plus((_609_i0).minus(new BigNumber((_614_v4).length)));
        let _615_v5;
        let _init15 = ((_616_v0) => function (_617_i1) {
          return _dafny.Seq.of(_616_v0);
        })(_610_v0);
        let _nw100 = Array((new BigNumber(21)).toNumber());
        for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw100.length); _i0_15++) {
          _nw100[_i0_15] = _init15(new BigNumber(_i0_15));
        }
        _615_v5 = _nw100;
        _615_v5 = _615_v5;
        let _618_v6;
        _618_v6 = _dafny.Seq.of(false);
        let _619_v7;
        _619_v7 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(true, !(_610_v0), _610_v0, (_618_v6)[_module.__default.safeIndex(_this.f13, new BigNumber((_618_v6).length))], globalState),_610_v0);
        let _620_v8;
        _620_v8 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p1).length),_609_i0));
        let _621_v9;
        let _nw101 = Array((new BigNumber(29)).toNumber());
        _nw101[(_dafny.ZERO).toNumber()] = new BigNumber((p1).length);
        _nw101[(_dafny.ONE).toNumber()] = (p2)[_module.__default.safeIndex(_609_i0, new BigNumber((p2).length))];
        _nw101[(new BigNumber(2)).toNumber()] = new BigNumber(908);
        _nw101[(new BigNumber(3)).toNumber()] = _609_i0;
        _nw101[(new BigNumber(4)).toNumber()] = new BigNumber((_619_v7).length);
        _nw101[(new BigNumber(5)).toNumber()] = (_this.f13).minus(_module.__default.fm2(_610_v0, globalState));
        _nw101[(new BigNumber(6)).toNumber()] = _609_i0;
        _nw101[(new BigNumber(7)).toNumber()] = _this.f13;
        _nw101[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(_609_i0);
        _nw101[(new BigNumber(9)).toNumber()] = new BigNumber(943);
        _nw101[(new BigNumber(10)).toNumber()] = _this.f13;
        _nw101[(new BigNumber(11)).toNumber()] = _this.f13;
        _nw101[(new BigNumber(12)).toNumber()] = new BigNumber((p1).length);
        _nw101[(new BigNumber(13)).toNumber()] = new BigNumber(917);
        _nw101[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_620_v8, _dafny.Seq.Create(_module.__default.abs(new BigNumber(681)), ((_622_i0) => function (_623_i2) {
          return _dafny.Map.Empty.slice().updateUnsafe(_622_i0,new BigNumber((_dafny.Set.fromElements(_this.f13)).length));
        })(_609_i0)))).length);
        _nw101[(new BigNumber(15)).toNumber()] = (new BigNumber(-545)).multipliedBy(_609_i0);
        _nw101[(new BigNumber(16)).toNumber()] = _609_i0;
        _nw101[(new BigNumber(17)).toNumber()] = _609_i0;
        _nw101[(new BigNumber(18)).toNumber()] = _this.f13;
        _nw101[(new BigNumber(19)).toNumber()] = _this.f13;
        _nw101[(new BigNumber(20)).toNumber()] = _609_i0;
        _nw101[(new BigNumber(21)).toNumber()] = _this.f13;
        _nw101[(new BigNumber(22)).toNumber()] = (_this.f13).minus(_this.f13);
        _nw101[(new BigNumber(23)).toNumber()] = _this.f13;
        _nw101[(new BigNumber(24)).toNumber()] = _this.f13;
        _nw101[(new BigNumber(25)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ngpktyvor"), p1)).length);
        _nw101[(new BigNumber(26)).toNumber()] = _609_i0;
        _nw101[(new BigNumber(27)).toNumber()] = _this.f13;
        _nw101[(new BigNumber(28)).toNumber()] = new BigNumber((_dafny.Seq.of(_610_v0)).length);
        _621_v9 = _nw101;
        let _index96 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_621_v9).length));
        (_621_v9)[_index96] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-155)), ((_624_v1) => function (_625_i3) {
          return _624_v1;
        })(_611_v1))).length);
        let _626_v10;
        _626_v10 = _module.D0.create_DC0(_610_v0);
        let _627_v11;
        _627_v11 = _dafny.Set.fromElements(_626_v10);
        let _628_v12;
        _628_v12 = _dafny.Map.Empty.slice().updateUnsafe(_610_v0,_627_v11);
        let _629_v13;
        _629_v13 = _dafny.Map.Empty.slice().updateUnsafe(_626_v10,new BigNumber(680));
        let _630_v15;
        _630_v15 = _dafny.Map.Empty.slice().updateUnsafe(false,(((_628_v12).contains(false)) ? ((_628_v12).get(false)) : (function () {
          let _coll15 = new _dafny.Set();
          for (const _compr_15 of (_629_v13).Keys.Elements) {
            let _631_v14 = _compr_15;
            if ((_629_v13).contains(_631_v14)) {
              _coll15.add(_631_v14);
            }
          }
          return _coll15;
        }())));
        let _632_v16;
        _632_v16 = _dafny.Set.fromElements(_610_v0, _610_v0, !(false), _610_v0, _610_v0);
        let _index97 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_621_v9).length));
        let _rhs89 = new BigNumber(((((_628_v12).contains(_610_v0)) ? ((_628_v12).get(_610_v0)) : (((_610_v0) ? (_627_v11) : (_627_v11))))).length);
        let _rhs90 = (_dafny.Set.fromElements(_610_v0, _610_v0)).IsProperSubsetOf(_632_v16);
        let _lhs67 = _621_v9;
        let _lhs68 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_621_v9).length));
        _lhs67[_lhs68] = _rhs89;
        _610_v0 = _rhs90;
        _610_v0 = _610_v0;
      }
      let _633_v17;
      _633_v17 = _dafny.Seq.of(p0);
      (_this).f13 = new BigNumber((_633_v17).length);
      let _634_v18;
      _634_v18 = true;
      let _rhs91 = !(_this.f13).isEqualTo(_this.f13);
      let _rhs92 = (_this.f13).isLessThan(_module.__default.safeDivisionInt(_this.f13, _this.f13));
      _634_v18 = _rhs91;
      _634_v18 = _rhs92;
      _634_v18 = false;
      let _hi3 = _module.__default.safeModuloInt(_this.f13, _this.f13);
      for (let _635_i4 = _this.f13; _635_i4.isLessThan(_hi3); _635_i4 = _635_i4.plus(_dafny.ONE)) {
        let _636_v19;
        _636_v19 = _dafny.Set.fromElements(_634_v18, false, true, _module.__default.fm0(_634_v18, _634_v18, _634_v18, _634_v18, globalState));
        if (!(!((_636_v19).IsSubsetOf(_636_v19)))) {
          let _637_v20;
          let _nw102 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _637_v20 = _nw102;
          let _638_v21;
          _638_v21 = new _dafny.CodePoint('l'.codePointAt(0));
          let _639_v22;
          _639_v22 = _dafny.Map.Empty.slice().updateUnsafe(_638_v21,p1);
          let _640_v23;
          _640_v23 = _module.D1.create_DC3(_638_v21);
          let _index98 = _module.__default.safeIndex(new BigNumber(281), new BigNumber((_637_v20).length));
          (_637_v20)[_index98] = (((_639_v22).contains((_640_v23).dtor_cf5)) ? ((_639_v22).get((_640_v23).dtor_cf5)) : (p1));
          let _index99 = _module.__default.safeIndex(new BigNumber(281), new BigNumber((_637_v20).length));
          (_637_v20)[_index99] = _module.__default.fm3(_634_v18, p1, _this.f13, globalState);
          let _641_v24;
          let _nw103 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
          _641_v24 = _nw103;
          let _index100 = _module.__default.safeIndex(new BigNumber(677), new BigNumber((_641_v24).length));
          (_641_v24)[_index100] = (_dafny.ZERO).minus(_635_i4);
          let _index101 = _module.__default.safeIndex(new BigNumber(677), new BigNumber((_641_v24).length));
          (_641_v24)[_index101] = _635_i4;
          let _642_v25;
          _642_v25 = _dafny.Seq.of(p2);
          _642_v25 = _642_v25;
          let _643_v26;
          _643_v26 = _dafny.Map.Empty.slice().updateUnsafe(!(_634_v18),_634_v18);
          let _644_v27;
          let _nw104 = Array((new BigNumber(21)).toNumber());
          _nw104[(_dafny.ZERO).toNumber()] = ((_634_v18) ? (_module.__default.fm0(_634_v18, _634_v18, _634_v18, !(_634_v18), globalState)) : (_634_v18));
          _nw104[(_dafny.ONE).toNumber()] = (_636_v19).IsProperSubsetOf(_636_v19);
          _nw104[(new BigNumber(2)).toNumber()] = !(_634_v18);
          _nw104[(new BigNumber(3)).toNumber()] = _634_v18;
          _nw104[(new BigNumber(4)).toNumber()] = _634_v18;
          _nw104[(new BigNumber(5)).toNumber()] = _634_v18;
          _nw104[(new BigNumber(6)).toNumber()] = _module.__default.fm0(_634_v18, _634_v18, _634_v18, _634_v18, globalState);
          _nw104[(new BigNumber(7)).toNumber()] = ((_634_v18) ? (_634_v18) : (_634_v18));
          _nw104[(new BigNumber(8)).toNumber()] = _module.__default.fm0(_634_v18, false, _634_v18, false, globalState);
          _nw104[(new BigNumber(9)).toNumber()] = _634_v18;
          _nw104[(new BigNumber(10)).toNumber()] = _634_v18;
          _nw104[(new BigNumber(11)).toNumber()] = (((_643_v26).contains(_634_v18)) ? ((_643_v26).get(_634_v18)) : (_634_v18));
          _nw104[(new BigNumber(12)).toNumber()] = (((_643_v26).contains(_634_v18)) ? ((_643_v26).get(_634_v18)) : (_634_v18));
          _nw104[(new BigNumber(13)).toNumber()] = _634_v18;
          _nw104[(new BigNumber(14)).toNumber()] = _634_v18;
          _nw104[(new BigNumber(15)).toNumber()] = _634_v18;
          _nw104[(new BigNumber(16)).toNumber()] = _634_v18;
          _nw104[(new BigNumber(17)).toNumber()] = _634_v18;
          _nw104[(new BigNumber(18)).toNumber()] = !(_634_v18);
          _nw104[(new BigNumber(19)).toNumber()] = _634_v18;
          _nw104[(new BigNumber(20)).toNumber()] = false;
          _644_v27 = _nw104;
          let _nw105 = Array((new BigNumber(17)).toNumber()).fill(false);
          _644_v27 = _nw105;
          let _645_v28;
          _645_v28 = _module.D1.create_DC5(_634_v18, _634_v18);
          let _646_v29;
          _646_v29 = _dafny.Seq.of(_645_v28);
          let _647_v30;
          let _nw106 = Array((new BigNumber(5)).toNumber()).fill(_module.D1.Default());
          _647_v30 = _nw106;
          let _648_v31;
          _648_v31 = _module.D2.create_DC7(_647_v30);
          let _649_v32;
          _649_v32 = _dafny.MultiSet.fromElements((_637_v20)[_module.__default.safeIndex(new BigNumber(281), new BigNumber((_637_v20).length))], p1);
          let _650_v33;
          _650_v33 = _module.D1.create_DC4(_649_v32, _635_i4);
          let _651_v34;
          _651_v34 = _dafny.Map.Empty.slice().updateUnsafe(_650_v33,_647_v30);
          let _652_v35;
          let _nw107 = Array((new BigNumber(29)).toNumber());
          _nw107[(_dafny.ZERO).toNumber()] = _647_v30;
          _nw107[(_dafny.ONE).toNumber()] = _647_v30;
          _nw107[(new BigNumber(2)).toNumber()] = _647_v30;
          _nw107[(new BigNumber(3)).toNumber()] = _647_v30;
          _nw107[(new BigNumber(4)).toNumber()] = _647_v30;
          _nw107[(new BigNumber(5)).toNumber()] = _647_v30;
          _nw107[(new BigNumber(6)).toNumber()] = _647_v30;
          _nw107[(new BigNumber(7)).toNumber()] = _647_v30;
          _nw107[(new BigNumber(8)).toNumber()] = _647_v30;
          _nw107[(new BigNumber(9)).toNumber()] = _647_v30;
          _nw107[(new BigNumber(10)).toNumber()] = _647_v30;
          _nw107[(new BigNumber(11)).toNumber()] = _647_v30;
          _nw107[(new BigNumber(12)).toNumber()] = _647_v30;
          _nw107[(new BigNumber(13)).toNumber()] = _647_v30;
          _nw107[(new BigNumber(14)).toNumber()] = ((!(_634_v18)) ? (_647_v30) : (_647_v30));
          _nw107[(new BigNumber(15)).toNumber()] = (_648_v31).dtor_cf11;
          _nw107[(new BigNumber(16)).toNumber()] = _647_v30;
          _nw107[(new BigNumber(17)).toNumber()] = (((_651_v34).contains(_650_v33)) ? ((_651_v34).get(_650_v33)) : (_647_v30));
          _nw107[(new BigNumber(18)).toNumber()] = _647_v30;
          _nw107[(new BigNumber(19)).toNumber()] = _647_v30;
          _nw107[(new BigNumber(20)).toNumber()] = _647_v30;
          _nw107[(new BigNumber(21)).toNumber()] = _647_v30;
          _nw107[(new BigNumber(22)).toNumber()] = _647_v30;
          _nw107[(new BigNumber(23)).toNumber()] = _647_v30;
          _nw107[(new BigNumber(24)).toNumber()] = _647_v30;
          _nw107[(new BigNumber(25)).toNumber()] = _647_v30;
          _nw107[(new BigNumber(26)).toNumber()] = _647_v30;
          _nw107[(new BigNumber(27)).toNumber()] = ((_634_v18) ? (_647_v30) : (_647_v30));
          _nw107[(new BigNumber(28)).toNumber()] = _647_v30;
          _652_v35 = _nw107;
          let _rhs93 = _dafny.Seq.update(_dafny.Seq.of(_module.D1.create_DC5(_634_v18, _634_v18)), _module.__default.safeIndex(_module.__default.safeDivisionInt(new BigNumber(620), _this.f13), new BigNumber((_dafny.Seq.of(_module.D1.create_DC5(_634_v18, _634_v18))).length)), _645_v28);
          let _rhs94 = !(_634_v18);
          let _rhs95 = _652_v35;
          _646_v29 = _rhs93;
          _634_v18 = _rhs94;
          _652_v35 = _rhs95;
        } else {
          let _653_v36;
          _653_v36 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(223),_634_v18);
          let _654_v37;
          _654_v37 = _module.D3.create_DC10(_634_v18, _this.f13, _634_v18, p0, _653_v36);
          let _655_v38;
          _655_v38 = new _dafny.CodePoint('i'.codePointAt(0));
          let _656_v39;
          _656_v39 = _dafny.Map.Empty.slice().updateUnsafe((_654_v37).dtor_cf17,_655_v38);
          _656_v39 = (_656_v39).update(p0, ((_634_v18) ? (_655_v38) : (_655_v38)));
          let _657_v40;
          let _nw108 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
          _657_v40 = _nw108;
          let _index102 = _module.__default.safeIndex(new BigNumber(425), new BigNumber((_657_v40).length));
          (_657_v40)[_index102] = _635_i4;
          let _index103 = _module.__default.safeIndex(new BigNumber(425), new BigNumber((_657_v40).length));
          let _rhs96 = _this.f13;
          let _rhs97 = _657_v40;
          let _rhs98 = _this.f13;
          let _lhs69 = _this;
          let _lhs70 = globalState;
          let _lhs71 = _657_v40;
          let _lhs72 = _module.__default.safeIndex(new BigNumber(425), new BigNumber((_657_v40).length));
          _lhs69.f13 = _rhs96;
          _lhs70.f10 = _rhs97;
          _lhs71[_lhs72] = _rhs98;
          let _658_v41;
          let _init16 = ((_659_v18) => function (_660_i5) {
            return _659_v18;
          })(_634_v18);
          let _nw109 = Array((new BigNumber(15)).toNumber());
          for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw109.length); _i0_16++) {
            _nw109[_i0_16] = _init16(new BigNumber(_i0_16));
          }
          _658_v41 = _nw109;
          _658_v41 = _658_v41;
          let _index104 = _module.__default.safeIndex(new BigNumber(425), new BigNumber((_657_v40).length));
          let _rhs99 = _634_v18;
          let _rhs100 = (_635_i4).multipliedBy(_module.__default.fm2(_634_v18, globalState));
          let _rhs101 = (_this.f13).plus(((_657_v40)[_module.__default.safeIndex(new BigNumber(425), new BigNumber((_657_v40).length))]).multipliedBy(_this.f13));
          let _rhs102 = _635_i4;
          let _rhs103 = _this.f13;
          let _lhs73 = _this;
          let _lhs74 = globalState;
          let _lhs75 = globalState;
          let _lhs76 = _657_v40;
          let _lhs77 = _module.__default.safeIndex(new BigNumber(425), new BigNumber((_657_v40).length));
          _634_v18 = _rhs99;
          _lhs73.f13 = _rhs100;
          _lhs74.f2 = _rhs101;
          _lhs75.f2 = _rhs102;
          _lhs76[_lhs77] = _rhs103;
          let _661_v42;
          _661_v42 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(new BigNumber(458), _this.f13),(new BigNumber((p2).length)).plus(_635_i4));
          _661_v42 = _661_v42;
        }
        let _662_v43;
        let _nw110 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Map.Empty);
        _662_v43 = _nw110;
        _662_v43 = _662_v43;
        let _663_v44;
        _663_v44 = new _dafny.CodePoint('a'.codePointAt(0));
        let _664_v45;
        _664_v45 = _dafny.Set.fromElements(_663_v44);
        let _665_v46;
        _665_v46 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p1).length),_634_v18);
        let _666_v47;
        _666_v47 = _module.D3.create_DC10(_634_v18, new BigNumber((_664_v45).length), _634_v18, (_633_v17)[_module.__default.safeIndex(_this.f13, new BigNumber((_633_v17).length))], _665_v46);
        if ((_666_v47).dtor_cf16) {
          let _667_v48;
          _667_v48 = _dafny.Map.Empty.slice().updateUnsafe(_634_v18,(_634_v18) || (_634_v18));
          _634_v18 = (((_667_v48).contains(!(!(_634_v18) || (_634_v18)))) ? ((_667_v48).get(!(!(_634_v18) || (_634_v18)))) : (false));
          (_this).f13 = _this.f13;
          let _668_v49;
          let _nw111 = Array((new BigNumber(17)).toNumber()).fill(false);
          _668_v49 = _nw111;
          let _nw112 = Array((new BigNumber(3)).toNumber()).fill(false);
          _668_v49 = _nw112;
          let _669_v50;
          _669_v50 = _dafny.MultiSet.fromElements(_635_i4);
          let _rhs104 = _this.f13;
          let _rhs105 = _this.f13;
          let _rhs106 = ((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(_635_i4)),_634_v18)).Merge((_665_v46).update(_this.f13, _634_v18))).contains((((_669_v50).contains(_module.__default.fm2(true, globalState))) ? ((_669_v50).get(_module.__default.fm2(true, globalState))) : (_635_i4)));
          let _rhs107 = new _dafny.CodePoint('c'.codePointAt(0));
          let _rhs108 = _634_v18;
          let _lhs78 = _this;
          let _lhs79 = globalState;
          _lhs78.f13 = _rhs104;
          _lhs79.f2 = _rhs105;
          _634_v18 = _rhs106;
          _663_v44 = _rhs107;
          _634_v18 = _rhs108;
          let _670_v51;
          let _nw113 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Map.Empty);
          _670_v51 = _nw113;
          let _671_v52;
          _671_v52 = _dafny.MultiSet.fromElements(_634_v18);
          _module.__default.m0(p0, _this.f13, _670_v51, new BigNumber((_671_v52).cardinality()), globalState);
        } else {
          let _index105 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((p0).length));
          (p0)[_index105] = _634_v18;
          let _672_v53;
          let _nw114 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Map.Empty);
          _672_v53 = _nw114;
          let _673_v54;
          _673_v54 = _dafny.Map.Empty.slice().updateUnsafe(p1,true);
          let _index106 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_672_v53).length));
          (_672_v53)[_index106] = _673_v54;
          let _index107 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((p0).length));
          let _index108 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_672_v53).length));
          let _rhs109 = _634_v18;
          let _rhs110 = _673_v54;
          let _lhs80 = p0;
          let _lhs81 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((p0).length));
          let _lhs82 = _672_v53;
          let _lhs83 = _module.__default.safeIndex(new BigNumber(729), new BigNumber((_672_v53).length));
          _lhs80[_lhs81] = _rhs109;
          _lhs82[_lhs83] = _rhs110;
          let _674_v55;
          let _nw115 = Array((new BigNumber(25)).toNumber()).fill([]);
          _674_v55 = _nw115;
          let _675_v56;
          let _nw116 = Array((new BigNumber(18)).toNumber()).fill(false);
          _675_v56 = _nw116;
          let _index109 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_674_v55).length));
          (_674_v55)[_index109] = _675_v56;
          let _index110 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_674_v55).length));
          (_674_v55)[_index110] = _675_v56;
          let _676_v57;
          _676_v57 = _dafny.MultiSet.fromElements((p0)[_module.__default.safeIndex(new BigNumber(573), new BigNumber((p0).length))], true, !((p0)[_module.__default.safeIndex(new BigNumber(573), new BigNumber((p0).length))]));
          _676_v57 = _676_v57;
          let _677_v58;
          let _nw117 = new _module.C1();
          _nw117.__ctor(_634_v18);
          _677_v58 = _nw117;
          let _678_v59;
          _678_v59 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-930)), ((_679_v44) => function (_680_i6) {
            return _679_v44;
          })(_663_v44)), p1);
          let _681_v60;
          _681_v60 = _dafny.Map.Empty.slice().updateUnsafe((p0)[_module.__default.safeIndex(new BigNumber(573), new BigNumber((p0).length))],true);
          let _682_v61;
          _682_v61 = _dafny.Map.Empty.slice().updateUnsafe(_677_v58.f14,new BigNumber((_681_v60).length));
          let _683_v62;
          _683_v62 = _dafny.Set.fromElements(new BigNumber((p2).length), (((_682_v61).contains(_module.__default.fm0(_634_v18, _677_v58.f14, _677_v58.f14, (p0)[_module.__default.safeIndex(new BigNumber(573), new BigNumber((p0).length))], globalState))) ? ((_682_v61).get(_module.__default.fm0(_634_v18, _677_v58.f14, _677_v58.f14, (p0)[_module.__default.safeIndex(new BigNumber(573), new BigNumber((p0).length))], globalState))) : (_635_i4)), _this.f13);
          let _684_v63;
          let _init17 = ((_685_i4) => function (_686_i7) {
            return (_686_i7).minus(_685_i4);
          })(_635_i4);
          let _nw118 = Array((new BigNumber(5)).toNumber());
          for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw118.length); _i0_17++) {
            _nw118[_i0_17] = _init17(new BigNumber(_i0_17));
          }
          _684_v63 = _nw118;
          let _687_v64;
          _687_v64 = _dafny.Seq.of(_684_v63, _684_v63, _684_v63, _684_v63);
          let _index111 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((p0).length));
          let _rhs111 = (_636_v19).IsDisjointFrom(_dafny.Set.fromElements(_634_v18));
          let _rhs112 = (p1)[_module.__default.safeIndex(new BigNumber((_683_v62).length), new BigNumber((p1).length))];
          let _rhs113 = new BigNumber((_dafny.Seq.Concat(_687_v64, _dafny.Seq.Concat(_dafny.Seq.of(_684_v63, _684_v63), _687_v64))).length);
          let _rhs114 = _dafny.Seq.update(_dafny.Seq.update(_678_v59, _module.__default.safeIndex(_635_i4, new BigNumber((_678_v59).length)), p1), _module.__default.safeIndex((_dafny.ZERO).minus(_this.f13), new BigNumber((_dafny.Seq.update(_678_v59, _module.__default.safeIndex(_635_i4, new BigNumber((_678_v59).length)), p1)).length)), _dafny.Seq.Concat(p1, p1));
          let _rhs115 = false;
          let _lhs84 = _this;
          let _lhs85 = p0;
          let _lhs86 = _module.__default.safeIndex(new BigNumber(573), new BigNumber((p0).length));
          _634_v18 = _rhs111;
          _663_v44 = _rhs112;
          _lhs84.f13 = _rhs113;
          _678_v59 = _rhs114;
          _lhs85[_lhs86] = _rhs115;
        }
        let _688_v65;
        _688_v65 = _dafny.Map.Empty.slice().updateUnsafe(_634_v18,_636_v19);
        let _689_v66;
        _689_v66 = _module.D4.create_DC12(_636_v19);
        let _690_v67;
        _690_v67 = _dafny.Map.Empty.slice().updateUnsafe(_635_i4,(_689_v66).dtor_cf20);
        let _691_v68;
        _691_v68 = _dafny.MultiSet.fromElements(_this.f13);
        let _692_v69;
        _692_v69 = _dafny.Seq.of(_634_v18, false);
        let _693_v70;
        _693_v70 = _dafny.Map.Empty.slice().updateUnsafe(_634_v18,(_692_v69)[_module.__default.safeIndex(_635_i4, new BigNumber((_692_v69).length))]);
        let _694_v71;
        _694_v71 = _dafny.Map.Empty.slice().updateUnsafe(!(_634_v18),_663_v44);
        let _695_v72;
        let _nw119 = Array((new BigNumber(29)).toNumber());
        _nw119[(_dafny.ZERO).toNumber()] = new BigNumber(647);
        _nw119[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.update(((_634_v18) ? (p1) : (p1)), _module.__default.safeIndex(_this.f13, new BigNumber((((_634_v18) ? (p1) : (p1))).length)), _663_v44)).length);
        _nw119[(new BigNumber(2)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(466)), function (_696_i8) {
          return new _dafny.CodePoint('b'.codePointAt(0));
        })).length), (p2)[_module.__default.safeIndex(_635_i4, new BigNumber((p2).length))]);
        _nw119[(new BigNumber(3)).toNumber()] = (_this.f13).minus(new BigNumber((_636_v19).length));
        _nw119[(new BigNumber(4)).toNumber()] = new BigNumber(((((_688_v65).contains(_634_v18)) ? ((_688_v65).get(_634_v18)) : ((((_690_v67).contains(_this.f13)) ? ((_690_v67).get(_this.f13)) : (_636_v19))))).length);
        _nw119[(new BigNumber(5)).toNumber()] = _635_i4;
        _nw119[(new BigNumber(6)).toNumber()] = _module.__default.safeModuloInt(_this.f13, new BigNumber(182));
        _nw119[(new BigNumber(7)).toNumber()] = _module.__default.safeModuloInt(_this.f13, _this.f13);
        _nw119[(new BigNumber(8)).toNumber()] = _635_i4;
        _nw119[(new BigNumber(9)).toNumber()] = _this.f13;
        _nw119[(new BigNumber(10)).toNumber()] = new BigNumber((p1).length);
        _nw119[(new BigNumber(11)).toNumber()] = (((_691_v68).contains(new BigNumber((_692_v69).length))) ? ((_691_v68).get(new BigNumber((_692_v69).length))) : (_635_i4));
        _nw119[(new BigNumber(12)).toNumber()] = _635_i4;
        _nw119[(new BigNumber(13)).toNumber()] = _this.f13;
        _nw119[(new BigNumber(14)).toNumber()] = _635_i4;
        _nw119[(new BigNumber(15)).toNumber()] = _635_i4;
        _nw119[(new BigNumber(16)).toNumber()] = new BigNumber(773);
        _nw119[(new BigNumber(17)).toNumber()] = new BigNumber((_693_v70).length);
        _nw119[(new BigNumber(18)).toNumber()] = _module.__default.fm4(globalState);
        _nw119[(new BigNumber(19)).toNumber()] = new BigNumber((_dafny.Seq.Concat(p1, p1)).length);
        _nw119[(new BigNumber(20)).toNumber()] = _this.f13;
        _nw119[(new BigNumber(21)).toNumber()] = _635_i4;
        _nw119[(new BigNumber(22)).toNumber()] = _635_i4;
        _nw119[(new BigNumber(23)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(_663_v44, _663_v44, _663_v44, (((_694_v71).contains(_634_v18)) ? ((_694_v71).get(_634_v18)) : (_663_v44)), _663_v44)).cardinality())));
        _nw119[(new BigNumber(24)).toNumber()] = _this.f13;
        _nw119[(new BigNumber(25)).toNumber()] = _this.f13;
        _nw119[(new BigNumber(26)).toNumber()] = _635_i4;
        _nw119[(new BigNumber(27)).toNumber()] = _635_i4;
        _nw119[(new BigNumber(28)).toNumber()] = _this.f13;
        _695_v72 = _nw119;
        let _index112 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_695_v72).length));
        (_695_v72)[_index112] = _this.f13;
        let _index113 = _module.__default.safeIndex(new BigNumber(878), new BigNumber((_695_v72).length));
        (_695_v72)[_index113] = _this.f13;
      }
      _634_v18 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_dafny.Seq.of(_this.f13), _dafny.Seq.update(p2, _module.__default.safeIndex(_this.f13, new BigNumber((p2).length)), _this.f13)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(802)), function (_697_i9) {
        return _this.f13;
      }));
      return;
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f11 = false;
      this._f12 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f11, f12) {
      let _this = this;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      return;
    }
    fm1(p0, p1, globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("ngc")).length), new BigNumber((_dafny.MultiSet.fromElements((_this).f11)).cardinality())), _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("ttwonxung")).length), new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(173)), function (_698_i0) {
        return (_this).f12;
      })))).cardinality())));
    };
    m1(p0, globalState) {
      let _this = this;
      let _699_v0;
      _699_v0 = _module.D0.create_DC0(p0);
      let _700_v1;
      _700_v1 = _module.D0.create_DC2(_699_v0);
      let _701_v2;
      let _nw120 = Array((new BigNumber(14)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _701_v2 = _nw120;
      let _702_v3;
      _702_v3 = _dafny.Map.Empty.slice().updateUnsafe(_700_v1,_701_v2);
      if ((_702_v3).contains(_module.D0.create_DC2(_699_v0))) {
        let _703_v4;
        let _nw121 = new _module.C2();
        _nw121.__ctor((_dafny.ZERO).minus((_this).f12));
        _703_v4 = _nw121;
        let _704_v5;
        _704_v5 = _dafny.MultiSet.fromElements((_703_v4.f13).minus((_this).f12), (_this).f12);
        _704_v5 = _704_v5;
        let _705_v6;
        let _nw122 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
        _705_v6 = _nw122;
        let _706_v7;
        _706_v7 = _dafny.Seq.of(p0);
        let _index114 = _module.__default.safeIndex(new BigNumber(559), new BigNumber((_705_v6).length));
        (_705_v6)[_index114] = new BigNumber((_706_v7).length);
        let _index115 = _module.__default.safeIndex(new BigNumber(559), new BigNumber((_705_v6).length));
        (_705_v6)[_index115] = _module.__default.safeModuloInt(new BigNumber(-878), new BigNumber((_module.__default.fm10(globalState)).length));
        let _707_v8;
        _707_v8 = false;
        let _708_v9;
        _708_v9 = _dafny.Set.fromElements((_705_v6)[_module.__default.safeIndex(new BigNumber(559), new BigNumber((_705_v6).length))]);
        _707_v8 = !(!(_708_v9).equals(_708_v9));
        let _709_v10;
        _709_v10 = new _dafny.CodePoint('e'.codePointAt(0));
        let _710_v11;
        _710_v11 = _dafny.Seq.UnicodeFromString("wdfkv");
        if (_module.__default.fm0(_dafny.Seq.contains(_710_v11, _709_v10), (((_this).f11) ? (_707_v8) : ((_this).f11)), (_this).f11, (_this).f11, globalState)) {
          (globalState).f2 = ((_705_v6)[_module.__default.safeIndex(new BigNumber(559), new BigNumber((_705_v6).length))]).multipliedBy(_703_v4.f13);
          _707_v8 = (new BigNumber((_706_v7).length)).isLessThan((_this).f12);
          let _711_v12;
          let _nw123 = Array((new BigNumber(23)).toNumber()).fill(false);
          _711_v12 = _nw123;
          let _index116 = _module.__default.safeIndex(new BigNumber(560), new BigNumber((_711_v12).length));
          (_711_v12)[_index116] = p0;
          let _712_v13;
          _712_v13 = _dafny.MultiSet.fromElements(_706_v7);
          let _index117 = _module.__default.safeIndex(new BigNumber(560), new BigNumber((_711_v12).length));
          (_711_v12)[_index117] = (_712_v13).IsSubsetOf((_712_v13).Union(_712_v13));
          (_703_v4).f13 = (_705_v6)[_module.__default.safeIndex(new BigNumber(559), new BigNumber((_705_v6).length))];
          let _index118 = _module.__default.safeIndex(new BigNumber(559), new BigNumber((_705_v6).length));
          (_705_v6)[_index118] = ((_this).f12).multipliedBy(_703_v4.f13);
        } else {
          _700_v1 = _700_v1;
          let _713_v14;
          let _nw124 = new _module.C1();
          _nw124.__ctor((_704_v5).IsProperSubsetOf(_704_v5));
          _713_v14 = _nw124;
          let _714_v15;
          _714_v15 = _dafny.Map.Empty.slice().updateUnsafe((_705_v6)[_module.__default.safeIndex(new BigNumber(559), new BigNumber((_705_v6).length))],_709_v10);
          let _715_v16;
          _715_v16 = _dafny.Map.Empty.slice().updateUnsafe((((_714_v15).contains((_this).f12)) ? ((_714_v15).get((_this).f12)) : (_709_v10)),_dafny.Seq.contains(_710_v11, _709_v10));
          _707_v8 = (((_715_v16).contains(new _dafny.CodePoint('o'.codePointAt(0)))) ? ((_715_v16).get(new _dafny.CodePoint('o'.codePointAt(0)))) : ((_this).f11));
          let _716_v17;
          let _nw125 = Array((new BigNumber(9)).toNumber());
          _nw125[(_dafny.ZERO).toNumber()] = (_707_v8) === ((_706_v7)[_module.__default.safeIndex(_703_v4.f13, new BigNumber((_706_v7).length))]);
          _nw125[(_dafny.ONE).toNumber()] = (_this).f11;
          _nw125[(new BigNumber(2)).toNumber()] = !(_713_v14.f14);
          _nw125[(new BigNumber(3)).toNumber()] = (_this).f11;
          _nw125[(new BigNumber(4)).toNumber()] = (_706_v7)[_module.__default.safeIndex(new BigNumber(-640), new BigNumber((_706_v7).length))];
          _nw125[(new BigNumber(5)).toNumber()] = _713_v14.f14;
          _nw125[(new BigNumber(6)).toNumber()] = p0;
          _nw125[(new BigNumber(7)).toNumber()] = _713_v14.f14;
          _nw125[(new BigNumber(8)).toNumber()] = (_this).f11;
          _716_v17 = _nw125;
          _716_v17 = _716_v17;
          let _717_v18;
          let _init18 = ((_718_v7) => function (_719_i0) {
            return _718_v7;
          })(_706_v7);
          let _nw126 = Array((new BigNumber(17)).toNumber());
          for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw126.length); _i0_18++) {
            _nw126[_i0_18] = _init18(new BigNumber(_i0_18));
          }
          _717_v18 = _nw126;
          _717_v18 = _717_v18;
        }
      } else {
        let _720_v19;
        _720_v19 = _dafny.Seq.of((_this).f12, (_this).f12);
        let _721_v20;
        _721_v20 = _dafny.Seq.of((_this).f12, new BigNumber((_720_v19).length), new BigNumber(515));
        let _722_v21;
        _722_v21 = _dafny.MultiSet.fromElements(_dafny.Seq.IsProperPrefixOf(_720_v19, _721_v20));
        _722_v21 = ((p0) ? (_722_v21) : (_722_v21));
        let _723_v22;
        let _nw127 = new _module.C2();
        _nw127.__ctor(_module.__default.fm4(globalState));
        _723_v22 = _nw127;
        _723_v22 = _723_v22;
        let _724_v23;
        _724_v23 = true;
        _724_v23 = _724_v23;
        let _725_v24;
        _725_v24 = _dafny.Map.Empty.slice().updateUnsafe(_723_v22.f13,((_this).f12).minus((_this).f12));
        let _rhs116 = _724_v23;
        let _rhs117 = p0;
        let _rhs118 = (new BigNumber(80)).isLessThan((_this).f12);
        let _rhs119 = p0;
        let _rhs120 = (_725_v24).Merge(_725_v24);
        _724_v23 = _rhs116;
        _724_v23 = _rhs117;
        _724_v23 = _rhs118;
        _724_v23 = _rhs119;
        _725_v24 = _rhs120;
        let _726_v25;
        let _nw128 = new _module.C1();
        _nw128.__ctor((_this).f11);
        _726_v25 = _nw128;
      }
      let _727_v26;
      _727_v26 = false;
      _727_v26 = !(p0);
      let _728_i1;
      _728_i1 = _dafny.ZERO;
      L1: {
        while (p0) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_728_i1)) {
              break L1;
            }
            _728_i1 = (_728_i1).plus(_dafny.ONE);
            let _729_v27;
            _729_v27 = _dafny.MultiSet.fromElements((_this).f11, (_this).f11);
            let _730_v28;
            _730_v28 = _dafny.Seq.UnicodeFromString("dcu");
            let _731_v29;
            _731_v29 = new _dafny.CodePoint('h'.codePointAt(0));
            let _732_v30;
            _732_v30 = _dafny.MultiSet.fromElements(_dafny.Seq.update(_730_v28, _module.__default.safeIndex((_this).f12, new BigNumber((_730_v28).length)), _731_v29), _dafny.Seq.UnicodeFromString("rbbmc"));
            let _733_v31;
            let _734_v32;
            let _735_v33;
            let _out27;
            let _out28;
            let _out29;
            let _outcollector9 = (_this).m2((new BigNumber((_729_v27).cardinality())).isLessThanOrEqualTo((_this).f12), (_732_v30).IsProperSubsetOf(_732_v30), globalState);
            _out27 = _outcollector9[0];
            _out28 = _outcollector9[1];
            _out29 = _outcollector9[2];
            _733_v31 = _out27;
            _734_v32 = _out28;
            _735_v33 = _out29;
            let _736_v34;
            _736_v34 = _dafny.Map.Empty.slice().updateUnsafe(!(_727_v26),_dafny.Set.fromElements((_this).f11, _733_v31, _733_v31));
            let _737_v35;
            _737_v35 = _dafny.Set.fromElements((_this).f11);
            let _738_v36;
            _738_v36 = _module.D4.create_DC12(_737_v35);
            _736_v34 = (_736_v34).update(p0, (_738_v36).dtor_cf20);
            _730_v28 = _dafny.Seq.Concat(_730_v28, _dafny.Seq.UnicodeFromString("ycmdjta"));
            let _739_v37;
            _739_v37 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(552),_727_v26);
            (globalState).f2 = ((_this).f12).plus(_module.__default.safeModuloInt((_this).f12, new BigNumber((_739_v37).length)));
          }
        }
      }
      let _hi4 = (_dafny.ZERO).minus((_this).f12);
      for (let _740_i2 = _module.__default.safeModuloInt((_this).f12, (_this).f12); _740_i2.isLessThan(_hi4); _740_i2 = _740_i2.plus(_dafny.ONE)) {
        let _741_v38;
        _741_v38 = _dafny.Seq.of(true);
        _741_v38 = _dafny.Seq.Concat(_dafny.Seq.update(_741_v38, _module.__default.safeIndex(_740_i2, new BigNumber((_741_v38).length)), p0), _dafny.Seq.update(_dafny.Seq.update(_741_v38, _module.__default.safeIndex((_this).f12, new BigNumber((_741_v38).length)), (_this).f11), _module.__default.safeIndex((_dafny.ZERO).minus(_740_i2), new BigNumber((_dafny.Seq.update(_741_v38, _module.__default.safeIndex((_this).f12, new BigNumber((_741_v38).length)), (_this).f11)).length)), !((_this).f11)));
        let _742_v39;
        let _init19 = ((_743_v26) => function (_744_i3) {
          return _743_v26;
        })(_727_v26);
        let _nw129 = Array((new BigNumber(15)).toNumber());
        for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw129.length); _i0_19++) {
          _nw129[_i0_19] = _init19(new BigNumber(_i0_19));
        }
        _742_v39 = _nw129;
        let _745_v40;
        _745_v40 = _dafny.Map.Empty.slice().updateUnsafe(_742_v39,_742_v39);
        let _746_v41;
        _746_v41 = _dafny.Seq.of((_this).f12);
        let _747_v42;
        _747_v42 = _dafny.MultiSet.fromElements((_this).f12, new BigNumber((_746_v41).length));
        let _748_v43;
        _748_v43 = _dafny.Map.Empty.slice().updateUnsafe(_727_v26,_747_v42);
        let _749_v44;
        _749_v44 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,(_this).f11);
        let _750_v45;
        _750_v45 = _dafny.Map.Empty.slice().updateUnsafe((_this).f12,p0);
        let _751_v46;
        _751_v46 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_749_v44).length),new BigNumber((_750_v45).length));
        let _752_v47;
        let _nw130 = Array((new BigNumber(9)).toNumber());
        _nw130[(_dafny.ZERO).toNumber()] = (_this).f12;
        _nw130[(_dafny.ONE).toNumber()] = _740_i2;
        _nw130[(new BigNumber(2)).toNumber()] = new BigNumber((_751_v46).length);
        _nw130[(new BigNumber(3)).toNumber()] = _740_i2;
        _nw130[(new BigNumber(4)).toNumber()] = (_this).f12;
        _nw130[(new BigNumber(5)).toNumber()] = _740_i2;
        _nw130[(new BigNumber(6)).toNumber()] = (_this).f12;
        _nw130[(new BigNumber(7)).toNumber()] = _740_i2;
        _nw130[(new BigNumber(8)).toNumber()] = _740_i2;
        _752_v47 = _nw130;
        let _753_v48;
        _753_v48 = new _dafny.CodePoint('e'.codePointAt(0));
        let _754_v49;
        _754_v49 = _dafny.Seq.of(_753_v48);
        let _755_v50;
        _755_v50 = _dafny.Seq.of(_747_v42);
        let _756_v51;
        _756_v51 = _module.D4.create_DC13(_753_v48);
        let _757_v52;
        _757_v52 = _dafny.MultiSet.fromElements(_756_v51);
        let _758_v53;
        _758_v53 = _757_v52;
        let _759_v54;
        _759_v54 = _dafny.Set.fromElements(((((_748_v43).contains(_727_v26)) ? ((_748_v43).get(_727_v26)) : (_dafny.MultiSet.fromElements(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_740_i2,_752_v47))).length))))).update(new BigNumber((_754_v49).length), _module.__default.abs(new BigNumber(-13))), (_747_v42).Union(_dafny.MultiSet.fromElements((_this).f12, (_this).f12)), _747_v42, (_755_v50)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_758_v53),(_this).f11)).length), _740_i2, (_this).f12)).cardinality()), new BigNumber((_755_v50).length))], (_747_v42).Difference(_747_v42));
        let _760_v55;
        _760_v55 = _dafny.Set.fromElements((_this).f11);
        let _rhs121 = ((_760_v55).Intersect(_760_v55)).IsDisjointFrom(_760_v55);
        let _rhs122 = _745_v40;
        let _rhs123 = _759_v54;
        _727_v26 = _rhs121;
        _745_v40 = _rhs122;
        _759_v54 = _rhs123;
        _750_v45 = (_750_v45).update(((_this).f12).minus((_this).f12), _727_v26);
        _727_v26 = true;
      }
      let _761_v56;
      _761_v56 = new _dafny.CodePoint('g'.codePointAt(0));
      if (_module.__default.fm0(!_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-689)), ((_762_v56) => function (_763_i4) {
        return _762_v56;
      })(_761_v56)), _761_v56), _727_v26, (_727_v26) && ((_this).f11), (_this).f11, globalState)) {
        let _764_v57;
        let _nw131 = new _module.C2();
        _nw131.__ctor(new BigNumber((_dafny.Seq.UnicodeFromString("my")).length));
        _764_v57 = _nw131;
        if ((_this).f11) {
          _727_v26 = !(_727_v26) || (!(!(p0)));
          let _765_v58;
          _765_v58 = _dafny.Seq.UnicodeFromString("xoku");
          let _766_v59;
          _766_v59 = _dafny.Set.fromElements(p0, !(!((_this).f11)), _727_v26, _727_v26, _727_v26);
          let _rhs124 = _dafny.Seq.UnicodeFromString("sedfiwcj");
          let _rhs125 = !(_766_v59).equals(_766_v59);
          _765_v58 = _rhs124;
          _727_v26 = _rhs125;
          let _767_v60;
          _767_v60 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(!(_727_v26))).length));
          _727_v26 = !(((_767_v60).Union(_dafny.Set.fromElements(new BigNumber(50), _764_v57.f13))).IsSubsetOf(_767_v60));
          _727_v26 = _727_v26;
          let _768_v61;
          _768_v61 = _module.D0.create_DC0(true);
          let _769_v62;
          _769_v62 = _dafny.Seq.of(!(p0) || ((_768_v61).dtor_cf0), !((_this).f11));
          _769_v62 = _769_v62;
        } else {
          let _770_v63;
          let _nw132 = new _module.C2();
          _nw132.__ctor(_764_v57.f13);
          _770_v63 = _nw132;
          let _771_v64;
          let _772_v65;
          let _773_v66;
          let _out30;
          let _out31;
          let _out32;
          let _outcollector10 = (_this).m2(_727_v26, _727_v26, globalState);
          _out30 = _outcollector10[0];
          _out31 = _outcollector10[1];
          _out32 = _outcollector10[2];
          _771_v64 = _out30;
          _772_v65 = _out31;
          _773_v66 = _out32;
          let _774_v67;
          _774_v67 = _dafny.Map.Empty.slice().updateUnsafe(_773_v66,p0);
          let _775_v68;
          _775_v68 = _dafny.Seq.UnicodeFromString("yfahqywud");
          let _776_v70;
          _776_v70 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(80)), function (_777_i5) {
            return new _dafny.CodePoint('l'.codePointAt(0));
          }));
          let _778_v71;
          _778_v71 = _dafny.MultiSet.fromElements(_764_v57.f13, new BigNumber((_774_v67).length), new BigNumber((_775_v68).length), new BigNumber((_dafny.Seq.of(new BigNumber((function () {
            let _coll16 = new _dafny.Map();
            for (const _compr_16 of (_776_v70).Elements) {
              let _779_v69 = _compr_16;
              if (_dafny.Seq.contains(_776_v70, _779_v69)) {
                _coll16.push([_779_v69,(_this).f11]);
              }
            }
            return _coll16;
          }()).length))).length));
          let _780_v72;
          let _nw133 = new _module.C1();
          _nw133.__ctor((_778_v71).IsSubsetOf(_module.__default.fm11(_module.__default.fm0(_727_v26, (_this).f11, p0, (_this).f11, globalState), _761_v56, globalState)));
          _780_v72 = _nw133;
          _780_v72 = _780_v72;
          let _781_v73;
          let _nw134 = new _module.C1();
          _nw134.__ctor(p0);
          _781_v73 = _nw134;
          (_780_v72).f14 = _771_v64;
        }
        (globalState).f2 = _module.__default.safeModuloInt(_764_v57.f13, (_this).f12);
        _727_v26 = _727_v26;
        let _782_v74;
        let _nw135 = new _module.C1();
        _nw135.__ctor((_764_v57.f13).isEqualTo(new BigNumber(184)));
        _782_v74 = _nw135;
      } else {
        let _783_v75;
        _783_v75 = _dafny.MultiSet.fromElements(_761_v56, _761_v56);
        _727_v26 = (_783_v75).IsDisjointFrom((_783_v75).Union(_783_v75));
        let _784_v76;
        let _nw136 = new _module.C2();
        _nw136.__ctor((_this).f12);
        _784_v76 = _nw136;
        let _785_v77;
        _785_v77 = _dafny.Seq.of((_this).f11, (_this).f11);
        let _786_v78;
        _786_v78 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_785_v77).length),(_this).f11);
        let _787_v79;
        _787_v79 = _dafny.MultiSet.fromElements((_this).f11, _727_v26, (_this).f11);
        let _788_v80;
        _788_v80 = _dafny.Seq.UnicodeFromString("onxfcxpf");
        let _789_v81;
        let _nw137 = Array((new BigNumber(6)).toNumber()).fill(false);
        _789_v81 = _nw137;
        let _790_v82;
        let _nw138 = Array((new BigNumber(15)).toNumber());
        _nw138[(_dafny.ZERO).toNumber()] = (_this).f12;
        _nw138[(_dafny.ONE).toNumber()] = (_this).f12;
        _nw138[(new BigNumber(2)).toNumber()] = new BigNumber((_787_v79).cardinality());
        _nw138[(new BigNumber(3)).toNumber()] = (new BigNumber(584)).minus(_784_v76.f13);
        _nw138[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus((_784_v76.f13).plus((_this).f12));
        _nw138[(new BigNumber(5)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_788_v80).length), _784_v76.f13);
        _nw138[(new BigNumber(6)).toNumber()] = (_this).f12;
        _nw138[(new BigNumber(7)).toNumber()] = ((_this).f12).multipliedBy(_784_v76.f13);
        _nw138[(new BigNumber(8)).toNumber()] = (_this).f12;
        _nw138[(new BigNumber(9)).toNumber()] = (((_this).f11) ? ((_this).f12) : (_784_v76.f13));
        _nw138[(new BigNumber(10)).toNumber()] = (_this).fm1(_dafny.Seq.UnicodeFromString("lnlmx"), _module.__default.fm2((_this).f11, globalState), globalState);
        _nw138[(new BigNumber(11)).toNumber()] = (_this).f12;
        _nw138[(new BigNumber(12)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_789_v81,true)).length);
        _nw138[(new BigNumber(13)).toNumber()] = (_this).f12;
        _nw138[(new BigNumber(14)).toNumber()] = (_this).f12;
        _790_v82 = _nw138;
        let _index119 = _module.__default.safeIndex(new BigNumber(700), new BigNumber((_790_v82).length));
        (_790_v82)[_index119] = _784_v76.f13;
        let _791_v83;
        let _nw139 = new _module.C0();
        _nw139.__ctor();
        _791_v83 = _nw139;
        let _792_v84;
        _792_v84 = _module.D7.create_DC16(_791_v83);
        let _793_v85;
        let _nw140 = Array((new BigNumber(27)).toNumber());
        _nw140[(_dafny.ZERO).toNumber()] = _791_v83;
        _nw140[(_dafny.ONE).toNumber()] = _791_v83;
        _nw140[(new BigNumber(2)).toNumber()] = _791_v83;
        _nw140[(new BigNumber(3)).toNumber()] = _791_v83;
        _nw140[(new BigNumber(4)).toNumber()] = _791_v83;
        _nw140[(new BigNumber(5)).toNumber()] = _791_v83;
        _nw140[(new BigNumber(6)).toNumber()] = _791_v83;
        _nw140[(new BigNumber(7)).toNumber()] = _791_v83;
        _nw140[(new BigNumber(8)).toNumber()] = _791_v83;
        _nw140[(new BigNumber(9)).toNumber()] = _791_v83;
        _nw140[(new BigNumber(10)).toNumber()] = _791_v83;
        _nw140[(new BigNumber(11)).toNumber()] = _791_v83;
        _nw140[(new BigNumber(12)).toNumber()] = _791_v83;
        _nw140[(new BigNumber(13)).toNumber()] = _791_v83;
        _nw140[(new BigNumber(14)).toNumber()] = _791_v83;
        _nw140[(new BigNumber(15)).toNumber()] = _791_v83;
        _nw140[(new BigNumber(16)).toNumber()] = _791_v83;
        _nw140[(new BigNumber(17)).toNumber()] = _791_v83;
        _nw140[(new BigNumber(18)).toNumber()] = _791_v83;
        _nw140[(new BigNumber(19)).toNumber()] = _791_v83;
        _nw140[(new BigNumber(20)).toNumber()] = _791_v83;
        _nw140[(new BigNumber(21)).toNumber()] = _791_v83;
        _nw140[(new BigNumber(22)).toNumber()] = _791_v83;
        _nw140[(new BigNumber(23)).toNumber()] = _791_v83;
        _nw140[(new BigNumber(24)).toNumber()] = (_792_v84).dtor_cf24;
        _nw140[(new BigNumber(25)).toNumber()] = _791_v83;
        _nw140[(new BigNumber(26)).toNumber()] = _791_v83;
        _793_v85 = _nw140;
        let _794_v86;
        _794_v86 = _dafny.Seq.of(_793_v85, _793_v85, _793_v85, _793_v85);
        let _index120 = _module.__default.safeIndex(new BigNumber(700), new BigNumber((_790_v82).length));
        let _rhs126 = (_module.D8.create_DC19(_784_v76)).dtor_cf27;
        let _rhs127 = (_this).f11;
        let _rhs128 = (_786_v78).Merge(((_module.D3.create_DC10(_727_v26, (_this).f12, p0, _789_v81, _786_v78)).dtor_cf18).Merge(_786_v78));
        let _rhs129 = (_this).f12;
        let _rhs130 = _dafny.Seq.Concat(_794_v86, _794_v86);
        let _lhs87 = _790_v82;
        let _lhs88 = _module.__default.safeIndex(new BigNumber(700), new BigNumber((_790_v82).length));
        _784_v76 = _rhs126;
        _727_v26 = _rhs127;
        _786_v78 = _rhs128;
        _lhs87[_lhs88] = _rhs129;
        _794_v86 = _rhs130;
        let _795_v87;
        _795_v87 = _module.D4.create_DC13(new _dafny.CodePoint('f'.codePointAt(0)));
        let _796_v88;
        _796_v88 = _dafny.Map.Empty.slice().updateUnsafe(p0,_795_v87);
        let _797_v89;
        _797_v89 = _dafny.Seq.of(_784_v76.f13, new BigNumber((_796_v88).length));
        let _798_v90;
        _798_v90 = _dafny.Seq.of(new BigNumber((_797_v89).length), _784_v76.f13);
        let _799_v91;
        _799_v91 = _dafny.Seq.of(new BigNumber((_798_v90).length));
        (_784_v76).m3(_789_v81, _788_v80, _dafny.Seq.Concat(_797_v89, _dafny.Seq.Create(_module.__default.abs(new BigNumber(911)), function (_800_i6) {
          return new BigNumber(-298);
        })), globalState);
        if (_727_v26) {
          (_784_v76).f13 = _784_v76.f13;
          let _rhs131 = (_784_v76.f13).multipliedBy(((_this).f12).plus(new BigNumber((_799_v91).length)));
          let _rhs132 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ebyufrei"), _788_v80);
          let _lhs89 = _784_v76;
          _lhs89.f13 = _rhs131;
          _788_v80 = _rhs132;
          (globalState).f2 = (_790_v82)[_module.__default.safeIndex(new BigNumber(700), new BigNumber((_790_v82).length))];
          _788_v80 = _dafny.Seq.update(_788_v80, _module.__default.safeIndex((_this).f12, new BigNumber((_788_v80).length)), _761_v56);
          let _index121 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_789_v81).length));
          (_789_v81)[_index121] = p0;
          let _index122 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_789_v81).length));
          (_789_v81)[_index122] = true;
        } else {
          let _801_v92;
          _801_v92 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("wtrbr"));
          let _802_v93;
          _802_v93 = _module.D1.create_DC4((_801_v92).Difference(_801_v92), (_dafny.ZERO).minus((_790_v82)[_module.__default.safeIndex(new BigNumber(700), new BigNumber((_790_v82).length))]));
          _802_v93 = _802_v93;
          (globalState).f2 = ((_790_v82)[_module.__default.safeIndex(new BigNumber(700), new BigNumber((_790_v82).length))]).plus((_this).f12);
          let _index123 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_789_v81).length));
          (_789_v81)[_index123] = p0;
          let _index124 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_789_v81).length));
          (_789_v81)[_index124] = !_dafny.areEqual(_dafny.Seq.UnicodeFromString("brsji"), _788_v80);
          (globalState).f2 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt(_784_v76.f13, new BigNumber((function () {
            let _coll17 = new _dafny.Map();
            for (const _compr_17 of (_786_v78).Keys.Elements) {
              let _803_v94 = _compr_17;
              if ((_786_v78).contains(_803_v94)) {
                _coll17.push([_module.__default.safeDivisionInt(_803_v94, (_this).f12),_761_v56]);
              }
            }
            return _coll17;
          }()).length))));
          let _804_v95;
          _804_v95 = _dafny.Map.Empty.slice().updateUnsafe(((_727_v26) ? (_784_v76.f13) : ((_790_v82)[_module.__default.safeIndex(new BigNumber(700), new BigNumber((_790_v82).length))])),_788_v80);
          _804_v95 = _804_v95;
        }
        let _805_v96;
        _805_v96 = _dafny.Set.fromElements(!(p0), (_this).f11);
        _727_v26 = (_805_v96).contains(!((_dafny.Set.fromElements((_797_v89)[_module.__default.safeIndex(_784_v76.f13, new BigNumber((_797_v89).length))], new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), (_this).f12)).IsDisjointFrom(_dafny.Set.fromElements((_790_v82)[_module.__default.safeIndex(new BigNumber(700), new BigNumber((_790_v82).length))]))));
      }
      let _806_v97;
      _806_v97 = _dafny.Seq.UnicodeFromString("ilyuyef");
      let _807_v98;
      _807_v98 = _dafny.MultiSet.fromElements(_727_v26);
      let _808_v99;
      _808_v99 = _dafny.Map.Empty.slice().updateUnsafe(_806_v97,new BigNumber((_807_v98).cardinality()));
      let _809_v100;
      let _nw141 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
      _809_v100 = _nw141;
      let _810_v101;
      _810_v101 = _dafny.Map.Empty.slice().updateUnsafe(_808_v99,_809_v100);
      let _811_v102;
      _811_v102 = _dafny.MultiSet.fromElements(new BigNumber(192));
      let _812_v103;
      _812_v103 = _dafny.MultiSet.fromElements(_811_v102, _811_v102, _dafny.MultiSet.fromElements((_this).f12), (_dafny.MultiSet.fromElements((_this).f12)).update((_this).f12, _module.__default.abs((_this).f12)), _811_v102);
      _810_v101 = (_810_v101).update(((_808_v99).update(_dafny.Seq.UnicodeFromString("ufaahfrp"), (_this).f12)).update(_dafny.Seq.UnicodeFromString("yaws"), (((_812_v103).contains(_811_v102)) ? ((_812_v103).get(_811_v102)) : (new BigNumber((_807_v98).cardinality())))), _809_v100);
      return;
    }
    m2(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Map.Empty;
      let r2 = false;
      let _813_v0;
      _813_v0 = _dafny.Seq.UnicodeFromString("iteghrvta");
      let _814_v1;
      _814_v1 = _dafny.Map.Empty.slice().updateUnsafe(_813_v0,(_this).f12);
      _814_v1 = (_814_v1).update(_813_v0, (_this).f12);
      let _815_v2;
      let _nw142 = Array((new BigNumber(23)).toNumber()).fill(false);
      _815_v2 = _nw142;
      let _816_v3;
      let _nw143 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Map.Empty);
      _816_v3 = _nw143;
      let _817_v4;
      _817_v4 = new _dafny.CodePoint('w'.codePointAt(0));
      let _818_v5;
      _818_v5 = _dafny.MultiSet.fromElements(_817_v4);
      let _819_v6;
      _819_v6 = _dafny.Seq.of(_module.__default.fm2(p0, globalState));
      _module.__default.m0(_815_v2, (_this).f12, _816_v3, (((_818_v5).contains(_817_v4)) ? ((_818_v5).get(_817_v4)) : (new BigNumber((_819_v6).length))), globalState);
      let _820_v7;
      let _nw144 = new _module.C0();
      _nw144.__ctor();
      _820_v7 = _nw144;
      let _821_v8;
      let _init20 = function (_822_i0) {
        return _module.__default.safeDivisionInt(_822_i0, (_this).f12);
      };
      let _nw145 = Array((new BigNumber(29)).toNumber());
      for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw145.length); _i0_20++) {
        _nw145[_i0_20] = _init20(new BigNumber(_i0_20));
      }
      _821_v8 = _nw145;
      (globalState).f10 = _821_v8;
      if (p0) {
        let _index125 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_815_v2).length));
        (_815_v2)[_index125] = (p0) === ((_this).f11);
        let _823_v9;
        _823_v9 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        let _824_v10;
        _824_v10 = _dafny.Set.fromElements((((_823_v9).contains(p1)) ? ((_823_v9).get(p1)) : (p1)), (_this).f11);
        let _825_v11;
        _825_v11 = _dafny.Map.Empty.slice().updateUnsafe(_824_v10,p0);
        let _index126 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_815_v2).length));
        (_815_v2)[_index126] = (((_825_v11).contains(_824_v10)) ? ((_825_v11).get(_824_v10)) : ((_this).f11));
        if ((_this).f11) {
          let _826_v12;
          _826_v12 = _dafny.MultiSet.fromElements(p1);
          r2 = !((_826_v12).IsProperSubsetOf(_826_v12)) || (p0);
          let _827_v13;
          _827_v13 = _dafny.MultiSet.fromElements(_824_v10);
          let _828_v14;
          _828_v14 = _dafny.Seq.of(!(!((_827_v13).IsDisjointFrom(_827_v13))));
          _828_v14 = _828_v14;
          let _index127 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_815_v2).length));
          let _rhs133 = _module.__default.fm6((_this).f12, p0, _module.__default.fm0((_this).f11, true, p1, (_this).f11, globalState), _817_v4, globalState);
          let _rhs134 = !((_815_v2)[_module.__default.safeIndex(new BigNumber(312), new BigNumber((_815_v2).length))]);
          let _lhs90 = _815_v2;
          let _lhs91 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_815_v2).length));
          _817_v4 = _rhs133;
          _lhs90[_lhs91] = _rhs134;
          let _829_v15;
          _829_v15 = _module.D1.create_DC5(p0, p0);
          r2 = (_820_v7).fm7(((p1) ? ((_829_v15).dtor_cf9) : ((_this).f11)), _module.D1.create_DC3(_817_v4), _813_v0, (_815_v2)[_module.__default.safeIndex(new BigNumber(312), new BigNumber((_815_v2).length))], globalState);
          let _index128 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_821_v8).length));
          (_821_v8)[_index128] = (_this).f12;
          let _index129 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_821_v8).length));
          (_821_v8)[_index129] = _module.__default.safeDivisionInt(_module.__default.safeModuloInt((_this).f12, (_this).f12), (_this).f12);
        } else {
          let _830_v16;
          let _nw146 = Array((new BigNumber(10)).toNumber()).fill(_module.D8.Default());
          _830_v16 = _nw146;
          let _831_v17;
          let _nw147 = new _module.C2();
          _nw147.__ctor((_this).f12);
          _831_v17 = _nw147;
          let _832_v18;
          _832_v18 = _module.D8.create_DC19(_831_v17);
          let _index130 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_830_v16).length));
          (_830_v16)[_index130] = ((p1) ? (_832_v18) : (_module.D8.create_DC19(_831_v17)));
          let _index131 = _module.__default.safeIndex(new BigNumber(588), new BigNumber((_830_v16).length));
          (_830_v16)[_index131] = _832_v18;
          let _833_v19;
          _833_v19 = _dafny.Map.Empty.slice().updateUnsafe(_831_v17.f13,new BigNumber((_813_v0).length));
          let _834_v20;
          _834_v20 = _dafny.Map.Empty.slice().updateUnsafe((_815_v2)[_module.__default.safeIndex(new BigNumber(312), new BigNumber((_815_v2).length))],_833_v19);
          _834_v20 = (_834_v20).update(false, _833_v19);
          let _835_v21;
          _835_v21 = _module.D7.create_DC17((_815_v2)[_module.__default.safeIndex(new BigNumber(312), new BigNumber((_815_v2).length))]);
          let _836_v22;
          _836_v22 = _module.D7.create_DC18(_835_v21);
          let _837_v23;
          _837_v23 = _module.D7.create_DC18(_835_v21);
          let _838_v24;
          _838_v24 = _dafny.Map.Empty.slice().updateUnsafe(_837_v23,(_815_v2)[_module.__default.safeIndex(new BigNumber(312), new BigNumber((_815_v2).length))]);
          let _index132 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_821_v8).length));
          (_821_v8)[_index132] = (_819_v6)[_module.__default.safeIndex(new BigNumber((_838_v24).length), new BigNumber((_819_v6).length))];
          let _index133 = _module.__default.safeIndex(new BigNumber(897), new BigNumber((_821_v8).length));
          (_821_v8)[_index133] = _831_v17.f13;
          let _839_v25;
          let _init21 = ((_840_v8) => function (_841_i1) {
            return (_841_i1).multipliedBy((_840_v8)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((_840_v8).length))]);
          })(_821_v8);
          let _nw148 = Array((new BigNumber(21)).toNumber());
          for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw148.length); _i0_21++) {
            _nw148[_i0_21] = _init21(new BigNumber(_i0_21));
          }
          _839_v25 = _nw148;
          let _index134 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_815_v2).length));
          let _rhs135 = _839_v25;
          let _rhs136 = _831_v17.f13;
          let _rhs137 = p0;
          let _rhs138 = (_821_v8)[_module.__default.safeIndex(new BigNumber(897), new BigNumber((_821_v8).length))];
          let _rhs139 = _dafny.Seq.Concat(_module.__default.fm12(_831_v17.f13, globalState), _dafny.Seq.update(_dafny.Seq.Concat(_819_v6, _819_v6), _module.__default.safeIndex((_this).f12, new BigNumber((_dafny.Seq.Concat(_819_v6, _819_v6)).length)), new BigNumber(406)));
          let _lhs92 = globalState;
          let _lhs93 = globalState;
          let _lhs94 = _815_v2;
          let _lhs95 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_815_v2).length));
          let _lhs96 = globalState;
          let _lhs97 = globalState;
          _lhs92.f10 = _rhs135;
          _lhs93.f2 = _rhs136;
          _lhs94[_lhs95] = _rhs137;
          _lhs96.f2 = _rhs138;
          _lhs97.f7 = _rhs139;
          let _842_v26;
          let _nw149 = Array((new BigNumber(26)).toNumber()).fill([]);
          _842_v26 = _nw149;
          let _843_v27;
          let _nw150 = Array((new BigNumber(12)).toNumber()).fill(false);
          _843_v27 = _nw150;
          let _index135 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_842_v26).length));
          (_842_v26)[_index135] = _843_v27;
          let _index136 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_842_v26).length));
          (_842_v26)[_index136] = _843_v27;
        }
        (globalState).f10 = _821_v8;
        _823_v9 = (_823_v9).update((_this).f11, (_815_v2)[_module.__default.safeIndex(new BigNumber(312), new BigNumber((_815_v2).length))]);
        let _844_v28;
        let _nw151 = new _module.C2();
        _nw151.__ctor(((_this).f12).multipliedBy((_this).f12));
        _844_v28 = _nw151;
      } else {
        let _845_v29;
        _845_v29 = _dafny.Map.Empty.slice().updateUnsafe((_this).f12,_821_v8);
        _845_v29 = _845_v29;
        (globalState).f7 = _819_v6;
        let _846_v30;
        let _nw152 = new _module.C2();
        _nw152.__ctor(new BigNumber((_813_v0).length));
        _846_v30 = _nw152;
        let _847_v31;
        _847_v31 = _dafny.Set.fromElements(_846_v30);
        (globalState).f2 = new BigNumber(((_dafny.Set.fromElements(_846_v30)).Difference(_847_v31)).length);
        let _848_v32;
        _848_v32 = _813_v0;
        let _849_v33;
        let _nw153 = Array((new BigNumber(13)).toNumber());
        _nw153[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat((_848_v32), _813_v0);
        _nw153[(_dafny.ONE).toNumber()] = _813_v0;
        _nw153[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_813_v0, _813_v0);
        _nw153[(new BigNumber(3)).toNumber()] = _813_v0;
        _nw153[(new BigNumber(4)).toNumber()] = _813_v0;
        _nw153[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("kctfg"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(908)), ((_850_v4) => function (_851_i2) {
          return _850_v4;
        })(_817_v4)));
        _nw153[(new BigNumber(6)).toNumber()] = _813_v0;
        _nw153[(new BigNumber(7)).toNumber()] = _813_v0;
        _nw153[(new BigNumber(8)).toNumber()] = _813_v0;
        _nw153[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_813_v0, _813_v0);
        _nw153[(new BigNumber(10)).toNumber()] = _813_v0;
        _nw153[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_813_v0, _module.__default.fm9(p1, globalState));
        _nw153[(new BigNumber(12)).toNumber()] = _813_v0;
        _849_v33 = _nw153;
        let _index137 = _module.__default.safeIndex(new BigNumber(651), new BigNumber((_849_v33).length));
        (_849_v33)[_index137] = _813_v0;
        let _index138 = _module.__default.safeIndex(new BigNumber(651), new BigNumber((_849_v33).length));
        (_849_v33)[_index138] = _813_v0;
        let _852_v34;
        _852_v34 = _dafny.Set.fromElements((_this).f12);
        let _853_v35;
        _853_v35 = _dafny.MultiSet.fromElements((_this).f11);
        let _854_v36;
        _854_v36 = _dafny.Map.Empty.slice().updateUnsafe((_this).f12,new BigNumber((_853_v35).cardinality()));
        _852_v34 = (function () {
          let _coll18 = new _dafny.Set();
          for (const _compr_18 of (_854_v36).Keys.Elements) {
            let _855_v37 = _compr_18;
            if ((_854_v36).contains(_855_v37)) {
              _coll18.add(_module.__default.safeModuloInt(_855_v37, (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(365)))).cardinality()))));
            }
          }
          return _coll18;
        }()).Difference(_852_v34);
      }
      (globalState).f2 = (_this).f12;
      r0 = (_this).f11;
      let _856_v38;
      _856_v38 = _dafny.Seq.of(!((_this).f11), p0, (_this).f11);
      let _857_v39;
      _857_v39 = _dafny.Map.Empty.slice().updateUnsafe(_856_v38,p0);
      r1 = _857_v39;
      r2 = _dafny.Seq.IsPrefixOf(_856_v38, _856_v38);
      return [r0, r1, r2];
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
