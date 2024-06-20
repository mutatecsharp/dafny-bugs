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
    static fm4(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(true, false, false, false, false)));
    };
    static fm6(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(984)), function (_0_i0) {
        return new _dafny.CodePoint('v'.codePointAt(0));
      }), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("iwtniohu"), _dafny.Seq.UnicodeFromString("aa")));
    };
    static fm7(p0, globalState) {
      return new _dafny.CodePoint('l'.codePointAt(0));
    };
    static fm10(p0, p1, p2, p3, globalState) {
      return (_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber(-104)), new BigNumber(964))).Difference((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('f'.codePointAt(0)))).length), new BigNumber(57))).Intersect(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("rjys")).length))));
    };
    static fm11(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(907)).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-984),new BigNumber((_dafny.Seq.UnicodeFromString("bsuch")).length))).length)),false);
    };
    static fm12(p0, p1, p2, p3, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(false, false)).length)))).cardinality()),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-992)), function (_1_i0) {
        return new BigNumber(970);
      })).length)))).length))).length)), new BigNumber(782)),(new _dafny.CodePoint('u'.codePointAt(0))));
    };
    static fm13(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("frnc"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-34)), function (_2_i0) {
        return new _dafny.CodePoint('v'.codePointAt(0));
      }));
    };
    static fm14(globalState) {
      return new _dafny.CodePoint('j'.codePointAt(0));
    };
    static fm15(p0, p1, p2, p3, globalState) {
      return ((_dafny.MultiSet.fromElements(true)).Difference(_dafny.MultiSet.fromElements(false, true))).Union((_dafny.MultiSet.fromElements(false, true)).Intersect(_dafny.MultiSet.fromElements(true, false)));
    };
    static fm16(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_dafny.Set.fromElements(false), _dafny.Set.fromElements(false, true, false, true, false)), _dafny.Seq.of(_dafny.Set.fromElements(true), _dafny.Set.fromElements(!(!(false))), _dafny.Set.fromElements(false))), (_module.D34.create_DC84(_dafny.Seq.Create(_module.__default.abs(new BigNumber(129)), function (_3_i0) {
  return _dafny.Set.fromElements(true, true);
}))).dtor_cf145);
    };
    static fm17(p0, globalState) {
      return ((_dafny.Set.fromElements(new BigNumber(751), new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()))).Union(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length), new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(935),true),true)).length)))).length), new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality())))).Difference(_dafny.Set.fromElements(new BigNumber(82), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(316), new BigNumber(-425))) {
          let _4_v0 = _compr_0;
          if (((new BigNumber(316)).isLessThanOrEqualTo(_4_v0)) && ((_4_v0).isLessThan(new BigNumber(-425)))) {
            _coll0.push([(_4_v0).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(833),new _dafny.CodePoint('b'.codePointAt(0)))).length))),false]);
          }
        }
        return _coll0;
      }()).length),new BigNumber(874))).length))));
    };
    static fm20(p0, p1, p2, p3, globalState) {
      if (true) {
        return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("awte"), _dafny.Seq.UnicodeFromString("skgl"));
      } else {
        return _dafny.Seq.UnicodeFromString("yr");
      }
    };
    static fm21(globalState) {
      let _source0 = _module.D1.create_DC4(false, !(true), new BigNumber((_dafny.Seq.UnicodeFromString("jee")).length));
      if (_source0.is_DC3) {
        let _5___mcc_h0 = (_source0).cf5;
        let _6_cf5 = _5___mcc_h0;
        return _module.D1.create_DC2(_dafny.Seq.Create(_module.__default.abs(new BigNumber(126)), ((_7_cf5) => function (_8_i0) {
  return new BigNumber((_7_cf5).length);
})(_6_cf5)));
      } else if (_source0.is_DC4) {
        let _9___mcc_h1 = (_source0).cf6;
        let _10___mcc_h2 = (_source0).cf7;
        let _11___mcc_h3 = (_source0).cf8;
        let _12_cf8 = _11___mcc_h3;
        let _13_cf7 = _10___mcc_h2;
        let _14_cf6 = _9___mcc_h1;
        return _module.D1.create_DC2(_dafny.Seq.of(new BigNumber(246)));
      } else if (_source0.is_DC2) {
        let _15___mcc_h4 = (_source0).cf4;
        let _16_cf4 = _15___mcc_h4;
        return _module.D1.create_DC2(_16_cf4);
      } else {
        let _17___mcc_h5 = (_source0).cf9;
        let _18_cf9 = _17___mcc_h5;
        return _module.D1.create_DC2(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("gkgim")).length), new BigNumber(-329)));
      }
    };
    static fm22(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("fihalg"), _dafny.Seq.UnicodeFromString("cileu")), _dafny.Seq.UnicodeFromString("urjn"));
    };
    static fm23(globalState) {
      if (false) {
        return _module.D0.create_DC0(!(true));
      } else {
        return _module.D0.create_DC0(false);
      }
    };
    static fm24(p0, globalState) {
      return (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,true), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(false,false))).Intersect((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,true))).Intersect(function () {
        let _coll1 = new _dafny.Set();
        for (const _compr_1 of (_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,true))).Elements) {
          let _19_v0 = _compr_1;
          if ((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,true))).contains(_19_v0)) {
            _coll1.add(_19_v0);
          }
        }
        return _coll1;
      }()));
    };
    static fm25(p0, p1, globalState) {
      return (_dafny.Set.fromElements(new BigNumber(510))).Intersect((_dafny.Set.fromElements(new BigNumber(239))).Union(_dafny.Set.fromElements(new BigNumber(179))));
    };
    static fm26(p0, globalState) {
      let _source1 = ((true) ? (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("alriycn")).length),new BigNumber(463))) : (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(896)), function (_20_i0) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      })).length),new BigNumber(486))));
      let _21___mcc_h0 = _source1;
      let _22_cf45 = _21___mcc_h0;
      return new _dafny.CodePoint('c'.codePointAt(0));
    };
    static fm29(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.MultiSet.fromElements(new _dafny.CodePoint('o'.codePointAt(0))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.MultiSet.fromElements(new _dafny.CodePoint('v'.codePointAt(0)))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.MultiSet.fromElements(new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('s'.codePointAt(0)))));
    };
    static fm30(globalState) {
      return _dafny.Set.fromElements(true);
    };
    static fm31(p0, p1, globalState) {
      return (_module.D6.create_DC18(new _dafny.CodePoint('l'.codePointAt(0)), new BigNumber(486), new BigNumber((_dafny.Set.fromElements(false)).length))).dtor_cf34;
    };
    static fm33(p0, p1, p2, p3, globalState) {
      let _source2 = _module.D6.create_DC18(new _dafny.CodePoint('p'.codePointAt(0)), new BigNumber((_dafny.Set.fromElements(new BigNumber(-530))).length), new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()));
      if (_source2.is_DC17) {
        let _23___mcc_h0 = (_source2).cf29;
        let _24___mcc_h1 = (_source2).cf30;
        let _25___mcc_h2 = (_source2).cf31;
        let _26___mcc_h3 = (_source2).cf32;
        let _27___mcc_h4 = (_source2).cf33;
        let _28_cf33 = _27___mcc_h4;
        let _29_cf32 = _26___mcc_h3;
        let _30_cf31 = _25___mcc_h2;
        let _31_cf30 = _24___mcc_h1;
        let _32_cf29 = _23___mcc_h0;
        return _30_cf31;
      } else if (_source2.is_DC18) {
        let _33___mcc_h5 = (_source2).cf34;
        let _34___mcc_h6 = (_source2).cf35;
        let _35___mcc_h7 = (_source2).cf36;
        let _36_cf36 = _35___mcc_h7;
        let _37_cf35 = _34___mcc_h6;
        let _38_cf34 = _33___mcc_h5;
        return _37_cf35;
      } else if (_source2.is_DC19) {
        let _39___mcc_h8 = (_source2).cf37;
        let _40___mcc_h9 = (_source2).cf38;
        let _41_cf38 = _40___mcc_h9;
        let _42_cf37 = _39___mcc_h8;
        return _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.of(function () {
          let _coll2 = new _dafny.Map();
          for (const _compr_2 of _dafny.IntegerRange(new BigNumber(-815), new BigNumber(685))) {
            let _43_v0 = _compr_2;
            if (((new BigNumber(-815)).isLessThanOrEqualTo(_43_v0)) && ((_43_v0).isLessThan(new BigNumber(685)))) {
              _coll2.push([_module.__default.safeModuloInt(_43_v0, new BigNumber(371)),(_module.D4.create_DC11(new BigNumber((_dafny.Seq.UnicodeFromString("fv")).length), _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(847),_41_cf38)).length)), _42_cf37)).dtor_cf18]);
            }
          }
          return _coll2;
        }())).length), new BigNumber(618));
      } else if (_source2.is_DC16) {
        let _44___mcc_h10 = (_source2).cf28;
        let _45_cf28 = _44___mcc_h10;
        return (_dafny.ZERO).minus(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.UnicodeFromString("dgejc"))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(true),_dafny.Seq.UnicodeFromString("vben")))).length));
      } else {
        let _46___mcc_h11 = (_source2).cf39;
        let _47_cf39 = _46___mcc_h11;
        return new BigNumber(408);
      }
    };
    static fm34(p0, p1, p2, p3, globalState) {
      return _dafny.Set.fromElements(((true) ? (false) : (false)));
    };
    static fm35(p0, globalState) {
      let _source3 = _module.D14.create_DC34(function () {
  let _coll3 = new _dafny.Map();
  for (const _compr_3 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(281)), function (_48_i0) {
    return _module.D4.create_DC11(new BigNumber(-429), _dafny.Set.fromElements(new BigNumber(694), new BigNumber(859)), true);
  }),new _dafny.CodePoint('q'.codePointAt(0)))).Keys.Elements) {
    let _49_v0 = _compr_3;
    if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(281)), function (_48_i0) {
      return _module.D4.create_DC11(new BigNumber(-429), _dafny.Set.fromElements(new BigNumber(694), new BigNumber(859)), true);
    }),new _dafny.CodePoint('q'.codePointAt(0)))).contains(_49_v0)) {
      _coll3.push([_49_v0,new BigNumber(-622)]);
    }
  }
  return _coll3;
}());
      if (_source3.is_DC35) {
        let _50___mcc_h0 = (_source3).cf62;
        let _51___mcc_h1 = (_source3).cf63;
        let _52___mcc_h2 = (_source3).cf64;
        let _53_cf64 = _52___mcc_h2;
        let _54_cf63 = _51___mcc_h1;
        let _55_cf62 = _50___mcc_h0;
        return _module.D3.create_DC8(_54_cf63, _54_cf63);
      } else if (_source3.is_DC36) {
        let _56___mcc_h3 = (_source3).cf65;
        let _57___mcc_h4 = (_source3).cf66;
        let _58___mcc_h5 = (_source3).cf67;
        let _59_cf67 = _58___mcc_h5;
        let _60_cf66 = _57___mcc_h4;
        let _61_cf65 = _56___mcc_h3;
        return _module.D3.create_DC8(!(_61_cf65), _61_cf65);
      } else {
        let _62___mcc_h6 = (_source3).cf61;
        let _63_cf61 = _62___mcc_h6;
        return _module.D3.create_DC8(!(false), !(true));
      }
    };
    static fm36(globalState) {
      let _source4 = _module.D6.create_DC20(_module.D6.create_DC18(new _dafny.CodePoint('m'.codePointAt(0)), new BigNumber(779), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(555))).length)));
      if (_source4.is_DC17) {
        let _64___mcc_h0 = (_source4).cf29;
        let _65___mcc_h1 = (_source4).cf30;
        let _66___mcc_h2 = (_source4).cf31;
        let _67___mcc_h3 = (_source4).cf32;
        let _68___mcc_h4 = (_source4).cf33;
        let _69_cf33 = _68___mcc_h4;
        let _70_cf32 = _67___mcc_h3;
        let _71_cf31 = _66___mcc_h2;
        let _72_cf30 = _65___mcc_h1;
        let _73_cf29 = _64___mcc_h0;
        return new _dafny.CodePoint('o'.codePointAt(0));
      } else if (_source4.is_DC18) {
        let _74___mcc_h5 = (_source4).cf34;
        let _75___mcc_h6 = (_source4).cf35;
        let _76___mcc_h7 = (_source4).cf36;
        let _77_cf36 = _76___mcc_h7;
        let _78_cf35 = _75___mcc_h6;
        let _79_cf34 = _74___mcc_h5;
        return new _dafny.CodePoint('x'.codePointAt(0));
      } else if (_source4.is_DC19) {
        let _80___mcc_h8 = (_source4).cf37;
        let _81___mcc_h9 = (_source4).cf38;
        let _82_cf38 = _81___mcc_h9;
        let _83_cf37 = _80___mcc_h8;
        return new _dafny.CodePoint('g'.codePointAt(0));
      } else if (_source4.is_DC16) {
        let _84___mcc_h10 = (_source4).cf28;
        let _85_cf28 = _84___mcc_h10;
        return new _dafny.CodePoint('s'.codePointAt(0));
      } else {
        let _86___mcc_h11 = (_source4).cf39;
        let _87_cf39 = _86___mcc_h11;
        return new _dafny.CodePoint('q'.codePointAt(0));
      }
    };
    static fm37(globalState) {
      return _dafny.Seq.Concat(((false) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(972)), function (_88_i0) {
        return new BigNumber(929);
      })) : (_dafny.Seq.of(new BigNumber(342), new BigNumber(642)))), _dafny.Seq.of(new BigNumber(208)));
    };
    static fm38(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(false, true, false, !(false)), _dafny.Seq.of(false, false, false));
    };
    static fm39(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(true, false, true, !(true), true)).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(335),!(true))).length)),((false) ? ((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(new BigNumber(-919))).length))) : (new BigNumber(247))));
    };
    static fm40(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(188),new BigNumber(-896)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(894),new BigNumber(511))), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(698)), function (_89_i0) {
        return function () {
          let _coll4 = new _dafny.Map();
          for (const _compr_4 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(631))).cardinality()),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(575)), function (_90_i1) {
            return new _dafny.CodePoint('v'.codePointAt(0));
          })).length))).Keys.Elements) {
            let _91_v0 = _compr_4;
            if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(631))).cardinality()),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(575)), function (_90_i1) {
              return new _dafny.CodePoint('v'.codePointAt(0));
            })).length))).contains(_91_v0)) {
              _coll4.push([(_91_v0).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,!(true))).length)),new BigNumber(696)]);
            }
          }
          return _coll4;
        }();
      }), _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(587),new BigNumber((_dafny.Seq.UnicodeFromString("cldmyuqk")).length)), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(false)).length),new BigNumber(179)))));
    };
    static fm43(p0, globalState) {
      if (true) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('b'.codePointAt(0));
      }
    };
    static fm44(p0, globalState) {
      return (function () {
        let _coll5 = new _dafny.Map();
        for (const _compr_5 of (_dafny.MultiSet.fromElements((_module.D17.create_DC44(new BigNumber(357), _dafny.Seq.of(false), new _dafny.CodePoint('t'.codePointAt(0)), true, new BigNumber(983))).dtor_cf84)).Elements) {
          let _92_v0 = _compr_5;
          if ((_dafny.MultiSet.fromElements((_module.D17.create_DC44(new BigNumber(357), _dafny.Seq.of(false), new _dafny.CodePoint('t'.codePointAt(0)), true, new BigNumber(983))).dtor_cf84)).contains(_92_v0)) {
            _coll5.push([(_92_v0).plus(new BigNumber(-17)),!(true)]);
          }
        }
        return _coll5;
      }()).Merge((_module.D35.create_DC87(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_module.D15.create_DC38(false, _dafny.Seq.Create(_module.__default.abs(new BigNumber(331)), function (_93_i0) {
  return new _dafny.CodePoint('x'.codePointAt(0));
}), true, !(true), new _dafny.CodePoint('i'.codePointAt(0)))).dtor_cf70).length),true))).dtor_cf149);
    };
    static fm45(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("m"), _dafny.Seq.UnicodeFromString("wgb")), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("lkwgqmac"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(878)), function (_94_i0) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      })));
    };
    static fm46(p0, p1, p2, globalState) {
      return (_dafny.MultiSet.fromElements(new BigNumber(347), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ye")).length)))).Intersect((_dafny.MultiSet.fromElements(new BigNumber(-658))).Difference(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(false)).length), new BigNumber(42))));
    };
    static fm47(p0, globalState) {
      if (!(true)) {
        return _dafny.Seq.of(_dafny.Seq.UnicodeFromString("bvpnsv"), _dafny.Seq.UnicodeFromString("f"));
      } else {
        return _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(420)), function (_95_i0) {
          return new _dafny.CodePoint('p'.codePointAt(0));
        }), _dafny.Seq.UnicodeFromString("ch"));
      }
    };
    static fm48(p0, globalState) {
      if ((false) || (true)) {
        return (_dafny.Map.Empty.slice().updateUnsafe(false,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(true),false));
      } else {
        return (_dafny.Map.Empty.slice().updateUnsafe(true,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(!(true)),true));
      }
    };
    static fm49(p0, globalState) {
      return _module.D6.create_DC17(true, (_dafny.MultiSet.fromElements(new BigNumber(-731))).IsSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((function () {
  let _coll6 = new _dafny.Map();
  for (const _compr_6 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(822)), function (_96_i0) {
    return new _dafny.CodePoint('d'.codePointAt(0));
  }),new BigNumber((function () {
    let _coll7 = new _dafny.Map();
    for (const _compr_7 of (_dafny.Seq.of(_dafny.Seq.of(new _dafny.CodePoint('r'.codePointAt(0))))).Elements) {
      let _97_v1 = _compr_7;
      if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.of(new _dafny.CodePoint('r'.codePointAt(0)))), _97_v1)) {
        _coll7.push([_97_v1,false]);
      }
    }
    return _coll7;
  }()).length))).Keys.Elements) {
    let _98_v0 = _compr_6;
    if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(822)), function (_96_i0) {
      return new _dafny.CodePoint('d'.codePointAt(0));
    }),new BigNumber((function () {
      let _coll8 = new _dafny.Map();
      for (const _compr_8 of (_dafny.Seq.of(_dafny.Seq.of(new _dafny.CodePoint('r'.codePointAt(0))))).Elements) {
        let _97_v1 = _compr_8;
        if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.of(new _dafny.CodePoint('r'.codePointAt(0)))), _97_v1)) {
          _coll8.push([_97_v1,false]);
        }
      }
      return _coll8;
    }()).length))).contains(_98_v0)) {
      _coll6.push([_98_v0,new BigNumber(-740)]);
    }
  }
  return _coll6;
}()).length)))), new BigNumber(330), (new BigNumber(3)).multipliedBy(new BigNumber(285)), _module.__default.safeDivisionInt(new BigNumber(-341), new BigNumber((_dafny.Seq.UnicodeFromString("sn")).length)));
    };
    static fm50(p0, p1, globalState) {
      return (_dafny.MultiSet.fromElements(true)).Difference(_dafny.MultiSet.fromElements(!(false)));
    };
    static fm51(globalState) {
      if (_dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(106)), function (_99_i0) {
        return new BigNumber(-181);
      }), _dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements((new _dafny.CodePoint('d'.codePointAt(0))))).cardinality()))).length), new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(516)), function (_100_i1) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      }))).cardinality()))).length)))) {
        return _module.D14.create_DC34(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-130)), function (_101_i2) {
  return _module.D4.create_DC11(new BigNumber(-556), _dafny.Set.fromElements(new BigNumber(253)), true);
}),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(5)), function (_102_i3) {
  return new _dafny.CodePoint('e'.codePointAt(0));
})).length)));
      } else {
        return _module.D14.create_DC34(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_module.D4.create_DC11(new BigNumber(636), function () {
  let _coll9 = new _dafny.Set();
  for (const _compr_9 of _dafny.IntegerRange(new BigNumber(712), new BigNumber(679))) {
    let _103_v0 = _compr_9;
    if (((new BigNumber(712)).isLessThanOrEqualTo(_103_v0)) && ((_103_v0).isLessThan(new BigNumber(679)))) {
      _coll9.add((_103_v0).minus(new BigNumber((_dafny.Seq.UnicodeFromString("iqsvf")).length)));
    }
  }
  return _coll9;
}(), true)),new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false, true, true, !(true))).length))).length)));
      }
    };
    static fm52(p0, globalState) {
      if (!((new BigNumber((function () {
        let _coll10 = new _dafny.Map();
        for (const _compr_10 of (function () {
          let _coll11 = new _dafny.Map();
          for (const _compr_11 of (function () {
            let _coll12 = new _dafny.Map();
            for (const _compr_12 of _dafny.IntegerRange(new BigNumber(419), new BigNumber(579))) {
              let _104_v2 = _compr_12;
              if (((new BigNumber(419)).isLessThanOrEqualTo(_104_v2)) && ((_104_v2).isLessThan(new BigNumber(579)))) {
                _coll12.push([_module.__default.safeDivisionInt(_104_v2, new BigNumber((_dafny.Seq.UnicodeFromString("sfqi")).length)),new BigNumber(971)]);
              }
            }
            return _coll12;
          }()).Keys.Elements) {
            let _105_v1 = _compr_11;
            if ((function () {
              let _coll13 = new _dafny.Map();
              for (const _compr_13 of _dafny.IntegerRange(new BigNumber(419), new BigNumber(579))) {
                let _104_v2 = _compr_13;
                if (((new BigNumber(419)).isLessThanOrEqualTo(_104_v2)) && ((_104_v2).isLessThan(new BigNumber(579)))) {
                  _coll13.push([_module.__default.safeDivisionInt(_104_v2, new BigNumber((_dafny.Seq.UnicodeFromString("sfqi")).length)),new BigNumber(971)]);
                }
              }
              return _coll13;
            }()).contains(_105_v1)) {
              _coll11.push([_module.__default.safeModuloInt(_105_v1, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(689)), function (_106_i0) {
                return new _dafny.CodePoint('f'.codePointAt(0));
              })).length)),true]);
            }
          }
          return _coll11;
        }()).Keys.Elements) {
          let _107_v0 = _compr_10;
          if ((function () {
            let _coll14 = new _dafny.Map();
            for (const _compr_14 of (function () {
              let _coll15 = new _dafny.Map();
              for (const _compr_15 of _dafny.IntegerRange(new BigNumber(419), new BigNumber(579))) {
                let _104_v2 = _compr_15;
                if (((new BigNumber(419)).isLessThanOrEqualTo(_104_v2)) && ((_104_v2).isLessThan(new BigNumber(579)))) {
                  _coll15.push([_module.__default.safeDivisionInt(_104_v2, new BigNumber((_dafny.Seq.UnicodeFromString("sfqi")).length)),new BigNumber(971)]);
                }
              }
              return _coll15;
            }()).Keys.Elements) {
              let _105_v1 = _compr_14;
              if ((function () {
                let _coll16 = new _dafny.Map();
                for (const _compr_16 of _dafny.IntegerRange(new BigNumber(419), new BigNumber(579))) {
                  let _104_v2 = _compr_16;
                  if (((new BigNumber(419)).isLessThanOrEqualTo(_104_v2)) && ((_104_v2).isLessThan(new BigNumber(579)))) {
                    _coll16.push([_module.__default.safeDivisionInt(_104_v2, new BigNumber((_dafny.Seq.UnicodeFromString("sfqi")).length)),new BigNumber(971)]);
                  }
                }
                return _coll16;
              }()).contains(_105_v1)) {
                _coll14.push([_module.__default.safeModuloInt(_105_v1, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(689)), function (_106_i0) {
                  return new _dafny.CodePoint('f'.codePointAt(0));
                })).length)),true]);
              }
            }
            return _coll14;
          }()).contains(_107_v0)) {
            _coll10.push([_module.__default.safeModuloInt(_107_v0, new BigNumber((_dafny.MultiSet.fromElements(!(false))).cardinality())),new BigNumber(-334)]);
          }
        }
        return _coll10;
      }()).length)).isLessThan(new BigNumber((_dafny.Seq.of(true)).length)))) {
        return _module.D8.create_DC23(_dafny.Seq.of(true));
      } else {
        return _module.D8.create_DC23(_dafny.Seq.of(!(!(true))));
      }
    };
    static fm53(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(((!(!(false))) ? (!(true)) : (false)),((false) ? (_module.D4.create_DC11(new BigNumber(859), function () {
  let _coll17 = new _dafny.Set();
  for (const _compr_17 of _dafny.IntegerRange(new BigNumber(845), new BigNumber(623))) {
    let _108_v0 = _compr_17;
    if (((new BigNumber(845)).isLessThanOrEqualTo(_108_v0)) && ((_108_v0).isLessThan(new BigNumber(623)))) {
      _coll17.add(_module.__default.safeDivisionInt(_108_v0, new BigNumber(-722)));
    }
  }
  return _coll17;
}(), false)) : (_module.D4.create_DC11(new BigNumber((_dafny.Set.fromElements(!(false), true, !(!(!(false))))).length), function () {
  let _coll18 = new _dafny.Set();
  for (const _compr_18 of _dafny.IntegerRange(new BigNumber(-473), new BigNumber(423))) {
    let _109_v1 = _compr_18;
    if (((new BigNumber(-473)).isLessThanOrEqualTo(_109_v1)) && ((_109_v1).isLessThan(new BigNumber(423)))) {
      _coll18.add((_109_v1).plus(new BigNumber(470)));
    }
  }
  return _coll18;
}(), !(true)))));
    };
    static fm54(globalState) {
      return function () {
        let _coll19 = new _dafny.Set();
        for (const _compr_19 of (_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(748)), function (_110_i0) {
          return new _dafny.CodePoint('e'.codePointAt(0));
        }))).Elements) {
          let _111_v0 = _compr_19;
          if ((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(748)), function (_110_i0) {
            return new _dafny.CodePoint('e'.codePointAt(0));
          }))).contains(_111_v0)) {
            _coll19.add(_111_v0);
          }
        }
        return _coll19;
      }();
    };
    static fm55(p0, globalState) {
      return _module.D6.create_DC20(((true) ? (_module.D6.create_DC19(!(false), new _dafny.CodePoint('i'.codePointAt(0)))) : (_module.D6.create_DC20(_module.D6.create_DC20(_module.D6.create_DC19(false, new _dafny.CodePoint('q'.codePointAt(0))))))));
    };
    static fm56(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-87)), function (_112_i0) {
        return _module.D14.create_DC36(false, _dafny.Seq.UnicodeFromString("mpbldi"), _dafny.Seq.UnicodeFromString("dccddbi"));
      }), _dafny.Seq.of(_module.D14.create_DC36(false, _dafny.Seq.Create(_module.__default.abs(new BigNumber(108)), function (_113_i1) {
  return new _dafny.CodePoint('i'.codePointAt(0));
}), _dafny.Seq.UnicodeFromString("xdxoru")), _module.D14.create_DC36(true, _dafny.Seq.UnicodeFromString("igbltr"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(982)), function (_114_i2) {
  return new _dafny.CodePoint('n'.codePointAt(0));
})))), _dafny.Seq.of(_module.D14.create_DC36(!(!(true)), _dafny.Seq.UnicodeFromString("d"), _dafny.Seq.UnicodeFromString("kwqsarg"))));
    };
    static fm57(p0, globalState) {
      return ((_dafny.MultiSet.fromElements(new _dafny.CodePoint('g'.codePointAt(0)))).Difference(_dafny.MultiSet.fromElements(new _dafny.CodePoint('s'.codePointAt(0))))).Difference((_dafny.MultiSet.fromElements(new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('q'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0)), new _dafny.CodePoint('m'.codePointAt(0)))).Union(_dafny.MultiSet.fromElements(new _dafny.CodePoint('i'.codePointAt(0)), new _dafny.CodePoint('r'.codePointAt(0)))));
    };
    static fm58(globalState) {
      return _dafny.Seq.of((_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(true)).length))).Union(_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()))), function () {
        let _coll20 = new _dafny.Set();
        for (const _compr_20 of _dafny.IntegerRange(new BigNumber(630), new BigNumber(790))) {
          let _115_v0 = _compr_20;
          if (((new BigNumber(630)).isLessThanOrEqualTo(_115_v0)) && ((_115_v0).isLessThan(new BigNumber(790)))) {
            _coll20.add((_115_v0).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("mnhgh"))).length)));
          }
        }
        return _coll20;
      }(), _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber((_dafny.Set.fromElements(false, false)).length))).length)));
    };
    static fm59(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),_dafny.Seq.of(!(true), true, !(true)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('b'.codePointAt(0)),(_module.D17.create_DC44((_module.D34.create_DC86(true, new BigNumber(906))).dtor_cf148, _dafny.Seq.of(true), new _dafny.CodePoint('p'.codePointAt(0)), !(true), new BigNumber((_dafny.Set.fromElements(false)).length))).dtor_cf85))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('h'.codePointAt(0)),_dafny.Seq.of(!(!(false)), true)));
    };
    static fm60(p0, p1, globalState) {
      let _source5 = _dafny.Seq.of(_module.D4.create_DC11(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), _dafny.Set.fromElements(new BigNumber(943), new BigNumber(-813), new BigNumber(-314)), !(false)));
      let _116___mcc_h0 = _source5;
      let _117_cf60 = _116___mcc_h0;
      return _module.D15.create_DC39(true, true);
    };
    static fm61(p0, globalState) {
      let _source6 = _module.D34.create_DC85(new _dafny.CodePoint('p'.codePointAt(0)));
      if (_source6.is_DC85) {
        let _118___mcc_h0 = (_source6).cf146;
        let _119_cf146 = _118___mcc_h0;
        return _dafny.MultiSet.fromElements(new BigNumber(857));
      } else if (_source6.is_DC86) {
        let _120___mcc_h1 = (_source6).cf147;
        let _121___mcc_h2 = (_source6).cf148;
        let _122_cf148 = _121___mcc_h2;
        let _123_cf147 = _120___mcc_h1;
        if (_123_cf147) {
          return _dafny.MultiSet.FromArray(_dafny.Seq.of(_122_cf148, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(214)), function (_124_i0) {
            return _dafny.Seq.UnicodeFromString("ey");
          })).length)));
        } else {
          return _dafny.MultiSet.fromElements(_122_cf148, _122_cf148);
        }
      } else {
        let _125___mcc_h3 = (_source6).cf145;
        let _126_cf145 = _125___mcc_h3;
        return _dafny.MultiSet.fromElements(new BigNumber(-923), new BigNumber((_dafny.Set.fromElements(true)).length), new BigNumber((_dafny.Seq.UnicodeFromString("qgxceec")).length), new BigNumber(738));
      }
    };
    static Main(__noArgsParameter) {
      let _127_v0;
      let _init0 = function (_128_i0) {
        return false;
      };
      let _nw0 = Array((new BigNumber(18)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw0.length); _i0_0++) {
        _nw0[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _127_v0 = _nw0;
      let _129_v1;
      _129_v1 = true;
      let _130_v2;
      _130_v2 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(4),_129_v1);
      let _131_globalState;
      let _nw1 = new _module.GlobalState();
      _nw1.__ctor(false, false, _127_v0, true, _dafny.Map.Empty.slice().updateUnsafe(_130_v2,!(_129_v1)), new BigNumber(-523), false, new BigNumber(-26), false);
      _131_globalState = _nw1;
      let _132_v3;
      let _nw2 = new _module.C1();
      _nw2.__ctor(_129_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(880)), function (_133_i1) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      }));
      _132_v3 = _nw2;
      let _134_v4;
      let _nw3 = Array((new BigNumber(10)).toNumber()).fill([]);
      _134_v4 = _nw3;
      let _135_v5;
      _135_v5 = _module.D17.create_DC43(_134_v4);
      let _136_v6;
      _136_v6 = _dafny.MultiSet.fromElements(_135_v5);
      let _137_v7;
      _137_v7 = _dafny.Map.Empty.slice().updateUnsafe(_136_v6,new BigNumber((_dafny.Seq.of(_129_v1)).length));
      let _138_v8;
      _138_v8 = new BigNumber(925);
      if ((_132_v3).fm32((((_137_v7).contains((_136_v6).update(_135_v5, _module.__default.abs(new BigNumber((_dafny.Seq.UnicodeFromString("rdlthb")).length))))) ? ((_137_v7).get((_136_v6).update(_135_v5, _module.__default.abs(new BigNumber((_dafny.Seq.UnicodeFromString("rdlthb")).length))))) : (_138_v8)), ((_129_v1) ? (new _dafny.CodePoint('j'.codePointAt(0))) : (new _dafny.CodePoint('l'.codePointAt(0)))), _131_globalState)) {
        let _139_v9;
        let _nw4 = new _module.C0();
        _nw4.__ctor();
        _139_v9 = _nw4;
        let _140_v10;
        _140_v10 = new _dafny.CodePoint('v'.codePointAt(0));
        _138_v8 = new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("pc"), _module.__default.safeIndex(new BigNumber(((_132_v3).f13).length), new BigNumber((_dafny.Seq.UnicodeFromString("pc")).length)), _140_v10)).length);
        (_131_globalState).f8 = !(!(_138_v8).isEqualTo(_module.__default.safeDivisionInt(_138_v8, new BigNumber(67))));
        let _141_v11;
        _141_v11 = _dafny.Map.Empty.slice().updateUnsafe(!((((_132_v3).f12) ? ((_132_v3).f12) : ((_132_v3).f12))),_138_v8);
        let _142_v12;
        _142_v12 = _dafny.Seq.of(_129_v1, (_132_v3).f12);
        let _143_v13;
        _143_v13 = _dafny.Set.fromElements(_138_v8, new BigNumber(521));
        _141_v11 = (_141_v11).update((_132_v3).f12, new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_140_v10,_142_v12)).Merge(_module.__default.fm59(new BigNumber((_143_v13).length), _138_v8, _131_globalState))).length));
        let _144_v14;
        _144_v14 = _module.D12.create_DC31((_132_v3).f12, (_132_v3).f13, (_132_v3).f13);
        let _145_v15;
        _145_v15 = _dafny.Map.Empty.slice().updateUnsafe(_138_v8,(_132_v3).f13);
        let _146_v16;
        _146_v16 = _module.D28.create_DC67(_module.__default.fm58(_131_globalState));
        let _147_v17;
        let _nw5 = Array((new BigNumber(20)).toNumber());
        _nw5[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat((_132_v3).f13, (_132_v3).f13);
        _nw5[(_dafny.ONE).toNumber()] = (_132_v3).f13;
        _nw5[(new BigNumber(2)).toNumber()] = (_132_v3).f13;
        _nw5[(new BigNumber(3)).toNumber()] = (_132_v3).f13;
        _nw5[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat((_132_v3).f13, (_132_v3).f13);
        _nw5[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat((_132_v3).f13, (_132_v3).f13);
        _nw5[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat((_132_v3).f13, (_144_v14).dtor_cf56);
        _nw5[(new BigNumber(7)).toNumber()] = (_132_v3).f13;
        _nw5[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat((_132_v3).f13, (_132_v3).f13);
        _nw5[(new BigNumber(9)).toNumber()] = _module.__default.fm22(_138_v8, (_132_v3).f12, false, _138_v8, _131_globalState);
        _nw5[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update((_132_v3).f13, _module.__default.safeIndex(_138_v8, new BigNumber(((_132_v3).f13).length)), _140_v10), _dafny.Seq.UnicodeFromString("bmidbbbbs"));
        _nw5[(new BigNumber(11)).toNumber()] = (_132_v3).f13;
        _nw5[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat((((_145_v15).contains((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_140_v10,(_module.__default.fm60(_146_v16, (_132_v3).f12, _131_globalState)).dtor_cf75)).length)))) ? ((_145_v15).get((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_140_v10,(_module.__default.fm60(_146_v16, (_132_v3).f12, _131_globalState)).dtor_cf75)).length)))) : (_dafny.Seq.UnicodeFromString("etfudx"))), (_132_v3).f13);
        _nw5[(new BigNumber(13)).toNumber()] = (_132_v3).f13;
        _nw5[(new BigNumber(14)).toNumber()] = (_132_v3).f13;
        _nw5[(new BigNumber(15)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(13)), ((_148_v10) => function (_149_i2) {
          return _148_v10;
        })(_140_v10));
        _nw5[(new BigNumber(16)).toNumber()] = _module.__default.fm13(_138_v8, _131_globalState);
        _nw5[(new BigNumber(17)).toNumber()] = _dafny.Seq.UnicodeFromString("ieoynisy");
        _nw5[(new BigNumber(18)).toNumber()] = (_132_v3).f13;
        _nw5[(new BigNumber(19)).toNumber()] = _module.__default.fm45((_132_v3).f12, _138_v8, (_132_v3).f12, _131_globalState);
        _147_v17 = _nw5;
        _147_v17 = _147_v17;
      } else {
        let _150_v18;
        let _nw6 = new _module.C0();
        _nw6.__ctor();
        _150_v18 = _nw6;
        let _151_v19;
        _151_v19 = _dafny.Map.Empty.slice().updateUnsafe(false,_138_v8);
        let _152_v20;
        _152_v20 = _dafny.Map.Empty.slice().updateUnsafe((_132_v3).f12,_129_v1);
        _151_v19 = (_151_v19).update(_129_v1, _module.__default.safeDivisionInt(_138_v8, (_dafny.ZERO).minus(new BigNumber((_152_v20).length))));
        if ((_132_v3).f12) {
          _138_v8 = (_138_v8).minus(_module.__default.safeModuloInt(_138_v8, _module.__default.fm33((_132_v3).f12, (_132_v3).f12, _138_v8, (_132_v3).f12, _131_globalState)));
          let _153_v21;
          _153_v21 = _dafny.MultiSet.fromElements((_132_v3).f12);
          let _154_v22;
          _154_v22 = _dafny.Map.Empty.slice().updateUnsafe(((_129_v1) ? ((_132_v3).f12) : (_129_v1)),_153_v21);
          _154_v22 = (_154_v22).update((_138_v8).isLessThan(_138_v8), _153_v21);
          let _155_v23;
          _155_v23 = _dafny.Seq.UnicodeFromString("ynqtib");
          let _156_v24;
          _156_v24 = _dafny.Set.fromElements(!((_132_v3).f12));
          let _157_v25;
          _157_v25 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm44((_132_v3).f12, _131_globalState),true);
          let _rhs0 = new BigNumber(((_156_v24).Difference(_module.__default.fm34(_138_v8, (_132_v3).f12, _129_v1, (((_157_v25).contains(_130_v2)) ? ((_157_v25).get(_130_v2)) : ((_132_v3).f12)), _131_globalState))).length);
          let _rhs1 = _dafny.Seq.Concat((((_132_v3).f12) ? ((_132_v3).f13) : (_dafny.Seq.UnicodeFromString("i"))), _dafny.Seq.Concat((_132_v3).f13, (_132_v3).f13));
          _138_v8 = _rhs0;
          _155_v23 = _rhs1;
          let _158_v26;
          let _nw7 = new _module.C5();
          _nw7.__ctor(_152_v20, (_132_v3).f12);
          _158_v26 = _nw7;
          _158_v26 = _158_v26;
          let _159_v27;
          _159_v27 = new _dafny.CodePoint('f'.codePointAt(0));
          _155_v23 = _dafny.Seq.update((_132_v3).f13, _module.__default.safeIndex(_138_v8, new BigNumber(((_132_v3).f13).length)), _159_v27);
        } else {
          _138_v8 = new BigNumber(-455);
          let _160_v28;
          let _nw8 = new _module.C3();
          _nw8.__ctor();
          _160_v28 = _nw8;
          let _161_v29;
          _161_v29 = new _dafny.CodePoint('o'.codePointAt(0));
          let _rhs2 = !(!((_138_v8).plus(_138_v8)).isEqualTo(_module.__default.safeModuloInt(_138_v8, _138_v8)));
          let _rhs3 = ((!((_132_v3).fm32(_138_v8, _161_v29, _131_globalState))) ? (_129_v1) : (_129_v1));
          let _rhs4 = _160_v28;
          let _rhs5 = _module.__default.safeDivisionInt(_138_v8, _138_v8);
          let _lhs0 = _131_globalState;
          _lhs0.f8 = _rhs2;
          _129_v1 = _rhs3;
          _160_v28 = _rhs4;
          _138_v8 = _rhs5;
          _138_v8 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_138_v8, _138_v8));
          _138_v8 = _module.__default.fm33((_132_v3).f12, (_132_v3).f12, _module.__default.fm33((_132_v3).f12, _129_v1, new BigNumber(590), (_132_v3).f12, _131_globalState), _129_v1, _131_globalState);
          _138_v8 = _138_v8;
        }
        let _162_v30;
        let _nw9 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
        _162_v30 = _nw9;
        let _index0 = _module.__default.safeIndex(new BigNumber(402), new BigNumber((_162_v30).length));
        (_162_v30)[_index0] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("mlch"), (_132_v3).f13)).length);
        let _index1 = _module.__default.safeIndex(new BigNumber(402), new BigNumber((_162_v30).length));
        (_162_v30)[_index1] = (_138_v8).minus(_module.__default.fm33(_129_v1, (_132_v3).f12, (_dafny.ZERO).minus(_138_v8), !(true), _131_globalState));
        let _163_v32;
        let _init1 = ((_164_v3) => function (_165_i3) {
          return function () {
            let _coll21 = new _dafny.Map();
            for (const _compr_21 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(271)), function (_166_i4) {
              return _dafny.Seq.Create(_module.__default.abs(new BigNumber(610)), function (_167_i5) {
                return new _dafny.CodePoint('s'.codePointAt(0));
              });
            })).Elements) {
              let _168_v31 = _compr_21;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(271)), function (_166_i4) {
                return _dafny.Seq.Create(_module.__default.abs(new BigNumber(610)), function (_167_i5) {
                  return new _dafny.CodePoint('s'.codePointAt(0));
                });
              }), _168_v31)) {
                _coll21.push([_168_v31,new BigNumber(((_164_v3).f13).length)]);
              }
            }
            return _coll21;
          }();
        })(_132_v3);
        let _nw10 = Array((new BigNumber(7)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw10.length); _i0_1++) {
          _nw10[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _163_v32 = _nw10;
        let _169_v33;
        _169_v33 = _dafny.Map.Empty.slice().updateUnsafe((_132_v3).f13,(_162_v30)[_module.__default.safeIndex(new BigNumber(402), new BigNumber((_162_v30).length))]);
        let _170_v34;
        _170_v34 = new _dafny.CodePoint('p'.codePointAt(0));
        let _index2 = _module.__default.safeIndex(new BigNumber(362), new BigNumber((_163_v32).length));
        (_163_v32)[_index2] = ((_169_v33).update(_dafny.Seq.of(_170_v34, _170_v34), new BigNumber((_dafny.MultiSet.fromElements(_138_v8)).cardinality()))).update((_132_v3).f13, _138_v8);
        let _171_v35;
        _171_v35 = _dafny.Map.Empty.slice().updateUnsafe(_162_v30,_138_v8);
        let _172_v36;
        _172_v36 = _dafny.MultiSet.fromElements((_132_v3).f12);
        let _index3 = _module.__default.safeIndex(new BigNumber(362), new BigNumber((_163_v32).length));
        let _rhs6 = ((((_171_v35)).equals(_171_v35)) ? (_138_v8) : (new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_172_v36).cardinality()))).cardinality())));
        let _rhs7 = ((_169_v33).update((_132_v3).f13, (_162_v30)[_module.__default.safeIndex(new BigNumber(402), new BigNumber((_162_v30).length))])).Merge((_dafny.Map.Empty.slice().updateUnsafe((_132_v3).f13,(_162_v30)[_module.__default.safeIndex(new BigNumber(402), new BigNumber((_162_v30).length))])).Merge((_169_v33).update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(697)), ((_173_v34) => function (_174_i6) {
          return _173_v34;
        })(_170_v34)), (_162_v30)[_module.__default.safeIndex(new BigNumber(402), new BigNumber((_162_v30).length))])));
        let _rhs8 = _170_v34;
        let _lhs1 = _163_v32;
        let _lhs2 = _module.__default.safeIndex(new BigNumber(362), new BigNumber((_163_v32).length));
        _138_v8 = _rhs6;
        _lhs1[_lhs2] = _rhs7;
        _170_v34 = _rhs8;
      }
      let _175_v37;
      _175_v37 = new _dafny.CodePoint('q'.codePointAt(0));
      _175_v37 = _175_v37;
      let _176_v38;
      _176_v38 = _dafny.Map.Empty.slice().updateUnsafe((_132_v3).f12,_129_v1);
      let _177_v39;
      let _nw11 = new _module.C6();
      _nw11.__ctor(_176_v38, true);
      _177_v39 = _nw11;
      let _178_v40;
      _178_v40 = _module.D22.create_DC56(_177_v39);
      let _179_v41;
      let _nw12 = Array((new BigNumber(26)).toNumber());
      _nw12[(_dafny.ZERO).toNumber()] = _177_v39;
      _nw12[(_dafny.ONE).toNumber()] = _177_v39;
      _nw12[(new BigNumber(2)).toNumber()] = _177_v39;
      _nw12[(new BigNumber(3)).toNumber()] = _177_v39;
      _nw12[(new BigNumber(4)).toNumber()] = _177_v39;
      _nw12[(new BigNumber(5)).toNumber()] = _177_v39;
      _nw12[(new BigNumber(6)).toNumber()] = _177_v39;
      _nw12[(new BigNumber(7)).toNumber()] = _177_v39;
      _nw12[(new BigNumber(8)).toNumber()] = _177_v39;
      _nw12[(new BigNumber(9)).toNumber()] = _177_v39;
      _nw12[(new BigNumber(10)).toNumber()] = _177_v39;
      _nw12[(new BigNumber(11)).toNumber()] = _177_v39;
      _nw12[(new BigNumber(12)).toNumber()] = _177_v39;
      _nw12[(new BigNumber(13)).toNumber()] = _177_v39;
      _nw12[(new BigNumber(14)).toNumber()] = _177_v39;
      _nw12[(new BigNumber(15)).toNumber()] = _177_v39;
      _nw12[(new BigNumber(16)).toNumber()] = _177_v39;
      _nw12[(new BigNumber(17)).toNumber()] = _177_v39;
      _nw12[(new BigNumber(18)).toNumber()] = _177_v39;
      _nw12[(new BigNumber(19)).toNumber()] = _177_v39;
      _nw12[(new BigNumber(20)).toNumber()] = (_178_v40).dtor_cf104;
      _nw12[(new BigNumber(21)).toNumber()] = _177_v39;
      _nw12[(new BigNumber(22)).toNumber()] = _177_v39;
      _nw12[(new BigNumber(23)).toNumber()] = _177_v39;
      _nw12[(new BigNumber(24)).toNumber()] = _177_v39;
      _nw12[(new BigNumber(25)).toNumber()] = (((_177_v39).f10) ? (_177_v39) : (_177_v39));
      _179_v41 = _nw12;
      let _index4 = _module.__default.safeIndex(new BigNumber(846), new BigNumber((_179_v41).length));
      (_179_v41)[_index4] = _177_v39;
      let _180_v42;
      _180_v42 = _dafny.Seq.of(_177_v39, _177_v39, _177_v39);
      let _index5 = _module.__default.safeIndex(new BigNumber(846), new BigNumber((_179_v41).length));
      (_179_v41)[_index5] = (_180_v42)[_module.__default.safeIndex(_module.__default.safeModuloInt(_138_v8, _138_v8), new BigNumber((_180_v42).length))];
      _138_v8 = new BigNumber(-556);
      let _181_v43;
      let _nw13 = new _module.C4();
      _nw13.__ctor((_177_v39).f9, (_177_v39).f10);
      _181_v43 = _nw13;
      let _182_v44;
      _182_v44 = _module.D30.create_DC73(_181_v43);
      _181_v43 = (((_132_v3).f12) ? (_181_v43) : ((_182_v44).dtor_cf127));
      let _183_v45;
      _183_v45 = _dafny.Seq.of(_138_v8);
      let _184_v46;
      _184_v46 = _module.D1.create_DC2(_183_v45);
      let _pat_let_tv0 = _132_v3;
      let _pat_let_tv1 = _138_v8;
      let _pat_let_tv2 = _138_v8;
      _138_v8 = function (_source7) {
        if (_source7.is_DC3) {
          let _185___mcc_h0 = (_source7).cf5;
          let _186_cf5 = _185___mcc_h0;
          return new BigNumber(138);
        } else if (_source7.is_DC4) {
          let _187___mcc_h1 = (_source7).cf6;
          let _188___mcc_h2 = (_source7).cf7;
          let _189___mcc_h3 = (_source7).cf8;
          let _190_cf8 = _189___mcc_h3;
          let _191_cf7 = _188___mcc_h2;
          let _192_cf6 = _187___mcc_h1;
          return new BigNumber((_dafny.Seq.of(_191_cf7, !(_191_cf7), (_pat_let_tv0).f12)).length);
        } else if (_source7.is_DC2) {
          let _193___mcc_h4 = (_source7).cf4;
          let _194_cf4 = _193___mcc_h4;
          return new BigNumber(558);
        } else {
          let _195___mcc_h5 = (_source7).cf9;
          let _196_cf9 = _195___mcc_h5;
          return _module.__default.safeModuloInt(_pat_let_tv1, _pat_let_tv2);
        }
      }(_184_v46);
      let _index6 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length));
      (_127_v0)[_index6] = (_132_v3).fm32(new BigNumber(((_132_v3).f13).length), _175_v37, _131_globalState);
      let _index7 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length));
      (_127_v0)[_index7] = (_132_v3).f12;
      if (!(_129_v1) || ((_177_v39).f10)) {
        if ((_177_v39).f10) {
          let _197_v47;
          let _nw14 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          _197_v47 = _nw14;
          let _198_v48;
          _198_v48 = _dafny.Map.Empty.slice().updateUnsafe(_197_v47,_138_v8);
          let _199_v49;
          _199_v49 = _198_v48;
          let _200_v50;
          _200_v50 = _dafny.Set.fromElements(_199_v49);
          let _201_v51;
          _201_v51 = _dafny.Seq.of(_200_v50);
          let _202_v52;
          let _nw15 = Array((new BigNumber(2)).toNumber());
          _nw15[(_dafny.ZERO).toNumber()] = _dafny.Seq.of(_200_v50);
          _nw15[(_dafny.ONE).toNumber()] = _201_v51;
          _202_v52 = _nw15;
          let _index8 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_202_v52).length));
          (_202_v52)[_index8] = _201_v51;
          let _index9 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_197_v47).length));
          (_197_v47)[_index9] = _138_v8;
          let _index10 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_202_v52).length));
          let _index11 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_197_v47).length));
          let _index12 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length));
          let _rhs9 = _dafny.Seq.Concat(_201_v51, _201_v51);
          let _rhs10 = (((_dafny.Map.Empty.slice().updateUnsafe(_138_v8,_175_v37)).contains(_138_v8)) ? (_138_v8) : (_138_v8));
          let _rhs11 = _module.__default.safeModuloInt((_183_v45)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_138_v8,_138_v8)).length), new BigNumber((_183_v45).length))], (_module.D8.create_DC24(true, _dafny.Set.fromElements(_197_v47, _197_v47, _197_v47), _138_v8)).dtor_cf44);
          let _rhs12 = _129_v1;
          let _lhs3 = _202_v52;
          let _lhs4 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_202_v52).length));
          let _lhs5 = _197_v47;
          let _lhs6 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_197_v47).length));
          let _lhs7 = _127_v0;
          let _lhs8 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length));
          _lhs3[_lhs4] = _rhs9;
          _lhs5[_lhs6] = _rhs10;
          _138_v8 = _rhs11;
          _lhs7[_lhs8] = _rhs12;
          let _203_v53;
          let _init2 = ((_204_v3) => function (_205_i7) {
            return (_204_v3).f13;
          })(_132_v3);
          let _nw16 = Array((new BigNumber(17)).toNumber());
          for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw16.length); _i0_2++) {
            _nw16[_i0_2] = _init2(new BigNumber(_i0_2));
          }
          _203_v53 = _nw16;
          let _index13 = _module.__default.safeIndex(new BigNumber(198), new BigNumber((_203_v53).length));
          (_203_v53)[_index13] = (_132_v3).f13;
          let _index14 = _module.__default.safeIndex(new BigNumber(198), new BigNumber((_203_v53).length));
          (_203_v53)[_index14] = (_132_v3).f13;
          let _206_v54;
          _206_v54 = _module.D30.create_DC74((_132_v3).f13);
          let _207_v55;
          _207_v55 = _module.D30.create_DC76(_module.D30.create_DC76(_206_v54));
          let _pat_let_tv3 = _206_v54;
          let _208_v56;
          let _nw17 = Array((new BigNumber(15)).toNumber());
          _nw17[(_dafny.ZERO).toNumber()] = _207_v55;
          _nw17[(_dafny.ONE).toNumber()] = _207_v55;
          _nw17[(new BigNumber(2)).toNumber()] = _207_v55;
          _nw17[(new BigNumber(3)).toNumber()] = _207_v55;
          _nw17[(new BigNumber(4)).toNumber()] = _207_v55;
          _nw17[(new BigNumber(5)).toNumber()] = _207_v55;
          _nw17[(new BigNumber(6)).toNumber()] = _207_v55;
          _nw17[(new BigNumber(7)).toNumber()] = _207_v55;
          _nw17[(new BigNumber(8)).toNumber()] = _207_v55;
          _nw17[(new BigNumber(9)).toNumber()] = _207_v55;
          _nw17[(new BigNumber(10)).toNumber()] = _207_v55;
          _nw17[(new BigNumber(11)).toNumber()] = function (_pat_let0_0) {
            return function (_209_dt__update__tmp_h0) {
              return function (_pat_let1_0) {
                return function (_210_dt__update_hcf130_h0) {
                  return _module.D30.create_DC76(_210_dt__update_hcf130_h0);
                }(_pat_let1_0);
              }(_pat_let_tv3);
            }(_pat_let0_0);
          }(_207_v55);
          _nw17[(new BigNumber(12)).toNumber()] = _207_v55;
          _nw17[(new BigNumber(13)).toNumber()] = _207_v55;
          _nw17[(new BigNumber(14)).toNumber()] = _207_v55;
          _208_v56 = _nw17;
          let _index15 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_208_v56).length));
          (_208_v56)[_index15] = _207_v55;
          let _211_v57;
          _211_v57 = _dafny.Seq.of((_177_v39).f10);
          let _212_v58;
          _212_v58 = _dafny.MultiSet.fromElements(!((_132_v3).f12), !((_177_v39).fm0(_211_v57, (_177_v39).f10, _138_v8, false, _131_globalState)), (_132_v3).f12);
          let _213_v59;
          let _nw18 = new _module.C2();
          _nw18.__ctor((_132_v3).f12, _127_v0, (_177_v39).f9, (_177_v39).f10);
          _213_v59 = _nw18;
          let _214_v60;
          _214_v60 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(955),_213_v59);
          let _215_v61;
          _215_v61 = _dafny.MultiSet.fromElements(new BigNumber((_214_v60).length), new BigNumber((_module.__default.fm15(true, (_177_v39).f10, _138_v8, _175_v37, _131_globalState)).cardinality()), (_197_v47)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_197_v47).length))]);
          let _216_v62;
          _216_v62 = _dafny.Seq.of(_132_v3);
          let _217_v63;
          _217_v63 = _dafny.Map.Empty.slice().updateUnsafe((_215_v61).Intersect(_module.__default.fm46((_177_v39).f10, (_132_v3).f12, (_132_v3).f12, _131_globalState)),_216_v62);
          let _218_v64;
          _218_v64 = _215_v61;
          let _index16 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_208_v56).length));
          let _rhs13 = !(_129_v1) || ((_127_v0)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length))]);
          let _rhs14 = _207_v55;
          let _rhs15 = (_132_v3).f12;
          let _rhs16 = _212_v58;
          let _rhs17 = (_dafny.Map.Empty.slice().updateUnsafe((_218_v64),_dafny.Seq.of(_132_v3))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_215_v61,_216_v62));
          let _lhs9 = _131_globalState;
          let _lhs10 = _208_v56;
          let _lhs11 = _module.__default.safeIndex(new BigNumber(307), new BigNumber((_208_v56).length));
          _lhs9.f8 = _rhs13;
          _lhs10[_lhs11] = _rhs14;
          _129_v1 = _rhs15;
          _212_v58 = _rhs16;
          _217_v63 = _rhs17;
          let _219_v65;
          _219_v65 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(865),new BigNumber((_dafny.Seq.update((_203_v53)[_module.__default.safeIndex(new BigNumber(198), new BigNumber((_203_v53).length))], _module.__default.safeIndex(_138_v8, new BigNumber(((_203_v53)[_module.__default.safeIndex(new BigNumber(198), new BigNumber((_203_v53).length))]).length)), _175_v37)).length));
          let _220_v66;
          _220_v66 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_132_v3).f13).length),(((_219_v65).contains(_138_v8)) ? ((_219_v65).get(_138_v8)) : (_138_v8)));
          let _221_v67;
          _221_v67 = _dafny.Seq.of((_220_v66).Merge(_219_v65));
          _221_v67 = _221_v67;
          let _index17 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_197_v47).length));
          (_197_v47)[_index17] = (_197_v47)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_197_v47).length))];
        } else {
          let _222_v68;
          _222_v68 = _dafny.Map.Empty.slice().updateUnsafe(_132_v3,_127_v0);
          let _223_v69;
          let _nw19 = Array((new BigNumber(27)).toNumber());
          _nw19[(_dafny.ZERO).toNumber()] = _127_v0;
          _nw19[(_dafny.ONE).toNumber()] = _127_v0;
          _nw19[(new BigNumber(2)).toNumber()] = _127_v0;
          _nw19[(new BigNumber(3)).toNumber()] = _127_v0;
          _nw19[(new BigNumber(4)).toNumber()] = _127_v0;
          _nw19[(new BigNumber(5)).toNumber()] = _127_v0;
          _nw19[(new BigNumber(6)).toNumber()] = _127_v0;
          _nw19[(new BigNumber(7)).toNumber()] = _127_v0;
          _nw19[(new BigNumber(8)).toNumber()] = _127_v0;
          _nw19[(new BigNumber(9)).toNumber()] = _127_v0;
          _nw19[(new BigNumber(10)).toNumber()] = _127_v0;
          _nw19[(new BigNumber(11)).toNumber()] = _127_v0;
          _nw19[(new BigNumber(12)).toNumber()] = _127_v0;
          _nw19[(new BigNumber(13)).toNumber()] = _127_v0;
          _nw19[(new BigNumber(14)).toNumber()] = _127_v0;
          _nw19[(new BigNumber(15)).toNumber()] = _127_v0;
          _nw19[(new BigNumber(16)).toNumber()] = (((_222_v68).contains(_132_v3)) ? ((_222_v68).get(_132_v3)) : (_127_v0));
          _nw19[(new BigNumber(17)).toNumber()] = _127_v0;
          _nw19[(new BigNumber(18)).toNumber()] = _127_v0;
          _nw19[(new BigNumber(19)).toNumber()] = _127_v0;
          _nw19[(new BigNumber(20)).toNumber()] = _127_v0;
          _nw19[(new BigNumber(21)).toNumber()] = _127_v0;
          _nw19[(new BigNumber(22)).toNumber()] = (((_177_v39).f10) ? (_127_v0) : (_127_v0));
          _nw19[(new BigNumber(23)).toNumber()] = _127_v0;
          _nw19[(new BigNumber(24)).toNumber()] = _127_v0;
          _nw19[(new BigNumber(25)).toNumber()] = _127_v0;
          _nw19[(new BigNumber(26)).toNumber()] = _127_v0;
          _223_v69 = _nw19;
          let _224_v70;
          _224_v70 = _module.D32.create_DC78(_223_v69);
          _223_v69 = (_224_v70).dtor_cf132;
          (_131_globalState).f8 = true;
          _130_v2 = (_130_v2).update(new BigNumber(((_132_v3).f13).length), (_132_v3).f12);
          let _225_v71;
          let _nw20 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
          _225_v71 = _nw20;
          let _index18 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_225_v71).length));
          (_225_v71)[_index18] = (_dafny.ZERO).minus(_138_v8);
          let _index19 = _module.__default.safeIndex(new BigNumber(126), new BigNumber((_225_v71).length));
          (_225_v71)[_index19] = new BigNumber(299);
          let _226_v72;
          _226_v72 = _dafny.MultiSet.fromElements((_225_v71)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_225_v71).length))], (_225_v71)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_225_v71).length))], (_225_v71)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_225_v71).length))], new BigNumber(178));
          let _227_v73;
          _227_v73 = _dafny.Set.fromElements(_138_v8, new BigNumber(((_226_v72).update(new BigNumber(((_132_v3).f13).length), _module.__default.abs((_225_v71)[_module.__default.safeIndex(new BigNumber(126), new BigNumber((_225_v71).length))]))).cardinality()));
          _227_v73 = (_227_v73).Union(_227_v73);
        }
        let _228_v74;
        let _229_v75;
        let _230_v76;
        let _231_v77;
        let _out0;
        let _out1;
        let _out2;
        let _out3;
        let _outcollector0 = (_181_v43).m0(_131_globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _out2 = _outcollector0[2];
        _out3 = _outcollector0[3];
        _228_v74 = _out0;
        _229_v75 = _out1;
        _230_v76 = _out2;
        _231_v77 = _out3;
        let _232_v78;
        let _233_v79;
        let _234_v80;
        let _235_v81;
        let _out4;
        let _out5;
        let _out6;
        let _out7;
        let _outcollector1 = (_181_v43).m0(_131_globalState);
        _out4 = _outcollector1[0];
        _out5 = _outcollector1[1];
        _out6 = _outcollector1[2];
        _out7 = _outcollector1[3];
        _232_v78 = _out4;
        _233_v79 = _out5;
        _234_v80 = _out6;
        _235_v81 = _out7;
        (_181_v43).m5(_131_globalState);
        let _236_v82;
        _236_v82 = _dafny.Set.fromElements(_module.__default.fm37(_131_globalState));
        _236_v82 = _236_v82;
      } else {
        (_131_globalState).f8 = _129_v1;
        let _237_v83;
        let _init3 = function (_238_i8) {
          return new _dafny.CodePoint('n'.codePointAt(0));
        };
        let _nw21 = Array((new BigNumber(3)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw21.length); _i0_3++) {
          _nw21[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _237_v83 = _nw21;
        let _239_v84;
        let _nw22 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
        _239_v84 = _nw22;
        let _index20 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_239_v84).length));
        (_239_v84)[_index20] = _138_v8;
        let _index21 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length));
        let _index22 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_239_v84).length));
        let _rhs18 = (_138_v8).minus(_138_v8);
        let _rhs19 = (_177_v39).f10;
        let _rhs20 = _237_v83;
        let _rhs21 = _138_v8;
        let _rhs22 = _138_v8;
        let _lhs12 = _127_v0;
        let _lhs13 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length));
        let _lhs14 = _239_v84;
        let _lhs15 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_239_v84).length));
        _138_v8 = _rhs18;
        _lhs12[_lhs13] = _rhs19;
        _237_v83 = _rhs20;
        _138_v8 = _rhs21;
        _lhs14[_lhs15] = _rhs22;
        let _240_v85;
        _240_v85 = _dafny.Seq.of((_module.D3.create_DC7(_239_v84)).dtor_cf11, _239_v84);
        let _241_v86;
        _241_v86 = _dafny.Set.fromElements((_240_v85)[_module.__default.safeIndex(_138_v8, new BigNumber((_240_v85).length))], _239_v84, _239_v84);
        if (!((_241_v86).IsDisjointFrom(_241_v86))) {
          _129_v1 = ((((_177_v39).f9).contains((_177_v39).f10)) ? (((_177_v39).f9).get((_177_v39).f10)) : ((new BigNumber((_dafny.Seq.UnicodeFromString("nlwuecg")).length)).isLessThan((_dafny.ZERO).minus((_239_v84)[_module.__default.safeIndex(new BigNumber(825), new BigNumber((_239_v84).length))]))));
          _138_v8 = (_dafny.ZERO).minus(_138_v8);
          let _index23 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_239_v84).length));
          (_239_v84)[_index23] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(19)), ((_242_v37) => function (_243_i9) {
            return _242_v37;
          })(_175_v37))).length);
          let _244_v87;
          _244_v87 = _dafny.Map.Empty.slice().updateUnsafe(!(_129_v1),_138_v8);
          let _245_v88;
          _245_v88 = _dafny.Set.fromElements((_132_v3).f13);
          _138_v8 = (_dafny.ZERO).minus((_module.__default.safeDivisionInt(new BigNumber((_244_v87).length), (_239_v84)[_module.__default.safeIndex(new BigNumber(825), new BigNumber((_239_v84).length))])).plus(new BigNumber((_245_v88).length)));
          _138_v8 = _module.__default.safeDivisionInt(_138_v8, (_138_v8).multipliedBy(_138_v8));
        } else {
          let _246_v89;
          _246_v89 = _dafny.Map.Empty.slice().updateUnsafe(_138_v8,_183_v45);
          let _247_v91;
          _247_v91 = _dafny.MultiSet.fromElements((_132_v3).f12, (_177_v39).f10, (_132_v3).f12);
          let _248_v92;
          _248_v92 = _dafny.Set.fromElements(new BigNumber((_247_v91).cardinality()), _138_v8);
          let _249_v96;
          _249_v96 = _dafny.MultiSet.fromElements((_239_v84)[_module.__default.safeIndex(new BigNumber(825), new BigNumber((_239_v84).length))]);
          let _250_v98;
          let _nw23 = Array((new BigNumber(16)).toNumber());
          _nw23[(_dafny.ZERO).toNumber()] = _246_v89;
          _nw23[(_dafny.ONE).toNumber()] = (function () {
            let _coll22 = new _dafny.Map();
            for (const _compr_22 of (_248_v92).Elements) {
              let _251_v90 = _compr_22;
              if ((_248_v92).contains(_251_v90)) {
                _coll22.push([(_251_v90).minus(_138_v8),_183_v45]);
              }
            }
            return _coll22;
          }()).Merge(_246_v89);
          _nw23[(new BigNumber(2)).toNumber()] = _246_v89;
          _nw23[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(380),_183_v45);
          _nw23[(new BigNumber(4)).toNumber()] = _246_v89;
          _nw23[(new BigNumber(5)).toNumber()] = (_246_v89).Merge((_module.D33.create_DC81(_246_v89)).dtor_cf139);
          _nw23[(new BigNumber(6)).toNumber()] = (function () {
            let _coll23 = new _dafny.Map();
            for (const _compr_23 of (_183_v45).Elements) {
              let _252_v93 = _compr_23;
              if (_dafny.Seq.contains(_183_v45, _252_v93)) {
                _coll23.push([_module.__default.safeModuloInt(_252_v93, (_239_v84)[_module.__default.safeIndex(new BigNumber(825), new BigNumber((_239_v84).length))]),_183_v45]);
              }
            }
            return _coll23;
          }()).Merge(_246_v89);
          _nw23[(new BigNumber(7)).toNumber()] = function () {
            let _coll24 = new _dafny.Map();
            for (const _compr_24 of _dafny.IntegerRange(new BigNumber(-33), new BigNumber(-393))) {
              let _253_v94 = _compr_24;
              if (((new BigNumber(-33)).isLessThanOrEqualTo(_253_v94)) && ((_253_v94).isLessThan(new BigNumber(-393)))) {
                _coll24.push([(_253_v94).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(603)), ((_254_v37) => function (_255_i10) {
                  return _254_v37;
                })(_175_v37))).length)),_183_v45]);
              }
            }
            return _coll24;
          }();
          _nw23[(new BigNumber(8)).toNumber()] = _246_v89;
          _nw23[(new BigNumber(9)).toNumber()] = _246_v89;
          _nw23[(new BigNumber(10)).toNumber()] = function () {
            let _coll25 = new _dafny.Map();
            for (const _compr_25 of (_249_v96).Elements) {
              let _256_v95 = _compr_25;
              if ((_249_v96).contains(_256_v95)) {
                _coll25.push([(_256_v95).multipliedBy((_239_v84)[_module.__default.safeIndex(new BigNumber(825), new BigNumber((_239_v84).length))]),_dafny.Seq.of(_138_v8, (_239_v84)[_module.__default.safeIndex(new BigNumber(825), new BigNumber((_239_v84).length))])]);
              }
            }
            return _coll25;
          }();
          _nw23[(new BigNumber(11)).toNumber()] = _246_v89;
          _nw23[(new BigNumber(12)).toNumber()] = _246_v89;
          _nw23[(new BigNumber(13)).toNumber()] = function () {
            let _coll26 = new _dafny.Map();
            for (const _compr_26 of (_249_v96).Elements) {
              let _257_v97 = _compr_26;
              if ((_249_v96).contains(_257_v97)) {
                _coll26.push([_module.__default.safeModuloInt(_257_v97, new BigNumber((_dafny.Seq.UnicodeFromString("fvnf")).length)),_183_v45]);
              }
            }
            return _coll26;
          }();
          _nw23[(new BigNumber(14)).toNumber()] = _246_v89;
          _nw23[(new BigNumber(15)).toNumber()] = _246_v89;
          _250_v98 = _nw23;
          let _index24 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_250_v98).length));
          (_250_v98)[_index24] = _246_v89;
          let _index25 = _module.__default.safeIndex(new BigNumber(633), new BigNumber((_250_v98).length));
          (_250_v98)[_index25] = ((false) ? ((_246_v89).Merge(_246_v89)) : ((_dafny.Map.Empty.slice().updateUnsafe(_138_v8,_183_v45)).update((_239_v84)[_module.__default.safeIndex(new BigNumber(825), new BigNumber((_239_v84).length))], _dafny.Seq.update(_183_v45, _module.__default.safeIndex((_239_v84)[_module.__default.safeIndex(new BigNumber(825), new BigNumber((_239_v84).length))], new BigNumber((_183_v45).length)), new BigNumber((_180_v42).length)))));
          let _rhs23 = (_dafny.ZERO).minus(_138_v8);
          let _rhs24 = !((_132_v3).f12);
          let _rhs25 = _239_v84;
          let _lhs16 = _131_globalState;
          _138_v8 = _rhs23;
          _lhs16.f8 = _rhs24;
          _239_v84 = _rhs25;
          let _index26 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length));
          (_127_v0)[_index26] = (_177_v39).f10;
          (_131_globalState).f8 = (_177_v39).f10;
          let _258_v99;
          _258_v99 = _dafny.Set.fromElements((_177_v39).f10);
          let _259_v100;
          let _nw24 = new _module.C7();
          _nw24.__ctor(_258_v99, (_177_v39).f9, (_132_v3).fm32((_239_v84)[_module.__default.safeIndex(new BigNumber(825), new BigNumber((_239_v84).length))], _175_v37, _131_globalState));
          _259_v100 = _nw24;
        }
        let _260_v101;
        _260_v101 = _dafny.Seq.of((_127_v0)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length))]);
        let _261_v102;
        _261_v102 = _dafny.Seq.of((_132_v3).f13);
        _138_v8 = (_module.__default.safeModuloInt(new BigNumber((_260_v101).length), _138_v8)).plus(new BigNumber((_261_v102).length));
        (_131_globalState).f8 = (_132_v3).f12;
      }
      if ((_127_v0)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length))]) {
        _138_v8 = (_138_v8).minus(new BigNumber(-444));
        _175_v37 = _175_v37;
        let _262_v103;
        _262_v103 = _dafny.Seq.UnicodeFromString("elcmh");
        let _index27 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length));
        let _rhs26 = !((new BigNumber(645)).isLessThan(_138_v8));
        let _rhs27 = _dafny.Seq.Concat((_132_v3).f13, (_132_v3).f13);
        let _lhs17 = _127_v0;
        let _lhs18 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length));
        _lhs17[_lhs18] = _rhs26;
        _262_v103 = _rhs27;
        (_181_v43).m5(_131_globalState);
        let _263_v104;
        _263_v104 = _dafny.Seq.of(_130_v2);
        _263_v104 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(395)), ((_264_v2) => function (_265_i11) {
          return (_264_v2).Merge(_264_v2);
        })(_130_v2));
      } else {
        _138_v8 = (_138_v8).minus(_138_v8);
        let _266_v105;
        let _nw25 = new _module.C0();
        _nw25.__ctor();
        _266_v105 = _nw25;
        let _267_v106;
        _267_v106 = _dafny.Seq.of(_127_v0);
        let _268_v107;
        let _nw26 = Array((new BigNumber(26)).toNumber());
        _nw26[(_dafny.ZERO).toNumber()] = _127_v0;
        _nw26[(_dafny.ONE).toNumber()] = _127_v0;
        _nw26[(new BigNumber(2)).toNumber()] = _127_v0;
        _nw26[(new BigNumber(3)).toNumber()] = _127_v0;
        _nw26[(new BigNumber(4)).toNumber()] = _127_v0;
        _nw26[(new BigNumber(5)).toNumber()] = _127_v0;
        _nw26[(new BigNumber(6)).toNumber()] = _127_v0;
        _nw26[(new BigNumber(7)).toNumber()] = _127_v0;
        _nw26[(new BigNumber(8)).toNumber()] = _127_v0;
        _nw26[(new BigNumber(9)).toNumber()] = _127_v0;
        _nw26[(new BigNumber(10)).toNumber()] = _127_v0;
        _nw26[(new BigNumber(11)).toNumber()] = _127_v0;
        _nw26[(new BigNumber(12)).toNumber()] = _127_v0;
        _nw26[(new BigNumber(13)).toNumber()] = _127_v0;
        _nw26[(new BigNumber(14)).toNumber()] = _127_v0;
        _nw26[(new BigNumber(15)).toNumber()] = (_267_v106)[_module.__default.safeIndex((_dafny.ZERO).minus(_138_v8), new BigNumber((_267_v106).length))];
        _nw26[(new BigNumber(16)).toNumber()] = _127_v0;
        _nw26[(new BigNumber(17)).toNumber()] = _127_v0;
        _nw26[(new BigNumber(18)).toNumber()] = _127_v0;
        _nw26[(new BigNumber(19)).toNumber()] = _127_v0;
        _nw26[(new BigNumber(20)).toNumber()] = _127_v0;
        _nw26[(new BigNumber(21)).toNumber()] = _127_v0;
        _nw26[(new BigNumber(22)).toNumber()] = _127_v0;
        _nw26[(new BigNumber(23)).toNumber()] = _127_v0;
        _nw26[(new BigNumber(24)).toNumber()] = _127_v0;
        _nw26[(new BigNumber(25)).toNumber()] = _127_v0;
        _268_v107 = _nw26;
        let _index28 = _module.__default.safeIndex(new BigNumber(934), new BigNumber((_268_v107).length));
        (_268_v107)[_index28] = _127_v0;
        let _index29 = _module.__default.safeIndex(new BigNumber(934), new BigNumber((_268_v107).length));
        (_268_v107)[_index29] = _127_v0;
        let _269_v108;
        let _nw27 = new _module.C1();
        _nw27.__ctor((_132_v3).f12, (_132_v3).f13);
        _269_v108 = _nw27;
        _269_v108 = _269_v108;
        let _270_v109;
        let _nw28 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Seq.of());
        _270_v109 = _nw28;
        let _index30 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_270_v109).length));
        (_270_v109)[_index30] = _267_v106;
        let _index31 = _module.__default.safeIndex(new BigNumber(974), new BigNumber((_270_v109).length));
        (_270_v109)[_index31] = _dafny.Seq.Concat(_dafny.Seq.of((_268_v107)[_module.__default.safeIndex(new BigNumber(934), new BigNumber((_268_v107).length))]), _dafny.Seq.Concat(_267_v106, _267_v106));
      }
      (_131_globalState).f8 = !((_177_v39).f10);
      let _271_v110;
      _271_v110 = _dafny.MultiSet.fromElements((_132_v3).f12, (_132_v3).f12);
      let _272_v111;
      let _nw29 = new _module.C1();
      _nw29.__ctor(!(new BigNumber(948)).isEqualTo(new BigNumber((_271_v110).cardinality())), _dafny.Seq.Concat((_132_v3).f13, (_132_v3).f13));
      _272_v111 = _nw29;
      let _273_v112;
      let _nw30 = new _module.C5();
      _nw30.__ctor(_dafny.Map.Empty.slice().updateUnsafe(_129_v1,true), (_127_v0)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length))]);
      _273_v112 = _nw30;
      let _274_v113;
      _274_v113 = _dafny.Map.Empty.slice().updateUnsafe(_273_v112,_138_v8);
      let _index32 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length));
      let _rhs28 = (_132_v3).fm32(_138_v8, _175_v37, _131_globalState);
      let _rhs29 = _272_v111;
      let _rhs30 = _129_v1;
      let _rhs31 = _module.__default.safeModuloInt(_module.__default.safeModuloInt(_138_v8, (((_274_v113).contains(_273_v112)) ? ((_274_v113).get(_273_v112)) : (_138_v8))), _138_v8);
      let _lhs19 = _131_globalState;
      let _lhs20 = _127_v0;
      let _lhs21 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length));
      _lhs19.f8 = _rhs28;
      _272_v111 = _rhs29;
      _lhs20[_lhs21] = _rhs30;
      _138_v8 = _rhs31;
      if ((_177_v39).f10) {
        let _index33 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length));
        (_127_v0)[_index33] = (_127_v0)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length))];
        let _275_v114;
        _275_v114 = _dafny.Map.Empty.slice().updateUnsafe(_129_v1,_dafny.MultiSet.FromArray(_183_v45));
        let _276_v115;
        _276_v115 = _dafny.MultiSet.fromElements(_138_v8, new BigNumber((_dafny.Seq.of((_177_v39).f10, (_132_v3).f12)).length));
        let _277_v116;
        _277_v116 = _dafny.Seq.of(_276_v115, _276_v115, (_276_v115).update(_138_v8, _module.__default.abs(_138_v8)), _276_v115, _276_v115);
        (_131_globalState).f8 = ((_275_v114).update((_127_v0)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length))], (_277_v116)[_module.__default.safeIndex(_138_v8, new BigNumber((_277_v116).length))])).contains((((_127_v0)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length))]) ? ((_127_v0)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length))]) : (true)));
        let _278_v117;
        let _nw31 = Array((new BigNumber(24)).toNumber()).fill([]);
        _278_v117 = _nw31;
        let _279_v118;
        _279_v118 = _module.D32.create_DC78(_278_v117);
        let _source8 = _279_v118;
        if (_source8.is_DC79) {
          let _280___mcc_h6 = (_source8).cf133;
          let _281___mcc_h7 = (_source8).cf134;
          let _282___mcc_h8 = (_source8).cf135;
          let _283_cf135 = _282___mcc_h8;
          let _284_cf134 = _281___mcc_h7;
          let _285_cf133 = _280___mcc_h6;
          let _286_v119;
          _286_v119 = _dafny.Seq.of(_129_v1);
          let _287_v120;
          _287_v120 = _dafny.Set.fromElements((_132_v3).f12, (_177_v39).fm0(_286_v119, (_132_v3).f12, _138_v8, (_127_v0)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length))], _131_globalState));
          _287_v120 = (_module.__default.fm34(_138_v8, false, (_177_v39).f10, true, _131_globalState)).Union(_287_v120);
          let _288_v121;
          _288_v121 = _module.D0.create_DC1(_138_v8, (_127_v0)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length))], _138_v8);
          let _289_v122;
          _289_v122 = _dafny.Seq.of(_272_v111, _272_v111);
          let _290_v123;
          _290_v123 = _dafny.Seq.of(_289_v122);
          let _291_v124;
          _291_v124 = _dafny.Map.Empty.slice().updateUnsafe(_138_v8,_175_v37);
          let _292_v125;
          let _nw32 = Array((new BigNumber(18)).toNumber());
          _nw32[(_dafny.ZERO).toNumber()] = _138_v8;
          _nw32[(_dafny.ONE).toNumber()] = _138_v8;
          _nw32[(new BigNumber(2)).toNumber()] = _138_v8;
          _nw32[(new BigNumber(3)).toNumber()] = _138_v8;
          _nw32[(new BigNumber(4)).toNumber()] = new BigNumber((_290_v123).length);
          _nw32[(new BigNumber(5)).toNumber()] = new BigNumber((_183_v45).length);
          _nw32[(new BigNumber(6)).toNumber()] = new BigNumber((_287_v120).length);
          _nw32[(new BigNumber(7)).toNumber()] = _138_v8;
          _nw32[(new BigNumber(8)).toNumber()] = _138_v8;
          _nw32[(new BigNumber(9)).toNumber()] = new BigNumber((_286_v119).length);
          _nw32[(new BigNumber(10)).toNumber()] = _138_v8;
          _nw32[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(_138_v8);
          _nw32[(new BigNumber(12)).toNumber()] = _138_v8;
          _nw32[(new BigNumber(13)).toNumber()] = new BigNumber((_291_v124).length);
          _nw32[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus(_138_v8);
          _nw32[(new BigNumber(15)).toNumber()] = _138_v8;
          _nw32[(new BigNumber(16)).toNumber()] = _138_v8;
          _nw32[(new BigNumber(17)).toNumber()] = _138_v8;
          _292_v125 = _nw32;
          let _293_v126;
          _293_v126 = _dafny.Set.fromElements(_292_v125, _292_v125, _292_v125, _292_v125, _292_v125);
          let _294_v127;
          _294_v127 = _module.D8.create_DC24((_288_v121).dtor_cf2, (_293_v126).Intersect(_293_v126), (_273_v112).fm1((_177_v39).f9, _131_globalState));
          _294_v127 = _294_v127;
          _175_v37 = _175_v37;
          let _295_v128;
          _295_v128 = _dafny.Set.fromElements(new BigNumber(921), _138_v8);
          let _index34 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length));
          (_127_v0)[_index34] = (_295_v128).IsSubsetOf((_295_v128).Intersect(_295_v128));
        } else if (_source8.is_DC80) {
          let _296___mcc_h9 = (_source8).cf136;
          let _297___mcc_h10 = (_source8).cf137;
          let _298___mcc_h11 = (_source8).cf138;
          let _299_cf138 = _298___mcc_h11;
          let _300_cf137 = _297___mcc_h10;
          let _301_cf136 = _296___mcc_h9;
          (_131_globalState).f8 = _129_v1;
          _278_v117 = _278_v117;
          _301_cf136 = _module.__default.safeDivisionInt(_299_cf138, _300_cf137);
          let _302_v129;
          _302_v129 = _dafny.Set.fromElements((_132_v3).f12, (_177_v39).f10);
          let _303_v130;
          _303_v130 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_302_v129).length),(_132_v3).f13);
          let _304_v131;
          _304_v131 = _dafny.Seq.of(_130_v2, (_module.__default.fm11(new BigNumber((_276_v115).cardinality()), new BigNumber((_183_v45).length), (_177_v39).f10, new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe((_177_v39).f10,(_179_v41)[_module.__default.safeIndex(new BigNumber(846), new BigNumber((_179_v41).length))])).update((_132_v3).f12, (_179_v41)[_module.__default.safeIndex(new BigNumber(846), new BigNumber((_179_v41).length))])).length), _131_globalState)).update(_300_cf137, _129_v1));
          let _305_v132;
          _305_v132 = _module.D16.create_DC40(_304_v131);
          let _pat_let_tv4 = _304_v131;
          let _306_v133;
          _306_v133 = _dafny.Map.Empty.slice().updateUnsafe(!_dafny.Seq.contains((((_303_v130).contains((_dafny.ZERO).minus(_300_cf137))) ? ((_303_v130).get((_dafny.ZERO).minus(_300_cf137))) : ((_132_v3).f13)), new _dafny.CodePoint('c'.codePointAt(0))),function (_pat_let2_0) {
            return function (_307_dt__update__tmp_h1) {
              return function (_pat_let3_0) {
                return function (_308_dt__update_hcf76_h0) {
                  return _module.D16.create_DC40(_308_dt__update_hcf76_h0);
                }(_pat_let3_0);
              }(_pat_let_tv4);
            }(_pat_let2_0);
          }(_305_v132));
          _306_v133 = (_306_v133).update(_dafny.Seq.IsPrefixOf(_183_v45, _dafny.Seq.of(new BigNumber(153))), _305_v132);
        } else {
          let _309___mcc_h12 = (_source8).cf132;
          let _310_cf132 = _309___mcc_h12;
          _175_v37 = _175_v37;
          let _311_v134;
          _311_v134 = _dafny.Seq.of((_177_v39).f10);
          (_131_globalState).f8 = (_311_v134)[_module.__default.safeIndex(_138_v8, new BigNumber((_311_v134).length))];
          let _312_v135;
          let _313_v136;
          let _314_v137;
          let _315_v138;
          let _out8;
          let _out9;
          let _out10;
          let _out11;
          let _outcollector2 = (_273_v112).m0(_131_globalState);
          _out8 = _outcollector2[0];
          _out9 = _outcollector2[1];
          _out10 = _outcollector2[2];
          _out11 = _outcollector2[3];
          _312_v135 = _out8;
          _313_v136 = _out9;
          _314_v137 = _out10;
          _315_v138 = _out11;
          _315_v138 = (_132_v3).f13;
        }
        let _316_v139;
        _316_v139 = _module.D25.create_DC63((_132_v3).f13, _138_v8);
        let _317_v140;
        _317_v140 = _dafny.Set.fromElements((_316_v139).dtor_cf113);
        let _318_v141;
        let _nw33 = Array((new BigNumber(25)).toNumber());
        _nw33[(_dafny.ZERO).toNumber()] = _138_v8;
        _nw33[(_dafny.ONE).toNumber()] = _138_v8;
        _nw33[(new BigNumber(2)).toNumber()] = _138_v8;
        _nw33[(new BigNumber(3)).toNumber()] = _138_v8;
        _nw33[(new BigNumber(4)).toNumber()] = _138_v8;
        _nw33[(new BigNumber(5)).toNumber()] = _138_v8;
        _nw33[(new BigNumber(6)).toNumber()] = _138_v8;
        _nw33[(new BigNumber(7)).toNumber()] = _138_v8;
        _nw33[(new BigNumber(8)).toNumber()] = _138_v8;
        _nw33[(new BigNumber(9)).toNumber()] = _138_v8;
        _nw33[(new BigNumber(10)).toNumber()] = _138_v8;
        _nw33[(new BigNumber(11)).toNumber()] = _138_v8;
        _nw33[(new BigNumber(12)).toNumber()] = _138_v8;
        _nw33[(new BigNumber(13)).toNumber()] = _138_v8;
        _nw33[(new BigNumber(14)).toNumber()] = _138_v8;
        _nw33[(new BigNumber(15)).toNumber()] = new BigNumber((_317_v140).length);
        _nw33[(new BigNumber(16)).toNumber()] = new BigNumber(((_132_v3).f13).length);
        _nw33[(new BigNumber(17)).toNumber()] = _138_v8;
        _nw33[(new BigNumber(18)).toNumber()] = _138_v8;
        _nw33[(new BigNumber(19)).toNumber()] = _138_v8;
        _nw33[(new BigNumber(20)).toNumber()] = _138_v8;
        _nw33[(new BigNumber(21)).toNumber()] = _module.__default.fm33(_129_v1, (_177_v39).f10, _138_v8, (_177_v39).f10, _131_globalState);
        _nw33[(new BigNumber(22)).toNumber()] = _138_v8;
        _nw33[(new BigNumber(23)).toNumber()] = new BigNumber(((_132_v3).f13).length);
        _nw33[(new BigNumber(24)).toNumber()] = _138_v8;
        _318_v141 = _nw33;
        let _319_v142;
        _319_v142 = _dafny.Set.fromElements(_318_v141);
        let _320_v143;
        _320_v143 = _module.D8.create_DC24((_177_v39).f10, _319_v142, _138_v8);
        let _321_v144;
        _321_v144 = _module.D16.create_DC42(_129_v1, (_320_v143).dtor_cf44, (_273_v112).fm19(_138_v8, _138_v8, _131_globalState));
        let _source9 = _321_v144;
        if (_source9.is_DC41) {
          let _322___mcc_h13 = (_source9).cf77;
          let _323___mcc_h14 = (_source9).cf78;
          let _324___mcc_h15 = (_source9).cf79;
          let _325_cf79 = _324___mcc_h15;
          let _326_cf78 = _323___mcc_h14;
          let _327_cf77 = _322___mcc_h13;
          _129_v1 = !(new BigNumber((((_176_v38).update((_177_v39).f10, (_127_v0)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length))])).Merge((_177_v39).f9)).length)).isEqualTo(((_dafny.ZERO).minus(_138_v8)).plus(_326_cf78));
          (_181_v43).m5(_131_globalState);
          let _328_v145;
          _328_v145 = _module.D21.create_DC53(_272_v111);
          let _pat_let_tv5 = _272_v111;
          _328_v145 = function (_pat_let4_0) {
            return function (_329_dt__update__tmp_h2) {
              return function (_pat_let5_0) {
                return function (_330_dt__update_hcf99_h0) {
                  return _module.D21.create_DC53(_330_dt__update_hcf99_h0);
                }(_pat_let5_0);
              }(_pat_let_tv5);
            }(_pat_let4_0);
          }(_328_v145);
          (_131_globalState).f8 = !_dafny.Seq.contains((_132_v3).f13, _175_v37);
        } else if (_source9.is_DC42) {
          let _331___mcc_h16 = (_source9).cf80;
          let _332___mcc_h17 = (_source9).cf81;
          let _333___mcc_h18 = (_source9).cf82;
          let _334_cf82 = _333___mcc_h18;
          let _335_cf81 = _332___mcc_h17;
          let _336_cf80 = _331___mcc_h16;
          let _337_v146;
          _337_v146 = _dafny.Map.Empty.slice().updateUnsafe(!((_132_v3).f12),((_dafny.ZERO).minus(_335_cf81)).minus(_334_cf82));
          let _rhs32 = (_127_v0)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length))];
          let _rhs33 = _337_v146;
          let _rhs34 = !(!((_132_v3).f12) || (!((_132_v3).f12))) || (!((_127_v0)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length))]) || ((_177_v39).f10));
          let _lhs22 = _131_globalState;
          let _lhs23 = _131_globalState;
          _lhs22.f8 = _rhs32;
          _337_v146 = _rhs33;
          _lhs23.f8 = _rhs34;
          _138_v8 = _334_cf82;
          let _338_v147;
          _338_v147 = _dafny.Set.fromElements(_127_v0);
          let _index35 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length));
          let _rhs35 = (_177_v39).f10;
          let _rhs36 = (_334_cf82).plus((_181_v43).fm1((_177_v39).f9, _131_globalState));
          let _rhs37 = _338_v147;
          let _lhs24 = _127_v0;
          let _lhs25 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length));
          _lhs24[_lhs25] = _rhs35;
          _335_cf81 = _rhs36;
          _338_v147 = _rhs37;
          let _339_v148;
          _339_v148 = _dafny.Seq.UnicodeFromString("kspfey");
          _339_v148 = _dafny.Seq.Concat((_132_v3).f13, _dafny.Seq.Create(_module.__default.abs(new BigNumber(784)), ((_340_v37) => function (_341_i12) {
            return _340_v37;
          })(_175_v37)));
        } else {
          let _342___mcc_h19 = (_source9).cf76;
          let _343_cf76 = _342___mcc_h19;
          let _index36 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length));
          (_127_v0)[_index36] = (_177_v39).f10;
          let _344_v149;
          let _345_v150;
          let _346_v151;
          let _347_v152;
          let _out12;
          let _out13;
          let _out14;
          let _out15;
          let _outcollector3 = (_181_v43).m0(_131_globalState);
          _out12 = _outcollector3[0];
          _out13 = _outcollector3[1];
          _out14 = _outcollector3[2];
          _out15 = _outcollector3[3];
          _344_v149 = _out12;
          _345_v150 = _out13;
          _346_v151 = _out14;
          _347_v152 = _out15;
          let _348_v153;
          let _349_v154;
          let _350_v155;
          let _351_v156;
          let _out16;
          let _out17;
          let _out18;
          let _out19;
          let _outcollector4 = (_132_v3).m0(_131_globalState);
          _out16 = _outcollector4[0];
          _out17 = _outcollector4[1];
          _out18 = _outcollector4[2];
          _out19 = _outcollector4[3];
          _348_v153 = _out16;
          _349_v154 = _out17;
          _350_v155 = _out18;
          _351_v156 = _out19;
          (_131_globalState).f8 = _345_v150;
        }
        let _352_v157;
        _352_v157 = _dafny.Map.Empty.slice().updateUnsafe(_138_v8,_138_v8);
        _138_v8 = _module.__default.safeModuloInt(_138_v8, _module.__default.safeModuloInt(new BigNumber((_352_v157).length), _138_v8));
      } else {
        let _353_v158;
        let _354_v159;
        let _355_v160;
        let _356_v161;
        let _out20;
        let _out21;
        let _out22;
        let _out23;
        let _outcollector5 = (_273_v112).m0(_131_globalState);
        _out20 = _outcollector5[0];
        _out21 = _outcollector5[1];
        _out22 = _outcollector5[2];
        _out23 = _outcollector5[3];
        _353_v158 = _out20;
        _354_v159 = _out21;
        _355_v160 = _out22;
        _356_v161 = _out23;
        let _357_v162;
        let _nw34 = new _module.C4();
        _nw34.__ctor(_176_v38, ((_132_v3).f12) || ((_177_v39).f10));
        _357_v162 = _nw34;
        let _358_v163;
        _358_v163 = _dafny.Map.Empty.slice().updateUnsafe(_272_v111,_129_v1);
        let _359_v164;
        _359_v164 = _dafny.Map.Empty.slice().updateUnsafe(_358_v163,_138_v8);
        let _360_v165;
        _360_v165 = _module.D32.create_DC80(_138_v8, _138_v8, new BigNumber(665));
        let _361_v166;
        _361_v166 = _dafny.MultiSet.fromElements(new BigNumber(879));
        let _362_v167;
        let _nw35 = Array((new BigNumber(3)).toNumber()).fill([]);
        _362_v167 = _nw35;
        let _363_v168;
        _363_v168 = _dafny.Map.Empty.slice().updateUnsafe(_362_v167,_356_v161);
        let _364_v169;
        let _nw36 = Array((new BigNumber(23)).toNumber());
        _nw36[(_dafny.ZERO).toNumber()] = _module.__default.fm22(_138_v8, _129_v1, (_177_v39).f10, _138_v8, _131_globalState);
        _nw36[(_dafny.ONE).toNumber()] = _353_v158;
        _nw36[(new BigNumber(2)).toNumber()] = _356_v161;
        _nw36[(new BigNumber(3)).toNumber()] = _dafny.Seq.update((_132_v3).f13, _module.__default.safeIndex(new BigNumber(((_132_v3).f13).length), new BigNumber(((_132_v3).f13).length)), _175_v37);
        _nw36[(new BigNumber(4)).toNumber()] = (_132_v3).f13;
        _nw36[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("yieoik");
        _nw36[(new BigNumber(6)).toNumber()] = _356_v161;
        _nw36[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat((_132_v3).f13, _353_v158);
        _nw36[(new BigNumber(8)).toNumber()] = _dafny.Seq.update((_132_v3).f13, _module.__default.safeIndex(new BigNumber((_359_v164).length), new BigNumber(((_132_v3).f13).length)), _175_v37);
        _nw36[(new BigNumber(9)).toNumber()] = _module.__default.fm13((_360_v165).dtor_cf138, _131_globalState);
        _nw36[(new BigNumber(10)).toNumber()] = _356_v161;
        _nw36[(new BigNumber(11)).toNumber()] = _dafny.Seq.update((_132_v3).f13, _module.__default.safeIndex(_138_v8, new BigNumber(((_132_v3).f13).length)), new _dafny.CodePoint('p'.codePointAt(0)));
        _nw36[(new BigNumber(12)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(476)), ((_365_v37) => function (_366_i13) {
          return _365_v37;
        })(_175_v37));
        _nw36[(new BigNumber(13)).toNumber()] = _dafny.Seq.UnicodeFromString("a");
        _nw36[(new BigNumber(14)).toNumber()] = _module.__default.fm45((_132_v3).f12, (((_361_v166).contains(_138_v8)) ? ((_361_v166).get(_138_v8)) : (_138_v8)), (_132_v3).f12, _131_globalState);
        _nw36[(new BigNumber(15)).toNumber()] = _353_v158;
        _nw36[(new BigNumber(16)).toNumber()] = _dafny.Seq.update((_132_v3).f13, _module.__default.safeIndex(new BigNumber((_353_v158).length), new BigNumber(((_132_v3).f13).length)), new _dafny.CodePoint('p'.codePointAt(0)));
        _nw36[(new BigNumber(17)).toNumber()] = _dafny.Seq.UnicodeFromString("iybtk");
        _nw36[(new BigNumber(18)).toNumber()] = _module.__default.fm45(_354_v159, _module.__default.fm33(_129_v1, _354_v159, (_dafny.ZERO).minus(_138_v8), (_127_v0)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length))], _131_globalState), !(!(!((_132_v3).f12))), _131_globalState);
        _nw36[(new BigNumber(19)).toNumber()] = (((_177_v39).f10) ? (_dafny.Seq.UnicodeFromString("wgwipapdd")) : ((_132_v3).f13));
        _nw36[(new BigNumber(20)).toNumber()] = _353_v158;
        _nw36[(new BigNumber(21)).toNumber()] = _dafny.Seq.Concat(_353_v158, _353_v158);
        _nw36[(new BigNumber(22)).toNumber()] = (((_363_v168).contains(_362_v167)) ? ((_363_v168).get(_362_v167)) : ((_132_v3).f13));
        _364_v169 = _nw36;
        let _367_v170;
        _367_v170 = _dafny.Map.Empty.slice().updateUnsafe(_138_v8,_364_v169);
        let _rhs38 = (_357_v162).fm1(_dafny.Map.Empty.slice().updateUnsafe((_177_v39).f10,(_127_v0)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length))]), _131_globalState);
        let _rhs39 = (_132_v3).f13;
        let _rhs40 = (((_367_v170).contains(new BigNumber(((_132_v3).f13).length))) ? ((_367_v170).get(new BigNumber(((_132_v3).f13).length))) : (_364_v169));
        let _rhs41 = _129_v1;
        _138_v8 = _rhs38;
        _356_v161 = _rhs39;
        _364_v169 = _rhs40;
        _354_v159 = _rhs41;
        let _368_v171;
        _368_v171 = _dafny.Map.Empty.slice().updateUnsafe((_132_v3).f12,_127_v0);
        _368_v171 = _dafny.Map.Empty.slice().updateUnsafe(!((_177_v39).f10) || (_129_v1),_127_v0);
        _129_v1 = (_177_v39).f10;
      }
      let _369_v172;
      _369_v172 = _dafny.Seq.of((_177_v39).f10, ((_129_v1) ? ((_177_v39).f10) : (!(_129_v1))), _dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(534)), ((_370_v37) => function (_371_i14) {
        return _370_v37;
      })(_175_v37)), (_132_v3).f13));
      if ((_369_v172)[_module.__default.safeIndex(_138_v8, new BigNumber((_369_v172).length))]) {
        let _372_v173;
        let _373_v174;
        let _out24;
        let _out25;
        let _outcollector6 = (_272_v111).m1(_175_v37, _131_globalState);
        _out24 = _outcollector6[0];
        _out25 = _outcollector6[1];
        _372_v173 = _out24;
        _373_v174 = _out25;
        (_181_v43).m5(_131_globalState);
        let _374_v175;
        let _375_v176;
        let _376_v177;
        let _377_v178;
        let _out26;
        let _out27;
        let _out28;
        let _out29;
        let _outcollector7 = (_273_v112).m0(_131_globalState);
        _out26 = _outcollector7[0];
        _out27 = _outcollector7[1];
        _out28 = _outcollector7[2];
        _out29 = _outcollector7[3];
        _374_v175 = _out26;
        _375_v176 = _out27;
        _376_v177 = _out28;
        _377_v178 = _out29;
        let _378_v179;
        let _nw37 = new _module.C0();
        _nw37.__ctor();
        _378_v179 = _nw37;
        let _379_v181;
        _379_v181 = _module.D33.create_DC81(function () {
  let _coll27 = new _dafny.Map();
  for (const _compr_27 of _dafny.IntegerRange(new BigNumber(547), new BigNumber(232))) {
    let _380_v180 = _compr_27;
    if (((new BigNumber(547)).isLessThanOrEqualTo(_380_v180)) && ((_380_v180).isLessThan(new BigNumber(232)))) {
      _coll27.push([(_380_v180).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(7)), ((_381_v37) => function (_382_i15) {
        return _381_v37;
      })(_175_v37))).length)),_183_v45]);
    }
  }
  return _coll27;
}());
        let _383_v182;
        _383_v182 = _module.D33.create_DC83(_379_v181);
        let _384_v183;
        _384_v183 = _module.D33.create_DC83(_379_v181);
        _384_v183 = _384_v183;
      } else {
        let _385_v184;
        _385_v184 = _dafny.Seq.UnicodeFromString("jytbikv");
        _385_v184 = _dafny.Seq.update(_dafny.Seq.update((_132_v3).f13, _module.__default.safeIndex(_138_v8, new BigNumber(((_132_v3).f13).length)), _175_v37), _module.__default.safeIndex(_138_v8, new BigNumber((_dafny.Seq.update((_132_v3).f13, _module.__default.safeIndex(_138_v8, new BigNumber(((_132_v3).f13).length)), _175_v37)).length)), _175_v37);
        let _386_v185;
        _386_v185 = _dafny.MultiSet.fromElements(_138_v8);
        _386_v185 = (_module.__default.fm61(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(986),_138_v8), _131_globalState));
        let _index37 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length));
        (_127_v0)[_index37] = !((_132_v3).f12) || (!(((_132_v3).f12) === ((_369_v172)[_module.__default.safeIndex(_138_v8, new BigNumber((_369_v172).length))])));
        _138_v8 = (new BigNumber((_dafny.Seq.Concat(_385_v184, _385_v184)).length)).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-69)), ((_387_v37) => function (_388_i16) {
          return _387_v37;
        })(_175_v37))).length));
        let _389_v186;
        _389_v186 = _module.D14.create_DC36(_129_v1, _385_v184, (_132_v3).f13);
        let _source10 = _389_v186;
        if (_source10.is_DC35) {
          let _390___mcc_h20 = (_source10).cf62;
          let _391___mcc_h21 = (_source10).cf63;
          let _392___mcc_h22 = (_source10).cf64;
          let _393_cf64 = _392___mcc_h22;
          let _394_cf63 = _391___mcc_h21;
          let _395_cf62 = _390___mcc_h20;
          let _396_v187;
          _396_v187 = _dafny.Map.Empty.slice().updateUnsafe(_138_v8,_395_cf62);
          _395_cf62 = (((_396_v187).contains((_module.D5.create_DC14(new BigNumber((_385_v184).length), (_132_v3).f12)).dtor_cf25)) ? ((_396_v187).get((_module.D5.create_DC14(new BigNumber((_385_v184).length), (_132_v3).f12)).dtor_cf25)) : (_395_cf62));
          _395_cf62 = _138_v8;
          let _397_v188;
          let _init4 = ((_398_v8) => function (_399_i17) {
            return (_399_i17).minus(_398_v8);
          })(_138_v8);
          let _nw38 = Array((new BigNumber(3)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw38.length); _i0_4++) {
            _nw38[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _397_v188 = _nw38;
          let _400_v189;
          _400_v189 = _dafny.Set.fromElements(_138_v8);
          let _index38 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_397_v188).length));
          (_397_v188)[_index38] = new BigNumber((_400_v189).length);
          let _401_v190;
          _401_v190 = _dafny.Map.Empty.slice().updateUnsafe(_395_cf62,_183_v45);
          let _index39 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length));
          let _index40 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_397_v188).length));
          let _rhs42 = (_132_v3).f12;
          let _rhs43 = !(_386_v185).contains(new BigNumber(((_132_v3).f13).length));
          let _rhs44 = _127_v0;
          let _rhs45 = (_dafny.ZERO).minus(new BigNumber(((_401_v190).Merge(function () {
            let _coll28 = new _dafny.Map();
            for (const _compr_28 of _dafny.IntegerRange(new BigNumber(492), new BigNumber(183))) {
              let _402_v191 = _compr_28;
              if (((new BigNumber(492)).isLessThanOrEqualTo(_402_v191)) && ((_402_v191).isLessThan(new BigNumber(183)))) {
                _coll28.push([_module.__default.safeDivisionInt(_402_v191, new BigNumber(-771)),_dafny.Seq.Create(_module.__default.abs(new BigNumber(996)), ((_403_cf62) => function (_404_i18) {
                  return _403_cf62;
                })(_395_cf62))]);
              }
            }
            return _coll28;
          }())).length));
          let _lhs26 = _127_v0;
          let _lhs27 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length));
          let _lhs28 = _131_globalState;
          let _lhs29 = _397_v188;
          let _lhs30 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_397_v188).length));
          _lhs26[_lhs27] = _rhs42;
          _lhs28.f8 = _rhs43;
          _127_v0 = _rhs44;
          _lhs29[_lhs30] = _rhs45;
          _395_cf62 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt(_395_cf62, new BigNumber(270)), _395_cf62);
        } else if (_source10.is_DC36) {
          let _405___mcc_h23 = (_source10).cf65;
          let _406___mcc_h24 = (_source10).cf66;
          let _407___mcc_h25 = (_source10).cf67;
          let _408_cf67 = _407___mcc_h25;
          let _409_cf66 = _406___mcc_h24;
          let _410_cf65 = _405___mcc_h23;
          let _411_v192;
          let _412_v193;
          let _out30;
          let _out31;
          let _outcollector8 = (_272_v111).m1(_175_v37, _131_globalState);
          _out30 = _outcollector8[0];
          _out31 = _outcollector8[1];
          _411_v192 = _out30;
          _412_v193 = _out31;
          let _413_v194;
          let _414_v195;
          let _out32;
          let _out33;
          let _outcollector9 = (_177_v39).m1(_175_v37, _131_globalState);
          _out32 = _outcollector9[0];
          _out33 = _outcollector9[1];
          _413_v194 = _out32;
          _414_v195 = _out33;
          _410_cf65 = _dafny.Seq.contains(_409_cf66, _175_v37);
          _369_v172 = (((_177_v39).f10) ? (_369_v172) : (_369_v172));
        } else {
          let _415___mcc_h26 = (_source10).cf61;
          let _416_cf61 = _415___mcc_h26;
          let _417_v196;
          _417_v196 = _dafny.Map.Empty.slice().updateUnsafe(_138_v8,_dafny.Seq.Concat((_132_v3).f13, (_132_v3).f13));
          _417_v196 = (_417_v196).update(_138_v8, _385_v184);
          let _418_v197;
          _418_v197 = _dafny.Map.Empty.slice().updateUnsafe(_175_v37,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_129_v1,_138_v8)).length));
          let _419_v198;
          _419_v198 = _dafny.Set.fromElements(_138_v8, (_dafny.ZERO).minus(_138_v8), new BigNumber((_dafny.Seq.update(_dafny.Seq.of((_132_v3).f12), _module.__default.safeIndex(new BigNumber((_418_v197).length), new BigNumber((_dafny.Seq.of((_132_v3).f12)).length)), !((_132_v3).f12))).length));
          (_131_globalState).f8 = (_419_v198).IsDisjointFrom(_419_v198);
          let _index41 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length));
          (_127_v0)[_index41] = (_138_v8).isLessThan(new BigNumber((_419_v198).length));
          let _index42 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length));
          let _rhs46 = (_132_v3).f12;
          let _rhs47 = _127_v0;
          let _lhs31 = _127_v0;
          let _lhs32 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length));
          _lhs31[_lhs32] = _rhs46;
          _127_v0 = _rhs47;
        }
      }
      let _source11 = _module.D16.create_DC41(!(_129_v1) || ((_127_v0)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length))]), (_dafny.ZERO).minus(_138_v8), (_127_v0)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length))]);
      if (_source11.is_DC41) {
        let _420___mcc_h27 = (_source11).cf77;
        let _421___mcc_h28 = (_source11).cf78;
        let _422___mcc_h29 = (_source11).cf79;
        let _423_cf79 = _422___mcc_h29;
        let _424_cf78 = _421___mcc_h28;
        let _425_cf77 = _420___mcc_h27;
        (_131_globalState).f8 = true;
        let _426_v199;
        _426_v199 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_129_v1,new BigNumber(((_132_v3).f13).length)), _dafny.Map.Empty.slice().updateUnsafe((_177_v39).f10,_138_v8));
        let _427_v200;
        _427_v200 = _dafny.Map.Empty.slice().updateUnsafe(_129_v1,_424_cf78);
        _138_v8 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_426_v199, _dafny.Seq.of(_427_v200)), _dafny.Seq.Concat(_426_v199, _dafny.Seq.of(_427_v200)))).length);
        _138_v8 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_424_cf78));
        let _428_v201;
        _428_v201 = _dafny.Set.fromElements(_424_cf78, new BigNumber((_369_v172).length));
        let _429_v202;
        _429_v202 = _dafny.Set.fromElements(new BigNumber((_428_v201).length), _138_v8, _424_cf78, _138_v8, _424_cf78);
        let _index43 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length));
        let _rhs48 = (_429_v202).IsSubsetOf(_428_v201);
        let _rhs49 = _424_cf78;
        let _rhs50 = _425_cf77;
        let _lhs33 = _131_globalState;
        let _lhs34 = _127_v0;
        let _lhs35 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length));
        _lhs33.f8 = _rhs48;
        _424_cf78 = _rhs49;
        _lhs34[_lhs35] = _rhs50;
      } else if (_source11.is_DC42) {
        let _430___mcc_h30 = (_source11).cf80;
        let _431___mcc_h31 = (_source11).cf81;
        let _432___mcc_h32 = (_source11).cf82;
        let _433_cf82 = _432___mcc_h32;
        let _434_cf81 = _431___mcc_h31;
        let _435_cf80 = _430___mcc_h30;
        let _436_v203;
        let _nw39 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
        _436_v203 = _nw39;
        let _index44 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_436_v203).length));
        (_436_v203)[_index44] = _434_cf81;
        let _437_v204;
        let _nw40 = Array((new BigNumber(12)).toNumber()).fill(_module.D12.Default());
        _437_v204 = _nw40;
        let _index45 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_436_v203).length));
        let _rhs51 = (_dafny.ZERO).minus(_433_cf82);
        let _rhs52 = _437_v204;
        let _lhs36 = _436_v203;
        let _lhs37 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_436_v203).length));
        _lhs36[_lhs37] = _rhs51;
        _437_v204 = _rhs52;
        let _438_v205;
        let _439_v206;
        let _440_v207;
        let _441_v208;
        let _out34;
        let _out35;
        let _out36;
        let _out37;
        let _outcollector10 = (_177_v39).m0(_131_globalState);
        _out34 = _outcollector10[0];
        _out35 = _outcollector10[1];
        _out36 = _outcollector10[2];
        _out37 = _outcollector10[3];
        _438_v205 = _out34;
        _439_v206 = _out35;
        _440_v207 = _out36;
        _441_v208 = _out37;
        let _442_v209;
        _442_v209 = _dafny.Set.fromElements((_177_v39).f10, (_132_v3).f12);
        let _443_v210;
        let _nw41 = new _module.C7();
        _nw41.__ctor((_442_v209).Union(_module.__default.fm34(_434_cf81, true, (_132_v3).f12, (_177_v39).f10, _131_globalState)), _176_v38, ((_439_v206) ? (_439_v206) : ((_127_v0)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length))])));
        _443_v210 = _nw41;
        let _444_v211;
        let _nw42 = Array((new BigNumber(25)).toNumber());
        _444_v211 = _nw42;
        let _445_v212;
        let _nw43 = new _module.C6();
        _nw43.__ctor(_176_v38, false);
        _445_v212 = _nw43;
        let _index46 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_444_v211).length));
        (_444_v211)[_index46] = _445_v212;
        let _446_v213;
        let _nw44 = Array((new BigNumber(24)).toNumber()).fill(_module.D1.Default());
        _446_v213 = _nw44;
        let _447_v214;
        _447_v214 = _dafny.Seq.of(_446_v213);
        let _448_v215;
        _448_v215 = _module.D15.create_DC37((_447_v214)[_module.__default.safeIndex(_138_v8, new BigNumber((_447_v214).length))]);
        let _449_v216;
        _449_v216 = _dafny.Seq.of(_dafny.Set.fromElements(_138_v8, _138_v8));
        let _450_v217;
        _450_v217 = _dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(true, (_132_v3).f12)).length), _138_v8);
        let _451_v218;
        _451_v218 = _dafny.Set.fromElements(_438_v205, (_132_v3).f13, (_132_v3).f13);
        let _index47 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_444_v211).length));
        let _index48 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_436_v203).length));
        let _index49 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length));
        let _rhs53 = _443_v210;
        let _rhs54 = _445_v212;
        let _rhs55 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_449_v216, _dafny.Seq.of(_450_v217))).length));
        let _rhs56 = _module.D15.create_DC37(_446_v213);
        let _rhs57 = !((_451_v218).contains(_dafny.Seq.Concat((_132_v3).f13, (_132_v3).f13)));
        let _lhs38 = _444_v211;
        let _lhs39 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_444_v211).length));
        let _lhs40 = _436_v203;
        let _lhs41 = _module.__default.safeIndex(new BigNumber(552), new BigNumber((_436_v203).length));
        let _lhs42 = _127_v0;
        let _lhs43 = _module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length));
        _443_v210 = _rhs53;
        _lhs38[_lhs39] = _rhs54;
        _lhs40[_lhs41] = _rhs55;
        _448_v215 = _rhs56;
        _lhs42[_lhs43] = _rhs57;
        _435_cf80 = (_127_v0)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length))];
      } else {
        let _452___mcc_h33 = (_source11).cf76;
        let _453_cf76 = _452___mcc_h33;
        _183_v45 = _183_v45;
        _138_v8 = _138_v8;
        _127_v0 = _127_v0;
        let _454_v219;
        _454_v219 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_132_v3).f13).length),_369_v172);
        _138_v8 = _module.__default.fm33((_127_v0)[_module.__default.safeIndex(new BigNumber(701), new BigNumber((_127_v0).length))], (new BigNumber(((((_454_v219).contains(new BigNumber(-795))) ? ((_454_v219).get(new BigNumber(-795))) : ((((_454_v219).contains(_138_v8)) ? ((_454_v219).get(_138_v8)) : (_369_v172))))).length)).isLessThanOrEqualTo(_138_v8), (_138_v8).multipliedBy(_138_v8), (_132_v3).f12, _131_globalState);
      }
      let _455_v220;
      let _456_v221;
      let _457_v222;
      let _458_v223;
      let _out38;
      let _out39;
      let _out40;
      let _out41;
      let _outcollector11 = (_177_v39).m0(_131_globalState);
      _out38 = _outcollector11[0];
      _out39 = _outcollector11[1];
      _out40 = _outcollector11[2];
      _out41 = _outcollector11[3];
      _455_v220 = _out38;
      _456_v221 = _out39;
      _457_v222 = _out40;
      _458_v223 = _out41;
      process.stdout.write(_dafny.toString((_127_v0)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v0)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v0)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v0)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v0)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v0)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v0)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v0)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v0)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v0)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v0)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v0)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v0)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v0)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v0)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v0)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v0)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_v0)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_129_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v2).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(4),true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_globalState).f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_131_globalState).f2)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_131_globalState).f2)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_131_globalState).f2)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_131_globalState).f2)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_131_globalState).f2)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_131_globalState).f2)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_131_globalState).f2)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_131_globalState).f2)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_131_globalState).f2)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_131_globalState).f2)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_131_globalState).f2)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_131_globalState).f2)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_131_globalState).f2)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_131_globalState).f2)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_131_globalState).f2)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_131_globalState).f2)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_131_globalState).f2)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_131_globalState).f2)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_globalState.f4).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(4),true),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_131_globalState).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_131_globalState.f8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_132_v3).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_132_v3).f13, _dafny.Seq.of(new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0)), new _dafny.CodePoint('n'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_136_v6).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_137_v7).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_138_v8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_175_v37));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_176_v38).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_177_v39).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_177_v39).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_178_v40).dtor_cf104).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_178_v40).dtor_cf104).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_179_v41)[_dafny.ZERO]).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_179_v41)[_dafny.ZERO]).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_179_v41)[_dafny.ONE]).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_179_v41)[_dafny.ONE]).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_179_v41)[new BigNumber(2)]).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_179_v41)[new BigNumber(2)]).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_179_v41)[new BigNumber(3)]).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_179_v41)[new BigNumber(3)]).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_179_v41)[new BigNumber(4)]).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_179_v41)[new BigNumber(4)]).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_179_v41)[new BigNumber(5)]).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_179_v41)[new BigNumber(5)]).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_179_v41)[new BigNumber(6)]).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_179_v41)[new BigNumber(6)]).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_179_v41)[new BigNumber(7)]).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_179_v41)[new BigNumber(7)]).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_179_v41)[new BigNumber(8)]).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_179_v41)[new BigNumber(8)]).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_179_v41)[new BigNumber(9)]).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_179_v41)[new BigNumber(9)]).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_179_v41)[new BigNumber(10)]).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_179_v41)[new BigNumber(10)]).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_179_v41)[new BigNumber(11)]).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_179_v41)[new BigNumber(11)]).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_179_v41)[new BigNumber(12)]).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_179_v41)[new BigNumber(12)]).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_179_v41)[new BigNumber(13)]).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_179_v41)[new BigNumber(13)]).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_179_v41)[new BigNumber(14)]).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_179_v41)[new BigNumber(14)]).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_179_v41)[new BigNumber(15)]).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_179_v41)[new BigNumber(15)]).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_179_v41)[new BigNumber(16)]).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_179_v41)[new BigNumber(16)]).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_179_v41)[new BigNumber(17)]).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_179_v41)[new BigNumber(17)]).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_179_v41)[new BigNumber(18)]).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_179_v41)[new BigNumber(18)]).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_179_v41)[new BigNumber(19)]).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_179_v41)[new BigNumber(19)]).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_179_v41)[new BigNumber(20)]).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_179_v41)[new BigNumber(20)]).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_179_v41)[new BigNumber(21)]).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_179_v41)[new BigNumber(21)]).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_179_v41)[new BigNumber(22)]).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_179_v41)[new BigNumber(22)]).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_179_v41)[new BigNumber(23)]).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_179_v41)[new BigNumber(23)]).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_179_v41)[new BigNumber(24)]).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_179_v41)[new BigNumber(24)]).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_179_v41)[new BigNumber(25)]).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_179_v41)[new BigNumber(25)]).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_180_v42).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_182_v44).dtor_cf127).f9).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_182_v44).dtor_cf127).f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_183_v45, _dafny.Seq.of(new BigNumber(-556)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_184_v46).dtor_cf4, _dafny.Seq.of(new BigNumber(-556)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_271_v110).equals(_dafny.MultiSet.fromElements(true, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_274_v113).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_369_v172, _dafny.Seq.of(true, true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_455_v220).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_456_v221));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_457_v222, _dafny.Seq.of(_dafny.Set.fromElements(true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_458_v223).toVerbatimString(false));
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
    get is_DC1() { return this.$tag === 0; }
    get is_DC0() { return this.$tag === 1; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ")";
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
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1) && this.cf2 === other.cf2 && _dafny.areEqual(this.cf3, other.cf3);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf0 === other.cf0;
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
    static create_DC3(cf5) {
      let $dt = new D1(0);
      $dt.cf5 = cf5;
      return $dt;
    }
    static create_DC4(cf6, cf7, cf8) {
      let $dt = new D1(1);
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC2(cf4) {
      let $dt = new D1(2);
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC5(cf9) {
      let $dt = new D1(3);
      $dt.cf9 = cf9;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get is_DC4() { return this.$tag === 1; }
    get is_DC2() { return this.$tag === 2; }
    get is_DC5() { return this.$tag === 3; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf9() { return this.cf9; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC3" + "(" + this.cf5.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 2) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 3) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf9) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf5, other.cf5);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf6 === other.cf6 && this.cf7 === other.cf7 && _dafny.areEqual(this.cf8, other.cf8);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf9, other.cf9);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC3(_dafny.Seq.UnicodeFromString(""));
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
    static create_DC6(cf10) {
      let $dt = new D2(0);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf10) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return new _dafny.CodePoint('D'.codePointAt(0));
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
    static create_DC8(cf12, cf13) {
      let $dt = new D3(0);
      $dt.cf12 = cf12;
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC7(cf11) {
      let $dt = new D3(1);
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC9(cf14) {
      let $dt = new D3(2);
      $dt.cf14 = cf14;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get is_DC9() { return this.$tag === 2; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf14() { return this.cf14; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf12) + ", " + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC7" + "(" + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf14) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf12 === other.cf12 && this.cf13 === other.cf13;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf11 === other.cf11;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf14, other.cf14);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC8(false, false);
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
    static create_DC11(cf16, cf17, cf18) {
      let $dt = new D4(0);
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC12(cf19, cf20, cf21, cf22, cf23) {
      let $dt = new D4(1);
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      return $dt;
    }
    static create_DC10(cf15) {
      let $dt = new D4(2);
      $dt.cf15 = cf15;
      return $dt;
    }
    get is_DC11() { return this.$tag === 0; }
    get is_DC12() { return this.$tag === 1; }
    get is_DC10() { return this.$tag === 2; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf15() { return this.cf15; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf15) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf16, other.cf16) && _dafny.areEqual(this.cf17, other.cf17) && this.cf18 === other.cf18;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf19 === other.cf19 && _dafny.areEqual(this.cf20, other.cf20) && _dafny.areEqual(this.cf21, other.cf21) && _dafny.areEqual(this.cf22, other.cf22) && _dafny.areEqual(this.cf23, other.cf23);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf15, other.cf15);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC11(_dafny.ZERO, _dafny.Set.Empty, false);
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
    static create_DC14(cf25, cf26) {
      let $dt = new D5(0);
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      return $dt;
    }
    static create_DC13(cf24) {
      let $dt = new D5(1);
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC15(cf27) {
      let $dt = new D5(2);
      $dt.cf27 = cf27;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get is_DC15() { return this.$tag === 2; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf27() { return this.cf27; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf27) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf25, other.cf25) && this.cf26 === other.cf26;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf24 === other.cf24;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf27, other.cf27);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC14(_dafny.ZERO, false);
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
    static create_DC17(cf29, cf30, cf31, cf32, cf33) {
      let $dt = new D6(0);
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC18(cf34, cf35, cf36) {
      let $dt = new D6(1);
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      return $dt;
    }
    static create_DC19(cf37, cf38) {
      let $dt = new D6(2);
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC16(cf28) {
      let $dt = new D6(3);
      $dt.cf28 = cf28;
      return $dt;
    }
    static create_DC20(cf39) {
      let $dt = new D6(4);
      $dt.cf39 = cf39;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC18() { return this.$tag === 1; }
    get is_DC19() { return this.$tag === 2; }
    get is_DC16() { return this.$tag === 3; }
    get is_DC20() { return this.$tag === 4; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf39() { return this.cf39; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC18" + "(" + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC19" + "(" + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 3) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf28) + ")";
      } else if (this.$tag === 4) {
        return "D6.DC20" + "(" + _dafny.toString(this.cf39) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf29 === other.cf29 && this.cf30 === other.cf30 && _dafny.areEqual(this.cf31, other.cf31) && _dafny.areEqual(this.cf32, other.cf32) && _dafny.areEqual(this.cf33, other.cf33);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf34, other.cf34) && _dafny.areEqual(this.cf35, other.cf35) && _dafny.areEqual(this.cf36, other.cf36);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf37 === other.cf37 && _dafny.areEqual(this.cf38, other.cf38);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf28, other.cf28);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf39, other.cf39);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC17(false, false, _dafny.ZERO, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC22() {
      let $dt = new D7(0);
      return $dt;
    }
    static create_DC21(cf40) {
      let $dt = new D7(1);
      $dt.cf40 = cf40;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get is_DC21() { return this.$tag === 1; }
    get dtor_cf40() { return this.cf40; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC22";
      } else if (this.$tag === 1) {
        return "D7.DC21" + "(" + _dafny.toString(this.cf40) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf40, other.cf40);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC22();
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
    static create_DC24(cf42, cf43, cf44) {
      let $dt = new D8(0);
      $dt.cf42 = cf42;
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      return $dt;
    }
    static create_DC23(cf41) {
      let $dt = new D8(1);
      $dt.cf41 = cf41;
      return $dt;
    }
    get is_DC24() { return this.$tag === 0; }
    get is_DC23() { return this.$tag === 1; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf41() { return this.cf41; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC24" + "(" + _dafny.toString(this.cf42) + ", " + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC23" + "(" + _dafny.toString(this.cf41) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf42 === other.cf42 && _dafny.areEqual(this.cf43, other.cf43) && _dafny.areEqual(this.cf44, other.cf44);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf41, other.cf41);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC24(false, _dafny.Set.Empty, _dafny.ZERO);
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
    static create_DC25(cf45) {
      let $dt = new D9(0);
      $dt.cf45 = cf45;
      return $dt;
    }
    get is_DC25() { return this.$tag === 0; }
    get dtor_cf45() { return this.cf45; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC25" + "(" + _dafny.toString(this.cf45) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf45, other.cf45);
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
          return D9.Default();
        }
      };
    }
  }

  $module.D10 = class D10 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC26(cf46) {
      let $dt = new D10(0);
      $dt.cf46 = cf46;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get dtor_cf46() { return this.cf46; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC26" + "(" + _dafny.toString(this.cf46) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf46, other.cf46);
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
    static create_DC28(cf48, cf49, cf50, cf51, cf52) {
      let $dt = new D11(0);
      $dt.cf48 = cf48;
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      return $dt;
    }
    static create_DC27(cf47) {
      let $dt = new D11(1);
      $dt.cf47 = cf47;
      return $dt;
    }
    static create_DC29(cf53) {
      let $dt = new D11(2);
      $dt.cf53 = cf53;
      return $dt;
    }
    get is_DC28() { return this.$tag === 0; }
    get is_DC27() { return this.$tag === 1; }
    get is_DC29() { return this.$tag === 2; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf53() { return this.cf53; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC28" + "(" + _dafny.toString(this.cf48) + ", " + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ", " + _dafny.toString(this.cf51) + ", " + this.cf52.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC27" + "(" + _dafny.toString(this.cf47) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC29" + "(" + _dafny.toString(this.cf53) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf48, other.cf48) && this.cf49 === other.cf49 && _dafny.areEqual(this.cf50, other.cf50) && this.cf51 === other.cf51 && _dafny.areEqual(this.cf52, other.cf52);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf47, other.cf47);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf53, other.cf53);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC28(_dafny.ZERO, false, _dafny.Map.Empty, false, _dafny.Seq.UnicodeFromString(""));
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
    static create_DC31(cf55, cf56, cf57) {
      let $dt = new D12(0);
      $dt.cf55 = cf55;
      $dt.cf56 = cf56;
      $dt.cf57 = cf57;
      return $dt;
    }
    static create_DC32(cf58, cf59) {
      let $dt = new D12(1);
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      return $dt;
    }
    static create_DC30(cf54) {
      let $dt = new D12(2);
      $dt.cf54 = cf54;
      return $dt;
    }
    get is_DC31() { return this.$tag === 0; }
    get is_DC32() { return this.$tag === 1; }
    get is_DC30() { return this.$tag === 2; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf54() { return this.cf54; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC31" + "(" + _dafny.toString(this.cf55) + ", " + this.cf56.toVerbatimString(true) + ", " + this.cf57.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC32" + "(" + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC30" + "(" + _dafny.toString(this.cf54) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf55 === other.cf55 && _dafny.areEqual(this.cf56, other.cf56) && _dafny.areEqual(this.cf57, other.cf57);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf58 === other.cf58 && _dafny.areEqual(this.cf59, other.cf59);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf54 === other.cf54;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC31(false, _dafny.Seq.UnicodeFromString(""), _dafny.Seq.UnicodeFromString(""));
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
    static create_DC33(cf60) {
      let $dt = new D13(0);
      $dt.cf60 = cf60;
      return $dt;
    }
    get is_DC33() { return this.$tag === 0; }
    get dtor_cf60() { return this.cf60; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC33" + "(" + _dafny.toString(this.cf60) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf60, other.cf60);
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
    static create_DC35(cf62, cf63, cf64) {
      let $dt = new D14(0);
      $dt.cf62 = cf62;
      $dt.cf63 = cf63;
      $dt.cf64 = cf64;
      return $dt;
    }
    static create_DC36(cf65, cf66, cf67) {
      let $dt = new D14(1);
      $dt.cf65 = cf65;
      $dt.cf66 = cf66;
      $dt.cf67 = cf67;
      return $dt;
    }
    static create_DC34(cf61) {
      let $dt = new D14(2);
      $dt.cf61 = cf61;
      return $dt;
    }
    get is_DC35() { return this.$tag === 0; }
    get is_DC36() { return this.$tag === 1; }
    get is_DC34() { return this.$tag === 2; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf61() { return this.cf61; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC35" + "(" + _dafny.toString(this.cf62) + ", " + _dafny.toString(this.cf63) + ", " + _dafny.toString(this.cf64) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC36" + "(" + _dafny.toString(this.cf65) + ", " + this.cf66.toVerbatimString(true) + ", " + this.cf67.toVerbatimString(true) + ")";
      } else if (this.$tag === 2) {
        return "D14.DC34" + "(" + _dafny.toString(this.cf61) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf62, other.cf62) && this.cf63 === other.cf63 && _dafny.areEqual(this.cf64, other.cf64);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf65 === other.cf65 && _dafny.areEqual(this.cf66, other.cf66) && _dafny.areEqual(this.cf67, other.cf67);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf61, other.cf61);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC35(_dafny.ZERO, false, new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC38(cf69, cf70, cf71, cf72, cf73) {
      let $dt = new D15(0);
      $dt.cf69 = cf69;
      $dt.cf70 = cf70;
      $dt.cf71 = cf71;
      $dt.cf72 = cf72;
      $dt.cf73 = cf73;
      return $dt;
    }
    static create_DC39(cf74, cf75) {
      let $dt = new D15(1);
      $dt.cf74 = cf74;
      $dt.cf75 = cf75;
      return $dt;
    }
    static create_DC37(cf68) {
      let $dt = new D15(2);
      $dt.cf68 = cf68;
      return $dt;
    }
    get is_DC38() { return this.$tag === 0; }
    get is_DC39() { return this.$tag === 1; }
    get is_DC37() { return this.$tag === 2; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf73() { return this.cf73; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf68() { return this.cf68; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC38" + "(" + _dafny.toString(this.cf69) + ", " + this.cf70.toVerbatimString(true) + ", " + _dafny.toString(this.cf71) + ", " + _dafny.toString(this.cf72) + ", " + _dafny.toString(this.cf73) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC39" + "(" + _dafny.toString(this.cf74) + ", " + _dafny.toString(this.cf75) + ")";
      } else if (this.$tag === 2) {
        return "D15.DC37" + "(" + _dafny.toString(this.cf68) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf69 === other.cf69 && _dafny.areEqual(this.cf70, other.cf70) && this.cf71 === other.cf71 && this.cf72 === other.cf72 && _dafny.areEqual(this.cf73, other.cf73);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf74 === other.cf74 && this.cf75 === other.cf75;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf68 === other.cf68;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC38(false, _dafny.Seq.UnicodeFromString(""), false, false, new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC41(cf77, cf78, cf79) {
      let $dt = new D16(0);
      $dt.cf77 = cf77;
      $dt.cf78 = cf78;
      $dt.cf79 = cf79;
      return $dt;
    }
    static create_DC42(cf80, cf81, cf82) {
      let $dt = new D16(1);
      $dt.cf80 = cf80;
      $dt.cf81 = cf81;
      $dt.cf82 = cf82;
      return $dt;
    }
    static create_DC40(cf76) {
      let $dt = new D16(2);
      $dt.cf76 = cf76;
      return $dt;
    }
    get is_DC41() { return this.$tag === 0; }
    get is_DC42() { return this.$tag === 1; }
    get is_DC40() { return this.$tag === 2; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf79() { return this.cf79; }
    get dtor_cf80() { return this.cf80; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf82() { return this.cf82; }
    get dtor_cf76() { return this.cf76; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC41" + "(" + _dafny.toString(this.cf77) + ", " + _dafny.toString(this.cf78) + ", " + _dafny.toString(this.cf79) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC42" + "(" + _dafny.toString(this.cf80) + ", " + _dafny.toString(this.cf81) + ", " + _dafny.toString(this.cf82) + ")";
      } else if (this.$tag === 2) {
        return "D16.DC40" + "(" + _dafny.toString(this.cf76) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf77 === other.cf77 && _dafny.areEqual(this.cf78, other.cf78) && this.cf79 === other.cf79;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf80 === other.cf80 && _dafny.areEqual(this.cf81, other.cf81) && _dafny.areEqual(this.cf82, other.cf82);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf76, other.cf76);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC41(false, _dafny.ZERO, false);
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
    static create_DC44(cf84, cf85, cf86, cf87, cf88) {
      let $dt = new D17(0);
      $dt.cf84 = cf84;
      $dt.cf85 = cf85;
      $dt.cf86 = cf86;
      $dt.cf87 = cf87;
      $dt.cf88 = cf88;
      return $dt;
    }
    static create_DC45() {
      let $dt = new D17(1);
      return $dt;
    }
    static create_DC43(cf83) {
      let $dt = new D17(2);
      $dt.cf83 = cf83;
      return $dt;
    }
    get is_DC44() { return this.$tag === 0; }
    get is_DC45() { return this.$tag === 1; }
    get is_DC43() { return this.$tag === 2; }
    get dtor_cf84() { return this.cf84; }
    get dtor_cf85() { return this.cf85; }
    get dtor_cf86() { return this.cf86; }
    get dtor_cf87() { return this.cf87; }
    get dtor_cf88() { return this.cf88; }
    get dtor_cf83() { return this.cf83; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC44" + "(" + _dafny.toString(this.cf84) + ", " + _dafny.toString(this.cf85) + ", " + _dafny.toString(this.cf86) + ", " + _dafny.toString(this.cf87) + ", " + _dafny.toString(this.cf88) + ")";
      } else if (this.$tag === 1) {
        return "D17.DC45";
      } else if (this.$tag === 2) {
        return "D17.DC43" + "(" + _dafny.toString(this.cf83) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf84, other.cf84) && _dafny.areEqual(this.cf85, other.cf85) && _dafny.areEqual(this.cf86, other.cf86) && this.cf87 === other.cf87 && _dafny.areEqual(this.cf88, other.cf88);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf83 === other.cf83;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D17.create_DC44(_dafny.ZERO, _dafny.Seq.of(), new _dafny.CodePoint('D'.codePointAt(0)), false, _dafny.ZERO);
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
    static create_DC47(cf90, cf91) {
      let $dt = new D18(0);
      $dt.cf90 = cf90;
      $dt.cf91 = cf91;
      return $dt;
    }
    static create_DC46(cf89) {
      let $dt = new D18(1);
      $dt.cf89 = cf89;
      return $dt;
    }
    get is_DC47() { return this.$tag === 0; }
    get is_DC46() { return this.$tag === 1; }
    get dtor_cf90() { return this.cf90; }
    get dtor_cf91() { return this.cf91; }
    get dtor_cf89() { return this.cf89; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC47" + "(" + _dafny.toString(this.cf90) + ", " + _dafny.toString(this.cf91) + ")";
      } else if (this.$tag === 1) {
        return "D18.DC46" + "(" + _dafny.toString(this.cf89) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf90 === other.cf90 && _dafny.areEqual(this.cf91, other.cf91);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf89, other.cf89);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC47(false, _module.D6.Default());
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
    static create_DC49(cf93) {
      let $dt = new D19(0);
      $dt.cf93 = cf93;
      return $dt;
    }
    static create_DC50(cf94, cf95, cf96) {
      let $dt = new D19(1);
      $dt.cf94 = cf94;
      $dt.cf95 = cf95;
      $dt.cf96 = cf96;
      return $dt;
    }
    static create_DC48(cf92) {
      let $dt = new D19(2);
      $dt.cf92 = cf92;
      return $dt;
    }
    get is_DC49() { return this.$tag === 0; }
    get is_DC50() { return this.$tag === 1; }
    get is_DC48() { return this.$tag === 2; }
    get dtor_cf93() { return this.cf93; }
    get dtor_cf94() { return this.cf94; }
    get dtor_cf95() { return this.cf95; }
    get dtor_cf96() { return this.cf96; }
    get dtor_cf92() { return this.cf92; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC49" + "(" + _dafny.toString(this.cf93) + ")";
      } else if (this.$tag === 1) {
        return "D19.DC50" + "(" + _dafny.toString(this.cf94) + ", " + _dafny.toString(this.cf95) + ", " + _dafny.toString(this.cf96) + ")";
      } else if (this.$tag === 2) {
        return "D19.DC48" + "(" + _dafny.toString(this.cf92) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf93 === other.cf93;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf94 === other.cf94 && _dafny.areEqual(this.cf95, other.cf95) && this.cf96 === other.cf96;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf92 === other.cf92;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D19.create_DC49(false);
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
    static create_DC52(cf98) {
      let $dt = new D20(0);
      $dt.cf98 = cf98;
      return $dt;
    }
    static create_DC51(cf97) {
      let $dt = new D20(1);
      $dt.cf97 = cf97;
      return $dt;
    }
    get is_DC52() { return this.$tag === 0; }
    get is_DC51() { return this.$tag === 1; }
    get dtor_cf98() { return this.cf98; }
    get dtor_cf97() { return this.cf97; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC52" + "(" + _dafny.toString(this.cf98) + ")";
      } else if (this.$tag === 1) {
        return "D20.DC51" + "(" + _dafny.toString(this.cf97) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf98, other.cf98);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf97 === other.cf97;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC52(_dafny.ZERO);
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
    static create_DC54(cf100, cf101, cf102) {
      let $dt = new D21(0);
      $dt.cf100 = cf100;
      $dt.cf101 = cf101;
      $dt.cf102 = cf102;
      return $dt;
    }
    static create_DC53(cf99) {
      let $dt = new D21(1);
      $dt.cf99 = cf99;
      return $dt;
    }
    static create_DC55(cf103) {
      let $dt = new D21(2);
      $dt.cf103 = cf103;
      return $dt;
    }
    get is_DC54() { return this.$tag === 0; }
    get is_DC53() { return this.$tag === 1; }
    get is_DC55() { return this.$tag === 2; }
    get dtor_cf100() { return this.cf100; }
    get dtor_cf101() { return this.cf101; }
    get dtor_cf102() { return this.cf102; }
    get dtor_cf99() { return this.cf99; }
    get dtor_cf103() { return this.cf103; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC54" + "(" + _dafny.toString(this.cf100) + ", " + _dafny.toString(this.cf101) + ", " + this.cf102.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D21.DC53" + "(" + _dafny.toString(this.cf99) + ")";
      } else if (this.$tag === 2) {
        return "D21.DC55" + "(" + _dafny.toString(this.cf103) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf100, other.cf100) && _dafny.areEqual(this.cf101, other.cf101) && _dafny.areEqual(this.cf102, other.cf102);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf99 === other.cf99;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf103, other.cf103);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D21.create_DC54(new _dafny.CodePoint('D'.codePointAt(0)), _dafny.ZERO, _dafny.Seq.UnicodeFromString(""));
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
    static create_DC57() {
      let $dt = new D22(0);
      return $dt;
    }
    static create_DC56(cf104) {
      let $dt = new D22(1);
      $dt.cf104 = cf104;
      return $dt;
    }
    get is_DC57() { return this.$tag === 0; }
    get is_DC56() { return this.$tag === 1; }
    get dtor_cf104() { return this.cf104; }
    toString() {
      if (this.$tag === 0) {
        return "D22.DC57";
      } else if (this.$tag === 1) {
        return "D22.DC56" + "(" + _dafny.toString(this.cf104) + ")";
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
        return other.$tag === 1 && this.cf104 === other.cf104;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D22.create_DC57();
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
    static create_DC59(cf106, cf107, cf108, cf109) {
      let $dt = new D23(0);
      $dt.cf106 = cf106;
      $dt.cf107 = cf107;
      $dt.cf108 = cf108;
      $dt.cf109 = cf109;
      return $dt;
    }
    static create_DC58(cf105) {
      let $dt = new D23(1);
      $dt.cf105 = cf105;
      return $dt;
    }
    get is_DC59() { return this.$tag === 0; }
    get is_DC58() { return this.$tag === 1; }
    get dtor_cf106() { return this.cf106; }
    get dtor_cf107() { return this.cf107; }
    get dtor_cf108() { return this.cf108; }
    get dtor_cf109() { return this.cf109; }
    get dtor_cf105() { return this.cf105; }
    toString() {
      if (this.$tag === 0) {
        return "D23.DC59" + "(" + _dafny.toString(this.cf106) + ", " + _dafny.toString(this.cf107) + ", " + _dafny.toString(this.cf108) + ", " + _dafny.toString(this.cf109) + ")";
      } else if (this.$tag === 1) {
        return "D23.DC58" + "(" + _dafny.toString(this.cf105) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf106, other.cf106) && this.cf107 === other.cf107 && _dafny.areEqual(this.cf108, other.cf108) && this.cf109 === other.cf109;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf105 === other.cf105;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D23.create_DC59(_module.D6.Default(), false, _dafny.ZERO, []);
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
    static create_DC60(cf110) {
      let $dt = new D24(0);
      $dt.cf110 = cf110;
      return $dt;
    }
    get is_DC60() { return this.$tag === 0; }
    get dtor_cf110() { return this.cf110; }
    toString() {
      if (this.$tag === 0) {
        return "D24.DC60" + "(" + _dafny.toString(this.cf110) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf110, other.cf110);
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
          return D24.Default();
        }
      };
    }
  }

  $module.D25 = class D25 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC62() {
      let $dt = new D25(0);
      return $dt;
    }
    static create_DC63(cf112, cf113) {
      let $dt = new D25(1);
      $dt.cf112 = cf112;
      $dt.cf113 = cf113;
      return $dt;
    }
    static create_DC61(cf111) {
      let $dt = new D25(2);
      $dt.cf111 = cf111;
      return $dt;
    }
    get is_DC62() { return this.$tag === 0; }
    get is_DC63() { return this.$tag === 1; }
    get is_DC61() { return this.$tag === 2; }
    get dtor_cf112() { return this.cf112; }
    get dtor_cf113() { return this.cf113; }
    get dtor_cf111() { return this.cf111; }
    toString() {
      if (this.$tag === 0) {
        return "D25.DC62";
      } else if (this.$tag === 1) {
        return "D25.DC63" + "(" + this.cf112.toVerbatimString(true) + ", " + _dafny.toString(this.cf113) + ")";
      } else if (this.$tag === 2) {
        return "D25.DC61" + "(" + _dafny.toString(this.cf111) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf112, other.cf112) && _dafny.areEqual(this.cf113, other.cf113);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf111 === other.cf111;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D25.create_DC62();
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
    static create_DC65(cf115, cf116, cf117) {
      let $dt = new D26(0);
      $dt.cf115 = cf115;
      $dt.cf116 = cf116;
      $dt.cf117 = cf117;
      return $dt;
    }
    static create_DC64(cf114) {
      let $dt = new D26(1);
      $dt.cf114 = cf114;
      return $dt;
    }
    get is_DC65() { return this.$tag === 0; }
    get is_DC64() { return this.$tag === 1; }
    get dtor_cf115() { return this.cf115; }
    get dtor_cf116() { return this.cf116; }
    get dtor_cf117() { return this.cf117; }
    get dtor_cf114() { return this.cf114; }
    toString() {
      if (this.$tag === 0) {
        return "D26.DC65" + "(" + _dafny.toString(this.cf115) + ", " + _dafny.toString(this.cf116) + ", " + _dafny.toString(this.cf117) + ")";
      } else if (this.$tag === 1) {
        return "D26.DC64" + "(" + _dafny.toString(this.cf114) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf115 === other.cf115 && this.cf116 === other.cf116 && this.cf117 === other.cf117;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf114 === other.cf114;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D26.create_DC65(false, [], false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D26.Default();
        }
      };
    }
  }

  $module.D27 = class D27 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC66(cf118) {
      let $dt = new D27(0);
      $dt.cf118 = cf118;
      return $dt;
    }
    get is_DC66() { return this.$tag === 0; }
    get dtor_cf118() { return this.cf118; }
    toString() {
      if (this.$tag === 0) {
        return "D27.DC66" + "(" + _dafny.toString(this.cf118) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf118, other.cf118);
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
          return D27.Default();
        }
      };
    }
  }

  $module.D28 = class D28 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC68(cf120, cf121, cf122) {
      let $dt = new D28(0);
      $dt.cf120 = cf120;
      $dt.cf121 = cf121;
      $dt.cf122 = cf122;
      return $dt;
    }
    static create_DC69() {
      let $dt = new D28(1);
      return $dt;
    }
    static create_DC70(cf123, cf124) {
      let $dt = new D28(2);
      $dt.cf123 = cf123;
      $dt.cf124 = cf124;
      return $dt;
    }
    static create_DC67(cf119) {
      let $dt = new D28(3);
      $dt.cf119 = cf119;
      return $dt;
    }
    static create_DC71(cf125) {
      let $dt = new D28(4);
      $dt.cf125 = cf125;
      return $dt;
    }
    get is_DC68() { return this.$tag === 0; }
    get is_DC69() { return this.$tag === 1; }
    get is_DC70() { return this.$tag === 2; }
    get is_DC67() { return this.$tag === 3; }
    get is_DC71() { return this.$tag === 4; }
    get dtor_cf120() { return this.cf120; }
    get dtor_cf121() { return this.cf121; }
    get dtor_cf122() { return this.cf122; }
    get dtor_cf123() { return this.cf123; }
    get dtor_cf124() { return this.cf124; }
    get dtor_cf119() { return this.cf119; }
    get dtor_cf125() { return this.cf125; }
    toString() {
      if (this.$tag === 0) {
        return "D28.DC68" + "(" + _dafny.toString(this.cf120) + ", " + _dafny.toString(this.cf121) + ", " + _dafny.toString(this.cf122) + ")";
      } else if (this.$tag === 1) {
        return "D28.DC69";
      } else if (this.$tag === 2) {
        return "D28.DC70" + "(" + _dafny.toString(this.cf123) + ", " + _dafny.toString(this.cf124) + ")";
      } else if (this.$tag === 3) {
        return "D28.DC67" + "(" + _dafny.toString(this.cf119) + ")";
      } else if (this.$tag === 4) {
        return "D28.DC71" + "(" + _dafny.toString(this.cf125) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf120, other.cf120) && this.cf121 === other.cf121 && _dafny.areEqual(this.cf122, other.cf122);
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf123 === other.cf123 && this.cf124 === other.cf124;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf119, other.cf119);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf125, other.cf125);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D28.create_DC68(_dafny.ZERO, false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D28.Default();
        }
      };
    }
  }

  $module.D29 = class D29 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC72(cf126) {
      let $dt = new D29(0);
      $dt.cf126 = cf126;
      return $dt;
    }
    get is_DC72() { return this.$tag === 0; }
    get dtor_cf126() { return this.cf126; }
    toString() {
      if (this.$tag === 0) {
        return "D29.DC72" + "(" + _dafny.toString(this.cf126) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf126, other.cf126);
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
          return D29.Default();
        }
      };
    }
  }

  $module.D30 = class D30 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC74(cf128) {
      let $dt = new D30(0);
      $dt.cf128 = cf128;
      return $dt;
    }
    static create_DC75(cf129) {
      let $dt = new D30(1);
      $dt.cf129 = cf129;
      return $dt;
    }
    static create_DC73(cf127) {
      let $dt = new D30(2);
      $dt.cf127 = cf127;
      return $dt;
    }
    static create_DC76(cf130) {
      let $dt = new D30(3);
      $dt.cf130 = cf130;
      return $dt;
    }
    get is_DC74() { return this.$tag === 0; }
    get is_DC75() { return this.$tag === 1; }
    get is_DC73() { return this.$tag === 2; }
    get is_DC76() { return this.$tag === 3; }
    get dtor_cf128() { return this.cf128; }
    get dtor_cf129() { return this.cf129; }
    get dtor_cf127() { return this.cf127; }
    get dtor_cf130() { return this.cf130; }
    toString() {
      if (this.$tag === 0) {
        return "D30.DC74" + "(" + this.cf128.toVerbatimString(true) + ")";
      } else if (this.$tag === 1) {
        return "D30.DC75" + "(" + _dafny.toString(this.cf129) + ")";
      } else if (this.$tag === 2) {
        return "D30.DC73" + "(" + _dafny.toString(this.cf127) + ")";
      } else if (this.$tag === 3) {
        return "D30.DC76" + "(" + _dafny.toString(this.cf130) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf128, other.cf128);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf129 === other.cf129;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf127 === other.cf127;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf130, other.cf130);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D30.create_DC74(_dafny.Seq.UnicodeFromString(""));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D30.Default();
        }
      };
    }
  }

  $module.D31 = class D31 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC77(cf131) {
      let $dt = new D31(0);
      $dt.cf131 = cf131;
      return $dt;
    }
    get is_DC77() { return this.$tag === 0; }
    get dtor_cf131() { return this.cf131; }
    toString() {
      if (this.$tag === 0) {
        return "D31.DC77" + "(" + _dafny.toString(this.cf131) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf131, other.cf131);
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
          return D31.Default();
        }
      };
    }
  }

  $module.D32 = class D32 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC79(cf133, cf134, cf135) {
      let $dt = new D32(0);
      $dt.cf133 = cf133;
      $dt.cf134 = cf134;
      $dt.cf135 = cf135;
      return $dt;
    }
    static create_DC80(cf136, cf137, cf138) {
      let $dt = new D32(1);
      $dt.cf136 = cf136;
      $dt.cf137 = cf137;
      $dt.cf138 = cf138;
      return $dt;
    }
    static create_DC78(cf132) {
      let $dt = new D32(2);
      $dt.cf132 = cf132;
      return $dt;
    }
    get is_DC79() { return this.$tag === 0; }
    get is_DC80() { return this.$tag === 1; }
    get is_DC78() { return this.$tag === 2; }
    get dtor_cf133() { return this.cf133; }
    get dtor_cf134() { return this.cf134; }
    get dtor_cf135() { return this.cf135; }
    get dtor_cf136() { return this.cf136; }
    get dtor_cf137() { return this.cf137; }
    get dtor_cf138() { return this.cf138; }
    get dtor_cf132() { return this.cf132; }
    toString() {
      if (this.$tag === 0) {
        return "D32.DC79" + "(" + _dafny.toString(this.cf133) + ", " + this.cf134.toVerbatimString(true) + ", " + _dafny.toString(this.cf135) + ")";
      } else if (this.$tag === 1) {
        return "D32.DC80" + "(" + _dafny.toString(this.cf136) + ", " + _dafny.toString(this.cf137) + ", " + _dafny.toString(this.cf138) + ")";
      } else if (this.$tag === 2) {
        return "D32.DC78" + "(" + _dafny.toString(this.cf132) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf133 === other.cf133 && _dafny.areEqual(this.cf134, other.cf134) && _dafny.areEqual(this.cf135, other.cf135);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf136, other.cf136) && _dafny.areEqual(this.cf137, other.cf137) && _dafny.areEqual(this.cf138, other.cf138);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf132 === other.cf132;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D32.create_DC79(false, _dafny.Seq.UnicodeFromString(""), _dafny.Seq.of());
    }
    static Rtd() {
      return class {
        static get Default() {
          return D32.Default();
        }
      };
    }
  }

  $module.D33 = class D33 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC82(cf140, cf141, cf142, cf143) {
      let $dt = new D33(0);
      $dt.cf140 = cf140;
      $dt.cf141 = cf141;
      $dt.cf142 = cf142;
      $dt.cf143 = cf143;
      return $dt;
    }
    static create_DC81(cf139) {
      let $dt = new D33(1);
      $dt.cf139 = cf139;
      return $dt;
    }
    static create_DC83(cf144) {
      let $dt = new D33(2);
      $dt.cf144 = cf144;
      return $dt;
    }
    get is_DC82() { return this.$tag === 0; }
    get is_DC81() { return this.$tag === 1; }
    get is_DC83() { return this.$tag === 2; }
    get dtor_cf140() { return this.cf140; }
    get dtor_cf141() { return this.cf141; }
    get dtor_cf142() { return this.cf142; }
    get dtor_cf143() { return this.cf143; }
    get dtor_cf139() { return this.cf139; }
    get dtor_cf144() { return this.cf144; }
    toString() {
      if (this.$tag === 0) {
        return "D33.DC82" + "(" + _dafny.toString(this.cf140) + ", " + _dafny.toString(this.cf141) + ", " + _dafny.toString(this.cf142) + ", " + _dafny.toString(this.cf143) + ")";
      } else if (this.$tag === 1) {
        return "D33.DC81" + "(" + _dafny.toString(this.cf139) + ")";
      } else if (this.$tag === 2) {
        return "D33.DC83" + "(" + _dafny.toString(this.cf144) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf140 === other.cf140 && this.cf141 === other.cf141 && this.cf142 === other.cf142 && _dafny.areEqual(this.cf143, other.cf143);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf139, other.cf139);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf144, other.cf144);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D33.create_DC82(false, false, null, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D33.Default();
        }
      };
    }
  }

  $module.D34 = class D34 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC85(cf146) {
      let $dt = new D34(0);
      $dt.cf146 = cf146;
      return $dt;
    }
    static create_DC86(cf147, cf148) {
      let $dt = new D34(1);
      $dt.cf147 = cf147;
      $dt.cf148 = cf148;
      return $dt;
    }
    static create_DC84(cf145) {
      let $dt = new D34(2);
      $dt.cf145 = cf145;
      return $dt;
    }
    get is_DC85() { return this.$tag === 0; }
    get is_DC86() { return this.$tag === 1; }
    get is_DC84() { return this.$tag === 2; }
    get dtor_cf146() { return this.cf146; }
    get dtor_cf147() { return this.cf147; }
    get dtor_cf148() { return this.cf148; }
    get dtor_cf145() { return this.cf145; }
    toString() {
      if (this.$tag === 0) {
        return "D34.DC85" + "(" + _dafny.toString(this.cf146) + ")";
      } else if (this.$tag === 1) {
        return "D34.DC86" + "(" + _dafny.toString(this.cf147) + ", " + _dafny.toString(this.cf148) + ")";
      } else if (this.$tag === 2) {
        return "D34.DC84" + "(" + _dafny.toString(this.cf145) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf146, other.cf146);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf147 === other.cf147 && _dafny.areEqual(this.cf148, other.cf148);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf145, other.cf145);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D34.create_DC85(new _dafny.CodePoint('D'.codePointAt(0)));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D34.Default();
        }
      };
    }
  }

  $module.D35 = class D35 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC88() {
      let $dt = new D35(0);
      return $dt;
    }
    static create_DC89(cf150, cf151, cf152) {
      let $dt = new D35(1);
      $dt.cf150 = cf150;
      $dt.cf151 = cf151;
      $dt.cf152 = cf152;
      return $dt;
    }
    static create_DC87(cf149) {
      let $dt = new D35(2);
      $dt.cf149 = cf149;
      return $dt;
    }
    get is_DC88() { return this.$tag === 0; }
    get is_DC89() { return this.$tag === 1; }
    get is_DC87() { return this.$tag === 2; }
    get dtor_cf150() { return this.cf150; }
    get dtor_cf151() { return this.cf151; }
    get dtor_cf152() { return this.cf152; }
    get dtor_cf149() { return this.cf149; }
    toString() {
      if (this.$tag === 0) {
        return "D35.DC88";
      } else if (this.$tag === 1) {
        return "D35.DC89" + "(" + _dafny.toString(this.cf150) + ", " + _dafny.toString(this.cf151) + ", " + _dafny.toString(this.cf152) + ")";
      } else if (this.$tag === 2) {
        return "D35.DC87" + "(" + _dafny.toString(this.cf149) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf150, other.cf150) && _dafny.areEqual(this.cf151, other.cf151) && _dafny.areEqual(this.cf152, other.cf152);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf149, other.cf149);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D35.create_DC88();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D35.Default();
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
      this.f4 = _dafny.Map.Empty;
      this.f8 = false;
      this._f0 = false;
      this._f1 = false;
      this._f2 = [];
      this._f3 = false;
      this._f5 = _dafny.ZERO;
      this._f6 = false;
      this._f7 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this).f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this).f8 = f8;
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
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_module.__default.safeDivisionInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.UnicodeFromString("aiohhcmxm"))).length), new BigNumber(-442))).isLessThanOrEqualTo((new BigNumber(-952)).minus(new BigNumber(834)));
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f12 = false;
      this._f13 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor(f12, f13) {
      let _this = this;
      (_this)._f12 = f12;
      (_this)._f13 = f13;
      return;
    }
    fm32(p0, p1, globalState) {
      let _this = this;
      return (_this).f12;
    };
    m0(globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = false;
      let r2 = _dafny.Seq.of();
      let r3 = _dafny.Seq.UnicodeFromString("");
      let _459_v0;
      let _nw45 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
      _459_v0 = _nw45;
      let _460_v1;
      _460_v1 = new BigNumber(261);
      let _index50 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_459_v0).length));
      (_459_v0)[_index50] = _module.__default.safeDivisionInt(_460_v1, _460_v1);
      let _index51 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_459_v0).length));
      (_459_v0)[_index51] = _460_v1;
      (globalState).f8 = (_this).f12;
      let _461_v2;
      _461_v2 = _dafny.Map.Empty.slice().updateUnsafe(_460_v1,(_this).f12);
      _461_v2 = (_461_v2).update((_module.__default.fm33((_this).f12, (_this).f12, _460_v1, true, globalState)).multipliedBy((_459_v0)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_459_v0).length))]), ((_this).f12) === (!((_this).f12)));
      let _462_v3;
      _462_v3 = new _dafny.CodePoint('m'.codePointAt(0));
      if ((_this).fm32((_459_v0)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_459_v0).length))], _462_v3, globalState)) {
        r0 = _dafny.Seq.UnicodeFromString("tsoaubs");
        let _463_v4;
        let _nw46 = new _module.C0();
        _nw46.__ctor();
        _463_v4 = _nw46;
        if (!((_459_v0)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_459_v0).length))]).isEqualTo((_459_v0)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_459_v0).length))])) {
          (globalState).f8 = !((_this).f12);
          let _464_v5;
          let _nw47 = Array((new BigNumber(27)).toNumber()).fill(false);
          _464_v5 = _nw47;
          let _index52 = _module.__default.safeIndex(new BigNumber(180), new BigNumber((_464_v5).length));
          (_464_v5)[_index52] = (((_461_v2).contains((_459_v0)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_459_v0).length))])) ? ((_461_v2).get((_459_v0)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_459_v0).length))])) : (false));
          let _465_v6;
          _465_v6 = _dafny.MultiSet.fromElements(_dafny.Seq.of((_this).f12));
          let _index53 = _module.__default.safeIndex(new BigNumber(180), new BigNumber((_464_v5).length));
          (_464_v5)[_index53] = (_465_v6).IsProperSubsetOf(_465_v6);
          let _nw48 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
          _459_v0 = _nw48;
          (globalState).f8 = (_464_v5)[_module.__default.safeIndex(new BigNumber(180), new BigNumber((_464_v5).length))];
          let _466_v7;
          let _nw49 = Array((_dafny.ONE).toNumber()).fill(_dafny.Set.Empty);
          _466_v7 = _nw49;
          _466_v7 = _466_v7;
        } else {
          let _index54 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_459_v0).length));
          (_459_v0)[_index54] = new BigNumber(606);
          let _467_v8;
          let _nw50 = new _module.C0();
          _nw50.__ctor();
          _467_v8 = _nw50;
          let _468_v9;
          _468_v9 = _dafny.MultiSet.fromElements(!((_this).f12));
          _460_v1 = _module.__default.fm33((_this).f12, !((_467_v8).fm5((_this).f12, _460_v1, (_459_v0)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_459_v0).length))], _468_v9, globalState)), _module.__default.safeModuloInt(_460_v1, _460_v1), (_this).f12, globalState);
          let _469_v10;
          _469_v10 = _dafny.Map.Empty.slice().updateUnsafe(_462_v3,(_this).f13);
          let _470_v11;
          _470_v11 = _dafny.Map.Empty.slice().updateUnsafe((_459_v0)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_459_v0).length))],_460_v1);
          let _rhs58 = (_469_v10).Merge((_469_v10).update(((_this).f13)[_module.__default.safeIndex(new BigNumber((_470_v11).length), new BigNumber(((_this).f13).length))], (_this).f13));
          let _rhs59 = (_460_v1).plus(_460_v1);
          let _rhs60 = (_this).f12;
          let _lhs44 = globalState;
          _469_v10 = _rhs58;
          _460_v1 = _rhs59;
          _lhs44.f8 = _rhs60;
          let _471_v12;
          let _init5 = function (_472_i0) {
            return (_this).f12;
          };
          let _nw51 = Array((new BigNumber(5)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw51.length); _i0_5++) {
            _nw51[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _471_v12 = _nw51;
          let _473_v13;
          _473_v13 = _dafny.MultiSet.fromElements(_471_v12);
          _473_v13 = ((_473_v13).Union(_473_v13)).Intersect(_473_v13);
        }
        let _474_v14;
        _474_v14 = _dafny.MultiSet.fromElements(new BigNumber(367), _module.__default.fm33((_this).f12, (_this).f12, (_459_v0)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_459_v0).length))], (_this).f12, globalState));
        let _475_v15;
        let _nw52 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _475_v15 = _nw52;
        let _index55 = _module.__default.safeIndex(new BigNumber(272), new BigNumber((_475_v15).length));
        (_475_v15)[_index55] = _dafny.Seq.Concat((_this).f13, _dafny.Seq.UnicodeFromString("ejfjtu"));
        let _index56 = _module.__default.safeIndex(new BigNumber(272), new BigNumber((_475_v15).length));
        let _rhs61 = ((_474_v14).Intersect(_dafny.MultiSet.fromElements(_460_v1, _460_v1))).Intersect(_474_v14);
        let _rhs62 = (_459_v0)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_459_v0).length))];
        let _rhs63 = ((true) ? ((_this).f13) : ((((_this).f12) ? ((_this).f13) : ((_this).f13))));
        let _lhs45 = _475_v15;
        let _lhs46 = _module.__default.safeIndex(new BigNumber(272), new BigNumber((_475_v15).length));
        _474_v14 = _rhs61;
        _460_v1 = _rhs62;
        _lhs45[_lhs46] = _rhs63;
        r1 = false;
      } else {
        let _476_v16;
        _476_v16 = _module.D3.create_DC7(_459_v0);
        _476_v16 = ((!((_this).f12) || ((_this).f12)) ? ((((_this).f12) ? (_module.D3.create_DC7(_459_v0)) : (_476_v16))) : (_476_v16));
        let _477_v17;
        _477_v17 = _dafny.Set.fromElements(_460_v1, _module.__default.fm33((_this).f12, (_this).f12, (_459_v0)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_459_v0).length))], (_this).f12, globalState));
        let _478_v18;
        _478_v18 = _dafny.Map.Empty.slice().updateUnsafe((_477_v17).IsDisjointFrom(_477_v17),(_460_v1).minus((_459_v0)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_459_v0).length))]));
        _478_v18 = (_478_v18).update((_this).f12, (_459_v0)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_459_v0).length))]);
        let _index57 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_459_v0).length));
        (_459_v0)[_index57] = (_dafny.ZERO).minus(_460_v1);
        let _479_v19;
        let _nw53 = Array((new BigNumber(6)).toNumber()).fill(false);
        _479_v19 = _nw53;
        let _index58 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_479_v19).length));
        (_479_v19)[_index58] = (_this).f12;
        let _index59 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_479_v19).length));
        (_479_v19)[_index59] = (_this).f12;
        let _480_v20;
        let _nw54 = Array((new BigNumber(4)).toNumber());
        _nw54[(_dafny.ZERO).toNumber()] = _459_v0;
        _nw54[(_dafny.ONE).toNumber()] = _459_v0;
        _nw54[(new BigNumber(2)).toNumber()] = _459_v0;
        _nw54[(new BigNumber(3)).toNumber()] = _459_v0;
        _480_v20 = _nw54;
        let _nw55 = Array((new BigNumber(28)).toNumber());
        _nw55[(_dafny.ZERO).toNumber()] = (_476_v16).dtor_cf11;
        _nw55[(_dafny.ONE).toNumber()] = _459_v0;
        _nw55[(new BigNumber(2)).toNumber()] = _459_v0;
        _nw55[(new BigNumber(3)).toNumber()] = _459_v0;
        _nw55[(new BigNumber(4)).toNumber()] = _459_v0;
        _nw55[(new BigNumber(5)).toNumber()] = _459_v0;
        _nw55[(new BigNumber(6)).toNumber()] = _459_v0;
        _nw55[(new BigNumber(7)).toNumber()] = _459_v0;
        _nw55[(new BigNumber(8)).toNumber()] = _459_v0;
        _nw55[(new BigNumber(9)).toNumber()] = _459_v0;
        _nw55[(new BigNumber(10)).toNumber()] = _459_v0;
        _nw55[(new BigNumber(11)).toNumber()] = _459_v0;
        _nw55[(new BigNumber(12)).toNumber()] = _459_v0;
        _nw55[(new BigNumber(13)).toNumber()] = _459_v0;
        _nw55[(new BigNumber(14)).toNumber()] = (((_479_v19)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_479_v19).length))]) ? (_459_v0) : (_459_v0));
        _nw55[(new BigNumber(15)).toNumber()] = _459_v0;
        _nw55[(new BigNumber(16)).toNumber()] = _459_v0;
        _nw55[(new BigNumber(17)).toNumber()] = _459_v0;
        _nw55[(new BigNumber(18)).toNumber()] = (((_479_v19)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_479_v19).length))]) ? (_459_v0) : (_459_v0));
        _nw55[(new BigNumber(19)).toNumber()] = _459_v0;
        _nw55[(new BigNumber(20)).toNumber()] = (((_479_v19)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_479_v19).length))]) ? (_459_v0) : (_459_v0));
        _nw55[(new BigNumber(21)).toNumber()] = _459_v0;
        _nw55[(new BigNumber(22)).toNumber()] = _459_v0;
        _nw55[(new BigNumber(23)).toNumber()] = _459_v0;
        _nw55[(new BigNumber(24)).toNumber()] = _459_v0;
        _nw55[(new BigNumber(25)).toNumber()] = _459_v0;
        _nw55[(new BigNumber(26)).toNumber()] = _459_v0;
        _nw55[(new BigNumber(27)).toNumber()] = (_module.D3.create_DC7(_459_v0)).dtor_cf11;
        _480_v20 = _nw55;
      }
      let _index60 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_459_v0).length));
      (_459_v0)[_index60] = _460_v1;
      let _index61 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_459_v0).length));
      (_459_v0)[_index61] = _460_v1;
      r0 = _dafny.Seq.update(_dafny.Seq.Concat((_this).f13, _dafny.Seq.Concat((_this).f13, (_this).f13)), _module.__default.safeIndex(new BigNumber(-935), new BigNumber((_dafny.Seq.Concat((_this).f13, _dafny.Seq.Concat((_this).f13, (_this).f13))).length)), _462_v3);
      r1 = (((_this).f12) ? ((_this).f12) : ((_this).f12));
      let _481_v21;
      _481_v21 = _dafny.Set.fromElements((_this).fm32(_460_v1, _462_v3, globalState), (_this).f12, false);
      let _482_v22;
      _482_v22 = _dafny.Seq.of(_481_v21);
      let _483_v23;
      _483_v23 = _dafny.Map.Empty.slice().updateUnsafe(_459_v0,(_459_v0)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_459_v0).length))]);
      r2 = _dafny.Seq.Concat(_dafny.Seq.of(_481_v21, (_482_v22)[_module.__default.safeIndex(_460_v1, new BigNumber((_482_v22).length))]), _dafny.Seq.of(_481_v21, _481_v21, _module.__default.fm34(new BigNumber((_483_v23).length), (_this).f12, (_this).f12, (_this).f12, globalState), _dafny.Set.fromElements(false)));
      r3 = _dafny.Seq.Concat((_this).f13, (_this).f13);
      return [r0, r1, r2, r3];
    }
    m1(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = false;
      let _484_v0;
      _484_v0 = new BigNumber(776);
      let _485_v1;
      _485_v1 = _dafny.Map.Empty.slice().updateUnsafe(_484_v0,(_this).f12);
      let _486_v2;
      _486_v2 = _module.D0.create_DC1(_484_v0, (_this).fm32(new BigNumber((_485_v1).length), p0, globalState), (_dafny.ZERO).minus(_484_v0));
      let _487_v4;
      _487_v4 = _dafny.Map.Empty.slice().updateUnsafe((_484_v0).isEqualTo((_486_v2).dtor_cf3),function () {
        let _coll29 = new _dafny.Map();
        for (const _compr_29 of _dafny.IntegerRange(new BigNumber(7), new BigNumber(695))) {
          let _488_v3 = _compr_29;
          if (((new BigNumber(7)).isLessThanOrEqualTo(_488_v3)) && ((_488_v3).isLessThan(new BigNumber(695)))) {
            _coll29.push([(_488_v3).multipliedBy(_484_v0),_484_v0]);
          }
        }
        return _coll29;
      }());
      let _489_v5;
      _489_v5 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(620)), function (_490_i0) {
        return new BigNumber(674);
      })).length),new BigNumber(54));
      _487_v4 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).fm32(_484_v0, p0, globalState)),_489_v5);
      let _491_v6;
      let _nw56 = new _module.C0();
      _nw56.__ctor();
      _491_v6 = _nw56;
      let _492_v7;
      _492_v7 = new _dafny.CodePoint('i'.codePointAt(0));
      let _pat_let_tv6 = p0;
      let _pat_let_tv7 = p0;
      let _pat_let_tv8 = p0;
      _492_v7 = function (_source12) {
        if (_source12.is_DC8) {
          let _493___mcc_h0 = (_source12).cf12;
          let _494___mcc_h1 = (_source12).cf13;
          let _495_cf13 = _494___mcc_h1;
          let _496_cf12 = _493___mcc_h0;
          if (_495_cf13) {
            return new _dafny.CodePoint('o'.codePointAt(0));
          } else {
            return _pat_let_tv6;
          }
        } else if (_source12.is_DC7) {
          let _497___mcc_h2 = (_source12).cf11;
          let _498_cf11 = _497___mcc_h2;
          return _pat_let_tv7;
        } else {
          let _499___mcc_h3 = (_source12).cf14;
          let _500_cf14 = _499___mcc_h3;
          return (_pat_let_tv8);
        }
      }(_module.__default.fm35((_this).f12, globalState));
      let _501_v8;
      _501_v8 = _dafny.MultiSet.fromElements((_this).f12);
      let _502_v9;
      _502_v9 = _module.D1.create_DC4((_491_v6).fm5((_this).f12, _484_v0, _484_v0, _501_v8, globalState), (_this).f12, (((_this).f12) ? (new BigNumber((_501_v8).cardinality())) : (_484_v0)));
      let _source13 = _502_v9;
      if (_source13.is_DC3) {
        let _503___mcc_h4 = (_source13).cf5;
        let _504_cf5 = _503___mcc_h4;
        let _505_v10;
        _505_v10 = _dafny.MultiSet.fromElements(_484_v0);
        let _506_v11;
        _506_v11 = _dafny.Map.Empty.slice().updateUnsafe((_this).f12,_505_v10);
        let _507_v12;
        _507_v12 = _dafny.Seq.of(_484_v0);
        _506_v11 = (_506_v11).update(!((_this).f12) || (true), (_dafny.MultiSet.FromArray(_507_v12)).Difference(_505_v10));
        (globalState).f8 = !(new BigNumber(924)).isEqualTo(new BigNumber((_501_v8).cardinality()));
        let _508_v13;
        let _nw57 = Array((new BigNumber(16)).toNumber());
        _nw57[(_dafny.ZERO).toNumber()] = (_this).f12;
        _nw57[(_dafny.ONE).toNumber()] = (_this).f12;
        _nw57[(new BigNumber(2)).toNumber()] = (_this).f12;
        _nw57[(new BigNumber(3)).toNumber()] = (_this).f12;
        _nw57[(new BigNumber(4)).toNumber()] = (_this).f12;
        _nw57[(new BigNumber(5)).toNumber()] = (new BigNumber(((_this).f13).length)).isLessThan(_484_v0);
        _nw57[(new BigNumber(6)).toNumber()] = (_this).f12;
        _nw57[(new BigNumber(7)).toNumber()] = false;
        _nw57[(new BigNumber(8)).toNumber()] = (((_this).f12) ? ((_this).f12) : ((_this).f12));
        _nw57[(new BigNumber(9)).toNumber()] = (_this).f12;
        _nw57[(new BigNumber(10)).toNumber()] = (_this).f12;
        _nw57[(new BigNumber(11)).toNumber()] = (_this).f12;
        _nw57[(new BigNumber(12)).toNumber()] = (_this).f12;
        _nw57[(new BigNumber(13)).toNumber()] = !((((_this).f12) ? ((_this).f12) : ((_this).f12)));
        _nw57[(new BigNumber(14)).toNumber()] = !(!((_491_v6).fm5((_this).f12, _484_v0, _484_v0, _501_v8, globalState)) || ((_this).f12));
        _nw57[(new BigNumber(15)).toNumber()] = ((_this).f12) || ((_this).f12);
        _508_v13 = _nw57;
        let _index62 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_508_v13).length));
        (_508_v13)[_index62] = (_this).f12;
        let _index63 = _module.__default.safeIndex(new BigNumber(257), new BigNumber((_508_v13).length));
        (_508_v13)[_index63] = !((_this).f12);
        let _509_v14;
        _509_v14 = _dafny.Seq.of((_508_v13)[_module.__default.safeIndex(new BigNumber(257), new BigNumber((_508_v13).length))], (_this).f12);
        _492_v7 = (_module.D4.create_DC12(true, new BigNumber((_509_v14).length), _module.__default.fm36(globalState), new BigNumber(739), _484_v0)).dtor_cf21;
      } else if (_source13.is_DC4) {
        let _510___mcc_h5 = (_source13).cf6;
        let _511___mcc_h6 = (_source13).cf7;
        let _512___mcc_h7 = (_source13).cf8;
        let _513_cf8 = _512___mcc_h7;
        let _514_cf7 = _511___mcc_h6;
        let _515_cf6 = _510___mcc_h5;
        _513_cf8 = new BigNumber(-389);
        let _516_v15;
        _516_v15 = _dafny.Set.fromElements(true, _515_cf6);
        _514_cf7 = (_516_v15).contains(_515_cf6);
        _491_v6 = _491_v6;
        let _517_v16;
        _517_v16 = _module.D0.create_DC0((_this).f12);
        let _518_v17;
        _518_v17 = _dafny.Seq.of(_514_cf7, (_517_v16).dtor_cf0);
        if ((_515_cf6) || ((_dafny.MultiSet.FromArray(_518_v17)).IsSubsetOf(_501_v8))) {
          r1 = ((_516_v15).Intersect(_dafny.Set.fromElements(_514_cf7))).IsSubsetOf(_516_v15);
          (globalState).f8 = _515_cf6;
          let _519_v18;
          _519_v18 = _dafny.Seq.of(_484_v0, _484_v0);
          _484_v0 = (_513_cf8).minus(new BigNumber((_519_v18).length));
          let _520_v19;
          let _nw58 = new _module.C0();
          _nw58.__ctor();
          _520_v19 = _nw58;
          let _rhs64 = _514_cf7;
          let _rhs65 = _514_cf7;
          _515_cf6 = _rhs64;
          r1 = _rhs65;
        } else {
          let _521_v20;
          _521_v20 = _dafny.Seq.UnicodeFromString("xrh");
          _521_v20 = _521_v20;
          _513_cf8 = ((_dafny.ZERO).minus(_513_cf8)).minus((((_489_v5).contains(_484_v0)) ? ((_489_v5).get(_484_v0)) : (_513_cf8)));
          let _522_v21;
          let _nw59 = Array((new BigNumber(3)).toNumber());
          _nw59[(_dafny.ZERO).toNumber()] = _491_v6;
          _nw59[(_dafny.ONE).toNumber()] = _491_v6;
          _nw59[(new BigNumber(2)).toNumber()] = _491_v6;
          _522_v21 = _nw59;
          let _index64 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_522_v21).length));
          (_522_v21)[_index64] = _491_v6;
          let _index65 = _module.__default.safeIndex(new BigNumber(222), new BigNumber((_522_v21).length));
          (_522_v21)[_index65] = _491_v6;
          (globalState).f8 = _515_cf6;
          let _523_v22;
          _523_v22 = _dafny.Seq.of(_513_cf8, (_module.__default.fm33(_515_cf6, (_this).f12, _513_cf8, (_this).f12, globalState)).minus(_513_cf8), _module.__default.safeModuloInt(new BigNumber((_521_v20).length), _484_v0));
          _523_v22 = _module.__default.fm37(globalState);
        }
      } else if (_source13.is_DC2) {
        let _524___mcc_h8 = (_source13).cf4;
        let _525_cf4 = _524___mcc_h8;
        _492_v7 = _492_v7;
        (globalState).f8 = ((_this).f12) === ((_this).f12);
        (globalState).f8 = (_this).f12;
        if ((_this).f12) {
          let _526_v23;
          _526_v23 = _dafny.Seq.of(true, (_this).f12);
          let _527_v24;
          _527_v24 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f12),(_this).f12);
          let _528_v25;
          _528_v25 = _dafny.Seq.of(_526_v23);
          let _529_v26;
          let _nw60 = Array((new BigNumber(18)).toNumber());
          _nw60[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_526_v23, _526_v23);
          _nw60[(_dafny.ONE).toNumber()] = _module.__default.fm38((((_527_v24).contains((_this).f12)) ? ((_527_v24).get((_this).f12)) : ((_this).f12)), globalState);
          _nw60[(new BigNumber(2)).toNumber()] = _526_v23;
          _nw60[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_526_v23, (_528_v25)[_module.__default.safeIndex(_484_v0, new BigNumber((_528_v25).length))]);
          _nw60[(new BigNumber(4)).toNumber()] = (((_this).f12) ? (_526_v23) : (_526_v23));
          _nw60[(new BigNumber(5)).toNumber()] = _526_v23;
          _nw60[(new BigNumber(6)).toNumber()] = _dafny.Seq.of((_this).f12, !((_this).f12));
          _nw60[(new BigNumber(7)).toNumber()] = _526_v23;
          _nw60[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_526_v23, _526_v23);
          _nw60[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_526_v23, _526_v23);
          _nw60[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_526_v23, _526_v23);
          _nw60[(new BigNumber(11)).toNumber()] = _526_v23;
          _nw60[(new BigNumber(12)).toNumber()] = _module.__default.fm38((_this).f12, globalState);
          _nw60[(new BigNumber(13)).toNumber()] = _526_v23;
          _nw60[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_526_v23, _526_v23);
          _nw60[(new BigNumber(15)).toNumber()] = _526_v23;
          _nw60[(new BigNumber(16)).toNumber()] = _dafny.Seq.of((_this).f12, !((_491_v6).fm5((_this).f12, _484_v0, _484_v0, _501_v8, globalState)), (_this).f12);
          _nw60[(new BigNumber(17)).toNumber()] = _526_v23;
          _529_v26 = _nw60;
          let _index66 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_529_v26).length));
          (_529_v26)[_index66] = _dafny.Seq.update(_526_v23, _module.__default.safeIndex((_dafny.ZERO).minus(_484_v0), new BigNumber((_526_v23).length)), (_this).f12);
          let _index67 = _module.__default.safeIndex(new BigNumber(226), new BigNumber((_529_v26).length));
          (_529_v26)[_index67] = _dafny.Seq.update(_dafny.Seq.Concat(_526_v23, _526_v23), _module.__default.safeIndex(_484_v0, new BigNumber((_dafny.Seq.Concat(_526_v23, _526_v23)).length)), (_this).f12);
          (globalState).f8 = ((_529_v26)[_module.__default.safeIndex(new BigNumber(226), new BigNumber((_529_v26).length))])[_module.__default.safeIndex(_484_v0, new BigNumber(((_529_v26)[_module.__default.safeIndex(new BigNumber(226), new BigNumber((_529_v26).length))]).length))];
          (globalState).f8 = (_this).f12;
          let _530_v27;
          let _nw61 = new _module.C0();
          _nw61.__ctor();
          _530_v27 = _nw61;
          _484_v0 = _484_v0;
        } else {
          _492_v7 = new _dafny.CodePoint('m'.codePointAt(0));
          _484_v0 = (((_484_v0).isEqualTo(_484_v0)) ? (_484_v0) : (_484_v0));
          let _531_v28;
          let _init6 = ((_532_v0) => function (_533_i1) {
            return (_533_i1).plus(_532_v0);
          })(_484_v0);
          let _nw62 = Array((new BigNumber(21)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw62.length); _i0_6++) {
            _nw62[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _531_v28 = _nw62;
          let _534_v29;
          _534_v29 = _dafny.Map.Empty.slice().updateUnsafe((_this).f12,_531_v28);
          let _535_v30;
          let _nw63 = Array((new BigNumber(16)).toNumber());
          _nw63[(_dafny.ZERO).toNumber()] = _531_v28;
          _nw63[(_dafny.ONE).toNumber()] = _531_v28;
          _nw63[(new BigNumber(2)).toNumber()] = _531_v28;
          _nw63[(new BigNumber(3)).toNumber()] = _531_v28;
          _nw63[(new BigNumber(4)).toNumber()] = _531_v28;
          _nw63[(new BigNumber(5)).toNumber()] = _531_v28;
          _nw63[(new BigNumber(6)).toNumber()] = _531_v28;
          _nw63[(new BigNumber(7)).toNumber()] = _531_v28;
          _nw63[(new BigNumber(8)).toNumber()] = _531_v28;
          _nw63[(new BigNumber(9)).toNumber()] = _531_v28;
          _nw63[(new BigNumber(10)).toNumber()] = _531_v28;
          _nw63[(new BigNumber(11)).toNumber()] = _531_v28;
          _nw63[(new BigNumber(12)).toNumber()] = _531_v28;
          _nw63[(new BigNumber(13)).toNumber()] = (((_534_v29).contains((_this).f12)) ? ((_534_v29).get((_this).f12)) : (_531_v28));
          _nw63[(new BigNumber(14)).toNumber()] = _531_v28;
          _nw63[(new BigNumber(15)).toNumber()] = _531_v28;
          _535_v30 = _nw63;
          let _index68 = _module.__default.safeIndex(new BigNumber(96), new BigNumber((_535_v30).length));
          (_535_v30)[_index68] = _531_v28;
          let _index69 = _module.__default.safeIndex(new BigNumber(96), new BigNumber((_535_v30).length));
          (_535_v30)[_index69] = _531_v28;
          let _536_v31;
          _536_v31 = _dafny.Seq.of((_this).f12, !((_this).f12));
          let _537_v32;
          _537_v32 = _dafny.Map.Empty.slice().updateUnsafe((_this).f12,(_this).f12);
          let _538_v33;
          _538_v33 = _module.D6.create_DC16(_537_v32);
          let _539_v34;
          _539_v34 = _dafny.Map.Empty.slice().updateUnsafe((_536_v31)[_module.__default.safeIndex(new BigNumber(((_538_v33).dtor_cf28).length), new BigNumber((_536_v31).length))],new BigNumber(-724));
          _539_v34 = (_539_v34).update(!((_this).f12), _484_v0);
          (globalState).f8 = (_484_v0).isLessThan(_484_v0);
        }
      } else {
        let _540___mcc_h9 = (_source13).cf9;
        let _541_cf9 = _540___mcc_h9;
        _484_v0 = _484_v0;
        r1 = false;
        let _542_v35;
        _542_v35 = _module.D7.create_DC21(_dafny.Set.fromElements((_this).f12));
        let _543_v36;
        _543_v36 = _dafny.Set.fromElements((_this).f12, (_this).f12, (_this).f12);
        let _544_v37;
        _544_v37 = _dafny.MultiSet.fromElements(_484_v0, new BigNumber(-245), new BigNumber((((_542_v35).dtor_cf40).Union(_543_v36)).length), _484_v0, _484_v0);
        _544_v37 = (_544_v37).Difference(_544_v37);
        if (!(_484_v0).isEqualTo(_484_v0)) {
          _485_v1 = (_485_v1).Merge((function () {
            let _coll30 = new _dafny.Map();
            for (const _compr_30 of (_544_v37).Elements) {
              let _545_v38 = _compr_30;
              if ((_544_v37).contains(_545_v38)) {
                _coll30.push([_module.__default.safeModuloInt(_545_v38, new BigNumber(((_this).f13).length)),(_this).f12]);
              }
            }
            return _coll30;
          }()).update(_484_v0, (_this).f12));
          let _546_v39;
          let _nw64 = Array((new BigNumber(26)).toNumber());
          _nw64[(_dafny.ZERO).toNumber()] = _492_v7;
          _nw64[(_dafny.ONE).toNumber()] = p0;
          _nw64[(new BigNumber(2)).toNumber()] = _492_v7;
          _nw64[(new BigNumber(3)).toNumber()] = _492_v7;
          _nw64[(new BigNumber(4)).toNumber()] = p0;
          _nw64[(new BigNumber(5)).toNumber()] = p0;
          _nw64[(new BigNumber(6)).toNumber()] = p0;
          _nw64[(new BigNumber(7)).toNumber()] = _492_v7;
          _nw64[(new BigNumber(8)).toNumber()] = _492_v7;
          _nw64[(new BigNumber(9)).toNumber()] = p0;
          _nw64[(new BigNumber(10)).toNumber()] = _492_v7;
          _nw64[(new BigNumber(11)).toNumber()] = p0;
          _nw64[(new BigNumber(12)).toNumber()] = new _dafny.CodePoint('v'.codePointAt(0));
          _nw64[(new BigNumber(13)).toNumber()] = new _dafny.CodePoint('m'.codePointAt(0));
          _nw64[(new BigNumber(14)).toNumber()] = _492_v7;
          _nw64[(new BigNumber(15)).toNumber()] = p0;
          _nw64[(new BigNumber(16)).toNumber()] = p0;
          _nw64[(new BigNumber(17)).toNumber()] = _492_v7;
          _nw64[(new BigNumber(18)).toNumber()] = p0;
          _nw64[(new BigNumber(19)).toNumber()] = p0;
          _nw64[(new BigNumber(20)).toNumber()] = p0;
          _nw64[(new BigNumber(21)).toNumber()] = p0;
          _nw64[(new BigNumber(22)).toNumber()] = _492_v7;
          _nw64[(new BigNumber(23)).toNumber()] = p0;
          _nw64[(new BigNumber(24)).toNumber()] = _492_v7;
          _nw64[(new BigNumber(25)).toNumber()] = _492_v7;
          _546_v39 = _nw64;
          let _index70 = _module.__default.safeIndex(new BigNumber(62), new BigNumber((_546_v39).length));
          (_546_v39)[_index70] = _492_v7;
          let _index71 = _module.__default.safeIndex(new BigNumber(62), new BigNumber((_546_v39).length));
          (_546_v39)[_index71] = new _dafny.CodePoint('b'.codePointAt(0));
          let _547_v40;
          _547_v40 = _dafny.Seq.of(new BigNumber(-832));
          let _rhs66 = (_this).f12;
          let _rhs67 = (new BigNumber((_543_v36).length)).multipliedBy(_module.__default.safeModuloInt(_484_v0, new BigNumber((_547_v40).length)));
          let _lhs47 = globalState;
          _lhs47.f8 = _rhs66;
          _484_v0 = _rhs67;
          let _548_v41;
          _548_v41 = _dafny.Seq.UnicodeFromString("rax");
          _548_v41 = (_this).f13;
          _485_v1 = (_485_v1).update(_484_v0, true);
        } else {
          let _549_v42;
          let _nw65 = new _module.C0();
          _nw65.__ctor();
          _549_v42 = _nw65;
          let _550_v43;
          _550_v43 = _dafny.Seq.of((_this).f12);
          _550_v43 = _550_v43;
          let _551_v44;
          _551_v44 = _dafny.Seq.UnicodeFromString("feeckpr");
          _551_v44 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(9)), ((_552_p0) => function (_553_i2) {
            return _552_p0;
          })(p0));
          let _554_v45;
          _554_v45 = _dafny.Map.Empty.slice().updateUnsafe((_this).f12,(_this).f12);
          let _555_v46;
          _555_v46 = _dafny.Set.fromElements(_module.__default.safeModuloInt(new BigNumber(-682), _module.__default.fm33((_this).f12, (_this).f12, new BigNumber((_554_v45).length), (_module.D1.create_DC4((_this).f12, (_this).f12, _484_v0)).dtor_cf6, globalState)));
          _555_v46 = function () {
            let _coll31 = new _dafny.Set();
            for (const _compr_31 of _dafny.IntegerRange(new BigNumber(825), new BigNumber(781))) {
              let _556_v47 = _compr_31;
              if (((new BigNumber(825)).isLessThanOrEqualTo(_556_v47)) && ((_556_v47).isLessThan(new BigNumber(781)))) {
                _coll31.add((_556_v47).multipliedBy(new BigNumber((_550_v43).length)));
              }
            }
            return _coll31;
          }();
          let _557_v48;
          let _nw66 = Array((new BigNumber(14)).toNumber()).fill(false);
          _557_v48 = _nw66;
          _557_v48 = _557_v48;
        }
      }
      let _558_v49;
      let _nw67 = new _module.C0();
      _nw67.__ctor();
      _558_v49 = _nw67;
      if ((_this).f12) {
        if ((_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("xplooehy")).length), _484_v0)).isLessThan(_484_v0)) {
          let _559_v50;
          _559_v50 = _module.D3.create_DC8((_this).f12, (_this).f12);
          _559_v50 = (((_this).f12) ? (_559_v50) : (function (_pat_let6_0) {
            return function (_560_dt__update__tmp_h0) {
              return function (_pat_let7_0) {
                return function (_561_dt__update_hcf12_h0) {
                  return _module.D3.create_DC8(_561_dt__update_hcf12_h0, (_560_dt__update__tmp_h0).dtor_cf13);
                }(_pat_let7_0);
              }((_this).f12);
            }(_pat_let6_0);
          }(_559_v50)));
          let _562_v51;
          _562_v51 = _dafny.Seq.of(_485_v1);
          let _563_v52;
          _563_v52 = _dafny.Map.Empty.slice().updateUnsafe(((_562_v51)[_module.__default.safeIndex(new BigNumber(((_this).f13).length), new BigNumber((_562_v51).length))]).update(new BigNumber(-24), (_this).f12),_484_v0);
          _484_v0 = new BigNumber(((_563_v52).update(_485_v1, (((_this).f12) ? (_484_v0) : (new BigNumber(-270))))).length);
          (globalState).f8 = (_491_v6).fm5((_this).f12, (_dafny.ZERO).minus(_484_v0), _module.__default.safeDivisionInt(_484_v0, _484_v0), _501_v8, globalState);
          let _564_v53;
          let _nw68 = new _module.C0();
          _nw68.__ctor();
          _564_v53 = _nw68;
          let _565_v55;
          _565_v55 = _dafny.Set.fromElements(_module.__default.fm33((_this).f12, (_this).f12, new BigNumber(953), (_this).f12, globalState), _484_v0, _484_v0, _484_v0, _484_v0);
          let _566_v56;
          _566_v56 = _dafny.MultiSet.fromElements(function () {
            let _coll32 = new _dafny.Set();
            for (const _compr_32 of _dafny.IntegerRange(new BigNumber(954), new BigNumber(800))) {
              let _567_v54 = _compr_32;
              if (((new BigNumber(954)).isLessThanOrEqualTo(_567_v54)) && ((_567_v54).isLessThan(new BigNumber(800)))) {
                _coll32.add((_567_v54).plus((((_489_v5).contains(_484_v0)) ? ((_489_v5).get(_484_v0)) : (_484_v0))));
              }
            }
            return _coll32;
          }(), _565_v55, _565_v55);
          _566_v56 = (_566_v56).Union(_566_v56);
        } else {
          let _568_v57;
          let _init7 = function (_569_i3) {
            return true;
          };
          let _nw69 = Array((new BigNumber(24)).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw69.length); _i0_7++) {
            _nw69[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _568_v57 = _nw69;
          let _570_v58;
          let _nw70 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
          _570_v58 = _nw70;
          let _571_v59;
          _571_v59 = _module.D8.create_DC23(_dafny.Seq.of((_this).f12, (_this).f12));
          let _index72 = _module.__default.safeIndex(new BigNumber(319), new BigNumber((_570_v58).length));
          (_570_v58)[_index72] = new BigNumber((_dafny.MultiSet.FromArray((_571_v59).dtor_cf41)).cardinality());
          let _572_v60;
          let _nw71 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _572_v60 = _nw71;
          let _573_v61;
          _573_v61 = _dafny.Map.Empty.slice().updateUnsafe(_572_v60,(_484_v0).multipliedBy(_module.__default.fm33((_this).f12, (_this).fm32(_484_v0, _492_v7, globalState), _484_v0, true, globalState)));
          let _index73 = _module.__default.safeIndex(new BigNumber(319), new BigNumber((_570_v58).length));
          let _rhs68 = _568_v57;
          let _rhs69 = (_this).f12;
          let _rhs70 = (((_573_v61).contains(_572_v60)) ? ((_573_v61).get(_572_v60)) : (_484_v0));
          let _lhs48 = _570_v58;
          let _lhs49 = _module.__default.safeIndex(new BigNumber(319), new BigNumber((_570_v58).length));
          _568_v57 = _rhs68;
          r1 = _rhs69;
          _lhs48[_lhs49] = _rhs70;
          let _574_v62;
          _574_v62 = _dafny.Set.fromElements(_570_v58, _570_v58);
          let _575_v63;
          _575_v63 = _module.D8.create_DC24((_this).f12, _574_v62, (new BigNumber(-247)).multipliedBy((_570_v58)[_module.__default.safeIndex(new BigNumber(319), new BigNumber((_570_v58).length))]));
          _575_v63 = _575_v63;
          let _576_v65;
          _576_v65 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("igpmkqt")).length), (_570_v58)[_module.__default.safeIndex(new BigNumber(319), new BigNumber((_570_v58).length))]);
          let _577_v66;
          _577_v66 = _dafny.Seq.of(_module.__default.fm39((_this).f12, (_dafny.ZERO).minus((_570_v58)[_module.__default.safeIndex(new BigNumber(319), new BigNumber((_570_v58).length))]), true, globalState), function () {
            let _coll33 = new _dafny.Map();
            for (const _compr_33 of (_576_v65).Elements) {
              let _578_v64 = _compr_33;
              if ((_576_v65).contains(_578_v64)) {
                _coll33.push([(_578_v64).multipliedBy((_570_v58)[_module.__default.safeIndex(new BigNumber(319), new BigNumber((_570_v58).length))]),(_570_v58)[_module.__default.safeIndex(new BigNumber(319), new BigNumber((_570_v58).length))]]);
              }
            }
            return _coll33;
          }(), _489_v5, _dafny.Map.Empty.slice().updateUnsafe((_570_v58)[_module.__default.safeIndex(new BigNumber(319), new BigNumber((_570_v58).length))],new BigNumber(90)));
          let _index74 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_568_v57).length));
          (_568_v57)[_index74] = (false) || (true);
          let _579_v67;
          _579_v67 = _dafny.Seq.of(_484_v0);
          let _580_v68;
          _580_v68 = _module.D5.create_DC14((_570_v58)[_module.__default.safeIndex(new BigNumber(319), new BigNumber((_570_v58).length))], (_this).f12);
          let _581_v69;
          _581_v69 = _dafny.Set.fromElements((_this).f12, (_this).f12);
          let _582_v70;
          _582_v70 = _dafny.Seq.of((_558_v49).fm5((_this).f12, (_570_v58)[_module.__default.safeIndex(new BigNumber(319), new BigNumber((_570_v58).length))], new BigNumber((_581_v69).length), _501_v8, globalState));
          let _583_v71;
          _583_v71 = _dafny.Seq.of(_582_v70);
          let _index75 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_568_v57).length));
          let _rhs71 = (_570_v58)[_module.__default.safeIndex(new BigNumber(319), new BigNumber((_570_v58).length))];
          let _rhs72 = _module.__default.fm40(_579_v67, _580_v68, new BigNumber(((_this).f13).length), globalState);
          let _rhs73 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_582_v70, _582_v70), (_583_v71)[_module.__default.safeIndex(new BigNumber(-736), new BigNumber((_583_v71).length))]);
          let _lhs50 = _568_v57;
          let _lhs51 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_568_v57).length));
          _484_v0 = _rhs71;
          _577_v66 = _rhs72;
          _lhs50[_lhs51] = _rhs73;
          let _index76 = _module.__default.safeIndex(new BigNumber(319), new BigNumber((_570_v58).length));
          (_570_v58)[_index76] = _module.__default.safeDivisionInt(new BigNumber(135), _484_v0);
          let _584_v72;
          let _nw72 = new _module.C0();
          _nw72.__ctor();
          _584_v72 = _nw72;
        }
        let _585_v73;
        let _init8 = function (_586_i4) {
          return (_this).f12;
        };
        let _nw73 = Array((new BigNumber(17)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw73.length); _i0_8++) {
          _nw73[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _585_v73 = _nw73;
        let _index77 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((_585_v73).length));
        (_585_v73)[_index77] = (_this).f12;
        let _index78 = _module.__default.safeIndex(new BigNumber(125), new BigNumber((_585_v73).length));
        (_585_v73)[_index78] = (_this).f12;
        let _587_v74;
        let _nw74 = new _module.C0();
        _nw74.__ctor();
        _587_v74 = _nw74;
        let _588_v75;
        _588_v75 = _dafny.Seq.of(new BigNumber((_dafny.Seq.update((_this).f13, _module.__default.safeIndex(_484_v0, new BigNumber(((_this).f13).length)), p0)).length), _484_v0);
        let _589_v76;
        _589_v76 = _dafny.Map.Empty.slice().updateUnsafe(_484_v0,_588_v75);
        let _590_v77;
        _590_v77 = _dafny.Map.Empty.slice().updateUnsafe((_this).f13,new BigNumber((_dafny.Seq.Concat(_588_v75, (((_589_v76).contains(_484_v0)) ? ((_589_v76).get(_484_v0)) : (_588_v75)))).length));
        _590_v77 = (_590_v77).update((_this).f13, (_484_v0).plus(new BigNumber((_588_v75).length)));
        _484_v0 = (_dafny.ZERO).minus((new BigNumber(144)).plus(_484_v0));
      } else {
        let _591_v78;
        _591_v78 = _dafny.Seq.of((_this).f12, (_this).f12);
        r1 = ((_591_v78)[_module.__default.safeIndex(_484_v0, new BigNumber((_591_v78).length))]) || (!_dafny.Seq.contains((_this).f13, p0));
        let _592_v79;
        _592_v79 = _module.D6.create_DC18(p0, _484_v0, (_484_v0).minus(_484_v0));
        _592_v79 = _592_v79;
        r1 = (((_485_v1).contains((((_489_v5).contains(_484_v0)) ? ((_489_v5).get(_484_v0)) : (new BigNumber(142))))) ? ((_485_v1).get((((_489_v5).contains(_484_v0)) ? ((_489_v5).get(_484_v0)) : (new BigNumber(142))))) : ((_this).f12));
        let _593_v80;
        let _nw75 = new _module.C0();
        _nw75.__ctor();
        _593_v80 = _nw75;
        _492_v7 = _492_v7;
      }
      let _594_v81;
      let _init9 = ((_595_v0) => function (_596_i5) {
        return (_596_i5).multipliedBy(_595_v0);
      })(_484_v0);
      let _nw76 = Array((new BigNumber(12)).toNumber());
      for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw76.length); _i0_9++) {
        _nw76[_i0_9] = _init9(new BigNumber(_i0_9));
      }
      _594_v81 = _nw76;
      let _597_v82;
      _597_v82 = _dafny.Map.Empty.slice().updateUnsafe(_594_v81,_dafny.Seq.UnicodeFromString("lhau"));
      r0 = _597_v82;
      r1 = _dafny.Seq.contains((_this).f13, (((_this).f12) ? (p0) : (p0)));
      return [r0, r1];
    }
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f9 = _dafny.Map.Empty;
      this._f10 = false;
      this.f14 = false;
      this._f15 = [];
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f10() {
      let _this = this;
      return _this._f10;
    };
    __ctor(f14, f15, f9, f10) {
      let _this = this;
      (_this).f14 = f14;
      (_this)._f15 = f15;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      return;
    }
    fm0(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !(new BigNumber(((function () {
        let _coll34 = new _dafny.Map();
        for (const _compr_34 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("dnq")).length),(_this).f9)).Keys.Elements) {
          let _598_v0 = _compr_34;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("dnq")).length),(_this).f9)).contains(_598_v0)) {
            _coll34.push([(_598_v0).multipliedBy(new BigNumber(688)),(_module.D0.create_DC1(new BigNumber((_dafny.Set.fromElements(_this.f14, (_this).f10)).length), _this.f14, (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(289)), function (_599_i0) {
  return new _dafny.CodePoint('m'.codePointAt(0));
})).length))).length)))).dtor_cf3]);
          }
        }
        return _coll34;
      }())).length)).isEqualTo(new BigNumber(301));
    };
    fm1(p0, globalState) {
      let _this = this;
      if (!(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("bnhgpf")).length),_module.D6.create_DC17((_this).f10, _this.f14, new BigNumber(55), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("xy")).length),new BigNumber((_dafny.Seq.UnicodeFromString("uyxnksw")).length))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(255),_this.f14)).length)))).length)).isEqualTo(new BigNumber(-987))) {
        return ((_dafny.ZERO).minus(new BigNumber(-395))).minus(new BigNumber((_dafny.Set.fromElements(_this.f14)).length));
      } else {
        return ((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(963),_this.f14)).length))).plus(new BigNumber(-169));
      }
    };
    fm41(p0, p1, p2, globalState) {
      let _this = this;
      return _dafny.Set.fromElements(_this.f14, _this.f14, (_dafny.MultiSet.fromElements(new BigNumber(196), new BigNumber(419), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("kepfnihg")).length),new BigNumber((function () {
        let _coll35 = new _dafny.Map();
        for (const _compr_35 of (_dafny.MultiSet.fromElements(new BigNumber(411))).Elements) {
          let _600_v0 = _compr_35;
          if ((_dafny.MultiSet.fromElements(new BigNumber(411))).contains(_600_v0)) {
            _coll35.push([(_600_v0).minus(new BigNumber(396)),new BigNumber(585)]);
          }
        }
        return _coll35;
      }()).length))).length))).equals(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("xeverqrey")).length))), !((new BigNumber(944)).isLessThanOrEqualTo(new BigNumber(289))));
    };
    fm42(p0, p1, p2, globalState) {
      let _this = this;
      return new BigNumber(24);
    };
    m0(globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = false;
      let r2 = _dafny.Seq.of();
      let r3 = _dafny.Seq.UnicodeFromString("");
      let _601_v0;
      _601_v0 = _dafny.MultiSet.fromElements(_this.f14);
      let _602_v1;
      _602_v1 = _dafny.Seq.of(_this.f14);
      let _603_v2;
      _603_v2 = new BigNumber(689);
      let _604_v3;
      _604_v3 = _dafny.Map.Empty.slice().updateUnsafe((_601_v0).IsDisjointFrom(_dafny.MultiSet.FromArray(_dafny.Seq.update(_602_v1, _module.__default.safeIndex((_this).fm42(_603_v2, _603_v2, _603_v2, globalState), new BigNumber((_602_v1).length)), (_this).f10))),_603_v2);
      let _605_v4;
      _605_v4 = _dafny.MultiSet.fromElements(_603_v2);
      _604_v3 = (_604_v3).update(((false) ? (true) : (_this.f14)), (_dafny.ZERO).minus((((_605_v4).contains((_dafny.ZERO).minus(_603_v2))) ? ((_605_v4).get((_dafny.ZERO).minus(_603_v2))) : (_603_v2))));
      (_this).f14 = (_this).f10;
      let _606_v5;
      let _init10 = function (_607_i0) {
        return new _dafny.CodePoint('f'.codePointAt(0));
      };
      let _nw77 = Array((new BigNumber(22)).toNumber());
      for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw77.length); _i0_10++) {
        _nw77[_i0_10] = _init10(new BigNumber(_i0_10));
      }
      _606_v5 = _nw77;
      let _index79 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_606_v5).length));
      (_606_v5)[_index79] = new _dafny.CodePoint('d'.codePointAt(0));
      let _index80 = _module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length));
      ((_this).f15)[_index80] = (_this).fm0(_602_v1, (_this).f10, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-427)), ((_608_v2) => function (_609_i1) {
        return _608_v2;
      })(_603_v2))).length), _this.f14, globalState);
      let _610_v7;
      _610_v7 = _dafny.Seq.of(_603_v2);
      let _611_v8;
      _611_v8 = _module.D6.create_DC17(_this.f14, (_this).f10, new BigNumber(-155), (_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((function () {
  let _coll36 = new _dafny.Map();
  for (const _compr_36 of (_610_v7).Elements) {
    let _612_v6 = _compr_36;
    if (_dafny.Seq.contains(_610_v7, _612_v6)) {
      _coll36.push([_module.__default.safeDivisionInt(_612_v6, _603_v2),_this.f14]);
    }
  }
  return _coll36;
}()).length)))), (_this).fm1((_this).f9, globalState));
      let _613_v9;
      _613_v9 = _module.D5.create_DC13((_this).f15);
      let _index81 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_606_v5).length));
      let _index82 = _module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length));
      let _rhs74 = (_611_v8).dtor_cf33;
      let _rhs75 = _module.__default.fm43((new BigNumber((_dafny.Seq.of(_613_v9)).length)).minus(new BigNumber(170)), globalState);
      let _rhs76 = _604_v3;
      let _rhs77 = !((_this).f10) || ((_this).f10);
      let _rhs78 = _603_v2;
      let _lhs52 = _606_v5;
      let _lhs53 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_606_v5).length));
      let _lhs54 = (_this).f15;
      let _lhs55 = _module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length));
      _603_v2 = _rhs74;
      _lhs52[_lhs53] = _rhs75;
      _604_v3 = _rhs76;
      _lhs54[_lhs55] = _rhs77;
      _603_v2 = _rhs78;
      let _614_i2;
      _614_i2 = _dafny.ZERO;
      L0: {
        while (!(((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))])) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_614_i2)) {
              break L0;
            }
            _614_i2 = (_614_i2).plus(_dafny.ONE);
            (globalState).f8 = (_603_v2).isLessThanOrEqualTo(_603_v2);
            r1 = ((_this).f10) || ((_this).f10);
            _603_v2 = _603_v2;
            let _index83 = _module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length));
            ((_this).f15)[_index83] = (false) && (_this.f14);
          }
        }
      }
      let _615_v10;
      _615_v10 = _dafny.Map.Empty.slice().updateUnsafe(_603_v2,((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))]);
      _615_v10 = _module.__default.fm44(false, globalState);
      if (((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))]) {
        let _616_v11;
        _616_v11 = _dafny.Seq.UnicodeFromString("wem");
        let _617_v12;
        _617_v12 = _dafny.Seq.of(_616_v11);
        let _618_v13;
        let _nw78 = new _module.C1();
        _nw78.__ctor(((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))], (_617_v12)[_module.__default.safeIndex(new BigNumber((_616_v11).length), new BigNumber((_617_v12).length))]);
        _618_v13 = _nw78;
        let _619_v14;
        let _nw79 = Array((new BigNumber(17)).toNumber()).fill([]);
        _619_v14 = _nw79;
        let _620_v15;
        _620_v15 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_602_v1);
        let _621_v16;
        let _nw80 = Array((new BigNumber(22)).toNumber());
        _nw80[(_dafny.ZERO).toNumber()] = (_618_v13).f13;
        _nw80[(_dafny.ONE).toNumber()] = _616_v11;
        _nw80[(new BigNumber(2)).toNumber()] = _module.__default.fm45(!(_this.f14), _603_v2, (_618_v13).f12, globalState);
        _nw80[(new BigNumber(3)).toNumber()] = _616_v11;
        _nw80[(new BigNumber(4)).toNumber()] = _module.__default.fm45((_602_v1)[_module.__default.safeIndex(_603_v2, new BigNumber((_602_v1).length))], _603_v2, _this.f14, globalState);
        _nw80[(new BigNumber(5)).toNumber()] = (_618_v13).f13;
        _nw80[(new BigNumber(6)).toNumber()] = _616_v11;
        _nw80[(new BigNumber(7)).toNumber()] = _616_v11;
        _nw80[(new BigNumber(8)).toNumber()] = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("cmnxwuk"), _module.__default.safeIndex(_603_v2, new BigNumber((_dafny.Seq.UnicodeFromString("cmnxwuk")).length)), (_606_v5)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_606_v5).length))]);
        _nw80[(new BigNumber(9)).toNumber()] = _616_v11;
        _nw80[(new BigNumber(10)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(486)), ((_622_v5) => function (_623_i3) {
          return (_622_v5)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_622_v5).length))];
        })(_606_v5));
        _nw80[(new BigNumber(11)).toNumber()] = (_618_v13).f13;
        _nw80[(new BigNumber(12)).toNumber()] = (_618_v13).f13;
        _nw80[(new BigNumber(13)).toNumber()] = _dafny.Seq.update((_618_v13).f13, _module.__default.safeIndex(new BigNumber((_620_v15).length), new BigNumber(((_618_v13).f13).length)), (_606_v5)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_606_v5).length))]);
        _nw80[(new BigNumber(14)).toNumber()] = (_618_v13).f13;
        _nw80[(new BigNumber(15)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(202)), ((_624_v5) => function (_625_i4) {
          return (_624_v5)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_624_v5).length))];
        })(_606_v5));
        _nw80[(new BigNumber(16)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(31)), ((_626_v5) => function (_627_i5) {
          return (_626_v5)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_626_v5).length))];
        })(_606_v5));
        _nw80[(new BigNumber(17)).toNumber()] = _616_v11;
        _nw80[(new BigNumber(18)).toNumber()] = _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.update(_616_v11, _module.__default.safeIndex(new BigNumber(((_618_v13).f13).length), new BigNumber((_616_v11).length)), new _dafny.CodePoint('o'.codePointAt(0))), _module.__default.safeIndex(_603_v2, new BigNumber((_dafny.Seq.update(_616_v11, _module.__default.safeIndex(new BigNumber(((_618_v13).f13).length), new BigNumber((_616_v11).length)), new _dafny.CodePoint('o'.codePointAt(0)))).length)), (_606_v5)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_606_v5).length))]), _module.__default.safeIndex(_603_v2, new BigNumber((_dafny.Seq.update(_dafny.Seq.update(_616_v11, _module.__default.safeIndex(new BigNumber(((_618_v13).f13).length), new BigNumber((_616_v11).length)), new _dafny.CodePoint('o'.codePointAt(0))), _module.__default.safeIndex(_603_v2, new BigNumber((_dafny.Seq.update(_616_v11, _module.__default.safeIndex(new BigNumber(((_618_v13).f13).length), new BigNumber((_616_v11).length)), new _dafny.CodePoint('o'.codePointAt(0)))).length)), (_606_v5)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_606_v5).length))])).length)), (_606_v5)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_606_v5).length))]);
        _nw80[(new BigNumber(19)).toNumber()] = _616_v11;
        _nw80[(new BigNumber(20)).toNumber()] = (_618_v13).f13;
        _nw80[(new BigNumber(21)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(591)), function (_628_i6) {
          return new _dafny.CodePoint('u'.codePointAt(0));
        });
        _621_v16 = _nw80;
        let _index84 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_619_v14).length));
        (_619_v14)[_index84] = _621_v16;
        let _index85 = _module.__default.safeIndex(new BigNumber(918), new BigNumber((_619_v14).length));
        (_619_v14)[_index85] = _621_v16;
        (globalState).f8 = (_this).fm0(_dafny.Seq.update(_602_v1, _module.__default.safeIndex(new BigNumber((_602_v1).length), new BigNumber((_602_v1).length)), _this.f14), (_618_v13).f12, _603_v2, _dafny.areEqual(_602_v1, _602_v1), globalState);
        let _629_v17;
        let _nw81 = new _module.C1();
        _nw81.__ctor(!(_this.f14), (_618_v13).f13);
        _629_v17 = _nw81;
        let _630_v18;
        _630_v18 = _module.D6.create_DC18((_606_v5)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_606_v5).length))], _603_v2, (_603_v2).plus(_603_v2));
        let _rhs79 = _630_v18;
        let _rhs80 = !(_dafny.Set.fromElements(_630_v18)).contains(_module.D6.create_DC18(((_629_v17).f13)[_module.__default.safeIndex(new BigNumber((_610_v7).length), new BigNumber(((_629_v17).f13).length))], _603_v2, _603_v2));
        let _lhs56 = _this;
        _630_v18 = _rhs79;
        _lhs56.f14 = _rhs80;
      } else {
        let _source14 = _module.D1.create_DC5(_module.D1.create_DC3(_dafny.Seq.UnicodeFromString("n")));
        if (_source14.is_DC3) {
          let _631___mcc_h0 = (_source14).cf5;
          let _632_cf5 = _631___mcc_h0;
          let _rhs81 = (_this).f10;
          let _rhs82 = ((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))];
          let _rhs83 = _603_v2;
          let _lhs57 = globalState;
          let _lhs58 = _this;
          _lhs57.f8 = _rhs81;
          _lhs58.f14 = _rhs82;
          _603_v2 = _rhs83;
          let _633_v19;
          let _nw82 = new _module.C1();
          _nw82.__ctor(_this.f14, _dafny.Seq.update(_632_cf5, _module.__default.safeIndex(_603_v2, new BigNumber((_632_cf5).length)), (_606_v5)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_606_v5).length))]));
          _633_v19 = _nw82;
          let _634_v20;
          _634_v20 = _dafny.Set.fromElements(_603_v2, _603_v2);
          let _index86 = _module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length));
          let _index87 = _module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length));
          let _rhs84 = _this.f14;
          let _rhs85 = (_634_v20).IsProperSubsetOf(_634_v20);
          let _lhs59 = (_this).f15;
          let _lhs60 = _module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length));
          let _lhs61 = (_this).f15;
          let _lhs62 = _module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length));
          _lhs59[_lhs60] = _rhs84;
          _lhs61[_lhs62] = _rhs85;
          let _635_v21;
          let _nw83 = Array((new BigNumber(23)).toNumber());
          _nw83[(_dafny.ZERO).toNumber()] = _this.f14;
          _nw83[(_dafny.ONE).toNumber()] = (_this).f10;
          _nw83[(new BigNumber(2)).toNumber()] = (_633_v19).f12;
          _nw83[(new BigNumber(3)).toNumber()] = (_633_v19).f12;
          _nw83[(new BigNumber(4)).toNumber()] = ((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))];
          _nw83[(new BigNumber(5)).toNumber()] = (_633_v19).f12;
          _nw83[(new BigNumber(6)).toNumber()] = (_633_v19).fm32(_603_v2, new _dafny.CodePoint('i'.codePointAt(0)), globalState);
          _nw83[(new BigNumber(7)).toNumber()] = true;
          _nw83[(new BigNumber(8)).toNumber()] = ((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))];
          _nw83[(new BigNumber(9)).toNumber()] = false;
          _nw83[(new BigNumber(10)).toNumber()] = _this.f14;
          _nw83[(new BigNumber(11)).toNumber()] = !((_633_v19).f12);
          _nw83[(new BigNumber(12)).toNumber()] = !((_this).f10);
          _nw83[(new BigNumber(13)).toNumber()] = !(true);
          _nw83[(new BigNumber(14)).toNumber()] = _this.f14;
          _nw83[(new BigNumber(15)).toNumber()] = ((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))];
          _nw83[(new BigNumber(16)).toNumber()] = ((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))];
          _nw83[(new BigNumber(17)).toNumber()] = (_633_v19).f12;
          _nw83[(new BigNumber(18)).toNumber()] = _this.f14;
          _nw83[(new BigNumber(19)).toNumber()] = ((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))];
          _nw83[(new BigNumber(20)).toNumber()] = (_633_v19).f12;
          _nw83[(new BigNumber(21)).toNumber()] = ((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))];
          _nw83[(new BigNumber(22)).toNumber()] = (_this).f10;
          _635_v21 = _nw83;
          let _636_v22;
          _636_v22 = _dafny.Map.Empty.slice().updateUnsafe(!((_633_v19).f12),_635_v21);
          _636_v22 = (((((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))]) ? (_636_v22) : (_636_v22))).Merge((_636_v22));
        } else if (_source14.is_DC4) {
          let _637___mcc_h1 = (_source14).cf6;
          let _638___mcc_h2 = (_source14).cf7;
          let _639___mcc_h3 = (_source14).cf8;
          let _640_cf8 = _639___mcc_h3;
          let _641_cf7 = _638___mcc_h2;
          let _642_cf6 = _637___mcc_h1;
          let _643_v23;
          _643_v23 = _module.D4.create_DC11((_this).fm1((_this).f9, globalState), _dafny.Set.fromElements(_640_cf8, _603_v2), _641_cf7);
          _640_cf8 = (_603_v2).multipliedBy((_643_v23).dtor_cf16);
          let _rhs86 = new BigNumber(-571);
          let _rhs87 = _641_cf7;
          let _lhs63 = globalState;
          _640_cf8 = _rhs86;
          _lhs63.f8 = _rhs87;
          let _644_v24;
          _644_v24 = _dafny.Map.Empty.slice().updateUnsafe(_603_v2,_module.__default.safeModuloInt(new BigNumber((_dafny.Set.fromElements(_640_cf8, _640_cf8)).length), _640_cf8));
          _644_v24 = (_644_v24).update((_this).fm42(_603_v2, (_dafny.ZERO).minus(new BigNumber(-380)), _603_v2, globalState), _module.__default.safeModuloInt(_640_cf8, _640_cf8));
          let _645_v25;
          let _nw84 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
          _645_v25 = _nw84;
          let _index88 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_645_v25).length));
          (_645_v25)[_index88] = _640_cf8;
          let _index89 = _module.__default.safeIndex(new BigNumber(405), new BigNumber((_645_v25).length));
          (_645_v25)[_index89] = (_this).fm1((_dafny.Map.Empty.slice().updateUnsafe((_this).fm0(_602_v1, !(_642_cf6), _603_v2, _641_cf7, globalState),_this.f14)).update(true, _642_cf6), globalState);
        } else if (_source14.is_DC2) {
          let _646___mcc_h4 = (_source14).cf4;
          let _647_cf4 = _646___mcc_h4;
          (globalState).f8 = (_603_v2).isLessThanOrEqualTo(_603_v2);
          let _648_v26;
          let _nw85 = new _module.C0();
          _nw85.__ctor();
          _648_v26 = _nw85;
          let _649_v27;
          let _init11 = ((_650_v2) => function (_651_i7) {
            return _module.__default.safeDivisionInt(_651_i7, (_dafny.ZERO).minus(_650_v2));
          })(_603_v2);
          let _nw86 = Array((new BigNumber(8)).toNumber());
          for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw86.length); _i0_11++) {
            _nw86[_i0_11] = _init11(new BigNumber(_i0_11));
          }
          _649_v27 = _nw86;
          let _652_v28;
          _652_v28 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm1((_this).f9, globalState),new BigNumber(-190));
          let _653_v29;
          _653_v29 = _dafny.Seq.UnicodeFromString("rnvidwis");
          let _654_v30;
          _654_v30 = _dafny.Map.Empty.slice().updateUnsafe(_653_v29,_603_v2);
          let _index90 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_649_v27).length));
          (_649_v27)[_index90] = ((((_652_v28).contains(new BigNumber((_653_v29).length))) ? ((_652_v28).get(new BigNumber((_653_v29).length))) : (_603_v2))).minus((((_654_v30).contains(_653_v29)) ? ((_654_v30).get(_653_v29)) : (_603_v2)));
          let _index91 = _module.__default.safeIndex(new BigNumber(935), new BigNumber((_649_v27).length));
          (_649_v27)[_index91] = _603_v2;
          let _655_v31;
          _655_v31 = _dafny.Map.Empty.slice().updateUnsafe(!(_603_v2).isEqualTo(_603_v2),((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))]);
          _655_v31 = (_655_v31).update(((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))], ((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))]);
        } else {
          let _656___mcc_h5 = (_source14).cf9;
          let _657_cf9 = _656___mcc_h5;
          _601_v0 = (_601_v0).Difference(_601_v0);
          let _658_v32;
          _658_v32 = _module.D0.create_DC0(_this.f14);
          (_this).f14 = ((((_this).f10) ? (_658_v32) : (_658_v32))).dtor_cf0;
          let _659_v33;
          let _nw87 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
          _659_v33 = _nw87;
          let _660_v34;
          _660_v34 = _dafny.Set.fromElements(_659_v33);
          let _661_v35;
          _661_v35 = _module.D8.create_DC24((_this).f10, _660_v34, _603_v2);
          let _662_v36;
          _662_v36 = _dafny.Seq.of(_661_v35);
          let _663_v37;
          _663_v37 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_662_v36, _dafny.Seq.update(_662_v36, _module.__default.safeIndex(new BigNumber(949), new BigNumber((_662_v36).length)), _661_v35)));
          _663_v37 = (_663_v37).Intersect(_663_v37);
          let _664_v38;
          let _665_v39;
          let _666_v40;
          let _out42;
          let _out43;
          let _out44;
          let _outcollector12 = (_this).m7(new BigNumber(511), _603_v2, globalState);
          _out42 = _outcollector12[0];
          _out43 = _outcollector12[1];
          _out44 = _outcollector12[2];
          _664_v38 = _out42;
          _665_v39 = _out43;
          _666_v40 = _out44;
        }
        let _667_v41;
        let _nw88 = Array((new BigNumber(15)).toNumber()).fill([]);
        _667_v41 = _nw88;
        let _668_v42;
        let _init12 = function (_669_i8) {
          return ((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))];
        };
        let _nw89 = Array((new BigNumber(6)).toNumber());
        for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw89.length); _i0_12++) {
          _nw89[_i0_12] = _init12(new BigNumber(_i0_12));
        }
        _668_v42 = _nw89;
        let _index92 = _module.__default.safeIndex(new BigNumber(651), new BigNumber((_667_v41).length));
        (_667_v41)[_index92] = _668_v42;
        let _670_v43;
        _670_v43 = _dafny.Seq.UnicodeFromString("lueejueq");
        let _671_v44;
        let _nw90 = Array((new BigNumber(8)).toNumber());
        _nw90[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_module.__default.fm45(_this.f14, new BigNumber((_601_v0).cardinality()), ((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))], globalState), _module.__default.safeIndex((_dafny.ZERO).minus(_603_v2), new BigNumber((_module.__default.fm45(_this.f14, new BigNumber((_601_v0).cardinality()), ((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))], globalState)).length)), (_606_v5)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_606_v5).length))]);
        _nw90[(_dafny.ONE).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(690)), ((_672_v5) => function (_673_i9) {
          return (_672_v5)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_672_v5).length))];
        })(_606_v5));
        _nw90[(new BigNumber(2)).toNumber()] = _670_v43;
        _nw90[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_670_v43, _670_v43);
        _nw90[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(584)), ((_674_v5) => function (_675_i10) {
          return (_674_v5)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_674_v5).length))];
        })(_606_v5)), _670_v43);
        _nw90[(new BigNumber(5)).toNumber()] = _670_v43;
        _nw90[(new BigNumber(6)).toNumber()] = _dafny.Seq.UnicodeFromString("cafhh");
        _nw90[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("wgxv"), _670_v43);
        _671_v44 = _nw90;
        let _index93 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_671_v44).length));
        (_671_v44)[_index93] = _670_v43;
        let _index94 = _module.__default.safeIndex(new BigNumber(651), new BigNumber((_667_v41).length));
        let _index95 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_671_v44).length));
        let _index96 = _module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length));
        let _rhs88 = ((((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))]) ? (_668_v42) : (_668_v42));
        let _rhs89 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("j"), _module.__default.fm45((_this).fm0(_602_v1, _this.f14, _603_v2, false, globalState), _603_v2, _this.f14, globalState));
        let _rhs90 = (_this).f10;
        let _rhs91 = ((_dafny.ZERO).minus(new BigNumber((_670_v43).length))).plus((_603_v2).plus(_603_v2));
        let _lhs64 = _667_v41;
        let _lhs65 = _module.__default.safeIndex(new BigNumber(651), new BigNumber((_667_v41).length));
        let _lhs66 = _671_v44;
        let _lhs67 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_671_v44).length));
        let _lhs68 = (_this).f15;
        let _lhs69 = _module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length));
        _lhs64[_lhs65] = _rhs88;
        _lhs66[_lhs67] = _rhs89;
        _lhs68[_lhs69] = _rhs90;
        _603_v2 = _rhs91;
        if ((((_615_v10).contains(_603_v2)) ? ((_615_v10).get(_603_v2)) : ((_603_v2).isEqualTo(_603_v2)))) {
          _613_v9 = _613_v9;
          let _676_v45;
          _676_v45 = _module.D1.create_DC3(_670_v43);
          let _pat_let_tv9 = _670_v43;
          _676_v45 = function (_pat_let8_0) {
            return function (_677_dt__update__tmp_h0) {
              return function (_pat_let9_0) {
                return function (_678_dt__update_hcf5_h0) {
                  return _module.D1.create_DC3(_678_dt__update_hcf5_h0);
                }(_pat_let9_0);
              }(_pat_let_tv9);
            }(_pat_let8_0);
          }(_module.D1.create_DC3(_670_v43));
          let _679_v46;
          let _nw91 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
          _679_v46 = _nw91;
          let _index97 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((_679_v46).length));
          (_679_v46)[_index97] = _603_v2;
          let _680_v47;
          let _init13 = ((_681_v2) => function (_682_i11) {
            return _module.D1.create_DC4(_this.f14, (_this).f10, _681_v2);
          })(_603_v2);
          let _nw92 = Array((new BigNumber(28)).toNumber());
          for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw92.length); _i0_13++) {
            _nw92[_i0_13] = _init13(new BigNumber(_i0_13));
          }
          _680_v47 = _nw92;
          let _index98 = _module.__default.safeIndex(new BigNumber(636), new BigNumber((_680_v47).length));
          (_680_v47)[_index98] = _module.D1.create_DC4(!(((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))]), _this.f14, _603_v2);
          let _683_v48;
          _683_v48 = _dafny.Seq.of(_668_v42, (_667_v41)[_module.__default.safeIndex(new BigNumber(651), new BigNumber((_667_v41).length))]);
          let _684_v49;
          _684_v49 = _module.D1.create_DC4(((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))], (((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))]) === (((((_this).f9).contains(((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))])) ? (((_this).f9).get(((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))])) : (false))), _603_v2);
          let _index99 = _module.__default.safeIndex(new BigNumber(651), new BigNumber((_667_v41).length));
          let _index100 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((_679_v46).length));
          let _index101 = _module.__default.safeIndex(new BigNumber(636), new BigNumber((_680_v47).length));
          let _rhs92 = (_this).fm1((_this).f9, globalState);
          let _rhs93 = (_683_v48)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements((_606_v5)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_606_v5).length))], (_606_v5)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_606_v5).length))], new _dafny.CodePoint('r'.codePointAt(0)))).length), new BigNumber((_683_v48).length))];
          let _rhs94 = _603_v2;
          let _rhs95 = _684_v49;
          let _rhs96 = new BigNumber((_602_v1).length);
          let _lhs70 = _667_v41;
          let _lhs71 = _module.__default.safeIndex(new BigNumber(651), new BigNumber((_667_v41).length));
          let _lhs72 = _679_v46;
          let _lhs73 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((_679_v46).length));
          let _lhs74 = _680_v47;
          let _lhs75 = _module.__default.safeIndex(new BigNumber(636), new BigNumber((_680_v47).length));
          _603_v2 = _rhs92;
          _lhs70[_lhs71] = _rhs93;
          _lhs72[_lhs73] = _rhs94;
          _lhs74[_lhs75] = _rhs95;
          _603_v2 = _rhs96;
          _603_v2 = _module.__default.safeDivisionInt(new BigNumber((_602_v1).length), ((_dafny.ZERO).minus((_this).fm42((_679_v46)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((_679_v46).length))], new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(((_671_v44)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_671_v44).length))]).length)),(_679_v46)[_module.__default.safeIndex(new BigNumber(108), new BigNumber((_679_v46).length))])).length), new BigNumber(-430), globalState))).plus(_603_v2));
          let _685_v50;
          let _686_v51;
          let _687_v52;
          let _out45;
          let _out46;
          let _out47;
          let _outcollector13 = (_this).m7(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("ou")).length), new BigNumber(((_671_v44)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_671_v44).length))]).length)), _603_v2, globalState);
          _out45 = _outcollector13[0];
          _out46 = _outcollector13[1];
          _out47 = _outcollector13[2];
          _685_v50 = _out45;
          _686_v51 = _out46;
          _687_v52 = _out47;
        } else {
          let _688_v53;
          _688_v53 = _dafny.Set.fromElements((_606_v5)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_606_v5).length))]);
          let _689_v54;
          _689_v54 = _dafny.Map.Empty.slice().updateUnsafe(_603_v2,_603_v2);
          let _690_v55;
          _690_v55 = _module.D1.create_DC4((_this).fm0(_dafny.Seq.of(((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))]), (_this).f10, new BigNumber(445), (_this).f10, globalState), _this.f14, _603_v2);
          let _691_v56;
          _691_v56 = _module.D5.create_DC14(_603_v2, _this.f14);
          let _692_v57;
          let _nw93 = Array((new BigNumber(12)).toNumber());
          _nw93[(_dafny.ZERO).toNumber()] = _603_v2;
          _nw93[(_dafny.ONE).toNumber()] = new BigNumber((_610_v7).length);
          _nw93[(new BigNumber(2)).toNumber()] = (_603_v2).multipliedBy(_603_v2);
          _nw93[(new BigNumber(3)).toNumber()] = _603_v2;
          _nw93[(new BigNumber(4)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_689_v54).length), _603_v2);
          _nw93[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(((_this.f14) ? (_603_v2) : (_603_v2)));
          _nw93[(new BigNumber(6)).toNumber()] = _603_v2;
          _nw93[(new BigNumber(7)).toNumber()] = _603_v2;
          _nw93[(new BigNumber(8)).toNumber()] = (_690_v55).dtor_cf8;
          _nw93[(new BigNumber(9)).toNumber()] = _603_v2;
          _nw93[(new BigNumber(10)).toNumber()] = (_691_v56).dtor_cf25;
          _nw93[(new BigNumber(11)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(364)), ((_693_v5) => function (_694_i12) {
            return (_693_v5)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_693_v5).length))];
          })(_606_v5))).length);
          _692_v57 = _nw93;
          let _index102 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_692_v57).length));
          (_692_v57)[_index102] = _603_v2;
          let _695_v58;
          _695_v58 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))]);
          let _696_v59;
          _696_v59 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))])).cardinality()),_695_v58);
          let _697_v61;
          _697_v61 = _dafny.MultiSet.fromElements((_this).f9);
          let _index103 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_692_v57).length));
          let _rhs97 = _module.__default.safeModuloInt((_dafny.ZERO).minus(_603_v2), _module.__default.safeDivisionInt(new BigNumber(((((_696_v59).contains(_603_v2)) ? ((_696_v59).get(_603_v2)) : (function () {
            let _coll37 = new _dafny.Map();
            for (const _compr_37 of (_697_v61).Elements) {
              let _698_v60 = _compr_37;
              if ((_697_v61).contains(_698_v60)) {
                _coll37.push([_698_v60,((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))]]);
              }
            }
            return _coll37;
          }()))).length), _603_v2));
          let _rhs98 = ((_688_v53).Difference(_688_v53)).Union((_dafny.Set.fromElements((_606_v5)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_606_v5).length))])).Difference(function () {
            let _coll38 = new _dafny.Set();
            for (const _compr_38 of (_670_v43).Elements) {
              let _699_v62 = _compr_38;
              if (_dafny.Seq.contains(_670_v43, _699_v62)) {
                _coll38.add(_699_v62);
              }
            }
            return _coll38;
          }()));
          let _rhs99 = _603_v2;
          let _rhs100 = (_dafny.ZERO).minus(new BigNumber((_601_v0).cardinality()));
          let _lhs76 = _692_v57;
          let _lhs77 = _module.__default.safeIndex(new BigNumber(698), new BigNumber((_692_v57).length));
          _603_v2 = _rhs97;
          _688_v53 = _rhs98;
          _603_v2 = _rhs99;
          _lhs76[_lhs77] = _rhs100;
          r0 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(67)), ((_700_v5) => function (_701_i13) {
            return (_700_v5)[_module.__default.safeIndex(new BigNumber(353), new BigNumber((_700_v5).length))];
          })(_606_v5));
          let _702_v63;
          _702_v63 = _module.D8.create_DC24(!(((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))]), _dafny.Set.fromElements(_692_v57), (_692_v57)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_692_v57).length))]);
          let _703_v64;
          _703_v64 = _dafny.Seq.of(_702_v63, _702_v63, _702_v63);
          let _704_v65;
          _704_v65 = _dafny.Set.fromElements(_692_v57);
          let _705_v66;
          _705_v66 = _dafny.Map.Empty.slice().updateUnsafe((_692_v57)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_692_v57).length))],_703_v64);
          let _706_v67;
          let _nw94 = Array((new BigNumber(21)).toNumber());
          _nw94[(_dafny.ZERO).toNumber()] = _703_v64;
          _nw94[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_703_v64, _dafny.Seq.of(_702_v63));
          _nw94[(new BigNumber(2)).toNumber()] = _703_v64;
          _nw94[(new BigNumber(3)).toNumber()] = _703_v64;
          _nw94[(new BigNumber(4)).toNumber()] = _703_v64;
          _nw94[(new BigNumber(5)).toNumber()] = _703_v64;
          _nw94[(new BigNumber(6)).toNumber()] = _dafny.Seq.of(_702_v63, _702_v63, _702_v63, _module.D8.create_DC24(_this.f14, _704_v65, (_692_v57)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_692_v57).length))]), _702_v63);
          _nw94[(new BigNumber(7)).toNumber()] = _dafny.Seq.of(_702_v63, _702_v63);
          _nw94[(new BigNumber(8)).toNumber()] = ((true) ? (_703_v64) : (_703_v64));
          _nw94[(new BigNumber(9)).toNumber()] = _703_v64;
          _nw94[(new BigNumber(10)).toNumber()] = ((((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))]) ? (_dafny.Seq.of(_702_v63, _702_v63, _702_v63)) : (_703_v64));
          _nw94[(new BigNumber(11)).toNumber()] = _dafny.Seq.update(_703_v64, _module.__default.safeIndex((_692_v57)[_module.__default.safeIndex(new BigNumber(698), new BigNumber((_692_v57).length))], new BigNumber((_703_v64).length)), _702_v63);
          _nw94[(new BigNumber(12)).toNumber()] = _703_v64;
          _nw94[(new BigNumber(13)).toNumber()] = _703_v64;
          _nw94[(new BigNumber(14)).toNumber()] = _703_v64;
          _nw94[(new BigNumber(15)).toNumber()] = _dafny.Seq.of(_702_v63, _702_v63, _702_v63);
          _nw94[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_703_v64, _dafny.Seq.of(_702_v63, _702_v63));
          _nw94[(new BigNumber(17)).toNumber()] = _703_v64;
          _nw94[(new BigNumber(18)).toNumber()] = _703_v64;
          _nw94[(new BigNumber(19)).toNumber()] = (((_705_v66).contains(_603_v2)) ? ((_705_v66).get(_603_v2)) : (_703_v64));
          _nw94[(new BigNumber(20)).toNumber()] = _703_v64;
          _706_v67 = _nw94;
          let _index104 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_706_v67).length));
          (_706_v67)[_index104] = _703_v64;
          let _707_v68;
          _707_v68 = _dafny.Map.Empty.slice().updateUnsafe(_601_v0,_dafny.Seq.of(_702_v63));
          let _pat_let_tv10 = _602_v1;
          let _pat_let_tv11 = _605_v4;
          let _pat_let_tv12 = globalState;
          let _pat_let_tv13 = _602_v1;
          let _pat_let_tv14 = _605_v4;
          let _pat_let_tv15 = globalState;
          let _index105 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_706_v67).length));
          (_706_v67)[_index105] = (((_707_v68).contains((_601_v0).update((function (_pat_let13_0) {
            return function (_711_dt__update__tmp_h1) {
              return function (_pat_let14_0) {
                return function (_712_dt__update_hcf7_h0) {
                  return function (_pat_let15_0) {
                    return function (_713_dt__update_hcf6_h0) {
                      return _module.D1.create_DC4(_713_dt__update_hcf6_h0, _712_dt__update_hcf7_h0, (_711_dt__update__tmp_h1).dtor_cf8);
                    }(_pat_let15_0);
                  }((_this).fm0(_pat_let_tv13, (_this).f10, new BigNumber((_pat_let_tv14).cardinality()), _this.f14, _pat_let_tv15));
                }(_pat_let14_0);
              }((_this).f10);
            }(_pat_let13_0);
          }(_690_v55)).dtor_cf6, _module.__default.abs((_this).fm1((_this).f9, globalState))))) ? ((_707_v68).get((_601_v0).update((function (_pat_let10_0) {
            return function (_708_dt__update__tmp_h2) {
              return function (_pat_let11_0) {
                return function (_709_dt__update_hcf7_h1) {
                  return function (_pat_let12_0) {
                    return function (_710_dt__update_hcf6_h1) {
                      return _module.D1.create_DC4(_710_dt__update_hcf6_h1, _709_dt__update_hcf7_h1, (_708_dt__update__tmp_h2).dtor_cf8);
                    }(_pat_let12_0);
                  }((_this).fm0(_pat_let_tv10, (_this).f10, new BigNumber((_pat_let_tv11).cardinality()), _this.f14, _pat_let_tv12));
                }(_pat_let11_0);
              }((_this).f10);
            }(_pat_let10_0);
          }(_690_v55)).dtor_cf6, _module.__default.abs((_this).fm1((_this).f9, globalState))))) : (_703_v64));
          let _714_v69;
          let _nw95 = Array((new BigNumber(7)).toNumber()).fill(_module.D8.Default());
          _714_v69 = _nw95;
          let _index106 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_714_v69).length));
          (_714_v69)[_index106] = _702_v63;
          let _index107 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_714_v69).length));
          (_714_v69)[_index107] = _702_v63;
          let _index108 = _module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length));
          ((_this).f15)[_index108] = !(_dafny.Seq.contains(_602_v1, ((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))]));
        }
        let _index109 = _module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length));
        let _rhs101 = (_this).f10;
        let _rhs102 = (new BigNumber(176)).minus(new BigNumber(((_this).f9).length));
        let _lhs78 = (_this).f15;
        let _lhs79 = _module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length));
        _lhs78[_lhs79] = _rhs101;
        _603_v2 = _rhs102;
        _603_v2 = (_dafny.ZERO).minus((_module.__default.safeDivisionInt(_603_v2, _603_v2)).minus(((((_this).f15)[_module.__default.safeIndex(new BigNumber(59), new BigNumber(((_this).f15).length))]) ? (_603_v2) : (_603_v2))));
      }
      let _715_v70;
      _715_v70 = _dafny.Seq.UnicodeFromString("ysxcsgfq");
      r0 = _715_v70;
      r1 = false;
      let _716_v71;
      _716_v71 = _dafny.Set.fromElements(_this.f14);
      let _717_v72;
      _717_v72 = _dafny.Seq.of((_716_v71).Union(_716_v71));
      r2 = _717_v72;
      r3 = _dafny.Seq.update(_715_v70, _module.__default.safeIndex((new BigNumber((_716_v71).length)).multipliedBy(_603_v2), new BigNumber((_715_v70).length)), new _dafny.CodePoint('v'.codePointAt(0)));
      return [r0, r1, r2, r3];
    }
    m1(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = false;
      let _718_v0;
      _718_v0 = new BigNumber(125);
      _718_v0 = new BigNumber(857);
      if ((_this.f14) && (!(_this.f14) || (_this.f14))) {
        let _index110 = _module.__default.safeIndex(new BigNumber(575), new BigNumber(((_this).f15).length));
        ((_this).f15)[_index110] = (_718_v0).isLessThan(new BigNumber(910));
        let _719_v1;
        _719_v1 = _dafny.Seq.of((true) === ((_this).f10), false, false);
        let _720_v2;
        _720_v2 = _dafny.MultiSet.fromElements(_718_v0);
        let _index111 = _module.__default.safeIndex(new BigNumber(575), new BigNumber(((_this).f15).length));
        let _rhs103 = true;
        let _rhs104 = _719_v1;
        let _rhs105 = (_this).f10;
        let _rhs106 = (_this).fm42(_718_v0, ((_dafny.ZERO).minus(_718_v0)).plus(_718_v0), new BigNumber((_720_v2).cardinality()), globalState);
        let _lhs80 = (_this).f15;
        let _lhs81 = _module.__default.safeIndex(new BigNumber(575), new BigNumber(((_this).f15).length));
        let _lhs82 = globalState;
        _lhs80[_lhs81] = _rhs103;
        _719_v1 = _rhs104;
        _lhs82.f8 = _rhs105;
        _718_v0 = _rhs106;
        _718_v0 = (_718_v0).minus(_718_v0);
        let _721_v3;
        _721_v3 = _dafny.Seq.UnicodeFromString("ll");
        let _722_v4;
        let _init14 = ((_723_v0) => function (_724_i0) {
          return (_724_i0).multipliedBy(_723_v0);
        })(_718_v0);
        let _nw96 = Array((new BigNumber(6)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw96.length); _i0_14++) {
          _nw96[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _722_v4 = _nw96;
        let _index112 = _module.__default.safeIndex(new BigNumber(727), new BigNumber((_722_v4).length));
        (_722_v4)[_index112] = (_dafny.ZERO).minus(_718_v0);
        let _index113 = _module.__default.safeIndex(new BigNumber(727), new BigNumber((_722_v4).length));
        let _rhs107 = _718_v0;
        let _rhs108 = _this.f14;
        let _rhs109 = (_this).f10;
        let _rhs110 = _721_v3;
        let _rhs111 = ((_718_v0).multipliedBy(new BigNumber(840))).multipliedBy(_module.__default.safeModuloInt(_718_v0, _718_v0));
        let _lhs83 = _this;
        let _lhs84 = globalState;
        let _lhs85 = _722_v4;
        let _lhs86 = _module.__default.safeIndex(new BigNumber(727), new BigNumber((_722_v4).length));
        _718_v0 = _rhs107;
        _lhs83.f14 = _rhs108;
        _lhs84.f8 = _rhs109;
        _721_v3 = _rhs110;
        _lhs85[_lhs86] = _rhs111;
        let _index114 = _module.__default.safeIndex(new BigNumber(727), new BigNumber((_722_v4).length));
        (_722_v4)[_index114] = ((_722_v4)[_module.__default.safeIndex(new BigNumber(727), new BigNumber((_722_v4).length))]).minus(_718_v0);
        let _725_v5;
        let _nw97 = new _module.C0();
        _nw97.__ctor();
        _725_v5 = _nw97;
      } else {
        let _726_v6;
        let _init15 = ((_727_v0) => function (_728_i1) {
          return _module.__default.safeDivisionInt(_728_i1, (_module.D5.create_DC14(_727_v0, (_this).f10)).dtor_cf25);
        })(_718_v0);
        let _nw98 = Array((new BigNumber(26)).toNumber());
        for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw98.length); _i0_15++) {
          _nw98[_i0_15] = _init15(new BigNumber(_i0_15));
        }
        _726_v6 = _nw98;
        _726_v6 = _726_v6;
        let _729_v7;
        _729_v7 = _dafny.Seq.of(_718_v0, _718_v0, new BigNumber(858));
        let _730_v8;
        _730_v8 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_module.__default.safeDivisionInt(_718_v0, new BigNumber((_729_v7).length)));
        _730_v8 = (_730_v8).update(true, (_dafny.ZERO).minus((_this).fm1((_this).f9, globalState)));
        let _731_v9;
        let _nw99 = Array((_dafny.ONE).toNumber()).fill([]);
        _731_v9 = _nw99;
        let _index115 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_731_v9).length));
        (_731_v9)[_index115] = _726_v6;
        let _732_v10;
        _732_v10 = _module.D3.create_DC7(_726_v6);
        let _index116 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_731_v9).length));
        (_731_v9)[_index116] = (_732_v10).dtor_cf11;
        let _733_v11;
        _733_v11 = _dafny.Seq.of(true);
        let _734_v12;
        _734_v12 = _module.D6.create_DC18(p0, new BigNumber((_733_v11).length), _718_v0);
        let _735_v13;
        _735_v13 = _module.D6.create_DC20(_module.D6.create_DC20(_734_v12));
        let _736_v14;
        _736_v14 = _dafny.Set.fromElements(function (_pat_let16_0) {
          return function (_737_dt__update__tmp_h0) {
            return function (_pat_let17_0) {
              return function (_738_dt__update_hcf39_h0) {
                return _module.D6.create_DC20(_738_dt__update_hcf39_h0);
              }(_pat_let17_0);
            }(_module.D6.create_DC16((_this).f9));
          }(_pat_let16_0);
        }(_735_v13));
        let _739_v15;
        _739_v15 = _dafny.Seq.of(_736_v14);
        let _arr0 = (_731_v9)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_731_v9).length))];
        let _index117 = _module.__default.safeIndex(new BigNumber(282), new BigNumber(((_731_v9)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_731_v9).length))]).length));
        _arr0[_index117] = _718_v0;
        let _arr1 = (_731_v9)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_731_v9).length))];
        let _index118 = _module.__default.safeIndex(new BigNumber(282), new BigNumber(((_731_v9)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_731_v9).length))]).length));
        let _index119 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_731_v9).length));
        let _rhs112 = ((_dafny.ZERO).minus((_718_v0).multipliedBy(_718_v0))).multipliedBy(_718_v0);
        let _rhs113 = _739_v15;
        let _rhs114 = (new BigNumber(195)).multipliedBy(_718_v0);
        let _rhs115 = (_732_v10).dtor_cf11;
        let _lhs87 = (_731_v9)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_731_v9).length))];
        let _lhs88 = _module.__default.safeIndex(new BigNumber(282), new BigNumber(((_731_v9)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_731_v9).length))]).length));
        let _lhs89 = _731_v9;
        let _lhs90 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_731_v9).length));
        _718_v0 = _rhs112;
        _739_v15 = _rhs113;
        _lhs87[_lhs88] = _rhs114;
        _lhs89[_lhs90] = _rhs115;
        (globalState).f8 = (_718_v0).isLessThan((((_this).f10) ? (((_731_v9)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_731_v9).length))])[_module.__default.safeIndex(new BigNumber(282), new BigNumber(((_731_v9)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_731_v9).length))]).length))]) : (new BigNumber((_module.__default.fm46((_this).f10, (_this).f10, (_this).f10, globalState)).cardinality()))));
      }
      let _740_v16;
      let _nw100 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Set.Empty);
      _740_v16 = _nw100;
      let _741_v17;
      _741_v17 = _dafny.Set.fromElements((_this).f10);
      let _index120 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_740_v16).length));
      (_740_v16)[_index120] = _741_v17;
      let _index121 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_740_v16).length));
      (_740_v16)[_index121] = _741_v17;
      let _742_v18;
      let _nw101 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Seq.of());
      _742_v18 = _nw101;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_742_v18).length))) {
        let _743_i2 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_743_i2)) && ((_743_i2).isLessThan(new BigNumber((_742_v18).length))))) {
          (_742_v18)[(_743_i2)] = _dafny.Seq.Concat(_dafny.Seq.of((_dafny.ZERO).minus((_dafny.ZERO).minus(_718_v0)), new BigNumber((_dafny.Seq.UnicodeFromString("upbj")).length), _718_v0), _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("kmewrgjgd")).length), _718_v0, new BigNumber(-56)));
        }
      }
      r1 = _this.f14;
      let _744_v19;
      _744_v19 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_718_v0);
      let _745_v20;
      _745_v20 = p0;
      let _746_v21;
      _746_v21 = _dafny.Map.Empty.slice().updateUnsafe(_745_v20,_744_v19);
      _744_v19 = (((_746_v21).contains(p0)) ? ((_746_v21).get(p0)) : (_dafny.Map.Empty.slice().updateUnsafe(_this.f14,new BigNumber(39))));
      let _747_v22;
      let _nw102 = Array((new BigNumber(7)).toNumber());
      _nw102[(_dafny.ZERO).toNumber()] = _718_v0;
      _nw102[(_dafny.ONE).toNumber()] = _718_v0;
      _nw102[(new BigNumber(2)).toNumber()] = _718_v0;
      _nw102[(new BigNumber(3)).toNumber()] = new BigNumber(-454);
      _nw102[(new BigNumber(4)).toNumber()] = _718_v0;
      _nw102[(new BigNumber(5)).toNumber()] = _718_v0;
      _nw102[(new BigNumber(6)).toNumber()] = _718_v0;
      _747_v22 = _nw102;
      let _748_v23;
      _748_v23 = _dafny.Seq.UnicodeFromString("by");
      let _749_v24;
      _749_v24 = _dafny.Map.Empty.slice().updateUnsafe(_747_v22,_748_v23);
      r0 = (_749_v24).update(_747_v22, _748_v23);
      let _750_v25;
      _750_v25 = _module.D1.create_DC3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(626)), ((_751_p0) => function (_752_i4) {
  return _751_p0;
})(p0)));
      r1 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(322)), ((_753_p0) => function (_754_i3) {
        return _753_p0;
      })(p0)), (_750_v25).dtor_cf5);
      return [r0, r1];
    }
    m7(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Map.Empty;
      let r2 = _module.D7.Default();
      r0 = (_this).f10;
      let _755_v0;
      let _nw103 = new _module.C0();
      _nw103.__ctor();
      _755_v0 = _nw103;
      let _756_i0;
      _756_i0 = _dafny.ZERO;
      L1: {
        while (!(_this.f14)) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_756_i0)) {
              break L1;
            }
            _756_i0 = (_756_i0).plus(_dafny.ONE);
            let _757_v1;
            _757_v1 = new BigNumber(891);
            _757_v1 = p0;
            let _758_v2;
            _758_v2 = new _dafny.CodePoint('o'.codePointAt(0));
            let _rhs116 = _758_v2;
            let _rhs117 = (_757_v1).multipliedBy(p1);
            let _rhs118 = p0;
            _758_v2 = _rhs116;
            _757_v1 = _rhs117;
            _757_v1 = _rhs118;
            _757_v1 = (_dafny.ZERO).minus(_757_v1);
            _757_v1 = (_dafny.ZERO).minus((_this).fm42(p0, p0, p0, globalState));
          }
        }
      }
      let _759_v3;
      _759_v3 = _dafny.Seq.of(p0);
      let _760_v4;
      _760_v4 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_759_v3)[_module.__default.safeIndex(new BigNumber(341), new BigNumber((_759_v3).length))]);
      let _761_v5;
      _761_v5 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,_760_v4);
      let _762_v6;
      _762_v6 = _dafny.Map.Empty.slice().updateUnsafe(_this.f14,p0);
      let _763_v7;
      _763_v7 = _dafny.MultiSet.fromElements(p1);
      let _764_v8;
      _764_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_763_v7).cardinality()),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("nixblma")).length)));
      _761_v5 = (_761_v5).update((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_this).f10,p0))).contains(_762_v6), _764_v8);
      r0 = (((_this).f10) || (_this.f14)) === (_this.f14);
      let _765_v9;
      let _nw104 = Array((new BigNumber(13)).toNumber());
      _nw104[(_dafny.ZERO).toNumber()] = _this.f14;
      _nw104[(_dafny.ONE).toNumber()] = (((_this).f10) ? ((_this).f10) : (_this.f14));
      _nw104[(new BigNumber(2)).toNumber()] = (_this).f10;
      _nw104[(new BigNumber(3)).toNumber()] = _this.f14;
      _nw104[(new BigNumber(4)).toNumber()] = _this.f14;
      _nw104[(new BigNumber(5)).toNumber()] = _this.f14;
      _nw104[(new BigNumber(6)).toNumber()] = (((_this).f10) ? ((_this).f10) : ((_this).f10));
      _nw104[(new BigNumber(7)).toNumber()] = (_this).f10;
      _nw104[(new BigNumber(8)).toNumber()] = (_this).f10;
      _nw104[(new BigNumber(9)).toNumber()] = (new BigNumber(-800)).isLessThan(new BigNumber(708));
      _nw104[(new BigNumber(10)).toNumber()] = !((_this).f10) || (_this.f14);
      _nw104[(new BigNumber(11)).toNumber()] = !((_this).f10) || (_this.f14);
      _nw104[(new BigNumber(12)).toNumber()] = _this.f14;
      _765_v9 = _nw104;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_765_v9).length))) {
        let _766_i1 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_766_i1)) && ((_766_i1).isLessThan(new BigNumber((_765_v9).length))))) {
          (_765_v9)[(_766_i1)] = !((new BigNumber((function () {
            let _coll39 = new _dafny.Set();
            for (const _compr_39 of (_dafny.Set.fromElements(p1)).Elements) {
              let _767_v10 = _compr_39;
              if ((_dafny.Set.fromElements(p1)).contains(_767_v10)) {
                _coll39.add((_767_v10).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(!(!(true))),(_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(324)), function (_768_i2) {
                  return new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length))).cardinality());
                })).length))))))).length)));
              }
            }
            return _coll39;
          }()).length)).isEqualTo(p0)) || ((_this).f10);
        }
      }
      r0 = _this.f14;
      r1 = _762_v6;
      let _769_v11;
      _769_v11 = _module.D7.create_DC22();
      r2 = _769_v11;
      return [r0, r1, r2];
    }
    get f15() {
      let _this = this;
      return _this._f15;
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
    fm27(p0, p1, p2, globalState) {
      let _this = this;
      return (!(false)) && ((new BigNumber(-159)).isEqualTo(new BigNumber(821)));
    };
    fm28(p0, p1, p2, p3, globalState) {
      let _this = this;
      return new BigNumber((function (_source15) {
        if (_source15.is_DC3) {
          let _770___mcc_h0 = (_source15).cf5;
          let _771_cf5 = _770___mcc_h0;
          return (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('f'.codePointAt(0)),!(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),false));
        } else if (_source15.is_DC4) {
          let _772___mcc_h1 = (_source15).cf6;
          let _773___mcc_h2 = (_source15).cf7;
          let _774___mcc_h3 = (_source15).cf8;
          let _775_cf8 = _774___mcc_h3;
          let _776_cf7 = _773___mcc_h2;
          let _777_cf6 = _772___mcc_h1;
          return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('y'.codePointAt(0)),!(_776_cf7));
        } else if (_source15.is_DC2) {
          let _778___mcc_h4 = (_source15).cf4;
          let _779_cf4 = _778___mcc_h4;
          return (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),true));
        } else {
          let _780___mcc_h5 = (_source15).cf9;
          let _781_cf9 = _780___mcc_h5;
          return (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),!(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),false));
        }
      }(_module.D1.create_DC3(_dafny.Seq.Create(_module.__default.abs(new BigNumber(725)), function (_782_i0) {
  return new _dafny.CodePoint('q'.codePointAt(0));
})))).length);
    };
    m0(globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = false;
      let r2 = _dafny.Seq.of();
      let r3 = _dafny.Seq.UnicodeFromString("");
      let _783_v0;
      _783_v0 = true;
      r1 = !(_783_v0) || (_783_v0);
      let _784_v1;
      let _nw105 = Array((new BigNumber(21)).toNumber()).fill(_dafny.ZERO);
      _784_v1 = _nw105;
      let _785_v2;
      _785_v2 = _dafny.Map.Empty.slice().updateUnsafe(_784_v1,_783_v0);
      _785_v2 = (_785_v2).update(_784_v1, (_this).fm27(_783_v0, _783_v0, _module.D3.create_DC8(_783_v0, _783_v0), globalState));
      let _786_v3;
      _786_v3 = new BigNumber(838);
      let _787_v4;
      _787_v4 = new _dafny.CodePoint('q'.codePointAt(0));
      let _788_v5;
      _788_v5 = _dafny.Seq.UnicodeFromString("gkun");
      let _789_v6;
      _789_v6 = _module.D1.create_DC3(_788_v5);
      let _790_v7;
      _790_v7 = _dafny.Map.Empty.slice().updateUnsafe(_789_v6,_786_v3);
      let _index122 = _module.__default.safeIndex(new BigNumber(968), new BigNumber((_784_v1).length));
      (_784_v1)[_index122] = new BigNumber(((_790_v7).update(_789_v6, _786_v3)).length);
      let _791_v8;
      _791_v8 = _module.D3.create_DC8(_783_v0, _783_v0);
      let _792_v9;
      _792_v9 = _dafny.MultiSet.fromElements((_this).fm27(_783_v0, _783_v0, _791_v8, globalState));
      let _index123 = _module.__default.safeIndex(new BigNumber(968), new BigNumber((_784_v1).length));
      let _rhs119 = new BigNumber(((_792_v9).Difference(_dafny.MultiSet.fromElements(true))).cardinality());
      let _rhs120 = _783_v0;
      let _rhs121 = (_788_v5)[_module.__default.safeIndex(_786_v3, new BigNumber((_788_v5).length))];
      let _rhs122 = (_786_v3).minus(_786_v3);
      let _lhs91 = globalState;
      let _lhs92 = _784_v1;
      let _lhs93 = _module.__default.safeIndex(new BigNumber(968), new BigNumber((_784_v1).length));
      _786_v3 = _rhs119;
      _lhs91.f8 = _rhs120;
      _787_v4 = _rhs121;
      _lhs92[_lhs93] = _rhs122;
      let _793_v10;
      let _init16 = ((_794_v0) => function (_795_i0) {
        return _794_v0;
      })(_783_v0);
      let _nw106 = Array((new BigNumber(10)).toNumber());
      for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw106.length); _i0_16++) {
        _nw106[_i0_16] = _init16(new BigNumber(_i0_16));
      }
      _793_v10 = _nw106;
      let _index124 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_793_v10).length));
      (_793_v10)[_index124] = _783_v0;
      let _796_v11;
      _796_v11 = _dafny.MultiSet.fromElements((_784_v1)[_module.__default.safeIndex(new BigNumber(968), new BigNumber((_784_v1).length))], (_784_v1)[_module.__default.safeIndex(new BigNumber(968), new BigNumber((_784_v1).length))]);
      let _797_v12;
      _797_v12 = _dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber((_788_v5).length), _786_v3), _796_v11);
      let _index125 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_793_v10).length));
      (_793_v10)[_index125] = !((_797_v12)[_module.__default.safeIndex(_786_v3, new BigNumber((_797_v12).length))]).contains(_786_v3);
      let _hi0 = (_786_v3).plus(_786_v3);
      for (let _798_i1 = ((_784_v1)[_module.__default.safeIndex(new BigNumber(968), new BigNumber((_784_v1).length))]).minus(new BigNumber((function () {
        let _coll40 = new _dafny.Set();
        for (const _compr_40 of _dafny.IntegerRange(new BigNumber(-59), new BigNumber(873))) {
          let _818_v13 = _compr_40;
          if (((new BigNumber(-59)).isLessThanOrEqualTo(_818_v13)) && ((_818_v13).isLessThan(new BigNumber(873)))) {
            _coll40.add((_818_v13).multipliedBy((_dafny.ZERO).minus((_784_v1)[_module.__default.safeIndex(new BigNumber(968), new BigNumber((_784_v1).length))])));
          }
        }
        return _coll40;
      }()).length)); _798_i1.isLessThan(_hi0); _798_i1 = _798_i1.plus(_dafny.ONE)) {
        let _799_v14;
        _799_v14 = _module.D1.create_DC4(_783_v0, _783_v0, _798_i1);
        if ((_799_v14).dtor_cf7) {
          let _index126 = _module.__default.safeIndex(new BigNumber(968), new BigNumber((_784_v1).length));
          (_784_v1)[_index126] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_798_i1));
          let _800_v15;
          let _nw107 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Map.Empty);
          _800_v15 = _nw107;
          let _nw108 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Map.Empty);
          _800_v15 = _nw108;
          let _801_v16;
          _801_v16 = _dafny.Seq.of(false);
          let _rhs123 = _801_v16;
          let _rhs124 = _783_v0;
          let _rhs125 = (_dafny.ZERO).minus(((_784_v1)[_module.__default.safeIndex(new BigNumber(968), new BigNumber((_784_v1).length))]).plus((_784_v1)[_module.__default.safeIndex(new BigNumber(968), new BigNumber((_784_v1).length))]));
          let _rhs126 = _dafny.Seq.Concat(_788_v5, _788_v5);
          _801_v16 = _rhs123;
          _783_v0 = _rhs124;
          _786_v3 = _rhs125;
          _788_v5 = _rhs126;
          let _802_v17;
          let _nw109 = new _module.C0();
          _nw109.__ctor();
          _802_v17 = _nw109;
          let _index127 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_793_v10).length));
          (_793_v10)[_index127] = (_this).fm27((_798_i1).isLessThanOrEqualTo((_784_v1)[_module.__default.safeIndex(new BigNumber(968), new BigNumber((_784_v1).length))]), false, _791_v8, globalState);
        } else {
          let _803_v18;
          let _init17 = ((_804_v14) => function (_805_i2) {
            return _804_v14;
          })(_799_v14);
          let _nw110 = Array((new BigNumber(16)).toNumber());
          for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw110.length); _i0_17++) {
            _nw110[_i0_17] = _init17(new BigNumber(_i0_17));
          }
          _803_v18 = _nw110;
          let _index128 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_803_v18).length));
          (_803_v18)[_index128] = _799_v14;
          let _index129 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_803_v18).length));
          (_803_v18)[_index129] = _799_v14;
          let _806_v19;
          let _nw111 = new _module.C0();
          _nw111.__ctor();
          _806_v19 = _nw111;
          let _index130 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_793_v10).length));
          (_793_v10)[_index130] = _783_v0;
          let _807_v20;
          _807_v20 = _dafny.MultiSet.fromElements(_787_v4);
          let _808_v21;
          _808_v21 = _dafny.Map.Empty.slice().updateUnsafe(_783_v0,_807_v20);
          let _809_v22;
          _809_v22 = _dafny.Set.fromElements(_798_i1);
          let _810_v23;
          _810_v23 = _dafny.Set.fromElements((_793_v10)[_module.__default.safeIndex(new BigNumber(470), new BigNumber((_793_v10).length))], (_793_v10)[_module.__default.safeIndex(new BigNumber(470), new BigNumber((_793_v10).length))]);
          let _811_v24;
          _811_v24 = _dafny.Seq.of(true, false);
          let _812_v25;
          _812_v25 = _dafny.Seq.of((_811_v24)[_module.__default.safeIndex(_798_i1, new BigNumber((_811_v24).length))]);
          let _813_v26;
          let _nw112 = Array((new BigNumber(19)).toNumber());
          _nw112[(_dafny.ZERO).toNumber()] = (_808_v21).Merge(_808_v21);
          _nw112[(_dafny.ONE).toNumber()] = _808_v21;
          _nw112[(new BigNumber(2)).toNumber()] = _808_v21;
          _nw112[(new BigNumber(3)).toNumber()] = (_808_v21).Merge(_808_v21);
          _nw112[(new BigNumber(4)).toNumber()] = _808_v21;
          _nw112[(new BigNumber(5)).toNumber()] = _808_v21;
          _nw112[(new BigNumber(6)).toNumber()] = _808_v21;
          _nw112[(new BigNumber(7)).toNumber()] = _808_v21;
          _nw112[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_783_v0,_807_v20);
          _nw112[(new BigNumber(9)).toNumber()] = _module.__default.fm29(_783_v0, true, _module.__default.fm30(globalState), _809_v22, globalState);
          _nw112[(new BigNumber(10)).toNumber()] = _module.__default.fm29(_783_v0, _783_v0, _810_v23, _809_v22, globalState);
          _nw112[(new BigNumber(11)).toNumber()] = (_808_v21).update(_783_v0, _807_v20);
          _nw112[(new BigNumber(12)).toNumber()] = (((_793_v10)[_module.__default.safeIndex(new BigNumber(470), new BigNumber((_793_v10).length))]) ? (_dafny.Map.Empty.slice().updateUnsafe(!((_812_v25)[_module.__default.safeIndex(_798_i1, new BigNumber((_812_v25).length))]),_807_v20)) : (_dafny.Map.Empty.slice().updateUnsafe(_783_v0,_807_v20)));
          _nw112[(new BigNumber(13)).toNumber()] = (_module.__default.fm29((_793_v10)[_module.__default.safeIndex(new BigNumber(470), new BigNumber((_793_v10).length))], _783_v0, _810_v23, _809_v22, globalState)).Merge(_808_v21);
          _nw112[(new BigNumber(14)).toNumber()] = (_808_v21).Merge(_dafny.Map.Empty.slice().updateUnsafe((_793_v10)[_module.__default.safeIndex(new BigNumber(470), new BigNumber((_793_v10).length))],_807_v20));
          _nw112[(new BigNumber(15)).toNumber()] = _808_v21;
          _nw112[(new BigNumber(16)).toNumber()] = (_module.D4.create_DC10(_dafny.Map.Empty.slice().updateUnsafe(_783_v0,_807_v20))).dtor_cf15;
          _nw112[(new BigNumber(17)).toNumber()] = _808_v21;
          _nw112[(new BigNumber(18)).toNumber()] = _808_v21;
          _813_v26 = _nw112;
          let _index131 = _module.__default.safeIndex(new BigNumber(489), new BigNumber((_813_v26).length));
          (_813_v26)[_index131] = _dafny.Map.Empty.slice().updateUnsafe(true,_dafny.MultiSet.fromElements(_787_v4));
          let _index132 = _module.__default.safeIndex(new BigNumber(489), new BigNumber((_813_v26).length));
          (_813_v26)[_index132] = _808_v21;
          _787_v4 = _787_v4;
        }
        let _814_v27;
        let _nw113 = new _module.C0();
        _nw113.__ctor();
        _814_v27 = _nw113;
        let _815_v28;
        let _nw114 = new _module.C0();
        _nw114.__ctor();
        _815_v28 = _nw114;
        let _816_v29;
        _816_v29 = _dafny.Seq.of((_793_v10)[_module.__default.safeIndex(new BigNumber(470), new BigNumber((_793_v10).length))]);
        let _817_v30;
        _817_v30 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-256),_816_v29);
        _817_v30 = (_817_v30).update(new BigNumber((_796_v11).cardinality()), _816_v29);
      }
      let _hi1 = (_784_v1)[_module.__default.safeIndex(new BigNumber(968), new BigNumber((_784_v1).length))];
      for (let _819_i3 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(838)), ((_828_v4) => function (_829_i4) {
        return _828_v4;
      })(_787_v4))).length)); _819_i3.isLessThan(_hi1); _819_i3 = _819_i3.plus(_dafny.ONE)) {
        _792_v9 = _792_v9;
        let _index133 = _module.__default.safeIndex(new BigNumber(470), new BigNumber((_793_v10).length));
        (_793_v10)[_index133] = (_793_v10)[_module.__default.safeIndex(new BigNumber(470), new BigNumber((_793_v10).length))];
        (globalState).f8 = ((_784_v1)[_module.__default.safeIndex(new BigNumber(968), new BigNumber((_784_v1).length))]).isLessThanOrEqualTo((_dafny.ZERO).minus(_819_i3));
        let _820_v31;
        _820_v31 = _dafny.Set.fromElements(_786_v3, _819_i3);
        let _821_v32;
        _821_v32 = _module.D4.create_DC12(!(_820_v31).contains(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-450)), ((_822_v4) => function (_823_i5) {
  return _822_v4;
})(_787_v4))).length)), _module.__default.safeDivisionInt(new BigNumber(303), _819_i3), _787_v4, _786_v3, new BigNumber(754));
        let _824_v33;
        _824_v33 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(338)), ((_825_v4) => function (_826_i6) {
          return _825_v4;
        })(_787_v4)),false);
        let _827_v34;
        _827_v34 = _module.D1.create_DC4(_783_v0, _783_v0, (_dafny.ZERO).minus(new BigNumber((_824_v33).length)));
        let _index134 = _module.__default.safeIndex(new BigNumber(968), new BigNumber((_784_v1).length));
        let _index135 = _module.__default.safeIndex(new BigNumber(968), new BigNumber((_784_v1).length));
        let _index136 = _module.__default.safeIndex(new BigNumber(968), new BigNumber((_784_v1).length));
        let _rhs127 = _821_v32;
        let _rhs128 = (_819_i3).multipliedBy((_this).fm28((_793_v10)[_module.__default.safeIndex(new BigNumber(470), new BigNumber((_793_v10).length))], _827_v34, _786_v3, new BigNumber(-108), globalState));
        let _rhs129 = _module.__default.safeDivisionInt(new BigNumber(261), ((_this).fm28(_783_v0, _827_v34, (_784_v1)[_module.__default.safeIndex(new BigNumber(968), new BigNumber((_784_v1).length))], _819_i3, globalState)).minus(_819_i3));
        let _rhs130 = _786_v3;
        let _lhs94 = _784_v1;
        let _lhs95 = _module.__default.safeIndex(new BigNumber(968), new BigNumber((_784_v1).length));
        let _lhs96 = _784_v1;
        let _lhs97 = _module.__default.safeIndex(new BigNumber(968), new BigNumber((_784_v1).length));
        let _lhs98 = _784_v1;
        let _lhs99 = _module.__default.safeIndex(new BigNumber(968), new BigNumber((_784_v1).length));
        _821_v32 = _rhs127;
        _lhs94[_lhs95] = _rhs128;
        _lhs96[_lhs97] = _rhs129;
        _lhs98[_lhs99] = _rhs130;
      }
      r0 = _dafny.Seq.update(_dafny.Seq.Concat(_788_v5, _788_v5), _module.__default.safeIndex((_784_v1)[_module.__default.safeIndex(new BigNumber(968), new BigNumber((_784_v1).length))], new BigNumber((_dafny.Seq.Concat(_788_v5, _788_v5)).length)), _787_v4);
      r1 = _783_v0;
      let _830_v35;
      _830_v35 = _dafny.Set.fromElements(false, (_793_v10)[_module.__default.safeIndex(new BigNumber(470), new BigNumber((_793_v10).length))]);
      r2 = _dafny.Seq.of(_830_v35);
      r3 = _dafny.Seq.Concat(_788_v5, _dafny.Seq.UnicodeFromString("letpy"));
      return [r0, r1, r2, r3];
    }
    m1(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = false;
      let _831_v0;
      let _nw115 = new _module.C0();
      _nw115.__ctor();
      _831_v0 = _nw115;
      let _832_v1;
      _832_v1 = true;
      let _833_v2;
      _833_v2 = new BigNumber(248);
      let _834_v3;
      _834_v3 = _dafny.Map.Empty.slice().updateUnsafe(_833_v2,_832_v1);
      let _835_v4;
      _835_v4 = _module.D4.create_DC12(_832_v1, _833_v2, _module.__default.fm31(_832_v1, _834_v3, globalState), _833_v2, _833_v2);
      let _pat_let_tv16 = _832_v1;
      let _pat_let_tv17 = _832_v1;
      if (function (_source16) {
        if (_source16.is_DC11) {
          let _836___mcc_h0 = (_source16).cf16;
          let _837___mcc_h1 = (_source16).cf17;
          let _838___mcc_h2 = (_source16).cf18;
          let _839_cf18 = _838___mcc_h2;
          let _840_cf17 = _837___mcc_h1;
          let _841_cf16 = _836___mcc_h0;
          return _pat_let_tv16;
        } else if (_source16.is_DC12) {
          let _842___mcc_h3 = (_source16).cf19;
          let _843___mcc_h4 = (_source16).cf20;
          let _844___mcc_h5 = (_source16).cf21;
          let _845___mcc_h6 = (_source16).cf22;
          let _846___mcc_h7 = (_source16).cf23;
          let _847_cf23 = _846___mcc_h7;
          let _848_cf22 = _845___mcc_h6;
          let _849_cf21 = _844___mcc_h5;
          let _850_cf20 = _843___mcc_h4;
          let _851_cf19 = _842___mcc_h3;
          return (_848_cf22).isLessThan(_850_cf20);
        } else {
          let _852___mcc_h8 = (_source16).cf15;
          let _853_cf15 = _852___mcc_h8;
          return _pat_let_tv17;
        }
      }(_835_v4)) {
        let _854_v6;
        let _init18 = ((_855_v1, _856_v2) => function (_857_i0) {
          return (_module.D4.create_DC11(new BigNumber(-957), function () {
  let _coll41 = new _dafny.Set();
  for (const _compr_41 of _dafny.IntegerRange(new BigNumber(-371), new BigNumber(330))) {
    let _858_v5 = _compr_41;
    if (((new BigNumber(-371)).isLessThanOrEqualTo(_858_v5)) && ((_858_v5).isLessThan(new BigNumber(330)))) {
      _coll41.add((_858_v5).plus(_856_v2));
    }
  }
  return _coll41;
}(), _855_v1)).dtor_cf18;
        })(_832_v1, _833_v2);
        let _nw116 = Array((new BigNumber(4)).toNumber());
        for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw116.length); _i0_18++) {
          _nw116[_i0_18] = _init18(new BigNumber(_i0_18));
        }
        _854_v6 = _nw116;
        let _index137 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_854_v6).length));
        (_854_v6)[_index137] = _832_v1;
        let _index138 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_854_v6).length));
        (_854_v6)[_index138] = _832_v1;
        let _index139 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_854_v6).length));
        (_854_v6)[_index139] = _832_v1;
        let _859_v7;
        let _nw117 = new _module.C0();
        _nw117.__ctor();
        _859_v7 = _nw117;
        let _860_v8;
        _860_v8 = _dafny.Seq.of(_832_v1);
        if (((_860_v8)[_module.__default.safeIndex(_833_v2, new BigNumber((_860_v8).length))]) || ((((_854_v6)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_854_v6).length))]) ? (_832_v1) : ((_854_v6)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_854_v6).length))])))) {
          let _index140 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_854_v6).length));
          (_854_v6)[_index140] = !(_832_v1);
          let _861_v9;
          _861_v9 = _dafny.Seq.UnicodeFromString("ewyay");
          _861_v9 = _dafny.Seq.Concat(_861_v9, _dafny.Seq.Concat(_861_v9, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-365)), ((_862_p0) => function (_863_i1) {
            return _862_p0;
          })(p0))));
          let _864_v10;
          _864_v10 = _dafny.Map.Empty.slice().updateUnsafe(_854_v6,_832_v1);
          let _865_v11;
          _865_v11 = _module.D5.create_DC13(_854_v6);
          _864_v10 = (_864_v10).update((_865_v11).dtor_cf24, (_854_v6)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_854_v6).length))]);
          let _866_v12;
          _866_v12 = _dafny.MultiSet.fromElements(false, _832_v1, true, (_854_v6)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_854_v6).length))], ((_dafny.ZERO).minus(_833_v2)).isEqualTo(_833_v2));
          _866_v12 = (_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_860_v8, _860_v8))).Intersect(_866_v12);
          let _867_v13;
          _867_v13 = _dafny.Map.Empty.slice().updateUnsafe(_833_v2,_831_v0);
          let _868_v14;
          _868_v14 = _dafny.Map.Empty.slice().updateUnsafe(_833_v2,new BigNumber((_867_v13).length));
          let _869_v15;
          _869_v15 = _module.D5.create_DC14(_833_v2, (_854_v6)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_854_v6).length))]);
          let _870_v16;
          _870_v16 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_854_v6)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_854_v6).length))],_dafny.Set.fromElements(_859_v7, _831_v0)),_dafny.Map.Empty.slice().updateUnsafe((_869_v15).dtor_cf25,_833_v2));
          let _871_v17;
          _871_v17 = _dafny.Set.fromElements(_831_v0);
          let _872_v18;
          _872_v18 = _dafny.Map.Empty.slice().updateUnsafe((_854_v6)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_854_v6).length))],_871_v17);
          let _873_v19;
          _873_v19 = _module.D1.create_DC4(_832_v1, (_854_v6)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_854_v6).length))], _833_v2);
          let _874_v20;
          _874_v20 = _dafny.MultiSet.fromElements(_833_v2, _833_v2);
          let _875_v21;
          let _nw118 = Array((new BigNumber(13)).toNumber());
          _nw118[(_dafny.ZERO).toNumber()] = _868_v14;
          _nw118[(_dafny.ONE).toNumber()] = (((_870_v16).contains(_872_v18)) ? ((_870_v16).get(_872_v18)) : (_868_v14));
          _nw118[(new BigNumber(2)).toNumber()] = _868_v14;
          _nw118[(new BigNumber(3)).toNumber()] = ((_868_v14).update(_833_v2, _833_v2)).update(_833_v2, _833_v2);
          _nw118[(new BigNumber(4)).toNumber()] = _868_v14;
          _nw118[(new BigNumber(5)).toNumber()] = _868_v14;
          _nw118[(new BigNumber(6)).toNumber()] = (_868_v14).update((_this).fm28((_854_v6)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_854_v6).length))], _873_v19, new BigNumber((_874_v20).cardinality()), _833_v2, globalState), _833_v2);
          _nw118[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_833_v2,_833_v2);
          _nw118[(new BigNumber(8)).toNumber()] = (_868_v14).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(136),new BigNumber(457)));
          _nw118[(new BigNumber(9)).toNumber()] = _868_v14;
          _nw118[(new BigNumber(10)).toNumber()] = _868_v14;
          _nw118[(new BigNumber(11)).toNumber()] = _868_v14;
          _nw118[(new BigNumber(12)).toNumber()] = _868_v14;
          _875_v21 = _nw118;
          _875_v21 = _875_v21;
        } else {
          let _876_v22;
          _876_v22 = _dafny.Seq.UnicodeFromString("nhxyjlt");
          let _index141 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_854_v6).length));
          (_854_v6)[_index141] = (_module.__default.safeModuloInt(new BigNumber((_876_v22).length), _833_v2)).isLessThan(new BigNumber((_dafny.Set.fromElements(_833_v2)).length));
          let _877_v23;
          _877_v23 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("hpdclvpfa"));
          _876_v22 = _dafny.Seq.Concat(_876_v22, (_877_v23)[_module.__default.safeIndex((_dafny.ZERO).minus(_833_v2), new BigNumber((_877_v23).length))]);
          let _878_v24;
          _878_v24 = _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(739)), ((_879_p0) => function (_880_i2) {
            return _879_p0;
          })(p0)));
          let _881_v25;
          _881_v25 = _dafny.Map.Empty.slice().updateUnsafe(_860_v8,false);
          let _882_v26;
          _882_v26 = _dafny.Map.Empty.slice().updateUnsafe(_833_v2,p0);
          let _883_v27;
          _883_v27 = _dafny.MultiSet.fromElements(_833_v2, _833_v2, _833_v2, new BigNumber((_882_v26).length));
          let _884_v28;
          _884_v28 = _dafny.MultiSet.fromElements(false, false);
          let _885_v29;
          _885_v29 = _dafny.Map.Empty.slice().updateUnsafe((_854_v6)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_854_v6).length))],!(_832_v1));
          let _886_v30;
          _886_v30 = _dafny.Seq.of(_833_v2);
          let _887_v31;
          let _nw119 = Array((new BigNumber(24)).toNumber());
          _nw119[(_dafny.ZERO).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_dafny.ZERO).minus(_833_v2)), _833_v2);
          _nw119[(_dafny.ONE).toNumber()] = _module.__default.safeDivisionInt(_833_v2, _833_v2);
          _nw119[(new BigNumber(2)).toNumber()] = new BigNumber(705);
          _nw119[(new BigNumber(3)).toNumber()] = _833_v2;
          _nw119[(new BigNumber(4)).toNumber()] = new BigNumber(559);
          _nw119[(new BigNumber(5)).toNumber()] = _833_v2;
          _nw119[(new BigNumber(6)).toNumber()] = _833_v2;
          _nw119[(new BigNumber(7)).toNumber()] = (_833_v2).minus(_833_v2);
          _nw119[(new BigNumber(8)).toNumber()] = new BigNumber((_878_v24).length);
          _nw119[(new BigNumber(9)).toNumber()] = _833_v2;
          _nw119[(new BigNumber(10)).toNumber()] = new BigNumber(((_881_v25).Merge(_881_v25)).length);
          _nw119[(new BigNumber(11)).toNumber()] = (((_883_v27).contains(new BigNumber((_884_v28).cardinality()))) ? ((_883_v27).get(new BigNumber((_884_v28).cardinality()))) : (_833_v2));
          _nw119[(new BigNumber(12)).toNumber()] = _833_v2;
          _nw119[(new BigNumber(13)).toNumber()] = ((_dafny.ZERO).minus(_833_v2)).multipliedBy(_833_v2);
          _nw119[(new BigNumber(14)).toNumber()] = new BigNumber(779);
          _nw119[(new BigNumber(15)).toNumber()] = new BigNumber((_860_v8).length);
          _nw119[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus((new BigNumber((_885_v29).length)).plus((_886_v30)[_module.__default.safeIndex(_833_v2, new BigNumber((_886_v30).length))]));
          _nw119[(new BigNumber(17)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("qryuduwux")).length);
          _nw119[(new BigNumber(18)).toNumber()] = _833_v2;
          _nw119[(new BigNumber(19)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_886_v30).length));
          _nw119[(new BigNumber(20)).toNumber()] = _833_v2;
          _nw119[(new BigNumber(21)).toNumber()] = _833_v2;
          _nw119[(new BigNumber(22)).toNumber()] = _833_v2;
          _nw119[(new BigNumber(23)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.UnicodeFromString("loerye")).length), _833_v2);
          _887_v31 = _nw119;
          let _888_v32;
          _888_v32 = _dafny.Seq.of(_887_v31, _887_v31, _887_v31);
          _887_v31 = (_888_v32)[_module.__default.safeIndex((_833_v2).minus(_833_v2), new BigNumber((_888_v32).length))];
          let _index142 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_887_v31).length));
          (_887_v31)[_index142] = _833_v2;
          let _889_v33;
          _889_v33 = _dafny.Set.fromElements((_854_v6)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_854_v6).length))], false);
          let _890_v34;
          _890_v34 = _dafny.Map.Empty.slice().updateUnsafe(_833_v2,_889_v33);
          let _index143 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_887_v31).length));
          let _index144 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_854_v6).length));
          let _rhs131 = _876_v22;
          let _rhs132 = (_833_v2).multipliedBy(((_832_v1) ? (_833_v2) : (new BigNumber(865))));
          let _rhs133 = !((_854_v6)[_module.__default.safeIndex(new BigNumber(306), new BigNumber((_854_v6).length))]);
          let _rhs134 = _890_v34;
          let _rhs135 = _833_v2;
          let _lhs100 = _887_v31;
          let _lhs101 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_887_v31).length));
          let _lhs102 = _854_v6;
          let _lhs103 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_854_v6).length));
          _876_v22 = _rhs131;
          _lhs100[_lhs101] = _rhs132;
          _lhs102[_lhs103] = _rhs133;
          _890_v34 = _rhs134;
          _833_v2 = _rhs135;
          let _891_v35;
          _891_v35 = _dafny.Map.Empty.slice().updateUnsafe(p0,_833_v2);
          let _index145 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_887_v31).length));
          (_887_v31)[_index145] = (((_891_v35).contains(_module.__default.fm31(!(true), _834_v3, globalState))) ? ((_891_v35).get(_module.__default.fm31(!(true), _834_v3, globalState))) : (new BigNumber(111)));
        }
        let _892_v36;
        let _nw120 = new _module.C0();
        _nw120.__ctor();
        _892_v36 = _nw120;
      } else {
        (globalState).f8 = _832_v1;
        (globalState).f8 = _832_v1;
        let _893_v37;
        let _nw121 = new _module.C1();
        _nw121.__ctor(_832_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(173)), function (_894_i3) {
          return new _dafny.CodePoint('f'.codePointAt(0));
        }));
        _893_v37 = _nw121;
        let _895_v38;
        let _init19 = ((_896_v1) => function (_897_i4) {
          return _896_v1;
        })(_832_v1);
        let _nw122 = Array((new BigNumber(13)).toNumber());
        for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw122.length); _i0_19++) {
          _nw122[_i0_19] = _init19(new BigNumber(_i0_19));
        }
        _895_v38 = _nw122;
        let _898_v39;
        _898_v39 = _dafny.Map.Empty.slice().updateUnsafe(true,_832_v1);
        let _899_v40;
        let _nw123 = new _module.C2();
        _nw123.__ctor(_832_v1, _895_v38, _898_v39, _832_v1);
        _899_v40 = _nw123;
        let _900_v41;
        _900_v41 = _dafny.Map.Empty.slice().updateUnsafe(_893_v37,_899_v40);
        _900_v41 = (_900_v41).update(_893_v37, _899_v40);
        _833_v2 = (_833_v2).plus(new BigNumber(752));
        let _901_v42;
        _901_v42 = _dafny.Map.Empty.slice().updateUnsafe(_895_v38,_833_v2);
        _901_v42 = (_901_v42).update(_895_v38, ((_832_v1) ? (_833_v2) : (_833_v2)));
      }
      let _rhs136 = !(_832_v1);
      let _rhs137 = (_dafny.ZERO).minus(_833_v2);
      r1 = _rhs136;
      _833_v2 = _rhs137;
      let _902_v43;
      _902_v43 = _dafny.Set.fromElements(_833_v2, _833_v2);
      let _903_v44;
      _903_v44 = _dafny.Seq.of((new BigNumber((_902_v43).length)).isLessThanOrEqualTo(new BigNumber(725)), (_833_v2).isLessThan(_833_v2));
      r1 = (_903_v44)[_module.__default.safeIndex(_833_v2, new BigNumber((_903_v44).length))];
      let _904_v45;
      let _init20 = ((_905_v2, _906_v44) => function (_907_i6) {
        return _module.__default.safeModuloInt(_907_i6, (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_905_v2,new BigNumber((_906_v44).length))).length)));
      })(_833_v2, _903_v44);
      let _nw124 = Array((new BigNumber(12)).toNumber());
      for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw124.length); _i0_20++) {
        _nw124[_i0_20] = _init20(new BigNumber(_i0_20));
      }
      _904_v45 = _nw124;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_904_v45).length))) {
        let _908_i5 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_908_i5)) && ((_908_i5).isLessThan(new BigNumber((_904_v45).length))))) {
          (_904_v45)[(_908_i5)] = _module.__default.safeModuloInt(_908_i5, _833_v2);
        }
      }
      (globalState).f8 = _832_v1;
      let _909_v46;
      _909_v46 = _dafny.Seq.UnicodeFromString("tgohvumtv");
      r0 = _dafny.Map.Empty.slice().updateUnsafe(((_832_v1) ? (_904_v45) : (_904_v45)),_dafny.Seq.Concat(_909_v46, _dafny.Seq.UnicodeFromString("kaeltdutj")));
      r1 = ((_832_v1) ? (_832_v1) : (_832_v1));
      return [r0, r1];
    }
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f9 = _dafny.Map.Empty;
      this._f10 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f10() {
      let _this = this;
      return _this._f10;
    };
    __ctor(f9, f10) {
      let _this = this;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      return;
    }
    fm0(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !((_this).f10) || ((_this).f10);
    };
    fm1(p0, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt(new BigNumber(-969), (new BigNumber(-383)).multipliedBy(new BigNumber(184)));
    };
    m0(globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = false;
      let r2 = _dafny.Seq.of();
      let r3 = _dafny.Seq.UnicodeFromString("");
      let _910_v0;
      _910_v0 = _dafny.Seq.UnicodeFromString("qlknnmx");
      let _911_v1;
      _911_v1 = _module.D1.create_DC3(_910_v0);
      let _912_v2;
      _912_v2 = _dafny.Seq.of((_this).f10);
      let _913_v3;
      _913_v3 = new BigNumber(30);
      let _pat_let_tv18 = _912_v2;
      let _pat_let_tv19 = globalState;
      let _pat_let_tv20 = _913_v3;
      let _pat_let_tv21 = globalState;
      let _914_v4;
      let _nw125 = Array((new BigNumber(19)).toNumber());
      _nw125[(_dafny.ZERO).toNumber()] = _911_v1;
      _nw125[(_dafny.ONE).toNumber()] = _911_v1;
      _nw125[(new BigNumber(2)).toNumber()] = _911_v1;
      _nw125[(new BigNumber(3)).toNumber()] = _911_v1;
      _nw125[(new BigNumber(4)).toNumber()] = _911_v1;
      _nw125[(new BigNumber(5)).toNumber()] = _911_v1;
      _nw125[(new BigNumber(6)).toNumber()] = _911_v1;
      _nw125[(new BigNumber(7)).toNumber()] = _911_v1;
      _nw125[(new BigNumber(8)).toNumber()] = _911_v1;
      _nw125[(new BigNumber(9)).toNumber()] = _911_v1;
      _nw125[(new BigNumber(10)).toNumber()] = function (_pat_let18_0) {
        return function (_915_dt__update__tmp_h0) {
          return function (_pat_let19_0) {
            return function (_918_dt__update_hcf5_h0) {
              return _module.D1.create_DC3(_918_dt__update_hcf5_h0);
            }(_pat_let19_0);
          }(_module.__default.fm22(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(588)), function (_916_i0) {
            return new _dafny.CodePoint('a'.codePointAt(0));
          })).length), (_this).f10, (_this).fm0(_pat_let_tv18, false, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-736)), function (_917_i1) {
            return new _dafny.CodePoint('g'.codePointAt(0));
          })).length), (_this).f10, _pat_let_tv19), _pat_let_tv20, _pat_let_tv21));
        }(_pat_let18_0);
      }(_911_v1);
      _nw125[(new BigNumber(11)).toNumber()] = _911_v1;
      _nw125[(new BigNumber(12)).toNumber()] = _module.D1.create_DC3(_910_v0);
      _nw125[(new BigNumber(13)).toNumber()] = _911_v1;
      _nw125[(new BigNumber(14)).toNumber()] = _911_v1;
      _nw125[(new BigNumber(15)).toNumber()] = _911_v1;
      _nw125[(new BigNumber(16)).toNumber()] = _911_v1;
      _nw125[(new BigNumber(17)).toNumber()] = _module.D1.create_DC3(_910_v0);
      _nw125[(new BigNumber(18)).toNumber()] = _911_v1;
      _914_v4 = _nw125;
      let _919_v5;
      _919_v5 = _dafny.Map.Empty.slice().updateUnsafe(_913_v3,_914_v4);
      _914_v4 = (((_this).f10) ? ((((_919_v5).contains(new BigNumber(67))) ? ((_919_v5).get(new BigNumber(67))) : (_914_v4))) : (_914_v4));
      _910_v0 = _910_v0;
      r1 = (_this).f10;
      (globalState).f8 = _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("bvcipvb"), _dafny.Seq.UnicodeFromString("ruf"));
      let _920_v6;
      let _nw126 = new _module.C0();
      _nw126.__ctor();
      _920_v6 = _nw126;
      _920_v6 = _920_v6;
      _913_v3 = _913_v3;
      let _921_v7;
      _921_v7 = _dafny.MultiSet.fromElements((_this).f10);
      let _922_v8;
      _922_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_921_v7).cardinality()),new BigNumber((_910_v0).length));
      r0 = _dafny.Seq.update(_910_v0, _module.__default.safeIndex(_913_v3, new BigNumber((_910_v0).length)), (_910_v0)[_module.__default.safeIndex(new BigNumber((_922_v8).length), new BigNumber((_910_v0).length))]);
      r1 = (_this).f10;
      r2 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(592)), function (_923_i2) {
        return _dafny.Set.fromElements((_this).f10, (_this).f10, !(true));
      });
      r3 = _910_v0;
      return [r0, r1, r2, r3];
    }
    m1(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = false;
      let _924_i0;
      _924_i0 = _dafny.ZERO;
      L2: {
        while (!(((((_this).f10) ? ((_this).f10) : ((_this).f10))) || ((((_this).f10) ? (!((_this).f10)) : ((_this).f10))))) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_924_i0)) {
              break L2;
            }
            _924_i0 = (_924_i0).plus(_dafny.ONE);
            let _925_v0;
            _925_v0 = _dafny.Seq.of((_this).f10);
            let _926_v1;
            _926_v1 = new BigNumber(169);
            let _927_v2;
            _927_v2 = _dafny.Map.Empty.slice().updateUnsafe(_926_v1,!((_this).f10));
            let _928_v3;
            _928_v3 = _dafny.Set.fromElements((_this).fm0(_925_v0, (_this).f10, new BigNumber((_927_v2).length), (_this).f10, globalState));
            let _929_v4;
            _929_v4 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f10) ? (false) : (!((_this).f10))),_928_v3);
            let _930_v5;
            let _nw127 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
            _930_v5 = _nw127;
            let _index146 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_930_v5).length));
            (_930_v5)[_index146] = _926_v1;
            let _index147 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_930_v5).length));
            let _rhs138 = (_929_v4).Merge(_929_v4);
            let _rhs139 = (_dafny.ZERO).minus(_926_v1);
            let _lhs104 = _930_v5;
            let _lhs105 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_930_v5).length));
            _929_v4 = _rhs138;
            _lhs104[_lhs105] = _rhs139;
            (globalState).f8 = (_this).f10;
            let _931_v6;
            _931_v6 = _dafny.MultiSet.fromElements(false);
            let _932_v7;
            _932_v7 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_926_v1)).cardinality()),_931_v6),_926_v1);
            let _933_v8;
            _933_v8 = _dafny.Map.Empty.slice().updateUnsafe((_930_v5)[_module.__default.safeIndex(new BigNumber(5), new BigNumber((_930_v5).length))],_931_v6);
            _932_v7 = (_932_v7).update((_933_v8).Merge(_933_v8), _926_v1);
            let _934_v10;
            _934_v10 = _dafny.Set.fromElements(p0);
            _926_v1 = ((_930_v5)[_module.__default.safeIndex(new BigNumber(5), new BigNumber((_930_v5).length))]).minus(((_this).fm1((_this).f9, globalState)).plus(new BigNumber(((function () {
              let _coll42 = new _dafny.Map();
              for (const _compr_42 of (_934_v10).Elements) {
                let _935_v9 = _compr_42;
                if ((_934_v10).contains(_935_v9)) {
                  _coll42.push([_935_v9,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(186)), ((_936_p0) => function (_937_i1) {
                    return _936_p0;
                  })(p0))).length)]);
                }
              }
              return _coll42;
            }()).update(p0, (_930_v5)[_module.__default.safeIndex(new BigNumber(5), new BigNumber((_930_v5).length))])).length)));
          }
        }
      }
      let _938_v11;
      _938_v11 = _dafny.Seq.of(_dafny.Set.fromElements((_this).f10, true));
      let _939_v12;
      _939_v12 = new BigNumber(-285);
      let _940_v13;
      _940_v13 = _dafny.Set.fromElements((_this).f10, (_this).f10);
      if (!((_938_v11)[_module.__default.safeIndex(_939_v12, new BigNumber((_938_v11).length))]).equals(_940_v13)) {
        (globalState).f8 = (_this).f10;
        r1 = (_this).f10;
        let _941_v14;
        _941_v14 = _module.D0.create_DC0((_this).f10);
        let _942_v15;
        _942_v15 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f10) ? (_941_v14) : (_941_v14)),!(!((_this).f10) || ((_this).f10)));
        _942_v15 = (_942_v15).update(_module.__default.fm23(globalState), (_this).f10);
        let _943_v16;
        _943_v16 = _dafny.Seq.UnicodeFromString("hlsgfllkl");
        if ((!(new BigNumber((_943_v16).length)).isEqualTo(_939_v12)) === ((_this).f10)) {
          let _944_v17;
          _944_v17 = _dafny.Map.Empty.slice().updateUnsafe(p0,_939_v12);
          _944_v17 = (_944_v17).update(p0, _939_v12);
          _939_v12 = _939_v12;
          let _945_v18;
          let _nw128 = Array((new BigNumber(29)).toNumber()).fill(false);
          _945_v18 = _nw128;
          _945_v18 = _945_v18;
          let _946_v19;
          let _nw129 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Map.Empty);
          _946_v19 = _nw129;
          let _947_v20;
          let _init21 = function (_948_i2) {
            return (_948_i2).plus(new BigNumber((_dafny.Seq.of((_this).f10)).length));
          };
          let _nw130 = Array((new BigNumber(29)).toNumber());
          for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw130.length); _i0_21++) {
            _nw130[_i0_21] = _init21(new BigNumber(_i0_21));
          }
          _947_v20 = _nw130;
          let _949_v21;
          _949_v21 = _module.D3.create_DC7(_947_v20);
          let _950_v22;
          _950_v22 = _dafny.Map.Empty.slice().updateUnsafe((_949_v21).dtor_cf11,(_this).fm1((_this).f9, globalState));
          let _index148 = _module.__default.safeIndex(new BigNumber(846), new BigNumber((_946_v19).length));
          (_946_v19)[_index148] = _950_v22;
          let _index149 = _module.__default.safeIndex(new BigNumber(846), new BigNumber((_946_v19).length));
          (_946_v19)[_index149] = (_dafny.Map.Empty.slice().updateUnsafe(_947_v20,(_this).fm1((_this).f9, globalState))).Merge(_950_v22);
          let _951_v23;
          _951_v23 = _dafny.Seq.of(new BigNumber(762));
          let _952_v24;
          _952_v24 = _dafny.Seq.of((_this).f10);
          _951_v23 = (((_this).fm0(_952_v24, (_this).f10, new BigNumber((_dafny.Seq.of(_939_v12, _939_v12)).length), !((_this).f10), globalState)) ? (_951_v23) : (_951_v23));
        } else {
          let _953_v25;
          let _nw131 = Array((new BigNumber(7)).toNumber()).fill(false);
          _953_v25 = _nw131;
          let _index150 = _module.__default.safeIndex(new BigNumber(728), new BigNumber((_953_v25).length));
          (_953_v25)[_index150] = !((_this).f10) || ((_this).f10);
          let _954_v26;
          _954_v26 = _dafny.Seq.of((_this).f10);
          let _955_v27;
          _955_v27 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_954_v26, _954_v26),new BigNumber((_dafny.Set.fromElements(_939_v12, _939_v12)).length));
          let _956_v28;
          _956_v28 = _dafny.Set.fromElements((_this).f9);
          let _index151 = _module.__default.safeIndex(new BigNumber(728), new BigNumber((_953_v25).length));
          let _rhs140 = _953_v25;
          let _rhs141 = (((_this).f10) ? ((_this).fm0(_954_v26, false, _939_v12, (_this).f10, globalState)) : ((_this).f10));
          let _rhs142 = _module.__default.safeModuloInt(_939_v12, _939_v12);
          let _rhs143 = (_module.__default.fm24((_this).f10, globalState)).IsDisjointFrom(_956_v28);
          let _rhs144 = _955_v27;
          let _lhs106 = _953_v25;
          let _lhs107 = _module.__default.safeIndex(new BigNumber(728), new BigNumber((_953_v25).length));
          let _lhs108 = globalState;
          _953_v25 = _rhs140;
          _lhs106[_lhs107] = _rhs141;
          _939_v12 = _rhs142;
          _lhs108.f8 = _rhs143;
          _955_v27 = _rhs144;
          let _957_v29;
          let _nw132 = new _module.C0();
          _nw132.__ctor();
          _957_v29 = _nw132;
          let _958_v30;
          _958_v30 = new _dafny.CodePoint('h'.codePointAt(0));
          let _959_v32;
          _959_v32 = _dafny.Map.Empty.slice().updateUnsafe((_953_v25)[_module.__default.safeIndex(new BigNumber(728), new BigNumber((_953_v25).length))],(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm25((_dafny.ZERO).minus(_939_v12), new BigNumber((_943_v16).length), globalState)).length),_939_v12)).Merge((_dafny.Map.Empty.slice().updateUnsafe(_939_v12,_939_v12)).update(new BigNumber((_943_v16).length), new BigNumber((function () {
            let _coll43 = new _dafny.Map();
            for (const _compr_43 of (_dafny.Map.Empty.slice().updateUnsafe(_939_v12,(_dafny.ZERO).minus(_939_v12))).Keys.Elements) {
              let _960_v31 = _compr_43;
              if ((_dafny.Map.Empty.slice().updateUnsafe(_939_v12,(_dafny.ZERO).minus(_939_v12))).contains(_960_v31)) {
                _coll43.push([_module.__default.safeModuloInt(_960_v31, _939_v12),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_939_v12,_939_v12)).length)]);
              }
            }
            return _coll43;
          }()).length))));
          let _961_v33;
          let _nw133 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
          _961_v33 = _nw133;
          let _962_v34;
          _962_v34 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(323),_939_v12);
          let _rhs145 = _958_v30;
          let _rhs146 = _dafny.Map.Empty.slice().updateUnsafe((_953_v25)[_module.__default.safeIndex(new BigNumber(728), new BigNumber((_953_v25).length))],(_962_v34).Merge(_962_v34));
          let _rhs147 = _961_v33;
          _958_v30 = _rhs145;
          _959_v32 = _rhs146;
          _961_v33 = _rhs147;
          let _963_v35;
          _963_v35 = _dafny.Map.Empty.slice().updateUnsafe(_958_v30,(_this).f10);
          let _964_v36;
          _964_v36 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm0(_954_v26, false, _939_v12, (_this).fm0(_954_v26, (_this).f10, _939_v12, (_this).f10, globalState), globalState),(_module.D1.create_DC4((_this).f10, (((_963_v35).contains(_958_v30)) ? ((_963_v35).get(_958_v30)) : ((_953_v25)[_module.__default.safeIndex(new BigNumber(728), new BigNumber((_953_v25).length))])), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("pojq")).length)))).dtor_cf7);
          _964_v36 = (_964_v36).update((_this).f10, true);
          let _965_v37;
          _965_v37 = _dafny.Seq.of(_943_v16);
          let _index152 = _module.__default.safeIndex(new BigNumber(728), new BigNumber((_953_v25).length));
          (_953_v25)[_index152] = _dafny.areEqual(_dafny.Seq.update(_965_v37, _module.__default.safeIndex(_939_v12, new BigNumber((_965_v37).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-464)), ((_966_v30) => function (_967_i3) {
            return _966_v30;
          })(_958_v30))), _965_v37);
        }
        let _968_v38;
        _968_v38 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_module.__default.safeModuloInt(_939_v12, new BigNumber(391))));
        _968_v38 = (_968_v38).Union(_968_v38);
      } else {
        let _969_v39;
        let _nw134 = new _module.C0();
        _nw134.__ctor();
        _969_v39 = _nw134;
        let _970_v40;
        _970_v40 = _dafny.Seq.of(_969_v39, _969_v39);
        let _971_v41;
        _971_v41 = _dafny.Seq.of((_this).f10, (_this).f10);
        let _rhs148 = (((_this).f10) ? (new BigNumber((_971_v41).length)) : (_939_v12));
        let _rhs149 = (_this).f10;
        let _rhs150 = (_this).f10;
        let _rhs151 = _dafny.Seq.Concat(_dafny.Seq.Concat(_970_v40, _970_v40), _970_v40);
        let _lhs109 = globalState;
        let _lhs110 = globalState;
        _939_v12 = _rhs148;
        _lhs109.f8 = _rhs149;
        _lhs110.f8 = _rhs150;
        _970_v40 = _rhs151;
        let _972_v42;
        let _nw135 = Array((new BigNumber(18)).toNumber()).fill(false);
        _972_v42 = _nw135;
        let _973_v43;
        _973_v43 = _dafny.Seq.UnicodeFromString("coavp");
        let _974_v44;
        _974_v44 = _dafny.MultiSet.fromElements(new BigNumber((_973_v43).length));
        let _index153 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_972_v42).length));
        (_972_v42)[_index153] = (_974_v44).equals(_974_v44);
        let _index154 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_972_v42).length));
        let _rhs152 = (_this).f10;
        let _rhs153 = (_this).f10;
        let _rhs154 = (_this).f10;
        let _lhs111 = globalState;
        let _lhs112 = _972_v42;
        let _lhs113 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_972_v42).length));
        let _lhs114 = globalState;
        _lhs111.f8 = _rhs152;
        _lhs112[_lhs113] = _rhs153;
        _lhs114.f8 = _rhs154;
        _939_v12 = _module.__default.safeDivisionInt((((_972_v42)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_972_v42).length))]) ? (_939_v12) : (_939_v12)), (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(323), _939_v12)));
        let _975_v45;
        let _nw136 = Array((new BigNumber(12)).toNumber());
        _nw136[(_dafny.ZERO).toNumber()] = p0;
        _nw136[(_dafny.ONE).toNumber()] = p0;
        _nw136[(new BigNumber(2)).toNumber()] = p0;
        _nw136[(new BigNumber(3)).toNumber()] = p0;
        _nw136[(new BigNumber(4)).toNumber()] = p0;
        _nw136[(new BigNumber(5)).toNumber()] = p0;
        _nw136[(new BigNumber(6)).toNumber()] = p0;
        _nw136[(new BigNumber(7)).toNumber()] = p0;
        _nw136[(new BigNumber(8)).toNumber()] = new _dafny.CodePoint('c'.codePointAt(0));
        _nw136[(new BigNumber(9)).toNumber()] = (((_972_v42)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_972_v42).length))]) ? (_module.__default.fm26(p0, globalState)) : (p0));
        _nw136[(new BigNumber(10)).toNumber()] = p0;
        _nw136[(new BigNumber(11)).toNumber()] = p0;
        _975_v45 = _nw136;
        _975_v45 = ((!((_this).f10)) ? (_975_v45) : (_975_v45));
        let _976_v46;
        _976_v46 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_dafny.Seq.UnicodeFromString("kxm"));
        let _977_v47;
        _977_v47 = _dafny.Seq.of((((_976_v46).contains((_972_v42)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_972_v42).length))])) ? ((_976_v46).get((_972_v42)[_module.__default.safeIndex(new BigNumber(765), new BigNumber((_972_v42).length))])) : (_973_v43)), _973_v43);
        (globalState).f8 = !_dafny.Seq.contains((_977_v47)[_module.__default.safeIndex(_939_v12, new BigNumber((_977_v47).length))], p0);
      }
      let _978_v48;
      _978_v48 = _dafny.Seq.of(_939_v12, (_939_v12).minus(_939_v12), _939_v12);
      _939_v12 = (_978_v48)[_module.__default.safeIndex(_939_v12, new BigNumber((_978_v48).length))];
      r1 = (_this).f10;
      let _979_v49;
      let _nw137 = new _module.C0();
      _nw137.__ctor();
      _979_v49 = _nw137;
      let _980_v50;
      _980_v50 = _dafny.Seq.UnicodeFromString("f");
      let _rhs155 = !(_dafny.areEqual(_980_v50, _980_v50));
      let _rhs156 = (_this).f10;
      r1 = _rhs155;
      r1 = _rhs156;
      let _981_v51;
      let _nw138 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
      _981_v51 = _nw138;
      let _982_v52;
      _982_v52 = _dafny.Map.Empty.slice().updateUnsafe(_981_v51,_dafny.Seq.UnicodeFromString("xxghgxtt"));
      r0 = (_982_v52).Merge(_dafny.Map.Empty.slice().updateUnsafe(_981_v51,_980_v50));
      r1 = (_this).f10;
      return [r0, r1];
    }
    m5(globalState) {
      let _this = this;
      let _983_v0;
      _983_v0 = _dafny.Seq.of((_this).f10, (_this).f10);
      (globalState).f8 = _dafny.Seq.contains(_983_v0, !((_this).f10));
      (globalState).f8 = (_this).f10;
      let _984_i0;
      _984_i0 = _dafny.ZERO;
      L3: {
        while ((_this).f10) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_984_i0)) {
              break L3;
            }
            _984_i0 = (_984_i0).plus(_dafny.ONE);
            if ((_this).f10) {
              let _985_v1;
              _985_v1 = new BigNumber(762);
              _985_v1 = ((_dafny.ZERO).minus(_985_v1)).minus(_985_v1);
              let _986_v2;
              _986_v2 = _dafny.Seq.UnicodeFromString("qqk");
              let _987_v3;
              let _nw139 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
              _987_v3 = _nw139;
              let _988_v4;
              _988_v4 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("nsdcjxlyo"),_987_v3);
              let _989_v5;
              let _990_v6;
              let _991_v7;
              let _992_v8;
              let _out48;
              let _out49;
              let _out50;
              let _out51;
              let _outcollector14 = (_this).m6(new BigNumber((_986_v2).length), (((_988_v4).contains(_986_v2)) ? ((_988_v4).get(_986_v2)) : (_987_v3)), (_985_v1).multipliedBy(_985_v1), globalState);
              _out48 = _outcollector14[0];
              _out49 = _outcollector14[1];
              _out50 = _outcollector14[2];
              _out51 = _outcollector14[3];
              _989_v5 = _out48;
              _990_v6 = _out49;
              _991_v7 = _out50;
              _992_v8 = _out51;
              let _993_v9;
              _993_v9 = new _dafny.CodePoint('f'.codePointAt(0));
              _993_v9 = _993_v9;
              _986_v2 = _dafny.Seq.of(_993_v9);
              let _994_v10;
              let _nw140 = new _module.C3();
              _nw140.__ctor();
              _994_v10 = _nw140;
              _994_v10 = _994_v10;
            } else {
              let _995_v11;
              _995_v11 = new BigNumber(153);
              _995_v11 = _995_v11;
              let _996_v12;
              let _init22 = function (_997_i1) {
                return !((_this).f10);
              };
              let _nw141 = Array((new BigNumber(7)).toNumber());
              for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw141.length); _i0_22++) {
                _nw141[_i0_22] = _init22(new BigNumber(_i0_22));
              }
              _996_v12 = _nw141;
              let _index155 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_996_v12).length));
              (_996_v12)[_index155] = (_this).f10;
              let _index156 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_996_v12).length));
              (_996_v12)[_index156] = !((_983_v0)[_module.__default.safeIndex(_995_v11, new BigNumber((_983_v0).length))]);
              let _998_v13;
              _998_v13 = _dafny.Seq.UnicodeFromString("xum");
              _998_v13 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-636)), ((_999_v12, _1000_v11) => function (_1001_i2) {
                return (_module.D4.create_DC12((_999_v12)[_module.__default.safeIndex(new BigNumber(684), new BigNumber((_999_v12).length))], new BigNumber(357), new _dafny.CodePoint('d'.codePointAt(0)), (_dafny.ZERO).minus(_1000_v11), _1000_v11)).dtor_cf21;
              })(_996_v12, _995_v11));
              let _1002_v14;
              let _nw142 = new _module.C1();
              _nw142.__ctor(_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(792)), function (_1003_i3) {
                return new _dafny.CodePoint('a'.codePointAt(0));
              }), _dafny.Seq.UnicodeFromString("sixbvx")), _998_v13);
              _1002_v14 = _nw142;
              let _1004_v15;
              _1004_v15 = _dafny.Seq.of(_995_v11, _995_v11, _995_v11, new BigNumber(369), new BigNumber((_dafny.Set.fromElements(_995_v11)).length));
              let _1005_v16;
              _1005_v16 = _dafny.Seq.of(_1004_v15);
              let _1006_v17;
              _1006_v17 = _dafny.Seq.of((_1005_v16)[_module.__default.safeIndex(_995_v11, new BigNumber((_1005_v16).length))]);
              let _1007_v18;
              _1007_v18 = _dafny.Seq.of((new BigNumber(((_1005_v16)[_module.__default.safeIndex(new BigNumber(((_1002_v14).f13).length), new BigNumber((_1005_v16).length))]).length)).minus(new BigNumber(650)));
              _1004_v15 = _1007_v18;
            }
            let _1008_v19;
            _1008_v19 = new BigNumber(745);
            let _1009_v20;
            _1009_v20 = _dafny.Set.fromElements(true);
            _1008_v19 = (((_1009_v20).IsProperSubsetOf(_1009_v20)) ? (_module.__default.safeDivisionInt(_1008_v19, new BigNumber(426))) : (_1008_v19));
            (globalState).f8 = (_this).f10;
            let _1010_v21;
            let _nw143 = Array((new BigNumber(25)).toNumber()).fill(false);
            _1010_v21 = _nw143;
            let _index157 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_1010_v21).length));
            (_1010_v21)[_index157] = (_this).f10;
            let _1011_v22;
            _1011_v22 = _dafny.Seq.of(_1008_v19);
            let _1012_v23;
            _1012_v23 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_dafny.ZERO).minus((_this).fm1((_this).f9, globalState)));
            let _1013_v24;
            _1013_v24 = _dafny.Set.fromElements(_1012_v23);
            let _index158 = _module.__default.safeIndex(new BigNumber(103), new BigNumber((_1010_v21).length));
            (_1010_v21)[_index158] = _dafny.Seq.IsPrefixOf(_1011_v22, _dafny.Seq.of((((_1012_v23).contains((_this).f10)) ? ((_1012_v23).get((_this).f10)) : (new BigNumber(-128))), new BigNumber(411), new BigNumber((_1013_v24).length), _1008_v19));
          }
        }
      }
      let _1014_i4;
      _1014_i4 = _dafny.ZERO;
      L4: {
        while ((_this).f10) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1014_i4)) {
              break L4;
            }
            _1014_i4 = (_1014_i4).plus(_dafny.ONE);
            if (((_this).f10) === (!((_this).f10))) {
              let _1015_v25;
              _1015_v25 = _dafny.Seq.UnicodeFromString("yuctjfpul");
              let _1016_v26;
              let _nw144 = new _module.C1();
              _nw144.__ctor((_this).f10, _1015_v25);
              _1016_v26 = _nw144;
              let _1017_v27;
              _1017_v27 = new BigNumber(729);
              let _1018_v28;
              _1018_v28 = _dafny.Map.Empty.slice().updateUnsafe(_1017_v27,(_this).f9);
              let _1019_v29;
              _1019_v29 = _dafny.Map.Empty.slice().updateUnsafe((_1018_v28).Merge(_1018_v28),(_1017_v27).plus(new BigNumber(((_1016_v26).f13).length)));
              let _1020_v30;
              _1020_v30 = _dafny.Map.Empty.slice().updateUnsafe((_1016_v26).f12,_1017_v27);
              _1019_v29 = (_1019_v29).update(_1018_v28, _module.__default.safeModuloInt(_1017_v27, (((_1020_v30).contains((_this).f10)) ? ((_1020_v30).get((_this).f10)) : (_1017_v27))));
              let _1021_v31;
              _1021_v31 = _dafny.Seq.of((_1016_v26).f13, (_1016_v26).f13, (_1016_v26).f13);
              _1015_v25 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("w"), _dafny.Seq.Concat((_1021_v31)[_module.__default.safeIndex(_1017_v27, new BigNumber((_1021_v31).length))], _1015_v25));
              _1017_v27 = new BigNumber((((_this).f9).Merge((_this).f9)).length);
              let _1022_v32;
              _1022_v32 = _dafny.MultiSet.fromElements(_1017_v27);
              _1020_v30 = (_1020_v30).update((_1022_v32).IsProperSubsetOf(_1022_v32), _1017_v27);
            } else {
              let _1023_v33;
              _1023_v33 = _dafny.Seq.UnicodeFromString("cckevci");
              let _1024_v34;
              _1024_v34 = new BigNumber(419);
              let _1025_v35;
              _1025_v35 = _dafny.Seq.of(_1024_v34);
              let _1026_v36;
              let _init23 = function (_1027_i5) {
                return new _dafny.CodePoint('n'.codePointAt(0));
              };
              let _nw145 = Array((new BigNumber(29)).toNumber());
              for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw145.length); _i0_23++) {
                _nw145[_i0_23] = _init23(new BigNumber(_i0_23));
              }
              _1026_v36 = _nw145;
              let _1028_v37;
              _1028_v37 = _dafny.MultiSet.fromElements(_1024_v34);
              let _rhs157 = _1023_v33;
              let _rhs158 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1025_v35, _dafny.Seq.of(_1024_v34, _1024_v34)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(750)), ((_1029_v34) => function (_1030_i6) {
                return _1029_v34;
              })(_1024_v34)));
              let _rhs159 = _1026_v36;
              let _rhs160 = _module.__default.safeDivisionInt(_1024_v34, _1024_v34);
              let _rhs161 = _module.__default.safeDivisionInt(_1024_v34, new BigNumber(((_1028_v37).Intersect(_1028_v37)).cardinality()));
              _1023_v33 = _rhs157;
              _1025_v35 = _rhs158;
              _1026_v36 = _rhs159;
              _1024_v34 = _rhs160;
              _1024_v34 = _rhs161;
              (globalState).f8 = (_this).f10;
              let _1031_v38;
              _1031_v38 = _dafny.MultiSet.fromElements(!((_this).f10));
              let _1032_v39;
              _1032_v39 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_983_v0), _1031_v38);
              let _1033_v40;
              _1033_v40 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,new _dafny.CodePoint('g'.codePointAt(0)));
              let _1034_v41;
              _1034_v41 = _dafny.MultiSet.fromElements(_module.__default.fm22(new BigNumber((_dafny.Seq.of(_1024_v34)).length), (_this).f10, (_this).fm0(_983_v0, (_this).f10, (_this).fm1((_this).f9, globalState), (_this).f10, globalState), _1024_v34, globalState));
              let _1035_v42;
              let _nw146 = Array((new BigNumber(5)).toNumber());
              _nw146[(_dafny.ZERO).toNumber()] = (_this).f10;
              _nw146[(_dafny.ONE).toNumber()] = (_this).f10;
              _nw146[(new BigNumber(2)).toNumber()] = (_this).f10;
              _nw146[(new BigNumber(3)).toNumber()] = !(false);
              _nw146[(new BigNumber(4)).toNumber()] = (_this).f10;
              _1035_v42 = _nw146;
              let _1036_v43;
              _1036_v43 = _dafny.Seq.of(_1035_v42);
              let _1037_v44;
              let _nw147 = Array((new BigNumber(25)).toNumber());
              _nw147[(_dafny.ZERO).toNumber()] = _1024_v34;
              _nw147[(_dafny.ONE).toNumber()] = (_1024_v34).multipliedBy(new BigNumber((_1023_v33).length));
              _nw147[(new BigNumber(2)).toNumber()] = _1024_v34;
              _nw147[(new BigNumber(3)).toNumber()] = _1024_v34;
              _nw147[(new BigNumber(4)).toNumber()] = new BigNumber(((_1032_v39)[_module.__default.safeIndex(new BigNumber(531), new BigNumber((_1032_v39).length))]).cardinality());
              _nw147[(new BigNumber(5)).toNumber()] = _1024_v34;
              _nw147[(new BigNumber(6)).toNumber()] = _1024_v34;
              _nw147[(new BigNumber(7)).toNumber()] = _1024_v34;
              _nw147[(new BigNumber(8)).toNumber()] = (_1024_v34).minus(new BigNumber(915));
              _nw147[(new BigNumber(9)).toNumber()] = _1024_v34;
              _nw147[(new BigNumber(10)).toNumber()] = _1024_v34;
              _nw147[(new BigNumber(11)).toNumber()] = (((_this).f10) ? (new BigNumber((_1033_v40).length)) : (_1024_v34));
              _nw147[(new BigNumber(12)).toNumber()] = _module.__default.safeModuloInt(_1024_v34, _1024_v34);
              _nw147[(new BigNumber(13)).toNumber()] = (_dafny.ZERO).minus(_1024_v34);
              _nw147[(new BigNumber(14)).toNumber()] = (((_this).f10) ? ((_dafny.ZERO).minus((_dafny.ZERO).minus(_1024_v34))) : (_1024_v34));
              _nw147[(new BigNumber(15)).toNumber()] = new BigNumber((_1034_v41).cardinality());
              _nw147[(new BigNumber(16)).toNumber()] = (_1024_v34).multipliedBy(new BigNumber(-334));
              _nw147[(new BigNumber(17)).toNumber()] = _1024_v34;
              _nw147[(new BigNumber(18)).toNumber()] = _1024_v34;
              _nw147[(new BigNumber(19)).toNumber()] = (((_1028_v37).contains((_dafny.ZERO).minus(_1024_v34))) ? ((_1028_v37).get((_dafny.ZERO).minus(_1024_v34))) : (_1024_v34));
              _nw147[(new BigNumber(20)).toNumber()] = (_1024_v34).multipliedBy(_1024_v34);
              _nw147[(new BigNumber(21)).toNumber()] = (_dafny.ZERO).minus(_1024_v34);
              _nw147[(new BigNumber(22)).toNumber()] = new BigNumber(((((_this).f10) ? (_1036_v43) : (_1036_v43))).length);
              _nw147[(new BigNumber(23)).toNumber()] = _1024_v34;
              _nw147[(new BigNumber(24)).toNumber()] = new BigNumber(236);
              _1037_v44 = _nw147;
              let _index159 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_1037_v44).length));
              (_1037_v44)[_index159] = _module.__default.safeModuloInt(_1024_v34, _1024_v34);
              let _1038_v45;
              _1038_v45 = new _dafny.CodePoint('u'.codePointAt(0));
              let _index160 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_1037_v44).length));
              (_1037_v44)[_index160] = ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_1023_v33, _module.__default.safeIndex(new BigNumber(284), new BigNumber((_1023_v33).length)), _1038_v45)).length))).minus((_dafny.ZERO).minus(_1024_v34));
              let _index161 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_1035_v42).length));
              (_1035_v42)[_index161] = (_this).f10;
              let _index162 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_1035_v42).length));
              let _rhs162 = (_module.__default.fm46((_this).f10, (_983_v0)[_module.__default.safeIndex(_1024_v34, new BigNumber((_983_v0).length))], !((_this).f10), globalState)).IsSubsetOf(_1028_v37);
              let _rhs163 = (_this).f10;
              let _lhs115 = _1035_v42;
              let _lhs116 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_1035_v42).length));
              let _lhs117 = globalState;
              _lhs115[_lhs116] = _rhs162;
              _lhs117.f8 = _rhs163;
              let _1039_v46;
              _1039_v46 = _dafny.Map.Empty.slice().updateUnsafe(_1024_v34,_1024_v34);
              _1039_v46 = (_1039_v46).update(new BigNumber(154), (_1037_v44)[_module.__default.safeIndex(new BigNumber(163), new BigNumber((_1037_v44).length))]);
            }
            let _1040_v47;
            let _nw148 = Array((new BigNumber(16)).toNumber()).fill(_dafny.ZERO);
            _1040_v47 = _nw148;
            let _1041_v48;
            _1041_v48 = new BigNumber(915);
            let _index163 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_1040_v47).length));
            (_1040_v47)[_index163] = _1041_v48;
            let _index164 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_1040_v47).length));
            (_1040_v47)[_index164] = _module.__default.safeDivisionInt((_1041_v48).minus(_1041_v48), _1041_v48);
            let _1042_v49;
            let _nw149 = Array((new BigNumber(6)).toNumber());
            _1042_v49 = _nw149;
            let _1043_v50;
            let _nw150 = Array((new BigNumber(3)).toNumber()).fill(false);
            _1043_v50 = _nw150;
            let _1044_v51;
            let _nw151 = new _module.C2();
            _nw151.__ctor(!((_this).f10), _1043_v50, (_this).f9, (_this).f10);
            _1044_v51 = _nw151;
            let _index165 = _module.__default.safeIndex(new BigNumber(479), new BigNumber((_1042_v49).length));
            (_1042_v49)[_index165] = _1044_v51;
            let _1045_v52;
            _1045_v52 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-785)), function (_1046_i7) {
              return new BigNumber((_dafny.Seq.UnicodeFromString("krxj")).length);
            }),_1044_v51);
            let _1047_v54;
            _1047_v54 = _dafny.Set.fromElements((_1044_v51).f10);
            let _1048_v55;
            _1048_v55 = _dafny.Seq.of(new BigNumber((function () {
              let _coll44 = new _dafny.Set();
              for (const _compr_44 of _dafny.IntegerRange(new BigNumber(525), new BigNumber(392))) {
                let _1049_v53 = _compr_44;
                if (((new BigNumber(525)).isLessThanOrEqualTo(_1049_v53)) && ((_1049_v53).isLessThan(new BigNumber(392)))) {
                  _coll44.add((_1049_v53).multipliedBy(_1041_v48));
                }
              }
              return _coll44;
            }()).length), new BigNumber((_1047_v54).length));
            let _1050_v56;
            _1050_v56 = _dafny.Seq.of(_1048_v55);
            let _index166 = _module.__default.safeIndex(new BigNumber(479), new BigNumber((_1042_v49).length));
            (_1042_v49)[_index166] = (((_1045_v52).contains((_1050_v56)[_module.__default.safeIndex((_1040_v47)[_module.__default.safeIndex(new BigNumber(505), new BigNumber((_1040_v47).length))], new BigNumber((_1050_v56).length))])) ? ((_1045_v52).get((_1050_v56)[_module.__default.safeIndex((_1040_v47)[_module.__default.safeIndex(new BigNumber(505), new BigNumber((_1040_v47).length))], new BigNumber((_1050_v56).length))])) : (_1044_v51));
            let _1051_v57;
            let _nw152 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.of());
            _1051_v57 = _nw152;
            let _1052_v58;
            let _nw153 = new _module.C3();
            _nw153.__ctor();
            _1052_v58 = _nw153;
            let _1053_v59;
            _1053_v59 = _dafny.Seq.of(_1052_v58);
            let _index167 = _module.__default.safeIndex(new BigNumber(984), new BigNumber((_1051_v57).length));
            (_1051_v57)[_index167] = _1053_v59;
            let _index168 = _module.__default.safeIndex(new BigNumber(984), new BigNumber((_1051_v57).length));
            (_1051_v57)[_index168] = _1053_v59;
          }
        }
      }
      if ((_this).f10) {
        let _1054_v60;
        _1054_v60 = _dafny.Seq.UnicodeFromString("tndwyncs");
        let _1055_v61;
        _1055_v61 = _dafny.MultiSet.fromElements(_1054_v60);
        let _1056_v62;
        _1056_v62 = new BigNumber(469);
        let _1057_v63;
        _1057_v63 = _module.D11.create_DC27((_1055_v61).update(_1054_v60, _module.__default.abs(_1056_v62)));
        if ((_dafny.MultiSet.fromElements(_1054_v60)).IsSubsetOf((_1057_v63).dtor_cf47)) {
          _1056_v62 = _1056_v62;
          let _1058_v64;
          let _nw154 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
          _1058_v64 = _nw154;
          let _1059_v65;
          _1059_v65 = _dafny.Set.fromElements((_this).f10, false);
          let _1060_v66;
          _1060_v66 = _module.D8.create_DC24(!((_this).f10), _dafny.Set.fromElements(_1058_v64), new BigNumber((_1059_v65).length));
          let _pat_let_tv22 = globalState;
          let _1061_v67;
          let _nw155 = Array((new BigNumber(8)).toNumber());
          _nw155[(_dafny.ZERO).toNumber()] = _1060_v66;
          _nw155[(_dafny.ONE).toNumber()] = _1060_v66;
          _nw155[(new BigNumber(2)).toNumber()] = _1060_v66;
          _nw155[(new BigNumber(3)).toNumber()] = _1060_v66;
          _nw155[(new BigNumber(4)).toNumber()] = _1060_v66;
          _nw155[(new BigNumber(5)).toNumber()] = _1060_v66;
          _nw155[(new BigNumber(6)).toNumber()] = _1060_v66;
          _nw155[(new BigNumber(7)).toNumber()] = function (_pat_let20_0) {
            return function (_1062_dt__update__tmp_h0) {
              return function (_pat_let21_0) {
                return function (_1063_dt__update_hcf44_h0) {
                  return _module.D8.create_DC24((_1062_dt__update__tmp_h0).dtor_cf42, (_1062_dt__update__tmp_h0).dtor_cf43, _1063_dt__update_hcf44_h0);
                }(_pat_let21_0);
              }((_this).fm1((_this).f9, _pat_let_tv22));
            }(_pat_let20_0);
          }(_1060_v66);
          _1061_v67 = _nw155;
          let _1064_v68;
          _1064_v68 = _dafny.Seq.of(_1061_v67, _1061_v67);
          let _1065_v69;
          _1065_v69 = _module.D12.create_DC30(_1061_v67);
          let _1066_v70;
          let _nw156 = Array((new BigNumber(26)).toNumber());
          _nw156[(_dafny.ZERO).toNumber()] = _1061_v67;
          _nw156[(_dafny.ONE).toNumber()] = _1061_v67;
          _nw156[(new BigNumber(2)).toNumber()] = _1061_v67;
          _nw156[(new BigNumber(3)).toNumber()] = _1061_v67;
          _nw156[(new BigNumber(4)).toNumber()] = _1061_v67;
          _nw156[(new BigNumber(5)).toNumber()] = _1061_v67;
          _nw156[(new BigNumber(6)).toNumber()] = _1061_v67;
          _nw156[(new BigNumber(7)).toNumber()] = _1061_v67;
          _nw156[(new BigNumber(8)).toNumber()] = _1061_v67;
          _nw156[(new BigNumber(9)).toNumber()] = (_1064_v68)[_module.__default.safeIndex(_1056_v62, new BigNumber((_1064_v68).length))];
          _nw156[(new BigNumber(10)).toNumber()] = (((_this).f10) ? (_1061_v67) : (_1061_v67));
          _nw156[(new BigNumber(11)).toNumber()] = ((false) ? (_1061_v67) : (_1061_v67));
          _nw156[(new BigNumber(12)).toNumber()] = _1061_v67;
          _nw156[(new BigNumber(13)).toNumber()] = (_1065_v69).dtor_cf54;
          _nw156[(new BigNumber(14)).toNumber()] = (((_this).f10) ? (_1061_v67) : (_1061_v67));
          _nw156[(new BigNumber(15)).toNumber()] = _1061_v67;
          _nw156[(new BigNumber(16)).toNumber()] = _1061_v67;
          _nw156[(new BigNumber(17)).toNumber()] = _1061_v67;
          _nw156[(new BigNumber(18)).toNumber()] = _1061_v67;
          _nw156[(new BigNumber(19)).toNumber()] = _1061_v67;
          _nw156[(new BigNumber(20)).toNumber()] = _1061_v67;
          _nw156[(new BigNumber(21)).toNumber()] = _1061_v67;
          _nw156[(new BigNumber(22)).toNumber()] = (((_this).f10) ? (_1061_v67) : (_1061_v67));
          _nw156[(new BigNumber(23)).toNumber()] = _1061_v67;
          _nw156[(new BigNumber(24)).toNumber()] = _1061_v67;
          _nw156[(new BigNumber(25)).toNumber()] = _1061_v67;
          _1066_v70 = _nw156;
          _1066_v70 = _1066_v70;
          (globalState).f8 = (_this).f10;
          let _1067_v71;
          let _nw157 = Array((new BigNumber(3)).toNumber()).fill(false);
          _1067_v71 = _nw157;
          let _index169 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_1067_v71).length));
          (_1067_v71)[_index169] = (_1056_v62).isLessThanOrEqualTo(_1056_v62);
          let _1068_v72;
          _1068_v72 = _dafny.Map.Empty.slice().updateUnsafe(_1054_v60,(_this).f10);
          let _1069_v73;
          _1069_v73 = _module.D1.create_DC4((_this).f10, (_this).f10, _1056_v62);
          let _1070_v74;
          _1070_v74 = _dafny.Map.Empty.slice().updateUnsafe(_1056_v62,(_1069_v73).dtor_cf8);
          let _index170 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_1067_v71).length));
          let _rhs164 = (((_1068_v72).contains(_1054_v60)) ? ((_1068_v72).get(_1054_v60)) : ((_this).f10));
          let _rhs165 = !((_this).f10);
          let _rhs166 = (_this).f10;
          let _rhs167 = !((_1056_v62).isLessThanOrEqualTo(new BigNumber(((_1070_v74).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1056_v62,_1056_v62))).length)));
          let _lhs118 = globalState;
          let _lhs119 = globalState;
          let _lhs120 = _1067_v71;
          let _lhs121 = _module.__default.safeIndex(new BigNumber(722), new BigNumber((_1067_v71).length));
          let _lhs122 = globalState;
          _lhs118.f8 = _rhs164;
          _lhs119.f8 = _rhs165;
          _lhs120[_lhs121] = _rhs166;
          _lhs122.f8 = _rhs167;
          let _1071_v75;
          _1071_v75 = _dafny.MultiSet.fromElements((_1067_v71)[_module.__default.safeIndex(new BigNumber(722), new BigNumber((_1067_v71).length))], (_this).f10);
          let _index171 = _module.__default.safeIndex(new BigNumber(14), new BigNumber((_1058_v64).length));
          (_1058_v64)[_index171] = (((_1071_v75).contains((_this).fm0(_983_v0, (_this).f10, _1056_v62, (_this).f10, globalState))) ? ((_1071_v75).get((_this).fm0(_983_v0, (_this).f10, _1056_v62, (_this).f10, globalState))) : (new BigNumber(819)));
          let _1072_v76;
          _1072_v76 = new _dafny.CodePoint('a'.codePointAt(0));
          let _1073_v77;
          _1073_v77 = _dafny.Map.Empty.slice().updateUnsafe(_1058_v64,_1056_v62);
          let _index172 = _module.__default.safeIndex(new BigNumber(14), new BigNumber((_1058_v64).length));
          (_1058_v64)[_index172] = _module.__default.safeModuloInt((_module.D6.create_DC18(_1072_v76, new BigNumber((_1054_v60).length), new BigNumber(142))).dtor_cf36, new BigNumber(((_1073_v77).Merge(_1073_v77)).length));
        } else {
          let _1074_v78;
          let _init24 = function (_1075_i8) {
            return (_this).f10;
          };
          let _nw158 = Array((new BigNumber(27)).toNumber());
          for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw158.length); _i0_24++) {
            _nw158[_i0_24] = _init24(new BigNumber(_i0_24));
          }
          _1074_v78 = _nw158;
          let _1076_v79;
          let _nw159 = new _module.C2();
          _nw159.__ctor((_this).f10, (_module.D5.create_DC13(_1074_v78)).dtor_cf24, (_this).f9, (_this).f10);
          _1076_v79 = _nw159;
          let _nw160 = new _module.C2();
          _nw160.__ctor((_this).fm0(_983_v0, (_1076_v79).f10, _1056_v62, true, globalState), _1074_v78, (_1076_v79).f9, (_1076_v79).f10);
          _1076_v79 = _nw160;
          let _1077_v80;
          let _nw161 = Array((new BigNumber(7)).toNumber());
          _nw161[(_dafny.ZERO).toNumber()] = _983_v0;
          _nw161[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_983_v0, _dafny.Seq.of((_this).f10, (_1076_v79).f10));
          _nw161[(new BigNumber(2)).toNumber()] = _983_v0;
          _nw161[(new BigNumber(3)).toNumber()] = _dafny.Seq.update(_983_v0, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber(-31)), new BigNumber((_983_v0).length)), !((_1076_v79).f10));
          _nw161[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_983_v0, _983_v0);
          _nw161[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_983_v0, _983_v0);
          _nw161[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_983_v0, _983_v0);
          _1077_v80 = _nw161;
          let _index173 = _module.__default.safeIndex(new BigNumber(812), new BigNumber((_1077_v80).length));
          (_1077_v80)[_index173] = _983_v0;
          let _index174 = _module.__default.safeIndex(new BigNumber(812), new BigNumber((_1077_v80).length));
          let _rhs168 = _983_v0;
          let _rhs169 = _1054_v60;
          let _rhs170 = (((_this).f10) ? (false) : ((_1076_v79).f10));
          let _lhs123 = _1077_v80;
          let _lhs124 = _module.__default.safeIndex(new BigNumber(812), new BigNumber((_1077_v80).length));
          let _lhs125 = globalState;
          _lhs123[_lhs124] = _rhs168;
          _1054_v60 = _rhs169;
          _lhs125.f8 = _rhs170;
          let _1078_v81;
          _1078_v81 = _dafny.Seq.of((_1077_v80)[_module.__default.safeIndex(new BigNumber(812), new BigNumber((_1077_v80).length))]);
          let _1079_v82;
          let _nw162 = new _module.C1();
          _nw162.__ctor(!((_this).f10), _1054_v60);
          _1079_v82 = _nw162;
          let _1080_v83;
          _1080_v83 = new _dafny.CodePoint('b'.codePointAt(0));
          let _rhs171 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_983_v0, (_1078_v81)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_1079_v82, _1079_v82)).length), new BigNumber((_1078_v81).length))]), _dafny.Seq.Concat((_1077_v80)[_module.__default.safeIndex(new BigNumber(812), new BigNumber((_1077_v80).length))], _module.__default.fm38((_1079_v82).f12, globalState)));
          let _rhs172 = !((_this).fm0((_1077_v80)[_module.__default.safeIndex(new BigNumber(812), new BigNumber((_1077_v80).length))], (_1076_v79).f10, _1056_v62, (_1076_v79).f10, globalState));
          let _rhs173 = (_1079_v82).fm32(_1056_v62, _1080_v83, globalState);
          let _lhs126 = globalState;
          let _lhs127 = globalState;
          let _lhs128 = globalState;
          _lhs126.f8 = _rhs171;
          _lhs127.f8 = _rhs172;
          _lhs128.f8 = _rhs173;
          let _1081_v84;
          let _nw163 = Array((new BigNumber(13)).toNumber());
          _nw163[(_dafny.ZERO).toNumber()] = _1080_v83;
          _nw163[(_dafny.ONE).toNumber()] = _1080_v83;
          _nw163[(new BigNumber(2)).toNumber()] = _1080_v83;
          _nw163[(new BigNumber(3)).toNumber()] = _1080_v83;
          _nw163[(new BigNumber(4)).toNumber()] = _1080_v83;
          _nw163[(new BigNumber(5)).toNumber()] = _1080_v83;
          _nw163[(new BigNumber(6)).toNumber()] = _1080_v83;
          _nw163[(new BigNumber(7)).toNumber()] = _1080_v83;
          _nw163[(new BigNumber(8)).toNumber()] = _1080_v83;
          _nw163[(new BigNumber(9)).toNumber()] = _1080_v83;
          _nw163[(new BigNumber(10)).toNumber()] = _1080_v83;
          _nw163[(new BigNumber(11)).toNumber()] = _1080_v83;
          _nw163[(new BigNumber(12)).toNumber()] = _1080_v83;
          _1081_v84 = _nw163;
          let _1082_v85;
          let _nw164 = Array((new BigNumber(5)).toNumber());
          _nw164[(_dafny.ZERO).toNumber()] = _1081_v84;
          _nw164[(_dafny.ONE).toNumber()] = _1081_v84;
          _nw164[(new BigNumber(2)).toNumber()] = _1081_v84;
          _nw164[(new BigNumber(3)).toNumber()] = _1081_v84;
          _nw164[(new BigNumber(4)).toNumber()] = _1081_v84;
          _1082_v85 = _nw164;
          let _1083_v86;
          _1083_v86 = _dafny.Map.Empty.slice().updateUnsafe(_1082_v85,(((_1079_v82).f12) ? ((_1076_v79).f10) : ((_1079_v82).f12)));
          let _1084_v87;
          _1084_v87 = _dafny.MultiSet.fromElements(_1080_v83);
          _1083_v86 = (_1083_v86).update(_1082_v85, (_1084_v87).IsSubsetOf(_1084_v87));
          let _1085_v88;
          _1085_v88 = _module.D5.create_DC13(_1074_v78);
          let _1086_v89;
          _1086_v89 = _dafny.Map.Empty.slice().updateUnsafe((_1079_v82).f12,(_1085_v88).dtor_cf24);
          let _1087_v90;
          _1087_v90 = _1086_v89;
          _1087_v90 = _1087_v90;
        }
        let _1088_v91;
        let _nw165 = new _module.C3();
        _nw165.__ctor();
        _1088_v91 = _nw165;
        let _1089_v92;
        _1089_v92 = _dafny.MultiSet.fromElements(_1088_v91, _1088_v91);
        let _1090_v93;
        _1090_v93 = _dafny.MultiSet.fromElements((_this).f10, (_this).f10, (_this).f10);
        let _1091_v94;
        _1091_v94 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_1090_v93);
        let _rhs174 = _1089_v92;
        let _rhs175 = _1091_v94;
        _1089_v92 = _rhs174;
        _1091_v94 = _rhs175;
        (globalState).f8 = (_this).f10;
        let _1092_v95;
        let _1093_v96;
        let _1094_v97;
        let _1095_v98;
        let _out52;
        let _out53;
        let _out54;
        let _out55;
        let _outcollector15 = (_1088_v91).m0(globalState);
        _out52 = _outcollector15[0];
        _out53 = _outcollector15[1];
        _out54 = _outcollector15[2];
        _out55 = _outcollector15[3];
        _1092_v95 = _out52;
        _1093_v96 = _out53;
        _1094_v97 = _out54;
        _1095_v98 = _out55;
        _1056_v62 = new BigNumber((_dafny.Seq.Concat(_module.__default.fm45(_1093_v96, _1056_v62, _1093_v96, globalState), _dafny.Seq.Concat(_1054_v60, _1095_v98))).length);
      } else {
        let _1096_v99;
        _1096_v99 = new BigNumber(295);
        (globalState).f8 = ((_this).fm1(_dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_this).f10), globalState)).isLessThanOrEqualTo((((_this).f10) ? (_1096_v99) : ((_dafny.ZERO).minus(_1096_v99))));
        let _1097_v100;
        let _nw166 = new _module.C3();
        _nw166.__ctor();
        _1097_v100 = _nw166;
        let _1098_v101;
        _1098_v101 = _dafny.Seq.of(_1097_v100);
        let _1099_v102;
        _1099_v102 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_1097_v100);
        let _1100_v103;
        let _nw167 = Array((new BigNumber(15)).toNumber());
        _nw167[(_dafny.ZERO).toNumber()] = _1097_v100;
        _nw167[(_dafny.ONE).toNumber()] = _1097_v100;
        _nw167[(new BigNumber(2)).toNumber()] = (_1098_v101)[_module.__default.safeIndex(_1096_v99, new BigNumber((_1098_v101).length))];
        _nw167[(new BigNumber(3)).toNumber()] = _1097_v100;
        _nw167[(new BigNumber(4)).toNumber()] = _1097_v100;
        _nw167[(new BigNumber(5)).toNumber()] = _1097_v100;
        _nw167[(new BigNumber(6)).toNumber()] = _1097_v100;
        _nw167[(new BigNumber(7)).toNumber()] = (((_this).f10) ? (_1097_v100) : (_1097_v100));
        _nw167[(new BigNumber(8)).toNumber()] = _1097_v100;
        _nw167[(new BigNumber(9)).toNumber()] = _1097_v100;
        _nw167[(new BigNumber(10)).toNumber()] = (((_1099_v102).contains((_this).f10)) ? ((_1099_v102).get((_this).f10)) : (_1097_v100));
        _nw167[(new BigNumber(11)).toNumber()] = _1097_v100;
        _nw167[(new BigNumber(12)).toNumber()] = _1097_v100;
        _nw167[(new BigNumber(13)).toNumber()] = _1097_v100;
        _nw167[(new BigNumber(14)).toNumber()] = _1097_v100;
        _1100_v103 = _nw167;
        let _index175 = _module.__default.safeIndex(new BigNumber(975), new BigNumber((_1100_v103).length));
        (_1100_v103)[_index175] = _1097_v100;
        let _index176 = _module.__default.safeIndex(new BigNumber(975), new BigNumber((_1100_v103).length));
        (_1100_v103)[_index176] = _1097_v100;
        let _1101_v104;
        _1101_v104 = _dafny.Seq.UnicodeFromString("smnhcyhf");
        (globalState).f8 = _dafny.Seq.IsPrefixOf(_1101_v104, _1101_v104);
        let _1102_v105;
        let _nw168 = new _module.C0();
        _nw168.__ctor();
        _1102_v105 = _nw168;
        let _1103_v106;
        _1103_v106 = new _dafny.CodePoint('s'.codePointAt(0));
        _1103_v106 = _1103_v106;
      }
      let _1104_v107;
      _1104_v107 = new BigNumber(521);
      _1104_v107 = (_dafny.ZERO).minus(_1104_v107);
      return;
    }
    m6(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Seq.of();
      let r3 = _dafny.ZERO;
      let _1105_v0;
      _1105_v0 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(467)), function (_1106_i0) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      }),((_dafny.ZERO).minus(p2)).plus(p0));
      let _1107_v1;
      _1107_v1 = _dafny.Seq.UnicodeFromString("veytmuod");
      _1105_v0 = (_1105_v0).update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("kxkg"), _dafny.Seq.update(_1107_v1, _module.__default.safeIndex((_this).fm1(_dafny.Map.Empty.slice().updateUnsafe(!((_this).f10),(_this).f10), globalState), new BigNumber((_1107_v1).length)), new _dafny.CodePoint('u'.codePointAt(0)))), p2);
      if (!((((_this).f10) ? ((_this).f10) : ((_this).f10)))) {
        (globalState).f8 = (_this).f10;
        let _1108_v2;
        _1108_v2 = _module.D0.create_DC0(false);
        let _1109_v3;
        _1109_v3 = _dafny.Seq.of((_this).f10, (_this).f10);
        let _1110_v4;
        _1110_v4 = _module.D8.create_DC23(_1109_v3);
        let _1111_v5;
        _1111_v5 = _dafny.Map.Empty.slice().updateUnsafe(_1108_v2,_1110_v4);
        r3 = new BigNumber(((_1111_v5).update(((true) ? (_1108_v2) : (_1108_v2)), _1110_v4)).length);
        let _1112_v6;
        _1112_v6 = new _dafny.CodePoint('s'.codePointAt(0));
        _1112_v6 = _1112_v6;
        let _1113_v7;
        _1113_v7 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f10),p2);
        _1113_v7 = (_1113_v7).update((_this).f10, new BigNumber((_module.__default.fm47((_this).f10, globalState)).length));
        let _1114_v8;
        _1114_v8 = _module.D6.create_DC16((_this).f9);
        _1114_v8 = _module.D6.create_DC16(_module.__default.fm48((_this).f10, globalState));
      } else {
        let _1115_v9;
        let _nw169 = new _module.C1();
        _nw169.__ctor((_this).f10, _1107_v1);
        _1115_v9 = _nw169;
        let _1116_v10;
        _1116_v10 = _dafny.Seq.of(!(true), (_1115_v9).f12);
        let _1117_v11;
        _1117_v11 = _module.D6.create_DC17((_1115_v9).f12, (_this).f10, new BigNumber(((_1115_v9).f13).length), new BigNumber(((_1115_v9).f13).length), new BigNumber((_dafny.Seq.of(p2, new BigNumber((_module.__default.fm45((_1115_v9).f12, p2, (_1115_v9).f12, globalState)).length), new BigNumber((_1116_v10).length), p0, p0)).length));
        let _1118_v12;
        _1118_v12 = _module.D6.create_DC20(_1117_v11);
        let _source17 = _1118_v12;
        if (_source17.is_DC17) {
          let _1119___mcc_h0 = (_source17).cf29;
          let _1120___mcc_h1 = (_source17).cf30;
          let _1121___mcc_h2 = (_source17).cf31;
          let _1122___mcc_h3 = (_source17).cf32;
          let _1123___mcc_h4 = (_source17).cf33;
          let _1124_cf33 = _1123___mcc_h4;
          let _1125_cf32 = _1122___mcc_h3;
          let _1126_cf31 = _1121___mcc_h2;
          let _1127_cf30 = _1120___mcc_h1;
          let _1128_cf29 = _1119___mcc_h0;
          let _1129_v13;
          let _init25 = ((_1130_cf29) => function (_1131_i1) {
            return _1130_cf29;
          })(_1128_cf29);
          let _nw170 = Array((new BigNumber(21)).toNumber());
          for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw170.length); _i0_25++) {
            _nw170[_i0_25] = _init25(new BigNumber(_i0_25));
          }
          _1129_v13 = _nw170;
          let _1132_v14;
          _1132_v14 = _dafny.Map.Empty.slice().updateUnsafe(!(_1127_cf30),_1129_v13);
          let _1133_v15;
          _1133_v15 = _dafny.Map.Empty.slice().updateUnsafe(_1124_cf33,_1128_cf29);
          _1132_v14 = (_1132_v14).update((((_1133_v15).contains((_dafny.ZERO).minus(_1126_cf31))) ? ((_1133_v15).get((_dafny.ZERO).minus(_1126_cf31))) : (_1127_cf30)), _1129_v13);
          (globalState).f8 = (_1115_v9).f12;
          let _1134_v16;
          _1134_v16 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_1115_v9).f13).length),_module.__default.fm33((_1115_v9).f12, _1128_cf29, new BigNumber(-327), _1128_cf29, globalState));
          let _1135_v17;
          _1135_v17 = _dafny.Seq.of(_1134_v16, _1134_v16, _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_1107_v1).length)));
          let _1136_v18;
          _1136_v18 = _dafny.Set.fromElements((_1115_v9).f12, false, (_this).fm0(_1116_v10, _1128_cf29, _1126_cf31, (_1115_v9).f12, globalState), !((_1115_v9).f12), true);
          let _1137_v19;
          _1137_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1136_v18,_1125_cf32);
          let _1138_v20;
          _1138_v20 = new _dafny.CodePoint('a'.codePointAt(0));
          let _1139_v21;
          _1139_v21 = _module.D4.create_DC12(_1127_cf30, new BigNumber(((_1115_v9).f13).length), _1138_v20, (_dafny.ZERO).minus(_1126_cf31), _1124_cf33);
          let _pat_let_tv23 = _1128_cf29;
          let _pat_let_tv24 = p2;
          _1127_cf30 = ((new BigNumber(((_1135_v17)[_module.__default.safeIndex(p0, new BigNumber((_1135_v17).length))]).length)).multipliedBy(new BigNumber((_1137_v19).length))).isEqualTo((function (_pat_let22_0) {
            return function (_1140_dt__update__tmp_h0) {
              return function (_pat_let23_0) {
                return function (_1141_dt__update_hcf19_h0) {
                  return function (_pat_let24_0) {
                    return function (_1142_dt__update_hcf23_h0) {
                      return _module.D4.create_DC12(_1141_dt__update_hcf19_h0, (_1140_dt__update__tmp_h0).dtor_cf20, (_1140_dt__update__tmp_h0).dtor_cf21, (_1140_dt__update__tmp_h0).dtor_cf22, _1142_dt__update_hcf23_h0);
                    }(_pat_let24_0);
                  }(_pat_let_tv24);
                }(_pat_let23_0);
              }(_pat_let_tv23);
            }(_pat_let22_0);
          }(_1139_v21)).dtor_cf20);
          let _1143_v22;
          _1143_v22 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_module.__default.fm33((_1115_v9).f12, (_1115_v9).f12, _1126_cf31, (_1115_v9).f12, globalState)).isLessThanOrEqualTo(_1125_cf32));
          _1143_v22 = _1143_v22;
        } else if (_source17.is_DC18) {
          let _1144___mcc_h5 = (_source17).cf34;
          let _1145___mcc_h6 = (_source17).cf35;
          let _1146___mcc_h7 = (_source17).cf36;
          let _1147_cf36 = _1146___mcc_h7;
          let _1148_cf35 = _1145___mcc_h6;
          let _1149_cf34 = _1144___mcc_h5;
          let _1150_v23;
          let _nw171 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.of());
          _1150_v23 = _nw171;
          let _nw172 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.of());
          _1150_v23 = _nw172;
          let _1151_v25;
          let _init26 = ((_1152_v9, _1153_v10, _1154_p0, _1155_cf36, _1156_cf35) => function (_1157_i2) {
            return (((_1152_v9).f12) ? (function () {
              let _coll45 = new _dafny.Set();
              for (const _compr_45 of (_dafny.MultiSet.fromElements(_1154_p0, new BigNumber(604), _1155_cf36, new BigNumber(916), _1156_cf35)).Elements) {
                let _1158_v24 = _compr_45;
                if ((_dafny.MultiSet.fromElements(_1154_p0, new BigNumber(604), _1155_cf36, new BigNumber(916), _1156_cf35)).contains(_1158_v24)) {
                  _coll45.add((_1158_v24).minus(new BigNumber(((_module.D8.create_DC23(_dafny.Seq.of(false))).dtor_cf41).length)));
                }
              }
              return _coll45;
            }()) : (_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_1153_v10),!((_1152_v9).f12))).length))));
          })(_1115_v9, _1116_v10, p0, _1147_cf36, _1148_cf35);
          let _nw173 = Array((new BigNumber(2)).toNumber());
          for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw173.length); _i0_26++) {
            _nw173[_i0_26] = _init26(new BigNumber(_i0_26));
          }
          _1151_v25 = _nw173;
          let _nw174 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Set.Empty);
          _1151_v25 = _nw174;
          (globalState).f8 = (p0).isLessThan((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(331)), ((_1159_p0) => function (_1160_i3) {
            return _1159_p0;
          })(p0))).length)).minus(p2));
          _1107_v1 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.update(_module.__default.fm22(_1148_cf35, (_1115_v9).f12, (_this).f10, p0, globalState), _module.__default.safeIndex(new BigNumber(959), new BigNumber((_module.__default.fm22(_1148_cf35, (_1115_v9).f12, (_this).f10, p0, globalState)).length)), _1149_cf34), (_1115_v9).f13), _dafny.Seq.Create(_module.__default.abs(new BigNumber(429)), ((_1161_cf34) => function (_1162_i4) {
            return _1161_cf34;
          })(_1149_cf34)));
        } else if (_source17.is_DC19) {
          let _1163___mcc_h8 = (_source17).cf37;
          let _1164___mcc_h9 = (_source17).cf38;
          let _1165_cf38 = _1164___mcc_h9;
          let _1166_cf37 = _1163___mcc_h8;
          r1 = p2;
          let _1167_v26;
          _1167_v26 = _dafny.Seq.of(p2);
          (globalState).f8 = (!(p0).isEqualTo((_1167_v26)[_module.__default.safeIndex(new BigNumber((_1116_v10).length), new BigNumber((_1167_v26).length))])) || ((_1115_v9).fm32((_this).fm1(((_this).f9).update(_1166_cf37, (_1115_v9).f12), globalState), _1165_cf38, globalState));
          let _1168_v27;
          let _nw175 = new _module.C1();
          _nw175.__ctor((_this).f10, _dafny.Seq.update((_1115_v9).f13, _module.__default.safeIndex(p0, new BigNumber(((_1115_v9).f13).length)), _1165_cf38));
          _1168_v27 = _nw175;
          (globalState).f8 = (p0).isEqualTo((_dafny.ZERO).minus(new BigNumber(-486)));
        } else if (_source17.is_DC16) {
          let _1169___mcc_h10 = (_source17).cf28;
          let _1170_cf28 = _1169___mcc_h10;
          let _1171_v28;
          _1171_v28 = _module.D7.create_DC22();
          let _1172_v29;
          let _nw176 = Array((new BigNumber(2)).toNumber()).fill(false);
          _1172_v29 = _nw176;
          let _1173_v30;
          _1173_v30 = _dafny.Set.fromElements((_this).f10, (_1115_v9).f12);
          let _1174_v31;
          _1174_v31 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f10);
          let _1175_v32;
          _1175_v32 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('e'.codePointAt(0)),(_1115_v9).f12);
          let _1176_v33;
          _1176_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1175_v32,_1173_v30);
          let _1177_v34;
          _1177_v34 = _dafny.Seq.of(p1, p1);
          let _rhs176 = _1171_v28;
          let _rhs177 = ((p2).multipliedBy(p2)).minus(p2);
          let _rhs178 = _1172_v29;
          let _rhs179 = (_module.__default.fm34(new BigNumber((_1174_v31).length), (_this).f10, (_1115_v9).f12, (_1115_v9).f12, globalState)).Union((((_1176_v33).contains(_1175_v32)) ? ((_1176_v33).get(_1175_v32)) : (_1173_v30)));
          let _rhs180 = new BigNumber((_1177_v34).length);
          _1171_v28 = _rhs176;
          r1 = _rhs177;
          _1172_v29 = _rhs178;
          _1173_v30 = _rhs179;
          r3 = _rhs180;
          let _1178_v35;
          _1178_v35 = _dafny.Seq.of(p2);
          let _1179_v36;
          _1179_v36 = _module.D1.create_DC2(_1178_v35);
          let _1180_v37;
          _1180_v37 = _dafny.Set.fromElements(_1179_v36);
          let _1181_v38;
          _1181_v38 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsPrefixOf((_1115_v9).f13, (_1115_v9).f13),_1180_v37);
          let _1182_v39;
          _1182_v39 = new _dafny.CodePoint('q'.codePointAt(0));
          let _1183_v40;
          _1183_v40 = _dafny.Map.Empty.slice().updateUnsafe(_1182_v39,_1182_v39);
          let _rhs181 = !(_1183_v40).contains(_1182_v39);
          let _rhs182 = (_1181_v38).Merge(_1181_v38);
          let _rhs183 = (_1173_v30).IsSubsetOf(_1173_v30);
          let _lhs129 = globalState;
          let _lhs130 = globalState;
          _lhs129.f8 = _rhs181;
          _1181_v38 = _rhs182;
          _lhs130.f8 = _rhs183;
          let _1184_v41;
          _1184_v41 = _dafny.Map.Empty.slice().updateUnsafe(!((_dafny.ZERO).minus(new BigNumber((_1174_v31).length))).isEqualTo(p0),(_dafny.ZERO).minus((p2).plus(p0)));
          _1184_v41 = _1184_v41;
          let _index177 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((_1172_v29).length));
          (_1172_v29)[_index177] = (_1115_v9).f12;
          let _index178 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((p1).length));
          (p1)[_index178] = p0;
          let _index179 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((_1172_v29).length));
          let _index180 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((p1).length));
          let _rhs184 = (p2).isLessThan(p0);
          let _rhs185 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).fm1(_1170_cf28, globalState)));
          let _rhs186 = p2;
          let _rhs187 = p0;
          let _rhs188 = (_1115_v9).f12;
          let _lhs131 = _1172_v29;
          let _lhs132 = _module.__default.safeIndex(new BigNumber(129), new BigNumber((_1172_v29).length));
          let _lhs133 = p1;
          let _lhs134 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((p1).length));
          let _lhs135 = globalState;
          _lhs131[_lhs132] = _rhs184;
          r3 = _rhs185;
          r3 = _rhs186;
          _lhs133[_lhs134] = _rhs187;
          _lhs135.f8 = _rhs188;
        } else {
          let _1185___mcc_h11 = (_source17).cf39;
          let _1186_cf39 = _1185___mcc_h11;
          let _1187_v42;
          _1187_v42 = _dafny.Set.fromElements((_1115_v9).f12, (_this).f10);
          let _1188_v43;
          _1188_v43 = _module.D3.create_DC8((_1187_v42).IsSubsetOf(_1187_v42), (_this).f10);
          _1188_v43 = _module.__default.fm35(false, globalState);
          let _1189_v44;
          _1189_v44 = _dafny.Map.Empty.slice().updateUnsafe((_1115_v9).f12,new BigNumber(((_this).f9).length));
          let _1190_v45;
          _1190_v45 = _dafny.Set.fromElements(new BigNumber((_1189_v44).length));
          let _1191_v46;
          _1191_v46 = _module.D4.create_DC11(new BigNumber((_1116_v10).length), _1190_v45, !((_1115_v9).f12));
          let _1192_v47;
          _1192_v47 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(309),p2);
          let _1193_v48;
          _1193_v48 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_1192_v47).length));
          let _pat_let_tv25 = _1190_v45;
          let _1194_v50;
          _1194_v50 = _dafny.Seq.of(_1191_v46, function (_pat_let25_0) {
            return function (_1199_dt__update__tmp_h1) {
              return function (_pat_let26_0) {
                return function (_1200_dt__update_hcf17_h0) {
                  return _module.D4.create_DC11((_1199_dt__update__tmp_h1).dtor_cf16, _1200_dt__update_hcf17_h0, (_1199_dt__update__tmp_h1).dtor_cf18);
                }(_pat_let26_0);
              }(_pat_let_tv25);
            }(_pat_let25_0);
          }(_module.D4.create_DC11(new BigNumber((_1192_v47).length), function () {
  let _coll46 = new _dafny.Set();
  for (const _compr_46 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-534)), ((_1195_p0) => function (_1196_i5) {
    return _1195_p0;
  })(p0))).Elements) {
    let _1197_v49 = _compr_46;
    if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-534)), ((_1198_p0) => function (_1196_i5) {
      return _1198_p0;
    })(p0)), _1197_v49)) {
      _coll46.add(_module.__default.safeModuloInt(_1197_v49, new BigNumber(-700)));
    }
  }
  return _coll46;
}(), (_1115_v9).f12)), _1191_v46);
          let _1201_v51;
          _1201_v51 = _1194_v50;
          let _1202_v52;
          _1202_v52 = _dafny.Map.Empty.slice().updateUnsafe(_1194_v50,new BigNumber(-803));
          let _rhs189 = !((_module.D14.create_DC34(_1202_v52)).dtor_cf61).contains((_1201_v51));
          let _rhs190 = p2;
          let _lhs136 = globalState;
          _lhs136.f8 = _rhs189;
          r1 = _rhs190;
          _1107_v1 = _dafny.Seq.Concat((_1115_v9).f13, _dafny.Seq.Concat(_1107_v1, (_1115_v9).f13));
          let _1203_v53;
          _1203_v53 = _dafny.Map.Empty.slice().updateUnsafe(_1201_v51,_1190_v45);
          _1189_v44 = (_1189_v44).update((p0).isLessThanOrEqualTo(new BigNumber((_1203_v53).length)), _module.__default.safeModuloInt(new BigNumber((_1107_v1).length), p0));
        }
        let _1204_v54;
        let _nw177 = new _module.C3();
        _nw177.__ctor();
        _1204_v54 = _nw177;
        let _1205_v55;
        let _init27 = ((_1206_v10, _1207_p2) => function (_1208_i6) {
          return (new BigNumber((_1206_v10).length)).isLessThan(_1207_p2);
        })(_1116_v10, p2);
        let _nw178 = Array((new BigNumber(27)).toNumber());
        for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw178.length); _i0_27++) {
          _nw178[_i0_27] = _init27(new BigNumber(_i0_27));
        }
        _1205_v55 = _nw178;
        let _index181 = _module.__default.safeIndex(new BigNumber(107), new BigNumber((_1205_v55).length));
        (_1205_v55)[_index181] = (_1115_v9).f12;
        let _index182 = _module.__default.safeIndex(new BigNumber(107), new BigNumber((_1205_v55).length));
        (_1205_v55)[_index182] = (_1116_v10)[_module.__default.safeIndex(p0, new BigNumber((_1116_v10).length))];
        let _index183 = _module.__default.safeIndex(new BigNumber(197), new BigNumber((p1).length));
        (p1)[_index183] = (new BigNumber(207)).plus(p0);
        let _index184 = _module.__default.safeIndex(new BigNumber(197), new BigNumber((p1).length));
        (p1)[_index184] = p2;
      }
      r1 = new BigNumber((_dafny.Seq.Concat(_module.__default.fm45(true, p2, true, globalState), _dafny.Seq.Concat(_1107_v1, _1107_v1))).length);
      let _1209_v56;
      let _nw179 = Array((new BigNumber(14)).toNumber()).fill(false);
      _1209_v56 = _nw179;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1209_v56).length))) {
        let _1210_i7 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1210_i7)) && ((_1210_i7).isLessThan(new BigNumber((_1209_v56).length))))) {
          (_1209_v56)[(_1210_i7)] = !((_this).f10);
        }
      }
      let _1211_v57;
      _1211_v57 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f10);
      let _1212_v58;
      _1212_v58 = _dafny.Seq.of((_this).f10);
      (globalState).f8 = (((_1211_v57).contains(p0)) ? ((_1211_v57).get(p0)) : (_dafny.Seq.contains(_1212_v58, (_this).f10)));
      let _1213_v59;
      let _nw180 = new _module.C2();
      _nw180.__ctor((_this).f10, _1209_v56, (_dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_this).f10)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,!(!((_this).f10)))), (_this).f10);
      _1213_v59 = _nw180;
      r0 = _dafny.Seq.of(new BigNumber(831));
      let _1214_v60;
      _1214_v60 = _dafny.Seq.of(p0, p0);
      let _1215_v61;
      _1215_v61 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
      let _1216_v62;
      _1216_v62 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_dafny.Seq.Concat(_1214_v60, _dafny.Seq.update(_dafny.Seq.of((_dafny.ZERO).minus(p0), (((_1215_v61).contains(new BigNumber((_1214_v60).length))) ? ((_1215_v61).get(new BigNumber((_1214_v60).length))) : (p2)), p2, new BigNumber((_dafny.Seq.UnicodeFromString("vyqvooyef")).length)), _module.__default.safeIndex((_1213_v59).fm42(p0, p2, p2, globalState), new BigNumber((_dafny.Seq.of((_dafny.ZERO).minus(p0), (((_1215_v61).contains(new BigNumber((_1214_v60).length))) ? ((_1215_v61).get(new BigNumber((_1214_v60).length))) : (p2)), p2, new BigNumber((_dafny.Seq.UnicodeFromString("vyqvooyef")).length))).length)), p0)));
      r1 = new BigNumber((_1216_v62).length);
      r2 = _module.__default.fm37(globalState);
      r3 = p2;
      return [r0, r1, r2, r3];
    }
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f9 = _dafny.Map.Empty;
      this._f10 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f10() {
      let _this = this;
      return _this._f10;
    };
    __ctor(f9, f10) {
      let _this = this;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      return;
    }
    fm0(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !((_module.D0.create_DC1(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements((_this).f10)).cardinality()))).length), (_this).f10, new BigNumber(-312))).dtor_cf2);
    };
    fm1(p0, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(138), new BigNumber((_dafny.Seq.of((_this).f10, (_this).f10)).length)));
    };
    fm18(p0, p1, globalState) {
      let _this = this;
      return ((_dafny.ZERO).minus((new BigNumber(-298)).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(133),_dafny.Seq.UnicodeFromString("n"))).length)))).isLessThan((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-447),(_this).f10)).length), new BigNumber(686))));
    };
    fm19(p0, p1, globalState) {
      let _this = this;
      return new BigNumber(((((_this).f10) ? ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_module.D1.create_DC5(_module.D1.create_DC2(_dafny.Seq.of(new BigNumber(725)))), _module.D1.create_DC5(_module.D1.create_DC4((_this).f10, (_this).f10, new BigNumber(103)))),_dafny.Map.Empty.slice().updateUnsafe((_this).f10,_module.D1.create_DC3(_dafny.Seq.UnicodeFromString("pppcgeiju"))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(322)), function (_1217_i0) {
        return _module.D1.create_DC5(_module.D1.create_DC3(_dafny.Seq.UnicodeFromString("wjqjefffv")));
      }),_dafny.Map.Empty.slice().updateUnsafe((_this).f10,_module.D1.create_DC3(_dafny.Seq.UnicodeFromString("ojakub")))))) : ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_module.D1.create_DC5(_module.D1.create_DC3(_dafny.Seq.UnicodeFromString("vlvw")))),_dafny.Map.Empty.slice().updateUnsafe((_this).f10,_module.D1.create_DC3(_dafny.Seq.UnicodeFromString("i"))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_module.D1.create_DC5(_module.D1.create_DC3(_dafny.Seq.UnicodeFromString("oa"))), _module.D1.create_DC5(_module.D1.create_DC3(_dafny.Seq.UnicodeFromString("lualcqxw")))),_dafny.Map.Empty.slice().updateUnsafe((_this).f10,_module.D1.create_DC3(_dafny.Seq.UnicodeFromString("jlo")))))))).length);
    };
    m0(globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = false;
      let r2 = _dafny.Seq.of();
      let r3 = _dafny.Seq.UnicodeFromString("");
      let _1218_v0;
      _1218_v0 = _dafny.Seq.UnicodeFromString("avx");
      let _1219_v1;
      _1219_v1 = new BigNumber(649);
      let _hi2 = _module.__default.safeModuloInt(_1219_v1, _1219_v1);
      for (let _1220_i0 = new BigNumber((_1218_v0).length); _1220_i0.isLessThan(_hi2); _1220_i0 = _1220_i0.plus(_dafny.ONE)) {
        let _1221_v2;
        let _nw181 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
        _1221_v2 = _nw181;
        let _index185 = _module.__default.safeIndex(new BigNumber(880), new BigNumber((_1221_v2).length));
        (_1221_v2)[_index185] = _1220_i0;
        let _1222_v3;
        _1222_v3 = _dafny.Seq.of(_1218_v0, _module.__default.fm20(_1220_i0, (_this).f10, (_this).f10, (_this).fm19(_1219_v1, _1220_i0, globalState), globalState));
        let _index186 = _module.__default.safeIndex(new BigNumber(880), new BigNumber((_1221_v2).length));
        (_1221_v2)[_index186] = new BigNumber(((_1222_v3)[_module.__default.safeIndex(_1220_i0, new BigNumber((_1222_v3).length))]).length);
        _1219_v1 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((new BigNumber(524)).minus((_1221_v2)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_1221_v2).length))]), _1219_v1));
        let _1223_v4;
        let _1224_v5;
        let _1225_v6;
        let _out56;
        let _out57;
        let _out58;
        let _outcollector16 = (_this).m4((_this).f10, (_this).f10, globalState);
        _out56 = _outcollector16[0];
        _out57 = _outcollector16[1];
        _out58 = _outcollector16[2];
        _1223_v4 = _out56;
        _1224_v5 = _out57;
        _1225_v6 = _out58;
        if ((_this).f10) {
          let _index187 = _module.__default.safeIndex(new BigNumber(230), new BigNumber((_1223_v4).length));
          (_1223_v4)[_index187] = true;
          let _1226_v7;
          _1226_v7 = _dafny.Set.fromElements((_this).f10);
          let _index188 = _module.__default.safeIndex(new BigNumber(230), new BigNumber((_1223_v4).length));
          (_1223_v4)[_index188] = (_1226_v7).IsSubsetOf(_1226_v7);
          let _1227_v8;
          _1227_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1220_i0,_dafny.Seq.of(new BigNumber((_dafny.Seq.of(_1219_v1)).length), _1219_v1, _1220_i0, _1219_v1));
          let _1228_v9;
          _1228_v9 = _dafny.Seq.of(_1220_i0);
          _1227_v8 = (_1227_v8).update(_1220_i0, _1228_v9);
          _1218_v0 = _1218_v0;
          r0 = _1218_v0;
          r1 = true;
        } else {
          let _index189 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_1223_v4).length));
          (_1223_v4)[_index189] = (_this).f10;
          let _index190 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_1223_v4).length));
          (_1223_v4)[_index190] = (_this).f10;
          let _1229_v10;
          _1229_v10 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm21(globalState),(_1220_i0).plus(new BigNumber((_1218_v0).length)));
          let _1230_v11;
          _1230_v11 = _dafny.Seq.of(new BigNumber(284), new BigNumber((_1218_v0).length));
          let _1231_v12;
          _1231_v12 = _module.D1.create_DC2(_1230_v11);
          _1229_v10 = (_1229_v10).update(_1231_v12, _1219_v1);
          let _index191 = _module.__default.safeIndex(new BigNumber(880), new BigNumber((_1221_v2).length));
          (_1221_v2)[_index191] = (_1230_v11)[_module.__default.safeIndex((_1221_v2)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_1221_v2).length))], new BigNumber((_1230_v11).length))];
          r0 = _dafny.Seq.Concat(_1218_v0, _1218_v0);
          (globalState).f8 = (_1223_v4)[_module.__default.safeIndex(new BigNumber(331), new BigNumber((_1223_v4).length))];
        }
      }
      _1219_v1 = _1219_v1;
      let _1232_v13;
      let _init28 = function (_1233_i1) {
        return !_dafny.Seq.contains(_dafny.Seq.of((_this).f10, (_this).f10, (_this).f10), (_this).f10);
      };
      let _nw182 = Array((new BigNumber(24)).toNumber());
      for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw182.length); _i0_28++) {
        _nw182[_i0_28] = _init28(new BigNumber(_i0_28));
      }
      _1232_v13 = _nw182;
      let _index192 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_1232_v13).length));
      (_1232_v13)[_index192] = ((((_this).f9).contains((_this).f10)) ? (((_this).f9).get((_this).f10)) : (true));
      let _index193 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_1232_v13).length));
      (_1232_v13)[_index193] = true;
      let _1234_v14;
      _1234_v14 = _dafny.Set.fromElements(_1218_v0, _1218_v0, _1218_v0);
      let _1235_v15;
      _1235_v15 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1234_v14).length),(_dafny.ZERO).minus(new BigNumber((_1218_v0).length)));
      let _1236_v16;
      let _nw183 = new _module.C2();
      _nw183.__ctor((_this).f10, _1232_v13, _dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_1232_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_1232_v13).length))]), (_1232_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_1232_v13).length))]);
      _1236_v16 = _nw183;
      let _1237_v17;
      _1237_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1235_v15,_1236_v16);
      let _1238_v19;
      _1238_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1219_v1,_1236_v16);
      _1237_v17 = (_1237_v17).update(function () {
        let _coll47 = new _dafny.Map();
        for (const _compr_47 of _dafny.IntegerRange(new BigNumber(218), new BigNumber(-929))) {
          let _1239_v18 = _compr_47;
          if (((new BigNumber(218)).isLessThanOrEqualTo(_1239_v18)) && ((_1239_v18).isLessThan(new BigNumber(-929)))) {
            _coll47.push([(_1239_v18).minus(new BigNumber(413)),_1219_v1]);
          }
        }
        return _coll47;
      }(), (((_1238_v19).contains(_1219_v1)) ? ((_1238_v19).get(_1219_v1)) : (_1236_v16)));
      let _1240_v20;
      let _nw184 = new _module.C1();
      _nw184.__ctor((_1232_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_1232_v13).length))], _1218_v0);
      _1240_v20 = _nw184;
      let _index194 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_1232_v13).length));
      let _rhs191 = !((_1232_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_1232_v13).length))]);
      let _rhs192 = (((_1232_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_1232_v13).length))]) ? (_1219_v1) : (_1219_v1));
      let _rhs193 = (_dafny.ZERO).minus(_1219_v1);
      let _rhs194 = _1240_v20;
      let _rhs195 = (_1232_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_1232_v13).length))];
      let _lhs137 = _1232_v13;
      let _lhs138 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_1232_v13).length));
      let _lhs139 = globalState;
      _lhs137[_lhs138] = _rhs191;
      _1219_v1 = _rhs192;
      _1219_v1 = _rhs193;
      _1240_v20 = _rhs194;
      _lhs139.f8 = _rhs195;
      let _1241_i2;
      _1241_i2 = _dafny.ZERO;
      L5: {
        while ((_1240_v20).f12) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1241_i2)) {
              break L5;
            }
            _1241_i2 = (_1241_i2).plus(_dafny.ONE);
            let _1242_v21;
            _1242_v21 = _dafny.MultiSet.fromElements(_1218_v0, _1218_v0);
            let _1243_v22;
            _1243_v22 = _module.D11.create_DC27(_1242_v21);
            let _1244_v23;
            _1244_v23 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_1243_v22);
            _1244_v23 = _1244_v23;
            let _1245_v24;
            _1245_v24 = _dafny.Map.Empty.slice().updateUnsafe((_1240_v20).f12,(_1240_v20).f12);
            _1245_v24 = (((_1240_v20).f12) ? ((_1236_v16).f9) : ((_1236_v16).f9));
            if ((_1240_v20).f12) {
              let _index195 = _module.__default.safeIndex(new BigNumber(795), new BigNumber((_1232_v13).length));
              (_1232_v13)[_index195] = ((_dafny.ZERO).minus(_1219_v1)).isLessThan((_1219_v1).plus(_1219_v1));
              let _1246_v25;
              _1246_v25 = _dafny.Seq.of(_1232_v13, _1232_v13, _1232_v13, _1232_v13);
              _1246_v25 = _1246_v25;
              (globalState).f8 = false;
              let _1247_v26;
              let _nw185 = new _module.C4();
              _nw185.__ctor((_1236_v16).f9, (_1236_v16).f10);
              _1247_v26 = _nw185;
              let _1248_v27;
              _1248_v27 = _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(113)), function (_1249_i3) {
                return new _dafny.CodePoint('v'.codePointAt(0));
              })).length));
              let _1250_v28;
              _1250_v28 = _dafny.Set.fromElements(new BigNumber((_1248_v27).length));
              let _1251_v29;
              _1251_v29 = _module.D6.create_DC17((_1236_v16).f10, true, _1219_v1, _1219_v1, new BigNumber((_1235_v15).length));
              let _1252_v30;
              _1252_v30 = _module.D4.create_DC11(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1247_v26,_1219_v1)).length), _1250_v28, !((_1251_v29).dtor_cf30));
              let _1253_v31;
              _1253_v31 = _dafny.Seq.of(_1252_v30, _1252_v30, _module.D4.create_DC11(new BigNumber(((_1236_v16).f9).length), _1250_v28, (_1232_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_1232_v13).length))]));
              let _1254_v32;
              _1254_v32 = _1253_v31;
              _1254_v32 = _1254_v32;
              let _1255_v33;
              _1255_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1219_v1,_1232_v13);
              _1255_v33 = (_1255_v33).update(_1219_v1, (((_1255_v33).contains(_1219_v1)) ? ((_1255_v33).get(_1219_v1)) : (_1232_v13)));
            } else {
              let _1256_v34;
              let _nw186 = Array((new BigNumber(27)).toNumber()).fill(_module.D1.Default());
              _1256_v34 = _nw186;
              let _1257_v35;
              _1257_v35 = _module.D1.create_DC4((_1232_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_1232_v13).length))], (_1236_v16).f10, _1219_v1);
              let _1258_v36;
              _1258_v36 = _dafny.Map.Empty.slice().updateUnsafe(_1256_v34,(_dafny.ZERO).minus((_1257_v35).dtor_cf8));
              let _1259_v37;
              _1259_v37 = _module.D15.create_DC37(_1256_v34);
              _1258_v36 = (_1258_v36).update((_1259_v37).dtor_cf68, new BigNumber(207));
              let _1260_v38;
              _1260_v38 = new _dafny.CodePoint('l'.codePointAt(0));
              let _1261_v39;
              let _1262_v40;
              let _out59;
              let _out60;
              let _outcollector17 = (_1236_v16).m1(_1260_v38, globalState);
              _out59 = _outcollector17[0];
              _out60 = _outcollector17[1];
              _1261_v39 = _out59;
              _1262_v40 = _out60;
              let _1263_v41;
              _1263_v41 = _dafny.Seq.of(_1262_v40);
              let _1264_v42;
              _1264_v42 = _dafny.MultiSet.fromElements(_1219_v1);
              let _1265_v43;
              _1265_v43 = _dafny.Seq.of(_1264_v42);
              let _1266_v44;
              _1266_v44 = _dafny.Map.Empty.slice().updateUnsafe((_1232_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_1232_v13).length))],_1236_v16);
              let _1267_v45;
              let _nw187 = Array((new BigNumber(23)).toNumber());
              _nw187[(_dafny.ZERO).toNumber()] = _1219_v1;
              _nw187[(_dafny.ONE).toNumber()] = _1219_v1;
              _nw187[(new BigNumber(2)).toNumber()] = _1219_v1;
              _nw187[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus((_this).fm19(_1219_v1, _1219_v1, globalState));
              _nw187[(new BigNumber(4)).toNumber()] = _1219_v1;
              _nw187[(new BigNumber(5)).toNumber()] = _1219_v1;
              _nw187[(new BigNumber(6)).toNumber()] = _1219_v1;
              _nw187[(new BigNumber(7)).toNumber()] = _1219_v1;
              _nw187[(new BigNumber(8)).toNumber()] = _1219_v1;
              _nw187[(new BigNumber(9)).toNumber()] = _1219_v1;
              _nw187[(new BigNumber(10)).toNumber()] = new BigNumber(((_1265_v43)[_module.__default.safeIndex(_1219_v1, new BigNumber((_1265_v43).length))]).cardinality());
              _nw187[(new BigNumber(11)).toNumber()] = new BigNumber(-273);
              _nw187[(new BigNumber(12)).toNumber()] = _1219_v1;
              _nw187[(new BigNumber(13)).toNumber()] = new BigNumber((_1266_v44).length);
              _nw187[(new BigNumber(14)).toNumber()] = _1219_v1;
              _nw187[(new BigNumber(15)).toNumber()] = _1219_v1;
              _nw187[(new BigNumber(16)).toNumber()] = _1219_v1;
              _nw187[(new BigNumber(17)).toNumber()] = new BigNumber(-283);
              _nw187[(new BigNumber(18)).toNumber()] = _1219_v1;
              _nw187[(new BigNumber(19)).toNumber()] = new BigNumber((_1235_v15).length);
              _nw187[(new BigNumber(20)).toNumber()] = _1219_v1;
              _nw187[(new BigNumber(21)).toNumber()] = _1219_v1;
              _nw187[(new BigNumber(22)).toNumber()] = new BigNumber(485);
              _1267_v45 = _nw187;
              let _1268_v46;
              _1268_v46 = _dafny.Map.Empty.slice().updateUnsafe(_1267_v45,(_1236_v16).f10);
              let _1269_v47;
              _1269_v47 = _dafny.Map.Empty.slice().updateUnsafe(_1260_v38,_1268_v46);
              (globalState).f8 = (_1236_v16).fm0(_dafny.Seq.Concat(_1263_v41, _1263_v41), (_module.__default.fm49(!((_1236_v16).f10), globalState)).dtor_cf29, new BigNumber(((((_1269_v47).contains(_1260_v38)) ? ((_1269_v47).get(_1260_v38)) : (_1268_v46))).length), (_1236_v16).fm0(_1263_v41, (_1236_v16).f10, _1219_v1, (_this).f10, globalState), globalState);
              let _1270_v48;
              _1270_v48 = _dafny.MultiSet.fromElements((_1236_v16).f10);
              let _1271_v49;
              _1271_v49 = _dafny.Map.Empty.slice().updateUnsafe(_1270_v48,_1232_v13);
              (globalState).f8 = !(new BigNumber((_1271_v49).length)).isEqualTo(_1219_v1);
              (globalState).f8 = false;
            }
            let _1272_v50;
            let _nw188 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
            _1272_v50 = _nw188;
            let _1273_v51;
            _1273_v51 = _dafny.Seq.of(_1272_v50, _1272_v50);
            let _1274_v52;
            _1274_v52 = _module.D12.create_DC32(_1272_v50, _1273_v51);
            let _1275_v53;
            _1275_v53 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_1219_v1),_1274_v52);
            let _1276_v54;
            _1276_v54 = _dafny.Map.Empty.slice().updateUnsafe(_1275_v53,_1219_v1);
            _1235_v15 = (_1235_v15).update(new BigNumber(((_1276_v54).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1275_v53,_1219_v1))).length), new BigNumber((_1218_v0).length));
          }
        }
      }
      r0 = _1218_v0;
      r1 = (_1236_v16).f10;
      let _1277_v55;
      _1277_v55 = _dafny.Set.fromElements((_this).f10, (_1240_v20).f12, (_module.D0.create_DC0((_1236_v16).f10)).dtor_cf0, !((_this).f10));
      let _1278_v56;
      _1278_v56 = _dafny.Seq.of(_1277_v55);
      let _1279_v57;
      _1279_v57 = _dafny.Seq.of((_this).f10);
      r2 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1278_v56, _dafny.Seq.Create(_module.__default.abs(new BigNumber(956)), ((_1280_v55) => function (_1281_i4) {
        return _1280_v55;
      })(_1277_v55))), _dafny.Seq.update(_1278_v56, _module.__default.safeIndex(new BigNumber(28), new BigNumber((_1278_v56).length)), _dafny.Set.fromElements((_1236_v16).fm0(_1279_v57, (_1236_v16).f10, _1219_v1, (_1232_v13)[_module.__default.safeIndex(new BigNumber(795), new BigNumber((_1232_v13).length))], globalState))));
      r3 = _module.__default.fm22(_1219_v1, (_this).f10, (_this).f10, new BigNumber(((_1240_v20).f13).length), globalState);
      return [r0, r1, r2, r3];
    }
    m1(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = false;
      let _1282_v0;
      _1282_v0 = new BigNumber(579);
      let _1283_i0;
      _1283_i0 = _dafny.ZERO;
      L6: {
        while ((_1282_v0).isLessThan(_1282_v0)) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1283_i0)) {
              break L6;
            }
            _1283_i0 = (_1283_i0).plus(_dafny.ONE);
            let _1284_v1;
            _1284_v1 = _dafny.Seq.UnicodeFromString("nllkwrs");
            let _rhs196 = (_this).f10;
            let _rhs197 = new BigNumber(994);
            let _rhs198 = (((_this).f10) ? (_dafny.Seq.UnicodeFromString("b")) : (_dafny.Seq.Concat(_1284_v1, _dafny.Seq.UnicodeFromString("aogmc"))));
            let _lhs140 = globalState;
            _lhs140.f8 = _rhs196;
            _1282_v0 = _rhs197;
            _1284_v1 = _rhs198;
            let _1285_v2;
            let _init29 = ((_1286_p0) => function (_1287_i1) {
              return _1286_p0;
            })(p0);
            let _nw189 = Array((new BigNumber(6)).toNumber());
            for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw189.length); _i0_29++) {
              _nw189[_i0_29] = _init29(new BigNumber(_i0_29));
            }
            _1285_v2 = _nw189;
            let _index196 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_1285_v2).length));
            (_1285_v2)[_index196] = p0;
            let _index197 = _module.__default.safeIndex(new BigNumber(475), new BigNumber((_1285_v2).length));
            (_1285_v2)[_index197] = p0;
            let _1288_v3;
            let _init30 = function (_1289_i2) {
              return ((((_this).f9).contains((_this).f10)) ? (((_this).f9).get((_this).f10)) : ((_this).f10));
            };
            let _nw190 = Array((new BigNumber(5)).toNumber());
            for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw190.length); _i0_30++) {
              _nw190[_i0_30] = _init30(new BigNumber(_i0_30));
            }
            _1288_v3 = _nw190;
            let _index198 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1288_v3).length));
            (_1288_v3)[_index198] = (_this).f10;
            let _index199 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1288_v3).length));
            (_1288_v3)[_index199] = (_this).f10;
            let _1290_v4;
            _1290_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-414),(_this).f10);
            _1290_v4 = (_1290_v4).Merge(_module.__default.fm44((_this).f10, globalState));
          }
        }
      }
      let _1291_v5;
      _1291_v5 = _dafny.MultiSet.fromElements((_this).f10);
      _1282_v0 = (_dafny.ZERO).minus(new BigNumber((((_1291_v5).update(!(true), _module.__default.abs(_1282_v0))).Difference(_dafny.MultiSet.fromElements((_this).f10))).cardinality()));
      (globalState).f8 = (_this).f10;
      let _1292_v6;
      _1292_v6 = _dafny.Seq.of(_1282_v0, _1282_v0);
      if (_dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_1292_v6, _dafny.Seq.of(_1282_v0)), _dafny.Seq.Concat(_1292_v6, _1292_v6))) {
        _1282_v0 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_1292_v6, _1292_v6), _1292_v6)).length);
        let _1293_v8;
        _1293_v8 = _dafny.Seq.UnicodeFromString("edykaowmo");
        let _1294_v9;
        let _nw191 = Array((new BigNumber(15)).toNumber());
        _nw191[(_dafny.ZERO).toNumber()] = _1282_v0;
        _nw191[(_dafny.ONE).toNumber()] = _1282_v0;
        _nw191[(new BigNumber(2)).toNumber()] = _1282_v0;
        _nw191[(new BigNumber(3)).toNumber()] = _1282_v0;
        _nw191[(new BigNumber(4)).toNumber()] = _1282_v0;
        _nw191[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1282_v0,(_this).f10)).length);
        _nw191[(new BigNumber(6)).toNumber()] = _1282_v0;
        _nw191[(new BigNumber(7)).toNumber()] = new BigNumber(-569);
        _nw191[(new BigNumber(8)).toNumber()] = _1282_v0;
        _nw191[(new BigNumber(9)).toNumber()] = new BigNumber((function () {
          let _coll48 = new _dafny.Set();
          for (const _compr_48 of _dafny.IntegerRange(new BigNumber(98), new BigNumber(915))) {
            let _1295_v7 = _compr_48;
            if (((new BigNumber(98)).isLessThanOrEqualTo(_1295_v7)) && ((_1295_v7).isLessThan(new BigNumber(915)))) {
              _coll48.add(_module.__default.safeModuloInt(_1295_v7, new BigNumber((_dafny.Seq.of((_this).f10)).length)));
            }
          }
          return _coll48;
        }()).length);
        _nw191[(new BigNumber(10)).toNumber()] = _1282_v0;
        _nw191[(new BigNumber(11)).toNumber()] = new BigNumber((_1293_v8).length);
        _nw191[(new BigNumber(12)).toNumber()] = _1282_v0;
        _nw191[(new BigNumber(13)).toNumber()] = _1282_v0;
        _nw191[(new BigNumber(14)).toNumber()] = _1282_v0;
        _1294_v9 = _nw191;
        let _1296_v10;
        _1296_v10 = _dafny.MultiSet.fromElements(_1294_v9);
        let _1297_v11;
        _1297_v11 = _dafny.Map.Empty.slice().updateUnsafe((_1296_v10).Difference(_1296_v10),_module.__default.fm33((_this).f10, (_this).f10, (_dafny.ZERO).minus(_1282_v0), !((_this).f10), globalState));
        _1297_v11 = (_1297_v11).update(_1296_v10, _1282_v0);
        let _1298_v12;
        let _nw192 = Array((new BigNumber(4)).toNumber());
        _nw192[(_dafny.ZERO).toNumber()] = (_this).f10;
        _nw192[(_dafny.ONE).toNumber()] = (_this).f10;
        _nw192[(new BigNumber(2)).toNumber()] = (_this).f10;
        _nw192[(new BigNumber(3)).toNumber()] = (_this).f10;
        _1298_v12 = _nw192;
        let _1299_v13;
        let _nw193 = new _module.C2();
        _nw193.__ctor((_1282_v0).isLessThanOrEqualTo(_1282_v0), _1298_v12, (_this).f9, (_1282_v0).isLessThan(_1282_v0));
        _1299_v13 = _nw193;
        let _1300_v14;
        _1300_v14 = _dafny.Set.fromElements((_this).f10);
        (_1299_v13).f14 = !((_1300_v14).IsSubsetOf(_1300_v14)) || ((_1282_v0).isLessThanOrEqualTo(new BigNumber(640)));
        _1294_v9 = _1294_v9;
      } else {
        _1282_v0 = _1282_v0;
        let _1301_v15;
        let _init31 = function (_1302_i3) {
          return (_1302_i3).plus(new BigNumber((_dafny.Seq.UnicodeFromString("uneixc")).length));
        };
        let _nw194 = Array((new BigNumber(4)).toNumber());
        for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw194.length); _i0_31++) {
          _nw194[_i0_31] = _init31(new BigNumber(_i0_31));
        }
        _1301_v15 = _nw194;
        let _1303_v16;
        _1303_v16 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_1301_v15);
        _1303_v16 = (_1303_v16).update((_this).f10, _1301_v15);
        let _index200 = _module.__default.safeIndex(new BigNumber(34), new BigNumber((_1301_v15).length));
        (_1301_v15)[_index200] = _1282_v0;
        let _index201 = _module.__default.safeIndex(new BigNumber(34), new BigNumber((_1301_v15).length));
        (_1301_v15)[_index201] = _1282_v0;
        let _1304_v17;
        let _nw195 = new _module.C0();
        _nw195.__ctor();
        _1304_v17 = _nw195;
        r1 = (((_1301_v15)[_module.__default.safeIndex(new BigNumber(34), new BigNumber((_1301_v15).length))]).multipliedBy(_1282_v0)).isEqualTo(_1282_v0);
      }
      let _1305_v18;
      let _init32 = function (_1306_i5) {
        return _module.__default.safeDivisionInt(_1306_i5, new BigNumber((_dafny.Seq.UnicodeFromString("kradmkld")).length));
      };
      let _nw196 = Array((new BigNumber(15)).toNumber());
      for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw196.length); _i0_32++) {
        _nw196[_i0_32] = _init32(new BigNumber(_i0_32));
      }
      _1305_v18 = _nw196;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1305_v18).length))) {
        let _1307_i4 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1307_i4)) && ((_1307_i4).isLessThan(new BigNumber((_1305_v18).length))))) {
          (_1305_v18)[(_1307_i4)] = _module.__default.safeDivisionInt(_1307_i4, _1282_v0);
        }
      }
      if ((_this).f10) {
        if ((_this).f10) {
          let _index202 = _module.__default.safeIndex(new BigNumber(160), new BigNumber((_1305_v18).length));
          (_1305_v18)[_index202] = (_1282_v0).multipliedBy(_1282_v0);
          let _1308_v19;
          _1308_v19 = _dafny.Seq.of((_this).f10);
          let _1309_v20;
          _1309_v20 = _dafny.Set.fromElements(!((_this).f10));
          let _1310_v21;
          _1310_v21 = _dafny.Seq.UnicodeFromString("cgysxnmd");
          let _1311_v22;
          let _nw197 = Array((new BigNumber(29)).toNumber());
          _nw197[(_dafny.ZERO).toNumber()] = true;
          _nw197[(_dafny.ONE).toNumber()] = (_this).fm18(_1309_v20, _1310_v21, globalState);
          _nw197[(new BigNumber(2)).toNumber()] = ((((_this).f9).contains((_this).f10)) ? (((_this).f9).get((_this).f10)) : ((_this).f10));
          _nw197[(new BigNumber(3)).toNumber()] = true;
          _nw197[(new BigNumber(4)).toNumber()] = false;
          _nw197[(new BigNumber(5)).toNumber()] = true;
          _nw197[(new BigNumber(6)).toNumber()] = true;
          _nw197[(new BigNumber(7)).toNumber()] = (_this).f10;
          _nw197[(new BigNumber(8)).toNumber()] = (_this).f10;
          _nw197[(new BigNumber(9)).toNumber()] = (_this).f10;
          _nw197[(new BigNumber(10)).toNumber()] = (_this).f10;
          _nw197[(new BigNumber(11)).toNumber()] = (_this).f10;
          _nw197[(new BigNumber(12)).toNumber()] = (_this).f10;
          _nw197[(new BigNumber(13)).toNumber()] = (_this).f10;
          _nw197[(new BigNumber(14)).toNumber()] = false;
          _nw197[(new BigNumber(15)).toNumber()] = (_this).f10;
          _nw197[(new BigNumber(16)).toNumber()] = (_this).f10;
          _nw197[(new BigNumber(17)).toNumber()] = (_this).f10;
          _nw197[(new BigNumber(18)).toNumber()] = false;
          _nw197[(new BigNumber(19)).toNumber()] = (_this).f10;
          _nw197[(new BigNumber(20)).toNumber()] = (_this).f10;
          _nw197[(new BigNumber(21)).toNumber()] = (_this).f10;
          _nw197[(new BigNumber(22)).toNumber()] = (_this).f10;
          _nw197[(new BigNumber(23)).toNumber()] = (_this).f10;
          _nw197[(new BigNumber(24)).toNumber()] = (_this).f10;
          _nw197[(new BigNumber(25)).toNumber()] = ((((_this).f9).contains((_this).f10)) ? (((_this).f9).get((_this).f10)) : ((_this).f10));
          _nw197[(new BigNumber(26)).toNumber()] = (_this).f10;
          _nw197[(new BigNumber(27)).toNumber()] = true;
          _nw197[(new BigNumber(28)).toNumber()] = true;
          _1311_v22 = _nw197;
          let _1312_v23;
          _1312_v23 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_1292_v6)[_module.__default.safeIndex(new BigNumber(293), new BigNumber((_1292_v6).length))]);
          let _1313_v24;
          let _nw198 = new _module.C2();
          _nw198.__ctor((_module.D1.create_DC4((_this).fm0(_1308_v19, true, _1282_v0, (_this).f10, globalState), !((_this).f10), new BigNumber((_1292_v6).length))).dtor_cf6, _1311_v22, (_this).f9, (_this).fm0(_1308_v19, (_this).f10, (((_1312_v23).contains(false)) ? ((_1312_v23).get(false)) : ((_1292_v6)[_module.__default.safeIndex(_1282_v0, new BigNumber((_1292_v6).length))])), (_this).f10, globalState));
          _1313_v24 = _nw198;
          let _index203 = _module.__default.safeIndex(new BigNumber(160), new BigNumber((_1305_v18).length));
          let _rhs199 = (_1282_v0).plus(_1282_v0);
          let _rhs200 = ((((_this).f10) === ((_this).f10)) ? (_1313_v24) : (_1313_v24));
          let _rhs201 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(7)), ((_1314_p0) => function (_1315_i6) {
            return _1314_p0;
          })(p0)), _1310_v21);
          let _lhs141 = _1305_v18;
          let _lhs142 = _module.__default.safeIndex(new BigNumber(160), new BigNumber((_1305_v18).length));
          _lhs141[_lhs142] = _rhs199;
          _1313_v24 = _rhs200;
          _1310_v21 = _rhs201;
          let _index204 = _module.__default.safeIndex(new BigNumber(160), new BigNumber((_1305_v18).length));
          let _index205 = _module.__default.safeIndex(new BigNumber(160), new BigNumber((_1305_v18).length));
          let _rhs202 = _1282_v0;
          let _rhs203 = (_this).f10;
          let _rhs204 = _module.__default.safeDivisionInt(_1282_v0, (_dafny.ZERO).minus((_this).fm19(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(347)), ((_1316_v0, _1317_v24, _1318_v18) => function (_1319_i7) {
            return new BigNumber((function () {
              let _coll49 = new _dafny.Map();
              for (const _compr_49 of (function () {
                let _coll50 = new _dafny.Map();
                for (const _compr_50 of _dafny.IntegerRange(new BigNumber(586), new BigNumber(962))) {
                  let _1320_v26 = _compr_50;
                  if (((new BigNumber(586)).isLessThanOrEqualTo(_1320_v26)) && ((_1320_v26).isLessThan(new BigNumber(962)))) {
                    _coll50.push([(_1320_v26).minus((_1318_v18)[_module.__default.safeIndex(new BigNumber(160), new BigNumber((_1318_v18).length))]),false]);
                  }
                }
                return _coll50;
              }()).Keys.Elements) {
                let _1321_v25 = _compr_49;
                if ((function () {
                  let _coll51 = new _dafny.Map();
                  for (const _compr_51 of _dafny.IntegerRange(new BigNumber(586), new BigNumber(962))) {
                    let _1320_v26 = _compr_51;
                    if (((new BigNumber(586)).isLessThanOrEqualTo(_1320_v26)) && ((_1320_v26).isLessThan(new BigNumber(962)))) {
                      _coll51.push([(_1320_v26).minus((_1318_v18)[_module.__default.safeIndex(new BigNumber(160), new BigNumber((_1318_v18).length))]),false]);
                    }
                  }
                  return _coll51;
                }()).contains(_1321_v25)) {
                  _coll49.push([_module.__default.safeModuloInt(_1321_v25, _1316_v0),(_1317_v24).f10]);
                }
              }
              return _coll49;
            }()).length);
          })(_1282_v0, _1313_v24, _1305_v18))).length), (_1305_v18)[_module.__default.safeIndex(new BigNumber(160), new BigNumber((_1305_v18).length))], globalState)));
          let _rhs205 = _1310_v21;
          let _rhs206 = _dafny.Seq.UnicodeFromString("cxs");
          let _lhs143 = _1305_v18;
          let _lhs144 = _module.__default.safeIndex(new BigNumber(160), new BigNumber((_1305_v18).length));
          let _lhs145 = globalState;
          let _lhs146 = _1305_v18;
          let _lhs147 = _module.__default.safeIndex(new BigNumber(160), new BigNumber((_1305_v18).length));
          _lhs143[_lhs144] = _rhs202;
          _lhs145.f8 = _rhs203;
          _lhs146[_lhs147] = _rhs204;
          _1310_v21 = _rhs205;
          _1310_v21 = _rhs206;
          let _1322_v27;
          let _nw199 = new _module.C3();
          _nw199.__ctor();
          _1322_v27 = _nw199;
          let _index206 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_1311_v22).length));
          (_1311_v22)[_index206] = !((_1313_v24).f10) || ((_1313_v24).f10);
          let _1323_v28;
          _1323_v28 = _module.D14.create_DC35(_1282_v0, (_1313_v24).f10, p0);
          let _index207 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((_1311_v22).length));
          (_1311_v22)[_index207] = (_1323_v28).dtor_cf63;
          let _1324_v29;
          let _nw200 = new _module.C4();
          _nw200.__ctor((_this).f9, (_1282_v0).isLessThan(_1282_v0));
          _1324_v29 = _nw200;
        } else {
          let _1325_v30;
          _1325_v30 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,true);
          let _1326_v31;
          _1326_v31 = _dafny.MultiSet.fromElements(_1282_v0);
          let _1327_v32;
          _1327_v32 = _dafny.Set.fromElements(_1282_v0, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1291_v5).cardinality()),_1282_v0)).length));
          _1325_v30 = (_1325_v30).update((_this).f10, (_dafny.MultiSet.fromElements(new BigNumber((_1327_v32).length), _1282_v0)).IsProperSubsetOf(_1326_v31));
          let _1328_v33;
          _1328_v33 = _dafny.Seq.of((_this).f10, (_this).f10);
          let _1329_v34;
          _1329_v34 = _dafny.Map.Empty.slice().updateUnsafe(_1328_v33,new BigNumber(953));
          let _1330_v35;
          _1330_v35 = _dafny.Map.Empty.slice().updateUnsafe((((_1329_v34).contains(_1328_v33)) ? ((_1329_v34).get(_1328_v33)) : (_1282_v0)),(_this).f10);
          let _1331_v36;
          _1331_v36 = _dafny.Set.fromElements((_module.D14.create_DC35(new BigNumber((_1330_v35).length), (_this).f10, p0)).dtor_cf63);
          (globalState).f8 = !(!((_this).fm18(_1331_v36, _dafny.Seq.UnicodeFromString("m"), globalState)));
          let _1332_v37;
          _1332_v37 = _dafny.Seq.UnicodeFromString("x");
          _1332_v37 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-22)), function (_1333_i8) {
            return new _dafny.CodePoint('n'.codePointAt(0));
          }), _1332_v37);
          let _1334_v38;
          let _nw201 = Array((new BigNumber(3)).toNumber()).fill(false);
          _1334_v38 = _nw201;
          let _index208 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_1334_v38).length));
          (_1334_v38)[_index208] = (_this).fm18(_1331_v36, _1332_v37, globalState);
          let _index209 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_1334_v38).length));
          (_1334_v38)[_index209] = (_this).f10;
          let _index210 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_1334_v38).length));
          (_1334_v38)[_index210] = (_module.__default.fm48((_this).f10, globalState)).contains((_this).fm18(_1331_v36, _1332_v37, globalState));
        }
        (globalState).f8 = (_this).f10;
        let _1335_v39;
        let _nw202 = Array((new BigNumber(5)).toNumber());
        _nw202[(_dafny.ZERO).toNumber()] = (((_this).f10) ? (_1305_v18) : (_1305_v18));
        _nw202[(_dafny.ONE).toNumber()] = _1305_v18;
        _nw202[(new BigNumber(2)).toNumber()] = _1305_v18;
        _nw202[(new BigNumber(3)).toNumber()] = _1305_v18;
        _nw202[(new BigNumber(4)).toNumber()] = _1305_v18;
        _1335_v39 = _nw202;
        let _index211 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_1335_v39).length));
        (_1335_v39)[_index211] = _1305_v18;
        let _index212 = _module.__default.safeIndex(new BigNumber(495), new BigNumber((_1335_v39).length));
        (_1335_v39)[_index212] = _1305_v18;
        let _1336_v40;
        let _nw203 = Array((new BigNumber(11)).toNumber()).fill(false);
        _1336_v40 = _nw203;
        let _1337_v41;
        _1337_v41 = _dafny.Set.fromElements((_this).f10);
        let _1338_v42;
        _1338_v42 = _dafny.Seq.UnicodeFromString("ksbootweo");
        let _index213 = _module.__default.safeIndex(new BigNumber(946), new BigNumber((_1336_v40).length));
        (_1336_v40)[_index213] = ((_this).fm18(_1337_v41, _1338_v42, globalState)) === ((_this).f10);
        let _1339_v43;
        _1339_v43 = _dafny.Seq.of(false, (_this).f10);
        let _1340_v44;
        _1340_v44 = _dafny.MultiSet.fromElements(_1282_v0, new BigNumber(55), _1282_v0);
        let _1341_v45;
        _1341_v45 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_1282_v0, _1282_v0)).cardinality()),p0);
        let _index214 = _module.__default.safeIndex(new BigNumber(946), new BigNumber((_1336_v40).length));
        (_1336_v40)[_index214] = (new BigNumber((_1338_v42).length)).isLessThanOrEqualTo((_module.D6.create_DC17((_this).fm0(_1339_v43, (_this).f10, new BigNumber((_1340_v44).cardinality()), (_this).f10, globalState), true, new BigNumber((_1341_v45).length), _1282_v0, _1282_v0)).dtor_cf31);
        if ((_module.D0.create_DC1(_1282_v0, ((((_this).f9).contains((_1336_v40)[_module.__default.safeIndex(new BigNumber(946), new BigNumber((_1336_v40).length))])) ? (((_this).f9).get((_1336_v40)[_module.__default.safeIndex(new BigNumber(946), new BigNumber((_1336_v40).length))])) : ((_1336_v40)[_module.__default.safeIndex(new BigNumber(946), new BigNumber((_1336_v40).length))])), _1282_v0)).dtor_cf2) {
          (globalState).f8 = !((_1336_v40)[_module.__default.safeIndex(new BigNumber(946), new BigNumber((_1336_v40).length))]);
          let _1342_v46;
          let _init33 = ((_1343_v40, _1344_v5) => function (_1345_i9) {
            return (((_1343_v40)[_module.__default.safeIndex(new BigNumber(946), new BigNumber((_1343_v40).length))]) ? (_1344_v5) : (_dafny.MultiSet.fromElements((_1343_v40)[_module.__default.safeIndex(new BigNumber(946), new BigNumber((_1343_v40).length))], (_1343_v40)[_module.__default.safeIndex(new BigNumber(946), new BigNumber((_1343_v40).length))])));
          })(_1336_v40, _1291_v5);
          let _nw204 = Array((new BigNumber(12)).toNumber());
          for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw204.length); _i0_33++) {
            _nw204[_i0_33] = _init33(new BigNumber(_i0_33));
          }
          _1342_v46 = _nw204;
          let _1346_v47;
          let _init34 = ((_1347_v43) => function (_1348_i10) {
            return _1347_v43;
          })(_1339_v43);
          let _nw205 = Array((new BigNumber(27)).toNumber());
          for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw205.length); _i0_34++) {
            _nw205[_i0_34] = _init34(new BigNumber(_i0_34));
          }
          _1346_v47 = _nw205;
          let _index215 = _module.__default.safeIndex(new BigNumber(134), new BigNumber((_1346_v47).length));
          (_1346_v47)[_index215] = _dafny.Seq.Concat(_1339_v43, _dafny.Seq.update(_1339_v43, _module.__default.safeIndex(_1282_v0, new BigNumber((_1339_v43).length)), (_1336_v40)[_module.__default.safeIndex(new BigNumber(946), new BigNumber((_1336_v40).length))]));
          let _1349_v48;
          let _nw206 = Array((new BigNumber(29)).toNumber()).fill([]);
          _1349_v48 = _nw206;
          let _index216 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_1349_v48).length));
          (_1349_v48)[_index216] = _1336_v40;
          let _index217 = _module.__default.safeIndex(new BigNumber(134), new BigNumber((_1346_v47).length));
          let _index218 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_1349_v48).length));
          let _rhs207 = _1342_v46;
          let _rhs208 = _1339_v43;
          let _rhs209 = _1336_v40;
          let _rhs210 = ((_this).f10) || ((_this).f10);
          let _rhs211 = _1282_v0;
          let _lhs148 = _1346_v47;
          let _lhs149 = _module.__default.safeIndex(new BigNumber(134), new BigNumber((_1346_v47).length));
          let _lhs150 = _1349_v48;
          let _lhs151 = _module.__default.safeIndex(new BigNumber(341), new BigNumber((_1349_v48).length));
          let _lhs152 = globalState;
          _1342_v46 = _rhs207;
          _lhs148[_lhs149] = _rhs208;
          _lhs150[_lhs151] = _rhs209;
          _lhs152.f8 = _rhs210;
          _1282_v0 = _rhs211;
          _1282_v0 = (new BigNumber(345)).multipliedBy(_1282_v0);
          let _1350_v49;
          _1350_v49 = _dafny.Set.fromElements(_1282_v0, _1282_v0);
          let _1351_v50;
          _1351_v50 = _module.D4.create_DC11(_1282_v0, _1350_v49, (_1336_v40)[_module.__default.safeIndex(new BigNumber(946), new BigNumber((_1336_v40).length))]);
          let _1352_v51;
          _1352_v51 = _dafny.Seq.of(_1351_v50, _1351_v50, _1351_v50);
          let _1353_v52;
          _1353_v52 = _dafny.Seq.Concat(_1352_v51, _1352_v51);
          _1353_v52 = _dafny.Seq.of(_1351_v50, _1351_v50, _1351_v50, _1351_v50, _1351_v50);
          let _arr2 = (_1335_v39)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_1335_v39).length))];
          let _index219 = _module.__default.safeIndex(new BigNumber(998), new BigNumber(((_1335_v39)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_1335_v39).length))]).length));
          _arr2[_index219] = _module.__default.safeModuloInt(_1282_v0, new BigNumber((_1338_v42).length));
          let _arr3 = (_1335_v39)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_1335_v39).length))];
          let _index220 = _module.__default.safeIndex(new BigNumber(998), new BigNumber(((_1335_v39)[_module.__default.safeIndex(new BigNumber(495), new BigNumber((_1335_v39).length))]).length));
          _arr3[_index220] = ((new BigNumber(((_this).f9).length)).multipliedBy((_dafny.ZERO).minus(new BigNumber((_1338_v42).length)))).plus(_1282_v0);
        } else {
          let _1354_v53;
          _1354_v53 = _dafny.Seq.of(_dafny.Seq.update(_1292_v6, _module.__default.safeIndex(_1282_v0, new BigNumber((_1292_v6).length)), _1282_v0));
          let _1355_v54;
          let _nw207 = new _module.C2();
          _nw207.__ctor(_dafny.Seq.contains(_1354_v53, _dafny.Seq.Create(_module.__default.abs(new BigNumber(263)), ((_1356_v0) => function (_1357_i11) {
            return _1356_v0;
          })(_1282_v0))), _1336_v40, ((_this).f9).Merge((_this).f9), (_this).f10);
          _1355_v54 = _nw207;
          let _rhs212 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_1282_v0));
          let _rhs213 = _1282_v0;
          let _rhs214 = _1282_v0;
          let _rhs215 = _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("tm"), _1338_v42);
          let _rhs216 = false;
          let _lhs153 = globalState;
          let _lhs154 = _1355_v54;
          _1282_v0 = _rhs212;
          _1282_v0 = _rhs213;
          _1282_v0 = _rhs214;
          _lhs153.f8 = _rhs215;
          _lhs154.f14 = _rhs216;
          let _index221 = _module.__default.safeIndex(new BigNumber(469), new BigNumber((_1305_v18).length));
          (_1305_v18)[_index221] = _1282_v0;
          let _index222 = _module.__default.safeIndex(new BigNumber(469), new BigNumber((_1305_v18).length));
          (_1305_v18)[_index222] = _1282_v0;
          let _1358_v55;
          _1358_v55 = _dafny.Seq.of(_module.__default.fm50((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements((_1305_v18)[_module.__default.safeIndex(new BigNumber(469), new BigNumber((_1305_v18).length))])).cardinality())), (_1305_v18)[_module.__default.safeIndex(new BigNumber(469), new BigNumber((_1305_v18).length))], globalState));
          _1358_v55 = _1358_v55;
          (globalState).f8 = _1355_v54.f14;
        }
      } else {
        let _1359_v56;
        _1359_v56 = _dafny.Seq.UnicodeFromString("ty");
        let _1360_v57;
        _1360_v57 = _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("yof"), _dafny.Seq.UnicodeFromString("uhu")), _dafny.Seq.Concat(_1359_v56, _1359_v56), _1359_v56, _dafny.Seq.Concat(_1359_v56, _1359_v56), _1359_v56);
        _1360_v57 = _dafny.Seq.of(_1359_v56, _dafny.Seq.UnicodeFromString("y"));
        let _1361_v58;
        _1361_v58 = _module.D14.create_DC36(false, _1359_v56, _1359_v56);
        let _1362_v59;
        _1362_v59 = _dafny.MultiSet.fromElements(_1361_v58);
        _1282_v0 = new BigNumber(((_1362_v59).update(_1361_v58, _module.__default.abs(_1282_v0))).cardinality());
        let _1363_v60;
        _1363_v60 = _module.D1.create_DC3(_1359_v56);
        let _1364_v61;
        _1364_v61 = _module.D1.create_DC5(_1363_v60);
        let _source18 = _1364_v61;
        if (_source18.is_DC3) {
          let _1365___mcc_h0 = (_source18).cf5;
          let _1366_cf5 = _1365___mcc_h0;
          let _1367_v62;
          _1367_v62 = _dafny.Seq.of(_1305_v18);
          let _1368_v63;
          _1368_v63 = _dafny.Set.fromElements(_1282_v0, _1282_v0);
          let _1369_v64;
          _1369_v64 = _module.D4.create_DC11(_1282_v0, _1368_v63, (_this).f10);
          let _1370_v65;
          _1370_v65 = _dafny.Map.Empty.slice().updateUnsafe(_1367_v62,((_1369_v64).dtor_cf16).minus(_1282_v0));
          _1370_v65 = (_1370_v65).update(_1367_v62, (_dafny.ZERO).minus(_1282_v0));
          let _1371_v66;
          _1371_v66 = _dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.of((_this).f10, (_this).f10));
          let _1372_v67;
          _1372_v67 = _dafny.Seq.of(false);
          _1371_v66 = (_1371_v66).update((_this).f10, _1372_v67);
          _1282_v0 = _module.__default.safeDivisionInt(new BigNumber((_1372_v67).length), (_this).fm1((_this).f9, globalState));
          let _1373_v68;
          let _init35 = function (_1374_i12) {
            return new _dafny.CodePoint('r'.codePointAt(0));
          };
          let _nw208 = Array((new BigNumber(5)).toNumber());
          for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw208.length); _i0_35++) {
            _nw208[_i0_35] = _init35(new BigNumber(_i0_35));
          }
          _1373_v68 = _nw208;
          let _index223 = _module.__default.safeIndex(new BigNumber(45), new BigNumber((_1373_v68).length));
          (_1373_v68)[_index223] = p0;
          let _index224 = _module.__default.safeIndex(new BigNumber(45), new BigNumber((_1373_v68).length));
          (_1373_v68)[_index224] = p0;
        } else if (_source18.is_DC4) {
          let _1375___mcc_h1 = (_source18).cf6;
          let _1376___mcc_h2 = (_source18).cf7;
          let _1377___mcc_h3 = (_source18).cf8;
          let _1378_cf8 = _1377___mcc_h3;
          let _1379_cf7 = _1376___mcc_h2;
          let _1380_cf6 = _1375___mcc_h1;
          let _1381_v69;
          _1381_v69 = _dafny.Set.fromElements(new BigNumber(976));
          let _1382_v70;
          _1382_v70 = _module.D4.create_DC11(_1282_v0, _1381_v69, (_this).f10);
          let _1383_v71;
          _1383_v71 = _dafny.Seq.of(_1382_v70);
          let _1384_v72;
          _1384_v72 = _dafny.Map.Empty.slice().updateUnsafe(_1383_v71,_1282_v0);
          let _1385_v73;
          _1385_v73 = _dafny.Seq.of(_1384_v72, _dafny.Map.Empty.slice().updateUnsafe(_1383_v71,(_dafny.ZERO).minus(_1282_v0)), _1384_v72, _1384_v72);
          let _1386_v74;
          _1386_v74 = _module.D14.create_DC34((_1385_v73)[_module.__default.safeIndex(new BigNumber(174), new BigNumber((_1385_v73).length))]);
          _1386_v74 = _module.__default.fm51(globalState);
          let _1387_v76;
          let _init36 = ((_1388_cf7, _1389_v0) => function (_1390_i13) {
            return function () {
              let _coll52 = new _dafny.Set();
              for (const _compr_52 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_1388_cf7),_1389_v0)).Keys.Elements) {
                let _1391_v75 = _compr_52;
                if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_1388_cf7),_1389_v0)).contains(_1391_v75)) {
                  _coll52.add(_1391_v75);
                }
              }
              return _coll52;
            }();
          })(_1379_cf7, _1282_v0);
          let _nw209 = Array((new BigNumber(19)).toNumber());
          for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw209.length); _i0_36++) {
            _nw209[_i0_36] = _init36(new BigNumber(_i0_36));
          }
          _1387_v76 = _nw209;
          let _1392_v77;
          _1392_v77 = _dafny.Seq.of(_1379_cf7);
          let _1393_v78;
          _1393_v78 = _dafny.Set.fromElements(_1392_v77);
          let _index225 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_1387_v76).length));
          (_1387_v76)[_index225] = _1393_v78;
          let _index226 = _module.__default.safeIndex(new BigNumber(831), new BigNumber((_1387_v76).length));
          (_1387_v76)[_index226] = _1393_v78;
          let _1394_v79;
          _1394_v79 = _dafny.Seq.of(_1305_v18, _1305_v18, _1305_v18, _1305_v18, _1305_v18);
          _1305_v18 = (_1394_v79)[_module.__default.safeIndex((_this).fm1(((_this).f9).update(false, !(((((_this).f9).contains(false)) ? (((_this).f9).get(false)) : (true)))), globalState), new BigNumber((_1394_v79).length))];
          (globalState).f8 = _1379_cf7;
        } else if (_source18.is_DC2) {
          let _1395___mcc_h4 = (_source18).cf4;
          let _1396_cf4 = _1395___mcc_h4;
          let _rhs217 = (_this).f10;
          let _rhs218 = (_this).f10;
          let _lhs155 = globalState;
          _lhs155.f8 = _rhs217;
          r1 = _rhs218;
          let _1397_v80;
          let _nw210 = new _module.C0();
          _nw210.__ctor();
          _1397_v80 = _nw210;
          let _index227 = _module.__default.safeIndex(new BigNumber(683), new BigNumber((_1305_v18).length));
          (_1305_v18)[_index227] = (_dafny.ZERO).minus((_this).fm19(_1282_v0, (_dafny.ZERO).minus(_1282_v0), globalState));
          let _index228 = _module.__default.safeIndex(new BigNumber(683), new BigNumber((_1305_v18).length));
          (_1305_v18)[_index228] = new BigNumber(-366);
          let _1398_v81;
          _1398_v81 = new _dafny.CodePoint('o'.codePointAt(0));
          _1398_v81 = _1398_v81;
        } else {
          let _1399___mcc_h5 = (_source18).cf9;
          let _1400_cf9 = _1399___mcc_h5;
          let _1401_v82;
          _1401_v82 = _dafny.Set.fromElements(_1305_v18);
          _1401_v82 = _1401_v82;
          (globalState).f8 = (_this).f10;
          _1282_v0 = _1282_v0;
          let _1402_v83;
          let _nw211 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1402_v83 = _nw211;
          let _index229 = _module.__default.safeIndex(new BigNumber(924), new BigNumber((_1402_v83).length));
          (_1402_v83)[_index229] = _1359_v56;
          let _index230 = _module.__default.safeIndex(new BigNumber(924), new BigNumber((_1402_v83).length));
          (_1402_v83)[_index230] = _1359_v56;
        }
        r1 = (_this).f10;
        let _index231 = _module.__default.safeIndex(new BigNumber(74), new BigNumber((_1305_v18).length));
        (_1305_v18)[_index231] = _module.__default.safeDivisionInt(new BigNumber((_1292_v6).length), _1282_v0);
        let _index232 = _module.__default.safeIndex(new BigNumber(74), new BigNumber((_1305_v18).length));
        (_1305_v18)[_index232] = new BigNumber(106);
      }
      let _1403_v84;
      _1403_v84 = _dafny.Seq.UnicodeFromString("dh");
      let _1404_v85;
      _1404_v85 = _dafny.Map.Empty.slice().updateUnsafe(_1305_v18,_dafny.Seq.Concat(_1403_v84, _1403_v84));
      r0 = _1404_v85;
      r1 = (_this).f10;
      return [r0, r1];
    }
    m4(p0, p1, globalState) {
      let _this = this;
      let r0 = [];
      let r1 = _dafny.Seq.of();
      let r2 = _dafny.Set.Empty;
      let _1405_v0;
      _1405_v0 = new BigNumber(49);
      let _hi3 = _1405_v0;
      for (let _1406_i0 = _module.__default.safeDivisionInt(_1405_v0, new BigNumber(146)); _1406_i0.isLessThan(_hi3); _1406_i0 = _1406_i0.plus(_dafny.ONE)) {
        let _1407_v1;
        _1407_v1 = _dafny.Seq.UnicodeFromString("itssvdb");
        _1407_v1 = _dafny.Seq.UnicodeFromString("mfa");
        let _source19 = _module.D6.create_DC17(p1, (_this).f10, _1406_i0, _1405_v0, _module.__default.safeDivisionInt(new BigNumber(576), _1406_i0));
        if (_source19.is_DC17) {
          let _1408___mcc_h0 = (_source19).cf29;
          let _1409___mcc_h1 = (_source19).cf30;
          let _1410___mcc_h2 = (_source19).cf31;
          let _1411___mcc_h3 = (_source19).cf32;
          let _1412___mcc_h4 = (_source19).cf33;
          let _1413_cf33 = _1412___mcc_h4;
          let _1414_cf32 = _1411___mcc_h3;
          let _1415_cf31 = _1410___mcc_h2;
          let _1416_cf30 = _1409___mcc_h1;
          let _1417_cf29 = _1408___mcc_h0;
          _1415_cf31 = (_1406_i0).plus(_1405_v0);
          let _1418_v2;
          _1418_v2 = new _dafny.CodePoint('t'.codePointAt(0));
          _1418_v2 = _1418_v2;
          let _1419_v3;
          _1419_v3 = _dafny.Seq.of(_1416_cf30);
          let _1420_v4;
          _1420_v4 = _module.D8.create_DC23(_dafny.Seq.of((_this).f10, p0, (_this).fm0(_1419_v3, p1, new BigNumber(313), false, globalState)));
          _1420_v4 = _1420_v4;
          let _1421_v5;
          _1421_v5 = _dafny.Set.fromElements((_this).f10);
          let _1422_v6;
          let _nw212 = Array((new BigNumber(4)).toNumber());
          _nw212[(_dafny.ZERO).toNumber()] = new BigNumber((_1421_v5).length);
          _nw212[(_dafny.ONE).toNumber()] = new BigNumber(-973);
          _nw212[(new BigNumber(2)).toNumber()] = _1414_cf32;
          _nw212[(new BigNumber(3)).toNumber()] = _1415_cf31;
          _1422_v6 = _nw212;
          let _1423_v7;
          _1423_v7 = _dafny.Set.fromElements(_1422_v6, _1422_v6);
          _1423_v7 = _1423_v7;
        } else if (_source19.is_DC18) {
          let _1424___mcc_h5 = (_source19).cf34;
          let _1425___mcc_h6 = (_source19).cf35;
          let _1426___mcc_h7 = (_source19).cf36;
          let _1427_cf36 = _1426___mcc_h7;
          let _1428_cf35 = _1425___mcc_h6;
          let _1429_cf34 = _1424___mcc_h5;
          let _1430_v8;
          _1430_v8 = _module.D14.create_DC35(_1427_cf36, !(true), _1429_cf34);
          _1427_cf36 = ((_1430_v8).dtor_cf62).minus(_1428_cf35);
          let _1431_v9;
          let _init37 = ((_1432_p1) => function (_1433_i1) {
            return _1432_p1;
          })(p1);
          let _nw213 = Array((new BigNumber(10)).toNumber());
          for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw213.length); _i0_37++) {
            _nw213[_i0_37] = _init37(new BigNumber(_i0_37));
          }
          _1431_v9 = _nw213;
          let _1434_v10;
          let _nw214 = new _module.C2();
          _nw214.__ctor(p1, _1431_v9, (_this).f9, p0);
          _1434_v10 = _nw214;
          let _1435_v11;
          let _1436_v12;
          let _out61;
          let _out62;
          let _outcollector18 = (_1434_v10).m1(_1429_cf34, globalState);
          _out61 = _outcollector18[0];
          _out62 = _outcollector18[1];
          _1435_v11 = _out61;
          _1436_v12 = _out62;
          let _1437_v13;
          _1437_v13 = _module.D0.create_DC1(_1406_i0, !(!(p0)), _1428_cf35);
          let _rhs219 = ((_1427_cf36).minus(new BigNumber((_dafny.Seq.of(_1406_i0)).length))).multipliedBy(_1406_i0);
          let _rhs220 = (_1437_v13).dtor_cf2;
          let _rhs221 = p1;
          let _rhs222 = (_this).f10;
          let _lhs156 = globalState;
          let _lhs157 = globalState;
          let _lhs158 = globalState;
          _1427_cf36 = _rhs219;
          _lhs156.f8 = _rhs220;
          _lhs157.f8 = _rhs221;
          _lhs158.f8 = _rhs222;
        } else if (_source19.is_DC19) {
          let _1438___mcc_h8 = (_source19).cf37;
          let _1439___mcc_h9 = (_source19).cf38;
          let _1440_cf38 = _1439___mcc_h9;
          let _1441_cf37 = _1438___mcc_h8;
          _1405_v0 = _1406_i0;
          let _1442_v14;
          _1442_v14 = _dafny.Seq.of((_this).f10, !((_this).f10));
          let _1443_v15;
          _1443_v15 = _dafny.Set.fromElements(_1442_v14);
          let _1444_v16;
          _1444_v16 = _dafny.MultiSet.fromElements(p1);
          let _1445_v17;
          _1445_v17 = _dafny.Seq.of(_1406_i0);
          let _1446_v18;
          _1446_v18 = _dafny.Set.fromElements((((_1444_v16).contains(_1441_cf37)) ? ((_1444_v16).get(_1441_cf37)) : (new BigNumber((_1445_v17).length))));
          _1441_cf37 = (_1446_v18).IsProperSubsetOf(_dafny.Set.fromElements(new BigNumber((_1443_v15).length)));
          let _1447_v19;
          _1447_v19 = _module.D8.create_DC23(_1442_v14);
          let _pat_let_tv26 = _1442_v14;
          let _pat_let_tv27 = _1442_v14;
          let _1448_v20;
          let _nw215 = Array((new BigNumber(28)).toNumber());
          _nw215[(_dafny.ZERO).toNumber()] = _1447_v19;
          _nw215[(_dafny.ONE).toNumber()] = _module.__default.fm52(!((_this).f10), globalState);
          _nw215[(new BigNumber(2)).toNumber()] = _1447_v19;
          _nw215[(new BigNumber(3)).toNumber()] = _module.D8.create_DC23(_1442_v14);
          _nw215[(new BigNumber(4)).toNumber()] = _1447_v19;
          _nw215[(new BigNumber(5)).toNumber()] = _1447_v19;
          _nw215[(new BigNumber(6)).toNumber()] = _1447_v19;
          _nw215[(new BigNumber(7)).toNumber()] = _1447_v19;
          _nw215[(new BigNumber(8)).toNumber()] = function (_pat_let27_0) {
            return function (_1449_dt__update__tmp_h0) {
              return function (_pat_let28_0) {
                return function (_1450_dt__update_hcf41_h0) {
                  return _module.D8.create_DC23(_1450_dt__update_hcf41_h0);
                }(_pat_let28_0);
              }(_pat_let_tv26);
            }(_pat_let27_0);
          }(_1447_v19);
          _nw215[(new BigNumber(9)).toNumber()] = _module.__default.fm52(_1441_cf37, globalState);
          _nw215[(new BigNumber(10)).toNumber()] = _1447_v19;
          _nw215[(new BigNumber(11)).toNumber()] = _1447_v19;
          _nw215[(new BigNumber(12)).toNumber()] = _1447_v19;
          _nw215[(new BigNumber(13)).toNumber()] = _1447_v19;
          _nw215[(new BigNumber(14)).toNumber()] = _1447_v19;
          _nw215[(new BigNumber(15)).toNumber()] = _1447_v19;
          _nw215[(new BigNumber(16)).toNumber()] = _1447_v19;
          _nw215[(new BigNumber(17)).toNumber()] = _1447_v19;
          _nw215[(new BigNumber(18)).toNumber()] = _1447_v19;
          _nw215[(new BigNumber(19)).toNumber()] = _1447_v19;
          _nw215[(new BigNumber(20)).toNumber()] = _1447_v19;
          _nw215[(new BigNumber(21)).toNumber()] = function (_pat_let29_0) {
            return function (_1451_dt__update__tmp_h1) {
              return function (_pat_let30_0) {
                return function (_1452_dt__update_hcf41_h1) {
                  return _module.D8.create_DC23(_1452_dt__update_hcf41_h1);
                }(_pat_let30_0);
              }(_pat_let_tv27);
            }(_pat_let29_0);
          }(_1447_v19);
          _nw215[(new BigNumber(22)).toNumber()] = _module.D8.create_DC23(_1442_v14);
          _nw215[(new BigNumber(23)).toNumber()] = _1447_v19;
          _nw215[(new BigNumber(24)).toNumber()] = _module.D8.create_DC23(_1442_v14);
          _nw215[(new BigNumber(25)).toNumber()] = _module.D8.create_DC23(_1442_v14);
          _nw215[(new BigNumber(26)).toNumber()] = _1447_v19;
          _nw215[(new BigNumber(27)).toNumber()] = _1447_v19;
          _1448_v20 = _nw215;
          let _1453_v21;
          _1453_v21 = _dafny.Seq.of(_1448_v20);
          let _rhs223 = _1453_v21;
          let _rhs224 = p1;
          let _rhs225 = _1406_i0;
          let _lhs159 = globalState;
          _1453_v21 = _rhs223;
          _lhs159.f8 = _rhs224;
          _1405_v0 = _rhs225;
          let _1454_v22;
          _1454_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1440_cf38,_1440_cf38);
          let _1455_v23;
          _1455_v23 = _module.D11.create_DC28(_1405_v0, _1441_cf37, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1446_v18).length),_1454_v22), (_this).f10, _dafny.Seq.update(_1407_v1, _module.__default.safeIndex(_1406_i0, new BigNumber((_1407_v1).length)), _1440_cf38));
          (globalState).f8 = !((_1405_v0).isLessThan(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1407_v1).length),_1440_cf38)).length))) || ((_1455_v23).dtor_cf51);
        } else if (_source19.is_DC16) {
          let _1456___mcc_h10 = (_source19).cf28;
          let _1457_cf28 = _1456___mcc_h10;
          let _1458_v24;
          let _nw216 = new _module.C3();
          _nw216.__ctor();
          _1458_v24 = _nw216;
          let _1459_v25;
          _1459_v25 = new _dafny.CodePoint('x'.codePointAt(0));
          let _1460_v26;
          _1460_v26 = _dafny.MultiSet.fromElements(_1406_i0);
          let _rhs226 = (_this).f10;
          let _rhs227 = (((_1460_v26).IsSubsetOf(_dafny.MultiSet.fromElements(_1405_v0, _1405_v0))) ? (_1458_v24) : (_1458_v24));
          let _rhs228 = p1;
          let _rhs229 = _1459_v25;
          let _lhs160 = globalState;
          let _lhs161 = globalState;
          _lhs160.f8 = _rhs226;
          _1458_v24 = _rhs227;
          _lhs161.f8 = _rhs228;
          _1459_v25 = _rhs229;
          (globalState).f8 = !((_1405_v0).isLessThan(new BigNumber((_1407_v1).length))) || (p1);
          let _1461_v27;
          _1461_v27 = _module.D7.create_DC21(_dafny.Set.fromElements(p1, false));
          let _1462_v28;
          _1462_v28 = _dafny.Set.fromElements((_this).f10, (_this).f10);
          _1461_v27 = _module.D7.create_DC21(_1462_v28);
          let _1463_v29;
          _1463_v29 = _dafny.Seq.of(_1460_v26, _dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-567)), ((_1464_v0) => function (_1465_i2) {
            return (_dafny.ZERO).minus(_1464_v0);
          })(_1405_v0))));
          let _1466_v30;
          _1466_v30 = _dafny.MultiSet.fromElements(p1, p0);
          let _1467_v31;
          _1467_v31 = _dafny.Set.fromElements(new BigNumber(628), new BigNumber((_1466_v30).cardinality()));
          let _1468_v32;
          _1468_v32 = _module.D4.create_DC11(new BigNumber(914), _1467_v31, (_this).f10);
          let _1469_v33;
          _1469_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1405_v0,p1);
          let _pat_let_tv28 = _1469_v33;
          let _pat_let_tv29 = _1469_v33;
          let _pat_let_tv30 = _1468_v32;
          _1405_v0 = (function (_pat_let31_0) {
            return function (_1470_dt__update__tmp_h2) {
              return function (_pat_let32_0) {
                return function (_1471_dt__update_hcf7_h0) {
                  return function (_pat_let33_0) {
                    return function (_1472_dt__update_hcf8_h0) {
                      return _module.D1.create_DC4((_1470_dt__update__tmp_h2).dtor_cf6, _1471_dt__update_hcf7_h0, _1472_dt__update_hcf8_h0);
                    }(_pat_let33_0);
                  }((_pat_let_tv30).dtor_cf16);
                }(_pat_let32_0);
              }((((_pat_let_tv29).contains(_1406_i0)) ? ((_pat_let_tv28).get(_1406_i0)) : ((_this).f10)));
            }(_pat_let31_0);
          }(_module.D1.create_DC4(p1, p1, new BigNumber((_1463_v29).length)))).dtor_cf8;
        } else {
          let _1473___mcc_h11 = (_source19).cf39;
          let _1474_cf39 = _1473___mcc_h11;
          let _1475_v34;
          _1475_v34 = _dafny.Seq.of(p1);
          (globalState).f8 = ((_this).fm0(_1475_v34, (_1475_v34)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("axewyqwhx")).length), new BigNumber((_1475_v34).length))], _1405_v0, (_this).f10, globalState)) && (p0);
          let _1476_v35;
          _1476_v35 = _dafny.Seq.of(_1405_v0, _1406_i0);
          let _1477_v36;
          _1477_v36 = _dafny.MultiSet.fromElements(p1, (_this).f10, (_this).f10);
          let _1478_v37;
          _1478_v37 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1477_v36).cardinality()),_1406_i0);
          let _1479_v38;
          let _nw217 = Array((new BigNumber(12)).toNumber());
          _nw217[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_1406_i0));
          _nw217[(_dafny.ONE).toNumber()] = (_this).fm1((_this).f9, globalState);
          _nw217[(new BigNumber(2)).toNumber()] = (new BigNumber((_1475_v34).length)).multipliedBy(new BigNumber((_dafny.MultiSet.FromArray(_1476_v35)).cardinality()));
          _nw217[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus((_this).fm19(_1405_v0, _1406_i0, globalState));
          _nw217[(new BigNumber(4)).toNumber()] = _1406_i0;
          _nw217[(new BigNumber(5)).toNumber()] = _1405_v0;
          _nw217[(new BigNumber(6)).toNumber()] = new BigNumber(((_1478_v37).Merge(_1478_v37)).length);
          _nw217[(new BigNumber(7)).toNumber()] = _1405_v0;
          _nw217[(new BigNumber(8)).toNumber()] = _1405_v0;
          _nw217[(new BigNumber(9)).toNumber()] = _1406_i0;
          _nw217[(new BigNumber(10)).toNumber()] = new BigNumber(-769);
          _nw217[(new BigNumber(11)).toNumber()] = (new BigNumber(197)).multipliedBy(_1406_i0);
          _1479_v38 = _nw217;
          let _index233 = _module.__default.safeIndex(new BigNumber(494), new BigNumber((_1479_v38).length));
          (_1479_v38)[_index233] = _module.__default.safeModuloInt(new BigNumber(543), new BigNumber(((_1478_v37).update(_1405_v0, _1406_i0)).length));
          let _1480_v39;
          _1480_v39 = _module.D5.create_DC14(_1405_v0, p1);
          let _1481_v40;
          _1481_v40 = _dafny.Seq.of(_1480_v39);
          let _index234 = _module.__default.safeIndex(new BigNumber(494), new BigNumber((_1479_v38).length));
          let _rhs230 = p0;
          let _rhs231 = (_this).f10;
          let _rhs232 = _module.__default.fm33(p0, _dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(868)), ((_1482_i0, _1483_p0) => function (_1484_i3) {
            return _module.D5.create_DC14(_1482_i0, _1483_p0);
          })(_1406_i0, p0)), _1481_v40), ((((_1477_v36).contains((_this).f10)) ? ((_1477_v36).get((_this).f10)) : (_1405_v0))).multipliedBy((_dafny.ZERO).minus(_1406_i0)), (_this).fm0(_1475_v34, p0, _1406_i0, (_this).f10, globalState), globalState);
          let _lhs162 = globalState;
          let _lhs163 = globalState;
          let _lhs164 = _1479_v38;
          let _lhs165 = _module.__default.safeIndex(new BigNumber(494), new BigNumber((_1479_v38).length));
          _lhs162.f8 = _rhs230;
          _lhs163.f8 = _rhs231;
          _lhs164[_lhs165] = _rhs232;
          let _1485_v41;
          _1485_v41 = new _dafny.CodePoint('o'.codePointAt(0));
          _1485_v41 = _1485_v41;
          let _1486_v42;
          let _nw218 = Array((new BigNumber(11)).toNumber());
          _nw218[(_dafny.ZERO).toNumber()] = _1407_v1;
          _nw218[(_dafny.ONE).toNumber()] = _dafny.Seq.of(_1485_v41, _1485_v41, _1485_v41);
          _nw218[(new BigNumber(2)).toNumber()] = _1407_v1;
          _nw218[(new BigNumber(3)).toNumber()] = _dafny.Seq.of(_1485_v41);
          _nw218[(new BigNumber(4)).toNumber()] = _1407_v1;
          _nw218[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_1407_v1, _1407_v1);
          _nw218[(new BigNumber(6)).toNumber()] = _1407_v1;
          _nw218[(new BigNumber(7)).toNumber()] = _module.__default.fm20(new BigNumber(753), p1, !(false), (_1479_v38)[_module.__default.safeIndex(new BigNumber(494), new BigNumber((_1479_v38).length))], globalState);
          _nw218[(new BigNumber(8)).toNumber()] = _1407_v1;
          _nw218[(new BigNumber(9)).toNumber()] = _dafny.Seq.of(_1485_v41);
          _nw218[(new BigNumber(10)).toNumber()] = _1407_v1;
          _1486_v42 = _nw218;
          let _index235 = _module.__default.safeIndex(new BigNumber(397), new BigNumber((_1486_v42).length));
          (_1486_v42)[_index235] = _1407_v1;
          let _index236 = _module.__default.safeIndex(new BigNumber(397), new BigNumber((_1486_v42).length));
          (_1486_v42)[_index236] = _dafny.Seq.Concat(_1407_v1, _1407_v1);
        }
        if (false) {
          let _1487_v43;
          _1487_v43 = _dafny.Map.Empty.slice().updateUnsafe(_1406_i0,(_this).f10);
          let _1488_v45;
          _1488_v45 = _dafny.Seq.of(_1487_v43, _1487_v43, (_1487_v43).Merge(function () {
            let _coll53 = new _dafny.Map();
            for (const _compr_53 of _dafny.IntegerRange(new BigNumber(185), new BigNumber(321))) {
              let _1489_v44 = _compr_53;
              if (((new BigNumber(185)).isLessThanOrEqualTo(_1489_v44)) && ((_1489_v44).isLessThan(new BigNumber(321)))) {
                _coll53.push([(_1489_v44).multipliedBy(new BigNumber(887)),p0]);
              }
            }
            return _coll53;
          }()), (_module.__default.fm44(p1, globalState)).Merge(_module.__default.fm44(p1, globalState)), _1487_v43);
          let _1490_v46;
          _1490_v46 = _module.D16.create_DC40(_1488_v45);
          _1488_v45 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_1487_v43), (_1490_v46).dtor_cf76), _dafny.Seq.Concat(_1488_v45, _1488_v45)), _module.__default.safeIndex((_dafny.ZERO).minus(_1406_i0), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_1487_v43), (_1490_v46).dtor_cf76), _dafny.Seq.Concat(_1488_v45, _1488_v45))).length)), _1487_v43);
          let _1491_v47;
          _1491_v47 = _dafny.Seq.of((_this).f10, p0, p1);
          (globalState).f8 = !((_this).fm0(_1491_v47, (p1) === ((_this).f10), _module.__default.safeDivisionInt(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-494))).cardinality()), _1405_v0), p0, globalState));
          let _1492_v48;
          let _nw219 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
          _1492_v48 = _nw219;
          let _1493_v49;
          _1493_v49 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),_1492_v48);
          let _1494_v50;
          _1494_v50 = new _dafny.CodePoint('u'.codePointAt(0));
          let _1495_v51;
          let _nw220 = new _module.C3();
          _nw220.__ctor();
          _1495_v51 = _nw220;
          let _1496_v52;
          _1496_v52 = _dafny.Map.Empty.slice().updateUnsafe((((_1493_v49).contains(_1494_v50)) ? ((_1493_v49).get(_1494_v50)) : (_1492_v48)),_1495_v51);
          _1496_v52 = ((_1496_v52).Merge(_1496_v52)).Merge(_1496_v52);
          _1494_v50 = _1494_v50;
          _1405_v0 = _1406_i0;
        } else {
          let _1497_v53;
          _1497_v53 = _module.D1.create_DC3(_1407_v1);
          let _1498_v54;
          _1498_v54 = _dafny.Seq.of((_1405_v0).plus(_1405_v0), new BigNumber(((_1497_v53).dtor_cf5).length));
          _1498_v54 = _1498_v54;
          _1405_v0 = (_dafny.ZERO).minus(((_dafny.ZERO).minus(new BigNumber((_1407_v1).length))).plus(_1406_i0));
          let _1499_v55;
          let _nw221 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
          _1499_v55 = _nw221;
          let _1500_v56;
          _1500_v56 = _dafny.Seq.of(p1, p0);
          let _1501_v57;
          _1501_v57 = _dafny.MultiSet.fromElements(_1500_v56, _1500_v56);
          let _1502_v58;
          _1502_v58 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("ayxnih")).length),_1406_i0);
          let _rhs233 = _1499_v55;
          let _rhs234 = (((_1406_i0).isLessThanOrEqualTo(_1405_v0)) ? (_module.__default.safeDivisionInt(_1405_v0, new BigNumber((_1501_v57).cardinality()))) : ((((_1502_v58).contains((_dafny.ZERO).minus(_1405_v0))) ? ((_1502_v58).get((_dafny.ZERO).minus(_1405_v0))) : (_1406_i0))));
          let _rhs235 = (_dafny.ZERO).minus(_1406_i0);
          _1499_v55 = _rhs233;
          _1405_v0 = _rhs234;
          _1405_v0 = _rhs235;
          let _1503_v59;
          _1503_v59 = _dafny.Set.fromElements(_dafny.Seq.Concat(_1498_v54, _1498_v54));
          _1405_v0 = new BigNumber((_1503_v59).length);
          let _1504_v60;
          _1504_v60 = _dafny.MultiSet.fromElements(false, true, false, p1, p1);
          let _1505_v61;
          _1505_v61 = _dafny.Set.fromElements(_1499_v55, _1499_v55, _1499_v55, _1499_v55, _1499_v55);
          let _1506_v62;
          _1506_v62 = _module.D8.create_DC24((_1504_v60).IsProperSubsetOf(_1504_v60), _1505_v61, _1405_v0);
          let _rhs236 = _1506_v62;
          let _rhs237 = p0;
          let _rhs238 = p0;
          let _rhs239 = (_1405_v0).plus(_1405_v0);
          let _rhs240 = (_this).f10;
          let _lhs166 = globalState;
          let _lhs167 = globalState;
          let _lhs168 = globalState;
          _1506_v62 = _rhs236;
          _lhs166.f8 = _rhs237;
          _lhs167.f8 = _rhs238;
          _1405_v0 = _rhs239;
          _lhs168.f8 = _rhs240;
        }
        let _1507_v63;
        let _nw222 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Map.Empty);
        _1507_v63 = _nw222;
        let _1508_v64;
        _1508_v64 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,new BigNumber((_1407_v1).length));
        let _1509_v65;
        _1509_v65 = _dafny.Seq.of(_1406_i0, _1406_i0);
        let _1510_v66;
        _1510_v66 = _dafny.Map.Empty.slice().updateUnsafe(_1508_v64,_1509_v65);
        let _index237 = _module.__default.safeIndex(new BigNumber(396), new BigNumber((_1507_v63).length));
        (_1507_v63)[_index237] = (_1510_v66).Merge(_1510_v66);
        let _1511_v67;
        _1511_v67 = _dafny.Seq.of(_1510_v66, _1510_v66, _dafny.Map.Empty.slice().updateUnsafe(_1508_v64,_1509_v65));
        let _index238 = _module.__default.safeIndex(new BigNumber(396), new BigNumber((_1507_v63).length));
        (_1507_v63)[_index238] = ((_1510_v66).Merge((_1511_v67)[_module.__default.safeIndex(_1405_v0, new BigNumber((_1511_v67).length))])).Merge(_1510_v66);
      }
      (globalState).f8 = p1;
      let _1512_v68;
      _1512_v68 = _dafny.MultiSet.fromElements((_1405_v0).minus(_1405_v0), (_this).fm1((_this).f9, globalState), new BigNumber(208), _1405_v0);
      _1512_v68 = (_1512_v68).update(_module.__default.safeDivisionInt(_1405_v0, _1405_v0), _module.__default.abs(_1405_v0));
      (globalState).f8 = p0;
      let _1513_v69;
      _1513_v69 = _module.D5.create_DC14(new BigNumber(75), p1);
      let _1514_v70;
      let _nw223 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
      _1514_v70 = _nw223;
      let _1515_v71;
      _1515_v71 = _dafny.Set.fromElements(_1514_v70);
      let _1516_v72;
      _1516_v72 = _dafny.Seq.of(_1405_v0);
      let _1517_v73;
      _1517_v73 = _dafny.Seq.UnicodeFromString("jeudkbik");
      let _1518_v74;
      _1518_v74 = _dafny.Map.Empty.slice().updateUnsafe(_1405_v0,p0);
      let _1519_v75;
      _1519_v75 = new _dafny.CodePoint('e'.codePointAt(0));
      let _1520_v76;
      let _nw224 = Array((new BigNumber(27)).toNumber());
      _nw224[(_dafny.ZERO).toNumber()] = _1405_v0;
      _nw224[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_1513_v69).dtor_cf25),new BigNumber((_1515_v71).length))).length);
      _nw224[(new BigNumber(2)).toNumber()] = _1405_v0;
      _nw224[(new BigNumber(3)).toNumber()] = _1405_v0;
      _nw224[(new BigNumber(4)).toNumber()] = new BigNumber(-813);
      _nw224[(new BigNumber(5)).toNumber()] = _1405_v0;
      _nw224[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_1516_v72, _1516_v72)).length));
      _nw224[(new BigNumber(7)).toNumber()] = _1405_v0;
      _nw224[(new BigNumber(8)).toNumber()] = new BigNumber((_1517_v73).length);
      _nw224[(new BigNumber(9)).toNumber()] = new BigNumber((_1518_v74).length);
      _nw224[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus(_1405_v0);
      _nw224[(new BigNumber(11)).toNumber()] = (_1405_v0).multipliedBy(new BigNumber(184));
      _nw224[(new BigNumber(12)).toNumber()] = _1405_v0;
      _nw224[(new BigNumber(13)).toNumber()] = _1405_v0;
      _nw224[(new BigNumber(14)).toNumber()] = _1405_v0;
      _nw224[(new BigNumber(15)).toNumber()] = _1405_v0;
      _nw224[(new BigNumber(16)).toNumber()] = (_1405_v0).minus(_1405_v0);
      _nw224[(new BigNumber(17)).toNumber()] = (_1516_v72)[_module.__default.safeIndex(_1405_v0, new BigNumber((_1516_v72).length))];
      _nw224[(new BigNumber(18)).toNumber()] = _1405_v0;
      _nw224[(new BigNumber(19)).toNumber()] = _module.__default.safeModuloInt(_1405_v0, _1405_v0);
      _nw224[(new BigNumber(20)).toNumber()] = new BigNumber(822);
      _nw224[(new BigNumber(21)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_1405_v0, _1405_v0));
      _nw224[(new BigNumber(22)).toNumber()] = new BigNumber((_1517_v73).length);
      _nw224[(new BigNumber(23)).toNumber()] = new BigNumber(104);
      _nw224[(new BigNumber(24)).toNumber()] = _1405_v0;
      _nw224[(new BigNumber(25)).toNumber()] = (_1405_v0).minus(new BigNumber(464));
      _nw224[(new BigNumber(26)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1519_v75,_dafny.Seq.update(_dafny.Seq.UnicodeFromString("nbqhra"), _module.__default.safeIndex(_1405_v0, new BigNumber((_dafny.Seq.UnicodeFromString("nbqhra")).length)), _1519_v75))).length);
      _1520_v76 = _nw224;
      let _index239 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_1514_v70).length));
      (_1514_v70)[_index239] = _1405_v0;
      let _1521_v77;
      _1521_v77 = _module.D6.create_DC18(new _dafny.CodePoint('m'.codePointAt(0)), _1405_v0, _1405_v0);
      let _pat_let_tv31 = p1;
      let _pat_let_tv32 = _1405_v0;
      let _pat_let_tv33 = _1405_v0;
      let _pat_let_tv34 = _1517_v73;
      let _pat_let_tv35 = _1517_v73;
      let _pat_let_tv36 = _1517_v73;
      let _index240 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_1514_v70).length));
      (_1514_v70)[_index240] = function (_source20) {
        if (_source20.is_DC17) {
          let _1522___mcc_h12 = (_source20).cf29;
          let _1523___mcc_h13 = (_source20).cf30;
          let _1524___mcc_h14 = (_source20).cf31;
          let _1525___mcc_h15 = (_source20).cf32;
          let _1526___mcc_h16 = (_source20).cf33;
          let _1527_cf33 = _1526___mcc_h16;
          let _1528_cf32 = _1525___mcc_h15;
          let _1529_cf31 = _1524___mcc_h14;
          let _1530_cf30 = _1523___mcc_h13;
          let _1531_cf29 = _1522___mcc_h12;
          return new BigNumber(((_dafny.Set.fromElements(_1530_cf30, _pat_let_tv31)).Intersect(_dafny.Set.fromElements(_1531_cf29))).length);
        } else if (_source20.is_DC18) {
          let _1532___mcc_h17 = (_source20).cf34;
          let _1533___mcc_h18 = (_source20).cf35;
          let _1534___mcc_h19 = (_source20).cf36;
          let _1535_cf36 = _1534___mcc_h19;
          let _1536_cf35 = _1533___mcc_h18;
          let _1537_cf34 = _1532___mcc_h17;
          return _pat_let_tv32;
        } else if (_source20.is_DC19) {
          let _1538___mcc_h20 = (_source20).cf37;
          let _1539___mcc_h21 = (_source20).cf38;
          let _1540_cf38 = _1539___mcc_h21;
          let _1541_cf37 = _1538___mcc_h20;
          return _pat_let_tv33;
        } else if (_source20.is_DC16) {
          let _1542___mcc_h22 = (_source20).cf28;
          let _1543_cf28 = _1542___mcc_h22;
          return (new BigNumber((_pat_let_tv34).length)).minus((_dafny.ZERO).minus(new BigNumber((_1543_cf28).length)));
        } else {
          let _1544___mcc_h23 = (_source20).cf39;
          let _1545_cf39 = _1544___mcc_h23;
          return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_pat_let_tv35, _pat_let_tv36)).length));
        }
      }(_1521_v77);
      for (const _guard_loop_5 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1520_v76).length))) {
        let _1546_i4 = _guard_loop_5;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1546_i4)) && ((_1546_i4).isLessThan(new BigNumber((_1520_v76).length))))) {
          (_1520_v76)[(_1546_i4)] = (_1546_i4).multipliedBy(new BigNumber(703));
        }
      }
      let _1547_v78;
      let _nw225 = Array((_dafny.ONE).toNumber()).fill(false);
      _1547_v78 = _nw225;
      r0 = _1547_v78;
      let _1548_v79;
      _1548_v79 = _module.D0.create_DC0(p1);
      let _1549_v80;
      _1549_v80 = _dafny.Seq.of(_1548_v79, _1548_v79, _1548_v79, _module.D0.create_DC0(p0));
      r1 = _1549_v80;
      let _1550_v82;
      _1550_v82 = _module.D4.create_DC11(new BigNumber(919), function () {
  let _coll54 = new _dafny.Set();
  for (const _compr_54 of _dafny.IntegerRange(new BigNumber(742), new BigNumber(-453))) {
    let _1551_v81 = _compr_54;
    if (((new BigNumber(742)).isLessThanOrEqualTo(_1551_v81)) && ((_1551_v81).isLessThan(new BigNumber(-453)))) {
      _coll54.add((_1551_v81).minus((_1514_v70)[_module.__default.safeIndex(new BigNumber(655), new BigNumber((_1514_v70).length))]));
    }
  }
  return _coll54;
}(), true);
      r2 = (((false) ? (_module.D4.create_DC11(_1405_v0, _dafny.Set.fromElements(new BigNumber((_1517_v73).length)), p1)) : (_1550_v82))).dtor_cf17;
      return [r0, r1, r2];
    }
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this._f9 = _dafny.Map.Empty;
      this._f10 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f10() {
      let _this = this;
      return _this._f10;
    };
    __ctor(f9, f10) {
      let _this = this;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      return;
    }
    fm0(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f10;
    };
    fm1(p0, globalState) {
      let _this = this;
      return new BigNumber(93);
    };
    fm8(globalState) {
      let _this = this;
      let _source21 = _module.D1.create_DC4((_this).f10, (_this).f10, new BigNumber(-971));
      if (_source21.is_DC3) {
        let _1552___mcc_h0 = (_source21).cf5;
        let _1553_cf5 = _1552___mcc_h0;
        return (new BigNumber((_1553_cf5).length)).minus(new BigNumber(-673));
      } else if (_source21.is_DC4) {
        let _1554___mcc_h1 = (_source21).cf6;
        let _1555___mcc_h2 = (_source21).cf7;
        let _1556___mcc_h3 = (_source21).cf8;
        let _1557_cf8 = _1556___mcc_h3;
        let _1558_cf7 = _1555___mcc_h2;
        let _1559_cf6 = _1554___mcc_h1;
        return _1557_cf8;
      } else if (_source21.is_DC2) {
        let _1560___mcc_h4 = (_source21).cf4;
        let _1561_cf4 = _1560___mcc_h4;
        return (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(862)), function (_1562_i0) {
          return new BigNumber((_dafny.Seq.of(_dafny.MultiSet.fromElements((_this).f10, (_module.D0.create_DC1(new BigNumber(576), (_this).f10, new BigNumber((_dafny.Seq.UnicodeFromString("qfvjs")).length))).dtor_cf2), _dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f10)))).length);
        })).length)).plus(new BigNumber((_dafny.Seq.of((_this).f10, true, (_this).f10, (_this).f10, (_this).f10)).length));
      } else {
        let _1563___mcc_h5 = (_source21).cf9;
        let _1564_cf9 = _1563___mcc_h5;
        return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(753)), function (_1565_i1) {
          return _dafny.Set.fromElements(new BigNumber(462), new BigNumber(3));
        })).length));
      }
    };
    fm9(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !((_this).f10);
    };
    m0(globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = false;
      let r2 = _dafny.Seq.of();
      let r3 = _dafny.Seq.UnicodeFromString("");
      if (!(!((_this).f10) || (true))) {
        let _1566_v0;
        let _init38 = function (_1567_i0) {
          return (_1567_i0).minus(new BigNumber(((_this).f9).length));
        };
        let _nw226 = Array((new BigNumber(21)).toNumber());
        for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw226.length); _i0_38++) {
          _nw226[_i0_38] = _init38(new BigNumber(_i0_38));
        }
        _1566_v0 = _nw226;
        let _1568_v1;
        _1568_v1 = _dafny.Set.fromElements((_this).f10, (_this).f10, (_this).f10);
        let _index241 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_1566_v0).length));
        (_1566_v0)[_index241] = (new BigNumber((_1568_v1).length)).minus(new BigNumber(((_this).f9).length));
        let _1569_v2;
        _1569_v2 = new BigNumber(-316);
        let _index242 = _module.__default.safeIndex(new BigNumber(333), new BigNumber((_1566_v0).length));
        (_1566_v0)[_index242] = (_1569_v2).multipliedBy(_1569_v2);
        let _1570_v4;
        let _init39 = ((_1571_v0) => function (_1572_i1) {
          return (function () {
            let _coll55 = new _dafny.Set();
            for (const _compr_55 of _dafny.IntegerRange(new BigNumber(249), new BigNumber(61))) {
              let _1573_v3 = _compr_55;
              if (((new BigNumber(249)).isLessThanOrEqualTo(_1573_v3)) && ((_1573_v3).isLessThan(new BigNumber(61)))) {
                _coll55.add((_1573_v3).multipliedBy((_1571_v0)[_module.__default.safeIndex(new BigNumber(333), new BigNumber((_1571_v0).length))]));
              }
            }
            return _coll55;
          }()).Difference(_dafny.Set.fromElements((_1571_v0)[_module.__default.safeIndex(new BigNumber(333), new BigNumber((_1571_v0).length))]));
        })(_1566_v0);
        let _nw227 = Array((new BigNumber(19)).toNumber());
        for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw227.length); _i0_39++) {
          _nw227[_i0_39] = _init39(new BigNumber(_i0_39));
        }
        _1570_v4 = _nw227;
        _1570_v4 = _1570_v4;
        let _1574_v5;
        let _init40 = function (_1575_i2) {
          return _module.D0.create_DC0((_this).f10);
        };
        let _nw228 = Array((new BigNumber(11)).toNumber());
        for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw228.length); _i0_40++) {
          _nw228[_i0_40] = _init40(new BigNumber(_i0_40));
        }
        _1574_v5 = _nw228;
        let _1576_v6;
        _1576_v6 = _module.D0.create_DC0((_this).f10);
        let _index243 = _module.__default.safeIndex(new BigNumber(302), new BigNumber((_1574_v5).length));
        (_1574_v5)[_index243] = function (_pat_let34_0) {
          return function (_1577_dt__update__tmp_h0) {
            return function (_pat_let35_0) {
              return function (_1578_dt__update_hcf0_h0) {
                return _module.D0.create_DC0(_1578_dt__update_hcf0_h0);
              }(_pat_let35_0);
            }(!(!((_this).f10)));
          }(_pat_let34_0);
        }(_1576_v6);
        let _index244 = _module.__default.safeIndex(new BigNumber(302), new BigNumber((_1574_v5).length));
        (_1574_v5)[_index244] = function (_pat_let36_0) {
          return function (_1579_dt__update__tmp_h1) {
            return function (_pat_let37_0) {
              return function (_1580_dt__update_hcf0_h1) {
                return _module.D0.create_DC0(_1580_dt__update_hcf0_h1);
              }(_pat_let37_0);
            }((_this).f10);
          }(_pat_let36_0);
        }(_1576_v6);
        let _1581_v7;
        let _nw229 = new _module.C0();
        _nw229.__ctor();
        _1581_v7 = _nw229;
        let _1582_v8;
        _1582_v8 = _dafny.Seq.UnicodeFromString("qbbn");
        let _1583_v9;
        _1583_v9 = _dafny.Map.Empty.slice().updateUnsafe((_1566_v0)[_module.__default.safeIndex(new BigNumber(333), new BigNumber((_1566_v0).length))],_1569_v2);
        let _1584_v10;
        _1584_v10 = _dafny.MultiSet.fromElements((_1566_v0)[_module.__default.safeIndex(new BigNumber(333), new BigNumber((_1566_v0).length))], new BigNumber(312));
        (globalState).f8 = (_module.__default.fm10((_this).f10, _1569_v2, _1582_v8, _1583_v9, globalState)).IsDisjointFrom(_1584_v10);
      } else {
        let _1585_v11;
        let _nw230 = new _module.C0();
        _nw230.__ctor();
        _1585_v11 = _nw230;
        let _1586_v12;
        let _init41 = function (_1587_i3) {
          return (_this).f10;
        };
        let _nw231 = Array((new BigNumber(21)).toNumber());
        for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw231.length); _i0_41++) {
          _nw231[_i0_41] = _init41(new BigNumber(_i0_41));
        }
        _1586_v12 = _nw231;
        let _index245 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_1586_v12).length));
        (_1586_v12)[_index245] = ((_this).f10) === ((_this).f10);
        let _1588_v13;
        _1588_v13 = new BigNumber(274);
        let _index246 = _module.__default.safeIndex(new BigNumber(294), new BigNumber((_1586_v12).length));
        (_1586_v12)[_index246] = (_1588_v13).isLessThan((_dafny.ZERO).minus(_1588_v13));
        let _1589_v14;
        _1589_v14 = _module.D1.create_DC4((_this).f10, false, (_this).fm1((_this).f9, globalState));
        let _rhs241 = _1589_v14;
        let _rhs242 = (_this).f10;
        let _rhs243 = _1588_v13;
        let _rhs244 = false;
        let _lhs169 = globalState;
        let _lhs170 = globalState;
        _1589_v14 = _rhs241;
        _lhs169.f8 = _rhs242;
        _1588_v13 = _rhs243;
        _lhs170.f8 = _rhs244;
        let _1590_v15;
        _1590_v15 = _dafny.Map.Empty.slice().updateUnsafe(_1588_v13,(_this).f9);
        if (((((_1590_v15).contains(_1588_v13)) ? ((_1590_v15).get(_1588_v13)) : ((_this).f9))).equals((_this).f9)) {
          let _1591_v16;
          _1591_v16 = _dafny.Map.Empty.slice().updateUnsafe(_1588_v13,_1588_v13);
          let _1592_v17;
          _1592_v17 = _dafny.MultiSet.fromElements((_this).f10, (_1586_v12)[_module.__default.safeIndex(new BigNumber(294), new BigNumber((_1586_v12).length))], (_1586_v12)[_module.__default.safeIndex(new BigNumber(294), new BigNumber((_1586_v12).length))], (_1586_v12)[_module.__default.safeIndex(new BigNumber(294), new BigNumber((_1586_v12).length))], (_1586_v12)[_module.__default.safeIndex(new BigNumber(294), new BigNumber((_1586_v12).length))]);
          _1588_v13 = (_1588_v13).multipliedBy((((_1591_v16).contains(new BigNumber((_1592_v17).cardinality()))) ? ((_1591_v16).get(new BigNumber((_1592_v17).cardinality()))) : (_1588_v13)));
          let _1593_v18;
          _1593_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1588_v13,(_1586_v12)[_module.__default.safeIndex(new BigNumber(294), new BigNumber((_1586_v12).length))]);
          let _1594_v19;
          _1594_v19 = _dafny.MultiSet.fromElements(_1588_v13);
          let _1595_v20;
          _1595_v20 = _dafny.Set.fromElements(_1594_v19);
          let _1596_v22;
          _1596_v22 = _dafny.Seq.of(_1588_v13, _1588_v13);
          let _1597_v23;
          let _nw232 = Array((new BigNumber(21)).toNumber());
          _nw232[(_dafny.ZERO).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_1588_v13,(_1586_v12)[_module.__default.safeIndex(new BigNumber(294), new BigNumber((_1586_v12).length))])).Merge(_1593_v18);
          _nw232[(_dafny.ONE).toNumber()] = _1593_v18;
          _nw232[(new BigNumber(2)).toNumber()] = (_module.__default.fm11(new BigNumber((_1595_v20).length), _1588_v13, false, _1588_v13, globalState)).Merge(_1593_v18);
          _nw232[(new BigNumber(3)).toNumber()] = (_1593_v18).Merge(_1593_v18);
          _nw232[(new BigNumber(4)).toNumber()] = _1593_v18;
          _nw232[(new BigNumber(5)).toNumber()] = _1593_v18;
          _nw232[(new BigNumber(6)).toNumber()] = _1593_v18;
          _nw232[(new BigNumber(7)).toNumber()] = _1593_v18;
          _nw232[(new BigNumber(8)).toNumber()] = _1593_v18;
          _nw232[(new BigNumber(9)).toNumber()] = _1593_v18;
          _nw232[(new BigNumber(10)).toNumber()] = function () {
            let _coll56 = new _dafny.Map();
            for (const _compr_56 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber(445))) {
              let _1598_v21 = _compr_56;
              if (((_dafny.ZERO).isLessThanOrEqualTo(_1598_v21)) && ((_1598_v21).isLessThan(new BigNumber(445)))) {
                _coll56.push([(_1598_v21).multipliedBy(new BigNumber(351)),(_1586_v12)[_module.__default.safeIndex(new BigNumber(294), new BigNumber((_1586_v12).length))]]);
              }
            }
            return _coll56;
          }();
          _nw232[(new BigNumber(11)).toNumber()] = _module.__default.fm11((_1596_v22)[_module.__default.safeIndex(new BigNumber((_1591_v16).length), new BigNumber((_1596_v22).length))], new BigNumber(439), (_1586_v12)[_module.__default.safeIndex(new BigNumber(294), new BigNumber((_1586_v12).length))], _1588_v13, globalState);
          _nw232[(new BigNumber(12)).toNumber()] = _1593_v18;
          _nw232[(new BigNumber(13)).toNumber()] = _1593_v18;
          _nw232[(new BigNumber(14)).toNumber()] = (_1593_v18).Merge(_1593_v18);
          _nw232[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1588_v13,(_1586_v12)[_module.__default.safeIndex(new BigNumber(294), new BigNumber((_1586_v12).length))]);
          _nw232[(new BigNumber(16)).toNumber()] = _1593_v18;
          _nw232[(new BigNumber(17)).toNumber()] = _1593_v18;
          _nw232[(new BigNumber(18)).toNumber()] = _module.__default.fm11(new BigNumber(-251), _1588_v13, (_this).f10, _1588_v13, globalState);
          _nw232[(new BigNumber(19)).toNumber()] = _1593_v18;
          _nw232[(new BigNumber(20)).toNumber()] = _1593_v18;
          _1597_v23 = _nw232;
          let _index247 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_1597_v23).length));
          (_1597_v23)[_index247] = (((_this).f10) ? (_1593_v18) : (_1593_v18));
          let _1599_v24;
          _1599_v24 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_1593_v18).Merge(_1593_v18));
          let _1600_v25;
          _1600_v25 = _dafny.Seq.UnicodeFromString("tnx");
          let _1601_v26;
          _1601_v26 = _dafny.Set.fromElements(_1600_v25);
          let _index248 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_1597_v23).length));
          let _rhs245 = new BigNumber(-464);
          let _rhs246 = (_1586_v12)[_module.__default.safeIndex(new BigNumber(294), new BigNumber((_1586_v12).length))];
          let _rhs247 = (((_1599_v24).contains(!((_1601_v26).IsSubsetOf(_1601_v26)))) ? ((_1599_v24).get(!((_1601_v26).IsSubsetOf(_1601_v26)))) : (_dafny.Map.Empty.slice().updateUnsafe(_1588_v13,(_1586_v12)[_module.__default.safeIndex(new BigNumber(294), new BigNumber((_1586_v12).length))])));
          let _rhs248 = !_dafny.areEqual(_1600_v25, _dafny.Seq.UnicodeFromString("xiesenboe"));
          let _lhs171 = _1597_v23;
          let _lhs172 = _module.__default.safeIndex(new BigNumber(985), new BigNumber((_1597_v23).length));
          _1588_v13 = _rhs245;
          r1 = _rhs246;
          _lhs171[_lhs172] = _rhs247;
          r1 = _rhs248;
          let _1602_v27;
          let _nw233 = new _module.C0();
          _nw233.__ctor();
          _1602_v27 = _nw233;
          let _1603_v28;
          let _nw234 = new _module.C0();
          _nw234.__ctor();
          _1603_v28 = _nw234;
          (globalState).f8 = (_1586_v12)[_module.__default.safeIndex(new BigNumber(294), new BigNumber((_1586_v12).length))];
        } else {
          let _1604_v29;
          _1604_v29 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm13(_1588_v13, globalState)).length),_module.__default.fm14(globalState));
          let _1605_v30;
          _1605_v30 = new _dafny.CodePoint('q'.codePointAt(0));
          let _1606_v33;
          _1606_v33 = _dafny.Seq.UnicodeFromString("qxxtkeihl");
          let _1607_v34;
          _1607_v34 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("rvcpxn"),_1588_v13);
          let _1608_v35;
          _1608_v35 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1607_v34).length),_dafny.Map.Empty.slice().updateUnsafe(_1588_v13,_1605_v30));
          let _1609_v36;
          _1609_v36 = _dafny.Seq.of(_1588_v13, _1588_v13);
          let _1610_v37;
          let _nw235 = Array((new BigNumber(19)).toNumber());
          _nw235[(_dafny.ZERO).toNumber()] = (_module.__default.fm12((_1586_v12)[_module.__default.safeIndex(new BigNumber(294), new BigNumber((_1586_v12).length))], _1588_v13, (_dafny.ZERO).minus(_1588_v13), (_1586_v12)[_module.__default.safeIndex(new BigNumber(294), new BigNumber((_1586_v12).length))], globalState)).Merge(_1604_v29);
          _nw235[(_dafny.ONE).toNumber()] = _1604_v29;
          _nw235[(new BigNumber(2)).toNumber()] = (_1604_v29).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1588_v13,_1605_v30));
          _nw235[(new BigNumber(3)).toNumber()] = _1604_v29;
          _nw235[(new BigNumber(4)).toNumber()] = (_1604_v29).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1588_v13,_1605_v30));
          _nw235[(new BigNumber(5)).toNumber()] = _1604_v29;
          _nw235[(new BigNumber(6)).toNumber()] = function () {
            let _coll57 = new _dafny.Map();
            for (const _compr_57 of _dafny.IntegerRange(new BigNumber(123), new BigNumber(-659))) {
              let _1611_v31 = _compr_57;
              if (((new BigNumber(123)).isLessThanOrEqualTo(_1611_v31)) && ((_1611_v31).isLessThan(new BigNumber(-659)))) {
                _coll57.push([(_1611_v31).plus(_1588_v13),_1605_v30]);
              }
            }
            return _coll57;
          }();
          _nw235[(new BigNumber(7)).toNumber()] = (_1604_v29).Merge(_1604_v29);
          _nw235[(new BigNumber(8)).toNumber()] = function () {
            let _coll58 = new _dafny.Map();
            for (const _compr_58 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(528)), function (_1612_i4) {
              return new BigNumber(360);
            })).Elements) {
              let _1613_v32 = _compr_58;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(528)), function (_1612_i4) {
                return new BigNumber(360);
              }), _1613_v32)) {
                _coll58.push([(_1613_v32).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements((_this).f10)).cardinality())),_1605_v30]);
              }
            }
            return _coll58;
          }();
          _nw235[(new BigNumber(9)).toNumber()] = _1604_v29;
          _nw235[(new BigNumber(10)).toNumber()] = (_module.__default.fm12(!((_this).f10), _1588_v13, new BigNumber((_1606_v33).length), (_this).f10, globalState)).Merge(_1604_v29);
          _nw235[(new BigNumber(11)).toNumber()] = _1604_v29;
          _nw235[(new BigNumber(12)).toNumber()] = (_1604_v29).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1588_v13,_1605_v30));
          _nw235[(new BigNumber(13)).toNumber()] = ((_1604_v29).update(_1588_v13, _1605_v30)).Merge((((_1608_v35).contains(_1588_v13)) ? ((_1608_v35).get(_1588_v13)) : (_module.__default.fm12(true, _1588_v13, _1588_v13, (_this).f10, globalState))));
          _nw235[(new BigNumber(14)).toNumber()] = _module.__default.fm12((_this).f10, new BigNumber((_1606_v33).length), new BigNumber((_1609_v36).length), (_this).f10, globalState);
          _nw235[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1588_v13,_1605_v30);
          _nw235[(new BigNumber(16)).toNumber()] = _1604_v29;
          _nw235[(new BigNumber(17)).toNumber()] = _1604_v29;
          _nw235[(new BigNumber(18)).toNumber()] = _1604_v29;
          _1610_v37 = _nw235;
          _1610_v37 = _1610_v37;
          _1606_v33 = _1606_v33;
          let _1614_v38;
          let _init42 = ((_1615_v13) => function (_1616_i5) {
            return (_1616_i5).minus(_1615_v13);
          })(_1588_v13);
          let _nw236 = Array((new BigNumber(25)).toNumber());
          for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw236.length); _i0_42++) {
            _nw236[_i0_42] = _init42(new BigNumber(_i0_42));
          }
          _1614_v38 = _nw236;
          let _index249 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_1614_v38).length));
          (_1614_v38)[_index249] = new BigNumber(940);
          let _index250 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_1614_v38).length));
          (_1614_v38)[_index250] = (new BigNumber(607)).minus(_1588_v13);
          (globalState).f8 = (_1586_v12)[_module.__default.safeIndex(new BigNumber(294), new BigNumber((_1586_v12).length))];
          let _1617_v39;
          _1617_v39 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(696),(_1586_v12)[_module.__default.safeIndex(new BigNumber(294), new BigNumber((_1586_v12).length))]);
          let _1618_v40;
          _1618_v40 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_module.__default.fm13(_1588_v13, globalState));
          let _rhs249 = !(((_this).f10) || ((_1586_v12)[_module.__default.safeIndex(new BigNumber(294), new BigNumber((_1586_v12).length))])) || ((_this).fm9((_1614_v38)[_module.__default.safeIndex(new BigNumber(758), new BigNumber((_1614_v38).length))], _1588_v13, _1588_v13, _1618_v40, globalState));
          let _rhs250 = (((_this).f10) ? (_1617_v39) : (_1617_v39));
          let _lhs173 = globalState;
          _lhs173.f8 = _rhs249;
          _1617_v39 = _rhs250;
        }
        (globalState).f8 = !(!((_1586_v12)[_module.__default.safeIndex(new BigNumber(294), new BigNumber((_1586_v12).length))])) || ((_this).f10);
      }
      let _1619_v41;
      _1619_v41 = _dafny.Seq.UnicodeFromString("fuy");
      (globalState).f8 = _dafny.Seq.IsPrefixOf(_1619_v41, _dafny.Seq.UnicodeFromString("d"));
      let _1620_v42;
      _1620_v42 = new BigNumber(-36);
      let _1621_v43;
      _1621_v43 = _dafny.MultiSet.fromElements(_1620_v42, new BigNumber(94));
      let _1622_v44;
      _1622_v44 = _dafny.MultiSet.fromElements(_1621_v43, _1621_v43);
      let _1623_i6;
      _1623_i6 = _dafny.ZERO;
      L7: {
        while ((_1622_v44).IsSubsetOf(_1622_v44)) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1623_i6)) {
              break L7;
            }
            _1623_i6 = (_1623_i6).plus(_dafny.ONE);
            let _1624_v45;
            let _nw237 = new _module.C0();
            _nw237.__ctor();
            _1624_v45 = _nw237;
            _1620_v42 = _1620_v42;
            let _1625_v46;
            _1625_v46 = _dafny.MultiSet.fromElements((_this).f10, !(false), (_this).f10, (_this).f10, (_this).f10);
            let _1626_v47;
            _1626_v47 = _dafny.Set.fromElements((_this).f10);
            let _1627_v48;
            _1627_v48 = _dafny.Seq.of(new BigNumber((_1626_v47).length));
            let _1628_v49;
            _1628_v49 = _dafny.Map.Empty.slice().updateUnsafe(_1620_v42,(_1627_v48)[_module.__default.safeIndex(new BigNumber(806), new BigNumber((_1627_v48).length))]);
            let _1629_v50;
            _1629_v50 = _dafny.Map.Empty.slice().updateUnsafe((_1625_v46).contains((_this).f10),(new BigNumber((_1619_v41).length)).multipliedBy((((_1628_v49).contains(new BigNumber((_dafny.Seq.of(_1620_v42, new BigNumber(296))).length))) ? ((_1628_v49).get(new BigNumber((_dafny.Seq.of(_1620_v42, new BigNumber(296))).length))) : (_1620_v42))));
            let _1630_v51;
            _1630_v51 = _dafny.Set.fromElements(new BigNumber(863));
            _1620_v42 = (((_1629_v50).contains((_this).f10)) ? ((_1629_v50).get((_this).f10)) : ((_1620_v42).multipliedBy(new BigNumber((_1630_v51).length))));
            let _1631_v52;
            _1631_v52 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_1619_v41);
            let _1632_v53;
            let _nw238 = Array((new BigNumber(25)).toNumber());
            _nw238[(_dafny.ZERO).toNumber()] = (_this).f10;
            _nw238[(_dafny.ONE).toNumber()] = (_this).f10;
            _nw238[(new BigNumber(2)).toNumber()] = (_this).f10;
            _nw238[(new BigNumber(3)).toNumber()] = true;
            _nw238[(new BigNumber(4)).toNumber()] = (_this).f10;
            _nw238[(new BigNumber(5)).toNumber()] = (_this).f10;
            _nw238[(new BigNumber(6)).toNumber()] = (_this).f10;
            _nw238[(new BigNumber(7)).toNumber()] = !((_this).f10);
            _nw238[(new BigNumber(8)).toNumber()] = (_this).f10;
            _nw238[(new BigNumber(9)).toNumber()] = (_this).f10;
            _nw238[(new BigNumber(10)).toNumber()] = true;
            _nw238[(new BigNumber(11)).toNumber()] = (_this).f10;
            _nw238[(new BigNumber(12)).toNumber()] = (_this).f10;
            _nw238[(new BigNumber(13)).toNumber()] = (_this).f10;
            _nw238[(new BigNumber(14)).toNumber()] = (_this).f10;
            _nw238[(new BigNumber(15)).toNumber()] = (_this).f10;
            _nw238[(new BigNumber(16)).toNumber()] = (_this).f10;
            _nw238[(new BigNumber(17)).toNumber()] = !((_this).f10);
            _nw238[(new BigNumber(18)).toNumber()] = (_this).f10;
            _nw238[(new BigNumber(19)).toNumber()] = (_this).fm9((_this).fm1((_this).f9, globalState), _1620_v42, _1620_v42, _1631_v52, globalState);
            _nw238[(new BigNumber(20)).toNumber()] = (_this).f10;
            _nw238[(new BigNumber(21)).toNumber()] = (_this).f10;
            _nw238[(new BigNumber(22)).toNumber()] = !((_this).f10);
            _nw238[(new BigNumber(23)).toNumber()] = (_this).f10;
            _nw238[(new BigNumber(24)).toNumber()] = (_this).f10;
            _1632_v53 = _nw238;
            let _1633_v54;
            _1633_v54 = _dafny.Seq.of(_1632_v53, _1632_v53);
            let _1634_v55;
            _1634_v55 = _dafny.MultiSet.fromElements((_1633_v54)[_module.__default.safeIndex(new BigNumber((_1627_v48).length), new BigNumber((_1633_v54).length))], (((_this).f10) ? (_1632_v53) : ((_1633_v54)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1620_v42,_1628_v49)).length), new BigNumber((_1633_v54).length))])), _1632_v53, _1632_v53, _1632_v53);
            let _rhs251 = _1634_v55;
            let _rhs252 = !(_1620_v42).isEqualTo(_1620_v42);
            let _lhs174 = globalState;
            _1634_v55 = _rhs251;
            _lhs174.f8 = _rhs252;
          }
        }
      }
      let _1635_v56;
      let _nw239 = Array((new BigNumber(10)).toNumber()).fill(false);
      _1635_v56 = _nw239;
      for (const _guard_loop_6 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1635_v56).length))) {
        let _1636_i7 = _guard_loop_6;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1636_i7)) && ((_1636_i7).isLessThan(new BigNumber((_1635_v56).length))))) {
          (_1635_v56)[(_1636_i7)] = (new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(!(!((_this).f10)),_1620_v42)).Merge(_dafny.Map.Empty.slice().updateUnsafe(!((_this).f10),_1620_v42))).length)).isLessThanOrEqualTo(new BigNumber(939));
        }
      }
      let _1637_v57;
      _1637_v57 = _dafny.Set.fromElements(_1620_v42);
      let _1638_v58;
      _1638_v58 = _dafny.MultiSet.fromElements(!((_this).f10), (_this).f10);
      let _1639_v59;
      _1639_v59 = _dafny.Seq.of(_1638_v58, _1638_v58, _1638_v58);
      let _1640_v60;
      _1640_v60 = _dafny.Map.Empty.slice().updateUnsafe(_1637_v57,_1639_v59);
      _1640_v60 = (_1640_v60).update((((_this).f10) ? (_1637_v57) : (function () {
        let _coll59 = new _dafny.Set();
        for (const _compr_59 of (_dafny.Seq.of(_1620_v42)).Elements) {
          let _1641_v61 = _compr_59;
          if (_dafny.Seq.contains(_dafny.Seq.of(_1620_v42), _1641_v61)) {
            _coll59.add(_module.__default.safeDivisionInt(_1641_v61, new BigNumber((_dafny.Seq.of(true, true, false, false)).length)));
          }
        }
        return _coll59;
      }())), _dafny.Seq.Concat(_dafny.Seq.of(_1638_v58), _1639_v59));
      let _1642_v62;
      _1642_v62 = _module.D0.create_DC1(_module.__default.safeModuloInt(new BigNumber(765), _1620_v42), (_this).f10, _1620_v42);
      let _source22 = _1642_v62;
      if (_source22.is_DC1) {
        let _1643___mcc_h0 = (_source22).cf1;
        let _1644___mcc_h1 = (_source22).cf2;
        let _1645___mcc_h2 = (_source22).cf3;
        let _1646_cf3 = _1645___mcc_h2;
        let _1647_cf2 = _1644___mcc_h1;
        let _1648_cf1 = _1643___mcc_h0;
        let _1649_v63;
        let _nw240 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
        _1649_v63 = _nw240;
        let _1650_v64;
        _1650_v64 = _module.D1.create_DC4(false, _1647_cf2, new BigNumber((_1619_v41).length));
        let _index251 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_1649_v63).length));
        (_1649_v63)[_index251] = (_1650_v64).dtor_cf8;
        let _index252 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_1649_v63).length));
        let _rhs253 = _1646_cf3;
        let _rhs254 = (_this).f10;
        let _rhs255 = _1620_v42;
        let _lhs175 = _1649_v63;
        let _lhs176 = _module.__default.safeIndex(new BigNumber(44), new BigNumber((_1649_v63).length));
        _1620_v42 = _rhs253;
        _1647_cf2 = _rhs254;
        _lhs175[_lhs176] = _rhs255;
        let _1651_v65;
        _1651_v65 = _dafny.Seq.of(_1619_v41);
        let _1652_v66;
        _1652_v66 = _module.D1.create_DC3((_1651_v65)[_module.__default.safeIndex(_1620_v42, new BigNumber((_1651_v65).length))]);
        _1652_v66 = _1652_v66;
        _1635_v56 = _1635_v56;
        _1647_cf2 = _1647_cf2;
      } else {
        let _1653___mcc_h3 = (_source22).cf0;
        let _1654_cf0 = _1653___mcc_h3;
        let _1655_v67;
        let _init43 = ((_1656_v58) => function (_1657_i8) {
          return (_1656_v58).Intersect(_1656_v58);
        })(_1638_v58);
        let _nw241 = Array((new BigNumber(16)).toNumber());
        for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw241.length); _i0_43++) {
          _nw241[_i0_43] = _init43(new BigNumber(_i0_43));
        }
        _1655_v67 = _nw241;
        let _index253 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_1655_v67).length));
        (_1655_v67)[_index253] = (_1638_v58).Union(_1638_v58);
        let _index254 = _module.__default.safeIndex(new BigNumber(723), new BigNumber((_1655_v67).length));
        (_1655_v67)[_index254] = _module.__default.fm15((_this).f10, (_this).f10, (_1620_v42).multipliedBy(new BigNumber((function () {
          let _coll60 = new _dafny.Map();
          for (const _compr_60 of _dafny.IntegerRange(new BigNumber(-192), new BigNumber(379))) {
            let _1658_v68 = _compr_60;
            if (((new BigNumber(-192)).isLessThanOrEqualTo(_1658_v68)) && ((_1658_v68).isLessThan(new BigNumber(379)))) {
              _coll60.push([(_1658_v68).plus(_1620_v42),_1654_cf0]);
            }
          }
          return _coll60;
        }()).length)), new _dafny.CodePoint('k'.codePointAt(0)), globalState);
        let _1659_v69;
        let _nw242 = new _module.C0();
        _nw242.__ctor();
        _1659_v69 = _nw242;
        let _1660_v70;
        let _init44 = ((_1661_cf0, _1662_v62) => function (_1663_i9) {
          return ((_1661_cf0) ? (_1662_v62) : (_1662_v62));
        })(_1654_cf0, _1642_v62);
        let _nw243 = Array((new BigNumber(25)).toNumber());
        for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw243.length); _i0_44++) {
          _nw243[_i0_44] = _init44(new BigNumber(_i0_44));
        }
        _1660_v70 = _nw243;
        let _index255 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((_1660_v70).length));
        (_1660_v70)[_index255] = _1642_v62;
        let _index256 = _module.__default.safeIndex(new BigNumber(108), new BigNumber((_1660_v70).length));
        (_1660_v70)[_index256] = _1642_v62;
        let _index257 = _module.__default.safeIndex(new BigNumber(767), new BigNumber((_1635_v56).length));
        (_1635_v56)[_index257] = (_this).f10;
        let _index258 = _module.__default.safeIndex(new BigNumber(767), new BigNumber((_1635_v56).length));
        (_1635_v56)[_index258] = true;
      }
      r0 = _dafny.Seq.update(_dafny.Seq.Concat(_1619_v41, _dafny.Seq.Create(_module.__default.abs(new BigNumber(363)), function (_1664_i10) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      })), _module.__default.safeIndex(((true) ? (_1620_v42) : ((_dafny.ZERO).minus(_1620_v42))), new BigNumber((_dafny.Seq.Concat(_1619_v41, _dafny.Seq.Create(_module.__default.abs(new BigNumber(363)), function (_1665_i10) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      }))).length)), (_1619_v41)[_module.__default.safeIndex(_1620_v42, new BigNumber((_1619_v41).length))]);
      r1 = (_this).f10;
      let _1666_v71;
      _1666_v71 = _dafny.Set.fromElements((_this).f10);
      let _1667_v72;
      _1667_v72 = _dafny.Seq.of(_1666_v71);
      let _1668_v73;
      _1668_v73 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_1667_v72);
      let _1669_v74;
      _1669_v74 = _dafny.Seq.of(_1620_v42);
      r2 = (((_1668_v73).contains((_this).f10)) ? ((_1668_v73).get((_this).f10)) : (_module.__default.fm16(new BigNumber((_1669_v74).length), _1620_v42, new BigNumber(913), globalState)));
      r3 = _1619_v41;
      return [r0, r1, r2, r3];
    }
    m1(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = false;
      let _1670_v0;
      let _nw244 = new _module.C0();
      _nw244.__ctor();
      _1670_v0 = _nw244;
      let _1671_v1;
      _1671_v1 = new BigNumber(699);
      let _1672_v2;
      _1672_v2 = _module.D1.create_DC4((_this).f10, (_this).f10, _1671_v1);
      r1 = function (_source23) {
        if (_source23.is_DC3) {
          let _1673___mcc_h0 = (_source23).cf5;
          let _1674_cf5 = _1673___mcc_h0;
          return (_this).f10;
        } else if (_source23.is_DC4) {
          let _1675___mcc_h1 = (_source23).cf6;
          let _1676___mcc_h2 = (_source23).cf7;
          let _1677___mcc_h3 = (_source23).cf8;
          let _1678_cf8 = _1677___mcc_h3;
          let _1679_cf7 = _1676___mcc_h2;
          let _1680_cf6 = _1675___mcc_h1;
          return (_dafny.MultiSet.fromElements(_1680_cf6, (_this).f10, (_this).f10, true)).IsDisjointFrom(_dafny.MultiSet.FromArray(_dafny.Seq.of(true)));
        } else if (_source23.is_DC2) {
          let _1681___mcc_h4 = (_source23).cf4;
          let _1682_cf4 = _1681___mcc_h4;
          return (_this).f10;
        } else {
          let _1683___mcc_h5 = (_source23).cf9;
          let _1684_cf9 = _1683___mcc_h5;
          return (_this).f10;
        }
      }(_module.D1.create_DC5(_1672_v2));
      let _1685_v3;
      _1685_v3 = _dafny.MultiSet.fromElements(_1671_v1, _1671_v1);
      let _1686_v4;
      _1686_v4 = _dafny.Seq.of(_1671_v1);
      let _1687_v5;
      _1687_v5 = _module.D1.create_DC4((_this).f10, (_this).f10, _1671_v1);
      let _1688_v6;
      _1688_v6 = _dafny.Seq.UnicodeFromString("qfpc");
      let _1689_v7;
      _1689_v7 = _dafny.Set.fromElements((_this).f10);
      let _1690_v8;
      _1690_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1671_v1,(_this).f10);
      let _1691_v9;
      let _nw245 = Array((new BigNumber(19)).toNumber());
      _nw245[(_dafny.ZERO).toNumber()] = new BigNumber(-961);
      _nw245[(_dafny.ONE).toNumber()] = _1671_v1;
      _nw245[(new BigNumber(2)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(810), _1671_v1);
      _nw245[(new BigNumber(3)).toNumber()] = new BigNumber(((_1685_v3).Difference(_1685_v3)).cardinality());
      _nw245[(new BigNumber(4)).toNumber()] = new BigNumber(139);
      _nw245[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_1686_v4, _1686_v4)).length);
      _nw245[(new BigNumber(6)).toNumber()] = (new BigNumber(-855)).plus((_1687_v5).dtor_cf8);
      _nw245[(new BigNumber(7)).toNumber()] = new BigNumber((_1688_v6).length);
      _nw245[(new BigNumber(8)).toNumber()] = _1671_v1;
      _nw245[(new BigNumber(9)).toNumber()] = new BigNumber((_1689_v7).length);
      _nw245[(new BigNumber(10)).toNumber()] = (new BigNumber(-393)).minus(_1671_v1);
      _nw245[(new BigNumber(11)).toNumber()] = _1671_v1;
      _nw245[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_1671_v1, _1671_v1));
      _nw245[(new BigNumber(13)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(869), _1671_v1);
      _nw245[(new BigNumber(14)).toNumber()] = _1671_v1;
      _nw245[(new BigNumber(15)).toNumber()] = _1671_v1;
      _nw245[(new BigNumber(16)).toNumber()] = _1671_v1;
      _nw245[(new BigNumber(17)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_1688_v6, _module.__default.safeIndex(new BigNumber((_1690_v8).length), new BigNumber((_1688_v6).length)), p0), _1688_v6)).length);
      _nw245[(new BigNumber(18)).toNumber()] = _1671_v1;
      _1691_v9 = _nw245;
      for (const _guard_loop_7 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1691_v9).length))) {
        let _1692_i0 = _guard_loop_7;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1692_i0)) && ((_1692_i0).isLessThan(new BigNumber((_1691_v9).length))))) {
          (_1691_v9)[(_1692_i0)] = (_1692_i0).plus(_1671_v1);
        }
      }
      let _hi4 = (_dafny.ZERO).minus(_1671_v1);
      for (let _1693_i1 = (_1671_v1).minus(_1671_v1); _1693_i1.isLessThan(_hi4); _1693_i1 = _1693_i1.plus(_dafny.ONE)) {
        (globalState).f8 = true;
        let _1694_v10;
        let _nw246 = new _module.C0();
        _nw246.__ctor();
        _1694_v10 = _nw246;
        let _1695_v11;
        _1695_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1691_v9,_1688_v6);
        let _1696_v12;
        let _nw247 = Array((new BigNumber(25)).toNumber());
        _nw247[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_1688_v6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(877)), function (_1697_i2) {
          return new _dafny.CodePoint('x'.codePointAt(0));
        }));
        _nw247[(_dafny.ONE).toNumber()] = _1688_v6;
        _nw247[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("amb"), _1688_v6);
        _nw247[(new BigNumber(3)).toNumber()] = _1688_v6;
        _nw247[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("xmfr"), _dafny.Seq.update(_1688_v6, _module.__default.safeIndex(_1671_v1, new BigNumber((_1688_v6).length)), p0));
        _nw247[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("bk");
        _nw247[(new BigNumber(6)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_1688_v6, _dafny.Seq.UnicodeFromString("bac")), _module.__default.safeIndex(_1671_v1, new BigNumber((_dafny.Seq.Concat(_1688_v6, _dafny.Seq.UnicodeFromString("bac"))).length)), p0);
        _nw247[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_1688_v6, _1688_v6);
        _nw247[(new BigNumber(8)).toNumber()] = _1688_v6;
        _nw247[(new BigNumber(9)).toNumber()] = _1688_v6;
        _nw247[(new BigNumber(10)).toNumber()] = _1688_v6;
        _nw247[(new BigNumber(11)).toNumber()] = _1688_v6;
        _nw247[(new BigNumber(12)).toNumber()] = _1688_v6;
        _nw247[(new BigNumber(13)).toNumber()] = _1688_v6;
        _nw247[(new BigNumber(14)).toNumber()] = _1688_v6;
        _nw247[(new BigNumber(15)).toNumber()] = (((_1695_v11).contains(_1691_v9)) ? ((_1695_v11).get(_1691_v9)) : (_1688_v6));
        _nw247[(new BigNumber(16)).toNumber()] = _module.__default.fm13(new BigNumber(592), globalState);
        _nw247[(new BigNumber(17)).toNumber()] = _dafny.Seq.Concat(_1688_v6, _1688_v6);
        _nw247[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_1688_v6, _1688_v6);
        _nw247[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_1688_v6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(644)), ((_1698_v6, _1699_v1) => function (_1700_i3) {
          return (_1698_v6)[_module.__default.safeIndex(new BigNumber((_dafny.Set.fromElements(_1699_v1, _1699_v1)).length), new BigNumber((_1698_v6).length))];
        })(_1688_v6, _1671_v1)));
        _nw247[(new BigNumber(20)).toNumber()] = _1688_v6;
        _nw247[(new BigNumber(21)).toNumber()] = _dafny.Seq.Concat(_1688_v6, _1688_v6);
        _nw247[(new BigNumber(22)).toNumber()] = _dafny.Seq.Concat(_1688_v6, _dafny.Seq.UnicodeFromString("om"));
        _nw247[(new BigNumber(23)).toNumber()] = _dafny.Seq.UnicodeFromString("lsvc");
        _nw247[(new BigNumber(24)).toNumber()] = _dafny.Seq.UnicodeFromString("ldunu");
        _1696_v12 = _nw247;
        let _index259 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_1696_v12).length));
        (_1696_v12)[_index259] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(786)), ((_1701_p0) => function (_1702_i4) {
          return _1701_p0;
        })(p0));
        let _index260 = _module.__default.safeIndex(new BigNumber(757), new BigNumber((_1691_v9).length));
        (_1691_v9)[_index260] = _1693_i1;
        let _index261 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_1696_v12).length));
        let _index262 = _module.__default.safeIndex(new BigNumber(757), new BigNumber((_1691_v9).length));
        let _rhs256 = _1671_v1;
        let _rhs257 = _dafny.Seq.Concat(_1688_v6, _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("gpkthj"), _1688_v6), _module.__default.safeIndex(_1671_v1, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("gpkthj"), _1688_v6)).length)), new _dafny.CodePoint('i'.codePointAt(0))));
        let _rhs258 = (_this).f10;
        let _rhs259 = (new BigNumber(-553)).multipliedBy((_1693_i1).multipliedBy(new BigNumber(583)));
        let _lhs177 = _1696_v12;
        let _lhs178 = _module.__default.safeIndex(new BigNumber(106), new BigNumber((_1696_v12).length));
        let _lhs179 = _1691_v9;
        let _lhs180 = _module.__default.safeIndex(new BigNumber(757), new BigNumber((_1691_v9).length));
        _1671_v1 = _rhs256;
        _lhs177[_lhs178] = _rhs257;
        r1 = _rhs258;
        _lhs179[_lhs180] = _rhs259;
        let _1703_v13;
        let _nw248 = new _module.C0();
        _nw248.__ctor();
        _1703_v13 = _nw248;
      }
      let _1704_v14;
      _1704_v14 = new _dafny.CodePoint('j'.codePointAt(0));
      _1704_v14 = _1704_v14;
      let _1705_v15;
      _1705_v15 = _module.D1.create_DC2(_dafny.Seq.update(_1686_v4, _module.__default.safeIndex(_1671_v1, new BigNumber((_1686_v4).length)), (_dafny.ZERO).minus(_1671_v1)));
      let _source24 = _1705_v15;
      if (_source24.is_DC3) {
        let _1706___mcc_h6 = (_source24).cf5;
        let _1707_cf5 = _1706___mcc_h6;
        let _1708_v16;
        _1708_v16 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('d'.codePointAt(0)),(_this).f10);
        let _1709_v17;
        _1709_v17 = _dafny.Seq.of((_this).f10);
        _1686_v4 = ((((((_this).f9).contains((_this).f10)) ? (((_this).f9).get((_this).f10)) : (!((((_1708_v16).contains(_1704_v14)) ? ((_1708_v16).get(_1704_v14)) : ((_1709_v17)[_module.__default.safeIndex(_1671_v1, new BigNumber((_1709_v17).length))])))))) ? (_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(149)), ((_1710_v1) => function (_1711_i5) {
          return (_dafny.ZERO).minus(_1710_v1);
        })(_1671_v1)), _dafny.Seq.of(new BigNumber((_dafny.Seq.of((_this).f10, (_this).f10)).length)))) : (_1686_v4));
        let _index263 = _module.__default.safeIndex(new BigNumber(393), new BigNumber((_1691_v9).length));
        (_1691_v9)[_index263] = _1671_v1;
        let _index264 = _module.__default.safeIndex(new BigNumber(393), new BigNumber((_1691_v9).length));
        (_1691_v9)[_index264] = (_dafny.ZERO).minus((_dafny.ZERO).minus(_1671_v1));
        _1690_v8 = _1690_v8;
        let _1712_v18;
        let _nw249 = Array((new BigNumber(4)).toNumber());
        _nw249[(_dafny.ZERO).toNumber()] = (_this).f10;
        _nw249[(_dafny.ONE).toNumber()] = (_this).f10;
        _nw249[(new BigNumber(2)).toNumber()] = (((_this).f10) ? (true) : ((_this).f10));
        _nw249[(new BigNumber(3)).toNumber()] = (new BigNumber(-255)).isLessThanOrEqualTo(new BigNumber(528));
        _1712_v18 = _nw249;
        let _index265 = _module.__default.safeIndex(new BigNumber(158), new BigNumber((_1712_v18).length));
        (_1712_v18)[_index265] = true;
        let _1713_v19;
        _1713_v19 = _dafny.Set.fromElements(_1671_v1);
        let _index266 = _module.__default.safeIndex(new BigNumber(158), new BigNumber((_1712_v18).length));
        (_1712_v18)[_index266] = (_module.__default.fm17(true, globalState)).IsSubsetOf((_1713_v19).Difference(_1713_v19));
      } else if (_source24.is_DC4) {
        let _1714___mcc_h7 = (_source24).cf6;
        let _1715___mcc_h8 = (_source24).cf7;
        let _1716___mcc_h9 = (_source24).cf8;
        let _1717_cf8 = _1716___mcc_h9;
        let _1718_cf7 = _1715___mcc_h8;
        let _1719_cf6 = _1714___mcc_h7;
        let _1720_v20;
        _1720_v20 = _module.D1.create_DC5(_1672_v2);
        let _1721_v21;
        _1721_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1720_v20,_module.D0.create_DC1((_dafny.ZERO).minus(_1671_v1), _1718_cf7, _1671_v1));
        let _1722_v22;
        _1722_v22 = _module.D0.create_DC0(_1718_cf7);
        let _1723_v23;
        _1723_v23 = _dafny.Seq.of(_1722_v22, _1722_v22);
        let _1724_v24;
        _1724_v24 = _dafny.Map.Empty.slice().updateUnsafe(_1718_cf7,new BigNumber(-583));
        let _1725_v25;
        _1725_v25 = _dafny.Seq.of((_this).f10);
        let _rhs260 = (new BigNumber((_dafny.Seq.Concat(_1723_v23, _1723_v23)).length)).isLessThan(_1671_v1);
        let _rhs261 = _1721_v21;
        let _rhs262 = ((_1719_cf6) ? (_1717_cf8) : ((((_1724_v24).contains((_1725_v25)[_module.__default.safeIndex(_1671_v1, new BigNumber((_1725_v25).length))])) ? ((_1724_v24).get((_1725_v25)[_module.__default.safeIndex(_1671_v1, new BigNumber((_1725_v25).length))])) : (new BigNumber((_1689_v7).length)))));
        _1718_cf7 = _rhs260;
        _1721_v21 = _rhs261;
        _1717_cf8 = _rhs262;
        let _1726_v26;
        let _nw250 = new _module.C0();
        _nw250.__ctor();
        _1726_v26 = _nw250;
        let _1727_v27;
        _1727_v27 = _dafny.Seq.of(_1691_v9);
        _1717_cf8 = (new BigNumber((_1727_v27).length)).plus(_1671_v1);
        let _1728_v28;
        _1728_v28 = _1704_v14;
        let _1729_v29;
        _1729_v29 = _module.D1.create_DC3(_dafny.Seq.update(_1688_v6, _module.__default.safeIndex(_1671_v1, new BigNumber((_1688_v6).length)), (_1728_v28)));
        let _source25 = _1729_v29;
        if (_source25.is_DC3) {
          let _1730___mcc_h12 = (_source25).cf5;
          let _1731_cf5 = _1730___mcc_h12;
          let _1732_v30;
          let _nw251 = new _module.C0();
          _nw251.__ctor();
          _1732_v30 = _nw251;
          _1728_v28 = _1704_v14;
          let _1733_v31;
          let _init45 = function (_1734_i6) {
            return (_this).f10;
          };
          let _nw252 = Array((new BigNumber(26)).toNumber());
          for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw252.length); _i0_45++) {
            _nw252[_i0_45] = _init45(new BigNumber(_i0_45));
          }
          _1733_v31 = _nw252;
          let _index267 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_1733_v31).length));
          (_1733_v31)[_index267] = _1718_cf7;
          let _1735_v32;
          _1735_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1719_cf6,_1691_v9);
          let _index268 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_1733_v31).length));
          let _rhs263 = _1719_cf6;
          let _rhs264 = (((_1735_v32).contains(false)) ? ((_1735_v32).get(false)) : ((_1727_v27)[_module.__default.safeIndex(_1671_v1, new BigNumber((_1727_v27).length))]));
          let _rhs265 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("fnhqhubuo"), _1688_v6)).length);
          let _rhs266 = _module.__default.fm14(globalState);
          let _lhs181 = _1733_v31;
          let _lhs182 = _module.__default.safeIndex(new BigNumber(138), new BigNumber((_1733_v31).length));
          _lhs181[_lhs182] = _rhs263;
          _1691_v9 = _rhs264;
          _1717_cf8 = _rhs265;
          _1704_v14 = _rhs266;
          _1690_v8 = _1690_v8;
        } else if (_source25.is_DC4) {
          let _1736___mcc_h13 = (_source25).cf6;
          let _1737___mcc_h14 = (_source25).cf7;
          let _1738___mcc_h15 = (_source25).cf8;
          let _1739_cf8 = _1738___mcc_h15;
          let _1740_cf7 = _1737___mcc_h14;
          let _1741_cf6 = _1736___mcc_h13;
          _1690_v8 = (_1690_v8).update(_module.__default.safeModuloInt(_1739_cf8, new BigNumber(472)), _1741_cf6);
          let _1742_v33;
          _1742_v33 = _dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(_1724_v24)).length));
          let _1743_v43;
          _1743_v43 = _dafny.Map.Empty.slice().updateUnsafe(_1740_cf7,_1670_v0);
          let _1744_v44;
          _1744_v44 = _dafny.Seq.of(function () {
            let _coll61 = new _dafny.Set();
            for (const _compr_61 of (_1686_v4).Elements) {
              let _1745_v41 = _compr_61;
              if (_dafny.Seq.contains(_1686_v4, _1745_v41)) {
                _coll61.add((_1745_v41).minus(new BigNumber(754)));
              }
            }
            return _coll61;
          }(), _1742_v33, function () {
            let _coll62 = new _dafny.Set();
            for (const _compr_62 of _dafny.IntegerRange(new BigNumber(487), new BigNumber(259))) {
              let _1746_v42 = _compr_62;
              if (((new BigNumber(487)).isLessThanOrEqualTo(_1746_v42)) && ((_1746_v42).isLessThan(new BigNumber(259)))) {
                _coll62.add((_1746_v42).minus(new BigNumber((_1688_v6).length)));
              }
            }
            return _coll62;
          }(), _1742_v33, _dafny.Set.fromElements(_1739_cf8, _1739_cf8, new BigNumber(-918), _1671_v1, new BigNumber((_1743_v43).length)));
          let _1747_v45;
          let _nw253 = Array((new BigNumber(13)).toNumber());
          _nw253[(_dafny.ZERO).toNumber()] = _1742_v33;
          _nw253[(_dafny.ONE).toNumber()] = _1742_v33;
          _nw253[(new BigNumber(2)).toNumber()] = _1742_v33;
          _nw253[(new BigNumber(3)).toNumber()] = function () {
            let _coll63 = new _dafny.Set();
            for (const _compr_63 of _dafny.IntegerRange(new BigNumber(600), new BigNumber(-636))) {
              let _1748_v34 = _compr_63;
              if (((new BigNumber(600)).isLessThanOrEqualTo(_1748_v34)) && ((_1748_v34).isLessThan(new BigNumber(-636)))) {
                _coll63.add(_module.__default.safeModuloInt(_1748_v34, new BigNumber((function () {
                  let _coll64 = new _dafny.Map();
                  for (const _compr_64 of (function () {
                    let _coll65 = new _dafny.Map();
                    for (const _compr_65 of (_dafny.MultiSet.fromElements(_1688_v6)).Elements) {
                      let _1749_v36 = _compr_65;
                      if ((_dafny.MultiSet.fromElements(_1688_v6)).contains(_1749_v36)) {
                        _coll65.push([_1749_v36,new BigNumber((_1688_v6).length)]);
                      }
                    }
                    return _coll65;
                  }()).Keys.Elements) {
                    let _1750_v35 = _compr_64;
                    if ((function () {
                      let _coll66 = new _dafny.Map();
                      for (const _compr_66 of (_dafny.MultiSet.fromElements(_1688_v6)).Elements) {
                        let _1749_v36 = _compr_66;
                        if ((_dafny.MultiSet.fromElements(_1688_v6)).contains(_1749_v36)) {
                          _coll66.push([_1749_v36,new BigNumber((_1688_v6).length)]);
                        }
                      }
                      return _coll66;
                    }()).contains(_1750_v35)) {
                      _coll64.push([_1750_v35,_1741_cf6]);
                    }
                  }
                  return _coll64;
                }()).length)));
              }
            }
            return _coll63;
          }();
          _nw253[(new BigNumber(4)).toNumber()] = function () {
            let _coll67 = new _dafny.Set();
            for (const _compr_67 of (_1685_v3).Elements) {
              let _1751_v37 = _compr_67;
              if ((_1685_v3).contains(_1751_v37)) {
                _coll67.add((_1751_v37).multipliedBy(new BigNumber(485)));
              }
            }
            return _coll67;
          }();
          _nw253[(new BigNumber(5)).toNumber()] = ((_1740_cf7) ? (_1742_v33) : (_1742_v33));
          _nw253[(new BigNumber(6)).toNumber()] = _1742_v33;
          _nw253[(new BigNumber(7)).toNumber()] = _1742_v33;
          _nw253[(new BigNumber(8)).toNumber()] = _1742_v33;
          _nw253[(new BigNumber(9)).toNumber()] = function () {
            let _coll68 = new _dafny.Set();
            for (const _compr_68 of (_1686_v4).Elements) {
              let _1752_v38 = _compr_68;
              if (_dafny.Seq.contains(_1686_v4, _1752_v38)) {
                _coll68.add(_module.__default.safeDivisionInt(_1752_v38, new BigNumber((function () {
                  let _coll69 = new _dafny.Map();
                  for (const _compr_69 of _dafny.IntegerRange(new BigNumber(305), new BigNumber(277))) {
                    let _1753_v39 = _compr_69;
                    if (((new BigNumber(305)).isLessThanOrEqualTo(_1753_v39)) && ((_1753_v39).isLessThan(new BigNumber(277)))) {
                      _coll69.push([_module.__default.safeModuloInt(_1753_v39, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(330)), function (_1755_i7) {
                        return new BigNumber(809);
                      })).length))),function () {
                        let _coll70 = new _dafny.Set();
                        for (const _compr_70 of _dafny.IntegerRange(new BigNumber(390), new BigNumber(962))) {
                          let _1754_v40 = _compr_70;
                          if (((new BigNumber(390)).isLessThanOrEqualTo(_1754_v40)) && ((_1754_v40).isLessThan(new BigNumber(962)))) {
                            _coll70.add((_1754_v40).minus(new BigNumber((_dafny.Seq.UnicodeFromString("h")).length)));
                          }
                        }
                        return _coll70;
                      }()]);
                    }
                  }
                  return _coll69;
                }()).length)));
              }
            }
            return _coll68;
          }();
          _nw253[(new BigNumber(10)).toNumber()] = (_1744_v44)[_module.__default.safeIndex(new BigNumber(307), new BigNumber((_1744_v44).length))];
          _nw253[(new BigNumber(11)).toNumber()] = _1742_v33;
          _nw253[(new BigNumber(12)).toNumber()] = _module.__default.fm17((_this).f10, globalState);
          _1747_v45 = _nw253;
          let _index269 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_1747_v45).length));
          (_1747_v45)[_index269] = _1742_v33;
          let _index270 = _module.__default.safeIndex(new BigNumber(739), new BigNumber((_1747_v45).length));
          (_1747_v45)[_index270] = _1742_v33;
          let _1756_v46;
          _1756_v46 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-407),_1739_cf8);
          let _1757_v47;
          let _nw254 = new _module.C4();
          _nw254.__ctor((_this).f9, _1741_cf6);
          _1757_v47 = _nw254;
          let _1758_v48;
          _1758_v48 = _dafny.Seq.of(_1757_v47);
          let _1759_v49;
          _1759_v49 = _dafny.Map.Empty.slice().updateUnsafe(_1741_cf6,_dafny.MultiSet.FromArray(_dafny.Seq.update(_1758_v48, _module.__default.safeIndex(_1671_v1, new BigNumber((_1758_v48).length)), _1757_v47)));
          _1756_v46 = (_1756_v46).update((_1671_v1).plus(_1739_cf8), ((_dafny.ZERO).minus(_1739_cf8)).multipliedBy(new BigNumber((_1759_v49).length)));
          let _1760_v50;
          _1760_v50 = _module.D6.create_DC17(_1741_cf6, (_this).f10, _1671_v1, _1739_cf8, new BigNumber((_1685_v3).cardinality()));
          let _1761_v51;
          let _nw255 = Array((new BigNumber(2)).toNumber());
          _nw255[(_dafny.ZERO).toNumber()] = (_1757_v47).f10;
          _nw255[(_dafny.ONE).toNumber()] = !(_1741_cf6);
          _1761_v51 = _nw255;
          let _1762_v52;
          _1762_v52 = _dafny.Map.Empty.slice().updateUnsafe(_1671_v1,_1761_v51);
          let _1763_v53;
          _1763_v53 = _dafny.Map.Empty.slice().updateUnsafe(_1762_v52,_1741_cf6);
          let _1764_v54;
          let _nw256 = Array((new BigNumber(21)).toNumber());
          _nw256[(_dafny.ZERO).toNumber()] = _1741_cf6;
          _nw256[(_dafny.ONE).toNumber()] = (_this).f10;
          _nw256[(new BigNumber(2)).toNumber()] = _1741_cf6;
          _nw256[(new BigNumber(3)).toNumber()] = !(_1740_cf7);
          _nw256[(new BigNumber(4)).toNumber()] = (_1760_v50).dtor_cf29;
          _nw256[(new BigNumber(5)).toNumber()] = (_this).f10;
          _nw256[(new BigNumber(6)).toNumber()] = true;
          _nw256[(new BigNumber(7)).toNumber()] = (_this).f10;
          _nw256[(new BigNumber(8)).toNumber()] = false;
          _nw256[(new BigNumber(9)).toNumber()] = _1741_cf6;
          _nw256[(new BigNumber(10)).toNumber()] = (((_1763_v53).contains(_1762_v52)) ? ((_1763_v53).get(_1762_v52)) : (true));
          _nw256[(new BigNumber(11)).toNumber()] = (_this).f10;
          _nw256[(new BigNumber(12)).toNumber()] = (_1757_v47).f10;
          _nw256[(new BigNumber(13)).toNumber()] = false;
          _nw256[(new BigNumber(14)).toNumber()] = false;
          _nw256[(new BigNumber(15)).toNumber()] = _1740_cf7;
          _nw256[(new BigNumber(16)).toNumber()] = (_this).f10;
          _nw256[(new BigNumber(17)).toNumber()] = true;
          _nw256[(new BigNumber(18)).toNumber()] = _1741_cf6;
          _nw256[(new BigNumber(19)).toNumber()] = (_1757_v47).f10;
          _nw256[(new BigNumber(20)).toNumber()] = _1741_cf6;
          _1764_v54 = _nw256;
          let _1765_v55;
          _1765_v55 = _dafny.Map.Empty.slice().updateUnsafe(_1741_cf6,_1764_v54);
          let _1766_v56;
          _1766_v56 = _1765_v55;
          let _1767_v57;
          let _nw257 = Array((new BigNumber(21)).toNumber());
          _nw257[(_dafny.ZERO).toNumber()] = _1766_v56;
          _nw257[(_dafny.ONE).toNumber()] = _1766_v56;
          _nw257[(new BigNumber(2)).toNumber()] = _1766_v56;
          _nw257[(new BigNumber(3)).toNumber()] = _1766_v56;
          _nw257[(new BigNumber(4)).toNumber()] = _1765_v55;
          _nw257[(new BigNumber(5)).toNumber()] = _1766_v56;
          _nw257[(new BigNumber(6)).toNumber()] = _1766_v56;
          _nw257[(new BigNumber(7)).toNumber()] = _1766_v56;
          _nw257[(new BigNumber(8)).toNumber()] = _1766_v56;
          _nw257[(new BigNumber(9)).toNumber()] = _1766_v56;
          _nw257[(new BigNumber(10)).toNumber()] = _1766_v56;
          _nw257[(new BigNumber(11)).toNumber()] = _1766_v56;
          _nw257[(new BigNumber(12)).toNumber()] = _1766_v56;
          _nw257[(new BigNumber(13)).toNumber()] = _1766_v56;
          _nw257[(new BigNumber(14)).toNumber()] = _1766_v56;
          _nw257[(new BigNumber(15)).toNumber()] = _1766_v56;
          _nw257[(new BigNumber(16)).toNumber()] = _1766_v56;
          _nw257[(new BigNumber(17)).toNumber()] = _1766_v56;
          _nw257[(new BigNumber(18)).toNumber()] = _1766_v56;
          _nw257[(new BigNumber(19)).toNumber()] = _1766_v56;
          _nw257[(new BigNumber(20)).toNumber()] = ((_1740_cf7) ? (_1766_v56) : (_1766_v56));
          _1767_v57 = _nw257;
          let _index271 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_1767_v57).length));
          (_1767_v57)[_index271] = _1766_v56;
          let _index272 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_1767_v57).length));
          (_1767_v57)[_index272] = _1766_v56;
        } else if (_source25.is_DC2) {
          let _1768___mcc_h16 = (_source25).cf4;
          let _1769_cf4 = _1768___mcc_h16;
          _1671_v1 = _1717_cf8;
          (globalState).f8 = _1719_cf6;
          let _1770_v58;
          let _nw258 = Array((new BigNumber(8)).toNumber());
          _1770_v58 = _nw258;
          let _index273 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_1770_v58).length));
          (_1770_v58)[_index273] = _1726_v26;
          let _index274 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_1770_v58).length));
          (_1770_v58)[_index274] = _1726_v26;
          let _1771_v59;
          _1771_v59 = _dafny.Map.Empty.slice().updateUnsafe(_1671_v1,_module.__default.fm13(_1671_v1, globalState));
          _1771_v59 = (_1771_v59).update(_1717_cf8, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(51)), function (_1772_i8) {
            return new _dafny.CodePoint('m'.codePointAt(0));
          }), _1688_v6));
        } else {
          let _1773___mcc_h17 = (_source25).cf9;
          let _1774_cf9 = _1773___mcc_h17;
          let _1775_v60;
          _1775_v60 = _module.D3.create_DC7(_1691_v9);
          let _pat_let_tv37 = _1691_v9;
          let _1776_v61;
          let _nw259 = Array((new BigNumber(8)).toNumber());
          _nw259[(_dafny.ZERO).toNumber()] = _1775_v60;
          _nw259[(_dafny.ONE).toNumber()] = function (_pat_let38_0) {
            return function (_1777_dt__update__tmp_h1) {
              return function (_pat_let39_0) {
                return function (_1778_dt__update_hcf11_h0) {
                  return _module.D3.create_DC7(_1778_dt__update_hcf11_h0);
                }(_pat_let39_0);
              }(_pat_let_tv37);
            }(_pat_let38_0);
          }(_1775_v60);
          _nw259[(new BigNumber(2)).toNumber()] = _1775_v60;
          _nw259[(new BigNumber(3)).toNumber()] = _1775_v60;
          _nw259[(new BigNumber(4)).toNumber()] = _1775_v60;
          _nw259[(new BigNumber(5)).toNumber()] = _1775_v60;
          _nw259[(new BigNumber(6)).toNumber()] = _1775_v60;
          _nw259[(new BigNumber(7)).toNumber()] = _1775_v60;
          _1776_v61 = _nw259;
          let _1779_v62;
          _1779_v62 = _dafny.Map.Empty.slice().updateUnsafe(_1776_v61,false);
          r1 = ((_1718_cf7) ? (_1719_cf6) : (!((_1779_v62).equals(_1779_v62))));
          let _1780_v63;
          let _nw260 = Array((_dafny.ONE).toNumber());
          _nw260[(_dafny.ZERO).toNumber()] = _1719_cf6;
          _1780_v63 = _nw260;
          let _index275 = _module.__default.safeIndex(new BigNumber(304), new BigNumber((_1780_v63).length));
          (_1780_v63)[_index275] = !((_this).fm0((_module.D8.create_DC23(_dafny.Seq.of(_1719_cf6))).dtor_cf41, false, _1671_v1, _1718_cf7, globalState));
          let _index276 = _module.__default.safeIndex(new BigNumber(304), new BigNumber((_1780_v63).length));
          (_1780_v63)[_index276] = !(_1718_cf7);
          let _1781_v65;
          _1781_v65 = _dafny.Map.Empty.slice().updateUnsafe(_1690_v8,function () {
            let _coll71 = new _dafny.Map();
            for (const _compr_71 of _dafny.IntegerRange(new BigNumber(-892), new BigNumber(990))) {
              let _1782_v64 = _compr_71;
              if (((new BigNumber(-892)).isLessThanOrEqualTo(_1782_v64)) && ((_1782_v64).isLessThan(new BigNumber(990)))) {
                _coll71.push([_module.__default.safeDivisionInt(_1782_v64, new BigNumber(547)),(_module.D14.create_DC35(_1671_v1, _1719_cf6, p0)).dtor_cf62]);
              }
            }
            return _coll71;
          }());
          _1671_v1 = new BigNumber((_1781_v65).length);
          _1725_v25 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of((_1780_v63)[_module.__default.safeIndex(new BigNumber(304), new BigNumber((_1780_v63).length))], (_this).f10), _1725_v25), _1725_v25);
        }
      } else if (_source24.is_DC2) {
        let _1783___mcc_h10 = (_source24).cf4;
        let _1784_cf4 = _1783___mcc_h10;
        _1671_v1 = _1671_v1;
        let _rhs267 = ((_dafny.ZERO).minus(_1671_v1)).isLessThan((_dafny.ZERO).minus((_this).fm1((_this).f9, globalState)));
        let _rhs268 = true;
        let _rhs269 = _1691_v9;
        let _rhs270 = !(_1671_v1).isEqualTo(new BigNumber((_1690_v8).length));
        let _lhs183 = globalState;
        let _lhs184 = globalState;
        let _lhs185 = globalState;
        _lhs183.f8 = _rhs267;
        _lhs184.f8 = _rhs268;
        _1691_v9 = _rhs269;
        _lhs185.f8 = _rhs270;
        let _1785_v66;
        _1785_v66 = _dafny.Seq.of((_this).f10, false);
        if ((_1785_v66)[_module.__default.safeIndex(_1671_v1, new BigNumber((_1785_v66).length))]) {
          let _1786_v67;
          _1786_v67 = _dafny.Map.Empty.slice().updateUnsafe((((_1690_v8).contains(new BigNumber((_dafny.Seq.of(_1691_v9, _1691_v9)).length))) ? ((_1690_v8).get(new BigNumber((_dafny.Seq.of(_1691_v9, _1691_v9)).length))) : ((_this).f10)),_1671_v1);
          let _1787_v68;
          _1787_v68 = _dafny.Map.Empty.slice().updateUnsafe(_1786_v67,_1785_v66);
          let _rhs271 = !(!(_1787_v68).equals(_1787_v68)) || (false);
          let _rhs272 = (_1690_v8).Merge((_1690_v8).Merge(_1690_v8));
          let _lhs186 = globalState;
          _lhs186.f8 = _rhs271;
          _1690_v8 = _rhs272;
          let _1788_v69;
          let _init46 = ((_1789_p0) => function (_1790_i9) {
            return _1789_p0;
          })(p0);
          let _nw261 = Array((new BigNumber(29)).toNumber());
          for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw261.length); _i0_46++) {
            _nw261[_i0_46] = _init46(new BigNumber(_i0_46));
          }
          _1788_v69 = _nw261;
          let _index277 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_1788_v69).length));
          (_1788_v69)[_index277] = _1704_v14;
          let _1791_v70;
          let _nw262 = new _module.C3();
          _nw262.__ctor();
          _1791_v70 = _nw262;
          let _1792_v71;
          _1792_v71 = _dafny.Map.Empty.slice().updateUnsafe(_1671_v1,_1791_v70);
          let _1793_v72;
          let _nw263 = new _module.C5();
          _nw263.__ctor(_dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_this).f10), (_this).f10);
          _1793_v72 = _nw263;
          let _1794_v73;
          _1794_v73 = _dafny.Map.Empty.slice().updateUnsafe(_1793_v72,_1671_v1);
          let _1795_v74;
          _1795_v74 = _dafny.Map.Empty.slice().updateUnsafe((_module.D6.create_DC17(false, (_this).f10, _1671_v1, new BigNumber(-929), new BigNumber((_1792_v71).length))).dtor_cf32,new BigNumber((_1794_v73).length));
          let _index278 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_1788_v69).length));
          let _rhs273 = (_1685_v3).IsProperSubsetOf(_module.__default.fm10((_this).f10, (_dafny.ZERO).minus(_1671_v1), _1688_v6, _1795_v74, globalState));
          let _rhs274 = p0;
          let _lhs187 = globalState;
          let _lhs188 = _1788_v69;
          let _lhs189 = _module.__default.safeIndex(new BigNumber(88), new BigNumber((_1788_v69).length));
          _lhs187.f8 = _rhs273;
          _lhs188[_lhs189] = _rhs274;
          let _1796_v75;
          _1796_v75 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_1784_cf4);
          _1671_v1 = new BigNumber((((_1796_v75).Merge(_1796_v75)).Merge((_1796_v75).Merge(_1796_v75))).length);
          let _1797_v76;
          let _nw264 = new _module.C3();
          _nw264.__ctor();
          _1797_v76 = _nw264;
          let _1798_v77;
          _1798_v77 = _dafny.MultiSet.fromElements((_this).f10, false);
          let _1799_v78;
          _1799_v78 = _module.D6.create_DC17((_this).f10, (_this).f10, (((_1798_v77).contains((_this).f10)) ? ((_1798_v77).get((_this).f10)) : (new BigNumber(667))), _1671_v1, _1671_v1);
          let _1800_v79;
          _1800_v79 = _module.D3.create_DC8((_this).f10, (_this).f10);
          let _pat_let_tv38 = _1797_v76;
          let _pat_let_tv39 = _1800_v79;
          let _pat_let_tv40 = globalState;
          let _pat_let_tv41 = globalState;
          let _pat_let_tv42 = _1671_v1;
          let _1801_v80;
          _1801_v80 = _dafny.Seq.of(_1786_v67, _dafny.Map.Empty.slice().updateUnsafe(false,(function (_pat_let40_0) {
            return function (_1802_dt__update__tmp_h2) {
              return function (_pat_let41_0) {
                return function (_1803_dt__update_hcf33_h0) {
                  return function (_pat_let42_0) {
                    return function (_1804_dt__update_hcf32_h0) {
                      return _module.D6.create_DC17((_1802_dt__update__tmp_h2).dtor_cf29, (_1802_dt__update__tmp_h2).dtor_cf30, (_1802_dt__update__tmp_h2).dtor_cf31, _1804_dt__update_hcf32_h0, _1803_dt__update_hcf33_h0);
                    }(_pat_let42_0);
                  }(_pat_let_tv42);
                }(_pat_let41_0);
              }(new BigNumber((_module.__default.fm46((_this).f10, (_pat_let_tv38).fm27((_this).f10, (_this).f10, _pat_let_tv39, _pat_let_tv40), (_this).f10, _pat_let_tv41)).cardinality()));
            }(_pat_let40_0);
          }(_1799_v78)).dtor_cf31), _1786_v67);
          _1801_v80 = _dafny.Seq.of((_1786_v67).Merge(_1786_v67));
        } else {
          let _1805_v81;
          let _nw265 = new _module.C0();
          _nw265.__ctor();
          _1805_v81 = _nw265;
          let _1806_v82;
          let _init47 = ((_1807_cf4) => function (_1808_i10) {
            return _1807_cf4;
          })(_1784_cf4);
          let _nw266 = Array((new BigNumber(16)).toNumber());
          for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw266.length); _i0_47++) {
            _nw266[_i0_47] = _init47(new BigNumber(_i0_47));
          }
          _1806_v82 = _nw266;
          let _index279 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_1806_v82).length));
          (_1806_v82)[_index279] = _dafny.Seq.of(_1671_v1);
          let _index280 = _module.__default.safeIndex(new BigNumber(941), new BigNumber((_1806_v82).length));
          (_1806_v82)[_index280] = _dafny.Seq.Concat(_dafny.Seq.update(_1784_cf4, _module.__default.safeIndex(new BigNumber(591), new BigNumber((_1784_cf4).length)), _1671_v1), _module.__default.fm37(globalState));
          _1704_v14 = _1704_v14;
          let _1809_v83;
          let _nw267 = Array((new BigNumber(13)).toNumber()).fill([]);
          _1809_v83 = _nw267;
          let _1810_v84;
          _1810_v84 = _module.D17.create_DC43(_1809_v83);
          let _1811_v85;
          let _nw268 = Array((new BigNumber(16)).toNumber());
          _nw268[(_dafny.ZERO).toNumber()] = _1809_v83;
          _nw268[(_dafny.ONE).toNumber()] = _1809_v83;
          _nw268[(new BigNumber(2)).toNumber()] = _1809_v83;
          _nw268[(new BigNumber(3)).toNumber()] = _1809_v83;
          _nw268[(new BigNumber(4)).toNumber()] = _1809_v83;
          _nw268[(new BigNumber(5)).toNumber()] = (_1810_v84).dtor_cf83;
          _nw268[(new BigNumber(6)).toNumber()] = _1809_v83;
          _nw268[(new BigNumber(7)).toNumber()] = _1809_v83;
          _nw268[(new BigNumber(8)).toNumber()] = _1809_v83;
          _nw268[(new BigNumber(9)).toNumber()] = _1809_v83;
          _nw268[(new BigNumber(10)).toNumber()] = _1809_v83;
          _nw268[(new BigNumber(11)).toNumber()] = _1809_v83;
          _nw268[(new BigNumber(12)).toNumber()] = _1809_v83;
          _nw268[(new BigNumber(13)).toNumber()] = _1809_v83;
          _nw268[(new BigNumber(14)).toNumber()] = (_1810_v84).dtor_cf83;
          _nw268[(new BigNumber(15)).toNumber()] = _1809_v83;
          _1811_v85 = _nw268;
          let _index281 = _module.__default.safeIndex(new BigNumber(874), new BigNumber((_1811_v85).length));
          (_1811_v85)[_index281] = _1809_v83;
          let _index282 = _module.__default.safeIndex(new BigNumber(874), new BigNumber((_1811_v85).length));
          (_1811_v85)[_index282] = _1809_v83;
          let _1812_v86;
          let _init48 = ((_1813_v6, _1814_v1) => function (_1815_i11) {
            return _dafny.Set.fromElements(new BigNumber((_1813_v6).length), new BigNumber(684), _1814_v1);
          })(_1688_v6, _1671_v1);
          let _nw269 = Array((new BigNumber(17)).toNumber());
          for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw269.length); _i0_48++) {
            _nw269[_i0_48] = _init48(new BigNumber(_i0_48));
          }
          _1812_v86 = _nw269;
          let _1816_v87;
          _1816_v87 = _dafny.Set.fromElements(_1671_v1, _1671_v1);
          let _1817_v88;
          _1817_v88 = _module.D4.create_DC11(_1671_v1, _1816_v87, (_this).f10);
          let _1818_v89;
          _1818_v89 = _dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1817_v88).dtor_cf16,new _dafny.CodePoint('p'.codePointAt(0)))).length), new BigNumber((_dafny.Seq.UnicodeFromString("bjnevp")).length));
          let _index283 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_1812_v86).length));
          (_1812_v86)[_index283] = (_1816_v87).Union(_1816_v87);
          let _index284 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_1812_v86).length));
          let _rhs275 = (((_1690_v8).contains(_1671_v1)) ? ((_1690_v8).get(_1671_v1)) : (_dafny.Seq.IsProperPrefixOf(_1688_v6, _1688_v6)));
          let _rhs276 = _module.__default.fm25(new BigNumber((_1688_v6).length), _1671_v1, globalState);
          let _lhs190 = globalState;
          let _lhs191 = _1812_v86;
          let _lhs192 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_1812_v86).length));
          _lhs190.f8 = _rhs275;
          _lhs191[_lhs192] = _rhs276;
        }
        let _1819_v90;
        _1819_v90 = _dafny.Set.fromElements(_1671_v1);
        _1819_v90 = _1819_v90;
      } else {
        let _1820___mcc_h11 = (_source24).cf9;
        let _1821_cf9 = _1820___mcc_h11;
        let _index285 = _module.__default.safeIndex(new BigNumber(863), new BigNumber((_1691_v9).length));
        (_1691_v9)[_index285] = _1671_v1;
        let _1822_v91;
        _1822_v91 = _dafny.Set.fromElements(_1688_v6, _dafny.Seq.UnicodeFromString("yp"), _1688_v6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-425)), ((_1823_v14) => function (_1824_i12) {
          return _1823_v14;
        })(_1704_v14)), _1688_v6);
        let _index286 = _module.__default.safeIndex(new BigNumber(863), new BigNumber((_1691_v9).length));
        let _rhs277 = (new BigNumber((_1822_v91).length)).minus(_1671_v1);
        let _rhs278 = (_1671_v1).plus(_1671_v1);
        let _lhs193 = _1691_v9;
        let _lhs194 = _module.__default.safeIndex(new BigNumber(863), new BigNumber((_1691_v9).length));
        _1671_v1 = _rhs277;
        _lhs193[_lhs194] = _rhs278;
        if ((_this).f10) {
          let _1825_v92;
          _1825_v92 = _module.D8.create_DC23(_dafny.Seq.of((_this).f10, (_this).f10));
          let _1826_v93;
          let _nw270 = Array((new BigNumber(16)).toNumber()).fill(false);
          _1826_v93 = _nw270;
          let _index287 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_1826_v93).length));
          (_1826_v93)[_index287] = (_this).f10;
          let _1827_v94;
          _1827_v94 = _dafny.Seq.of((_this).f10, (_1671_v1).isEqualTo(new BigNumber(91)));
          let _index288 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_1826_v93).length));
          let _index289 = _module.__default.safeIndex(new BigNumber(863), new BigNumber((_1691_v9).length));
          let _rhs279 = _1825_v92;
          let _rhs280 = (_1671_v1).isLessThan(_module.__default.safeDivisionInt(_1671_v1, _1671_v1));
          let _rhs281 = _dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(983)), ((_1828_v1) => function (_1829_i13) {
            return _1828_v1;
          })(_1671_v1)), _1671_v1);
          let _rhs282 = _1671_v1;
          let _rhs283 = (_1827_v94)[_module.__default.safeIndex(new BigNumber((_1688_v6).length), new BigNumber((_1827_v94).length))];
          let _lhs195 = _1826_v93;
          let _lhs196 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_1826_v93).length));
          let _lhs197 = globalState;
          let _lhs198 = _1691_v9;
          let _lhs199 = _module.__default.safeIndex(new BigNumber(863), new BigNumber((_1691_v9).length));
          let _lhs200 = globalState;
          _1825_v92 = _rhs279;
          _lhs195[_lhs196] = _rhs280;
          _lhs197.f8 = _rhs281;
          _lhs198[_lhs199] = _rhs282;
          _lhs200.f8 = _rhs283;
          let _1830_v95;
          _1830_v95 = _module.D14.create_DC35(_1671_v1, (_this).f10, _1704_v14);
          _1704_v14 = (_1830_v95).dtor_cf64;
          (globalState).f8 = !_dafny.areEqual(_1688_v6, _dafny.Seq.Concat(_1688_v6, _1688_v6));
          let _1831_v96;
          let _nw271 = new _module.C3();
          _nw271.__ctor();
          _1831_v96 = _nw271;
          let _1832_v97;
          _1832_v97 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm53(_1671_v1, globalState),_1831_v96);
          let _index290 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_1826_v93).length));
          let _index291 = _module.__default.safeIndex(new BigNumber(863), new BigNumber((_1691_v9).length));
          let _index292 = _module.__default.safeIndex(new BigNumber(863), new BigNumber((_1691_v9).length));
          let _rhs284 = !((_this).f10) || ((_this).f10);
          let _rhs285 = (_module.__default.safeModuloInt((_1691_v9)[_module.__default.safeIndex(new BigNumber(863), new BigNumber((_1691_v9).length))], _1671_v1)).isLessThan((_1691_v9)[_module.__default.safeIndex(new BigNumber(863), new BigNumber((_1691_v9).length))]);
          let _rhs286 = !(_1832_v97).equals(_1832_v97);
          let _rhs287 = ((_1691_v9)[_module.__default.safeIndex(new BigNumber(863), new BigNumber((_1691_v9).length))]).minus(new BigNumber(471));
          let _rhs288 = (_1691_v9)[_module.__default.safeIndex(new BigNumber(863), new BigNumber((_1691_v9).length))];
          let _lhs201 = _1826_v93;
          let _lhs202 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_1826_v93).length));
          let _lhs203 = globalState;
          let _lhs204 = globalState;
          let _lhs205 = _1691_v9;
          let _lhs206 = _module.__default.safeIndex(new BigNumber(863), new BigNumber((_1691_v9).length));
          let _lhs207 = _1691_v9;
          let _lhs208 = _module.__default.safeIndex(new BigNumber(863), new BigNumber((_1691_v9).length));
          _lhs201[_lhs202] = _rhs284;
          _lhs203.f8 = _rhs285;
          _lhs204.f8 = _rhs286;
          _lhs205[_lhs206] = _rhs287;
          _lhs207[_lhs208] = _rhs288;
          let _1833_v98;
          _1833_v98 = _dafny.Map.Empty.slice().updateUnsafe(_1671_v1,(_this).f9);
          let _1834_v99;
          _1834_v99 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_1686_v4);
          let _1835_v100;
          let _nw272 = new _module.C2();
          _nw272.__ctor((_1826_v93)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_1826_v93).length))], _1826_v93, (_dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_1826_v93)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_1826_v93).length))])).Merge((((_1833_v98).contains(new BigNumber(((((_1834_v99).contains(!((_1826_v93)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_1826_v93).length))]))) ? ((_1834_v99).get(!((_1826_v93)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_1826_v93).length))]))) : (_1686_v4))).length))) ? ((_1833_v98).get(new BigNumber(((((_1834_v99).contains(!((_1826_v93)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_1826_v93).length))]))) ? ((_1834_v99).get(!((_1826_v93)[_module.__default.safeIndex(new BigNumber(442), new BigNumber((_1826_v93).length))]))) : (_1686_v4))).length))) : ((_this).f9))), (_this).f10);
          _1835_v100 = _nw272;
        } else {
          let _1836_v101;
          let _nw273 = new _module.C3();
          _nw273.__ctor();
          _1836_v101 = _nw273;
          let _index293 = _module.__default.safeIndex(new BigNumber(863), new BigNumber((_1691_v9).length));
          (_1691_v9)[_index293] = ((_1691_v9)[_module.__default.safeIndex(new BigNumber(863), new BigNumber((_1691_v9).length))]).minus((_1691_v9)[_module.__default.safeIndex(new BigNumber(863), new BigNumber((_1691_v9).length))]);
          _1685_v3 = (_1685_v3).Union((_1685_v3).Union(_1685_v3));
          let _1837_v102;
          let _nw274 = new _module.C4();
          _nw274.__ctor((_this).f9, ((_1691_v9)[_module.__default.safeIndex(new BigNumber(863), new BigNumber((_1691_v9).length))]).isLessThan(_1671_v1));
          _1837_v102 = _nw274;
          _1837_v102 = _1837_v102;
          let _index294 = _module.__default.safeIndex(new BigNumber(863), new BigNumber((_1691_v9).length));
          (_1691_v9)[_index294] = (_1691_v9)[_module.__default.safeIndex(new BigNumber(863), new BigNumber((_1691_v9).length))];
        }
        let _1838_v103;
        let _nw275 = Array((_dafny.ONE).toNumber()).fill(false);
        _1838_v103 = _nw275;
        let _index295 = _module.__default.safeIndex(new BigNumber(734), new BigNumber((_1838_v103).length));
        (_1838_v103)[_index295] = (_this).f10;
        let _index296 = _module.__default.safeIndex(new BigNumber(734), new BigNumber((_1838_v103).length));
        (_1838_v103)[_index296] = (_this).f10;
        let _1839_v104;
        let _init49 = ((_1840_v6) => function (_1841_i14) {
          return (_1841_i14).plus((_dafny.ZERO).minus(new BigNumber((_1840_v6).length)));
        })(_1688_v6);
        let _nw276 = Array((new BigNumber(6)).toNumber());
        for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw276.length); _i0_49++) {
          _nw276[_i0_49] = _init49(new BigNumber(_i0_49));
        }
        _1839_v104 = _nw276;
        let _1842_v105;
        _1842_v105 = _dafny.Set.fromElements(_1839_v104, _1839_v104);
        let _1843_v106;
        _1843_v106 = _dafny.Map.Empty.slice().updateUnsafe(_1842_v105,(_1838_v103)[_module.__default.safeIndex(new BigNumber(734), new BigNumber((_1838_v103).length))]);
        _1843_v106 = (_1843_v106).update((((_1838_v103)[_module.__default.safeIndex(new BigNumber(734), new BigNumber((_1838_v103).length))]) ? (_1842_v105) : (_1842_v105)), !((_1838_v103)[_module.__default.safeIndex(new BigNumber(734), new BigNumber((_1838_v103).length))]));
      }
      let _1844_v107;
      _1844_v107 = _dafny.Map.Empty.slice().updateUnsafe(_1691_v9,_module.__default.fm20(_1671_v1, (_this).f10, (_this).f10, _1671_v1, globalState));
      r0 = _1844_v107;
      r1 = !((((_this).f10) ? (false) : ((_this).f10)));
      return [r0, r1];
    }
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
      this._f9 = _dafny.Map.Empty;
      this._f10 = false;
      this._f11 = _dafny.Set.Empty;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f10() {
      let _this = this;
      return _this._f10;
    };
    __ctor(f11, f9, f10) {
      let _this = this;
      (_this)._f11 = f11;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      return;
    }
    fm0(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (_this).f10;
    };
    fm1(p0, globalState) {
      let _this = this;
      return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of((_this).f10)).length));
    };
    fm2(globalState) {
      let _this = this;
      return (_module.__default.safeDivisionInt((_dafny.ZERO).minus(new BigNumber(((_this).f11).length)), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(73))).length), (_dafny.ZERO).minus(new BigNumber(-305)), (_dafny.ZERO).minus(new BigNumber(-142)))).cardinality()))).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.of((_module.D0.create_DC0(!(!((_this).f10)))).dtor_cf0)).length));
    };
    fm3(p0, p1, p2, p3, globalState) {
      let _this = this;
      return new BigNumber(((_module.D1.create_DC2(_dafny.Seq.of(new BigNumber((((_this).f9).update(true, (_this).f10)).length)))).dtor_cf4).length);
    };
    m0(globalState) {
      let _this = this;
      let r0 = _dafny.Seq.UnicodeFromString("");
      let r1 = false;
      let r2 = _dafny.Seq.of();
      let r3 = _dafny.Seq.UnicodeFromString("");
      let _1845_v0;
      _1845_v0 = new BigNumber(233);
      let _1846_v1;
      _1846_v1 = _dafny.Set.fromElements(_1845_v0, _1845_v0);
      let _1847_v2;
      _1847_v2 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(!((_this).f10),!((_this).f10)));
      let _hi5 = (((_this).f10) ? (new BigNumber((_1847_v2).length)) : ((_dafny.ZERO).minus(_1845_v0)));
      for (let _1848_i0 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_1846_v1).length), new BigNumber(-93))); _1848_i0.isLessThan(_hi5); _1848_i0 = _1848_i0.plus(_dafny.ONE)) {
        let _1849_v3;
        _1849_v3 = _dafny.Seq.UnicodeFromString("bt");
        let _1850_v4;
        _1850_v4 = _module.D0.create_DC1(_1845_v0, (_this).f10, (new BigNumber((_1849_v3).length)).plus(_1845_v0));
        let _1851_v6;
        _1851_v6 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_1848_i0);
        let _1852_v7;
        let _nw277 = Array((new BigNumber(6)).toNumber());
        _nw277[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((function () {
          let _coll72 = new _dafny.Map();
          for (const _compr_72 of _dafny.IntegerRange(new BigNumber(23), new BigNumber(718))) {
            let _1853_v5 = _compr_72;
            if (((new BigNumber(23)).isLessThanOrEqualTo(_1853_v5)) && ((_1853_v5).isLessThan(new BigNumber(718)))) {
              _coll72.push([(_1853_v5).plus(_1845_v0),_1848_i0]);
            }
          }
          return _coll72;
        }()).length)));
        _nw277[(_dafny.ONE).toNumber()] = ((_dafny.ZERO).minus(new BigNumber((_1849_v3).length))).minus((((_1851_v6).contains((_this).f10)) ? ((_1851_v6).get((_this).f10)) : (_1845_v0)));
        _nw277[(new BigNumber(2)).toNumber()] = new BigNumber(169);
        _nw277[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(_1845_v0);
        _nw277[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_1849_v3, _dafny.Seq.UnicodeFromString("mootujt"))).length);
        _nw277[(new BigNumber(5)).toNumber()] = _1845_v0;
        _1852_v7 = _nw277;
        let _1854_v8;
        _1854_v8 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(729)), ((_1855_v3) => function (_1856_i1) {
          return new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_1855_v3, _dafny.Seq.UnicodeFromString("ywcvn")))).cardinality());
        })(_1849_v3)),(_this).f10);
        let _index297 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_1852_v7).length));
        (_1852_v7)[_index297] = new BigNumber(((_1854_v8).Merge(_1854_v8)).length);
        let _1857_v9;
        let _nw278 = Array((new BigNumber(10)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1857_v9 = _nw278;
        let _1858_v10;
        _1858_v10 = _dafny.Seq.of((_this).f10, (_this).f10);
        let _1859_v11;
        _1859_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1857_v9,_dafny.Seq.update(_1858_v10, _module.__default.safeIndex((_dafny.ZERO).minus(_1845_v0), new BigNumber((_1858_v10).length)), (_this).f10));
        let _1860_v12;
        _1860_v12 = new _dafny.CodePoint('c'.codePointAt(0));
        let _1861_v13;
        _1861_v13 = _dafny.Seq.of((_1859_v11).Merge((_1859_v11).update(_1857_v9, _module.__default.fm4(_1860_v12, globalState))));
        let _1862_v14;
        _1862_v14 = _dafny.Seq.of(_dafny.Seq.of(_1848_i0));
        let _index298 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_1852_v7).length));
        let _rhs289 = true;
        let _rhs290 = _1850_v4;
        let _rhs291 = new BigNumber(((_1861_v13)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_1862_v14).length)), new BigNumber((_1861_v13).length))]).length);
        let _lhs209 = globalState;
        let _lhs210 = _1852_v7;
        let _lhs211 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_1852_v7).length));
        _lhs209.f8 = _rhs289;
        _1850_v4 = _rhs290;
        _lhs210[_lhs211] = _rhs291;
        let _1863_v15;
        _1863_v15 = _dafny.Seq.of((_1852_v7)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_1852_v7).length))]);
        let _1864_v16;
        _1864_v16 = _module.D1.create_DC2(_1863_v15);
        let _source26 = _module.D1.create_DC5(_1864_v16);
        if (_source26.is_DC3) {
          let _1865___mcc_h0 = (_source26).cf5;
          let _1866_cf5 = _1865___mcc_h0;
          (globalState).f8 = (_1846_v1).IsSubsetOf((_1846_v1).Union(_1846_v1));
          (globalState).f8 = (_module.__default.safeDivisionInt((_dafny.ZERO).minus(_1845_v0), (_dafny.ZERO).minus((_1863_v15)[_module.__default.safeIndex(_1845_v0, new BigNumber((_1863_v15).length))]))).isLessThanOrEqualTo(_1845_v0);
          let _1867_v17;
          let _nw279 = new _module.C0();
          _nw279.__ctor();
          _1867_v17 = _nw279;
          let _1868_v18;
          let _nw280 = new _module.C0();
          _nw280.__ctor();
          _1868_v18 = _nw280;
        } else if (_source26.is_DC4) {
          let _1869___mcc_h1 = (_source26).cf6;
          let _1870___mcc_h2 = (_source26).cf7;
          let _1871___mcc_h3 = (_source26).cf8;
          let _1872_cf8 = _1871___mcc_h3;
          let _1873_cf7 = _1870___mcc_h2;
          let _1874_cf6 = _1869___mcc_h1;
          let _index299 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_1852_v7).length));
          (_1852_v7)[_index299] = _1845_v0;
          let _1875_v19;
          let _init50 = ((_1876_v3) => function (_1877_i2) {
            return _1876_v3;
          })(_1849_v3);
          let _nw281 = Array((new BigNumber(15)).toNumber());
          for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw281.length); _i0_50++) {
            _nw281[_i0_50] = _init50(new BigNumber(_i0_50));
          }
          _1875_v19 = _nw281;
          let _1878_v20;
          _1878_v20 = _module.D1.create_DC3(_1849_v3);
          let _index300 = _module.__default.safeIndex(new BigNumber(769), new BigNumber((_1875_v19).length));
          (_1875_v19)[_index300] = (_1878_v20).dtor_cf5;
          let _index301 = _module.__default.safeIndex(new BigNumber(769), new BigNumber((_1875_v19).length));
          (_1875_v19)[_index301] = _dafny.Seq.Concat(((!((_this).f10)) ? (_1849_v3) : (_1849_v3)), _1849_v3);
          (globalState).f8 = _1873_cf7;
          let _1879_v21;
          _1879_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1845_v0,_1845_v0);
          let _1880_v22;
          _1880_v22 = _module.D0.create_DC0(!(_1874_cf6));
          let _1881_v23;
          _1881_v23 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_1848_i0).minus(new BigNumber((_1879_v21).length))),_1880_v22);
          _1881_v23 = (_1881_v23).update(new BigNumber(308), function (_pat_let43_0) {
            return function (_1882_dt__update__tmp_h0) {
              return function (_pat_let44_0) {
                return function (_1883_dt__update_hcf0_h0) {
                  return _module.D0.create_DC0(_1883_dt__update_hcf0_h0);
                }(_pat_let44_0);
              }((_this).f10);
            }(_pat_let43_0);
          }(_1880_v22));
        } else if (_source26.is_DC2) {
          let _1884___mcc_h4 = (_source26).cf4;
          let _1885_cf4 = _1884___mcc_h4;
          let _index302 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_1852_v7).length));
          (_1852_v7)[_index302] = _1845_v0;
          let _1886_v24;
          _1886_v24 = _dafny.MultiSet.fromElements(new BigNumber(58));
          let _index303 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_1852_v7).length));
          let _rhs292 = (_1858_v10)[_module.__default.safeIndex(_1845_v0, new BigNumber((_1858_v10).length))];
          let _rhs293 = _1848_i0;
          let _rhs294 = !((_this).f10) || ((new BigNumber((_1858_v10).length)).isLessThan(new BigNumber((_1886_v24).cardinality())));
          let _rhs295 = (_1845_v0).multipliedBy(_1845_v0);
          let _lhs212 = _1852_v7;
          let _lhs213 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_1852_v7).length));
          let _lhs214 = globalState;
          r1 = _rhs292;
          _lhs212[_lhs213] = _rhs293;
          _lhs214.f8 = _rhs294;
          _1845_v0 = _rhs295;
          let _1887_v25;
          _1887_v25 = _module.D1.create_DC4((_this).fm2(globalState), (_1846_v1).equals(_1846_v1), _1848_i0);
          _1887_v25 = (((_this).f10) ? (_1887_v25) : (_1887_v25));
          let _index304 = _module.__default.safeIndex(new BigNumber(812), new BigNumber((_1857_v9).length));
          (_1857_v9)[_index304] = _1860_v12;
          let _1888_v26;
          _1888_v26 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,(_this).f10);
          let _index305 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_1852_v7).length));
          let _index306 = _module.__default.safeIndex(new BigNumber(812), new BigNumber((_1857_v9).length));
          let _index307 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_1852_v7).length));
          let _rhs296 = _module.__default.fm6(_1845_v0, _1845_v0, _1888_v26, globalState);
          let _rhs297 = _module.__default.safeDivisionInt((_1852_v7)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_1852_v7).length))], ((_dafny.ZERO).minus((_1852_v7)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_1852_v7).length))])).multipliedBy(new BigNumber(678)));
          let _rhs298 = _1848_i0;
          let _rhs299 = _module.__default.fm7(_1845_v0, globalState);
          let _rhs300 = (_dafny.ZERO).minus((_1852_v7)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_1852_v7).length))]);
          let _lhs215 = _1852_v7;
          let _lhs216 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_1852_v7).length));
          let _lhs217 = _1857_v9;
          let _lhs218 = _module.__default.safeIndex(new BigNumber(812), new BigNumber((_1857_v9).length));
          let _lhs219 = _1852_v7;
          let _lhs220 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_1852_v7).length));
          r0 = _rhs296;
          _1845_v0 = _rhs297;
          _lhs215[_lhs216] = _rhs298;
          _lhs217[_lhs218] = _rhs299;
          _lhs219[_lhs220] = _rhs300;
        } else {
          let _1889___mcc_h5 = (_source26).cf9;
          let _1890_cf9 = _1889___mcc_h5;
          let _1891_v27;
          _1891_v27 = _module.D1.create_DC5(_1864_v16);
          _1891_v27 = _1891_v27;
          let _1892_v28;
          _1892_v28 = _dafny.MultiSet.fromElements((_this).f10);
          let _1893_v29;
          let _init51 = function (_1894_i3) {
            return (_this).f10;
          };
          let _nw282 = Array((new BigNumber(5)).toNumber());
          for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw282.length); _i0_51++) {
            _nw282[_i0_51] = _init51(new BigNumber(_i0_51));
          }
          _1893_v29 = _nw282;
          let _1895_v30;
          let _nw283 = new _module.C2();
          _nw283.__ctor((_this).f10, _1893_v29, _dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_this).f10), false);
          _1895_v30 = _nw283;
          let _1896_v31;
          _1896_v31 = _dafny.MultiSet.fromElements(_1895_v30);
          let _1897_v32;
          _1897_v32 = _dafny.Map.Empty.slice().updateUnsafe((_1896_v31).update(_1895_v30, _module.__default.abs(new BigNumber((_1858_v10).length))),_1863_v15);
          let _rhs301 = (new BigNumber(((_1897_v32).Merge((_dafny.Map.Empty.slice().updateUnsafe(_1896_v31,_1863_v15)).update(_dafny.MultiSet.fromElements(_1895_v30, _1895_v30), _1863_v15))).length)).plus((_dafny.ZERO).minus((_1852_v7)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_1852_v7).length))]));
          let _rhs302 = (_1892_v28).Intersect(_1892_v28);
          let _rhs303 = _1852_v7;
          let _rhs304 = !((_1895_v30).f10);
          let _lhs221 = globalState;
          _1845_v0 = _rhs301;
          _1892_v28 = _rhs302;
          _1852_v7 = _rhs303;
          _lhs221.f8 = _rhs304;
          let _1898_v33;
          let _nw284 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1898_v33 = _nw284;
          let _1899_v34;
          _1899_v34 = _dafny.Map.Empty.slice().updateUnsafe((_1895_v30).f10,_dafny.Map.Empty.slice().updateUnsafe((_this).f11,!((_1895_v30).f10)));
          let _1900_v35;
          _1900_v35 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,(_this).f10);
          let _index308 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_1898_v33).length));
          (_1898_v33)[_index308] = _dafny.Seq.Concat(_module.__default.fm6(_1845_v0, _1848_i0, (((_1899_v34).contains(true)) ? ((_1899_v34).get(true)) : (_1900_v35)), globalState), _1849_v3);
          let _1901_v36;
          let _nw285 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Map.Empty);
          _1901_v36 = _nw285;
          let _1902_v37;
          _1902_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1860_v12,_1848_i0);
          let _index309 = _module.__default.safeIndex(new BigNumber(710), new BigNumber((_1901_v36).length));
          (_1901_v36)[_index309] = _1902_v37;
          let _1903_v38;
          _1903_v38 = _module.D3.create_DC7(_1852_v7);
          let _1904_v39;
          _1904_v39 = _module.D3.create_DC9(_1903_v38);
          let _1905_v41;
          _1905_v41 = _dafny.Map.Empty.slice().updateUnsafe(_1860_v12,_1860_v12);
          let _1906_v42;
          _1906_v42 = _module.D18.create_DC46(_1905_v41);
          let _index310 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_1898_v33).length));
          let _index311 = _module.__default.safeIndex(new BigNumber(710), new BigNumber((_1901_v36).length));
          let _rhs305 = _1849_v3;
          let _rhs306 = function () {
            let _coll73 = new _dafny.Map();
            for (const _compr_73 of ((_1906_v42).dtor_cf89).Keys.Elements) {
              let _1907_v40 = _compr_73;
              if (((_1906_v42).dtor_cf89).contains(_1907_v40)) {
                _coll73.push([_1907_v40,(_1848_i0).plus(new BigNumber(51))]);
              }
            }
            return _coll73;
          }();
          let _rhs307 = _1904_v39;
          let _rhs308 = (_this).f10;
          let _lhs222 = _1898_v33;
          let _lhs223 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_1898_v33).length));
          let _lhs224 = _1901_v36;
          let _lhs225 = _module.__default.safeIndex(new BigNumber(710), new BigNumber((_1901_v36).length));
          _lhs222[_lhs223] = _rhs305;
          _lhs224[_lhs225] = _rhs306;
          _1904_v39 = _rhs307;
          r1 = _rhs308;
          let _1908_v43;
          _1908_v43 = _1860_v12;
          let _1909_v44;
          _1909_v44 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f10) ? (_1908_v43) : (_1908_v43)),(_1898_v33)[_module.__default.safeIndex(new BigNumber(136), new BigNumber((_1898_v33).length))]);
          _1909_v44 = (_dafny.Map.Empty.slice().updateUnsafe(_1908_v43,(_1898_v33)[_module.__default.safeIndex(new BigNumber(136), new BigNumber((_1898_v33).length))])).Merge(_1909_v44);
        }
        let _1910_v45;
        let _init52 = ((_1911_v10) => function (_1912_i4) {
          return _1911_v10;
        })(_1858_v10);
        let _nw286 = Array((new BigNumber(7)).toNumber());
        for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw286.length); _i0_52++) {
          _nw286[_i0_52] = _init52(new BigNumber(_i0_52));
        }
        _1910_v45 = _nw286;
        let _index312 = _module.__default.safeIndex(new BigNumber(503), new BigNumber((_1910_v45).length));
        (_1910_v45)[_index312] = _1858_v10;
        let _index313 = _module.__default.safeIndex(new BigNumber(503), new BigNumber((_1910_v45).length));
        let _rhs309 = _module.__default.safeModuloInt(_1845_v0, _module.__default.fm33((_this).f10, (_this).f10, _1848_i0, (_this).f10, globalState));
        let _rhs310 = new _dafny.CodePoint('w'.codePointAt(0));
        let _rhs311 = _dafny.Seq.Concat(_1858_v10, _dafny.Seq.of((_this).fm2(globalState), (_this).f10, (_this).f10));
        let _lhs226 = _1910_v45;
        let _lhs227 = _module.__default.safeIndex(new BigNumber(503), new BigNumber((_1910_v45).length));
        _1845_v0 = _rhs309;
        _1860_v12 = _rhs310;
        _lhs226[_lhs227] = _rhs311;
        let _1913_v46;
        let _nw287 = new _module.C0();
        _nw287.__ctor();
        _1913_v46 = _nw287;
      }
      let _1914_v47;
      let _nw288 = new _module.C3();
      _nw288.__ctor();
      _1914_v47 = _nw288;
      let _1915_v48;
      _1915_v48 = _dafny.MultiSet.fromElements(_1914_v47, _1914_v47, _1914_v47);
      let _1916_v49;
      _1916_v49 = _dafny.Map.Empty.slice().updateUnsafe(_1845_v0,_1915_v48);
      let _hi6 = new BigNumber((_1916_v49).length);
      for (let _1917_i5 = _module.__default.safeDivisionInt(_1845_v0, _1845_v0); _1917_i5.isLessThan(_hi6); _1917_i5 = _1917_i5.plus(_dafny.ONE)) {
        r1 = (_this).f10;
        if (((_this).f10) === ((_this).f10)) {
          let _1918_v50;
          let _nw289 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
          _1918_v50 = _nw289;
          let _index314 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_1918_v50).length));
          (_1918_v50)[_index314] = _1917_i5;
          let _index315 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_1918_v50).length));
          (_1918_v50)[_index315] = _1917_i5;
          (globalState).f8 = (_this).fm2(globalState);
          let _index316 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_1918_v50).length));
          (_1918_v50)[_index316] = (_dafny.ZERO).minus(((!(!((_this).f10) || ((_this).f10))) ? (_1917_i5) : ((_1918_v50)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_1918_v50).length))])));
          let _1919_v51;
          _1919_v51 = _dafny.Seq.of((_1918_v50)[_module.__default.safeIndex(new BigNumber(216), new BigNumber((_1918_v50).length))]);
          let _1920_v52;
          _1920_v52 = _dafny.Seq.UnicodeFromString("p");
          _1919_v51 = _dafny.Seq.of((_1845_v0).multipliedBy(new BigNumber((_1920_v52).length)));
          (globalState).f8 = (_this).f10;
        } else {
          let _1921_v53;
          _1921_v53 = new _dafny.CodePoint('x'.codePointAt(0));
          let _1922_v54;
          _1922_v54 = _module.D6.create_DC19((_this).f10, _1921_v53);
          let _pat_let_tv43 = _1921_v53;
          (globalState).f8 = (function (_pat_let45_0) {
            return function (_1923_dt__update__tmp_h1) {
              return function (_pat_let46_0) {
                return function (_1924_dt__update_hcf38_h0) {
                  return _module.D6.create_DC19((_1923_dt__update__tmp_h1).dtor_cf37, _1924_dt__update_hcf38_h0);
                }(_pat_let46_0);
              }(_pat_let_tv43);
            }(_pat_let45_0);
          }(_1922_v54)).dtor_cf37;
          let _1925_v55;
          let _nw290 = new _module.C3();
          _nw290.__ctor();
          _1925_v55 = _nw290;
          let _1926_v56;
          let _nw291 = new _module.C0();
          _nw291.__ctor();
          _1926_v56 = _nw291;
          _1845_v0 = _module.__default.safeDivisionInt(_1845_v0, (_dafny.ZERO).minus((_dafny.ZERO).minus(_1845_v0)));
          let _1927_v57;
          _1927_v57 = _dafny.Map.Empty.slice().updateUnsafe(_1845_v0,_1917_i5);
          let _1928_v58;
          _1928_v58 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_1927_v57);
          let _1929_v59;
          _1929_v59 = _dafny.Seq.UnicodeFromString("whn");
          let _1930_v60;
          _1930_v60 = _dafny.Map.Empty.slice().updateUnsafe(_1845_v0,new BigNumber((_1929_v59).length));
          _1928_v58 = (_1928_v58).update((_this).f10, _1930_v60);
        }
        if ((_this).f10) {
          r1 = (_this).f10;
          let _1931_v61;
          let _init53 = function (_1932_i6) {
            return (_1932_i6).multipliedBy(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-602)), function (_1933_i7) {
              return new _dafny.CodePoint('i'.codePointAt(0));
            })).length));
          };
          let _nw292 = Array((new BigNumber(9)).toNumber());
          for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw292.length); _i0_53++) {
            _nw292[_i0_53] = _init53(new BigNumber(_i0_53));
          }
          _1931_v61 = _nw292;
          let _1934_v62;
          _1934_v62 = _dafny.Seq.of(_1931_v61);
          let _1935_v63;
          _1935_v63 = _module.D12.create_DC32(_1931_v61, _dafny.Seq.update(_1934_v62, _module.__default.safeIndex(new BigNumber(851), new BigNumber((_1934_v62).length)), _1931_v61));
          _1935_v63 = _1935_v63;
          let _1936_v64;
          _1936_v64 = _dafny.MultiSet.fromElements((_this).f10);
          let _1937_v65;
          _1937_v65 = _dafny.MultiSet.fromElements(new BigNumber(((_this).f9).length));
          let _1938_v66;
          _1938_v66 = new _dafny.CodePoint('a'.codePointAt(0));
          let _1939_v67;
          _1939_v67 = _dafny.MultiSet.fromElements(_1938_v66);
          let _1940_v68;
          _1940_v68 = _dafny.Seq.UnicodeFromString("dokjnmwo");
          let _1941_v69;
          _1941_v69 = _dafny.Map.Empty.slice().updateUnsafe(_1917_i5,_1940_v68);
          let _1942_v70;
          _1942_v70 = _dafny.Seq.of(_1940_v68);
          let _1943_v71;
          let _nw293 = Array((new BigNumber(20)).toNumber());
          _nw293[(_dafny.ZERO).toNumber()] = (_this).f10;
          _nw293[(_dafny.ONE).toNumber()] = (_this).f10;
          _nw293[(new BigNumber(2)).toNumber()] = (_1917_i5).isLessThanOrEqualTo(_1845_v0);
          _nw293[(new BigNumber(3)).toNumber()] = (_this).f10;
          _nw293[(new BigNumber(4)).toNumber()] = !((_this).f10) || ((_this).f10);
          _nw293[(new BigNumber(5)).toNumber()] = (_this).f10;
          _nw293[(new BigNumber(6)).toNumber()] = ((((_1936_v64).contains(true)) ? ((_1936_v64).get(true)) : (_1845_v0))).isLessThanOrEqualTo(_1917_i5);
          _nw293[(new BigNumber(7)).toNumber()] = (_this).f10;
          _nw293[(new BigNumber(8)).toNumber()] = (_this).f10;
          _nw293[(new BigNumber(9)).toNumber()] = (_this).f10;
          _nw293[(new BigNumber(10)).toNumber()] = !((_1937_v65).IsProperSubsetOf(_1937_v65));
          _nw293[(new BigNumber(11)).toNumber()] = (_1939_v67).IsProperSubsetOf(_1939_v67);
          _nw293[(new BigNumber(12)).toNumber()] = !_dafny.Seq.contains((((_1941_v69).contains(_1917_i5)) ? ((_1941_v69).get(_1917_i5)) : ((_1942_v70)[_module.__default.safeIndex((_dafny.ZERO).minus(_1845_v0), new BigNumber((_1942_v70).length))])), _1938_v66);
          _nw293[(new BigNumber(13)).toNumber()] = (_this).f10;
          _nw293[(new BigNumber(14)).toNumber()] = (_this).f10;
          _nw293[(new BigNumber(15)).toNumber()] = !(true);
          _nw293[(new BigNumber(16)).toNumber()] = !((_this).f10);
          _nw293[(new BigNumber(17)).toNumber()] = (_this).f10;
          _nw293[(new BigNumber(18)).toNumber()] = true;
          _nw293[(new BigNumber(19)).toNumber()] = ((false) ? ((_this).f10) : ((_this).f10));
          _1943_v71 = _nw293;
          let _index317 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_1943_v71).length));
          (_1943_v71)[_index317] = (_this).f10;
          let _1944_v72;
          _1944_v72 = _module.D12.create_DC31((_this).f10, _1940_v68, _dafny.Seq.Create(_module.__default.abs(new BigNumber(287)), ((_1945_v66) => function (_1946_i8) {
  return _1945_v66;
})(_1938_v66)));
          let _pat_let_tv44 = _1940_v68;
          let _pat_let_tv45 = _1845_v0;
          let _pat_let_tv46 = _1938_v66;
          let _1947_v73;
          let _nw294 = new _module.C2();
          _nw294.__ctor(true, _1943_v71, (_this).f9, (function (_pat_let47_0) {
            return function (_1948_dt__update__tmp_h2) {
              return function (_pat_let48_0) {
                return function (_1949_dt__update_hcf57_h0) {
                  return function (_pat_let49_0) {
                    return function (_1950_dt__update_hcf56_h0) {
                      return _module.D12.create_DC31((_1948_dt__update__tmp_h2).dtor_cf55, _1950_dt__update_hcf56_h0, _1949_dt__update_hcf57_h0);
                    }(_pat_let49_0);
                  }(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("cguqrqn"), _module.__default.safeIndex(_pat_let_tv45, new BigNumber((_dafny.Seq.UnicodeFromString("cguqrqn")).length)), _pat_let_tv46));
                }(_pat_let48_0);
              }(_pat_let_tv44);
            }(_pat_let47_0);
          }(_1944_v72)).dtor_cf55);
          _1947_v73 = _nw294;
          let _1951_v74;
          _1951_v74 = _module.D16.create_DC41((_1947_v73).f10, new BigNumber(563), (_this).f10);
          let _1952_v75;
          _1952_v75 = _dafny.Seq.of(new BigNumber(576), (_1951_v74).dtor_cf78);
          let _1953_v76;
          _1953_v76 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_dafny.Seq.of(_1952_v75, _dafny.Seq.Create(_module.__default.abs(new BigNumber(194)), function (_1954_i9) {
            return new BigNumber((_dafny.Seq.UnicodeFromString("i")).length);
          }))));
          let _1955_v77;
          _1955_v77 = _dafny.MultiSet.fromElements(_1952_v75, _1952_v75);
          let _index318 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_1943_v71).length));
          let _rhs312 = (_1955_v77).IsProperSubsetOf((_1953_v76)[_module.__default.safeIndex((((_1937_v65).contains(_1917_i5)) ? ((_1937_v65).get(_1917_i5)) : (_1845_v0)), new BigNumber((_1953_v76).length))]);
          let _rhs313 = _1947_v73;
          let _lhs228 = _1943_v71;
          let _lhs229 = _module.__default.safeIndex(new BigNumber(298), new BigNumber((_1943_v71).length));
          _lhs228[_lhs229] = _rhs312;
          _1947_v73 = _rhs313;
          let _1956_v78;
          let _nw295 = new _module.C2();
          _nw295.__ctor((_1943_v71)[_module.__default.safeIndex(new BigNumber(298), new BigNumber((_1943_v71).length))], _1943_v71, (_1947_v73).f9, (_1943_v71)[_module.__default.safeIndex(new BigNumber(298), new BigNumber((_1943_v71).length))]);
          _1956_v78 = _nw295;
          let _1957_v79;
          let _nw296 = Array((new BigNumber(18)).toNumber());
          _nw296[(_dafny.ZERO).toNumber()] = _1956_v78;
          _nw296[(_dafny.ONE).toNumber()] = _1956_v78;
          _nw296[(new BigNumber(2)).toNumber()] = _1956_v78;
          _nw296[(new BigNumber(3)).toNumber()] = _1956_v78;
          _nw296[(new BigNumber(4)).toNumber()] = _1956_v78;
          _nw296[(new BigNumber(5)).toNumber()] = _1956_v78;
          _nw296[(new BigNumber(6)).toNumber()] = _1956_v78;
          _nw296[(new BigNumber(7)).toNumber()] = _1956_v78;
          _nw296[(new BigNumber(8)).toNumber()] = ((_1956_v78.f14) ? (_1956_v78) : (_1956_v78));
          _nw296[(new BigNumber(9)).toNumber()] = _1956_v78;
          _nw296[(new BigNumber(10)).toNumber()] = _1956_v78;
          _nw296[(new BigNumber(11)).toNumber()] = _1956_v78;
          _nw296[(new BigNumber(12)).toNumber()] = _1956_v78;
          _nw296[(new BigNumber(13)).toNumber()] = _1956_v78;
          _nw296[(new BigNumber(14)).toNumber()] = _1956_v78;
          _nw296[(new BigNumber(15)).toNumber()] = _1956_v78;
          _nw296[(new BigNumber(16)).toNumber()] = _1956_v78;
          _nw296[(new BigNumber(17)).toNumber()] = _1956_v78;
          _1957_v79 = _nw296;
          let _index319 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_1957_v79).length));
          (_1957_v79)[_index319] = _1956_v78;
          let _index320 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_1957_v79).length));
          (_1957_v79)[_index320] = _1956_v78;
          _1936_v64 = (_1936_v64).Union((_dafny.MultiSet.fromElements(_1956_v78.f14, (_1947_v73).f10)).Intersect(_1936_v64));
        } else {
          r1 = (_this).f10;
          (globalState).f8 = !((_this).f10);
          let _1958_v80;
          _1958_v80 = _dafny.Seq.of(new BigNumber(623));
          let _1959_v81;
          _1959_v81 = _dafny.Seq.of(_1917_i5, new BigNumber((_1958_v80).length), new BigNumber(227));
          let _1960_v82;
          let _nw297 = Array((new BigNumber(8)).toNumber());
          _nw297[(_dafny.ZERO).toNumber()] = (_this).f10;
          _nw297[(_dafny.ONE).toNumber()] = (_this).f10;
          _nw297[(new BigNumber(2)).toNumber()] = (_1845_v0).isEqualTo(_1845_v0);
          _nw297[(new BigNumber(3)).toNumber()] = (_this).f10;
          _nw297[(new BigNumber(4)).toNumber()] = (_this).f10;
          _nw297[(new BigNumber(5)).toNumber()] = (new BigNumber(((_this).f11).length)).isLessThan(new BigNumber((_1958_v80).length));
          _nw297[(new BigNumber(6)).toNumber()] = (_this).f10;
          _nw297[(new BigNumber(7)).toNumber()] = (_this).f10;
          _1960_v82 = _nw297;
          let _index321 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_1960_v82).length));
          (_1960_v82)[_index321] = (_this).f10;
          let _index322 = _module.__default.safeIndex(new BigNumber(176), new BigNumber((_1960_v82).length));
          (_1960_v82)[_index322] = ((_this).f10) || ((_this).f10);
          let _rhs314 = _1917_i5;
          let _rhs315 = _1960_v82;
          let _rhs316 = (_1960_v82)[_module.__default.safeIndex(new BigNumber(176), new BigNumber((_1960_v82).length))];
          let _rhs317 = _1960_v82;
          let _lhs230 = globalState;
          _1845_v0 = _rhs314;
          _1960_v82 = _rhs315;
          _lhs230.f8 = _rhs316;
          _1960_v82 = _rhs317;
          _1845_v0 = new BigNumber((_module.__default.fm54(globalState)).length);
        }
        let _1961_v83;
        _1961_v83 = new _dafny.CodePoint('f'.codePointAt(0));
        let _1962_v84;
        _1962_v84 = _dafny.Seq.of((_this).fm1(_dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_this).f10), globalState));
        let _1963_v85;
        let _nw298 = Array((new BigNumber(4)).toNumber());
        _nw298[(_dafny.ZERO).toNumber()] = (_1962_v84)[_module.__default.safeIndex(_1845_v0, new BigNumber((_1962_v84).length))];
        _nw298[(_dafny.ONE).toNumber()] = _1917_i5;
        _nw298[(new BigNumber(2)).toNumber()] = _1845_v0;
        _nw298[(new BigNumber(3)).toNumber()] = _1845_v0;
        _1963_v85 = _nw298;
        let _1964_v86;
        _1964_v86 = _dafny.Map.Empty.slice().updateUnsafe(_1963_v85,(_this).f10);
        let _1965_v87;
        _1965_v87 = _dafny.Seq.of(_1964_v86);
        let _1966_v88;
        _1966_v88 = _dafny.Map.Empty.slice().updateUnsafe(_1961_v83,(_1965_v87)[_module.__default.safeIndex(_1845_v0, new BigNumber((_1965_v87).length))]);
        let _1967_v89;
        _1967_v89 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_1963_v85);
        let _1968_v90;
        _1968_v90 = _dafny.Seq.of(!((_this).f10), (_this).f10, !((_this).f10), true, (_this).f10);
        let _1969_v91;
        _1969_v91 = _dafny.Seq.UnicodeFromString("uqmwx");
        let _1970_v92;
        _1970_v92 = _dafny.Map.Empty.slice().updateUnsafe(_1917_i5,(_1964_v86).update((((_1967_v89).contains((_this).f10)) ? ((_1967_v89).get((_this).f10)) : (_1963_v85)), (_1968_v90)[_module.__default.safeIndex(new BigNumber((_1969_v91).length), new BigNumber((_1968_v90).length))]));
        let _1971_v93;
        let _nw299 = new _module.C5();
        _nw299.__ctor((_this).f9, (_this).f10);
        _1971_v93 = _nw299;
        let _1972_v94;
        let _nw300 = new _module.C6();
        _nw300.__ctor((_this).f9, true);
        _1972_v94 = _nw300;
        let _1973_v95;
        _1973_v95 = _dafny.Map.Empty.slice().updateUnsafe(_1971_v93,_1972_v94);
        let _1974_v96;
        _1974_v96 = _module.D8.create_DC24((_this).f10, _dafny.Set.fromElements(_1963_v85, _1963_v85), new BigNumber((_1973_v95).length));
        if ((((((_1966_v88).contains(_1961_v83)) ? ((_1966_v88).get(_1961_v83)) : (_1964_v86))).Merge((((_1970_v92).contains((_1974_v96).dtor_cf44)) ? ((_1970_v92).get((_1974_v96).dtor_cf44)) : (_1964_v86)))).equals((_1964_v86).Merge(_1964_v86))) {
          (globalState).f8 = (_this).f10;
          let _1975_v97;
          let _nw301 = Array((new BigNumber(18)).toNumber()).fill(false);
          _1975_v97 = _nw301;
          let _index323 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_1975_v97).length));
          (_1975_v97)[_index323] = !((_this).f10);
          let _1976_v98;
          _1976_v98 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_1917_i5),(_this).f10);
          let _1977_v99;
          _1977_v99 = _module.D16.create_DC42((_this).f10, new BigNumber((_dafny.Seq.UnicodeFromString("im")).length), _1917_i5);
          let _index324 = _module.__default.safeIndex(new BigNumber(212), new BigNumber((_1975_v97).length));
          (_1975_v97)[_index324] = ((((((_1976_v98).contains(_1846_v1)) ? ((_1976_v98).get(_1846_v1)) : ((_this).f10))) ? (_1977_v99) : (_1977_v99))).dtor_cf80;
          let _1978_v101;
          let _nw302 = new _module.C2();
          _nw302.__ctor((_1975_v97)[_module.__default.safeIndex(new BigNumber(212), new BigNumber((_1975_v97).length))], _1975_v97, (_this).f9, (_1968_v90)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).fm3(new BigNumber((function () {
            let _coll74 = new _dafny.Set();
            for (const _compr_74 of _dafny.IntegerRange(new BigNumber(47), new BigNumber(633))) {
              let _1979_v100 = _compr_74;
              if (((new BigNumber(47)).isLessThanOrEqualTo(_1979_v100)) && ((_1979_v100).isLessThan(new BigNumber(633)))) {
                _coll74.add((_1979_v100).plus(_1845_v0));
              }
            }
            return _coll74;
          }()).length), new BigNumber(113), _1961_v83, _1845_v0, globalState)), new BigNumber((_1968_v90).length))]);
          _1978_v101 = _nw302;
          let _1980_v102;
          _1980_v102 = _dafny.Map.Empty.slice().updateUnsafe((_1975_v97)[_module.__default.safeIndex(new BigNumber(212), new BigNumber((_1975_v97).length))],_1978_v101);
          _1980_v102 = _1980_v102;
          _1845_v0 = _1917_i5;
          _1845_v0 = _module.__default.safeDivisionInt(_1917_i5, _1845_v0);
        } else {
          let _1981_v103;
          _1981_v103 = _dafny.MultiSet.fromElements((_this).f10, true, (_this).f10, ((_this).f10) && ((_this).f10));
          _1845_v0 = (((_1981_v103).contains((_this).f10)) ? ((_1981_v103).get((_this).f10)) : (_1845_v0));
          _1963_v85 = _1963_v85;
          let _1982_v104;
          let _nw303 = Array((new BigNumber(22)).toNumber());
          _1982_v104 = _nw303;
          let _1983_v105;
          _1983_v105 = _module.D19.create_DC48(_1982_v104);
          let _1984_v106;
          let _nw304 = Array((new BigNumber(6)).toNumber());
          _nw304[(_dafny.ZERO).toNumber()] = _1982_v104;
          _nw304[(_dafny.ONE).toNumber()] = _1982_v104;
          _nw304[(new BigNumber(2)).toNumber()] = _1982_v104;
          _nw304[(new BigNumber(3)).toNumber()] = _1982_v104;
          _nw304[(new BigNumber(4)).toNumber()] = _1982_v104;
          _nw304[(new BigNumber(5)).toNumber()] = (_1983_v105).dtor_cf92;
          _1984_v106 = _nw304;
          _1984_v106 = _1984_v106;
          let _1985_v107;
          let _nw305 = new _module.C1();
          _nw305.__ctor((_this).f10, _1969_v91);
          _1985_v107 = _nw305;
          let _rhs318 = _1985_v107;
          let _rhs319 = new BigNumber((_dafny.Seq.Concat(_1962_v84, _dafny.Seq.Create(_module.__default.abs(new BigNumber(313)), ((_1986_v84, _1987_v0) => function (_1988_i10) {
            return (_1986_v84)[_module.__default.safeIndex(_1987_v0, new BigNumber((_1986_v84).length))];
          })(_1962_v84, _1845_v0)))).length);
          let _rhs320 = (_this).fm1((_this).f9, globalState);
          let _rhs321 = _dafny.Seq.of(new BigNumber(13));
          _1985_v107 = _rhs318;
          _1845_v0 = _rhs319;
          _1845_v0 = _rhs320;
          _1962_v84 = _rhs321;
          let _index325 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_1963_v85).length));
          (_1963_v85)[_index325] = (_1917_i5).minus(_1845_v0);
          let _index326 = _module.__default.safeIndex(new BigNumber(136), new BigNumber((_1963_v85).length));
          (_1963_v85)[_index326] = (_dafny.ZERO).minus(_1845_v0);
        }
      }
      let _1989_v108;
      _1989_v108 = _dafny.Seq.of((_this).f10);
      let _1990_v109;
      _1990_v109 = _dafny.Seq.of(((_this).fm0(_1989_v108, (_this).f10, _1845_v0, (_this).f10, globalState)) || ((_this).f10));
      let _1991_v110;
      _1991_v110 = _module.D15.create_DC39(!((_this).f10), true);
      let _pat_let_tv47 = _1989_v108;
      let _pat_let_tv48 = _1990_v109;
      let _pat_let_tv49 = _1990_v109;
      _1990_v109 = function (_source27) {
        if (_source27.is_DC38) {
          let _1992___mcc_h6 = (_source27).cf69;
          let _1993___mcc_h7 = (_source27).cf70;
          let _1994___mcc_h8 = (_source27).cf71;
          let _1995___mcc_h9 = (_source27).cf72;
          let _1996___mcc_h10 = (_source27).cf73;
          let _1997_cf73 = _1996___mcc_h10;
          let _1998_cf72 = _1995___mcc_h9;
          let _1999_cf71 = _1994___mcc_h8;
          let _2000_cf70 = _1993___mcc_h7;
          let _2001_cf69 = _1992___mcc_h6;
          return (_module.D8.create_DC23(_pat_let_tv47)).dtor_cf41;
        } else if (_source27.is_DC39) {
          let _2002___mcc_h11 = (_source27).cf74;
          let _2003___mcc_h12 = (_source27).cf75;
          let _2004_cf75 = _2003___mcc_h12;
          let _2005_cf74 = _2002___mcc_h11;
          return _pat_let_tv48;
        } else {
          let _2006___mcc_h13 = (_source27).cf68;
          let _2007_cf68 = _2006___mcc_h13;
          return _pat_let_tv49;
        }
      }((((_this).f10) ? (_1991_v110) : (_module.D15.create_DC39((_this).f10, (_this).f10))));
      let _2008_v111;
      let _nw306 = Array((new BigNumber(29)).toNumber()).fill(false);
      _2008_v111 = _nw306;
      for (const _guard_loop_8 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2008_v111).length))) {
        let _2009_i11 = _guard_loop_8;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2009_i11)) && ((_2009_i11).isLessThan(new BigNumber((_2008_v111).length))))) {
          (_2008_v111)[(_2009_i11)] = ((_this).f10) || ((_this).f10);
        }
      }
      let _2010_i12;
      _2010_i12 = _dafny.ZERO;
      L8: {
        while ((_this).f10) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2010_i12)) {
              break L8;
            }
            _2010_i12 = (_2010_i12).plus(_dafny.ONE);
            let _2011_v112;
            let _nw307 = new _module.C2();
            _nw307.__ctor((_this).f10, _2008_v111, (_this).f9, (_this).f10);
            _2011_v112 = _nw307;
            let _2012_v113;
            _2012_v113 = new _dafny.CodePoint('q'.codePointAt(0));
            _1845_v0 = ((_dafny.ZERO).minus((_1845_v0).multipliedBy(_1845_v0))).plus((((_this).f10) ? ((_this).fm3(new BigNumber((_dafny.Seq.of(_2011_v112)).length), new BigNumber(((_this).f11).length), _2012_v113, _1845_v0, globalState)) : (_1845_v0)));
            let _2013_v114;
            let _nw308 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
            _2013_v114 = _nw308;
            let _2014_v115;
            _2014_v115 = _dafny.Seq.UnicodeFromString("fie");
            let _index327 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_2013_v114).length));
            (_2013_v114)[_index327] = _2014_v115;
            let _2015_v116;
            _2015_v116 = _module.D14.create_DC36((_this).f10, _2014_v115, _dafny.Seq.Create(_module.__default.abs(new BigNumber(425)), ((_2016_v113) => function (_2017_i13) {
  return _2016_v113;
})(_2012_v113)));
            let _index328 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_2013_v114).length));
            (_2013_v114)[_index328] = _dafny.Seq.Concat(_2014_v115, ((true) ? (_module.__default.fm20(_1845_v0, _2011_v112.f14, !((_this).f10), _1845_v0, globalState)) : ((_2015_v116).dtor_cf66)));
            _1845_v0 = _1845_v0;
            if ((_this).f10) {
              _1845_v0 = new BigNumber(609);
              let _2018_v117;
              let _nw309 = new _module.C4();
              _nw309.__ctor((_this).f9, true);
              _2018_v117 = _nw309;
              let _2019_v118;
              _2019_v118 = _dafny.Map.Empty.slice().updateUnsafe(_2011_v112.f14,_2018_v117);
              let _2020_v119;
              _2020_v119 = _dafny.MultiSet.fromElements(_1845_v0, new BigNumber((_2019_v118).length));
              let _index329 = _module.__default.safeIndex(new BigNumber(982), new BigNumber(((_2011_v112).f15).length));
              ((_2011_v112).f15)[_index329] = (_dafny.MultiSet.fromElements(_1845_v0, _1845_v0)).IsProperSubsetOf(_2020_v119);
              let _index330 = _module.__default.safeIndex(new BigNumber(982), new BigNumber(((_2011_v112).f15).length));
              ((_2011_v112).f15)[_index330] = true;
              let _2021_v120;
              let _init54 = ((_2022_v0) => function (_2023_i14) {
                return _module.__default.safeDivisionInt(_2023_i14, _2022_v0);
              })(_1845_v0);
              let _nw310 = Array((new BigNumber(22)).toNumber());
              for (let _i0_54 = 0; _i0_54 < new BigNumber(_nw310.length); _i0_54++) {
                _nw310[_i0_54] = _init54(new BigNumber(_i0_54));
              }
              _2021_v120 = _nw310;
              let _index331 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_2021_v120).length));
              (_2021_v120)[_index331] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(687)), ((_2024_v113) => function (_2025_i15) {
                return _2024_v113;
              })(_2012_v113))).length));
              let _2026_v121;
              let _init55 = ((_2027_v0, _2028_v1, _2029_v112, _2030_v113) => function (_2031_i16) {
                return _dafny.Map.Empty.slice().updateUnsafe(_2027_v0,_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_module.D4.create_DC11(_2027_v0, _2028_v1, _2029_v112.f14), _module.D4.create_DC11(_2027_v0, _2028_v1, ((_2029_v112).f15)[_module.__default.safeIndex(new BigNumber(982), new BigNumber(((_2029_v112).f15).length))]), _module.D4.create_DC11(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(847)), ((_2032_v113) => function (_2033_i17) {
  return _2032_v113;
})(_2030_v113))).length), _2028_v1, true)),_2027_v0));
              })(_1845_v0, _1846_v1, _2011_v112, _2012_v113);
              let _nw311 = Array((new BigNumber(3)).toNumber());
              for (let _i0_55 = 0; _i0_55 < new BigNumber(_nw311.length); _i0_55++) {
                _nw311[_i0_55] = _init55(new BigNumber(_i0_55));
              }
              _2026_v121 = _nw311;
              let _2034_v122;
              _2034_v122 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-656)), ((_2035_v0, _2036_v1, _2037_v112) => function (_2038_i18) {
                return _module.D4.create_DC11(_2035_v0, _2036_v1, _2037_v112.f14);
              })(_1845_v0, _1846_v1, _2011_v112));
              let _2039_v123;
              _2039_v123 = _dafny.Map.Empty.slice().updateUnsafe(_2034_v122,_1845_v0);
              let _2040_v124;
              _2040_v124 = _dafny.Map.Empty.slice().updateUnsafe(_1845_v0,_2039_v123);
              let _index332 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_2026_v121).length));
              (_2026_v121)[_index332] = _2040_v124;
              let _index333 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_2021_v120).length));
              let _index334 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_2026_v121).length));
              let _rhs322 = (_1845_v0).plus(_1845_v0);
              let _rhs323 = _2012_v113;
              let _rhs324 = (_2040_v124).Merge(_2040_v124);
              let _lhs231 = _2021_v120;
              let _lhs232 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_2021_v120).length));
              let _lhs233 = _2026_v121;
              let _lhs234 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_2026_v121).length));
              _lhs231[_lhs232] = _rhs322;
              _2012_v113 = _rhs323;
              _lhs233[_lhs234] = _rhs324;
              let _2041_v125;
              _2041_v125 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(313),_1845_v0);
              let _2042_v126;
              _2042_v126 = _dafny.Seq.of(new BigNumber((_dafny.Seq.of((((_2041_v125).contains((_2021_v120)[_module.__default.safeIndex(new BigNumber(894), new BigNumber((_2021_v120).length))])) ? ((_2041_v125).get((_2021_v120)[_module.__default.safeIndex(new BigNumber(894), new BigNumber((_2021_v120).length))])) : (_1845_v0)), _1845_v0)).length));
              let _2043_v127;
              _2043_v127 = _dafny.Map.Empty.slice().updateUnsafe((_2042_v126)[_module.__default.safeIndex(new BigNumber((_1989_v108).length), new BigNumber((_2042_v126).length))],true);
              (_2011_v112).f14 = !((_2043_v127).Merge(_2043_v127)).equals(_2043_v127);
              let _2044_v128;
              _2044_v128 = _module.D6.create_DC19((_this).f10, _2012_v113);
              let _2045_v129;
              _2045_v129 = _dafny.MultiSet.fromElements(_2044_v128, _2044_v128, _2044_v128);
              let _2046_v130;
              _2046_v130 = _dafny.Seq.of(_2045_v129);
              let _rhs325 = (_2013_v114)[_module.__default.safeIndex(new BigNumber(571), new BigNumber((_2013_v114).length))];
              let _rhs326 = (_2046_v130)[_module.__default.safeIndex(_1845_v0, new BigNumber((_2046_v130).length))];
              _2014_v115 = _rhs325;
              _2045_v129 = _rhs326;
            } else {
              let _2047_v131;
              let _nw312 = Array((new BigNumber(23)).toNumber()).fill(_module.D1.Default());
              _2047_v131 = _nw312;
              let _2048_v132;
              _2048_v132 = _dafny.Set.fromElements(_2047_v131, _2047_v131, _2047_v131, _2047_v131);
              _2048_v132 = _2048_v132;
              let _2049_v133;
              let _nw313 = new _module.C5();
              _nw313.__ctor(_module.__default.fm48(_2011_v112.f14, globalState), _2011_v112.f14);
              _2049_v133 = _nw313;
              let _2050_v134;
              let _nw314 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
              _2050_v134 = _nw314;
              _2050_v134 = _2050_v134;
              let _2051_v135;
              _2051_v135 = _dafny.Set.fromElements((_2013_v114)[_module.__default.safeIndex(new BigNumber(571), new BigNumber((_2013_v114).length))]);
              _2051_v135 = (_2051_v135).Difference(_2051_v135);
              let _2052_v136;
              let _nw315 = Array((new BigNumber(15)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
              _2052_v136 = _nw315;
              let _index335 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_2052_v136).length));
              (_2052_v136)[_index335] = new _dafny.CodePoint('a'.codePointAt(0));
              let _index336 = _module.__default.safeIndex(new BigNumber(781), new BigNumber((_2052_v136).length));
              (_2052_v136)[_index336] = _2012_v113;
            }
          }
        }
      }
      r1 = (_this).f10;
      let _2053_v137;
      _2053_v137 = _dafny.Seq.UnicodeFromString("y");
      r0 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2053_v137, _2053_v137), _2053_v137);
      let _2054_v138;
      _2054_v138 = _dafny.Seq.of(_1845_v0);
      let _2055_v139;
      _2055_v139 = new _dafny.CodePoint('f'.codePointAt(0));
      let _2056_v140;
      _2056_v140 = _dafny.Map.Empty.slice().updateUnsafe(_1845_v0,(_this).f10);
      let _2057_v141;
      _2057_v141 = _module.D14.create_DC35(new BigNumber((_2056_v140).length), (_this).f10, _2055_v139);
      r1 = !(((_dafny.ZERO).minus(_1845_v0)).plus((_this).fm3(new BigNumber((_2054_v138).length), new BigNumber((_1990_v109).length), _2055_v139, _1845_v0, globalState))).isEqualTo((_dafny.ZERO).minus((_2057_v141).dtor_cf62));
      let _2058_v142;
      _2058_v142 = _module.D7.create_DC21((_this).f11);
      let _2059_v143;
      _2059_v143 = _dafny.Seq.of(_dafny.Set.fromElements((_this).f10, (_this).f10, (_this).f10), (_this).f11, (function (_pat_let50_0) {
        return function (_2060_dt__update__tmp_h3) {
          return function (_pat_let51_0) {
            return function (_2061_dt__update_hcf40_h0) {
              return _module.D7.create_DC21(_2061_dt__update_hcf40_h0);
            }(_pat_let51_0);
          }(_dafny.Set.fromElements((_this).f10));
        }(_pat_let50_0);
      }(_2058_v142)).dtor_cf40);
      let _2062_v144;
      _2062_v144 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(375)), function (_2063_i19) {
        return (_this).f11;
      }), _dafny.Seq.update(_dafny.Seq.Concat(_2059_v143, _2059_v143), _module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,(_this).f9)).length), new BigNumber((_dafny.Seq.Concat(_2059_v143, _2059_v143)).length)), (_this).f11), _2059_v143, _dafny.Seq.of((_this).f11, (_this).f11), _dafny.Seq.of((_this).f11, _dafny.Set.fromElements((_this).f10), (_this).f11, (_this).f11));
      r2 = (_2062_v144)[_module.__default.safeIndex(_module.__default.safeModuloInt(_1845_v0, new BigNumber((_2054_v138).length)), new BigNumber((_2062_v144).length))];
      r3 = _dafny.Seq.UnicodeFromString("xjnxgcyay");
      return [r0, r1, r2, r3];
    }
    m1(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = false;
      let _2064_v0;
      let _nw316 = Array((new BigNumber(2)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _2064_v0 = _nw316;
      for (const _guard_loop_9 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2064_v0).length))) {
        let _2065_i0 = _guard_loop_9;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2065_i0)) && ((_2065_i0).isLessThan(new BigNumber((_2064_v0).length))))) {
          (_2064_v0)[(_2065_i0)] = p0;
        }
      }
      if ((_this).f10) {
        let _2066_v1;
        let _nw317 = Array((new BigNumber(8)).toNumber()).fill(false);
        _2066_v1 = _nw317;
        _2066_v1 = _2066_v1;
        let _2067_v2;
        _2067_v2 = new _dafny.CodePoint('y'.codePointAt(0));
        _2067_v2 = p0;
        r1 = (((_this).f10) ? ((_this).f10) : (true));
        let _2068_v3;
        let _nw318 = Array((new BigNumber(6)).toNumber()).fill([]);
        _2068_v3 = _nw318;
        _2068_v3 = _2068_v3;
        let _2069_v4;
        _2069_v4 = new BigNumber(128);
        _2069_v4 = ((!(_2069_v4).isEqualTo(_2069_v4)) ? (new BigNumber(662)) : (_2069_v4));
      } else {
        let _2070_v5;
        _2070_v5 = new BigNumber(817);
        _2070_v5 = _2070_v5;
        if ((_this).f10) {
          let _2071_v6;
          let _init56 = function (_2072_i1) {
            return (_this).f10;
          };
          let _nw319 = Array((new BigNumber(6)).toNumber());
          for (let _i0_56 = 0; _i0_56 < new BigNumber(_nw319.length); _i0_56++) {
            _nw319[_i0_56] = _init56(new BigNumber(_i0_56));
          }
          _2071_v6 = _nw319;
          let _2073_v7;
          let _nw320 = Array((new BigNumber(2)).toNumber());
          _nw320[(_dafny.ZERO).toNumber()] = _2071_v6;
          _nw320[(_dafny.ONE).toNumber()] = _2071_v6;
          _2073_v7 = _nw320;
          let _index337 = _module.__default.safeIndex(new BigNumber(517), new BigNumber((_2071_v6).length));
          (_2071_v6)[_index337] = (_this).f10;
          let _2074_v8;
          _2074_v8 = _dafny.Seq.UnicodeFromString("gfnvwqdiy");
          let _2075_v9;
          _2075_v9 = _module.D14.create_DC36(!((_this).f10), _2074_v8, _2074_v8);
          let _index338 = _module.__default.safeIndex(new BigNumber(517), new BigNumber((_2071_v6).length));
          let _rhs327 = _2073_v7;
          let _rhs328 = (_2075_v9).dtor_cf65;
          let _rhs329 = !((_this).f10);
          let _rhs330 = new BigNumber(-820);
          let _lhs235 = _2071_v6;
          let _lhs236 = _module.__default.safeIndex(new BigNumber(517), new BigNumber((_2071_v6).length));
          let _lhs237 = globalState;
          _2073_v7 = _rhs327;
          _lhs235[_lhs236] = _rhs328;
          _lhs237.f8 = _rhs329;
          _2070_v5 = _rhs330;
          _2070_v5 = new BigNumber(557);
          r1 = (_this).f10;
          let _2076_v10;
          _2076_v10 = _dafny.Seq.of((_2071_v6)[_module.__default.safeIndex(new BigNumber(517), new BigNumber((_2071_v6).length))]);
          (globalState).f8 = (_dafny.Map.Empty.slice().updateUnsafe(_2070_v5,_dafny.Seq.update(_2076_v10, _module.__default.safeIndex(_2070_v5, new BigNumber((_2076_v10).length)), (_this).fm2(globalState)))).equals(_dafny.Map.Empty.slice().updateUnsafe((_this).fm1((_this).f9, globalState),_dafny.Seq.update(_2076_v10, _module.__default.safeIndex(_2070_v5, new BigNumber((_2076_v10).length)), (_this).f10)));
          (globalState).f8 = (_module.D15.create_DC39((_this).f10, true)).dtor_cf74;
        } else {
          _2070_v5 = _2070_v5;
          (globalState).f8 = !((_this).f10);
          (globalState).f8 = (_this).f10;
          let _2077_v11;
          let _nw321 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
          _2077_v11 = _nw321;
          let _index339 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_2077_v11).length));
          (_2077_v11)[_index339] = (_dafny.ZERO).minus((_2070_v5).minus(_2070_v5));
          let _index340 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_2077_v11).length));
          (_2077_v11)[_index340] = new BigNumber(224);
          let _2078_v12;
          _2078_v12 = _dafny.Seq.UnicodeFromString("uppopviv");
          (globalState).f8 = !((_module.__default.safeModuloInt(new BigNumber((_2078_v12).length), new BigNumber((_dafny.Set.fromElements((_this).f10, !((_this).f10))).length))).isLessThan(new BigNumber((_2078_v12).length)));
        }
        _2070_v5 = _module.__default.safeModuloInt(_2070_v5, _2070_v5);
        let _rhs331 = _2070_v5;
        let _rhs332 = _2070_v5;
        let _rhs333 = (_this).f10;
        let _lhs238 = globalState;
        _2070_v5 = _rhs331;
        _2070_v5 = _rhs332;
        _lhs238.f8 = _rhs333;
        let _2079_v13;
        _2079_v13 = _dafny.Seq.of((_this).f10, true);
        if (!(!(!(_dafny.Seq.IsPrefixOf(_2079_v13, _dafny.Seq.of((_this).f10))))) || ((_this).f10)) {
          let _2080_v14;
          let _init57 = function (_2081_i2) {
            return _module.__default.safeModuloInt(_2081_i2, new BigNumber((_dafny.Seq.UnicodeFromString("eqk")).length));
          };
          let _nw322 = Array((new BigNumber(13)).toNumber());
          for (let _i0_57 = 0; _i0_57 < new BigNumber(_nw322.length); _i0_57++) {
            _nw322[_i0_57] = _init57(new BigNumber(_i0_57));
          }
          _2080_v14 = _nw322;
          _2080_v14 = _2080_v14;
          let _2082_v15;
          _2082_v15 = _module.D0.create_DC0((_this).f10);
          let _2083_v16;
          _2083_v16 = _dafny.Map.Empty.slice().updateUnsafe(!((_2082_v15).dtor_cf0),_2070_v5);
          let _2084_v17;
          _2084_v17 = _dafny.Map.Empty.slice().updateUnsafe(_2070_v5,true);
          _2083_v16 = (_2083_v16).update(!((((_2084_v17).contains(_2070_v5)) ? ((_2084_v17).get(_2070_v5)) : ((_this).f10))), _module.__default.fm33((_this).f10, (_2079_v13)[_module.__default.safeIndex(_2070_v5, new BigNumber((_2079_v13).length))], new BigNumber(704), !((_this).f10), globalState));
          let _2085_v18;
          let _nw323 = new _module.C3();
          _nw323.__ctor();
          _2085_v18 = _nw323;
          let _2086_v19;
          let _nw324 = new _module.C4();
          _nw324.__ctor((_this).f9, (_this).f10);
          _2086_v19 = _nw324;
          let _2087_v20;
          _2087_v20 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_2086_v19);
          let _2088_v21;
          _2088_v21 = _dafny.Map.Empty.slice().updateUnsafe(_2085_v18,_2087_v20);
          _2088_v21 = (_2088_v21).update(_2085_v18, _2087_v20);
          let _2089_v22;
          _2089_v22 = new _dafny.CodePoint('r'.codePointAt(0));
          _2089_v22 = _2089_v22;
          let _2090_v23;
          _2090_v23 = _dafny.Seq.UnicodeFromString("sta");
          let _2091_v24;
          _2091_v24 = _dafny.Seq.of(_2090_v23);
          let _2092_v25;
          _2092_v25 = _2089_v22;
          let _2093_v26;
          _2093_v26 = _module.D15.create_DC38(!((_this).f10), (_2091_v24)[_module.__default.safeIndex(_2070_v5, new BigNumber((_2091_v24).length))], true, (_this).f10, _2092_v25);
          let _2094_v27;
          _2094_v27 = _dafny.Seq.of(_2084_v17, _2084_v17, _2084_v17, _2084_v17);
          let _2095_v28;
          let _nw325 = Array((new BigNumber(5)).toNumber()).fill(false);
          _2095_v28 = _nw325;
          let _2096_v29;
          _2096_v29 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_2095_v28);
          let _2097_v30;
          let _nw326 = new _module.C2();
          _nw326.__ctor((_this).f10, (((_2096_v29).contains((_this).f10)) ? ((_2096_v29).get((_this).f10)) : (_2095_v28)), _dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_this).f10), true);
          _2097_v30 = _nw326;
          let _2098_v31;
          _2098_v31 = _dafny.MultiSet.fromElements(_2097_v30);
          let _2099_v32;
          _2099_v32 = _module.D4.create_DC12((_this).f10, _2070_v5, _module.__default.fm31((_this).f10, (_2094_v27)[_module.__default.safeIndex(new BigNumber((_2098_v31).cardinality()), new BigNumber((_2094_v27).length))], globalState), (_dafny.ZERO).minus(_2070_v5), new BigNumber((_2090_v23).length));
          let _2100_v33;
          _2100_v33 = _dafny.Map.Empty.slice().updateUnsafe(_2093_v26,_2099_v32);
          let _2101_v34;
          _2101_v34 = _module.D6.create_DC18(p0, _2070_v5, _2070_v5);
          let _2102_v35;
          _2102_v35 = _module.D6.create_DC20(_2101_v34);
          let _2103_v36;
          _2103_v36 = _dafny.Seq.of(_module.__default.fm55((_this).f10, globalState));
          let _2104_v37;
          let _nw327 = Array((new BigNumber(22)).toNumber());
          _nw327[(_dafny.ZERO).toNumber()] = (_this).f10;
          _nw327[(_dafny.ONE).toNumber()] = (_this).f10;
          _nw327[(new BigNumber(2)).toNumber()] = (_this).f10;
          _nw327[(new BigNumber(3)).toNumber()] = !((_this).f10);
          _nw327[(new BigNumber(4)).toNumber()] = (_this).f10;
          _nw327[(new BigNumber(5)).toNumber()] = !((_this).f10);
          _nw327[(new BigNumber(6)).toNumber()] = (_this).f10;
          _nw327[(new BigNumber(7)).toNumber()] = (_this).f10;
          _nw327[(new BigNumber(8)).toNumber()] = (_this).f10;
          _nw327[(new BigNumber(9)).toNumber()] = !(_2100_v33).equals(_dafny.Map.Empty.slice().updateUnsafe(_2093_v26,_2099_v32));
          _nw327[(new BigNumber(10)).toNumber()] = (_this).f10;
          _nw327[(new BigNumber(11)).toNumber()] = (_this).f10;
          _nw327[(new BigNumber(12)).toNumber()] = (new BigNumber(((_this).f9).length)).isLessThan(_2070_v5);
          _nw327[(new BigNumber(13)).toNumber()] = !((_2097_v30.f14) || ((_this).f10));
          _nw327[(new BigNumber(14)).toNumber()] = (_this).f10;
          _nw327[(new BigNumber(15)).toNumber()] = (_2097_v30.f14) || (_2097_v30.f14);
          _nw327[(new BigNumber(16)).toNumber()] = !_dafny.Seq.contains(_2103_v36, _module.D6.create_DC20(_2102_v35));
          _nw327[(new BigNumber(17)).toNumber()] = (_this).f10;
          _nw327[(new BigNumber(18)).toNumber()] = !(!((_dafny.MultiSet.FromArray(_2079_v13)).update(_2097_v30.f14, _module.__default.abs(_2070_v5))).contains(false));
          _nw327[(new BigNumber(19)).toNumber()] = _2097_v30.f14;
          _nw327[(new BigNumber(20)).toNumber()] = (_this).f10;
          _nw327[(new BigNumber(21)).toNumber()] = _2097_v30.f14;
          _2104_v37 = _nw327;
          let _2105_v38;
          _2105_v38 = _dafny.MultiSet.fromElements(_2070_v5);
          let _index341 = _module.__default.safeIndex(new BigNumber(123), new BigNumber((_2104_v37).length));
          (_2104_v37)[_index341] = (_2105_v38).IsDisjointFrom(_2105_v38);
          let _index342 = _module.__default.safeIndex(new BigNumber(123), new BigNumber((_2104_v37).length));
          (_2104_v37)[_index342] = (_this).fm0(_module.__default.fm4(new _dafny.CodePoint('o'.codePointAt(0)), globalState), _2097_v30.f14, _2070_v5, (_this).f10, globalState);
        } else {
          let _2106_v39;
          let _nw328 = Array((new BigNumber(16)).toNumber()).fill([]);
          _2106_v39 = _nw328;
          let _2107_v40;
          let _init58 = ((_2108_v5) => function (_2109_i3) {
            return (_2109_i3).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2108_v5,(_this).f10)).length));
          })(_2070_v5);
          let _nw329 = Array((new BigNumber(26)).toNumber());
          for (let _i0_58 = 0; _i0_58 < new BigNumber(_nw329.length); _i0_58++) {
            _nw329[_i0_58] = _init58(new BigNumber(_i0_58));
          }
          _2107_v40 = _nw329;
          let _index343 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((_2106_v39).length));
          (_2106_v39)[_index343] = _2107_v40;
          let _2110_v41;
          _2110_v41 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_2107_v40);
          let _index344 = _module.__default.safeIndex(new BigNumber(893), new BigNumber((_2106_v39).length));
          (_2106_v39)[_index344] = (((_2110_v41).contains((_this).f10)) ? ((_2110_v41).get((_this).f10)) : (_2107_v40));
          let _2111_v42;
          _2111_v42 = _module.D6.create_DC18(p0, _2070_v5, new BigNumber(((_this).f9).length));
          let _arr4 = (_2106_v39)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((_2106_v39).length))];
          let _index345 = _module.__default.safeIndex(new BigNumber(675), new BigNumber(((_2106_v39)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((_2106_v39).length))]).length));
          _arr4[_index345] = (_2111_v42).dtor_cf35;
          let _2112_v43;
          let _nw330 = new _module.C4();
          _nw330.__ctor((_this).f9, (_this).f10);
          _2112_v43 = _nw330;
          let _2113_v44;
          _2113_v44 = _dafny.Map.Empty.slice().updateUnsafe(_2070_v5,_2112_v43);
          let _2114_v45;
          let _nw331 = Array((new BigNumber(16)).toNumber());
          _nw331[(_dafny.ZERO).toNumber()] = _2112_v43;
          _nw331[(_dafny.ONE).toNumber()] = _2112_v43;
          _nw331[(new BigNumber(2)).toNumber()] = _2112_v43;
          _nw331[(new BigNumber(3)).toNumber()] = _2112_v43;
          _nw331[(new BigNumber(4)).toNumber()] = _2112_v43;
          _nw331[(new BigNumber(5)).toNumber()] = _2112_v43;
          _nw331[(new BigNumber(6)).toNumber()] = _2112_v43;
          _nw331[(new BigNumber(7)).toNumber()] = _2112_v43;
          _nw331[(new BigNumber(8)).toNumber()] = _2112_v43;
          _nw331[(new BigNumber(9)).toNumber()] = (((_2113_v44).contains(_2070_v5)) ? ((_2113_v44).get(_2070_v5)) : (_2112_v43));
          _nw331[(new BigNumber(10)).toNumber()] = _2112_v43;
          _nw331[(new BigNumber(11)).toNumber()] = _2112_v43;
          _nw331[(new BigNumber(12)).toNumber()] = _2112_v43;
          _nw331[(new BigNumber(13)).toNumber()] = _2112_v43;
          _nw331[(new BigNumber(14)).toNumber()] = _2112_v43;
          _nw331[(new BigNumber(15)).toNumber()] = _2112_v43;
          _2114_v45 = _nw331;
          let _index346 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_2114_v45).length));
          (_2114_v45)[_index346] = _2112_v43;
          let _arr5 = (_2106_v39)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((_2106_v39).length))];
          let _index347 = _module.__default.safeIndex(new BigNumber(675), new BigNumber(((_2106_v39)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((_2106_v39).length))]).length));
          let _index348 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_2114_v45).length));
          let _rhs334 = (_this).f10;
          let _rhs335 = _2070_v5;
          let _rhs336 = _2112_v43;
          let _lhs239 = globalState;
          let _lhs240 = (_2106_v39)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((_2106_v39).length))];
          let _lhs241 = _module.__default.safeIndex(new BigNumber(675), new BigNumber(((_2106_v39)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((_2106_v39).length))]).length));
          let _lhs242 = _2114_v45;
          let _lhs243 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_2114_v45).length));
          _lhs239.f8 = _rhs334;
          _lhs240[_lhs241] = _rhs335;
          _lhs242[_lhs243] = _rhs336;
          let _2115_v46;
          let _init59 = function (_2116_i4) {
            return !((_this).f10);
          };
          let _nw332 = Array((new BigNumber(2)).toNumber());
          for (let _i0_59 = 0; _i0_59 < new BigNumber(_nw332.length); _i0_59++) {
            _nw332[_i0_59] = _init59(new BigNumber(_i0_59));
          }
          _2115_v46 = _nw332;
          let _index349 = _module.__default.safeIndex(new BigNumber(976), new BigNumber((_2115_v46).length));
          (_2115_v46)[_index349] = (_this).f10;
          let _index350 = _module.__default.safeIndex(new BigNumber(976), new BigNumber((_2115_v46).length));
          (_2115_v46)[_index350] = (new BigNumber(210)).isLessThanOrEqualTo(((_2106_v39)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((_2106_v39).length))])[_module.__default.safeIndex(new BigNumber(675), new BigNumber(((_2106_v39)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((_2106_v39).length))]).length))]);
          r1 = (_2115_v46)[_module.__default.safeIndex(new BigNumber(976), new BigNumber((_2115_v46).length))];
          let _init60 = ((_2117_v39, _2118_v5) => function (_2119_i5) {
            return !(function () {
              let _coll75 = new _dafny.Map();
              for (const _compr_75 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(239)), ((_2120_v5) => function (_2121_i6) {
                return _2120_v5;
              })(_2118_v5))).Elements) {
                let _2122_v47 = _compr_75;
                if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(239)), ((_2123_v5) => function (_2121_i6) {
                  return _2123_v5;
                })(_2118_v5)), _2122_v47)) {
                  _coll75.push([_module.__default.safeModuloInt(_2122_v47, ((_2117_v39)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((_2117_v39).length))])[_module.__default.safeIndex(new BigNumber(675), new BigNumber(((_2117_v39)[_module.__default.safeIndex(new BigNumber(893), new BigNumber((_2117_v39).length))]).length))]),new BigNumber((_dafny.Seq.UnicodeFromString("cxlqhjcb")).length)]);
                }
              }
              return _coll75;
            }()).contains(new BigNumber(671));
          })(_2106_v39, _2070_v5);
          let _nw333 = Array((new BigNumber(19)).toNumber());
          for (let _i0_60 = 0; _i0_60 < new BigNumber(_nw333.length); _i0_60++) {
            _nw333[_i0_60] = _init60(new BigNumber(_i0_60));
          }
          _2115_v46 = _nw333;
        }
      }
      let _2124_v48;
      _2124_v48 = new BigNumber(-908);
      let _rhs337 = _2124_v48;
      let _rhs338 = (!((_this).fm2(globalState)) || ((_this).f10)) || ((_this).f10);
      _2124_v48 = _rhs337;
      r1 = _rhs338;
      let _2125_v49;
      _2125_v49 = _dafny.Seq.of((_this).f10);
      if ((_2125_v49)[_module.__default.safeIndex(_2124_v48, new BigNumber((_2125_v49).length))]) {
        let _2126_v50;
        let _init61 = function (_2127_i7) {
          return (_this).f10;
        };
        let _nw334 = Array((_dafny.ONE).toNumber());
        for (let _i0_61 = 0; _i0_61 < new BigNumber(_nw334.length); _i0_61++) {
          _nw334[_i0_61] = _init61(new BigNumber(_i0_61));
        }
        _2126_v50 = _nw334;
        let _2128_v51;
        _2128_v51 = _dafny.MultiSet.fromElements((_this).f10);
        let _index351 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_2126_v50).length));
        (_2126_v50)[_index351] = (_2128_v51).IsSubsetOf(_dafny.MultiSet.FromArray(_2125_v49));
        let _2129_v52;
        let _init62 = ((_2130_v49) => function (_2131_i8) {
          return _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_2130_v49);
        })(_2125_v49);
        let _nw335 = Array((new BigNumber(4)).toNumber());
        for (let _i0_62 = 0; _i0_62 < new BigNumber(_nw335.length); _i0_62++) {
          _nw335[_i0_62] = _init62(new BigNumber(_i0_62));
        }
        _2129_v52 = _nw335;
        let _2132_v53;
        _2132_v53 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_2125_v49);
        let _index352 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_2129_v52).length));
        (_2129_v52)[_index352] = (_2132_v53).Merge(_2132_v53);
        let _2133_v54;
        _2133_v54 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_dafny.ZERO).minus(_2124_v48));
        let _index353 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_2126_v50).length));
        let _index354 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_2129_v52).length));
        let _rhs339 = !((_this).f10) || ((_this).f10);
        let _rhs340 = ((_this).f10) === ((new BigNumber((_2133_v54).length)).isLessThan(_2124_v48));
        let _rhs341 = ((_2132_v53).Merge((_2132_v53).update(!((_this).fm2(globalState)), _2125_v49))).Merge((_2132_v53).update((_this).f10, _2125_v49));
        let _rhs342 = (_2124_v48).minus(_2124_v48);
        let _lhs244 = _2126_v50;
        let _lhs245 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_2126_v50).length));
        let _lhs246 = _2129_v52;
        let _lhs247 = _module.__default.safeIndex(new BigNumber(181), new BigNumber((_2129_v52).length));
        _lhs244[_lhs245] = _rhs339;
        r1 = _rhs340;
        _lhs246[_lhs247] = _rhs341;
        _2124_v48 = _rhs342;
        let _2134_v55;
        _2134_v55 = _module.D6.create_DC16((_this).f9);
        let _2135_v56;
        let _nw336 = Array((new BigNumber(15)).toNumber());
        _nw336[(_dafny.ZERO).toNumber()] = (_this).f9;
        _nw336[(_dafny.ONE).toNumber()] = ((_this).f9).Merge((_this).f9);
        _nw336[(new BigNumber(2)).toNumber()] = (_module.__default.fm48((_2126_v50)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_2126_v50).length))], globalState)).update((_2126_v50)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_2126_v50).length))], (_this).f10);
        _nw336[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_2126_v50)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_2126_v50).length))],(_2126_v50)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_2126_v50).length))]);
        _nw336[(new BigNumber(4)).toNumber()] = ((_this).f9).update((_2126_v50)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_2126_v50).length))], (_2126_v50)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_2126_v50).length))]);
        _nw336[(new BigNumber(5)).toNumber()] = (_this).f9;
        _nw336[(new BigNumber(6)).toNumber()] = ((_this).f9).Merge((_this).f9);
        _nw336[(new BigNumber(7)).toNumber()] = (_2134_v55).dtor_cf28;
        _nw336[(new BigNumber(8)).toNumber()] = (_this).f9;
        _nw336[(new BigNumber(9)).toNumber()] = ((_this).f9).update((_2126_v50)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_2126_v50).length))], (_this).f10);
        _nw336[(new BigNumber(10)).toNumber()] = (_this).f9;
        _nw336[(new BigNumber(11)).toNumber()] = (_this).f9;
        _nw336[(new BigNumber(12)).toNumber()] = ((_this).f9).Merge((_this).f9);
        _nw336[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(!((_2126_v50)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_2126_v50).length))]),(_2126_v50)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_2126_v50).length))]);
        _nw336[(new BigNumber(14)).toNumber()] = (_this).f9;
        _2135_v56 = _nw336;
        let _2136_v57;
        _2136_v57 = _dafny.Seq.of((_this).f9);
        let _index355 = _module.__default.safeIndex(new BigNumber(801), new BigNumber((_2135_v56).length));
        (_2135_v56)[_index355] = ((_2136_v57)[_module.__default.safeIndex(_2124_v48, new BigNumber((_2136_v57).length))]).Merge((_this).f9);
        let _index356 = _module.__default.safeIndex(new BigNumber(801), new BigNumber((_2135_v56).length));
        (_2135_v56)[_index356] = (_this).f9;
        let _2137_v58;
        _2137_v58 = _module.D15.create_DC39(!((_this).f10) || (!((_2126_v50)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_2126_v50).length))])), (_2126_v50)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_2126_v50).length))]);
        let _source28 = _2137_v58;
        if (_source28.is_DC38) {
          let _2138___mcc_h0 = (_source28).cf69;
          let _2139___mcc_h1 = (_source28).cf70;
          let _2140___mcc_h2 = (_source28).cf71;
          let _2141___mcc_h3 = (_source28).cf72;
          let _2142___mcc_h4 = (_source28).cf73;
          let _2143_cf73 = _2142___mcc_h4;
          let _2144_cf72 = _2141___mcc_h3;
          let _2145_cf71 = _2140___mcc_h2;
          let _2146_cf70 = _2139___mcc_h1;
          let _2147_cf69 = _2138___mcc_h0;
          let _2148_v59;
          _2148_v59 = _dafny.Map.Empty.slice().updateUnsafe(_2126_v50,_2144_cf72);
          _2148_v59 = ((_2148_v59).Merge((_2148_v59).update(_2126_v50, !((_this).f10)))).Merge(_2148_v59);
          _2145_cf71 = _2144_cf72;
          let _2149_v60;
          let _nw337 = new _module.C0();
          _nw337.__ctor();
          _2149_v60 = _nw337;
          _2124_v48 = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_2124_v48), _2124_v48);
        } else if (_source28.is_DC39) {
          let _2150___mcc_h5 = (_source28).cf74;
          let _2151___mcc_h6 = (_source28).cf75;
          let _2152_cf75 = _2151___mcc_h6;
          let _2153_cf74 = _2150___mcc_h5;
          let _index357 = _module.__default.safeIndex(new BigNumber(191), new BigNumber((_2126_v50).length));
          (_2126_v50)[_index357] = (_this).f10;
          let _2154_v61;
          _2154_v61 = _dafny.Seq.UnicodeFromString("dsijjtbga");
          let _2155_v62;
          let _nw338 = new _module.C1();
          _nw338.__ctor((_this).fm2(globalState), _2154_v61);
          _2155_v62 = _nw338;
          let _2156_v63;
          let _nw339 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2156_v63 = _nw339;
          let _index358 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_2156_v63).length));
          (_2156_v63)[_index358] = _2154_v61;
          let _2157_v64;
          _2157_v64 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_2155_v62).f13);
          let _index359 = _module.__default.safeIndex(new BigNumber(295), new BigNumber((_2156_v63).length));
          (_2156_v63)[_index359] = _dafny.Seq.Concat(_dafny.Seq.Concat((((_2157_v64).contains(p0)) ? ((_2157_v64).get(p0)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-651)), ((_2158_p0) => function (_2159_i9) {
            return _2158_p0;
          })(p0)))), _2154_v61), (_2155_v62).f13);
          let _2160_v65;
          let _2161_v66;
          let _2162_v67;
          let _2163_v68;
          let _out63;
          let _out64;
          let _out65;
          let _out66;
          let _outcollector19 = (_2155_v62).m0(globalState);
          _out63 = _outcollector19[0];
          _out64 = _outcollector19[1];
          _out65 = _outcollector19[2];
          _out66 = _outcollector19[3];
          _2160_v65 = _out63;
          _2161_v66 = _out64;
          _2162_v67 = _out65;
          _2163_v68 = _out66;
        } else {
          let _2164___mcc_h7 = (_source28).cf68;
          let _2165_cf68 = _2164___mcc_h7;
          let _2166_v69;
          _2166_v69 = new _dafny.CodePoint('r'.codePointAt(0));
          let _2167_v70;
          _2167_v70 = _dafny.Map.Empty.slice().updateUnsafe(_2124_v48,(_this).f10);
          _2166_v69 = _module.__default.fm31((_2126_v50)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_2126_v50).length))], _2167_v70, globalState);
          let _2168_v71;
          _2168_v71 = _dafny.Set.fromElements(_2124_v48, _2124_v48);
          let _2169_v72;
          _2169_v72 = _dafny.Map.Empty.slice().updateUnsafe(_2168_v71,_2064_v0);
          _2124_v48 = new BigNumber(((_2169_v72).Merge(_2169_v72)).length);
          r1 = (_2126_v50)[_module.__default.safeIndex(new BigNumber(191), new BigNumber((_2126_v50).length))];
          let _2170_v73;
          let _nw340 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
          _2170_v73 = _nw340;
          let _2171_v74;
          _2171_v74 = _dafny.Map.Empty.slice().updateUnsafe(_2124_v48,_2170_v73);
          _2171_v74 = (_2171_v74).update((_2124_v48).multipliedBy(_2124_v48), _2170_v73);
        }
        _2124_v48 = _2124_v48;
        let _2172_v75;
        let _init63 = ((_2173_v48) => function (_2174_i10) {
          return (_2174_i10).plus(_2173_v48);
        })(_2124_v48);
        let _nw341 = Array((new BigNumber(28)).toNumber());
        for (let _i0_63 = 0; _i0_63 < new BigNumber(_nw341.length); _i0_63++) {
          _nw341[_i0_63] = _init63(new BigNumber(_i0_63));
        }
        _2172_v75 = _nw341;
        let _2175_v76;
        _2175_v76 = _dafny.Map.Empty.slice().updateUnsafe(_2126_v50,_2172_v75);
        let _2176_v77;
        _2176_v77 = _dafny.Map.Empty.slice().updateUnsafe(_2172_v75,(_2175_v76).Merge(_2175_v76));
        _2176_v77 = (_2176_v77).update(_2172_v75, _dafny.Map.Empty.slice().updateUnsafe(_2126_v50,_2172_v75));
      } else {
        let _2177_v78;
        _2177_v78 = _module.D0.create_DC1(new BigNumber(598), (_this).f10, (_dafny.ZERO).minus(_2124_v48));
        let _2178_v79;
        _2178_v79 = _dafny.Map.Empty.slice().updateUnsafe(_2124_v48,(_this).f11);
        let _2179_v80;
        let _nw342 = Array((new BigNumber(22)).toNumber());
        _nw342[(_dafny.ZERO).toNumber()] = (_2177_v78).dtor_cf2;
        _nw342[(_dafny.ONE).toNumber()] = (_this).f10;
        _nw342[(new BigNumber(2)).toNumber()] = (_this).f10;
        _nw342[(new BigNumber(3)).toNumber()] = (_this).f10;
        _nw342[(new BigNumber(4)).toNumber()] = (_this).f10;
        _nw342[(new BigNumber(5)).toNumber()] = (_this).f10;
        _nw342[(new BigNumber(6)).toNumber()] = (_this).f10;
        _nw342[(new BigNumber(7)).toNumber()] = (_this).f10;
        _nw342[(new BigNumber(8)).toNumber()] = (_this).f10;
        _nw342[(new BigNumber(9)).toNumber()] = (_this).f10;
        _nw342[(new BigNumber(10)).toNumber()] = !((_this).f10);
        _nw342[(new BigNumber(11)).toNumber()] = (_this).fm2(globalState);
        _nw342[(new BigNumber(12)).toNumber()] = (_this).f10;
        _nw342[(new BigNumber(13)).toNumber()] = !(!((_this).f10));
        _nw342[(new BigNumber(14)).toNumber()] = (_this).f10;
        _nw342[(new BigNumber(15)).toNumber()] = (_this).f10;
        _nw342[(new BigNumber(16)).toNumber()] = !((_this).f10);
        _nw342[(new BigNumber(17)).toNumber()] = (_this).f10;
        _nw342[(new BigNumber(18)).toNumber()] = ((_this).f11).IsDisjointFrom((((_2178_v79).contains(_2124_v48)) ? ((_2178_v79).get(_2124_v48)) : ((_this).f11)));
        _nw342[(new BigNumber(19)).toNumber()] = false;
        _nw342[(new BigNumber(20)).toNumber()] = (_this).f10;
        _nw342[(new BigNumber(21)).toNumber()] = (_this).f10;
        _2179_v80 = _nw342;
        let _index360 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_2179_v80).length));
        (_2179_v80)[_index360] = (_this).f10;
        let _index361 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_2179_v80).length));
        (_2179_v80)[_index361] = ((_this).f10) || (false);
        let _2180_v81;
        _2180_v81 = new _dafny.CodePoint('x'.codePointAt(0));
        _2180_v81 = p0;
        let _2181_v82;
        _2181_v82 = _dafny.Seq.UnicodeFromString("tf");
        if (_dafny.Seq.IsPrefixOf(_2181_v82, _2181_v82)) {
          let _2182_v83;
          let _nw343 = Array((_dafny.ONE).toNumber()).fill([]);
          _2182_v83 = _nw343;
          let _index362 = _module.__default.safeIndex(new BigNumber(268), new BigNumber((_2182_v83).length));
          (_2182_v83)[_index362] = _2179_v80;
          let _index363 = _module.__default.safeIndex(new BigNumber(268), new BigNumber((_2182_v83).length));
          (_2182_v83)[_index363] = _2179_v80;
          let _2183_v84;
          let _nw344 = new _module.C3();
          _nw344.__ctor();
          _2183_v84 = _nw344;
          _2124_v48 = _module.__default.safeDivisionInt(_2124_v48, _2124_v48);
          let _2184_v85;
          _2184_v85 = _dafny.MultiSet.fromElements(new BigNumber(997), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f10,_2124_v48)).length));
          let _2185_v86;
          _2185_v86 = _dafny.Map.Empty.slice().updateUnsafe((_2184_v85).update(_2124_v48, _module.__default.abs(_2124_v48)),_2125_v49);
          let _index364 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_2179_v80).length));
          (_2179_v80)[_index364] = ((_2124_v48).isLessThanOrEqualTo(_2124_v48)) || (!(!(_dafny.Seq.IsPrefixOf((((_2185_v86).contains(_2184_v85)) ? ((_2185_v86).get(_2184_v85)) : (_2125_v49)), _2125_v49))));
          let _2186_v87;
          _2186_v87 = _dafny.Seq.of(_2124_v48, _2124_v48);
          let _2187_v88;
          let _nw345 = Array((new BigNumber(15)).toNumber());
          _nw345[(_dafny.ZERO).toNumber()] = _2184_v85;
          _nw345[(_dafny.ONE).toNumber()] = _2184_v85;
          _nw345[(new BigNumber(2)).toNumber()] = _2184_v85;
          _nw345[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.FromArray(_2186_v87);
          _nw345[(new BigNumber(4)).toNumber()] = _module.__default.fm46((_2179_v80)[_module.__default.safeIndex(new BigNumber(524), new BigNumber((_2179_v80).length))], (_this).f10, (_2179_v80)[_module.__default.safeIndex(new BigNumber(524), new BigNumber((_2179_v80).length))], globalState);
          _nw345[(new BigNumber(5)).toNumber()] = _2184_v85;
          _nw345[(new BigNumber(6)).toNumber()] = _2184_v85;
          _nw345[(new BigNumber(7)).toNumber()] = _2184_v85;
          _nw345[(new BigNumber(8)).toNumber()] = _2184_v85;
          _nw345[(new BigNumber(9)).toNumber()] = _dafny.MultiSet.fromElements(_2124_v48);
          _nw345[(new BigNumber(10)).toNumber()] = _2184_v85;
          _nw345[(new BigNumber(11)).toNumber()] = _2184_v85;
          _nw345[(new BigNumber(12)).toNumber()] = _dafny.MultiSet.fromElements(_2124_v48, _2124_v48);
          _nw345[(new BigNumber(13)).toNumber()] = _2184_v85;
          _nw345[(new BigNumber(14)).toNumber()] = _dafny.MultiSet.FromArray(_2186_v87);
          _2187_v88 = _nw345;
          let _2188_v89;
          _2188_v89 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_2187_v88);
          let _2189_v90;
          _2189_v90 = _module.D20.create_DC51(_2187_v88);
          let _2190_v91;
          let _nw346 = Array((new BigNumber(14)).toNumber());
          _nw346[(_dafny.ZERO).toNumber()] = (((_2188_v89).contains((_this).f10)) ? ((_2188_v89).get((_this).f10)) : (_2187_v88));
          _nw346[(_dafny.ONE).toNumber()] = (_2189_v90).dtor_cf97;
          _nw346[(new BigNumber(2)).toNumber()] = _2187_v88;
          _nw346[(new BigNumber(3)).toNumber()] = (((_2188_v89).contains(true)) ? ((_2188_v89).get(true)) : (_2187_v88));
          _nw346[(new BigNumber(4)).toNumber()] = _2187_v88;
          _nw346[(new BigNumber(5)).toNumber()] = _2187_v88;
          _nw346[(new BigNumber(6)).toNumber()] = _2187_v88;
          _nw346[(new BigNumber(7)).toNumber()] = _2187_v88;
          _nw346[(new BigNumber(8)).toNumber()] = _2187_v88;
          _nw346[(new BigNumber(9)).toNumber()] = _2187_v88;
          _nw346[(new BigNumber(10)).toNumber()] = _2187_v88;
          _nw346[(new BigNumber(11)).toNumber()] = _2187_v88;
          _nw346[(new BigNumber(12)).toNumber()] = _2187_v88;
          _nw346[(new BigNumber(13)).toNumber()] = (((_2179_v80)[_module.__default.safeIndex(new BigNumber(524), new BigNumber((_2179_v80).length))]) ? (_2187_v88) : (_2187_v88));
          _2190_v91 = _nw346;
          let _index365 = _module.__default.safeIndex(new BigNumber(45), new BigNumber((_2190_v91).length));
          (_2190_v91)[_index365] = _2187_v88;
          let _2191_v92;
          _2191_v92 = _dafny.Seq.of(_2187_v88);
          let _index366 = _module.__default.safeIndex(new BigNumber(45), new BigNumber((_2190_v91).length));
          (_2190_v91)[_index366] = (_2191_v92)[_module.__default.safeIndex((new BigNumber(84)).minus(new BigNumber((_2125_v49).length)), new BigNumber((_2191_v92).length))];
        } else {
          let _2192_v93;
          _2192_v93 = _dafny.Map.Empty.slice().updateUnsafe((_2179_v80)[_module.__default.safeIndex(new BigNumber(524), new BigNumber((_2179_v80).length))],true);
          _2192_v93 = ((_this).f9).Merge((_module.D6.create_DC16((_this).f9)).dtor_cf28);
          let _2193_v94;
          _2193_v94 = _dafny.Seq.of(_2124_v48);
          let _2194_v95;
          let _init64 = ((_2195_v48) => function (_2196_i11) {
            return (_2196_i11).minus(_2195_v48);
          })(_2124_v48);
          let _nw347 = Array((new BigNumber(22)).toNumber());
          for (let _i0_64 = 0; _i0_64 < new BigNumber(_nw347.length); _i0_64++) {
            _nw347[_i0_64] = _init64(new BigNumber(_i0_64));
          }
          _2194_v95 = _nw347;
          let _2197_v96;
          _2197_v96 = _dafny.Map.Empty.slice().updateUnsafe(_2124_v48,_2124_v48);
          let _2198_v97;
          _2198_v97 = _dafny.Map.Empty.slice().updateUnsafe(_2194_v95,(((_2197_v96).contains(_2124_v48)) ? ((_2197_v96).get(_2124_v48)) : (_2124_v48)));
          let _2199_v98;
          _2199_v98 = _dafny.Map.Empty.slice().updateUnsafe(_module.D4.create_DC12(true, _2124_v48, _2180_v81, _2124_v48, _2124_v48),_2124_v48);
          let _2200_v99;
          _2200_v99 = _module.D4.create_DC12((_this).f10, _2124_v48, _2180_v81, _2124_v48, _2124_v48);
          let _2201_v100;
          _2201_v100 = _dafny.Seq.of(_2193_v94, _2193_v94);
          let _2202_v101;
          let _nw348 = Array((new BigNumber(6)).toNumber());
          _nw348[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Set.fromElements(_module.__default.fm17((_2179_v80)[_module.__default.safeIndex(new BigNumber(524), new BigNumber((_2179_v80).length))], globalState))).length);
          _nw348[(_dafny.ONE).toNumber()] = (_2193_v94)[_module.__default.safeIndex(_2124_v48, new BigNumber((_2193_v94).length))];
          _nw348[(new BigNumber(2)).toNumber()] = (new BigNumber((_2198_v97).length)).multipliedBy((((_2199_v98).contains(_2200_v99)) ? ((_2199_v98).get(_2200_v99)) : (_2124_v48)));
          _nw348[(new BigNumber(3)).toNumber()] = _2124_v48;
          _nw348[(new BigNumber(4)).toNumber()] = new BigNumber(-343);
          _nw348[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(327)), ((_2203_v94) => function (_2204_i12) {
            return _2203_v94;
          })(_2193_v94)), _2201_v100)).length);
          _2202_v101 = _nw348;
          let _index367 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_2194_v95).length));
          (_2194_v95)[_index367] = new BigNumber((_dafny.Seq.Concat(_2181_v82, _2181_v82)).length);
          let _2205_v102;
          _2205_v102 = _dafny.MultiSet.fromElements((_2179_v80)[_module.__default.safeIndex(new BigNumber(524), new BigNumber((_2179_v80).length))]);
          let _2206_v103;
          _2206_v103 = _dafny.MultiSet.fromElements(_2124_v48, new BigNumber((_dafny.Seq.update(_2181_v82, _module.__default.safeIndex((_dafny.ZERO).minus(_2124_v48), new BigNumber((_2181_v82).length)), _2180_v81)).length), _2124_v48);
          let _2207_v104;
          _2207_v104 = _dafny.Seq.of(_2206_v103, _2206_v103);
          let _2208_v105;
          _2208_v105 = _dafny.Seq.of(_2206_v103, _2206_v103, (_2207_v104)[_module.__default.safeIndex(new BigNumber((_2181_v82).length), new BigNumber((_2207_v104).length))], _module.__default.fm46((_2179_v80)[_module.__default.safeIndex(new BigNumber(524), new BigNumber((_2179_v80).length))], true, (_2179_v80)[_module.__default.safeIndex(new BigNumber(524), new BigNumber((_2179_v80).length))], globalState));
          let _index368 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_2179_v80).length));
          let _index369 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_2194_v95).length));
          let _rhs343 = _dafny.Seq.Concat(_2181_v82, _dafny.Seq.Concat(_2181_v82, _2181_v82));
          let _rhs344 = true;
          let _rhs345 = (((_this).f10) ? ((((_2197_v96).contains(new BigNumber(799))) ? ((_2197_v96).get(new BigNumber(799))) : ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_2193_v94, _module.__default.safeIndex(new BigNumber(232), new BigNumber((_2193_v94).length)), (_dafny.ZERO).minus(new BigNumber((_2205_v102).cardinality())))).length))))) : (_module.__default.safeDivisionInt(_2124_v48, (_this).fm3(new BigNumber((_dafny.Seq.UnicodeFromString("fscb")).length), _2124_v48, _2180_v81, new BigNumber((_2208_v105).length), globalState))));
          let _rhs346 = p0;
          let _rhs347 = (((_this).f10) ? (_2124_v48) : (new BigNumber(81)));
          let _lhs248 = _2179_v80;
          let _lhs249 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_2179_v80).length));
          let _lhs250 = _2194_v95;
          let _lhs251 = _module.__default.safeIndex(new BigNumber(571), new BigNumber((_2194_v95).length));
          _2181_v82 = _rhs343;
          _lhs248[_lhs249] = _rhs344;
          _2124_v48 = _rhs345;
          _2180_v81 = _rhs346;
          _lhs250[_lhs251] = _rhs347;
          let _2209_v106;
          _2209_v106 = _module.D5.create_DC13(_2179_v80);
          let _2210_v107;
          _2210_v107 = _dafny.Set.fromElements(_2209_v106, _2209_v106, _2209_v106);
          let _2211_v108;
          _2211_v108 = _dafny.Map.Empty.slice().updateUnsafe(_2180_v81,_dafny.Set.fromElements(_2209_v106));
          let _2212_v109;
          _2212_v109 = _2180_v81;
          _2210_v107 = (((_2211_v108).contains(_2212_v109)) ? ((_2211_v108).get(_2212_v109)) : (_2210_v107));
          _2124_v48 = ((_2194_v95)[_module.__default.safeIndex(new BigNumber(571), new BigNumber((_2194_v95).length))]).multipliedBy(_2124_v48);
          let _2213_v110;
          _2213_v110 = _dafny.Map.Empty.slice().updateUnsafe(_2124_v48,_2202_v101);
          _2202_v101 = (((_2213_v110).contains(_module.__default.safeDivisionInt(new BigNumber(570), _2124_v48))) ? ((_2213_v110).get(_module.__default.safeDivisionInt(new BigNumber(570), _2124_v48))) : (_2194_v95));
        }
        let _2214_v111;
        _2214_v111 = _dafny.Map.Empty.slice().updateUnsafe((_2179_v80)[_module.__default.safeIndex(new BigNumber(524), new BigNumber((_2179_v80).length))],_2124_v48);
        _2214_v111 = (_2214_v111).update((_this).f10, _2124_v48);
        let _index370 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_2179_v80).length));
        (_2179_v80)[_index370] = (_this).f10;
      }
      let _2215_v112;
      let _2216_v113;
      let _out67;
      let _out68;
      let _outcollector20 = (_this).m2(globalState);
      _out67 = _outcollector20[0];
      _out68 = _outcollector20[1];
      _2215_v112 = _out67;
      _2216_v113 = _out68;
      let _2217_v114;
      _2217_v114 = new _dafny.CodePoint('j'.codePointAt(0));
      _2217_v114 = _2217_v114;
      let _2218_v115;
      let _nw349 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
      _2218_v115 = _nw349;
      let _2219_v116;
      _2219_v116 = _dafny.Map.Empty.slice().updateUnsafe(_2218_v115,_dafny.Seq.UnicodeFromString("gwjllvov"));
      r0 = _2219_v116;
      r1 = ((_this).f10) || (true);
      return [r0, r1];
    }
    m2(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Seq.UnicodeFromString("");
      let _2220_v0;
      _2220_v0 = new BigNumber(-474);
      let _2221_v1;
      _2221_v1 = _dafny.MultiSet.fromElements(_2220_v0);
      let _2222_v2;
      _2222_v2 = _dafny.MultiSet.fromElements(new BigNumber((_2221_v1).cardinality()));
      (globalState).f8 = (_2222_v2).IsProperSubsetOf((_dafny.MultiSet.fromElements(_2220_v0)).Intersect(_2221_v1));
      if ((_this).f10) {
        let _2223_v3;
        let _nw350 = new _module.C5();
        _nw350.__ctor(((_this).f9).update(false, (_this).f10), !((_this).f10));
        _2223_v3 = _nw350;
        let _2224_v4;
        _2224_v4 = _dafny.Seq.UnicodeFromString("r");
        let _2225_v5;
        let _nw351 = new _module.C1();
        _nw351.__ctor((_this).f10, _2224_v4);
        _2225_v5 = _nw351;
        let _2226_v6;
        _2226_v6 = _module.D21.create_DC53(_2225_v5);
        let _2227_v7;
        let _nw352 = Array((new BigNumber(12)).toNumber());
        _nw352[(_dafny.ZERO).toNumber()] = _2225_v5;
        _nw352[(_dafny.ONE).toNumber()] = _2225_v5;
        _nw352[(new BigNumber(2)).toNumber()] = _2225_v5;
        _nw352[(new BigNumber(3)).toNumber()] = _2225_v5;
        _nw352[(new BigNumber(4)).toNumber()] = (_2226_v6).dtor_cf99;
        _nw352[(new BigNumber(5)).toNumber()] = _2225_v5;
        _nw352[(new BigNumber(6)).toNumber()] = _2225_v5;
        _nw352[(new BigNumber(7)).toNumber()] = _2225_v5;
        _nw352[(new BigNumber(8)).toNumber()] = _2225_v5;
        _nw352[(new BigNumber(9)).toNumber()] = _2225_v5;
        _nw352[(new BigNumber(10)).toNumber()] = _2225_v5;
        _nw352[(new BigNumber(11)).toNumber()] = _2225_v5;
        _2227_v7 = _nw352;
        let _index371 = _module.__default.safeIndex(new BigNumber(962), new BigNumber((_2227_v7).length));
        (_2227_v7)[_index371] = _2225_v5;
        let _2228_v8;
        _2228_v8 = new _dafny.CodePoint('m'.codePointAt(0));
        let _2229_v9;
        _2229_v9 = _module.D6.create_DC19(true, _2228_v8);
        let _index372 = _module.__default.safeIndex(new BigNumber(962), new BigNumber((_2227_v7).length));
        let _rhs348 = (((_2220_v0).isLessThanOrEqualTo(_2220_v0)) ? (_2225_v5) : (_2225_v5));
        let _rhs349 = (!((_this).f10)) === (false);
        let _rhs350 = _2225_v5;
        let _rhs351 = _2229_v9;
        let _lhs252 = globalState;
        let _lhs253 = _2227_v7;
        let _lhs254 = _module.__default.safeIndex(new BigNumber(962), new BigNumber((_2227_v7).length));
        _2225_v5 = _rhs348;
        _lhs252.f8 = _rhs349;
        _lhs253[_lhs254] = _rhs350;
        _2229_v9 = _rhs351;
        let _2230_v10;
        let _init65 = ((_2231_v0) => function (_2232_i0) {
          return (_2232_i0).multipliedBy(_2231_v0);
        })(_2220_v0);
        let _nw353 = Array((new BigNumber(22)).toNumber());
        for (let _i0_65 = 0; _i0_65 < new BigNumber(_nw353.length); _i0_65++) {
          _nw353[_i0_65] = _init65(new BigNumber(_i0_65));
        }
        _2230_v10 = _nw353;
        let _2233_v11;
        _2233_v11 = _dafny.MultiSet.fromElements((_2223_v3).fm18((_this).f11, _dafny.Seq.UnicodeFromString("boc"), globalState), true);
        let _2234_v12;
        _2234_v12 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("ejnrlofe")).length),new BigNumber((_2233_v11).cardinality()));
        let _index373 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_2230_v10).length));
        (_2230_v10)[_index373] = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.of((_this).f10, (_this).f10, (_this).f10)).length), new BigNumber((_2234_v12).length));
        let _2235_v13;
        let _nw354 = new _module.C4();
        _nw354.__ctor((_this).f9, (_this).f10);
        _2235_v13 = _nw354;
        let _2236_v14;
        let _nw355 = Array((new BigNumber(4)).toNumber());
        _nw355[(_dafny.ZERO).toNumber()] = _2235_v13;
        _nw355[(_dafny.ONE).toNumber()] = (_module.D22.create_DC56(_2235_v13)).dtor_cf104;
        _nw355[(new BigNumber(2)).toNumber()] = (((_this).f10) ? (_2235_v13) : (_2235_v13));
        _nw355[(new BigNumber(3)).toNumber()] = _2235_v13;
        _2236_v14 = _nw355;
        let _index374 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_2236_v14).length));
        (_2236_v14)[_index374] = _2235_v13;
        let _2237_v15;
        _2237_v15 = _dafny.Seq.of((_this).f10);
        let _index375 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_2230_v10).length));
        let _index376 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_2236_v14).length));
        let _rhs352 = ((_2223_v3).fm1((_this).f9, globalState)).multipliedBy(new BigNumber(122));
        let _rhs353 = _2220_v0;
        let _rhs354 = (_this).fm1(((_2235_v13).f9).update((_2235_v13).f10, (_2235_v13).f10), globalState);
        let _rhs355 = (new BigNumber(604)).isEqualTo(new BigNumber((_2237_v15).length));
        let _rhs356 = _2235_v13;
        let _lhs255 = _2230_v10;
        let _lhs256 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_2230_v10).length));
        let _lhs257 = globalState;
        let _lhs258 = _2236_v14;
        let _lhs259 = _module.__default.safeIndex(new BigNumber(884), new BigNumber((_2236_v14).length));
        _lhs255[_lhs256] = _rhs352;
        _2220_v0 = _rhs353;
        r0 = _rhs354;
        _lhs257.f8 = _rhs355;
        _lhs258[_lhs259] = _rhs356;
        (globalState).f8 = false;
        _2224_v4 = _dafny.Seq.update(_2224_v4, _module.__default.safeIndex((_2230_v10)[_module.__default.safeIndex(new BigNumber(347), new BigNumber((_2230_v10).length))], new BigNumber((_2224_v4).length)), _2228_v8);
      } else {
        let _2238_v16;
        let _init66 = ((_2239_v0) => function (_2240_i1) {
          return (_2240_i1).multipliedBy(_2239_v0);
        })(_2220_v0);
        let _nw356 = Array((new BigNumber(19)).toNumber());
        for (let _i0_66 = 0; _i0_66 < new BigNumber(_nw356.length); _i0_66++) {
          _nw356[_i0_66] = _init66(new BigNumber(_i0_66));
        }
        _2238_v16 = _nw356;
        let _index377 = _module.__default.safeIndex(new BigNumber(237), new BigNumber((_2238_v16).length));
        (_2238_v16)[_index377] = _2220_v0;
        let _index378 = _module.__default.safeIndex(new BigNumber(237), new BigNumber((_2238_v16).length));
        (_2238_v16)[_index378] = (_dafny.ZERO).minus(_2220_v0);
        let _2241_v17;
        _2241_v17 = _dafny.Seq.UnicodeFromString("f");
        let _2242_v18;
        _2242_v18 = _module.D14.create_DC36((_this).f10, _2241_v17, _2241_v17);
        let _2243_v19;
        _2243_v19 = _dafny.Seq.of(_2242_v18);
        let _2244_v20;
        _2244_v20 = new _dafny.CodePoint('e'.codePointAt(0));
        let _2245_v21;
        let _nw357 = Array((new BigNumber(17)).toNumber());
        _nw357[(_dafny.ZERO).toNumber()] = _dafny.Seq.of(_2242_v18, _2242_v18, _2242_v18);
        _nw357[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_module.__default.fm56((_this).f10, globalState), _module.__default.fm56(true, globalState));
        _nw357[(new BigNumber(2)).toNumber()] = _2243_v19;
        _nw357[(new BigNumber(3)).toNumber()] = _2243_v19;
        _nw357[(new BigNumber(4)).toNumber()] = _dafny.Seq.of(_2242_v18, _module.D14.create_DC36((_this).f10, _2241_v17, _dafny.Seq.Create(_module.__default.abs(new BigNumber(629)), function (_2246_i2) {
  return new _dafny.CodePoint('v'.codePointAt(0));
})));
        _nw357[(new BigNumber(5)).toNumber()] = _dafny.Seq.of(_2242_v18);
        _nw357[(new BigNumber(6)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_2242_v18), _2243_v19);
        _nw357[(new BigNumber(7)).toNumber()] = _2243_v19;
        _nw357[(new BigNumber(8)).toNumber()] = _2243_v19;
        _nw357[(new BigNumber(9)).toNumber()] = _2243_v19;
        _nw357[(new BigNumber(10)).toNumber()] = _2243_v19;
        _nw357[(new BigNumber(11)).toNumber()] = _2243_v19;
        _nw357[(new BigNumber(12)).toNumber()] = _2243_v19;
        _nw357[(new BigNumber(13)).toNumber()] = _module.__default.fm56((_this).f10, globalState);
        _nw357[(new BigNumber(14)).toNumber()] = _2243_v19;
        _nw357[(new BigNumber(15)).toNumber()] = _2243_v19;
        _nw357[(new BigNumber(16)).toNumber()] = _dafny.Seq.of(_module.D14.create_DC36((_this).f10, _dafny.Seq.update(_2241_v17, _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("lmmm")).length), new BigNumber((_2241_v17).length)), _2244_v20), _2241_v17));
        _2245_v21 = _nw357;
        let _pat_let_tv50 = _2245_v21;
        let _index379 = _module.__default.safeIndex(new BigNumber(237), new BigNumber((_2238_v16).length));
        let _rhs357 = (_2238_v16)[_module.__default.safeIndex(new BigNumber(237), new BigNumber((_2238_v16).length))];
        let _rhs358 = (((_this).f10) ? ((function (_pat_let52_0) {
          return function (_2247_dt__update__tmp_h0) {
            return function (_pat_let53_0) {
              return function (_2248_dt__update_hcf105_h0) {
                return _module.D23.create_DC58(_2248_dt__update_hcf105_h0);
              }(_pat_let53_0);
            }(_pat_let_tv50);
          }(_pat_let52_0);
        }(_module.D23.create_DC58(_2245_v21))).dtor_cf105) : (_2245_v21));
        let _lhs260 = _2238_v16;
        let _lhs261 = _module.__default.safeIndex(new BigNumber(237), new BigNumber((_2238_v16).length));
        _lhs260[_lhs261] = _rhs357;
        _2245_v21 = _rhs358;
        let _2249_v22;
        let _nw358 = Array((new BigNumber(6)).toNumber()).fill(_module.D15.Default());
        _2249_v22 = _nw358;
        let _2250_v23;
        _2250_v23 = _module.D19.create_DC49(false);
        let _2251_v24;
        _2251_v24 = _dafny.Seq.of((_this).f10, (_2250_v23).dtor_cf93, (_this).f10, (_this).f10, (_this).f10);
        let _2252_v25;
        _2252_v25 = _module.D15.create_DC39((_this).f10, (_2251_v24)[_module.__default.safeIndex((_2238_v16)[_module.__default.safeIndex(new BigNumber(237), new BigNumber((_2238_v16).length))], new BigNumber((_2251_v24).length))]);
        let _index380 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_2249_v22).length));
        (_2249_v22)[_index380] = _2252_v25;
        let _2253_v26;
        _2253_v26 = _dafny.MultiSet.fromElements((_this).f10, (_this).f10, (_this).f10, (_this).f10);
        let _pat_let_tv51 = _2253_v26;
        let _pat_let_tv52 = _2251_v24;
        let _index381 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_2249_v22).length));
        let _index382 = _module.__default.safeIndex(new BigNumber(237), new BigNumber((_2238_v16).length));
        let _index383 = _module.__default.safeIndex(new BigNumber(237), new BigNumber((_2238_v16).length));
        let _rhs359 = function (_pat_let54_0) {
          return function (_2254_dt__update__tmp_h1) {
            return function (_pat_let55_0) {
              return function (_2255_dt__update_hcf74_h0) {
                return _module.D15.create_DC39(_2255_dt__update_hcf74_h0, (_2254_dt__update__tmp_h1).dtor_cf75);
              }(_pat_let55_0);
            }((_pat_let_tv51).IsDisjointFrom(_dafny.MultiSet.FromArray(_pat_let_tv52)));
          }(_pat_let54_0);
        }(_2252_v25);
        let _rhs360 = (_this).f10;
        let _rhs361 = new BigNumber(((_this).f9).length);
        let _rhs362 = _2220_v0;
        let _lhs262 = _2249_v22;
        let _lhs263 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_2249_v22).length));
        let _lhs264 = globalState;
        let _lhs265 = _2238_v16;
        let _lhs266 = _module.__default.safeIndex(new BigNumber(237), new BigNumber((_2238_v16).length));
        let _lhs267 = _2238_v16;
        let _lhs268 = _module.__default.safeIndex(new BigNumber(237), new BigNumber((_2238_v16).length));
        _lhs262[_lhs263] = _rhs359;
        _lhs264.f8 = _rhs360;
        _lhs265[_lhs266] = _rhs361;
        _lhs267[_lhs268] = _rhs362;
        let _2256_v27;
        let _nw359 = new _module.C5();
        _nw359.__ctor((_this).f9, (_this).f10);
        _2256_v27 = _nw359;
        let _2257_v28;
        let _nw360 = Array((new BigNumber(18)).toNumber()).fill(false);
        _2257_v28 = _nw360;
        let _index384 = _module.__default.safeIndex(new BigNumber(533), new BigNumber((_2257_v28).length));
        (_2257_v28)[_index384] = (_this).f10;
        let _index385 = _module.__default.safeIndex(new BigNumber(533), new BigNumber((_2257_v28).length));
        (_2257_v28)[_index385] = (_this).f10;
      }
      let _2258_v29;
      _2258_v29 = _dafny.Seq.UnicodeFromString("glcese");
      if (!(!(!(_dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_2258_v29, _dafny.Seq.UnicodeFromString("nhghrw")), _2258_v29))))) {
        let _2259_v30;
        let _nw361 = Array((new BigNumber(26)).toNumber()).fill(false);
        _2259_v30 = _nw361;
        let _2260_v31;
        _2260_v31 = _dafny.MultiSet.fromElements((_this).f10);
        let _index386 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((_2259_v30).length));
        (_2259_v30)[_index386] = (_2260_v31).IsProperSubsetOf(_2260_v31);
        let _index387 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((_2259_v30).length));
        let _rhs363 = (_this).f10;
        let _rhs364 = _2220_v0;
        let _rhs365 = _module.__default.safeDivisionInt(_2220_v0, new BigNumber(((_this).f11).length));
        let _lhs269 = _2259_v30;
        let _lhs270 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((_2259_v30).length));
        _lhs269[_lhs270] = _rhs363;
        r0 = _rhs364;
        r0 = _rhs365;
        let _2261_v32;
        _2261_v32 = _dafny.MultiSet.fromElements(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("kiarxiha")).length)));
        (globalState).f8 = !((_2261_v32).IsSubsetOf(_2261_v32));
        r0 = _2220_v0;
        if ((_2259_v30)[_module.__default.safeIndex(new BigNumber(412), new BigNumber((_2259_v30).length))]) {
          let _2262_v33;
          let _nw362 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
          _2262_v33 = _nw362;
          let _2263_v34;
          _2263_v34 = _dafny.Seq.of(_2220_v0, _2220_v0, _2220_v0, new BigNumber(93));
          let _index388 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_2262_v33).length));
          (_2262_v33)[_index388] = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_2263_v34, _2263_v34), _module.__default.safeIndex((_2263_v34)[_module.__default.safeIndex(_2220_v0, new BigNumber((_2263_v34).length))], new BigNumber((_dafny.Seq.Concat(_2263_v34, _2263_v34)).length)), _2220_v0)).length);
          let _index389 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_2262_v33).length));
          (_2262_v33)[_index389] = (_2263_v34)[_module.__default.safeIndex(_2220_v0, new BigNumber((_2263_v34).length))];
          let _index390 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_2262_v33).length));
          (_2262_v33)[_index390] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_2263_v34, _2263_v34), _2263_v34)).length);
          let _2264_v35;
          _2264_v35 = _dafny.Set.fromElements((_2262_v33)[_module.__default.safeIndex(new BigNumber(696), new BigNumber((_2262_v33).length))], new BigNumber(815));
          let _2265_v36;
          _2265_v36 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2264_v35).length),(_2262_v33)[_module.__default.safeIndex(new BigNumber(696), new BigNumber((_2262_v33).length))]);
          let _2266_v37;
          _2266_v37 = _dafny.Seq.of((_2259_v30)[_module.__default.safeIndex(new BigNumber(412), new BigNumber((_2259_v30).length))]);
          let _2267_v38;
          _2267_v38 = new _dafny.CodePoint('v'.codePointAt(0));
          let _2268_v40;
          _2268_v40 = _dafny.Seq.of(_2264_v35);
          (globalState).f8 = !((_2265_v36).update((_dafny.ZERO).minus(new BigNumber((_module.__default.fm15((_this).f10, !((_2266_v37)[_module.__default.safeIndex(new BigNumber(-274), new BigNumber((_2266_v37).length))]), new BigNumber((_2264_v35).length), _2267_v38, globalState)).cardinality())), _2220_v0)).equals(function () {
            let _coll76 = new _dafny.Map();
            for (const _compr_76 of ((_2268_v40)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("ykixtfsa")).length), new BigNumber((_2268_v40).length))]).Elements) {
              let _2269_v39 = _compr_76;
              if (((_2268_v40)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("ykixtfsa")).length), new BigNumber((_2268_v40).length))]).contains(_2269_v39)) {
                _coll76.push([(_2269_v39).plus((_2262_v33)[_module.__default.safeIndex(new BigNumber(696), new BigNumber((_2262_v33).length))]),_2220_v0]);
              }
            }
            return _coll76;
          }());
          let _2270_v41;
          let _nw363 = new _module.C4();
          _nw363.__ctor((_this).f9, (_this).f10);
          _2270_v41 = _nw363;
          let _2271_v42;
          _2271_v42 = _dafny.Seq.of(_2270_v41);
          let _2272_v43;
          _2272_v43 = _dafny.Map.Empty.slice().updateUnsafe((((_2265_v36).contains(_2220_v0)) ? ((_2265_v36).get(_2220_v0)) : ((_2262_v33)[_module.__default.safeIndex(new BigNumber(696), new BigNumber((_2262_v33).length))])),(_2271_v42)[_module.__default.safeIndex(_2220_v0, new BigNumber((_2271_v42).length))]);
          _2272_v43 = _2272_v43;
          let _2273_v44;
          let _nw364 = new _module.C2();
          _nw364.__ctor((_2259_v30)[_module.__default.safeIndex(new BigNumber(412), new BigNumber((_2259_v30).length))], _2259_v30, (_this).f9, (_this).fm0(_2266_v37, (_2259_v30)[_module.__default.safeIndex(new BigNumber(412), new BigNumber((_2259_v30).length))], _2220_v0, (_2270_v41).fm0(_2266_v37, (_2259_v30)[_module.__default.safeIndex(new BigNumber(412), new BigNumber((_2259_v30).length))], new BigNumber(939), false, globalState), globalState));
          _2273_v44 = _nw364;
          _2273_v44 = _2273_v44;
        } else {
          let _2274_v45;
          _2274_v45 = _dafny.Map.Empty.slice().updateUnsafe(!(((_this).f11).IsProperSubsetOf((_this).f11)),(new BigNumber(237)).isEqualTo(_2220_v0));
          _2274_v45 = _2274_v45;
          _2220_v0 = _2220_v0;
          let _2275_v46;
          _2275_v46 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_2258_v29);
          let _2276_v47;
          let _nw365 = new _module.C1();
          _nw365.__ctor((_2259_v30)[_module.__default.safeIndex(new BigNumber(412), new BigNumber((_2259_v30).length))], _dafny.Seq.Concat((((_2275_v46).contains((_2259_v30)[_module.__default.safeIndex(new BigNumber(412), new BigNumber((_2259_v30).length))])) ? ((_2275_v46).get((_2259_v30)[_module.__default.safeIndex(new BigNumber(412), new BigNumber((_2259_v30).length))])) : (_2258_v29)), _2258_v29));
          _2276_v47 = _nw365;
          let _2277_v48;
          let _nw366 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
          _2277_v48 = _nw366;
          let _index391 = _module.__default.safeIndex(new BigNumber(767), new BigNumber((_2277_v48).length));
          (_2277_v48)[_index391] = _2220_v0;
          let _index392 = _module.__default.safeIndex(new BigNumber(767), new BigNumber((_2277_v48).length));
          (_2277_v48)[_index392] = _2220_v0;
          let _2278_v49;
          let _nw367 = new _module.C5();
          _nw367.__ctor((_2274_v45).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,(_2259_v30)[_module.__default.safeIndex(new BigNumber(412), new BigNumber((_2259_v30).length))])), ((_2277_v48)[_module.__default.safeIndex(new BigNumber(767), new BigNumber((_2277_v48).length))]).isEqualTo((_2277_v48)[_module.__default.safeIndex(new BigNumber(767), new BigNumber((_2277_v48).length))]));
          _2278_v49 = _nw367;
        }
        let _2279_v50;
        let _nw368 = Array((new BigNumber(10)).toNumber());
        _nw368[(_dafny.ZERO).toNumber()] = _module.__default.fm33((_2259_v30)[_module.__default.safeIndex(new BigNumber(412), new BigNumber((_2259_v30).length))], (_this).f10, _2220_v0, (_this).f10, globalState);
        _nw368[(_dafny.ONE).toNumber()] = (_2220_v0).multipliedBy(_2220_v0);
        _nw368[(new BigNumber(2)).toNumber()] = _2220_v0;
        _nw368[(new BigNumber(3)).toNumber()] = _2220_v0;
        _nw368[(new BigNumber(4)).toNumber()] = _2220_v0;
        _nw368[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(_2220_v0);
        _nw368[(new BigNumber(6)).toNumber()] = _module.__default.safeModuloInt(_2220_v0, _2220_v0);
        _nw368[(new BigNumber(7)).toNumber()] = _2220_v0;
        _nw368[(new BigNumber(8)).toNumber()] = _2220_v0;
        _nw368[(new BigNumber(9)).toNumber()] = _module.__default.fm33((_2259_v30)[_module.__default.safeIndex(new BigNumber(412), new BigNumber((_2259_v30).length))], (_this).f10, _2220_v0, true, globalState);
        _2279_v50 = _nw368;
        _2279_v50 = _2279_v50;
      } else {
        let _2280_v51;
        _2280_v51 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f10) === ((_this).f10),!((_this).f10));
        _2280_v51 = (_2280_v51).update((_this).f10, (_this).f10);
        let _2281_v52;
        let _2282_v53;
        let _2283_v54;
        let _2284_v55;
        let _out69;
        let _out70;
        let _out71;
        let _out72;
        let _outcollector21 = (_this).m3(new BigNumber(272), globalState);
        _out69 = _outcollector21[0];
        _out70 = _outcollector21[1];
        _out71 = _outcollector21[2];
        _out72 = _outcollector21[3];
        _2281_v52 = _out69;
        _2282_v53 = _out70;
        _2283_v54 = _out71;
        _2284_v55 = _out72;
        (globalState).f8 = (_this).f10;
        let _2285_v56;
        _2285_v56 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_2283_v54),(_this).f10),_2220_v0);
        let _2286_v57;
        _2286_v57 = _dafny.Seq.of(new BigNumber((_2285_v56).length));
        let _2287_v58;
        _2287_v58 = _dafny.Map.Empty.slice().updateUnsafe(_2282_v53,_module.D1.create_DC2(_2286_v57));
        let _2288_v59;
        _2288_v59 = new _dafny.CodePoint('c'.codePointAt(0));
        _2284_v55 = _dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(_2258_v29, _2258_v29), _module.__default.safeIndex(new BigNumber((_2287_v58).length), new BigNumber((_dafny.Seq.Concat(_2258_v29, _2258_v29)).length)), (_module.D14.create_DC35(_2220_v0, _2282_v53, _2288_v59)).dtor_cf64), _module.__default.safeIndex(_2283_v54, new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_2258_v29, _2258_v29), _module.__default.safeIndex(new BigNumber((_2287_v58).length), new BigNumber((_dafny.Seq.Concat(_2258_v29, _2258_v29)).length)), (_module.D14.create_DC35(_2220_v0, _2282_v53, _2288_v59)).dtor_cf64)).length)), new _dafny.CodePoint('a'.codePointAt(0)));
        let _2289_v60;
        let _init67 = function (_2290_i3) {
          return (_this).f10;
        };
        let _nw369 = Array((_dafny.ONE).toNumber());
        for (let _i0_67 = 0; _i0_67 < new BigNumber(_nw369.length); _i0_67++) {
          _nw369[_i0_67] = _init67(new BigNumber(_i0_67));
        }
        _2289_v60 = _nw369;
        let _2291_v61;
        _2291_v61 = _dafny.Seq.of(false, _2282_v53);
        let _2292_v62;
        let _nw370 = new _module.C2();
        _nw370.__ctor(false, _2289_v60, (_this).f9, ((_this).fm2(globalState)) || ((_2291_v61)[_module.__default.safeIndex((_dafny.ZERO).minus(_2283_v54), new BigNumber((_2291_v61).length))]));
        _2292_v62 = _nw370;
      }
      let _2293_v63;
      _2293_v63 = _dafny.Map.Empty.slice().updateUnsafe(_2222_v2,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(761)), function (_2294_i5) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      })).length));
      let _hi7 = new BigNumber((_2293_v63).length);
      for (let _2295_i4 = _2220_v0; _2295_i4.isLessThan(_hi7); _2295_i4 = _2295_i4.plus(_dafny.ONE)) {
        (globalState).f8 = !(_2220_v0).isEqualTo(_2295_i4);
        r0 = (_module.__default.safeModuloInt(new BigNumber((_2258_v29).length), new BigNumber((_2258_v29).length))).multipliedBy(_2295_i4);
        r1 = _dafny.Seq.Concat(_2258_v29, _dafny.Seq.UnicodeFromString("cemlbkmst"));
        let _2296_v64;
        _2296_v64 = new _dafny.CodePoint('b'.codePointAt(0));
        let _2297_v65;
        _2297_v65 = _module.D1.create_DC3(_2258_v29);
        let _2298_v66;
        let _nw371 = Array((new BigNumber(19)).toNumber());
        _nw371[(_dafny.ZERO).toNumber()] = _2258_v29;
        _nw371[(_dafny.ONE).toNumber()] = _2258_v29;
        _nw371[(new BigNumber(2)).toNumber()] = _2258_v29;
        _nw371[(new BigNumber(3)).toNumber()] = _2258_v29;
        _nw371[(new BigNumber(4)).toNumber()] = _2258_v29;
        _nw371[(new BigNumber(5)).toNumber()] = _dafny.Seq.update(_dafny.Seq.update(_2258_v29, _module.__default.safeIndex(_2220_v0, new BigNumber((_2258_v29).length)), new _dafny.CodePoint('q'.codePointAt(0))), _module.__default.safeIndex(_2295_i4, new BigNumber((_dafny.Seq.update(_2258_v29, _module.__default.safeIndex(_2220_v0, new BigNumber((_2258_v29).length)), new _dafny.CodePoint('q'.codePointAt(0)))).length)), _2296_v64);
        _nw371[(new BigNumber(6)).toNumber()] = _dafny.Seq.of(new _dafny.CodePoint('t'.codePointAt(0)));
        _nw371[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_2258_v29, _2258_v29);
        _nw371[(new BigNumber(8)).toNumber()] = _2258_v29;
        _nw371[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_2258_v29, _2258_v29);
        _nw371[(new BigNumber(10)).toNumber()] = _module.__default.fm13(_2220_v0, globalState);
        _nw371[(new BigNumber(11)).toNumber()] = (_2297_v65).dtor_cf5;
        _nw371[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_2258_v29, _2258_v29);
        _nw371[(new BigNumber(13)).toNumber()] = _2258_v29;
        _nw371[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_2258_v29, _dafny.Seq.of(_2296_v64));
        _nw371[(new BigNumber(15)).toNumber()] = _dafny.Seq.of(_2296_v64);
        _nw371[(new BigNumber(16)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(524)), function (_2299_i6) {
          return new _dafny.CodePoint('h'.codePointAt(0));
        });
        _nw371[(new BigNumber(17)).toNumber()] = _dafny.Seq.Concat(_2258_v29, _2258_v29);
        _nw371[(new BigNumber(18)).toNumber()] = _2258_v29;
        _2298_v66 = _nw371;
        let _index393 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_2298_v66).length));
        (_2298_v66)[_index393] = _2258_v29;
        let _2300_v67;
        let _nw372 = new _module.C6();
        _nw372.__ctor(_dafny.Map.Empty.slice().updateUnsafe((_this).f10,false), !((_this).f10));
        _2300_v67 = _nw372;
        let _2301_v68;
        _2301_v68 = _dafny.Set.fromElements(_2295_i4, _2295_i4);
        let _index394 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_2298_v66).length));
        let _rhs366 = _dafny.MultiSet.fromElements(_2220_v0, new BigNumber((_dafny.Seq.of(!((_this).f10), (_this).f10)).length), (((_this).f10) ? (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2300_v67,!((_this).f10))).length)) : (_2295_i4)), _2295_i4, new BigNumber((_2301_v68).length));
        let _rhs367 = _module.__default.safeDivisionInt(_2295_i4, (_dafny.ZERO).minus(_2295_i4));
        let _rhs368 = _dafny.Seq.update(_2258_v29, _module.__default.safeIndex(_2295_i4, new BigNumber((_2258_v29).length)), new _dafny.CodePoint('x'.codePointAt(0)));
        let _rhs369 = (_this).f10;
        let _rhs370 = _2296_v64;
        let _lhs271 = _2298_v66;
        let _lhs272 = _module.__default.safeIndex(new BigNumber(90), new BigNumber((_2298_v66).length));
        let _lhs273 = globalState;
        _2222_v2 = _rhs366;
        _2220_v0 = _rhs367;
        _lhs271[_lhs272] = _rhs368;
        _lhs273.f8 = _rhs369;
        _2296_v64 = _rhs370;
      }
      let _2302_v69;
      _2302_v69 = _dafny.Seq.of((_this).f10);
      let _2303_v70;
      let _nw373 = Array((new BigNumber(13)).toNumber());
      _nw373[(_dafny.ZERO).toNumber()] = true;
      _nw373[(_dafny.ONE).toNumber()] = (_this).f10;
      _nw373[(new BigNumber(2)).toNumber()] = (_this).f10;
      _nw373[(new BigNumber(3)).toNumber()] = (_this).f10;
      _nw373[(new BigNumber(4)).toNumber()] = (_this).f10;
      _nw373[(new BigNumber(5)).toNumber()] = (_this).f10;
      _nw373[(new BigNumber(6)).toNumber()] = (_this).fm0(_2302_v69, (_this).f10, new BigNumber((_module.__default.fm25(_2220_v0, new BigNumber(435), globalState)).length), (_this).f10, globalState);
      _nw373[(new BigNumber(7)).toNumber()] = (_this).f10;
      _nw373[(new BigNumber(8)).toNumber()] = (_this).f10;
      _nw373[(new BigNumber(9)).toNumber()] = (_this).f10;
      _nw373[(new BigNumber(10)).toNumber()] = (_this).f10;
      _nw373[(new BigNumber(11)).toNumber()] = (_this).f10;
      _nw373[(new BigNumber(12)).toNumber()] = (_this).f10;
      _2303_v70 = _nw373;
      let _2304_v71;
      _2304_v71 = _dafny.Seq.of(_2303_v70);
      let _2305_v72;
      _2305_v72 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(_2220_v0)),(_this).f10);
      let _2306_v73;
      _2306_v73 = _module.D5.create_DC13((_2304_v71)[_module.__default.safeIndex(new BigNumber((_2305_v72).length), new BigNumber((_2304_v71).length))]);
      let _source29 = _2306_v73;
      if (_source29.is_DC14) {
        let _2307___mcc_h0 = (_source29).cf25;
        let _2308___mcc_h1 = (_source29).cf26;
        let _2309_cf26 = _2308___mcc_h1;
        let _2310_cf25 = _2307___mcc_h0;
        let _2311_v74;
        _2311_v74 = _dafny.Seq.of((_this).f11, (_this).f11, (_this).f11, (_this).f11);
        let _2312_v75;
        let _nw374 = Array((new BigNumber(17)).toNumber());
        _nw374[(_dafny.ZERO).toNumber()] = _module.__default.fm30(globalState);
        _nw374[(_dafny.ONE).toNumber()] = (_dafny.Set.fromElements(_2309_cf26, (_this).fm0(_2302_v69, _2309_cf26, new BigNumber(497), (_this).f10, globalState), _2309_cf26, _2309_cf26, _2309_cf26)).Intersect(_dafny.Set.fromElements((_this).f10, false, _2309_cf26));
        _nw374[(new BigNumber(2)).toNumber()] = _dafny.Set.fromElements(_2309_cf26, _2309_cf26, _2309_cf26);
        _nw374[(new BigNumber(3)).toNumber()] = ((_2311_v74)[_module.__default.safeIndex(_2310_cf25, new BigNumber((_2311_v74).length))]).Difference((_this).f11);
        _nw374[(new BigNumber(4)).toNumber()] = ((_this).f11).Union((_this).f11);
        _nw374[(new BigNumber(5)).toNumber()] = (_this).f11;
        _nw374[(new BigNumber(6)).toNumber()] = (_this).f11;
        _nw374[(new BigNumber(7)).toNumber()] = (_this).f11;
        _nw374[(new BigNumber(8)).toNumber()] = (_this).f11;
        _nw374[(new BigNumber(9)).toNumber()] = (_this).f11;
        _nw374[(new BigNumber(10)).toNumber()] = (_this).f11;
        _nw374[(new BigNumber(11)).toNumber()] = (_this).f11;
        _nw374[(new BigNumber(12)).toNumber()] = _dafny.Set.fromElements((_this).f10);
        _nw374[(new BigNumber(13)).toNumber()] = (_this).f11;
        _nw374[(new BigNumber(14)).toNumber()] = ((_this).f11).Difference((_this).f11);
        _nw374[(new BigNumber(15)).toNumber()] = (_this).f11;
        _nw374[(new BigNumber(16)).toNumber()] = (_this).f11;
        _2312_v75 = _nw374;
        _2312_v75 = _2312_v75;
        let _2313_v76;
        let _nw375 = new _module.C5();
        _nw375.__ctor((_this).f9, _2309_cf26);
        _2313_v76 = _nw375;
        let _2314_v77;
        let _nw376 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
        _2314_v77 = _nw376;
        let _index395 = _module.__default.safeIndex(new BigNumber(543), new BigNumber((_2314_v77).length));
        (_2314_v77)[_index395] = _2220_v0;
        let _index396 = _module.__default.safeIndex(new BigNumber(543), new BigNumber((_2314_v77).length));
        (_2314_v77)[_index396] = _module.__default.safeModuloInt(new BigNumber((_2258_v29).length), _2310_cf25);
        _2310_cf25 = (_dafny.ZERO).minus((_2220_v0).minus(_2220_v0));
      } else if (_source29.is_DC13) {
        let _2315___mcc_h2 = (_source29).cf24;
        let _2316_cf24 = _2315___mcc_h2;
        let _2317_v78;
        _2317_v78 = new _dafny.CodePoint('c'.codePointAt(0));
        let _2318_v79;
        _2318_v79 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm36(globalState),_2220_v0);
        let _2319_v80;
        _2319_v80 = _dafny.Set.fromElements((_this).fm3((_dafny.ZERO).minus(_2220_v0), _2220_v0, _2317_v78, new BigNumber((_2318_v79).length), globalState));
        let _2320_v81;
        _2320_v81 = _dafny.Seq.of((((_this).f10) ? (_dafny.Set.fromElements(_2220_v0, _2220_v0)) : (_dafny.Set.fromElements(_2220_v0, _2220_v0))), _dafny.Set.fromElements(new BigNumber(803), _2220_v0), _2319_v80);
        _2320_v81 = _dafny.Seq.Concat(_2320_v81, _2320_v81);
        _2220_v0 = _2220_v0;
        if ((_this).fm2(globalState)) {
          r0 = _2220_v0;
          let _2321_v82;
          let _nw377 = Array((new BigNumber(5)).toNumber()).fill(_module.D15.Default());
          _2321_v82 = _nw377;
          let _2322_v83;
          _2322_v83 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm3(_2220_v0, _2220_v0, _2317_v78, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_2317_v78,!((_this).f10))).length), globalState),_2321_v82);
          let _2323_v84;
          let _nw378 = Array((new BigNumber(6)).toNumber());
          _nw378[(_dafny.ZERO).toNumber()] = (_2322_v83).Merge(_2322_v83);
          _nw378[(_dafny.ONE).toNumber()] = (_2322_v83).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2220_v0,_2321_v82));
          _nw378[(new BigNumber(2)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_2220_v0,_2321_v82)).Merge(_2322_v83);
          _nw378[(new BigNumber(3)).toNumber()] = (_2322_v83).Merge(_2322_v83);
          _nw378[(new BigNumber(4)).toNumber()] = _2322_v83;
          _nw378[(new BigNumber(5)).toNumber()] = _2322_v83;
          _2323_v84 = _nw378;
          let _index397 = _module.__default.safeIndex(new BigNumber(576), new BigNumber((_2323_v84).length));
          (_2323_v84)[_index397] = _2322_v83;
          let _2324_v85;
          let _nw379 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
          _2324_v85 = _nw379;
          let _index398 = _module.__default.safeIndex(new BigNumber(576), new BigNumber((_2323_v84).length));
          let _rhs371 = (((_this).f10) ? (_dafny.Map.Empty.slice().updateUnsafe(_2220_v0,_2321_v82)) : (_2322_v83));
          let _rhs372 = _2317_v78;
          let _rhs373 = _2324_v85;
          let _rhs374 = _2220_v0;
          let _lhs274 = _2323_v84;
          let _lhs275 = _module.__default.safeIndex(new BigNumber(576), new BigNumber((_2323_v84).length));
          _lhs274[_lhs275] = _rhs371;
          _2317_v78 = _rhs372;
          _2324_v85 = _rhs373;
          _2220_v0 = _rhs374;
          let _index399 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_2316_cf24).length));
          (_2316_cf24)[_index399] = false;
          let _2325_v86;
          _2325_v86 = _dafny.Map.Empty.slice().updateUnsafe(_2324_v85,true);
          let _2326_v87;
          _2326_v87 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(427)), ((_2327_v78) => function (_2328_i7) {
            return _2327_v78;
          })(_2317_v78)),(((_2325_v86).contains(_2324_v85)) ? ((_2325_v86).get(_2324_v85)) : ((_this).f10)));
          let _index400 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_2316_cf24).length));
          let _rhs375 = (((_2326_v87).contains(_dafny.Seq.update((_module.D1.create_DC3(_2258_v29)).dtor_cf5, _module.__default.safeIndex(_2220_v0, new BigNumber(((_module.D1.create_DC3(_2258_v29)).dtor_cf5).length)), new _dafny.CodePoint('d'.codePointAt(0))))) ? ((_2326_v87).get(_dafny.Seq.update((_module.D1.create_DC3(_2258_v29)).dtor_cf5, _module.__default.safeIndex(_2220_v0, new BigNumber(((_module.D1.create_DC3(_2258_v29)).dtor_cf5).length)), new _dafny.CodePoint('d'.codePointAt(0))))) : ((_this).f10));
          let _rhs376 = !((!((_this).f10)) || ((_this).f10));
          let _rhs377 = _2220_v0;
          let _rhs378 = (_this).f10;
          let _lhs276 = _2316_cf24;
          let _lhs277 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_2316_cf24).length));
          let _lhs278 = globalState;
          let _lhs279 = globalState;
          _lhs276[_lhs277] = _rhs375;
          _lhs278.f8 = _rhs376;
          r0 = _rhs377;
          _lhs279.f8 = _rhs378;
          let _2329_v88;
          let _nw380 = new _module.C6();
          _nw380.__ctor((_this).f9, (_this).f10);
          _2329_v88 = _nw380;
          let _index401 = _module.__default.safeIndex(new BigNumber(384), new BigNumber((_2324_v85).length));
          (_2324_v85)[_index401] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_2220_v0, _2220_v0));
          let _2330_v89;
          _2330_v89 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f10,_2258_v29)).length));
          let _index402 = _module.__default.safeIndex(new BigNumber(384), new BigNumber((_2324_v85).length));
          (_2324_v85)[_index402] = (((_2330_v89).contains((_this).f10)) ? ((_2330_v89).get((_this).f10)) : (_2220_v0));
        } else {
          r1 = (((_this).f10) ? (_2258_v29) : (_dafny.Seq.Concat(_2258_v29, _2258_v29)));
          let _2331_v90;
          _2331_v90 = _module.D22.create_DC57();
          _2331_v90 = _2331_v90;
          r0 = (_dafny.ZERO).minus(((((_this).f10) ? (_2220_v0) : (_2220_v0))).plus(new BigNumber((_2258_v29).length)));
          let _2332_v91;
          let _nw381 = new _module.C1();
          _nw381.__ctor((_this).f10, _dafny.Seq.UnicodeFromString("v"));
          _2332_v91 = _nw381;
          let _2333_v92;
          _2333_v92 = _dafny.Seq.of(_2332_v91, _2332_v91);
          let _2334_v93;
          _2334_v93 = _dafny.Set.fromElements(_dafny.Seq.Concat(_dafny.Seq.update(_2333_v92, _module.__default.safeIndex(_2220_v0, new BigNumber((_2333_v92).length)), _2332_v91), _2333_v92), _dafny.Seq.Concat(_2333_v92, _dafny.Seq.update(_2333_v92, _module.__default.safeIndex(new BigNumber((_2258_v29).length), new BigNumber((_2333_v92).length)), _2332_v91)));
          _2334_v93 = _2334_v93;
          _2220_v0 = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_2320_v81, _2320_v81)).length)));
        }
        (globalState).f8 = (_this).f10;
      } else {
        let _2335___mcc_h3 = (_source29).cf27;
        let _2336_cf27 = _2335___mcc_h3;
        let _2337_v94;
        _2337_v94 = _dafny.Seq.of((_dafny.ZERO).minus(_2220_v0), _2220_v0);
        _2337_v94 = _2337_v94;
        let _2338_v95;
        let _nw382 = new _module.C1();
        _nw382.__ctor((_this).f10, _2258_v29);
        _2338_v95 = _nw382;
        _2338_v95 = (((_this).fm2(globalState)) ? (_2338_v95) : (_2338_v95));
        let _2339_v96;
        _2339_v96 = new _dafny.CodePoint('i'.codePointAt(0));
        let _2340_v97;
        _2340_v97 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm36(globalState),_2339_v96);
        (globalState).f8 = !(new BigNumber((_2340_v97).length)).isEqualTo(_2220_v0);
        _2220_v0 = (_2220_v0).minus((new BigNumber(33)).minus(_2220_v0));
      }
      let _2341_v98;
      let _nw383 = new _module.C2();
      _nw383.__ctor((_2302_v69)[_module.__default.safeIndex(_2220_v0, new BigNumber((_2302_v69).length))], _2303_v70, (_dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_this).f10)).Merge((_this).f9), (_this).f10);
      _2341_v98 = _nw383;
      r0 = new BigNumber((_module.__default.fm46(false, (_2222_v2).IsDisjointFrom(_2222_v2), _2341_v98.f14, globalState)).cardinality());
      let _2342_v99;
      _2342_v99 = new _dafny.CodePoint('g'.codePointAt(0));
      r1 = _dafny.Seq.update(_2258_v29, _module.__default.safeIndex(new BigNumber(((_this).f11).length), new BigNumber((_2258_v29).length)), _2342_v99);
      return [r0, r1];
    }
    m3(p0, globalState) {
      let _this = this;
      let r0 = _dafny.MultiSet.Empty;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let r3 = _dafny.Seq.UnicodeFromString("");
      let _2343_v0;
      _2343_v0 = new _dafny.CodePoint('x'.codePointAt(0));
      let _2344_v1;
      _2344_v1 = _module.D6.create_DC18(_2343_v0, p0, p0);
      let _2345_v2;
      _2345_v2 = _module.D18.create_DC47((_this).f10, _2344_v1);
      let _2346_v3;
      _2346_v3 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_module.__default.safeModuloInt(p0, new BigNumber(958))),_2345_v2);
      let _2347_v4;
      _2347_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,p0);
      let _2348_v5;
      _2348_v5 = _dafny.MultiSet.fromElements((_this).f10, (_this).f10, (_this).f10, (_this).f10);
      let _2349_v6;
      _2349_v6 = _dafny.MultiSet.fromElements(new BigNumber((_2348_v5).cardinality()));
      let _2350_v7;
      _2350_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_2347_v4)).length),new BigNumber((_2349_v6).cardinality()));
      let _rhs379 = ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2350_v7).length),_2345_v2)).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,_2345_v2))).Merge((_dafny.Map.Empty.slice().updateUnsafe(p0,_2345_v2)).Merge(_2346_v3));
      let _rhs380 = p0;
      _2346_v3 = _rhs379;
      r2 = _rhs380;
      let _2351_i0;
      _2351_i0 = _dafny.ZERO;
      L9: {
        while ((_this).f10) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2351_i0)) {
              break L9;
            }
            _2351_i0 = (_2351_i0).plus(_dafny.ONE);
            let _2352_v8;
            let _init68 = function (_2353_i1) {
              return _dafny.Seq.UnicodeFromString("dvupvsxfh");
            };
            let _nw384 = Array((new BigNumber(9)).toNumber());
            for (let _i0_68 = 0; _i0_68 < new BigNumber(_nw384.length); _i0_68++) {
              _nw384[_i0_68] = _init68(new BigNumber(_i0_68));
            }
            _2352_v8 = _nw384;
            let _2354_v9;
            _2354_v9 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2352_v8);
            _2354_v9 = (_2354_v9).update((_dafny.ZERO).minus(_module.__default.safeModuloInt(p0, p0)), (((_this).f10) ? (_2352_v8) : (_2352_v8)));
            r2 = (new BigNumber(-261)).minus(p0);
            r1 = (_this).f10;
            let _2355_v10;
            let _nw385 = new _module.C5();
            _nw385.__ctor((_this).f9, false);
            _2355_v10 = _nw385;
          }
        }
      }
      r2 = _module.__default.safeModuloInt((((_this).f10) ? (p0) : (p0)), _module.__default.safeDivisionInt(p0, p0));
      if ((_this).f10) {
        r2 = (_this).fm1((_this).f9, globalState);
        let _2356_v11;
        let _nw386 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
        _2356_v11 = _nw386;
        let _2357_v12;
        let _nw387 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Seq.of());
        _2357_v12 = _nw387;
        let _index403 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_2357_v12).length));
        (_2357_v12)[_index403] = _dafny.Seq.of((_this).f10, (_this).f10);
        let _2358_v13;
        let _nw388 = new _module.C4();
        _nw388.__ctor(((_this).f9).update((_this).f10, (_this).f10), (_this).f10);
        _2358_v13 = _nw388;
        let _index404 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_2356_v11).length));
        (_2356_v11)[_index404] = p0;
        let _2359_v14;
        _2359_v14 = _dafny.MultiSet.fromElements(_2343_v0);
        let _2360_v15;
        _2360_v15 = _dafny.Seq.of(((_dafny.ZERO).minus(p0)).isLessThanOrEqualTo(p0), (_2359_v14).IsDisjointFrom(_2359_v14));
        let _index405 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_2357_v12).length));
        let _index406 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_2356_v11).length));
        let _rhs381 = _2356_v11;
        let _rhs382 = _2360_v15;
        let _rhs383 = _2358_v13;
        let _rhs384 = (((!(false)) ? (p0) : (p0))).multipliedBy((((_2349_v6).contains(p0)) ? ((_2349_v6).get(p0)) : (p0)));
        let _lhs280 = _2357_v12;
        let _lhs281 = _module.__default.safeIndex(new BigNumber(484), new BigNumber((_2357_v12).length));
        let _lhs282 = _2356_v11;
        let _lhs283 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_2356_v11).length));
        _2356_v11 = _rhs381;
        _lhs280[_lhs281] = _rhs382;
        _2358_v13 = _rhs383;
        _lhs282[_lhs283] = _rhs384;
        let _2361_v16;
        _2361_v16 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(155)).isLessThanOrEqualTo((_dafny.ZERO).minus((_2356_v11)[_module.__default.safeIndex(new BigNumber(311), new BigNumber((_2356_v11).length))])),_2343_v0);
        _2343_v0 = (((_2361_v16).contains(false)) ? ((_2361_v16).get(false)) : (_2343_v0));
        let _2362_v17;
        _2362_v17 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f10) ? ((_this).f10) : (true)),_2348_v5);
        _2362_v17 = (_2362_v17).update((_this).f10, _2348_v5);
        let _2363_v18;
        let _init69 = function (_2364_i2) {
          return (_this).f10;
        };
        let _nw389 = Array((new BigNumber(20)).toNumber());
        for (let _i0_69 = 0; _i0_69 < new BigNumber(_nw389.length); _i0_69++) {
          _nw389[_i0_69] = _init69(new BigNumber(_i0_69));
        }
        _2363_v18 = _nw389;
        _2363_v18 = ((true) ? (_2363_v18) : (_2363_v18));
      } else {
        let _2365_v19;
        _2365_v19 = _module.D19.create_DC49(true);
        let _2366_v20;
        _2366_v20 = _module.D1.create_DC4((_2365_v19).dtor_cf93, (_this).f10, p0);
        let _2367_v21;
        _2367_v21 = _dafny.Seq.of((_this).f10);
        let _2368_v22;
        _2368_v22 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2367_v21);
        let _2369_v23;
        _2369_v23 = _dafny.Set.fromElements(new _dafny.CodePoint('d'.codePointAt(0)));
        let _2370_v24;
        _2370_v24 = _dafny.Map.Empty.slice().updateUnsafe(((_2366_v20).dtor_cf8).minus(new BigNumber(((((_2368_v22).contains(p0)) ? ((_2368_v22).get(p0)) : (_2367_v21))).length)),(_2369_v23).IsSubsetOf(_2369_v23));
        _2370_v24 = (_2370_v24).update(p0, (_this).f10);
        let _2371_v25;
        let _nw390 = Array((new BigNumber(10)).toNumber()).fill(false);
        _2371_v25 = _nw390;
        let _2372_v26;
        _2372_v26 = _module.D6.create_DC16((_this).f9);
        let _2373_v27;
        _2373_v27 = _dafny.Set.fromElements(new BigNumber(-117));
        let _2374_v28;
        _2374_v28 = _module.D23.create_DC59(_2372_v26, (_this).f10, new BigNumber((_2373_v27).length), (_module.D5.create_DC13((_module.D5.create_DC13(_2371_v25)).dtor_cf24)).dtor_cf24);
        let _2375_v29;
        let _nw391 = Array((new BigNumber(5)).toNumber());
        _nw391[(_dafny.ZERO).toNumber()] = _2371_v25;
        _nw391[(_dafny.ONE).toNumber()] = _2371_v25;
        _nw391[(new BigNumber(2)).toNumber()] = _2371_v25;
        _nw391[(new BigNumber(3)).toNumber()] = (_2374_v28).dtor_cf109;
        _nw391[(new BigNumber(4)).toNumber()] = _2371_v25;
        _2375_v29 = _nw391;
        _2375_v29 = _2375_v29;
        let _2376_v30;
        _2376_v30 = _dafny.Seq.of(p0);
        let _2377_v31;
        let _nw392 = Array((new BigNumber(2)).toNumber());
        _nw392[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_2376_v30, _2376_v30);
        _nw392[(_dafny.ONE).toNumber()] = _2376_v30;
        _2377_v31 = _nw392;
        let _index407 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((_2377_v31).length));
        (_2377_v31)[_index407] = _dafny.Seq.Concat(_2376_v30, _dafny.Seq.of(p0));
        let _index408 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((_2377_v31).length));
        (_2377_v31)[_index408] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(423)), function (_2378_i3) {
          return new BigNumber(-744);
        }), _2376_v30);
        _2370_v24 = (_2370_v24).update((_dafny.ZERO).minus(p0), (_this).f10);
        _2367_v21 = _dafny.Seq.update(_2367_v21, _module.__default.safeIndex(p0, new BigNumber((_2367_v21).length)), (_this).f10);
      }
      let _hi8 = new BigNumber(643);
      for (let _2379_i4 = p0; _2379_i4.isLessThan(_hi8); _2379_i4 = _2379_i4.plus(_dafny.ONE)) {
        let _2380_v32;
        _2380_v32 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,_2379_i4);
        let _2381_v33;
        _2381_v33 = _dafny.Seq.of(new BigNumber((_2380_v32).length));
        let _source30 = _module.D16.create_DC42((_this).f10, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f10,!(false))).length), (((_this).f10) ? (new BigNumber((_dafny.MultiSet.FromArray(_2381_v33)).cardinality())) : (_2379_i4)));
        if (_source30.is_DC41) {
          let _2382___mcc_h0 = (_source30).cf77;
          let _2383___mcc_h1 = (_source30).cf78;
          let _2384___mcc_h2 = (_source30).cf79;
          let _2385_cf79 = _2384___mcc_h2;
          let _2386_cf78 = _2383___mcc_h1;
          let _2387_cf77 = _2382___mcc_h0;
          let _2388_v34;
          let _nw393 = Array((_dafny.ONE).toNumber()).fill(false);
          _2388_v34 = _nw393;
          let _2389_v35;
          let _nw394 = new _module.C2();
          _nw394.__ctor(_2385_cf79, _2388_v34, (_this).f9, _2387_cf77);
          _2389_v35 = _nw394;
          let _2390_v36;
          _2390_v36 = _dafny.Seq.of(_2389_v35, _2389_v35);
          _2390_v36 = _2390_v36;
          let _index409 = _module.__default.safeIndex(new BigNumber(545), new BigNumber((_2388_v34).length));
          (_2388_v34)[_index409] = _2387_cf77;
          let _index410 = _module.__default.safeIndex(new BigNumber(545), new BigNumber((_2388_v34).length));
          (_2388_v34)[_index410] = !((_this).f10);
          let _2391_v37;
          let _nw395 = Array((new BigNumber(11)).toNumber());
          _nw395[(_dafny.ZERO).toNumber()] = _2343_v0;
          _nw395[(_dafny.ONE).toNumber()] = _2343_v0;
          _nw395[(new BigNumber(2)).toNumber()] = _2343_v0;
          _nw395[(new BigNumber(3)).toNumber()] = _2343_v0;
          _nw395[(new BigNumber(4)).toNumber()] = _2343_v0;
          _nw395[(new BigNumber(5)).toNumber()] = _2343_v0;
          _nw395[(new BigNumber(6)).toNumber()] = _2343_v0;
          _nw395[(new BigNumber(7)).toNumber()] = _2343_v0;
          _nw395[(new BigNumber(8)).toNumber()] = _2343_v0;
          _nw395[(new BigNumber(9)).toNumber()] = _2343_v0;
          _nw395[(new BigNumber(10)).toNumber()] = new _dafny.CodePoint('y'.codePointAt(0));
          _2391_v37 = _nw395;
          let _2392_v38;
          _2392_v38 = _module.D25.create_DC61(_2391_v37);
          let _2393_v39;
          _2393_v39 = _dafny.Seq.of(_2391_v37, _2391_v37, _2391_v37, _2391_v37);
          let _2394_v40;
          _2394_v40 = _dafny.Map.Empty.slice().updateUnsafe((_2388_v34)[_module.__default.safeIndex(new BigNumber(545), new BigNumber((_2388_v34).length))],_2380_v32);
          let _2395_v41;
          let _nw396 = Array((new BigNumber(13)).toNumber());
          _nw396[(_dafny.ZERO).toNumber()] = _2391_v37;
          _nw396[(_dafny.ONE).toNumber()] = (_2392_v38).dtor_cf111;
          _nw396[(new BigNumber(2)).toNumber()] = _2391_v37;
          _nw396[(new BigNumber(3)).toNumber()] = _2391_v37;
          _nw396[(new BigNumber(4)).toNumber()] = _2391_v37;
          _nw396[(new BigNumber(5)).toNumber()] = (_2393_v39)[_module.__default.safeIndex(new BigNumber((_2394_v40).length), new BigNumber((_2393_v39).length))];
          _nw396[(new BigNumber(6)).toNumber()] = _2391_v37;
          _nw396[(new BigNumber(7)).toNumber()] = _2391_v37;
          _nw396[(new BigNumber(8)).toNumber()] = _2391_v37;
          _nw396[(new BigNumber(9)).toNumber()] = _2391_v37;
          _nw396[(new BigNumber(10)).toNumber()] = _2391_v37;
          _nw396[(new BigNumber(11)).toNumber()] = _2391_v37;
          _nw396[(new BigNumber(12)).toNumber()] = (((_2389_v35).f10) ? (_2391_v37) : (_2391_v37));
          _2395_v41 = _nw396;
          let _index411 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_2395_v41).length));
          (_2395_v41)[_index411] = ((_2385_cf79) ? (_2391_v37) : (_2391_v37));
          let _index412 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_2395_v41).length));
          let _rhs385 = _2379_i4;
          let _rhs386 = _2391_v37;
          let _lhs284 = _2395_v41;
          let _lhs285 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_2395_v41).length));
          r2 = _rhs385;
          _lhs284[_lhs285] = _rhs386;
          let _2396_v42;
          let _nw397 = new _module.C4();
          _nw397.__ctor((_2389_v35).f9, (_this).f10);
          _2396_v42 = _nw397;
        } else if (_source30.is_DC42) {
          let _2397___mcc_h3 = (_source30).cf80;
          let _2398___mcc_h4 = (_source30).cf81;
          let _2399___mcc_h5 = (_source30).cf82;
          let _2400_cf82 = _2399___mcc_h5;
          let _2401_cf81 = _2398___mcc_h4;
          let _2402_cf80 = _2397___mcc_h3;
          let _2403_v43;
          let _nw398 = new _module.C1();
          _nw398.__ctor(_2402_cf80, _dafny.Seq.UnicodeFromString("knxgb"));
          _2403_v43 = _nw398;
          let _2404_v44;
          _2404_v44 = _module.D26.create_DC64(_2403_v43);
          let _2405_v45;
          _2405_v45 = _dafny.Set.fromElements(_2403_v43, _2403_v43, _2403_v43, _2403_v43, (_2404_v44).dtor_cf114);
          let _2406_v47;
          let _nw399 = Array((new BigNumber(25)).toNumber());
          _nw399[(_dafny.ZERO).toNumber()] = ((_this).f10) && ((_this).f10);
          _nw399[(_dafny.ONE).toNumber()] = _2402_cf80;
          _nw399[(new BigNumber(2)).toNumber()] = _2402_cf80;
          _nw399[(new BigNumber(3)).toNumber()] = !(_2402_cf80) || (_2402_cf80);
          _nw399[(new BigNumber(4)).toNumber()] = true;
          _nw399[(new BigNumber(5)).toNumber()] = (_2405_v45).IsDisjointFrom(_dafny.Set.fromElements(_2403_v43));
          _nw399[(new BigNumber(6)).toNumber()] = false;
          _nw399[(new BigNumber(7)).toNumber()] = (_this).f10;
          _nw399[(new BigNumber(8)).toNumber()] = (_2403_v43).f12;
          _nw399[(new BigNumber(9)).toNumber()] = true;
          _nw399[(new BigNumber(10)).toNumber()] = (_this).f10;
          _nw399[(new BigNumber(11)).toNumber()] = (function () {
            let _coll77 = new _dafny.Set();
            for (const _compr_77 of _dafny.IntegerRange(new BigNumber(707), new BigNumber(683))) {
              let _2407_v46 = _compr_77;
              if (((new BigNumber(707)).isLessThanOrEqualTo(_2407_v46)) && ((_2407_v46).isLessThan(new BigNumber(683)))) {
                _coll77.add(_module.__default.safeModuloInt(_2407_v46, _2379_i4));
              }
            }
            return _coll77;
          }()).IsSubsetOf(_module.__default.fm25(_2401_cf81, _2400_cf82, globalState));
          _nw399[(new BigNumber(12)).toNumber()] = (_2400_cf82).isLessThanOrEqualTo(p0);
          _nw399[(new BigNumber(13)).toNumber()] = _2402_cf80;
          _nw399[(new BigNumber(14)).toNumber()] = (_2403_v43).f12;
          _nw399[(new BigNumber(15)).toNumber()] = !(_2402_cf80) || (_2402_cf80);
          _nw399[(new BigNumber(16)).toNumber()] = (_2401_cf81).isEqualTo((_this).fm1(_dafny.Map.Empty.slice().updateUnsafe((_2403_v43).f12,_2402_cf80), globalState));
          _nw399[(new BigNumber(17)).toNumber()] = (_this).f10;
          _nw399[(new BigNumber(18)).toNumber()] = _2402_cf80;
          _nw399[(new BigNumber(19)).toNumber()] = (_2403_v43).f12;
          _nw399[(new BigNumber(20)).toNumber()] = _2402_cf80;
          _nw399[(new BigNumber(21)).toNumber()] = ((_2402_cf80) ? ((_this).f10) : (!((_this).f10)));
          _nw399[(new BigNumber(22)).toNumber()] = _2402_cf80;
          _nw399[(new BigNumber(23)).toNumber()] = (_this).f10;
          _nw399[(new BigNumber(24)).toNumber()] = (_2402_cf80) || ((_this).f10);
          _2406_v47 = _nw399;
          _2406_v47 = _2406_v47;
          (globalState).f8 = _2402_cf80;
          let _2408_v48;
          _2408_v48 = _dafny.Map.Empty.slice().updateUnsafe(((!((_2403_v43).fm32(_2379_i4, _2343_v0, globalState))) ? (new BigNumber(630)) : (_2401_cf81)),_2406_v47);
          _2406_v47 = (((_2408_v48).contains(_module.__default.safeModuloInt(new BigNumber(524), p0))) ? ((_2408_v48).get(_module.__default.safeModuloInt(new BigNumber(524), p0))) : (_2406_v47));
          let _2409_v49;
          _2409_v49 = _dafny.Seq.of((_this).f10, _2402_cf80);
          let _2410_v50;
          _2410_v50 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('c'.codePointAt(0)),(_this).f10);
          let _2411_v51;
          let _nw400 = new _module.C4();
          _nw400.__ctor(_module.__default.fm48((_this).fm0(_2409_v49, (_2403_v43).f12, new BigNumber((_2410_v50).length), (_this).f10, globalState), globalState), _2402_cf80);
          _2411_v51 = _nw400;
          let _2412_v52;
          _2412_v52 = _dafny.Map.Empty.slice().updateUnsafe((_2403_v43).f13,new BigNumber((_2409_v49).length));
          let _2413_v53;
          let _nw401 = new _module.C1();
          _nw401.__ctor((_dafny.Set.fromElements(_2401_cf81, (_2381_v33)[_module.__default.safeIndex(new BigNumber((_2409_v49).length), new BigNumber((_2381_v33).length))], (_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus((((_2412_v52).contains((_2403_v43).f13)) ? ((_2412_v52).get((_2403_v43).f13)) : (p0))))))).contains(_2401_cf81), (_2403_v43).f13);
          _2413_v53 = _nw401;
          let _2414_v54;
          _2414_v54 = _dafny.Map.Empty.slice().updateUnsafe((_2403_v43).f12,_2406_v47);
          let _2415_v55;
          _2415_v55 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_2403_v43).f13).length),(_2403_v43).f13);
          let _rhs387 = _2411_v51;
          let _rhs388 = _2413_v53;
          let _rhs389 = ((_dafny.MultiSet.fromElements(_2401_cf81, _2379_i4, new BigNumber(((((_2415_v55).contains((_2381_v33)[_module.__default.safeIndex(_2400_cf82, new BigNumber((_2381_v33).length))])) ? ((_2415_v55).get((_2381_v33)[_module.__default.safeIndex(_2400_cf82, new BigNumber((_2381_v33).length))])) : ((_2403_v43).f13))).length), _2379_i4)).Union(_dafny.MultiSet.fromElements(_2379_i4))).Intersect(_module.__default.fm46(_2402_cf80, (_2403_v43).f12, !((_2403_v43).f12), globalState));
          let _rhs390 = ((((_2411_v51).fm0(_2409_v49, (_2403_v43).f12, _2379_i4, (_2403_v43).f12, globalState)) ? (_2379_i4) : (_2400_cf82))).multipliedBy(_module.__default.safeDivisionInt(p0, new BigNumber(((_2403_v43).f13).length)));
          let _rhs391 = _2414_v54;
          _2411_v51 = _rhs387;
          _2413_v53 = _rhs388;
          _2349_v6 = _rhs389;
          _2401_cf81 = _rhs390;
          _2414_v54 = _rhs391;
        } else {
          let _2416___mcc_h6 = (_source30).cf76;
          let _2417_cf76 = _2416___mcc_h6;
          let _2418_v56;
          let _nw402 = new _module.C0();
          _nw402.__ctor();
          _2418_v56 = _nw402;
          _2348_v5 = _2348_v5;
          let _2419_v57;
          let _nw403 = new _module.C0();
          _nw403.__ctor();
          _2419_v57 = _nw403;
          r2 = p0;
        }
        _2349_v6 = (_2349_v6).Difference(_2349_v6);
        if (!((_this).f10)) {
          let _2420_v58;
          _2420_v58 = _dafny.MultiSet.fromElements(_2343_v0, _2343_v0);
          let _2421_v59;
          _2421_v59 = _dafny.Seq.of(_2343_v0, _2343_v0);
          let _2422_v60;
          _2422_v60 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2420_v58);
          let _2423_v61;
          _2423_v61 = _dafny.Seq.of(_2420_v58);
          let _2424_v62;
          let _nw404 = Array((_dafny.ONE).toNumber()).fill(false);
          _2424_v62 = _nw404;
          let _2425_v63;
          _2425_v63 = _dafny.MultiSet.fromElements(_2424_v62);
          let _2426_v64;
          _2426_v64 = _2420_v58;
          let _2427_v65;
          let _nw405 = Array((new BigNumber(28)).toNumber());
          _nw405[(_dafny.ZERO).toNumber()] = (_dafny.MultiSet.fromElements(_2343_v0, _2343_v0, _2343_v0)).Difference(_2420_v58);
          _nw405[(_dafny.ONE).toNumber()] = _dafny.MultiSet.FromArray(_2421_v59);
          _nw405[(new BigNumber(2)).toNumber()] = _2420_v58;
          _nw405[(new BigNumber(3)).toNumber()] = (((_2422_v60).contains(p0)) ? ((_2422_v60).get(p0)) : ((((_2422_v60).contains(_2379_i4)) ? ((_2422_v60).get(_2379_i4)) : (_2420_v58))));
          _nw405[(new BigNumber(4)).toNumber()] = _dafny.MultiSet.fromElements(_2343_v0, _2343_v0);
          _nw405[(new BigNumber(5)).toNumber()] = _2420_v58;
          _nw405[(new BigNumber(6)).toNumber()] = _2420_v58;
          _nw405[(new BigNumber(7)).toNumber()] = _2420_v58;
          _nw405[(new BigNumber(8)).toNumber()] = (_2420_v58).update(new _dafny.CodePoint('i'.codePointAt(0)), _module.__default.abs(_2379_i4));
          _nw405[(new BigNumber(9)).toNumber()] = (_2423_v61)[_module.__default.safeIndex(new BigNumber((_2425_v63).cardinality()), new BigNumber((_2423_v61).length))];
          _nw405[(new BigNumber(10)).toNumber()] = (_2420_v58).update(_2343_v0, _module.__default.abs(new BigNumber(274)));
          _nw405[(new BigNumber(11)).toNumber()] = _2420_v58;
          _nw405[(new BigNumber(12)).toNumber()] = _2420_v58;
          _nw405[(new BigNumber(13)).toNumber()] = (_2426_v64);
          _nw405[(new BigNumber(14)).toNumber()] = (_2420_v58).Difference(_dafny.MultiSet.FromArray(_2421_v59));
          _nw405[(new BigNumber(15)).toNumber()] = _2420_v58;
          _nw405[(new BigNumber(16)).toNumber()] = (_2420_v58).Union(_dafny.MultiSet.FromArray(_2421_v59));
          _nw405[(new BigNumber(17)).toNumber()] = _2420_v58;
          _nw405[(new BigNumber(18)).toNumber()] = _2420_v58;
          _nw405[(new BigNumber(19)).toNumber()] = _2420_v58;
          _nw405[(new BigNumber(20)).toNumber()] = _2420_v58;
          _nw405[(new BigNumber(21)).toNumber()] = _2420_v58;
          _nw405[(new BigNumber(22)).toNumber()] = _2420_v58;
          _nw405[(new BigNumber(23)).toNumber()] = _2420_v58;
          _nw405[(new BigNumber(24)).toNumber()] = _2420_v58;
          _nw405[(new BigNumber(25)).toNumber()] = (_2420_v58).Union(_2420_v58);
          _nw405[(new BigNumber(26)).toNumber()] = _2420_v58;
          _nw405[(new BigNumber(27)).toNumber()] = _2420_v58;
          _2427_v65 = _nw405;
          let _2428_v66;
          _2428_v66 = _dafny.Seq.of((_this).f10);
          let _index413 = _module.__default.safeIndex(new BigNumber(999), new BigNumber((_2427_v65).length));
          (_2427_v65)[_index413] = (_2420_v58).Intersect(_module.__default.fm57(new BigNumber((_2428_v66).length), globalState));
          let _index414 = _module.__default.safeIndex(new BigNumber(999), new BigNumber((_2427_v65).length));
          (_2427_v65)[_index414] = ((_2420_v58).Intersect(_2420_v58)).Difference((((_this).f10) ? ((_2423_v61)[_module.__default.safeIndex(p0, new BigNumber((_2423_v61).length))]) : (_2420_v58)));
          let _2429_v67;
          _2429_v67 = _module.D6.create_DC16((_this).f9);
          let _2430_v68;
          let _nw406 = new _module.C5();
          _nw406.__ctor((_2429_v67).dtor_cf28, (_this).f10);
          _2430_v68 = _nw406;
          let _2431_v69;
          let _nw407 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _2431_v69 = _nw407;
          let _2432_v70;
          _2432_v70 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2421_v59);
          let _2433_v71;
          _2433_v71 = _dafny.Map.Empty.slice().updateUnsafe(_2431_v69,_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(733)), ((_2434_v0) => function (_2435_i5) {
            return _2434_v0;
          })(_2343_v0)), (((_2432_v70).contains(p0)) ? ((_2432_v70).get(p0)) : (_2421_v59))));
          let _rhs392 = (((_2433_v71).contains(_2431_v69)) ? ((_2433_v71).get(_2431_v69)) : (_dafny.Seq.UnicodeFromString("sxeiw")));
          let _rhs393 = _2430_v68;
          _2421_v59 = _rhs392;
          _2430_v68 = _rhs393;
          _2343_v0 = _2343_v0;
          let _nw408 = Array((new BigNumber(28)).toNumber()).fill(false);
          _2424_v62 = _nw408;
          r2 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(_2343_v0), _dafny.Seq.of(_2343_v0, _module.__default.fm36(globalState), _2343_v0)), _dafny.Seq.Concat(_2421_v59, _2421_v59))).length);
        } else {
          let _2436_v72;
          _2436_v72 = _dafny.Seq.UnicodeFromString("p");
          let _2437_v73;
          let _nw409 = new _module.C1();
          _nw409.__ctor((_this).f10, _2436_v72);
          _2437_v73 = _nw409;
          let _2438_v74;
          _2438_v74 = _module.D26.create_DC64(_2437_v73);
          let _2439_v75;
          let _nw410 = new _module.C3();
          _nw410.__ctor();
          _2439_v75 = _nw410;
          let _pat_let_tv53 = _2437_v73;
          let _rhs394 = _2379_i4;
          let _rhs395 = function (_pat_let56_0) {
            return function (_2440_dt__update__tmp_h0) {
              return function (_pat_let57_0) {
                return function (_2441_dt__update_hcf114_h0) {
                  return _module.D26.create_DC64(_2441_dt__update_hcf114_h0);
                }(_pat_let57_0);
              }(_pat_let_tv53);
            }(_pat_let56_0);
          }(_module.D26.create_DC64(_2437_v73));
          let _rhs396 = (_2437_v73).f12;
          let _rhs397 = ((((_this).f10) === ((_2437_v73).f12)) ? (_2439_v75) : (_2439_v75));
          let _lhs286 = globalState;
          r2 = _rhs394;
          _2438_v74 = _rhs395;
          _lhs286.f8 = _rhs396;
          _2439_v75 = _rhs397;
          r2 = p0;
          let _2442_v76;
          let _nw411 = new _module.C0();
          _nw411.__ctor();
          _2442_v76 = _nw411;
          r2 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_2379_i4, new BigNumber(((_2437_v73).f13).length)));
          let _2443_v77;
          let _init70 = ((_2444_p0) => function (_2445_i6) {
            return (_2445_i6).multipliedBy(_2444_p0);
          })(p0);
          let _nw412 = Array((new BigNumber(27)).toNumber());
          for (let _i0_70 = 0; _i0_70 < new BigNumber(_nw412.length); _i0_70++) {
            _nw412[_i0_70] = _init70(new BigNumber(_i0_70));
          }
          _2443_v77 = _nw412;
          _2443_v77 = _2443_v77;
        }
        let _2446_v79;
        _2446_v79 = _dafny.Set.fromElements(p0);
        let _2447_v80;
        _2447_v80 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_module.D28.create_DC67(_dafny.Seq.update(_dafny.Seq.of(_dafny.Set.fromElements(p0), function () {
  let _coll78 = new _dafny.Set();
  for (const _compr_78 of _dafny.IntegerRange(new BigNumber(-767), new BigNumber(994))) {
    let _2448_v78 = _compr_78;
    if (((new BigNumber(-767)).isLessThanOrEqualTo(_2448_v78)) && ((_2448_v78).isLessThan(new BigNumber(994)))) {
      _coll78.add(_module.__default.safeModuloInt(_2448_v78, _2379_i4));
    }
  }
  return _coll78;
}(), _2446_v79, _2446_v79, _2446_v79), _module.__default.safeIndex(_2379_i4, new BigNumber((_dafny.Seq.of(_dafny.Set.fromElements(p0), function () {
  let _coll79 = new _dafny.Set();
  for (const _compr_79 of _dafny.IntegerRange(new BigNumber(-767), new BigNumber(994))) {
    let _2449_v78 = _compr_79;
    if (((new BigNumber(-767)).isLessThanOrEqualTo(_2449_v78)) && ((_2449_v78).isLessThan(new BigNumber(994)))) {
      _coll79.add(_module.__default.safeModuloInt(_2449_v78, _2379_i4));
    }
  }
  return _coll79;
}(), _2446_v79, _2446_v79, _2446_v79)).length)), _2446_v79))).dtor_cf119);
        let _2450_v81;
        _2450_v81 = _dafny.Seq.of(_2446_v79);
        if (!_dafny.areEqual(_module.__default.fm58(globalState), (((_2447_v80).contains(new BigNumber(-615))) ? ((_2447_v80).get(new BigNumber(-615))) : (_2450_v81)))) {
          (globalState).f8 = (_this).f10;
          let _2451_v82;
          _2451_v82 = _dafny.Seq.of((_this).f11, (_this).f11, (_this).f11);
          let _rhs398 = true;
          let _rhs399 = (new BigNumber((_2451_v82).length)).minus(_2379_i4);
          let _rhs400 = (_this).f10;
          let _rhs401 = _2379_i4;
          let _lhs287 = globalState;
          r1 = _rhs398;
          r2 = _rhs399;
          _lhs287.f8 = _rhs400;
          r2 = _rhs401;
          let _2452_v83;
          _2452_v83 = _dafny.Seq.of((_this).f10);
          (globalState).f8 = (p0).isLessThanOrEqualTo(new BigNumber((_2452_v83).length));
          r2 = p0;
          let _2453_v84;
          let _nw413 = new _module.C1();
          _nw413.__ctor(!((_this).f10), _dafny.Seq.UnicodeFromString("ukpsmds"));
          _2453_v84 = _nw413;
          let _2454_v85;
          _2454_v85 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2453_v84);
          _2454_v85 = (_2454_v85).update(new BigNumber((_dafny.MultiSet.FromArray((((_this).f10) ? (_2452_v83) : (_2452_v83)))).cardinality()), _2453_v84);
        } else {
          r2 = (p0).multipliedBy((_this).fm1((_this).f9, globalState));
          let _2455_v86;
          _2455_v86 = _dafny.Seq.of((_this).f10);
          let _2456_v87;
          let _nw414 = new _module.C0();
          _nw414.__ctor();
          _2456_v87 = _nw414;
          let _2457_v88;
          _2457_v88 = _dafny.MultiSet.fromElements(_2456_v87);
          let _2458_v89;
          let _nw415 = Array((new BigNumber(22)).toNumber());
          _nw415[(_dafny.ZERO).toNumber()] = !((_this).f10);
          _nw415[(_dafny.ONE).toNumber()] = (_this).f10;
          _nw415[(new BigNumber(2)).toNumber()] = (_this).f10;
          _nw415[(new BigNumber(3)).toNumber()] = (_this).f10;
          _nw415[(new BigNumber(4)).toNumber()] = (_this).f10;
          _nw415[(new BigNumber(5)).toNumber()] = (_this).f10;
          _nw415[(new BigNumber(6)).toNumber()] = !(!((_this).f10)) || ((_this).f10);
          _nw415[(new BigNumber(7)).toNumber()] = (_2379_i4).isEqualTo(new BigNumber((_2455_v86).length));
          _nw415[(new BigNumber(8)).toNumber()] = (_this).f10;
          _nw415[(new BigNumber(9)).toNumber()] = !((_this).f10);
          _nw415[(new BigNumber(10)).toNumber()] = (_this).f10;
          _nw415[(new BigNumber(11)).toNumber()] = (_2457_v88).IsProperSubsetOf(_2457_v88);
          _nw415[(new BigNumber(12)).toNumber()] = (_this).f10;
          _nw415[(new BigNumber(13)).toNumber()] = !((_this).f10);
          _nw415[(new BigNumber(14)).toNumber()] = (_this).f10;
          _nw415[(new BigNumber(15)).toNumber()] = !((_this).f10);
          _nw415[(new BigNumber(16)).toNumber()] = (_this).f10;
          _nw415[(new BigNumber(17)).toNumber()] = (_this).f10;
          _nw415[(new BigNumber(18)).toNumber()] = (_this).f10;
          _nw415[(new BigNumber(19)).toNumber()] = (_2455_v86)[_module.__default.safeIndex(_2379_i4, new BigNumber((_2455_v86).length))];
          _nw415[(new BigNumber(20)).toNumber()] = (_this).f10;
          _nw415[(new BigNumber(21)).toNumber()] = (_this).f10;
          _2458_v89 = _nw415;
          let _index415 = _module.__default.safeIndex(new BigNumber(184), new BigNumber((_2458_v89).length));
          (_2458_v89)[_index415] = (_this).f10;
          let _index416 = _module.__default.safeIndex(new BigNumber(184), new BigNumber((_2458_v89).length));
          (_2458_v89)[_index416] = (_this).f10;
          (globalState).f8 = ((_2458_v89)[_module.__default.safeIndex(new BigNumber(184), new BigNumber((_2458_v89).length))]) === ((_2458_v89)[_module.__default.safeIndex(new BigNumber(184), new BigNumber((_2458_v89).length))]);
          let _2459_v90;
          _2459_v90 = _module.D7.create_DC21((_this).f11);
          let _2460_v91;
          _2460_v91 = _module.D7.create_DC21((_2459_v90).dtor_cf40);
          let _2461_v92;
          _2461_v92 = _dafny.Map.Empty.slice().updateUnsafe((((_2458_v89)[_module.__default.safeIndex(new BigNumber(184), new BigNumber((_2458_v89).length))]) ? ((_this).f10) : ((_this).f10)),((_2460_v91).dtor_cf40).IsProperSubsetOf((_this).f11));
          _2461_v92 = _2461_v92;
          let _index417 = _module.__default.safeIndex(new BigNumber(184), new BigNumber((_2458_v89).length));
          (_2458_v89)[_index417] = !((_2458_v89)[_module.__default.safeIndex(new BigNumber(184), new BigNumber((_2458_v89).length))]) || ((_this).f10);
        }
      }
      let _2462_v93;
      _2462_v93 = _dafny.Seq.of(p0, new BigNumber(-457), p0);
      let _2463_v94;
      let _nw416 = Array((new BigNumber(10)).toNumber());
      _nw416[(_dafny.ZERO).toNumber()] = new BigNumber(864);
      _nw416[(_dafny.ONE).toNumber()] = p0;
      _nw416[(new BigNumber(2)).toNumber()] = p0;
      _nw416[(new BigNumber(3)).toNumber()] = (_2462_v93)[_module.__default.safeIndex((_dafny.ZERO).minus(_module.__default.fm33(!((_this).f10), (_this).f10, p0, (_this).f10, globalState)), new BigNumber((_2462_v93).length))];
      _nw416[(new BigNumber(4)).toNumber()] = p0;
      _nw416[(new BigNumber(5)).toNumber()] = p0;
      _nw416[(new BigNumber(6)).toNumber()] = p0;
      _nw416[(new BigNumber(7)).toNumber()] = p0;
      _nw416[(new BigNumber(8)).toNumber()] = p0;
      _nw416[(new BigNumber(9)).toNumber()] = p0;
      _2463_v94 = _nw416;
      let _2464_v95;
      _2464_v95 = _dafny.MultiSet.fromElements(_2463_v94, _2463_v94, _2463_v94);
      let _hi9 = (p0).minus(p0);
      for (let _2465_i7 = (_dafny.ZERO).minus(new BigNumber((_2464_v95).cardinality())); _2465_i7.isLessThan(_hi9); _2465_i7 = _2465_i7.plus(_dafny.ONE)) {
        let _2466_v96;
        let _nw417 = new _module.C5();
        _nw417.__ctor((_dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_this).f10)).update((_this).f10, !((_this).f10)), !((_this).f10));
        _2466_v96 = _nw417;
        _2466_v96 = _2466_v96;
        let _2467_v97;
        _2467_v97 = _dafny.Seq.UnicodeFromString("ax");
        let _2468_v98;
        _2468_v98 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2467_v97).length),(_this).f10);
        let _rhs402 = (_2466_v96).fm18((_this).f11, _dafny.Seq.Concat(_2467_v97, _dafny.Seq.update(_2467_v97, _module.__default.safeIndex(p0, new BigNumber((_2467_v97).length)), _module.__default.fm31((_this).f10, _2468_v98, globalState))), globalState);
        let _rhs403 = _2465_i7;
        let _rhs404 = (_this).f10;
        let _lhs288 = globalState;
        r1 = _rhs402;
        r2 = _rhs403;
        _lhs288.f8 = _rhs404;
        let _index418 = _module.__default.safeIndex(new BigNumber(130), new BigNumber((_2463_v94).length));
        (_2463_v94)[_index418] = _2465_i7;
        let _2469_v99;
        _2469_v99 = _dafny.Map.Empty.slice().updateUnsafe((_this).f10,(_2462_v93)[_module.__default.safeIndex(p0, new BigNumber((_2462_v93).length))]);
        let _index419 = _module.__default.safeIndex(new BigNumber(130), new BigNumber((_2463_v94).length));
        (_2463_v94)[_index419] = (_2465_i7).minus((new BigNumber((_2469_v99).length)).minus((_dafny.ZERO).minus(_2465_i7)));
        let _2470_v100;
        let _nw418 = new _module.C0();
        _nw418.__ctor();
        _2470_v100 = _nw418;
      }
      let _2471_v101;
      _2471_v101 = _dafny.Seq.UnicodeFromString("gbx");
      r0 = (_module.__default.fm10((_this).f10, new BigNumber((_2462_v93).length), _2471_v101, _2350_v7, globalState)).Difference(_2349_v6);
      r1 = true;
      let _2472_v102;
      _2472_v102 = _dafny.Seq.of((_this).f10);
      r2 = new BigNumber((_dafny.Seq.update(_dafny.Seq.of((((_this).f10) ? (!((_this).f10)) : (false))), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.of((((_this).f10) ? (!((_this).f10)) : (false)))).length)), _dafny.Seq.contains(_2472_v102, !((_this).f10)))).length);
      r3 = _2471_v101;
      return [r0, r1, r2, r3];
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
